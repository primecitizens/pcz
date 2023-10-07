import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/virtualkeyboard", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_FeatureRestrictions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 5, "autoCompleteEnabled" in x ? true : false);
        A.store.Bool(ptr + 0, x["autoCompleteEnabled"] ? true : false);
        A.store.Bool(ptr + 6, "autoCorrectEnabled" in x ? true : false);
        A.store.Bool(ptr + 1, x["autoCorrectEnabled"] ? true : false);
        A.store.Bool(ptr + 7, "handwritingEnabled" in x ? true : false);
        A.store.Bool(ptr + 2, x["handwritingEnabled"] ? true : false);
        A.store.Bool(ptr + 8, "spellCheckEnabled" in x ? true : false);
        A.store.Bool(ptr + 3, x["spellCheckEnabled"] ? true : false);
        A.store.Bool(ptr + 9, "voiceInputEnabled" in x ? true : false);
        A.store.Bool(ptr + 4, x["voiceInputEnabled"] ? true : false);
      }
    },
    "load_FeatureRestrictions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 5)) {
        x["autoCompleteEnabled"] = A.load.Bool(ptr + 0);
      } else {
        delete x["autoCompleteEnabled"];
      }
      if (A.load.Bool(ptr + 6)) {
        x["autoCorrectEnabled"] = A.load.Bool(ptr + 1);
      } else {
        delete x["autoCorrectEnabled"];
      }
      if (A.load.Bool(ptr + 7)) {
        x["handwritingEnabled"] = A.load.Bool(ptr + 2);
      } else {
        delete x["handwritingEnabled"];
      }
      if (A.load.Bool(ptr + 8)) {
        x["spellCheckEnabled"] = A.load.Bool(ptr + 3);
      } else {
        delete x["spellCheckEnabled"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["voiceInputEnabled"] = A.load.Bool(ptr + 4);
      } else {
        delete x["voiceInputEnabled"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_RestrictFeatures": (): heap.Ref<boolean> => {
      if (WEBEXT?.virtualKeyboard && "restrictFeatures" in WEBEXT?.virtualKeyboard) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RestrictFeatures": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.virtualKeyboard.restrictFeatures);
    },
    "call_RestrictFeatures": (retPtr: Pointer, restrictions: Pointer): void => {
      const restrictions_ffi = {};

      if (A.load.Bool(restrictions + 5)) {
        restrictions_ffi["autoCompleteEnabled"] = A.load.Bool(restrictions + 0);
      }
      if (A.load.Bool(restrictions + 6)) {
        restrictions_ffi["autoCorrectEnabled"] = A.load.Bool(restrictions + 1);
      }
      if (A.load.Bool(restrictions + 7)) {
        restrictions_ffi["handwritingEnabled"] = A.load.Bool(restrictions + 2);
      }
      if (A.load.Bool(restrictions + 8)) {
        restrictions_ffi["spellCheckEnabled"] = A.load.Bool(restrictions + 3);
      }
      if (A.load.Bool(restrictions + 9)) {
        restrictions_ffi["voiceInputEnabled"] = A.load.Bool(restrictions + 4);
      }

      const _ret = WEBEXT.virtualKeyboard.restrictFeatures(restrictions_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RestrictFeatures": (retPtr: Pointer, errPtr: Pointer, restrictions: Pointer): heap.Ref<boolean> => {
      try {
        const restrictions_ffi = {};

        if (A.load.Bool(restrictions + 5)) {
          restrictions_ffi["autoCompleteEnabled"] = A.load.Bool(restrictions + 0);
        }
        if (A.load.Bool(restrictions + 6)) {
          restrictions_ffi["autoCorrectEnabled"] = A.load.Bool(restrictions + 1);
        }
        if (A.load.Bool(restrictions + 7)) {
          restrictions_ffi["handwritingEnabled"] = A.load.Bool(restrictions + 2);
        }
        if (A.load.Bool(restrictions + 8)) {
          restrictions_ffi["spellCheckEnabled"] = A.load.Bool(restrictions + 3);
        }
        if (A.load.Bool(restrictions + 9)) {
          restrictions_ffi["voiceInputEnabled"] = A.load.Bool(restrictions + 4);
        }

        const _ret = WEBEXT.virtualKeyboard.restrictFeatures(restrictions_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
