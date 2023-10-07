import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/pageaction", (A: Application) => {
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

    "store_SetIconArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 24, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Int64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 25, true);
        A.store.Bool(ptr + 24, "iconIndex" in x ? true : false);
        A.store.Int64(ptr + 0, x["iconIndex"] === undefined ? 0 : (x["iconIndex"] as number));
        A.store.Ref(ptr + 8, x["imageData"]);
        A.store.Ref(ptr + 12, x["path"]);
        A.store.Int64(ptr + 16, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_SetIconArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["iconIndex"] = A.load.Int64(ptr + 0);
      } else {
        delete x["iconIndex"];
      }
      x["imageData"] = A.load.Ref(ptr + 8, undefined);
      x["path"] = A.load.Ref(ptr + 12, undefined);
      x["tabId"] = A.load.Int64(ptr + 16);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetPopupArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["popup"]);
        A.store.Int64(ptr + 8, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_SetPopupArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["popup"] = A.load.Ref(ptr + 0, undefined);
      x["tabId"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetTitleArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Int64(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Ref(ptr + 8, x["title"]);
      }
    },
    "load_SetTitleArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["tabId"] = A.load.Int64(ptr + 0);
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
    "has_GetPopup": (): heap.Ref<boolean> => {
      if (WEBEXT?.pageAction && "getPopup" in WEBEXT?.pageAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPopup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pageAction.getPopup);
    },
    "call_GetPopup": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 8)) {
        details_ffi["tabId"] = A.load.Int64(details + 0);
      }

      const _ret = WEBEXT.pageAction.getPopup(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPopup": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 8)) {
          details_ffi["tabId"] = A.load.Int64(details + 0);
        }

        const _ret = WEBEXT.pageAction.getPopup(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetTitle": (): heap.Ref<boolean> => {
      if (WEBEXT?.pageAction && "getTitle" in WEBEXT?.pageAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetTitle": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pageAction.getTitle);
    },
    "call_GetTitle": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 8)) {
        details_ffi["tabId"] = A.load.Int64(details + 0);
      }

      const _ret = WEBEXT.pageAction.getTitle(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetTitle": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 8)) {
          details_ffi["tabId"] = A.load.Int64(details + 0);
        }

        const _ret = WEBEXT.pageAction.getTitle(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Hide": (): heap.Ref<boolean> => {
      if (WEBEXT?.pageAction && "hide" in WEBEXT?.pageAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Hide": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pageAction.hide);
    },
    "call_Hide": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.pageAction.hide(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Hide": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.pageAction.hide(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.pageAction?.onClicked && "addListener" in WEBEXT?.pageAction?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pageAction.onClicked.addListener);
    },
    "call_OnClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.pageAction.onClicked.addListener(A.H.get<object>(callback));
    },
    "try_OnClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.pageAction.onClicked.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.pageAction?.onClicked && "removeListener" in WEBEXT?.pageAction?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pageAction.onClicked.removeListener);
    },
    "call_OffClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.pageAction.onClicked.removeListener(A.H.get<object>(callback));
    },
    "try_OffClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.pageAction.onClicked.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.pageAction?.onClicked && "hasListener" in WEBEXT?.pageAction?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pageAction.onClicked.hasListener);
    },
    "call_HasOnClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.pageAction.onClicked.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.pageAction.onClicked.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetIcon": (): heap.Ref<boolean> => {
      if (WEBEXT?.pageAction && "setIcon" in WEBEXT?.pageAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetIcon": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pageAction.setIcon);
    },
    "call_SetIcon": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 24)) {
        details_ffi["iconIndex"] = A.load.Int64(details + 0);
      }
      details_ffi["imageData"] = A.load.Ref(details + 8, undefined);
      details_ffi["path"] = A.load.Ref(details + 12, undefined);
      details_ffi["tabId"] = A.load.Int64(details + 16);

      const _ret = WEBEXT.pageAction.setIcon(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetIcon": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 24)) {
          details_ffi["iconIndex"] = A.load.Int64(details + 0);
        }
        details_ffi["imageData"] = A.load.Ref(details + 8, undefined);
        details_ffi["path"] = A.load.Ref(details + 12, undefined);
        details_ffi["tabId"] = A.load.Int64(details + 16);

        const _ret = WEBEXT.pageAction.setIcon(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPopup": (): heap.Ref<boolean> => {
      if (WEBEXT?.pageAction && "setPopup" in WEBEXT?.pageAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPopup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pageAction.setPopup);
    },
    "call_SetPopup": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["popup"] = A.load.Ref(details + 0, undefined);
      details_ffi["tabId"] = A.load.Int64(details + 8);

      const _ret = WEBEXT.pageAction.setPopup(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetPopup": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["popup"] = A.load.Ref(details + 0, undefined);
        details_ffi["tabId"] = A.load.Int64(details + 8);

        const _ret = WEBEXT.pageAction.setPopup(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetTitle": (): heap.Ref<boolean> => {
      if (WEBEXT?.pageAction && "setTitle" in WEBEXT?.pageAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetTitle": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pageAction.setTitle);
    },
    "call_SetTitle": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["tabId"] = A.load.Int64(details + 0);
      details_ffi["title"] = A.load.Ref(details + 8, undefined);

      const _ret = WEBEXT.pageAction.setTitle(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetTitle": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["tabId"] = A.load.Int64(details + 0);
        details_ffi["title"] = A.load.Ref(details + 8, undefined);

        const _ret = WEBEXT.pageAction.setTitle(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Show": (): heap.Ref<boolean> => {
      if (WEBEXT?.pageAction && "show" in WEBEXT?.pageAction) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Show": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.pageAction.show);
    },
    "call_Show": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.pageAction.show(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Show": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.pageAction.show(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
