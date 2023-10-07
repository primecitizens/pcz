import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/filemanagerprivateinternal", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_EntryDescription": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Ref(ptr + 0, x["fileSystemName"]);
        A.store.Ref(ptr + 4, x["fileSystemRoot"]);
        A.store.Ref(ptr + 8, x["fileFullPath"]);
        A.store.Bool(ptr + 13, "fileIsDirectory" in x ? true : false);
        A.store.Bool(ptr + 12, x["fileIsDirectory"] ? true : false);
      }
    },
    "load_EntryDescription": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemName"] = A.load.Ref(ptr + 0, undefined);
      x["fileSystemRoot"] = A.load.Ref(ptr + 4, undefined);
      x["fileFullPath"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 13)) {
        x["fileIsDirectory"] = A.load.Bool(ptr + 12);
      } else {
        delete x["fileIsDirectory"];
      }
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
        A.store.Ref(ptr + 0, x["destinationFolderUrl"]);
        A.store.Ref(ptr + 4, x["password"]);
        A.store.Bool(ptr + 9, "showNotification" in x ? true : false);
        A.store.Bool(ptr + 8, x["showNotification"] ? true : false);
      }
    },
    "load_IOTaskParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["destinationFolderUrl"] = A.load.Ref(ptr + 0, undefined);
      x["password"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 9)) {
        x["showNotification"] = A.load.Bool(ptr + 8);
      } else {
        delete x["showNotification"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ParsedTrashInfoFile": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 33, false);

        A.store.Bool(ptr + 0 + 14, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Ref(ptr + 0 + 8, undefined);
        A.store.Bool(ptr + 0 + 13, false);
        A.store.Bool(ptr + 0 + 12, false);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 32, false);
        A.store.Float64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 33, true);

        if (typeof x["restoreEntry"] === "undefined") {
          A.store.Bool(ptr + 0 + 14, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Ref(ptr + 0 + 8, undefined);
          A.store.Bool(ptr + 0 + 13, false);
          A.store.Bool(ptr + 0 + 12, false);
        } else {
          A.store.Bool(ptr + 0 + 14, true);
          A.store.Ref(ptr + 0 + 0, x["restoreEntry"]["fileSystemName"]);
          A.store.Ref(ptr + 0 + 4, x["restoreEntry"]["fileSystemRoot"]);
          A.store.Ref(ptr + 0 + 8, x["restoreEntry"]["fileFullPath"]);
          A.store.Bool(ptr + 0 + 13, "fileIsDirectory" in x["restoreEntry"] ? true : false);
          A.store.Bool(ptr + 0 + 12, x["restoreEntry"]["fileIsDirectory"] ? true : false);
        }
        A.store.Ref(ptr + 16, x["trashInfoFileName"]);
        A.store.Bool(ptr + 32, "deletionDate" in x ? true : false);
        A.store.Float64(ptr + 24, x["deletionDate"] === undefined ? 0 : (x["deletionDate"] as number));
      }
    },
    "load_ParsedTrashInfoFile": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 14)) {
        x["restoreEntry"] = {};
        x["restoreEntry"]["fileSystemName"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["restoreEntry"]["fileSystemRoot"] = A.load.Ref(ptr + 0 + 4, undefined);
        x["restoreEntry"]["fileFullPath"] = A.load.Ref(ptr + 0 + 8, undefined);
        if (A.load.Bool(ptr + 0 + 13)) {
          x["restoreEntry"]["fileIsDirectory"] = A.load.Bool(ptr + 0 + 12);
        } else {
          delete x["restoreEntry"]["fileIsDirectory"];
        }
      } else {
        delete x["restoreEntry"];
      }
      x["trashInfoFileName"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 32)) {
        x["deletionDate"] = A.load.Float64(ptr + 24);
      } else {
        delete x["deletionDate"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SearchFilesParams": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Ref(ptr + 0, x["rootUrl"]);
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
    "load_SearchFilesParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["rootUrl"] = A.load.Ref(ptr + 0, undefined);
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
    "has_AddFileWatch": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "addFileWatch" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddFileWatch": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.addFileWatch);
    },
    "call_AddFileWatch": (retPtr: Pointer, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.addFileWatch(A.H.get<object>(url));
      A.store.Ref(retPtr, _ret);
    },
    "try_AddFileWatch": (retPtr: Pointer, errPtr: Pointer, url: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.addFileWatch(A.H.get<object>(url));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ComputeChecksum": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "computeChecksum" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ComputeChecksum": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.computeChecksum);
    },
    "call_ComputeChecksum": (retPtr: Pointer, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.computeChecksum(A.H.get<object>(url));
      A.store.Ref(retPtr, _ret);
    },
    "try_ComputeChecksum": (retPtr: Pointer, errPtr: Pointer, url: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.computeChecksum(A.H.get<object>(url));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ExecuteCustomAction": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "executeCustomAction" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ExecuteCustomAction": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.executeCustomAction);
    },
    "call_ExecuteCustomAction": (retPtr: Pointer, urls: heap.Ref<object>, actionId: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.executeCustomAction(
        A.H.get<object>(urls),
        A.H.get<object>(actionId)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_ExecuteCustomAction": (
      retPtr: Pointer,
      errPtr: Pointer,
      urls: heap.Ref<object>,
      actionId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.executeCustomAction(
          A.H.get<object>(urls),
          A.H.get<object>(actionId)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ExecuteTask": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "executeTask" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ExecuteTask": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.executeTask);
    },
    "call_ExecuteTask": (retPtr: Pointer, descriptor: Pointer, urls: heap.Ref<object>): void => {
      const descriptor_ffi = {};

      descriptor_ffi["appId"] = A.load.Ref(descriptor + 0, undefined);
      descriptor_ffi["taskType"] = A.load.Ref(descriptor + 4, undefined);
      descriptor_ffi["actionId"] = A.load.Ref(descriptor + 8, undefined);

      const _ret = WEBEXT.fileManagerPrivateInternal.executeTask(descriptor_ffi, A.H.get<object>(urls));
      A.store.Ref(retPtr, _ret);
    },
    "try_ExecuteTask": (
      retPtr: Pointer,
      errPtr: Pointer,
      descriptor: Pointer,
      urls: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const descriptor_ffi = {};

        descriptor_ffi["appId"] = A.load.Ref(descriptor + 0, undefined);
        descriptor_ffi["taskType"] = A.load.Ref(descriptor + 4, undefined);
        descriptor_ffi["actionId"] = A.load.Ref(descriptor + 8, undefined);

        const _ret = WEBEXT.fileManagerPrivateInternal.executeTask(descriptor_ffi, A.H.get<object>(urls));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetContentMetadata": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "getContentMetadata" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetContentMetadata": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.getContentMetadata);
    },
    "call_GetContentMetadata": (
      retPtr: Pointer,
      blobUUID: heap.Ref<object>,
      mimeType: heap.Ref<object>,
      includeImages: heap.Ref<boolean>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.getContentMetadata(
        A.H.get<object>(blobUUID),
        A.H.get<object>(mimeType),
        includeImages === A.H.TRUE
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_GetContentMetadata": (
      retPtr: Pointer,
      errPtr: Pointer,
      blobUUID: heap.Ref<object>,
      mimeType: heap.Ref<object>,
      includeImages: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.getContentMetadata(
          A.H.get<object>(blobUUID),
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
      if (WEBEXT?.fileManagerPrivateInternal && "getContentMimeType" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetContentMimeType": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.getContentMimeType);
    },
    "call_GetContentMimeType": (retPtr: Pointer, blobUUID: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.getContentMimeType(A.H.get<object>(blobUUID));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetContentMimeType": (retPtr: Pointer, errPtr: Pointer, blobUUID: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.getContentMimeType(A.H.get<object>(blobUUID));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCrostiniSharedPaths": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "getCrostiniSharedPaths" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCrostiniSharedPaths": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.getCrostiniSharedPaths);
    },
    "call_GetCrostiniSharedPaths": (
      retPtr: Pointer,
      observeFirstForSession: heap.Ref<boolean>,
      vmName: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.getCrostiniSharedPaths(
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
        const _ret = WEBEXT.fileManagerPrivateInternal.getCrostiniSharedPaths(
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
      if (WEBEXT?.fileManagerPrivateInternal && "getCustomActions" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCustomActions": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.getCustomActions);
    },
    "call_GetCustomActions": (retPtr: Pointer, urls: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.getCustomActions(A.H.get<object>(urls));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCustomActions": (retPtr: Pointer, errPtr: Pointer, urls: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.getCustomActions(A.H.get<object>(urls));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDirectorySize": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "getDirectorySize" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDirectorySize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.getDirectorySize);
    },
    "call_GetDirectorySize": (retPtr: Pointer, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.getDirectorySize(A.H.get<object>(url));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDirectorySize": (retPtr: Pointer, errPtr: Pointer, url: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.getDirectorySize(A.H.get<object>(url));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDisallowedTransfers": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "getDisallowedTransfers" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDisallowedTransfers": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.getDisallowedTransfers);
    },
    "call_GetDisallowedTransfers": (
      retPtr: Pointer,
      entries: heap.Ref<object>,
      destinationEntry: heap.Ref<object>,
      isMove: heap.Ref<boolean>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.getDisallowedTransfers(
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
        const _ret = WEBEXT.fileManagerPrivateInternal.getDisallowedTransfers(
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
    "has_GetDlpMetadata": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "getDlpMetadata" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDlpMetadata": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.getDlpMetadata);
    },
    "call_GetDlpMetadata": (retPtr: Pointer, entries: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.getDlpMetadata(A.H.get<object>(entries));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDlpMetadata": (retPtr: Pointer, errPtr: Pointer, entries: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.getDlpMetadata(A.H.get<object>(entries));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDriveQuotaMetadata": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "getDriveQuotaMetadata" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDriveQuotaMetadata": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.getDriveQuotaMetadata);
    },
    "call_GetDriveQuotaMetadata": (retPtr: Pointer, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.getDriveQuotaMetadata(A.H.get<object>(url));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDriveQuotaMetadata": (retPtr: Pointer, errPtr: Pointer, url: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.getDriveQuotaMetadata(A.H.get<object>(url));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetEntryProperties": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "getEntryProperties" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetEntryProperties": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.getEntryProperties);
    },
    "call_GetEntryProperties": (retPtr: Pointer, urls: heap.Ref<object>, names: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.getEntryProperties(A.H.get<object>(urls), A.H.get<object>(names));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetEntryProperties": (
      retPtr: Pointer,
      errPtr: Pointer,
      urls: heap.Ref<object>,
      names: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.getEntryProperties(
          A.H.get<object>(urls),
          A.H.get<object>(names)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetFileTasks": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "getFileTasks" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetFileTasks": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.getFileTasks);
    },
    "call_GetFileTasks": (retPtr: Pointer, urls: heap.Ref<object>, dlpSourceUrls: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.getFileTasks(
        A.H.get<object>(urls),
        A.H.get<object>(dlpSourceUrls)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_GetFileTasks": (
      retPtr: Pointer,
      errPtr: Pointer,
      urls: heap.Ref<object>,
      dlpSourceUrls: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.getFileTasks(
          A.H.get<object>(urls),
          A.H.get<object>(dlpSourceUrls)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetLinuxPackageInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "getLinuxPackageInfo" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetLinuxPackageInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.getLinuxPackageInfo);
    },
    "call_GetLinuxPackageInfo": (retPtr: Pointer, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.getLinuxPackageInfo(A.H.get<object>(url));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetLinuxPackageInfo": (retPtr: Pointer, errPtr: Pointer, url: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.getLinuxPackageInfo(A.H.get<object>(url));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetMimeType": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "getMimeType" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetMimeType": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.getMimeType);
    },
    "call_GetMimeType": (retPtr: Pointer, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.getMimeType(A.H.get<object>(url));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetMimeType": (retPtr: Pointer, errPtr: Pointer, url: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.getMimeType(A.H.get<object>(url));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetRecentFiles": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "getRecentFiles" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetRecentFiles": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.getRecentFiles);
    },
    "call_GetRecentFiles": (
      retPtr: Pointer,
      restriction: number,
      file_category: number,
      invalidate_cache: heap.Ref<boolean>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.getRecentFiles(
        restriction > 0 && restriction <= 2 ? ["any_source", "native_source"][restriction - 1] : undefined,
        file_category > 0 && file_category <= 5
          ? ["all", "audio", "image", "video", "document"][file_category - 1]
          : undefined,
        invalidate_cache === A.H.TRUE
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_GetRecentFiles": (
      retPtr: Pointer,
      errPtr: Pointer,
      restriction: number,
      file_category: number,
      invalidate_cache: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.getRecentFiles(
          restriction > 0 && restriction <= 2 ? ["any_source", "native_source"][restriction - 1] : undefined,
          file_category > 0 && file_category <= 5
            ? ["all", "audio", "image", "video", "document"][file_category - 1]
            : undefined,
          invalidate_cache === A.H.TRUE
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetVolumeRoot": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "getVolumeRoot" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetVolumeRoot": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.getVolumeRoot);
    },
    "call_GetVolumeRoot": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["volumeId"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 5)) {
        options_ffi["writable"] = A.load.Bool(options + 4);
      }

      const _ret = WEBEXT.fileManagerPrivateInternal.getVolumeRoot(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetVolumeRoot": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["volumeId"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 5)) {
          options_ffi["writable"] = A.load.Bool(options + 4);
        }

        const _ret = WEBEXT.fileManagerPrivateInternal.getVolumeRoot(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ImportCrostiniImage": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "importCrostiniImage" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ImportCrostiniImage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.importCrostiniImage);
    },
    "call_ImportCrostiniImage": (retPtr: Pointer, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.importCrostiniImage(A.H.get<object>(url));
    },
    "try_ImportCrostiniImage": (retPtr: Pointer, errPtr: Pointer, url: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.importCrostiniImage(A.H.get<object>(url));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InstallLinuxPackage": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "installLinuxPackage" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InstallLinuxPackage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.installLinuxPackage);
    },
    "call_InstallLinuxPackage": (retPtr: Pointer, url: heap.Ref<object>, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.installLinuxPackage(
        A.H.get<object>(url),
        A.H.get<object>(callback)
      );
    },
    "try_InstallLinuxPackage": (
      retPtr: Pointer,
      errPtr: Pointer,
      url: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.installLinuxPackage(
          A.H.get<object>(url),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InvokeSharesheet": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "invokeSharesheet" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InvokeSharesheet": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.invokeSharesheet);
    },
    "call_InvokeSharesheet": (
      retPtr: Pointer,
      urls: heap.Ref<object>,
      launchSource: number,
      dlpSourceUrls: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.invokeSharesheet(
        A.H.get<object>(urls),
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
      urls: heap.Ref<object>,
      launchSource: number,
      dlpSourceUrls: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.invokeSharesheet(
          A.H.get<object>(urls),
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
    "has_ParseTrashInfoFiles": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "parseTrashInfoFiles" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ParseTrashInfoFiles": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.parseTrashInfoFiles);
    },
    "call_ParseTrashInfoFiles": (retPtr: Pointer, urls: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.parseTrashInfoFiles(A.H.get<object>(urls));
      A.store.Ref(retPtr, _ret);
    },
    "try_ParseTrashInfoFiles": (retPtr: Pointer, errPtr: Pointer, urls: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.parseTrashInfoFiles(A.H.get<object>(urls));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_PinDriveFile": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "pinDriveFile" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_PinDriveFile": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.pinDriveFile);
    },
    "call_PinDriveFile": (retPtr: Pointer, url: heap.Ref<object>, pin: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.pinDriveFile(A.H.get<object>(url), pin === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_PinDriveFile": (
      retPtr: Pointer,
      errPtr: Pointer,
      url: heap.Ref<object>,
      pin: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.pinDriveFile(A.H.get<object>(url), pin === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveFileWatch": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "removeFileWatch" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveFileWatch": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.removeFileWatch);
    },
    "call_RemoveFileWatch": (retPtr: Pointer, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.removeFileWatch(A.H.get<object>(url));
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveFileWatch": (retPtr: Pointer, errPtr: Pointer, url: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.removeFileWatch(A.H.get<object>(url));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ResolveIsolatedEntries": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "resolveIsolatedEntries" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ResolveIsolatedEntries": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.resolveIsolatedEntries);
    },
    "call_ResolveIsolatedEntries": (retPtr: Pointer, urls: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.resolveIsolatedEntries(A.H.get<object>(urls));
      A.store.Ref(retPtr, _ret);
    },
    "try_ResolveIsolatedEntries": (retPtr: Pointer, errPtr: Pointer, urls: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.resolveIsolatedEntries(A.H.get<object>(urls));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SearchFiles": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "searchFiles" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SearchFiles": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.searchFiles);
    },
    "call_SearchFiles": (retPtr: Pointer, searchParams: Pointer): void => {
      const searchParams_ffi = {};

      searchParams_ffi["rootUrl"] = A.load.Ref(searchParams + 0, undefined);
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

      const _ret = WEBEXT.fileManagerPrivateInternal.searchFiles(searchParams_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SearchFiles": (retPtr: Pointer, errPtr: Pointer, searchParams: Pointer): heap.Ref<boolean> => {
      try {
        const searchParams_ffi = {};

        searchParams_ffi["rootUrl"] = A.load.Ref(searchParams + 0, undefined);
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

        const _ret = WEBEXT.fileManagerPrivateInternal.searchFiles(searchParams_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetDefaultTask": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "setDefaultTask" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetDefaultTask": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.setDefaultTask);
    },
    "call_SetDefaultTask": (
      retPtr: Pointer,
      descriptor: Pointer,
      urls: heap.Ref<object>,
      mimeTypes: heap.Ref<object>
    ): void => {
      const descriptor_ffi = {};

      descriptor_ffi["appId"] = A.load.Ref(descriptor + 0, undefined);
      descriptor_ffi["taskType"] = A.load.Ref(descriptor + 4, undefined);
      descriptor_ffi["actionId"] = A.load.Ref(descriptor + 8, undefined);

      const _ret = WEBEXT.fileManagerPrivateInternal.setDefaultTask(
        descriptor_ffi,
        A.H.get<object>(urls),
        A.H.get<object>(mimeTypes)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_SetDefaultTask": (
      retPtr: Pointer,
      errPtr: Pointer,
      descriptor: Pointer,
      urls: heap.Ref<object>,
      mimeTypes: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const descriptor_ffi = {};

        descriptor_ffi["appId"] = A.load.Ref(descriptor + 0, undefined);
        descriptor_ffi["taskType"] = A.load.Ref(descriptor + 4, undefined);
        descriptor_ffi["actionId"] = A.load.Ref(descriptor + 8, undefined);

        const _ret = WEBEXT.fileManagerPrivateInternal.setDefaultTask(
          descriptor_ffi,
          A.H.get<object>(urls),
          A.H.get<object>(mimeTypes)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SharePathsWithCrostini": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "sharePathsWithCrostini" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SharePathsWithCrostini": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.sharePathsWithCrostini);
    },
    "call_SharePathsWithCrostini": (
      retPtr: Pointer,
      vmName: heap.Ref<object>,
      urls: heap.Ref<object>,
      persist: heap.Ref<boolean>
    ): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.sharePathsWithCrostini(
        A.H.get<object>(vmName),
        A.H.get<object>(urls),
        persist === A.H.TRUE
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_SharePathsWithCrostini": (
      retPtr: Pointer,
      errPtr: Pointer,
      vmName: heap.Ref<object>,
      urls: heap.Ref<object>,
      persist: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.sharePathsWithCrostini(
          A.H.get<object>(vmName),
          A.H.get<object>(urls),
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
      if (WEBEXT?.fileManagerPrivateInternal && "sharesheetHasTargets" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SharesheetHasTargets": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.sharesheetHasTargets);
    },
    "call_SharesheetHasTargets": (retPtr: Pointer, urls: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.sharesheetHasTargets(A.H.get<object>(urls));
      A.store.Ref(retPtr, _ret);
    },
    "try_SharesheetHasTargets": (retPtr: Pointer, errPtr: Pointer, urls: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.sharesheetHasTargets(A.H.get<object>(urls));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartIOTask": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "startIOTask" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartIOTask": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.startIOTask);
    },
    "call_StartIOTask": (retPtr: Pointer, type: number, urls: heap.Ref<object>, params: Pointer): void => {
      const params_ffi = {};

      params_ffi["destinationFolderUrl"] = A.load.Ref(params + 0, undefined);
      params_ffi["password"] = A.load.Ref(params + 4, undefined);
      if (A.load.Bool(params + 9)) {
        params_ffi["showNotification"] = A.load.Bool(params + 8);
      }

      const _ret = WEBEXT.fileManagerPrivateInternal.startIOTask(
        type > 0 && type <= 9
          ? ["copy", "delete", "empty_trash", "extract", "move", "restore", "restore_to_destination", "trash", "zip"][
              type - 1
            ]
          : undefined,
        A.H.get<object>(urls),
        params_ffi
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_StartIOTask": (
      retPtr: Pointer,
      errPtr: Pointer,
      type: number,
      urls: heap.Ref<object>,
      params: Pointer
    ): heap.Ref<boolean> => {
      try {
        const params_ffi = {};

        params_ffi["destinationFolderUrl"] = A.load.Ref(params + 0, undefined);
        params_ffi["password"] = A.load.Ref(params + 4, undefined);
        if (A.load.Bool(params + 9)) {
          params_ffi["showNotification"] = A.load.Bool(params + 8);
        }

        const _ret = WEBEXT.fileManagerPrivateInternal.startIOTask(
          type > 0 && type <= 9
            ? ["copy", "delete", "empty_trash", "extract", "move", "restore", "restore_to_destination", "trash", "zip"][
                type - 1
              ]
            : undefined,
          A.H.get<object>(urls),
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
      if (WEBEXT?.fileManagerPrivateInternal && "toggleAddedToHoldingSpace" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ToggleAddedToHoldingSpace": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.toggleAddedToHoldingSpace);
    },
    "call_ToggleAddedToHoldingSpace": (retPtr: Pointer, urls: heap.Ref<object>, add: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.toggleAddedToHoldingSpace(A.H.get<object>(urls), add === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_ToggleAddedToHoldingSpace": (
      retPtr: Pointer,
      errPtr: Pointer,
      urls: heap.Ref<object>,
      add: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.toggleAddedToHoldingSpace(
          A.H.get<object>(urls),
          add === A.H.TRUE
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UnsharePathWithCrostini": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "unsharePathWithCrostini" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UnsharePathWithCrostini": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.unsharePathWithCrostini);
    },
    "call_UnsharePathWithCrostini": (retPtr: Pointer, vmName: heap.Ref<object>, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.unsharePathWithCrostini(
        A.H.get<object>(vmName),
        A.H.get<object>(url)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_UnsharePathWithCrostini": (
      retPtr: Pointer,
      errPtr: Pointer,
      vmName: heap.Ref<object>,
      url: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.unsharePathWithCrostini(
          A.H.get<object>(vmName),
          A.H.get<object>(url)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ValidatePathNameLength": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileManagerPrivateInternal && "validatePathNameLength" in WEBEXT?.fileManagerPrivateInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ValidatePathNameLength": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileManagerPrivateInternal.validatePathNameLength);
    },
    "call_ValidatePathNameLength": (retPtr: Pointer, parentUrl: heap.Ref<object>, name: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileManagerPrivateInternal.validatePathNameLength(
        A.H.get<object>(parentUrl),
        A.H.get<object>(name)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_ValidatePathNameLength": (
      retPtr: Pointer,
      errPtr: Pointer,
      parentUrl: heap.Ref<object>,
      name: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileManagerPrivateInternal.validatePathNameLength(
          A.H.get<object>(parentUrl),
          A.H.get<object>(name)
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
