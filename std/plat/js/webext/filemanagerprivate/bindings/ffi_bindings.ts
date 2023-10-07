import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/filemanagerprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_IconSet": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["icon16x16Url"]);
        A.store.Ref(ptr + 4, x["icon32x32Url"]);
      }
    },
    "load_IconSet": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["icon16x16Url"] = A.load.Ref(ptr + 0, undefined);
      x["icon32x32Url"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AndroidApp": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);

        A.store.Bool(ptr + 12 + 8, false);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 4, undefined);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["packageName"]);
        A.store.Ref(ptr + 8, x["activityName"]);

        if (typeof x["iconSet"] === "undefined") {
          A.store.Bool(ptr + 12 + 8, false);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 4, undefined);
        } else {
          A.store.Bool(ptr + 12 + 8, true);
          A.store.Ref(ptr + 12 + 0, x["iconSet"]["icon16x16Url"]);
          A.store.Ref(ptr + 12 + 4, x["iconSet"]["icon32x32Url"]);
        }
      }
    },
    "load_AndroidApp": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["packageName"] = A.load.Ref(ptr + 4, undefined);
      x["activityName"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 12 + 8)) {
        x["iconSet"] = {};
        x["iconSet"]["icon16x16Url"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["iconSet"]["icon32x32Url"] = A.load.Ref(ptr + 12 + 4, undefined);
      } else {
        delete x["iconSet"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AttachedImages": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["data"]);
        A.store.Ref(ptr + 4, x["type"]);
      }
    },
    "load_AttachedImages": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["data"] = A.load.Ref(ptr + 0, undefined);
      x["type"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_BulkPinStage": (ref: heap.Ref<string>): number => {
      const idx = [
        "stopped",
        "paused_offline",
        "paused_battery_saver",
        "getting_free_space",
        "listing_files",
        "syncing",
        "success",
        "not_enough_space",
        "cannot_get_free_space",
        "cannot_list_files",
        "cannot_enable_docs_offline",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_BulkPinProgress": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 65, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 57, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 58, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 59, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Bool(ptr + 60, false);
        A.store.Float64(ptr + 32, 0);
        A.store.Bool(ptr + 61, false);
        A.store.Int32(ptr + 40, 0);
        A.store.Bool(ptr + 62, false);
        A.store.Int32(ptr + 44, 0);
        A.store.Bool(ptr + 63, false);
        A.store.Float64(ptr + 48, 0);
        A.store.Bool(ptr + 64, false);
        A.store.Bool(ptr + 56, false);
      } else {
        A.store.Bool(ptr + 65, true);
        A.store.Enum(
          ptr + 0,
          [
            "stopped",
            "paused_offline",
            "paused_battery_saver",
            "getting_free_space",
            "listing_files",
            "syncing",
            "success",
            "not_enough_space",
            "cannot_get_free_space",
            "cannot_list_files",
            "cannot_enable_docs_offline",
          ].indexOf(x["stage"] as string)
        );
        A.store.Bool(ptr + 57, "freeSpaceBytes" in x ? true : false);
        A.store.Float64(ptr + 8, x["freeSpaceBytes"] === undefined ? 0 : (x["freeSpaceBytes"] as number));
        A.store.Bool(ptr + 58, "requiredSpaceBytes" in x ? true : false);
        A.store.Float64(ptr + 16, x["requiredSpaceBytes"] === undefined ? 0 : (x["requiredSpaceBytes"] as number));
        A.store.Bool(ptr + 59, "bytesToPin" in x ? true : false);
        A.store.Float64(ptr + 24, x["bytesToPin"] === undefined ? 0 : (x["bytesToPin"] as number));
        A.store.Bool(ptr + 60, "pinnedBytes" in x ? true : false);
        A.store.Float64(ptr + 32, x["pinnedBytes"] === undefined ? 0 : (x["pinnedBytes"] as number));
        A.store.Bool(ptr + 61, "filesToPin" in x ? true : false);
        A.store.Int32(ptr + 40, x["filesToPin"] === undefined ? 0 : (x["filesToPin"] as number));
        A.store.Bool(ptr + 62, "listedFiles" in x ? true : false);
        A.store.Int32(ptr + 44, x["listedFiles"] === undefined ? 0 : (x["listedFiles"] as number));
        A.store.Bool(ptr + 63, "remainingSeconds" in x ? true : false);
        A.store.Float64(ptr + 48, x["remainingSeconds"] === undefined ? 0 : (x["remainingSeconds"] as number));
        A.store.Bool(ptr + 64, "emptiedQueue" in x ? true : false);
        A.store.Bool(ptr + 56, x["emptiedQueue"] ? true : false);
      }
    },
    "load_BulkPinProgress": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["stage"] = A.load.Enum(ptr + 0, [
        "stopped",
        "paused_offline",
        "paused_battery_saver",
        "getting_free_space",
        "listing_files",
        "syncing",
        "success",
        "not_enough_space",
        "cannot_get_free_space",
        "cannot_list_files",
        "cannot_enable_docs_offline",
      ]);
      if (A.load.Bool(ptr + 57)) {
        x["freeSpaceBytes"] = A.load.Float64(ptr + 8);
      } else {
        delete x["freeSpaceBytes"];
      }
      if (A.load.Bool(ptr + 58)) {
        x["requiredSpaceBytes"] = A.load.Float64(ptr + 16);
      } else {
        delete x["requiredSpaceBytes"];
      }
      if (A.load.Bool(ptr + 59)) {
        x["bytesToPin"] = A.load.Float64(ptr + 24);
      } else {
        delete x["bytesToPin"];
      }
      if (A.load.Bool(ptr + 60)) {
        x["pinnedBytes"] = A.load.Float64(ptr + 32);
      } else {
        delete x["pinnedBytes"];
      }
      if (A.load.Bool(ptr + 61)) {
        x["filesToPin"] = A.load.Int32(ptr + 40);
      } else {
        delete x["filesToPin"];
      }
      if (A.load.Bool(ptr + 62)) {
        x["listedFiles"] = A.load.Int32(ptr + 44);
      } else {
        delete x["listedFiles"];
      }
      if (A.load.Bool(ptr + 63)) {
        x["remainingSeconds"] = A.load.Float64(ptr + 48);
      } else {
        delete x["remainingSeconds"];
      }
      if (A.load.Bool(ptr + 64)) {
        x["emptiedQueue"] = A.load.Bool(ptr + 56);
      } else {
        delete x["emptiedQueue"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ChangeType": (ref: heap.Ref<string>): number => {
      const idx = ["add_or_update", "delete"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ConflictPauseParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 5, false);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Ref(ptr + 0, x["conflictName"]);
        A.store.Bool(ptr + 12, "conflictIsDirectory" in x ? true : false);
        A.store.Bool(ptr + 4, x["conflictIsDirectory"] ? true : false);
        A.store.Bool(ptr + 13, "conflictMultiple" in x ? true : false);
        A.store.Bool(ptr + 5, x["conflictMultiple"] ? true : false);
        A.store.Ref(ptr + 8, x["conflictTargetUrl"]);
      }
    },
    "load_ConflictPauseParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["conflictName"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["conflictIsDirectory"] = A.load.Bool(ptr + 4);
      } else {
        delete x["conflictIsDirectory"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["conflictMultiple"] = A.load.Bool(ptr + 5);
      } else {
        delete x["conflictMultiple"];
      }
      x["conflictTargetUrl"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ConflictResumeParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 6, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 6, true);
        A.store.Ref(ptr + 0, x["conflictResolve"]);
        A.store.Bool(ptr + 5, "conflictApplyToAll" in x ? true : false);
        A.store.Bool(ptr + 4, x["conflictApplyToAll"] ? true : false);
      }
    },
    "load_ConflictResumeParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["conflictResolve"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 5)) {
        x["conflictApplyToAll"] = A.load.Bool(ptr + 4);
      } else {
        delete x["conflictApplyToAll"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_CrostiniEventType": (ref: heap.Ref<string>): number => {
      const idx = ["enable", "disable", "share", "unshare", "drop_failed_plugin_vm_directory_not_shared"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },

    "store_CrostiniEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Enum(
          ptr + 0,
          ["enable", "disable", "share", "unshare", "drop_failed_plugin_vm_directory_not_shared"].indexOf(
            x["eventType"] as string
          )
        );
        A.store.Ref(ptr + 4, x["vmName"]);
        A.store.Ref(ptr + 8, x["containerName"]);
        A.store.Ref(ptr + 12, x["entries"]);
      }
    },
    "load_CrostiniEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["eventType"] = A.load.Enum(ptr + 0, [
        "enable",
        "disable",
        "share",
        "unshare",
        "drop_failed_plugin_vm_directory_not_shared",
      ]);
      x["vmName"] = A.load.Ref(ptr + 4, undefined);
      x["containerName"] = A.load.Ref(ptr + 8, undefined);
      x["entries"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DeviceConnectionState": (ref: heap.Ref<string>): number => {
      const idx = ["OFFLINE", "ONLINE"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_DeviceEventType": (ref: heap.Ref<string>): number => {
      const idx = [
        "disabled",
        "removed",
        "hard_unplugged",
        "format_start",
        "format_success",
        "format_fail",
        "rename_start",
        "rename_success",
        "rename_fail",
        "partition_start",
        "partition_success",
        "partition_fail",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DeviceEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Enum(
          ptr + 0,
          [
            "disabled",
            "removed",
            "hard_unplugged",
            "format_start",
            "format_success",
            "format_fail",
            "rename_start",
            "rename_success",
            "rename_fail",
            "partition_start",
            "partition_success",
            "partition_fail",
          ].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 4, x["devicePath"]);
        A.store.Ref(ptr + 8, x["deviceLabel"]);
      }
    },
    "load_DeviceEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, [
        "disabled",
        "removed",
        "hard_unplugged",
        "format_start",
        "format_success",
        "format_fail",
        "rename_start",
        "rename_success",
        "rename_fail",
        "partition_start",
        "partition_success",
        "partition_fail",
      ]);
      x["devicePath"] = A.load.Ref(ptr + 4, undefined);
      x["deviceLabel"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DeviceType": (ref: heap.Ref<string>): number => {
      const idx = ["usb", "sd", "optical", "mobile", "unknown"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_VolumeType": (ref: heap.Ref<string>): number => {
      const idx = [
        "drive",
        "downloads",
        "removable",
        "archive",
        "provided",
        "mtp",
        "media_view",
        "crostini",
        "android_files",
        "documents_provider",
        "testing",
        "smb",
        "system_internal",
        "guest_os",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DialogCallerInformation": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["url"]);
        A.store.Enum(
          ptr + 4,
          [
            "drive",
            "downloads",
            "removable",
            "archive",
            "provided",
            "mtp",
            "media_view",
            "crostini",
            "android_files",
            "documents_provider",
            "testing",
            "smb",
            "system_internal",
            "guest_os",
          ].indexOf(x["component"] as string)
        );
      }
    },
    "load_DialogCallerInformation": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      x["component"] = A.load.Enum(ptr + 4, [
        "drive",
        "downloads",
        "removable",
        "archive",
        "provided",
        "mtp",
        "media_view",
        "crostini",
        "android_files",
        "documents_provider",
        "testing",
        "smb",
        "system_internal",
        "guest_os",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DlpLevel": (ref: heap.Ref<string>): number => {
      const idx = ["report", "warn", "block", "allow"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DlpMetadata": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 5, false);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["sourceUrl"]);
        A.store.Bool(ptr + 6, "isDlpRestricted" in x ? true : false);
        A.store.Bool(ptr + 4, x["isDlpRestricted"] ? true : false);
        A.store.Bool(ptr + 7, "isRestrictedForDestination" in x ? true : false);
        A.store.Bool(ptr + 5, x["isRestrictedForDestination"] ? true : false);
      }
    },
    "load_DlpMetadata": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["sourceUrl"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 6)) {
        x["isDlpRestricted"] = A.load.Bool(ptr + 4);
      } else {
        delete x["isDlpRestricted"];
      }
      if (A.load.Bool(ptr + 7)) {
        x["isRestrictedForDestination"] = A.load.Bool(ptr + 5);
      } else {
        delete x["isRestrictedForDestination"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DlpRestrictionDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Enum(ptr + 0, ["report", "warn", "block", "allow"].indexOf(x["level"] as string));
        A.store.Ref(ptr + 4, x["urls"]);
        A.store.Ref(ptr + 8, x["components"]);
      }
    },
    "load_DlpRestrictionDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["level"] = A.load.Enum(ptr + 0, ["report", "warn", "block", "allow"]);
      x["urls"] = A.load.Ref(ptr + 4, undefined);
      x["components"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DriveConfirmDialogType": (ref: heap.Ref<string>): number => {
      const idx = ["enable_docs_offline"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DriveConfirmDialogEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["enable_docs_offline"].indexOf(x["type"] as string));
        A.store.Ref(ptr + 4, x["fileUrl"]);
      }
    },
    "load_DriveConfirmDialogEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, ["enable_docs_offline"]);
      x["fileUrl"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DriveConnectionStateType": (ref: heap.Ref<string>): number => {
      const idx = ["OFFLINE", "METERED", "ONLINE"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_DriveOfflineReason": (ref: heap.Ref<string>): number => {
      const idx = ["NOT_READY", "NO_NETWORK", "NO_SERVICE"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DriveConnectionState": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["OFFLINE", "METERED", "ONLINE"].indexOf(x["type"] as string));
        A.store.Enum(ptr + 4, ["NOT_READY", "NO_NETWORK", "NO_SERVICE"].indexOf(x["reason"] as string));
      }
    },
    "load_DriveConnectionState": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, ["OFFLINE", "METERED", "ONLINE"]);
      x["reason"] = A.load.Enum(ptr + 4, ["NOT_READY", "NO_NETWORK", "NO_SERVICE"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DriveDialogResult": (ref: heap.Ref<string>): number => {
      const idx = ["not_displayed", "accept", "reject", "dismiss"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DriveMetadataSearchResult": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Ref(ptr + 0, x["entry"]);
        A.store.Ref(ptr + 4, x["highlightedBaseName"]);
        A.store.Bool(ptr + 9, "availableOffline" in x ? true : false);
        A.store.Bool(ptr + 8, x["availableOffline"] ? true : false);
      }
    },
    "load_DriveMetadataSearchResult": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["entry"] = A.load.Ref(ptr + 0, undefined);
      x["highlightedBaseName"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 9)) {
        x["availableOffline"] = A.load.Bool(ptr + 8);
      } else {
        delete x["availableOffline"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_UserType": (ref: heap.Ref<string>): number => {
      const idx = ["unmanaged", "organization"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DriveQuotaMetadata": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 35, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 32, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 33, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 34, false);
        A.store.Bool(ptr + 24, false);
        A.store.Ref(ptr + 28, undefined);
      } else {
        A.store.Bool(ptr + 35, true);
        A.store.Enum(ptr + 0, ["unmanaged", "organization"].indexOf(x["userType"] as string));
        A.store.Bool(ptr + 32, "usedBytes" in x ? true : false);
        A.store.Float64(ptr + 8, x["usedBytes"] === undefined ? 0 : (x["usedBytes"] as number));
        A.store.Bool(ptr + 33, "totalBytes" in x ? true : false);
        A.store.Float64(ptr + 16, x["totalBytes"] === undefined ? 0 : (x["totalBytes"] as number));
        A.store.Bool(ptr + 34, "organizationLimitExceeded" in x ? true : false);
        A.store.Bool(ptr + 24, x["organizationLimitExceeded"] ? true : false);
        A.store.Ref(ptr + 28, x["organizationName"]);
      }
    },
    "load_DriveQuotaMetadata": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["userType"] = A.load.Enum(ptr + 0, ["unmanaged", "organization"]);
      if (A.load.Bool(ptr + 32)) {
        x["usedBytes"] = A.load.Float64(ptr + 8);
      } else {
        delete x["usedBytes"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["totalBytes"] = A.load.Float64(ptr + 16);
      } else {
        delete x["totalBytes"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["organizationLimitExceeded"] = A.load.Bool(ptr + 24);
      } else {
        delete x["organizationLimitExceeded"];
      }
      x["organizationName"] = A.load.Ref(ptr + 28, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DriveShareType": (ref: heap.Ref<string>): number => {
      const idx = ["can_edit", "can_comment", "can_view"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_DriveSyncErrorType": (ref: heap.Ref<string>): number => {
      const idx = [
        "delete_without_permission",
        "service_unavailable",
        "no_server_space",
        "no_server_space_organization",
        "no_local_space",
        "no_shared_drive_space",
        "misc",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DriveSyncErrorEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Enum(
          ptr + 0,
          [
            "delete_without_permission",
            "service_unavailable",
            "no_server_space",
            "no_server_space_organization",
            "no_local_space",
            "no_shared_drive_space",
            "misc",
          ].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 4, x["fileUrl"]);
        A.store.Ref(ptr + 8, x["sharedDrive"]);
      }
    },
    "load_DriveSyncErrorEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, [
        "delete_without_permission",
        "service_unavailable",
        "no_server_space",
        "no_server_space_organization",
        "no_local_space",
        "no_shared_drive_space",
        "misc",
      ]);
      x["fileUrl"] = A.load.Ref(ptr + 4, undefined);
      x["sharedDrive"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RecentDateBucket": (ref: heap.Ref<string>): number => {
      const idx = [
        "today",
        "yesterday",
        "earlier_this_week",
        "earlier_this_month",
        "earlier_this_year",
        "older",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SyncStatus": (ref: heap.Ref<string>): number => {
      const idx = ["not_found", "queued", "in_progress", "completed", "error"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_EntryProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 140, false);
        A.store.Bool(ptr + 113, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 114, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 115, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Enum(ptr + 24, -1);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Bool(ptr + 116, false);
        A.store.Int32(ptr + 36, 0);
        A.store.Bool(ptr + 117, false);
        A.store.Int32(ptr + 40, 0);
        A.store.Bool(ptr + 118, false);
        A.store.Int32(ptr + 44, 0);
        A.store.Bool(ptr + 119, false);
        A.store.Bool(ptr + 48, false);
        A.store.Bool(ptr + 120, false);
        A.store.Bool(ptr + 49, false);
        A.store.Bool(ptr + 121, false);
        A.store.Bool(ptr + 50, false);
        A.store.Bool(ptr + 122, false);
        A.store.Bool(ptr + 51, false);
        A.store.Bool(ptr + 123, false);
        A.store.Bool(ptr + 52, false);
        A.store.Bool(ptr + 124, false);
        A.store.Bool(ptr + 53, false);
        A.store.Ref(ptr + 56, undefined);
        A.store.Ref(ptr + 60, undefined);
        A.store.Bool(ptr + 125, false);
        A.store.Bool(ptr + 64, false);
        A.store.Bool(ptr + 126, false);
        A.store.Bool(ptr + 65, false);
        A.store.Bool(ptr + 127, false);
        A.store.Bool(ptr + 66, false);
        A.store.Ref(ptr + 68, undefined);
        A.store.Ref(ptr + 72, undefined);
        A.store.Ref(ptr + 76, undefined);
        A.store.Bool(ptr + 128, false);
        A.store.Bool(ptr + 80, false);
        A.store.Bool(ptr + 129, false);
        A.store.Bool(ptr + 81, false);
        A.store.Bool(ptr + 130, false);
        A.store.Bool(ptr + 82, false);
        A.store.Bool(ptr + 131, false);
        A.store.Bool(ptr + 83, false);
        A.store.Bool(ptr + 132, false);
        A.store.Bool(ptr + 84, false);
        A.store.Bool(ptr + 133, false);
        A.store.Bool(ptr + 85, false);
        A.store.Bool(ptr + 134, false);
        A.store.Bool(ptr + 86, false);
        A.store.Bool(ptr + 135, false);
        A.store.Bool(ptr + 87, false);
        A.store.Bool(ptr + 136, false);
        A.store.Bool(ptr + 88, false);
        A.store.Enum(ptr + 92, -1);
        A.store.Bool(ptr + 137, false);
        A.store.Float64(ptr + 96, 0);
        A.store.Bool(ptr + 138, false);
        A.store.Float64(ptr + 104, 0);
        A.store.Bool(ptr + 139, false);
        A.store.Bool(ptr + 112, false);
      } else {
        A.store.Bool(ptr + 140, true);
        A.store.Bool(ptr + 113, "size" in x ? true : false);
        A.store.Float64(ptr + 0, x["size"] === undefined ? 0 : (x["size"] as number));
        A.store.Bool(ptr + 114, "modificationTime" in x ? true : false);
        A.store.Float64(ptr + 8, x["modificationTime"] === undefined ? 0 : (x["modificationTime"] as number));
        A.store.Bool(ptr + 115, "modificationByMeTime" in x ? true : false);
        A.store.Float64(ptr + 16, x["modificationByMeTime"] === undefined ? 0 : (x["modificationByMeTime"] as number));
        A.store.Enum(
          ptr + 24,
          ["today", "yesterday", "earlier_this_week", "earlier_this_month", "earlier_this_year", "older"].indexOf(
            x["recentDateBucket"] as string
          )
        );
        A.store.Ref(ptr + 28, x["thumbnailUrl"]);
        A.store.Ref(ptr + 32, x["croppedThumbnailUrl"]);
        A.store.Bool(ptr + 116, "imageWidth" in x ? true : false);
        A.store.Int32(ptr + 36, x["imageWidth"] === undefined ? 0 : (x["imageWidth"] as number));
        A.store.Bool(ptr + 117, "imageHeight" in x ? true : false);
        A.store.Int32(ptr + 40, x["imageHeight"] === undefined ? 0 : (x["imageHeight"] as number));
        A.store.Bool(ptr + 118, "imageRotation" in x ? true : false);
        A.store.Int32(ptr + 44, x["imageRotation"] === undefined ? 0 : (x["imageRotation"] as number));
        A.store.Bool(ptr + 119, "pinned" in x ? true : false);
        A.store.Bool(ptr + 48, x["pinned"] ? true : false);
        A.store.Bool(ptr + 120, "present" in x ? true : false);
        A.store.Bool(ptr + 49, x["present"] ? true : false);
        A.store.Bool(ptr + 121, "hosted" in x ? true : false);
        A.store.Bool(ptr + 50, x["hosted"] ? true : false);
        A.store.Bool(ptr + 122, "availableOffline" in x ? true : false);
        A.store.Bool(ptr + 51, x["availableOffline"] ? true : false);
        A.store.Bool(ptr + 123, "availableWhenMetered" in x ? true : false);
        A.store.Bool(ptr + 52, x["availableWhenMetered"] ? true : false);
        A.store.Bool(ptr + 124, "dirty" in x ? true : false);
        A.store.Bool(ptr + 53, x["dirty"] ? true : false);
        A.store.Ref(ptr + 56, x["customIconUrl"]);
        A.store.Ref(ptr + 60, x["contentMimeType"]);
        A.store.Bool(ptr + 125, "sharedWithMe" in x ? true : false);
        A.store.Bool(ptr + 64, x["sharedWithMe"] ? true : false);
        A.store.Bool(ptr + 126, "shared" in x ? true : false);
        A.store.Bool(ptr + 65, x["shared"] ? true : false);
        A.store.Bool(ptr + 127, "starred" in x ? true : false);
        A.store.Bool(ptr + 66, x["starred"] ? true : false);
        A.store.Ref(ptr + 68, x["externalFileUrl"]);
        A.store.Ref(ptr + 72, x["alternateUrl"]);
        A.store.Ref(ptr + 76, x["shareUrl"]);
        A.store.Bool(ptr + 128, "canCopy" in x ? true : false);
        A.store.Bool(ptr + 80, x["canCopy"] ? true : false);
        A.store.Bool(ptr + 129, "canDelete" in x ? true : false);
        A.store.Bool(ptr + 81, x["canDelete"] ? true : false);
        A.store.Bool(ptr + 130, "canRename" in x ? true : false);
        A.store.Bool(ptr + 82, x["canRename"] ? true : false);
        A.store.Bool(ptr + 131, "canAddChildren" in x ? true : false);
        A.store.Bool(ptr + 83, x["canAddChildren"] ? true : false);
        A.store.Bool(ptr + 132, "canShare" in x ? true : false);
        A.store.Bool(ptr + 84, x["canShare"] ? true : false);
        A.store.Bool(ptr + 133, "canPin" in x ? true : false);
        A.store.Bool(ptr + 85, x["canPin"] ? true : false);
        A.store.Bool(ptr + 134, "isMachineRoot" in x ? true : false);
        A.store.Bool(ptr + 86, x["isMachineRoot"] ? true : false);
        A.store.Bool(ptr + 135, "isExternalMedia" in x ? true : false);
        A.store.Bool(ptr + 87, x["isExternalMedia"] ? true : false);
        A.store.Bool(ptr + 136, "isArbitrarySyncFolder" in x ? true : false);
        A.store.Bool(ptr + 88, x["isArbitrarySyncFolder"] ? true : false);
        A.store.Enum(
          ptr + 92,
          ["not_found", "queued", "in_progress", "completed", "error"].indexOf(x["syncStatus"] as string)
        );
        A.store.Bool(ptr + 137, "progress" in x ? true : false);
        A.store.Float64(ptr + 96, x["progress"] === undefined ? 0 : (x["progress"] as number));
        A.store.Bool(ptr + 138, "syncCompletedTime" in x ? true : false);
        A.store.Float64(ptr + 104, x["syncCompletedTime"] === undefined ? 0 : (x["syncCompletedTime"] as number));
        A.store.Bool(ptr + 139, "shortcut" in x ? true : false);
        A.store.Bool(ptr + 112, x["shortcut"] ? true : false);
      }
    },
    "load_EntryProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 113)) {
        x["size"] = A.load.Float64(ptr + 0);
      } else {
        delete x["size"];
      }
      if (A.load.Bool(ptr + 114)) {
        x["modificationTime"] = A.load.Float64(ptr + 8);
      } else {
        delete x["modificationTime"];
      }
      if (A.load.Bool(ptr + 115)) {
        x["modificationByMeTime"] = A.load.Float64(ptr + 16);
      } else {
        delete x["modificationByMeTime"];
      }
      x["recentDateBucket"] = A.load.Enum(ptr + 24, [
        "today",
        "yesterday",
        "earlier_this_week",
        "earlier_this_month",
        "earlier_this_year",
        "older",
      ]);
      x["thumbnailUrl"] = A.load.Ref(ptr + 28, undefined);
      x["croppedThumbnailUrl"] = A.load.Ref(ptr + 32, undefined);
      if (A.load.Bool(ptr + 116)) {
        x["imageWidth"] = A.load.Int32(ptr + 36);
      } else {
        delete x["imageWidth"];
      }
      if (A.load.Bool(ptr + 117)) {
        x["imageHeight"] = A.load.Int32(ptr + 40);
      } else {
        delete x["imageHeight"];
      }
      if (A.load.Bool(ptr + 118)) {
        x["imageRotation"] = A.load.Int32(ptr + 44);
      } else {
        delete x["imageRotation"];
      }
      if (A.load.Bool(ptr + 119)) {
        x["pinned"] = A.load.Bool(ptr + 48);
      } else {
        delete x["pinned"];
      }
      if (A.load.Bool(ptr + 120)) {
        x["present"] = A.load.Bool(ptr + 49);
      } else {
        delete x["present"];
      }
      if (A.load.Bool(ptr + 121)) {
        x["hosted"] = A.load.Bool(ptr + 50);
      } else {
        delete x["hosted"];
      }
      if (A.load.Bool(ptr + 122)) {
        x["availableOffline"] = A.load.Bool(ptr + 51);
      } else {
        delete x["availableOffline"];
      }
      if (A.load.Bool(ptr + 123)) {
        x["availableWhenMetered"] = A.load.Bool(ptr + 52);
      } else {
        delete x["availableWhenMetered"];
      }
      if (A.load.Bool(ptr + 124)) {
        x["dirty"] = A.load.Bool(ptr + 53);
      } else {
        delete x["dirty"];
      }
      x["customIconUrl"] = A.load.Ref(ptr + 56, undefined);
      x["contentMimeType"] = A.load.Ref(ptr + 60, undefined);
      if (A.load.Bool(ptr + 125)) {
        x["sharedWithMe"] = A.load.Bool(ptr + 64);
      } else {
        delete x["sharedWithMe"];
      }
      if (A.load.Bool(ptr + 126)) {
        x["shared"] = A.load.Bool(ptr + 65);
      } else {
        delete x["shared"];
      }
      if (A.load.Bool(ptr + 127)) {
        x["starred"] = A.load.Bool(ptr + 66);
      } else {
        delete x["starred"];
      }
      x["externalFileUrl"] = A.load.Ref(ptr + 68, undefined);
      x["alternateUrl"] = A.load.Ref(ptr + 72, undefined);
      x["shareUrl"] = A.load.Ref(ptr + 76, undefined);
      if (A.load.Bool(ptr + 128)) {
        x["canCopy"] = A.load.Bool(ptr + 80);
      } else {
        delete x["canCopy"];
      }
      if (A.load.Bool(ptr + 129)) {
        x["canDelete"] = A.load.Bool(ptr + 81);
      } else {
        delete x["canDelete"];
      }
      if (A.load.Bool(ptr + 130)) {
        x["canRename"] = A.load.Bool(ptr + 82);
      } else {
        delete x["canRename"];
      }
      if (A.load.Bool(ptr + 131)) {
        x["canAddChildren"] = A.load.Bool(ptr + 83);
      } else {
        delete x["canAddChildren"];
      }
      if (A.load.Bool(ptr + 132)) {
        x["canShare"] = A.load.Bool(ptr + 84);
      } else {
        delete x["canShare"];
      }
      if (A.load.Bool(ptr + 133)) {
        x["canPin"] = A.load.Bool(ptr + 85);
      } else {
        delete x["canPin"];
      }
      if (A.load.Bool(ptr + 134)) {
        x["isMachineRoot"] = A.load.Bool(ptr + 86);
      } else {
        delete x["isMachineRoot"];
      }
      if (A.load.Bool(ptr + 135)) {
        x["isExternalMedia"] = A.load.Bool(ptr + 87);
      } else {
        delete x["isExternalMedia"];
      }
      if (A.load.Bool(ptr + 136)) {
        x["isArbitrarySyncFolder"] = A.load.Bool(ptr + 88);
      } else {
        delete x["isArbitrarySyncFolder"];
      }
      x["syncStatus"] = A.load.Enum(ptr + 92, ["not_found", "queued", "in_progress", "completed", "error"]);
      if (A.load.Bool(ptr + 137)) {
        x["progress"] = A.load.Float64(ptr + 96);
      } else {
        delete x["progress"];
      }
      if (A.load.Bool(ptr + 138)) {
        x["syncCompletedTime"] = A.load.Float64(ptr + 104);
      } else {
        delete x["syncCompletedTime"];
      }
      if (A.load.Bool(ptr + 139)) {
        x["shortcut"] = A.load.Bool(ptr + 112);
      } else {
        delete x["shortcut"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_EntryPropertyName": (ref: heap.Ref<string>): number => {
      const idx = [
        "size",
        "modificationTime",
        "modificationByMeTime",
        "thumbnailUrl",
        "croppedThumbnailUrl",
        "imageWidth",
        "imageHeight",
        "imageRotation",
        "pinned",
        "present",
        "hosted",
        "availableOffline",
        "availableWhenMetered",
        "dirty",
        "customIconUrl",
        "contentMimeType",
        "sharedWithMe",
        "shared",
        "starred",
        "externalFileUrl",
        "alternateUrl",
        "shareUrl",
        "canCopy",
        "canDelete",
        "canRename",
        "canAddChildren",
        "canShare",
        "canPin",
        "isMachineRoot",
        "isExternalMedia",
        "isArbitrarySyncFolder",
        "syncStatus",
        "progress",
        "shortcut",
        "syncCompletedTime",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_TaskResult": (ref: heap.Ref<string>): number => {
      const idx = ["opened", "message_sent", "failed", "empty", "failed_plugin_vm_directory_not_shared"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_FileCategory": (ref: heap.Ref<string>): number => {
      const idx = ["all", "audio", "image", "video", "document"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_FileChange": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["url"]);
        A.store.Ref(ptr + 4, x["changes"]);
      }
    },
    "load_FileChange": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      x["changes"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_FileSystemProviderAction": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["title"]);
      }
    },
    "load_FileSystemProviderAction": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["title"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_FileTaskDescriptor": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["appId"]);
        A.store.Ref(ptr + 4, x["taskType"]);
        A.store.Ref(ptr + 8, x["actionId"]);
      }
    },
    "load_FileTaskDescriptor": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["appId"] = A.load.Ref(ptr + 0, undefined);
      x["taskType"] = A.load.Ref(ptr + 4, undefined);
      x["actionId"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_FileTask": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 30, false);

        A.store.Bool(ptr + 0 + 12, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Ref(ptr + 0 + 8, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Bool(ptr + 27, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 29, false);
        A.store.Bool(ptr + 26, false);
      } else {
        A.store.Bool(ptr + 30, true);

        if (typeof x["descriptor"] === "undefined") {
          A.store.Bool(ptr + 0 + 12, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Ref(ptr + 0 + 8, undefined);
        } else {
          A.store.Bool(ptr + 0 + 12, true);
          A.store.Ref(ptr + 0 + 0, x["descriptor"]["appId"]);
          A.store.Ref(ptr + 0 + 4, x["descriptor"]["taskType"]);
          A.store.Ref(ptr + 0 + 8, x["descriptor"]["actionId"]);
        }
        A.store.Ref(ptr + 16, x["title"]);
        A.store.Ref(ptr + 20, x["iconUrl"]);
        A.store.Bool(ptr + 27, "isDefault" in x ? true : false);
        A.store.Bool(ptr + 24, x["isDefault"] ? true : false);
        A.store.Bool(ptr + 28, "isGenericFileHandler" in x ? true : false);
        A.store.Bool(ptr + 25, x["isGenericFileHandler"] ? true : false);
        A.store.Bool(ptr + 29, "isDlpBlocked" in x ? true : false);
        A.store.Bool(ptr + 26, x["isDlpBlocked"] ? true : false);
      }
    },
    "load_FileTask": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 12)) {
        x["descriptor"] = {};
        x["descriptor"]["appId"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["descriptor"]["taskType"] = A.load.Ref(ptr + 0 + 4, undefined);
        x["descriptor"]["actionId"] = A.load.Ref(ptr + 0 + 8, undefined);
      } else {
        delete x["descriptor"];
      }
      x["title"] = A.load.Ref(ptr + 16, undefined);
      x["iconUrl"] = A.load.Ref(ptr + 20, undefined);
      if (A.load.Bool(ptr + 27)) {
        x["isDefault"] = A.load.Bool(ptr + 24);
      } else {
        delete x["isDefault"];
      }
      if (A.load.Bool(ptr + 28)) {
        x["isGenericFileHandler"] = A.load.Bool(ptr + 25);
      } else {
        delete x["isGenericFileHandler"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["isDlpBlocked"] = A.load.Bool(ptr + 26);
      } else {
        delete x["isDlpBlocked"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_TransferState": (ref: heap.Ref<string>): number => {
      const idx = ["in_progress", "queued", "completed", "failed"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_FileTransferStatus": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 35, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Bool(ptr + 30, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 31, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 32, false);
        A.store.Int32(ptr + 24, 0);
        A.store.Bool(ptr + 33, false);
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 34, false);
        A.store.Bool(ptr + 29, false);
      } else {
        A.store.Bool(ptr + 35, true);
        A.store.Ref(ptr + 0, x["fileUrl"]);
        A.store.Enum(ptr + 4, ["in_progress", "queued", "completed", "failed"].indexOf(x["transferState"] as string));
        A.store.Bool(ptr + 30, "processed" in x ? true : false);
        A.store.Float64(ptr + 8, x["processed"] === undefined ? 0 : (x["processed"] as number));
        A.store.Bool(ptr + 31, "total" in x ? true : false);
        A.store.Float64(ptr + 16, x["total"] === undefined ? 0 : (x["total"] as number));
        A.store.Bool(ptr + 32, "numTotalJobs" in x ? true : false);
        A.store.Int32(ptr + 24, x["numTotalJobs"] === undefined ? 0 : (x["numTotalJobs"] as number));
        A.store.Bool(ptr + 33, "showNotification" in x ? true : false);
        A.store.Bool(ptr + 28, x["showNotification"] ? true : false);
        A.store.Bool(ptr + 34, "hideWhenZeroJobs" in x ? true : false);
        A.store.Bool(ptr + 29, x["hideWhenZeroJobs"] ? true : false);
      }
    },
    "load_FileTransferStatus": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileUrl"] = A.load.Ref(ptr + 0, undefined);
      x["transferState"] = A.load.Enum(ptr + 4, ["in_progress", "queued", "completed", "failed"]);
      if (A.load.Bool(ptr + 30)) {
        x["processed"] = A.load.Float64(ptr + 8);
      } else {
        delete x["processed"];
      }
      if (A.load.Bool(ptr + 31)) {
        x["total"] = A.load.Float64(ptr + 16);
      } else {
        delete x["total"];
      }
      if (A.load.Bool(ptr + 32)) {
        x["numTotalJobs"] = A.load.Int32(ptr + 24);
      } else {
        delete x["numTotalJobs"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["showNotification"] = A.load.Bool(ptr + 28);
      } else {
        delete x["showNotification"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["hideWhenZeroJobs"] = A.load.Bool(ptr + 29);
      } else {
        delete x["hideWhenZeroJobs"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_FileWatchEventType": (ref: heap.Ref<string>): number => {
      const idx = ["changed", "error"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_FileWatchEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Enum(ptr + 0, ["changed", "error"].indexOf(x["eventType"] as string));
        A.store.Ref(ptr + 4, x["entry"]);
        A.store.Ref(ptr + 8, x["changedFiles"]);
      }
    },
    "load_FileWatchEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["eventType"] = A.load.Enum(ptr + 0, ["changed", "error"]);
      x["entry"] = A.load.Ref(ptr + 4, undefined);
      x["changedFiles"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_FormatFileSystemType": (ref: heap.Ref<string>): number => {
      const idx = ["vfat", "exfat", "ntfs"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_StreamInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["type"]);
        A.store.Ref(ptr + 4, x["tags"]);
      }
    },
    "load_StreamInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Ref(ptr + 0, undefined);
      x["tags"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MediaMetadata": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 78, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 72, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 73, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 74, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 75, false);
        A.store.Int32(ptr + 24, 0);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
        A.store.Bool(ptr + 76, false);
        A.store.Int32(ptr + 44, 0);
        A.store.Ref(ptr + 48, undefined);
        A.store.Ref(ptr + 52, undefined);
        A.store.Ref(ptr + 56, undefined);
        A.store.Bool(ptr + 77, false);
        A.store.Int32(ptr + 60, 0);
        A.store.Ref(ptr + 64, undefined);
        A.store.Ref(ptr + 68, undefined);
      } else {
        A.store.Bool(ptr + 78, true);
        A.store.Ref(ptr + 0, x["mimeType"]);
        A.store.Bool(ptr + 72, "height" in x ? true : false);
        A.store.Int32(ptr + 4, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Bool(ptr + 73, "width" in x ? true : false);
        A.store.Int32(ptr + 8, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Bool(ptr + 74, "duration" in x ? true : false);
        A.store.Float64(ptr + 16, x["duration"] === undefined ? 0 : (x["duration"] as number));
        A.store.Bool(ptr + 75, "rotation" in x ? true : false);
        A.store.Int32(ptr + 24, x["rotation"] === undefined ? 0 : (x["rotation"] as number));
        A.store.Ref(ptr + 28, x["album"]);
        A.store.Ref(ptr + 32, x["artist"]);
        A.store.Ref(ptr + 36, x["comment"]);
        A.store.Ref(ptr + 40, x["copyright"]);
        A.store.Bool(ptr + 76, "disc" in x ? true : false);
        A.store.Int32(ptr + 44, x["disc"] === undefined ? 0 : (x["disc"] as number));
        A.store.Ref(ptr + 48, x["genre"]);
        A.store.Ref(ptr + 52, x["language"]);
        A.store.Ref(ptr + 56, x["title"]);
        A.store.Bool(ptr + 77, "track" in x ? true : false);
        A.store.Int32(ptr + 60, x["track"] === undefined ? 0 : (x["track"] as number));
        A.store.Ref(ptr + 64, x["rawTags"]);
        A.store.Ref(ptr + 68, x["attachedImages"]);
      }
    },
    "load_MediaMetadata": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["mimeType"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 72)) {
        x["height"] = A.load.Int32(ptr + 4);
      } else {
        delete x["height"];
      }
      if (A.load.Bool(ptr + 73)) {
        x["width"] = A.load.Int32(ptr + 8);
      } else {
        delete x["width"];
      }
      if (A.load.Bool(ptr + 74)) {
        x["duration"] = A.load.Float64(ptr + 16);
      } else {
        delete x["duration"];
      }
      if (A.load.Bool(ptr + 75)) {
        x["rotation"] = A.load.Int32(ptr + 24);
      } else {
        delete x["rotation"];
      }
      x["album"] = A.load.Ref(ptr + 28, undefined);
      x["artist"] = A.load.Ref(ptr + 32, undefined);
      x["comment"] = A.load.Ref(ptr + 36, undefined);
      x["copyright"] = A.load.Ref(ptr + 40, undefined);
      if (A.load.Bool(ptr + 76)) {
        x["disc"] = A.load.Int32(ptr + 44);
      } else {
        delete x["disc"];
      }
      x["genre"] = A.load.Ref(ptr + 48, undefined);
      x["language"] = A.load.Ref(ptr + 52, undefined);
      x["title"] = A.load.Ref(ptr + 56, undefined);
      if (A.load.Bool(ptr + 77)) {
        x["track"] = A.load.Int32(ptr + 60);
      } else {
        delete x["track"];
      }
      x["rawTags"] = A.load.Ref(ptr + 64, undefined);
      x["attachedImages"] = A.load.Ref(ptr + 68, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PolicyDefaultHandlerStatus": (ref: heap.Ref<string>): number => {
      const idx = ["default_handler_assigned_by_policy", "incorrect_assignment"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ResultingTasks": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["tasks"]);
        A.store.Enum(
          ptr + 4,
          ["default_handler_assigned_by_policy", "incorrect_assignment"].indexOf(
            x["policyDefaultHandlerStatus"] as string
          )
        );
      }
    },
    "load_ResultingTasks": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["tasks"] = A.load.Ref(ptr + 0, undefined);
      x["policyDefaultHandlerStatus"] = A.load.Enum(ptr + 4, [
        "default_handler_assigned_by_policy",
        "incorrect_assignment",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_LinuxPackageInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["version"]);
        A.store.Ref(ptr + 8, x["summary"]);
        A.store.Ref(ptr + 12, x["description"]);
      }
    },
    "load_LinuxPackageInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["version"] = A.load.Ref(ptr + 4, undefined);
      x["summary"] = A.load.Ref(ptr + 8, undefined);
      x["description"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Preferences": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 51, false);
        A.store.Bool(ptr + 41, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 42, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 43, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 44, false);
        A.store.Bool(ptr + 3, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 45, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 46, false);
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 47, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 48, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Bool(ptr + 49, false);
        A.store.Float64(ptr + 32, 0);
        A.store.Bool(ptr + 50, false);
        A.store.Bool(ptr + 40, false);
      } else {
        A.store.Bool(ptr + 51, true);
        A.store.Bool(ptr + 41, "driveEnabled" in x ? true : false);
        A.store.Bool(ptr + 0, x["driveEnabled"] ? true : false);
        A.store.Bool(ptr + 42, "driveSyncEnabledOnMeteredNetwork" in x ? true : false);
        A.store.Bool(ptr + 1, x["driveSyncEnabledOnMeteredNetwork"] ? true : false);
        A.store.Bool(ptr + 43, "searchSuggestEnabled" in x ? true : false);
        A.store.Bool(ptr + 2, x["searchSuggestEnabled"] ? true : false);
        A.store.Bool(ptr + 44, "use24hourClock" in x ? true : false);
        A.store.Bool(ptr + 3, x["use24hourClock"] ? true : false);
        A.store.Ref(ptr + 4, x["timezone"]);
        A.store.Bool(ptr + 45, "arcEnabled" in x ? true : false);
        A.store.Bool(ptr + 8, x["arcEnabled"] ? true : false);
        A.store.Bool(ptr + 46, "arcRemovableMediaAccessEnabled" in x ? true : false);
        A.store.Bool(ptr + 9, x["arcRemovableMediaAccessEnabled"] ? true : false);
        A.store.Ref(ptr + 12, x["folderShortcuts"]);
        A.store.Bool(ptr + 47, "trashEnabled" in x ? true : false);
        A.store.Bool(ptr + 16, x["trashEnabled"] ? true : false);
        A.store.Bool(ptr + 48, "officeFileMovedOneDrive" in x ? true : false);
        A.store.Float64(
          ptr + 24,
          x["officeFileMovedOneDrive"] === undefined ? 0 : (x["officeFileMovedOneDrive"] as number)
        );
        A.store.Bool(ptr + 49, "officeFileMovedGoogleDrive" in x ? true : false);
        A.store.Float64(
          ptr + 32,
          x["officeFileMovedGoogleDrive"] === undefined ? 0 : (x["officeFileMovedGoogleDrive"] as number)
        );
        A.store.Bool(ptr + 50, "driveFsBulkPinningEnabled" in x ? true : false);
        A.store.Bool(ptr + 40, x["driveFsBulkPinningEnabled"] ? true : false);
      }
    },
    "load_Preferences": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 41)) {
        x["driveEnabled"] = A.load.Bool(ptr + 0);
      } else {
        delete x["driveEnabled"];
      }
      if (A.load.Bool(ptr + 42)) {
        x["driveSyncEnabledOnMeteredNetwork"] = A.load.Bool(ptr + 1);
      } else {
        delete x["driveSyncEnabledOnMeteredNetwork"];
      }
      if (A.load.Bool(ptr + 43)) {
        x["searchSuggestEnabled"] = A.load.Bool(ptr + 2);
      } else {
        delete x["searchSuggestEnabled"];
      }
      if (A.load.Bool(ptr + 44)) {
        x["use24hourClock"] = A.load.Bool(ptr + 3);
      } else {
        delete x["use24hourClock"];
      }
      x["timezone"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 45)) {
        x["arcEnabled"] = A.load.Bool(ptr + 8);
      } else {
        delete x["arcEnabled"];
      }
      if (A.load.Bool(ptr + 46)) {
        x["arcRemovableMediaAccessEnabled"] = A.load.Bool(ptr + 9);
      } else {
        delete x["arcRemovableMediaAccessEnabled"];
      }
      x["folderShortcuts"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 47)) {
        x["trashEnabled"] = A.load.Bool(ptr + 16);
      } else {
        delete x["trashEnabled"];
      }
      if (A.load.Bool(ptr + 48)) {
        x["officeFileMovedOneDrive"] = A.load.Float64(ptr + 24);
      } else {
        delete x["officeFileMovedOneDrive"];
      }
      if (A.load.Bool(ptr + 49)) {
        x["officeFileMovedGoogleDrive"] = A.load.Float64(ptr + 32);
      } else {
        delete x["officeFileMovedGoogleDrive"];
      }
      if (A.load.Bool(ptr + 50)) {
        x["driveFsBulkPinningEnabled"] = A.load.Bool(ptr + 40);
      } else {
        delete x["driveFsBulkPinningEnabled"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ProfileInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Ref(ptr + 0, x["profileId"]);
        A.store.Ref(ptr + 4, x["displayName"]);
        A.store.Bool(ptr + 9, "isCurrentProfile" in x ? true : false);
        A.store.Bool(ptr + 8, x["isCurrentProfile"] ? true : false);
      }
    },
    "load_ProfileInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["profileId"] = A.load.Ref(ptr + 0, undefined);
      x["displayName"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 9)) {
        x["isCurrentProfile"] = A.load.Bool(ptr + 8);
      } else {
        delete x["isCurrentProfile"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ProviderSource": (ref: heap.Ref<string>): number => {
      const idx = ["file", "device", "network"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Provider": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 31, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 4 + 8, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Ref(ptr + 4 + 4, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 29, false);
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 30, false);
        A.store.Bool(ptr + 22, false);
        A.store.Enum(ptr + 24, -1);
      } else {
        A.store.Bool(ptr + 31, true);
        A.store.Ref(ptr + 0, x["providerId"]);

        if (typeof x["iconSet"] === "undefined") {
          A.store.Bool(ptr + 4 + 8, false);
          A.store.Ref(ptr + 4 + 0, undefined);
          A.store.Ref(ptr + 4 + 4, undefined);
        } else {
          A.store.Bool(ptr + 4 + 8, true);
          A.store.Ref(ptr + 4 + 0, x["iconSet"]["icon16x16Url"]);
          A.store.Ref(ptr + 4 + 4, x["iconSet"]["icon32x32Url"]);
        }
        A.store.Ref(ptr + 16, x["name"]);
        A.store.Bool(ptr + 28, "configurable" in x ? true : false);
        A.store.Bool(ptr + 20, x["configurable"] ? true : false);
        A.store.Bool(ptr + 29, "watchable" in x ? true : false);
        A.store.Bool(ptr + 21, x["watchable"] ? true : false);
        A.store.Bool(ptr + 30, "multipleMounts" in x ? true : false);
        A.store.Bool(ptr + 22, x["multipleMounts"] ? true : false);
        A.store.Enum(ptr + 24, ["file", "device", "network"].indexOf(x["source"] as string));
      }
    },
    "load_Provider": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["providerId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 4 + 8)) {
        x["iconSet"] = {};
        x["iconSet"]["icon16x16Url"] = A.load.Ref(ptr + 4 + 0, undefined);
        x["iconSet"]["icon32x32Url"] = A.load.Ref(ptr + 4 + 4, undefined);
      } else {
        delete x["iconSet"];
      }
      x["name"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 28)) {
        x["configurable"] = A.load.Bool(ptr + 20);
      } else {
        delete x["configurable"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["watchable"] = A.load.Bool(ptr + 21);
      } else {
        delete x["watchable"];
      }
      if (A.load.Bool(ptr + 30)) {
        x["multipleMounts"] = A.load.Bool(ptr + 22);
      } else {
        delete x["multipleMounts"];
      }
      x["source"] = A.load.Enum(ptr + 24, ["file", "device", "network"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MountPointSizeStats": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Bool(ptr + 16, "totalSize" in x ? true : false);
        A.store.Float64(ptr + 0, x["totalSize"] === undefined ? 0 : (x["totalSize"] as number));
        A.store.Bool(ptr + 17, "remainingSize" in x ? true : false);
        A.store.Float64(ptr + 8, x["remainingSize"] === undefined ? 0 : (x["remainingSize"] as number));
      }
    },
    "load_MountPointSizeStats": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["totalSize"] = A.load.Float64(ptr + 0);
      } else {
        delete x["totalSize"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["remainingSize"] = A.load.Float64(ptr + 8);
      } else {
        delete x["remainingSize"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_Source": (ref: heap.Ref<string>): number => {
      const idx = ["file", "device", "network", "system"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_MountError": (ref: heap.Ref<string>): number => {
      const idx = [
        "success",
        "in_progress",
        "unknown_error",
        "internal_error",
        "invalid_argument",
        "invalid_path",
        "path_already_mounted",
        "path_not_mounted",
        "directory_creation_failed",
        "invalid_mount_options",
        "insufficient_permissions",
        "mount_program_not_found",
        "mount_program_failed",
        "invalid_device_path",
        "unknown_filesystem",
        "unsupported_filesystem",
        "need_password",
        "cancelled",
        "busy",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_MountContext": (ref: heap.Ref<string>): number => {
      const idx = ["user", "auto"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_VmType": (ref: heap.Ref<string>): number => {
      const idx = ["termina", "plugin_vm", "borealis", "bruschetta", "arcvm"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_VolumeMetadata": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 103, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Enum(ptr + 12, -1);
        A.store.Ref(ptr + 16, undefined);

        A.store.Bool(ptr + 20 + 10, false);
        A.store.Ref(ptr + 20 + 0, undefined);
        A.store.Ref(ptr + 20 + 4, undefined);
        A.store.Bool(ptr + 20 + 9, false);
        A.store.Bool(ptr + 20 + 8, false);
        A.store.Ref(ptr + 32, undefined);
        A.store.Enum(ptr + 36, -1);
        A.store.Enum(ptr + 40, -1);
        A.store.Ref(ptr + 44, undefined);
        A.store.Bool(ptr + 96, false);
        A.store.Bool(ptr + 48, false);
        A.store.Bool(ptr + 97, false);
        A.store.Bool(ptr + 49, false);
        A.store.Bool(ptr + 98, false);
        A.store.Bool(ptr + 50, false);
        A.store.Bool(ptr + 99, false);
        A.store.Bool(ptr + 51, false);
        A.store.Bool(ptr + 100, false);
        A.store.Bool(ptr + 52, false);
        A.store.Bool(ptr + 101, false);
        A.store.Bool(ptr + 53, false);
        A.store.Enum(ptr + 56, -1);
        A.store.Enum(ptr + 60, -1);
        A.store.Ref(ptr + 64, undefined);

        A.store.Bool(ptr + 68 + 8, false);
        A.store.Ref(ptr + 68 + 0, undefined);
        A.store.Ref(ptr + 68 + 4, undefined);
        A.store.Ref(ptr + 80, undefined);
        A.store.Ref(ptr + 84, undefined);
        A.store.Bool(ptr + 102, false);
        A.store.Bool(ptr + 88, false);
        A.store.Enum(ptr + 92, -1);
      } else {
        A.store.Bool(ptr + 103, true);
        A.store.Ref(ptr + 0, x["volumeId"]);
        A.store.Ref(ptr + 4, x["fileSystemId"]);
        A.store.Ref(ptr + 8, x["providerId"]);
        A.store.Enum(ptr + 12, ["file", "device", "network", "system"].indexOf(x["source"] as string));
        A.store.Ref(ptr + 16, x["volumeLabel"]);

        if (typeof x["profile"] === "undefined") {
          A.store.Bool(ptr + 20 + 10, false);
          A.store.Ref(ptr + 20 + 0, undefined);
          A.store.Ref(ptr + 20 + 4, undefined);
          A.store.Bool(ptr + 20 + 9, false);
          A.store.Bool(ptr + 20 + 8, false);
        } else {
          A.store.Bool(ptr + 20 + 10, true);
          A.store.Ref(ptr + 20 + 0, x["profile"]["profileId"]);
          A.store.Ref(ptr + 20 + 4, x["profile"]["displayName"]);
          A.store.Bool(ptr + 20 + 9, "isCurrentProfile" in x["profile"] ? true : false);
          A.store.Bool(ptr + 20 + 8, x["profile"]["isCurrentProfile"] ? true : false);
        }
        A.store.Ref(ptr + 32, x["sourcePath"]);
        A.store.Enum(
          ptr + 36,
          [
            "drive",
            "downloads",
            "removable",
            "archive",
            "provided",
            "mtp",
            "media_view",
            "crostini",
            "android_files",
            "documents_provider",
            "testing",
            "smb",
            "system_internal",
            "guest_os",
          ].indexOf(x["volumeType"] as string)
        );
        A.store.Enum(ptr + 40, ["usb", "sd", "optical", "mobile", "unknown"].indexOf(x["deviceType"] as string));
        A.store.Ref(ptr + 44, x["devicePath"]);
        A.store.Bool(ptr + 96, "isParentDevice" in x ? true : false);
        A.store.Bool(ptr + 48, x["isParentDevice"] ? true : false);
        A.store.Bool(ptr + 97, "isReadOnly" in x ? true : false);
        A.store.Bool(ptr + 49, x["isReadOnly"] ? true : false);
        A.store.Bool(ptr + 98, "isReadOnlyRemovableDevice" in x ? true : false);
        A.store.Bool(ptr + 50, x["isReadOnlyRemovableDevice"] ? true : false);
        A.store.Bool(ptr + 99, "hasMedia" in x ? true : false);
        A.store.Bool(ptr + 51, x["hasMedia"] ? true : false);
        A.store.Bool(ptr + 100, "configurable" in x ? true : false);
        A.store.Bool(ptr + 52, x["configurable"] ? true : false);
        A.store.Bool(ptr + 101, "watchable" in x ? true : false);
        A.store.Bool(ptr + 53, x["watchable"] ? true : false);
        A.store.Enum(
          ptr + 56,
          [
            "success",
            "in_progress",
            "unknown_error",
            "internal_error",
            "invalid_argument",
            "invalid_path",
            "path_already_mounted",
            "path_not_mounted",
            "directory_creation_failed",
            "invalid_mount_options",
            "insufficient_permissions",
            "mount_program_not_found",
            "mount_program_failed",
            "invalid_device_path",
            "unknown_filesystem",
            "unsupported_filesystem",
            "need_password",
            "cancelled",
            "busy",
          ].indexOf(x["mountCondition"] as string)
        );
        A.store.Enum(ptr + 60, ["user", "auto"].indexOf(x["mountContext"] as string));
        A.store.Ref(ptr + 64, x["diskFileSystemType"]);

        if (typeof x["iconSet"] === "undefined") {
          A.store.Bool(ptr + 68 + 8, false);
          A.store.Ref(ptr + 68 + 0, undefined);
          A.store.Ref(ptr + 68 + 4, undefined);
        } else {
          A.store.Bool(ptr + 68 + 8, true);
          A.store.Ref(ptr + 68 + 0, x["iconSet"]["icon16x16Url"]);
          A.store.Ref(ptr + 68 + 4, x["iconSet"]["icon32x32Url"]);
        }
        A.store.Ref(ptr + 80, x["driveLabel"]);
        A.store.Ref(ptr + 84, x["remoteMountPath"]);
        A.store.Bool(ptr + 102, "hidden" in x ? true : false);
        A.store.Bool(ptr + 88, x["hidden"] ? true : false);
        A.store.Enum(
          ptr + 92,
          ["termina", "plugin_vm", "borealis", "bruschetta", "arcvm"].indexOf(x["vmType"] as string)
        );
      }
    },
    "load_VolumeMetadata": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["volumeId"] = A.load.Ref(ptr + 0, undefined);
      x["fileSystemId"] = A.load.Ref(ptr + 4, undefined);
      x["providerId"] = A.load.Ref(ptr + 8, undefined);
      x["source"] = A.load.Enum(ptr + 12, ["file", "device", "network", "system"]);
      x["volumeLabel"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 20 + 10)) {
        x["profile"] = {};
        x["profile"]["profileId"] = A.load.Ref(ptr + 20 + 0, undefined);
        x["profile"]["displayName"] = A.load.Ref(ptr + 20 + 4, undefined);
        if (A.load.Bool(ptr + 20 + 9)) {
          x["profile"]["isCurrentProfile"] = A.load.Bool(ptr + 20 + 8);
        } else {
          delete x["profile"]["isCurrentProfile"];
        }
      } else {
        delete x["profile"];
      }
      x["sourcePath"] = A.load.Ref(ptr + 32, undefined);
      x["volumeType"] = A.load.Enum(ptr + 36, [
        "drive",
        "downloads",
        "removable",
        "archive",
        "provided",
        "mtp",
        "media_view",
        "crostini",
        "android_files",
        "documents_provider",
        "testing",
        "smb",
        "system_internal",
        "guest_os",
      ]);
      x["deviceType"] = A.load.Enum(ptr + 40, ["usb", "sd", "optical", "mobile", "unknown"]);
      x["devicePath"] = A.load.Ref(ptr + 44, undefined);
      if (A.load.Bool(ptr + 96)) {
        x["isParentDevice"] = A.load.Bool(ptr + 48);
      } else {
        delete x["isParentDevice"];
      }
      if (A.load.Bool(ptr + 97)) {
        x["isReadOnly"] = A.load.Bool(ptr + 49);
      } else {
        delete x["isReadOnly"];
      }
      if (A.load.Bool(ptr + 98)) {
        x["isReadOnlyRemovableDevice"] = A.load.Bool(ptr + 50);
      } else {
        delete x["isReadOnlyRemovableDevice"];
      }
      if (A.load.Bool(ptr + 99)) {
        x["hasMedia"] = A.load.Bool(ptr + 51);
      } else {
        delete x["hasMedia"];
      }
      if (A.load.Bool(ptr + 100)) {
        x["configurable"] = A.load.Bool(ptr + 52);
      } else {
        delete x["configurable"];
      }
      if (A.load.Bool(ptr + 101)) {
        x["watchable"] = A.load.Bool(ptr + 53);
      } else {
        delete x["watchable"];
      }
      x["mountCondition"] = A.load.Enum(ptr + 56, [
        "success",
        "in_progress",
        "unknown_error",
        "internal_error",
        "invalid_argument",
        "invalid_path",
        "path_already_mounted",
        "path_not_mounted",
        "directory_creation_failed",
        "invalid_mount_options",
        "insufficient_permissions",
        "mount_program_not_found",
        "mount_program_failed",
        "invalid_device_path",
        "unknown_filesystem",
        "unsupported_filesystem",
        "need_password",
        "cancelled",
        "busy",
      ]);
      x["mountContext"] = A.load.Enum(ptr + 60, ["user", "auto"]);
      x["diskFileSystemType"] = A.load.Ref(ptr + 64, undefined);
      if (A.load.Bool(ptr + 68 + 8)) {
        x["iconSet"] = {};
        x["iconSet"]["icon16x16Url"] = A.load.Ref(ptr + 68 + 0, undefined);
        x["iconSet"]["icon32x32Url"] = A.load.Ref(ptr + 68 + 4, undefined);
      } else {
        delete x["iconSet"];
      }
      x["driveLabel"] = A.load.Ref(ptr + 80, undefined);
      x["remoteMountPath"] = A.load.Ref(ptr + 84, undefined);
      if (A.load.Bool(ptr + 102)) {
        x["hidden"] = A.load.Bool(ptr + 88);
      } else {
        delete x["hidden"];
      }
      x["vmType"] = A.load.Enum(ptr + 92, ["termina", "plugin_vm", "borealis", "bruschetta", "arcvm"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetVolumeRootOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 6, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 6, true);
        A.store.Ref(ptr + 0, x["volumeId"]);
        A.store.Bool(ptr + 5, "writable" in x ? true : false);
        A.store.Bool(ptr + 4, x["writable"] ? true : false);
      }
    },
    "load_GetVolumeRootOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["volumeId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 5)) {
        x["writable"] = A.load.Bool(ptr + 4);
      } else {
        delete x["writable"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_HoldingSpaceState": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["itemUrls"]);
      }
    },
    "load_HoldingSpaceState": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["itemUrls"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_IOTaskParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Ref(ptr + 0, x["destinationFolder"]);
        A.store.Ref(ptr + 4, x["password"]);
        A.store.Bool(ptr + 9, "showNotification" in x ? true : false);
        A.store.Bool(ptr + 8, x["showNotification"] ? true : false);
      }
    },
    "load_IOTaskParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["destinationFolder"] = A.load.Ref(ptr + 0, undefined);
      x["password"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 9)) {
        x["showNotification"] = A.load.Bool(ptr + 8);
      } else {
        delete x["showNotification"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_IOTaskState": (ref: heap.Ref<string>): number => {
      const idx = [
        "queued",
        "scanning",
        "in_progress",
        "paused",
        "success",
        "error",
        "need_password",
        "cancelled",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_IOTaskType": (ref: heap.Ref<string>): number => {
      const idx = [
        "copy",
        "delete",
        "empty_trash",
        "extract",
        "move",
        "restore",
        "restore_to_destination",
        "trash",
        "zip",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_InspectionType": (ref: heap.Ref<string>): number => {
      const idx = ["normal", "console", "element", "background"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_InstallLinuxPackageResponse": (ref: heap.Ref<string>): number => {
      const idx = ["started", "failed", "install_already_active"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_MountableGuest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "id" in x ? true : false);
        A.store.Int32(ptr + 0, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Ref(ptr + 4, x["displayName"]);
        A.store.Enum(
          ptr + 8,
          ["termina", "plugin_vm", "borealis", "bruschetta", "arcvm"].indexOf(x["vmType"] as string)
        );
      }
    },
    "load_MountableGuest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["id"] = A.load.Int32(ptr + 0);
      } else {
        delete x["id"];
      }
      x["displayName"] = A.load.Ref(ptr + 4, undefined);
      x["vmType"] = A.load.Enum(ptr + 8, ["termina", "plugin_vm", "borealis", "bruschetta", "arcvm"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_MountCompletedEventType": (ref: heap.Ref<string>): number => {
      const idx = ["mount", "unmount"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_MountCompletedEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 114, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Enum(ptr + 4, -1);

        A.store.Bool(ptr + 8 + 103, false);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Ref(ptr + 8 + 4, undefined);
        A.store.Ref(ptr + 8 + 8, undefined);
        A.store.Enum(ptr + 8 + 12, -1);
        A.store.Ref(ptr + 8 + 16, undefined);

        A.store.Bool(ptr + 8 + 20 + 10, false);
        A.store.Ref(ptr + 8 + 20 + 0, undefined);
        A.store.Ref(ptr + 8 + 20 + 4, undefined);
        A.store.Bool(ptr + 8 + 20 + 9, false);
        A.store.Bool(ptr + 8 + 20 + 8, false);
        A.store.Ref(ptr + 8 + 32, undefined);
        A.store.Enum(ptr + 8 + 36, -1);
        A.store.Enum(ptr + 8 + 40, -1);
        A.store.Ref(ptr + 8 + 44, undefined);
        A.store.Bool(ptr + 8 + 96, false);
        A.store.Bool(ptr + 8 + 48, false);
        A.store.Bool(ptr + 8 + 97, false);
        A.store.Bool(ptr + 8 + 49, false);
        A.store.Bool(ptr + 8 + 98, false);
        A.store.Bool(ptr + 8 + 50, false);
        A.store.Bool(ptr + 8 + 99, false);
        A.store.Bool(ptr + 8 + 51, false);
        A.store.Bool(ptr + 8 + 100, false);
        A.store.Bool(ptr + 8 + 52, false);
        A.store.Bool(ptr + 8 + 101, false);
        A.store.Bool(ptr + 8 + 53, false);
        A.store.Enum(ptr + 8 + 56, -1);
        A.store.Enum(ptr + 8 + 60, -1);
        A.store.Ref(ptr + 8 + 64, undefined);

        A.store.Bool(ptr + 8 + 68 + 8, false);
        A.store.Ref(ptr + 8 + 68 + 0, undefined);
        A.store.Ref(ptr + 8 + 68 + 4, undefined);
        A.store.Ref(ptr + 8 + 80, undefined);
        A.store.Ref(ptr + 8 + 84, undefined);
        A.store.Bool(ptr + 8 + 102, false);
        A.store.Bool(ptr + 8 + 88, false);
        A.store.Enum(ptr + 8 + 92, -1);
        A.store.Bool(ptr + 113, false);
        A.store.Bool(ptr + 112, false);
      } else {
        A.store.Bool(ptr + 114, true);
        A.store.Enum(ptr + 0, ["mount", "unmount"].indexOf(x["eventType"] as string));
        A.store.Enum(
          ptr + 4,
          [
            "success",
            "in_progress",
            "unknown_error",
            "internal_error",
            "invalid_argument",
            "invalid_path",
            "path_already_mounted",
            "path_not_mounted",
            "directory_creation_failed",
            "invalid_mount_options",
            "insufficient_permissions",
            "mount_program_not_found",
            "mount_program_failed",
            "invalid_device_path",
            "unknown_filesystem",
            "unsupported_filesystem",
            "need_password",
            "cancelled",
            "busy",
          ].indexOf(x["status"] as string)
        );

        if (typeof x["volumeMetadata"] === "undefined") {
          A.store.Bool(ptr + 8 + 103, false);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Ref(ptr + 8 + 4, undefined);
          A.store.Ref(ptr + 8 + 8, undefined);
          A.store.Enum(ptr + 8 + 12, -1);
          A.store.Ref(ptr + 8 + 16, undefined);

          A.store.Bool(ptr + 8 + 20 + 10, false);
          A.store.Ref(ptr + 8 + 20 + 0, undefined);
          A.store.Ref(ptr + 8 + 20 + 4, undefined);
          A.store.Bool(ptr + 8 + 20 + 9, false);
          A.store.Bool(ptr + 8 + 20 + 8, false);
          A.store.Ref(ptr + 8 + 32, undefined);
          A.store.Enum(ptr + 8 + 36, -1);
          A.store.Enum(ptr + 8 + 40, -1);
          A.store.Ref(ptr + 8 + 44, undefined);
          A.store.Bool(ptr + 8 + 96, false);
          A.store.Bool(ptr + 8 + 48, false);
          A.store.Bool(ptr + 8 + 97, false);
          A.store.Bool(ptr + 8 + 49, false);
          A.store.Bool(ptr + 8 + 98, false);
          A.store.Bool(ptr + 8 + 50, false);
          A.store.Bool(ptr + 8 + 99, false);
          A.store.Bool(ptr + 8 + 51, false);
          A.store.Bool(ptr + 8 + 100, false);
          A.store.Bool(ptr + 8 + 52, false);
          A.store.Bool(ptr + 8 + 101, false);
          A.store.Bool(ptr + 8 + 53, false);
          A.store.Enum(ptr + 8 + 56, -1);
          A.store.Enum(ptr + 8 + 60, -1);
          A.store.Ref(ptr + 8 + 64, undefined);

          A.store.Bool(ptr + 8 + 68 + 8, false);
          A.store.Ref(ptr + 8 + 68 + 0, undefined);
          A.store.Ref(ptr + 8 + 68 + 4, undefined);
          A.store.Ref(ptr + 8 + 80, undefined);
          A.store.Ref(ptr + 8 + 84, undefined);
          A.store.Bool(ptr + 8 + 102, false);
          A.store.Bool(ptr + 8 + 88, false);
          A.store.Enum(ptr + 8 + 92, -1);
        } else {
          A.store.Bool(ptr + 8 + 103, true);
          A.store.Ref(ptr + 8 + 0, x["volumeMetadata"]["volumeId"]);
          A.store.Ref(ptr + 8 + 4, x["volumeMetadata"]["fileSystemId"]);
          A.store.Ref(ptr + 8 + 8, x["volumeMetadata"]["providerId"]);
          A.store.Enum(
            ptr + 8 + 12,
            ["file", "device", "network", "system"].indexOf(x["volumeMetadata"]["source"] as string)
          );
          A.store.Ref(ptr + 8 + 16, x["volumeMetadata"]["volumeLabel"]);

          if (typeof x["volumeMetadata"]["profile"] === "undefined") {
            A.store.Bool(ptr + 8 + 20 + 10, false);
            A.store.Ref(ptr + 8 + 20 + 0, undefined);
            A.store.Ref(ptr + 8 + 20 + 4, undefined);
            A.store.Bool(ptr + 8 + 20 + 9, false);
            A.store.Bool(ptr + 8 + 20 + 8, false);
          } else {
            A.store.Bool(ptr + 8 + 20 + 10, true);
            A.store.Ref(ptr + 8 + 20 + 0, x["volumeMetadata"]["profile"]["profileId"]);
            A.store.Ref(ptr + 8 + 20 + 4, x["volumeMetadata"]["profile"]["displayName"]);
            A.store.Bool(ptr + 8 + 20 + 9, "isCurrentProfile" in x["volumeMetadata"]["profile"] ? true : false);
            A.store.Bool(ptr + 8 + 20 + 8, x["volumeMetadata"]["profile"]["isCurrentProfile"] ? true : false);
          }
          A.store.Ref(ptr + 8 + 32, x["volumeMetadata"]["sourcePath"]);
          A.store.Enum(
            ptr + 8 + 36,
            [
              "drive",
              "downloads",
              "removable",
              "archive",
              "provided",
              "mtp",
              "media_view",
              "crostini",
              "android_files",
              "documents_provider",
              "testing",
              "smb",
              "system_internal",
              "guest_os",
            ].indexOf(x["volumeMetadata"]["volumeType"] as string)
          );
          A.store.Enum(
            ptr + 8 + 40,
            ["usb", "sd", "optical", "mobile", "unknown"].indexOf(x["volumeMetadata"]["deviceType"] as string)
          );
          A.store.Ref(ptr + 8 + 44, x["volumeMetadata"]["devicePath"]);
          A.store.Bool(ptr + 8 + 96, "isParentDevice" in x["volumeMetadata"] ? true : false);
          A.store.Bool(ptr + 8 + 48, x["volumeMetadata"]["isParentDevice"] ? true : false);
          A.store.Bool(ptr + 8 + 97, "isReadOnly" in x["volumeMetadata"] ? true : false);
          A.store.Bool(ptr + 8 + 49, x["volumeMetadata"]["isReadOnly"] ? true : false);
          A.store.Bool(ptr + 8 + 98, "isReadOnlyRemovableDevice" in x["volumeMetadata"] ? true : false);
          A.store.Bool(ptr + 8 + 50, x["volumeMetadata"]["isReadOnlyRemovableDevice"] ? true : false);
          A.store.Bool(ptr + 8 + 99, "hasMedia" in x["volumeMetadata"] ? true : false);
          A.store.Bool(ptr + 8 + 51, x["volumeMetadata"]["hasMedia"] ? true : false);
          A.store.Bool(ptr + 8 + 100, "configurable" in x["volumeMetadata"] ? true : false);
          A.store.Bool(ptr + 8 + 52, x["volumeMetadata"]["configurable"] ? true : false);
          A.store.Bool(ptr + 8 + 101, "watchable" in x["volumeMetadata"] ? true : false);
          A.store.Bool(ptr + 8 + 53, x["volumeMetadata"]["watchable"] ? true : false);
          A.store.Enum(
            ptr + 8 + 56,
            [
              "success",
              "in_progress",
              "unknown_error",
              "internal_error",
              "invalid_argument",
              "invalid_path",
              "path_already_mounted",
              "path_not_mounted",
              "directory_creation_failed",
              "invalid_mount_options",
              "insufficient_permissions",
              "mount_program_not_found",
              "mount_program_failed",
              "invalid_device_path",
              "unknown_filesystem",
              "unsupported_filesystem",
              "need_password",
              "cancelled",
              "busy",
            ].indexOf(x["volumeMetadata"]["mountCondition"] as string)
          );
          A.store.Enum(ptr + 8 + 60, ["user", "auto"].indexOf(x["volumeMetadata"]["mountContext"] as string));
          A.store.Ref(ptr + 8 + 64, x["volumeMetadata"]["diskFileSystemType"]);

          if (typeof x["volumeMetadata"]["iconSet"] === "undefined") {
            A.store.Bool(ptr + 8 + 68 + 8, false);
            A.store.Ref(ptr + 8 + 68 + 0, undefined);
            A.store.Ref(ptr + 8 + 68 + 4, undefined);
          } else {
            A.store.Bool(ptr + 8 + 68 + 8, true);
            A.store.Ref(ptr + 8 + 68 + 0, x["volumeMetadata"]["iconSet"]["icon16x16Url"]);
            A.store.Ref(ptr + 8 + 68 + 4, x["volumeMetadata"]["iconSet"]["icon32x32Url"]);
          }
          A.store.Ref(ptr + 8 + 80, x["volumeMetadata"]["driveLabel"]);
          A.store.Ref(ptr + 8 + 84, x["volumeMetadata"]["remoteMountPath"]);
          A.store.Bool(ptr + 8 + 102, "hidden" in x["volumeMetadata"] ? true : false);
          A.store.Bool(ptr + 8 + 88, x["volumeMetadata"]["hidden"] ? true : false);
          A.store.Enum(
            ptr + 8 + 92,
            ["termina", "plugin_vm", "borealis", "bruschetta", "arcvm"].indexOf(x["volumeMetadata"]["vmType"] as string)
          );
        }
        A.store.Bool(ptr + 113, "shouldNotify" in x ? true : false);
        A.store.Bool(ptr + 112, x["shouldNotify"] ? true : false);
      }
    },
    "load_MountCompletedEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["eventType"] = A.load.Enum(ptr + 0, ["mount", "unmount"]);
      x["status"] = A.load.Enum(ptr + 4, [
        "success",
        "in_progress",
        "unknown_error",
        "internal_error",
        "invalid_argument",
        "invalid_path",
        "path_already_mounted",
        "path_not_mounted",
        "directory_creation_failed",
        "invalid_mount_options",
        "insufficient_permissions",
        "mount_program_not_found",
        "mount_program_failed",
        "invalid_device_path",
        "unknown_filesystem",
        "unsupported_filesystem",
        "need_password",
        "cancelled",
        "busy",
      ]);
      if (A.load.Bool(ptr + 8 + 103)) {
        x["volumeMetadata"] = {};
        x["volumeMetadata"]["volumeId"] = A.load.Ref(ptr + 8 + 0, undefined);
        x["volumeMetadata"]["fileSystemId"] = A.load.Ref(ptr + 8 + 4, undefined);
        x["volumeMetadata"]["providerId"] = A.load.Ref(ptr + 8 + 8, undefined);
        x["volumeMetadata"]["source"] = A.load.Enum(ptr + 8 + 12, ["file", "device", "network", "system"]);
        x["volumeMetadata"]["volumeLabel"] = A.load.Ref(ptr + 8 + 16, undefined);
        if (A.load.Bool(ptr + 8 + 20 + 10)) {
          x["volumeMetadata"]["profile"] = {};
          x["volumeMetadata"]["profile"]["profileId"] = A.load.Ref(ptr + 8 + 20 + 0, undefined);
          x["volumeMetadata"]["profile"]["displayName"] = A.load.Ref(ptr + 8 + 20 + 4, undefined);
          if (A.load.Bool(ptr + 8 + 20 + 9)) {
            x["volumeMetadata"]["profile"]["isCurrentProfile"] = A.load.Bool(ptr + 8 + 20 + 8);
          } else {
            delete x["volumeMetadata"]["profile"]["isCurrentProfile"];
          }
        } else {
          delete x["volumeMetadata"]["profile"];
        }
        x["volumeMetadata"]["sourcePath"] = A.load.Ref(ptr + 8 + 32, undefined);
        x["volumeMetadata"]["volumeType"] = A.load.Enum(ptr + 8 + 36, [
          "drive",
          "downloads",
          "removable",
          "archive",
          "provided",
          "mtp",
          "media_view",
          "crostini",
          "android_files",
          "documents_provider",
          "testing",
          "smb",
          "system_internal",
          "guest_os",
        ]);
        x["volumeMetadata"]["deviceType"] = A.load.Enum(ptr + 8 + 40, ["usb", "sd", "optical", "mobile", "unknown"]);
        x["volumeMetadata"]["devicePath"] = A.load.Ref(ptr + 8 + 44, undefined);
        if (A.load.Bool(ptr + 8 + 96)) {
          x["volumeMetadata"]["isParentDevice"] = A.load.Bool(ptr + 8 + 48);
        } else {
          delete x["volumeMetadata"]["isParentDevice"];
        }
        if (A.load.Bool(ptr + 8 + 97)) {
          x["volumeMetadata"]["isReadOnly"] = A.load.Bool(ptr + 8 + 49);
        } else {
          delete x["volumeMetadata"]["isReadOnly"];
        }
        if (A.load.Bool(ptr + 8 + 98)) {
          x["volumeMetadata"]["isReadOnlyRemovableDevice"] = A.load.Bool(ptr + 8 + 50);
        } else {
          delete x["volumeMetadata"]["isReadOnlyRemovableDevice"];
        }
        if (A.load.Bool(ptr + 8 + 99)) {
          x["volumeMetadata"]["hasMedia"] = A.load.Bool(ptr + 8 + 51);
        } else {
          delete x["volumeMetadata"]["hasMedia"];
        }
        if (A.load.Bool(ptr + 8 + 100)) {
          x["volumeMetadata"]["configurable"] = A.load.Bool(ptr + 8 + 52);
        } else {
          delete x["volumeMetadata"]["configurable"];
        }
        if (A.load.Bool(ptr + 8 + 101)) {
          x["volumeMetadata"]["watchable"] = A.load.Bool(ptr + 8 + 53);
        } else {
          delete x["volumeMetadata"]["watchable"];
        }
        x["volumeMetadata"]["mountCondition"] = A.load.Enum(ptr + 8 + 56, [
          "success",
          "in_progress",
          "unknown_error",
          "internal_error",
          "invalid_argument",
          "invalid_path",
          "path_already_mounted",
          "path_not_mounted",
          "directory_creation_failed",
          "invalid_mount_options",
          "insufficient_permissions",
          "mount_program_not_found",
          "mount_program_failed",
          "invalid_device_path",
          "unknown_filesystem",
          "unsupported_filesystem",
          "need_password",
          "cancelled",
          "busy",
        ]);
        x["volumeMetadata"]["mountContext"] = A.load.Enum(ptr + 8 + 60, ["user", "auto"]);
        x["volumeMetadata"]["diskFileSystemType"] = A.load.Ref(ptr + 8 + 64, undefined);
        if (A.load.Bool(ptr + 8 + 68 + 8)) {
          x["volumeMetadata"]["iconSet"] = {};
          x["volumeMetadata"]["iconSet"]["icon16x16Url"] = A.load.Ref(ptr + 8 + 68 + 0, undefined);
          x["volumeMetadata"]["iconSet"]["icon32x32Url"] = A.load.Ref(ptr + 8 + 68 + 4, undefined);
        } else {
          delete x["volumeMetadata"]["iconSet"];
        }
        x["volumeMetadata"]["driveLabel"] = A.load.Ref(ptr + 8 + 80, undefined);
        x["volumeMetadata"]["remoteMountPath"] = A.load.Ref(ptr + 8 + 84, undefined);
        if (A.load.Bool(ptr + 8 + 102)) {
          x["volumeMetadata"]["hidden"] = A.load.Bool(ptr + 8 + 88);
        } else {
          delete x["volumeMetadata"]["hidden"];
        }
        x["volumeMetadata"]["vmType"] = A.load.Enum(ptr + 8 + 92, [
          "termina",
          "plugin_vm",
          "borealis",
          "bruschetta",
          "arcvm",
        ]);
      } else {
        delete x["volumeMetadata"];
      }
      if (A.load.Bool(ptr + 113)) {
        x["shouldNotify"] = A.load.Bool(ptr + 112);
      } else {
        delete x["shouldNotify"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OpenWindowParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["currentDirectoryURL"]);
        A.store.Ref(ptr + 4, x["selectionURL"]);
      }
    },
    "load_OpenWindowParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["currentDirectoryURL"] = A.load.Ref(ptr + 0, undefined);
      x["selectionURL"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ParsedTrashInfoFile": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["restoreEntry"]);
        A.store.Ref(ptr + 4, x["trashInfoFileName"]);
        A.store.Bool(ptr + 16, "deletionDate" in x ? true : false);
        A.store.Float64(ptr + 8, x["deletionDate"] === undefined ? 0 : (x["deletionDate"] as number));
      }
    },
    "load_ParsedTrashInfoFile": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["restoreEntry"] = A.load.Ref(ptr + 0, undefined);
      x["trashInfoFileName"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["deletionDate"] = A.load.Float64(ptr + 8);
      } else {
        delete x["deletionDate"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PolicyErrorType": (ref: heap.Ref<string>): number => {
      const idx = ["dlp", "enterprise_connectors", "dlp_warning_timeout"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PolicyPauseParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Enum(ptr + 0, ["dlp", "enterprise_connectors", "dlp_warning_timeout"].indexOf(x["type"] as string));
        A.store.Bool(ptr + 12, "policyFileCount" in x ? true : false);
        A.store.Int32(ptr + 4, x["policyFileCount"] === undefined ? 0 : (x["policyFileCount"] as number));
        A.store.Ref(ptr + 8, x["fileName"]);
      }
    },
    "load_PolicyPauseParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, ["dlp", "enterprise_connectors", "dlp_warning_timeout"]);
      if (A.load.Bool(ptr + 12)) {
        x["policyFileCount"] = A.load.Int32(ptr + 4);
      } else {
        delete x["policyFileCount"];
      }
      x["fileName"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PauseParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 30, false);

        A.store.Bool(ptr + 0 + 14, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Bool(ptr + 0 + 12, false);
        A.store.Bool(ptr + 0 + 4, false);
        A.store.Bool(ptr + 0 + 13, false);
        A.store.Bool(ptr + 0 + 5, false);
        A.store.Ref(ptr + 0 + 8, undefined);

        A.store.Bool(ptr + 16 + 13, false);
        A.store.Enum(ptr + 16 + 0, -1);
        A.store.Bool(ptr + 16 + 12, false);
        A.store.Int32(ptr + 16 + 4, 0);
        A.store.Ref(ptr + 16 + 8, undefined);
      } else {
        A.store.Bool(ptr + 30, true);

        if (typeof x["conflictParams"] === "undefined") {
          A.store.Bool(ptr + 0 + 14, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Bool(ptr + 0 + 12, false);
          A.store.Bool(ptr + 0 + 4, false);
          A.store.Bool(ptr + 0 + 13, false);
          A.store.Bool(ptr + 0 + 5, false);
          A.store.Ref(ptr + 0 + 8, undefined);
        } else {
          A.store.Bool(ptr + 0 + 14, true);
          A.store.Ref(ptr + 0 + 0, x["conflictParams"]["conflictName"]);
          A.store.Bool(ptr + 0 + 12, "conflictIsDirectory" in x["conflictParams"] ? true : false);
          A.store.Bool(ptr + 0 + 4, x["conflictParams"]["conflictIsDirectory"] ? true : false);
          A.store.Bool(ptr + 0 + 13, "conflictMultiple" in x["conflictParams"] ? true : false);
          A.store.Bool(ptr + 0 + 5, x["conflictParams"]["conflictMultiple"] ? true : false);
          A.store.Ref(ptr + 0 + 8, x["conflictParams"]["conflictTargetUrl"]);
        }

        if (typeof x["policyParams"] === "undefined") {
          A.store.Bool(ptr + 16 + 13, false);
          A.store.Enum(ptr + 16 + 0, -1);
          A.store.Bool(ptr + 16 + 12, false);
          A.store.Int32(ptr + 16 + 4, 0);
          A.store.Ref(ptr + 16 + 8, undefined);
        } else {
          A.store.Bool(ptr + 16 + 13, true);
          A.store.Enum(
            ptr + 16 + 0,
            ["dlp", "enterprise_connectors", "dlp_warning_timeout"].indexOf(x["policyParams"]["type"] as string)
          );
          A.store.Bool(ptr + 16 + 12, "policyFileCount" in x["policyParams"] ? true : false);
          A.store.Int32(
            ptr + 16 + 4,
            x["policyParams"]["policyFileCount"] === undefined ? 0 : (x["policyParams"]["policyFileCount"] as number)
          );
          A.store.Ref(ptr + 16 + 8, x["policyParams"]["fileName"]);
        }
      }
    },
    "load_PauseParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 14)) {
        x["conflictParams"] = {};
        x["conflictParams"]["conflictName"] = A.load.Ref(ptr + 0 + 0, undefined);
        if (A.load.Bool(ptr + 0 + 12)) {
          x["conflictParams"]["conflictIsDirectory"] = A.load.Bool(ptr + 0 + 4);
        } else {
          delete x["conflictParams"]["conflictIsDirectory"];
        }
        if (A.load.Bool(ptr + 0 + 13)) {
          x["conflictParams"]["conflictMultiple"] = A.load.Bool(ptr + 0 + 5);
        } else {
          delete x["conflictParams"]["conflictMultiple"];
        }
        x["conflictParams"]["conflictTargetUrl"] = A.load.Ref(ptr + 0 + 8, undefined);
      } else {
        delete x["conflictParams"];
      }
      if (A.load.Bool(ptr + 16 + 13)) {
        x["policyParams"] = {};
        x["policyParams"]["type"] = A.load.Enum(ptr + 16 + 0, ["dlp", "enterprise_connectors", "dlp_warning_timeout"]);
        if (A.load.Bool(ptr + 16 + 12)) {
          x["policyParams"]["policyFileCount"] = A.load.Int32(ptr + 16 + 4);
        } else {
          delete x["policyParams"]["policyFileCount"];
        }
        x["policyParams"]["fileName"] = A.load.Ref(ptr + 16 + 8, undefined);
      } else {
        delete x["policyParams"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PolicyDialogType": (ref: heap.Ref<string>): number => {
      const idx = ["warning", "error"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PolicyError": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Enum(ptr + 0, ["dlp", "enterprise_connectors", "dlp_warning_timeout"].indexOf(x["type"] as string));
        A.store.Bool(ptr + 12, "policyFileCount" in x ? true : false);
        A.store.Int32(ptr + 4, x["policyFileCount"] === undefined ? 0 : (x["policyFileCount"] as number));
        A.store.Ref(ptr + 8, x["fileName"]);
      }
    },
    "load_PolicyError": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, ["dlp", "enterprise_connectors", "dlp_warning_timeout"]);
      if (A.load.Bool(ptr + 12)) {
        x["policyFileCount"] = A.load.Int32(ptr + 4);
      } else {
        delete x["policyFileCount"];
      }
      x["fileName"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PolicyResumeParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["dlp", "enterprise_connectors", "dlp_warning_timeout"].indexOf(x["type"] as string));
      }
    },
    "load_PolicyResumeParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, ["dlp", "enterprise_connectors", "dlp_warning_timeout"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PreferencesChange": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 2, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 9, "driveSyncEnabledOnMeteredNetwork" in x ? true : false);
        A.store.Bool(ptr + 0, x["driveSyncEnabledOnMeteredNetwork"] ? true : false);
        A.store.Bool(ptr + 10, "arcEnabled" in x ? true : false);
        A.store.Bool(ptr + 1, x["arcEnabled"] ? true : false);
        A.store.Bool(ptr + 11, "arcRemovableMediaAccessEnabled" in x ? true : false);
        A.store.Bool(ptr + 2, x["arcRemovableMediaAccessEnabled"] ? true : false);
        A.store.Ref(ptr + 4, x["folderShortcuts"]);
        A.store.Bool(ptr + 12, "driveFsBulkPinningEnabled" in x ? true : false);
        A.store.Bool(ptr + 8, x["driveFsBulkPinningEnabled"] ? true : false);
      }
    },
    "load_PreferencesChange": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 9)) {
        x["driveSyncEnabledOnMeteredNetwork"] = A.load.Bool(ptr + 0);
      } else {
        delete x["driveSyncEnabledOnMeteredNetwork"];
      }
      if (A.load.Bool(ptr + 10)) {
        x["arcEnabled"] = A.load.Bool(ptr + 1);
      } else {
        delete x["arcEnabled"];
      }
      if (A.load.Bool(ptr + 11)) {
        x["arcRemovableMediaAccessEnabled"] = A.load.Bool(ptr + 2);
      } else {
        delete x["arcRemovableMediaAccessEnabled"];
      }
      x["folderShortcuts"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["driveFsBulkPinningEnabled"] = A.load.Bool(ptr + 8);
      } else {
        delete x["driveFsBulkPinningEnabled"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ProgressStatus": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 124, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Enum(ptr + 4, -1);

        A.store.Bool(ptr + 8 + 13, false);
        A.store.Enum(ptr + 8 + 0, -1);
        A.store.Bool(ptr + 8 + 12, false);
        A.store.Int32(ptr + 8 + 4, 0);
        A.store.Ref(ptr + 8 + 8, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Bool(ptr + 116, false);
        A.store.Int32(ptr + 28, 0);
        A.store.Bool(ptr + 117, false);
        A.store.Int32(ptr + 32, 0);
        A.store.Ref(ptr + 36, undefined);
        A.store.Bool(ptr + 118, false);
        A.store.Int32(ptr + 40, 0);
        A.store.Bool(ptr + 119, false);
        A.store.Int32(ptr + 44, 0);
        A.store.Bool(ptr + 120, false);
        A.store.Int32(ptr + 48, 0);
        A.store.Bool(ptr + 121, false);
        A.store.Float64(ptr + 56, 0);
        A.store.Bool(ptr + 122, false);
        A.store.Int32(ptr + 64, 0);
        A.store.Bool(ptr + 123, false);
        A.store.Bool(ptr + 68, false);
        A.store.Ref(ptr + 72, undefined);

        A.store.Bool(ptr + 76 + 30, false);

        A.store.Bool(ptr + 76 + 0 + 14, false);
        A.store.Ref(ptr + 76 + 0 + 0, undefined);
        A.store.Bool(ptr + 76 + 0 + 12, false);
        A.store.Bool(ptr + 76 + 0 + 4, false);
        A.store.Bool(ptr + 76 + 0 + 13, false);
        A.store.Bool(ptr + 76 + 0 + 5, false);
        A.store.Ref(ptr + 76 + 0 + 8, undefined);

        A.store.Bool(ptr + 76 + 16 + 13, false);
        A.store.Enum(ptr + 76 + 16 + 0, -1);
        A.store.Bool(ptr + 76 + 16 + 12, false);
        A.store.Int32(ptr + 76 + 16 + 4, 0);
        A.store.Ref(ptr + 76 + 16 + 8, undefined);
        A.store.Ref(ptr + 108, undefined);
        A.store.Ref(ptr + 112, undefined);
      } else {
        A.store.Bool(ptr + 124, true);
        A.store.Enum(
          ptr + 0,
          [
            "copy",
            "delete",
            "empty_trash",
            "extract",
            "move",
            "restore",
            "restore_to_destination",
            "trash",
            "zip",
          ].indexOf(x["type"] as string)
        );
        A.store.Enum(
          ptr + 4,
          ["queued", "scanning", "in_progress", "paused", "success", "error", "need_password", "cancelled"].indexOf(
            x["state"] as string
          )
        );

        if (typeof x["policyError"] === "undefined") {
          A.store.Bool(ptr + 8 + 13, false);
          A.store.Enum(ptr + 8 + 0, -1);
          A.store.Bool(ptr + 8 + 12, false);
          A.store.Int32(ptr + 8 + 4, 0);
          A.store.Ref(ptr + 8 + 8, undefined);
        } else {
          A.store.Bool(ptr + 8 + 13, true);
          A.store.Enum(
            ptr + 8 + 0,
            ["dlp", "enterprise_connectors", "dlp_warning_timeout"].indexOf(x["policyError"]["type"] as string)
          );
          A.store.Bool(ptr + 8 + 12, "policyFileCount" in x["policyError"] ? true : false);
          A.store.Int32(
            ptr + 8 + 4,
            x["policyError"]["policyFileCount"] === undefined ? 0 : (x["policyError"]["policyFileCount"] as number)
          );
          A.store.Ref(ptr + 8 + 8, x["policyError"]["fileName"]);
        }
        A.store.Ref(ptr + 24, x["sourceName"]);
        A.store.Bool(ptr + 116, "numRemainingItems" in x ? true : false);
        A.store.Int32(ptr + 28, x["numRemainingItems"] === undefined ? 0 : (x["numRemainingItems"] as number));
        A.store.Bool(ptr + 117, "itemCount" in x ? true : false);
        A.store.Int32(ptr + 32, x["itemCount"] === undefined ? 0 : (x["itemCount"] as number));
        A.store.Ref(ptr + 36, x["destinationName"]);
        A.store.Bool(ptr + 118, "bytesTransferred" in x ? true : false);
        A.store.Int32(ptr + 40, x["bytesTransferred"] === undefined ? 0 : (x["bytesTransferred"] as number));
        A.store.Bool(ptr + 119, "totalBytes" in x ? true : false);
        A.store.Int32(ptr + 44, x["totalBytes"] === undefined ? 0 : (x["totalBytes"] as number));
        A.store.Bool(ptr + 120, "taskId" in x ? true : false);
        A.store.Int32(ptr + 48, x["taskId"] === undefined ? 0 : (x["taskId"] as number));
        A.store.Bool(ptr + 121, "remainingSeconds" in x ? true : false);
        A.store.Float64(ptr + 56, x["remainingSeconds"] === undefined ? 0 : (x["remainingSeconds"] as number));
        A.store.Bool(ptr + 122, "sourcesScanned" in x ? true : false);
        A.store.Int32(ptr + 64, x["sourcesScanned"] === undefined ? 0 : (x["sourcesScanned"] as number));
        A.store.Bool(ptr + 123, "showNotification" in x ? true : false);
        A.store.Bool(ptr + 68, x["showNotification"] ? true : false);
        A.store.Ref(ptr + 72, x["errorName"]);

        if (typeof x["pauseParams"] === "undefined") {
          A.store.Bool(ptr + 76 + 30, false);

          A.store.Bool(ptr + 76 + 0 + 14, false);
          A.store.Ref(ptr + 76 + 0 + 0, undefined);
          A.store.Bool(ptr + 76 + 0 + 12, false);
          A.store.Bool(ptr + 76 + 0 + 4, false);
          A.store.Bool(ptr + 76 + 0 + 13, false);
          A.store.Bool(ptr + 76 + 0 + 5, false);
          A.store.Ref(ptr + 76 + 0 + 8, undefined);

          A.store.Bool(ptr + 76 + 16 + 13, false);
          A.store.Enum(ptr + 76 + 16 + 0, -1);
          A.store.Bool(ptr + 76 + 16 + 12, false);
          A.store.Int32(ptr + 76 + 16 + 4, 0);
          A.store.Ref(ptr + 76 + 16 + 8, undefined);
        } else {
          A.store.Bool(ptr + 76 + 30, true);

          if (typeof x["pauseParams"]["conflictParams"] === "undefined") {
            A.store.Bool(ptr + 76 + 0 + 14, false);
            A.store.Ref(ptr + 76 + 0 + 0, undefined);
            A.store.Bool(ptr + 76 + 0 + 12, false);
            A.store.Bool(ptr + 76 + 0 + 4, false);
            A.store.Bool(ptr + 76 + 0 + 13, false);
            A.store.Bool(ptr + 76 + 0 + 5, false);
            A.store.Ref(ptr + 76 + 0 + 8, undefined);
          } else {
            A.store.Bool(ptr + 76 + 0 + 14, true);
            A.store.Ref(ptr + 76 + 0 + 0, x["pauseParams"]["conflictParams"]["conflictName"]);
            A.store.Bool(ptr + 76 + 0 + 12, "conflictIsDirectory" in x["pauseParams"]["conflictParams"] ? true : false);
            A.store.Bool(ptr + 76 + 0 + 4, x["pauseParams"]["conflictParams"]["conflictIsDirectory"] ? true : false);
            A.store.Bool(ptr + 76 + 0 + 13, "conflictMultiple" in x["pauseParams"]["conflictParams"] ? true : false);
            A.store.Bool(ptr + 76 + 0 + 5, x["pauseParams"]["conflictParams"]["conflictMultiple"] ? true : false);
            A.store.Ref(ptr + 76 + 0 + 8, x["pauseParams"]["conflictParams"]["conflictTargetUrl"]);
          }

          if (typeof x["pauseParams"]["policyParams"] === "undefined") {
            A.store.Bool(ptr + 76 + 16 + 13, false);
            A.store.Enum(ptr + 76 + 16 + 0, -1);
            A.store.Bool(ptr + 76 + 16 + 12, false);
            A.store.Int32(ptr + 76 + 16 + 4, 0);
            A.store.Ref(ptr + 76 + 16 + 8, undefined);
          } else {
            A.store.Bool(ptr + 76 + 16 + 13, true);
            A.store.Enum(
              ptr + 76 + 16 + 0,
              ["dlp", "enterprise_connectors", "dlp_warning_timeout"].indexOf(
                x["pauseParams"]["policyParams"]["type"] as string
              )
            );
            A.store.Bool(ptr + 76 + 16 + 12, "policyFileCount" in x["pauseParams"]["policyParams"] ? true : false);
            A.store.Int32(
              ptr + 76 + 16 + 4,
              x["pauseParams"]["policyParams"]["policyFileCount"] === undefined
                ? 0
                : (x["pauseParams"]["policyParams"]["policyFileCount"] as number)
            );
            A.store.Ref(ptr + 76 + 16 + 8, x["pauseParams"]["policyParams"]["fileName"]);
          }
        }
        A.store.Ref(ptr + 108, x["outputs"]);
        A.store.Ref(ptr + 112, x["destinationVolumeId"]);
      }
    },
    "load_ProgressStatus": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, [
        "copy",
        "delete",
        "empty_trash",
        "extract",
        "move",
        "restore",
        "restore_to_destination",
        "trash",
        "zip",
      ]);
      x["state"] = A.load.Enum(ptr + 4, [
        "queued",
        "scanning",
        "in_progress",
        "paused",
        "success",
        "error",
        "need_password",
        "cancelled",
      ]);
      if (A.load.Bool(ptr + 8 + 13)) {
        x["policyError"] = {};
        x["policyError"]["type"] = A.load.Enum(ptr + 8 + 0, ["dlp", "enterprise_connectors", "dlp_warning_timeout"]);
        if (A.load.Bool(ptr + 8 + 12)) {
          x["policyError"]["policyFileCount"] = A.load.Int32(ptr + 8 + 4);
        } else {
          delete x["policyError"]["policyFileCount"];
        }
        x["policyError"]["fileName"] = A.load.Ref(ptr + 8 + 8, undefined);
      } else {
        delete x["policyError"];
      }
      x["sourceName"] = A.load.Ref(ptr + 24, undefined);
      if (A.load.Bool(ptr + 116)) {
        x["numRemainingItems"] = A.load.Int32(ptr + 28);
      } else {
        delete x["numRemainingItems"];
      }
      if (A.load.Bool(ptr + 117)) {
        x["itemCount"] = A.load.Int32(ptr + 32);
      } else {
        delete x["itemCount"];
      }
      x["destinationName"] = A.load.Ref(ptr + 36, undefined);
      if (A.load.Bool(ptr + 118)) {
        x["bytesTransferred"] = A.load.Int32(ptr + 40);
      } else {
        delete x["bytesTransferred"];
      }
      if (A.load.Bool(ptr + 119)) {
        x["totalBytes"] = A.load.Int32(ptr + 44);
      } else {
        delete x["totalBytes"];
      }
      if (A.load.Bool(ptr + 120)) {
        x["taskId"] = A.load.Int32(ptr + 48);
      } else {
        delete x["taskId"];
      }
      if (A.load.Bool(ptr + 121)) {
        x["remainingSeconds"] = A.load.Float64(ptr + 56);
      } else {
        delete x["remainingSeconds"];
      }
      if (A.load.Bool(ptr + 122)) {
        x["sourcesScanned"] = A.load.Int32(ptr + 64);
      } else {
        delete x["sourcesScanned"];
      }
      if (A.load.Bool(ptr + 123)) {
        x["showNotification"] = A.load.Bool(ptr + 68);
      } else {
        delete x["showNotification"];
      }
      x["errorName"] = A.load.Ref(ptr + 72, undefined);
      if (A.load.Bool(ptr + 76 + 30)) {
        x["pauseParams"] = {};
        if (A.load.Bool(ptr + 76 + 0 + 14)) {
          x["pauseParams"]["conflictParams"] = {};
          x["pauseParams"]["conflictParams"]["conflictName"] = A.load.Ref(ptr + 76 + 0 + 0, undefined);
          if (A.load.Bool(ptr + 76 + 0 + 12)) {
            x["pauseParams"]["conflictParams"]["conflictIsDirectory"] = A.load.Bool(ptr + 76 + 0 + 4);
          } else {
            delete x["pauseParams"]["conflictParams"]["conflictIsDirectory"];
          }
          if (A.load.Bool(ptr + 76 + 0 + 13)) {
            x["pauseParams"]["conflictParams"]["conflictMultiple"] = A.load.Bool(ptr + 76 + 0 + 5);
          } else {
            delete x["pauseParams"]["conflictParams"]["conflictMultiple"];
          }
          x["pauseParams"]["conflictParams"]["conflictTargetUrl"] = A.load.Ref(ptr + 76 + 0 + 8, undefined);
        } else {
          delete x["pauseParams"]["conflictParams"];
        }
        if (A.load.Bool(ptr + 76 + 16 + 13)) {
          x["pauseParams"]["policyParams"] = {};
          x["pauseParams"]["policyParams"]["type"] = A.load.Enum(ptr + 76 + 16 + 0, [
            "dlp",
            "enterprise_connectors",
            "dlp_warning_timeout",
          ]);
          if (A.load.Bool(ptr + 76 + 16 + 12)) {
            x["pauseParams"]["policyParams"]["policyFileCount"] = A.load.Int32(ptr + 76 + 16 + 4);
          } else {
            delete x["pauseParams"]["policyParams"]["policyFileCount"];
          }
          x["pauseParams"]["policyParams"]["fileName"] = A.load.Ref(ptr + 76 + 16 + 8, undefined);
        } else {
          delete x["pauseParams"]["policyParams"];
        }
      } else {
        delete x["pauseParams"];
      }
      x["outputs"] = A.load.Ref(ptr + 108, undefined);
      x["destinationVolumeId"] = A.load.Ref(ptr + 112, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ResumeParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);

        A.store.Bool(ptr + 0 + 6, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Bool(ptr + 0 + 5, false);
        A.store.Bool(ptr + 0 + 4, false);

        A.store.Bool(ptr + 8 + 4, false);
        A.store.Enum(ptr + 8 + 0, -1);
      } else {
        A.store.Bool(ptr + 13, true);

        if (typeof x["conflictParams"] === "undefined") {
          A.store.Bool(ptr + 0 + 6, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Bool(ptr + 0 + 5, false);
          A.store.Bool(ptr + 0 + 4, false);
        } else {
          A.store.Bool(ptr + 0 + 6, true);
          A.store.Ref(ptr + 0 + 0, x["conflictParams"]["conflictResolve"]);
          A.store.Bool(ptr + 0 + 5, "conflictApplyToAll" in x["conflictParams"] ? true : false);
          A.store.Bool(ptr + 0 + 4, x["conflictParams"]["conflictApplyToAll"] ? true : false);
        }

        if (typeof x["policyParams"] === "undefined") {
          A.store.Bool(ptr + 8 + 4, false);
          A.store.Enum(ptr + 8 + 0, -1);
        } else {
          A.store.Bool(ptr + 8 + 4, true);
          A.store.Enum(
            ptr + 8 + 0,
            ["dlp", "enterprise_connectors", "dlp_warning_timeout"].indexOf(x["policyParams"]["type"] as string)
          );
        }
      }
    },
    "load_ResumeParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 6)) {
        x["conflictParams"] = {};
        x["conflictParams"]["conflictResolve"] = A.load.Ref(ptr + 0 + 0, undefined);
        if (A.load.Bool(ptr + 0 + 5)) {
          x["conflictParams"]["conflictApplyToAll"] = A.load.Bool(ptr + 0 + 4);
        } else {
          delete x["conflictParams"]["conflictApplyToAll"];
        }
      } else {
        delete x["conflictParams"];
      }
      if (A.load.Bool(ptr + 8 + 4)) {
        x["policyParams"] = {};
        x["policyParams"]["type"] = A.load.Enum(ptr + 8 + 0, ["dlp", "enterprise_connectors", "dlp_warning_timeout"]);
      } else {
        delete x["policyParams"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_SearchType": (ref: heap.Ref<string>): number => {
      const idx = ["EXCLUDE_DIRECTORIES", "SHARED_WITH_ME", "OFFLINE", "ALL"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SearchMetadataParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 30, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Bool(ptr + 28, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 29, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Enum(ptr + 24, -1);
      } else {
        A.store.Bool(ptr + 30, true);
        A.store.Ref(ptr + 0, x["rootDir"]);
        A.store.Ref(ptr + 4, x["query"]);
        A.store.Enum(
          ptr + 8,
          ["EXCLUDE_DIRECTORIES", "SHARED_WITH_ME", "OFFLINE", "ALL"].indexOf(x["types"] as string)
        );
        A.store.Bool(ptr + 28, "maxResults" in x ? true : false);
        A.store.Int32(ptr + 12, x["maxResults"] === undefined ? 0 : (x["maxResults"] as number));
        A.store.Bool(ptr + 29, "modifiedTimestamp" in x ? true : false);
        A.store.Float64(ptr + 16, x["modifiedTimestamp"] === undefined ? 0 : (x["modifiedTimestamp"] as number));
        A.store.Enum(ptr + 24, ["all", "audio", "image", "video", "document"].indexOf(x["category"] as string));
      }
    },
    "load_SearchMetadataParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["rootDir"] = A.load.Ref(ptr + 0, undefined);
      x["query"] = A.load.Ref(ptr + 4, undefined);
      x["types"] = A.load.Enum(ptr + 8, ["EXCLUDE_DIRECTORIES", "SHARED_WITH_ME", "OFFLINE", "ALL"]);
      if (A.load.Bool(ptr + 28)) {
        x["maxResults"] = A.load.Int32(ptr + 12);
      } else {
        delete x["maxResults"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["modifiedTimestamp"] = A.load.Float64(ptr + 16);
      } else {
        delete x["modifiedTimestamp"];
      }
      x["category"] = A.load.Enum(ptr + 24, ["all", "audio", "image", "video", "document"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SearchParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Bool(ptr + 20, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Ref(ptr + 0, x["query"]);
        A.store.Enum(ptr + 4, ["all", "audio", "image", "video", "document"].indexOf(x["category"] as string));
        A.store.Bool(ptr + 20, "modifiedTimestamp" in x ? true : false);
        A.store.Float64(ptr + 8, x["modifiedTimestamp"] === undefined ? 0 : (x["modifiedTimestamp"] as number));
        A.store.Ref(ptr + 16, x["nextFeed"]);
      }
    },
    "load_SearchParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["query"] = A.load.Ref(ptr + 0, undefined);
      x["category"] = A.load.Enum(ptr + 4, ["all", "audio", "image", "video", "document"]);
      if (A.load.Bool(ptr + 20)) {
        x["modifiedTimestamp"] = A.load.Float64(ptr + 8);
      } else {
        delete x["modifiedTimestamp"];
      }
      x["nextFeed"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_SharesheetLaunchSource": (ref: heap.Ref<string>): number => {
      const idx = ["context_menu", "sharesheet_button", "unknown"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SourceRestriction": (ref: heap.Ref<string>): number => {
      const idx = ["any_source", "native_source"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SyncState": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["fileUrl"]);
        A.store.Enum(
          ptr + 4,
          ["not_found", "queued", "in_progress", "completed", "error"].indexOf(x["syncStatus"] as string)
        );
        A.store.Bool(ptr + 16, "progress" in x ? true : false);
        A.store.Float64(ptr + 8, x["progress"] === undefined ? 0 : (x["progress"] as number));
      }
    },
    "load_SyncState": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileUrl"] = A.load.Ref(ptr + 0, undefined);
      x["syncStatus"] = A.load.Enum(ptr + 4, ["not_found", "queued", "in_progress", "completed", "error"]);
      if (A.load.Bool(ptr + 16)) {
        x["progress"] = A.load.Float64(ptr + 8);
      } else {
        delete x["progress"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ZoomOperationType": (ref: heap.Ref<string>): number => {
      const idx = ["in", "out", "reset"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_AddFileWatch": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "addFileWatch" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddFileWatch": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.addFileWatch);
    },
    "call_AddFileWatch": (retPtr: Pointer, entry: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.addFileWatch(A.H.get<object>(entry));
      A.store.Ref(retPtr, _ret);
    },
    "try_AddFileWatch": (retPtr: Pointer, errPtr: Pointer, entry: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.addFileWatch(A.H.get<object>(entry));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AddMount": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "addMount" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddMount": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.addMount);
    },
    "call_AddMount": (retPtr: Pointer, fileUrl: heap.Ref<object>, password: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.addMount(A.H.get<object>(fileUrl), A.H.get<object>(password));
      A.store.Ref(retPtr, _ret);
    },
    "try_AddMount": (
      retPtr: Pointer,
      errPtr: Pointer,
      fileUrl: heap.Ref<object>,
      password: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.addMount(A.H.get<object>(fileUrl), A.H.get<object>(password));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AddProvidedFileSystem": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "addProvidedFileSystem" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddProvidedFileSystem": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.addProvidedFileSystem);
    },
    "call_AddProvidedFileSystem": (retPtr: Pointer, providerId: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.addProvidedFileSystem(A.H.get<object>(providerId));
      A.store.Ref(retPtr, _ret);
    },
    "try_AddProvidedFileSystem": (
      retPtr: Pointer,
      errPtr: Pointer,
      providerId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.addProvidedFileSystem(A.H.get<object>(providerId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CalculateBulkPinRequiredSpace": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "calculateBulkPinRequiredSpace" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CalculateBulkPinRequiredSpace": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.calculateBulkPinRequiredSpace);
    },
    "call_CalculateBulkPinRequiredSpace": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.calculateBulkPinRequiredSpace();
      A.store.Ref(retPtr, _ret);
    },
    "try_CalculateBulkPinRequiredSpace": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.calculateBulkPinRequiredSpace();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CancelDialog": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "cancelDialog" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CancelDialog": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.cancelDialog);
    },
    "call_CancelDialog": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.cancelDialog();
    },
    "try_CancelDialog": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.cancelDialog();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CancelIOTask": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "cancelIOTask" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CancelIOTask": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.cancelIOTask);
    },
    "call_CancelIOTask": (retPtr: Pointer, taskId: number): void => {
      const _ret = WEBEXT.fileManagerPrivate.cancelIOTask(taskId);
    },
    "try_CancelIOTask": (retPtr: Pointer, errPtr: Pointer, taskId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.cancelIOTask(taskId);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CancelMounting": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "cancelMounting" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CancelMounting": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.cancelMounting);
    },
    "call_CancelMounting": (retPtr: Pointer, fileUrl: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.cancelMounting(A.H.get<object>(fileUrl));
      A.store.Ref(retPtr, _ret);
    },
    "try_CancelMounting": (retPtr: Pointer, errPtr: Pointer, fileUrl: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.cancelMounting(A.H.get<object>(fileUrl));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ComputeChecksum": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "computeChecksum" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ComputeChecksum": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.computeChecksum);
    },
    "call_ComputeChecksum": (retPtr: Pointer, entry: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.computeChecksum(A.H.get<object>(entry));
      A.store.Ref(retPtr, _ret);
    },
    "try_ComputeChecksum": (retPtr: Pointer, errPtr: Pointer, entry: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.computeChecksum(A.H.get<object>(entry));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ConfigureVolume": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "configureVolume" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ConfigureVolume": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.configureVolume);
    },
    "call_ConfigureVolume": (retPtr: Pointer, volumeId: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.configureVolume(A.H.get<object>(volumeId));
      A.store.Ref(retPtr, _ret);
    },
    "try_ConfigureVolume": (retPtr: Pointer, errPtr: Pointer, volumeId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.configureVolume(A.H.get<object>(volumeId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DismissIOTask": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "dismissIOTask" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DismissIOTask": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.dismissIOTask);
    },
    "call_DismissIOTask": (retPtr: Pointer, taskId: number): void => {
      const _ret = WEBEXT.fileManagerPrivate.dismissIOTask(taskId);
      A.store.Ref(retPtr, _ret);
    },
    "try_DismissIOTask": (retPtr: Pointer, errPtr: Pointer, taskId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.dismissIOTask(taskId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_EnableExternalFileScheme": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "enableExternalFileScheme" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EnableExternalFileScheme": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.enableExternalFileScheme);
    },
    "call_EnableExternalFileScheme": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.enableExternalFileScheme();
    },
    "try_EnableExternalFileScheme": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.enableExternalFileScheme();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ExecuteCustomAction": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "executeCustomAction" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ExecuteCustomAction": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.executeCustomAction);
    },
    "call_ExecuteCustomAction": (retPtr: Pointer, entries: heap.Ref<object>, actionId: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.executeCustomAction(A.H.get<object>(entries), A.H.get<object>(actionId));
      A.store.Ref(retPtr, _ret);
    },
    "try_ExecuteCustomAction": (
      retPtr: Pointer,
      errPtr: Pointer,
      entries: heap.Ref<object>,
      actionId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.executeCustomAction(A.H.get<object>(entries), A.H.get<object>(actionId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ExecuteTask": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "executeTask" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ExecuteTask": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.executeTask);
    },
    "call_ExecuteTask": (retPtr: Pointer, descriptor: Pointer, entries: heap.Ref<object>): void => {
      const descriptor_ffi = {};

      descriptor_ffi["appId"] = A.load.Ref(descriptor + 0, undefined);
      descriptor_ffi["taskType"] = A.load.Ref(descriptor + 4, undefined);
      descriptor_ffi["actionId"] = A.load.Ref(descriptor + 8, undefined);

      const _ret = WEBEXT.fileManagerPrivate.executeTask(descriptor_ffi, A.H.get<object>(entries));
      A.store.Ref(retPtr, _ret);
    },
    "try_ExecuteTask": (
      retPtr: Pointer,
      errPtr: Pointer,
      descriptor: Pointer,
      entries: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const descriptor_ffi = {};

        descriptor_ffi["appId"] = A.load.Ref(descriptor + 0, undefined);
        descriptor_ffi["taskType"] = A.load.Ref(descriptor + 4, undefined);
        descriptor_ffi["actionId"] = A.load.Ref(descriptor + 8, undefined);

        const _ret = WEBEXT.fileManagerPrivate.executeTask(descriptor_ffi, A.H.get<object>(entries));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_FormatVolume": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "formatVolume" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_FormatVolume": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.formatVolume);
    },
    "call_FormatVolume": (
      retPtr: Pointer,
      volumeId: heap.Ref<object>,
      filesystem: number,
      volumeLabel: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivate.formatVolume(
        A.H.get<object>(volumeId),
        filesystem > 0 && filesystem <= 3 ? ["vfat", "exfat", "ntfs"][filesystem - 1] : undefined,
        A.H.get<object>(volumeLabel)
      );
    },
    "try_FormatVolume": (
      retPtr: Pointer,
      errPtr: Pointer,
      volumeId: heap.Ref<object>,
      filesystem: number,
      volumeLabel: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.formatVolume(
          A.H.get<object>(volumeId),
          filesystem > 0 && filesystem <= 3 ? ["vfat", "exfat", "ntfs"][filesystem - 1] : undefined,
          A.H.get<object>(volumeLabel)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAndroidPickerApps": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getAndroidPickerApps" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAndroidPickerApps": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getAndroidPickerApps);
    },
    "call_GetAndroidPickerApps": (retPtr: Pointer, extensions: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.getAndroidPickerApps(A.H.get<object>(extensions));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAndroidPickerApps": (retPtr: Pointer, errPtr: Pointer, extensions: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getAndroidPickerApps(A.H.get<object>(extensions));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetBulkPinProgress": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getBulkPinProgress" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetBulkPinProgress": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getBulkPinProgress);
    },
    "call_GetBulkPinProgress": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.getBulkPinProgress();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetBulkPinProgress": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getBulkPinProgress();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetContentMetadata": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getContentMetadata" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetContentMetadata": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getContentMetadata);
    },
    "call_GetContentMetadata": (
      retPtr: Pointer,
      fileEntry: heap.Ref<object>,
      mimeType: heap.Ref<object>,
      includeImages: heap.Ref<boolean>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivate.getContentMetadata(
        A.H.get<object>(fileEntry),
        A.H.get<object>(mimeType),
        includeImages === A.H.TRUE
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_GetContentMetadata": (
      retPtr: Pointer,
      errPtr: Pointer,
      fileEntry: heap.Ref<object>,
      mimeType: heap.Ref<object>,
      includeImages: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getContentMetadata(
          A.H.get<object>(fileEntry),
          A.H.get<object>(mimeType),
          includeImages === A.H.TRUE
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetContentMimeType": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getContentMimeType" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetContentMimeType": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getContentMimeType);
    },
    "call_GetContentMimeType": (retPtr: Pointer, fileEntry: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.getContentMimeType(A.H.get<object>(fileEntry));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetContentMimeType": (retPtr: Pointer, errPtr: Pointer, fileEntry: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getContentMimeType(A.H.get<object>(fileEntry));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCrostiniSharedPaths": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getCrostiniSharedPaths" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCrostiniSharedPaths": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getCrostiniSharedPaths);
    },
    "call_GetCrostiniSharedPaths": (
      retPtr: Pointer,
      observeFirstForSession: heap.Ref<boolean>,
      vmName: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivate.getCrostiniSharedPaths(
        observeFirstForSession === A.H.TRUE,
        A.H.get<object>(vmName),
        A.H.get<object>(callback)
      );
    },
    "try_GetCrostiniSharedPaths": (
      retPtr: Pointer,
      errPtr: Pointer,
      observeFirstForSession: heap.Ref<boolean>,
      vmName: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getCrostiniSharedPaths(
          observeFirstForSession === A.H.TRUE,
          A.H.get<object>(vmName),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCustomActions": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getCustomActions" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCustomActions": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getCustomActions);
    },
    "call_GetCustomActions": (retPtr: Pointer, entries: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.getCustomActions(A.H.get<object>(entries));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCustomActions": (retPtr: Pointer, errPtr: Pointer, entries: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getCustomActions(A.H.get<object>(entries));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDeviceConnectionState": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getDeviceConnectionState" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDeviceConnectionState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getDeviceConnectionState);
    },
    "call_GetDeviceConnectionState": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.getDeviceConnectionState(A.H.get<object>(callback));
    },
    "try_GetDeviceConnectionState": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getDeviceConnectionState(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDialogCaller": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getDialogCaller" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDialogCaller": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getDialogCaller);
    },
    "call_GetDialogCaller": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.getDialogCaller();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDialogCaller": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getDialogCaller();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDirectorySize": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getDirectorySize" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDirectorySize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getDirectorySize);
    },
    "call_GetDirectorySize": (retPtr: Pointer, entry: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.getDirectorySize(A.H.get<object>(entry));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDirectorySize": (retPtr: Pointer, errPtr: Pointer, entry: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getDirectorySize(A.H.get<object>(entry));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDisallowedTransfers": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getDisallowedTransfers" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDisallowedTransfers": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getDisallowedTransfers);
    },
    "call_GetDisallowedTransfers": (
      retPtr: Pointer,
      entries: heap.Ref<object>,
      destinationEntry: heap.Ref<object>,
      isMove: heap.Ref<boolean>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivate.getDisallowedTransfers(
        A.H.get<object>(entries),
        A.H.get<object>(destinationEntry),
        isMove === A.H.TRUE
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDisallowedTransfers": (
      retPtr: Pointer,
      errPtr: Pointer,
      entries: heap.Ref<object>,
      destinationEntry: heap.Ref<object>,
      isMove: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getDisallowedTransfers(
          A.H.get<object>(entries),
          A.H.get<object>(destinationEntry),
          isMove === A.H.TRUE
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDlpBlockedComponents": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getDlpBlockedComponents" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDlpBlockedComponents": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getDlpBlockedComponents);
    },
    "call_GetDlpBlockedComponents": (retPtr: Pointer, sourceUrl: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.getDlpBlockedComponents(A.H.get<object>(sourceUrl));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDlpBlockedComponents": (
      retPtr: Pointer,
      errPtr: Pointer,
      sourceUrl: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getDlpBlockedComponents(A.H.get<object>(sourceUrl));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDlpMetadata": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getDlpMetadata" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDlpMetadata": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getDlpMetadata);
    },
    "call_GetDlpMetadata": (retPtr: Pointer, entries: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.getDlpMetadata(A.H.get<object>(entries));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDlpMetadata": (retPtr: Pointer, errPtr: Pointer, entries: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getDlpMetadata(A.H.get<object>(entries));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDlpRestrictionDetails": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getDlpRestrictionDetails" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDlpRestrictionDetails": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getDlpRestrictionDetails);
    },
    "call_GetDlpRestrictionDetails": (retPtr: Pointer, sourceUrl: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.getDlpRestrictionDetails(A.H.get<object>(sourceUrl));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDlpRestrictionDetails": (
      retPtr: Pointer,
      errPtr: Pointer,
      sourceUrl: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getDlpRestrictionDetails(A.H.get<object>(sourceUrl));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDriveConnectionState": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getDriveConnectionState" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDriveConnectionState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getDriveConnectionState);
    },
    "call_GetDriveConnectionState": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.getDriveConnectionState();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDriveConnectionState": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getDriveConnectionState();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDriveQuotaMetadata": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getDriveQuotaMetadata" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDriveQuotaMetadata": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getDriveQuotaMetadata);
    },
    "call_GetDriveQuotaMetadata": (retPtr: Pointer, entry: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.getDriveQuotaMetadata(A.H.get<object>(entry));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDriveQuotaMetadata": (retPtr: Pointer, errPtr: Pointer, entry: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getDriveQuotaMetadata(A.H.get<object>(entry));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetEntryProperties": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getEntryProperties" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetEntryProperties": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getEntryProperties);
    },
    "call_GetEntryProperties": (retPtr: Pointer, entries: heap.Ref<object>, names: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.getEntryProperties(A.H.get<object>(entries), A.H.get<object>(names));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetEntryProperties": (
      retPtr: Pointer,
      errPtr: Pointer,
      entries: heap.Ref<object>,
      names: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getEntryProperties(A.H.get<object>(entries), A.H.get<object>(names));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetFileTasks": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getFileTasks" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetFileTasks": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getFileTasks);
    },
    "call_GetFileTasks": (retPtr: Pointer, entries: heap.Ref<object>, dlpSourceUrls: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.getFileTasks(A.H.get<object>(entries), A.H.get<object>(dlpSourceUrls));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetFileTasks": (
      retPtr: Pointer,
      errPtr: Pointer,
      entries: heap.Ref<object>,
      dlpSourceUrls: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getFileTasks(A.H.get<object>(entries), A.H.get<object>(dlpSourceUrls));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetHoldingSpaceState": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getHoldingSpaceState" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetHoldingSpaceState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getHoldingSpaceState);
    },
    "call_GetHoldingSpaceState": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.getHoldingSpaceState();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetHoldingSpaceState": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getHoldingSpaceState();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetLinuxPackageInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getLinuxPackageInfo" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetLinuxPackageInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getLinuxPackageInfo);
    },
    "call_GetLinuxPackageInfo": (retPtr: Pointer, entry: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.getLinuxPackageInfo(A.H.get<object>(entry));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetLinuxPackageInfo": (retPtr: Pointer, errPtr: Pointer, entry: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getLinuxPackageInfo(A.H.get<object>(entry));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetMimeType": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getMimeType" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetMimeType": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getMimeType);
    },
    "call_GetMimeType": (retPtr: Pointer, entry: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.getMimeType(A.H.get<object>(entry));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetMimeType": (retPtr: Pointer, errPtr: Pointer, entry: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getMimeType(A.H.get<object>(entry));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPreferences": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getPreferences" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPreferences": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getPreferences);
    },
    "call_GetPreferences": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.getPreferences();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPreferences": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getPreferences();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetProfiles": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getProfiles" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetProfiles": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getProfiles);
    },
    "call_GetProfiles": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.getProfiles(A.H.get<object>(callback));
    },
    "try_GetProfiles": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getProfiles(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetProviders": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getProviders" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetProviders": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getProviders);
    },
    "call_GetProviders": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.getProviders();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetProviders": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getProviders();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetRecentFiles": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getRecentFiles" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetRecentFiles": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getRecentFiles);
    },
    "call_GetRecentFiles": (
      retPtr: Pointer,
      restriction: number,
      fileCategory: number,
      invalidateCache: heap.Ref<boolean>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivate.getRecentFiles(
        restriction > 0 && restriction <= 2 ? ["any_source", "native_source"][restriction - 1] : undefined,
        fileCategory > 0 && fileCategory <= 5
          ? ["all", "audio", "image", "video", "document"][fileCategory - 1]
          : undefined,
        invalidateCache === A.H.TRUE
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_GetRecentFiles": (
      retPtr: Pointer,
      errPtr: Pointer,
      restriction: number,
      fileCategory: number,
      invalidateCache: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getRecentFiles(
          restriction > 0 && restriction <= 2 ? ["any_source", "native_source"][restriction - 1] : undefined,
          fileCategory > 0 && fileCategory <= 5
            ? ["all", "audio", "image", "video", "document"][fileCategory - 1]
            : undefined,
          invalidateCache === A.H.TRUE
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSizeStats": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getSizeStats" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSizeStats": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getSizeStats);
    },
    "call_GetSizeStats": (retPtr: Pointer, volumeId: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.getSizeStats(A.H.get<object>(volumeId));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSizeStats": (retPtr: Pointer, errPtr: Pointer, volumeId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getSizeStats(A.H.get<object>(volumeId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetStrings": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getStrings" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetStrings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getStrings);
    },
    "call_GetStrings": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.getStrings();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetStrings": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getStrings();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetVolumeMetadataList": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getVolumeMetadataList" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetVolumeMetadataList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getVolumeMetadataList);
    },
    "call_GetVolumeMetadataList": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.getVolumeMetadataList();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetVolumeMetadataList": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.getVolumeMetadataList();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetVolumeRoot": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "getVolumeRoot" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetVolumeRoot": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.getVolumeRoot);
    },
    "call_GetVolumeRoot": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["volumeId"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 5)) {
        options_ffi["writable"] = A.load.Bool(options + 4);
      }

      const _ret = WEBEXT.fileManagerPrivate.getVolumeRoot(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetVolumeRoot": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["volumeId"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 5)) {
          options_ffi["writable"] = A.load.Bool(options + 4);
        }

        const _ret = WEBEXT.fileManagerPrivate.getVolumeRoot(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GrantAccess": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "grantAccess" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GrantAccess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.grantAccess);
    },
    "call_GrantAccess": (retPtr: Pointer, entryUrls: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.grantAccess(A.H.get<object>(entryUrls));
      A.store.Ref(retPtr, _ret);
    },
    "try_GrantAccess": (retPtr: Pointer, errPtr: Pointer, entryUrls: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.grantAccess(A.H.get<object>(entryUrls));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ImportCrostiniImage": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "importCrostiniImage" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ImportCrostiniImage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.importCrostiniImage);
    },
    "call_ImportCrostiniImage": (retPtr: Pointer, entry: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.importCrostiniImage(A.H.get<object>(entry));
    },
    "try_ImportCrostiniImage": (retPtr: Pointer, errPtr: Pointer, entry: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.importCrostiniImage(A.H.get<object>(entry));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InstallLinuxPackage": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "installLinuxPackage" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InstallLinuxPackage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.installLinuxPackage);
    },
    "call_InstallLinuxPackage": (retPtr: Pointer, entry: heap.Ref<object>, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.installLinuxPackage(A.H.get<object>(entry), A.H.get<object>(callback));
    },
    "try_InstallLinuxPackage": (
      retPtr: Pointer,
      errPtr: Pointer,
      entry: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.installLinuxPackage(A.H.get<object>(entry), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InvokeSharesheet": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "invokeSharesheet" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InvokeSharesheet": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.invokeSharesheet);
    },
    "call_InvokeSharesheet": (
      retPtr: Pointer,
      entries: heap.Ref<object>,
      launchSource: number,
      dlpSourceUrls: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivate.invokeSharesheet(
        A.H.get<object>(entries),
        launchSource > 0 && launchSource <= 3
          ? ["context_menu", "sharesheet_button", "unknown"][launchSource - 1]
          : undefined,
        A.H.get<object>(dlpSourceUrls)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_InvokeSharesheet": (
      retPtr: Pointer,
      errPtr: Pointer,
      entries: heap.Ref<object>,
      launchSource: number,
      dlpSourceUrls: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.invokeSharesheet(
          A.H.get<object>(entries),
          launchSource > 0 && launchSource <= 3
            ? ["context_menu", "sharesheet_button", "unknown"][launchSource - 1]
            : undefined,
          A.H.get<object>(dlpSourceUrls)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsTabletModeEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "isTabletModeEnabled" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsTabletModeEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.isTabletModeEnabled);
    },
    "call_IsTabletModeEnabled": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.isTabletModeEnabled();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsTabletModeEnabled": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.isTabletModeEnabled();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ListMountableGuests": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "listMountableGuests" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ListMountableGuests": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.listMountableGuests);
    },
    "call_ListMountableGuests": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.listMountableGuests();
      A.store.Ref(retPtr, _ret);
    },
    "try_ListMountableGuests": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.listMountableGuests();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_MountCrostini": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "mountCrostini" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MountCrostini": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.mountCrostini);
    },
    "call_MountCrostini": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.mountCrostini();
      A.store.Ref(retPtr, _ret);
    },
    "try_MountCrostini": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.mountCrostini();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_MountGuest": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "mountGuest" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MountGuest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.mountGuest);
    },
    "call_MountGuest": (retPtr: Pointer, id: number): void => {
      const _ret = WEBEXT.fileManagerPrivate.mountGuest(id);
      A.store.Ref(retPtr, _ret);
    },
    "try_MountGuest": (retPtr: Pointer, errPtr: Pointer, id: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.mountGuest(id);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_NotifyDriveDialogResult": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "notifyDriveDialogResult" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_NotifyDriveDialogResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.notifyDriveDialogResult);
    },
    "call_NotifyDriveDialogResult": (retPtr: Pointer, result: number): void => {
      const _ret = WEBEXT.fileManagerPrivate.notifyDriveDialogResult(
        result > 0 && result <= 4 ? ["not_displayed", "accept", "reject", "dismiss"][result - 1] : undefined
      );
    },
    "try_NotifyDriveDialogResult": (retPtr: Pointer, errPtr: Pointer, result: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.notifyDriveDialogResult(
          result > 0 && result <= 4 ? ["not_displayed", "accept", "reject", "dismiss"][result - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAppsUpdated": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate?.onAppsUpdated && "addListener" in WEBEXT?.fileManagerPrivate?.onAppsUpdated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAppsUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onAppsUpdated.addListener);
    },
    "call_OnAppsUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onAppsUpdated.addListener(A.H.get<object>(callback));
    },
    "try_OnAppsUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onAppsUpdated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAppsUpdated": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate?.onAppsUpdated && "removeListener" in WEBEXT?.fileManagerPrivate?.onAppsUpdated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAppsUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onAppsUpdated.removeListener);
    },
    "call_OffAppsUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onAppsUpdated.removeListener(A.H.get<object>(callback));
    },
    "try_OffAppsUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onAppsUpdated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAppsUpdated": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate?.onAppsUpdated && "hasListener" in WEBEXT?.fileManagerPrivate?.onAppsUpdated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAppsUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onAppsUpdated.hasListener);
    },
    "call_HasOnAppsUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onAppsUpdated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAppsUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onAppsUpdated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnBulkPinProgress": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onBulkPinProgress &&
        "addListener" in WEBEXT?.fileManagerPrivate?.onBulkPinProgress
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnBulkPinProgress": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onBulkPinProgress.addListener);
    },
    "call_OnBulkPinProgress": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onBulkPinProgress.addListener(A.H.get<object>(callback));
    },
    "try_OnBulkPinProgress": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onBulkPinProgress.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffBulkPinProgress": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onBulkPinProgress &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onBulkPinProgress
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffBulkPinProgress": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onBulkPinProgress.removeListener);
    },
    "call_OffBulkPinProgress": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onBulkPinProgress.removeListener(A.H.get<object>(callback));
    },
    "try_OffBulkPinProgress": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onBulkPinProgress.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnBulkPinProgress": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onBulkPinProgress &&
        "hasListener" in WEBEXT?.fileManagerPrivate?.onBulkPinProgress
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnBulkPinProgress": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onBulkPinProgress.hasListener);
    },
    "call_HasOnBulkPinProgress": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onBulkPinProgress.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnBulkPinProgress": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onBulkPinProgress.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCrostiniChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onCrostiniChanged &&
        "addListener" in WEBEXT?.fileManagerPrivate?.onCrostiniChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCrostiniChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onCrostiniChanged.addListener);
    },
    "call_OnCrostiniChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onCrostiniChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnCrostiniChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onCrostiniChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCrostiniChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onCrostiniChanged &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onCrostiniChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCrostiniChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onCrostiniChanged.removeListener);
    },
    "call_OffCrostiniChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onCrostiniChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffCrostiniChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onCrostiniChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCrostiniChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onCrostiniChanged &&
        "hasListener" in WEBEXT?.fileManagerPrivate?.onCrostiniChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCrostiniChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onCrostiniChanged.hasListener);
    },
    "call_HasOnCrostiniChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onCrostiniChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCrostiniChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onCrostiniChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeviceChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate?.onDeviceChanged && "addListener" in WEBEXT?.fileManagerPrivate?.onDeviceChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeviceChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDeviceChanged.addListener);
    },
    "call_OnDeviceChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDeviceChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnDeviceChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDeviceChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeviceChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDeviceChanged &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onDeviceChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeviceChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDeviceChanged.removeListener);
    },
    "call_OffDeviceChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDeviceChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeviceChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDeviceChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeviceChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate?.onDeviceChanged && "hasListener" in WEBEXT?.fileManagerPrivate?.onDeviceChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeviceChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDeviceChanged.hasListener);
    },
    "call_HasOnDeviceChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDeviceChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeviceChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDeviceChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeviceConnectionStatusChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDeviceConnectionStatusChanged &&
        "addListener" in WEBEXT?.fileManagerPrivate?.onDeviceConnectionStatusChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeviceConnectionStatusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.addListener);
    },
    "call_OnDeviceConnectionStatusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnDeviceConnectionStatusChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeviceConnectionStatusChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDeviceConnectionStatusChanged &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onDeviceConnectionStatusChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeviceConnectionStatusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.removeListener);
    },
    "call_OffDeviceConnectionStatusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeviceConnectionStatusChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeviceConnectionStatusChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDeviceConnectionStatusChanged &&
        "hasListener" in WEBEXT?.fileManagerPrivate?.onDeviceConnectionStatusChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeviceConnectionStatusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.hasListener);
    },
    "call_HasOnDeviceConnectionStatusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeviceConnectionStatusChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDeviceConnectionStatusChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDirectoryChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDirectoryChanged &&
        "addListener" in WEBEXT?.fileManagerPrivate?.onDirectoryChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDirectoryChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDirectoryChanged.addListener);
    },
    "call_OnDirectoryChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDirectoryChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnDirectoryChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDirectoryChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDirectoryChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDirectoryChanged &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onDirectoryChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDirectoryChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDirectoryChanged.removeListener);
    },
    "call_OffDirectoryChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDirectoryChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffDirectoryChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDirectoryChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDirectoryChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDirectoryChanged &&
        "hasListener" in WEBEXT?.fileManagerPrivate?.onDirectoryChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDirectoryChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDirectoryChanged.hasListener);
    },
    "call_HasOnDirectoryChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDirectoryChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDirectoryChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDirectoryChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDriveConfirmDialog": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDriveConfirmDialog &&
        "addListener" in WEBEXT?.fileManagerPrivate?.onDriveConfirmDialog
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDriveConfirmDialog": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDriveConfirmDialog.addListener);
    },
    "call_OnDriveConfirmDialog": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDriveConfirmDialog.addListener(A.H.get<object>(callback));
    },
    "try_OnDriveConfirmDialog": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDriveConfirmDialog.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDriveConfirmDialog": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDriveConfirmDialog &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onDriveConfirmDialog
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDriveConfirmDialog": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDriveConfirmDialog.removeListener);
    },
    "call_OffDriveConfirmDialog": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDriveConfirmDialog.removeListener(A.H.get<object>(callback));
    },
    "try_OffDriveConfirmDialog": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDriveConfirmDialog.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDriveConfirmDialog": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDriveConfirmDialog &&
        "hasListener" in WEBEXT?.fileManagerPrivate?.onDriveConfirmDialog
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDriveConfirmDialog": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDriveConfirmDialog.hasListener);
    },
    "call_HasOnDriveConfirmDialog": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDriveConfirmDialog.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDriveConfirmDialog": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDriveConfirmDialog.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDriveConnectionStatusChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDriveConnectionStatusChanged &&
        "addListener" in WEBEXT?.fileManagerPrivate?.onDriveConnectionStatusChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDriveConnectionStatusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.addListener);
    },
    "call_OnDriveConnectionStatusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnDriveConnectionStatusChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDriveConnectionStatusChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDriveConnectionStatusChanged &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onDriveConnectionStatusChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDriveConnectionStatusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.removeListener);
    },
    "call_OffDriveConnectionStatusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffDriveConnectionStatusChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDriveConnectionStatusChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDriveConnectionStatusChanged &&
        "hasListener" in WEBEXT?.fileManagerPrivate?.onDriveConnectionStatusChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDriveConnectionStatusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.hasListener);
    },
    "call_HasOnDriveConnectionStatusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDriveConnectionStatusChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDriveConnectionStatusChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDriveSyncError": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDriveSyncError &&
        "addListener" in WEBEXT?.fileManagerPrivate?.onDriveSyncError
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDriveSyncError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDriveSyncError.addListener);
    },
    "call_OnDriveSyncError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDriveSyncError.addListener(A.H.get<object>(callback));
    },
    "try_OnDriveSyncError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDriveSyncError.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDriveSyncError": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDriveSyncError &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onDriveSyncError
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDriveSyncError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDriveSyncError.removeListener);
    },
    "call_OffDriveSyncError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDriveSyncError.removeListener(A.H.get<object>(callback));
    },
    "try_OffDriveSyncError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDriveSyncError.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDriveSyncError": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onDriveSyncError &&
        "hasListener" in WEBEXT?.fileManagerPrivate?.onDriveSyncError
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDriveSyncError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onDriveSyncError.hasListener);
    },
    "call_HasOnDriveSyncError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onDriveSyncError.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDriveSyncError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onDriveSyncError.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnFileTransfersUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onFileTransfersUpdated &&
        "addListener" in WEBEXT?.fileManagerPrivate?.onFileTransfersUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnFileTransfersUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onFileTransfersUpdated.addListener);
    },
    "call_OnFileTransfersUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onFileTransfersUpdated.addListener(A.H.get<object>(callback));
    },
    "try_OnFileTransfersUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onFileTransfersUpdated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffFileTransfersUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onFileTransfersUpdated &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onFileTransfersUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffFileTransfersUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onFileTransfersUpdated.removeListener);
    },
    "call_OffFileTransfersUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onFileTransfersUpdated.removeListener(A.H.get<object>(callback));
    },
    "try_OffFileTransfersUpdated": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onFileTransfersUpdated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnFileTransfersUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onFileTransfersUpdated &&
        "hasListener" in WEBEXT?.fileManagerPrivate?.onFileTransfersUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnFileTransfersUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onFileTransfersUpdated.hasListener);
    },
    "call_HasOnFileTransfersUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onFileTransfersUpdated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnFileTransfersUpdated": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onFileTransfersUpdated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnIOTaskProgressStatus": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onIOTaskProgressStatus &&
        "addListener" in WEBEXT?.fileManagerPrivate?.onIOTaskProgressStatus
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnIOTaskProgressStatus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.addListener);
    },
    "call_OnIOTaskProgressStatus": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.addListener(A.H.get<object>(callback));
    },
    "try_OnIOTaskProgressStatus": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffIOTaskProgressStatus": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onIOTaskProgressStatus &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onIOTaskProgressStatus
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffIOTaskProgressStatus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.removeListener);
    },
    "call_OffIOTaskProgressStatus": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.removeListener(A.H.get<object>(callback));
    },
    "try_OffIOTaskProgressStatus": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnIOTaskProgressStatus": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onIOTaskProgressStatus &&
        "hasListener" in WEBEXT?.fileManagerPrivate?.onIOTaskProgressStatus
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnIOTaskProgressStatus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.hasListener);
    },
    "call_HasOnIOTaskProgressStatus": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnIOTaskProgressStatus": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onIOTaskProgressStatus.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnIndividualFileTransfersUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onIndividualFileTransfersUpdated &&
        "addListener" in WEBEXT?.fileManagerPrivate?.onIndividualFileTransfersUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnIndividualFileTransfersUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.addListener);
    },
    "call_OnIndividualFileTransfersUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.addListener(A.H.get<object>(callback));
    },
    "try_OnIndividualFileTransfersUpdated": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffIndividualFileTransfersUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onIndividualFileTransfersUpdated &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onIndividualFileTransfersUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffIndividualFileTransfersUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.removeListener);
    },
    "call_OffIndividualFileTransfersUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.removeListener(A.H.get<object>(callback));
    },
    "try_OffIndividualFileTransfersUpdated": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnIndividualFileTransfersUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onIndividualFileTransfersUpdated &&
        "hasListener" in WEBEXT?.fileManagerPrivate?.onIndividualFileTransfersUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnIndividualFileTransfersUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.hasListener);
    },
    "call_HasOnIndividualFileTransfersUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnIndividualFileTransfersUpdated": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onIndividualFileTransfersUpdated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMountCompleted": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onMountCompleted &&
        "addListener" in WEBEXT?.fileManagerPrivate?.onMountCompleted
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMountCompleted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onMountCompleted.addListener);
    },
    "call_OnMountCompleted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onMountCompleted.addListener(A.H.get<object>(callback));
    },
    "try_OnMountCompleted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onMountCompleted.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMountCompleted": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onMountCompleted &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onMountCompleted
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMountCompleted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onMountCompleted.removeListener);
    },
    "call_OffMountCompleted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onMountCompleted.removeListener(A.H.get<object>(callback));
    },
    "try_OffMountCompleted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onMountCompleted.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMountCompleted": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onMountCompleted &&
        "hasListener" in WEBEXT?.fileManagerPrivate?.onMountCompleted
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMountCompleted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onMountCompleted.hasListener);
    },
    "call_HasOnMountCompleted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onMountCompleted.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMountCompleted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onMountCompleted.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMountableGuestsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onMountableGuestsChanged &&
        "addListener" in WEBEXT?.fileManagerPrivate?.onMountableGuestsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMountableGuestsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onMountableGuestsChanged.addListener);
    },
    "call_OnMountableGuestsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onMountableGuestsChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnMountableGuestsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onMountableGuestsChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMountableGuestsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onMountableGuestsChanged &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onMountableGuestsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMountableGuestsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onMountableGuestsChanged.removeListener);
    },
    "call_OffMountableGuestsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onMountableGuestsChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffMountableGuestsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onMountableGuestsChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMountableGuestsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onMountableGuestsChanged &&
        "hasListener" in WEBEXT?.fileManagerPrivate?.onMountableGuestsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMountableGuestsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onMountableGuestsChanged.hasListener);
    },
    "call_HasOnMountableGuestsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onMountableGuestsChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMountableGuestsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onMountableGuestsChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPinTransfersUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onPinTransfersUpdated &&
        "addListener" in WEBEXT?.fileManagerPrivate?.onPinTransfersUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPinTransfersUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onPinTransfersUpdated.addListener);
    },
    "call_OnPinTransfersUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onPinTransfersUpdated.addListener(A.H.get<object>(callback));
    },
    "try_OnPinTransfersUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onPinTransfersUpdated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPinTransfersUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onPinTransfersUpdated &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onPinTransfersUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPinTransfersUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onPinTransfersUpdated.removeListener);
    },
    "call_OffPinTransfersUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onPinTransfersUpdated.removeListener(A.H.get<object>(callback));
    },
    "try_OffPinTransfersUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onPinTransfersUpdated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPinTransfersUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onPinTransfersUpdated &&
        "hasListener" in WEBEXT?.fileManagerPrivate?.onPinTransfersUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPinTransfersUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onPinTransfersUpdated.hasListener);
    },
    "call_HasOnPinTransfersUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onPinTransfersUpdated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPinTransfersUpdated": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onPinTransfersUpdated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPreferencesChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onPreferencesChanged &&
        "addListener" in WEBEXT?.fileManagerPrivate?.onPreferencesChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPreferencesChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onPreferencesChanged.addListener);
    },
    "call_OnPreferencesChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onPreferencesChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnPreferencesChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onPreferencesChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPreferencesChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onPreferencesChanged &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onPreferencesChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPreferencesChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onPreferencesChanged.removeListener);
    },
    "call_OffPreferencesChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onPreferencesChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffPreferencesChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onPreferencesChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPreferencesChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onPreferencesChanged &&
        "hasListener" in WEBEXT?.fileManagerPrivate?.onPreferencesChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPreferencesChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onPreferencesChanged.hasListener);
    },
    "call_HasOnPreferencesChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onPreferencesChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPreferencesChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onPreferencesChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTabletModeChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onTabletModeChanged &&
        "addListener" in WEBEXT?.fileManagerPrivate?.onTabletModeChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTabletModeChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onTabletModeChanged.addListener);
    },
    "call_OnTabletModeChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onTabletModeChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnTabletModeChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onTabletModeChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffTabletModeChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onTabletModeChanged &&
        "removeListener" in WEBEXT?.fileManagerPrivate?.onTabletModeChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffTabletModeChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onTabletModeChanged.removeListener);
    },
    "call_OffTabletModeChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onTabletModeChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffTabletModeChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onTabletModeChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnTabletModeChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileManagerPrivate?.onTabletModeChanged &&
        "hasListener" in WEBEXT?.fileManagerPrivate?.onTabletModeChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnTabletModeChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.onTabletModeChanged.hasListener);
    },
    "call_HasOnTabletModeChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.onTabletModeChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnTabletModeChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.onTabletModeChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenInspector": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "openInspector" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenInspector": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.openInspector);
    },
    "call_OpenInspector": (retPtr: Pointer, type: number): void => {
      const _ret = WEBEXT.fileManagerPrivate.openInspector(
        type > 0 && type <= 4 ? ["normal", "console", "element", "background"][type - 1] : undefined
      );
    },
    "try_OpenInspector": (retPtr: Pointer, errPtr: Pointer, type: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.openInspector(
          type > 0 && type <= 4 ? ["normal", "console", "element", "background"][type - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenManageSyncSettings": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "openManageSyncSettings" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenManageSyncSettings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.openManageSyncSettings);
    },
    "call_OpenManageSyncSettings": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.openManageSyncSettings();
    },
    "try_OpenManageSyncSettings": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.openManageSyncSettings();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenSettingsSubpage": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "openSettingsSubpage" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenSettingsSubpage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.openSettingsSubpage);
    },
    "call_OpenSettingsSubpage": (retPtr: Pointer, sub_page: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.openSettingsSubpage(A.H.get<object>(sub_page));
    },
    "try_OpenSettingsSubpage": (retPtr: Pointer, errPtr: Pointer, sub_page: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.openSettingsSubpage(A.H.get<object>(sub_page));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenURL": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "openURL" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenURL": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.openURL);
    },
    "call_OpenURL": (retPtr: Pointer, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.openURL(A.H.get<object>(url));
    },
    "try_OpenURL": (retPtr: Pointer, errPtr: Pointer, url: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.openURL(A.H.get<object>(url));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenWindow": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "openWindow" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenWindow": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.openWindow);
    },
    "call_OpenWindow": (retPtr: Pointer, params: Pointer): void => {
      const params_ffi = {};

      params_ffi["currentDirectoryURL"] = A.load.Ref(params + 0, undefined);
      params_ffi["selectionURL"] = A.load.Ref(params + 4, undefined);

      const _ret = WEBEXT.fileManagerPrivate.openWindow(params_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_OpenWindow": (retPtr: Pointer, errPtr: Pointer, params: Pointer): heap.Ref<boolean> => {
      try {
        const params_ffi = {};

        params_ffi["currentDirectoryURL"] = A.load.Ref(params + 0, undefined);
        params_ffi["selectionURL"] = A.load.Ref(params + 4, undefined);

        const _ret = WEBEXT.fileManagerPrivate.openWindow(params_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ParseTrashInfoFiles": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "parseTrashInfoFiles" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ParseTrashInfoFiles": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.parseTrashInfoFiles);
    },
    "call_ParseTrashInfoFiles": (retPtr: Pointer, entries: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.parseTrashInfoFiles(A.H.get<object>(entries));
      A.store.Ref(retPtr, _ret);
    },
    "try_ParseTrashInfoFiles": (retPtr: Pointer, errPtr: Pointer, entries: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.parseTrashInfoFiles(A.H.get<object>(entries));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_PinDriveFile": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "pinDriveFile" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_PinDriveFile": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.pinDriveFile);
    },
    "call_PinDriveFile": (retPtr: Pointer, entry: heap.Ref<object>, pin: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.fileManagerPrivate.pinDriveFile(A.H.get<object>(entry), pin === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_PinDriveFile": (
      retPtr: Pointer,
      errPtr: Pointer,
      entry: heap.Ref<object>,
      pin: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.pinDriveFile(A.H.get<object>(entry), pin === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_PollDriveHostedFilePinStates": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "pollDriveHostedFilePinStates" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_PollDriveHostedFilePinStates": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.pollDriveHostedFilePinStates);
    },
    "call_PollDriveHostedFilePinStates": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.pollDriveHostedFilePinStates();
    },
    "try_PollDriveHostedFilePinStates": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.pollDriveHostedFilePinStates();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ProgressPausedTasks": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "progressPausedTasks" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ProgressPausedTasks": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.progressPausedTasks);
    },
    "call_ProgressPausedTasks": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.progressPausedTasks();
    },
    "try_ProgressPausedTasks": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.progressPausedTasks();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveFileWatch": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "removeFileWatch" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveFileWatch": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.removeFileWatch);
    },
    "call_RemoveFileWatch": (retPtr: Pointer, entry: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.removeFileWatch(A.H.get<object>(entry));
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveFileWatch": (retPtr: Pointer, errPtr: Pointer, entry: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.removeFileWatch(A.H.get<object>(entry));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveMount": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "removeMount" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveMount": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.removeMount);
    },
    "call_RemoveMount": (retPtr: Pointer, volumeId: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.removeMount(A.H.get<object>(volumeId));
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveMount": (retPtr: Pointer, errPtr: Pointer, volumeId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.removeMount(A.H.get<object>(volumeId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RenameVolume": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "renameVolume" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RenameVolume": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.renameVolume);
    },
    "call_RenameVolume": (retPtr: Pointer, volumeId: heap.Ref<object>, newName: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.renameVolume(A.H.get<object>(volumeId), A.H.get<object>(newName));
    },
    "try_RenameVolume": (
      retPtr: Pointer,
      errPtr: Pointer,
      volumeId: heap.Ref<object>,
      newName: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.renameVolume(A.H.get<object>(volumeId), A.H.get<object>(newName));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ResolveIsolatedEntries": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "resolveIsolatedEntries" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ResolveIsolatedEntries": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.resolveIsolatedEntries);
    },
    "call_ResolveIsolatedEntries": (retPtr: Pointer, entries: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.resolveIsolatedEntries(A.H.get<object>(entries));
      A.store.Ref(retPtr, _ret);
    },
    "try_ResolveIsolatedEntries": (retPtr: Pointer, errPtr: Pointer, entries: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.resolveIsolatedEntries(A.H.get<object>(entries));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ResumeIOTask": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "resumeIOTask" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ResumeIOTask": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.resumeIOTask);
    },
    "call_ResumeIOTask": (retPtr: Pointer, taskId: number, params: Pointer): void => {
      const params_ffi = {};

      if (A.load.Bool(params + 0 + 6)) {
        params_ffi["conflictParams"] = {};
        params_ffi["conflictParams"]["conflictResolve"] = A.load.Ref(params + 0 + 0, undefined);
        if (A.load.Bool(params + 0 + 5)) {
          params_ffi["conflictParams"]["conflictApplyToAll"] = A.load.Bool(params + 0 + 4);
        }
      }
      if (A.load.Bool(params + 8 + 4)) {
        params_ffi["policyParams"] = {};
        params_ffi["policyParams"]["type"] = A.load.Enum(params + 8 + 0, [
          "dlp",
          "enterprise_connectors",
          "dlp_warning_timeout",
        ]);
      }

      const _ret = WEBEXT.fileManagerPrivate.resumeIOTask(taskId, params_ffi);
    },
    "try_ResumeIOTask": (retPtr: Pointer, errPtr: Pointer, taskId: number, params: Pointer): heap.Ref<boolean> => {
      try {
        const params_ffi = {};

        if (A.load.Bool(params + 0 + 6)) {
          params_ffi["conflictParams"] = {};
          params_ffi["conflictParams"]["conflictResolve"] = A.load.Ref(params + 0 + 0, undefined);
          if (A.load.Bool(params + 0 + 5)) {
            params_ffi["conflictParams"]["conflictApplyToAll"] = A.load.Bool(params + 0 + 4);
          }
        }
        if (A.load.Bool(params + 8 + 4)) {
          params_ffi["policyParams"] = {};
          params_ffi["policyParams"]["type"] = A.load.Enum(params + 8 + 0, [
            "dlp",
            "enterprise_connectors",
            "dlp_warning_timeout",
          ]);
        }

        const _ret = WEBEXT.fileManagerPrivate.resumeIOTask(taskId, params_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SearchDrive": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "searchDrive" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SearchDrive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.searchDrive);
    },
    "call_SearchDrive": (retPtr: Pointer, searchParams: Pointer, callback: heap.Ref<object>): void => {
      const searchParams_ffi = {};

      searchParams_ffi["query"] = A.load.Ref(searchParams + 0, undefined);
      searchParams_ffi["category"] = A.load.Enum(searchParams + 4, ["all", "audio", "image", "video", "document"]);
      if (A.load.Bool(searchParams + 20)) {
        searchParams_ffi["modifiedTimestamp"] = A.load.Float64(searchParams + 8);
      }
      searchParams_ffi["nextFeed"] = A.load.Ref(searchParams + 16, undefined);

      const _ret = WEBEXT.fileManagerPrivate.searchDrive(searchParams_ffi, A.H.get<object>(callback));
    },
    "try_SearchDrive": (
      retPtr: Pointer,
      errPtr: Pointer,
      searchParams: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const searchParams_ffi = {};

        searchParams_ffi["query"] = A.load.Ref(searchParams + 0, undefined);
        searchParams_ffi["category"] = A.load.Enum(searchParams + 4, ["all", "audio", "image", "video", "document"]);
        if (A.load.Bool(searchParams + 20)) {
          searchParams_ffi["modifiedTimestamp"] = A.load.Float64(searchParams + 8);
        }
        searchParams_ffi["nextFeed"] = A.load.Ref(searchParams + 16, undefined);

        const _ret = WEBEXT.fileManagerPrivate.searchDrive(searchParams_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SearchDriveMetadata": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "searchDriveMetadata" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SearchDriveMetadata": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.searchDriveMetadata);
    },
    "call_SearchDriveMetadata": (retPtr: Pointer, searchParams: Pointer): void => {
      const searchParams_ffi = {};

      searchParams_ffi["rootDir"] = A.load.Ref(searchParams + 0, undefined);
      searchParams_ffi["query"] = A.load.Ref(searchParams + 4, undefined);
      searchParams_ffi["types"] = A.load.Enum(searchParams + 8, [
        "EXCLUDE_DIRECTORIES",
        "SHARED_WITH_ME",
        "OFFLINE",
        "ALL",
      ]);
      if (A.load.Bool(searchParams + 28)) {
        searchParams_ffi["maxResults"] = A.load.Int32(searchParams + 12);
      }
      if (A.load.Bool(searchParams + 29)) {
        searchParams_ffi["modifiedTimestamp"] = A.load.Float64(searchParams + 16);
      }
      searchParams_ffi["category"] = A.load.Enum(searchParams + 24, ["all", "audio", "image", "video", "document"]);

      const _ret = WEBEXT.fileManagerPrivate.searchDriveMetadata(searchParams_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SearchDriveMetadata": (retPtr: Pointer, errPtr: Pointer, searchParams: Pointer): heap.Ref<boolean> => {
      try {
        const searchParams_ffi = {};

        searchParams_ffi["rootDir"] = A.load.Ref(searchParams + 0, undefined);
        searchParams_ffi["query"] = A.load.Ref(searchParams + 4, undefined);
        searchParams_ffi["types"] = A.load.Enum(searchParams + 8, [
          "EXCLUDE_DIRECTORIES",
          "SHARED_WITH_ME",
          "OFFLINE",
          "ALL",
        ]);
        if (A.load.Bool(searchParams + 28)) {
          searchParams_ffi["maxResults"] = A.load.Int32(searchParams + 12);
        }
        if (A.load.Bool(searchParams + 29)) {
          searchParams_ffi["modifiedTimestamp"] = A.load.Float64(searchParams + 16);
        }
        searchParams_ffi["category"] = A.load.Enum(searchParams + 24, ["all", "audio", "image", "video", "document"]);

        const _ret = WEBEXT.fileManagerPrivate.searchDriveMetadata(searchParams_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SearchFiles": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "searchFiles" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SearchFiles": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.searchFiles);
    },
    "call_SearchFiles": (retPtr: Pointer, searchParams: Pointer): void => {
      const searchParams_ffi = {};

      searchParams_ffi["rootDir"] = A.load.Ref(searchParams + 0, undefined);
      searchParams_ffi["query"] = A.load.Ref(searchParams + 4, undefined);
      searchParams_ffi["types"] = A.load.Enum(searchParams + 8, [
        "EXCLUDE_DIRECTORIES",
        "SHARED_WITH_ME",
        "OFFLINE",
        "ALL",
      ]);
      if (A.load.Bool(searchParams + 28)) {
        searchParams_ffi["maxResults"] = A.load.Int32(searchParams + 12);
      }
      if (A.load.Bool(searchParams + 29)) {
        searchParams_ffi["modifiedTimestamp"] = A.load.Float64(searchParams + 16);
      }
      searchParams_ffi["category"] = A.load.Enum(searchParams + 24, ["all", "audio", "image", "video", "document"]);

      const _ret = WEBEXT.fileManagerPrivate.searchFiles(searchParams_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SearchFiles": (retPtr: Pointer, errPtr: Pointer, searchParams: Pointer): heap.Ref<boolean> => {
      try {
        const searchParams_ffi = {};

        searchParams_ffi["rootDir"] = A.load.Ref(searchParams + 0, undefined);
        searchParams_ffi["query"] = A.load.Ref(searchParams + 4, undefined);
        searchParams_ffi["types"] = A.load.Enum(searchParams + 8, [
          "EXCLUDE_DIRECTORIES",
          "SHARED_WITH_ME",
          "OFFLINE",
          "ALL",
        ]);
        if (A.load.Bool(searchParams + 28)) {
          searchParams_ffi["maxResults"] = A.load.Int32(searchParams + 12);
        }
        if (A.load.Bool(searchParams + 29)) {
          searchParams_ffi["modifiedTimestamp"] = A.load.Float64(searchParams + 16);
        }
        searchParams_ffi["category"] = A.load.Enum(searchParams + 24, ["all", "audio", "image", "video", "document"]);

        const _ret = WEBEXT.fileManagerPrivate.searchFiles(searchParams_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SearchFilesByHashes": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "searchFilesByHashes" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SearchFilesByHashes": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.searchFilesByHashes);
    },
    "call_SearchFilesByHashes": (retPtr: Pointer, volumeId: heap.Ref<object>, hashList: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.searchFilesByHashes(A.H.get<object>(volumeId), A.H.get<object>(hashList));
      A.store.Ref(retPtr, _ret);
    },
    "try_SearchFilesByHashes": (
      retPtr: Pointer,
      errPtr: Pointer,
      volumeId: heap.Ref<object>,
      hashList: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.searchFilesByHashes(
          A.H.get<object>(volumeId),
          A.H.get<object>(hashList)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SelectAndroidPickerApp": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "selectAndroidPickerApp" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SelectAndroidPickerApp": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.selectAndroidPickerApp);
    },
    "call_SelectAndroidPickerApp": (retPtr: Pointer, androidApp: Pointer): void => {
      const androidApp_ffi = {};

      androidApp_ffi["name"] = A.load.Ref(androidApp + 0, undefined);
      androidApp_ffi["packageName"] = A.load.Ref(androidApp + 4, undefined);
      androidApp_ffi["activityName"] = A.load.Ref(androidApp + 8, undefined);
      if (A.load.Bool(androidApp + 12 + 8)) {
        androidApp_ffi["iconSet"] = {};
        androidApp_ffi["iconSet"]["icon16x16Url"] = A.load.Ref(androidApp + 12 + 0, undefined);
        androidApp_ffi["iconSet"]["icon32x32Url"] = A.load.Ref(androidApp + 12 + 4, undefined);
      }

      const _ret = WEBEXT.fileManagerPrivate.selectAndroidPickerApp(androidApp_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SelectAndroidPickerApp": (retPtr: Pointer, errPtr: Pointer, androidApp: Pointer): heap.Ref<boolean> => {
      try {
        const androidApp_ffi = {};

        androidApp_ffi["name"] = A.load.Ref(androidApp + 0, undefined);
        androidApp_ffi["packageName"] = A.load.Ref(androidApp + 4, undefined);
        androidApp_ffi["activityName"] = A.load.Ref(androidApp + 8, undefined);
        if (A.load.Bool(androidApp + 12 + 8)) {
          androidApp_ffi["iconSet"] = {};
          androidApp_ffi["iconSet"]["icon16x16Url"] = A.load.Ref(androidApp + 12 + 0, undefined);
          androidApp_ffi["iconSet"]["icon32x32Url"] = A.load.Ref(androidApp + 12 + 4, undefined);
        }

        const _ret = WEBEXT.fileManagerPrivate.selectAndroidPickerApp(androidApp_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SelectFile": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "selectFile" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SelectFile": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.selectFile);
    },
    "call_SelectFile": (
      retPtr: Pointer,
      selectedPath: heap.Ref<object>,
      index: number,
      forOpening: heap.Ref<boolean>,
      shouldReturnLocalPath: heap.Ref<boolean>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivate.selectFile(
        A.H.get<object>(selectedPath),
        index,
        forOpening === A.H.TRUE,
        shouldReturnLocalPath === A.H.TRUE
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_SelectFile": (
      retPtr: Pointer,
      errPtr: Pointer,
      selectedPath: heap.Ref<object>,
      index: number,
      forOpening: heap.Ref<boolean>,
      shouldReturnLocalPath: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.selectFile(
          A.H.get<object>(selectedPath),
          index,
          forOpening === A.H.TRUE,
          shouldReturnLocalPath === A.H.TRUE
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SelectFiles": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "selectFiles" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SelectFiles": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.selectFiles);
    },
    "call_SelectFiles": (
      retPtr: Pointer,
      selectedPaths: heap.Ref<object>,
      shouldReturnLocalPath: heap.Ref<boolean>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivate.selectFiles(
        A.H.get<object>(selectedPaths),
        shouldReturnLocalPath === A.H.TRUE
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_SelectFiles": (
      retPtr: Pointer,
      errPtr: Pointer,
      selectedPaths: heap.Ref<object>,
      shouldReturnLocalPath: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.selectFiles(
          A.H.get<object>(selectedPaths),
          shouldReturnLocalPath === A.H.TRUE
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendFeedback": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "sendFeedback" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendFeedback": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.sendFeedback);
    },
    "call_SendFeedback": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileManagerPrivate.sendFeedback();
    },
    "try_SendFeedback": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.sendFeedback();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetDefaultTask": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "setDefaultTask" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetDefaultTask": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.setDefaultTask);
    },
    "call_SetDefaultTask": (
      retPtr: Pointer,
      descriptor: Pointer,
      entries: heap.Ref<object>,
      mimeTypes: heap.Ref<object>
    ): void => {
      const descriptor_ffi = {};

      descriptor_ffi["appId"] = A.load.Ref(descriptor + 0, undefined);
      descriptor_ffi["taskType"] = A.load.Ref(descriptor + 4, undefined);
      descriptor_ffi["actionId"] = A.load.Ref(descriptor + 8, undefined);

      const _ret = WEBEXT.fileManagerPrivate.setDefaultTask(
        descriptor_ffi,
        A.H.get<object>(entries),
        A.H.get<object>(mimeTypes)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_SetDefaultTask": (
      retPtr: Pointer,
      errPtr: Pointer,
      descriptor: Pointer,
      entries: heap.Ref<object>,
      mimeTypes: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const descriptor_ffi = {};

        descriptor_ffi["appId"] = A.load.Ref(descriptor + 0, undefined);
        descriptor_ffi["taskType"] = A.load.Ref(descriptor + 4, undefined);
        descriptor_ffi["actionId"] = A.load.Ref(descriptor + 8, undefined);

        const _ret = WEBEXT.fileManagerPrivate.setDefaultTask(
          descriptor_ffi,
          A.H.get<object>(entries),
          A.H.get<object>(mimeTypes)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPreferences": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "setPreferences" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPreferences": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.setPreferences);
    },
    "call_SetPreferences": (retPtr: Pointer, changeInfo: Pointer): void => {
      const changeInfo_ffi = {};

      if (A.load.Bool(changeInfo + 9)) {
        changeInfo_ffi["driveSyncEnabledOnMeteredNetwork"] = A.load.Bool(changeInfo + 0);
      }
      if (A.load.Bool(changeInfo + 10)) {
        changeInfo_ffi["arcEnabled"] = A.load.Bool(changeInfo + 1);
      }
      if (A.load.Bool(changeInfo + 11)) {
        changeInfo_ffi["arcRemovableMediaAccessEnabled"] = A.load.Bool(changeInfo + 2);
      }
      changeInfo_ffi["folderShortcuts"] = A.load.Ref(changeInfo + 4, undefined);
      if (A.load.Bool(changeInfo + 12)) {
        changeInfo_ffi["driveFsBulkPinningEnabled"] = A.load.Bool(changeInfo + 8);
      }

      const _ret = WEBEXT.fileManagerPrivate.setPreferences(changeInfo_ffi);
    },
    "try_SetPreferences": (retPtr: Pointer, errPtr: Pointer, changeInfo: Pointer): heap.Ref<boolean> => {
      try {
        const changeInfo_ffi = {};

        if (A.load.Bool(changeInfo + 9)) {
          changeInfo_ffi["driveSyncEnabledOnMeteredNetwork"] = A.load.Bool(changeInfo + 0);
        }
        if (A.load.Bool(changeInfo + 10)) {
          changeInfo_ffi["arcEnabled"] = A.load.Bool(changeInfo + 1);
        }
        if (A.load.Bool(changeInfo + 11)) {
          changeInfo_ffi["arcRemovableMediaAccessEnabled"] = A.load.Bool(changeInfo + 2);
        }
        changeInfo_ffi["folderShortcuts"] = A.load.Ref(changeInfo + 4, undefined);
        if (A.load.Bool(changeInfo + 12)) {
          changeInfo_ffi["driveFsBulkPinningEnabled"] = A.load.Bool(changeInfo + 8);
        }

        const _ret = WEBEXT.fileManagerPrivate.setPreferences(changeInfo_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SharePathsWithCrostini": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "sharePathsWithCrostini" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SharePathsWithCrostini": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.sharePathsWithCrostini);
    },
    "call_SharePathsWithCrostini": (
      retPtr: Pointer,
      vmName: heap.Ref<object>,
      entries: heap.Ref<object>,
      persist: heap.Ref<boolean>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivate.sharePathsWithCrostini(
        A.H.get<object>(vmName),
        A.H.get<object>(entries),
        persist === A.H.TRUE
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_SharePathsWithCrostini": (
      retPtr: Pointer,
      errPtr: Pointer,
      vmName: heap.Ref<object>,
      entries: heap.Ref<object>,
      persist: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.sharePathsWithCrostini(
          A.H.get<object>(vmName),
          A.H.get<object>(entries),
          persist === A.H.TRUE
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SharesheetHasTargets": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "sharesheetHasTargets" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SharesheetHasTargets": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.sharesheetHasTargets);
    },
    "call_SharesheetHasTargets": (retPtr: Pointer, entries: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.sharesheetHasTargets(A.H.get<object>(entries));
      A.store.Ref(retPtr, _ret);
    },
    "try_SharesheetHasTargets": (retPtr: Pointer, errPtr: Pointer, entries: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.sharesheetHasTargets(A.H.get<object>(entries));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ShowPolicyDialog": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "showPolicyDialog" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ShowPolicyDialog": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.showPolicyDialog);
    },
    "call_ShowPolicyDialog": (retPtr: Pointer, taskId: number, type: number): void => {
      const _ret = WEBEXT.fileManagerPrivate.showPolicyDialog(
        taskId,
        type > 0 && type <= 2 ? ["warning", "error"][type - 1] : undefined
      );
    },
    "try_ShowPolicyDialog": (retPtr: Pointer, errPtr: Pointer, taskId: number, type: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.showPolicyDialog(
          taskId,
          type > 0 && type <= 2 ? ["warning", "error"][type - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SinglePartitionFormat": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "singlePartitionFormat" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SinglePartitionFormat": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.singlePartitionFormat);
    },
    "call_SinglePartitionFormat": (
      retPtr: Pointer,
      deviceStoragePath: heap.Ref<object>,
      filesystem: number,
      volumeLabel: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivate.singlePartitionFormat(
        A.H.get<object>(deviceStoragePath),
        filesystem > 0 && filesystem <= 3 ? ["vfat", "exfat", "ntfs"][filesystem - 1] : undefined,
        A.H.get<object>(volumeLabel)
      );
    },
    "try_SinglePartitionFormat": (
      retPtr: Pointer,
      errPtr: Pointer,
      deviceStoragePath: heap.Ref<object>,
      filesystem: number,
      volumeLabel: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.singlePartitionFormat(
          A.H.get<object>(deviceStoragePath),
          filesystem > 0 && filesystem <= 3 ? ["vfat", "exfat", "ntfs"][filesystem - 1] : undefined,
          A.H.get<object>(volumeLabel)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartIOTask": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "startIOTask" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartIOTask": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.startIOTask);
    },
    "call_StartIOTask": (retPtr: Pointer, type: number, entries: heap.Ref<object>, params: Pointer): void => {
      const params_ffi = {};

      params_ffi["destinationFolder"] = A.load.Ref(params + 0, undefined);
      params_ffi["password"] = A.load.Ref(params + 4, undefined);
      if (A.load.Bool(params + 9)) {
        params_ffi["showNotification"] = A.load.Bool(params + 8);
      }

      const _ret = WEBEXT.fileManagerPrivate.startIOTask(
        type > 0 && type <= 9
          ? ["copy", "delete", "empty_trash", "extract", "move", "restore", "restore_to_destination", "trash", "zip"][
              type - 1
            ]
          : undefined,
        A.H.get<object>(entries),
        params_ffi
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_StartIOTask": (
      retPtr: Pointer,
      errPtr: Pointer,
      type: number,
      entries: heap.Ref<object>,
      params: Pointer
    ): heap.Ref<boolean> => {
      try {
        const params_ffi = {};

        params_ffi["destinationFolder"] = A.load.Ref(params + 0, undefined);
        params_ffi["password"] = A.load.Ref(params + 4, undefined);
        if (A.load.Bool(params + 9)) {
          params_ffi["showNotification"] = A.load.Bool(params + 8);
        }

        const _ret = WEBEXT.fileManagerPrivate.startIOTask(
          type > 0 && type <= 9
            ? ["copy", "delete", "empty_trash", "extract", "move", "restore", "restore_to_destination", "trash", "zip"][
                type - 1
              ]
            : undefined,
          A.H.get<object>(entries),
          params_ffi
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ToggleAddedToHoldingSpace": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "toggleAddedToHoldingSpace" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ToggleAddedToHoldingSpace": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.toggleAddedToHoldingSpace);
    },
    "call_ToggleAddedToHoldingSpace": (retPtr: Pointer, entries: heap.Ref<object>, added: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.fileManagerPrivate.toggleAddedToHoldingSpace(A.H.get<object>(entries), added === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_ToggleAddedToHoldingSpace": (
      retPtr: Pointer,
      errPtr: Pointer,
      entries: heap.Ref<object>,
      added: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.toggleAddedToHoldingSpace(A.H.get<object>(entries), added === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UnsharePathWithCrostini": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "unsharePathWithCrostini" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UnsharePathWithCrostini": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.unsharePathWithCrostini);
    },
    "call_UnsharePathWithCrostini": (retPtr: Pointer, vmName: heap.Ref<object>, entry: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.unsharePathWithCrostini(A.H.get<object>(vmName), A.H.get<object>(entry));
      A.store.Ref(retPtr, _ret);
    },
    "try_UnsharePathWithCrostini": (
      retPtr: Pointer,
      errPtr: Pointer,
      vmName: heap.Ref<object>,
      entry: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.unsharePathWithCrostini(A.H.get<object>(vmName), A.H.get<object>(entry));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ValidatePathNameLength": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "validatePathNameLength" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ValidatePathNameLength": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.validatePathNameLength);
    },
    "call_ValidatePathNameLength": (retPtr: Pointer, parentEntry: heap.Ref<object>, name: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivate.validatePathNameLength(
        A.H.get<object>(parentEntry),
        A.H.get<object>(name)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_ValidatePathNameLength": (
      retPtr: Pointer,
      errPtr: Pointer,
      parentEntry: heap.Ref<object>,
      name: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.validatePathNameLength(
          A.H.get<object>(parentEntry),
          A.H.get<object>(name)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Zoom": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivate && "zoom" in WEBEXT?.fileManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Zoom": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivate.zoom);
    },
    "call_Zoom": (retPtr: Pointer, operation: number): void => {
      const _ret = WEBEXT.fileManagerPrivate.zoom(
        operation > 0 && operation <= 3 ? ["in", "out", "reset"][operation - 1] : undefined
      );
    },
    "try_Zoom": (retPtr: Pointer, errPtr: Pointer, operation: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivate.zoom(
          operation > 0 && operation <= 3 ? ["in", "out", "reset"][operation - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
