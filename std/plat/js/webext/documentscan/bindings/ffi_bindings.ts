import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/documentscan", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_ScanResults": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["dataUrls"]);
        A.store.Ref(ptr + 4, x["mimeType"]);
      }
    },
    "load_ScanResults": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["dataUrls"] = A.load.Ref(ptr + 0, undefined);
      x["mimeType"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ScanOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Ref(ptr + 0, x["mimeTypes"]);
        A.store.Bool(ptr + 8, "maxImages" in x ? true : false);
        A.store.Int32(ptr + 4, x["maxImages"] === undefined ? 0 : (x["maxImages"] as number));
      }
    },
    "load_ScanOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["mimeTypes"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 8)) {
        x["maxImages"] = A.load.Int32(ptr + 4);
      } else {
        delete x["maxImages"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Scan": (): heap.Ref<boolean> => {
      if (WEBEXT?.documentScan && "scan" in WEBEXT?.documentScan) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Scan": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.documentScan.scan);
    },
    "call_Scan": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["mimeTypes"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 8)) {
        options_ffi["maxImages"] = A.load.Int32(options + 4);
      }

      const _ret = WEBEXT.documentScan.scan(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Scan": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["mimeTypes"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 8)) {
          options_ffi["maxImages"] = A.load.Int32(options + 4);
        }

        const _ret = WEBEXT.documentScan.scan(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
