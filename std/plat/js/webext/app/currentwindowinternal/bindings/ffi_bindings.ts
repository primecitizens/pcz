import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/app/currentwindowinternal", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Bounds": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Bool(ptr + 16, "left" in x ? true : false);
        A.store.Int32(ptr + 0, x["left"] === undefined ? 0 : (x["left"] as number));
        A.store.Bool(ptr + 17, "top" in x ? true : false);
        A.store.Int32(ptr + 4, x["top"] === undefined ? 0 : (x["top"] as number));
        A.store.Bool(ptr + 18, "width" in x ? true : false);
        A.store.Int32(ptr + 8, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Bool(ptr + 19, "height" in x ? true : false);
        A.store.Int32(ptr + 12, x["height"] === undefined ? 0 : (x["height"] as number));
      }
    },
    "load_Bounds": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["left"] = A.load.Int32(ptr + 0);
      } else {
        delete x["left"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["top"] = A.load.Int32(ptr + 4);
      } else {
        delete x["top"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["width"] = A.load.Int32(ptr + 8);
      } else {
        delete x["width"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["height"] = A.load.Int32(ptr + 12);
      } else {
        delete x["height"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RegionRect": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Bool(ptr + 16, "left" in x ? true : false);
        A.store.Int32(ptr + 0, x["left"] === undefined ? 0 : (x["left"] as number));
        A.store.Bool(ptr + 17, "top" in x ? true : false);
        A.store.Int32(ptr + 4, x["top"] === undefined ? 0 : (x["top"] as number));
        A.store.Bool(ptr + 18, "width" in x ? true : false);
        A.store.Int32(ptr + 8, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Bool(ptr + 19, "height" in x ? true : false);
        A.store.Int32(ptr + 12, x["height"] === undefined ? 0 : (x["height"] as number));
      }
    },
    "load_RegionRect": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["left"] = A.load.Int32(ptr + 0);
      } else {
        delete x["left"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["top"] = A.load.Int32(ptr + 4);
      } else {
        delete x["top"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["width"] = A.load.Int32(ptr + 8);
      } else {
        delete x["width"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["height"] = A.load.Int32(ptr + 12);
      } else {
        delete x["height"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Region": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["rects"]);
      }
    },
    "load_Region": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["rects"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SizeConstraints": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Bool(ptr + 16, "minWidth" in x ? true : false);
        A.store.Int32(ptr + 0, x["minWidth"] === undefined ? 0 : (x["minWidth"] as number));
        A.store.Bool(ptr + 17, "minHeight" in x ? true : false);
        A.store.Int32(ptr + 4, x["minHeight"] === undefined ? 0 : (x["minHeight"] as number));
        A.store.Bool(ptr + 18, "maxWidth" in x ? true : false);
        A.store.Int32(ptr + 8, x["maxWidth"] === undefined ? 0 : (x["maxWidth"] as number));
        A.store.Bool(ptr + 19, "maxHeight" in x ? true : false);
        A.store.Int32(ptr + 12, x["maxHeight"] === undefined ? 0 : (x["maxHeight"] as number));
      }
    },
    "load_SizeConstraints": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["minWidth"] = A.load.Int32(ptr + 0);
      } else {
        delete x["minWidth"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["minHeight"] = A.load.Int32(ptr + 4);
      } else {
        delete x["minHeight"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["maxWidth"] = A.load.Int32(ptr + 8);
      } else {
        delete x["maxWidth"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["maxHeight"] = A.load.Int32(ptr + 12);
      } else {
        delete x["maxHeight"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_ClearAttention": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "clearAttention" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ClearAttention": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.clearAttention);
    },
    "call_ClearAttention": (retPtr: Pointer): void => {
      const _ret = WEBEXT.app.currentWindowInternal.clearAttention();
    },
    "try_ClearAttention": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.clearAttention();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DrawAttention": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "drawAttention" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DrawAttention": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.drawAttention);
    },
    "call_DrawAttention": (retPtr: Pointer): void => {
      const _ret = WEBEXT.app.currentWindowInternal.drawAttention();
    },
    "try_DrawAttention": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.drawAttention();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Focus": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "focus" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Focus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.focus);
    },
    "call_Focus": (retPtr: Pointer): void => {
      const _ret = WEBEXT.app.currentWindowInternal.focus();
    },
    "try_Focus": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.focus();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Fullscreen": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "fullscreen" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Fullscreen": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.fullscreen);
    },
    "call_Fullscreen": (retPtr: Pointer): void => {
      const _ret = WEBEXT.app.currentWindowInternal.fullscreen();
    },
    "try_Fullscreen": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.fullscreen();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Hide": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "hide" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Hide": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.hide);
    },
    "call_Hide": (retPtr: Pointer): void => {
      const _ret = WEBEXT.app.currentWindowInternal.hide();
    },
    "try_Hide": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.hide();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Maximize": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "maximize" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Maximize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.maximize);
    },
    "call_Maximize": (retPtr: Pointer): void => {
      const _ret = WEBEXT.app.currentWindowInternal.maximize();
    },
    "try_Maximize": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.maximize();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Minimize": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "minimize" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Minimize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.minimize);
    },
    "call_Minimize": (retPtr: Pointer): void => {
      const _ret = WEBEXT.app.currentWindowInternal.minimize();
    },
    "try_Minimize": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.minimize();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAlphaEnabledChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onAlphaEnabledChanged &&
        "addListener" in WEBEXT?.app?.currentWindowInternal?.onAlphaEnabledChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAlphaEnabledChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.addListener);
    },
    "call_OnAlphaEnabledChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnAlphaEnabledChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAlphaEnabledChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onAlphaEnabledChanged &&
        "removeListener" in WEBEXT?.app?.currentWindowInternal?.onAlphaEnabledChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAlphaEnabledChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.removeListener);
    },
    "call_OffAlphaEnabledChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffAlphaEnabledChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAlphaEnabledChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onAlphaEnabledChanged &&
        "hasListener" in WEBEXT?.app?.currentWindowInternal?.onAlphaEnabledChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAlphaEnabledChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.hasListener);
    },
    "call_HasOnAlphaEnabledChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAlphaEnabledChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnBoundsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onBoundsChanged &&
        "addListener" in WEBEXT?.app?.currentWindowInternal?.onBoundsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onBoundsChanged.addListener);
    },
    "call_OnBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onBoundsChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnBoundsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onBoundsChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffBoundsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onBoundsChanged &&
        "removeListener" in WEBEXT?.app?.currentWindowInternal?.onBoundsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onBoundsChanged.removeListener);
    },
    "call_OffBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onBoundsChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffBoundsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onBoundsChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnBoundsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onBoundsChanged &&
        "hasListener" in WEBEXT?.app?.currentWindowInternal?.onBoundsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onBoundsChanged.hasListener);
    },
    "call_HasOnBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onBoundsChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnBoundsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onBoundsChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnClosed": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onClosed &&
        "addListener" in WEBEXT?.app?.currentWindowInternal?.onClosed
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClosed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onClosed.addListener);
    },
    "call_OnClosed": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onClosed.addListener(A.H.get<object>(callback));
    },
    "try_OnClosed": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onClosed.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClosed": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onClosed &&
        "removeListener" in WEBEXT?.app?.currentWindowInternal?.onClosed
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClosed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onClosed.removeListener);
    },
    "call_OffClosed": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onClosed.removeListener(A.H.get<object>(callback));
    },
    "try_OffClosed": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onClosed.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClosed": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onClosed &&
        "hasListener" in WEBEXT?.app?.currentWindowInternal?.onClosed
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClosed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onClosed.hasListener);
    },
    "call_HasOnClosed": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onClosed.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClosed": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onClosed.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnFullscreened": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onFullscreened &&
        "addListener" in WEBEXT?.app?.currentWindowInternal?.onFullscreened
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnFullscreened": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onFullscreened.addListener);
    },
    "call_OnFullscreened": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onFullscreened.addListener(A.H.get<object>(callback));
    },
    "try_OnFullscreened": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onFullscreened.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffFullscreened": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onFullscreened &&
        "removeListener" in WEBEXT?.app?.currentWindowInternal?.onFullscreened
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffFullscreened": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onFullscreened.removeListener);
    },
    "call_OffFullscreened": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onFullscreened.removeListener(A.H.get<object>(callback));
    },
    "try_OffFullscreened": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onFullscreened.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnFullscreened": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onFullscreened &&
        "hasListener" in WEBEXT?.app?.currentWindowInternal?.onFullscreened
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnFullscreened": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onFullscreened.hasListener);
    },
    "call_HasOnFullscreened": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onFullscreened.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnFullscreened": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onFullscreened.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMaximized": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onMaximized &&
        "addListener" in WEBEXT?.app?.currentWindowInternal?.onMaximized
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMaximized": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onMaximized.addListener);
    },
    "call_OnMaximized": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onMaximized.addListener(A.H.get<object>(callback));
    },
    "try_OnMaximized": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onMaximized.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMaximized": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onMaximized &&
        "removeListener" in WEBEXT?.app?.currentWindowInternal?.onMaximized
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMaximized": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onMaximized.removeListener);
    },
    "call_OffMaximized": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onMaximized.removeListener(A.H.get<object>(callback));
    },
    "try_OffMaximized": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onMaximized.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMaximized": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onMaximized &&
        "hasListener" in WEBEXT?.app?.currentWindowInternal?.onMaximized
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMaximized": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onMaximized.hasListener);
    },
    "call_HasOnMaximized": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onMaximized.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMaximized": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onMaximized.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMinimized": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onMinimized &&
        "addListener" in WEBEXT?.app?.currentWindowInternal?.onMinimized
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMinimized": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onMinimized.addListener);
    },
    "call_OnMinimized": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onMinimized.addListener(A.H.get<object>(callback));
    },
    "try_OnMinimized": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onMinimized.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMinimized": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onMinimized &&
        "removeListener" in WEBEXT?.app?.currentWindowInternal?.onMinimized
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMinimized": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onMinimized.removeListener);
    },
    "call_OffMinimized": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onMinimized.removeListener(A.H.get<object>(callback));
    },
    "try_OffMinimized": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onMinimized.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMinimized": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onMinimized &&
        "hasListener" in WEBEXT?.app?.currentWindowInternal?.onMinimized
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMinimized": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onMinimized.hasListener);
    },
    "call_HasOnMinimized": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onMinimized.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMinimized": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onMinimized.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRestored": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onRestored &&
        "addListener" in WEBEXT?.app?.currentWindowInternal?.onRestored
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRestored": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onRestored.addListener);
    },
    "call_OnRestored": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onRestored.addListener(A.H.get<object>(callback));
    },
    "try_OnRestored": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onRestored.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRestored": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onRestored &&
        "removeListener" in WEBEXT?.app?.currentWindowInternal?.onRestored
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRestored": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onRestored.removeListener);
    },
    "call_OffRestored": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onRestored.removeListener(A.H.get<object>(callback));
    },
    "try_OffRestored": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onRestored.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRestored": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.currentWindowInternal?.onRestored &&
        "hasListener" in WEBEXT?.app?.currentWindowInternal?.onRestored
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRestored": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.onRestored.hasListener);
    },
    "call_HasOnRestored": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.onRestored.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRestored": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.onRestored.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Restore": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "restore" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Restore": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.restore);
    },
    "call_Restore": (retPtr: Pointer): void => {
      const _ret = WEBEXT.app.currentWindowInternal.restore();
    },
    "try_Restore": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.restore();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetActivateOnPointer": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "setActivateOnPointer" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetActivateOnPointer": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.setActivateOnPointer);
    },
    "call_SetActivateOnPointer": (retPtr: Pointer, activate_on_pointer: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.setActivateOnPointer(activate_on_pointer === A.H.TRUE);
    },
    "try_SetActivateOnPointer": (
      retPtr: Pointer,
      errPtr: Pointer,
      activate_on_pointer: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.setActivateOnPointer(activate_on_pointer === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetAlwaysOnTop": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "setAlwaysOnTop" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetAlwaysOnTop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.setAlwaysOnTop);
    },
    "call_SetAlwaysOnTop": (retPtr: Pointer, always_on_top: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.setAlwaysOnTop(always_on_top === A.H.TRUE);
    },
    "try_SetAlwaysOnTop": (retPtr: Pointer, errPtr: Pointer, always_on_top: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.setAlwaysOnTop(always_on_top === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetBounds": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "setBounds" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetBounds": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.setBounds);
    },
    "call_SetBounds": (retPtr: Pointer, boundsType: heap.Ref<object>, bounds: Pointer): void => {
      const bounds_ffi = {};

      if (A.load.Bool(bounds + 16)) {
        bounds_ffi["left"] = A.load.Int32(bounds + 0);
      }
      if (A.load.Bool(bounds + 17)) {
        bounds_ffi["top"] = A.load.Int32(bounds + 4);
      }
      if (A.load.Bool(bounds + 18)) {
        bounds_ffi["width"] = A.load.Int32(bounds + 8);
      }
      if (A.load.Bool(bounds + 19)) {
        bounds_ffi["height"] = A.load.Int32(bounds + 12);
      }

      const _ret = WEBEXT.app.currentWindowInternal.setBounds(A.H.get<object>(boundsType), bounds_ffi);
    },
    "try_SetBounds": (
      retPtr: Pointer,
      errPtr: Pointer,
      boundsType: heap.Ref<object>,
      bounds: Pointer
    ): heap.Ref<boolean> => {
      try {
        const bounds_ffi = {};

        if (A.load.Bool(bounds + 16)) {
          bounds_ffi["left"] = A.load.Int32(bounds + 0);
        }
        if (A.load.Bool(bounds + 17)) {
          bounds_ffi["top"] = A.load.Int32(bounds + 4);
        }
        if (A.load.Bool(bounds + 18)) {
          bounds_ffi["width"] = A.load.Int32(bounds + 8);
        }
        if (A.load.Bool(bounds + 19)) {
          bounds_ffi["height"] = A.load.Int32(bounds + 12);
        }

        const _ret = WEBEXT.app.currentWindowInternal.setBounds(A.H.get<object>(boundsType), bounds_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetIcon": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "setIcon" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetIcon": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.setIcon);
    },
    "call_SetIcon": (retPtr: Pointer, icon_url: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.setIcon(A.H.get<object>(icon_url));
    },
    "try_SetIcon": (retPtr: Pointer, errPtr: Pointer, icon_url: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.setIcon(A.H.get<object>(icon_url));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetShape": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "setShape" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetShape": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.setShape);
    },
    "call_SetShape": (retPtr: Pointer, region: Pointer): void => {
      const region_ffi = {};

      region_ffi["rects"] = A.load.Ref(region + 0, undefined);

      const _ret = WEBEXT.app.currentWindowInternal.setShape(region_ffi);
    },
    "try_SetShape": (retPtr: Pointer, errPtr: Pointer, region: Pointer): heap.Ref<boolean> => {
      try {
        const region_ffi = {};

        region_ffi["rects"] = A.load.Ref(region + 0, undefined);

        const _ret = WEBEXT.app.currentWindowInternal.setShape(region_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetSizeConstraints": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "setSizeConstraints" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetSizeConstraints": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.setSizeConstraints);
    },
    "call_SetSizeConstraints": (retPtr: Pointer, boundsType: heap.Ref<object>, constraints: Pointer): void => {
      const constraints_ffi = {};

      if (A.load.Bool(constraints + 16)) {
        constraints_ffi["minWidth"] = A.load.Int32(constraints + 0);
      }
      if (A.load.Bool(constraints + 17)) {
        constraints_ffi["minHeight"] = A.load.Int32(constraints + 4);
      }
      if (A.load.Bool(constraints + 18)) {
        constraints_ffi["maxWidth"] = A.load.Int32(constraints + 8);
      }
      if (A.load.Bool(constraints + 19)) {
        constraints_ffi["maxHeight"] = A.load.Int32(constraints + 12);
      }

      const _ret = WEBEXT.app.currentWindowInternal.setSizeConstraints(A.H.get<object>(boundsType), constraints_ffi);
    },
    "try_SetSizeConstraints": (
      retPtr: Pointer,
      errPtr: Pointer,
      boundsType: heap.Ref<object>,
      constraints: Pointer
    ): heap.Ref<boolean> => {
      try {
        const constraints_ffi = {};

        if (A.load.Bool(constraints + 16)) {
          constraints_ffi["minWidth"] = A.load.Int32(constraints + 0);
        }
        if (A.load.Bool(constraints + 17)) {
          constraints_ffi["minHeight"] = A.load.Int32(constraints + 4);
        }
        if (A.load.Bool(constraints + 18)) {
          constraints_ffi["maxWidth"] = A.load.Int32(constraints + 8);
        }
        if (A.load.Bool(constraints + 19)) {
          constraints_ffi["maxHeight"] = A.load.Int32(constraints + 12);
        }

        const _ret = WEBEXT.app.currentWindowInternal.setSizeConstraints(A.H.get<object>(boundsType), constraints_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetVisibleOnAllWorkspaces": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "setVisibleOnAllWorkspaces" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetVisibleOnAllWorkspaces": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.setVisibleOnAllWorkspaces);
    },
    "call_SetVisibleOnAllWorkspaces": (retPtr: Pointer, always_visible: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.setVisibleOnAllWorkspaces(always_visible === A.H.TRUE);
    },
    "try_SetVisibleOnAllWorkspaces": (
      retPtr: Pointer,
      errPtr: Pointer,
      always_visible: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.setVisibleOnAllWorkspaces(always_visible === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Show": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.currentWindowInternal && "show" in WEBEXT?.app?.currentWindowInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Show": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.currentWindowInternal.show);
    },
    "call_Show": (retPtr: Pointer, focused: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.app.currentWindowInternal.show(focused === A.H.TRUE);
    },
    "try_Show": (retPtr: Pointer, errPtr: Pointer, focused: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.currentWindowInternal.show(focused === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
