import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/webrequestinternal", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_AddEventListenerOptions": (ref: heap.Ref<string>): number => {
      const idx = [
        "requestHeaders",
        "responseHeaders",
        "blocking",
        "asyncBlocking",
        "requestBody",
        "extraHeaders",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_AddEventListener": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequestInternal && "addEventListener" in WEBEXT?.webRequestInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddEventListener": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequestInternal.addEventListener);
    },
    "call_AddEventListener": (
      retPtr: Pointer,
      callback: heap.Ref<object>,
      filter: Pointer,
      extraInfoSpec: heap.Ref<object>,
      eventName: heap.Ref<object>,
      subEventName: heap.Ref<object>,
      webViewInstanceId: number
    ): void => {
      const filter_ffi = {};

      if (A.load.Bool(filter + 24)) {
        filter_ffi["tabId"] = A.load.Int64(filter + 0);
      }
      filter_ffi["types"] = A.load.Ref(filter + 8, undefined);
      filter_ffi["urls"] = A.load.Ref(filter + 12, undefined);
      if (A.load.Bool(filter + 25)) {
        filter_ffi["windowId"] = A.load.Int64(filter + 16);
      }

      const _ret = WEBEXT.webRequestInternal.addEventListener(
        A.H.get<object>(callback),
        filter_ffi,
        A.H.get<object>(extraInfoSpec),
        A.H.get<object>(eventName),
        A.H.get<object>(subEventName),
        webViewInstanceId
      );
    },
    "try_AddEventListener": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>,
      filter: Pointer,
      extraInfoSpec: heap.Ref<object>,
      eventName: heap.Ref<object>,
      subEventName: heap.Ref<object>,
      webViewInstanceId: number
    ): heap.Ref<boolean> => {
      try {
        const filter_ffi = {};

        if (A.load.Bool(filter + 24)) {
          filter_ffi["tabId"] = A.load.Int64(filter + 0);
        }
        filter_ffi["types"] = A.load.Ref(filter + 8, undefined);
        filter_ffi["urls"] = A.load.Ref(filter + 12, undefined);
        if (A.load.Bool(filter + 25)) {
          filter_ffi["windowId"] = A.load.Int64(filter + 16);
        }

        const _ret = WEBEXT.webRequestInternal.addEventListener(
          A.H.get<object>(callback),
          filter_ffi,
          A.H.get<object>(extraInfoSpec),
          A.H.get<object>(eventName),
          A.H.get<object>(subEventName),
          webViewInstanceId
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_EventHandled": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequestInternal && "eventHandled" in WEBEXT?.webRequestInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EventHandled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequestInternal.eventHandled);
    },
    "call_EventHandled": (
      retPtr: Pointer,
      eventName: heap.Ref<object>,
      subEventName: heap.Ref<object>,
      requestId: heap.Ref<object>,
      webViewInstanceId: number,
      response: Pointer
    ): void => {
      const response_ffi = {};

      if (A.load.Bool(response + 0 + 8)) {
        response_ffi["authCredentials"] = {};
        response_ffi["authCredentials"]["password"] = A.load.Ref(response + 0 + 0, undefined);
        response_ffi["authCredentials"]["username"] = A.load.Ref(response + 0 + 4, undefined);
      }
      if (A.load.Bool(response + 24)) {
        response_ffi["cancel"] = A.load.Bool(response + 9);
      }
      response_ffi["redirectUrl"] = A.load.Ref(response + 12, undefined);
      response_ffi["requestHeaders"] = A.load.Ref(response + 16, undefined);
      response_ffi["responseHeaders"] = A.load.Ref(response + 20, undefined);

      const _ret = WEBEXT.webRequestInternal.eventHandled(
        A.H.get<object>(eventName),
        A.H.get<object>(subEventName),
        A.H.get<object>(requestId),
        webViewInstanceId,
        response_ffi
      );
    },
    "try_EventHandled": (
      retPtr: Pointer,
      errPtr: Pointer,
      eventName: heap.Ref<object>,
      subEventName: heap.Ref<object>,
      requestId: heap.Ref<object>,
      webViewInstanceId: number,
      response: Pointer
    ): heap.Ref<boolean> => {
      try {
        const response_ffi = {};

        if (A.load.Bool(response + 0 + 8)) {
          response_ffi["authCredentials"] = {};
          response_ffi["authCredentials"]["password"] = A.load.Ref(response + 0 + 0, undefined);
          response_ffi["authCredentials"]["username"] = A.load.Ref(response + 0 + 4, undefined);
        }
        if (A.load.Bool(response + 24)) {
          response_ffi["cancel"] = A.load.Bool(response + 9);
        }
        response_ffi["redirectUrl"] = A.load.Ref(response + 12, undefined);
        response_ffi["requestHeaders"] = A.load.Ref(response + 16, undefined);
        response_ffi["responseHeaders"] = A.load.Ref(response + 20, undefined);

        const _ret = WEBEXT.webRequestInternal.eventHandled(
          A.H.get<object>(eventName),
          A.H.get<object>(subEventName),
          A.H.get<object>(requestId),
          webViewInstanceId,
          response_ffi
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
