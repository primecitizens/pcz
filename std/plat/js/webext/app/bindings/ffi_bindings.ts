import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/app", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_DOMWindow": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_DOMWindow": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Details": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_Details": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_InstallStateType": (ref: heap.Ref<string>): number => {
      const idx = ["not_installed", "installed", "disabled"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_RunningStateType": (ref: heap.Ref<string>): number => {
      const idx = ["running", "cannot_run", "ready_to_run"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_GetDetails": (): heap.Ref<boolean> => {
      if (WEBEXT?.app && "getDetails" in WEBEXT?.app) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDetails": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.getDetails);
    },
    "call_GetDetails": (retPtr: Pointer): void => {
      const _ret = WEBEXT.app.getDetails();
      if (typeof _ret === "undefined") {
        A.store.Bool(retPtr + 0, false);
      } else {
        A.store.Bool(retPtr + 0, true);
      }
    },
    "try_GetDetails": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.getDetails();
        if (typeof _ret === "undefined") {
          A.store.Bool(retPtr + 0, false);
        } else {
          A.store.Bool(retPtr + 0, true);
        }
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetIsInstalled": (): heap.Ref<boolean> => {
      if (WEBEXT?.app && "getIsInstalled" in WEBEXT?.app) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetIsInstalled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.getIsInstalled);
    },
    "call_GetIsInstalled": (retPtr: Pointer): void => {
      const _ret = WEBEXT.app.getIsInstalled();
      A.store.Bool(retPtr, _ret);
    },
    "try_GetIsInstalled": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.getIsInstalled();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InstallState": (): heap.Ref<boolean> => {
      if (WEBEXT?.app && "installState" in WEBEXT?.app) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InstallState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.installState);
    },
    "call_InstallState": (retPtr: Pointer): void => {
      const _ret = WEBEXT.app.installState();
      A.store.Ref(retPtr, _ret);
    },
    "try_InstallState": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.installState();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunningState": (): heap.Ref<boolean> => {
      if (WEBEXT?.app && "runningState" in WEBEXT?.app) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunningState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.runningState);
    },
    "call_RunningState": (retPtr: Pointer): void => {
      const _ret = WEBEXT.app.runningState();
      A.store.Enum(retPtr, ["running", "cannot_run", "ready_to_run"].indexOf(_ret as string));
    },
    "try_RunningState": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.runningState();
        A.store.Enum(retPtr, ["running", "cannot_run", "ready_to_run"].indexOf(_ret as string));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
