import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/i18n", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_DetectLanguageReturnTypeFieldLanguagesElem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["language"]);
        A.store.Int64(ptr + 8, x["percentage"] === undefined ? 0 : (x["percentage"] as number));
      }
    },
    "load_DetectLanguageReturnTypeFieldLanguagesElem": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["language"] = A.load.Ref(ptr + 0, undefined);
      x["percentage"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DetectLanguageReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Bool(ptr + 0, x["isReliable"] ? true : false);
        A.store.Ref(ptr + 4, x["languages"]);
      }
    },
    "load_DetectLanguageReturnType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["isReliable"] = A.load.Bool(ptr + 0);
      x["languages"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetMessageArgOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 1, "escapeLt" in x ? true : false);
        A.store.Bool(ptr + 0, x["escapeLt"] ? true : false);
      }
    },
    "load_GetMessageArgOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 1)) {
        x["escapeLt"] = A.load.Bool(ptr + 0);
      } else {
        delete x["escapeLt"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_DetectLanguage": (): heap.Ref<boolean> => {
      if (WEBEXT?.i18n && "detectLanguage" in WEBEXT?.i18n) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DetectLanguage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.i18n.detectLanguage);
    },
    "call_DetectLanguage": (retPtr: Pointer, text: heap.Ref<object>): void => {
      const _ret = WEBEXT.i18n.detectLanguage(A.H.get<object>(text));
      A.store.Ref(retPtr, _ret);
    },
    "try_DetectLanguage": (retPtr: Pointer, errPtr: Pointer, text: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.i18n.detectLanguage(A.H.get<object>(text));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAcceptLanguages": (): heap.Ref<boolean> => {
      if (WEBEXT?.i18n && "getAcceptLanguages" in WEBEXT?.i18n) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAcceptLanguages": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.i18n.getAcceptLanguages);
    },
    "call_GetAcceptLanguages": (retPtr: Pointer): void => {
      const _ret = WEBEXT.i18n.getAcceptLanguages();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAcceptLanguages": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.i18n.getAcceptLanguages();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.i18n && "getMessage" in WEBEXT?.i18n) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.i18n.getMessage);
    },
    "call_GetMessage": (
      retPtr: Pointer,
      messageName: heap.Ref<object>,
      substitutions: heap.Ref<object>,
      options: Pointer
    ): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 1)) {
        options_ffi["escapeLt"] = A.load.Bool(options + 0);
      }

      const _ret = WEBEXT.i18n.getMessage(A.H.get<object>(messageName), A.H.get<object>(substitutions), options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetMessage": (
      retPtr: Pointer,
      errPtr: Pointer,
      messageName: heap.Ref<object>,
      substitutions: heap.Ref<object>,
      options: Pointer
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 1)) {
          options_ffi["escapeLt"] = A.load.Bool(options + 0);
        }

        const _ret = WEBEXT.i18n.getMessage(A.H.get<object>(messageName), A.H.get<object>(substitutions), options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetUILanguage": (): heap.Ref<boolean> => {
      if (WEBEXT?.i18n && "getUILanguage" in WEBEXT?.i18n) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetUILanguage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.i18n.getUILanguage);
    },
    "call_GetUILanguage": (retPtr: Pointer): void => {
      const _ret = WEBEXT.i18n.getUILanguage();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetUILanguage": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.i18n.getUILanguage();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
