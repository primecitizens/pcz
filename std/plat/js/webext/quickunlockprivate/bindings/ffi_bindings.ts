import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/quickunlockprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_CredentialProblem": (ref: heap.Ref<string>): number => {
      const idx = ["TOO_SHORT", "TOO_LONG", "TOO_WEAK", "CONTAINS_NONDIGIT"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_CredentialCheck": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["errors"]);
        A.store.Ref(ptr + 4, x["warnings"]);
      }
    },
    "load_CredentialCheck": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["errors"] = A.load.Ref(ptr + 0, undefined);
      x["warnings"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CredentialRequirements": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "minLength" in x ? true : false);
        A.store.Int32(ptr + 0, x["minLength"] === undefined ? 0 : (x["minLength"] as number));
        A.store.Bool(ptr + 9, "maxLength" in x ? true : false);
        A.store.Int32(ptr + 4, x["maxLength"] === undefined ? 0 : (x["maxLength"] as number));
      }
    },
    "load_CredentialRequirements": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["minLength"] = A.load.Int32(ptr + 0);
      } else {
        delete x["minLength"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["maxLength"] = A.load.Int32(ptr + 4);
      } else {
        delete x["maxLength"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_QuickUnlockMode": (ref: heap.Ref<string>): number => {
      const idx = ["PIN"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_TokenInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Ref(ptr + 0, x["token"]);
        A.store.Bool(ptr + 8, "lifetimeSeconds" in x ? true : false);
        A.store.Int32(ptr + 4, x["lifetimeSeconds"] === undefined ? 0 : (x["lifetimeSeconds"] as number));
      }
    },
    "load_TokenInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["token"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 8)) {
        x["lifetimeSeconds"] = A.load.Int32(ptr + 4);
      } else {
        delete x["lifetimeSeconds"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_CanAuthenticatePin": (): heap.Ref<boolean> => {
      if (WEBEXT?.quickUnlockPrivate && "canAuthenticatePin" in WEBEXT?.quickUnlockPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CanAuthenticatePin": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.quickUnlockPrivate.canAuthenticatePin);
    },
    "call_CanAuthenticatePin": (retPtr: Pointer): void => {
      const _ret = WEBEXT.quickUnlockPrivate.canAuthenticatePin();
      A.store.Ref(retPtr, _ret);
    },
    "try_CanAuthenticatePin": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.quickUnlockPrivate.canAuthenticatePin();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CheckCredential": (): heap.Ref<boolean> => {
      if (WEBEXT?.quickUnlockPrivate && "checkCredential" in WEBEXT?.quickUnlockPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CheckCredential": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.quickUnlockPrivate.checkCredential);
    },
    "call_CheckCredential": (retPtr: Pointer, mode: number, credential: heap.Ref<object>): void => {
      const _ret = WEBEXT.quickUnlockPrivate.checkCredential(
        mode > 0 && mode <= 1 ? ["PIN"][mode - 1] : undefined,
        A.H.get<object>(credential)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_CheckCredential": (
      retPtr: Pointer,
      errPtr: Pointer,
      mode: number,
      credential: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.quickUnlockPrivate.checkCredential(
          mode > 0 && mode <= 1 ? ["PIN"][mode - 1] : undefined,
          A.H.get<object>(credential)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetActiveModes": (): heap.Ref<boolean> => {
      if (WEBEXT?.quickUnlockPrivate && "getActiveModes" in WEBEXT?.quickUnlockPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetActiveModes": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.quickUnlockPrivate.getActiveModes);
    },
    "call_GetActiveModes": (retPtr: Pointer): void => {
      const _ret = WEBEXT.quickUnlockPrivate.getActiveModes();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetActiveModes": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.quickUnlockPrivate.getActiveModes();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAuthToken": (): heap.Ref<boolean> => {
      if (WEBEXT?.quickUnlockPrivate && "getAuthToken" in WEBEXT?.quickUnlockPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAuthToken": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.quickUnlockPrivate.getAuthToken);
    },
    "call_GetAuthToken": (retPtr: Pointer, accountPassword: heap.Ref<object>): void => {
      const _ret = WEBEXT.quickUnlockPrivate.getAuthToken(A.H.get<object>(accountPassword));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAuthToken": (retPtr: Pointer, errPtr: Pointer, accountPassword: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.quickUnlockPrivate.getAuthToken(A.H.get<object>(accountPassword));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAvailableModes": (): heap.Ref<boolean> => {
      if (WEBEXT?.quickUnlockPrivate && "getAvailableModes" in WEBEXT?.quickUnlockPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAvailableModes": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.quickUnlockPrivate.getAvailableModes);
    },
    "call_GetAvailableModes": (retPtr: Pointer): void => {
      const _ret = WEBEXT.quickUnlockPrivate.getAvailableModes();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAvailableModes": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.quickUnlockPrivate.getAvailableModes();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCredentialRequirements": (): heap.Ref<boolean> => {
      if (WEBEXT?.quickUnlockPrivate && "getCredentialRequirements" in WEBEXT?.quickUnlockPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCredentialRequirements": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.quickUnlockPrivate.getCredentialRequirements);
    },
    "call_GetCredentialRequirements": (retPtr: Pointer, mode: number): void => {
      const _ret = WEBEXT.quickUnlockPrivate.getCredentialRequirements(
        mode > 0 && mode <= 1 ? ["PIN"][mode - 1] : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCredentialRequirements": (retPtr: Pointer, errPtr: Pointer, mode: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.quickUnlockPrivate.getCredentialRequirements(
          mode > 0 && mode <= 1 ? ["PIN"][mode - 1] : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnActiveModesChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.quickUnlockPrivate?.onActiveModesChanged &&
        "addListener" in WEBEXT?.quickUnlockPrivate?.onActiveModesChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnActiveModesChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.quickUnlockPrivate.onActiveModesChanged.addListener);
    },
    "call_OnActiveModesChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.quickUnlockPrivate.onActiveModesChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnActiveModesChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.quickUnlockPrivate.onActiveModesChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffActiveModesChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.quickUnlockPrivate?.onActiveModesChanged &&
        "removeListener" in WEBEXT?.quickUnlockPrivate?.onActiveModesChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffActiveModesChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.quickUnlockPrivate.onActiveModesChanged.removeListener);
    },
    "call_OffActiveModesChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.quickUnlockPrivate.onActiveModesChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffActiveModesChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.quickUnlockPrivate.onActiveModesChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnActiveModesChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.quickUnlockPrivate?.onActiveModesChanged &&
        "hasListener" in WEBEXT?.quickUnlockPrivate?.onActiveModesChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnActiveModesChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.quickUnlockPrivate.onActiveModesChanged.hasListener);
    },
    "call_HasOnActiveModesChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.quickUnlockPrivate.onActiveModesChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnActiveModesChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.quickUnlockPrivate.onActiveModesChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetLockScreenEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.quickUnlockPrivate && "setLockScreenEnabled" in WEBEXT?.quickUnlockPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetLockScreenEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.quickUnlockPrivate.setLockScreenEnabled);
    },
    "call_SetLockScreenEnabled": (retPtr: Pointer, token: heap.Ref<object>, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.quickUnlockPrivate.setLockScreenEnabled(A.H.get<object>(token), enabled === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetLockScreenEnabled": (
      retPtr: Pointer,
      errPtr: Pointer,
      token: heap.Ref<object>,
      enabled: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.quickUnlockPrivate.setLockScreenEnabled(A.H.get<object>(token), enabled === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetModes": (): heap.Ref<boolean> => {
      if (WEBEXT?.quickUnlockPrivate && "setModes" in WEBEXT?.quickUnlockPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetModes": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.quickUnlockPrivate.setModes);
    },
    "call_SetModes": (
      retPtr: Pointer,
      token: heap.Ref<object>,
      modes: heap.Ref<object>,
      credentials: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.quickUnlockPrivate.setModes(
        A.H.get<object>(token),
        A.H.get<object>(modes),
        A.H.get<object>(credentials)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_SetModes": (
      retPtr: Pointer,
      errPtr: Pointer,
      token: heap.Ref<object>,
      modes: heap.Ref<object>,
      credentials: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.quickUnlockPrivate.setModes(
          A.H.get<object>(token),
          A.H.get<object>(modes),
          A.H.get<object>(credentials)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPinAutosubmitEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.quickUnlockPrivate && "setPinAutosubmitEnabled" in WEBEXT?.quickUnlockPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPinAutosubmitEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.quickUnlockPrivate.setPinAutosubmitEnabled);
    },
    "call_SetPinAutosubmitEnabled": (
      retPtr: Pointer,
      token: heap.Ref<object>,
      pin: heap.Ref<object>,
      enabled: heap.Ref<boolean>
    ): void => {
      const _ret = WEBEXT.quickUnlockPrivate.setPinAutosubmitEnabled(
        A.H.get<object>(token),
        A.H.get<object>(pin),
        enabled === A.H.TRUE
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_SetPinAutosubmitEnabled": (
      retPtr: Pointer,
      errPtr: Pointer,
      token: heap.Ref<object>,
      pin: heap.Ref<object>,
      enabled: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.quickUnlockPrivate.setPinAutosubmitEnabled(
          A.H.get<object>(token),
          A.H.get<object>(pin),
          enabled === A.H.TRUE
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
