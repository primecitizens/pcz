import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/declarativecontent", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_PageStateMatcherInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeContent.PageStateMatcher"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PageStateMatcher": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 94, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Bool(ptr + 93, false);
        A.store.Bool(ptr + 8, false);

        A.store.Bool(ptr + 12 + 0, false);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Ref(ptr + 12 + 0, undefined);
      } else {
        A.store.Bool(ptr + 94, true);
        A.store.Ref(ptr + 0, x["css"]);
        A.store.Enum(ptr + 4, ["declarativeContent.PageStateMatcher"].indexOf(x["instanceType"] as string));
        A.store.Bool(ptr + 93, "isBookmarked" in x ? true : false);
        A.store.Bool(ptr + 8, x["isBookmarked"] ? true : false);

        if (typeof x["pageUrl"] === "undefined") {
          A.store.Bool(ptr + 12 + 0, false);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Ref(ptr + 12 + 0, undefined);
        } else {
          A.store.Bool(ptr + 12 + 0, true);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["hostContains"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["hostEquals"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["hostPrefix"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["hostSuffix"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["originAndPathMatches"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["pathContains"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["pathEquals"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["pathPrefix"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["pathSuffix"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["ports"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["queryContains"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["queryEquals"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["queryPrefix"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["querySuffix"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["schemes"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["urlContains"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["urlEquals"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["urlMatches"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["urlPrefix"]);
          A.store.Ref(ptr + 12 + 0, x["pageUrl"]["urlSuffix"]);
        }
      }
    },
    "load_PageStateMatcher": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["css"] = A.load.Ref(ptr + 0, undefined);
      x["instanceType"] = A.load.Enum(ptr + 4, ["declarativeContent.PageStateMatcher"]);
      if (A.load.Bool(ptr + 93)) {
        x["isBookmarked"] = A.load.Bool(ptr + 8);
      } else {
        delete x["isBookmarked"];
      }
      if (A.load.Bool(ptr + 12 + 0)) {
        x["pageUrl"] = {};
        x["pageUrl"]["hostContains"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["hostEquals"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["hostPrefix"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["hostSuffix"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["originAndPathMatches"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["pathContains"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["pathEquals"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["pathPrefix"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["pathSuffix"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["ports"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["queryContains"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["queryEquals"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["queryPrefix"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["querySuffix"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["schemes"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["urlContains"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["urlEquals"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["urlMatches"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["urlPrefix"] = A.load.Ref(ptr + 12 + 0, undefined);
        x["pageUrl"]["urlSuffix"] = A.load.Ref(ptr + 12 + 0, undefined);
      } else {
        delete x["pageUrl"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RequestContentScriptInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeContent.RequestContentScript"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RequestContentScript": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
      } else {
        A.store.Bool(ptr + 19, true);
        A.store.Bool(ptr + 17, "allFrames" in x ? true : false);
        A.store.Bool(ptr + 0, x["allFrames"] ? true : false);
        A.store.Ref(ptr + 4, x["css"]);
        A.store.Enum(ptr + 8, ["declarativeContent.RequestContentScript"].indexOf(x["instanceType"] as string));
        A.store.Ref(ptr + 12, x["js"]);
        A.store.Bool(ptr + 18, "matchAboutBlank" in x ? true : false);
        A.store.Bool(ptr + 16, x["matchAboutBlank"] ? true : false);
      }
    },
    "load_RequestContentScript": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 17)) {
        x["allFrames"] = A.load.Bool(ptr + 0);
      } else {
        delete x["allFrames"];
      }
      x["css"] = A.load.Ref(ptr + 4, undefined);
      x["instanceType"] = A.load.Enum(ptr + 8, ["declarativeContent.RequestContentScript"]);
      x["js"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 18)) {
        x["matchAboutBlank"] = A.load.Bool(ptr + 16);
      } else {
        delete x["matchAboutBlank"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_SetIconInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeContent.SetIcon"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SetIcon": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["imageData"]);
        A.store.Enum(ptr + 4, ["declarativeContent.SetIcon"].indexOf(x["instanceType"] as string));
      }
    },
    "load_SetIcon": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["imageData"] = A.load.Ref(ptr + 0, undefined);
      x["instanceType"] = A.load.Enum(ptr + 4, ["declarativeContent.SetIcon"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ShowActionInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeContent.ShowAction"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ShowAction": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["declarativeContent.ShowAction"].indexOf(x["instanceType"] as string));
      }
    },
    "load_ShowAction": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["instanceType"] = A.load.Enum(ptr + 0, ["declarativeContent.ShowAction"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ShowPageActionInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeContent.ShowPageAction"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ShowPageAction": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["declarativeContent.ShowPageAction"].indexOf(x["instanceType"] as string));
      }
    },
    "load_ShowPageAction": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["instanceType"] = A.load.Enum(ptr + 0, ["declarativeContent.ShowPageAction"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_OnPageChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeContent?.onPageChanged && "addListener" in WEBEXT?.declarativeContent?.onPageChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPageChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeContent.onPageChanged.addListener);
    },
    "call_OnPageChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.declarativeContent.onPageChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnPageChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.declarativeContent.onPageChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPageChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeContent?.onPageChanged && "removeListener" in WEBEXT?.declarativeContent?.onPageChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPageChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeContent.onPageChanged.removeListener);
    },
    "call_OffPageChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.declarativeContent.onPageChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffPageChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.declarativeContent.onPageChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPageChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeContent?.onPageChanged && "hasListener" in WEBEXT?.declarativeContent?.onPageChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPageChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeContent.onPageChanged.hasListener);
    },
    "call_HasOnPageChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.declarativeContent.onPageChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPageChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.declarativeContent.onPageChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
