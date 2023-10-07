import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/appviewtag", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_EmbedRequest_Allow": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "allow" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EmbedRequest_Allow": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).allow);
    },

    "call_EmbedRequest_Allow": (self: heap.Ref<object>, retPtr: Pointer, url: heap.Ref<object>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.allow(A.H.get<object>(url));
    },
    "try_EmbedRequest_Allow": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      url: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.allow(A.H.get<object>(url));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_EmbedRequest_Deny": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "deny" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EmbedRequest_Deny": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).deny);
    },

    "call_EmbedRequest_Deny": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.deny();
    },
    "try_EmbedRequest_Deny": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.deny();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_EmbedRequest_EmbedderId": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "embedderId" in thiz) {
        const val = thiz.embedderId;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_EmbedRequest_EmbedderId": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "embedderId", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_EmbedRequest_Data": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "data" in thiz) {
        const val = thiz.data;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_EmbedRequest_Data": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "data", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "has_Connect": (): heap.Ref<boolean> => {
      if (WEBEXT?.appviewTag && "connect" in WEBEXT?.appviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Connect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.appviewTag.connect);
    },
    "call_Connect": (
      retPtr: Pointer,
      app: heap.Ref<object>,
      data: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.appviewTag.connect(A.H.get<object>(app), A.H.get<object>(data), A.H.get<object>(callback));
    },
    "try_Connect": (
      retPtr: Pointer,
      errPtr: Pointer,
      app: heap.Ref<object>,
      data: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.appviewTag.connect(A.H.get<object>(app), A.H.get<object>(data), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
