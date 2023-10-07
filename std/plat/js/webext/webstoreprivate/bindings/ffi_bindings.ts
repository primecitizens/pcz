import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/webstoreprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_ExtensionInstallStatus": (ref: heap.Ref<string>): number => {
      const idx = [
        "can_request",
        "request_pending",
        "blocked_by_policy",
        "installable",
        "enabled",
        "disabled",
        "terminated",
        "blacklisted",
        "custodian_approval_required",
        "force_installed",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_GetBrowserLoginReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["login"]);
      }
    },
    "load_GetBrowserLoginReturnType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["login"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_Result": (ref: heap.Ref<string>): number => {
      const idx = [
        "",
        "success",
        "user_gesture_required",
        "unknown_error",
        "feature_disabled",
        "unsupported_extension_type",
        "missing_dependencies",
        "install_error",
        "user_cancelled",
        "invalid_id",
        "blacklisted",
        "blocked_by_policy",
        "install_in_progress",
        "launch_in_progress",
        "manifest_error",
        "icon_error",
        "invalid_icon_url",
        "already_installed",
        "blocked_for_child_account",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_WebGlStatus": (ref: heap.Ref<string>): number => {
      const idx = ["webgl_allowed", "webgl_blocked"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_BeginInstallWithManifest3": (): heap.Ref<boolean> => {
      if (WEBEXT?.webstorePrivate && "beginInstallWithManifest3" in WEBEXT?.webstorePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_BeginInstallWithManifest3": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webstorePrivate.beginInstallWithManifest3);
    },
    "call_BeginInstallWithManifest3": (retPtr: Pointer, details: heap.Ref<object>): void => {
      const _ret = WEBEXT.webstorePrivate.beginInstallWithManifest3(A.H.get<object>(details));
      A.store.Ref(retPtr, _ret);
    },
    "try_BeginInstallWithManifest3": (
      retPtr: Pointer,
      errPtr: Pointer,
      details: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webstorePrivate.beginInstallWithManifest3(A.H.get<object>(details));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CompleteInstall": (): heap.Ref<boolean> => {
      if (WEBEXT?.webstorePrivate && "completeInstall" in WEBEXT?.webstorePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CompleteInstall": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webstorePrivate.completeInstall);
    },
    "call_CompleteInstall": (retPtr: Pointer, expected_id: heap.Ref<object>): void => {
      const _ret = WEBEXT.webstorePrivate.completeInstall(A.H.get<object>(expected_id));
      A.store.Ref(retPtr, _ret);
    },
    "try_CompleteInstall": (retPtr: Pointer, errPtr: Pointer, expected_id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webstorePrivate.completeInstall(A.H.get<object>(expected_id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_EnableAppLauncher": (): heap.Ref<boolean> => {
      if (WEBEXT?.webstorePrivate && "enableAppLauncher" in WEBEXT?.webstorePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EnableAppLauncher": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webstorePrivate.enableAppLauncher);
    },
    "call_EnableAppLauncher": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webstorePrivate.enableAppLauncher();
      A.store.Ref(retPtr, _ret);
    },
    "try_EnableAppLauncher": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webstorePrivate.enableAppLauncher();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetBrowserLogin": (): heap.Ref<boolean> => {
      if (WEBEXT?.webstorePrivate && "getBrowserLogin" in WEBEXT?.webstorePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetBrowserLogin": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webstorePrivate.getBrowserLogin);
    },
    "call_GetBrowserLogin": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webstorePrivate.getBrowserLogin();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetBrowserLogin": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webstorePrivate.getBrowserLogin();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetExtensionStatus": (): heap.Ref<boolean> => {
      if (WEBEXT?.webstorePrivate && "getExtensionStatus" in WEBEXT?.webstorePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetExtensionStatus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webstorePrivate.getExtensionStatus);
    },
    "call_GetExtensionStatus": (retPtr: Pointer, id: heap.Ref<object>, manifest: heap.Ref<object>): void => {
      const _ret = WEBEXT.webstorePrivate.getExtensionStatus(A.H.get<object>(id), A.H.get<object>(manifest));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetExtensionStatus": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<object>,
      manifest: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webstorePrivate.getExtensionStatus(A.H.get<object>(id), A.H.get<object>(manifest));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetIsLauncherEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.webstorePrivate && "getIsLauncherEnabled" in WEBEXT?.webstorePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetIsLauncherEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webstorePrivate.getIsLauncherEnabled);
    },
    "call_GetIsLauncherEnabled": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webstorePrivate.getIsLauncherEnabled();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetIsLauncherEnabled": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webstorePrivate.getIsLauncherEnabled();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetReferrerChain": (): heap.Ref<boolean> => {
      if (WEBEXT?.webstorePrivate && "getReferrerChain" in WEBEXT?.webstorePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetReferrerChain": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webstorePrivate.getReferrerChain);
    },
    "call_GetReferrerChain": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webstorePrivate.getReferrerChain();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetReferrerChain": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webstorePrivate.getReferrerChain();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetStoreLogin": (): heap.Ref<boolean> => {
      if (WEBEXT?.webstorePrivate && "getStoreLogin" in WEBEXT?.webstorePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetStoreLogin": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webstorePrivate.getStoreLogin);
    },
    "call_GetStoreLogin": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webstorePrivate.getStoreLogin();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetStoreLogin": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webstorePrivate.getStoreLogin();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetWebGLStatus": (): heap.Ref<boolean> => {
      if (WEBEXT?.webstorePrivate && "getWebGLStatus" in WEBEXT?.webstorePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetWebGLStatus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webstorePrivate.getWebGLStatus);
    },
    "call_GetWebGLStatus": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webstorePrivate.getWebGLStatus();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetWebGLStatus": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webstorePrivate.getWebGLStatus();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Install": (): heap.Ref<boolean> => {
      if (WEBEXT?.webstorePrivate && "install" in WEBEXT?.webstorePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Install": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webstorePrivate.install);
    },
    "call_Install": (retPtr: Pointer, expected_id: heap.Ref<object>): void => {
      const _ret = WEBEXT.webstorePrivate.install(A.H.get<object>(expected_id));
      A.store.Ref(retPtr, _ret);
    },
    "try_Install": (retPtr: Pointer, errPtr: Pointer, expected_id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webstorePrivate.install(A.H.get<object>(expected_id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsInIncognitoMode": (): heap.Ref<boolean> => {
      if (WEBEXT?.webstorePrivate && "isInIncognitoMode" in WEBEXT?.webstorePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsInIncognitoMode": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webstorePrivate.isInIncognitoMode);
    },
    "call_IsInIncognitoMode": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webstorePrivate.isInIncognitoMode();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsInIncognitoMode": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webstorePrivate.isInIncognitoMode();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsPendingCustodianApproval": (): heap.Ref<boolean> => {
      if (WEBEXT?.webstorePrivate && "isPendingCustodianApproval" in WEBEXT?.webstorePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsPendingCustodianApproval": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webstorePrivate.isPendingCustodianApproval);
    },
    "call_IsPendingCustodianApproval": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.webstorePrivate.isPendingCustodianApproval(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_IsPendingCustodianApproval": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webstorePrivate.isPendingCustodianApproval(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetStoreLogin": (): heap.Ref<boolean> => {
      if (WEBEXT?.webstorePrivate && "setStoreLogin" in WEBEXT?.webstorePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetStoreLogin": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webstorePrivate.setStoreLogin);
    },
    "call_SetStoreLogin": (retPtr: Pointer, login: heap.Ref<object>): void => {
      const _ret = WEBEXT.webstorePrivate.setStoreLogin(A.H.get<object>(login));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetStoreLogin": (retPtr: Pointer, errPtr: Pointer, login: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webstorePrivate.setStoreLogin(A.H.get<object>(login));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
