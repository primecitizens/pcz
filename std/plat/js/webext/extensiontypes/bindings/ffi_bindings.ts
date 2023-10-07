import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/extensiontypes", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_CSSOrigin": (ref: heap.Ref<string>): number => {
      const idx = ["author", "user"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ExecutionWorld": (ref: heap.Ref<string>): number => {
      const idx = ["ISOLATED", "MAIN"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_RunAt": (ref: heap.Ref<string>): number => {
      const idx = ["document_start", "document_end", "document_idle"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DeleteInjectionDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 26, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Bool(ptr + 27, false);
        A.store.Bool(ptr + 24, false);
      } else {
        A.store.Bool(ptr + 28, true);
        A.store.Bool(ptr + 25, "allFrames" in x ? true : false);
        A.store.Bool(ptr + 0, x["allFrames"] ? true : false);
        A.store.Ref(ptr + 4, x["code"]);
        A.store.Enum(ptr + 8, ["author", "user"].indexOf(x["cssOrigin"] as string));
        A.store.Ref(ptr + 12, x["file"]);
        A.store.Bool(ptr + 26, "frameId" in x ? true : false);
        A.store.Int64(ptr + 16, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Bool(ptr + 27, "matchAboutBlank" in x ? true : false);
        A.store.Bool(ptr + 24, x["matchAboutBlank"] ? true : false);
      }
    },
    "load_DeleteInjectionDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 25)) {
        x["allFrames"] = A.load.Bool(ptr + 0);
      } else {
        delete x["allFrames"];
      }
      x["code"] = A.load.Ref(ptr + 4, undefined);
      x["cssOrigin"] = A.load.Enum(ptr + 8, ["author", "user"]);
      x["file"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 26)) {
        x["frameId"] = A.load.Int64(ptr + 16);
      } else {
        delete x["frameId"];
      }
      if (A.load.Bool(ptr + 27)) {
        x["matchAboutBlank"] = A.load.Bool(ptr + 24);
      } else {
        delete x["matchAboutBlank"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DocumentLifecycle": (ref: heap.Ref<string>): number => {
      const idx = ["prerender", "active", "cached", "pending_deletion"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_FrameType": (ref: heap.Ref<string>): number => {
      const idx = ["outermost_frame", "fenced_frame", "sub_frame"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ImageFormat": (ref: heap.Ref<string>): number => {
      const idx = ["jpeg", "png"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ImageDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Enum(ptr + 0, ["jpeg", "png"].indexOf(x["format"] as string));
        A.store.Bool(ptr + 16, "quality" in x ? true : false);
        A.store.Int64(ptr + 8, x["quality"] === undefined ? 0 : (x["quality"] as number));
      }
    },
    "load_ImageDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["format"] = A.load.Enum(ptr + 0, ["jpeg", "png"]);
      if (A.load.Bool(ptr + 16)) {
        x["quality"] = A.load.Int64(ptr + 8);
      } else {
        delete x["quality"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_InjectDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 35, false);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 33, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Bool(ptr + 34, false);
        A.store.Bool(ptr + 24, false);
        A.store.Enum(ptr + 28, -1);
      } else {
        A.store.Bool(ptr + 35, true);
        A.store.Bool(ptr + 32, "allFrames" in x ? true : false);
        A.store.Bool(ptr + 0, x["allFrames"] ? true : false);
        A.store.Ref(ptr + 4, x["code"]);
        A.store.Enum(ptr + 8, ["author", "user"].indexOf(x["cssOrigin"] as string));
        A.store.Ref(ptr + 12, x["file"]);
        A.store.Bool(ptr + 33, "frameId" in x ? true : false);
        A.store.Int64(ptr + 16, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Bool(ptr + 34, "matchAboutBlank" in x ? true : false);
        A.store.Bool(ptr + 24, x["matchAboutBlank"] ? true : false);
        A.store.Enum(ptr + 28, ["document_start", "document_end", "document_idle"].indexOf(x["runAt"] as string));
      }
    },
    "load_InjectDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 32)) {
        x["allFrames"] = A.load.Bool(ptr + 0);
      } else {
        delete x["allFrames"];
      }
      x["code"] = A.load.Ref(ptr + 4, undefined);
      x["cssOrigin"] = A.load.Enum(ptr + 8, ["author", "user"]);
      x["file"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 33)) {
        x["frameId"] = A.load.Int64(ptr + 16);
      } else {
        delete x["frameId"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["matchAboutBlank"] = A.load.Bool(ptr + 24);
      } else {
        delete x["matchAboutBlank"];
      }
      x["runAt"] = A.load.Enum(ptr + 28, ["document_start", "document_end", "document_idle"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
  };
});
