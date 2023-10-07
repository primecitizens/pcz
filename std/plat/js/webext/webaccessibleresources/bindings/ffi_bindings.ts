import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/webaccessibleresources", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_WebAccessibleResource": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Ref(ptr + 0, x["resources"]);
        A.store.Ref(ptr + 4, x["matches"]);
        A.store.Ref(ptr + 8, x["extension_ids"]);
        A.store.Bool(ptr + 13, "use_dynamic_url" in x ? true : false);
        A.store.Bool(ptr + 12, x["use_dynamic_url"] ? true : false);
      }
    },
    "load_WebAccessibleResource": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["resources"] = A.load.Ref(ptr + 0, undefined);
      x["matches"] = A.load.Ref(ptr + 4, undefined);
      x["extension_ids"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 13)) {
        x["use_dynamic_url"] = A.load.Bool(ptr + 12);
      } else {
        delete x["use_dynamic_url"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManifestKeys": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["web_accessible_resources"]);
      }
    },
    "load_ManifestKeys": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["web_accessible_resources"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
  };
});
