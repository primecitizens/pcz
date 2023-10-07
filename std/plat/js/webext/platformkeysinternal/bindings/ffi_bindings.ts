import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/platformkeysinternal", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_GetPublicKey": (): heap.Ref<boolean> => {
      if (WEBEXT?.platformKeysInternal && "getPublicKey" in WEBEXT?.platformKeysInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPublicKey": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.platformKeysInternal.getPublicKey);
    },
    "call_GetPublicKey": (
      retPtr: Pointer,
      certificate: heap.Ref<object>,
      algorithmName: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.platformKeysInternal.getPublicKey(
        A.H.get<object>(certificate),
        A.H.get<object>(algorithmName),
        A.H.get<object>(callback)
      );
    },
    "try_GetPublicKey": (
      retPtr: Pointer,
      errPtr: Pointer,
      certificate: heap.Ref<object>,
      algorithmName: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.platformKeysInternal.getPublicKey(
          A.H.get<object>(certificate),
          A.H.get<object>(algorithmName),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPublicKeyBySpki": (): heap.Ref<boolean> => {
      if (WEBEXT?.platformKeysInternal && "getPublicKeyBySpki" in WEBEXT?.platformKeysInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPublicKeyBySpki": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.platformKeysInternal.getPublicKeyBySpki);
    },
    "call_GetPublicKeyBySpki": (
      retPtr: Pointer,
      publicKeySpkiDer: heap.Ref<object>,
      algorithmName: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.platformKeysInternal.getPublicKeyBySpki(
        A.H.get<object>(publicKeySpkiDer),
        A.H.get<object>(algorithmName),
        A.H.get<object>(callback)
      );
    },
    "try_GetPublicKeyBySpki": (
      retPtr: Pointer,
      errPtr: Pointer,
      publicKeySpkiDer: heap.Ref<object>,
      algorithmName: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.platformKeysInternal.getPublicKeyBySpki(
          A.H.get<object>(publicKeySpkiDer),
          A.H.get<object>(algorithmName),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SelectClientCertificates": (): heap.Ref<boolean> => {
      if (WEBEXT?.platformKeysInternal && "selectClientCertificates" in WEBEXT?.platformKeysInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SelectClientCertificates": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.platformKeysInternal.selectClientCertificates);
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

      const _ret = WEBEXT.platformKeysInternal.selectClientCertificates(details_ffi, A.H.get<object>(callback));
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

        const _ret = WEBEXT.platformKeysInternal.selectClientCertificates(details_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Sign": (): heap.Ref<boolean> => {
      if (WEBEXT?.platformKeysInternal && "sign" in WEBEXT?.platformKeysInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Sign": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.platformKeysInternal.sign);
    },
    "call_Sign": (
      retPtr: Pointer,
      tokenId: heap.Ref<object>,
      publicKey: heap.Ref<object>,
      algorithmName: heap.Ref<object>,
      hashAlgorithmName: heap.Ref<object>,
      data: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.platformKeysInternal.sign(
        A.H.get<object>(tokenId),
        A.H.get<object>(publicKey),
        A.H.get<object>(algorithmName),
        A.H.get<object>(hashAlgorithmName),
        A.H.get<object>(data),
        A.H.get<object>(callback)
      );
    },
    "try_Sign": (
      retPtr: Pointer,
      errPtr: Pointer,
      tokenId: heap.Ref<object>,
      publicKey: heap.Ref<object>,
      algorithmName: heap.Ref<object>,
      hashAlgorithmName: heap.Ref<object>,
      data: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.platformKeysInternal.sign(
          A.H.get<object>(tokenId),
          A.H.get<object>(publicKey),
          A.H.get<object>(algorithmName),
          A.H.get<object>(hashAlgorithmName),
          A.H.get<object>(data),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
