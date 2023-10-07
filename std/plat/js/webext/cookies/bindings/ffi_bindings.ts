import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/cookies", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_SameSiteStatus": (ref: heap.Ref<string>): number => {
      const idx = ["no_restriction", "lax", "strict", "unspecified"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Cookie": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 45, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 44, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Enum(ptr + 28, -1);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 33, false);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
      } else {
        A.store.Bool(ptr + 45, true);
        A.store.Ref(ptr + 0, x["domain"]);
        A.store.Bool(ptr + 44, "expirationDate" in x ? true : false);
        A.store.Float64(ptr + 8, x["expirationDate"] === undefined ? 0 : (x["expirationDate"] as number));
        A.store.Bool(ptr + 16, x["hostOnly"] ? true : false);
        A.store.Bool(ptr + 17, x["httpOnly"] ? true : false);
        A.store.Ref(ptr + 20, x["name"]);
        A.store.Ref(ptr + 24, x["path"]);
        A.store.Enum(ptr + 28, ["no_restriction", "lax", "strict", "unspecified"].indexOf(x["sameSite"] as string));
        A.store.Bool(ptr + 32, x["secure"] ? true : false);
        A.store.Bool(ptr + 33, x["session"] ? true : false);
        A.store.Ref(ptr + 36, x["storeId"]);
        A.store.Ref(ptr + 40, x["value"]);
      }
    },
    "load_Cookie": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["domain"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 44)) {
        x["expirationDate"] = A.load.Float64(ptr + 8);
      } else {
        delete x["expirationDate"];
      }
      x["hostOnly"] = A.load.Bool(ptr + 16);
      x["httpOnly"] = A.load.Bool(ptr + 17);
      x["name"] = A.load.Ref(ptr + 20, undefined);
      x["path"] = A.load.Ref(ptr + 24, undefined);
      x["sameSite"] = A.load.Enum(ptr + 28, ["no_restriction", "lax", "strict", "unspecified"]);
      x["secure"] = A.load.Bool(ptr + 32);
      x["session"] = A.load.Bool(ptr + 33);
      x["storeId"] = A.load.Ref(ptr + 36, undefined);
      x["value"] = A.load.Ref(ptr + 40, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CookieDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["storeId"]);
        A.store.Ref(ptr + 8, x["url"]);
      }
    },
    "load_CookieDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["storeId"] = A.load.Ref(ptr + 4, undefined);
      x["url"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CookieStore": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["tabIds"]);
      }
    },
    "load_CookieStore": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["tabIds"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetAllArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 26, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
      } else {
        A.store.Bool(ptr + 26, true);
        A.store.Ref(ptr + 0, x["domain"]);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Ref(ptr + 8, x["path"]);
        A.store.Bool(ptr + 24, "secure" in x ? true : false);
        A.store.Bool(ptr + 12, x["secure"] ? true : false);
        A.store.Bool(ptr + 25, "session" in x ? true : false);
        A.store.Bool(ptr + 13, x["session"] ? true : false);
        A.store.Ref(ptr + 16, x["storeId"]);
        A.store.Ref(ptr + 20, x["url"]);
      }
    },
    "load_GetAllArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["domain"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      x["path"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 24)) {
        x["secure"] = A.load.Bool(ptr + 12);
      } else {
        delete x["secure"];
      }
      if (A.load.Bool(ptr + 25)) {
        x["session"] = A.load.Bool(ptr + 13);
      } else {
        delete x["session"];
      }
      x["storeId"] = A.load.Ref(ptr + 16, undefined);
      x["url"] = A.load.Ref(ptr + 20, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OnChangedCause": (ref: heap.Ref<string>): number => {
      const idx = ["evicted", "expired", "explicit", "expired_overwrite", "overwrite"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OnChangedArgChangeInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 55, false);
        A.store.Enum(ptr + 0, -1);

        A.store.Bool(ptr + 8 + 45, false);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Bool(ptr + 8 + 44, false);
        A.store.Float64(ptr + 8 + 8, 0);
        A.store.Bool(ptr + 8 + 16, false);
        A.store.Bool(ptr + 8 + 17, false);
        A.store.Ref(ptr + 8 + 20, undefined);
        A.store.Ref(ptr + 8 + 24, undefined);
        A.store.Enum(ptr + 8 + 28, -1);
        A.store.Bool(ptr + 8 + 32, false);
        A.store.Bool(ptr + 8 + 33, false);
        A.store.Ref(ptr + 8 + 36, undefined);
        A.store.Ref(ptr + 8 + 40, undefined);
        A.store.Bool(ptr + 54, false);
      } else {
        A.store.Bool(ptr + 55, true);
        A.store.Enum(
          ptr + 0,
          ["evicted", "expired", "explicit", "expired_overwrite", "overwrite"].indexOf(x["cause"] as string)
        );

        if (typeof x["cookie"] === "undefined") {
          A.store.Bool(ptr + 8 + 45, false);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Bool(ptr + 8 + 44, false);
          A.store.Float64(ptr + 8 + 8, 0);
          A.store.Bool(ptr + 8 + 16, false);
          A.store.Bool(ptr + 8 + 17, false);
          A.store.Ref(ptr + 8 + 20, undefined);
          A.store.Ref(ptr + 8 + 24, undefined);
          A.store.Enum(ptr + 8 + 28, -1);
          A.store.Bool(ptr + 8 + 32, false);
          A.store.Bool(ptr + 8 + 33, false);
          A.store.Ref(ptr + 8 + 36, undefined);
          A.store.Ref(ptr + 8 + 40, undefined);
        } else {
          A.store.Bool(ptr + 8 + 45, true);
          A.store.Ref(ptr + 8 + 0, x["cookie"]["domain"]);
          A.store.Bool(ptr + 8 + 44, "expirationDate" in x["cookie"] ? true : false);
          A.store.Float64(
            ptr + 8 + 8,
            x["cookie"]["expirationDate"] === undefined ? 0 : (x["cookie"]["expirationDate"] as number)
          );
          A.store.Bool(ptr + 8 + 16, x["cookie"]["hostOnly"] ? true : false);
          A.store.Bool(ptr + 8 + 17, x["cookie"]["httpOnly"] ? true : false);
          A.store.Ref(ptr + 8 + 20, x["cookie"]["name"]);
          A.store.Ref(ptr + 8 + 24, x["cookie"]["path"]);
          A.store.Enum(
            ptr + 8 + 28,
            ["no_restriction", "lax", "strict", "unspecified"].indexOf(x["cookie"]["sameSite"] as string)
          );
          A.store.Bool(ptr + 8 + 32, x["cookie"]["secure"] ? true : false);
          A.store.Bool(ptr + 8 + 33, x["cookie"]["session"] ? true : false);
          A.store.Ref(ptr + 8 + 36, x["cookie"]["storeId"]);
          A.store.Ref(ptr + 8 + 40, x["cookie"]["value"]);
        }
        A.store.Bool(ptr + 54, x["removed"] ? true : false);
      }
    },
    "load_OnChangedArgChangeInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["cause"] = A.load.Enum(ptr + 0, ["evicted", "expired", "explicit", "expired_overwrite", "overwrite"]);
      if (A.load.Bool(ptr + 8 + 45)) {
        x["cookie"] = {};
        x["cookie"]["domain"] = A.load.Ref(ptr + 8 + 0, undefined);
        if (A.load.Bool(ptr + 8 + 44)) {
          x["cookie"]["expirationDate"] = A.load.Float64(ptr + 8 + 8);
        } else {
          delete x["cookie"]["expirationDate"];
        }
        x["cookie"]["hostOnly"] = A.load.Bool(ptr + 8 + 16);
        x["cookie"]["httpOnly"] = A.load.Bool(ptr + 8 + 17);
        x["cookie"]["name"] = A.load.Ref(ptr + 8 + 20, undefined);
        x["cookie"]["path"] = A.load.Ref(ptr + 8 + 24, undefined);
        x["cookie"]["sameSite"] = A.load.Enum(ptr + 8 + 28, ["no_restriction", "lax", "strict", "unspecified"]);
        x["cookie"]["secure"] = A.load.Bool(ptr + 8 + 32);
        x["cookie"]["session"] = A.load.Bool(ptr + 8 + 33);
        x["cookie"]["storeId"] = A.load.Ref(ptr + 8 + 36, undefined);
        x["cookie"]["value"] = A.load.Ref(ptr + 8 + 40, undefined);
      } else {
        delete x["cookie"];
      }
      x["removed"] = A.load.Bool(ptr + 54);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RemoveReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["storeId"]);
        A.store.Ref(ptr + 8, x["url"]);
      }
    },
    "load_RemoveReturnType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["storeId"] = A.load.Ref(ptr + 4, undefined);
      x["url"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 51, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 48, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 49, false);
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Enum(ptr + 28, -1);
        A.store.Bool(ptr + 50, false);
        A.store.Bool(ptr + 32, false);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
      } else {
        A.store.Bool(ptr + 51, true);
        A.store.Ref(ptr + 0, x["domain"]);
        A.store.Bool(ptr + 48, "expirationDate" in x ? true : false);
        A.store.Float64(ptr + 8, x["expirationDate"] === undefined ? 0 : (x["expirationDate"] as number));
        A.store.Bool(ptr + 49, "httpOnly" in x ? true : false);
        A.store.Bool(ptr + 16, x["httpOnly"] ? true : false);
        A.store.Ref(ptr + 20, x["name"]);
        A.store.Ref(ptr + 24, x["path"]);
        A.store.Enum(ptr + 28, ["no_restriction", "lax", "strict", "unspecified"].indexOf(x["sameSite"] as string));
        A.store.Bool(ptr + 50, "secure" in x ? true : false);
        A.store.Bool(ptr + 32, x["secure"] ? true : false);
        A.store.Ref(ptr + 36, x["storeId"]);
        A.store.Ref(ptr + 40, x["url"]);
        A.store.Ref(ptr + 44, x["value"]);
      }
    },
    "load_SetArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["domain"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 48)) {
        x["expirationDate"] = A.load.Float64(ptr + 8);
      } else {
        delete x["expirationDate"];
      }
      if (A.load.Bool(ptr + 49)) {
        x["httpOnly"] = A.load.Bool(ptr + 16);
      } else {
        delete x["httpOnly"];
      }
      x["name"] = A.load.Ref(ptr + 20, undefined);
      x["path"] = A.load.Ref(ptr + 24, undefined);
      x["sameSite"] = A.load.Enum(ptr + 28, ["no_restriction", "lax", "strict", "unspecified"]);
      if (A.load.Bool(ptr + 50)) {
        x["secure"] = A.load.Bool(ptr + 32);
      } else {
        delete x["secure"];
      }
      x["storeId"] = A.load.Ref(ptr + 36, undefined);
      x["url"] = A.load.Ref(ptr + 40, undefined);
      x["value"] = A.load.Ref(ptr + 44, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Get": (): heap.Ref<boolean> => {
      if (WEBEXT?.cookies && "get" in WEBEXT?.cookies) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Get": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.cookies.get);
    },
    "call_Get": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["name"] = A.load.Ref(details + 0, undefined);
      details_ffi["storeId"] = A.load.Ref(details + 4, undefined);
      details_ffi["url"] = A.load.Ref(details + 8, undefined);

      const _ret = WEBEXT.cookies.get(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Get": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["name"] = A.load.Ref(details + 0, undefined);
        details_ffi["storeId"] = A.load.Ref(details + 4, undefined);
        details_ffi["url"] = A.load.Ref(details + 8, undefined);

        const _ret = WEBEXT.cookies.get(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAll": (): heap.Ref<boolean> => {
      if (WEBEXT?.cookies && "getAll" in WEBEXT?.cookies) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.cookies.getAll);
    },
    "call_GetAll": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["domain"] = A.load.Ref(details + 0, undefined);
      details_ffi["name"] = A.load.Ref(details + 4, undefined);
      details_ffi["path"] = A.load.Ref(details + 8, undefined);
      if (A.load.Bool(details + 24)) {
        details_ffi["secure"] = A.load.Bool(details + 12);
      }
      if (A.load.Bool(details + 25)) {
        details_ffi["session"] = A.load.Bool(details + 13);
      }
      details_ffi["storeId"] = A.load.Ref(details + 16, undefined);
      details_ffi["url"] = A.load.Ref(details + 20, undefined);

      const _ret = WEBEXT.cookies.getAll(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAll": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["domain"] = A.load.Ref(details + 0, undefined);
        details_ffi["name"] = A.load.Ref(details + 4, undefined);
        details_ffi["path"] = A.load.Ref(details + 8, undefined);
        if (A.load.Bool(details + 24)) {
          details_ffi["secure"] = A.load.Bool(details + 12);
        }
        if (A.load.Bool(details + 25)) {
          details_ffi["session"] = A.load.Bool(details + 13);
        }
        details_ffi["storeId"] = A.load.Ref(details + 16, undefined);
        details_ffi["url"] = A.load.Ref(details + 20, undefined);

        const _ret = WEBEXT.cookies.getAll(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAllCookieStores": (): heap.Ref<boolean> => {
      if (WEBEXT?.cookies && "getAllCookieStores" in WEBEXT?.cookies) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAllCookieStores": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.cookies.getAllCookieStores);
    },
    "call_GetAllCookieStores": (retPtr: Pointer): void => {
      const _ret = WEBEXT.cookies.getAllCookieStores();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAllCookieStores": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.cookies.getAllCookieStores();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.cookies?.onChanged && "addListener" in WEBEXT?.cookies?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.cookies.onChanged.addListener);
    },
    "call_OnChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.cookies.onChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.cookies.onChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.cookies?.onChanged && "removeListener" in WEBEXT?.cookies?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.cookies.onChanged.removeListener);
    },
    "call_OffChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.cookies.onChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.cookies.onChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.cookies?.onChanged && "hasListener" in WEBEXT?.cookies?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.cookies.onChanged.hasListener);
    },
    "call_HasOnChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.cookies.onChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.cookies.onChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Remove": (): heap.Ref<boolean> => {
      if (WEBEXT?.cookies && "remove" in WEBEXT?.cookies) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Remove": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.cookies.remove);
    },
    "call_Remove": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["name"] = A.load.Ref(details + 0, undefined);
      details_ffi["storeId"] = A.load.Ref(details + 4, undefined);
      details_ffi["url"] = A.load.Ref(details + 8, undefined);

      const _ret = WEBEXT.cookies.remove(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Remove": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["name"] = A.load.Ref(details + 0, undefined);
        details_ffi["storeId"] = A.load.Ref(details + 4, undefined);
        details_ffi["url"] = A.load.Ref(details + 8, undefined);

        const _ret = WEBEXT.cookies.remove(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Set": (): heap.Ref<boolean> => {
      if (WEBEXT?.cookies && "set" in WEBEXT?.cookies) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Set": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.cookies.set);
    },
    "call_Set": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["domain"] = A.load.Ref(details + 0, undefined);
      if (A.load.Bool(details + 48)) {
        details_ffi["expirationDate"] = A.load.Float64(details + 8);
      }
      if (A.load.Bool(details + 49)) {
        details_ffi["httpOnly"] = A.load.Bool(details + 16);
      }
      details_ffi["name"] = A.load.Ref(details + 20, undefined);
      details_ffi["path"] = A.load.Ref(details + 24, undefined);
      details_ffi["sameSite"] = A.load.Enum(details + 28, ["no_restriction", "lax", "strict", "unspecified"]);
      if (A.load.Bool(details + 50)) {
        details_ffi["secure"] = A.load.Bool(details + 32);
      }
      details_ffi["storeId"] = A.load.Ref(details + 36, undefined);
      details_ffi["url"] = A.load.Ref(details + 40, undefined);
      details_ffi["value"] = A.load.Ref(details + 44, undefined);

      const _ret = WEBEXT.cookies.set(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Set": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["domain"] = A.load.Ref(details + 0, undefined);
        if (A.load.Bool(details + 48)) {
          details_ffi["expirationDate"] = A.load.Float64(details + 8);
        }
        if (A.load.Bool(details + 49)) {
          details_ffi["httpOnly"] = A.load.Bool(details + 16);
        }
        details_ffi["name"] = A.load.Ref(details + 20, undefined);
        details_ffi["path"] = A.load.Ref(details + 24, undefined);
        details_ffi["sameSite"] = A.load.Enum(details + 28, ["no_restriction", "lax", "strict", "unspecified"]);
        if (A.load.Bool(details + 50)) {
          details_ffi["secure"] = A.load.Bool(details + 32);
        }
        details_ffi["storeId"] = A.load.Ref(details + 36, undefined);
        details_ffi["url"] = A.load.Ref(details + 40, undefined);
        details_ffi["value"] = A.load.Ref(details + 44, undefined);

        const _ret = WEBEXT.cookies.set(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
