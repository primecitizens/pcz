import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/crashreportprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_ErrorInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 34, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 32, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Bool(ptr + 33, false);
        A.store.Int32(ptr + 20, 0);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
      } else {
        A.store.Bool(ptr + 34, true);
        A.store.Ref(ptr + 0, x["message"]);
        A.store.Ref(ptr + 4, x["url"]);
        A.store.Ref(ptr + 8, x["product"]);
        A.store.Ref(ptr + 12, x["version"]);
        A.store.Bool(ptr + 32, "lineNumber" in x ? true : false);
        A.store.Int32(ptr + 16, x["lineNumber"] === undefined ? 0 : (x["lineNumber"] as number));
        A.store.Bool(ptr + 33, "columnNumber" in x ? true : false);
        A.store.Int32(ptr + 20, x["columnNumber"] === undefined ? 0 : (x["columnNumber"] as number));
        A.store.Ref(ptr + 24, x["debugId"]);
        A.store.Ref(ptr + 28, x["stackTrace"]);
      }
    },
    "load_ErrorInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["message"] = A.load.Ref(ptr + 0, undefined);
      x["url"] = A.load.Ref(ptr + 4, undefined);
      x["product"] = A.load.Ref(ptr + 8, undefined);
      x["version"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 32)) {
        x["lineNumber"] = A.load.Int32(ptr + 16);
      } else {
        delete x["lineNumber"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["columnNumber"] = A.load.Int32(ptr + 20);
      } else {
        delete x["columnNumber"];
      }
      x["debugId"] = A.load.Ref(ptr + 24, undefined);
      x["stackTrace"] = A.load.Ref(ptr + 28, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_ReportError": (): heap.Ref<boolean> => {
      if (WEBEXT?.crashReportPrivate && "reportError" in WEBEXT?.crashReportPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.crashReportPrivate.reportError);
    },
    "call_ReportError": (retPtr: Pointer, info: Pointer): void => {
      const info_ffi = {};

      info_ffi["message"] = A.load.Ref(info + 0, undefined);
      info_ffi["url"] = A.load.Ref(info + 4, undefined);
      info_ffi["product"] = A.load.Ref(info + 8, undefined);
      info_ffi["version"] = A.load.Ref(info + 12, undefined);
      if (A.load.Bool(info + 32)) {
        info_ffi["lineNumber"] = A.load.Int32(info + 16);
      }
      if (A.load.Bool(info + 33)) {
        info_ffi["columnNumber"] = A.load.Int32(info + 20);
      }
      info_ffi["debugId"] = A.load.Ref(info + 24, undefined);
      info_ffi["stackTrace"] = A.load.Ref(info + 28, undefined);

      const _ret = WEBEXT.crashReportPrivate.reportError(info_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ReportError": (retPtr: Pointer, errPtr: Pointer, info: Pointer): heap.Ref<boolean> => {
      try {
        const info_ffi = {};

        info_ffi["message"] = A.load.Ref(info + 0, undefined);
        info_ffi["url"] = A.load.Ref(info + 4, undefined);
        info_ffi["product"] = A.load.Ref(info + 8, undefined);
        info_ffi["version"] = A.load.Ref(info + 12, undefined);
        if (A.load.Bool(info + 32)) {
          info_ffi["lineNumber"] = A.load.Int32(info + 16);
        }
        if (A.load.Bool(info + 33)) {
          info_ffi["columnNumber"] = A.load.Int32(info + 20);
        }
        info_ffi["debugId"] = A.load.Ref(info + 24, undefined);
        info_ffi["stackTrace"] = A.load.Ref(info + 28, undefined);

        const _ret = WEBEXT.crashReportPrivate.reportError(info_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
