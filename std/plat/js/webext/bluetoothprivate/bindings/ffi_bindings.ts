import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/bluetoothprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_ConnectResultType": (ref: heap.Ref<string>): number => {
      const idx = [
        "alreadyConnected",
        "authCanceled",
        "authFailed",
        "authRejected",
        "authTimeout",
        "failed",
        "inProgress",
        "success",
        "unknownError",
        "unsupportedDevice",
        "notReady",
        "alreadyExists",
        "notConnected",
        "doesNotExist",
        "invalidArgs",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_TransportType": (ref: heap.Ref<string>): number => {
      const idx = ["le", "bredr", "dual"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DiscoveryFilter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 12, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Enum(ptr + 0, ["le", "bredr", "dual"].indexOf(x["transport"] as string));
        A.store.Ref(ptr + 4, x["uuids"]);
        A.store.Bool(ptr + 16, "rssi" in x ? true : false);
        A.store.Int32(ptr + 8, x["rssi"] === undefined ? 0 : (x["rssi"] as number));
        A.store.Bool(ptr + 17, "pathloss" in x ? true : false);
        A.store.Int32(ptr + 12, x["pathloss"] === undefined ? 0 : (x["pathloss"] as number));
      }
    },
    "load_DiscoveryFilter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["transport"] = A.load.Enum(ptr + 0, ["le", "bredr", "dual"]);
      x["uuids"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["rssi"] = A.load.Int32(ptr + 8);
      } else {
        delete x["rssi"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["pathloss"] = A.load.Int32(ptr + 12);
      } else {
        delete x["pathloss"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_NewAdapterState": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 5, false);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Bool(ptr + 6, "powered" in x ? true : false);
        A.store.Bool(ptr + 4, x["powered"] ? true : false);
        A.store.Bool(ptr + 7, "discoverable" in x ? true : false);
        A.store.Bool(ptr + 5, x["discoverable"] ? true : false);
      }
    },
    "load_NewAdapterState": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 6)) {
        x["powered"] = A.load.Bool(ptr + 4);
      } else {
        delete x["powered"];
      }
      if (A.load.Bool(ptr + 7)) {
        x["discoverable"] = A.load.Bool(ptr + 5);
      } else {
        delete x["discoverable"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PairingEventType": (ref: heap.Ref<string>): number => {
      const idx = [
        "requestPincode",
        "displayPincode",
        "requestPasskey",
        "displayPasskey",
        "keysEntered",
        "confirmPasskey",
        "requestAuthorization",
        "complete",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PairingEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 86, false);
        A.store.Enum(ptr + 0, -1);

        A.store.Bool(ptr + 4 + 67, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Ref(ptr + 4 + 4, undefined);
        A.store.Bool(ptr + 4 + 56, false);
        A.store.Int32(ptr + 4 + 8, 0);
        A.store.Enum(ptr + 4 + 12, -1);
        A.store.Bool(ptr + 4 + 57, false);
        A.store.Int32(ptr + 4 + 16, 0);
        A.store.Bool(ptr + 4 + 58, false);
        A.store.Int32(ptr + 4 + 20, 0);
        A.store.Bool(ptr + 4 + 59, false);
        A.store.Int32(ptr + 4 + 24, 0);
        A.store.Enum(ptr + 4 + 28, -1);
        A.store.Bool(ptr + 4 + 60, false);
        A.store.Bool(ptr + 4 + 32, false);
        A.store.Bool(ptr + 4 + 61, false);
        A.store.Bool(ptr + 4 + 33, false);
        A.store.Bool(ptr + 4 + 62, false);
        A.store.Bool(ptr + 4 + 34, false);
        A.store.Bool(ptr + 4 + 63, false);
        A.store.Bool(ptr + 4 + 35, false);
        A.store.Ref(ptr + 4 + 36, undefined);
        A.store.Bool(ptr + 4 + 64, false);
        A.store.Int32(ptr + 4 + 40, 0);
        A.store.Bool(ptr + 4 + 65, false);
        A.store.Int32(ptr + 4 + 44, 0);
        A.store.Enum(ptr + 4 + 48, -1);
        A.store.Bool(ptr + 4 + 66, false);
        A.store.Int32(ptr + 4 + 52, 0);
        A.store.Ref(ptr + 72, undefined);
        A.store.Bool(ptr + 84, false);
        A.store.Int32(ptr + 76, 0);
        A.store.Bool(ptr + 85, false);
        A.store.Int32(ptr + 80, 0);
      } else {
        A.store.Bool(ptr + 86, true);
        A.store.Enum(
          ptr + 0,
          [
            "requestPincode",
            "displayPincode",
            "requestPasskey",
            "displayPasskey",
            "keysEntered",
            "confirmPasskey",
            "requestAuthorization",
            "complete",
          ].indexOf(x["pairing"] as string)
        );

        if (typeof x["device"] === "undefined") {
          A.store.Bool(ptr + 4 + 67, false);
          A.store.Ref(ptr + 4 + 0, undefined);
          A.store.Ref(ptr + 4 + 4, undefined);
          A.store.Bool(ptr + 4 + 56, false);
          A.store.Int32(ptr + 4 + 8, 0);
          A.store.Enum(ptr + 4 + 12, -1);
          A.store.Bool(ptr + 4 + 57, false);
          A.store.Int32(ptr + 4 + 16, 0);
          A.store.Bool(ptr + 4 + 58, false);
          A.store.Int32(ptr + 4 + 20, 0);
          A.store.Bool(ptr + 4 + 59, false);
          A.store.Int32(ptr + 4 + 24, 0);
          A.store.Enum(ptr + 4 + 28, -1);
          A.store.Bool(ptr + 4 + 60, false);
          A.store.Bool(ptr + 4 + 32, false);
          A.store.Bool(ptr + 4 + 61, false);
          A.store.Bool(ptr + 4 + 33, false);
          A.store.Bool(ptr + 4 + 62, false);
          A.store.Bool(ptr + 4 + 34, false);
          A.store.Bool(ptr + 4 + 63, false);
          A.store.Bool(ptr + 4 + 35, false);
          A.store.Ref(ptr + 4 + 36, undefined);
          A.store.Bool(ptr + 4 + 64, false);
          A.store.Int32(ptr + 4 + 40, 0);
          A.store.Bool(ptr + 4 + 65, false);
          A.store.Int32(ptr + 4 + 44, 0);
          A.store.Enum(ptr + 4 + 48, -1);
          A.store.Bool(ptr + 4 + 66, false);
          A.store.Int32(ptr + 4 + 52, 0);
        } else {
          A.store.Bool(ptr + 4 + 67, true);
          A.store.Ref(ptr + 4 + 0, x["device"]["address"]);
          A.store.Ref(ptr + 4 + 4, x["device"]["name"]);
          A.store.Bool(ptr + 4 + 56, "deviceClass" in x["device"] ? true : false);
          A.store.Int32(
            ptr + 4 + 8,
            x["device"]["deviceClass"] === undefined ? 0 : (x["device"]["deviceClass"] as number)
          );
          A.store.Enum(ptr + 4 + 12, ["bluetooth", "usb"].indexOf(x["device"]["vendorIdSource"] as string));
          A.store.Bool(ptr + 4 + 57, "vendorId" in x["device"] ? true : false);
          A.store.Int32(ptr + 4 + 16, x["device"]["vendorId"] === undefined ? 0 : (x["device"]["vendorId"] as number));
          A.store.Bool(ptr + 4 + 58, "productId" in x["device"] ? true : false);
          A.store.Int32(
            ptr + 4 + 20,
            x["device"]["productId"] === undefined ? 0 : (x["device"]["productId"] as number)
          );
          A.store.Bool(ptr + 4 + 59, "deviceId" in x["device"] ? true : false);
          A.store.Int32(ptr + 4 + 24, x["device"]["deviceId"] === undefined ? 0 : (x["device"]["deviceId"] as number));
          A.store.Enum(
            ptr + 4 + 28,
            [
              "computer",
              "phone",
              "modem",
              "audio",
              "carAudio",
              "video",
              "peripheral",
              "joystick",
              "gamepad",
              "keyboard",
              "mouse",
              "tablet",
              "keyboardMouseCombo",
            ].indexOf(x["device"]["type"] as string)
          );
          A.store.Bool(ptr + 4 + 60, "paired" in x["device"] ? true : false);
          A.store.Bool(ptr + 4 + 32, x["device"]["paired"] ? true : false);
          A.store.Bool(ptr + 4 + 61, "connected" in x["device"] ? true : false);
          A.store.Bool(ptr + 4 + 33, x["device"]["connected"] ? true : false);
          A.store.Bool(ptr + 4 + 62, "connecting" in x["device"] ? true : false);
          A.store.Bool(ptr + 4 + 34, x["device"]["connecting"] ? true : false);
          A.store.Bool(ptr + 4 + 63, "connectable" in x["device"] ? true : false);
          A.store.Bool(ptr + 4 + 35, x["device"]["connectable"] ? true : false);
          A.store.Ref(ptr + 4 + 36, x["device"]["uuids"]);
          A.store.Bool(ptr + 4 + 64, "inquiryRssi" in x["device"] ? true : false);
          A.store.Int32(
            ptr + 4 + 40,
            x["device"]["inquiryRssi"] === undefined ? 0 : (x["device"]["inquiryRssi"] as number)
          );
          A.store.Bool(ptr + 4 + 65, "inquiryTxPower" in x["device"] ? true : false);
          A.store.Int32(
            ptr + 4 + 44,
            x["device"]["inquiryTxPower"] === undefined ? 0 : (x["device"]["inquiryTxPower"] as number)
          );
          A.store.Enum(ptr + 4 + 48, ["invalid", "classic", "le", "dual"].indexOf(x["device"]["transport"] as string));
          A.store.Bool(ptr + 4 + 66, "batteryPercentage" in x["device"] ? true : false);
          A.store.Int32(
            ptr + 4 + 52,
            x["device"]["batteryPercentage"] === undefined ? 0 : (x["device"]["batteryPercentage"] as number)
          );
        }
        A.store.Ref(ptr + 72, x["pincode"]);
        A.store.Bool(ptr + 84, "passkey" in x ? true : false);
        A.store.Int32(ptr + 76, x["passkey"] === undefined ? 0 : (x["passkey"] as number));
        A.store.Bool(ptr + 85, "enteredKey" in x ? true : false);
        A.store.Int32(ptr + 80, x["enteredKey"] === undefined ? 0 : (x["enteredKey"] as number));
      }
    },
    "load_PairingEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["pairing"] = A.load.Enum(ptr + 0, [
        "requestPincode",
        "displayPincode",
        "requestPasskey",
        "displayPasskey",
        "keysEntered",
        "confirmPasskey",
        "requestAuthorization",
        "complete",
      ]);
      if (A.load.Bool(ptr + 4 + 67)) {
        x["device"] = {};
        x["device"]["address"] = A.load.Ref(ptr + 4 + 0, undefined);
        x["device"]["name"] = A.load.Ref(ptr + 4 + 4, undefined);
        if (A.load.Bool(ptr + 4 + 56)) {
          x["device"]["deviceClass"] = A.load.Int32(ptr + 4 + 8);
        } else {
          delete x["device"]["deviceClass"];
        }
        x["device"]["vendorIdSource"] = A.load.Enum(ptr + 4 + 12, ["bluetooth", "usb"]);
        if (A.load.Bool(ptr + 4 + 57)) {
          x["device"]["vendorId"] = A.load.Int32(ptr + 4 + 16);
        } else {
          delete x["device"]["vendorId"];
        }
        if (A.load.Bool(ptr + 4 + 58)) {
          x["device"]["productId"] = A.load.Int32(ptr + 4 + 20);
        } else {
          delete x["device"]["productId"];
        }
        if (A.load.Bool(ptr + 4 + 59)) {
          x["device"]["deviceId"] = A.load.Int32(ptr + 4 + 24);
        } else {
          delete x["device"]["deviceId"];
        }
        x["device"]["type"] = A.load.Enum(ptr + 4 + 28, [
          "computer",
          "phone",
          "modem",
          "audio",
          "carAudio",
          "video",
          "peripheral",
          "joystick",
          "gamepad",
          "keyboard",
          "mouse",
          "tablet",
          "keyboardMouseCombo",
        ]);
        if (A.load.Bool(ptr + 4 + 60)) {
          x["device"]["paired"] = A.load.Bool(ptr + 4 + 32);
        } else {
          delete x["device"]["paired"];
        }
        if (A.load.Bool(ptr + 4 + 61)) {
          x["device"]["connected"] = A.load.Bool(ptr + 4 + 33);
        } else {
          delete x["device"]["connected"];
        }
        if (A.load.Bool(ptr + 4 + 62)) {
          x["device"]["connecting"] = A.load.Bool(ptr + 4 + 34);
        } else {
          delete x["device"]["connecting"];
        }
        if (A.load.Bool(ptr + 4 + 63)) {
          x["device"]["connectable"] = A.load.Bool(ptr + 4 + 35);
        } else {
          delete x["device"]["connectable"];
        }
        x["device"]["uuids"] = A.load.Ref(ptr + 4 + 36, undefined);
        if (A.load.Bool(ptr + 4 + 64)) {
          x["device"]["inquiryRssi"] = A.load.Int32(ptr + 4 + 40);
        } else {
          delete x["device"]["inquiryRssi"];
        }
        if (A.load.Bool(ptr + 4 + 65)) {
          x["device"]["inquiryTxPower"] = A.load.Int32(ptr + 4 + 44);
        } else {
          delete x["device"]["inquiryTxPower"];
        }
        x["device"]["transport"] = A.load.Enum(ptr + 4 + 48, ["invalid", "classic", "le", "dual"]);
        if (A.load.Bool(ptr + 4 + 66)) {
          x["device"]["batteryPercentage"] = A.load.Int32(ptr + 4 + 52);
        } else {
          delete x["device"]["batteryPercentage"];
        }
      } else {
        delete x["device"];
      }
      x["pincode"] = A.load.Ref(ptr + 72, undefined);
      if (A.load.Bool(ptr + 84)) {
        x["passkey"] = A.load.Int32(ptr + 76);
      } else {
        delete x["passkey"];
      }
      if (A.load.Bool(ptr + 85)) {
        x["enteredKey"] = A.load.Int32(ptr + 80);
      } else {
        delete x["enteredKey"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PairingResponse": (ref: heap.Ref<string>): number => {
      const idx = ["confirm", "reject", "cancel"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SetPairingResponseOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 81, false);

        A.store.Bool(ptr + 0 + 67, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Bool(ptr + 0 + 56, false);
        A.store.Int32(ptr + 0 + 8, 0);
        A.store.Enum(ptr + 0 + 12, -1);
        A.store.Bool(ptr + 0 + 57, false);
        A.store.Int32(ptr + 0 + 16, 0);
        A.store.Bool(ptr + 0 + 58, false);
        A.store.Int32(ptr + 0 + 20, 0);
        A.store.Bool(ptr + 0 + 59, false);
        A.store.Int32(ptr + 0 + 24, 0);
        A.store.Enum(ptr + 0 + 28, -1);
        A.store.Bool(ptr + 0 + 60, false);
        A.store.Bool(ptr + 0 + 32, false);
        A.store.Bool(ptr + 0 + 61, false);
        A.store.Bool(ptr + 0 + 33, false);
        A.store.Bool(ptr + 0 + 62, false);
        A.store.Bool(ptr + 0 + 34, false);
        A.store.Bool(ptr + 0 + 63, false);
        A.store.Bool(ptr + 0 + 35, false);
        A.store.Ref(ptr + 0 + 36, undefined);
        A.store.Bool(ptr + 0 + 64, false);
        A.store.Int32(ptr + 0 + 40, 0);
        A.store.Bool(ptr + 0 + 65, false);
        A.store.Int32(ptr + 0 + 44, 0);
        A.store.Enum(ptr + 0 + 48, -1);
        A.store.Bool(ptr + 0 + 66, false);
        A.store.Int32(ptr + 0 + 52, 0);
        A.store.Enum(ptr + 68, -1);
        A.store.Ref(ptr + 72, undefined);
        A.store.Bool(ptr + 80, false);
        A.store.Int32(ptr + 76, 0);
      } else {
        A.store.Bool(ptr + 81, true);

        if (typeof x["device"] === "undefined") {
          A.store.Bool(ptr + 0 + 67, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Bool(ptr + 0 + 56, false);
          A.store.Int32(ptr + 0 + 8, 0);
          A.store.Enum(ptr + 0 + 12, -1);
          A.store.Bool(ptr + 0 + 57, false);
          A.store.Int32(ptr + 0 + 16, 0);
          A.store.Bool(ptr + 0 + 58, false);
          A.store.Int32(ptr + 0 + 20, 0);
          A.store.Bool(ptr + 0 + 59, false);
          A.store.Int32(ptr + 0 + 24, 0);
          A.store.Enum(ptr + 0 + 28, -1);
          A.store.Bool(ptr + 0 + 60, false);
          A.store.Bool(ptr + 0 + 32, false);
          A.store.Bool(ptr + 0 + 61, false);
          A.store.Bool(ptr + 0 + 33, false);
          A.store.Bool(ptr + 0 + 62, false);
          A.store.Bool(ptr + 0 + 34, false);
          A.store.Bool(ptr + 0 + 63, false);
          A.store.Bool(ptr + 0 + 35, false);
          A.store.Ref(ptr + 0 + 36, undefined);
          A.store.Bool(ptr + 0 + 64, false);
          A.store.Int32(ptr + 0 + 40, 0);
          A.store.Bool(ptr + 0 + 65, false);
          A.store.Int32(ptr + 0 + 44, 0);
          A.store.Enum(ptr + 0 + 48, -1);
          A.store.Bool(ptr + 0 + 66, false);
          A.store.Int32(ptr + 0 + 52, 0);
        } else {
          A.store.Bool(ptr + 0 + 67, true);
          A.store.Ref(ptr + 0 + 0, x["device"]["address"]);
          A.store.Ref(ptr + 0 + 4, x["device"]["name"]);
          A.store.Bool(ptr + 0 + 56, "deviceClass" in x["device"] ? true : false);
          A.store.Int32(
            ptr + 0 + 8,
            x["device"]["deviceClass"] === undefined ? 0 : (x["device"]["deviceClass"] as number)
          );
          A.store.Enum(ptr + 0 + 12, ["bluetooth", "usb"].indexOf(x["device"]["vendorIdSource"] as string));
          A.store.Bool(ptr + 0 + 57, "vendorId" in x["device"] ? true : false);
          A.store.Int32(ptr + 0 + 16, x["device"]["vendorId"] === undefined ? 0 : (x["device"]["vendorId"] as number));
          A.store.Bool(ptr + 0 + 58, "productId" in x["device"] ? true : false);
          A.store.Int32(
            ptr + 0 + 20,
            x["device"]["productId"] === undefined ? 0 : (x["device"]["productId"] as number)
          );
          A.store.Bool(ptr + 0 + 59, "deviceId" in x["device"] ? true : false);
          A.store.Int32(ptr + 0 + 24, x["device"]["deviceId"] === undefined ? 0 : (x["device"]["deviceId"] as number));
          A.store.Enum(
            ptr + 0 + 28,
            [
              "computer",
              "phone",
              "modem",
              "audio",
              "carAudio",
              "video",
              "peripheral",
              "joystick",
              "gamepad",
              "keyboard",
              "mouse",
              "tablet",
              "keyboardMouseCombo",
            ].indexOf(x["device"]["type"] as string)
          );
          A.store.Bool(ptr + 0 + 60, "paired" in x["device"] ? true : false);
          A.store.Bool(ptr + 0 + 32, x["device"]["paired"] ? true : false);
          A.store.Bool(ptr + 0 + 61, "connected" in x["device"] ? true : false);
          A.store.Bool(ptr + 0 + 33, x["device"]["connected"] ? true : false);
          A.store.Bool(ptr + 0 + 62, "connecting" in x["device"] ? true : false);
          A.store.Bool(ptr + 0 + 34, x["device"]["connecting"] ? true : false);
          A.store.Bool(ptr + 0 + 63, "connectable" in x["device"] ? true : false);
          A.store.Bool(ptr + 0 + 35, x["device"]["connectable"] ? true : false);
          A.store.Ref(ptr + 0 + 36, x["device"]["uuids"]);
          A.store.Bool(ptr + 0 + 64, "inquiryRssi" in x["device"] ? true : false);
          A.store.Int32(
            ptr + 0 + 40,
            x["device"]["inquiryRssi"] === undefined ? 0 : (x["device"]["inquiryRssi"] as number)
          );
          A.store.Bool(ptr + 0 + 65, "inquiryTxPower" in x["device"] ? true : false);
          A.store.Int32(
            ptr + 0 + 44,
            x["device"]["inquiryTxPower"] === undefined ? 0 : (x["device"]["inquiryTxPower"] as number)
          );
          A.store.Enum(ptr + 0 + 48, ["invalid", "classic", "le", "dual"].indexOf(x["device"]["transport"] as string));
          A.store.Bool(ptr + 0 + 66, "batteryPercentage" in x["device"] ? true : false);
          A.store.Int32(
            ptr + 0 + 52,
            x["device"]["batteryPercentage"] === undefined ? 0 : (x["device"]["batteryPercentage"] as number)
          );
        }
        A.store.Enum(ptr + 68, ["confirm", "reject", "cancel"].indexOf(x["response"] as string));
        A.store.Ref(ptr + 72, x["pincode"]);
        A.store.Bool(ptr + 80, "passkey" in x ? true : false);
        A.store.Int32(ptr + 76, x["passkey"] === undefined ? 0 : (x["passkey"] as number));
      }
    },
    "load_SetPairingResponseOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 67)) {
        x["device"] = {};
        x["device"]["address"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["device"]["name"] = A.load.Ref(ptr + 0 + 4, undefined);
        if (A.load.Bool(ptr + 0 + 56)) {
          x["device"]["deviceClass"] = A.load.Int32(ptr + 0 + 8);
        } else {
          delete x["device"]["deviceClass"];
        }
        x["device"]["vendorIdSource"] = A.load.Enum(ptr + 0 + 12, ["bluetooth", "usb"]);
        if (A.load.Bool(ptr + 0 + 57)) {
          x["device"]["vendorId"] = A.load.Int32(ptr + 0 + 16);
        } else {
          delete x["device"]["vendorId"];
        }
        if (A.load.Bool(ptr + 0 + 58)) {
          x["device"]["productId"] = A.load.Int32(ptr + 0 + 20);
        } else {
          delete x["device"]["productId"];
        }
        if (A.load.Bool(ptr + 0 + 59)) {
          x["device"]["deviceId"] = A.load.Int32(ptr + 0 + 24);
        } else {
          delete x["device"]["deviceId"];
        }
        x["device"]["type"] = A.load.Enum(ptr + 0 + 28, [
          "computer",
          "phone",
          "modem",
          "audio",
          "carAudio",
          "video",
          "peripheral",
          "joystick",
          "gamepad",
          "keyboard",
          "mouse",
          "tablet",
          "keyboardMouseCombo",
        ]);
        if (A.load.Bool(ptr + 0 + 60)) {
          x["device"]["paired"] = A.load.Bool(ptr + 0 + 32);
        } else {
          delete x["device"]["paired"];
        }
        if (A.load.Bool(ptr + 0 + 61)) {
          x["device"]["connected"] = A.load.Bool(ptr + 0 + 33);
        } else {
          delete x["device"]["connected"];
        }
        if (A.load.Bool(ptr + 0 + 62)) {
          x["device"]["connecting"] = A.load.Bool(ptr + 0 + 34);
        } else {
          delete x["device"]["connecting"];
        }
        if (A.load.Bool(ptr + 0 + 63)) {
          x["device"]["connectable"] = A.load.Bool(ptr + 0 + 35);
        } else {
          delete x["device"]["connectable"];
        }
        x["device"]["uuids"] = A.load.Ref(ptr + 0 + 36, undefined);
        if (A.load.Bool(ptr + 0 + 64)) {
          x["device"]["inquiryRssi"] = A.load.Int32(ptr + 0 + 40);
        } else {
          delete x["device"]["inquiryRssi"];
        }
        if (A.load.Bool(ptr + 0 + 65)) {
          x["device"]["inquiryTxPower"] = A.load.Int32(ptr + 0 + 44);
        } else {
          delete x["device"]["inquiryTxPower"];
        }
        x["device"]["transport"] = A.load.Enum(ptr + 0 + 48, ["invalid", "classic", "le", "dual"]);
        if (A.load.Bool(ptr + 0 + 66)) {
          x["device"]["batteryPercentage"] = A.load.Int32(ptr + 0 + 52);
        } else {
          delete x["device"]["batteryPercentage"];
        }
      } else {
        delete x["device"];
      }
      x["response"] = A.load.Enum(ptr + 68, ["confirm", "reject", "cancel"]);
      x["pincode"] = A.load.Ref(ptr + 72, undefined);
      if (A.load.Bool(ptr + 80)) {
        x["passkey"] = A.load.Int32(ptr + 76);
      } else {
        delete x["passkey"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Connect": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothPrivate && "connect" in WEBEXT?.bluetoothPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Connect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.connect);
    },
    "call_Connect": (retPtr: Pointer, deviceAddress: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothPrivate.connect(A.H.get<object>(deviceAddress));
      A.store.Ref(retPtr, _ret);
    },
    "try_Connect": (retPtr: Pointer, errPtr: Pointer, deviceAddress: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothPrivate.connect(A.H.get<object>(deviceAddress));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DisconnectAll": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothPrivate && "disconnectAll" in WEBEXT?.bluetoothPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DisconnectAll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.disconnectAll);
    },
    "call_DisconnectAll": (retPtr: Pointer, deviceAddress: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothPrivate.disconnectAll(A.H.get<object>(deviceAddress));
      A.store.Ref(retPtr, _ret);
    },
    "try_DisconnectAll": (retPtr: Pointer, errPtr: Pointer, deviceAddress: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothPrivate.disconnectAll(A.H.get<object>(deviceAddress));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ForgetDevice": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothPrivate && "forgetDevice" in WEBEXT?.bluetoothPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ForgetDevice": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.forgetDevice);
    },
    "call_ForgetDevice": (retPtr: Pointer, deviceAddress: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothPrivate.forgetDevice(A.H.get<object>(deviceAddress));
      A.store.Ref(retPtr, _ret);
    },
    "try_ForgetDevice": (retPtr: Pointer, errPtr: Pointer, deviceAddress: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothPrivate.forgetDevice(A.H.get<object>(deviceAddress));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeviceAddressChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothPrivate?.onDeviceAddressChanged &&
        "addListener" in WEBEXT?.bluetoothPrivate?.onDeviceAddressChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeviceAddressChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.onDeviceAddressChanged.addListener);
    },
    "call_OnDeviceAddressChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothPrivate.onDeviceAddressChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnDeviceAddressChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothPrivate.onDeviceAddressChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeviceAddressChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothPrivate?.onDeviceAddressChanged &&
        "removeListener" in WEBEXT?.bluetoothPrivate?.onDeviceAddressChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeviceAddressChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.onDeviceAddressChanged.removeListener);
    },
    "call_OffDeviceAddressChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothPrivate.onDeviceAddressChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeviceAddressChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothPrivate.onDeviceAddressChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeviceAddressChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothPrivate?.onDeviceAddressChanged &&
        "hasListener" in WEBEXT?.bluetoothPrivate?.onDeviceAddressChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeviceAddressChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.onDeviceAddressChanged.hasListener);
    },
    "call_HasOnDeviceAddressChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothPrivate.onDeviceAddressChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeviceAddressChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothPrivate.onDeviceAddressChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPairing": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothPrivate?.onPairing && "addListener" in WEBEXT?.bluetoothPrivate?.onPairing) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPairing": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.onPairing.addListener);
    },
    "call_OnPairing": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothPrivate.onPairing.addListener(A.H.get<object>(callback));
    },
    "try_OnPairing": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothPrivate.onPairing.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPairing": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothPrivate?.onPairing && "removeListener" in WEBEXT?.bluetoothPrivate?.onPairing) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPairing": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.onPairing.removeListener);
    },
    "call_OffPairing": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothPrivate.onPairing.removeListener(A.H.get<object>(callback));
    },
    "try_OffPairing": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothPrivate.onPairing.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPairing": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothPrivate?.onPairing && "hasListener" in WEBEXT?.bluetoothPrivate?.onPairing) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPairing": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.onPairing.hasListener);
    },
    "call_HasOnPairing": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothPrivate.onPairing.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPairing": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothPrivate.onPairing.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Pair": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothPrivate && "pair" in WEBEXT?.bluetoothPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Pair": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.pair);
    },
    "call_Pair": (retPtr: Pointer, deviceAddress: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothPrivate.pair(A.H.get<object>(deviceAddress));
      A.store.Ref(retPtr, _ret);
    },
    "try_Pair": (retPtr: Pointer, errPtr: Pointer, deviceAddress: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothPrivate.pair(A.H.get<object>(deviceAddress));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordDeviceSelection": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothPrivate && "recordDeviceSelection" in WEBEXT?.bluetoothPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordDeviceSelection": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.recordDeviceSelection);
    },
    "call_RecordDeviceSelection": (
      retPtr: Pointer,
      selectionDurationMs: number,
      wasPaired: heap.Ref<boolean>,
      transport: number
    ): void => {
      const _ret = WEBEXT.bluetoothPrivate.recordDeviceSelection(
        selectionDurationMs,
        wasPaired === A.H.TRUE,
        transport > 0 && transport <= 4 ? ["invalid", "classic", "le", "dual"][transport - 1] : undefined
      );
    },
    "try_RecordDeviceSelection": (
      retPtr: Pointer,
      errPtr: Pointer,
      selectionDurationMs: number,
      wasPaired: heap.Ref<boolean>,
      transport: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothPrivate.recordDeviceSelection(
          selectionDurationMs,
          wasPaired === A.H.TRUE,
          transport > 0 && transport <= 4 ? ["invalid", "classic", "le", "dual"][transport - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordPairing": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothPrivate && "recordPairing" in WEBEXT?.bluetoothPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordPairing": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.recordPairing);
    },
    "call_RecordPairing": (retPtr: Pointer, transport: number, pairingDurationMs: number, result: number): void => {
      const _ret = WEBEXT.bluetoothPrivate.recordPairing(
        transport > 0 && transport <= 4 ? ["invalid", "classic", "le", "dual"][transport - 1] : undefined,
        pairingDurationMs,
        result > 0 && result <= 15
          ? [
              "alreadyConnected",
              "authCanceled",
              "authFailed",
              "authRejected",
              "authTimeout",
              "failed",
              "inProgress",
              "success",
              "unknownError",
              "unsupportedDevice",
              "notReady",
              "alreadyExists",
              "notConnected",
              "doesNotExist",
              "invalidArgs",
            ][result - 1]
          : undefined
      );
    },
    "try_RecordPairing": (
      retPtr: Pointer,
      errPtr: Pointer,
      transport: number,
      pairingDurationMs: number,
      result: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothPrivate.recordPairing(
          transport > 0 && transport <= 4 ? ["invalid", "classic", "le", "dual"][transport - 1] : undefined,
          pairingDurationMs,
          result > 0 && result <= 15
            ? [
                "alreadyConnected",
                "authCanceled",
                "authFailed",
                "authRejected",
                "authTimeout",
                "failed",
                "inProgress",
                "success",
                "unknownError",
                "unsupportedDevice",
                "notReady",
                "alreadyExists",
                "notConnected",
                "doesNotExist",
                "invalidArgs",
              ][result - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordReconnection": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothPrivate && "recordReconnection" in WEBEXT?.bluetoothPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordReconnection": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.recordReconnection);
    },
    "call_RecordReconnection": (retPtr: Pointer, result: number): void => {
      const _ret = WEBEXT.bluetoothPrivate.recordReconnection(
        result > 0 && result <= 15
          ? [
              "alreadyConnected",
              "authCanceled",
              "authFailed",
              "authRejected",
              "authTimeout",
              "failed",
              "inProgress",
              "success",
              "unknownError",
              "unsupportedDevice",
              "notReady",
              "alreadyExists",
              "notConnected",
              "doesNotExist",
              "invalidArgs",
            ][result - 1]
          : undefined
      );
    },
    "try_RecordReconnection": (retPtr: Pointer, errPtr: Pointer, result: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothPrivate.recordReconnection(
          result > 0 && result <= 15
            ? [
                "alreadyConnected",
                "authCanceled",
                "authFailed",
                "authRejected",
                "authTimeout",
                "failed",
                "inProgress",
                "success",
                "unknownError",
                "unsupportedDevice",
                "notReady",
                "alreadyExists",
                "notConnected",
                "doesNotExist",
                "invalidArgs",
              ][result - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetAdapterState": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothPrivate && "setAdapterState" in WEBEXT?.bluetoothPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetAdapterState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.setAdapterState);
    },
    "call_SetAdapterState": (retPtr: Pointer, adapterState: Pointer): void => {
      const adapterState_ffi = {};

      adapterState_ffi["name"] = A.load.Ref(adapterState + 0, undefined);
      if (A.load.Bool(adapterState + 6)) {
        adapterState_ffi["powered"] = A.load.Bool(adapterState + 4);
      }
      if (A.load.Bool(adapterState + 7)) {
        adapterState_ffi["discoverable"] = A.load.Bool(adapterState + 5);
      }

      const _ret = WEBEXT.bluetoothPrivate.setAdapterState(adapterState_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetAdapterState": (retPtr: Pointer, errPtr: Pointer, adapterState: Pointer): heap.Ref<boolean> => {
      try {
        const adapterState_ffi = {};

        adapterState_ffi["name"] = A.load.Ref(adapterState + 0, undefined);
        if (A.load.Bool(adapterState + 6)) {
          adapterState_ffi["powered"] = A.load.Bool(adapterState + 4);
        }
        if (A.load.Bool(adapterState + 7)) {
          adapterState_ffi["discoverable"] = A.load.Bool(adapterState + 5);
        }

        const _ret = WEBEXT.bluetoothPrivate.setAdapterState(adapterState_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetDiscoveryFilter": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothPrivate && "setDiscoveryFilter" in WEBEXT?.bluetoothPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetDiscoveryFilter": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.setDiscoveryFilter);
    },
    "call_SetDiscoveryFilter": (retPtr: Pointer, discoveryFilter: Pointer): void => {
      const discoveryFilter_ffi = {};

      discoveryFilter_ffi["transport"] = A.load.Enum(discoveryFilter + 0, ["le", "bredr", "dual"]);
      discoveryFilter_ffi["uuids"] = A.load.Ref(discoveryFilter + 4, undefined);
      if (A.load.Bool(discoveryFilter + 16)) {
        discoveryFilter_ffi["rssi"] = A.load.Int32(discoveryFilter + 8);
      }
      if (A.load.Bool(discoveryFilter + 17)) {
        discoveryFilter_ffi["pathloss"] = A.load.Int32(discoveryFilter + 12);
      }

      const _ret = WEBEXT.bluetoothPrivate.setDiscoveryFilter(discoveryFilter_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetDiscoveryFilter": (retPtr: Pointer, errPtr: Pointer, discoveryFilter: Pointer): heap.Ref<boolean> => {
      try {
        const discoveryFilter_ffi = {};

        discoveryFilter_ffi["transport"] = A.load.Enum(discoveryFilter + 0, ["le", "bredr", "dual"]);
        discoveryFilter_ffi["uuids"] = A.load.Ref(discoveryFilter + 4, undefined);
        if (A.load.Bool(discoveryFilter + 16)) {
          discoveryFilter_ffi["rssi"] = A.load.Int32(discoveryFilter + 8);
        }
        if (A.load.Bool(discoveryFilter + 17)) {
          discoveryFilter_ffi["pathloss"] = A.load.Int32(discoveryFilter + 12);
        }

        const _ret = WEBEXT.bluetoothPrivate.setDiscoveryFilter(discoveryFilter_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPairingResponse": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothPrivate && "setPairingResponse" in WEBEXT?.bluetoothPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPairingResponse": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothPrivate.setPairingResponse);
    },
    "call_SetPairingResponse": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 0 + 67)) {
        options_ffi["device"] = {};
        options_ffi["device"]["address"] = A.load.Ref(options + 0 + 0, undefined);
        options_ffi["device"]["name"] = A.load.Ref(options + 0 + 4, undefined);
        if (A.load.Bool(options + 0 + 56)) {
          options_ffi["device"]["deviceClass"] = A.load.Int32(options + 0 + 8);
        }
        options_ffi["device"]["vendorIdSource"] = A.load.Enum(options + 0 + 12, ["bluetooth", "usb"]);
        if (A.load.Bool(options + 0 + 57)) {
          options_ffi["device"]["vendorId"] = A.load.Int32(options + 0 + 16);
        }
        if (A.load.Bool(options + 0 + 58)) {
          options_ffi["device"]["productId"] = A.load.Int32(options + 0 + 20);
        }
        if (A.load.Bool(options + 0 + 59)) {
          options_ffi["device"]["deviceId"] = A.load.Int32(options + 0 + 24);
        }
        options_ffi["device"]["type"] = A.load.Enum(options + 0 + 28, [
          "computer",
          "phone",
          "modem",
          "audio",
          "carAudio",
          "video",
          "peripheral",
          "joystick",
          "gamepad",
          "keyboard",
          "mouse",
          "tablet",
          "keyboardMouseCombo",
        ]);
        if (A.load.Bool(options + 0 + 60)) {
          options_ffi["device"]["paired"] = A.load.Bool(options + 0 + 32);
        }
        if (A.load.Bool(options + 0 + 61)) {
          options_ffi["device"]["connected"] = A.load.Bool(options + 0 + 33);
        }
        if (A.load.Bool(options + 0 + 62)) {
          options_ffi["device"]["connecting"] = A.load.Bool(options + 0 + 34);
        }
        if (A.load.Bool(options + 0 + 63)) {
          options_ffi["device"]["connectable"] = A.load.Bool(options + 0 + 35);
        }
        options_ffi["device"]["uuids"] = A.load.Ref(options + 0 + 36, undefined);
        if (A.load.Bool(options + 0 + 64)) {
          options_ffi["device"]["inquiryRssi"] = A.load.Int32(options + 0 + 40);
        }
        if (A.load.Bool(options + 0 + 65)) {
          options_ffi["device"]["inquiryTxPower"] = A.load.Int32(options + 0 + 44);
        }
        options_ffi["device"]["transport"] = A.load.Enum(options + 0 + 48, ["invalid", "classic", "le", "dual"]);
        if (A.load.Bool(options + 0 + 66)) {
          options_ffi["device"]["batteryPercentage"] = A.load.Int32(options + 0 + 52);
        }
      }
      options_ffi["response"] = A.load.Enum(options + 68, ["confirm", "reject", "cancel"]);
      options_ffi["pincode"] = A.load.Ref(options + 72, undefined);
      if (A.load.Bool(options + 80)) {
        options_ffi["passkey"] = A.load.Int32(options + 76);
      }

      const _ret = WEBEXT.bluetoothPrivate.setPairingResponse(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetPairingResponse": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 0 + 67)) {
          options_ffi["device"] = {};
          options_ffi["device"]["address"] = A.load.Ref(options + 0 + 0, undefined);
          options_ffi["device"]["name"] = A.load.Ref(options + 0 + 4, undefined);
          if (A.load.Bool(options + 0 + 56)) {
            options_ffi["device"]["deviceClass"] = A.load.Int32(options + 0 + 8);
          }
          options_ffi["device"]["vendorIdSource"] = A.load.Enum(options + 0 + 12, ["bluetooth", "usb"]);
          if (A.load.Bool(options + 0 + 57)) {
            options_ffi["device"]["vendorId"] = A.load.Int32(options + 0 + 16);
          }
          if (A.load.Bool(options + 0 + 58)) {
            options_ffi["device"]["productId"] = A.load.Int32(options + 0 + 20);
          }
          if (A.load.Bool(options + 0 + 59)) {
            options_ffi["device"]["deviceId"] = A.load.Int32(options + 0 + 24);
          }
          options_ffi["device"]["type"] = A.load.Enum(options + 0 + 28, [
            "computer",
            "phone",
            "modem",
            "audio",
            "carAudio",
            "video",
            "peripheral",
            "joystick",
            "gamepad",
            "keyboard",
            "mouse",
            "tablet",
            "keyboardMouseCombo",
          ]);
          if (A.load.Bool(options + 0 + 60)) {
            options_ffi["device"]["paired"] = A.load.Bool(options + 0 + 32);
          }
          if (A.load.Bool(options + 0 + 61)) {
            options_ffi["device"]["connected"] = A.load.Bool(options + 0 + 33);
          }
          if (A.load.Bool(options + 0 + 62)) {
            options_ffi["device"]["connecting"] = A.load.Bool(options + 0 + 34);
          }
          if (A.load.Bool(options + 0 + 63)) {
            options_ffi["device"]["connectable"] = A.load.Bool(options + 0 + 35);
          }
          options_ffi["device"]["uuids"] = A.load.Ref(options + 0 + 36, undefined);
          if (A.load.Bool(options + 0 + 64)) {
            options_ffi["device"]["inquiryRssi"] = A.load.Int32(options + 0 + 40);
          }
          if (A.load.Bool(options + 0 + 65)) {
            options_ffi["device"]["inquiryTxPower"] = A.load.Int32(options + 0 + 44);
          }
          options_ffi["device"]["transport"] = A.load.Enum(options + 0 + 48, ["invalid", "classic", "le", "dual"]);
          if (A.load.Bool(options + 0 + 66)) {
            options_ffi["device"]["batteryPercentage"] = A.load.Int32(options + 0 + 52);
          }
        }
        options_ffi["response"] = A.load.Enum(options + 68, ["confirm", "reject", "cancel"]);
        options_ffi["pincode"] = A.load.Ref(options + 72, undefined);
        if (A.load.Bool(options + 80)) {
          options_ffi["passkey"] = A.load.Int32(options + 76);
        }

        const _ret = WEBEXT.bluetoothPrivate.setPairingResponse(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
