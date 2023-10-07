import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/readinglist", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AddEntryOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Ref(ptr + 4, x["title"]);
        A.store.Bool(ptr + 9, "hasBeenRead" in x ? true : false);
        A.store.Bool(ptr + 8, x["hasBeenRead"] ? true : false);
      }
    },
    "load_AddEntryOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      x["title"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 9)) {
        x["hasBeenRead"] = A.load.Bool(ptr + 8);
      } else {
        delete x["hasBeenRead"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ReadingListEntry": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 35, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 33, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 34, false);
        A.store.Float64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 35, true);
        A.store.Ref(ptr + 0, x["url"]);
        A.store.Ref(ptr + 4, x["title"]);
        A.store.Bool(ptr + 32, "hasBeenRead" in x ? true : false);
        A.store.Bool(ptr + 8, x["hasBeenRead"] ? true : false);
        A.store.Bool(ptr + 33, "lastUpdateTime" in x ? true : false);
        A.store.Float64(ptr + 16, x["lastUpdateTime"] === undefined ? 0 : (x["lastUpdateTime"] as number));
        A.store.Bool(ptr + 34, "creationTime" in x ? true : false);
        A.store.Float64(ptr + 24, x["creationTime"] === undefined ? 0 : (x["creationTime"] as number));
      }
    },
    "load_ReadingListEntry": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      x["title"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 32)) {
        x["hasBeenRead"] = A.load.Bool(ptr + 8);
      } else {
        delete x["hasBeenRead"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["lastUpdateTime"] = A.load.Float64(ptr + 16);
      } else {
        delete x["lastUpdateTime"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["creationTime"] = A.load.Float64(ptr + 24);
      } else {
        delete x["creationTime"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_QueryInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Ref(ptr + 4, x["title"]);
        A.store.Bool(ptr + 9, "hasBeenRead" in x ? true : false);
        A.store.Bool(ptr + 8, x["hasBeenRead"] ? true : false);
      }
    },
    "load_QueryInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      x["title"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 9)) {
        x["hasBeenRead"] = A.load.Bool(ptr + 8);
      } else {
        delete x["hasBeenRead"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RemoveOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["url"]);
      }
    },
    "load_RemoveOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UpdateEntryOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Ref(ptr + 4, x["title"]);
        A.store.Bool(ptr + 9, "hasBeenRead" in x ? true : false);
        A.store.Bool(ptr + 8, x["hasBeenRead"] ? true : false);
      }
    },
    "load_UpdateEntryOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      x["title"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 9)) {
        x["hasBeenRead"] = A.load.Bool(ptr + 8);
      } else {
        delete x["hasBeenRead"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_AddEntry": (): heap.Ref<boolean> => {
      if (WEBEXT?.readingList && "addEntry" in WEBEXT?.readingList) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddEntry": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.readingList.addEntry);
    },
    "call_AddEntry": (retPtr: Pointer, entry: Pointer): void => {
      const entry_ffi = {};

      entry_ffi["url"] = A.load.Ref(entry + 0, undefined);
      entry_ffi["title"] = A.load.Ref(entry + 4, undefined);
      if (A.load.Bool(entry + 9)) {
        entry_ffi["hasBeenRead"] = A.load.Bool(entry + 8);
      }

      const _ret = WEBEXT.readingList.addEntry(entry_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_AddEntry": (retPtr: Pointer, errPtr: Pointer, entry: Pointer): heap.Ref<boolean> => {
      try {
        const entry_ffi = {};

        entry_ffi["url"] = A.load.Ref(entry + 0, undefined);
        entry_ffi["title"] = A.load.Ref(entry + 4, undefined);
        if (A.load.Bool(entry + 9)) {
          entry_ffi["hasBeenRead"] = A.load.Bool(entry + 8);
        }

        const _ret = WEBEXT.readingList.addEntry(entry_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnEntryAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.readingList?.onEntryAdded && "addListener" in WEBEXT?.readingList?.onEntryAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnEntryAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.readingList.onEntryAdded.addListener);
    },
    "call_OnEntryAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.readingList.onEntryAdded.addListener(A.H.get<object>(callback));
    },
    "try_OnEntryAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.readingList.onEntryAdded.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffEntryAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.readingList?.onEntryAdded && "removeListener" in WEBEXT?.readingList?.onEntryAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffEntryAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.readingList.onEntryAdded.removeListener);
    },
    "call_OffEntryAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.readingList.onEntryAdded.removeListener(A.H.get<object>(callback));
    },
    "try_OffEntryAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.readingList.onEntryAdded.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnEntryAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.readingList?.onEntryAdded && "hasListener" in WEBEXT?.readingList?.onEntryAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnEntryAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.readingList.onEntryAdded.hasListener);
    },
    "call_HasOnEntryAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.readingList.onEntryAdded.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnEntryAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.readingList.onEntryAdded.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnEntryRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.readingList?.onEntryRemoved && "addListener" in WEBEXT?.readingList?.onEntryRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnEntryRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.readingList.onEntryRemoved.addListener);
    },
    "call_OnEntryRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.readingList.onEntryRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnEntryRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.readingList.onEntryRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffEntryRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.readingList?.onEntryRemoved && "removeListener" in WEBEXT?.readingList?.onEntryRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffEntryRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.readingList.onEntryRemoved.removeListener);
    },
    "call_OffEntryRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.readingList.onEntryRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffEntryRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.readingList.onEntryRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnEntryRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.readingList?.onEntryRemoved && "hasListener" in WEBEXT?.readingList?.onEntryRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnEntryRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.readingList.onEntryRemoved.hasListener);
    },
    "call_HasOnEntryRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.readingList.onEntryRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnEntryRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.readingList.onEntryRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnEntryUpdated": (): heap.Ref<boolean> => {
      if (WEBEXT?.readingList?.onEntryUpdated && "addListener" in WEBEXT?.readingList?.onEntryUpdated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnEntryUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.readingList.onEntryUpdated.addListener);
    },
    "call_OnEntryUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.readingList.onEntryUpdated.addListener(A.H.get<object>(callback));
    },
    "try_OnEntryUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.readingList.onEntryUpdated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffEntryUpdated": (): heap.Ref<boolean> => {
      if (WEBEXT?.readingList?.onEntryUpdated && "removeListener" in WEBEXT?.readingList?.onEntryUpdated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffEntryUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.readingList.onEntryUpdated.removeListener);
    },
    "call_OffEntryUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.readingList.onEntryUpdated.removeListener(A.H.get<object>(callback));
    },
    "try_OffEntryUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.readingList.onEntryUpdated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnEntryUpdated": (): heap.Ref<boolean> => {
      if (WEBEXT?.readingList?.onEntryUpdated && "hasListener" in WEBEXT?.readingList?.onEntryUpdated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnEntryUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.readingList.onEntryUpdated.hasListener);
    },
    "call_HasOnEntryUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.readingList.onEntryUpdated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnEntryUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.readingList.onEntryUpdated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Query": (): heap.Ref<boolean> => {
      if (WEBEXT?.readingList && "query" in WEBEXT?.readingList) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Query": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.readingList.query);
    },
    "call_Query": (retPtr: Pointer, info: Pointer): void => {
      const info_ffi = {};

      info_ffi["url"] = A.load.Ref(info + 0, undefined);
      info_ffi["title"] = A.load.Ref(info + 4, undefined);
      if (A.load.Bool(info + 9)) {
        info_ffi["hasBeenRead"] = A.load.Bool(info + 8);
      }

      const _ret = WEBEXT.readingList.query(info_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Query": (retPtr: Pointer, errPtr: Pointer, info: Pointer): heap.Ref<boolean> => {
      try {
        const info_ffi = {};

        info_ffi["url"] = A.load.Ref(info + 0, undefined);
        info_ffi["title"] = A.load.Ref(info + 4, undefined);
        if (A.load.Bool(info + 9)) {
          info_ffi["hasBeenRead"] = A.load.Bool(info + 8);
        }

        const _ret = WEBEXT.readingList.query(info_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveEntry": (): heap.Ref<boolean> => {
      if (WEBEXT?.readingList && "removeEntry" in WEBEXT?.readingList) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveEntry": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.readingList.removeEntry);
    },
    "call_RemoveEntry": (retPtr: Pointer, info: Pointer): void => {
      const info_ffi = {};

      info_ffi["url"] = A.load.Ref(info + 0, undefined);

      const _ret = WEBEXT.readingList.removeEntry(info_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveEntry": (retPtr: Pointer, errPtr: Pointer, info: Pointer): heap.Ref<boolean> => {
      try {
        const info_ffi = {};

        info_ffi["url"] = A.load.Ref(info + 0, undefined);

        const _ret = WEBEXT.readingList.removeEntry(info_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateEntry": (): heap.Ref<boolean> => {
      if (WEBEXT?.readingList && "updateEntry" in WEBEXT?.readingList) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateEntry": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.readingList.updateEntry);
    },
    "call_UpdateEntry": (retPtr: Pointer, info: Pointer): void => {
      const info_ffi = {};

      info_ffi["url"] = A.load.Ref(info + 0, undefined);
      info_ffi["title"] = A.load.Ref(info + 4, undefined);
      if (A.load.Bool(info + 9)) {
        info_ffi["hasBeenRead"] = A.load.Bool(info + 8);
      }

      const _ret = WEBEXT.readingList.updateEntry(info_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_UpdateEntry": (retPtr: Pointer, errPtr: Pointer, info: Pointer): heap.Ref<boolean> => {
      try {
        const info_ffi = {};

        info_ffi["url"] = A.load.Ref(info + 0, undefined);
        info_ffi["title"] = A.load.Ref(info + 4, undefined);
        if (A.load.Bool(info + 9)) {
          info_ffi["hasBeenRead"] = A.load.Bool(info + 8);
        }

        const _ret = WEBEXT.readingList.updateEntry(info_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
