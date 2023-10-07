import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/desktopcapture", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_ChooseDesktopMediaArgCallbackArgOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 1, true);
        A.store.Bool(ptr + 0, x["canRequestAudioTrack"] ? true : false);
      }
    },
    "load_ChooseDesktopMediaArgCallbackArgOptions": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["canRequestAudioTrack"] = A.load.Bool(ptr + 0);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_SelfCapturePreferenceEnum": (ref: heap.Ref<string>): number => {
      const idx = ["include", "exclude"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SystemAudioPreferenceEnum": (ref: heap.Ref<string>): number => {
      const idx = ["include", "exclude"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ChooseDesktopMediaArgOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Enum(ptr + 0, ["include", "exclude"].indexOf(x["selfBrowserSurface"] as string));
        A.store.Bool(ptr + 12, "suppressLocalAudioPlaybackIntended" in x ? true : false);
        A.store.Bool(ptr + 4, x["suppressLocalAudioPlaybackIntended"] ? true : false);
        A.store.Enum(ptr + 8, ["include", "exclude"].indexOf(x["systemAudio"] as string));
      }
    },
    "load_ChooseDesktopMediaArgOptions": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["selfBrowserSurface"] = A.load.Enum(ptr + 0, ["include", "exclude"]);
      if (A.load.Bool(ptr + 12)) {
        x["suppressLocalAudioPlaybackIntended"] = A.load.Bool(ptr + 4);
      } else {
        delete x["suppressLocalAudioPlaybackIntended"];
      }
      x["systemAudio"] = A.load.Enum(ptr + 8, ["include", "exclude"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DesktopCaptureSourceType": (ref: heap.Ref<string>): number => {
      const idx = ["screen", "window", "tab", "audio"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_CancelChooseDesktopMedia": (): heap.Ref<boolean> => {
      if (WEBEXT?.desktopCapture && "cancelChooseDesktopMedia" in WEBEXT?.desktopCapture) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CancelChooseDesktopMedia": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.desktopCapture.cancelChooseDesktopMedia);
    },
    "call_CancelChooseDesktopMedia": (retPtr: Pointer, desktopMediaRequestId: number): void => {
      const _ret = WEBEXT.desktopCapture.cancelChooseDesktopMedia(desktopMediaRequestId);
    },
    "try_CancelChooseDesktopMedia": (
      retPtr: Pointer,
      errPtr: Pointer,
      desktopMediaRequestId: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.desktopCapture.cancelChooseDesktopMedia(desktopMediaRequestId);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ChooseDesktopMedia": (): heap.Ref<boolean> => {
      if (WEBEXT?.desktopCapture && "chooseDesktopMedia" in WEBEXT?.desktopCapture) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ChooseDesktopMedia": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.desktopCapture.chooseDesktopMedia);
    },
    "call_ChooseDesktopMedia": (
      retPtr: Pointer,
      sources: heap.Ref<object>,
      targetTab: Pointer,
      options: Pointer,
      callback: heap.Ref<object>
    ): void => {
      const targetTab_ffi = {};

      targetTab_ffi["active"] = A.load.Bool(targetTab + 0);
      if (A.load.Bool(targetTab + 0)) {
        targetTab_ffi["audible"] = A.load.Bool(targetTab + 0);
      }
      targetTab_ffi["autoDiscardable"] = A.load.Bool(targetTab + 0);
      targetTab_ffi["discarded"] = A.load.Bool(targetTab + 0);
      targetTab_ffi["favIconUrl"] = A.load.Ref(targetTab + 0, undefined);
      targetTab_ffi["groupId"] = A.load.Int0(targetTab + 0);
      if (A.load.Bool(targetTab + 0)) {
        targetTab_ffi["height"] = A.load.Int0(targetTab + 0);
      }
      targetTab_ffi["highlighted"] = A.load.Bool(targetTab + 0);
      if (A.load.Bool(targetTab + 0)) {
        targetTab_ffi["id"] = A.load.Int0(targetTab + 0);
      }
      targetTab_ffi["incognito"] = A.load.Bool(targetTab + 0);
      targetTab_ffi["index"] = A.load.Int0(targetTab + 0);
      if (A.load.Bool(targetTab + 0 + 0)) {
        targetTab_ffi["mutedInfo"] = {};
        targetTab_ffi["mutedInfo"]["extensionId"] = A.load.Ref(targetTab + 0 + 0, undefined);
        targetTab_ffi["mutedInfo"]["muted"] = A.load.Bool(targetTab + 0 + 0);
        targetTab_ffi["mutedInfo"]["reason"] = A.load.Enum(targetTab + 0 + 0, ["user", "capture", "extension"]);
      }
      if (A.load.Bool(targetTab + 0)) {
        targetTab_ffi["openerTabId"] = A.load.Int0(targetTab + 0);
      }
      targetTab_ffi["pendingUrl"] = A.load.Ref(targetTab + 0, undefined);
      targetTab_ffi["pinned"] = A.load.Bool(targetTab + 0);
      targetTab_ffi["selected"] = A.load.Bool(targetTab + 0);
      targetTab_ffi["sessionId"] = A.load.Ref(targetTab + 0, undefined);
      targetTab_ffi["status"] = A.load.Enum(targetTab + 0, ["unloaded", "loading", "complete"]);
      targetTab_ffi["title"] = A.load.Ref(targetTab + 0, undefined);
      targetTab_ffi["url"] = A.load.Ref(targetTab + 0, undefined);
      if (A.load.Bool(targetTab + 0)) {
        targetTab_ffi["width"] = A.load.Int0(targetTab + 0);
      }
      targetTab_ffi["windowId"] = A.load.Int0(targetTab + 0);

      const options_ffi = {};

      options_ffi["selfBrowserSurface"] = A.load.Enum(options + 0, ["include", "exclude"]);
      if (A.load.Bool(options + 12)) {
        options_ffi["suppressLocalAudioPlaybackIntended"] = A.load.Bool(options + 4);
      }
      options_ffi["systemAudio"] = A.load.Enum(options + 8, ["include", "exclude"]);

      const _ret = WEBEXT.desktopCapture.chooseDesktopMedia(
        A.H.get<object>(sources),
        targetTab_ffi,
        options_ffi,
        A.H.get<object>(callback)
      );
      A.store.Int64(retPtr, _ret);
    },
    "try_ChooseDesktopMedia": (
      retPtr: Pointer,
      errPtr: Pointer,
      sources: heap.Ref<object>,
      targetTab: Pointer,
      options: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const targetTab_ffi = {};

        targetTab_ffi["active"] = A.load.Bool(targetTab + 0);
        if (A.load.Bool(targetTab + 0)) {
          targetTab_ffi["audible"] = A.load.Bool(targetTab + 0);
        }
        targetTab_ffi["autoDiscardable"] = A.load.Bool(targetTab + 0);
        targetTab_ffi["discarded"] = A.load.Bool(targetTab + 0);
        targetTab_ffi["favIconUrl"] = A.load.Ref(targetTab + 0, undefined);
        targetTab_ffi["groupId"] = A.load.Int0(targetTab + 0);
        if (A.load.Bool(targetTab + 0)) {
          targetTab_ffi["height"] = A.load.Int0(targetTab + 0);
        }
        targetTab_ffi["highlighted"] = A.load.Bool(targetTab + 0);
        if (A.load.Bool(targetTab + 0)) {
          targetTab_ffi["id"] = A.load.Int0(targetTab + 0);
        }
        targetTab_ffi["incognito"] = A.load.Bool(targetTab + 0);
        targetTab_ffi["index"] = A.load.Int0(targetTab + 0);
        if (A.load.Bool(targetTab + 0 + 0)) {
          targetTab_ffi["mutedInfo"] = {};
          targetTab_ffi["mutedInfo"]["extensionId"] = A.load.Ref(targetTab + 0 + 0, undefined);
          targetTab_ffi["mutedInfo"]["muted"] = A.load.Bool(targetTab + 0 + 0);
          targetTab_ffi["mutedInfo"]["reason"] = A.load.Enum(targetTab + 0 + 0, ["user", "capture", "extension"]);
        }
        if (A.load.Bool(targetTab + 0)) {
          targetTab_ffi["openerTabId"] = A.load.Int0(targetTab + 0);
        }
        targetTab_ffi["pendingUrl"] = A.load.Ref(targetTab + 0, undefined);
        targetTab_ffi["pinned"] = A.load.Bool(targetTab + 0);
        targetTab_ffi["selected"] = A.load.Bool(targetTab + 0);
        targetTab_ffi["sessionId"] = A.load.Ref(targetTab + 0, undefined);
        targetTab_ffi["status"] = A.load.Enum(targetTab + 0, ["unloaded", "loading", "complete"]);
        targetTab_ffi["title"] = A.load.Ref(targetTab + 0, undefined);
        targetTab_ffi["url"] = A.load.Ref(targetTab + 0, undefined);
        if (A.load.Bool(targetTab + 0)) {
          targetTab_ffi["width"] = A.load.Int0(targetTab + 0);
        }
        targetTab_ffi["windowId"] = A.load.Int0(targetTab + 0);

        const options_ffi = {};

        options_ffi["selfBrowserSurface"] = A.load.Enum(options + 0, ["include", "exclude"]);
        if (A.load.Bool(options + 12)) {
          options_ffi["suppressLocalAudioPlaybackIntended"] = A.load.Bool(options + 4);
        }
        options_ffi["systemAudio"] = A.load.Enum(options + 8, ["include", "exclude"]);

        const _ret = WEBEXT.desktopCapture.chooseDesktopMedia(
          A.H.get<object>(sources),
          targetTab_ffi,
          options_ffi,
          A.H.get<object>(callback)
        );
        A.store.Int64(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
