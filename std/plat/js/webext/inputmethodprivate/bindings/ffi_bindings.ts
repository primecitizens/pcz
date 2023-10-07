import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/inputmethodprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_AutoCapitalizeType": (ref: heap.Ref<string>): number => {
      const idx = ["off", "characters", "words", "sentences"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_FinishComposingTextArgParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Int64(ptr + 0, x["contextID"] === undefined ? 0 : (x["contextID"] as number));
      }
    },
    "load_FinishComposingTextArgParameters": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["contextID"] = A.load.Int64(ptr + 0);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_FocusReason": (ref: heap.Ref<string>): number => {
      const idx = ["mouse", "touch", "pen", "other"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_GetInputMethodConfigReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 1, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 0, x["isImeMenuActivated"] ? true : false);
        A.store.Bool(ptr + 1, x["isPhysicalKeyboardAutocorrectEnabled"] ? true : false);
      }
    },
    "load_GetInputMethodConfigReturnType": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["isImeMenuActivated"] = A.load.Bool(ptr + 0);
      x["isPhysicalKeyboardAutocorrectEnabled"] = A.load.Bool(ptr + 1);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetInputMethodsReturnTypeElem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["indicator"]);
        A.store.Ref(ptr + 8, x["name"]);
      }
    },
    "load_GetInputMethodsReturnTypeElem": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["indicator"] = A.load.Ref(ptr + 4, undefined);
      x["name"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetSurroundingTextReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["after"]);
        A.store.Ref(ptr + 4, x["before"]);
        A.store.Ref(ptr + 8, x["selected"]);
      }
    },
    "load_GetSurroundingTextReturnType": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["after"] = A.load.Ref(ptr + 0, undefined);
      x["before"] = A.load.Ref(ptr + 4, undefined);
      x["selected"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetTextFieldBoundsArgParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Int64(ptr + 0, x["contextID"] === undefined ? 0 : (x["contextID"] as number));
      }
    },
    "load_GetTextFieldBoundsArgParameters": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["contextID"] = A.load.Int64(ptr + 0);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetTextFieldBoundsReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Int64(ptr + 8, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Int64(ptr + 16, x["x"] === undefined ? 0 : (x["x"] as number));
        A.store.Int64(ptr + 24, x["y"] === undefined ? 0 : (x["y"] as number));
      }
    },
    "load_GetTextFieldBoundsReturnType": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["height"] = A.load.Int64(ptr + 0);
      x["width"] = A.load.Int64(ptr + 8);
      x["x"] = A.load.Int64(ptr + 16);
      x["y"] = A.load.Int64(ptr + 24);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_InputModeType": (ref: heap.Ref<string>): number => {
      const idx = ["noKeyboard", "text", "tel", "url", "email", "numeric", "decimal", "search"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_InputContextType": (ref: heap.Ref<string>): number => {
      const idx = ["text", "search", "tel", "url", "email", "number", "password", "null"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_InputContext": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 40, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 9, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Enum(ptr + 24, -1);
        A.store.Enum(ptr + 28, -1);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 33, false);
        A.store.Enum(ptr + 36, -1);
      } else {
        A.store.Bool(ptr + 40, true);
        A.store.Ref(ptr + 0, x["appKey"]);
        A.store.Enum(ptr + 4, ["off", "characters", "words", "sentences"].indexOf(x["autoCapitalize"] as string));
        A.store.Bool(ptr + 8, x["autoComplete"] ? true : false);
        A.store.Bool(ptr + 9, x["autoCorrect"] ? true : false);
        A.store.Int64(ptr + 16, x["contextID"] === undefined ? 0 : (x["contextID"] as number));
        A.store.Enum(ptr + 24, ["mouse", "touch", "pen", "other"].indexOf(x["focusReason"] as string));
        A.store.Enum(
          ptr + 28,
          ["noKeyboard", "text", "tel", "url", "email", "numeric", "decimal", "search"].indexOf(x["mode"] as string)
        );
        A.store.Bool(ptr + 32, x["shouldDoLearning"] ? true : false);
        A.store.Bool(ptr + 33, x["spellCheck"] ? true : false);
        A.store.Enum(
          ptr + 36,
          ["text", "search", "tel", "url", "email", "number", "password", "null"].indexOf(x["type"] as string)
        );
      }
    },
    "load_InputContext": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["appKey"] = A.load.Ref(ptr + 0, undefined);
      x["autoCapitalize"] = A.load.Enum(ptr + 4, ["off", "characters", "words", "sentences"]);
      x["autoComplete"] = A.load.Bool(ptr + 8);
      x["autoCorrect"] = A.load.Bool(ptr + 9);
      x["contextID"] = A.load.Int64(ptr + 16);
      x["focusReason"] = A.load.Enum(ptr + 24, ["mouse", "touch", "pen", "other"]);
      x["mode"] = A.load.Enum(ptr + 28, ["noKeyboard", "text", "tel", "url", "email", "numeric", "decimal", "search"]);
      x["shouldDoLearning"] = A.load.Bool(ptr + 32);
      x["spellCheck"] = A.load.Bool(ptr + 33);
      x["type"] = A.load.Enum(ptr + 36, ["text", "search", "tel", "url", "email", "number", "password", "null"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_InputMethodSettingsFieldPinyinFuzzyConfig": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 22, false);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 23, false);
        A.store.Bool(ptr + 11, false);
      } else {
        A.store.Bool(ptr + 24, true);
        A.store.Bool(ptr + 12, "an_ang" in x ? true : false);
        A.store.Bool(ptr + 0, x["an_ang"] ? true : false);
        A.store.Bool(ptr + 13, "c_ch" in x ? true : false);
        A.store.Bool(ptr + 1, x["c_ch"] ? true : false);
        A.store.Bool(ptr + 14, "en_eng" in x ? true : false);
        A.store.Bool(ptr + 2, x["en_eng"] ? true : false);
        A.store.Bool(ptr + 15, "f_h" in x ? true : false);
        A.store.Bool(ptr + 3, x["f_h"] ? true : false);
        A.store.Bool(ptr + 16, "ian_iang" in x ? true : false);
        A.store.Bool(ptr + 4, x["ian_iang"] ? true : false);
        A.store.Bool(ptr + 17, "in_ing" in x ? true : false);
        A.store.Bool(ptr + 5, x["in_ing"] ? true : false);
        A.store.Bool(ptr + 18, "k_g" in x ? true : false);
        A.store.Bool(ptr + 6, x["k_g"] ? true : false);
        A.store.Bool(ptr + 19, "l_n" in x ? true : false);
        A.store.Bool(ptr + 7, x["l_n"] ? true : false);
        A.store.Bool(ptr + 20, "r_l" in x ? true : false);
        A.store.Bool(ptr + 8, x["r_l"] ? true : false);
        A.store.Bool(ptr + 21, "s_sh" in x ? true : false);
        A.store.Bool(ptr + 9, x["s_sh"] ? true : false);
        A.store.Bool(ptr + 22, "uan_uang" in x ? true : false);
        A.store.Bool(ptr + 10, x["uan_uang"] ? true : false);
        A.store.Bool(ptr + 23, "z_zh" in x ? true : false);
        A.store.Bool(ptr + 11, x["z_zh"] ? true : false);
      }
    },
    "load_InputMethodSettingsFieldPinyinFuzzyConfig": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["an_ang"] = A.load.Bool(ptr + 0);
      } else {
        delete x["an_ang"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["c_ch"] = A.load.Bool(ptr + 1);
      } else {
        delete x["c_ch"];
      }
      if (A.load.Bool(ptr + 14)) {
        x["en_eng"] = A.load.Bool(ptr + 2);
      } else {
        delete x["en_eng"];
      }
      if (A.load.Bool(ptr + 15)) {
        x["f_h"] = A.load.Bool(ptr + 3);
      } else {
        delete x["f_h"];
      }
      if (A.load.Bool(ptr + 16)) {
        x["ian_iang"] = A.load.Bool(ptr + 4);
      } else {
        delete x["ian_iang"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["in_ing"] = A.load.Bool(ptr + 5);
      } else {
        delete x["in_ing"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["k_g"] = A.load.Bool(ptr + 6);
      } else {
        delete x["k_g"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["l_n"] = A.load.Bool(ptr + 7);
      } else {
        delete x["l_n"];
      }
      if (A.load.Bool(ptr + 20)) {
        x["r_l"] = A.load.Bool(ptr + 8);
      } else {
        delete x["r_l"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["s_sh"] = A.load.Bool(ptr + 9);
      } else {
        delete x["s_sh"];
      }
      if (A.load.Bool(ptr + 22)) {
        x["uan_uang"] = A.load.Bool(ptr + 10);
      } else {
        delete x["uan_uang"];
      }
      if (A.load.Bool(ptr + 23)) {
        x["z_zh"] = A.load.Bool(ptr + 11);
      } else {
        delete x["z_zh"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_InputMethodSettings": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 138, false);
        A.store.Bool(ptr + 108, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 109, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 110, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 111, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 112, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 113, false);
        A.store.Bool(ptr + 5, false);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 114, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 115, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 116, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Bool(ptr + 117, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 118, false);
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 119, false);
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 120, false);
        A.store.Bool(ptr + 27, false);
        A.store.Bool(ptr + 121, false);
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 122, false);
        A.store.Bool(ptr + 29, false);
        A.store.Bool(ptr + 123, false);
        A.store.Bool(ptr + 30, false);
        A.store.Bool(ptr + 124, false);
        A.store.Bool(ptr + 31, false);
        A.store.Bool(ptr + 125, false);
        A.store.Bool(ptr + 32, false);

        A.store.Bool(ptr + 33 + 24, false);
        A.store.Bool(ptr + 33 + 12, false);
        A.store.Bool(ptr + 33 + 0, false);
        A.store.Bool(ptr + 33 + 13, false);
        A.store.Bool(ptr + 33 + 1, false);
        A.store.Bool(ptr + 33 + 14, false);
        A.store.Bool(ptr + 33 + 2, false);
        A.store.Bool(ptr + 33 + 15, false);
        A.store.Bool(ptr + 33 + 3, false);
        A.store.Bool(ptr + 33 + 16, false);
        A.store.Bool(ptr + 33 + 4, false);
        A.store.Bool(ptr + 33 + 17, false);
        A.store.Bool(ptr + 33 + 5, false);
        A.store.Bool(ptr + 33 + 18, false);
        A.store.Bool(ptr + 33 + 6, false);
        A.store.Bool(ptr + 33 + 19, false);
        A.store.Bool(ptr + 33 + 7, false);
        A.store.Bool(ptr + 33 + 20, false);
        A.store.Bool(ptr + 33 + 8, false);
        A.store.Bool(ptr + 33 + 21, false);
        A.store.Bool(ptr + 33 + 9, false);
        A.store.Bool(ptr + 33 + 22, false);
        A.store.Bool(ptr + 33 + 10, false);
        A.store.Bool(ptr + 33 + 23, false);
        A.store.Bool(ptr + 33 + 11, false);
        A.store.Bool(ptr + 126, false);
        A.store.Bool(ptr + 58, false);
        A.store.Bool(ptr + 127, false);
        A.store.Bool(ptr + 59, false);
        A.store.Bool(ptr + 128, false);
        A.store.Bool(ptr + 60, false);
        A.store.Bool(ptr + 129, false);
        A.store.Bool(ptr + 61, false);
        A.store.Bool(ptr + 130, false);
        A.store.Bool(ptr + 62, false);
        A.store.Bool(ptr + 131, false);
        A.store.Bool(ptr + 63, false);
        A.store.Bool(ptr + 132, false);
        A.store.Bool(ptr + 64, false);
        A.store.Bool(ptr + 133, false);
        A.store.Bool(ptr + 65, false);
        A.store.Bool(ptr + 134, false);
        A.store.Bool(ptr + 66, false);
        A.store.Bool(ptr + 135, false);
        A.store.Int64(ptr + 72, 0);
        A.store.Bool(ptr + 136, false);
        A.store.Bool(ptr + 80, false);
        A.store.Ref(ptr + 84, undefined);
        A.store.Ref(ptr + 88, undefined);
        A.store.Bool(ptr + 137, false);
        A.store.Int64(ptr + 96, 0);
        A.store.Ref(ptr + 104, undefined);
      } else {
        A.store.Bool(ptr + 138, true);
        A.store.Bool(ptr + 108, "enableCompletion" in x ? true : false);
        A.store.Bool(ptr + 0, x["enableCompletion"] ? true : false);
        A.store.Bool(ptr + 109, "enableDoubleSpacePeriod" in x ? true : false);
        A.store.Bool(ptr + 1, x["enableDoubleSpacePeriod"] ? true : false);
        A.store.Bool(ptr + 110, "enableGestureTyping" in x ? true : false);
        A.store.Bool(ptr + 2, x["enableGestureTyping"] ? true : false);
        A.store.Bool(ptr + 111, "enablePrediction" in x ? true : false);
        A.store.Bool(ptr + 3, x["enablePrediction"] ? true : false);
        A.store.Bool(ptr + 112, "enableSoundOnKeypress" in x ? true : false);
        A.store.Bool(ptr + 4, x["enableSoundOnKeypress"] ? true : false);
        A.store.Bool(ptr + 113, "koreanEnableSyllableInput" in x ? true : false);
        A.store.Bool(ptr + 5, x["koreanEnableSyllableInput"] ? true : false);
        A.store.Ref(ptr + 8, x["koreanKeyboardLayout"]);
        A.store.Bool(ptr + 114, "koreanShowHangulCandidate" in x ? true : false);
        A.store.Bool(ptr + 12, x["koreanShowHangulCandidate"] ? true : false);
        A.store.Bool(ptr + 115, "physicalKeyboardAutoCorrectionEnabledByDefault" in x ? true : false);
        A.store.Bool(ptr + 13, x["physicalKeyboardAutoCorrectionEnabledByDefault"] ? true : false);
        A.store.Bool(ptr + 116, "physicalKeyboardAutoCorrectionLevel" in x ? true : false);
        A.store.Int64(
          ptr + 16,
          x["physicalKeyboardAutoCorrectionLevel"] === undefined
            ? 0
            : (x["physicalKeyboardAutoCorrectionLevel"] as number)
        );
        A.store.Bool(ptr + 117, "physicalKeyboardEnableCapitalization" in x ? true : false);
        A.store.Bool(ptr + 24, x["physicalKeyboardEnableCapitalization"] ? true : false);
        A.store.Bool(ptr + 118, "physicalKeyboardEnableDiacriticsOnLongpress" in x ? true : false);
        A.store.Bool(ptr + 25, x["physicalKeyboardEnableDiacriticsOnLongpress"] ? true : false);
        A.store.Bool(ptr + 119, "physicalKeyboardEnablePredictiveWriting" in x ? true : false);
        A.store.Bool(ptr + 26, x["physicalKeyboardEnablePredictiveWriting"] ? true : false);
        A.store.Bool(ptr + 120, "pinyinChinesePunctuation" in x ? true : false);
        A.store.Bool(ptr + 27, x["pinyinChinesePunctuation"] ? true : false);
        A.store.Bool(ptr + 121, "pinyinDefaultChinese" in x ? true : false);
        A.store.Bool(ptr + 28, x["pinyinDefaultChinese"] ? true : false);
        A.store.Bool(ptr + 122, "pinyinEnableFuzzy" in x ? true : false);
        A.store.Bool(ptr + 29, x["pinyinEnableFuzzy"] ? true : false);
        A.store.Bool(ptr + 123, "pinyinEnableLowerPaging" in x ? true : false);
        A.store.Bool(ptr + 30, x["pinyinEnableLowerPaging"] ? true : false);
        A.store.Bool(ptr + 124, "pinyinEnableUpperPaging" in x ? true : false);
        A.store.Bool(ptr + 31, x["pinyinEnableUpperPaging"] ? true : false);
        A.store.Bool(ptr + 125, "pinyinFullWidthCharacter" in x ? true : false);
        A.store.Bool(ptr + 32, x["pinyinFullWidthCharacter"] ? true : false);

        if (typeof x["pinyinFuzzyConfig"] === "undefined") {
          A.store.Bool(ptr + 33 + 24, false);
          A.store.Bool(ptr + 33 + 12, false);
          A.store.Bool(ptr + 33 + 0, false);
          A.store.Bool(ptr + 33 + 13, false);
          A.store.Bool(ptr + 33 + 1, false);
          A.store.Bool(ptr + 33 + 14, false);
          A.store.Bool(ptr + 33 + 2, false);
          A.store.Bool(ptr + 33 + 15, false);
          A.store.Bool(ptr + 33 + 3, false);
          A.store.Bool(ptr + 33 + 16, false);
          A.store.Bool(ptr + 33 + 4, false);
          A.store.Bool(ptr + 33 + 17, false);
          A.store.Bool(ptr + 33 + 5, false);
          A.store.Bool(ptr + 33 + 18, false);
          A.store.Bool(ptr + 33 + 6, false);
          A.store.Bool(ptr + 33 + 19, false);
          A.store.Bool(ptr + 33 + 7, false);
          A.store.Bool(ptr + 33 + 20, false);
          A.store.Bool(ptr + 33 + 8, false);
          A.store.Bool(ptr + 33 + 21, false);
          A.store.Bool(ptr + 33 + 9, false);
          A.store.Bool(ptr + 33 + 22, false);
          A.store.Bool(ptr + 33 + 10, false);
          A.store.Bool(ptr + 33 + 23, false);
          A.store.Bool(ptr + 33 + 11, false);
        } else {
          A.store.Bool(ptr + 33 + 24, true);
          A.store.Bool(ptr + 33 + 12, "an_ang" in x["pinyinFuzzyConfig"] ? true : false);
          A.store.Bool(ptr + 33 + 0, x["pinyinFuzzyConfig"]["an_ang"] ? true : false);
          A.store.Bool(ptr + 33 + 13, "c_ch" in x["pinyinFuzzyConfig"] ? true : false);
          A.store.Bool(ptr + 33 + 1, x["pinyinFuzzyConfig"]["c_ch"] ? true : false);
          A.store.Bool(ptr + 33 + 14, "en_eng" in x["pinyinFuzzyConfig"] ? true : false);
          A.store.Bool(ptr + 33 + 2, x["pinyinFuzzyConfig"]["en_eng"] ? true : false);
          A.store.Bool(ptr + 33 + 15, "f_h" in x["pinyinFuzzyConfig"] ? true : false);
          A.store.Bool(ptr + 33 + 3, x["pinyinFuzzyConfig"]["f_h"] ? true : false);
          A.store.Bool(ptr + 33 + 16, "ian_iang" in x["pinyinFuzzyConfig"] ? true : false);
          A.store.Bool(ptr + 33 + 4, x["pinyinFuzzyConfig"]["ian_iang"] ? true : false);
          A.store.Bool(ptr + 33 + 17, "in_ing" in x["pinyinFuzzyConfig"] ? true : false);
          A.store.Bool(ptr + 33 + 5, x["pinyinFuzzyConfig"]["in_ing"] ? true : false);
          A.store.Bool(ptr + 33 + 18, "k_g" in x["pinyinFuzzyConfig"] ? true : false);
          A.store.Bool(ptr + 33 + 6, x["pinyinFuzzyConfig"]["k_g"] ? true : false);
          A.store.Bool(ptr + 33 + 19, "l_n" in x["pinyinFuzzyConfig"] ? true : false);
          A.store.Bool(ptr + 33 + 7, x["pinyinFuzzyConfig"]["l_n"] ? true : false);
          A.store.Bool(ptr + 33 + 20, "r_l" in x["pinyinFuzzyConfig"] ? true : false);
          A.store.Bool(ptr + 33 + 8, x["pinyinFuzzyConfig"]["r_l"] ? true : false);
          A.store.Bool(ptr + 33 + 21, "s_sh" in x["pinyinFuzzyConfig"] ? true : false);
          A.store.Bool(ptr + 33 + 9, x["pinyinFuzzyConfig"]["s_sh"] ? true : false);
          A.store.Bool(ptr + 33 + 22, "uan_uang" in x["pinyinFuzzyConfig"] ? true : false);
          A.store.Bool(ptr + 33 + 10, x["pinyinFuzzyConfig"]["uan_uang"] ? true : false);
          A.store.Bool(ptr + 33 + 23, "z_zh" in x["pinyinFuzzyConfig"] ? true : false);
          A.store.Bool(ptr + 33 + 11, x["pinyinFuzzyConfig"]["z_zh"] ? true : false);
        }
        A.store.Bool(ptr + 126, "vietnameseTelexAllowFlexibleDiacritics" in x ? true : false);
        A.store.Bool(ptr + 58, x["vietnameseTelexAllowFlexibleDiacritics"] ? true : false);
        A.store.Bool(ptr + 127, "vietnameseTelexInsertDoubleHornOnUo" in x ? true : false);
        A.store.Bool(ptr + 59, x["vietnameseTelexInsertDoubleHornOnUo"] ? true : false);
        A.store.Bool(ptr + 128, "vietnameseTelexInsertUHornOnW" in x ? true : false);
        A.store.Bool(ptr + 60, x["vietnameseTelexInsertUHornOnW"] ? true : false);
        A.store.Bool(ptr + 129, "vietnameseTelexNewStyleToneMarkPlacement" in x ? true : false);
        A.store.Bool(ptr + 61, x["vietnameseTelexNewStyleToneMarkPlacement"] ? true : false);
        A.store.Bool(ptr + 130, "vietnameseTelexShowUnderline" in x ? true : false);
        A.store.Bool(ptr + 62, x["vietnameseTelexShowUnderline"] ? true : false);
        A.store.Bool(ptr + 131, "vietnameseVniAllowFlexibleDiacritics" in x ? true : false);
        A.store.Bool(ptr + 63, x["vietnameseVniAllowFlexibleDiacritics"] ? true : false);
        A.store.Bool(ptr + 132, "vietnameseVniInsertDoubleHornOnUo" in x ? true : false);
        A.store.Bool(ptr + 64, x["vietnameseVniInsertDoubleHornOnUo"] ? true : false);
        A.store.Bool(ptr + 133, "vietnameseVniNewStyleToneMarkPlacement" in x ? true : false);
        A.store.Bool(ptr + 65, x["vietnameseVniNewStyleToneMarkPlacement"] ? true : false);
        A.store.Bool(ptr + 134, "vietnameseVniShowUnderline" in x ? true : false);
        A.store.Bool(ptr + 66, x["vietnameseVniShowUnderline"] ? true : false);
        A.store.Bool(ptr + 135, "virtualKeyboardAutoCorrectionLevel" in x ? true : false);
        A.store.Int64(
          ptr + 72,
          x["virtualKeyboardAutoCorrectionLevel"] === undefined
            ? 0
            : (x["virtualKeyboardAutoCorrectionLevel"] as number)
        );
        A.store.Bool(ptr + 136, "virtualKeyboardEnableCapitalization" in x ? true : false);
        A.store.Bool(ptr + 80, x["virtualKeyboardEnableCapitalization"] ? true : false);
        A.store.Ref(ptr + 84, x["xkbLayout"]);
        A.store.Ref(ptr + 88, x["zhuyinKeyboardLayout"]);
        A.store.Bool(ptr + 137, "zhuyinPageSize" in x ? true : false);
        A.store.Int64(ptr + 96, x["zhuyinPageSize"] === undefined ? 0 : (x["zhuyinPageSize"] as number));
        A.store.Ref(ptr + 104, x["zhuyinSelectKeys"]);
      }
    },
    "load_InputMethodSettings": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 108)) {
        x["enableCompletion"] = A.load.Bool(ptr + 0);
      } else {
        delete x["enableCompletion"];
      }
      if (A.load.Bool(ptr + 109)) {
        x["enableDoubleSpacePeriod"] = A.load.Bool(ptr + 1);
      } else {
        delete x["enableDoubleSpacePeriod"];
      }
      if (A.load.Bool(ptr + 110)) {
        x["enableGestureTyping"] = A.load.Bool(ptr + 2);
      } else {
        delete x["enableGestureTyping"];
      }
      if (A.load.Bool(ptr + 111)) {
        x["enablePrediction"] = A.load.Bool(ptr + 3);
      } else {
        delete x["enablePrediction"];
      }
      if (A.load.Bool(ptr + 112)) {
        x["enableSoundOnKeypress"] = A.load.Bool(ptr + 4);
      } else {
        delete x["enableSoundOnKeypress"];
      }
      if (A.load.Bool(ptr + 113)) {
        x["koreanEnableSyllableInput"] = A.load.Bool(ptr + 5);
      } else {
        delete x["koreanEnableSyllableInput"];
      }
      x["koreanKeyboardLayout"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 114)) {
        x["koreanShowHangulCandidate"] = A.load.Bool(ptr + 12);
      } else {
        delete x["koreanShowHangulCandidate"];
      }
      if (A.load.Bool(ptr + 115)) {
        x["physicalKeyboardAutoCorrectionEnabledByDefault"] = A.load.Bool(ptr + 13);
      } else {
        delete x["physicalKeyboardAutoCorrectionEnabledByDefault"];
      }
      if (A.load.Bool(ptr + 116)) {
        x["physicalKeyboardAutoCorrectionLevel"] = A.load.Int64(ptr + 16);
      } else {
        delete x["physicalKeyboardAutoCorrectionLevel"];
      }
      if (A.load.Bool(ptr + 117)) {
        x["physicalKeyboardEnableCapitalization"] = A.load.Bool(ptr + 24);
      } else {
        delete x["physicalKeyboardEnableCapitalization"];
      }
      if (A.load.Bool(ptr + 118)) {
        x["physicalKeyboardEnableDiacriticsOnLongpress"] = A.load.Bool(ptr + 25);
      } else {
        delete x["physicalKeyboardEnableDiacriticsOnLongpress"];
      }
      if (A.load.Bool(ptr + 119)) {
        x["physicalKeyboardEnablePredictiveWriting"] = A.load.Bool(ptr + 26);
      } else {
        delete x["physicalKeyboardEnablePredictiveWriting"];
      }
      if (A.load.Bool(ptr + 120)) {
        x["pinyinChinesePunctuation"] = A.load.Bool(ptr + 27);
      } else {
        delete x["pinyinChinesePunctuation"];
      }
      if (A.load.Bool(ptr + 121)) {
        x["pinyinDefaultChinese"] = A.load.Bool(ptr + 28);
      } else {
        delete x["pinyinDefaultChinese"];
      }
      if (A.load.Bool(ptr + 122)) {
        x["pinyinEnableFuzzy"] = A.load.Bool(ptr + 29);
      } else {
        delete x["pinyinEnableFuzzy"];
      }
      if (A.load.Bool(ptr + 123)) {
        x["pinyinEnableLowerPaging"] = A.load.Bool(ptr + 30);
      } else {
        delete x["pinyinEnableLowerPaging"];
      }
      if (A.load.Bool(ptr + 124)) {
        x["pinyinEnableUpperPaging"] = A.load.Bool(ptr + 31);
      } else {
        delete x["pinyinEnableUpperPaging"];
      }
      if (A.load.Bool(ptr + 125)) {
        x["pinyinFullWidthCharacter"] = A.load.Bool(ptr + 32);
      } else {
        delete x["pinyinFullWidthCharacter"];
      }
      if (A.load.Bool(ptr + 33 + 24)) {
        x["pinyinFuzzyConfig"] = {};
        if (A.load.Bool(ptr + 33 + 12)) {
          x["pinyinFuzzyConfig"]["an_ang"] = A.load.Bool(ptr + 33 + 0);
        } else {
          delete x["pinyinFuzzyConfig"]["an_ang"];
        }
        if (A.load.Bool(ptr + 33 + 13)) {
          x["pinyinFuzzyConfig"]["c_ch"] = A.load.Bool(ptr + 33 + 1);
        } else {
          delete x["pinyinFuzzyConfig"]["c_ch"];
        }
        if (A.load.Bool(ptr + 33 + 14)) {
          x["pinyinFuzzyConfig"]["en_eng"] = A.load.Bool(ptr + 33 + 2);
        } else {
          delete x["pinyinFuzzyConfig"]["en_eng"];
        }
        if (A.load.Bool(ptr + 33 + 15)) {
          x["pinyinFuzzyConfig"]["f_h"] = A.load.Bool(ptr + 33 + 3);
        } else {
          delete x["pinyinFuzzyConfig"]["f_h"];
        }
        if (A.load.Bool(ptr + 33 + 16)) {
          x["pinyinFuzzyConfig"]["ian_iang"] = A.load.Bool(ptr + 33 + 4);
        } else {
          delete x["pinyinFuzzyConfig"]["ian_iang"];
        }
        if (A.load.Bool(ptr + 33 + 17)) {
          x["pinyinFuzzyConfig"]["in_ing"] = A.load.Bool(ptr + 33 + 5);
        } else {
          delete x["pinyinFuzzyConfig"]["in_ing"];
        }
        if (A.load.Bool(ptr + 33 + 18)) {
          x["pinyinFuzzyConfig"]["k_g"] = A.load.Bool(ptr + 33 + 6);
        } else {
          delete x["pinyinFuzzyConfig"]["k_g"];
        }
        if (A.load.Bool(ptr + 33 + 19)) {
          x["pinyinFuzzyConfig"]["l_n"] = A.load.Bool(ptr + 33 + 7);
        } else {
          delete x["pinyinFuzzyConfig"]["l_n"];
        }
        if (A.load.Bool(ptr + 33 + 20)) {
          x["pinyinFuzzyConfig"]["r_l"] = A.load.Bool(ptr + 33 + 8);
        } else {
          delete x["pinyinFuzzyConfig"]["r_l"];
        }
        if (A.load.Bool(ptr + 33 + 21)) {
          x["pinyinFuzzyConfig"]["s_sh"] = A.load.Bool(ptr + 33 + 9);
        } else {
          delete x["pinyinFuzzyConfig"]["s_sh"];
        }
        if (A.load.Bool(ptr + 33 + 22)) {
          x["pinyinFuzzyConfig"]["uan_uang"] = A.load.Bool(ptr + 33 + 10);
        } else {
          delete x["pinyinFuzzyConfig"]["uan_uang"];
        }
        if (A.load.Bool(ptr + 33 + 23)) {
          x["pinyinFuzzyConfig"]["z_zh"] = A.load.Bool(ptr + 33 + 11);
        } else {
          delete x["pinyinFuzzyConfig"]["z_zh"];
        }
      } else {
        delete x["pinyinFuzzyConfig"];
      }
      if (A.load.Bool(ptr + 126)) {
        x["vietnameseTelexAllowFlexibleDiacritics"] = A.load.Bool(ptr + 58);
      } else {
        delete x["vietnameseTelexAllowFlexibleDiacritics"];
      }
      if (A.load.Bool(ptr + 127)) {
        x["vietnameseTelexInsertDoubleHornOnUo"] = A.load.Bool(ptr + 59);
      } else {
        delete x["vietnameseTelexInsertDoubleHornOnUo"];
      }
      if (A.load.Bool(ptr + 128)) {
        x["vietnameseTelexInsertUHornOnW"] = A.load.Bool(ptr + 60);
      } else {
        delete x["vietnameseTelexInsertUHornOnW"];
      }
      if (A.load.Bool(ptr + 129)) {
        x["vietnameseTelexNewStyleToneMarkPlacement"] = A.load.Bool(ptr + 61);
      } else {
        delete x["vietnameseTelexNewStyleToneMarkPlacement"];
      }
      if (A.load.Bool(ptr + 130)) {
        x["vietnameseTelexShowUnderline"] = A.load.Bool(ptr + 62);
      } else {
        delete x["vietnameseTelexShowUnderline"];
      }
      if (A.load.Bool(ptr + 131)) {
        x["vietnameseVniAllowFlexibleDiacritics"] = A.load.Bool(ptr + 63);
      } else {
        delete x["vietnameseVniAllowFlexibleDiacritics"];
      }
      if (A.load.Bool(ptr + 132)) {
        x["vietnameseVniInsertDoubleHornOnUo"] = A.load.Bool(ptr + 64);
      } else {
        delete x["vietnameseVniInsertDoubleHornOnUo"];
      }
      if (A.load.Bool(ptr + 133)) {
        x["vietnameseVniNewStyleToneMarkPlacement"] = A.load.Bool(ptr + 65);
      } else {
        delete x["vietnameseVniNewStyleToneMarkPlacement"];
      }
      if (A.load.Bool(ptr + 134)) {
        x["vietnameseVniShowUnderline"] = A.load.Bool(ptr + 66);
      } else {
        delete x["vietnameseVniShowUnderline"];
      }
      if (A.load.Bool(ptr + 135)) {
        x["virtualKeyboardAutoCorrectionLevel"] = A.load.Int64(ptr + 72);
      } else {
        delete x["virtualKeyboardAutoCorrectionLevel"];
      }
      if (A.load.Bool(ptr + 136)) {
        x["virtualKeyboardEnableCapitalization"] = A.load.Bool(ptr + 80);
      } else {
        delete x["virtualKeyboardEnableCapitalization"];
      }
      x["xkbLayout"] = A.load.Ref(ptr + 84, undefined);
      x["zhuyinKeyboardLayout"] = A.load.Ref(ptr + 88, undefined);
      if (A.load.Bool(ptr + 137)) {
        x["zhuyinPageSize"] = A.load.Int64(ptr + 96);
      } else {
        delete x["zhuyinPageSize"];
      }
      x["zhuyinSelectKeys"] = A.load.Ref(ptr + 104, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_MenuItemStyle": (ref: heap.Ref<string>): number => {
      const idx = ["check", "radio", "separator"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_MenuItem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 1, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Enum(ptr + 12, -1);
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 16, false);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Bool(ptr + 17, "checked" in x ? true : false);
        A.store.Bool(ptr + 0, x["checked"] ? true : false);
        A.store.Bool(ptr + 18, "enabled" in x ? true : false);
        A.store.Bool(ptr + 1, x["enabled"] ? true : false);
        A.store.Ref(ptr + 4, x["id"]);
        A.store.Ref(ptr + 8, x["label"]);
        A.store.Enum(ptr + 12, ["check", "radio", "separator"].indexOf(x["style"] as string));
        A.store.Bool(ptr + 19, "visible" in x ? true : false);
        A.store.Bool(ptr + 16, x["visible"] ? true : false);
      }
    },
    "load_MenuItem": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 17)) {
        x["checked"] = A.load.Bool(ptr + 0);
      } else {
        delete x["checked"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["enabled"] = A.load.Bool(ptr + 1);
      } else {
        delete x["enabled"];
      }
      x["id"] = A.load.Ref(ptr + 4, undefined);
      x["label"] = A.load.Ref(ptr + 8, undefined);
      x["style"] = A.load.Enum(ptr + 12, ["check", "radio", "separator"]);
      if (A.load.Bool(ptr + 19)) {
        x["visible"] = A.load.Bool(ptr + 16);
      } else {
        delete x["visible"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnAutocorrectArgParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Int64(ptr + 16, 0);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 28, true);
        A.store.Int64(ptr + 0, x["contextID"] === undefined ? 0 : (x["contextID"] as number));
        A.store.Ref(ptr + 8, x["correctedWord"]);
        A.store.Int64(ptr + 16, x["startIndex"] === undefined ? 0 : (x["startIndex"] as number));
        A.store.Ref(ptr + 24, x["typedWord"]);
      }
    },
    "load_OnAutocorrectArgParameters": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["contextID"] = A.load.Int64(ptr + 0);
      x["correctedWord"] = A.load.Ref(ptr + 8, undefined);
      x["startIndex"] = A.load.Int64(ptr + 16);
      x["typedWord"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnCaretBoundsChangedArgCaretBounds": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 32, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
        A.store.Int64(ptr + 16, 0);
        A.store.Int64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 32, true);
        A.store.Int64(ptr + 0, x["h"] === undefined ? 0 : (x["h"] as number));
        A.store.Int64(ptr + 8, x["w"] === undefined ? 0 : (x["w"] as number));
        A.store.Int64(ptr + 16, x["x"] === undefined ? 0 : (x["x"] as number));
        A.store.Int64(ptr + 24, x["y"] === undefined ? 0 : (x["y"] as number));
      }
    },
    "load_OnCaretBoundsChangedArgCaretBounds": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["h"] = A.load.Int64(ptr + 0);
      x["w"] = A.load.Int64(ptr + 8);
      x["x"] = A.load.Int64(ptr + 16);
      x["y"] = A.load.Int64(ptr + 24);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_UnderlineStyle": (ref: heap.Ref<string>): number => {
      const idx = ["underline", "doubleUnderline", "noUnderline"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SetCompositionRangeArgParametersFieldSegmentsElem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Int64(ptr + 0, x["end"] === undefined ? 0 : (x["end"] as number));
        A.store.Int64(ptr + 8, x["start"] === undefined ? 0 : (x["start"] as number));
        A.store.Enum(ptr + 16, ["underline", "doubleUnderline", "noUnderline"].indexOf(x["style"] as string));
      }
    },
    "load_SetCompositionRangeArgParametersFieldSegmentsElem": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["end"] = A.load.Int64(ptr + 0);
      x["start"] = A.load.Int64(ptr + 8);
      x["style"] = A.load.Enum(ptr + 16, ["underline", "doubleUnderline", "noUnderline"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetCompositionRangeArgParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 32, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Int64(ptr + 16, 0);
        A.store.Int64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 32, true);
        A.store.Int64(ptr + 0, x["contextID"] === undefined ? 0 : (x["contextID"] as number));
        A.store.Ref(ptr + 8, x["segments"]);
        A.store.Int64(ptr + 16, x["selectionAfter"] === undefined ? 0 : (x["selectionAfter"] as number));
        A.store.Int64(ptr + 24, x["selectionBefore"] === undefined ? 0 : (x["selectionBefore"] as number));
      }
    },
    "load_SetCompositionRangeArgParameters": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["contextID"] = A.load.Int64(ptr + 0);
      x["segments"] = A.load.Ref(ptr + 8, undefined);
      x["selectionAfter"] = A.load.Int64(ptr + 16);
      x["selectionBefore"] = A.load.Int64(ptr + 24);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_AddWordToDictionary": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "addWordToDictionary" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddWordToDictionary": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.addWordToDictionary);
    },
    "call_AddWordToDictionary": (retPtr: Pointer, word: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.addWordToDictionary(A.H.get<object>(word));
      A.store.Ref(retPtr, _ret);
    },
    "try_AddWordToDictionary": (retPtr: Pointer, errPtr: Pointer, word: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.addWordToDictionary(A.H.get<object>(word));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_FetchAllDictionaryWords": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "fetchAllDictionaryWords" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_FetchAllDictionaryWords": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.fetchAllDictionaryWords);
    },
    "call_FetchAllDictionaryWords": (retPtr: Pointer): void => {
      const _ret = WEBEXT.inputMethodPrivate.fetchAllDictionaryWords();
      A.store.Ref(retPtr, _ret);
    },
    "try_FetchAllDictionaryWords": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.fetchAllDictionaryWords();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_FinishComposingText": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "finishComposingText" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_FinishComposingText": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.finishComposingText);
    },
    "call_FinishComposingText": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["contextID"] = A.load.Int64(parameters + 0);

      const _ret = WEBEXT.inputMethodPrivate.finishComposingText(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_FinishComposingText": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["contextID"] = A.load.Int64(parameters + 0);

        const _ret = WEBEXT.inputMethodPrivate.finishComposingText(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCurrentInputMethod": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "getCurrentInputMethod" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCurrentInputMethod": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.getCurrentInputMethod);
    },
    "call_GetCurrentInputMethod": (retPtr: Pointer): void => {
      const _ret = WEBEXT.inputMethodPrivate.getCurrentInputMethod();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCurrentInputMethod": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.getCurrentInputMethod();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetInputMethodConfig": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "getInputMethodConfig" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetInputMethodConfig": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.getInputMethodConfig);
    },
    "call_GetInputMethodConfig": (retPtr: Pointer): void => {
      const _ret = WEBEXT.inputMethodPrivate.getInputMethodConfig();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetInputMethodConfig": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.getInputMethodConfig();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetInputMethods": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "getInputMethods" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetInputMethods": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.getInputMethods);
    },
    "call_GetInputMethods": (retPtr: Pointer): void => {
      const _ret = WEBEXT.inputMethodPrivate.getInputMethods();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetInputMethods": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.getInputMethods();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSettings": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "getSettings" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSettings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.getSettings);
    },
    "call_GetSettings": (retPtr: Pointer, engineID: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.getSettings(A.H.get<object>(engineID));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSettings": (retPtr: Pointer, errPtr: Pointer, engineID: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.getSettings(A.H.get<object>(engineID));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSurroundingText": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "getSurroundingText" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSurroundingText": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.getSurroundingText);
    },
    "call_GetSurroundingText": (retPtr: Pointer, beforeLength: number, afterLength: number): void => {
      const _ret = WEBEXT.inputMethodPrivate.getSurroundingText(beforeLength, afterLength);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSurroundingText": (
      retPtr: Pointer,
      errPtr: Pointer,
      beforeLength: number,
      afterLength: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.getSurroundingText(beforeLength, afterLength);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetTextFieldBounds": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "getTextFieldBounds" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetTextFieldBounds": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.getTextFieldBounds);
    },
    "call_GetTextFieldBounds": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["contextID"] = A.load.Int64(parameters + 0);

      const _ret = WEBEXT.inputMethodPrivate.getTextFieldBounds(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetTextFieldBounds": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["contextID"] = A.load.Int64(parameters + 0);

        const _ret = WEBEXT.inputMethodPrivate.getTextFieldBounds(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HideInputView": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "hideInputView" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HideInputView": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.hideInputView);
    },
    "call_HideInputView": (retPtr: Pointer): void => {
      const _ret = WEBEXT.inputMethodPrivate.hideInputView();
      A.store.Ref(retPtr, _ret);
    },
    "try_HideInputView": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.hideInputView();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_NotifyInputMethodReadyForTesting": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "notifyInputMethodReadyForTesting" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_NotifyInputMethodReadyForTesting": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.notifyInputMethodReadyForTesting);
    },
    "call_NotifyInputMethodReadyForTesting": (retPtr: Pointer): void => {
      const _ret = WEBEXT.inputMethodPrivate.notifyInputMethodReadyForTesting();
    },
    "try_NotifyInputMethodReadyForTesting": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.notifyInputMethodReadyForTesting();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAutocorrect": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "onAutocorrect" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAutocorrect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onAutocorrect);
    },
    "call_OnAutocorrect": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["contextID"] = A.load.Int64(parameters + 0);
      parameters_ffi["correctedWord"] = A.load.Ref(parameters + 8, undefined);
      parameters_ffi["startIndex"] = A.load.Int64(parameters + 16);
      parameters_ffi["typedWord"] = A.load.Ref(parameters + 24, undefined);

      const _ret = WEBEXT.inputMethodPrivate.onAutocorrect(parameters_ffi);
    },
    "try_OnAutocorrect": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["contextID"] = A.load.Int64(parameters + 0);
        parameters_ffi["correctedWord"] = A.load.Ref(parameters + 8, undefined);
        parameters_ffi["startIndex"] = A.load.Int64(parameters + 16);
        parameters_ffi["typedWord"] = A.load.Ref(parameters + 24, undefined);

        const _ret = WEBEXT.inputMethodPrivate.onAutocorrect(parameters_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCaretBoundsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onCaretBoundsChanged &&
        "addListener" in WEBEXT?.inputMethodPrivate?.onCaretBoundsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCaretBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onCaretBoundsChanged.addListener);
    },
    "call_OnCaretBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onCaretBoundsChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnCaretBoundsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onCaretBoundsChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCaretBoundsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onCaretBoundsChanged &&
        "removeListener" in WEBEXT?.inputMethodPrivate?.onCaretBoundsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCaretBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onCaretBoundsChanged.removeListener);
    },
    "call_OffCaretBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onCaretBoundsChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffCaretBoundsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onCaretBoundsChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCaretBoundsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onCaretBoundsChanged &&
        "hasListener" in WEBEXT?.inputMethodPrivate?.onCaretBoundsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCaretBoundsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onCaretBoundsChanged.hasListener);
    },
    "call_HasOnCaretBoundsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onCaretBoundsChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCaretBoundsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onCaretBoundsChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate?.onChanged && "addListener" in WEBEXT?.inputMethodPrivate?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onChanged.addListener);
    },
    "call_OnChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate?.onChanged && "removeListener" in WEBEXT?.inputMethodPrivate?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onChanged.removeListener);
    },
    "call_OffChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate?.onChanged && "hasListener" in WEBEXT?.inputMethodPrivate?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onChanged.hasListener);
    },
    "call_HasOnChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDictionaryChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onDictionaryChanged &&
        "addListener" in WEBEXT?.inputMethodPrivate?.onDictionaryChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDictionaryChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onDictionaryChanged.addListener);
    },
    "call_OnDictionaryChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onDictionaryChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnDictionaryChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onDictionaryChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDictionaryChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onDictionaryChanged &&
        "removeListener" in WEBEXT?.inputMethodPrivate?.onDictionaryChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDictionaryChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onDictionaryChanged.removeListener);
    },
    "call_OffDictionaryChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onDictionaryChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffDictionaryChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onDictionaryChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDictionaryChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onDictionaryChanged &&
        "hasListener" in WEBEXT?.inputMethodPrivate?.onDictionaryChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDictionaryChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onDictionaryChanged.hasListener);
    },
    "call_HasOnDictionaryChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onDictionaryChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDictionaryChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onDictionaryChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDictionaryLoaded": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onDictionaryLoaded &&
        "addListener" in WEBEXT?.inputMethodPrivate?.onDictionaryLoaded
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDictionaryLoaded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onDictionaryLoaded.addListener);
    },
    "call_OnDictionaryLoaded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onDictionaryLoaded.addListener(A.H.get<object>(callback));
    },
    "try_OnDictionaryLoaded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onDictionaryLoaded.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDictionaryLoaded": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onDictionaryLoaded &&
        "removeListener" in WEBEXT?.inputMethodPrivate?.onDictionaryLoaded
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDictionaryLoaded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onDictionaryLoaded.removeListener);
    },
    "call_OffDictionaryLoaded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onDictionaryLoaded.removeListener(A.H.get<object>(callback));
    },
    "try_OffDictionaryLoaded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onDictionaryLoaded.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDictionaryLoaded": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onDictionaryLoaded &&
        "hasListener" in WEBEXT?.inputMethodPrivate?.onDictionaryLoaded
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDictionaryLoaded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onDictionaryLoaded.hasListener);
    },
    "call_HasOnDictionaryLoaded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onDictionaryLoaded.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDictionaryLoaded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onDictionaryLoaded.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnFocus": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate?.onFocus && "addListener" in WEBEXT?.inputMethodPrivate?.onFocus) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnFocus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onFocus.addListener);
    },
    "call_OnFocus": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onFocus.addListener(A.H.get<object>(callback));
    },
    "try_OnFocus": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onFocus.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffFocus": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate?.onFocus && "removeListener" in WEBEXT?.inputMethodPrivate?.onFocus) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffFocus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onFocus.removeListener);
    },
    "call_OffFocus": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onFocus.removeListener(A.H.get<object>(callback));
    },
    "try_OffFocus": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onFocus.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnFocus": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate?.onFocus && "hasListener" in WEBEXT?.inputMethodPrivate?.onFocus) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnFocus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onFocus.hasListener);
    },
    "call_HasOnFocus": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onFocus.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnFocus": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onFocus.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnImeMenuActivationChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onImeMenuActivationChanged &&
        "addListener" in WEBEXT?.inputMethodPrivate?.onImeMenuActivationChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnImeMenuActivationChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.addListener);
    },
    "call_OnImeMenuActivationChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnImeMenuActivationChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffImeMenuActivationChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onImeMenuActivationChanged &&
        "removeListener" in WEBEXT?.inputMethodPrivate?.onImeMenuActivationChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffImeMenuActivationChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.removeListener);
    },
    "call_OffImeMenuActivationChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffImeMenuActivationChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnImeMenuActivationChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onImeMenuActivationChanged &&
        "hasListener" in WEBEXT?.inputMethodPrivate?.onImeMenuActivationChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnImeMenuActivationChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.hasListener);
    },
    "call_HasOnImeMenuActivationChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnImeMenuActivationChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnImeMenuItemsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onImeMenuItemsChanged &&
        "addListener" in WEBEXT?.inputMethodPrivate?.onImeMenuItemsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnImeMenuItemsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.addListener);
    },
    "call_OnImeMenuItemsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnImeMenuItemsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffImeMenuItemsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onImeMenuItemsChanged &&
        "removeListener" in WEBEXT?.inputMethodPrivate?.onImeMenuItemsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffImeMenuItemsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.removeListener);
    },
    "call_OffImeMenuItemsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffImeMenuItemsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnImeMenuItemsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onImeMenuItemsChanged &&
        "hasListener" in WEBEXT?.inputMethodPrivate?.onImeMenuItemsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnImeMenuItemsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.hasListener);
    },
    "call_HasOnImeMenuItemsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnImeMenuItemsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnImeMenuListChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onImeMenuListChanged &&
        "addListener" in WEBEXT?.inputMethodPrivate?.onImeMenuListChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnImeMenuListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onImeMenuListChanged.addListener);
    },
    "call_OnImeMenuListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onImeMenuListChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnImeMenuListChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onImeMenuListChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffImeMenuListChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onImeMenuListChanged &&
        "removeListener" in WEBEXT?.inputMethodPrivate?.onImeMenuListChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffImeMenuListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onImeMenuListChanged.removeListener);
    },
    "call_OffImeMenuListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onImeMenuListChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffImeMenuListChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onImeMenuListChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnImeMenuListChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onImeMenuListChanged &&
        "hasListener" in WEBEXT?.inputMethodPrivate?.onImeMenuListChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnImeMenuListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onImeMenuListChanged.hasListener);
    },
    "call_HasOnImeMenuListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onImeMenuListChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnImeMenuListChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onImeMenuListChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnInputMethodOptionsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onInputMethodOptionsChanged &&
        "addListener" in WEBEXT?.inputMethodPrivate?.onInputMethodOptionsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnInputMethodOptionsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.addListener);
    },
    "call_OnInputMethodOptionsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnInputMethodOptionsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffInputMethodOptionsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onInputMethodOptionsChanged &&
        "removeListener" in WEBEXT?.inputMethodPrivate?.onInputMethodOptionsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffInputMethodOptionsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.removeListener);
    },
    "call_OffInputMethodOptionsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffInputMethodOptionsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnInputMethodOptionsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onInputMethodOptionsChanged &&
        "hasListener" in WEBEXT?.inputMethodPrivate?.onInputMethodOptionsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnInputMethodOptionsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.hasListener);
    },
    "call_HasOnInputMethodOptionsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnInputMethodOptionsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnScreenProjectionChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onScreenProjectionChanged &&
        "addListener" in WEBEXT?.inputMethodPrivate?.onScreenProjectionChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnScreenProjectionChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onScreenProjectionChanged.addListener);
    },
    "call_OnScreenProjectionChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onScreenProjectionChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnScreenProjectionChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onScreenProjectionChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffScreenProjectionChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onScreenProjectionChanged &&
        "removeListener" in WEBEXT?.inputMethodPrivate?.onScreenProjectionChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffScreenProjectionChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onScreenProjectionChanged.removeListener);
    },
    "call_OffScreenProjectionChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onScreenProjectionChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffScreenProjectionChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onScreenProjectionChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnScreenProjectionChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onScreenProjectionChanged &&
        "hasListener" in WEBEXT?.inputMethodPrivate?.onScreenProjectionChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnScreenProjectionChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onScreenProjectionChanged.hasListener);
    },
    "call_HasOnScreenProjectionChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onScreenProjectionChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnScreenProjectionChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onScreenProjectionChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSettingsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onSettingsChanged &&
        "addListener" in WEBEXT?.inputMethodPrivate?.onSettingsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSettingsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onSettingsChanged.addListener);
    },
    "call_OnSettingsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onSettingsChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnSettingsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onSettingsChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSettingsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onSettingsChanged &&
        "removeListener" in WEBEXT?.inputMethodPrivate?.onSettingsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSettingsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onSettingsChanged.removeListener);
    },
    "call_OffSettingsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onSettingsChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffSettingsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onSettingsChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSettingsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onSettingsChanged &&
        "hasListener" in WEBEXT?.inputMethodPrivate?.onSettingsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSettingsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onSettingsChanged.hasListener);
    },
    "call_HasOnSettingsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onSettingsChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSettingsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onSettingsChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSuggestionsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onSuggestionsChanged &&
        "addListener" in WEBEXT?.inputMethodPrivate?.onSuggestionsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSuggestionsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onSuggestionsChanged.addListener);
    },
    "call_OnSuggestionsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onSuggestionsChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnSuggestionsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onSuggestionsChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSuggestionsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onSuggestionsChanged &&
        "removeListener" in WEBEXT?.inputMethodPrivate?.onSuggestionsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSuggestionsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onSuggestionsChanged.removeListener);
    },
    "call_OffSuggestionsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onSuggestionsChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffSuggestionsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onSuggestionsChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSuggestionsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.inputMethodPrivate?.onSuggestionsChanged &&
        "hasListener" in WEBEXT?.inputMethodPrivate?.onSuggestionsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSuggestionsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onSuggestionsChanged.hasListener);
    },
    "call_HasOnSuggestionsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onSuggestionsChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSuggestionsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onSuggestionsChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTouch": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate?.onTouch && "addListener" in WEBEXT?.inputMethodPrivate?.onTouch) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTouch": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onTouch.addListener);
    },
    "call_OnTouch": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onTouch.addListener(A.H.get<object>(callback));
    },
    "try_OnTouch": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onTouch.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffTouch": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate?.onTouch && "removeListener" in WEBEXT?.inputMethodPrivate?.onTouch) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffTouch": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onTouch.removeListener);
    },
    "call_OffTouch": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onTouch.removeListener(A.H.get<object>(callback));
    },
    "try_OffTouch": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onTouch.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnTouch": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate?.onTouch && "hasListener" in WEBEXT?.inputMethodPrivate?.onTouch) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnTouch": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.onTouch.hasListener);
    },
    "call_HasOnTouch": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.onTouch.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnTouch": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.onTouch.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenOptionsPage": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "openOptionsPage" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenOptionsPage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.openOptionsPage);
    },
    "call_OpenOptionsPage": (retPtr: Pointer, inputMethodId: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.openOptionsPage(A.H.get<object>(inputMethodId));
    },
    "try_OpenOptionsPage": (retPtr: Pointer, errPtr: Pointer, inputMethodId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.openOptionsPage(A.H.get<object>(inputMethodId));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Reset": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "reset" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Reset": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.reset);
    },
    "call_Reset": (retPtr: Pointer): void => {
      const _ret = WEBEXT.inputMethodPrivate.reset();
    },
    "try_Reset": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.reset();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetCompositionRange": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "setCompositionRange" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetCompositionRange": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.setCompositionRange);
    },
    "call_SetCompositionRange": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["contextID"] = A.load.Int64(parameters + 0);
      parameters_ffi["segments"] = A.load.Ref(parameters + 8, undefined);
      parameters_ffi["selectionAfter"] = A.load.Int64(parameters + 16);
      parameters_ffi["selectionBefore"] = A.load.Int64(parameters + 24);

      const _ret = WEBEXT.inputMethodPrivate.setCompositionRange(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetCompositionRange": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["contextID"] = A.load.Int64(parameters + 0);
        parameters_ffi["segments"] = A.load.Ref(parameters + 8, undefined);
        parameters_ffi["selectionAfter"] = A.load.Int64(parameters + 16);
        parameters_ffi["selectionBefore"] = A.load.Int64(parameters + 24);

        const _ret = WEBEXT.inputMethodPrivate.setCompositionRange(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetCurrentInputMethod": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "setCurrentInputMethod" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetCurrentInputMethod": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.setCurrentInputMethod);
    },
    "call_SetCurrentInputMethod": (retPtr: Pointer, inputMethodId: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.setCurrentInputMethod(A.H.get<object>(inputMethodId));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetCurrentInputMethod": (
      retPtr: Pointer,
      errPtr: Pointer,
      inputMethodId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.setCurrentInputMethod(A.H.get<object>(inputMethodId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetSettings": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "setSettings" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetSettings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.setSettings);
    },
    "call_SetSettings": (retPtr: Pointer, engineID: heap.Ref<object>, settings: Pointer): void => {
      const settings_ffi = {};

      if (A.load.Bool(settings + 108)) {
        settings_ffi["enableCompletion"] = A.load.Bool(settings + 0);
      }
      if (A.load.Bool(settings + 109)) {
        settings_ffi["enableDoubleSpacePeriod"] = A.load.Bool(settings + 1);
      }
      if (A.load.Bool(settings + 110)) {
        settings_ffi["enableGestureTyping"] = A.load.Bool(settings + 2);
      }
      if (A.load.Bool(settings + 111)) {
        settings_ffi["enablePrediction"] = A.load.Bool(settings + 3);
      }
      if (A.load.Bool(settings + 112)) {
        settings_ffi["enableSoundOnKeypress"] = A.load.Bool(settings + 4);
      }
      if (A.load.Bool(settings + 113)) {
        settings_ffi["koreanEnableSyllableInput"] = A.load.Bool(settings + 5);
      }
      settings_ffi["koreanKeyboardLayout"] = A.load.Ref(settings + 8, undefined);
      if (A.load.Bool(settings + 114)) {
        settings_ffi["koreanShowHangulCandidate"] = A.load.Bool(settings + 12);
      }
      if (A.load.Bool(settings + 115)) {
        settings_ffi["physicalKeyboardAutoCorrectionEnabledByDefault"] = A.load.Bool(settings + 13);
      }
      if (A.load.Bool(settings + 116)) {
        settings_ffi["physicalKeyboardAutoCorrectionLevel"] = A.load.Int64(settings + 16);
      }
      if (A.load.Bool(settings + 117)) {
        settings_ffi["physicalKeyboardEnableCapitalization"] = A.load.Bool(settings + 24);
      }
      if (A.load.Bool(settings + 118)) {
        settings_ffi["physicalKeyboardEnableDiacriticsOnLongpress"] = A.load.Bool(settings + 25);
      }
      if (A.load.Bool(settings + 119)) {
        settings_ffi["physicalKeyboardEnablePredictiveWriting"] = A.load.Bool(settings + 26);
      }
      if (A.load.Bool(settings + 120)) {
        settings_ffi["pinyinChinesePunctuation"] = A.load.Bool(settings + 27);
      }
      if (A.load.Bool(settings + 121)) {
        settings_ffi["pinyinDefaultChinese"] = A.load.Bool(settings + 28);
      }
      if (A.load.Bool(settings + 122)) {
        settings_ffi["pinyinEnableFuzzy"] = A.load.Bool(settings + 29);
      }
      if (A.load.Bool(settings + 123)) {
        settings_ffi["pinyinEnableLowerPaging"] = A.load.Bool(settings + 30);
      }
      if (A.load.Bool(settings + 124)) {
        settings_ffi["pinyinEnableUpperPaging"] = A.load.Bool(settings + 31);
      }
      if (A.load.Bool(settings + 125)) {
        settings_ffi["pinyinFullWidthCharacter"] = A.load.Bool(settings + 32);
      }
      if (A.load.Bool(settings + 33 + 24)) {
        settings_ffi["pinyinFuzzyConfig"] = {};
        if (A.load.Bool(settings + 33 + 12)) {
          settings_ffi["pinyinFuzzyConfig"]["an_ang"] = A.load.Bool(settings + 33 + 0);
        }
        if (A.load.Bool(settings + 33 + 13)) {
          settings_ffi["pinyinFuzzyConfig"]["c_ch"] = A.load.Bool(settings + 33 + 1);
        }
        if (A.load.Bool(settings + 33 + 14)) {
          settings_ffi["pinyinFuzzyConfig"]["en_eng"] = A.load.Bool(settings + 33 + 2);
        }
        if (A.load.Bool(settings + 33 + 15)) {
          settings_ffi["pinyinFuzzyConfig"]["f_h"] = A.load.Bool(settings + 33 + 3);
        }
        if (A.load.Bool(settings + 33 + 16)) {
          settings_ffi["pinyinFuzzyConfig"]["ian_iang"] = A.load.Bool(settings + 33 + 4);
        }
        if (A.load.Bool(settings + 33 + 17)) {
          settings_ffi["pinyinFuzzyConfig"]["in_ing"] = A.load.Bool(settings + 33 + 5);
        }
        if (A.load.Bool(settings + 33 + 18)) {
          settings_ffi["pinyinFuzzyConfig"]["k_g"] = A.load.Bool(settings + 33 + 6);
        }
        if (A.load.Bool(settings + 33 + 19)) {
          settings_ffi["pinyinFuzzyConfig"]["l_n"] = A.load.Bool(settings + 33 + 7);
        }
        if (A.load.Bool(settings + 33 + 20)) {
          settings_ffi["pinyinFuzzyConfig"]["r_l"] = A.load.Bool(settings + 33 + 8);
        }
        if (A.load.Bool(settings + 33 + 21)) {
          settings_ffi["pinyinFuzzyConfig"]["s_sh"] = A.load.Bool(settings + 33 + 9);
        }
        if (A.load.Bool(settings + 33 + 22)) {
          settings_ffi["pinyinFuzzyConfig"]["uan_uang"] = A.load.Bool(settings + 33 + 10);
        }
        if (A.load.Bool(settings + 33 + 23)) {
          settings_ffi["pinyinFuzzyConfig"]["z_zh"] = A.load.Bool(settings + 33 + 11);
        }
      }
      if (A.load.Bool(settings + 126)) {
        settings_ffi["vietnameseTelexAllowFlexibleDiacritics"] = A.load.Bool(settings + 58);
      }
      if (A.load.Bool(settings + 127)) {
        settings_ffi["vietnameseTelexInsertDoubleHornOnUo"] = A.load.Bool(settings + 59);
      }
      if (A.load.Bool(settings + 128)) {
        settings_ffi["vietnameseTelexInsertUHornOnW"] = A.load.Bool(settings + 60);
      }
      if (A.load.Bool(settings + 129)) {
        settings_ffi["vietnameseTelexNewStyleToneMarkPlacement"] = A.load.Bool(settings + 61);
      }
      if (A.load.Bool(settings + 130)) {
        settings_ffi["vietnameseTelexShowUnderline"] = A.load.Bool(settings + 62);
      }
      if (A.load.Bool(settings + 131)) {
        settings_ffi["vietnameseVniAllowFlexibleDiacritics"] = A.load.Bool(settings + 63);
      }
      if (A.load.Bool(settings + 132)) {
        settings_ffi["vietnameseVniInsertDoubleHornOnUo"] = A.load.Bool(settings + 64);
      }
      if (A.load.Bool(settings + 133)) {
        settings_ffi["vietnameseVniNewStyleToneMarkPlacement"] = A.load.Bool(settings + 65);
      }
      if (A.load.Bool(settings + 134)) {
        settings_ffi["vietnameseVniShowUnderline"] = A.load.Bool(settings + 66);
      }
      if (A.load.Bool(settings + 135)) {
        settings_ffi["virtualKeyboardAutoCorrectionLevel"] = A.load.Int64(settings + 72);
      }
      if (A.load.Bool(settings + 136)) {
        settings_ffi["virtualKeyboardEnableCapitalization"] = A.load.Bool(settings + 80);
      }
      settings_ffi["xkbLayout"] = A.load.Ref(settings + 84, undefined);
      settings_ffi["zhuyinKeyboardLayout"] = A.load.Ref(settings + 88, undefined);
      if (A.load.Bool(settings + 137)) {
        settings_ffi["zhuyinPageSize"] = A.load.Int64(settings + 96);
      }
      settings_ffi["zhuyinSelectKeys"] = A.load.Ref(settings + 104, undefined);

      const _ret = WEBEXT.inputMethodPrivate.setSettings(A.H.get<object>(engineID), settings_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetSettings": (
      retPtr: Pointer,
      errPtr: Pointer,
      engineID: heap.Ref<object>,
      settings: Pointer
    ): heap.Ref<boolean> => {
      try {
        const settings_ffi = {};

        if (A.load.Bool(settings + 108)) {
          settings_ffi["enableCompletion"] = A.load.Bool(settings + 0);
        }
        if (A.load.Bool(settings + 109)) {
          settings_ffi["enableDoubleSpacePeriod"] = A.load.Bool(settings + 1);
        }
        if (A.load.Bool(settings + 110)) {
          settings_ffi["enableGestureTyping"] = A.load.Bool(settings + 2);
        }
        if (A.load.Bool(settings + 111)) {
          settings_ffi["enablePrediction"] = A.load.Bool(settings + 3);
        }
        if (A.load.Bool(settings + 112)) {
          settings_ffi["enableSoundOnKeypress"] = A.load.Bool(settings + 4);
        }
        if (A.load.Bool(settings + 113)) {
          settings_ffi["koreanEnableSyllableInput"] = A.load.Bool(settings + 5);
        }
        settings_ffi["koreanKeyboardLayout"] = A.load.Ref(settings + 8, undefined);
        if (A.load.Bool(settings + 114)) {
          settings_ffi["koreanShowHangulCandidate"] = A.load.Bool(settings + 12);
        }
        if (A.load.Bool(settings + 115)) {
          settings_ffi["physicalKeyboardAutoCorrectionEnabledByDefault"] = A.load.Bool(settings + 13);
        }
        if (A.load.Bool(settings + 116)) {
          settings_ffi["physicalKeyboardAutoCorrectionLevel"] = A.load.Int64(settings + 16);
        }
        if (A.load.Bool(settings + 117)) {
          settings_ffi["physicalKeyboardEnableCapitalization"] = A.load.Bool(settings + 24);
        }
        if (A.load.Bool(settings + 118)) {
          settings_ffi["physicalKeyboardEnableDiacriticsOnLongpress"] = A.load.Bool(settings + 25);
        }
        if (A.load.Bool(settings + 119)) {
          settings_ffi["physicalKeyboardEnablePredictiveWriting"] = A.load.Bool(settings + 26);
        }
        if (A.load.Bool(settings + 120)) {
          settings_ffi["pinyinChinesePunctuation"] = A.load.Bool(settings + 27);
        }
        if (A.load.Bool(settings + 121)) {
          settings_ffi["pinyinDefaultChinese"] = A.load.Bool(settings + 28);
        }
        if (A.load.Bool(settings + 122)) {
          settings_ffi["pinyinEnableFuzzy"] = A.load.Bool(settings + 29);
        }
        if (A.load.Bool(settings + 123)) {
          settings_ffi["pinyinEnableLowerPaging"] = A.load.Bool(settings + 30);
        }
        if (A.load.Bool(settings + 124)) {
          settings_ffi["pinyinEnableUpperPaging"] = A.load.Bool(settings + 31);
        }
        if (A.load.Bool(settings + 125)) {
          settings_ffi["pinyinFullWidthCharacter"] = A.load.Bool(settings + 32);
        }
        if (A.load.Bool(settings + 33 + 24)) {
          settings_ffi["pinyinFuzzyConfig"] = {};
          if (A.load.Bool(settings + 33 + 12)) {
            settings_ffi["pinyinFuzzyConfig"]["an_ang"] = A.load.Bool(settings + 33 + 0);
          }
          if (A.load.Bool(settings + 33 + 13)) {
            settings_ffi["pinyinFuzzyConfig"]["c_ch"] = A.load.Bool(settings + 33 + 1);
          }
          if (A.load.Bool(settings + 33 + 14)) {
            settings_ffi["pinyinFuzzyConfig"]["en_eng"] = A.load.Bool(settings + 33 + 2);
          }
          if (A.load.Bool(settings + 33 + 15)) {
            settings_ffi["pinyinFuzzyConfig"]["f_h"] = A.load.Bool(settings + 33 + 3);
          }
          if (A.load.Bool(settings + 33 + 16)) {
            settings_ffi["pinyinFuzzyConfig"]["ian_iang"] = A.load.Bool(settings + 33 + 4);
          }
          if (A.load.Bool(settings + 33 + 17)) {
            settings_ffi["pinyinFuzzyConfig"]["in_ing"] = A.load.Bool(settings + 33 + 5);
          }
          if (A.load.Bool(settings + 33 + 18)) {
            settings_ffi["pinyinFuzzyConfig"]["k_g"] = A.load.Bool(settings + 33 + 6);
          }
          if (A.load.Bool(settings + 33 + 19)) {
            settings_ffi["pinyinFuzzyConfig"]["l_n"] = A.load.Bool(settings + 33 + 7);
          }
          if (A.load.Bool(settings + 33 + 20)) {
            settings_ffi["pinyinFuzzyConfig"]["r_l"] = A.load.Bool(settings + 33 + 8);
          }
          if (A.load.Bool(settings + 33 + 21)) {
            settings_ffi["pinyinFuzzyConfig"]["s_sh"] = A.load.Bool(settings + 33 + 9);
          }
          if (A.load.Bool(settings + 33 + 22)) {
            settings_ffi["pinyinFuzzyConfig"]["uan_uang"] = A.load.Bool(settings + 33 + 10);
          }
          if (A.load.Bool(settings + 33 + 23)) {
            settings_ffi["pinyinFuzzyConfig"]["z_zh"] = A.load.Bool(settings + 33 + 11);
          }
        }
        if (A.load.Bool(settings + 126)) {
          settings_ffi["vietnameseTelexAllowFlexibleDiacritics"] = A.load.Bool(settings + 58);
        }
        if (A.load.Bool(settings + 127)) {
          settings_ffi["vietnameseTelexInsertDoubleHornOnUo"] = A.load.Bool(settings + 59);
        }
        if (A.load.Bool(settings + 128)) {
          settings_ffi["vietnameseTelexInsertUHornOnW"] = A.load.Bool(settings + 60);
        }
        if (A.load.Bool(settings + 129)) {
          settings_ffi["vietnameseTelexNewStyleToneMarkPlacement"] = A.load.Bool(settings + 61);
        }
        if (A.load.Bool(settings + 130)) {
          settings_ffi["vietnameseTelexShowUnderline"] = A.load.Bool(settings + 62);
        }
        if (A.load.Bool(settings + 131)) {
          settings_ffi["vietnameseVniAllowFlexibleDiacritics"] = A.load.Bool(settings + 63);
        }
        if (A.load.Bool(settings + 132)) {
          settings_ffi["vietnameseVniInsertDoubleHornOnUo"] = A.load.Bool(settings + 64);
        }
        if (A.load.Bool(settings + 133)) {
          settings_ffi["vietnameseVniNewStyleToneMarkPlacement"] = A.load.Bool(settings + 65);
        }
        if (A.load.Bool(settings + 134)) {
          settings_ffi["vietnameseVniShowUnderline"] = A.load.Bool(settings + 66);
        }
        if (A.load.Bool(settings + 135)) {
          settings_ffi["virtualKeyboardAutoCorrectionLevel"] = A.load.Int64(settings + 72);
        }
        if (A.load.Bool(settings + 136)) {
          settings_ffi["virtualKeyboardEnableCapitalization"] = A.load.Bool(settings + 80);
        }
        settings_ffi["xkbLayout"] = A.load.Ref(settings + 84, undefined);
        settings_ffi["zhuyinKeyboardLayout"] = A.load.Ref(settings + 88, undefined);
        if (A.load.Bool(settings + 137)) {
          settings_ffi["zhuyinPageSize"] = A.load.Int64(settings + 96);
        }
        settings_ffi["zhuyinSelectKeys"] = A.load.Ref(settings + 104, undefined);

        const _ret = WEBEXT.inputMethodPrivate.setSettings(A.H.get<object>(engineID), settings_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetXkbLayout": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "setXkbLayout" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetXkbLayout": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.setXkbLayout);
    },
    "call_SetXkbLayout": (retPtr: Pointer, xkb_name: heap.Ref<object>): void => {
      const _ret = WEBEXT.inputMethodPrivate.setXkbLayout(A.H.get<object>(xkb_name));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetXkbLayout": (retPtr: Pointer, errPtr: Pointer, xkb_name: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.setXkbLayout(A.H.get<object>(xkb_name));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ShowInputView": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "showInputView" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ShowInputView": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.showInputView);
    },
    "call_ShowInputView": (retPtr: Pointer): void => {
      const _ret = WEBEXT.inputMethodPrivate.showInputView();
      A.store.Ref(retPtr, _ret);
    },
    "try_ShowInputView": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.showInputView();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SwitchToLastUsedInputMethod": (): heap.Ref<boolean> => {
      if (WEBEXT?.inputMethodPrivate && "switchToLastUsedInputMethod" in WEBEXT?.inputMethodPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SwitchToLastUsedInputMethod": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.inputMethodPrivate.switchToLastUsedInputMethod);
    },
    "call_SwitchToLastUsedInputMethod": (retPtr: Pointer): void => {
      const _ret = WEBEXT.inputMethodPrivate.switchToLastUsedInputMethod();
      A.store.Ref(retPtr, _ret);
    },
    "try_SwitchToLastUsedInputMethod": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.inputMethodPrivate.switchToLastUsedInputMethod();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
