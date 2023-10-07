import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/manifesttypes", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_ChromeSettingsOverridesFieldSearchProvider": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 57, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 20, false);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Bool(ptr + 56, false);
        A.store.Int64(ptr + 32, 0);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Ref(ptr + 48, undefined);
        A.store.Ref(ptr + 52, undefined);
      } else {
        A.store.Bool(ptr + 57, true);
        A.store.Ref(ptr + 0, x["alternate_urls"]);
        A.store.Ref(ptr + 4, x["encoding"]);
        A.store.Ref(ptr + 8, x["favicon_url"]);
        A.store.Ref(ptr + 12, x["image_url"]);
        A.store.Ref(ptr + 16, x["image_url_post_params"]);
        A.store.Bool(ptr + 20, x["is_default"] ? true : false);
        A.store.Ref(ptr + 24, x["keyword"]);
        A.store.Ref(ptr + 28, x["name"]);
        A.store.Bool(ptr + 56, "prepopulated_id" in x ? true : false);
        A.store.Int64(ptr + 32, x["prepopulated_id"] === undefined ? 0 : (x["prepopulated_id"] as number));
        A.store.Ref(ptr + 40, x["search_url"]);
        A.store.Ref(ptr + 44, x["search_url_post_params"]);
        A.store.Ref(ptr + 48, x["suggest_url"]);
        A.store.Ref(ptr + 52, x["suggest_url_post_params"]);
      }
    },
    "load_ChromeSettingsOverridesFieldSearchProvider": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["alternate_urls"] = A.load.Ref(ptr + 0, undefined);
      x["encoding"] = A.load.Ref(ptr + 4, undefined);
      x["favicon_url"] = A.load.Ref(ptr + 8, undefined);
      x["image_url"] = A.load.Ref(ptr + 12, undefined);
      x["image_url_post_params"] = A.load.Ref(ptr + 16, undefined);
      x["is_default"] = A.load.Bool(ptr + 20);
      x["keyword"] = A.load.Ref(ptr + 24, undefined);
      x["name"] = A.load.Ref(ptr + 28, undefined);
      if (A.load.Bool(ptr + 56)) {
        x["prepopulated_id"] = A.load.Int64(ptr + 32);
      } else {
        delete x["prepopulated_id"];
      }
      x["search_url"] = A.load.Ref(ptr + 40, undefined);
      x["search_url_post_params"] = A.load.Ref(ptr + 44, undefined);
      x["suggest_url"] = A.load.Ref(ptr + 48, undefined);
      x["suggest_url_post_params"] = A.load.Ref(ptr + 52, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ChromeSettingsOverrides": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 72, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 8 + 57, false);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Ref(ptr + 8 + 4, undefined);
        A.store.Ref(ptr + 8 + 8, undefined);
        A.store.Ref(ptr + 8 + 12, undefined);
        A.store.Ref(ptr + 8 + 16, undefined);
        A.store.Bool(ptr + 8 + 20, false);
        A.store.Ref(ptr + 8 + 24, undefined);
        A.store.Ref(ptr + 8 + 28, undefined);
        A.store.Bool(ptr + 8 + 56, false);
        A.store.Int64(ptr + 8 + 32, 0);
        A.store.Ref(ptr + 8 + 40, undefined);
        A.store.Ref(ptr + 8 + 44, undefined);
        A.store.Ref(ptr + 8 + 48, undefined);
        A.store.Ref(ptr + 8 + 52, undefined);
        A.store.Ref(ptr + 68, undefined);
      } else {
        A.store.Bool(ptr + 72, true);
        A.store.Ref(ptr + 0, x["homepage"]);

        if (typeof x["search_provider"] === "undefined") {
          A.store.Bool(ptr + 8 + 57, false);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Ref(ptr + 8 + 4, undefined);
          A.store.Ref(ptr + 8 + 8, undefined);
          A.store.Ref(ptr + 8 + 12, undefined);
          A.store.Ref(ptr + 8 + 16, undefined);
          A.store.Bool(ptr + 8 + 20, false);
          A.store.Ref(ptr + 8 + 24, undefined);
          A.store.Ref(ptr + 8 + 28, undefined);
          A.store.Bool(ptr + 8 + 56, false);
          A.store.Int64(ptr + 8 + 32, 0);
          A.store.Ref(ptr + 8 + 40, undefined);
          A.store.Ref(ptr + 8 + 44, undefined);
          A.store.Ref(ptr + 8 + 48, undefined);
          A.store.Ref(ptr + 8 + 52, undefined);
        } else {
          A.store.Bool(ptr + 8 + 57, true);
          A.store.Ref(ptr + 8 + 0, x["search_provider"]["alternate_urls"]);
          A.store.Ref(ptr + 8 + 4, x["search_provider"]["encoding"]);
          A.store.Ref(ptr + 8 + 8, x["search_provider"]["favicon_url"]);
          A.store.Ref(ptr + 8 + 12, x["search_provider"]["image_url"]);
          A.store.Ref(ptr + 8 + 16, x["search_provider"]["image_url_post_params"]);
          A.store.Bool(ptr + 8 + 20, x["search_provider"]["is_default"] ? true : false);
          A.store.Ref(ptr + 8 + 24, x["search_provider"]["keyword"]);
          A.store.Ref(ptr + 8 + 28, x["search_provider"]["name"]);
          A.store.Bool(ptr + 8 + 56, "prepopulated_id" in x["search_provider"] ? true : false);
          A.store.Int64(
            ptr + 8 + 32,
            x["search_provider"]["prepopulated_id"] === undefined
              ? 0
              : (x["search_provider"]["prepopulated_id"] as number)
          );
          A.store.Ref(ptr + 8 + 40, x["search_provider"]["search_url"]);
          A.store.Ref(ptr + 8 + 44, x["search_provider"]["search_url_post_params"]);
          A.store.Ref(ptr + 8 + 48, x["search_provider"]["suggest_url"]);
          A.store.Ref(ptr + 8 + 52, x["search_provider"]["suggest_url_post_params"]);
        }
        A.store.Ref(ptr + 68, x["startup_pages"]);
      }
    },
    "load_ChromeSettingsOverrides": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["homepage"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 8 + 57)) {
        x["search_provider"] = {};
        x["search_provider"]["alternate_urls"] = A.load.Ref(ptr + 8 + 0, undefined);
        x["search_provider"]["encoding"] = A.load.Ref(ptr + 8 + 4, undefined);
        x["search_provider"]["favicon_url"] = A.load.Ref(ptr + 8 + 8, undefined);
        x["search_provider"]["image_url"] = A.load.Ref(ptr + 8 + 12, undefined);
        x["search_provider"]["image_url_post_params"] = A.load.Ref(ptr + 8 + 16, undefined);
        x["search_provider"]["is_default"] = A.load.Bool(ptr + 8 + 20);
        x["search_provider"]["keyword"] = A.load.Ref(ptr + 8 + 24, undefined);
        x["search_provider"]["name"] = A.load.Ref(ptr + 8 + 28, undefined);
        if (A.load.Bool(ptr + 8 + 56)) {
          x["search_provider"]["prepopulated_id"] = A.load.Int64(ptr + 8 + 32);
        } else {
          delete x["search_provider"]["prepopulated_id"];
        }
        x["search_provider"]["search_url"] = A.load.Ref(ptr + 8 + 40, undefined);
        x["search_provider"]["search_url_post_params"] = A.load.Ref(ptr + 8 + 44, undefined);
        x["search_provider"]["suggest_url"] = A.load.Ref(ptr + 8 + 48, undefined);
        x["search_provider"]["suggest_url_post_params"] = A.load.Ref(ptr + 8 + 52, undefined);
      } else {
        delete x["search_provider"];
      }
      x["startup_pages"] = A.load.Ref(ptr + 68, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_FileSystemProviderSource": (ref: heap.Ref<string>): number => {
      const idx = ["file", "device", "network"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_FileSystemProviderCapabilities": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 1, false);
        A.store.Enum(ptr + 4, -1);
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Bool(ptr + 9, "configurable" in x ? true : false);
        A.store.Bool(ptr + 0, x["configurable"] ? true : false);
        A.store.Bool(ptr + 10, "multiple_mounts" in x ? true : false);
        A.store.Bool(ptr + 1, x["multiple_mounts"] ? true : false);
        A.store.Enum(ptr + 4, ["file", "device", "network"].indexOf(x["source"] as string));
        A.store.Bool(ptr + 11, "watchable" in x ? true : false);
        A.store.Bool(ptr + 8, x["watchable"] ? true : false);
      }
    },
    "load_FileSystemProviderCapabilities": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 9)) {
        x["configurable"] = A.load.Bool(ptr + 0);
      } else {
        delete x["configurable"];
      }
      if (A.load.Bool(ptr + 10)) {
        x["multiple_mounts"] = A.load.Bool(ptr + 1);
      } else {
        delete x["multiple_mounts"];
      }
      x["source"] = A.load.Enum(ptr + 4, ["file", "device", "network"]);
      if (A.load.Bool(ptr + 11)) {
        x["watchable"] = A.load.Bool(ptr + 8);
      } else {
        delete x["watchable"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
  };
});
