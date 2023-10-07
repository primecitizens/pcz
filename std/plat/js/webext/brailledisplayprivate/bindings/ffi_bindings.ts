import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/brailledisplayprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_DisplayState": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 18, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 19, false);
        A.store.Int32(ptr + 12, 0);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Bool(ptr + 16, "available" in x ? true : false);
        A.store.Bool(ptr + 0, x["available"] ? true : false);
        A.store.Bool(ptr + 17, "textRowCount" in x ? true : false);
        A.store.Int32(ptr + 4, x["textRowCount"] === undefined ? 0 : (x["textRowCount"] as number));
        A.store.Bool(ptr + 18, "textColumnCount" in x ? true : false);
        A.store.Int32(ptr + 8, x["textColumnCount"] === undefined ? 0 : (x["textColumnCount"] as number));
        A.store.Bool(ptr + 19, "cellSize" in x ? true : false);
        A.store.Int32(ptr + 12, x["cellSize"] === undefined ? 0 : (x["cellSize"] as number));
      }
    },
    "load_DisplayState": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["available"] = A.load.Bool(ptr + 0);
      } else {
        delete x["available"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["textRowCount"] = A.load.Int32(ptr + 4);
      } else {
        delete x["textRowCount"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["textColumnCount"] = A.load.Int32(ptr + 8);
      } else {
        delete x["textColumnCount"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["cellSize"] = A.load.Int32(ptr + 12);
      } else {
        delete x["cellSize"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_KeyCommand": (ref: heap.Ref<string>): number => {
      const idx = [
        "line_up",
        "line_down",
        "pan_left",
        "pan_right",
        "top",
        "bottom",
        "routing",
        "secondary_routing",
        "dots",
        "chord",
        "standard_key",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_KeyEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 30, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 24, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 25, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 27, false);
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 22, false);
        A.store.Bool(ptr + 29, false);
        A.store.Bool(ptr + 23, false);
      } else {
        A.store.Bool(ptr + 30, true);
        A.store.Enum(
          ptr + 0,
          [
            "line_up",
            "line_down",
            "pan_left",
            "pan_right",
            "top",
            "bottom",
            "routing",
            "secondary_routing",
            "dots",
            "chord",
            "standard_key",
          ].indexOf(x["command"] as string)
        );
        A.store.Bool(ptr + 24, "displayPosition" in x ? true : false);
        A.store.Int32(ptr + 4, x["displayPosition"] === undefined ? 0 : (x["displayPosition"] as number));
        A.store.Bool(ptr + 25, "brailleDots" in x ? true : false);
        A.store.Int32(ptr + 8, x["brailleDots"] === undefined ? 0 : (x["brailleDots"] as number));
        A.store.Ref(ptr + 12, x["standardKeyCode"]);
        A.store.Ref(ptr + 16, x["standardKeyChar"]);
        A.store.Bool(ptr + 26, "spaceKey" in x ? true : false);
        A.store.Bool(ptr + 20, x["spaceKey"] ? true : false);
        A.store.Bool(ptr + 27, "altKey" in x ? true : false);
        A.store.Bool(ptr + 21, x["altKey"] ? true : false);
        A.store.Bool(ptr + 28, "shiftKey" in x ? true : false);
        A.store.Bool(ptr + 22, x["shiftKey"] ? true : false);
        A.store.Bool(ptr + 29, "ctrlKey" in x ? true : false);
        A.store.Bool(ptr + 23, x["ctrlKey"] ? true : false);
      }
    },
    "load_KeyEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["command"] = A.load.Enum(ptr + 0, [
        "line_up",
        "line_down",
        "pan_left",
        "pan_right",
        "top",
        "bottom",
        "routing",
        "secondary_routing",
        "dots",
        "chord",
        "standard_key",
      ]);
      if (A.load.Bool(ptr + 24)) {
        x["displayPosition"] = A.load.Int32(ptr + 4);
      } else {
        delete x["displayPosition"];
      }
      if (A.load.Bool(ptr + 25)) {
        x["brailleDots"] = A.load.Int32(ptr + 8);
      } else {
        delete x["brailleDots"];
      }
      x["standardKeyCode"] = A.load.Ref(ptr + 12, undefined);
      x["standardKeyChar"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 26)) {
        x["spaceKey"] = A.load.Bool(ptr + 20);
      } else {
        delete x["spaceKey"];
      }
      if (A.load.Bool(ptr + 27)) {
        x["altKey"] = A.load.Bool(ptr + 21);
      } else {
        delete x["altKey"];
      }
      if (A.load.Bool(ptr + 28)) {
        x["shiftKey"] = A.load.Bool(ptr + 22);
      } else {
        delete x["shiftKey"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["ctrlKey"] = A.load.Bool(ptr + 23);
      } else {
        delete x["ctrlKey"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetDisplayState": (): heap.Ref<boolean> => {
      if (WEBEXT?.brailleDisplayPrivate && "getDisplayState" in WEBEXT?.brailleDisplayPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDisplayState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.brailleDisplayPrivate.getDisplayState);
    },
    "call_GetDisplayState": (retPtr: Pointer): void => {
      const _ret = WEBEXT.brailleDisplayPrivate.getDisplayState();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDisplayState": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.brailleDisplayPrivate.getDisplayState();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDisplayStateChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.brailleDisplayPrivate?.onDisplayStateChanged &&
        "addListener" in WEBEXT?.brailleDisplayPrivate?.onDisplayStateChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDisplayStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.addListener);
    },
    "call_OnDisplayStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnDisplayStateChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDisplayStateChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.brailleDisplayPrivate?.onDisplayStateChanged &&
        "removeListener" in WEBEXT?.brailleDisplayPrivate?.onDisplayStateChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDisplayStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.removeListener);
    },
    "call_OffDisplayStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffDisplayStateChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDisplayStateChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.brailleDisplayPrivate?.onDisplayStateChanged &&
        "hasListener" in WEBEXT?.brailleDisplayPrivate?.onDisplayStateChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDisplayStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.hasListener);
    },
    "call_HasOnDisplayStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDisplayStateChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnKeyEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.brailleDisplayPrivate?.onKeyEvent && "addListener" in WEBEXT?.brailleDisplayPrivate?.onKeyEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnKeyEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.brailleDisplayPrivate.onKeyEvent.addListener);
    },
    "call_OnKeyEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.brailleDisplayPrivate.onKeyEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnKeyEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.brailleDisplayPrivate.onKeyEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffKeyEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.brailleDisplayPrivate?.onKeyEvent && "removeListener" in WEBEXT?.brailleDisplayPrivate?.onKeyEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffKeyEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.brailleDisplayPrivate.onKeyEvent.removeListener);
    },
    "call_OffKeyEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.brailleDisplayPrivate.onKeyEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffKeyEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.brailleDisplayPrivate.onKeyEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnKeyEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.brailleDisplayPrivate?.onKeyEvent && "hasListener" in WEBEXT?.brailleDisplayPrivate?.onKeyEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnKeyEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.brailleDisplayPrivate.onKeyEvent.hasListener);
    },
    "call_HasOnKeyEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.brailleDisplayPrivate.onKeyEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnKeyEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.brailleDisplayPrivate.onKeyEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateBluetoothBrailleDisplayAddress": (): heap.Ref<boolean> => {
      if (WEBEXT?.brailleDisplayPrivate && "updateBluetoothBrailleDisplayAddress" in WEBEXT?.brailleDisplayPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateBluetoothBrailleDisplayAddress": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.brailleDisplayPrivate.updateBluetoothBrailleDisplayAddress);
    },
    "call_UpdateBluetoothBrailleDisplayAddress": (retPtr: Pointer, address: heap.Ref<object>): void => {
      const _ret = WEBEXT.brailleDisplayPrivate.updateBluetoothBrailleDisplayAddress(A.H.get<object>(address));
    },
    "try_UpdateBluetoothBrailleDisplayAddress": (
      retPtr: Pointer,
      errPtr: Pointer,
      address: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.brailleDisplayPrivate.updateBluetoothBrailleDisplayAddress(A.H.get<object>(address));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_WriteDots": (): heap.Ref<boolean> => {
      if (WEBEXT?.brailleDisplayPrivate && "writeDots" in WEBEXT?.brailleDisplayPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_WriteDots": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.brailleDisplayPrivate.writeDots);
    },
    "call_WriteDots": (retPtr: Pointer, cells: heap.Ref<object>, columns: number, rows: number): void => {
      const _ret = WEBEXT.brailleDisplayPrivate.writeDots(A.H.get<object>(cells), columns, rows);
    },
    "try_WriteDots": (
      retPtr: Pointer,
      errPtr: Pointer,
      cells: heap.Ref<object>,
      columns: number,
      rows: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.brailleDisplayPrivate.writeDots(A.H.get<object>(cells), columns, rows);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
