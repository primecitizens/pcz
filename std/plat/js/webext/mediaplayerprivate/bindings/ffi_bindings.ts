import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/mediaplayerprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_OnNextTrack": (): heap.Ref<boolean> => {
      if (WEBEXT?.mediaPlayerPrivate?.onNextTrack && "addListener" in WEBEXT?.mediaPlayerPrivate?.onNextTrack) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnNextTrack": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPlayerPrivate.onNextTrack.addListener);
    },
    "call_OnNextTrack": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mediaPlayerPrivate.onNextTrack.addListener(A.H.get<object>(callback));
    },
    "try_OnNextTrack": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mediaPlayerPrivate.onNextTrack.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffNextTrack": (): heap.Ref<boolean> => {
      if (WEBEXT?.mediaPlayerPrivate?.onNextTrack && "removeListener" in WEBEXT?.mediaPlayerPrivate?.onNextTrack) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffNextTrack": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPlayerPrivate.onNextTrack.removeListener);
    },
    "call_OffNextTrack": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mediaPlayerPrivate.onNextTrack.removeListener(A.H.get<object>(callback));
    },
    "try_OffNextTrack": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mediaPlayerPrivate.onNextTrack.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnNextTrack": (): heap.Ref<boolean> => {
      if (WEBEXT?.mediaPlayerPrivate?.onNextTrack && "hasListener" in WEBEXT?.mediaPlayerPrivate?.onNextTrack) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnNextTrack": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPlayerPrivate.onNextTrack.hasListener);
    },
    "call_HasOnNextTrack": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mediaPlayerPrivate.onNextTrack.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnNextTrack": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mediaPlayerPrivate.onNextTrack.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPrevTrack": (): heap.Ref<boolean> => {
      if (WEBEXT?.mediaPlayerPrivate?.onPrevTrack && "addListener" in WEBEXT?.mediaPlayerPrivate?.onPrevTrack) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPrevTrack": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPlayerPrivate.onPrevTrack.addListener);
    },
    "call_OnPrevTrack": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mediaPlayerPrivate.onPrevTrack.addListener(A.H.get<object>(callback));
    },
    "try_OnPrevTrack": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mediaPlayerPrivate.onPrevTrack.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPrevTrack": (): heap.Ref<boolean> => {
      if (WEBEXT?.mediaPlayerPrivate?.onPrevTrack && "removeListener" in WEBEXT?.mediaPlayerPrivate?.onPrevTrack) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPrevTrack": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPlayerPrivate.onPrevTrack.removeListener);
    },
    "call_OffPrevTrack": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mediaPlayerPrivate.onPrevTrack.removeListener(A.H.get<object>(callback));
    },
    "try_OffPrevTrack": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mediaPlayerPrivate.onPrevTrack.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPrevTrack": (): heap.Ref<boolean> => {
      if (WEBEXT?.mediaPlayerPrivate?.onPrevTrack && "hasListener" in WEBEXT?.mediaPlayerPrivate?.onPrevTrack) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPrevTrack": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPlayerPrivate.onPrevTrack.hasListener);
    },
    "call_HasOnPrevTrack": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mediaPlayerPrivate.onPrevTrack.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPrevTrack": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mediaPlayerPrivate.onPrevTrack.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTogglePlayState": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.mediaPlayerPrivate?.onTogglePlayState &&
        "addListener" in WEBEXT?.mediaPlayerPrivate?.onTogglePlayState
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTogglePlayState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPlayerPrivate.onTogglePlayState.addListener);
    },
    "call_OnTogglePlayState": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mediaPlayerPrivate.onTogglePlayState.addListener(A.H.get<object>(callback));
    },
    "try_OnTogglePlayState": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mediaPlayerPrivate.onTogglePlayState.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffTogglePlayState": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.mediaPlayerPrivate?.onTogglePlayState &&
        "removeListener" in WEBEXT?.mediaPlayerPrivate?.onTogglePlayState
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffTogglePlayState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPlayerPrivate.onTogglePlayState.removeListener);
    },
    "call_OffTogglePlayState": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mediaPlayerPrivate.onTogglePlayState.removeListener(A.H.get<object>(callback));
    },
    "try_OffTogglePlayState": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mediaPlayerPrivate.onTogglePlayState.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnTogglePlayState": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.mediaPlayerPrivate?.onTogglePlayState &&
        "hasListener" in WEBEXT?.mediaPlayerPrivate?.onTogglePlayState
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnTogglePlayState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPlayerPrivate.onTogglePlayState.hasListener);
    },
    "call_HasOnTogglePlayState": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mediaPlayerPrivate.onTogglePlayState.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnTogglePlayState": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mediaPlayerPrivate.onTogglePlayState.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
