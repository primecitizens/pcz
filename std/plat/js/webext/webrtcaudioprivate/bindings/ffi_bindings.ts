import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/webrtcaudioprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_SinkInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 14, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 13, false);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["sinkId"]);
        A.store.Ref(ptr + 4, x["sinkLabel"]);
        A.store.Bool(ptr + 14, "sampleRate" in x ? true : false);
        A.store.Int32(ptr + 8, x["sampleRate"] === undefined ? 0 : (x["sampleRate"] as number));
        A.store.Bool(ptr + 15, "isReady" in x ? true : false);
        A.store.Bool(ptr + 12, x["isReady"] ? true : false);
        A.store.Bool(ptr + 16, "isDefault" in x ? true : false);
        A.store.Bool(ptr + 13, x["isDefault"] ? true : false);
      }
    },
    "load_SinkInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["sinkId"] = A.load.Ref(ptr + 0, undefined);
      x["sinkLabel"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 14)) {
        x["sampleRate"] = A.load.Int32(ptr + 8);
      } else {
        delete x["sampleRate"];
      }
      if (A.load.Bool(ptr + 15)) {
        x["isReady"] = A.load.Bool(ptr + 12);
      } else {
        delete x["isReady"];
      }
      if (A.load.Bool(ptr + 16)) {
        x["isDefault"] = A.load.Bool(ptr + 13);
      } else {
        delete x["isDefault"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetAssociatedSink": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcAudioPrivate && "getAssociatedSink" in WEBEXT?.webrtcAudioPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAssociatedSink": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcAudioPrivate.getAssociatedSink);
    },
    "call_GetAssociatedSink": (
      retPtr: Pointer,
      securityOrigin: heap.Ref<object>,
      sourceIdInOrigin: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.webrtcAudioPrivate.getAssociatedSink(
        A.H.get<object>(securityOrigin),
        A.H.get<object>(sourceIdInOrigin)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAssociatedSink": (
      retPtr: Pointer,
      errPtr: Pointer,
      securityOrigin: heap.Ref<object>,
      sourceIdInOrigin: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webrtcAudioPrivate.getAssociatedSink(
          A.H.get<object>(securityOrigin),
          A.H.get<object>(sourceIdInOrigin)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSinks": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcAudioPrivate && "getSinks" in WEBEXT?.webrtcAudioPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSinks": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcAudioPrivate.getSinks);
    },
    "call_GetSinks": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webrtcAudioPrivate.getSinks();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSinks": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webrtcAudioPrivate.getSinks();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSinksChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcAudioPrivate?.onSinksChanged && "addListener" in WEBEXT?.webrtcAudioPrivate?.onSinksChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSinksChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcAudioPrivate.onSinksChanged.addListener);
    },
    "call_OnSinksChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webrtcAudioPrivate.onSinksChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnSinksChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webrtcAudioPrivate.onSinksChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSinksChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.webrtcAudioPrivate?.onSinksChanged &&
        "removeListener" in WEBEXT?.webrtcAudioPrivate?.onSinksChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSinksChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcAudioPrivate.onSinksChanged.removeListener);
    },
    "call_OffSinksChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webrtcAudioPrivate.onSinksChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffSinksChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webrtcAudioPrivate.onSinksChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSinksChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcAudioPrivate?.onSinksChanged && "hasListener" in WEBEXT?.webrtcAudioPrivate?.onSinksChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSinksChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcAudioPrivate.onSinksChanged.hasListener);
    },
    "call_HasOnSinksChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webrtcAudioPrivate.onSinksChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSinksChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webrtcAudioPrivate.onSinksChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
