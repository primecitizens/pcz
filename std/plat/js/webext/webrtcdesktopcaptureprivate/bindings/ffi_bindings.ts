import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/webrtcdesktopcaptureprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_DesktopCaptureSourceType": (ref: heap.Ref<string>): number => {
      const idx = ["screen", "window", "tab"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RequestInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "guestProcessId" in x ? true : false);
        A.store.Int32(ptr + 0, x["guestProcessId"] === undefined ? 0 : (x["guestProcessId"] as number));
        A.store.Bool(ptr + 9, "guestRenderFrameId" in x ? true : false);
        A.store.Int32(ptr + 4, x["guestRenderFrameId"] === undefined ? 0 : (x["guestRenderFrameId"] as number));
      }
    },
    "load_RequestInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["guestProcessId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["guestProcessId"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["guestRenderFrameId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["guestRenderFrameId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_CancelChooseDesktopMedia": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcDesktopCapturePrivate && "cancelChooseDesktopMedia" in WEBEXT?.webrtcDesktopCapturePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CancelChooseDesktopMedia": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcDesktopCapturePrivate.cancelChooseDesktopMedia);
    },
    "call_CancelChooseDesktopMedia": (retPtr: Pointer, desktopMediaRequestId: number): void => {
      const _ret = WEBEXT.webrtcDesktopCapturePrivate.cancelChooseDesktopMedia(desktopMediaRequestId);
    },
    "try_CancelChooseDesktopMedia": (
      retPtr: Pointer,
      errPtr: Pointer,
      desktopMediaRequestId: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webrtcDesktopCapturePrivate.cancelChooseDesktopMedia(desktopMediaRequestId);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ChooseDesktopMedia": (): heap.Ref<boolean> => {
      if (WEBEXT?.webrtcDesktopCapturePrivate && "chooseDesktopMedia" in WEBEXT?.webrtcDesktopCapturePrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ChooseDesktopMedia": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webrtcDesktopCapturePrivate.chooseDesktopMedia);
    },
    "call_ChooseDesktopMedia": (
      retPtr: Pointer,
      sources: heap.Ref<object>,
      request: Pointer,
      callback: heap.Ref<object>
    ): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 8)) {
        request_ffi["guestProcessId"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 9)) {
        request_ffi["guestRenderFrameId"] = A.load.Int32(request + 4);
      }

      const _ret = WEBEXT.webrtcDesktopCapturePrivate.chooseDesktopMedia(
        A.H.get<object>(sources),
        request_ffi,
        A.H.get<object>(callback)
      );
      A.store.Int32(retPtr, _ret);
    },
    "try_ChooseDesktopMedia": (
      retPtr: Pointer,
      errPtr: Pointer,
      sources: heap.Ref<object>,
      request: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 8)) {
          request_ffi["guestProcessId"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 9)) {
          request_ffi["guestRenderFrameId"] = A.load.Int32(request + 4);
        }

        const _ret = WEBEXT.webrtcDesktopCapturePrivate.chooseDesktopMedia(
          A.H.get<object>(sources),
          request_ffi,
          A.H.get<object>(callback)
        );
        A.store.Int32(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
