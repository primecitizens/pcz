import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/contextmenus", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "get_ACTION_MENU_TOP_LEVEL_LIMIT": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.contextMenus && "ACTION_MENU_TOP_LEVEL_LIMIT" in WEBEXT?.contextMenus) {
        const val = WEBEXT.contextMenus.ACTION_MENU_TOP_LEVEL_LIMIT;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_ACTION_MENU_TOP_LEVEL_LIMIT": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.contextMenus, "ACTION_MENU_TOP_LEVEL_LIMIT", A.H.get<object>(val), WEBEXT.contextMenus)
        ? A.H.TRUE
        : A.H.FALSE;
    },

    "store_OnClickData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 52, false);
        A.store.Bool(ptr + 49, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 50, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Bool(ptr + 51, false);
        A.store.Bool(ptr + 48, false);
      } else {
        A.store.Bool(ptr + 52, true);
        A.store.Bool(ptr + 49, "checked" in x ? true : false);
        A.store.Bool(ptr + 0, x["checked"] ? true : false);
        A.store.Bool(ptr + 1, x["editable"] ? true : false);
        A.store.Bool(ptr + 50, "frameId" in x ? true : false);
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Ref(ptr + 16, x["frameUrl"]);
        A.store.Ref(ptr + 20, x["linkUrl"]);
        A.store.Ref(ptr + 24, x["mediaType"]);
        A.store.Ref(ptr + 28, x["menuItemId"]);
        A.store.Ref(ptr + 32, x["pageUrl"]);
        A.store.Ref(ptr + 36, x["parentMenuItemId"]);
        A.store.Ref(ptr + 40, x["selectionText"]);
        A.store.Ref(ptr + 44, x["srcUrl"]);
        A.store.Bool(ptr + 51, "wasChecked" in x ? true : false);
        A.store.Bool(ptr + 48, x["wasChecked"] ? true : false);
      }
    },
    "load_OnClickData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 49)) {
        x["checked"] = A.load.Bool(ptr + 0);
      } else {
        delete x["checked"];
      }
      x["editable"] = A.load.Bool(ptr + 1);
      if (A.load.Bool(ptr + 50)) {
        x["frameId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["frameId"];
      }
      x["frameUrl"] = A.load.Ref(ptr + 16, undefined);
      x["linkUrl"] = A.load.Ref(ptr + 20, undefined);
      x["mediaType"] = A.load.Ref(ptr + 24, undefined);
      x["menuItemId"] = A.load.Ref(ptr + 28, undefined);
      x["pageUrl"] = A.load.Ref(ptr + 32, undefined);
      x["parentMenuItemId"] = A.load.Ref(ptr + 36, undefined);
      x["selectionText"] = A.load.Ref(ptr + 40, undefined);
      x["srcUrl"] = A.load.Ref(ptr + 44, undefined);
      if (A.load.Bool(ptr + 51)) {
        x["wasChecked"] = A.load.Bool(ptr + 48);
      } else {
        delete x["wasChecked"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ItemType": (ref: heap.Ref<string>): number => {
      const idx = ["normal", "checkbox", "radio", "separator"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ContextType": (ref: heap.Ref<string>): number => {
      const idx = [
        "all",
        "page",
        "frame",
        "selection",
        "link",
        "editable",
        "image",
        "video",
        "audio",
        "launcher",
        "browser_action",
        "page_action",
        "action",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_CreateArgCreateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 44, false);
        A.store.Bool(ptr + 41, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 42, false);
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Enum(ptr + 36, -1);
        A.store.Bool(ptr + 43, false);
        A.store.Bool(ptr + 40, false);
      } else {
        A.store.Bool(ptr + 44, true);
        A.store.Bool(ptr + 41, "checked" in x ? true : false);
        A.store.Bool(ptr + 0, x["checked"] ? true : false);
        A.store.Ref(ptr + 4, x["contexts"]);
        A.store.Ref(ptr + 8, x["documentUrlPatterns"]);
        A.store.Bool(ptr + 42, "enabled" in x ? true : false);
        A.store.Bool(ptr + 12, x["enabled"] ? true : false);
        A.store.Ref(ptr + 16, x["id"]);
        A.store.Ref(ptr + 20, x["onclick"]);
        A.store.Ref(ptr + 24, x["parentId"]);
        A.store.Ref(ptr + 28, x["targetUrlPatterns"]);
        A.store.Ref(ptr + 32, x["title"]);
        A.store.Enum(ptr + 36, ["normal", "checkbox", "radio", "separator"].indexOf(x["type"] as string));
        A.store.Bool(ptr + 43, "visible" in x ? true : false);
        A.store.Bool(ptr + 40, x["visible"] ? true : false);
      }
    },
    "load_CreateArgCreateProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 41)) {
        x["checked"] = A.load.Bool(ptr + 0);
      } else {
        delete x["checked"];
      }
      x["contexts"] = A.load.Ref(ptr + 4, undefined);
      x["documentUrlPatterns"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 42)) {
        x["enabled"] = A.load.Bool(ptr + 12);
      } else {
        delete x["enabled"];
      }
      x["id"] = A.load.Ref(ptr + 16, undefined);
      x["onclick"] = A.load.Ref(ptr + 20, undefined);
      x["parentId"] = A.load.Ref(ptr + 24, undefined);
      x["targetUrlPatterns"] = A.load.Ref(ptr + 28, undefined);
      x["title"] = A.load.Ref(ptr + 32, undefined);
      x["type"] = A.load.Enum(ptr + 36, ["normal", "checkbox", "radio", "separator"]);
      if (A.load.Bool(ptr + 43)) {
        x["visible"] = A.load.Bool(ptr + 40);
      } else {
        delete x["visible"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UpdateArgUpdateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 40, false);
        A.store.Bool(ptr + 37, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 38, false);
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Enum(ptr + 32, -1);
        A.store.Bool(ptr + 39, false);
        A.store.Bool(ptr + 36, false);
      } else {
        A.store.Bool(ptr + 40, true);
        A.store.Bool(ptr + 37, "checked" in x ? true : false);
        A.store.Bool(ptr + 0, x["checked"] ? true : false);
        A.store.Ref(ptr + 4, x["contexts"]);
        A.store.Ref(ptr + 8, x["documentUrlPatterns"]);
        A.store.Bool(ptr + 38, "enabled" in x ? true : false);
        A.store.Bool(ptr + 12, x["enabled"] ? true : false);
        A.store.Ref(ptr + 16, x["onclick"]);
        A.store.Ref(ptr + 20, x["parentId"]);
        A.store.Ref(ptr + 24, x["targetUrlPatterns"]);
        A.store.Ref(ptr + 28, x["title"]);
        A.store.Enum(ptr + 32, ["normal", "checkbox", "radio", "separator"].indexOf(x["type"] as string));
        A.store.Bool(ptr + 39, "visible" in x ? true : false);
        A.store.Bool(ptr + 36, x["visible"] ? true : false);
      }
    },
    "load_UpdateArgUpdateProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 37)) {
        x["checked"] = A.load.Bool(ptr + 0);
      } else {
        delete x["checked"];
      }
      x["contexts"] = A.load.Ref(ptr + 4, undefined);
      x["documentUrlPatterns"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 38)) {
        x["enabled"] = A.load.Bool(ptr + 12);
      } else {
        delete x["enabled"];
      }
      x["onclick"] = A.load.Ref(ptr + 16, undefined);
      x["parentId"] = A.load.Ref(ptr + 20, undefined);
      x["targetUrlPatterns"] = A.load.Ref(ptr + 24, undefined);
      x["title"] = A.load.Ref(ptr + 28, undefined);
      x["type"] = A.load.Enum(ptr + 32, ["normal", "checkbox", "radio", "separator"]);
      if (A.load.Bool(ptr + 39)) {
        x["visible"] = A.load.Bool(ptr + 36);
      } else {
        delete x["visible"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Create": (): heap.Ref<boolean> => {
      if (WEBEXT?.contextMenus && "create" in WEBEXT?.contextMenus) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Create": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.contextMenus.create);
    },
    "call_Create": (retPtr: Pointer, createProperties: Pointer, callback: heap.Ref<object>): void => {
      const createProperties_ffi = {};

      if (A.load.Bool(createProperties + 41)) {
        createProperties_ffi["checked"] = A.load.Bool(createProperties + 0);
      }
      createProperties_ffi["contexts"] = A.load.Ref(createProperties + 4, undefined);
      createProperties_ffi["documentUrlPatterns"] = A.load.Ref(createProperties + 8, undefined);
      if (A.load.Bool(createProperties + 42)) {
        createProperties_ffi["enabled"] = A.load.Bool(createProperties + 12);
      }
      createProperties_ffi["id"] = A.load.Ref(createProperties + 16, undefined);
      createProperties_ffi["onclick"] = A.load.Ref(createProperties + 20, undefined);
      createProperties_ffi["parentId"] = A.load.Ref(createProperties + 24, undefined);
      createProperties_ffi["targetUrlPatterns"] = A.load.Ref(createProperties + 28, undefined);
      createProperties_ffi["title"] = A.load.Ref(createProperties + 32, undefined);
      createProperties_ffi["type"] = A.load.Enum(createProperties + 36, ["normal", "checkbox", "radio", "separator"]);
      if (A.load.Bool(createProperties + 43)) {
        createProperties_ffi["visible"] = A.load.Bool(createProperties + 40);
      }

      const _ret = WEBEXT.contextMenus.create(createProperties_ffi, A.H.get<object>(callback));
      A.store.Ref(retPtr, _ret);
    },
    "try_Create": (
      retPtr: Pointer,
      errPtr: Pointer,
      createProperties: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const createProperties_ffi = {};

        if (A.load.Bool(createProperties + 41)) {
          createProperties_ffi["checked"] = A.load.Bool(createProperties + 0);
        }
        createProperties_ffi["contexts"] = A.load.Ref(createProperties + 4, undefined);
        createProperties_ffi["documentUrlPatterns"] = A.load.Ref(createProperties + 8, undefined);
        if (A.load.Bool(createProperties + 42)) {
          createProperties_ffi["enabled"] = A.load.Bool(createProperties + 12);
        }
        createProperties_ffi["id"] = A.load.Ref(createProperties + 16, undefined);
        createProperties_ffi["onclick"] = A.load.Ref(createProperties + 20, undefined);
        createProperties_ffi["parentId"] = A.load.Ref(createProperties + 24, undefined);
        createProperties_ffi["targetUrlPatterns"] = A.load.Ref(createProperties + 28, undefined);
        createProperties_ffi["title"] = A.load.Ref(createProperties + 32, undefined);
        createProperties_ffi["type"] = A.load.Enum(createProperties + 36, ["normal", "checkbox", "radio", "separator"]);
        if (A.load.Bool(createProperties + 43)) {
          createProperties_ffi["visible"] = A.load.Bool(createProperties + 40);
        }

        const _ret = WEBEXT.contextMenus.create(createProperties_ffi, A.H.get<object>(callback));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.contextMenus?.onClicked && "addListener" in WEBEXT?.contextMenus?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.contextMenus.onClicked.addListener);
    },
    "call_OnClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.contextMenus.onClicked.addListener(A.H.get<object>(callback));
    },
    "try_OnClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.contextMenus.onClicked.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.contextMenus?.onClicked && "removeListener" in WEBEXT?.contextMenus?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.contextMenus.onClicked.removeListener);
    },
    "call_OffClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.contextMenus.onClicked.removeListener(A.H.get<object>(callback));
    },
    "try_OffClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.contextMenus.onClicked.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.contextMenus?.onClicked && "hasListener" in WEBEXT?.contextMenus?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.contextMenus.onClicked.hasListener);
    },
    "call_HasOnClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.contextMenus.onClicked.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.contextMenus.onClicked.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Remove": (): heap.Ref<boolean> => {
      if (WEBEXT?.contextMenus && "remove" in WEBEXT?.contextMenus) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Remove": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.contextMenus.remove);
    },
    "call_Remove": (retPtr: Pointer, menuItemId: heap.Ref<any>, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.contextMenus.remove(A.H.get<any>(menuItemId), A.H.get<object>(callback));
    },
    "try_Remove": (
      retPtr: Pointer,
      errPtr: Pointer,
      menuItemId: heap.Ref<any>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.contextMenus.remove(A.H.get<any>(menuItemId), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveAll": (): heap.Ref<boolean> => {
      if (WEBEXT?.contextMenus && "removeAll" in WEBEXT?.contextMenus) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveAll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.contextMenus.removeAll);
    },
    "call_RemoveAll": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.contextMenus.removeAll(A.H.get<object>(callback));
    },
    "try_RemoveAll": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.contextMenus.removeAll(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Update": (): heap.Ref<boolean> => {
      if (WEBEXT?.contextMenus && "update" in WEBEXT?.contextMenus) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Update": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.contextMenus.update);
    },
    "call_Update": (
      retPtr: Pointer,
      id: heap.Ref<any>,
      updateProperties: Pointer,
      callback: heap.Ref<object>
    ): void => {
      const updateProperties_ffi = {};

      if (A.load.Bool(updateProperties + 37)) {
        updateProperties_ffi["checked"] = A.load.Bool(updateProperties + 0);
      }
      updateProperties_ffi["contexts"] = A.load.Ref(updateProperties + 4, undefined);
      updateProperties_ffi["documentUrlPatterns"] = A.load.Ref(updateProperties + 8, undefined);
      if (A.load.Bool(updateProperties + 38)) {
        updateProperties_ffi["enabled"] = A.load.Bool(updateProperties + 12);
      }
      updateProperties_ffi["onclick"] = A.load.Ref(updateProperties + 16, undefined);
      updateProperties_ffi["parentId"] = A.load.Ref(updateProperties + 20, undefined);
      updateProperties_ffi["targetUrlPatterns"] = A.load.Ref(updateProperties + 24, undefined);
      updateProperties_ffi["title"] = A.load.Ref(updateProperties + 28, undefined);
      updateProperties_ffi["type"] = A.load.Enum(updateProperties + 32, ["normal", "checkbox", "radio", "separator"]);
      if (A.load.Bool(updateProperties + 39)) {
        updateProperties_ffi["visible"] = A.load.Bool(updateProperties + 36);
      }

      const _ret = WEBEXT.contextMenus.update(A.H.get<any>(id), updateProperties_ffi, A.H.get<object>(callback));
    },
    "try_Update": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<any>,
      updateProperties: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const updateProperties_ffi = {};

        if (A.load.Bool(updateProperties + 37)) {
          updateProperties_ffi["checked"] = A.load.Bool(updateProperties + 0);
        }
        updateProperties_ffi["contexts"] = A.load.Ref(updateProperties + 4, undefined);
        updateProperties_ffi["documentUrlPatterns"] = A.load.Ref(updateProperties + 8, undefined);
        if (A.load.Bool(updateProperties + 38)) {
          updateProperties_ffi["enabled"] = A.load.Bool(updateProperties + 12);
        }
        updateProperties_ffi["onclick"] = A.load.Ref(updateProperties + 16, undefined);
        updateProperties_ffi["parentId"] = A.load.Ref(updateProperties + 20, undefined);
        updateProperties_ffi["targetUrlPatterns"] = A.load.Ref(updateProperties + 24, undefined);
        updateProperties_ffi["title"] = A.load.Ref(updateProperties + 28, undefined);
        updateProperties_ffi["type"] = A.load.Enum(updateProperties + 32, ["normal", "checkbox", "radio", "separator"]);
        if (A.load.Bool(updateProperties + 39)) {
          updateProperties_ffi["visible"] = A.load.Bool(updateProperties + 36);
        }

        const _ret = WEBEXT.contextMenus.update(A.H.get<any>(id), updateProperties_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
