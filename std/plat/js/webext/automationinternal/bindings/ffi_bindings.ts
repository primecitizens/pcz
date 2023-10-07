import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/automationinternal", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AXEventParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 40, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 36, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 37, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 38, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Bool(ptr + 39, false);
        A.store.Int32(ptr + 32, 0);
      } else {
        A.store.Bool(ptr + 40, true);
        A.store.Ref(ptr + 0, x["treeID"]);
        A.store.Bool(ptr + 36, "targetID" in x ? true : false);
        A.store.Int32(ptr + 4, x["targetID"] === undefined ? 0 : (x["targetID"] as number));
        A.store.Ref(ptr + 8, x["eventType"]);
        A.store.Ref(ptr + 12, x["eventFrom"]);
        A.store.Bool(ptr + 37, "mouseX" in x ? true : false);
        A.store.Float64(ptr + 16, x["mouseX"] === undefined ? 0 : (x["mouseX"] as number));
        A.store.Bool(ptr + 38, "mouseY" in x ? true : false);
        A.store.Float64(ptr + 24, x["mouseY"] === undefined ? 0 : (x["mouseY"] as number));
        A.store.Bool(ptr + 39, "actionRequestID" in x ? true : false);
        A.store.Int32(ptr + 32, x["actionRequestID"] === undefined ? 0 : (x["actionRequestID"] as number));
      }
    },
    "load_AXEventParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["treeID"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 36)) {
        x["targetID"] = A.load.Int32(ptr + 4);
      } else {
        delete x["targetID"];
      }
      x["eventType"] = A.load.Ref(ptr + 8, undefined);
      x["eventFrom"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 37)) {
        x["mouseX"] = A.load.Float64(ptr + 16);
      } else {
        delete x["mouseX"];
      }
      if (A.load.Bool(ptr + 38)) {
        x["mouseY"] = A.load.Float64(ptr + 24);
      } else {
        delete x["mouseY"];
      }
      if (A.load.Bool(ptr + 39)) {
        x["actionRequestID"] = A.load.Int32(ptr + 32);
      } else {
        delete x["actionRequestID"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AXTextLocationParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 39, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 32, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 33, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 34, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 35, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Bool(ptr + 36, false);
        A.store.Int32(ptr + 20, 0);
        A.store.Bool(ptr + 37, false);
        A.store.Int32(ptr + 24, 0);
        A.store.Bool(ptr + 38, false);
        A.store.Int32(ptr + 28, 0);
      } else {
        A.store.Bool(ptr + 39, true);
        A.store.Ref(ptr + 0, x["treeID"]);
        A.store.Bool(ptr + 32, "nodeID" in x ? true : false);
        A.store.Int32(ptr + 4, x["nodeID"] === undefined ? 0 : (x["nodeID"] as number));
        A.store.Bool(ptr + 33, "result" in x ? true : false);
        A.store.Bool(ptr + 8, x["result"] ? true : false);
        A.store.Bool(ptr + 34, "left" in x ? true : false);
        A.store.Int32(ptr + 12, x["left"] === undefined ? 0 : (x["left"] as number));
        A.store.Bool(ptr + 35, "top" in x ? true : false);
        A.store.Int32(ptr + 16, x["top"] === undefined ? 0 : (x["top"] as number));
        A.store.Bool(ptr + 36, "width" in x ? true : false);
        A.store.Int32(ptr + 20, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Bool(ptr + 37, "height" in x ? true : false);
        A.store.Int32(ptr + 24, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Bool(ptr + 38, "requestID" in x ? true : false);
        A.store.Int32(ptr + 28, x["requestID"] === undefined ? 0 : (x["requestID"] as number));
      }
    },
    "load_AXTextLocationParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["treeID"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 32)) {
        x["nodeID"] = A.load.Int32(ptr + 4);
      } else {
        delete x["nodeID"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["result"] = A.load.Bool(ptr + 8);
      } else {
        delete x["result"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["left"] = A.load.Int32(ptr + 12);
      } else {
        delete x["left"];
      }
      if (A.load.Bool(ptr + 35)) {
        x["top"] = A.load.Int32(ptr + 16);
      } else {
        delete x["top"];
      }
      if (A.load.Bool(ptr + 36)) {
        x["width"] = A.load.Int32(ptr + 20);
      } else {
        delete x["width"];
      }
      if (A.load.Bool(ptr + 37)) {
        x["height"] = A.load.Int32(ptr + 24);
      } else {
        delete x["height"];
      }
      if (A.load.Bool(ptr + 38)) {
        x["requestID"] = A.load.Int32(ptr + 28);
      } else {
        delete x["requestID"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetImageDataParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "maxWidth" in x ? true : false);
        A.store.Int32(ptr + 0, x["maxWidth"] === undefined ? 0 : (x["maxWidth"] as number));
        A.store.Bool(ptr + 9, "maxHeight" in x ? true : false);
        A.store.Int32(ptr + 4, x["maxHeight"] === undefined ? 0 : (x["maxHeight"] as number));
      }
    },
    "load_GetImageDataParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["maxWidth"] = A.load.Int32(ptr + 0);
      } else {
        delete x["maxWidth"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["maxHeight"] = A.load.Int32(ptr + 4);
      } else {
        delete x["maxHeight"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetTextLocationDataParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "startIndex" in x ? true : false);
        A.store.Int32(ptr + 0, x["startIndex"] === undefined ? 0 : (x["startIndex"] as number));
        A.store.Bool(ptr + 9, "endIndex" in x ? true : false);
        A.store.Int32(ptr + 4, x["endIndex"] === undefined ? 0 : (x["endIndex"] as number));
      }
    },
    "load_GetTextLocationDataParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["startIndex"] = A.load.Int32(ptr + 0);
      } else {
        delete x["startIndex"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["endIndex"] = A.load.Int32(ptr + 4);
      } else {
        delete x["endIndex"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_HitTestParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Bool(ptr + 12, "x" in x ? true : false);
        A.store.Int32(ptr + 0, x["x"] === undefined ? 0 : (x["x"] as number));
        A.store.Bool(ptr + 13, "y" in x ? true : false);
        A.store.Int32(ptr + 4, x["y"] === undefined ? 0 : (x["y"] as number));
        A.store.Ref(ptr + 8, x["eventToFire"]);
      }
    },
    "load_HitTestParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["x"] = A.load.Int32(ptr + 0);
      } else {
        delete x["x"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["y"] = A.load.Int32(ptr + 4);
      } else {
        delete x["y"];
      }
      x["eventToFire"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PerformActionRequiredParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 12, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Ref(ptr + 0, x["treeID"]);
        A.store.Bool(ptr + 16, "automationNodeID" in x ? true : false);
        A.store.Int32(ptr + 4, x["automationNodeID"] === undefined ? 0 : (x["automationNodeID"] as number));
        A.store.Ref(ptr + 8, x["actionType"]);
        A.store.Bool(ptr + 17, "requestID" in x ? true : false);
        A.store.Int32(ptr + 12, x["requestID"] === undefined ? 0 : (x["requestID"] as number));
      }
    },
    "load_PerformActionRequiredParams": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["treeID"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["automationNodeID"] = A.load.Int32(ptr + 4);
      } else {
        delete x["automationNodeID"];
      }
      x["actionType"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 17)) {
        x["requestID"] = A.load.Int32(ptr + 12);
      } else {
        delete x["requestID"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PerformCustomActionParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "customActionID" in x ? true : false);
        A.store.Int32(ptr + 0, x["customActionID"] === undefined ? 0 : (x["customActionID"] as number));
      }
    },
    "load_PerformCustomActionParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["customActionID"] = A.load.Int32(ptr + 0);
      } else {
        delete x["customActionID"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ReplaceSelectedTextParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["value"]);
      }
    },
    "load_ReplaceSelectedTextParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["value"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ScrollToPointParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "x" in x ? true : false);
        A.store.Int32(ptr + 0, x["x"] === undefined ? 0 : (x["x"] as number));
        A.store.Bool(ptr + 9, "y" in x ? true : false);
        A.store.Int32(ptr + 4, x["y"] === undefined ? 0 : (x["y"] as number));
      }
    },
    "load_ScrollToPointParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["x"] = A.load.Int32(ptr + 0);
      } else {
        delete x["x"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["y"] = A.load.Int32(ptr + 4);
      } else {
        delete x["y"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ScrollToPositionAtRowColumnParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "row" in x ? true : false);
        A.store.Int32(ptr + 0, x["row"] === undefined ? 0 : (x["row"] as number));
        A.store.Bool(ptr + 9, "column" in x ? true : false);
        A.store.Int32(ptr + 4, x["column"] === undefined ? 0 : (x["column"] as number));
      }
    },
    "load_ScrollToPositionAtRowColumnParams": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["row"] = A.load.Int32(ptr + 0);
      } else {
        delete x["row"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["column"] = A.load.Int32(ptr + 4);
      } else {
        delete x["column"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetScrollOffsetParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "x" in x ? true : false);
        A.store.Int32(ptr + 0, x["x"] === undefined ? 0 : (x["x"] as number));
        A.store.Bool(ptr + 9, "y" in x ? true : false);
        A.store.Int32(ptr + 4, x["y"] === undefined ? 0 : (x["y"] as number));
      }
    },
    "load_SetScrollOffsetParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["x"] = A.load.Int32(ptr + 0);
      } else {
        delete x["x"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["y"] = A.load.Int32(ptr + 4);
      } else {
        delete x["y"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetSelectionParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 14, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 15, true);
        A.store.Bool(ptr + 12, "focusNodeID" in x ? true : false);
        A.store.Int32(ptr + 0, x["focusNodeID"] === undefined ? 0 : (x["focusNodeID"] as number));
        A.store.Bool(ptr + 13, "anchorOffset" in x ? true : false);
        A.store.Int32(ptr + 4, x["anchorOffset"] === undefined ? 0 : (x["anchorOffset"] as number));
        A.store.Bool(ptr + 14, "focusOffset" in x ? true : false);
        A.store.Int32(ptr + 8, x["focusOffset"] === undefined ? 0 : (x["focusOffset"] as number));
      }
    },
    "load_SetSelectionParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["focusNodeID"] = A.load.Int32(ptr + 0);
      } else {
        delete x["focusNodeID"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["anchorOffset"] = A.load.Int32(ptr + 4);
      } else {
        delete x["anchorOffset"];
      }
      if (A.load.Bool(ptr + 14)) {
        x["focusOffset"] = A.load.Int32(ptr + 8);
      } else {
        delete x["focusOffset"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetValueParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["value"]);
      }
    },
    "load_SetValueParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["value"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_DisableDesktop": (): heap.Ref<boolean> => {
      if (WEBEXT?.automationInternal && "disableDesktop" in WEBEXT?.automationInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DisableDesktop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.disableDesktop);
    },
    "call_DisableDesktop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.disableDesktop(A.H.get<object>(callback));
    },
    "try_DisableDesktop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.disableDesktop(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_EnableDesktop": (): heap.Ref<boolean> => {
      if (WEBEXT?.automationInternal && "enableDesktop" in WEBEXT?.automationInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EnableDesktop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.enableDesktop);
    },
    "call_EnableDesktop": (retPtr: Pointer): void => {
      const _ret = WEBEXT.automationInternal.enableDesktop();
      A.store.Ref(retPtr, _ret);
    },
    "try_EnableDesktop": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.enableDesktop();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_EnableTree": (): heap.Ref<boolean> => {
      if (WEBEXT?.automationInternal && "enableTree" in WEBEXT?.automationInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EnableTree": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.enableTree);
    },
    "call_EnableTree": (retPtr: Pointer, tree_id: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.enableTree(A.H.get<object>(tree_id));
    },
    "try_EnableTree": (retPtr: Pointer, errPtr: Pointer, tree_id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.enableTree(A.H.get<object>(tree_id));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAccessibilityEvent": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onAccessibilityEvent &&
        "addListener" in WEBEXT?.automationInternal?.onAccessibilityEvent
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAccessibilityEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onAccessibilityEvent.addListener);
    },
    "call_OnAccessibilityEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onAccessibilityEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnAccessibilityEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onAccessibilityEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAccessibilityEvent": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onAccessibilityEvent &&
        "removeListener" in WEBEXT?.automationInternal?.onAccessibilityEvent
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAccessibilityEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onAccessibilityEvent.removeListener);
    },
    "call_OffAccessibilityEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onAccessibilityEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffAccessibilityEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onAccessibilityEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAccessibilityEvent": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onAccessibilityEvent &&
        "hasListener" in WEBEXT?.automationInternal?.onAccessibilityEvent
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAccessibilityEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onAccessibilityEvent.hasListener);
    },
    "call_HasOnAccessibilityEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onAccessibilityEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAccessibilityEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onAccessibilityEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAccessibilityTreeDestroyed": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onAccessibilityTreeDestroyed &&
        "addListener" in WEBEXT?.automationInternal?.onAccessibilityTreeDestroyed
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAccessibilityTreeDestroyed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onAccessibilityTreeDestroyed.addListener);
    },
    "call_OnAccessibilityTreeDestroyed": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onAccessibilityTreeDestroyed.addListener(A.H.get<object>(callback));
    },
    "try_OnAccessibilityTreeDestroyed": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onAccessibilityTreeDestroyed.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAccessibilityTreeDestroyed": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onAccessibilityTreeDestroyed &&
        "removeListener" in WEBEXT?.automationInternal?.onAccessibilityTreeDestroyed
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAccessibilityTreeDestroyed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onAccessibilityTreeDestroyed.removeListener);
    },
    "call_OffAccessibilityTreeDestroyed": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onAccessibilityTreeDestroyed.removeListener(A.H.get<object>(callback));
    },
    "try_OffAccessibilityTreeDestroyed": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onAccessibilityTreeDestroyed.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAccessibilityTreeDestroyed": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onAccessibilityTreeDestroyed &&
        "hasListener" in WEBEXT?.automationInternal?.onAccessibilityTreeDestroyed
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAccessibilityTreeDestroyed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onAccessibilityTreeDestroyed.hasListener);
    },
    "call_HasOnAccessibilityTreeDestroyed": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onAccessibilityTreeDestroyed.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAccessibilityTreeDestroyed": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onAccessibilityTreeDestroyed.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAccessibilityTreeSerializationError": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onAccessibilityTreeSerializationError &&
        "addListener" in WEBEXT?.automationInternal?.onAccessibilityTreeSerializationError
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAccessibilityTreeSerializationError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onAccessibilityTreeSerializationError.addListener);
    },
    "call_OnAccessibilityTreeSerializationError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onAccessibilityTreeSerializationError.addListener(
        A.H.get<object>(callback)
      );
    },
    "try_OnAccessibilityTreeSerializationError": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onAccessibilityTreeSerializationError.addListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAccessibilityTreeSerializationError": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onAccessibilityTreeSerializationError &&
        "removeListener" in WEBEXT?.automationInternal?.onAccessibilityTreeSerializationError
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAccessibilityTreeSerializationError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onAccessibilityTreeSerializationError.removeListener);
    },
    "call_OffAccessibilityTreeSerializationError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onAccessibilityTreeSerializationError.removeListener(
        A.H.get<object>(callback)
      );
    },
    "try_OffAccessibilityTreeSerializationError": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onAccessibilityTreeSerializationError.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAccessibilityTreeSerializationError": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onAccessibilityTreeSerializationError &&
        "hasListener" in WEBEXT?.automationInternal?.onAccessibilityTreeSerializationError
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAccessibilityTreeSerializationError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onAccessibilityTreeSerializationError.hasListener);
    },
    "call_HasOnAccessibilityTreeSerializationError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onAccessibilityTreeSerializationError.hasListener(
        A.H.get<object>(callback)
      );
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAccessibilityTreeSerializationError": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onAccessibilityTreeSerializationError.hasListener(
          A.H.get<object>(callback)
        );
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnActionResult": (): heap.Ref<boolean> => {
      if (WEBEXT?.automationInternal?.onActionResult && "addListener" in WEBEXT?.automationInternal?.onActionResult) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnActionResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onActionResult.addListener);
    },
    "call_OnActionResult": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onActionResult.addListener(A.H.get<object>(callback));
    },
    "try_OnActionResult": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onActionResult.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffActionResult": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onActionResult &&
        "removeListener" in WEBEXT?.automationInternal?.onActionResult
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffActionResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onActionResult.removeListener);
    },
    "call_OffActionResult": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onActionResult.removeListener(A.H.get<object>(callback));
    },
    "try_OffActionResult": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onActionResult.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnActionResult": (): heap.Ref<boolean> => {
      if (WEBEXT?.automationInternal?.onActionResult && "hasListener" in WEBEXT?.automationInternal?.onActionResult) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnActionResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onActionResult.hasListener);
    },
    "call_HasOnActionResult": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onActionResult.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnActionResult": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onActionResult.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAllAutomationEventListenersRemoved": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onAllAutomationEventListenersRemoved &&
        "addListener" in WEBEXT?.automationInternal?.onAllAutomationEventListenersRemoved
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAllAutomationEventListenersRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.addListener);
    },
    "call_OnAllAutomationEventListenersRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.addListener(
        A.H.get<object>(callback)
      );
    },
    "try_OnAllAutomationEventListenersRemoved": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.addListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAllAutomationEventListenersRemoved": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onAllAutomationEventListenersRemoved &&
        "removeListener" in WEBEXT?.automationInternal?.onAllAutomationEventListenersRemoved
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAllAutomationEventListenersRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.removeListener);
    },
    "call_OffAllAutomationEventListenersRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.removeListener(
        A.H.get<object>(callback)
      );
    },
    "try_OffAllAutomationEventListenersRemoved": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAllAutomationEventListenersRemoved": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onAllAutomationEventListenersRemoved &&
        "hasListener" in WEBEXT?.automationInternal?.onAllAutomationEventListenersRemoved
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAllAutomationEventListenersRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.hasListener);
    },
    "call_HasOnAllAutomationEventListenersRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.hasListener(
        A.H.get<object>(callback)
      );
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAllAutomationEventListenersRemoved": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.hasListener(
          A.H.get<object>(callback)
        );
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnChildTreeID": (): heap.Ref<boolean> => {
      if (WEBEXT?.automationInternal?.onChildTreeID && "addListener" in WEBEXT?.automationInternal?.onChildTreeID) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnChildTreeID": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onChildTreeID.addListener);
    },
    "call_OnChildTreeID": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onChildTreeID.addListener(A.H.get<object>(callback));
    },
    "try_OnChildTreeID": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onChildTreeID.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffChildTreeID": (): heap.Ref<boolean> => {
      if (WEBEXT?.automationInternal?.onChildTreeID && "removeListener" in WEBEXT?.automationInternal?.onChildTreeID) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffChildTreeID": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onChildTreeID.removeListener);
    },
    "call_OffChildTreeID": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onChildTreeID.removeListener(A.H.get<object>(callback));
    },
    "try_OffChildTreeID": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onChildTreeID.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnChildTreeID": (): heap.Ref<boolean> => {
      if (WEBEXT?.automationInternal?.onChildTreeID && "hasListener" in WEBEXT?.automationInternal?.onChildTreeID) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnChildTreeID": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onChildTreeID.hasListener);
    },
    "call_HasOnChildTreeID": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onChildTreeID.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnChildTreeID": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onChildTreeID.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnGetTextLocationResult": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onGetTextLocationResult &&
        "addListener" in WEBEXT?.automationInternal?.onGetTextLocationResult
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnGetTextLocationResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onGetTextLocationResult.addListener);
    },
    "call_OnGetTextLocationResult": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onGetTextLocationResult.addListener(A.H.get<object>(callback));
    },
    "try_OnGetTextLocationResult": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onGetTextLocationResult.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffGetTextLocationResult": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onGetTextLocationResult &&
        "removeListener" in WEBEXT?.automationInternal?.onGetTextLocationResult
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffGetTextLocationResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onGetTextLocationResult.removeListener);
    },
    "call_OffGetTextLocationResult": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onGetTextLocationResult.removeListener(A.H.get<object>(callback));
    },
    "try_OffGetTextLocationResult": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onGetTextLocationResult.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnGetTextLocationResult": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onGetTextLocationResult &&
        "hasListener" in WEBEXT?.automationInternal?.onGetTextLocationResult
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnGetTextLocationResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onGetTextLocationResult.hasListener);
    },
    "call_HasOnGetTextLocationResult": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onGetTextLocationResult.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnGetTextLocationResult": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onGetTextLocationResult.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnNodesRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.automationInternal?.onNodesRemoved && "addListener" in WEBEXT?.automationInternal?.onNodesRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnNodesRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onNodesRemoved.addListener);
    },
    "call_OnNodesRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onNodesRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnNodesRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onNodesRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffNodesRemoved": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.automationInternal?.onNodesRemoved &&
        "removeListener" in WEBEXT?.automationInternal?.onNodesRemoved
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffNodesRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onNodesRemoved.removeListener);
    },
    "call_OffNodesRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onNodesRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffNodesRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onNodesRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnNodesRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.automationInternal?.onNodesRemoved && "hasListener" in WEBEXT?.automationInternal?.onNodesRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnNodesRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onNodesRemoved.hasListener);
    },
    "call_HasOnNodesRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onNodesRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnNodesRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onNodesRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTreeChange": (): heap.Ref<boolean> => {
      if (WEBEXT?.automationInternal?.onTreeChange && "addListener" in WEBEXT?.automationInternal?.onTreeChange) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTreeChange": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onTreeChange.addListener);
    },
    "call_OnTreeChange": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onTreeChange.addListener(A.H.get<object>(callback));
    },
    "try_OnTreeChange": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onTreeChange.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffTreeChange": (): heap.Ref<boolean> => {
      if (WEBEXT?.automationInternal?.onTreeChange && "removeListener" in WEBEXT?.automationInternal?.onTreeChange) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffTreeChange": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onTreeChange.removeListener);
    },
    "call_OffTreeChange": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onTreeChange.removeListener(A.H.get<object>(callback));
    },
    "try_OffTreeChange": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onTreeChange.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnTreeChange": (): heap.Ref<boolean> => {
      if (WEBEXT?.automationInternal?.onTreeChange && "hasListener" in WEBEXT?.automationInternal?.onTreeChange) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnTreeChange": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.onTreeChange.hasListener);
    },
    "call_HasOnTreeChange": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automationInternal.onTreeChange.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnTreeChange": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automationInternal.onTreeChange.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_PerformAction": (): heap.Ref<boolean> => {
      if (WEBEXT?.automationInternal && "performAction" in WEBEXT?.automationInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_PerformAction": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automationInternal.performAction);
    },
    "call_PerformAction": (retPtr: Pointer, args: Pointer, opt_args: heap.Ref<object>): void => {
      const args_ffi = {};

      args_ffi["treeID"] = A.load.Ref(args + 0, undefined);
      if (A.load.Bool(args + 16)) {
        args_ffi["automationNodeID"] = A.load.Int32(args + 4);
      }
      args_ffi["actionType"] = A.load.Ref(args + 8, undefined);
      if (A.load.Bool(args + 17)) {
        args_ffi["requestID"] = A.load.Int32(args + 12);
      }

      const _ret = WEBEXT.automationInternal.performAction(args_ffi, A.H.get<object>(opt_args));
    },
    "try_PerformAction": (
      retPtr: Pointer,
      errPtr: Pointer,
      args: Pointer,
      opt_args: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const args_ffi = {};

        args_ffi["treeID"] = A.load.Ref(args + 0, undefined);
        if (A.load.Bool(args + 16)) {
          args_ffi["automationNodeID"] = A.load.Int32(args + 4);
        }
        args_ffi["actionType"] = A.load.Ref(args + 8, undefined);
        if (A.load.Bool(args + 17)) {
          args_ffi["requestID"] = A.load.Int32(args + 12);
        }

        const _ret = WEBEXT.automationInternal.performAction(args_ffi, A.H.get<object>(opt_args));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
