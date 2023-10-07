import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/filebrowserhandler", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_FileHandlerExecuteEventDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["entries"]);
        A.store.Bool(ptr + 16, "tab_id" in x ? true : false);
        A.store.Int64(ptr + 8, x["tab_id"] === undefined ? 0 : (x["tab_id"] as number));
      }
    },
    "load_FileHandlerExecuteEventDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["entries"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["tab_id"] = A.load.Int64(ptr + 8);
      } else {
        delete x["tab_id"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_OnExecute": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileBrowserHandler?.onExecute && "addListener" in WEBEXT?.fileBrowserHandler?.onExecute) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnExecute": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileBrowserHandler.onExecute.addListener);
    },
    "call_OnExecute": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileBrowserHandler.onExecute.addListener(A.H.get<object>(callback));
    },
    "try_OnExecute": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileBrowserHandler.onExecute.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffExecute": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileBrowserHandler?.onExecute && "removeListener" in WEBEXT?.fileBrowserHandler?.onExecute) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffExecute": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileBrowserHandler.onExecute.removeListener);
    },
    "call_OffExecute": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileBrowserHandler.onExecute.removeListener(A.H.get<object>(callback));
    },
    "try_OffExecute": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileBrowserHandler.onExecute.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnExecute": (): heap.Ref<boolean> => {
      if (WEBEXT?.fileBrowserHandler?.onExecute && "hasListener" in WEBEXT?.fileBrowserHandler?.onExecute) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnExecute": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.fileBrowserHandler.onExecute.hasListener);
    },
    "call_HasOnExecute": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.fileBrowserHandler.onExecute.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnExecute": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.fileBrowserHandler.onExecute.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
