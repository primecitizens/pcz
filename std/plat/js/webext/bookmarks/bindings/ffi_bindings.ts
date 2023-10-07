import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/bookmarks", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_BookmarkTreeNodeUnmodifiable": (ref: heap.Ref<string>): number => {
      const idx = ["managed"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_BookmarkTreeNode": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 68, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 64, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 65, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 66, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Ref(ptr + 32, undefined);
        A.store.Bool(ptr + 67, false);
        A.store.Int64(ptr + 40, 0);
        A.store.Ref(ptr + 48, undefined);
        A.store.Ref(ptr + 52, undefined);
        A.store.Enum(ptr + 56, -1);
        A.store.Ref(ptr + 60, undefined);
      } else {
        A.store.Bool(ptr + 68, true);
        A.store.Ref(ptr + 0, x["children"]);
        A.store.Bool(ptr + 64, "dateAdded" in x ? true : false);
        A.store.Float64(ptr + 8, x["dateAdded"] === undefined ? 0 : (x["dateAdded"] as number));
        A.store.Bool(ptr + 65, "dateGroupModified" in x ? true : false);
        A.store.Float64(ptr + 16, x["dateGroupModified"] === undefined ? 0 : (x["dateGroupModified"] as number));
        A.store.Bool(ptr + 66, "dateLastUsed" in x ? true : false);
        A.store.Float64(ptr + 24, x["dateLastUsed"] === undefined ? 0 : (x["dateLastUsed"] as number));
        A.store.Ref(ptr + 32, x["id"]);
        A.store.Bool(ptr + 67, "index" in x ? true : false);
        A.store.Int64(ptr + 40, x["index"] === undefined ? 0 : (x["index"] as number));
        A.store.Ref(ptr + 48, x["parentId"]);
        A.store.Ref(ptr + 52, x["title"]);
        A.store.Enum(ptr + 56, ["managed"].indexOf(x["unmodifiable"] as string));
        A.store.Ref(ptr + 60, x["url"]);
      }
    },
    "load_BookmarkTreeNode": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["children"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 64)) {
        x["dateAdded"] = A.load.Float64(ptr + 8);
      } else {
        delete x["dateAdded"];
      }
      if (A.load.Bool(ptr + 65)) {
        x["dateGroupModified"] = A.load.Float64(ptr + 16);
      } else {
        delete x["dateGroupModified"];
      }
      if (A.load.Bool(ptr + 66)) {
        x["dateLastUsed"] = A.load.Float64(ptr + 24);
      } else {
        delete x["dateLastUsed"];
      }
      x["id"] = A.load.Ref(ptr + 32, undefined);
      if (A.load.Bool(ptr + 67)) {
        x["index"] = A.load.Int64(ptr + 40);
      } else {
        delete x["index"];
      }
      x["parentId"] = A.load.Ref(ptr + 48, undefined);
      x["title"] = A.load.Ref(ptr + 52, undefined);
      x["unmodifiable"] = A.load.Enum(ptr + 56, ["managed"]);
      x["url"] = A.load.Ref(ptr + 60, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CreateDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 20, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Bool(ptr + 20, "index" in x ? true : false);
        A.store.Int64(ptr + 0, x["index"] === undefined ? 0 : (x["index"] as number));
        A.store.Ref(ptr + 8, x["parentId"]);
        A.store.Ref(ptr + 12, x["title"]);
        A.store.Ref(ptr + 16, x["url"]);
      }
    },
    "load_CreateDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 20)) {
        x["index"] = A.load.Int64(ptr + 0);
      } else {
        delete x["index"];
      }
      x["parentId"] = A.load.Ref(ptr + 8, undefined);
      x["title"] = A.load.Ref(ptr + 12, undefined);
      x["url"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "get_MAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks && "MAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE" in WEBEXT?.bookmarks) {
        const val = WEBEXT.bookmarks.MAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_MAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.bookmarks,
        "MAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE",
        A.H.get<object>(val),
        WEBEXT.bookmarks
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_MAX_WRITE_OPERATIONS_PER_HOUR": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks && "MAX_WRITE_OPERATIONS_PER_HOUR" in WEBEXT?.bookmarks) {
        const val = WEBEXT.bookmarks.MAX_WRITE_OPERATIONS_PER_HOUR;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_MAX_WRITE_OPERATIONS_PER_HOUR": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.bookmarks, "MAX_WRITE_OPERATIONS_PER_HOUR", A.H.get<object>(val), WEBEXT.bookmarks)
        ? A.H.TRUE
        : A.H.FALSE;
    },

    "store_MoveArgDestination": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "index" in x ? true : false);
        A.store.Int64(ptr + 0, x["index"] === undefined ? 0 : (x["index"] as number));
        A.store.Ref(ptr + 8, x["parentId"]);
      }
    },
    "load_MoveArgDestination": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["index"] = A.load.Int64(ptr + 0);
      } else {
        delete x["index"];
      }
      x["parentId"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnChangedArgChangeInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["title"]);
        A.store.Ref(ptr + 4, x["url"]);
      }
    },
    "load_OnChangedArgChangeInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["title"] = A.load.Ref(ptr + 0, undefined);
      x["url"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnChildrenReorderedArgReorderInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["childIds"]);
      }
    },
    "load_OnChildrenReorderedArgReorderInfo": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["childIds"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnMovedArgMoveInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 24, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
      } else {
        A.store.Bool(ptr + 24, true);
        A.store.Int64(ptr + 0, x["index"] === undefined ? 0 : (x["index"] as number));
        A.store.Int64(ptr + 8, x["oldIndex"] === undefined ? 0 : (x["oldIndex"] as number));
        A.store.Ref(ptr + 16, x["oldParentId"]);
        A.store.Ref(ptr + 20, x["parentId"]);
      }
    },
    "load_OnMovedArgMoveInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["index"] = A.load.Int64(ptr + 0);
      x["oldIndex"] = A.load.Int64(ptr + 8);
      x["oldParentId"] = A.load.Ref(ptr + 16, undefined);
      x["parentId"] = A.load.Ref(ptr + 20, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnRemovedArgRemoveInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 84, false);
        A.store.Int64(ptr + 0, 0);

        A.store.Bool(ptr + 8 + 68, false);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Bool(ptr + 8 + 64, false);
        A.store.Float64(ptr + 8 + 8, 0);
        A.store.Bool(ptr + 8 + 65, false);
        A.store.Float64(ptr + 8 + 16, 0);
        A.store.Bool(ptr + 8 + 66, false);
        A.store.Float64(ptr + 8 + 24, 0);
        A.store.Ref(ptr + 8 + 32, undefined);
        A.store.Bool(ptr + 8 + 67, false);
        A.store.Int64(ptr + 8 + 40, 0);
        A.store.Ref(ptr + 8 + 48, undefined);
        A.store.Ref(ptr + 8 + 52, undefined);
        A.store.Enum(ptr + 8 + 56, -1);
        A.store.Ref(ptr + 8 + 60, undefined);
        A.store.Ref(ptr + 80, undefined);
      } else {
        A.store.Bool(ptr + 84, true);
        A.store.Int64(ptr + 0, x["index"] === undefined ? 0 : (x["index"] as number));

        if (typeof x["node"] === "undefined") {
          A.store.Bool(ptr + 8 + 68, false);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Bool(ptr + 8 + 64, false);
          A.store.Float64(ptr + 8 + 8, 0);
          A.store.Bool(ptr + 8 + 65, false);
          A.store.Float64(ptr + 8 + 16, 0);
          A.store.Bool(ptr + 8 + 66, false);
          A.store.Float64(ptr + 8 + 24, 0);
          A.store.Ref(ptr + 8 + 32, undefined);
          A.store.Bool(ptr + 8 + 67, false);
          A.store.Int64(ptr + 8 + 40, 0);
          A.store.Ref(ptr + 8 + 48, undefined);
          A.store.Ref(ptr + 8 + 52, undefined);
          A.store.Enum(ptr + 8 + 56, -1);
          A.store.Ref(ptr + 8 + 60, undefined);
        } else {
          A.store.Bool(ptr + 8 + 68, true);
          A.store.Ref(ptr + 8 + 0, x["node"]["children"]);
          A.store.Bool(ptr + 8 + 64, "dateAdded" in x["node"] ? true : false);
          A.store.Float64(ptr + 8 + 8, x["node"]["dateAdded"] === undefined ? 0 : (x["node"]["dateAdded"] as number));
          A.store.Bool(ptr + 8 + 65, "dateGroupModified" in x["node"] ? true : false);
          A.store.Float64(
            ptr + 8 + 16,
            x["node"]["dateGroupModified"] === undefined ? 0 : (x["node"]["dateGroupModified"] as number)
          );
          A.store.Bool(ptr + 8 + 66, "dateLastUsed" in x["node"] ? true : false);
          A.store.Float64(
            ptr + 8 + 24,
            x["node"]["dateLastUsed"] === undefined ? 0 : (x["node"]["dateLastUsed"] as number)
          );
          A.store.Ref(ptr + 8 + 32, x["node"]["id"]);
          A.store.Bool(ptr + 8 + 67, "index" in x["node"] ? true : false);
          A.store.Int64(ptr + 8 + 40, x["node"]["index"] === undefined ? 0 : (x["node"]["index"] as number));
          A.store.Ref(ptr + 8 + 48, x["node"]["parentId"]);
          A.store.Ref(ptr + 8 + 52, x["node"]["title"]);
          A.store.Enum(ptr + 8 + 56, ["managed"].indexOf(x["node"]["unmodifiable"] as string));
          A.store.Ref(ptr + 8 + 60, x["node"]["url"]);
        }
        A.store.Ref(ptr + 80, x["parentId"]);
      }
    },
    "load_OnRemovedArgRemoveInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["index"] = A.load.Int64(ptr + 0);
      if (A.load.Bool(ptr + 8 + 68)) {
        x["node"] = {};
        x["node"]["children"] = A.load.Ref(ptr + 8 + 0, undefined);
        if (A.load.Bool(ptr + 8 + 64)) {
          x["node"]["dateAdded"] = A.load.Float64(ptr + 8 + 8);
        } else {
          delete x["node"]["dateAdded"];
        }
        if (A.load.Bool(ptr + 8 + 65)) {
          x["node"]["dateGroupModified"] = A.load.Float64(ptr + 8 + 16);
        } else {
          delete x["node"]["dateGroupModified"];
        }
        if (A.load.Bool(ptr + 8 + 66)) {
          x["node"]["dateLastUsed"] = A.load.Float64(ptr + 8 + 24);
        } else {
          delete x["node"]["dateLastUsed"];
        }
        x["node"]["id"] = A.load.Ref(ptr + 8 + 32, undefined);
        if (A.load.Bool(ptr + 8 + 67)) {
          x["node"]["index"] = A.load.Int64(ptr + 8 + 40);
        } else {
          delete x["node"]["index"];
        }
        x["node"]["parentId"] = A.load.Ref(ptr + 8 + 48, undefined);
        x["node"]["title"] = A.load.Ref(ptr + 8 + 52, undefined);
        x["node"]["unmodifiable"] = A.load.Enum(ptr + 8 + 56, ["managed"]);
        x["node"]["url"] = A.load.Ref(ptr + 8 + 60, undefined);
      } else {
        delete x["node"];
      }
      x["parentId"] = A.load.Ref(ptr + 80, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SearchArgQueryChoice1": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["query"]);
        A.store.Ref(ptr + 4, x["title"]);
        A.store.Ref(ptr + 8, x["url"]);
      }
    },
    "load_SearchArgQueryChoice1": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["query"] = A.load.Ref(ptr + 0, undefined);
      x["title"] = A.load.Ref(ptr + 4, undefined);
      x["url"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UpdateArgChanges": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["title"]);
        A.store.Ref(ptr + 4, x["url"]);
      }
    },
    "load_UpdateArgChanges": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["title"] = A.load.Ref(ptr + 0, undefined);
      x["url"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Create": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks && "create" in WEBEXT?.bookmarks) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Create": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.create);
    },
    "call_Create": (retPtr: Pointer, bookmark: Pointer): void => {
      const bookmark_ffi = {};

      if (A.load.Bool(bookmark + 20)) {
        bookmark_ffi["index"] = A.load.Int64(bookmark + 0);
      }
      bookmark_ffi["parentId"] = A.load.Ref(bookmark + 8, undefined);
      bookmark_ffi["title"] = A.load.Ref(bookmark + 12, undefined);
      bookmark_ffi["url"] = A.load.Ref(bookmark + 16, undefined);

      const _ret = WEBEXT.bookmarks.create(bookmark_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Create": (retPtr: Pointer, errPtr: Pointer, bookmark: Pointer): heap.Ref<boolean> => {
      try {
        const bookmark_ffi = {};

        if (A.load.Bool(bookmark + 20)) {
          bookmark_ffi["index"] = A.load.Int64(bookmark + 0);
        }
        bookmark_ffi["parentId"] = A.load.Ref(bookmark + 8, undefined);
        bookmark_ffi["title"] = A.load.Ref(bookmark + 12, undefined);
        bookmark_ffi["url"] = A.load.Ref(bookmark + 16, undefined);

        const _ret = WEBEXT.bookmarks.create(bookmark_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Get": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks && "get" in WEBEXT?.bookmarks) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Get": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.get);
    },
    "call_Get": (retPtr: Pointer, idOrIdList: heap.Ref<any>): void => {
      const _ret = WEBEXT.bookmarks.get(A.H.get<any>(idOrIdList));
      A.store.Ref(retPtr, _ret);
    },
    "try_Get": (retPtr: Pointer, errPtr: Pointer, idOrIdList: heap.Ref<any>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.get(A.H.get<any>(idOrIdList));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetChildren": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks && "getChildren" in WEBEXT?.bookmarks) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetChildren": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.getChildren);
    },
    "call_GetChildren": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.getChildren(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetChildren": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.getChildren(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetRecent": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks && "getRecent" in WEBEXT?.bookmarks) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetRecent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.getRecent);
    },
    "call_GetRecent": (retPtr: Pointer, numberOfItems: number): void => {
      const _ret = WEBEXT.bookmarks.getRecent(numberOfItems);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetRecent": (retPtr: Pointer, errPtr: Pointer, numberOfItems: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.getRecent(numberOfItems);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSubTree": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks && "getSubTree" in WEBEXT?.bookmarks) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSubTree": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.getSubTree);
    },
    "call_GetSubTree": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.getSubTree(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSubTree": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.getSubTree(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetTree": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks && "getTree" in WEBEXT?.bookmarks) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetTree": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.getTree);
    },
    "call_GetTree": (retPtr: Pointer): void => {
      const _ret = WEBEXT.bookmarks.getTree();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetTree": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.getTree();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Move": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks && "move" in WEBEXT?.bookmarks) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Move": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.move);
    },
    "call_Move": (retPtr: Pointer, id: heap.Ref<object>, destination: Pointer): void => {
      const destination_ffi = {};

      if (A.load.Bool(destination + 12)) {
        destination_ffi["index"] = A.load.Int64(destination + 0);
      }
      destination_ffi["parentId"] = A.load.Ref(destination + 8, undefined);

      const _ret = WEBEXT.bookmarks.move(A.H.get<object>(id), destination_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Move": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>, destination: Pointer): heap.Ref<boolean> => {
      try {
        const destination_ffi = {};

        if (A.load.Bool(destination + 12)) {
          destination_ffi["index"] = A.load.Int64(destination + 0);
        }
        destination_ffi["parentId"] = A.load.Ref(destination + 8, undefined);

        const _ret = WEBEXT.bookmarks.move(A.H.get<object>(id), destination_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onChanged && "addListener" in WEBEXT?.bookmarks?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onChanged.addListener);
    },
    "call_OnChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onChanged && "removeListener" in WEBEXT?.bookmarks?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onChanged.removeListener);
    },
    "call_OffChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onChanged && "hasListener" in WEBEXT?.bookmarks?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onChanged.hasListener);
    },
    "call_HasOnChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnChildrenReordered": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onChildrenReordered && "addListener" in WEBEXT?.bookmarks?.onChildrenReordered) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnChildrenReordered": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onChildrenReordered.addListener);
    },
    "call_OnChildrenReordered": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onChildrenReordered.addListener(A.H.get<object>(callback));
    },
    "try_OnChildrenReordered": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onChildrenReordered.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffChildrenReordered": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onChildrenReordered && "removeListener" in WEBEXT?.bookmarks?.onChildrenReordered) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffChildrenReordered": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onChildrenReordered.removeListener);
    },
    "call_OffChildrenReordered": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onChildrenReordered.removeListener(A.H.get<object>(callback));
    },
    "try_OffChildrenReordered": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onChildrenReordered.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnChildrenReordered": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onChildrenReordered && "hasListener" in WEBEXT?.bookmarks?.onChildrenReordered) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnChildrenReordered": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onChildrenReordered.hasListener);
    },
    "call_HasOnChildrenReordered": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onChildrenReordered.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnChildrenReordered": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onChildrenReordered.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onCreated && "addListener" in WEBEXT?.bookmarks?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onCreated.addListener);
    },
    "call_OnCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onCreated.addListener(A.H.get<object>(callback));
    },
    "try_OnCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onCreated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onCreated && "removeListener" in WEBEXT?.bookmarks?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onCreated.removeListener);
    },
    "call_OffCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onCreated.removeListener(A.H.get<object>(callback));
    },
    "try_OffCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onCreated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onCreated && "hasListener" in WEBEXT?.bookmarks?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onCreated.hasListener);
    },
    "call_HasOnCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onCreated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onCreated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnImportBegan": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onImportBegan && "addListener" in WEBEXT?.bookmarks?.onImportBegan) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnImportBegan": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onImportBegan.addListener);
    },
    "call_OnImportBegan": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onImportBegan.addListener(A.H.get<object>(callback));
    },
    "try_OnImportBegan": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onImportBegan.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffImportBegan": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onImportBegan && "removeListener" in WEBEXT?.bookmarks?.onImportBegan) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffImportBegan": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onImportBegan.removeListener);
    },
    "call_OffImportBegan": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onImportBegan.removeListener(A.H.get<object>(callback));
    },
    "try_OffImportBegan": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onImportBegan.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnImportBegan": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onImportBegan && "hasListener" in WEBEXT?.bookmarks?.onImportBegan) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnImportBegan": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onImportBegan.hasListener);
    },
    "call_HasOnImportBegan": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onImportBegan.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnImportBegan": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onImportBegan.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnImportEnded": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onImportEnded && "addListener" in WEBEXT?.bookmarks?.onImportEnded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnImportEnded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onImportEnded.addListener);
    },
    "call_OnImportEnded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onImportEnded.addListener(A.H.get<object>(callback));
    },
    "try_OnImportEnded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onImportEnded.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffImportEnded": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onImportEnded && "removeListener" in WEBEXT?.bookmarks?.onImportEnded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffImportEnded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onImportEnded.removeListener);
    },
    "call_OffImportEnded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onImportEnded.removeListener(A.H.get<object>(callback));
    },
    "try_OffImportEnded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onImportEnded.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnImportEnded": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onImportEnded && "hasListener" in WEBEXT?.bookmarks?.onImportEnded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnImportEnded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onImportEnded.hasListener);
    },
    "call_HasOnImportEnded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onImportEnded.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnImportEnded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onImportEnded.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onMoved && "addListener" in WEBEXT?.bookmarks?.onMoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onMoved.addListener);
    },
    "call_OnMoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onMoved.addListener(A.H.get<object>(callback));
    },
    "try_OnMoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onMoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onMoved && "removeListener" in WEBEXT?.bookmarks?.onMoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onMoved.removeListener);
    },
    "call_OffMoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onMoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffMoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onMoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onMoved && "hasListener" in WEBEXT?.bookmarks?.onMoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onMoved.hasListener);
    },
    "call_HasOnMoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onMoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onMoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onRemoved && "addListener" in WEBEXT?.bookmarks?.onRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onRemoved.addListener);
    },
    "call_OnRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onRemoved && "removeListener" in WEBEXT?.bookmarks?.onRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onRemoved.removeListener);
    },
    "call_OffRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks?.onRemoved && "hasListener" in WEBEXT?.bookmarks?.onRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.onRemoved.hasListener);
    },
    "call_HasOnRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.onRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.onRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Remove": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks && "remove" in WEBEXT?.bookmarks) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Remove": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.remove);
    },
    "call_Remove": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.remove(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_Remove": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.remove(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveTree": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks && "removeTree" in WEBEXT?.bookmarks) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveTree": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.removeTree);
    },
    "call_RemoveTree": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarks.removeTree(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveTree": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.removeTree(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Search": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks && "search" in WEBEXT?.bookmarks) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Search": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.search);
    },
    "call_Search": (retPtr: Pointer, query: heap.Ref<any>): void => {
      const _ret = WEBEXT.bookmarks.search(A.H.get<any>(query));
      A.store.Ref(retPtr, _ret);
    },
    "try_Search": (retPtr: Pointer, errPtr: Pointer, query: heap.Ref<any>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarks.search(A.H.get<any>(query));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Update": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarks && "update" in WEBEXT?.bookmarks) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Update": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarks.update);
    },
    "call_Update": (retPtr: Pointer, id: heap.Ref<object>, changes: Pointer): void => {
      const changes_ffi = {};

      changes_ffi["title"] = A.load.Ref(changes + 0, undefined);
      changes_ffi["url"] = A.load.Ref(changes + 4, undefined);

      const _ret = WEBEXT.bookmarks.update(A.H.get<object>(id), changes_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Update": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>, changes: Pointer): heap.Ref<boolean> => {
      try {
        const changes_ffi = {};

        changes_ffi["title"] = A.load.Ref(changes + 0, undefined);
        changes_ffi["url"] = A.load.Ref(changes + 4, undefined);

        const _ret = WEBEXT.bookmarks.update(A.H.get<object>(id), changes_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
