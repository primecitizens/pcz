import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/requirements", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_ThreeDFeature": (ref: heap.Ref<string>): number => {
      const idx = ["webgl", "css3d"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
  };
});
