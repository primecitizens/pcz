import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/storage", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_AccessLevel": (ref: heap.Ref<string>): number => {
      const idx = ["TRUSTED_CONTEXTS", "TRUSTED_AND_UNTRUSTED_CONTEXTS"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SetAccessLevelArgAccessOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(
          ptr + 0,
          ["TRUSTED_CONTEXTS", "TRUSTED_AND_UNTRUSTED_CONTEXTS"].indexOf(x["accessLevel"] as string)
        );
      }
    },
    "load_SetAccessLevelArgAccessOptions": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["accessLevel"] = A.load.Enum(ptr + 0, ["TRUSTED_CONTEXTS", "TRUSTED_AND_UNTRUSTED_CONTEXTS"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "has_StorageArea_Get": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "get" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StorageArea_Get": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).get);
    },

    "call_StorageArea_Get": (self: heap.Ref<object>, retPtr: Pointer, keys: heap.Ref<any>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.get(A.H.get<any>(keys));
      A.store.Ref(retPtr, _ret);
    },
    "try_StorageArea_Get": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      keys: heap.Ref<any>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.get(A.H.get<any>(keys));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StorageArea_Get1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "get" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StorageArea_Get1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).get);
    },

    "call_StorageArea_Get1": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.get();
      A.store.Ref(retPtr, _ret);
    },
    "try_StorageArea_Get1": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.get();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StorageArea_GetBytesInUse": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "getBytesInUse" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StorageArea_GetBytesInUse": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).getBytesInUse);
    },

    "call_StorageArea_GetBytesInUse": (self: heap.Ref<object>, retPtr: Pointer, keys: heap.Ref<any>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.getBytesInUse(A.H.get<any>(keys));
      A.store.Ref(retPtr, _ret);
    },
    "try_StorageArea_GetBytesInUse": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      keys: heap.Ref<any>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.getBytesInUse(A.H.get<any>(keys));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StorageArea_GetBytesInUse1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "getBytesInUse" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StorageArea_GetBytesInUse1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).getBytesInUse);
    },

    "call_StorageArea_GetBytesInUse1": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.getBytesInUse();
      A.store.Ref(retPtr, _ret);
    },
    "try_StorageArea_GetBytesInUse1": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.getBytesInUse();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StorageArea_Set": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "set" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StorageArea_Set": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).set);
    },

    "call_StorageArea_Set": (self: heap.Ref<object>, retPtr: Pointer, items: heap.Ref<object>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.set(A.H.get<object>(items));
      A.store.Ref(retPtr, _ret);
    },
    "try_StorageArea_Set": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      items: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.set(A.H.get<object>(items));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StorageArea_Remove": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "remove" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StorageArea_Remove": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).remove);
    },

    "call_StorageArea_Remove": (self: heap.Ref<object>, retPtr: Pointer, keys: heap.Ref<any>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.remove(A.H.get<any>(keys));
      A.store.Ref(retPtr, _ret);
    },
    "try_StorageArea_Remove": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      keys: heap.Ref<any>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.remove(A.H.get<any>(keys));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StorageArea_Clear": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "clear" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StorageArea_Clear": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).clear);
    },

    "call_StorageArea_Clear": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.clear();
      A.store.Ref(retPtr, _ret);
    },
    "try_StorageArea_Clear": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.clear();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StorageArea_SetAccessLevel": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "setAccessLevel" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StorageArea_SetAccessLevel": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).setAccessLevel);
    },

    "call_StorageArea_SetAccessLevel": (self: heap.Ref<object>, retPtr: Pointer, accessOptions: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const accessOptions_ffi = {};

      accessOptions_ffi["accessLevel"] = A.load.Enum(accessOptions + 0, [
        "TRUSTED_CONTEXTS",
        "TRUSTED_AND_UNTRUSTED_CONTEXTS",
      ]);
      const _ret = thiz.setAccessLevel(accessOptions_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_StorageArea_SetAccessLevel": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      accessOptions: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const accessOptions_ffi = {};

        accessOptions_ffi["accessLevel"] = A.load.Enum(accessOptions + 0, [
          "TRUSTED_CONTEXTS",
          "TRUSTED_AND_UNTRUSTED_CONTEXTS",
        ]);
        const _ret = thiz.setAccessLevel(accessOptions_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },

    "store_StorageChange": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["newValue"]);
        A.store.Ref(ptr + 4, x["oldValue"]);
      }
    },
    "load_StorageChange": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["newValue"] = A.load.Ref(ptr + 0, undefined);
      x["oldValue"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "get_Local": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.storage && "local" in WEBEXT?.storage) {
        const val = WEBEXT.storage.local;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Local": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.storage, "local", A.H.get<object>(val), WEBEXT.storage) ? A.H.TRUE : A.H.FALSE;
    },
    "get_Managed": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.storage && "managed" in WEBEXT?.storage) {
        const val = WEBEXT.storage.managed;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Managed": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.storage, "managed", A.H.get<object>(val), WEBEXT.storage) ? A.H.TRUE : A.H.FALSE;
    },
    "has_OnChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.storage?.onChanged && "addListener" in WEBEXT?.storage?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.storage.onChanged.addListener);
    },
    "call_OnChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.storage.onChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.storage.onChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.storage?.onChanged && "removeListener" in WEBEXT?.storage?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.storage.onChanged.removeListener);
    },
    "call_OffChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.storage.onChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.storage.onChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.storage?.onChanged && "hasListener" in WEBEXT?.storage?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.storage.onChanged.hasListener);
    },
    "call_HasOnChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.storage.onChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.storage.onChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_Session": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.storage && "session" in WEBEXT?.storage) {
        const val = WEBEXT.storage.session;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Session": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.storage, "session", A.H.get<object>(val), WEBEXT.storage) ? A.H.TRUE : A.H.FALSE;
    },
    "get_Sync": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.storage && "sync" in WEBEXT?.storage) {
        const val = WEBEXT.storage.sync;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Sync": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.storage, "sync", A.H.get<object>(val), WEBEXT.storage) ? A.H.TRUE : A.H.FALSE;
    },
  };
});
