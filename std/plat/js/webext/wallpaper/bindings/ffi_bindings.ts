import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/wallpaper", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_WallpaperLayout": (ref: heap.Ref<string>): number => {
      const idx = ["STRETCH", "CENTER", "CENTER_CROPPED"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SetWallpaperArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Ref(ptr + 0, x["data"]);
        A.store.Ref(ptr + 4, x["filename"]);
        A.store.Enum(ptr + 8, ["STRETCH", "CENTER", "CENTER_CROPPED"].indexOf(x["layout"] as string));
        A.store.Bool(ptr + 20, "thumbnail" in x ? true : false);
        A.store.Bool(ptr + 12, x["thumbnail"] ? true : false);
        A.store.Ref(ptr + 16, x["url"]);
      }
    },
    "load_SetWallpaperArgDetails": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["data"] = A.load.Ref(ptr + 0, undefined);
      x["filename"] = A.load.Ref(ptr + 4, undefined);
      x["layout"] = A.load.Enum(ptr + 8, ["STRETCH", "CENTER", "CENTER_CROPPED"]);
      if (A.load.Bool(ptr + 20)) {
        x["thumbnail"] = A.load.Bool(ptr + 12);
      } else {
        delete x["thumbnail"];
      }
      x["url"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_SetWallpaper": (): heap.Ref<boolean> => {
      if (WEBEXT?.wallpaper && "setWallpaper" in WEBEXT?.wallpaper) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetWallpaper": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wallpaper.setWallpaper);
    },
    "call_SetWallpaper": (retPtr: Pointer, details: Pointer): void => {
      const details_ffi = {};

      details_ffi["data"] = A.load.Ref(details + 0, undefined);
      details_ffi["filename"] = A.load.Ref(details + 4, undefined);
      details_ffi["layout"] = A.load.Enum(details + 8, ["STRETCH", "CENTER", "CENTER_CROPPED"]);
      if (A.load.Bool(details + 20)) {
        details_ffi["thumbnail"] = A.load.Bool(details + 12);
      }
      details_ffi["url"] = A.load.Ref(details + 16, undefined);

      const _ret = WEBEXT.wallpaper.setWallpaper(details_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetWallpaper": (retPtr: Pointer, errPtr: Pointer, details: Pointer): heap.Ref<boolean> => {
      try {
        const details_ffi = {};

        details_ffi["data"] = A.load.Ref(details + 0, undefined);
        details_ffi["filename"] = A.load.Ref(details + 4, undefined);
        details_ffi["layout"] = A.load.Enum(details + 8, ["STRETCH", "CENTER", "CENTER_CROPPED"]);
        if (A.load.Bool(details + 20)) {
          details_ffi["thumbnail"] = A.load.Bool(details + 12);
        }
        details_ffi["url"] = A.load.Ref(details + 16, undefined);

        const _ret = WEBEXT.wallpaper.setWallpaper(details_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
