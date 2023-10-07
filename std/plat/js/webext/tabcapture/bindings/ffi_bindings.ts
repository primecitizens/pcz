import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/tabcapture", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_TabCaptureState": (ref: heap.Ref<string>): number => {
      const idx = ["pending", "active", "stopped", "error"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_CaptureInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Enum(ptr + 4, -1);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 11, true);
        A.store.Bool(ptr + 9, "tabId" in x ? true : false);
        A.store.Int32(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Enum(ptr + 4, ["pending", "active", "stopped", "error"].indexOf(x["status"] as string));
        A.store.Bool(ptr + 10, "fullscreen" in x ? true : false);
        A.store.Bool(ptr + 8, x["fullscreen"] ? true : false);
      }
    },
    "load_CaptureInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 9)) {
        x["tabId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["tabId"];
      }
      x["status"] = A.load.Enum(ptr + 4, ["pending", "active", "stopped", "error"]);
      if (A.load.Bool(ptr + 10)) {
        x["fullscreen"] = A.load.Bool(ptr + 8);
      } else {
        delete x["fullscreen"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MediaStreamConstraint": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["mandatory"]);
        A.store.Ref(ptr + 4, x["optional"]);
      }
    },
    "load_MediaStreamConstraint": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["mandatory"] = A.load.Ref(ptr + 0, undefined);
      x["optional"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CaptureOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 34, false);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 33, false);
        A.store.Bool(ptr + 1, false);

        A.store.Bool(ptr + 4 + 8, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Ref(ptr + 4 + 4, undefined);

        A.store.Bool(ptr + 16 + 8, false);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 4, undefined);
        A.store.Ref(ptr + 28, undefined);
      } else {
        A.store.Bool(ptr + 34, true);
        A.store.Bool(ptr + 32, "audio" in x ? true : false);
        A.store.Bool(ptr + 0, x["audio"] ? true : false);
        A.store.Bool(ptr + 33, "video" in x ? true : false);
        A.store.Bool(ptr + 1, x["video"] ? true : false);

        if (typeof x["audioConstraints"] === "undefined") {
          A.store.Bool(ptr + 4 + 8, false);
          A.store.Ref(ptr + 4 + 0, undefined);
          A.store.Ref(ptr + 4 + 4, undefined);
        } else {
          A.store.Bool(ptr + 4 + 8, true);
          A.store.Ref(ptr + 4 + 0, x["audioConstraints"]["mandatory"]);
          A.store.Ref(ptr + 4 + 4, x["audioConstraints"]["optional"]);
        }

        if (typeof x["videoConstraints"] === "undefined") {
          A.store.Bool(ptr + 16 + 8, false);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 4, undefined);
        } else {
          A.store.Bool(ptr + 16 + 8, true);
          A.store.Ref(ptr + 16 + 0, x["videoConstraints"]["mandatory"]);
          A.store.Ref(ptr + 16 + 4, x["videoConstraints"]["optional"]);
        }
        A.store.Ref(ptr + 28, x["presentationId"]);
      }
    },
    "load_CaptureOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 32)) {
        x["audio"] = A.load.Bool(ptr + 0);
      } else {
        delete x["audio"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["video"] = A.load.Bool(ptr + 1);
      } else {
        delete x["video"];
      }
      if (A.load.Bool(ptr + 4 + 8)) {
        x["audioConstraints"] = {};
        x["audioConstraints"]["mandatory"] = A.load.Ref(ptr + 4 + 0, undefined);
        x["audioConstraints"]["optional"] = A.load.Ref(ptr + 4 + 4, undefined);
      } else {
        delete x["audioConstraints"];
      }
      if (A.load.Bool(ptr + 16 + 8)) {
        x["videoConstraints"] = {};
        x["videoConstraints"]["mandatory"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["videoConstraints"]["optional"] = A.load.Ref(ptr + 16 + 4, undefined);
      } else {
        delete x["videoConstraints"];
      }
      x["presentationId"] = A.load.Ref(ptr + 28, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetMediaStreamOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "consumerTabId" in x ? true : false);
        A.store.Int32(ptr + 0, x["consumerTabId"] === undefined ? 0 : (x["consumerTabId"] as number));
        A.store.Bool(ptr + 9, "targetTabId" in x ? true : false);
        A.store.Int32(ptr + 4, x["targetTabId"] === undefined ? 0 : (x["targetTabId"] as number));
      }
    },
    "load_GetMediaStreamOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["consumerTabId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["consumerTabId"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["targetTabId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["targetTabId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Capture": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabCapture && "capture" in WEBEXT?.tabCapture) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Capture": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabCapture.capture);
    },
    "call_Capture": (retPtr: Pointer, options: Pointer, callback: heap.Ref<object>): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 32)) {
        options_ffi["audio"] = A.load.Bool(options + 0);
      }
      if (A.load.Bool(options + 33)) {
        options_ffi["video"] = A.load.Bool(options + 1);
      }
      if (A.load.Bool(options + 4 + 8)) {
        options_ffi["audioConstraints"] = {};
        options_ffi["audioConstraints"]["mandatory"] = A.load.Ref(options + 4 + 0, undefined);
        options_ffi["audioConstraints"]["optional"] = A.load.Ref(options + 4 + 4, undefined);
      }
      if (A.load.Bool(options + 16 + 8)) {
        options_ffi["videoConstraints"] = {};
        options_ffi["videoConstraints"]["mandatory"] = A.load.Ref(options + 16 + 0, undefined);
        options_ffi["videoConstraints"]["optional"] = A.load.Ref(options + 16 + 4, undefined);
      }
      options_ffi["presentationId"] = A.load.Ref(options + 28, undefined);

      const _ret = WEBEXT.tabCapture.capture(options_ffi, A.H.get<object>(callback));
    },
    "try_Capture": (
      retPtr: Pointer,
      errPtr: Pointer,
      options: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 32)) {
          options_ffi["audio"] = A.load.Bool(options + 0);
        }
        if (A.load.Bool(options + 33)) {
          options_ffi["video"] = A.load.Bool(options + 1);
        }
        if (A.load.Bool(options + 4 + 8)) {
          options_ffi["audioConstraints"] = {};
          options_ffi["audioConstraints"]["mandatory"] = A.load.Ref(options + 4 + 0, undefined);
          options_ffi["audioConstraints"]["optional"] = A.load.Ref(options + 4 + 4, undefined);
        }
        if (A.load.Bool(options + 16 + 8)) {
          options_ffi["videoConstraints"] = {};
          options_ffi["videoConstraints"]["mandatory"] = A.load.Ref(options + 16 + 0, undefined);
          options_ffi["videoConstraints"]["optional"] = A.load.Ref(options + 16 + 4, undefined);
        }
        options_ffi["presentationId"] = A.load.Ref(options + 28, undefined);

        const _ret = WEBEXT.tabCapture.capture(options_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCapturedTabs": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabCapture && "getCapturedTabs" in WEBEXT?.tabCapture) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCapturedTabs": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabCapture.getCapturedTabs);
    },
    "call_GetCapturedTabs": (retPtr: Pointer): void => {
      const _ret = WEBEXT.tabCapture.getCapturedTabs();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCapturedTabs": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabCapture.getCapturedTabs();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetMediaStreamId": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabCapture && "getMediaStreamId" in WEBEXT?.tabCapture) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetMediaStreamId": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabCapture.getMediaStreamId);
    },
    "call_GetMediaStreamId": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 8)) {
        options_ffi["consumerTabId"] = A.load.Int32(options + 0);
      }
      if (A.load.Bool(options + 9)) {
        options_ffi["targetTabId"] = A.load.Int32(options + 4);
      }

      const _ret = WEBEXT.tabCapture.getMediaStreamId(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetMediaStreamId": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 8)) {
          options_ffi["consumerTabId"] = A.load.Int32(options + 0);
        }
        if (A.load.Bool(options + 9)) {
          options_ffi["targetTabId"] = A.load.Int32(options + 4);
        }

        const _ret = WEBEXT.tabCapture.getMediaStreamId(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnStatusChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabCapture?.onStatusChanged && "addListener" in WEBEXT?.tabCapture?.onStatusChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnStatusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabCapture.onStatusChanged.addListener);
    },
    "call_OnStatusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabCapture.onStatusChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnStatusChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabCapture.onStatusChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffStatusChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabCapture?.onStatusChanged && "removeListener" in WEBEXT?.tabCapture?.onStatusChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffStatusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabCapture.onStatusChanged.removeListener);
    },
    "call_OffStatusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabCapture.onStatusChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffStatusChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabCapture.onStatusChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnStatusChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabCapture?.onStatusChanged && "hasListener" in WEBEXT?.tabCapture?.onStatusChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnStatusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabCapture.onStatusChanged.hasListener);
    },
    "call_HasOnStatusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabCapture.onStatusChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnStatusChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabCapture.onStatusChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
