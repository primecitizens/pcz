import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/sharedstorageprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_Get": (): heap.Ref<boolean> => {
      if (WEBEXT?.sharedStoragePrivate && "get" in WEBEXT?.sharedStoragePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Get": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sharedStoragePrivate.get);
    },
    "call_Get": (retPtr: Pointer): void => {
      const _ret = WEBEXT.sharedStoragePrivate.get();
      A.store.Ref(retPtr, _ret);
    },
    "try_Get": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sharedStoragePrivate.get();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Remove": (): heap.Ref<boolean> => {
      if (WEBEXT?.sharedStoragePrivate && "remove" in WEBEXT?.sharedStoragePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Remove": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sharedStoragePrivate.remove);
    },
    "call_Remove": (retPtr: Pointer, keys: heap.Ref<object>): void => {
      const _ret = WEBEXT.sharedStoragePrivate.remove(A.H.get<object>(keys));
      A.store.Ref(retPtr, _ret);
    },
    "try_Remove": (retPtr: Pointer, errPtr: Pointer, keys: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sharedStoragePrivate.remove(A.H.get<object>(keys));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Set": (): heap.Ref<boolean> => {
      if (WEBEXT?.sharedStoragePrivate && "set" in WEBEXT?.sharedStoragePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Set": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sharedStoragePrivate.set);
    },
    "call_Set": (retPtr: Pointer, items: heap.Ref<object>): void => {
      const _ret = WEBEXT.sharedStoragePrivate.set(A.H.get<object>(items));
      A.store.Ref(retPtr, _ret);
    },
    "try_Set": (retPtr: Pointer, errPtr: Pointer, items: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sharedStoragePrivate.set(A.H.get<object>(items));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
