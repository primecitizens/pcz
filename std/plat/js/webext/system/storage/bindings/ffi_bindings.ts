import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/system/storage", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_EjectDeviceResultCode": (ref: heap.Ref<string>): number => {
      const idx = ["success", "in_use", "no_such_device", "failure"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_StorageAvailableCapacityInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Bool(ptr + 16, "availableCapacity" in x ? true : false);
        A.store.Float64(ptr + 8, x["availableCapacity"] === undefined ? 0 : (x["availableCapacity"] as number));
      }
    },
    "load_StorageAvailableCapacityInfo": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["availableCapacity"] = A.load.Float64(ptr + 8);
      } else {
        delete x["availableCapacity"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_StorageUnitType": (ref: heap.Ref<string>): number => {
      const idx = ["fixed", "removable", "unknown"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_StorageUnitInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Bool(ptr + 24, false);
        A.store.Float64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 25, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Enum(ptr + 8, ["fixed", "removable", "unknown"].indexOf(x["type"] as string));
        A.store.Bool(ptr + 24, "capacity" in x ? true : false);
        A.store.Float64(ptr + 16, x["capacity"] === undefined ? 0 : (x["capacity"] as number));
      }
    },
    "load_StorageUnitInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      x["type"] = A.load.Enum(ptr + 8, ["fixed", "removable", "unknown"]);
      if (A.load.Bool(ptr + 24)) {
        x["capacity"] = A.load.Float64(ptr + 16);
      } else {
        delete x["capacity"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_EjectDevice": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.storage && "ejectDevice" in WEBEXT?.system?.storage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EjectDevice": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.storage.ejectDevice);
    },
    "call_EjectDevice": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.storage.ejectDevice(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_EjectDevice": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.storage.ejectDevice(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAvailableCapacity": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.storage && "getAvailableCapacity" in WEBEXT?.system?.storage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAvailableCapacity": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.storage.getAvailableCapacity);
    },
    "call_GetAvailableCapacity": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.storage.getAvailableCapacity(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAvailableCapacity": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.storage.getAvailableCapacity(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.storage && "getInfo" in WEBEXT?.system?.storage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.storage.getInfo);
    },
    "call_GetInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.system.storage.getInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.storage.getInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAttached": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.storage?.onAttached && "addListener" in WEBEXT?.system?.storage?.onAttached) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAttached": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.storage.onAttached.addListener);
    },
    "call_OnAttached": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.storage.onAttached.addListener(A.H.get<object>(callback));
    },
    "try_OnAttached": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.storage.onAttached.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAttached": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.storage?.onAttached && "removeListener" in WEBEXT?.system?.storage?.onAttached) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAttached": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.storage.onAttached.removeListener);
    },
    "call_OffAttached": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.storage.onAttached.removeListener(A.H.get<object>(callback));
    },
    "try_OffAttached": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.storage.onAttached.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAttached": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.storage?.onAttached && "hasListener" in WEBEXT?.system?.storage?.onAttached) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAttached": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.storage.onAttached.hasListener);
    },
    "call_HasOnAttached": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.storage.onAttached.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAttached": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.storage.onAttached.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDetached": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.storage?.onDetached && "addListener" in WEBEXT?.system?.storage?.onDetached) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDetached": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.storage.onDetached.addListener);
    },
    "call_OnDetached": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.storage.onDetached.addListener(A.H.get<object>(callback));
    },
    "try_OnDetached": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.storage.onDetached.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDetached": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.storage?.onDetached && "removeListener" in WEBEXT?.system?.storage?.onDetached) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDetached": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.storage.onDetached.removeListener);
    },
    "call_OffDetached": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.storage.onDetached.removeListener(A.H.get<object>(callback));
    },
    "try_OffDetached": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.storage.onDetached.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDetached": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.storage?.onDetached && "hasListener" in WEBEXT?.system?.storage?.onDetached) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDetached": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.storage.onDetached.hasListener);
    },
    "call_HasOnDetached": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.storage.onDetached.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDetached": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.storage.onDetached.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
