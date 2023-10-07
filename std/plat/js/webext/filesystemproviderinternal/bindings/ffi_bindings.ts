import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/filesystemproviderinternal", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_GetActionsRequestedSuccess": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystemProviderInternal && "getActionsRequestedSuccess" in WEBEXT?.fileSystemProviderInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetActionsRequestedSuccess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProviderInternal.getActionsRequestedSuccess);
    },
    "call_GetActionsRequestedSuccess": (
      retPtr: Pointer,
      fileSystemId: heap.Ref<object>,
      requestId: number,
      actions: heap.Ref<object>,
      executionTime: number
    ): void => {
      const _ret = WEBEXT.fileSystemProviderInternal.getActionsRequestedSuccess(
        A.H.get<object>(fileSystemId),
        requestId,
        A.H.get<object>(actions),
        executionTime
      );
    },
    "try_GetActionsRequestedSuccess": (
      retPtr: Pointer,
      errPtr: Pointer,
      fileSystemId: heap.Ref<object>,
      requestId: number,
      actions: heap.Ref<object>,
      executionTime: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProviderInternal.getActionsRequestedSuccess(
          A.H.get<object>(fileSystemId),
          requestId,
          A.H.get<object>(actions),
          executionTime
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetMetadataRequestedSuccess": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystemProviderInternal && "getMetadataRequestedSuccess" in WEBEXT?.fileSystemProviderInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetMetadataRequestedSuccess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProviderInternal.getMetadataRequestedSuccess);
    },
    "call_GetMetadataRequestedSuccess": (
      retPtr: Pointer,
      fileSystemId: heap.Ref<object>,
      requestId: number,
      metadata: Pointer,
      executionTime: number
    ): void => {
      const metadata_ffi = {};

      if (A.load.Bool(metadata + 37)) {
        metadata_ffi["isDirectory"] = A.load.Bool(metadata + 0);
      }
      metadata_ffi["name"] = A.load.Ref(metadata + 4, undefined);
      if (A.load.Bool(metadata + 38)) {
        metadata_ffi["size"] = A.load.Float64(metadata + 8);
      }
      metadata_ffi["modificationTime"] = A.load.Ref(metadata + 16, undefined);
      metadata_ffi["mimeType"] = A.load.Ref(metadata + 20, undefined);
      metadata_ffi["thumbnail"] = A.load.Ref(metadata + 24, undefined);
      if (A.load.Bool(metadata + 28 + 8)) {
        metadata_ffi["cloudIdentifier"] = {};
        metadata_ffi["cloudIdentifier"]["providerName"] = A.load.Ref(metadata + 28 + 0, undefined);
        metadata_ffi["cloudIdentifier"]["id"] = A.load.Ref(metadata + 28 + 4, undefined);
      }

      const _ret = WEBEXT.fileSystemProviderInternal.getMetadataRequestedSuccess(
        A.H.get<object>(fileSystemId),
        requestId,
        metadata_ffi,
        executionTime
      );
    },
    "try_GetMetadataRequestedSuccess": (
      retPtr: Pointer,
      errPtr: Pointer,
      fileSystemId: heap.Ref<object>,
      requestId: number,
      metadata: Pointer,
      executionTime: number
    ): heap.Ref<boolean> => {
      try {
        const metadata_ffi = {};

        if (A.load.Bool(metadata + 37)) {
          metadata_ffi["isDirectory"] = A.load.Bool(metadata + 0);
        }
        metadata_ffi["name"] = A.load.Ref(metadata + 4, undefined);
        if (A.load.Bool(metadata + 38)) {
          metadata_ffi["size"] = A.load.Float64(metadata + 8);
        }
        metadata_ffi["modificationTime"] = A.load.Ref(metadata + 16, undefined);
        metadata_ffi["mimeType"] = A.load.Ref(metadata + 20, undefined);
        metadata_ffi["thumbnail"] = A.load.Ref(metadata + 24, undefined);
        if (A.load.Bool(metadata + 28 + 8)) {
          metadata_ffi["cloudIdentifier"] = {};
          metadata_ffi["cloudIdentifier"]["providerName"] = A.load.Ref(metadata + 28 + 0, undefined);
          metadata_ffi["cloudIdentifier"]["id"] = A.load.Ref(metadata + 28 + 4, undefined);
        }

        const _ret = WEBEXT.fileSystemProviderInternal.getMetadataRequestedSuccess(
          A.H.get<object>(fileSystemId),
          requestId,
          metadata_ffi,
          executionTime
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OperationRequestedError": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystemProviderInternal && "operationRequestedError" in WEBEXT?.fileSystemProviderInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OperationRequestedError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProviderInternal.operationRequestedError);
    },
    "call_OperationRequestedError": (
      retPtr: Pointer,
      fileSystemId: heap.Ref<object>,
      requestId: number,
      error: number,
      executionTime: number
    ): void => {
      const _ret = WEBEXT.fileSystemProviderInternal.operationRequestedError(
        A.H.get<object>(fileSystemId),
        requestId,
        error > 0 && error <= 17
          ? [
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
            ][error - 1]
          : undefined,
        executionTime
      );
    },
    "try_OperationRequestedError": (
      retPtr: Pointer,
      errPtr: Pointer,
      fileSystemId: heap.Ref<object>,
      requestId: number,
      error: number,
      executionTime: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProviderInternal.operationRequestedError(
          A.H.get<object>(fileSystemId),
          requestId,
          error > 0 && error <= 17
            ? [
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
              ][error - 1]
            : undefined,
          executionTime
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OperationRequestedSuccess": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystemProviderInternal && "operationRequestedSuccess" in WEBEXT?.fileSystemProviderInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OperationRequestedSuccess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProviderInternal.operationRequestedSuccess);
    },
    "call_OperationRequestedSuccess": (
      retPtr: Pointer,
      fileSystemId: heap.Ref<object>,
      requestId: number,
      executionTime: number
    ): void => {
      const _ret = WEBEXT.fileSystemProviderInternal.operationRequestedSuccess(
        A.H.get<object>(fileSystemId),
        requestId,
        executionTime
      );
    },
    "try_OperationRequestedSuccess": (
      retPtr: Pointer,
      errPtr: Pointer,
      fileSystemId: heap.Ref<object>,
      requestId: number,
      executionTime: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProviderInternal.operationRequestedSuccess(
          A.H.get<object>(fileSystemId),
          requestId,
          executionTime
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReadDirectoryRequestedSuccess": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystemProviderInternal && "readDirectoryRequestedSuccess" in WEBEXT?.fileSystemProviderInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReadDirectoryRequestedSuccess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProviderInternal.readDirectoryRequestedSuccess);
    },
    "call_ReadDirectoryRequestedSuccess": (
      retPtr: Pointer,
      fileSystemId: heap.Ref<object>,
      requestId: number,
      entries: heap.Ref<object>,
      hasMore: heap.Ref<boolean>,
      executionTime: number
    ): void => {
      const _ret = WEBEXT.fileSystemProviderInternal.readDirectoryRequestedSuccess(
        A.H.get<object>(fileSystemId),
        requestId,
        A.H.get<object>(entries),
        hasMore === A.H.TRUE,
        executionTime
      );
    },
    "try_ReadDirectoryRequestedSuccess": (
      retPtr: Pointer,
      errPtr: Pointer,
      fileSystemId: heap.Ref<object>,
      requestId: number,
      entries: heap.Ref<object>,
      hasMore: heap.Ref<boolean>,
      executionTime: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProviderInternal.readDirectoryRequestedSuccess(
          A.H.get<object>(fileSystemId),
          requestId,
          A.H.get<object>(entries),
          hasMore === A.H.TRUE,
          executionTime
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReadFileRequestedSuccess": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystemProviderInternal && "readFileRequestedSuccess" in WEBEXT?.fileSystemProviderInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReadFileRequestedSuccess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProviderInternal.readFileRequestedSuccess);
    },
    "call_ReadFileRequestedSuccess": (
      retPtr: Pointer,
      fileSystemId: heap.Ref<object>,
      requestId: number,
      data: heap.Ref<object>,
      hasMore: heap.Ref<boolean>,
      executionTime: number
    ): void => {
      const _ret = WEBEXT.fileSystemProviderInternal.readFileRequestedSuccess(
        A.H.get<object>(fileSystemId),
        requestId,
        A.H.get<object>(data),
        hasMore === A.H.TRUE,
        executionTime
      );
    },
    "try_ReadFileRequestedSuccess": (
      retPtr: Pointer,
      errPtr: Pointer,
      fileSystemId: heap.Ref<object>,
      requestId: number,
      data: heap.Ref<object>,
      hasMore: heap.Ref<boolean>,
      executionTime: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProviderInternal.readFileRequestedSuccess(
          A.H.get<object>(fileSystemId),
          requestId,
          A.H.get<object>(data),
          hasMore === A.H.TRUE,
          executionTime
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RespondToMountRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystemProviderInternal && "respondToMountRequest" in WEBEXT?.fileSystemProviderInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RespondToMountRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProviderInternal.respondToMountRequest);
    },
    "call_RespondToMountRequest": (retPtr: Pointer, requestId: number, error: number, executionTime: number): void => {
      const _ret = WEBEXT.fileSystemProviderInternal.respondToMountRequest(
        requestId,
        error > 0 && error <= 17
          ? [
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
            ][error - 1]
          : undefined,
        executionTime
      );
    },
    "try_RespondToMountRequest": (
      retPtr: Pointer,
      errPtr: Pointer,
      requestId: number,
      error: number,
      executionTime: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProviderInternal.respondToMountRequest(
          requestId,
          error > 0 && error <= 17
            ? [
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
              ][error - 1]
            : undefined,
          executionTime
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UnmountRequestedSuccess": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileSystemProviderInternal && "unmountRequestedSuccess" in WEBEXT?.fileSystemProviderInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UnmountRequestedSuccess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileSystemProviderInternal.unmountRequestedSuccess);
    },
    "call_UnmountRequestedSuccess": (
      retPtr: Pointer,
      fileSystemId: heap.Ref<object>,
      requestId: number,
      executionTime: number
    ): void => {
      const _ret = WEBEXT.fileSystemProviderInternal.unmountRequestedSuccess(
        A.H.get<object>(fileSystemId),
        requestId,
        executionTime
      );
    },
    "try_UnmountRequestedSuccess": (
      retPtr: Pointer,
      errPtr: Pointer,
      fileSystemId: heap.Ref<object>,
      requestId: number,
      executionTime: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileSystemProviderInternal.unmountRequestedSuccess(
          A.H.get<object>(fileSystemId),
          requestId,
          executionTime
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
