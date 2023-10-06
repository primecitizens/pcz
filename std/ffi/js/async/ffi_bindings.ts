// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

import { Application, CallbackFunc, Pointer, heap, importModule } from "@ffi";

importModule("ffi/js/async", (A: Application) => {
  interface PromiseExt {
    synchronous?: boolean;
    ok?: boolean;
    err?: boolean;
    data?: any;
  }

  const EXT_PROP = "_ext_";

  return {
    "promise": (fn: heap.Ref<CallbackFunc>): heap.Ref<Promise<any>> => {
      const executor = A.H.get<CallbackFunc>(fn);
      return A.H.push(
        new Promise<any>((resolve, reject) => {
          executor(this, resolve, reject);
        })
      );
    },
    "bg": (ref: heap.Ref<Promise<any>>): heap.Ref<boolean> => {
      const p = A.H.get<Promise<any>>(ref);
      if (!p) {
        return A.H.FALSE;
      }

      if (typeof p[EXT_PROP] === "object") {
        return A.H.TRUE;
      }

      const ext: PromiseExt = {
        ok: false,
        err: false,
        data: undefined,
      };

      A.H.replace(
        ref,
        Object.defineProperty(
          p.then(
            (value: any): any => {
              if (!ext.err && typeof ext.data === "undefined") {
                ext.data = value;
                ext.ok = true;
              }
              return value;
            },
            (reason: any): any => {
              if (!ext.ok && typeof ext.data === "undefined") {
                ext.data = reason;
                ext.err = true;
              }
              return reason;
            }
          ),
          EXT_PROP,
          {
            get: (): PromiseExt => {
              return ext;
            },
          }
        )
      );

      return A.H.TRUE;
    },
    "resolved": (take: heap.Ref<boolean>, data: heap.Ref<any>): heap.Ref<Promise<any>> => {
      return A.H.push(Promise.resolve(take === A.H.TRUE ? A.H.take(data) : A.H.get(data)));
    },
    "rejected": (take: heap.Ref<boolean>, reason: heap.Ref<any>): heap.Ref<Promise<any>> => {
      return A.H.push(Promise.reject(take === A.H.TRUE ? A.H.take(reason) : A.H.get(reason)));
    },
    "allok": (take: heap.Ref<boolean>, ptr: Pointer, n: number): heap.Ref<Promise<any[]>> => {
      return A.H.push(Promise.all(A.load.Refs(ptr, n, take === A.H.TRUE)));
    },
    "all": (free: heap.Ref<boolean>, ptr: Pointer, n: number): heap.Ref<Promise<PromiseSettledResult<any>[]>> => {
      return A.H.push(Promise.allSettled<any[]>(A.load.Refs(ptr, n, free === A.H.TRUE)));
    },
    "first": (free: heap.Ref<boolean>, ptr: Pointer, n: number): heap.Ref<Promise<any>> => {
      return A.H.push(Promise.race(A.load.Refs(ptr, n, free === A.H.TRUE)));
    },
    "then": (
      p: heap.Ref<Promise<any>>,
      onfulfilled: heap.Ref<(value: any) => any>,
      onrejected: heap.Ref<(reason: any) => any>,
      onfinally: heap.Ref<() => void>
    ): void => {
      let x = A.H.get<Promise<any>>(p).then(
        onfulfilled !== A.H.UNDEFINED ? A.H.get(onfulfilled) : null,
        onrejected !== A.H.UNDEFINED ? A.H.get(onrejected) : null
      );
      if (onfinally !== A.H.UNDEFINED) {
        // available since ES2018, supported by all browsers but IE and Opera Mini.
        x = x.finally(A.H.get(onfinally));
      }
      A.H.replace(p, x);
    },
    "shouldwait": (
      take: heap.Ref<boolean>,
      ref: heap.Ref<Promise<any>>,
      pFulfilled: Pointer,
      pRef: Pointer
    ): heap.Ref<boolean> => {
      const promise = take === A.H.TRUE ? A.H.take<Promise<any>>(ref) : A.H.get<Promise<any>>(ref);
      if (!promise) {
        return A.H.FALSE;
      }

      let ext = promise[EXT_PROP] as PromiseExt;
      if (typeof ext !== "undefined") {
        // is a background promise, check whether finished
        if (ext.ok) {
          A.store.Bool(pFulfilled, true);
          A.store.Ref(pRef, ext.data);
          return A.H.FALSE; // do not wait and do not call "takewaited"
        }

        if (ext.err) {
          A.store.Ref(pRef, ext.data);
          return A.H.FALSE; // do not wait
        }
      }

      // is a plain/unfinished promise so we have no idea whether
      // the promise has been resolved, attach callbacks to check.
      ext = ext ? ext : {};

      // defensive: ensure not waiting for promise callback that has
      // fired synchronously
      ext.synchronous = true;

      promise.then(
        (value: any): any => {
          ext.data = value;
          ext.ok = true;
          if (ext.synchronous) {
            ext.synchronous = false;
          } else {
            A.call(0, 0);
          }
          return value;
        },
        (reason: any): any => {
          ext.data = reason;
          ext.err = true;

          if (ext.synchronous) {
            ext.synchronous = false;
          } else {
            A.call(0, 0);
          }
          return reason;
        }
      );
      if (ext.synchronous) {
        A.store.Ref(pRef, ext);
        ext.synchronous = false;
        return A.H.TRUE; // need to wait and call "takewaited"
      }

      if (ext.ok) {
        A.store.Bool(pFulfilled, true);
        A.store.Ref(pRef, ext.data);
        return A.H.FALSE; // do not wait and do not call "takewaited"
      }

      if (ext.err) {
        A.store.Ref(pRef, ext.data);
        return A.H.FALSE; // do not wait and do not call "takewaited"
      }

      throw new Error("unreachable");
    },
    "takewaited": (ref: heap.Ref<PromiseExt>, pDataRef: Pointer): heap.Ref<boolean> => {
      const ext = A.H.take<PromiseExt>(ref);
      if (!ext.ok && !ext.err) {
        throw new Error("unexpected state: promise unfinished");
      }

      A.store.Ref(pDataRef, ext.data);
      return ext.ok ? A.H.TRUE : A.H.FALSE;
    },
  };
});
