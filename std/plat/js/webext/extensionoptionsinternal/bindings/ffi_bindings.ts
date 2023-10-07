import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/extensionoptionsinternal", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_PreferredSizeChangedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Bool(ptr + 16, "width" in x ? true : false);
        A.store.Float64(ptr + 0, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Bool(ptr + 17, "height" in x ? true : false);
        A.store.Float64(ptr + 8, x["height"] === undefined ? 0 : (x["height"] as number));
      }
    },
    "load_PreferredSizeChangedOptions": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["width"] = A.load.Float64(ptr + 0);
      } else {
        delete x["width"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["height"] = A.load.Float64(ptr + 8);
      } else {
        delete x["height"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SizeChangedOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 18, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 19, false);
        A.store.Int32(ptr + 12, 0);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Bool(ptr + 16, "oldWidth" in x ? true : false);
        A.store.Int32(ptr + 0, x["oldWidth"] === undefined ? 0 : (x["oldWidth"] as number));
        A.store.Bool(ptr + 17, "oldHeight" in x ? true : false);
        A.store.Int32(ptr + 4, x["oldHeight"] === undefined ? 0 : (x["oldHeight"] as number));
        A.store.Bool(ptr + 18, "newWidth" in x ? true : false);
        A.store.Int32(ptr + 8, x["newWidth"] === undefined ? 0 : (x["newWidth"] as number));
        A.store.Bool(ptr + 19, "newHeight" in x ? true : false);
        A.store.Int32(ptr + 12, x["newHeight"] === undefined ? 0 : (x["newHeight"] as number));
      }
    },
    "load_SizeChangedOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["oldWidth"] = A.load.Int32(ptr + 0);
      } else {
        delete x["oldWidth"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["oldHeight"] = A.load.Int32(ptr + 4);
      } else {
        delete x["oldHeight"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["newWidth"] = A.load.Int32(ptr + 8);
      } else {
        delete x["newWidth"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["newHeight"] = A.load.Int32(ptr + 12);
      } else {
        delete x["newHeight"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_OnClose": (): heap.Ref<boolean> => {
      if (WEBEXT?.extensionOptionsInternal?.onClose && "addListener" in WEBEXT?.extensionOptionsInternal?.onClose) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClose": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extensionOptionsInternal.onClose.addListener);
    },
    "call_OnClose": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.extensionOptionsInternal.onClose.addListener(A.H.get<object>(callback));
    },
    "try_OnClose": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extensionOptionsInternal.onClose.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClose": (): heap.Ref<boolean> => {
      if (WEBEXT?.extensionOptionsInternal?.onClose && "removeListener" in WEBEXT?.extensionOptionsInternal?.onClose) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClose": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extensionOptionsInternal.onClose.removeListener);
    },
    "call_OffClose": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.extensionOptionsInternal.onClose.removeListener(A.H.get<object>(callback));
    },
    "try_OffClose": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extensionOptionsInternal.onClose.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClose": (): heap.Ref<boolean> => {
      if (WEBEXT?.extensionOptionsInternal?.onClose && "hasListener" in WEBEXT?.extensionOptionsInternal?.onClose) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClose": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extensionOptionsInternal.onClose.hasListener);
    },
    "call_HasOnClose": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.extensionOptionsInternal.onClose.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClose": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extensionOptionsInternal.onClose.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnLoad": (): heap.Ref<boolean> => {
      if (WEBEXT?.extensionOptionsInternal?.onLoad && "addListener" in WEBEXT?.extensionOptionsInternal?.onLoad) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnLoad": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extensionOptionsInternal.onLoad.addListener);
    },
    "call_OnLoad": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.extensionOptionsInternal.onLoad.addListener(A.H.get<object>(callback));
    },
    "try_OnLoad": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extensionOptionsInternal.onLoad.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffLoad": (): heap.Ref<boolean> => {
      if (WEBEXT?.extensionOptionsInternal?.onLoad && "removeListener" in WEBEXT?.extensionOptionsInternal?.onLoad) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffLoad": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extensionOptionsInternal.onLoad.removeListener);
    },
    "call_OffLoad": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.extensionOptionsInternal.onLoad.removeListener(A.H.get<object>(callback));
    },
    "try_OffLoad": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extensionOptionsInternal.onLoad.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnLoad": (): heap.Ref<boolean> => {
      if (WEBEXT?.extensionOptionsInternal?.onLoad && "hasListener" in WEBEXT?.extensionOptionsInternal?.onLoad) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnLoad": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extensionOptionsInternal.onLoad.hasListener);
    },
    "call_HasOnLoad": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.extensionOptionsInternal.onLoad.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnLoad": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extensionOptionsInternal.onLoad.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPreferredSizeChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.extensionOptionsInternal?.onPreferredSizeChanged &&
        "addListener" in WEBEXT?.extensionOptionsInternal?.onPreferredSizeChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPreferredSizeChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.addListener);
    },
    "call_OnPreferredSizeChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnPreferredSizeChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPreferredSizeChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.extensionOptionsInternal?.onPreferredSizeChanged &&
        "removeListener" in WEBEXT?.extensionOptionsInternal?.onPreferredSizeChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPreferredSizeChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.removeListener);
    },
    "call_OffPreferredSizeChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffPreferredSizeChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPreferredSizeChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.extensionOptionsInternal?.onPreferredSizeChanged &&
        "hasListener" in WEBEXT?.extensionOptionsInternal?.onPreferredSizeChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPreferredSizeChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.hasListener);
    },
    "call_HasOnPreferredSizeChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPreferredSizeChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
