import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/diagnostics", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_SendPacketResult": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["ip"]);
        A.store.Bool(ptr + 16, "latency" in x ? true : false);
        A.store.Float64(ptr + 8, x["latency"] === undefined ? 0 : (x["latency"] as number));
      }
    },
    "load_SendPacketResult": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["ip"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["latency"] = A.load.Float64(ptr + 8);
      } else {
        delete x["latency"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SendPacketOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 19, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 18, false);
        A.store.Int32(ptr + 12, 0);
      } else {
        A.store.Bool(ptr + 19, true);
        A.store.Ref(ptr + 0, x["ip"]);
        A.store.Bool(ptr + 16, "ttl" in x ? true : false);
        A.store.Int32(ptr + 4, x["ttl"] === undefined ? 0 : (x["ttl"] as number));
        A.store.Bool(ptr + 17, "timeout" in x ? true : false);
        A.store.Int32(ptr + 8, x["timeout"] === undefined ? 0 : (x["timeout"] as number));
        A.store.Bool(ptr + 18, "size" in x ? true : false);
        A.store.Int32(ptr + 12, x["size"] === undefined ? 0 : (x["size"] as number));
      }
    },
    "load_SendPacketOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["ip"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["ttl"] = A.load.Int32(ptr + 4);
      } else {
        delete x["ttl"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["timeout"] = A.load.Int32(ptr + 8);
      } else {
        delete x["timeout"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["size"] = A.load.Int32(ptr + 12);
      } else {
        delete x["size"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_SendPacket": (): heap.Ref<boolean> => {
      if (WEBEXT?.diagnostics && "sendPacket" in WEBEXT?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendPacket": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.diagnostics.sendPacket);
    },
    "call_SendPacket": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["ip"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 16)) {
        options_ffi["ttl"] = A.load.Int32(options + 4);
      }
      if (A.load.Bool(options + 17)) {
        options_ffi["timeout"] = A.load.Int32(options + 8);
      }
      if (A.load.Bool(options + 18)) {
        options_ffi["size"] = A.load.Int32(options + 12);
      }

      const _ret = WEBEXT.diagnostics.sendPacket(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SendPacket": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["ip"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 16)) {
          options_ffi["ttl"] = A.load.Int32(options + 4);
        }
        if (A.load.Bool(options + 17)) {
          options_ffi["timeout"] = A.load.Int32(options + 8);
        }
        if (A.load.Bool(options + 18)) {
          options_ffi["size"] = A.load.Int32(options + 12);
        }

        const _ret = WEBEXT.diagnostics.sendPacket(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
