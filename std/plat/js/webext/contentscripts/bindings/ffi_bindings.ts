import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/contentscripts", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_RunAt": (ref: heap.Ref<string>): number => {
      const idx = ["document_idle", "document_start", "document_end"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ContentScript": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 39, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 36, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 37, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 38, false);
        A.store.Bool(ptr + 18, false);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Enum(ptr + 28, -1);
        A.store.Enum(ptr + 32, -1);
      } else {
        A.store.Bool(ptr + 39, true);
        A.store.Ref(ptr + 0, x["matches"]);
        A.store.Ref(ptr + 4, x["exclude_matches"]);
        A.store.Ref(ptr + 8, x["css"]);
        A.store.Ref(ptr + 12, x["js"]);
        A.store.Bool(ptr + 36, "all_frames" in x ? true : false);
        A.store.Bool(ptr + 16, x["all_frames"] ? true : false);
        A.store.Bool(ptr + 37, "match_origin_as_fallback" in x ? true : false);
        A.store.Bool(ptr + 17, x["match_origin_as_fallback"] ? true : false);
        A.store.Bool(ptr + 38, "match_about_blank" in x ? true : false);
        A.store.Bool(ptr + 18, x["match_about_blank"] ? true : false);
        A.store.Ref(ptr + 20, x["include_globs"]);
        A.store.Ref(ptr + 24, x["exclude_globs"]);
        A.store.Enum(ptr + 28, ["document_idle", "document_start", "document_end"].indexOf(x["run_at"] as string));
        A.store.Enum(ptr + 32, ["ISOLATED", "MAIN"].indexOf(x["world"] as string));
      }
    },
    "load_ContentScript": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["matches"] = A.load.Ref(ptr + 0, undefined);
      x["exclude_matches"] = A.load.Ref(ptr + 4, undefined);
      x["css"] = A.load.Ref(ptr + 8, undefined);
      x["js"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 36)) {
        x["all_frames"] = A.load.Bool(ptr + 16);
      } else {
        delete x["all_frames"];
      }
      if (A.load.Bool(ptr + 37)) {
        x["match_origin_as_fallback"] = A.load.Bool(ptr + 17);
      } else {
        delete x["match_origin_as_fallback"];
      }
      if (A.load.Bool(ptr + 38)) {
        x["match_about_blank"] = A.load.Bool(ptr + 18);
      } else {
        delete x["match_about_blank"];
      }
      x["include_globs"] = A.load.Ref(ptr + 20, undefined);
      x["exclude_globs"] = A.load.Ref(ptr + 24, undefined);
      x["run_at"] = A.load.Enum(ptr + 28, ["document_idle", "document_start", "document_end"]);
      x["world"] = A.load.Enum(ptr + 32, ["ISOLATED", "MAIN"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManifestKeys": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["content_scripts"]);
      }
    },
    "load_ManifestKeys": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["content_scripts"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
  };
});
