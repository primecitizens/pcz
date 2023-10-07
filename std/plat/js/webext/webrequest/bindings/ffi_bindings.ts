import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/webrequest", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_BlockingResponseFieldAuthCredentials": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["password"]);
        A.store.Ref(ptr + 4, x["username"]);
      }
    },
    "load_BlockingResponseFieldAuthCredentials": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["password"] = A.load.Ref(ptr + 0, undefined);
      x["username"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_HttpHeadersElem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["binaryValue"]);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Ref(ptr + 8, x["value"]);
      }
    },
    "load_HttpHeadersElem": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["binaryValue"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      x["value"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_BlockingResponse": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);

        A.store.Bool(ptr + 0 + 8, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
      } else {
        A.store.Bool(ptr + 25, true);

        if (typeof x["authCredentials"] === "undefined") {
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
        } else {
          A.store.Bool(ptr + 0 + 8, true);
          A.store.Ref(ptr + 0 + 0, x["authCredentials"]["password"]);
          A.store.Ref(ptr + 0 + 4, x["authCredentials"]["username"]);
        }
        A.store.Bool(ptr + 24, "cancel" in x ? true : false);
        A.store.Bool(ptr + 9, x["cancel"] ? true : false);
        A.store.Ref(ptr + 12, x["redirectUrl"]);
        A.store.Ref(ptr + 16, x["requestHeaders"]);
        A.store.Ref(ptr + 20, x["responseHeaders"]);
      }
    },
    "load_BlockingResponse": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 8)) {
        x["authCredentials"] = {};
        x["authCredentials"]["password"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["authCredentials"]["username"] = A.load.Ref(ptr + 0 + 4, undefined);
      } else {
        delete x["authCredentials"];
      }
      if (A.load.Bool(ptr + 24)) {
        x["cancel"] = A.load.Bool(ptr + 9);
      } else {
        delete x["cancel"];
      }
      x["redirectUrl"] = A.load.Ref(ptr + 12, undefined);
      x["requestHeaders"] = A.load.Ref(ptr + 16, undefined);
      x["responseHeaders"] = A.load.Ref(ptr + 20, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_IgnoredActionType": (ref: heap.Ref<string>): number => {
      const idx = ["redirect", "request_headers", "response_headers", "auth_credentials"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "get_MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest && "MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES" in WEBEXT?.webRequest) {
        const val = WEBEXT.webRequest.MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.webRequest,
        "MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES",
        A.H.get<object>(val),
        WEBEXT.webRequest
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },

    "store_OnActionIgnoredArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(
          ptr + 0,
          ["redirect", "request_headers", "response_headers", "auth_credentials"].indexOf(x["action"] as string)
        );
        A.store.Ref(ptr + 4, x["requestId"]);
      }
    },
    "load_OnActionIgnoredArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["action"] = A.load.Enum(ptr + 0, ["redirect", "request_headers", "response_headers", "auth_credentials"]);
      x["requestId"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnAuthRequiredArgDetailsFieldChallenger": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["host"]);
        A.store.Int64(ptr + 8, x["port"] === undefined ? 0 : (x["port"] as number));
      }
    },
    "load_OnAuthRequiredArgDetailsFieldChallenger": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["host"] = A.load.Ref(ptr + 0, undefined);
      x["port"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ResourceType": (ref: heap.Ref<string>): number => {
      const idx = [
        "main_frame",
        "sub_frame",
        "stylesheet",
        "script",
        "image",
        "font",
        "object",
        "xmlhttprequest",
        "ping",
        "csp_report",
        "media",
        "websocket",
        "webbundle",
        "other",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OnAuthRequiredArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 128, false);

        A.store.Bool(ptr + 0 + 16, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Int64(ptr + 0 + 8, 0);
        A.store.Ref(ptr + 20, undefined);
        A.store.Enum(ptr + 24, -1);
        A.store.Int64(ptr + 32, 0);
        A.store.Enum(ptr + 40, -1);
        A.store.Ref(ptr + 44, undefined);
        A.store.Bool(ptr + 48, false);
        A.store.Ref(ptr + 52, undefined);
        A.store.Ref(ptr + 56, undefined);
        A.store.Int64(ptr + 64, 0);
        A.store.Ref(ptr + 72, undefined);
        A.store.Ref(ptr + 76, undefined);
        A.store.Ref(ptr + 80, undefined);
        A.store.Ref(ptr + 84, undefined);
        A.store.Int64(ptr + 88, 0);
        A.store.Ref(ptr + 96, undefined);
        A.store.Int64(ptr + 104, 0);
        A.store.Float64(ptr + 112, 0);
        A.store.Enum(ptr + 120, -1);
        A.store.Ref(ptr + 124, undefined);
      } else {
        A.store.Bool(ptr + 128, true);

        if (typeof x["challenger"] === "undefined") {
          A.store.Bool(ptr + 0 + 16, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Int64(ptr + 0 + 8, 0);
        } else {
          A.store.Bool(ptr + 0 + 16, true);
          A.store.Ref(ptr + 0 + 0, x["challenger"]["host"]);
          A.store.Int64(ptr + 0 + 8, x["challenger"]["port"] === undefined ? 0 : (x["challenger"]["port"] as number));
        }
        A.store.Ref(ptr + 20, x["documentId"]);
        A.store.Enum(
          ptr + 24,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Int64(ptr + 32, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 40, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Ref(ptr + 44, x["initiator"]);
        A.store.Bool(ptr + 48, x["isProxy"] ? true : false);
        A.store.Ref(ptr + 52, x["method"]);
        A.store.Ref(ptr + 56, x["parentDocumentId"]);
        A.store.Int64(ptr + 64, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Ref(ptr + 72, x["realm"]);
        A.store.Ref(ptr + 76, x["requestId"]);
        A.store.Ref(ptr + 80, x["responseHeaders"]);
        A.store.Ref(ptr + 84, x["scheme"]);
        A.store.Int64(ptr + 88, x["statusCode"] === undefined ? 0 : (x["statusCode"] as number));
        A.store.Ref(ptr + 96, x["statusLine"]);
        A.store.Int64(ptr + 104, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 112, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Enum(
          ptr + 120,
          [
            "main_frame",
            "sub_frame",
            "stylesheet",
            "script",
            "image",
            "font",
            "object",
            "xmlhttprequest",
            "ping",
            "csp_report",
            "media",
            "websocket",
            "webbundle",
            "other",
          ].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 124, x["url"]);
      }
    },
    "load_OnAuthRequiredArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 16)) {
        x["challenger"] = {};
        x["challenger"]["host"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["challenger"]["port"] = A.load.Int64(ptr + 0 + 8);
      } else {
        delete x["challenger"];
      }
      x["documentId"] = A.load.Ref(ptr + 20, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 24, ["prerender", "active", "cached", "pending_deletion"]);
      x["frameId"] = A.load.Int64(ptr + 32);
      x["frameType"] = A.load.Enum(ptr + 40, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["initiator"] = A.load.Ref(ptr + 44, undefined);
      x["isProxy"] = A.load.Bool(ptr + 48);
      x["method"] = A.load.Ref(ptr + 52, undefined);
      x["parentDocumentId"] = A.load.Ref(ptr + 56, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 64);
      x["realm"] = A.load.Ref(ptr + 72, undefined);
      x["requestId"] = A.load.Ref(ptr + 76, undefined);
      x["responseHeaders"] = A.load.Ref(ptr + 80, undefined);
      x["scheme"] = A.load.Ref(ptr + 84, undefined);
      x["statusCode"] = A.load.Int64(ptr + 88);
      x["statusLine"] = A.load.Ref(ptr + 96, undefined);
      x["tabId"] = A.load.Int64(ptr + 104);
      x["timeStamp"] = A.load.Float64(ptr + 112);
      x["type"] = A.load.Enum(ptr + 120, [
        "main_frame",
        "sub_frame",
        "stylesheet",
        "script",
        "image",
        "font",
        "object",
        "xmlhttprequest",
        "ping",
        "csp_report",
        "media",
        "websocket",
        "webbundle",
        "other",
      ]);
      x["url"] = A.load.Ref(ptr + 124, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OnAuthRequiredOptions": (ref: heap.Ref<string>): number => {
      const idx = ["responseHeaders", "blocking", "asyncBlocking", "extraHeaders"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OnBeforeRedirectArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 104, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Bool(ptr + 20, false);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Int64(ptr + 40, 0);
        A.store.Ref(ptr + 48, undefined);
        A.store.Ref(ptr + 52, undefined);
        A.store.Ref(ptr + 56, undefined);
        A.store.Int64(ptr + 64, 0);
        A.store.Ref(ptr + 72, undefined);
        A.store.Int64(ptr + 80, 0);
        A.store.Float64(ptr + 88, 0);
        A.store.Enum(ptr + 96, -1);
        A.store.Ref(ptr + 100, undefined);
      } else {
        A.store.Bool(ptr + 104, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Bool(ptr + 20, x["fromCache"] ? true : false);
        A.store.Ref(ptr + 24, x["initiator"]);
        A.store.Ref(ptr + 28, x["ip"]);
        A.store.Ref(ptr + 32, x["method"]);
        A.store.Ref(ptr + 36, x["parentDocumentId"]);
        A.store.Int64(ptr + 40, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Ref(ptr + 48, x["redirectUrl"]);
        A.store.Ref(ptr + 52, x["requestId"]);
        A.store.Ref(ptr + 56, x["responseHeaders"]);
        A.store.Int64(ptr + 64, x["statusCode"] === undefined ? 0 : (x["statusCode"] as number));
        A.store.Ref(ptr + 72, x["statusLine"]);
        A.store.Int64(ptr + 80, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 88, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Enum(
          ptr + 96,
          [
            "main_frame",
            "sub_frame",
            "stylesheet",
            "script",
            "image",
            "font",
            "object",
            "xmlhttprequest",
            "ping",
            "csp_report",
            "media",
            "websocket",
            "webbundle",
            "other",
          ].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 100, x["url"]);
      }
    },
    "load_OnBeforeRedirectArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["frameId"] = A.load.Int64(ptr + 8);
      x["frameType"] = A.load.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["fromCache"] = A.load.Bool(ptr + 20);
      x["initiator"] = A.load.Ref(ptr + 24, undefined);
      x["ip"] = A.load.Ref(ptr + 28, undefined);
      x["method"] = A.load.Ref(ptr + 32, undefined);
      x["parentDocumentId"] = A.load.Ref(ptr + 36, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 40);
      x["redirectUrl"] = A.load.Ref(ptr + 48, undefined);
      x["requestId"] = A.load.Ref(ptr + 52, undefined);
      x["responseHeaders"] = A.load.Ref(ptr + 56, undefined);
      x["statusCode"] = A.load.Int64(ptr + 64);
      x["statusLine"] = A.load.Ref(ptr + 72, undefined);
      x["tabId"] = A.load.Int64(ptr + 80);
      x["timeStamp"] = A.load.Float64(ptr + 88);
      x["type"] = A.load.Enum(ptr + 96, [
        "main_frame",
        "sub_frame",
        "stylesheet",
        "script",
        "image",
        "font",
        "object",
        "xmlhttprequest",
        "ping",
        "csp_report",
        "media",
        "websocket",
        "webbundle",
        "other",
      ]);
      x["url"] = A.load.Ref(ptr + 100, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OnBeforeRedirectOptions": (ref: heap.Ref<string>): number => {
      const idx = ["responseHeaders", "extraHeaders"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_UploadData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["bytes"]);
        A.store.Ref(ptr + 4, x["file"]);
      }
    },
    "load_UploadData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["bytes"] = A.load.Ref(ptr + 0, undefined);
      x["file"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnBeforeRequestArgDetailsFieldRequestBody": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["error"]);
        A.store.Ref(ptr + 4, x["formData"]);
        A.store.Ref(ptr + 8, x["raw"]);
      }
    },
    "load_OnBeforeRequestArgDetailsFieldRequestBody": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["error"] = A.load.Ref(ptr + 0, undefined);
      x["formData"] = A.load.Ref(ptr + 4, undefined);
      x["raw"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnBeforeRequestArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 88, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Int64(ptr + 32, 0);

        A.store.Bool(ptr + 40 + 12, false);
        A.store.Ref(ptr + 40 + 0, undefined);
        A.store.Ref(ptr + 40 + 4, undefined);
        A.store.Ref(ptr + 40 + 8, undefined);
        A.store.Ref(ptr + 56, undefined);
        A.store.Int64(ptr + 64, 0);
        A.store.Float64(ptr + 72, 0);
        A.store.Enum(ptr + 80, -1);
        A.store.Ref(ptr + 84, undefined);
      } else {
        A.store.Bool(ptr + 88, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Ref(ptr + 20, x["initiator"]);
        A.store.Ref(ptr + 24, x["method"]);
        A.store.Ref(ptr + 28, x["parentDocumentId"]);
        A.store.Int64(ptr + 32, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));

        if (typeof x["requestBody"] === "undefined") {
          A.store.Bool(ptr + 40 + 12, false);
          A.store.Ref(ptr + 40 + 0, undefined);
          A.store.Ref(ptr + 40 + 4, undefined);
          A.store.Ref(ptr + 40 + 8, undefined);
        } else {
          A.store.Bool(ptr + 40 + 12, true);
          A.store.Ref(ptr + 40 + 0, x["requestBody"]["error"]);
          A.store.Ref(ptr + 40 + 4, x["requestBody"]["formData"]);
          A.store.Ref(ptr + 40 + 8, x["requestBody"]["raw"]);
        }
        A.store.Ref(ptr + 56, x["requestId"]);
        A.store.Int64(ptr + 64, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 72, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Enum(
          ptr + 80,
          [
            "main_frame",
            "sub_frame",
            "stylesheet",
            "script",
            "image",
            "font",
            "object",
            "xmlhttprequest",
            "ping",
            "csp_report",
            "media",
            "websocket",
            "webbundle",
            "other",
          ].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 84, x["url"]);
      }
    },
    "load_OnBeforeRequestArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["frameId"] = A.load.Int64(ptr + 8);
      x["frameType"] = A.load.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["initiator"] = A.load.Ref(ptr + 20, undefined);
      x["method"] = A.load.Ref(ptr + 24, undefined);
      x["parentDocumentId"] = A.load.Ref(ptr + 28, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 32);
      if (A.load.Bool(ptr + 40 + 12)) {
        x["requestBody"] = {};
        x["requestBody"]["error"] = A.load.Ref(ptr + 40 + 0, undefined);
        x["requestBody"]["formData"] = A.load.Ref(ptr + 40 + 4, undefined);
        x["requestBody"]["raw"] = A.load.Ref(ptr + 40 + 8, undefined);
      } else {
        delete x["requestBody"];
      }
      x["requestId"] = A.load.Ref(ptr + 56, undefined);
      x["tabId"] = A.load.Int64(ptr + 64);
      x["timeStamp"] = A.load.Float64(ptr + 72);
      x["type"] = A.load.Enum(ptr + 80, [
        "main_frame",
        "sub_frame",
        "stylesheet",
        "script",
        "image",
        "font",
        "object",
        "xmlhttprequest",
        "ping",
        "csp_report",
        "media",
        "websocket",
        "webbundle",
        "other",
      ]);
      x["url"] = A.load.Ref(ptr + 84, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OnBeforeRequestOptions": (ref: heap.Ref<string>): number => {
      const idx = ["blocking", "requestBody", "extraHeaders"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OnBeforeSendHeadersArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 72, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Int64(ptr + 32, 0);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Int64(ptr + 48, 0);
        A.store.Float64(ptr + 56, 0);
        A.store.Enum(ptr + 64, -1);
        A.store.Ref(ptr + 68, undefined);
      } else {
        A.store.Bool(ptr + 72, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Ref(ptr + 20, x["initiator"]);
        A.store.Ref(ptr + 24, x["method"]);
        A.store.Ref(ptr + 28, x["parentDocumentId"]);
        A.store.Int64(ptr + 32, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Ref(ptr + 40, x["requestHeaders"]);
        A.store.Ref(ptr + 44, x["requestId"]);
        A.store.Int64(ptr + 48, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 56, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Enum(
          ptr + 64,
          [
            "main_frame",
            "sub_frame",
            "stylesheet",
            "script",
            "image",
            "font",
            "object",
            "xmlhttprequest",
            "ping",
            "csp_report",
            "media",
            "websocket",
            "webbundle",
            "other",
          ].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 68, x["url"]);
      }
    },
    "load_OnBeforeSendHeadersArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["frameId"] = A.load.Int64(ptr + 8);
      x["frameType"] = A.load.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["initiator"] = A.load.Ref(ptr + 20, undefined);
      x["method"] = A.load.Ref(ptr + 24, undefined);
      x["parentDocumentId"] = A.load.Ref(ptr + 28, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 32);
      x["requestHeaders"] = A.load.Ref(ptr + 40, undefined);
      x["requestId"] = A.load.Ref(ptr + 44, undefined);
      x["tabId"] = A.load.Int64(ptr + 48);
      x["timeStamp"] = A.load.Float64(ptr + 56);
      x["type"] = A.load.Enum(ptr + 64, [
        "main_frame",
        "sub_frame",
        "stylesheet",
        "script",
        "image",
        "font",
        "object",
        "xmlhttprequest",
        "ping",
        "csp_report",
        "media",
        "websocket",
        "webbundle",
        "other",
      ]);
      x["url"] = A.load.Ref(ptr + 68, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OnBeforeSendHeadersOptions": (ref: heap.Ref<string>): number => {
      const idx = ["requestHeaders", "blocking", "extraHeaders"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OnCompletedArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 96, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Bool(ptr + 20, false);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Int64(ptr + 40, 0);
        A.store.Ref(ptr + 48, undefined);
        A.store.Ref(ptr + 52, undefined);
        A.store.Int64(ptr + 56, 0);
        A.store.Ref(ptr + 64, undefined);
        A.store.Int64(ptr + 72, 0);
        A.store.Float64(ptr + 80, 0);
        A.store.Enum(ptr + 88, -1);
        A.store.Ref(ptr + 92, undefined);
      } else {
        A.store.Bool(ptr + 96, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Bool(ptr + 20, x["fromCache"] ? true : false);
        A.store.Ref(ptr + 24, x["initiator"]);
        A.store.Ref(ptr + 28, x["ip"]);
        A.store.Ref(ptr + 32, x["method"]);
        A.store.Ref(ptr + 36, x["parentDocumentId"]);
        A.store.Int64(ptr + 40, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Ref(ptr + 48, x["requestId"]);
        A.store.Ref(ptr + 52, x["responseHeaders"]);
        A.store.Int64(ptr + 56, x["statusCode"] === undefined ? 0 : (x["statusCode"] as number));
        A.store.Ref(ptr + 64, x["statusLine"]);
        A.store.Int64(ptr + 72, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 80, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Enum(
          ptr + 88,
          [
            "main_frame",
            "sub_frame",
            "stylesheet",
            "script",
            "image",
            "font",
            "object",
            "xmlhttprequest",
            "ping",
            "csp_report",
            "media",
            "websocket",
            "webbundle",
            "other",
          ].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 92, x["url"]);
      }
    },
    "load_OnCompletedArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["frameId"] = A.load.Int64(ptr + 8);
      x["frameType"] = A.load.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["fromCache"] = A.load.Bool(ptr + 20);
      x["initiator"] = A.load.Ref(ptr + 24, undefined);
      x["ip"] = A.load.Ref(ptr + 28, undefined);
      x["method"] = A.load.Ref(ptr + 32, undefined);
      x["parentDocumentId"] = A.load.Ref(ptr + 36, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 40);
      x["requestId"] = A.load.Ref(ptr + 48, undefined);
      x["responseHeaders"] = A.load.Ref(ptr + 52, undefined);
      x["statusCode"] = A.load.Int64(ptr + 56);
      x["statusLine"] = A.load.Ref(ptr + 64, undefined);
      x["tabId"] = A.load.Int64(ptr + 72);
      x["timeStamp"] = A.load.Float64(ptr + 80);
      x["type"] = A.load.Enum(ptr + 88, [
        "main_frame",
        "sub_frame",
        "stylesheet",
        "script",
        "image",
        "font",
        "object",
        "xmlhttprequest",
        "ping",
        "csp_report",
        "media",
        "websocket",
        "webbundle",
        "other",
      ]);
      x["url"] = A.load.Ref(ptr + 92, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OnCompletedOptions": (ref: heap.Ref<string>): number => {
      const idx = ["responseHeaders", "extraHeaders"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OnErrorOccurredArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 88, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
        A.store.Int64(ptr + 16, 0);
        A.store.Enum(ptr + 24, -1);
        A.store.Bool(ptr + 28, false);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Int64(ptr + 48, 0);
        A.store.Ref(ptr + 56, undefined);
        A.store.Int64(ptr + 64, 0);
        A.store.Float64(ptr + 72, 0);
        A.store.Enum(ptr + 80, -1);
        A.store.Ref(ptr + 84, undefined);
      } else {
        A.store.Bool(ptr + 88, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Ref(ptr + 8, x["error"]);
        A.store.Int64(ptr + 16, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 24, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Bool(ptr + 28, x["fromCache"] ? true : false);
        A.store.Ref(ptr + 32, x["initiator"]);
        A.store.Ref(ptr + 36, x["ip"]);
        A.store.Ref(ptr + 40, x["method"]);
        A.store.Ref(ptr + 44, x["parentDocumentId"]);
        A.store.Int64(ptr + 48, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Ref(ptr + 56, x["requestId"]);
        A.store.Int64(ptr + 64, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 72, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Enum(
          ptr + 80,
          [
            "main_frame",
            "sub_frame",
            "stylesheet",
            "script",
            "image",
            "font",
            "object",
            "xmlhttprequest",
            "ping",
            "csp_report",
            "media",
            "websocket",
            "webbundle",
            "other",
          ].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 84, x["url"]);
      }
    },
    "load_OnErrorOccurredArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["error"] = A.load.Ref(ptr + 8, undefined);
      x["frameId"] = A.load.Int64(ptr + 16);
      x["frameType"] = A.load.Enum(ptr + 24, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["fromCache"] = A.load.Bool(ptr + 28);
      x["initiator"] = A.load.Ref(ptr + 32, undefined);
      x["ip"] = A.load.Ref(ptr + 36, undefined);
      x["method"] = A.load.Ref(ptr + 40, undefined);
      x["parentDocumentId"] = A.load.Ref(ptr + 44, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 48);
      x["requestId"] = A.load.Ref(ptr + 56, undefined);
      x["tabId"] = A.load.Int64(ptr + 64);
      x["timeStamp"] = A.load.Float64(ptr + 72);
      x["type"] = A.load.Enum(ptr + 80, [
        "main_frame",
        "sub_frame",
        "stylesheet",
        "script",
        "image",
        "font",
        "object",
        "xmlhttprequest",
        "ping",
        "csp_report",
        "media",
        "websocket",
        "webbundle",
        "other",
      ]);
      x["url"] = A.load.Ref(ptr + 84, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OnErrorOccurredOptions": (ref: heap.Ref<string>): number => {
      const idx = ["extraHeaders"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OnHeadersReceivedArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 88, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Int64(ptr + 32, 0);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Int64(ptr + 48, 0);
        A.store.Ref(ptr + 56, undefined);
        A.store.Int64(ptr + 64, 0);
        A.store.Float64(ptr + 72, 0);
        A.store.Enum(ptr + 80, -1);
        A.store.Ref(ptr + 84, undefined);
      } else {
        A.store.Bool(ptr + 88, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Ref(ptr + 20, x["initiator"]);
        A.store.Ref(ptr + 24, x["method"]);
        A.store.Ref(ptr + 28, x["parentDocumentId"]);
        A.store.Int64(ptr + 32, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Ref(ptr + 40, x["requestId"]);
        A.store.Ref(ptr + 44, x["responseHeaders"]);
        A.store.Int64(ptr + 48, x["statusCode"] === undefined ? 0 : (x["statusCode"] as number));
        A.store.Ref(ptr + 56, x["statusLine"]);
        A.store.Int64(ptr + 64, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 72, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Enum(
          ptr + 80,
          [
            "main_frame",
            "sub_frame",
            "stylesheet",
            "script",
            "image",
            "font",
            "object",
            "xmlhttprequest",
            "ping",
            "csp_report",
            "media",
            "websocket",
            "webbundle",
            "other",
          ].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 84, x["url"]);
      }
    },
    "load_OnHeadersReceivedArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["frameId"] = A.load.Int64(ptr + 8);
      x["frameType"] = A.load.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["initiator"] = A.load.Ref(ptr + 20, undefined);
      x["method"] = A.load.Ref(ptr + 24, undefined);
      x["parentDocumentId"] = A.load.Ref(ptr + 28, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 32);
      x["requestId"] = A.load.Ref(ptr + 40, undefined);
      x["responseHeaders"] = A.load.Ref(ptr + 44, undefined);
      x["statusCode"] = A.load.Int64(ptr + 48);
      x["statusLine"] = A.load.Ref(ptr + 56, undefined);
      x["tabId"] = A.load.Int64(ptr + 64);
      x["timeStamp"] = A.load.Float64(ptr + 72);
      x["type"] = A.load.Enum(ptr + 80, [
        "main_frame",
        "sub_frame",
        "stylesheet",
        "script",
        "image",
        "font",
        "object",
        "xmlhttprequest",
        "ping",
        "csp_report",
        "media",
        "websocket",
        "webbundle",
        "other",
      ]);
      x["url"] = A.load.Ref(ptr + 84, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OnHeadersReceivedOptions": (ref: heap.Ref<string>): number => {
      const idx = ["blocking", "responseHeaders", "extraHeaders"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OnResponseStartedArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 96, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Bool(ptr + 20, false);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Int64(ptr + 40, 0);
        A.store.Ref(ptr + 48, undefined);
        A.store.Ref(ptr + 52, undefined);
        A.store.Int64(ptr + 56, 0);
        A.store.Ref(ptr + 64, undefined);
        A.store.Int64(ptr + 72, 0);
        A.store.Float64(ptr + 80, 0);
        A.store.Enum(ptr + 88, -1);
        A.store.Ref(ptr + 92, undefined);
      } else {
        A.store.Bool(ptr + 96, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Bool(ptr + 20, x["fromCache"] ? true : false);
        A.store.Ref(ptr + 24, x["initiator"]);
        A.store.Ref(ptr + 28, x["ip"]);
        A.store.Ref(ptr + 32, x["method"]);
        A.store.Ref(ptr + 36, x["parentDocumentId"]);
        A.store.Int64(ptr + 40, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Ref(ptr + 48, x["requestId"]);
        A.store.Ref(ptr + 52, x["responseHeaders"]);
        A.store.Int64(ptr + 56, x["statusCode"] === undefined ? 0 : (x["statusCode"] as number));
        A.store.Ref(ptr + 64, x["statusLine"]);
        A.store.Int64(ptr + 72, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 80, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Enum(
          ptr + 88,
          [
            "main_frame",
            "sub_frame",
            "stylesheet",
            "script",
            "image",
            "font",
            "object",
            "xmlhttprequest",
            "ping",
            "csp_report",
            "media",
            "websocket",
            "webbundle",
            "other",
          ].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 92, x["url"]);
      }
    },
    "load_OnResponseStartedArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["frameId"] = A.load.Int64(ptr + 8);
      x["frameType"] = A.load.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["fromCache"] = A.load.Bool(ptr + 20);
      x["initiator"] = A.load.Ref(ptr + 24, undefined);
      x["ip"] = A.load.Ref(ptr + 28, undefined);
      x["method"] = A.load.Ref(ptr + 32, undefined);
      x["parentDocumentId"] = A.load.Ref(ptr + 36, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 40);
      x["requestId"] = A.load.Ref(ptr + 48, undefined);
      x["responseHeaders"] = A.load.Ref(ptr + 52, undefined);
      x["statusCode"] = A.load.Int64(ptr + 56);
      x["statusLine"] = A.load.Ref(ptr + 64, undefined);
      x["tabId"] = A.load.Int64(ptr + 72);
      x["timeStamp"] = A.load.Float64(ptr + 80);
      x["type"] = A.load.Enum(ptr + 88, [
        "main_frame",
        "sub_frame",
        "stylesheet",
        "script",
        "image",
        "font",
        "object",
        "xmlhttprequest",
        "ping",
        "csp_report",
        "media",
        "websocket",
        "webbundle",
        "other",
      ]);
      x["url"] = A.load.Ref(ptr + 92, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OnResponseStartedOptions": (ref: heap.Ref<string>): number => {
      const idx = ["responseHeaders", "extraHeaders"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OnSendHeadersArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 72, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Int64(ptr + 32, 0);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Int64(ptr + 48, 0);
        A.store.Float64(ptr + 56, 0);
        A.store.Enum(ptr + 64, -1);
        A.store.Ref(ptr + 68, undefined);
      } else {
        A.store.Bool(ptr + 72, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Ref(ptr + 20, x["initiator"]);
        A.store.Ref(ptr + 24, x["method"]);
        A.store.Ref(ptr + 28, x["parentDocumentId"]);
        A.store.Int64(ptr + 32, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Ref(ptr + 40, x["requestHeaders"]);
        A.store.Ref(ptr + 44, x["requestId"]);
        A.store.Int64(ptr + 48, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 56, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Enum(
          ptr + 64,
          [
            "main_frame",
            "sub_frame",
            "stylesheet",
            "script",
            "image",
            "font",
            "object",
            "xmlhttprequest",
            "ping",
            "csp_report",
            "media",
            "websocket",
            "webbundle",
            "other",
          ].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 68, x["url"]);
      }
    },
    "load_OnSendHeadersArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["frameId"] = A.load.Int64(ptr + 8);
      x["frameType"] = A.load.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["initiator"] = A.load.Ref(ptr + 20, undefined);
      x["method"] = A.load.Ref(ptr + 24, undefined);
      x["parentDocumentId"] = A.load.Ref(ptr + 28, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 32);
      x["requestHeaders"] = A.load.Ref(ptr + 40, undefined);
      x["requestId"] = A.load.Ref(ptr + 44, undefined);
      x["tabId"] = A.load.Int64(ptr + 48);
      x["timeStamp"] = A.load.Float64(ptr + 56);
      x["type"] = A.load.Enum(ptr + 64, [
        "main_frame",
        "sub_frame",
        "stylesheet",
        "script",
        "image",
        "font",
        "object",
        "xmlhttprequest",
        "ping",
        "csp_report",
        "media",
        "websocket",
        "webbundle",
        "other",
      ]);
      x["url"] = A.load.Ref(ptr + 68, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OnSendHeadersOptions": (ref: heap.Ref<string>): number => {
      const idx = ["requestHeaders", "extraHeaders"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RequestFilter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 24, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 25, false);
        A.store.Int64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 26, true);
        A.store.Bool(ptr + 24, "tabId" in x ? true : false);
        A.store.Int64(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Ref(ptr + 8, x["types"]);
        A.store.Ref(ptr + 12, x["urls"]);
        A.store.Bool(ptr + 25, "windowId" in x ? true : false);
        A.store.Int64(ptr + 16, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_RequestFilter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["tabId"] = A.load.Int64(ptr + 0);
      } else {
        delete x["tabId"];
      }
      x["types"] = A.load.Ref(ptr + 8, undefined);
      x["urls"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 25)) {
        x["windowId"] = A.load.Int64(ptr + 16);
      } else {
        delete x["windowId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_HandlerBehaviorChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest && "handlerBehaviorChanged" in WEBEXT?.webRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HandlerBehaviorChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.handlerBehaviorChanged);
    },
    "call_HandlerBehaviorChanged": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webRequest.handlerBehaviorChanged();
      A.store.Ref(retPtr, _ret);
    },
    "try_HandlerBehaviorChanged": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.handlerBehaviorChanged();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnActionIgnored": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onActionIgnored && "addListener" in WEBEXT?.webRequest?.onActionIgnored) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnActionIgnored": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onActionIgnored.addListener);
    },
    "call_OnActionIgnored": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onActionIgnored.addListener(A.H.get<object>(callback));
    },
    "try_OnActionIgnored": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onActionIgnored.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffActionIgnored": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onActionIgnored && "removeListener" in WEBEXT?.webRequest?.onActionIgnored) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffActionIgnored": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onActionIgnored.removeListener);
    },
    "call_OffActionIgnored": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onActionIgnored.removeListener(A.H.get<object>(callback));
    },
    "try_OffActionIgnored": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onActionIgnored.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnActionIgnored": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onActionIgnored && "hasListener" in WEBEXT?.webRequest?.onActionIgnored) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnActionIgnored": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onActionIgnored.hasListener);
    },
    "call_HasOnActionIgnored": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onActionIgnored.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnActionIgnored": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onActionIgnored.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAuthRequired": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onAuthRequired && "addListener" in WEBEXT?.webRequest?.onAuthRequired) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAuthRequired": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onAuthRequired.addListener);
    },
    "call_OnAuthRequired": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onAuthRequired.addListener(A.H.get<object>(callback));
    },
    "try_OnAuthRequired": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onAuthRequired.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAuthRequired": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onAuthRequired && "removeListener" in WEBEXT?.webRequest?.onAuthRequired) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAuthRequired": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onAuthRequired.removeListener);
    },
    "call_OffAuthRequired": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onAuthRequired.removeListener(A.H.get<object>(callback));
    },
    "try_OffAuthRequired": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onAuthRequired.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAuthRequired": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onAuthRequired && "hasListener" in WEBEXT?.webRequest?.onAuthRequired) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAuthRequired": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onAuthRequired.hasListener);
    },
    "call_HasOnAuthRequired": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onAuthRequired.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAuthRequired": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onAuthRequired.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnBeforeRedirect": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onBeforeRedirect && "addListener" in WEBEXT?.webRequest?.onBeforeRedirect) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnBeforeRedirect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onBeforeRedirect.addListener);
    },
    "call_OnBeforeRedirect": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onBeforeRedirect.addListener(A.H.get<object>(callback));
    },
    "try_OnBeforeRedirect": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onBeforeRedirect.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffBeforeRedirect": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onBeforeRedirect && "removeListener" in WEBEXT?.webRequest?.onBeforeRedirect) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffBeforeRedirect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onBeforeRedirect.removeListener);
    },
    "call_OffBeforeRedirect": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onBeforeRedirect.removeListener(A.H.get<object>(callback));
    },
    "try_OffBeforeRedirect": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onBeforeRedirect.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnBeforeRedirect": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onBeforeRedirect && "hasListener" in WEBEXT?.webRequest?.onBeforeRedirect) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnBeforeRedirect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onBeforeRedirect.hasListener);
    },
    "call_HasOnBeforeRedirect": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onBeforeRedirect.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnBeforeRedirect": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onBeforeRedirect.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnBeforeRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onBeforeRequest && "addListener" in WEBEXT?.webRequest?.onBeforeRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnBeforeRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onBeforeRequest.addListener);
    },
    "call_OnBeforeRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onBeforeRequest.addListener(A.H.get<object>(callback));
    },
    "try_OnBeforeRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onBeforeRequest.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffBeforeRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onBeforeRequest && "removeListener" in WEBEXT?.webRequest?.onBeforeRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffBeforeRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onBeforeRequest.removeListener);
    },
    "call_OffBeforeRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onBeforeRequest.removeListener(A.H.get<object>(callback));
    },
    "try_OffBeforeRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onBeforeRequest.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnBeforeRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onBeforeRequest && "hasListener" in WEBEXT?.webRequest?.onBeforeRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnBeforeRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onBeforeRequest.hasListener);
    },
    "call_HasOnBeforeRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onBeforeRequest.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnBeforeRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onBeforeRequest.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnBeforeSendHeaders": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onBeforeSendHeaders && "addListener" in WEBEXT?.webRequest?.onBeforeSendHeaders) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnBeforeSendHeaders": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onBeforeSendHeaders.addListener);
    },
    "call_OnBeforeSendHeaders": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onBeforeSendHeaders.addListener(A.H.get<object>(callback));
    },
    "try_OnBeforeSendHeaders": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onBeforeSendHeaders.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffBeforeSendHeaders": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onBeforeSendHeaders && "removeListener" in WEBEXT?.webRequest?.onBeforeSendHeaders) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffBeforeSendHeaders": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onBeforeSendHeaders.removeListener);
    },
    "call_OffBeforeSendHeaders": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onBeforeSendHeaders.removeListener(A.H.get<object>(callback));
    },
    "try_OffBeforeSendHeaders": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onBeforeSendHeaders.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnBeforeSendHeaders": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onBeforeSendHeaders && "hasListener" in WEBEXT?.webRequest?.onBeforeSendHeaders) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnBeforeSendHeaders": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onBeforeSendHeaders.hasListener);
    },
    "call_HasOnBeforeSendHeaders": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onBeforeSendHeaders.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnBeforeSendHeaders": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onBeforeSendHeaders.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCompleted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onCompleted && "addListener" in WEBEXT?.webRequest?.onCompleted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCompleted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onCompleted.addListener);
    },
    "call_OnCompleted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onCompleted.addListener(A.H.get<object>(callback));
    },
    "try_OnCompleted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onCompleted.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCompleted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onCompleted && "removeListener" in WEBEXT?.webRequest?.onCompleted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCompleted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onCompleted.removeListener);
    },
    "call_OffCompleted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onCompleted.removeListener(A.H.get<object>(callback));
    },
    "try_OffCompleted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onCompleted.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCompleted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onCompleted && "hasListener" in WEBEXT?.webRequest?.onCompleted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCompleted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onCompleted.hasListener);
    },
    "call_HasOnCompleted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onCompleted.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCompleted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onCompleted.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnErrorOccurred": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onErrorOccurred && "addListener" in WEBEXT?.webRequest?.onErrorOccurred) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnErrorOccurred": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onErrorOccurred.addListener);
    },
    "call_OnErrorOccurred": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onErrorOccurred.addListener(A.H.get<object>(callback));
    },
    "try_OnErrorOccurred": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onErrorOccurred.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffErrorOccurred": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onErrorOccurred && "removeListener" in WEBEXT?.webRequest?.onErrorOccurred) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffErrorOccurred": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onErrorOccurred.removeListener);
    },
    "call_OffErrorOccurred": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onErrorOccurred.removeListener(A.H.get<object>(callback));
    },
    "try_OffErrorOccurred": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onErrorOccurred.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnErrorOccurred": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onErrorOccurred && "hasListener" in WEBEXT?.webRequest?.onErrorOccurred) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnErrorOccurred": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onErrorOccurred.hasListener);
    },
    "call_HasOnErrorOccurred": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onErrorOccurred.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnErrorOccurred": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onErrorOccurred.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnHeadersReceived": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onHeadersReceived && "addListener" in WEBEXT?.webRequest?.onHeadersReceived) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnHeadersReceived": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onHeadersReceived.addListener);
    },
    "call_OnHeadersReceived": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onHeadersReceived.addListener(A.H.get<object>(callback));
    },
    "try_OnHeadersReceived": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onHeadersReceived.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffHeadersReceived": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onHeadersReceived && "removeListener" in WEBEXT?.webRequest?.onHeadersReceived) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffHeadersReceived": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onHeadersReceived.removeListener);
    },
    "call_OffHeadersReceived": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onHeadersReceived.removeListener(A.H.get<object>(callback));
    },
    "try_OffHeadersReceived": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onHeadersReceived.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnHeadersReceived": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onHeadersReceived && "hasListener" in WEBEXT?.webRequest?.onHeadersReceived) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnHeadersReceived": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onHeadersReceived.hasListener);
    },
    "call_HasOnHeadersReceived": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onHeadersReceived.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnHeadersReceived": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onHeadersReceived.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnResponseStarted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onResponseStarted && "addListener" in WEBEXT?.webRequest?.onResponseStarted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnResponseStarted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onResponseStarted.addListener);
    },
    "call_OnResponseStarted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onResponseStarted.addListener(A.H.get<object>(callback));
    },
    "try_OnResponseStarted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onResponseStarted.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffResponseStarted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onResponseStarted && "removeListener" in WEBEXT?.webRequest?.onResponseStarted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffResponseStarted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onResponseStarted.removeListener);
    },
    "call_OffResponseStarted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onResponseStarted.removeListener(A.H.get<object>(callback));
    },
    "try_OffResponseStarted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onResponseStarted.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnResponseStarted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onResponseStarted && "hasListener" in WEBEXT?.webRequest?.onResponseStarted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnResponseStarted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onResponseStarted.hasListener);
    },
    "call_HasOnResponseStarted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onResponseStarted.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnResponseStarted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onResponseStarted.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSendHeaders": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onSendHeaders && "addListener" in WEBEXT?.webRequest?.onSendHeaders) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSendHeaders": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onSendHeaders.addListener);
    },
    "call_OnSendHeaders": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onSendHeaders.addListener(A.H.get<object>(callback));
    },
    "try_OnSendHeaders": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onSendHeaders.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSendHeaders": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onSendHeaders && "removeListener" in WEBEXT?.webRequest?.onSendHeaders) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSendHeaders": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onSendHeaders.removeListener);
    },
    "call_OffSendHeaders": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onSendHeaders.removeListener(A.H.get<object>(callback));
    },
    "try_OffSendHeaders": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onSendHeaders.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSendHeaders": (): heap.Ref<boolean> => {
      if (WEBEXT?.webRequest?.onSendHeaders && "hasListener" in WEBEXT?.webRequest?.onSendHeaders) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSendHeaders": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webRequest.onSendHeaders.hasListener);
    },
    "call_HasOnSendHeaders": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webRequest.onSendHeaders.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSendHeaders": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webRequest.onSendHeaders.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
