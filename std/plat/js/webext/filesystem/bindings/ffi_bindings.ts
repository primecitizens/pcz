import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/filesystem", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AcceptOption": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["description"]);
        A.store.Ref(ptr + 4, x["mimeTypes"]);
        A.store.Ref(ptr + 8, x["extensions"]);
      }
    },
    "load_AcceptOption": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["description"] = A.load.Ref(ptr + 0, undefined);
      x["mimeTypes"] = A.load.Ref(ptr + 4, undefined);
      x["extensions"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ChooseEntryType": (ref: heap.Ref<string>): number => {
      const idx = ["openFile", "openWritableFile", "saveFile", "openDirectory"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ChooseEntryOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 13, false);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Enum(
          ptr + 0,
          ["openFile", "openWritableFile", "saveFile", "openDirectory"].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 4, x["suggestedName"]);
        A.store.Ref(ptr + 8, x["accepts"]);
        A.store.Bool(ptr + 14, "acceptsAllTypes" in x ? true : false);
        A.store.Bool(ptr + 12, x["acceptsAllTypes"] ? true : false);
        A.store.Bool(ptr + 15, "acceptsMultiple" in x ? true : false);
        A.store.Bool(ptr + 13, x["acceptsMultiple"] ? true : false);
      }
    },
    "load_ChooseEntryOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, ["openFile", "openWritableFile", "saveFile", "openDirectory"]);
      x["suggestedName"] = A.load.Ref(ptr + 4, undefined);
      x["accepts"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 14)) {
        x["acceptsAllTypes"] = A.load.Bool(ptr + 12);
      } else {
        delete x["acceptsAllTypes"];
      }
      if (A.load.Bool(ptr + 15)) {
        x["acceptsMultiple"] = A.load.Bool(ptr + 13);
      } else {
        delete x["acceptsMultiple"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Volume": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_Volume": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["volumeId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 5)) {
        x["writable"] = A.load.Bool(ptr + 4);
      } else {
        delete x["writable"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RequestFileSystemOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_RequestFileSystemOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["volumeId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 5)) {
        x["writable"] = A.load.Bool(ptr + 4);
      } else {
        delete x["writable"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_VolumeListChangedEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["volumes"]);
      }
    },
    "load_VolumeListChangedEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["volumes"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_ChooseEntry": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystem && "chooseEntry" in WEBEXT?.fileSystem) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ChooseEntry": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystem.chooseEntry);
    },
    "call_ChooseEntry": (retPtr: Pointer, options: Pointer, callback: heap.Ref<object>): void => {
      const options_ffi = {};

      options_ffi["type"] = A.load.Enum(options + 0, ["openFile", "openWritableFile", "saveFile", "openDirectory"]);
      options_ffi["suggestedName"] = A.load.Ref(options + 4, undefined);
      options_ffi["accepts"] = A.load.Ref(options + 8, undefined);
      if (A.load.Bool(options + 14)) {
        options_ffi["acceptsAllTypes"] = A.load.Bool(options + 12);
      }
      if (A.load.Bool(options + 15)) {
        options_ffi["acceptsMultiple"] = A.load.Bool(options + 13);
      }

      const _ret = WEBEXT.fileSystem.chooseEntry(options_ffi, A.H.get<object>(callback));
    },
    "try_ChooseEntry": (
      retPtr: Pointer,
      errPtr: Pointer,
      options: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["type"] = A.load.Enum(options + 0, ["openFile", "openWritableFile", "saveFile", "openDirectory"]);
        options_ffi["suggestedName"] = A.load.Ref(options + 4, undefined);
        options_ffi["accepts"] = A.load.Ref(options + 8, undefined);
        if (A.load.Bool(options + 14)) {
          options_ffi["acceptsAllTypes"] = A.load.Bool(options + 12);
        }
        if (A.load.Bool(options + 15)) {
          options_ffi["acceptsMultiple"] = A.load.Bool(options + 13);
        }

        const _ret = WEBEXT.fileSystem.chooseEntry(options_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDisplayPath": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystem && "getDisplayPath" in WEBEXT?.fileSystem) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDisplayPath": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystem.getDisplayPath);
    },
    "call_GetDisplayPath": (retPtr: Pointer, entry: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystem.getDisplayPath(A.H.get<object>(entry));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDisplayPath": (retPtr: Pointer, errPtr: Pointer, entry: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystem.getDisplayPath(A.H.get<object>(entry));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetVolumeList": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystem && "getVolumeList" in WEBEXT?.fileSystem) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetVolumeList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystem.getVolumeList);
    },
    "call_GetVolumeList": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileSystem.getVolumeList();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetVolumeList": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystem.getVolumeList();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetWritableEntry": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystem && "getWritableEntry" in WEBEXT?.fileSystem) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetWritableEntry": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystem.getWritableEntry);
    },
    "call_GetWritableEntry": (retPtr: Pointer, entry: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystem.getWritableEntry(A.H.get<object>(entry));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetWritableEntry": (retPtr: Pointer, errPtr: Pointer, entry: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystem.getWritableEntry(A.H.get<object>(entry));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsRestorable": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystem && "isRestorable" in WEBEXT?.fileSystem) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsRestorable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystem.isRestorable);
    },
    "call_IsRestorable": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystem.isRestorable(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_IsRestorable": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystem.isRestorable(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsWritableEntry": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystem && "isWritableEntry" in WEBEXT?.fileSystem) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsWritableEntry": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystem.isWritableEntry);
    },
    "call_IsWritableEntry": (retPtr: Pointer, entry: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystem.isWritableEntry(A.H.get<object>(entry));
      A.store.Ref(retPtr, _ret);
    },
    "try_IsWritableEntry": (retPtr: Pointer, errPtr: Pointer, entry: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystem.isWritableEntry(A.H.get<object>(entry));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnVolumeListChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystem?.onVolumeListChanged && "addListener" in WEBEXT?.fileSystem?.onVolumeListChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnVolumeListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystem.onVolumeListChanged.addListener);
    },
    "call_OnVolumeListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystem.onVolumeListChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnVolumeListChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystem.onVolumeListChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffVolumeListChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystem?.onVolumeListChanged && "removeListener" in WEBEXT?.fileSystem?.onVolumeListChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffVolumeListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystem.onVolumeListChanged.removeListener);
    },
    "call_OffVolumeListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystem.onVolumeListChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffVolumeListChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystem.onVolumeListChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnVolumeListChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystem?.onVolumeListChanged && "hasListener" in WEBEXT?.fileSystem?.onVolumeListChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnVolumeListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystem.onVolumeListChanged.hasListener);
    },
    "call_HasOnVolumeListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystem.onVolumeListChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnVolumeListChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystem.onVolumeListChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RequestFileSystem": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystem && "requestFileSystem" in WEBEXT?.fileSystem) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RequestFileSystem": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystem.requestFileSystem);
    },
    "call_RequestFileSystem": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["volumeId"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 5)) {
        options_ffi["writable"] = A.load.Bool(options + 4);
      }

      const _ret = WEBEXT.fileSystem.requestFileSystem(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RequestFileSystem": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["volumeId"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 5)) {
          options_ffi["writable"] = A.load.Bool(options + 4);
        }

        const _ret = WEBEXT.fileSystem.requestFileSystem(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RestoreEntry": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystem && "restoreEntry" in WEBEXT?.fileSystem) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RestoreEntry": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystem.restoreEntry);
    },
    "call_RestoreEntry": (retPtr: Pointer, id: heap.Ref<object>, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystem.restoreEntry(A.H.get<object>(id), A.H.get<object>(callback));
    },
    "try_RestoreEntry": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystem.restoreEntry(A.H.get<object>(id), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RetainEntry": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystem && "retainEntry" in WEBEXT?.fileSystem) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RetainEntry": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystem.retainEntry);
    },
    "call_RetainEntry": (retPtr: Pointer, entry: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystem.retainEntry(A.H.get<object>(entry));
      A.store.Ref(retPtr, _ret);
    },
    "try_RetainEntry": (retPtr: Pointer, errPtr: Pointer, entry: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystem.retainEntry(A.H.get<object>(entry));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
