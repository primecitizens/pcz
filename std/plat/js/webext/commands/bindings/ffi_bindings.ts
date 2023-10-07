import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/commands", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Command": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["description"]);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Ref(ptr + 8, x["shortcut"]);
      }
    },
    "load_Command": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["description"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      x["shortcut"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetAll": (): heap.Ref<boolean> => {
      if (WEBEXT?.commands && "getAll" in WEBEXT?.commands) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.commands.getAll);
    },
    "call_GetAll": (retPtr: Pointer): void => {
      const _ret = WEBEXT.commands.getAll();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAll": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.commands.getAll();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCommand": (): heap.Ref<boolean> => {
      if (WEBEXT?.commands?.onCommand && "addListener" in WEBEXT?.commands?.onCommand) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCommand": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.commands.onCommand.addListener);
    },
    "call_OnCommand": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.commands.onCommand.addListener(A.H.get<object>(callback));
    },
    "try_OnCommand": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.commands.onCommand.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCommand": (): heap.Ref<boolean> => {
      if (WEBEXT?.commands?.onCommand && "removeListener" in WEBEXT?.commands?.onCommand) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCommand": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.commands.onCommand.removeListener);
    },
    "call_OffCommand": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.commands.onCommand.removeListener(A.H.get<object>(callback));
    },
    "try_OffCommand": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.commands.onCommand.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCommand": (): heap.Ref<boolean> => {
      if (WEBEXT?.commands?.onCommand && "hasListener" in WEBEXT?.commands?.onCommand) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCommand": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.commands.onCommand.hasListener);
    },
    "call_HasOnCommand": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.commands.onCommand.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCommand": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.commands.onCommand.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
