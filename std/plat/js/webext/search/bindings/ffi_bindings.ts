import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/search", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_Disposition": (ref: heap.Ref<string>): number => {
      const idx = ["CURRENT_TAB", "NEW_TAB", "NEW_WINDOW"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_QueryInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Ref(ptr + 0, x["text"]);
        A.store.Enum(ptr + 4, ["CURRENT_TAB", "NEW_TAB", "NEW_WINDOW"].indexOf(x["disposition"] as string));
        A.store.Bool(ptr + 12, "tabId" in x ? true : false);
        A.store.Int32(ptr + 8, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
      }
    },
    "load_QueryInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["text"] = A.load.Ref(ptr + 0, undefined);
      x["disposition"] = A.load.Enum(ptr + 4, ["CURRENT_TAB", "NEW_TAB", "NEW_WINDOW"]);
      if (A.load.Bool(ptr + 12)) {
        x["tabId"] = A.load.Int32(ptr + 8);
      } else {
        delete x["tabId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Query": (): heap.Ref<boolean> => {
      if (WEBEXT?.search && "query" in WEBEXT?.search) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Query": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.search.query);
    },
    "call_Query": (retPtr: Pointer, queryInfo: Pointer): void => {
      const queryInfo_ffi = {};

      queryInfo_ffi["text"] = A.load.Ref(queryInfo + 0, undefined);
      queryInfo_ffi["disposition"] = A.load.Enum(queryInfo + 4, ["CURRENT_TAB", "NEW_TAB", "NEW_WINDOW"]);
      if (A.load.Bool(queryInfo + 12)) {
        queryInfo_ffi["tabId"] = A.load.Int32(queryInfo + 8);
      }

      const _ret = WEBEXT.search.query(queryInfo_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Query": (retPtr: Pointer, errPtr: Pointer, queryInfo: Pointer): heap.Ref<boolean> => {
      try {
        const queryInfo_ffi = {};

        queryInfo_ffi["text"] = A.load.Ref(queryInfo + 0, undefined);
        queryInfo_ffi["disposition"] = A.load.Enum(queryInfo + 4, ["CURRENT_TAB", "NEW_TAB", "NEW_WINDOW"]);
        if (A.load.Bool(queryInfo + 12)) {
          queryInfo_ffi["tabId"] = A.load.Int32(queryInfo + 8);
        }

        const _ret = WEBEXT.search.query(queryInfo_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
