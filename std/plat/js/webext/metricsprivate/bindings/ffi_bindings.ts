import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/metricsprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_HistogramBucket": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 24, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
        A.store.Int64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 24, true);
        A.store.Int64(ptr + 0, x["count"] === undefined ? 0 : (x["count"] as number));
        A.store.Int64(ptr + 8, x["max"] === undefined ? 0 : (x["max"] as number));
        A.store.Int64(ptr + 16, x["min"] === undefined ? 0 : (x["min"] as number));
      }
    },
    "load_HistogramBucket": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["count"] = A.load.Int64(ptr + 0);
      x["max"] = A.load.Int64(ptr + 8);
      x["min"] = A.load.Int64(ptr + 16);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Histogram": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["buckets"]);
        A.store.Float64(ptr + 8, x["sum"] === undefined ? 0 : (x["sum"] as number));
      }
    },
    "load_Histogram": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["buckets"] = A.load.Ref(ptr + 0, undefined);
      x["sum"] = A.load.Float64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_MetricTypeType": (ref: heap.Ref<string>): number => {
      const idx = ["histogram-log", "histogram-linear"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_MetricType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 36, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Int64(ptr + 24, 0);
        A.store.Enum(ptr + 32, -1);
      } else {
        A.store.Bool(ptr + 36, true);
        A.store.Int64(ptr + 0, x["buckets"] === undefined ? 0 : (x["buckets"] as number));
        A.store.Int64(ptr + 8, x["max"] === undefined ? 0 : (x["max"] as number));
        A.store.Ref(ptr + 16, x["metricName"]);
        A.store.Int64(ptr + 24, x["min"] === undefined ? 0 : (x["min"] as number));
        A.store.Enum(ptr + 32, ["histogram-log", "histogram-linear"].indexOf(x["type"] as string));
      }
    },
    "load_MetricType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["buckets"] = A.load.Int64(ptr + 0);
      x["max"] = A.load.Int64(ptr + 8);
      x["metricName"] = A.load.Ref(ptr + 16, undefined);
      x["min"] = A.load.Int64(ptr + 24);
      x["type"] = A.load.Enum(ptr + 32, ["histogram-log", "histogram-linear"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetFieldTrial": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "getFieldTrial" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetFieldTrial": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.getFieldTrial);
    },
    "call_GetFieldTrial": (retPtr: Pointer, name: heap.Ref<object>): void => {
      const _ret = WEBEXT.metricsPrivate.getFieldTrial(A.H.get<object>(name));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetFieldTrial": (retPtr: Pointer, errPtr: Pointer, name: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.getFieldTrial(A.H.get<object>(name));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetHistogram": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "getHistogram" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetHistogram": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.getHistogram);
    },
    "call_GetHistogram": (retPtr: Pointer, name: heap.Ref<object>): void => {
      const _ret = WEBEXT.metricsPrivate.getHistogram(A.H.get<object>(name));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetHistogram": (retPtr: Pointer, errPtr: Pointer, name: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.getHistogram(A.H.get<object>(name));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetIsCrashReportingEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "getIsCrashReportingEnabled" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetIsCrashReportingEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.getIsCrashReportingEnabled);
    },
    "call_GetIsCrashReportingEnabled": (retPtr: Pointer): void => {
      const _ret = WEBEXT.metricsPrivate.getIsCrashReportingEnabled();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetIsCrashReportingEnabled": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.getIsCrashReportingEnabled();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetVariationParams": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "getVariationParams" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetVariationParams": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.getVariationParams);
    },
    "call_GetVariationParams": (retPtr: Pointer, name: heap.Ref<object>): void => {
      const _ret = WEBEXT.metricsPrivate.getVariationParams(A.H.get<object>(name));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetVariationParams": (retPtr: Pointer, errPtr: Pointer, name: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.getVariationParams(A.H.get<object>(name));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordBoolean": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "recordBoolean" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordBoolean": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.recordBoolean);
    },
    "call_RecordBoolean": (retPtr: Pointer, metricName: heap.Ref<object>, value: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.metricsPrivate.recordBoolean(A.H.get<object>(metricName), value === A.H.TRUE);
    },
    "try_RecordBoolean": (
      retPtr: Pointer,
      errPtr: Pointer,
      metricName: heap.Ref<object>,
      value: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.recordBoolean(A.H.get<object>(metricName), value === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordCount": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "recordCount" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordCount": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.recordCount);
    },
    "call_RecordCount": (retPtr: Pointer, metricName: heap.Ref<object>, value: number): void => {
      const _ret = WEBEXT.metricsPrivate.recordCount(A.H.get<object>(metricName), value);
    },
    "try_RecordCount": (
      retPtr: Pointer,
      errPtr: Pointer,
      metricName: heap.Ref<object>,
      value: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.recordCount(A.H.get<object>(metricName), value);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordEnumerationValue": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "recordEnumerationValue" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordEnumerationValue": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.recordEnumerationValue);
    },
    "call_RecordEnumerationValue": (
      retPtr: Pointer,
      metricName: heap.Ref<object>,
      value: number,
      enumSize: number
    ): void => {
      const _ret = WEBEXT.metricsPrivate.recordEnumerationValue(A.H.get<object>(metricName), value, enumSize);
    },
    "try_RecordEnumerationValue": (
      retPtr: Pointer,
      errPtr: Pointer,
      metricName: heap.Ref<object>,
      value: number,
      enumSize: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.recordEnumerationValue(A.H.get<object>(metricName), value, enumSize);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordLongTime": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "recordLongTime" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordLongTime": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.recordLongTime);
    },
    "call_RecordLongTime": (retPtr: Pointer, metricName: heap.Ref<object>, value: number): void => {
      const _ret = WEBEXT.metricsPrivate.recordLongTime(A.H.get<object>(metricName), value);
    },
    "try_RecordLongTime": (
      retPtr: Pointer,
      errPtr: Pointer,
      metricName: heap.Ref<object>,
      value: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.recordLongTime(A.H.get<object>(metricName), value);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordMediumCount": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "recordMediumCount" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordMediumCount": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.recordMediumCount);
    },
    "call_RecordMediumCount": (retPtr: Pointer, metricName: heap.Ref<object>, value: number): void => {
      const _ret = WEBEXT.metricsPrivate.recordMediumCount(A.H.get<object>(metricName), value);
    },
    "try_RecordMediumCount": (
      retPtr: Pointer,
      errPtr: Pointer,
      metricName: heap.Ref<object>,
      value: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.recordMediumCount(A.H.get<object>(metricName), value);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordMediumTime": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "recordMediumTime" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordMediumTime": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.recordMediumTime);
    },
    "call_RecordMediumTime": (retPtr: Pointer, metricName: heap.Ref<object>, value: number): void => {
      const _ret = WEBEXT.metricsPrivate.recordMediumTime(A.H.get<object>(metricName), value);
    },
    "try_RecordMediumTime": (
      retPtr: Pointer,
      errPtr: Pointer,
      metricName: heap.Ref<object>,
      value: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.recordMediumTime(A.H.get<object>(metricName), value);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordPercentage": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "recordPercentage" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordPercentage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.recordPercentage);
    },
    "call_RecordPercentage": (retPtr: Pointer, metricName: heap.Ref<object>, value: number): void => {
      const _ret = WEBEXT.metricsPrivate.recordPercentage(A.H.get<object>(metricName), value);
    },
    "try_RecordPercentage": (
      retPtr: Pointer,
      errPtr: Pointer,
      metricName: heap.Ref<object>,
      value: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.recordPercentage(A.H.get<object>(metricName), value);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordSmallCount": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "recordSmallCount" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordSmallCount": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.recordSmallCount);
    },
    "call_RecordSmallCount": (retPtr: Pointer, metricName: heap.Ref<object>, value: number): void => {
      const _ret = WEBEXT.metricsPrivate.recordSmallCount(A.H.get<object>(metricName), value);
    },
    "try_RecordSmallCount": (
      retPtr: Pointer,
      errPtr: Pointer,
      metricName: heap.Ref<object>,
      value: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.recordSmallCount(A.H.get<object>(metricName), value);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordSparseValue": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "recordSparseValue" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordSparseValue": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.recordSparseValue);
    },
    "call_RecordSparseValue": (retPtr: Pointer, metricName: heap.Ref<object>, value: number): void => {
      const _ret = WEBEXT.metricsPrivate.recordSparseValue(A.H.get<object>(metricName), value);
    },
    "try_RecordSparseValue": (
      retPtr: Pointer,
      errPtr: Pointer,
      metricName: heap.Ref<object>,
      value: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.recordSparseValue(A.H.get<object>(metricName), value);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordSparseValueWithHashMetricName": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "recordSparseValueWithHashMetricName" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordSparseValueWithHashMetricName": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.recordSparseValueWithHashMetricName);
    },
    "call_RecordSparseValueWithHashMetricName": (
      retPtr: Pointer,
      metricName: heap.Ref<object>,
      value: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.metricsPrivate.recordSparseValueWithHashMetricName(
        A.H.get<object>(metricName),
        A.H.get<object>(value)
      );
    },
    "try_RecordSparseValueWithHashMetricName": (
      retPtr: Pointer,
      errPtr: Pointer,
      metricName: heap.Ref<object>,
      value: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.recordSparseValueWithHashMetricName(
          A.H.get<object>(metricName),
          A.H.get<object>(value)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordSparseValueWithPersistentHash": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "recordSparseValueWithPersistentHash" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordSparseValueWithPersistentHash": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.recordSparseValueWithPersistentHash);
    },
    "call_RecordSparseValueWithPersistentHash": (
      retPtr: Pointer,
      metricName: heap.Ref<object>,
      value: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.metricsPrivate.recordSparseValueWithPersistentHash(
        A.H.get<object>(metricName),
        A.H.get<object>(value)
      );
    },
    "try_RecordSparseValueWithPersistentHash": (
      retPtr: Pointer,
      errPtr: Pointer,
      metricName: heap.Ref<object>,
      value: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.recordSparseValueWithPersistentHash(
          A.H.get<object>(metricName),
          A.H.get<object>(value)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordTime": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "recordTime" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordTime": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.recordTime);
    },
    "call_RecordTime": (retPtr: Pointer, metricName: heap.Ref<object>, value: number): void => {
      const _ret = WEBEXT.metricsPrivate.recordTime(A.H.get<object>(metricName), value);
    },
    "try_RecordTime": (
      retPtr: Pointer,
      errPtr: Pointer,
      metricName: heap.Ref<object>,
      value: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.recordTime(A.H.get<object>(metricName), value);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordUserAction": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "recordUserAction" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordUserAction": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.recordUserAction);
    },
    "call_RecordUserAction": (retPtr: Pointer, name: heap.Ref<object>): void => {
      const _ret = WEBEXT.metricsPrivate.recordUserAction(A.H.get<object>(name));
    },
    "try_RecordUserAction": (retPtr: Pointer, errPtr: Pointer, name: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivate.recordUserAction(A.H.get<object>(name));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordValue": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivate && "recordValue" in WEBEXT?.metricsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordValue": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivate.recordValue);
    },
    "call_RecordValue": (retPtr: Pointer, metric: Pointer, value: number): void => {
      const metric_ffi = {};

      metric_ffi["buckets"] = A.load.Int64(metric + 0);
      metric_ffi["max"] = A.load.Int64(metric + 8);
      metric_ffi["metricName"] = A.load.Ref(metric + 16, undefined);
      metric_ffi["min"] = A.load.Int64(metric + 24);
      metric_ffi["type"] = A.load.Enum(metric + 32, ["histogram-log", "histogram-linear"]);

      const _ret = WEBEXT.metricsPrivate.recordValue(metric_ffi, value);
    },
    "try_RecordValue": (retPtr: Pointer, errPtr: Pointer, metric: Pointer, value: number): heap.Ref<boolean> => {
      try {
        const metric_ffi = {};

        metric_ffi["buckets"] = A.load.Int64(metric + 0);
        metric_ffi["max"] = A.load.Int64(metric + 8);
        metric_ffi["metricName"] = A.load.Ref(metric + 16, undefined);
        metric_ffi["min"] = A.load.Int64(metric + 24);
        metric_ffi["type"] = A.load.Enum(metric + 32, ["histogram-log", "histogram-linear"]);

        const _ret = WEBEXT.metricsPrivate.recordValue(metric_ffi, value);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
