import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/systemlog", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_MessageOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["message"]);
      }
    },
    "load_MessageOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["message"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Add": (): heap.Ref<boolean> => {
      if (WEBEXT?.systemLog && "add" in WEBEXT?.systemLog) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Add": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.systemLog.add);
    },
    "call_Add": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["message"] = A.load.Ref(options + 0, undefined);

      const _ret = WEBEXT.systemLog.add(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Add": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["message"] = A.load.Ref(options + 0, undefined);

        const _ret = WEBEXT.systemLog.add(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
