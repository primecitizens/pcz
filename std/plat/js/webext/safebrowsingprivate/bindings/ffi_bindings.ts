import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/safebrowsingprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_DangerousDownloadInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["url"]);
        A.store.Ref(ptr + 4, x["fileName"]);
        A.store.Ref(ptr + 8, x["downloadDigestSha256"]);
        A.store.Ref(ptr + 12, x["userName"]);
      }
    },
    "load_DangerousDownloadInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      x["fileName"] = A.load.Ref(ptr + 4, undefined);
      x["downloadDigestSha256"] = A.load.Ref(ptr + 8, undefined);
      x["userName"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_URLType": (ref: heap.Ref<string>): number => {
      const idx = [
        "EVENT_URL",
        "LANDING_PAGE",
        "LANDING_REFERRER",
        "CLIENT_REDIRECT",
        "RECENT_NAVIGATION",
        "REFERRER",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ServerRedirect": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["url"]);
      }
    },
    "load_ServerRedirect": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_NavigationInitiation": (ref: heap.Ref<string>): number => {
      const idx = [
        "BROWSER_INITIATED",
        "RENDERER_INITIATED_WITHOUT_USER_GESTURE",
        "RENDERER_INITIATED_WITH_USER_GESTURE",
        "COPY_PASTE_USER_INITIATED",
        "NOTIFICATION_INITIATED",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ReferrerChainEntry": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 58, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Bool(ptr + 52, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 53, false);
        A.store.Float64(ptr + 32, 0);
        A.store.Ref(ptr + 40, undefined);
        A.store.Enum(ptr + 44, -1);
        A.store.Bool(ptr + 54, false);
        A.store.Bool(ptr + 48, false);
        A.store.Bool(ptr + 55, false);
        A.store.Bool(ptr + 49, false);
        A.store.Bool(ptr + 56, false);
        A.store.Bool(ptr + 50, false);
        A.store.Bool(ptr + 57, false);
        A.store.Bool(ptr + 51, false);
      } else {
        A.store.Bool(ptr + 58, true);
        A.store.Ref(ptr + 0, x["url"]);
        A.store.Ref(ptr + 4, x["mainFrameUrl"]);
        A.store.Enum(
          ptr + 8,
          ["EVENT_URL", "LANDING_PAGE", "LANDING_REFERRER", "CLIENT_REDIRECT", "RECENT_NAVIGATION", "REFERRER"].indexOf(
            x["urlType"] as string
          )
        );
        A.store.Ref(ptr + 12, x["ipAddresses"]);
        A.store.Ref(ptr + 16, x["referrerUrl"]);
        A.store.Ref(ptr + 20, x["referrerMainFrameUrl"]);
        A.store.Bool(ptr + 52, "isRetargeting" in x ? true : false);
        A.store.Bool(ptr + 24, x["isRetargeting"] ? true : false);
        A.store.Bool(ptr + 53, "navigationTimeMs" in x ? true : false);
        A.store.Float64(ptr + 32, x["navigationTimeMs"] === undefined ? 0 : (x["navigationTimeMs"] as number));
        A.store.Ref(ptr + 40, x["serverRedirectChain"]);
        A.store.Enum(
          ptr + 44,
          [
            "BROWSER_INITIATED",
            "RENDERER_INITIATED_WITHOUT_USER_GESTURE",
            "RENDERER_INITIATED_WITH_USER_GESTURE",
            "COPY_PASTE_USER_INITIATED",
            "NOTIFICATION_INITIATED",
          ].indexOf(x["navigationInitiation"] as string)
        );
        A.store.Bool(ptr + 54, "maybeLaunchedByExternalApp" in x ? true : false);
        A.store.Bool(ptr + 48, x["maybeLaunchedByExternalApp"] ? true : false);
        A.store.Bool(ptr + 55, "isSubframeUrlRemoved" in x ? true : false);
        A.store.Bool(ptr + 49, x["isSubframeUrlRemoved"] ? true : false);
        A.store.Bool(ptr + 56, "isSubframeReferrerUrlRemoved" in x ? true : false);
        A.store.Bool(ptr + 50, x["isSubframeReferrerUrlRemoved"] ? true : false);
        A.store.Bool(ptr + 57, "isUrlRemovedByPolicy" in x ? true : false);
        A.store.Bool(ptr + 51, x["isUrlRemovedByPolicy"] ? true : false);
      }
    },
    "load_ReferrerChainEntry": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      x["mainFrameUrl"] = A.load.Ref(ptr + 4, undefined);
      x["urlType"] = A.load.Enum(ptr + 8, [
        "EVENT_URL",
        "LANDING_PAGE",
        "LANDING_REFERRER",
        "CLIENT_REDIRECT",
        "RECENT_NAVIGATION",
        "REFERRER",
      ]);
      x["ipAddresses"] = A.load.Ref(ptr + 12, undefined);
      x["referrerUrl"] = A.load.Ref(ptr + 16, undefined);
      x["referrerMainFrameUrl"] = A.load.Ref(ptr + 20, undefined);
      if (A.load.Bool(ptr + 52)) {
        x["isRetargeting"] = A.load.Bool(ptr + 24);
      } else {
        delete x["isRetargeting"];
      }
      if (A.load.Bool(ptr + 53)) {
        x["navigationTimeMs"] = A.load.Float64(ptr + 32);
      } else {
        delete x["navigationTimeMs"];
      }
      x["serverRedirectChain"] = A.load.Ref(ptr + 40, undefined);
      x["navigationInitiation"] = A.load.Enum(ptr + 44, [
        "BROWSER_INITIATED",
        "RENDERER_INITIATED_WITHOUT_USER_GESTURE",
        "RENDERER_INITIATED_WITH_USER_GESTURE",
        "COPY_PASTE_USER_INITIATED",
        "NOTIFICATION_INITIATED",
      ]);
      if (A.load.Bool(ptr + 54)) {
        x["maybeLaunchedByExternalApp"] = A.load.Bool(ptr + 48);
      } else {
        delete x["maybeLaunchedByExternalApp"];
      }
      if (A.load.Bool(ptr + 55)) {
        x["isSubframeUrlRemoved"] = A.load.Bool(ptr + 49);
      } else {
        delete x["isSubframeUrlRemoved"];
      }
      if (A.load.Bool(ptr + 56)) {
        x["isSubframeReferrerUrlRemoved"] = A.load.Bool(ptr + 50);
      } else {
        delete x["isSubframeReferrerUrlRemoved"];
      }
      if (A.load.Bool(ptr + 57)) {
        x["isUrlRemovedByPolicy"] = A.load.Bool(ptr + 51);
      } else {
        delete x["isUrlRemovedByPolicy"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_InterstitialInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["url"]);
        A.store.Ref(ptr + 4, x["reason"]);
        A.store.Ref(ptr + 8, x["netErrorCode"]);
        A.store.Ref(ptr + 12, x["userName"]);
      }
    },
    "load_InterstitialInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      x["reason"] = A.load.Ref(ptr + 4, undefined);
      x["netErrorCode"] = A.load.Ref(ptr + 8, undefined);
      x["userName"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PolicySpecifiedPasswordReuse": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Ref(ptr + 0, x["url"]);
        A.store.Ref(ptr + 4, x["userName"]);
        A.store.Bool(ptr + 9, "isPhishingUrl" in x ? true : false);
        A.store.Bool(ptr + 8, x["isPhishingUrl"] ? true : false);
      }
    },
    "load_PolicySpecifiedPasswordReuse": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      x["userName"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 9)) {
        x["isPhishingUrl"] = A.load.Bool(ptr + 8);
      } else {
        delete x["isPhishingUrl"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetReferrerChain": (): heap.Ref<boolean> => {
      if (WEBEXT?.safeBrowsingPrivate && "getReferrerChain" in WEBEXT?.safeBrowsingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetReferrerChain": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.getReferrerChain);
    },
    "call_GetReferrerChain": (retPtr: Pointer, tabId: number): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.getReferrerChain(tabId);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetReferrerChain": (retPtr: Pointer, errPtr: Pointer, tabId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.getReferrerChain(tabId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDangerousDownloadOpened": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.safeBrowsingPrivate?.onDangerousDownloadOpened &&
        "addListener" in WEBEXT?.safeBrowsingPrivate?.onDangerousDownloadOpened
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDangerousDownloadOpened": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.addListener);
    },
    "call_OnDangerousDownloadOpened": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.addListener(A.H.get<object>(callback));
    },
    "try_OnDangerousDownloadOpened": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDangerousDownloadOpened": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.safeBrowsingPrivate?.onDangerousDownloadOpened &&
        "removeListener" in WEBEXT?.safeBrowsingPrivate?.onDangerousDownloadOpened
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDangerousDownloadOpened": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.removeListener);
    },
    "call_OffDangerousDownloadOpened": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.removeListener(A.H.get<object>(callback));
    },
    "try_OffDangerousDownloadOpened": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDangerousDownloadOpened": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.safeBrowsingPrivate?.onDangerousDownloadOpened &&
        "hasListener" in WEBEXT?.safeBrowsingPrivate?.onDangerousDownloadOpened
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDangerousDownloadOpened": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.hasListener);
    },
    "call_HasOnDangerousDownloadOpened": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDangerousDownloadOpened": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPolicySpecifiedPasswordChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.safeBrowsingPrivate?.onPolicySpecifiedPasswordChanged &&
        "addListener" in WEBEXT?.safeBrowsingPrivate?.onPolicySpecifiedPasswordChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPolicySpecifiedPasswordChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.addListener);
    },
    "call_OnPolicySpecifiedPasswordChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnPolicySpecifiedPasswordChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPolicySpecifiedPasswordChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.safeBrowsingPrivate?.onPolicySpecifiedPasswordChanged &&
        "removeListener" in WEBEXT?.safeBrowsingPrivate?.onPolicySpecifiedPasswordChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPolicySpecifiedPasswordChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.removeListener);
    },
    "call_OffPolicySpecifiedPasswordChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.removeListener(
        A.H.get<object>(callback)
      );
    },
    "try_OffPolicySpecifiedPasswordChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPolicySpecifiedPasswordChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.safeBrowsingPrivate?.onPolicySpecifiedPasswordChanged &&
        "hasListener" in WEBEXT?.safeBrowsingPrivate?.onPolicySpecifiedPasswordChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPolicySpecifiedPasswordChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.hasListener);
    },
    "call_HasOnPolicySpecifiedPasswordChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPolicySpecifiedPasswordChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPolicySpecifiedPasswordReuseDetected": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.safeBrowsingPrivate?.onPolicySpecifiedPasswordReuseDetected &&
        "addListener" in WEBEXT?.safeBrowsingPrivate?.onPolicySpecifiedPasswordReuseDetected
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPolicySpecifiedPasswordReuseDetected": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.addListener);
    },
    "call_OnPolicySpecifiedPasswordReuseDetected": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.addListener(
        A.H.get<object>(callback)
      );
    },
    "try_OnPolicySpecifiedPasswordReuseDetected": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.addListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPolicySpecifiedPasswordReuseDetected": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.safeBrowsingPrivate?.onPolicySpecifiedPasswordReuseDetected &&
        "removeListener" in WEBEXT?.safeBrowsingPrivate?.onPolicySpecifiedPasswordReuseDetected
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPolicySpecifiedPasswordReuseDetected": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.removeListener);
    },
    "call_OffPolicySpecifiedPasswordReuseDetected": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.removeListener(
        A.H.get<object>(callback)
      );
    },
    "try_OffPolicySpecifiedPasswordReuseDetected": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPolicySpecifiedPasswordReuseDetected": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.safeBrowsingPrivate?.onPolicySpecifiedPasswordReuseDetected &&
        "hasListener" in WEBEXT?.safeBrowsingPrivate?.onPolicySpecifiedPasswordReuseDetected
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPolicySpecifiedPasswordReuseDetected": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.hasListener);
    },
    "call_HasOnPolicySpecifiedPasswordReuseDetected": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.hasListener(
        A.H.get<object>(callback)
      );
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPolicySpecifiedPasswordReuseDetected": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.hasListener(
          A.H.get<object>(callback)
        );
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSecurityInterstitialProceeded": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.safeBrowsingPrivate?.onSecurityInterstitialProceeded &&
        "addListener" in WEBEXT?.safeBrowsingPrivate?.onSecurityInterstitialProceeded
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSecurityInterstitialProceeded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.addListener);
    },
    "call_OnSecurityInterstitialProceeded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.addListener(A.H.get<object>(callback));
    },
    "try_OnSecurityInterstitialProceeded": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSecurityInterstitialProceeded": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.safeBrowsingPrivate?.onSecurityInterstitialProceeded &&
        "removeListener" in WEBEXT?.safeBrowsingPrivate?.onSecurityInterstitialProceeded
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSecurityInterstitialProceeded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.removeListener);
    },
    "call_OffSecurityInterstitialProceeded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.removeListener(A.H.get<object>(callback));
    },
    "try_OffSecurityInterstitialProceeded": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSecurityInterstitialProceeded": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.safeBrowsingPrivate?.onSecurityInterstitialProceeded &&
        "hasListener" in WEBEXT?.safeBrowsingPrivate?.onSecurityInterstitialProceeded
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSecurityInterstitialProceeded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.hasListener);
    },
    "call_HasOnSecurityInterstitialProceeded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSecurityInterstitialProceeded": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSecurityInterstitialShown": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.safeBrowsingPrivate?.onSecurityInterstitialShown &&
        "addListener" in WEBEXT?.safeBrowsingPrivate?.onSecurityInterstitialShown
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSecurityInterstitialShown": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.addListener);
    },
    "call_OnSecurityInterstitialShown": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.addListener(A.H.get<object>(callback));
    },
    "try_OnSecurityInterstitialShown": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSecurityInterstitialShown": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.safeBrowsingPrivate?.onSecurityInterstitialShown &&
        "removeListener" in WEBEXT?.safeBrowsingPrivate?.onSecurityInterstitialShown
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSecurityInterstitialShown": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.removeListener);
    },
    "call_OffSecurityInterstitialShown": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.removeListener(A.H.get<object>(callback));
    },
    "try_OffSecurityInterstitialShown": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSecurityInterstitialShown": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.safeBrowsingPrivate?.onSecurityInterstitialShown &&
        "hasListener" in WEBEXT?.safeBrowsingPrivate?.onSecurityInterstitialShown
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSecurityInterstitialShown": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.hasListener);
    },
    "call_HasOnSecurityInterstitialShown": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSecurityInterstitialShown": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
