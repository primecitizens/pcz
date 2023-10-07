import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/identity", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AccountInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["id"]);
      }
    },
    "load_AccountInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_AccountStatus": (ref: heap.Ref<string>): number => {
      const idx = ["SYNC", "ANY"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_GetAuthTokenResult": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["token"]);
        A.store.Ref(ptr + 4, x["grantedScopes"]);
      }
    },
    "load_GetAuthTokenResult": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["token"] = A.load.Ref(ptr + 0, undefined);
      x["grantedScopes"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ProfileUserInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["email"]);
        A.store.Ref(ptr + 4, x["id"]);
      }
    },
    "load_ProfileUserInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["email"] = A.load.Ref(ptr + 0, undefined);
      x["id"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_InvalidTokenDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["token"]);
      }
    },
    "load_InvalidTokenDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["token"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ProfileDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["SYNC", "ANY"].indexOf(x["accountStatus"] as string));
      }
    },
    "load_ProfileDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["accountStatus"] = A.load.Enum(ptr + 0, ["SYNC", "ANY"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TokenDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 0, false);

        A.store.Bool(ptr + 4 + 4, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
      } else {
        A.store.Bool(ptr + 19, true);
        A.store.Bool(ptr + 17, "interactive" in x ? true : false);
        A.store.Bool(ptr + 0, x["interactive"] ? true : false);

        if (typeof x["account"] === "undefined") {
          A.store.Bool(ptr + 4 + 4, false);
          A.store.Ref(ptr + 4 + 0, undefined);
        } else {
          A.store.Bool(ptr + 4 + 4, true);
          A.store.Ref(ptr + 4 + 0, x["account"]["id"]);
        }
        A.store.Ref(ptr + 12, x["scopes"]);
        A.store.Bool(ptr + 18, "enableGranularPermissions" in x ? true : false);
        A.store.Bool(ptr + 16, x["enableGranularPermissions"] ? true : false);
      }
    },
    "load_TokenDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 17)) {
        x["interactive"] = A.load.Bool(ptr + 0);
      } else {
        delete x["interactive"];
      }
      if (A.load.Bool(ptr + 4 + 4)) {
        x["account"] = {};
        x["account"]["id"] = A.load.Ref(ptr + 4 + 0, undefined);
      } else {
        delete x["account"];
      }
      x["scopes"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 18)) {
        x["enableGranularPermissions"] = A.load.Bool(ptr + 16);
      } else {
        delete x["enableGranularPermissions"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_WebAuthFlowDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 15, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 14, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 15, true);
        A.store.Ref(ptr + 0, x["url"]);
        A.store.Bool(ptr + 12, "interactive" in x ? true : false);
        A.store.Bool(ptr + 4, x["interactive"] ? true : false);
        A.store.Bool(ptr + 13, "abortOnLoadForNonInteractive" in x ? true : false);
        A.store.Bool(ptr + 5, x["abortOnLoadForNonInteractive"] ? true : false);
        A.store.Bool(ptr + 14, "timeoutMsForNonInteractive" in x ? true : false);
        A.store.Int32(
          ptr + 8,
          x["timeoutMsForNonInteractive"] === undefined ? 0 : (x["timeoutMsForNonInteractive"] as number)
        );
      }
    },
    "load_WebAuthFlowDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["interactive"] = A.load.Bool(ptr + 4);
      } else {
        delete x["interactive"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["abortOnLoadForNonInteractive"] = A.load.Bool(ptr + 5);
      } else {
        delete x["abortOnLoadForNonInteractive"];
      }
      if (A.load.Bool(ptr + 14)) {
        x["timeoutMsForNonInteractive"] = A.load.Int32(ptr + 8);
      } else {
        delete x["timeoutMsForNonInteractive"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_ClearAllCachedAuthTokens": (): heap.Ref<boolean> => {
      if (WEBEXT?.identity && "clearAllCachedAuthTokens" in WEBEXT?.identity) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ClearAllCachedAuthTokens": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.identity.clearAllCachedAuthTokens);
    },
    "call_ClearAllCachedAuthTokens": (retPtr: Pointer): void => {
      const _ret = WEBEXT.identity.clearAllCachedAuthTokens();
      A.store.Ref(retPtr, _ret);
    },
    "try_ClearAllCachedAuthTokens": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.identity.clearAllCachedAuthTokens();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAccounts": (): heap.Ref<boolean> => {
      if (WEBEXT?.identity && "getAccounts" in WEBEXT?.identity) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAccounts": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.identity.getAccounts);
    },
    "call_GetAccounts": (retPtr: Pointer): void => {
      const _ret = WEBEXT.identity.getAccounts();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAccounts": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.identity.getAccounts();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAuthToken": (): heap.Ref<boolean> => {
      if (WEBEXT?.identity && "getAuthToken" in WEBEXT?.identity) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAuthToken": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.identity.getAuthToken);
    },
    "call_GetAuthToken": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 17)) {
        details_ffi["interactive"] = A.load.Bool(details + 0);
      }
      if (A.load.Bool(details + 4 + 4)) {
        details_ffi["account"] = {};
        details_ffi["account"]["id"] = A.load.Ref(details + 4 + 0, undefined);
      }
      details_ffi["scopes"] = A.load.Ref(details + 12, undefined);
      if (A.load.Bool(details + 18)) {
        details_ffi["enableGranularPermissions"] = A.load.Bool(details + 16);
      }

      const _ret = WEBEXT.identity.getAuthToken(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAuthToken": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 17)) {
          details_ffi["interactive"] = A.load.Bool(details + 0);
        }
        if (A.load.Bool(details + 4 + 4)) {
          details_ffi["account"] = {};
          details_ffi["account"]["id"] = A.load.Ref(details + 4 + 0, undefined);
        }
        details_ffi["scopes"] = A.load.Ref(details + 12, undefined);
        if (A.load.Bool(details + 18)) {
          details_ffi["enableGranularPermissions"] = A.load.Bool(details + 16);
        }

        const _ret = WEBEXT.identity.getAuthToken(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetProfileUserInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.identity && "getProfileUserInfo" in WEBEXT?.identity) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetProfileUserInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.identity.getProfileUserInfo);
    },
    "call_GetProfileUserInfo": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["accountStatus"] = A.load.Enum(details + 0, ["SYNC", "ANY"]);

      const _ret = WEBEXT.identity.getProfileUserInfo(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetProfileUserInfo": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["accountStatus"] = A.load.Enum(details + 0, ["SYNC", "ANY"]);

        const _ret = WEBEXT.identity.getProfileUserInfo(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetRedirectURL": (): heap.Ref<boolean> => {
      if (WEBEXT?.identity && "getRedirectURL" in WEBEXT?.identity) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetRedirectURL": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.identity.getRedirectURL);
    },
    "call_GetRedirectURL": (retPtr: Pointer, path: heap.Ref<object>): void => {
      const _ret = WEBEXT.identity.getRedirectURL(A.H.get<object>(path));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetRedirectURL": (retPtr: Pointer, errPtr: Pointer, path: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.identity.getRedirectURL(A.H.get<object>(path));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LaunchWebAuthFlow": (): heap.Ref<boolean> => {
      if (WEBEXT?.identity && "launchWebAuthFlow" in WEBEXT?.identity) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LaunchWebAuthFlow": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.identity.launchWebAuthFlow);
    },
    "call_LaunchWebAuthFlow": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["url"] = A.load.Ref(details + 0, undefined);
      if (A.load.Bool(details + 12)) {
        details_ffi["interactive"] = A.load.Bool(details + 4);
      }
      if (A.load.Bool(details + 13)) {
        details_ffi["abortOnLoadForNonInteractive"] = A.load.Bool(details + 5);
      }
      if (A.load.Bool(details + 14)) {
        details_ffi["timeoutMsForNonInteractive"] = A.load.Int32(details + 8);
      }

      const _ret = WEBEXT.identity.launchWebAuthFlow(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_LaunchWebAuthFlow": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["url"] = A.load.Ref(details + 0, undefined);
        if (A.load.Bool(details + 12)) {
          details_ffi["interactive"] = A.load.Bool(details + 4);
        }
        if (A.load.Bool(details + 13)) {
          details_ffi["abortOnLoadForNonInteractive"] = A.load.Bool(details + 5);
        }
        if (A.load.Bool(details + 14)) {
          details_ffi["timeoutMsForNonInteractive"] = A.load.Int32(details + 8);
        }

        const _ret = WEBEXT.identity.launchWebAuthFlow(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSignInChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.identity?.onSignInChanged && "addListener" in WEBEXT?.identity?.onSignInChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSignInChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.identity.onSignInChanged.addListener);
    },
    "call_OnSignInChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.identity.onSignInChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnSignInChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.identity.onSignInChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSignInChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.identity?.onSignInChanged && "removeListener" in WEBEXT?.identity?.onSignInChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSignInChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.identity.onSignInChanged.removeListener);
    },
    "call_OffSignInChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.identity.onSignInChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffSignInChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.identity.onSignInChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSignInChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.identity?.onSignInChanged && "hasListener" in WEBEXT?.identity?.onSignInChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSignInChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.identity.onSignInChanged.hasListener);
    },
    "call_HasOnSignInChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.identity.onSignInChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSignInChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.identity.onSignInChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveCachedAuthToken": (): heap.Ref<boolean> => {
      if (WEBEXT?.identity && "removeCachedAuthToken" in WEBEXT?.identity) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveCachedAuthToken": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.identity.removeCachedAuthToken);
    },
    "call_RemoveCachedAuthToken": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["token"] = A.load.Ref(details + 0, undefined);

      const _ret = WEBEXT.identity.removeCachedAuthToken(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveCachedAuthToken": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["token"] = A.load.Ref(details + 0, undefined);

        const _ret = WEBEXT.identity.removeCachedAuthToken(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
