import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/enterprise/platformkeysinternal", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Hash": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["name"]);
      }
    },
    "load_Hash": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Algorithm": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 24, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);

        A.store.Bool(ptr + 12 + 4, false);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 20, undefined);
      } else {
        A.store.Bool(ptr + 25, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Bool(ptr + 24, "modulusLength" in x ? true : false);
        A.store.Int32(ptr + 4, x["modulusLength"] === undefined ? 0 : (x["modulusLength"] as number));
        A.store.Ref(ptr + 8, x["publicExponent"]);

        if (typeof x["hash"] === "undefined") {
          A.store.Bool(ptr + 12 + 4, false);
          A.store.Ref(ptr + 12 + 0, undefined);
        } else {
          A.store.Bool(ptr + 12 + 4, true);
          A.store.Ref(ptr + 12 + 0, x["hash"]["name"]);
        }
        A.store.Ref(ptr + 20, x["namedCurve"]);
      }
    },
    "load_Algorithm": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 24)) {
        x["modulusLength"] = A.load.Int32(ptr + 4);
      } else {
        delete x["modulusLength"];
      }
      x["publicExponent"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 12 + 4)) {
        x["hash"] = {};
        x["hash"]["name"] = A.load.Ref(ptr + 12 + 0, undefined);
      } else {
        delete x["hash"];
      }
      x["namedCurve"] = A.load.Ref(ptr + 20, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GenerateKey": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.platformKeysInternal && "generateKey" in WEBEXT?.enterprise?.platformKeysInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GenerateKey": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.platformKeysInternal.generateKey);
    },
    "call_GenerateKey": (
      retPtr: Pointer,
      tokenId: heap.Ref<object>,
      algorithm: Pointer,
      softwareBacked: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): void => {
      const algorithm_ffi = {};

      algorithm_ffi["name"] = A.load.Ref(algorithm + 0, undefined);
      if (A.load.Bool(algorithm + 24)) {
        algorithm_ffi["modulusLength"] = A.load.Int32(algorithm + 4);
      }
      algorithm_ffi["publicExponent"] = A.load.Ref(algorithm + 8, undefined);
      if (A.load.Bool(algorithm + 12 + 4)) {
        algorithm_ffi["hash"] = {};
        algorithm_ffi["hash"]["name"] = A.load.Ref(algorithm + 12 + 0, undefined);
      }
      algorithm_ffi["namedCurve"] = A.load.Ref(algorithm + 20, undefined);

      const _ret = WEBEXT.enterprise.platformKeysInternal.generateKey(
        A.H.get<object>(tokenId),
        algorithm_ffi,
        softwareBacked === A.H.TRUE,
        A.H.get<object>(callback)
      );
    },
    "try_GenerateKey": (
      retPtr: Pointer,
      errPtr: Pointer,
      tokenId: heap.Ref<object>,
      algorithm: Pointer,
      softwareBacked: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const algorithm_ffi = {};

        algorithm_ffi["name"] = A.load.Ref(algorithm + 0, undefined);
        if (A.load.Bool(algorithm + 24)) {
          algorithm_ffi["modulusLength"] = A.load.Int32(algorithm + 4);
        }
        algorithm_ffi["publicExponent"] = A.load.Ref(algorithm + 8, undefined);
        if (A.load.Bool(algorithm + 12 + 4)) {
          algorithm_ffi["hash"] = {};
          algorithm_ffi["hash"]["name"] = A.load.Ref(algorithm + 12 + 0, undefined);
        }
        algorithm_ffi["namedCurve"] = A.load.Ref(algorithm + 20, undefined);

        const _ret = WEBEXT.enterprise.platformKeysInternal.generateKey(
          A.H.get<object>(tokenId),
          algorithm_ffi,
          softwareBacked === A.H.TRUE,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetTokens": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.platformKeysInternal && "getTokens" in WEBEXT?.enterprise?.platformKeysInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetTokens": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.platformKeysInternal.getTokens);
    },
    "call_GetTokens": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.enterprise.platformKeysInternal.getTokens(A.H.get<object>(callback));
    },
    "try_GetTokens": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.platformKeysInternal.getTokens(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
