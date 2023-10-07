import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/enterprise/hardwareplatform", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_HardwarePlatformInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["model"]);
        A.store.Ref(ptr + 4, x["manufacturer"]);
      }
    },
    "load_HardwarePlatformInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["model"] = A.load.Ref(ptr + 0, undefined);
      x["manufacturer"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetHardwarePlatformInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.hardwarePlatform && "getHardwarePlatformInfo" in WEBEXT?.enterprise?.hardwarePlatform) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetHardwarePlatformInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.hardwarePlatform.getHardwarePlatformInfo);
    },
    "call_GetHardwarePlatformInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.enterprise.hardwarePlatform.getHardwarePlatformInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetHardwarePlatformInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.hardwarePlatform.getHardwarePlatformInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
