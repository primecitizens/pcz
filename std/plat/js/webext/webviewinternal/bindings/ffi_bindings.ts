import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/webviewinternal", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
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

    "store_DataTypeSet": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_DataTypeSet": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
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

    "store_FindArgCallbackArgResultsFieldSelectionRect": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_FindArgCallbackArgResultsFieldSelectionRect": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["height"] = A.load.Int64(ptr + 0);
      x["left"] = A.load.Int64(ptr + 8);
      x["top"] = A.load.Int64(ptr + 16);
      x["width"] = A.load.Int64(ptr + 24);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_FindArgCallbackArgResults": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_FindArgCallbackArgResults": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
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

    "store_FindArgOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_FindArgOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
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
    "constof_ZoomMode": (ref: heap.Ref<string>): number => {
      const idx = ["per-origin", "per-view", "disabled"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "get_MAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND": (retPtr: Pointer): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "MAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND" in WEBEXT?.webViewInternal) {
        const val = WEBEXT.webViewInternal.MAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_MAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND": (val: heap.Ref<object>): heap.Ref<boolean> => {
      return Reflect.set(
        WEBEXT.webViewInternal,
        "MAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND",
        A.H.get<object>(val),
        WEBEXT.webViewInternal
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },

    "store_RemovalOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_RemovalOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["since"] = A.load.Float64(ptr + 0);
      } else {
        delete x["since"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_SetPermissionAction": (ref: heap.Ref<string>): number => {
      const idx = ["allow", "deny", "default"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_StopFindingAction": (ref: heap.Ref<string>): number => {
      const idx = ["clear", "keep", "activate"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_AddContentScripts": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "addContentScripts" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddContentScripts": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.addContentScripts);
    },
    "call_AddContentScripts": (retPtr: Pointer, instanceId: number, contentScriptList: heap.Ref<object>): void => {
      const _ret = WEBEXT.webViewInternal.addContentScripts(instanceId, A.H.get<object>(contentScriptList));
    },
    "try_AddContentScripts": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      contentScriptList: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.addContentScripts(instanceId, A.H.get<object>(contentScriptList));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CaptureVisibleRegion": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "captureVisibleRegion" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CaptureVisibleRegion": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.captureVisibleRegion);
    },
    "call_CaptureVisibleRegion": (
      retPtr: Pointer,
      instanceId: number,
      options: Pointer,
      callback: heap.Ref<object>
    ): void => {
      const options_ffi = {};

      options_ffi["format"] = A.load.Enum(options + 0, ["jpeg", "png"]);
      if (A.load.Bool(options + 16)) {
        options_ffi["quality"] = A.load.Int64(options + 8);
      }

      const _ret = WEBEXT.webViewInternal.captureVisibleRegion(instanceId, options_ffi, A.H.get<object>(callback));
    },
    "try_CaptureVisibleRegion": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      options: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["format"] = A.load.Enum(options + 0, ["jpeg", "png"]);
        if (A.load.Bool(options + 16)) {
          options_ffi["quality"] = A.load.Int64(options + 8);
        }

        const _ret = WEBEXT.webViewInternal.captureVisibleRegion(instanceId, options_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ClearData": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "clearData" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ClearData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.clearData);
    },
    "call_ClearData": (
      retPtr: Pointer,
      instanceId: number,
      options: Pointer,
      dataToRemove: Pointer,
      callback: heap.Ref<object>
    ): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 8)) {
        options_ffi["since"] = A.load.Float64(options + 0);
      }

      const dataToRemove_ffi = {};

      if (A.load.Bool(dataToRemove + 9)) {
        dataToRemove_ffi["appcache"] = A.load.Bool(dataToRemove + 0);
      }
      if (A.load.Bool(dataToRemove + 10)) {
        dataToRemove_ffi["cache"] = A.load.Bool(dataToRemove + 1);
      }
      if (A.load.Bool(dataToRemove + 11)) {
        dataToRemove_ffi["cookies"] = A.load.Bool(dataToRemove + 2);
      }
      if (A.load.Bool(dataToRemove + 12)) {
        dataToRemove_ffi["fileSystems"] = A.load.Bool(dataToRemove + 3);
      }
      if (A.load.Bool(dataToRemove + 13)) {
        dataToRemove_ffi["indexedDB"] = A.load.Bool(dataToRemove + 4);
      }
      if (A.load.Bool(dataToRemove + 14)) {
        dataToRemove_ffi["localStorage"] = A.load.Bool(dataToRemove + 5);
      }
      if (A.load.Bool(dataToRemove + 15)) {
        dataToRemove_ffi["persistentCookies"] = A.load.Bool(dataToRemove + 6);
      }
      if (A.load.Bool(dataToRemove + 16)) {
        dataToRemove_ffi["sessionCookies"] = A.load.Bool(dataToRemove + 7);
      }
      if (A.load.Bool(dataToRemove + 17)) {
        dataToRemove_ffi["webSQL"] = A.load.Bool(dataToRemove + 8);
      }

      const _ret = WEBEXT.webViewInternal.clearData(
        instanceId,
        options_ffi,
        dataToRemove_ffi,
        A.H.get<object>(callback)
      );
    },
    "try_ClearData": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      options: Pointer,
      dataToRemove: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 8)) {
          options_ffi["since"] = A.load.Float64(options + 0);
        }

        const dataToRemove_ffi = {};

        if (A.load.Bool(dataToRemove + 9)) {
          dataToRemove_ffi["appcache"] = A.load.Bool(dataToRemove + 0);
        }
        if (A.load.Bool(dataToRemove + 10)) {
          dataToRemove_ffi["cache"] = A.load.Bool(dataToRemove + 1);
        }
        if (A.load.Bool(dataToRemove + 11)) {
          dataToRemove_ffi["cookies"] = A.load.Bool(dataToRemove + 2);
        }
        if (A.load.Bool(dataToRemove + 12)) {
          dataToRemove_ffi["fileSystems"] = A.load.Bool(dataToRemove + 3);
        }
        if (A.load.Bool(dataToRemove + 13)) {
          dataToRemove_ffi["indexedDB"] = A.load.Bool(dataToRemove + 4);
        }
        if (A.load.Bool(dataToRemove + 14)) {
          dataToRemove_ffi["localStorage"] = A.load.Bool(dataToRemove + 5);
        }
        if (A.load.Bool(dataToRemove + 15)) {
          dataToRemove_ffi["persistentCookies"] = A.load.Bool(dataToRemove + 6);
        }
        if (A.load.Bool(dataToRemove + 16)) {
          dataToRemove_ffi["sessionCookies"] = A.load.Bool(dataToRemove + 7);
        }
        if (A.load.Bool(dataToRemove + 17)) {
          dataToRemove_ffi["webSQL"] = A.load.Bool(dataToRemove + 8);
        }

        const _ret = WEBEXT.webViewInternal.clearData(
          instanceId,
          options_ffi,
          dataToRemove_ffi,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ExecuteScript": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "executeScript" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ExecuteScript": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.executeScript);
    },
    "call_ExecuteScript": (
      retPtr: Pointer,
      instanceId: number,
      src: heap.Ref<object>,
      details: Pointer,
      callback: heap.Ref<object>
    ): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 32)) {
        details_ffi["allFrames"] = A.load.Bool(details + 0);
      }
      details_ffi["code"] = A.load.Ref(details + 4, undefined);
      details_ffi["cssOrigin"] = A.load.Enum(details + 8, ["author", "user"]);
      details_ffi["file"] = A.load.Ref(details + 12, undefined);
      if (A.load.Bool(details + 33)) {
        details_ffi["frameId"] = A.load.Int64(details + 16);
      }
      if (A.load.Bool(details + 34)) {
        details_ffi["matchAboutBlank"] = A.load.Bool(details + 24);
      }
      details_ffi["runAt"] = A.load.Enum(details + 28, ["document_start", "document_end", "document_idle"]);

      const _ret = WEBEXT.webViewInternal.executeScript(
        instanceId,
        A.H.get<object>(src),
        details_ffi,
        A.H.get<object>(callback)
      );
    },
    "try_ExecuteScript": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      src: heap.Ref<object>,
      details: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 32)) {
          details_ffi["allFrames"] = A.load.Bool(details + 0);
        }
        details_ffi["code"] = A.load.Ref(details + 4, undefined);
        details_ffi["cssOrigin"] = A.load.Enum(details + 8, ["author", "user"]);
        details_ffi["file"] = A.load.Ref(details + 12, undefined);
        if (A.load.Bool(details + 33)) {
          details_ffi["frameId"] = A.load.Int64(details + 16);
        }
        if (A.load.Bool(details + 34)) {
          details_ffi["matchAboutBlank"] = A.load.Bool(details + 24);
        }
        details_ffi["runAt"] = A.load.Enum(details + 28, ["document_start", "document_end", "document_idle"]);

        const _ret = WEBEXT.webViewInternal.executeScript(
          instanceId,
          A.H.get<object>(src),
          details_ffi,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Find": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "find" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Find": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.find);
    },
    "call_Find": (
      retPtr: Pointer,
      instanceId: number,
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

      const _ret = WEBEXT.webViewInternal.find(
        instanceId,
        A.H.get<object>(searchText),
        options_ffi,
        A.H.get<object>(callback)
      );
    },
    "try_Find": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
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

        const _ret = WEBEXT.webViewInternal.find(
          instanceId,
          A.H.get<object>(searchText),
          options_ffi,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAudioState": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "getAudioState" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAudioState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.getAudioState);
    },
    "call_GetAudioState": (retPtr: Pointer, instanceId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webViewInternal.getAudioState(instanceId, A.H.get<object>(callback));
    },
    "try_GetAudioState": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.getAudioState(instanceId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetZoom": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "getZoom" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetZoom": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.getZoom);
    },
    "call_GetZoom": (retPtr: Pointer, instanceId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webViewInternal.getZoom(instanceId, A.H.get<object>(callback));
    },
    "try_GetZoom": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.getZoom(instanceId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetZoomMode": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "getZoomMode" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetZoomMode": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.getZoomMode);
    },
    "call_GetZoomMode": (retPtr: Pointer, instanceId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webViewInternal.getZoomMode(instanceId, A.H.get<object>(callback));
    },
    "try_GetZoomMode": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.getZoomMode(instanceId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Go": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "go" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Go": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.go);
    },
    "call_Go": (retPtr: Pointer, instanceId: number, relativeIndex: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webViewInternal.go(instanceId, relativeIndex, A.H.get<object>(callback));
    },
    "try_Go": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      relativeIndex: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.go(instanceId, relativeIndex, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InsertCSS": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "insertCSS" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InsertCSS": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.insertCSS);
    },
    "call_InsertCSS": (
      retPtr: Pointer,
      instanceId: number,
      src: heap.Ref<object>,
      details: Pointer,
      callback: heap.Ref<object>
    ): void => {
      const details_ffi = {};

      if (A.load.Bool(details + 32)) {
        details_ffi["allFrames"] = A.load.Bool(details + 0);
      }
      details_ffi["code"] = A.load.Ref(details + 4, undefined);
      details_ffi["cssOrigin"] = A.load.Enum(details + 8, ["author", "user"]);
      details_ffi["file"] = A.load.Ref(details + 12, undefined);
      if (A.load.Bool(details + 33)) {
        details_ffi["frameId"] = A.load.Int64(details + 16);
      }
      if (A.load.Bool(details + 34)) {
        details_ffi["matchAboutBlank"] = A.load.Bool(details + 24);
      }
      details_ffi["runAt"] = A.load.Enum(details + 28, ["document_start", "document_end", "document_idle"]);

      const _ret = WEBEXT.webViewInternal.insertCSS(
        instanceId,
        A.H.get<object>(src),
        details_ffi,
        A.H.get<object>(callback)
      );
    },
    "try_InsertCSS": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      src: heap.Ref<object>,
      details: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        if (A.load.Bool(details + 32)) {
          details_ffi["allFrames"] = A.load.Bool(details + 0);
        }
        details_ffi["code"] = A.load.Ref(details + 4, undefined);
        details_ffi["cssOrigin"] = A.load.Enum(details + 8, ["author", "user"]);
        details_ffi["file"] = A.load.Ref(details + 12, undefined);
        if (A.load.Bool(details + 33)) {
          details_ffi["frameId"] = A.load.Int64(details + 16);
        }
        if (A.load.Bool(details + 34)) {
          details_ffi["matchAboutBlank"] = A.load.Bool(details + 24);
        }
        details_ffi["runAt"] = A.load.Enum(details + 28, ["document_start", "document_end", "document_idle"]);

        const _ret = WEBEXT.webViewInternal.insertCSS(
          instanceId,
          A.H.get<object>(src),
          details_ffi,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsAudioMuted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "isAudioMuted" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsAudioMuted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.isAudioMuted);
    },
    "call_IsAudioMuted": (retPtr: Pointer, instanceId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webViewInternal.isAudioMuted(instanceId, A.H.get<object>(callback));
    },
    "try_IsAudioMuted": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.isAudioMuted(instanceId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsSpatialNavigationEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "isSpatialNavigationEnabled" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsSpatialNavigationEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.isSpatialNavigationEnabled);
    },
    "call_IsSpatialNavigationEnabled": (retPtr: Pointer, instanceId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webViewInternal.isSpatialNavigationEnabled(instanceId, A.H.get<object>(callback));
    },
    "try_IsSpatialNavigationEnabled": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.isSpatialNavigationEnabled(instanceId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LoadDataWithBaseUrl": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "loadDataWithBaseUrl" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LoadDataWithBaseUrl": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.loadDataWithBaseUrl);
    },
    "call_LoadDataWithBaseUrl": (
      retPtr: Pointer,
      instanceId: number,
      dataUrl: heap.Ref<object>,
      baseUrl: heap.Ref<object>,
      virtualUrl: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.webViewInternal.loadDataWithBaseUrl(
        instanceId,
        A.H.get<object>(dataUrl),
        A.H.get<object>(baseUrl),
        A.H.get<object>(virtualUrl),
        A.H.get<object>(callback)
      );
    },
    "try_LoadDataWithBaseUrl": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      dataUrl: heap.Ref<object>,
      baseUrl: heap.Ref<object>,
      virtualUrl: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.loadDataWithBaseUrl(
          instanceId,
          A.H.get<object>(dataUrl),
          A.H.get<object>(baseUrl),
          A.H.get<object>(virtualUrl),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Navigate": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "navigate" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Navigate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.navigate);
    },
    "call_Navigate": (retPtr: Pointer, instanceId: number, src: heap.Ref<object>): void => {
      const _ret = WEBEXT.webViewInternal.navigate(instanceId, A.H.get<object>(src));
    },
    "try_Navigate": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      src: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.navigate(instanceId, A.H.get<object>(src));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OverrideUserAgent": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "overrideUserAgent" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OverrideUserAgent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.overrideUserAgent);
    },
    "call_OverrideUserAgent": (retPtr: Pointer, instanceId: number, userAgentOverride: heap.Ref<object>): void => {
      const _ret = WEBEXT.webViewInternal.overrideUserAgent(instanceId, A.H.get<object>(userAgentOverride));
    },
    "try_OverrideUserAgent": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      userAgentOverride: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.overrideUserAgent(instanceId, A.H.get<object>(userAgentOverride));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Reload": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "reload" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Reload": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.reload);
    },
    "call_Reload": (retPtr: Pointer, instanceId: number): void => {
      const _ret = WEBEXT.webViewInternal.reload(instanceId);
    },
    "try_Reload": (retPtr: Pointer, errPtr: Pointer, instanceId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.reload(instanceId);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveContentScripts": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "removeContentScripts" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveContentScripts": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.removeContentScripts);
    },
    "call_RemoveContentScripts": (retPtr: Pointer, instanceId: number, scriptNameList: heap.Ref<object>): void => {
      const _ret = WEBEXT.webViewInternal.removeContentScripts(instanceId, A.H.get<object>(scriptNameList));
    },
    "try_RemoveContentScripts": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      scriptNameList: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.removeContentScripts(instanceId, A.H.get<object>(scriptNameList));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetAllowScaling": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "setAllowScaling" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetAllowScaling": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.setAllowScaling);
    },
    "call_SetAllowScaling": (retPtr: Pointer, instanceId: number, allow: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.webViewInternal.setAllowScaling(instanceId, allow === A.H.TRUE);
    },
    "try_SetAllowScaling": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      allow: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.setAllowScaling(instanceId, allow === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetAllowTransparency": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "setAllowTransparency" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetAllowTransparency": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.setAllowTransparency);
    },
    "call_SetAllowTransparency": (retPtr: Pointer, instanceId: number, allow: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.webViewInternal.setAllowTransparency(instanceId, allow === A.H.TRUE);
    },
    "try_SetAllowTransparency": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      allow: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.setAllowTransparency(instanceId, allow === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetAudioMuted": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "setAudioMuted" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetAudioMuted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.setAudioMuted);
    },
    "call_SetAudioMuted": (retPtr: Pointer, instanceId: number, mute: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.webViewInternal.setAudioMuted(instanceId, mute === A.H.TRUE);
    },
    "try_SetAudioMuted": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      mute: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.setAudioMuted(instanceId, mute === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetName": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "setName" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetName": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.setName);
    },
    "call_SetName": (retPtr: Pointer, instanceId: number, frameName: heap.Ref<object>): void => {
      const _ret = WEBEXT.webViewInternal.setName(instanceId, A.H.get<object>(frameName));
    },
    "try_SetName": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      frameName: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.setName(instanceId, A.H.get<object>(frameName));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPermission": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "setPermission" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPermission": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.setPermission);
    },
    "call_SetPermission": (
      retPtr: Pointer,
      instanceId: number,
      requestId: number,
      action: number,
      userInput: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.webViewInternal.setPermission(
        instanceId,
        requestId,
        action > 0 && action <= 3 ? ["allow", "deny", "default"][action - 1] : undefined,
        A.H.get<object>(userInput),
        A.H.get<object>(callback)
      );
    },
    "try_SetPermission": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      requestId: number,
      action: number,
      userInput: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.setPermission(
          instanceId,
          requestId,
          action > 0 && action <= 3 ? ["allow", "deny", "default"][action - 1] : undefined,
          A.H.get<object>(userInput),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetSpatialNavigationEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "setSpatialNavigationEnabled" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetSpatialNavigationEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.setSpatialNavigationEnabled);
    },
    "call_SetSpatialNavigationEnabled": (
      retPtr: Pointer,
      instanceId: number,
      spatialNavEnabled: heap.Ref<boolean>
    ): void => {
      const _ret = WEBEXT.webViewInternal.setSpatialNavigationEnabled(instanceId, spatialNavEnabled === A.H.TRUE);
    },
    "try_SetSpatialNavigationEnabled": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      spatialNavEnabled: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.setSpatialNavigationEnabled(instanceId, spatialNavEnabled === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetZoom": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "setZoom" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetZoom": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.setZoom);
    },
    "call_SetZoom": (retPtr: Pointer, instanceId: number, zoomFactor: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webViewInternal.setZoom(instanceId, zoomFactor, A.H.get<object>(callback));
    },
    "try_SetZoom": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      zoomFactor: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.setZoom(instanceId, zoomFactor, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetZoomMode": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "setZoomMode" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetZoomMode": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.setZoomMode);
    },
    "call_SetZoomMode": (retPtr: Pointer, instanceId: number, ZoomMode: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.webViewInternal.setZoomMode(
        instanceId,
        ZoomMode > 0 && ZoomMode <= 3 ? ["per-origin", "per-view", "disabled"][ZoomMode - 1] : undefined,
        A.H.get<object>(callback)
      );
    },
    "try_SetZoomMode": (
      retPtr: Pointer,
      errPtr: Pointer,
      instanceId: number,
      ZoomMode: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.setZoomMode(
          instanceId,
          ZoomMode > 0 && ZoomMode <= 3 ? ["per-origin", "per-view", "disabled"][ZoomMode - 1] : undefined,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Stop": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "stop" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Stop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.stop);
    },
    "call_Stop": (retPtr: Pointer, instanceId: number): void => {
      const _ret = WEBEXT.webViewInternal.stop(instanceId);
    },
    "try_Stop": (retPtr: Pointer, errPtr: Pointer, instanceId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.stop(instanceId);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StopFinding": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "stopFinding" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StopFinding": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.stopFinding);
    },
    "call_StopFinding": (retPtr: Pointer, instanceId: number, action: number): void => {
      const _ret = WEBEXT.webViewInternal.stopFinding(
        instanceId,
        action > 0 && action <= 3 ? ["clear", "keep", "activate"][action - 1] : undefined
      );
    },
    "try_StopFinding": (retPtr: Pointer, errPtr: Pointer, instanceId: number, action: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.stopFinding(
          instanceId,
          action > 0 && action <= 3 ? ["clear", "keep", "activate"][action - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Terminate": (): heap.Ref<boolean> => {
      if (WEBEXT?.webViewInternal && "terminate" in WEBEXT?.webViewInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Terminate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webViewInternal.terminate);
    },
    "call_Terminate": (retPtr: Pointer, instanceId: number): void => {
      const _ret = WEBEXT.webViewInternal.terminate(instanceId);
    },
    "try_Terminate": (retPtr: Pointer, errPtr: Pointer, instanceId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webViewInternal.terminate(instanceId);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
