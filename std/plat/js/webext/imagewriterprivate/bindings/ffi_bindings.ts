import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/imagewriterprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_RemovableStorageDevice": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 27, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 25, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 24, false);
      } else {
        A.store.Bool(ptr + 27, true);
        A.store.Ref(ptr + 0, x["storageUnitId"]);
        A.store.Bool(ptr + 25, "capacity" in x ? true : false);
        A.store.Float64(ptr + 8, x["capacity"] === undefined ? 0 : (x["capacity"] as number));
        A.store.Ref(ptr + 16, x["vendor"]);
        A.store.Ref(ptr + 20, x["model"]);
        A.store.Bool(ptr + 26, "removable" in x ? true : false);
        A.store.Bool(ptr + 24, x["removable"] ? true : false);
      }
    },
    "load_RemovableStorageDevice": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["storageUnitId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 25)) {
        x["capacity"] = A.load.Float64(ptr + 8);
      } else {
        delete x["capacity"];
      }
      x["vendor"] = A.load.Ref(ptr + 16, undefined);
      x["model"] = A.load.Ref(ptr + 20, undefined);
      if (A.load.Bool(ptr + 26)) {
        x["removable"] = A.load.Bool(ptr + 24);
      } else {
        delete x["removable"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_Stage": (ref: heap.Ref<string>): number => {
      const idx = ["confirmation", "download", "verifyDownload", "unzip", "write", "verifyWrite", "unknown"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ProgressInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Enum(
          ptr + 0,
          ["confirmation", "download", "verifyDownload", "unzip", "write", "verifyWrite", "unknown"].indexOf(
            x["stage"] as string
          )
        );
        A.store.Bool(ptr + 8, "percentComplete" in x ? true : false);
        A.store.Int32(ptr + 4, x["percentComplete"] === undefined ? 0 : (x["percentComplete"] as number));
      }
    },
    "load_ProgressInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["stage"] = A.load.Enum(ptr + 0, [
        "confirmation",
        "download",
        "verifyDownload",
        "unzip",
        "write",
        "verifyWrite",
        "unknown",
      ]);
      if (A.load.Bool(ptr + 8)) {
        x["percentComplete"] = A.load.Int32(ptr + 4);
      } else {
        delete x["percentComplete"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UrlWriteOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 6, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 6, true);
        A.store.Ref(ptr + 0, x["imageHash"]);
        A.store.Bool(ptr + 5, "saveAsDownload" in x ? true : false);
        A.store.Bool(ptr + 4, x["saveAsDownload"] ? true : false);
      }
    },
    "load_UrlWriteOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["imageHash"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 5)) {
        x["saveAsDownload"] = A.load.Bool(ptr + 4);
      } else {
        delete x["saveAsDownload"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_CancelWrite": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageWriterPrivate && "cancelWrite" in WEBEXT?.imageWriterPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CancelWrite": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.cancelWrite);
    },
    "call_CancelWrite": (retPtr: Pointer): void => {
      const _ret = WEBEXT.imageWriterPrivate.cancelWrite();
      A.store.Ref(retPtr, _ret);
    },
    "try_CancelWrite": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.cancelWrite();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DestroyPartitions": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageWriterPrivate && "destroyPartitions" in WEBEXT?.imageWriterPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DestroyPartitions": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.destroyPartitions);
    },
    "call_DestroyPartitions": (retPtr: Pointer, storageUnitId: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.destroyPartitions(A.H.get<object>(storageUnitId));
      A.store.Ref(retPtr, _ret);
    },
    "try_DestroyPartitions": (retPtr: Pointer, errPtr: Pointer, storageUnitId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.destroyPartitions(A.H.get<object>(storageUnitId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ListRemovableStorageDevices": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageWriterPrivate && "listRemovableStorageDevices" in WEBEXT?.imageWriterPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ListRemovableStorageDevices": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.listRemovableStorageDevices);
    },
    "call_ListRemovableStorageDevices": (retPtr: Pointer): void => {
      const _ret = WEBEXT.imageWriterPrivate.listRemovableStorageDevices();
      A.store.Ref(retPtr, _ret);
    },
    "try_ListRemovableStorageDevices": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.listRemovableStorageDevices();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeviceInserted": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.imageWriterPrivate?.onDeviceInserted &&
        "addListener" in WEBEXT?.imageWriterPrivate?.onDeviceInserted
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeviceInserted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.onDeviceInserted.addListener);
    },
    "call_OnDeviceInserted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.onDeviceInserted.addListener(A.H.get<object>(callback));
    },
    "try_OnDeviceInserted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.onDeviceInserted.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeviceInserted": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.imageWriterPrivate?.onDeviceInserted &&
        "removeListener" in WEBEXT?.imageWriterPrivate?.onDeviceInserted
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeviceInserted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.onDeviceInserted.removeListener);
    },
    "call_OffDeviceInserted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.onDeviceInserted.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeviceInserted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.onDeviceInserted.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeviceInserted": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.imageWriterPrivate?.onDeviceInserted &&
        "hasListener" in WEBEXT?.imageWriterPrivate?.onDeviceInserted
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeviceInserted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.onDeviceInserted.hasListener);
    },
    "call_HasOnDeviceInserted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.onDeviceInserted.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeviceInserted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.onDeviceInserted.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeviceRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageWriterPrivate?.onDeviceRemoved && "addListener" in WEBEXT?.imageWriterPrivate?.onDeviceRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeviceRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.onDeviceRemoved.addListener);
    },
    "call_OnDeviceRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.onDeviceRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnDeviceRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.onDeviceRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeviceRemoved": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.imageWriterPrivate?.onDeviceRemoved &&
        "removeListener" in WEBEXT?.imageWriterPrivate?.onDeviceRemoved
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeviceRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.onDeviceRemoved.removeListener);
    },
    "call_OffDeviceRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.onDeviceRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeviceRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.onDeviceRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeviceRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageWriterPrivate?.onDeviceRemoved && "hasListener" in WEBEXT?.imageWriterPrivate?.onDeviceRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeviceRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.onDeviceRemoved.hasListener);
    },
    "call_HasOnDeviceRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.onDeviceRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeviceRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.onDeviceRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnWriteComplete": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageWriterPrivate?.onWriteComplete && "addListener" in WEBEXT?.imageWriterPrivate?.onWriteComplete) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnWriteComplete": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.onWriteComplete.addListener);
    },
    "call_OnWriteComplete": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.onWriteComplete.addListener(A.H.get<object>(callback));
    },
    "try_OnWriteComplete": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.onWriteComplete.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffWriteComplete": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.imageWriterPrivate?.onWriteComplete &&
        "removeListener" in WEBEXT?.imageWriterPrivate?.onWriteComplete
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffWriteComplete": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.onWriteComplete.removeListener);
    },
    "call_OffWriteComplete": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.onWriteComplete.removeListener(A.H.get<object>(callback));
    },
    "try_OffWriteComplete": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.onWriteComplete.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnWriteComplete": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageWriterPrivate?.onWriteComplete && "hasListener" in WEBEXT?.imageWriterPrivate?.onWriteComplete) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnWriteComplete": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.onWriteComplete.hasListener);
    },
    "call_HasOnWriteComplete": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.onWriteComplete.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnWriteComplete": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.onWriteComplete.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnWriteError": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageWriterPrivate?.onWriteError && "addListener" in WEBEXT?.imageWriterPrivate?.onWriteError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnWriteError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.onWriteError.addListener);
    },
    "call_OnWriteError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.onWriteError.addListener(A.H.get<object>(callback));
    },
    "try_OnWriteError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.onWriteError.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffWriteError": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageWriterPrivate?.onWriteError && "removeListener" in WEBEXT?.imageWriterPrivate?.onWriteError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffWriteError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.onWriteError.removeListener);
    },
    "call_OffWriteError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.onWriteError.removeListener(A.H.get<object>(callback));
    },
    "try_OffWriteError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.onWriteError.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnWriteError": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageWriterPrivate?.onWriteError && "hasListener" in WEBEXT?.imageWriterPrivate?.onWriteError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnWriteError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.onWriteError.hasListener);
    },
    "call_HasOnWriteError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.onWriteError.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnWriteError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.onWriteError.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnWriteProgress": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageWriterPrivate?.onWriteProgress && "addListener" in WEBEXT?.imageWriterPrivate?.onWriteProgress) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnWriteProgress": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.onWriteProgress.addListener);
    },
    "call_OnWriteProgress": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.onWriteProgress.addListener(A.H.get<object>(callback));
    },
    "try_OnWriteProgress": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.onWriteProgress.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffWriteProgress": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.imageWriterPrivate?.onWriteProgress &&
        "removeListener" in WEBEXT?.imageWriterPrivate?.onWriteProgress
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffWriteProgress": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.onWriteProgress.removeListener);
    },
    "call_OffWriteProgress": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.onWriteProgress.removeListener(A.H.get<object>(callback));
    },
    "try_OffWriteProgress": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.onWriteProgress.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnWriteProgress": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageWriterPrivate?.onWriteProgress && "hasListener" in WEBEXT?.imageWriterPrivate?.onWriteProgress) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnWriteProgress": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.onWriteProgress.hasListener);
    },
    "call_HasOnWriteProgress": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.onWriteProgress.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnWriteProgress": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.onWriteProgress.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_WriteFromFile": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageWriterPrivate && "writeFromFile" in WEBEXT?.imageWriterPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_WriteFromFile": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.writeFromFile);
    },
    "call_WriteFromFile": (retPtr: Pointer, storageUnitId: heap.Ref<object>, fileEntry: heap.Ref<object>): void => {
      const _ret = WEBEXT.imageWriterPrivate.writeFromFile(A.H.get<object>(storageUnitId), A.H.get<object>(fileEntry));
      A.store.Ref(retPtr, _ret);
    },
    "try_WriteFromFile": (
      retPtr: Pointer,
      errPtr: Pointer,
      storageUnitId: heap.Ref<object>,
      fileEntry: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageWriterPrivate.writeFromFile(
          A.H.get<object>(storageUnitId),
          A.H.get<object>(fileEntry)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_WriteFromUrl": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageWriterPrivate && "writeFromUrl" in WEBEXT?.imageWriterPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_WriteFromUrl": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageWriterPrivate.writeFromUrl);
    },
    "call_WriteFromUrl": (
      retPtr: Pointer,
      storageUnitId: heap.Ref<object>,
      imageUrl: heap.Ref<object>,
      options: Pointer
    ): void => {
      const options_ffi = {};

      options_ffi["imageHash"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 5)) {
        options_ffi["saveAsDownload"] = A.load.Bool(options + 4);
      }

      const _ret = WEBEXT.imageWriterPrivate.writeFromUrl(
        A.H.get<object>(storageUnitId),
        A.H.get<object>(imageUrl),
        options_ffi
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_WriteFromUrl": (
      retPtr: Pointer,
      errPtr: Pointer,
      storageUnitId: heap.Ref<object>,
      imageUrl: heap.Ref<object>,
      options: Pointer
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["imageHash"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 5)) {
          options_ffi["saveAsDownload"] = A.load.Bool(options + 4);
        }

        const _ret = WEBEXT.imageWriterPrivate.writeFromUrl(
          A.H.get<object>(storageUnitId),
          A.H.get<object>(imageUrl),
          options_ffi
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
