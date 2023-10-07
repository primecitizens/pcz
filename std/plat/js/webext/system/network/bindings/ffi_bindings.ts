import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/system/network", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_NetworkInterface": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["address"]);
        A.store.Bool(ptr + 12, "prefixLength" in x ? true : false);
        A.store.Int32(ptr + 8, x["prefixLength"] === undefined ? 0 : (x["prefixLength"] as number));
      }
    },
    "load_NetworkInterface": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["address"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["prefixLength"] = A.load.Int32(ptr + 8);
      } else {
        delete x["prefixLength"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetNetworkInterfaces": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.network && "getNetworkInterfaces" in WEBEXT?.system?.network) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetNetworkInterfaces": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.network.getNetworkInterfaces);
    },
    "call_GetNetworkInterfaces": (retPtr: Pointer): void => {
      const _ret = WEBEXT.system.network.getNetworkInterfaces();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetNetworkInterfaces": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.network.getNetworkInterfaces();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
