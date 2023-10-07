import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/virtualkeyboardprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Bounds": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 32, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
        A.store.Int64(ptr + 16, 0);
        A.store.Int64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 32, true);
        A.store.Int64(ptr + 0, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Int64(ptr + 8, x["left"] === undefined ? 0 : (x["left"] as number));
        A.store.Int64(ptr + 16, x["top"] === undefined ? 0 : (x["top"] as number));
        A.store.Int64(ptr + 24, x["width"] === undefined ? 0 : (x["width"] as number));
      }
    },
    "load_Bounds": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["height"] = A.load.Int64(ptr + 0);
      x["left"] = A.load.Int64(ptr + 8);
      x["top"] = A.load.Int64(ptr + 16);
      x["width"] = A.load.Int64(ptr + 24);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DisplayFormat": (ref: heap.Ref<string>): number => {
      const idx = ["text", "png", "html", "file"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ClipboardItem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 24, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Float64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 24, true);
        A.store.Enum(ptr + 0, ["text", "png", "html", "file"].indexOf(x["displayFormat"] as string));
        A.store.Ref(ptr + 4, x["id"]);
        A.store.Ref(ptr + 8, x["imageData"]);
        A.store.Ref(ptr + 12, x["textData"]);
        A.store.Float64(ptr + 16, x["timeCopied"] === undefined ? 0 : (x["timeCopied"] as number));
      }
    },
    "load_ClipboardItem": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["displayFormat"] = A.load.Enum(ptr + 0, ["text", "png", "html", "file"]);
      x["id"] = A.load.Ref(ptr + 4, undefined);
      x["imageData"] = A.load.Ref(ptr + 8, undefined);
      x["textData"] = A.load.Ref(ptr + 12, undefined);
      x["timeCopied"] = A.load.Float64(ptr + 16);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_KeyboardMode": (ref: heap.Ref<string>): number => {
      const idx = ["FULL_WIDTH", "FLOATING"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ContainerBehaviorOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 40, false);

        A.store.Bool(ptr + 0 + 32, false);
        A.store.Int64(ptr + 0 + 0, 0);
        A.store.Int64(ptr + 0 + 8, 0);
        A.store.Int64(ptr + 0 + 16, 0);
        A.store.Int64(ptr + 0 + 24, 0);
        A.store.Enum(ptr + 36, -1);
      } else {
        A.store.Bool(ptr + 40, true);

        if (typeof x["bounds"] === "undefined") {
          A.store.Bool(ptr + 0 + 32, false);
          A.store.Int64(ptr + 0 + 0, 0);
          A.store.Int64(ptr + 0 + 8, 0);
          A.store.Int64(ptr + 0 + 16, 0);
          A.store.Int64(ptr + 0 + 24, 0);
        } else {
          A.store.Bool(ptr + 0 + 32, true);
          A.store.Int64(ptr + 0 + 0, x["bounds"]["height"] === undefined ? 0 : (x["bounds"]["height"] as number));
          A.store.Int64(ptr + 0 + 8, x["bounds"]["left"] === undefined ? 0 : (x["bounds"]["left"] as number));
          A.store.Int64(ptr + 0 + 16, x["bounds"]["top"] === undefined ? 0 : (x["bounds"]["top"] as number));
          A.store.Int64(ptr + 0 + 24, x["bounds"]["width"] === undefined ? 0 : (x["bounds"]["width"] as number));
        }
        A.store.Enum(ptr + 36, ["FULL_WIDTH", "FLOATING"].indexOf(x["mode"] as string));
      }
    },
    "load_ContainerBehaviorOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 32)) {
        x["bounds"] = {};
        x["bounds"]["height"] = A.load.Int64(ptr + 0 + 0);
        x["bounds"]["left"] = A.load.Int64(ptr + 0 + 8);
        x["bounds"]["top"] = A.load.Int64(ptr + 0 + 16);
        x["bounds"]["width"] = A.load.Int64(ptr + 0 + 24);
      } else {
        delete x["bounds"];
      }
      x["mode"] = A.load.Enum(ptr + 36, ["FULL_WIDTH", "FLOATING"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetClipboardHistoryArgOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["itemIds"]);
      }
    },
    "load_GetClipboardHistoryArgOptions": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["itemIds"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_KeyboardConfig": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Bool(ptr + 0, x["a11ymode"] ? true : false);
        A.store.Ref(ptr + 4, x["features"]);
        A.store.Bool(ptr + 8, x["hotrodmode"] ? true : false);
        A.store.Ref(ptr + 12, x["layout"]);
      }
    },
    "load_KeyboardConfig": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["a11ymode"] = A.load.Bool(ptr + 0);
      x["features"] = A.load.Ref(ptr + 4, undefined);
      x["hotrodmode"] = A.load.Bool(ptr + 8);
      x["layout"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_KeyboardState": (ref: heap.Ref<string>): number => {
      const idx = ["ENABLED", "DISABLED", "AUTO"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_VirtualKeyboardEventType": (ref: heap.Ref<string>): number => {
      const idx = ["keyup", "keydown"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_VirtualKeyboardEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 37, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 36, false);
        A.store.Int64(ptr + 24, 0);
        A.store.Enum(ptr + 32, -1);
      } else {
        A.store.Bool(ptr + 37, true);
        A.store.Int64(ptr + 0, x["charValue"] === undefined ? 0 : (x["charValue"] as number));
        A.store.Int64(ptr + 8, x["keyCode"] === undefined ? 0 : (x["keyCode"] as number));
        A.store.Ref(ptr + 16, x["keyName"]);
        A.store.Bool(ptr + 36, "modifiers" in x ? true : false);
        A.store.Int64(ptr + 24, x["modifiers"] === undefined ? 0 : (x["modifiers"] as number));
        A.store.Enum(ptr + 32, ["keyup", "keydown"].indexOf(x["type"] as string));
      }
    },
    "load_VirtualKeyboardEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["charValue"] = A.load.Int64(ptr + 0);
      x["keyCode"] = A.load.Int64(ptr + 8);
      x["keyName"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 36)) {
        x["modifiers"] = A.load.Int64(ptr + 24);
      } else {
        delete x["modifiers"];
      }
      x["type"] = A.load.Enum(ptr + 32, ["keyup", "keydown"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_DeleteClipboardItem": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "deleteClipboardItem" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DeleteClipboardItem": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.deleteClipboardItem);
    },
    "call_DeleteClipboardItem": (retPtr: Pointer, itemId: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.deleteClipboardItem(A.H.get<object>(itemId));
    },
    "try_DeleteClipboardItem": (retPtr: Pointer, errPtr: Pointer, itemId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.deleteClipboardItem(A.H.get<object>(itemId));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetClipboardHistory": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "getClipboardHistory" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetClipboardHistory": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.getClipboardHistory);
    },
    "call_GetClipboardHistory": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["itemIds"] = A.load.Ref(options + 0, undefined);

      const _ret = WEBEXT.virtualKeyboardPrivate.getClipboardHistory(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetClipboardHistory": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["itemIds"] = A.load.Ref(options + 0, undefined);

        const _ret = WEBEXT.virtualKeyboardPrivate.getClipboardHistory(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetKeyboardConfig": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "getKeyboardConfig" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetKeyboardConfig": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.getKeyboardConfig);
    },
    "call_GetKeyboardConfig": (retPtr: Pointer): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.getKeyboardConfig();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetKeyboardConfig": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.getKeyboardConfig();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HideKeyboard": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "hideKeyboard" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HideKeyboard": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.hideKeyboard);
    },
    "call_HideKeyboard": (retPtr: Pointer): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.hideKeyboard();
      A.store.Ref(retPtr, _ret);
    },
    "try_HideKeyboard": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.hideKeyboard();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InsertText": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "insertText" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InsertText": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.insertText);
    },
    "call_InsertText": (retPtr: Pointer, text: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.insertText(A.H.get<object>(text));
      A.store.Ref(retPtr, _ret);
    },
    "try_InsertText": (retPtr: Pointer, errPtr: Pointer, text: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.insertText(A.H.get<object>(text));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_KeyboardLoaded": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "keyboardLoaded" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_KeyboardLoaded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.keyboardLoaded);
    },
    "call_KeyboardLoaded": (retPtr: Pointer): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.keyboardLoaded();
      A.store.Ref(retPtr, _ret);
    },
    "try_KeyboardLoaded": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.keyboardLoaded();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LockKeyboard": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "lockKeyboard" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LockKeyboard": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.lockKeyboard);
    },
    "call_LockKeyboard": (retPtr: Pointer, lock: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.lockKeyboard(lock === A.H.TRUE);
    },
    "try_LockKeyboard": (retPtr: Pointer, errPtr: Pointer, lock: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.lockKeyboard(lock === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnBoundsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onBoundsChanged &&
        "addListener" in WEBEXT?.virtualKeyboardPrivate?.onBoundsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onBoundsChanged.addListener);
    },
    "call_OnBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onBoundsChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnBoundsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onBoundsChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffBoundsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onBoundsChanged &&
        "removeListener" in WEBEXT?.virtualKeyboardPrivate?.onBoundsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onBoundsChanged.removeListener);
    },
    "call_OffBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onBoundsChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffBoundsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onBoundsChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnBoundsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onBoundsChanged &&
        "hasListener" in WEBEXT?.virtualKeyboardPrivate?.onBoundsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onBoundsChanged.hasListener);
    },
    "call_HasOnBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onBoundsChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnBoundsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onBoundsChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnClipboardHistoryChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onClipboardHistoryChanged &&
        "addListener" in WEBEXT?.virtualKeyboardPrivate?.onClipboardHistoryChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClipboardHistoryChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.addListener);
    },
    "call_OnClipboardHistoryChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnClipboardHistoryChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClipboardHistoryChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onClipboardHistoryChanged &&
        "removeListener" in WEBEXT?.virtualKeyboardPrivate?.onClipboardHistoryChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClipboardHistoryChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.removeListener);
    },
    "call_OffClipboardHistoryChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffClipboardHistoryChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClipboardHistoryChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onClipboardHistoryChanged &&
        "hasListener" in WEBEXT?.virtualKeyboardPrivate?.onClipboardHistoryChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClipboardHistoryChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.hasListener);
    },
    "call_HasOnClipboardHistoryChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClipboardHistoryChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnClipboardItemUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onClipboardItemUpdated &&
        "addListener" in WEBEXT?.virtualKeyboardPrivate?.onClipboardItemUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClipboardItemUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.addListener);
    },
    "call_OnClipboardItemUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.addListener(A.H.get<object>(callback));
    },
    "try_OnClipboardItemUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClipboardItemUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onClipboardItemUpdated &&
        "removeListener" in WEBEXT?.virtualKeyboardPrivate?.onClipboardItemUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClipboardItemUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.removeListener);
    },
    "call_OffClipboardItemUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.removeListener(A.H.get<object>(callback));
    },
    "try_OffClipboardItemUpdated": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClipboardItemUpdated": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onClipboardItemUpdated &&
        "hasListener" in WEBEXT?.virtualKeyboardPrivate?.onClipboardItemUpdated
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClipboardItemUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.hasListener);
    },
    "call_HasOnClipboardItemUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClipboardItemUpdated": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnColorProviderChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onColorProviderChanged &&
        "addListener" in WEBEXT?.virtualKeyboardPrivate?.onColorProviderChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnColorProviderChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.addListener);
    },
    "call_OnColorProviderChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnColorProviderChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffColorProviderChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onColorProviderChanged &&
        "removeListener" in WEBEXT?.virtualKeyboardPrivate?.onColorProviderChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffColorProviderChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.removeListener);
    },
    "call_OffColorProviderChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffColorProviderChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnColorProviderChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onColorProviderChanged &&
        "hasListener" in WEBEXT?.virtualKeyboardPrivate?.onColorProviderChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnColorProviderChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.hasListener);
    },
    "call_HasOnColorProviderChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnColorProviderChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnKeyboardClosed": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onKeyboardClosed &&
        "addListener" in WEBEXT?.virtualKeyboardPrivate?.onKeyboardClosed
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnKeyboardClosed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.addListener);
    },
    "call_OnKeyboardClosed": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.addListener(A.H.get<object>(callback));
    },
    "try_OnKeyboardClosed": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffKeyboardClosed": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onKeyboardClosed &&
        "removeListener" in WEBEXT?.virtualKeyboardPrivate?.onKeyboardClosed
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffKeyboardClosed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.removeListener);
    },
    "call_OffKeyboardClosed": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.removeListener(A.H.get<object>(callback));
    },
    "try_OffKeyboardClosed": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnKeyboardClosed": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onKeyboardClosed &&
        "hasListener" in WEBEXT?.virtualKeyboardPrivate?.onKeyboardClosed
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnKeyboardClosed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.hasListener);
    },
    "call_HasOnKeyboardClosed": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnKeyboardClosed": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnKeyboardConfigChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onKeyboardConfigChanged &&
        "addListener" in WEBEXT?.virtualKeyboardPrivate?.onKeyboardConfigChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnKeyboardConfigChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.addListener);
    },
    "call_OnKeyboardConfigChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnKeyboardConfigChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffKeyboardConfigChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onKeyboardConfigChanged &&
        "removeListener" in WEBEXT?.virtualKeyboardPrivate?.onKeyboardConfigChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffKeyboardConfigChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.removeListener);
    },
    "call_OffKeyboardConfigChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffKeyboardConfigChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnKeyboardConfigChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.virtualKeyboardPrivate?.onKeyboardConfigChanged &&
        "hasListener" in WEBEXT?.virtualKeyboardPrivate?.onKeyboardConfigChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnKeyboardConfigChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.hasListener);
    },
    "call_HasOnKeyboardConfigChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnKeyboardConfigChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenSettings": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "openSettings" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenSettings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.openSettings);
    },
    "call_OpenSettings": (retPtr: Pointer): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.openSettings();
    },
    "try_OpenSettings": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.openSettings();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenSuggestionSettings": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "openSuggestionSettings" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenSuggestionSettings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.openSuggestionSettings);
    },
    "call_OpenSuggestionSettings": (retPtr: Pointer): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.openSuggestionSettings();
    },
    "try_OpenSuggestionSettings": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.openSuggestionSettings();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_PasteClipboardItem": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "pasteClipboardItem" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_PasteClipboardItem": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.pasteClipboardItem);
    },
    "call_PasteClipboardItem": (retPtr: Pointer, itemId: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.pasteClipboardItem(A.H.get<object>(itemId));
    },
    "try_PasteClipboardItem": (retPtr: Pointer, errPtr: Pointer, itemId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.pasteClipboardItem(A.H.get<object>(itemId));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendKeyEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "sendKeyEvent" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendKeyEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.sendKeyEvent);
    },
    "call_SendKeyEvent": (retPtr: Pointer, keyEvent: Pointer): void => {
      const keyEvent_ffi = {};

      keyEvent_ffi["charValue"] = A.load.Int64(keyEvent + 0);
      keyEvent_ffi["keyCode"] = A.load.Int64(keyEvent + 8);
      keyEvent_ffi["keyName"] = A.load.Ref(keyEvent + 16, undefined);
      if (A.load.Bool(keyEvent + 36)) {
        keyEvent_ffi["modifiers"] = A.load.Int64(keyEvent + 24);
      }
      keyEvent_ffi["type"] = A.load.Enum(keyEvent + 32, ["keyup", "keydown"]);

      const _ret = WEBEXT.virtualKeyboardPrivate.sendKeyEvent(keyEvent_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SendKeyEvent": (retPtr: Pointer, errPtr: Pointer, keyEvent: Pointer): heap.Ref<boolean> => {
      try {
        const keyEvent_ffi = {};

        keyEvent_ffi["charValue"] = A.load.Int64(keyEvent + 0);
        keyEvent_ffi["keyCode"] = A.load.Int64(keyEvent + 8);
        keyEvent_ffi["keyName"] = A.load.Ref(keyEvent + 16, undefined);
        if (A.load.Bool(keyEvent + 36)) {
          keyEvent_ffi["modifiers"] = A.load.Int64(keyEvent + 24);
        }
        keyEvent_ffi["type"] = A.load.Enum(keyEvent + 32, ["keyup", "keydown"]);

        const _ret = WEBEXT.virtualKeyboardPrivate.sendKeyEvent(keyEvent_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetAreaToRemainOnScreen": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "setAreaToRemainOnScreen" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetAreaToRemainOnScreen": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.setAreaToRemainOnScreen);
    },
    "call_SetAreaToRemainOnScreen": (retPtr: Pointer, bounds: Pointer): void => {
      const bounds_ffi = {};

      bounds_ffi["height"] = A.load.Int64(bounds + 0);
      bounds_ffi["left"] = A.load.Int64(bounds + 8);
      bounds_ffi["top"] = A.load.Int64(bounds + 16);
      bounds_ffi["width"] = A.load.Int64(bounds + 24);

      const _ret = WEBEXT.virtualKeyboardPrivate.setAreaToRemainOnScreen(bounds_ffi);
    },
    "try_SetAreaToRemainOnScreen": (retPtr: Pointer, errPtr: Pointer, bounds: Pointer): heap.Ref<boolean> => {
      try {
        const bounds_ffi = {};

        bounds_ffi["height"] = A.load.Int64(bounds + 0);
        bounds_ffi["left"] = A.load.Int64(bounds + 8);
        bounds_ffi["top"] = A.load.Int64(bounds + 16);
        bounds_ffi["width"] = A.load.Int64(bounds + 24);

        const _ret = WEBEXT.virtualKeyboardPrivate.setAreaToRemainOnScreen(bounds_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetContainerBehavior": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "setContainerBehavior" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetContainerBehavior": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.setContainerBehavior);
    },
    "call_SetContainerBehavior": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 0 + 32)) {
        options_ffi["bounds"] = {};
        options_ffi["bounds"]["height"] = A.load.Int64(options + 0 + 0);
        options_ffi["bounds"]["left"] = A.load.Int64(options + 0 + 8);
        options_ffi["bounds"]["top"] = A.load.Int64(options + 0 + 16);
        options_ffi["bounds"]["width"] = A.load.Int64(options + 0 + 24);
      }
      options_ffi["mode"] = A.load.Enum(options + 36, ["FULL_WIDTH", "FLOATING"]);

      const _ret = WEBEXT.virtualKeyboardPrivate.setContainerBehavior(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetContainerBehavior": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 0 + 32)) {
          options_ffi["bounds"] = {};
          options_ffi["bounds"]["height"] = A.load.Int64(options + 0 + 0);
          options_ffi["bounds"]["left"] = A.load.Int64(options + 0 + 8);
          options_ffi["bounds"]["top"] = A.load.Int64(options + 0 + 16);
          options_ffi["bounds"]["width"] = A.load.Int64(options + 0 + 24);
        }
        options_ffi["mode"] = A.load.Enum(options + 36, ["FULL_WIDTH", "FLOATING"]);

        const _ret = WEBEXT.virtualKeyboardPrivate.setContainerBehavior(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetDraggableArea": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "setDraggableArea" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetDraggableArea": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.setDraggableArea);
    },
    "call_SetDraggableArea": (retPtr: Pointer, bounds: Pointer): void => {
      const bounds_ffi = {};

      bounds_ffi["height"] = A.load.Int64(bounds + 0);
      bounds_ffi["left"] = A.load.Int64(bounds + 8);
      bounds_ffi["top"] = A.load.Int64(bounds + 16);
      bounds_ffi["width"] = A.load.Int64(bounds + 24);

      const _ret = WEBEXT.virtualKeyboardPrivate.setDraggableArea(bounds_ffi);
    },
    "try_SetDraggableArea": (retPtr: Pointer, errPtr: Pointer, bounds: Pointer): heap.Ref<boolean> => {
      try {
        const bounds_ffi = {};

        bounds_ffi["height"] = A.load.Int64(bounds + 0);
        bounds_ffi["left"] = A.load.Int64(bounds + 8);
        bounds_ffi["top"] = A.load.Int64(bounds + 16);
        bounds_ffi["width"] = A.load.Int64(bounds + 24);

        const _ret = WEBEXT.virtualKeyboardPrivate.setDraggableArea(bounds_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetHitTestBounds": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "setHitTestBounds" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetHitTestBounds": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.setHitTestBounds);
    },
    "call_SetHitTestBounds": (retPtr: Pointer, boundsList: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.setHitTestBounds(A.H.get<object>(boundsList));
    },
    "try_SetHitTestBounds": (retPtr: Pointer, errPtr: Pointer, boundsList: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.setHitTestBounds(A.H.get<object>(boundsList));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetHotrodKeyboard": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "setHotrodKeyboard" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetHotrodKeyboard": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.setHotrodKeyboard);
    },
    "call_SetHotrodKeyboard": (retPtr: Pointer, enable: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.setHotrodKeyboard(enable === A.H.TRUE);
    },
    "try_SetHotrodKeyboard": (retPtr: Pointer, errPtr: Pointer, enable: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.setHotrodKeyboard(enable === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetKeyboardState": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "setKeyboardState" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetKeyboardState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.setKeyboardState);
    },
    "call_SetKeyboardState": (retPtr: Pointer, state: number): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.setKeyboardState(
        state > 0 && state <= 3 ? ["ENABLED", "DISABLED", "AUTO"][state - 1] : undefined
      );
    },
    "try_SetKeyboardState": (retPtr: Pointer, errPtr: Pointer, state: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.setKeyboardState(
          state > 0 && state <= 3 ? ["ENABLED", "DISABLED", "AUTO"][state - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetOccludedBounds": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "setOccludedBounds" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetOccludedBounds": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.setOccludedBounds);
    },
    "call_SetOccludedBounds": (retPtr: Pointer, boundsList: heap.Ref<object>): void => {
      const _ret = WEBEXT.virtualKeyboardPrivate.setOccludedBounds(A.H.get<object>(boundsList));
    },
    "try_SetOccludedBounds": (retPtr: Pointer, errPtr: Pointer, boundsList: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.virtualKeyboardPrivate.setOccludedBounds(A.H.get<object>(boundsList));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetWindowBoundsInScreen": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboardPrivate && "setWindowBoundsInScreen" in WEBEXT?.virtualKeyboardPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetWindowBoundsInScreen": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboardPrivate.setWindowBoundsInScreen);
    },
    "call_SetWindowBoundsInScreen": (retPtr: Pointer, bounds: Pointer): void => {
      const bounds_ffi = {};

      bounds_ffi["height"] = A.load.Int64(bounds + 0);
      bounds_ffi["left"] = A.load.Int64(bounds + 8);
      bounds_ffi["top"] = A.load.Int64(bounds + 16);
      bounds_ffi["width"] = A.load.Int64(bounds + 24);

      const _ret = WEBEXT.virtualKeyboardPrivate.setWindowBoundsInScreen(bounds_ffi);
    },
    "try_SetWindowBoundsInScreen": (retPtr: Pointer, errPtr: Pointer, bounds: Pointer): heap.Ref<boolean> => {
      try {
        const bounds_ffi = {};

        bounds_ffi["height"] = A.load.Int64(bounds + 0);
        bounds_ffi["left"] = A.load.Int64(bounds + 8);
        bounds_ffi["top"] = A.load.Int64(bounds + 16);
        bounds_ffi["width"] = A.load.Int64(bounds + 24);

        const _ret = WEBEXT.virtualKeyboardPrivate.setWindowBoundsInScreen(bounds_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
