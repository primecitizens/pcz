import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/printerprovider", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_PrintError": (ref: heap.Ref<string>): number => {
      const idx = ["OK", "FAILED", "INVALID_TICKET", "INVALID_DATA"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PrintJob": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Ref(ptr + 0, x["printerId"]);
        A.store.Ref(ptr + 4, x["title"]);
        A.store.Ref(ptr + 8, x["ticket"]);
        A.store.Ref(ptr + 12, x["contentType"]);
        A.store.Ref(ptr + 16, x["document"]);
      }
    },
    "load_PrintJob": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["printerId"] = A.load.Ref(ptr + 0, undefined);
      x["title"] = A.load.Ref(ptr + 4, undefined);
      x["ticket"] = A.load.Ref(ptr + 8, undefined);
      x["contentType"] = A.load.Ref(ptr + 12, undefined);
      x["document"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PrinterInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Ref(ptr + 8, x["description"]);
      }
    },
    "load_PrinterInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      x["description"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_OnGetCapabilityRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.printerProvider?.onGetCapabilityRequested &&
        "addListener" in WEBEXT?.printerProvider?.onGetCapabilityRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnGetCapabilityRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProvider.onGetCapabilityRequested.addListener);
    },
    "call_OnGetCapabilityRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printerProvider.onGetCapabilityRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnGetCapabilityRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProvider.onGetCapabilityRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffGetCapabilityRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.printerProvider?.onGetCapabilityRequested &&
        "removeListener" in WEBEXT?.printerProvider?.onGetCapabilityRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffGetCapabilityRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProvider.onGetCapabilityRequested.removeListener);
    },
    "call_OffGetCapabilityRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printerProvider.onGetCapabilityRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffGetCapabilityRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProvider.onGetCapabilityRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnGetCapabilityRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.printerProvider?.onGetCapabilityRequested &&
        "hasListener" in WEBEXT?.printerProvider?.onGetCapabilityRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnGetCapabilityRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProvider.onGetCapabilityRequested.hasListener);
    },
    "call_HasOnGetCapabilityRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printerProvider.onGetCapabilityRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnGetCapabilityRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProvider.onGetCapabilityRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnGetPrintersRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.printerProvider?.onGetPrintersRequested &&
        "addListener" in WEBEXT?.printerProvider?.onGetPrintersRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnGetPrintersRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProvider.onGetPrintersRequested.addListener);
    },
    "call_OnGetPrintersRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printerProvider.onGetPrintersRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnGetPrintersRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProvider.onGetPrintersRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffGetPrintersRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.printerProvider?.onGetPrintersRequested &&
        "removeListener" in WEBEXT?.printerProvider?.onGetPrintersRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffGetPrintersRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProvider.onGetPrintersRequested.removeListener);
    },
    "call_OffGetPrintersRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printerProvider.onGetPrintersRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffGetPrintersRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProvider.onGetPrintersRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnGetPrintersRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.printerProvider?.onGetPrintersRequested &&
        "hasListener" in WEBEXT?.printerProvider?.onGetPrintersRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnGetPrintersRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProvider.onGetPrintersRequested.hasListener);
    },
    "call_HasOnGetPrintersRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printerProvider.onGetPrintersRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnGetPrintersRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProvider.onGetPrintersRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnGetUsbPrinterInfoRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.printerProvider?.onGetUsbPrinterInfoRequested &&
        "addListener" in WEBEXT?.printerProvider?.onGetUsbPrinterInfoRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnGetUsbPrinterInfoRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.addListener);
    },
    "call_OnGetUsbPrinterInfoRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnGetUsbPrinterInfoRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffGetUsbPrinterInfoRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.printerProvider?.onGetUsbPrinterInfoRequested &&
        "removeListener" in WEBEXT?.printerProvider?.onGetUsbPrinterInfoRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffGetUsbPrinterInfoRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.removeListener);
    },
    "call_OffGetUsbPrinterInfoRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffGetUsbPrinterInfoRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnGetUsbPrinterInfoRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.printerProvider?.onGetUsbPrinterInfoRequested &&
        "hasListener" in WEBEXT?.printerProvider?.onGetUsbPrinterInfoRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnGetUsbPrinterInfoRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.hasListener);
    },
    "call_HasOnGetUsbPrinterInfoRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnGetUsbPrinterInfoRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPrintRequested": (): heap.Ref<boolean> => {
      if (WEBEXT?.printerProvider?.onPrintRequested && "addListener" in WEBEXT?.printerProvider?.onPrintRequested) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPrintRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProvider.onPrintRequested.addListener);
    },
    "call_OnPrintRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printerProvider.onPrintRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnPrintRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProvider.onPrintRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPrintRequested": (): heap.Ref<boolean> => {
      if (WEBEXT?.printerProvider?.onPrintRequested && "removeListener" in WEBEXT?.printerProvider?.onPrintRequested) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPrintRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProvider.onPrintRequested.removeListener);
    },
    "call_OffPrintRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printerProvider.onPrintRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffPrintRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProvider.onPrintRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPrintRequested": (): heap.Ref<boolean> => {
      if (WEBEXT?.printerProvider?.onPrintRequested && "hasListener" in WEBEXT?.printerProvider?.onPrintRequested) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPrintRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProvider.onPrintRequested.hasListener);
    },
    "call_HasOnPrintRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printerProvider.onPrintRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPrintRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProvider.onPrintRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
