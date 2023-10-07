import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/activitylogprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_ExtensionActivityType": (ref: heap.Ref<string>): number => {
      const idx = ["api_call", "api_event", "content_script", "dom_access", "dom_event", "web_request"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ExtensionActivityDomVerb": (ref: heap.Ref<string>): number => {
      const idx = ["getter", "setter", "method", "inserted", "xhr", "webrequest", "modified"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ExtensionActivityFieldOther": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Enum(
          ptr + 0,
          ["getter", "setter", "method", "inserted", "xhr", "webrequest", "modified"].indexOf(x["domVerb"] as string)
        );
        A.store.Ref(ptr + 4, x["extra"]);
        A.store.Bool(ptr + 16, "prerender" in x ? true : false);
        A.store.Bool(ptr + 8, x["prerender"] ? true : false);
        A.store.Ref(ptr + 12, x["webRequest"]);
      }
    },
    "load_ExtensionActivityFieldOther": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["domVerb"] = A.load.Enum(ptr + 0, ["getter", "setter", "method", "inserted", "xhr", "webrequest", "modified"]);
      x["extra"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["prerender"] = A.load.Bool(ptr + 8);
      } else {
        delete x["prerender"];
      }
      x["webRequest"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ExtensionActivity": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 74, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 72, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Ref(ptr + 32, undefined);

        A.store.Bool(ptr + 36 + 17, false);
        A.store.Enum(ptr + 36 + 0, -1);
        A.store.Ref(ptr + 36 + 4, undefined);
        A.store.Bool(ptr + 36 + 16, false);
        A.store.Bool(ptr + 36 + 8, false);
        A.store.Ref(ptr + 36 + 12, undefined);
        A.store.Ref(ptr + 56, undefined);
        A.store.Ref(ptr + 60, undefined);
        A.store.Bool(ptr + 73, false);
        A.store.Float64(ptr + 64, 0);
      } else {
        A.store.Bool(ptr + 74, true);
        A.store.Ref(ptr + 0, x["activityId"]);
        A.store.Enum(
          ptr + 4,
          ["api_call", "api_event", "content_script", "dom_access", "dom_event", "web_request"].indexOf(
            x["activityType"] as string
          )
        );
        A.store.Ref(ptr + 8, x["apiCall"]);
        A.store.Ref(ptr + 12, x["argUrl"]);
        A.store.Ref(ptr + 16, x["args"]);
        A.store.Bool(ptr + 72, "count" in x ? true : false);
        A.store.Float64(ptr + 24, x["count"] === undefined ? 0 : (x["count"] as number));
        A.store.Ref(ptr + 32, x["extensionId"]);

        if (typeof x["other"] === "undefined") {
          A.store.Bool(ptr + 36 + 17, false);
          A.store.Enum(ptr + 36 + 0, -1);
          A.store.Ref(ptr + 36 + 4, undefined);
          A.store.Bool(ptr + 36 + 16, false);
          A.store.Bool(ptr + 36 + 8, false);
          A.store.Ref(ptr + 36 + 12, undefined);
        } else {
          A.store.Bool(ptr + 36 + 17, true);
          A.store.Enum(
            ptr + 36 + 0,
            ["getter", "setter", "method", "inserted", "xhr", "webrequest", "modified"].indexOf(
              x["other"]["domVerb"] as string
            )
          );
          A.store.Ref(ptr + 36 + 4, x["other"]["extra"]);
          A.store.Bool(ptr + 36 + 16, "prerender" in x["other"] ? true : false);
          A.store.Bool(ptr + 36 + 8, x["other"]["prerender"] ? true : false);
          A.store.Ref(ptr + 36 + 12, x["other"]["webRequest"]);
        }
        A.store.Ref(ptr + 56, x["pageTitle"]);
        A.store.Ref(ptr + 60, x["pageUrl"]);
        A.store.Bool(ptr + 73, "time" in x ? true : false);
        A.store.Float64(ptr + 64, x["time"] === undefined ? 0 : (x["time"] as number));
      }
    },
    "load_ExtensionActivity": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["activityId"] = A.load.Ref(ptr + 0, undefined);
      x["activityType"] = A.load.Enum(ptr + 4, [
        "api_call",
        "api_event",
        "content_script",
        "dom_access",
        "dom_event",
        "web_request",
      ]);
      x["apiCall"] = A.load.Ref(ptr + 8, undefined);
      x["argUrl"] = A.load.Ref(ptr + 12, undefined);
      x["args"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 72)) {
        x["count"] = A.load.Float64(ptr + 24);
      } else {
        delete x["count"];
      }
      x["extensionId"] = A.load.Ref(ptr + 32, undefined);
      if (A.load.Bool(ptr + 36 + 17)) {
        x["other"] = {};
        x["other"]["domVerb"] = A.load.Enum(ptr + 36 + 0, [
          "getter",
          "setter",
          "method",
          "inserted",
          "xhr",
          "webrequest",
          "modified",
        ]);
        x["other"]["extra"] = A.load.Ref(ptr + 36 + 4, undefined);
        if (A.load.Bool(ptr + 36 + 16)) {
          x["other"]["prerender"] = A.load.Bool(ptr + 36 + 8);
        } else {
          delete x["other"]["prerender"];
        }
        x["other"]["webRequest"] = A.load.Ref(ptr + 36 + 12, undefined);
      } else {
        delete x["other"];
      }
      x["pageTitle"] = A.load.Ref(ptr + 56, undefined);
      x["pageUrl"] = A.load.Ref(ptr + 60, undefined);
      if (A.load.Bool(ptr + 73)) {
        x["time"] = A.load.Float64(ptr + 64);
      } else {
        delete x["time"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ActivityResultSet": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["activities"]);
      }
    },
    "load_ActivityResultSet": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["activities"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ExtensionActivityFilter": (ref: heap.Ref<string>): number => {
      const idx = ["api_call", "api_event", "content_script", "dom_access", "dom_event", "web_request", "any"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Filter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 33, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 32, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
      } else {
        A.store.Bool(ptr + 33, true);
        A.store.Enum(
          ptr + 0,
          ["api_call", "api_event", "content_script", "dom_access", "dom_event", "web_request", "any"].indexOf(
            x["activityType"] as string
          )
        );
        A.store.Ref(ptr + 4, x["apiCall"]);
        A.store.Ref(ptr + 8, x["argUrl"]);
        A.store.Bool(ptr + 32, "daysAgo" in x ? true : false);
        A.store.Int64(ptr + 16, x["daysAgo"] === undefined ? 0 : (x["daysAgo"] as number));
        A.store.Ref(ptr + 24, x["extensionId"]);
        A.store.Ref(ptr + 28, x["pageUrl"]);
      }
    },
    "load_Filter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["activityType"] = A.load.Enum(ptr + 0, [
        "api_call",
        "api_event",
        "content_script",
        "dom_access",
        "dom_event",
        "web_request",
        "any",
      ]);
      x["apiCall"] = A.load.Ref(ptr + 4, undefined);
      x["argUrl"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 32)) {
        x["daysAgo"] = A.load.Int64(ptr + 16);
      } else {
        delete x["daysAgo"];
      }
      x["extensionId"] = A.load.Ref(ptr + 24, undefined);
      x["pageUrl"] = A.load.Ref(ptr + 28, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_DeleteActivities": (): heap.Ref<boolean> => {
      if (WEBEXT?.activityLogPrivate && "deleteActivities" in WEBEXT?.activityLogPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DeleteActivities": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.activityLogPrivate.deleteActivities);
    },
    "call_DeleteActivities": (retPtr: Pointer, activityIds: heap.Ref<object>): void => {
      const _ret = WEBEXT.activityLogPrivate.deleteActivities(A.H.get<object>(activityIds));
      A.store.Ref(retPtr, _ret);
    },
    "try_DeleteActivities": (retPtr: Pointer, errPtr: Pointer, activityIds: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.activityLogPrivate.deleteActivities(A.H.get<object>(activityIds));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DeleteActivitiesByExtension": (): heap.Ref<boolean> => {
      if (WEBEXT?.activityLogPrivate && "deleteActivitiesByExtension" in WEBEXT?.activityLogPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DeleteActivitiesByExtension": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.activityLogPrivate.deleteActivitiesByExtension);
    },
    "call_DeleteActivitiesByExtension": (retPtr: Pointer, extensionId: heap.Ref<object>): void => {
      const _ret = WEBEXT.activityLogPrivate.deleteActivitiesByExtension(A.H.get<object>(extensionId));
      A.store.Ref(retPtr, _ret);
    },
    "try_DeleteActivitiesByExtension": (
      retPtr: Pointer,
      errPtr: Pointer,
      extensionId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.activityLogPrivate.deleteActivitiesByExtension(A.H.get<object>(extensionId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DeleteDatabase": (): heap.Ref<boolean> => {
      if (WEBEXT?.activityLogPrivate && "deleteDatabase" in WEBEXT?.activityLogPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DeleteDatabase": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.activityLogPrivate.deleteDatabase);
    },
    "call_DeleteDatabase": (retPtr: Pointer): void => {
      const _ret = WEBEXT.activityLogPrivate.deleteDatabase();
    },
    "try_DeleteDatabase": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.activityLogPrivate.deleteDatabase();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DeleteUrls": (): heap.Ref<boolean> => {
      if (WEBEXT?.activityLogPrivate && "deleteUrls" in WEBEXT?.activityLogPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DeleteUrls": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.activityLogPrivate.deleteUrls);
    },
    "call_DeleteUrls": (retPtr: Pointer, urls: heap.Ref<object>): void => {
      const _ret = WEBEXT.activityLogPrivate.deleteUrls(A.H.get<object>(urls));
    },
    "try_DeleteUrls": (retPtr: Pointer, errPtr: Pointer, urls: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.activityLogPrivate.deleteUrls(A.H.get<object>(urls));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetExtensionActivities": (): heap.Ref<boolean> => {
      if (WEBEXT?.activityLogPrivate && "getExtensionActivities" in WEBEXT?.activityLogPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetExtensionActivities": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.activityLogPrivate.getExtensionActivities);
    },
    "call_GetExtensionActivities": (retPtr: Pointer, filter: Pointer): void => {
      const filter_ffi = {};

      filter_ffi["activityType"] = A.load.Enum(filter + 0, [
        "api_call",
        "api_event",
        "content_script",
        "dom_access",
        "dom_event",
        "web_request",
        "any",
      ]);
      filter_ffi["apiCall"] = A.load.Ref(filter + 4, undefined);
      filter_ffi["argUrl"] = A.load.Ref(filter + 8, undefined);
      if (A.load.Bool(filter + 32)) {
        filter_ffi["daysAgo"] = A.load.Int64(filter + 16);
      }
      filter_ffi["extensionId"] = A.load.Ref(filter + 24, undefined);
      filter_ffi["pageUrl"] = A.load.Ref(filter + 28, undefined);

      const _ret = WEBEXT.activityLogPrivate.getExtensionActivities(filter_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetExtensionActivities": (retPtr: Pointer, errPtr: Pointer, filter: Pointer): heap.Ref<boolean> => {
      try {
        const filter_ffi = {};

        filter_ffi["activityType"] = A.load.Enum(filter + 0, [
          "api_call",
          "api_event",
          "content_script",
          "dom_access",
          "dom_event",
          "web_request",
          "any",
        ]);
        filter_ffi["apiCall"] = A.load.Ref(filter + 4, undefined);
        filter_ffi["argUrl"] = A.load.Ref(filter + 8, undefined);
        if (A.load.Bool(filter + 32)) {
          filter_ffi["daysAgo"] = A.load.Int64(filter + 16);
        }
        filter_ffi["extensionId"] = A.load.Ref(filter + 24, undefined);
        filter_ffi["pageUrl"] = A.load.Ref(filter + 28, undefined);

        const _ret = WEBEXT.activityLogPrivate.getExtensionActivities(filter_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnExtensionActivity": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.activityLogPrivate?.onExtensionActivity &&
        "addListener" in WEBEXT?.activityLogPrivate?.onExtensionActivity
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnExtensionActivity": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.activityLogPrivate.onExtensionActivity.addListener);
    },
    "call_OnExtensionActivity": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.activityLogPrivate.onExtensionActivity.addListener(A.H.get<object>(callback));
    },
    "try_OnExtensionActivity": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.activityLogPrivate.onExtensionActivity.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffExtensionActivity": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.activityLogPrivate?.onExtensionActivity &&
        "removeListener" in WEBEXT?.activityLogPrivate?.onExtensionActivity
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffExtensionActivity": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.activityLogPrivate.onExtensionActivity.removeListener);
    },
    "call_OffExtensionActivity": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.activityLogPrivate.onExtensionActivity.removeListener(A.H.get<object>(callback));
    },
    "try_OffExtensionActivity": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.activityLogPrivate.onExtensionActivity.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnExtensionActivity": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.activityLogPrivate?.onExtensionActivity &&
        "hasListener" in WEBEXT?.activityLogPrivate?.onExtensionActivity
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnExtensionActivity": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.activityLogPrivate.onExtensionActivity.hasListener);
    },
    "call_HasOnExtensionActivity": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.activityLogPrivate.onExtensionActivity.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnExtensionActivity": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.activityLogPrivate.onExtensionActivity.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
