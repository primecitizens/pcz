import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/gcm", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "get_MAX_MESSAGE_SIZE": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.gcm && "MAX_MESSAGE_SIZE" in WEBEXT?.gcm) {
        const val = WEBEXT.gcm.MAX_MESSAGE_SIZE;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_MAX_MESSAGE_SIZE": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.gcm, "MAX_MESSAGE_SIZE", A.H.get<object>(val), WEBEXT.gcm) ? A.H.TRUE : A.H.FALSE;
    },

    "store_OnMessageArgMessage": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["collapseKey"]);
        A.store.Ref(ptr + 4, x["data"]);
        A.store.Ref(ptr + 8, x["from"]);
      }
    },
    "load_OnMessageArgMessage": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["collapseKey"] = A.load.Ref(ptr + 0, undefined);
      x["data"] = A.load.Ref(ptr + 4, undefined);
      x["from"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnSendErrorArgError": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["details"]);
        A.store.Ref(ptr + 4, x["errorMessage"]);
        A.store.Ref(ptr + 8, x["messageId"]);
      }
    },
    "load_OnSendErrorArgError": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["details"] = A.load.Ref(ptr + 0, undefined);
      x["errorMessage"] = A.load.Ref(ptr + 4, undefined);
      x["messageId"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SendArgMessage": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 24, false);
        A.store.Int64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 25, true);
        A.store.Ref(ptr + 0, x["data"]);
        A.store.Ref(ptr + 4, x["destinationId"]);
        A.store.Ref(ptr + 8, x["messageId"]);
        A.store.Bool(ptr + 24, "timeToLive" in x ? true : false);
        A.store.Int64(ptr + 16, x["timeToLive"] === undefined ? 0 : (x["timeToLive"] as number));
      }
    },
    "load_SendArgMessage": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["data"] = A.load.Ref(ptr + 0, undefined);
      x["destinationId"] = A.load.Ref(ptr + 4, undefined);
      x["messageId"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 24)) {
        x["timeToLive"] = A.load.Int64(ptr + 16);
      } else {
        delete x["timeToLive"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_OnMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.gcm?.onMessage && "addListener" in WEBEXT?.gcm?.onMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.gcm.onMessage.addListener);
    },
    "call_OnMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.gcm.onMessage.addListener(A.H.get<object>(callback));
    },
    "try_OnMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.gcm.onMessage.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.gcm?.onMessage && "removeListener" in WEBEXT?.gcm?.onMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.gcm.onMessage.removeListener);
    },
    "call_OffMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.gcm.onMessage.removeListener(A.H.get<object>(callback));
    },
    "try_OffMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.gcm.onMessage.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.gcm?.onMessage && "hasListener" in WEBEXT?.gcm?.onMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.gcm.onMessage.hasListener);
    },
    "call_HasOnMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.gcm.onMessage.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.gcm.onMessage.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMessagesDeleted": (): heap.Ref<boolean> => {
      if (WEBEXT?.gcm?.onMessagesDeleted && "addListener" in WEBEXT?.gcm?.onMessagesDeleted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMessagesDeleted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.gcm.onMessagesDeleted.addListener);
    },
    "call_OnMessagesDeleted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.gcm.onMessagesDeleted.addListener(A.H.get<object>(callback));
    },
    "try_OnMessagesDeleted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.gcm.onMessagesDeleted.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMessagesDeleted": (): heap.Ref<boolean> => {
      if (WEBEXT?.gcm?.onMessagesDeleted && "removeListener" in WEBEXT?.gcm?.onMessagesDeleted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMessagesDeleted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.gcm.onMessagesDeleted.removeListener);
    },
    "call_OffMessagesDeleted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.gcm.onMessagesDeleted.removeListener(A.H.get<object>(callback));
    },
    "try_OffMessagesDeleted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.gcm.onMessagesDeleted.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMessagesDeleted": (): heap.Ref<boolean> => {
      if (WEBEXT?.gcm?.onMessagesDeleted && "hasListener" in WEBEXT?.gcm?.onMessagesDeleted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMessagesDeleted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.gcm.onMessagesDeleted.hasListener);
    },
    "call_HasOnMessagesDeleted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.gcm.onMessagesDeleted.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMessagesDeleted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.gcm.onMessagesDeleted.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSendError": (): heap.Ref<boolean> => {
      if (WEBEXT?.gcm?.onSendError && "addListener" in WEBEXT?.gcm?.onSendError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSendError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.gcm.onSendError.addListener);
    },
    "call_OnSendError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.gcm.onSendError.addListener(A.H.get<object>(callback));
    },
    "try_OnSendError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.gcm.onSendError.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSendError": (): heap.Ref<boolean> => {
      if (WEBEXT?.gcm?.onSendError && "removeListener" in WEBEXT?.gcm?.onSendError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSendError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.gcm.onSendError.removeListener);
    },
    "call_OffSendError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.gcm.onSendError.removeListener(A.H.get<object>(callback));
    },
    "try_OffSendError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.gcm.onSendError.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSendError": (): heap.Ref<boolean> => {
      if (WEBEXT?.gcm?.onSendError && "hasListener" in WEBEXT?.gcm?.onSendError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSendError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.gcm.onSendError.hasListener);
    },
    "call_HasOnSendError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.gcm.onSendError.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSendError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.gcm.onSendError.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Register": (): heap.Ref<boolean> => {
      if (WEBEXT?.gcm && "register" in WEBEXT?.gcm) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Register": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.gcm.register);
    },
    "call_Register": (retPtr: Pointer, senderIds: heap.Ref<object>): void => {
      const _ret = WEBEXT.gcm.register(A.H.get<object>(senderIds));
      A.store.Ref(retPtr, _ret);
    },
    "try_Register": (retPtr: Pointer, errPtr: Pointer, senderIds: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.gcm.register(A.H.get<object>(senderIds));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Send": (): heap.Ref<boolean> => {
      if (WEBEXT?.gcm && "send" in WEBEXT?.gcm) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Send": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.gcm.send);
    },
    "call_Send": (retPtr: Pointer, message: Pointer): void => {
      const message_ffi = {};

      message_ffi["data"] = A.load.Ref(message + 0, undefined);
      message_ffi["destinationId"] = A.load.Ref(message + 4, undefined);
      message_ffi["messageId"] = A.load.Ref(message + 8, undefined);
      if (A.load.Bool(message + 24)) {
        message_ffi["timeToLive"] = A.load.Int64(message + 16);
      }

      const _ret = WEBEXT.gcm.send(message_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Send": (retPtr: Pointer, errPtr: Pointer, message: Pointer): heap.Ref<boolean> => {
      try {
        const message_ffi = {};

        message_ffi["data"] = A.load.Ref(message + 0, undefined);
        message_ffi["destinationId"] = A.load.Ref(message + 4, undefined);
        message_ffi["messageId"] = A.load.Ref(message + 8, undefined);
        if (A.load.Bool(message + 24)) {
          message_ffi["timeToLive"] = A.load.Int64(message + 16);
        }

        const _ret = WEBEXT.gcm.send(message_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Unregister": (): heap.Ref<boolean> => {
      if (WEBEXT?.gcm && "unregister" in WEBEXT?.gcm) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Unregister": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.gcm.unregister);
    },
    "call_Unregister": (retPtr: Pointer): void => {
      const _ret = WEBEXT.gcm.unregister();
      A.store.Ref(retPtr, _ret);
    },
    "try_Unregister": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.gcm.unregister();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
