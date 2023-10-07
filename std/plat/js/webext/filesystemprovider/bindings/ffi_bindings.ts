import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/filesystemprovider", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AbortRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 12, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Bool(ptr + 13, "operationRequestId" in x ? true : false);
        A.store.Int32(ptr + 8, x["operationRequestId"] === undefined ? 0 : (x["operationRequestId"] as number));
      }
    },
    "load_AbortRequestedOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["operationRequestId"] = A.load.Int32(ptr + 8);
      } else {
        delete x["operationRequestId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Action": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_Action": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["title"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AddWatcherRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 15, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 12, false);
      } else {
        A.store.Bool(ptr + 15, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 13, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Ref(ptr + 8, x["entryPath"]);
        A.store.Bool(ptr + 14, "recursive" in x ? true : false);
        A.store.Bool(ptr + 12, x["recursive"] ? true : false);
      }
    },
    "load_AddWatcherRequestedOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 13)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      x["entryPath"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 14)) {
        x["recursive"] = A.load.Bool(ptr + 12);
      } else {
        delete x["recursive"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ChangeType": (ref: heap.Ref<string>): number => {
      const idx = ["CHANGED", "DELETED"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Change": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["entryPath"]);
        A.store.Enum(ptr + 4, ["CHANGED", "DELETED"].indexOf(x["changeType"] as string));
      }
    },
    "load_Change": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["entryPath"] = A.load.Ref(ptr + 0, undefined);
      x["changeType"] = A.load.Enum(ptr + 4, ["CHANGED", "DELETED"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CloseFileRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 12, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Bool(ptr + 13, "openRequestId" in x ? true : false);
        A.store.Int32(ptr + 8, x["openRequestId"] === undefined ? 0 : (x["openRequestId"] as number));
      }
    },
    "load_CloseFileRequestedOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["openRequestId"] = A.load.Int32(ptr + 8);
      } else {
        delete x["openRequestId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CloudIdentifier": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["providerName"]);
        A.store.Ref(ptr + 4, x["id"]);
      }
    },
    "load_CloudIdentifier": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["providerName"] = A.load.Ref(ptr + 0, undefined);
      x["id"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_CommonActionId": (ref: heap.Ref<string>): number => {
      const idx = ["SAVE_FOR_OFFLINE", "OFFLINE_NOT_NECESSARY", "SHARE"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ConfigureRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 8, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
      }
    },
    "load_ConfigureRequestedOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 8)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CopyEntryRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 16, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Ref(ptr + 8, x["sourcePath"]);
        A.store.Ref(ptr + 12, x["targetPath"]);
      }
    },
    "load_CopyEntryRequestedOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      x["sourcePath"] = A.load.Ref(ptr + 8, undefined);
      x["targetPath"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CreateDirectoryRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 15, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 12, false);
      } else {
        A.store.Bool(ptr + 15, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 13, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Ref(ptr + 8, x["directoryPath"]);
        A.store.Bool(ptr + 14, "recursive" in x ? true : false);
        A.store.Bool(ptr + 12, x["recursive"] ? true : false);
      }
    },
    "load_CreateDirectoryRequestedOptions": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 13)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      x["directoryPath"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 14)) {
        x["recursive"] = A.load.Bool(ptr + 12);
      } else {
        delete x["recursive"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CreateFileRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 12, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Ref(ptr + 8, x["filePath"]);
      }
    },
    "load_CreateFileRequestedOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      x["filePath"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DeleteEntryRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 15, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 12, false);
      } else {
        A.store.Bool(ptr + 15, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 13, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Ref(ptr + 8, x["entryPath"]);
        A.store.Bool(ptr + 14, "recursive" in x ? true : false);
        A.store.Bool(ptr + 12, x["recursive"] ? true : false);
      }
    },
    "load_DeleteEntryRequestedOptions": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 13)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      x["entryPath"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 14)) {
        x["recursive"] = A.load.Bool(ptr + 12);
      } else {
        delete x["recursive"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_EntryMetadata": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 39, false);
        A.store.Bool(ptr + 37, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 38, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);

        A.store.Bool(ptr + 28 + 8, false);
        A.store.Ref(ptr + 28 + 0, undefined);
        A.store.Ref(ptr + 28 + 4, undefined);
      } else {
        A.store.Bool(ptr + 39, true);
        A.store.Bool(ptr + 37, "isDirectory" in x ? true : false);
        A.store.Bool(ptr + 0, x["isDirectory"] ? true : false);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Bool(ptr + 38, "size" in x ? true : false);
        A.store.Float64(ptr + 8, x["size"] === undefined ? 0 : (x["size"] as number));
        A.store.Ref(ptr + 16, x["modificationTime"]);
        A.store.Ref(ptr + 20, x["mimeType"]);
        A.store.Ref(ptr + 24, x["thumbnail"]);

        if (typeof x["cloudIdentifier"] === "undefined") {
          A.store.Bool(ptr + 28 + 8, false);
          A.store.Ref(ptr + 28 + 0, undefined);
          A.store.Ref(ptr + 28 + 4, undefined);
        } else {
          A.store.Bool(ptr + 28 + 8, true);
          A.store.Ref(ptr + 28 + 0, x["cloudIdentifier"]["providerName"]);
          A.store.Ref(ptr + 28 + 4, x["cloudIdentifier"]["id"]);
        }
      }
    },
    "load_EntryMetadata": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 37)) {
        x["isDirectory"] = A.load.Bool(ptr + 0);
      } else {
        delete x["isDirectory"];
      }
      x["name"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 38)) {
        x["size"] = A.load.Float64(ptr + 8);
      } else {
        delete x["size"];
      }
      x["modificationTime"] = A.load.Ref(ptr + 16, undefined);
      x["mimeType"] = A.load.Ref(ptr + 20, undefined);
      x["thumbnail"] = A.load.Ref(ptr + 24, undefined);
      if (A.load.Bool(ptr + 28 + 8)) {
        x["cloudIdentifier"] = {};
        x["cloudIdentifier"]["providerName"] = A.load.Ref(ptr + 28 + 0, undefined);
        x["cloudIdentifier"]["id"] = A.load.Ref(ptr + 28 + 4, undefined);
      } else {
        delete x["cloudIdentifier"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ExecuteActionRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 16, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Ref(ptr + 8, x["entryPaths"]);
        A.store.Ref(ptr + 12, x["actionId"]);
      }
    },
    "load_ExecuteActionRequestedOptions": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      x["entryPaths"] = A.load.Ref(ptr + 8, undefined);
      x["actionId"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OpenFileMode": (ref: heap.Ref<string>): number => {
      const idx = ["READ", "WRITE"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OpenedFile": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "openRequestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["openRequestId"] === undefined ? 0 : (x["openRequestId"] as number));
        A.store.Ref(ptr + 4, x["filePath"]);
        A.store.Enum(ptr + 8, ["READ", "WRITE"].indexOf(x["mode"] as string));
      }
    },
    "load_OpenedFile": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["openRequestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["openRequestId"];
      }
      x["filePath"] = A.load.Ref(ptr + 4, undefined);
      x["mode"] = A.load.Enum(ptr + 8, ["READ", "WRITE"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Watcher": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Ref(ptr + 0, x["entryPath"]);
        A.store.Bool(ptr + 12, "recursive" in x ? true : false);
        A.store.Bool(ptr + 4, x["recursive"] ? true : false);
        A.store.Ref(ptr + 8, x["lastTag"]);
      }
    },
    "load_Watcher": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["entryPath"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["recursive"] = A.load.Bool(ptr + 4);
      } else {
        delete x["recursive"];
      }
      x["lastTag"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_FileSystemInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 31, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 29, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 30, false);
        A.store.Bool(ptr + 20, false);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 31, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Ref(ptr + 4, x["displayName"]);
        A.store.Bool(ptr + 28, "writable" in x ? true : false);
        A.store.Bool(ptr + 8, x["writable"] ? true : false);
        A.store.Bool(ptr + 29, "openedFilesLimit" in x ? true : false);
        A.store.Int32(ptr + 12, x["openedFilesLimit"] === undefined ? 0 : (x["openedFilesLimit"] as number));
        A.store.Ref(ptr + 16, x["openedFiles"]);
        A.store.Bool(ptr + 30, "supportsNotifyTag" in x ? true : false);
        A.store.Bool(ptr + 20, x["supportsNotifyTag"] ? true : false);
        A.store.Ref(ptr + 24, x["watchers"]);
      }
    },
    "load_FileSystemInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      x["displayName"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 28)) {
        x["writable"] = A.load.Bool(ptr + 8);
      } else {
        delete x["writable"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["openedFilesLimit"] = A.load.Int32(ptr + 12);
      } else {
        delete x["openedFilesLimit"];
      }
      x["openedFiles"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 30)) {
        x["supportsNotifyTag"] = A.load.Bool(ptr + 20);
      } else {
        delete x["supportsNotifyTag"];
      }
      x["watchers"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetActionsRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 12, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Ref(ptr + 8, x["entryPaths"]);
      }
    },
    "load_GetActionsRequestedOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      x["entryPaths"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetMetadataRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 27, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 19, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 22, false);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 23, false);
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 18, false);
      } else {
        A.store.Bool(ptr + 27, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 19, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Ref(ptr + 8, x["entryPath"]);
        A.store.Bool(ptr + 20, "isDirectory" in x ? true : false);
        A.store.Bool(ptr + 12, x["isDirectory"] ? true : false);
        A.store.Bool(ptr + 21, "name" in x ? true : false);
        A.store.Bool(ptr + 13, x["name"] ? true : false);
        A.store.Bool(ptr + 22, "size" in x ? true : false);
        A.store.Bool(ptr + 14, x["size"] ? true : false);
        A.store.Bool(ptr + 23, "modificationTime" in x ? true : false);
        A.store.Bool(ptr + 15, x["modificationTime"] ? true : false);
        A.store.Bool(ptr + 24, "mimeType" in x ? true : false);
        A.store.Bool(ptr + 16, x["mimeType"] ? true : false);
        A.store.Bool(ptr + 25, "thumbnail" in x ? true : false);
        A.store.Bool(ptr + 17, x["thumbnail"] ? true : false);
        A.store.Bool(ptr + 26, "cloudIdentifier" in x ? true : false);
        A.store.Bool(ptr + 18, x["cloudIdentifier"] ? true : false);
      }
    },
    "load_GetMetadataRequestedOptions": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 19)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      x["entryPath"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 20)) {
        x["isDirectory"] = A.load.Bool(ptr + 12);
      } else {
        delete x["isDirectory"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["name"] = A.load.Bool(ptr + 13);
      } else {
        delete x["name"];
      }
      if (A.load.Bool(ptr + 22)) {
        x["size"] = A.load.Bool(ptr + 14);
      } else {
        delete x["size"];
      }
      if (A.load.Bool(ptr + 23)) {
        x["modificationTime"] = A.load.Bool(ptr + 15);
      } else {
        delete x["modificationTime"];
      }
      if (A.load.Bool(ptr + 24)) {
        x["mimeType"] = A.load.Bool(ptr + 16);
      } else {
        delete x["mimeType"];
      }
      if (A.load.Bool(ptr + 25)) {
        x["thumbnail"] = A.load.Bool(ptr + 17);
      } else {
        delete x["thumbnail"];
      }
      if (A.load.Bool(ptr + 26)) {
        x["cloudIdentifier"] = A.load.Bool(ptr + 18);
      } else {
        delete x["cloudIdentifier"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MountOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 22, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 19, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 17, false);
      } else {
        A.store.Bool(ptr + 22, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Ref(ptr + 4, x["displayName"]);
        A.store.Bool(ptr + 18, "writable" in x ? true : false);
        A.store.Bool(ptr + 8, x["writable"] ? true : false);
        A.store.Bool(ptr + 19, "openedFilesLimit" in x ? true : false);
        A.store.Int32(ptr + 12, x["openedFilesLimit"] === undefined ? 0 : (x["openedFilesLimit"] as number));
        A.store.Bool(ptr + 20, "supportsNotifyTag" in x ? true : false);
        A.store.Bool(ptr + 16, x["supportsNotifyTag"] ? true : false);
        A.store.Bool(ptr + 21, "persistent" in x ? true : false);
        A.store.Bool(ptr + 17, x["persistent"] ? true : false);
      }
    },
    "load_MountOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      x["displayName"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 18)) {
        x["writable"] = A.load.Bool(ptr + 8);
      } else {
        delete x["writable"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["openedFilesLimit"] = A.load.Int32(ptr + 12);
      } else {
        delete x["openedFilesLimit"];
      }
      if (A.load.Bool(ptr + 20)) {
        x["supportsNotifyTag"] = A.load.Bool(ptr + 16);
      } else {
        delete x["supportsNotifyTag"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["persistent"] = A.load.Bool(ptr + 17);
      } else {
        delete x["persistent"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MoveEntryRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 16, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Ref(ptr + 8, x["sourcePath"]);
        A.store.Ref(ptr + 12, x["targetPath"]);
      }
    },
    "load_MoveEntryRequestedOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      x["sourcePath"] = A.load.Ref(ptr + 8, undefined);
      x["targetPath"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_NotifyOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 12, -1);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
      } else {
        A.store.Bool(ptr + 25, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Ref(ptr + 4, x["observedPath"]);
        A.store.Bool(ptr + 24, "recursive" in x ? true : false);
        A.store.Bool(ptr + 8, x["recursive"] ? true : false);
        A.store.Enum(ptr + 12, ["CHANGED", "DELETED"].indexOf(x["changeType"] as string));
        A.store.Ref(ptr + 16, x["changes"]);
        A.store.Ref(ptr + 20, x["tag"]);
      }
    },
    "load_NotifyOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      x["observedPath"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 24)) {
        x["recursive"] = A.load.Bool(ptr + 8);
      } else {
        delete x["recursive"];
      }
      x["changeType"] = A.load.Enum(ptr + 12, ["CHANGED", "DELETED"]);
      x["changes"] = A.load.Ref(ptr + 16, undefined);
      x["tag"] = A.load.Ref(ptr + 20, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OpenFileRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Enum(ptr + 12, -1);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 16, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Ref(ptr + 8, x["filePath"]);
        A.store.Enum(ptr + 12, ["READ", "WRITE"].indexOf(x["mode"] as string));
      }
    },
    "load_OpenFileRequestedOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      x["filePath"] = A.load.Ref(ptr + 8, undefined);
      x["mode"] = A.load.Enum(ptr + 12, ["READ", "WRITE"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ProviderError": (ref: heap.Ref<string>): number => {
      const idx = [
        "OK",
        "FAILED",
        "IN_USE",
        "EXISTS",
        "NOT_FOUND",
        "ACCESS_DENIED",
        "TOO_MANY_OPENED",
        "NO_MEMORY",
        "NO_SPACE",
        "NOT_A_DIRECTORY",
        "INVALID_OPERATION",
        "SECURITY",
        "ABORT",
        "NOT_A_FILE",
        "NOT_EMPTY",
        "INVALID_URL",
        "IO",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ReadDirectoryRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 18, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 22, false);
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 23, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 17, false);
      } else {
        A.store.Bool(ptr + 25, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 18, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Ref(ptr + 8, x["directoryPath"]);
        A.store.Bool(ptr + 19, "isDirectory" in x ? true : false);
        A.store.Bool(ptr + 12, x["isDirectory"] ? true : false);
        A.store.Bool(ptr + 20, "name" in x ? true : false);
        A.store.Bool(ptr + 13, x["name"] ? true : false);
        A.store.Bool(ptr + 21, "size" in x ? true : false);
        A.store.Bool(ptr + 14, x["size"] ? true : false);
        A.store.Bool(ptr + 22, "modificationTime" in x ? true : false);
        A.store.Bool(ptr + 15, x["modificationTime"] ? true : false);
        A.store.Bool(ptr + 23, "mimeType" in x ? true : false);
        A.store.Bool(ptr + 16, x["mimeType"] ? true : false);
        A.store.Bool(ptr + 24, "thumbnail" in x ? true : false);
        A.store.Bool(ptr + 17, x["thumbnail"] ? true : false);
      }
    },
    "load_ReadDirectoryRequestedOptions": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 18)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      x["directoryPath"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 19)) {
        x["isDirectory"] = A.load.Bool(ptr + 12);
      } else {
        delete x["isDirectory"];
      }
      if (A.load.Bool(ptr + 20)) {
        x["name"] = A.load.Bool(ptr + 13);
      } else {
        delete x["name"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["size"] = A.load.Bool(ptr + 14);
      } else {
        delete x["size"];
      }
      if (A.load.Bool(ptr + 22)) {
        x["modificationTime"] = A.load.Bool(ptr + 15);
      } else {
        delete x["modificationTime"];
      }
      if (A.load.Bool(ptr + 23)) {
        x["mimeType"] = A.load.Bool(ptr + 16);
      } else {
        delete x["mimeType"];
      }
      if (A.load.Bool(ptr + 24)) {
        x["thumbnail"] = A.load.Bool(ptr + 17);
      } else {
        delete x["thumbnail"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ReadFileRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 36, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 32, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 33, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 34, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 35, false);
        A.store.Float64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 36, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 32, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Bool(ptr + 33, "openRequestId" in x ? true : false);
        A.store.Int32(ptr + 8, x["openRequestId"] === undefined ? 0 : (x["openRequestId"] as number));
        A.store.Bool(ptr + 34, "offset" in x ? true : false);
        A.store.Float64(ptr + 16, x["offset"] === undefined ? 0 : (x["offset"] as number));
        A.store.Bool(ptr + 35, "length" in x ? true : false);
        A.store.Float64(ptr + 24, x["length"] === undefined ? 0 : (x["length"] as number));
      }
    },
    "load_ReadFileRequestedOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 32)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["openRequestId"] = A.load.Int32(ptr + 8);
      } else {
        delete x["openRequestId"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["offset"] = A.load.Float64(ptr + 16);
      } else {
        delete x["offset"];
      }
      if (A.load.Bool(ptr + 35)) {
        x["length"] = A.load.Float64(ptr + 24);
      } else {
        delete x["length"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RemoveWatcherRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 15, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 12, false);
      } else {
        A.store.Bool(ptr + 15, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 13, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Ref(ptr + 8, x["entryPath"]);
        A.store.Bool(ptr + 14, "recursive" in x ? true : false);
        A.store.Bool(ptr + 12, x["recursive"] ? true : false);
      }
    },
    "load_RemoveWatcherRequestedOptions": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 13)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      x["entryPath"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 14)) {
        x["recursive"] = A.load.Bool(ptr + 12);
      } else {
        delete x["recursive"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TruncateRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 26, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 24, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 25, false);
        A.store.Float64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 26, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 24, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Ref(ptr + 8, x["filePath"]);
        A.store.Bool(ptr + 25, "length" in x ? true : false);
        A.store.Float64(ptr + 16, x["length"] === undefined ? 0 : (x["length"] as number));
      }
    },
    "load_TruncateRequestedOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 24)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      x["filePath"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 25)) {
        x["length"] = A.load.Float64(ptr + 16);
      } else {
        delete x["length"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UnmountOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
      }
    },
    "load_UnmountOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UnmountRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 8, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
      }
    },
    "load_UnmountRequestedOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 8)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_WriteFileRequestedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 31, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 28, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 29, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 30, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 31, true);
        A.store.Ref(ptr + 0, x["fileSystemId"]);
        A.store.Bool(ptr + 28, "requestId" in x ? true : false);
        A.store.Int32(ptr + 4, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Bool(ptr + 29, "openRequestId" in x ? true : false);
        A.store.Int32(ptr + 8, x["openRequestId"] === undefined ? 0 : (x["openRequestId"] as number));
        A.store.Bool(ptr + 30, "offset" in x ? true : false);
        A.store.Float64(ptr + 16, x["offset"] === undefined ? 0 : (x["offset"] as number));
        A.store.Ref(ptr + 24, x["data"]);
      }
    },
    "load_WriteFileRequestedOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fileSystemId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 28)) {
        x["requestId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["requestId"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["openRequestId"] = A.load.Int32(ptr + 8);
      } else {
        delete x["openRequestId"];
      }
      if (A.load.Bool(ptr + 30)) {
        x["offset"] = A.load.Float64(ptr + 16);
      } else {
        delete x["offset"];
      }
      x["data"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Get": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystemProvider && "get" in WEBEXT?.fileSystemProvider) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Get": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.get);
    },
    "call_Get": (retPtr: Pointer, fileSystemId: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.get(A.H.get<object>(fileSystemId));
      A.store.Ref(retPtr, _ret);
    },
    "try_Get": (retPtr: Pointer, errPtr: Pointer, fileSystemId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.get(A.H.get<object>(fileSystemId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAll": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystemProvider && "getAll" in WEBEXT?.fileSystemProvider) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.getAll);
    },
    "call_GetAll": (retPtr: Pointer): void => {
      const _ret = WEBEXT.fileSystemProvider.getAll();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAll": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.getAll();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Mount": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystemProvider && "mount" in WEBEXT?.fileSystemProvider) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Mount": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.mount);
    },
    "call_Mount": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["fileSystemId"] = A.load.Ref(options + 0, undefined);
      options_ffi["displayName"] = A.load.Ref(options + 4, undefined);
      if (A.load.Bool(options + 18)) {
        options_ffi["writable"] = A.load.Bool(options + 8);
      }
      if (A.load.Bool(options + 19)) {
        options_ffi["openedFilesLimit"] = A.load.Int32(options + 12);
      }
      if (A.load.Bool(options + 20)) {
        options_ffi["supportsNotifyTag"] = A.load.Bool(options + 16);
      }
      if (A.load.Bool(options + 21)) {
        options_ffi["persistent"] = A.load.Bool(options + 17);
      }

      const _ret = WEBEXT.fileSystemProvider.mount(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Mount": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["fileSystemId"] = A.load.Ref(options + 0, undefined);
        options_ffi["displayName"] = A.load.Ref(options + 4, undefined);
        if (A.load.Bool(options + 18)) {
          options_ffi["writable"] = A.load.Bool(options + 8);
        }
        if (A.load.Bool(options + 19)) {
          options_ffi["openedFilesLimit"] = A.load.Int32(options + 12);
        }
        if (A.load.Bool(options + 20)) {
          options_ffi["supportsNotifyTag"] = A.load.Bool(options + 16);
        }
        if (A.load.Bool(options + 21)) {
          options_ffi["persistent"] = A.load.Bool(options + 17);
        }

        const _ret = WEBEXT.fileSystemProvider.mount(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Notify": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystemProvider && "notify" in WEBEXT?.fileSystemProvider) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Notify": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.notify);
    },
    "call_Notify": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["fileSystemId"] = A.load.Ref(options + 0, undefined);
      options_ffi["observedPath"] = A.load.Ref(options + 4, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["recursive"] = A.load.Bool(options + 8);
      }
      options_ffi["changeType"] = A.load.Enum(options + 12, ["CHANGED", "DELETED"]);
      options_ffi["changes"] = A.load.Ref(options + 16, undefined);
      options_ffi["tag"] = A.load.Ref(options + 20, undefined);

      const _ret = WEBEXT.fileSystemProvider.notify(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Notify": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["fileSystemId"] = A.load.Ref(options + 0, undefined);
        options_ffi["observedPath"] = A.load.Ref(options + 4, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["recursive"] = A.load.Bool(options + 8);
        }
        options_ffi["changeType"] = A.load.Enum(options + 12, ["CHANGED", "DELETED"]);
        options_ffi["changes"] = A.load.Ref(options + 16, undefined);
        options_ffi["tag"] = A.load.Ref(options + 20, undefined);

        const _ret = WEBEXT.fileSystemProvider.notify(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAbortRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onAbortRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onAbortRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAbortRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onAbortRequested.addListener);
    },
    "call_OnAbortRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onAbortRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnAbortRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onAbortRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAbortRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onAbortRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onAbortRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAbortRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onAbortRequested.removeListener);
    },
    "call_OffAbortRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onAbortRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffAbortRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onAbortRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAbortRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onAbortRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onAbortRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAbortRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onAbortRequested.hasListener);
    },
    "call_HasOnAbortRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onAbortRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAbortRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onAbortRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAddWatcherRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onAddWatcherRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onAddWatcherRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAddWatcherRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onAddWatcherRequested.addListener);
    },
    "call_OnAddWatcherRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onAddWatcherRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnAddWatcherRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onAddWatcherRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAddWatcherRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onAddWatcherRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onAddWatcherRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAddWatcherRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onAddWatcherRequested.removeListener);
    },
    "call_OffAddWatcherRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onAddWatcherRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffAddWatcherRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onAddWatcherRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAddWatcherRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onAddWatcherRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onAddWatcherRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAddWatcherRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onAddWatcherRequested.hasListener);
    },
    "call_HasOnAddWatcherRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onAddWatcherRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAddWatcherRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onAddWatcherRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCloseFileRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onCloseFileRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onCloseFileRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCloseFileRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onCloseFileRequested.addListener);
    },
    "call_OnCloseFileRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onCloseFileRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnCloseFileRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onCloseFileRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCloseFileRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onCloseFileRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onCloseFileRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCloseFileRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onCloseFileRequested.removeListener);
    },
    "call_OffCloseFileRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onCloseFileRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffCloseFileRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onCloseFileRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCloseFileRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onCloseFileRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onCloseFileRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCloseFileRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onCloseFileRequested.hasListener);
    },
    "call_HasOnCloseFileRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onCloseFileRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCloseFileRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onCloseFileRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnConfigureRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onConfigureRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onConfigureRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnConfigureRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onConfigureRequested.addListener);
    },
    "call_OnConfigureRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onConfigureRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnConfigureRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onConfigureRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffConfigureRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onConfigureRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onConfigureRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffConfigureRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onConfigureRequested.removeListener);
    },
    "call_OffConfigureRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onConfigureRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffConfigureRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onConfigureRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnConfigureRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onConfigureRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onConfigureRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnConfigureRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onConfigureRequested.hasListener);
    },
    "call_HasOnConfigureRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onConfigureRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnConfigureRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onConfigureRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCopyEntryRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onCopyEntryRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onCopyEntryRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCopyEntryRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onCopyEntryRequested.addListener);
    },
    "call_OnCopyEntryRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onCopyEntryRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnCopyEntryRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onCopyEntryRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCopyEntryRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onCopyEntryRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onCopyEntryRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCopyEntryRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onCopyEntryRequested.removeListener);
    },
    "call_OffCopyEntryRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onCopyEntryRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffCopyEntryRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onCopyEntryRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCopyEntryRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onCopyEntryRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onCopyEntryRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCopyEntryRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onCopyEntryRequested.hasListener);
    },
    "call_HasOnCopyEntryRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onCopyEntryRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCopyEntryRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onCopyEntryRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCreateDirectoryRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onCreateDirectoryRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onCreateDirectoryRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCreateDirectoryRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onCreateDirectoryRequested.addListener);
    },
    "call_OnCreateDirectoryRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onCreateDirectoryRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnCreateDirectoryRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onCreateDirectoryRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCreateDirectoryRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onCreateDirectoryRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onCreateDirectoryRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCreateDirectoryRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onCreateDirectoryRequested.removeListener);
    },
    "call_OffCreateDirectoryRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onCreateDirectoryRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffCreateDirectoryRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onCreateDirectoryRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCreateDirectoryRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onCreateDirectoryRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onCreateDirectoryRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCreateDirectoryRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onCreateDirectoryRequested.hasListener);
    },
    "call_HasOnCreateDirectoryRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onCreateDirectoryRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCreateDirectoryRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onCreateDirectoryRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCreateFileRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onCreateFileRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onCreateFileRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCreateFileRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onCreateFileRequested.addListener);
    },
    "call_OnCreateFileRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onCreateFileRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnCreateFileRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onCreateFileRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCreateFileRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onCreateFileRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onCreateFileRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCreateFileRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onCreateFileRequested.removeListener);
    },
    "call_OffCreateFileRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onCreateFileRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffCreateFileRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onCreateFileRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCreateFileRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onCreateFileRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onCreateFileRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCreateFileRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onCreateFileRequested.hasListener);
    },
    "call_HasOnCreateFileRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onCreateFileRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCreateFileRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onCreateFileRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeleteEntryRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onDeleteEntryRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onDeleteEntryRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeleteEntryRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onDeleteEntryRequested.addListener);
    },
    "call_OnDeleteEntryRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onDeleteEntryRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnDeleteEntryRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onDeleteEntryRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeleteEntryRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onDeleteEntryRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onDeleteEntryRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeleteEntryRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onDeleteEntryRequested.removeListener);
    },
    "call_OffDeleteEntryRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onDeleteEntryRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeleteEntryRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onDeleteEntryRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeleteEntryRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onDeleteEntryRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onDeleteEntryRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeleteEntryRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onDeleteEntryRequested.hasListener);
    },
    "call_HasOnDeleteEntryRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onDeleteEntryRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeleteEntryRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onDeleteEntryRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnExecuteActionRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onExecuteActionRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onExecuteActionRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnExecuteActionRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onExecuteActionRequested.addListener);
    },
    "call_OnExecuteActionRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onExecuteActionRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnExecuteActionRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onExecuteActionRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffExecuteActionRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onExecuteActionRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onExecuteActionRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffExecuteActionRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onExecuteActionRequested.removeListener);
    },
    "call_OffExecuteActionRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onExecuteActionRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffExecuteActionRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onExecuteActionRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnExecuteActionRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onExecuteActionRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onExecuteActionRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnExecuteActionRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onExecuteActionRequested.hasListener);
    },
    "call_HasOnExecuteActionRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onExecuteActionRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnExecuteActionRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onExecuteActionRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnGetActionsRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onGetActionsRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onGetActionsRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnGetActionsRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onGetActionsRequested.addListener);
    },
    "call_OnGetActionsRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onGetActionsRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnGetActionsRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onGetActionsRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffGetActionsRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onGetActionsRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onGetActionsRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffGetActionsRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onGetActionsRequested.removeListener);
    },
    "call_OffGetActionsRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onGetActionsRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffGetActionsRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onGetActionsRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnGetActionsRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onGetActionsRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onGetActionsRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnGetActionsRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onGetActionsRequested.hasListener);
    },
    "call_HasOnGetActionsRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onGetActionsRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnGetActionsRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onGetActionsRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnGetMetadataRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onGetMetadataRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onGetMetadataRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnGetMetadataRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onGetMetadataRequested.addListener);
    },
    "call_OnGetMetadataRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onGetMetadataRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnGetMetadataRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onGetMetadataRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffGetMetadataRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onGetMetadataRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onGetMetadataRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffGetMetadataRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onGetMetadataRequested.removeListener);
    },
    "call_OffGetMetadataRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onGetMetadataRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffGetMetadataRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onGetMetadataRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnGetMetadataRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onGetMetadataRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onGetMetadataRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnGetMetadataRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onGetMetadataRequested.hasListener);
    },
    "call_HasOnGetMetadataRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onGetMetadataRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnGetMetadataRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onGetMetadataRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMountRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onMountRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onMountRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMountRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onMountRequested.addListener);
    },
    "call_OnMountRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onMountRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnMountRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onMountRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMountRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onMountRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onMountRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMountRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onMountRequested.removeListener);
    },
    "call_OffMountRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onMountRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffMountRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onMountRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMountRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onMountRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onMountRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMountRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onMountRequested.hasListener);
    },
    "call_HasOnMountRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onMountRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMountRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onMountRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMoveEntryRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onMoveEntryRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onMoveEntryRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMoveEntryRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onMoveEntryRequested.addListener);
    },
    "call_OnMoveEntryRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onMoveEntryRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnMoveEntryRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onMoveEntryRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMoveEntryRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onMoveEntryRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onMoveEntryRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMoveEntryRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onMoveEntryRequested.removeListener);
    },
    "call_OffMoveEntryRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onMoveEntryRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffMoveEntryRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onMoveEntryRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMoveEntryRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onMoveEntryRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onMoveEntryRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMoveEntryRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onMoveEntryRequested.hasListener);
    },
    "call_HasOnMoveEntryRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onMoveEntryRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMoveEntryRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onMoveEntryRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnOpenFileRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onOpenFileRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onOpenFileRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnOpenFileRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onOpenFileRequested.addListener);
    },
    "call_OnOpenFileRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onOpenFileRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnOpenFileRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onOpenFileRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffOpenFileRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onOpenFileRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onOpenFileRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffOpenFileRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onOpenFileRequested.removeListener);
    },
    "call_OffOpenFileRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onOpenFileRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffOpenFileRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onOpenFileRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnOpenFileRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onOpenFileRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onOpenFileRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnOpenFileRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onOpenFileRequested.hasListener);
    },
    "call_HasOnOpenFileRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onOpenFileRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnOpenFileRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onOpenFileRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnReadDirectoryRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onReadDirectoryRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onReadDirectoryRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnReadDirectoryRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onReadDirectoryRequested.addListener);
    },
    "call_OnReadDirectoryRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onReadDirectoryRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnReadDirectoryRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onReadDirectoryRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffReadDirectoryRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onReadDirectoryRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onReadDirectoryRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffReadDirectoryRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onReadDirectoryRequested.removeListener);
    },
    "call_OffReadDirectoryRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onReadDirectoryRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffReadDirectoryRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onReadDirectoryRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnReadDirectoryRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onReadDirectoryRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onReadDirectoryRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnReadDirectoryRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onReadDirectoryRequested.hasListener);
    },
    "call_HasOnReadDirectoryRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onReadDirectoryRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnReadDirectoryRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onReadDirectoryRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnReadFileRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onReadFileRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onReadFileRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnReadFileRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onReadFileRequested.addListener);
    },
    "call_OnReadFileRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onReadFileRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnReadFileRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onReadFileRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffReadFileRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onReadFileRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onReadFileRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffReadFileRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onReadFileRequested.removeListener);
    },
    "call_OffReadFileRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onReadFileRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffReadFileRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onReadFileRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnReadFileRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onReadFileRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onReadFileRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnReadFileRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onReadFileRequested.hasListener);
    },
    "call_HasOnReadFileRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onReadFileRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnReadFileRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onReadFileRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRemoveWatcherRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onRemoveWatcherRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onRemoveWatcherRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRemoveWatcherRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onRemoveWatcherRequested.addListener);
    },
    "call_OnRemoveWatcherRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onRemoveWatcherRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnRemoveWatcherRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onRemoveWatcherRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRemoveWatcherRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onRemoveWatcherRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onRemoveWatcherRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRemoveWatcherRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onRemoveWatcherRequested.removeListener);
    },
    "call_OffRemoveWatcherRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onRemoveWatcherRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffRemoveWatcherRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onRemoveWatcherRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRemoveWatcherRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onRemoveWatcherRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onRemoveWatcherRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRemoveWatcherRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onRemoveWatcherRequested.hasListener);
    },
    "call_HasOnRemoveWatcherRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onRemoveWatcherRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRemoveWatcherRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onRemoveWatcherRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTruncateRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onTruncateRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onTruncateRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTruncateRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onTruncateRequested.addListener);
    },
    "call_OnTruncateRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onTruncateRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnTruncateRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onTruncateRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffTruncateRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onTruncateRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onTruncateRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffTruncateRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onTruncateRequested.removeListener);
    },
    "call_OffTruncateRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onTruncateRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffTruncateRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onTruncateRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnTruncateRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onTruncateRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onTruncateRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnTruncateRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onTruncateRequested.hasListener);
    },
    "call_HasOnTruncateRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onTruncateRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnTruncateRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onTruncateRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnUnmountRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onUnmountRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onUnmountRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnUnmountRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onUnmountRequested.addListener);
    },
    "call_OnUnmountRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onUnmountRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnUnmountRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onUnmountRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffUnmountRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onUnmountRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onUnmountRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffUnmountRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onUnmountRequested.removeListener);
    },
    "call_OffUnmountRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onUnmountRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffUnmountRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onUnmountRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnUnmountRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onUnmountRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onUnmountRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnUnmountRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onUnmountRequested.hasListener);
    },
    "call_HasOnUnmountRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onUnmountRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnUnmountRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onUnmountRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnWriteFileRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onWriteFileRequested &&
        "addListener" in WEBEXT?.fileSystemProvider?.onWriteFileRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnWriteFileRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onWriteFileRequested.addListener);
    },
    "call_OnWriteFileRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onWriteFileRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnWriteFileRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onWriteFileRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffWriteFileRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onWriteFileRequested &&
        "removeListener" in WEBEXT?.fileSystemProvider?.onWriteFileRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffWriteFileRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onWriteFileRequested.removeListener);
    },
    "call_OffWriteFileRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onWriteFileRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffWriteFileRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onWriteFileRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnWriteFileRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.fileSystemProvider?.onWriteFileRequested &&
        "hasListener" in WEBEXT?.fileSystemProvider?.onWriteFileRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnWriteFileRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.onWriteFileRequested.hasListener);
    },
    "call_HasOnWriteFileRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileSystemProvider.onWriteFileRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnWriteFileRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProvider.onWriteFileRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Unmount": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystemProvider && "unmount" in WEBEXT?.fileSystemProvider) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Unmount": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProvider.unmount);
    },
    "call_Unmount": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["fileSystemId"] = A.load.Ref(options + 0, undefined);

      const _ret = WEBEXT.fileSystemProvider.unmount(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Unmount": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["fileSystemId"] = A.load.Ref(options + 0, undefined);

        const _ret = WEBEXT.fileSystemProvider.unmount(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
