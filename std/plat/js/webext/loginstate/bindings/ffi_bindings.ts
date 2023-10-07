import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/loginstate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_ProfileType": (ref: heap.Ref<string>): number => {
      const idx = ["SIGNIN_PROFILE", "USER_PROFILE"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SessionState": (ref: heap.Ref<string>): number => {
      const idx = [
        "UNKNOWN",
        "IN_OOBE_SCREEN",
        "IN_LOGIN_SCREEN",
        "IN_SESSION",
        "IN_LOCK_SCREEN",
        "IN_RMA_SCREEN",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_GetProfileType": (): heap.Ref<boolean> => {
      if (WEBEXT?.loginState && "getProfileType" in WEBEXT?.loginState) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetProfileType": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.loginState.getProfileType);
    },
    "call_GetProfileType": (retPtr: Pointer): void => {
      const _ret = WEBEXT.loginState.getProfileType();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetProfileType": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.loginState.getProfileType();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSessionState": (): heap.Ref<boolean> => {
      if (WEBEXT?.loginState && "getSessionState" in WEBEXT?.loginState) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSessionState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.loginState.getSessionState);
    },
    "call_GetSessionState": (retPtr: Pointer): void => {
      const _ret = WEBEXT.loginState.getSessionState();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSessionState": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.loginState.getSessionState();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSessionStateChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.loginState?.onSessionStateChanged && "addListener" in WEBEXT?.loginState?.onSessionStateChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSessionStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.loginState.onSessionStateChanged.addListener);
    },
    "call_OnSessionStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.loginState.onSessionStateChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnSessionStateChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.loginState.onSessionStateChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSessionStateChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.loginState?.onSessionStateChanged && "removeListener" in WEBEXT?.loginState?.onSessionStateChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSessionStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.loginState.onSessionStateChanged.removeListener);
    },
    "call_OffSessionStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.loginState.onSessionStateChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffSessionStateChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.loginState.onSessionStateChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSessionStateChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.loginState?.onSessionStateChanged && "hasListener" in WEBEXT?.loginState?.onSessionStateChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSessionStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.loginState.onSessionStateChanged.hasListener);
    },
    "call_HasOnSessionStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.loginState.onSessionStateChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSessionStateChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.loginState.onSessionStateChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
