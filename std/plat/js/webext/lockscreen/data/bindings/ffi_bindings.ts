import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/lockscreen/data", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_DataItemInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["id"]);
      }
    },
    "load_DataItemInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DataItemsAvailableEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 1, "wasLocked" in x ? true : false);
        A.store.Bool(ptr + 0, x["wasLocked"] ? true : false);
      }
    },
    "load_DataItemsAvailableEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 1)) {
        x["wasLocked"] = A.load.Bool(ptr + 0);
      } else {
        delete x["wasLocked"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Create": (): heap.Ref<boolean> => {
      if (WEBEXT?.lockScreen?.data && "create" in WEBEXT?.lockScreen?.data) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Create": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.lockScreen.data.create);
    },
    "call_Create": (retPtr: Pointer): void => {
      const _ret = WEBEXT.lockScreen.data.create();
      A.store.Ref(retPtr, _ret);
    },
    "try_Create": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.lockScreen.data.create();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Delete": (): heap.Ref<boolean> => {
      if (WEBEXT?.lockScreen?.data && "delete" in WEBEXT?.lockScreen?.data) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Delete": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.lockScreen.data.delete);
    },
    "call_Delete": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.lockScreen.data.delete(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_Delete": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.lockScreen.data.delete(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAll": (): heap.Ref<boolean> => {
      if (WEBEXT?.lockScreen?.data && "getAll" in WEBEXT?.lockScreen?.data) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.lockScreen.data.getAll);
    },
    "call_GetAll": (retPtr: Pointer): void => {
      const _ret = WEBEXT.lockScreen.data.getAll();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAll": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.lockScreen.data.getAll();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetContent": (): heap.Ref<boolean> => {
      if (WEBEXT?.lockScreen?.data && "getContent" in WEBEXT?.lockScreen?.data) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetContent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.lockScreen.data.getContent);
    },
    "call_GetContent": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.lockScreen.data.getContent(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetContent": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.lockScreen.data.getContent(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDataItemsAvailable": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.lockScreen?.data?.onDataItemsAvailable &&
        "addListener" in WEBEXT?.lockScreen?.data?.onDataItemsAvailable
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDataItemsAvailable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.lockScreen.data.onDataItemsAvailable.addListener);
    },
    "call_OnDataItemsAvailable": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.lockScreen.data.onDataItemsAvailable.addListener(A.H.get<object>(callback));
    },
    "try_OnDataItemsAvailable": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.lockScreen.data.onDataItemsAvailable.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDataItemsAvailable": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.lockScreen?.data?.onDataItemsAvailable &&
        "removeListener" in WEBEXT?.lockScreen?.data?.onDataItemsAvailable
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDataItemsAvailable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.lockScreen.data.onDataItemsAvailable.removeListener);
    },
    "call_OffDataItemsAvailable": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.lockScreen.data.onDataItemsAvailable.removeListener(A.H.get<object>(callback));
    },
    "try_OffDataItemsAvailable": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.lockScreen.data.onDataItemsAvailable.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDataItemsAvailable": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.lockScreen?.data?.onDataItemsAvailable &&
        "hasListener" in WEBEXT?.lockScreen?.data?.onDataItemsAvailable
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDataItemsAvailable": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.lockScreen.data.onDataItemsAvailable.hasListener);
    },
    "call_HasOnDataItemsAvailable": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.lockScreen.data.onDataItemsAvailable.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDataItemsAvailable": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.lockScreen.data.onDataItemsAvailable.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetContent": (): heap.Ref<boolean> => {
      if (WEBEXT?.lockScreen?.data && "setContent" in WEBEXT?.lockScreen?.data) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetContent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.lockScreen.data.setContent);
    },
    "call_SetContent": (retPtr: Pointer, id: heap.Ref<object>, data: heap.Ref<object>): void => {
      const _ret = WEBEXT.lockScreen.data.setContent(A.H.get<object>(id), A.H.get<object>(data));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetContent": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<object>,
      data: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.lockScreen.data.setContent(A.H.get<object>(id), A.H.get<object>(data));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
