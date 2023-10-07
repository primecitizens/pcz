import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/tabs", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_ConnectArgConnectInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 20, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Bool(ptr + 20, "frameId" in x ? true : false);
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Ref(ptr + 16, x["name"]);
      }
    },
    "load_ConnectArgConnectInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 20)) {
        x["frameId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["frameId"];
      }
      x["name"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_MutedInfoReason": (ref: heap.Ref<string>): number => {
      const idx = ["user", "capture", "extension"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_MutedInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["extensionId"]);
        A.store.Bool(ptr + 4, x["muted"] ? true : false);
        A.store.Enum(ptr + 8, ["user", "capture", "extension"].indexOf(x["reason"] as string));
      }
    },
    "load_MutedInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["extensionId"] = A.load.Ref(ptr + 0, undefined);
      x["muted"] = A.load.Bool(ptr + 4);
      x["reason"] = A.load.Enum(ptr + 8, ["user", "capture", "extension"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_TabStatus": (ref: heap.Ref<string>): number => {
      const idx = ["unloaded", "loading", "complete"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Tab": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 125, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 120, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 3, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Int64(ptr + 8, 0);
        A.store.Bool(ptr + 121, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 122, false);
        A.store.Int64(ptr + 32, 0);
        A.store.Bool(ptr + 40, false);
        A.store.Int64(ptr + 48, 0);

        A.store.Bool(ptr + 56 + 12, false);
        A.store.Ref(ptr + 56 + 0, undefined);
        A.store.Bool(ptr + 56 + 4, false);
        A.store.Enum(ptr + 56 + 8, -1);
        A.store.Bool(ptr + 123, false);
        A.store.Int64(ptr + 72, 0);
        A.store.Ref(ptr + 80, undefined);
        A.store.Bool(ptr + 84, false);
        A.store.Bool(ptr + 85, false);
        A.store.Ref(ptr + 88, undefined);
        A.store.Enum(ptr + 92, -1);
        A.store.Ref(ptr + 96, undefined);
        A.store.Ref(ptr + 100, undefined);
        A.store.Bool(ptr + 124, false);
        A.store.Int64(ptr + 104, 0);
        A.store.Int64(ptr + 112, 0);
      } else {
        A.store.Bool(ptr + 125, true);
        A.store.Bool(ptr + 0, x["active"] ? true : false);
        A.store.Bool(ptr + 120, "audible" in x ? true : false);
        A.store.Bool(ptr + 1, x["audible"] ? true : false);
        A.store.Bool(ptr + 2, x["autoDiscardable"] ? true : false);
        A.store.Bool(ptr + 3, x["discarded"] ? true : false);
        A.store.Ref(ptr + 4, x["favIconUrl"]);
        A.store.Int64(ptr + 8, x["groupId"] === undefined ? 0 : (x["groupId"] as number));
        A.store.Bool(ptr + 121, "height" in x ? true : false);
        A.store.Int64(ptr + 16, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Bool(ptr + 24, x["highlighted"] ? true : false);
        A.store.Bool(ptr + 122, "id" in x ? true : false);
        A.store.Int64(ptr + 32, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Bool(ptr + 40, x["incognito"] ? true : false);
        A.store.Int64(ptr + 48, x["index"] === undefined ? 0 : (x["index"] as number));

        if (typeof x["mutedInfo"] === "undefined") {
          A.store.Bool(ptr + 56 + 12, false);
          A.store.Ref(ptr + 56 + 0, undefined);
          A.store.Bool(ptr + 56 + 4, false);
          A.store.Enum(ptr + 56 + 8, -1);
        } else {
          A.store.Bool(ptr + 56 + 12, true);
          A.store.Ref(ptr + 56 + 0, x["mutedInfo"]["extensionId"]);
          A.store.Bool(ptr + 56 + 4, x["mutedInfo"]["muted"] ? true : false);
          A.store.Enum(ptr + 56 + 8, ["user", "capture", "extension"].indexOf(x["mutedInfo"]["reason"] as string));
        }
        A.store.Bool(ptr + 123, "openerTabId" in x ? true : false);
        A.store.Int64(ptr + 72, x["openerTabId"] === undefined ? 0 : (x["openerTabId"] as number));
        A.store.Ref(ptr + 80, x["pendingUrl"]);
        A.store.Bool(ptr + 84, x["pinned"] ? true : false);
        A.store.Bool(ptr + 85, x["selected"] ? true : false);
        A.store.Ref(ptr + 88, x["sessionId"]);
        A.store.Enum(ptr + 92, ["unloaded", "loading", "complete"].indexOf(x["status"] as string));
        A.store.Ref(ptr + 96, x["title"]);
        A.store.Ref(ptr + 100, x["url"]);
        A.store.Bool(ptr + 124, "width" in x ? true : false);
        A.store.Int64(ptr + 104, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Int64(ptr + 112, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_Tab": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["active"] = A.load.Bool(ptr + 0);
      if (A.load.Bool(ptr + 120)) {
        x["audible"] = A.load.Bool(ptr + 1);
      } else {
        delete x["audible"];
      }
      x["autoDiscardable"] = A.load.Bool(ptr + 2);
      x["discarded"] = A.load.Bool(ptr + 3);
      x["favIconUrl"] = A.load.Ref(ptr + 4, undefined);
      x["groupId"] = A.load.Int64(ptr + 8);
      if (A.load.Bool(ptr + 121)) {
        x["height"] = A.load.Int64(ptr + 16);
      } else {
        delete x["height"];
      }
      x["highlighted"] = A.load.Bool(ptr + 24);
      if (A.load.Bool(ptr + 122)) {
        x["id"] = A.load.Int64(ptr + 32);
      } else {
        delete x["id"];
      }
      x["incognito"] = A.load.Bool(ptr + 40);
      x["index"] = A.load.Int64(ptr + 48);
      if (A.load.Bool(ptr + 56 + 12)) {
        x["mutedInfo"] = {};
        x["mutedInfo"]["extensionId"] = A.load.Ref(ptr + 56 + 0, undefined);
        x["mutedInfo"]["muted"] = A.load.Bool(ptr + 56 + 4);
        x["mutedInfo"]["reason"] = A.load.Enum(ptr + 56 + 8, ["user", "capture", "extension"]);
      } else {
        delete x["mutedInfo"];
      }
      if (A.load.Bool(ptr + 123)) {
        x["openerTabId"] = A.load.Int64(ptr + 72);
      } else {
        delete x["openerTabId"];
      }
      x["pendingUrl"] = A.load.Ref(ptr + 80, undefined);
      x["pinned"] = A.load.Bool(ptr + 84);
      x["selected"] = A.load.Bool(ptr + 85);
      x["sessionId"] = A.load.Ref(ptr + 88, undefined);
      x["status"] = A.load.Enum(ptr + 92, ["unloaded", "loading", "complete"]);
      x["title"] = A.load.Ref(ptr + 96, undefined);
      x["url"] = A.load.Ref(ptr + 100, undefined);
      if (A.load.Bool(ptr + 124)) {
        x["width"] = A.load.Int64(ptr + 104);
      } else {
        delete x["width"];
      }
      x["windowId"] = A.load.Int64(ptr + 112);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CreateArgCreateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 46, false);
        A.store.Bool(ptr + 40, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 41, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Bool(ptr + 42, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Bool(ptr + 43, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 44, false);
        A.store.Bool(ptr + 25, false);
        A.store.Ref(ptr + 28, undefined);
        A.store.Bool(ptr + 45, false);
        A.store.Int64(ptr + 32, 0);
      } else {
        A.store.Bool(ptr + 46, true);
        A.store.Bool(ptr + 40, "active" in x ? true : false);
        A.store.Bool(ptr + 0, x["active"] ? true : false);
        A.store.Bool(ptr + 41, "index" in x ? true : false);
        A.store.Int64(ptr + 8, x["index"] === undefined ? 0 : (x["index"] as number));
        A.store.Bool(ptr + 42, "openerTabId" in x ? true : false);
        A.store.Int64(ptr + 16, x["openerTabId"] === undefined ? 0 : (x["openerTabId"] as number));
        A.store.Bool(ptr + 43, "pinned" in x ? true : false);
        A.store.Bool(ptr + 24, x["pinned"] ? true : false);
        A.store.Bool(ptr + 44, "selected" in x ? true : false);
        A.store.Bool(ptr + 25, x["selected"] ? true : false);
        A.store.Ref(ptr + 28, x["url"]);
        A.store.Bool(ptr + 45, "windowId" in x ? true : false);
        A.store.Int64(ptr + 32, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_CreateArgCreateProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 40)) {
        x["active"] = A.load.Bool(ptr + 0);
      } else {
        delete x["active"];
      }
      if (A.load.Bool(ptr + 41)) {
        x["index"] = A.load.Int64(ptr + 8);
      } else {
        delete x["index"];
      }
      if (A.load.Bool(ptr + 42)) {
        x["openerTabId"] = A.load.Int64(ptr + 16);
      } else {
        delete x["openerTabId"];
      }
      if (A.load.Bool(ptr + 43)) {
        x["pinned"] = A.load.Bool(ptr + 24);
      } else {
        delete x["pinned"];
      }
      if (A.load.Bool(ptr + 44)) {
        x["selected"] = A.load.Bool(ptr + 25);
      } else {
        delete x["selected"];
      }
      x["url"] = A.load.Ref(ptr + 28, undefined);
      if (A.load.Bool(ptr + 45)) {
        x["windowId"] = A.load.Int64(ptr + 32);
      } else {
        delete x["windowId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GroupArgOptionsFieldCreateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "windowId" in x ? true : false);
        A.store.Int64(ptr + 0, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_GroupArgOptionsFieldCreateProperties": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["windowId"] = A.load.Int64(ptr + 0);
      } else {
        delete x["windowId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GroupArgOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 29, false);

        A.store.Bool(ptr + 0 + 9, false);
        A.store.Bool(ptr + 0 + 8, false);
        A.store.Int64(ptr + 0 + 0, 0);
        A.store.Bool(ptr + 28, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 29, true);

        if (typeof x["createProperties"] === "undefined") {
          A.store.Bool(ptr + 0 + 9, false);
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Int64(ptr + 0 + 0, 0);
        } else {
          A.store.Bool(ptr + 0 + 9, true);
          A.store.Bool(ptr + 0 + 8, "windowId" in x["createProperties"] ? true : false);
          A.store.Int64(
            ptr + 0 + 0,
            x["createProperties"]["windowId"] === undefined ? 0 : (x["createProperties"]["windowId"] as number)
          );
        }
        A.store.Bool(ptr + 28, "groupId" in x ? true : false);
        A.store.Int64(ptr + 16, x["groupId"] === undefined ? 0 : (x["groupId"] as number));
        A.store.Ref(ptr + 24, x["tabIds"]);
      }
    },
    "load_GroupArgOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 9)) {
        x["createProperties"] = {};
        if (A.load.Bool(ptr + 0 + 8)) {
          x["createProperties"]["windowId"] = A.load.Int64(ptr + 0 + 0);
        } else {
          delete x["createProperties"]["windowId"];
        }
      } else {
        delete x["createProperties"];
      }
      if (A.load.Bool(ptr + 28)) {
        x["groupId"] = A.load.Int64(ptr + 16);
      } else {
        delete x["groupId"];
      }
      x["tabIds"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_HighlightArgHighlightInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["tabs"]);
        A.store.Bool(ptr + 16, "windowId" in x ? true : false);
        A.store.Int64(ptr + 8, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_HighlightArgHighlightInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["tabs"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["windowId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["windowId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "get_MAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "MAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND" in WEBEXT?.tabs) {
        const val = WEBEXT.tabs.MAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_MAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.tabs, "MAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND", A.H.get<object>(val), WEBEXT.tabs)
        ? A.H.TRUE
        : A.H.FALSE;
    },

    "store_MoveArgMoveProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Int64(ptr + 0, x["index"] === undefined ? 0 : (x["index"] as number));
        A.store.Bool(ptr + 16, "windowId" in x ? true : false);
        A.store.Int64(ptr + 8, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_MoveArgMoveProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["index"] = A.load.Int64(ptr + 0);
      if (A.load.Bool(ptr + 16)) {
        x["windowId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["windowId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnActivatedArgActiveInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Int64(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Int64(ptr + 8, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_OnActivatedArgActiveInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["tabId"] = A.load.Int64(ptr + 0);
      x["windowId"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnActiveChangedArgSelectInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Int64(ptr + 0, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_OnActiveChangedArgSelectInfo": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["windowId"] = A.load.Int64(ptr + 0);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnAttachedArgAttachInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Int64(ptr + 0, x["newPosition"] === undefined ? 0 : (x["newPosition"] as number));
        A.store.Int64(ptr + 8, x["newWindowId"] === undefined ? 0 : (x["newWindowId"] as number));
      }
    },
    "load_OnAttachedArgAttachInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["newPosition"] = A.load.Int64(ptr + 0);
      x["newWindowId"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnDetachedArgDetachInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Int64(ptr + 0, x["oldPosition"] === undefined ? 0 : (x["oldPosition"] as number));
        A.store.Int64(ptr + 8, x["oldWindowId"] === undefined ? 0 : (x["oldWindowId"] as number));
      }
    },
    "load_OnDetachedArgDetachInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["oldPosition"] = A.load.Int64(ptr + 0);
      x["oldWindowId"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnHighlightChangedArgSelectInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["tabIds"]);
        A.store.Int64(ptr + 8, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_OnHighlightChangedArgSelectInfo": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["tabIds"] = A.load.Ref(ptr + 0, undefined);
      x["windowId"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnHighlightedArgHighlightInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["tabIds"]);
        A.store.Int64(ptr + 8, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_OnHighlightedArgHighlightInfo": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["tabIds"] = A.load.Ref(ptr + 0, undefined);
      x["windowId"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnMovedArgMoveInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 24, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
        A.store.Int64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 24, true);
        A.store.Int64(ptr + 0, x["fromIndex"] === undefined ? 0 : (x["fromIndex"] as number));
        A.store.Int64(ptr + 8, x["toIndex"] === undefined ? 0 : (x["toIndex"] as number));
        A.store.Int64(ptr + 16, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_OnMovedArgMoveInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fromIndex"] = A.load.Int64(ptr + 0);
      x["toIndex"] = A.load.Int64(ptr + 8);
      x["windowId"] = A.load.Int64(ptr + 16);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnRemovedArgRemoveInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 0, false);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Bool(ptr + 0, x["isWindowClosing"] ? true : false);
        A.store.Int64(ptr + 8, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_OnRemovedArgRemoveInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["isWindowClosing"] = A.load.Bool(ptr + 0);
      x["windowId"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnSelectionChangedArgSelectInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Int64(ptr + 0, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_OnSelectionChangedArgSelectInfo": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["windowId"] = A.load.Int64(ptr + 0);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnUpdatedArgChangeInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 49, false);
        A.store.Bool(ptr + 44, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 45, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 46, false);
        A.store.Bool(ptr + 2, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 47, false);
        A.store.Int64(ptr + 8, 0);

        A.store.Bool(ptr + 16 + 12, false);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Bool(ptr + 16 + 4, false);
        A.store.Enum(ptr + 16 + 8, -1);
        A.store.Bool(ptr + 48, false);
        A.store.Bool(ptr + 29, false);
        A.store.Enum(ptr + 32, -1);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
      } else {
        A.store.Bool(ptr + 49, true);
        A.store.Bool(ptr + 44, "audible" in x ? true : false);
        A.store.Bool(ptr + 0, x["audible"] ? true : false);
        A.store.Bool(ptr + 45, "autoDiscardable" in x ? true : false);
        A.store.Bool(ptr + 1, x["autoDiscardable"] ? true : false);
        A.store.Bool(ptr + 46, "discarded" in x ? true : false);
        A.store.Bool(ptr + 2, x["discarded"] ? true : false);
        A.store.Ref(ptr + 4, x["favIconUrl"]);
        A.store.Bool(ptr + 47, "groupId" in x ? true : false);
        A.store.Int64(ptr + 8, x["groupId"] === undefined ? 0 : (x["groupId"] as number));

        if (typeof x["mutedInfo"] === "undefined") {
          A.store.Bool(ptr + 16 + 12, false);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Bool(ptr + 16 + 4, false);
          A.store.Enum(ptr + 16 + 8, -1);
        } else {
          A.store.Bool(ptr + 16 + 12, true);
          A.store.Ref(ptr + 16 + 0, x["mutedInfo"]["extensionId"]);
          A.store.Bool(ptr + 16 + 4, x["mutedInfo"]["muted"] ? true : false);
          A.store.Enum(ptr + 16 + 8, ["user", "capture", "extension"].indexOf(x["mutedInfo"]["reason"] as string));
        }
        A.store.Bool(ptr + 48, "pinned" in x ? true : false);
        A.store.Bool(ptr + 29, x["pinned"] ? true : false);
        A.store.Enum(ptr + 32, ["unloaded", "loading", "complete"].indexOf(x["status"] as string));
        A.store.Ref(ptr + 36, x["title"]);
        A.store.Ref(ptr + 40, x["url"]);
      }
    },
    "load_OnUpdatedArgChangeInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 44)) {
        x["audible"] = A.load.Bool(ptr + 0);
      } else {
        delete x["audible"];
      }
      if (A.load.Bool(ptr + 45)) {
        x["autoDiscardable"] = A.load.Bool(ptr + 1);
      } else {
        delete x["autoDiscardable"];
      }
      if (A.load.Bool(ptr + 46)) {
        x["discarded"] = A.load.Bool(ptr + 2);
      } else {
        delete x["discarded"];
      }
      x["favIconUrl"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 47)) {
        x["groupId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["groupId"];
      }
      if (A.load.Bool(ptr + 16 + 12)) {
        x["mutedInfo"] = {};
        x["mutedInfo"]["extensionId"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["mutedInfo"]["muted"] = A.load.Bool(ptr + 16 + 4);
        x["mutedInfo"]["reason"] = A.load.Enum(ptr + 16 + 8, ["user", "capture", "extension"]);
      } else {
        delete x["mutedInfo"];
      }
      if (A.load.Bool(ptr + 48)) {
        x["pinned"] = A.load.Bool(ptr + 29);
      } else {
        delete x["pinned"];
      }
      x["status"] = A.load.Enum(ptr + 32, ["unloaded", "loading", "complete"]);
      x["title"] = A.load.Ref(ptr + 36, undefined);
      x["url"] = A.load.Ref(ptr + 40, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ZoomSettingsMode": (ref: heap.Ref<string>): number => {
      const idx = ["automatic", "manual", "disabled"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ZoomSettingsScope": (ref: heap.Ref<string>): number => {
      const idx = ["per-origin", "per-tab"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ZoomSettings": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Enum(ptr + 8, -1);
        A.store.Enum(ptr + 12, -1);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Bool(ptr + 16, "defaultZoomFactor" in x ? true : false);
        A.store.Float64(ptr + 0, x["defaultZoomFactor"] === undefined ? 0 : (x["defaultZoomFactor"] as number));
        A.store.Enum(ptr + 8, ["automatic", "manual", "disabled"].indexOf(x["mode"] as string));
        A.store.Enum(ptr + 12, ["per-origin", "per-tab"].indexOf(x["scope"] as string));
      }
    },
    "load_ZoomSettings": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["defaultZoomFactor"] = A.load.Float64(ptr + 0);
      } else {
        delete x["defaultZoomFactor"];
      }
      x["mode"] = A.load.Enum(ptr + 8, ["automatic", "manual", "disabled"]);
      x["scope"] = A.load.Enum(ptr + 12, ["per-origin", "per-tab"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnZoomChangeArgZoomChangeInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 42, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Float64(ptr + 8, 0);
        A.store.Int64(ptr + 16, 0);

        A.store.Bool(ptr + 24 + 17, false);
        A.store.Bool(ptr + 24 + 16, false);
        A.store.Float64(ptr + 24 + 0, 0);
        A.store.Enum(ptr + 24 + 8, -1);
        A.store.Enum(ptr + 24 + 12, -1);
      } else {
        A.store.Bool(ptr + 42, true);
        A.store.Float64(ptr + 0, x["newZoomFactor"] === undefined ? 0 : (x["newZoomFactor"] as number));
        A.store.Float64(ptr + 8, x["oldZoomFactor"] === undefined ? 0 : (x["oldZoomFactor"] as number));
        A.store.Int64(ptr + 16, x["tabId"] === undefined ? 0 : (x["tabId"] as number));

        if (typeof x["zoomSettings"] === "undefined") {
          A.store.Bool(ptr + 24 + 17, false);
          A.store.Bool(ptr + 24 + 16, false);
          A.store.Float64(ptr + 24 + 0, 0);
          A.store.Enum(ptr + 24 + 8, -1);
          A.store.Enum(ptr + 24 + 12, -1);
        } else {
          A.store.Bool(ptr + 24 + 17, true);
          A.store.Bool(ptr + 24 + 16, "defaultZoomFactor" in x["zoomSettings"] ? true : false);
          A.store.Float64(
            ptr + 24 + 0,
            x["zoomSettings"]["defaultZoomFactor"] === undefined
              ? 0
              : (x["zoomSettings"]["defaultZoomFactor"] as number)
          );
          A.store.Enum(ptr + 24 + 8, ["automatic", "manual", "disabled"].indexOf(x["zoomSettings"]["mode"] as string));
          A.store.Enum(ptr + 24 + 12, ["per-origin", "per-tab"].indexOf(x["zoomSettings"]["scope"] as string));
        }
      }
    },
    "load_OnZoomChangeArgZoomChangeInfo": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["newZoomFactor"] = A.load.Float64(ptr + 0);
      x["oldZoomFactor"] = A.load.Float64(ptr + 8);
      x["tabId"] = A.load.Int64(ptr + 16);
      if (A.load.Bool(ptr + 24 + 17)) {
        x["zoomSettings"] = {};
        if (A.load.Bool(ptr + 24 + 16)) {
          x["zoomSettings"]["defaultZoomFactor"] = A.load.Float64(ptr + 24 + 0);
        } else {
          delete x["zoomSettings"]["defaultZoomFactor"];
        }
        x["zoomSettings"]["mode"] = A.load.Enum(ptr + 24 + 8, ["automatic", "manual", "disabled"]);
        x["zoomSettings"]["scope"] = A.load.Enum(ptr + 24 + 12, ["per-origin", "per-tab"]);
      } else {
        delete x["zoomSettings"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_WindowType": (ref: heap.Ref<string>): number => {
      const idx = ["normal", "popup", "panel", "app", "devtools"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_QueryArgQueryInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 72, false);
        A.store.Bool(ptr + 60, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 61, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 62, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 63, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 64, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 65, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Bool(ptr + 66, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 67, false);
        A.store.Int64(ptr + 24, 0);
        A.store.Bool(ptr + 68, false);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 69, false);
        A.store.Bool(ptr + 33, false);
        A.store.Bool(ptr + 70, false);
        A.store.Bool(ptr + 34, false);
        A.store.Enum(ptr + 36, -1);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Bool(ptr + 71, false);
        A.store.Int64(ptr + 48, 0);
        A.store.Enum(ptr + 56, -1);
      } else {
        A.store.Bool(ptr + 72, true);
        A.store.Bool(ptr + 60, "active" in x ? true : false);
        A.store.Bool(ptr + 0, x["active"] ? true : false);
        A.store.Bool(ptr + 61, "audible" in x ? true : false);
        A.store.Bool(ptr + 1, x["audible"] ? true : false);
        A.store.Bool(ptr + 62, "autoDiscardable" in x ? true : false);
        A.store.Bool(ptr + 2, x["autoDiscardable"] ? true : false);
        A.store.Bool(ptr + 63, "currentWindow" in x ? true : false);
        A.store.Bool(ptr + 3, x["currentWindow"] ? true : false);
        A.store.Bool(ptr + 64, "discarded" in x ? true : false);
        A.store.Bool(ptr + 4, x["discarded"] ? true : false);
        A.store.Bool(ptr + 65, "groupId" in x ? true : false);
        A.store.Int64(ptr + 8, x["groupId"] === undefined ? 0 : (x["groupId"] as number));
        A.store.Bool(ptr + 66, "highlighted" in x ? true : false);
        A.store.Bool(ptr + 16, x["highlighted"] ? true : false);
        A.store.Bool(ptr + 67, "index" in x ? true : false);
        A.store.Int64(ptr + 24, x["index"] === undefined ? 0 : (x["index"] as number));
        A.store.Bool(ptr + 68, "lastFocusedWindow" in x ? true : false);
        A.store.Bool(ptr + 32, x["lastFocusedWindow"] ? true : false);
        A.store.Bool(ptr + 69, "muted" in x ? true : false);
        A.store.Bool(ptr + 33, x["muted"] ? true : false);
        A.store.Bool(ptr + 70, "pinned" in x ? true : false);
        A.store.Bool(ptr + 34, x["pinned"] ? true : false);
        A.store.Enum(ptr + 36, ["unloaded", "loading", "complete"].indexOf(x["status"] as string));
        A.store.Ref(ptr + 40, x["title"]);
        A.store.Ref(ptr + 44, x["url"]);
        A.store.Bool(ptr + 71, "windowId" in x ? true : false);
        A.store.Int64(ptr + 48, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
        A.store.Enum(ptr + 56, ["normal", "popup", "panel", "app", "devtools"].indexOf(x["windowType"] as string));
      }
    },
    "load_QueryArgQueryInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 60)) {
        x["active"] = A.load.Bool(ptr + 0);
      } else {
        delete x["active"];
      }
      if (A.load.Bool(ptr + 61)) {
        x["audible"] = A.load.Bool(ptr + 1);
      } else {
        delete x["audible"];
      }
      if (A.load.Bool(ptr + 62)) {
        x["autoDiscardable"] = A.load.Bool(ptr + 2);
      } else {
        delete x["autoDiscardable"];
      }
      if (A.load.Bool(ptr + 63)) {
        x["currentWindow"] = A.load.Bool(ptr + 3);
      } else {
        delete x["currentWindow"];
      }
      if (A.load.Bool(ptr + 64)) {
        x["discarded"] = A.load.Bool(ptr + 4);
      } else {
        delete x["discarded"];
      }
      if (A.load.Bool(ptr + 65)) {
        x["groupId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["groupId"];
      }
      if (A.load.Bool(ptr + 66)) {
        x["highlighted"] = A.load.Bool(ptr + 16);
      } else {
        delete x["highlighted"];
      }
      if (A.load.Bool(ptr + 67)) {
        x["index"] = A.load.Int64(ptr + 24);
      } else {
        delete x["index"];
      }
      if (A.load.Bool(ptr + 68)) {
        x["lastFocusedWindow"] = A.load.Bool(ptr + 32);
      } else {
        delete x["lastFocusedWindow"];
      }
      if (A.load.Bool(ptr + 69)) {
        x["muted"] = A.load.Bool(ptr + 33);
      } else {
        delete x["muted"];
      }
      if (A.load.Bool(ptr + 70)) {
        x["pinned"] = A.load.Bool(ptr + 34);
      } else {
        delete x["pinned"];
      }
      x["status"] = A.load.Enum(ptr + 36, ["unloaded", "loading", "complete"]);
      x["title"] = A.load.Ref(ptr + 40, undefined);
      x["url"] = A.load.Ref(ptr + 44, undefined);
      if (A.load.Bool(ptr + 71)) {
        x["windowId"] = A.load.Int64(ptr + 48);
      } else {
        delete x["windowId"];
      }
      x["windowType"] = A.load.Enum(ptr + 56, ["normal", "popup", "panel", "app", "devtools"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ReloadArgReloadProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 1, "bypassCache" in x ? true : false);
        A.store.Bool(ptr + 0, x["bypassCache"] ? true : false);
      }
    },
    "load_ReloadArgReloadProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 1)) {
        x["bypassCache"] = A.load.Bool(ptr + 0);
      } else {
        delete x["bypassCache"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SendMessageArgOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Bool(ptr + 16, "frameId" in x ? true : false);
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
      }
    },
    "load_SendMessageArgOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["frameId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["frameId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "get_TAB_ID_NONE": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "TAB_ID_NONE" in WEBEXT?.tabs) {
        const val = WEBEXT.tabs.TAB_ID_NONE;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_TAB_ID_NONE": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.tabs, "TAB_ID_NONE", A.H.get<object>(val), WEBEXT.tabs) ? A.H.TRUE : A.H.FALSE;
    },

    "store_UpdateArgUpdateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 31, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 27, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 28, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Bool(ptr + 29, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 30, false);
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 20, undefined);
      } else {
        A.store.Bool(ptr + 31, true);
        A.store.Bool(ptr + 24, "active" in x ? true : false);
        A.store.Bool(ptr + 0, x["active"] ? true : false);
        A.store.Bool(ptr + 25, "autoDiscardable" in x ? true : false);
        A.store.Bool(ptr + 1, x["autoDiscardable"] ? true : false);
        A.store.Bool(ptr + 26, "highlighted" in x ? true : false);
        A.store.Bool(ptr + 2, x["highlighted"] ? true : false);
        A.store.Bool(ptr + 27, "muted" in x ? true : false);
        A.store.Bool(ptr + 3, x["muted"] ? true : false);
        A.store.Bool(ptr + 28, "openerTabId" in x ? true : false);
        A.store.Int64(ptr + 8, x["openerTabId"] === undefined ? 0 : (x["openerTabId"] as number));
        A.store.Bool(ptr + 29, "pinned" in x ? true : false);
        A.store.Bool(ptr + 16, x["pinned"] ? true : false);
        A.store.Bool(ptr + 30, "selected" in x ? true : false);
        A.store.Bool(ptr + 17, x["selected"] ? true : false);
        A.store.Ref(ptr + 20, x["url"]);
      }
    },
    "load_UpdateArgUpdateProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["active"] = A.load.Bool(ptr + 0);
      } else {
        delete x["active"];
      }
      if (A.load.Bool(ptr + 25)) {
        x["autoDiscardable"] = A.load.Bool(ptr + 1);
      } else {
        delete x["autoDiscardable"];
      }
      if (A.load.Bool(ptr + 26)) {
        x["highlighted"] = A.load.Bool(ptr + 2);
      } else {
        delete x["highlighted"];
      }
      if (A.load.Bool(ptr + 27)) {
        x["muted"] = A.load.Bool(ptr + 3);
      } else {
        delete x["muted"];
      }
      if (A.load.Bool(ptr + 28)) {
        x["openerTabId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["openerTabId"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["pinned"] = A.load.Bool(ptr + 16);
      } else {
        delete x["pinned"];
      }
      if (A.load.Bool(ptr + 30)) {
        x["selected"] = A.load.Bool(ptr + 17);
      } else {
        delete x["selected"];
      }
      x["url"] = A.load.Ref(ptr + 20, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_CaptureVisibleTab": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "captureVisibleTab" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CaptureVisibleTab": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.captureVisibleTab);
    },
    "call_CaptureVisibleTab": (retPtr: Pointer, windowId: number, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["format"] = A.load.Enum(options + 0, ["jpeg", "png"]);
      if (A.load.Bool(options + 16)) {
        options_ffi["quality"] = A.load.Int64(options + 8);
      }

      const _ret = WEBEXT.tabs.captureVisibleTab(windowId, options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_CaptureVisibleTab": (
      retPtr: Pointer,
      errPtr: Pointer,
      windowId: number,
      options: Pointer
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["format"] = A.load.Enum(options + 0, ["jpeg", "png"]);
        if (A.load.Bool(options + 16)) {
          options_ffi["quality"] = A.load.Int64(options + 8);
        }

        const _ret = WEBEXT.tabs.captureVisibleTab(windowId, options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Create": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "create" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Create": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.create);
    },
    "call_Create": (retPtr: Pointer, createProperties: Pointer): void => {
      const createProperties_ffi = {};

      if (A.load.Bool(createProperties + 40)) {
        createProperties_ffi["active"] = A.load.Bool(createProperties + 0);
      }
      if (A.load.Bool(createProperties + 41)) {
        createProperties_ffi["index"] = A.load.Int64(createProperties + 8);
      }
      if (A.load.Bool(createProperties + 42)) {
        createProperties_ffi["openerTabId"] = A.load.Int64(createProperties + 16);
      }
      if (A.load.Bool(createProperties + 43)) {
        createProperties_ffi["pinned"] = A.load.Bool(createProperties + 24);
      }
      if (A.load.Bool(createProperties + 44)) {
        createProperties_ffi["selected"] = A.load.Bool(createProperties + 25);
      }
      createProperties_ffi["url"] = A.load.Ref(createProperties + 28, undefined);
      if (A.load.Bool(createProperties + 45)) {
        createProperties_ffi["windowId"] = A.load.Int64(createProperties + 32);
      }

      const _ret = WEBEXT.tabs.create(createProperties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Create": (retPtr: Pointer, errPtr: Pointer, createProperties: Pointer): heap.Ref<boolean> => {
      try {
        const createProperties_ffi = {};

        if (A.load.Bool(createProperties + 40)) {
          createProperties_ffi["active"] = A.load.Bool(createProperties + 0);
        }
        if (A.load.Bool(createProperties + 41)) {
          createProperties_ffi["index"] = A.load.Int64(createProperties + 8);
        }
        if (A.load.Bool(createProperties + 42)) {
          createProperties_ffi["openerTabId"] = A.load.Int64(createProperties + 16);
        }
        if (A.load.Bool(createProperties + 43)) {
          createProperties_ffi["pinned"] = A.load.Bool(createProperties + 24);
        }
        if (A.load.Bool(createProperties + 44)) {
          createProperties_ffi["selected"] = A.load.Bool(createProperties + 25);
        }
        createProperties_ffi["url"] = A.load.Ref(createProperties + 28, undefined);
        if (A.load.Bool(createProperties + 45)) {
          createProperties_ffi["windowId"] = A.load.Int64(createProperties + 32);
        }

        const _ret = WEBEXT.tabs.create(createProperties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DetectLanguage": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "detectLanguage" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DetectLanguage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.detectLanguage);
    },
    "call_DetectLanguage": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.tabs.detectLanguage(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_DetectLanguage": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.detectLanguage(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Discard": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "discard" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Discard": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.discard);
    },
    "call_Discard": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.tabs.discard(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Discard": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.discard(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Duplicate": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "duplicate" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Duplicate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.duplicate);
    },
    "call_Duplicate": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.tabs.duplicate(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Duplicate": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.duplicate(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ExecuteScript": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "executeScript" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ExecuteScript": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.executeScript);
    },
    "call_ExecuteScript": (retPtr: Pointer, tabId: number, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 32)) {
        details_ffi["allFrames"] = A.load.Bool(details + 0);
      }
      details_ffi["code"] = A.load.Ref(details + 4, undefined);
      details_ffi["cssOrigin"] = A.load.Enum(details + 8, ["author", "user"]);
      details_ffi["file"] = A.load.Ref(details + 12, undefined);
      if (A.load.Bool(details + 33)) {
        details_ffi["frameId"] = A.load.Int64(details + 16);
      }
      if (A.load.Bool(details + 34)) {
        details_ffi["matchAboutBlank"] = A.load.Bool(details + 24);
      }
      details_ffi["runAt"] = A.load.Enum(details + 28, ["document_start", "document_end", "document_idle"]);

      const _ret = WEBEXT.tabs.executeScript(tabId, details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ExecuteScript": (retPtr: Pointer, errPtr: Pointer, tabId: number, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 32)) {
          details_ffi["allFrames"] = A.load.Bool(details + 0);
        }
        details_ffi["code"] = A.load.Ref(details + 4, undefined);
        details_ffi["cssOrigin"] = A.load.Enum(details + 8, ["author", "user"]);
        details_ffi["file"] = A.load.Ref(details + 12, undefined);
        if (A.load.Bool(details + 33)) {
          details_ffi["frameId"] = A.load.Int64(details + 16);
        }
        if (A.load.Bool(details + 34)) {
          details_ffi["matchAboutBlank"] = A.load.Bool(details + 24);
        }
        details_ffi["runAt"] = A.load.Enum(details + 28, ["document_start", "document_end", "document_idle"]);

        const _ret = WEBEXT.tabs.executeScript(tabId, details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Get": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "get" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Get": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.get);
    },
    "call_Get": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.tabs.get(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Get": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.get(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAllInWindow": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "getAllInWindow" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAllInWindow": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.getAllInWindow);
    },
    "call_GetAllInWindow": (retPtr: Pointer, windowId: number): void => {
      const _ret = WEBEXT.tabs.getAllInWindow(windowId);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAllInWindow": (retPtr: Pointer, errPtr: Pointer, windowId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.getAllInWindow(windowId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCurrent": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "getCurrent" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCurrent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.getCurrent);
    },
    "call_GetCurrent": (retPtr: Pointer): void => {
      const _ret = WEBEXT.tabs.getCurrent();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCurrent": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.getCurrent();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSelected": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "getSelected" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSelected": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.getSelected);
    },
    "call_GetSelected": (retPtr: Pointer, windowId: number): void => {
      const _ret = WEBEXT.tabs.getSelected(windowId);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSelected": (retPtr: Pointer, errPtr: Pointer, windowId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.getSelected(windowId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetZoom": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "getZoom" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetZoom": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.getZoom);
    },
    "call_GetZoom": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.tabs.getZoom(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetZoom": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.getZoom(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetZoomSettings": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "getZoomSettings" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetZoomSettings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.getZoomSettings);
    },
    "call_GetZoomSettings": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.tabs.getZoomSettings(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetZoomSettings": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.getZoomSettings(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GoBack": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "goBack" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GoBack": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.goBack);
    },
    "call_GoBack": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.tabs.goBack(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_GoBack": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.goBack(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GoForward": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "goForward" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GoForward": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.goForward);
    },
    "call_GoForward": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.tabs.goForward(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_GoForward": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.goForward(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Group": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "group" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Group": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.group);
    },
    "call_Group": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 0 + 9)) {
        options_ffi["createProperties"] = {};
        if (A.load.Bool(options + 0 + 8)) {
          options_ffi["createProperties"]["windowId"] = A.load.Int64(options + 0 + 0);
        }
      }
      if (A.load.Bool(options + 28)) {
        options_ffi["groupId"] = A.load.Int64(options + 16);
      }
      options_ffi["tabIds"] = A.load.Ref(options + 24, undefined);

      const _ret = WEBEXT.tabs.group(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Group": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 0 + 9)) {
          options_ffi["createProperties"] = {};
          if (A.load.Bool(options + 0 + 8)) {
            options_ffi["createProperties"]["windowId"] = A.load.Int64(options + 0 + 0);
          }
        }
        if (A.load.Bool(options + 28)) {
          options_ffi["groupId"] = A.load.Int64(options + 16);
        }
        options_ffi["tabIds"] = A.load.Ref(options + 24, undefined);

        const _ret = WEBEXT.tabs.group(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InsertCSS": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "insertCSS" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InsertCSS": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.insertCSS);
    },
    "call_InsertCSS": (retPtr: Pointer, tabId: number, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 32)) {
        details_ffi["allFrames"] = A.load.Bool(details + 0);
      }
      details_ffi["code"] = A.load.Ref(details + 4, undefined);
      details_ffi["cssOrigin"] = A.load.Enum(details + 8, ["author", "user"]);
      details_ffi["file"] = A.load.Ref(details + 12, undefined);
      if (A.load.Bool(details + 33)) {
        details_ffi["frameId"] = A.load.Int64(details + 16);
      }
      if (A.load.Bool(details + 34)) {
        details_ffi["matchAboutBlank"] = A.load.Bool(details + 24);
      }
      details_ffi["runAt"] = A.load.Enum(details + 28, ["document_start", "document_end", "document_idle"]);

      const _ret = WEBEXT.tabs.insertCSS(tabId, details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_InsertCSS": (retPtr: Pointer, errPtr: Pointer, tabId: number, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 32)) {
          details_ffi["allFrames"] = A.load.Bool(details + 0);
        }
        details_ffi["code"] = A.load.Ref(details + 4, undefined);
        details_ffi["cssOrigin"] = A.load.Enum(details + 8, ["author", "user"]);
        details_ffi["file"] = A.load.Ref(details + 12, undefined);
        if (A.load.Bool(details + 33)) {
          details_ffi["frameId"] = A.load.Int64(details + 16);
        }
        if (A.load.Bool(details + 34)) {
          details_ffi["matchAboutBlank"] = A.load.Bool(details + 24);
        }
        details_ffi["runAt"] = A.load.Enum(details + 28, ["document_start", "document_end", "document_idle"]);

        const _ret = WEBEXT.tabs.insertCSS(tabId, details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Move": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "move" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Move": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.move);
    },
    "call_Move": (retPtr: Pointer, tabIds: heap.Ref<any>, moveProperties: Pointer): void => {
      const moveProperties_ffi = {};

      moveProperties_ffi["index"] = A.load.Int64(moveProperties + 0);
      if (A.load.Bool(moveProperties + 16)) {
        moveProperties_ffi["windowId"] = A.load.Int64(moveProperties + 8);
      }

      const _ret = WEBEXT.tabs.move(A.H.get<any>(tabIds), moveProperties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Move": (
      retPtr: Pointer,
      errPtr: Pointer,
      tabIds: heap.Ref<any>,
      moveProperties: Pointer
    ): heap.Ref<boolean> => {
      try {
        const moveProperties_ffi = {};

        moveProperties_ffi["index"] = A.load.Int64(moveProperties + 0);
        if (A.load.Bool(moveProperties + 16)) {
          moveProperties_ffi["windowId"] = A.load.Int64(moveProperties + 8);
        }

        const _ret = WEBEXT.tabs.move(A.H.get<any>(tabIds), moveProperties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnActivated": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onActivated && "addListener" in WEBEXT?.tabs?.onActivated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnActivated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onActivated.addListener);
    },
    "call_OnActivated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onActivated.addListener(A.H.get<object>(callback));
    },
    "try_OnActivated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onActivated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffActivated": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onActivated && "removeListener" in WEBEXT?.tabs?.onActivated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffActivated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onActivated.removeListener);
    },
    "call_OffActivated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onActivated.removeListener(A.H.get<object>(callback));
    },
    "try_OffActivated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onActivated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnActivated": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onActivated && "hasListener" in WEBEXT?.tabs?.onActivated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnActivated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onActivated.hasListener);
    },
    "call_HasOnActivated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onActivated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnActivated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onActivated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnActiveChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onActiveChanged && "addListener" in WEBEXT?.tabs?.onActiveChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnActiveChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onActiveChanged.addListener);
    },
    "call_OnActiveChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onActiveChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnActiveChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onActiveChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffActiveChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onActiveChanged && "removeListener" in WEBEXT?.tabs?.onActiveChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffActiveChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onActiveChanged.removeListener);
    },
    "call_OffActiveChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onActiveChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffActiveChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onActiveChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnActiveChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onActiveChanged && "hasListener" in WEBEXT?.tabs?.onActiveChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnActiveChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onActiveChanged.hasListener);
    },
    "call_HasOnActiveChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onActiveChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnActiveChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onActiveChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAttached": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onAttached && "addListener" in WEBEXT?.tabs?.onAttached) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAttached": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onAttached.addListener);
    },
    "call_OnAttached": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onAttached.addListener(A.H.get<object>(callback));
    },
    "try_OnAttached": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onAttached.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAttached": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onAttached && "removeListener" in WEBEXT?.tabs?.onAttached) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAttached": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onAttached.removeListener);
    },
    "call_OffAttached": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onAttached.removeListener(A.H.get<object>(callback));
    },
    "try_OffAttached": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onAttached.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAttached": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onAttached && "hasListener" in WEBEXT?.tabs?.onAttached) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAttached": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onAttached.hasListener);
    },
    "call_HasOnAttached": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onAttached.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAttached": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onAttached.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onCreated && "addListener" in WEBEXT?.tabs?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onCreated.addListener);
    },
    "call_OnCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onCreated.addListener(A.H.get<object>(callback));
    },
    "try_OnCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onCreated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onCreated && "removeListener" in WEBEXT?.tabs?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onCreated.removeListener);
    },
    "call_OffCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onCreated.removeListener(A.H.get<object>(callback));
    },
    "try_OffCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onCreated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onCreated && "hasListener" in WEBEXT?.tabs?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onCreated.hasListener);
    },
    "call_HasOnCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onCreated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onCreated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDetached": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onDetached && "addListener" in WEBEXT?.tabs?.onDetached) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDetached": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onDetached.addListener);
    },
    "call_OnDetached": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onDetached.addListener(A.H.get<object>(callback));
    },
    "try_OnDetached": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onDetached.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDetached": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onDetached && "removeListener" in WEBEXT?.tabs?.onDetached) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDetached": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onDetached.removeListener);
    },
    "call_OffDetached": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onDetached.removeListener(A.H.get<object>(callback));
    },
    "try_OffDetached": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onDetached.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDetached": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onDetached && "hasListener" in WEBEXT?.tabs?.onDetached) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDetached": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onDetached.hasListener);
    },
    "call_HasOnDetached": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onDetached.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDetached": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onDetached.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnHighlightChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onHighlightChanged && "addListener" in WEBEXT?.tabs?.onHighlightChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnHighlightChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onHighlightChanged.addListener);
    },
    "call_OnHighlightChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onHighlightChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnHighlightChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onHighlightChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffHighlightChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onHighlightChanged && "removeListener" in WEBEXT?.tabs?.onHighlightChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffHighlightChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onHighlightChanged.removeListener);
    },
    "call_OffHighlightChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onHighlightChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffHighlightChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onHighlightChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnHighlightChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onHighlightChanged && "hasListener" in WEBEXT?.tabs?.onHighlightChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnHighlightChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onHighlightChanged.hasListener);
    },
    "call_HasOnHighlightChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onHighlightChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnHighlightChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onHighlightChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnHighlighted": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onHighlighted && "addListener" in WEBEXT?.tabs?.onHighlighted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnHighlighted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onHighlighted.addListener);
    },
    "call_OnHighlighted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onHighlighted.addListener(A.H.get<object>(callback));
    },
    "try_OnHighlighted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onHighlighted.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffHighlighted": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onHighlighted && "removeListener" in WEBEXT?.tabs?.onHighlighted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffHighlighted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onHighlighted.removeListener);
    },
    "call_OffHighlighted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onHighlighted.removeListener(A.H.get<object>(callback));
    },
    "try_OffHighlighted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onHighlighted.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnHighlighted": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onHighlighted && "hasListener" in WEBEXT?.tabs?.onHighlighted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnHighlighted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onHighlighted.hasListener);
    },
    "call_HasOnHighlighted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onHighlighted.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnHighlighted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onHighlighted.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onMoved && "addListener" in WEBEXT?.tabs?.onMoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onMoved.addListener);
    },
    "call_OnMoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onMoved.addListener(A.H.get<object>(callback));
    },
    "try_OnMoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onMoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onMoved && "removeListener" in WEBEXT?.tabs?.onMoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onMoved.removeListener);
    },
    "call_OffMoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onMoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffMoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onMoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onMoved && "hasListener" in WEBEXT?.tabs?.onMoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onMoved.hasListener);
    },
    "call_HasOnMoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onMoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onMoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onRemoved && "addListener" in WEBEXT?.tabs?.onRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onRemoved.addListener);
    },
    "call_OnRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onRemoved && "removeListener" in WEBEXT?.tabs?.onRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onRemoved.removeListener);
    },
    "call_OffRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onRemoved && "hasListener" in WEBEXT?.tabs?.onRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onRemoved.hasListener);
    },
    "call_HasOnRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnReplaced": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onReplaced && "addListener" in WEBEXT?.tabs?.onReplaced) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnReplaced": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onReplaced.addListener);
    },
    "call_OnReplaced": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onReplaced.addListener(A.H.get<object>(callback));
    },
    "try_OnReplaced": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onReplaced.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffReplaced": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onReplaced && "removeListener" in WEBEXT?.tabs?.onReplaced) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffReplaced": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onReplaced.removeListener);
    },
    "call_OffReplaced": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onReplaced.removeListener(A.H.get<object>(callback));
    },
    "try_OffReplaced": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onReplaced.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnReplaced": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onReplaced && "hasListener" in WEBEXT?.tabs?.onReplaced) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnReplaced": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onReplaced.hasListener);
    },
    "call_HasOnReplaced": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onReplaced.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnReplaced": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onReplaced.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSelectionChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onSelectionChanged && "addListener" in WEBEXT?.tabs?.onSelectionChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSelectionChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onSelectionChanged.addListener);
    },
    "call_OnSelectionChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onSelectionChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnSelectionChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onSelectionChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSelectionChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onSelectionChanged && "removeListener" in WEBEXT?.tabs?.onSelectionChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSelectionChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onSelectionChanged.removeListener);
    },
    "call_OffSelectionChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onSelectionChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffSelectionChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onSelectionChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSelectionChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onSelectionChanged && "hasListener" in WEBEXT?.tabs?.onSelectionChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSelectionChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onSelectionChanged.hasListener);
    },
    "call_HasOnSelectionChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onSelectionChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSelectionChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onSelectionChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnUpdated": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onUpdated && "addListener" in WEBEXT?.tabs?.onUpdated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onUpdated.addListener);
    },
    "call_OnUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onUpdated.addListener(A.H.get<object>(callback));
    },
    "try_OnUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onUpdated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffUpdated": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onUpdated && "removeListener" in WEBEXT?.tabs?.onUpdated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onUpdated.removeListener);
    },
    "call_OffUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onUpdated.removeListener(A.H.get<object>(callback));
    },
    "try_OffUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onUpdated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnUpdated": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onUpdated && "hasListener" in WEBEXT?.tabs?.onUpdated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onUpdated.hasListener);
    },
    "call_HasOnUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onUpdated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onUpdated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnZoomChange": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onZoomChange && "addListener" in WEBEXT?.tabs?.onZoomChange) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnZoomChange": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onZoomChange.addListener);
    },
    "call_OnZoomChange": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onZoomChange.addListener(A.H.get<object>(callback));
    },
    "try_OnZoomChange": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onZoomChange.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffZoomChange": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onZoomChange && "removeListener" in WEBEXT?.tabs?.onZoomChange) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffZoomChange": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onZoomChange.removeListener);
    },
    "call_OffZoomChange": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onZoomChange.removeListener(A.H.get<object>(callback));
    },
    "try_OffZoomChange": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onZoomChange.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnZoomChange": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs?.onZoomChange && "hasListener" in WEBEXT?.tabs?.onZoomChange) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnZoomChange": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.onZoomChange.hasListener);
    },
    "call_HasOnZoomChange": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.onZoomChange.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnZoomChange": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.onZoomChange.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Query": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "query" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Query": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.query);
    },
    "call_Query": (retPtr: Pointer, queryInfo: Pointer): void => {
      const queryInfo_ffi = {};

      if (A.load.Bool(queryInfo + 60)) {
        queryInfo_ffi["active"] = A.load.Bool(queryInfo + 0);
      }
      if (A.load.Bool(queryInfo + 61)) {
        queryInfo_ffi["audible"] = A.load.Bool(queryInfo + 1);
      }
      if (A.load.Bool(queryInfo + 62)) {
        queryInfo_ffi["autoDiscardable"] = A.load.Bool(queryInfo + 2);
      }
      if (A.load.Bool(queryInfo + 63)) {
        queryInfo_ffi["currentWindow"] = A.load.Bool(queryInfo + 3);
      }
      if (A.load.Bool(queryInfo + 64)) {
        queryInfo_ffi["discarded"] = A.load.Bool(queryInfo + 4);
      }
      if (A.load.Bool(queryInfo + 65)) {
        queryInfo_ffi["groupId"] = A.load.Int64(queryInfo + 8);
      }
      if (A.load.Bool(queryInfo + 66)) {
        queryInfo_ffi["highlighted"] = A.load.Bool(queryInfo + 16);
      }
      if (A.load.Bool(queryInfo + 67)) {
        queryInfo_ffi["index"] = A.load.Int64(queryInfo + 24);
      }
      if (A.load.Bool(queryInfo + 68)) {
        queryInfo_ffi["lastFocusedWindow"] = A.load.Bool(queryInfo + 32);
      }
      if (A.load.Bool(queryInfo + 69)) {
        queryInfo_ffi["muted"] = A.load.Bool(queryInfo + 33);
      }
      if (A.load.Bool(queryInfo + 70)) {
        queryInfo_ffi["pinned"] = A.load.Bool(queryInfo + 34);
      }
      queryInfo_ffi["status"] = A.load.Enum(queryInfo + 36, ["unloaded", "loading", "complete"]);
      queryInfo_ffi["title"] = A.load.Ref(queryInfo + 40, undefined);
      queryInfo_ffi["url"] = A.load.Ref(queryInfo + 44, undefined);
      if (A.load.Bool(queryInfo + 71)) {
        queryInfo_ffi["windowId"] = A.load.Int64(queryInfo + 48);
      }
      queryInfo_ffi["windowType"] = A.load.Enum(queryInfo + 56, ["normal", "popup", "panel", "app", "devtools"]);

      const _ret = WEBEXT.tabs.query(queryInfo_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Query": (retPtr: Pointer, errPtr: Pointer, queryInfo: Pointer): heap.Ref<boolean> => {
      try {
        const queryInfo_ffi = {};

        if (A.load.Bool(queryInfo + 60)) {
          queryInfo_ffi["active"] = A.load.Bool(queryInfo + 0);
        }
        if (A.load.Bool(queryInfo + 61)) {
          queryInfo_ffi["audible"] = A.load.Bool(queryInfo + 1);
        }
        if (A.load.Bool(queryInfo + 62)) {
          queryInfo_ffi["autoDiscardable"] = A.load.Bool(queryInfo + 2);
        }
        if (A.load.Bool(queryInfo + 63)) {
          queryInfo_ffi["currentWindow"] = A.load.Bool(queryInfo + 3);
        }
        if (A.load.Bool(queryInfo + 64)) {
          queryInfo_ffi["discarded"] = A.load.Bool(queryInfo + 4);
        }
        if (A.load.Bool(queryInfo + 65)) {
          queryInfo_ffi["groupId"] = A.load.Int64(queryInfo + 8);
        }
        if (A.load.Bool(queryInfo + 66)) {
          queryInfo_ffi["highlighted"] = A.load.Bool(queryInfo + 16);
        }
        if (A.load.Bool(queryInfo + 67)) {
          queryInfo_ffi["index"] = A.load.Int64(queryInfo + 24);
        }
        if (A.load.Bool(queryInfo + 68)) {
          queryInfo_ffi["lastFocusedWindow"] = A.load.Bool(queryInfo + 32);
        }
        if (A.load.Bool(queryInfo + 69)) {
          queryInfo_ffi["muted"] = A.load.Bool(queryInfo + 33);
        }
        if (A.load.Bool(queryInfo + 70)) {
          queryInfo_ffi["pinned"] = A.load.Bool(queryInfo + 34);
        }
        queryInfo_ffi["status"] = A.load.Enum(queryInfo + 36, ["unloaded", "loading", "complete"]);
        queryInfo_ffi["title"] = A.load.Ref(queryInfo + 40, undefined);
        queryInfo_ffi["url"] = A.load.Ref(queryInfo + 44, undefined);
        if (A.load.Bool(queryInfo + 71)) {
          queryInfo_ffi["windowId"] = A.load.Int64(queryInfo + 48);
        }
        queryInfo_ffi["windowType"] = A.load.Enum(queryInfo + 56, ["normal", "popup", "panel", "app", "devtools"]);

        const _ret = WEBEXT.tabs.query(queryInfo_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Reload": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "reload" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Reload": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.reload);
    },
    "call_Reload": (retPtr: Pointer, tabId: number, reloadProperties: Pointer): void => {
      const reloadProperties_ffi = {};

      if (A.load.Bool(reloadProperties + 1)) {
        reloadProperties_ffi["bypassCache"] = A.load.Bool(reloadProperties + 0);
      }

      const _ret = WEBEXT.tabs.reload(tabId, reloadProperties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Reload": (retPtr: Pointer, errPtr: Pointer, tabId: number, reloadProperties: Pointer): heap.Ref<boolean> => {
      try {
        const reloadProperties_ffi = {};

        if (A.load.Bool(reloadProperties + 1)) {
          reloadProperties_ffi["bypassCache"] = A.load.Bool(reloadProperties + 0);
        }

        const _ret = WEBEXT.tabs.reload(tabId, reloadProperties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Remove": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "remove" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Remove": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.remove);
    },
    "call_Remove": (retPtr: Pointer, tabIds: heap.Ref<any>): void => {
      const _ret = WEBEXT.tabs.remove(A.H.get<any>(tabIds));
      A.store.Ref(retPtr, _ret);
    },
    "try_Remove": (retPtr: Pointer, errPtr: Pointer, tabIds: heap.Ref<any>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.remove(A.H.get<any>(tabIds));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveCSS": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "removeCSS" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveCSS": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.removeCSS);
    },
    "call_RemoveCSS": (retPtr: Pointer, tabId: number, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 25)) {
        details_ffi["allFrames"] = A.load.Bool(details + 0);
      }
      details_ffi["code"] = A.load.Ref(details + 4, undefined);
      details_ffi["cssOrigin"] = A.load.Enum(details + 8, ["author", "user"]);
      details_ffi["file"] = A.load.Ref(details + 12, undefined);
      if (A.load.Bool(details + 26)) {
        details_ffi["frameId"] = A.load.Int64(details + 16);
      }
      if (A.load.Bool(details + 27)) {
        details_ffi["matchAboutBlank"] = A.load.Bool(details + 24);
      }

      const _ret = WEBEXT.tabs.removeCSS(tabId, details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveCSS": (retPtr: Pointer, errPtr: Pointer, tabId: number, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 25)) {
          details_ffi["allFrames"] = A.load.Bool(details + 0);
        }
        details_ffi["code"] = A.load.Ref(details + 4, undefined);
        details_ffi["cssOrigin"] = A.load.Enum(details + 8, ["author", "user"]);
        details_ffi["file"] = A.load.Ref(details + 12, undefined);
        if (A.load.Bool(details + 26)) {
          details_ffi["frameId"] = A.load.Int64(details + 16);
        }
        if (A.load.Bool(details + 27)) {
          details_ffi["matchAboutBlank"] = A.load.Bool(details + 24);
        }

        const _ret = WEBEXT.tabs.removeCSS(tabId, details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "sendMessage" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.sendMessage);
    },
    "call_SendMessage": (retPtr: Pointer, tabId: number, message: heap.Ref<object>, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["documentId"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 16)) {
        options_ffi["frameId"] = A.load.Int64(options + 8);
      }

      const _ret = WEBEXT.tabs.sendMessage(tabId, A.H.get<object>(message), options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SendMessage": (
      retPtr: Pointer,
      errPtr: Pointer,
      tabId: number,
      message: heap.Ref<object>,
      options: Pointer
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["documentId"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 16)) {
          options_ffi["frameId"] = A.load.Int64(options + 8);
        }

        const _ret = WEBEXT.tabs.sendMessage(tabId, A.H.get<object>(message), options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "sendRequest" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.sendRequest);
    },
    "call_SendRequest": (retPtr: Pointer, tabId: number, request: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabs.sendRequest(tabId, A.H.get<object>(request));
      A.store.Ref(retPtr, _ret);
    },
    "try_SendRequest": (
      retPtr: Pointer,
      errPtr: Pointer,
      tabId: number,
      request: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.sendRequest(tabId, A.H.get<object>(request));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetZoom": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "setZoom" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetZoom": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.setZoom);
    },
    "call_SetZoom": (retPtr: Pointer, tabId: number, zoomFactor: number): void => {
      const _ret = WEBEXT.tabs.setZoom(tabId, zoomFactor);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetZoom": (retPtr: Pointer, errPtr: Pointer, tabId: number, zoomFactor: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.setZoom(tabId, zoomFactor);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetZoomSettings": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "setZoomSettings" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetZoomSettings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.setZoomSettings);
    },
    "call_SetZoomSettings": (retPtr: Pointer, tabId: number, zoomSettings: Pointer): void => {
      const zoomSettings_ffi = {};

      if (A.load.Bool(zoomSettings + 16)) {
        zoomSettings_ffi["defaultZoomFactor"] = A.load.Float64(zoomSettings + 0);
      }
      zoomSettings_ffi["mode"] = A.load.Enum(zoomSettings + 8, ["automatic", "manual", "disabled"]);
      zoomSettings_ffi["scope"] = A.load.Enum(zoomSettings + 12, ["per-origin", "per-tab"]);

      const _ret = WEBEXT.tabs.setZoomSettings(tabId, zoomSettings_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetZoomSettings": (
      retPtr: Pointer,
      errPtr: Pointer,
      tabId: number,
      zoomSettings: Pointer
    ): heap.Ref<boolean> => {
      try {
        const zoomSettings_ffi = {};

        if (A.load.Bool(zoomSettings + 16)) {
          zoomSettings_ffi["defaultZoomFactor"] = A.load.Float64(zoomSettings + 0);
        }
        zoomSettings_ffi["mode"] = A.load.Enum(zoomSettings + 8, ["automatic", "manual", "disabled"]);
        zoomSettings_ffi["scope"] = A.load.Enum(zoomSettings + 12, ["per-origin", "per-tab"]);

        const _ret = WEBEXT.tabs.setZoomSettings(tabId, zoomSettings_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Ungroup": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "ungroup" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Ungroup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.ungroup);
    },
    "call_Ungroup": (retPtr: Pointer, tabIds: heap.Ref<any>): void => {
      const _ret = WEBEXT.tabs.ungroup(A.H.get<any>(tabIds));
      A.store.Ref(retPtr, _ret);
    },
    "try_Ungroup": (retPtr: Pointer, errPtr: Pointer, tabIds: heap.Ref<any>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabs.ungroup(A.H.get<any>(tabIds));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Update": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "update" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Update": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.update);
    },
    "call_Update": (retPtr: Pointer, tabId: number, updateProperties: Pointer): void => {
      const updateProperties_ffi = {};

      if (A.load.Bool(updateProperties + 24)) {
        updateProperties_ffi["active"] = A.load.Bool(updateProperties + 0);
      }
      if (A.load.Bool(updateProperties + 25)) {
        updateProperties_ffi["autoDiscardable"] = A.load.Bool(updateProperties + 1);
      }
      if (A.load.Bool(updateProperties + 26)) {
        updateProperties_ffi["highlighted"] = A.load.Bool(updateProperties + 2);
      }
      if (A.load.Bool(updateProperties + 27)) {
        updateProperties_ffi["muted"] = A.load.Bool(updateProperties + 3);
      }
      if (A.load.Bool(updateProperties + 28)) {
        updateProperties_ffi["openerTabId"] = A.load.Int64(updateProperties + 8);
      }
      if (A.load.Bool(updateProperties + 29)) {
        updateProperties_ffi["pinned"] = A.load.Bool(updateProperties + 16);
      }
      if (A.load.Bool(updateProperties + 30)) {
        updateProperties_ffi["selected"] = A.load.Bool(updateProperties + 17);
      }
      updateProperties_ffi["url"] = A.load.Ref(updateProperties + 20, undefined);

      const _ret = WEBEXT.tabs.update(tabId, updateProperties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Update": (retPtr: Pointer, errPtr: Pointer, tabId: number, updateProperties: Pointer): heap.Ref<boolean> => {
      try {
        const updateProperties_ffi = {};

        if (A.load.Bool(updateProperties + 24)) {
          updateProperties_ffi["active"] = A.load.Bool(updateProperties + 0);
        }
        if (A.load.Bool(updateProperties + 25)) {
          updateProperties_ffi["autoDiscardable"] = A.load.Bool(updateProperties + 1);
        }
        if (A.load.Bool(updateProperties + 26)) {
          updateProperties_ffi["highlighted"] = A.load.Bool(updateProperties + 2);
        }
        if (A.load.Bool(updateProperties + 27)) {
          updateProperties_ffi["muted"] = A.load.Bool(updateProperties + 3);
        }
        if (A.load.Bool(updateProperties + 28)) {
          updateProperties_ffi["openerTabId"] = A.load.Int64(updateProperties + 8);
        }
        if (A.load.Bool(updateProperties + 29)) {
          updateProperties_ffi["pinned"] = A.load.Bool(updateProperties + 16);
        }
        if (A.load.Bool(updateProperties + 30)) {
          updateProperties_ffi["selected"] = A.load.Bool(updateProperties + 17);
        }
        updateProperties_ffi["url"] = A.load.Ref(updateProperties + 20, undefined);

        const _ret = WEBEXT.tabs.update(tabId, updateProperties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
