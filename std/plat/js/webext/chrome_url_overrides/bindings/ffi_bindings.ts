import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/chrome_url_overrides", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_UrlOverrideInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Ref(ptr + 0, x["newtab"]);
        A.store.Ref(ptr + 4, x["bookmarks"]);
        A.store.Ref(ptr + 8, x["history"]);
        A.store.Ref(ptr + 12, x["activationmessage"]);
        A.store.Ref(ptr + 16, x["keyboard"]);
      }
    },
    "load_UrlOverrideInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["newtab"] = A.load.Ref(ptr + 0, undefined);
      x["bookmarks"] = A.load.Ref(ptr + 4, undefined);
      x["history"] = A.load.Ref(ptr + 8, undefined);
      x["activationmessage"] = A.load.Ref(ptr + 12, undefined);
      x["keyboard"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManifestKeys": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);

        A.store.Bool(ptr + 0 + 20, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Ref(ptr + 0 + 8, undefined);
        A.store.Ref(ptr + 0 + 12, undefined);
        A.store.Ref(ptr + 0 + 16, undefined);
      } else {
        A.store.Bool(ptr + 21, true);

        if (typeof x["chrome_url_overrides"] === "undefined") {
          A.store.Bool(ptr + 0 + 20, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Ref(ptr + 0 + 8, undefined);
          A.store.Ref(ptr + 0 + 12, undefined);
          A.store.Ref(ptr + 0 + 16, undefined);
        } else {
          A.store.Bool(ptr + 0 + 20, true);
          A.store.Ref(ptr + 0 + 0, x["chrome_url_overrides"]["newtab"]);
          A.store.Ref(ptr + 0 + 4, x["chrome_url_overrides"]["bookmarks"]);
          A.store.Ref(ptr + 0 + 8, x["chrome_url_overrides"]["history"]);
          A.store.Ref(ptr + 0 + 12, x["chrome_url_overrides"]["activationmessage"]);
          A.store.Ref(ptr + 0 + 16, x["chrome_url_overrides"]["keyboard"]);
        }
      }
    },
    "load_ManifestKeys": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 20)) {
        x["chrome_url_overrides"] = {};
        x["chrome_url_overrides"]["newtab"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["chrome_url_overrides"]["bookmarks"] = A.load.Ref(ptr + 0 + 4, undefined);
        x["chrome_url_overrides"]["history"] = A.load.Ref(ptr + 0 + 8, undefined);
        x["chrome_url_overrides"]["activationmessage"] = A.load.Ref(ptr + 0 + 12, undefined);
        x["chrome_url_overrides"]["keyboard"] = A.load.Ref(ptr + 0 + 16, undefined);
      } else {
        delete x["chrome_url_overrides"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
  };
});
