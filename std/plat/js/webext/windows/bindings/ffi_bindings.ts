import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/windows", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_WindowState": (ref: heap.Ref<string>): number => {
      const idx = ["normal", "minimized", "maximized", "fullscreen", "locked-fullscreen"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_CreateType": (ref: heap.Ref<string>): number => {
      const idx = ["normal", "popup", "panel"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_CreateArgCreateData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 80, false);
        A.store.Bool(ptr + 72, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 73, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Bool(ptr + 74, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 75, false);
        A.store.Int64(ptr + 24, 0);
        A.store.Bool(ptr + 76, false);
        A.store.Bool(ptr + 32, false);
        A.store.Enum(ptr + 36, -1);
        A.store.Bool(ptr + 77, false);
        A.store.Int64(ptr + 40, 0);
        A.store.Bool(ptr + 78, false);
        A.store.Int64(ptr + 48, 0);
        A.store.Enum(ptr + 56, -1);
        A.store.Ref(ptr + 60, undefined);
        A.store.Bool(ptr + 79, false);
        A.store.Int64(ptr + 64, 0);
      } else {
        A.store.Bool(ptr + 80, true);
        A.store.Bool(ptr + 72, "focused" in x ? true : false);
        A.store.Bool(ptr + 0, x["focused"] ? true : false);
        A.store.Bool(ptr + 73, "height" in x ? true : false);
        A.store.Int64(ptr + 8, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Bool(ptr + 74, "incognito" in x ? true : false);
        A.store.Bool(ptr + 16, x["incognito"] ? true : false);
        A.store.Bool(ptr + 75, "left" in x ? true : false);
        A.store.Int64(ptr + 24, x["left"] === undefined ? 0 : (x["left"] as number));
        A.store.Bool(ptr + 76, "setSelfAsOpener" in x ? true : false);
        A.store.Bool(ptr + 32, x["setSelfAsOpener"] ? true : false);
        A.store.Enum(
          ptr + 36,
          ["normal", "minimized", "maximized", "fullscreen", "locked-fullscreen"].indexOf(x["state"] as string)
        );
        A.store.Bool(ptr + 77, "tabId" in x ? true : false);
        A.store.Int64(ptr + 40, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Bool(ptr + 78, "top" in x ? true : false);
        A.store.Int64(ptr + 48, x["top"] === undefined ? 0 : (x["top"] as number));
        A.store.Enum(ptr + 56, ["normal", "popup", "panel"].indexOf(x["type"] as string));
        A.store.Ref(ptr + 60, x["url"]);
        A.store.Bool(ptr + 79, "width" in x ? true : false);
        A.store.Int64(ptr + 64, x["width"] === undefined ? 0 : (x["width"] as number));
      }
    },
    "load_CreateArgCreateData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 72)) {
        x["focused"] = A.load.Bool(ptr + 0);
      } else {
        delete x["focused"];
      }
      if (A.load.Bool(ptr + 73)) {
        x["height"] = A.load.Int64(ptr + 8);
      } else {
        delete x["height"];
      }
      if (A.load.Bool(ptr + 74)) {
        x["incognito"] = A.load.Bool(ptr + 16);
      } else {
        delete x["incognito"];
      }
      if (A.load.Bool(ptr + 75)) {
        x["left"] = A.load.Int64(ptr + 24);
      } else {
        delete x["left"];
      }
      if (A.load.Bool(ptr + 76)) {
        x["setSelfAsOpener"] = A.load.Bool(ptr + 32);
      } else {
        delete x["setSelfAsOpener"];
      }
      x["state"] = A.load.Enum(ptr + 36, ["normal", "minimized", "maximized", "fullscreen", "locked-fullscreen"]);
      if (A.load.Bool(ptr + 77)) {
        x["tabId"] = A.load.Int64(ptr + 40);
      } else {
        delete x["tabId"];
      }
      if (A.load.Bool(ptr + 78)) {
        x["top"] = A.load.Int64(ptr + 48);
      } else {
        delete x["top"];
      }
      x["type"] = A.load.Enum(ptr + 56, ["normal", "popup", "panel"]);
      x["url"] = A.load.Ref(ptr + 60, undefined);
      if (A.load.Bool(ptr + 79)) {
        x["width"] = A.load.Int64(ptr + 64);
      } else {
        delete x["width"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_WindowType": (ref: heap.Ref<string>): number => {
      const idx = ["normal", "popup", "panel", "app", "devtools"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Window": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 85, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 80, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Bool(ptr + 81, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 82, false);
        A.store.Int64(ptr + 32, 0);
        A.store.Ref(ptr + 40, undefined);
        A.store.Enum(ptr + 44, -1);
        A.store.Ref(ptr + 48, undefined);
        A.store.Bool(ptr + 83, false);
        A.store.Int64(ptr + 56, 0);
        A.store.Enum(ptr + 64, -1);
        A.store.Bool(ptr + 84, false);
        A.store.Int64(ptr + 72, 0);
      } else {
        A.store.Bool(ptr + 85, true);
        A.store.Bool(ptr + 0, x["alwaysOnTop"] ? true : false);
        A.store.Bool(ptr + 1, x["focused"] ? true : false);
        A.store.Bool(ptr + 80, "height" in x ? true : false);
        A.store.Int64(ptr + 8, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Bool(ptr + 81, "id" in x ? true : false);
        A.store.Int64(ptr + 16, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Bool(ptr + 24, x["incognito"] ? true : false);
        A.store.Bool(ptr + 82, "left" in x ? true : false);
        A.store.Int64(ptr + 32, x["left"] === undefined ? 0 : (x["left"] as number));
        A.store.Ref(ptr + 40, x["sessionId"]);
        A.store.Enum(
          ptr + 44,
          ["normal", "minimized", "maximized", "fullscreen", "locked-fullscreen"].indexOf(x["state"] as string)
        );
        A.store.Ref(ptr + 48, x["tabs"]);
        A.store.Bool(ptr + 83, "top" in x ? true : false);
        A.store.Int64(ptr + 56, x["top"] === undefined ? 0 : (x["top"] as number));
        A.store.Enum(ptr + 64, ["normal", "popup", "panel", "app", "devtools"].indexOf(x["type"] as string));
        A.store.Bool(ptr + 84, "width" in x ? true : false);
        A.store.Int64(ptr + 72, x["width"] === undefined ? 0 : (x["width"] as number));
      }
    },
    "load_Window": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["alwaysOnTop"] = A.load.Bool(ptr + 0);
      x["focused"] = A.load.Bool(ptr + 1);
      if (A.load.Bool(ptr + 80)) {
        x["height"] = A.load.Int64(ptr + 8);
      } else {
        delete x["height"];
      }
      if (A.load.Bool(ptr + 81)) {
        x["id"] = A.load.Int64(ptr + 16);
      } else {
        delete x["id"];
      }
      x["incognito"] = A.load.Bool(ptr + 24);
      if (A.load.Bool(ptr + 82)) {
        x["left"] = A.load.Int64(ptr + 32);
      } else {
        delete x["left"];
      }
      x["sessionId"] = A.load.Ref(ptr + 40, undefined);
      x["state"] = A.load.Enum(ptr + 44, ["normal", "minimized", "maximized", "fullscreen", "locked-fullscreen"]);
      x["tabs"] = A.load.Ref(ptr + 48, undefined);
      if (A.load.Bool(ptr + 83)) {
        x["top"] = A.load.Int64(ptr + 56);
      } else {
        delete x["top"];
      }
      x["type"] = A.load.Enum(ptr + 64, ["normal", "popup", "panel", "app", "devtools"]);
      if (A.load.Bool(ptr + 84)) {
        x["width"] = A.load.Int64(ptr + 72);
      } else {
        delete x["width"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_QueryOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "populate" in x ? true : false);
        A.store.Bool(ptr + 0, x["populate"] ? true : false);
        A.store.Ref(ptr + 4, x["windowTypes"]);
      }
    },
    "load_QueryOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["populate"] = A.load.Bool(ptr + 0);
      } else {
        delete x["populate"];
      }
      x["windowTypes"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UpdateArgUpdateInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 54, false);
        A.store.Bool(ptr + 48, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 49, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 50, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Bool(ptr + 51, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Enum(ptr + 24, -1);
        A.store.Bool(ptr + 52, false);
        A.store.Int64(ptr + 32, 0);
        A.store.Bool(ptr + 53, false);
        A.store.Int64(ptr + 40, 0);
      } else {
        A.store.Bool(ptr + 54, true);
        A.store.Bool(ptr + 48, "drawAttention" in x ? true : false);
        A.store.Bool(ptr + 0, x["drawAttention"] ? true : false);
        A.store.Bool(ptr + 49, "focused" in x ? true : false);
        A.store.Bool(ptr + 1, x["focused"] ? true : false);
        A.store.Bool(ptr + 50, "height" in x ? true : false);
        A.store.Int64(ptr + 8, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Bool(ptr + 51, "left" in x ? true : false);
        A.store.Int64(ptr + 16, x["left"] === undefined ? 0 : (x["left"] as number));
        A.store.Enum(
          ptr + 24,
          ["normal", "minimized", "maximized", "fullscreen", "locked-fullscreen"].indexOf(x["state"] as string)
        );
        A.store.Bool(ptr + 52, "top" in x ? true : false);
        A.store.Int64(ptr + 32, x["top"] === undefined ? 0 : (x["top"] as number));
        A.store.Bool(ptr + 53, "width" in x ? true : false);
        A.store.Int64(ptr + 40, x["width"] === undefined ? 0 : (x["width"] as number));
      }
    },
    "load_UpdateArgUpdateInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 48)) {
        x["drawAttention"] = A.load.Bool(ptr + 0);
      } else {
        delete x["drawAttention"];
      }
      if (A.load.Bool(ptr + 49)) {
        x["focused"] = A.load.Bool(ptr + 1);
      } else {
        delete x["focused"];
      }
      if (A.load.Bool(ptr + 50)) {
        x["height"] = A.load.Int64(ptr + 8);
      } else {
        delete x["height"];
      }
      if (A.load.Bool(ptr + 51)) {
        x["left"] = A.load.Int64(ptr + 16);
      } else {
        delete x["left"];
      }
      x["state"] = A.load.Enum(ptr + 24, ["normal", "minimized", "maximized", "fullscreen", "locked-fullscreen"]);
      if (A.load.Bool(ptr + 52)) {
        x["top"] = A.load.Int64(ptr + 32);
      } else {
        delete x["top"];
      }
      if (A.load.Bool(ptr + 53)) {
        x["width"] = A.load.Int64(ptr + 40);
      } else {
        delete x["width"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "get_WINDOW_ID_CURRENT": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.windows && "WINDOW_ID_CURRENT" in WEBEXT?.windows) {
        const val = WEBEXT.windows.WINDOW_ID_CURRENT;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_WINDOW_ID_CURRENT": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.windows, "WINDOW_ID_CURRENT", A.H.get<object>(val), WEBEXT.windows)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_WINDOW_ID_NONE": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.windows && "WINDOW_ID_NONE" in WEBEXT?.windows) {
        const val = WEBEXT.windows.WINDOW_ID_NONE;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_WINDOW_ID_NONE": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.windows, "WINDOW_ID_NONE", A.H.get<object>(val), WEBEXT.windows) ? A.H.TRUE : A.H.FALSE;
    },
    "has_Create": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows && "create" in WEBEXT?.windows) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Create": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.create);
    },
    "call_Create": (retPtr: Pointer, createData: Pointer): void => {
      const createData_ffi = {};

      if (A.load.Bool(createData + 72)) {
        createData_ffi["focused"] = A.load.Bool(createData + 0);
      }
      if (A.load.Bool(createData + 73)) {
        createData_ffi["height"] = A.load.Int64(createData + 8);
      }
      if (A.load.Bool(createData + 74)) {
        createData_ffi["incognito"] = A.load.Bool(createData + 16);
      }
      if (A.load.Bool(createData + 75)) {
        createData_ffi["left"] = A.load.Int64(createData + 24);
      }
      if (A.load.Bool(createData + 76)) {
        createData_ffi["setSelfAsOpener"] = A.load.Bool(createData + 32);
      }
      createData_ffi["state"] = A.load.Enum(createData + 36, [
        "normal",
        "minimized",
        "maximized",
        "fullscreen",
        "locked-fullscreen",
      ]);
      if (A.load.Bool(createData + 77)) {
        createData_ffi["tabId"] = A.load.Int64(createData + 40);
      }
      if (A.load.Bool(createData + 78)) {
        createData_ffi["top"] = A.load.Int64(createData + 48);
      }
      createData_ffi["type"] = A.load.Enum(createData + 56, ["normal", "popup", "panel"]);
      createData_ffi["url"] = A.load.Ref(createData + 60, undefined);
      if (A.load.Bool(createData + 79)) {
        createData_ffi["width"] = A.load.Int64(createData + 64);
      }

      const _ret = WEBEXT.windows.create(createData_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Create": (retPtr: Pointer, errPtr: Pointer, createData: Pointer): heap.Ref<boolean> => {
      try {
        const createData_ffi = {};

        if (A.load.Bool(createData + 72)) {
          createData_ffi["focused"] = A.load.Bool(createData + 0);
        }
        if (A.load.Bool(createData + 73)) {
          createData_ffi["height"] = A.load.Int64(createData + 8);
        }
        if (A.load.Bool(createData + 74)) {
          createData_ffi["incognito"] = A.load.Bool(createData + 16);
        }
        if (A.load.Bool(createData + 75)) {
          createData_ffi["left"] = A.load.Int64(createData + 24);
        }
        if (A.load.Bool(createData + 76)) {
          createData_ffi["setSelfAsOpener"] = A.load.Bool(createData + 32);
        }
        createData_ffi["state"] = A.load.Enum(createData + 36, [
          "normal",
          "minimized",
          "maximized",
          "fullscreen",
          "locked-fullscreen",
        ]);
        if (A.load.Bool(createData + 77)) {
          createData_ffi["tabId"] = A.load.Int64(createData + 40);
        }
        if (A.load.Bool(createData + 78)) {
          createData_ffi["top"] = A.load.Int64(createData + 48);
        }
        createData_ffi["type"] = A.load.Enum(createData + 56, ["normal", "popup", "panel"]);
        createData_ffi["url"] = A.load.Ref(createData + 60, undefined);
        if (A.load.Bool(createData + 79)) {
          createData_ffi["width"] = A.load.Int64(createData + 64);
        }

        const _ret = WEBEXT.windows.create(createData_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Get": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows && "get" in WEBEXT?.windows) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Get": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.get);
    },
    "call_Get": (retPtr: Pointer, windowId: number, queryOptions: Pointer): void => {
      const queryOptions_ffi = {};

      if (A.load.Bool(queryOptions + 8)) {
        queryOptions_ffi["populate"] = A.load.Bool(queryOptions + 0);
      }
      queryOptions_ffi["windowTypes"] = A.load.Ref(queryOptions + 4, undefined);

      const _ret = WEBEXT.windows.get(windowId, queryOptions_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Get": (retPtr: Pointer, errPtr: Pointer, windowId: number, queryOptions: Pointer): heap.Ref<boolean> => {
      try {
        const queryOptions_ffi = {};

        if (A.load.Bool(queryOptions + 8)) {
          queryOptions_ffi["populate"] = A.load.Bool(queryOptions + 0);
        }
        queryOptions_ffi["windowTypes"] = A.load.Ref(queryOptions + 4, undefined);

        const _ret = WEBEXT.windows.get(windowId, queryOptions_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAll": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows && "getAll" in WEBEXT?.windows) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.getAll);
    },
    "call_GetAll": (retPtr: Pointer, queryOptions: Pointer): void => {
      const queryOptions_ffi = {};

      if (A.load.Bool(queryOptions + 8)) {
        queryOptions_ffi["populate"] = A.load.Bool(queryOptions + 0);
      }
      queryOptions_ffi["windowTypes"] = A.load.Ref(queryOptions + 4, undefined);

      const _ret = WEBEXT.windows.getAll(queryOptions_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAll": (retPtr: Pointer, errPtr: Pointer, queryOptions: Pointer): heap.Ref<boolean> => {
      try {
        const queryOptions_ffi = {};

        if (A.load.Bool(queryOptions + 8)) {
          queryOptions_ffi["populate"] = A.load.Bool(queryOptions + 0);
        }
        queryOptions_ffi["windowTypes"] = A.load.Ref(queryOptions + 4, undefined);

        const _ret = WEBEXT.windows.getAll(queryOptions_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCurrent": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows && "getCurrent" in WEBEXT?.windows) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCurrent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.getCurrent);
    },
    "call_GetCurrent": (retPtr: Pointer, queryOptions: Pointer): void => {
      const queryOptions_ffi = {};

      if (A.load.Bool(queryOptions + 8)) {
        queryOptions_ffi["populate"] = A.load.Bool(queryOptions + 0);
      }
      queryOptions_ffi["windowTypes"] = A.load.Ref(queryOptions + 4, undefined);

      const _ret = WEBEXT.windows.getCurrent(queryOptions_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCurrent": (retPtr: Pointer, errPtr: Pointer, queryOptions: Pointer): heap.Ref<boolean> => {
      try {
        const queryOptions_ffi = {};

        if (A.load.Bool(queryOptions + 8)) {
          queryOptions_ffi["populate"] = A.load.Bool(queryOptions + 0);
        }
        queryOptions_ffi["windowTypes"] = A.load.Ref(queryOptions + 4, undefined);

        const _ret = WEBEXT.windows.getCurrent(queryOptions_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetLastFocused": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows && "getLastFocused" in WEBEXT?.windows) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetLastFocused": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.getLastFocused);
    },
    "call_GetLastFocused": (retPtr: Pointer, queryOptions: Pointer): void => {
      const queryOptions_ffi = {};

      if (A.load.Bool(queryOptions + 8)) {
        queryOptions_ffi["populate"] = A.load.Bool(queryOptions + 0);
      }
      queryOptions_ffi["windowTypes"] = A.load.Ref(queryOptions + 4, undefined);

      const _ret = WEBEXT.windows.getLastFocused(queryOptions_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetLastFocused": (retPtr: Pointer, errPtr: Pointer, queryOptions: Pointer): heap.Ref<boolean> => {
      try {
        const queryOptions_ffi = {};

        if (A.load.Bool(queryOptions + 8)) {
          queryOptions_ffi["populate"] = A.load.Bool(queryOptions + 0);
        }
        queryOptions_ffi["windowTypes"] = A.load.Ref(queryOptions + 4, undefined);

        const _ret = WEBEXT.windows.getLastFocused(queryOptions_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnBoundsChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows?.onBoundsChanged && "addListener" in WEBEXT?.windows?.onBoundsChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.onBoundsChanged.addListener);
    },
    "call_OnBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.windows.onBoundsChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnBoundsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.windows.onBoundsChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffBoundsChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows?.onBoundsChanged && "removeListener" in WEBEXT?.windows?.onBoundsChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.onBoundsChanged.removeListener);
    },
    "call_OffBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.windows.onBoundsChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffBoundsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.windows.onBoundsChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnBoundsChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows?.onBoundsChanged && "hasListener" in WEBEXT?.windows?.onBoundsChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.onBoundsChanged.hasListener);
    },
    "call_HasOnBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.windows.onBoundsChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnBoundsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.windows.onBoundsChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows?.onCreated && "addListener" in WEBEXT?.windows?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.onCreated.addListener);
    },
    "call_OnCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.windows.onCreated.addListener(A.H.get<object>(callback));
    },
    "try_OnCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.windows.onCreated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows?.onCreated && "removeListener" in WEBEXT?.windows?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.onCreated.removeListener);
    },
    "call_OffCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.windows.onCreated.removeListener(A.H.get<object>(callback));
    },
    "try_OffCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.windows.onCreated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows?.onCreated && "hasListener" in WEBEXT?.windows?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.onCreated.hasListener);
    },
    "call_HasOnCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.windows.onCreated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.windows.onCreated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnFocusChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows?.onFocusChanged && "addListener" in WEBEXT?.windows?.onFocusChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnFocusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.onFocusChanged.addListener);
    },
    "call_OnFocusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.windows.onFocusChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnFocusChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.windows.onFocusChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffFocusChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows?.onFocusChanged && "removeListener" in WEBEXT?.windows?.onFocusChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffFocusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.onFocusChanged.removeListener);
    },
    "call_OffFocusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.windows.onFocusChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffFocusChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.windows.onFocusChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnFocusChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows?.onFocusChanged && "hasListener" in WEBEXT?.windows?.onFocusChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnFocusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.onFocusChanged.hasListener);
    },
    "call_HasOnFocusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.windows.onFocusChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnFocusChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.windows.onFocusChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows?.onRemoved && "addListener" in WEBEXT?.windows?.onRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.onRemoved.addListener);
    },
    "call_OnRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.windows.onRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.windows.onRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows?.onRemoved && "removeListener" in WEBEXT?.windows?.onRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.onRemoved.removeListener);
    },
    "call_OffRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.windows.onRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.windows.onRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows?.onRemoved && "hasListener" in WEBEXT?.windows?.onRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.onRemoved.hasListener);
    },
    "call_HasOnRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.windows.onRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.windows.onRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Remove": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows && "remove" in WEBEXT?.windows) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Remove": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.remove);
    },
    "call_Remove": (retPtr: Pointer, windowId: number): void => {
      const _ret = WEBEXT.windows.remove(windowId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Remove": (retPtr: Pointer, errPtr: Pointer, windowId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.windows.remove(windowId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Update": (): heap.Ref<boolean> => {
      if (WEBEXT?.windows && "update" in WEBEXT?.windows) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Update": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.windows.update);
    },
    "call_Update": (retPtr: Pointer, windowId: number, updateInfo: Pointer): void => {
      const updateInfo_ffi = {};

      if (A.load.Bool(updateInfo + 48)) {
        updateInfo_ffi["drawAttention"] = A.load.Bool(updateInfo + 0);
      }
      if (A.load.Bool(updateInfo + 49)) {
        updateInfo_ffi["focused"] = A.load.Bool(updateInfo + 1);
      }
      if (A.load.Bool(updateInfo + 50)) {
        updateInfo_ffi["height"] = A.load.Int64(updateInfo + 8);
      }
      if (A.load.Bool(updateInfo + 51)) {
        updateInfo_ffi["left"] = A.load.Int64(updateInfo + 16);
      }
      updateInfo_ffi["state"] = A.load.Enum(updateInfo + 24, [
        "normal",
        "minimized",
        "maximized",
        "fullscreen",
        "locked-fullscreen",
      ]);
      if (A.load.Bool(updateInfo + 52)) {
        updateInfo_ffi["top"] = A.load.Int64(updateInfo + 32);
      }
      if (A.load.Bool(updateInfo + 53)) {
        updateInfo_ffi["width"] = A.load.Int64(updateInfo + 40);
      }

      const _ret = WEBEXT.windows.update(windowId, updateInfo_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Update": (retPtr: Pointer, errPtr: Pointer, windowId: number, updateInfo: Pointer): heap.Ref<boolean> => {
      try {
        const updateInfo_ffi = {};

        if (A.load.Bool(updateInfo + 48)) {
          updateInfo_ffi["drawAttention"] = A.load.Bool(updateInfo + 0);
        }
        if (A.load.Bool(updateInfo + 49)) {
          updateInfo_ffi["focused"] = A.load.Bool(updateInfo + 1);
        }
        if (A.load.Bool(updateInfo + 50)) {
          updateInfo_ffi["height"] = A.load.Int64(updateInfo + 8);
        }
        if (A.load.Bool(updateInfo + 51)) {
          updateInfo_ffi["left"] = A.load.Int64(updateInfo + 16);
        }
        updateInfo_ffi["state"] = A.load.Enum(updateInfo + 24, [
          "normal",
          "minimized",
          "maximized",
          "fullscreen",
          "locked-fullscreen",
        ]);
        if (A.load.Bool(updateInfo + 52)) {
          updateInfo_ffi["top"] = A.load.Int64(updateInfo + 32);
        }
        if (A.load.Bool(updateInfo + 53)) {
          updateInfo_ffi["width"] = A.load.Int64(updateInfo + 40);
        }

        const _ret = WEBEXT.windows.update(windowId, updateInfo_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
