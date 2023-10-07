import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/sidepanel", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_GetPanelOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "tabId" in x ? true : false);
        A.store.Int32(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_GetPanelOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["tabId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["tabId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SidePanel": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["default_path"]);
      }
    },
    "load_SidePanel": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["default_path"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManifestKeys": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);

        A.store.Bool(ptr + 0 + 4, false);
        A.store.Ref(ptr + 0 + 0, undefined);
      } else {
        A.store.Bool(ptr + 5, true);

        if (typeof x["side_panel"] === "undefined") {
          A.store.Bool(ptr + 0 + 4, false);
          A.store.Ref(ptr + 0 + 0, undefined);
        } else {
          A.store.Bool(ptr + 0 + 4, true);
          A.store.Ref(ptr + 0 + 0, x["side_panel"]["default_path"]);
        }
      }
    },
    "load_ManifestKeys": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 4)) {
        x["side_panel"] = {};
        x["side_panel"]["default_path"] = A.load.Ref(ptr + 0 + 0, undefined);
      } else {
        delete x["side_panel"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OpenOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "windowId" in x ? true : false);
        A.store.Int32(ptr + 0, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
        A.store.Bool(ptr + 9, "tabId" in x ? true : false);
        A.store.Int32(ptr + 4, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_OpenOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["windowId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["windowId"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["tabId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["tabId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PanelBehavior": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 1, "openPanelOnActionClick" in x ? true : false);
        A.store.Bool(ptr + 0, x["openPanelOnActionClick"] ? true : false);
      }
    },
    "load_PanelBehavior": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 1)) {
        x["openPanelOnActionClick"] = A.load.Bool(ptr + 0);
      } else {
        delete x["openPanelOnActionClick"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PanelOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 11, true);
        A.store.Bool(ptr + 9, "tabId" in x ? true : false);
        A.store.Int32(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Ref(ptr + 4, x["path"]);
        A.store.Bool(ptr + 10, "enabled" in x ? true : false);
        A.store.Bool(ptr + 8, x["enabled"] ? true : false);
      }
    },
    "load_PanelOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 9)) {
        x["tabId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["tabId"];
      }
      x["path"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 10)) {
        x["enabled"] = A.load.Bool(ptr + 8);
      } else {
        delete x["enabled"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetOptions": (): heap.Ref<boolean> => {
      if (WEBEXT?.sidePanel && "getOptions" in WEBEXT?.sidePanel) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetOptions": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sidePanel.getOptions);
    },
    "call_GetOptions": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 4)) {
        options_ffi["tabId"] = A.load.Int32(options + 0);
      }

      const _ret = WEBEXT.sidePanel.getOptions(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetOptions": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 4)) {
          options_ffi["tabId"] = A.load.Int32(options + 0);
        }

        const _ret = WEBEXT.sidePanel.getOptions(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPanelBehavior": (): heap.Ref<boolean> => {
      if (WEBEXT?.sidePanel && "getPanelBehavior" in WEBEXT?.sidePanel) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPanelBehavior": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sidePanel.getPanelBehavior);
    },
    "call_GetPanelBehavior": (retPtr: Pointer): void => {
      const _ret = WEBEXT.sidePanel.getPanelBehavior();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPanelBehavior": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sidePanel.getPanelBehavior();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Open": (): heap.Ref<boolean> => {
      if (WEBEXT?.sidePanel && "open" in WEBEXT?.sidePanel) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Open": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sidePanel.open);
    },
    "call_Open": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 8)) {
        options_ffi["windowId"] = A.load.Int32(options + 0);
      }
      if (A.load.Bool(options + 9)) {
        options_ffi["tabId"] = A.load.Int32(options + 4);
      }

      const _ret = WEBEXT.sidePanel.open(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Open": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 8)) {
          options_ffi["windowId"] = A.load.Int32(options + 0);
        }
        if (A.load.Bool(options + 9)) {
          options_ffi["tabId"] = A.load.Int32(options + 4);
        }

        const _ret = WEBEXT.sidePanel.open(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetOptions": (): heap.Ref<boolean> => {
      if (WEBEXT?.sidePanel && "setOptions" in WEBEXT?.sidePanel) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetOptions": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sidePanel.setOptions);
    },
    "call_SetOptions": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 9)) {
        options_ffi["tabId"] = A.load.Int32(options + 0);
      }
      options_ffi["path"] = A.load.Ref(options + 4, undefined);
      if (A.load.Bool(options + 10)) {
        options_ffi["enabled"] = A.load.Bool(options + 8);
      }

      const _ret = WEBEXT.sidePanel.setOptions(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetOptions": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 9)) {
          options_ffi["tabId"] = A.load.Int32(options + 0);
        }
        options_ffi["path"] = A.load.Ref(options + 4, undefined);
        if (A.load.Bool(options + 10)) {
          options_ffi["enabled"] = A.load.Bool(options + 8);
        }

        const _ret = WEBEXT.sidePanel.setOptions(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPanelBehavior": (): heap.Ref<boolean> => {
      if (WEBEXT?.sidePanel && "setPanelBehavior" in WEBEXT?.sidePanel) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPanelBehavior": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sidePanel.setPanelBehavior);
    },
    "call_SetPanelBehavior": (retPtr: Pointer, behavior: Pointer): void => {
      const behavior_ffi = {};

      if (A.load.Bool(behavior + 1)) {
        behavior_ffi["openPanelOnActionClick"] = A.load.Bool(behavior + 0);
      }

      const _ret = WEBEXT.sidePanel.setPanelBehavior(behavior_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetPanelBehavior": (retPtr: Pointer, errPtr: Pointer, behavior: Pointer): heap.Ref<boolean> => {
      try {
        const behavior_ffi = {};

        if (A.load.Bool(behavior + 1)) {
          behavior_ffi["openPanelOnActionClick"] = A.load.Bool(behavior + 0);
        }

        const _ret = WEBEXT.sidePanel.setPanelBehavior(behavior_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
