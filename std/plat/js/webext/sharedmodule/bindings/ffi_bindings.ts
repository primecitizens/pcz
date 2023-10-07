import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/sharedmodule", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Export": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["allowlist"]);
      }
    },
    "load_Export": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["allowlist"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Import": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["minimum_version"]);
      }
    },
    "load_Import": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["minimum_version"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManifestKeys": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 4 + 4, false);
        A.store.Ref(ptr + 4 + 0, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Ref(ptr + 0, x["import"]);

        if (typeof x["export"] === "undefined") {
          A.store.Bool(ptr + 4 + 4, false);
          A.store.Ref(ptr + 4 + 0, undefined);
        } else {
          A.store.Bool(ptr + 4 + 4, true);
          A.store.Ref(ptr + 4 + 0, x["export"]["allowlist"]);
        }
      }
    },
    "load_ManifestKeys": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["import"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 4 + 4)) {
        x["export"] = {};
        x["export"]["allowlist"] = A.load.Ref(ptr + 4 + 0, undefined);
      } else {
        delete x["export"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
  };
});
