import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/pagecapture", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_SaveAsMHTMLArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Int64(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_SaveAsMHTMLArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["tabId"] = A.load.Int64(ptr + 0);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_SaveAsMHTML": (): heap.Ref<boolean> => {
      if (WEBEXT?.pageCapture && "saveAsMHTML" in WEBEXT?.pageCapture) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SaveAsMHTML": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pageCapture.saveAsMHTML);
    },
    "call_SaveAsMHTML": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["tabId"] = A.load.Int64(details + 0);

      const _ret = WEBEXT.pageCapture.saveAsMHTML(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SaveAsMHTML": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["tabId"] = A.load.Int64(details + 0);

        const _ret = WEBEXT.pageCapture.saveAsMHTML(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
