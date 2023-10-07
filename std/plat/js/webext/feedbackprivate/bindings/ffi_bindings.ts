import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/feedbackprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AttachedFile": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["data"]);
      }
    },
    "load_AttachedFile": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["data"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_FeedbackFlow": (ref: heap.Ref<string>): number => {
      const idx = ["regular", "login", "sadTabCrash", "googleInternal"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_LogsMapEntry": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["key"]);
        A.store.Ref(ptr + 4, x["value"]);
      }
    },
    "load_LogsMapEntry": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["key"] = A.load.Ref(ptr + 0, undefined);
      x["value"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_FeedbackInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 89, false);

        A.store.Bool(ptr + 0 + 8, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Bool(ptr + 77, false);
        A.store.Int32(ptr + 32, 0);
        A.store.Ref(ptr + 36, undefined);
        A.store.Bool(ptr + 78, false);
        A.store.Int32(ptr + 40, 0);
        A.store.Ref(ptr + 44, undefined);
        A.store.Bool(ptr + 79, false);
        A.store.Bool(ptr + 48, false);
        A.store.Enum(ptr + 52, -1);
        A.store.Ref(ptr + 56, undefined);
        A.store.Ref(ptr + 60, undefined);
        A.store.Bool(ptr + 80, false);
        A.store.Bool(ptr + 64, false);
        A.store.Bool(ptr + 81, false);
        A.store.Bool(ptr + 65, false);
        A.store.Bool(ptr + 82, false);
        A.store.Bool(ptr + 66, false);
        A.store.Bool(ptr + 83, false);
        A.store.Bool(ptr + 67, false);
        A.store.Bool(ptr + 84, false);
        A.store.Bool(ptr + 68, false);
        A.store.Bool(ptr + 85, false);
        A.store.Bool(ptr + 69, false);
        A.store.Bool(ptr + 86, false);
        A.store.Bool(ptr + 70, false);
        A.store.Bool(ptr + 87, false);
        A.store.Bool(ptr + 71, false);
        A.store.Ref(ptr + 72, undefined);
        A.store.Bool(ptr + 88, false);
        A.store.Bool(ptr + 76, false);
      } else {
        A.store.Bool(ptr + 89, true);

        if (typeof x["attachedFile"] === "undefined") {
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
        } else {
          A.store.Bool(ptr + 0 + 8, true);
          A.store.Ref(ptr + 0 + 0, x["attachedFile"]["name"]);
          A.store.Ref(ptr + 0 + 4, x["attachedFile"]["data"]);
        }
        A.store.Ref(ptr + 12, x["categoryTag"]);
        A.store.Ref(ptr + 16, x["description"]);
        A.store.Ref(ptr + 20, x["descriptionPlaceholder"]);
        A.store.Ref(ptr + 24, x["email"]);
        A.store.Ref(ptr + 28, x["pageUrl"]);
        A.store.Bool(ptr + 77, "productId" in x ? true : false);
        A.store.Int32(ptr + 32, x["productId"] === undefined ? 0 : (x["productId"] as number));
        A.store.Ref(ptr + 36, x["screenshot"]);
        A.store.Bool(ptr + 78, "traceId" in x ? true : false);
        A.store.Int32(ptr + 40, x["traceId"] === undefined ? 0 : (x["traceId"] as number));
        A.store.Ref(ptr + 44, x["systemInformation"]);
        A.store.Bool(ptr + 79, "sendHistograms" in x ? true : false);
        A.store.Bool(ptr + 48, x["sendHistograms"] ? true : false);
        A.store.Enum(ptr + 52, ["regular", "login", "sadTabCrash", "googleInternal"].indexOf(x["flow"] as string));
        A.store.Ref(ptr + 56, x["attachedFileBlobUuid"]);
        A.store.Ref(ptr + 60, x["screenshotBlobUuid"]);
        A.store.Bool(ptr + 80, "useSystemWindowFrame" in x ? true : false);
        A.store.Bool(ptr + 64, x["useSystemWindowFrame"] ? true : false);
        A.store.Bool(ptr + 81, "sendBluetoothLogs" in x ? true : false);
        A.store.Bool(ptr + 65, x["sendBluetoothLogs"] ? true : false);
        A.store.Bool(ptr + 82, "sendTabTitles" in x ? true : false);
        A.store.Bool(ptr + 66, x["sendTabTitles"] ? true : false);
        A.store.Bool(ptr + 83, "assistantDebugInfoAllowed" in x ? true : false);
        A.store.Bool(ptr + 67, x["assistantDebugInfoAllowed"] ? true : false);
        A.store.Bool(ptr + 84, "fromAssistant" in x ? true : false);
        A.store.Bool(ptr + 68, x["fromAssistant"] ? true : false);
        A.store.Bool(ptr + 85, "includeBluetoothLogs" in x ? true : false);
        A.store.Bool(ptr + 69, x["includeBluetoothLogs"] ? true : false);
        A.store.Bool(ptr + 86, "showQuestionnaire" in x ? true : false);
        A.store.Bool(ptr + 70, x["showQuestionnaire"] ? true : false);
        A.store.Bool(ptr + 87, "fromAutofill" in x ? true : false);
        A.store.Bool(ptr + 71, x["fromAutofill"] ? true : false);
        A.store.Ref(ptr + 72, x["autofillMetadata"]);
        A.store.Bool(ptr + 88, "sendAutofillMetadata" in x ? true : false);
        A.store.Bool(ptr + 76, x["sendAutofillMetadata"] ? true : false);
      }
    },
    "load_FeedbackInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 8)) {
        x["attachedFile"] = {};
        x["attachedFile"]["name"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["attachedFile"]["data"] = A.load.Ref(ptr + 0 + 4, undefined);
      } else {
        delete x["attachedFile"];
      }
      x["categoryTag"] = A.load.Ref(ptr + 12, undefined);
      x["description"] = A.load.Ref(ptr + 16, undefined);
      x["descriptionPlaceholder"] = A.load.Ref(ptr + 20, undefined);
      x["email"] = A.load.Ref(ptr + 24, undefined);
      x["pageUrl"] = A.load.Ref(ptr + 28, undefined);
      if (A.load.Bool(ptr + 77)) {
        x["productId"] = A.load.Int32(ptr + 32);
      } else {
        delete x["productId"];
      }
      x["screenshot"] = A.load.Ref(ptr + 36, undefined);
      if (A.load.Bool(ptr + 78)) {
        x["traceId"] = A.load.Int32(ptr + 40);
      } else {
        delete x["traceId"];
      }
      x["systemInformation"] = A.load.Ref(ptr + 44, undefined);
      if (A.load.Bool(ptr + 79)) {
        x["sendHistograms"] = A.load.Bool(ptr + 48);
      } else {
        delete x["sendHistograms"];
      }
      x["flow"] = A.load.Enum(ptr + 52, ["regular", "login", "sadTabCrash", "googleInternal"]);
      x["attachedFileBlobUuid"] = A.load.Ref(ptr + 56, undefined);
      x["screenshotBlobUuid"] = A.load.Ref(ptr + 60, undefined);
      if (A.load.Bool(ptr + 80)) {
        x["useSystemWindowFrame"] = A.load.Bool(ptr + 64);
      } else {
        delete x["useSystemWindowFrame"];
      }
      if (A.load.Bool(ptr + 81)) {
        x["sendBluetoothLogs"] = A.load.Bool(ptr + 65);
      } else {
        delete x["sendBluetoothLogs"];
      }
      if (A.load.Bool(ptr + 82)) {
        x["sendTabTitles"] = A.load.Bool(ptr + 66);
      } else {
        delete x["sendTabTitles"];
      }
      if (A.load.Bool(ptr + 83)) {
        x["assistantDebugInfoAllowed"] = A.load.Bool(ptr + 67);
      } else {
        delete x["assistantDebugInfoAllowed"];
      }
      if (A.load.Bool(ptr + 84)) {
        x["fromAssistant"] = A.load.Bool(ptr + 68);
      } else {
        delete x["fromAssistant"];
      }
      if (A.load.Bool(ptr + 85)) {
        x["includeBluetoothLogs"] = A.load.Bool(ptr + 69);
      } else {
        delete x["includeBluetoothLogs"];
      }
      if (A.load.Bool(ptr + 86)) {
        x["showQuestionnaire"] = A.load.Bool(ptr + 70);
      } else {
        delete x["showQuestionnaire"];
      }
      if (A.load.Bool(ptr + 87)) {
        x["fromAutofill"] = A.load.Bool(ptr + 71);
      } else {
        delete x["fromAutofill"];
      }
      x["autofillMetadata"] = A.load.Ref(ptr + 72, undefined);
      if (A.load.Bool(ptr + 88)) {
        x["sendAutofillMetadata"] = A.load.Bool(ptr + 76);
      } else {
        delete x["sendAutofillMetadata"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_FeedbackSource": (ref: heap.Ref<string>): number => {
      const idx = ["quickoffice"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_LandingPageType": (ref: heap.Ref<string>): number => {
      const idx = ["normal", "techstop", "noLandingPage"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_LogSource": (ref: heap.Ref<string>): number => {
      const idx = [
        "messages",
        "uiLatest",
        "drmModetest",
        "lsusb",
        "atrusLog",
        "netLog",
        "eventLog",
        "updateEngineLog",
        "powerdLatest",
        "powerdPrevious",
        "lspci",
        "ifconfig",
        "uptime",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ReadLogSourceResult": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "readerId" in x ? true : false);
        A.store.Int32(ptr + 0, x["readerId"] === undefined ? 0 : (x["readerId"] as number));
        A.store.Ref(ptr + 4, x["logLines"]);
      }
    },
    "load_ReadLogSourceResult": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["readerId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["readerId"];
      }
      x["logLines"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ReadLogSourceParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Enum(
          ptr + 0,
          [
            "messages",
            "uiLatest",
            "drmModetest",
            "lsusb",
            "atrusLog",
            "netLog",
            "eventLog",
            "updateEngineLog",
            "powerdLatest",
            "powerdPrevious",
            "lspci",
            "ifconfig",
            "uptime",
          ].indexOf(x["source"] as string)
        );
        A.store.Bool(ptr + 12, "incremental" in x ? true : false);
        A.store.Bool(ptr + 4, x["incremental"] ? true : false);
        A.store.Bool(ptr + 13, "readerId" in x ? true : false);
        A.store.Int32(ptr + 8, x["readerId"] === undefined ? 0 : (x["readerId"] as number));
      }
    },
    "load_ReadLogSourceParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["source"] = A.load.Enum(ptr + 0, [
        "messages",
        "uiLatest",
        "drmModetest",
        "lsusb",
        "atrusLog",
        "netLog",
        "eventLog",
        "updateEngineLog",
        "powerdLatest",
        "powerdPrevious",
        "lspci",
        "ifconfig",
        "uptime",
      ]);
      if (A.load.Bool(ptr + 12)) {
        x["incremental"] = A.load.Bool(ptr + 4);
      } else {
        delete x["incremental"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["readerId"] = A.load.Int32(ptr + 8);
      } else {
        delete x["readerId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_Status": (ref: heap.Ref<string>): number => {
      const idx = ["success", "delayed"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SendFeedbackResult": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["success", "delayed"].indexOf(x["status"] as string));
        A.store.Enum(ptr + 4, ["normal", "techstop", "noLandingPage"].indexOf(x["landingPageType"] as string));
      }
    },
    "load_SendFeedbackResult": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["status"] = A.load.Enum(ptr + 0, ["success", "delayed"]);
      x["landingPageType"] = A.load.Enum(ptr + 4, ["normal", "techstop", "noLandingPage"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetSystemInformation": (): heap.Ref<boolean> => {
      if (WEBEXT?.feedbackPrivate && "getSystemInformation" in WEBEXT?.feedbackPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSystemInformation": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.feedbackPrivate.getSystemInformation);
    },
    "call_GetSystemInformation": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.feedbackPrivate.getSystemInformation(A.H.get<object>(callback));
    },
    "try_GetSystemInformation": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.feedbackPrivate.getSystemInformation(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetUserEmail": (): heap.Ref<boolean> => {
      if (WEBEXT?.feedbackPrivate && "getUserEmail" in WEBEXT?.feedbackPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetUserEmail": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.feedbackPrivate.getUserEmail);
    },
    "call_GetUserEmail": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.feedbackPrivate.getUserEmail(A.H.get<object>(callback));
    },
    "try_GetUserEmail": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.feedbackPrivate.getUserEmail(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnFeedbackRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.feedbackPrivate?.onFeedbackRequested &&
        "addListener" in WEBEXT?.feedbackPrivate?.onFeedbackRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnFeedbackRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.feedbackPrivate.onFeedbackRequested.addListener);
    },
    "call_OnFeedbackRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.feedbackPrivate.onFeedbackRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnFeedbackRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.feedbackPrivate.onFeedbackRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffFeedbackRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.feedbackPrivate?.onFeedbackRequested &&
        "removeListener" in WEBEXT?.feedbackPrivate?.onFeedbackRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffFeedbackRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.feedbackPrivate.onFeedbackRequested.removeListener);
    },
    "call_OffFeedbackRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.feedbackPrivate.onFeedbackRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffFeedbackRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.feedbackPrivate.onFeedbackRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnFeedbackRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.feedbackPrivate?.onFeedbackRequested &&
        "hasListener" in WEBEXT?.feedbackPrivate?.onFeedbackRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnFeedbackRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.feedbackPrivate.onFeedbackRequested.hasListener);
    },
    "call_HasOnFeedbackRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.feedbackPrivate.onFeedbackRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnFeedbackRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.feedbackPrivate.onFeedbackRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenFeedback": (): heap.Ref<boolean> => {
      if (WEBEXT?.feedbackPrivate && "openFeedback" in WEBEXT?.feedbackPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenFeedback": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.feedbackPrivate.openFeedback);
    },
    "call_OpenFeedback": (retPtr: Pointer, source: number): void => {
      const _ret = WEBEXT.feedbackPrivate.openFeedback(
        source > 0 && source <= 1 ? ["quickoffice"][source - 1] : undefined
      );
    },
    "try_OpenFeedback": (retPtr: Pointer, errPtr: Pointer, source: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.feedbackPrivate.openFeedback(
          source > 0 && source <= 1 ? ["quickoffice"][source - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReadLogSource": (): heap.Ref<boolean> => {
      if (WEBEXT?.feedbackPrivate && "readLogSource" in WEBEXT?.feedbackPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReadLogSource": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.feedbackPrivate.readLogSource);
    },
    "call_ReadLogSource": (retPtr: Pointer, params: Pointer, callback: heap.Ref<object>): void => {
      const params_ffi = {};

      params_ffi["source"] = A.load.Enum(params + 0, [
        "messages",
        "uiLatest",
        "drmModetest",
        "lsusb",
        "atrusLog",
        "netLog",
        "eventLog",
        "updateEngineLog",
        "powerdLatest",
        "powerdPrevious",
        "lspci",
        "ifconfig",
        "uptime",
      ]);
      if (A.load.Bool(params + 12)) {
        params_ffi["incremental"] = A.load.Bool(params + 4);
      }
      if (A.load.Bool(params + 13)) {
        params_ffi["readerId"] = A.load.Int32(params + 8);
      }

      const _ret = WEBEXT.feedbackPrivate.readLogSource(params_ffi, A.H.get<object>(callback));
    },
    "try_ReadLogSource": (
      retPtr: Pointer,
      errPtr: Pointer,
      params: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const params_ffi = {};

        params_ffi["source"] = A.load.Enum(params + 0, [
          "messages",
          "uiLatest",
          "drmModetest",
          "lsusb",
          "atrusLog",
          "netLog",
          "eventLog",
          "updateEngineLog",
          "powerdLatest",
          "powerdPrevious",
          "lspci",
          "ifconfig",
          "uptime",
        ]);
        if (A.load.Bool(params + 12)) {
          params_ffi["incremental"] = A.load.Bool(params + 4);
        }
        if (A.load.Bool(params + 13)) {
          params_ffi["readerId"] = A.load.Int32(params + 8);
        }

        const _ret = WEBEXT.feedbackPrivate.readLogSource(params_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendFeedback": (): heap.Ref<boolean> => {
      if (WEBEXT?.feedbackPrivate && "sendFeedback" in WEBEXT?.feedbackPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendFeedback": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.feedbackPrivate.sendFeedback);
    },
    "call_SendFeedback": (
      retPtr: Pointer,
      feedback: Pointer,
      loadSystemInfo: heap.Ref<boolean>,
      formOpenTime: number
    ): void => {
      const feedback_ffi = {};

      if (A.load.Bool(feedback + 0 + 8)) {
        feedback_ffi["attachedFile"] = {};
        feedback_ffi["attachedFile"]["name"] = A.load.Ref(feedback + 0 + 0, undefined);
        feedback_ffi["attachedFile"]["data"] = A.load.Ref(feedback + 0 + 4, undefined);
      }
      feedback_ffi["categoryTag"] = A.load.Ref(feedback + 12, undefined);
      feedback_ffi["description"] = A.load.Ref(feedback + 16, undefined);
      feedback_ffi["descriptionPlaceholder"] = A.load.Ref(feedback + 20, undefined);
      feedback_ffi["email"] = A.load.Ref(feedback + 24, undefined);
      feedback_ffi["pageUrl"] = A.load.Ref(feedback + 28, undefined);
      if (A.load.Bool(feedback + 77)) {
        feedback_ffi["productId"] = A.load.Int32(feedback + 32);
      }
      feedback_ffi["screenshot"] = A.load.Ref(feedback + 36, undefined);
      if (A.load.Bool(feedback + 78)) {
        feedback_ffi["traceId"] = A.load.Int32(feedback + 40);
      }
      feedback_ffi["systemInformation"] = A.load.Ref(feedback + 44, undefined);
      if (A.load.Bool(feedback + 79)) {
        feedback_ffi["sendHistograms"] = A.load.Bool(feedback + 48);
      }
      feedback_ffi["flow"] = A.load.Enum(feedback + 52, ["regular", "login", "sadTabCrash", "googleInternal"]);
      feedback_ffi["attachedFileBlobUuid"] = A.load.Ref(feedback + 56, undefined);
      feedback_ffi["screenshotBlobUuid"] = A.load.Ref(feedback + 60, undefined);
      if (A.load.Bool(feedback + 80)) {
        feedback_ffi["useSystemWindowFrame"] = A.load.Bool(feedback + 64);
      }
      if (A.load.Bool(feedback + 81)) {
        feedback_ffi["sendBluetoothLogs"] = A.load.Bool(feedback + 65);
      }
      if (A.load.Bool(feedback + 82)) {
        feedback_ffi["sendTabTitles"] = A.load.Bool(feedback + 66);
      }
      if (A.load.Bool(feedback + 83)) {
        feedback_ffi["assistantDebugInfoAllowed"] = A.load.Bool(feedback + 67);
      }
      if (A.load.Bool(feedback + 84)) {
        feedback_ffi["fromAssistant"] = A.load.Bool(feedback + 68);
      }
      if (A.load.Bool(feedback + 85)) {
        feedback_ffi["includeBluetoothLogs"] = A.load.Bool(feedback + 69);
      }
      if (A.load.Bool(feedback + 86)) {
        feedback_ffi["showQuestionnaire"] = A.load.Bool(feedback + 70);
      }
      if (A.load.Bool(feedback + 87)) {
        feedback_ffi["fromAutofill"] = A.load.Bool(feedback + 71);
      }
      feedback_ffi["autofillMetadata"] = A.load.Ref(feedback + 72, undefined);
      if (A.load.Bool(feedback + 88)) {
        feedback_ffi["sendAutofillMetadata"] = A.load.Bool(feedback + 76);
      }

      const _ret = WEBEXT.feedbackPrivate.sendFeedback(feedback_ffi, loadSystemInfo === A.H.TRUE, formOpenTime);
      A.store.Ref(retPtr, _ret);
    },
    "try_SendFeedback": (
      retPtr: Pointer,
      errPtr: Pointer,
      feedback: Pointer,
      loadSystemInfo: heap.Ref<boolean>,
      formOpenTime: number
    ): heap.Ref<boolean> => {
      try {
        const feedback_ffi = {};

        if (A.load.Bool(feedback + 0 + 8)) {
          feedback_ffi["attachedFile"] = {};
          feedback_ffi["attachedFile"]["name"] = A.load.Ref(feedback + 0 + 0, undefined);
          feedback_ffi["attachedFile"]["data"] = A.load.Ref(feedback + 0 + 4, undefined);
        }
        feedback_ffi["categoryTag"] = A.load.Ref(feedback + 12, undefined);
        feedback_ffi["description"] = A.load.Ref(feedback + 16, undefined);
        feedback_ffi["descriptionPlaceholder"] = A.load.Ref(feedback + 20, undefined);
        feedback_ffi["email"] = A.load.Ref(feedback + 24, undefined);
        feedback_ffi["pageUrl"] = A.load.Ref(feedback + 28, undefined);
        if (A.load.Bool(feedback + 77)) {
          feedback_ffi["productId"] = A.load.Int32(feedback + 32);
        }
        feedback_ffi["screenshot"] = A.load.Ref(feedback + 36, undefined);
        if (A.load.Bool(feedback + 78)) {
          feedback_ffi["traceId"] = A.load.Int32(feedback + 40);
        }
        feedback_ffi["systemInformation"] = A.load.Ref(feedback + 44, undefined);
        if (A.load.Bool(feedback + 79)) {
          feedback_ffi["sendHistograms"] = A.load.Bool(feedback + 48);
        }
        feedback_ffi["flow"] = A.load.Enum(feedback + 52, ["regular", "login", "sadTabCrash", "googleInternal"]);
        feedback_ffi["attachedFileBlobUuid"] = A.load.Ref(feedback + 56, undefined);
        feedback_ffi["screenshotBlobUuid"] = A.load.Ref(feedback + 60, undefined);
        if (A.load.Bool(feedback + 80)) {
          feedback_ffi["useSystemWindowFrame"] = A.load.Bool(feedback + 64);
        }
        if (A.load.Bool(feedback + 81)) {
          feedback_ffi["sendBluetoothLogs"] = A.load.Bool(feedback + 65);
        }
        if (A.load.Bool(feedback + 82)) {
          feedback_ffi["sendTabTitles"] = A.load.Bool(feedback + 66);
        }
        if (A.load.Bool(feedback + 83)) {
          feedback_ffi["assistantDebugInfoAllowed"] = A.load.Bool(feedback + 67);
        }
        if (A.load.Bool(feedback + 84)) {
          feedback_ffi["fromAssistant"] = A.load.Bool(feedback + 68);
        }
        if (A.load.Bool(feedback + 85)) {
          feedback_ffi["includeBluetoothLogs"] = A.load.Bool(feedback + 69);
        }
        if (A.load.Bool(feedback + 86)) {
          feedback_ffi["showQuestionnaire"] = A.load.Bool(feedback + 70);
        }
        if (A.load.Bool(feedback + 87)) {
          feedback_ffi["fromAutofill"] = A.load.Bool(feedback + 71);
        }
        feedback_ffi["autofillMetadata"] = A.load.Ref(feedback + 72, undefined);
        if (A.load.Bool(feedback + 88)) {
          feedback_ffi["sendAutofillMetadata"] = A.load.Bool(feedback + 76);
        }

        const _ret = WEBEXT.feedbackPrivate.sendFeedback(feedback_ffi, loadSystemInfo === A.H.TRUE, formOpenTime);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
