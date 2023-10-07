import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/system/memory", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_MemoryInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Bool(ptr + 16, "capacity" in x ? true : false);
        A.store.Float64(ptr + 0, x["capacity"] === undefined ? 0 : (x["capacity"] as number));
        A.store.Bool(ptr + 17, "availableCapacity" in x ? true : false);
        A.store.Float64(ptr + 8, x["availableCapacity"] === undefined ? 0 : (x["availableCapacity"] as number));
      }
    },
    "load_MemoryInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["capacity"] = A.load.Float64(ptr + 0);
      } else {
        delete x["capacity"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["availableCapacity"] = A.load.Float64(ptr + 8);
      } else {
        delete x["availableCapacity"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.memory && "getInfo" in WEBEXT?.system?.memory) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.memory.getInfo);
    },
    "call_GetInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.system.memory.getInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.memory.getInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
