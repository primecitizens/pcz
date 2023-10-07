import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/loginscreenui", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_ShowOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 6, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 6, true);
        A.store.Ref(ptr + 0, x["url"]);
        A.store.Bool(ptr + 5, "userCanClose" in x ? true : false);
        A.store.Bool(ptr + 4, x["userCanClose"] ? true : false);
      }
    },
    "load_ShowOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 5)) {
        x["userCanClose"] = A.load.Bool(ptr + 4);
      } else {
        delete x["userCanClose"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Close": (): heap.Ref<boolean> => {
      if (WEBEXT?.loginScreenUi && "close" in WEBEXT?.loginScreenUi) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Close": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.loginScreenUi.close);
    },
    "call_Close": (retPtr: Pointer): void => {
      const _ret = WEBEXT.loginScreenUi.close();
      A.store.Ref(retPtr, _ret);
    },
    "try_Close": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.loginScreenUi.close();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Show": (): heap.Ref<boolean> => {
      if (WEBEXT?.loginScreenUi && "show" in WEBEXT?.loginScreenUi) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Show": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.loginScreenUi.show);
    },
    "call_Show": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["url"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 5)) {
        options_ffi["userCanClose"] = A.load.Bool(options + 4);
      }

      const _ret = WEBEXT.loginScreenUi.show(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Show": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["url"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 5)) {
          options_ffi["userCanClose"] = A.load.Bool(options + 4);
        }

        const _ret = WEBEXT.loginScreenUi.show(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
