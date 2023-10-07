import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/printingmetrics", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_ColorMode": (ref: heap.Ref<string>): number => {
      const idx = ["BLACK_AND_WHITE", "COLOR"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_DuplexMode": (ref: heap.Ref<string>): number => {
      const idx = ["ONE_SIDED", "TWO_SIDED_LONG_EDGE", "TWO_SIDED_SHORT_EDGE"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PrintJobSource": (ref: heap.Ref<string>): number => {
      const idx = ["PRINT_PREVIEW", "ANDROID_APP", "EXTENSION"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PrintJobStatus": (ref: heap.Ref<string>): number => {
      const idx = ["FAILED", "CANCELED", "PRINTED"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PrinterSource": (ref: heap.Ref<string>): number => {
      const idx = ["USER", "POLICY"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Printer": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["uri"]);
        A.store.Enum(ptr + 8, ["USER", "POLICY"].indexOf(x["source"] as string));
      }
    },
    "load_Printer": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["uri"] = A.load.Ref(ptr + 4, undefined);
      x["source"] = A.load.Enum(ptr + 8, ["USER", "POLICY"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MediaSize": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Bool(ptr + 12, "width" in x ? true : false);
        A.store.Int32(ptr + 0, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Bool(ptr + 13, "height" in x ? true : false);
        A.store.Int32(ptr + 4, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Ref(ptr + 8, x["vendorId"]);
      }
    },
    "load_MediaSize": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["width"] = A.load.Int32(ptr + 0);
      } else {
        delete x["width"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["height"] = A.load.Int32(ptr + 4);
      } else {
        delete x["height"];
      }
      x["vendorId"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PrintSettings": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 29, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Enum(ptr + 4, -1);

        A.store.Bool(ptr + 8 + 14, false);
        A.store.Bool(ptr + 8 + 12, false);
        A.store.Int32(ptr + 8 + 0, 0);
        A.store.Bool(ptr + 8 + 13, false);
        A.store.Int32(ptr + 8 + 4, 0);
        A.store.Ref(ptr + 8 + 8, undefined);
        A.store.Bool(ptr + 28, false);
        A.store.Int32(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 29, true);
        A.store.Enum(ptr + 0, ["BLACK_AND_WHITE", "COLOR"].indexOf(x["color"] as string));
        A.store.Enum(
          ptr + 4,
          ["ONE_SIDED", "TWO_SIDED_LONG_EDGE", "TWO_SIDED_SHORT_EDGE"].indexOf(x["duplex"] as string)
        );

        if (typeof x["mediaSize"] === "undefined") {
          A.store.Bool(ptr + 8 + 14, false);
          A.store.Bool(ptr + 8 + 12, false);
          A.store.Int32(ptr + 8 + 0, 0);
          A.store.Bool(ptr + 8 + 13, false);
          A.store.Int32(ptr + 8 + 4, 0);
          A.store.Ref(ptr + 8 + 8, undefined);
        } else {
          A.store.Bool(ptr + 8 + 14, true);
          A.store.Bool(ptr + 8 + 12, "width" in x["mediaSize"] ? true : false);
          A.store.Int32(ptr + 8 + 0, x["mediaSize"]["width"] === undefined ? 0 : (x["mediaSize"]["width"] as number));
          A.store.Bool(ptr + 8 + 13, "height" in x["mediaSize"] ? true : false);
          A.store.Int32(ptr + 8 + 4, x["mediaSize"]["height"] === undefined ? 0 : (x["mediaSize"]["height"] as number));
          A.store.Ref(ptr + 8 + 8, x["mediaSize"]["vendorId"]);
        }
        A.store.Bool(ptr + 28, "copies" in x ? true : false);
        A.store.Int32(ptr + 24, x["copies"] === undefined ? 0 : (x["copies"] as number));
      }
    },
    "load_PrintSettings": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["color"] = A.load.Enum(ptr + 0, ["BLACK_AND_WHITE", "COLOR"]);
      x["duplex"] = A.load.Enum(ptr + 4, ["ONE_SIDED", "TWO_SIDED_LONG_EDGE", "TWO_SIDED_SHORT_EDGE"]);
      if (A.load.Bool(ptr + 8 + 14)) {
        x["mediaSize"] = {};
        if (A.load.Bool(ptr + 8 + 12)) {
          x["mediaSize"]["width"] = A.load.Int32(ptr + 8 + 0);
        } else {
          delete x["mediaSize"]["width"];
        }
        if (A.load.Bool(ptr + 8 + 13)) {
          x["mediaSize"]["height"] = A.load.Int32(ptr + 8 + 4);
        } else {
          delete x["mediaSize"]["height"];
        }
        x["mediaSize"]["vendorId"] = A.load.Ref(ptr + 8 + 8, undefined);
      } else {
        delete x["mediaSize"];
      }
      if (A.load.Bool(ptr + 28)) {
        x["copies"] = A.load.Int32(ptr + 24);
      } else {
        delete x["copies"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PrintJobInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 99, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Ref(ptr + 12, undefined);
        A.store.Enum(ptr + 16, -1);
        A.store.Bool(ptr + 96, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Bool(ptr + 97, false);
        A.store.Float64(ptr + 32, 0);

        A.store.Bool(ptr + 40 + 12, false);
        A.store.Ref(ptr + 40 + 0, undefined);
        A.store.Ref(ptr + 40 + 4, undefined);
        A.store.Enum(ptr + 40 + 8, -1);

        A.store.Bool(ptr + 56 + 29, false);
        A.store.Enum(ptr + 56 + 0, -1);
        A.store.Enum(ptr + 56 + 4, -1);

        A.store.Bool(ptr + 56 + 8 + 14, false);
        A.store.Bool(ptr + 56 + 8 + 12, false);
        A.store.Int32(ptr + 56 + 8 + 0, 0);
        A.store.Bool(ptr + 56 + 8 + 13, false);
        A.store.Int32(ptr + 56 + 8 + 4, 0);
        A.store.Ref(ptr + 56 + 8 + 8, undefined);
        A.store.Bool(ptr + 56 + 28, false);
        A.store.Int32(ptr + 56 + 24, 0);
        A.store.Bool(ptr + 98, false);
        A.store.Int32(ptr + 88, 0);
        A.store.Enum(ptr + 92, -1);
      } else {
        A.store.Bool(ptr + 99, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["title"]);
        A.store.Enum(ptr + 8, ["PRINT_PREVIEW", "ANDROID_APP", "EXTENSION"].indexOf(x["source"] as string));
        A.store.Ref(ptr + 12, x["sourceId"]);
        A.store.Enum(ptr + 16, ["FAILED", "CANCELED", "PRINTED"].indexOf(x["status"] as string));
        A.store.Bool(ptr + 96, "creationTime" in x ? true : false);
        A.store.Float64(ptr + 24, x["creationTime"] === undefined ? 0 : (x["creationTime"] as number));
        A.store.Bool(ptr + 97, "completionTime" in x ? true : false);
        A.store.Float64(ptr + 32, x["completionTime"] === undefined ? 0 : (x["completionTime"] as number));

        if (typeof x["printer"] === "undefined") {
          A.store.Bool(ptr + 40 + 12, false);
          A.store.Ref(ptr + 40 + 0, undefined);
          A.store.Ref(ptr + 40 + 4, undefined);
          A.store.Enum(ptr + 40 + 8, -1);
        } else {
          A.store.Bool(ptr + 40 + 12, true);
          A.store.Ref(ptr + 40 + 0, x["printer"]["name"]);
          A.store.Ref(ptr + 40 + 4, x["printer"]["uri"]);
          A.store.Enum(ptr + 40 + 8, ["USER", "POLICY"].indexOf(x["printer"]["source"] as string));
        }

        if (typeof x["settings"] === "undefined") {
          A.store.Bool(ptr + 56 + 29, false);
          A.store.Enum(ptr + 56 + 0, -1);
          A.store.Enum(ptr + 56 + 4, -1);

          A.store.Bool(ptr + 56 + 8 + 14, false);
          A.store.Bool(ptr + 56 + 8 + 12, false);
          A.store.Int32(ptr + 56 + 8 + 0, 0);
          A.store.Bool(ptr + 56 + 8 + 13, false);
          A.store.Int32(ptr + 56 + 8 + 4, 0);
          A.store.Ref(ptr + 56 + 8 + 8, undefined);
          A.store.Bool(ptr + 56 + 28, false);
          A.store.Int32(ptr + 56 + 24, 0);
        } else {
          A.store.Bool(ptr + 56 + 29, true);
          A.store.Enum(ptr + 56 + 0, ["BLACK_AND_WHITE", "COLOR"].indexOf(x["settings"]["color"] as string));
          A.store.Enum(
            ptr + 56 + 4,
            ["ONE_SIDED", "TWO_SIDED_LONG_EDGE", "TWO_SIDED_SHORT_EDGE"].indexOf(x["settings"]["duplex"] as string)
          );

          if (typeof x["settings"]["mediaSize"] === "undefined") {
            A.store.Bool(ptr + 56 + 8 + 14, false);
            A.store.Bool(ptr + 56 + 8 + 12, false);
            A.store.Int32(ptr + 56 + 8 + 0, 0);
            A.store.Bool(ptr + 56 + 8 + 13, false);
            A.store.Int32(ptr + 56 + 8 + 4, 0);
            A.store.Ref(ptr + 56 + 8 + 8, undefined);
          } else {
            A.store.Bool(ptr + 56 + 8 + 14, true);
            A.store.Bool(ptr + 56 + 8 + 12, "width" in x["settings"]["mediaSize"] ? true : false);
            A.store.Int32(
              ptr + 56 + 8 + 0,
              x["settings"]["mediaSize"]["width"] === undefined ? 0 : (x["settings"]["mediaSize"]["width"] as number)
            );
            A.store.Bool(ptr + 56 + 8 + 13, "height" in x["settings"]["mediaSize"] ? true : false);
            A.store.Int32(
              ptr + 56 + 8 + 4,
              x["settings"]["mediaSize"]["height"] === undefined ? 0 : (x["settings"]["mediaSize"]["height"] as number)
            );
            A.store.Ref(ptr + 56 + 8 + 8, x["settings"]["mediaSize"]["vendorId"]);
          }
          A.store.Bool(ptr + 56 + 28, "copies" in x["settings"] ? true : false);
          A.store.Int32(ptr + 56 + 24, x["settings"]["copies"] === undefined ? 0 : (x["settings"]["copies"] as number));
        }
        A.store.Bool(ptr + 98, "numberOfPages" in x ? true : false);
        A.store.Int32(ptr + 88, x["numberOfPages"] === undefined ? 0 : (x["numberOfPages"] as number));
        A.store.Enum(
          ptr + 92,
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
          ].indexOf(x["printer_status"] as string)
        );
      }
    },
    "load_PrintJobInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["title"] = A.load.Ref(ptr + 4, undefined);
      x["source"] = A.load.Enum(ptr + 8, ["PRINT_PREVIEW", "ANDROID_APP", "EXTENSION"]);
      x["sourceId"] = A.load.Ref(ptr + 12, undefined);
      x["status"] = A.load.Enum(ptr + 16, ["FAILED", "CANCELED", "PRINTED"]);
      if (A.load.Bool(ptr + 96)) {
        x["creationTime"] = A.load.Float64(ptr + 24);
      } else {
        delete x["creationTime"];
      }
      if (A.load.Bool(ptr + 97)) {
        x["completionTime"] = A.load.Float64(ptr + 32);
      } else {
        delete x["completionTime"];
      }
      if (A.load.Bool(ptr + 40 + 12)) {
        x["printer"] = {};
        x["printer"]["name"] = A.load.Ref(ptr + 40 + 0, undefined);
        x["printer"]["uri"] = A.load.Ref(ptr + 40 + 4, undefined);
        x["printer"]["source"] = A.load.Enum(ptr + 40 + 8, ["USER", "POLICY"]);
      } else {
        delete x["printer"];
      }
      if (A.load.Bool(ptr + 56 + 29)) {
        x["settings"] = {};
        x["settings"]["color"] = A.load.Enum(ptr + 56 + 0, ["BLACK_AND_WHITE", "COLOR"]);
        x["settings"]["duplex"] = A.load.Enum(ptr + 56 + 4, [
          "ONE_SIDED",
          "TWO_SIDED_LONG_EDGE",
          "TWO_SIDED_SHORT_EDGE",
        ]);
        if (A.load.Bool(ptr + 56 + 8 + 14)) {
          x["settings"]["mediaSize"] = {};
          if (A.load.Bool(ptr + 56 + 8 + 12)) {
            x["settings"]["mediaSize"]["width"] = A.load.Int32(ptr + 56 + 8 + 0);
          } else {
            delete x["settings"]["mediaSize"]["width"];
          }
          if (A.load.Bool(ptr + 56 + 8 + 13)) {
            x["settings"]["mediaSize"]["height"] = A.load.Int32(ptr + 56 + 8 + 4);
          } else {
            delete x["settings"]["mediaSize"]["height"];
          }
          x["settings"]["mediaSize"]["vendorId"] = A.load.Ref(ptr + 56 + 8 + 8, undefined);
        } else {
          delete x["settings"]["mediaSize"];
        }
        if (A.load.Bool(ptr + 56 + 28)) {
          x["settings"]["copies"] = A.load.Int32(ptr + 56 + 24);
        } else {
          delete x["settings"]["copies"];
        }
      } else {
        delete x["settings"];
      }
      if (A.load.Bool(ptr + 98)) {
        x["numberOfPages"] = A.load.Int32(ptr + 88);
      } else {
        delete x["numberOfPages"];
      }
      x["printer_status"] = A.load.Enum(ptr + 92, [
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
    "has_GetPrintJobs": (): heap.Ref<boolean> => {
      if (WEBEXT?.printingMetrics && "getPrintJobs" in WEBEXT?.printingMetrics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPrintJobs": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printingMetrics.getPrintJobs);
    },
    "call_GetPrintJobs": (retPtr: Pointer): void => {
      const _ret = WEBEXT.printingMetrics.getPrintJobs();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPrintJobs": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printingMetrics.getPrintJobs();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPrintJobFinished": (): heap.Ref<boolean> => {
      if (WEBEXT?.printingMetrics?.onPrintJobFinished && "addListener" in WEBEXT?.printingMetrics?.onPrintJobFinished) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPrintJobFinished": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printingMetrics.onPrintJobFinished.addListener);
    },
    "call_OnPrintJobFinished": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printingMetrics.onPrintJobFinished.addListener(A.H.get<object>(callback));
    },
    "try_OnPrintJobFinished": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printingMetrics.onPrintJobFinished.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPrintJobFinished": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.printingMetrics?.onPrintJobFinished &&
        "removeListener" in WEBEXT?.printingMetrics?.onPrintJobFinished
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPrintJobFinished": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printingMetrics.onPrintJobFinished.removeListener);
    },
    "call_OffPrintJobFinished": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printingMetrics.onPrintJobFinished.removeListener(A.H.get<object>(callback));
    },
    "try_OffPrintJobFinished": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printingMetrics.onPrintJobFinished.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPrintJobFinished": (): heap.Ref<boolean> => {
      if (WEBEXT?.printingMetrics?.onPrintJobFinished && "hasListener" in WEBEXT?.printingMetrics?.onPrintJobFinished) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPrintJobFinished": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.printingMetrics.onPrintJobFinished.hasListener);
    },
    "call_HasOnPrintJobFinished": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.printingMetrics.onPrintJobFinished.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPrintJobFinished": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.printingMetrics.onPrintJobFinished.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
