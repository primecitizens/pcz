import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/offscreen", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_Reason": (ref: heap.Ref<string>): number => {
      const idx = [
        "TESTING",
        "AUDIO_PLAYBACK",
        "IFRAME_SCRIPTING",
        "DOM_SCRAPING",
        "BLOBS",
        "DOM_PARSER",
        "USER_MEDIA",
        "DISPLAY_MEDIA",
        "WEB_RTC",
        "CLIPBOARD",
        "LOCAL_STORAGE",
        "WORKERS",
        "BATTERY_STATUS",
        "MATCH_MEDIA",
        "GEOLOCATION",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_CreateParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["reasons"]);
        A.store.Ref(ptr + 4, x["url"]);
        A.store.Ref(ptr + 8, x["justification"]);
      }
    },
    "load_CreateParameters": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["reasons"] = A.load.Ref(ptr + 0, undefined);
      x["url"] = A.load.Ref(ptr + 4, undefined);
      x["justification"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_CloseDocument": (): heap.Ref<boolean> => {
      if (WEBEXT?.offscreen && "closeDocument" in WEBEXT?.offscreen) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CloseDocument": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.offscreen.closeDocument);
    },
    "call_CloseDocument": (retPtr: Pointer): void => {
      const _ret = WEBEXT.offscreen.closeDocument();
      A.store.Ref(retPtr, _ret);
    },
    "try_CloseDocument": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.offscreen.closeDocument();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CreateDocument": (): heap.Ref<boolean> => {
      if (WEBEXT?.offscreen && "createDocument" in WEBEXT?.offscreen) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CreateDocument": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.offscreen.createDocument);
    },
    "call_CreateDocument": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["reasons"] = A.load.Ref(parameters + 0, undefined);
      parameters_ffi["url"] = A.load.Ref(parameters + 4, undefined);
      parameters_ffi["justification"] = A.load.Ref(parameters + 8, undefined);

      const _ret = WEBEXT.offscreen.createDocument(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_CreateDocument": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["reasons"] = A.load.Ref(parameters + 0, undefined);
        parameters_ffi["url"] = A.load.Ref(parameters + 4, undefined);
        parameters_ffi["justification"] = A.load.Ref(parameters + 8, undefined);

        const _ret = WEBEXT.offscreen.createDocument(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasDocument": (): heap.Ref<boolean> => {
      if (WEBEXT?.offscreen && "hasDocument" in WEBEXT?.offscreen) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasDocument": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.offscreen.hasDocument);
    },
    "call_HasDocument": (retPtr: Pointer): void => {
      const _ret = WEBEXT.offscreen.hasDocument();
      A.store.Ref(retPtr, _ret);
    },
    "try_HasDocument": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.offscreen.hasDocument();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
