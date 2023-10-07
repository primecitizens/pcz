import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/cecprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_DisplayCecPowerState": (ref: heap.Ref<string>): number => {
      const idx = [
        "error",
        "adapterNotConfigured",
        "noDevice",
        "on",
        "standby",
        "transitioningToOn",
        "transitioningToStandby",
        "unknown",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_QueryDisplayCecPowerState": (): heap.Ref<boolean> => {
      if (WEBEXT?.cecPrivate && "queryDisplayCecPowerState" in WEBEXT?.cecPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_QueryDisplayCecPowerState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.cecPrivate.queryDisplayCecPowerState);
    },
    "call_QueryDisplayCecPowerState": (retPtr: Pointer): void => {
      const _ret = WEBEXT.cecPrivate.queryDisplayCecPowerState();
      A.store.Ref(retPtr, _ret);
    },
    "try_QueryDisplayCecPowerState": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.cecPrivate.queryDisplayCecPowerState();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendStandBy": (): heap.Ref<boolean> => {
      if (WEBEXT?.cecPrivate && "sendStandBy" in WEBEXT?.cecPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendStandBy": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.cecPrivate.sendStandBy);
    },
    "call_SendStandBy": (retPtr: Pointer): void => {
      const _ret = WEBEXT.cecPrivate.sendStandBy();
      A.store.Ref(retPtr, _ret);
    },
    "try_SendStandBy": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.cecPrivate.sendStandBy();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendWakeUp": (): heap.Ref<boolean> => {
      if (WEBEXT?.cecPrivate && "sendWakeUp" in WEBEXT?.cecPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendWakeUp": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.cecPrivate.sendWakeUp);
    },
    "call_SendWakeUp": (retPtr: Pointer): void => {
      const _ret = WEBEXT.cecPrivate.sendWakeUp();
      A.store.Ref(retPtr, _ret);
    },
    "try_SendWakeUp": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.cecPrivate.sendWakeUp();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
