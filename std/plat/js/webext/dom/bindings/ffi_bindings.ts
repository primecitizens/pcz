import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/dom", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_OpenOrClosedShadowRoot": (): heap.Ref<boolean> => {
      if (WEBEXT?.dom && "openOrClosedShadowRoot" in WEBEXT?.dom) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenOrClosedShadowRoot": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.dom.openOrClosedShadowRoot);
    },
    "call_OpenOrClosedShadowRoot": (retPtr: Pointer, element: heap.Ref<object>): void => {
      const _ret = WEBEXT.dom.openOrClosedShadowRoot(A.H.get<object>(element));
      A.store.Ref(retPtr, _ret);
    },
    "try_OpenOrClosedShadowRoot": (retPtr: Pointer, errPtr: Pointer, element: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.dom.openOrClosedShadowRoot(A.H.get<object>(element));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
