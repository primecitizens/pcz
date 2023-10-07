import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/history", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_DeleteRangeArgRange": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Float64(ptr + 0, x["endTime"] === undefined ? 0 : (x["endTime"] as number));
        A.store.Float64(ptr + 8, x["startTime"] === undefined ? 0 : (x["startTime"] as number));
      }
    },
    "load_DeleteRangeArgRange": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["endTime"] = A.load.Float64(ptr + 0);
      x["startTime"] = A.load.Float64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_HistoryItem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 51, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 48, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 49, false);
        A.store.Int64(ptr + 24, 0);
        A.store.Ref(ptr + 32, undefined);
        A.store.Bool(ptr + 50, false);
        A.store.Int64(ptr + 40, 0);
      } else {
        A.store.Bool(ptr + 51, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Bool(ptr + 48, "lastVisitTime" in x ? true : false);
        A.store.Float64(ptr + 8, x["lastVisitTime"] === undefined ? 0 : (x["lastVisitTime"] as number));
        A.store.Ref(ptr + 16, x["title"]);
        A.store.Bool(ptr + 49, "typedCount" in x ? true : false);
        A.store.Int64(ptr + 24, x["typedCount"] === undefined ? 0 : (x["typedCount"] as number));
        A.store.Ref(ptr + 32, x["url"]);
        A.store.Bool(ptr + 50, "visitCount" in x ? true : false);
        A.store.Int64(ptr + 40, x["visitCount"] === undefined ? 0 : (x["visitCount"] as number));
      }
    },
    "load_HistoryItem": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 48)) {
        x["lastVisitTime"] = A.load.Float64(ptr + 8);
      } else {
        delete x["lastVisitTime"];
      }
      x["title"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 49)) {
        x["typedCount"] = A.load.Int64(ptr + 24);
      } else {
        delete x["typedCount"];
      }
      x["url"] = A.load.Ref(ptr + 32, undefined);
      if (A.load.Bool(ptr + 50)) {
        x["visitCount"] = A.load.Int64(ptr + 40);
      } else {
        delete x["visitCount"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnVisitRemovedArgRemoved": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Bool(ptr + 0, x["allHistory"] ? true : false);
        A.store.Ref(ptr + 4, x["urls"]);
      }
    },
    "load_OnVisitRemovedArgRemoved": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["allHistory"] = A.load.Bool(ptr + 0);
      x["urls"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SearchArgQuery": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 31, false);
        A.store.Bool(ptr + 28, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 29, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Bool(ptr + 30, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 31, true);
        A.store.Bool(ptr + 28, "endTime" in x ? true : false);
        A.store.Float64(ptr + 0, x["endTime"] === undefined ? 0 : (x["endTime"] as number));
        A.store.Bool(ptr + 29, "maxResults" in x ? true : false);
        A.store.Int64(ptr + 8, x["maxResults"] === undefined ? 0 : (x["maxResults"] as number));
        A.store.Bool(ptr + 30, "startTime" in x ? true : false);
        A.store.Float64(ptr + 16, x["startTime"] === undefined ? 0 : (x["startTime"] as number));
        A.store.Ref(ptr + 24, x["text"]);
      }
    },
    "load_SearchArgQuery": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 28)) {
        x["endTime"] = A.load.Float64(ptr + 0);
      } else {
        delete x["endTime"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["maxResults"] = A.load.Int64(ptr + 8);
      } else {
        delete x["maxResults"];
      }
      if (A.load.Bool(ptr + 30)) {
        x["startTime"] = A.load.Float64(ptr + 16);
      } else {
        delete x["startTime"];
      }
      x["text"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_TransitionType": (ref: heap.Ref<string>): number => {
      const idx = [
        "link",
        "typed",
        "auto_bookmark",
        "auto_subframe",
        "manual_subframe",
        "generated",
        "auto_toplevel",
        "form_submit",
        "reload",
        "keyword",
        "keyword_generated",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_UrlDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["url"]);
      }
    },
    "load_UrlDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_VisitItem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 33, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 8, undefined);
        A.store.Enum(ptr + 12, -1);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 32, false);
        A.store.Float64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 33, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Bool(ptr + 4, x["isLocal"] ? true : false);
        A.store.Ref(ptr + 8, x["referringVisitId"]);
        A.store.Enum(
          ptr + 12,
          [
            "link",
            "typed",
            "auto_bookmark",
            "auto_subframe",
            "manual_subframe",
            "generated",
            "auto_toplevel",
            "form_submit",
            "reload",
            "keyword",
            "keyword_generated",
          ].indexOf(x["transition"] as string)
        );
        A.store.Ref(ptr + 16, x["visitId"]);
        A.store.Bool(ptr + 32, "visitTime" in x ? true : false);
        A.store.Float64(ptr + 24, x["visitTime"] === undefined ? 0 : (x["visitTime"] as number));
      }
    },
    "load_VisitItem": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["isLocal"] = A.load.Bool(ptr + 4);
      x["referringVisitId"] = A.load.Ref(ptr + 8, undefined);
      x["transition"] = A.load.Enum(ptr + 12, [
        "link",
        "typed",
        "auto_bookmark",
        "auto_subframe",
        "manual_subframe",
        "generated",
        "auto_toplevel",
        "form_submit",
        "reload",
        "keyword",
        "keyword_generated",
      ]);
      x["visitId"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 32)) {
        x["visitTime"] = A.load.Float64(ptr + 24);
      } else {
        delete x["visitTime"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_AddUrl": (): heap.Ref<boolean> => {
      if (WEBEXT?.history && "addUrl" in WEBEXT?.history) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddUrl": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.history.addUrl);
    },
    "call_AddUrl": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["url"] = A.load.Ref(details + 0, undefined);

      const _ret = WEBEXT.history.addUrl(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_AddUrl": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["url"] = A.load.Ref(details + 0, undefined);

        const _ret = WEBEXT.history.addUrl(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DeleteAll": (): heap.Ref<boolean> => {
      if (WEBEXT?.history && "deleteAll" in WEBEXT?.history) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DeleteAll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.history.deleteAll);
    },
    "call_DeleteAll": (retPtr: Pointer): void => {
      const _ret = WEBEXT.history.deleteAll();
      A.store.Ref(retPtr, _ret);
    },
    "try_DeleteAll": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.history.deleteAll();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DeleteRange": (): heap.Ref<boolean> => {
      if (WEBEXT?.history && "deleteRange" in WEBEXT?.history) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DeleteRange": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.history.deleteRange);
    },
    "call_DeleteRange": (retPtr: Pointer, range: Pointer): void => {
      const range_ffi = {};

      range_ffi["endTime"] = A.load.Float64(range + 0);
      range_ffi["startTime"] = A.load.Float64(range + 8);

      const _ret = WEBEXT.history.deleteRange(range_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_DeleteRange": (retPtr: Pointer, errPtr: Pointer, range: Pointer): heap.Ref<boolean> => {
      try {
        const range_ffi = {};

        range_ffi["endTime"] = A.load.Float64(range + 0);
        range_ffi["startTime"] = A.load.Float64(range + 8);

        const _ret = WEBEXT.history.deleteRange(range_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DeleteUrl": (): heap.Ref<boolean> => {
      if (WEBEXT?.history && "deleteUrl" in WEBEXT?.history) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DeleteUrl": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.history.deleteUrl);
    },
    "call_DeleteUrl": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["url"] = A.load.Ref(details + 0, undefined);

      const _ret = WEBEXT.history.deleteUrl(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_DeleteUrl": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["url"] = A.load.Ref(details + 0, undefined);

        const _ret = WEBEXT.history.deleteUrl(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetVisits": (): heap.Ref<boolean> => {
      if (WEBEXT?.history && "getVisits" in WEBEXT?.history) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetVisits": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.history.getVisits);
    },
    "call_GetVisits": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["url"] = A.load.Ref(details + 0, undefined);

      const _ret = WEBEXT.history.getVisits(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetVisits": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["url"] = A.load.Ref(details + 0, undefined);

        const _ret = WEBEXT.history.getVisits(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnVisitRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.history?.onVisitRemoved && "addListener" in WEBEXT?.history?.onVisitRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnVisitRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.history.onVisitRemoved.addListener);
    },
    "call_OnVisitRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.history.onVisitRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnVisitRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.history.onVisitRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffVisitRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.history?.onVisitRemoved && "removeListener" in WEBEXT?.history?.onVisitRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffVisitRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.history.onVisitRemoved.removeListener);
    },
    "call_OffVisitRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.history.onVisitRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffVisitRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.history.onVisitRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnVisitRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.history?.onVisitRemoved && "hasListener" in WEBEXT?.history?.onVisitRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnVisitRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.history.onVisitRemoved.hasListener);
    },
    "call_HasOnVisitRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.history.onVisitRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnVisitRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.history.onVisitRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnVisited": (): heap.Ref<boolean> => {
      if (WEBEXT?.history?.onVisited && "addListener" in WEBEXT?.history?.onVisited) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnVisited": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.history.onVisited.addListener);
    },
    "call_OnVisited": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.history.onVisited.addListener(A.H.get<object>(callback));
    },
    "try_OnVisited": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.history.onVisited.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffVisited": (): heap.Ref<boolean> => {
      if (WEBEXT?.history?.onVisited && "removeListener" in WEBEXT?.history?.onVisited) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffVisited": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.history.onVisited.removeListener);
    },
    "call_OffVisited": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.history.onVisited.removeListener(A.H.get<object>(callback));
    },
    "try_OffVisited": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.history.onVisited.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnVisited": (): heap.Ref<boolean> => {
      if (WEBEXT?.history?.onVisited && "hasListener" in WEBEXT?.history?.onVisited) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnVisited": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.history.onVisited.hasListener);
    },
    "call_HasOnVisited": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.history.onVisited.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnVisited": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.history.onVisited.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Search": (): heap.Ref<boolean> => {
      if (WEBEXT?.history && "search" in WEBEXT?.history) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Search": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.history.search);
    },
    "call_Search": (retPtr: Pointer, query: Pointer): void => {
      const query_ffi = {};

      if (A.load.Bool(query + 28)) {
        query_ffi["endTime"] = A.load.Float64(query + 0);
      }
      if (A.load.Bool(query + 29)) {
        query_ffi["maxResults"] = A.load.Int64(query + 8);
      }
      if (A.load.Bool(query + 30)) {
        query_ffi["startTime"] = A.load.Float64(query + 16);
      }
      query_ffi["text"] = A.load.Ref(query + 24, undefined);

      const _ret = WEBEXT.history.search(query_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Search": (retPtr: Pointer, errPtr: Pointer, query: Pointer): heap.Ref<boolean> => {
      try {
        const query_ffi = {};

        if (A.load.Bool(query + 28)) {
          query_ffi["endTime"] = A.load.Float64(query + 0);
        }
        if (A.load.Bool(query + 29)) {
          query_ffi["maxResults"] = A.load.Int64(query + 8);
        }
        if (A.load.Bool(query + 30)) {
          query_ffi["startTime"] = A.load.Float64(query + 16);
        }
        query_ffi["text"] = A.load.Ref(query + 24, undefined);

        const _ret = WEBEXT.history.search(query_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
