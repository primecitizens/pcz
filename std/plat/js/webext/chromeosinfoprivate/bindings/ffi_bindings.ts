import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/chromeosinfoprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_AssistantStatus": (ref: heap.Ref<string>): number => {
      const idx = ["unsupported", "supported"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_DeviceType": (ref: heap.Ref<string>): number => {
      const idx = ["chromebase", "chromebit", "chromebook", "chromebox", "chromedevice"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ManagedDeviceStatus": (ref: heap.Ref<string>): number => {
      const idx = ["managed", "not managed"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PlayStoreStatus": (ref: heap.Ref<string>): number => {
      const idx = ["not available", "available", "enabled"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SessionType": (ref: heap.Ref<string>): number => {
      const idx = ["normal", "kiosk", "public session"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_StylusStatus": (ref: heap.Ref<string>): number => {
      const idx = ["unsupported", "supported", "seen"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_GetReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 97, false);
        A.store.Bool(ptr + 80, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 81, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 82, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 83, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 84, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 85, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 86, false);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 87, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 88, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 89, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 90, false);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 91, false);
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 92, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 93, false);
        A.store.Bool(ptr + 13, false);
        A.store.Enum(ptr + 16, -1);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Enum(ptr + 32, -1);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Bool(ptr + 94, false);
        A.store.Bool(ptr + 48, false);
        A.store.Bool(ptr + 95, false);
        A.store.Bool(ptr + 49, false);
        A.store.Enum(ptr + 52, -1);
        A.store.Enum(ptr + 56, -1);
        A.store.Bool(ptr + 96, false);
        A.store.Bool(ptr + 60, false);
        A.store.Enum(ptr + 64, -1);
        A.store.Enum(ptr + 68, -1);
        A.store.Ref(ptr + 72, undefined);
        A.store.Ref(ptr + 76, undefined);
      } else {
        A.store.Bool(ptr + 97, true);
        A.store.Bool(ptr + 80, "a11yAutoClickEnabled" in x ? true : false);
        A.store.Bool(ptr + 0, x["a11yAutoClickEnabled"] ? true : false);
        A.store.Bool(ptr + 81, "a11yCaretHighlightEnabled" in x ? true : false);
        A.store.Bool(ptr + 1, x["a11yCaretHighlightEnabled"] ? true : false);
        A.store.Bool(ptr + 82, "a11yCursorColorEnabled" in x ? true : false);
        A.store.Bool(ptr + 2, x["a11yCursorColorEnabled"] ? true : false);
        A.store.Bool(ptr + 83, "a11yCursorHighlightEnabled" in x ? true : false);
        A.store.Bool(ptr + 3, x["a11yCursorHighlightEnabled"] ? true : false);
        A.store.Bool(ptr + 84, "a11yDockedMagnifierEnabled" in x ? true : false);
        A.store.Bool(ptr + 4, x["a11yDockedMagnifierEnabled"] ? true : false);
        A.store.Bool(ptr + 85, "a11yFocusHighlightEnabled" in x ? true : false);
        A.store.Bool(ptr + 5, x["a11yFocusHighlightEnabled"] ? true : false);
        A.store.Bool(ptr + 86, "a11yHighContrastEnabled" in x ? true : false);
        A.store.Bool(ptr + 6, x["a11yHighContrastEnabled"] ? true : false);
        A.store.Bool(ptr + 87, "a11yLargeCursorEnabled" in x ? true : false);
        A.store.Bool(ptr + 7, x["a11yLargeCursorEnabled"] ? true : false);
        A.store.Bool(ptr + 88, "a11yScreenMagnifierEnabled" in x ? true : false);
        A.store.Bool(ptr + 8, x["a11yScreenMagnifierEnabled"] ? true : false);
        A.store.Bool(ptr + 89, "a11ySelectToSpeakEnabled" in x ? true : false);
        A.store.Bool(ptr + 9, x["a11ySelectToSpeakEnabled"] ? true : false);
        A.store.Bool(ptr + 90, "a11ySpokenFeedbackEnabled" in x ? true : false);
        A.store.Bool(ptr + 10, x["a11ySpokenFeedbackEnabled"] ? true : false);
        A.store.Bool(ptr + 91, "a11yStickyKeysEnabled" in x ? true : false);
        A.store.Bool(ptr + 11, x["a11yStickyKeysEnabled"] ? true : false);
        A.store.Bool(ptr + 92, "a11ySwitchAccessEnabled" in x ? true : false);
        A.store.Bool(ptr + 12, x["a11ySwitchAccessEnabled"] ? true : false);
        A.store.Bool(ptr + 93, "a11yVirtualKeyboardEnabled" in x ? true : false);
        A.store.Bool(ptr + 13, x["a11yVirtualKeyboardEnabled"] ? true : false);
        A.store.Enum(ptr + 16, ["unsupported", "supported"].indexOf(x["assistantStatus"] as string));
        A.store.Ref(ptr + 20, x["board"]);
        A.store.Ref(ptr + 24, x["clientId"]);
        A.store.Ref(ptr + 28, x["customizationId"]);
        A.store.Enum(
          ptr + 32,
          ["chromebase", "chromebit", "chromebook", "chromebox", "chromedevice"].indexOf(x["deviceType"] as string)
        );
        A.store.Ref(ptr + 36, x["homeProvider"]);
        A.store.Ref(ptr + 40, x["hwid"]);
        A.store.Ref(ptr + 44, x["initialLocale"]);
        A.store.Bool(ptr + 94, "isMeetDevice" in x ? true : false);
        A.store.Bool(ptr + 48, x["isMeetDevice"] ? true : false);
        A.store.Bool(ptr + 95, "isOwner" in x ? true : false);
        A.store.Bool(ptr + 49, x["isOwner"] ? true : false);
        A.store.Enum(ptr + 52, ["managed", "not managed"].indexOf(x["managedDeviceStatus"] as string));
        A.store.Enum(ptr + 56, ["not available", "available", "enabled"].indexOf(x["playStoreStatus"] as string));
        A.store.Bool(ptr + 96, "sendFunctionKeys" in x ? true : false);
        A.store.Bool(ptr + 60, x["sendFunctionKeys"] ? true : false);
        A.store.Enum(ptr + 64, ["normal", "kiosk", "public session"].indexOf(x["sessionType"] as string));
        A.store.Enum(ptr + 68, ["unsupported", "supported", "seen"].indexOf(x["stylusStatus"] as string));
        A.store.Ref(ptr + 72, x["supportedTimezones"]);
        A.store.Ref(ptr + 76, x["timezone"]);
      }
    },
    "load_GetReturnType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 80)) {
        x["a11yAutoClickEnabled"] = A.load.Bool(ptr + 0);
      } else {
        delete x["a11yAutoClickEnabled"];
      }
      if (A.load.Bool(ptr + 81)) {
        x["a11yCaretHighlightEnabled"] = A.load.Bool(ptr + 1);
      } else {
        delete x["a11yCaretHighlightEnabled"];
      }
      if (A.load.Bool(ptr + 82)) {
        x["a11yCursorColorEnabled"] = A.load.Bool(ptr + 2);
      } else {
        delete x["a11yCursorColorEnabled"];
      }
      if (A.load.Bool(ptr + 83)) {
        x["a11yCursorHighlightEnabled"] = A.load.Bool(ptr + 3);
      } else {
        delete x["a11yCursorHighlightEnabled"];
      }
      if (A.load.Bool(ptr + 84)) {
        x["a11yDockedMagnifierEnabled"] = A.load.Bool(ptr + 4);
      } else {
        delete x["a11yDockedMagnifierEnabled"];
      }
      if (A.load.Bool(ptr + 85)) {
        x["a11yFocusHighlightEnabled"] = A.load.Bool(ptr + 5);
      } else {
        delete x["a11yFocusHighlightEnabled"];
      }
      if (A.load.Bool(ptr + 86)) {
        x["a11yHighContrastEnabled"] = A.load.Bool(ptr + 6);
      } else {
        delete x["a11yHighContrastEnabled"];
      }
      if (A.load.Bool(ptr + 87)) {
        x["a11yLargeCursorEnabled"] = A.load.Bool(ptr + 7);
      } else {
        delete x["a11yLargeCursorEnabled"];
      }
      if (A.load.Bool(ptr + 88)) {
        x["a11yScreenMagnifierEnabled"] = A.load.Bool(ptr + 8);
      } else {
        delete x["a11yScreenMagnifierEnabled"];
      }
      if (A.load.Bool(ptr + 89)) {
        x["a11ySelectToSpeakEnabled"] = A.load.Bool(ptr + 9);
      } else {
        delete x["a11ySelectToSpeakEnabled"];
      }
      if (A.load.Bool(ptr + 90)) {
        x["a11ySpokenFeedbackEnabled"] = A.load.Bool(ptr + 10);
      } else {
        delete x["a11ySpokenFeedbackEnabled"];
      }
      if (A.load.Bool(ptr + 91)) {
        x["a11yStickyKeysEnabled"] = A.load.Bool(ptr + 11);
      } else {
        delete x["a11yStickyKeysEnabled"];
      }
      if (A.load.Bool(ptr + 92)) {
        x["a11ySwitchAccessEnabled"] = A.load.Bool(ptr + 12);
      } else {
        delete x["a11ySwitchAccessEnabled"];
      }
      if (A.load.Bool(ptr + 93)) {
        x["a11yVirtualKeyboardEnabled"] = A.load.Bool(ptr + 13);
      } else {
        delete x["a11yVirtualKeyboardEnabled"];
      }
      x["assistantStatus"] = A.load.Enum(ptr + 16, ["unsupported", "supported"]);
      x["board"] = A.load.Ref(ptr + 20, undefined);
      x["clientId"] = A.load.Ref(ptr + 24, undefined);
      x["customizationId"] = A.load.Ref(ptr + 28, undefined);
      x["deviceType"] = A.load.Enum(ptr + 32, ["chromebase", "chromebit", "chromebook", "chromebox", "chromedevice"]);
      x["homeProvider"] = A.load.Ref(ptr + 36, undefined);
      x["hwid"] = A.load.Ref(ptr + 40, undefined);
      x["initialLocale"] = A.load.Ref(ptr + 44, undefined);
      if (A.load.Bool(ptr + 94)) {
        x["isMeetDevice"] = A.load.Bool(ptr + 48);
      } else {
        delete x["isMeetDevice"];
      }
      if (A.load.Bool(ptr + 95)) {
        x["isOwner"] = A.load.Bool(ptr + 49);
      } else {
        delete x["isOwner"];
      }
      x["managedDeviceStatus"] = A.load.Enum(ptr + 52, ["managed", "not managed"]);
      x["playStoreStatus"] = A.load.Enum(ptr + 56, ["not available", "available", "enabled"]);
      if (A.load.Bool(ptr + 96)) {
        x["sendFunctionKeys"] = A.load.Bool(ptr + 60);
      } else {
        delete x["sendFunctionKeys"];
      }
      x["sessionType"] = A.load.Enum(ptr + 64, ["normal", "kiosk", "public session"]);
      x["stylusStatus"] = A.load.Enum(ptr + 68, ["unsupported", "supported", "seen"]);
      x["supportedTimezones"] = A.load.Ref(ptr + 72, undefined);
      x["timezone"] = A.load.Ref(ptr + 76, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PropertyName": (ref: heap.Ref<string>): number => {
      const idx = [
        "timezone",
        "a11yLargeCursorEnabled",
        "a11yStickyKeysEnabled",
        "a11ySpokenFeedbackEnabled",
        "a11yHighContrastEnabled",
        "a11yScreenMagnifierEnabled",
        "a11yAutoClickEnabled",
        "a11yVirtualKeyboardEnabled",
        "a11yCaretHighlightEnabled",
        "a11yCursorHighlightEnabled",
        "a11yFocusHighlightEnabled",
        "a11ySelectToSpeakEnabled",
        "a11ySwitchAccessEnabled",
        "a11yCursorColorEnabled",
        "a11yDockedMagnifierEnabled",
        "sendFunctionKeys",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_Get": (): heap.Ref<boolean> => {
      if (WEBEXT?.chromeosInfoPrivate && "get" in WEBEXT?.chromeosInfoPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Get": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.chromeosInfoPrivate.get);
    },
    "call_Get": (retPtr: Pointer, propertyNames: heap.Ref<object>): void => {
      const _ret = WEBEXT.chromeosInfoPrivate.get(A.H.get<object>(propertyNames));
      A.store.Ref(retPtr, _ret);
    },
    "try_Get": (retPtr: Pointer, errPtr: Pointer, propertyNames: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.chromeosInfoPrivate.get(A.H.get<object>(propertyNames));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsTabletModeEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.chromeosInfoPrivate && "isTabletModeEnabled" in WEBEXT?.chromeosInfoPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsTabletModeEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.chromeosInfoPrivate.isTabletModeEnabled);
    },
    "call_IsTabletModeEnabled": (retPtr: Pointer): void => {
      const _ret = WEBEXT.chromeosInfoPrivate.isTabletModeEnabled();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsTabletModeEnabled": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.chromeosInfoPrivate.isTabletModeEnabled();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Set": (): heap.Ref<boolean> => {
      if (WEBEXT?.chromeosInfoPrivate && "set" in WEBEXT?.chromeosInfoPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Set": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.chromeosInfoPrivate.set);
    },
    "call_Set": (retPtr: Pointer, propertyName: number, propertyValue: heap.Ref<object>): void => {
      const _ret = WEBEXT.chromeosInfoPrivate.set(
        propertyName > 0 && propertyName <= 16
          ? [
              "timezone",
              "a11yLargeCursorEnabled",
              "a11yStickyKeysEnabled",
              "a11ySpokenFeedbackEnabled",
              "a11yHighContrastEnabled",
              "a11yScreenMagnifierEnabled",
              "a11yAutoClickEnabled",
              "a11yVirtualKeyboardEnabled",
              "a11yCaretHighlightEnabled",
              "a11yCursorHighlightEnabled",
              "a11yFocusHighlightEnabled",
              "a11ySelectToSpeakEnabled",
              "a11ySwitchAccessEnabled",
              "a11yCursorColorEnabled",
              "a11yDockedMagnifierEnabled",
              "sendFunctionKeys",
            ][propertyName - 1]
          : undefined,
        A.H.get<object>(propertyValue)
      );
    },
    "try_Set": (
      retPtr: Pointer,
      errPtr: Pointer,
      propertyName: number,
      propertyValue: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.chromeosInfoPrivate.set(
          propertyName > 0 && propertyName <= 16
            ? [
                "timezone",
                "a11yLargeCursorEnabled",
                "a11yStickyKeysEnabled",
                "a11ySpokenFeedbackEnabled",
                "a11yHighContrastEnabled",
                "a11yScreenMagnifierEnabled",
                "a11yAutoClickEnabled",
                "a11yVirtualKeyboardEnabled",
                "a11yCaretHighlightEnabled",
                "a11yCursorHighlightEnabled",
                "a11yFocusHighlightEnabled",
                "a11ySelectToSpeakEnabled",
                "a11ySwitchAccessEnabled",
                "a11yCursorColorEnabled",
                "a11yDockedMagnifierEnabled",
                "sendFunctionKeys",
              ][propertyName - 1]
            : undefined,
          A.H.get<object>(propertyValue)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
