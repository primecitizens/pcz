import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/tabgroups", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_Color": (ref: heap.Ref<string>): number => {
      const idx = ["grey", "blue", "red", "yellow", "green", "pink", "purple", "cyan", "orange"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_MoveArgMoveProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Int64(ptr + 0, x["index"] === undefined ? 0 : (x["index"] as number));
        A.store.Bool(ptr + 16, "windowId" in x ? true : false);
        A.store.Int64(ptr + 8, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_MoveArgMoveProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["index"] = A.load.Int64(ptr + 0);
      if (A.load.Bool(ptr + 16)) {
        x["windowId"] = A.load.Int64(ptr + 8);
      } else {
        delete x["windowId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_QueryArgQueryInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 0, false);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 25, false);
        A.store.Int64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 26, true);
        A.store.Bool(ptr + 24, "collapsed" in x ? true : false);
        A.store.Bool(ptr + 0, x["collapsed"] ? true : false);
        A.store.Enum(
          ptr + 4,
          ["grey", "blue", "red", "yellow", "green", "pink", "purple", "cyan", "orange"].indexOf(x["color"] as string)
        );
        A.store.Ref(ptr + 8, x["title"]);
        A.store.Bool(ptr + 25, "windowId" in x ? true : false);
        A.store.Int64(ptr + 16, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_QueryArgQueryInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["collapsed"] = A.load.Bool(ptr + 0);
      } else {
        delete x["collapsed"];
      }
      x["color"] = A.load.Enum(ptr + 4, ["grey", "blue", "red", "yellow", "green", "pink", "purple", "cyan", "orange"]);
      x["title"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 25)) {
        x["windowId"] = A.load.Int64(ptr + 16);
      } else {
        delete x["windowId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "get_TAB_GROUP_ID_NONE": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups && "TAB_GROUP_ID_NONE" in WEBEXT?.tabGroups) {
        const val = WEBEXT.tabGroups.TAB_GROUP_ID_NONE;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_TAB_GROUP_ID_NONE": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.tabGroups, "TAB_GROUP_ID_NONE", A.H.get<object>(val), WEBEXT.tabGroups)
        ? A.H.TRUE
        : A.H.FALSE;
    },

    "store_TabGroup": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 0, false);
        A.store.Enum(ptr + 4, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Int64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 32, true);
        A.store.Bool(ptr + 0, x["collapsed"] ? true : false);
        A.store.Enum(
          ptr + 4,
          ["grey", "blue", "red", "yellow", "green", "pink", "purple", "cyan", "orange"].indexOf(x["color"] as string)
        );
        A.store.Int64(ptr + 8, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Ref(ptr + 16, x["title"]);
        A.store.Int64(ptr + 24, x["windowId"] === undefined ? 0 : (x["windowId"] as number));
      }
    },
    "load_TabGroup": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["collapsed"] = A.load.Bool(ptr + 0);
      x["color"] = A.load.Enum(ptr + 4, ["grey", "blue", "red", "yellow", "green", "pink", "purple", "cyan", "orange"]);
      x["id"] = A.load.Int64(ptr + 8);
      x["title"] = A.load.Ref(ptr + 16, undefined);
      x["windowId"] = A.load.Int64(ptr + 24);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UpdateArgUpdateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 0, false);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "collapsed" in x ? true : false);
        A.store.Bool(ptr + 0, x["collapsed"] ? true : false);
        A.store.Enum(
          ptr + 4,
          ["grey", "blue", "red", "yellow", "green", "pink", "purple", "cyan", "orange"].indexOf(x["color"] as string)
        );
        A.store.Ref(ptr + 8, x["title"]);
      }
    },
    "load_UpdateArgUpdateProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["collapsed"] = A.load.Bool(ptr + 0);
      } else {
        delete x["collapsed"];
      }
      x["color"] = A.load.Enum(ptr + 4, ["grey", "blue", "red", "yellow", "green", "pink", "purple", "cyan", "orange"]);
      x["title"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Get": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups && "get" in WEBEXT?.tabGroups) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Get": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.get);
    },
    "call_Get": (retPtr: Pointer, groupId: number): void => {
      const _ret = WEBEXT.tabGroups.get(groupId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Get": (retPtr: Pointer, errPtr: Pointer, groupId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabGroups.get(groupId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Move": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups && "move" in WEBEXT?.tabGroups) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Move": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.move);
    },
    "call_Move": (retPtr: Pointer, groupId: number, moveProperties: Pointer): void => {
      const moveProperties_ffi = {};

      moveProperties_ffi["index"] = A.load.Int64(moveProperties + 0);
      if (A.load.Bool(moveProperties + 16)) {
        moveProperties_ffi["windowId"] = A.load.Int64(moveProperties + 8);
      }

      const _ret = WEBEXT.tabGroups.move(groupId, moveProperties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Move": (retPtr: Pointer, errPtr: Pointer, groupId: number, moveProperties: Pointer): heap.Ref<boolean> => {
      try {
        const moveProperties_ffi = {};

        moveProperties_ffi["index"] = A.load.Int64(moveProperties + 0);
        if (A.load.Bool(moveProperties + 16)) {
          moveProperties_ffi["windowId"] = A.load.Int64(moveProperties + 8);
        }

        const _ret = WEBEXT.tabGroups.move(groupId, moveProperties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups?.onCreated && "addListener" in WEBEXT?.tabGroups?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.onCreated.addListener);
    },
    "call_OnCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabGroups.onCreated.addListener(A.H.get<object>(callback));
    },
    "try_OnCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabGroups.onCreated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups?.onCreated && "removeListener" in WEBEXT?.tabGroups?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.onCreated.removeListener);
    },
    "call_OffCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabGroups.onCreated.removeListener(A.H.get<object>(callback));
    },
    "try_OffCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabGroups.onCreated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups?.onCreated && "hasListener" in WEBEXT?.tabGroups?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.onCreated.hasListener);
    },
    "call_HasOnCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabGroups.onCreated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabGroups.onCreated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups?.onMoved && "addListener" in WEBEXT?.tabGroups?.onMoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.onMoved.addListener);
    },
    "call_OnMoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabGroups.onMoved.addListener(A.H.get<object>(callback));
    },
    "try_OnMoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabGroups.onMoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups?.onMoved && "removeListener" in WEBEXT?.tabGroups?.onMoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.onMoved.removeListener);
    },
    "call_OffMoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabGroups.onMoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffMoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabGroups.onMoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups?.onMoved && "hasListener" in WEBEXT?.tabGroups?.onMoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.onMoved.hasListener);
    },
    "call_HasOnMoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabGroups.onMoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabGroups.onMoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups?.onRemoved && "addListener" in WEBEXT?.tabGroups?.onRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.onRemoved.addListener);
    },
    "call_OnRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabGroups.onRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabGroups.onRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups?.onRemoved && "removeListener" in WEBEXT?.tabGroups?.onRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.onRemoved.removeListener);
    },
    "call_OffRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabGroups.onRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabGroups.onRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups?.onRemoved && "hasListener" in WEBEXT?.tabGroups?.onRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.onRemoved.hasListener);
    },
    "call_HasOnRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabGroups.onRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabGroups.onRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnUpdated": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups?.onUpdated && "addListener" in WEBEXT?.tabGroups?.onUpdated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.onUpdated.addListener);
    },
    "call_OnUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabGroups.onUpdated.addListener(A.H.get<object>(callback));
    },
    "try_OnUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabGroups.onUpdated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffUpdated": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups?.onUpdated && "removeListener" in WEBEXT?.tabGroups?.onUpdated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.onUpdated.removeListener);
    },
    "call_OffUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabGroups.onUpdated.removeListener(A.H.get<object>(callback));
    },
    "try_OffUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabGroups.onUpdated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnUpdated": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups?.onUpdated && "hasListener" in WEBEXT?.tabGroups?.onUpdated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnUpdated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.onUpdated.hasListener);
    },
    "call_HasOnUpdated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tabGroups.onUpdated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnUpdated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tabGroups.onUpdated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Query": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups && "query" in WEBEXT?.tabGroups) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Query": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.query);
    },
    "call_Query": (retPtr: Pointer, queryInfo: Pointer): void => {
      const queryInfo_ffi = {};

      if (A.load.Bool(queryInfo + 24)) {
        queryInfo_ffi["collapsed"] = A.load.Bool(queryInfo + 0);
      }
      queryInfo_ffi["color"] = A.load.Enum(queryInfo + 4, [
        "grey",
        "blue",
        "red",
        "yellow",
        "green",
        "pink",
        "purple",
        "cyan",
        "orange",
      ]);
      queryInfo_ffi["title"] = A.load.Ref(queryInfo + 8, undefined);
      if (A.load.Bool(queryInfo + 25)) {
        queryInfo_ffi["windowId"] = A.load.Int64(queryInfo + 16);
      }

      const _ret = WEBEXT.tabGroups.query(queryInfo_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Query": (retPtr: Pointer, errPtr: Pointer, queryInfo: Pointer): heap.Ref<boolean> => {
      try {
        const queryInfo_ffi = {};

        if (A.load.Bool(queryInfo + 24)) {
          queryInfo_ffi["collapsed"] = A.load.Bool(queryInfo + 0);
        }
        queryInfo_ffi["color"] = A.load.Enum(queryInfo + 4, [
          "grey",
          "blue",
          "red",
          "yellow",
          "green",
          "pink",
          "purple",
          "cyan",
          "orange",
        ]);
        queryInfo_ffi["title"] = A.load.Ref(queryInfo + 8, undefined);
        if (A.load.Bool(queryInfo + 25)) {
          queryInfo_ffi["windowId"] = A.load.Int64(queryInfo + 16);
        }

        const _ret = WEBEXT.tabGroups.query(queryInfo_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Update": (): heap.Ref<boolean> => {
      if (WEBEXT?.tabGroups && "update" in WEBEXT?.tabGroups) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Update": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tabGroups.update);
    },
    "call_Update": (retPtr: Pointer, groupId: number, updateProperties: Pointer): void => {
      const updateProperties_ffi = {};

      if (A.load.Bool(updateProperties + 12)) {
        updateProperties_ffi["collapsed"] = A.load.Bool(updateProperties + 0);
      }
      updateProperties_ffi["color"] = A.load.Enum(updateProperties + 4, [
        "grey",
        "blue",
        "red",
        "yellow",
        "green",
        "pink",
        "purple",
        "cyan",
        "orange",
      ]);
      updateProperties_ffi["title"] = A.load.Ref(updateProperties + 8, undefined);

      const _ret = WEBEXT.tabGroups.update(groupId, updateProperties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Update": (retPtr: Pointer, errPtr: Pointer, groupId: number, updateProperties: Pointer): heap.Ref<boolean> => {
      try {
        const updateProperties_ffi = {};

        if (A.load.Bool(updateProperties + 12)) {
          updateProperties_ffi["collapsed"] = A.load.Bool(updateProperties + 0);
        }
        updateProperties_ffi["color"] = A.load.Enum(updateProperties + 4, [
          "grey",
          "blue",
          "red",
          "yellow",
          "green",
          "pink",
          "purple",
          "cyan",
          "orange",
        ]);
        updateProperties_ffi["title"] = A.load.Ref(updateProperties + 8, undefined);

        const _ret = WEBEXT.tabGroups.update(groupId, updateProperties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
