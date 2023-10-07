import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/privacy", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_IPHandlingPolicy": (ref: heap.Ref<string>): number => {
      const idx = [
        "default",
        "default_public_and_private_interfaces",
        "default_public_interface_only",
        "disable_non_proxied_udp",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_NetworkProperty": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["networkPredictionEnabled"]);
        A.store.Ref(ptr + 4, x["webRTCIPHandlingPolicy"]);
      }
    },
    "load_NetworkProperty": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["networkPredictionEnabled"] = A.load.Ref(ptr + 0, undefined);
      x["webRTCIPHandlingPolicy"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ServicesProperty": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 40, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
      } else {
        A.store.Bool(ptr + 40, true);
        A.store.Ref(ptr + 0, x["alternateErrorPagesEnabled"]);
        A.store.Ref(ptr + 4, x["autofillAddressEnabled"]);
        A.store.Ref(ptr + 8, x["autofillCreditCardEnabled"]);
        A.store.Ref(ptr + 12, x["autofillEnabled"]);
        A.store.Ref(ptr + 16, x["passwordSavingEnabled"]);
        A.store.Ref(ptr + 20, x["safeBrowsingEnabled"]);
        A.store.Ref(ptr + 24, x["safeBrowsingExtendedReportingEnabled"]);
        A.store.Ref(ptr + 28, x["searchSuggestEnabled"]);
        A.store.Ref(ptr + 32, x["spellingServiceEnabled"]);
        A.store.Ref(ptr + 36, x["translationServiceEnabled"]);
      }
    },
    "load_ServicesProperty": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["alternateErrorPagesEnabled"] = A.load.Ref(ptr + 0, undefined);
      x["autofillAddressEnabled"] = A.load.Ref(ptr + 4, undefined);
      x["autofillCreditCardEnabled"] = A.load.Ref(ptr + 8, undefined);
      x["autofillEnabled"] = A.load.Ref(ptr + 12, undefined);
      x["passwordSavingEnabled"] = A.load.Ref(ptr + 16, undefined);
      x["safeBrowsingEnabled"] = A.load.Ref(ptr + 20, undefined);
      x["safeBrowsingExtendedReportingEnabled"] = A.load.Ref(ptr + 24, undefined);
      x["searchSuggestEnabled"] = A.load.Ref(ptr + 28, undefined);
      x["spellingServiceEnabled"] = A.load.Ref(ptr + 32, undefined);
      x["translationServiceEnabled"] = A.load.Ref(ptr + 36, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_WebsitesProperty": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 32, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
      } else {
        A.store.Bool(ptr + 32, true);
        A.store.Ref(ptr + 0, x["adMeasurementEnabled"]);
        A.store.Ref(ptr + 4, x["doNotTrackEnabled"]);
        A.store.Ref(ptr + 8, x["fledgeEnabled"]);
        A.store.Ref(ptr + 12, x["hyperlinkAuditingEnabled"]);
        A.store.Ref(ptr + 16, x["protectedContentEnabled"]);
        A.store.Ref(ptr + 20, x["referrersEnabled"]);
        A.store.Ref(ptr + 24, x["thirdPartyCookiesAllowed"]);
        A.store.Ref(ptr + 28, x["topicsEnabled"]);
      }
    },
    "load_WebsitesProperty": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["adMeasurementEnabled"] = A.load.Ref(ptr + 0, undefined);
      x["doNotTrackEnabled"] = A.load.Ref(ptr + 4, undefined);
      x["fledgeEnabled"] = A.load.Ref(ptr + 8, undefined);
      x["hyperlinkAuditingEnabled"] = A.load.Ref(ptr + 12, undefined);
      x["protectedContentEnabled"] = A.load.Ref(ptr + 16, undefined);
      x["referrersEnabled"] = A.load.Ref(ptr + 20, undefined);
      x["thirdPartyCookiesAllowed"] = A.load.Ref(ptr + 24, undefined);
      x["topicsEnabled"] = A.load.Ref(ptr + 28, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "get_Network": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.privacy && "network" in WEBEXT?.privacy) {
        const val = WEBEXT.privacy.network;
        if (typeof val === "undefined") {
          A.store.Bool(retPtr + 8, false);
          A.store.Ref(retPtr + 0, undefined);
          A.store.Ref(retPtr + 4, undefined);
        } else {
          A.store.Bool(retPtr + 8, true);
          A.store.Ref(retPtr + 0, val["networkPredictionEnabled"]);
          A.store.Ref(retPtr + 4, val["webRTCIPHandlingPolicy"]);
        }
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Network": (val: Pointer): heap.Ref<boolean> => {
      const val_ffi = {};

      val_ffi["networkPredictionEnabled"] = A.load.Ref(val + 0, undefined);
      val_ffi["webRTCIPHandlingPolicy"] = A.load.Ref(val + 4, undefined);
      return Reflect.set(WEBEXT.privacy, "network", val_ffi, WEBEXT.privacy) ? A.H.TRUE : A.H.FALSE;
    },
    "get_Services": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.privacy && "services" in WEBEXT?.privacy) {
        const val = WEBEXT.privacy.services;
        if (typeof val === "undefined") {
          A.store.Bool(retPtr + 40, false);
          A.store.Ref(retPtr + 0, undefined);
          A.store.Ref(retPtr + 4, undefined);
          A.store.Ref(retPtr + 8, undefined);
          A.store.Ref(retPtr + 12, undefined);
          A.store.Ref(retPtr + 16, undefined);
          A.store.Ref(retPtr + 20, undefined);
          A.store.Ref(retPtr + 24, undefined);
          A.store.Ref(retPtr + 28, undefined);
          A.store.Ref(retPtr + 32, undefined);
          A.store.Ref(retPtr + 36, undefined);
        } else {
          A.store.Bool(retPtr + 40, true);
          A.store.Ref(retPtr + 0, val["alternateErrorPagesEnabled"]);
          A.store.Ref(retPtr + 4, val["autofillAddressEnabled"]);
          A.store.Ref(retPtr + 8, val["autofillCreditCardEnabled"]);
          A.store.Ref(retPtr + 12, val["autofillEnabled"]);
          A.store.Ref(retPtr + 16, val["passwordSavingEnabled"]);
          A.store.Ref(retPtr + 20, val["safeBrowsingEnabled"]);
          A.store.Ref(retPtr + 24, val["safeBrowsingExtendedReportingEnabled"]);
          A.store.Ref(retPtr + 28, val["searchSuggestEnabled"]);
          A.store.Ref(retPtr + 32, val["spellingServiceEnabled"]);
          A.store.Ref(retPtr + 36, val["translationServiceEnabled"]);
        }
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Services": (val: Pointer): heap.Ref<boolean> => {
      const val_ffi = {};

      val_ffi["alternateErrorPagesEnabled"] = A.load.Ref(val + 0, undefined);
      val_ffi["autofillAddressEnabled"] = A.load.Ref(val + 4, undefined);
      val_ffi["autofillCreditCardEnabled"] = A.load.Ref(val + 8, undefined);
      val_ffi["autofillEnabled"] = A.load.Ref(val + 12, undefined);
      val_ffi["passwordSavingEnabled"] = A.load.Ref(val + 16, undefined);
      val_ffi["safeBrowsingEnabled"] = A.load.Ref(val + 20, undefined);
      val_ffi["safeBrowsingExtendedReportingEnabled"] = A.load.Ref(val + 24, undefined);
      val_ffi["searchSuggestEnabled"] = A.load.Ref(val + 28, undefined);
      val_ffi["spellingServiceEnabled"] = A.load.Ref(val + 32, undefined);
      val_ffi["translationServiceEnabled"] = A.load.Ref(val + 36, undefined);
      return Reflect.set(WEBEXT.privacy, "services", val_ffi, WEBEXT.privacy) ? A.H.TRUE : A.H.FALSE;
    },
    "get_Websites": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.privacy && "websites" in WEBEXT?.privacy) {
        const val = WEBEXT.privacy.websites;
        if (typeof val === "undefined") {
          A.store.Bool(retPtr + 32, false);
          A.store.Ref(retPtr + 0, undefined);
          A.store.Ref(retPtr + 4, undefined);
          A.store.Ref(retPtr + 8, undefined);
          A.store.Ref(retPtr + 12, undefined);
          A.store.Ref(retPtr + 16, undefined);
          A.store.Ref(retPtr + 20, undefined);
          A.store.Ref(retPtr + 24, undefined);
          A.store.Ref(retPtr + 28, undefined);
        } else {
          A.store.Bool(retPtr + 32, true);
          A.store.Ref(retPtr + 0, val["adMeasurementEnabled"]);
          A.store.Ref(retPtr + 4, val["doNotTrackEnabled"]);
          A.store.Ref(retPtr + 8, val["fledgeEnabled"]);
          A.store.Ref(retPtr + 12, val["hyperlinkAuditingEnabled"]);
          A.store.Ref(retPtr + 16, val["protectedContentEnabled"]);
          A.store.Ref(retPtr + 20, val["referrersEnabled"]);
          A.store.Ref(retPtr + 24, val["thirdPartyCookiesAllowed"]);
          A.store.Ref(retPtr + 28, val["topicsEnabled"]);
        }
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Websites": (val: Pointer): heap.Ref<boolean> => {
      const val_ffi = {};

      val_ffi["adMeasurementEnabled"] = A.load.Ref(val + 0, undefined);
      val_ffi["doNotTrackEnabled"] = A.load.Ref(val + 4, undefined);
      val_ffi["fledgeEnabled"] = A.load.Ref(val + 8, undefined);
      val_ffi["hyperlinkAuditingEnabled"] = A.load.Ref(val + 12, undefined);
      val_ffi["protectedContentEnabled"] = A.load.Ref(val + 16, undefined);
      val_ffi["referrersEnabled"] = A.load.Ref(val + 20, undefined);
      val_ffi["thirdPartyCookiesAllowed"] = A.load.Ref(val + 24, undefined);
      val_ffi["topicsEnabled"] = A.load.Ref(val + 28, undefined);
      return Reflect.set(WEBEXT.privacy, "websites", val_ffi, WEBEXT.privacy) ? A.H.TRUE : A.H.FALSE;
    },
  };
});
