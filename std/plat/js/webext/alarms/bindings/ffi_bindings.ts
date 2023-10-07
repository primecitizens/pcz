import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/alarms", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Alarm": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 26, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 24, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 25, false);
        A.store.Float64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 26, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Bool(ptr + 24, "scheduledTime" in x ? true : false);
        A.store.Float64(ptr + 8, x["scheduledTime"] === undefined ? 0 : (x["scheduledTime"] as number));
        A.store.Bool(ptr + 25, "periodInMinutes" in x ? true : false);
        A.store.Float64(ptr + 16, x["periodInMinutes"] === undefined ? 0 : (x["periodInMinutes"] as number));
      }
    },
    "load_Alarm": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 24)) {
        x["scheduledTime"] = A.load.Float64(ptr + 8);
      } else {
        delete x["scheduledTime"];
      }
      if (A.load.Bool(ptr + 25)) {
        x["periodInMinutes"] = A.load.Float64(ptr + 16);
      } else {
        delete x["periodInMinutes"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AlarmCreateInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 27, false);
        A.store.Bool(ptr + 24, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 25, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 26, false);
        A.store.Float64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 27, true);
        A.store.Bool(ptr + 24, "when" in x ? true : false);
        A.store.Float64(ptr + 0, x["when"] === undefined ? 0 : (x["when"] as number));
        A.store.Bool(ptr + 25, "delayInMinutes" in x ? true : false);
        A.store.Float64(ptr + 8, x["delayInMinutes"] === undefined ? 0 : (x["delayInMinutes"] as number));
        A.store.Bool(ptr + 26, "periodInMinutes" in x ? true : false);
        A.store.Float64(ptr + 16, x["periodInMinutes"] === undefined ? 0 : (x["periodInMinutes"] as number));
      }
    },
    "load_AlarmCreateInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["when"] = A.load.Float64(ptr + 0);
      } else {
        delete x["when"];
      }
      if (A.load.Bool(ptr + 25)) {
        x["delayInMinutes"] = A.load.Float64(ptr + 8);
      } else {
        delete x["delayInMinutes"];
      }
      if (A.load.Bool(ptr + 26)) {
        x["periodInMinutes"] = A.load.Float64(ptr + 16);
      } else {
        delete x["periodInMinutes"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Clear": (): heap.Ref<boolean> => {
      if (WEBEXT?.alarms && "clear" in WEBEXT?.alarms) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Clear": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.alarms.clear);
    },
    "call_Clear": (retPtr: Pointer, name: heap.Ref<object>): void => {
      const _ret = WEBEXT.alarms.clear(A.H.get<object>(name));
      A.store.Ref(retPtr, _ret);
    },
    "try_Clear": (retPtr: Pointer, errPtr: Pointer, name: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.alarms.clear(A.H.get<object>(name));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ClearAll": (): heap.Ref<boolean> => {
      if (WEBEXT?.alarms && "clearAll" in WEBEXT?.alarms) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ClearAll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.alarms.clearAll);
    },
    "call_ClearAll": (retPtr: Pointer): void => {
      const _ret = WEBEXT.alarms.clearAll();
      A.store.Ref(retPtr, _ret);
    },
    "try_ClearAll": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.alarms.clearAll();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Create": (): heap.Ref<boolean> => {
      if (WEBEXT?.alarms && "create" in WEBEXT?.alarms) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Create": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.alarms.create);
    },
    "call_Create": (retPtr: Pointer, name: heap.Ref<object>, alarmInfo: Pointer): void => {
      const alarmInfo_ffi = {};

      if (A.load.Bool(alarmInfo + 24)) {
        alarmInfo_ffi["when"] = A.load.Float64(alarmInfo + 0);
      }
      if (A.load.Bool(alarmInfo + 25)) {
        alarmInfo_ffi["delayInMinutes"] = A.load.Float64(alarmInfo + 8);
      }
      if (A.load.Bool(alarmInfo + 26)) {
        alarmInfo_ffi["periodInMinutes"] = A.load.Float64(alarmInfo + 16);
      }

      const _ret = WEBEXT.alarms.create(A.H.get<object>(name), alarmInfo_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Create": (retPtr: Pointer, errPtr: Pointer, name: heap.Ref<object>, alarmInfo: Pointer): heap.Ref<boolean> => {
      try {
        const alarmInfo_ffi = {};

        if (A.load.Bool(alarmInfo + 24)) {
          alarmInfo_ffi["when"] = A.load.Float64(alarmInfo + 0);
        }
        if (A.load.Bool(alarmInfo + 25)) {
          alarmInfo_ffi["delayInMinutes"] = A.load.Float64(alarmInfo + 8);
        }
        if (A.load.Bool(alarmInfo + 26)) {
          alarmInfo_ffi["periodInMinutes"] = A.load.Float64(alarmInfo + 16);
        }

        const _ret = WEBEXT.alarms.create(A.H.get<object>(name), alarmInfo_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Get": (): heap.Ref<boolean> => {
      if (WEBEXT?.alarms && "get" in WEBEXT?.alarms) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Get": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.alarms.get);
    },
    "call_Get": (retPtr: Pointer, name: heap.Ref<object>): void => {
      const _ret = WEBEXT.alarms.get(A.H.get<object>(name));
      A.store.Ref(retPtr, _ret);
    },
    "try_Get": (retPtr: Pointer, errPtr: Pointer, name: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.alarms.get(A.H.get<object>(name));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAll": (): heap.Ref<boolean> => {
      if (WEBEXT?.alarms && "getAll" in WEBEXT?.alarms) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.alarms.getAll);
    },
    "call_GetAll": (retPtr: Pointer): void => {
      const _ret = WEBEXT.alarms.getAll();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAll": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.alarms.getAll();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAlarm": (): heap.Ref<boolean> => {
      if (WEBEXT?.alarms?.onAlarm && "addListener" in WEBEXT?.alarms?.onAlarm) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAlarm": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.alarms.onAlarm.addListener);
    },
    "call_OnAlarm": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.alarms.onAlarm.addListener(A.H.get<object>(callback));
    },
    "try_OnAlarm": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.alarms.onAlarm.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAlarm": (): heap.Ref<boolean> => {
      if (WEBEXT?.alarms?.onAlarm && "removeListener" in WEBEXT?.alarms?.onAlarm) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAlarm": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.alarms.onAlarm.removeListener);
    },
    "call_OffAlarm": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.alarms.onAlarm.removeListener(A.H.get<object>(callback));
    },
    "try_OffAlarm": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.alarms.onAlarm.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAlarm": (): heap.Ref<boolean> => {
      if (WEBEXT?.alarms?.onAlarm && "hasListener" in WEBEXT?.alarms?.onAlarm) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAlarm": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.alarms.onAlarm.hasListener);
    },
    "call_HasOnAlarm": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.alarms.onAlarm.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAlarm": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.alarms.onAlarm.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
