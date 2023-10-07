import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/incognito", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_IncognitoMode": (ref: heap.Ref<string>): number => {
      const idx = ["split", "spanning", "not_allowed"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
  };
});
