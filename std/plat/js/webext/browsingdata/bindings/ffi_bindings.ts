import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/browsingdata", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_DataTypeSet": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 30, false);
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 22, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 23, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 27, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 29, false);
        A.store.Bool(ptr + 14, false);
      } else {
        A.store.Bool(ptr + 30, true);
        A.store.Bool(ptr + 15, "appcache" in x ? true : false);
        A.store.Bool(ptr + 0, x["appcache"] ? true : false);
        A.store.Bool(ptr + 16, "cache" in x ? true : false);
        A.store.Bool(ptr + 1, x["cache"] ? true : false);
        A.store.Bool(ptr + 17, "cacheStorage" in x ? true : false);
        A.store.Bool(ptr + 2, x["cacheStorage"] ? true : false);
        A.store.Bool(ptr + 18, "cookies" in x ? true : false);
        A.store.Bool(ptr + 3, x["cookies"] ? true : false);
        A.store.Bool(ptr + 19, "downloads" in x ? true : false);
        A.store.Bool(ptr + 4, x["downloads"] ? true : false);
        A.store.Bool(ptr + 20, "fileSystems" in x ? true : false);
        A.store.Bool(ptr + 5, x["fileSystems"] ? true : false);
        A.store.Bool(ptr + 21, "formData" in x ? true : false);
        A.store.Bool(ptr + 6, x["formData"] ? true : false);
        A.store.Bool(ptr + 22, "history" in x ? true : false);
        A.store.Bool(ptr + 7, x["history"] ? true : false);
        A.store.Bool(ptr + 23, "indexedDB" in x ? true : false);
        A.store.Bool(ptr + 8, x["indexedDB"] ? true : false);
        A.store.Bool(ptr + 24, "localStorage" in x ? true : false);
        A.store.Bool(ptr + 9, x["localStorage"] ? true : false);
        A.store.Bool(ptr + 25, "passwords" in x ? true : false);
        A.store.Bool(ptr + 10, x["passwords"] ? true : false);
        A.store.Bool(ptr + 26, "pluginData" in x ? true : false);
        A.store.Bool(ptr + 11, x["pluginData"] ? true : false);
        A.store.Bool(ptr + 27, "serverBoundCertificates" in x ? true : false);
        A.store.Bool(ptr + 12, x["serverBoundCertificates"] ? true : false);
        A.store.Bool(ptr + 28, "serviceWorkers" in x ? true : false);
        A.store.Bool(ptr + 13, x["serviceWorkers"] ? true : false);
        A.store.Bool(ptr + 29, "webSQL" in x ? true : false);
        A.store.Bool(ptr + 14, x["webSQL"] ? true : false);
      }
    },
    "load_DataTypeSet": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 15)) {
        x["appcache"] = A.load.Bool(ptr + 0);
      } else {
        delete x["appcache"];
      }
      if (A.load.Bool(ptr + 16)) {
        x["cache"] = A.load.Bool(ptr + 1);
      } else {
        delete x["cache"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["cacheStorage"] = A.load.Bool(ptr + 2);
      } else {
        delete x["cacheStorage"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["cookies"] = A.load.Bool(ptr + 3);
      } else {
        delete x["cookies"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["downloads"] = A.load.Bool(ptr + 4);
      } else {
        delete x["downloads"];
      }
      if (A.load.Bool(ptr + 20)) {
        x["fileSystems"] = A.load.Bool(ptr + 5);
      } else {
        delete x["fileSystems"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["formData"] = A.load.Bool(ptr + 6);
      } else {
        delete x["formData"];
      }
      if (A.load.Bool(ptr + 22)) {
        x["history"] = A.load.Bool(ptr + 7);
      } else {
        delete x["history"];
      }
      if (A.load.Bool(ptr + 23)) {
        x["indexedDB"] = A.load.Bool(ptr + 8);
      } else {
        delete x["indexedDB"];
      }
      if (A.load.Bool(ptr + 24)) {
        x["localStorage"] = A.load.Bool(ptr + 9);
      } else {
        delete x["localStorage"];
      }
      if (A.load.Bool(ptr + 25)) {
        x["passwords"] = A.load.Bool(ptr + 10);
      } else {
        delete x["passwords"];
      }
      if (A.load.Bool(ptr + 26)) {
        x["pluginData"] = A.load.Bool(ptr + 11);
      } else {
        delete x["pluginData"];
      }
      if (A.load.Bool(ptr + 27)) {
        x["serverBoundCertificates"] = A.load.Bool(ptr + 12);
      } else {
        delete x["serverBoundCertificates"];
      }
      if (A.load.Bool(ptr + 28)) {
        x["serviceWorkers"] = A.load.Bool(ptr + 13);
      } else {
        delete x["serviceWorkers"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["webSQL"] = A.load.Bool(ptr + 14);
      } else {
        delete x["webSQL"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RemovalOptionsFieldOriginTypes": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 2, false);
      } else {
        A.store.Bool(ptr + 6, true);
        A.store.Bool(ptr + 3, "extension" in x ? true : false);
        A.store.Bool(ptr + 0, x["extension"] ? true : false);
        A.store.Bool(ptr + 4, "protectedWeb" in x ? true : false);
        A.store.Bool(ptr + 1, x["protectedWeb"] ? true : false);
        A.store.Bool(ptr + 5, "unprotectedWeb" in x ? true : false);
        A.store.Bool(ptr + 2, x["unprotectedWeb"] ? true : false);
      }
    },
    "load_RemovalOptionsFieldOriginTypes": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 3)) {
        x["extension"] = A.load.Bool(ptr + 0);
      } else {
        delete x["extension"];
      }
      if (A.load.Bool(ptr + 4)) {
        x["protectedWeb"] = A.load.Bool(ptr + 1);
      } else {
        delete x["protectedWeb"];
      }
      if (A.load.Bool(ptr + 5)) {
        x["unprotectedWeb"] = A.load.Bool(ptr + 2);
      } else {
        delete x["unprotectedWeb"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RemovalOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 4 + 6, false);
        A.store.Bool(ptr + 4 + 3, false);
        A.store.Bool(ptr + 4 + 0, false);
        A.store.Bool(ptr + 4 + 4, false);
        A.store.Bool(ptr + 4 + 1, false);
        A.store.Bool(ptr + 4 + 5, false);
        A.store.Bool(ptr + 4 + 2, false);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 24, false);
        A.store.Float64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 25, true);
        A.store.Ref(ptr + 0, x["excludeOrigins"]);

        if (typeof x["originTypes"] === "undefined") {
          A.store.Bool(ptr + 4 + 6, false);
          A.store.Bool(ptr + 4 + 3, false);
          A.store.Bool(ptr + 4 + 0, false);
          A.store.Bool(ptr + 4 + 4, false);
          A.store.Bool(ptr + 4 + 1, false);
          A.store.Bool(ptr + 4 + 5, false);
          A.store.Bool(ptr + 4 + 2, false);
        } else {
          A.store.Bool(ptr + 4 + 6, true);
          A.store.Bool(ptr + 4 + 3, "extension" in x["originTypes"] ? true : false);
          A.store.Bool(ptr + 4 + 0, x["originTypes"]["extension"] ? true : false);
          A.store.Bool(ptr + 4 + 4, "protectedWeb" in x["originTypes"] ? true : false);
          A.store.Bool(ptr + 4 + 1, x["originTypes"]["protectedWeb"] ? true : false);
          A.store.Bool(ptr + 4 + 5, "unprotectedWeb" in x["originTypes"] ? true : false);
          A.store.Bool(ptr + 4 + 2, x["originTypes"]["unprotectedWeb"] ? true : false);
        }
        A.store.Ref(ptr + 12, x["origins"]);
        A.store.Bool(ptr + 24, "since" in x ? true : false);
        A.store.Float64(ptr + 16, x["since"] === undefined ? 0 : (x["since"] as number));
      }
    },
    "load_RemovalOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["excludeOrigins"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 4 + 6)) {
        x["originTypes"] = {};
        if (A.load.Bool(ptr + 4 + 3)) {
          x["originTypes"]["extension"] = A.load.Bool(ptr + 4 + 0);
        } else {
          delete x["originTypes"]["extension"];
        }
        if (A.load.Bool(ptr + 4 + 4)) {
          x["originTypes"]["protectedWeb"] = A.load.Bool(ptr + 4 + 1);
        } else {
          delete x["originTypes"]["protectedWeb"];
        }
        if (A.load.Bool(ptr + 4 + 5)) {
          x["originTypes"]["unprotectedWeb"] = A.load.Bool(ptr + 4 + 2);
        } else {
          delete x["originTypes"]["unprotectedWeb"];
        }
      } else {
        delete x["originTypes"];
      }
      x["origins"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 24)) {
        x["since"] = A.load.Float64(ptr + 16);
      } else {
        delete x["since"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SettingsReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 90, false);

        A.store.Bool(ptr + 0 + 30, false);
        A.store.Bool(ptr + 0 + 15, false);
        A.store.Bool(ptr + 0 + 0, false);
        A.store.Bool(ptr + 0 + 16, false);
        A.store.Bool(ptr + 0 + 1, false);
        A.store.Bool(ptr + 0 + 17, false);
        A.store.Bool(ptr + 0 + 2, false);
        A.store.Bool(ptr + 0 + 18, false);
        A.store.Bool(ptr + 0 + 3, false);
        A.store.Bool(ptr + 0 + 19, false);
        A.store.Bool(ptr + 0 + 4, false);
        A.store.Bool(ptr + 0 + 20, false);
        A.store.Bool(ptr + 0 + 5, false);
        A.store.Bool(ptr + 0 + 21, false);
        A.store.Bool(ptr + 0 + 6, false);
        A.store.Bool(ptr + 0 + 22, false);
        A.store.Bool(ptr + 0 + 7, false);
        A.store.Bool(ptr + 0 + 23, false);
        A.store.Bool(ptr + 0 + 8, false);
        A.store.Bool(ptr + 0 + 24, false);
        A.store.Bool(ptr + 0 + 9, false);
        A.store.Bool(ptr + 0 + 25, false);
        A.store.Bool(ptr + 0 + 10, false);
        A.store.Bool(ptr + 0 + 26, false);
        A.store.Bool(ptr + 0 + 11, false);
        A.store.Bool(ptr + 0 + 27, false);
        A.store.Bool(ptr + 0 + 12, false);
        A.store.Bool(ptr + 0 + 28, false);
        A.store.Bool(ptr + 0 + 13, false);
        A.store.Bool(ptr + 0 + 29, false);
        A.store.Bool(ptr + 0 + 14, false);

        A.store.Bool(ptr + 31 + 30, false);
        A.store.Bool(ptr + 31 + 15, false);
        A.store.Bool(ptr + 31 + 0, false);
        A.store.Bool(ptr + 31 + 16, false);
        A.store.Bool(ptr + 31 + 1, false);
        A.store.Bool(ptr + 31 + 17, false);
        A.store.Bool(ptr + 31 + 2, false);
        A.store.Bool(ptr + 31 + 18, false);
        A.store.Bool(ptr + 31 + 3, false);
        A.store.Bool(ptr + 31 + 19, false);
        A.store.Bool(ptr + 31 + 4, false);
        A.store.Bool(ptr + 31 + 20, false);
        A.store.Bool(ptr + 31 + 5, false);
        A.store.Bool(ptr + 31 + 21, false);
        A.store.Bool(ptr + 31 + 6, false);
        A.store.Bool(ptr + 31 + 22, false);
        A.store.Bool(ptr + 31 + 7, false);
        A.store.Bool(ptr + 31 + 23, false);
        A.store.Bool(ptr + 31 + 8, false);
        A.store.Bool(ptr + 31 + 24, false);
        A.store.Bool(ptr + 31 + 9, false);
        A.store.Bool(ptr + 31 + 25, false);
        A.store.Bool(ptr + 31 + 10, false);
        A.store.Bool(ptr + 31 + 26, false);
        A.store.Bool(ptr + 31 + 11, false);
        A.store.Bool(ptr + 31 + 27, false);
        A.store.Bool(ptr + 31 + 12, false);
        A.store.Bool(ptr + 31 + 28, false);
        A.store.Bool(ptr + 31 + 13, false);
        A.store.Bool(ptr + 31 + 29, false);
        A.store.Bool(ptr + 31 + 14, false);

        A.store.Bool(ptr + 64 + 25, false);
        A.store.Ref(ptr + 64 + 0, undefined);

        A.store.Bool(ptr + 64 + 4 + 6, false);
        A.store.Bool(ptr + 64 + 4 + 3, false);
        A.store.Bool(ptr + 64 + 4 + 0, false);
        A.store.Bool(ptr + 64 + 4 + 4, false);
        A.store.Bool(ptr + 64 + 4 + 1, false);
        A.store.Bool(ptr + 64 + 4 + 5, false);
        A.store.Bool(ptr + 64 + 4 + 2, false);
        A.store.Ref(ptr + 64 + 12, undefined);
        A.store.Bool(ptr + 64 + 24, false);
        A.store.Float64(ptr + 64 + 16, 0);
      } else {
        A.store.Bool(ptr + 90, true);

        if (typeof x["dataRemovalPermitted"] === "undefined") {
          A.store.Bool(ptr + 0 + 30, false);
          A.store.Bool(ptr + 0 + 15, false);
          A.store.Bool(ptr + 0 + 0, false);
          A.store.Bool(ptr + 0 + 16, false);
          A.store.Bool(ptr + 0 + 1, false);
          A.store.Bool(ptr + 0 + 17, false);
          A.store.Bool(ptr + 0 + 2, false);
          A.store.Bool(ptr + 0 + 18, false);
          A.store.Bool(ptr + 0 + 3, false);
          A.store.Bool(ptr + 0 + 19, false);
          A.store.Bool(ptr + 0 + 4, false);
          A.store.Bool(ptr + 0 + 20, false);
          A.store.Bool(ptr + 0 + 5, false);
          A.store.Bool(ptr + 0 + 21, false);
          A.store.Bool(ptr + 0 + 6, false);
          A.store.Bool(ptr + 0 + 22, false);
          A.store.Bool(ptr + 0 + 7, false);
          A.store.Bool(ptr + 0 + 23, false);
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Bool(ptr + 0 + 24, false);
          A.store.Bool(ptr + 0 + 9, false);
          A.store.Bool(ptr + 0 + 25, false);
          A.store.Bool(ptr + 0 + 10, false);
          A.store.Bool(ptr + 0 + 26, false);
          A.store.Bool(ptr + 0 + 11, false);
          A.store.Bool(ptr + 0 + 27, false);
          A.store.Bool(ptr + 0 + 12, false);
          A.store.Bool(ptr + 0 + 28, false);
          A.store.Bool(ptr + 0 + 13, false);
          A.store.Bool(ptr + 0 + 29, false);
          A.store.Bool(ptr + 0 + 14, false);
        } else {
          A.store.Bool(ptr + 0 + 30, true);
          A.store.Bool(ptr + 0 + 15, "appcache" in x["dataRemovalPermitted"] ? true : false);
          A.store.Bool(ptr + 0 + 0, x["dataRemovalPermitted"]["appcache"] ? true : false);
          A.store.Bool(ptr + 0 + 16, "cache" in x["dataRemovalPermitted"] ? true : false);
          A.store.Bool(ptr + 0 + 1, x["dataRemovalPermitted"]["cache"] ? true : false);
          A.store.Bool(ptr + 0 + 17, "cacheStorage" in x["dataRemovalPermitted"] ? true : false);
          A.store.Bool(ptr + 0 + 2, x["dataRemovalPermitted"]["cacheStorage"] ? true : false);
          A.store.Bool(ptr + 0 + 18, "cookies" in x["dataRemovalPermitted"] ? true : false);
          A.store.Bool(ptr + 0 + 3, x["dataRemovalPermitted"]["cookies"] ? true : false);
          A.store.Bool(ptr + 0 + 19, "downloads" in x["dataRemovalPermitted"] ? true : false);
          A.store.Bool(ptr + 0 + 4, x["dataRemovalPermitted"]["downloads"] ? true : false);
          A.store.Bool(ptr + 0 + 20, "fileSystems" in x["dataRemovalPermitted"] ? true : false);
          A.store.Bool(ptr + 0 + 5, x["dataRemovalPermitted"]["fileSystems"] ? true : false);
          A.store.Bool(ptr + 0 + 21, "formData" in x["dataRemovalPermitted"] ? true : false);
          A.store.Bool(ptr + 0 + 6, x["dataRemovalPermitted"]["formData"] ? true : false);
          A.store.Bool(ptr + 0 + 22, "history" in x["dataRemovalPermitted"] ? true : false);
          A.store.Bool(ptr + 0 + 7, x["dataRemovalPermitted"]["history"] ? true : false);
          A.store.Bool(ptr + 0 + 23, "indexedDB" in x["dataRemovalPermitted"] ? true : false);
          A.store.Bool(ptr + 0 + 8, x["dataRemovalPermitted"]["indexedDB"] ? true : false);
          A.store.Bool(ptr + 0 + 24, "localStorage" in x["dataRemovalPermitted"] ? true : false);
          A.store.Bool(ptr + 0 + 9, x["dataRemovalPermitted"]["localStorage"] ? true : false);
          A.store.Bool(ptr + 0 + 25, "passwords" in x["dataRemovalPermitted"] ? true : false);
          A.store.Bool(ptr + 0 + 10, x["dataRemovalPermitted"]["passwords"] ? true : false);
          A.store.Bool(ptr + 0 + 26, "pluginData" in x["dataRemovalPermitted"] ? true : false);
          A.store.Bool(ptr + 0 + 11, x["dataRemovalPermitted"]["pluginData"] ? true : false);
          A.store.Bool(ptr + 0 + 27, "serverBoundCertificates" in x["dataRemovalPermitted"] ? true : false);
          A.store.Bool(ptr + 0 + 12, x["dataRemovalPermitted"]["serverBoundCertificates"] ? true : false);
          A.store.Bool(ptr + 0 + 28, "serviceWorkers" in x["dataRemovalPermitted"] ? true : false);
          A.store.Bool(ptr + 0 + 13, x["dataRemovalPermitted"]["serviceWorkers"] ? true : false);
          A.store.Bool(ptr + 0 + 29, "webSQL" in x["dataRemovalPermitted"] ? true : false);
          A.store.Bool(ptr + 0 + 14, x["dataRemovalPermitted"]["webSQL"] ? true : false);
        }

        if (typeof x["dataToRemove"] === "undefined") {
          A.store.Bool(ptr + 31 + 30, false);
          A.store.Bool(ptr + 31 + 15, false);
          A.store.Bool(ptr + 31 + 0, false);
          A.store.Bool(ptr + 31 + 16, false);
          A.store.Bool(ptr + 31 + 1, false);
          A.store.Bool(ptr + 31 + 17, false);
          A.store.Bool(ptr + 31 + 2, false);
          A.store.Bool(ptr + 31 + 18, false);
          A.store.Bool(ptr + 31 + 3, false);
          A.store.Bool(ptr + 31 + 19, false);
          A.store.Bool(ptr + 31 + 4, false);
          A.store.Bool(ptr + 31 + 20, false);
          A.store.Bool(ptr + 31 + 5, false);
          A.store.Bool(ptr + 31 + 21, false);
          A.store.Bool(ptr + 31 + 6, false);
          A.store.Bool(ptr + 31 + 22, false);
          A.store.Bool(ptr + 31 + 7, false);
          A.store.Bool(ptr + 31 + 23, false);
          A.store.Bool(ptr + 31 + 8, false);
          A.store.Bool(ptr + 31 + 24, false);
          A.store.Bool(ptr + 31 + 9, false);
          A.store.Bool(ptr + 31 + 25, false);
          A.store.Bool(ptr + 31 + 10, false);
          A.store.Bool(ptr + 31 + 26, false);
          A.store.Bool(ptr + 31 + 11, false);
          A.store.Bool(ptr + 31 + 27, false);
          A.store.Bool(ptr + 31 + 12, false);
          A.store.Bool(ptr + 31 + 28, false);
          A.store.Bool(ptr + 31 + 13, false);
          A.store.Bool(ptr + 31 + 29, false);
          A.store.Bool(ptr + 31 + 14, false);
        } else {
          A.store.Bool(ptr + 31 + 30, true);
          A.store.Bool(ptr + 31 + 15, "appcache" in x["dataToRemove"] ? true : false);
          A.store.Bool(ptr + 31 + 0, x["dataToRemove"]["appcache"] ? true : false);
          A.store.Bool(ptr + 31 + 16, "cache" in x["dataToRemove"] ? true : false);
          A.store.Bool(ptr + 31 + 1, x["dataToRemove"]["cache"] ? true : false);
          A.store.Bool(ptr + 31 + 17, "cacheStorage" in x["dataToRemove"] ? true : false);
          A.store.Bool(ptr + 31 + 2, x["dataToRemove"]["cacheStorage"] ? true : false);
          A.store.Bool(ptr + 31 + 18, "cookies" in x["dataToRemove"] ? true : false);
          A.store.Bool(ptr + 31 + 3, x["dataToRemove"]["cookies"] ? true : false);
          A.store.Bool(ptr + 31 + 19, "downloads" in x["dataToRemove"] ? true : false);
          A.store.Bool(ptr + 31 + 4, x["dataToRemove"]["downloads"] ? true : false);
          A.store.Bool(ptr + 31 + 20, "fileSystems" in x["dataToRemove"] ? true : false);
          A.store.Bool(ptr + 31 + 5, x["dataToRemove"]["fileSystems"] ? true : false);
          A.store.Bool(ptr + 31 + 21, "formData" in x["dataToRemove"] ? true : false);
          A.store.Bool(ptr + 31 + 6, x["dataToRemove"]["formData"] ? true : false);
          A.store.Bool(ptr + 31 + 22, "history" in x["dataToRemove"] ? true : false);
          A.store.Bool(ptr + 31 + 7, x["dataToRemove"]["history"] ? true : false);
          A.store.Bool(ptr + 31 + 23, "indexedDB" in x["dataToRemove"] ? true : false);
          A.store.Bool(ptr + 31 + 8, x["dataToRemove"]["indexedDB"] ? true : false);
          A.store.Bool(ptr + 31 + 24, "localStorage" in x["dataToRemove"] ? true : false);
          A.store.Bool(ptr + 31 + 9, x["dataToRemove"]["localStorage"] ? true : false);
          A.store.Bool(ptr + 31 + 25, "passwords" in x["dataToRemove"] ? true : false);
          A.store.Bool(ptr + 31 + 10, x["dataToRemove"]["passwords"] ? true : false);
          A.store.Bool(ptr + 31 + 26, "pluginData" in x["dataToRemove"] ? true : false);
          A.store.Bool(ptr + 31 + 11, x["dataToRemove"]["pluginData"] ? true : false);
          A.store.Bool(ptr + 31 + 27, "serverBoundCertificates" in x["dataToRemove"] ? true : false);
          A.store.Bool(ptr + 31 + 12, x["dataToRemove"]["serverBoundCertificates"] ? true : false);
          A.store.Bool(ptr + 31 + 28, "serviceWorkers" in x["dataToRemove"] ? true : false);
          A.store.Bool(ptr + 31 + 13, x["dataToRemove"]["serviceWorkers"] ? true : false);
          A.store.Bool(ptr + 31 + 29, "webSQL" in x["dataToRemove"] ? true : false);
          A.store.Bool(ptr + 31 + 14, x["dataToRemove"]["webSQL"] ? true : false);
        }

        if (typeof x["options"] === "undefined") {
          A.store.Bool(ptr + 64 + 25, false);
          A.store.Ref(ptr + 64 + 0, undefined);

          A.store.Bool(ptr + 64 + 4 + 6, false);
          A.store.Bool(ptr + 64 + 4 + 3, false);
          A.store.Bool(ptr + 64 + 4 + 0, false);
          A.store.Bool(ptr + 64 + 4 + 4, false);
          A.store.Bool(ptr + 64 + 4 + 1, false);
          A.store.Bool(ptr + 64 + 4 + 5, false);
          A.store.Bool(ptr + 64 + 4 + 2, false);
          A.store.Ref(ptr + 64 + 12, undefined);
          A.store.Bool(ptr + 64 + 24, false);
          A.store.Float64(ptr + 64 + 16, 0);
        } else {
          A.store.Bool(ptr + 64 + 25, true);
          A.store.Ref(ptr + 64 + 0, x["options"]["excludeOrigins"]);

          if (typeof x["options"]["originTypes"] === "undefined") {
            A.store.Bool(ptr + 64 + 4 + 6, false);
            A.store.Bool(ptr + 64 + 4 + 3, false);
            A.store.Bool(ptr + 64 + 4 + 0, false);
            A.store.Bool(ptr + 64 + 4 + 4, false);
            A.store.Bool(ptr + 64 + 4 + 1, false);
            A.store.Bool(ptr + 64 + 4 + 5, false);
            A.store.Bool(ptr + 64 + 4 + 2, false);
          } else {
            A.store.Bool(ptr + 64 + 4 + 6, true);
            A.store.Bool(ptr + 64 + 4 + 3, "extension" in x["options"]["originTypes"] ? true : false);
            A.store.Bool(ptr + 64 + 4 + 0, x["options"]["originTypes"]["extension"] ? true : false);
            A.store.Bool(ptr + 64 + 4 + 4, "protectedWeb" in x["options"]["originTypes"] ? true : false);
            A.store.Bool(ptr + 64 + 4 + 1, x["options"]["originTypes"]["protectedWeb"] ? true : false);
            A.store.Bool(ptr + 64 + 4 + 5, "unprotectedWeb" in x["options"]["originTypes"] ? true : false);
            A.store.Bool(ptr + 64 + 4 + 2, x["options"]["originTypes"]["unprotectedWeb"] ? true : false);
          }
          A.store.Ref(ptr + 64 + 12, x["options"]["origins"]);
          A.store.Bool(ptr + 64 + 24, "since" in x["options"] ? true : false);
          A.store.Float64(ptr + 64 + 16, x["options"]["since"] === undefined ? 0 : (x["options"]["since"] as number));
        }
      }
    },
    "load_SettingsReturnType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 30)) {
        x["dataRemovalPermitted"] = {};
        if (A.load.Bool(ptr + 0 + 15)) {
          x["dataRemovalPermitted"]["appcache"] = A.load.Bool(ptr + 0 + 0);
        } else {
          delete x["dataRemovalPermitted"]["appcache"];
        }
        if (A.load.Bool(ptr + 0 + 16)) {
          x["dataRemovalPermitted"]["cache"] = A.load.Bool(ptr + 0 + 1);
        } else {
          delete x["dataRemovalPermitted"]["cache"];
        }
        if (A.load.Bool(ptr + 0 + 17)) {
          x["dataRemovalPermitted"]["cacheStorage"] = A.load.Bool(ptr + 0 + 2);
        } else {
          delete x["dataRemovalPermitted"]["cacheStorage"];
        }
        if (A.load.Bool(ptr + 0 + 18)) {
          x["dataRemovalPermitted"]["cookies"] = A.load.Bool(ptr + 0 + 3);
        } else {
          delete x["dataRemovalPermitted"]["cookies"];
        }
        if (A.load.Bool(ptr + 0 + 19)) {
          x["dataRemovalPermitted"]["downloads"] = A.load.Bool(ptr + 0 + 4);
        } else {
          delete x["dataRemovalPermitted"]["downloads"];
        }
        if (A.load.Bool(ptr + 0 + 20)) {
          x["dataRemovalPermitted"]["fileSystems"] = A.load.Bool(ptr + 0 + 5);
        } else {
          delete x["dataRemovalPermitted"]["fileSystems"];
        }
        if (A.load.Bool(ptr + 0 + 21)) {
          x["dataRemovalPermitted"]["formData"] = A.load.Bool(ptr + 0 + 6);
        } else {
          delete x["dataRemovalPermitted"]["formData"];
        }
        if (A.load.Bool(ptr + 0 + 22)) {
          x["dataRemovalPermitted"]["history"] = A.load.Bool(ptr + 0 + 7);
        } else {
          delete x["dataRemovalPermitted"]["history"];
        }
        if (A.load.Bool(ptr + 0 + 23)) {
          x["dataRemovalPermitted"]["indexedDB"] = A.load.Bool(ptr + 0 + 8);
        } else {
          delete x["dataRemovalPermitted"]["indexedDB"];
        }
        if (A.load.Bool(ptr + 0 + 24)) {
          x["dataRemovalPermitted"]["localStorage"] = A.load.Bool(ptr + 0 + 9);
        } else {
          delete x["dataRemovalPermitted"]["localStorage"];
        }
        if (A.load.Bool(ptr + 0 + 25)) {
          x["dataRemovalPermitted"]["passwords"] = A.load.Bool(ptr + 0 + 10);
        } else {
          delete x["dataRemovalPermitted"]["passwords"];
        }
        if (A.load.Bool(ptr + 0 + 26)) {
          x["dataRemovalPermitted"]["pluginData"] = A.load.Bool(ptr + 0 + 11);
        } else {
          delete x["dataRemovalPermitted"]["pluginData"];
        }
        if (A.load.Bool(ptr + 0 + 27)) {
          x["dataRemovalPermitted"]["serverBoundCertificates"] = A.load.Bool(ptr + 0 + 12);
        } else {
          delete x["dataRemovalPermitted"]["serverBoundCertificates"];
        }
        if (A.load.Bool(ptr + 0 + 28)) {
          x["dataRemovalPermitted"]["serviceWorkers"] = A.load.Bool(ptr + 0 + 13);
        } else {
          delete x["dataRemovalPermitted"]["serviceWorkers"];
        }
        if (A.load.Bool(ptr + 0 + 29)) {
          x["dataRemovalPermitted"]["webSQL"] = A.load.Bool(ptr + 0 + 14);
        } else {
          delete x["dataRemovalPermitted"]["webSQL"];
        }
      } else {
        delete x["dataRemovalPermitted"];
      }
      if (A.load.Bool(ptr + 31 + 30)) {
        x["dataToRemove"] = {};
        if (A.load.Bool(ptr + 31 + 15)) {
          x["dataToRemove"]["appcache"] = A.load.Bool(ptr + 31 + 0);
        } else {
          delete x["dataToRemove"]["appcache"];
        }
        if (A.load.Bool(ptr + 31 + 16)) {
          x["dataToRemove"]["cache"] = A.load.Bool(ptr + 31 + 1);
        } else {
          delete x["dataToRemove"]["cache"];
        }
        if (A.load.Bool(ptr + 31 + 17)) {
          x["dataToRemove"]["cacheStorage"] = A.load.Bool(ptr + 31 + 2);
        } else {
          delete x["dataToRemove"]["cacheStorage"];
        }
        if (A.load.Bool(ptr + 31 + 18)) {
          x["dataToRemove"]["cookies"] = A.load.Bool(ptr + 31 + 3);
        } else {
          delete x["dataToRemove"]["cookies"];
        }
        if (A.load.Bool(ptr + 31 + 19)) {
          x["dataToRemove"]["downloads"] = A.load.Bool(ptr + 31 + 4);
        } else {
          delete x["dataToRemove"]["downloads"];
        }
        if (A.load.Bool(ptr + 31 + 20)) {
          x["dataToRemove"]["fileSystems"] = A.load.Bool(ptr + 31 + 5);
        } else {
          delete x["dataToRemove"]["fileSystems"];
        }
        if (A.load.Bool(ptr + 31 + 21)) {
          x["dataToRemove"]["formData"] = A.load.Bool(ptr + 31 + 6);
        } else {
          delete x["dataToRemove"]["formData"];
        }
        if (A.load.Bool(ptr + 31 + 22)) {
          x["dataToRemove"]["history"] = A.load.Bool(ptr + 31 + 7);
        } else {
          delete x["dataToRemove"]["history"];
        }
        if (A.load.Bool(ptr + 31 + 23)) {
          x["dataToRemove"]["indexedDB"] = A.load.Bool(ptr + 31 + 8);
        } else {
          delete x["dataToRemove"]["indexedDB"];
        }
        if (A.load.Bool(ptr + 31 + 24)) {
          x["dataToRemove"]["localStorage"] = A.load.Bool(ptr + 31 + 9);
        } else {
          delete x["dataToRemove"]["localStorage"];
        }
        if (A.load.Bool(ptr + 31 + 25)) {
          x["dataToRemove"]["passwords"] = A.load.Bool(ptr + 31 + 10);
        } else {
          delete x["dataToRemove"]["passwords"];
        }
        if (A.load.Bool(ptr + 31 + 26)) {
          x["dataToRemove"]["pluginData"] = A.load.Bool(ptr + 31 + 11);
        } else {
          delete x["dataToRemove"]["pluginData"];
        }
        if (A.load.Bool(ptr + 31 + 27)) {
          x["dataToRemove"]["serverBoundCertificates"] = A.load.Bool(ptr + 31 + 12);
        } else {
          delete x["dataToRemove"]["serverBoundCertificates"];
        }
        if (A.load.Bool(ptr + 31 + 28)) {
          x["dataToRemove"]["serviceWorkers"] = A.load.Bool(ptr + 31 + 13);
        } else {
          delete x["dataToRemove"]["serviceWorkers"];
        }
        if (A.load.Bool(ptr + 31 + 29)) {
          x["dataToRemove"]["webSQL"] = A.load.Bool(ptr + 31 + 14);
        } else {
          delete x["dataToRemove"]["webSQL"];
        }
      } else {
        delete x["dataToRemove"];
      }
      if (A.load.Bool(ptr + 64 + 25)) {
        x["options"] = {};
        x["options"]["excludeOrigins"] = A.load.Ref(ptr + 64 + 0, undefined);
        if (A.load.Bool(ptr + 64 + 4 + 6)) {
          x["options"]["originTypes"] = {};
          if (A.load.Bool(ptr + 64 + 4 + 3)) {
            x["options"]["originTypes"]["extension"] = A.load.Bool(ptr + 64 + 4 + 0);
          } else {
            delete x["options"]["originTypes"]["extension"];
          }
          if (A.load.Bool(ptr + 64 + 4 + 4)) {
            x["options"]["originTypes"]["protectedWeb"] = A.load.Bool(ptr + 64 + 4 + 1);
          } else {
            delete x["options"]["originTypes"]["protectedWeb"];
          }
          if (A.load.Bool(ptr + 64 + 4 + 5)) {
            x["options"]["originTypes"]["unprotectedWeb"] = A.load.Bool(ptr + 64 + 4 + 2);
          } else {
            delete x["options"]["originTypes"]["unprotectedWeb"];
          }
        } else {
          delete x["options"]["originTypes"];
        }
        x["options"]["origins"] = A.load.Ref(ptr + 64 + 12, undefined);
        if (A.load.Bool(ptr + 64 + 24)) {
          x["options"]["since"] = A.load.Float64(ptr + 64 + 16);
        } else {
          delete x["options"]["since"];
        }
      } else {
        delete x["options"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Remove": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "remove" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Remove": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.remove);
    },
    "call_Remove": (retPtr: Pointer, options: Pointer, dataToRemove: Pointer): void => {
      const options_ffi = {};

      options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 6)) {
        options_ffi["originTypes"] = {};
        if (A.load.Bool(options + 4 + 3)) {
          options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
        }
        if (A.load.Bool(options + 4 + 5)) {
          options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
        }
      }
      options_ffi["origins"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["since"] = A.load.Float64(options + 16);
      }

      const dataToRemove_ffi = {};

      if (A.load.Bool(dataToRemove + 15)) {
        dataToRemove_ffi["appcache"] = A.load.Bool(dataToRemove + 0);
      }
      if (A.load.Bool(dataToRemove + 16)) {
        dataToRemove_ffi["cache"] = A.load.Bool(dataToRemove + 1);
      }
      if (A.load.Bool(dataToRemove + 17)) {
        dataToRemove_ffi["cacheStorage"] = A.load.Bool(dataToRemove + 2);
      }
      if (A.load.Bool(dataToRemove + 18)) {
        dataToRemove_ffi["cookies"] = A.load.Bool(dataToRemove + 3);
      }
      if (A.load.Bool(dataToRemove + 19)) {
        dataToRemove_ffi["downloads"] = A.load.Bool(dataToRemove + 4);
      }
      if (A.load.Bool(dataToRemove + 20)) {
        dataToRemove_ffi["fileSystems"] = A.load.Bool(dataToRemove + 5);
      }
      if (A.load.Bool(dataToRemove + 21)) {
        dataToRemove_ffi["formData"] = A.load.Bool(dataToRemove + 6);
      }
      if (A.load.Bool(dataToRemove + 22)) {
        dataToRemove_ffi["history"] = A.load.Bool(dataToRemove + 7);
      }
      if (A.load.Bool(dataToRemove + 23)) {
        dataToRemove_ffi["indexedDB"] = A.load.Bool(dataToRemove + 8);
      }
      if (A.load.Bool(dataToRemove + 24)) {
        dataToRemove_ffi["localStorage"] = A.load.Bool(dataToRemove + 9);
      }
      if (A.load.Bool(dataToRemove + 25)) {
        dataToRemove_ffi["passwords"] = A.load.Bool(dataToRemove + 10);
      }
      if (A.load.Bool(dataToRemove + 26)) {
        dataToRemove_ffi["pluginData"] = A.load.Bool(dataToRemove + 11);
      }
      if (A.load.Bool(dataToRemove + 27)) {
        dataToRemove_ffi["serverBoundCertificates"] = A.load.Bool(dataToRemove + 12);
      }
      if (A.load.Bool(dataToRemove + 28)) {
        dataToRemove_ffi["serviceWorkers"] = A.load.Bool(dataToRemove + 13);
      }
      if (A.load.Bool(dataToRemove + 29)) {
        dataToRemove_ffi["webSQL"] = A.load.Bool(dataToRemove + 14);
      }

      const _ret = WEBEXT.browsingData.remove(options_ffi, dataToRemove_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Remove": (retPtr: Pointer, errPtr: Pointer, options: Pointer, dataToRemove: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 6)) {
          options_ffi["originTypes"] = {};
          if (A.load.Bool(options + 4 + 3)) {
            options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 4)) {
            options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
          }
          if (A.load.Bool(options + 4 + 5)) {
            options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
          }
        }
        options_ffi["origins"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["since"] = A.load.Float64(options + 16);
        }

        const dataToRemove_ffi = {};

        if (A.load.Bool(dataToRemove + 15)) {
          dataToRemove_ffi["appcache"] = A.load.Bool(dataToRemove + 0);
        }
        if (A.load.Bool(dataToRemove + 16)) {
          dataToRemove_ffi["cache"] = A.load.Bool(dataToRemove + 1);
        }
        if (A.load.Bool(dataToRemove + 17)) {
          dataToRemove_ffi["cacheStorage"] = A.load.Bool(dataToRemove + 2);
        }
        if (A.load.Bool(dataToRemove + 18)) {
          dataToRemove_ffi["cookies"] = A.load.Bool(dataToRemove + 3);
        }
        if (A.load.Bool(dataToRemove + 19)) {
          dataToRemove_ffi["downloads"] = A.load.Bool(dataToRemove + 4);
        }
        if (A.load.Bool(dataToRemove + 20)) {
          dataToRemove_ffi["fileSystems"] = A.load.Bool(dataToRemove + 5);
        }
        if (A.load.Bool(dataToRemove + 21)) {
          dataToRemove_ffi["formData"] = A.load.Bool(dataToRemove + 6);
        }
        if (A.load.Bool(dataToRemove + 22)) {
          dataToRemove_ffi["history"] = A.load.Bool(dataToRemove + 7);
        }
        if (A.load.Bool(dataToRemove + 23)) {
          dataToRemove_ffi["indexedDB"] = A.load.Bool(dataToRemove + 8);
        }
        if (A.load.Bool(dataToRemove + 24)) {
          dataToRemove_ffi["localStorage"] = A.load.Bool(dataToRemove + 9);
        }
        if (A.load.Bool(dataToRemove + 25)) {
          dataToRemove_ffi["passwords"] = A.load.Bool(dataToRemove + 10);
        }
        if (A.load.Bool(dataToRemove + 26)) {
          dataToRemove_ffi["pluginData"] = A.load.Bool(dataToRemove + 11);
        }
        if (A.load.Bool(dataToRemove + 27)) {
          dataToRemove_ffi["serverBoundCertificates"] = A.load.Bool(dataToRemove + 12);
        }
        if (A.load.Bool(dataToRemove + 28)) {
          dataToRemove_ffi["serviceWorkers"] = A.load.Bool(dataToRemove + 13);
        }
        if (A.load.Bool(dataToRemove + 29)) {
          dataToRemove_ffi["webSQL"] = A.load.Bool(dataToRemove + 14);
        }

        const _ret = WEBEXT.browsingData.remove(options_ffi, dataToRemove_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveAppcache": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "removeAppcache" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveAppcache": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.removeAppcache);
    },
    "call_RemoveAppcache": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 6)) {
        options_ffi["originTypes"] = {};
        if (A.load.Bool(options + 4 + 3)) {
          options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
        }
        if (A.load.Bool(options + 4 + 5)) {
          options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
        }
      }
      options_ffi["origins"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["since"] = A.load.Float64(options + 16);
      }

      const _ret = WEBEXT.browsingData.removeAppcache(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveAppcache": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 6)) {
          options_ffi["originTypes"] = {};
          if (A.load.Bool(options + 4 + 3)) {
            options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 4)) {
            options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
          }
          if (A.load.Bool(options + 4 + 5)) {
            options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
          }
        }
        options_ffi["origins"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["since"] = A.load.Float64(options + 16);
        }

        const _ret = WEBEXT.browsingData.removeAppcache(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveCache": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "removeCache" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveCache": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.removeCache);
    },
    "call_RemoveCache": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 6)) {
        options_ffi["originTypes"] = {};
        if (A.load.Bool(options + 4 + 3)) {
          options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
        }
        if (A.load.Bool(options + 4 + 5)) {
          options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
        }
      }
      options_ffi["origins"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["since"] = A.load.Float64(options + 16);
      }

      const _ret = WEBEXT.browsingData.removeCache(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveCache": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 6)) {
          options_ffi["originTypes"] = {};
          if (A.load.Bool(options + 4 + 3)) {
            options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 4)) {
            options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
          }
          if (A.load.Bool(options + 4 + 5)) {
            options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
          }
        }
        options_ffi["origins"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["since"] = A.load.Float64(options + 16);
        }

        const _ret = WEBEXT.browsingData.removeCache(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveCacheStorage": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "removeCacheStorage" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveCacheStorage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.removeCacheStorage);
    },
    "call_RemoveCacheStorage": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 6)) {
        options_ffi["originTypes"] = {};
        if (A.load.Bool(options + 4 + 3)) {
          options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
        }
        if (A.load.Bool(options + 4 + 5)) {
          options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
        }
      }
      options_ffi["origins"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["since"] = A.load.Float64(options + 16);
      }

      const _ret = WEBEXT.browsingData.removeCacheStorage(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveCacheStorage": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 6)) {
          options_ffi["originTypes"] = {};
          if (A.load.Bool(options + 4 + 3)) {
            options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 4)) {
            options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
          }
          if (A.load.Bool(options + 4 + 5)) {
            options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
          }
        }
        options_ffi["origins"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["since"] = A.load.Float64(options + 16);
        }

        const _ret = WEBEXT.browsingData.removeCacheStorage(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveCookies": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "removeCookies" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveCookies": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.removeCookies);
    },
    "call_RemoveCookies": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 6)) {
        options_ffi["originTypes"] = {};
        if (A.load.Bool(options + 4 + 3)) {
          options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
        }
        if (A.load.Bool(options + 4 + 5)) {
          options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
        }
      }
      options_ffi["origins"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["since"] = A.load.Float64(options + 16);
      }

      const _ret = WEBEXT.browsingData.removeCookies(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveCookies": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 6)) {
          options_ffi["originTypes"] = {};
          if (A.load.Bool(options + 4 + 3)) {
            options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 4)) {
            options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
          }
          if (A.load.Bool(options + 4 + 5)) {
            options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
          }
        }
        options_ffi["origins"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["since"] = A.load.Float64(options + 16);
        }

        const _ret = WEBEXT.browsingData.removeCookies(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveDownloads": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "removeDownloads" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveDownloads": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.removeDownloads);
    },
    "call_RemoveDownloads": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 6)) {
        options_ffi["originTypes"] = {};
        if (A.load.Bool(options + 4 + 3)) {
          options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
        }
        if (A.load.Bool(options + 4 + 5)) {
          options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
        }
      }
      options_ffi["origins"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["since"] = A.load.Float64(options + 16);
      }

      const _ret = WEBEXT.browsingData.removeDownloads(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveDownloads": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 6)) {
          options_ffi["originTypes"] = {};
          if (A.load.Bool(options + 4 + 3)) {
            options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 4)) {
            options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
          }
          if (A.load.Bool(options + 4 + 5)) {
            options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
          }
        }
        options_ffi["origins"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["since"] = A.load.Float64(options + 16);
        }

        const _ret = WEBEXT.browsingData.removeDownloads(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveFileSystems": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "removeFileSystems" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveFileSystems": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.removeFileSystems);
    },
    "call_RemoveFileSystems": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 6)) {
        options_ffi["originTypes"] = {};
        if (A.load.Bool(options + 4 + 3)) {
          options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
        }
        if (A.load.Bool(options + 4 + 5)) {
          options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
        }
      }
      options_ffi["origins"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["since"] = A.load.Float64(options + 16);
      }

      const _ret = WEBEXT.browsingData.removeFileSystems(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveFileSystems": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 6)) {
          options_ffi["originTypes"] = {};
          if (A.load.Bool(options + 4 + 3)) {
            options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 4)) {
            options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
          }
          if (A.load.Bool(options + 4 + 5)) {
            options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
          }
        }
        options_ffi["origins"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["since"] = A.load.Float64(options + 16);
        }

        const _ret = WEBEXT.browsingData.removeFileSystems(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveFormData": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "removeFormData" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveFormData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.removeFormData);
    },
    "call_RemoveFormData": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 6)) {
        options_ffi["originTypes"] = {};
        if (A.load.Bool(options + 4 + 3)) {
          options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
        }
        if (A.load.Bool(options + 4 + 5)) {
          options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
        }
      }
      options_ffi["origins"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["since"] = A.load.Float64(options + 16);
      }

      const _ret = WEBEXT.browsingData.removeFormData(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveFormData": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 6)) {
          options_ffi["originTypes"] = {};
          if (A.load.Bool(options + 4 + 3)) {
            options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 4)) {
            options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
          }
          if (A.load.Bool(options + 4 + 5)) {
            options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
          }
        }
        options_ffi["origins"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["since"] = A.load.Float64(options + 16);
        }

        const _ret = WEBEXT.browsingData.removeFormData(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveHistory": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "removeHistory" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveHistory": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.removeHistory);
    },
    "call_RemoveHistory": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 6)) {
        options_ffi["originTypes"] = {};
        if (A.load.Bool(options + 4 + 3)) {
          options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
        }
        if (A.load.Bool(options + 4 + 5)) {
          options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
        }
      }
      options_ffi["origins"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["since"] = A.load.Float64(options + 16);
      }

      const _ret = WEBEXT.browsingData.removeHistory(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveHistory": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 6)) {
          options_ffi["originTypes"] = {};
          if (A.load.Bool(options + 4 + 3)) {
            options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 4)) {
            options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
          }
          if (A.load.Bool(options + 4 + 5)) {
            options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
          }
        }
        options_ffi["origins"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["since"] = A.load.Float64(options + 16);
        }

        const _ret = WEBEXT.browsingData.removeHistory(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveIndexedDB": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "removeIndexedDB" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveIndexedDB": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.removeIndexedDB);
    },
    "call_RemoveIndexedDB": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 6)) {
        options_ffi["originTypes"] = {};
        if (A.load.Bool(options + 4 + 3)) {
          options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
        }
        if (A.load.Bool(options + 4 + 5)) {
          options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
        }
      }
      options_ffi["origins"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["since"] = A.load.Float64(options + 16);
      }

      const _ret = WEBEXT.browsingData.removeIndexedDB(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveIndexedDB": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 6)) {
          options_ffi["originTypes"] = {};
          if (A.load.Bool(options + 4 + 3)) {
            options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 4)) {
            options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
          }
          if (A.load.Bool(options + 4 + 5)) {
            options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
          }
        }
        options_ffi["origins"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["since"] = A.load.Float64(options + 16);
        }

        const _ret = WEBEXT.browsingData.removeIndexedDB(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveLocalStorage": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "removeLocalStorage" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveLocalStorage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.removeLocalStorage);
    },
    "call_RemoveLocalStorage": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 6)) {
        options_ffi["originTypes"] = {};
        if (A.load.Bool(options + 4 + 3)) {
          options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
        }
        if (A.load.Bool(options + 4 + 5)) {
          options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
        }
      }
      options_ffi["origins"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["since"] = A.load.Float64(options + 16);
      }

      const _ret = WEBEXT.browsingData.removeLocalStorage(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveLocalStorage": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 6)) {
          options_ffi["originTypes"] = {};
          if (A.load.Bool(options + 4 + 3)) {
            options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 4)) {
            options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
          }
          if (A.load.Bool(options + 4 + 5)) {
            options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
          }
        }
        options_ffi["origins"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["since"] = A.load.Float64(options + 16);
        }

        const _ret = WEBEXT.browsingData.removeLocalStorage(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemovePasswords": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "removePasswords" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemovePasswords": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.removePasswords);
    },
    "call_RemovePasswords": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 6)) {
        options_ffi["originTypes"] = {};
        if (A.load.Bool(options + 4 + 3)) {
          options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
        }
        if (A.load.Bool(options + 4 + 5)) {
          options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
        }
      }
      options_ffi["origins"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["since"] = A.load.Float64(options + 16);
      }

      const _ret = WEBEXT.browsingData.removePasswords(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemovePasswords": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 6)) {
          options_ffi["originTypes"] = {};
          if (A.load.Bool(options + 4 + 3)) {
            options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 4)) {
            options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
          }
          if (A.load.Bool(options + 4 + 5)) {
            options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
          }
        }
        options_ffi["origins"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["since"] = A.load.Float64(options + 16);
        }

        const _ret = WEBEXT.browsingData.removePasswords(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemovePluginData": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "removePluginData" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemovePluginData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.removePluginData);
    },
    "call_RemovePluginData": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 6)) {
        options_ffi["originTypes"] = {};
        if (A.load.Bool(options + 4 + 3)) {
          options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
        }
        if (A.load.Bool(options + 4 + 5)) {
          options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
        }
      }
      options_ffi["origins"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["since"] = A.load.Float64(options + 16);
      }

      const _ret = WEBEXT.browsingData.removePluginData(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemovePluginData": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 6)) {
          options_ffi["originTypes"] = {};
          if (A.load.Bool(options + 4 + 3)) {
            options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 4)) {
            options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
          }
          if (A.load.Bool(options + 4 + 5)) {
            options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
          }
        }
        options_ffi["origins"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["since"] = A.load.Float64(options + 16);
        }

        const _ret = WEBEXT.browsingData.removePluginData(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveServiceWorkers": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "removeServiceWorkers" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveServiceWorkers": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.removeServiceWorkers);
    },
    "call_RemoveServiceWorkers": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 6)) {
        options_ffi["originTypes"] = {};
        if (A.load.Bool(options + 4 + 3)) {
          options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
        }
        if (A.load.Bool(options + 4 + 5)) {
          options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
        }
      }
      options_ffi["origins"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["since"] = A.load.Float64(options + 16);
      }

      const _ret = WEBEXT.browsingData.removeServiceWorkers(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveServiceWorkers": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 6)) {
          options_ffi["originTypes"] = {};
          if (A.load.Bool(options + 4 + 3)) {
            options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 4)) {
            options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
          }
          if (A.load.Bool(options + 4 + 5)) {
            options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
          }
        }
        options_ffi["origins"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["since"] = A.load.Float64(options + 16);
        }

        const _ret = WEBEXT.browsingData.removeServiceWorkers(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveWebSQL": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "removeWebSQL" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveWebSQL": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.removeWebSQL);
    },
    "call_RemoveWebSQL": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 4 + 6)) {
        options_ffi["originTypes"] = {};
        if (A.load.Bool(options + 4 + 3)) {
          options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
        }
        if (A.load.Bool(options + 4 + 4)) {
          options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
        }
        if (A.load.Bool(options + 4 + 5)) {
          options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
        }
      }
      options_ffi["origins"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 24)) {
        options_ffi["since"] = A.load.Float64(options + 16);
      }

      const _ret = WEBEXT.browsingData.removeWebSQL(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveWebSQL": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["excludeOrigins"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 4 + 6)) {
          options_ffi["originTypes"] = {};
          if (A.load.Bool(options + 4 + 3)) {
            options_ffi["originTypes"]["extension"] = A.load.Bool(options + 4 + 0);
          }
          if (A.load.Bool(options + 4 + 4)) {
            options_ffi["originTypes"]["protectedWeb"] = A.load.Bool(options + 4 + 1);
          }
          if (A.load.Bool(options + 4 + 5)) {
            options_ffi["originTypes"]["unprotectedWeb"] = A.load.Bool(options + 4 + 2);
          }
        }
        options_ffi["origins"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 24)) {
          options_ffi["since"] = A.load.Float64(options + 16);
        }

        const _ret = WEBEXT.browsingData.removeWebSQL(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Settings": (): heap.Ref<boolean> => {
      if (WEBEXT?.browsingData && "settings" in WEBEXT?.browsingData) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Settings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.browsingData.settings);
    },
    "call_Settings": (retPtr: Pointer): void => {
      const _ret = WEBEXT.browsingData.settings();
      A.store.Ref(retPtr, _ret);
    },
    "try_Settings": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.browsingData.settings();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
