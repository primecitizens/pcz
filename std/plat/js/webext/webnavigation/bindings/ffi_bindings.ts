import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/webnavigation", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_GetAllFramesArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Int64(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_GetAllFramesArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["tabId"] = A.load.Int64(ptr + 0);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetAllFramesReturnTypeElem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 52, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Enum(ptr + 24, -1);
        A.store.Ref(ptr + 28, undefined);
        A.store.Int64(ptr + 32, 0);
        A.store.Int64(ptr + 40, 0);
        A.store.Ref(ptr + 48, undefined);
      } else {
        A.store.Bool(ptr + 52, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Bool(ptr + 8, x["errorOccurred"] ? true : false);
        A.store.Int64(ptr + 16, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 24, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Ref(ptr + 28, x["parentDocumentId"]);
        A.store.Int64(ptr + 32, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Int64(ptr + 40, x["processId"] === undefined ? 0 : (x["processId"] as number));
        A.store.Ref(ptr + 48, x["url"]);
      }
    },
    "load_GetAllFramesReturnTypeElem": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["errorOccurred"] = A.load.Bool(ptr + 8);
      x["frameId"] = A.load.Int64(ptr + 16);
      x["frameType"] = A.load.Enum(ptr + 24, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["parentDocumentId"] = A.load.Ref(ptr + 28, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 32);
      x["processId"] = A.load.Int64(ptr + 40);
      x["url"] = A.load.Ref(ptr + 48, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetFrameArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 35, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 32, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Bool(ptr + 33, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Bool(ptr + 34, false);
        A.store.Int64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 35, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Bool(ptr + 32, "frameId" in x ? true : false);
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Bool(ptr + 33, "processId" in x ? true : false);
        A.store.Int64(ptr + 16, x["processId"] === undefined ? 0 : (x["processId"] as number));
        A.store.Bool(ptr + 34, "tabId" in x ? true : false);
        A.store.Int64(ptr + 24, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_GetFrameArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 32)) {
        x["frameId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["frameId"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["processId"] = A.load.Int64(ptr + 16);
      } else {
        delete x["processId"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["tabId"] = A.load.Int64(ptr + 24);
      } else {
        delete x["tabId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetFrameReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 36, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 12, -1);
        A.store.Ref(ptr + 16, undefined);
        A.store.Int64(ptr + 24, 0);
        A.store.Ref(ptr + 32, undefined);
      } else {
        A.store.Bool(ptr + 36, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Bool(ptr + 8, x["errorOccurred"] ? true : false);
        A.store.Enum(ptr + 12, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Ref(ptr + 16, x["parentDocumentId"]);
        A.store.Int64(ptr + 24, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Ref(ptr + 32, x["url"]);
      }
    },
    "load_GetFrameReturnType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["errorOccurred"] = A.load.Bool(ptr + 8);
      x["frameType"] = A.load.Enum(ptr + 12, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["parentDocumentId"] = A.load.Ref(ptr + 16, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 24);
      x["url"] = A.load.Ref(ptr + 32, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnBeforeNavigateArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 60, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Ref(ptr + 20, undefined);
        A.store.Int64(ptr + 24, 0);
        A.store.Int64(ptr + 32, 0);
        A.store.Int64(ptr + 40, 0);
        A.store.Float64(ptr + 48, 0);
        A.store.Ref(ptr + 56, undefined);
      } else {
        A.store.Bool(ptr + 60, true);
        A.store.Enum(
          ptr + 0,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Ref(ptr + 20, x["parentDocumentId"]);
        A.store.Int64(ptr + 24, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Int64(ptr + 32, x["processId"] === undefined ? 0 : (x["processId"] as number));
        A.store.Int64(ptr + 40, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 48, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Ref(ptr + 56, x["url"]);
      }
    },
    "load_OnBeforeNavigateArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentLifecycle"] = A.load.Enum(ptr + 0, ["prerender", "active", "cached", "pending_deletion"]);
      x["frameId"] = A.load.Int64(ptr + 8);
      x["frameType"] = A.load.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["parentDocumentId"] = A.load.Ref(ptr + 20, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 24);
      x["processId"] = A.load.Int64(ptr + 32);
      x["tabId"] = A.load.Int64(ptr + 40);
      x["timeStamp"] = A.load.Float64(ptr + 48);
      x["url"] = A.load.Ref(ptr + 56, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_TransitionQualifier": (ref: heap.Ref<string>): number => {
      const idx = ["client_redirect", "server_redirect", "forward_back", "from_address_bar"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_TransitionType": (ref: heap.Ref<string>): number => {
      const idx = [
        "link",
        "typed",
        "auto_bookmark",
        "auto_subframe",
        "manual_subframe",
        "generated",
        "start_page",
        "form_submit",
        "reload",
        "keyword",
        "keyword_generated",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OnCommittedArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 68, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Ref(ptr + 20, undefined);
        A.store.Int64(ptr + 24, 0);
        A.store.Int64(ptr + 32, 0);
        A.store.Int64(ptr + 40, 0);
        A.store.Float64(ptr + 48, 0);
        A.store.Ref(ptr + 56, undefined);
        A.store.Enum(ptr + 60, -1);
        A.store.Ref(ptr + 64, undefined);
      } else {
        A.store.Bool(ptr + 68, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Ref(ptr + 20, x["parentDocumentId"]);
        A.store.Int64(ptr + 24, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Int64(ptr + 32, x["processId"] === undefined ? 0 : (x["processId"] as number));
        A.store.Int64(ptr + 40, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 48, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Ref(ptr + 56, x["transitionQualifiers"]);
        A.store.Enum(
          ptr + 60,
          [
            "link",
            "typed",
            "auto_bookmark",
            "auto_subframe",
            "manual_subframe",
            "generated",
            "start_page",
            "form_submit",
            "reload",
            "keyword",
            "keyword_generated",
          ].indexOf(x["transitionType"] as string)
        );
        A.store.Ref(ptr + 64, x["url"]);
      }
    },
    "load_OnCommittedArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["frameId"] = A.load.Int64(ptr + 8);
      x["frameType"] = A.load.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["parentDocumentId"] = A.load.Ref(ptr + 20, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 24);
      x["processId"] = A.load.Int64(ptr + 32);
      x["tabId"] = A.load.Int64(ptr + 40);
      x["timeStamp"] = A.load.Float64(ptr + 48);
      x["transitionQualifiers"] = A.load.Ref(ptr + 56, undefined);
      x["transitionType"] = A.load.Enum(ptr + 60, [
        "link",
        "typed",
        "auto_bookmark",
        "auto_subframe",
        "manual_subframe",
        "generated",
        "start_page",
        "form_submit",
        "reload",
        "keyword",
        "keyword_generated",
      ]);
      x["url"] = A.load.Ref(ptr + 64, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnCompletedArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 60, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Ref(ptr + 20, undefined);
        A.store.Int64(ptr + 24, 0);
        A.store.Int64(ptr + 32, 0);
        A.store.Int64(ptr + 40, 0);
        A.store.Float64(ptr + 48, 0);
        A.store.Ref(ptr + 56, undefined);
      } else {
        A.store.Bool(ptr + 60, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Ref(ptr + 20, x["parentDocumentId"]);
        A.store.Int64(ptr + 24, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Int64(ptr + 32, x["processId"] === undefined ? 0 : (x["processId"] as number));
        A.store.Int64(ptr + 40, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 48, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Ref(ptr + 56, x["url"]);
      }
    },
    "load_OnCompletedArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["frameId"] = A.load.Int64(ptr + 8);
      x["frameType"] = A.load.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["parentDocumentId"] = A.load.Ref(ptr + 20, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 24);
      x["processId"] = A.load.Int64(ptr + 32);
      x["tabId"] = A.load.Int64(ptr + 40);
      x["timeStamp"] = A.load.Float64(ptr + 48);
      x["url"] = A.load.Ref(ptr + 56, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnCreatedNavigationTargetArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 44, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
        A.store.Int64(ptr + 16, 0);
        A.store.Int64(ptr + 24, 0);
        A.store.Float64(ptr + 32, 0);
        A.store.Ref(ptr + 40, undefined);
      } else {
        A.store.Bool(ptr + 44, true);
        A.store.Int64(ptr + 0, x["sourceFrameId"] === undefined ? 0 : (x["sourceFrameId"] as number));
        A.store.Int64(ptr + 8, x["sourceProcessId"] === undefined ? 0 : (x["sourceProcessId"] as number));
        A.store.Int64(ptr + 16, x["sourceTabId"] === undefined ? 0 : (x["sourceTabId"] as number));
        A.store.Int64(ptr + 24, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 32, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Ref(ptr + 40, x["url"]);
      }
    },
    "load_OnCreatedNavigationTargetArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["sourceFrameId"] = A.load.Int64(ptr + 0);
      x["sourceProcessId"] = A.load.Int64(ptr + 8);
      x["sourceTabId"] = A.load.Int64(ptr + 16);
      x["tabId"] = A.load.Int64(ptr + 24);
      x["timeStamp"] = A.load.Float64(ptr + 32);
      x["url"] = A.load.Ref(ptr + 40, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnDOMContentLoadedArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 60, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Ref(ptr + 20, undefined);
        A.store.Int64(ptr + 24, 0);
        A.store.Int64(ptr + 32, 0);
        A.store.Int64(ptr + 40, 0);
        A.store.Float64(ptr + 48, 0);
        A.store.Ref(ptr + 56, undefined);
      } else {
        A.store.Bool(ptr + 60, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Ref(ptr + 20, x["parentDocumentId"]);
        A.store.Int64(ptr + 24, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Int64(ptr + 32, x["processId"] === undefined ? 0 : (x["processId"] as number));
        A.store.Int64(ptr + 40, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 48, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Ref(ptr + 56, x["url"]);
      }
    },
    "load_OnDOMContentLoadedArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["frameId"] = A.load.Int64(ptr + 8);
      x["frameType"] = A.load.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["parentDocumentId"] = A.load.Ref(ptr + 20, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 24);
      x["processId"] = A.load.Int64(ptr + 32);
      x["tabId"] = A.load.Int64(ptr + 40);
      x["timeStamp"] = A.load.Float64(ptr + 48);
      x["url"] = A.load.Ref(ptr + 56, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnErrorOccurredArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 68, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
        A.store.Int64(ptr + 16, 0);
        A.store.Enum(ptr + 24, -1);
        A.store.Ref(ptr + 28, undefined);
        A.store.Int64(ptr + 32, 0);
        A.store.Int64(ptr + 40, 0);
        A.store.Int64(ptr + 48, 0);
        A.store.Float64(ptr + 56, 0);
        A.store.Ref(ptr + 64, undefined);
      } else {
        A.store.Bool(ptr + 68, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Ref(ptr + 8, x["error"]);
        A.store.Int64(ptr + 16, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 24, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Ref(ptr + 28, x["parentDocumentId"]);
        A.store.Int64(ptr + 32, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Int64(ptr + 40, x["processId"] === undefined ? 0 : (x["processId"] as number));
        A.store.Int64(ptr + 48, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 56, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Ref(ptr + 64, x["url"]);
      }
    },
    "load_OnErrorOccurredArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["error"] = A.load.Ref(ptr + 8, undefined);
      x["frameId"] = A.load.Int64(ptr + 16);
      x["frameType"] = A.load.Enum(ptr + 24, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["parentDocumentId"] = A.load.Ref(ptr + 28, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 32);
      x["processId"] = A.load.Int64(ptr + 40);
      x["tabId"] = A.load.Int64(ptr + 48);
      x["timeStamp"] = A.load.Float64(ptr + 56);
      x["url"] = A.load.Ref(ptr + 64, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnHistoryStateUpdatedArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 68, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Ref(ptr + 20, undefined);
        A.store.Int64(ptr + 24, 0);
        A.store.Int64(ptr + 32, 0);
        A.store.Int64(ptr + 40, 0);
        A.store.Float64(ptr + 48, 0);
        A.store.Ref(ptr + 56, undefined);
        A.store.Enum(ptr + 60, -1);
        A.store.Ref(ptr + 64, undefined);
      } else {
        A.store.Bool(ptr + 68, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Ref(ptr + 20, x["parentDocumentId"]);
        A.store.Int64(ptr + 24, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Int64(ptr + 32, x["processId"] === undefined ? 0 : (x["processId"] as number));
        A.store.Int64(ptr + 40, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 48, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Ref(ptr + 56, x["transitionQualifiers"]);
        A.store.Enum(
          ptr + 60,
          [
            "link",
            "typed",
            "auto_bookmark",
            "auto_subframe",
            "manual_subframe",
            "generated",
            "start_page",
            "form_submit",
            "reload",
            "keyword",
            "keyword_generated",
          ].indexOf(x["transitionType"] as string)
        );
        A.store.Ref(ptr + 64, x["url"]);
      }
    },
    "load_OnHistoryStateUpdatedArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["frameId"] = A.load.Int64(ptr + 8);
      x["frameType"] = A.load.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["parentDocumentId"] = A.load.Ref(ptr + 20, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 24);
      x["processId"] = A.load.Int64(ptr + 32);
      x["tabId"] = A.load.Int64(ptr + 40);
      x["timeStamp"] = A.load.Float64(ptr + 48);
      x["transitionQualifiers"] = A.load.Ref(ptr + 56, undefined);
      x["transitionType"] = A.load.Enum(ptr + 60, [
        "link",
        "typed",
        "auto_bookmark",
        "auto_subframe",
        "manual_subframe",
        "generated",
        "start_page",
        "form_submit",
        "reload",
        "keyword",
        "keyword_generated",
      ]);
      x["url"] = A.load.Ref(ptr + 64, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnReferenceFragmentUpdatedArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 68, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Ref(ptr + 20, undefined);
        A.store.Int64(ptr + 24, 0);
        A.store.Int64(ptr + 32, 0);
        A.store.Int64(ptr + 40, 0);
        A.store.Float64(ptr + 48, 0);
        A.store.Ref(ptr + 56, undefined);
        A.store.Enum(ptr + 60, -1);
        A.store.Ref(ptr + 64, undefined);
      } else {
        A.store.Bool(ptr + 68, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Ref(ptr + 20, x["parentDocumentId"]);
        A.store.Int64(ptr + 24, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Int64(ptr + 32, x["processId"] === undefined ? 0 : (x["processId"] as number));
        A.store.Int64(ptr + 40, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 48, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Ref(ptr + 56, x["transitionQualifiers"]);
        A.store.Enum(
          ptr + 60,
          [
            "link",
            "typed",
            "auto_bookmark",
            "auto_subframe",
            "manual_subframe",
            "generated",
            "start_page",
            "form_submit",
            "reload",
            "keyword",
            "keyword_generated",
          ].indexOf(x["transitionType"] as string)
        );
        A.store.Ref(ptr + 64, x["url"]);
      }
    },
    "load_OnReferenceFragmentUpdatedArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["frameId"] = A.load.Int64(ptr + 8);
      x["frameType"] = A.load.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["parentDocumentId"] = A.load.Ref(ptr + 20, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 24);
      x["processId"] = A.load.Int64(ptr + 32);
      x["tabId"] = A.load.Int64(ptr + 40);
      x["timeStamp"] = A.load.Float64(ptr + 48);
      x["transitionQualifiers"] = A.load.Ref(ptr + 56, undefined);
      x["transitionType"] = A.load.Enum(ptr + 60, [
        "link",
        "typed",
        "auto_bookmark",
        "auto_subframe",
        "manual_subframe",
        "generated",
        "start_page",
        "form_submit",
        "reload",
        "keyword",
        "keyword_generated",
      ]);
      x["url"] = A.load.Ref(ptr + 64, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnTabReplacedArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 24, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
        A.store.Float64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 24, true);
        A.store.Int64(ptr + 0, x["replacedTabId"] === undefined ? 0 : (x["replacedTabId"] as number));
        A.store.Int64(ptr + 8, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 16, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
      }
    },
    "load_OnTabReplacedArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["replacedTabId"] = A.load.Int64(ptr + 0);
      x["tabId"] = A.load.Int64(ptr + 8);
      x["timeStamp"] = A.load.Float64(ptr + 16);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetAllFrames": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation && "getAllFrames" in WEBEXT?.webNavigation) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAllFrames": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.getAllFrames);
    },
    "call_GetAllFrames": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["tabId"] = A.load.Int64(details + 0);

      const _ret = WEBEXT.webNavigation.getAllFrames(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAllFrames": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["tabId"] = A.load.Int64(details + 0);

        const _ret = WEBEXT.webNavigation.getAllFrames(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetFrame": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation && "getFrame" in WEBEXT?.webNavigation) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetFrame": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.getFrame);
    },
    "call_GetFrame": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["documentId"] = A.load.Ref(details + 0, undefined);
      if (A.load.Bool(details + 32)) {
        details_ffi["frameId"] = A.load.Int64(details + 8);
      }
      if (A.load.Bool(details + 33)) {
        details_ffi["processId"] = A.load.Int64(details + 16);
      }
      if (A.load.Bool(details + 34)) {
        details_ffi["tabId"] = A.load.Int64(details + 24);
      }

      const _ret = WEBEXT.webNavigation.getFrame(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetFrame": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["documentId"] = A.load.Ref(details + 0, undefined);
        if (A.load.Bool(details + 32)) {
          details_ffi["frameId"] = A.load.Int64(details + 8);
        }
        if (A.load.Bool(details + 33)) {
          details_ffi["processId"] = A.load.Int64(details + 16);
        }
        if (A.load.Bool(details + 34)) {
          details_ffi["tabId"] = A.load.Int64(details + 24);
        }

        const _ret = WEBEXT.webNavigation.getFrame(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnBeforeNavigate": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onBeforeNavigate && "addListener" in WEBEXT?.webNavigation?.onBeforeNavigate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnBeforeNavigate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onBeforeNavigate.addListener);
    },
    "call_OnBeforeNavigate": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onBeforeNavigate.addListener(A.H.get<object>(callback));
    },
    "try_OnBeforeNavigate": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onBeforeNavigate.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffBeforeNavigate": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onBeforeNavigate && "removeListener" in WEBEXT?.webNavigation?.onBeforeNavigate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffBeforeNavigate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onBeforeNavigate.removeListener);
    },
    "call_OffBeforeNavigate": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onBeforeNavigate.removeListener(A.H.get<object>(callback));
    },
    "try_OffBeforeNavigate": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onBeforeNavigate.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnBeforeNavigate": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onBeforeNavigate && "hasListener" in WEBEXT?.webNavigation?.onBeforeNavigate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnBeforeNavigate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onBeforeNavigate.hasListener);
    },
    "call_HasOnBeforeNavigate": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onBeforeNavigate.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnBeforeNavigate": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onBeforeNavigate.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCommitted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onCommitted && "addListener" in WEBEXT?.webNavigation?.onCommitted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCommitted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onCommitted.addListener);
    },
    "call_OnCommitted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onCommitted.addListener(A.H.get<object>(callback));
    },
    "try_OnCommitted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onCommitted.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCommitted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onCommitted && "removeListener" in WEBEXT?.webNavigation?.onCommitted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCommitted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onCommitted.removeListener);
    },
    "call_OffCommitted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onCommitted.removeListener(A.H.get<object>(callback));
    },
    "try_OffCommitted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onCommitted.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCommitted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onCommitted && "hasListener" in WEBEXT?.webNavigation?.onCommitted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCommitted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onCommitted.hasListener);
    },
    "call_HasOnCommitted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onCommitted.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCommitted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onCommitted.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCompleted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onCompleted && "addListener" in WEBEXT?.webNavigation?.onCompleted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCompleted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onCompleted.addListener);
    },
    "call_OnCompleted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onCompleted.addListener(A.H.get<object>(callback));
    },
    "try_OnCompleted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onCompleted.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCompleted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onCompleted && "removeListener" in WEBEXT?.webNavigation?.onCompleted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCompleted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onCompleted.removeListener);
    },
    "call_OffCompleted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onCompleted.removeListener(A.H.get<object>(callback));
    },
    "try_OffCompleted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onCompleted.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCompleted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onCompleted && "hasListener" in WEBEXT?.webNavigation?.onCompleted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCompleted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onCompleted.hasListener);
    },
    "call_HasOnCompleted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onCompleted.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCompleted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onCompleted.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCreatedNavigationTarget": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webNavigation?.onCreatedNavigationTarget &&
        "addListener" in WEBEXT?.webNavigation?.onCreatedNavigationTarget
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCreatedNavigationTarget": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onCreatedNavigationTarget.addListener);
    },
    "call_OnCreatedNavigationTarget": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onCreatedNavigationTarget.addListener(A.H.get<object>(callback));
    },
    "try_OnCreatedNavigationTarget": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onCreatedNavigationTarget.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCreatedNavigationTarget": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webNavigation?.onCreatedNavigationTarget &&
        "removeListener" in WEBEXT?.webNavigation?.onCreatedNavigationTarget
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCreatedNavigationTarget": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onCreatedNavigationTarget.removeListener);
    },
    "call_OffCreatedNavigationTarget": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onCreatedNavigationTarget.removeListener(A.H.get<object>(callback));
    },
    "try_OffCreatedNavigationTarget": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onCreatedNavigationTarget.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCreatedNavigationTarget": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webNavigation?.onCreatedNavigationTarget &&
        "hasListener" in WEBEXT?.webNavigation?.onCreatedNavigationTarget
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCreatedNavigationTarget": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onCreatedNavigationTarget.hasListener);
    },
    "call_HasOnCreatedNavigationTarget": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onCreatedNavigationTarget.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCreatedNavigationTarget": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onCreatedNavigationTarget.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDOMContentLoaded": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onDOMContentLoaded && "addListener" in WEBEXT?.webNavigation?.onDOMContentLoaded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDOMContentLoaded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onDOMContentLoaded.addListener);
    },
    "call_OnDOMContentLoaded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onDOMContentLoaded.addListener(A.H.get<object>(callback));
    },
    "try_OnDOMContentLoaded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onDOMContentLoaded.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDOMContentLoaded": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onDOMContentLoaded && "removeListener" in WEBEXT?.webNavigation?.onDOMContentLoaded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDOMContentLoaded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onDOMContentLoaded.removeListener);
    },
    "call_OffDOMContentLoaded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onDOMContentLoaded.removeListener(A.H.get<object>(callback));
    },
    "try_OffDOMContentLoaded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onDOMContentLoaded.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDOMContentLoaded": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onDOMContentLoaded && "hasListener" in WEBEXT?.webNavigation?.onDOMContentLoaded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDOMContentLoaded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onDOMContentLoaded.hasListener);
    },
    "call_HasOnDOMContentLoaded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onDOMContentLoaded.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDOMContentLoaded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onDOMContentLoaded.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnErrorOccurred": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onErrorOccurred && "addListener" in WEBEXT?.webNavigation?.onErrorOccurred) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnErrorOccurred": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onErrorOccurred.addListener);
    },
    "call_OnErrorOccurred": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onErrorOccurred.addListener(A.H.get<object>(callback));
    },
    "try_OnErrorOccurred": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onErrorOccurred.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffErrorOccurred": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onErrorOccurred && "removeListener" in WEBEXT?.webNavigation?.onErrorOccurred) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffErrorOccurred": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onErrorOccurred.removeListener);
    },
    "call_OffErrorOccurred": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onErrorOccurred.removeListener(A.H.get<object>(callback));
    },
    "try_OffErrorOccurred": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onErrorOccurred.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnErrorOccurred": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onErrorOccurred && "hasListener" in WEBEXT?.webNavigation?.onErrorOccurred) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnErrorOccurred": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onErrorOccurred.hasListener);
    },
    "call_HasOnErrorOccurred": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onErrorOccurred.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnErrorOccurred": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onErrorOccurred.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnHistoryStateUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webNavigation?.onHistoryStateUpdated &&
        "addListener" in WEBEXT?.webNavigation?.onHistoryStateUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnHistoryStateUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onHistoryStateUpdated.addListener);
    },
    "call_OnHistoryStateUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onHistoryStateUpdated.addListener(A.H.get<object>(callback));
    },
    "try_OnHistoryStateUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onHistoryStateUpdated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffHistoryStateUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webNavigation?.onHistoryStateUpdated &&
        "removeListener" in WEBEXT?.webNavigation?.onHistoryStateUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffHistoryStateUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onHistoryStateUpdated.removeListener);
    },
    "call_OffHistoryStateUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onHistoryStateUpdated.removeListener(A.H.get<object>(callback));
    },
    "try_OffHistoryStateUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onHistoryStateUpdated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnHistoryStateUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webNavigation?.onHistoryStateUpdated &&
        "hasListener" in WEBEXT?.webNavigation?.onHistoryStateUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnHistoryStateUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onHistoryStateUpdated.hasListener);
    },
    "call_HasOnHistoryStateUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onHistoryStateUpdated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnHistoryStateUpdated": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onHistoryStateUpdated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnReferenceFragmentUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webNavigation?.onReferenceFragmentUpdated &&
        "addListener" in WEBEXT?.webNavigation?.onReferenceFragmentUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnReferenceFragmentUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onReferenceFragmentUpdated.addListener);
    },
    "call_OnReferenceFragmentUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onReferenceFragmentUpdated.addListener(A.H.get<object>(callback));
    },
    "try_OnReferenceFragmentUpdated": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onReferenceFragmentUpdated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffReferenceFragmentUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webNavigation?.onReferenceFragmentUpdated &&
        "removeListener" in WEBEXT?.webNavigation?.onReferenceFragmentUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffReferenceFragmentUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onReferenceFragmentUpdated.removeListener);
    },
    "call_OffReferenceFragmentUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onReferenceFragmentUpdated.removeListener(A.H.get<object>(callback));
    },
    "try_OffReferenceFragmentUpdated": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onReferenceFragmentUpdated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnReferenceFragmentUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webNavigation?.onReferenceFragmentUpdated &&
        "hasListener" in WEBEXT?.webNavigation?.onReferenceFragmentUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnReferenceFragmentUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onReferenceFragmentUpdated.hasListener);
    },
    "call_HasOnReferenceFragmentUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onReferenceFragmentUpdated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnReferenceFragmentUpdated": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onReferenceFragmentUpdated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTabReplaced": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onTabReplaced && "addListener" in WEBEXT?.webNavigation?.onTabReplaced) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTabReplaced": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onTabReplaced.addListener);
    },
    "call_OnTabReplaced": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onTabReplaced.addListener(A.H.get<object>(callback));
    },
    "try_OnTabReplaced": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onTabReplaced.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffTabReplaced": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onTabReplaced && "removeListener" in WEBEXT?.webNavigation?.onTabReplaced) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffTabReplaced": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onTabReplaced.removeListener);
    },
    "call_OffTabReplaced": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onTabReplaced.removeListener(A.H.get<object>(callback));
    },
    "try_OffTabReplaced": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onTabReplaced.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnTabReplaced": (): heap.Ref<boolean> => {
      if (WEBEXT?.webNavigation?.onTabReplaced && "hasListener" in WEBEXT?.webNavigation?.onTabReplaced) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnTabReplaced": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webNavigation.onTabReplaced.hasListener);
    },
    "call_HasOnTabReplaced": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webNavigation.onTabReplaced.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnTabReplaced": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webNavigation.onTabReplaced.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
