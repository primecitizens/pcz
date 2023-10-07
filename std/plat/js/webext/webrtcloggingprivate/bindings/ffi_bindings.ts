import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/webrtcloggingprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_MetaDataEntry": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["key"]);
        A.store.Ref(ptr + 4, x["value"]);
      }
    },
    "load_MetaDataEntry": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["key"] = A.load.Ref(ptr + 0, undefined);
      x["value"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RecordingInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 5, false);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["prefixPath"]);
        A.store.Bool(ptr + 6, "didStop" in x ? true : false);
        A.store.Bool(ptr + 4, x["didStop"] ? true : false);
        A.store.Bool(ptr + 7, "didManualStop" in x ? true : false);
        A.store.Bool(ptr + 5, x["didManualStop"] ? true : false);
      }
    },
    "load_RecordingInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["prefixPath"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 6)) {
        x["didStop"] = A.load.Bool(ptr + 4);
      } else {
        delete x["didStop"];
      }
      if (A.load.Bool(ptr + 7)) {
        x["didManualStop"] = A.load.Bool(ptr + 5);
      } else {
        delete x["didManualStop"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RequestInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 10, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Bool(ptr + 9, "tabId" in x ? true : false);
        A.store.Int32(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Bool(ptr + 10, "guestProcessId" in x ? true : false);
        A.store.Int32(ptr + 4, x["guestProcessId"] === undefined ? 0 : (x["guestProcessId"] as number));
        A.store.Bool(ptr + 11, "targetWebview" in x ? true : false);
        A.store.Bool(ptr + 8, x["targetWebview"] ? true : false);
      }
    },
    "load_RequestInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 9)) {
        x["tabId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["tabId"];
      }
      if (A.load.Bool(ptr + 10)) {
        x["guestProcessId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["guestProcessId"];
      }
      if (A.load.Bool(ptr + 11)) {
        x["targetWebview"] = A.load.Bool(ptr + 8);
      } else {
        delete x["targetWebview"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_StartEventLoggingResult": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["logId"]);
      }
    },
    "load_StartEventLoggingResult": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["logId"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UploadResult": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["reportId"]);
      }
    },
    "load_UploadResult": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["reportId"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Discard": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcLoggingPrivate && "discard" in WEBEXT?.webrtcLoggingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Discard": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcLoggingPrivate.discard);
    },
    "call_Discard": (
      retPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 9)) {
        request_ffi["tabId"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 10)) {
        request_ffi["guestProcessId"] = A.load.Int32(request + 4);
      }
      if (A.load.Bool(request + 11)) {
        request_ffi["targetWebview"] = A.load.Bool(request + 8);
      }

      const _ret = WEBEXT.webrtcLoggingPrivate.discard(
        request_ffi,
        A.H.get<object>(securityOrigin),
        A.H.get<object>(callback)
      );
    },
    "try_Discard": (
      retPtr: Pointer,
      errPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 9)) {
          request_ffi["tabId"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 10)) {
          request_ffi["guestProcessId"] = A.load.Int32(request + 4);
        }
        if (A.load.Bool(request + 11)) {
          request_ffi["targetWebview"] = A.load.Bool(request + 8);
        }

        const _ret = WEBEXT.webrtcLoggingPrivate.discard(
          request_ffi,
          A.H.get<object>(securityOrigin),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetLogsDirectory": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcLoggingPrivate && "getLogsDirectory" in WEBEXT?.webrtcLoggingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetLogsDirectory": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcLoggingPrivate.getLogsDirectory);
    },
    "call_GetLogsDirectory": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webrtcLoggingPrivate.getLogsDirectory(A.H.get<object>(callback));
    },
    "try_GetLogsDirectory": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webrtcLoggingPrivate.getLogsDirectory(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetMetaData": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcLoggingPrivate && "setMetaData" in WEBEXT?.webrtcLoggingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetMetaData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcLoggingPrivate.setMetaData);
    },
    "call_SetMetaData": (
      retPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      metaData: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 9)) {
        request_ffi["tabId"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 10)) {
        request_ffi["guestProcessId"] = A.load.Int32(request + 4);
      }
      if (A.load.Bool(request + 11)) {
        request_ffi["targetWebview"] = A.load.Bool(request + 8);
      }

      const _ret = WEBEXT.webrtcLoggingPrivate.setMetaData(
        request_ffi,
        A.H.get<object>(securityOrigin),
        A.H.get<object>(metaData),
        A.H.get<object>(callback)
      );
    },
    "try_SetMetaData": (
      retPtr: Pointer,
      errPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      metaData: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 9)) {
          request_ffi["tabId"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 10)) {
          request_ffi["guestProcessId"] = A.load.Int32(request + 4);
        }
        if (A.load.Bool(request + 11)) {
          request_ffi["targetWebview"] = A.load.Bool(request + 8);
        }

        const _ret = WEBEXT.webrtcLoggingPrivate.setMetaData(
          request_ffi,
          A.H.get<object>(securityOrigin),
          A.H.get<object>(metaData),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetUploadOnRenderClose": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcLoggingPrivate && "setUploadOnRenderClose" in WEBEXT?.webrtcLoggingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetUploadOnRenderClose": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcLoggingPrivate.setUploadOnRenderClose);
    },
    "call_SetUploadOnRenderClose": (
      retPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      shouldUpload: heap.Ref<boolean>
    ): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 9)) {
        request_ffi["tabId"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 10)) {
        request_ffi["guestProcessId"] = A.load.Int32(request + 4);
      }
      if (A.load.Bool(request + 11)) {
        request_ffi["targetWebview"] = A.load.Bool(request + 8);
      }

      const _ret = WEBEXT.webrtcLoggingPrivate.setUploadOnRenderClose(
        request_ffi,
        A.H.get<object>(securityOrigin),
        shouldUpload === A.H.TRUE
      );
    },
    "try_SetUploadOnRenderClose": (
      retPtr: Pointer,
      errPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      shouldUpload: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 9)) {
          request_ffi["tabId"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 10)) {
          request_ffi["guestProcessId"] = A.load.Int32(request + 4);
        }
        if (A.load.Bool(request + 11)) {
          request_ffi["targetWebview"] = A.load.Bool(request + 8);
        }

        const _ret = WEBEXT.webrtcLoggingPrivate.setUploadOnRenderClose(
          request_ffi,
          A.H.get<object>(securityOrigin),
          shouldUpload === A.H.TRUE
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Start": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcLoggingPrivate && "start" in WEBEXT?.webrtcLoggingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Start": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcLoggingPrivate.start);
    },
    "call_Start": (
      retPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 9)) {
        request_ffi["tabId"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 10)) {
        request_ffi["guestProcessId"] = A.load.Int32(request + 4);
      }
      if (A.load.Bool(request + 11)) {
        request_ffi["targetWebview"] = A.load.Bool(request + 8);
      }

      const _ret = WEBEXT.webrtcLoggingPrivate.start(
        request_ffi,
        A.H.get<object>(securityOrigin),
        A.H.get<object>(callback)
      );
    },
    "try_Start": (
      retPtr: Pointer,
      errPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 9)) {
          request_ffi["tabId"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 10)) {
          request_ffi["guestProcessId"] = A.load.Int32(request + 4);
        }
        if (A.load.Bool(request + 11)) {
          request_ffi["targetWebview"] = A.load.Bool(request + 8);
        }

        const _ret = WEBEXT.webrtcLoggingPrivate.start(
          request_ffi,
          A.H.get<object>(securityOrigin),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartAudioDebugRecordings": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcLoggingPrivate && "startAudioDebugRecordings" in WEBEXT?.webrtcLoggingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartAudioDebugRecordings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcLoggingPrivate.startAudioDebugRecordings);
    },
    "call_StartAudioDebugRecordings": (
      retPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      seconds: number,
      callback: heap.Ref<object>
    ): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 9)) {
        request_ffi["tabId"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 10)) {
        request_ffi["guestProcessId"] = A.load.Int32(request + 4);
      }
      if (A.load.Bool(request + 11)) {
        request_ffi["targetWebview"] = A.load.Bool(request + 8);
      }

      const _ret = WEBEXT.webrtcLoggingPrivate.startAudioDebugRecordings(
        request_ffi,
        A.H.get<object>(securityOrigin),
        seconds,
        A.H.get<object>(callback)
      );
    },
    "try_StartAudioDebugRecordings": (
      retPtr: Pointer,
      errPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      seconds: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 9)) {
          request_ffi["tabId"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 10)) {
          request_ffi["guestProcessId"] = A.load.Int32(request + 4);
        }
        if (A.load.Bool(request + 11)) {
          request_ffi["targetWebview"] = A.load.Bool(request + 8);
        }

        const _ret = WEBEXT.webrtcLoggingPrivate.startAudioDebugRecordings(
          request_ffi,
          A.H.get<object>(securityOrigin),
          seconds,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartEventLogging": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcLoggingPrivate && "startEventLogging" in WEBEXT?.webrtcLoggingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartEventLogging": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcLoggingPrivate.startEventLogging);
    },
    "call_StartEventLogging": (
      retPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      sessionId: heap.Ref<object>,
      maxLogSizeBytes: number,
      outputPeriodMs: number,
      webAppId: number,
      callback: heap.Ref<object>
    ): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 9)) {
        request_ffi["tabId"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 10)) {
        request_ffi["guestProcessId"] = A.load.Int32(request + 4);
      }
      if (A.load.Bool(request + 11)) {
        request_ffi["targetWebview"] = A.load.Bool(request + 8);
      }

      const _ret = WEBEXT.webrtcLoggingPrivate.startEventLogging(
        request_ffi,
        A.H.get<object>(securityOrigin),
        A.H.get<object>(sessionId),
        maxLogSizeBytes,
        outputPeriodMs,
        webAppId,
        A.H.get<object>(callback)
      );
    },
    "try_StartEventLogging": (
      retPtr: Pointer,
      errPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      sessionId: heap.Ref<object>,
      maxLogSizeBytes: number,
      outputPeriodMs: number,
      webAppId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 9)) {
          request_ffi["tabId"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 10)) {
          request_ffi["guestProcessId"] = A.load.Int32(request + 4);
        }
        if (A.load.Bool(request + 11)) {
          request_ffi["targetWebview"] = A.load.Bool(request + 8);
        }

        const _ret = WEBEXT.webrtcLoggingPrivate.startEventLogging(
          request_ffi,
          A.H.get<object>(securityOrigin),
          A.H.get<object>(sessionId),
          maxLogSizeBytes,
          outputPeriodMs,
          webAppId,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartRtpDump": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcLoggingPrivate && "startRtpDump" in WEBEXT?.webrtcLoggingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartRtpDump": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcLoggingPrivate.startRtpDump);
    },
    "call_StartRtpDump": (
      retPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      incoming: heap.Ref<boolean>,
      outgoing: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 9)) {
        request_ffi["tabId"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 10)) {
        request_ffi["guestProcessId"] = A.load.Int32(request + 4);
      }
      if (A.load.Bool(request + 11)) {
        request_ffi["targetWebview"] = A.load.Bool(request + 8);
      }

      const _ret = WEBEXT.webrtcLoggingPrivate.startRtpDump(
        request_ffi,
        A.H.get<object>(securityOrigin),
        incoming === A.H.TRUE,
        outgoing === A.H.TRUE,
        A.H.get<object>(callback)
      );
    },
    "try_StartRtpDump": (
      retPtr: Pointer,
      errPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      incoming: heap.Ref<boolean>,
      outgoing: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 9)) {
          request_ffi["tabId"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 10)) {
          request_ffi["guestProcessId"] = A.load.Int32(request + 4);
        }
        if (A.load.Bool(request + 11)) {
          request_ffi["targetWebview"] = A.load.Bool(request + 8);
        }

        const _ret = WEBEXT.webrtcLoggingPrivate.startRtpDump(
          request_ffi,
          A.H.get<object>(securityOrigin),
          incoming === A.H.TRUE,
          outgoing === A.H.TRUE,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Stop": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcLoggingPrivate && "stop" in WEBEXT?.webrtcLoggingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Stop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcLoggingPrivate.stop);
    },
    "call_Stop": (
      retPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 9)) {
        request_ffi["tabId"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 10)) {
        request_ffi["guestProcessId"] = A.load.Int32(request + 4);
      }
      if (A.load.Bool(request + 11)) {
        request_ffi["targetWebview"] = A.load.Bool(request + 8);
      }

      const _ret = WEBEXT.webrtcLoggingPrivate.stop(
        request_ffi,
        A.H.get<object>(securityOrigin),
        A.H.get<object>(callback)
      );
    },
    "try_Stop": (
      retPtr: Pointer,
      errPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 9)) {
          request_ffi["tabId"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 10)) {
          request_ffi["guestProcessId"] = A.load.Int32(request + 4);
        }
        if (A.load.Bool(request + 11)) {
          request_ffi["targetWebview"] = A.load.Bool(request + 8);
        }

        const _ret = WEBEXT.webrtcLoggingPrivate.stop(
          request_ffi,
          A.H.get<object>(securityOrigin),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StopAudioDebugRecordings": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcLoggingPrivate && "stopAudioDebugRecordings" in WEBEXT?.webrtcLoggingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StopAudioDebugRecordings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcLoggingPrivate.stopAudioDebugRecordings);
    },
    "call_StopAudioDebugRecordings": (
      retPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 9)) {
        request_ffi["tabId"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 10)) {
        request_ffi["guestProcessId"] = A.load.Int32(request + 4);
      }
      if (A.load.Bool(request + 11)) {
        request_ffi["targetWebview"] = A.load.Bool(request + 8);
      }

      const _ret = WEBEXT.webrtcLoggingPrivate.stopAudioDebugRecordings(
        request_ffi,
        A.H.get<object>(securityOrigin),
        A.H.get<object>(callback)
      );
    },
    "try_StopAudioDebugRecordings": (
      retPtr: Pointer,
      errPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 9)) {
          request_ffi["tabId"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 10)) {
          request_ffi["guestProcessId"] = A.load.Int32(request + 4);
        }
        if (A.load.Bool(request + 11)) {
          request_ffi["targetWebview"] = A.load.Bool(request + 8);
        }

        const _ret = WEBEXT.webrtcLoggingPrivate.stopAudioDebugRecordings(
          request_ffi,
          A.H.get<object>(securityOrigin),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StopRtpDump": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcLoggingPrivate && "stopRtpDump" in WEBEXT?.webrtcLoggingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StopRtpDump": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcLoggingPrivate.stopRtpDump);
    },
    "call_StopRtpDump": (
      retPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      incoming: heap.Ref<boolean>,
      outgoing: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 9)) {
        request_ffi["tabId"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 10)) {
        request_ffi["guestProcessId"] = A.load.Int32(request + 4);
      }
      if (A.load.Bool(request + 11)) {
        request_ffi["targetWebview"] = A.load.Bool(request + 8);
      }

      const _ret = WEBEXT.webrtcLoggingPrivate.stopRtpDump(
        request_ffi,
        A.H.get<object>(securityOrigin),
        incoming === A.H.TRUE,
        outgoing === A.H.TRUE,
        A.H.get<object>(callback)
      );
    },
    "try_StopRtpDump": (
      retPtr: Pointer,
      errPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      incoming: heap.Ref<boolean>,
      outgoing: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 9)) {
          request_ffi["tabId"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 10)) {
          request_ffi["guestProcessId"] = A.load.Int32(request + 4);
        }
        if (A.load.Bool(request + 11)) {
          request_ffi["targetWebview"] = A.load.Bool(request + 8);
        }

        const _ret = WEBEXT.webrtcLoggingPrivate.stopRtpDump(
          request_ffi,
          A.H.get<object>(securityOrigin),
          incoming === A.H.TRUE,
          outgoing === A.H.TRUE,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Store": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcLoggingPrivate && "store" in WEBEXT?.webrtcLoggingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Store": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcLoggingPrivate.store);
    },
    "call_Store": (
      retPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      logId: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 9)) {
        request_ffi["tabId"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 10)) {
        request_ffi["guestProcessId"] = A.load.Int32(request + 4);
      }
      if (A.load.Bool(request + 11)) {
        request_ffi["targetWebview"] = A.load.Bool(request + 8);
      }

      const _ret = WEBEXT.webrtcLoggingPrivate.store(
        request_ffi,
        A.H.get<object>(securityOrigin),
        A.H.get<object>(logId),
        A.H.get<object>(callback)
      );
    },
    "try_Store": (
      retPtr: Pointer,
      errPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      logId: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 9)) {
          request_ffi["tabId"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 10)) {
          request_ffi["guestProcessId"] = A.load.Int32(request + 4);
        }
        if (A.load.Bool(request + 11)) {
          request_ffi["targetWebview"] = A.load.Bool(request + 8);
        }

        const _ret = WEBEXT.webrtcLoggingPrivate.store(
          request_ffi,
          A.H.get<object>(securityOrigin),
          A.H.get<object>(logId),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Upload": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcLoggingPrivate && "upload" in WEBEXT?.webrtcLoggingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Upload": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcLoggingPrivate.upload);
    },
    "call_Upload": (
      retPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 9)) {
        request_ffi["tabId"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 10)) {
        request_ffi["guestProcessId"] = A.load.Int32(request + 4);
      }
      if (A.load.Bool(request + 11)) {
        request_ffi["targetWebview"] = A.load.Bool(request + 8);
      }

      const _ret = WEBEXT.webrtcLoggingPrivate.upload(
        request_ffi,
        A.H.get<object>(securityOrigin),
        A.H.get<object>(callback)
      );
    },
    "try_Upload": (
      retPtr: Pointer,
      errPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 9)) {
          request_ffi["tabId"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 10)) {
          request_ffi["guestProcessId"] = A.load.Int32(request + 4);
        }
        if (A.load.Bool(request + 11)) {
          request_ffi["targetWebview"] = A.load.Bool(request + 8);
        }

        const _ret = WEBEXT.webrtcLoggingPrivate.upload(
          request_ffi,
          A.H.get<object>(securityOrigin),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UploadStored": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcLoggingPrivate && "uploadStored" in WEBEXT?.webrtcLoggingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UploadStored": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcLoggingPrivate.uploadStored);
    },
    "call_UploadStored": (
      retPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      logId: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 9)) {
        request_ffi["tabId"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 10)) {
        request_ffi["guestProcessId"] = A.load.Int32(request + 4);
      }
      if (A.load.Bool(request + 11)) {
        request_ffi["targetWebview"] = A.load.Bool(request + 8);
      }

      const _ret = WEBEXT.webrtcLoggingPrivate.uploadStored(
        request_ffi,
        A.H.get<object>(securityOrigin),
        A.H.get<object>(logId),
        A.H.get<object>(callback)
      );
    },
    "try_UploadStored": (
      retPtr: Pointer,
      errPtr: Pointer,
      request: Pointer,
      securityOrigin: heap.Ref<object>,
      logId: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 9)) {
          request_ffi["tabId"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 10)) {
          request_ffi["guestProcessId"] = A.load.Int32(request + 4);
        }
        if (A.load.Bool(request + 11)) {
          request_ffi["targetWebview"] = A.load.Bool(request + 8);
        }

        const _ret = WEBEXT.webrtcLoggingPrivate.uploadStored(
          request_ffi,
          A.H.get<object>(securityOrigin),
          A.H.get<object>(logId),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
