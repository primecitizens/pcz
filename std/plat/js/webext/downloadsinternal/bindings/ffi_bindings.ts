import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/downloadsinternal", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_DetermineFilename": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloadsInternal && "determineFilename" in WEBEXT?.downloadsInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DetermineFilename": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloadsInternal.determineFilename);
    },
    "call_DetermineFilename": (
      retPtr: Pointer,
      downloadId: number,
      filename: heap.Ref<object>,
      conflict_action: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.downloadsInternal.determineFilename(
        downloadId,
        A.H.get<object>(filename),
        A.H.get<object>(conflict_action)
      );
    },
    "try_DetermineFilename": (
      retPtr: Pointer,
      errPtr: Pointer,
      downloadId: number,
      filename: heap.Ref<object>,
      conflict_action: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloadsInternal.determineFilename(
          downloadId,
          A.H.get<object>(filename),
          A.H.get<object>(conflict_action)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
