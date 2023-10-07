import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/idltest", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_GetArrayBuffer": (): heap.Ref<boolean> => {
      if (WEBEXT?.idltest && "getArrayBuffer" in WEBEXT?.idltest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetArrayBuffer": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.idltest.getArrayBuffer);
    },
    "call_GetArrayBuffer": (retPtr: Pointer): void => {
      const _ret = WEBEXT.idltest.getArrayBuffer();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetArrayBuffer": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.idltest.getArrayBuffer();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_NocompileFunc": (): heap.Ref<boolean> => {
      if (WEBEXT?.idltest && "nocompileFunc" in WEBEXT?.idltest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_NocompileFunc": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.idltest.nocompileFunc);
    },
    "call_NocompileFunc": (retPtr: Pointer, switch_: number): void => {
      const _ret = WEBEXT.idltest.nocompileFunc(switch_);
    },
    "try_NocompileFunc": (retPtr: Pointer, errPtr: Pointer, switch_: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.idltest.nocompileFunc(switch_);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendArrayBuffer": (): heap.Ref<boolean> => {
      if (WEBEXT?.idltest && "sendArrayBuffer" in WEBEXT?.idltest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendArrayBuffer": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.idltest.sendArrayBuffer);
    },
    "call_SendArrayBuffer": (retPtr: Pointer, input: heap.Ref<object>): void => {
      const _ret = WEBEXT.idltest.sendArrayBuffer(A.H.get<object>(input));
      A.store.Ref(retPtr, _ret);
    },
    "try_SendArrayBuffer": (retPtr: Pointer, errPtr: Pointer, input: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.idltest.sendArrayBuffer(A.H.get<object>(input));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendArrayBufferView": (): heap.Ref<boolean> => {
      if (WEBEXT?.idltest && "sendArrayBufferView" in WEBEXT?.idltest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendArrayBufferView": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.idltest.sendArrayBufferView);
    },
    "call_SendArrayBufferView": (retPtr: Pointer, input: heap.Ref<object>): void => {
      const _ret = WEBEXT.idltest.sendArrayBufferView(A.H.get<object>(input));
      A.store.Ref(retPtr, _ret);
    },
    "try_SendArrayBufferView": (retPtr: Pointer, errPtr: Pointer, input: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.idltest.sendArrayBufferView(A.H.get<object>(input));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
