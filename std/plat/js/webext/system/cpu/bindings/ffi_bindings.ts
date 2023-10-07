import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/system/cpu", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_CpuTime": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 36, false);
        A.store.Bool(ptr + 32, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 33, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 34, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 35, false);
        A.store.Float64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 36, true);
        A.store.Bool(ptr + 32, "user" in x ? true : false);
        A.store.Float64(ptr + 0, x["user"] === undefined ? 0 : (x["user"] as number));
        A.store.Bool(ptr + 33, "kernel" in x ? true : false);
        A.store.Float64(ptr + 8, x["kernel"] === undefined ? 0 : (x["kernel"] as number));
        A.store.Bool(ptr + 34, "idle" in x ? true : false);
        A.store.Float64(ptr + 16, x["idle"] === undefined ? 0 : (x["idle"] as number));
        A.store.Bool(ptr + 35, "total" in x ? true : false);
        A.store.Float64(ptr + 24, x["total"] === undefined ? 0 : (x["total"] as number));
      }
    },
    "load_CpuTime": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 32)) {
        x["user"] = A.load.Float64(ptr + 0);
      } else {
        delete x["user"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["kernel"] = A.load.Float64(ptr + 8);
      } else {
        delete x["kernel"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["idle"] = A.load.Float64(ptr + 16);
      } else {
        delete x["idle"];
      }
      if (A.load.Bool(ptr + 35)) {
        x["total"] = A.load.Float64(ptr + 24);
      } else {
        delete x["total"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ProcessorInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 37, false);

        A.store.Bool(ptr + 0 + 36, false);
        A.store.Bool(ptr + 0 + 32, false);
        A.store.Float64(ptr + 0 + 0, 0);
        A.store.Bool(ptr + 0 + 33, false);
        A.store.Float64(ptr + 0 + 8, 0);
        A.store.Bool(ptr + 0 + 34, false);
        A.store.Float64(ptr + 0 + 16, 0);
        A.store.Bool(ptr + 0 + 35, false);
        A.store.Float64(ptr + 0 + 24, 0);
      } else {
        A.store.Bool(ptr + 37, true);

        if (typeof x["usage"] === "undefined") {
          A.store.Bool(ptr + 0 + 36, false);
          A.store.Bool(ptr + 0 + 32, false);
          A.store.Float64(ptr + 0 + 0, 0);
          A.store.Bool(ptr + 0 + 33, false);
          A.store.Float64(ptr + 0 + 8, 0);
          A.store.Bool(ptr + 0 + 34, false);
          A.store.Float64(ptr + 0 + 16, 0);
          A.store.Bool(ptr + 0 + 35, false);
          A.store.Float64(ptr + 0 + 24, 0);
        } else {
          A.store.Bool(ptr + 0 + 36, true);
          A.store.Bool(ptr + 0 + 32, "user" in x["usage"] ? true : false);
          A.store.Float64(ptr + 0 + 0, x["usage"]["user"] === undefined ? 0 : (x["usage"]["user"] as number));
          A.store.Bool(ptr + 0 + 33, "kernel" in x["usage"] ? true : false);
          A.store.Float64(ptr + 0 + 8, x["usage"]["kernel"] === undefined ? 0 : (x["usage"]["kernel"] as number));
          A.store.Bool(ptr + 0 + 34, "idle" in x["usage"] ? true : false);
          A.store.Float64(ptr + 0 + 16, x["usage"]["idle"] === undefined ? 0 : (x["usage"]["idle"] as number));
          A.store.Bool(ptr + 0 + 35, "total" in x["usage"] ? true : false);
          A.store.Float64(ptr + 0 + 24, x["usage"]["total"] === undefined ? 0 : (x["usage"]["total"] as number));
        }
      }
    },
    "load_ProcessorInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 36)) {
        x["usage"] = {};
        if (A.load.Bool(ptr + 0 + 32)) {
          x["usage"]["user"] = A.load.Float64(ptr + 0 + 0);
        } else {
          delete x["usage"]["user"];
        }
        if (A.load.Bool(ptr + 0 + 33)) {
          x["usage"]["kernel"] = A.load.Float64(ptr + 0 + 8);
        } else {
          delete x["usage"]["kernel"];
        }
        if (A.load.Bool(ptr + 0 + 34)) {
          x["usage"]["idle"] = A.load.Float64(ptr + 0 + 16);
        } else {
          delete x["usage"]["idle"];
        }
        if (A.load.Bool(ptr + 0 + 35)) {
          x["usage"]["total"] = A.load.Float64(ptr + 0 + 24);
        } else {
          delete x["usage"]["total"];
        }
      } else {
        delete x["usage"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CpuInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 24, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
      } else {
        A.store.Bool(ptr + 25, true);
        A.store.Bool(ptr + 24, "numOfProcessors" in x ? true : false);
        A.store.Int32(ptr + 0, x["numOfProcessors"] === undefined ? 0 : (x["numOfProcessors"] as number));
        A.store.Ref(ptr + 4, x["archName"]);
        A.store.Ref(ptr + 8, x["modelName"]);
        A.store.Ref(ptr + 12, x["features"]);
        A.store.Ref(ptr + 16, x["processors"]);
        A.store.Ref(ptr + 20, x["temperatures"]);
      }
    },
    "load_CpuInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["numOfProcessors"] = A.load.Int32(ptr + 0);
      } else {
        delete x["numOfProcessors"];
      }
      x["archName"] = A.load.Ref(ptr + 4, undefined);
      x["modelName"] = A.load.Ref(ptr + 8, undefined);
      x["features"] = A.load.Ref(ptr + 12, undefined);
      x["processors"] = A.load.Ref(ptr + 16, undefined);
      x["temperatures"] = A.load.Ref(ptr + 20, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.cpu && "getInfo" in WEBEXT?.system?.cpu) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.cpu.getInfo);
    },
    "call_GetInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.system.cpu.getInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.cpu.getInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
