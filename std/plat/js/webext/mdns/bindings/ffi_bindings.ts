import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/mdns", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_MDnsService": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["serviceName"]);
        A.store.Ref(ptr + 4, x["serviceHostPort"]);
        A.store.Ref(ptr + 8, x["ipAddress"]);
        A.store.Ref(ptr + 12, x["serviceData"]);
      }
    },
    "load_MDnsService": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["serviceName"] = A.load.Ref(ptr + 0, undefined);
      x["serviceHostPort"] = A.load.Ref(ptr + 4, undefined);
      x["ipAddress"] = A.load.Ref(ptr + 8, undefined);
      x["serviceData"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "has_Properties_MAX_SERVICE_INSTANCES_PER_EVENT": (self: heap.Ref<any>): heap.Ref<boolean> => {
      if (WEBEXT?.mdns && "MAX_SERVICE_INSTANCES_PER_EVENT" in WEBEXT?.mdns) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Properties_MAX_SERVICE_INSTANCES_PER_EVENT": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mdns.MAX_SERVICE_INSTANCES_PER_EVENT);
    },

    "call_Properties_MAX_SERVICE_INSTANCES_PER_EVENT": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = Reflect.apply(WEBEXT.mdns.MAX_SERVICE_INSTANCES_PER_EVENT, thiz, []);
      A.store.Int32(retPtr, _ret);
    },
    "try_Properties_MAX_SERVICE_INSTANCES_PER_EVENT": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = Reflect.apply(WEBEXT.mdns.MAX_SERVICE_INSTANCES_PER_EVENT, thiz, []);
        A.store.Int32(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ForceDiscovery": (): heap.Ref<boolean> => {
      if (WEBEXT?.mdns && "forceDiscovery" in WEBEXT?.mdns) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ForceDiscovery": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mdns.forceDiscovery);
    },
    "call_ForceDiscovery": (retPtr: Pointer): void => {
      const _ret = WEBEXT.mdns.forceDiscovery();
      A.store.Ref(retPtr, _ret);
    },
    "try_ForceDiscovery": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mdns.forceDiscovery();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnServiceList": (): heap.Ref<boolean> => {
      if (WEBEXT?.mdns?.onServiceList && "addListener" in WEBEXT?.mdns?.onServiceList) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnServiceList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mdns.onServiceList.addListener);
    },
    "call_OnServiceList": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mdns.onServiceList.addListener(A.H.get<object>(callback));
    },
    "try_OnServiceList": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mdns.onServiceList.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffServiceList": (): heap.Ref<boolean> => {
      if (WEBEXT?.mdns?.onServiceList && "removeListener" in WEBEXT?.mdns?.onServiceList) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffServiceList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mdns.onServiceList.removeListener);
    },
    "call_OffServiceList": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mdns.onServiceList.removeListener(A.H.get<object>(callback));
    },
    "try_OffServiceList": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mdns.onServiceList.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnServiceList": (): heap.Ref<boolean> => {
      if (WEBEXT?.mdns?.onServiceList && "hasListener" in WEBEXT?.mdns?.onServiceList) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnServiceList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mdns.onServiceList.hasListener);
    },
    "call_HasOnServiceList": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mdns.onServiceList.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnServiceList": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mdns.onServiceList.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
