import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/management", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_ExtensionDisabledReason": (ref: heap.Ref<string>): number => {
      const idx = ["unknown", "permissions_increase"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_LaunchType": (ref: heap.Ref<string>): number => {
      const idx = ["OPEN_AS_REGULAR_TAB", "OPEN_AS_PINNED_TAB", "OPEN_AS_WINDOW", "OPEN_FULL_SCREEN"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },

    "store_IconInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Int64(ptr + 0, x["size"] === undefined ? 0 : (x["size"] as number));
        A.store.Ref(ptr + 8, x["url"]);
      }
    },
    "load_IconInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["size"] = A.load.Int64(ptr + 0);
      x["url"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ExtensionInstallType": (ref: heap.Ref<string>): number => {
      const idx = ["admin", "development", "normal", "sideload", "other"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ExtensionType": (ref: heap.Ref<string>): number => {
      const idx = [
        "extension",
        "hosted_app",
        "packaged_app",
        "legacy_packaged_app",
        "theme",
        "login_screen_extension",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ExtensionInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 89, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Enum(ptr + 12, -1);
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Enum(ptr + 36, -1);
        A.store.Bool(ptr + 40, false);
        A.store.Enum(ptr + 44, -1);
        A.store.Bool(ptr + 48, false);
        A.store.Bool(ptr + 88, false);
        A.store.Bool(ptr + 49, false);
        A.store.Ref(ptr + 52, undefined);
        A.store.Bool(ptr + 56, false);
        A.store.Ref(ptr + 60, undefined);
        A.store.Ref(ptr + 64, undefined);
        A.store.Ref(ptr + 68, undefined);
        A.store.Enum(ptr + 72, -1);
        A.store.Ref(ptr + 76, undefined);
        A.store.Ref(ptr + 80, undefined);
        A.store.Ref(ptr + 84, undefined);
      } else {
        A.store.Bool(ptr + 89, true);
        A.store.Ref(ptr + 0, x["appLaunchUrl"]);
        A.store.Ref(ptr + 4, x["availableLaunchTypes"]);
        A.store.Ref(ptr + 8, x["description"]);
        A.store.Enum(ptr + 12, ["unknown", "permissions_increase"].indexOf(x["disabledReason"] as string));
        A.store.Bool(ptr + 16, x["enabled"] ? true : false);
        A.store.Ref(ptr + 20, x["homepageUrl"]);
        A.store.Ref(ptr + 24, x["hostPermissions"]);
        A.store.Ref(ptr + 28, x["icons"]);
        A.store.Ref(ptr + 32, x["id"]);
        A.store.Enum(
          ptr + 36,
          ["admin", "development", "normal", "sideload", "other"].indexOf(x["installType"] as string)
        );
        A.store.Bool(ptr + 40, x["isApp"] ? true : false);
        A.store.Enum(
          ptr + 44,
          ["OPEN_AS_REGULAR_TAB", "OPEN_AS_PINNED_TAB", "OPEN_AS_WINDOW", "OPEN_FULL_SCREEN"].indexOf(
            x["launchType"] as string
          )
        );
        A.store.Bool(ptr + 48, x["mayDisable"] ? true : false);
        A.store.Bool(ptr + 88, "mayEnable" in x ? true : false);
        A.store.Bool(ptr + 49, x["mayEnable"] ? true : false);
        A.store.Ref(ptr + 52, x["name"]);
        A.store.Bool(ptr + 56, x["offlineEnabled"] ? true : false);
        A.store.Ref(ptr + 60, x["optionsUrl"]);
        A.store.Ref(ptr + 64, x["permissions"]);
        A.store.Ref(ptr + 68, x["shortName"]);
        A.store.Enum(
          ptr + 72,
          ["extension", "hosted_app", "packaged_app", "legacy_packaged_app", "theme", "login_screen_extension"].indexOf(
            x["type"] as string
          )
        );
        A.store.Ref(ptr + 76, x["updateUrl"]);
        A.store.Ref(ptr + 80, x["version"]);
        A.store.Ref(ptr + 84, x["versionName"]);
      }
    },
    "load_ExtensionInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["appLaunchUrl"] = A.load.Ref(ptr + 0, undefined);
      x["availableLaunchTypes"] = A.load.Ref(ptr + 4, undefined);
      x["description"] = A.load.Ref(ptr + 8, undefined);
      x["disabledReason"] = A.load.Enum(ptr + 12, ["unknown", "permissions_increase"]);
      x["enabled"] = A.load.Bool(ptr + 16);
      x["homepageUrl"] = A.load.Ref(ptr + 20, undefined);
      x["hostPermissions"] = A.load.Ref(ptr + 24, undefined);
      x["icons"] = A.load.Ref(ptr + 28, undefined);
      x["id"] = A.load.Ref(ptr + 32, undefined);
      x["installType"] = A.load.Enum(ptr + 36, ["admin", "development", "normal", "sideload", "other"]);
      x["isApp"] = A.load.Bool(ptr + 40);
      x["launchType"] = A.load.Enum(ptr + 44, [
        "OPEN_AS_REGULAR_TAB",
        "OPEN_AS_PINNED_TAB",
        "OPEN_AS_WINDOW",
        "OPEN_FULL_SCREEN",
      ]);
      x["mayDisable"] = A.load.Bool(ptr + 48);
      if (A.load.Bool(ptr + 88)) {
        x["mayEnable"] = A.load.Bool(ptr + 49);
      } else {
        delete x["mayEnable"];
      }
      x["name"] = A.load.Ref(ptr + 52, undefined);
      x["offlineEnabled"] = A.load.Bool(ptr + 56);
      x["optionsUrl"] = A.load.Ref(ptr + 60, undefined);
      x["permissions"] = A.load.Ref(ptr + 64, undefined);
      x["shortName"] = A.load.Ref(ptr + 68, undefined);
      x["type"] = A.load.Enum(ptr + 72, [
        "extension",
        "hosted_app",
        "packaged_app",
        "legacy_packaged_app",
        "theme",
        "login_screen_extension",
      ]);
      x["updateUrl"] = A.load.Ref(ptr + 76, undefined);
      x["version"] = A.load.Ref(ptr + 80, undefined);
      x["versionName"] = A.load.Ref(ptr + 84, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UninstallOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 1, "showConfirmDialog" in x ? true : false);
        A.store.Bool(ptr + 0, x["showConfirmDialog"] ? true : false);
      }
    },
    "load_UninstallOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 1)) {
        x["showConfirmDialog"] = A.load.Bool(ptr + 0);
      } else {
        delete x["showConfirmDialog"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_CreateAppShortcut": (): heap.Ref<boolean> => {
      if (WEBEXT?.management && "createAppShortcut" in WEBEXT?.management) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CreateAppShortcut": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.createAppShortcut);
    },
    "call_CreateAppShortcut": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.createAppShortcut(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_CreateAppShortcut": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.createAppShortcut(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GenerateAppForLink": (): heap.Ref<boolean> => {
      if (WEBEXT?.management && "generateAppForLink" in WEBEXT?.management) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GenerateAppForLink": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.generateAppForLink);
    },
    "call_GenerateAppForLink": (retPtr: Pointer, url: heap.Ref<object>, title: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.generateAppForLink(A.H.get<object>(url), A.H.get<object>(title));
      A.store.Ref(retPtr, _ret);
    },
    "try_GenerateAppForLink": (
      retPtr: Pointer,
      errPtr: Pointer,
      url: heap.Ref<object>,
      title: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.generateAppForLink(A.H.get<object>(url), A.H.get<object>(title));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Get": (): heap.Ref<boolean> => {
      if (WEBEXT?.management && "get" in WEBEXT?.management) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Get": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.get);
    },
    "call_Get": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.get(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_Get": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.get(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAll": (): heap.Ref<boolean> => {
      if (WEBEXT?.management && "getAll" in WEBEXT?.management) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.getAll);
    },
    "call_GetAll": (retPtr: Pointer): void => {
      const _ret = WEBEXT.management.getAll();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAll": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.getAll();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPermissionWarningsById": (): heap.Ref<boolean> => {
      if (WEBEXT?.management && "getPermissionWarningsById" in WEBEXT?.management) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPermissionWarningsById": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.getPermissionWarningsById);
    },
    "call_GetPermissionWarningsById": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.getPermissionWarningsById(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPermissionWarningsById": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.getPermissionWarningsById(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPermissionWarningsByManifest": (): heap.Ref<boolean> => {
      if (WEBEXT?.management && "getPermissionWarningsByManifest" in WEBEXT?.management) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPermissionWarningsByManifest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.getPermissionWarningsByManifest);
    },
    "call_GetPermissionWarningsByManifest": (retPtr: Pointer, manifestStr: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.getPermissionWarningsByManifest(A.H.get<object>(manifestStr));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPermissionWarningsByManifest": (
      retPtr: Pointer,
      errPtr: Pointer,
      manifestStr: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.getPermissionWarningsByManifest(A.H.get<object>(manifestStr));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSelf": (): heap.Ref<boolean> => {
      if (WEBEXT?.management && "getSelf" in WEBEXT?.management) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSelf": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.getSelf);
    },
    "call_GetSelf": (retPtr: Pointer): void => {
      const _ret = WEBEXT.management.getSelf();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSelf": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.getSelf();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InstallReplacementWebApp": (): heap.Ref<boolean> => {
      if (WEBEXT?.management && "installReplacementWebApp" in WEBEXT?.management) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InstallReplacementWebApp": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.installReplacementWebApp);
    },
    "call_InstallReplacementWebApp": (retPtr: Pointer): void => {
      const _ret = WEBEXT.management.installReplacementWebApp();
      A.store.Ref(retPtr, _ret);
    },
    "try_InstallReplacementWebApp": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.installReplacementWebApp();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LaunchApp": (): heap.Ref<boolean> => {
      if (WEBEXT?.management && "launchApp" in WEBEXT?.management) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LaunchApp": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.launchApp);
    },
    "call_LaunchApp": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.launchApp(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_LaunchApp": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.launchApp(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDisabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.management?.onDisabled && "addListener" in WEBEXT?.management?.onDisabled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDisabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.onDisabled.addListener);
    },
    "call_OnDisabled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.onDisabled.addListener(A.H.get<object>(callback));
    },
    "try_OnDisabled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.onDisabled.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDisabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.management?.onDisabled && "removeListener" in WEBEXT?.management?.onDisabled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDisabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.onDisabled.removeListener);
    },
    "call_OffDisabled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.onDisabled.removeListener(A.H.get<object>(callback));
    },
    "try_OffDisabled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.onDisabled.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDisabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.management?.onDisabled && "hasListener" in WEBEXT?.management?.onDisabled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDisabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.onDisabled.hasListener);
    },
    "call_HasOnDisabled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.onDisabled.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDisabled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.onDisabled.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.management?.onEnabled && "addListener" in WEBEXT?.management?.onEnabled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.onEnabled.addListener);
    },
    "call_OnEnabled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.onEnabled.addListener(A.H.get<object>(callback));
    },
    "try_OnEnabled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.onEnabled.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.management?.onEnabled && "removeListener" in WEBEXT?.management?.onEnabled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.onEnabled.removeListener);
    },
    "call_OffEnabled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.onEnabled.removeListener(A.H.get<object>(callback));
    },
    "try_OffEnabled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.onEnabled.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.management?.onEnabled && "hasListener" in WEBEXT?.management?.onEnabled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.onEnabled.hasListener);
    },
    "call_HasOnEnabled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.onEnabled.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnEnabled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.onEnabled.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnInstalled": (): heap.Ref<boolean> => {
      if (WEBEXT?.management?.onInstalled && "addListener" in WEBEXT?.management?.onInstalled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnInstalled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.onInstalled.addListener);
    },
    "call_OnInstalled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.onInstalled.addListener(A.H.get<object>(callback));
    },
    "try_OnInstalled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.onInstalled.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffInstalled": (): heap.Ref<boolean> => {
      if (WEBEXT?.management?.onInstalled && "removeListener" in WEBEXT?.management?.onInstalled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffInstalled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.onInstalled.removeListener);
    },
    "call_OffInstalled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.onInstalled.removeListener(A.H.get<object>(callback));
    },
    "try_OffInstalled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.onInstalled.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnInstalled": (): heap.Ref<boolean> => {
      if (WEBEXT?.management?.onInstalled && "hasListener" in WEBEXT?.management?.onInstalled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnInstalled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.onInstalled.hasListener);
    },
    "call_HasOnInstalled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.onInstalled.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnInstalled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.onInstalled.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnUninstalled": (): heap.Ref<boolean> => {
      if (WEBEXT?.management?.onUninstalled && "addListener" in WEBEXT?.management?.onUninstalled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnUninstalled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.onUninstalled.addListener);
    },
    "call_OnUninstalled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.onUninstalled.addListener(A.H.get<object>(callback));
    },
    "try_OnUninstalled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.onUninstalled.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffUninstalled": (): heap.Ref<boolean> => {
      if (WEBEXT?.management?.onUninstalled && "removeListener" in WEBEXT?.management?.onUninstalled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffUninstalled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.onUninstalled.removeListener);
    },
    "call_OffUninstalled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.onUninstalled.removeListener(A.H.get<object>(callback));
    },
    "try_OffUninstalled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.onUninstalled.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnUninstalled": (): heap.Ref<boolean> => {
      if (WEBEXT?.management?.onUninstalled && "hasListener" in WEBEXT?.management?.onUninstalled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnUninstalled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.onUninstalled.hasListener);
    },
    "call_HasOnUninstalled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.management.onUninstalled.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnUninstalled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.onUninstalled.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.management && "setEnabled" in WEBEXT?.management) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.setEnabled);
    },
    "call_SetEnabled": (retPtr: Pointer, id: heap.Ref<object>, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.management.setEnabled(A.H.get<object>(id), enabled === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetEnabled": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<object>,
      enabled: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.setEnabled(A.H.get<object>(id), enabled === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetLaunchType": (): heap.Ref<boolean> => {
      if (WEBEXT?.management && "setLaunchType" in WEBEXT?.management) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetLaunchType": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.setLaunchType);
    },
    "call_SetLaunchType": (retPtr: Pointer, id: heap.Ref<object>, launchType: number): void => {
      const _ret = WEBEXT.management.setLaunchType(
        A.H.get<object>(id),
        launchType > 0 && launchType <= 4
          ? ["OPEN_AS_REGULAR_TAB", "OPEN_AS_PINNED_TAB", "OPEN_AS_WINDOW", "OPEN_FULL_SCREEN"][launchType - 1]
          : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_SetLaunchType": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<object>,
      launchType: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.management.setLaunchType(
          A.H.get<object>(id),
          launchType > 0 && launchType <= 4
            ? ["OPEN_AS_REGULAR_TAB", "OPEN_AS_PINNED_TAB", "OPEN_AS_WINDOW", "OPEN_FULL_SCREEN"][launchType - 1]
            : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Uninstall": (): heap.Ref<boolean> => {
      if (WEBEXT?.management && "uninstall" in WEBEXT?.management) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Uninstall": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.uninstall);
    },
    "call_Uninstall": (retPtr: Pointer, id: heap.Ref<object>, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 1)) {
        options_ffi["showConfirmDialog"] = A.load.Bool(options + 0);
      }

      const _ret = WEBEXT.management.uninstall(A.H.get<object>(id), options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Uninstall": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 1)) {
          options_ffi["showConfirmDialog"] = A.load.Bool(options + 0);
        }

        const _ret = WEBEXT.management.uninstall(A.H.get<object>(id), options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UninstallSelf": (): heap.Ref<boolean> => {
      if (WEBEXT?.management && "uninstallSelf" in WEBEXT?.management) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UninstallSelf": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.management.uninstallSelf);
    },
    "call_UninstallSelf": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 1)) {
        options_ffi["showConfirmDialog"] = A.load.Bool(options + 0);
      }

      const _ret = WEBEXT.management.uninstallSelf(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_UninstallSelf": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 1)) {
          options_ffi["showConfirmDialog"] = A.load.Bool(options + 0);
        }

        const _ret = WEBEXT.management.uninstallSelf(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
