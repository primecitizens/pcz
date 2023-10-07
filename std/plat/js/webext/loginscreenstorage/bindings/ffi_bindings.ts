import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/loginscreenstorage", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_RetrieveCredentials": (): heap.Ref<boolean> => {
      if (WEBEXT?.loginScreenStorage && "retrieveCredentials" in WEBEXT?.loginScreenStorage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RetrieveCredentials": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.loginScreenStorage.retrieveCredentials);
    },
    "call_RetrieveCredentials": (retPtr: Pointer): void => {
      const _ret = WEBEXT.loginScreenStorage.retrieveCredentials();
      A.store.Ref(retPtr, _ret);
    },
    "try_RetrieveCredentials": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.loginScreenStorage.retrieveCredentials();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RetrievePersistentData": (): heap.Ref<boolean> => {
      if (WEBEXT?.loginScreenStorage && "retrievePersistentData" in WEBEXT?.loginScreenStorage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RetrievePersistentData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.loginScreenStorage.retrievePersistentData);
    },
    "call_RetrievePersistentData": (retPtr: Pointer, ownerId: heap.Ref<object>): void => {
      const _ret = WEBEXT.loginScreenStorage.retrievePersistentData(A.H.get<object>(ownerId));
      A.store.Ref(retPtr, _ret);
    },
    "try_RetrievePersistentData": (retPtr: Pointer, errPtr: Pointer, ownerId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.loginScreenStorage.retrievePersistentData(A.H.get<object>(ownerId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StoreCredentials": (): heap.Ref<boolean> => {
      if (WEBEXT?.loginScreenStorage && "storeCredentials" in WEBEXT?.loginScreenStorage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StoreCredentials": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.loginScreenStorage.storeCredentials);
    },
    "call_StoreCredentials": (retPtr: Pointer, extensionId: heap.Ref<object>, credentials: heap.Ref<object>): void => {
      const _ret = WEBEXT.loginScreenStorage.storeCredentials(
        A.H.get<object>(extensionId),
        A.H.get<object>(credentials)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_StoreCredentials": (
      retPtr: Pointer,
      errPtr: Pointer,
      extensionId: heap.Ref<object>,
      credentials: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.loginScreenStorage.storeCredentials(
          A.H.get<object>(extensionId),
          A.H.get<object>(credentials)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StorePersistentData": (): heap.Ref<boolean> => {
      if (WEBEXT?.loginScreenStorage && "storePersistentData" in WEBEXT?.loginScreenStorage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StorePersistentData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.loginScreenStorage.storePersistentData);
    },
    "call_StorePersistentData": (retPtr: Pointer, extensionIds: heap.Ref<object>, data: heap.Ref<object>): void => {
      const _ret = WEBEXT.loginScreenStorage.storePersistentData(A.H.get<object>(extensionIds), A.H.get<object>(data));
      A.store.Ref(retPtr, _ret);
    },
    "try_StorePersistentData": (
      retPtr: Pointer,
      errPtr: Pointer,
      extensionIds: heap.Ref<object>,
      data: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.loginScreenStorage.storePersistentData(
          A.H.get<object>(extensionIds),
          A.H.get<object>(data)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
