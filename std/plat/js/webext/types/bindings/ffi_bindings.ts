import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/types", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_LevelOfControl": (ref: heap.Ref<string>): number => {
      const idx = [
        "not_controllable",
        "controlled_by_other_extensions",
        "controllable_by_this_extension",
        "controlled_by_this_extension",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_GetReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 0, false);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "incognitoSpecific" in x ? true : false);
        A.store.Bool(ptr + 0, x["incognitoSpecific"] ? true : false);
        A.store.Enum(
          ptr + 4,
          [
            "not_controllable",
            "controlled_by_other_extensions",
            "controllable_by_this_extension",
            "controlled_by_this_extension",
          ].indexOf(x["levelOfControl"] as string)
        );
        A.store.Ref(ptr + 8, x["value"]);
      }
    },
    "load_GetReturnType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["incognitoSpecific"] = A.load.Bool(ptr + 0);
      } else {
        delete x["incognitoSpecific"];
      }
      x["levelOfControl"] = A.load.Enum(ptr + 4, [
        "not_controllable",
        "controlled_by_other_extensions",
        "controllable_by_this_extension",
        "controlled_by_this_extension",
      ]);
      x["value"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 1, "incognito" in x ? true : false);
        A.store.Bool(ptr + 0, x["incognito"] ? true : false);
      }
    },
    "load_GetArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 1)) {
        x["incognito"] = A.load.Bool(ptr + 0);
      } else {
        delete x["incognito"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ChromeSettingScope": (ref: heap.Ref<string>): number => {
      const idx = ["regular", "regular_only", "incognito_persistent", "incognito_session_only"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SetArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(
          ptr + 0,
          ["regular", "regular_only", "incognito_persistent", "incognito_session_only"].indexOf(x["scope"] as string)
        );
        A.store.Ref(ptr + 4, x["value"]);
      }
    },
    "load_SetArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["scope"] = A.load.Enum(ptr + 0, ["regular", "regular_only", "incognito_persistent", "incognito_session_only"]);
      x["value"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ClearArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(
          ptr + 0,
          ["regular", "regular_only", "incognito_persistent", "incognito_session_only"].indexOf(x["scope"] as string)
        );
      }
    },
    "load_ClearArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["scope"] = A.load.Enum(ptr + 0, ["regular", "regular_only", "incognito_persistent", "incognito_session_only"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "has_ChromeSetting_Get": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "get" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ChromeSetting_Get": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).get);
    },

    "call_ChromeSetting_Get": (self: heap.Ref<object>, retPtr: Pointer, details: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const details_ffi = {};

      if (A.load.Bool(details + 1)) {
        details_ffi["incognito"] = A.load.Bool(details + 0);
      }
      const _ret = thiz.get(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ChromeSetting_Get": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      details: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const details_ffi = {};

        if (A.load.Bool(details + 1)) {
          details_ffi["incognito"] = A.load.Bool(details + 0);
        }
        const _ret = thiz.get(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ChromeSetting_Set": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "set" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ChromeSetting_Set": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).set);
    },

    "call_ChromeSetting_Set": (self: heap.Ref<object>, retPtr: Pointer, details: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const details_ffi = {};

      details_ffi["scope"] = A.load.Enum(details + 0, [
        "regular",
        "regular_only",
        "incognito_persistent",
        "incognito_session_only",
      ]);
      details_ffi["value"] = A.load.Ref(details + 4, undefined);
      const _ret = thiz.set(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ChromeSetting_Set": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      details: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const details_ffi = {};

        details_ffi["scope"] = A.load.Enum(details + 0, [
          "regular",
          "regular_only",
          "incognito_persistent",
          "incognito_session_only",
        ]);
        details_ffi["value"] = A.load.Ref(details + 4, undefined);
        const _ret = thiz.set(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ChromeSetting_Clear": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "clear" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ChromeSetting_Clear": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).clear);
    },

    "call_ChromeSetting_Clear": (self: heap.Ref<object>, retPtr: Pointer, details: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const details_ffi = {};

      details_ffi["scope"] = A.load.Enum(details + 0, [
        "regular",
        "regular_only",
        "incognito_persistent",
        "incognito_session_only",
      ]);
      const _ret = thiz.clear(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ChromeSetting_Clear": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      details: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const details_ffi = {};

        details_ffi["scope"] = A.load.Enum(details + 0, [
          "regular",
          "regular_only",
          "incognito_persistent",
          "incognito_session_only",
        ]);
        const _ret = thiz.clear(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
