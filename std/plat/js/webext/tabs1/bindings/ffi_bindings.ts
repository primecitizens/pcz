import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/tabs1", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_Connect": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "connect" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Connect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.connect);
    },
    "call_Connect": (retPtr: Pointer, tabId: number, connectInfo: Pointer): void => {
      const connectInfo_ffi = {};

      connectInfo_ffi["documentId"] = A.load.Ref(connectInfo + 0, undefined);
      if (A.load.Bool(connectInfo + 20)) {
        connectInfo_ffi["frameId"] = A.load.Int64(connectInfo + 8);
      }
      connectInfo_ffi["name"] = A.load.Ref(connectInfo + 16, undefined);

      const _ret = WEBEXT.tabs.connect(tabId, connectInfo_ffi);
      if (typeof _ret === "undefined") {
        A.store.Bool(retPtr + 204, false);
        A.store.Ref(retPtr + 0, undefined);
        A.store.Ref(retPtr + 4, undefined);
        A.store.Ref(retPtr + 8, undefined);

        A.store.Bool(retPtr + 16 + 187, false);
        A.store.Ref(retPtr + 16 + 0, undefined);
        A.store.Ref(retPtr + 16 + 4, undefined);
        A.store.Bool(retPtr + 16 + 184, false);
        A.store.Int64(retPtr + 16 + 8, 0);
        A.store.Bool(retPtr + 16 + 185, false);
        A.store.Int64(retPtr + 16 + 16, 0);
        A.store.Bool(retPtr + 16 + 186, false);
        A.store.Int64(retPtr + 16 + 24, 0);
        A.store.Ref(retPtr + 16 + 32, undefined);
        A.store.Ref(retPtr + 16 + 36, undefined);
        A.store.Ref(retPtr + 16 + 40, undefined);

        A.store.Bool(retPtr + 16 + 48 + 125, false);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Bool(retPtr + 16 + 48 + 120, false);
        A.store.Bool(retPtr + 16 + 48 + 1, false);
        A.store.Bool(retPtr + 16 + 48 + 2, false);
        A.store.Bool(retPtr + 16 + 48 + 3, false);
        A.store.Ref(retPtr + 16 + 48 + 4, undefined);
        A.store.Int64(retPtr + 16 + 48 + 8, 0);
        A.store.Bool(retPtr + 16 + 48 + 121, false);
        A.store.Int64(retPtr + 16 + 48 + 16, 0);
        A.store.Bool(retPtr + 16 + 48 + 24, false);
        A.store.Bool(retPtr + 16 + 48 + 122, false);
        A.store.Int64(retPtr + 16 + 48 + 32, 0);
        A.store.Bool(retPtr + 16 + 48 + 40, false);
        A.store.Int64(retPtr + 16 + 48 + 48, 0);

        A.store.Bool(retPtr + 16 + 48 + 56 + 12, false);
        A.store.Ref(retPtr + 16 + 48 + 56 + 0, undefined);
        A.store.Bool(retPtr + 16 + 48 + 56 + 4, false);
        A.store.Enum(retPtr + 16 + 48 + 56 + 8, -1);
        A.store.Bool(retPtr + 16 + 48 + 123, false);
        A.store.Int64(retPtr + 16 + 48 + 72, 0);
        A.store.Ref(retPtr + 16 + 48 + 80, undefined);
        A.store.Bool(retPtr + 16 + 48 + 84, false);
        A.store.Bool(retPtr + 16 + 48 + 85, false);
        A.store.Ref(retPtr + 16 + 48 + 88, undefined);
        A.store.Enum(retPtr + 16 + 48 + 92, -1);
        A.store.Ref(retPtr + 16 + 48 + 96, undefined);
        A.store.Ref(retPtr + 16 + 48 + 100, undefined);
        A.store.Bool(retPtr + 16 + 48 + 124, false);
        A.store.Int64(retPtr + 16 + 48 + 104, 0);
        A.store.Int64(retPtr + 16 + 48 + 112, 0);
        A.store.Ref(retPtr + 16 + 176, undefined);
        A.store.Ref(retPtr + 16 + 180, undefined);
      } else {
        A.store.Bool(retPtr + 204, true);
        A.store.Ref(retPtr + 0, _ret["disconnect"]);
        A.store.Ref(retPtr + 4, _ret["name"]);
        A.store.Ref(retPtr + 8, _ret["postMessage"]);

        if (typeof _ret["sender"] === "undefined") {
          A.store.Bool(retPtr + 16 + 187, false);
          A.store.Ref(retPtr + 16 + 0, undefined);
          A.store.Ref(retPtr + 16 + 4, undefined);
          A.store.Bool(retPtr + 16 + 184, false);
          A.store.Int64(retPtr + 16 + 8, 0);
          A.store.Bool(retPtr + 16 + 185, false);
          A.store.Int64(retPtr + 16 + 16, 0);
          A.store.Bool(retPtr + 16 + 186, false);
          A.store.Int64(retPtr + 16 + 24, 0);
          A.store.Ref(retPtr + 16 + 32, undefined);
          A.store.Ref(retPtr + 16 + 36, undefined);
          A.store.Ref(retPtr + 16 + 40, undefined);

          A.store.Bool(retPtr + 16 + 48 + 125, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 120, false);
          A.store.Bool(retPtr + 16 + 48 + 1, false);
          A.store.Bool(retPtr + 16 + 48 + 2, false);
          A.store.Bool(retPtr + 16 + 48 + 3, false);
          A.store.Ref(retPtr + 16 + 48 + 4, undefined);
          A.store.Int64(retPtr + 16 + 48 + 8, 0);
          A.store.Bool(retPtr + 16 + 48 + 121, false);
          A.store.Int64(retPtr + 16 + 48 + 16, 0);
          A.store.Bool(retPtr + 16 + 48 + 24, false);
          A.store.Bool(retPtr + 16 + 48 + 122, false);
          A.store.Int64(retPtr + 16 + 48 + 32, 0);
          A.store.Bool(retPtr + 16 + 48 + 40, false);
          A.store.Int64(retPtr + 16 + 48 + 48, 0);

          A.store.Bool(retPtr + 16 + 48 + 56 + 12, false);
          A.store.Ref(retPtr + 16 + 48 + 56 + 0, undefined);
          A.store.Bool(retPtr + 16 + 48 + 56 + 4, false);
          A.store.Enum(retPtr + 16 + 48 + 56 + 8, -1);
          A.store.Bool(retPtr + 16 + 48 + 123, false);
          A.store.Int64(retPtr + 16 + 48 + 72, 0);
          A.store.Ref(retPtr + 16 + 48 + 80, undefined);
          A.store.Bool(retPtr + 16 + 48 + 84, false);
          A.store.Bool(retPtr + 16 + 48 + 85, false);
          A.store.Ref(retPtr + 16 + 48 + 88, undefined);
          A.store.Enum(retPtr + 16 + 48 + 92, -1);
          A.store.Ref(retPtr + 16 + 48 + 96, undefined);
          A.store.Ref(retPtr + 16 + 48 + 100, undefined);
          A.store.Bool(retPtr + 16 + 48 + 124, false);
          A.store.Int64(retPtr + 16 + 48 + 104, 0);
          A.store.Int64(retPtr + 16 + 48 + 112, 0);
          A.store.Ref(retPtr + 16 + 176, undefined);
          A.store.Ref(retPtr + 16 + 180, undefined);
        } else {
          A.store.Bool(retPtr + 16 + 187, true);
          A.store.Ref(retPtr + 16 + 0, _ret["sender"]["documentId"]);
          A.store.Ref(retPtr + 16 + 4, _ret["sender"]["documentLifecycle"]);
          A.store.Bool(retPtr + 16 + 184, "frameId" in _ret["sender"] ? true : false);
          A.store.Int64(
            retPtr + 16 + 8,
            _ret["sender"]["frameId"] === undefined ? 0 : (_ret["sender"]["frameId"] as number)
          );
          A.store.Bool(retPtr + 16 + 185, "guestProcessId" in _ret["sender"] ? true : false);
          A.store.Int64(
            retPtr + 16 + 16,
            _ret["sender"]["guestProcessId"] === undefined ? 0 : (_ret["sender"]["guestProcessId"] as number)
          );
          A.store.Bool(retPtr + 16 + 186, "guestRenderFrameRoutingId" in _ret["sender"] ? true : false);
          A.store.Int64(
            retPtr + 16 + 24,
            _ret["sender"]["guestRenderFrameRoutingId"] === undefined
              ? 0
              : (_ret["sender"]["guestRenderFrameRoutingId"] as number)
          );
          A.store.Ref(retPtr + 16 + 32, _ret["sender"]["id"]);
          A.store.Ref(retPtr + 16 + 36, _ret["sender"]["nativeApplication"]);
          A.store.Ref(retPtr + 16 + 40, _ret["sender"]["origin"]);

          if (typeof _ret["sender"]["tab"] === "undefined") {
            A.store.Bool(retPtr + 16 + 48 + 125, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 120, false);
            A.store.Bool(retPtr + 16 + 48 + 1, false);
            A.store.Bool(retPtr + 16 + 48 + 2, false);
            A.store.Bool(retPtr + 16 + 48 + 3, false);
            A.store.Ref(retPtr + 16 + 48 + 4, undefined);
            A.store.Int64(retPtr + 16 + 48 + 8, 0);
            A.store.Bool(retPtr + 16 + 48 + 121, false);
            A.store.Int64(retPtr + 16 + 48 + 16, 0);
            A.store.Bool(retPtr + 16 + 48 + 24, false);
            A.store.Bool(retPtr + 16 + 48 + 122, false);
            A.store.Int64(retPtr + 16 + 48 + 32, 0);
            A.store.Bool(retPtr + 16 + 48 + 40, false);
            A.store.Int64(retPtr + 16 + 48 + 48, 0);

            A.store.Bool(retPtr + 16 + 48 + 56 + 12, false);
            A.store.Ref(retPtr + 16 + 48 + 56 + 0, undefined);
            A.store.Bool(retPtr + 16 + 48 + 56 + 4, false);
            A.store.Enum(retPtr + 16 + 48 + 56 + 8, -1);
            A.store.Bool(retPtr + 16 + 48 + 123, false);
            A.store.Int64(retPtr + 16 + 48 + 72, 0);
            A.store.Ref(retPtr + 16 + 48 + 80, undefined);
            A.store.Bool(retPtr + 16 + 48 + 84, false);
            A.store.Bool(retPtr + 16 + 48 + 85, false);
            A.store.Ref(retPtr + 16 + 48 + 88, undefined);
            A.store.Enum(retPtr + 16 + 48 + 92, -1);
            A.store.Ref(retPtr + 16 + 48 + 96, undefined);
            A.store.Ref(retPtr + 16 + 48 + 100, undefined);
            A.store.Bool(retPtr + 16 + 48 + 124, false);
            A.store.Int64(retPtr + 16 + 48 + 104, 0);
            A.store.Int64(retPtr + 16 + 48 + 112, 0);
          } else {
            A.store.Bool(retPtr + 16 + 48 + 125, true);
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["active"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 120, "audible" in _ret["sender"]["tab"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 1, _ret["sender"]["tab"]["audible"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 2, _ret["sender"]["tab"]["autoDiscardable"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 3, _ret["sender"]["tab"]["discarded"] ? true : false);
            A.store.Ref(retPtr + 16 + 48 + 4, _ret["sender"]["tab"]["favIconUrl"]);
            A.store.Int64(
              retPtr + 16 + 48 + 8,
              _ret["sender"]["tab"]["groupId"] === undefined ? 0 : (_ret["sender"]["tab"]["groupId"] as number)
            );
            A.store.Bool(retPtr + 16 + 48 + 121, "height" in _ret["sender"]["tab"] ? true : false);
            A.store.Int64(
              retPtr + 16 + 48 + 16,
              _ret["sender"]["tab"]["height"] === undefined ? 0 : (_ret["sender"]["tab"]["height"] as number)
            );
            A.store.Bool(retPtr + 16 + 48 + 24, _ret["sender"]["tab"]["highlighted"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 122, "id" in _ret["sender"]["tab"] ? true : false);
            A.store.Int64(
              retPtr + 16 + 48 + 32,
              _ret["sender"]["tab"]["id"] === undefined ? 0 : (_ret["sender"]["tab"]["id"] as number)
            );
            A.store.Bool(retPtr + 16 + 48 + 40, _ret["sender"]["tab"]["incognito"] ? true : false);
            A.store.Int64(
              retPtr + 16 + 48 + 48,
              _ret["sender"]["tab"]["index"] === undefined ? 0 : (_ret["sender"]["tab"]["index"] as number)
            );

            if (typeof _ret["sender"]["tab"]["mutedInfo"] === "undefined") {
              A.store.Bool(retPtr + 16 + 48 + 56 + 12, false);
              A.store.Ref(retPtr + 16 + 48 + 56 + 0, undefined);
              A.store.Bool(retPtr + 16 + 48 + 56 + 4, false);
              A.store.Enum(retPtr + 16 + 48 + 56 + 8, -1);
            } else {
              A.store.Bool(retPtr + 16 + 48 + 56 + 12, true);
              A.store.Ref(retPtr + 16 + 48 + 56 + 0, _ret["sender"]["tab"]["mutedInfo"]["extensionId"]);
              A.store.Bool(retPtr + 16 + 48 + 56 + 4, _ret["sender"]["tab"]["mutedInfo"]["muted"] ? true : false);
              A.store.Enum(
                retPtr + 16 + 48 + 56 + 8,
                ["user", "capture", "extension"].indexOf(_ret["sender"]["tab"]["mutedInfo"]["reason"] as string)
              );
            }
            A.store.Bool(retPtr + 16 + 48 + 123, "openerTabId" in _ret["sender"]["tab"] ? true : false);
            A.store.Int64(
              retPtr + 16 + 48 + 72,
              _ret["sender"]["tab"]["openerTabId"] === undefined ? 0 : (_ret["sender"]["tab"]["openerTabId"] as number)
            );
            A.store.Ref(retPtr + 16 + 48 + 80, _ret["sender"]["tab"]["pendingUrl"]);
            A.store.Bool(retPtr + 16 + 48 + 84, _ret["sender"]["tab"]["pinned"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 85, _ret["sender"]["tab"]["selected"] ? true : false);
            A.store.Ref(retPtr + 16 + 48 + 88, _ret["sender"]["tab"]["sessionId"]);
            A.store.Enum(
              retPtr + 16 + 48 + 92,
              ["unloaded", "loading", "complete"].indexOf(_ret["sender"]["tab"]["status"] as string)
            );
            A.store.Ref(retPtr + 16 + 48 + 96, _ret["sender"]["tab"]["title"]);
            A.store.Ref(retPtr + 16 + 48 + 100, _ret["sender"]["tab"]["url"]);
            A.store.Bool(retPtr + 16 + 48 + 124, "width" in _ret["sender"]["tab"] ? true : false);
            A.store.Int64(
              retPtr + 16 + 48 + 104,
              _ret["sender"]["tab"]["width"] === undefined ? 0 : (_ret["sender"]["tab"]["width"] as number)
            );
            A.store.Int64(
              retPtr + 16 + 48 + 112,
              _ret["sender"]["tab"]["windowId"] === undefined ? 0 : (_ret["sender"]["tab"]["windowId"] as number)
            );
          }
          A.store.Ref(retPtr + 16 + 176, _ret["sender"]["tlsChannelId"]);
          A.store.Ref(retPtr + 16 + 180, _ret["sender"]["url"]);
        }
      }
    },
    "try_Connect": (retPtr: Pointer, errPtr: Pointer, tabId: number, connectInfo: Pointer): heap.Ref<boolean> => {
      try {
        const connectInfo_ffi = {};

        connectInfo_ffi["documentId"] = A.load.Ref(connectInfo + 0, undefined);
        if (A.load.Bool(connectInfo + 20)) {
          connectInfo_ffi["frameId"] = A.load.Int64(connectInfo + 8);
        }
        connectInfo_ffi["name"] = A.load.Ref(connectInfo + 16, undefined);

        const _ret = WEBEXT.tabs.connect(tabId, connectInfo_ffi);
        if (typeof _ret === "undefined") {
          A.store.Bool(retPtr + 204, false);
          A.store.Ref(retPtr + 0, undefined);
          A.store.Ref(retPtr + 4, undefined);
          A.store.Ref(retPtr + 8, undefined);

          A.store.Bool(retPtr + 16 + 187, false);
          A.store.Ref(retPtr + 16 + 0, undefined);
          A.store.Ref(retPtr + 16 + 4, undefined);
          A.store.Bool(retPtr + 16 + 184, false);
          A.store.Int64(retPtr + 16 + 8, 0);
          A.store.Bool(retPtr + 16 + 185, false);
          A.store.Int64(retPtr + 16 + 16, 0);
          A.store.Bool(retPtr + 16 + 186, false);
          A.store.Int64(retPtr + 16 + 24, 0);
          A.store.Ref(retPtr + 16 + 32, undefined);
          A.store.Ref(retPtr + 16 + 36, undefined);
          A.store.Ref(retPtr + 16 + 40, undefined);

          A.store.Bool(retPtr + 16 + 48 + 125, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 120, false);
          A.store.Bool(retPtr + 16 + 48 + 1, false);
          A.store.Bool(retPtr + 16 + 48 + 2, false);
          A.store.Bool(retPtr + 16 + 48 + 3, false);
          A.store.Ref(retPtr + 16 + 48 + 4, undefined);
          A.store.Int64(retPtr + 16 + 48 + 8, 0);
          A.store.Bool(retPtr + 16 + 48 + 121, false);
          A.store.Int64(retPtr + 16 + 48 + 16, 0);
          A.store.Bool(retPtr + 16 + 48 + 24, false);
          A.store.Bool(retPtr + 16 + 48 + 122, false);
          A.store.Int64(retPtr + 16 + 48 + 32, 0);
          A.store.Bool(retPtr + 16 + 48 + 40, false);
          A.store.Int64(retPtr + 16 + 48 + 48, 0);

          A.store.Bool(retPtr + 16 + 48 + 56 + 12, false);
          A.store.Ref(retPtr + 16 + 48 + 56 + 0, undefined);
          A.store.Bool(retPtr + 16 + 48 + 56 + 4, false);
          A.store.Enum(retPtr + 16 + 48 + 56 + 8, -1);
          A.store.Bool(retPtr + 16 + 48 + 123, false);
          A.store.Int64(retPtr + 16 + 48 + 72, 0);
          A.store.Ref(retPtr + 16 + 48 + 80, undefined);
          A.store.Bool(retPtr + 16 + 48 + 84, false);
          A.store.Bool(retPtr + 16 + 48 + 85, false);
          A.store.Ref(retPtr + 16 + 48 + 88, undefined);
          A.store.Enum(retPtr + 16 + 48 + 92, -1);
          A.store.Ref(retPtr + 16 + 48 + 96, undefined);
          A.store.Ref(retPtr + 16 + 48 + 100, undefined);
          A.store.Bool(retPtr + 16 + 48 + 124, false);
          A.store.Int64(retPtr + 16 + 48 + 104, 0);
          A.store.Int64(retPtr + 16 + 48 + 112, 0);
          A.store.Ref(retPtr + 16 + 176, undefined);
          A.store.Ref(retPtr + 16 + 180, undefined);
        } else {
          A.store.Bool(retPtr + 204, true);
          A.store.Ref(retPtr + 0, _ret["disconnect"]);
          A.store.Ref(retPtr + 4, _ret["name"]);
          A.store.Ref(retPtr + 8, _ret["postMessage"]);

          if (typeof _ret["sender"] === "undefined") {
            A.store.Bool(retPtr + 16 + 187, false);
            A.store.Ref(retPtr + 16 + 0, undefined);
            A.store.Ref(retPtr + 16 + 4, undefined);
            A.store.Bool(retPtr + 16 + 184, false);
            A.store.Int64(retPtr + 16 + 8, 0);
            A.store.Bool(retPtr + 16 + 185, false);
            A.store.Int64(retPtr + 16 + 16, 0);
            A.store.Bool(retPtr + 16 + 186, false);
            A.store.Int64(retPtr + 16 + 24, 0);
            A.store.Ref(retPtr + 16 + 32, undefined);
            A.store.Ref(retPtr + 16 + 36, undefined);
            A.store.Ref(retPtr + 16 + 40, undefined);

            A.store.Bool(retPtr + 16 + 48 + 125, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 120, false);
            A.store.Bool(retPtr + 16 + 48 + 1, false);
            A.store.Bool(retPtr + 16 + 48 + 2, false);
            A.store.Bool(retPtr + 16 + 48 + 3, false);
            A.store.Ref(retPtr + 16 + 48 + 4, undefined);
            A.store.Int64(retPtr + 16 + 48 + 8, 0);
            A.store.Bool(retPtr + 16 + 48 + 121, false);
            A.store.Int64(retPtr + 16 + 48 + 16, 0);
            A.store.Bool(retPtr + 16 + 48 + 24, false);
            A.store.Bool(retPtr + 16 + 48 + 122, false);
            A.store.Int64(retPtr + 16 + 48 + 32, 0);
            A.store.Bool(retPtr + 16 + 48 + 40, false);
            A.store.Int64(retPtr + 16 + 48 + 48, 0);

            A.store.Bool(retPtr + 16 + 48 + 56 + 12, false);
            A.store.Ref(retPtr + 16 + 48 + 56 + 0, undefined);
            A.store.Bool(retPtr + 16 + 48 + 56 + 4, false);
            A.store.Enum(retPtr + 16 + 48 + 56 + 8, -1);
            A.store.Bool(retPtr + 16 + 48 + 123, false);
            A.store.Int64(retPtr + 16 + 48 + 72, 0);
            A.store.Ref(retPtr + 16 + 48 + 80, undefined);
            A.store.Bool(retPtr + 16 + 48 + 84, false);
            A.store.Bool(retPtr + 16 + 48 + 85, false);
            A.store.Ref(retPtr + 16 + 48 + 88, undefined);
            A.store.Enum(retPtr + 16 + 48 + 92, -1);
            A.store.Ref(retPtr + 16 + 48 + 96, undefined);
            A.store.Ref(retPtr + 16 + 48 + 100, undefined);
            A.store.Bool(retPtr + 16 + 48 + 124, false);
            A.store.Int64(retPtr + 16 + 48 + 104, 0);
            A.store.Int64(retPtr + 16 + 48 + 112, 0);
            A.store.Ref(retPtr + 16 + 176, undefined);
            A.store.Ref(retPtr + 16 + 180, undefined);
          } else {
            A.store.Bool(retPtr + 16 + 187, true);
            A.store.Ref(retPtr + 16 + 0, _ret["sender"]["documentId"]);
            A.store.Ref(retPtr + 16 + 4, _ret["sender"]["documentLifecycle"]);
            A.store.Bool(retPtr + 16 + 184, "frameId" in _ret["sender"] ? true : false);
            A.store.Int64(
              retPtr + 16 + 8,
              _ret["sender"]["frameId"] === undefined ? 0 : (_ret["sender"]["frameId"] as number)
            );
            A.store.Bool(retPtr + 16 + 185, "guestProcessId" in _ret["sender"] ? true : false);
            A.store.Int64(
              retPtr + 16 + 16,
              _ret["sender"]["guestProcessId"] === undefined ? 0 : (_ret["sender"]["guestProcessId"] as number)
            );
            A.store.Bool(retPtr + 16 + 186, "guestRenderFrameRoutingId" in _ret["sender"] ? true : false);
            A.store.Int64(
              retPtr + 16 + 24,
              _ret["sender"]["guestRenderFrameRoutingId"] === undefined
                ? 0
                : (_ret["sender"]["guestRenderFrameRoutingId"] as number)
            );
            A.store.Ref(retPtr + 16 + 32, _ret["sender"]["id"]);
            A.store.Ref(retPtr + 16 + 36, _ret["sender"]["nativeApplication"]);
            A.store.Ref(retPtr + 16 + 40, _ret["sender"]["origin"]);

            if (typeof _ret["sender"]["tab"] === "undefined") {
              A.store.Bool(retPtr + 16 + 48 + 125, false);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Bool(retPtr + 16 + 48 + 120, false);
              A.store.Bool(retPtr + 16 + 48 + 1, false);
              A.store.Bool(retPtr + 16 + 48 + 2, false);
              A.store.Bool(retPtr + 16 + 48 + 3, false);
              A.store.Ref(retPtr + 16 + 48 + 4, undefined);
              A.store.Int64(retPtr + 16 + 48 + 8, 0);
              A.store.Bool(retPtr + 16 + 48 + 121, false);
              A.store.Int64(retPtr + 16 + 48 + 16, 0);
              A.store.Bool(retPtr + 16 + 48 + 24, false);
              A.store.Bool(retPtr + 16 + 48 + 122, false);
              A.store.Int64(retPtr + 16 + 48 + 32, 0);
              A.store.Bool(retPtr + 16 + 48 + 40, false);
              A.store.Int64(retPtr + 16 + 48 + 48, 0);

              A.store.Bool(retPtr + 16 + 48 + 56 + 12, false);
              A.store.Ref(retPtr + 16 + 48 + 56 + 0, undefined);
              A.store.Bool(retPtr + 16 + 48 + 56 + 4, false);
              A.store.Enum(retPtr + 16 + 48 + 56 + 8, -1);
              A.store.Bool(retPtr + 16 + 48 + 123, false);
              A.store.Int64(retPtr + 16 + 48 + 72, 0);
              A.store.Ref(retPtr + 16 + 48 + 80, undefined);
              A.store.Bool(retPtr + 16 + 48 + 84, false);
              A.store.Bool(retPtr + 16 + 48 + 85, false);
              A.store.Ref(retPtr + 16 + 48 + 88, undefined);
              A.store.Enum(retPtr + 16 + 48 + 92, -1);
              A.store.Ref(retPtr + 16 + 48 + 96, undefined);
              A.store.Ref(retPtr + 16 + 48 + 100, undefined);
              A.store.Bool(retPtr + 16 + 48 + 124, false);
              A.store.Int64(retPtr + 16 + 48 + 104, 0);
              A.store.Int64(retPtr + 16 + 48 + 112, 0);
            } else {
              A.store.Bool(retPtr + 16 + 48 + 125, true);
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["active"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 120, "audible" in _ret["sender"]["tab"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 1, _ret["sender"]["tab"]["audible"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 2, _ret["sender"]["tab"]["autoDiscardable"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 3, _ret["sender"]["tab"]["discarded"] ? true : false);
              A.store.Ref(retPtr + 16 + 48 + 4, _ret["sender"]["tab"]["favIconUrl"]);
              A.store.Int64(
                retPtr + 16 + 48 + 8,
                _ret["sender"]["tab"]["groupId"] === undefined ? 0 : (_ret["sender"]["tab"]["groupId"] as number)
              );
              A.store.Bool(retPtr + 16 + 48 + 121, "height" in _ret["sender"]["tab"] ? true : false);
              A.store.Int64(
                retPtr + 16 + 48 + 16,
                _ret["sender"]["tab"]["height"] === undefined ? 0 : (_ret["sender"]["tab"]["height"] as number)
              );
              A.store.Bool(retPtr + 16 + 48 + 24, _ret["sender"]["tab"]["highlighted"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 122, "id" in _ret["sender"]["tab"] ? true : false);
              A.store.Int64(
                retPtr + 16 + 48 + 32,
                _ret["sender"]["tab"]["id"] === undefined ? 0 : (_ret["sender"]["tab"]["id"] as number)
              );
              A.store.Bool(retPtr + 16 + 48 + 40, _ret["sender"]["tab"]["incognito"] ? true : false);
              A.store.Int64(
                retPtr + 16 + 48 + 48,
                _ret["sender"]["tab"]["index"] === undefined ? 0 : (_ret["sender"]["tab"]["index"] as number)
              );

              if (typeof _ret["sender"]["tab"]["mutedInfo"] === "undefined") {
                A.store.Bool(retPtr + 16 + 48 + 56 + 12, false);
                A.store.Ref(retPtr + 16 + 48 + 56 + 0, undefined);
                A.store.Bool(retPtr + 16 + 48 + 56 + 4, false);
                A.store.Enum(retPtr + 16 + 48 + 56 + 8, -1);
              } else {
                A.store.Bool(retPtr + 16 + 48 + 56 + 12, true);
                A.store.Ref(retPtr + 16 + 48 + 56 + 0, _ret["sender"]["tab"]["mutedInfo"]["extensionId"]);
                A.store.Bool(retPtr + 16 + 48 + 56 + 4, _ret["sender"]["tab"]["mutedInfo"]["muted"] ? true : false);
                A.store.Enum(
                  retPtr + 16 + 48 + 56 + 8,
                  ["user", "capture", "extension"].indexOf(_ret["sender"]["tab"]["mutedInfo"]["reason"] as string)
                );
              }
              A.store.Bool(retPtr + 16 + 48 + 123, "openerTabId" in _ret["sender"]["tab"] ? true : false);
              A.store.Int64(
                retPtr + 16 + 48 + 72,
                _ret["sender"]["tab"]["openerTabId"] === undefined
                  ? 0
                  : (_ret["sender"]["tab"]["openerTabId"] as number)
              );
              A.store.Ref(retPtr + 16 + 48 + 80, _ret["sender"]["tab"]["pendingUrl"]);
              A.store.Bool(retPtr + 16 + 48 + 84, _ret["sender"]["tab"]["pinned"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 85, _ret["sender"]["tab"]["selected"] ? true : false);
              A.store.Ref(retPtr + 16 + 48 + 88, _ret["sender"]["tab"]["sessionId"]);
              A.store.Enum(
                retPtr + 16 + 48 + 92,
                ["unloaded", "loading", "complete"].indexOf(_ret["sender"]["tab"]["status"] as string)
              );
              A.store.Ref(retPtr + 16 + 48 + 96, _ret["sender"]["tab"]["title"]);
              A.store.Ref(retPtr + 16 + 48 + 100, _ret["sender"]["tab"]["url"]);
              A.store.Bool(retPtr + 16 + 48 + 124, "width" in _ret["sender"]["tab"] ? true : false);
              A.store.Int64(
                retPtr + 16 + 48 + 104,
                _ret["sender"]["tab"]["width"] === undefined ? 0 : (_ret["sender"]["tab"]["width"] as number)
              );
              A.store.Int64(
                retPtr + 16 + 48 + 112,
                _ret["sender"]["tab"]["windowId"] === undefined ? 0 : (_ret["sender"]["tab"]["windowId"] as number)
              );
            }
            A.store.Ref(retPtr + 16 + 176, _ret["sender"]["tlsChannelId"]);
            A.store.Ref(retPtr + 16 + 180, _ret["sender"]["url"]);
          }
        }
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Highlight": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabs && "highlight" in WEBEXT?.tabs) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Highlight": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabs.highlight);
    },
    "call_Highlight": (retPtr: Pointer, highlightInfo: Pointer): void => {
      const highlightInfo_ffi = {};

      highlightInfo_ffi["tabs"] = A.load.Ref(highlightInfo + 0, undefined);
      if (A.load.Bool(highlightInfo + 16)) {
        highlightInfo_ffi["windowId"] = A.load.Int64(highlightInfo + 8);
      }

      const _ret = WEBEXT.tabs.highlight(highlightInfo_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Highlight": (retPtr: Pointer, errPtr: Pointer, highlightInfo: Pointer): heap.Ref<boolean> => {
      try {
        const highlightInfo_ffi = {};

        highlightInfo_ffi["tabs"] = A.load.Ref(highlightInfo + 0, undefined);
        if (A.load.Bool(highlightInfo + 16)) {
          highlightInfo_ffi["windowId"] = A.load.Int64(highlightInfo + 8);
        }

        const _ret = WEBEXT.tabs.highlight(highlightInfo_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
