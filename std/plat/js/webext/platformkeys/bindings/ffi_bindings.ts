import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/platformkeys", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_ClientCertificateType": (ref: heap.Ref<string>): number => {
      const idx = ["rsaSign", "ecdsaSign"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ClientCertificateRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["certificateTypes"]);
        A.store.Ref(ptr + 4, x["certificateAuthorities"]);
      }
    },
    "load_ClientCertificateRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["certificateTypes"] = A.load.Ref(ptr + 0, undefined);
      x["certificateAuthorities"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Match": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["certificate"]);
        A.store.Ref(ptr + 4, x["keyAlgorithm"]);
      }
    },
    "load_Match": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["certificate"] = A.load.Ref(ptr + 0, undefined);
      x["keyAlgorithm"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SelectDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);

        A.store.Bool(ptr + 0 + 8, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 16, false);
      } else {
        A.store.Bool(ptr + 18, true);

        if (typeof x["request"] === "undefined") {
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
        } else {
          A.store.Bool(ptr + 0 + 8, true);
          A.store.Ref(ptr + 0 + 0, x["request"]["certificateTypes"]);
          A.store.Ref(ptr + 0 + 4, x["request"]["certificateAuthorities"]);
        }
        A.store.Ref(ptr + 12, x["clientCerts"]);
        A.store.Bool(ptr + 17, "interactive" in x ? true : false);
        A.store.Bool(ptr + 16, x["interactive"] ? true : false);
      }
    },
    "load_SelectDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 8)) {
        x["request"] = {};
        x["request"]["certificateTypes"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["request"]["certificateAuthorities"] = A.load.Ref(ptr + 0 + 4, undefined);
      } else {
        delete x["request"];
      }
      x["clientCerts"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 17)) {
        x["interactive"] = A.load.Bool(ptr + 16);
      } else {
        delete x["interactive"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_VerificationResult": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "trusted" in x ? true : false);
        A.store.Bool(ptr + 0, x["trusted"] ? true : false);
        A.store.Ref(ptr + 4, x["debug_errors"]);
      }
    },
    "load_VerificationResult": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["trusted"] = A.load.Bool(ptr + 0);
      } else {
        delete x["trusted"];
      }
      x["debug_errors"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_VerificationDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["serverCertificateChain"]);
        A.store.Ref(ptr + 4, x["hostname"]);
      }
    },
    "load_VerificationDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["serverCertificateChain"] = A.load.Ref(ptr + 0, undefined);
      x["hostname"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetKeyPair": (): heap.Ref<boolean> => {
      if (WEBEXT?.platformKeys && "getKeyPair" in WEBEXT?.platformKeys) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetKeyPair": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.platformKeys.getKeyPair);
    },
    "call_GetKeyPair": (
      retPtr: Pointer,
      certificate: heap.Ref<object>,
      parameters: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.platformKeys.getKeyPair(
        A.H.get<object>(certificate),
        A.H.get<object>(parameters),
        A.H.get<object>(callback)
      );
    },
    "try_GetKeyPair": (
      retPtr: Pointer,
      errPtr: Pointer,
      certificate: heap.Ref<object>,
      parameters: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.platformKeys.getKeyPair(
          A.H.get<object>(certificate),
          A.H.get<object>(parameters),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetKeyPairBySpki": (): heap.Ref<boolean> => {
      if (WEBEXT?.platformKeys && "getKeyPairBySpki" in WEBEXT?.platformKeys) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetKeyPairBySpki": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.platformKeys.getKeyPairBySpki);
    },
    "call_GetKeyPairBySpki": (
      retPtr: Pointer,
      publicKeySpkiDer: heap.Ref<object>,
      parameters: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.platformKeys.getKeyPairBySpki(
        A.H.get<object>(publicKeySpkiDer),
        A.H.get<object>(parameters),
        A.H.get<object>(callback)
      );
    },
    "try_GetKeyPairBySpki": (
      retPtr: Pointer,
      errPtr: Pointer,
      publicKeySpkiDer: heap.Ref<object>,
      parameters: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.platformKeys.getKeyPairBySpki(
          A.H.get<object>(publicKeySpkiDer),
          A.H.get<object>(parameters),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SelectClientCertificates": (): heap.Ref<boolean> => {
      if (WEBEXT?.platformKeys && "selectClientCertificates" in WEBEXT?.platformKeys) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SelectClientCertificates": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.platformKeys.selectClientCertificates);
    },
    "call_SelectClientCertificates": (retPtr: Pointer, details: Pointer, callback: heap.Ref<object>): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 0 + 8)) {
        details_ffi["request"] = {};
        details_ffi["request"]["certificateTypes"] = A.load.Ref(details + 0 + 0, undefined);
        details_ffi["request"]["certificateAuthorities"] = A.load.Ref(details + 0 + 4, undefined);
      }
      details_ffi["clientCerts"] = A.load.Ref(details + 12, undefined);
      if (A.load.Bool(details + 17)) {
        details_ffi["interactive"] = A.load.Bool(details + 16);
      }

      const _ret = WEBEXT.platformKeys.selectClientCertificates(details_ffi, A.H.get<object>(callback));
    },
    "try_SelectClientCertificates": (
      retPtr: Pointer,
      errPtr: Pointer,
      details: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 0 + 8)) {
          details_ffi["request"] = {};
          details_ffi["request"]["certificateTypes"] = A.load.Ref(details + 0 + 0, undefined);
          details_ffi["request"]["certificateAuthorities"] = A.load.Ref(details + 0 + 4, undefined);
        }
        details_ffi["clientCerts"] = A.load.Ref(details + 12, undefined);
        if (A.load.Bool(details + 17)) {
          details_ffi["interactive"] = A.load.Bool(details + 16);
        }

        const _ret = WEBEXT.platformKeys.selectClientCertificates(details_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SubtleCrypto": (): heap.Ref<boolean> => {
      if (WEBEXT?.platformKeys && "subtleCrypto" in WEBEXT?.platformKeys) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SubtleCrypto": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.platformKeys.subtleCrypto);
    },
    "call_SubtleCrypto": (retPtr: Pointer): void => {
      const _ret = WEBEXT.platformKeys.subtleCrypto();
      A.store.Ref(retPtr, _ret);
    },
    "try_SubtleCrypto": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.platformKeys.subtleCrypto();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_VerifyTLSServerCertificate": (): heap.Ref<boolean> => {
      if (WEBEXT?.platformKeys && "verifyTLSServerCertificate" in WEBEXT?.platformKeys) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_VerifyTLSServerCertificate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.platformKeys.verifyTLSServerCertificate);
    },
    "call_VerifyTLSServerCertificate": (retPtr: Pointer, details: Pointer, callback: heap.Ref<object>): void => {
      const details_ffi = {};

      details_ffi["serverCertificateChain"] = A.load.Ref(details + 0, undefined);
      details_ffi["hostname"] = A.load.Ref(details + 4, undefined);

      const _ret = WEBEXT.platformKeys.verifyTLSServerCertificate(details_ffi, A.H.get<object>(callback));
    },
    "try_VerifyTLSServerCertificate": (
      retPtr: Pointer,
      errPtr: Pointer,
      details: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["serverCertificateChain"] = A.load.Ref(details + 0, undefined);
        details_ffi["hostname"] = A.load.Ref(details + 4, undefined);

        const _ret = WEBEXT.platformKeys.verifyTLSServerCertificate(details_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
