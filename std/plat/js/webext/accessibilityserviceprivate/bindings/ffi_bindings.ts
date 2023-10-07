import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/accessibilityserviceprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_SpeakSelectedText": (): heap.Ref<boolean> => {
      if (WEBEXT?.accessibilityServicePrivate && "speakSelectedText" in WEBEXT?.accessibilityServicePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SpeakSelectedText": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.accessibilityServicePrivate.speakSelectedText);
    },
    "call_SpeakSelectedText": (retPtr: Pointer): void => {
      const _ret = WEBEXT.accessibilityServicePrivate.speakSelectedText();
      A.store.Ref(retPtr, _ret);
    },
    "try_SpeakSelectedText": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.accessibilityServicePrivate.speakSelectedText();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
