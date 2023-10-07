import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/debugger", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Debuggee": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 20, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Ref(ptr + 0, x["extensionId"]);
        A.store.Bool(ptr + 20, "tabId" in x ? true : false);
        A.store.Int64(ptr + 8, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Ref(ptr + 16, x["targetId"]);
      }
    },
    "load_Debuggee": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["extensionId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 20)) {
        x["tabId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["tabId"];
      }
      x["targetId"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DetachReason": (ref: heap.Ref<string>): number => {
      const idx = ["target_closed", "canceled_by_user"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_TargetInfoType": (ref: heap.Ref<string>): number => {
      const idx = ["page", "background_page", "worker", "other"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_TargetInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 37, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 36, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Ref(ptr + 24, undefined);
        A.store.Enum(ptr + 28, -1);
        A.store.Ref(ptr + 32, undefined);
      } else {
        A.store.Bool(ptr + 37, true);
        A.store.Bool(ptr + 0, x["attached"] ? true : false);
        A.store.Ref(ptr + 4, x["extensionId"]);
        A.store.Ref(ptr + 8, x["faviconUrl"]);
        A.store.Ref(ptr + 12, x["id"]);
        A.store.Bool(ptr + 36, "tabId" in x ? true : false);
        A.store.Int64(ptr + 16, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Ref(ptr + 24, x["title"]);
        A.store.Enum(ptr + 28, ["page", "background_page", "worker", "other"].indexOf(x["type"] as string));
        A.store.Ref(ptr + 32, x["url"]);
      }
    },
    "load_TargetInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["attached"] = A.load.Bool(ptr + 0);
      x["extensionId"] = A.load.Ref(ptr + 4, undefined);
      x["faviconUrl"] = A.load.Ref(ptr + 8, undefined);
      x["id"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 36)) {
        x["tabId"] = A.load.Int64(ptr + 16);
      } else {
        delete x["tabId"];
      }
      x["title"] = A.load.Ref(ptr + 24, undefined);
      x["type"] = A.load.Enum(ptr + 28, ["page", "background_page", "worker", "other"]);
      x["url"] = A.load.Ref(ptr + 32, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Attach": (): heap.Ref<boolean> => {
      if (WEBEXT?.debugger && "attach" in WEBEXT?.debugger) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Attach": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.debugger.attach);
    },
    "call_Attach": (retPtr: Pointer, target: Pointer, requiredVersion: heap.Ref<object>): void => {
      const target_ffi = {};

      target_ffi["extensionId"] = A.load.Ref(target + 0, undefined);
      if (A.load.Bool(target + 20)) {
        target_ffi["tabId"] = A.load.Int64(target + 8);
      }
      target_ffi["targetId"] = A.load.Ref(target + 16, undefined);

      const _ret = WEBEXT.debugger.attach(target_ffi, A.H.get<object>(requiredVersion));
      A.store.Ref(retPtr, _ret);
    },
    "try_Attach": (
      retPtr: Pointer,
      errPtr: Pointer,
      target: Pointer,
      requiredVersion: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const target_ffi = {};

        target_ffi["extensionId"] = A.load.Ref(target + 0, undefined);
        if (A.load.Bool(target + 20)) {
          target_ffi["tabId"] = A.load.Int64(target + 8);
        }
        target_ffi["targetId"] = A.load.Ref(target + 16, undefined);

        const _ret = WEBEXT.debugger.attach(target_ffi, A.H.get<object>(requiredVersion));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Detach": (): heap.Ref<boolean> => {
      if (WEBEXT?.debugger && "detach" in WEBEXT?.debugger) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Detach": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.debugger.detach);
    },
    "call_Detach": (retPtr: Pointer, target: Pointer): void => {
      const target_ffi = {};

      target_ffi["extensionId"] = A.load.Ref(target + 0, undefined);
      if (A.load.Bool(target + 20)) {
        target_ffi["tabId"] = A.load.Int64(target + 8);
      }
      target_ffi["targetId"] = A.load.Ref(target + 16, undefined);

      const _ret = WEBEXT.debugger.detach(target_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Detach": (retPtr: Pointer, errPtr: Pointer, target: Pointer): heap.Ref<boolean> => {
      try {
        const target_ffi = {};

        target_ffi["extensionId"] = A.load.Ref(target + 0, undefined);
        if (A.load.Bool(target + 20)) {
          target_ffi["tabId"] = A.load.Int64(target + 8);
        }
        target_ffi["targetId"] = A.load.Ref(target + 16, undefined);

        const _ret = WEBEXT.debugger.detach(target_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetTargets": (): heap.Ref<boolean> => {
      if (WEBEXT?.debugger && "getTargets" in WEBEXT?.debugger) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetTargets": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.debugger.getTargets);
    },
    "call_GetTargets": (retPtr: Pointer): void => {
      const _ret = WEBEXT.debugger.getTargets();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetTargets": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.debugger.getTargets();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDetach": (): heap.Ref<boolean> => {
      if (WEBEXT?.debugger?.onDetach && "addListener" in WEBEXT?.debugger?.onDetach) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDetach": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.debugger.onDetach.addListener);
    },
    "call_OnDetach": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.debugger.onDetach.addListener(A.H.get<object>(callback));
    },
    "try_OnDetach": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.debugger.onDetach.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDetach": (): heap.Ref<boolean> => {
      if (WEBEXT?.debugger?.onDetach && "removeListener" in WEBEXT?.debugger?.onDetach) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDetach": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.debugger.onDetach.removeListener);
    },
    "call_OffDetach": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.debugger.onDetach.removeListener(A.H.get<object>(callback));
    },
    "try_OffDetach": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.debugger.onDetach.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDetach": (): heap.Ref<boolean> => {
      if (WEBEXT?.debugger?.onDetach && "hasListener" in WEBEXT?.debugger?.onDetach) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDetach": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.debugger.onDetach.hasListener);
    },
    "call_HasOnDetach": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.debugger.onDetach.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDetach": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.debugger.onDetach.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.debugger?.onEvent && "addListener" in WEBEXT?.debugger?.onEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.debugger.onEvent.addListener);
    },
    "call_OnEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.debugger.onEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.debugger.onEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.debugger?.onEvent && "removeListener" in WEBEXT?.debugger?.onEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.debugger.onEvent.removeListener);
    },
    "call_OffEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.debugger.onEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.debugger.onEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.debugger?.onEvent && "hasListener" in WEBEXT?.debugger?.onEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.debugger.onEvent.hasListener);
    },
    "call_HasOnEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.debugger.onEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.debugger.onEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendCommand": (): heap.Ref<boolean> => {
      if (WEBEXT?.debugger && "sendCommand" in WEBEXT?.debugger) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendCommand": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.debugger.sendCommand);
    },
    "call_SendCommand": (
      retPtr: Pointer,
      target: Pointer,
      method: heap.Ref<object>,
      commandParams: heap.Ref<object>
    ): void => {
      const target_ffi = {};

      target_ffi["extensionId"] = A.load.Ref(target + 0, undefined);
      if (A.load.Bool(target + 20)) {
        target_ffi["tabId"] = A.load.Int64(target + 8);
      }
      target_ffi["targetId"] = A.load.Ref(target + 16, undefined);

      const _ret = WEBEXT.debugger.sendCommand(target_ffi, A.H.get<object>(method), A.H.get<object>(commandParams));
      A.store.Ref(retPtr, _ret);
    },
    "try_SendCommand": (
      retPtr: Pointer,
      errPtr: Pointer,
      target: Pointer,
      method: heap.Ref<object>,
      commandParams: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const target_ffi = {};

        target_ffi["extensionId"] = A.load.Ref(target + 0, undefined);
        if (A.load.Bool(target + 20)) {
          target_ffi["tabId"] = A.load.Int64(target + 8);
        }
        target_ffi["targetId"] = A.load.Ref(target + 16, undefined);

        const _ret = WEBEXT.debugger.sendCommand(target_ffi, A.H.get<object>(method), A.H.get<object>(commandParams));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
