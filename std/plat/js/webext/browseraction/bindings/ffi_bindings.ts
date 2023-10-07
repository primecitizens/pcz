import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/browseraction", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_ImageDataType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_ImageDataType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

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
    "has_Disable": (): heap.Ref<boolean> => {
      if (WEBEXT?.browserAction && "disable" in WEBEXT?.browserAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Disable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browserAction.disable);
    },
    "call_Disable": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.browserAction.disable(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Disable": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.browserAction.disable(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Enable": (): heap.Ref<boolean> => {
      if (WEBEXT?.browserAction && "enable" in WEBEXT?.browserAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Enable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browserAction.enable);
    },
    "call_Enable": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.browserAction.enable(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Enable": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.browserAction.enable(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetBadgeBackgroundColor": (): heap.Ref<boolean> => {
      if (WEBEXT?.browserAction && "getBadgeBackgroundColor" in WEBEXT?.browserAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetBadgeBackgroundColor": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browserAction.getBadgeBackgroundColor);
    },
    "call_GetBadgeBackgroundColor": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 8)) {
        details_ffi["tabId"] = A.load.Int64(details + 0);
      }

      const _ret = WEBEXT.browserAction.getBadgeBackgroundColor(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetBadgeBackgroundColor": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 8)) {
          details_ffi["tabId"] = A.load.Int64(details + 0);
        }

        const _ret = WEBEXT.browserAction.getBadgeBackgroundColor(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetBadgeText": (): heap.Ref<boolean> => {
      if (WEBEXT?.browserAction && "getBadgeText" in WEBEXT?.browserAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetBadgeText": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browserAction.getBadgeText);
    },
    "call_GetBadgeText": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 8)) {
        details_ffi["tabId"] = A.load.Int64(details + 0);
      }

      const _ret = WEBEXT.browserAction.getBadgeText(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetBadgeText": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 8)) {
          details_ffi["tabId"] = A.load.Int64(details + 0);
        }

        const _ret = WEBEXT.browserAction.getBadgeText(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPopup": (): heap.Ref<boolean> => {
      if (WEBEXT?.browserAction && "getPopup" in WEBEXT?.browserAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPopup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browserAction.getPopup);
    },
    "call_GetPopup": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 8)) {
        details_ffi["tabId"] = A.load.Int64(details + 0);
      }

      const _ret = WEBEXT.browserAction.getPopup(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPopup": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 8)) {
          details_ffi["tabId"] = A.load.Int64(details + 0);
        }

        const _ret = WEBEXT.browserAction.getPopup(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetTitle": (): heap.Ref<boolean> => {
      if (WEBEXT?.browserAction && "getTitle" in WEBEXT?.browserAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetTitle": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browserAction.getTitle);
    },
    "call_GetTitle": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 8)) {
        details_ffi["tabId"] = A.load.Int64(details + 0);
      }

      const _ret = WEBEXT.browserAction.getTitle(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetTitle": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 8)) {
          details_ffi["tabId"] = A.load.Int64(details + 0);
        }

        const _ret = WEBEXT.browserAction.getTitle(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.browserAction?.onClicked && "addListener" in WEBEXT?.browserAction?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browserAction.onClicked.addListener);
    },
    "call_OnClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.browserAction.onClicked.addListener(A.H.get<object>(callback));
    },
    "try_OnClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.browserAction.onClicked.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.browserAction?.onClicked && "removeListener" in WEBEXT?.browserAction?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browserAction.onClicked.removeListener);
    },
    "call_OffClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.browserAction.onClicked.removeListener(A.H.get<object>(callback));
    },
    "try_OffClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.browserAction.onClicked.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.browserAction?.onClicked && "hasListener" in WEBEXT?.browserAction?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browserAction.onClicked.hasListener);
    },
    "call_HasOnClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.browserAction.onClicked.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.browserAction.onClicked.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenPopup": (): heap.Ref<boolean> => {
      if (WEBEXT?.browserAction && "openPopup" in WEBEXT?.browserAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenPopup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browserAction.openPopup);
    },
    "call_OpenPopup": (retPtr: Pointer): void => {
      const _ret = WEBEXT.browserAction.openPopup();
      A.store.Ref(retPtr, _ret);
    },
    "try_OpenPopup": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.browserAction.openPopup();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetBadgeBackgroundColor": (): heap.Ref<boolean> => {
      if (WEBEXT?.browserAction && "setBadgeBackgroundColor" in WEBEXT?.browserAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetBadgeBackgroundColor": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browserAction.setBadgeBackgroundColor);
    },
    "call_SetBadgeBackgroundColor": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["color"] = A.load.Ref(details + 0, undefined);
      if (A.load.Bool(details + 16)) {
        details_ffi["tabId"] = A.load.Int64(details + 8);
      }

      const _ret = WEBEXT.browserAction.setBadgeBackgroundColor(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetBadgeBackgroundColor": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["color"] = A.load.Ref(details + 0, undefined);
        if (A.load.Bool(details + 16)) {
          details_ffi["tabId"] = A.load.Int64(details + 8);
        }

        const _ret = WEBEXT.browserAction.setBadgeBackgroundColor(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetBadgeText": (): heap.Ref<boolean> => {
      if (WEBEXT?.browserAction && "setBadgeText" in WEBEXT?.browserAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetBadgeText": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browserAction.setBadgeText);
    },
    "call_SetBadgeText": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 12)) {
        details_ffi["tabId"] = A.load.Int64(details + 0);
      }
      details_ffi["text"] = A.load.Ref(details + 8, undefined);

      const _ret = WEBEXT.browserAction.setBadgeText(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetBadgeText": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 12)) {
          details_ffi["tabId"] = A.load.Int64(details + 0);
        }
        details_ffi["text"] = A.load.Ref(details + 8, undefined);

        const _ret = WEBEXT.browserAction.setBadgeText(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetIcon": (): heap.Ref<boolean> => {
      if (WEBEXT?.browserAction && "setIcon" in WEBEXT?.browserAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetIcon": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browserAction.setIcon);
    },
    "call_SetIcon": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["imageData"] = A.load.Ref(details + 0, undefined);
      details_ffi["path"] = A.load.Ref(details + 4, undefined);
      if (A.load.Bool(details + 16)) {
        details_ffi["tabId"] = A.load.Int64(details + 8);
      }

      const _ret = WEBEXT.browserAction.setIcon(details_ffi);
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

        const _ret = WEBEXT.browserAction.setIcon(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPopup": (): heap.Ref<boolean> => {
      if (WEBEXT?.browserAction && "setPopup" in WEBEXT?.browserAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPopup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browserAction.setPopup);
    },
    "call_SetPopup": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["popup"] = A.load.Ref(details + 0, undefined);
      if (A.load.Bool(details + 16)) {
        details_ffi["tabId"] = A.load.Int64(details + 8);
      }

      const _ret = WEBEXT.browserAction.setPopup(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetPopup": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["popup"] = A.load.Ref(details + 0, undefined);
        if (A.load.Bool(details + 16)) {
          details_ffi["tabId"] = A.load.Int64(details + 8);
        }

        const _ret = WEBEXT.browserAction.setPopup(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetTitle": (): heap.Ref<boolean> => {
      if (WEBEXT?.browserAction && "setTitle" in WEBEXT?.browserAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetTitle": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browserAction.setTitle);
    },
    "call_SetTitle": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 12)) {
        details_ffi["tabId"] = A.load.Int64(details + 0);
      }
      details_ffi["title"] = A.load.Ref(details + 8, undefined);

      const _ret = WEBEXT.browserAction.setTitle(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetTitle": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 12)) {
          details_ffi["tabId"] = A.load.Int64(details + 0);
        }
        details_ffi["title"] = A.load.Ref(details + 8, undefined);

        const _ret = WEBEXT.browserAction.setTitle(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
