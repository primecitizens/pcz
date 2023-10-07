import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/commandlineprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_HasSwitch": (): heap.Ref<boolean> => {
      if (WEBEXT?.commandLinePrivate && "hasSwitch" in WEBEXT?.commandLinePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasSwitch": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.commandLinePrivate.hasSwitch);
    },
    "call_HasSwitch": (retPtr: Pointer, name: heap.Ref<object>): void => {
      const _ret = WEBEXT.commandLinePrivate.hasSwitch(A.H.get<object>(name));
      A.store.Ref(retPtr, _ret);
    },
    "try_HasSwitch": (retPtr: Pointer, errPtr: Pointer, name: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.commandLinePrivate.hasSwitch(A.H.get<object>(name));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
