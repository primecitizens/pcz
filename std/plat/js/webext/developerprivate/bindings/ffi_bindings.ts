import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/developerprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AccessModifier": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 1, false);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Bool(ptr + 2, "isEnabled" in x ? true : false);
        A.store.Bool(ptr + 0, x["isEnabled"] ? true : false);
        A.store.Bool(ptr + 3, "isActive" in x ? true : false);
        A.store.Bool(ptr + 1, x["isActive"] ? true : false);
      }
    },
    "load_AccessModifier": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 2)) {
        x["isEnabled"] = A.load.Bool(ptr + 0);
      } else {
        delete x["isEnabled"];
      }
      if (A.load.Bool(ptr + 3)) {
        x["isActive"] = A.load.Bool(ptr + 1);
      } else {
        delete x["isActive"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_CommandScope": (ref: heap.Ref<string>): number => {
      const idx = ["GLOBAL", "CHROME"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Command": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 23, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 12, false);
        A.store.Enum(ptr + 16, -1);
        A.store.Bool(ptr + 22, false);
        A.store.Bool(ptr + 20, false);
      } else {
        A.store.Bool(ptr + 23, true);
        A.store.Ref(ptr + 0, x["description"]);
        A.store.Ref(ptr + 4, x["keybinding"]);
        A.store.Ref(ptr + 8, x["name"]);
        A.store.Bool(ptr + 21, "isActive" in x ? true : false);
        A.store.Bool(ptr + 12, x["isActive"] ? true : false);
        A.store.Enum(ptr + 16, ["GLOBAL", "CHROME"].indexOf(x["scope"] as string));
        A.store.Bool(ptr + 22, "isExtensionAction" in x ? true : false);
        A.store.Bool(ptr + 20, x["isExtensionAction"] ? true : false);
      }
    },
    "load_Command": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["description"] = A.load.Ref(ptr + 0, undefined);
      x["keybinding"] = A.load.Ref(ptr + 4, undefined);
      x["name"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 21)) {
        x["isActive"] = A.load.Bool(ptr + 12);
      } else {
        delete x["isActive"];
      }
      x["scope"] = A.load.Enum(ptr + 16, ["GLOBAL", "CHROME"]);
      if (A.load.Bool(ptr + 22)) {
        x["isExtensionAction"] = A.load.Bool(ptr + 20);
      } else {
        delete x["isExtensionAction"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ControlledInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["text"]);
      }
    },
    "load_ControlledInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["text"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ErrorType": (ref: heap.Ref<string>): number => {
      const idx = ["MANIFEST", "RUNTIME"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DeleteExtensionErrorsProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["extensionId"]);
        A.store.Ref(ptr + 4, x["errorIds"]);
        A.store.Enum(ptr + 8, ["MANIFEST", "RUNTIME"].indexOf(x["type"] as string));
      }
    },
    "load_DeleteExtensionErrorsProperties": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["extensionId"] = A.load.Ref(ptr + 0, undefined);
      x["errorIds"] = A.load.Ref(ptr + 4, undefined);
      x["type"] = A.load.Enum(ptr + 8, ["MANIFEST", "RUNTIME"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DependentExtension": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["name"]);
      }
    },
    "load_DependentExtension": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DisableReasons": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 7, false);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Bool(ptr + 8, "suspiciousInstall" in x ? true : false);
        A.store.Bool(ptr + 0, x["suspiciousInstall"] ? true : false);
        A.store.Bool(ptr + 9, "corruptInstall" in x ? true : false);
        A.store.Bool(ptr + 1, x["corruptInstall"] ? true : false);
        A.store.Bool(ptr + 10, "updateRequired" in x ? true : false);
        A.store.Bool(ptr + 2, x["updateRequired"] ? true : false);
        A.store.Bool(ptr + 11, "publishedInStoreRequired" in x ? true : false);
        A.store.Bool(ptr + 3, x["publishedInStoreRequired"] ? true : false);
        A.store.Bool(ptr + 12, "blockedByPolicy" in x ? true : false);
        A.store.Bool(ptr + 4, x["blockedByPolicy"] ? true : false);
        A.store.Bool(ptr + 13, "reloading" in x ? true : false);
        A.store.Bool(ptr + 5, x["reloading"] ? true : false);
        A.store.Bool(ptr + 14, "custodianApprovalRequired" in x ? true : false);
        A.store.Bool(ptr + 6, x["custodianApprovalRequired"] ? true : false);
        A.store.Bool(ptr + 15, "parentDisabledPermissions" in x ? true : false);
        A.store.Bool(ptr + 7, x["parentDisabledPermissions"] ? true : false);
      }
    },
    "load_DisableReasons": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["suspiciousInstall"] = A.load.Bool(ptr + 0);
      } else {
        delete x["suspiciousInstall"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["corruptInstall"] = A.load.Bool(ptr + 1);
      } else {
        delete x["corruptInstall"];
      }
      if (A.load.Bool(ptr + 10)) {
        x["updateRequired"] = A.load.Bool(ptr + 2);
      } else {
        delete x["updateRequired"];
      }
      if (A.load.Bool(ptr + 11)) {
        x["publishedInStoreRequired"] = A.load.Bool(ptr + 3);
      } else {
        delete x["publishedInStoreRequired"];
      }
      if (A.load.Bool(ptr + 12)) {
        x["blockedByPolicy"] = A.load.Bool(ptr + 4);
      } else {
        delete x["blockedByPolicy"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["reloading"] = A.load.Bool(ptr + 5);
      } else {
        delete x["reloading"];
      }
      if (A.load.Bool(ptr + 14)) {
        x["custodianApprovalRequired"] = A.load.Bool(ptr + 6);
      } else {
        delete x["custodianApprovalRequired"];
      }
      if (A.load.Bool(ptr + 15)) {
        x["parentDisabledPermissions"] = A.load.Bool(ptr + 7);
      } else {
        delete x["parentDisabledPermissions"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ErrorFileSource": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["beforeHighlight"]);
        A.store.Ref(ptr + 4, x["highlight"]);
        A.store.Ref(ptr + 8, x["afterHighlight"]);
      }
    },
    "load_ErrorFileSource": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["beforeHighlight"] = A.load.Ref(ptr + 0, undefined);
      x["highlight"] = A.load.Ref(ptr + 4, undefined);
      x["afterHighlight"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ErrorLevel": (ref: heap.Ref<string>): number => {
      const idx = ["LOG", "WARN", "ERROR"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_EventType": (ref: heap.Ref<string>): number => {
      const idx = [
        "INSTALLED",
        "UNINSTALLED",
        "LOADED",
        "UNLOADED",
        "VIEW_REGISTERED",
        "VIEW_UNREGISTERED",
        "ERROR_ADDED",
        "ERRORS_REMOVED",
        "PREFS_CHANGED",
        "WARNINGS_CHANGED",
        "COMMAND_ADDED",
        "COMMAND_REMOVED",
        "PERMISSIONS_CHANGED",
        "SERVICE_WORKER_STARTED",
        "SERVICE_WORKER_STOPPED",
        "CONFIGURATION_CHANGED",
        "PINNED_ACTIONS_CHANGED",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SafetyCheckStrings": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["panelString"]);
        A.store.Ref(ptr + 4, x["detailString"]);
      }
    },
    "load_SafetyCheckStrings": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["panelString"] = A.load.Ref(ptr + 0, undefined);
      x["detailString"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_HomePage": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 6, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 6, true);
        A.store.Ref(ptr + 0, x["url"]);
        A.store.Bool(ptr + 5, "specified" in x ? true : false);
        A.store.Bool(ptr + 4, x["specified"] ? true : false);
      }
    },
    "load_HomePage": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 5)) {
        x["specified"] = A.load.Bool(ptr + 4);
      } else {
        delete x["specified"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_Location": (ref: heap.Ref<string>): number => {
      const idx = ["FROM_STORE", "UNPACKED", "THIRD_PARTY", "INSTALLED_BY_DEFAULT", "UNKNOWN"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ManifestError": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 34, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 33, false);
        A.store.Int32(ptr + 20, 0);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
      } else {
        A.store.Bool(ptr + 34, true);
        A.store.Enum(ptr + 0, ["MANIFEST", "RUNTIME"].indexOf(x["type"] as string));
        A.store.Ref(ptr + 4, x["extensionId"]);
        A.store.Bool(ptr + 32, "fromIncognito" in x ? true : false);
        A.store.Bool(ptr + 8, x["fromIncognito"] ? true : false);
        A.store.Ref(ptr + 12, x["source"]);
        A.store.Ref(ptr + 16, x["message"]);
        A.store.Bool(ptr + 33, "id" in x ? true : false);
        A.store.Int32(ptr + 20, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Ref(ptr + 24, x["manifestKey"]);
        A.store.Ref(ptr + 28, x["manifestSpecific"]);
      }
    },
    "load_ManifestError": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, ["MANIFEST", "RUNTIME"]);
      x["extensionId"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 32)) {
        x["fromIncognito"] = A.load.Bool(ptr + 8);
      } else {
        delete x["fromIncognito"];
      }
      x["source"] = A.load.Ref(ptr + 12, undefined);
      x["message"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 33)) {
        x["id"] = A.load.Int32(ptr + 20);
      } else {
        delete x["id"];
      }
      x["manifestKey"] = A.load.Ref(ptr + 24, undefined);
      x["manifestSpecific"] = A.load.Ref(ptr + 28, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OptionsPage": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "openInTab" in x ? true : false);
        A.store.Bool(ptr + 0, x["openInTab"] ? true : false);
        A.store.Ref(ptr + 4, x["url"]);
      }
    },
    "load_OptionsPage": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["openInTab"] = A.load.Bool(ptr + 0);
      } else {
        delete x["openInTab"];
      }
      x["url"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Permission": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["message"]);
        A.store.Ref(ptr + 4, x["submessages"]);
      }
    },
    "load_Permission": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["message"] = A.load.Ref(ptr + 0, undefined);
      x["submessages"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_HostAccess": (ref: heap.Ref<string>): number => {
      const idx = ["ON_CLICK", "ON_SPECIFIC_SITES", "ON_ALL_SITES"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SiteControl": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 6, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 6, true);
        A.store.Ref(ptr + 0, x["host"]);
        A.store.Bool(ptr + 5, "granted" in x ? true : false);
        A.store.Bool(ptr + 4, x["granted"] ? true : false);
      }
    },
    "load_SiteControl": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["host"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 5)) {
        x["granted"] = A.load.Bool(ptr + 4);
      } else {
        delete x["granted"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RuntimeHostPermissions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 0, false);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "hasAllHosts" in x ? true : false);
        A.store.Bool(ptr + 0, x["hasAllHosts"] ? true : false);
        A.store.Enum(ptr + 4, ["ON_CLICK", "ON_SPECIFIC_SITES", "ON_ALL_SITES"].indexOf(x["hostAccess"] as string));
        A.store.Ref(ptr + 8, x["hosts"]);
      }
    },
    "load_RuntimeHostPermissions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["hasAllHosts"] = A.load.Bool(ptr + 0);
      } else {
        delete x["hasAllHosts"];
      }
      x["hostAccess"] = A.load.Enum(ptr + 4, ["ON_CLICK", "ON_SPECIFIC_SITES", "ON_ALL_SITES"]);
      x["hosts"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Permissions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 4 + 13, false);
        A.store.Bool(ptr + 4 + 12, false);
        A.store.Bool(ptr + 4 + 0, false);
        A.store.Enum(ptr + 4 + 4, -1);
        A.store.Ref(ptr + 4 + 8, undefined);
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 18, false);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Ref(ptr + 0, x["simplePermissions"]);

        if (typeof x["runtimeHostPermissions"] === "undefined") {
          A.store.Bool(ptr + 4 + 13, false);
          A.store.Bool(ptr + 4 + 12, false);
          A.store.Bool(ptr + 4 + 0, false);
          A.store.Enum(ptr + 4 + 4, -1);
          A.store.Ref(ptr + 4 + 8, undefined);
        } else {
          A.store.Bool(ptr + 4 + 13, true);
          A.store.Bool(ptr + 4 + 12, "hasAllHosts" in x["runtimeHostPermissions"] ? true : false);
          A.store.Bool(ptr + 4 + 0, x["runtimeHostPermissions"]["hasAllHosts"] ? true : false);
          A.store.Enum(
            ptr + 4 + 4,
            ["ON_CLICK", "ON_SPECIFIC_SITES", "ON_ALL_SITES"].indexOf(
              x["runtimeHostPermissions"]["hostAccess"] as string
            )
          );
          A.store.Ref(ptr + 4 + 8, x["runtimeHostPermissions"]["hosts"]);
        }
        A.store.Bool(ptr + 19, "canAccessSiteData" in x ? true : false);
        A.store.Bool(ptr + 18, x["canAccessSiteData"] ? true : false);
      }
    },
    "load_Permissions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["simplePermissions"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 4 + 13)) {
        x["runtimeHostPermissions"] = {};
        if (A.load.Bool(ptr + 4 + 12)) {
          x["runtimeHostPermissions"]["hasAllHosts"] = A.load.Bool(ptr + 4 + 0);
        } else {
          delete x["runtimeHostPermissions"]["hasAllHosts"];
        }
        x["runtimeHostPermissions"]["hostAccess"] = A.load.Enum(ptr + 4 + 4, [
          "ON_CLICK",
          "ON_SPECIFIC_SITES",
          "ON_ALL_SITES",
        ]);
        x["runtimeHostPermissions"]["hosts"] = A.load.Ref(ptr + 4 + 8, undefined);
      } else {
        delete x["runtimeHostPermissions"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["canAccessSiteData"] = A.load.Bool(ptr + 18);
      } else {
        delete x["canAccessSiteData"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_StackFrame": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Bool(ptr + 16, "lineNumber" in x ? true : false);
        A.store.Int32(ptr + 0, x["lineNumber"] === undefined ? 0 : (x["lineNumber"] as number));
        A.store.Bool(ptr + 17, "columnNumber" in x ? true : false);
        A.store.Int32(ptr + 4, x["columnNumber"] === undefined ? 0 : (x["columnNumber"] as number));
        A.store.Ref(ptr + 8, x["url"]);
        A.store.Ref(ptr + 12, x["functionName"]);
      }
    },
    "load_StackFrame": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["lineNumber"] = A.load.Int32(ptr + 0);
      } else {
        delete x["lineNumber"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["columnNumber"] = A.load.Int32(ptr + 4);
      } else {
        delete x["columnNumber"];
      }
      x["url"] = A.load.Ref(ptr + 8, undefined);
      x["functionName"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RuntimeError": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 58, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 52, false);
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 53, false);
        A.store.Int32(ptr + 20, 0);
        A.store.Enum(ptr + 24, -1);
        A.store.Ref(ptr + 28, undefined);
        A.store.Bool(ptr + 54, false);
        A.store.Int32(ptr + 32, 0);
        A.store.Bool(ptr + 55, false);
        A.store.Int32(ptr + 36, 0);
        A.store.Bool(ptr + 56, false);
        A.store.Int32(ptr + 40, 0);
        A.store.Bool(ptr + 57, false);
        A.store.Bool(ptr + 44, false);
        A.store.Ref(ptr + 48, undefined);
      } else {
        A.store.Bool(ptr + 58, true);
        A.store.Enum(ptr + 0, ["MANIFEST", "RUNTIME"].indexOf(x["type"] as string));
        A.store.Ref(ptr + 4, x["extensionId"]);
        A.store.Bool(ptr + 52, "fromIncognito" in x ? true : false);
        A.store.Bool(ptr + 8, x["fromIncognito"] ? true : false);
        A.store.Ref(ptr + 12, x["source"]);
        A.store.Ref(ptr + 16, x["message"]);
        A.store.Bool(ptr + 53, "id" in x ? true : false);
        A.store.Int32(ptr + 20, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Enum(ptr + 24, ["LOG", "WARN", "ERROR"].indexOf(x["severity"] as string));
        A.store.Ref(ptr + 28, x["contextUrl"]);
        A.store.Bool(ptr + 54, "occurrences" in x ? true : false);
        A.store.Int32(ptr + 32, x["occurrences"] === undefined ? 0 : (x["occurrences"] as number));
        A.store.Bool(ptr + 55, "renderViewId" in x ? true : false);
        A.store.Int32(ptr + 36, x["renderViewId"] === undefined ? 0 : (x["renderViewId"] as number));
        A.store.Bool(ptr + 56, "renderProcessId" in x ? true : false);
        A.store.Int32(ptr + 40, x["renderProcessId"] === undefined ? 0 : (x["renderProcessId"] as number));
        A.store.Bool(ptr + 57, "canInspect" in x ? true : false);
        A.store.Bool(ptr + 44, x["canInspect"] ? true : false);
        A.store.Ref(ptr + 48, x["stackTrace"]);
      }
    },
    "load_RuntimeError": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, ["MANIFEST", "RUNTIME"]);
      x["extensionId"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 52)) {
        x["fromIncognito"] = A.load.Bool(ptr + 8);
      } else {
        delete x["fromIncognito"];
      }
      x["source"] = A.load.Ref(ptr + 12, undefined);
      x["message"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 53)) {
        x["id"] = A.load.Int32(ptr + 20);
      } else {
        delete x["id"];
      }
      x["severity"] = A.load.Enum(ptr + 24, ["LOG", "WARN", "ERROR"]);
      x["contextUrl"] = A.load.Ref(ptr + 28, undefined);
      if (A.load.Bool(ptr + 54)) {
        x["occurrences"] = A.load.Int32(ptr + 32);
      } else {
        delete x["occurrences"];
      }
      if (A.load.Bool(ptr + 55)) {
        x["renderViewId"] = A.load.Int32(ptr + 36);
      } else {
        delete x["renderViewId"];
      }
      if (A.load.Bool(ptr + 56)) {
        x["renderProcessId"] = A.load.Int32(ptr + 40);
      } else {
        delete x["renderProcessId"];
      }
      if (A.load.Bool(ptr + 57)) {
        x["canInspect"] = A.load.Bool(ptr + 44);
      } else {
        delete x["canInspect"];
      }
      x["stackTrace"] = A.load.Ref(ptr + 48, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ExtensionState": (ref: heap.Ref<string>): number => {
      const idx = ["ENABLED", "DISABLED", "TERMINATED", "BLACKLISTED"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ExtensionType": (ref: heap.Ref<string>): number => {
      const idx = ["HOSTED_APP", "PLATFORM_APP", "LEGACY_PACKAGED_APP", "EXTENSION", "THEME", "SHARED_MODULE"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ViewType": (ref: heap.Ref<string>): number => {
      const idx = [
        "APP_WINDOW",
        "BACKGROUND_CONTENTS",
        "COMPONENT",
        "EXTENSION_BACKGROUND_PAGE",
        "EXTENSION_DIALOG",
        "EXTENSION_GUEST",
        "EXTENSION_POPUP",
        "EXTENSION_SERVICE_WORKER_BACKGROUND",
        "TAB_CONTENTS",
        "OFFSCREEN_DOCUMENT",
        "EXTENSION_SIDE_PANEL",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ExtensionView": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 24, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 20, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 21, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 22, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 23, false);
        A.store.Bool(ptr + 13, false);
        A.store.Enum(ptr + 16, -1);
      } else {
        A.store.Bool(ptr + 24, true);
        A.store.Ref(ptr + 0, x["url"]);
        A.store.Bool(ptr + 20, "renderProcessId" in x ? true : false);
        A.store.Int32(ptr + 4, x["renderProcessId"] === undefined ? 0 : (x["renderProcessId"] as number));
        A.store.Bool(ptr + 21, "renderViewId" in x ? true : false);
        A.store.Int32(ptr + 8, x["renderViewId"] === undefined ? 0 : (x["renderViewId"] as number));
        A.store.Bool(ptr + 22, "incognito" in x ? true : false);
        A.store.Bool(ptr + 12, x["incognito"] ? true : false);
        A.store.Bool(ptr + 23, "isIframe" in x ? true : false);
        A.store.Bool(ptr + 13, x["isIframe"] ? true : false);
        A.store.Enum(
          ptr + 16,
          [
            "APP_WINDOW",
            "BACKGROUND_CONTENTS",
            "COMPONENT",
            "EXTENSION_BACKGROUND_PAGE",
            "EXTENSION_DIALOG",
            "EXTENSION_GUEST",
            "EXTENSION_POPUP",
            "EXTENSION_SERVICE_WORKER_BACKGROUND",
            "TAB_CONTENTS",
            "OFFSCREEN_DOCUMENT",
            "EXTENSION_SIDE_PANEL",
          ].indexOf(x["type"] as string)
        );
      }
    },
    "load_ExtensionView": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 20)) {
        x["renderProcessId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["renderProcessId"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["renderViewId"] = A.load.Int32(ptr + 8);
      } else {
        delete x["renderViewId"];
      }
      if (A.load.Bool(ptr + 22)) {
        x["incognito"] = A.load.Bool(ptr + 12);
      } else {
        delete x["incognito"];
      }
      if (A.load.Bool(ptr + 23)) {
        x["isIframe"] = A.load.Bool(ptr + 13);
      } else {
        delete x["isIframe"];
      }
      x["type"] = A.load.Enum(ptr + 16, [
        "APP_WINDOW",
        "BACKGROUND_CONTENTS",
        "COMPONENT",
        "EXTENSION_BACKGROUND_PAGE",
        "EXTENSION_DIALOG",
        "EXTENSION_GUEST",
        "EXTENSION_POPUP",
        "EXTENSION_SERVICE_WORKER_BACKGROUND",
        "TAB_CONTENTS",
        "OFFSCREEN_DOCUMENT",
        "EXTENSION_SIDE_PANEL",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ExtensionInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 215, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 4 + 8, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Ref(ptr + 4 + 4, undefined);
        A.store.Ref(ptr + 16, undefined);

        A.store.Bool(ptr + 20 + 4, false);
        A.store.Ref(ptr + 20 + 0, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);

        A.store.Bool(ptr + 36 + 16, false);
        A.store.Bool(ptr + 36 + 8, false);
        A.store.Bool(ptr + 36 + 0, false);
        A.store.Bool(ptr + 36 + 9, false);
        A.store.Bool(ptr + 36 + 1, false);
        A.store.Bool(ptr + 36 + 10, false);
        A.store.Bool(ptr + 36 + 2, false);
        A.store.Bool(ptr + 36 + 11, false);
        A.store.Bool(ptr + 36 + 3, false);
        A.store.Bool(ptr + 36 + 12, false);
        A.store.Bool(ptr + 36 + 4, false);
        A.store.Bool(ptr + 36 + 13, false);
        A.store.Bool(ptr + 36 + 5, false);
        A.store.Bool(ptr + 36 + 14, false);
        A.store.Bool(ptr + 36 + 6, false);
        A.store.Bool(ptr + 36 + 15, false);
        A.store.Bool(ptr + 36 + 7, false);

        A.store.Bool(ptr + 53 + 4, false);
        A.store.Bool(ptr + 53 + 2, false);
        A.store.Bool(ptr + 53 + 0, false);
        A.store.Bool(ptr + 53 + 3, false);
        A.store.Bool(ptr + 53 + 1, false);

        A.store.Bool(ptr + 58 + 4, false);
        A.store.Bool(ptr + 58 + 2, false);
        A.store.Bool(ptr + 58 + 0, false);
        A.store.Bool(ptr + 58 + 3, false);
        A.store.Bool(ptr + 58 + 1, false);

        A.store.Bool(ptr + 64 + 6, false);
        A.store.Ref(ptr + 64 + 0, undefined);
        A.store.Bool(ptr + 64 + 5, false);
        A.store.Bool(ptr + 64 + 4, false);
        A.store.Ref(ptr + 72, undefined);
        A.store.Ref(ptr + 76, undefined);

        A.store.Bool(ptr + 80 + 4, false);
        A.store.Bool(ptr + 80 + 2, false);
        A.store.Bool(ptr + 80 + 0, false);
        A.store.Bool(ptr + 80 + 3, false);
        A.store.Bool(ptr + 80 + 1, false);
        A.store.Ref(ptr + 88, undefined);
        A.store.Ref(ptr + 92, undefined);
        A.store.Enum(ptr + 96, -1);
        A.store.Ref(ptr + 100, undefined);
        A.store.Ref(ptr + 104, undefined);
        A.store.Ref(ptr + 108, undefined);
        A.store.Bool(ptr + 208, false);
        A.store.Bool(ptr + 112, false);
        A.store.Ref(ptr + 116, undefined);
        A.store.Bool(ptr + 209, false);
        A.store.Bool(ptr + 120, false);

        A.store.Bool(ptr + 124 + 9, false);
        A.store.Bool(ptr + 124 + 8, false);
        A.store.Bool(ptr + 124 + 0, false);
        A.store.Ref(ptr + 124 + 4, undefined);
        A.store.Ref(ptr + 136, undefined);

        A.store.Bool(ptr + 140 + 20, false);
        A.store.Ref(ptr + 140 + 0, undefined);

        A.store.Bool(ptr + 140 + 4 + 13, false);
        A.store.Bool(ptr + 140 + 4 + 12, false);
        A.store.Bool(ptr + 140 + 4 + 0, false);
        A.store.Enum(ptr + 140 + 4 + 4, -1);
        A.store.Ref(ptr + 140 + 4 + 8, undefined);
        A.store.Bool(ptr + 140 + 19, false);
        A.store.Bool(ptr + 140 + 18, false);
        A.store.Ref(ptr + 164, undefined);
        A.store.Ref(ptr + 168, undefined);
        A.store.Ref(ptr + 172, undefined);
        A.store.Enum(ptr + 176, -1);
        A.store.Enum(ptr + 180, -1);
        A.store.Ref(ptr + 184, undefined);
        A.store.Bool(ptr + 210, false);
        A.store.Bool(ptr + 188, false);
        A.store.Ref(ptr + 192, undefined);
        A.store.Ref(ptr + 196, undefined);
        A.store.Ref(ptr + 200, undefined);
        A.store.Bool(ptr + 211, false);
        A.store.Bool(ptr + 204, false);
        A.store.Bool(ptr + 212, false);
        A.store.Bool(ptr + 205, false);
        A.store.Bool(ptr + 213, false);
        A.store.Bool(ptr + 206, false);
        A.store.Bool(ptr + 214, false);
        A.store.Bool(ptr + 207, false);
      } else {
        A.store.Bool(ptr + 215, true);
        A.store.Ref(ptr + 0, x["blacklistText"]);

        if (typeof x["safetyCheckText"] === "undefined") {
          A.store.Bool(ptr + 4 + 8, false);
          A.store.Ref(ptr + 4 + 0, undefined);
          A.store.Ref(ptr + 4 + 4, undefined);
        } else {
          A.store.Bool(ptr + 4 + 8, true);
          A.store.Ref(ptr + 4 + 0, x["safetyCheckText"]["panelString"]);
          A.store.Ref(ptr + 4 + 4, x["safetyCheckText"]["detailString"]);
        }
        A.store.Ref(ptr + 16, x["commands"]);

        if (typeof x["controlledInfo"] === "undefined") {
          A.store.Bool(ptr + 20 + 4, false);
          A.store.Ref(ptr + 20 + 0, undefined);
        } else {
          A.store.Bool(ptr + 20 + 4, true);
          A.store.Ref(ptr + 20 + 0, x["controlledInfo"]["text"]);
        }
        A.store.Ref(ptr + 28, x["dependentExtensions"]);
        A.store.Ref(ptr + 32, x["description"]);

        if (typeof x["disableReasons"] === "undefined") {
          A.store.Bool(ptr + 36 + 16, false);
          A.store.Bool(ptr + 36 + 8, false);
          A.store.Bool(ptr + 36 + 0, false);
          A.store.Bool(ptr + 36 + 9, false);
          A.store.Bool(ptr + 36 + 1, false);
          A.store.Bool(ptr + 36 + 10, false);
          A.store.Bool(ptr + 36 + 2, false);
          A.store.Bool(ptr + 36 + 11, false);
          A.store.Bool(ptr + 36 + 3, false);
          A.store.Bool(ptr + 36 + 12, false);
          A.store.Bool(ptr + 36 + 4, false);
          A.store.Bool(ptr + 36 + 13, false);
          A.store.Bool(ptr + 36 + 5, false);
          A.store.Bool(ptr + 36 + 14, false);
          A.store.Bool(ptr + 36 + 6, false);
          A.store.Bool(ptr + 36 + 15, false);
          A.store.Bool(ptr + 36 + 7, false);
        } else {
          A.store.Bool(ptr + 36 + 16, true);
          A.store.Bool(ptr + 36 + 8, "suspiciousInstall" in x["disableReasons"] ? true : false);
          A.store.Bool(ptr + 36 + 0, x["disableReasons"]["suspiciousInstall"] ? true : false);
          A.store.Bool(ptr + 36 + 9, "corruptInstall" in x["disableReasons"] ? true : false);
          A.store.Bool(ptr + 36 + 1, x["disableReasons"]["corruptInstall"] ? true : false);
          A.store.Bool(ptr + 36 + 10, "updateRequired" in x["disableReasons"] ? true : false);
          A.store.Bool(ptr + 36 + 2, x["disableReasons"]["updateRequired"] ? true : false);
          A.store.Bool(ptr + 36 + 11, "publishedInStoreRequired" in x["disableReasons"] ? true : false);
          A.store.Bool(ptr + 36 + 3, x["disableReasons"]["publishedInStoreRequired"] ? true : false);
          A.store.Bool(ptr + 36 + 12, "blockedByPolicy" in x["disableReasons"] ? true : false);
          A.store.Bool(ptr + 36 + 4, x["disableReasons"]["blockedByPolicy"] ? true : false);
          A.store.Bool(ptr + 36 + 13, "reloading" in x["disableReasons"] ? true : false);
          A.store.Bool(ptr + 36 + 5, x["disableReasons"]["reloading"] ? true : false);
          A.store.Bool(ptr + 36 + 14, "custodianApprovalRequired" in x["disableReasons"] ? true : false);
          A.store.Bool(ptr + 36 + 6, x["disableReasons"]["custodianApprovalRequired"] ? true : false);
          A.store.Bool(ptr + 36 + 15, "parentDisabledPermissions" in x["disableReasons"] ? true : false);
          A.store.Bool(ptr + 36 + 7, x["disableReasons"]["parentDisabledPermissions"] ? true : false);
        }

        if (typeof x["errorCollection"] === "undefined") {
          A.store.Bool(ptr + 53 + 4, false);
          A.store.Bool(ptr + 53 + 2, false);
          A.store.Bool(ptr + 53 + 0, false);
          A.store.Bool(ptr + 53 + 3, false);
          A.store.Bool(ptr + 53 + 1, false);
        } else {
          A.store.Bool(ptr + 53 + 4, true);
          A.store.Bool(ptr + 53 + 2, "isEnabled" in x["errorCollection"] ? true : false);
          A.store.Bool(ptr + 53 + 0, x["errorCollection"]["isEnabled"] ? true : false);
          A.store.Bool(ptr + 53 + 3, "isActive" in x["errorCollection"] ? true : false);
          A.store.Bool(ptr + 53 + 1, x["errorCollection"]["isActive"] ? true : false);
        }

        if (typeof x["fileAccess"] === "undefined") {
          A.store.Bool(ptr + 58 + 4, false);
          A.store.Bool(ptr + 58 + 2, false);
          A.store.Bool(ptr + 58 + 0, false);
          A.store.Bool(ptr + 58 + 3, false);
          A.store.Bool(ptr + 58 + 1, false);
        } else {
          A.store.Bool(ptr + 58 + 4, true);
          A.store.Bool(ptr + 58 + 2, "isEnabled" in x["fileAccess"] ? true : false);
          A.store.Bool(ptr + 58 + 0, x["fileAccess"]["isEnabled"] ? true : false);
          A.store.Bool(ptr + 58 + 3, "isActive" in x["fileAccess"] ? true : false);
          A.store.Bool(ptr + 58 + 1, x["fileAccess"]["isActive"] ? true : false);
        }

        if (typeof x["homePage"] === "undefined") {
          A.store.Bool(ptr + 64 + 6, false);
          A.store.Ref(ptr + 64 + 0, undefined);
          A.store.Bool(ptr + 64 + 5, false);
          A.store.Bool(ptr + 64 + 4, false);
        } else {
          A.store.Bool(ptr + 64 + 6, true);
          A.store.Ref(ptr + 64 + 0, x["homePage"]["url"]);
          A.store.Bool(ptr + 64 + 5, "specified" in x["homePage"] ? true : false);
          A.store.Bool(ptr + 64 + 4, x["homePage"]["specified"] ? true : false);
        }
        A.store.Ref(ptr + 72, x["iconUrl"]);
        A.store.Ref(ptr + 76, x["id"]);

        if (typeof x["incognitoAccess"] === "undefined") {
          A.store.Bool(ptr + 80 + 4, false);
          A.store.Bool(ptr + 80 + 2, false);
          A.store.Bool(ptr + 80 + 0, false);
          A.store.Bool(ptr + 80 + 3, false);
          A.store.Bool(ptr + 80 + 1, false);
        } else {
          A.store.Bool(ptr + 80 + 4, true);
          A.store.Bool(ptr + 80 + 2, "isEnabled" in x["incognitoAccess"] ? true : false);
          A.store.Bool(ptr + 80 + 0, x["incognitoAccess"]["isEnabled"] ? true : false);
          A.store.Bool(ptr + 80 + 3, "isActive" in x["incognitoAccess"] ? true : false);
          A.store.Bool(ptr + 80 + 1, x["incognitoAccess"]["isActive"] ? true : false);
        }
        A.store.Ref(ptr + 88, x["installWarnings"]);
        A.store.Ref(ptr + 92, x["launchUrl"]);
        A.store.Enum(
          ptr + 96,
          ["FROM_STORE", "UNPACKED", "THIRD_PARTY", "INSTALLED_BY_DEFAULT", "UNKNOWN"].indexOf(x["location"] as string)
        );
        A.store.Ref(ptr + 100, x["locationText"]);
        A.store.Ref(ptr + 104, x["manifestErrors"]);
        A.store.Ref(ptr + 108, x["manifestHomePageUrl"]);
        A.store.Bool(ptr + 208, "mustRemainInstalled" in x ? true : false);
        A.store.Bool(ptr + 112, x["mustRemainInstalled"] ? true : false);
        A.store.Ref(ptr + 116, x["name"]);
        A.store.Bool(ptr + 209, "offlineEnabled" in x ? true : false);
        A.store.Bool(ptr + 120, x["offlineEnabled"] ? true : false);

        if (typeof x["optionsPage"] === "undefined") {
          A.store.Bool(ptr + 124 + 9, false);
          A.store.Bool(ptr + 124 + 8, false);
          A.store.Bool(ptr + 124 + 0, false);
          A.store.Ref(ptr + 124 + 4, undefined);
        } else {
          A.store.Bool(ptr + 124 + 9, true);
          A.store.Bool(ptr + 124 + 8, "openInTab" in x["optionsPage"] ? true : false);
          A.store.Bool(ptr + 124 + 0, x["optionsPage"]["openInTab"] ? true : false);
          A.store.Ref(ptr + 124 + 4, x["optionsPage"]["url"]);
        }
        A.store.Ref(ptr + 136, x["path"]);

        if (typeof x["permissions"] === "undefined") {
          A.store.Bool(ptr + 140 + 20, false);
          A.store.Ref(ptr + 140 + 0, undefined);

          A.store.Bool(ptr + 140 + 4 + 13, false);
          A.store.Bool(ptr + 140 + 4 + 12, false);
          A.store.Bool(ptr + 140 + 4 + 0, false);
          A.store.Enum(ptr + 140 + 4 + 4, -1);
          A.store.Ref(ptr + 140 + 4 + 8, undefined);
          A.store.Bool(ptr + 140 + 19, false);
          A.store.Bool(ptr + 140 + 18, false);
        } else {
          A.store.Bool(ptr + 140 + 20, true);
          A.store.Ref(ptr + 140 + 0, x["permissions"]["simplePermissions"]);

          if (typeof x["permissions"]["runtimeHostPermissions"] === "undefined") {
            A.store.Bool(ptr + 140 + 4 + 13, false);
            A.store.Bool(ptr + 140 + 4 + 12, false);
            A.store.Bool(ptr + 140 + 4 + 0, false);
            A.store.Enum(ptr + 140 + 4 + 4, -1);
            A.store.Ref(ptr + 140 + 4 + 8, undefined);
          } else {
            A.store.Bool(ptr + 140 + 4 + 13, true);
            A.store.Bool(
              ptr + 140 + 4 + 12,
              "hasAllHosts" in x["permissions"]["runtimeHostPermissions"] ? true : false
            );
            A.store.Bool(ptr + 140 + 4 + 0, x["permissions"]["runtimeHostPermissions"]["hasAllHosts"] ? true : false);
            A.store.Enum(
              ptr + 140 + 4 + 4,
              ["ON_CLICK", "ON_SPECIFIC_SITES", "ON_ALL_SITES"].indexOf(
                x["permissions"]["runtimeHostPermissions"]["hostAccess"] as string
              )
            );
            A.store.Ref(ptr + 140 + 4 + 8, x["permissions"]["runtimeHostPermissions"]["hosts"]);
          }
          A.store.Bool(ptr + 140 + 19, "canAccessSiteData" in x["permissions"] ? true : false);
          A.store.Bool(ptr + 140 + 18, x["permissions"]["canAccessSiteData"] ? true : false);
        }
        A.store.Ref(ptr + 164, x["prettifiedPath"]);
        A.store.Ref(ptr + 168, x["runtimeErrors"]);
        A.store.Ref(ptr + 172, x["runtimeWarnings"]);
        A.store.Enum(ptr + 176, ["ENABLED", "DISABLED", "TERMINATED", "BLACKLISTED"].indexOf(x["state"] as string));
        A.store.Enum(
          ptr + 180,
          ["HOSTED_APP", "PLATFORM_APP", "LEGACY_PACKAGED_APP", "EXTENSION", "THEME", "SHARED_MODULE"].indexOf(
            x["type"] as string
          )
        );
        A.store.Ref(ptr + 184, x["updateUrl"]);
        A.store.Bool(ptr + 210, "userMayModify" in x ? true : false);
        A.store.Bool(ptr + 188, x["userMayModify"] ? true : false);
        A.store.Ref(ptr + 192, x["version"]);
        A.store.Ref(ptr + 196, x["views"]);
        A.store.Ref(ptr + 200, x["webStoreUrl"]);
        A.store.Bool(ptr + 211, "showSafeBrowsingAllowlistWarning" in x ? true : false);
        A.store.Bool(ptr + 204, x["showSafeBrowsingAllowlistWarning"] ? true : false);
        A.store.Bool(ptr + 212, "showAccessRequestsInToolbar" in x ? true : false);
        A.store.Bool(ptr + 205, x["showAccessRequestsInToolbar"] ? true : false);
        A.store.Bool(ptr + 213, "acknowledgeSafetyCheckWarning" in x ? true : false);
        A.store.Bool(ptr + 206, x["acknowledgeSafetyCheckWarning"] ? true : false);
        A.store.Bool(ptr + 214, "pinnedToToolbar" in x ? true : false);
        A.store.Bool(ptr + 207, x["pinnedToToolbar"] ? true : false);
      }
    },
    "load_ExtensionInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["blacklistText"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 4 + 8)) {
        x["safetyCheckText"] = {};
        x["safetyCheckText"]["panelString"] = A.load.Ref(ptr + 4 + 0, undefined);
        x["safetyCheckText"]["detailString"] = A.load.Ref(ptr + 4 + 4, undefined);
      } else {
        delete x["safetyCheckText"];
      }
      x["commands"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 20 + 4)) {
        x["controlledInfo"] = {};
        x["controlledInfo"]["text"] = A.load.Ref(ptr + 20 + 0, undefined);
      } else {
        delete x["controlledInfo"];
      }
      x["dependentExtensions"] = A.load.Ref(ptr + 28, undefined);
      x["description"] = A.load.Ref(ptr + 32, undefined);
      if (A.load.Bool(ptr + 36 + 16)) {
        x["disableReasons"] = {};
        if (A.load.Bool(ptr + 36 + 8)) {
          x["disableReasons"]["suspiciousInstall"] = A.load.Bool(ptr + 36 + 0);
        } else {
          delete x["disableReasons"]["suspiciousInstall"];
        }
        if (A.load.Bool(ptr + 36 + 9)) {
          x["disableReasons"]["corruptInstall"] = A.load.Bool(ptr + 36 + 1);
        } else {
          delete x["disableReasons"]["corruptInstall"];
        }
        if (A.load.Bool(ptr + 36 + 10)) {
          x["disableReasons"]["updateRequired"] = A.load.Bool(ptr + 36 + 2);
        } else {
          delete x["disableReasons"]["updateRequired"];
        }
        if (A.load.Bool(ptr + 36 + 11)) {
          x["disableReasons"]["publishedInStoreRequired"] = A.load.Bool(ptr + 36 + 3);
        } else {
          delete x["disableReasons"]["publishedInStoreRequired"];
        }
        if (A.load.Bool(ptr + 36 + 12)) {
          x["disableReasons"]["blockedByPolicy"] = A.load.Bool(ptr + 36 + 4);
        } else {
          delete x["disableReasons"]["blockedByPolicy"];
        }
        if (A.load.Bool(ptr + 36 + 13)) {
          x["disableReasons"]["reloading"] = A.load.Bool(ptr + 36 + 5);
        } else {
          delete x["disableReasons"]["reloading"];
        }
        if (A.load.Bool(ptr + 36 + 14)) {
          x["disableReasons"]["custodianApprovalRequired"] = A.load.Bool(ptr + 36 + 6);
        } else {
          delete x["disableReasons"]["custodianApprovalRequired"];
        }
        if (A.load.Bool(ptr + 36 + 15)) {
          x["disableReasons"]["parentDisabledPermissions"] = A.load.Bool(ptr + 36 + 7);
        } else {
          delete x["disableReasons"]["parentDisabledPermissions"];
        }
      } else {
        delete x["disableReasons"];
      }
      if (A.load.Bool(ptr + 53 + 4)) {
        x["errorCollection"] = {};
        if (A.load.Bool(ptr + 53 + 2)) {
          x["errorCollection"]["isEnabled"] = A.load.Bool(ptr + 53 + 0);
        } else {
          delete x["errorCollection"]["isEnabled"];
        }
        if (A.load.Bool(ptr + 53 + 3)) {
          x["errorCollection"]["isActive"] = A.load.Bool(ptr + 53 + 1);
        } else {
          delete x["errorCollection"]["isActive"];
        }
      } else {
        delete x["errorCollection"];
      }
      if (A.load.Bool(ptr + 58 + 4)) {
        x["fileAccess"] = {};
        if (A.load.Bool(ptr + 58 + 2)) {
          x["fileAccess"]["isEnabled"] = A.load.Bool(ptr + 58 + 0);
        } else {
          delete x["fileAccess"]["isEnabled"];
        }
        if (A.load.Bool(ptr + 58 + 3)) {
          x["fileAccess"]["isActive"] = A.load.Bool(ptr + 58 + 1);
        } else {
          delete x["fileAccess"]["isActive"];
        }
      } else {
        delete x["fileAccess"];
      }
      if (A.load.Bool(ptr + 64 + 6)) {
        x["homePage"] = {};
        x["homePage"]["url"] = A.load.Ref(ptr + 64 + 0, undefined);
        if (A.load.Bool(ptr + 64 + 5)) {
          x["homePage"]["specified"] = A.load.Bool(ptr + 64 + 4);
        } else {
          delete x["homePage"]["specified"];
        }
      } else {
        delete x["homePage"];
      }
      x["iconUrl"] = A.load.Ref(ptr + 72, undefined);
      x["id"] = A.load.Ref(ptr + 76, undefined);
      if (A.load.Bool(ptr + 80 + 4)) {
        x["incognitoAccess"] = {};
        if (A.load.Bool(ptr + 80 + 2)) {
          x["incognitoAccess"]["isEnabled"] = A.load.Bool(ptr + 80 + 0);
        } else {
          delete x["incognitoAccess"]["isEnabled"];
        }
        if (A.load.Bool(ptr + 80 + 3)) {
          x["incognitoAccess"]["isActive"] = A.load.Bool(ptr + 80 + 1);
        } else {
          delete x["incognitoAccess"]["isActive"];
        }
      } else {
        delete x["incognitoAccess"];
      }
      x["installWarnings"] = A.load.Ref(ptr + 88, undefined);
      x["launchUrl"] = A.load.Ref(ptr + 92, undefined);
      x["location"] = A.load.Enum(ptr + 96, [
        "FROM_STORE",
        "UNPACKED",
        "THIRD_PARTY",
        "INSTALLED_BY_DEFAULT",
        "UNKNOWN",
      ]);
      x["locationText"] = A.load.Ref(ptr + 100, undefined);
      x["manifestErrors"] = A.load.Ref(ptr + 104, undefined);
      x["manifestHomePageUrl"] = A.load.Ref(ptr + 108, undefined);
      if (A.load.Bool(ptr + 208)) {
        x["mustRemainInstalled"] = A.load.Bool(ptr + 112);
      } else {
        delete x["mustRemainInstalled"];
      }
      x["name"] = A.load.Ref(ptr + 116, undefined);
      if (A.load.Bool(ptr + 209)) {
        x["offlineEnabled"] = A.load.Bool(ptr + 120);
      } else {
        delete x["offlineEnabled"];
      }
      if (A.load.Bool(ptr + 124 + 9)) {
        x["optionsPage"] = {};
        if (A.load.Bool(ptr + 124 + 8)) {
          x["optionsPage"]["openInTab"] = A.load.Bool(ptr + 124 + 0);
        } else {
          delete x["optionsPage"]["openInTab"];
        }
        x["optionsPage"]["url"] = A.load.Ref(ptr + 124 + 4, undefined);
      } else {
        delete x["optionsPage"];
      }
      x["path"] = A.load.Ref(ptr + 136, undefined);
      if (A.load.Bool(ptr + 140 + 20)) {
        x["permissions"] = {};
        x["permissions"]["simplePermissions"] = A.load.Ref(ptr + 140 + 0, undefined);
        if (A.load.Bool(ptr + 140 + 4 + 13)) {
          x["permissions"]["runtimeHostPermissions"] = {};
          if (A.load.Bool(ptr + 140 + 4 + 12)) {
            x["permissions"]["runtimeHostPermissions"]["hasAllHosts"] = A.load.Bool(ptr + 140 + 4 + 0);
          } else {
            delete x["permissions"]["runtimeHostPermissions"]["hasAllHosts"];
          }
          x["permissions"]["runtimeHostPermissions"]["hostAccess"] = A.load.Enum(ptr + 140 + 4 + 4, [
            "ON_CLICK",
            "ON_SPECIFIC_SITES",
            "ON_ALL_SITES",
          ]);
          x["permissions"]["runtimeHostPermissions"]["hosts"] = A.load.Ref(ptr + 140 + 4 + 8, undefined);
        } else {
          delete x["permissions"]["runtimeHostPermissions"];
        }
        if (A.load.Bool(ptr + 140 + 19)) {
          x["permissions"]["canAccessSiteData"] = A.load.Bool(ptr + 140 + 18);
        } else {
          delete x["permissions"]["canAccessSiteData"];
        }
      } else {
        delete x["permissions"];
      }
      x["prettifiedPath"] = A.load.Ref(ptr + 164, undefined);
      x["runtimeErrors"] = A.load.Ref(ptr + 168, undefined);
      x["runtimeWarnings"] = A.load.Ref(ptr + 172, undefined);
      x["state"] = A.load.Enum(ptr + 176, ["ENABLED", "DISABLED", "TERMINATED", "BLACKLISTED"]);
      x["type"] = A.load.Enum(ptr + 180, [
        "HOSTED_APP",
        "PLATFORM_APP",
        "LEGACY_PACKAGED_APP",
        "EXTENSION",
        "THEME",
        "SHARED_MODULE",
      ]);
      x["updateUrl"] = A.load.Ref(ptr + 184, undefined);
      if (A.load.Bool(ptr + 210)) {
        x["userMayModify"] = A.load.Bool(ptr + 188);
      } else {
        delete x["userMayModify"];
      }
      x["version"] = A.load.Ref(ptr + 192, undefined);
      x["views"] = A.load.Ref(ptr + 196, undefined);
      x["webStoreUrl"] = A.load.Ref(ptr + 200, undefined);
      if (A.load.Bool(ptr + 211)) {
        x["showSafeBrowsingAllowlistWarning"] = A.load.Bool(ptr + 204);
      } else {
        delete x["showSafeBrowsingAllowlistWarning"];
      }
      if (A.load.Bool(ptr + 212)) {
        x["showAccessRequestsInToolbar"] = A.load.Bool(ptr + 205);
      } else {
        delete x["showAccessRequestsInToolbar"];
      }
      if (A.load.Bool(ptr + 213)) {
        x["acknowledgeSafetyCheckWarning"] = A.load.Bool(ptr + 206);
      } else {
        delete x["acknowledgeSafetyCheckWarning"];
      }
      if (A.load.Bool(ptr + 214)) {
        x["pinnedToToolbar"] = A.load.Bool(ptr + 207);
      } else {
        delete x["pinnedToToolbar"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_EventData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 224, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);

        A.store.Bool(ptr + 8 + 215, false);
        A.store.Ref(ptr + 8 + 0, undefined);

        A.store.Bool(ptr + 8 + 4 + 8, false);
        A.store.Ref(ptr + 8 + 4 + 0, undefined);
        A.store.Ref(ptr + 8 + 4 + 4, undefined);
        A.store.Ref(ptr + 8 + 16, undefined);

        A.store.Bool(ptr + 8 + 20 + 4, false);
        A.store.Ref(ptr + 8 + 20 + 0, undefined);
        A.store.Ref(ptr + 8 + 28, undefined);
        A.store.Ref(ptr + 8 + 32, undefined);

        A.store.Bool(ptr + 8 + 36 + 16, false);
        A.store.Bool(ptr + 8 + 36 + 8, false);
        A.store.Bool(ptr + 8 + 36 + 0, false);
        A.store.Bool(ptr + 8 + 36 + 9, false);
        A.store.Bool(ptr + 8 + 36 + 1, false);
        A.store.Bool(ptr + 8 + 36 + 10, false);
        A.store.Bool(ptr + 8 + 36 + 2, false);
        A.store.Bool(ptr + 8 + 36 + 11, false);
        A.store.Bool(ptr + 8 + 36 + 3, false);
        A.store.Bool(ptr + 8 + 36 + 12, false);
        A.store.Bool(ptr + 8 + 36 + 4, false);
        A.store.Bool(ptr + 8 + 36 + 13, false);
        A.store.Bool(ptr + 8 + 36 + 5, false);
        A.store.Bool(ptr + 8 + 36 + 14, false);
        A.store.Bool(ptr + 8 + 36 + 6, false);
        A.store.Bool(ptr + 8 + 36 + 15, false);
        A.store.Bool(ptr + 8 + 36 + 7, false);

        A.store.Bool(ptr + 8 + 53 + 4, false);
        A.store.Bool(ptr + 8 + 53 + 2, false);
        A.store.Bool(ptr + 8 + 53 + 0, false);
        A.store.Bool(ptr + 8 + 53 + 3, false);
        A.store.Bool(ptr + 8 + 53 + 1, false);

        A.store.Bool(ptr + 8 + 58 + 4, false);
        A.store.Bool(ptr + 8 + 58 + 2, false);
        A.store.Bool(ptr + 8 + 58 + 0, false);
        A.store.Bool(ptr + 8 + 58 + 3, false);
        A.store.Bool(ptr + 8 + 58 + 1, false);

        A.store.Bool(ptr + 8 + 64 + 6, false);
        A.store.Ref(ptr + 8 + 64 + 0, undefined);
        A.store.Bool(ptr + 8 + 64 + 5, false);
        A.store.Bool(ptr + 8 + 64 + 4, false);
        A.store.Ref(ptr + 8 + 72, undefined);
        A.store.Ref(ptr + 8 + 76, undefined);

        A.store.Bool(ptr + 8 + 80 + 4, false);
        A.store.Bool(ptr + 8 + 80 + 2, false);
        A.store.Bool(ptr + 8 + 80 + 0, false);
        A.store.Bool(ptr + 8 + 80 + 3, false);
        A.store.Bool(ptr + 8 + 80 + 1, false);
        A.store.Ref(ptr + 8 + 88, undefined);
        A.store.Ref(ptr + 8 + 92, undefined);
        A.store.Enum(ptr + 8 + 96, -1);
        A.store.Ref(ptr + 8 + 100, undefined);
        A.store.Ref(ptr + 8 + 104, undefined);
        A.store.Ref(ptr + 8 + 108, undefined);
        A.store.Bool(ptr + 8 + 208, false);
        A.store.Bool(ptr + 8 + 112, false);
        A.store.Ref(ptr + 8 + 116, undefined);
        A.store.Bool(ptr + 8 + 209, false);
        A.store.Bool(ptr + 8 + 120, false);

        A.store.Bool(ptr + 8 + 124 + 9, false);
        A.store.Bool(ptr + 8 + 124 + 8, false);
        A.store.Bool(ptr + 8 + 124 + 0, false);
        A.store.Ref(ptr + 8 + 124 + 4, undefined);
        A.store.Ref(ptr + 8 + 136, undefined);

        A.store.Bool(ptr + 8 + 140 + 20, false);
        A.store.Ref(ptr + 8 + 140 + 0, undefined);

        A.store.Bool(ptr + 8 + 140 + 4 + 13, false);
        A.store.Bool(ptr + 8 + 140 + 4 + 12, false);
        A.store.Bool(ptr + 8 + 140 + 4 + 0, false);
        A.store.Enum(ptr + 8 + 140 + 4 + 4, -1);
        A.store.Ref(ptr + 8 + 140 + 4 + 8, undefined);
        A.store.Bool(ptr + 8 + 140 + 19, false);
        A.store.Bool(ptr + 8 + 140 + 18, false);
        A.store.Ref(ptr + 8 + 164, undefined);
        A.store.Ref(ptr + 8 + 168, undefined);
        A.store.Ref(ptr + 8 + 172, undefined);
        A.store.Enum(ptr + 8 + 176, -1);
        A.store.Enum(ptr + 8 + 180, -1);
        A.store.Ref(ptr + 8 + 184, undefined);
        A.store.Bool(ptr + 8 + 210, false);
        A.store.Bool(ptr + 8 + 188, false);
        A.store.Ref(ptr + 8 + 192, undefined);
        A.store.Ref(ptr + 8 + 196, undefined);
        A.store.Ref(ptr + 8 + 200, undefined);
        A.store.Bool(ptr + 8 + 211, false);
        A.store.Bool(ptr + 8 + 204, false);
        A.store.Bool(ptr + 8 + 212, false);
        A.store.Bool(ptr + 8 + 205, false);
        A.store.Bool(ptr + 8 + 213, false);
        A.store.Bool(ptr + 8 + 206, false);
        A.store.Bool(ptr + 8 + 214, false);
        A.store.Bool(ptr + 8 + 207, false);
      } else {
        A.store.Bool(ptr + 224, true);
        A.store.Enum(
          ptr + 0,
          [
            "INSTALLED",
            "UNINSTALLED",
            "LOADED",
            "UNLOADED",
            "VIEW_REGISTERED",
            "VIEW_UNREGISTERED",
            "ERROR_ADDED",
            "ERRORS_REMOVED",
            "PREFS_CHANGED",
            "WARNINGS_CHANGED",
            "COMMAND_ADDED",
            "COMMAND_REMOVED",
            "PERMISSIONS_CHANGED",
            "SERVICE_WORKER_STARTED",
            "SERVICE_WORKER_STOPPED",
            "CONFIGURATION_CHANGED",
            "PINNED_ACTIONS_CHANGED",
          ].indexOf(x["event_type"] as string)
        );
        A.store.Ref(ptr + 4, x["item_id"]);

        if (typeof x["extensionInfo"] === "undefined") {
          A.store.Bool(ptr + 8 + 215, false);
          A.store.Ref(ptr + 8 + 0, undefined);

          A.store.Bool(ptr + 8 + 4 + 8, false);
          A.store.Ref(ptr + 8 + 4 + 0, undefined);
          A.store.Ref(ptr + 8 + 4 + 4, undefined);
          A.store.Ref(ptr + 8 + 16, undefined);

          A.store.Bool(ptr + 8 + 20 + 4, false);
          A.store.Ref(ptr + 8 + 20 + 0, undefined);
          A.store.Ref(ptr + 8 + 28, undefined);
          A.store.Ref(ptr + 8 + 32, undefined);

          A.store.Bool(ptr + 8 + 36 + 16, false);
          A.store.Bool(ptr + 8 + 36 + 8, false);
          A.store.Bool(ptr + 8 + 36 + 0, false);
          A.store.Bool(ptr + 8 + 36 + 9, false);
          A.store.Bool(ptr + 8 + 36 + 1, false);
          A.store.Bool(ptr + 8 + 36 + 10, false);
          A.store.Bool(ptr + 8 + 36 + 2, false);
          A.store.Bool(ptr + 8 + 36 + 11, false);
          A.store.Bool(ptr + 8 + 36 + 3, false);
          A.store.Bool(ptr + 8 + 36 + 12, false);
          A.store.Bool(ptr + 8 + 36 + 4, false);
          A.store.Bool(ptr + 8 + 36 + 13, false);
          A.store.Bool(ptr + 8 + 36 + 5, false);
          A.store.Bool(ptr + 8 + 36 + 14, false);
          A.store.Bool(ptr + 8 + 36 + 6, false);
          A.store.Bool(ptr + 8 + 36 + 15, false);
          A.store.Bool(ptr + 8 + 36 + 7, false);

          A.store.Bool(ptr + 8 + 53 + 4, false);
          A.store.Bool(ptr + 8 + 53 + 2, false);
          A.store.Bool(ptr + 8 + 53 + 0, false);
          A.store.Bool(ptr + 8 + 53 + 3, false);
          A.store.Bool(ptr + 8 + 53 + 1, false);

          A.store.Bool(ptr + 8 + 58 + 4, false);
          A.store.Bool(ptr + 8 + 58 + 2, false);
          A.store.Bool(ptr + 8 + 58 + 0, false);
          A.store.Bool(ptr + 8 + 58 + 3, false);
          A.store.Bool(ptr + 8 + 58 + 1, false);

          A.store.Bool(ptr + 8 + 64 + 6, false);
          A.store.Ref(ptr + 8 + 64 + 0, undefined);
          A.store.Bool(ptr + 8 + 64 + 5, false);
          A.store.Bool(ptr + 8 + 64 + 4, false);
          A.store.Ref(ptr + 8 + 72, undefined);
          A.store.Ref(ptr + 8 + 76, undefined);

          A.store.Bool(ptr + 8 + 80 + 4, false);
          A.store.Bool(ptr + 8 + 80 + 2, false);
          A.store.Bool(ptr + 8 + 80 + 0, false);
          A.store.Bool(ptr + 8 + 80 + 3, false);
          A.store.Bool(ptr + 8 + 80 + 1, false);
          A.store.Ref(ptr + 8 + 88, undefined);
          A.store.Ref(ptr + 8 + 92, undefined);
          A.store.Enum(ptr + 8 + 96, -1);
          A.store.Ref(ptr + 8 + 100, undefined);
          A.store.Ref(ptr + 8 + 104, undefined);
          A.store.Ref(ptr + 8 + 108, undefined);
          A.store.Bool(ptr + 8 + 208, false);
          A.store.Bool(ptr + 8 + 112, false);
          A.store.Ref(ptr + 8 + 116, undefined);
          A.store.Bool(ptr + 8 + 209, false);
          A.store.Bool(ptr + 8 + 120, false);

          A.store.Bool(ptr + 8 + 124 + 9, false);
          A.store.Bool(ptr + 8 + 124 + 8, false);
          A.store.Bool(ptr + 8 + 124 + 0, false);
          A.store.Ref(ptr + 8 + 124 + 4, undefined);
          A.store.Ref(ptr + 8 + 136, undefined);

          A.store.Bool(ptr + 8 + 140 + 20, false);
          A.store.Ref(ptr + 8 + 140 + 0, undefined);

          A.store.Bool(ptr + 8 + 140 + 4 + 13, false);
          A.store.Bool(ptr + 8 + 140 + 4 + 12, false);
          A.store.Bool(ptr + 8 + 140 + 4 + 0, false);
          A.store.Enum(ptr + 8 + 140 + 4 + 4, -1);
          A.store.Ref(ptr + 8 + 140 + 4 + 8, undefined);
          A.store.Bool(ptr + 8 + 140 + 19, false);
          A.store.Bool(ptr + 8 + 140 + 18, false);
          A.store.Ref(ptr + 8 + 164, undefined);
          A.store.Ref(ptr + 8 + 168, undefined);
          A.store.Ref(ptr + 8 + 172, undefined);
          A.store.Enum(ptr + 8 + 176, -1);
          A.store.Enum(ptr + 8 + 180, -1);
          A.store.Ref(ptr + 8 + 184, undefined);
          A.store.Bool(ptr + 8 + 210, false);
          A.store.Bool(ptr + 8 + 188, false);
          A.store.Ref(ptr + 8 + 192, undefined);
          A.store.Ref(ptr + 8 + 196, undefined);
          A.store.Ref(ptr + 8 + 200, undefined);
          A.store.Bool(ptr + 8 + 211, false);
          A.store.Bool(ptr + 8 + 204, false);
          A.store.Bool(ptr + 8 + 212, false);
          A.store.Bool(ptr + 8 + 205, false);
          A.store.Bool(ptr + 8 + 213, false);
          A.store.Bool(ptr + 8 + 206, false);
          A.store.Bool(ptr + 8 + 214, false);
          A.store.Bool(ptr + 8 + 207, false);
        } else {
          A.store.Bool(ptr + 8 + 215, true);
          A.store.Ref(ptr + 8 + 0, x["extensionInfo"]["blacklistText"]);

          if (typeof x["extensionInfo"]["safetyCheckText"] === "undefined") {
            A.store.Bool(ptr + 8 + 4 + 8, false);
            A.store.Ref(ptr + 8 + 4 + 0, undefined);
            A.store.Ref(ptr + 8 + 4 + 4, undefined);
          } else {
            A.store.Bool(ptr + 8 + 4 + 8, true);
            A.store.Ref(ptr + 8 + 4 + 0, x["extensionInfo"]["safetyCheckText"]["panelString"]);
            A.store.Ref(ptr + 8 + 4 + 4, x["extensionInfo"]["safetyCheckText"]["detailString"]);
          }
          A.store.Ref(ptr + 8 + 16, x["extensionInfo"]["commands"]);

          if (typeof x["extensionInfo"]["controlledInfo"] === "undefined") {
            A.store.Bool(ptr + 8 + 20 + 4, false);
            A.store.Ref(ptr + 8 + 20 + 0, undefined);
          } else {
            A.store.Bool(ptr + 8 + 20 + 4, true);
            A.store.Ref(ptr + 8 + 20 + 0, x["extensionInfo"]["controlledInfo"]["text"]);
          }
          A.store.Ref(ptr + 8 + 28, x["extensionInfo"]["dependentExtensions"]);
          A.store.Ref(ptr + 8 + 32, x["extensionInfo"]["description"]);

          if (typeof x["extensionInfo"]["disableReasons"] === "undefined") {
            A.store.Bool(ptr + 8 + 36 + 16, false);
            A.store.Bool(ptr + 8 + 36 + 8, false);
            A.store.Bool(ptr + 8 + 36 + 0, false);
            A.store.Bool(ptr + 8 + 36 + 9, false);
            A.store.Bool(ptr + 8 + 36 + 1, false);
            A.store.Bool(ptr + 8 + 36 + 10, false);
            A.store.Bool(ptr + 8 + 36 + 2, false);
            A.store.Bool(ptr + 8 + 36 + 11, false);
            A.store.Bool(ptr + 8 + 36 + 3, false);
            A.store.Bool(ptr + 8 + 36 + 12, false);
            A.store.Bool(ptr + 8 + 36 + 4, false);
            A.store.Bool(ptr + 8 + 36 + 13, false);
            A.store.Bool(ptr + 8 + 36 + 5, false);
            A.store.Bool(ptr + 8 + 36 + 14, false);
            A.store.Bool(ptr + 8 + 36 + 6, false);
            A.store.Bool(ptr + 8 + 36 + 15, false);
            A.store.Bool(ptr + 8 + 36 + 7, false);
          } else {
            A.store.Bool(ptr + 8 + 36 + 16, true);
            A.store.Bool(ptr + 8 + 36 + 8, "suspiciousInstall" in x["extensionInfo"]["disableReasons"] ? true : false);
            A.store.Bool(ptr + 8 + 36 + 0, x["extensionInfo"]["disableReasons"]["suspiciousInstall"] ? true : false);
            A.store.Bool(ptr + 8 + 36 + 9, "corruptInstall" in x["extensionInfo"]["disableReasons"] ? true : false);
            A.store.Bool(ptr + 8 + 36 + 1, x["extensionInfo"]["disableReasons"]["corruptInstall"] ? true : false);
            A.store.Bool(ptr + 8 + 36 + 10, "updateRequired" in x["extensionInfo"]["disableReasons"] ? true : false);
            A.store.Bool(ptr + 8 + 36 + 2, x["extensionInfo"]["disableReasons"]["updateRequired"] ? true : false);
            A.store.Bool(
              ptr + 8 + 36 + 11,
              "publishedInStoreRequired" in x["extensionInfo"]["disableReasons"] ? true : false
            );
            A.store.Bool(
              ptr + 8 + 36 + 3,
              x["extensionInfo"]["disableReasons"]["publishedInStoreRequired"] ? true : false
            );
            A.store.Bool(ptr + 8 + 36 + 12, "blockedByPolicy" in x["extensionInfo"]["disableReasons"] ? true : false);
            A.store.Bool(ptr + 8 + 36 + 4, x["extensionInfo"]["disableReasons"]["blockedByPolicy"] ? true : false);
            A.store.Bool(ptr + 8 + 36 + 13, "reloading" in x["extensionInfo"]["disableReasons"] ? true : false);
            A.store.Bool(ptr + 8 + 36 + 5, x["extensionInfo"]["disableReasons"]["reloading"] ? true : false);
            A.store.Bool(
              ptr + 8 + 36 + 14,
              "custodianApprovalRequired" in x["extensionInfo"]["disableReasons"] ? true : false
            );
            A.store.Bool(
              ptr + 8 + 36 + 6,
              x["extensionInfo"]["disableReasons"]["custodianApprovalRequired"] ? true : false
            );
            A.store.Bool(
              ptr + 8 + 36 + 15,
              "parentDisabledPermissions" in x["extensionInfo"]["disableReasons"] ? true : false
            );
            A.store.Bool(
              ptr + 8 + 36 + 7,
              x["extensionInfo"]["disableReasons"]["parentDisabledPermissions"] ? true : false
            );
          }

          if (typeof x["extensionInfo"]["errorCollection"] === "undefined") {
            A.store.Bool(ptr + 8 + 53 + 4, false);
            A.store.Bool(ptr + 8 + 53 + 2, false);
            A.store.Bool(ptr + 8 + 53 + 0, false);
            A.store.Bool(ptr + 8 + 53 + 3, false);
            A.store.Bool(ptr + 8 + 53 + 1, false);
          } else {
            A.store.Bool(ptr + 8 + 53 + 4, true);
            A.store.Bool(ptr + 8 + 53 + 2, "isEnabled" in x["extensionInfo"]["errorCollection"] ? true : false);
            A.store.Bool(ptr + 8 + 53 + 0, x["extensionInfo"]["errorCollection"]["isEnabled"] ? true : false);
            A.store.Bool(ptr + 8 + 53 + 3, "isActive" in x["extensionInfo"]["errorCollection"] ? true : false);
            A.store.Bool(ptr + 8 + 53 + 1, x["extensionInfo"]["errorCollection"]["isActive"] ? true : false);
          }

          if (typeof x["extensionInfo"]["fileAccess"] === "undefined") {
            A.store.Bool(ptr + 8 + 58 + 4, false);
            A.store.Bool(ptr + 8 + 58 + 2, false);
            A.store.Bool(ptr + 8 + 58 + 0, false);
            A.store.Bool(ptr + 8 + 58 + 3, false);
            A.store.Bool(ptr + 8 + 58 + 1, false);
          } else {
            A.store.Bool(ptr + 8 + 58 + 4, true);
            A.store.Bool(ptr + 8 + 58 + 2, "isEnabled" in x["extensionInfo"]["fileAccess"] ? true : false);
            A.store.Bool(ptr + 8 + 58 + 0, x["extensionInfo"]["fileAccess"]["isEnabled"] ? true : false);
            A.store.Bool(ptr + 8 + 58 + 3, "isActive" in x["extensionInfo"]["fileAccess"] ? true : false);
            A.store.Bool(ptr + 8 + 58 + 1, x["extensionInfo"]["fileAccess"]["isActive"] ? true : false);
          }

          if (typeof x["extensionInfo"]["homePage"] === "undefined") {
            A.store.Bool(ptr + 8 + 64 + 6, false);
            A.store.Ref(ptr + 8 + 64 + 0, undefined);
            A.store.Bool(ptr + 8 + 64 + 5, false);
            A.store.Bool(ptr + 8 + 64 + 4, false);
          } else {
            A.store.Bool(ptr + 8 + 64 + 6, true);
            A.store.Ref(ptr + 8 + 64 + 0, x["extensionInfo"]["homePage"]["url"]);
            A.store.Bool(ptr + 8 + 64 + 5, "specified" in x["extensionInfo"]["homePage"] ? true : false);
            A.store.Bool(ptr + 8 + 64 + 4, x["extensionInfo"]["homePage"]["specified"] ? true : false);
          }
          A.store.Ref(ptr + 8 + 72, x["extensionInfo"]["iconUrl"]);
          A.store.Ref(ptr + 8 + 76, x["extensionInfo"]["id"]);

          if (typeof x["extensionInfo"]["incognitoAccess"] === "undefined") {
            A.store.Bool(ptr + 8 + 80 + 4, false);
            A.store.Bool(ptr + 8 + 80 + 2, false);
            A.store.Bool(ptr + 8 + 80 + 0, false);
            A.store.Bool(ptr + 8 + 80 + 3, false);
            A.store.Bool(ptr + 8 + 80 + 1, false);
          } else {
            A.store.Bool(ptr + 8 + 80 + 4, true);
            A.store.Bool(ptr + 8 + 80 + 2, "isEnabled" in x["extensionInfo"]["incognitoAccess"] ? true : false);
            A.store.Bool(ptr + 8 + 80 + 0, x["extensionInfo"]["incognitoAccess"]["isEnabled"] ? true : false);
            A.store.Bool(ptr + 8 + 80 + 3, "isActive" in x["extensionInfo"]["incognitoAccess"] ? true : false);
            A.store.Bool(ptr + 8 + 80 + 1, x["extensionInfo"]["incognitoAccess"]["isActive"] ? true : false);
          }
          A.store.Ref(ptr + 8 + 88, x["extensionInfo"]["installWarnings"]);
          A.store.Ref(ptr + 8 + 92, x["extensionInfo"]["launchUrl"]);
          A.store.Enum(
            ptr + 8 + 96,
            ["FROM_STORE", "UNPACKED", "THIRD_PARTY", "INSTALLED_BY_DEFAULT", "UNKNOWN"].indexOf(
              x["extensionInfo"]["location"] as string
            )
          );
          A.store.Ref(ptr + 8 + 100, x["extensionInfo"]["locationText"]);
          A.store.Ref(ptr + 8 + 104, x["extensionInfo"]["manifestErrors"]);
          A.store.Ref(ptr + 8 + 108, x["extensionInfo"]["manifestHomePageUrl"]);
          A.store.Bool(ptr + 8 + 208, "mustRemainInstalled" in x["extensionInfo"] ? true : false);
          A.store.Bool(ptr + 8 + 112, x["extensionInfo"]["mustRemainInstalled"] ? true : false);
          A.store.Ref(ptr + 8 + 116, x["extensionInfo"]["name"]);
          A.store.Bool(ptr + 8 + 209, "offlineEnabled" in x["extensionInfo"] ? true : false);
          A.store.Bool(ptr + 8 + 120, x["extensionInfo"]["offlineEnabled"] ? true : false);

          if (typeof x["extensionInfo"]["optionsPage"] === "undefined") {
            A.store.Bool(ptr + 8 + 124 + 9, false);
            A.store.Bool(ptr + 8 + 124 + 8, false);
            A.store.Bool(ptr + 8 + 124 + 0, false);
            A.store.Ref(ptr + 8 + 124 + 4, undefined);
          } else {
            A.store.Bool(ptr + 8 + 124 + 9, true);
            A.store.Bool(ptr + 8 + 124 + 8, "openInTab" in x["extensionInfo"]["optionsPage"] ? true : false);
            A.store.Bool(ptr + 8 + 124 + 0, x["extensionInfo"]["optionsPage"]["openInTab"] ? true : false);
            A.store.Ref(ptr + 8 + 124 + 4, x["extensionInfo"]["optionsPage"]["url"]);
          }
          A.store.Ref(ptr + 8 + 136, x["extensionInfo"]["path"]);

          if (typeof x["extensionInfo"]["permissions"] === "undefined") {
            A.store.Bool(ptr + 8 + 140 + 20, false);
            A.store.Ref(ptr + 8 + 140 + 0, undefined);

            A.store.Bool(ptr + 8 + 140 + 4 + 13, false);
            A.store.Bool(ptr + 8 + 140 + 4 + 12, false);
            A.store.Bool(ptr + 8 + 140 + 4 + 0, false);
            A.store.Enum(ptr + 8 + 140 + 4 + 4, -1);
            A.store.Ref(ptr + 8 + 140 + 4 + 8, undefined);
            A.store.Bool(ptr + 8 + 140 + 19, false);
            A.store.Bool(ptr + 8 + 140 + 18, false);
          } else {
            A.store.Bool(ptr + 8 + 140 + 20, true);
            A.store.Ref(ptr + 8 + 140 + 0, x["extensionInfo"]["permissions"]["simplePermissions"]);

            if (typeof x["extensionInfo"]["permissions"]["runtimeHostPermissions"] === "undefined") {
              A.store.Bool(ptr + 8 + 140 + 4 + 13, false);
              A.store.Bool(ptr + 8 + 140 + 4 + 12, false);
              A.store.Bool(ptr + 8 + 140 + 4 + 0, false);
              A.store.Enum(ptr + 8 + 140 + 4 + 4, -1);
              A.store.Ref(ptr + 8 + 140 + 4 + 8, undefined);
            } else {
              A.store.Bool(ptr + 8 + 140 + 4 + 13, true);
              A.store.Bool(
                ptr + 8 + 140 + 4 + 12,
                "hasAllHosts" in x["extensionInfo"]["permissions"]["runtimeHostPermissions"] ? true : false
              );
              A.store.Bool(
                ptr + 8 + 140 + 4 + 0,
                x["extensionInfo"]["permissions"]["runtimeHostPermissions"]["hasAllHosts"] ? true : false
              );
              A.store.Enum(
                ptr + 8 + 140 + 4 + 4,
                ["ON_CLICK", "ON_SPECIFIC_SITES", "ON_ALL_SITES"].indexOf(
                  x["extensionInfo"]["permissions"]["runtimeHostPermissions"]["hostAccess"] as string
                )
              );
              A.store.Ref(ptr + 8 + 140 + 4 + 8, x["extensionInfo"]["permissions"]["runtimeHostPermissions"]["hosts"]);
            }
            A.store.Bool(ptr + 8 + 140 + 19, "canAccessSiteData" in x["extensionInfo"]["permissions"] ? true : false);
            A.store.Bool(ptr + 8 + 140 + 18, x["extensionInfo"]["permissions"]["canAccessSiteData"] ? true : false);
          }
          A.store.Ref(ptr + 8 + 164, x["extensionInfo"]["prettifiedPath"]);
          A.store.Ref(ptr + 8 + 168, x["extensionInfo"]["runtimeErrors"]);
          A.store.Ref(ptr + 8 + 172, x["extensionInfo"]["runtimeWarnings"]);
          A.store.Enum(
            ptr + 8 + 176,
            ["ENABLED", "DISABLED", "TERMINATED", "BLACKLISTED"].indexOf(x["extensionInfo"]["state"] as string)
          );
          A.store.Enum(
            ptr + 8 + 180,
            ["HOSTED_APP", "PLATFORM_APP", "LEGACY_PACKAGED_APP", "EXTENSION", "THEME", "SHARED_MODULE"].indexOf(
              x["extensionInfo"]["type"] as string
            )
          );
          A.store.Ref(ptr + 8 + 184, x["extensionInfo"]["updateUrl"]);
          A.store.Bool(ptr + 8 + 210, "userMayModify" in x["extensionInfo"] ? true : false);
          A.store.Bool(ptr + 8 + 188, x["extensionInfo"]["userMayModify"] ? true : false);
          A.store.Ref(ptr + 8 + 192, x["extensionInfo"]["version"]);
          A.store.Ref(ptr + 8 + 196, x["extensionInfo"]["views"]);
          A.store.Ref(ptr + 8 + 200, x["extensionInfo"]["webStoreUrl"]);
          A.store.Bool(ptr + 8 + 211, "showSafeBrowsingAllowlistWarning" in x["extensionInfo"] ? true : false);
          A.store.Bool(ptr + 8 + 204, x["extensionInfo"]["showSafeBrowsingAllowlistWarning"] ? true : false);
          A.store.Bool(ptr + 8 + 212, "showAccessRequestsInToolbar" in x["extensionInfo"] ? true : false);
          A.store.Bool(ptr + 8 + 205, x["extensionInfo"]["showAccessRequestsInToolbar"] ? true : false);
          A.store.Bool(ptr + 8 + 213, "acknowledgeSafetyCheckWarning" in x["extensionInfo"] ? true : false);
          A.store.Bool(ptr + 8 + 206, x["extensionInfo"]["acknowledgeSafetyCheckWarning"] ? true : false);
          A.store.Bool(ptr + 8 + 214, "pinnedToToolbar" in x["extensionInfo"] ? true : false);
          A.store.Bool(ptr + 8 + 207, x["extensionInfo"]["pinnedToToolbar"] ? true : false);
        }
      }
    },
    "load_EventData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["event_type"] = A.load.Enum(ptr + 0, [
        "INSTALLED",
        "UNINSTALLED",
        "LOADED",
        "UNLOADED",
        "VIEW_REGISTERED",
        "VIEW_UNREGISTERED",
        "ERROR_ADDED",
        "ERRORS_REMOVED",
        "PREFS_CHANGED",
        "WARNINGS_CHANGED",
        "COMMAND_ADDED",
        "COMMAND_REMOVED",
        "PERMISSIONS_CHANGED",
        "SERVICE_WORKER_STARTED",
        "SERVICE_WORKER_STOPPED",
        "CONFIGURATION_CHANGED",
        "PINNED_ACTIONS_CHANGED",
      ]);
      x["item_id"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 8 + 215)) {
        x["extensionInfo"] = {};
        x["extensionInfo"]["blacklistText"] = A.load.Ref(ptr + 8 + 0, undefined);
        if (A.load.Bool(ptr + 8 + 4 + 8)) {
          x["extensionInfo"]["safetyCheckText"] = {};
          x["extensionInfo"]["safetyCheckText"]["panelString"] = A.load.Ref(ptr + 8 + 4 + 0, undefined);
          x["extensionInfo"]["safetyCheckText"]["detailString"] = A.load.Ref(ptr + 8 + 4 + 4, undefined);
        } else {
          delete x["extensionInfo"]["safetyCheckText"];
        }
        x["extensionInfo"]["commands"] = A.load.Ref(ptr + 8 + 16, undefined);
        if (A.load.Bool(ptr + 8 + 20 + 4)) {
          x["extensionInfo"]["controlledInfo"] = {};
          x["extensionInfo"]["controlledInfo"]["text"] = A.load.Ref(ptr + 8 + 20 + 0, undefined);
        } else {
          delete x["extensionInfo"]["controlledInfo"];
        }
        x["extensionInfo"]["dependentExtensions"] = A.load.Ref(ptr + 8 + 28, undefined);
        x["extensionInfo"]["description"] = A.load.Ref(ptr + 8 + 32, undefined);
        if (A.load.Bool(ptr + 8 + 36 + 16)) {
          x["extensionInfo"]["disableReasons"] = {};
          if (A.load.Bool(ptr + 8 + 36 + 8)) {
            x["extensionInfo"]["disableReasons"]["suspiciousInstall"] = A.load.Bool(ptr + 8 + 36 + 0);
          } else {
            delete x["extensionInfo"]["disableReasons"]["suspiciousInstall"];
          }
          if (A.load.Bool(ptr + 8 + 36 + 9)) {
            x["extensionInfo"]["disableReasons"]["corruptInstall"] = A.load.Bool(ptr + 8 + 36 + 1);
          } else {
            delete x["extensionInfo"]["disableReasons"]["corruptInstall"];
          }
          if (A.load.Bool(ptr + 8 + 36 + 10)) {
            x["extensionInfo"]["disableReasons"]["updateRequired"] = A.load.Bool(ptr + 8 + 36 + 2);
          } else {
            delete x["extensionInfo"]["disableReasons"]["updateRequired"];
          }
          if (A.load.Bool(ptr + 8 + 36 + 11)) {
            x["extensionInfo"]["disableReasons"]["publishedInStoreRequired"] = A.load.Bool(ptr + 8 + 36 + 3);
          } else {
            delete x["extensionInfo"]["disableReasons"]["publishedInStoreRequired"];
          }
          if (A.load.Bool(ptr + 8 + 36 + 12)) {
            x["extensionInfo"]["disableReasons"]["blockedByPolicy"] = A.load.Bool(ptr + 8 + 36 + 4);
          } else {
            delete x["extensionInfo"]["disableReasons"]["blockedByPolicy"];
          }
          if (A.load.Bool(ptr + 8 + 36 + 13)) {
            x["extensionInfo"]["disableReasons"]["reloading"] = A.load.Bool(ptr + 8 + 36 + 5);
          } else {
            delete x["extensionInfo"]["disableReasons"]["reloading"];
          }
          if (A.load.Bool(ptr + 8 + 36 + 14)) {
            x["extensionInfo"]["disableReasons"]["custodianApprovalRequired"] = A.load.Bool(ptr + 8 + 36 + 6);
          } else {
            delete x["extensionInfo"]["disableReasons"]["custodianApprovalRequired"];
          }
          if (A.load.Bool(ptr + 8 + 36 + 15)) {
            x["extensionInfo"]["disableReasons"]["parentDisabledPermissions"] = A.load.Bool(ptr + 8 + 36 + 7);
          } else {
            delete x["extensionInfo"]["disableReasons"]["parentDisabledPermissions"];
          }
        } else {
          delete x["extensionInfo"]["disableReasons"];
        }
        if (A.load.Bool(ptr + 8 + 53 + 4)) {
          x["extensionInfo"]["errorCollection"] = {};
          if (A.load.Bool(ptr + 8 + 53 + 2)) {
            x["extensionInfo"]["errorCollection"]["isEnabled"] = A.load.Bool(ptr + 8 + 53 + 0);
          } else {
            delete x["extensionInfo"]["errorCollection"]["isEnabled"];
          }
          if (A.load.Bool(ptr + 8 + 53 + 3)) {
            x["extensionInfo"]["errorCollection"]["isActive"] = A.load.Bool(ptr + 8 + 53 + 1);
          } else {
            delete x["extensionInfo"]["errorCollection"]["isActive"];
          }
        } else {
          delete x["extensionInfo"]["errorCollection"];
        }
        if (A.load.Bool(ptr + 8 + 58 + 4)) {
          x["extensionInfo"]["fileAccess"] = {};
          if (A.load.Bool(ptr + 8 + 58 + 2)) {
            x["extensionInfo"]["fileAccess"]["isEnabled"] = A.load.Bool(ptr + 8 + 58 + 0);
          } else {
            delete x["extensionInfo"]["fileAccess"]["isEnabled"];
          }
          if (A.load.Bool(ptr + 8 + 58 + 3)) {
            x["extensionInfo"]["fileAccess"]["isActive"] = A.load.Bool(ptr + 8 + 58 + 1);
          } else {
            delete x["extensionInfo"]["fileAccess"]["isActive"];
          }
        } else {
          delete x["extensionInfo"]["fileAccess"];
        }
        if (A.load.Bool(ptr + 8 + 64 + 6)) {
          x["extensionInfo"]["homePage"] = {};
          x["extensionInfo"]["homePage"]["url"] = A.load.Ref(ptr + 8 + 64 + 0, undefined);
          if (A.load.Bool(ptr + 8 + 64 + 5)) {
            x["extensionInfo"]["homePage"]["specified"] = A.load.Bool(ptr + 8 + 64 + 4);
          } else {
            delete x["extensionInfo"]["homePage"]["specified"];
          }
        } else {
          delete x["extensionInfo"]["homePage"];
        }
        x["extensionInfo"]["iconUrl"] = A.load.Ref(ptr + 8 + 72, undefined);
        x["extensionInfo"]["id"] = A.load.Ref(ptr + 8 + 76, undefined);
        if (A.load.Bool(ptr + 8 + 80 + 4)) {
          x["extensionInfo"]["incognitoAccess"] = {};
          if (A.load.Bool(ptr + 8 + 80 + 2)) {
            x["extensionInfo"]["incognitoAccess"]["isEnabled"] = A.load.Bool(ptr + 8 + 80 + 0);
          } else {
            delete x["extensionInfo"]["incognitoAccess"]["isEnabled"];
          }
          if (A.load.Bool(ptr + 8 + 80 + 3)) {
            x["extensionInfo"]["incognitoAccess"]["isActive"] = A.load.Bool(ptr + 8 + 80 + 1);
          } else {
            delete x["extensionInfo"]["incognitoAccess"]["isActive"];
          }
        } else {
          delete x["extensionInfo"]["incognitoAccess"];
        }
        x["extensionInfo"]["installWarnings"] = A.load.Ref(ptr + 8 + 88, undefined);
        x["extensionInfo"]["launchUrl"] = A.load.Ref(ptr + 8 + 92, undefined);
        x["extensionInfo"]["location"] = A.load.Enum(ptr + 8 + 96, [
          "FROM_STORE",
          "UNPACKED",
          "THIRD_PARTY",
          "INSTALLED_BY_DEFAULT",
          "UNKNOWN",
        ]);
        x["extensionInfo"]["locationText"] = A.load.Ref(ptr + 8 + 100, undefined);
        x["extensionInfo"]["manifestErrors"] = A.load.Ref(ptr + 8 + 104, undefined);
        x["extensionInfo"]["manifestHomePageUrl"] = A.load.Ref(ptr + 8 + 108, undefined);
        if (A.load.Bool(ptr + 8 + 208)) {
          x["extensionInfo"]["mustRemainInstalled"] = A.load.Bool(ptr + 8 + 112);
        } else {
          delete x["extensionInfo"]["mustRemainInstalled"];
        }
        x["extensionInfo"]["name"] = A.load.Ref(ptr + 8 + 116, undefined);
        if (A.load.Bool(ptr + 8 + 209)) {
          x["extensionInfo"]["offlineEnabled"] = A.load.Bool(ptr + 8 + 120);
        } else {
          delete x["extensionInfo"]["offlineEnabled"];
        }
        if (A.load.Bool(ptr + 8 + 124 + 9)) {
          x["extensionInfo"]["optionsPage"] = {};
          if (A.load.Bool(ptr + 8 + 124 + 8)) {
            x["extensionInfo"]["optionsPage"]["openInTab"] = A.load.Bool(ptr + 8 + 124 + 0);
          } else {
            delete x["extensionInfo"]["optionsPage"]["openInTab"];
          }
          x["extensionInfo"]["optionsPage"]["url"] = A.load.Ref(ptr + 8 + 124 + 4, undefined);
        } else {
          delete x["extensionInfo"]["optionsPage"];
        }
        x["extensionInfo"]["path"] = A.load.Ref(ptr + 8 + 136, undefined);
        if (A.load.Bool(ptr + 8 + 140 + 20)) {
          x["extensionInfo"]["permissions"] = {};
          x["extensionInfo"]["permissions"]["simplePermissions"] = A.load.Ref(ptr + 8 + 140 + 0, undefined);
          if (A.load.Bool(ptr + 8 + 140 + 4 + 13)) {
            x["extensionInfo"]["permissions"]["runtimeHostPermissions"] = {};
            if (A.load.Bool(ptr + 8 + 140 + 4 + 12)) {
              x["extensionInfo"]["permissions"]["runtimeHostPermissions"]["hasAllHosts"] = A.load.Bool(
                ptr + 8 + 140 + 4 + 0
              );
            } else {
              delete x["extensionInfo"]["permissions"]["runtimeHostPermissions"]["hasAllHosts"];
            }
            x["extensionInfo"]["permissions"]["runtimeHostPermissions"]["hostAccess"] = A.load.Enum(
              ptr + 8 + 140 + 4 + 4,
              ["ON_CLICK", "ON_SPECIFIC_SITES", "ON_ALL_SITES"]
            );
            x["extensionInfo"]["permissions"]["runtimeHostPermissions"]["hosts"] = A.load.Ref(
              ptr + 8 + 140 + 4 + 8,
              undefined
            );
          } else {
            delete x["extensionInfo"]["permissions"]["runtimeHostPermissions"];
          }
          if (A.load.Bool(ptr + 8 + 140 + 19)) {
            x["extensionInfo"]["permissions"]["canAccessSiteData"] = A.load.Bool(ptr + 8 + 140 + 18);
          } else {
            delete x["extensionInfo"]["permissions"]["canAccessSiteData"];
          }
        } else {
          delete x["extensionInfo"]["permissions"];
        }
        x["extensionInfo"]["prettifiedPath"] = A.load.Ref(ptr + 8 + 164, undefined);
        x["extensionInfo"]["runtimeErrors"] = A.load.Ref(ptr + 8 + 168, undefined);
        x["extensionInfo"]["runtimeWarnings"] = A.load.Ref(ptr + 8 + 172, undefined);
        x["extensionInfo"]["state"] = A.load.Enum(ptr + 8 + 176, ["ENABLED", "DISABLED", "TERMINATED", "BLACKLISTED"]);
        x["extensionInfo"]["type"] = A.load.Enum(ptr + 8 + 180, [
          "HOSTED_APP",
          "PLATFORM_APP",
          "LEGACY_PACKAGED_APP",
          "EXTENSION",
          "THEME",
          "SHARED_MODULE",
        ]);
        x["extensionInfo"]["updateUrl"] = A.load.Ref(ptr + 8 + 184, undefined);
        if (A.load.Bool(ptr + 8 + 210)) {
          x["extensionInfo"]["userMayModify"] = A.load.Bool(ptr + 8 + 188);
        } else {
          delete x["extensionInfo"]["userMayModify"];
        }
        x["extensionInfo"]["version"] = A.load.Ref(ptr + 8 + 192, undefined);
        x["extensionInfo"]["views"] = A.load.Ref(ptr + 8 + 196, undefined);
        x["extensionInfo"]["webStoreUrl"] = A.load.Ref(ptr + 8 + 200, undefined);
        if (A.load.Bool(ptr + 8 + 211)) {
          x["extensionInfo"]["showSafeBrowsingAllowlistWarning"] = A.load.Bool(ptr + 8 + 204);
        } else {
          delete x["extensionInfo"]["showSafeBrowsingAllowlistWarning"];
        }
        if (A.load.Bool(ptr + 8 + 212)) {
          x["extensionInfo"]["showAccessRequestsInToolbar"] = A.load.Bool(ptr + 8 + 205);
        } else {
          delete x["extensionInfo"]["showAccessRequestsInToolbar"];
        }
        if (A.load.Bool(ptr + 8 + 213)) {
          x["extensionInfo"]["acknowledgeSafetyCheckWarning"] = A.load.Bool(ptr + 8 + 206);
        } else {
          delete x["extensionInfo"]["acknowledgeSafetyCheckWarning"];
        }
        if (A.load.Bool(ptr + 8 + 214)) {
          x["extensionInfo"]["pinnedToToolbar"] = A.load.Bool(ptr + 8 + 207);
        } else {
          delete x["extensionInfo"]["pinnedToToolbar"];
        }
      } else {
        delete x["extensionInfo"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ExtensionCommandUpdate": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["extensionId"]);
        A.store.Ref(ptr + 4, x["commandName"]);
        A.store.Enum(ptr + 8, ["GLOBAL", "CHROME"].indexOf(x["scope"] as string));
        A.store.Ref(ptr + 12, x["keybinding"]);
      }
    },
    "load_ExtensionCommandUpdate": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["extensionId"] = A.load.Ref(ptr + 0, undefined);
      x["commandName"] = A.load.Ref(ptr + 4, undefined);
      x["scope"] = A.load.Enum(ptr + 8, ["GLOBAL", "CHROME"]);
      x["keybinding"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ExtensionConfigurationUpdate": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 6, false);
        A.store.Enum(ptr + 8, -1);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 14, false);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Ref(ptr + 0, x["extensionId"]);
        A.store.Bool(ptr + 15, "fileAccess" in x ? true : false);
        A.store.Bool(ptr + 4, x["fileAccess"] ? true : false);
        A.store.Bool(ptr + 16, "incognitoAccess" in x ? true : false);
        A.store.Bool(ptr + 5, x["incognitoAccess"] ? true : false);
        A.store.Bool(ptr + 17, "errorCollection" in x ? true : false);
        A.store.Bool(ptr + 6, x["errorCollection"] ? true : false);
        A.store.Enum(ptr + 8, ["ON_CLICK", "ON_SPECIFIC_SITES", "ON_ALL_SITES"].indexOf(x["hostAccess"] as string));
        A.store.Bool(ptr + 18, "showAccessRequestsInToolbar" in x ? true : false);
        A.store.Bool(ptr + 12, x["showAccessRequestsInToolbar"] ? true : false);
        A.store.Bool(ptr + 19, "acknowledgeSafetyCheckWarning" in x ? true : false);
        A.store.Bool(ptr + 13, x["acknowledgeSafetyCheckWarning"] ? true : false);
        A.store.Bool(ptr + 20, "pinnedToToolbar" in x ? true : false);
        A.store.Bool(ptr + 14, x["pinnedToToolbar"] ? true : false);
      }
    },
    "load_ExtensionConfigurationUpdate": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["extensionId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 15)) {
        x["fileAccess"] = A.load.Bool(ptr + 4);
      } else {
        delete x["fileAccess"];
      }
      if (A.load.Bool(ptr + 16)) {
        x["incognitoAccess"] = A.load.Bool(ptr + 5);
      } else {
        delete x["incognitoAccess"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["errorCollection"] = A.load.Bool(ptr + 6);
      } else {
        delete x["errorCollection"];
      }
      x["hostAccess"] = A.load.Enum(ptr + 8, ["ON_CLICK", "ON_SPECIFIC_SITES", "ON_ALL_SITES"]);
      if (A.load.Bool(ptr + 18)) {
        x["showAccessRequestsInToolbar"] = A.load.Bool(ptr + 12);
      } else {
        delete x["showAccessRequestsInToolbar"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["acknowledgeSafetyCheckWarning"] = A.load.Bool(ptr + 13);
      } else {
        delete x["acknowledgeSafetyCheckWarning"];
      }
      if (A.load.Bool(ptr + 20)) {
        x["pinnedToToolbar"] = A.load.Bool(ptr + 14);
      } else {
        delete x["pinnedToToolbar"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ExtensionSiteAccessUpdate": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Enum(ptr + 4, ["ON_CLICK", "ON_SPECIFIC_SITES", "ON_ALL_SITES"].indexOf(x["siteAccess"] as string));
      }
    },
    "load_ExtensionSiteAccessUpdate": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["siteAccess"] = A.load.Enum(ptr + 4, ["ON_CLICK", "ON_SPECIFIC_SITES", "ON_ALL_SITES"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_FileType": (ref: heap.Ref<string>): number => {
      const idx = ["LOAD", "PEM"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_GetExtensionsInfoOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 1, false);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Bool(ptr + 2, "includeDisabled" in x ? true : false);
        A.store.Bool(ptr + 0, x["includeDisabled"] ? true : false);
        A.store.Bool(ptr + 3, "includeTerminated" in x ? true : false);
        A.store.Bool(ptr + 1, x["includeTerminated"] ? true : false);
      }
    },
    "load_GetExtensionsInfoOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 2)) {
        x["includeDisabled"] = A.load.Bool(ptr + 0);
      } else {
        delete x["includeDisabled"];
      }
      if (A.load.Bool(ptr + 3)) {
        x["includeTerminated"] = A.load.Bool(ptr + 1);
      } else {
        delete x["includeTerminated"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MatchingExtensionInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Enum(ptr + 4, ["ON_CLICK", "ON_SPECIFIC_SITES", "ON_ALL_SITES"].indexOf(x["siteAccess"] as string));
        A.store.Bool(ptr + 9, "canRequestAllSites" in x ? true : false);
        A.store.Bool(ptr + 8, x["canRequestAllSites"] ? true : false);
      }
    },
    "load_MatchingExtensionInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["siteAccess"] = A.load.Enum(ptr + 4, ["ON_CLICK", "ON_SPECIFIC_SITES", "ON_ALL_SITES"]);
      if (A.load.Bool(ptr + 9)) {
        x["canRequestAllSites"] = A.load.Bool(ptr + 8);
      } else {
        delete x["canRequestAllSites"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ProjectInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["name"]);
      }
    },
    "load_ProjectInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_InspectOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Ref(ptr + 0, x["extension_id"]);
        A.store.Ref(ptr + 4, x["render_process_id"]);
        A.store.Ref(ptr + 8, x["render_view_id"]);
        A.store.Bool(ptr + 13, "incognito" in x ? true : false);
        A.store.Bool(ptr + 12, x["incognito"] ? true : false);
      }
    },
    "load_InspectOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["extension_id"] = A.load.Ref(ptr + 0, undefined);
      x["render_process_id"] = A.load.Ref(ptr + 4, undefined);
      x["render_view_id"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 13)) {
        x["incognito"] = A.load.Bool(ptr + 12);
      } else {
        delete x["incognito"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_InstallWarning": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["message"]);
      }
    },
    "load_InstallWarning": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["message"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ItemType": (ref: heap.Ref<string>): number => {
      const idx = ["hosted_app", "packaged_app", "legacy_packaged_app", "extension", "theme"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ItemInspectView": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 14, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 15, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 13, false);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Ref(ptr + 0, x["path"]);
        A.store.Bool(ptr + 14, "render_process_id" in x ? true : false);
        A.store.Int32(ptr + 4, x["render_process_id"] === undefined ? 0 : (x["render_process_id"] as number));
        A.store.Bool(ptr + 15, "render_view_id" in x ? true : false);
        A.store.Int32(ptr + 8, x["render_view_id"] === undefined ? 0 : (x["render_view_id"] as number));
        A.store.Bool(ptr + 16, "incognito" in x ? true : false);
        A.store.Bool(ptr + 12, x["incognito"] ? true : false);
        A.store.Bool(ptr + 17, "generatedBackgroundPage" in x ? true : false);
        A.store.Bool(ptr + 13, x["generatedBackgroundPage"] ? true : false);
      }
    },
    "load_ItemInspectView": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["path"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 14)) {
        x["render_process_id"] = A.load.Int32(ptr + 4);
      } else {
        delete x["render_process_id"];
      }
      if (A.load.Bool(ptr + 15)) {
        x["render_view_id"] = A.load.Int32(ptr + 8);
      } else {
        delete x["render_view_id"];
      }
      if (A.load.Bool(ptr + 16)) {
        x["incognito"] = A.load.Bool(ptr + 12);
      } else {
        delete x["incognito"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["generatedBackgroundPage"] = A.load.Bool(ptr + 13);
      } else {
        delete x["generatedBackgroundPage"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ItemInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 88, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 76, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 77, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 78, false);
        A.store.Bool(ptr + 18, false);
        A.store.Enum(ptr + 20, -1);
        A.store.Bool(ptr + 79, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 80, false);
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 81, false);
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 82, false);
        A.store.Bool(ptr + 27, false);
        A.store.Bool(ptr + 83, false);
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 84, false);
        A.store.Bool(ptr + 29, false);
        A.store.Bool(ptr + 85, false);
        A.store.Bool(ptr + 30, false);
        A.store.Bool(ptr + 86, false);
        A.store.Bool(ptr + 31, false);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Ref(ptr + 48, undefined);
        A.store.Ref(ptr + 52, undefined);
        A.store.Ref(ptr + 56, undefined);
        A.store.Ref(ptr + 60, undefined);
        A.store.Ref(ptr + 64, undefined);
        A.store.Bool(ptr + 87, false);
        A.store.Bool(ptr + 68, false);
        A.store.Ref(ptr + 72, undefined);
      } else {
        A.store.Bool(ptr + 88, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Ref(ptr + 8, x["version"]);
        A.store.Ref(ptr + 12, x["description"]);
        A.store.Bool(ptr + 76, "may_disable" in x ? true : false);
        A.store.Bool(ptr + 16, x["may_disable"] ? true : false);
        A.store.Bool(ptr + 77, "enabled" in x ? true : false);
        A.store.Bool(ptr + 17, x["enabled"] ? true : false);
        A.store.Bool(ptr + 78, "isApp" in x ? true : false);
        A.store.Bool(ptr + 18, x["isApp"] ? true : false);
        A.store.Enum(
          ptr + 20,
          ["hosted_app", "packaged_app", "legacy_packaged_app", "extension", "theme"].indexOf(x["type"] as string)
        );
        A.store.Bool(ptr + 79, "allow_activity" in x ? true : false);
        A.store.Bool(ptr + 24, x["allow_activity"] ? true : false);
        A.store.Bool(ptr + 80, "allow_file_access" in x ? true : false);
        A.store.Bool(ptr + 25, x["allow_file_access"] ? true : false);
        A.store.Bool(ptr + 81, "wants_file_access" in x ? true : false);
        A.store.Bool(ptr + 26, x["wants_file_access"] ? true : false);
        A.store.Bool(ptr + 82, "incognito_enabled" in x ? true : false);
        A.store.Bool(ptr + 27, x["incognito_enabled"] ? true : false);
        A.store.Bool(ptr + 83, "is_unpacked" in x ? true : false);
        A.store.Bool(ptr + 28, x["is_unpacked"] ? true : false);
        A.store.Bool(ptr + 84, "allow_reload" in x ? true : false);
        A.store.Bool(ptr + 29, x["allow_reload"] ? true : false);
        A.store.Bool(ptr + 85, "terminated" in x ? true : false);
        A.store.Bool(ptr + 30, x["terminated"] ? true : false);
        A.store.Bool(ptr + 86, "allow_incognito" in x ? true : false);
        A.store.Bool(ptr + 31, x["allow_incognito"] ? true : false);
        A.store.Ref(ptr + 32, x["icon_url"]);
        A.store.Ref(ptr + 36, x["path"]);
        A.store.Ref(ptr + 40, x["options_url"]);
        A.store.Ref(ptr + 44, x["app_launch_url"]);
        A.store.Ref(ptr + 48, x["homepage_url"]);
        A.store.Ref(ptr + 52, x["update_url"]);
        A.store.Ref(ptr + 56, x["install_warnings"]);
        A.store.Ref(ptr + 60, x["manifest_errors"]);
        A.store.Ref(ptr + 64, x["runtime_errors"]);
        A.store.Bool(ptr + 87, "offline_enabled" in x ? true : false);
        A.store.Bool(ptr + 68, x["offline_enabled"] ? true : false);
        A.store.Ref(ptr + 72, x["views"]);
      }
    },
    "load_ItemInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      x["version"] = A.load.Ref(ptr + 8, undefined);
      x["description"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 76)) {
        x["may_disable"] = A.load.Bool(ptr + 16);
      } else {
        delete x["may_disable"];
      }
      if (A.load.Bool(ptr + 77)) {
        x["enabled"] = A.load.Bool(ptr + 17);
      } else {
        delete x["enabled"];
      }
      if (A.load.Bool(ptr + 78)) {
        x["isApp"] = A.load.Bool(ptr + 18);
      } else {
        delete x["isApp"];
      }
      x["type"] = A.load.Enum(ptr + 20, ["hosted_app", "packaged_app", "legacy_packaged_app", "extension", "theme"]);
      if (A.load.Bool(ptr + 79)) {
        x["allow_activity"] = A.load.Bool(ptr + 24);
      } else {
        delete x["allow_activity"];
      }
      if (A.load.Bool(ptr + 80)) {
        x["allow_file_access"] = A.load.Bool(ptr + 25);
      } else {
        delete x["allow_file_access"];
      }
      if (A.load.Bool(ptr + 81)) {
        x["wants_file_access"] = A.load.Bool(ptr + 26);
      } else {
        delete x["wants_file_access"];
      }
      if (A.load.Bool(ptr + 82)) {
        x["incognito_enabled"] = A.load.Bool(ptr + 27);
      } else {
        delete x["incognito_enabled"];
      }
      if (A.load.Bool(ptr + 83)) {
        x["is_unpacked"] = A.load.Bool(ptr + 28);
      } else {
        delete x["is_unpacked"];
      }
      if (A.load.Bool(ptr + 84)) {
        x["allow_reload"] = A.load.Bool(ptr + 29);
      } else {
        delete x["allow_reload"];
      }
      if (A.load.Bool(ptr + 85)) {
        x["terminated"] = A.load.Bool(ptr + 30);
      } else {
        delete x["terminated"];
      }
      if (A.load.Bool(ptr + 86)) {
        x["allow_incognito"] = A.load.Bool(ptr + 31);
      } else {
        delete x["allow_incognito"];
      }
      x["icon_url"] = A.load.Ref(ptr + 32, undefined);
      x["path"] = A.load.Ref(ptr + 36, undefined);
      x["options_url"] = A.load.Ref(ptr + 40, undefined);
      x["app_launch_url"] = A.load.Ref(ptr + 44, undefined);
      x["homepage_url"] = A.load.Ref(ptr + 48, undefined);
      x["update_url"] = A.load.Ref(ptr + 52, undefined);
      x["install_warnings"] = A.load.Ref(ptr + 56, undefined);
      x["manifest_errors"] = A.load.Ref(ptr + 60, undefined);
      x["runtime_errors"] = A.load.Ref(ptr + 64, undefined);
      if (A.load.Bool(ptr + 87)) {
        x["offline_enabled"] = A.load.Bool(ptr + 68);
      } else {
        delete x["offline_enabled"];
      }
      x["views"] = A.load.Ref(ptr + 72, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_LoadError": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);

        A.store.Bool(ptr + 8 + 12, false);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Ref(ptr + 8 + 4, undefined);
        A.store.Ref(ptr + 8 + 8, undefined);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 28, true);
        A.store.Ref(ptr + 0, x["error"]);
        A.store.Ref(ptr + 4, x["path"]);

        if (typeof x["source"] === "undefined") {
          A.store.Bool(ptr + 8 + 12, false);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Ref(ptr + 8 + 4, undefined);
          A.store.Ref(ptr + 8 + 8, undefined);
        } else {
          A.store.Bool(ptr + 8 + 12, true);
          A.store.Ref(ptr + 8 + 0, x["source"]["beforeHighlight"]);
          A.store.Ref(ptr + 8 + 4, x["source"]["highlight"]);
          A.store.Ref(ptr + 8 + 8, x["source"]["afterHighlight"]);
        }
        A.store.Ref(ptr + 24, x["retryGuid"]);
      }
    },
    "load_LoadError": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["error"] = A.load.Ref(ptr + 0, undefined);
      x["path"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 8 + 12)) {
        x["source"] = {};
        x["source"]["beforeHighlight"] = A.load.Ref(ptr + 8 + 0, undefined);
        x["source"]["highlight"] = A.load.Ref(ptr + 8 + 4, undefined);
        x["source"]["afterHighlight"] = A.load.Ref(ptr + 8 + 8, undefined);
      } else {
        delete x["source"];
      }
      x["retryGuid"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_LoadUnpackedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 1, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Bool(ptr + 9, "failQuietly" in x ? true : false);
        A.store.Bool(ptr + 0, x["failQuietly"] ? true : false);
        A.store.Bool(ptr + 10, "populateError" in x ? true : false);
        A.store.Bool(ptr + 1, x["populateError"] ? true : false);
        A.store.Ref(ptr + 4, x["retryGuid"]);
        A.store.Bool(ptr + 11, "useDraggedPath" in x ? true : false);
        A.store.Bool(ptr + 8, x["useDraggedPath"] ? true : false);
      }
    },
    "load_LoadUnpackedOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 9)) {
        x["failQuietly"] = A.load.Bool(ptr + 0);
      } else {
        delete x["failQuietly"];
      }
      if (A.load.Bool(ptr + 10)) {
        x["populateError"] = A.load.Bool(ptr + 1);
      } else {
        delete x["populateError"];
      }
      x["retryGuid"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 11)) {
        x["useDraggedPath"] = A.load.Bool(ptr + 8);
      } else {
        delete x["useDraggedPath"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OpenDevToolsProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 34, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 28, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 29, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 30, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 31, false);
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 32, false);
        A.store.Int32(ptr + 20, 0);
        A.store.Bool(ptr + 33, false);
        A.store.Int32(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 34, true);
        A.store.Ref(ptr + 0, x["extensionId"]);
        A.store.Bool(ptr + 28, "renderViewId" in x ? true : false);
        A.store.Int32(ptr + 4, x["renderViewId"] === undefined ? 0 : (x["renderViewId"] as number));
        A.store.Bool(ptr + 29, "renderProcessId" in x ? true : false);
        A.store.Int32(ptr + 8, x["renderProcessId"] === undefined ? 0 : (x["renderProcessId"] as number));
        A.store.Bool(ptr + 30, "isServiceWorker" in x ? true : false);
        A.store.Bool(ptr + 12, x["isServiceWorker"] ? true : false);
        A.store.Bool(ptr + 31, "incognito" in x ? true : false);
        A.store.Bool(ptr + 13, x["incognito"] ? true : false);
        A.store.Ref(ptr + 16, x["url"]);
        A.store.Bool(ptr + 32, "lineNumber" in x ? true : false);
        A.store.Int32(ptr + 20, x["lineNumber"] === undefined ? 0 : (x["lineNumber"] as number));
        A.store.Bool(ptr + 33, "columnNumber" in x ? true : false);
        A.store.Int32(ptr + 24, x["columnNumber"] === undefined ? 0 : (x["columnNumber"] as number));
      }
    },
    "load_OpenDevToolsProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["extensionId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 28)) {
        x["renderViewId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["renderViewId"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["renderProcessId"] = A.load.Int32(ptr + 8);
      } else {
        delete x["renderProcessId"];
      }
      if (A.load.Bool(ptr + 30)) {
        x["isServiceWorker"] = A.load.Bool(ptr + 12);
      } else {
        delete x["isServiceWorker"];
      }
      if (A.load.Bool(ptr + 31)) {
        x["incognito"] = A.load.Bool(ptr + 13);
      } else {
        delete x["incognito"];
      }
      x["url"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 32)) {
        x["lineNumber"] = A.load.Int32(ptr + 20);
      } else {
        delete x["lineNumber"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["columnNumber"] = A.load.Int32(ptr + 24);
      } else {
        delete x["columnNumber"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PackStatus": (ref: heap.Ref<string>): number => {
      const idx = ["SUCCESS", "ERROR", "WARNING"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PackDirectoryResponse": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 20, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Enum(ptr + 16, -1);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Ref(ptr + 0, x["message"]);
        A.store.Ref(ptr + 4, x["item_path"]);
        A.store.Ref(ptr + 8, x["pem_path"]);
        A.store.Bool(ptr + 20, "override_flags" in x ? true : false);
        A.store.Int32(ptr + 12, x["override_flags"] === undefined ? 0 : (x["override_flags"] as number));
        A.store.Enum(ptr + 16, ["SUCCESS", "ERROR", "WARNING"].indexOf(x["status"] as string));
      }
    },
    "load_PackDirectoryResponse": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["message"] = A.load.Ref(ptr + 0, undefined);
      x["item_path"] = A.load.Ref(ptr + 4, undefined);
      x["pem_path"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 20)) {
        x["override_flags"] = A.load.Int32(ptr + 12);
      } else {
        delete x["override_flags"];
      }
      x["status"] = A.load.Enum(ptr + 16, ["SUCCESS", "ERROR", "WARNING"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ProfileConfigurationUpdate": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 1, "inDeveloperMode" in x ? true : false);
        A.store.Bool(ptr + 0, x["inDeveloperMode"] ? true : false);
      }
    },
    "load_ProfileConfigurationUpdate": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 1)) {
        x["inDeveloperMode"] = A.load.Bool(ptr + 0);
      } else {
        delete x["inDeveloperMode"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ProfileInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 5, "canLoadUnpacked" in x ? true : false);
        A.store.Bool(ptr + 0, x["canLoadUnpacked"] ? true : false);
        A.store.Bool(ptr + 6, "inDeveloperMode" in x ? true : false);
        A.store.Bool(ptr + 1, x["inDeveloperMode"] ? true : false);
        A.store.Bool(ptr + 7, "isDeveloperModeControlledByPolicy" in x ? true : false);
        A.store.Bool(ptr + 2, x["isDeveloperModeControlledByPolicy"] ? true : false);
        A.store.Bool(ptr + 8, "isIncognitoAvailable" in x ? true : false);
        A.store.Bool(ptr + 3, x["isIncognitoAvailable"] ? true : false);
        A.store.Bool(ptr + 9, "isChildAccount" in x ? true : false);
        A.store.Bool(ptr + 4, x["isChildAccount"] ? true : false);
      }
    },
    "load_ProfileInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 5)) {
        x["canLoadUnpacked"] = A.load.Bool(ptr + 0);
      } else {
        delete x["canLoadUnpacked"];
      }
      if (A.load.Bool(ptr + 6)) {
        x["inDeveloperMode"] = A.load.Bool(ptr + 1);
      } else {
        delete x["inDeveloperMode"];
      }
      if (A.load.Bool(ptr + 7)) {
        x["isDeveloperModeControlledByPolicy"] = A.load.Bool(ptr + 2);
      } else {
        delete x["isDeveloperModeControlledByPolicy"];
      }
      if (A.load.Bool(ptr + 8)) {
        x["isIncognitoAvailable"] = A.load.Bool(ptr + 3);
      } else {
        delete x["isIncognitoAvailable"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["isChildAccount"] = A.load.Bool(ptr + 4);
      } else {
        delete x["isChildAccount"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ReloadOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 1, false);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Bool(ptr + 2, "failQuietly" in x ? true : false);
        A.store.Bool(ptr + 0, x["failQuietly"] ? true : false);
        A.store.Bool(ptr + 3, "populateErrorForUnpacked" in x ? true : false);
        A.store.Bool(ptr + 1, x["populateErrorForUnpacked"] ? true : false);
      }
    },
    "load_ReloadOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 2)) {
        x["failQuietly"] = A.load.Bool(ptr + 0);
      } else {
        delete x["failQuietly"];
      }
      if (A.load.Bool(ptr + 3)) {
        x["populateErrorForUnpacked"] = A.load.Bool(ptr + 1);
      } else {
        delete x["populateErrorForUnpacked"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RequestFileSourceResponse": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Ref(ptr + 0, x["highlight"]);
        A.store.Ref(ptr + 4, x["beforeHighlight"]);
        A.store.Ref(ptr + 8, x["afterHighlight"]);
        A.store.Ref(ptr + 12, x["title"]);
        A.store.Ref(ptr + 16, x["message"]);
      }
    },
    "load_RequestFileSourceResponse": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["highlight"] = A.load.Ref(ptr + 0, undefined);
      x["beforeHighlight"] = A.load.Ref(ptr + 4, undefined);
      x["afterHighlight"] = A.load.Ref(ptr + 8, undefined);
      x["title"] = A.load.Ref(ptr + 12, undefined);
      x["message"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RequestFileSourceProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 24, false);
        A.store.Int32(ptr + 20, 0);
      } else {
        A.store.Bool(ptr + 25, true);
        A.store.Ref(ptr + 0, x["extensionId"]);
        A.store.Ref(ptr + 4, x["pathSuffix"]);
        A.store.Ref(ptr + 8, x["message"]);
        A.store.Ref(ptr + 12, x["manifestKey"]);
        A.store.Ref(ptr + 16, x["manifestSpecific"]);
        A.store.Bool(ptr + 24, "lineNumber" in x ? true : false);
        A.store.Int32(ptr + 20, x["lineNumber"] === undefined ? 0 : (x["lineNumber"] as number));
      }
    },
    "load_RequestFileSourceProperties": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["extensionId"] = A.load.Ref(ptr + 0, undefined);
      x["pathSuffix"] = A.load.Ref(ptr + 4, undefined);
      x["message"] = A.load.Ref(ptr + 8, undefined);
      x["manifestKey"] = A.load.Ref(ptr + 12, undefined);
      x["manifestSpecific"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 24)) {
        x["lineNumber"] = A.load.Int32(ptr + 20);
      } else {
        delete x["lineNumber"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_SelectType": (ref: heap.Ref<string>): number => {
      const idx = ["FILE", "FOLDER"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SiteSet": (ref: heap.Ref<string>): number => {
      const idx = ["USER_PERMITTED", "USER_RESTRICTED", "EXTENSION_SPECIFIED"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SiteInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Enum(
          ptr + 0,
          ["USER_PERMITTED", "USER_RESTRICTED", "EXTENSION_SPECIFIED"].indexOf(x["siteSet"] as string)
        );
        A.store.Bool(ptr + 12, "numExtensions" in x ? true : false);
        A.store.Int32(ptr + 4, x["numExtensions"] === undefined ? 0 : (x["numExtensions"] as number));
        A.store.Ref(ptr + 8, x["site"]);
      }
    },
    "load_SiteInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["siteSet"] = A.load.Enum(ptr + 0, ["USER_PERMITTED", "USER_RESTRICTED", "EXTENSION_SPECIFIED"]);
      if (A.load.Bool(ptr + 12)) {
        x["numExtensions"] = A.load.Int32(ptr + 4);
      } else {
        delete x["numExtensions"];
      }
      x["site"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SiteGroup": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Ref(ptr + 0, x["etldPlusOne"]);
        A.store.Bool(ptr + 12, "numExtensions" in x ? true : false);
        A.store.Int32(ptr + 4, x["numExtensions"] === undefined ? 0 : (x["numExtensions"] as number));
        A.store.Ref(ptr + 8, x["sites"]);
      }
    },
    "load_SiteGroup": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["etldPlusOne"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["numExtensions"] = A.load.Int32(ptr + 4);
      } else {
        delete x["numExtensions"];
      }
      x["sites"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UserSiteSettings": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["permittedSites"]);
        A.store.Ref(ptr + 4, x["restrictedSites"]);
      }
    },
    "load_UserSiteSettings": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["permittedSites"] = A.load.Ref(ptr + 0, undefined);
      x["restrictedSites"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UserSiteSettingsOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(
          ptr + 0,
          ["USER_PERMITTED", "USER_RESTRICTED", "EXTENSION_SPECIFIED"].indexOf(x["siteSet"] as string)
        );
        A.store.Ref(ptr + 4, x["hosts"]);
      }
    },
    "load_UserSiteSettingsOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["siteSet"] = A.load.Enum(ptr + 0, ["USER_PERMITTED", "USER_RESTRICTED", "EXTENSION_SPECIFIED"]);
      x["hosts"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_AddHostPermission": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "addHostPermission" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddHostPermission": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.addHostPermission);
    },
    "call_AddHostPermission": (retPtr: Pointer, extensionId: heap.Ref<object>, host: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.addHostPermission(A.H.get<object>(extensionId), A.H.get<object>(host));
      A.store.Ref(retPtr, _ret);
    },
    "try_AddHostPermission": (
      retPtr: Pointer,
      errPtr: Pointer,
      extensionId: heap.Ref<object>,
      host: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.addHostPermission(A.H.get<object>(extensionId), A.H.get<object>(host));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AddUserSpecifiedSites": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "addUserSpecifiedSites" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddUserSpecifiedSites": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.addUserSpecifiedSites);
    },
    "call_AddUserSpecifiedSites": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["siteSet"] = A.load.Enum(options + 0, ["USER_PERMITTED", "USER_RESTRICTED", "EXTENSION_SPECIFIED"]);
      options_ffi["hosts"] = A.load.Ref(options + 4, undefined);

      const _ret = WEBEXT.developerPrivate.addUserSpecifiedSites(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_AddUserSpecifiedSites": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["siteSet"] = A.load.Enum(options + 0, ["USER_PERMITTED", "USER_RESTRICTED", "EXTENSION_SPECIFIED"]);
        options_ffi["hosts"] = A.load.Ref(options + 4, undefined);

        const _ret = WEBEXT.developerPrivate.addUserSpecifiedSites(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutoUpdate": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "autoUpdate" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutoUpdate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.autoUpdate);
    },
    "call_AutoUpdate": (retPtr: Pointer): void => {
      const _ret = WEBEXT.developerPrivate.autoUpdate();
      A.store.Ref(retPtr, _ret);
    },
    "try_AutoUpdate": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.autoUpdate();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ChoosePath": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "choosePath" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ChoosePath": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.choosePath);
    },
    "call_ChoosePath": (retPtr: Pointer, selectType: number, fileType: number): void => {
      const _ret = WEBEXT.developerPrivate.choosePath(
        selectType > 0 && selectType <= 2 ? ["FILE", "FOLDER"][selectType - 1] : undefined,
        fileType > 0 && fileType <= 2 ? ["LOAD", "PEM"][fileType - 1] : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_ChoosePath": (retPtr: Pointer, errPtr: Pointer, selectType: number, fileType: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.choosePath(
          selectType > 0 && selectType <= 2 ? ["FILE", "FOLDER"][selectType - 1] : undefined,
          fileType > 0 && fileType <= 2 ? ["LOAD", "PEM"][fileType - 1] : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DeleteExtensionErrors": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "deleteExtensionErrors" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DeleteExtensionErrors": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.deleteExtensionErrors);
    },
    "call_DeleteExtensionErrors": (retPtr: Pointer, properties: Pointer): void => {
      const properties_ffi = {};

      properties_ffi["extensionId"] = A.load.Ref(properties + 0, undefined);
      properties_ffi["errorIds"] = A.load.Ref(properties + 4, undefined);
      properties_ffi["type"] = A.load.Enum(properties + 8, ["MANIFEST", "RUNTIME"]);

      const _ret = WEBEXT.developerPrivate.deleteExtensionErrors(properties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_DeleteExtensionErrors": (retPtr: Pointer, errPtr: Pointer, properties: Pointer): heap.Ref<boolean> => {
      try {
        const properties_ffi = {};

        properties_ffi["extensionId"] = A.load.Ref(properties + 0, undefined);
        properties_ffi["errorIds"] = A.load.Ref(properties + 4, undefined);
        properties_ffi["type"] = A.load.Enum(properties + 8, ["MANIFEST", "RUNTIME"]);

        const _ret = WEBEXT.developerPrivate.deleteExtensionErrors(properties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetExtensionInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "getExtensionInfo" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetExtensionInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.getExtensionInfo);
    },
    "call_GetExtensionInfo": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.getExtensionInfo(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetExtensionInfo": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.getExtensionInfo(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetExtensionSize": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "getExtensionSize" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetExtensionSize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.getExtensionSize);
    },
    "call_GetExtensionSize": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.getExtensionSize(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetExtensionSize": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.getExtensionSize(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetExtensionsInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "getExtensionsInfo" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetExtensionsInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.getExtensionsInfo);
    },
    "call_GetExtensionsInfo": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 2)) {
        options_ffi["includeDisabled"] = A.load.Bool(options + 0);
      }
      if (A.load.Bool(options + 3)) {
        options_ffi["includeTerminated"] = A.load.Bool(options + 1);
      }

      const _ret = WEBEXT.developerPrivate.getExtensionsInfo(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetExtensionsInfo": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 2)) {
          options_ffi["includeDisabled"] = A.load.Bool(options + 0);
        }
        if (A.load.Bool(options + 3)) {
          options_ffi["includeTerminated"] = A.load.Bool(options + 1);
        }

        const _ret = WEBEXT.developerPrivate.getExtensionsInfo(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetMatchingExtensionsForSite": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "getMatchingExtensionsForSite" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetMatchingExtensionsForSite": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.getMatchingExtensionsForSite);
    },
    "call_GetMatchingExtensionsForSite": (retPtr: Pointer, site: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.getMatchingExtensionsForSite(A.H.get<object>(site));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetMatchingExtensionsForSite": (
      retPtr: Pointer,
      errPtr: Pointer,
      site: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.getMatchingExtensionsForSite(A.H.get<object>(site));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetProfileConfiguration": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "getProfileConfiguration" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetProfileConfiguration": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.getProfileConfiguration);
    },
    "call_GetProfileConfiguration": (retPtr: Pointer): void => {
      const _ret = WEBEXT.developerPrivate.getProfileConfiguration();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetProfileConfiguration": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.getProfileConfiguration();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetUserAndExtensionSitesByEtld": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "getUserAndExtensionSitesByEtld" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetUserAndExtensionSitesByEtld": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.getUserAndExtensionSitesByEtld);
    },
    "call_GetUserAndExtensionSitesByEtld": (retPtr: Pointer): void => {
      const _ret = WEBEXT.developerPrivate.getUserAndExtensionSitesByEtld();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetUserAndExtensionSitesByEtld": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.getUserAndExtensionSitesByEtld();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetUserSiteSettings": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "getUserSiteSettings" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetUserSiteSettings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.getUserSiteSettings);
    },
    "call_GetUserSiteSettings": (retPtr: Pointer): void => {
      const _ret = WEBEXT.developerPrivate.getUserSiteSettings();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetUserSiteSettings": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.getUserSiteSettings();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Inspect": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "inspect" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Inspect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.inspect);
    },
    "call_Inspect": (retPtr: Pointer, options: Pointer, callback: heap.Ref<object>): void => {
      const options_ffi = {};

      options_ffi["extension_id"] = A.load.Ref(options + 0, undefined);
      options_ffi["render_process_id"] = A.load.Ref(options + 4, undefined);
      options_ffi["render_view_id"] = A.load.Ref(options + 8, undefined);
      if (A.load.Bool(options + 13)) {
        options_ffi["incognito"] = A.load.Bool(options + 12);
      }

      const _ret = WEBEXT.developerPrivate.inspect(options_ffi, A.H.get<object>(callback));
    },
    "try_Inspect": (
      retPtr: Pointer,
      errPtr: Pointer,
      options: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["extension_id"] = A.load.Ref(options + 0, undefined);
        options_ffi["render_process_id"] = A.load.Ref(options + 4, undefined);
        options_ffi["render_view_id"] = A.load.Ref(options + 8, undefined);
        if (A.load.Bool(options + 13)) {
          options_ffi["incognito"] = A.load.Bool(options + 12);
        }

        const _ret = WEBEXT.developerPrivate.inspect(options_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InstallDroppedFile": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "installDroppedFile" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InstallDroppedFile": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.installDroppedFile);
    },
    "call_InstallDroppedFile": (retPtr: Pointer): void => {
      const _ret = WEBEXT.developerPrivate.installDroppedFile();
      A.store.Ref(retPtr, _ret);
    },
    "try_InstallDroppedFile": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.installDroppedFile();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsProfileManaged": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "isProfileManaged" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsProfileManaged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.isProfileManaged);
    },
    "call_IsProfileManaged": (retPtr: Pointer): void => {
      const _ret = WEBEXT.developerPrivate.isProfileManaged();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsProfileManaged": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.isProfileManaged();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LoadDirectory": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "loadDirectory" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LoadDirectory": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.loadDirectory);
    },
    "call_LoadDirectory": (retPtr: Pointer, directory: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.loadDirectory(A.H.get<object>(directory));
      A.store.Ref(retPtr, _ret);
    },
    "try_LoadDirectory": (retPtr: Pointer, errPtr: Pointer, directory: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.loadDirectory(A.H.get<object>(directory));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LoadUnpacked": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "loadUnpacked" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LoadUnpacked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.loadUnpacked);
    },
    "call_LoadUnpacked": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 9)) {
        options_ffi["failQuietly"] = A.load.Bool(options + 0);
      }
      if (A.load.Bool(options + 10)) {
        options_ffi["populateError"] = A.load.Bool(options + 1);
      }
      options_ffi["retryGuid"] = A.load.Ref(options + 4, undefined);
      if (A.load.Bool(options + 11)) {
        options_ffi["useDraggedPath"] = A.load.Bool(options + 8);
      }

      const _ret = WEBEXT.developerPrivate.loadUnpacked(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_LoadUnpacked": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 9)) {
          options_ffi["failQuietly"] = A.load.Bool(options + 0);
        }
        if (A.load.Bool(options + 10)) {
          options_ffi["populateError"] = A.load.Bool(options + 1);
        }
        options_ffi["retryGuid"] = A.load.Ref(options + 4, undefined);
        if (A.load.Bool(options + 11)) {
          options_ffi["useDraggedPath"] = A.load.Bool(options + 8);
        }

        const _ret = WEBEXT.developerPrivate.loadUnpacked(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_NotifyDragInstallInProgress": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "notifyDragInstallInProgress" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_NotifyDragInstallInProgress": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.notifyDragInstallInProgress);
    },
    "call_NotifyDragInstallInProgress": (retPtr: Pointer): void => {
      const _ret = WEBEXT.developerPrivate.notifyDragInstallInProgress();
    },
    "try_NotifyDragInstallInProgress": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.notifyDragInstallInProgress();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnItemStateChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.developerPrivate?.onItemStateChanged &&
        "addListener" in WEBEXT?.developerPrivate?.onItemStateChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnItemStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.onItemStateChanged.addListener);
    },
    "call_OnItemStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.onItemStateChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnItemStateChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.onItemStateChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffItemStateChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.developerPrivate?.onItemStateChanged &&
        "removeListener" in WEBEXT?.developerPrivate?.onItemStateChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffItemStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.onItemStateChanged.removeListener);
    },
    "call_OffItemStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.onItemStateChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffItemStateChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.onItemStateChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnItemStateChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.developerPrivate?.onItemStateChanged &&
        "hasListener" in WEBEXT?.developerPrivate?.onItemStateChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnItemStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.onItemStateChanged.hasListener);
    },
    "call_HasOnItemStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.onItemStateChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnItemStateChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.onItemStateChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnProfileStateChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.developerPrivate?.onProfileStateChanged &&
        "addListener" in WEBEXT?.developerPrivate?.onProfileStateChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnProfileStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.onProfileStateChanged.addListener);
    },
    "call_OnProfileStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.onProfileStateChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnProfileStateChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.onProfileStateChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffProfileStateChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.developerPrivate?.onProfileStateChanged &&
        "removeListener" in WEBEXT?.developerPrivate?.onProfileStateChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffProfileStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.onProfileStateChanged.removeListener);
    },
    "call_OffProfileStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.onProfileStateChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffProfileStateChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.onProfileStateChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnProfileStateChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.developerPrivate?.onProfileStateChanged &&
        "hasListener" in WEBEXT?.developerPrivate?.onProfileStateChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnProfileStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.onProfileStateChanged.hasListener);
    },
    "call_HasOnProfileStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.onProfileStateChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnProfileStateChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.onProfileStateChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnUserSiteSettingsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.developerPrivate?.onUserSiteSettingsChanged &&
        "addListener" in WEBEXT?.developerPrivate?.onUserSiteSettingsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnUserSiteSettingsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.onUserSiteSettingsChanged.addListener);
    },
    "call_OnUserSiteSettingsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.onUserSiteSettingsChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnUserSiteSettingsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.onUserSiteSettingsChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffUserSiteSettingsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.developerPrivate?.onUserSiteSettingsChanged &&
        "removeListener" in WEBEXT?.developerPrivate?.onUserSiteSettingsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffUserSiteSettingsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.onUserSiteSettingsChanged.removeListener);
    },
    "call_OffUserSiteSettingsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.onUserSiteSettingsChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffUserSiteSettingsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.onUserSiteSettingsChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnUserSiteSettingsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.developerPrivate?.onUserSiteSettingsChanged &&
        "hasListener" in WEBEXT?.developerPrivate?.onUserSiteSettingsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnUserSiteSettingsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.onUserSiteSettingsChanged.hasListener);
    },
    "call_HasOnUserSiteSettingsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.onUserSiteSettingsChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnUserSiteSettingsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.onUserSiteSettingsChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenDevTools": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "openDevTools" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenDevTools": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.openDevTools);
    },
    "call_OpenDevTools": (retPtr: Pointer, properties: Pointer): void => {
      const properties_ffi = {};

      properties_ffi["extensionId"] = A.load.Ref(properties + 0, undefined);
      if (A.load.Bool(properties + 28)) {
        properties_ffi["renderViewId"] = A.load.Int32(properties + 4);
      }
      if (A.load.Bool(properties + 29)) {
        properties_ffi["renderProcessId"] = A.load.Int32(properties + 8);
      }
      if (A.load.Bool(properties + 30)) {
        properties_ffi["isServiceWorker"] = A.load.Bool(properties + 12);
      }
      if (A.load.Bool(properties + 31)) {
        properties_ffi["incognito"] = A.load.Bool(properties + 13);
      }
      properties_ffi["url"] = A.load.Ref(properties + 16, undefined);
      if (A.load.Bool(properties + 32)) {
        properties_ffi["lineNumber"] = A.load.Int32(properties + 20);
      }
      if (A.load.Bool(properties + 33)) {
        properties_ffi["columnNumber"] = A.load.Int32(properties + 24);
      }

      const _ret = WEBEXT.developerPrivate.openDevTools(properties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_OpenDevTools": (retPtr: Pointer, errPtr: Pointer, properties: Pointer): heap.Ref<boolean> => {
      try {
        const properties_ffi = {};

        properties_ffi["extensionId"] = A.load.Ref(properties + 0, undefined);
        if (A.load.Bool(properties + 28)) {
          properties_ffi["renderViewId"] = A.load.Int32(properties + 4);
        }
        if (A.load.Bool(properties + 29)) {
          properties_ffi["renderProcessId"] = A.load.Int32(properties + 8);
        }
        if (A.load.Bool(properties + 30)) {
          properties_ffi["isServiceWorker"] = A.load.Bool(properties + 12);
        }
        if (A.load.Bool(properties + 31)) {
          properties_ffi["incognito"] = A.load.Bool(properties + 13);
        }
        properties_ffi["url"] = A.load.Ref(properties + 16, undefined);
        if (A.load.Bool(properties + 32)) {
          properties_ffi["lineNumber"] = A.load.Int32(properties + 20);
        }
        if (A.load.Bool(properties + 33)) {
          properties_ffi["columnNumber"] = A.load.Int32(properties + 24);
        }

        const _ret = WEBEXT.developerPrivate.openDevTools(properties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_PackDirectory": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "packDirectory" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_PackDirectory": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.packDirectory);
    },
    "call_PackDirectory": (
      retPtr: Pointer,
      path: heap.Ref<object>,
      privateKeyPath: heap.Ref<object>,
      flags: number
    ): void => {
      const _ret = WEBEXT.developerPrivate.packDirectory(A.H.get<object>(path), A.H.get<object>(privateKeyPath), flags);
      A.store.Ref(retPtr, _ret);
    },
    "try_PackDirectory": (
      retPtr: Pointer,
      errPtr: Pointer,
      path: heap.Ref<object>,
      privateKeyPath: heap.Ref<object>,
      flags: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.packDirectory(
          A.H.get<object>(path),
          A.H.get<object>(privateKeyPath),
          flags
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Reload": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "reload" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Reload": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.reload);
    },
    "call_Reload": (retPtr: Pointer, extensionId: heap.Ref<object>, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 2)) {
        options_ffi["failQuietly"] = A.load.Bool(options + 0);
      }
      if (A.load.Bool(options + 3)) {
        options_ffi["populateErrorForUnpacked"] = A.load.Bool(options + 1);
      }

      const _ret = WEBEXT.developerPrivate.reload(A.H.get<object>(extensionId), options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Reload": (
      retPtr: Pointer,
      errPtr: Pointer,
      extensionId: heap.Ref<object>,
      options: Pointer
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 2)) {
          options_ffi["failQuietly"] = A.load.Bool(options + 0);
        }
        if (A.load.Bool(options + 3)) {
          options_ffi["populateErrorForUnpacked"] = A.load.Bool(options + 1);
        }

        const _ret = WEBEXT.developerPrivate.reload(A.H.get<object>(extensionId), options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveHostPermission": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "removeHostPermission" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveHostPermission": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.removeHostPermission);
    },
    "call_RemoveHostPermission": (retPtr: Pointer, extensionId: heap.Ref<object>, host: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.removeHostPermission(A.H.get<object>(extensionId), A.H.get<object>(host));
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveHostPermission": (
      retPtr: Pointer,
      errPtr: Pointer,
      extensionId: heap.Ref<object>,
      host: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.removeHostPermission(A.H.get<object>(extensionId), A.H.get<object>(host));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveMultipleExtensions": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "removeMultipleExtensions" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveMultipleExtensions": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.removeMultipleExtensions);
    },
    "call_RemoveMultipleExtensions": (retPtr: Pointer, extensionIds: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.removeMultipleExtensions(A.H.get<object>(extensionIds));
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveMultipleExtensions": (
      retPtr: Pointer,
      errPtr: Pointer,
      extensionIds: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.removeMultipleExtensions(A.H.get<object>(extensionIds));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveUserSpecifiedSites": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "removeUserSpecifiedSites" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveUserSpecifiedSites": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.removeUserSpecifiedSites);
    },
    "call_RemoveUserSpecifiedSites": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["siteSet"] = A.load.Enum(options + 0, ["USER_PERMITTED", "USER_RESTRICTED", "EXTENSION_SPECIFIED"]);
      options_ffi["hosts"] = A.load.Ref(options + 4, undefined);

      const _ret = WEBEXT.developerPrivate.removeUserSpecifiedSites(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveUserSpecifiedSites": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["siteSet"] = A.load.Enum(options + 0, ["USER_PERMITTED", "USER_RESTRICTED", "EXTENSION_SPECIFIED"]);
        options_ffi["hosts"] = A.load.Ref(options + 4, undefined);

        const _ret = WEBEXT.developerPrivate.removeUserSpecifiedSites(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RepairExtension": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "repairExtension" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RepairExtension": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.repairExtension);
    },
    "call_RepairExtension": (retPtr: Pointer, extensionId: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.repairExtension(A.H.get<object>(extensionId));
      A.store.Ref(retPtr, _ret);
    },
    "try_RepairExtension": (retPtr: Pointer, errPtr: Pointer, extensionId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.repairExtension(A.H.get<object>(extensionId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RequestFileSource": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "requestFileSource" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RequestFileSource": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.requestFileSource);
    },
    "call_RequestFileSource": (retPtr: Pointer, properties: Pointer): void => {
      const properties_ffi = {};

      properties_ffi["extensionId"] = A.load.Ref(properties + 0, undefined);
      properties_ffi["pathSuffix"] = A.load.Ref(properties + 4, undefined);
      properties_ffi["message"] = A.load.Ref(properties + 8, undefined);
      properties_ffi["manifestKey"] = A.load.Ref(properties + 12, undefined);
      properties_ffi["manifestSpecific"] = A.load.Ref(properties + 16, undefined);
      if (A.load.Bool(properties + 24)) {
        properties_ffi["lineNumber"] = A.load.Int32(properties + 20);
      }

      const _ret = WEBEXT.developerPrivate.requestFileSource(properties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RequestFileSource": (retPtr: Pointer, errPtr: Pointer, properties: Pointer): heap.Ref<boolean> => {
      try {
        const properties_ffi = {};

        properties_ffi["extensionId"] = A.load.Ref(properties + 0, undefined);
        properties_ffi["pathSuffix"] = A.load.Ref(properties + 4, undefined);
        properties_ffi["message"] = A.load.Ref(properties + 8, undefined);
        properties_ffi["manifestKey"] = A.load.Ref(properties + 12, undefined);
        properties_ffi["manifestSpecific"] = A.load.Ref(properties + 16, undefined);
        if (A.load.Bool(properties + 24)) {
          properties_ffi["lineNumber"] = A.load.Int32(properties + 20);
        }

        const _ret = WEBEXT.developerPrivate.requestFileSource(properties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetShortcutHandlingSuspended": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "setShortcutHandlingSuspended" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetShortcutHandlingSuspended": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.setShortcutHandlingSuspended);
    },
    "call_SetShortcutHandlingSuspended": (retPtr: Pointer, isSuspended: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.developerPrivate.setShortcutHandlingSuspended(isSuspended === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetShortcutHandlingSuspended": (
      retPtr: Pointer,
      errPtr: Pointer,
      isSuspended: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.setShortcutHandlingSuspended(isSuspended === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ShowOptions": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "showOptions" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ShowOptions": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.showOptions);
    },
    "call_ShowOptions": (retPtr: Pointer, extensionId: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.showOptions(A.H.get<object>(extensionId));
      A.store.Ref(retPtr, _ret);
    },
    "try_ShowOptions": (retPtr: Pointer, errPtr: Pointer, extensionId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.showOptions(A.H.get<object>(extensionId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ShowPath": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "showPath" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ShowPath": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.showPath);
    },
    "call_ShowPath": (retPtr: Pointer, extensionId: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.showPath(A.H.get<object>(extensionId));
      A.store.Ref(retPtr, _ret);
    },
    "try_ShowPath": (retPtr: Pointer, errPtr: Pointer, extensionId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.showPath(A.H.get<object>(extensionId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateExtensionCommand": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "updateExtensionCommand" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateExtensionCommand": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.updateExtensionCommand);
    },
    "call_UpdateExtensionCommand": (retPtr: Pointer, update: Pointer): void => {
      const update_ffi = {};

      update_ffi["extensionId"] = A.load.Ref(update + 0, undefined);
      update_ffi["commandName"] = A.load.Ref(update + 4, undefined);
      update_ffi["scope"] = A.load.Enum(update + 8, ["GLOBAL", "CHROME"]);
      update_ffi["keybinding"] = A.load.Ref(update + 12, undefined);

      const _ret = WEBEXT.developerPrivate.updateExtensionCommand(update_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_UpdateExtensionCommand": (retPtr: Pointer, errPtr: Pointer, update: Pointer): heap.Ref<boolean> => {
      try {
        const update_ffi = {};

        update_ffi["extensionId"] = A.load.Ref(update + 0, undefined);
        update_ffi["commandName"] = A.load.Ref(update + 4, undefined);
        update_ffi["scope"] = A.load.Enum(update + 8, ["GLOBAL", "CHROME"]);
        update_ffi["keybinding"] = A.load.Ref(update + 12, undefined);

        const _ret = WEBEXT.developerPrivate.updateExtensionCommand(update_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateExtensionConfiguration": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "updateExtensionConfiguration" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateExtensionConfiguration": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.updateExtensionConfiguration);
    },
    "call_UpdateExtensionConfiguration": (retPtr: Pointer, update: Pointer): void => {
      const update_ffi = {};

      update_ffi["extensionId"] = A.load.Ref(update + 0, undefined);
      if (A.load.Bool(update + 15)) {
        update_ffi["fileAccess"] = A.load.Bool(update + 4);
      }
      if (A.load.Bool(update + 16)) {
        update_ffi["incognitoAccess"] = A.load.Bool(update + 5);
      }
      if (A.load.Bool(update + 17)) {
        update_ffi["errorCollection"] = A.load.Bool(update + 6);
      }
      update_ffi["hostAccess"] = A.load.Enum(update + 8, ["ON_CLICK", "ON_SPECIFIC_SITES", "ON_ALL_SITES"]);
      if (A.load.Bool(update + 18)) {
        update_ffi["showAccessRequestsInToolbar"] = A.load.Bool(update + 12);
      }
      if (A.load.Bool(update + 19)) {
        update_ffi["acknowledgeSafetyCheckWarning"] = A.load.Bool(update + 13);
      }
      if (A.load.Bool(update + 20)) {
        update_ffi["pinnedToToolbar"] = A.load.Bool(update + 14);
      }

      const _ret = WEBEXT.developerPrivate.updateExtensionConfiguration(update_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_UpdateExtensionConfiguration": (retPtr: Pointer, errPtr: Pointer, update: Pointer): heap.Ref<boolean> => {
      try {
        const update_ffi = {};

        update_ffi["extensionId"] = A.load.Ref(update + 0, undefined);
        if (A.load.Bool(update + 15)) {
          update_ffi["fileAccess"] = A.load.Bool(update + 4);
        }
        if (A.load.Bool(update + 16)) {
          update_ffi["incognitoAccess"] = A.load.Bool(update + 5);
        }
        if (A.load.Bool(update + 17)) {
          update_ffi["errorCollection"] = A.load.Bool(update + 6);
        }
        update_ffi["hostAccess"] = A.load.Enum(update + 8, ["ON_CLICK", "ON_SPECIFIC_SITES", "ON_ALL_SITES"]);
        if (A.load.Bool(update + 18)) {
          update_ffi["showAccessRequestsInToolbar"] = A.load.Bool(update + 12);
        }
        if (A.load.Bool(update + 19)) {
          update_ffi["acknowledgeSafetyCheckWarning"] = A.load.Bool(update + 13);
        }
        if (A.load.Bool(update + 20)) {
          update_ffi["pinnedToToolbar"] = A.load.Bool(update + 14);
        }

        const _ret = WEBEXT.developerPrivate.updateExtensionConfiguration(update_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateProfileConfiguration": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "updateProfileConfiguration" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateProfileConfiguration": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.updateProfileConfiguration);
    },
    "call_UpdateProfileConfiguration": (retPtr: Pointer, update: Pointer): void => {
      const update_ffi = {};

      if (A.load.Bool(update + 1)) {
        update_ffi["inDeveloperMode"] = A.load.Bool(update + 0);
      }

      const _ret = WEBEXT.developerPrivate.updateProfileConfiguration(update_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_UpdateProfileConfiguration": (retPtr: Pointer, errPtr: Pointer, update: Pointer): heap.Ref<boolean> => {
      try {
        const update_ffi = {};

        if (A.load.Bool(update + 1)) {
          update_ffi["inDeveloperMode"] = A.load.Bool(update + 0);
        }

        const _ret = WEBEXT.developerPrivate.updateProfileConfiguration(update_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateSiteAccess": (): heap.Ref<boolean> => {
      if (WEBEXT?.developerPrivate && "updateSiteAccess" in WEBEXT?.developerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateSiteAccess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.developerPrivate.updateSiteAccess);
    },
    "call_UpdateSiteAccess": (retPtr: Pointer, site: heap.Ref<object>, updates: heap.Ref<object>): void => {
      const _ret = WEBEXT.developerPrivate.updateSiteAccess(A.H.get<object>(site), A.H.get<object>(updates));
      A.store.Ref(retPtr, _ret);
    },
    "try_UpdateSiteAccess": (
      retPtr: Pointer,
      errPtr: Pointer,
      site: heap.Ref<object>,
      updates: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.developerPrivate.updateSiteAccess(A.H.get<object>(site), A.H.get<object>(updates));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
