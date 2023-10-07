import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/pdfviewerprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_IsAllowedLocalFileAccess": (): heap.Ref<boolean> => {
      if (WEBEXT?.pdfViewerPrivate && "isAllowedLocalFileAccess" in WEBEXT?.pdfViewerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsAllowedLocalFileAccess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pdfViewerPrivate.isAllowedLocalFileAccess);
    },
    "call_IsAllowedLocalFileAccess": (retPtr: Pointer, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.pdfViewerPrivate.isAllowedLocalFileAccess(A.H.get<object>(url));
      A.store.Ref(retPtr, _ret);
    },
    "try_IsAllowedLocalFileAccess": (retPtr: Pointer, errPtr: Pointer, url: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.pdfViewerPrivate.isAllowedLocalFileAccess(A.H.get<object>(url));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsPdfOcrAlwaysActive": (): heap.Ref<boolean> => {
      if (WEBEXT?.pdfViewerPrivate && "isPdfOcrAlwaysActive" in WEBEXT?.pdfViewerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsPdfOcrAlwaysActive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pdfViewerPrivate.isPdfOcrAlwaysActive);
    },
    "call_IsPdfOcrAlwaysActive": (retPtr: Pointer): void => {
      const _ret = WEBEXT.pdfViewerPrivate.isPdfOcrAlwaysActive();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsPdfOcrAlwaysActive": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.pdfViewerPrivate.isPdfOcrAlwaysActive();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPdfOcrPrefChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.pdfViewerPrivate?.onPdfOcrPrefChanged &&
        "addListener" in WEBEXT?.pdfViewerPrivate?.onPdfOcrPrefChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPdfOcrPrefChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.addListener);
    },
    "call_OnPdfOcrPrefChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnPdfOcrPrefChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPdfOcrPrefChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.pdfViewerPrivate?.onPdfOcrPrefChanged &&
        "removeListener" in WEBEXT?.pdfViewerPrivate?.onPdfOcrPrefChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPdfOcrPrefChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.removeListener);
    },
    "call_OffPdfOcrPrefChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffPdfOcrPrefChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPdfOcrPrefChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.pdfViewerPrivate?.onPdfOcrPrefChanged &&
        "hasListener" in WEBEXT?.pdfViewerPrivate?.onPdfOcrPrefChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPdfOcrPrefChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.hasListener);
    },
    "call_HasOnPdfOcrPrefChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPdfOcrPrefChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPdfOcrPref": (): heap.Ref<boolean> => {
      if (WEBEXT?.pdfViewerPrivate && "setPdfOcrPref" in WEBEXT?.pdfViewerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPdfOcrPref": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pdfViewerPrivate.setPdfOcrPref);
    },
    "call_SetPdfOcrPref": (retPtr: Pointer, value: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.pdfViewerPrivate.setPdfOcrPref(value === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetPdfOcrPref": (retPtr: Pointer, errPtr: Pointer, value: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.pdfViewerPrivate.setPdfOcrPref(value === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
