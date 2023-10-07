import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/settingsprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_ControlledBy": (ref: heap.Ref<string>): number => {
      const idx = [
        "DEVICE_POLICY",
        "USER_POLICY",
        "OWNER",
        "PRIMARY_USER",
        "EXTENSION",
        "PARENT",
        "CHILD_RESTRICTION",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_Enforcement": (ref: heap.Ref<string>): number => {
      const idx = ["ENFORCED", "RECOMMENDED", "PARENT_SUPERVISED"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PrefType": (ref: heap.Ref<string>): number => {
      const idx = ["BOOLEAN", "NUMBER", "STRING", "URL", "LIST", "DICTIONARY"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PrefObject": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 43, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
        A.store.Enum(ptr + 12, -1);
        A.store.Ref(ptr + 16, undefined);
        A.store.Enum(ptr + 20, -1);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Bool(ptr + 41, false);
        A.store.Bool(ptr + 32, false);
        A.store.Ref(ptr + 36, undefined);
        A.store.Bool(ptr + 42, false);
        A.store.Bool(ptr + 40, false);
      } else {
        A.store.Bool(ptr + 43, true);
        A.store.Ref(ptr + 0, x["key"]);
        A.store.Enum(
          ptr + 4,
          ["BOOLEAN", "NUMBER", "STRING", "URL", "LIST", "DICTIONARY"].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 8, x["value"]);
        A.store.Enum(
          ptr + 12,
          ["DEVICE_POLICY", "USER_POLICY", "OWNER", "PRIMARY_USER", "EXTENSION", "PARENT", "CHILD_RESTRICTION"].indexOf(
            x["controlledBy"] as string
          )
        );
        A.store.Ref(ptr + 16, x["controlledByName"]);
        A.store.Enum(ptr + 20, ["ENFORCED", "RECOMMENDED", "PARENT_SUPERVISED"].indexOf(x["enforcement"] as string));
        A.store.Ref(ptr + 24, x["recommendedValue"]);
        A.store.Ref(ptr + 28, x["userSelectableValues"]);
        A.store.Bool(ptr + 41, "userControlDisabled" in x ? true : false);
        A.store.Bool(ptr + 32, x["userControlDisabled"] ? true : false);
        A.store.Ref(ptr + 36, x["extensionId"]);
        A.store.Bool(ptr + 42, "extensionCanBeDisabled" in x ? true : false);
        A.store.Bool(ptr + 40, x["extensionCanBeDisabled"] ? true : false);
      }
    },
    "load_PrefObject": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["key"] = A.load.Ref(ptr + 0, undefined);
      x["type"] = A.load.Enum(ptr + 4, ["BOOLEAN", "NUMBER", "STRING", "URL", "LIST", "DICTIONARY"]);
      x["value"] = A.load.Ref(ptr + 8, undefined);
      x["controlledBy"] = A.load.Enum(ptr + 12, [
        "DEVICE_POLICY",
        "USER_POLICY",
        "OWNER",
        "PRIMARY_USER",
        "EXTENSION",
        "PARENT",
        "CHILD_RESTRICTION",
      ]);
      x["controlledByName"] = A.load.Ref(ptr + 16, undefined);
      x["enforcement"] = A.load.Enum(ptr + 20, ["ENFORCED", "RECOMMENDED", "PARENT_SUPERVISED"]);
      x["recommendedValue"] = A.load.Ref(ptr + 24, undefined);
      x["userSelectableValues"] = A.load.Ref(ptr + 28, undefined);
      if (A.load.Bool(ptr + 41)) {
        x["userControlDisabled"] = A.load.Bool(ptr + 32);
      } else {
        delete x["userControlDisabled"];
      }
      x["extensionId"] = A.load.Ref(ptr + 36, undefined);
      if (A.load.Bool(ptr + 42)) {
        x["extensionCanBeDisabled"] = A.load.Bool(ptr + 40);
      } else {
        delete x["extensionCanBeDisabled"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetAllPrefs": (): heap.Ref<boolean> => {
      if (WEBEXT?.settingsPrivate && "getAllPrefs" in WEBEXT?.settingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAllPrefs": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.settingsPrivate.getAllPrefs);
    },
    "call_GetAllPrefs": (retPtr: Pointer): void => {
      const _ret = WEBEXT.settingsPrivate.getAllPrefs();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAllPrefs": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.settingsPrivate.getAllPrefs();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDefaultZoom": (): heap.Ref<boolean> => {
      if (WEBEXT?.settingsPrivate && "getDefaultZoom" in WEBEXT?.settingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDefaultZoom": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.settingsPrivate.getDefaultZoom);
    },
    "call_GetDefaultZoom": (retPtr: Pointer): void => {
      const _ret = WEBEXT.settingsPrivate.getDefaultZoom();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDefaultZoom": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.settingsPrivate.getDefaultZoom();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPref": (): heap.Ref<boolean> => {
      if (WEBEXT?.settingsPrivate && "getPref" in WEBEXT?.settingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPref": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.settingsPrivate.getPref);
    },
    "call_GetPref": (retPtr: Pointer, name: heap.Ref<object>): void => {
      const _ret = WEBEXT.settingsPrivate.getPref(A.H.get<object>(name));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPref": (retPtr: Pointer, errPtr: Pointer, name: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.settingsPrivate.getPref(A.H.get<object>(name));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPrefsChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.settingsPrivate?.onPrefsChanged && "addListener" in WEBEXT?.settingsPrivate?.onPrefsChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPrefsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.settingsPrivate.onPrefsChanged.addListener);
    },
    "call_OnPrefsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.settingsPrivate.onPrefsChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnPrefsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.settingsPrivate.onPrefsChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPrefsChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.settingsPrivate?.onPrefsChanged && "removeListener" in WEBEXT?.settingsPrivate?.onPrefsChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPrefsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.settingsPrivate.onPrefsChanged.removeListener);
    },
    "call_OffPrefsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.settingsPrivate.onPrefsChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffPrefsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.settingsPrivate.onPrefsChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPrefsChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.settingsPrivate?.onPrefsChanged && "hasListener" in WEBEXT?.settingsPrivate?.onPrefsChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPrefsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.settingsPrivate.onPrefsChanged.hasListener);
    },
    "call_HasOnPrefsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.settingsPrivate.onPrefsChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPrefsChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.settingsPrivate.onPrefsChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetDefaultZoom": (): heap.Ref<boolean> => {
      if (WEBEXT?.settingsPrivate && "setDefaultZoom" in WEBEXT?.settingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetDefaultZoom": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.settingsPrivate.setDefaultZoom);
    },
    "call_SetDefaultZoom": (retPtr: Pointer, zoom: number): void => {
      const _ret = WEBEXT.settingsPrivate.setDefaultZoom(zoom);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetDefaultZoom": (retPtr: Pointer, errPtr: Pointer, zoom: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.settingsPrivate.setDefaultZoom(zoom);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPref": (): heap.Ref<boolean> => {
      if (WEBEXT?.settingsPrivate && "setPref" in WEBEXT?.settingsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPref": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.settingsPrivate.setPref);
    },
    "call_SetPref": (
      retPtr: Pointer,
      name: heap.Ref<object>,
      value: heap.Ref<object>,
      pageId: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.settingsPrivate.setPref(
        A.H.get<object>(name),
        A.H.get<object>(value),
        A.H.get<object>(pageId)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_SetPref": (
      retPtr: Pointer,
      errPtr: Pointer,
      name: heap.Ref<object>,
      value: heap.Ref<object>,
      pageId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.settingsPrivate.setPref(
          A.H.get<object>(name),
          A.H.get<object>(value),
          A.H.get<object>(pageId)
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
