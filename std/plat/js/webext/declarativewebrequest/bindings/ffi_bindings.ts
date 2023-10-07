import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/declarativewebrequest", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_RequestCookie": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["value"]);
      }
    },
    "load_RequestCookie": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["value"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_AddRequestCookieInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.AddRequestCookie"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_AddRequestCookie": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);

        A.store.Bool(ptr + 0 + 8, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Enum(ptr + 12, -1);
      } else {
        A.store.Bool(ptr + 16, true);

        if (typeof x["cookie"] === "undefined") {
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
        } else {
          A.store.Bool(ptr + 0 + 8, true);
          A.store.Ref(ptr + 0 + 0, x["cookie"]["name"]);
          A.store.Ref(ptr + 0 + 4, x["cookie"]["value"]);
        }
        A.store.Enum(ptr + 12, ["declarativeWebRequest.AddRequestCookie"].indexOf(x["instanceType"] as string));
      }
    },
    "load_AddRequestCookie": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 8)) {
        x["cookie"] = {};
        x["cookie"]["name"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["cookie"]["value"] = A.load.Ref(ptr + 0 + 4, undefined);
      } else {
        delete x["cookie"];
      }
      x["instanceType"] = A.load.Enum(ptr + 12, ["declarativeWebRequest.AddRequestCookie"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ResponseCookie": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 41, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 40, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
      } else {
        A.store.Bool(ptr + 41, true);
        A.store.Ref(ptr + 0, x["domain"]);
        A.store.Ref(ptr + 4, x["expires"]);
        A.store.Ref(ptr + 8, x["httpOnly"]);
        A.store.Bool(ptr + 40, "maxAge" in x ? true : false);
        A.store.Float64(ptr + 16, x["maxAge"] === undefined ? 0 : (x["maxAge"] as number));
        A.store.Ref(ptr + 24, x["name"]);
        A.store.Ref(ptr + 28, x["path"]);
        A.store.Ref(ptr + 32, x["secure"]);
        A.store.Ref(ptr + 36, x["value"]);
      }
    },
    "load_ResponseCookie": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["domain"] = A.load.Ref(ptr + 0, undefined);
      x["expires"] = A.load.Ref(ptr + 4, undefined);
      x["httpOnly"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 40)) {
        x["maxAge"] = A.load.Float64(ptr + 16);
      } else {
        delete x["maxAge"];
      }
      x["name"] = A.load.Ref(ptr + 24, undefined);
      x["path"] = A.load.Ref(ptr + 28, undefined);
      x["secure"] = A.load.Ref(ptr + 32, undefined);
      x["value"] = A.load.Ref(ptr + 36, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_AddResponseCookieInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.AddResponseCookie"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_AddResponseCookie": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 48, false);

        A.store.Bool(ptr + 0 + 41, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Ref(ptr + 0 + 8, undefined);
        A.store.Bool(ptr + 0 + 40, false);
        A.store.Float64(ptr + 0 + 16, 0);
        A.store.Ref(ptr + 0 + 24, undefined);
        A.store.Ref(ptr + 0 + 28, undefined);
        A.store.Ref(ptr + 0 + 32, undefined);
        A.store.Ref(ptr + 0 + 36, undefined);
        A.store.Enum(ptr + 44, -1);
      } else {
        A.store.Bool(ptr + 48, true);

        if (typeof x["cookie"] === "undefined") {
          A.store.Bool(ptr + 0 + 41, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Ref(ptr + 0 + 8, undefined);
          A.store.Bool(ptr + 0 + 40, false);
          A.store.Float64(ptr + 0 + 16, 0);
          A.store.Ref(ptr + 0 + 24, undefined);
          A.store.Ref(ptr + 0 + 28, undefined);
          A.store.Ref(ptr + 0 + 32, undefined);
          A.store.Ref(ptr + 0 + 36, undefined);
        } else {
          A.store.Bool(ptr + 0 + 41, true);
          A.store.Ref(ptr + 0 + 0, x["cookie"]["domain"]);
          A.store.Ref(ptr + 0 + 4, x["cookie"]["expires"]);
          A.store.Ref(ptr + 0 + 8, x["cookie"]["httpOnly"]);
          A.store.Bool(ptr + 0 + 40, "maxAge" in x["cookie"] ? true : false);
          A.store.Float64(ptr + 0 + 16, x["cookie"]["maxAge"] === undefined ? 0 : (x["cookie"]["maxAge"] as number));
          A.store.Ref(ptr + 0 + 24, x["cookie"]["name"]);
          A.store.Ref(ptr + 0 + 28, x["cookie"]["path"]);
          A.store.Ref(ptr + 0 + 32, x["cookie"]["secure"]);
          A.store.Ref(ptr + 0 + 36, x["cookie"]["value"]);
        }
        A.store.Enum(ptr + 44, ["declarativeWebRequest.AddResponseCookie"].indexOf(x["instanceType"] as string));
      }
    },
    "load_AddResponseCookie": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 41)) {
        x["cookie"] = {};
        x["cookie"]["domain"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["cookie"]["expires"] = A.load.Ref(ptr + 0 + 4, undefined);
        x["cookie"]["httpOnly"] = A.load.Ref(ptr + 0 + 8, undefined);
        if (A.load.Bool(ptr + 0 + 40)) {
          x["cookie"]["maxAge"] = A.load.Float64(ptr + 0 + 16);
        } else {
          delete x["cookie"]["maxAge"];
        }
        x["cookie"]["name"] = A.load.Ref(ptr + 0 + 24, undefined);
        x["cookie"]["path"] = A.load.Ref(ptr + 0 + 28, undefined);
        x["cookie"]["secure"] = A.load.Ref(ptr + 0 + 32, undefined);
        x["cookie"]["value"] = A.load.Ref(ptr + 0 + 36, undefined);
      } else {
        delete x["cookie"];
      }
      x["instanceType"] = A.load.Enum(ptr + 44, ["declarativeWebRequest.AddResponseCookie"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_AddResponseHeaderInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.AddResponseHeader"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_AddResponseHeader": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Enum(ptr + 0, ["declarativeWebRequest.AddResponseHeader"].indexOf(x["instanceType"] as string));
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Ref(ptr + 8, x["value"]);
      }
    },
    "load_AddResponseHeader": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["instanceType"] = A.load.Enum(ptr + 0, ["declarativeWebRequest.AddResponseHeader"]);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      x["value"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_CancelRequestInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.CancelRequest"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_CancelRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["declarativeWebRequest.CancelRequest"].indexOf(x["instanceType"] as string));
      }
    },
    "load_CancelRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["instanceType"] = A.load.Enum(ptr + 0, ["declarativeWebRequest.CancelRequest"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_EditRequestCookieInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.EditRequestCookie"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_EditRequestCookie": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);

        A.store.Bool(ptr + 0 + 8, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Enum(ptr + 12, -1);

        A.store.Bool(ptr + 16 + 8, false);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 4, undefined);
      } else {
        A.store.Bool(ptr + 25, true);

        if (typeof x["filter"] === "undefined") {
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
        } else {
          A.store.Bool(ptr + 0 + 8, true);
          A.store.Ref(ptr + 0 + 0, x["filter"]["name"]);
          A.store.Ref(ptr + 0 + 4, x["filter"]["value"]);
        }
        A.store.Enum(ptr + 12, ["declarativeWebRequest.EditRequestCookie"].indexOf(x["instanceType"] as string));

        if (typeof x["modification"] === "undefined") {
          A.store.Bool(ptr + 16 + 8, false);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 4, undefined);
        } else {
          A.store.Bool(ptr + 16 + 8, true);
          A.store.Ref(ptr + 16 + 0, x["modification"]["name"]);
          A.store.Ref(ptr + 16 + 4, x["modification"]["value"]);
        }
      }
    },
    "load_EditRequestCookie": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 8)) {
        x["filter"] = {};
        x["filter"]["name"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["filter"]["value"] = A.load.Ref(ptr + 0 + 4, undefined);
      } else {
        delete x["filter"];
      }
      x["instanceType"] = A.load.Enum(ptr + 12, ["declarativeWebRequest.EditRequestCookie"]);
      if (A.load.Bool(ptr + 16 + 8)) {
        x["modification"] = {};
        x["modification"]["name"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["modification"]["value"] = A.load.Ref(ptr + 16 + 4, undefined);
      } else {
        delete x["modification"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_FilterResponseCookie": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 64, false);
        A.store.Bool(ptr + 60, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Bool(ptr + 61, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Bool(ptr + 62, false);
        A.store.Float64(ptr + 32, 0);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Ref(ptr + 48, undefined);
        A.store.Bool(ptr + 63, false);
        A.store.Bool(ptr + 52, false);
        A.store.Ref(ptr + 56, undefined);
      } else {
        A.store.Bool(ptr + 64, true);
        A.store.Bool(ptr + 60, "ageLowerBound" in x ? true : false);
        A.store.Int64(ptr + 0, x["ageLowerBound"] === undefined ? 0 : (x["ageLowerBound"] as number));
        A.store.Bool(ptr + 61, "ageUpperBound" in x ? true : false);
        A.store.Int64(ptr + 8, x["ageUpperBound"] === undefined ? 0 : (x["ageUpperBound"] as number));
        A.store.Ref(ptr + 16, x["domain"]);
        A.store.Ref(ptr + 20, x["expires"]);
        A.store.Ref(ptr + 24, x["httpOnly"]);
        A.store.Bool(ptr + 62, "maxAge" in x ? true : false);
        A.store.Float64(ptr + 32, x["maxAge"] === undefined ? 0 : (x["maxAge"] as number));
        A.store.Ref(ptr + 40, x["name"]);
        A.store.Ref(ptr + 44, x["path"]);
        A.store.Ref(ptr + 48, x["secure"]);
        A.store.Bool(ptr + 63, "sessionCookie" in x ? true : false);
        A.store.Bool(ptr + 52, x["sessionCookie"] ? true : false);
        A.store.Ref(ptr + 56, x["value"]);
      }
    },
    "load_FilterResponseCookie": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 60)) {
        x["ageLowerBound"] = A.load.Int64(ptr + 0);
      } else {
        delete x["ageLowerBound"];
      }
      if (A.load.Bool(ptr + 61)) {
        x["ageUpperBound"] = A.load.Int64(ptr + 8);
      } else {
        delete x["ageUpperBound"];
      }
      x["domain"] = A.load.Ref(ptr + 16, undefined);
      x["expires"] = A.load.Ref(ptr + 20, undefined);
      x["httpOnly"] = A.load.Ref(ptr + 24, undefined);
      if (A.load.Bool(ptr + 62)) {
        x["maxAge"] = A.load.Float64(ptr + 32);
      } else {
        delete x["maxAge"];
      }
      x["name"] = A.load.Ref(ptr + 40, undefined);
      x["path"] = A.load.Ref(ptr + 44, undefined);
      x["secure"] = A.load.Ref(ptr + 48, undefined);
      if (A.load.Bool(ptr + 63)) {
        x["sessionCookie"] = A.load.Bool(ptr + 52);
      } else {
        delete x["sessionCookie"];
      }
      x["value"] = A.load.Ref(ptr + 56, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_EditResponseCookieInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.EditResponseCookie"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_EditResponseCookie": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 114, false);

        A.store.Bool(ptr + 0 + 64, false);
        A.store.Bool(ptr + 0 + 60, false);
        A.store.Int64(ptr + 0 + 0, 0);
        A.store.Bool(ptr + 0 + 61, false);
        A.store.Int64(ptr + 0 + 8, 0);
        A.store.Ref(ptr + 0 + 16, undefined);
        A.store.Ref(ptr + 0 + 20, undefined);
        A.store.Ref(ptr + 0 + 24, undefined);
        A.store.Bool(ptr + 0 + 62, false);
        A.store.Float64(ptr + 0 + 32, 0);
        A.store.Ref(ptr + 0 + 40, undefined);
        A.store.Ref(ptr + 0 + 44, undefined);
        A.store.Ref(ptr + 0 + 48, undefined);
        A.store.Bool(ptr + 0 + 63, false);
        A.store.Bool(ptr + 0 + 52, false);
        A.store.Ref(ptr + 0 + 56, undefined);
        A.store.Enum(ptr + 68, -1);

        A.store.Bool(ptr + 72 + 41, false);
        A.store.Ref(ptr + 72 + 0, undefined);
        A.store.Ref(ptr + 72 + 4, undefined);
        A.store.Ref(ptr + 72 + 8, undefined);
        A.store.Bool(ptr + 72 + 40, false);
        A.store.Float64(ptr + 72 + 16, 0);
        A.store.Ref(ptr + 72 + 24, undefined);
        A.store.Ref(ptr + 72 + 28, undefined);
        A.store.Ref(ptr + 72 + 32, undefined);
        A.store.Ref(ptr + 72 + 36, undefined);
      } else {
        A.store.Bool(ptr + 114, true);

        if (typeof x["filter"] === "undefined") {
          A.store.Bool(ptr + 0 + 64, false);
          A.store.Bool(ptr + 0 + 60, false);
          A.store.Int64(ptr + 0 + 0, 0);
          A.store.Bool(ptr + 0 + 61, false);
          A.store.Int64(ptr + 0 + 8, 0);
          A.store.Ref(ptr + 0 + 16, undefined);
          A.store.Ref(ptr + 0 + 20, undefined);
          A.store.Ref(ptr + 0 + 24, undefined);
          A.store.Bool(ptr + 0 + 62, false);
          A.store.Float64(ptr + 0 + 32, 0);
          A.store.Ref(ptr + 0 + 40, undefined);
          A.store.Ref(ptr + 0 + 44, undefined);
          A.store.Ref(ptr + 0 + 48, undefined);
          A.store.Bool(ptr + 0 + 63, false);
          A.store.Bool(ptr + 0 + 52, false);
          A.store.Ref(ptr + 0 + 56, undefined);
        } else {
          A.store.Bool(ptr + 0 + 64, true);
          A.store.Bool(ptr + 0 + 60, "ageLowerBound" in x["filter"] ? true : false);
          A.store.Int64(
            ptr + 0 + 0,
            x["filter"]["ageLowerBound"] === undefined ? 0 : (x["filter"]["ageLowerBound"] as number)
          );
          A.store.Bool(ptr + 0 + 61, "ageUpperBound" in x["filter"] ? true : false);
          A.store.Int64(
            ptr + 0 + 8,
            x["filter"]["ageUpperBound"] === undefined ? 0 : (x["filter"]["ageUpperBound"] as number)
          );
          A.store.Ref(ptr + 0 + 16, x["filter"]["domain"]);
          A.store.Ref(ptr + 0 + 20, x["filter"]["expires"]);
          A.store.Ref(ptr + 0 + 24, x["filter"]["httpOnly"]);
          A.store.Bool(ptr + 0 + 62, "maxAge" in x["filter"] ? true : false);
          A.store.Float64(ptr + 0 + 32, x["filter"]["maxAge"] === undefined ? 0 : (x["filter"]["maxAge"] as number));
          A.store.Ref(ptr + 0 + 40, x["filter"]["name"]);
          A.store.Ref(ptr + 0 + 44, x["filter"]["path"]);
          A.store.Ref(ptr + 0 + 48, x["filter"]["secure"]);
          A.store.Bool(ptr + 0 + 63, "sessionCookie" in x["filter"] ? true : false);
          A.store.Bool(ptr + 0 + 52, x["filter"]["sessionCookie"] ? true : false);
          A.store.Ref(ptr + 0 + 56, x["filter"]["value"]);
        }
        A.store.Enum(ptr + 68, ["declarativeWebRequest.EditResponseCookie"].indexOf(x["instanceType"] as string));

        if (typeof x["modification"] === "undefined") {
          A.store.Bool(ptr + 72 + 41, false);
          A.store.Ref(ptr + 72 + 0, undefined);
          A.store.Ref(ptr + 72 + 4, undefined);
          A.store.Ref(ptr + 72 + 8, undefined);
          A.store.Bool(ptr + 72 + 40, false);
          A.store.Float64(ptr + 72 + 16, 0);
          A.store.Ref(ptr + 72 + 24, undefined);
          A.store.Ref(ptr + 72 + 28, undefined);
          A.store.Ref(ptr + 72 + 32, undefined);
          A.store.Ref(ptr + 72 + 36, undefined);
        } else {
          A.store.Bool(ptr + 72 + 41, true);
          A.store.Ref(ptr + 72 + 0, x["modification"]["domain"]);
          A.store.Ref(ptr + 72 + 4, x["modification"]["expires"]);
          A.store.Ref(ptr + 72 + 8, x["modification"]["httpOnly"]);
          A.store.Bool(ptr + 72 + 40, "maxAge" in x["modification"] ? true : false);
          A.store.Float64(
            ptr + 72 + 16,
            x["modification"]["maxAge"] === undefined ? 0 : (x["modification"]["maxAge"] as number)
          );
          A.store.Ref(ptr + 72 + 24, x["modification"]["name"]);
          A.store.Ref(ptr + 72 + 28, x["modification"]["path"]);
          A.store.Ref(ptr + 72 + 32, x["modification"]["secure"]);
          A.store.Ref(ptr + 72 + 36, x["modification"]["value"]);
        }
      }
    },
    "load_EditResponseCookie": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 64)) {
        x["filter"] = {};
        if (A.load.Bool(ptr + 0 + 60)) {
          x["filter"]["ageLowerBound"] = A.load.Int64(ptr + 0 + 0);
        } else {
          delete x["filter"]["ageLowerBound"];
        }
        if (A.load.Bool(ptr + 0 + 61)) {
          x["filter"]["ageUpperBound"] = A.load.Int64(ptr + 0 + 8);
        } else {
          delete x["filter"]["ageUpperBound"];
        }
        x["filter"]["domain"] = A.load.Ref(ptr + 0 + 16, undefined);
        x["filter"]["expires"] = A.load.Ref(ptr + 0 + 20, undefined);
        x["filter"]["httpOnly"] = A.load.Ref(ptr + 0 + 24, undefined);
        if (A.load.Bool(ptr + 0 + 62)) {
          x["filter"]["maxAge"] = A.load.Float64(ptr + 0 + 32);
        } else {
          delete x["filter"]["maxAge"];
        }
        x["filter"]["name"] = A.load.Ref(ptr + 0 + 40, undefined);
        x["filter"]["path"] = A.load.Ref(ptr + 0 + 44, undefined);
        x["filter"]["secure"] = A.load.Ref(ptr + 0 + 48, undefined);
        if (A.load.Bool(ptr + 0 + 63)) {
          x["filter"]["sessionCookie"] = A.load.Bool(ptr + 0 + 52);
        } else {
          delete x["filter"]["sessionCookie"];
        }
        x["filter"]["value"] = A.load.Ref(ptr + 0 + 56, undefined);
      } else {
        delete x["filter"];
      }
      x["instanceType"] = A.load.Enum(ptr + 68, ["declarativeWebRequest.EditResponseCookie"]);
      if (A.load.Bool(ptr + 72 + 41)) {
        x["modification"] = {};
        x["modification"]["domain"] = A.load.Ref(ptr + 72 + 0, undefined);
        x["modification"]["expires"] = A.load.Ref(ptr + 72 + 4, undefined);
        x["modification"]["httpOnly"] = A.load.Ref(ptr + 72 + 8, undefined);
        if (A.load.Bool(ptr + 72 + 40)) {
          x["modification"]["maxAge"] = A.load.Float64(ptr + 72 + 16);
        } else {
          delete x["modification"]["maxAge"];
        }
        x["modification"]["name"] = A.load.Ref(ptr + 72 + 24, undefined);
        x["modification"]["path"] = A.load.Ref(ptr + 72 + 28, undefined);
        x["modification"]["secure"] = A.load.Ref(ptr + 72 + 32, undefined);
        x["modification"]["value"] = A.load.Ref(ptr + 72 + 36, undefined);
      } else {
        delete x["modification"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_HeaderFilter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 32, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
      } else {
        A.store.Bool(ptr + 32, true);
        A.store.Ref(ptr + 0, x["nameContains"]);
        A.store.Ref(ptr + 4, x["nameEquals"]);
        A.store.Ref(ptr + 8, x["namePrefix"]);
        A.store.Ref(ptr + 12, x["nameSuffix"]);
        A.store.Ref(ptr + 16, x["valueContains"]);
        A.store.Ref(ptr + 20, x["valueEquals"]);
        A.store.Ref(ptr + 24, x["valuePrefix"]);
        A.store.Ref(ptr + 28, x["valueSuffix"]);
      }
    },
    "load_HeaderFilter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["nameContains"] = A.load.Ref(ptr + 0, undefined);
      x["nameEquals"] = A.load.Ref(ptr + 4, undefined);
      x["namePrefix"] = A.load.Ref(ptr + 8, undefined);
      x["nameSuffix"] = A.load.Ref(ptr + 12, undefined);
      x["valueContains"] = A.load.Ref(ptr + 16, undefined);
      x["valueEquals"] = A.load.Ref(ptr + 20, undefined);
      x["valuePrefix"] = A.load.Ref(ptr + 24, undefined);
      x["valueSuffix"] = A.load.Ref(ptr + 28, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_IgnoreRulesInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.IgnoreRules"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_IgnoreRules": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["hasTag"]);
        A.store.Enum(ptr + 4, ["declarativeWebRequest.IgnoreRules"].indexOf(x["instanceType"] as string));
        A.store.Bool(ptr + 16, "lowerPriorityThan" in x ? true : false);
        A.store.Int64(ptr + 8, x["lowerPriorityThan"] === undefined ? 0 : (x["lowerPriorityThan"] as number));
      }
    },
    "load_IgnoreRules": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["hasTag"] = A.load.Ref(ptr + 0, undefined);
      x["instanceType"] = A.load.Enum(ptr + 4, ["declarativeWebRequest.IgnoreRules"]);
      if (A.load.Bool(ptr + 16)) {
        x["lowerPriorityThan"] = A.load.Int64(ptr + 8);
      } else {
        delete x["lowerPriorityThan"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_Stage": (ref: heap.Ref<string>): number => {
      const idx = ["onBeforeRequest", "onBeforeSendHeaders", "onHeadersReceived", "onAuthRequired"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OnMessageArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 72, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Int64(ptr + 32, 0);
        A.store.Ref(ptr + 40, undefined);
        A.store.Enum(ptr + 44, -1);
        A.store.Int64(ptr + 48, 0);
        A.store.Float64(ptr + 56, 0);
        A.store.Enum(ptr + 64, -1);
        A.store.Ref(ptr + 68, undefined);
      } else {
        A.store.Bool(ptr + 72, true);
        A.store.Ref(ptr + 0, x["documentId"]);
        A.store.Enum(
          ptr + 4,
          ["prerender", "active", "cached", "pending_deletion"].indexOf(x["documentLifecycle"] as string)
        );
        A.store.Int64(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(x["frameType"] as string));
        A.store.Ref(ptr + 20, x["message"]);
        A.store.Ref(ptr + 24, x["method"]);
        A.store.Ref(ptr + 28, x["parentDocumentId"]);
        A.store.Int64(ptr + 32, x["parentFrameId"] === undefined ? 0 : (x["parentFrameId"] as number));
        A.store.Ref(ptr + 40, x["requestId"]);
        A.store.Enum(
          ptr + 44,
          ["onBeforeRequest", "onBeforeSendHeaders", "onHeadersReceived", "onAuthRequired"].indexOf(
            x["stage"] as string
          )
        );
        A.store.Int64(ptr + 48, x["tabId"] === undefined ? 0 : (x["tabId"] as number));
        A.store.Float64(ptr + 56, x["timeStamp"] === undefined ? 0 : (x["timeStamp"] as number));
        A.store.Enum(
          ptr + 64,
          [
            "main_frame",
            "sub_frame",
            "stylesheet",
            "script",
            "image",
            "font",
            "object",
            "xmlhttprequest",
            "ping",
            "csp_report",
            "media",
            "websocket",
            "webbundle",
            "other",
          ].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 68, x["url"]);
      }
    },
    "load_OnMessageArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["documentId"] = A.load.Ref(ptr + 0, undefined);
      x["documentLifecycle"] = A.load.Enum(ptr + 4, ["prerender", "active", "cached", "pending_deletion"]);
      x["frameId"] = A.load.Int64(ptr + 8);
      x["frameType"] = A.load.Enum(ptr + 16, ["outermost_frame", "fenced_frame", "sub_frame"]);
      x["message"] = A.load.Ref(ptr + 20, undefined);
      x["method"] = A.load.Ref(ptr + 24, undefined);
      x["parentDocumentId"] = A.load.Ref(ptr + 28, undefined);
      x["parentFrameId"] = A.load.Int64(ptr + 32);
      x["requestId"] = A.load.Ref(ptr + 40, undefined);
      x["stage"] = A.load.Enum(ptr + 44, [
        "onBeforeRequest",
        "onBeforeSendHeaders",
        "onHeadersReceived",
        "onAuthRequired",
      ]);
      x["tabId"] = A.load.Int64(ptr + 48);
      x["timeStamp"] = A.load.Float64(ptr + 56);
      x["type"] = A.load.Enum(ptr + 64, [
        "main_frame",
        "sub_frame",
        "stylesheet",
        "script",
        "image",
        "font",
        "object",
        "xmlhttprequest",
        "ping",
        "csp_report",
        "media",
        "websocket",
        "webbundle",
        "other",
      ]);
      x["url"] = A.load.Ref(ptr + 68, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RedirectByRegExInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.RedirectByRegEx"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RedirectByRegEx": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["from"]);
        A.store.Enum(ptr + 4, ["declarativeWebRequest.RedirectByRegEx"].indexOf(x["instanceType"] as string));
        A.store.Ref(ptr + 8, x["to"]);
      }
    },
    "load_RedirectByRegEx": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["from"] = A.load.Ref(ptr + 0, undefined);
      x["instanceType"] = A.load.Enum(ptr + 4, ["declarativeWebRequest.RedirectByRegEx"]);
      x["to"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RedirectRequestInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.RedirectRequest"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RedirectRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["declarativeWebRequest.RedirectRequest"].indexOf(x["instanceType"] as string));
        A.store.Ref(ptr + 4, x["redirectUrl"]);
      }
    },
    "load_RedirectRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["instanceType"] = A.load.Enum(ptr + 0, ["declarativeWebRequest.RedirectRequest"]);
      x["redirectUrl"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RedirectToEmptyDocumentInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.RedirectToEmptyDocument"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RedirectToEmptyDocument": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["declarativeWebRequest.RedirectToEmptyDocument"].indexOf(x["instanceType"] as string));
      }
    },
    "load_RedirectToEmptyDocument": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["instanceType"] = A.load.Enum(ptr + 0, ["declarativeWebRequest.RedirectToEmptyDocument"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RedirectToTransparentImageInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.RedirectToTransparentImage"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RedirectToTransparentImage": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(
          ptr + 0,
          ["declarativeWebRequest.RedirectToTransparentImage"].indexOf(x["instanceType"] as string)
        );
      }
    },
    "load_RedirectToTransparentImage": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["instanceType"] = A.load.Enum(ptr + 0, ["declarativeWebRequest.RedirectToTransparentImage"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RemoveRequestCookieInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.RemoveRequestCookie"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RemoveRequestCookie": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);

        A.store.Bool(ptr + 0 + 8, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Enum(ptr + 12, -1);
      } else {
        A.store.Bool(ptr + 16, true);

        if (typeof x["filter"] === "undefined") {
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
        } else {
          A.store.Bool(ptr + 0 + 8, true);
          A.store.Ref(ptr + 0 + 0, x["filter"]["name"]);
          A.store.Ref(ptr + 0 + 4, x["filter"]["value"]);
        }
        A.store.Enum(ptr + 12, ["declarativeWebRequest.RemoveRequestCookie"].indexOf(x["instanceType"] as string));
      }
    },
    "load_RemoveRequestCookie": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 8)) {
        x["filter"] = {};
        x["filter"]["name"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["filter"]["value"] = A.load.Ref(ptr + 0 + 4, undefined);
      } else {
        delete x["filter"];
      }
      x["instanceType"] = A.load.Enum(ptr + 12, ["declarativeWebRequest.RemoveRequestCookie"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RemoveRequestHeaderInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.RemoveRequestHeader"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RemoveRequestHeader": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["declarativeWebRequest.RemoveRequestHeader"].indexOf(x["instanceType"] as string));
        A.store.Ref(ptr + 4, x["name"]);
      }
    },
    "load_RemoveRequestHeader": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["instanceType"] = A.load.Enum(ptr + 0, ["declarativeWebRequest.RemoveRequestHeader"]);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RemoveResponseCookieInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.RemoveResponseCookie"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RemoveResponseCookie": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 72, false);

        A.store.Bool(ptr + 0 + 64, false);
        A.store.Bool(ptr + 0 + 60, false);
        A.store.Int64(ptr + 0 + 0, 0);
        A.store.Bool(ptr + 0 + 61, false);
        A.store.Int64(ptr + 0 + 8, 0);
        A.store.Ref(ptr + 0 + 16, undefined);
        A.store.Ref(ptr + 0 + 20, undefined);
        A.store.Ref(ptr + 0 + 24, undefined);
        A.store.Bool(ptr + 0 + 62, false);
        A.store.Float64(ptr + 0 + 32, 0);
        A.store.Ref(ptr + 0 + 40, undefined);
        A.store.Ref(ptr + 0 + 44, undefined);
        A.store.Ref(ptr + 0 + 48, undefined);
        A.store.Bool(ptr + 0 + 63, false);
        A.store.Bool(ptr + 0 + 52, false);
        A.store.Ref(ptr + 0 + 56, undefined);
        A.store.Enum(ptr + 68, -1);
      } else {
        A.store.Bool(ptr + 72, true);

        if (typeof x["filter"] === "undefined") {
          A.store.Bool(ptr + 0 + 64, false);
          A.store.Bool(ptr + 0 + 60, false);
          A.store.Int64(ptr + 0 + 0, 0);
          A.store.Bool(ptr + 0 + 61, false);
          A.store.Int64(ptr + 0 + 8, 0);
          A.store.Ref(ptr + 0 + 16, undefined);
          A.store.Ref(ptr + 0 + 20, undefined);
          A.store.Ref(ptr + 0 + 24, undefined);
          A.store.Bool(ptr + 0 + 62, false);
          A.store.Float64(ptr + 0 + 32, 0);
          A.store.Ref(ptr + 0 + 40, undefined);
          A.store.Ref(ptr + 0 + 44, undefined);
          A.store.Ref(ptr + 0 + 48, undefined);
          A.store.Bool(ptr + 0 + 63, false);
          A.store.Bool(ptr + 0 + 52, false);
          A.store.Ref(ptr + 0 + 56, undefined);
        } else {
          A.store.Bool(ptr + 0 + 64, true);
          A.store.Bool(ptr + 0 + 60, "ageLowerBound" in x["filter"] ? true : false);
          A.store.Int64(
            ptr + 0 + 0,
            x["filter"]["ageLowerBound"] === undefined ? 0 : (x["filter"]["ageLowerBound"] as number)
          );
          A.store.Bool(ptr + 0 + 61, "ageUpperBound" in x["filter"] ? true : false);
          A.store.Int64(
            ptr + 0 + 8,
            x["filter"]["ageUpperBound"] === undefined ? 0 : (x["filter"]["ageUpperBound"] as number)
          );
          A.store.Ref(ptr + 0 + 16, x["filter"]["domain"]);
          A.store.Ref(ptr + 0 + 20, x["filter"]["expires"]);
          A.store.Ref(ptr + 0 + 24, x["filter"]["httpOnly"]);
          A.store.Bool(ptr + 0 + 62, "maxAge" in x["filter"] ? true : false);
          A.store.Float64(ptr + 0 + 32, x["filter"]["maxAge"] === undefined ? 0 : (x["filter"]["maxAge"] as number));
          A.store.Ref(ptr + 0 + 40, x["filter"]["name"]);
          A.store.Ref(ptr + 0 + 44, x["filter"]["path"]);
          A.store.Ref(ptr + 0 + 48, x["filter"]["secure"]);
          A.store.Bool(ptr + 0 + 63, "sessionCookie" in x["filter"] ? true : false);
          A.store.Bool(ptr + 0 + 52, x["filter"]["sessionCookie"] ? true : false);
          A.store.Ref(ptr + 0 + 56, x["filter"]["value"]);
        }
        A.store.Enum(ptr + 68, ["declarativeWebRequest.RemoveResponseCookie"].indexOf(x["instanceType"] as string));
      }
    },
    "load_RemoveResponseCookie": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 64)) {
        x["filter"] = {};
        if (A.load.Bool(ptr + 0 + 60)) {
          x["filter"]["ageLowerBound"] = A.load.Int64(ptr + 0 + 0);
        } else {
          delete x["filter"]["ageLowerBound"];
        }
        if (A.load.Bool(ptr + 0 + 61)) {
          x["filter"]["ageUpperBound"] = A.load.Int64(ptr + 0 + 8);
        } else {
          delete x["filter"]["ageUpperBound"];
        }
        x["filter"]["domain"] = A.load.Ref(ptr + 0 + 16, undefined);
        x["filter"]["expires"] = A.load.Ref(ptr + 0 + 20, undefined);
        x["filter"]["httpOnly"] = A.load.Ref(ptr + 0 + 24, undefined);
        if (A.load.Bool(ptr + 0 + 62)) {
          x["filter"]["maxAge"] = A.load.Float64(ptr + 0 + 32);
        } else {
          delete x["filter"]["maxAge"];
        }
        x["filter"]["name"] = A.load.Ref(ptr + 0 + 40, undefined);
        x["filter"]["path"] = A.load.Ref(ptr + 0 + 44, undefined);
        x["filter"]["secure"] = A.load.Ref(ptr + 0 + 48, undefined);
        if (A.load.Bool(ptr + 0 + 63)) {
          x["filter"]["sessionCookie"] = A.load.Bool(ptr + 0 + 52);
        } else {
          delete x["filter"]["sessionCookie"];
        }
        x["filter"]["value"] = A.load.Ref(ptr + 0 + 56, undefined);
      } else {
        delete x["filter"];
      }
      x["instanceType"] = A.load.Enum(ptr + 68, ["declarativeWebRequest.RemoveResponseCookie"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RemoveResponseHeaderInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.RemoveResponseHeader"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RemoveResponseHeader": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Enum(ptr + 0, ["declarativeWebRequest.RemoveResponseHeader"].indexOf(x["instanceType"] as string));
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Ref(ptr + 8, x["value"]);
      }
    },
    "load_RemoveResponseHeader": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["instanceType"] = A.load.Enum(ptr + 0, ["declarativeWebRequest.RemoveResponseHeader"]);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      x["value"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RequestMatcherInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.RequestMatcher"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RequestMatcher": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 206, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);

        A.store.Bool(ptr + 16 + 0, false);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Enum(ptr + 100, -1);
        A.store.Ref(ptr + 104, undefined);
        A.store.Ref(ptr + 108, undefined);
        A.store.Ref(ptr + 112, undefined);
        A.store.Ref(ptr + 116, undefined);
        A.store.Bool(ptr + 205, false);
        A.store.Bool(ptr + 120, false);

        A.store.Bool(ptr + 124 + 0, false);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
        A.store.Ref(ptr + 124 + 0, undefined);
      } else {
        A.store.Bool(ptr + 206, true);
        A.store.Ref(ptr + 0, x["contentType"]);
        A.store.Ref(ptr + 4, x["excludeContentType"]);
        A.store.Ref(ptr + 8, x["excludeRequestHeaders"]);
        A.store.Ref(ptr + 12, x["excludeResponseHeaders"]);

        if (typeof x["firstPartyForCookiesUrl"] === "undefined") {
          A.store.Bool(ptr + 16 + 0, false);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 0, undefined);
        } else {
          A.store.Bool(ptr + 16 + 0, true);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["hostContains"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["hostEquals"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["hostPrefix"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["hostSuffix"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["originAndPathMatches"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["pathContains"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["pathEquals"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["pathPrefix"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["pathSuffix"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["ports"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["queryContains"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["queryEquals"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["queryPrefix"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["querySuffix"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["schemes"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["urlContains"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["urlEquals"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["urlMatches"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["urlPrefix"]);
          A.store.Ref(ptr + 16 + 0, x["firstPartyForCookiesUrl"]["urlSuffix"]);
        }
        A.store.Enum(ptr + 100, ["declarativeWebRequest.RequestMatcher"].indexOf(x["instanceType"] as string));
        A.store.Ref(ptr + 104, x["requestHeaders"]);
        A.store.Ref(ptr + 108, x["resourceType"]);
        A.store.Ref(ptr + 112, x["responseHeaders"]);
        A.store.Ref(ptr + 116, x["stages"]);
        A.store.Bool(ptr + 205, "thirdPartyForCookies" in x ? true : false);
        A.store.Bool(ptr + 120, x["thirdPartyForCookies"] ? true : false);

        if (typeof x["url"] === "undefined") {
          A.store.Bool(ptr + 124 + 0, false);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
          A.store.Ref(ptr + 124 + 0, undefined);
        } else {
          A.store.Bool(ptr + 124 + 0, true);
          A.store.Ref(ptr + 124 + 0, x["url"]["hostContains"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["hostEquals"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["hostPrefix"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["hostSuffix"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["originAndPathMatches"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["pathContains"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["pathEquals"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["pathPrefix"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["pathSuffix"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["ports"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["queryContains"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["queryEquals"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["queryPrefix"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["querySuffix"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["schemes"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["urlContains"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["urlEquals"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["urlMatches"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["urlPrefix"]);
          A.store.Ref(ptr + 124 + 0, x["url"]["urlSuffix"]);
        }
      }
    },
    "load_RequestMatcher": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["contentType"] = A.load.Ref(ptr + 0, undefined);
      x["excludeContentType"] = A.load.Ref(ptr + 4, undefined);
      x["excludeRequestHeaders"] = A.load.Ref(ptr + 8, undefined);
      x["excludeResponseHeaders"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 16 + 0)) {
        x["firstPartyForCookiesUrl"] = {};
        x["firstPartyForCookiesUrl"]["hostContains"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["hostEquals"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["hostPrefix"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["hostSuffix"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["originAndPathMatches"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["pathContains"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["pathEquals"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["pathPrefix"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["pathSuffix"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["ports"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["queryContains"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["queryEquals"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["queryPrefix"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["querySuffix"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["schemes"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["urlContains"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["urlEquals"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["urlMatches"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["urlPrefix"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["firstPartyForCookiesUrl"]["urlSuffix"] = A.load.Ref(ptr + 16 + 0, undefined);
      } else {
        delete x["firstPartyForCookiesUrl"];
      }
      x["instanceType"] = A.load.Enum(ptr + 100, ["declarativeWebRequest.RequestMatcher"]);
      x["requestHeaders"] = A.load.Ref(ptr + 104, undefined);
      x["resourceType"] = A.load.Ref(ptr + 108, undefined);
      x["responseHeaders"] = A.load.Ref(ptr + 112, undefined);
      x["stages"] = A.load.Ref(ptr + 116, undefined);
      if (A.load.Bool(ptr + 205)) {
        x["thirdPartyForCookies"] = A.load.Bool(ptr + 120);
      } else {
        delete x["thirdPartyForCookies"];
      }
      if (A.load.Bool(ptr + 124 + 0)) {
        x["url"] = {};
        x["url"]["hostContains"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["hostEquals"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["hostPrefix"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["hostSuffix"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["originAndPathMatches"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["pathContains"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["pathEquals"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["pathPrefix"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["pathSuffix"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["ports"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["queryContains"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["queryEquals"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["queryPrefix"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["querySuffix"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["schemes"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["urlContains"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["urlEquals"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["urlMatches"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["urlPrefix"] = A.load.Ref(ptr + 124 + 0, undefined);
        x["url"]["urlSuffix"] = A.load.Ref(ptr + 124 + 0, undefined);
      } else {
        delete x["url"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_SendMessageToExtensionInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.SendMessageToExtension"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SendMessageToExtension": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["declarativeWebRequest.SendMessageToExtension"].indexOf(x["instanceType"] as string));
        A.store.Ref(ptr + 4, x["message"]);
      }
    },
    "load_SendMessageToExtension": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["instanceType"] = A.load.Enum(ptr + 0, ["declarativeWebRequest.SendMessageToExtension"]);
      x["message"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_SetRequestHeaderInstanceType": (ref: heap.Ref<string>): number => {
      const idx = ["declarativeWebRequest.SetRequestHeader"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SetRequestHeader": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Enum(ptr + 0, ["declarativeWebRequest.SetRequestHeader"].indexOf(x["instanceType"] as string));
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Ref(ptr + 8, x["value"]);
      }
    },
    "load_SetRequestHeader": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["instanceType"] = A.load.Enum(ptr + 0, ["declarativeWebRequest.SetRequestHeader"]);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      x["value"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_OnMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeWebRequest?.onMessage && "addListener" in WEBEXT?.declarativeWebRequest?.onMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeWebRequest.onMessage.addListener);
    },
    "call_OnMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.declarativeWebRequest.onMessage.addListener(A.H.get<object>(callback));
    },
    "try_OnMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.declarativeWebRequest.onMessage.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeWebRequest?.onMessage && "removeListener" in WEBEXT?.declarativeWebRequest?.onMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeWebRequest.onMessage.removeListener);
    },
    "call_OffMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.declarativeWebRequest.onMessage.removeListener(A.H.get<object>(callback));
    },
    "try_OffMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.declarativeWebRequest.onMessage.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeWebRequest?.onMessage && "hasListener" in WEBEXT?.declarativeWebRequest?.onMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeWebRequest.onMessage.hasListener);
    },
    "call_HasOnMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.declarativeWebRequest.onMessage.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.declarativeWebRequest.onMessage.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeWebRequest?.onRequest && "addListener" in WEBEXT?.declarativeWebRequest?.onRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeWebRequest.onRequest.addListener);
    },
    "call_OnRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.declarativeWebRequest.onRequest.addListener(A.H.get<object>(callback));
    },
    "try_OnRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.declarativeWebRequest.onRequest.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeWebRequest?.onRequest && "removeListener" in WEBEXT?.declarativeWebRequest?.onRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeWebRequest.onRequest.removeListener);
    },
    "call_OffRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.declarativeWebRequest.onRequest.removeListener(A.H.get<object>(callback));
    },
    "try_OffRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.declarativeWebRequest.onRequest.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRequest": (): heap.Ref<boolean> => {
      if (WEBEXT?.declarativeWebRequest?.onRequest && "hasListener" in WEBEXT?.declarativeWebRequest?.onRequest) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.declarativeWebRequest.onRequest.hasListener);
    },
    "call_HasOnRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.declarativeWebRequest.onRequest.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRequest": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.declarativeWebRequest.onRequest.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
