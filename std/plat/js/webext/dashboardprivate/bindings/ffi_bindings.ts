import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/dashboardprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_Result": (ref: heap.Ref<string>): number => {
      const idx = [
        "",
        "unknown_error",
        "user_cancelled",
        "invalid_id",
        "manifest_error",
        "icon_error",
        "invalid_icon_url",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ShowPermissionPromptForDelegatedInstallArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Ref(ptr + 0, x["delegatedUser"]);
        A.store.Ref(ptr + 4, x["iconUrl"]);
        A.store.Ref(ptr + 8, x["id"]);
        A.store.Ref(ptr + 12, x["localizedName"]);
        A.store.Ref(ptr + 16, x["manifest"]);
      }
    },
    "load_ShowPermissionPromptForDelegatedInstallArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["delegatedUser"] = A.load.Ref(ptr + 0, undefined);
      x["iconUrl"] = A.load.Ref(ptr + 4, undefined);
      x["id"] = A.load.Ref(ptr + 8, undefined);
      x["localizedName"] = A.load.Ref(ptr + 12, undefined);
      x["manifest"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_ShowPermissionPromptForDelegatedInstall": (): heap.Ref<boolean> => {
      if (WEBEXT?.dashboardPrivate && "showPermissionPromptForDelegatedInstall" in WEBEXT?.dashboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ShowPermissionPromptForDelegatedInstall": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.dashboardPrivate.showPermissionPromptForDelegatedInstall);
    },
    "call_ShowPermissionPromptForDelegatedInstall": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["delegatedUser"] = A.load.Ref(details + 0, undefined);
      details_ffi["iconUrl"] = A.load.Ref(details + 4, undefined);
      details_ffi["id"] = A.load.Ref(details + 8, undefined);
      details_ffi["localizedName"] = A.load.Ref(details + 12, undefined);
      details_ffi["manifest"] = A.load.Ref(details + 16, undefined);

      const _ret = WEBEXT.dashboardPrivate.showPermissionPromptForDelegatedInstall(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ShowPermissionPromptForDelegatedInstall": (
      retPtr: Pointer,
      errPtr: Pointer,
      details: Pointer
    ): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["delegatedUser"] = A.load.Ref(details + 0, undefined);
        details_ffi["iconUrl"] = A.load.Ref(details + 4, undefined);
        details_ffi["id"] = A.load.Ref(details + 8, undefined);
        details_ffi["localizedName"] = A.load.Ref(details + 12, undefined);
        details_ffi["manifest"] = A.load.Ref(details + 16, undefined);

        const _ret = WEBEXT.dashboardPrivate.showPermissionPromptForDelegatedInstall(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
