import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/dns", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_ResolveCallbackResolveInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "resultCode" in x ? true : false);
        A.store.Int32(ptr + 0, x["resultCode"] === undefined ? 0 : (x["resultCode"] as number));
        A.store.Ref(ptr + 4, x["address"]);
      }
    },
    "load_ResolveCallbackResolveInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["resultCode"] = A.load.Int32(ptr + 0);
      } else {
        delete x["resultCode"];
      }
      x["address"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Resolve": (): heap.Ref<boolean> => {
      if (WEBEXT?.dns && "resolve" in WEBEXT?.dns) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Resolve": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.dns.resolve);
    },
    "call_Resolve": (retPtr: Pointer, hostname: heap.Ref<object>): void => {
      const _ret = WEBEXT.dns.resolve(A.H.get<object>(hostname));
      A.store.Ref(retPtr, _ret);
    },
    "try_Resolve": (retPtr: Pointer, errPtr: Pointer, hostname: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.dns.resolve(A.H.get<object>(hostname));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
