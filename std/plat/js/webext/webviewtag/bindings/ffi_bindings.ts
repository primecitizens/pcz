import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/webviewtag", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AttachArgWebview": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_AttachArgWebview": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ClearDataOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Float64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "since" in x ? true : false);
        A.store.Float64(ptr + 0, x["since"] === undefined ? 0 : (x["since"] as number));
      }
    },
    "load_ClearDataOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["since"] = A.load.Float64(ptr + 0);
      } else {
        delete x["since"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ClearDataTypeSet": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Bool(ptr + 9, "appcache" in x ? true : false);
        A.store.Bool(ptr + 0, x["appcache"] ? true : false);
        A.store.Bool(ptr + 10, "cache" in x ? true : false);
        A.store.Bool(ptr + 1, x["cache"] ? true : false);
        A.store.Bool(ptr + 11, "cookies" in x ? true : false);
        A.store.Bool(ptr + 2, x["cookies"] ? true : false);
        A.store.Bool(ptr + 12, "fileSystems" in x ? true : false);
        A.store.Bool(ptr + 3, x["fileSystems"] ? true : false);
        A.store.Bool(ptr + 13, "indexedDB" in x ? true : false);
        A.store.Bool(ptr + 4, x["indexedDB"] ? true : false);
        A.store.Bool(ptr + 14, "localStorage" in x ? true : false);
        A.store.Bool(ptr + 5, x["localStorage"] ? true : false);
        A.store.Bool(ptr + 15, "persistentCookies" in x ? true : false);
        A.store.Bool(ptr + 6, x["persistentCookies"] ? true : false);
        A.store.Bool(ptr + 16, "sessionCookies" in x ? true : false);
        A.store.Bool(ptr + 7, x["sessionCookies"] ? true : false);
        A.store.Bool(ptr + 17, "webSQL" in x ? true : false);
        A.store.Bool(ptr + 8, x["webSQL"] ? true : false);
      }
    },
    "load_ClearDataTypeSet": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 9)) {
        x["appcache"] = A.load.Bool(ptr + 0);
      } else {
        delete x["appcache"];
      }
      if (A.load.Bool(ptr + 10)) {
        x["cache"] = A.load.Bool(ptr + 1);
      } else {
        delete x["cache"];
      }
      if (A.load.Bool(ptr + 11)) {
        x["cookies"] = A.load.Bool(ptr + 2);
      } else {
        delete x["cookies"];
      }
      if (A.load.Bool(ptr + 12)) {
        x["fileSystems"] = A.load.Bool(ptr + 3);
      } else {
        delete x["fileSystems"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["indexedDB"] = A.load.Bool(ptr + 4);
      } else {
        delete x["indexedDB"];
      }
      if (A.load.Bool(ptr + 14)) {
        x["localStorage"] = A.load.Bool(ptr + 5);
      } else {
        delete x["localStorage"];
      }
      if (A.load.Bool(ptr + 15)) {
        x["persistentCookies"] = A.load.Bool(ptr + 6);
      } else {
        delete x["persistentCookies"];
      }
      if (A.load.Bool(ptr + 16)) {
        x["sessionCookies"] = A.load.Bool(ptr + 7);
      } else {
        delete x["sessionCookies"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["webSQL"] = A.load.Bool(ptr + 8);
      } else {
        delete x["webSQL"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_InjectionItems": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["code"]);
        A.store.Ref(ptr + 4, x["files"]);
      }
    },
    "load_InjectionItems": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["code"] = A.load.Ref(ptr + 0, undefined);
      x["files"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ContentScriptDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 54, false);
        A.store.Bool(ptr + 52, false);
        A.store.Bool(ptr + 0, false);

        A.store.Bool(ptr + 4 + 8, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Ref(ptr + 4 + 4, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);

        A.store.Bool(ptr + 28 + 8, false);
        A.store.Ref(ptr + 28 + 0, undefined);
        A.store.Ref(ptr + 28 + 4, undefined);
        A.store.Bool(ptr + 53, false);
        A.store.Bool(ptr + 37, false);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Enum(ptr + 48, -1);
      } else {
        A.store.Bool(ptr + 54, true);
        A.store.Bool(ptr + 52, "all_frames" in x ? true : false);
        A.store.Bool(ptr + 0, x["all_frames"] ? true : false);

        if (typeof x["css"] === "undefined") {
          A.store.Bool(ptr + 4 + 8, false);
          A.store.Ref(ptr + 4 + 0, undefined);
          A.store.Ref(ptr + 4 + 4, undefined);
        } else {
          A.store.Bool(ptr + 4 + 8, true);
          A.store.Ref(ptr + 4 + 0, x["css"]["code"]);
          A.store.Ref(ptr + 4 + 4, x["css"]["files"]);
        }
        A.store.Ref(ptr + 16, x["exclude_globs"]);
        A.store.Ref(ptr + 20, x["exclude_matches"]);
        A.store.Ref(ptr + 24, x["include_globs"]);

        if (typeof x["js"] === "undefined") {
          A.store.Bool(ptr + 28 + 8, false);
          A.store.Ref(ptr + 28 + 0, undefined);
          A.store.Ref(ptr + 28 + 4, undefined);
        } else {
          A.store.Bool(ptr + 28 + 8, true);
          A.store.Ref(ptr + 28 + 0, x["js"]["code"]);
          A.store.Ref(ptr + 28 + 4, x["js"]["files"]);
        }
        A.store.Bool(ptr + 53, "match_about_blank" in x ? true : false);
        A.store.Bool(ptr + 37, x["match_about_blank"] ? true : false);
        A.store.Ref(ptr + 40, x["matches"]);
        A.store.Ref(ptr + 44, x["name"]);
        A.store.Enum(ptr + 48, ["document_start", "document_end", "document_idle"].indexOf(x["run_at"] as string));
      }
    },
    "load_ContentScriptDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 52)) {
        x["all_frames"] = A.load.Bool(ptr + 0);
      } else {
        delete x["all_frames"];
      }
      if (A.load.Bool(ptr + 4 + 8)) {
        x["css"] = {};
        x["css"]["code"] = A.load.Ref(ptr + 4 + 0, undefined);
        x["css"]["files"] = A.load.Ref(ptr + 4 + 4, undefined);
      } else {
        delete x["css"];
      }
      x["exclude_globs"] = A.load.Ref(ptr + 16, undefined);
      x["exclude_matches"] = A.load.Ref(ptr + 20, undefined);
      x["include_globs"] = A.load.Ref(ptr + 24, undefined);
      if (A.load.Bool(ptr + 28 + 8)) {
        x["js"] = {};
        x["js"]["code"] = A.load.Ref(ptr + 28 + 0, undefined);
        x["js"]["files"] = A.load.Ref(ptr + 28 + 4, undefined);
      } else {
        delete x["js"];
      }
      if (A.load.Bool(ptr + 53)) {
        x["match_about_blank"] = A.load.Bool(ptr + 37);
      } else {
        delete x["match_about_blank"];
      }
      x["matches"] = A.load.Ref(ptr + 40, undefined);
      x["name"] = A.load.Ref(ptr + 44, undefined);
      x["run_at"] = A.load.Enum(ptr + 48, ["document_start", "document_end", "document_idle"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "has_ContentWindowType_PostMessage": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "postMessage" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContentWindowType_PostMessage": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).postMessage);
    },

    "call_ContentWindowType_PostMessage": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      message: heap.Ref<object>,
      targetOrigin: heap.Ref<object>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.postMessage(A.H.get<object>(message), A.H.get<object>(targetOrigin));
    },
    "try_ContentWindowType_PostMessage": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      message: heap.Ref<object>,
      targetOrigin: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.postMessage(A.H.get<object>(message), A.H.get<object>(targetOrigin));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "constof_ContextType": (ref: heap.Ref<string>): number => {
      const idx = ["all", "page", "frame", "selection", "link", "editable", "image", "video", "audio"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ContextMenuCreateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 42, false);
        A.store.Bool(ptr + 40, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 41, false);
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Enum(ptr + 36, -1);
      } else {
        A.store.Bool(ptr + 42, true);
        A.store.Bool(ptr + 40, "checked" in x ? true : false);
        A.store.Bool(ptr + 0, x["checked"] ? true : false);
        A.store.Ref(ptr + 4, x["contexts"]);
        A.store.Ref(ptr + 8, x["documentUrlPatterns"]);
        A.store.Bool(ptr + 41, "enabled" in x ? true : false);
        A.store.Bool(ptr + 12, x["enabled"] ? true : false);
        A.store.Ref(ptr + 16, x["id"]);
        A.store.Ref(ptr + 20, x["onclick"]);
        A.store.Ref(ptr + 24, x["parentId"]);
        A.store.Ref(ptr + 28, x["targetUrlPatterns"]);
        A.store.Ref(ptr + 32, x["title"]);
        A.store.Enum(ptr + 36, ["normal", "checkbox", "radio", "separator"].indexOf(x["type"] as string));
      }
    },
    "load_ContextMenuCreateProperties": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 40)) {
        x["checked"] = A.load.Bool(ptr + 0);
      } else {
        delete x["checked"];
      }
      x["contexts"] = A.load.Ref(ptr + 4, undefined);
      x["documentUrlPatterns"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 41)) {
        x["enabled"] = A.load.Bool(ptr + 12);
      } else {
        delete x["enabled"];
      }
      x["id"] = A.load.Ref(ptr + 16, undefined);
      x["onclick"] = A.load.Ref(ptr + 20, undefined);
      x["parentId"] = A.load.Ref(ptr + 24, undefined);
      x["targetUrlPatterns"] = A.load.Ref(ptr + 28, undefined);
      x["title"] = A.load.Ref(ptr + 32, undefined);
      x["type"] = A.load.Enum(ptr + 36, ["normal", "checkbox", "radio", "separator"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ContextMenuUpdateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 38, false);
        A.store.Bool(ptr + 36, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 37, false);
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Enum(ptr + 32, -1);
      } else {
        A.store.Bool(ptr + 38, true);
        A.store.Bool(ptr + 36, "checked" in x ? true : false);
        A.store.Bool(ptr + 0, x["checked"] ? true : false);
        A.store.Ref(ptr + 4, x["contexts"]);
        A.store.Ref(ptr + 8, x["documentUrlPatterns"]);
        A.store.Bool(ptr + 37, "enabled" in x ? true : false);
        A.store.Bool(ptr + 12, x["enabled"] ? true : false);
        A.store.Ref(ptr + 16, x["onclick"]);
        A.store.Ref(ptr + 20, x["parentId"]);
        A.store.Ref(ptr + 24, x["targetUrlPatterns"]);
        A.store.Ref(ptr + 28, x["title"]);
        A.store.Enum(ptr + 32, ["normal", "checkbox", "radio", "separator"].indexOf(x["type"] as string));
      }
    },
    "load_ContextMenuUpdateProperties": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 36)) {
        x["checked"] = A.load.Bool(ptr + 0);
      } else {
        delete x["checked"];
      }
      x["contexts"] = A.load.Ref(ptr + 4, undefined);
      x["documentUrlPatterns"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 37)) {
        x["enabled"] = A.load.Bool(ptr + 12);
      } else {
        delete x["enabled"];
      }
      x["onclick"] = A.load.Ref(ptr + 16, undefined);
      x["parentId"] = A.load.Ref(ptr + 20, undefined);
      x["targetUrlPatterns"] = A.load.Ref(ptr + 24, undefined);
      x["title"] = A.load.Ref(ptr + 28, undefined);
      x["type"] = A.load.Enum(ptr + 32, ["normal", "checkbox", "radio", "separator"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "has_ContextMenusType_Create": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "create" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContextMenusType_Create": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).create);
    },

    "call_ContextMenusType_Create": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      createProperties: Pointer,
      callback: heap.Ref<object>
    ): void => {
      const thiz = A.H.get<any>(self);

      const createProperties_ffi = {};

      if (A.load.Bool(createProperties + 40)) {
        createProperties_ffi["checked"] = A.load.Bool(createProperties + 0);
      }
      createProperties_ffi["contexts"] = A.load.Ref(createProperties + 4, undefined);
      createProperties_ffi["documentUrlPatterns"] = A.load.Ref(createProperties + 8, undefined);
      if (A.load.Bool(createProperties + 41)) {
        createProperties_ffi["enabled"] = A.load.Bool(createProperties + 12);
      }
      createProperties_ffi["id"] = A.load.Ref(createProperties + 16, undefined);
      createProperties_ffi["onclick"] = A.load.Ref(createProperties + 20, undefined);
      createProperties_ffi["parentId"] = A.load.Ref(createProperties + 24, undefined);
      createProperties_ffi["targetUrlPatterns"] = A.load.Ref(createProperties + 28, undefined);
      createProperties_ffi["title"] = A.load.Ref(createProperties + 32, undefined);
      createProperties_ffi["type"] = A.load.Enum(createProperties + 36, ["normal", "checkbox", "radio", "separator"]);

      const _ret = thiz.create(createProperties_ffi, A.H.get<object>(callback));
      A.store.Ref(retPtr, _ret);
    },
    "try_ContextMenusType_Create": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      createProperties: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const createProperties_ffi = {};

        if (A.load.Bool(createProperties + 40)) {
          createProperties_ffi["checked"] = A.load.Bool(createProperties + 0);
        }
        createProperties_ffi["contexts"] = A.load.Ref(createProperties + 4, undefined);
        createProperties_ffi["documentUrlPatterns"] = A.load.Ref(createProperties + 8, undefined);
        if (A.load.Bool(createProperties + 41)) {
          createProperties_ffi["enabled"] = A.load.Bool(createProperties + 12);
        }
        createProperties_ffi["id"] = A.load.Ref(createProperties + 16, undefined);
        createProperties_ffi["onclick"] = A.load.Ref(createProperties + 20, undefined);
        createProperties_ffi["parentId"] = A.load.Ref(createProperties + 24, undefined);
        createProperties_ffi["targetUrlPatterns"] = A.load.Ref(createProperties + 28, undefined);
        createProperties_ffi["title"] = A.load.Ref(createProperties + 32, undefined);
        createProperties_ffi["type"] = A.load.Enum(createProperties + 36, ["normal", "checkbox", "radio", "separator"]);

        const _ret = thiz.create(createProperties_ffi, A.H.get<object>(callback));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ContextMenusType_Create1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "create" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContextMenusType_Create1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).create);
    },

    "call_ContextMenusType_Create1": (self: heap.Ref<object>, retPtr: Pointer, createProperties: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const createProperties_ffi = {};

      if (A.load.Bool(createProperties + 40)) {
        createProperties_ffi["checked"] = A.load.Bool(createProperties + 0);
      }
      createProperties_ffi["contexts"] = A.load.Ref(createProperties + 4, undefined);
      createProperties_ffi["documentUrlPatterns"] = A.load.Ref(createProperties + 8, undefined);
      if (A.load.Bool(createProperties + 41)) {
        createProperties_ffi["enabled"] = A.load.Bool(createProperties + 12);
      }
      createProperties_ffi["id"] = A.load.Ref(createProperties + 16, undefined);
      createProperties_ffi["onclick"] = A.load.Ref(createProperties + 20, undefined);
      createProperties_ffi["parentId"] = A.load.Ref(createProperties + 24, undefined);
      createProperties_ffi["targetUrlPatterns"] = A.load.Ref(createProperties + 28, undefined);
      createProperties_ffi["title"] = A.load.Ref(createProperties + 32, undefined);
      createProperties_ffi["type"] = A.load.Enum(createProperties + 36, ["normal", "checkbox", "radio", "separator"]);
      const _ret = thiz.create(createProperties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ContextMenusType_Create1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      createProperties: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const createProperties_ffi = {};

        if (A.load.Bool(createProperties + 40)) {
          createProperties_ffi["checked"] = A.load.Bool(createProperties + 0);
        }
        createProperties_ffi["contexts"] = A.load.Ref(createProperties + 4, undefined);
        createProperties_ffi["documentUrlPatterns"] = A.load.Ref(createProperties + 8, undefined);
        if (A.load.Bool(createProperties + 41)) {
          createProperties_ffi["enabled"] = A.load.Bool(createProperties + 12);
        }
        createProperties_ffi["id"] = A.load.Ref(createProperties + 16, undefined);
        createProperties_ffi["onclick"] = A.load.Ref(createProperties + 20, undefined);
        createProperties_ffi["parentId"] = A.load.Ref(createProperties + 24, undefined);
        createProperties_ffi["targetUrlPatterns"] = A.load.Ref(createProperties + 28, undefined);
        createProperties_ffi["title"] = A.load.Ref(createProperties + 32, undefined);
        createProperties_ffi["type"] = A.load.Enum(createProperties + 36, ["normal", "checkbox", "radio", "separator"]);
        const _ret = thiz.create(createProperties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ContextMenusType_Update": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "update" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContextMenusType_Update": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).update);
    },

    "call_ContextMenusType_Update": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      id: heap.Ref<any>,
      updateProperties: Pointer,
      callback: heap.Ref<object>
    ): void => {
      const thiz = A.H.get<any>(self);

      const updateProperties_ffi = {};

      if (A.load.Bool(updateProperties + 36)) {
        updateProperties_ffi["checked"] = A.load.Bool(updateProperties + 0);
      }
      updateProperties_ffi["contexts"] = A.load.Ref(updateProperties + 4, undefined);
      updateProperties_ffi["documentUrlPatterns"] = A.load.Ref(updateProperties + 8, undefined);
      if (A.load.Bool(updateProperties + 37)) {
        updateProperties_ffi["enabled"] = A.load.Bool(updateProperties + 12);
      }
      updateProperties_ffi["onclick"] = A.load.Ref(updateProperties + 16, undefined);
      updateProperties_ffi["parentId"] = A.load.Ref(updateProperties + 20, undefined);
      updateProperties_ffi["targetUrlPatterns"] = A.load.Ref(updateProperties + 24, undefined);
      updateProperties_ffi["title"] = A.load.Ref(updateProperties + 28, undefined);
      updateProperties_ffi["type"] = A.load.Enum(updateProperties + 32, ["normal", "checkbox", "radio", "separator"]);

      const _ret = thiz.update(A.H.get<any>(id), updateProperties_ffi, A.H.get<object>(callback));
    },
    "try_ContextMenusType_Update": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<any>,
      updateProperties: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const updateProperties_ffi = {};

        if (A.load.Bool(updateProperties + 36)) {
          updateProperties_ffi["checked"] = A.load.Bool(updateProperties + 0);
        }
        updateProperties_ffi["contexts"] = A.load.Ref(updateProperties + 4, undefined);
        updateProperties_ffi["documentUrlPatterns"] = A.load.Ref(updateProperties + 8, undefined);
        if (A.load.Bool(updateProperties + 37)) {
          updateProperties_ffi["enabled"] = A.load.Bool(updateProperties + 12);
        }
        updateProperties_ffi["onclick"] = A.load.Ref(updateProperties + 16, undefined);
        updateProperties_ffi["parentId"] = A.load.Ref(updateProperties + 20, undefined);
        updateProperties_ffi["targetUrlPatterns"] = A.load.Ref(updateProperties + 24, undefined);
        updateProperties_ffi["title"] = A.load.Ref(updateProperties + 28, undefined);
        updateProperties_ffi["type"] = A.load.Enum(updateProperties + 32, ["normal", "checkbox", "radio", "separator"]);

        const _ret = thiz.update(A.H.get<any>(id), updateProperties_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ContextMenusType_Update1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "update" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContextMenusType_Update1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).update);
    },

    "call_ContextMenusType_Update1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      id: heap.Ref<any>,
      updateProperties: Pointer
    ): void => {
      const thiz = A.H.get<any>(self);

      const updateProperties_ffi = {};

      if (A.load.Bool(updateProperties + 36)) {
        updateProperties_ffi["checked"] = A.load.Bool(updateProperties + 0);
      }
      updateProperties_ffi["contexts"] = A.load.Ref(updateProperties + 4, undefined);
      updateProperties_ffi["documentUrlPatterns"] = A.load.Ref(updateProperties + 8, undefined);
      if (A.load.Bool(updateProperties + 37)) {
        updateProperties_ffi["enabled"] = A.load.Bool(updateProperties + 12);
      }
      updateProperties_ffi["onclick"] = A.load.Ref(updateProperties + 16, undefined);
      updateProperties_ffi["parentId"] = A.load.Ref(updateProperties + 20, undefined);
      updateProperties_ffi["targetUrlPatterns"] = A.load.Ref(updateProperties + 24, undefined);
      updateProperties_ffi["title"] = A.load.Ref(updateProperties + 28, undefined);
      updateProperties_ffi["type"] = A.load.Enum(updateProperties + 32, ["normal", "checkbox", "radio", "separator"]);
      const _ret = thiz.update(A.H.get<any>(id), updateProperties_ffi);
    },
    "try_ContextMenusType_Update1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<any>,
      updateProperties: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const updateProperties_ffi = {};

        if (A.load.Bool(updateProperties + 36)) {
          updateProperties_ffi["checked"] = A.load.Bool(updateProperties + 0);
        }
        updateProperties_ffi["contexts"] = A.load.Ref(updateProperties + 4, undefined);
        updateProperties_ffi["documentUrlPatterns"] = A.load.Ref(updateProperties + 8, undefined);
        if (A.load.Bool(updateProperties + 37)) {
          updateProperties_ffi["enabled"] = A.load.Bool(updateProperties + 12);
        }
        updateProperties_ffi["onclick"] = A.load.Ref(updateProperties + 16, undefined);
        updateProperties_ffi["parentId"] = A.load.Ref(updateProperties + 20, undefined);
        updateProperties_ffi["targetUrlPatterns"] = A.load.Ref(updateProperties + 24, undefined);
        updateProperties_ffi["title"] = A.load.Ref(updateProperties + 28, undefined);
        updateProperties_ffi["type"] = A.load.Enum(updateProperties + 32, ["normal", "checkbox", "radio", "separator"]);
        const _ret = thiz.update(A.H.get<any>(id), updateProperties_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ContextMenusType_Remove": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "remove" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContextMenusType_Remove": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).remove);
    },

    "call_ContextMenusType_Remove": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      menuItemId: heap.Ref<any>,
      callback: heap.Ref<object>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.remove(A.H.get<any>(menuItemId), A.H.get<object>(callback));
    },
    "try_ContextMenusType_Remove": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      menuItemId: heap.Ref<any>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.remove(A.H.get<any>(menuItemId), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ContextMenusType_Remove1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "remove" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContextMenusType_Remove1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).remove);
    },

    "call_ContextMenusType_Remove1": (self: heap.Ref<object>, retPtr: Pointer, menuItemId: heap.Ref<any>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.remove(A.H.get<any>(menuItemId));
    },
    "try_ContextMenusType_Remove1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      menuItemId: heap.Ref<any>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.remove(A.H.get<any>(menuItemId));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ContextMenusType_RemoveAll": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "removeAll" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContextMenusType_RemoveAll": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).removeAll);
    },

    "call_ContextMenusType_RemoveAll": (self: heap.Ref<object>, retPtr: Pointer, callback: heap.Ref<object>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.removeAll(A.H.get<object>(callback));
    },
    "try_ContextMenusType_RemoveAll": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.removeAll(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ContextMenusType_RemoveAll1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "removeAll" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContextMenusType_RemoveAll1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).removeAll);
    },

    "call_ContextMenusType_RemoveAll1": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.removeAll();
    },
    "try_ContextMenusType_RemoveAll1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.removeAll();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "constof_DialogArgMessageType": (ref: heap.Ref<string>): number => {
      const idx = ["alert", "confirm", "prompt"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "has_DialogController_Ok": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "ok" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DialogController_Ok": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).ok);
    },

    "call_DialogController_Ok": (self: heap.Ref<object>, retPtr: Pointer, response: heap.Ref<object>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.ok(A.H.get<object>(response));
    },
    "try_DialogController_Ok": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      response: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.ok(A.H.get<object>(response));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DialogController_Ok1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "ok" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DialogController_Ok1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).ok);
    },

    "call_DialogController_Ok1": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.ok();
    },
    "try_DialogController_Ok1": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.ok();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DialogController_Cancel": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "cancel" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DialogController_Cancel": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).cancel);
    },

    "call_DialogController_Cancel": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.cancel();
    },
    "try_DialogController_Cancel": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.cancel();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },

    "has_DownloadPermissionRequest_Allow": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "allow" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DownloadPermissionRequest_Allow": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).allow);
    },

    "call_DownloadPermissionRequest_Allow": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.allow();
    },
    "try_DownloadPermissionRequest_Allow": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.allow();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DownloadPermissionRequest_Deny": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "deny" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DownloadPermissionRequest_Deny": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).deny);
    },

    "call_DownloadPermissionRequest_Deny": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.deny();
    },
    "try_DownloadPermissionRequest_Deny": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.deny();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_DownloadPermissionRequest_RequestMethod": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "requestMethod" in thiz) {
        const val = thiz.requestMethod;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_DownloadPermissionRequest_RequestMethod": (
      self: heap.Ref<object>,
      val: heap.Ref<object>
    ): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "requestMethod", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_DownloadPermissionRequest_Url": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "url" in thiz) {
        const val = thiz.url;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_DownloadPermissionRequest_Url": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "url", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "constof_ExitArgReason": (ref: heap.Ref<string>): number => {
      const idx = ["normal", "abnormal", "crash", "kill"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "has_FileSystemPermissionRequest_Allow": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "allow" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_FileSystemPermissionRequest_Allow": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).allow);
    },

    "call_FileSystemPermissionRequest_Allow": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.allow();
    },
    "try_FileSystemPermissionRequest_Allow": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.allow();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_FileSystemPermissionRequest_Deny": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "deny" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_FileSystemPermissionRequest_Deny": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).deny);
    },

    "call_FileSystemPermissionRequest_Deny": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.deny();
    },
    "try_FileSystemPermissionRequest_Deny": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.deny();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_FileSystemPermissionRequest_Url": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "url" in thiz) {
        const val = thiz.url;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_FileSystemPermissionRequest_Url": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "url", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },

    "store_SelectionRect": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 32, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
        A.store.Int64(ptr + 16, 0);
        A.store.Int64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 32, true);
        A.store.Int64(ptr + 0, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Int64(ptr + 8, x["left"] === undefined ? 0 : (x["left"] as number));
        A.store.Int64(ptr + 16, x["top"] === undefined ? 0 : (x["top"] as number));
        A.store.Int64(ptr + 24, x["width"] === undefined ? 0 : (x["width"] as number));
      }
    },
    "load_SelectionRect": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["height"] = A.load.Int64(ptr + 0);
      x["left"] = A.load.Int64(ptr + 8);
      x["top"] = A.load.Int64(ptr + 16);
      x["width"] = A.load.Int64(ptr + 24);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_FindCallbackResults": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 57, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 16, 0);

        A.store.Bool(ptr + 24 + 32, false);
        A.store.Int64(ptr + 24 + 0, 0);
        A.store.Int64(ptr + 24 + 8, 0);
        A.store.Int64(ptr + 24 + 16, 0);
        A.store.Int64(ptr + 24 + 24, 0);
      } else {
        A.store.Bool(ptr + 57, true);
        A.store.Int64(ptr + 0, x["activeMatchOrdinal"] === undefined ? 0 : (x["activeMatchOrdinal"] as number));
        A.store.Bool(ptr + 8, x["canceled"] ? true : false);
        A.store.Int64(ptr + 16, x["numberOfMatches"] === undefined ? 0 : (x["numberOfMatches"] as number));

        if (typeof x["selectionRect"] === "undefined") {
          A.store.Bool(ptr + 24 + 32, false);
          A.store.Int64(ptr + 24 + 0, 0);
          A.store.Int64(ptr + 24 + 8, 0);
          A.store.Int64(ptr + 24 + 16, 0);
          A.store.Int64(ptr + 24 + 24, 0);
        } else {
          A.store.Bool(ptr + 24 + 32, true);
          A.store.Int64(
            ptr + 24 + 0,
            x["selectionRect"]["height"] === undefined ? 0 : (x["selectionRect"]["height"] as number)
          );
          A.store.Int64(
            ptr + 24 + 8,
            x["selectionRect"]["left"] === undefined ? 0 : (x["selectionRect"]["left"] as number)
          );
          A.store.Int64(
            ptr + 24 + 16,
            x["selectionRect"]["top"] === undefined ? 0 : (x["selectionRect"]["top"] as number)
          );
          A.store.Int64(
            ptr + 24 + 24,
            x["selectionRect"]["width"] === undefined ? 0 : (x["selectionRect"]["width"] as number)
          );
        }
      }
    },
    "load_FindCallbackResults": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["activeMatchOrdinal"] = A.load.Int64(ptr + 0);
      x["canceled"] = A.load.Bool(ptr + 8);
      x["numberOfMatches"] = A.load.Int64(ptr + 16);
      if (A.load.Bool(ptr + 24 + 32)) {
        x["selectionRect"] = {};
        x["selectionRect"]["height"] = A.load.Int64(ptr + 24 + 0);
        x["selectionRect"]["left"] = A.load.Int64(ptr + 24 + 8);
        x["selectionRect"]["top"] = A.load.Int64(ptr + 24 + 16);
        x["selectionRect"]["width"] = A.load.Int64(ptr + 24 + 24);
      } else {
        delete x["selectionRect"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_FindOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 1, false);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Bool(ptr + 2, "backward" in x ? true : false);
        A.store.Bool(ptr + 0, x["backward"] ? true : false);
        A.store.Bool(ptr + 3, "matchCase" in x ? true : false);
        A.store.Bool(ptr + 1, x["matchCase"] ? true : false);
      }
    },
    "load_FindOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 2)) {
        x["backward"] = A.load.Bool(ptr + 0);
      } else {
        delete x["backward"];
      }
      if (A.load.Bool(ptr + 3)) {
        x["matchCase"] = A.load.Bool(ptr + 1);
      } else {
        delete x["matchCase"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "has_FullscreenPermissionRequest_Allow": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "allow" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_FullscreenPermissionRequest_Allow": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).allow);
    },

    "call_FullscreenPermissionRequest_Allow": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.allow();
    },
    "try_FullscreenPermissionRequest_Allow": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.allow();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_FullscreenPermissionRequest_Deny": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "deny" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_FullscreenPermissionRequest_Deny": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).deny);
    },

    "call_FullscreenPermissionRequest_Deny": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.deny();
    },
    "try_FullscreenPermissionRequest_Deny": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.deny();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_FullscreenPermissionRequest_Origin": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "origin" in thiz) {
        const val = thiz.origin;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_FullscreenPermissionRequest_Origin": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "origin", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },

    "has_GeolocationPermissionRequest_Allow": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "allow" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GeolocationPermissionRequest_Allow": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).allow);
    },

    "call_GeolocationPermissionRequest_Allow": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.allow();
    },
    "try_GeolocationPermissionRequest_Allow": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.allow();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GeolocationPermissionRequest_Deny": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "deny" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GeolocationPermissionRequest_Deny": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).deny);
    },

    "call_GeolocationPermissionRequest_Deny": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.deny();
    },
    "try_GeolocationPermissionRequest_Deny": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.deny();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_GeolocationPermissionRequest_Url": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "url" in thiz) {
        const val = thiz.url;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_GeolocationPermissionRequest_Url": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "url", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "constof_ZoomMode": (ref: heap.Ref<string>): number => {
      const idx = ["per-origin", "per-view", "disabled"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_InjectDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_InjectDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["code"] = A.load.Ref(ptr + 0, undefined);
      x["file"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "has_LoadPluginPermissionRequest_Allow": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "allow" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LoadPluginPermissionRequest_Allow": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).allow);
    },

    "call_LoadPluginPermissionRequest_Allow": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.allow();
    },
    "try_LoadPluginPermissionRequest_Allow": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.allow();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LoadPluginPermissionRequest_Deny": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "deny" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LoadPluginPermissionRequest_Deny": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).deny);
    },

    "call_LoadPluginPermissionRequest_Deny": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.deny();
    },
    "try_LoadPluginPermissionRequest_Deny": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.deny();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_LoadPluginPermissionRequest_Identifier": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "identifier" in thiz) {
        const val = thiz.identifier;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_LoadPluginPermissionRequest_Identifier": (
      self: heap.Ref<object>,
      val: heap.Ref<object>
    ): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "identifier", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_LoadPluginPermissionRequest_Name": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "name" in thiz) {
        const val = thiz.name;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_LoadPluginPermissionRequest_Name": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "name", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "constof_LoadabortArgReason": (ref: heap.Ref<string>): number => {
      const idx = [
        "ERR_ABORTED",
        "ERR_INVALID_URL",
        "ERR_DISALLOWED_URL_SCHEME",
        "ERR_BLOCKED_BY_CLIENT",
        "ERR_ADDRESS_UNREACHABLE",
        "ERR_EMPTY_RESPONSE",
        "ERR_FILE_NOT_FOUND",
        "ERR_UNKNOWN_URL_SCHEME",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "has_MediaPermissionRequest_Allow": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "allow" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MediaPermissionRequest_Allow": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).allow);
    },

    "call_MediaPermissionRequest_Allow": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.allow();
    },
    "try_MediaPermissionRequest_Allow": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.allow();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_MediaPermissionRequest_Deny": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "deny" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MediaPermissionRequest_Deny": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).deny);
    },

    "call_MediaPermissionRequest_Deny": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.deny();
    },
    "try_MediaPermissionRequest_Deny": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.deny();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_MediaPermissionRequest_Url": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "url" in thiz) {
        const val = thiz.url;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_MediaPermissionRequest_Url": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "url", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },

    "has_NewWindow_Attach": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "attach" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_NewWindow_Attach": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).attach);
    },

    "call_NewWindow_Attach": (self: heap.Ref<object>, retPtr: Pointer, webview: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const webview_ffi = {};

      const _ret = thiz.attach(webview_ffi);
    },
    "try_NewWindow_Attach": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      webview: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const webview_ffi = {};

        const _ret = thiz.attach(webview_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_NewWindow_Discard": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "discard" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_NewWindow_Discard": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).discard);
    },

    "call_NewWindow_Discard": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.discard();
    },
    "try_NewWindow_Discard": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.discard();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "constof_NewwindowArgWindowOpenDisposition": (ref: heap.Ref<string>): number => {
      const idx = [
        "ignore",
        "save_to_disk",
        "current_tab",
        "new_background_tab",
        "new_foreground_tab",
        "new_window",
        "new_popup",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PermissionrequestArgPermission": (ref: heap.Ref<string>): number => {
      const idx = ["media", "geolocation", "pointerLock", "download", "loadplugin", "filesystem", "fullscreen"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PermissionrequestArgRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_PermissionrequestArgRequest": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "has_PointerLockPermissionRequest_Allow": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "allow" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_PointerLockPermissionRequest_Allow": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).allow);
    },

    "call_PointerLockPermissionRequest_Allow": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.allow();
    },
    "try_PointerLockPermissionRequest_Allow": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.allow();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_PointerLockPermissionRequest_Deny": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "deny" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_PointerLockPermissionRequest_Deny": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).deny);
    },

    "call_PointerLockPermissionRequest_Deny": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.deny();
    },
    "try_PointerLockPermissionRequest_Deny": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.deny();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_PointerLockPermissionRequest_LastUnlockedBySelf": (
      self: heap.Ref<any>,
      retPtr: Pointer
    ): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "lastUnlockedBySelf" in thiz) {
        const val = thiz.lastUnlockedBySelf;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_PointerLockPermissionRequest_LastUnlockedBySelf": (
      self: heap.Ref<object>,
      val: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "lastUnlockedBySelf", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_PointerLockPermissionRequest_Url": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "url" in thiz) {
        const val = thiz.url;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_PointerLockPermissionRequest_Url": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "url", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_PointerLockPermissionRequest_UserGesture": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "userGesture" in thiz) {
        const val = thiz.userGesture;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_PointerLockPermissionRequest_UserGesture": (
      self: heap.Ref<object>,
      val: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "userGesture", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "constof_StopFindingArgAction": (ref: heap.Ref<string>): number => {
      const idx = ["clear", "keep", "activate"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_WebRequestEventInterface": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_WebRequestEventInterface": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_AddContentScripts": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "addContentScripts" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddContentScripts": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.addContentScripts);
    },
    "call_AddContentScripts": (retPtr: Pointer, contentScriptList: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.addContentScripts(A.H.get<object>(contentScriptList));
    },
    "try_AddContentScripts": (
      retPtr: Pointer,
      errPtr: Pointer,
      contentScriptList: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.addContentScripts(A.H.get<object>(contentScriptList));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Back": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "back" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Back": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.back);
    },
    "call_Back": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.back(A.H.get<object>(callback));
    },
    "try_Back": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.back(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CanGoBack": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "canGoBack" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CanGoBack": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.canGoBack);
    },
    "call_CanGoBack": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webviewTag.canGoBack();
      A.store.Bool(retPtr, _ret);
    },
    "try_CanGoBack": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.canGoBack();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CanGoForward": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "canGoForward" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CanGoForward": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.canGoForward);
    },
    "call_CanGoForward": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webviewTag.canGoForward();
      A.store.Bool(retPtr, _ret);
    },
    "try_CanGoForward": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.canGoForward();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CaptureVisibleRegion": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "captureVisibleRegion" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CaptureVisibleRegion": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.captureVisibleRegion);
    },
    "call_CaptureVisibleRegion": (retPtr: Pointer, options: Pointer, callback: heap.Ref<object>): void => {
      const options_ffi = {};

      options_ffi["format"] = A.load.Enum(options + 0, ["jpeg", "png"]);
      if (A.load.Bool(options + 16)) {
        options_ffi["quality"] = A.load.Int64(options + 8);
      }

      const _ret = WEBEXT.webviewTag.captureVisibleRegion(options_ffi, A.H.get<object>(callback));
    },
    "try_CaptureVisibleRegion": (
      retPtr: Pointer,
      errPtr: Pointer,
      options: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["format"] = A.load.Enum(options + 0, ["jpeg", "png"]);
        if (A.load.Bool(options + 16)) {
          options_ffi["quality"] = A.load.Int64(options + 8);
        }

        const _ret = WEBEXT.webviewTag.captureVisibleRegion(options_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ClearData": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "clearData" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ClearData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.clearData);
    },
    "call_ClearData": (retPtr: Pointer, options: Pointer, types: Pointer, callback: heap.Ref<object>): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 8)) {
        options_ffi["since"] = A.load.Float64(options + 0);
      }

      const types_ffi = {};

      if (A.load.Bool(types + 9)) {
        types_ffi["appcache"] = A.load.Bool(types + 0);
      }
      if (A.load.Bool(types + 10)) {
        types_ffi["cache"] = A.load.Bool(types + 1);
      }
      if (A.load.Bool(types + 11)) {
        types_ffi["cookies"] = A.load.Bool(types + 2);
      }
      if (A.load.Bool(types + 12)) {
        types_ffi["fileSystems"] = A.load.Bool(types + 3);
      }
      if (A.load.Bool(types + 13)) {
        types_ffi["indexedDB"] = A.load.Bool(types + 4);
      }
      if (A.load.Bool(types + 14)) {
        types_ffi["localStorage"] = A.load.Bool(types + 5);
      }
      if (A.load.Bool(types + 15)) {
        types_ffi["persistentCookies"] = A.load.Bool(types + 6);
      }
      if (A.load.Bool(types + 16)) {
        types_ffi["sessionCookies"] = A.load.Bool(types + 7);
      }
      if (A.load.Bool(types + 17)) {
        types_ffi["webSQL"] = A.load.Bool(types + 8);
      }

      const _ret = WEBEXT.webviewTag.clearData(options_ffi, types_ffi, A.H.get<object>(callback));
    },
    "try_ClearData": (
      retPtr: Pointer,
      errPtr: Pointer,
      options: Pointer,
      types: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 8)) {
          options_ffi["since"] = A.load.Float64(options + 0);
        }

        const types_ffi = {};

        if (A.load.Bool(types + 9)) {
          types_ffi["appcache"] = A.load.Bool(types + 0);
        }
        if (A.load.Bool(types + 10)) {
          types_ffi["cache"] = A.load.Bool(types + 1);
        }
        if (A.load.Bool(types + 11)) {
          types_ffi["cookies"] = A.load.Bool(types + 2);
        }
        if (A.load.Bool(types + 12)) {
          types_ffi["fileSystems"] = A.load.Bool(types + 3);
        }
        if (A.load.Bool(types + 13)) {
          types_ffi["indexedDB"] = A.load.Bool(types + 4);
        }
        if (A.load.Bool(types + 14)) {
          types_ffi["localStorage"] = A.load.Bool(types + 5);
        }
        if (A.load.Bool(types + 15)) {
          types_ffi["persistentCookies"] = A.load.Bool(types + 6);
        }
        if (A.load.Bool(types + 16)) {
          types_ffi["sessionCookies"] = A.load.Bool(types + 7);
        }
        if (A.load.Bool(types + 17)) {
          types_ffi["webSQL"] = A.load.Bool(types + 8);
        }

        const _ret = WEBEXT.webviewTag.clearData(options_ffi, types_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnClose": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.close && "addListener" in WEBEXT?.webviewTag?.close) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClose": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.close.addListener);
    },
    "call_OnClose": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.close.addListener(A.H.get<object>(callback));
    },
    "try_OnClose": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.close.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClose": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.close && "removeListener" in WEBEXT?.webviewTag?.close) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClose": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.close.removeListener);
    },
    "call_OffClose": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.close.removeListener(A.H.get<object>(callback));
    },
    "try_OffClose": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.close.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClose": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.close && "hasListener" in WEBEXT?.webviewTag?.close) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClose": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.close.hasListener);
    },
    "call_HasOnClose": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.close.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClose": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.close.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnConsolemessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.consolemessage && "addListener" in WEBEXT?.webviewTag?.consolemessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnConsolemessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.consolemessage.addListener);
    },
    "call_OnConsolemessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.consolemessage.addListener(A.H.get<object>(callback));
    },
    "try_OnConsolemessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.consolemessage.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffConsolemessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.consolemessage && "removeListener" in WEBEXT?.webviewTag?.consolemessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffConsolemessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.consolemessage.removeListener);
    },
    "call_OffConsolemessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.consolemessage.removeListener(A.H.get<object>(callback));
    },
    "try_OffConsolemessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.consolemessage.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnConsolemessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.consolemessage && "hasListener" in WEBEXT?.webviewTag?.consolemessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnConsolemessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.consolemessage.hasListener);
    },
    "call_HasOnConsolemessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.consolemessage.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnConsolemessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.consolemessage.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_ContentWindow": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "contentWindow" in WEBEXT?.webviewTag) {
        const val = WEBEXT.webviewTag.contentWindow;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_ContentWindow": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.webviewTag, "contentWindow", A.H.get<object>(val), WEBEXT.webviewTag)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "has_OnContentload": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.contentload && "addListener" in WEBEXT?.webviewTag?.contentload) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnContentload": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.contentload.addListener);
    },
    "call_OnContentload": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.contentload.addListener(A.H.get<object>(callback));
    },
    "try_OnContentload": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.contentload.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffContentload": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.contentload && "removeListener" in WEBEXT?.webviewTag?.contentload) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffContentload": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.contentload.removeListener);
    },
    "call_OffContentload": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.contentload.removeListener(A.H.get<object>(callback));
    },
    "try_OffContentload": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.contentload.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnContentload": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.contentload && "hasListener" in WEBEXT?.webviewTag?.contentload) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnContentload": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.contentload.hasListener);
    },
    "call_HasOnContentload": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.contentload.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnContentload": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.contentload.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_ContextMenus": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "contextMenus" in WEBEXT?.webviewTag) {
        const val = WEBEXT.webviewTag.contextMenus;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_ContextMenus": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(WEBEXT.webviewTag, "contextMenus", A.H.get<object>(val), WEBEXT.webviewTag)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "has_OnDialog": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.dialog && "addListener" in WEBEXT?.webviewTag?.dialog) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDialog": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.dialog.addListener);
    },
    "call_OnDialog": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.dialog.addListener(A.H.get<object>(callback));
    },
    "try_OnDialog": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.dialog.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDialog": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.dialog && "removeListener" in WEBEXT?.webviewTag?.dialog) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDialog": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.dialog.removeListener);
    },
    "call_OffDialog": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.dialog.removeListener(A.H.get<object>(callback));
    },
    "try_OffDialog": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.dialog.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDialog": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.dialog && "hasListener" in WEBEXT?.webviewTag?.dialog) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDialog": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.dialog.hasListener);
    },
    "call_HasOnDialog": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.dialog.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDialog": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.dialog.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ExecuteScript": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "executeScript" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ExecuteScript": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.executeScript);
    },
    "call_ExecuteScript": (retPtr: Pointer, details: Pointer, callback: heap.Ref<object>): void => {
      const details_ffi = {};

      details_ffi["code"] = A.load.Ref(details + 0, undefined);
      details_ffi["file"] = A.load.Ref(details + 4, undefined);

      const _ret = WEBEXT.webviewTag.executeScript(details_ffi, A.H.get<object>(callback));
    },
    "try_ExecuteScript": (
      retPtr: Pointer,
      errPtr: Pointer,
      details: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["code"] = A.load.Ref(details + 0, undefined);
        details_ffi["file"] = A.load.Ref(details + 4, undefined);

        const _ret = WEBEXT.webviewTag.executeScript(details_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnExit": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.exit && "addListener" in WEBEXT?.webviewTag?.exit) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnExit": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.exit.addListener);
    },
    "call_OnExit": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.exit.addListener(A.H.get<object>(callback));
    },
    "try_OnExit": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.exit.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffExit": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.exit && "removeListener" in WEBEXT?.webviewTag?.exit) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffExit": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.exit.removeListener);
    },
    "call_OffExit": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.exit.removeListener(A.H.get<object>(callback));
    },
    "try_OffExit": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.exit.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnExit": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.exit && "hasListener" in WEBEXT?.webviewTag?.exit) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnExit": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.exit.hasListener);
    },
    "call_HasOnExit": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.exit.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnExit": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.exit.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Find": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "find" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Find": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.find);
    },
    "call_Find": (
      retPtr: Pointer,
      searchText: heap.Ref<object>,
      options: Pointer,
      callback: heap.Ref<object>
    ): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 2)) {
        options_ffi["backward"] = A.load.Bool(options + 0);
      }
      if (A.load.Bool(options + 3)) {
        options_ffi["matchCase"] = A.load.Bool(options + 1);
      }

      const _ret = WEBEXT.webviewTag.find(A.H.get<object>(searchText), options_ffi, A.H.get<object>(callback));
    },
    "try_Find": (
      retPtr: Pointer,
      errPtr: Pointer,
      searchText: heap.Ref<object>,
      options: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 2)) {
          options_ffi["backward"] = A.load.Bool(options + 0);
        }
        if (A.load.Bool(options + 3)) {
          options_ffi["matchCase"] = A.load.Bool(options + 1);
        }

        const _ret = WEBEXT.webviewTag.find(A.H.get<object>(searchText), options_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnFindupdate": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.findupdate && "addListener" in WEBEXT?.webviewTag?.findupdate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnFindupdate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.findupdate.addListener);
    },
    "call_OnFindupdate": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.findupdate.addListener(A.H.get<object>(callback));
    },
    "try_OnFindupdate": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.findupdate.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffFindupdate": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.findupdate && "removeListener" in WEBEXT?.webviewTag?.findupdate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffFindupdate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.findupdate.removeListener);
    },
    "call_OffFindupdate": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.findupdate.removeListener(A.H.get<object>(callback));
    },
    "try_OffFindupdate": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.findupdate.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnFindupdate": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.findupdate && "hasListener" in WEBEXT?.webviewTag?.findupdate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnFindupdate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.findupdate.hasListener);
    },
    "call_HasOnFindupdate": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.findupdate.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnFindupdate": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.findupdate.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Forward": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "forward" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Forward": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.forward);
    },
    "call_Forward": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.forward(A.H.get<object>(callback));
    },
    "try_Forward": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.forward(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAudioState": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "getAudioState" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAudioState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.getAudioState);
    },
    "call_GetAudioState": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.getAudioState(A.H.get<object>(callback));
    },
    "try_GetAudioState": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.getAudioState(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetProcessId": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "getProcessId" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetProcessId": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.getProcessId);
    },
    "call_GetProcessId": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webviewTag.getProcessId();
      A.store.Int64(retPtr, _ret);
    },
    "try_GetProcessId": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.getProcessId();
        A.store.Int64(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetUserAgent": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "getUserAgent" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetUserAgent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.getUserAgent);
    },
    "call_GetUserAgent": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webviewTag.getUserAgent();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetUserAgent": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.getUserAgent();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetZoom": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "getZoom" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetZoom": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.getZoom);
    },
    "call_GetZoom": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.getZoom(A.H.get<object>(callback));
    },
    "try_GetZoom": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.getZoom(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetZoomMode": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "getZoomMode" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetZoomMode": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.getZoomMode);
    },
    "call_GetZoomMode": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.getZoomMode(A.H.get<object>(callback));
    },
    "try_GetZoomMode": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.getZoomMode(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Go": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "go" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Go": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.go);
    },
    "call_Go": (retPtr: Pointer, relativeIndex: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.go(relativeIndex, A.H.get<object>(callback));
    },
    "try_Go": (
      retPtr: Pointer,
      errPtr: Pointer,
      relativeIndex: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.go(relativeIndex, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InsertCSS": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "insertCSS" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InsertCSS": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.insertCSS);
    },
    "call_InsertCSS": (retPtr: Pointer, details: Pointer, callback: heap.Ref<object>): void => {
      const details_ffi = {};

      details_ffi["code"] = A.load.Ref(details + 0, undefined);
      details_ffi["file"] = A.load.Ref(details + 4, undefined);

      const _ret = WEBEXT.webviewTag.insertCSS(details_ffi, A.H.get<object>(callback));
    },
    "try_InsertCSS": (
      retPtr: Pointer,
      errPtr: Pointer,
      details: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["code"] = A.load.Ref(details + 0, undefined);
        details_ffi["file"] = A.load.Ref(details + 4, undefined);

        const _ret = WEBEXT.webviewTag.insertCSS(details_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsAudioMuted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "isAudioMuted" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsAudioMuted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.isAudioMuted);
    },
    "call_IsAudioMuted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.isAudioMuted(A.H.get<object>(callback));
    },
    "try_IsAudioMuted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.isAudioMuted(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsSpatialNavigationEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "isSpatialNavigationEnabled" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsSpatialNavigationEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.isSpatialNavigationEnabled);
    },
    "call_IsSpatialNavigationEnabled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.isSpatialNavigationEnabled(A.H.get<object>(callback));
    },
    "try_IsSpatialNavigationEnabled": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.isSpatialNavigationEnabled(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsUserAgentOverridden": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "isUserAgentOverridden" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsUserAgentOverridden": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.isUserAgentOverridden);
    },
    "call_IsUserAgentOverridden": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webviewTag.isUserAgentOverridden();
    },
    "try_IsUserAgentOverridden": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.isUserAgentOverridden();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LoadDataWithBaseUrl": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "loadDataWithBaseUrl" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LoadDataWithBaseUrl": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadDataWithBaseUrl);
    },
    "call_LoadDataWithBaseUrl": (
      retPtr: Pointer,
      dataUrl: heap.Ref<object>,
      baseUrl: heap.Ref<object>,
      virtualUrl: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.webviewTag.loadDataWithBaseUrl(
        A.H.get<object>(dataUrl),
        A.H.get<object>(baseUrl),
        A.H.get<object>(virtualUrl)
      );
    },
    "try_LoadDataWithBaseUrl": (
      retPtr: Pointer,
      errPtr: Pointer,
      dataUrl: heap.Ref<object>,
      baseUrl: heap.Ref<object>,
      virtualUrl: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadDataWithBaseUrl(
          A.H.get<object>(dataUrl),
          A.H.get<object>(baseUrl),
          A.H.get<object>(virtualUrl)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnLoadabort": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.loadabort && "addListener" in WEBEXT?.webviewTag?.loadabort) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnLoadabort": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadabort.addListener);
    },
    "call_OnLoadabort": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.loadabort.addListener(A.H.get<object>(callback));
    },
    "try_OnLoadabort": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadabort.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffLoadabort": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.loadabort && "removeListener" in WEBEXT?.webviewTag?.loadabort) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffLoadabort": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadabort.removeListener);
    },
    "call_OffLoadabort": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.loadabort.removeListener(A.H.get<object>(callback));
    },
    "try_OffLoadabort": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadabort.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnLoadabort": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.loadabort && "hasListener" in WEBEXT?.webviewTag?.loadabort) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnLoadabort": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadabort.hasListener);
    },
    "call_HasOnLoadabort": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.loadabort.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnLoadabort": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadabort.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnLoadcommit": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.loadcommit && "addListener" in WEBEXT?.webviewTag?.loadcommit) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnLoadcommit": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadcommit.addListener);
    },
    "call_OnLoadcommit": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.loadcommit.addListener(A.H.get<object>(callback));
    },
    "try_OnLoadcommit": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadcommit.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffLoadcommit": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.loadcommit && "removeListener" in WEBEXT?.webviewTag?.loadcommit) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffLoadcommit": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadcommit.removeListener);
    },
    "call_OffLoadcommit": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.loadcommit.removeListener(A.H.get<object>(callback));
    },
    "try_OffLoadcommit": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadcommit.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnLoadcommit": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.loadcommit && "hasListener" in WEBEXT?.webviewTag?.loadcommit) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnLoadcommit": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadcommit.hasListener);
    },
    "call_HasOnLoadcommit": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.loadcommit.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnLoadcommit": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadcommit.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnLoadredirect": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.loadredirect && "addListener" in WEBEXT?.webviewTag?.loadredirect) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnLoadredirect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadredirect.addListener);
    },
    "call_OnLoadredirect": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.loadredirect.addListener(A.H.get<object>(callback));
    },
    "try_OnLoadredirect": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadredirect.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffLoadredirect": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.loadredirect && "removeListener" in WEBEXT?.webviewTag?.loadredirect) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffLoadredirect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadredirect.removeListener);
    },
    "call_OffLoadredirect": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.loadredirect.removeListener(A.H.get<object>(callback));
    },
    "try_OffLoadredirect": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadredirect.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnLoadredirect": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.loadredirect && "hasListener" in WEBEXT?.webviewTag?.loadredirect) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnLoadredirect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadredirect.hasListener);
    },
    "call_HasOnLoadredirect": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.loadredirect.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnLoadredirect": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadredirect.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnLoadstart": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.loadstart && "addListener" in WEBEXT?.webviewTag?.loadstart) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnLoadstart": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadstart.addListener);
    },
    "call_OnLoadstart": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.loadstart.addListener(A.H.get<object>(callback));
    },
    "try_OnLoadstart": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadstart.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffLoadstart": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.loadstart && "removeListener" in WEBEXT?.webviewTag?.loadstart) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffLoadstart": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadstart.removeListener);
    },
    "call_OffLoadstart": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.loadstart.removeListener(A.H.get<object>(callback));
    },
    "try_OffLoadstart": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadstart.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnLoadstart": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.loadstart && "hasListener" in WEBEXT?.webviewTag?.loadstart) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnLoadstart": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadstart.hasListener);
    },
    "call_HasOnLoadstart": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.loadstart.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnLoadstart": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadstart.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnLoadstop": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.loadstop && "addListener" in WEBEXT?.webviewTag?.loadstop) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnLoadstop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadstop.addListener);
    },
    "call_OnLoadstop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.loadstop.addListener(A.H.get<object>(callback));
    },
    "try_OnLoadstop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadstop.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffLoadstop": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.loadstop && "removeListener" in WEBEXT?.webviewTag?.loadstop) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffLoadstop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadstop.removeListener);
    },
    "call_OffLoadstop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.loadstop.removeListener(A.H.get<object>(callback));
    },
    "try_OffLoadstop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadstop.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnLoadstop": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.loadstop && "hasListener" in WEBEXT?.webviewTag?.loadstop) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnLoadstop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.loadstop.hasListener);
    },
    "call_HasOnLoadstop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.loadstop.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnLoadstop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.loadstop.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnNewwindow": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.newwindow && "addListener" in WEBEXT?.webviewTag?.newwindow) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnNewwindow": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.newwindow.addListener);
    },
    "call_OnNewwindow": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.newwindow.addListener(A.H.get<object>(callback));
    },
    "try_OnNewwindow": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.newwindow.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffNewwindow": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.newwindow && "removeListener" in WEBEXT?.webviewTag?.newwindow) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffNewwindow": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.newwindow.removeListener);
    },
    "call_OffNewwindow": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.newwindow.removeListener(A.H.get<object>(callback));
    },
    "try_OffNewwindow": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.newwindow.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnNewwindow": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.newwindow && "hasListener" in WEBEXT?.webviewTag?.newwindow) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnNewwindow": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.newwindow.hasListener);
    },
    "call_HasOnNewwindow": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.newwindow.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnNewwindow": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.newwindow.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPermissionrequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.permissionrequest && "addListener" in WEBEXT?.webviewTag?.permissionrequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPermissionrequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.permissionrequest.addListener);
    },
    "call_OnPermissionrequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.permissionrequest.addListener(A.H.get<object>(callback));
    },
    "try_OnPermissionrequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.permissionrequest.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPermissionrequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.permissionrequest && "removeListener" in WEBEXT?.webviewTag?.permissionrequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPermissionrequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.permissionrequest.removeListener);
    },
    "call_OffPermissionrequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.permissionrequest.removeListener(A.H.get<object>(callback));
    },
    "try_OffPermissionrequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.permissionrequest.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPermissionrequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.permissionrequest && "hasListener" in WEBEXT?.webviewTag?.permissionrequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPermissionrequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.permissionrequest.hasListener);
    },
    "call_HasOnPermissionrequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.permissionrequest.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPermissionrequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.permissionrequest.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Print": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "print" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Print": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.print);
    },
    "call_Print": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webviewTag.print();
    },
    "try_Print": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.print();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Reload": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "reload" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Reload": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.reload);
    },
    "call_Reload": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webviewTag.reload();
    },
    "try_Reload": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.reload();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveContentScripts": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "removeContentScripts" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveContentScripts": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.removeContentScripts);
    },
    "call_RemoveContentScripts": (retPtr: Pointer, scriptNameList: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.removeContentScripts(A.H.get<object>(scriptNameList));
    },
    "try_RemoveContentScripts": (
      retPtr: Pointer,
      errPtr: Pointer,
      scriptNameList: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.removeContentScripts(A.H.get<object>(scriptNameList));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_Request": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "request" in WEBEXT?.webviewTag) {
        const val = WEBEXT.webviewTag.request;
        if (typeof val === "undefined") {
          A.store.Bool(retPtr + 0, false);
        } else {
          A.store.Bool(retPtr + 0, true);
        }
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_Request": (val: Pointer): heap.Ref<boolean> => {
      const val_ffi = {};

      return Reflect.set(WEBEXT.webviewTag, "request", val_ffi, WEBEXT.webviewTag) ? A.H.TRUE : A.H.FALSE;
    },
    "has_OnResponsive": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.responsive && "addListener" in WEBEXT?.webviewTag?.responsive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnResponsive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.responsive.addListener);
    },
    "call_OnResponsive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.responsive.addListener(A.H.get<object>(callback));
    },
    "try_OnResponsive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.responsive.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffResponsive": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.responsive && "removeListener" in WEBEXT?.webviewTag?.responsive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffResponsive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.responsive.removeListener);
    },
    "call_OffResponsive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.responsive.removeListener(A.H.get<object>(callback));
    },
    "try_OffResponsive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.responsive.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnResponsive": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.responsive && "hasListener" in WEBEXT?.webviewTag?.responsive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnResponsive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.responsive.hasListener);
    },
    "call_HasOnResponsive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.responsive.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnResponsive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.responsive.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetAudioMuted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "setAudioMuted" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetAudioMuted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.setAudioMuted);
    },
    "call_SetAudioMuted": (retPtr: Pointer, mute: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.webviewTag.setAudioMuted(mute === A.H.TRUE);
    },
    "try_SetAudioMuted": (retPtr: Pointer, errPtr: Pointer, mute: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.setAudioMuted(mute === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetSpatialNavigationEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "setSpatialNavigationEnabled" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetSpatialNavigationEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.setSpatialNavigationEnabled);
    },
    "call_SetSpatialNavigationEnabled": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.webviewTag.setSpatialNavigationEnabled(enabled === A.H.TRUE);
    },
    "try_SetSpatialNavigationEnabled": (
      retPtr: Pointer,
      errPtr: Pointer,
      enabled: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.setSpatialNavigationEnabled(enabled === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetUserAgentOverride": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "setUserAgentOverride" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetUserAgentOverride": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.setUserAgentOverride);
    },
    "call_SetUserAgentOverride": (retPtr: Pointer, userAgent: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.setUserAgentOverride(A.H.get<object>(userAgent));
    },
    "try_SetUserAgentOverride": (retPtr: Pointer, errPtr: Pointer, userAgent: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.setUserAgentOverride(A.H.get<object>(userAgent));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetZoom": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "setZoom" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetZoom": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.setZoom);
    },
    "call_SetZoom": (retPtr: Pointer, zoomFactor: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.setZoom(zoomFactor, A.H.get<object>(callback));
    },
    "try_SetZoom": (
      retPtr: Pointer,
      errPtr: Pointer,
      zoomFactor: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.setZoom(zoomFactor, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetZoomMode": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "setZoomMode" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetZoomMode": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.setZoomMode);
    },
    "call_SetZoomMode": (retPtr: Pointer, ZoomMode: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.setZoomMode(
        ZoomMode > 0 && ZoomMode <= 3 ? ["per-origin", "per-view", "disabled"][ZoomMode - 1] : undefined,
        A.H.get<object>(callback)
      );
    },
    "try_SetZoomMode": (
      retPtr: Pointer,
      errPtr: Pointer,
      ZoomMode: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.setZoomMode(
          ZoomMode > 0 && ZoomMode <= 3 ? ["per-origin", "per-view", "disabled"][ZoomMode - 1] : undefined,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSizechanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.sizechanged && "addListener" in WEBEXT?.webviewTag?.sizechanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSizechanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.sizechanged.addListener);
    },
    "call_OnSizechanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.sizechanged.addListener(A.H.get<object>(callback));
    },
    "try_OnSizechanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.sizechanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSizechanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.sizechanged && "removeListener" in WEBEXT?.webviewTag?.sizechanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSizechanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.sizechanged.removeListener);
    },
    "call_OffSizechanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.sizechanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffSizechanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.sizechanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSizechanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.sizechanged && "hasListener" in WEBEXT?.webviewTag?.sizechanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSizechanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.sizechanged.hasListener);
    },
    "call_HasOnSizechanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.sizechanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSizechanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.sizechanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Stop": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "stop" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Stop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.stop);
    },
    "call_Stop": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webviewTag.stop();
    },
    "try_Stop": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.stop();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StopFinding": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "stopFinding" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StopFinding": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.stopFinding);
    },
    "call_StopFinding": (retPtr: Pointer, action: number): void => {
      const _ret = WEBEXT.webviewTag.stopFinding(
        action > 0 && action <= 3 ? ["clear", "keep", "activate"][action - 1] : undefined
      );
    },
    "try_StopFinding": (retPtr: Pointer, errPtr: Pointer, action: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.stopFinding(
          action > 0 && action <= 3 ? ["clear", "keep", "activate"][action - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Terminate": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag && "terminate" in WEBEXT?.webviewTag) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Terminate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.terminate);
    },
    "call_Terminate": (retPtr: Pointer): void => {
      const _ret = WEBEXT.webviewTag.terminate();
    },
    "try_Terminate": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.terminate();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnUnresponsive": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.unresponsive && "addListener" in WEBEXT?.webviewTag?.unresponsive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnUnresponsive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.unresponsive.addListener);
    },
    "call_OnUnresponsive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.unresponsive.addListener(A.H.get<object>(callback));
    },
    "try_OnUnresponsive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.unresponsive.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffUnresponsive": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.unresponsive && "removeListener" in WEBEXT?.webviewTag?.unresponsive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffUnresponsive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.unresponsive.removeListener);
    },
    "call_OffUnresponsive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.unresponsive.removeListener(A.H.get<object>(callback));
    },
    "try_OffUnresponsive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.unresponsive.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnUnresponsive": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.unresponsive && "hasListener" in WEBEXT?.webviewTag?.unresponsive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnUnresponsive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.unresponsive.hasListener);
    },
    "call_HasOnUnresponsive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.unresponsive.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnUnresponsive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.unresponsive.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnZoomchange": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.zoomchange && "addListener" in WEBEXT?.webviewTag?.zoomchange) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnZoomchange": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.zoomchange.addListener);
    },
    "call_OnZoomchange": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.zoomchange.addListener(A.H.get<object>(callback));
    },
    "try_OnZoomchange": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.zoomchange.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffZoomchange": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.zoomchange && "removeListener" in WEBEXT?.webviewTag?.zoomchange) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffZoomchange": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.zoomchange.removeListener);
    },
    "call_OffZoomchange": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.zoomchange.removeListener(A.H.get<object>(callback));
    },
    "try_OffZoomchange": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.zoomchange.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnZoomchange": (): heap.Ref<boolean> => {
      if (WEBEXT?.webviewTag?.zoomchange && "hasListener" in WEBEXT?.webviewTag?.zoomchange) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnZoomchange": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webviewTag.zoomchange.hasListener);
    },
    "call_HasOnZoomchange": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webviewTag.zoomchange.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnZoomchange": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webviewTag.zoomchange.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
