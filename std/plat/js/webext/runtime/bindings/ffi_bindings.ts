import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/runtime", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_ConnectArgConnectInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "includeTlsChannelId" in x ? true : false);
        A.store.Bool(ptr + 0, x["includeTlsChannelId"] ? true : false);
        A.store.Ref(ptr + 4, x["name"]);
      }
    },
    "load_ConnectArgConnectInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["includeTlsChannelId"] = A.load.Bool(ptr + 0);
      } else {
        delete x["includeTlsChannelId"];
      }
      x["name"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ContextType": (ref: heap.Ref<string>): number => {
      const idx = ["TAB", "POPUP", "BACKGROUND", "OFFSCREEN_DOCUMENT", "SIDE_PANEL"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ContextFilter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 37, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Bool(ptr + 36, false);
        A.store.Bool(ptr + 24, false);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
      } else {
        A.store.Bool(ptr + 37, true);
        A.store.Ref(ptr + 0, x["contextIds"]);
        A.store.Ref(ptr + 4, x["contextTypes"]);
        A.store.Ref(ptr + 8, x["documentIds"]);
        A.store.Ref(ptr + 12, x["documentOrigins"]);
        A.store.Ref(ptr + 16, x["documentUrls"]);
        A.store.Ref(ptr + 20, x["frameIds"]);
        A.store.Bool(ptr + 36, "incognito" in x ? true : false);
        A.store.Bool(ptr + 24, x["incognito"] ? true : false);
        A.store.Ref(ptr + 28, x["tabIds"]);
        A.store.Ref(ptr + 32, x["windowIds"]);
      }
    },
    "load_ContextFilter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["contextIds"] = A.load.Ref(ptr + 0, undefined);
      x["contextTypes"] = A.load.Ref(ptr + 4, undefined);
      x["documentIds"] = A.load.Ref(ptr + 8, undefined);
      x["documentOrigins"] = A.load.Ref(ptr + 12, undefined);
      x["documentUrls"] = A.load.Ref(ptr + 16, undefined);
      x["frameIds"] = A.load.Ref(ptr + 20, undefined);
      if (A.load.Bool(ptr + 36)) {
        x["incognito"] = A.load.Bool(ptr + 24);
      } else {
        delete x["incognito"];
      }
      x["tabIds"] = A.load.Ref(ptr + 28, undefined);
      x["windowIds"] = A.load.Ref(ptr + 32, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ExtensionContext": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 56, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Int64(ptr + 24, 0);
        A.store.Bool(ptr + 32, false);
        A.store.Int64(ptr + 40, 0);
        A.store.Int64(ptr + 48, 0);
      } else {
        A.store.Bool(ptr + 56, true);
        A.store.Ref(ptr + 0, x["contextId"]);
        A.store.Enum(
          ptr + 4,
          ["TAB", "POPUP", "BACKGROUND", "OFFSCREEN_DOCUMENT", "SIDE_PANEL"].indexOf(x["contextType"] as string)
        );
        A.store.Ref(ptr + 8, x["documentId"]);
        A.store.Ref(ptr + 12, x["documentOrigin"]);
        A.store.Ref(ptr + 16, x["documentUrl"]);
        A.store.Int64(ptr + 24, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Bool(ptr + 32, x["incognito"] ? true : false);
        A.store.Int64(ptr + 40, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Int64(ptr + 48, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_ExtensionContext": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["contextId"] = A.load.Ref(ptr + 0, undefined);
      x["contextType"] = A.load.Enum(ptr + 4, ["TAB", "POPUP", "BACKGROUND", "OFFSCREEN_DOCUMENT", "SIDE_PANEL"]);
      x["documentId"] = A.load.Ref(ptr + 8, undefined);
      x["documentOrigin"] = A.load.Ref(ptr + 12, undefined);
      x["documentUrl"] = A.load.Ref(ptr + 16, undefined);
      x["frameId"] = A.load.Int64(ptr + 24);
      x["incognito"] = A.load.Bool(ptr + 32);
      x["tabId"] = A.load.Int64(ptr + 40);
      x["windowId"] = A.load.Int64(ptr + 48);
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

    "store_MessageSender": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 187, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 184, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Bool(ptr + 185, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Bool(ptr + 186, false);
        A.store.Int64(ptr + 24, 0);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);

        A.store.Bool(ptr + 48 + 0, false);
        A.store.Bool(ptr + 48 + 0, false);
        A.store.Bool(ptr + 48 + 0, false);
        A.store.Bool(ptr + 48 + 0, false);
        A.store.Bool(ptr + 48 + 0, false);
        A.store.Bool(ptr + 48 + 0, false);
        A.store.Ref(ptr + 48 + 0, undefined);
        A.store.Int0(ptr + 48 + 0, 0);
        A.store.Bool(ptr + 48 + 0, false);
        A.store.Int0(ptr + 48 + 0, 0);
        A.store.Bool(ptr + 48 + 0, false);
        A.store.Bool(ptr + 48 + 0, false);
        A.store.Int0(ptr + 48 + 0, 0);
        A.store.Bool(ptr + 48 + 0, false);
        A.store.Int0(ptr + 48 + 0, 0);

        A.store.Bool(ptr + 48 + 0 + 0, false);
        A.store.Ref(ptr + 48 + 0 + 0, undefined);
        A.store.Bool(ptr + 48 + 0 + 0, false);
        A.store.Enum(ptr + 48 + 0 + 0, -1);
        A.store.Bool(ptr + 48 + 0, false);
        A.store.Int0(ptr + 48 + 0, 0);
        A.store.Ref(ptr + 48 + 0, undefined);
        A.store.Bool(ptr + 48 + 0, false);
        A.store.Bool(ptr + 48 + 0, false);
        A.store.Ref(ptr + 48 + 0, undefined);
        A.store.Enum(ptr + 48 + 0, -1);
        A.store.Ref(ptr + 48 + 0, undefined);
        A.store.Ref(ptr + 48 + 0, undefined);
        A.store.Bool(ptr + 48 + 0, false);
        A.store.Int0(ptr + 48 + 0, 0);
        A.store.Int0(ptr + 48 + 0, 0);
        A.store.Ref(ptr + 176, undefined);
        A.store.Ref(ptr + 180, undefined);
      } else {
        A.store.Bool(ptr + 187, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Ref(ptr + 4, x["documentLifecycle"]);
        A.store.Bool(ptr + 184, "frameId" in x ? true : false);
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Bool(ptr + 185, "guestProcessId" in x ? true : false);
        A.store.Int64(ptr + 16, x["guestProcessId"] === undefined ? 0 : (x["guestProcessId"] as number));
        A.store.Bool(ptr + 186, "guestRenderFrameRoutingId" in x ? true : false);
        A.store.Int64(
          ptr + 24,
          x["guestRenderFrameRoutingId"] === undefined ? 0 : (x["guestRenderFrameRoutingId"] as number)
        );
        A.store.Ref(ptr + 32, x["id"]);
        A.store.Ref(ptr + 36, x["nativeApplication"]);
        A.store.Ref(ptr + 40, x["origin"]);

        if (typeof x["tab"] === "undefined") {
          A.store.Bool(ptr + 48 + 0, false);
          A.store.Bool(ptr + 48 + 0, false);
          A.store.Bool(ptr + 48 + 0, false);
          A.store.Bool(ptr + 48 + 0, false);
          A.store.Bool(ptr + 48 + 0, false);
          A.store.Bool(ptr + 48 + 0, false);
          A.store.Ref(ptr + 48 + 0, undefined);
          A.store.Int0(ptr + 48 + 0, 0);
          A.store.Bool(ptr + 48 + 0, false);
          A.store.Int0(ptr + 48 + 0, 0);
          A.store.Bool(ptr + 48 + 0, false);
          A.store.Bool(ptr + 48 + 0, false);
          A.store.Int0(ptr + 48 + 0, 0);
          A.store.Bool(ptr + 48 + 0, false);
          A.store.Int0(ptr + 48 + 0, 0);

          A.store.Bool(ptr + 48 + 0 + 0, false);
          A.store.Ref(ptr + 48 + 0 + 0, undefined);
          A.store.Bool(ptr + 48 + 0 + 0, false);
          A.store.Enum(ptr + 48 + 0 + 0, -1);
          A.store.Bool(ptr + 48 + 0, false);
          A.store.Int0(ptr + 48 + 0, 0);
          A.store.Ref(ptr + 48 + 0, undefined);
          A.store.Bool(ptr + 48 + 0, false);
          A.store.Bool(ptr + 48 + 0, false);
          A.store.Ref(ptr + 48 + 0, undefined);
          A.store.Enum(ptr + 48 + 0, -1);
          A.store.Ref(ptr + 48 + 0, undefined);
          A.store.Ref(ptr + 48 + 0, undefined);
          A.store.Bool(ptr + 48 + 0, false);
          A.store.Int0(ptr + 48 + 0, 0);
          A.store.Int0(ptr + 48 + 0, 0);
        } else {
          A.store.Bool(ptr + 48 + 0, true);
          A.store.Bool(ptr + 48 + 0, x["tab"]["active"] ? true : false);
          A.store.Bool(ptr + 48 + 0, "audible" in x["tab"] ? true : false);
          A.store.Bool(ptr + 48 + 0, x["tab"]["audible"] ? true : false);
          A.store.Bool(ptr + 48 + 0, x["tab"]["autoDiscardable"] ? true : false);
          A.store.Bool(ptr + 48 + 0, x["tab"]["discarded"] ? true : false);
          A.store.Ref(ptr + 48 + 0, x["tab"]["favIconUrl"]);
          A.store.Int0(ptr + 48 + 0, x["tab"]["groupId"] === undefined ? 0 : (x["tab"]["groupId"] as number));
          A.store.Bool(ptr + 48 + 0, "height" in x["tab"] ? true : false);
          A.store.Int0(ptr + 48 + 0, x["tab"]["height"] === undefined ? 0 : (x["tab"]["height"] as number));
          A.store.Bool(ptr + 48 + 0, x["tab"]["highlighted"] ? true : false);
          A.store.Bool(ptr + 48 + 0, "id" in x["tab"] ? true : false);
          A.store.Int0(ptr + 48 + 0, x["tab"]["id"] === undefined ? 0 : (x["tab"]["id"] as number));
          A.store.Bool(ptr + 48 + 0, x["tab"]["incognito"] ? true : false);
          A.store.Int0(ptr + 48 + 0, x["tab"]["index"] === undefined ? 0 : (x["tab"]["index"] as number));

          if (typeof x["tab"]["mutedInfo"] === "undefined") {
            A.store.Bool(ptr + 48 + 0 + 0, false);
            A.store.Ref(ptr + 48 + 0 + 0, undefined);
            A.store.Bool(ptr + 48 + 0 + 0, false);
            A.store.Enum(ptr + 48 + 0 + 0, -1);
          } else {
            A.store.Bool(ptr + 48 + 0 + 0, true);
            A.store.Ref(ptr + 48 + 0 + 0, x["tab"]["mutedInfo"]["extensionId"]);
            A.store.Bool(ptr + 48 + 0 + 0, x["tab"]["mutedInfo"]["muted"] ? true : false);
            A.store.Enum(
              ptr + 48 + 0 + 0,
              ["user", "capture", "extension"].indexOf(x["tab"]["mutedInfo"]["reason"] as string)
            );
          }
          A.store.Bool(ptr + 48 + 0, "openerTabId" in x["tab"] ? true : false);
          A.store.Int0(ptr + 48 + 0, x["tab"]["openerTabId"] === undefined ? 0 : (x["tab"]["openerTabId"] as number));
          A.store.Ref(ptr + 48 + 0, x["tab"]["pendingUrl"]);
          A.store.Bool(ptr + 48 + 0, x["tab"]["pinned"] ? true : false);
          A.store.Bool(ptr + 48 + 0, x["tab"]["selected"] ? true : false);
          A.store.Ref(ptr + 48 + 0, x["tab"]["sessionId"]);
          A.store.Enum(ptr + 48 + 0, ["unloaded", "loading", "complete"].indexOf(x["tab"]["status"] as string));
          A.store.Ref(ptr + 48 + 0, x["tab"]["title"]);
          A.store.Ref(ptr + 48 + 0, x["tab"]["url"]);
          A.store.Bool(ptr + 48 + 0, "width" in x["tab"] ? true : false);
          A.store.Int0(ptr + 48 + 0, x["tab"]["width"] === undefined ? 0 : (x["tab"]["width"] as number));
          A.store.Int0(ptr + 48 + 0, x["tab"]["windowId"] === undefined ? 0 : (x["tab"]["windowId"] as number));
        }
        A.store.Ref(ptr + 176, x["tlsChannelId"]);
        A.store.Ref(ptr + 180, x["url"]);
      }
    },
    "load_MessageSender": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 184)) {
        x["frameId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["frameId"];
      }
      if (A.load.Bool(ptr + 185)) {
        x["guestProcessId"] = A.load.Int64(ptr + 16);
      } else {
        delete x["guestProcessId"];
      }
      if (A.load.Bool(ptr + 186)) {
        x["guestRenderFrameRoutingId"] = A.load.Int64(ptr + 24);
      } else {
        delete x["guestRenderFrameRoutingId"];
      }
      x["id"] = A.load.Ref(ptr + 32, undefined);
      x["nativeApplication"] = A.load.Ref(ptr + 36, undefined);
      x["origin"] = A.load.Ref(ptr + 40, undefined);
      if (A.load.Bool(ptr + 48 + 0)) {
        x["tab"] = {};
        x["tab"]["active"] = A.load.Bool(ptr + 48 + 0);
        if (A.load.Bool(ptr + 48 + 0)) {
          x["tab"]["audible"] = A.load.Bool(ptr + 48 + 0);
        } else {
          delete x["tab"]["audible"];
        }
        x["tab"]["autoDiscardable"] = A.load.Bool(ptr + 48 + 0);
        x["tab"]["discarded"] = A.load.Bool(ptr + 48 + 0);
        x["tab"]["favIconUrl"] = A.load.Ref(ptr + 48 + 0, undefined);
        x["tab"]["groupId"] = A.load.Int0(ptr + 48 + 0);
        if (A.load.Bool(ptr + 48 + 0)) {
          x["tab"]["height"] = A.load.Int0(ptr + 48 + 0);
        } else {
          delete x["tab"]["height"];
        }
        x["tab"]["highlighted"] = A.load.Bool(ptr + 48 + 0);
        if (A.load.Bool(ptr + 48 + 0)) {
          x["tab"]["id"] = A.load.Int0(ptr + 48 + 0);
        } else {
          delete x["tab"]["id"];
        }
        x["tab"]["incognito"] = A.load.Bool(ptr + 48 + 0);
        x["tab"]["index"] = A.load.Int0(ptr + 48 + 0);
        if (A.load.Bool(ptr + 48 + 0 + 0)) {
          x["tab"]["mutedInfo"] = {};
          x["tab"]["mutedInfo"]["extensionId"] = A.load.Ref(ptr + 48 + 0 + 0, undefined);
          x["tab"]["mutedInfo"]["muted"] = A.load.Bool(ptr + 48 + 0 + 0);
          x["tab"]["mutedInfo"]["reason"] = A.load.Enum(ptr + 48 + 0 + 0, ["user", "capture", "extension"]);
        } else {
          delete x["tab"]["mutedInfo"];
        }
        if (A.load.Bool(ptr + 48 + 0)) {
          x["tab"]["openerTabId"] = A.load.Int0(ptr + 48 + 0);
        } else {
          delete x["tab"]["openerTabId"];
        }
        x["tab"]["pendingUrl"] = A.load.Ref(ptr + 48 + 0, undefined);
        x["tab"]["pinned"] = A.load.Bool(ptr + 48 + 0);
        x["tab"]["selected"] = A.load.Bool(ptr + 48 + 0);
        x["tab"]["sessionId"] = A.load.Ref(ptr + 48 + 0, undefined);
        x["tab"]["status"] = A.load.Enum(ptr + 48 + 0, ["unloaded", "loading", "complete"]);
        x["tab"]["title"] = A.load.Ref(ptr + 48 + 0, undefined);
        x["tab"]["url"] = A.load.Ref(ptr + 48 + 0, undefined);
        if (A.load.Bool(ptr + 48 + 0)) {
          x["tab"]["width"] = A.load.Int0(ptr + 48 + 0);
        } else {
          delete x["tab"]["width"];
        }
        x["tab"]["windowId"] = A.load.Int0(ptr + 48 + 0);
      } else {
        delete x["tab"];
      }
      x["tlsChannelId"] = A.load.Ref(ptr + 176, undefined);
      x["url"] = A.load.Ref(ptr + 180, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OnInstalledReason": (ref: heap.Ref<string>): number => {
      const idx = ["install", "update", "chrome_update", "shared_module_update"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OnInstalledArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["previousVersion"]);
        A.store.Enum(
          ptr + 8,
          ["install", "update", "chrome_update", "shared_module_update"].indexOf(x["reason"] as string)
        );
      }
    },
    "load_OnInstalledArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["previousVersion"] = A.load.Ref(ptr + 4, undefined);
      x["reason"] = A.load.Enum(ptr + 8, ["install", "update", "chrome_update", "shared_module_update"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OnRestartRequiredReason": (ref: heap.Ref<string>): number => {
      const idx = ["app_update", "os_update", "periodic"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PlatformArch": (ref: heap.Ref<string>): number => {
      const idx = ["arm", "arm64", "x86-32", "x86-64", "mips", "mips64"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PlatformNaclArch": (ref: heap.Ref<string>): number => {
      const idx = ["arm", "x86-32", "x86-64", "mips", "mips64"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PlatformOs": (ref: heap.Ref<string>): number => {
      const idx = ["mac", "win", "android", "cros", "linux", "openbsd", "fuchsia"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PlatformInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Enum(ptr + 4, -1);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Enum(ptr + 0, ["arm", "arm64", "x86-32", "x86-64", "mips", "mips64"].indexOf(x["arch"] as string));
        A.store.Enum(ptr + 4, ["arm", "x86-32", "x86-64", "mips", "mips64"].indexOf(x["nacl_arch"] as string));
        A.store.Enum(
          ptr + 8,
          ["mac", "win", "android", "cros", "linux", "openbsd", "fuchsia"].indexOf(x["os"] as string)
        );
      }
    },
    "load_PlatformInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["arch"] = A.load.Enum(ptr + 0, ["arm", "arm64", "x86-32", "x86-64", "mips", "mips64"]);
      x["nacl_arch"] = A.load.Enum(ptr + 4, ["arm", "x86-32", "x86-64", "mips", "mips64"]);
      x["os"] = A.load.Enum(ptr + 8, ["mac", "win", "android", "cros", "linux", "openbsd", "fuchsia"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Port": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 204, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);

        A.store.Bool(ptr + 16 + 187, false);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 4, undefined);
        A.store.Bool(ptr + 16 + 184, false);
        A.store.Int64(ptr + 16 + 8, 0);
        A.store.Bool(ptr + 16 + 185, false);
        A.store.Int64(ptr + 16 + 16, 0);
        A.store.Bool(ptr + 16 + 186, false);
        A.store.Int64(ptr + 16 + 24, 0);
        A.store.Ref(ptr + 16 + 32, undefined);
        A.store.Ref(ptr + 16 + 36, undefined);
        A.store.Ref(ptr + 16 + 40, undefined);

        A.store.Bool(ptr + 16 + 48 + 0, false);
        A.store.Bool(ptr + 16 + 48 + 0, false);
        A.store.Bool(ptr + 16 + 48 + 0, false);
        A.store.Bool(ptr + 16 + 48 + 0, false);
        A.store.Bool(ptr + 16 + 48 + 0, false);
        A.store.Bool(ptr + 16 + 48 + 0, false);
        A.store.Ref(ptr + 16 + 48 + 0, undefined);
        A.store.Int0(ptr + 16 + 48 + 0, 0);
        A.store.Bool(ptr + 16 + 48 + 0, false);
        A.store.Int0(ptr + 16 + 48 + 0, 0);
        A.store.Bool(ptr + 16 + 48 + 0, false);
        A.store.Bool(ptr + 16 + 48 + 0, false);
        A.store.Int0(ptr + 16 + 48 + 0, 0);
        A.store.Bool(ptr + 16 + 48 + 0, false);
        A.store.Int0(ptr + 16 + 48 + 0, 0);

        A.store.Bool(ptr + 16 + 48 + 0 + 0, false);
        A.store.Ref(ptr + 16 + 48 + 0 + 0, undefined);
        A.store.Bool(ptr + 16 + 48 + 0 + 0, false);
        A.store.Enum(ptr + 16 + 48 + 0 + 0, -1);
        A.store.Bool(ptr + 16 + 48 + 0, false);
        A.store.Int0(ptr + 16 + 48 + 0, 0);
        A.store.Ref(ptr + 16 + 48 + 0, undefined);
        A.store.Bool(ptr + 16 + 48 + 0, false);
        A.store.Bool(ptr + 16 + 48 + 0, false);
        A.store.Ref(ptr + 16 + 48 + 0, undefined);
        A.store.Enum(ptr + 16 + 48 + 0, -1);
        A.store.Ref(ptr + 16 + 48 + 0, undefined);
        A.store.Ref(ptr + 16 + 48 + 0, undefined);
        A.store.Bool(ptr + 16 + 48 + 0, false);
        A.store.Int0(ptr + 16 + 48 + 0, 0);
        A.store.Int0(ptr + 16 + 48 + 0, 0);
        A.store.Ref(ptr + 16 + 176, undefined);
        A.store.Ref(ptr + 16 + 180, undefined);
      } else {
        A.store.Bool(ptr + 204, true);
        A.store.Ref(ptr + 0, x["disconnect"]);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Ref(ptr + 8, x["postMessage"]);

        if (typeof x["sender"] === "undefined") {
          A.store.Bool(ptr + 16 + 187, false);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 4, undefined);
          A.store.Bool(ptr + 16 + 184, false);
          A.store.Int64(ptr + 16 + 8, 0);
          A.store.Bool(ptr + 16 + 185, false);
          A.store.Int64(ptr + 16 + 16, 0);
          A.store.Bool(ptr + 16 + 186, false);
          A.store.Int64(ptr + 16 + 24, 0);
          A.store.Ref(ptr + 16 + 32, undefined);
          A.store.Ref(ptr + 16 + 36, undefined);
          A.store.Ref(ptr + 16 + 40, undefined);

          A.store.Bool(ptr + 16 + 48 + 0, false);
          A.store.Bool(ptr + 16 + 48 + 0, false);
          A.store.Bool(ptr + 16 + 48 + 0, false);
          A.store.Bool(ptr + 16 + 48 + 0, false);
          A.store.Bool(ptr + 16 + 48 + 0, false);
          A.store.Bool(ptr + 16 + 48 + 0, false);
          A.store.Ref(ptr + 16 + 48 + 0, undefined);
          A.store.Int0(ptr + 16 + 48 + 0, 0);
          A.store.Bool(ptr + 16 + 48 + 0, false);
          A.store.Int0(ptr + 16 + 48 + 0, 0);
          A.store.Bool(ptr + 16 + 48 + 0, false);
          A.store.Bool(ptr + 16 + 48 + 0, false);
          A.store.Int0(ptr + 16 + 48 + 0, 0);
          A.store.Bool(ptr + 16 + 48 + 0, false);
          A.store.Int0(ptr + 16 + 48 + 0, 0);

          A.store.Bool(ptr + 16 + 48 + 0 + 0, false);
          A.store.Ref(ptr + 16 + 48 + 0 + 0, undefined);
          A.store.Bool(ptr + 16 + 48 + 0 + 0, false);
          A.store.Enum(ptr + 16 + 48 + 0 + 0, -1);
          A.store.Bool(ptr + 16 + 48 + 0, false);
          A.store.Int0(ptr + 16 + 48 + 0, 0);
          A.store.Ref(ptr + 16 + 48 + 0, undefined);
          A.store.Bool(ptr + 16 + 48 + 0, false);
          A.store.Bool(ptr + 16 + 48 + 0, false);
          A.store.Ref(ptr + 16 + 48 + 0, undefined);
          A.store.Enum(ptr + 16 + 48 + 0, -1);
          A.store.Ref(ptr + 16 + 48 + 0, undefined);
          A.store.Ref(ptr + 16 + 48 + 0, undefined);
          A.store.Bool(ptr + 16 + 48 + 0, false);
          A.store.Int0(ptr + 16 + 48 + 0, 0);
          A.store.Int0(ptr + 16 + 48 + 0, 0);
          A.store.Ref(ptr + 16 + 176, undefined);
          A.store.Ref(ptr + 16 + 180, undefined);
        } else {
          A.store.Bool(ptr + 16 + 187, true);
          A.store.Ref(ptr + 16 + 0, x["sender"]["documentId"]);
          A.store.Ref(ptr + 16 + 4, x["sender"]["documentLifecycle"]);
          A.store.Bool(ptr + 16 + 184, "frameId" in x["sender"] ? true : false);
          A.store.Int64(ptr + 16 + 8, x["sender"]["frameId"] === undefined ? 0 : (x["sender"]["frameId"] as number));
          A.store.Bool(ptr + 16 + 185, "guestProcessId" in x["sender"] ? true : false);
          A.store.Int64(
            ptr + 16 + 16,
            x["sender"]["guestProcessId"] === undefined ? 0 : (x["sender"]["guestProcessId"] as number)
          );
          A.store.Bool(ptr + 16 + 186, "guestRenderFrameRoutingId" in x["sender"] ? true : false);
          A.store.Int64(
            ptr + 16 + 24,
            x["sender"]["guestRenderFrameRoutingId"] === undefined
              ? 0
              : (x["sender"]["guestRenderFrameRoutingId"] as number)
          );
          A.store.Ref(ptr + 16 + 32, x["sender"]["id"]);
          A.store.Ref(ptr + 16 + 36, x["sender"]["nativeApplication"]);
          A.store.Ref(ptr + 16 + 40, x["sender"]["origin"]);

          if (typeof x["sender"]["tab"] === "undefined") {
            A.store.Bool(ptr + 16 + 48 + 0, false);
            A.store.Bool(ptr + 16 + 48 + 0, false);
            A.store.Bool(ptr + 16 + 48 + 0, false);
            A.store.Bool(ptr + 16 + 48 + 0, false);
            A.store.Bool(ptr + 16 + 48 + 0, false);
            A.store.Bool(ptr + 16 + 48 + 0, false);
            A.store.Ref(ptr + 16 + 48 + 0, undefined);
            A.store.Int0(ptr + 16 + 48 + 0, 0);
            A.store.Bool(ptr + 16 + 48 + 0, false);
            A.store.Int0(ptr + 16 + 48 + 0, 0);
            A.store.Bool(ptr + 16 + 48 + 0, false);
            A.store.Bool(ptr + 16 + 48 + 0, false);
            A.store.Int0(ptr + 16 + 48 + 0, 0);
            A.store.Bool(ptr + 16 + 48 + 0, false);
            A.store.Int0(ptr + 16 + 48 + 0, 0);

            A.store.Bool(ptr + 16 + 48 + 0 + 0, false);
            A.store.Ref(ptr + 16 + 48 + 0 + 0, undefined);
            A.store.Bool(ptr + 16 + 48 + 0 + 0, false);
            A.store.Enum(ptr + 16 + 48 + 0 + 0, -1);
            A.store.Bool(ptr + 16 + 48 + 0, false);
            A.store.Int0(ptr + 16 + 48 + 0, 0);
            A.store.Ref(ptr + 16 + 48 + 0, undefined);
            A.store.Bool(ptr + 16 + 48 + 0, false);
            A.store.Bool(ptr + 16 + 48 + 0, false);
            A.store.Ref(ptr + 16 + 48 + 0, undefined);
            A.store.Enum(ptr + 16 + 48 + 0, -1);
            A.store.Ref(ptr + 16 + 48 + 0, undefined);
            A.store.Ref(ptr + 16 + 48 + 0, undefined);
            A.store.Bool(ptr + 16 + 48 + 0, false);
            A.store.Int0(ptr + 16 + 48 + 0, 0);
            A.store.Int0(ptr + 16 + 48 + 0, 0);
          } else {
            A.store.Bool(ptr + 16 + 48 + 0, true);
            A.store.Bool(ptr + 16 + 48 + 0, x["sender"]["tab"]["active"] ? true : false);
            A.store.Bool(ptr + 16 + 48 + 0, "audible" in x["sender"]["tab"] ? true : false);
            A.store.Bool(ptr + 16 + 48 + 0, x["sender"]["tab"]["audible"] ? true : false);
            A.store.Bool(ptr + 16 + 48 + 0, x["sender"]["tab"]["autoDiscardable"] ? true : false);
            A.store.Bool(ptr + 16 + 48 + 0, x["sender"]["tab"]["discarded"] ? true : false);
            A.store.Ref(ptr + 16 + 48 + 0, x["sender"]["tab"]["favIconUrl"]);
            A.store.Int0(
              ptr + 16 + 48 + 0,
              x["sender"]["tab"]["groupId"] === undefined ? 0 : (x["sender"]["tab"]["groupId"] as number)
            );
            A.store.Bool(ptr + 16 + 48 + 0, "height" in x["sender"]["tab"] ? true : false);
            A.store.Int0(
              ptr + 16 + 48 + 0,
              x["sender"]["tab"]["height"] === undefined ? 0 : (x["sender"]["tab"]["height"] as number)
            );
            A.store.Bool(ptr + 16 + 48 + 0, x["sender"]["tab"]["highlighted"] ? true : false);
            A.store.Bool(ptr + 16 + 48 + 0, "id" in x["sender"]["tab"] ? true : false);
            A.store.Int0(
              ptr + 16 + 48 + 0,
              x["sender"]["tab"]["id"] === undefined ? 0 : (x["sender"]["tab"]["id"] as number)
            );
            A.store.Bool(ptr + 16 + 48 + 0, x["sender"]["tab"]["incognito"] ? true : false);
            A.store.Int0(
              ptr + 16 + 48 + 0,
              x["sender"]["tab"]["index"] === undefined ? 0 : (x["sender"]["tab"]["index"] as number)
            );

            if (typeof x["sender"]["tab"]["mutedInfo"] === "undefined") {
              A.store.Bool(ptr + 16 + 48 + 0 + 0, false);
              A.store.Ref(ptr + 16 + 48 + 0 + 0, undefined);
              A.store.Bool(ptr + 16 + 48 + 0 + 0, false);
              A.store.Enum(ptr + 16 + 48 + 0 + 0, -1);
            } else {
              A.store.Bool(ptr + 16 + 48 + 0 + 0, true);
              A.store.Ref(ptr + 16 + 48 + 0 + 0, x["sender"]["tab"]["mutedInfo"]["extensionId"]);
              A.store.Bool(ptr + 16 + 48 + 0 + 0, x["sender"]["tab"]["mutedInfo"]["muted"] ? true : false);
              A.store.Enum(
                ptr + 16 + 48 + 0 + 0,
                ["user", "capture", "extension"].indexOf(x["sender"]["tab"]["mutedInfo"]["reason"] as string)
              );
            }
            A.store.Bool(ptr + 16 + 48 + 0, "openerTabId" in x["sender"]["tab"] ? true : false);
            A.store.Int0(
              ptr + 16 + 48 + 0,
              x["sender"]["tab"]["openerTabId"] === undefined ? 0 : (x["sender"]["tab"]["openerTabId"] as number)
            );
            A.store.Ref(ptr + 16 + 48 + 0, x["sender"]["tab"]["pendingUrl"]);
            A.store.Bool(ptr + 16 + 48 + 0, x["sender"]["tab"]["pinned"] ? true : false);
            A.store.Bool(ptr + 16 + 48 + 0, x["sender"]["tab"]["selected"] ? true : false);
            A.store.Ref(ptr + 16 + 48 + 0, x["sender"]["tab"]["sessionId"]);
            A.store.Enum(
              ptr + 16 + 48 + 0,
              ["unloaded", "loading", "complete"].indexOf(x["sender"]["tab"]["status"] as string)
            );
            A.store.Ref(ptr + 16 + 48 + 0, x["sender"]["tab"]["title"]);
            A.store.Ref(ptr + 16 + 48 + 0, x["sender"]["tab"]["url"]);
            A.store.Bool(ptr + 16 + 48 + 0, "width" in x["sender"]["tab"] ? true : false);
            A.store.Int0(
              ptr + 16 + 48 + 0,
              x["sender"]["tab"]["width"] === undefined ? 0 : (x["sender"]["tab"]["width"] as number)
            );
            A.store.Int0(
              ptr + 16 + 48 + 0,
              x["sender"]["tab"]["windowId"] === undefined ? 0 : (x["sender"]["tab"]["windowId"] as number)
            );
          }
          A.store.Ref(ptr + 16 + 176, x["sender"]["tlsChannelId"]);
          A.store.Ref(ptr + 16 + 180, x["sender"]["url"]);
        }
      }
    },
    "load_Port": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["disconnect"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      x["postMessage"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 16 + 187)) {
        x["sender"] = {};
        x["sender"]["documentId"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["sender"]["documentLifecycle"] = A.load.Ref(ptr + 16 + 4, undefined);
        if (A.load.Bool(ptr + 16 + 184)) {
          x["sender"]["frameId"] = A.load.Int64(ptr + 16 + 8);
        } else {
          delete x["sender"]["frameId"];
        }
        if (A.load.Bool(ptr + 16 + 185)) {
          x["sender"]["guestProcessId"] = A.load.Int64(ptr + 16 + 16);
        } else {
          delete x["sender"]["guestProcessId"];
        }
        if (A.load.Bool(ptr + 16 + 186)) {
          x["sender"]["guestRenderFrameRoutingId"] = A.load.Int64(ptr + 16 + 24);
        } else {
          delete x["sender"]["guestRenderFrameRoutingId"];
        }
        x["sender"]["id"] = A.load.Ref(ptr + 16 + 32, undefined);
        x["sender"]["nativeApplication"] = A.load.Ref(ptr + 16 + 36, undefined);
        x["sender"]["origin"] = A.load.Ref(ptr + 16 + 40, undefined);
        if (A.load.Bool(ptr + 16 + 48 + 0)) {
          x["sender"]["tab"] = {};
          x["sender"]["tab"]["active"] = A.load.Bool(ptr + 16 + 48 + 0);
          if (A.load.Bool(ptr + 16 + 48 + 0)) {
            x["sender"]["tab"]["audible"] = A.load.Bool(ptr + 16 + 48 + 0);
          } else {
            delete x["sender"]["tab"]["audible"];
          }
          x["sender"]["tab"]["autoDiscardable"] = A.load.Bool(ptr + 16 + 48 + 0);
          x["sender"]["tab"]["discarded"] = A.load.Bool(ptr + 16 + 48 + 0);
          x["sender"]["tab"]["favIconUrl"] = A.load.Ref(ptr + 16 + 48 + 0, undefined);
          x["sender"]["tab"]["groupId"] = A.load.Int0(ptr + 16 + 48 + 0);
          if (A.load.Bool(ptr + 16 + 48 + 0)) {
            x["sender"]["tab"]["height"] = A.load.Int0(ptr + 16 + 48 + 0);
          } else {
            delete x["sender"]["tab"]["height"];
          }
          x["sender"]["tab"]["highlighted"] = A.load.Bool(ptr + 16 + 48 + 0);
          if (A.load.Bool(ptr + 16 + 48 + 0)) {
            x["sender"]["tab"]["id"] = A.load.Int0(ptr + 16 + 48 + 0);
          } else {
            delete x["sender"]["tab"]["id"];
          }
          x["sender"]["tab"]["incognito"] = A.load.Bool(ptr + 16 + 48 + 0);
          x["sender"]["tab"]["index"] = A.load.Int0(ptr + 16 + 48 + 0);
          if (A.load.Bool(ptr + 16 + 48 + 0 + 0)) {
            x["sender"]["tab"]["mutedInfo"] = {};
            x["sender"]["tab"]["mutedInfo"]["extensionId"] = A.load.Ref(ptr + 16 + 48 + 0 + 0, undefined);
            x["sender"]["tab"]["mutedInfo"]["muted"] = A.load.Bool(ptr + 16 + 48 + 0 + 0);
            x["sender"]["tab"]["mutedInfo"]["reason"] = A.load.Enum(ptr + 16 + 48 + 0 + 0, [
              "user",
              "capture",
              "extension",
            ]);
          } else {
            delete x["sender"]["tab"]["mutedInfo"];
          }
          if (A.load.Bool(ptr + 16 + 48 + 0)) {
            x["sender"]["tab"]["openerTabId"] = A.load.Int0(ptr + 16 + 48 + 0);
          } else {
            delete x["sender"]["tab"]["openerTabId"];
          }
          x["sender"]["tab"]["pendingUrl"] = A.load.Ref(ptr + 16 + 48 + 0, undefined);
          x["sender"]["tab"]["pinned"] = A.load.Bool(ptr + 16 + 48 + 0);
          x["sender"]["tab"]["selected"] = A.load.Bool(ptr + 16 + 48 + 0);
          x["sender"]["tab"]["sessionId"] = A.load.Ref(ptr + 16 + 48 + 0, undefined);
          x["sender"]["tab"]["status"] = A.load.Enum(ptr + 16 + 48 + 0, ["unloaded", "loading", "complete"]);
          x["sender"]["tab"]["title"] = A.load.Ref(ptr + 16 + 48 + 0, undefined);
          x["sender"]["tab"]["url"] = A.load.Ref(ptr + 16 + 48 + 0, undefined);
          if (A.load.Bool(ptr + 16 + 48 + 0)) {
            x["sender"]["tab"]["width"] = A.load.Int0(ptr + 16 + 48 + 0);
          } else {
            delete x["sender"]["tab"]["width"];
          }
          x["sender"]["tab"]["windowId"] = A.load.Int0(ptr + 16 + 48 + 0);
        } else {
          delete x["sender"]["tab"];
        }
        x["sender"]["tlsChannelId"] = A.load.Ref(ptr + 16 + 176, undefined);
        x["sender"]["url"] = A.load.Ref(ptr + 16 + 180, undefined);
      } else {
        delete x["sender"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RequestUpdateCheckStatus": (ref: heap.Ref<string>): number => {
      const idx = ["throttled", "no_update", "update_available"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RequestUpdateCheckReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["throttled", "no_update", "update_available"].indexOf(x["status"] as string));
        A.store.Ref(ptr + 4, x["version"]);
      }
    },
    "load_RequestUpdateCheckReturnType": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["status"] = A.load.Enum(ptr + 0, ["throttled", "no_update", "update_available"]);
      x["version"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SendMessageArgOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 1, "includeTlsChannelId" in x ? true : false);
        A.store.Bool(ptr + 0, x["includeTlsChannelId"] ? true : false);
      }
    },
    "load_SendMessageArgOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 1)) {
        x["includeTlsChannelId"] = A.load.Bool(ptr + 0);
      } else {
        delete x["includeTlsChannelId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Connect": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "connect" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Connect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.connect);
    },
    "call_Connect": (retPtr: Pointer, extensionId: heap.Ref<object>, connectInfo: Pointer): void => {
      const connectInfo_ffi = {};

      if (A.load.Bool(connectInfo + 8)) {
        connectInfo_ffi["includeTlsChannelId"] = A.load.Bool(connectInfo + 0);
      }
      connectInfo_ffi["name"] = A.load.Ref(connectInfo + 4, undefined);

      const _ret = WEBEXT.runtime.connect(A.H.get<object>(extensionId), connectInfo_ffi);
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

        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Ref(retPtr + 16 + 48 + 0, undefined);
        A.store.Int0(retPtr + 16 + 48 + 0, 0);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Int0(retPtr + 16 + 48 + 0, 0);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Int0(retPtr + 16 + 48 + 0, 0);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Int0(retPtr + 16 + 48 + 0, 0);

        A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
        A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
        A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
        A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Int0(retPtr + 16 + 48 + 0, 0);
        A.store.Ref(retPtr + 16 + 48 + 0, undefined);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Ref(retPtr + 16 + 48 + 0, undefined);
        A.store.Enum(retPtr + 16 + 48 + 0, -1);
        A.store.Ref(retPtr + 16 + 48 + 0, undefined);
        A.store.Ref(retPtr + 16 + 48 + 0, undefined);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Int0(retPtr + 16 + 48 + 0, 0);
        A.store.Int0(retPtr + 16 + 48 + 0, 0);
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

          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);

          A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
          A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
          A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
          A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Enum(retPtr + 16 + 48 + 0, -1);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
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
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);

            A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
            A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
            A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
            A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Enum(retPtr + 16 + 48 + 0, -1);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
          } else {
            A.store.Bool(retPtr + 16 + 48 + 0, true);
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["active"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 0, "audible" in _ret["sender"]["tab"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["audible"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["autoDiscardable"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["discarded"] ? true : false);
            A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["favIconUrl"]);
            A.store.Int0(
              retPtr + 16 + 48 + 0,
              _ret["sender"]["tab"]["groupId"] === undefined ? 0 : (_ret["sender"]["tab"]["groupId"] as number)
            );
            A.store.Bool(retPtr + 16 + 48 + 0, "height" in _ret["sender"]["tab"] ? true : false);
            A.store.Int0(
              retPtr + 16 + 48 + 0,
              _ret["sender"]["tab"]["height"] === undefined ? 0 : (_ret["sender"]["tab"]["height"] as number)
            );
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["highlighted"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 0, "id" in _ret["sender"]["tab"] ? true : false);
            A.store.Int0(
              retPtr + 16 + 48 + 0,
              _ret["sender"]["tab"]["id"] === undefined ? 0 : (_ret["sender"]["tab"]["id"] as number)
            );
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["incognito"] ? true : false);
            A.store.Int0(
              retPtr + 16 + 48 + 0,
              _ret["sender"]["tab"]["index"] === undefined ? 0 : (_ret["sender"]["tab"]["index"] as number)
            );

            if (typeof _ret["sender"]["tab"]["mutedInfo"] === "undefined") {
              A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
              A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
              A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
              A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
            } else {
              A.store.Bool(retPtr + 16 + 48 + 0 + 0, true);
              A.store.Ref(retPtr + 16 + 48 + 0 + 0, _ret["sender"]["tab"]["mutedInfo"]["extensionId"]);
              A.store.Bool(retPtr + 16 + 48 + 0 + 0, _ret["sender"]["tab"]["mutedInfo"]["muted"] ? true : false);
              A.store.Enum(
                retPtr + 16 + 48 + 0 + 0,
                ["user", "capture", "extension"].indexOf(_ret["sender"]["tab"]["mutedInfo"]["reason"] as string)
              );
            }
            A.store.Bool(retPtr + 16 + 48 + 0, "openerTabId" in _ret["sender"]["tab"] ? true : false);
            A.store.Int0(
              retPtr + 16 + 48 + 0,
              _ret["sender"]["tab"]["openerTabId"] === undefined ? 0 : (_ret["sender"]["tab"]["openerTabId"] as number)
            );
            A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["pendingUrl"]);
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["pinned"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["selected"] ? true : false);
            A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["sessionId"]);
            A.store.Enum(
              retPtr + 16 + 48 + 0,
              ["unloaded", "loading", "complete"].indexOf(_ret["sender"]["tab"]["status"] as string)
            );
            A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["title"]);
            A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["url"]);
            A.store.Bool(retPtr + 16 + 48 + 0, "width" in _ret["sender"]["tab"] ? true : false);
            A.store.Int0(
              retPtr + 16 + 48 + 0,
              _ret["sender"]["tab"]["width"] === undefined ? 0 : (_ret["sender"]["tab"]["width"] as number)
            );
            A.store.Int0(
              retPtr + 16 + 48 + 0,
              _ret["sender"]["tab"]["windowId"] === undefined ? 0 : (_ret["sender"]["tab"]["windowId"] as number)
            );
          }
          A.store.Ref(retPtr + 16 + 176, _ret["sender"]["tlsChannelId"]);
          A.store.Ref(retPtr + 16 + 180, _ret["sender"]["url"]);
        }
      }
    },
    "try_Connect": (
      retPtr: Pointer,
      errPtr: Pointer,
      extensionId: heap.Ref<object>,
      connectInfo: Pointer
    ): heap.Ref<boolean> => {
      try {
        const connectInfo_ffi = {};

        if (A.load.Bool(connectInfo + 8)) {
          connectInfo_ffi["includeTlsChannelId"] = A.load.Bool(connectInfo + 0);
        }
        connectInfo_ffi["name"] = A.load.Ref(connectInfo + 4, undefined);

        const _ret = WEBEXT.runtime.connect(A.H.get<object>(extensionId), connectInfo_ffi);
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

          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);

          A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
          A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
          A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
          A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Enum(retPtr + 16 + 48 + 0, -1);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
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

            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);

            A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
            A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
            A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
            A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Enum(retPtr + 16 + 48 + 0, -1);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
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
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Ref(retPtr + 16 + 48 + 0, undefined);
              A.store.Int0(retPtr + 16 + 48 + 0, 0);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Int0(retPtr + 16 + 48 + 0, 0);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Int0(retPtr + 16 + 48 + 0, 0);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Int0(retPtr + 16 + 48 + 0, 0);

              A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
              A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
              A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
              A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Int0(retPtr + 16 + 48 + 0, 0);
              A.store.Ref(retPtr + 16 + 48 + 0, undefined);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Ref(retPtr + 16 + 48 + 0, undefined);
              A.store.Enum(retPtr + 16 + 48 + 0, -1);
              A.store.Ref(retPtr + 16 + 48 + 0, undefined);
              A.store.Ref(retPtr + 16 + 48 + 0, undefined);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Int0(retPtr + 16 + 48 + 0, 0);
              A.store.Int0(retPtr + 16 + 48 + 0, 0);
            } else {
              A.store.Bool(retPtr + 16 + 48 + 0, true);
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["active"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 0, "audible" in _ret["sender"]["tab"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["audible"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["autoDiscardable"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["discarded"] ? true : false);
              A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["favIconUrl"]);
              A.store.Int0(
                retPtr + 16 + 48 + 0,
                _ret["sender"]["tab"]["groupId"] === undefined ? 0 : (_ret["sender"]["tab"]["groupId"] as number)
              );
              A.store.Bool(retPtr + 16 + 48 + 0, "height" in _ret["sender"]["tab"] ? true : false);
              A.store.Int0(
                retPtr + 16 + 48 + 0,
                _ret["sender"]["tab"]["height"] === undefined ? 0 : (_ret["sender"]["tab"]["height"] as number)
              );
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["highlighted"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 0, "id" in _ret["sender"]["tab"] ? true : false);
              A.store.Int0(
                retPtr + 16 + 48 + 0,
                _ret["sender"]["tab"]["id"] === undefined ? 0 : (_ret["sender"]["tab"]["id"] as number)
              );
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["incognito"] ? true : false);
              A.store.Int0(
                retPtr + 16 + 48 + 0,
                _ret["sender"]["tab"]["index"] === undefined ? 0 : (_ret["sender"]["tab"]["index"] as number)
              );

              if (typeof _ret["sender"]["tab"]["mutedInfo"] === "undefined") {
                A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
                A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
                A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
                A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
              } else {
                A.store.Bool(retPtr + 16 + 48 + 0 + 0, true);
                A.store.Ref(retPtr + 16 + 48 + 0 + 0, _ret["sender"]["tab"]["mutedInfo"]["extensionId"]);
                A.store.Bool(retPtr + 16 + 48 + 0 + 0, _ret["sender"]["tab"]["mutedInfo"]["muted"] ? true : false);
                A.store.Enum(
                  retPtr + 16 + 48 + 0 + 0,
                  ["user", "capture", "extension"].indexOf(_ret["sender"]["tab"]["mutedInfo"]["reason"] as string)
                );
              }
              A.store.Bool(retPtr + 16 + 48 + 0, "openerTabId" in _ret["sender"]["tab"] ? true : false);
              A.store.Int0(
                retPtr + 16 + 48 + 0,
                _ret["sender"]["tab"]["openerTabId"] === undefined
                  ? 0
                  : (_ret["sender"]["tab"]["openerTabId"] as number)
              );
              A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["pendingUrl"]);
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["pinned"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["selected"] ? true : false);
              A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["sessionId"]);
              A.store.Enum(
                retPtr + 16 + 48 + 0,
                ["unloaded", "loading", "complete"].indexOf(_ret["sender"]["tab"]["status"] as string)
              );
              A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["title"]);
              A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["url"]);
              A.store.Bool(retPtr + 16 + 48 + 0, "width" in _ret["sender"]["tab"] ? true : false);
              A.store.Int0(
                retPtr + 16 + 48 + 0,
                _ret["sender"]["tab"]["width"] === undefined ? 0 : (_ret["sender"]["tab"]["width"] as number)
              );
              A.store.Int0(
                retPtr + 16 + 48 + 0,
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
    "has_ConnectNative": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "connectNative" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ConnectNative": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.connectNative);
    },
    "call_ConnectNative": (retPtr: Pointer, application: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.connectNative(A.H.get<object>(application));
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

        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Ref(retPtr + 16 + 48 + 0, undefined);
        A.store.Int0(retPtr + 16 + 48 + 0, 0);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Int0(retPtr + 16 + 48 + 0, 0);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Int0(retPtr + 16 + 48 + 0, 0);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Int0(retPtr + 16 + 48 + 0, 0);

        A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
        A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
        A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
        A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Int0(retPtr + 16 + 48 + 0, 0);
        A.store.Ref(retPtr + 16 + 48 + 0, undefined);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Ref(retPtr + 16 + 48 + 0, undefined);
        A.store.Enum(retPtr + 16 + 48 + 0, -1);
        A.store.Ref(retPtr + 16 + 48 + 0, undefined);
        A.store.Ref(retPtr + 16 + 48 + 0, undefined);
        A.store.Bool(retPtr + 16 + 48 + 0, false);
        A.store.Int0(retPtr + 16 + 48 + 0, 0);
        A.store.Int0(retPtr + 16 + 48 + 0, 0);
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

          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);

          A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
          A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
          A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
          A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Enum(retPtr + 16 + 48 + 0, -1);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
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
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);

            A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
            A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
            A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
            A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Enum(retPtr + 16 + 48 + 0, -1);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
          } else {
            A.store.Bool(retPtr + 16 + 48 + 0, true);
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["active"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 0, "audible" in _ret["sender"]["tab"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["audible"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["autoDiscardable"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["discarded"] ? true : false);
            A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["favIconUrl"]);
            A.store.Int0(
              retPtr + 16 + 48 + 0,
              _ret["sender"]["tab"]["groupId"] === undefined ? 0 : (_ret["sender"]["tab"]["groupId"] as number)
            );
            A.store.Bool(retPtr + 16 + 48 + 0, "height" in _ret["sender"]["tab"] ? true : false);
            A.store.Int0(
              retPtr + 16 + 48 + 0,
              _ret["sender"]["tab"]["height"] === undefined ? 0 : (_ret["sender"]["tab"]["height"] as number)
            );
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["highlighted"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 0, "id" in _ret["sender"]["tab"] ? true : false);
            A.store.Int0(
              retPtr + 16 + 48 + 0,
              _ret["sender"]["tab"]["id"] === undefined ? 0 : (_ret["sender"]["tab"]["id"] as number)
            );
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["incognito"] ? true : false);
            A.store.Int0(
              retPtr + 16 + 48 + 0,
              _ret["sender"]["tab"]["index"] === undefined ? 0 : (_ret["sender"]["tab"]["index"] as number)
            );

            if (typeof _ret["sender"]["tab"]["mutedInfo"] === "undefined") {
              A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
              A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
              A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
              A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
            } else {
              A.store.Bool(retPtr + 16 + 48 + 0 + 0, true);
              A.store.Ref(retPtr + 16 + 48 + 0 + 0, _ret["sender"]["tab"]["mutedInfo"]["extensionId"]);
              A.store.Bool(retPtr + 16 + 48 + 0 + 0, _ret["sender"]["tab"]["mutedInfo"]["muted"] ? true : false);
              A.store.Enum(
                retPtr + 16 + 48 + 0 + 0,
                ["user", "capture", "extension"].indexOf(_ret["sender"]["tab"]["mutedInfo"]["reason"] as string)
              );
            }
            A.store.Bool(retPtr + 16 + 48 + 0, "openerTabId" in _ret["sender"]["tab"] ? true : false);
            A.store.Int0(
              retPtr + 16 + 48 + 0,
              _ret["sender"]["tab"]["openerTabId"] === undefined ? 0 : (_ret["sender"]["tab"]["openerTabId"] as number)
            );
            A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["pendingUrl"]);
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["pinned"] ? true : false);
            A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["selected"] ? true : false);
            A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["sessionId"]);
            A.store.Enum(
              retPtr + 16 + 48 + 0,
              ["unloaded", "loading", "complete"].indexOf(_ret["sender"]["tab"]["status"] as string)
            );
            A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["title"]);
            A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["url"]);
            A.store.Bool(retPtr + 16 + 48 + 0, "width" in _ret["sender"]["tab"] ? true : false);
            A.store.Int0(
              retPtr + 16 + 48 + 0,
              _ret["sender"]["tab"]["width"] === undefined ? 0 : (_ret["sender"]["tab"]["width"] as number)
            );
            A.store.Int0(
              retPtr + 16 + 48 + 0,
              _ret["sender"]["tab"]["windowId"] === undefined ? 0 : (_ret["sender"]["tab"]["windowId"] as number)
            );
          }
          A.store.Ref(retPtr + 16 + 176, _ret["sender"]["tlsChannelId"]);
          A.store.Ref(retPtr + 16 + 180, _ret["sender"]["url"]);
        }
      }
    },
    "try_ConnectNative": (retPtr: Pointer, errPtr: Pointer, application: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.connectNative(A.H.get<object>(application));
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

          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);

          A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
          A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
          A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
          A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Enum(retPtr + 16 + 48 + 0, -1);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Ref(retPtr + 16 + 48 + 0, undefined);
          A.store.Bool(retPtr + 16 + 48 + 0, false);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
          A.store.Int0(retPtr + 16 + 48 + 0, 0);
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

            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);

            A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
            A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
            A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
            A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Enum(retPtr + 16 + 48 + 0, -1);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Ref(retPtr + 16 + 48 + 0, undefined);
            A.store.Bool(retPtr + 16 + 48 + 0, false);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
            A.store.Int0(retPtr + 16 + 48 + 0, 0);
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
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Ref(retPtr + 16 + 48 + 0, undefined);
              A.store.Int0(retPtr + 16 + 48 + 0, 0);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Int0(retPtr + 16 + 48 + 0, 0);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Int0(retPtr + 16 + 48 + 0, 0);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Int0(retPtr + 16 + 48 + 0, 0);

              A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
              A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
              A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
              A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Int0(retPtr + 16 + 48 + 0, 0);
              A.store.Ref(retPtr + 16 + 48 + 0, undefined);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Ref(retPtr + 16 + 48 + 0, undefined);
              A.store.Enum(retPtr + 16 + 48 + 0, -1);
              A.store.Ref(retPtr + 16 + 48 + 0, undefined);
              A.store.Ref(retPtr + 16 + 48 + 0, undefined);
              A.store.Bool(retPtr + 16 + 48 + 0, false);
              A.store.Int0(retPtr + 16 + 48 + 0, 0);
              A.store.Int0(retPtr + 16 + 48 + 0, 0);
            } else {
              A.store.Bool(retPtr + 16 + 48 + 0, true);
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["active"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 0, "audible" in _ret["sender"]["tab"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["audible"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["autoDiscardable"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["discarded"] ? true : false);
              A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["favIconUrl"]);
              A.store.Int0(
                retPtr + 16 + 48 + 0,
                _ret["sender"]["tab"]["groupId"] === undefined ? 0 : (_ret["sender"]["tab"]["groupId"] as number)
              );
              A.store.Bool(retPtr + 16 + 48 + 0, "height" in _ret["sender"]["tab"] ? true : false);
              A.store.Int0(
                retPtr + 16 + 48 + 0,
                _ret["sender"]["tab"]["height"] === undefined ? 0 : (_ret["sender"]["tab"]["height"] as number)
              );
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["highlighted"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 0, "id" in _ret["sender"]["tab"] ? true : false);
              A.store.Int0(
                retPtr + 16 + 48 + 0,
                _ret["sender"]["tab"]["id"] === undefined ? 0 : (_ret["sender"]["tab"]["id"] as number)
              );
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["incognito"] ? true : false);
              A.store.Int0(
                retPtr + 16 + 48 + 0,
                _ret["sender"]["tab"]["index"] === undefined ? 0 : (_ret["sender"]["tab"]["index"] as number)
              );

              if (typeof _ret["sender"]["tab"]["mutedInfo"] === "undefined") {
                A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
                A.store.Ref(retPtr + 16 + 48 + 0 + 0, undefined);
                A.store.Bool(retPtr + 16 + 48 + 0 + 0, false);
                A.store.Enum(retPtr + 16 + 48 + 0 + 0, -1);
              } else {
                A.store.Bool(retPtr + 16 + 48 + 0 + 0, true);
                A.store.Ref(retPtr + 16 + 48 + 0 + 0, _ret["sender"]["tab"]["mutedInfo"]["extensionId"]);
                A.store.Bool(retPtr + 16 + 48 + 0 + 0, _ret["sender"]["tab"]["mutedInfo"]["muted"] ? true : false);
                A.store.Enum(
                  retPtr + 16 + 48 + 0 + 0,
                  ["user", "capture", "extension"].indexOf(_ret["sender"]["tab"]["mutedInfo"]["reason"] as string)
                );
              }
              A.store.Bool(retPtr + 16 + 48 + 0, "openerTabId" in _ret["sender"]["tab"] ? true : false);
              A.store.Int0(
                retPtr + 16 + 48 + 0,
                _ret["sender"]["tab"]["openerTabId"] === undefined
                  ? 0
                  : (_ret["sender"]["tab"]["openerTabId"] as number)
              );
              A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["pendingUrl"]);
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["pinned"] ? true : false);
              A.store.Bool(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["selected"] ? true : false);
              A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["sessionId"]);
              A.store.Enum(
                retPtr + 16 + 48 + 0,
                ["unloaded", "loading", "complete"].indexOf(_ret["sender"]["tab"]["status"] as string)
              );
              A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["title"]);
              A.store.Ref(retPtr + 16 + 48 + 0, _ret["sender"]["tab"]["url"]);
              A.store.Bool(retPtr + 16 + 48 + 0, "width" in _ret["sender"]["tab"] ? true : false);
              A.store.Int0(
                retPtr + 16 + 48 + 0,
                _ret["sender"]["tab"]["width"] === undefined ? 0 : (_ret["sender"]["tab"]["width"] as number)
              );
              A.store.Int0(
                retPtr + 16 + 48 + 0,
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
    "has_GetBackgroundPage": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "getBackgroundPage" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetBackgroundPage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.getBackgroundPage);
    },
    "call_GetBackgroundPage": (retPtr: Pointer): void => {
      const _ret = WEBEXT.runtime.getBackgroundPage();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetBackgroundPage": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.getBackgroundPage();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetContexts": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "getContexts" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetContexts": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.getContexts);
    },
    "call_GetContexts": (retPtr: Pointer, filter: Pointer): void => {
      const filter_ffi = {};

      filter_ffi["contextIds"] = A.load.Ref(filter + 0, undefined);
      filter_ffi["contextTypes"] = A.load.Ref(filter + 4, undefined);
      filter_ffi["documentIds"] = A.load.Ref(filter + 8, undefined);
      filter_ffi["documentOrigins"] = A.load.Ref(filter + 12, undefined);
      filter_ffi["documentUrls"] = A.load.Ref(filter + 16, undefined);
      filter_ffi["frameIds"] = A.load.Ref(filter + 20, undefined);
      if (A.load.Bool(filter + 36)) {
        filter_ffi["incognito"] = A.load.Bool(filter + 24);
      }
      filter_ffi["tabIds"] = A.load.Ref(filter + 28, undefined);
      filter_ffi["windowIds"] = A.load.Ref(filter + 32, undefined);

      const _ret = WEBEXT.runtime.getContexts(filter_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetContexts": (retPtr: Pointer, errPtr: Pointer, filter: Pointer): heap.Ref<boolean> => {
      try {
        const filter_ffi = {};

        filter_ffi["contextIds"] = A.load.Ref(filter + 0, undefined);
        filter_ffi["contextTypes"] = A.load.Ref(filter + 4, undefined);
        filter_ffi["documentIds"] = A.load.Ref(filter + 8, undefined);
        filter_ffi["documentOrigins"] = A.load.Ref(filter + 12, undefined);
        filter_ffi["documentUrls"] = A.load.Ref(filter + 16, undefined);
        filter_ffi["frameIds"] = A.load.Ref(filter + 20, undefined);
        if (A.load.Bool(filter + 36)) {
          filter_ffi["incognito"] = A.load.Bool(filter + 24);
        }
        filter_ffi["tabIds"] = A.load.Ref(filter + 28, undefined);
        filter_ffi["windowIds"] = A.load.Ref(filter + 32, undefined);

        const _ret = WEBEXT.runtime.getContexts(filter_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetManifest": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "getManifest" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetManifest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.getManifest);
    },
    "call_GetManifest": (retPtr: Pointer): void => {
      const _ret = WEBEXT.runtime.getManifest();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetManifest": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.getManifest();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPackageDirectoryEntry": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "getPackageDirectoryEntry" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPackageDirectoryEntry": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.getPackageDirectoryEntry);
    },
    "call_GetPackageDirectoryEntry": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.getPackageDirectoryEntry(A.H.get<object>(callback));
    },
    "try_GetPackageDirectoryEntry": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.getPackageDirectoryEntry(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPlatformInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "getPlatformInfo" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPlatformInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.getPlatformInfo);
    },
    "call_GetPlatformInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.runtime.getPlatformInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPlatformInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.getPlatformInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetURL": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "getURL" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetURL": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.getURL);
    },
    "call_GetURL": (retPtr: Pointer, path: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.getURL(A.H.get<object>(path));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetURL": (retPtr: Pointer, errPtr: Pointer, path: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.getURL(A.H.get<object>(path));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_Id": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "id" in WEBEXT?.runtime) {
        const val = WEBEXT.runtime.id;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Id": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.runtime, "id", A.H.get<object>(val), WEBEXT.runtime) ? A.H.TRUE : A.H.FALSE;
    },
    "get_LastError": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "lastError" in WEBEXT?.runtime) {
        const val = WEBEXT.runtime.lastError;
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
      return Reflect.set(WEBEXT.runtime, "lastError", val_ffi, WEBEXT.runtime) ? A.H.TRUE : A.H.FALSE;
    },
    "has_OnBrowserUpdateAvailable": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onBrowserUpdateAvailable && "addListener" in WEBEXT?.runtime?.onBrowserUpdateAvailable) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnBrowserUpdateAvailable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onBrowserUpdateAvailable.addListener);
    },
    "call_OnBrowserUpdateAvailable": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onBrowserUpdateAvailable.addListener(A.H.get<object>(callback));
    },
    "try_OnBrowserUpdateAvailable": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onBrowserUpdateAvailable.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffBrowserUpdateAvailable": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onBrowserUpdateAvailable && "removeListener" in WEBEXT?.runtime?.onBrowserUpdateAvailable) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffBrowserUpdateAvailable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onBrowserUpdateAvailable.removeListener);
    },
    "call_OffBrowserUpdateAvailable": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onBrowserUpdateAvailable.removeListener(A.H.get<object>(callback));
    },
    "try_OffBrowserUpdateAvailable": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onBrowserUpdateAvailable.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnBrowserUpdateAvailable": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onBrowserUpdateAvailable && "hasListener" in WEBEXT?.runtime?.onBrowserUpdateAvailable) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnBrowserUpdateAvailable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onBrowserUpdateAvailable.hasListener);
    },
    "call_HasOnBrowserUpdateAvailable": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onBrowserUpdateAvailable.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnBrowserUpdateAvailable": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onBrowserUpdateAvailable.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnConnect": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onConnect && "addListener" in WEBEXT?.runtime?.onConnect) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnConnect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onConnect.addListener);
    },
    "call_OnConnect": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onConnect.addListener(A.H.get<object>(callback));
    },
    "try_OnConnect": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onConnect.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffConnect": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onConnect && "removeListener" in WEBEXT?.runtime?.onConnect) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffConnect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onConnect.removeListener);
    },
    "call_OffConnect": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onConnect.removeListener(A.H.get<object>(callback));
    },
    "try_OffConnect": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onConnect.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnConnect": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onConnect && "hasListener" in WEBEXT?.runtime?.onConnect) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnConnect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onConnect.hasListener);
    },
    "call_HasOnConnect": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onConnect.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnConnect": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onConnect.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnConnectExternal": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onConnectExternal && "addListener" in WEBEXT?.runtime?.onConnectExternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnConnectExternal": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onConnectExternal.addListener);
    },
    "call_OnConnectExternal": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onConnectExternal.addListener(A.H.get<object>(callback));
    },
    "try_OnConnectExternal": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onConnectExternal.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffConnectExternal": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onConnectExternal && "removeListener" in WEBEXT?.runtime?.onConnectExternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffConnectExternal": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onConnectExternal.removeListener);
    },
    "call_OffConnectExternal": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onConnectExternal.removeListener(A.H.get<object>(callback));
    },
    "try_OffConnectExternal": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onConnectExternal.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnConnectExternal": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onConnectExternal && "hasListener" in WEBEXT?.runtime?.onConnectExternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnConnectExternal": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onConnectExternal.hasListener);
    },
    "call_HasOnConnectExternal": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onConnectExternal.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnConnectExternal": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onConnectExternal.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnConnectNative": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onConnectNative && "addListener" in WEBEXT?.runtime?.onConnectNative) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnConnectNative": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onConnectNative.addListener);
    },
    "call_OnConnectNative": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onConnectNative.addListener(A.H.get<object>(callback));
    },
    "try_OnConnectNative": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onConnectNative.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffConnectNative": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onConnectNative && "removeListener" in WEBEXT?.runtime?.onConnectNative) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffConnectNative": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onConnectNative.removeListener);
    },
    "call_OffConnectNative": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onConnectNative.removeListener(A.H.get<object>(callback));
    },
    "try_OffConnectNative": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onConnectNative.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnConnectNative": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onConnectNative && "hasListener" in WEBEXT?.runtime?.onConnectNative) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnConnectNative": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onConnectNative.hasListener);
    },
    "call_HasOnConnectNative": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onConnectNative.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnConnectNative": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onConnectNative.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnInstalled": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onInstalled && "addListener" in WEBEXT?.runtime?.onInstalled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnInstalled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onInstalled.addListener);
    },
    "call_OnInstalled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onInstalled.addListener(A.H.get<object>(callback));
    },
    "try_OnInstalled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onInstalled.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffInstalled": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onInstalled && "removeListener" in WEBEXT?.runtime?.onInstalled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffInstalled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onInstalled.removeListener);
    },
    "call_OffInstalled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onInstalled.removeListener(A.H.get<object>(callback));
    },
    "try_OffInstalled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onInstalled.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnInstalled": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onInstalled && "hasListener" in WEBEXT?.runtime?.onInstalled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnInstalled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onInstalled.hasListener);
    },
    "call_HasOnInstalled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onInstalled.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnInstalled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onInstalled.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onMessage && "addListener" in WEBEXT?.runtime?.onMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onMessage.addListener);
    },
    "call_OnMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onMessage.addListener(A.H.get<object>(callback));
    },
    "try_OnMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onMessage.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onMessage && "removeListener" in WEBEXT?.runtime?.onMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onMessage.removeListener);
    },
    "call_OffMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onMessage.removeListener(A.H.get<object>(callback));
    },
    "try_OffMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onMessage.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onMessage && "hasListener" in WEBEXT?.runtime?.onMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onMessage.hasListener);
    },
    "call_HasOnMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onMessage.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onMessage.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMessageExternal": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onMessageExternal && "addListener" in WEBEXT?.runtime?.onMessageExternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMessageExternal": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onMessageExternal.addListener);
    },
    "call_OnMessageExternal": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onMessageExternal.addListener(A.H.get<object>(callback));
    },
    "try_OnMessageExternal": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onMessageExternal.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMessageExternal": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onMessageExternal && "removeListener" in WEBEXT?.runtime?.onMessageExternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMessageExternal": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onMessageExternal.removeListener);
    },
    "call_OffMessageExternal": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onMessageExternal.removeListener(A.H.get<object>(callback));
    },
    "try_OffMessageExternal": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onMessageExternal.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMessageExternal": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onMessageExternal && "hasListener" in WEBEXT?.runtime?.onMessageExternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMessageExternal": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onMessageExternal.hasListener);
    },
    "call_HasOnMessageExternal": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onMessageExternal.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMessageExternal": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onMessageExternal.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRestartRequired": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onRestartRequired && "addListener" in WEBEXT?.runtime?.onRestartRequired) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRestartRequired": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onRestartRequired.addListener);
    },
    "call_OnRestartRequired": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onRestartRequired.addListener(A.H.get<object>(callback));
    },
    "try_OnRestartRequired": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onRestartRequired.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRestartRequired": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onRestartRequired && "removeListener" in WEBEXT?.runtime?.onRestartRequired) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRestartRequired": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onRestartRequired.removeListener);
    },
    "call_OffRestartRequired": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onRestartRequired.removeListener(A.H.get<object>(callback));
    },
    "try_OffRestartRequired": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onRestartRequired.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRestartRequired": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onRestartRequired && "hasListener" in WEBEXT?.runtime?.onRestartRequired) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRestartRequired": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onRestartRequired.hasListener);
    },
    "call_HasOnRestartRequired": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onRestartRequired.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRestartRequired": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onRestartRequired.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnStartup": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onStartup && "addListener" in WEBEXT?.runtime?.onStartup) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnStartup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onStartup.addListener);
    },
    "call_OnStartup": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onStartup.addListener(A.H.get<object>(callback));
    },
    "try_OnStartup": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onStartup.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffStartup": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onStartup && "removeListener" in WEBEXT?.runtime?.onStartup) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffStartup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onStartup.removeListener);
    },
    "call_OffStartup": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onStartup.removeListener(A.H.get<object>(callback));
    },
    "try_OffStartup": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onStartup.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnStartup": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onStartup && "hasListener" in WEBEXT?.runtime?.onStartup) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnStartup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onStartup.hasListener);
    },
    "call_HasOnStartup": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onStartup.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnStartup": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onStartup.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSuspend": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onSuspend && "addListener" in WEBEXT?.runtime?.onSuspend) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSuspend": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onSuspend.addListener);
    },
    "call_OnSuspend": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onSuspend.addListener(A.H.get<object>(callback));
    },
    "try_OnSuspend": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onSuspend.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSuspend": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onSuspend && "removeListener" in WEBEXT?.runtime?.onSuspend) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSuspend": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onSuspend.removeListener);
    },
    "call_OffSuspend": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onSuspend.removeListener(A.H.get<object>(callback));
    },
    "try_OffSuspend": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onSuspend.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSuspend": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onSuspend && "hasListener" in WEBEXT?.runtime?.onSuspend) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSuspend": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onSuspend.hasListener);
    },
    "call_HasOnSuspend": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onSuspend.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSuspend": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onSuspend.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSuspendCanceled": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onSuspendCanceled && "addListener" in WEBEXT?.runtime?.onSuspendCanceled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSuspendCanceled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onSuspendCanceled.addListener);
    },
    "call_OnSuspendCanceled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onSuspendCanceled.addListener(A.H.get<object>(callback));
    },
    "try_OnSuspendCanceled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onSuspendCanceled.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSuspendCanceled": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onSuspendCanceled && "removeListener" in WEBEXT?.runtime?.onSuspendCanceled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSuspendCanceled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onSuspendCanceled.removeListener);
    },
    "call_OffSuspendCanceled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onSuspendCanceled.removeListener(A.H.get<object>(callback));
    },
    "try_OffSuspendCanceled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onSuspendCanceled.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSuspendCanceled": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onSuspendCanceled && "hasListener" in WEBEXT?.runtime?.onSuspendCanceled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSuspendCanceled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onSuspendCanceled.hasListener);
    },
    "call_HasOnSuspendCanceled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onSuspendCanceled.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSuspendCanceled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onSuspendCanceled.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnUpdateAvailable": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onUpdateAvailable && "addListener" in WEBEXT?.runtime?.onUpdateAvailable) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnUpdateAvailable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onUpdateAvailable.addListener);
    },
    "call_OnUpdateAvailable": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onUpdateAvailable.addListener(A.H.get<object>(callback));
    },
    "try_OnUpdateAvailable": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onUpdateAvailable.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffUpdateAvailable": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onUpdateAvailable && "removeListener" in WEBEXT?.runtime?.onUpdateAvailable) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffUpdateAvailable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onUpdateAvailable.removeListener);
    },
    "call_OffUpdateAvailable": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onUpdateAvailable.removeListener(A.H.get<object>(callback));
    },
    "try_OffUpdateAvailable": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onUpdateAvailable.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnUpdateAvailable": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onUpdateAvailable && "hasListener" in WEBEXT?.runtime?.onUpdateAvailable) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnUpdateAvailable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onUpdateAvailable.hasListener);
    },
    "call_HasOnUpdateAvailable": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onUpdateAvailable.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnUpdateAvailable": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onUpdateAvailable.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnUserScriptConnect": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onUserScriptConnect && "addListener" in WEBEXT?.runtime?.onUserScriptConnect) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnUserScriptConnect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onUserScriptConnect.addListener);
    },
    "call_OnUserScriptConnect": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onUserScriptConnect.addListener(A.H.get<object>(callback));
    },
    "try_OnUserScriptConnect": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onUserScriptConnect.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffUserScriptConnect": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onUserScriptConnect && "removeListener" in WEBEXT?.runtime?.onUserScriptConnect) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffUserScriptConnect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onUserScriptConnect.removeListener);
    },
    "call_OffUserScriptConnect": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onUserScriptConnect.removeListener(A.H.get<object>(callback));
    },
    "try_OffUserScriptConnect": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onUserScriptConnect.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnUserScriptConnect": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onUserScriptConnect && "hasListener" in WEBEXT?.runtime?.onUserScriptConnect) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnUserScriptConnect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onUserScriptConnect.hasListener);
    },
    "call_HasOnUserScriptConnect": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onUserScriptConnect.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnUserScriptConnect": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onUserScriptConnect.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnUserScriptMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onUserScriptMessage && "addListener" in WEBEXT?.runtime?.onUserScriptMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnUserScriptMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onUserScriptMessage.addListener);
    },
    "call_OnUserScriptMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onUserScriptMessage.addListener(A.H.get<object>(callback));
    },
    "try_OnUserScriptMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onUserScriptMessage.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffUserScriptMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onUserScriptMessage && "removeListener" in WEBEXT?.runtime?.onUserScriptMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffUserScriptMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onUserScriptMessage.removeListener);
    },
    "call_OffUserScriptMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onUserScriptMessage.removeListener(A.H.get<object>(callback));
    },
    "try_OffUserScriptMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onUserScriptMessage.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnUserScriptMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime?.onUserScriptMessage && "hasListener" in WEBEXT?.runtime?.onUserScriptMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnUserScriptMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.onUserScriptMessage.hasListener);
    },
    "call_HasOnUserScriptMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.onUserScriptMessage.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnUserScriptMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.onUserScriptMessage.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenOptionsPage": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "openOptionsPage" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenOptionsPage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.openOptionsPage);
    },
    "call_OpenOptionsPage": (retPtr: Pointer): void => {
      const _ret = WEBEXT.runtime.openOptionsPage();
      A.store.Ref(retPtr, _ret);
    },
    "try_OpenOptionsPage": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.openOptionsPage();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Reload": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "reload" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Reload": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.reload);
    },
    "call_Reload": (retPtr: Pointer): void => {
      const _ret = WEBEXT.runtime.reload();
    },
    "try_Reload": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.reload();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RequestUpdateCheck": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "requestUpdateCheck" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RequestUpdateCheck": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.requestUpdateCheck);
    },
    "call_RequestUpdateCheck": (retPtr: Pointer): void => {
      const _ret = WEBEXT.runtime.requestUpdateCheck();
      A.store.Ref(retPtr, _ret);
    },
    "try_RequestUpdateCheck": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.requestUpdateCheck();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Restart": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "restart" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Restart": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.restart);
    },
    "call_Restart": (retPtr: Pointer): void => {
      const _ret = WEBEXT.runtime.restart();
    },
    "try_Restart": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.restart();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RestartAfterDelay": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "restartAfterDelay" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RestartAfterDelay": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.restartAfterDelay);
    },
    "call_RestartAfterDelay": (retPtr: Pointer, seconds: number): void => {
      const _ret = WEBEXT.runtime.restartAfterDelay(seconds);
      A.store.Ref(retPtr, _ret);
    },
    "try_RestartAfterDelay": (retPtr: Pointer, errPtr: Pointer, seconds: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.restartAfterDelay(seconds);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "sendMessage" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.sendMessage);
    },
    "call_SendMessage": (
      retPtr: Pointer,
      extensionId: heap.Ref<object>,
      message: heap.Ref<object>,
      options: Pointer
    ): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 1)) {
        options_ffi["includeTlsChannelId"] = A.load.Bool(options + 0);
      }

      const _ret = WEBEXT.runtime.sendMessage(A.H.get<object>(extensionId), A.H.get<object>(message), options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SendMessage": (
      retPtr: Pointer,
      errPtr: Pointer,
      extensionId: heap.Ref<object>,
      message: heap.Ref<object>,
      options: Pointer
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 1)) {
          options_ffi["includeTlsChannelId"] = A.load.Bool(options + 0);
        }

        const _ret = WEBEXT.runtime.sendMessage(A.H.get<object>(extensionId), A.H.get<object>(message), options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendNativeMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "sendNativeMessage" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendNativeMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.sendNativeMessage);
    },
    "call_SendNativeMessage": (retPtr: Pointer, application: heap.Ref<object>, message: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.sendNativeMessage(A.H.get<object>(application), A.H.get<object>(message));
      A.store.Ref(retPtr, _ret);
    },
    "try_SendNativeMessage": (
      retPtr: Pointer,
      errPtr: Pointer,
      application: heap.Ref<object>,
      message: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.sendNativeMessage(A.H.get<object>(application), A.H.get<object>(message));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetUninstallURL": (): heap.Ref<boolean> => {
      if (WEBEXT?.runtime && "setUninstallURL" in WEBEXT?.runtime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetUninstallURL": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.runtime.setUninstallURL);
    },
    "call_SetUninstallURL": (retPtr: Pointer, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.runtime.setUninstallURL(A.H.get<object>(url));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetUninstallURL": (retPtr: Pointer, errPtr: Pointer, url: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.runtime.setUninstallURL(A.H.get<object>(url));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
