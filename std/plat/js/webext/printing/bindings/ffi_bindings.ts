import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/printing", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_PrinterStatus": (ref: heap.Ref<string>): number => {
      const idx = [
        "DOOR_OPEN",
        "TRAY_MISSING",
        "OUT_OF_INK",
        "OUT_OF_PAPER",
        "OUTPUT_FULL",
        "PAPER_JAM",
        "GENERIC_ISSUE",
        "STOPPED",
        "UNREACHABLE",
        "EXPIRED_CERTIFICATE",
        "AVAILABLE",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_GetPrinterInfoResponse": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["capabilities"]);
        A.store.Enum(
          ptr + 4,
          [
            "DOOR_OPEN",
            "TRAY_MISSING",
            "OUT_OF_INK",
            "OUT_OF_PAPER",
            "OUTPUT_FULL",
            "PAPER_JAM",
            "GENERIC_ISSUE",
            "STOPPED",
            "UNREACHABLE",
            "EXPIRED_CERTIFICATE",
            "AVAILABLE",
          ].indexOf(x["status"] as string)
        );
      }
    },
    "load_GetPrinterInfoResponse": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["capabilities"] = A.load.Ref(ptr + 0, undefined);
      x["status"] = A.load.Enum(ptr + 4, [
        "DOOR_OPEN",
        "TRAY_MISSING",
        "OUT_OF_INK",
        "OUT_OF_PAPER",
        "OUTPUT_FULL",
        "PAPER_JAM",
        "GENERIC_ISSUE",
        "STOPPED",
        "UNREACHABLE",
        "EXPIRED_CERTIFICATE",
        "AVAILABLE",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PrinterSource": (ref: heap.Ref<string>): number => {
      const idx = ["USER", "POLICY"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Printer": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 30, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Enum(ptr + 16, -1);
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 29, false);
        A.store.Int32(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 30, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Ref(ptr + 8, x["description"]);
        A.store.Ref(ptr + 12, x["uri"]);
        A.store.Enum(ptr + 16, ["USER", "POLICY"].indexOf(x["source"] as string));
        A.store.Bool(ptr + 28, "isDefault" in x ? true : false);
        A.store.Bool(ptr + 20, x["isDefault"] ? true : false);
        A.store.Bool(ptr + 29, "recentlyUsedRank" in x ? true : false);
        A.store.Int32(ptr + 24, x["recentlyUsedRank"] === undefined ? 0 : (x["recentlyUsedRank"] as number));
      }
    },
    "load_Printer": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      x["description"] = A.load.Ref(ptr + 8, undefined);
      x["uri"] = A.load.Ref(ptr + 12, undefined);
      x["source"] = A.load.Enum(ptr + 16, ["USER", "POLICY"]);
      if (A.load.Bool(ptr + 28)) {
        x["isDefault"] = A.load.Bool(ptr + 20);
      } else {
        delete x["isDefault"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["recentlyUsedRank"] = A.load.Int32(ptr + 24);
      } else {
        delete x["recentlyUsedRank"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_JobStatus": (ref: heap.Ref<string>): number => {
      const idx = ["PENDING", "IN_PROGRESS", "FAILED", "CANCELED", "PRINTED"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "has_Properties_MAX_SUBMIT_JOB_CALLS_PER_MINUTE": (self: heap.Ref<any>): heap.Ref<boolean> => {
      if (WEBEXT?.printing && "MAX_SUBMIT_JOB_CALLS_PER_MINUTE" in WEBEXT?.printing) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Properties_MAX_SUBMIT_JOB_CALLS_PER_MINUTE": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printing.MAX_SUBMIT_JOB_CALLS_PER_MINUTE);
    },

    "call_Properties_MAX_SUBMIT_JOB_CALLS_PER_MINUTE": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = Reflect.apply(WEBEXT.printing.MAX_SUBMIT_JOB_CALLS_PER_MINUTE, thiz, []);
      A.store.Int32(retPtr, _ret);
    },
    "try_Properties_MAX_SUBMIT_JOB_CALLS_PER_MINUTE": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = Reflect.apply(WEBEXT.printing.MAX_SUBMIT_JOB_CALLS_PER_MINUTE, thiz, []);
        A.store.Int32(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Properties_MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE": (self: heap.Ref<any>): heap.Ref<boolean> => {
      if (WEBEXT?.printing && "MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE" in WEBEXT?.printing) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Properties_MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printing.MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE);
    },

    "call_Properties_MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = Reflect.apply(WEBEXT.printing.MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE, thiz, []);
      A.store.Int32(retPtr, _ret);
    },
    "try_Properties_MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = Reflect.apply(WEBEXT.printing.MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE, thiz, []);
        A.store.Int32(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "constof_SubmitJobStatus": (ref: heap.Ref<string>): number => {
      const idx = ["OK", "USER_REJECTED"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SubmitJobResponse": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["OK", "USER_REJECTED"].indexOf(x["status"] as string));
        A.store.Ref(ptr + 4, x["jobId"]);
      }
    },
    "load_SubmitJobResponse": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["status"] = A.load.Enum(ptr + 0, ["OK", "USER_REJECTED"]);
      x["jobId"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SubmitJobRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);

        A.store.Bool(ptr + 0 + 20, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Ref(ptr + 0 + 8, undefined);
        A.store.Ref(ptr + 0 + 12, undefined);
        A.store.Ref(ptr + 0 + 16, undefined);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 28, true);

        if (typeof x["job"] === "undefined") {
          A.store.Bool(ptr + 0 + 20, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Ref(ptr + 0 + 8, undefined);
          A.store.Ref(ptr + 0 + 12, undefined);
          A.store.Ref(ptr + 0 + 16, undefined);
        } else {
          A.store.Bool(ptr + 0 + 20, true);
          A.store.Ref(ptr + 0 + 0, x["job"]["printerId"]);
          A.store.Ref(ptr + 0 + 4, x["job"]["title"]);
          A.store.Ref(ptr + 0 + 8, x["job"]["ticket"]);
          A.store.Ref(ptr + 0 + 12, x["job"]["contentType"]);
          A.store.Ref(ptr + 0 + 16, x["job"]["document"]);
        }
        A.store.Ref(ptr + 24, x["documentBlobUuid"]);
      }
    },
    "load_SubmitJobRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 20)) {
        x["job"] = {};
        x["job"]["printerId"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["job"]["title"] = A.load.Ref(ptr + 0 + 4, undefined);
        x["job"]["ticket"] = A.load.Ref(ptr + 0 + 8, undefined);
        x["job"]["contentType"] = A.load.Ref(ptr + 0 + 12, undefined);
        x["job"]["document"] = A.load.Ref(ptr + 0 + 16, undefined);
      } else {
        delete x["job"];
      }
      x["documentBlobUuid"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_CancelJob": (): heap.Ref<boolean> => {
      if (WEBEXT?.printing && "cancelJob" in WEBEXT?.printing) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CancelJob": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printing.cancelJob);
    },
    "call_CancelJob": (retPtr: Pointer, jobId: heap.Ref<object>): void => {
      const _ret = WEBEXT.printing.cancelJob(A.H.get<object>(jobId));
      A.store.Ref(retPtr, _ret);
    },
    "try_CancelJob": (retPtr: Pointer, errPtr: Pointer, jobId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printing.cancelJob(A.H.get<object>(jobId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPrinterInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.printing && "getPrinterInfo" in WEBEXT?.printing) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPrinterInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printing.getPrinterInfo);
    },
    "call_GetPrinterInfo": (retPtr: Pointer, printerId: heap.Ref<object>): void => {
      const _ret = WEBEXT.printing.getPrinterInfo(A.H.get<object>(printerId));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPrinterInfo": (retPtr: Pointer, errPtr: Pointer, printerId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printing.getPrinterInfo(A.H.get<object>(printerId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPrinters": (): heap.Ref<boolean> => {
      if (WEBEXT?.printing && "getPrinters" in WEBEXT?.printing) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPrinters": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printing.getPrinters);
    },
    "call_GetPrinters": (retPtr: Pointer): void => {
      const _ret = WEBEXT.printing.getPrinters();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPrinters": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printing.getPrinters();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnJobStatusChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.printing?.onJobStatusChanged && "addListener" in WEBEXT?.printing?.onJobStatusChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnJobStatusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printing.onJobStatusChanged.addListener);
    },
    "call_OnJobStatusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printing.onJobStatusChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnJobStatusChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printing.onJobStatusChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffJobStatusChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.printing?.onJobStatusChanged && "removeListener" in WEBEXT?.printing?.onJobStatusChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffJobStatusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printing.onJobStatusChanged.removeListener);
    },
    "call_OffJobStatusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printing.onJobStatusChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffJobStatusChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printing.onJobStatusChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnJobStatusChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.printing?.onJobStatusChanged && "hasListener" in WEBEXT?.printing?.onJobStatusChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnJobStatusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printing.onJobStatusChanged.hasListener);
    },
    "call_HasOnJobStatusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printing.onJobStatusChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnJobStatusChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printing.onJobStatusChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SubmitJob": (): heap.Ref<boolean> => {
      if (WEBEXT?.printing && "submitJob" in WEBEXT?.printing) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SubmitJob": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printing.submitJob);
    },
    "call_SubmitJob": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 0 + 20)) {
        request_ffi["job"] = {};
        request_ffi["job"]["printerId"] = A.load.Ref(request + 0 + 0, undefined);
        request_ffi["job"]["title"] = A.load.Ref(request + 0 + 4, undefined);
        request_ffi["job"]["ticket"] = A.load.Ref(request + 0 + 8, undefined);
        request_ffi["job"]["contentType"] = A.load.Ref(request + 0 + 12, undefined);
        request_ffi["job"]["document"] = A.load.Ref(request + 0 + 16, undefined);
      }
      request_ffi["documentBlobUuid"] = A.load.Ref(request + 24, undefined);

      const _ret = WEBEXT.printing.submitJob(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SubmitJob": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 0 + 20)) {
          request_ffi["job"] = {};
          request_ffi["job"]["printerId"] = A.load.Ref(request + 0 + 0, undefined);
          request_ffi["job"]["title"] = A.load.Ref(request + 0 + 4, undefined);
          request_ffi["job"]["ticket"] = A.load.Ref(request + 0 + 8, undefined);
          request_ffi["job"]["contentType"] = A.load.Ref(request + 0 + 12, undefined);
          request_ffi["job"]["document"] = A.load.Ref(request + 0 + 16, undefined);
        }
        request_ffi["documentBlobUuid"] = A.load.Ref(request + 24, undefined);

        const _ret = WEBEXT.printing.submitJob(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
