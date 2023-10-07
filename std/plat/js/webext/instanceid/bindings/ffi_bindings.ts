import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/instanceid", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_DeleteTokenArgDeleteTokenParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["authorizedEntity"]);
        A.store.Ref(ptr + 4, x["scope"]);
      }
    },
    "load_DeleteTokenArgDeleteTokenParams": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["authorizedEntity"] = A.load.Ref(ptr + 0, undefined);
      x["scope"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetTokenArgGetTokenParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["authorizedEntity"]);
        A.store.Ref(ptr + 4, x["options"]);
        A.store.Ref(ptr + 8, x["scope"]);
      }
    },
    "load_GetTokenArgGetTokenParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["authorizedEntity"] = A.load.Ref(ptr + 0, undefined);
      x["options"] = A.load.Ref(ptr + 4, undefined);
      x["scope"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_DeleteID": (): heap.Ref<boolean> => {
      if (WEBEXT?.instanceID && "deleteID" in WEBEXT?.instanceID) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DeleteID": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.instanceID.deleteID);
    },
    "call_DeleteID": (retPtr: Pointer): void => {
      const _ret = WEBEXT.instanceID.deleteID();
      A.store.Ref(retPtr, _ret);
    },
    "try_DeleteID": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.instanceID.deleteID();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DeleteToken": (): heap.Ref<boolean> => {
      if (WEBEXT?.instanceID && "deleteToken" in WEBEXT?.instanceID) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DeleteToken": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.instanceID.deleteToken);
    },
    "call_DeleteToken": (retPtr: Pointer, deleteTokenParams: Pointer): void => {
      const deleteTokenParams_ffi = {};

      deleteTokenParams_ffi["authorizedEntity"] = A.load.Ref(deleteTokenParams + 0, undefined);
      deleteTokenParams_ffi["scope"] = A.load.Ref(deleteTokenParams + 4, undefined);

      const _ret = WEBEXT.instanceID.deleteToken(deleteTokenParams_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_DeleteToken": (retPtr: Pointer, errPtr: Pointer, deleteTokenParams: Pointer): heap.Ref<boolean> => {
      try {
        const deleteTokenParams_ffi = {};

        deleteTokenParams_ffi["authorizedEntity"] = A.load.Ref(deleteTokenParams + 0, undefined);
        deleteTokenParams_ffi["scope"] = A.load.Ref(deleteTokenParams + 4, undefined);

        const _ret = WEBEXT.instanceID.deleteToken(deleteTokenParams_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCreationTime": (): heap.Ref<boolean> => {
      if (WEBEXT?.instanceID && "getCreationTime" in WEBEXT?.instanceID) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCreationTime": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.instanceID.getCreationTime);
    },
    "call_GetCreationTime": (retPtr: Pointer): void => {
      const _ret = WEBEXT.instanceID.getCreationTime();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCreationTime": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.instanceID.getCreationTime();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetID": (): heap.Ref<boolean> => {
      if (WEBEXT?.instanceID && "getID" in WEBEXT?.instanceID) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetID": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.instanceID.getID);
    },
    "call_GetID": (retPtr: Pointer): void => {
      const _ret = WEBEXT.instanceID.getID();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetID": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.instanceID.getID();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetToken": (): heap.Ref<boolean> => {
      if (WEBEXT?.instanceID && "getToken" in WEBEXT?.instanceID) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetToken": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.instanceID.getToken);
    },
    "call_GetToken": (retPtr: Pointer, getTokenParams: Pointer): void => {
      const getTokenParams_ffi = {};

      getTokenParams_ffi["authorizedEntity"] = A.load.Ref(getTokenParams + 0, undefined);
      getTokenParams_ffi["options"] = A.load.Ref(getTokenParams + 4, undefined);
      getTokenParams_ffi["scope"] = A.load.Ref(getTokenParams + 8, undefined);

      const _ret = WEBEXT.instanceID.getToken(getTokenParams_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetToken": (retPtr: Pointer, errPtr: Pointer, getTokenParams: Pointer): heap.Ref<boolean> => {
      try {
        const getTokenParams_ffi = {};

        getTokenParams_ffi["authorizedEntity"] = A.load.Ref(getTokenParams + 0, undefined);
        getTokenParams_ffi["options"] = A.load.Ref(getTokenParams + 4, undefined);
        getTokenParams_ffi["scope"] = A.load.Ref(getTokenParams + 8, undefined);

        const _ret = WEBEXT.instanceID.getToken(getTokenParams_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTokenRefresh": (): heap.Ref<boolean> => {
      if (WEBEXT?.instanceID?.onTokenRefresh && "addListener" in WEBEXT?.instanceID?.onTokenRefresh) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTokenRefresh": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.instanceID.onTokenRefresh.addListener);
    },
    "call_OnTokenRefresh": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.instanceID.onTokenRefresh.addListener(A.H.get<object>(callback));
    },
    "try_OnTokenRefresh": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.instanceID.onTokenRefresh.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffTokenRefresh": (): heap.Ref<boolean> => {
      if (WEBEXT?.instanceID?.onTokenRefresh && "removeListener" in WEBEXT?.instanceID?.onTokenRefresh) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffTokenRefresh": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.instanceID.onTokenRefresh.removeListener);
    },
    "call_OffTokenRefresh": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.instanceID.onTokenRefresh.removeListener(A.H.get<object>(callback));
    },
    "try_OffTokenRefresh": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.instanceID.onTokenRefresh.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnTokenRefresh": (): heap.Ref<boolean> => {
      if (WEBEXT?.instanceID?.onTokenRefresh && "hasListener" in WEBEXT?.instanceID?.onTokenRefresh) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnTokenRefresh": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.instanceID.onTokenRefresh.hasListener);
    },
    "call_HasOnTokenRefresh": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.instanceID.onTokenRefresh.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnTokenRefresh": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.instanceID.onTokenRefresh.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
