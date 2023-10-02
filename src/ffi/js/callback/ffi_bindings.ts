import { Application, Pointer, heap, importModule } from "@ffi";

importModule("ffi/js/callback", (A: Application) => {
  type ReturnType = any | PromiseLike<any>;

  interface CallbackContext {
    dispFnPC: number;
    targetPC: number;
    dispRecv: number;
    retRef: heap.Ref<ReturnType>;
    args: heap.Ref<any>[];
    resolveFn: heap.Ref<() => void>;
  }

  const RET_PLACEHOLDER = "_CBRET_";
  const PROMISE_PLACEHOLDER = "_CBPROM_";

  return {
    "func": (dispFnPC: number, dispRecv: number, targetPC: number): heap.Ref<Function> => {
      dispFnPC >>>= 0;
      dispRecv >>>= 0;
      targetPC >>>= 0;
      const runid = A.runid;

      // NOTE: do not use arrow function, whose `this` would always be the Application.
      return A.H.push(function (): ReturnType {
        // NOTE: on calling go function, the app could have exited,
        // in that case, the app may has something like active `setInterval` routine
        // not cancelled.
        if (A.exited || runid != A.runid) {
          throw new Error(`unexpected function call after program exit: pc=${targetPC}`);
        }

        if (arguments.length > 15) {
          throw new Error("too many args to callback function");
        }

        const heapArgs = new Array(arguments.length + 1);
        heapArgs[0] = A.H.push(this);
        for (let i = 0; i < arguments.length; i++) {
          heapArgs[i + 1] = A.H.push(arguments[i]);
        }

        const context: CallbackContext = {
          dispFnPC,
          dispRecv,
          targetPC,
          args: heapArgs,
          retRef: A.H.push(RET_PLACEHOLDER), // cannot push the `undefined` value
          resolveFn: A.H.push(PROMISE_PLACEHOLDER), // see below
        };
        A.H.replace(context.retRef, undefined);

        let resolved = false;
        // the callback may return due to awaiting,
        // so we have to wait for the actual return in a promise.
        const p: Promise<ReturnType> = new Promise<boolean>((resolve: (value: boolean) => void) => {
          A.H.replace(context.resolveFn, () => {
            resolved = true;
            A.H.free(context.resolveFn);
            resolve(true);
          });
        }).then(() => {
          heapArgs.forEach((x) => A.H.free(x));
          if (A.exited) {
            A.resolveExitPromise();
          }
          return A.H.take(context.retRef);
        });

        A.call(A.H.push(context), 0);

        // The callback may also return synchronously, in that case we
        // should always return the value;
        //
        // resolved can only be true when the callbacked returned synchronously.
        if (resolved) {
          return A.H.get(context.retRef);
        }

        return p;
      });
    },
    "ctx": (ref: heap.Ref<CallbackContext>, pCallbackContext: Pointer): void => {
      const cb = A.H.take<CallbackContext>(ref);

      // offsets MUST agree with ../callback.go#type:CallbackContext
      pCallbackContext >>>= 0;
      A.store.Uint32(pCallbackContext + 0, cb.dispFnPC);
      A.store.Uint32(pCallbackContext + 4, cb.dispRecv);
      A.store.Uint32(pCallbackContext + 8, cb.targetPC);
      A.store.Uint32(pCallbackContext + 12, cb.retRef);
      A.store.Uint32(pCallbackContext + 16, cb.resolveFn);
      if (!cb.args || cb.args.length === 0) {
        A.store.Uint32(pCallbackContext + 20, 0);
        return;
      }

      A.store.Uint32(pCallbackContext + 20, cb.args.length);
      pCallbackContext += 24;
      // at most 15 args as guarded by the 'func'
      for (let i = 0; i < cb.args.length; i++) {
        A.store.Uint32(pCallbackContext + i * 4, cb.args[i]);
      }
    },
  };
});
