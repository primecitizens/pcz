import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/app/window", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_ContentBounds": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_ContentBounds": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
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

    "has_Bounds_SetPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "setPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Bounds_SetPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).setPosition);
    },

    "call_Bounds_SetPosition": (self: heap.Ref<object>, retPtr: Pointer, left: number, top: number): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.setPosition(left, top);
    },
    "try_Bounds_SetPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      left: number,
      top: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.setPosition(left, top);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Bounds_SetSize": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "setSize" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Bounds_SetSize": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).setSize);
    },

    "call_Bounds_SetSize": (self: heap.Ref<object>, retPtr: Pointer, width: number, height: number): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.setSize(width, height);
    },
    "try_Bounds_SetSize": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      width: number,
      height: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.setSize(width, height);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Bounds_SetMinimumSize": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "setMinimumSize" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Bounds_SetMinimumSize": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).setMinimumSize);
    },

    "call_Bounds_SetMinimumSize": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      minWidth: number,
      minHeight: number
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.setMinimumSize(minWidth, minHeight);
    },
    "try_Bounds_SetMinimumSize": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      minWidth: number,
      minHeight: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.setMinimumSize(minWidth, minHeight);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Bounds_SetMaximumSize": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "setMaximumSize" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Bounds_SetMaximumSize": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).setMaximumSize);
    },

    "call_Bounds_SetMaximumSize": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      maxWidth: number,
      maxHeight: number
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.setMaximumSize(maxWidth, maxHeight);
    },
    "try_Bounds_SetMaximumSize": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      maxWidth: number,
      maxHeight: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.setMaximumSize(maxWidth, maxHeight);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_Bounds_Left": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "left" in thiz) {
        const val = thiz.left;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Bounds_Left": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "left", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_Bounds_Top": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "top" in thiz) {
        const val = thiz.top;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Bounds_Top": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "top", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_Bounds_Width": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "width" in thiz) {
        const val = thiz.width;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Bounds_Width": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "width", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_Bounds_Height": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "height" in thiz) {
        const val = thiz.height;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Bounds_Height": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "height", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_Bounds_MinWidth": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "minWidth" in thiz) {
        const val = thiz.minWidth;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Bounds_MinWidth": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "minWidth", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_Bounds_MinHeight": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "minHeight" in thiz) {
        const val = thiz.minHeight;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Bounds_MinHeight": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "minHeight", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_Bounds_MaxWidth": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "maxWidth" in thiz) {
        const val = thiz.maxWidth;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Bounds_MaxWidth": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "maxWidth", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_Bounds_MaxHeight": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "maxHeight" in thiz) {
        const val = thiz.maxHeight;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Bounds_MaxHeight": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "maxHeight", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },

    "has_AppWindow_Focus": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "focus" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_Focus": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).focus);
    },

    "call_AppWindow_Focus": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.focus();
    },
    "try_AppWindow_Focus": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.focus();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_Fullscreen": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "fullscreen" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_Fullscreen": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).fullscreen);
    },

    "call_AppWindow_Fullscreen": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.fullscreen();
    },
    "try_AppWindow_Fullscreen": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.fullscreen();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_IsFullscreen": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "isFullscreen" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_IsFullscreen": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).isFullscreen);
    },

    "call_AppWindow_IsFullscreen": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.isFullscreen();
      A.store.Bool(retPtr, _ret);
    },
    "try_AppWindow_IsFullscreen": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.isFullscreen();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_Minimize": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "minimize" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_Minimize": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).minimize);
    },

    "call_AppWindow_Minimize": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.minimize();
    },
    "try_AppWindow_Minimize": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.minimize();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_IsMinimized": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "isMinimized" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_IsMinimized": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).isMinimized);
    },

    "call_AppWindow_IsMinimized": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.isMinimized();
      A.store.Bool(retPtr, _ret);
    },
    "try_AppWindow_IsMinimized": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.isMinimized();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_Maximize": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "maximize" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_Maximize": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).maximize);
    },

    "call_AppWindow_Maximize": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.maximize();
    },
    "try_AppWindow_Maximize": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.maximize();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_IsMaximized": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "isMaximized" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_IsMaximized": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).isMaximized);
    },

    "call_AppWindow_IsMaximized": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.isMaximized();
      A.store.Bool(retPtr, _ret);
    },
    "try_AppWindow_IsMaximized": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.isMaximized();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_Restore": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "restore" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_Restore": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).restore);
    },

    "call_AppWindow_Restore": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.restore();
    },
    "try_AppWindow_Restore": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.restore();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_MoveTo": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveTo" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_MoveTo": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveTo);
    },

    "call_AppWindow_MoveTo": (self: heap.Ref<object>, retPtr: Pointer, left: number, top: number): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveTo(left, top);
    },
    "try_AppWindow_MoveTo": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      left: number,
      top: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveTo(left, top);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_ResizeTo": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "resizeTo" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_ResizeTo": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).resizeTo);
    },

    "call_AppWindow_ResizeTo": (self: heap.Ref<object>, retPtr: Pointer, width: number, height: number): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.resizeTo(width, height);
    },
    "try_AppWindow_ResizeTo": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      width: number,
      height: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.resizeTo(width, height);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_DrawAttention": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "drawAttention" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_DrawAttention": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).drawAttention);
    },

    "call_AppWindow_DrawAttention": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.drawAttention();
    },
    "try_AppWindow_DrawAttention": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.drawAttention();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_ClearAttention": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "clearAttention" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_ClearAttention": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).clearAttention);
    },

    "call_AppWindow_ClearAttention": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.clearAttention();
    },
    "try_AppWindow_ClearAttention": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.clearAttention();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_Close": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "close" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_Close": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).close);
    },

    "call_AppWindow_Close": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.close();
    },
    "try_AppWindow_Close": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.close();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_Show": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "show" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_Show": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).show);
    },

    "call_AppWindow_Show": (self: heap.Ref<object>, retPtr: Pointer, focused: heap.Ref<boolean>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.show(focused === A.H.TRUE);
    },
    "try_AppWindow_Show": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      focused: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.show(focused === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_Show1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "show" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_Show1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).show);
    },

    "call_AppWindow_Show1": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.show();
    },
    "try_AppWindow_Show1": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.show();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_Hide": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "hide" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_Hide": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).hide);
    },

    "call_AppWindow_Hide": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.hide();
    },
    "try_AppWindow_Hide": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.hide();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_GetBounds": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "getBounds" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_GetBounds": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).getBounds);
    },

    "call_AppWindow_GetBounds": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.getBounds();
      if (typeof _ret === "undefined") {
        A.store.Bool(retPtr + 20, false);
        A.store.Bool(retPtr + 16, false);
        A.store.Int32(retPtr + 0, 0);
        A.store.Bool(retPtr + 17, false);
        A.store.Int32(retPtr + 4, 0);
        A.store.Bool(retPtr + 18, false);
        A.store.Int32(retPtr + 8, 0);
        A.store.Bool(retPtr + 19, false);
        A.store.Int32(retPtr + 12, 0);
      } else {
        A.store.Bool(retPtr + 20, true);
        A.store.Bool(retPtr + 16, "left" in _ret ? true : false);
        A.store.Int32(retPtr + 0, _ret["left"] === undefined ? 0 : (_ret["left"] as number));
        A.store.Bool(retPtr + 17, "top" in _ret ? true : false);
        A.store.Int32(retPtr + 4, _ret["top"] === undefined ? 0 : (_ret["top"] as number));
        A.store.Bool(retPtr + 18, "width" in _ret ? true : false);
        A.store.Int32(retPtr + 8, _ret["width"] === undefined ? 0 : (_ret["width"] as number));
        A.store.Bool(retPtr + 19, "height" in _ret ? true : false);
        A.store.Int32(retPtr + 12, _ret["height"] === undefined ? 0 : (_ret["height"] as number));
      }
    },
    "try_AppWindow_GetBounds": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.getBounds();
        if (typeof _ret === "undefined") {
          A.store.Bool(retPtr + 20, false);
          A.store.Bool(retPtr + 16, false);
          A.store.Int32(retPtr + 0, 0);
          A.store.Bool(retPtr + 17, false);
          A.store.Int32(retPtr + 4, 0);
          A.store.Bool(retPtr + 18, false);
          A.store.Int32(retPtr + 8, 0);
          A.store.Bool(retPtr + 19, false);
          A.store.Int32(retPtr + 12, 0);
        } else {
          A.store.Bool(retPtr + 20, true);
          A.store.Bool(retPtr + 16, "left" in _ret ? true : false);
          A.store.Int32(retPtr + 0, _ret["left"] === undefined ? 0 : (_ret["left"] as number));
          A.store.Bool(retPtr + 17, "top" in _ret ? true : false);
          A.store.Int32(retPtr + 4, _ret["top"] === undefined ? 0 : (_ret["top"] as number));
          A.store.Bool(retPtr + 18, "width" in _ret ? true : false);
          A.store.Int32(retPtr + 8, _ret["width"] === undefined ? 0 : (_ret["width"] as number));
          A.store.Bool(retPtr + 19, "height" in _ret ? true : false);
          A.store.Int32(retPtr + 12, _ret["height"] === undefined ? 0 : (_ret["height"] as number));
        }
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_SetBounds": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "setBounds" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_SetBounds": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).setBounds);
    },

    "call_AppWindow_SetBounds": (self: heap.Ref<object>, retPtr: Pointer, bounds: Pointer): void => {
      const thiz = A.H.get<any>(self);

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
      const _ret = thiz.setBounds(bounds_ffi);
    },
    "try_AppWindow_SetBounds": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      bounds: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

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
        const _ret = thiz.setBounds(bounds_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_SetIcon": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "setIcon" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_SetIcon": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).setIcon);
    },

    "call_AppWindow_SetIcon": (self: heap.Ref<object>, retPtr: Pointer, iconUrl: heap.Ref<object>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.setIcon(A.H.get<object>(iconUrl));
    },
    "try_AppWindow_SetIcon": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      iconUrl: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.setIcon(A.H.get<object>(iconUrl));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_IsAlwaysOnTop": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "isAlwaysOnTop" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_IsAlwaysOnTop": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).isAlwaysOnTop);
    },

    "call_AppWindow_IsAlwaysOnTop": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.isAlwaysOnTop();
      A.store.Bool(retPtr, _ret);
    },
    "try_AppWindow_IsAlwaysOnTop": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.isAlwaysOnTop();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_SetAlwaysOnTop": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "setAlwaysOnTop" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_SetAlwaysOnTop": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).setAlwaysOnTop);
    },

    "call_AppWindow_SetAlwaysOnTop": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      alwaysOnTop: heap.Ref<boolean>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.setAlwaysOnTop(alwaysOnTop === A.H.TRUE);
    },
    "try_AppWindow_SetAlwaysOnTop": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      alwaysOnTop: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.setAlwaysOnTop(alwaysOnTop === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_AlphaEnabled": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "alphaEnabled" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_AlphaEnabled": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).alphaEnabled);
    },

    "call_AppWindow_AlphaEnabled": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.alphaEnabled();
      A.store.Bool(retPtr, _ret);
    },
    "try_AppWindow_AlphaEnabled": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.alphaEnabled();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AppWindow_SetVisibleOnAllWorkspaces": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "setVisibleOnAllWorkspaces" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AppWindow_SetVisibleOnAllWorkspaces": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).setVisibleOnAllWorkspaces);
    },

    "call_AppWindow_SetVisibleOnAllWorkspaces": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      alwaysVisible: heap.Ref<boolean>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.setVisibleOnAllWorkspaces(alwaysVisible === A.H.TRUE);
    },
    "try_AppWindow_SetVisibleOnAllWorkspaces": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      alwaysVisible: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.setVisibleOnAllWorkspaces(alwaysVisible === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_AppWindow_HasFrameColor": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "hasFrameColor" in thiz) {
        const val = thiz.hasFrameColor;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AppWindow_HasFrameColor": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "hasFrameColor", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AppWindow_ActiveFrameColor": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "activeFrameColor" in thiz) {
        const val = thiz.activeFrameColor;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AppWindow_ActiveFrameColor": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "activeFrameColor", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AppWindow_InactiveFrameColor": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "inactiveFrameColor" in thiz) {
        const val = thiz.inactiveFrameColor;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AppWindow_InactiveFrameColor": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "inactiveFrameColor", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AppWindow_ContentWindow": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "contentWindow" in thiz) {
        const val = thiz.contentWindow;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AppWindow_ContentWindow": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "contentWindow", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AppWindow_Id": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "id" in thiz) {
        const val = thiz.id;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AppWindow_Id": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "id", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AppWindow_InnerBounds": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "innerBounds" in thiz) {
        const val = thiz.innerBounds;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AppWindow_InnerBounds": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "innerBounds", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AppWindow_OuterBounds": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "outerBounds" in thiz) {
        const val = thiz.outerBounds;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AppWindow_OuterBounds": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "outerBounds", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },

    "store_BoundsSpecification": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 40, false);
        A.store.Bool(ptr + 32, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 33, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 34, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 35, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 36, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Bool(ptr + 37, false);
        A.store.Int32(ptr + 20, 0);
        A.store.Bool(ptr + 38, false);
        A.store.Int32(ptr + 24, 0);
        A.store.Bool(ptr + 39, false);
        A.store.Int32(ptr + 28, 0);
      } else {
        A.store.Bool(ptr + 40, true);
        A.store.Bool(ptr + 32, "left" in x ? true : false);
        A.store.Int32(ptr + 0, x["left"] === undefined ? 0 : (x["left"] as number));
        A.store.Bool(ptr + 33, "top" in x ? true : false);
        A.store.Int32(ptr + 4, x["top"] === undefined ? 0 : (x["top"] as number));
        A.store.Bool(ptr + 34, "width" in x ? true : false);
        A.store.Int32(ptr + 8, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Bool(ptr + 35, "height" in x ? true : false);
        A.store.Int32(ptr + 12, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Bool(ptr + 36, "minWidth" in x ? true : false);
        A.store.Int32(ptr + 16, x["minWidth"] === undefined ? 0 : (x["minWidth"] as number));
        A.store.Bool(ptr + 37, "minHeight" in x ? true : false);
        A.store.Int32(ptr + 20, x["minHeight"] === undefined ? 0 : (x["minHeight"] as number));
        A.store.Bool(ptr + 38, "maxWidth" in x ? true : false);
        A.store.Int32(ptr + 24, x["maxWidth"] === undefined ? 0 : (x["maxWidth"] as number));
        A.store.Bool(ptr + 39, "maxHeight" in x ? true : false);
        A.store.Int32(ptr + 28, x["maxHeight"] === undefined ? 0 : (x["maxHeight"] as number));
      }
    },
    "load_BoundsSpecification": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 32)) {
        x["left"] = A.load.Int32(ptr + 0);
      } else {
        delete x["left"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["top"] = A.load.Int32(ptr + 4);
      } else {
        delete x["top"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["width"] = A.load.Int32(ptr + 8);
      } else {
        delete x["width"];
      }
      if (A.load.Bool(ptr + 35)) {
        x["height"] = A.load.Int32(ptr + 12);
      } else {
        delete x["height"];
      }
      if (A.load.Bool(ptr + 36)) {
        x["minWidth"] = A.load.Int32(ptr + 16);
      } else {
        delete x["minWidth"];
      }
      if (A.load.Bool(ptr + 37)) {
        x["minHeight"] = A.load.Int32(ptr + 20);
      } else {
        delete x["minHeight"];
      }
      if (A.load.Bool(ptr + 38)) {
        x["maxWidth"] = A.load.Int32(ptr + 24);
      } else {
        delete x["maxWidth"];
      }
      if (A.load.Bool(ptr + 39)) {
        x["maxHeight"] = A.load.Int32(ptr + 28);
      } else {
        delete x["maxHeight"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_WindowType": (ref: heap.Ref<string>): number => {
      const idx = ["shell", "panel"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_FrameOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["type"]);
        A.store.Ref(ptr + 4, x["color"]);
        A.store.Ref(ptr + 8, x["activeColor"]);
        A.store.Ref(ptr + 12, x["inactiveColor"]);
      }
    },
    "load_FrameOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Ref(ptr + 0, undefined);
      x["color"] = A.load.Ref(ptr + 4, undefined);
      x["activeColor"] = A.load.Ref(ptr + 8, undefined);
      x["inactiveColor"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_State": (ref: heap.Ref<string>): number => {
      const idx = ["normal", "fullscreen", "maximized", "minimized"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_CreateWindowOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 217, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 4 + 40, false);
        A.store.Bool(ptr + 4 + 32, false);
        A.store.Int32(ptr + 4 + 0, 0);
        A.store.Bool(ptr + 4 + 33, false);
        A.store.Int32(ptr + 4 + 4, 0);
        A.store.Bool(ptr + 4 + 34, false);
        A.store.Int32(ptr + 4 + 8, 0);
        A.store.Bool(ptr + 4 + 35, false);
        A.store.Int32(ptr + 4 + 12, 0);
        A.store.Bool(ptr + 4 + 36, false);
        A.store.Int32(ptr + 4 + 16, 0);
        A.store.Bool(ptr + 4 + 37, false);
        A.store.Int32(ptr + 4 + 20, 0);
        A.store.Bool(ptr + 4 + 38, false);
        A.store.Int32(ptr + 4 + 24, 0);
        A.store.Bool(ptr + 4 + 39, false);
        A.store.Int32(ptr + 4 + 28, 0);

        A.store.Bool(ptr + 48 + 40, false);
        A.store.Bool(ptr + 48 + 32, false);
        A.store.Int32(ptr + 48 + 0, 0);
        A.store.Bool(ptr + 48 + 33, false);
        A.store.Int32(ptr + 48 + 4, 0);
        A.store.Bool(ptr + 48 + 34, false);
        A.store.Int32(ptr + 48 + 8, 0);
        A.store.Bool(ptr + 48 + 35, false);
        A.store.Int32(ptr + 48 + 12, 0);
        A.store.Bool(ptr + 48 + 36, false);
        A.store.Int32(ptr + 48 + 16, 0);
        A.store.Bool(ptr + 48 + 37, false);
        A.store.Int32(ptr + 48 + 20, 0);
        A.store.Bool(ptr + 48 + 38, false);
        A.store.Int32(ptr + 48 + 24, 0);
        A.store.Bool(ptr + 48 + 39, false);
        A.store.Int32(ptr + 48 + 28, 0);
        A.store.Bool(ptr + 196, false);
        A.store.Int32(ptr + 92, 0);
        A.store.Bool(ptr + 197, false);
        A.store.Int32(ptr + 96, 0);
        A.store.Bool(ptr + 198, false);
        A.store.Int32(ptr + 100, 0);
        A.store.Bool(ptr + 199, false);
        A.store.Int32(ptr + 104, 0);
        A.store.Bool(ptr + 200, false);
        A.store.Int32(ptr + 108, 0);
        A.store.Bool(ptr + 201, false);
        A.store.Int32(ptr + 112, 0);
        A.store.Bool(ptr + 202, false);
        A.store.Int32(ptr + 116, 0);
        A.store.Bool(ptr + 203, false);
        A.store.Int32(ptr + 120, 0);
        A.store.Bool(ptr + 204, false);
        A.store.Int32(ptr + 124, 0);
        A.store.Bool(ptr + 205, false);
        A.store.Int32(ptr + 128, 0);
        A.store.Bool(ptr + 206, false);
        A.store.Int32(ptr + 132, 0);
        A.store.Bool(ptr + 207, false);
        A.store.Int32(ptr + 136, 0);
        A.store.Enum(ptr + 140, -1);
        A.store.Bool(ptr + 208, false);
        A.store.Bool(ptr + 144, false);
        A.store.Bool(ptr + 209, false);
        A.store.Bool(ptr + 145, false);
        A.store.Ref(ptr + 148, undefined);
        A.store.Ref(ptr + 152, undefined);

        A.store.Bool(ptr + 156 + 20, false);
        A.store.Bool(ptr + 156 + 16, false);
        A.store.Int32(ptr + 156 + 0, 0);
        A.store.Bool(ptr + 156 + 17, false);
        A.store.Int32(ptr + 156 + 4, 0);
        A.store.Bool(ptr + 156 + 18, false);
        A.store.Int32(ptr + 156 + 8, 0);
        A.store.Bool(ptr + 156 + 19, false);
        A.store.Int32(ptr + 156 + 12, 0);
        A.store.Bool(ptr + 210, false);
        A.store.Bool(ptr + 177, false);
        A.store.Enum(ptr + 180, -1);
        A.store.Bool(ptr + 211, false);
        A.store.Bool(ptr + 184, false);
        A.store.Bool(ptr + 212, false);
        A.store.Bool(ptr + 185, false);
        A.store.Bool(ptr + 213, false);
        A.store.Bool(ptr + 186, false);
        A.store.Bool(ptr + 214, false);
        A.store.Bool(ptr + 187, false);
        A.store.Bool(ptr + 215, false);
        A.store.Bool(ptr + 188, false);
        A.store.Bool(ptr + 216, false);
        A.store.Bool(ptr + 189, false);
        A.store.Enum(ptr + 192, -1);
      } else {
        A.store.Bool(ptr + 217, true);
        A.store.Ref(ptr + 0, x["id"]);

        if (typeof x["innerBounds"] === "undefined") {
          A.store.Bool(ptr + 4 + 40, false);
          A.store.Bool(ptr + 4 + 32, false);
          A.store.Int32(ptr + 4 + 0, 0);
          A.store.Bool(ptr + 4 + 33, false);
          A.store.Int32(ptr + 4 + 4, 0);
          A.store.Bool(ptr + 4 + 34, false);
          A.store.Int32(ptr + 4 + 8, 0);
          A.store.Bool(ptr + 4 + 35, false);
          A.store.Int32(ptr + 4 + 12, 0);
          A.store.Bool(ptr + 4 + 36, false);
          A.store.Int32(ptr + 4 + 16, 0);
          A.store.Bool(ptr + 4 + 37, false);
          A.store.Int32(ptr + 4 + 20, 0);
          A.store.Bool(ptr + 4 + 38, false);
          A.store.Int32(ptr + 4 + 24, 0);
          A.store.Bool(ptr + 4 + 39, false);
          A.store.Int32(ptr + 4 + 28, 0);
        } else {
          A.store.Bool(ptr + 4 + 40, true);
          A.store.Bool(ptr + 4 + 32, "left" in x["innerBounds"] ? true : false);
          A.store.Int32(ptr + 4 + 0, x["innerBounds"]["left"] === undefined ? 0 : (x["innerBounds"]["left"] as number));
          A.store.Bool(ptr + 4 + 33, "top" in x["innerBounds"] ? true : false);
          A.store.Int32(ptr + 4 + 4, x["innerBounds"]["top"] === undefined ? 0 : (x["innerBounds"]["top"] as number));
          A.store.Bool(ptr + 4 + 34, "width" in x["innerBounds"] ? true : false);
          A.store.Int32(
            ptr + 4 + 8,
            x["innerBounds"]["width"] === undefined ? 0 : (x["innerBounds"]["width"] as number)
          );
          A.store.Bool(ptr + 4 + 35, "height" in x["innerBounds"] ? true : false);
          A.store.Int32(
            ptr + 4 + 12,
            x["innerBounds"]["height"] === undefined ? 0 : (x["innerBounds"]["height"] as number)
          );
          A.store.Bool(ptr + 4 + 36, "minWidth" in x["innerBounds"] ? true : false);
          A.store.Int32(
            ptr + 4 + 16,
            x["innerBounds"]["minWidth"] === undefined ? 0 : (x["innerBounds"]["minWidth"] as number)
          );
          A.store.Bool(ptr + 4 + 37, "minHeight" in x["innerBounds"] ? true : false);
          A.store.Int32(
            ptr + 4 + 20,
            x["innerBounds"]["minHeight"] === undefined ? 0 : (x["innerBounds"]["minHeight"] as number)
          );
          A.store.Bool(ptr + 4 + 38, "maxWidth" in x["innerBounds"] ? true : false);
          A.store.Int32(
            ptr + 4 + 24,
            x["innerBounds"]["maxWidth"] === undefined ? 0 : (x["innerBounds"]["maxWidth"] as number)
          );
          A.store.Bool(ptr + 4 + 39, "maxHeight" in x["innerBounds"] ? true : false);
          A.store.Int32(
            ptr + 4 + 28,
            x["innerBounds"]["maxHeight"] === undefined ? 0 : (x["innerBounds"]["maxHeight"] as number)
          );
        }

        if (typeof x["outerBounds"] === "undefined") {
          A.store.Bool(ptr + 48 + 40, false);
          A.store.Bool(ptr + 48 + 32, false);
          A.store.Int32(ptr + 48 + 0, 0);
          A.store.Bool(ptr + 48 + 33, false);
          A.store.Int32(ptr + 48 + 4, 0);
          A.store.Bool(ptr + 48 + 34, false);
          A.store.Int32(ptr + 48 + 8, 0);
          A.store.Bool(ptr + 48 + 35, false);
          A.store.Int32(ptr + 48 + 12, 0);
          A.store.Bool(ptr + 48 + 36, false);
          A.store.Int32(ptr + 48 + 16, 0);
          A.store.Bool(ptr + 48 + 37, false);
          A.store.Int32(ptr + 48 + 20, 0);
          A.store.Bool(ptr + 48 + 38, false);
          A.store.Int32(ptr + 48 + 24, 0);
          A.store.Bool(ptr + 48 + 39, false);
          A.store.Int32(ptr + 48 + 28, 0);
        } else {
          A.store.Bool(ptr + 48 + 40, true);
          A.store.Bool(ptr + 48 + 32, "left" in x["outerBounds"] ? true : false);
          A.store.Int32(
            ptr + 48 + 0,
            x["outerBounds"]["left"] === undefined ? 0 : (x["outerBounds"]["left"] as number)
          );
          A.store.Bool(ptr + 48 + 33, "top" in x["outerBounds"] ? true : false);
          A.store.Int32(ptr + 48 + 4, x["outerBounds"]["top"] === undefined ? 0 : (x["outerBounds"]["top"] as number));
          A.store.Bool(ptr + 48 + 34, "width" in x["outerBounds"] ? true : false);
          A.store.Int32(
            ptr + 48 + 8,
            x["outerBounds"]["width"] === undefined ? 0 : (x["outerBounds"]["width"] as number)
          );
          A.store.Bool(ptr + 48 + 35, "height" in x["outerBounds"] ? true : false);
          A.store.Int32(
            ptr + 48 + 12,
            x["outerBounds"]["height"] === undefined ? 0 : (x["outerBounds"]["height"] as number)
          );
          A.store.Bool(ptr + 48 + 36, "minWidth" in x["outerBounds"] ? true : false);
          A.store.Int32(
            ptr + 48 + 16,
            x["outerBounds"]["minWidth"] === undefined ? 0 : (x["outerBounds"]["minWidth"] as number)
          );
          A.store.Bool(ptr + 48 + 37, "minHeight" in x["outerBounds"] ? true : false);
          A.store.Int32(
            ptr + 48 + 20,
            x["outerBounds"]["minHeight"] === undefined ? 0 : (x["outerBounds"]["minHeight"] as number)
          );
          A.store.Bool(ptr + 48 + 38, "maxWidth" in x["outerBounds"] ? true : false);
          A.store.Int32(
            ptr + 48 + 24,
            x["outerBounds"]["maxWidth"] === undefined ? 0 : (x["outerBounds"]["maxWidth"] as number)
          );
          A.store.Bool(ptr + 48 + 39, "maxHeight" in x["outerBounds"] ? true : false);
          A.store.Int32(
            ptr + 48 + 28,
            x["outerBounds"]["maxHeight"] === undefined ? 0 : (x["outerBounds"]["maxHeight"] as number)
          );
        }
        A.store.Bool(ptr + 196, "defaultWidth" in x ? true : false);
        A.store.Int32(ptr + 92, x["defaultWidth"] === undefined ? 0 : (x["defaultWidth"] as number));
        A.store.Bool(ptr + 197, "defaultHeight" in x ? true : false);
        A.store.Int32(ptr + 96, x["defaultHeight"] === undefined ? 0 : (x["defaultHeight"] as number));
        A.store.Bool(ptr + 198, "defaultLeft" in x ? true : false);
        A.store.Int32(ptr + 100, x["defaultLeft"] === undefined ? 0 : (x["defaultLeft"] as number));
        A.store.Bool(ptr + 199, "defaultTop" in x ? true : false);
        A.store.Int32(ptr + 104, x["defaultTop"] === undefined ? 0 : (x["defaultTop"] as number));
        A.store.Bool(ptr + 200, "width" in x ? true : false);
        A.store.Int32(ptr + 108, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Bool(ptr + 201, "height" in x ? true : false);
        A.store.Int32(ptr + 112, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Bool(ptr + 202, "left" in x ? true : false);
        A.store.Int32(ptr + 116, x["left"] === undefined ? 0 : (x["left"] as number));
        A.store.Bool(ptr + 203, "top" in x ? true : false);
        A.store.Int32(ptr + 120, x["top"] === undefined ? 0 : (x["top"] as number));
        A.store.Bool(ptr + 204, "minWidth" in x ? true : false);
        A.store.Int32(ptr + 124, x["minWidth"] === undefined ? 0 : (x["minWidth"] as number));
        A.store.Bool(ptr + 205, "minHeight" in x ? true : false);
        A.store.Int32(ptr + 128, x["minHeight"] === undefined ? 0 : (x["minHeight"] as number));
        A.store.Bool(ptr + 206, "maxWidth" in x ? true : false);
        A.store.Int32(ptr + 132, x["maxWidth"] === undefined ? 0 : (x["maxWidth"] as number));
        A.store.Bool(ptr + 207, "maxHeight" in x ? true : false);
        A.store.Int32(ptr + 136, x["maxHeight"] === undefined ? 0 : (x["maxHeight"] as number));
        A.store.Enum(ptr + 140, ["shell", "panel"].indexOf(x["type"] as string));
        A.store.Bool(ptr + 208, "ime" in x ? true : false);
        A.store.Bool(ptr + 144, x["ime"] ? true : false);
        A.store.Bool(ptr + 209, "showInShelf" in x ? true : false);
        A.store.Bool(ptr + 145, x["showInShelf"] ? true : false);
        A.store.Ref(ptr + 148, x["icon"]);
        A.store.Ref(ptr + 152, x["frame"]);

        if (typeof x["bounds"] === "undefined") {
          A.store.Bool(ptr + 156 + 20, false);
          A.store.Bool(ptr + 156 + 16, false);
          A.store.Int32(ptr + 156 + 0, 0);
          A.store.Bool(ptr + 156 + 17, false);
          A.store.Int32(ptr + 156 + 4, 0);
          A.store.Bool(ptr + 156 + 18, false);
          A.store.Int32(ptr + 156 + 8, 0);
          A.store.Bool(ptr + 156 + 19, false);
          A.store.Int32(ptr + 156 + 12, 0);
        } else {
          A.store.Bool(ptr + 156 + 20, true);
          A.store.Bool(ptr + 156 + 16, "left" in x["bounds"] ? true : false);
          A.store.Int32(ptr + 156 + 0, x["bounds"]["left"] === undefined ? 0 : (x["bounds"]["left"] as number));
          A.store.Bool(ptr + 156 + 17, "top" in x["bounds"] ? true : false);
          A.store.Int32(ptr + 156 + 4, x["bounds"]["top"] === undefined ? 0 : (x["bounds"]["top"] as number));
          A.store.Bool(ptr + 156 + 18, "width" in x["bounds"] ? true : false);
          A.store.Int32(ptr + 156 + 8, x["bounds"]["width"] === undefined ? 0 : (x["bounds"]["width"] as number));
          A.store.Bool(ptr + 156 + 19, "height" in x["bounds"] ? true : false);
          A.store.Int32(ptr + 156 + 12, x["bounds"]["height"] === undefined ? 0 : (x["bounds"]["height"] as number));
        }
        A.store.Bool(ptr + 210, "alphaEnabled" in x ? true : false);
        A.store.Bool(ptr + 177, x["alphaEnabled"] ? true : false);
        A.store.Enum(ptr + 180, ["normal", "fullscreen", "maximized", "minimized"].indexOf(x["state"] as string));
        A.store.Bool(ptr + 211, "hidden" in x ? true : false);
        A.store.Bool(ptr + 184, x["hidden"] ? true : false);
        A.store.Bool(ptr + 212, "resizable" in x ? true : false);
        A.store.Bool(ptr + 185, x["resizable"] ? true : false);
        A.store.Bool(ptr + 213, "singleton" in x ? true : false);
        A.store.Bool(ptr + 186, x["singleton"] ? true : false);
        A.store.Bool(ptr + 214, "alwaysOnTop" in x ? true : false);
        A.store.Bool(ptr + 187, x["alwaysOnTop"] ? true : false);
        A.store.Bool(ptr + 215, "focused" in x ? true : false);
        A.store.Bool(ptr + 188, x["focused"] ? true : false);
        A.store.Bool(ptr + 216, "visibleOnAllWorkspaces" in x ? true : false);
        A.store.Bool(ptr + 189, x["visibleOnAllWorkspaces"] ? true : false);
        A.store.Enum(ptr + 192, ["new_note"].indexOf(x["lockScreenAction"] as string));
      }
    },
    "load_CreateWindowOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 4 + 40)) {
        x["innerBounds"] = {};
        if (A.load.Bool(ptr + 4 + 32)) {
          x["innerBounds"]["left"] = A.load.Int32(ptr + 4 + 0);
        } else {
          delete x["innerBounds"]["left"];
        }
        if (A.load.Bool(ptr + 4 + 33)) {
          x["innerBounds"]["top"] = A.load.Int32(ptr + 4 + 4);
        } else {
          delete x["innerBounds"]["top"];
        }
        if (A.load.Bool(ptr + 4 + 34)) {
          x["innerBounds"]["width"] = A.load.Int32(ptr + 4 + 8);
        } else {
          delete x["innerBounds"]["width"];
        }
        if (A.load.Bool(ptr + 4 + 35)) {
          x["innerBounds"]["height"] = A.load.Int32(ptr + 4 + 12);
        } else {
          delete x["innerBounds"]["height"];
        }
        if (A.load.Bool(ptr + 4 + 36)) {
          x["innerBounds"]["minWidth"] = A.load.Int32(ptr + 4 + 16);
        } else {
          delete x["innerBounds"]["minWidth"];
        }
        if (A.load.Bool(ptr + 4 + 37)) {
          x["innerBounds"]["minHeight"] = A.load.Int32(ptr + 4 + 20);
        } else {
          delete x["innerBounds"]["minHeight"];
        }
        if (A.load.Bool(ptr + 4 + 38)) {
          x["innerBounds"]["maxWidth"] = A.load.Int32(ptr + 4 + 24);
        } else {
          delete x["innerBounds"]["maxWidth"];
        }
        if (A.load.Bool(ptr + 4 + 39)) {
          x["innerBounds"]["maxHeight"] = A.load.Int32(ptr + 4 + 28);
        } else {
          delete x["innerBounds"]["maxHeight"];
        }
      } else {
        delete x["innerBounds"];
      }
      if (A.load.Bool(ptr + 48 + 40)) {
        x["outerBounds"] = {};
        if (A.load.Bool(ptr + 48 + 32)) {
          x["outerBounds"]["left"] = A.load.Int32(ptr + 48 + 0);
        } else {
          delete x["outerBounds"]["left"];
        }
        if (A.load.Bool(ptr + 48 + 33)) {
          x["outerBounds"]["top"] = A.load.Int32(ptr + 48 + 4);
        } else {
          delete x["outerBounds"]["top"];
        }
        if (A.load.Bool(ptr + 48 + 34)) {
          x["outerBounds"]["width"] = A.load.Int32(ptr + 48 + 8);
        } else {
          delete x["outerBounds"]["width"];
        }
        if (A.load.Bool(ptr + 48 + 35)) {
          x["outerBounds"]["height"] = A.load.Int32(ptr + 48 + 12);
        } else {
          delete x["outerBounds"]["height"];
        }
        if (A.load.Bool(ptr + 48 + 36)) {
          x["outerBounds"]["minWidth"] = A.load.Int32(ptr + 48 + 16);
        } else {
          delete x["outerBounds"]["minWidth"];
        }
        if (A.load.Bool(ptr + 48 + 37)) {
          x["outerBounds"]["minHeight"] = A.load.Int32(ptr + 48 + 20);
        } else {
          delete x["outerBounds"]["minHeight"];
        }
        if (A.load.Bool(ptr + 48 + 38)) {
          x["outerBounds"]["maxWidth"] = A.load.Int32(ptr + 48 + 24);
        } else {
          delete x["outerBounds"]["maxWidth"];
        }
        if (A.load.Bool(ptr + 48 + 39)) {
          x["outerBounds"]["maxHeight"] = A.load.Int32(ptr + 48 + 28);
        } else {
          delete x["outerBounds"]["maxHeight"];
        }
      } else {
        delete x["outerBounds"];
      }
      if (A.load.Bool(ptr + 196)) {
        x["defaultWidth"] = A.load.Int32(ptr + 92);
      } else {
        delete x["defaultWidth"];
      }
      if (A.load.Bool(ptr + 197)) {
        x["defaultHeight"] = A.load.Int32(ptr + 96);
      } else {
        delete x["defaultHeight"];
      }
      if (A.load.Bool(ptr + 198)) {
        x["defaultLeft"] = A.load.Int32(ptr + 100);
      } else {
        delete x["defaultLeft"];
      }
      if (A.load.Bool(ptr + 199)) {
        x["defaultTop"] = A.load.Int32(ptr + 104);
      } else {
        delete x["defaultTop"];
      }
      if (A.load.Bool(ptr + 200)) {
        x["width"] = A.load.Int32(ptr + 108);
      } else {
        delete x["width"];
      }
      if (A.load.Bool(ptr + 201)) {
        x["height"] = A.load.Int32(ptr + 112);
      } else {
        delete x["height"];
      }
      if (A.load.Bool(ptr + 202)) {
        x["left"] = A.load.Int32(ptr + 116);
      } else {
        delete x["left"];
      }
      if (A.load.Bool(ptr + 203)) {
        x["top"] = A.load.Int32(ptr + 120);
      } else {
        delete x["top"];
      }
      if (A.load.Bool(ptr + 204)) {
        x["minWidth"] = A.load.Int32(ptr + 124);
      } else {
        delete x["minWidth"];
      }
      if (A.load.Bool(ptr + 205)) {
        x["minHeight"] = A.load.Int32(ptr + 128);
      } else {
        delete x["minHeight"];
      }
      if (A.load.Bool(ptr + 206)) {
        x["maxWidth"] = A.load.Int32(ptr + 132);
      } else {
        delete x["maxWidth"];
      }
      if (A.load.Bool(ptr + 207)) {
        x["maxHeight"] = A.load.Int32(ptr + 136);
      } else {
        delete x["maxHeight"];
      }
      x["type"] = A.load.Enum(ptr + 140, ["shell", "panel"]);
      if (A.load.Bool(ptr + 208)) {
        x["ime"] = A.load.Bool(ptr + 144);
      } else {
        delete x["ime"];
      }
      if (A.load.Bool(ptr + 209)) {
        x["showInShelf"] = A.load.Bool(ptr + 145);
      } else {
        delete x["showInShelf"];
      }
      x["icon"] = A.load.Ref(ptr + 148, undefined);
      x["frame"] = A.load.Ref(ptr + 152, undefined);
      if (A.load.Bool(ptr + 156 + 20)) {
        x["bounds"] = {};
        if (A.load.Bool(ptr + 156 + 16)) {
          x["bounds"]["left"] = A.load.Int32(ptr + 156 + 0);
        } else {
          delete x["bounds"]["left"];
        }
        if (A.load.Bool(ptr + 156 + 17)) {
          x["bounds"]["top"] = A.load.Int32(ptr + 156 + 4);
        } else {
          delete x["bounds"]["top"];
        }
        if (A.load.Bool(ptr + 156 + 18)) {
          x["bounds"]["width"] = A.load.Int32(ptr + 156 + 8);
        } else {
          delete x["bounds"]["width"];
        }
        if (A.load.Bool(ptr + 156 + 19)) {
          x["bounds"]["height"] = A.load.Int32(ptr + 156 + 12);
        } else {
          delete x["bounds"]["height"];
        }
      } else {
        delete x["bounds"];
      }
      if (A.load.Bool(ptr + 210)) {
        x["alphaEnabled"] = A.load.Bool(ptr + 177);
      } else {
        delete x["alphaEnabled"];
      }
      x["state"] = A.load.Enum(ptr + 180, ["normal", "fullscreen", "maximized", "minimized"]);
      if (A.load.Bool(ptr + 211)) {
        x["hidden"] = A.load.Bool(ptr + 184);
      } else {
        delete x["hidden"];
      }
      if (A.load.Bool(ptr + 212)) {
        x["resizable"] = A.load.Bool(ptr + 185);
      } else {
        delete x["resizable"];
      }
      if (A.load.Bool(ptr + 213)) {
        x["singleton"] = A.load.Bool(ptr + 186);
      } else {
        delete x["singleton"];
      }
      if (A.load.Bool(ptr + 214)) {
        x["alwaysOnTop"] = A.load.Bool(ptr + 187);
      } else {
        delete x["alwaysOnTop"];
      }
      if (A.load.Bool(ptr + 215)) {
        x["focused"] = A.load.Bool(ptr + 188);
      } else {
        delete x["focused"];
      }
      if (A.load.Bool(ptr + 216)) {
        x["visibleOnAllWorkspaces"] = A.load.Bool(ptr + 189);
      } else {
        delete x["visibleOnAllWorkspaces"];
      }
      x["lockScreenAction"] = A.load.Enum(ptr + 192, ["new_note"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_CanSetVisibleOnAllWorkspaces": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window && "canSetVisibleOnAllWorkspaces" in WEBEXT?.app?.window) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CanSetVisibleOnAllWorkspaces": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.canSetVisibleOnAllWorkspaces);
    },
    "call_CanSetVisibleOnAllWorkspaces": (retPtr: Pointer): void => {
      const _ret = WEBEXT.app.window.canSetVisibleOnAllWorkspaces();
      A.store.Bool(retPtr, _ret);
    },
    "try_CanSetVisibleOnAllWorkspaces": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.canSetVisibleOnAllWorkspaces();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Create": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window && "create" in WEBEXT?.app?.window) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Create": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.create);
    },
    "call_Create": (retPtr: Pointer, url: heap.Ref<object>, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["id"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 40)) {
        options_ffi["innerBounds"] = {};
        if (A.load.Bool(options + 4 + 32)) {
          options_ffi["innerBounds"]["left"] = A.load.Int32(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 33)) {
          options_ffi["innerBounds"]["top"] = A.load.Int32(options + 4 + 4);
        }
        if (A.load.Bool(options + 4 + 34)) {
          options_ffi["innerBounds"]["width"] = A.load.Int32(options + 4 + 8);
        }
        if (A.load.Bool(options + 4 + 35)) {
          options_ffi["innerBounds"]["height"] = A.load.Int32(options + 4 + 12);
        }
        if (A.load.Bool(options + 4 + 36)) {
          options_ffi["innerBounds"]["minWidth"] = A.load.Int32(options + 4 + 16);
        }
        if (A.load.Bool(options + 4 + 37)) {
          options_ffi["innerBounds"]["minHeight"] = A.load.Int32(options + 4 + 20);
        }
        if (A.load.Bool(options + 4 + 38)) {
          options_ffi["innerBounds"]["maxWidth"] = A.load.Int32(options + 4 + 24);
        }
        if (A.load.Bool(options + 4 + 39)) {
          options_ffi["innerBounds"]["maxHeight"] = A.load.Int32(options + 4 + 28);
        }
      }
      if (A.load.Bool(options + 48 + 40)) {
        options_ffi["outerBounds"] = {};
        if (A.load.Bool(options + 48 + 32)) {
          options_ffi["outerBounds"]["left"] = A.load.Int32(options + 48 + 0);
        }
        if (A.load.Bool(options + 48 + 33)) {
          options_ffi["outerBounds"]["top"] = A.load.Int32(options + 48 + 4);
        }
        if (A.load.Bool(options + 48 + 34)) {
          options_ffi["outerBounds"]["width"] = A.load.Int32(options + 48 + 8);
        }
        if (A.load.Bool(options + 48 + 35)) {
          options_ffi["outerBounds"]["height"] = A.load.Int32(options + 48 + 12);
        }
        if (A.load.Bool(options + 48 + 36)) {
          options_ffi["outerBounds"]["minWidth"] = A.load.Int32(options + 48 + 16);
        }
        if (A.load.Bool(options + 48 + 37)) {
          options_ffi["outerBounds"]["minHeight"] = A.load.Int32(options + 48 + 20);
        }
        if (A.load.Bool(options + 48 + 38)) {
          options_ffi["outerBounds"]["maxWidth"] = A.load.Int32(options + 48 + 24);
        }
        if (A.load.Bool(options + 48 + 39)) {
          options_ffi["outerBounds"]["maxHeight"] = A.load.Int32(options + 48 + 28);
        }
      }
      if (A.load.Bool(options + 196)) {
        options_ffi["defaultWidth"] = A.load.Int32(options + 92);
      }
      if (A.load.Bool(options + 197)) {
        options_ffi["defaultHeight"] = A.load.Int32(options + 96);
      }
      if (A.load.Bool(options + 198)) {
        options_ffi["defaultLeft"] = A.load.Int32(options + 100);
      }
      if (A.load.Bool(options + 199)) {
        options_ffi["defaultTop"] = A.load.Int32(options + 104);
      }
      if (A.load.Bool(options + 200)) {
        options_ffi["width"] = A.load.Int32(options + 108);
      }
      if (A.load.Bool(options + 201)) {
        options_ffi["height"] = A.load.Int32(options + 112);
      }
      if (A.load.Bool(options + 202)) {
        options_ffi["left"] = A.load.Int32(options + 116);
      }
      if (A.load.Bool(options + 203)) {
        options_ffi["top"] = A.load.Int32(options + 120);
      }
      if (A.load.Bool(options + 204)) {
        options_ffi["minWidth"] = A.load.Int32(options + 124);
      }
      if (A.load.Bool(options + 205)) {
        options_ffi["minHeight"] = A.load.Int32(options + 128);
      }
      if (A.load.Bool(options + 206)) {
        options_ffi["maxWidth"] = A.load.Int32(options + 132);
      }
      if (A.load.Bool(options + 207)) {
        options_ffi["maxHeight"] = A.load.Int32(options + 136);
      }
      options_ffi["type"] = A.load.Enum(options + 140, ["shell", "panel"]);
      if (A.load.Bool(options + 208)) {
        options_ffi["ime"] = A.load.Bool(options + 144);
      }
      if (A.load.Bool(options + 209)) {
        options_ffi["showInShelf"] = A.load.Bool(options + 145);
      }
      options_ffi["icon"] = A.load.Ref(options + 148, undefined);
      options_ffi["frame"] = A.load.Ref(options + 152, undefined);
      if (A.load.Bool(options + 156 + 20)) {
        options_ffi["bounds"] = {};
        if (A.load.Bool(options + 156 + 16)) {
          options_ffi["bounds"]["left"] = A.load.Int32(options + 156 + 0);
        }
        if (A.load.Bool(options + 156 + 17)) {
          options_ffi["bounds"]["top"] = A.load.Int32(options + 156 + 4);
        }
        if (A.load.Bool(options + 156 + 18)) {
          options_ffi["bounds"]["width"] = A.load.Int32(options + 156 + 8);
        }
        if (A.load.Bool(options + 156 + 19)) {
          options_ffi["bounds"]["height"] = A.load.Int32(options + 156 + 12);
        }
      }
      if (A.load.Bool(options + 210)) {
        options_ffi["alphaEnabled"] = A.load.Bool(options + 177);
      }
      options_ffi["state"] = A.load.Enum(options + 180, ["normal", "fullscreen", "maximized", "minimized"]);
      if (A.load.Bool(options + 211)) {
        options_ffi["hidden"] = A.load.Bool(options + 184);
      }
      if (A.load.Bool(options + 212)) {
        options_ffi["resizable"] = A.load.Bool(options + 185);
      }
      if (A.load.Bool(options + 213)) {
        options_ffi["singleton"] = A.load.Bool(options + 186);
      }
      if (A.load.Bool(options + 214)) {
        options_ffi["alwaysOnTop"] = A.load.Bool(options + 187);
      }
      if (A.load.Bool(options + 215)) {
        options_ffi["focused"] = A.load.Bool(options + 188);
      }
      if (A.load.Bool(options + 216)) {
        options_ffi["visibleOnAllWorkspaces"] = A.load.Bool(options + 189);
      }
      options_ffi["lockScreenAction"] = A.load.Enum(options + 192, ["new_note"]);

      const _ret = WEBEXT.app.window.create(A.H.get<object>(url), options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Create": (retPtr: Pointer, errPtr: Pointer, url: heap.Ref<object>, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["id"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 40)) {
          options_ffi["innerBounds"] = {};
          if (A.load.Bool(options + 4 + 32)) {
            options_ffi["innerBounds"]["left"] = A.load.Int32(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 33)) {
            options_ffi["innerBounds"]["top"] = A.load.Int32(options + 4 + 4);
          }
          if (A.load.Bool(options + 4 + 34)) {
            options_ffi["innerBounds"]["width"] = A.load.Int32(options + 4 + 8);
          }
          if (A.load.Bool(options + 4 + 35)) {
            options_ffi["innerBounds"]["height"] = A.load.Int32(options + 4 + 12);
          }
          if (A.load.Bool(options + 4 + 36)) {
            options_ffi["innerBounds"]["minWidth"] = A.load.Int32(options + 4 + 16);
          }
          if (A.load.Bool(options + 4 + 37)) {
            options_ffi["innerBounds"]["minHeight"] = A.load.Int32(options + 4 + 20);
          }
          if (A.load.Bool(options + 4 + 38)) {
            options_ffi["innerBounds"]["maxWidth"] = A.load.Int32(options + 4 + 24);
          }
          if (A.load.Bool(options + 4 + 39)) {
            options_ffi["innerBounds"]["maxHeight"] = A.load.Int32(options + 4 + 28);
          }
        }
        if (A.load.Bool(options + 48 + 40)) {
          options_ffi["outerBounds"] = {};
          if (A.load.Bool(options + 48 + 32)) {
            options_ffi["outerBounds"]["left"] = A.load.Int32(options + 48 + 0);
          }
          if (A.load.Bool(options + 48 + 33)) {
            options_ffi["outerBounds"]["top"] = A.load.Int32(options + 48 + 4);
          }
          if (A.load.Bool(options + 48 + 34)) {
            options_ffi["outerBounds"]["width"] = A.load.Int32(options + 48 + 8);
          }
          if (A.load.Bool(options + 48 + 35)) {
            options_ffi["outerBounds"]["height"] = A.load.Int32(options + 48 + 12);
          }
          if (A.load.Bool(options + 48 + 36)) {
            options_ffi["outerBounds"]["minWidth"] = A.load.Int32(options + 48 + 16);
          }
          if (A.load.Bool(options + 48 + 37)) {
            options_ffi["outerBounds"]["minHeight"] = A.load.Int32(options + 48 + 20);
          }
          if (A.load.Bool(options + 48 + 38)) {
            options_ffi["outerBounds"]["maxWidth"] = A.load.Int32(options + 48 + 24);
          }
          if (A.load.Bool(options + 48 + 39)) {
            options_ffi["outerBounds"]["maxHeight"] = A.load.Int32(options + 48 + 28);
          }
        }
        if (A.load.Bool(options + 196)) {
          options_ffi["defaultWidth"] = A.load.Int32(options + 92);
        }
        if (A.load.Bool(options + 197)) {
          options_ffi["defaultHeight"] = A.load.Int32(options + 96);
        }
        if (A.load.Bool(options + 198)) {
          options_ffi["defaultLeft"] = A.load.Int32(options + 100);
        }
        if (A.load.Bool(options + 199)) {
          options_ffi["defaultTop"] = A.load.Int32(options + 104);
        }
        if (A.load.Bool(options + 200)) {
          options_ffi["width"] = A.load.Int32(options + 108);
        }
        if (A.load.Bool(options + 201)) {
          options_ffi["height"] = A.load.Int32(options + 112);
        }
        if (A.load.Bool(options + 202)) {
          options_ffi["left"] = A.load.Int32(options + 116);
        }
        if (A.load.Bool(options + 203)) {
          options_ffi["top"] = A.load.Int32(options + 120);
        }
        if (A.load.Bool(options + 204)) {
          options_ffi["minWidth"] = A.load.Int32(options + 124);
        }
        if (A.load.Bool(options + 205)) {
          options_ffi["minHeight"] = A.load.Int32(options + 128);
        }
        if (A.load.Bool(options + 206)) {
          options_ffi["maxWidth"] = A.load.Int32(options + 132);
        }
        if (A.load.Bool(options + 207)) {
          options_ffi["maxHeight"] = A.load.Int32(options + 136);
        }
        options_ffi["type"] = A.load.Enum(options + 140, ["shell", "panel"]);
        if (A.load.Bool(options + 208)) {
          options_ffi["ime"] = A.load.Bool(options + 144);
        }
        if (A.load.Bool(options + 209)) {
          options_ffi["showInShelf"] = A.load.Bool(options + 145);
        }
        options_ffi["icon"] = A.load.Ref(options + 148, undefined);
        options_ffi["frame"] = A.load.Ref(options + 152, undefined);
        if (A.load.Bool(options + 156 + 20)) {
          options_ffi["bounds"] = {};
          if (A.load.Bool(options + 156 + 16)) {
            options_ffi["bounds"]["left"] = A.load.Int32(options + 156 + 0);
          }
          if (A.load.Bool(options + 156 + 17)) {
            options_ffi["bounds"]["top"] = A.load.Int32(options + 156 + 4);
          }
          if (A.load.Bool(options + 156 + 18)) {
            options_ffi["bounds"]["width"] = A.load.Int32(options + 156 + 8);
          }
          if (A.load.Bool(options + 156 + 19)) {
            options_ffi["bounds"]["height"] = A.load.Int32(options + 156 + 12);
          }
        }
        if (A.load.Bool(options + 210)) {
          options_ffi["alphaEnabled"] = A.load.Bool(options + 177);
        }
        options_ffi["state"] = A.load.Enum(options + 180, ["normal", "fullscreen", "maximized", "minimized"]);
        if (A.load.Bool(options + 211)) {
          options_ffi["hidden"] = A.load.Bool(options + 184);
        }
        if (A.load.Bool(options + 212)) {
          options_ffi["resizable"] = A.load.Bool(options + 185);
        }
        if (A.load.Bool(options + 213)) {
          options_ffi["singleton"] = A.load.Bool(options + 186);
        }
        if (A.load.Bool(options + 214)) {
          options_ffi["alwaysOnTop"] = A.load.Bool(options + 187);
        }
        if (A.load.Bool(options + 215)) {
          options_ffi["focused"] = A.load.Bool(options + 188);
        }
        if (A.load.Bool(options + 216)) {
          options_ffi["visibleOnAllWorkspaces"] = A.load.Bool(options + 189);
        }
        options_ffi["lockScreenAction"] = A.load.Enum(options + 192, ["new_note"]);

        const _ret = WEBEXT.app.window.create(A.H.get<object>(url), options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Current": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window && "current" in WEBEXT?.app?.window) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Current": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.current);
    },
    "call_Current": (retPtr: Pointer): void => {
      const _ret = WEBEXT.app.window.current();
      A.store.Ref(retPtr, _ret);
    },
    "try_Current": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.current();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Get": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window && "get" in WEBEXT?.app?.window) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Get": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.get);
    },
    "call_Get": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.get(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_Get": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.get(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAll": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window && "getAll" in WEBEXT?.app?.window) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.getAll);
    },
    "call_GetAll": (retPtr: Pointer): void => {
      const _ret = WEBEXT.app.window.getAll();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAll": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.getAll();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InitializeAppWindow": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window && "initializeAppWindow" in WEBEXT?.app?.window) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InitializeAppWindow": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.initializeAppWindow);
    },
    "call_InitializeAppWindow": (retPtr: Pointer, state: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.initializeAppWindow(A.H.get<object>(state));
    },
    "try_InitializeAppWindow": (retPtr: Pointer, errPtr: Pointer, state: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.initializeAppWindow(A.H.get<object>(state));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAlphaEnabledChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onAlphaEnabledChanged && "addListener" in WEBEXT?.app?.window?.onAlphaEnabledChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAlphaEnabledChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onAlphaEnabledChanged.addListener);
    },
    "call_OnAlphaEnabledChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onAlphaEnabledChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnAlphaEnabledChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onAlphaEnabledChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAlphaEnabledChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.app?.window?.onAlphaEnabledChanged &&
        "removeListener" in WEBEXT?.app?.window?.onAlphaEnabledChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAlphaEnabledChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onAlphaEnabledChanged.removeListener);
    },
    "call_OffAlphaEnabledChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onAlphaEnabledChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffAlphaEnabledChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onAlphaEnabledChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAlphaEnabledChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onAlphaEnabledChanged && "hasListener" in WEBEXT?.app?.window?.onAlphaEnabledChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAlphaEnabledChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onAlphaEnabledChanged.hasListener);
    },
    "call_HasOnAlphaEnabledChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onAlphaEnabledChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAlphaEnabledChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onAlphaEnabledChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnBoundsChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onBoundsChanged && "addListener" in WEBEXT?.app?.window?.onBoundsChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onBoundsChanged.addListener);
    },
    "call_OnBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onBoundsChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnBoundsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onBoundsChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffBoundsChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onBoundsChanged && "removeListener" in WEBEXT?.app?.window?.onBoundsChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onBoundsChanged.removeListener);
    },
    "call_OffBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onBoundsChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffBoundsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onBoundsChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnBoundsChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onBoundsChanged && "hasListener" in WEBEXT?.app?.window?.onBoundsChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onBoundsChanged.hasListener);
    },
    "call_HasOnBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onBoundsChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnBoundsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onBoundsChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnClosed": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onClosed && "addListener" in WEBEXT?.app?.window?.onClosed) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClosed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onClosed.addListener);
    },
    "call_OnClosed": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onClosed.addListener(A.H.get<object>(callback));
    },
    "try_OnClosed": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onClosed.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClosed": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onClosed && "removeListener" in WEBEXT?.app?.window?.onClosed) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClosed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onClosed.removeListener);
    },
    "call_OffClosed": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onClosed.removeListener(A.H.get<object>(callback));
    },
    "try_OffClosed": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onClosed.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClosed": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onClosed && "hasListener" in WEBEXT?.app?.window?.onClosed) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClosed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onClosed.hasListener);
    },
    "call_HasOnClosed": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onClosed.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClosed": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onClosed.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnFullscreened": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onFullscreened && "addListener" in WEBEXT?.app?.window?.onFullscreened) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnFullscreened": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onFullscreened.addListener);
    },
    "call_OnFullscreened": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onFullscreened.addListener(A.H.get<object>(callback));
    },
    "try_OnFullscreened": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onFullscreened.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffFullscreened": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onFullscreened && "removeListener" in WEBEXT?.app?.window?.onFullscreened) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffFullscreened": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onFullscreened.removeListener);
    },
    "call_OffFullscreened": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onFullscreened.removeListener(A.H.get<object>(callback));
    },
    "try_OffFullscreened": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onFullscreened.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnFullscreened": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onFullscreened && "hasListener" in WEBEXT?.app?.window?.onFullscreened) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnFullscreened": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onFullscreened.hasListener);
    },
    "call_HasOnFullscreened": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onFullscreened.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnFullscreened": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onFullscreened.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMaximized": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onMaximized && "addListener" in WEBEXT?.app?.window?.onMaximized) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMaximized": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onMaximized.addListener);
    },
    "call_OnMaximized": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onMaximized.addListener(A.H.get<object>(callback));
    },
    "try_OnMaximized": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onMaximized.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMaximized": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onMaximized && "removeListener" in WEBEXT?.app?.window?.onMaximized) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMaximized": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onMaximized.removeListener);
    },
    "call_OffMaximized": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onMaximized.removeListener(A.H.get<object>(callback));
    },
    "try_OffMaximized": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onMaximized.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMaximized": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onMaximized && "hasListener" in WEBEXT?.app?.window?.onMaximized) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMaximized": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onMaximized.hasListener);
    },
    "call_HasOnMaximized": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onMaximized.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMaximized": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onMaximized.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMinimized": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onMinimized && "addListener" in WEBEXT?.app?.window?.onMinimized) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMinimized": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onMinimized.addListener);
    },
    "call_OnMinimized": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onMinimized.addListener(A.H.get<object>(callback));
    },
    "try_OnMinimized": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onMinimized.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMinimized": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onMinimized && "removeListener" in WEBEXT?.app?.window?.onMinimized) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMinimized": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onMinimized.removeListener);
    },
    "call_OffMinimized": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onMinimized.removeListener(A.H.get<object>(callback));
    },
    "try_OffMinimized": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onMinimized.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMinimized": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onMinimized && "hasListener" in WEBEXT?.app?.window?.onMinimized) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMinimized": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onMinimized.hasListener);
    },
    "call_HasOnMinimized": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onMinimized.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMinimized": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onMinimized.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRestored": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onRestored && "addListener" in WEBEXT?.app?.window?.onRestored) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRestored": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onRestored.addListener);
    },
    "call_OnRestored": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onRestored.addListener(A.H.get<object>(callback));
    },
    "try_OnRestored": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onRestored.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRestored": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onRestored && "removeListener" in WEBEXT?.app?.window?.onRestored) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRestored": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onRestored.removeListener);
    },
    "call_OffRestored": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onRestored.removeListener(A.H.get<object>(callback));
    },
    "try_OffRestored": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onRestored.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRestored": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.window?.onRestored && "hasListener" in WEBEXT?.app?.window?.onRestored) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRestored": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.window.onRestored.hasListener);
    },
    "call_HasOnRestored": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.window.onRestored.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRestored": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.window.onRestored.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
