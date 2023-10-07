import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/echoprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_GetUserConsentArgConsentRequester": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["origin"]);
        A.store.Ref(ptr + 4, x["serviceName"]);
        A.store.Bool(ptr + 16, "tabId" in x ? true : false);
        A.store.Int64(ptr + 8, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_GetUserConsentArgConsentRequester": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["origin"] = A.load.Ref(ptr + 0, undefined);
      x["serviceName"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["tabId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["tabId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetOfferInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.echoPrivate && "getOfferInfo" in WEBEXT?.echoPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetOfferInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.echoPrivate.getOfferInfo);
    },
    "call_GetOfferInfo": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.echoPrivate.getOfferInfo(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetOfferInfo": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.echoPrivate.getOfferInfo(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetOobeTimestamp": (): heap.Ref<boolean> => {
      if (WEBEXT?.echoPrivate && "getOobeTimestamp" in WEBEXT?.echoPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetOobeTimestamp": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.echoPrivate.getOobeTimestamp);
    },
    "call_GetOobeTimestamp": (retPtr: Pointer): void => {
      const _ret = WEBEXT.echoPrivate.getOobeTimestamp();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetOobeTimestamp": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.echoPrivate.getOobeTimestamp();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetRegistrationCode": (): heap.Ref<boolean> => {
      if (WEBEXT?.echoPrivate && "getRegistrationCode" in WEBEXT?.echoPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetRegistrationCode": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.echoPrivate.getRegistrationCode);
    },
    "call_GetRegistrationCode": (retPtr: Pointer, type: heap.Ref<object>): void => {
      const _ret = WEBEXT.echoPrivate.getRegistrationCode(A.H.get<object>(type));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetRegistrationCode": (retPtr: Pointer, errPtr: Pointer, type: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.echoPrivate.getRegistrationCode(A.H.get<object>(type));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetUserConsent": (): heap.Ref<boolean> => {
      if (WEBEXT?.echoPrivate && "getUserConsent" in WEBEXT?.echoPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetUserConsent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.echoPrivate.getUserConsent);
    },
    "call_GetUserConsent": (retPtr: Pointer, consentRequester: Pointer): void => {
      const consentRequester_ffi = {};

      consentRequester_ffi["origin"] = A.load.Ref(consentRequester + 0, undefined);
      consentRequester_ffi["serviceName"] = A.load.Ref(consentRequester + 4, undefined);
      if (A.load.Bool(consentRequester + 16)) {
        consentRequester_ffi["tabId"] = A.load.Int64(consentRequester + 8);
      }

      const _ret = WEBEXT.echoPrivate.getUserConsent(consentRequester_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetUserConsent": (retPtr: Pointer, errPtr: Pointer, consentRequester: Pointer): heap.Ref<boolean> => {
      try {
        const consentRequester_ffi = {};

        consentRequester_ffi["origin"] = A.load.Ref(consentRequester + 0, undefined);
        consentRequester_ffi["serviceName"] = A.load.Ref(consentRequester + 4, undefined);
        if (A.load.Bool(consentRequester + 16)) {
          consentRequester_ffi["tabId"] = A.load.Int64(consentRequester + 8);
        }

        const _ret = WEBEXT.echoPrivate.getUserConsent(consentRequester_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetOfferInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.echoPrivate && "setOfferInfo" in WEBEXT?.echoPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetOfferInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.echoPrivate.setOfferInfo);
    },
    "call_SetOfferInfo": (retPtr: Pointer, id: heap.Ref<object>, offerInfo: heap.Ref<object>): void => {
      const _ret = WEBEXT.echoPrivate.setOfferInfo(A.H.get<object>(id), A.H.get<object>(offerInfo));
    },
    "try_SetOfferInfo": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<object>,
      offerInfo: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.echoPrivate.setOfferInfo(A.H.get<object>(id), A.H.get<object>(offerInfo));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
