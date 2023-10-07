import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/webauthenticationproxy", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_CreateRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "requestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Ref(ptr + 4, x["requestDetailsJson"]);
      }
    },
    "load_CreateRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["requestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["requestId"];
      }
      x["requestDetailsJson"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DOMExceptionDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["message"]);
      }
    },
    "load_DOMExceptionDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["message"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CreateResponseDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 20, false);
        A.store.Int32(ptr + 0, 0);

        A.store.Bool(ptr + 4 + 8, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Ref(ptr + 4 + 4, undefined);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Bool(ptr + 20, "requestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["requestId"] === undefined ? 0 : (x["requestId"] as number));

        if (typeof x["error"] === "undefined") {
          A.store.Bool(ptr + 4 + 8, false);
          A.store.Ref(ptr + 4 + 0, undefined);
          A.store.Ref(ptr + 4 + 4, undefined);
        } else {
          A.store.Bool(ptr + 4 + 8, true);
          A.store.Ref(ptr + 4 + 0, x["error"]["name"]);
          A.store.Ref(ptr + 4 + 4, x["error"]["message"]);
        }
        A.store.Ref(ptr + 16, x["responseJson"]);
      }
    },
    "load_CreateResponseDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 20)) {
        x["requestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["requestId"];
      }
      if (A.load.Bool(ptr + 4 + 8)) {
        x["error"] = {};
        x["error"]["name"] = A.load.Ref(ptr + 4 + 0, undefined);
        x["error"]["message"] = A.load.Ref(ptr + 4 + 4, undefined);
      } else {
        delete x["error"];
      }
      x["responseJson"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "requestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Ref(ptr + 4, x["requestDetailsJson"]);
      }
    },
    "load_GetRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["requestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["requestId"];
      }
      x["requestDetailsJson"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetResponseDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 20, false);
        A.store.Int32(ptr + 0, 0);

        A.store.Bool(ptr + 4 + 8, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Ref(ptr + 4 + 4, undefined);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Bool(ptr + 20, "requestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["requestId"] === undefined ? 0 : (x["requestId"] as number));

        if (typeof x["error"] === "undefined") {
          A.store.Bool(ptr + 4 + 8, false);
          A.store.Ref(ptr + 4 + 0, undefined);
          A.store.Ref(ptr + 4 + 4, undefined);
        } else {
          A.store.Bool(ptr + 4 + 8, true);
          A.store.Ref(ptr + 4 + 0, x["error"]["name"]);
          A.store.Ref(ptr + 4 + 4, x["error"]["message"]);
        }
        A.store.Ref(ptr + 16, x["responseJson"]);
      }
    },
    "load_GetResponseDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 20)) {
        x["requestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["requestId"];
      }
      if (A.load.Bool(ptr + 4 + 8)) {
        x["error"] = {};
        x["error"]["name"] = A.load.Ref(ptr + 4 + 0, undefined);
        x["error"]["message"] = A.load.Ref(ptr + 4 + 4, undefined);
      } else {
        delete x["error"];
      }
      x["responseJson"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_IsUvpaaRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "requestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
      }
    },
    "load_IsUvpaaRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["requestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["requestId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_IsUvpaaResponseDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 5, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 7, true);
        A.store.Bool(ptr + 5, "requestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Bool(ptr + 6, "isUvpaa" in x ? true : false);
        A.store.Bool(ptr + 4, x["isUvpaa"] ? true : false);
      }
    },
    "load_IsUvpaaResponseDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 5)) {
        x["requestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["requestId"];
      }
      if (A.load.Bool(ptr + 6)) {
        x["isUvpaa"] = A.load.Bool(ptr + 4);
      } else {
        delete x["isUvpaa"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Attach": (): heap.Ref<boolean> => {
      if (WEBEXT?.webAuthenticationProxy && "attach" in WEBEXT?.webAuthenticationProxy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Attach": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.attach);
    },
    "call_Attach": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webAuthenticationProxy.attach();
      A.store.Ref(retPtr, _ret);
    },
    "try_Attach": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.attach();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CompleteCreateRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.webAuthenticationProxy && "completeCreateRequest" in WEBEXT?.webAuthenticationProxy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CompleteCreateRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.completeCreateRequest);
    },
    "call_CompleteCreateRequest": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 20)) {
        details_ffi["requestId"] = A.load.Int32(details + 0);
      }
      if (A.load.Bool(details + 4 + 8)) {
        details_ffi["error"] = {};
        details_ffi["error"]["name"] = A.load.Ref(details + 4 + 0, undefined);
        details_ffi["error"]["message"] = A.load.Ref(details + 4 + 4, undefined);
      }
      details_ffi["responseJson"] = A.load.Ref(details + 16, undefined);

      const _ret = WEBEXT.webAuthenticationProxy.completeCreateRequest(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_CompleteCreateRequest": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 20)) {
          details_ffi["requestId"] = A.load.Int32(details + 0);
        }
        if (A.load.Bool(details + 4 + 8)) {
          details_ffi["error"] = {};
          details_ffi["error"]["name"] = A.load.Ref(details + 4 + 0, undefined);
          details_ffi["error"]["message"] = A.load.Ref(details + 4 + 4, undefined);
        }
        details_ffi["responseJson"] = A.load.Ref(details + 16, undefined);

        const _ret = WEBEXT.webAuthenticationProxy.completeCreateRequest(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CompleteGetRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.webAuthenticationProxy && "completeGetRequest" in WEBEXT?.webAuthenticationProxy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CompleteGetRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.completeGetRequest);
    },
    "call_CompleteGetRequest": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 20)) {
        details_ffi["requestId"] = A.load.Int32(details + 0);
      }
      if (A.load.Bool(details + 4 + 8)) {
        details_ffi["error"] = {};
        details_ffi["error"]["name"] = A.load.Ref(details + 4 + 0, undefined);
        details_ffi["error"]["message"] = A.load.Ref(details + 4 + 4, undefined);
      }
      details_ffi["responseJson"] = A.load.Ref(details + 16, undefined);

      const _ret = WEBEXT.webAuthenticationProxy.completeGetRequest(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_CompleteGetRequest": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 20)) {
          details_ffi["requestId"] = A.load.Int32(details + 0);
        }
        if (A.load.Bool(details + 4 + 8)) {
          details_ffi["error"] = {};
          details_ffi["error"]["name"] = A.load.Ref(details + 4 + 0, undefined);
          details_ffi["error"]["message"] = A.load.Ref(details + 4 + 4, undefined);
        }
        details_ffi["responseJson"] = A.load.Ref(details + 16, undefined);

        const _ret = WEBEXT.webAuthenticationProxy.completeGetRequest(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CompleteIsUvpaaRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.webAuthenticationProxy && "completeIsUvpaaRequest" in WEBEXT?.webAuthenticationProxy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CompleteIsUvpaaRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.completeIsUvpaaRequest);
    },
    "call_CompleteIsUvpaaRequest": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 5)) {
        details_ffi["requestId"] = A.load.Int32(details + 0);
      }
      if (A.load.Bool(details + 6)) {
        details_ffi["isUvpaa"] = A.load.Bool(details + 4);
      }

      const _ret = WEBEXT.webAuthenticationProxy.completeIsUvpaaRequest(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_CompleteIsUvpaaRequest": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 5)) {
          details_ffi["requestId"] = A.load.Int32(details + 0);
        }
        if (A.load.Bool(details + 6)) {
          details_ffi["isUvpaa"] = A.load.Bool(details + 4);
        }

        const _ret = WEBEXT.webAuthenticationProxy.completeIsUvpaaRequest(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Detach": (): heap.Ref<boolean> => {
      if (WEBEXT?.webAuthenticationProxy && "detach" in WEBEXT?.webAuthenticationProxy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Detach": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.detach);
    },
    "call_Detach": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webAuthenticationProxy.detach();
      A.store.Ref(retPtr, _ret);
    },
    "try_Detach": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.detach();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCreateRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webAuthenticationProxy?.onCreateRequest &&
        "addListener" in WEBEXT?.webAuthenticationProxy?.onCreateRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCreateRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.onCreateRequest.addListener);
    },
    "call_OnCreateRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webAuthenticationProxy.onCreateRequest.addListener(A.H.get<object>(callback));
    },
    "try_OnCreateRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.onCreateRequest.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCreateRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webAuthenticationProxy?.onCreateRequest &&
        "removeListener" in WEBEXT?.webAuthenticationProxy?.onCreateRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCreateRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.onCreateRequest.removeListener);
    },
    "call_OffCreateRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webAuthenticationProxy.onCreateRequest.removeListener(A.H.get<object>(callback));
    },
    "try_OffCreateRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.onCreateRequest.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCreateRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webAuthenticationProxy?.onCreateRequest &&
        "hasListener" in WEBEXT?.webAuthenticationProxy?.onCreateRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCreateRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.onCreateRequest.hasListener);
    },
    "call_HasOnCreateRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webAuthenticationProxy.onCreateRequest.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCreateRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.onCreateRequest.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnGetRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webAuthenticationProxy?.onGetRequest &&
        "addListener" in WEBEXT?.webAuthenticationProxy?.onGetRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnGetRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.onGetRequest.addListener);
    },
    "call_OnGetRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webAuthenticationProxy.onGetRequest.addListener(A.H.get<object>(callback));
    },
    "try_OnGetRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.onGetRequest.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffGetRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webAuthenticationProxy?.onGetRequest &&
        "removeListener" in WEBEXT?.webAuthenticationProxy?.onGetRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffGetRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.onGetRequest.removeListener);
    },
    "call_OffGetRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webAuthenticationProxy.onGetRequest.removeListener(A.H.get<object>(callback));
    },
    "try_OffGetRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.onGetRequest.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnGetRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webAuthenticationProxy?.onGetRequest &&
        "hasListener" in WEBEXT?.webAuthenticationProxy?.onGetRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnGetRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.onGetRequest.hasListener);
    },
    "call_HasOnGetRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webAuthenticationProxy.onGetRequest.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnGetRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.onGetRequest.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnIsUvpaaRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webAuthenticationProxy?.onIsUvpaaRequest &&
        "addListener" in WEBEXT?.webAuthenticationProxy?.onIsUvpaaRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnIsUvpaaRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.addListener);
    },
    "call_OnIsUvpaaRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.addListener(A.H.get<object>(callback));
    },
    "try_OnIsUvpaaRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffIsUvpaaRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webAuthenticationProxy?.onIsUvpaaRequest &&
        "removeListener" in WEBEXT?.webAuthenticationProxy?.onIsUvpaaRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffIsUvpaaRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.removeListener);
    },
    "call_OffIsUvpaaRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.removeListener(A.H.get<object>(callback));
    },
    "try_OffIsUvpaaRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnIsUvpaaRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webAuthenticationProxy?.onIsUvpaaRequest &&
        "hasListener" in WEBEXT?.webAuthenticationProxy?.onIsUvpaaRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnIsUvpaaRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.hasListener);
    },
    "call_HasOnIsUvpaaRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnIsUvpaaRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRemoteSessionStateChange": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webAuthenticationProxy?.onRemoteSessionStateChange &&
        "addListener" in WEBEXT?.webAuthenticationProxy?.onRemoteSessionStateChange
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRemoteSessionStateChange": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.addListener);
    },
    "call_OnRemoteSessionStateChange": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.addListener(A.H.get<object>(callback));
    },
    "try_OnRemoteSessionStateChange": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRemoteSessionStateChange": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webAuthenticationProxy?.onRemoteSessionStateChange &&
        "removeListener" in WEBEXT?.webAuthenticationProxy?.onRemoteSessionStateChange
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRemoteSessionStateChange": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.removeListener);
    },
    "call_OffRemoteSessionStateChange": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.removeListener(A.H.get<object>(callback));
    },
    "try_OffRemoteSessionStateChange": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRemoteSessionStateChange": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webAuthenticationProxy?.onRemoteSessionStateChange &&
        "hasListener" in WEBEXT?.webAuthenticationProxy?.onRemoteSessionStateChange
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRemoteSessionStateChange": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.hasListener);
    },
    "call_HasOnRemoteSessionStateChange": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRemoteSessionStateChange": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRequestCanceled": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webAuthenticationProxy?.onRequestCanceled &&
        "addListener" in WEBEXT?.webAuthenticationProxy?.onRequestCanceled
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRequestCanceled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.onRequestCanceled.addListener);
    },
    "call_OnRequestCanceled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webAuthenticationProxy.onRequestCanceled.addListener(A.H.get<object>(callback));
    },
    "try_OnRequestCanceled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.onRequestCanceled.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRequestCanceled": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webAuthenticationProxy?.onRequestCanceled &&
        "removeListener" in WEBEXT?.webAuthenticationProxy?.onRequestCanceled
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRequestCanceled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.onRequestCanceled.removeListener);
    },
    "call_OffRequestCanceled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webAuthenticationProxy.onRequestCanceled.removeListener(A.H.get<object>(callback));
    },
    "try_OffRequestCanceled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.onRequestCanceled.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRequestCanceled": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webAuthenticationProxy?.onRequestCanceled &&
        "hasListener" in WEBEXT?.webAuthenticationProxy?.onRequestCanceled
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRequestCanceled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webAuthenticationProxy.onRequestCanceled.hasListener);
    },
    "call_HasOnRequestCanceled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webAuthenticationProxy.onRequestCanceled.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRequestCanceled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webAuthenticationProxy.onRequestCanceled.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
