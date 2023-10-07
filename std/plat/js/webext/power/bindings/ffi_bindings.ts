import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/power", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_Level": (ref: heap.Ref<string>): number => {
      const idx = ["system", "display"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_ReleaseKeepAwake": (): heap.Ref<boolean> => {
      if (WEBEXT?.power && "releaseKeepAwake" in WEBEXT?.power) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReleaseKeepAwake": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.power.releaseKeepAwake);
    },
    "call_ReleaseKeepAwake": (retPtr: Pointer): void => {
      const _ret = WEBEXT.power.releaseKeepAwake();
    },
    "try_ReleaseKeepAwake": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.power.releaseKeepAwake();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReportActivity": (): heap.Ref<boolean> => {
      if (WEBEXT?.power && "reportActivity" in WEBEXT?.power) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportActivity": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.power.reportActivity);
    },
    "call_ReportActivity": (retPtr: Pointer): void => {
      const _ret = WEBEXT.power.reportActivity();
      A.store.Ref(retPtr, _ret);
    },
    "try_ReportActivity": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.power.reportActivity();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RequestKeepAwake": (): heap.Ref<boolean> => {
      if (WEBEXT?.power && "requestKeepAwake" in WEBEXT?.power) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RequestKeepAwake": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.power.requestKeepAwake);
    },
    "call_RequestKeepAwake": (retPtr: Pointer, level: number): void => {
      const _ret = WEBEXT.power.requestKeepAwake(
        level > 0 && level <= 2 ? ["system", "display"][level - 1] : undefined
      );
    },
    "try_RequestKeepAwake": (retPtr: Pointer, errPtr: Pointer, level: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.power.requestKeepAwake(
          level > 0 && level <= 2 ? ["system", "display"][level - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
