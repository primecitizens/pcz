import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/mojoprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_RequireAsync": (): heap.Ref<boolean> => {
      if (WEBEXT?.mojoPrivate && "requireAsync" in WEBEXT?.mojoPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RequireAsync": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mojoPrivate.requireAsync);
    },
    "call_RequireAsync": (retPtr: Pointer, name: heap.Ref<object>): void => {
      const _ret = WEBEXT.mojoPrivate.requireAsync(A.H.get<object>(name));
      A.store.Ref(retPtr, _ret);
    },
    "try_RequireAsync": (retPtr: Pointer, errPtr: Pointer, name: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mojoPrivate.requireAsync(A.H.get<object>(name));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
