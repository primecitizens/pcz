import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/certificateprovider", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_Algorithm": (ref: heap.Ref<string>): number => {
      const idx = [
        "RSASSA_PKCS1_v1_5_MD5_SHA1",
        "RSASSA_PKCS1_v1_5_SHA1",
        "RSASSA_PKCS1_v1_5_SHA256",
        "RSASSA_PKCS1_v1_5_SHA384",
        "RSASSA_PKCS1_v1_5_SHA512",
        "RSASSA_PSS_SHA256",
        "RSASSA_PSS_SHA384",
        "RSASSA_PSS_SHA512",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_Hash": (ref: heap.Ref<string>): number => {
      const idx = ["MD5_SHA1", "SHA1", "SHA256", "SHA384", "SHA512"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_CertificateInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["certificate"]);
        A.store.Ref(ptr + 4, x["supportedHashes"]);
      }
    },
    "load_CertificateInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["certificate"] = A.load.Ref(ptr + 0, undefined);
      x["supportedHashes"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CertificatesUpdateRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "certificatesRequestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["certificatesRequestId"] === undefined ? 0 : (x["certificatesRequestId"] as number));
      }
    },
    "load_CertificatesUpdateRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["certificatesRequestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["certificatesRequestId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ClientCertificateInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["certificateChain"]);
        A.store.Ref(ptr + 4, x["supportedAlgorithms"]);
      }
    },
    "load_ClientCertificateInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["certificateChain"] = A.load.Ref(ptr + 0, undefined);
      x["supportedAlgorithms"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_Error": (ref: heap.Ref<string>): number => {
      const idx = ["GENERAL_ERROR"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PinRequestErrorType": (ref: heap.Ref<string>): number => {
      const idx = ["INVALID_PIN", "INVALID_PUK", "MAX_ATTEMPTS_EXCEEDED", "UNKNOWN_ERROR"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PinRequestType": (ref: heap.Ref<string>): number => {
      const idx = ["PIN", "PUK"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PinResponseDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["userInput"]);
      }
    },
    "load_PinResponseDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["userInput"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ReportSignatureDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "signRequestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["signRequestId"] === undefined ? 0 : (x["signRequestId"] as number));
        A.store.Enum(ptr + 4, ["GENERAL_ERROR"].indexOf(x["error"] as string));
        A.store.Ref(ptr + 8, x["signature"]);
      }
    },
    "load_ReportSignatureDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["signRequestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["signRequestId"];
      }
      x["error"] = A.load.Enum(ptr + 4, ["GENERAL_ERROR"]);
      x["signature"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RequestPinDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Enum(ptr + 4, -1);
        A.store.Enum(ptr + 8, -1);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 12, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Bool(ptr + 16, "signRequestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["signRequestId"] === undefined ? 0 : (x["signRequestId"] as number));
        A.store.Enum(ptr + 4, ["PIN", "PUK"].indexOf(x["requestType"] as string));
        A.store.Enum(
          ptr + 8,
          ["INVALID_PIN", "INVALID_PUK", "MAX_ATTEMPTS_EXCEEDED", "UNKNOWN_ERROR"].indexOf(x["errorType"] as string)
        );
        A.store.Bool(ptr + 17, "attemptsLeft" in x ? true : false);
        A.store.Int32(ptr + 12, x["attemptsLeft"] === undefined ? 0 : (x["attemptsLeft"] as number));
      }
    },
    "load_RequestPinDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["signRequestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["signRequestId"];
      }
      x["requestType"] = A.load.Enum(ptr + 4, ["PIN", "PUK"]);
      x["errorType"] = A.load.Enum(ptr + 8, ["INVALID_PIN", "INVALID_PUK", "MAX_ATTEMPTS_EXCEEDED", "UNKNOWN_ERROR"]);
      if (A.load.Bool(ptr + 17)) {
        x["attemptsLeft"] = A.load.Int32(ptr + 12);
      } else {
        delete x["attemptsLeft"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetCertificatesDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "certificatesRequestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["certificatesRequestId"] === undefined ? 0 : (x["certificatesRequestId"] as number));
        A.store.Enum(ptr + 4, ["GENERAL_ERROR"].indexOf(x["error"] as string));
        A.store.Ref(ptr + 8, x["clientCertificates"]);
      }
    },
    "load_SetCertificatesDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["certificatesRequestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["certificatesRequestId"];
      }
      x["error"] = A.load.Enum(ptr + 4, ["GENERAL_ERROR"]);
      x["clientCertificates"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SignRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Bool(ptr + 16, "signRequestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["signRequestId"] === undefined ? 0 : (x["signRequestId"] as number));
        A.store.Ref(ptr + 4, x["digest"]);
        A.store.Enum(ptr + 8, ["MD5_SHA1", "SHA1", "SHA256", "SHA384", "SHA512"].indexOf(x["hash"] as string));
        A.store.Ref(ptr + 12, x["certificate"]);
      }
    },
    "load_SignRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["signRequestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["signRequestId"];
      }
      x["digest"] = A.load.Ref(ptr + 4, undefined);
      x["hash"] = A.load.Enum(ptr + 8, ["MD5_SHA1", "SHA1", "SHA256", "SHA384", "SHA512"]);
      x["certificate"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SignatureRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Bool(ptr + 16, "signRequestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["signRequestId"] === undefined ? 0 : (x["signRequestId"] as number));
        A.store.Ref(ptr + 4, x["input"]);
        A.store.Enum(
          ptr + 8,
          [
            "RSASSA_PKCS1_v1_5_MD5_SHA1",
            "RSASSA_PKCS1_v1_5_SHA1",
            "RSASSA_PKCS1_v1_5_SHA256",
            "RSASSA_PKCS1_v1_5_SHA384",
            "RSASSA_PKCS1_v1_5_SHA512",
            "RSASSA_PSS_SHA256",
            "RSASSA_PSS_SHA384",
            "RSASSA_PSS_SHA512",
          ].indexOf(x["algorithm"] as string)
        );
        A.store.Ref(ptr + 12, x["certificate"]);
      }
    },
    "load_SignatureRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["signRequestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["signRequestId"];
      }
      x["input"] = A.load.Ref(ptr + 4, undefined);
      x["algorithm"] = A.load.Enum(ptr + 8, [
        "RSASSA_PKCS1_v1_5_MD5_SHA1",
        "RSASSA_PKCS1_v1_5_SHA1",
        "RSASSA_PKCS1_v1_5_SHA256",
        "RSASSA_PKCS1_v1_5_SHA384",
        "RSASSA_PKCS1_v1_5_SHA512",
        "RSASSA_PSS_SHA256",
        "RSASSA_PSS_SHA384",
        "RSASSA_PSS_SHA512",
      ]);
      x["certificate"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_StopPinRequestDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "signRequestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["signRequestId"] === undefined ? 0 : (x["signRequestId"] as number));
        A.store.Enum(
          ptr + 4,
          ["INVALID_PIN", "INVALID_PUK", "MAX_ATTEMPTS_EXCEEDED", "UNKNOWN_ERROR"].indexOf(x["errorType"] as string)
        );
      }
    },
    "load_StopPinRequestDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["signRequestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["signRequestId"];
      }
      x["errorType"] = A.load.Enum(ptr + 4, ["INVALID_PIN", "INVALID_PUK", "MAX_ATTEMPTS_EXCEEDED", "UNKNOWN_ERROR"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_OnCertificatesRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.certificateProvider?.onCertificatesRequested &&
        "addListener" in WEBEXT?.certificateProvider?.onCertificatesRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCertificatesRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.onCertificatesRequested.addListener);
    },
    "call_OnCertificatesRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.certificateProvider.onCertificatesRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnCertificatesRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.certificateProvider.onCertificatesRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCertificatesRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.certificateProvider?.onCertificatesRequested &&
        "removeListener" in WEBEXT?.certificateProvider?.onCertificatesRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCertificatesRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.onCertificatesRequested.removeListener);
    },
    "call_OffCertificatesRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.certificateProvider.onCertificatesRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffCertificatesRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.certificateProvider.onCertificatesRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCertificatesRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.certificateProvider?.onCertificatesRequested &&
        "hasListener" in WEBEXT?.certificateProvider?.onCertificatesRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCertificatesRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.onCertificatesRequested.hasListener);
    },
    "call_HasOnCertificatesRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.certificateProvider.onCertificatesRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCertificatesRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.certificateProvider.onCertificatesRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCertificatesUpdateRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.certificateProvider?.onCertificatesUpdateRequested &&
        "addListener" in WEBEXT?.certificateProvider?.onCertificatesUpdateRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCertificatesUpdateRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.onCertificatesUpdateRequested.addListener);
    },
    "call_OnCertificatesUpdateRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.certificateProvider.onCertificatesUpdateRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnCertificatesUpdateRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.certificateProvider.onCertificatesUpdateRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCertificatesUpdateRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.certificateProvider?.onCertificatesUpdateRequested &&
        "removeListener" in WEBEXT?.certificateProvider?.onCertificatesUpdateRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCertificatesUpdateRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.onCertificatesUpdateRequested.removeListener);
    },
    "call_OffCertificatesUpdateRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.certificateProvider.onCertificatesUpdateRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffCertificatesUpdateRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.certificateProvider.onCertificatesUpdateRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCertificatesUpdateRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.certificateProvider?.onCertificatesUpdateRequested &&
        "hasListener" in WEBEXT?.certificateProvider?.onCertificatesUpdateRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCertificatesUpdateRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.onCertificatesUpdateRequested.hasListener);
    },
    "call_HasOnCertificatesUpdateRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.certificateProvider.onCertificatesUpdateRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCertificatesUpdateRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.certificateProvider.onCertificatesUpdateRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSignDigestRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.certificateProvider?.onSignDigestRequested &&
        "addListener" in WEBEXT?.certificateProvider?.onSignDigestRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSignDigestRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.onSignDigestRequested.addListener);
    },
    "call_OnSignDigestRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.certificateProvider.onSignDigestRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnSignDigestRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.certificateProvider.onSignDigestRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSignDigestRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.certificateProvider?.onSignDigestRequested &&
        "removeListener" in WEBEXT?.certificateProvider?.onSignDigestRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSignDigestRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.onSignDigestRequested.removeListener);
    },
    "call_OffSignDigestRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.certificateProvider.onSignDigestRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffSignDigestRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.certificateProvider.onSignDigestRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSignDigestRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.certificateProvider?.onSignDigestRequested &&
        "hasListener" in WEBEXT?.certificateProvider?.onSignDigestRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSignDigestRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.onSignDigestRequested.hasListener);
    },
    "call_HasOnSignDigestRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.certificateProvider.onSignDigestRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSignDigestRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.certificateProvider.onSignDigestRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSignatureRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.certificateProvider?.onSignatureRequested &&
        "addListener" in WEBEXT?.certificateProvider?.onSignatureRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSignatureRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.onSignatureRequested.addListener);
    },
    "call_OnSignatureRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.certificateProvider.onSignatureRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnSignatureRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.certificateProvider.onSignatureRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSignatureRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.certificateProvider?.onSignatureRequested &&
        "removeListener" in WEBEXT?.certificateProvider?.onSignatureRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSignatureRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.onSignatureRequested.removeListener);
    },
    "call_OffSignatureRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.certificateProvider.onSignatureRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffSignatureRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.certificateProvider.onSignatureRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSignatureRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.certificateProvider?.onSignatureRequested &&
        "hasListener" in WEBEXT?.certificateProvider?.onSignatureRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSignatureRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.onSignatureRequested.hasListener);
    },
    "call_HasOnSignatureRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.certificateProvider.onSignatureRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSignatureRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.certificateProvider.onSignatureRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReportSignature": (): heap.Ref<boolean> => {
      if (WEBEXT?.certificateProvider && "reportSignature" in WEBEXT?.certificateProvider) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportSignature": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.reportSignature);
    },
    "call_ReportSignature": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 12)) {
        details_ffi["signRequestId"] = A.load.Int32(details + 0);
      }
      details_ffi["error"] = A.load.Enum(details + 4, ["GENERAL_ERROR"]);
      details_ffi["signature"] = A.load.Ref(details + 8, undefined);

      const _ret = WEBEXT.certificateProvider.reportSignature(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ReportSignature": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 12)) {
          details_ffi["signRequestId"] = A.load.Int32(details + 0);
        }
        details_ffi["error"] = A.load.Enum(details + 4, ["GENERAL_ERROR"]);
        details_ffi["signature"] = A.load.Ref(details + 8, undefined);

        const _ret = WEBEXT.certificateProvider.reportSignature(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RequestPin": (): heap.Ref<boolean> => {
      if (WEBEXT?.certificateProvider && "requestPin" in WEBEXT?.certificateProvider) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RequestPin": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.requestPin);
    },
    "call_RequestPin": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 16)) {
        details_ffi["signRequestId"] = A.load.Int32(details + 0);
      }
      details_ffi["requestType"] = A.load.Enum(details + 4, ["PIN", "PUK"]);
      details_ffi["errorType"] = A.load.Enum(details + 8, [
        "INVALID_PIN",
        "INVALID_PUK",
        "MAX_ATTEMPTS_EXCEEDED",
        "UNKNOWN_ERROR",
      ]);
      if (A.load.Bool(details + 17)) {
        details_ffi["attemptsLeft"] = A.load.Int32(details + 12);
      }

      const _ret = WEBEXT.certificateProvider.requestPin(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RequestPin": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 16)) {
          details_ffi["signRequestId"] = A.load.Int32(details + 0);
        }
        details_ffi["requestType"] = A.load.Enum(details + 4, ["PIN", "PUK"]);
        details_ffi["errorType"] = A.load.Enum(details + 8, [
          "INVALID_PIN",
          "INVALID_PUK",
          "MAX_ATTEMPTS_EXCEEDED",
          "UNKNOWN_ERROR",
        ]);
        if (A.load.Bool(details + 17)) {
          details_ffi["attemptsLeft"] = A.load.Int32(details + 12);
        }

        const _ret = WEBEXT.certificateProvider.requestPin(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetCertificates": (): heap.Ref<boolean> => {
      if (WEBEXT?.certificateProvider && "setCertificates" in WEBEXT?.certificateProvider) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetCertificates": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.setCertificates);
    },
    "call_SetCertificates": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 12)) {
        details_ffi["certificatesRequestId"] = A.load.Int32(details + 0);
      }
      details_ffi["error"] = A.load.Enum(details + 4, ["GENERAL_ERROR"]);
      details_ffi["clientCertificates"] = A.load.Ref(details + 8, undefined);

      const _ret = WEBEXT.certificateProvider.setCertificates(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetCertificates": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 12)) {
          details_ffi["certificatesRequestId"] = A.load.Int32(details + 0);
        }
        details_ffi["error"] = A.load.Enum(details + 4, ["GENERAL_ERROR"]);
        details_ffi["clientCertificates"] = A.load.Ref(details + 8, undefined);

        const _ret = WEBEXT.certificateProvider.setCertificates(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StopPinRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.certificateProvider && "stopPinRequest" in WEBEXT?.certificateProvider) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StopPinRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProvider.stopPinRequest);
    },
    "call_StopPinRequest": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 8)) {
        details_ffi["signRequestId"] = A.load.Int32(details + 0);
      }
      details_ffi["errorType"] = A.load.Enum(details + 4, [
        "INVALID_PIN",
        "INVALID_PUK",
        "MAX_ATTEMPTS_EXCEEDED",
        "UNKNOWN_ERROR",
      ]);

      const _ret = WEBEXT.certificateProvider.stopPinRequest(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_StopPinRequest": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 8)) {
          details_ffi["signRequestId"] = A.load.Int32(details + 0);
        }
        details_ffi["errorType"] = A.load.Enum(details + 4, [
          "INVALID_PIN",
          "INVALID_PUK",
          "MAX_ATTEMPTS_EXCEEDED",
          "UNKNOWN_ERROR",
        ]);

        const _ret = WEBEXT.certificateProvider.stopPinRequest(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
