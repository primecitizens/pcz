import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/enterprise/platformkeys", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_Algorithm": (ref: heap.Ref<string>): number => {
      const idx = ["RSA", "ECDSA"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RegisterKeyOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["RSA", "ECDSA"].indexOf(x["algorithm"] as string));
      }
    },
    "load_RegisterKeyOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["algorithm"] = A.load.Enum(ptr + 0, ["RSA", "ECDSA"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_Scope": (ref: heap.Ref<string>): number => {
      const idx = ["USER", "MACHINE"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ChallengeKeyOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 4 + 4, false);
        A.store.Enum(ptr + 4 + 0, -1);
        A.store.Enum(ptr + 12, -1);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["challenge"]);

        if (typeof x["registerKey"] === "undefined") {
          A.store.Bool(ptr + 4 + 4, false);
          A.store.Enum(ptr + 4 + 0, -1);
        } else {
          A.store.Bool(ptr + 4 + 4, true);
          A.store.Enum(ptr + 4 + 0, ["RSA", "ECDSA"].indexOf(x["registerKey"]["algorithm"] as string));
        }
        A.store.Enum(ptr + 12, ["USER", "MACHINE"].indexOf(x["scope"] as string));
      }
    },
    "load_ChallengeKeyOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["challenge"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 4 + 4)) {
        x["registerKey"] = {};
        x["registerKey"]["algorithm"] = A.load.Enum(ptr + 4 + 0, ["RSA", "ECDSA"]);
      } else {
        delete x["registerKey"];
      }
      x["scope"] = A.load.Enum(ptr + 12, ["USER", "MACHINE"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Token": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["subtleCrypto"]);
        A.store.Ref(ptr + 8, x["softwareBackedSubtleCrypto"]);
      }
    },
    "load_Token": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["subtleCrypto"] = A.load.Ref(ptr + 4, undefined);
      x["softwareBackedSubtleCrypto"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_ChallengeKey": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.platformKeys && "challengeKey" in WEBEXT?.enterprise?.platformKeys) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ChallengeKey": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.platformKeys.challengeKey);
    },
    "call_ChallengeKey": (retPtr: Pointer, options: Pointer, callback: heap.Ref<object>): void => {
      const options_ffi = {};

      options_ffi["challenge"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 4)) {
        options_ffi["registerKey"] = {};
        options_ffi["registerKey"]["algorithm"] = A.load.Enum(options + 4 + 0, ["RSA", "ECDSA"]);
      }
      options_ffi["scope"] = A.load.Enum(options + 12, ["USER", "MACHINE"]);

      const _ret = WEBEXT.enterprise.platformKeys.challengeKey(options_ffi, A.H.get<object>(callback));
    },
    "try_ChallengeKey": (
      retPtr: Pointer,
      errPtr: Pointer,
      options: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["challenge"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["registerKey"] = {};
          options_ffi["registerKey"]["algorithm"] = A.load.Enum(options + 4 + 0, ["RSA", "ECDSA"]);
        }
        options_ffi["scope"] = A.load.Enum(options + 12, ["USER", "MACHINE"]);

        const _ret = WEBEXT.enterprise.platformKeys.challengeKey(options_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ChallengeMachineKey": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.platformKeys && "challengeMachineKey" in WEBEXT?.enterprise?.platformKeys) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ChallengeMachineKey": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.platformKeys.challengeMachineKey);
    },
    "call_ChallengeMachineKey": (
      retPtr: Pointer,
      challenge: heap.Ref<object>,
      registerKey: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.enterprise.platformKeys.challengeMachineKey(
        A.H.get<object>(challenge),
        registerKey === A.H.TRUE,
        A.H.get<object>(callback)
      );
    },
    "try_ChallengeMachineKey": (
      retPtr: Pointer,
      errPtr: Pointer,
      challenge: heap.Ref<object>,
      registerKey: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.platformKeys.challengeMachineKey(
          A.H.get<object>(challenge),
          registerKey === A.H.TRUE,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ChallengeUserKey": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.platformKeys && "challengeUserKey" in WEBEXT?.enterprise?.platformKeys) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ChallengeUserKey": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.platformKeys.challengeUserKey);
    },
    "call_ChallengeUserKey": (
      retPtr: Pointer,
      challenge: heap.Ref<object>,
      registerKey: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.enterprise.platformKeys.challengeUserKey(
        A.H.get<object>(challenge),
        registerKey === A.H.TRUE,
        A.H.get<object>(callback)
      );
    },
    "try_ChallengeUserKey": (
      retPtr: Pointer,
      errPtr: Pointer,
      challenge: heap.Ref<object>,
      registerKey: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.platformKeys.challengeUserKey(
          A.H.get<object>(challenge),
          registerKey === A.H.TRUE,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCertificates": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.platformKeys && "getCertificates" in WEBEXT?.enterprise?.platformKeys) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCertificates": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.platformKeys.getCertificates);
    },
    "call_GetCertificates": (retPtr: Pointer, tokenId: heap.Ref<object>, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.enterprise.platformKeys.getCertificates(A.H.get<object>(tokenId), A.H.get<object>(callback));
    },
    "try_GetCertificates": (
      retPtr: Pointer,
      errPtr: Pointer,
      tokenId: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.platformKeys.getCertificates(
          A.H.get<object>(tokenId),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetTokens": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.platformKeys && "getTokens" in WEBEXT?.enterprise?.platformKeys) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetTokens": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.platformKeys.getTokens);
    },
    "call_GetTokens": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.enterprise.platformKeys.getTokens(A.H.get<object>(callback));
    },
    "try_GetTokens": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.platformKeys.getTokens(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ImportCertificate": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.platformKeys && "importCertificate" in WEBEXT?.enterprise?.platformKeys) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ImportCertificate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.platformKeys.importCertificate);
    },
    "call_ImportCertificate": (
      retPtr: Pointer,
      tokenId: heap.Ref<object>,
      certificate: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.enterprise.platformKeys.importCertificate(
        A.H.get<object>(tokenId),
        A.H.get<object>(certificate),
        A.H.get<object>(callback)
      );
    },
    "try_ImportCertificate": (
      retPtr: Pointer,
      errPtr: Pointer,
      tokenId: heap.Ref<object>,
      certificate: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.platformKeys.importCertificate(
          A.H.get<object>(tokenId),
          A.H.get<object>(certificate),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveCertificate": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.platformKeys && "removeCertificate" in WEBEXT?.enterprise?.platformKeys) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveCertificate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.platformKeys.removeCertificate);
    },
    "call_RemoveCertificate": (
      retPtr: Pointer,
      tokenId: heap.Ref<object>,
      certificate: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.enterprise.platformKeys.removeCertificate(
        A.H.get<object>(tokenId),
        A.H.get<object>(certificate),
        A.H.get<object>(callback)
      );
    },
    "try_RemoveCertificate": (
      retPtr: Pointer,
      errPtr: Pointer,
      tokenId: heap.Ref<object>,
      certificate: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.platformKeys.removeCertificate(
          A.H.get<object>(tokenId),
          A.H.get<object>(certificate),
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
