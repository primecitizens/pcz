import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/chromewebviewinternal", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_ContextMenuItem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Int64(ptr + 0, x["commandId"] === undefined ? 0 : (x["commandId"] as number));
        A.store.Ref(ptr + 8, x["label"]);
      }
    },
    "load_ContextMenuItem": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["commandId"] = A.load.Int64(ptr + 0);
      x["label"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ContextMenusCreateArgCreateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_ContextMenusCreateArgCreateProperties": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
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

    "store_ContextMenusUpdateArgUpdateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_ContextMenusUpdateArgUpdateProperties": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
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

    "store_OnShowArgEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["preventDefault"]);
      }
    },
    "load_OnShowArgEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["preventDefault"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_ContextMenusCreate": (): heap.Ref<boolean> => {
      if (WEBEXT?.chromeWebViewInternal && "contextMenusCreate" in WEBEXT?.chromeWebViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContextMenusCreate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.chromeWebViewInternal.contextMenusCreate);
    },
    "call_ContextMenusCreate": (
      retPtr: Pointer,
      instanceId: number,
      createProperties: Pointer,
      callback: heap.Ref<object>
    ): void => {
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

      const _ret = WEBEXT.chromeWebViewInternal.contextMenusCreate(
        instanceId,
        createProperties_ffi,
        A.H.get<object>(callback)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_ContextMenusCreate": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
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

        const _ret = WEBEXT.chromeWebViewInternal.contextMenusCreate(
          instanceId,
          createProperties_ffi,
          A.H.get<object>(callback)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ContextMenusRemove": (): heap.Ref<boolean> => {
      if (WEBEXT?.chromeWebViewInternal && "contextMenusRemove" in WEBEXT?.chromeWebViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContextMenusRemove": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.chromeWebViewInternal.contextMenusRemove);
    },
    "call_ContextMenusRemove": (
      retPtr: Pointer,
      instanceId: number,
      menuItemId: heap.Ref<any>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.chromeWebViewInternal.contextMenusRemove(
        instanceId,
        A.H.get<any>(menuItemId),
        A.H.get<object>(callback)
      );
    },
    "try_ContextMenusRemove": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      menuItemId: heap.Ref<any>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.chromeWebViewInternal.contextMenusRemove(
          instanceId,
          A.H.get<any>(menuItemId),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ContextMenusRemoveAll": (): heap.Ref<boolean> => {
      if (WEBEXT?.chromeWebViewInternal && "contextMenusRemoveAll" in WEBEXT?.chromeWebViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContextMenusRemoveAll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.chromeWebViewInternal.contextMenusRemoveAll);
    },
    "call_ContextMenusRemoveAll": (retPtr: Pointer, instanceId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.chromeWebViewInternal.contextMenusRemoveAll(instanceId, A.H.get<object>(callback));
    },
    "try_ContextMenusRemoveAll": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.chromeWebViewInternal.contextMenusRemoveAll(instanceId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ContextMenusUpdate": (): heap.Ref<boolean> => {
      if (WEBEXT?.chromeWebViewInternal && "contextMenusUpdate" in WEBEXT?.chromeWebViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContextMenusUpdate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.chromeWebViewInternal.contextMenusUpdate);
    },
    "call_ContextMenusUpdate": (
      retPtr: Pointer,
      instanceId: number,
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

      const _ret = WEBEXT.chromeWebViewInternal.contextMenusUpdate(
        instanceId,
        A.H.get<any>(id),
        updateProperties_ffi,
        A.H.get<object>(callback)
      );
    },
    "try_ContextMenusUpdate": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
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

        const _ret = WEBEXT.chromeWebViewInternal.contextMenusUpdate(
          instanceId,
          A.H.get<any>(id),
          updateProperties_ffi,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.chromeWebViewInternal?.onClicked && "addListener" in WEBEXT?.chromeWebViewInternal?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.chromeWebViewInternal.onClicked.addListener);
    },
    "call_OnClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.chromeWebViewInternal.onClicked.addListener(A.H.get<object>(callback));
    },
    "try_OnClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.chromeWebViewInternal.onClicked.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.chromeWebViewInternal?.onClicked && "removeListener" in WEBEXT?.chromeWebViewInternal?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.chromeWebViewInternal.onClicked.removeListener);
    },
    "call_OffClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.chromeWebViewInternal.onClicked.removeListener(A.H.get<object>(callback));
    },
    "try_OffClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.chromeWebViewInternal.onClicked.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.chromeWebViewInternal?.onClicked && "hasListener" in WEBEXT?.chromeWebViewInternal?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.chromeWebViewInternal.onClicked.hasListener);
    },
    "call_HasOnClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.chromeWebViewInternal.onClicked.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.chromeWebViewInternal.onClicked.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnShow": (): heap.Ref<boolean> => {
      if (WEBEXT?.chromeWebViewInternal?.onShow && "addListener" in WEBEXT?.chromeWebViewInternal?.onShow) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnShow": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.chromeWebViewInternal.onShow.addListener);
    },
    "call_OnShow": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.chromeWebViewInternal.onShow.addListener(A.H.get<object>(callback));
    },
    "try_OnShow": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.chromeWebViewInternal.onShow.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffShow": (): heap.Ref<boolean> => {
      if (WEBEXT?.chromeWebViewInternal?.onShow && "removeListener" in WEBEXT?.chromeWebViewInternal?.onShow) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffShow": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.chromeWebViewInternal.onShow.removeListener);
    },
    "call_OffShow": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.chromeWebViewInternal.onShow.removeListener(A.H.get<object>(callback));
    },
    "try_OffShow": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.chromeWebViewInternal.onShow.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnShow": (): heap.Ref<boolean> => {
      if (WEBEXT?.chromeWebViewInternal?.onShow && "hasListener" in WEBEXT?.chromeWebViewInternal?.onShow) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnShow": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.chromeWebViewInternal.onShow.hasListener);
    },
    "call_HasOnShow": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.chromeWebViewInternal.onShow.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnShow": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.chromeWebViewInternal.onShow.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ShowContextMenu": (): heap.Ref<boolean> => {
      if (WEBEXT?.chromeWebViewInternal && "showContextMenu" in WEBEXT?.chromeWebViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ShowContextMenu": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.chromeWebViewInternal.showContextMenu);
    },
    "call_ShowContextMenu": (
      retPtr: Pointer,
      instanceId: number,
      requestId: number,
      itemsToShow: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.chromeWebViewInternal.showContextMenu(instanceId, requestId, A.H.get<object>(itemsToShow));
    },
    "try_ShowContextMenu": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      requestId: number,
      itemsToShow: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.chromeWebViewInternal.showContextMenu(instanceId, requestId, A.H.get<object>(itemsToShow));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
