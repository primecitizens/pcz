import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/action", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_OpenPopupOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_OpenPopupOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["windowId"] = A.load.Int64(ptr + 0);
      } else {
        delete x["windowId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetBadgeBackgroundColorArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["color"]);
        A.store.Bool(ptr + 16, "tabId" in x ? true : false);
        A.store.Int64(ptr + 8, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_SetBadgeBackgroundColorArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["color"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["tabId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["tabId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetBadgeTextArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "tabId" in x ? true : false);
        A.store.Int64(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Ref(ptr + 8, x["text"]);
      }
    },
    "load_SetBadgeTextArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["tabId"] = A.load.Int64(ptr + 0);
      } else {
        delete x["tabId"];
      }
      x["text"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetBadgeTextColorArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["color"]);
        A.store.Bool(ptr + 16, "tabId" in x ? true : false);
        A.store.Int64(ptr + 8, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_SetBadgeTextColorArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["color"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["tabId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["tabId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetIconArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["imageData"]);
        A.store.Ref(ptr + 4, x["path"]);
        A.store.Bool(ptr + 16, "tabId" in x ? true : false);
        A.store.Int64(ptr + 8, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_SetIconArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["imageData"] = A.load.Ref(ptr + 0, undefined);
      x["path"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["tabId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["tabId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetPopupArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["popup"]);
        A.store.Bool(ptr + 16, "tabId" in x ? true : false);
        A.store.Int64(ptr + 8, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_SetPopupArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["popup"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["tabId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["tabId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetTitleArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "tabId" in x ? true : false);
        A.store.Int64(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Ref(ptr + 8, x["title"]);
      }
    },
    "load_SetTitleArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["tabId"] = A.load.Int64(ptr + 0);
      } else {
        delete x["tabId"];
      }
      x["title"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TabDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "tabId" in x ? true : false);
        A.store.Int64(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_TabDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["tabId"] = A.load.Int64(ptr + 0);
      } else {
        delete x["tabId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UserSettings": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 1, true);
        A.store.Bool(ptr + 0, x["isOnToolbar"] ? true : false);
      }
    },
    "load_UserSettings": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["isOnToolbar"] = A.load.Bool(ptr + 0);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Disable": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "disable" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Disable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.disable);
    },
    "call_Disable": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.action.disable(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Disable": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.action.disable(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Enable": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "enable" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Enable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.enable);
    },
    "call_Enable": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.action.enable(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Enable": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.action.enable(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetBadgeBackgroundColor": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "getBadgeBackgroundColor" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetBadgeBackgroundColor": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.getBadgeBackgroundColor);
    },
    "call_GetBadgeBackgroundColor": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 8)) {
        details_ffi["tabId"] = A.load.Int64(details + 0);
      }

      const _ret = WEBEXT.action.getBadgeBackgroundColor(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetBadgeBackgroundColor": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 8)) {
          details_ffi["tabId"] = A.load.Int64(details + 0);
        }

        const _ret = WEBEXT.action.getBadgeBackgroundColor(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetBadgeText": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "getBadgeText" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetBadgeText": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.getBadgeText);
    },
    "call_GetBadgeText": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 8)) {
        details_ffi["tabId"] = A.load.Int64(details + 0);
      }

      const _ret = WEBEXT.action.getBadgeText(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetBadgeText": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 8)) {
          details_ffi["tabId"] = A.load.Int64(details + 0);
        }

        const _ret = WEBEXT.action.getBadgeText(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetBadgeTextColor": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "getBadgeTextColor" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetBadgeTextColor": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.getBadgeTextColor);
    },
    "call_GetBadgeTextColor": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 8)) {
        details_ffi["tabId"] = A.load.Int64(details + 0);
      }

      const _ret = WEBEXT.action.getBadgeTextColor(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetBadgeTextColor": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 8)) {
          details_ffi["tabId"] = A.load.Int64(details + 0);
        }

        const _ret = WEBEXT.action.getBadgeTextColor(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPopup": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "getPopup" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPopup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.getPopup);
    },
    "call_GetPopup": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 8)) {
        details_ffi["tabId"] = A.load.Int64(details + 0);
      }

      const _ret = WEBEXT.action.getPopup(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPopup": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 8)) {
          details_ffi["tabId"] = A.load.Int64(details + 0);
        }

        const _ret = WEBEXT.action.getPopup(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetTitle": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "getTitle" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetTitle": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.getTitle);
    },
    "call_GetTitle": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 8)) {
        details_ffi["tabId"] = A.load.Int64(details + 0);
      }

      const _ret = WEBEXT.action.getTitle(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetTitle": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 8)) {
          details_ffi["tabId"] = A.load.Int64(details + 0);
        }

        const _ret = WEBEXT.action.getTitle(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetUserSettings": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "getUserSettings" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetUserSettings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.getUserSettings);
    },
    "call_GetUserSettings": (retPtr: Pointer): void => {
      const _ret = WEBEXT.action.getUserSettings();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetUserSettings": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.action.getUserSettings();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "isEnabled" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.isEnabled);
    },
    "call_IsEnabled": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.action.isEnabled(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_IsEnabled": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.action.isEnabled(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.action?.onClicked && "addListener" in WEBEXT?.action?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.onClicked.addListener);
    },
    "call_OnClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.action.onClicked.addListener(A.H.get<object>(callback));
    },
    "try_OnClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.action.onClicked.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.action?.onClicked && "removeListener" in WEBEXT?.action?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.onClicked.removeListener);
    },
    "call_OffClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.action.onClicked.removeListener(A.H.get<object>(callback));
    },
    "try_OffClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.action.onClicked.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.action?.onClicked && "hasListener" in WEBEXT?.action?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.onClicked.hasListener);
    },
    "call_HasOnClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.action.onClicked.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.action.onClicked.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenPopup": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "openPopup" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenPopup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.openPopup);
    },
    "call_OpenPopup": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 8)) {
        options_ffi["windowId"] = A.load.Int64(options + 0);
      }

      const _ret = WEBEXT.action.openPopup(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_OpenPopup": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 8)) {
          options_ffi["windowId"] = A.load.Int64(options + 0);
        }

        const _ret = WEBEXT.action.openPopup(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetBadgeBackgroundColor": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "setBadgeBackgroundColor" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetBadgeBackgroundColor": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.setBadgeBackgroundColor);
    },
    "call_SetBadgeBackgroundColor": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["color"] = A.load.Ref(details + 0, undefined);
      if (A.load.Bool(details + 16)) {
        details_ffi["tabId"] = A.load.Int64(details + 8);
      }

      const _ret = WEBEXT.action.setBadgeBackgroundColor(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetBadgeBackgroundColor": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["color"] = A.load.Ref(details + 0, undefined);
        if (A.load.Bool(details + 16)) {
          details_ffi["tabId"] = A.load.Int64(details + 8);
        }

        const _ret = WEBEXT.action.setBadgeBackgroundColor(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetBadgeText": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "setBadgeText" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetBadgeText": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.setBadgeText);
    },
    "call_SetBadgeText": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 12)) {
        details_ffi["tabId"] = A.load.Int64(details + 0);
      }
      details_ffi["text"] = A.load.Ref(details + 8, undefined);

      const _ret = WEBEXT.action.setBadgeText(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetBadgeText": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 12)) {
          details_ffi["tabId"] = A.load.Int64(details + 0);
        }
        details_ffi["text"] = A.load.Ref(details + 8, undefined);

        const _ret = WEBEXT.action.setBadgeText(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetBadgeTextColor": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "setBadgeTextColor" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetBadgeTextColor": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.setBadgeTextColor);
    },
    "call_SetBadgeTextColor": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["color"] = A.load.Ref(details + 0, undefined);
      if (A.load.Bool(details + 16)) {
        details_ffi["tabId"] = A.load.Int64(details + 8);
      }

      const _ret = WEBEXT.action.setBadgeTextColor(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetBadgeTextColor": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["color"] = A.load.Ref(details + 0, undefined);
        if (A.load.Bool(details + 16)) {
          details_ffi["tabId"] = A.load.Int64(details + 8);
        }

        const _ret = WEBEXT.action.setBadgeTextColor(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetIcon": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "setIcon" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetIcon": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.setIcon);
    },
    "call_SetIcon": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["imageData"] = A.load.Ref(details + 0, undefined);
      details_ffi["path"] = A.load.Ref(details + 4, undefined);
      if (A.load.Bool(details + 16)) {
        details_ffi["tabId"] = A.load.Int64(details + 8);
      }

      const _ret = WEBEXT.action.setIcon(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetIcon": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["imageData"] = A.load.Ref(details + 0, undefined);
        details_ffi["path"] = A.load.Ref(details + 4, undefined);
        if (A.load.Bool(details + 16)) {
          details_ffi["tabId"] = A.load.Int64(details + 8);
        }

        const _ret = WEBEXT.action.setIcon(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPopup": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "setPopup" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPopup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.setPopup);
    },
    "call_SetPopup": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["popup"] = A.load.Ref(details + 0, undefined);
      if (A.load.Bool(details + 16)) {
        details_ffi["tabId"] = A.load.Int64(details + 8);
      }

      const _ret = WEBEXT.action.setPopup(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetPopup": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["popup"] = A.load.Ref(details + 0, undefined);
        if (A.load.Bool(details + 16)) {
          details_ffi["tabId"] = A.load.Int64(details + 8);
        }

        const _ret = WEBEXT.action.setPopup(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetTitle": (): heap.Ref<boolean> => {
      if (WEBEXT?.action && "setTitle" in WEBEXT?.action) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetTitle": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.action.setTitle);
    },
    "call_SetTitle": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 12)) {
        details_ffi["tabId"] = A.load.Int64(details + 0);
      }
      details_ffi["title"] = A.load.Ref(details + 8, undefined);

      const _ret = WEBEXT.action.setTitle(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetTitle": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 12)) {
          details_ffi["tabId"] = A.load.Int64(details + 0);
        }
        details_ffi["title"] = A.load.Ref(details + 8, undefined);

        const _ret = WEBEXT.action.setTitle(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
