import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/enterprise/networkingattributes", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_NetworkDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["macAddress"]);
        A.store.Ref(ptr + 4, x["ipv4"]);
        A.store.Ref(ptr + 8, x["ipv6"]);
      }
    },
    "load_NetworkDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["macAddress"] = A.load.Ref(ptr + 0, undefined);
      x["ipv4"] = A.load.Ref(ptr + 4, undefined);
      x["ipv6"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetNetworkDetails": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.networkingAttributes && "getNetworkDetails" in WEBEXT?.enterprise?.networkingAttributes) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetNetworkDetails": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.networkingAttributes.getNetworkDetails);
    },
    "call_GetNetworkDetails": (retPtr: Pointer): void => {
      const _ret = WEBEXT.enterprise.networkingAttributes.getNetworkDetails();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetNetworkDetails": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.networkingAttributes.getNetworkDetails();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
