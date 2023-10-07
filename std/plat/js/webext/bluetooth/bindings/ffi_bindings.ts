import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/bluetooth", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AdapterState": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 10, false);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Ref(ptr + 0, x["address"]);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Bool(ptr + 11, "powered" in x ? true : false);
        A.store.Bool(ptr + 8, x["powered"] ? true : false);
        A.store.Bool(ptr + 12, "available" in x ? true : false);
        A.store.Bool(ptr + 9, x["available"] ? true : false);
        A.store.Bool(ptr + 13, "discovering" in x ? true : false);
        A.store.Bool(ptr + 10, x["discovering"] ? true : false);
      }
    },
    "load_AdapterState": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["address"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 11)) {
        x["powered"] = A.load.Bool(ptr + 8);
      } else {
        delete x["powered"];
      }
      if (A.load.Bool(ptr + 12)) {
        x["available"] = A.load.Bool(ptr + 9);
      } else {
        delete x["available"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["discovering"] = A.load.Bool(ptr + 10);
      } else {
        delete x["discovering"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_FilterType": (ref: heap.Ref<string>): number => {
      const idx = ["all", "known"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_BluetoothFilter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Enum(ptr + 0, ["all", "known"].indexOf(x["filterType"] as string));
        A.store.Bool(ptr + 8, "limit" in x ? true : false);
        A.store.Int32(ptr + 4, x["limit"] === undefined ? 0 : (x["limit"] as number));
      }
    },
    "load_BluetoothFilter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["filterType"] = A.load.Enum(ptr + 0, ["all", "known"]);
      if (A.load.Bool(ptr + 8)) {
        x["limit"] = A.load.Int32(ptr + 4);
      } else {
        delete x["limit"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_VendorIdSource": (ref: heap.Ref<string>): number => {
      const idx = ["bluetooth", "usb"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_DeviceType": (ref: heap.Ref<string>): number => {
      const idx = [
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
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_Transport": (ref: heap.Ref<string>): number => {
      const idx = ["invalid", "classic", "le", "dual"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Device": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 67, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 56, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Enum(ptr + 12, -1);
        A.store.Bool(ptr + 57, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Bool(ptr + 58, false);
        A.store.Int32(ptr + 20, 0);
        A.store.Bool(ptr + 59, false);
        A.store.Int32(ptr + 24, 0);
        A.store.Enum(ptr + 28, -1);
        A.store.Bool(ptr + 60, false);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 61, false);
        A.store.Bool(ptr + 33, false);
        A.store.Bool(ptr + 62, false);
        A.store.Bool(ptr + 34, false);
        A.store.Bool(ptr + 63, false);
        A.store.Bool(ptr + 35, false);
        A.store.Ref(ptr + 36, undefined);
        A.store.Bool(ptr + 64, false);
        A.store.Int32(ptr + 40, 0);
        A.store.Bool(ptr + 65, false);
        A.store.Int32(ptr + 44, 0);
        A.store.Enum(ptr + 48, -1);
        A.store.Bool(ptr + 66, false);
        A.store.Int32(ptr + 52, 0);
      } else {
        A.store.Bool(ptr + 67, true);
        A.store.Ref(ptr + 0, x["address"]);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Bool(ptr + 56, "deviceClass" in x ? true : false);
        A.store.Int32(ptr + 8, x["deviceClass"] === undefined ? 0 : (x["deviceClass"] as number));
        A.store.Enum(ptr + 12, ["bluetooth", "usb"].indexOf(x["vendorIdSource"] as string));
        A.store.Bool(ptr + 57, "vendorId" in x ? true : false);
        A.store.Int32(ptr + 16, x["vendorId"] === undefined ? 0 : (x["vendorId"] as number));
        A.store.Bool(ptr + 58, "productId" in x ? true : false);
        A.store.Int32(ptr + 20, x["productId"] === undefined ? 0 : (x["productId"] as number));
        A.store.Bool(ptr + 59, "deviceId" in x ? true : false);
        A.store.Int32(ptr + 24, x["deviceId"] === undefined ? 0 : (x["deviceId"] as number));
        A.store.Enum(
          ptr + 28,
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
          ].indexOf(x["type"] as string)
        );
        A.store.Bool(ptr + 60, "paired" in x ? true : false);
        A.store.Bool(ptr + 32, x["paired"] ? true : false);
        A.store.Bool(ptr + 61, "connected" in x ? true : false);
        A.store.Bool(ptr + 33, x["connected"] ? true : false);
        A.store.Bool(ptr + 62, "connecting" in x ? true : false);
        A.store.Bool(ptr + 34, x["connecting"] ? true : false);
        A.store.Bool(ptr + 63, "connectable" in x ? true : false);
        A.store.Bool(ptr + 35, x["connectable"] ? true : false);
        A.store.Ref(ptr + 36, x["uuids"]);
        A.store.Bool(ptr + 64, "inquiryRssi" in x ? true : false);
        A.store.Int32(ptr + 40, x["inquiryRssi"] === undefined ? 0 : (x["inquiryRssi"] as number));
        A.store.Bool(ptr + 65, "inquiryTxPower" in x ? true : false);
        A.store.Int32(ptr + 44, x["inquiryTxPower"] === undefined ? 0 : (x["inquiryTxPower"] as number));
        A.store.Enum(ptr + 48, ["invalid", "classic", "le", "dual"].indexOf(x["transport"] as string));
        A.store.Bool(ptr + 66, "batteryPercentage" in x ? true : false);
        A.store.Int32(ptr + 52, x["batteryPercentage"] === undefined ? 0 : (x["batteryPercentage"] as number));
      }
    },
    "load_Device": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["address"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 56)) {
        x["deviceClass"] = A.load.Int32(ptr + 8);
      } else {
        delete x["deviceClass"];
      }
      x["vendorIdSource"] = A.load.Enum(ptr + 12, ["bluetooth", "usb"]);
      if (A.load.Bool(ptr + 57)) {
        x["vendorId"] = A.load.Int32(ptr + 16);
      } else {
        delete x["vendorId"];
      }
      if (A.load.Bool(ptr + 58)) {
        x["productId"] = A.load.Int32(ptr + 20);
      } else {
        delete x["productId"];
      }
      if (A.load.Bool(ptr + 59)) {
        x["deviceId"] = A.load.Int32(ptr + 24);
      } else {
        delete x["deviceId"];
      }
      x["type"] = A.load.Enum(ptr + 28, [
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
      if (A.load.Bool(ptr + 60)) {
        x["paired"] = A.load.Bool(ptr + 32);
      } else {
        delete x["paired"];
      }
      if (A.load.Bool(ptr + 61)) {
        x["connected"] = A.load.Bool(ptr + 33);
      } else {
        delete x["connected"];
      }
      if (A.load.Bool(ptr + 62)) {
        x["connecting"] = A.load.Bool(ptr + 34);
      } else {
        delete x["connecting"];
      }
      if (A.load.Bool(ptr + 63)) {
        x["connectable"] = A.load.Bool(ptr + 35);
      } else {
        delete x["connectable"];
      }
      x["uuids"] = A.load.Ref(ptr + 36, undefined);
      if (A.load.Bool(ptr + 64)) {
        x["inquiryRssi"] = A.load.Int32(ptr + 40);
      } else {
        delete x["inquiryRssi"];
      }
      if (A.load.Bool(ptr + 65)) {
        x["inquiryTxPower"] = A.load.Int32(ptr + 44);
      } else {
        delete x["inquiryTxPower"];
      }
      x["transport"] = A.load.Enum(ptr + 48, ["invalid", "classic", "le", "dual"]);
      if (A.load.Bool(ptr + 66)) {
        x["batteryPercentage"] = A.load.Int32(ptr + 52);
      } else {
        delete x["batteryPercentage"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetAdapterState": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth && "getAdapterState" in WEBEXT?.bluetooth) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAdapterState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.getAdapterState);
    },
    "call_GetAdapterState": (retPtr: Pointer): void => {
      const _ret = WEBEXT.bluetooth.getAdapterState();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAdapterState": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.getAdapterState();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDevice": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth && "getDevice" in WEBEXT?.bluetooth) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDevice": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.getDevice);
    },
    "call_GetDevice": (retPtr: Pointer, deviceAddress: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetooth.getDevice(A.H.get<object>(deviceAddress));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDevice": (retPtr: Pointer, errPtr: Pointer, deviceAddress: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.getDevice(A.H.get<object>(deviceAddress));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDevices": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth && "getDevices" in WEBEXT?.bluetooth) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDevices": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.getDevices);
    },
    "call_GetDevices": (retPtr: Pointer, filter: Pointer): void => {
      const filter_ffi = {};

      filter_ffi["filterType"] = A.load.Enum(filter + 0, ["all", "known"]);
      if (A.load.Bool(filter + 8)) {
        filter_ffi["limit"] = A.load.Int32(filter + 4);
      }

      const _ret = WEBEXT.bluetooth.getDevices(filter_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDevices": (retPtr: Pointer, errPtr: Pointer, filter: Pointer): heap.Ref<boolean> => {
      try {
        const filter_ffi = {};

        filter_ffi["filterType"] = A.load.Enum(filter + 0, ["all", "known"]);
        if (A.load.Bool(filter + 8)) {
          filter_ffi["limit"] = A.load.Int32(filter + 4);
        }

        const _ret = WEBEXT.bluetooth.getDevices(filter_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAdapterStateChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth?.onAdapterStateChanged && "addListener" in WEBEXT?.bluetooth?.onAdapterStateChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAdapterStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.onAdapterStateChanged.addListener);
    },
    "call_OnAdapterStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetooth.onAdapterStateChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnAdapterStateChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.onAdapterStateChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAdapterStateChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth?.onAdapterStateChanged && "removeListener" in WEBEXT?.bluetooth?.onAdapterStateChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAdapterStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.onAdapterStateChanged.removeListener);
    },
    "call_OffAdapterStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetooth.onAdapterStateChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffAdapterStateChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.onAdapterStateChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAdapterStateChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth?.onAdapterStateChanged && "hasListener" in WEBEXT?.bluetooth?.onAdapterStateChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAdapterStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.onAdapterStateChanged.hasListener);
    },
    "call_HasOnAdapterStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetooth.onAdapterStateChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAdapterStateChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.onAdapterStateChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeviceAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth?.onDeviceAdded && "addListener" in WEBEXT?.bluetooth?.onDeviceAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeviceAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.onDeviceAdded.addListener);
    },
    "call_OnDeviceAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetooth.onDeviceAdded.addListener(A.H.get<object>(callback));
    },
    "try_OnDeviceAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.onDeviceAdded.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeviceAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth?.onDeviceAdded && "removeListener" in WEBEXT?.bluetooth?.onDeviceAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeviceAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.onDeviceAdded.removeListener);
    },
    "call_OffDeviceAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetooth.onDeviceAdded.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeviceAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.onDeviceAdded.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeviceAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth?.onDeviceAdded && "hasListener" in WEBEXT?.bluetooth?.onDeviceAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeviceAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.onDeviceAdded.hasListener);
    },
    "call_HasOnDeviceAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetooth.onDeviceAdded.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeviceAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.onDeviceAdded.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeviceChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth?.onDeviceChanged && "addListener" in WEBEXT?.bluetooth?.onDeviceChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeviceChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.onDeviceChanged.addListener);
    },
    "call_OnDeviceChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetooth.onDeviceChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnDeviceChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.onDeviceChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeviceChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth?.onDeviceChanged && "removeListener" in WEBEXT?.bluetooth?.onDeviceChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeviceChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.onDeviceChanged.removeListener);
    },
    "call_OffDeviceChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetooth.onDeviceChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeviceChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.onDeviceChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeviceChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth?.onDeviceChanged && "hasListener" in WEBEXT?.bluetooth?.onDeviceChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeviceChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.onDeviceChanged.hasListener);
    },
    "call_HasOnDeviceChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetooth.onDeviceChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeviceChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.onDeviceChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeviceRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth?.onDeviceRemoved && "addListener" in WEBEXT?.bluetooth?.onDeviceRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeviceRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.onDeviceRemoved.addListener);
    },
    "call_OnDeviceRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetooth.onDeviceRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnDeviceRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.onDeviceRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeviceRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth?.onDeviceRemoved && "removeListener" in WEBEXT?.bluetooth?.onDeviceRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeviceRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.onDeviceRemoved.removeListener);
    },
    "call_OffDeviceRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetooth.onDeviceRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeviceRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.onDeviceRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeviceRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth?.onDeviceRemoved && "hasListener" in WEBEXT?.bluetooth?.onDeviceRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeviceRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.onDeviceRemoved.hasListener);
    },
    "call_HasOnDeviceRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetooth.onDeviceRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeviceRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.onDeviceRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartDiscovery": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth && "startDiscovery" in WEBEXT?.bluetooth) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartDiscovery": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.startDiscovery);
    },
    "call_StartDiscovery": (retPtr: Pointer): void => {
      const _ret = WEBEXT.bluetooth.startDiscovery();
      A.store.Ref(retPtr, _ret);
    },
    "try_StartDiscovery": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.startDiscovery();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StopDiscovery": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetooth && "stopDiscovery" in WEBEXT?.bluetooth) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StopDiscovery": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetooth.stopDiscovery);
    },
    "call_StopDiscovery": (retPtr: Pointer): void => {
      const _ret = WEBEXT.bluetooth.stopDiscovery();
      A.store.Ref(retPtr, _ret);
    },
    "try_StopDiscovery": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetooth.stopDiscovery();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
