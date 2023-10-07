import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/oauth2", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_OAuth2Info": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "auto_approve" in x ? true : false);
        A.store.Bool(ptr + 0, x["auto_approve"] ? true : false);
        A.store.Ref(ptr + 4, x["client_id"]);
        A.store.Ref(ptr + 8, x["scopes"]);
      }
    },
    "load_OAuth2Info": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["auto_approve"] = A.load.Bool(ptr + 0);
      } else {
        delete x["auto_approve"];
      }
      x["client_id"] = A.load.Ref(ptr + 4, undefined);
      x["scopes"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManifestKeys": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);

        A.store.Bool(ptr + 0 + 13, false);
        A.store.Bool(ptr + 0 + 12, false);
        A.store.Bool(ptr + 0 + 0, false);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Ref(ptr + 0 + 8, undefined);
      } else {
        A.store.Bool(ptr + 14, true);

        if (typeof x["oauth2"] === "undefined") {
          A.store.Bool(ptr + 0 + 13, false);
          A.store.Bool(ptr + 0 + 12, false);
          A.store.Bool(ptr + 0 + 0, false);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Ref(ptr + 0 + 8, undefined);
        } else {
          A.store.Bool(ptr + 0 + 13, true);
          A.store.Bool(ptr + 0 + 12, "auto_approve" in x["oauth2"] ? true : false);
          A.store.Bool(ptr + 0 + 0, x["oauth2"]["auto_approve"] ? true : false);
          A.store.Ref(ptr + 0 + 4, x["oauth2"]["client_id"]);
          A.store.Ref(ptr + 0 + 8, x["oauth2"]["scopes"]);
        }
      }
    },
    "load_ManifestKeys": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 13)) {
        x["oauth2"] = {};
        if (A.load.Bool(ptr + 0 + 12)) {
          x["oauth2"]["auto_approve"] = A.load.Bool(ptr + 0 + 0);
        } else {
          delete x["oauth2"]["auto_approve"];
        }
        x["oauth2"]["client_id"] = A.load.Ref(ptr + 0 + 4, undefined);
        x["oauth2"]["scopes"] = A.load.Ref(ptr + 0 + 8, undefined);
      } else {
        delete x["oauth2"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
  };
});
