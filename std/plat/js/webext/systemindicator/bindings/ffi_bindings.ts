import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/systemindicator", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_SetIconDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["path"]);
        A.store.Ref(ptr + 4, x["imageData"]);
      }
    },
    "load_SetIconDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["path"] = A.load.Ref(ptr + 0, undefined);
      x["imageData"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Disable": (): heap.Ref<boolean> => {
      if (WEBEXT?.systemIndicator && "disable" in WEBEXT?.systemIndicator) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Disable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.systemIndicator.disable);
    },
    "call_Disable": (retPtr: Pointer): void => {
      const _ret = WEBEXT.systemIndicator.disable();
    },
    "try_Disable": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.systemIndicator.disable();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Enable": (): heap.Ref<boolean> => {
      if (WEBEXT?.systemIndicator && "enable" in WEBEXT?.systemIndicator) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Enable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.systemIndicator.enable);
    },
    "call_Enable": (retPtr: Pointer): void => {
      const _ret = WEBEXT.systemIndicator.enable();
    },
    "try_Enable": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.systemIndicator.enable();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.systemIndicator?.onClicked && "addListener" in WEBEXT?.systemIndicator?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.systemIndicator.onClicked.addListener);
    },
    "call_OnClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.systemIndicator.onClicked.addListener(A.H.get<object>(callback));
    },
    "try_OnClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.systemIndicator.onClicked.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.systemIndicator?.onClicked && "removeListener" in WEBEXT?.systemIndicator?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.systemIndicator.onClicked.removeListener);
    },
    "call_OffClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.systemIndicator.onClicked.removeListener(A.H.get<object>(callback));
    },
    "try_OffClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.systemIndicator.onClicked.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.systemIndicator?.onClicked && "hasListener" in WEBEXT?.systemIndicator?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.systemIndicator.onClicked.hasListener);
    },
    "call_HasOnClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.systemIndicator.onClicked.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.systemIndicator.onClicked.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetIcon": (): heap.Ref<boolean> => {
      if (WEBEXT?.systemIndicator && "setIcon" in WEBEXT?.systemIndicator) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetIcon": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.systemIndicator.setIcon);
    },
    "call_SetIcon": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["path"] = A.load.Ref(details + 0, undefined);
      details_ffi["imageData"] = A.load.Ref(details + 4, undefined);

      const _ret = WEBEXT.systemIndicator.setIcon(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetIcon": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["path"] = A.load.Ref(details + 0, undefined);
        details_ffi["imageData"] = A.load.Ref(details + 4, undefined);

        const _ret = WEBEXT.systemIndicator.setIcon(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
