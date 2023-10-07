import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/languagesettingsprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_InputMethod": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 22, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 18, false);
      } else {
        A.store.Bool(ptr + 22, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["displayName"]);
        A.store.Ref(ptr + 8, x["languageCodes"]);
        A.store.Ref(ptr + 12, x["tags"]);
        A.store.Bool(ptr + 19, "enabled" in x ? true : false);
        A.store.Bool(ptr + 16, x["enabled"] ? true : false);
        A.store.Bool(ptr + 20, "hasOptionsPage" in x ? true : false);
        A.store.Bool(ptr + 17, x["hasOptionsPage"] ? true : false);
        A.store.Bool(ptr + 21, "isProhibitedByPolicy" in x ? true : false);
        A.store.Bool(ptr + 18, x["isProhibitedByPolicy"] ? true : false);
      }
    },
    "load_InputMethod": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["displayName"] = A.load.Ref(ptr + 4, undefined);
      x["languageCodes"] = A.load.Ref(ptr + 8, undefined);
      x["tags"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 19)) {
        x["enabled"] = A.load.Bool(ptr + 16);
      } else {
        delete x["enabled"];
      }
      if (A.load.Bool(ptr + 20)) {
        x["hasOptionsPage"] = A.load.Bool(ptr + 17);
      } else {
        delete x["hasOptionsPage"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["isProhibitedByPolicy"] = A.load.Bool(ptr + 18);
      } else {
        delete x["isProhibitedByPolicy"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_InputMethodLists": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["componentExtensionImes"]);
        A.store.Ref(ptr + 4, x["thirdPartyExtensionImes"]);
      }
    },
    "load_InputMethodLists": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["componentExtensionImes"] = A.load.Ref(ptr + 0, undefined);
      x["thirdPartyExtensionImes"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Language": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 15, false);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Ref(ptr + 0, x["code"]);
        A.store.Ref(ptr + 4, x["displayName"]);
        A.store.Ref(ptr + 8, x["nativeDisplayName"]);
        A.store.Bool(ptr + 16, "supportsUI" in x ? true : false);
        A.store.Bool(ptr + 12, x["supportsUI"] ? true : false);
        A.store.Bool(ptr + 17, "supportsSpellcheck" in x ? true : false);
        A.store.Bool(ptr + 13, x["supportsSpellcheck"] ? true : false);
        A.store.Bool(ptr + 18, "supportsTranslate" in x ? true : false);
        A.store.Bool(ptr + 14, x["supportsTranslate"] ? true : false);
        A.store.Bool(ptr + 19, "isProhibitedLanguage" in x ? true : false);
        A.store.Bool(ptr + 15, x["isProhibitedLanguage"] ? true : false);
      }
    },
    "load_Language": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["code"] = A.load.Ref(ptr + 0, undefined);
      x["displayName"] = A.load.Ref(ptr + 4, undefined);
      x["nativeDisplayName"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["supportsUI"] = A.load.Bool(ptr + 12);
      } else {
        delete x["supportsUI"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["supportsSpellcheck"] = A.load.Bool(ptr + 13);
      } else {
        delete x["supportsSpellcheck"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["supportsTranslate"] = A.load.Bool(ptr + 14);
      } else {
        delete x["supportsTranslate"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["isProhibitedLanguage"] = A.load.Bool(ptr + 15);
      } else {
        delete x["isProhibitedLanguage"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SpellcheckDictionaryStatus": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 6, false);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Ref(ptr + 0, x["languageCode"]);
        A.store.Bool(ptr + 7, "isReady" in x ? true : false);
        A.store.Bool(ptr + 4, x["isReady"] ? true : false);
        A.store.Bool(ptr + 8, "isDownloading" in x ? true : false);
        A.store.Bool(ptr + 5, x["isDownloading"] ? true : false);
        A.store.Bool(ptr + 9, "downloadFailed" in x ? true : false);
        A.store.Bool(ptr + 6, x["downloadFailed"] ? true : false);
      }
    },
    "load_SpellcheckDictionaryStatus": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["languageCode"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 7)) {
        x["isReady"] = A.load.Bool(ptr + 4);
      } else {
        delete x["isReady"];
      }
      if (A.load.Bool(ptr + 8)) {
        x["isDownloading"] = A.load.Bool(ptr + 5);
      } else {
        delete x["isDownloading"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["downloadFailed"] = A.load.Bool(ptr + 6);
      } else {
        delete x["downloadFailed"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_MoveType": (ref: heap.Ref<string>): number => {
      const idx = ["TOP", "UP", "DOWN", "UNKNOWN"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_AddInputMethod": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "addInputMethod" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddInputMethod": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.addInputMethod);
    },
    "call_AddInputMethod": (retPtr: Pointer, inputMethodId: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.addInputMethod(A.H.get<object>(inputMethodId));
    },
    "try_AddInputMethod": (retPtr: Pointer, errPtr: Pointer, inputMethodId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.addInputMethod(A.H.get<object>(inputMethodId));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AddSpellcheckWord": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "addSpellcheckWord" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddSpellcheckWord": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.addSpellcheckWord);
    },
    "call_AddSpellcheckWord": (retPtr: Pointer, word: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.addSpellcheckWord(A.H.get<object>(word));
    },
    "try_AddSpellcheckWord": (retPtr: Pointer, errPtr: Pointer, word: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.addSpellcheckWord(A.H.get<object>(word));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DisableLanguage": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "disableLanguage" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DisableLanguage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.disableLanguage);
    },
    "call_DisableLanguage": (retPtr: Pointer, languageCode: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.disableLanguage(A.H.get<object>(languageCode));
    },
    "try_DisableLanguage": (retPtr: Pointer, errPtr: Pointer, languageCode: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.disableLanguage(A.H.get<object>(languageCode));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_EnableLanguage": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "enableLanguage" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EnableLanguage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.enableLanguage);
    },
    "call_EnableLanguage": (retPtr: Pointer, languageCode: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.enableLanguage(A.H.get<object>(languageCode));
    },
    "try_EnableLanguage": (retPtr: Pointer, errPtr: Pointer, languageCode: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.enableLanguage(A.H.get<object>(languageCode));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAlwaysTranslateLanguages": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "getAlwaysTranslateLanguages" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAlwaysTranslateLanguages": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.getAlwaysTranslateLanguages);
    },
    "call_GetAlwaysTranslateLanguages": (retPtr: Pointer): void => {
      const _ret = WEBEXT.languageSettingsPrivate.getAlwaysTranslateLanguages();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAlwaysTranslateLanguages": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.getAlwaysTranslateLanguages();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetInputMethodLists": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "getInputMethodLists" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetInputMethodLists": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.getInputMethodLists);
    },
    "call_GetInputMethodLists": (retPtr: Pointer): void => {
      const _ret = WEBEXT.languageSettingsPrivate.getInputMethodLists();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetInputMethodLists": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.getInputMethodLists();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetLanguageList": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "getLanguageList" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetLanguageList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.getLanguageList);
    },
    "call_GetLanguageList": (retPtr: Pointer): void => {
      const _ret = WEBEXT.languageSettingsPrivate.getLanguageList();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetLanguageList": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.getLanguageList();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetNeverTranslateLanguages": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "getNeverTranslateLanguages" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetNeverTranslateLanguages": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.getNeverTranslateLanguages);
    },
    "call_GetNeverTranslateLanguages": (retPtr: Pointer): void => {
      const _ret = WEBEXT.languageSettingsPrivate.getNeverTranslateLanguages();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetNeverTranslateLanguages": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.getNeverTranslateLanguages();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSpellcheckDictionaryStatuses": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "getSpellcheckDictionaryStatuses" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSpellcheckDictionaryStatuses": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.getSpellcheckDictionaryStatuses);
    },
    "call_GetSpellcheckDictionaryStatuses": (retPtr: Pointer): void => {
      const _ret = WEBEXT.languageSettingsPrivate.getSpellcheckDictionaryStatuses();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSpellcheckDictionaryStatuses": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.getSpellcheckDictionaryStatuses();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSpellcheckWords": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "getSpellcheckWords" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSpellcheckWords": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.getSpellcheckWords);
    },
    "call_GetSpellcheckWords": (retPtr: Pointer): void => {
      const _ret = WEBEXT.languageSettingsPrivate.getSpellcheckWords();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSpellcheckWords": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.getSpellcheckWords();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetTranslateTargetLanguage": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "getTranslateTargetLanguage" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetTranslateTargetLanguage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.getTranslateTargetLanguage);
    },
    "call_GetTranslateTargetLanguage": (retPtr: Pointer): void => {
      const _ret = WEBEXT.languageSettingsPrivate.getTranslateTargetLanguage();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetTranslateTargetLanguage": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.getTranslateTargetLanguage();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_MoveLanguage": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "moveLanguage" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MoveLanguage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.moveLanguage);
    },
    "call_MoveLanguage": (retPtr: Pointer, languageCode: heap.Ref<object>, moveType: number): void => {
      const _ret = WEBEXT.languageSettingsPrivate.moveLanguage(
        A.H.get<object>(languageCode),
        moveType > 0 && moveType <= 4 ? ["TOP", "UP", "DOWN", "UNKNOWN"][moveType - 1] : undefined
      );
    },
    "try_MoveLanguage": (
      retPtr: Pointer,
      errPtr: Pointer,
      languageCode: heap.Ref<object>,
      moveType: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.moveLanguage(
          A.H.get<object>(languageCode),
          moveType > 0 && moveType <= 4 ? ["TOP", "UP", "DOWN", "UNKNOWN"][moveType - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCustomDictionaryChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.languageSettingsPrivate?.onCustomDictionaryChanged &&
        "addListener" in WEBEXT?.languageSettingsPrivate?.onCustomDictionaryChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCustomDictionaryChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.addListener);
    },
    "call_OnCustomDictionaryChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnCustomDictionaryChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCustomDictionaryChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.languageSettingsPrivate?.onCustomDictionaryChanged &&
        "removeListener" in WEBEXT?.languageSettingsPrivate?.onCustomDictionaryChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCustomDictionaryChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.removeListener);
    },
    "call_OffCustomDictionaryChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffCustomDictionaryChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCustomDictionaryChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.languageSettingsPrivate?.onCustomDictionaryChanged &&
        "hasListener" in WEBEXT?.languageSettingsPrivate?.onCustomDictionaryChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCustomDictionaryChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.hasListener);
    },
    "call_HasOnCustomDictionaryChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCustomDictionaryChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnInputMethodAdded": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.languageSettingsPrivate?.onInputMethodAdded &&
        "addListener" in WEBEXT?.languageSettingsPrivate?.onInputMethodAdded
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnInputMethodAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.onInputMethodAdded.addListener);
    },
    "call_OnInputMethodAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.onInputMethodAdded.addListener(A.H.get<object>(callback));
    },
    "try_OnInputMethodAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.onInputMethodAdded.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffInputMethodAdded": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.languageSettingsPrivate?.onInputMethodAdded &&
        "removeListener" in WEBEXT?.languageSettingsPrivate?.onInputMethodAdded
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffInputMethodAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.onInputMethodAdded.removeListener);
    },
    "call_OffInputMethodAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.onInputMethodAdded.removeListener(A.H.get<object>(callback));
    },
    "try_OffInputMethodAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.onInputMethodAdded.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnInputMethodAdded": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.languageSettingsPrivate?.onInputMethodAdded &&
        "hasListener" in WEBEXT?.languageSettingsPrivate?.onInputMethodAdded
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnInputMethodAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.onInputMethodAdded.hasListener);
    },
    "call_HasOnInputMethodAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.onInputMethodAdded.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnInputMethodAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.onInputMethodAdded.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnInputMethodRemoved": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.languageSettingsPrivate?.onInputMethodRemoved &&
        "addListener" in WEBEXT?.languageSettingsPrivate?.onInputMethodRemoved
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnInputMethodRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.onInputMethodRemoved.addListener);
    },
    "call_OnInputMethodRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.onInputMethodRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnInputMethodRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.onInputMethodRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffInputMethodRemoved": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.languageSettingsPrivate?.onInputMethodRemoved &&
        "removeListener" in WEBEXT?.languageSettingsPrivate?.onInputMethodRemoved
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffInputMethodRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.onInputMethodRemoved.removeListener);
    },
    "call_OffInputMethodRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.onInputMethodRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffInputMethodRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.onInputMethodRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnInputMethodRemoved": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.languageSettingsPrivate?.onInputMethodRemoved &&
        "hasListener" in WEBEXT?.languageSettingsPrivate?.onInputMethodRemoved
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnInputMethodRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.onInputMethodRemoved.hasListener);
    },
    "call_HasOnInputMethodRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.onInputMethodRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnInputMethodRemoved": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.onInputMethodRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSpellcheckDictionariesChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.languageSettingsPrivate?.onSpellcheckDictionariesChanged &&
        "addListener" in WEBEXT?.languageSettingsPrivate?.onSpellcheckDictionariesChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSpellcheckDictionariesChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.addListener);
    },
    "call_OnSpellcheckDictionariesChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.addListener(
        A.H.get<object>(callback)
      );
    },
    "try_OnSpellcheckDictionariesChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.addListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSpellcheckDictionariesChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.languageSettingsPrivate?.onSpellcheckDictionariesChanged &&
        "removeListener" in WEBEXT?.languageSettingsPrivate?.onSpellcheckDictionariesChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSpellcheckDictionariesChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.removeListener);
    },
    "call_OffSpellcheckDictionariesChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.removeListener(
        A.H.get<object>(callback)
      );
    },
    "try_OffSpellcheckDictionariesChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSpellcheckDictionariesChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.languageSettingsPrivate?.onSpellcheckDictionariesChanged &&
        "hasListener" in WEBEXT?.languageSettingsPrivate?.onSpellcheckDictionariesChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSpellcheckDictionariesChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.hasListener);
    },
    "call_HasOnSpellcheckDictionariesChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.hasListener(
        A.H.get<object>(callback)
      );
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSpellcheckDictionariesChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.hasListener(
          A.H.get<object>(callback)
        );
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveInputMethod": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "removeInputMethod" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveInputMethod": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.removeInputMethod);
    },
    "call_RemoveInputMethod": (retPtr: Pointer, inputMethodId: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.removeInputMethod(A.H.get<object>(inputMethodId));
    },
    "try_RemoveInputMethod": (retPtr: Pointer, errPtr: Pointer, inputMethodId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.removeInputMethod(A.H.get<object>(inputMethodId));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveSpellcheckWord": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "removeSpellcheckWord" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveSpellcheckWord": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.removeSpellcheckWord);
    },
    "call_RemoveSpellcheckWord": (retPtr: Pointer, word: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.removeSpellcheckWord(A.H.get<object>(word));
    },
    "try_RemoveSpellcheckWord": (retPtr: Pointer, errPtr: Pointer, word: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.removeSpellcheckWord(A.H.get<object>(word));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RetryDownloadDictionary": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "retryDownloadDictionary" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RetryDownloadDictionary": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.retryDownloadDictionary);
    },
    "call_RetryDownloadDictionary": (retPtr: Pointer, languageCode: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.retryDownloadDictionary(A.H.get<object>(languageCode));
    },
    "try_RetryDownloadDictionary": (
      retPtr: Pointer,
      errPtr: Pointer,
      languageCode: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.retryDownloadDictionary(A.H.get<object>(languageCode));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetEnableTranslationForLanguage": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "setEnableTranslationForLanguage" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetEnableTranslationForLanguage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.setEnableTranslationForLanguage);
    },
    "call_SetEnableTranslationForLanguage": (
      retPtr: Pointer,
      languageCode: heap.Ref<object>,
      enable: heap.Ref<boolean>
    ): void => {
      const _ret = WEBEXT.languageSettingsPrivate.setEnableTranslationForLanguage(
        A.H.get<object>(languageCode),
        enable === A.H.TRUE
      );
    },
    "try_SetEnableTranslationForLanguage": (
      retPtr: Pointer,
      errPtr: Pointer,
      languageCode: heap.Ref<object>,
      enable: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.setEnableTranslationForLanguage(
          A.H.get<object>(languageCode),
          enable === A.H.TRUE
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetLanguageAlwaysTranslateState": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "setLanguageAlwaysTranslateState" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetLanguageAlwaysTranslateState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.setLanguageAlwaysTranslateState);
    },
    "call_SetLanguageAlwaysTranslateState": (
      retPtr: Pointer,
      languageCode: heap.Ref<object>,
      alwaysTranslate: heap.Ref<boolean>
    ): void => {
      const _ret = WEBEXT.languageSettingsPrivate.setLanguageAlwaysTranslateState(
        A.H.get<object>(languageCode),
        alwaysTranslate === A.H.TRUE
      );
    },
    "try_SetLanguageAlwaysTranslateState": (
      retPtr: Pointer,
      errPtr: Pointer,
      languageCode: heap.Ref<object>,
      alwaysTranslate: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.setLanguageAlwaysTranslateState(
          A.H.get<object>(languageCode),
          alwaysTranslate === A.H.TRUE
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetTranslateTargetLanguage": (): heap.Ref<boolean> => {
      if (WEBEXT?.languageSettingsPrivate && "setTranslateTargetLanguage" in WEBEXT?.languageSettingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetTranslateTargetLanguage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.languageSettingsPrivate.setTranslateTargetLanguage);
    },
    "call_SetTranslateTargetLanguage": (retPtr: Pointer, languageCode: heap.Ref<object>): void => {
      const _ret = WEBEXT.languageSettingsPrivate.setTranslateTargetLanguage(A.H.get<object>(languageCode));
    },
    "try_SetTranslateTargetLanguage": (
      retPtr: Pointer,
      errPtr: Pointer,
      languageCode: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.languageSettingsPrivate.setTranslateTargetLanguage(A.H.get<object>(languageCode));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
