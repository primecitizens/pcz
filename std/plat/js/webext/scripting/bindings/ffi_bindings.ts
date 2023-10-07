import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/scripting", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_InjectionTarget": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 12, false);
      } else {
        A.store.Bool(ptr + 15, true);
        A.store.Bool(ptr + 13, "tabId" in x ? true : false);
        A.store.Int32(ptr + 0, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Ref(ptr + 4, x["frameIds"]);
        A.store.Ref(ptr + 8, x["documentIds"]);
        A.store.Bool(ptr + 14, "allFrames" in x ? true : false);
        A.store.Bool(ptr + 12, x["allFrames"] ? true : false);
      }
    },
    "load_InjectionTarget": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 13)) {
        x["tabId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["tabId"];
      }
      x["frameIds"] = A.load.Ref(ptr + 4, undefined);
      x["documentIds"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 14)) {
        x["allFrames"] = A.load.Bool(ptr + 12);
      } else {
        delete x["allFrames"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_StyleOrigin": (ref: heap.Ref<string>): number => {
      const idx = ["AUTHOR", "USER"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_CSSInjection": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);

        A.store.Bool(ptr + 0 + 15, false);
        A.store.Bool(ptr + 0 + 13, false);
        A.store.Int32(ptr + 0 + 0, 0);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Ref(ptr + 0 + 8, undefined);
        A.store.Bool(ptr + 0 + 14, false);
        A.store.Bool(ptr + 0 + 12, false);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Enum(ptr + 24, -1);
      } else {
        A.store.Bool(ptr + 28, true);

        if (typeof x["target"] === "undefined") {
          A.store.Bool(ptr + 0 + 15, false);
          A.store.Bool(ptr + 0 + 13, false);
          A.store.Int32(ptr + 0 + 0, 0);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Ref(ptr + 0 + 8, undefined);
          A.store.Bool(ptr + 0 + 14, false);
          A.store.Bool(ptr + 0 + 12, false);
        } else {
          A.store.Bool(ptr + 0 + 15, true);
          A.store.Bool(ptr + 0 + 13, "tabId" in x["target"] ? true : false);
          A.store.Int32(ptr + 0 + 0, x["target"]["tabId"] === undefined ? 0 : (x["target"]["tabId"] as number));
          A.store.Ref(ptr + 0 + 4, x["target"]["frameIds"]);
          A.store.Ref(ptr + 0 + 8, x["target"]["documentIds"]);
          A.store.Bool(ptr + 0 + 14, "allFrames" in x["target"] ? true : false);
          A.store.Bool(ptr + 0 + 12, x["target"]["allFrames"] ? true : false);
        }
        A.store.Ref(ptr + 16, x["css"]);
        A.store.Ref(ptr + 20, x["files"]);
        A.store.Enum(ptr + 24, ["AUTHOR", "USER"].indexOf(x["origin"] as string));
      }
    },
    "load_CSSInjection": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 15)) {
        x["target"] = {};
        if (A.load.Bool(ptr + 0 + 13)) {
          x["target"]["tabId"] = A.load.Int32(ptr + 0 + 0);
        } else {
          delete x["target"]["tabId"];
        }
        x["target"]["frameIds"] = A.load.Ref(ptr + 0 + 4, undefined);
        x["target"]["documentIds"] = A.load.Ref(ptr + 0 + 8, undefined);
        if (A.load.Bool(ptr + 0 + 14)) {
          x["target"]["allFrames"] = A.load.Bool(ptr + 0 + 12);
        } else {
          delete x["target"]["allFrames"];
        }
      } else {
        delete x["target"];
      }
      x["css"] = A.load.Ref(ptr + 16, undefined);
      x["files"] = A.load.Ref(ptr + 20, undefined);
      x["origin"] = A.load.Enum(ptr + 24, ["AUTHOR", "USER"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ContentScriptFilter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["ids"]);
      }
    },
    "load_ContentScriptFilter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["ids"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ExecutionWorld": (ref: heap.Ref<string>): number => {
      const idx = ["ISOLATED", "MAIN"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RegisteredContentScript": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 39, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 36, false);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 37, false);
        A.store.Bool(ptr + 21, false);
        A.store.Enum(ptr + 24, -1);
        A.store.Bool(ptr + 38, false);
        A.store.Bool(ptr + 28, false);
        A.store.Enum(ptr + 32, -1);
      } else {
        A.store.Bool(ptr + 39, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["matches"]);
        A.store.Ref(ptr + 8, x["excludeMatches"]);
        A.store.Ref(ptr + 12, x["css"]);
        A.store.Ref(ptr + 16, x["js"]);
        A.store.Bool(ptr + 36, "allFrames" in x ? true : false);
        A.store.Bool(ptr + 20, x["allFrames"] ? true : false);
        A.store.Bool(ptr + 37, "matchOriginAsFallback" in x ? true : false);
        A.store.Bool(ptr + 21, x["matchOriginAsFallback"] ? true : false);
        A.store.Enum(ptr + 24, ["document_start", "document_end", "document_idle"].indexOf(x["runAt"] as string));
        A.store.Bool(ptr + 38, "persistAcrossSessions" in x ? true : false);
        A.store.Bool(ptr + 28, x["persistAcrossSessions"] ? true : false);
        A.store.Enum(ptr + 32, ["ISOLATED", "MAIN"].indexOf(x["world"] as string));
      }
    },
    "load_RegisteredContentScript": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["matches"] = A.load.Ref(ptr + 4, undefined);
      x["excludeMatches"] = A.load.Ref(ptr + 8, undefined);
      x["css"] = A.load.Ref(ptr + 12, undefined);
      x["js"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 36)) {
        x["allFrames"] = A.load.Bool(ptr + 20);
      } else {
        delete x["allFrames"];
      }
      if (A.load.Bool(ptr + 37)) {
        x["matchOriginAsFallback"] = A.load.Bool(ptr + 21);
      } else {
        delete x["matchOriginAsFallback"];
      }
      x["runAt"] = A.load.Enum(ptr + 24, ["document_start", "document_end", "document_idle"]);
      if (A.load.Bool(ptr + 38)) {
        x["persistAcrossSessions"] = A.load.Bool(ptr + 28);
      } else {
        delete x["persistAcrossSessions"];
      }
      x["world"] = A.load.Enum(ptr + 32, ["ISOLATED", "MAIN"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_InjectionResult": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Ref(ptr + 0, x["result"]);
        A.store.Bool(ptr + 12, "frameId" in x ? true : false);
        A.store.Int32(ptr + 4, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Ref(ptr + 8, x["documentId"]);
      }
    },
    "load_InjectionResult": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["result"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["frameId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["frameId"];
      }
      x["documentId"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "has_Properties_GlobalParams": (self: heap.Ref<any>): heap.Ref<boolean> => {
      if (WEBEXT?.scripting && "globalParams" in WEBEXT?.scripting) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Properties_GlobalParams": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.scripting.globalParams);
    },

    "call_Properties_GlobalParams": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = Reflect.apply(WEBEXT.scripting.globalParams, thiz, []);
      A.store.Int32(retPtr, _ret);
    },
    "try_Properties_GlobalParams": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = Reflect.apply(WEBEXT.scripting.globalParams, thiz, []);
        A.store.Int32(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },

    "store_ScriptInjection": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 38, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);

        A.store.Bool(ptr + 16 + 15, false);
        A.store.Bool(ptr + 16 + 13, false);
        A.store.Int32(ptr + 16 + 0, 0);
        A.store.Ref(ptr + 16 + 4, undefined);
        A.store.Ref(ptr + 16 + 8, undefined);
        A.store.Bool(ptr + 16 + 14, false);
        A.store.Bool(ptr + 16 + 12, false);
        A.store.Enum(ptr + 32, -1);
        A.store.Bool(ptr + 37, false);
        A.store.Bool(ptr + 36, false);
      } else {
        A.store.Bool(ptr + 38, true);
        A.store.Ref(ptr + 0, x["func"]);
        A.store.Ref(ptr + 4, x["args"]);
        A.store.Ref(ptr + 8, x["function"]);
        A.store.Ref(ptr + 12, x["files"]);

        if (typeof x["target"] === "undefined") {
          A.store.Bool(ptr + 16 + 15, false);
          A.store.Bool(ptr + 16 + 13, false);
          A.store.Int32(ptr + 16 + 0, 0);
          A.store.Ref(ptr + 16 + 4, undefined);
          A.store.Ref(ptr + 16 + 8, undefined);
          A.store.Bool(ptr + 16 + 14, false);
          A.store.Bool(ptr + 16 + 12, false);
        } else {
          A.store.Bool(ptr + 16 + 15, true);
          A.store.Bool(ptr + 16 + 13, "tabId" in x["target"] ? true : false);
          A.store.Int32(ptr + 16 + 0, x["target"]["tabId"] === undefined ? 0 : (x["target"]["tabId"] as number));
          A.store.Ref(ptr + 16 + 4, x["target"]["frameIds"]);
          A.store.Ref(ptr + 16 + 8, x["target"]["documentIds"]);
          A.store.Bool(ptr + 16 + 14, "allFrames" in x["target"] ? true : false);
          A.store.Bool(ptr + 16 + 12, x["target"]["allFrames"] ? true : false);
        }
        A.store.Enum(ptr + 32, ["ISOLATED", "MAIN"].indexOf(x["world"] as string));
        A.store.Bool(ptr + 37, "injectImmediately" in x ? true : false);
        A.store.Bool(ptr + 36, x["injectImmediately"] ? true : false);
      }
    },
    "load_ScriptInjection": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["func"] = A.load.Ref(ptr + 0, undefined);
      x["args"] = A.load.Ref(ptr + 4, undefined);
      x["function"] = A.load.Ref(ptr + 8, undefined);
      x["files"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 16 + 15)) {
        x["target"] = {};
        if (A.load.Bool(ptr + 16 + 13)) {
          x["target"]["tabId"] = A.load.Int32(ptr + 16 + 0);
        } else {
          delete x["target"]["tabId"];
        }
        x["target"]["frameIds"] = A.load.Ref(ptr + 16 + 4, undefined);
        x["target"]["documentIds"] = A.load.Ref(ptr + 16 + 8, undefined);
        if (A.load.Bool(ptr + 16 + 14)) {
          x["target"]["allFrames"] = A.load.Bool(ptr + 16 + 12);
        } else {
          delete x["target"]["allFrames"];
        }
      } else {
        delete x["target"];
      }
      x["world"] = A.load.Enum(ptr + 32, ["ISOLATED", "MAIN"]);
      if (A.load.Bool(ptr + 37)) {
        x["injectImmediately"] = A.load.Bool(ptr + 36);
      } else {
        delete x["injectImmediately"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_ExecuteScript": (): heap.Ref<boolean> => {
      if (WEBEXT?.scripting && "executeScript" in WEBEXT?.scripting) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ExecuteScript": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.scripting.executeScript);
    },
    "call_ExecuteScript": (retPtr: Pointer, injection: Pointer): void => {
      const injection_ffi = {};

      injection_ffi["func"] = A.load.Ref(injection + 0, undefined);
      injection_ffi["args"] = A.load.Ref(injection + 4, undefined);
      injection_ffi["function"] = A.load.Ref(injection + 8, undefined);
      injection_ffi["files"] = A.load.Ref(injection + 12, undefined);
      if (A.load.Bool(injection + 16 + 15)) {
        injection_ffi["target"] = {};
        if (A.load.Bool(injection + 16 + 13)) {
          injection_ffi["target"]["tabId"] = A.load.Int32(injection + 16 + 0);
        }
        injection_ffi["target"]["frameIds"] = A.load.Ref(injection + 16 + 4, undefined);
        injection_ffi["target"]["documentIds"] = A.load.Ref(injection + 16 + 8, undefined);
        if (A.load.Bool(injection + 16 + 14)) {
          injection_ffi["target"]["allFrames"] = A.load.Bool(injection + 16 + 12);
        }
      }
      injection_ffi["world"] = A.load.Enum(injection + 32, ["ISOLATED", "MAIN"]);
      if (A.load.Bool(injection + 37)) {
        injection_ffi["injectImmediately"] = A.load.Bool(injection + 36);
      }

      const _ret = WEBEXT.scripting.executeScript(injection_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ExecuteScript": (retPtr: Pointer, errPtr: Pointer, injection: Pointer): heap.Ref<boolean> => {
      try {
        const injection_ffi = {};

        injection_ffi["func"] = A.load.Ref(injection + 0, undefined);
        injection_ffi["args"] = A.load.Ref(injection + 4, undefined);
        injection_ffi["function"] = A.load.Ref(injection + 8, undefined);
        injection_ffi["files"] = A.load.Ref(injection + 12, undefined);
        if (A.load.Bool(injection + 16 + 15)) {
          injection_ffi["target"] = {};
          if (A.load.Bool(injection + 16 + 13)) {
            injection_ffi["target"]["tabId"] = A.load.Int32(injection + 16 + 0);
          }
          injection_ffi["target"]["frameIds"] = A.load.Ref(injection + 16 + 4, undefined);
          injection_ffi["target"]["documentIds"] = A.load.Ref(injection + 16 + 8, undefined);
          if (A.load.Bool(injection + 16 + 14)) {
            injection_ffi["target"]["allFrames"] = A.load.Bool(injection + 16 + 12);
          }
        }
        injection_ffi["world"] = A.load.Enum(injection + 32, ["ISOLATED", "MAIN"]);
        if (A.load.Bool(injection + 37)) {
          injection_ffi["injectImmediately"] = A.load.Bool(injection + 36);
        }

        const _ret = WEBEXT.scripting.executeScript(injection_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetRegisteredContentScripts": (): heap.Ref<boolean> => {
      if (WEBEXT?.scripting && "getRegisteredContentScripts" in WEBEXT?.scripting) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetRegisteredContentScripts": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.scripting.getRegisteredContentScripts);
    },
    "call_GetRegisteredContentScripts": (retPtr: Pointer, filter: Pointer): void => {
      const filter_ffi = {};

      filter_ffi["ids"] = A.load.Ref(filter + 0, undefined);

      const _ret = WEBEXT.scripting.getRegisteredContentScripts(filter_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetRegisteredContentScripts": (retPtr: Pointer, errPtr: Pointer, filter: Pointer): heap.Ref<boolean> => {
      try {
        const filter_ffi = {};

        filter_ffi["ids"] = A.load.Ref(filter + 0, undefined);

        const _ret = WEBEXT.scripting.getRegisteredContentScripts(filter_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InsertCSS": (): heap.Ref<boolean> => {
      if (WEBEXT?.scripting && "insertCSS" in WEBEXT?.scripting) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InsertCSS": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.scripting.insertCSS);
    },
    "call_InsertCSS": (retPtr: Pointer, injection: Pointer): void => {
      const injection_ffi = {};

      if (A.load.Bool(injection + 0 + 15)) {
        injection_ffi["target"] = {};
        if (A.load.Bool(injection + 0 + 13)) {
          injection_ffi["target"]["tabId"] = A.load.Int32(injection + 0 + 0);
        }
        injection_ffi["target"]["frameIds"] = A.load.Ref(injection + 0 + 4, undefined);
        injection_ffi["target"]["documentIds"] = A.load.Ref(injection + 0 + 8, undefined);
        if (A.load.Bool(injection + 0 + 14)) {
          injection_ffi["target"]["allFrames"] = A.load.Bool(injection + 0 + 12);
        }
      }
      injection_ffi["css"] = A.load.Ref(injection + 16, undefined);
      injection_ffi["files"] = A.load.Ref(injection + 20, undefined);
      injection_ffi["origin"] = A.load.Enum(injection + 24, ["AUTHOR", "USER"]);

      const _ret = WEBEXT.scripting.insertCSS(injection_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_InsertCSS": (retPtr: Pointer, errPtr: Pointer, injection: Pointer): heap.Ref<boolean> => {
      try {
        const injection_ffi = {};

        if (A.load.Bool(injection + 0 + 15)) {
          injection_ffi["target"] = {};
          if (A.load.Bool(injection + 0 + 13)) {
            injection_ffi["target"]["tabId"] = A.load.Int32(injection + 0 + 0);
          }
          injection_ffi["target"]["frameIds"] = A.load.Ref(injection + 0 + 4, undefined);
          injection_ffi["target"]["documentIds"] = A.load.Ref(injection + 0 + 8, undefined);
          if (A.load.Bool(injection + 0 + 14)) {
            injection_ffi["target"]["allFrames"] = A.load.Bool(injection + 0 + 12);
          }
        }
        injection_ffi["css"] = A.load.Ref(injection + 16, undefined);
        injection_ffi["files"] = A.load.Ref(injection + 20, undefined);
        injection_ffi["origin"] = A.load.Enum(injection + 24, ["AUTHOR", "USER"]);

        const _ret = WEBEXT.scripting.insertCSS(injection_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RegisterContentScripts": (): heap.Ref<boolean> => {
      if (WEBEXT?.scripting && "registerContentScripts" in WEBEXT?.scripting) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RegisterContentScripts": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.scripting.registerContentScripts);
    },
    "call_RegisterContentScripts": (retPtr: Pointer, scripts: heap.Ref<object>): void => {
      const _ret = WEBEXT.scripting.registerContentScripts(A.H.get<object>(scripts));
      A.store.Ref(retPtr, _ret);
    },
    "try_RegisterContentScripts": (retPtr: Pointer, errPtr: Pointer, scripts: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.scripting.registerContentScripts(A.H.get<object>(scripts));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveCSS": (): heap.Ref<boolean> => {
      if (WEBEXT?.scripting && "removeCSS" in WEBEXT?.scripting) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveCSS": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.scripting.removeCSS);
    },
    "call_RemoveCSS": (retPtr: Pointer, injection: Pointer): void => {
      const injection_ffi = {};

      if (A.load.Bool(injection + 0 + 15)) {
        injection_ffi["target"] = {};
        if (A.load.Bool(injection + 0 + 13)) {
          injection_ffi["target"]["tabId"] = A.load.Int32(injection + 0 + 0);
        }
        injection_ffi["target"]["frameIds"] = A.load.Ref(injection + 0 + 4, undefined);
        injection_ffi["target"]["documentIds"] = A.load.Ref(injection + 0 + 8, undefined);
        if (A.load.Bool(injection + 0 + 14)) {
          injection_ffi["target"]["allFrames"] = A.load.Bool(injection + 0 + 12);
        }
      }
      injection_ffi["css"] = A.load.Ref(injection + 16, undefined);
      injection_ffi["files"] = A.load.Ref(injection + 20, undefined);
      injection_ffi["origin"] = A.load.Enum(injection + 24, ["AUTHOR", "USER"]);

      const _ret = WEBEXT.scripting.removeCSS(injection_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveCSS": (retPtr: Pointer, errPtr: Pointer, injection: Pointer): heap.Ref<boolean> => {
      try {
        const injection_ffi = {};

        if (A.load.Bool(injection + 0 + 15)) {
          injection_ffi["target"] = {};
          if (A.load.Bool(injection + 0 + 13)) {
            injection_ffi["target"]["tabId"] = A.load.Int32(injection + 0 + 0);
          }
          injection_ffi["target"]["frameIds"] = A.load.Ref(injection + 0 + 4, undefined);
          injection_ffi["target"]["documentIds"] = A.load.Ref(injection + 0 + 8, undefined);
          if (A.load.Bool(injection + 0 + 14)) {
            injection_ffi["target"]["allFrames"] = A.load.Bool(injection + 0 + 12);
          }
        }
        injection_ffi["css"] = A.load.Ref(injection + 16, undefined);
        injection_ffi["files"] = A.load.Ref(injection + 20, undefined);
        injection_ffi["origin"] = A.load.Enum(injection + 24, ["AUTHOR", "USER"]);

        const _ret = WEBEXT.scripting.removeCSS(injection_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UnregisterContentScripts": (): heap.Ref<boolean> => {
      if (WEBEXT?.scripting && "unregisterContentScripts" in WEBEXT?.scripting) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UnregisterContentScripts": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.scripting.unregisterContentScripts);
    },
    "call_UnregisterContentScripts": (retPtr: Pointer, filter: Pointer): void => {
      const filter_ffi = {};

      filter_ffi["ids"] = A.load.Ref(filter + 0, undefined);

      const _ret = WEBEXT.scripting.unregisterContentScripts(filter_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_UnregisterContentScripts": (retPtr: Pointer, errPtr: Pointer, filter: Pointer): heap.Ref<boolean> => {
      try {
        const filter_ffi = {};

        filter_ffi["ids"] = A.load.Ref(filter + 0, undefined);

        const _ret = WEBEXT.scripting.unregisterContentScripts(filter_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateContentScripts": (): heap.Ref<boolean> => {
      if (WEBEXT?.scripting && "updateContentScripts" in WEBEXT?.scripting) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateContentScripts": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.scripting.updateContentScripts);
    },
    "call_UpdateContentScripts": (retPtr: Pointer, scripts: heap.Ref<object>): void => {
      const _ret = WEBEXT.scripting.updateContentScripts(A.H.get<object>(scripts));
      A.store.Ref(retPtr, _ret);
    },
    "try_UpdateContentScripts": (retPtr: Pointer, errPtr: Pointer, scripts: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.scripting.updateContentScripts(A.H.get<object>(scripts));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
