import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/processes", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Cache": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Bool(ptr + 16, "size" in x ? true : false);
        A.store.Float64(ptr + 0, x["size"] === undefined ? 0 : (x["size"] as number));
        A.store.Bool(ptr + 17, "liveSize" in x ? true : false);
        A.store.Float64(ptr + 8, x["liveSize"] === undefined ? 0 : (x["liveSize"] as number));
      }
    },
    "load_Cache": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["size"] = A.load.Float64(ptr + 0);
      } else {
        delete x["size"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["liveSize"] = A.load.Float64(ptr + 8);
      } else {
        delete x["liveSize"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ProcessType": (ref: heap.Ref<string>): number => {
      const idx = [
        "browser",
        "renderer",
        "extension",
        "notification",
        "plugin",
        "worker",
        "nacl",
        "service_worker",
        "utility",
        "gpu",
        "other",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_TaskInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Ref(ptr + 0, x["title"]);
        A.store.Bool(ptr + 8, "tabId" in x ? true : false);
        A.store.Int32(ptr + 4, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_TaskInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["title"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 8)) {
        x["tabId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["tabId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Process": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 148, false);
        A.store.Bool(ptr + 139, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 140, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Enum(ptr + 8, -1);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 141, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Ref(ptr + 20, undefined);
        A.store.Bool(ptr + 142, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Bool(ptr + 143, false);
        A.store.Float64(ptr + 32, 0);
        A.store.Bool(ptr + 144, false);
        A.store.Float64(ptr + 40, 0);
        A.store.Bool(ptr + 145, false);
        A.store.Float64(ptr + 48, 0);
        A.store.Bool(ptr + 146, false);
        A.store.Float64(ptr + 56, 0);
        A.store.Bool(ptr + 147, false);
        A.store.Float64(ptr + 64, 0);

        A.store.Bool(ptr + 72 + 18, false);
        A.store.Bool(ptr + 72 + 16, false);
        A.store.Float64(ptr + 72 + 0, 0);
        A.store.Bool(ptr + 72 + 17, false);
        A.store.Float64(ptr + 72 + 8, 0);

        A.store.Bool(ptr + 96 + 18, false);
        A.store.Bool(ptr + 96 + 16, false);
        A.store.Float64(ptr + 96 + 0, 0);
        A.store.Bool(ptr + 96 + 17, false);
        A.store.Float64(ptr + 96 + 8, 0);

        A.store.Bool(ptr + 120 + 18, false);
        A.store.Bool(ptr + 120 + 16, false);
        A.store.Float64(ptr + 120 + 0, 0);
        A.store.Bool(ptr + 120 + 17, false);
        A.store.Float64(ptr + 120 + 8, 0);
      } else {
        A.store.Bool(ptr + 148, true);
        A.store.Bool(ptr + 139, "id" in x ? true : false);
        A.store.Int32(ptr + 0, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Bool(ptr + 140, "osProcessId" in x ? true : false);
        A.store.Int32(ptr + 4, x["osProcessId"] === undefined ? 0 : (x["osProcessId"] as number));
        A.store.Enum(
          ptr + 8,
          [
            "browser",
            "renderer",
            "extension",
            "notification",
            "plugin",
            "worker",
            "nacl",
            "service_worker",
            "utility",
            "gpu",
            "other",
          ].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 12, x["profile"]);
        A.store.Bool(ptr + 141, "naclDebugPort" in x ? true : false);
        A.store.Int32(ptr + 16, x["naclDebugPort"] === undefined ? 0 : (x["naclDebugPort"] as number));
        A.store.Ref(ptr + 20, x["tasks"]);
        A.store.Bool(ptr + 142, "cpu" in x ? true : false);
        A.store.Float64(ptr + 24, x["cpu"] === undefined ? 0 : (x["cpu"] as number));
        A.store.Bool(ptr + 143, "network" in x ? true : false);
        A.store.Float64(ptr + 32, x["network"] === undefined ? 0 : (x["network"] as number));
        A.store.Bool(ptr + 144, "privateMemory" in x ? true : false);
        A.store.Float64(ptr + 40, x["privateMemory"] === undefined ? 0 : (x["privateMemory"] as number));
        A.store.Bool(ptr + 145, "jsMemoryAllocated" in x ? true : false);
        A.store.Float64(ptr + 48, x["jsMemoryAllocated"] === undefined ? 0 : (x["jsMemoryAllocated"] as number));
        A.store.Bool(ptr + 146, "jsMemoryUsed" in x ? true : false);
        A.store.Float64(ptr + 56, x["jsMemoryUsed"] === undefined ? 0 : (x["jsMemoryUsed"] as number));
        A.store.Bool(ptr + 147, "sqliteMemory" in x ? true : false);
        A.store.Float64(ptr + 64, x["sqliteMemory"] === undefined ? 0 : (x["sqliteMemory"] as number));

        if (typeof x["imageCache"] === "undefined") {
          A.store.Bool(ptr + 72 + 18, false);
          A.store.Bool(ptr + 72 + 16, false);
          A.store.Float64(ptr + 72 + 0, 0);
          A.store.Bool(ptr + 72 + 17, false);
          A.store.Float64(ptr + 72 + 8, 0);
        } else {
          A.store.Bool(ptr + 72 + 18, true);
          A.store.Bool(ptr + 72 + 16, "size" in x["imageCache"] ? true : false);
          A.store.Float64(
            ptr + 72 + 0,
            x["imageCache"]["size"] === undefined ? 0 : (x["imageCache"]["size"] as number)
          );
          A.store.Bool(ptr + 72 + 17, "liveSize" in x["imageCache"] ? true : false);
          A.store.Float64(
            ptr + 72 + 8,
            x["imageCache"]["liveSize"] === undefined ? 0 : (x["imageCache"]["liveSize"] as number)
          );
        }

        if (typeof x["scriptCache"] === "undefined") {
          A.store.Bool(ptr + 96 + 18, false);
          A.store.Bool(ptr + 96 + 16, false);
          A.store.Float64(ptr + 96 + 0, 0);
          A.store.Bool(ptr + 96 + 17, false);
          A.store.Float64(ptr + 96 + 8, 0);
        } else {
          A.store.Bool(ptr + 96 + 18, true);
          A.store.Bool(ptr + 96 + 16, "size" in x["scriptCache"] ? true : false);
          A.store.Float64(
            ptr + 96 + 0,
            x["scriptCache"]["size"] === undefined ? 0 : (x["scriptCache"]["size"] as number)
          );
          A.store.Bool(ptr + 96 + 17, "liveSize" in x["scriptCache"] ? true : false);
          A.store.Float64(
            ptr + 96 + 8,
            x["scriptCache"]["liveSize"] === undefined ? 0 : (x["scriptCache"]["liveSize"] as number)
          );
        }

        if (typeof x["cssCache"] === "undefined") {
          A.store.Bool(ptr + 120 + 18, false);
          A.store.Bool(ptr + 120 + 16, false);
          A.store.Float64(ptr + 120 + 0, 0);
          A.store.Bool(ptr + 120 + 17, false);
          A.store.Float64(ptr + 120 + 8, 0);
        } else {
          A.store.Bool(ptr + 120 + 18, true);
          A.store.Bool(ptr + 120 + 16, "size" in x["cssCache"] ? true : false);
          A.store.Float64(ptr + 120 + 0, x["cssCache"]["size"] === undefined ? 0 : (x["cssCache"]["size"] as number));
          A.store.Bool(ptr + 120 + 17, "liveSize" in x["cssCache"] ? true : false);
          A.store.Float64(
            ptr + 120 + 8,
            x["cssCache"]["liveSize"] === undefined ? 0 : (x["cssCache"]["liveSize"] as number)
          );
        }
      }
    },
    "load_Process": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 139)) {
        x["id"] = A.load.Int32(ptr + 0);
      } else {
        delete x["id"];
      }
      if (A.load.Bool(ptr + 140)) {
        x["osProcessId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["osProcessId"];
      }
      x["type"] = A.load.Enum(ptr + 8, [
        "browser",
        "renderer",
        "extension",
        "notification",
        "plugin",
        "worker",
        "nacl",
        "service_worker",
        "utility",
        "gpu",
        "other",
      ]);
      x["profile"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 141)) {
        x["naclDebugPort"] = A.load.Int32(ptr + 16);
      } else {
        delete x["naclDebugPort"];
      }
      x["tasks"] = A.load.Ref(ptr + 20, undefined);
      if (A.load.Bool(ptr + 142)) {
        x["cpu"] = A.load.Float64(ptr + 24);
      } else {
        delete x["cpu"];
      }
      if (A.load.Bool(ptr + 143)) {
        x["network"] = A.load.Float64(ptr + 32);
      } else {
        delete x["network"];
      }
      if (A.load.Bool(ptr + 144)) {
        x["privateMemory"] = A.load.Float64(ptr + 40);
      } else {
        delete x["privateMemory"];
      }
      if (A.load.Bool(ptr + 145)) {
        x["jsMemoryAllocated"] = A.load.Float64(ptr + 48);
      } else {
        delete x["jsMemoryAllocated"];
      }
      if (A.load.Bool(ptr + 146)) {
        x["jsMemoryUsed"] = A.load.Float64(ptr + 56);
      } else {
        delete x["jsMemoryUsed"];
      }
      if (A.load.Bool(ptr + 147)) {
        x["sqliteMemory"] = A.load.Float64(ptr + 64);
      } else {
        delete x["sqliteMemory"];
      }
      if (A.load.Bool(ptr + 72 + 18)) {
        x["imageCache"] = {};
        if (A.load.Bool(ptr + 72 + 16)) {
          x["imageCache"]["size"] = A.load.Float64(ptr + 72 + 0);
        } else {
          delete x["imageCache"]["size"];
        }
        if (A.load.Bool(ptr + 72 + 17)) {
          x["imageCache"]["liveSize"] = A.load.Float64(ptr + 72 + 8);
        } else {
          delete x["imageCache"]["liveSize"];
        }
      } else {
        delete x["imageCache"];
      }
      if (A.load.Bool(ptr + 96 + 18)) {
        x["scriptCache"] = {};
        if (A.load.Bool(ptr + 96 + 16)) {
          x["scriptCache"]["size"] = A.load.Float64(ptr + 96 + 0);
        } else {
          delete x["scriptCache"]["size"];
        }
        if (A.load.Bool(ptr + 96 + 17)) {
          x["scriptCache"]["liveSize"] = A.load.Float64(ptr + 96 + 8);
        } else {
          delete x["scriptCache"]["liveSize"];
        }
      } else {
        delete x["scriptCache"];
      }
      if (A.load.Bool(ptr + 120 + 18)) {
        x["cssCache"] = {};
        if (A.load.Bool(ptr + 120 + 16)) {
          x["cssCache"]["size"] = A.load.Float64(ptr + 120 + 0);
        } else {
          delete x["cssCache"]["size"];
        }
        if (A.load.Bool(ptr + 120 + 17)) {
          x["cssCache"]["liveSize"] = A.load.Float64(ptr + 120 + 8);
        } else {
          delete x["cssCache"]["liveSize"];
        }
      } else {
        delete x["cssCache"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetProcessIdForTab": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes && "getProcessIdForTab" in WEBEXT?.processes) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetProcessIdForTab": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.getProcessIdForTab);
    },
    "call_GetProcessIdForTab": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.processes.getProcessIdForTab(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetProcessIdForTab": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.getProcessIdForTab(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetProcessInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes && "getProcessInfo" in WEBEXT?.processes) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetProcessInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.getProcessInfo);
    },
    "call_GetProcessInfo": (retPtr: Pointer, processIds: heap.Ref<any>, includeMemory: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.processes.getProcessInfo(A.H.get<any>(processIds), includeMemory === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetProcessInfo": (
      retPtr: Pointer,
      errPtr: Pointer,
      processIds: heap.Ref<any>,
      includeMemory: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.getProcessInfo(A.H.get<any>(processIds), includeMemory === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes?.onCreated && "addListener" in WEBEXT?.processes?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.onCreated.addListener);
    },
    "call_OnCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.processes.onCreated.addListener(A.H.get<object>(callback));
    },
    "try_OnCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.onCreated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes?.onCreated && "removeListener" in WEBEXT?.processes?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.onCreated.removeListener);
    },
    "call_OffCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.processes.onCreated.removeListener(A.H.get<object>(callback));
    },
    "try_OffCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.onCreated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes?.onCreated && "hasListener" in WEBEXT?.processes?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.onCreated.hasListener);
    },
    "call_HasOnCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.processes.onCreated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.onCreated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnExited": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes?.onExited && "addListener" in WEBEXT?.processes?.onExited) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnExited": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.onExited.addListener);
    },
    "call_OnExited": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.processes.onExited.addListener(A.H.get<object>(callback));
    },
    "try_OnExited": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.onExited.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffExited": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes?.onExited && "removeListener" in WEBEXT?.processes?.onExited) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffExited": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.onExited.removeListener);
    },
    "call_OffExited": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.processes.onExited.removeListener(A.H.get<object>(callback));
    },
    "try_OffExited": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.onExited.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnExited": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes?.onExited && "hasListener" in WEBEXT?.processes?.onExited) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnExited": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.onExited.hasListener);
    },
    "call_HasOnExited": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.processes.onExited.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnExited": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.onExited.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnUnresponsive": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes?.onUnresponsive && "addListener" in WEBEXT?.processes?.onUnresponsive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnUnresponsive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.onUnresponsive.addListener);
    },
    "call_OnUnresponsive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.processes.onUnresponsive.addListener(A.H.get<object>(callback));
    },
    "try_OnUnresponsive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.onUnresponsive.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffUnresponsive": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes?.onUnresponsive && "removeListener" in WEBEXT?.processes?.onUnresponsive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffUnresponsive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.onUnresponsive.removeListener);
    },
    "call_OffUnresponsive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.processes.onUnresponsive.removeListener(A.H.get<object>(callback));
    },
    "try_OffUnresponsive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.onUnresponsive.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnUnresponsive": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes?.onUnresponsive && "hasListener" in WEBEXT?.processes?.onUnresponsive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnUnresponsive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.onUnresponsive.hasListener);
    },
    "call_HasOnUnresponsive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.processes.onUnresponsive.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnUnresponsive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.onUnresponsive.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnUpdated": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes?.onUpdated && "addListener" in WEBEXT?.processes?.onUpdated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.onUpdated.addListener);
    },
    "call_OnUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.processes.onUpdated.addListener(A.H.get<object>(callback));
    },
    "try_OnUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.onUpdated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffUpdated": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes?.onUpdated && "removeListener" in WEBEXT?.processes?.onUpdated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.onUpdated.removeListener);
    },
    "call_OffUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.processes.onUpdated.removeListener(A.H.get<object>(callback));
    },
    "try_OffUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.onUpdated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnUpdated": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes?.onUpdated && "hasListener" in WEBEXT?.processes?.onUpdated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.onUpdated.hasListener);
    },
    "call_HasOnUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.processes.onUpdated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.onUpdated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnUpdatedWithMemory": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes?.onUpdatedWithMemory && "addListener" in WEBEXT?.processes?.onUpdatedWithMemory) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnUpdatedWithMemory": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.onUpdatedWithMemory.addListener);
    },
    "call_OnUpdatedWithMemory": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.processes.onUpdatedWithMemory.addListener(A.H.get<object>(callback));
    },
    "try_OnUpdatedWithMemory": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.onUpdatedWithMemory.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffUpdatedWithMemory": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes?.onUpdatedWithMemory && "removeListener" in WEBEXT?.processes?.onUpdatedWithMemory) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffUpdatedWithMemory": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.onUpdatedWithMemory.removeListener);
    },
    "call_OffUpdatedWithMemory": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.processes.onUpdatedWithMemory.removeListener(A.H.get<object>(callback));
    },
    "try_OffUpdatedWithMemory": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.onUpdatedWithMemory.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnUpdatedWithMemory": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes?.onUpdatedWithMemory && "hasListener" in WEBEXT?.processes?.onUpdatedWithMemory) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnUpdatedWithMemory": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.onUpdatedWithMemory.hasListener);
    },
    "call_HasOnUpdatedWithMemory": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.processes.onUpdatedWithMemory.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnUpdatedWithMemory": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.onUpdatedWithMemory.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Terminate": (): heap.Ref<boolean> => {
      if (WEBEXT?.processes && "terminate" in WEBEXT?.processes) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Terminate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.processes.terminate);
    },
    "call_Terminate": (retPtr: Pointer, processId: number): void => {
      const _ret = WEBEXT.processes.terminate(processId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Terminate": (retPtr: Pointer, errPtr: Pointer, processId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.processes.terminate(processId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
