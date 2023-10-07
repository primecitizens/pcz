import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/permissions", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Permissions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["origins"]);
        A.store.Ref(ptr + 4, x["permissions"]);
      }
    },
    "load_Permissions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["origins"] = A.load.Ref(ptr + 0, undefined);
      x["permissions"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Contains": (): heap.Ref<boolean> => {
      if (WEBEXT?.permissions && "contains" in WEBEXT?.permissions) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Contains": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.permissions.contains);
    },
    "call_Contains": (retPtr: Pointer, permissions: Pointer): void => {
      const permissions_ffi = {};

      permissions_ffi["origins"] = A.load.Ref(permissions + 0, undefined);
      permissions_ffi["permissions"] = A.load.Ref(permissions + 4, undefined);

      const _ret = WEBEXT.permissions.contains(permissions_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Contains": (retPtr: Pointer, errPtr: Pointer, permissions: Pointer): heap.Ref<boolean> => {
      try {
        const permissions_ffi = {};

        permissions_ffi["origins"] = A.load.Ref(permissions + 0, undefined);
        permissions_ffi["permissions"] = A.load.Ref(permissions + 4, undefined);

        const _ret = WEBEXT.permissions.contains(permissions_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAll": (): heap.Ref<boolean> => {
      if (WEBEXT?.permissions && "getAll" in WEBEXT?.permissions) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.permissions.getAll);
    },
    "call_GetAll": (retPtr: Pointer): void => {
      const _ret = WEBEXT.permissions.getAll();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAll": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.permissions.getAll();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.permissions?.onAdded && "addListener" in WEBEXT?.permissions?.onAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.permissions.onAdded.addListener);
    },
    "call_OnAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.permissions.onAdded.addListener(A.H.get<object>(callback));
    },
    "try_OnAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.permissions.onAdded.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.permissions?.onAdded && "removeListener" in WEBEXT?.permissions?.onAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.permissions.onAdded.removeListener);
    },
    "call_OffAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.permissions.onAdded.removeListener(A.H.get<object>(callback));
    },
    "try_OffAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.permissions.onAdded.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.permissions?.onAdded && "hasListener" in WEBEXT?.permissions?.onAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.permissions.onAdded.hasListener);
    },
    "call_HasOnAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.permissions.onAdded.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.permissions.onAdded.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.permissions?.onRemoved && "addListener" in WEBEXT?.permissions?.onRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.permissions.onRemoved.addListener);
    },
    "call_OnRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.permissions.onRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.permissions.onRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.permissions?.onRemoved && "removeListener" in WEBEXT?.permissions?.onRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.permissions.onRemoved.removeListener);
    },
    "call_OffRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.permissions.onRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.permissions.onRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.permissions?.onRemoved && "hasListener" in WEBEXT?.permissions?.onRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.permissions.onRemoved.hasListener);
    },
    "call_HasOnRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.permissions.onRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.permissions.onRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Remove": (): heap.Ref<boolean> => {
      if (WEBEXT?.permissions && "remove" in WEBEXT?.permissions) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Remove": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.permissions.remove);
    },
    "call_Remove": (retPtr: Pointer, permissions: Pointer): void => {
      const permissions_ffi = {};

      permissions_ffi["origins"] = A.load.Ref(permissions + 0, undefined);
      permissions_ffi["permissions"] = A.load.Ref(permissions + 4, undefined);

      const _ret = WEBEXT.permissions.remove(permissions_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Remove": (retPtr: Pointer, errPtr: Pointer, permissions: Pointer): heap.Ref<boolean> => {
      try {
        const permissions_ffi = {};

        permissions_ffi["origins"] = A.load.Ref(permissions + 0, undefined);
        permissions_ffi["permissions"] = A.load.Ref(permissions + 4, undefined);

        const _ret = WEBEXT.permissions.remove(permissions_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Request": (): heap.Ref<boolean> => {
      if (WEBEXT?.permissions && "request" in WEBEXT?.permissions) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Request": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.permissions.request);
    },
    "call_Request": (retPtr: Pointer, permissions: Pointer): void => {
      const permissions_ffi = {};

      permissions_ffi["origins"] = A.load.Ref(permissions + 0, undefined);
      permissions_ffi["permissions"] = A.load.Ref(permissions + 4, undefined);

      const _ret = WEBEXT.permissions.request(permissions_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Request": (retPtr: Pointer, errPtr: Pointer, permissions: Pointer): heap.Ref<boolean> => {
      try {
        const permissions_ffi = {};

        permissions_ffi["origins"] = A.load.Ref(permissions + 0, undefined);
        permissions_ffi["permissions"] = A.load.Ref(permissions + 4, undefined);

        const _ret = WEBEXT.permissions.request(permissions_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
