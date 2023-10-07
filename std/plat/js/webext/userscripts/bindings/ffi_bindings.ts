import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/userscripts", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_ScriptSource": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["code"]);
        A.store.Ref(ptr + 4, x["file"]);
      }
    },
    "load_ScriptSource": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["code"] = A.load.Ref(ptr + 0, undefined);
      x["file"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RegisteredUserScript": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Enum(ptr + 20, -1);
      } else {
        A.store.Bool(ptr + 25, true);
        A.store.Bool(ptr + 24, "allFrames" in x ? true : false);
        A.store.Bool(ptr + 0, x["allFrames"] ? true : false);
        A.store.Ref(ptr + 4, x["excludeMatches"]);
        A.store.Ref(ptr + 8, x["id"]);
        A.store.Ref(ptr + 12, x["js"]);
        A.store.Ref(ptr + 16, x["matches"]);
        A.store.Enum(ptr + 20, ["document_start", "document_end", "document_idle"].indexOf(x["runAt"] as string));
      }
    },
    "load_RegisteredUserScript": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["allFrames"] = A.load.Bool(ptr + 0);
      } else {
        delete x["allFrames"];
      }
      x["excludeMatches"] = A.load.Ref(ptr + 4, undefined);
      x["id"] = A.load.Ref(ptr + 8, undefined);
      x["js"] = A.load.Ref(ptr + 12, undefined);
      x["matches"] = A.load.Ref(ptr + 16, undefined);
      x["runAt"] = A.load.Enum(ptr + 20, ["document_start", "document_end", "document_idle"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UserScriptFilter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["ids"]);
      }
    },
    "load_UserScriptFilter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["ids"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetScripts": (): heap.Ref<boolean> => {
      if (WEBEXT?.userScripts && "getScripts" in WEBEXT?.userScripts) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetScripts": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.userScripts.getScripts);
    },
    "call_GetScripts": (retPtr: Pointer, filter: Pointer): void => {
      const filter_ffi = {};

      filter_ffi["ids"] = A.load.Ref(filter + 0, undefined);

      const _ret = WEBEXT.userScripts.getScripts(filter_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetScripts": (retPtr: Pointer, errPtr: Pointer, filter: Pointer): heap.Ref<boolean> => {
      try {
        const filter_ffi = {};

        filter_ffi["ids"] = A.load.Ref(filter + 0, undefined);

        const _ret = WEBEXT.userScripts.getScripts(filter_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Register": (): heap.Ref<boolean> => {
      if (WEBEXT?.userScripts && "register" in WEBEXT?.userScripts) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Register": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.userScripts.register);
    },
    "call_Register": (retPtr: Pointer, scripts: heap.Ref<object>): void => {
      const _ret = WEBEXT.userScripts.register(A.H.get<object>(scripts));
      A.store.Ref(retPtr, _ret);
    },
    "try_Register": (retPtr: Pointer, errPtr: Pointer, scripts: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.userScripts.register(A.H.get<object>(scripts));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Unregister": (): heap.Ref<boolean> => {
      if (WEBEXT?.userScripts && "unregister" in WEBEXT?.userScripts) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Unregister": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.userScripts.unregister);
    },
    "call_Unregister": (retPtr: Pointer, filter: Pointer): void => {
      const filter_ffi = {};

      filter_ffi["ids"] = A.load.Ref(filter + 0, undefined);

      const _ret = WEBEXT.userScripts.unregister(filter_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Unregister": (retPtr: Pointer, errPtr: Pointer, filter: Pointer): heap.Ref<boolean> => {
      try {
        const filter_ffi = {};

        filter_ffi["ids"] = A.load.Ref(filter + 0, undefined);

        const _ret = WEBEXT.userScripts.unregister(filter_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
