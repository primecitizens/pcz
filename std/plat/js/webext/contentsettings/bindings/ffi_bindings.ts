import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/contentsettings", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_AutoVerifyContentSetting": (ref: heap.Ref<string>): number => {
      const idx = ["allow", "block"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_CameraContentSetting": (ref: heap.Ref<string>): number => {
      const idx = ["allow", "block", "ask"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_Scope": (ref: heap.Ref<string>): number => {
      const idx = ["regular", "incognito_session_only"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ClearArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["regular", "incognito_session_only"].indexOf(x["scope"] as string));
      }
    },
    "load_ClearArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["scope"] = A.load.Enum(ptr + 0, ["regular", "incognito_session_only"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["setting"]);
      }
    },
    "load_GetReturnType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["setting"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ResourceIdentifier": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["description"]);
        A.store.Ref(ptr + 4, x["id"]);
      }
    },
    "load_ResourceIdentifier": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["description"] = A.load.Ref(ptr + 0, undefined);
      x["id"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);

        A.store.Bool(ptr + 8 + 8, false);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Ref(ptr + 8 + 4, undefined);
        A.store.Ref(ptr + 20, undefined);
      } else {
        A.store.Bool(ptr + 25, true);
        A.store.Bool(ptr + 24, "incognito" in x ? true : false);
        A.store.Bool(ptr + 0, x["incognito"] ? true : false);
        A.store.Ref(ptr + 4, x["primaryUrl"]);

        if (typeof x["resourceIdentifier"] === "undefined") {
          A.store.Bool(ptr + 8 + 8, false);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Ref(ptr + 8 + 4, undefined);
        } else {
          A.store.Bool(ptr + 8 + 8, true);
          A.store.Ref(ptr + 8 + 0, x["resourceIdentifier"]["description"]);
          A.store.Ref(ptr + 8 + 4, x["resourceIdentifier"]["id"]);
        }
        A.store.Ref(ptr + 20, x["secondaryUrl"]);
      }
    },
    "load_GetArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["incognito"] = A.load.Bool(ptr + 0);
      } else {
        delete x["incognito"];
      }
      x["primaryUrl"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 8 + 8)) {
        x["resourceIdentifier"] = {};
        x["resourceIdentifier"]["description"] = A.load.Ref(ptr + 8 + 0, undefined);
        x["resourceIdentifier"]["id"] = A.load.Ref(ptr + 8 + 4, undefined);
      } else {
        delete x["resourceIdentifier"];
      }
      x["secondaryUrl"] = A.load.Ref(ptr + 20, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 4 + 8, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Ref(ptr + 4 + 4, undefined);
        A.store.Enum(ptr + 16, -1);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 28, true);
        A.store.Ref(ptr + 0, x["primaryPattern"]);

        if (typeof x["resourceIdentifier"] === "undefined") {
          A.store.Bool(ptr + 4 + 8, false);
          A.store.Ref(ptr + 4 + 0, undefined);
          A.store.Ref(ptr + 4 + 4, undefined);
        } else {
          A.store.Bool(ptr + 4 + 8, true);
          A.store.Ref(ptr + 4 + 0, x["resourceIdentifier"]["description"]);
          A.store.Ref(ptr + 4 + 4, x["resourceIdentifier"]["id"]);
        }
        A.store.Enum(ptr + 16, ["regular", "incognito_session_only"].indexOf(x["scope"] as string));
        A.store.Ref(ptr + 20, x["secondaryPattern"]);
        A.store.Ref(ptr + 24, x["setting"]);
      }
    },
    "load_SetArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["primaryPattern"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 4 + 8)) {
        x["resourceIdentifier"] = {};
        x["resourceIdentifier"]["description"] = A.load.Ref(ptr + 4 + 0, undefined);
        x["resourceIdentifier"]["id"] = A.load.Ref(ptr + 4 + 4, undefined);
      } else {
        delete x["resourceIdentifier"];
      }
      x["scope"] = A.load.Enum(ptr + 16, ["regular", "incognito_session_only"]);
      x["secondaryPattern"] = A.load.Ref(ptr + 20, undefined);
      x["setting"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "has_ContentSetting_Clear": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "clear" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContentSetting_Clear": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).clear);
    },

    "call_ContentSetting_Clear": (self: heap.Ref<object>, retPtr: Pointer, details: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const details_ffi = {};

      details_ffi["scope"] = A.load.Enum(details + 0, ["regular", "incognito_session_only"]);
      const _ret = thiz.clear(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ContentSetting_Clear": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      details: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const details_ffi = {};

        details_ffi["scope"] = A.load.Enum(details + 0, ["regular", "incognito_session_only"]);
        const _ret = thiz.clear(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ContentSetting_Get": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "get" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContentSetting_Get": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).get);
    },

    "call_ContentSetting_Get": (self: heap.Ref<object>, retPtr: Pointer, details: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const details_ffi = {};

      if (A.load.Bool(details + 24)) {
        details_ffi["incognito"] = A.load.Bool(details + 0);
      }
      details_ffi["primaryUrl"] = A.load.Ref(details + 4, undefined);
      if (A.load.Bool(details + 8 + 8)) {
        details_ffi["resourceIdentifier"] = {};
        details_ffi["resourceIdentifier"]["description"] = A.load.Ref(details + 8 + 0, undefined);
        details_ffi["resourceIdentifier"]["id"] = A.load.Ref(details + 8 + 4, undefined);
      }
      details_ffi["secondaryUrl"] = A.load.Ref(details + 20, undefined);
      const _ret = thiz.get(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ContentSetting_Get": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      details: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const details_ffi = {};

        if (A.load.Bool(details + 24)) {
          details_ffi["incognito"] = A.load.Bool(details + 0);
        }
        details_ffi["primaryUrl"] = A.load.Ref(details + 4, undefined);
        if (A.load.Bool(details + 8 + 8)) {
          details_ffi["resourceIdentifier"] = {};
          details_ffi["resourceIdentifier"]["description"] = A.load.Ref(details + 8 + 0, undefined);
          details_ffi["resourceIdentifier"]["id"] = A.load.Ref(details + 8 + 4, undefined);
        }
        details_ffi["secondaryUrl"] = A.load.Ref(details + 20, undefined);
        const _ret = thiz.get(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ContentSetting_Set": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "set" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContentSetting_Set": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).set);
    },

    "call_ContentSetting_Set": (self: heap.Ref<object>, retPtr: Pointer, details: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const details_ffi = {};

      details_ffi["primaryPattern"] = A.load.Ref(details + 0, undefined);
      if (A.load.Bool(details + 4 + 8)) {
        details_ffi["resourceIdentifier"] = {};
        details_ffi["resourceIdentifier"]["description"] = A.load.Ref(details + 4 + 0, undefined);
        details_ffi["resourceIdentifier"]["id"] = A.load.Ref(details + 4 + 4, undefined);
      }
      details_ffi["scope"] = A.load.Enum(details + 16, ["regular", "incognito_session_only"]);
      details_ffi["secondaryPattern"] = A.load.Ref(details + 20, undefined);
      details_ffi["setting"] = A.load.Ref(details + 24, undefined);
      const _ret = thiz.set(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ContentSetting_Set": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      details: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const details_ffi = {};

        details_ffi["primaryPattern"] = A.load.Ref(details + 0, undefined);
        if (A.load.Bool(details + 4 + 8)) {
          details_ffi["resourceIdentifier"] = {};
          details_ffi["resourceIdentifier"]["description"] = A.load.Ref(details + 4 + 0, undefined);
          details_ffi["resourceIdentifier"]["id"] = A.load.Ref(details + 4 + 4, undefined);
        }
        details_ffi["scope"] = A.load.Enum(details + 16, ["regular", "incognito_session_only"]);
        details_ffi["secondaryPattern"] = A.load.Ref(details + 20, undefined);
        details_ffi["setting"] = A.load.Ref(details + 24, undefined);
        const _ret = thiz.set(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ContentSetting_GetResourceIdentifiers": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "getResourceIdentifiers" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContentSetting_GetResourceIdentifiers": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).getResourceIdentifiers);
    },

    "call_ContentSetting_GetResourceIdentifiers": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.getResourceIdentifiers();
      A.store.Ref(retPtr, _ret);
    },
    "try_ContentSetting_GetResourceIdentifiers": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.getResourceIdentifiers();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "constof_CookiesContentSetting": (ref: heap.Ref<string>): number => {
      const idx = ["allow", "block", "session_only"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_FullscreenContentSetting": (ref: heap.Ref<string>): number => {
      const idx = ["allow"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ImagesContentSetting": (ref: heap.Ref<string>): number => {
      const idx = ["allow", "block"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_JavascriptContentSetting": (ref: heap.Ref<string>): number => {
      const idx = ["allow", "block"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_LocationContentSetting": (ref: heap.Ref<string>): number => {
      const idx = ["allow", "block", "ask"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_MicrophoneContentSetting": (ref: heap.Ref<string>): number => {
      const idx = ["allow", "block", "ask"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_MouselockContentSetting": (ref: heap.Ref<string>): number => {
      const idx = ["allow"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_MultipleAutomaticDownloadsContentSetting": (ref: heap.Ref<string>): number => {
      const idx = ["allow", "block", "ask"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_NotificationsContentSetting": (ref: heap.Ref<string>): number => {
      const idx = ["allow", "block", "ask"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PluginsContentSetting": (ref: heap.Ref<string>): number => {
      const idx = ["block"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PopupsContentSetting": (ref: heap.Ref<string>): number => {
      const idx = ["allow", "block"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PpapiBrokerContentSetting": (ref: heap.Ref<string>): number => {
      const idx = ["block"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "get_AutoVerify": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.contentSettings && "autoVerify" in WEBEXT?.contentSettings) {
        const val = WEBEXT.contentSettings.autoVerify;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutoVerify": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.contentSettings, "autoVerify", A.H.get<object>(val), WEBEXT.contentSettings)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_AutomaticDownloads": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.contentSettings && "automaticDownloads" in WEBEXT?.contentSettings) {
        const val = WEBEXT.contentSettings.automaticDownloads;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomaticDownloads": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.contentSettings, "automaticDownloads", A.H.get<object>(val), WEBEXT.contentSettings)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_Camera": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.contentSettings && "camera" in WEBEXT?.contentSettings) {
        const val = WEBEXT.contentSettings.camera;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Camera": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.contentSettings, "camera", A.H.get<object>(val), WEBEXT.contentSettings)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_Cookies": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.contentSettings && "cookies" in WEBEXT?.contentSettings) {
        const val = WEBEXT.contentSettings.cookies;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Cookies": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.contentSettings, "cookies", A.H.get<object>(val), WEBEXT.contentSettings)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_Fullscreen": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.contentSettings && "fullscreen" in WEBEXT?.contentSettings) {
        const val = WEBEXT.contentSettings.fullscreen;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Fullscreen": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.contentSettings, "fullscreen", A.H.get<object>(val), WEBEXT.contentSettings)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_Images": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.contentSettings && "images" in WEBEXT?.contentSettings) {
        const val = WEBEXT.contentSettings.images;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Images": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.contentSettings, "images", A.H.get<object>(val), WEBEXT.contentSettings)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_Javascript": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.contentSettings && "javascript" in WEBEXT?.contentSettings) {
        const val = WEBEXT.contentSettings.javascript;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Javascript": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.contentSettings, "javascript", A.H.get<object>(val), WEBEXT.contentSettings)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_Location": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.contentSettings && "location" in WEBEXT?.contentSettings) {
        const val = WEBEXT.contentSettings.location;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Location": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.contentSettings, "location", A.H.get<object>(val), WEBEXT.contentSettings)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_Microphone": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.contentSettings && "microphone" in WEBEXT?.contentSettings) {
        const val = WEBEXT.contentSettings.microphone;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Microphone": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.contentSettings, "microphone", A.H.get<object>(val), WEBEXT.contentSettings)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_Mouselock": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.contentSettings && "mouselock" in WEBEXT?.contentSettings) {
        const val = WEBEXT.contentSettings.mouselock;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Mouselock": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.contentSettings, "mouselock", A.H.get<object>(val), WEBEXT.contentSettings)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_Notifications": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.contentSettings && "notifications" in WEBEXT?.contentSettings) {
        const val = WEBEXT.contentSettings.notifications;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Notifications": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.contentSettings, "notifications", A.H.get<object>(val), WEBEXT.contentSettings)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_Plugins": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.contentSettings && "plugins" in WEBEXT?.contentSettings) {
        const val = WEBEXT.contentSettings.plugins;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Plugins": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.contentSettings, "plugins", A.H.get<object>(val), WEBEXT.contentSettings)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_Popups": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.contentSettings && "popups" in WEBEXT?.contentSettings) {
        const val = WEBEXT.contentSettings.popups;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Popups": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.contentSettings, "popups", A.H.get<object>(val), WEBEXT.contentSettings)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_UnsandboxedPlugins": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.contentSettings && "unsandboxedPlugins" in WEBEXT?.contentSettings) {
        const val = WEBEXT.contentSettings.unsandboxedPlugins;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_UnsandboxedPlugins": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.contentSettings, "unsandboxedPlugins", A.H.get<object>(val), WEBEXT.contentSettings)
        ? A.H.TRUE
        : A.H.FALSE;
    },
  };
});
