import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/clipboard", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_DataItemType": (ref: heap.Ref<string>): number => {
      const idx = ["textPlain", "textHtml"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_AdditionalDataItem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["textPlain", "textHtml"].indexOf(x["type"] as string));
        A.store.Ref(ptr + 4, x["data"]);
      }
    },
    "load_AdditionalDataItem": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, ["textPlain", "textHtml"]);
      x["data"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ImageType": (ref: heap.Ref<string>): number => {
      const idx = ["png", "jpeg"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_OnClipboardDataChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.clipboard?.onClipboardDataChanged && "addListener" in WEBEXT?.clipboard?.onClipboardDataChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClipboardDataChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.clipboard.onClipboardDataChanged.addListener);
    },
    "call_OnClipboardDataChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.clipboard.onClipboardDataChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnClipboardDataChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.clipboard.onClipboardDataChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClipboardDataChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.clipboard?.onClipboardDataChanged && "removeListener" in WEBEXT?.clipboard?.onClipboardDataChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClipboardDataChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.clipboard.onClipboardDataChanged.removeListener);
    },
    "call_OffClipboardDataChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.clipboard.onClipboardDataChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffClipboardDataChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.clipboard.onClipboardDataChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClipboardDataChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.clipboard?.onClipboardDataChanged && "hasListener" in WEBEXT?.clipboard?.onClipboardDataChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClipboardDataChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.clipboard.onClipboardDataChanged.hasListener);
    },
    "call_HasOnClipboardDataChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.clipboard.onClipboardDataChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClipboardDataChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.clipboard.onClipboardDataChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetImageData": (): heap.Ref<boolean> => {
      if (WEBEXT?.clipboard && "setImageData" in WEBEXT?.clipboard) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetImageData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.clipboard.setImageData);
    },
    "call_SetImageData": (
      retPtr: Pointer,
      imageData: heap.Ref<object>,
      type: number,
      additionalItems: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.clipboard.setImageData(
        A.H.get<object>(imageData),
        type > 0 && type <= 2 ? ["png", "jpeg"][type - 1] : undefined,
        A.H.get<object>(additionalItems)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_SetImageData": (
      retPtr: Pointer,
      errPtr: Pointer,
      imageData: heap.Ref<object>,
      type: number,
      additionalItems: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.clipboard.setImageData(
          A.H.get<object>(imageData),
          type > 0 && type <= 2 ? ["png", "jpeg"][type - 1] : undefined,
          A.H.get<object>(additionalItems)
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
