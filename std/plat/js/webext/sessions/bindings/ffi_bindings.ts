import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/sessions", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Session": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 222, false);
        A.store.Int64(ptr + 0, 0);

        A.store.Bool(ptr + 8 + 0, false);
        A.store.Bool(ptr + 8 + 0, false);
        A.store.Bool(ptr + 8 + 0, false);
        A.store.Bool(ptr + 8 + 0, false);
        A.store.Bool(ptr + 8 + 0, false);
        A.store.Bool(ptr + 8 + 0, false);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Int0(ptr + 8 + 0, 0);
        A.store.Bool(ptr + 8 + 0, false);
        A.store.Int0(ptr + 8 + 0, 0);
        A.store.Bool(ptr + 8 + 0, false);
        A.store.Bool(ptr + 8 + 0, false);
        A.store.Int0(ptr + 8 + 0, 0);
        A.store.Bool(ptr + 8 + 0, false);
        A.store.Int0(ptr + 8 + 0, 0);

        A.store.Bool(ptr + 8 + 0 + 0, false);
        A.store.Ref(ptr + 8 + 0 + 0, undefined);
        A.store.Bool(ptr + 8 + 0 + 0, false);
        A.store.Enum(ptr + 8 + 0 + 0, -1);
        A.store.Bool(ptr + 8 + 0, false);
        A.store.Int0(ptr + 8 + 0, 0);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Bool(ptr + 8 + 0, false);
        A.store.Bool(ptr + 8 + 0, false);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Enum(ptr + 8 + 0, -1);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Bool(ptr + 8 + 0, false);
        A.store.Int0(ptr + 8 + 0, 0);
        A.store.Int0(ptr + 8 + 0, 0);

        A.store.Bool(ptr + 136 + 0, false);
        A.store.Bool(ptr + 136 + 0, false);
        A.store.Bool(ptr + 136 + 0, false);
        A.store.Bool(ptr + 136 + 0, false);
        A.store.Int0(ptr + 136 + 0, 0);
        A.store.Bool(ptr + 136 + 0, false);
        A.store.Int0(ptr + 136 + 0, 0);
        A.store.Bool(ptr + 136 + 0, false);
        A.store.Bool(ptr + 136 + 0, false);
        A.store.Int0(ptr + 136 + 0, 0);
        A.store.Ref(ptr + 136 + 0, undefined);
        A.store.Enum(ptr + 136 + 0, -1);
        A.store.Ref(ptr + 136 + 0, undefined);
        A.store.Bool(ptr + 136 + 0, false);
        A.store.Int0(ptr + 136 + 0, 0);
        A.store.Enum(ptr + 136 + 0, -1);
        A.store.Bool(ptr + 136 + 0, false);
        A.store.Int0(ptr + 136 + 0, 0);
      } else {
        A.store.Bool(ptr + 222, true);
        A.store.Int64(ptr + 0, x["lastModified"] === undefined ? 0 : (x["lastModified"] as number));

        if (typeof x["tab"] === "undefined") {
          A.store.Bool(ptr + 8 + 0, false);
          A.store.Bool(ptr + 8 + 0, false);
          A.store.Bool(ptr + 8 + 0, false);
          A.store.Bool(ptr + 8 + 0, false);
          A.store.Bool(ptr + 8 + 0, false);
          A.store.Bool(ptr + 8 + 0, false);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Int0(ptr + 8 + 0, 0);
          A.store.Bool(ptr + 8 + 0, false);
          A.store.Int0(ptr + 8 + 0, 0);
          A.store.Bool(ptr + 8 + 0, false);
          A.store.Bool(ptr + 8 + 0, false);
          A.store.Int0(ptr + 8 + 0, 0);
          A.store.Bool(ptr + 8 + 0, false);
          A.store.Int0(ptr + 8 + 0, 0);

          A.store.Bool(ptr + 8 + 0 + 0, false);
          A.store.Ref(ptr + 8 + 0 + 0, undefined);
          A.store.Bool(ptr + 8 + 0 + 0, false);
          A.store.Enum(ptr + 8 + 0 + 0, -1);
          A.store.Bool(ptr + 8 + 0, false);
          A.store.Int0(ptr + 8 + 0, 0);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Bool(ptr + 8 + 0, false);
          A.store.Bool(ptr + 8 + 0, false);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Enum(ptr + 8 + 0, -1);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Bool(ptr + 8 + 0, false);
          A.store.Int0(ptr + 8 + 0, 0);
          A.store.Int0(ptr + 8 + 0, 0);
        } else {
          A.store.Bool(ptr + 8 + 0, true);
          A.store.Bool(ptr + 8 + 0, x["tab"]["active"] ? true : false);
          A.store.Bool(ptr + 8 + 0, "audible" in x["tab"] ? true : false);
          A.store.Bool(ptr + 8 + 0, x["tab"]["audible"] ? true : false);
          A.store.Bool(ptr + 8 + 0, x["tab"]["autoDiscardable"] ? true : false);
          A.store.Bool(ptr + 8 + 0, x["tab"]["discarded"] ? true : false);
          A.store.Ref(ptr + 8 + 0, x["tab"]["favIconUrl"]);
          A.store.Int0(ptr + 8 + 0, x["tab"]["groupId"] === undefined ? 0 : (x["tab"]["groupId"] as number));
          A.store.Bool(ptr + 8 + 0, "height" in x["tab"] ? true : false);
          A.store.Int0(ptr + 8 + 0, x["tab"]["height"] === undefined ? 0 : (x["tab"]["height"] as number));
          A.store.Bool(ptr + 8 + 0, x["tab"]["highlighted"] ? true : false);
          A.store.Bool(ptr + 8 + 0, "id" in x["tab"] ? true : false);
          A.store.Int0(ptr + 8 + 0, x["tab"]["id"] === undefined ? 0 : (x["tab"]["id"] as number));
          A.store.Bool(ptr + 8 + 0, x["tab"]["incognito"] ? true : false);
          A.store.Int0(ptr + 8 + 0, x["tab"]["index"] === undefined ? 0 : (x["tab"]["index"] as number));

          if (typeof x["tab"]["mutedInfo"] === "undefined") {
            A.store.Bool(ptr + 8 + 0 + 0, false);
            A.store.Ref(ptr + 8 + 0 + 0, undefined);
            A.store.Bool(ptr + 8 + 0 + 0, false);
            A.store.Enum(ptr + 8 + 0 + 0, -1);
          } else {
            A.store.Bool(ptr + 8 + 0 + 0, true);
            A.store.Ref(ptr + 8 + 0 + 0, x["tab"]["mutedInfo"]["extensionId"]);
            A.store.Bool(ptr + 8 + 0 + 0, x["tab"]["mutedInfo"]["muted"] ? true : false);
            A.store.Enum(
              ptr + 8 + 0 + 0,
              ["user", "capture", "extension"].indexOf(x["tab"]["mutedInfo"]["reason"] as string)
            );
          }
          A.store.Bool(ptr + 8 + 0, "openerTabId" in x["tab"] ? true : false);
          A.store.Int0(ptr + 8 + 0, x["tab"]["openerTabId"] === undefined ? 0 : (x["tab"]["openerTabId"] as number));
          A.store.Ref(ptr + 8 + 0, x["tab"]["pendingUrl"]);
          A.store.Bool(ptr + 8 + 0, x["tab"]["pinned"] ? true : false);
          A.store.Bool(ptr + 8 + 0, x["tab"]["selected"] ? true : false);
          A.store.Ref(ptr + 8 + 0, x["tab"]["sessionId"]);
          A.store.Enum(ptr + 8 + 0, ["unloaded", "loading", "complete"].indexOf(x["tab"]["status"] as string));
          A.store.Ref(ptr + 8 + 0, x["tab"]["title"]);
          A.store.Ref(ptr + 8 + 0, x["tab"]["url"]);
          A.store.Bool(ptr + 8 + 0, "width" in x["tab"] ? true : false);
          A.store.Int0(ptr + 8 + 0, x["tab"]["width"] === undefined ? 0 : (x["tab"]["width"] as number));
          A.store.Int0(ptr + 8 + 0, x["tab"]["windowId"] === undefined ? 0 : (x["tab"]["windowId"] as number));
        }

        if (typeof x["window"] === "undefined") {
          A.store.Bool(ptr + 136 + 0, false);
          A.store.Bool(ptr + 136 + 0, false);
          A.store.Bool(ptr + 136 + 0, false);
          A.store.Bool(ptr + 136 + 0, false);
          A.store.Int0(ptr + 136 + 0, 0);
          A.store.Bool(ptr + 136 + 0, false);
          A.store.Int0(ptr + 136 + 0, 0);
          A.store.Bool(ptr + 136 + 0, false);
          A.store.Bool(ptr + 136 + 0, false);
          A.store.Int0(ptr + 136 + 0, 0);
          A.store.Ref(ptr + 136 + 0, undefined);
          A.store.Enum(ptr + 136 + 0, -1);
          A.store.Ref(ptr + 136 + 0, undefined);
          A.store.Bool(ptr + 136 + 0, false);
          A.store.Int0(ptr + 136 + 0, 0);
          A.store.Enum(ptr + 136 + 0, -1);
          A.store.Bool(ptr + 136 + 0, false);
          A.store.Int0(ptr + 136 + 0, 0);
        } else {
          A.store.Bool(ptr + 136 + 0, true);
          A.store.Bool(ptr + 136 + 0, x["window"]["alwaysOnTop"] ? true : false);
          A.store.Bool(ptr + 136 + 0, x["window"]["focused"] ? true : false);
          A.store.Bool(ptr + 136 + 0, "height" in x["window"] ? true : false);
          A.store.Int0(ptr + 136 + 0, x["window"]["height"] === undefined ? 0 : (x["window"]["height"] as number));
          A.store.Bool(ptr + 136 + 0, "id" in x["window"] ? true : false);
          A.store.Int0(ptr + 136 + 0, x["window"]["id"] === undefined ? 0 : (x["window"]["id"] as number));
          A.store.Bool(ptr + 136 + 0, x["window"]["incognito"] ? true : false);
          A.store.Bool(ptr + 136 + 0, "left" in x["window"] ? true : false);
          A.store.Int0(ptr + 136 + 0, x["window"]["left"] === undefined ? 0 : (x["window"]["left"] as number));
          A.store.Ref(ptr + 136 + 0, x["window"]["sessionId"]);
          A.store.Enum(
            ptr + 136 + 0,
            ["normal", "minimized", "maximized", "fullscreen", "locked-fullscreen"].indexOf(
              x["window"]["state"] as string
            )
          );
          A.store.Ref(ptr + 136 + 0, x["window"]["tabs"]);
          A.store.Bool(ptr + 136 + 0, "top" in x["window"] ? true : false);
          A.store.Int0(ptr + 136 + 0, x["window"]["top"] === undefined ? 0 : (x["window"]["top"] as number));
          A.store.Enum(
            ptr + 136 + 0,
            ["normal", "popup", "panel", "app", "devtools"].indexOf(x["window"]["type"] as string)
          );
          A.store.Bool(ptr + 136 + 0, "width" in x["window"] ? true : false);
          A.store.Int0(ptr + 136 + 0, x["window"]["width"] === undefined ? 0 : (x["window"]["width"] as number));
        }
      }
    },
    "load_Session": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["lastModified"] = A.load.Int64(ptr + 0);
      if (A.load.Bool(ptr + 8 + 0)) {
        x["tab"] = {};
        x["tab"]["active"] = A.load.Bool(ptr + 8 + 0);
        if (A.load.Bool(ptr + 8 + 0)) {
          x["tab"]["audible"] = A.load.Bool(ptr + 8 + 0);
        } else {
          delete x["tab"]["audible"];
        }
        x["tab"]["autoDiscardable"] = A.load.Bool(ptr + 8 + 0);
        x["tab"]["discarded"] = A.load.Bool(ptr + 8 + 0);
        x["tab"]["favIconUrl"] = A.load.Ref(ptr + 8 + 0, undefined);
        x["tab"]["groupId"] = A.load.Int0(ptr + 8 + 0);
        if (A.load.Bool(ptr + 8 + 0)) {
          x["tab"]["height"] = A.load.Int0(ptr + 8 + 0);
        } else {
          delete x["tab"]["height"];
        }
        x["tab"]["highlighted"] = A.load.Bool(ptr + 8 + 0);
        if (A.load.Bool(ptr + 8 + 0)) {
          x["tab"]["id"] = A.load.Int0(ptr + 8 + 0);
        } else {
          delete x["tab"]["id"];
        }
        x["tab"]["incognito"] = A.load.Bool(ptr + 8 + 0);
        x["tab"]["index"] = A.load.Int0(ptr + 8 + 0);
        if (A.load.Bool(ptr + 8 + 0 + 0)) {
          x["tab"]["mutedInfo"] = {};
          x["tab"]["mutedInfo"]["extensionId"] = A.load.Ref(ptr + 8 + 0 + 0, undefined);
          x["tab"]["mutedInfo"]["muted"] = A.load.Bool(ptr + 8 + 0 + 0);
          x["tab"]["mutedInfo"]["reason"] = A.load.Enum(ptr + 8 + 0 + 0, ["user", "capture", "extension"]);
        } else {
          delete x["tab"]["mutedInfo"];
        }
        if (A.load.Bool(ptr + 8 + 0)) {
          x["tab"]["openerTabId"] = A.load.Int0(ptr + 8 + 0);
        } else {
          delete x["tab"]["openerTabId"];
        }
        x["tab"]["pendingUrl"] = A.load.Ref(ptr + 8 + 0, undefined);
        x["tab"]["pinned"] = A.load.Bool(ptr + 8 + 0);
        x["tab"]["selected"] = A.load.Bool(ptr + 8 + 0);
        x["tab"]["sessionId"] = A.load.Ref(ptr + 8 + 0, undefined);
        x["tab"]["status"] = A.load.Enum(ptr + 8 + 0, ["unloaded", "loading", "complete"]);
        x["tab"]["title"] = A.load.Ref(ptr + 8 + 0, undefined);
        x["tab"]["url"] = A.load.Ref(ptr + 8 + 0, undefined);
        if (A.load.Bool(ptr + 8 + 0)) {
          x["tab"]["width"] = A.load.Int0(ptr + 8 + 0);
        } else {
          delete x["tab"]["width"];
        }
        x["tab"]["windowId"] = A.load.Int0(ptr + 8 + 0);
      } else {
        delete x["tab"];
      }
      if (A.load.Bool(ptr + 136 + 0)) {
        x["window"] = {};
        x["window"]["alwaysOnTop"] = A.load.Bool(ptr + 136 + 0);
        x["window"]["focused"] = A.load.Bool(ptr + 136 + 0);
        if (A.load.Bool(ptr + 136 + 0)) {
          x["window"]["height"] = A.load.Int0(ptr + 136 + 0);
        } else {
          delete x["window"]["height"];
        }
        if (A.load.Bool(ptr + 136 + 0)) {
          x["window"]["id"] = A.load.Int0(ptr + 136 + 0);
        } else {
          delete x["window"]["id"];
        }
        x["window"]["incognito"] = A.load.Bool(ptr + 136 + 0);
        if (A.load.Bool(ptr + 136 + 0)) {
          x["window"]["left"] = A.load.Int0(ptr + 136 + 0);
        } else {
          delete x["window"]["left"];
        }
        x["window"]["sessionId"] = A.load.Ref(ptr + 136 + 0, undefined);
        x["window"]["state"] = A.load.Enum(ptr + 136 + 0, [
          "normal",
          "minimized",
          "maximized",
          "fullscreen",
          "locked-fullscreen",
        ]);
        x["window"]["tabs"] = A.load.Ref(ptr + 136 + 0, undefined);
        if (A.load.Bool(ptr + 136 + 0)) {
          x["window"]["top"] = A.load.Int0(ptr + 136 + 0);
        } else {
          delete x["window"]["top"];
        }
        x["window"]["type"] = A.load.Enum(ptr + 136 + 0, ["normal", "popup", "panel", "app", "devtools"]);
        if (A.load.Bool(ptr + 136 + 0)) {
          x["window"]["width"] = A.load.Int0(ptr + 136 + 0);
        } else {
          delete x["window"]["width"];
        }
      } else {
        delete x["window"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Device": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["deviceName"]);
        A.store.Ref(ptr + 4, x["info"]);
        A.store.Ref(ptr + 8, x["sessions"]);
      }
    },
    "load_Device": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["deviceName"] = A.load.Ref(ptr + 0, undefined);
      x["info"] = A.load.Ref(ptr + 4, undefined);
      x["sessions"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Filter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "maxResults" in x ? true : false);
        A.store.Int64(ptr + 0, x["maxResults"] === undefined ? 0 : (x["maxResults"] as number));
      }
    },
    "load_Filter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["maxResults"] = A.load.Int64(ptr + 0);
      } else {
        delete x["maxResults"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "get_MAX_SESSION_RESULTS": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.sessions && "MAX_SESSION_RESULTS" in WEBEXT?.sessions) {
        const val = WEBEXT.sessions.MAX_SESSION_RESULTS;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_MAX_SESSION_RESULTS": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.sessions, "MAX_SESSION_RESULTS", A.H.get<object>(val), WEBEXT.sessions)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "has_GetDevices": (): heap.Ref<boolean> => {
      if (WEBEXT?.sessions && "getDevices" in WEBEXT?.sessions) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDevices": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sessions.getDevices);
    },
    "call_GetDevices": (retPtr: Pointer, filter: Pointer): void => {
      const filter_ffi = {};

      if (A.load.Bool(filter + 8)) {
        filter_ffi["maxResults"] = A.load.Int64(filter + 0);
      }

      const _ret = WEBEXT.sessions.getDevices(filter_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDevices": (retPtr: Pointer, errPtr: Pointer, filter: Pointer): heap.Ref<boolean> => {
      try {
        const filter_ffi = {};

        if (A.load.Bool(filter + 8)) {
          filter_ffi["maxResults"] = A.load.Int64(filter + 0);
        }

        const _ret = WEBEXT.sessions.getDevices(filter_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetRecentlyClosed": (): heap.Ref<boolean> => {
      if (WEBEXT?.sessions && "getRecentlyClosed" in WEBEXT?.sessions) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetRecentlyClosed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sessions.getRecentlyClosed);
    },
    "call_GetRecentlyClosed": (retPtr: Pointer, filter: Pointer): void => {
      const filter_ffi = {};

      if (A.load.Bool(filter + 8)) {
        filter_ffi["maxResults"] = A.load.Int64(filter + 0);
      }

      const _ret = WEBEXT.sessions.getRecentlyClosed(filter_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetRecentlyClosed": (retPtr: Pointer, errPtr: Pointer, filter: Pointer): heap.Ref<boolean> => {
      try {
        const filter_ffi = {};

        if (A.load.Bool(filter + 8)) {
          filter_ffi["maxResults"] = A.load.Int64(filter + 0);
        }

        const _ret = WEBEXT.sessions.getRecentlyClosed(filter_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.sessions?.onChanged && "addListener" in WEBEXT?.sessions?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sessions.onChanged.addListener);
    },
    "call_OnChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sessions.onChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sessions.onChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.sessions?.onChanged && "removeListener" in WEBEXT?.sessions?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sessions.onChanged.removeListener);
    },
    "call_OffChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sessions.onChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sessions.onChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.sessions?.onChanged && "hasListener" in WEBEXT?.sessions?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sessions.onChanged.hasListener);
    },
    "call_HasOnChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sessions.onChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sessions.onChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Restore": (): heap.Ref<boolean> => {
      if (WEBEXT?.sessions && "restore" in WEBEXT?.sessions) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Restore": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sessions.restore);
    },
    "call_Restore": (retPtr: Pointer, sessionId: heap.Ref<object>): void => {
      const _ret = WEBEXT.sessions.restore(A.H.get<object>(sessionId));
      A.store.Ref(retPtr, _ret);
    },
    "try_Restore": (retPtr: Pointer, errPtr: Pointer, sessionId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sessions.restore(A.H.get<object>(sessionId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
