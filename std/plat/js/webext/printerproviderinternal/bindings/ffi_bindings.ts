import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/printerproviderinternal", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_PrintError": (ref: heap.Ref<string>): number => {
      const idx = ["OK", "FAILED", "INVALID_TICKET", "INVALID_DATA"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_GetPrintData": (): heap.Ref<boolean> => {
      if (WEBEXT?.printerProviderInternal && "getPrintData" in WEBEXT?.printerProviderInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPrintData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProviderInternal.getPrintData);
    },
    "call_GetPrintData": (retPtr: Pointer, requestId: number): void => {
      const _ret = WEBEXT.printerProviderInternal.getPrintData(requestId);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPrintData": (retPtr: Pointer, errPtr: Pointer, requestId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProviderInternal.getPrintData(requestId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReportPrintResult": (): heap.Ref<boolean> => {
      if (WEBEXT?.printerProviderInternal && "reportPrintResult" in WEBEXT?.printerProviderInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportPrintResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProviderInternal.reportPrintResult);
    },
    "call_ReportPrintResult": (retPtr: Pointer, request_id: number, error: number): void => {
      const _ret = WEBEXT.printerProviderInternal.reportPrintResult(
        request_id,
        error > 0 && error <= 4 ? ["OK", "FAILED", "INVALID_TICKET", "INVALID_DATA"][error - 1] : undefined
      );
    },
    "try_ReportPrintResult": (
      retPtr: Pointer,
      errPtr: Pointer,
      request_id: number,
      error: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProviderInternal.reportPrintResult(
          request_id,
          error > 0 && error <= 4 ? ["OK", "FAILED", "INVALID_TICKET", "INVALID_DATA"][error - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReportPrinterCapability": (): heap.Ref<boolean> => {
      if (WEBEXT?.printerProviderInternal && "reportPrinterCapability" in WEBEXT?.printerProviderInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportPrinterCapability": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProviderInternal.reportPrinterCapability);
    },
    "call_ReportPrinterCapability": (retPtr: Pointer, request_id: number, capability: heap.Ref<object>): void => {
      const _ret = WEBEXT.printerProviderInternal.reportPrinterCapability(request_id, A.H.get<object>(capability));
    },
    "try_ReportPrinterCapability": (
      retPtr: Pointer,
      errPtr: Pointer,
      request_id: number,
      capability: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProviderInternal.reportPrinterCapability(request_id, A.H.get<object>(capability));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReportPrinters": (): heap.Ref<boolean> => {
      if (WEBEXT?.printerProviderInternal && "reportPrinters" in WEBEXT?.printerProviderInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportPrinters": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProviderInternal.reportPrinters);
    },
    "call_ReportPrinters": (retPtr: Pointer, requestId: number, printers: heap.Ref<object>): void => {
      const _ret = WEBEXT.printerProviderInternal.reportPrinters(requestId, A.H.get<object>(printers));
    },
    "try_ReportPrinters": (
      retPtr: Pointer,
      errPtr: Pointer,
      requestId: number,
      printers: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printerProviderInternal.reportPrinters(requestId, A.H.get<object>(printers));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReportUsbPrinterInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.printerProviderInternal && "reportUsbPrinterInfo" in WEBEXT?.printerProviderInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportUsbPrinterInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printerProviderInternal.reportUsbPrinterInfo);
    },
    "call_ReportUsbPrinterInfo": (retPtr: Pointer, requestId: number, printerInfo: Pointer): void => {
      const printerInfo_ffi = {};

      printerInfo_ffi["id"] = A.load.Ref(printerInfo + 0, undefined);
      printerInfo_ffi["name"] = A.load.Ref(printerInfo + 4, undefined);
      printerInfo_ffi["description"] = A.load.Ref(printerInfo + 8, undefined);

      const _ret = WEBEXT.printerProviderInternal.reportUsbPrinterInfo(requestId, printerInfo_ffi);
    },
    "try_ReportUsbPrinterInfo": (
      retPtr: Pointer,
      errPtr: Pointer,
      requestId: number,
      printerInfo: Pointer
    ): heap.Ref<boolean> => {
      try {
        const printerInfo_ffi = {};

        printerInfo_ffi["id"] = A.load.Ref(printerInfo + 0, undefined);
        printerInfo_ffi["name"] = A.load.Ref(printerInfo + 4, undefined);
        printerInfo_ffi["description"] = A.load.Ref(printerInfo + 8, undefined);

        const _ret = WEBEXT.printerProviderInternal.reportUsbPrinterInfo(requestId, printerInfo_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
