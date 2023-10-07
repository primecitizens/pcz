import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/app/runtime", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_ActionType": (ref: heap.Ref<string>): number => {
      const idx = ["new_note"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ActionData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 5, false);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["new_note"].indexOf(x["actionType"] as string));
        A.store.Bool(ptr + 6, "isLockScreenAction" in x ? true : false);
        A.store.Bool(ptr + 4, x["isLockScreenAction"] ? true : false);
        A.store.Bool(ptr + 7, "restoreLastActionState" in x ? true : false);
        A.store.Bool(ptr + 5, x["restoreLastActionState"] ? true : false);
      }
    },
    "load_ActionData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["actionType"] = A.load.Enum(ptr + 0, ["new_note"]);
      if (A.load.Bool(ptr + 6)) {
        x["isLockScreenAction"] = A.load.Bool(ptr + 4);
      } else {
        delete x["isLockScreenAction"];
      }
      if (A.load.Bool(ptr + 7)) {
        x["restoreLastActionState"] = A.load.Bool(ptr + 5);
      } else {
        delete x["restoreLastActionState"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "has_EmbedRequest_Allow": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "allow" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EmbedRequest_Allow": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).allow);
    },

    "call_EmbedRequest_Allow": (self: heap.Ref<object>, retPtr: Pointer, url: heap.Ref<object>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.allow(A.H.get<object>(url));
    },
    "try_EmbedRequest_Allow": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      url: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.allow(A.H.get<object>(url));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_EmbedRequest_Deny": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "deny" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EmbedRequest_Deny": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).deny);
    },

    "call_EmbedRequest_Deny": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.deny();
    },
    "try_EmbedRequest_Deny": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.deny();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_EmbedRequest_EmbedderId": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "embedderId" in thiz) {
        const val = thiz.embedderId;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_EmbedRequest_EmbedderId": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "embedderId", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_EmbedRequest_Data": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "data" in thiz) {
        const val = thiz.data;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_EmbedRequest_Data": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "data", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },

    "store_LaunchItem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["entry"]);
        A.store.Ref(ptr + 4, x["type"]);
      }
    },
    "load_LaunchItem": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["entry"] = A.load.Ref(ptr + 0, undefined);
      x["type"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_LaunchSource": (ref: heap.Ref<string>): number => {
      const idx = [
        "untracked",
        "app_launcher",
        "new_tab_page",
        "reload",
        "restart",
        "load_and_launch",
        "command_line",
        "file_handler",
        "url_handler",
        "system_tray",
        "about_page",
        "keyboard",
        "extensions_page",
        "management_api",
        "ephemeral_app",
        "background",
        "kiosk",
        "chrome_internal",
        "test",
        "installed_notification",
        "context_menu",
        "arc",
        "intent_url",
        "app_home_page",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_LaunchData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 36, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 33, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 34, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 35, false);
        A.store.Bool(ptr + 18, false);
        A.store.Enum(ptr + 20, -1);

        A.store.Bool(ptr + 24 + 8, false);
        A.store.Enum(ptr + 24 + 0, -1);
        A.store.Bool(ptr + 24 + 6, false);
        A.store.Bool(ptr + 24 + 4, false);
        A.store.Bool(ptr + 24 + 7, false);
        A.store.Bool(ptr + 24 + 5, false);
      } else {
        A.store.Bool(ptr + 36, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["items"]);
        A.store.Ref(ptr + 8, x["url"]);
        A.store.Ref(ptr + 12, x["referrerUrl"]);
        A.store.Bool(ptr + 33, "isDemoSession" in x ? true : false);
        A.store.Bool(ptr + 16, x["isDemoSession"] ? true : false);
        A.store.Bool(ptr + 34, "isKioskSession" in x ? true : false);
        A.store.Bool(ptr + 17, x["isKioskSession"] ? true : false);
        A.store.Bool(ptr + 35, "isPublicSession" in x ? true : false);
        A.store.Bool(ptr + 18, x["isPublicSession"] ? true : false);
        A.store.Enum(
          ptr + 20,
          [
            "untracked",
            "app_launcher",
            "new_tab_page",
            "reload",
            "restart",
            "load_and_launch",
            "command_line",
            "file_handler",
            "url_handler",
            "system_tray",
            "about_page",
            "keyboard",
            "extensions_page",
            "management_api",
            "ephemeral_app",
            "background",
            "kiosk",
            "chrome_internal",
            "test",
            "installed_notification",
            "context_menu",
            "arc",
            "intent_url",
            "app_home_page",
          ].indexOf(x["source"] as string)
        );

        if (typeof x["actionData"] === "undefined") {
          A.store.Bool(ptr + 24 + 8, false);
          A.store.Enum(ptr + 24 + 0, -1);
          A.store.Bool(ptr + 24 + 6, false);
          A.store.Bool(ptr + 24 + 4, false);
          A.store.Bool(ptr + 24 + 7, false);
          A.store.Bool(ptr + 24 + 5, false);
        } else {
          A.store.Bool(ptr + 24 + 8, true);
          A.store.Enum(ptr + 24 + 0, ["new_note"].indexOf(x["actionData"]["actionType"] as string));
          A.store.Bool(ptr + 24 + 6, "isLockScreenAction" in x["actionData"] ? true : false);
          A.store.Bool(ptr + 24 + 4, x["actionData"]["isLockScreenAction"] ? true : false);
          A.store.Bool(ptr + 24 + 7, "restoreLastActionState" in x["actionData"] ? true : false);
          A.store.Bool(ptr + 24 + 5, x["actionData"]["restoreLastActionState"] ? true : false);
        }
      }
    },
    "load_LaunchData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["items"] = A.load.Ref(ptr + 4, undefined);
      x["url"] = A.load.Ref(ptr + 8, undefined);
      x["referrerUrl"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 33)) {
        x["isDemoSession"] = A.load.Bool(ptr + 16);
      } else {
        delete x["isDemoSession"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["isKioskSession"] = A.load.Bool(ptr + 17);
      } else {
        delete x["isKioskSession"];
      }
      if (A.load.Bool(ptr + 35)) {
        x["isPublicSession"] = A.load.Bool(ptr + 18);
      } else {
        delete x["isPublicSession"];
      }
      x["source"] = A.load.Enum(ptr + 20, [
        "untracked",
        "app_launcher",
        "new_tab_page",
        "reload",
        "restart",
        "load_and_launch",
        "command_line",
        "file_handler",
        "url_handler",
        "system_tray",
        "about_page",
        "keyboard",
        "extensions_page",
        "management_api",
        "ephemeral_app",
        "background",
        "kiosk",
        "chrome_internal",
        "test",
        "installed_notification",
        "context_menu",
        "arc",
        "intent_url",
        "app_home_page",
      ]);
      if (A.load.Bool(ptr + 24 + 8)) {
        x["actionData"] = {};
        x["actionData"]["actionType"] = A.load.Enum(ptr + 24 + 0, ["new_note"]);
        if (A.load.Bool(ptr + 24 + 6)) {
          x["actionData"]["isLockScreenAction"] = A.load.Bool(ptr + 24 + 4);
        } else {
          delete x["actionData"]["isLockScreenAction"];
        }
        if (A.load.Bool(ptr + 24 + 7)) {
          x["actionData"]["restoreLastActionState"] = A.load.Bool(ptr + 24 + 5);
        } else {
          delete x["actionData"]["restoreLastActionState"];
        }
      } else {
        delete x["actionData"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_OnEmbedRequested": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.runtime?.onEmbedRequested && "addListener" in WEBEXT?.app?.runtime?.onEmbedRequested) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnEmbedRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.runtime.onEmbedRequested.addListener);
    },
    "call_OnEmbedRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.runtime.onEmbedRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnEmbedRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.runtime.onEmbedRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffEmbedRequested": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.runtime?.onEmbedRequested && "removeListener" in WEBEXT?.app?.runtime?.onEmbedRequested) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffEmbedRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.runtime.onEmbedRequested.removeListener);
    },
    "call_OffEmbedRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.runtime.onEmbedRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffEmbedRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.runtime.onEmbedRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnEmbedRequested": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.runtime?.onEmbedRequested && "hasListener" in WEBEXT?.app?.runtime?.onEmbedRequested) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnEmbedRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.runtime.onEmbedRequested.hasListener);
    },
    "call_HasOnEmbedRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.runtime.onEmbedRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnEmbedRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.runtime.onEmbedRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnLaunched": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.runtime?.onLaunched && "addListener" in WEBEXT?.app?.runtime?.onLaunched) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnLaunched": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.runtime.onLaunched.addListener);
    },
    "call_OnLaunched": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.runtime.onLaunched.addListener(A.H.get<object>(callback));
    },
    "try_OnLaunched": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.runtime.onLaunched.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffLaunched": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.runtime?.onLaunched && "removeListener" in WEBEXT?.app?.runtime?.onLaunched) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffLaunched": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.runtime.onLaunched.removeListener);
    },
    "call_OffLaunched": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.runtime.onLaunched.removeListener(A.H.get<object>(callback));
    },
    "try_OffLaunched": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.runtime.onLaunched.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnLaunched": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.runtime?.onLaunched && "hasListener" in WEBEXT?.app?.runtime?.onLaunched) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnLaunched": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.runtime.onLaunched.hasListener);
    },
    "call_HasOnLaunched": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.runtime.onLaunched.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnLaunched": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.runtime.onLaunched.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRestarted": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.runtime?.onRestarted && "addListener" in WEBEXT?.app?.runtime?.onRestarted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRestarted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.runtime.onRestarted.addListener);
    },
    "call_OnRestarted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.runtime.onRestarted.addListener(A.H.get<object>(callback));
    },
    "try_OnRestarted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.runtime.onRestarted.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRestarted": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.runtime?.onRestarted && "removeListener" in WEBEXT?.app?.runtime?.onRestarted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRestarted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.runtime.onRestarted.removeListener);
    },
    "call_OffRestarted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.runtime.onRestarted.removeListener(A.H.get<object>(callback));
    },
    "try_OffRestarted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.runtime.onRestarted.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRestarted": (): heap.Ref<boolean> => {
      if (WEBEXT?.app?.runtime?.onRestarted && "hasListener" in WEBEXT?.app?.runtime?.onRestarted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRestarted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.app.runtime.onRestarted.hasListener);
    },
    "call_HasOnRestarted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.app.runtime.onRestarted.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRestarted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.app.runtime.onRestarted.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
