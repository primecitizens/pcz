import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/guestviewinternal", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Size": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Int64(ptr + 0, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Int64(ptr + 8, x["width"] === undefined ? 0 : (x["width"] as number));
      }
    },
    "load_Size": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["height"] = A.load.Int64(ptr + 0);
      x["width"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SizeParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 74, false);
        A.store.Bool(ptr + 73, false);
        A.store.Bool(ptr + 0, false);

        A.store.Bool(ptr + 8 + 16, false);
        A.store.Int64(ptr + 8 + 0, 0);
        A.store.Int64(ptr + 8 + 8, 0);

        A.store.Bool(ptr + 32 + 16, false);
        A.store.Int64(ptr + 32 + 0, 0);
        A.store.Int64(ptr + 32 + 8, 0);

        A.store.Bool(ptr + 56 + 16, false);
        A.store.Int64(ptr + 56 + 0, 0);
        A.store.Int64(ptr + 56 + 8, 0);
      } else {
        A.store.Bool(ptr + 74, true);
        A.store.Bool(ptr + 73, "enableAutoSize" in x ? true : false);
        A.store.Bool(ptr + 0, x["enableAutoSize"] ? true : false);

        if (typeof x["max"] === "undefined") {
          A.store.Bool(ptr + 8 + 16, false);
          A.store.Int64(ptr + 8 + 0, 0);
          A.store.Int64(ptr + 8 + 8, 0);
        } else {
          A.store.Bool(ptr + 8 + 16, true);
          A.store.Int64(ptr + 8 + 0, x["max"]["height"] === undefined ? 0 : (x["max"]["height"] as number));
          A.store.Int64(ptr + 8 + 8, x["max"]["width"] === undefined ? 0 : (x["max"]["width"] as number));
        }

        if (typeof x["min"] === "undefined") {
          A.store.Bool(ptr + 32 + 16, false);
          A.store.Int64(ptr + 32 + 0, 0);
          A.store.Int64(ptr + 32 + 8, 0);
        } else {
          A.store.Bool(ptr + 32 + 16, true);
          A.store.Int64(ptr + 32 + 0, x["min"]["height"] === undefined ? 0 : (x["min"]["height"] as number));
          A.store.Int64(ptr + 32 + 8, x["min"]["width"] === undefined ? 0 : (x["min"]["width"] as number));
        }

        if (typeof x["normal"] === "undefined") {
          A.store.Bool(ptr + 56 + 16, false);
          A.store.Int64(ptr + 56 + 0, 0);
          A.store.Int64(ptr + 56 + 8, 0);
        } else {
          A.store.Bool(ptr + 56 + 16, true);
          A.store.Int64(ptr + 56 + 0, x["normal"]["height"] === undefined ? 0 : (x["normal"]["height"] as number));
          A.store.Int64(ptr + 56 + 8, x["normal"]["width"] === undefined ? 0 : (x["normal"]["width"] as number));
        }
      }
    },
    "load_SizeParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 73)) {
        x["enableAutoSize"] = A.load.Bool(ptr + 0);
      } else {
        delete x["enableAutoSize"];
      }
      if (A.load.Bool(ptr + 8 + 16)) {
        x["max"] = {};
        x["max"]["height"] = A.load.Int64(ptr + 8 + 0);
        x["max"]["width"] = A.load.Int64(ptr + 8 + 8);
      } else {
        delete x["max"];
      }
      if (A.load.Bool(ptr + 32 + 16)) {
        x["min"] = {};
        x["min"]["height"] = A.load.Int64(ptr + 32 + 0);
        x["min"]["width"] = A.load.Int64(ptr + 32 + 8);
      } else {
        delete x["min"];
      }
      if (A.load.Bool(ptr + 56 + 16)) {
        x["normal"] = {};
        x["normal"]["height"] = A.load.Int64(ptr + 56 + 0);
        x["normal"]["width"] = A.load.Int64(ptr + 56 + 8);
      } else {
        delete x["normal"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_CreateGuest": (): heap.Ref<boolean> => {
      if (WEBEXT?.guestViewInternal && "createGuest" in WEBEXT?.guestViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CreateGuest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.guestViewInternal.createGuest);
    },
    "call_CreateGuest": (
      retPtr: Pointer,
      viewType: heap.Ref<object>,
      ownerRoutingId: number,
      createParams: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.guestViewInternal.createGuest(
        A.H.get<object>(viewType),
        ownerRoutingId,
        A.H.get<object>(createParams)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_CreateGuest": (
      retPtr: Pointer,
      errPtr: Pointer,
      viewType: heap.Ref<object>,
      ownerRoutingId: number,
      createParams: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.guestViewInternal.createGuest(
          A.H.get<object>(viewType),
          ownerRoutingId,
          A.H.get<object>(createParams)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DestroyUnattachedGuest": (): heap.Ref<boolean> => {
      if (WEBEXT?.guestViewInternal && "destroyUnattachedGuest" in WEBEXT?.guestViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DestroyUnattachedGuest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.guestViewInternal.destroyUnattachedGuest);
    },
    "call_DestroyUnattachedGuest": (retPtr: Pointer, instanceId: number): void => {
      const _ret = WEBEXT.guestViewInternal.destroyUnattachedGuest(instanceId);
    },
    "try_DestroyUnattachedGuest": (retPtr: Pointer, errPtr: Pointer, instanceId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.guestViewInternal.destroyUnattachedGuest(instanceId);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetSize": (): heap.Ref<boolean> => {
      if (WEBEXT?.guestViewInternal && "setSize" in WEBEXT?.guestViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetSize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.guestViewInternal.setSize);
    },
    "call_SetSize": (retPtr: Pointer, instanceId: number, params: Pointer): void => {
      const params_ffi = {};

      if (A.load.Bool(params + 73)) {
        params_ffi["enableAutoSize"] = A.load.Bool(params + 0);
      }
      if (A.load.Bool(params + 8 + 16)) {
        params_ffi["max"] = {};
        params_ffi["max"]["height"] = A.load.Int64(params + 8 + 0);
        params_ffi["max"]["width"] = A.load.Int64(params + 8 + 8);
      }
      if (A.load.Bool(params + 32 + 16)) {
        params_ffi["min"] = {};
        params_ffi["min"]["height"] = A.load.Int64(params + 32 + 0);
        params_ffi["min"]["width"] = A.load.Int64(params + 32 + 8);
      }
      if (A.load.Bool(params + 56 + 16)) {
        params_ffi["normal"] = {};
        params_ffi["normal"]["height"] = A.load.Int64(params + 56 + 0);
        params_ffi["normal"]["width"] = A.load.Int64(params + 56 + 8);
      }

      const _ret = WEBEXT.guestViewInternal.setSize(instanceId, params_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetSize": (retPtr: Pointer, errPtr: Pointer, instanceId: number, params: Pointer): heap.Ref<boolean> => {
      try {
        const params_ffi = {};

        if (A.load.Bool(params + 73)) {
          params_ffi["enableAutoSize"] = A.load.Bool(params + 0);
        }
        if (A.load.Bool(params + 8 + 16)) {
          params_ffi["max"] = {};
          params_ffi["max"]["height"] = A.load.Int64(params + 8 + 0);
          params_ffi["max"]["width"] = A.load.Int64(params + 8 + 8);
        }
        if (A.load.Bool(params + 32 + 16)) {
          params_ffi["min"] = {};
          params_ffi["min"]["height"] = A.load.Int64(params + 32 + 0);
          params_ffi["min"]["width"] = A.load.Int64(params + 32 + 8);
        }
        if (A.load.Bool(params + 56 + 16)) {
          params_ffi["normal"] = {};
          params_ffi["normal"]["height"] = A.load.Int64(params + 56 + 0);
          params_ffi["normal"]["width"] = A.load.Int64(params + 56 + 8);
        }

        const _ret = WEBEXT.guestViewInternal.setSize(instanceId, params_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
