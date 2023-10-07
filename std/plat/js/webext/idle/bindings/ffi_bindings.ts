import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/idle", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_IdleState": (ref: heap.Ref<string>): number => {
      const idx = ["active", "idle", "locked"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_GetAutoLockDelay": (): heap.Ref<boolean> => {
      if (WEBEXT?.idle && "getAutoLockDelay" in WEBEXT?.idle) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAutoLockDelay": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.idle.getAutoLockDelay);
    },
    "call_GetAutoLockDelay": (retPtr: Pointer): void => {
      const _ret = WEBEXT.idle.getAutoLockDelay();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAutoLockDelay": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.idle.getAutoLockDelay();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnStateChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.idle?.onStateChanged && "addListener" in WEBEXT?.idle?.onStateChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.idle.onStateChanged.addListener);
    },
    "call_OnStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.idle.onStateChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnStateChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.idle.onStateChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffStateChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.idle?.onStateChanged && "removeListener" in WEBEXT?.idle?.onStateChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.idle.onStateChanged.removeListener);
    },
    "call_OffStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.idle.onStateChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffStateChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.idle.onStateChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnStateChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.idle?.onStateChanged && "hasListener" in WEBEXT?.idle?.onStateChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.idle.onStateChanged.hasListener);
    },
    "call_HasOnStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.idle.onStateChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnStateChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.idle.onStateChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_QueryState": (): heap.Ref<boolean> => {
      if (WEBEXT?.idle && "queryState" in WEBEXT?.idle) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_QueryState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.idle.queryState);
    },
    "call_QueryState": (retPtr: Pointer, detectionIntervalInSeconds: number): void => {
      const _ret = WEBEXT.idle.queryState(detectionIntervalInSeconds);
      A.store.Ref(retPtr, _ret);
    },
    "try_QueryState": (retPtr: Pointer, errPtr: Pointer, detectionIntervalInSeconds: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.idle.queryState(detectionIntervalInSeconds);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetDetectionInterval": (): heap.Ref<boolean> => {
      if (WEBEXT?.idle && "setDetectionInterval" in WEBEXT?.idle) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetDetectionInterval": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.idle.setDetectionInterval);
    },
    "call_SetDetectionInterval": (retPtr: Pointer, intervalInSeconds: number): void => {
      const _ret = WEBEXT.idle.setDetectionInterval(intervalInSeconds);
    },
    "try_SetDetectionInterval": (retPtr: Pointer, errPtr: Pointer, intervalInSeconds: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.idle.setDetectionInterval(intervalInSeconds);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
