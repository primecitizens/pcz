import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/login", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_SamlUserSessionProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["email"]);
        A.store.Ref(ptr + 4, x["gaiaId"]);
        A.store.Ref(ptr + 8, x["password"]);
        A.store.Ref(ptr + 12, x["oauthCode"]);
      }
    },
    "load_SamlUserSessionProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["email"] = A.load.Ref(ptr + 0, undefined);
      x["gaiaId"] = A.load.Ref(ptr + 4, undefined);
      x["password"] = A.load.Ref(ptr + 8, undefined);
      x["oauthCode"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_EndSharedSession": (): heap.Ref<boolean> => {
      if (WEBEXT?.login && "endSharedSession" in WEBEXT?.login) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EndSharedSession": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.endSharedSession);
    },
    "call_EndSharedSession": (retPtr: Pointer): void => {
      const _ret = WEBEXT.login.endSharedSession();
      A.store.Ref(retPtr, _ret);
    },
    "try_EndSharedSession": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.endSharedSession();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_EnterSharedSession": (): heap.Ref<boolean> => {
      if (WEBEXT?.login && "enterSharedSession" in WEBEXT?.login) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EnterSharedSession": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.enterSharedSession);
    },
    "call_EnterSharedSession": (retPtr: Pointer, password: heap.Ref<object>): void => {
      const _ret = WEBEXT.login.enterSharedSession(A.H.get<object>(password));
      A.store.Ref(retPtr, _ret);
    },
    "try_EnterSharedSession": (retPtr: Pointer, errPtr: Pointer, password: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.enterSharedSession(A.H.get<object>(password));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ExitCurrentSession": (): heap.Ref<boolean> => {
      if (WEBEXT?.login && "exitCurrentSession" in WEBEXT?.login) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ExitCurrentSession": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.exitCurrentSession);
    },
    "call_ExitCurrentSession": (retPtr: Pointer, dataForNextLoginAttempt: heap.Ref<object>): void => {
      const _ret = WEBEXT.login.exitCurrentSession(A.H.get<object>(dataForNextLoginAttempt));
      A.store.Ref(retPtr, _ret);
    },
    "try_ExitCurrentSession": (
      retPtr: Pointer,
      errPtr: Pointer,
      dataForNextLoginAttempt: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.exitCurrentSession(A.H.get<object>(dataForNextLoginAttempt));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_FetchDataForNextLoginAttempt": (): heap.Ref<boolean> => {
      if (WEBEXT?.login && "fetchDataForNextLoginAttempt" in WEBEXT?.login) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_FetchDataForNextLoginAttempt": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.fetchDataForNextLoginAttempt);
    },
    "call_FetchDataForNextLoginAttempt": (retPtr: Pointer): void => {
      const _ret = WEBEXT.login.fetchDataForNextLoginAttempt();
      A.store.Ref(retPtr, _ret);
    },
    "try_FetchDataForNextLoginAttempt": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.fetchDataForNextLoginAttempt();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LaunchManagedGuestSession": (): heap.Ref<boolean> => {
      if (WEBEXT?.login && "launchManagedGuestSession" in WEBEXT?.login) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LaunchManagedGuestSession": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.launchManagedGuestSession);
    },
    "call_LaunchManagedGuestSession": (retPtr: Pointer, password: heap.Ref<object>): void => {
      const _ret = WEBEXT.login.launchManagedGuestSession(A.H.get<object>(password));
      A.store.Ref(retPtr, _ret);
    },
    "try_LaunchManagedGuestSession": (
      retPtr: Pointer,
      errPtr: Pointer,
      password: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.launchManagedGuestSession(A.H.get<object>(password));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LaunchSamlUserSession": (): heap.Ref<boolean> => {
      if (WEBEXT?.login && "launchSamlUserSession" in WEBEXT?.login) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LaunchSamlUserSession": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.launchSamlUserSession);
    },
    "call_LaunchSamlUserSession": (retPtr: Pointer, properties: Pointer): void => {
      const properties_ffi = {};

      properties_ffi["email"] = A.load.Ref(properties + 0, undefined);
      properties_ffi["gaiaId"] = A.load.Ref(properties + 4, undefined);
      properties_ffi["password"] = A.load.Ref(properties + 8, undefined);
      properties_ffi["oauthCode"] = A.load.Ref(properties + 12, undefined);

      const _ret = WEBEXT.login.launchSamlUserSession(properties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_LaunchSamlUserSession": (retPtr: Pointer, errPtr: Pointer, properties: Pointer): heap.Ref<boolean> => {
      try {
        const properties_ffi = {};

        properties_ffi["email"] = A.load.Ref(properties + 0, undefined);
        properties_ffi["gaiaId"] = A.load.Ref(properties + 4, undefined);
        properties_ffi["password"] = A.load.Ref(properties + 8, undefined);
        properties_ffi["oauthCode"] = A.load.Ref(properties + 12, undefined);

        const _ret = WEBEXT.login.launchSamlUserSession(properties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LaunchSharedManagedGuestSession": (): heap.Ref<boolean> => {
      if (WEBEXT?.login && "launchSharedManagedGuestSession" in WEBEXT?.login) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LaunchSharedManagedGuestSession": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.launchSharedManagedGuestSession);
    },
    "call_LaunchSharedManagedGuestSession": (retPtr: Pointer, password: heap.Ref<object>): void => {
      const _ret = WEBEXT.login.launchSharedManagedGuestSession(A.H.get<object>(password));
      A.store.Ref(retPtr, _ret);
    },
    "try_LaunchSharedManagedGuestSession": (
      retPtr: Pointer,
      errPtr: Pointer,
      password: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.launchSharedManagedGuestSession(A.H.get<object>(password));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LockCurrentSession": (): heap.Ref<boolean> => {
      if (WEBEXT?.login && "lockCurrentSession" in WEBEXT?.login) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LockCurrentSession": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.lockCurrentSession);
    },
    "call_LockCurrentSession": (retPtr: Pointer): void => {
      const _ret = WEBEXT.login.lockCurrentSession();
      A.store.Ref(retPtr, _ret);
    },
    "try_LockCurrentSession": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.lockCurrentSession();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LockManagedGuestSession": (): heap.Ref<boolean> => {
      if (WEBEXT?.login && "lockManagedGuestSession" in WEBEXT?.login) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LockManagedGuestSession": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.lockManagedGuestSession);
    },
    "call_LockManagedGuestSession": (retPtr: Pointer): void => {
      const _ret = WEBEXT.login.lockManagedGuestSession();
      A.store.Ref(retPtr, _ret);
    },
    "try_LockManagedGuestSession": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.lockManagedGuestSession();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_NotifyExternalLogoutDone": (): heap.Ref<boolean> => {
      if (WEBEXT?.login && "notifyExternalLogoutDone" in WEBEXT?.login) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_NotifyExternalLogoutDone": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.notifyExternalLogoutDone);
    },
    "call_NotifyExternalLogoutDone": (retPtr: Pointer): void => {
      const _ret = WEBEXT.login.notifyExternalLogoutDone();
      A.store.Ref(retPtr, _ret);
    },
    "try_NotifyExternalLogoutDone": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.notifyExternalLogoutDone();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnExternalLogoutDone": (): heap.Ref<boolean> => {
      if (WEBEXT?.login?.onExternalLogoutDone && "addListener" in WEBEXT?.login?.onExternalLogoutDone) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnExternalLogoutDone": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.onExternalLogoutDone.addListener);
    },
    "call_OnExternalLogoutDone": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.login.onExternalLogoutDone.addListener(A.H.get<object>(callback));
    },
    "try_OnExternalLogoutDone": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.onExternalLogoutDone.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffExternalLogoutDone": (): heap.Ref<boolean> => {
      if (WEBEXT?.login?.onExternalLogoutDone && "removeListener" in WEBEXT?.login?.onExternalLogoutDone) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffExternalLogoutDone": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.onExternalLogoutDone.removeListener);
    },
    "call_OffExternalLogoutDone": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.login.onExternalLogoutDone.removeListener(A.H.get<object>(callback));
    },
    "try_OffExternalLogoutDone": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.onExternalLogoutDone.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnExternalLogoutDone": (): heap.Ref<boolean> => {
      if (WEBEXT?.login?.onExternalLogoutDone && "hasListener" in WEBEXT?.login?.onExternalLogoutDone) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnExternalLogoutDone": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.onExternalLogoutDone.hasListener);
    },
    "call_HasOnExternalLogoutDone": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.login.onExternalLogoutDone.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnExternalLogoutDone": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.onExternalLogoutDone.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRequestExternalLogout": (): heap.Ref<boolean> => {
      if (WEBEXT?.login?.onRequestExternalLogout && "addListener" in WEBEXT?.login?.onRequestExternalLogout) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRequestExternalLogout": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.onRequestExternalLogout.addListener);
    },
    "call_OnRequestExternalLogout": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.login.onRequestExternalLogout.addListener(A.H.get<object>(callback));
    },
    "try_OnRequestExternalLogout": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.onRequestExternalLogout.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRequestExternalLogout": (): heap.Ref<boolean> => {
      if (WEBEXT?.login?.onRequestExternalLogout && "removeListener" in WEBEXT?.login?.onRequestExternalLogout) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRequestExternalLogout": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.onRequestExternalLogout.removeListener);
    },
    "call_OffRequestExternalLogout": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.login.onRequestExternalLogout.removeListener(A.H.get<object>(callback));
    },
    "try_OffRequestExternalLogout": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.onRequestExternalLogout.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRequestExternalLogout": (): heap.Ref<boolean> => {
      if (WEBEXT?.login?.onRequestExternalLogout && "hasListener" in WEBEXT?.login?.onRequestExternalLogout) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRequestExternalLogout": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.onRequestExternalLogout.hasListener);
    },
    "call_HasOnRequestExternalLogout": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.login.onRequestExternalLogout.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRequestExternalLogout": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.onRequestExternalLogout.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RequestExternalLogout": (): heap.Ref<boolean> => {
      if (WEBEXT?.login && "requestExternalLogout" in WEBEXT?.login) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RequestExternalLogout": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.requestExternalLogout);
    },
    "call_RequestExternalLogout": (retPtr: Pointer): void => {
      const _ret = WEBEXT.login.requestExternalLogout();
      A.store.Ref(retPtr, _ret);
    },
    "try_RequestExternalLogout": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.requestExternalLogout();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetDataForNextLoginAttempt": (): heap.Ref<boolean> => {
      if (WEBEXT?.login && "setDataForNextLoginAttempt" in WEBEXT?.login) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetDataForNextLoginAttempt": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.setDataForNextLoginAttempt);
    },
    "call_SetDataForNextLoginAttempt": (retPtr: Pointer, dataForNextLoginAttempt: heap.Ref<object>): void => {
      const _ret = WEBEXT.login.setDataForNextLoginAttempt(A.H.get<object>(dataForNextLoginAttempt));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetDataForNextLoginAttempt": (
      retPtr: Pointer,
      errPtr: Pointer,
      dataForNextLoginAttempt: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.setDataForNextLoginAttempt(A.H.get<object>(dataForNextLoginAttempt));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UnlockCurrentSession": (): heap.Ref<boolean> => {
      if (WEBEXT?.login && "unlockCurrentSession" in WEBEXT?.login) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UnlockCurrentSession": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.unlockCurrentSession);
    },
    "call_UnlockCurrentSession": (retPtr: Pointer, password: heap.Ref<object>): void => {
      const _ret = WEBEXT.login.unlockCurrentSession(A.H.get<object>(password));
      A.store.Ref(retPtr, _ret);
    },
    "try_UnlockCurrentSession": (retPtr: Pointer, errPtr: Pointer, password: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.unlockCurrentSession(A.H.get<object>(password));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UnlockManagedGuestSession": (): heap.Ref<boolean> => {
      if (WEBEXT?.login && "unlockManagedGuestSession" in WEBEXT?.login) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UnlockManagedGuestSession": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.unlockManagedGuestSession);
    },
    "call_UnlockManagedGuestSession": (retPtr: Pointer, password: heap.Ref<object>): void => {
      const _ret = WEBEXT.login.unlockManagedGuestSession(A.H.get<object>(password));
      A.store.Ref(retPtr, _ret);
    },
    "try_UnlockManagedGuestSession": (
      retPtr: Pointer,
      errPtr: Pointer,
      password: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.unlockManagedGuestSession(A.H.get<object>(password));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UnlockSharedSession": (): heap.Ref<boolean> => {
      if (WEBEXT?.login && "unlockSharedSession" in WEBEXT?.login) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UnlockSharedSession": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.login.unlockSharedSession);
    },
    "call_UnlockSharedSession": (retPtr: Pointer, password: heap.Ref<object>): void => {
      const _ret = WEBEXT.login.unlockSharedSession(A.H.get<object>(password));
      A.store.Ref(retPtr, _ret);
    },
    "try_UnlockSharedSession": (retPtr: Pointer, errPtr: Pointer, password: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.login.unlockSharedSession(A.H.get<object>(password));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
