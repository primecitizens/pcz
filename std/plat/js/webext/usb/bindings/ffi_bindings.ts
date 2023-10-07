import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/usb", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_TransferType": (ref: heap.Ref<string>): number => {
      const idx = ["control", "interrupt", "isochronous", "bulk"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_Direction": (ref: heap.Ref<string>): number => {
      const idx = ["in", "out"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SynchronizationType": (ref: heap.Ref<string>): number => {
      const idx = ["asynchronous", "adaptive", "synchronous"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_UsageType": (ref: heap.Ref<string>): number => {
      const idx = ["data", "feedback", "explicitFeedback", "periodic", "notification"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_EndpointDescriptor": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 35, false);
        A.store.Bool(ptr + 32, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Enum(ptr + 4, -1);
        A.store.Enum(ptr + 8, -1);
        A.store.Bool(ptr + 33, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Enum(ptr + 20, -1);
        A.store.Bool(ptr + 34, false);
        A.store.Int32(ptr + 24, 0);
        A.store.Ref(ptr + 28, undefined);
      } else {
        A.store.Bool(ptr + 35, true);
        A.store.Bool(ptr + 32, "address" in x ? true : false);
        A.store.Int32(ptr + 0, x["address"] === undefined ? 0 : (x["address"] as number));
        A.store.Enum(ptr + 4, ["control", "interrupt", "isochronous", "bulk"].indexOf(x["type"] as string));
        A.store.Enum(ptr + 8, ["in", "out"].indexOf(x["direction"] as string));
        A.store.Bool(ptr + 33, "maximumPacketSize" in x ? true : false);
        A.store.Int32(ptr + 12, x["maximumPacketSize"] === undefined ? 0 : (x["maximumPacketSize"] as number));
        A.store.Enum(ptr + 16, ["asynchronous", "adaptive", "synchronous"].indexOf(x["synchronization"] as string));
        A.store.Enum(
          ptr + 20,
          ["data", "feedback", "explicitFeedback", "periodic", "notification"].indexOf(x["usage"] as string)
        );
        A.store.Bool(ptr + 34, "pollingInterval" in x ? true : false);
        A.store.Int32(ptr + 24, x["pollingInterval"] === undefined ? 0 : (x["pollingInterval"] as number));
        A.store.Ref(ptr + 28, x["extra_data"]);
      }
    },
    "load_EndpointDescriptor": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 32)) {
        x["address"] = A.load.Int32(ptr + 0);
      } else {
        delete x["address"];
      }
      x["type"] = A.load.Enum(ptr + 4, ["control", "interrupt", "isochronous", "bulk"]);
      x["direction"] = A.load.Enum(ptr + 8, ["in", "out"]);
      if (A.load.Bool(ptr + 33)) {
        x["maximumPacketSize"] = A.load.Int32(ptr + 12);
      } else {
        delete x["maximumPacketSize"];
      }
      x["synchronization"] = A.load.Enum(ptr + 16, ["asynchronous", "adaptive", "synchronous"]);
      x["usage"] = A.load.Enum(ptr + 20, ["data", "feedback", "explicitFeedback", "periodic", "notification"]);
      if (A.load.Bool(ptr + 34)) {
        x["pollingInterval"] = A.load.Int32(ptr + 24);
      } else {
        delete x["pollingInterval"];
      }
      x["extra_data"] = A.load.Ref(ptr + 28, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_InterfaceDescriptor": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 37, false);
        A.store.Bool(ptr + 32, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 33, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 34, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 35, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 36, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
      } else {
        A.store.Bool(ptr + 37, true);
        A.store.Bool(ptr + 32, "interfaceNumber" in x ? true : false);
        A.store.Int32(ptr + 0, x["interfaceNumber"] === undefined ? 0 : (x["interfaceNumber"] as number));
        A.store.Bool(ptr + 33, "alternateSetting" in x ? true : false);
        A.store.Int32(ptr + 4, x["alternateSetting"] === undefined ? 0 : (x["alternateSetting"] as number));
        A.store.Bool(ptr + 34, "interfaceClass" in x ? true : false);
        A.store.Int32(ptr + 8, x["interfaceClass"] === undefined ? 0 : (x["interfaceClass"] as number));
        A.store.Bool(ptr + 35, "interfaceSubclass" in x ? true : false);
        A.store.Int32(ptr + 12, x["interfaceSubclass"] === undefined ? 0 : (x["interfaceSubclass"] as number));
        A.store.Bool(ptr + 36, "interfaceProtocol" in x ? true : false);
        A.store.Int32(ptr + 16, x["interfaceProtocol"] === undefined ? 0 : (x["interfaceProtocol"] as number));
        A.store.Ref(ptr + 20, x["description"]);
        A.store.Ref(ptr + 24, x["endpoints"]);
        A.store.Ref(ptr + 28, x["extra_data"]);
      }
    },
    "load_InterfaceDescriptor": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 32)) {
        x["interfaceNumber"] = A.load.Int32(ptr + 0);
      } else {
        delete x["interfaceNumber"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["alternateSetting"] = A.load.Int32(ptr + 4);
      } else {
        delete x["alternateSetting"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["interfaceClass"] = A.load.Int32(ptr + 8);
      } else {
        delete x["interfaceClass"];
      }
      if (A.load.Bool(ptr + 35)) {
        x["interfaceSubclass"] = A.load.Int32(ptr + 12);
      } else {
        delete x["interfaceSubclass"];
      }
      if (A.load.Bool(ptr + 36)) {
        x["interfaceProtocol"] = A.load.Int32(ptr + 16);
      } else {
        delete x["interfaceProtocol"];
      }
      x["description"] = A.load.Ref(ptr + 20, undefined);
      x["endpoints"] = A.load.Ref(ptr + 24, undefined);
      x["extra_data"] = A.load.Ref(ptr + 28, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ConfigDescriptor": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 33, false);
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 29, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 30, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 31, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 32, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 33, true);
        A.store.Bool(ptr + 28, "active" in x ? true : false);
        A.store.Bool(ptr + 0, x["active"] ? true : false);
        A.store.Bool(ptr + 29, "configurationValue" in x ? true : false);
        A.store.Int32(ptr + 4, x["configurationValue"] === undefined ? 0 : (x["configurationValue"] as number));
        A.store.Ref(ptr + 8, x["description"]);
        A.store.Bool(ptr + 30, "selfPowered" in x ? true : false);
        A.store.Bool(ptr + 12, x["selfPowered"] ? true : false);
        A.store.Bool(ptr + 31, "remoteWakeup" in x ? true : false);
        A.store.Bool(ptr + 13, x["remoteWakeup"] ? true : false);
        A.store.Bool(ptr + 32, "maxPower" in x ? true : false);
        A.store.Int32(ptr + 16, x["maxPower"] === undefined ? 0 : (x["maxPower"] as number));
        A.store.Ref(ptr + 20, x["interfaces"]);
        A.store.Ref(ptr + 24, x["extra_data"]);
      }
    },
    "load_ConfigDescriptor": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 28)) {
        x["active"] = A.load.Bool(ptr + 0);
      } else {
        delete x["active"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["configurationValue"] = A.load.Int32(ptr + 4);
      } else {
        delete x["configurationValue"];
      }
      x["description"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 30)) {
        x["selfPowered"] = A.load.Bool(ptr + 12);
      } else {
        delete x["selfPowered"];
      }
      if (A.load.Bool(ptr + 31)) {
        x["remoteWakeup"] = A.load.Bool(ptr + 13);
      } else {
        delete x["remoteWakeup"];
      }
      if (A.load.Bool(ptr + 32)) {
        x["maxPower"] = A.load.Int32(ptr + 16);
      } else {
        delete x["maxPower"];
      }
      x["interfaces"] = A.load.Ref(ptr + 20, undefined);
      x["extra_data"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ConnectionHandle": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Bool(ptr + 12, "handle" in x ? true : false);
        A.store.Int32(ptr + 0, x["handle"] === undefined ? 0 : (x["handle"] as number));
        A.store.Bool(ptr + 13, "vendorId" in x ? true : false);
        A.store.Int32(ptr + 4, x["vendorId"] === undefined ? 0 : (x["vendorId"] as number));
        A.store.Bool(ptr + 14, "productId" in x ? true : false);
        A.store.Int32(ptr + 8, x["productId"] === undefined ? 0 : (x["productId"] as number));
      }
    },
    "load_ConnectionHandle": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["handle"] = A.load.Int32(ptr + 0);
      } else {
        delete x["handle"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["vendorId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["vendorId"];
      }
      if (A.load.Bool(ptr + 14)) {
        x["productId"] = A.load.Int32(ptr + 8);
      } else {
        delete x["productId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_Recipient": (ref: heap.Ref<string>): number => {
      const idx = ["device", "_interface", "endpoint", "other"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_RequestType": (ref: heap.Ref<string>): number => {
      const idx = ["standard", "class", "vendor", "reserved"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ControlTransferInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 41, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Enum(ptr + 4, -1);
        A.store.Enum(ptr + 8, -1);
        A.store.Bool(ptr + 36, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 37, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Bool(ptr + 38, false);
        A.store.Int32(ptr + 20, 0);
        A.store.Bool(ptr + 39, false);
        A.store.Int32(ptr + 24, 0);
        A.store.Ref(ptr + 28, undefined);
        A.store.Bool(ptr + 40, false);
        A.store.Int32(ptr + 32, 0);
      } else {
        A.store.Bool(ptr + 41, true);
        A.store.Enum(ptr + 0, ["in", "out"].indexOf(x["direction"] as string));
        A.store.Enum(ptr + 4, ["device", "_interface", "endpoint", "other"].indexOf(x["recipient"] as string));
        A.store.Enum(ptr + 8, ["standard", "class", "vendor", "reserved"].indexOf(x["requestType"] as string));
        A.store.Bool(ptr + 36, "request" in x ? true : false);
        A.store.Int32(ptr + 12, x["request"] === undefined ? 0 : (x["request"] as number));
        A.store.Bool(ptr + 37, "value" in x ? true : false);
        A.store.Int32(ptr + 16, x["value"] === undefined ? 0 : (x["value"] as number));
        A.store.Bool(ptr + 38, "index" in x ? true : false);
        A.store.Int32(ptr + 20, x["index"] === undefined ? 0 : (x["index"] as number));
        A.store.Bool(ptr + 39, "length" in x ? true : false);
        A.store.Int32(ptr + 24, x["length"] === undefined ? 0 : (x["length"] as number));
        A.store.Ref(ptr + 28, x["data"]);
        A.store.Bool(ptr + 40, "timeout" in x ? true : false);
        A.store.Int32(ptr + 32, x["timeout"] === undefined ? 0 : (x["timeout"] as number));
      }
    },
    "load_ControlTransferInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["direction"] = A.load.Enum(ptr + 0, ["in", "out"]);
      x["recipient"] = A.load.Enum(ptr + 4, ["device", "_interface", "endpoint", "other"]);
      x["requestType"] = A.load.Enum(ptr + 8, ["standard", "class", "vendor", "reserved"]);
      if (A.load.Bool(ptr + 36)) {
        x["request"] = A.load.Int32(ptr + 12);
      } else {
        delete x["request"];
      }
      if (A.load.Bool(ptr + 37)) {
        x["value"] = A.load.Int32(ptr + 16);
      } else {
        delete x["value"];
      }
      if (A.load.Bool(ptr + 38)) {
        x["index"] = A.load.Int32(ptr + 20);
      } else {
        delete x["index"];
      }
      if (A.load.Bool(ptr + 39)) {
        x["length"] = A.load.Int32(ptr + 24);
      } else {
        delete x["length"];
      }
      x["data"] = A.load.Ref(ptr + 28, undefined);
      if (A.load.Bool(ptr + 40)) {
        x["timeout"] = A.load.Int32(ptr + 32);
      } else {
        delete x["timeout"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Device": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 28, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 29, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 30, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 31, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 32, true);
        A.store.Bool(ptr + 28, "device" in x ? true : false);
        A.store.Int32(ptr + 0, x["device"] === undefined ? 0 : (x["device"] as number));
        A.store.Bool(ptr + 29, "vendorId" in x ? true : false);
        A.store.Int32(ptr + 4, x["vendorId"] === undefined ? 0 : (x["vendorId"] as number));
        A.store.Bool(ptr + 30, "productId" in x ? true : false);
        A.store.Int32(ptr + 8, x["productId"] === undefined ? 0 : (x["productId"] as number));
        A.store.Bool(ptr + 31, "version" in x ? true : false);
        A.store.Int32(ptr + 12, x["version"] === undefined ? 0 : (x["version"] as number));
        A.store.Ref(ptr + 16, x["productName"]);
        A.store.Ref(ptr + 20, x["manufacturerName"]);
        A.store.Ref(ptr + 24, x["serialNumber"]);
      }
    },
    "load_Device": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 28)) {
        x["device"] = A.load.Int32(ptr + 0);
      } else {
        delete x["device"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["vendorId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["vendorId"];
      }
      if (A.load.Bool(ptr + 30)) {
        x["productId"] = A.load.Int32(ptr + 8);
      } else {
        delete x["productId"];
      }
      if (A.load.Bool(ptr + 31)) {
        x["version"] = A.load.Int32(ptr + 12);
      } else {
        delete x["version"];
      }
      x["productName"] = A.load.Ref(ptr + 16, undefined);
      x["manufacturerName"] = A.load.Ref(ptr + 20, undefined);
      x["serialNumber"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DeviceFilter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 20, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 21, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 22, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 23, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 24, false);
        A.store.Int32(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 25, true);
        A.store.Bool(ptr + 20, "vendorId" in x ? true : false);
        A.store.Int32(ptr + 0, x["vendorId"] === undefined ? 0 : (x["vendorId"] as number));
        A.store.Bool(ptr + 21, "productId" in x ? true : false);
        A.store.Int32(ptr + 4, x["productId"] === undefined ? 0 : (x["productId"] as number));
        A.store.Bool(ptr + 22, "interfaceClass" in x ? true : false);
        A.store.Int32(ptr + 8, x["interfaceClass"] === undefined ? 0 : (x["interfaceClass"] as number));
        A.store.Bool(ptr + 23, "interfaceSubclass" in x ? true : false);
        A.store.Int32(ptr + 12, x["interfaceSubclass"] === undefined ? 0 : (x["interfaceSubclass"] as number));
        A.store.Bool(ptr + 24, "interfaceProtocol" in x ? true : false);
        A.store.Int32(ptr + 16, x["interfaceProtocol"] === undefined ? 0 : (x["interfaceProtocol"] as number));
      }
    },
    "load_DeviceFilter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 20)) {
        x["vendorId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["vendorId"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["productId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["productId"];
      }
      if (A.load.Bool(ptr + 22)) {
        x["interfaceClass"] = A.load.Int32(ptr + 8);
      } else {
        delete x["interfaceClass"];
      }
      if (A.load.Bool(ptr + 23)) {
        x["interfaceSubclass"] = A.load.Int32(ptr + 12);
      } else {
        delete x["interfaceSubclass"];
      }
      if (A.load.Bool(ptr + 24)) {
        x["interfaceProtocol"] = A.load.Int32(ptr + 16);
      } else {
        delete x["interfaceProtocol"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DevicePromptOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "multiple" in x ? true : false);
        A.store.Bool(ptr + 0, x["multiple"] ? true : false);
        A.store.Ref(ptr + 4, x["filters"]);
      }
    },
    "load_DevicePromptOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["multiple"] = A.load.Bool(ptr + 0);
      } else {
        delete x["multiple"];
      }
      x["filters"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_EnumerateDevicesAndRequestAccessOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Bool(ptr + 12, "vendorId" in x ? true : false);
        A.store.Int32(ptr + 0, x["vendorId"] === undefined ? 0 : (x["vendorId"] as number));
        A.store.Bool(ptr + 13, "productId" in x ? true : false);
        A.store.Int32(ptr + 4, x["productId"] === undefined ? 0 : (x["productId"] as number));
        A.store.Bool(ptr + 14, "interfaceId" in x ? true : false);
        A.store.Int32(ptr + 8, x["interfaceId"] === undefined ? 0 : (x["interfaceId"] as number));
      }
    },
    "load_EnumerateDevicesAndRequestAccessOptions": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["vendorId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["vendorId"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["productId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["productId"];
      }
      if (A.load.Bool(ptr + 14)) {
        x["interfaceId"] = A.load.Int32(ptr + 8);
      } else {
        delete x["interfaceId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_EnumerateDevicesOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Bool(ptr + 12, "vendorId" in x ? true : false);
        A.store.Int32(ptr + 0, x["vendorId"] === undefined ? 0 : (x["vendorId"] as number));
        A.store.Bool(ptr + 13, "productId" in x ? true : false);
        A.store.Int32(ptr + 4, x["productId"] === undefined ? 0 : (x["productId"] as number));
        A.store.Ref(ptr + 8, x["filters"]);
      }
    },
    "load_EnumerateDevicesOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["vendorId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["vendorId"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["productId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["productId"];
      }
      x["filters"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GenericTransferInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 23, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 20, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 21, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 22, false);
        A.store.Int32(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 23, true);
        A.store.Enum(ptr + 0, ["in", "out"].indexOf(x["direction"] as string));
        A.store.Bool(ptr + 20, "endpoint" in x ? true : false);
        A.store.Int32(ptr + 4, x["endpoint"] === undefined ? 0 : (x["endpoint"] as number));
        A.store.Bool(ptr + 21, "length" in x ? true : false);
        A.store.Int32(ptr + 8, x["length"] === undefined ? 0 : (x["length"] as number));
        A.store.Ref(ptr + 12, x["data"]);
        A.store.Bool(ptr + 22, "timeout" in x ? true : false);
        A.store.Int32(ptr + 16, x["timeout"] === undefined ? 0 : (x["timeout"] as number));
      }
    },
    "load_GenericTransferInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["direction"] = A.load.Enum(ptr + 0, ["in", "out"]);
      if (A.load.Bool(ptr + 20)) {
        x["endpoint"] = A.load.Int32(ptr + 4);
      } else {
        delete x["endpoint"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["length"] = A.load.Int32(ptr + 8);
      } else {
        delete x["length"];
      }
      x["data"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 22)) {
        x["timeout"] = A.load.Int32(ptr + 16);
      } else {
        delete x["timeout"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_IsochronousTransferInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 34, false);

        A.store.Bool(ptr + 0 + 23, false);
        A.store.Enum(ptr + 0 + 0, -1);
        A.store.Bool(ptr + 0 + 20, false);
        A.store.Int32(ptr + 0 + 4, 0);
        A.store.Bool(ptr + 0 + 21, false);
        A.store.Int32(ptr + 0 + 8, 0);
        A.store.Ref(ptr + 0 + 12, undefined);
        A.store.Bool(ptr + 0 + 22, false);
        A.store.Int32(ptr + 0 + 16, 0);
        A.store.Bool(ptr + 32, false);
        A.store.Int32(ptr + 24, 0);
        A.store.Bool(ptr + 33, false);
        A.store.Int32(ptr + 28, 0);
      } else {
        A.store.Bool(ptr + 34, true);

        if (typeof x["transferInfo"] === "undefined") {
          A.store.Bool(ptr + 0 + 23, false);
          A.store.Enum(ptr + 0 + 0, -1);
          A.store.Bool(ptr + 0 + 20, false);
          A.store.Int32(ptr + 0 + 4, 0);
          A.store.Bool(ptr + 0 + 21, false);
          A.store.Int32(ptr + 0 + 8, 0);
          A.store.Ref(ptr + 0 + 12, undefined);
          A.store.Bool(ptr + 0 + 22, false);
          A.store.Int32(ptr + 0 + 16, 0);
        } else {
          A.store.Bool(ptr + 0 + 23, true);
          A.store.Enum(ptr + 0 + 0, ["in", "out"].indexOf(x["transferInfo"]["direction"] as string));
          A.store.Bool(ptr + 0 + 20, "endpoint" in x["transferInfo"] ? true : false);
          A.store.Int32(
            ptr + 0 + 4,
            x["transferInfo"]["endpoint"] === undefined ? 0 : (x["transferInfo"]["endpoint"] as number)
          );
          A.store.Bool(ptr + 0 + 21, "length" in x["transferInfo"] ? true : false);
          A.store.Int32(
            ptr + 0 + 8,
            x["transferInfo"]["length"] === undefined ? 0 : (x["transferInfo"]["length"] as number)
          );
          A.store.Ref(ptr + 0 + 12, x["transferInfo"]["data"]);
          A.store.Bool(ptr + 0 + 22, "timeout" in x["transferInfo"] ? true : false);
          A.store.Int32(
            ptr + 0 + 16,
            x["transferInfo"]["timeout"] === undefined ? 0 : (x["transferInfo"]["timeout"] as number)
          );
        }
        A.store.Bool(ptr + 32, "packets" in x ? true : false);
        A.store.Int32(ptr + 24, x["packets"] === undefined ? 0 : (x["packets"] as number));
        A.store.Bool(ptr + 33, "packetLength" in x ? true : false);
        A.store.Int32(ptr + 28, x["packetLength"] === undefined ? 0 : (x["packetLength"] as number));
      }
    },
    "load_IsochronousTransferInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 23)) {
        x["transferInfo"] = {};
        x["transferInfo"]["direction"] = A.load.Enum(ptr + 0 + 0, ["in", "out"]);
        if (A.load.Bool(ptr + 0 + 20)) {
          x["transferInfo"]["endpoint"] = A.load.Int32(ptr + 0 + 4);
        } else {
          delete x["transferInfo"]["endpoint"];
        }
        if (A.load.Bool(ptr + 0 + 21)) {
          x["transferInfo"]["length"] = A.load.Int32(ptr + 0 + 8);
        } else {
          delete x["transferInfo"]["length"];
        }
        x["transferInfo"]["data"] = A.load.Ref(ptr + 0 + 12, undefined);
        if (A.load.Bool(ptr + 0 + 22)) {
          x["transferInfo"]["timeout"] = A.load.Int32(ptr + 0 + 16);
        } else {
          delete x["transferInfo"]["timeout"];
        }
      } else {
        delete x["transferInfo"];
      }
      if (A.load.Bool(ptr + 32)) {
        x["packets"] = A.load.Int32(ptr + 24);
      } else {
        delete x["packets"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["packetLength"] = A.load.Int32(ptr + 28);
      } else {
        delete x["packetLength"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TransferResultInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "resultCode" in x ? true : false);
        A.store.Int32(ptr + 0, x["resultCode"] === undefined ? 0 : (x["resultCode"] as number));
        A.store.Ref(ptr + 4, x["data"]);
      }
    },
    "load_TransferResultInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["resultCode"] = A.load.Int32(ptr + 0);
      } else {
        delete x["resultCode"];
      }
      x["data"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_BulkTransfer": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "bulkTransfer" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_BulkTransfer": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.bulkTransfer);
    },
    "call_BulkTransfer": (retPtr: Pointer, handle: Pointer, transferInfo: Pointer): void => {
      const handle_ffi = {};

      if (A.load.Bool(handle + 12)) {
        handle_ffi["handle"] = A.load.Int32(handle + 0);
      }
      if (A.load.Bool(handle + 13)) {
        handle_ffi["vendorId"] = A.load.Int32(handle + 4);
      }
      if (A.load.Bool(handle + 14)) {
        handle_ffi["productId"] = A.load.Int32(handle + 8);
      }

      const transferInfo_ffi = {};

      transferInfo_ffi["direction"] = A.load.Enum(transferInfo + 0, ["in", "out"]);
      if (A.load.Bool(transferInfo + 20)) {
        transferInfo_ffi["endpoint"] = A.load.Int32(transferInfo + 4);
      }
      if (A.load.Bool(transferInfo + 21)) {
        transferInfo_ffi["length"] = A.load.Int32(transferInfo + 8);
      }
      transferInfo_ffi["data"] = A.load.Ref(transferInfo + 12, undefined);
      if (A.load.Bool(transferInfo + 22)) {
        transferInfo_ffi["timeout"] = A.load.Int32(transferInfo + 16);
      }

      const _ret = WEBEXT.usb.bulkTransfer(handle_ffi, transferInfo_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_BulkTransfer": (
      retPtr: Pointer,
      errPtr: Pointer,
      handle: Pointer,
      transferInfo: Pointer
    ): heap.Ref<boolean> => {
      try {
        const handle_ffi = {};

        if (A.load.Bool(handle + 12)) {
          handle_ffi["handle"] = A.load.Int32(handle + 0);
        }
        if (A.load.Bool(handle + 13)) {
          handle_ffi["vendorId"] = A.load.Int32(handle + 4);
        }
        if (A.load.Bool(handle + 14)) {
          handle_ffi["productId"] = A.load.Int32(handle + 8);
        }

        const transferInfo_ffi = {};

        transferInfo_ffi["direction"] = A.load.Enum(transferInfo + 0, ["in", "out"]);
        if (A.load.Bool(transferInfo + 20)) {
          transferInfo_ffi["endpoint"] = A.load.Int32(transferInfo + 4);
        }
        if (A.load.Bool(transferInfo + 21)) {
          transferInfo_ffi["length"] = A.load.Int32(transferInfo + 8);
        }
        transferInfo_ffi["data"] = A.load.Ref(transferInfo + 12, undefined);
        if (A.load.Bool(transferInfo + 22)) {
          transferInfo_ffi["timeout"] = A.load.Int32(transferInfo + 16);
        }

        const _ret = WEBEXT.usb.bulkTransfer(handle_ffi, transferInfo_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ClaimInterface": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "claimInterface" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ClaimInterface": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.claimInterface);
    },
    "call_ClaimInterface": (retPtr: Pointer, handle: Pointer, interfaceNumber: number): void => {
      const handle_ffi = {};

      if (A.load.Bool(handle + 12)) {
        handle_ffi["handle"] = A.load.Int32(handle + 0);
      }
      if (A.load.Bool(handle + 13)) {
        handle_ffi["vendorId"] = A.load.Int32(handle + 4);
      }
      if (A.load.Bool(handle + 14)) {
        handle_ffi["productId"] = A.load.Int32(handle + 8);
      }

      const _ret = WEBEXT.usb.claimInterface(handle_ffi, interfaceNumber);
      A.store.Ref(retPtr, _ret);
    },
    "try_ClaimInterface": (
      retPtr: Pointer,
      errPtr: Pointer,
      handle: Pointer,
      interfaceNumber: number
    ): heap.Ref<boolean> => {
      try {
        const handle_ffi = {};

        if (A.load.Bool(handle + 12)) {
          handle_ffi["handle"] = A.load.Int32(handle + 0);
        }
        if (A.load.Bool(handle + 13)) {
          handle_ffi["vendorId"] = A.load.Int32(handle + 4);
        }
        if (A.load.Bool(handle + 14)) {
          handle_ffi["productId"] = A.load.Int32(handle + 8);
        }

        const _ret = WEBEXT.usb.claimInterface(handle_ffi, interfaceNumber);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CloseDevice": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "closeDevice" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CloseDevice": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.closeDevice);
    },
    "call_CloseDevice": (retPtr: Pointer, handle: Pointer): void => {
      const handle_ffi = {};

      if (A.load.Bool(handle + 12)) {
        handle_ffi["handle"] = A.load.Int32(handle + 0);
      }
      if (A.load.Bool(handle + 13)) {
        handle_ffi["vendorId"] = A.load.Int32(handle + 4);
      }
      if (A.load.Bool(handle + 14)) {
        handle_ffi["productId"] = A.load.Int32(handle + 8);
      }

      const _ret = WEBEXT.usb.closeDevice(handle_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_CloseDevice": (retPtr: Pointer, errPtr: Pointer, handle: Pointer): heap.Ref<boolean> => {
      try {
        const handle_ffi = {};

        if (A.load.Bool(handle + 12)) {
          handle_ffi["handle"] = A.load.Int32(handle + 0);
        }
        if (A.load.Bool(handle + 13)) {
          handle_ffi["vendorId"] = A.load.Int32(handle + 4);
        }
        if (A.load.Bool(handle + 14)) {
          handle_ffi["productId"] = A.load.Int32(handle + 8);
        }

        const _ret = WEBEXT.usb.closeDevice(handle_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ControlTransfer": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "controlTransfer" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ControlTransfer": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.controlTransfer);
    },
    "call_ControlTransfer": (retPtr: Pointer, handle: Pointer, transferInfo: Pointer): void => {
      const handle_ffi = {};

      if (A.load.Bool(handle + 12)) {
        handle_ffi["handle"] = A.load.Int32(handle + 0);
      }
      if (A.load.Bool(handle + 13)) {
        handle_ffi["vendorId"] = A.load.Int32(handle + 4);
      }
      if (A.load.Bool(handle + 14)) {
        handle_ffi["productId"] = A.load.Int32(handle + 8);
      }

      const transferInfo_ffi = {};

      transferInfo_ffi["direction"] = A.load.Enum(transferInfo + 0, ["in", "out"]);
      transferInfo_ffi["recipient"] = A.load.Enum(transferInfo + 4, ["device", "_interface", "endpoint", "other"]);
      transferInfo_ffi["requestType"] = A.load.Enum(transferInfo + 8, ["standard", "class", "vendor", "reserved"]);
      if (A.load.Bool(transferInfo + 36)) {
        transferInfo_ffi["request"] = A.load.Int32(transferInfo + 12);
      }
      if (A.load.Bool(transferInfo + 37)) {
        transferInfo_ffi["value"] = A.load.Int32(transferInfo + 16);
      }
      if (A.load.Bool(transferInfo + 38)) {
        transferInfo_ffi["index"] = A.load.Int32(transferInfo + 20);
      }
      if (A.load.Bool(transferInfo + 39)) {
        transferInfo_ffi["length"] = A.load.Int32(transferInfo + 24);
      }
      transferInfo_ffi["data"] = A.load.Ref(transferInfo + 28, undefined);
      if (A.load.Bool(transferInfo + 40)) {
        transferInfo_ffi["timeout"] = A.load.Int32(transferInfo + 32);
      }

      const _ret = WEBEXT.usb.controlTransfer(handle_ffi, transferInfo_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ControlTransfer": (
      retPtr: Pointer,
      errPtr: Pointer,
      handle: Pointer,
      transferInfo: Pointer
    ): heap.Ref<boolean> => {
      try {
        const handle_ffi = {};

        if (A.load.Bool(handle + 12)) {
          handle_ffi["handle"] = A.load.Int32(handle + 0);
        }
        if (A.load.Bool(handle + 13)) {
          handle_ffi["vendorId"] = A.load.Int32(handle + 4);
        }
        if (A.load.Bool(handle + 14)) {
          handle_ffi["productId"] = A.load.Int32(handle + 8);
        }

        const transferInfo_ffi = {};

        transferInfo_ffi["direction"] = A.load.Enum(transferInfo + 0, ["in", "out"]);
        transferInfo_ffi["recipient"] = A.load.Enum(transferInfo + 4, ["device", "_interface", "endpoint", "other"]);
        transferInfo_ffi["requestType"] = A.load.Enum(transferInfo + 8, ["standard", "class", "vendor", "reserved"]);
        if (A.load.Bool(transferInfo + 36)) {
          transferInfo_ffi["request"] = A.load.Int32(transferInfo + 12);
        }
        if (A.load.Bool(transferInfo + 37)) {
          transferInfo_ffi["value"] = A.load.Int32(transferInfo + 16);
        }
        if (A.load.Bool(transferInfo + 38)) {
          transferInfo_ffi["index"] = A.load.Int32(transferInfo + 20);
        }
        if (A.load.Bool(transferInfo + 39)) {
          transferInfo_ffi["length"] = A.load.Int32(transferInfo + 24);
        }
        transferInfo_ffi["data"] = A.load.Ref(transferInfo + 28, undefined);
        if (A.load.Bool(transferInfo + 40)) {
          transferInfo_ffi["timeout"] = A.load.Int32(transferInfo + 32);
        }

        const _ret = WEBEXT.usb.controlTransfer(handle_ffi, transferInfo_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_FindDevices": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "findDevices" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_FindDevices": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.findDevices);
    },
    "call_FindDevices": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 12)) {
        options_ffi["vendorId"] = A.load.Int32(options + 0);
      }
      if (A.load.Bool(options + 13)) {
        options_ffi["productId"] = A.load.Int32(options + 4);
      }
      if (A.load.Bool(options + 14)) {
        options_ffi["interfaceId"] = A.load.Int32(options + 8);
      }

      const _ret = WEBEXT.usb.findDevices(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_FindDevices": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 12)) {
          options_ffi["vendorId"] = A.load.Int32(options + 0);
        }
        if (A.load.Bool(options + 13)) {
          options_ffi["productId"] = A.load.Int32(options + 4);
        }
        if (A.load.Bool(options + 14)) {
          options_ffi["interfaceId"] = A.load.Int32(options + 8);
        }

        const _ret = WEBEXT.usb.findDevices(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetConfiguration": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "getConfiguration" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetConfiguration": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.getConfiguration);
    },
    "call_GetConfiguration": (retPtr: Pointer, handle: Pointer): void => {
      const handle_ffi = {};

      if (A.load.Bool(handle + 12)) {
        handle_ffi["handle"] = A.load.Int32(handle + 0);
      }
      if (A.load.Bool(handle + 13)) {
        handle_ffi["vendorId"] = A.load.Int32(handle + 4);
      }
      if (A.load.Bool(handle + 14)) {
        handle_ffi["productId"] = A.load.Int32(handle + 8);
      }

      const _ret = WEBEXT.usb.getConfiguration(handle_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetConfiguration": (retPtr: Pointer, errPtr: Pointer, handle: Pointer): heap.Ref<boolean> => {
      try {
        const handle_ffi = {};

        if (A.load.Bool(handle + 12)) {
          handle_ffi["handle"] = A.load.Int32(handle + 0);
        }
        if (A.load.Bool(handle + 13)) {
          handle_ffi["vendorId"] = A.load.Int32(handle + 4);
        }
        if (A.load.Bool(handle + 14)) {
          handle_ffi["productId"] = A.load.Int32(handle + 8);
        }

        const _ret = WEBEXT.usb.getConfiguration(handle_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetConfigurations": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "getConfigurations" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetConfigurations": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.getConfigurations);
    },
    "call_GetConfigurations": (retPtr: Pointer, device: Pointer): void => {
      const device_ffi = {};

      if (A.load.Bool(device + 28)) {
        device_ffi["device"] = A.load.Int32(device + 0);
      }
      if (A.load.Bool(device + 29)) {
        device_ffi["vendorId"] = A.load.Int32(device + 4);
      }
      if (A.load.Bool(device + 30)) {
        device_ffi["productId"] = A.load.Int32(device + 8);
      }
      if (A.load.Bool(device + 31)) {
        device_ffi["version"] = A.load.Int32(device + 12);
      }
      device_ffi["productName"] = A.load.Ref(device + 16, undefined);
      device_ffi["manufacturerName"] = A.load.Ref(device + 20, undefined);
      device_ffi["serialNumber"] = A.load.Ref(device + 24, undefined);

      const _ret = WEBEXT.usb.getConfigurations(device_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetConfigurations": (retPtr: Pointer, errPtr: Pointer, device: Pointer): heap.Ref<boolean> => {
      try {
        const device_ffi = {};

        if (A.load.Bool(device + 28)) {
          device_ffi["device"] = A.load.Int32(device + 0);
        }
        if (A.load.Bool(device + 29)) {
          device_ffi["vendorId"] = A.load.Int32(device + 4);
        }
        if (A.load.Bool(device + 30)) {
          device_ffi["productId"] = A.load.Int32(device + 8);
        }
        if (A.load.Bool(device + 31)) {
          device_ffi["version"] = A.load.Int32(device + 12);
        }
        device_ffi["productName"] = A.load.Ref(device + 16, undefined);
        device_ffi["manufacturerName"] = A.load.Ref(device + 20, undefined);
        device_ffi["serialNumber"] = A.load.Ref(device + 24, undefined);

        const _ret = WEBEXT.usb.getConfigurations(device_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDevices": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "getDevices" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDevices": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.getDevices);
    },
    "call_GetDevices": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 12)) {
        options_ffi["vendorId"] = A.load.Int32(options + 0);
      }
      if (A.load.Bool(options + 13)) {
        options_ffi["productId"] = A.load.Int32(options + 4);
      }
      options_ffi["filters"] = A.load.Ref(options + 8, undefined);

      const _ret = WEBEXT.usb.getDevices(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDevices": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 12)) {
          options_ffi["vendorId"] = A.load.Int32(options + 0);
        }
        if (A.load.Bool(options + 13)) {
          options_ffi["productId"] = A.load.Int32(options + 4);
        }
        options_ffi["filters"] = A.load.Ref(options + 8, undefined);

        const _ret = WEBEXT.usb.getDevices(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetUserSelectedDevices": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "getUserSelectedDevices" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetUserSelectedDevices": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.getUserSelectedDevices);
    },
    "call_GetUserSelectedDevices": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 8)) {
        options_ffi["multiple"] = A.load.Bool(options + 0);
      }
      options_ffi["filters"] = A.load.Ref(options + 4, undefined);

      const _ret = WEBEXT.usb.getUserSelectedDevices(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetUserSelectedDevices": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 8)) {
          options_ffi["multiple"] = A.load.Bool(options + 0);
        }
        options_ffi["filters"] = A.load.Ref(options + 4, undefined);

        const _ret = WEBEXT.usb.getUserSelectedDevices(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InterruptTransfer": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "interruptTransfer" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InterruptTransfer": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.interruptTransfer);
    },
    "call_InterruptTransfer": (retPtr: Pointer, handle: Pointer, transferInfo: Pointer): void => {
      const handle_ffi = {};

      if (A.load.Bool(handle + 12)) {
        handle_ffi["handle"] = A.load.Int32(handle + 0);
      }
      if (A.load.Bool(handle + 13)) {
        handle_ffi["vendorId"] = A.load.Int32(handle + 4);
      }
      if (A.load.Bool(handle + 14)) {
        handle_ffi["productId"] = A.load.Int32(handle + 8);
      }

      const transferInfo_ffi = {};

      transferInfo_ffi["direction"] = A.load.Enum(transferInfo + 0, ["in", "out"]);
      if (A.load.Bool(transferInfo + 20)) {
        transferInfo_ffi["endpoint"] = A.load.Int32(transferInfo + 4);
      }
      if (A.load.Bool(transferInfo + 21)) {
        transferInfo_ffi["length"] = A.load.Int32(transferInfo + 8);
      }
      transferInfo_ffi["data"] = A.load.Ref(transferInfo + 12, undefined);
      if (A.load.Bool(transferInfo + 22)) {
        transferInfo_ffi["timeout"] = A.load.Int32(transferInfo + 16);
      }

      const _ret = WEBEXT.usb.interruptTransfer(handle_ffi, transferInfo_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_InterruptTransfer": (
      retPtr: Pointer,
      errPtr: Pointer,
      handle: Pointer,
      transferInfo: Pointer
    ): heap.Ref<boolean> => {
      try {
        const handle_ffi = {};

        if (A.load.Bool(handle + 12)) {
          handle_ffi["handle"] = A.load.Int32(handle + 0);
        }
        if (A.load.Bool(handle + 13)) {
          handle_ffi["vendorId"] = A.load.Int32(handle + 4);
        }
        if (A.load.Bool(handle + 14)) {
          handle_ffi["productId"] = A.load.Int32(handle + 8);
        }

        const transferInfo_ffi = {};

        transferInfo_ffi["direction"] = A.load.Enum(transferInfo + 0, ["in", "out"]);
        if (A.load.Bool(transferInfo + 20)) {
          transferInfo_ffi["endpoint"] = A.load.Int32(transferInfo + 4);
        }
        if (A.load.Bool(transferInfo + 21)) {
          transferInfo_ffi["length"] = A.load.Int32(transferInfo + 8);
        }
        transferInfo_ffi["data"] = A.load.Ref(transferInfo + 12, undefined);
        if (A.load.Bool(transferInfo + 22)) {
          transferInfo_ffi["timeout"] = A.load.Int32(transferInfo + 16);
        }

        const _ret = WEBEXT.usb.interruptTransfer(handle_ffi, transferInfo_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsochronousTransfer": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "isochronousTransfer" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsochronousTransfer": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.isochronousTransfer);
    },
    "call_IsochronousTransfer": (retPtr: Pointer, handle: Pointer, transferInfo: Pointer): void => {
      const handle_ffi = {};

      if (A.load.Bool(handle + 12)) {
        handle_ffi["handle"] = A.load.Int32(handle + 0);
      }
      if (A.load.Bool(handle + 13)) {
        handle_ffi["vendorId"] = A.load.Int32(handle + 4);
      }
      if (A.load.Bool(handle + 14)) {
        handle_ffi["productId"] = A.load.Int32(handle + 8);
      }

      const transferInfo_ffi = {};

      if (A.load.Bool(transferInfo + 0 + 23)) {
        transferInfo_ffi["transferInfo"] = {};
        transferInfo_ffi["transferInfo"]["direction"] = A.load.Enum(transferInfo + 0 + 0, ["in", "out"]);
        if (A.load.Bool(transferInfo + 0 + 20)) {
          transferInfo_ffi["transferInfo"]["endpoint"] = A.load.Int32(transferInfo + 0 + 4);
        }
        if (A.load.Bool(transferInfo + 0 + 21)) {
          transferInfo_ffi["transferInfo"]["length"] = A.load.Int32(transferInfo + 0 + 8);
        }
        transferInfo_ffi["transferInfo"]["data"] = A.load.Ref(transferInfo + 0 + 12, undefined);
        if (A.load.Bool(transferInfo + 0 + 22)) {
          transferInfo_ffi["transferInfo"]["timeout"] = A.load.Int32(transferInfo + 0 + 16);
        }
      }
      if (A.load.Bool(transferInfo + 32)) {
        transferInfo_ffi["packets"] = A.load.Int32(transferInfo + 24);
      }
      if (A.load.Bool(transferInfo + 33)) {
        transferInfo_ffi["packetLength"] = A.load.Int32(transferInfo + 28);
      }

      const _ret = WEBEXT.usb.isochronousTransfer(handle_ffi, transferInfo_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_IsochronousTransfer": (
      retPtr: Pointer,
      errPtr: Pointer,
      handle: Pointer,
      transferInfo: Pointer
    ): heap.Ref<boolean> => {
      try {
        const handle_ffi = {};

        if (A.load.Bool(handle + 12)) {
          handle_ffi["handle"] = A.load.Int32(handle + 0);
        }
        if (A.load.Bool(handle + 13)) {
          handle_ffi["vendorId"] = A.load.Int32(handle + 4);
        }
        if (A.load.Bool(handle + 14)) {
          handle_ffi["productId"] = A.load.Int32(handle + 8);
        }

        const transferInfo_ffi = {};

        if (A.load.Bool(transferInfo + 0 + 23)) {
          transferInfo_ffi["transferInfo"] = {};
          transferInfo_ffi["transferInfo"]["direction"] = A.load.Enum(transferInfo + 0 + 0, ["in", "out"]);
          if (A.load.Bool(transferInfo + 0 + 20)) {
            transferInfo_ffi["transferInfo"]["endpoint"] = A.load.Int32(transferInfo + 0 + 4);
          }
          if (A.load.Bool(transferInfo + 0 + 21)) {
            transferInfo_ffi["transferInfo"]["length"] = A.load.Int32(transferInfo + 0 + 8);
          }
          transferInfo_ffi["transferInfo"]["data"] = A.load.Ref(transferInfo + 0 + 12, undefined);
          if (A.load.Bool(transferInfo + 0 + 22)) {
            transferInfo_ffi["transferInfo"]["timeout"] = A.load.Int32(transferInfo + 0 + 16);
          }
        }
        if (A.load.Bool(transferInfo + 32)) {
          transferInfo_ffi["packets"] = A.load.Int32(transferInfo + 24);
        }
        if (A.load.Bool(transferInfo + 33)) {
          transferInfo_ffi["packetLength"] = A.load.Int32(transferInfo + 28);
        }

        const _ret = WEBEXT.usb.isochronousTransfer(handle_ffi, transferInfo_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ListInterfaces": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "listInterfaces" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ListInterfaces": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.listInterfaces);
    },
    "call_ListInterfaces": (retPtr: Pointer, handle: Pointer): void => {
      const handle_ffi = {};

      if (A.load.Bool(handle + 12)) {
        handle_ffi["handle"] = A.load.Int32(handle + 0);
      }
      if (A.load.Bool(handle + 13)) {
        handle_ffi["vendorId"] = A.load.Int32(handle + 4);
      }
      if (A.load.Bool(handle + 14)) {
        handle_ffi["productId"] = A.load.Int32(handle + 8);
      }

      const _ret = WEBEXT.usb.listInterfaces(handle_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ListInterfaces": (retPtr: Pointer, errPtr: Pointer, handle: Pointer): heap.Ref<boolean> => {
      try {
        const handle_ffi = {};

        if (A.load.Bool(handle + 12)) {
          handle_ffi["handle"] = A.load.Int32(handle + 0);
        }
        if (A.load.Bool(handle + 13)) {
          handle_ffi["vendorId"] = A.load.Int32(handle + 4);
        }
        if (A.load.Bool(handle + 14)) {
          handle_ffi["productId"] = A.load.Int32(handle + 8);
        }

        const _ret = WEBEXT.usb.listInterfaces(handle_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeviceAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb?.onDeviceAdded && "addListener" in WEBEXT?.usb?.onDeviceAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeviceAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.onDeviceAdded.addListener);
    },
    "call_OnDeviceAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.usb.onDeviceAdded.addListener(A.H.get<object>(callback));
    },
    "try_OnDeviceAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.usb.onDeviceAdded.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeviceAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb?.onDeviceAdded && "removeListener" in WEBEXT?.usb?.onDeviceAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeviceAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.onDeviceAdded.removeListener);
    },
    "call_OffDeviceAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.usb.onDeviceAdded.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeviceAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.usb.onDeviceAdded.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeviceAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb?.onDeviceAdded && "hasListener" in WEBEXT?.usb?.onDeviceAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeviceAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.onDeviceAdded.hasListener);
    },
    "call_HasOnDeviceAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.usb.onDeviceAdded.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeviceAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.usb.onDeviceAdded.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeviceRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb?.onDeviceRemoved && "addListener" in WEBEXT?.usb?.onDeviceRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeviceRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.onDeviceRemoved.addListener);
    },
    "call_OnDeviceRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.usb.onDeviceRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnDeviceRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.usb.onDeviceRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeviceRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb?.onDeviceRemoved && "removeListener" in WEBEXT?.usb?.onDeviceRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeviceRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.onDeviceRemoved.removeListener);
    },
    "call_OffDeviceRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.usb.onDeviceRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeviceRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.usb.onDeviceRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeviceRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb?.onDeviceRemoved && "hasListener" in WEBEXT?.usb?.onDeviceRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeviceRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.onDeviceRemoved.hasListener);
    },
    "call_HasOnDeviceRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.usb.onDeviceRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeviceRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.usb.onDeviceRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenDevice": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "openDevice" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenDevice": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.openDevice);
    },
    "call_OpenDevice": (retPtr: Pointer, device: Pointer): void => {
      const device_ffi = {};

      if (A.load.Bool(device + 28)) {
        device_ffi["device"] = A.load.Int32(device + 0);
      }
      if (A.load.Bool(device + 29)) {
        device_ffi["vendorId"] = A.load.Int32(device + 4);
      }
      if (A.load.Bool(device + 30)) {
        device_ffi["productId"] = A.load.Int32(device + 8);
      }
      if (A.load.Bool(device + 31)) {
        device_ffi["version"] = A.load.Int32(device + 12);
      }
      device_ffi["productName"] = A.load.Ref(device + 16, undefined);
      device_ffi["manufacturerName"] = A.load.Ref(device + 20, undefined);
      device_ffi["serialNumber"] = A.load.Ref(device + 24, undefined);

      const _ret = WEBEXT.usb.openDevice(device_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_OpenDevice": (retPtr: Pointer, errPtr: Pointer, device: Pointer): heap.Ref<boolean> => {
      try {
        const device_ffi = {};

        if (A.load.Bool(device + 28)) {
          device_ffi["device"] = A.load.Int32(device + 0);
        }
        if (A.load.Bool(device + 29)) {
          device_ffi["vendorId"] = A.load.Int32(device + 4);
        }
        if (A.load.Bool(device + 30)) {
          device_ffi["productId"] = A.load.Int32(device + 8);
        }
        if (A.load.Bool(device + 31)) {
          device_ffi["version"] = A.load.Int32(device + 12);
        }
        device_ffi["productName"] = A.load.Ref(device + 16, undefined);
        device_ffi["manufacturerName"] = A.load.Ref(device + 20, undefined);
        device_ffi["serialNumber"] = A.load.Ref(device + 24, undefined);

        const _ret = WEBEXT.usb.openDevice(device_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReleaseInterface": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "releaseInterface" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReleaseInterface": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.releaseInterface);
    },
    "call_ReleaseInterface": (retPtr: Pointer, handle: Pointer, interfaceNumber: number): void => {
      const handle_ffi = {};

      if (A.load.Bool(handle + 12)) {
        handle_ffi["handle"] = A.load.Int32(handle + 0);
      }
      if (A.load.Bool(handle + 13)) {
        handle_ffi["vendorId"] = A.load.Int32(handle + 4);
      }
      if (A.load.Bool(handle + 14)) {
        handle_ffi["productId"] = A.load.Int32(handle + 8);
      }

      const _ret = WEBEXT.usb.releaseInterface(handle_ffi, interfaceNumber);
      A.store.Ref(retPtr, _ret);
    },
    "try_ReleaseInterface": (
      retPtr: Pointer,
      errPtr: Pointer,
      handle: Pointer,
      interfaceNumber: number
    ): heap.Ref<boolean> => {
      try {
        const handle_ffi = {};

        if (A.load.Bool(handle + 12)) {
          handle_ffi["handle"] = A.load.Int32(handle + 0);
        }
        if (A.load.Bool(handle + 13)) {
          handle_ffi["vendorId"] = A.load.Int32(handle + 4);
        }
        if (A.load.Bool(handle + 14)) {
          handle_ffi["productId"] = A.load.Int32(handle + 8);
        }

        const _ret = WEBEXT.usb.releaseInterface(handle_ffi, interfaceNumber);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RequestAccess": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "requestAccess" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RequestAccess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.requestAccess);
    },
    "call_RequestAccess": (retPtr: Pointer, device: Pointer, interfaceId: number): void => {
      const device_ffi = {};

      if (A.load.Bool(device + 28)) {
        device_ffi["device"] = A.load.Int32(device + 0);
      }
      if (A.load.Bool(device + 29)) {
        device_ffi["vendorId"] = A.load.Int32(device + 4);
      }
      if (A.load.Bool(device + 30)) {
        device_ffi["productId"] = A.load.Int32(device + 8);
      }
      if (A.load.Bool(device + 31)) {
        device_ffi["version"] = A.load.Int32(device + 12);
      }
      device_ffi["productName"] = A.load.Ref(device + 16, undefined);
      device_ffi["manufacturerName"] = A.load.Ref(device + 20, undefined);
      device_ffi["serialNumber"] = A.load.Ref(device + 24, undefined);

      const _ret = WEBEXT.usb.requestAccess(device_ffi, interfaceId);
      A.store.Ref(retPtr, _ret);
    },
    "try_RequestAccess": (
      retPtr: Pointer,
      errPtr: Pointer,
      device: Pointer,
      interfaceId: number
    ): heap.Ref<boolean> => {
      try {
        const device_ffi = {};

        if (A.load.Bool(device + 28)) {
          device_ffi["device"] = A.load.Int32(device + 0);
        }
        if (A.load.Bool(device + 29)) {
          device_ffi["vendorId"] = A.load.Int32(device + 4);
        }
        if (A.load.Bool(device + 30)) {
          device_ffi["productId"] = A.load.Int32(device + 8);
        }
        if (A.load.Bool(device + 31)) {
          device_ffi["version"] = A.load.Int32(device + 12);
        }
        device_ffi["productName"] = A.load.Ref(device + 16, undefined);
        device_ffi["manufacturerName"] = A.load.Ref(device + 20, undefined);
        device_ffi["serialNumber"] = A.load.Ref(device + 24, undefined);

        const _ret = WEBEXT.usb.requestAccess(device_ffi, interfaceId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ResetDevice": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "resetDevice" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ResetDevice": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.resetDevice);
    },
    "call_ResetDevice": (retPtr: Pointer, handle: Pointer): void => {
      const handle_ffi = {};

      if (A.load.Bool(handle + 12)) {
        handle_ffi["handle"] = A.load.Int32(handle + 0);
      }
      if (A.load.Bool(handle + 13)) {
        handle_ffi["vendorId"] = A.load.Int32(handle + 4);
      }
      if (A.load.Bool(handle + 14)) {
        handle_ffi["productId"] = A.load.Int32(handle + 8);
      }

      const _ret = WEBEXT.usb.resetDevice(handle_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ResetDevice": (retPtr: Pointer, errPtr: Pointer, handle: Pointer): heap.Ref<boolean> => {
      try {
        const handle_ffi = {};

        if (A.load.Bool(handle + 12)) {
          handle_ffi["handle"] = A.load.Int32(handle + 0);
        }
        if (A.load.Bool(handle + 13)) {
          handle_ffi["vendorId"] = A.load.Int32(handle + 4);
        }
        if (A.load.Bool(handle + 14)) {
          handle_ffi["productId"] = A.load.Int32(handle + 8);
        }

        const _ret = WEBEXT.usb.resetDevice(handle_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetConfiguration": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "setConfiguration" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetConfiguration": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.setConfiguration);
    },
    "call_SetConfiguration": (retPtr: Pointer, handle: Pointer, configurationValue: number): void => {
      const handle_ffi = {};

      if (A.load.Bool(handle + 12)) {
        handle_ffi["handle"] = A.load.Int32(handle + 0);
      }
      if (A.load.Bool(handle + 13)) {
        handle_ffi["vendorId"] = A.load.Int32(handle + 4);
      }
      if (A.load.Bool(handle + 14)) {
        handle_ffi["productId"] = A.load.Int32(handle + 8);
      }

      const _ret = WEBEXT.usb.setConfiguration(handle_ffi, configurationValue);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetConfiguration": (
      retPtr: Pointer,
      errPtr: Pointer,
      handle: Pointer,
      configurationValue: number
    ): heap.Ref<boolean> => {
      try {
        const handle_ffi = {};

        if (A.load.Bool(handle + 12)) {
          handle_ffi["handle"] = A.load.Int32(handle + 0);
        }
        if (A.load.Bool(handle + 13)) {
          handle_ffi["vendorId"] = A.load.Int32(handle + 4);
        }
        if (A.load.Bool(handle + 14)) {
          handle_ffi["productId"] = A.load.Int32(handle + 8);
        }

        const _ret = WEBEXT.usb.setConfiguration(handle_ffi, configurationValue);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetInterfaceAlternateSetting": (): heap.Ref<boolean> => {
      if (WEBEXT?.usb && "setInterfaceAlternateSetting" in WEBEXT?.usb) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetInterfaceAlternateSetting": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usb.setInterfaceAlternateSetting);
    },
    "call_SetInterfaceAlternateSetting": (
      retPtr: Pointer,
      handle: Pointer,
      interfaceNumber: number,
      alternateSetting: number
    ): void => {
      const handle_ffi = {};

      if (A.load.Bool(handle + 12)) {
        handle_ffi["handle"] = A.load.Int32(handle + 0);
      }
      if (A.load.Bool(handle + 13)) {
        handle_ffi["vendorId"] = A.load.Int32(handle + 4);
      }
      if (A.load.Bool(handle + 14)) {
        handle_ffi["productId"] = A.load.Int32(handle + 8);
      }

      const _ret = WEBEXT.usb.setInterfaceAlternateSetting(handle_ffi, interfaceNumber, alternateSetting);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetInterfaceAlternateSetting": (
      retPtr: Pointer,
      errPtr: Pointer,
      handle: Pointer,
      interfaceNumber: number,
      alternateSetting: number
    ): heap.Ref<boolean> => {
      try {
        const handle_ffi = {};

        if (A.load.Bool(handle + 12)) {
          handle_ffi["handle"] = A.load.Int32(handle + 0);
        }
        if (A.load.Bool(handle + 13)) {
          handle_ffi["vendorId"] = A.load.Int32(handle + 4);
        }
        if (A.load.Bool(handle + 14)) {
          handle_ffi["productId"] = A.load.Int32(handle + 8);
        }

        const _ret = WEBEXT.usb.setInterfaceAlternateSetting(handle_ffi, interfaceNumber, alternateSetting);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
