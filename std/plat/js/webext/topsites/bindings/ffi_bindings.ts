import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/topsites", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_MostVisitedURL": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["title"]);
        A.store.Ref(ptr + 4, x["url"]);
      }
    },
    "load_MostVisitedURL": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["title"] = A.load.Ref(ptr + 0, undefined);
      x["url"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Get": (): heap.Ref<boolean> => {
      if (WEBEXT?.topSites && "get" in WEBEXT?.topSites) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Get": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.topSites.get);
    },
    "call_Get": (retPtr: Pointer): void => {
      const _ret = WEBEXT.topSites.get();
      A.store.Ref(retPtr, _ret);
    },
    "try_Get": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.topSites.get();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
