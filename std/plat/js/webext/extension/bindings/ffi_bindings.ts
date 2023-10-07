import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/extension", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_ViewType": (ref: heap.Ref<string>): number => {
      const idx = ["tab", "popup"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_GetViewsArgFetchProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 24, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Enum(ptr + 8, -1);
        A.store.Bool(ptr + 25, false);
        A.store.Int64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 26, true);
        A.store.Bool(ptr + 24, "tabId" in x ? true : false);
        A.store.Int64(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Enum(ptr + 8, ["tab", "popup"].indexOf(x["type"] as string));
        A.store.Bool(ptr + 25, "windowId" in x ? true : false);
        A.store.Int64(ptr + 16, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_GetViewsArgFetchProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["tabId"] = A.load.Int64(ptr + 0);
      } else {
        delete x["tabId"];
      }
      x["type"] = A.load.Enum(ptr + 8, ["tab", "popup"]);
      if (A.load.Bool(ptr + 25)) {
        x["windowId"] = A.load.Int64(ptr + 16);
      } else {
        delete x["windowId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_LastErrorProperty": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["message"]);
      }
    },
    "load_LastErrorProperty": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["message"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetBackgroundPage": (): heap.Ref<boolean> => {
      if (WEBEXT?.extension && "getBackgroundPage" in WEBEXT?.extension) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetBackgroundPage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extension.getBackgroundPage);
    },
    "call_GetBackgroundPage": (retPtr: Pointer): void => {
      const _ret = WEBEXT.extension.getBackgroundPage();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetBackgroundPage": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extension.getBackgroundPage();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetExtensionTabs": (): heap.Ref<boolean> => {
      if (WEBEXT?.extension && "getExtensionTabs" in WEBEXT?.extension) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetExtensionTabs": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extension.getExtensionTabs);
    },
    "call_GetExtensionTabs": (retPtr: Pointer, windowId: number): void => {
      const _ret = WEBEXT.extension.getExtensionTabs(windowId);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetExtensionTabs": (retPtr: Pointer, errPtr: Pointer, windowId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extension.getExtensionTabs(windowId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetURL": (): heap.Ref<boolean> => {
      if (WEBEXT?.extension && "getURL" in WEBEXT?.extension) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetURL": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extension.getURL);
    },
    "call_GetURL": (retPtr: Pointer, path: heap.Ref<object>): void => {
      const _ret = WEBEXT.extension.getURL(A.H.get<object>(path));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetURL": (retPtr: Pointer, errPtr: Pointer, path: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extension.getURL(A.H.get<object>(path));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetViews": (): heap.Ref<boolean> => {
      if (WEBEXT?.extension && "getViews" in WEBEXT?.extension) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetViews": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extension.getViews);
    },
    "call_GetViews": (retPtr: Pointer, fetchProperties: Pointer): void => {
      const fetchProperties_ffi = {};

      if (A.load.Bool(fetchProperties + 24)) {
        fetchProperties_ffi["tabId"] = A.load.Int64(fetchProperties + 0);
      }
      fetchProperties_ffi["type"] = A.load.Enum(fetchProperties + 8, ["tab", "popup"]);
      if (A.load.Bool(fetchProperties + 25)) {
        fetchProperties_ffi["windowId"] = A.load.Int64(fetchProperties + 16);
      }

      const _ret = WEBEXT.extension.getViews(fetchProperties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetViews": (retPtr: Pointer, errPtr: Pointer, fetchProperties: Pointer): heap.Ref<boolean> => {
      try {
        const fetchProperties_ffi = {};

        if (A.load.Bool(fetchProperties + 24)) {
          fetchProperties_ffi["tabId"] = A.load.Int64(fetchProperties + 0);
        }
        fetchProperties_ffi["type"] = A.load.Enum(fetchProperties + 8, ["tab", "popup"]);
        if (A.load.Bool(fetchProperties + 25)) {
          fetchProperties_ffi["windowId"] = A.load.Int64(fetchProperties + 16);
        }

        const _ret = WEBEXT.extension.getViews(fetchProperties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_InIncognitoContext": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.extension && "inIncognitoContext" in WEBEXT?.extension) {
        const val = WEBEXT.extension.inIncognitoContext;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_InIncognitoContext": (val: heap.Ref<boolean>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.extension, "inIncognitoContext", val === A.H.TRUE, WEBEXT.extension)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "has_IsAllowedFileSchemeAccess": (): heap.Ref<boolean> => {
      if (WEBEXT?.extension && "isAllowedFileSchemeAccess" in WEBEXT?.extension) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsAllowedFileSchemeAccess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extension.isAllowedFileSchemeAccess);
    },
    "call_IsAllowedFileSchemeAccess": (retPtr: Pointer): void => {
      const _ret = WEBEXT.extension.isAllowedFileSchemeAccess();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsAllowedFileSchemeAccess": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extension.isAllowedFileSchemeAccess();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsAllowedIncognitoAccess": (): heap.Ref<boolean> => {
      if (WEBEXT?.extension && "isAllowedIncognitoAccess" in WEBEXT?.extension) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsAllowedIncognitoAccess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extension.isAllowedIncognitoAccess);
    },
    "call_IsAllowedIncognitoAccess": (retPtr: Pointer): void => {
      const _ret = WEBEXT.extension.isAllowedIncognitoAccess();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsAllowedIncognitoAccess": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extension.isAllowedIncognitoAccess();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_LastError": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.extension && "lastError" in WEBEXT?.extension) {
        const val = WEBEXT.extension.lastError;
        if (typeof val === "undefined") {
          A.store.Bool(retPtr + 4, false);
          A.store.Ref(retPtr + 0, undefined);
        } else {
          A.store.Bool(retPtr + 4, true);
          A.store.Ref(retPtr + 0, val["message"]);
        }
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_LastError": (val: Pointer): heap.Ref<boolean> => {
      const val_ffi = {};

      val_ffi["message"] = A.load.Ref(val + 0, undefined);
      return Reflect.set(WEBEXT.extension, "lastError", val_ffi, WEBEXT.extension) ? A.H.TRUE : A.H.FALSE;
    },
    "has_OnRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.extension?.onRequest && "addListener" in WEBEXT?.extension?.onRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extension.onRequest.addListener);
    },
    "call_OnRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.extension.onRequest.addListener(A.H.get<object>(callback));
    },
    "try_OnRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extension.onRequest.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.extension?.onRequest && "removeListener" in WEBEXT?.extension?.onRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extension.onRequest.removeListener);
    },
    "call_OffRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.extension.onRequest.removeListener(A.H.get<object>(callback));
    },
    "try_OffRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extension.onRequest.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.extension?.onRequest && "hasListener" in WEBEXT?.extension?.onRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extension.onRequest.hasListener);
    },
    "call_HasOnRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.extension.onRequest.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extension.onRequest.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRequestExternal": (): heap.Ref<boolean> => {
      if (WEBEXT?.extension?.onRequestExternal && "addListener" in WEBEXT?.extension?.onRequestExternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRequestExternal": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extension.onRequestExternal.addListener);
    },
    "call_OnRequestExternal": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.extension.onRequestExternal.addListener(A.H.get<object>(callback));
    },
    "try_OnRequestExternal": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extension.onRequestExternal.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRequestExternal": (): heap.Ref<boolean> => {
      if (WEBEXT?.extension?.onRequestExternal && "removeListener" in WEBEXT?.extension?.onRequestExternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRequestExternal": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extension.onRequestExternal.removeListener);
    },
    "call_OffRequestExternal": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.extension.onRequestExternal.removeListener(A.H.get<object>(callback));
    },
    "try_OffRequestExternal": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extension.onRequestExternal.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRequestExternal": (): heap.Ref<boolean> => {
      if (WEBEXT?.extension?.onRequestExternal && "hasListener" in WEBEXT?.extension?.onRequestExternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRequestExternal": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extension.onRequestExternal.hasListener);
    },
    "call_HasOnRequestExternal": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.extension.onRequestExternal.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRequestExternal": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extension.onRequestExternal.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.extension && "sendRequest" in WEBEXT?.extension) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extension.sendRequest);
    },
    "call_SendRequest": (retPtr: Pointer, extensionId: heap.Ref<object>, request: heap.Ref<object>): void => {
      const _ret = WEBEXT.extension.sendRequest(A.H.get<object>(extensionId), A.H.get<object>(request));
      A.store.Ref(retPtr, _ret);
    },
    "try_SendRequest": (
      retPtr: Pointer,
      errPtr: Pointer,
      extensionId: heap.Ref<object>,
      request: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extension.sendRequest(A.H.get<object>(extensionId), A.H.get<object>(request));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetUpdateUrlData": (): heap.Ref<boolean> => {
      if (WEBEXT?.extension && "setUpdateUrlData" in WEBEXT?.extension) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetUpdateUrlData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extension.setUpdateUrlData);
    },
    "call_SetUpdateUrlData": (retPtr: Pointer, data: heap.Ref<object>): void => {
      const _ret = WEBEXT.extension.setUpdateUrlData(A.H.get<object>(data));
    },
    "try_SetUpdateUrlData": (retPtr: Pointer, errPtr: Pointer, data: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extension.setUpdateUrlData(A.H.get<object>(data));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
