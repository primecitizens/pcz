import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/mimehandlerprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_StreamInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 23, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 21, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 22, false);
        A.store.Bool(ptr + 20, false);
      } else {
        A.store.Bool(ptr + 23, true);
        A.store.Ref(ptr + 0, x["mimeType"]);
        A.store.Ref(ptr + 4, x["originalUrl"]);
        A.store.Ref(ptr + 8, x["streamUrl"]);
        A.store.Bool(ptr + 21, "tabId" in x ? true : false);
        A.store.Int32(ptr + 12, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Ref(ptr + 16, x["responseHeaders"]);
        A.store.Bool(ptr + 22, "embedded" in x ? true : false);
        A.store.Bool(ptr + 20, x["embedded"] ? true : false);
      }
    },
    "load_StreamInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["mimeType"] = A.load.Ref(ptr + 0, undefined);
      x["originalUrl"] = A.load.Ref(ptr + 4, undefined);
      x["streamUrl"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 21)) {
        x["tabId"] = A.load.Int32(ptr + 12);
      } else {
        delete x["tabId"];
      }
      x["responseHeaders"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 22)) {
        x["embedded"] = A.load.Bool(ptr + 20);
      } else {
        delete x["embedded"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PdfPluginAttributes": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 9, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 11, true);
        A.store.Bool(ptr + 9, "backgroundColor" in x ? true : false);
        A.store.Float64(ptr + 0, x["backgroundColor"] === undefined ? 0 : (x["backgroundColor"] as number));
        A.store.Bool(ptr + 10, "allowJavascript" in x ? true : false);
        A.store.Bool(ptr + 8, x["allowJavascript"] ? true : false);
      }
    },
    "load_PdfPluginAttributes": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 9)) {
        x["backgroundColor"] = A.load.Float64(ptr + 0);
      } else {
        delete x["backgroundColor"];
      }
      if (A.load.Bool(ptr + 10)) {
        x["allowJavascript"] = A.load.Bool(ptr + 8);
      } else {
        delete x["allowJavascript"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetStreamInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.mimeHandlerPrivate && "getStreamInfo" in WEBEXT?.mimeHandlerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetStreamInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mimeHandlerPrivate.getStreamInfo);
    },
    "call_GetStreamInfo": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mimeHandlerPrivate.getStreamInfo(A.H.get<object>(callback));
    },
    "try_GetStreamInfo": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mimeHandlerPrivate.getStreamInfo(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSave": (): heap.Ref<boolean> => {
      if (WEBEXT?.mimeHandlerPrivate?.onSave && "addListener" in WEBEXT?.mimeHandlerPrivate?.onSave) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSave": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mimeHandlerPrivate.onSave.addListener);
    },
    "call_OnSave": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mimeHandlerPrivate.onSave.addListener(A.H.get<object>(callback));
    },
    "try_OnSave": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mimeHandlerPrivate.onSave.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSave": (): heap.Ref<boolean> => {
      if (WEBEXT?.mimeHandlerPrivate?.onSave && "removeListener" in WEBEXT?.mimeHandlerPrivate?.onSave) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSave": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mimeHandlerPrivate.onSave.removeListener);
    },
    "call_OffSave": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mimeHandlerPrivate.onSave.removeListener(A.H.get<object>(callback));
    },
    "try_OffSave": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mimeHandlerPrivate.onSave.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSave": (): heap.Ref<boolean> => {
      if (WEBEXT?.mimeHandlerPrivate?.onSave && "hasListener" in WEBEXT?.mimeHandlerPrivate?.onSave) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSave": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mimeHandlerPrivate.onSave.hasListener);
    },
    "call_HasOnSave": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mimeHandlerPrivate.onSave.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSave": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mimeHandlerPrivate.onSave.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPdfPluginAttributes": (): heap.Ref<boolean> => {
      if (WEBEXT?.mimeHandlerPrivate && "setPdfPluginAttributes" in WEBEXT?.mimeHandlerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPdfPluginAttributes": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mimeHandlerPrivate.setPdfPluginAttributes);
    },
    "call_SetPdfPluginAttributes": (retPtr: Pointer, pdfPluginAttributes: Pointer): void => {
      const pdfPluginAttributes_ffi = {};

      if (A.load.Bool(pdfPluginAttributes + 9)) {
        pdfPluginAttributes_ffi["backgroundColor"] = A.load.Float64(pdfPluginAttributes + 0);
      }
      if (A.load.Bool(pdfPluginAttributes + 10)) {
        pdfPluginAttributes_ffi["allowJavascript"] = A.load.Bool(pdfPluginAttributes + 8);
      }

      const _ret = WEBEXT.mimeHandlerPrivate.setPdfPluginAttributes(pdfPluginAttributes_ffi);
    },
    "try_SetPdfPluginAttributes": (
      retPtr: Pointer,
      errPtr: Pointer,
      pdfPluginAttributes: Pointer
    ): heap.Ref<boolean> => {
      try {
        const pdfPluginAttributes_ffi = {};

        if (A.load.Bool(pdfPluginAttributes + 9)) {
          pdfPluginAttributes_ffi["backgroundColor"] = A.load.Float64(pdfPluginAttributes + 0);
        }
        if (A.load.Bool(pdfPluginAttributes + 10)) {
          pdfPluginAttributes_ffi["allowJavascript"] = A.load.Bool(pdfPluginAttributes + 8);
        }

        const _ret = WEBEXT.mimeHandlerPrivate.setPdfPluginAttributes(pdfPluginAttributes_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetShowBeforeUnloadDialog": (): heap.Ref<boolean> => {
      if (WEBEXT?.mimeHandlerPrivate && "setShowBeforeUnloadDialog" in WEBEXT?.mimeHandlerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetShowBeforeUnloadDialog": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mimeHandlerPrivate.setShowBeforeUnloadDialog);
    },
    "call_SetShowBeforeUnloadDialog": (
      retPtr: Pointer,
      showDialog: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.mimeHandlerPrivate.setShowBeforeUnloadDialog(
        showDialog === A.H.TRUE,
        A.H.get<object>(callback)
      );
    },
    "try_SetShowBeforeUnloadDialog": (
      retPtr: Pointer,
      errPtr: Pointer,
      showDialog: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mimeHandlerPrivate.setShowBeforeUnloadDialog(
          showDialog === A.H.TRUE,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
