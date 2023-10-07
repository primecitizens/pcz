import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/crossoriginisolation", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_ResponseHeader": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["value"]);
      }
    },
    "load_ResponseHeader": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["value"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManifestKeys": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);

        A.store.Bool(ptr + 0 + 4, false);
        A.store.Ref(ptr + 0 + 0, undefined);

        A.store.Bool(ptr + 8 + 4, false);
        A.store.Ref(ptr + 8 + 0, undefined);
      } else {
        A.store.Bool(ptr + 13, true);

        if (typeof x["cross_origin_embedder_policy"] === "undefined") {
          A.store.Bool(ptr + 0 + 4, false);
          A.store.Ref(ptr + 0 + 0, undefined);
        } else {
          A.store.Bool(ptr + 0 + 4, true);
          A.store.Ref(ptr + 0 + 0, x["cross_origin_embedder_policy"]["value"]);
        }

        if (typeof x["cross_origin_opener_policy"] === "undefined") {
          A.store.Bool(ptr + 8 + 4, false);
          A.store.Ref(ptr + 8 + 0, undefined);
        } else {
          A.store.Bool(ptr + 8 + 4, true);
          A.store.Ref(ptr + 8 + 0, x["cross_origin_opener_policy"]["value"]);
        }
      }
    },
    "load_ManifestKeys": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 4)) {
        x["cross_origin_embedder_policy"] = {};
        x["cross_origin_embedder_policy"]["value"] = A.load.Ref(ptr + 0 + 0, undefined);
      } else {
        delete x["cross_origin_embedder_policy"];
      }
      if (A.load.Bool(ptr + 8 + 4)) {
        x["cross_origin_opener_policy"] = {};
        x["cross_origin_opener_policy"]["value"] = A.load.Ref(ptr + 8 + 0, undefined);
      } else {
        delete x["cross_origin_opener_policy"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
  };
});
