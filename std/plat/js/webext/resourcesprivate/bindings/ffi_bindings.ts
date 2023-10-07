import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/resourcesprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_Component": (ref: heap.Ref<string>): number => {
      const idx = ["identity", "pdf"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_GetStrings": (): heap.Ref<boolean> => {
      if (WEBEXT?.resourcesPrivate && "getStrings" in WEBEXT?.resourcesPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetStrings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.resourcesPrivate.getStrings);
    },
    "call_GetStrings": (retPtr: Pointer, component: number): void => {
      const _ret = WEBEXT.resourcesPrivate.getStrings(
        component > 0 && component <= 2 ? ["identity", "pdf"][component - 1] : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_GetStrings": (retPtr: Pointer, errPtr: Pointer, component: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.resourcesPrivate.getStrings(
          component > 0 && component <= 2 ? ["identity", "pdf"][component - 1] : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
