import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/appviewguestinternal", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_AttachFrame": (): heap.Ref<boolean> => {
      if (WEBEXT?.appViewGuestInternal && "attachFrame" in WEBEXT?.appViewGuestInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AttachFrame": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.appViewGuestInternal.attachFrame);
    },
    "call_AttachFrame": (retPtr: Pointer, url: heap.Ref<object>, guestInstanceId: number): void => {
      const _ret = WEBEXT.appViewGuestInternal.attachFrame(A.H.get<object>(url), guestInstanceId);
      A.store.Ref(retPtr, _ret);
    },
    "try_AttachFrame": (
      retPtr: Pointer,
      errPtr: Pointer,
      url: heap.Ref<object>,
      guestInstanceId: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.appViewGuestInternal.attachFrame(A.H.get<object>(url), guestInstanceId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DenyRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.appViewGuestInternal && "denyRequest" in WEBEXT?.appViewGuestInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DenyRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.appViewGuestInternal.denyRequest);
    },
    "call_DenyRequest": (retPtr: Pointer, guestInstanceId: number): void => {
      const _ret = WEBEXT.appViewGuestInternal.denyRequest(guestInstanceId);
    },
    "try_DenyRequest": (retPtr: Pointer, errPtr: Pointer, guestInstanceId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.appViewGuestInternal.denyRequest(guestInstanceId);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
