import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/metricsprivateindividualapis", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_RecordEnumerationValue": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivateIndividualApis && "recordEnumerationValue" in WEBEXT?.metricsPrivateIndividualApis) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordEnumerationValue": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivateIndividualApis.recordEnumerationValue);
    },
    "call_RecordEnumerationValue": (
      retPtr: Pointer,
      metricName: heap.Ref<object>,
      value: number,
      enumSize: number
    ): void => {
      const _ret = WEBEXT.metricsPrivateIndividualApis.recordEnumerationValue(
        A.H.get<object>(metricName),
        value,
        enumSize
      );
    },
    "try_RecordEnumerationValue": (
      retPtr: Pointer,
      errPtr: Pointer,
      metricName: heap.Ref<object>,
      value: number,
      enumSize: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivateIndividualApis.recordEnumerationValue(
          A.H.get<object>(metricName),
          value,
          enumSize
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordMediumTime": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivateIndividualApis && "recordMediumTime" in WEBEXT?.metricsPrivateIndividualApis) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordMediumTime": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivateIndividualApis.recordMediumTime);
    },
    "call_RecordMediumTime": (retPtr: Pointer, metricName: heap.Ref<object>, value: number): void => {
      const _ret = WEBEXT.metricsPrivateIndividualApis.recordMediumTime(A.H.get<object>(metricName), value);
    },
    "try_RecordMediumTime": (
      retPtr: Pointer,
      errPtr: Pointer,
      metricName: heap.Ref<object>,
      value: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivateIndividualApis.recordMediumTime(A.H.get<object>(metricName), value);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordUserAction": (): heap.Ref<boolean> => {
      if (WEBEXT?.metricsPrivateIndividualApis && "recordUserAction" in WEBEXT?.metricsPrivateIndividualApis) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordUserAction": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.metricsPrivateIndividualApis.recordUserAction);
    },
    "call_RecordUserAction": (retPtr: Pointer, name: heap.Ref<object>): void => {
      const _ret = WEBEXT.metricsPrivateIndividualApis.recordUserAction(A.H.get<object>(name));
    },
    "try_RecordUserAction": (retPtr: Pointer, errPtr: Pointer, name: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.metricsPrivateIndividualApis.recordUserAction(A.H.get<object>(name));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
