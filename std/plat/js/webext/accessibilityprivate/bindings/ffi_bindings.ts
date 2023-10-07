import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/accessibilityprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_AcceleratorAction": (ref: heap.Ref<string>): number => {
      const idx = ["focusPreviousPane", "focusNextPane"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_AccessibilityFeature": (ref: heap.Ref<string>): number => {
      const idx = [
        "googleTtsLanguagePacks",
        "dictationContextChecking",
        "chromevoxSettingsMigration",
        "gameFaceIntegration",
        "googleTtsHighQualityVoices",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_AlertInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["message"]);
      }
    },
    "load_AlertInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["message"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_AssistiveTechnologyType": (ref: heap.Ref<string>): number => {
      const idx = ["chromeVox", "selectToSpeak", "switchAccess", "autoClick", "magnifier", "dictation"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_DictationBubbleHintType": (ref: heap.Ref<string>): number => {
      const idx = ["trySaying", "type", "delete", "selectAll", "undo", "help", "unselect", "copy"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_DictationBubbleIconType": (ref: heap.Ref<string>): number => {
      const idx = ["hidden", "standby", "macroSuccess", "macroFail"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DictationBubbleProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 12, false);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Ref(ptr + 0, x["hints"]);
        A.store.Enum(ptr + 4, ["hidden", "standby", "macroSuccess", "macroFail"].indexOf(x["icon"] as string));
        A.store.Ref(ptr + 8, x["text"]);
        A.store.Bool(ptr + 12, x["visible"] ? true : false);
      }
    },
    "load_DictationBubbleProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["hints"] = A.load.Ref(ptr + 0, undefined);
      x["icon"] = A.load.Enum(ptr + 4, ["hidden", "standby", "macroSuccess", "macroFail"]);
      x["text"] = A.load.Ref(ptr + 8, undefined);
      x["visible"] = A.load.Bool(ptr + 12);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DlcType": (ref: heap.Ref<string>): number => {
      const idx = [
        "ttsBnBd",
        "ttsCsCz",
        "ttsDaDk",
        "ttsDeDe",
        "ttsElGr",
        "ttsEnAu",
        "ttsEnGb",
        "ttsEnUs",
        "ttsEsEs",
        "ttsEsUs",
        "ttsFiFi",
        "ttsFilPh",
        "ttsFrFr",
        "ttsHiIn",
        "ttsHuHu",
        "ttsIdId",
        "ttsItIt",
        "ttsJaJp",
        "ttsKmKh",
        "ttsKoKr",
        "ttsNbNo",
        "ttsNeNp",
        "ttsNlNl",
        "ttsPlPl",
        "ttsPtBr",
        "ttsSiLk",
        "ttsSkSk",
        "ttsSvSe",
        "ttsThTh",
        "ttsTrTr",
        "ttsUkUa",
        "ttsViVn",
        "ttsYueHk",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ScreenRect": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_ScreenRect": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["height"] = A.load.Int64(ptr + 0);
      x["left"] = A.load.Int64(ptr + 8);
      x["top"] = A.load.Int64(ptr + 16);
      x["width"] = A.load.Int64(ptr + 24);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_FocusRingStackingOrder": (ref: heap.Ref<string>): number => {
      const idx = ["aboveAccessibilityBubbles", "belowAccessibilityBubbles"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_FocusType": (ref: heap.Ref<string>): number => {
      const idx = ["glow", "solid", "dashed"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_FocusRingInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Enum(ptr + 20, -1);
        A.store.Enum(ptr + 24, -1);
      } else {
        A.store.Bool(ptr + 28, true);
        A.store.Ref(ptr + 0, x["backgroundColor"]);
        A.store.Ref(ptr + 4, x["color"]);
        A.store.Ref(ptr + 8, x["id"]);
        A.store.Ref(ptr + 12, x["rects"]);
        A.store.Ref(ptr + 16, x["secondaryColor"]);
        A.store.Enum(
          ptr + 20,
          ["aboveAccessibilityBubbles", "belowAccessibilityBubbles"].indexOf(x["stackingOrder"] as string)
        );
        A.store.Enum(ptr + 24, ["glow", "solid", "dashed"].indexOf(x["type"] as string));
      }
    },
    "load_FocusRingInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["backgroundColor"] = A.load.Ref(ptr + 0, undefined);
      x["color"] = A.load.Ref(ptr + 4, undefined);
      x["id"] = A.load.Ref(ptr + 8, undefined);
      x["rects"] = A.load.Ref(ptr + 12, undefined);
      x["secondaryColor"] = A.load.Ref(ptr + 16, undefined);
      x["stackingOrder"] = A.load.Enum(ptr + 20, ["aboveAccessibilityBubbles", "belowAccessibilityBubbles"]);
      x["type"] = A.load.Enum(ptr + 24, ["glow", "solid", "dashed"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_Gesture": (ref: heap.Ref<string>): number => {
      const idx = [
        "click",
        "swipeLeft1",
        "swipeUp1",
        "swipeRight1",
        "swipeDown1",
        "swipeLeft2",
        "swipeUp2",
        "swipeRight2",
        "swipeDown2",
        "swipeLeft3",
        "swipeUp3",
        "swipeRight3",
        "swipeDown3",
        "swipeLeft4",
        "swipeUp4",
        "swipeRight4",
        "swipeDown4",
        "tap2",
        "tap3",
        "tap4",
        "touchExplore",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "get_IS_DEFAULT_EVENT_SOURCE_TOUCH": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "IS_DEFAULT_EVENT_SOURCE_TOUCH" in WEBEXT?.accessibilityPrivate) {
        const val = WEBEXT.accessibilityPrivate.IS_DEFAULT_EVENT_SOURCE_TOUCH;
        A.store.Int64(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_IS_DEFAULT_EVENT_SOURCE_TOUCH": (val: number): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.accessibilityPrivate, "IS_DEFAULT_EVENT_SOURCE_TOUCH", val, WEBEXT.accessibilityPrivate)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "constof_MagnifierCommand": (ref: heap.Ref<string>): number => {
      const idx = ["moveStop", "moveUp", "moveDown", "moveLeft", "moveRight"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PointScanPoint": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Float64(ptr + 0, x["x"] === undefined ? 0 : (x["x"] as number));
        A.store.Float64(ptr + 8, x["y"] === undefined ? 0 : (x["y"] as number));
      }
    },
    "load_PointScanPoint": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["x"] = A.load.Float64(ptr + 0);
      x["y"] = A.load.Float64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PointScanState": (ref: heap.Ref<string>): number => {
      const idx = ["start", "stop"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PumpkinData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 52, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Ref(ptr + 48, undefined);
      } else {
        A.store.Bool(ptr + 52, true);
        A.store.Ref(ptr + 0, x["de_de_action_config_binarypb"]);
        A.store.Ref(ptr + 4, x["de_de_pumpkin_config_binarypb"]);
        A.store.Ref(ptr + 8, x["en_us_action_config_binarypb"]);
        A.store.Ref(ptr + 12, x["en_us_pumpkin_config_binarypb"]);
        A.store.Ref(ptr + 16, x["es_es_action_config_binarypb"]);
        A.store.Ref(ptr + 20, x["es_es_pumpkin_config_binarypb"]);
        A.store.Ref(ptr + 24, x["fr_fr_action_config_binarypb"]);
        A.store.Ref(ptr + 28, x["fr_fr_pumpkin_config_binarypb"]);
        A.store.Ref(ptr + 32, x["it_it_action_config_binarypb"]);
        A.store.Ref(ptr + 36, x["it_it_pumpkin_config_binarypb"]);
        A.store.Ref(ptr + 40, x["js_pumpkin_tagger_bin_js"]);
        A.store.Ref(ptr + 44, x["tagger_wasm_main_js"]);
        A.store.Ref(ptr + 48, x["tagger_wasm_main_wasm"]);
      }
    },
    "load_PumpkinData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["de_de_action_config_binarypb"] = A.load.Ref(ptr + 0, undefined);
      x["de_de_pumpkin_config_binarypb"] = A.load.Ref(ptr + 4, undefined);
      x["en_us_action_config_binarypb"] = A.load.Ref(ptr + 8, undefined);
      x["en_us_pumpkin_config_binarypb"] = A.load.Ref(ptr + 12, undefined);
      x["es_es_action_config_binarypb"] = A.load.Ref(ptr + 16, undefined);
      x["es_es_pumpkin_config_binarypb"] = A.load.Ref(ptr + 20, undefined);
      x["fr_fr_action_config_binarypb"] = A.load.Ref(ptr + 24, undefined);
      x["fr_fr_pumpkin_config_binarypb"] = A.load.Ref(ptr + 28, undefined);
      x["it_it_action_config_binarypb"] = A.load.Ref(ptr + 32, undefined);
      x["it_it_pumpkin_config_binarypb"] = A.load.Ref(ptr + 36, undefined);
      x["js_pumpkin_tagger_bin_js"] = A.load.Ref(ptr + 40, undefined);
      x["tagger_wasm_main_js"] = A.load.Ref(ptr + 44, undefined);
      x["tagger_wasm_main_wasm"] = A.load.Ref(ptr + 48, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ScreenPoint": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Int64(ptr + 0, x["x"] === undefined ? 0 : (x["x"] as number));
        A.store.Int64(ptr + 8, x["y"] === undefined ? 0 : (x["y"] as number));
      }
    },
    "load_ScreenPoint": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["x"] = A.load.Int64(ptr + 0);
      x["y"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_SelectToSpeakPanelAction": (ref: heap.Ref<string>): number => {
      const idx = [
        "previousParagraph",
        "previousSentence",
        "pause",
        "resume",
        "nextSentence",
        "nextParagraph",
        "exit",
        "changeSpeed",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SelectToSpeakState": (ref: heap.Ref<string>): number => {
      const idx = ["selecting", "speaking", "inactive"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SetNativeChromeVoxResponse": (ref: heap.Ref<string>): number => {
      const idx = [
        "success",
        "talkbackNotInstalled",
        "windowNotFound",
        "failure",
        "needDeprecationConfirmation",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SwitchAccessBubble": (ref: heap.Ref<string>): number => {
      const idx = ["backButton", "menu"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SwitchAccessCommand": (ref: heap.Ref<string>): number => {
      const idx = ["select", "next", "previous"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SwitchAccessMenuAction": (ref: heap.Ref<string>): number => {
      const idx = [
        "copy",
        "cut",
        "decrement",
        "dictation",
        "endTextSelection",
        "increment",
        "itemScan",
        "jumpToBeginningOfText",
        "jumpToEndOfText",
        "keyboard",
        "leftClick",
        "moveBackwardOneCharOfText",
        "moveBackwardOneWordOfText",
        "moveCursor",
        "moveDownOneLineOfText",
        "moveForwardOneCharOfText",
        "moveForwardOneWordOfText",
        "moveUpOneLineOfText",
        "paste",
        "pointScan",
        "rightClick",
        "scrollDown",
        "scrollLeft",
        "scrollRight",
        "scrollUp",
        "select",
        "settings",
        "startTextSelection",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SyntheticKeyboardModifiers": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 3, false);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Bool(ptr + 4, "alt" in x ? true : false);
        A.store.Bool(ptr + 0, x["alt"] ? true : false);
        A.store.Bool(ptr + 5, "ctrl" in x ? true : false);
        A.store.Bool(ptr + 1, x["ctrl"] ? true : false);
        A.store.Bool(ptr + 6, "search" in x ? true : false);
        A.store.Bool(ptr + 2, x["search"] ? true : false);
        A.store.Bool(ptr + 7, "shift" in x ? true : false);
        A.store.Bool(ptr + 3, x["shift"] ? true : false);
      }
    },
    "load_SyntheticKeyboardModifiers": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["alt"] = A.load.Bool(ptr + 0);
      } else {
        delete x["alt"];
      }
      if (A.load.Bool(ptr + 5)) {
        x["ctrl"] = A.load.Bool(ptr + 1);
      } else {
        delete x["ctrl"];
      }
      if (A.load.Bool(ptr + 6)) {
        x["search"] = A.load.Bool(ptr + 2);
      } else {
        delete x["search"];
      }
      if (A.load.Bool(ptr + 7)) {
        x["shift"] = A.load.Bool(ptr + 3);
      } else {
        delete x["shift"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_SyntheticKeyboardEventType": (ref: heap.Ref<string>): number => {
      const idx = ["keyup", "keydown"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SyntheticKeyboardEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 24, false);
        A.store.Int64(ptr + 0, 0);

        A.store.Bool(ptr + 8 + 8, false);
        A.store.Bool(ptr + 8 + 4, false);
        A.store.Bool(ptr + 8 + 0, false);
        A.store.Bool(ptr + 8 + 5, false);
        A.store.Bool(ptr + 8 + 1, false);
        A.store.Bool(ptr + 8 + 6, false);
        A.store.Bool(ptr + 8 + 2, false);
        A.store.Bool(ptr + 8 + 7, false);
        A.store.Bool(ptr + 8 + 3, false);
        A.store.Enum(ptr + 20, -1);
      } else {
        A.store.Bool(ptr + 24, true);
        A.store.Int64(ptr + 0, x["keyCode"] === undefined ? 0 : (x["keyCode"] as number));

        if (typeof x["modifiers"] === "undefined") {
          A.store.Bool(ptr + 8 + 8, false);
          A.store.Bool(ptr + 8 + 4, false);
          A.store.Bool(ptr + 8 + 0, false);
          A.store.Bool(ptr + 8 + 5, false);
          A.store.Bool(ptr + 8 + 1, false);
          A.store.Bool(ptr + 8 + 6, false);
          A.store.Bool(ptr + 8 + 2, false);
          A.store.Bool(ptr + 8 + 7, false);
          A.store.Bool(ptr + 8 + 3, false);
        } else {
          A.store.Bool(ptr + 8 + 8, true);
          A.store.Bool(ptr + 8 + 4, "alt" in x["modifiers"] ? true : false);
          A.store.Bool(ptr + 8 + 0, x["modifiers"]["alt"] ? true : false);
          A.store.Bool(ptr + 8 + 5, "ctrl" in x["modifiers"] ? true : false);
          A.store.Bool(ptr + 8 + 1, x["modifiers"]["ctrl"] ? true : false);
          A.store.Bool(ptr + 8 + 6, "search" in x["modifiers"] ? true : false);
          A.store.Bool(ptr + 8 + 2, x["modifiers"]["search"] ? true : false);
          A.store.Bool(ptr + 8 + 7, "shift" in x["modifiers"] ? true : false);
          A.store.Bool(ptr + 8 + 3, x["modifiers"]["shift"] ? true : false);
        }
        A.store.Enum(ptr + 20, ["keyup", "keydown"].indexOf(x["type"] as string));
      }
    },
    "load_SyntheticKeyboardEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["keyCode"] = A.load.Int64(ptr + 0);
      if (A.load.Bool(ptr + 8 + 8)) {
        x["modifiers"] = {};
        if (A.load.Bool(ptr + 8 + 4)) {
          x["modifiers"]["alt"] = A.load.Bool(ptr + 8 + 0);
        } else {
          delete x["modifiers"]["alt"];
        }
        if (A.load.Bool(ptr + 8 + 5)) {
          x["modifiers"]["ctrl"] = A.load.Bool(ptr + 8 + 1);
        } else {
          delete x["modifiers"]["ctrl"];
        }
        if (A.load.Bool(ptr + 8 + 6)) {
          x["modifiers"]["search"] = A.load.Bool(ptr + 8 + 2);
        } else {
          delete x["modifiers"]["search"];
        }
        if (A.load.Bool(ptr + 8 + 7)) {
          x["modifiers"]["shift"] = A.load.Bool(ptr + 8 + 3);
        } else {
          delete x["modifiers"]["shift"];
        }
      } else {
        delete x["modifiers"];
      }
      x["type"] = A.load.Enum(ptr + 20, ["keyup", "keydown"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_SyntheticMouseEventButton": (ref: heap.Ref<string>): number => {
      const idx = ["left", "middle", "right", "back", "foward"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SyntheticMouseEventType": (ref: heap.Ref<string>): number => {
      const idx = ["press", "release", "drag", "move", "enter", "exit"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SyntheticMouseEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 33, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 8, -1);
        A.store.Int64(ptr + 16, 0);
        A.store.Int64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 33, true);
        A.store.Enum(ptr + 0, ["left", "middle", "right", "back", "foward"].indexOf(x["mouseButton"] as string));
        A.store.Bool(ptr + 32, "touchAccessibility" in x ? true : false);
        A.store.Bool(ptr + 4, x["touchAccessibility"] ? true : false);
        A.store.Enum(ptr + 8, ["press", "release", "drag", "move", "enter", "exit"].indexOf(x["type"] as string));
        A.store.Int64(ptr + 16, x["x"] === undefined ? 0 : (x["x"] as number));
        A.store.Int64(ptr + 24, x["y"] === undefined ? 0 : (x["y"] as number));
      }
    },
    "load_SyntheticMouseEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["mouseButton"] = A.load.Enum(ptr + 0, ["left", "middle", "right", "back", "foward"]);
      if (A.load.Bool(ptr + 32)) {
        x["touchAccessibility"] = A.load.Bool(ptr + 4);
      } else {
        delete x["touchAccessibility"];
      }
      x["type"] = A.load.Enum(ptr + 8, ["press", "release", "drag", "move", "enter", "exit"]);
      x["x"] = A.load.Int64(ptr + 16);
      x["y"] = A.load.Int64(ptr + 24);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ToastType": (ref: heap.Ref<string>): number => {
      const idx = ["dictationNoFocusedTextField"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_DarkenScreen": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "darkenScreen" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DarkenScreen": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.darkenScreen);
    },
    "call_DarkenScreen": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.accessibilityPrivate.darkenScreen(enabled === A.H.TRUE);
    },
    "try_DarkenScreen": (retPtr: Pointer, errPtr: Pointer, enabled: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.darkenScreen(enabled === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_EnableMouseEvents": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "enableMouseEvents" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EnableMouseEvents": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.enableMouseEvents);
    },
    "call_EnableMouseEvents": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.accessibilityPrivate.enableMouseEvents(enabled === A.H.TRUE);
    },
    "try_EnableMouseEvents": (retPtr: Pointer, errPtr: Pointer, enabled: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.enableMouseEvents(enabled === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ForwardKeyEventsToSwitchAccess": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "forwardKeyEventsToSwitchAccess" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ForwardKeyEventsToSwitchAccess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.forwardKeyEventsToSwitchAccess);
    },
    "call_ForwardKeyEventsToSwitchAccess": (retPtr: Pointer, shouldForward: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.accessibilityPrivate.forwardKeyEventsToSwitchAccess(shouldForward === A.H.TRUE);
    },
    "try_ForwardKeyEventsToSwitchAccess": (
      retPtr: Pointer,
      errPtr: Pointer,
      shouldForward: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.forwardKeyEventsToSwitchAccess(shouldForward === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetBatteryDescription": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "getBatteryDescription" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetBatteryDescription": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.getBatteryDescription);
    },
    "call_GetBatteryDescription": (retPtr: Pointer): void => {
      const _ret = WEBEXT.accessibilityPrivate.getBatteryDescription();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetBatteryDescription": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.getBatteryDescription();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDisplayNameForLocale": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "getDisplayNameForLocale" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDisplayNameForLocale": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.getDisplayNameForLocale);
    },
    "call_GetDisplayNameForLocale": (
      retPtr: Pointer,
      localeCodeToTranslate: heap.Ref<object>,
      displayLocaleCode: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.accessibilityPrivate.getDisplayNameForLocale(
        A.H.get<object>(localeCodeToTranslate),
        A.H.get<object>(displayLocaleCode)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDisplayNameForLocale": (
      retPtr: Pointer,
      errPtr: Pointer,
      localeCodeToTranslate: heap.Ref<object>,
      displayLocaleCode: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.getDisplayNameForLocale(
          A.H.get<object>(localeCodeToTranslate),
          A.H.get<object>(displayLocaleCode)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDlcContents": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "getDlcContents" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDlcContents": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.getDlcContents);
    },
    "call_GetDlcContents": (retPtr: Pointer, dlc: number): void => {
      const _ret = WEBEXT.accessibilityPrivate.getDlcContents(
        dlc > 0 && dlc <= 33
          ? [
              "ttsBnBd",
              "ttsCsCz",
              "ttsDaDk",
              "ttsDeDe",
              "ttsElGr",
              "ttsEnAu",
              "ttsEnGb",
              "ttsEnUs",
              "ttsEsEs",
              "ttsEsUs",
              "ttsFiFi",
              "ttsFilPh",
              "ttsFrFr",
              "ttsHiIn",
              "ttsHuHu",
              "ttsIdId",
              "ttsItIt",
              "ttsJaJp",
              "ttsKmKh",
              "ttsKoKr",
              "ttsNbNo",
              "ttsNeNp",
              "ttsNlNl",
              "ttsPlPl",
              "ttsPtBr",
              "ttsSiLk",
              "ttsSkSk",
              "ttsSvSe",
              "ttsThTh",
              "ttsTrTr",
              "ttsUkUa",
              "ttsViVn",
              "ttsYueHk",
            ][dlc - 1]
          : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDlcContents": (retPtr: Pointer, errPtr: Pointer, dlc: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.getDlcContents(
          dlc > 0 && dlc <= 33
            ? [
                "ttsBnBd",
                "ttsCsCz",
                "ttsDaDk",
                "ttsDeDe",
                "ttsElGr",
                "ttsEnAu",
                "ttsEnGb",
                "ttsEnUs",
                "ttsEsEs",
                "ttsEsUs",
                "ttsFiFi",
                "ttsFilPh",
                "ttsFrFr",
                "ttsHiIn",
                "ttsHuHu",
                "ttsIdId",
                "ttsItIt",
                "ttsJaJp",
                "ttsKmKh",
                "ttsKoKr",
                "ttsNbNo",
                "ttsNeNp",
                "ttsNlNl",
                "ttsPlPl",
                "ttsPtBr",
                "ttsSiLk",
                "ttsSkSk",
                "ttsSvSe",
                "ttsThTh",
                "ttsTrTr",
                "ttsUkUa",
                "ttsViVn",
                "ttsYueHk",
              ][dlc - 1]
            : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetLocalizedDomKeyStringForKeyCode": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "getLocalizedDomKeyStringForKeyCode" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetLocalizedDomKeyStringForKeyCode": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.getLocalizedDomKeyStringForKeyCode);
    },
    "call_GetLocalizedDomKeyStringForKeyCode": (retPtr: Pointer, keyCode: number): void => {
      const _ret = WEBEXT.accessibilityPrivate.getLocalizedDomKeyStringForKeyCode(keyCode);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetLocalizedDomKeyStringForKeyCode": (
      retPtr: Pointer,
      errPtr: Pointer,
      keyCode: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.getLocalizedDomKeyStringForKeyCode(keyCode);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HandleScrollableBoundsForPointFound": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "handleScrollableBoundsForPointFound" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HandleScrollableBoundsForPointFound": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.handleScrollableBoundsForPointFound);
    },
    "call_HandleScrollableBoundsForPointFound": (retPtr: Pointer, rect: Pointer): void => {
      const rect_ffi = {};

      rect_ffi["height"] = A.load.Int64(rect + 0);
      rect_ffi["left"] = A.load.Int64(rect + 8);
      rect_ffi["top"] = A.load.Int64(rect + 16);
      rect_ffi["width"] = A.load.Int64(rect + 24);

      const _ret = WEBEXT.accessibilityPrivate.handleScrollableBoundsForPointFound(rect_ffi);
    },
    "try_HandleScrollableBoundsForPointFound": (retPtr: Pointer, errPtr: Pointer, rect: Pointer): heap.Ref<boolean> => {
      try {
        const rect_ffi = {};

        rect_ffi["height"] = A.load.Int64(rect + 0);
        rect_ffi["left"] = A.load.Int64(rect + 8);
        rect_ffi["top"] = A.load.Int64(rect + 16);
        rect_ffi["width"] = A.load.Int64(rect + 24);

        const _ret = WEBEXT.accessibilityPrivate.handleScrollableBoundsForPointFound(rect_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InstallPumpkinForDictation": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "installPumpkinForDictation" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InstallPumpkinForDictation": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.installPumpkinForDictation);
    },
    "call_InstallPumpkinForDictation": (retPtr: Pointer): void => {
      const _ret = WEBEXT.accessibilityPrivate.installPumpkinForDictation();
      A.store.Ref(retPtr, _ret);
    },
    "try_InstallPumpkinForDictation": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.installPumpkinForDictation();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsFeatureEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "isFeatureEnabled" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsFeatureEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.isFeatureEnabled);
    },
    "call_IsFeatureEnabled": (retPtr: Pointer, feature: number): void => {
      const _ret = WEBEXT.accessibilityPrivate.isFeatureEnabled(
        feature > 0 && feature <= 5
          ? [
              "googleTtsLanguagePacks",
              "dictationContextChecking",
              "chromevoxSettingsMigration",
              "gameFaceIntegration",
              "googleTtsHighQualityVoices",
            ][feature - 1]
          : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_IsFeatureEnabled": (retPtr: Pointer, errPtr: Pointer, feature: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.isFeatureEnabled(
          feature > 0 && feature <= 5
            ? [
                "googleTtsLanguagePacks",
                "dictationContextChecking",
                "chromevoxSettingsMigration",
                "gameFaceIntegration",
                "googleTtsHighQualityVoices",
              ][feature - 1]
            : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsLacrosPrimary": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "isLacrosPrimary" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsLacrosPrimary": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.isLacrosPrimary);
    },
    "call_IsLacrosPrimary": (retPtr: Pointer): void => {
      const _ret = WEBEXT.accessibilityPrivate.isLacrosPrimary();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsLacrosPrimary": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.isLacrosPrimary();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_MagnifierCenterOnPoint": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "magnifierCenterOnPoint" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MagnifierCenterOnPoint": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.magnifierCenterOnPoint);
    },
    "call_MagnifierCenterOnPoint": (retPtr: Pointer, point: Pointer): void => {
      const point_ffi = {};

      point_ffi["x"] = A.load.Int64(point + 0);
      point_ffi["y"] = A.load.Int64(point + 8);

      const _ret = WEBEXT.accessibilityPrivate.magnifierCenterOnPoint(point_ffi);
    },
    "try_MagnifierCenterOnPoint": (retPtr: Pointer, errPtr: Pointer, point: Pointer): heap.Ref<boolean> => {
      try {
        const point_ffi = {};

        point_ffi["x"] = A.load.Int64(point + 0);
        point_ffi["y"] = A.load.Int64(point + 8);

        const _ret = WEBEXT.accessibilityPrivate.magnifierCenterOnPoint(point_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_MoveMagnifierToRect": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "moveMagnifierToRect" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MoveMagnifierToRect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.moveMagnifierToRect);
    },
    "call_MoveMagnifierToRect": (retPtr: Pointer, rect: Pointer): void => {
      const rect_ffi = {};

      rect_ffi["height"] = A.load.Int64(rect + 0);
      rect_ffi["left"] = A.load.Int64(rect + 8);
      rect_ffi["top"] = A.load.Int64(rect + 16);
      rect_ffi["width"] = A.load.Int64(rect + 24);

      const _ret = WEBEXT.accessibilityPrivate.moveMagnifierToRect(rect_ffi);
    },
    "try_MoveMagnifierToRect": (retPtr: Pointer, errPtr: Pointer, rect: Pointer): heap.Ref<boolean> => {
      try {
        const rect_ffi = {};

        rect_ffi["height"] = A.load.Int64(rect + 0);
        rect_ffi["left"] = A.load.Int64(rect + 8);
        rect_ffi["top"] = A.load.Int64(rect + 16);
        rect_ffi["width"] = A.load.Int64(rect + 24);

        const _ret = WEBEXT.accessibilityPrivate.moveMagnifierToRect(rect_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAccessibilityGesture": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onAccessibilityGesture &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onAccessibilityGesture
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAccessibilityGesture": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onAccessibilityGesture.addListener);
    },
    "call_OnAccessibilityGesture": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onAccessibilityGesture.addListener(A.H.get<object>(callback));
    },
    "try_OnAccessibilityGesture": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onAccessibilityGesture.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAccessibilityGesture": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onAccessibilityGesture &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onAccessibilityGesture
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAccessibilityGesture": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onAccessibilityGesture.removeListener);
    },
    "call_OffAccessibilityGesture": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onAccessibilityGesture.removeListener(A.H.get<object>(callback));
    },
    "try_OffAccessibilityGesture": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onAccessibilityGesture.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAccessibilityGesture": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onAccessibilityGesture &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onAccessibilityGesture
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAccessibilityGesture": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onAccessibilityGesture.hasListener);
    },
    "call_HasOnAccessibilityGesture": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onAccessibilityGesture.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAccessibilityGesture": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onAccessibilityGesture.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAnnounceForAccessibility": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onAnnounceForAccessibility &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onAnnounceForAccessibility
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAnnounceForAccessibility": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.addListener);
    },
    "call_OnAnnounceForAccessibility": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.addListener(A.H.get<object>(callback));
    },
    "try_OnAnnounceForAccessibility": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAnnounceForAccessibility": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onAnnounceForAccessibility &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onAnnounceForAccessibility
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAnnounceForAccessibility": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.removeListener);
    },
    "call_OffAnnounceForAccessibility": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.removeListener(A.H.get<object>(callback));
    },
    "try_OffAnnounceForAccessibility": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAnnounceForAccessibility": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onAnnounceForAccessibility &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onAnnounceForAccessibility
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAnnounceForAccessibility": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.hasListener);
    },
    "call_HasOnAnnounceForAccessibility": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAnnounceForAccessibility": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCustomSpokenFeedbackToggled": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onCustomSpokenFeedbackToggled &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onCustomSpokenFeedbackToggled
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCustomSpokenFeedbackToggled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.addListener);
    },
    "call_OnCustomSpokenFeedbackToggled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.addListener(A.H.get<object>(callback));
    },
    "try_OnCustomSpokenFeedbackToggled": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCustomSpokenFeedbackToggled": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onCustomSpokenFeedbackToggled &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onCustomSpokenFeedbackToggled
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCustomSpokenFeedbackToggled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.removeListener);
    },
    "call_OffCustomSpokenFeedbackToggled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.removeListener(A.H.get<object>(callback));
    },
    "try_OffCustomSpokenFeedbackToggled": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCustomSpokenFeedbackToggled": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onCustomSpokenFeedbackToggled &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onCustomSpokenFeedbackToggled
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCustomSpokenFeedbackToggled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.hasListener);
    },
    "call_HasOnCustomSpokenFeedbackToggled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCustomSpokenFeedbackToggled": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnIntroduceChromeVox": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onIntroduceChromeVox &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onIntroduceChromeVox
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnIntroduceChromeVox": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onIntroduceChromeVox.addListener);
    },
    "call_OnIntroduceChromeVox": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onIntroduceChromeVox.addListener(A.H.get<object>(callback));
    },
    "try_OnIntroduceChromeVox": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onIntroduceChromeVox.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffIntroduceChromeVox": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onIntroduceChromeVox &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onIntroduceChromeVox
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffIntroduceChromeVox": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onIntroduceChromeVox.removeListener);
    },
    "call_OffIntroduceChromeVox": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onIntroduceChromeVox.removeListener(A.H.get<object>(callback));
    },
    "try_OffIntroduceChromeVox": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onIntroduceChromeVox.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnIntroduceChromeVox": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onIntroduceChromeVox &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onIntroduceChromeVox
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnIntroduceChromeVox": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onIntroduceChromeVox.hasListener);
    },
    "call_HasOnIntroduceChromeVox": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onIntroduceChromeVox.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnIntroduceChromeVox": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onIntroduceChromeVox.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMagnifierBoundsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onMagnifierBoundsChanged &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onMagnifierBoundsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMagnifierBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.addListener);
    },
    "call_OnMagnifierBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnMagnifierBoundsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMagnifierBoundsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onMagnifierBoundsChanged &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onMagnifierBoundsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMagnifierBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.removeListener);
    },
    "call_OffMagnifierBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffMagnifierBoundsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMagnifierBoundsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onMagnifierBoundsChanged &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onMagnifierBoundsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMagnifierBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.hasListener);
    },
    "call_HasOnMagnifierBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMagnifierBoundsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMagnifierCommand": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onMagnifierCommand &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onMagnifierCommand
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMagnifierCommand": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onMagnifierCommand.addListener);
    },
    "call_OnMagnifierCommand": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onMagnifierCommand.addListener(A.H.get<object>(callback));
    },
    "try_OnMagnifierCommand": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onMagnifierCommand.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMagnifierCommand": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onMagnifierCommand &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onMagnifierCommand
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMagnifierCommand": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onMagnifierCommand.removeListener);
    },
    "call_OffMagnifierCommand": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onMagnifierCommand.removeListener(A.H.get<object>(callback));
    },
    "try_OffMagnifierCommand": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onMagnifierCommand.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMagnifierCommand": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onMagnifierCommand &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onMagnifierCommand
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMagnifierCommand": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onMagnifierCommand.hasListener);
    },
    "call_HasOnMagnifierCommand": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onMagnifierCommand.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMagnifierCommand": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onMagnifierCommand.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPointScanSet": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onPointScanSet &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onPointScanSet
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPointScanSet": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onPointScanSet.addListener);
    },
    "call_OnPointScanSet": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onPointScanSet.addListener(A.H.get<object>(callback));
    },
    "try_OnPointScanSet": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onPointScanSet.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPointScanSet": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onPointScanSet &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onPointScanSet
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPointScanSet": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onPointScanSet.removeListener);
    },
    "call_OffPointScanSet": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onPointScanSet.removeListener(A.H.get<object>(callback));
    },
    "try_OffPointScanSet": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onPointScanSet.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPointScanSet": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onPointScanSet &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onPointScanSet
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPointScanSet": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onPointScanSet.hasListener);
    },
    "call_HasOnPointScanSet": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onPointScanSet.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPointScanSet": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onPointScanSet.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnScrollableBoundsForPointRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onScrollableBoundsForPointRequested &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onScrollableBoundsForPointRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnScrollableBoundsForPointRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.addListener);
    },
    "call_OnScrollableBoundsForPointRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.addListener(
        A.H.get<object>(callback)
      );
    },
    "try_OnScrollableBoundsForPointRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.addListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffScrollableBoundsForPointRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onScrollableBoundsForPointRequested &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onScrollableBoundsForPointRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffScrollableBoundsForPointRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.removeListener);
    },
    "call_OffScrollableBoundsForPointRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.removeListener(
        A.H.get<object>(callback)
      );
    },
    "try_OffScrollableBoundsForPointRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnScrollableBoundsForPointRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onScrollableBoundsForPointRequested &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onScrollableBoundsForPointRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnScrollableBoundsForPointRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.hasListener);
    },
    "call_HasOnScrollableBoundsForPointRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.hasListener(
        A.H.get<object>(callback)
      );
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnScrollableBoundsForPointRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.hasListener(
          A.H.get<object>(callback)
        );
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSelectToSpeakContextMenuClicked": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSelectToSpeakContextMenuClicked &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onSelectToSpeakContextMenuClicked
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSelectToSpeakContextMenuClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.addListener);
    },
    "call_OnSelectToSpeakContextMenuClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.addListener(A.H.get<object>(callback));
    },
    "try_OnSelectToSpeakContextMenuClicked": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.addListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSelectToSpeakContextMenuClicked": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSelectToSpeakContextMenuClicked &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onSelectToSpeakContextMenuClicked
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSelectToSpeakContextMenuClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.removeListener);
    },
    "call_OffSelectToSpeakContextMenuClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.removeListener(
        A.H.get<object>(callback)
      );
    },
    "try_OffSelectToSpeakContextMenuClicked": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSelectToSpeakContextMenuClicked": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSelectToSpeakContextMenuClicked &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onSelectToSpeakContextMenuClicked
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSelectToSpeakContextMenuClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.hasListener);
    },
    "call_HasOnSelectToSpeakContextMenuClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSelectToSpeakContextMenuClicked": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.hasListener(
          A.H.get<object>(callback)
        );
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSelectToSpeakKeysPressedChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSelectToSpeakKeysPressedChanged &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onSelectToSpeakKeysPressedChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSelectToSpeakKeysPressedChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.addListener);
    },
    "call_OnSelectToSpeakKeysPressedChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnSelectToSpeakKeysPressedChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.addListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSelectToSpeakKeysPressedChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSelectToSpeakKeysPressedChanged &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onSelectToSpeakKeysPressedChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSelectToSpeakKeysPressedChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.removeListener);
    },
    "call_OffSelectToSpeakKeysPressedChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.removeListener(
        A.H.get<object>(callback)
      );
    },
    "try_OffSelectToSpeakKeysPressedChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSelectToSpeakKeysPressedChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSelectToSpeakKeysPressedChanged &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onSelectToSpeakKeysPressedChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSelectToSpeakKeysPressedChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.hasListener);
    },
    "call_HasOnSelectToSpeakKeysPressedChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSelectToSpeakKeysPressedChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.hasListener(
          A.H.get<object>(callback)
        );
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSelectToSpeakMouseChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSelectToSpeakMouseChanged &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onSelectToSpeakMouseChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSelectToSpeakMouseChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.addListener);
    },
    "call_OnSelectToSpeakMouseChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnSelectToSpeakMouseChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSelectToSpeakMouseChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSelectToSpeakMouseChanged &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onSelectToSpeakMouseChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSelectToSpeakMouseChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.removeListener);
    },
    "call_OffSelectToSpeakMouseChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffSelectToSpeakMouseChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSelectToSpeakMouseChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSelectToSpeakMouseChanged &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onSelectToSpeakMouseChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSelectToSpeakMouseChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.hasListener);
    },
    "call_HasOnSelectToSpeakMouseChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSelectToSpeakMouseChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSelectToSpeakPanelAction": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSelectToSpeakPanelAction &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onSelectToSpeakPanelAction
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSelectToSpeakPanelAction": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.addListener);
    },
    "call_OnSelectToSpeakPanelAction": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.addListener(A.H.get<object>(callback));
    },
    "try_OnSelectToSpeakPanelAction": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSelectToSpeakPanelAction": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSelectToSpeakPanelAction &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onSelectToSpeakPanelAction
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSelectToSpeakPanelAction": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.removeListener);
    },
    "call_OffSelectToSpeakPanelAction": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.removeListener(A.H.get<object>(callback));
    },
    "try_OffSelectToSpeakPanelAction": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSelectToSpeakPanelAction": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSelectToSpeakPanelAction &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onSelectToSpeakPanelAction
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSelectToSpeakPanelAction": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.hasListener);
    },
    "call_HasOnSelectToSpeakPanelAction": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSelectToSpeakPanelAction": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSelectToSpeakStateChangeRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSelectToSpeakStateChangeRequested &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onSelectToSpeakStateChangeRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSelectToSpeakStateChangeRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.addListener);
    },
    "call_OnSelectToSpeakStateChangeRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.addListener(
        A.H.get<object>(callback)
      );
    },
    "try_OnSelectToSpeakStateChangeRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.addListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSelectToSpeakStateChangeRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSelectToSpeakStateChangeRequested &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onSelectToSpeakStateChangeRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSelectToSpeakStateChangeRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.removeListener);
    },
    "call_OffSelectToSpeakStateChangeRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.removeListener(
        A.H.get<object>(callback)
      );
    },
    "try_OffSelectToSpeakStateChangeRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSelectToSpeakStateChangeRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSelectToSpeakStateChangeRequested &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onSelectToSpeakStateChangeRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSelectToSpeakStateChangeRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.hasListener);
    },
    "call_HasOnSelectToSpeakStateChangeRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.hasListener(
        A.H.get<object>(callback)
      );
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSelectToSpeakStateChangeRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.hasListener(
          A.H.get<object>(callback)
        );
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnShowChromeVoxTutorial": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onShowChromeVoxTutorial &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onShowChromeVoxTutorial
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnShowChromeVoxTutorial": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.addListener);
    },
    "call_OnShowChromeVoxTutorial": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.addListener(A.H.get<object>(callback));
    },
    "try_OnShowChromeVoxTutorial": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffShowChromeVoxTutorial": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onShowChromeVoxTutorial &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onShowChromeVoxTutorial
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffShowChromeVoxTutorial": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.removeListener);
    },
    "call_OffShowChromeVoxTutorial": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.removeListener(A.H.get<object>(callback));
    },
    "try_OffShowChromeVoxTutorial": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnShowChromeVoxTutorial": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onShowChromeVoxTutorial &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onShowChromeVoxTutorial
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnShowChromeVoxTutorial": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.hasListener);
    },
    "call_HasOnShowChromeVoxTutorial": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnShowChromeVoxTutorial": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSwitchAccessCommand": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSwitchAccessCommand &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onSwitchAccessCommand
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSwitchAccessCommand": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSwitchAccessCommand.addListener);
    },
    "call_OnSwitchAccessCommand": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSwitchAccessCommand.addListener(A.H.get<object>(callback));
    },
    "try_OnSwitchAccessCommand": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSwitchAccessCommand.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSwitchAccessCommand": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSwitchAccessCommand &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onSwitchAccessCommand
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSwitchAccessCommand": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSwitchAccessCommand.removeListener);
    },
    "call_OffSwitchAccessCommand": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSwitchAccessCommand.removeListener(A.H.get<object>(callback));
    },
    "try_OffSwitchAccessCommand": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSwitchAccessCommand.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSwitchAccessCommand": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onSwitchAccessCommand &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onSwitchAccessCommand
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSwitchAccessCommand": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onSwitchAccessCommand.hasListener);
    },
    "call_HasOnSwitchAccessCommand": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onSwitchAccessCommand.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSwitchAccessCommand": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onSwitchAccessCommand.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnToggleDictation": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onToggleDictation &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onToggleDictation
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnToggleDictation": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onToggleDictation.addListener);
    },
    "call_OnToggleDictation": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onToggleDictation.addListener(A.H.get<object>(callback));
    },
    "try_OnToggleDictation": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onToggleDictation.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffToggleDictation": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onToggleDictation &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onToggleDictation
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffToggleDictation": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onToggleDictation.removeListener);
    },
    "call_OffToggleDictation": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onToggleDictation.removeListener(A.H.get<object>(callback));
    },
    "try_OffToggleDictation": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onToggleDictation.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnToggleDictation": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onToggleDictation &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onToggleDictation
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnToggleDictation": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onToggleDictation.hasListener);
    },
    "call_HasOnToggleDictation": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onToggleDictation.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnToggleDictation": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onToggleDictation.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTwoFingerTouchStart": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onTwoFingerTouchStart &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onTwoFingerTouchStart
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTwoFingerTouchStart": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.addListener);
    },
    "call_OnTwoFingerTouchStart": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.addListener(A.H.get<object>(callback));
    },
    "try_OnTwoFingerTouchStart": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffTwoFingerTouchStart": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onTwoFingerTouchStart &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onTwoFingerTouchStart
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffTwoFingerTouchStart": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.removeListener);
    },
    "call_OffTwoFingerTouchStart": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.removeListener(A.H.get<object>(callback));
    },
    "try_OffTwoFingerTouchStart": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnTwoFingerTouchStart": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onTwoFingerTouchStart &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onTwoFingerTouchStart
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnTwoFingerTouchStart": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.hasListener);
    },
    "call_HasOnTwoFingerTouchStart": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnTwoFingerTouchStart": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTwoFingerTouchStop": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onTwoFingerTouchStop &&
        "addListener" in WEBEXT?.accessibilityPrivate?.onTwoFingerTouchStop
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTwoFingerTouchStop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.addListener);
    },
    "call_OnTwoFingerTouchStop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.addListener(A.H.get<object>(callback));
    },
    "try_OnTwoFingerTouchStop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffTwoFingerTouchStop": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onTwoFingerTouchStop &&
        "removeListener" in WEBEXT?.accessibilityPrivate?.onTwoFingerTouchStop
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffTwoFingerTouchStop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.removeListener);
    },
    "call_OffTwoFingerTouchStop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.removeListener(A.H.get<object>(callback));
    },
    "try_OffTwoFingerTouchStop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnTwoFingerTouchStop": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.accessibilityPrivate?.onTwoFingerTouchStop &&
        "hasListener" in WEBEXT?.accessibilityPrivate?.onTwoFingerTouchStop
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnTwoFingerTouchStop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.hasListener);
    },
    "call_HasOnTwoFingerTouchStop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnTwoFingerTouchStop": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenSettingsSubpage": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "openSettingsSubpage" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenSettingsSubpage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.openSettingsSubpage);
    },
    "call_OpenSettingsSubpage": (retPtr: Pointer, subpage: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.openSettingsSubpage(A.H.get<object>(subpage));
    },
    "try_OpenSettingsSubpage": (retPtr: Pointer, errPtr: Pointer, subpage: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.openSettingsSubpage(A.H.get<object>(subpage));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_PerformAcceleratorAction": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "performAcceleratorAction" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_PerformAcceleratorAction": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.performAcceleratorAction);
    },
    "call_PerformAcceleratorAction": (retPtr: Pointer, acceleratorAction: number): void => {
      const _ret = WEBEXT.accessibilityPrivate.performAcceleratorAction(
        acceleratorAction > 0 && acceleratorAction <= 2
          ? ["focusPreviousPane", "focusNextPane"][acceleratorAction - 1]
          : undefined
      );
    },
    "try_PerformAcceleratorAction": (
      retPtr: Pointer,
      errPtr: Pointer,
      acceleratorAction: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.performAcceleratorAction(
          acceleratorAction > 0 && acceleratorAction <= 2
            ? ["focusPreviousPane", "focusNextPane"][acceleratorAction - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendSyntheticKeyEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "sendSyntheticKeyEvent" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendSyntheticKeyEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.sendSyntheticKeyEvent);
    },
    "call_SendSyntheticKeyEvent": (retPtr: Pointer, keyEvent: Pointer, useRewriters: heap.Ref<boolean>): void => {
      const keyEvent_ffi = {};

      keyEvent_ffi["keyCode"] = A.load.Int64(keyEvent + 0);
      if (A.load.Bool(keyEvent + 8 + 8)) {
        keyEvent_ffi["modifiers"] = {};
        if (A.load.Bool(keyEvent + 8 + 4)) {
          keyEvent_ffi["modifiers"]["alt"] = A.load.Bool(keyEvent + 8 + 0);
        }
        if (A.load.Bool(keyEvent + 8 + 5)) {
          keyEvent_ffi["modifiers"]["ctrl"] = A.load.Bool(keyEvent + 8 + 1);
        }
        if (A.load.Bool(keyEvent + 8 + 6)) {
          keyEvent_ffi["modifiers"]["search"] = A.load.Bool(keyEvent + 8 + 2);
        }
        if (A.load.Bool(keyEvent + 8 + 7)) {
          keyEvent_ffi["modifiers"]["shift"] = A.load.Bool(keyEvent + 8 + 3);
        }
      }
      keyEvent_ffi["type"] = A.load.Enum(keyEvent + 20, ["keyup", "keydown"]);

      const _ret = WEBEXT.accessibilityPrivate.sendSyntheticKeyEvent(keyEvent_ffi, useRewriters === A.H.TRUE);
    },
    "try_SendSyntheticKeyEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      keyEvent: Pointer,
      useRewriters: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const keyEvent_ffi = {};

        keyEvent_ffi["keyCode"] = A.load.Int64(keyEvent + 0);
        if (A.load.Bool(keyEvent + 8 + 8)) {
          keyEvent_ffi["modifiers"] = {};
          if (A.load.Bool(keyEvent + 8 + 4)) {
            keyEvent_ffi["modifiers"]["alt"] = A.load.Bool(keyEvent + 8 + 0);
          }
          if (A.load.Bool(keyEvent + 8 + 5)) {
            keyEvent_ffi["modifiers"]["ctrl"] = A.load.Bool(keyEvent + 8 + 1);
          }
          if (A.load.Bool(keyEvent + 8 + 6)) {
            keyEvent_ffi["modifiers"]["search"] = A.load.Bool(keyEvent + 8 + 2);
          }
          if (A.load.Bool(keyEvent + 8 + 7)) {
            keyEvent_ffi["modifiers"]["shift"] = A.load.Bool(keyEvent + 8 + 3);
          }
        }
        keyEvent_ffi["type"] = A.load.Enum(keyEvent + 20, ["keyup", "keydown"]);

        const _ret = WEBEXT.accessibilityPrivate.sendSyntheticKeyEvent(keyEvent_ffi, useRewriters === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendSyntheticMouseEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "sendSyntheticMouseEvent" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendSyntheticMouseEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.sendSyntheticMouseEvent);
    },
    "call_SendSyntheticMouseEvent": (retPtr: Pointer, mouseEvent: Pointer): void => {
      const mouseEvent_ffi = {};

      mouseEvent_ffi["mouseButton"] = A.load.Enum(mouseEvent + 0, ["left", "middle", "right", "back", "foward"]);
      if (A.load.Bool(mouseEvent + 32)) {
        mouseEvent_ffi["touchAccessibility"] = A.load.Bool(mouseEvent + 4);
      }
      mouseEvent_ffi["type"] = A.load.Enum(mouseEvent + 8, ["press", "release", "drag", "move", "enter", "exit"]);
      mouseEvent_ffi["x"] = A.load.Int64(mouseEvent + 16);
      mouseEvent_ffi["y"] = A.load.Int64(mouseEvent + 24);

      const _ret = WEBEXT.accessibilityPrivate.sendSyntheticMouseEvent(mouseEvent_ffi);
    },
    "try_SendSyntheticMouseEvent": (retPtr: Pointer, errPtr: Pointer, mouseEvent: Pointer): heap.Ref<boolean> => {
      try {
        const mouseEvent_ffi = {};

        mouseEvent_ffi["mouseButton"] = A.load.Enum(mouseEvent + 0, ["left", "middle", "right", "back", "foward"]);
        if (A.load.Bool(mouseEvent + 32)) {
          mouseEvent_ffi["touchAccessibility"] = A.load.Bool(mouseEvent + 4);
        }
        mouseEvent_ffi["type"] = A.load.Enum(mouseEvent + 8, ["press", "release", "drag", "move", "enter", "exit"]);
        mouseEvent_ffi["x"] = A.load.Int64(mouseEvent + 16);
        mouseEvent_ffi["y"] = A.load.Int64(mouseEvent + 24);

        const _ret = WEBEXT.accessibilityPrivate.sendSyntheticMouseEvent(mouseEvent_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetFocusRings": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "setFocusRings" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetFocusRings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.setFocusRings);
    },
    "call_SetFocusRings": (retPtr: Pointer, focusRings: heap.Ref<object>, atType: number): void => {
      const _ret = WEBEXT.accessibilityPrivate.setFocusRings(
        A.H.get<object>(focusRings),
        atType > 0 && atType <= 6
          ? ["chromeVox", "selectToSpeak", "switchAccess", "autoClick", "magnifier", "dictation"][atType - 1]
          : undefined
      );
    },
    "try_SetFocusRings": (
      retPtr: Pointer,
      errPtr: Pointer,
      focusRings: heap.Ref<object>,
      atType: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.setFocusRings(
          A.H.get<object>(focusRings),
          atType > 0 && atType <= 6
            ? ["chromeVox", "selectToSpeak", "switchAccess", "autoClick", "magnifier", "dictation"][atType - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetHighlights": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "setHighlights" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetHighlights": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.setHighlights);
    },
    "call_SetHighlights": (retPtr: Pointer, rects: heap.Ref<object>, color: heap.Ref<object>): void => {
      const _ret = WEBEXT.accessibilityPrivate.setHighlights(A.H.get<object>(rects), A.H.get<object>(color));
    },
    "try_SetHighlights": (
      retPtr: Pointer,
      errPtr: Pointer,
      rects: heap.Ref<object>,
      color: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.setHighlights(A.H.get<object>(rects), A.H.get<object>(color));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetKeyboardListener": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "setKeyboardListener" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetKeyboardListener": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.setKeyboardListener);
    },
    "call_SetKeyboardListener": (retPtr: Pointer, enabled: heap.Ref<boolean>, capture: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.accessibilityPrivate.setKeyboardListener(enabled === A.H.TRUE, capture === A.H.TRUE);
    },
    "try_SetKeyboardListener": (
      retPtr: Pointer,
      errPtr: Pointer,
      enabled: heap.Ref<boolean>,
      capture: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.setKeyboardListener(enabled === A.H.TRUE, capture === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetNativeAccessibilityEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "setNativeAccessibilityEnabled" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetNativeAccessibilityEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.setNativeAccessibilityEnabled);
    },
    "call_SetNativeAccessibilityEnabled": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.accessibilityPrivate.setNativeAccessibilityEnabled(enabled === A.H.TRUE);
    },
    "try_SetNativeAccessibilityEnabled": (
      retPtr: Pointer,
      errPtr: Pointer,
      enabled: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.setNativeAccessibilityEnabled(enabled === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetNativeChromeVoxArcSupportForCurrentApp": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "setNativeChromeVoxArcSupportForCurrentApp" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetNativeChromeVoxArcSupportForCurrentApp": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.setNativeChromeVoxArcSupportForCurrentApp);
    },
    "call_SetNativeChromeVoxArcSupportForCurrentApp": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.accessibilityPrivate.setNativeChromeVoxArcSupportForCurrentApp(enabled === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetNativeChromeVoxArcSupportForCurrentApp": (
      retPtr: Pointer,
      errPtr: Pointer,
      enabled: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.setNativeChromeVoxArcSupportForCurrentApp(enabled === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPointScanState": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "setPointScanState" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPointScanState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.setPointScanState);
    },
    "call_SetPointScanState": (retPtr: Pointer, state: number): void => {
      const _ret = WEBEXT.accessibilityPrivate.setPointScanState(
        state > 0 && state <= 2 ? ["start", "stop"][state - 1] : undefined
      );
    },
    "try_SetPointScanState": (retPtr: Pointer, errPtr: Pointer, state: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.setPointScanState(
          state > 0 && state <= 2 ? ["start", "stop"][state - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetSelectToSpeakState": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "setSelectToSpeakState" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetSelectToSpeakState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.setSelectToSpeakState);
    },
    "call_SetSelectToSpeakState": (retPtr: Pointer, state: number): void => {
      const _ret = WEBEXT.accessibilityPrivate.setSelectToSpeakState(
        state > 0 && state <= 3 ? ["selecting", "speaking", "inactive"][state - 1] : undefined
      );
    },
    "try_SetSelectToSpeakState": (retPtr: Pointer, errPtr: Pointer, state: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.setSelectToSpeakState(
          state > 0 && state <= 3 ? ["selecting", "speaking", "inactive"][state - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetVirtualKeyboardVisible": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "setVirtualKeyboardVisible" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetVirtualKeyboardVisible": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.setVirtualKeyboardVisible);
    },
    "call_SetVirtualKeyboardVisible": (retPtr: Pointer, isVisible: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.accessibilityPrivate.setVirtualKeyboardVisible(isVisible === A.H.TRUE);
    },
    "try_SetVirtualKeyboardVisible": (
      retPtr: Pointer,
      errPtr: Pointer,
      isVisible: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.setVirtualKeyboardVisible(isVisible === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ShowConfirmationDialog": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "showConfirmationDialog" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ShowConfirmationDialog": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.showConfirmationDialog);
    },
    "call_ShowConfirmationDialog": (
      retPtr: Pointer,
      title: heap.Ref<object>,
      description: heap.Ref<object>,
      cancelName: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.accessibilityPrivate.showConfirmationDialog(
        A.H.get<object>(title),
        A.H.get<object>(description),
        A.H.get<object>(cancelName)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_ShowConfirmationDialog": (
      retPtr: Pointer,
      errPtr: Pointer,
      title: heap.Ref<object>,
      description: heap.Ref<object>,
      cancelName: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.showConfirmationDialog(
          A.H.get<object>(title),
          A.H.get<object>(description),
          A.H.get<object>(cancelName)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ShowToast": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "showToast" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ShowToast": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.showToast);
    },
    "call_ShowToast": (retPtr: Pointer, type: number): void => {
      const _ret = WEBEXT.accessibilityPrivate.showToast(
        type > 0 && type <= 1 ? ["dictationNoFocusedTextField"][type - 1] : undefined
      );
    },
    "try_ShowToast": (retPtr: Pointer, errPtr: Pointer, type: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.showToast(
          type > 0 && type <= 1 ? ["dictationNoFocusedTextField"][type - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SilenceSpokenFeedback": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "silenceSpokenFeedback" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SilenceSpokenFeedback": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.silenceSpokenFeedback);
    },
    "call_SilenceSpokenFeedback": (retPtr: Pointer): void => {
      const _ret = WEBEXT.accessibilityPrivate.silenceSpokenFeedback();
    },
    "try_SilenceSpokenFeedback": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.silenceSpokenFeedback();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ToggleDictation": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "toggleDictation" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ToggleDictation": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.toggleDictation);
    },
    "call_ToggleDictation": (retPtr: Pointer): void => {
      const _ret = WEBEXT.accessibilityPrivate.toggleDictation();
    },
    "try_ToggleDictation": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityPrivate.toggleDictation();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateDictationBubble": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "updateDictationBubble" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateDictationBubble": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.updateDictationBubble);
    },
    "call_UpdateDictationBubble": (retPtr: Pointer, properties: Pointer): void => {
      const properties_ffi = {};

      properties_ffi["hints"] = A.load.Ref(properties + 0, undefined);
      properties_ffi["icon"] = A.load.Enum(properties + 4, ["hidden", "standby", "macroSuccess", "macroFail"]);
      properties_ffi["text"] = A.load.Ref(properties + 8, undefined);
      properties_ffi["visible"] = A.load.Bool(properties + 12);

      const _ret = WEBEXT.accessibilityPrivate.updateDictationBubble(properties_ffi);
    },
    "try_UpdateDictationBubble": (retPtr: Pointer, errPtr: Pointer, properties: Pointer): heap.Ref<boolean> => {
      try {
        const properties_ffi = {};

        properties_ffi["hints"] = A.load.Ref(properties + 0, undefined);
        properties_ffi["icon"] = A.load.Enum(properties + 4, ["hidden", "standby", "macroSuccess", "macroFail"]);
        properties_ffi["text"] = A.load.Ref(properties + 8, undefined);
        properties_ffi["visible"] = A.load.Bool(properties + 12);

        const _ret = WEBEXT.accessibilityPrivate.updateDictationBubble(properties_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateSelectToSpeakPanel": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "updateSelectToSpeakPanel" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateSelectToSpeakPanel": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.updateSelectToSpeakPanel);
    },
    "call_UpdateSelectToSpeakPanel": (
      retPtr: Pointer,
      show: heap.Ref<boolean>,
      anchor: Pointer,
      isPaused: heap.Ref<boolean>,
      speed: number
    ): void => {
      const anchor_ffi = {};

      anchor_ffi["height"] = A.load.Int64(anchor + 0);
      anchor_ffi["left"] = A.load.Int64(anchor + 8);
      anchor_ffi["top"] = A.load.Int64(anchor + 16);
      anchor_ffi["width"] = A.load.Int64(anchor + 24);

      const _ret = WEBEXT.accessibilityPrivate.updateSelectToSpeakPanel(
        show === A.H.TRUE,
        anchor_ffi,
        isPaused === A.H.TRUE,
        speed
      );
    },
    "try_UpdateSelectToSpeakPanel": (
      retPtr: Pointer,
      errPtr: Pointer,
      show: heap.Ref<boolean>,
      anchor: Pointer,
      isPaused: heap.Ref<boolean>,
      speed: number
    ): heap.Ref<boolean> => {
      try {
        const anchor_ffi = {};

        anchor_ffi["height"] = A.load.Int64(anchor + 0);
        anchor_ffi["left"] = A.load.Int64(anchor + 8);
        anchor_ffi["top"] = A.load.Int64(anchor + 16);
        anchor_ffi["width"] = A.load.Int64(anchor + 24);

        const _ret = WEBEXT.accessibilityPrivate.updateSelectToSpeakPanel(
          show === A.H.TRUE,
          anchor_ffi,
          isPaused === A.H.TRUE,
          speed
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateSwitchAccessBubble": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityPrivate && "updateSwitchAccessBubble" in WEBEXT?.accessibilityPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateSwitchAccessBubble": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityPrivate.updateSwitchAccessBubble);
    },
    "call_UpdateSwitchAccessBubble": (
      retPtr: Pointer,
      bubble: number,
      show: heap.Ref<boolean>,
      anchor: Pointer,
      actions: heap.Ref<object>
    ): void => {
      const anchor_ffi = {};

      anchor_ffi["height"] = A.load.Int64(anchor + 0);
      anchor_ffi["left"] = A.load.Int64(anchor + 8);
      anchor_ffi["top"] = A.load.Int64(anchor + 16);
      anchor_ffi["width"] = A.load.Int64(anchor + 24);

      const _ret = WEBEXT.accessibilityPrivate.updateSwitchAccessBubble(
        bubble > 0 && bubble <= 2 ? ["backButton", "menu"][bubble - 1] : undefined,
        show === A.H.TRUE,
        anchor_ffi,
        A.H.get<object>(actions)
      );
    },
    "try_UpdateSwitchAccessBubble": (
      retPtr: Pointer,
      errPtr: Pointer,
      bubble: number,
      show: heap.Ref<boolean>,
      anchor: Pointer,
      actions: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const anchor_ffi = {};

        anchor_ffi["height"] = A.load.Int64(anchor + 0);
        anchor_ffi["left"] = A.load.Int64(anchor + 8);
        anchor_ffi["top"] = A.load.Int64(anchor + 16);
        anchor_ffi["width"] = A.load.Int64(anchor + 24);

        const _ret = WEBEXT.accessibilityPrivate.updateSwitchAccessBubble(
          bubble > 0 && bubble <= 2 ? ["backButton", "menu"][bubble - 1] : undefined,
          show === A.H.TRUE,
          anchor_ffi,
          A.H.get<object>(actions)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
