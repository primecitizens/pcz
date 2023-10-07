import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/systemprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_GetIncognitoModeAvailabilityValue": (ref: heap.Ref<string>): number => {
      const idx = ["enabled", "disabled", "forced"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_UpdateStatusState": (ref: heap.Ref<string>): number => {
      const idx = ["NotAvailable", "Updating", "NeedRestart"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_UpdateStatus": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Float64(ptr + 0, x["downloadProgress"] === undefined ? 0 : (x["downloadProgress"] as number));
        A.store.Enum(ptr + 8, ["NotAvailable", "Updating", "NeedRestart"].indexOf(x["state"] as string));
      }
    },
    "load_UpdateStatus": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["downloadProgress"] = A.load.Float64(ptr + 0);
      x["state"] = A.load.Enum(ptr + 8, ["NotAvailable", "Updating", "NeedRestart"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetApiKey": (): heap.Ref<boolean> => {
      if (WEBEXT?.systemPrivate && "getApiKey" in WEBEXT?.systemPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetApiKey": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.systemPrivate.getApiKey);
    },
    "call_GetApiKey": (retPtr: Pointer): void => {
      const _ret = WEBEXT.systemPrivate.getApiKey();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetApiKey": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.systemPrivate.getApiKey();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetIncognitoModeAvailability": (): heap.Ref<boolean> => {
      if (WEBEXT?.systemPrivate && "getIncognitoModeAvailability" in WEBEXT?.systemPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetIncognitoModeAvailability": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.systemPrivate.getIncognitoModeAvailability);
    },
    "call_GetIncognitoModeAvailability": (retPtr: Pointer): void => {
      const _ret = WEBEXT.systemPrivate.getIncognitoModeAvailability();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetIncognitoModeAvailability": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.systemPrivate.getIncognitoModeAvailability();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetUpdateStatus": (): heap.Ref<boolean> => {
      if (WEBEXT?.systemPrivate && "getUpdateStatus" in WEBEXT?.systemPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetUpdateStatus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.systemPrivate.getUpdateStatus);
    },
    "call_GetUpdateStatus": (retPtr: Pointer): void => {
      const _ret = WEBEXT.systemPrivate.getUpdateStatus();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetUpdateStatus": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.systemPrivate.getUpdateStatus();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
