import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/chromeos/events", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_AudioJackDeviceType": (ref: heap.Ref<string>): number => {
      const idx = ["headphone", "microphone"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_AudioJackEvent": (ref: heap.Ref<string>): number => {
      const idx = ["connected", "disconnected"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_AudioJackEventInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["connected", "disconnected"].indexOf(x["event"] as string));
        A.store.Enum(ptr + 4, ["headphone", "microphone"].indexOf(x["deviceType"] as string));
      }
    },
    "load_AudioJackEventInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["event"] = A.load.Enum(ptr + 0, ["connected", "disconnected"]);
      x["deviceType"] = A.load.Enum(ptr + 4, ["headphone", "microphone"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DisplayInputType": (ref: heap.Ref<string>): number => {
      const idx = ["unknown", "digital", "analog"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_EventCategory": (ref: heap.Ref<string>): number => {
      const idx = [
        "audio_jack",
        "lid",
        "usb",
        "sd_card",
        "power",
        "keyboard_diagnostic",
        "stylus_garage",
        "touchpad_button",
        "touchpad_touch",
        "touchpad_connected",
        "touchscreen_touch",
        "touchscreen_connected",
        "external_display",
        "stylus_touch",
        "stylus_connected",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_EventSupportStatus": (ref: heap.Ref<string>): number => {
      const idx = ["supported", "unsupported"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_EventSupportStatusInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["supported", "unsupported"].indexOf(x["status"] as string));
      }
    },
    "load_EventSupportStatusInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["status"] = A.load.Enum(ptr + 0, ["supported", "unsupported"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ExternalDisplayEvent": (ref: heap.Ref<string>): number => {
      const idx = ["connected", "disconnected"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ExternalDisplayInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 65, false);
        A.store.Bool(ptr + 56, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 57, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 58, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 59, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 60, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Ref(ptr + 24, undefined);
        A.store.Bool(ptr + 61, false);
        A.store.Int32(ptr + 28, 0);
        A.store.Bool(ptr + 62, false);
        A.store.Int32(ptr + 32, 0);
        A.store.Bool(ptr + 63, false);
        A.store.Int32(ptr + 36, 0);
        A.store.Bool(ptr + 64, false);
        A.store.Int32(ptr + 40, 0);
        A.store.Ref(ptr + 44, undefined);
        A.store.Enum(ptr + 48, -1);
        A.store.Ref(ptr + 52, undefined);
      } else {
        A.store.Bool(ptr + 65, true);
        A.store.Bool(ptr + 56, "displayWidth" in x ? true : false);
        A.store.Int32(ptr + 0, x["displayWidth"] === undefined ? 0 : (x["displayWidth"] as number));
        A.store.Bool(ptr + 57, "displayHeight" in x ? true : false);
        A.store.Int32(ptr + 4, x["displayHeight"] === undefined ? 0 : (x["displayHeight"] as number));
        A.store.Bool(ptr + 58, "resolutionHorizontal" in x ? true : false);
        A.store.Int32(ptr + 8, x["resolutionHorizontal"] === undefined ? 0 : (x["resolutionHorizontal"] as number));
        A.store.Bool(ptr + 59, "resolutionVertical" in x ? true : false);
        A.store.Int32(ptr + 12, x["resolutionVertical"] === undefined ? 0 : (x["resolutionVertical"] as number));
        A.store.Bool(ptr + 60, "refreshRate" in x ? true : false);
        A.store.Float64(ptr + 16, x["refreshRate"] === undefined ? 0 : (x["refreshRate"] as number));
        A.store.Ref(ptr + 24, x["manufacturer"]);
        A.store.Bool(ptr + 61, "modelId" in x ? true : false);
        A.store.Int32(ptr + 28, x["modelId"] === undefined ? 0 : (x["modelId"] as number));
        A.store.Bool(ptr + 62, "serialNumber" in x ? true : false);
        A.store.Int32(ptr + 32, x["serialNumber"] === undefined ? 0 : (x["serialNumber"] as number));
        A.store.Bool(ptr + 63, "manufactureWeek" in x ? true : false);
        A.store.Int32(ptr + 36, x["manufactureWeek"] === undefined ? 0 : (x["manufactureWeek"] as number));
        A.store.Bool(ptr + 64, "manufactureYear" in x ? true : false);
        A.store.Int32(ptr + 40, x["manufactureYear"] === undefined ? 0 : (x["manufactureYear"] as number));
        A.store.Ref(ptr + 44, x["edidVersion"]);
        A.store.Enum(ptr + 48, ["unknown", "digital", "analog"].indexOf(x["inputType"] as string));
        A.store.Ref(ptr + 52, x["displayName"]);
      }
    },
    "load_ExternalDisplayInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 56)) {
        x["displayWidth"] = A.load.Int32(ptr + 0);
      } else {
        delete x["displayWidth"];
      }
      if (A.load.Bool(ptr + 57)) {
        x["displayHeight"] = A.load.Int32(ptr + 4);
      } else {
        delete x["displayHeight"];
      }
      if (A.load.Bool(ptr + 58)) {
        x["resolutionHorizontal"] = A.load.Int32(ptr + 8);
      } else {
        delete x["resolutionHorizontal"];
      }
      if (A.load.Bool(ptr + 59)) {
        x["resolutionVertical"] = A.load.Int32(ptr + 12);
      } else {
        delete x["resolutionVertical"];
      }
      if (A.load.Bool(ptr + 60)) {
        x["refreshRate"] = A.load.Float64(ptr + 16);
      } else {
        delete x["refreshRate"];
      }
      x["manufacturer"] = A.load.Ref(ptr + 24, undefined);
      if (A.load.Bool(ptr + 61)) {
        x["modelId"] = A.load.Int32(ptr + 28);
      } else {
        delete x["modelId"];
      }
      if (A.load.Bool(ptr + 62)) {
        x["serialNumber"] = A.load.Int32(ptr + 32);
      } else {
        delete x["serialNumber"];
      }
      if (A.load.Bool(ptr + 63)) {
        x["manufactureWeek"] = A.load.Int32(ptr + 36);
      } else {
        delete x["manufactureWeek"];
      }
      if (A.load.Bool(ptr + 64)) {
        x["manufactureYear"] = A.load.Int32(ptr + 40);
      } else {
        delete x["manufactureYear"];
      }
      x["edidVersion"] = A.load.Ref(ptr + 44, undefined);
      x["inputType"] = A.load.Enum(ptr + 48, ["unknown", "digital", "analog"]);
      x["displayName"] = A.load.Ref(ptr + 52, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ExternalDisplayEventInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 74, false);
        A.store.Enum(ptr + 0, -1);

        A.store.Bool(ptr + 8 + 65, false);
        A.store.Bool(ptr + 8 + 56, false);
        A.store.Int32(ptr + 8 + 0, 0);
        A.store.Bool(ptr + 8 + 57, false);
        A.store.Int32(ptr + 8 + 4, 0);
        A.store.Bool(ptr + 8 + 58, false);
        A.store.Int32(ptr + 8 + 8, 0);
        A.store.Bool(ptr + 8 + 59, false);
        A.store.Int32(ptr + 8 + 12, 0);
        A.store.Bool(ptr + 8 + 60, false);
        A.store.Float64(ptr + 8 + 16, 0);
        A.store.Ref(ptr + 8 + 24, undefined);
        A.store.Bool(ptr + 8 + 61, false);
        A.store.Int32(ptr + 8 + 28, 0);
        A.store.Bool(ptr + 8 + 62, false);
        A.store.Int32(ptr + 8 + 32, 0);
        A.store.Bool(ptr + 8 + 63, false);
        A.store.Int32(ptr + 8 + 36, 0);
        A.store.Bool(ptr + 8 + 64, false);
        A.store.Int32(ptr + 8 + 40, 0);
        A.store.Ref(ptr + 8 + 44, undefined);
        A.store.Enum(ptr + 8 + 48, -1);
        A.store.Ref(ptr + 8 + 52, undefined);
      } else {
        A.store.Bool(ptr + 74, true);
        A.store.Enum(ptr + 0, ["connected", "disconnected"].indexOf(x["event"] as string));

        if (typeof x["displayInfo"] === "undefined") {
          A.store.Bool(ptr + 8 + 65, false);
          A.store.Bool(ptr + 8 + 56, false);
          A.store.Int32(ptr + 8 + 0, 0);
          A.store.Bool(ptr + 8 + 57, false);
          A.store.Int32(ptr + 8 + 4, 0);
          A.store.Bool(ptr + 8 + 58, false);
          A.store.Int32(ptr + 8 + 8, 0);
          A.store.Bool(ptr + 8 + 59, false);
          A.store.Int32(ptr + 8 + 12, 0);
          A.store.Bool(ptr + 8 + 60, false);
          A.store.Float64(ptr + 8 + 16, 0);
          A.store.Ref(ptr + 8 + 24, undefined);
          A.store.Bool(ptr + 8 + 61, false);
          A.store.Int32(ptr + 8 + 28, 0);
          A.store.Bool(ptr + 8 + 62, false);
          A.store.Int32(ptr + 8 + 32, 0);
          A.store.Bool(ptr + 8 + 63, false);
          A.store.Int32(ptr + 8 + 36, 0);
          A.store.Bool(ptr + 8 + 64, false);
          A.store.Int32(ptr + 8 + 40, 0);
          A.store.Ref(ptr + 8 + 44, undefined);
          A.store.Enum(ptr + 8 + 48, -1);
          A.store.Ref(ptr + 8 + 52, undefined);
        } else {
          A.store.Bool(ptr + 8 + 65, true);
          A.store.Bool(ptr + 8 + 56, "displayWidth" in x["displayInfo"] ? true : false);
          A.store.Int32(
            ptr + 8 + 0,
            x["displayInfo"]["displayWidth"] === undefined ? 0 : (x["displayInfo"]["displayWidth"] as number)
          );
          A.store.Bool(ptr + 8 + 57, "displayHeight" in x["displayInfo"] ? true : false);
          A.store.Int32(
            ptr + 8 + 4,
            x["displayInfo"]["displayHeight"] === undefined ? 0 : (x["displayInfo"]["displayHeight"] as number)
          );
          A.store.Bool(ptr + 8 + 58, "resolutionHorizontal" in x["displayInfo"] ? true : false);
          A.store.Int32(
            ptr + 8 + 8,
            x["displayInfo"]["resolutionHorizontal"] === undefined
              ? 0
              : (x["displayInfo"]["resolutionHorizontal"] as number)
          );
          A.store.Bool(ptr + 8 + 59, "resolutionVertical" in x["displayInfo"] ? true : false);
          A.store.Int32(
            ptr + 8 + 12,
            x["displayInfo"]["resolutionVertical"] === undefined
              ? 0
              : (x["displayInfo"]["resolutionVertical"] as number)
          );
          A.store.Bool(ptr + 8 + 60, "refreshRate" in x["displayInfo"] ? true : false);
          A.store.Float64(
            ptr + 8 + 16,
            x["displayInfo"]["refreshRate"] === undefined ? 0 : (x["displayInfo"]["refreshRate"] as number)
          );
          A.store.Ref(ptr + 8 + 24, x["displayInfo"]["manufacturer"]);
          A.store.Bool(ptr + 8 + 61, "modelId" in x["displayInfo"] ? true : false);
          A.store.Int32(
            ptr + 8 + 28,
            x["displayInfo"]["modelId"] === undefined ? 0 : (x["displayInfo"]["modelId"] as number)
          );
          A.store.Bool(ptr + 8 + 62, "serialNumber" in x["displayInfo"] ? true : false);
          A.store.Int32(
            ptr + 8 + 32,
            x["displayInfo"]["serialNumber"] === undefined ? 0 : (x["displayInfo"]["serialNumber"] as number)
          );
          A.store.Bool(ptr + 8 + 63, "manufactureWeek" in x["displayInfo"] ? true : false);
          A.store.Int32(
            ptr + 8 + 36,
            x["displayInfo"]["manufactureWeek"] === undefined ? 0 : (x["displayInfo"]["manufactureWeek"] as number)
          );
          A.store.Bool(ptr + 8 + 64, "manufactureYear" in x["displayInfo"] ? true : false);
          A.store.Int32(
            ptr + 8 + 40,
            x["displayInfo"]["manufactureYear"] === undefined ? 0 : (x["displayInfo"]["manufactureYear"] as number)
          );
          A.store.Ref(ptr + 8 + 44, x["displayInfo"]["edidVersion"]);
          A.store.Enum(ptr + 8 + 48, ["unknown", "digital", "analog"].indexOf(x["displayInfo"]["inputType"] as string));
          A.store.Ref(ptr + 8 + 52, x["displayInfo"]["displayName"]);
        }
      }
    },
    "load_ExternalDisplayEventInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["event"] = A.load.Enum(ptr + 0, ["connected", "disconnected"]);
      if (A.load.Bool(ptr + 8 + 65)) {
        x["displayInfo"] = {};
        if (A.load.Bool(ptr + 8 + 56)) {
          x["displayInfo"]["displayWidth"] = A.load.Int32(ptr + 8 + 0);
        } else {
          delete x["displayInfo"]["displayWidth"];
        }
        if (A.load.Bool(ptr + 8 + 57)) {
          x["displayInfo"]["displayHeight"] = A.load.Int32(ptr + 8 + 4);
        } else {
          delete x["displayInfo"]["displayHeight"];
        }
        if (A.load.Bool(ptr + 8 + 58)) {
          x["displayInfo"]["resolutionHorizontal"] = A.load.Int32(ptr + 8 + 8);
        } else {
          delete x["displayInfo"]["resolutionHorizontal"];
        }
        if (A.load.Bool(ptr + 8 + 59)) {
          x["displayInfo"]["resolutionVertical"] = A.load.Int32(ptr + 8 + 12);
        } else {
          delete x["displayInfo"]["resolutionVertical"];
        }
        if (A.load.Bool(ptr + 8 + 60)) {
          x["displayInfo"]["refreshRate"] = A.load.Float64(ptr + 8 + 16);
        } else {
          delete x["displayInfo"]["refreshRate"];
        }
        x["displayInfo"]["manufacturer"] = A.load.Ref(ptr + 8 + 24, undefined);
        if (A.load.Bool(ptr + 8 + 61)) {
          x["displayInfo"]["modelId"] = A.load.Int32(ptr + 8 + 28);
        } else {
          delete x["displayInfo"]["modelId"];
        }
        if (A.load.Bool(ptr + 8 + 62)) {
          x["displayInfo"]["serialNumber"] = A.load.Int32(ptr + 8 + 32);
        } else {
          delete x["displayInfo"]["serialNumber"];
        }
        if (A.load.Bool(ptr + 8 + 63)) {
          x["displayInfo"]["manufactureWeek"] = A.load.Int32(ptr + 8 + 36);
        } else {
          delete x["displayInfo"]["manufactureWeek"];
        }
        if (A.load.Bool(ptr + 8 + 64)) {
          x["displayInfo"]["manufactureYear"] = A.load.Int32(ptr + 8 + 40);
        } else {
          delete x["displayInfo"]["manufactureYear"];
        }
        x["displayInfo"]["edidVersion"] = A.load.Ref(ptr + 8 + 44, undefined);
        x["displayInfo"]["inputType"] = A.load.Enum(ptr + 8 + 48, ["unknown", "digital", "analog"]);
        x["displayInfo"]["displayName"] = A.load.Ref(ptr + 8 + 52, undefined);
      } else {
        delete x["displayInfo"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_InputTouchButton": (ref: heap.Ref<string>): number => {
      const idx = ["left", "middle", "right"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_InputTouchButtonState": (ref: heap.Ref<string>): number => {
      const idx = ["pressed", "released"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_KeyboardConnectionType": (ref: heap.Ref<string>): number => {
      const idx = ["internal", "usb", "bluetooth", "unknown"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PhysicalKeyboardLayout": (ref: heap.Ref<string>): number => {
      const idx = ["unknown", "chrome_os"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_MechanicalKeyboardLayout": (ref: heap.Ref<string>): number => {
      const idx = ["unknown", "ansi", "iso", "jis"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_KeyboardNumberPadPresence": (ref: heap.Ref<string>): number => {
      const idx = ["unknown", "present", "not_present"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_KeyboardTopRowKey": (ref: heap.Ref<string>): number => {
      const idx = [
        "no_key",
        "unknown",
        "back",
        "forward",
        "refresh",
        "fullscreen",
        "overview",
        "screenshot",
        "screen_brightness_down",
        "screen_brightness_up",
        "privacy_screen_toggle",
        "microphone_mute",
        "volume_mute",
        "volume_down",
        "volume_up",
        "keyboard_backlight_toggle",
        "keyboard_backlight_down",
        "keyboard_backlight_up",
        "next_track",
        "previous_track",
        "play_pause",
        "screen_mirror",
        "delete",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_KeyboardTopRightKey": (ref: heap.Ref<string>): number => {
      const idx = ["unknown", "power", "lock", "control_panel"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_KeyboardInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 39, false);
        A.store.Bool(ptr + 37, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
        A.store.Enum(ptr + 12, -1);
        A.store.Enum(ptr + 16, -1);
        A.store.Ref(ptr + 20, undefined);
        A.store.Enum(ptr + 24, -1);
        A.store.Ref(ptr + 28, undefined);
        A.store.Enum(ptr + 32, -1);
        A.store.Bool(ptr + 38, false);
        A.store.Bool(ptr + 36, false);
      } else {
        A.store.Bool(ptr + 39, true);
        A.store.Bool(ptr + 37, "id" in x ? true : false);
        A.store.Int32(ptr + 0, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Enum(ptr + 4, ["internal", "usb", "bluetooth", "unknown"].indexOf(x["connectionType"] as string));
        A.store.Ref(ptr + 8, x["name"]);
        A.store.Enum(ptr + 12, ["unknown", "chrome_os"].indexOf(x["physicalLayout"] as string));
        A.store.Enum(ptr + 16, ["unknown", "ansi", "iso", "jis"].indexOf(x["mechanicalLayout"] as string));
        A.store.Ref(ptr + 20, x["regionCode"]);
        A.store.Enum(ptr + 24, ["unknown", "present", "not_present"].indexOf(x["numberPadPresent"] as string));
        A.store.Ref(ptr + 28, x["topRowKeys"]);
        A.store.Enum(ptr + 32, ["unknown", "power", "lock", "control_panel"].indexOf(x["topRightKey"] as string));
        A.store.Bool(ptr + 38, "hasAssistantKey" in x ? true : false);
        A.store.Bool(ptr + 36, x["hasAssistantKey"] ? true : false);
      }
    },
    "load_KeyboardInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 37)) {
        x["id"] = A.load.Int32(ptr + 0);
      } else {
        delete x["id"];
      }
      x["connectionType"] = A.load.Enum(ptr + 4, ["internal", "usb", "bluetooth", "unknown"]);
      x["name"] = A.load.Ref(ptr + 8, undefined);
      x["physicalLayout"] = A.load.Enum(ptr + 12, ["unknown", "chrome_os"]);
      x["mechanicalLayout"] = A.load.Enum(ptr + 16, ["unknown", "ansi", "iso", "jis"]);
      x["regionCode"] = A.load.Ref(ptr + 20, undefined);
      x["numberPadPresent"] = A.load.Enum(ptr + 24, ["unknown", "present", "not_present"]);
      x["topRowKeys"] = A.load.Ref(ptr + 28, undefined);
      x["topRightKey"] = A.load.Enum(ptr + 32, ["unknown", "power", "lock", "control_panel"]);
      if (A.load.Bool(ptr + 38)) {
        x["hasAssistantKey"] = A.load.Bool(ptr + 36);
      } else {
        delete x["hasAssistantKey"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_KeyboardDiagnosticEventInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 48, false);

        A.store.Bool(ptr + 0 + 39, false);
        A.store.Bool(ptr + 0 + 37, false);
        A.store.Int32(ptr + 0 + 0, 0);
        A.store.Enum(ptr + 0 + 4, -1);
        A.store.Ref(ptr + 0 + 8, undefined);
        A.store.Enum(ptr + 0 + 12, -1);
        A.store.Enum(ptr + 0 + 16, -1);
        A.store.Ref(ptr + 0 + 20, undefined);
        A.store.Enum(ptr + 0 + 24, -1);
        A.store.Ref(ptr + 0 + 28, undefined);
        A.store.Enum(ptr + 0 + 32, -1);
        A.store.Bool(ptr + 0 + 38, false);
        A.store.Bool(ptr + 0 + 36, false);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
      } else {
        A.store.Bool(ptr + 48, true);

        if (typeof x["keyboardInfo"] === "undefined") {
          A.store.Bool(ptr + 0 + 39, false);
          A.store.Bool(ptr + 0 + 37, false);
          A.store.Int32(ptr + 0 + 0, 0);
          A.store.Enum(ptr + 0 + 4, -1);
          A.store.Ref(ptr + 0 + 8, undefined);
          A.store.Enum(ptr + 0 + 12, -1);
          A.store.Enum(ptr + 0 + 16, -1);
          A.store.Ref(ptr + 0 + 20, undefined);
          A.store.Enum(ptr + 0 + 24, -1);
          A.store.Ref(ptr + 0 + 28, undefined);
          A.store.Enum(ptr + 0 + 32, -1);
          A.store.Bool(ptr + 0 + 38, false);
          A.store.Bool(ptr + 0 + 36, false);
        } else {
          A.store.Bool(ptr + 0 + 39, true);
          A.store.Bool(ptr + 0 + 37, "id" in x["keyboardInfo"] ? true : false);
          A.store.Int32(ptr + 0 + 0, x["keyboardInfo"]["id"] === undefined ? 0 : (x["keyboardInfo"]["id"] as number));
          A.store.Enum(
            ptr + 0 + 4,
            ["internal", "usb", "bluetooth", "unknown"].indexOf(x["keyboardInfo"]["connectionType"] as string)
          );
          A.store.Ref(ptr + 0 + 8, x["keyboardInfo"]["name"]);
          A.store.Enum(ptr + 0 + 12, ["unknown", "chrome_os"].indexOf(x["keyboardInfo"]["physicalLayout"] as string));
          A.store.Enum(
            ptr + 0 + 16,
            ["unknown", "ansi", "iso", "jis"].indexOf(x["keyboardInfo"]["mechanicalLayout"] as string)
          );
          A.store.Ref(ptr + 0 + 20, x["keyboardInfo"]["regionCode"]);
          A.store.Enum(
            ptr + 0 + 24,
            ["unknown", "present", "not_present"].indexOf(x["keyboardInfo"]["numberPadPresent"] as string)
          );
          A.store.Ref(ptr + 0 + 28, x["keyboardInfo"]["topRowKeys"]);
          A.store.Enum(
            ptr + 0 + 32,
            ["unknown", "power", "lock", "control_panel"].indexOf(x["keyboardInfo"]["topRightKey"] as string)
          );
          A.store.Bool(ptr + 0 + 38, "hasAssistantKey" in x["keyboardInfo"] ? true : false);
          A.store.Bool(ptr + 0 + 36, x["keyboardInfo"]["hasAssistantKey"] ? true : false);
        }
        A.store.Ref(ptr + 40, x["testedKeys"]);
        A.store.Ref(ptr + 44, x["testedTopRowKeys"]);
      }
    },
    "load_KeyboardDiagnosticEventInfo": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 39)) {
        x["keyboardInfo"] = {};
        if (A.load.Bool(ptr + 0 + 37)) {
          x["keyboardInfo"]["id"] = A.load.Int32(ptr + 0 + 0);
        } else {
          delete x["keyboardInfo"]["id"];
        }
        x["keyboardInfo"]["connectionType"] = A.load.Enum(ptr + 0 + 4, ["internal", "usb", "bluetooth", "unknown"]);
        x["keyboardInfo"]["name"] = A.load.Ref(ptr + 0 + 8, undefined);
        x["keyboardInfo"]["physicalLayout"] = A.load.Enum(ptr + 0 + 12, ["unknown", "chrome_os"]);
        x["keyboardInfo"]["mechanicalLayout"] = A.load.Enum(ptr + 0 + 16, ["unknown", "ansi", "iso", "jis"]);
        x["keyboardInfo"]["regionCode"] = A.load.Ref(ptr + 0 + 20, undefined);
        x["keyboardInfo"]["numberPadPresent"] = A.load.Enum(ptr + 0 + 24, ["unknown", "present", "not_present"]);
        x["keyboardInfo"]["topRowKeys"] = A.load.Ref(ptr + 0 + 28, undefined);
        x["keyboardInfo"]["topRightKey"] = A.load.Enum(ptr + 0 + 32, ["unknown", "power", "lock", "control_panel"]);
        if (A.load.Bool(ptr + 0 + 38)) {
          x["keyboardInfo"]["hasAssistantKey"] = A.load.Bool(ptr + 0 + 36);
        } else {
          delete x["keyboardInfo"]["hasAssistantKey"];
        }
      } else {
        delete x["keyboardInfo"];
      }
      x["testedKeys"] = A.load.Ref(ptr + 40, undefined);
      x["testedTopRowKeys"] = A.load.Ref(ptr + 44, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_LidEvent": (ref: heap.Ref<string>): number => {
      const idx = ["closed", "opened"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_LidEventInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["closed", "opened"].indexOf(x["event"] as string));
      }
    },
    "load_LidEventInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["event"] = A.load.Enum(ptr + 0, ["closed", "opened"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PowerEvent": (ref: heap.Ref<string>): number => {
      const idx = ["ac_inserted", "ac_removed", "os_suspend", "os_resume"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PowerEventInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["ac_inserted", "ac_removed", "os_suspend", "os_resume"].indexOf(x["event"] as string));
      }
    },
    "load_PowerEventInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["event"] = A.load.Enum(ptr + 0, ["ac_inserted", "ac_removed", "os_suspend", "os_resume"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_SdCardEvent": (ref: heap.Ref<string>): number => {
      const idx = ["connected", "disconnected"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SdCardEventInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["connected", "disconnected"].indexOf(x["event"] as string));
      }
    },
    "load_SdCardEventInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["event"] = A.load.Enum(ptr + 0, ["connected", "disconnected"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_StylusConnectedEventInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 14, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 15, true);
        A.store.Bool(ptr + 12, "max_x" in x ? true : false);
        A.store.Int32(ptr + 0, x["max_x"] === undefined ? 0 : (x["max_x"] as number));
        A.store.Bool(ptr + 13, "max_y" in x ? true : false);
        A.store.Int32(ptr + 4, x["max_y"] === undefined ? 0 : (x["max_y"] as number));
        A.store.Bool(ptr + 14, "max_pressure" in x ? true : false);
        A.store.Int32(ptr + 8, x["max_pressure"] === undefined ? 0 : (x["max_pressure"] as number));
      }
    },
    "load_StylusConnectedEventInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["max_x"] = A.load.Int32(ptr + 0);
      } else {
        delete x["max_x"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["max_y"] = A.load.Int32(ptr + 4);
      } else {
        delete x["max_y"];
      }
      if (A.load.Bool(ptr + 14)) {
        x["max_pressure"] = A.load.Int32(ptr + 8);
      } else {
        delete x["max_pressure"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_StylusGarageEvent": (ref: heap.Ref<string>): number => {
      const idx = ["inserted", "removed"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_StylusGarageEventInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["inserted", "removed"].indexOf(x["event"] as string));
      }
    },
    "load_StylusGarageEventInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["event"] = A.load.Enum(ptr + 0, ["inserted", "removed"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_StylusTouchPointInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 14, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 15, true);
        A.store.Bool(ptr + 12, "x" in x ? true : false);
        A.store.Int32(ptr + 0, x["x"] === undefined ? 0 : (x["x"] as number));
        A.store.Bool(ptr + 13, "y" in x ? true : false);
        A.store.Int32(ptr + 4, x["y"] === undefined ? 0 : (x["y"] as number));
        A.store.Bool(ptr + 14, "pressure" in x ? true : false);
        A.store.Int32(ptr + 8, x["pressure"] === undefined ? 0 : (x["pressure"] as number));
      }
    },
    "load_StylusTouchPointInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["x"] = A.load.Int32(ptr + 0);
      } else {
        delete x["x"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["y"] = A.load.Int32(ptr + 4);
      } else {
        delete x["y"];
      }
      if (A.load.Bool(ptr + 14)) {
        x["pressure"] = A.load.Int32(ptr + 8);
      } else {
        delete x["pressure"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_StylusTouchEventInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);

        A.store.Bool(ptr + 0 + 15, false);
        A.store.Bool(ptr + 0 + 12, false);
        A.store.Int32(ptr + 0 + 0, 0);
        A.store.Bool(ptr + 0 + 13, false);
        A.store.Int32(ptr + 0 + 4, 0);
        A.store.Bool(ptr + 0 + 14, false);
        A.store.Int32(ptr + 0 + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);

        if (typeof x["touch_point"] === "undefined") {
          A.store.Bool(ptr + 0 + 15, false);
          A.store.Bool(ptr + 0 + 12, false);
          A.store.Int32(ptr + 0 + 0, 0);
          A.store.Bool(ptr + 0 + 13, false);
          A.store.Int32(ptr + 0 + 4, 0);
          A.store.Bool(ptr + 0 + 14, false);
          A.store.Int32(ptr + 0 + 8, 0);
        } else {
          A.store.Bool(ptr + 0 + 15, true);
          A.store.Bool(ptr + 0 + 12, "x" in x["touch_point"] ? true : false);
          A.store.Int32(ptr + 0 + 0, x["touch_point"]["x"] === undefined ? 0 : (x["touch_point"]["x"] as number));
          A.store.Bool(ptr + 0 + 13, "y" in x["touch_point"] ? true : false);
          A.store.Int32(ptr + 0 + 4, x["touch_point"]["y"] === undefined ? 0 : (x["touch_point"]["y"] as number));
          A.store.Bool(ptr + 0 + 14, "pressure" in x["touch_point"] ? true : false);
          A.store.Int32(
            ptr + 0 + 8,
            x["touch_point"]["pressure"] === undefined ? 0 : (x["touch_point"]["pressure"] as number)
          );
        }
      }
    },
    "load_StylusTouchEventInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 15)) {
        x["touch_point"] = {};
        if (A.load.Bool(ptr + 0 + 12)) {
          x["touch_point"]["x"] = A.load.Int32(ptr + 0 + 0);
        } else {
          delete x["touch_point"]["x"];
        }
        if (A.load.Bool(ptr + 0 + 13)) {
          x["touch_point"]["y"] = A.load.Int32(ptr + 0 + 4);
        } else {
          delete x["touch_point"]["y"];
        }
        if (A.load.Bool(ptr + 0 + 14)) {
          x["touch_point"]["pressure"] = A.load.Int32(ptr + 0 + 8);
        } else {
          delete x["touch_point"]["pressure"];
        }
      } else {
        delete x["touch_point"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TouchPointInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 30, false);
        A.store.Bool(ptr + 24, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 25, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 26, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 27, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 28, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Bool(ptr + 29, false);
        A.store.Int32(ptr + 20, 0);
      } else {
        A.store.Bool(ptr + 30, true);
        A.store.Bool(ptr + 24, "trackingId" in x ? true : false);
        A.store.Int32(ptr + 0, x["trackingId"] === undefined ? 0 : (x["trackingId"] as number));
        A.store.Bool(ptr + 25, "x" in x ? true : false);
        A.store.Int32(ptr + 4, x["x"] === undefined ? 0 : (x["x"] as number));
        A.store.Bool(ptr + 26, "y" in x ? true : false);
        A.store.Int32(ptr + 8, x["y"] === undefined ? 0 : (x["y"] as number));
        A.store.Bool(ptr + 27, "pressure" in x ? true : false);
        A.store.Int32(ptr + 12, x["pressure"] === undefined ? 0 : (x["pressure"] as number));
        A.store.Bool(ptr + 28, "touchMajor" in x ? true : false);
        A.store.Int32(ptr + 16, x["touchMajor"] === undefined ? 0 : (x["touchMajor"] as number));
        A.store.Bool(ptr + 29, "touchMinor" in x ? true : false);
        A.store.Int32(ptr + 20, x["touchMinor"] === undefined ? 0 : (x["touchMinor"] as number));
      }
    },
    "load_TouchPointInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["trackingId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["trackingId"];
      }
      if (A.load.Bool(ptr + 25)) {
        x["x"] = A.load.Int32(ptr + 4);
      } else {
        delete x["x"];
      }
      if (A.load.Bool(ptr + 26)) {
        x["y"] = A.load.Int32(ptr + 8);
      } else {
        delete x["y"];
      }
      if (A.load.Bool(ptr + 27)) {
        x["pressure"] = A.load.Int32(ptr + 12);
      } else {
        delete x["pressure"];
      }
      if (A.load.Bool(ptr + 28)) {
        x["touchMajor"] = A.load.Int32(ptr + 16);
      } else {
        delete x["touchMajor"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["touchMinor"] = A.load.Int32(ptr + 20);
      } else {
        delete x["touchMinor"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TouchpadButtonEventInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["left", "middle", "right"].indexOf(x["button"] as string));
        A.store.Enum(ptr + 4, ["pressed", "released"].indexOf(x["state"] as string));
      }
    },
    "load_TouchpadButtonEventInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["button"] = A.load.Enum(ptr + 0, ["left", "middle", "right"]);
      x["state"] = A.load.Enum(ptr + 4, ["pressed", "released"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TouchpadConnectedEventInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 18, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 19, true);
        A.store.Bool(ptr + 16, "maxX" in x ? true : false);
        A.store.Int32(ptr + 0, x["maxX"] === undefined ? 0 : (x["maxX"] as number));
        A.store.Bool(ptr + 17, "maxY" in x ? true : false);
        A.store.Int32(ptr + 4, x["maxY"] === undefined ? 0 : (x["maxY"] as number));
        A.store.Bool(ptr + 18, "maxPressure" in x ? true : false);
        A.store.Int32(ptr + 8, x["maxPressure"] === undefined ? 0 : (x["maxPressure"] as number));
        A.store.Ref(ptr + 12, x["buttons"]);
      }
    },
    "load_TouchpadConnectedEventInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["maxX"] = A.load.Int32(ptr + 0);
      } else {
        delete x["maxX"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["maxY"] = A.load.Int32(ptr + 4);
      } else {
        delete x["maxY"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["maxPressure"] = A.load.Int32(ptr + 8);
      } else {
        delete x["maxPressure"];
      }
      x["buttons"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TouchpadTouchEventInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["touchPoints"]);
      }
    },
    "load_TouchpadTouchEventInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["touchPoints"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TouchscreenConnectedEventInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 14, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 15, true);
        A.store.Bool(ptr + 12, "maxX" in x ? true : false);
        A.store.Int32(ptr + 0, x["maxX"] === undefined ? 0 : (x["maxX"] as number));
        A.store.Bool(ptr + 13, "maxY" in x ? true : false);
        A.store.Int32(ptr + 4, x["maxY"] === undefined ? 0 : (x["maxY"] as number));
        A.store.Bool(ptr + 14, "maxPressure" in x ? true : false);
        A.store.Int32(ptr + 8, x["maxPressure"] === undefined ? 0 : (x["maxPressure"] as number));
      }
    },
    "load_TouchscreenConnectedEventInfo": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["maxX"] = A.load.Int32(ptr + 0);
      } else {
        delete x["maxX"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["maxY"] = A.load.Int32(ptr + 4);
      } else {
        delete x["maxY"];
      }
      if (A.load.Bool(ptr + 14)) {
        x["maxPressure"] = A.load.Int32(ptr + 8);
      } else {
        delete x["maxPressure"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TouchscreenTouchEventInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["touchPoints"]);
      }
    },
    "load_TouchscreenTouchEventInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["touchPoints"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_UsbEvent": (ref: heap.Ref<string>): number => {
      const idx = ["connected", "disconnected"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_UsbEventInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 26, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 24, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 25, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Enum(ptr + 20, -1);
      } else {
        A.store.Bool(ptr + 26, true);
        A.store.Ref(ptr + 0, x["vendor"]);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Bool(ptr + 24, "vid" in x ? true : false);
        A.store.Int32(ptr + 8, x["vid"] === undefined ? 0 : (x["vid"] as number));
        A.store.Bool(ptr + 25, "pid" in x ? true : false);
        A.store.Int32(ptr + 12, x["pid"] === undefined ? 0 : (x["pid"] as number));
        A.store.Ref(ptr + 16, x["categories"]);
        A.store.Enum(ptr + 20, ["connected", "disconnected"].indexOf(x["event"] as string));
      }
    },
    "load_UsbEventInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["vendor"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 24)) {
        x["vid"] = A.load.Int32(ptr + 8);
      } else {
        delete x["vid"];
      }
      if (A.load.Bool(ptr + 25)) {
        x["pid"] = A.load.Int32(ptr + 12);
      } else {
        delete x["pid"];
      }
      x["categories"] = A.load.Ref(ptr + 16, undefined);
      x["event"] = A.load.Enum(ptr + 20, ["connected", "disconnected"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_IsEventSupported": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events && "isEventSupported" in WEBEXT?.os?.events) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsEventSupported": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.isEventSupported);
    },
    "call_IsEventSupported": (retPtr: Pointer, category: number): void => {
      const _ret = WEBEXT.os.events.isEventSupported(
        category > 0 && category <= 15
          ? [
              "audio_jack",
              "lid",
              "usb",
              "sd_card",
              "power",
              "keyboard_diagnostic",
              "stylus_garage",
              "touchpad_button",
              "touchpad_touch",
              "touchpad_connected",
              "touchscreen_touch",
              "touchscreen_connected",
              "external_display",
              "stylus_touch",
              "stylus_connected",
            ][category - 1]
          : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_IsEventSupported": (retPtr: Pointer, errPtr: Pointer, category: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.isEventSupported(
          category > 0 && category <= 15
            ? [
                "audio_jack",
                "lid",
                "usb",
                "sd_card",
                "power",
                "keyboard_diagnostic",
                "stylus_garage",
                "touchpad_button",
                "touchpad_touch",
                "touchpad_connected",
                "touchscreen_touch",
                "touchscreen_connected",
                "external_display",
                "stylus_touch",
                "stylus_connected",
              ][category - 1]
            : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAudioJackEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onAudioJackEvent && "addListener" in WEBEXT?.os?.events?.onAudioJackEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAudioJackEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onAudioJackEvent.addListener);
    },
    "call_OnAudioJackEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onAudioJackEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnAudioJackEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onAudioJackEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAudioJackEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onAudioJackEvent && "removeListener" in WEBEXT?.os?.events?.onAudioJackEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAudioJackEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onAudioJackEvent.removeListener);
    },
    "call_OffAudioJackEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onAudioJackEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffAudioJackEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onAudioJackEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAudioJackEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onAudioJackEvent && "hasListener" in WEBEXT?.os?.events?.onAudioJackEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAudioJackEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onAudioJackEvent.hasListener);
    },
    "call_HasOnAudioJackEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onAudioJackEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAudioJackEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onAudioJackEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnExternalDisplayEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onExternalDisplayEvent && "addListener" in WEBEXT?.os?.events?.onExternalDisplayEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnExternalDisplayEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onExternalDisplayEvent.addListener);
    },
    "call_OnExternalDisplayEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onExternalDisplayEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnExternalDisplayEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onExternalDisplayEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffExternalDisplayEvent": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.events?.onExternalDisplayEvent &&
        "removeListener" in WEBEXT?.os?.events?.onExternalDisplayEvent
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffExternalDisplayEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onExternalDisplayEvent.removeListener);
    },
    "call_OffExternalDisplayEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onExternalDisplayEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffExternalDisplayEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onExternalDisplayEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnExternalDisplayEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onExternalDisplayEvent && "hasListener" in WEBEXT?.os?.events?.onExternalDisplayEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnExternalDisplayEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onExternalDisplayEvent.hasListener);
    },
    "call_HasOnExternalDisplayEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onExternalDisplayEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnExternalDisplayEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onExternalDisplayEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnKeyboardDiagnosticEvent": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.events?.onKeyboardDiagnosticEvent &&
        "addListener" in WEBEXT?.os?.events?.onKeyboardDiagnosticEvent
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnKeyboardDiagnosticEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onKeyboardDiagnosticEvent.addListener);
    },
    "call_OnKeyboardDiagnosticEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onKeyboardDiagnosticEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnKeyboardDiagnosticEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onKeyboardDiagnosticEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffKeyboardDiagnosticEvent": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.events?.onKeyboardDiagnosticEvent &&
        "removeListener" in WEBEXT?.os?.events?.onKeyboardDiagnosticEvent
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffKeyboardDiagnosticEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onKeyboardDiagnosticEvent.removeListener);
    },
    "call_OffKeyboardDiagnosticEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onKeyboardDiagnosticEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffKeyboardDiagnosticEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onKeyboardDiagnosticEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnKeyboardDiagnosticEvent": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.events?.onKeyboardDiagnosticEvent &&
        "hasListener" in WEBEXT?.os?.events?.onKeyboardDiagnosticEvent
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnKeyboardDiagnosticEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onKeyboardDiagnosticEvent.hasListener);
    },
    "call_HasOnKeyboardDiagnosticEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onKeyboardDiagnosticEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnKeyboardDiagnosticEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onKeyboardDiagnosticEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnLidEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onLidEvent && "addListener" in WEBEXT?.os?.events?.onLidEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnLidEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onLidEvent.addListener);
    },
    "call_OnLidEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onLidEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnLidEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onLidEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffLidEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onLidEvent && "removeListener" in WEBEXT?.os?.events?.onLidEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffLidEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onLidEvent.removeListener);
    },
    "call_OffLidEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onLidEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffLidEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onLidEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnLidEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onLidEvent && "hasListener" in WEBEXT?.os?.events?.onLidEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnLidEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onLidEvent.hasListener);
    },
    "call_HasOnLidEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onLidEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnLidEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onLidEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPowerEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onPowerEvent && "addListener" in WEBEXT?.os?.events?.onPowerEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPowerEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onPowerEvent.addListener);
    },
    "call_OnPowerEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onPowerEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnPowerEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onPowerEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPowerEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onPowerEvent && "removeListener" in WEBEXT?.os?.events?.onPowerEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPowerEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onPowerEvent.removeListener);
    },
    "call_OffPowerEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onPowerEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffPowerEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onPowerEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPowerEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onPowerEvent && "hasListener" in WEBEXT?.os?.events?.onPowerEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPowerEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onPowerEvent.hasListener);
    },
    "call_HasOnPowerEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onPowerEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPowerEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onPowerEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSdCardEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onSdCardEvent && "addListener" in WEBEXT?.os?.events?.onSdCardEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSdCardEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onSdCardEvent.addListener);
    },
    "call_OnSdCardEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onSdCardEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnSdCardEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onSdCardEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSdCardEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onSdCardEvent && "removeListener" in WEBEXT?.os?.events?.onSdCardEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSdCardEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onSdCardEvent.removeListener);
    },
    "call_OffSdCardEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onSdCardEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffSdCardEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onSdCardEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSdCardEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onSdCardEvent && "hasListener" in WEBEXT?.os?.events?.onSdCardEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSdCardEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onSdCardEvent.hasListener);
    },
    "call_HasOnSdCardEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onSdCardEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSdCardEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onSdCardEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnStylusConnectedEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onStylusConnectedEvent && "addListener" in WEBEXT?.os?.events?.onStylusConnectedEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnStylusConnectedEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onStylusConnectedEvent.addListener);
    },
    "call_OnStylusConnectedEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onStylusConnectedEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnStylusConnectedEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onStylusConnectedEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffStylusConnectedEvent": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.events?.onStylusConnectedEvent &&
        "removeListener" in WEBEXT?.os?.events?.onStylusConnectedEvent
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffStylusConnectedEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onStylusConnectedEvent.removeListener);
    },
    "call_OffStylusConnectedEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onStylusConnectedEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffStylusConnectedEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onStylusConnectedEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnStylusConnectedEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onStylusConnectedEvent && "hasListener" in WEBEXT?.os?.events?.onStylusConnectedEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnStylusConnectedEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onStylusConnectedEvent.hasListener);
    },
    "call_HasOnStylusConnectedEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onStylusConnectedEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnStylusConnectedEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onStylusConnectedEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnStylusGarageEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onStylusGarageEvent && "addListener" in WEBEXT?.os?.events?.onStylusGarageEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnStylusGarageEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onStylusGarageEvent.addListener);
    },
    "call_OnStylusGarageEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onStylusGarageEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnStylusGarageEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onStylusGarageEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffStylusGarageEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onStylusGarageEvent && "removeListener" in WEBEXT?.os?.events?.onStylusGarageEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffStylusGarageEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onStylusGarageEvent.removeListener);
    },
    "call_OffStylusGarageEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onStylusGarageEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffStylusGarageEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onStylusGarageEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnStylusGarageEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onStylusGarageEvent && "hasListener" in WEBEXT?.os?.events?.onStylusGarageEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnStylusGarageEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onStylusGarageEvent.hasListener);
    },
    "call_HasOnStylusGarageEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onStylusGarageEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnStylusGarageEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onStylusGarageEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnStylusTouchEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onStylusTouchEvent && "addListener" in WEBEXT?.os?.events?.onStylusTouchEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnStylusTouchEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onStylusTouchEvent.addListener);
    },
    "call_OnStylusTouchEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onStylusTouchEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnStylusTouchEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onStylusTouchEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffStylusTouchEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onStylusTouchEvent && "removeListener" in WEBEXT?.os?.events?.onStylusTouchEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffStylusTouchEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onStylusTouchEvent.removeListener);
    },
    "call_OffStylusTouchEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onStylusTouchEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffStylusTouchEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onStylusTouchEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnStylusTouchEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onStylusTouchEvent && "hasListener" in WEBEXT?.os?.events?.onStylusTouchEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnStylusTouchEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onStylusTouchEvent.hasListener);
    },
    "call_HasOnStylusTouchEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onStylusTouchEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnStylusTouchEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onStylusTouchEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTouchpadButtonEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onTouchpadButtonEvent && "addListener" in WEBEXT?.os?.events?.onTouchpadButtonEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTouchpadButtonEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onTouchpadButtonEvent.addListener);
    },
    "call_OnTouchpadButtonEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onTouchpadButtonEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnTouchpadButtonEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onTouchpadButtonEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffTouchpadButtonEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onTouchpadButtonEvent && "removeListener" in WEBEXT?.os?.events?.onTouchpadButtonEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffTouchpadButtonEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onTouchpadButtonEvent.removeListener);
    },
    "call_OffTouchpadButtonEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onTouchpadButtonEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffTouchpadButtonEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onTouchpadButtonEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnTouchpadButtonEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onTouchpadButtonEvent && "hasListener" in WEBEXT?.os?.events?.onTouchpadButtonEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnTouchpadButtonEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onTouchpadButtonEvent.hasListener);
    },
    "call_HasOnTouchpadButtonEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onTouchpadButtonEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnTouchpadButtonEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onTouchpadButtonEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTouchpadConnectedEvent": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.events?.onTouchpadConnectedEvent &&
        "addListener" in WEBEXT?.os?.events?.onTouchpadConnectedEvent
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTouchpadConnectedEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onTouchpadConnectedEvent.addListener);
    },
    "call_OnTouchpadConnectedEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onTouchpadConnectedEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnTouchpadConnectedEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onTouchpadConnectedEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffTouchpadConnectedEvent": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.events?.onTouchpadConnectedEvent &&
        "removeListener" in WEBEXT?.os?.events?.onTouchpadConnectedEvent
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffTouchpadConnectedEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onTouchpadConnectedEvent.removeListener);
    },
    "call_OffTouchpadConnectedEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onTouchpadConnectedEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffTouchpadConnectedEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onTouchpadConnectedEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnTouchpadConnectedEvent": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.events?.onTouchpadConnectedEvent &&
        "hasListener" in WEBEXT?.os?.events?.onTouchpadConnectedEvent
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnTouchpadConnectedEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onTouchpadConnectedEvent.hasListener);
    },
    "call_HasOnTouchpadConnectedEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onTouchpadConnectedEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnTouchpadConnectedEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onTouchpadConnectedEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTouchpadTouchEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onTouchpadTouchEvent && "addListener" in WEBEXT?.os?.events?.onTouchpadTouchEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTouchpadTouchEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onTouchpadTouchEvent.addListener);
    },
    "call_OnTouchpadTouchEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onTouchpadTouchEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnTouchpadTouchEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onTouchpadTouchEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffTouchpadTouchEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onTouchpadTouchEvent && "removeListener" in WEBEXT?.os?.events?.onTouchpadTouchEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffTouchpadTouchEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onTouchpadTouchEvent.removeListener);
    },
    "call_OffTouchpadTouchEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onTouchpadTouchEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffTouchpadTouchEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onTouchpadTouchEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnTouchpadTouchEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onTouchpadTouchEvent && "hasListener" in WEBEXT?.os?.events?.onTouchpadTouchEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnTouchpadTouchEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onTouchpadTouchEvent.hasListener);
    },
    "call_HasOnTouchpadTouchEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onTouchpadTouchEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnTouchpadTouchEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onTouchpadTouchEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTouchscreenConnectedEvent": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.events?.onTouchscreenConnectedEvent &&
        "addListener" in WEBEXT?.os?.events?.onTouchscreenConnectedEvent
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTouchscreenConnectedEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onTouchscreenConnectedEvent.addListener);
    },
    "call_OnTouchscreenConnectedEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onTouchscreenConnectedEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnTouchscreenConnectedEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onTouchscreenConnectedEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffTouchscreenConnectedEvent": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.events?.onTouchscreenConnectedEvent &&
        "removeListener" in WEBEXT?.os?.events?.onTouchscreenConnectedEvent
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffTouchscreenConnectedEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onTouchscreenConnectedEvent.removeListener);
    },
    "call_OffTouchscreenConnectedEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onTouchscreenConnectedEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffTouchscreenConnectedEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onTouchscreenConnectedEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnTouchscreenConnectedEvent": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.events?.onTouchscreenConnectedEvent &&
        "hasListener" in WEBEXT?.os?.events?.onTouchscreenConnectedEvent
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnTouchscreenConnectedEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onTouchscreenConnectedEvent.hasListener);
    },
    "call_HasOnTouchscreenConnectedEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onTouchscreenConnectedEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnTouchscreenConnectedEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onTouchscreenConnectedEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTouchscreenTouchEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onTouchscreenTouchEvent && "addListener" in WEBEXT?.os?.events?.onTouchscreenTouchEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTouchscreenTouchEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onTouchscreenTouchEvent.addListener);
    },
    "call_OnTouchscreenTouchEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onTouchscreenTouchEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnTouchscreenTouchEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onTouchscreenTouchEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffTouchscreenTouchEvent": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.events?.onTouchscreenTouchEvent &&
        "removeListener" in WEBEXT?.os?.events?.onTouchscreenTouchEvent
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffTouchscreenTouchEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onTouchscreenTouchEvent.removeListener);
    },
    "call_OffTouchscreenTouchEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onTouchscreenTouchEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffTouchscreenTouchEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onTouchscreenTouchEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnTouchscreenTouchEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onTouchscreenTouchEvent && "hasListener" in WEBEXT?.os?.events?.onTouchscreenTouchEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnTouchscreenTouchEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onTouchscreenTouchEvent.hasListener);
    },
    "call_HasOnTouchscreenTouchEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onTouchscreenTouchEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnTouchscreenTouchEvent": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onTouchscreenTouchEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnUsbEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onUsbEvent && "addListener" in WEBEXT?.os?.events?.onUsbEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnUsbEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onUsbEvent.addListener);
    },
    "call_OnUsbEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onUsbEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnUsbEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onUsbEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffUsbEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onUsbEvent && "removeListener" in WEBEXT?.os?.events?.onUsbEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffUsbEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onUsbEvent.removeListener);
    },
    "call_OffUsbEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onUsbEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffUsbEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onUsbEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnUsbEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events?.onUsbEvent && "hasListener" in WEBEXT?.os?.events?.onUsbEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnUsbEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.onUsbEvent.hasListener);
    },
    "call_HasOnUsbEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.events.onUsbEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnUsbEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.onUsbEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartCapturingEvents": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events && "startCapturingEvents" in WEBEXT?.os?.events) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartCapturingEvents": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.startCapturingEvents);
    },
    "call_StartCapturingEvents": (retPtr: Pointer, category: number): void => {
      const _ret = WEBEXT.os.events.startCapturingEvents(
        category > 0 && category <= 15
          ? [
              "audio_jack",
              "lid",
              "usb",
              "sd_card",
              "power",
              "keyboard_diagnostic",
              "stylus_garage",
              "touchpad_button",
              "touchpad_touch",
              "touchpad_connected",
              "touchscreen_touch",
              "touchscreen_connected",
              "external_display",
              "stylus_touch",
              "stylus_connected",
            ][category - 1]
          : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_StartCapturingEvents": (retPtr: Pointer, errPtr: Pointer, category: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.startCapturingEvents(
          category > 0 && category <= 15
            ? [
                "audio_jack",
                "lid",
                "usb",
                "sd_card",
                "power",
                "keyboard_diagnostic",
                "stylus_garage",
                "touchpad_button",
                "touchpad_touch",
                "touchpad_connected",
                "touchscreen_touch",
                "touchscreen_connected",
                "external_display",
                "stylus_touch",
                "stylus_connected",
              ][category - 1]
            : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StopCapturingEvents": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.events && "stopCapturingEvents" in WEBEXT?.os?.events) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StopCapturingEvents": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.events.stopCapturingEvents);
    },
    "call_StopCapturingEvents": (retPtr: Pointer, category: number): void => {
      const _ret = WEBEXT.os.events.stopCapturingEvents(
        category > 0 && category <= 15
          ? [
              "audio_jack",
              "lid",
              "usb",
              "sd_card",
              "power",
              "keyboard_diagnostic",
              "stylus_garage",
              "touchpad_button",
              "touchpad_touch",
              "touchpad_connected",
              "touchscreen_touch",
              "touchscreen_connected",
              "external_display",
              "stylus_touch",
              "stylus_connected",
            ][category - 1]
          : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_StopCapturingEvents": (retPtr: Pointer, errPtr: Pointer, category: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.events.stopCapturingEvents(
          category > 0 && category <= 15
            ? [
                "audio_jack",
                "lid",
                "usb",
                "sd_card",
                "power",
                "keyboard_diagnostic",
                "stylus_garage",
                "touchpad_button",
                "touchpad_touch",
                "touchpad_connected",
                "touchscreen_touch",
                "touchscreen_connected",
                "external_display",
                "stylus_touch",
                "stylus_connected",
              ][category - 1]
            : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
