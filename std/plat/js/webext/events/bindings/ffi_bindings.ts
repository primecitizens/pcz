import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/events", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Rule": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 29, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 28, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 29, true);
        A.store.Ref(ptr + 0, x["actions"]);
        A.store.Ref(ptr + 4, x["conditions"]);
        A.store.Ref(ptr + 8, x["id"]);
        A.store.Bool(ptr + 28, "priority" in x ? true : false);
        A.store.Int64(ptr + 16, x["priority"] === undefined ? 0 : (x["priority"] as number));
        A.store.Ref(ptr + 24, x["tags"]);
      }
    },
    "load_Rule": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["actions"] = A.load.Ref(ptr + 0, undefined);
      x["conditions"] = A.load.Ref(ptr + 4, undefined);
      x["id"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 28)) {
        x["priority"] = A.load.Int64(ptr + 16);
      } else {
        delete x["priority"];
      }
      x["tags"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "has_Event_AddListener": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "addListener" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Event_AddListener": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).addListener);
    },

    "call_Event_AddListener": (self: heap.Ref<object>, retPtr: Pointer, callback: heap.Ref<object>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.addListener(A.H.get<object>(callback));
    },
    "try_Event_AddListener": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Event_RemoveListener": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "removeListener" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Event_RemoveListener": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).removeListener);
    },

    "call_Event_RemoveListener": (self: heap.Ref<object>, retPtr: Pointer, callback: heap.Ref<object>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.removeListener(A.H.get<object>(callback));
    },
    "try_Event_RemoveListener": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Event_HasListener": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "hasListener" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Event_HasListener": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).hasListener);
    },

    "call_Event_HasListener": (self: heap.Ref<object>, retPtr: Pointer, callback: heap.Ref<object>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_Event_HasListener": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Event_HasListeners": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "hasListeners" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Event_HasListeners": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).hasListeners);
    },

    "call_Event_HasListeners": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.hasListeners();
      A.store.Bool(retPtr, _ret);
    },
    "try_Event_HasListeners": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.hasListeners();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Event_AddRules": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "addRules" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Event_AddRules": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).addRules);
    },

    "call_Event_AddRules": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      eventName: heap.Ref<object>,
      webViewInstanceId: number,
      rules: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.addRules(
        A.H.get<object>(eventName),
        webViewInstanceId,
        A.H.get<object>(rules),
        A.H.get<object>(callback)
      );
    },
    "try_Event_AddRules": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      eventName: heap.Ref<object>,
      webViewInstanceId: number,
      rules: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.addRules(
          A.H.get<object>(eventName),
          webViewInstanceId,
          A.H.get<object>(rules),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Event_AddRules1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "addRules" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Event_AddRules1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).addRules);
    },

    "call_Event_AddRules1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      eventName: heap.Ref<object>,
      webViewInstanceId: number,
      rules: heap.Ref<object>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.addRules(A.H.get<object>(eventName), webViewInstanceId, A.H.get<object>(rules));
    },
    "try_Event_AddRules1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      eventName: heap.Ref<object>,
      webViewInstanceId: number,
      rules: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.addRules(A.H.get<object>(eventName), webViewInstanceId, A.H.get<object>(rules));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Event_GetRules": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "getRules" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Event_GetRules": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).getRules);
    },

    "call_Event_GetRules": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      eventName: heap.Ref<object>,
      webViewInstanceId: number,
      ruleIdentifiers: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.getRules(
        A.H.get<object>(eventName),
        webViewInstanceId,
        A.H.get<object>(ruleIdentifiers),
        A.H.get<object>(callback)
      );
    },
    "try_Event_GetRules": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      eventName: heap.Ref<object>,
      webViewInstanceId: number,
      ruleIdentifiers: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.getRules(
          A.H.get<object>(eventName),
          webViewInstanceId,
          A.H.get<object>(ruleIdentifiers),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Event_GetRules1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "getRules" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Event_GetRules1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).getRules);
    },

    "call_Event_GetRules1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      eventName: heap.Ref<object>,
      webViewInstanceId: number
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.getRules(A.H.get<object>(eventName), webViewInstanceId);
    },
    "try_Event_GetRules1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      eventName: heap.Ref<object>,
      webViewInstanceId: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.getRules(A.H.get<object>(eventName), webViewInstanceId);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Event_RemoveRules": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "removeRules" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Event_RemoveRules": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).removeRules);
    },

    "call_Event_RemoveRules": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      eventName: heap.Ref<object>,
      webViewInstanceId: number,
      ruleIdentifiers: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.removeRules(
        A.H.get<object>(eventName),
        webViewInstanceId,
        A.H.get<object>(ruleIdentifiers),
        A.H.get<object>(callback)
      );
    },
    "try_Event_RemoveRules": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      eventName: heap.Ref<object>,
      webViewInstanceId: number,
      ruleIdentifiers: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.removeRules(
          A.H.get<object>(eventName),
          webViewInstanceId,
          A.H.get<object>(ruleIdentifiers),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Event_RemoveRules1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "removeRules" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Event_RemoveRules1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).removeRules);
    },

    "call_Event_RemoveRules1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      eventName: heap.Ref<object>,
      webViewInstanceId: number,
      ruleIdentifiers: heap.Ref<object>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.removeRules(A.H.get<object>(eventName), webViewInstanceId, A.H.get<object>(ruleIdentifiers));
    },
    "try_Event_RemoveRules1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      eventName: heap.Ref<object>,
      webViewInstanceId: number,
      ruleIdentifiers: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.removeRules(A.H.get<object>(eventName), webViewInstanceId, A.H.get<object>(ruleIdentifiers));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Event_RemoveRules2": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "removeRules" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Event_RemoveRules2": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).removeRules);
    },

    "call_Event_RemoveRules2": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      eventName: heap.Ref<object>,
      webViewInstanceId: number
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.removeRules(A.H.get<object>(eventName), webViewInstanceId);
    },
    "try_Event_RemoveRules2": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      eventName: heap.Ref<object>,
      webViewInstanceId: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.removeRules(A.H.get<object>(eventName), webViewInstanceId);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },

    "store_UrlFilter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 80, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Ref(ptr + 48, undefined);
        A.store.Ref(ptr + 52, undefined);
        A.store.Ref(ptr + 56, undefined);
        A.store.Ref(ptr + 60, undefined);
        A.store.Ref(ptr + 64, undefined);
        A.store.Ref(ptr + 68, undefined);
        A.store.Ref(ptr + 72, undefined);
        A.store.Ref(ptr + 76, undefined);
      } else {
        A.store.Bool(ptr + 80, true);
        A.store.Ref(ptr + 0, x["hostContains"]);
        A.store.Ref(ptr + 4, x["hostEquals"]);
        A.store.Ref(ptr + 8, x["hostPrefix"]);
        A.store.Ref(ptr + 12, x["hostSuffix"]);
        A.store.Ref(ptr + 16, x["originAndPathMatches"]);
        A.store.Ref(ptr + 20, x["pathContains"]);
        A.store.Ref(ptr + 24, x["pathEquals"]);
        A.store.Ref(ptr + 28, x["pathPrefix"]);
        A.store.Ref(ptr + 32, x["pathSuffix"]);
        A.store.Ref(ptr + 36, x["ports"]);
        A.store.Ref(ptr + 40, x["queryContains"]);
        A.store.Ref(ptr + 44, x["queryEquals"]);
        A.store.Ref(ptr + 48, x["queryPrefix"]);
        A.store.Ref(ptr + 52, x["querySuffix"]);
        A.store.Ref(ptr + 56, x["schemes"]);
        A.store.Ref(ptr + 60, x["urlContains"]);
        A.store.Ref(ptr + 64, x["urlEquals"]);
        A.store.Ref(ptr + 68, x["urlMatches"]);
        A.store.Ref(ptr + 72, x["urlPrefix"]);
        A.store.Ref(ptr + 76, x["urlSuffix"]);
      }
    },
    "load_UrlFilter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["hostContains"] = A.load.Ref(ptr + 0, undefined);
      x["hostEquals"] = A.load.Ref(ptr + 4, undefined);
      x["hostPrefix"] = A.load.Ref(ptr + 8, undefined);
      x["hostSuffix"] = A.load.Ref(ptr + 12, undefined);
      x["originAndPathMatches"] = A.load.Ref(ptr + 16, undefined);
      x["pathContains"] = A.load.Ref(ptr + 20, undefined);
      x["pathEquals"] = A.load.Ref(ptr + 24, undefined);
      x["pathPrefix"] = A.load.Ref(ptr + 28, undefined);
      x["pathSuffix"] = A.load.Ref(ptr + 32, undefined);
      x["ports"] = A.load.Ref(ptr + 36, undefined);
      x["queryContains"] = A.load.Ref(ptr + 40, undefined);
      x["queryEquals"] = A.load.Ref(ptr + 44, undefined);
      x["queryPrefix"] = A.load.Ref(ptr + 48, undefined);
      x["querySuffix"] = A.load.Ref(ptr + 52, undefined);
      x["schemes"] = A.load.Ref(ptr + 56, undefined);
      x["urlContains"] = A.load.Ref(ptr + 60, undefined);
      x["urlEquals"] = A.load.Ref(ptr + 64, undefined);
      x["urlMatches"] = A.load.Ref(ptr + 68, undefined);
      x["urlPrefix"] = A.load.Ref(ptr + 72, undefined);
      x["urlSuffix"] = A.load.Ref(ptr + 76, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
  };
});
