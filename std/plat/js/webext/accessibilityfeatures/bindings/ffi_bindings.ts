import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/accessibilityfeatures", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "get_AnimationPolicy": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "animationPolicy" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.animationPolicy;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AnimationPolicy": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.accessibilityFeatures,
        "animationPolicy",
        A.H.get<object>(val),
        WEBEXT.accessibilityFeatures
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_Autoclick": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "autoclick" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.autoclick;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Autoclick": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.accessibilityFeatures, "autoclick", A.H.get<object>(val), WEBEXT.accessibilityFeatures)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_CaretHighlight": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "caretHighlight" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.caretHighlight;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_CaretHighlight": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.accessibilityFeatures,
        "caretHighlight",
        A.H.get<object>(val),
        WEBEXT.accessibilityFeatures
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_CursorColor": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "cursorColor" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.cursorColor;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_CursorColor": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.accessibilityFeatures,
        "cursorColor",
        A.H.get<object>(val),
        WEBEXT.accessibilityFeatures
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_CursorHighlight": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "cursorHighlight" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.cursorHighlight;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_CursorHighlight": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.accessibilityFeatures,
        "cursorHighlight",
        A.H.get<object>(val),
        WEBEXT.accessibilityFeatures
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_Dictation": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "dictation" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.dictation;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Dictation": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.accessibilityFeatures, "dictation", A.H.get<object>(val), WEBEXT.accessibilityFeatures)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_DockedMagnifier": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "dockedMagnifier" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.dockedMagnifier;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_DockedMagnifier": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.accessibilityFeatures,
        "dockedMagnifier",
        A.H.get<object>(val),
        WEBEXT.accessibilityFeatures
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_FocusHighlight": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "focusHighlight" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.focusHighlight;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_FocusHighlight": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.accessibilityFeatures,
        "focusHighlight",
        A.H.get<object>(val),
        WEBEXT.accessibilityFeatures
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_HighContrast": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "highContrast" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.highContrast;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_HighContrast": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.accessibilityFeatures,
        "highContrast",
        A.H.get<object>(val),
        WEBEXT.accessibilityFeatures
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_LargeCursor": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "largeCursor" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.largeCursor;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_LargeCursor": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.accessibilityFeatures,
        "largeCursor",
        A.H.get<object>(val),
        WEBEXT.accessibilityFeatures
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_ScreenMagnifier": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "screenMagnifier" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.screenMagnifier;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_ScreenMagnifier": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.accessibilityFeatures,
        "screenMagnifier",
        A.H.get<object>(val),
        WEBEXT.accessibilityFeatures
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_SelectToSpeak": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "selectToSpeak" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.selectToSpeak;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_SelectToSpeak": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.accessibilityFeatures,
        "selectToSpeak",
        A.H.get<object>(val),
        WEBEXT.accessibilityFeatures
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_SpokenFeedback": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "spokenFeedback" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.spokenFeedback;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_SpokenFeedback": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.accessibilityFeatures,
        "spokenFeedback",
        A.H.get<object>(val),
        WEBEXT.accessibilityFeatures
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_StickyKeys": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "stickyKeys" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.stickyKeys;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_StickyKeys": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.accessibilityFeatures, "stickyKeys", A.H.get<object>(val), WEBEXT.accessibilityFeatures)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_SwitchAccess": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "switchAccess" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.switchAccess;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_SwitchAccess": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.accessibilityFeatures,
        "switchAccess",
        A.H.get<object>(val),
        WEBEXT.accessibilityFeatures
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_VirtualKeyboard": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityFeatures && "virtualKeyboard" in WEBEXT?.accessibilityFeatures) {
        const val = WEBEXT.accessibilityFeatures.virtualKeyboard;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_VirtualKeyboard": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.accessibilityFeatures,
        "virtualKeyboard",
        A.H.get<object>(val),
        WEBEXT.accessibilityFeatures
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
  };
});
