import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/bluetoothlowenergy", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_AdvertisementType": (ref: heap.Ref<string>): number => {
      const idx = ["broadcast", "peripheral"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ManufacturerData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "id" in x ? true : false);
        A.store.Int32(ptr + 0, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Ref(ptr + 4, x["data"]);
      }
    },
    "load_ManufacturerData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["id"] = A.load.Int32(ptr + 0);
      } else {
        delete x["id"];
      }
      x["data"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ServiceData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["uuid"]);
        A.store.Ref(ptr + 4, x["data"]);
      }
    },
    "load_ServiceData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["uuid"] = A.load.Ref(ptr + 0, undefined);
      x["data"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Advertisement": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Enum(ptr + 0, ["broadcast", "peripheral"].indexOf(x["type"] as string));
        A.store.Ref(ptr + 4, x["serviceUuids"]);
        A.store.Ref(ptr + 8, x["manufacturerData"]);
        A.store.Ref(ptr + 12, x["solicitUuids"]);
        A.store.Ref(ptr + 16, x["serviceData"]);
      }
    },
    "load_Advertisement": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, ["broadcast", "peripheral"]);
      x["serviceUuids"] = A.load.Ref(ptr + 4, undefined);
      x["manufacturerData"] = A.load.Ref(ptr + 8, undefined);
      x["solicitUuids"] = A.load.Ref(ptr + 12, undefined);
      x["serviceData"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Service": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["uuid"]);
        A.store.Bool(ptr + 16, "isPrimary" in x ? true : false);
        A.store.Bool(ptr + 4, x["isPrimary"] ? true : false);
        A.store.Ref(ptr + 8, x["instanceId"]);
        A.store.Ref(ptr + 12, x["deviceAddress"]);
      }
    },
    "load_Service": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["uuid"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["isPrimary"] = A.load.Bool(ptr + 4);
      } else {
        delete x["isPrimary"];
      }
      x["instanceId"] = A.load.Ref(ptr + 8, undefined);
      x["deviceAddress"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_CharacteristicProperty": (ref: heap.Ref<string>): number => {
      const idx = [
        "broadcast",
        "read",
        "writeWithoutResponse",
        "write",
        "notify",
        "indicate",
        "authenticatedSignedWrites",
        "extendedProperties",
        "reliableWrite",
        "writableAuxiliaries",
        "encryptRead",
        "encryptWrite",
        "encryptAuthenticatedRead",
        "encryptAuthenticatedWrite",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Characteristic": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 36, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 4 + 17, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Bool(ptr + 4 + 16, false);
        A.store.Bool(ptr + 4 + 4, false);
        A.store.Ref(ptr + 4 + 8, undefined);
        A.store.Ref(ptr + 4 + 12, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
      } else {
        A.store.Bool(ptr + 36, true);
        A.store.Ref(ptr + 0, x["uuid"]);

        if (typeof x["service"] === "undefined") {
          A.store.Bool(ptr + 4 + 17, false);
          A.store.Ref(ptr + 4 + 0, undefined);
          A.store.Bool(ptr + 4 + 16, false);
          A.store.Bool(ptr + 4 + 4, false);
          A.store.Ref(ptr + 4 + 8, undefined);
          A.store.Ref(ptr + 4 + 12, undefined);
        } else {
          A.store.Bool(ptr + 4 + 17, true);
          A.store.Ref(ptr + 4 + 0, x["service"]["uuid"]);
          A.store.Bool(ptr + 4 + 16, "isPrimary" in x["service"] ? true : false);
          A.store.Bool(ptr + 4 + 4, x["service"]["isPrimary"] ? true : false);
          A.store.Ref(ptr + 4 + 8, x["service"]["instanceId"]);
          A.store.Ref(ptr + 4 + 12, x["service"]["deviceAddress"]);
        }
        A.store.Ref(ptr + 24, x["properties"]);
        A.store.Ref(ptr + 28, x["instanceId"]);
        A.store.Ref(ptr + 32, x["value"]);
      }
    },
    "load_Characteristic": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["uuid"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 4 + 17)) {
        x["service"] = {};
        x["service"]["uuid"] = A.load.Ref(ptr + 4 + 0, undefined);
        if (A.load.Bool(ptr + 4 + 16)) {
          x["service"]["isPrimary"] = A.load.Bool(ptr + 4 + 4);
        } else {
          delete x["service"]["isPrimary"];
        }
        x["service"]["instanceId"] = A.load.Ref(ptr + 4 + 8, undefined);
        x["service"]["deviceAddress"] = A.load.Ref(ptr + 4 + 12, undefined);
      } else {
        delete x["service"];
      }
      x["properties"] = A.load.Ref(ptr + 24, undefined);
      x["instanceId"] = A.load.Ref(ptr + 28, undefined);
      x["value"] = A.load.Ref(ptr + 32, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ConnectProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 1, "persistent" in x ? true : false);
        A.store.Bool(ptr + 0, x["persistent"] ? true : false);
      }
    },
    "load_ConnectProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 1)) {
        x["persistent"] = A.load.Bool(ptr + 0);
      } else {
        delete x["persistent"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DescriptorPermission": (ref: heap.Ref<string>): number => {
      const idx = [
        "read",
        "write",
        "encryptedRead",
        "encryptedWrite",
        "encryptedAuthenticatedRead",
        "encryptedAuthenticatedWrite",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Descriptor": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 56, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 4 + 36, false);
        A.store.Ref(ptr + 4 + 0, undefined);

        A.store.Bool(ptr + 4 + 4 + 17, false);
        A.store.Ref(ptr + 4 + 4 + 0, undefined);
        A.store.Bool(ptr + 4 + 4 + 16, false);
        A.store.Bool(ptr + 4 + 4 + 4, false);
        A.store.Ref(ptr + 4 + 4 + 8, undefined);
        A.store.Ref(ptr + 4 + 4 + 12, undefined);
        A.store.Ref(ptr + 4 + 24, undefined);
        A.store.Ref(ptr + 4 + 28, undefined);
        A.store.Ref(ptr + 4 + 32, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Ref(ptr + 48, undefined);
        A.store.Ref(ptr + 52, undefined);
      } else {
        A.store.Bool(ptr + 56, true);
        A.store.Ref(ptr + 0, x["uuid"]);

        if (typeof x["characteristic"] === "undefined") {
          A.store.Bool(ptr + 4 + 36, false);
          A.store.Ref(ptr + 4 + 0, undefined);

          A.store.Bool(ptr + 4 + 4 + 17, false);
          A.store.Ref(ptr + 4 + 4 + 0, undefined);
          A.store.Bool(ptr + 4 + 4 + 16, false);
          A.store.Bool(ptr + 4 + 4 + 4, false);
          A.store.Ref(ptr + 4 + 4 + 8, undefined);
          A.store.Ref(ptr + 4 + 4 + 12, undefined);
          A.store.Ref(ptr + 4 + 24, undefined);
          A.store.Ref(ptr + 4 + 28, undefined);
          A.store.Ref(ptr + 4 + 32, undefined);
        } else {
          A.store.Bool(ptr + 4 + 36, true);
          A.store.Ref(ptr + 4 + 0, x["characteristic"]["uuid"]);

          if (typeof x["characteristic"]["service"] === "undefined") {
            A.store.Bool(ptr + 4 + 4 + 17, false);
            A.store.Ref(ptr + 4 + 4 + 0, undefined);
            A.store.Bool(ptr + 4 + 4 + 16, false);
            A.store.Bool(ptr + 4 + 4 + 4, false);
            A.store.Ref(ptr + 4 + 4 + 8, undefined);
            A.store.Ref(ptr + 4 + 4 + 12, undefined);
          } else {
            A.store.Bool(ptr + 4 + 4 + 17, true);
            A.store.Ref(ptr + 4 + 4 + 0, x["characteristic"]["service"]["uuid"]);
            A.store.Bool(ptr + 4 + 4 + 16, "isPrimary" in x["characteristic"]["service"] ? true : false);
            A.store.Bool(ptr + 4 + 4 + 4, x["characteristic"]["service"]["isPrimary"] ? true : false);
            A.store.Ref(ptr + 4 + 4 + 8, x["characteristic"]["service"]["instanceId"]);
            A.store.Ref(ptr + 4 + 4 + 12, x["characteristic"]["service"]["deviceAddress"]);
          }
          A.store.Ref(ptr + 4 + 24, x["characteristic"]["properties"]);
          A.store.Ref(ptr + 4 + 28, x["characteristic"]["instanceId"]);
          A.store.Ref(ptr + 4 + 32, x["characteristic"]["value"]);
        }
        A.store.Ref(ptr + 44, x["permissions"]);
        A.store.Ref(ptr + 48, x["instanceId"]);
        A.store.Ref(ptr + 52, x["value"]);
      }
    },
    "load_Descriptor": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["uuid"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 4 + 36)) {
        x["characteristic"] = {};
        x["characteristic"]["uuid"] = A.load.Ref(ptr + 4 + 0, undefined);
        if (A.load.Bool(ptr + 4 + 4 + 17)) {
          x["characteristic"]["service"] = {};
          x["characteristic"]["service"]["uuid"] = A.load.Ref(ptr + 4 + 4 + 0, undefined);
          if (A.load.Bool(ptr + 4 + 4 + 16)) {
            x["characteristic"]["service"]["isPrimary"] = A.load.Bool(ptr + 4 + 4 + 4);
          } else {
            delete x["characteristic"]["service"]["isPrimary"];
          }
          x["characteristic"]["service"]["instanceId"] = A.load.Ref(ptr + 4 + 4 + 8, undefined);
          x["characteristic"]["service"]["deviceAddress"] = A.load.Ref(ptr + 4 + 4 + 12, undefined);
        } else {
          delete x["characteristic"]["service"];
        }
        x["characteristic"]["properties"] = A.load.Ref(ptr + 4 + 24, undefined);
        x["characteristic"]["instanceId"] = A.load.Ref(ptr + 4 + 28, undefined);
        x["characteristic"]["value"] = A.load.Ref(ptr + 4 + 32, undefined);
      } else {
        delete x["characteristic"];
      }
      x["permissions"] = A.load.Ref(ptr + 44, undefined);
      x["instanceId"] = A.load.Ref(ptr + 48, undefined);
      x["value"] = A.load.Ref(ptr + 52, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Device": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Ref(ptr + 0, x["address"]);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Bool(ptr + 12, "deviceClass" in x ? true : false);
        A.store.Int32(ptr + 8, x["deviceClass"] === undefined ? 0 : (x["deviceClass"] as number));
      }
    },
    "load_Device": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["address"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["deviceClass"] = A.load.Int32(ptr + 8);
      } else {
        delete x["deviceClass"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Notification": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 6, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 6, true);
        A.store.Ref(ptr + 0, x["value"]);
        A.store.Bool(ptr + 5, "shouldIndicate" in x ? true : false);
        A.store.Bool(ptr + 4, x["shouldIndicate"] ? true : false);
      }
    },
    "load_Notification": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["value"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 5)) {
        x["shouldIndicate"] = A.load.Bool(ptr + 4);
      } else {
        delete x["shouldIndicate"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_NotificationProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 1, "persistent" in x ? true : false);
        A.store.Bool(ptr + 0, x["persistent"] ? true : false);
      }
    },
    "load_NotificationProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 1)) {
        x["persistent"] = A.load.Bool(ptr + 0);
      } else {
        delete x["persistent"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Request": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 24, false);
        A.store.Int32(ptr + 0, 0);

        A.store.Bool(ptr + 4 + 13, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Ref(ptr + 4 + 4, undefined);
        A.store.Bool(ptr + 4 + 12, false);
        A.store.Int32(ptr + 4 + 8, 0);
        A.store.Ref(ptr + 20, undefined);
      } else {
        A.store.Bool(ptr + 25, true);
        A.store.Bool(ptr + 24, "requestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["requestId"] === undefined ? 0 : (x["requestId"] as number));

        if (typeof x["device"] === "undefined") {
          A.store.Bool(ptr + 4 + 13, false);
          A.store.Ref(ptr + 4 + 0, undefined);
          A.store.Ref(ptr + 4 + 4, undefined);
          A.store.Bool(ptr + 4 + 12, false);
          A.store.Int32(ptr + 4 + 8, 0);
        } else {
          A.store.Bool(ptr + 4 + 13, true);
          A.store.Ref(ptr + 4 + 0, x["device"]["address"]);
          A.store.Ref(ptr + 4 + 4, x["device"]["name"]);
          A.store.Bool(ptr + 4 + 12, "deviceClass" in x["device"] ? true : false);
          A.store.Int32(
            ptr + 4 + 8,
            x["device"]["deviceClass"] === undefined ? 0 : (x["device"]["deviceClass"] as number)
          );
        }
        A.store.Ref(ptr + 20, x["value"]);
      }
    },
    "load_Request": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["requestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["requestId"];
      }
      if (A.load.Bool(ptr + 4 + 13)) {
        x["device"] = {};
        x["device"]["address"] = A.load.Ref(ptr + 4 + 0, undefined);
        x["device"]["name"] = A.load.Ref(ptr + 4 + 4, undefined);
        if (A.load.Bool(ptr + 4 + 12)) {
          x["device"]["deviceClass"] = A.load.Int32(ptr + 4 + 8);
        } else {
          delete x["device"]["deviceClass"];
        }
      } else {
        delete x["device"];
      }
      x["value"] = A.load.Ref(ptr + 20, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Response": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Bool(ptr + 12, "requestId" in x ? true : false);
        A.store.Int32(ptr + 0, x["requestId"] === undefined ? 0 : (x["requestId"] as number));
        A.store.Bool(ptr + 13, "isError" in x ? true : false);
        A.store.Bool(ptr + 4, x["isError"] ? true : false);
        A.store.Ref(ptr + 8, x["value"]);
      }
    },
    "load_Response": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["requestId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["requestId"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["isError"] = A.load.Bool(ptr + 4);
      } else {
        delete x["isError"];
      }
      x["value"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Connect": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "connect" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Connect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.connect);
    },
    "call_Connect": (retPtr: Pointer, deviceAddress: heap.Ref<object>, properties: Pointer): void => {
      const properties_ffi = {};

      if (A.load.Bool(properties + 1)) {
        properties_ffi["persistent"] = A.load.Bool(properties + 0);
      }

      const _ret = WEBEXT.bluetoothLowEnergy.connect(A.H.get<object>(deviceAddress), properties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Connect": (
      retPtr: Pointer,
      errPtr: Pointer,
      deviceAddress: heap.Ref<object>,
      properties: Pointer
    ): heap.Ref<boolean> => {
      try {
        const properties_ffi = {};

        if (A.load.Bool(properties + 1)) {
          properties_ffi["persistent"] = A.load.Bool(properties + 0);
        }

        const _ret = WEBEXT.bluetoothLowEnergy.connect(A.H.get<object>(deviceAddress), properties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CreateCharacteristic": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "createCharacteristic" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CreateCharacteristic": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.createCharacteristic);
    },
    "call_CreateCharacteristic": (retPtr: Pointer, characteristic: Pointer, serviceId: heap.Ref<object>): void => {
      const characteristic_ffi = {};

      characteristic_ffi["uuid"] = A.load.Ref(characteristic + 0, undefined);
      if (A.load.Bool(characteristic + 4 + 17)) {
        characteristic_ffi["service"] = {};
        characteristic_ffi["service"]["uuid"] = A.load.Ref(characteristic + 4 + 0, undefined);
        if (A.load.Bool(characteristic + 4 + 16)) {
          characteristic_ffi["service"]["isPrimary"] = A.load.Bool(characteristic + 4 + 4);
        }
        characteristic_ffi["service"]["instanceId"] = A.load.Ref(characteristic + 4 + 8, undefined);
        characteristic_ffi["service"]["deviceAddress"] = A.load.Ref(characteristic + 4 + 12, undefined);
      }
      characteristic_ffi["properties"] = A.load.Ref(characteristic + 24, undefined);
      characteristic_ffi["instanceId"] = A.load.Ref(characteristic + 28, undefined);
      characteristic_ffi["value"] = A.load.Ref(characteristic + 32, undefined);

      const _ret = WEBEXT.bluetoothLowEnergy.createCharacteristic(characteristic_ffi, A.H.get<object>(serviceId));
      A.store.Ref(retPtr, _ret);
    },
    "try_CreateCharacteristic": (
      retPtr: Pointer,
      errPtr: Pointer,
      characteristic: Pointer,
      serviceId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const characteristic_ffi = {};

        characteristic_ffi["uuid"] = A.load.Ref(characteristic + 0, undefined);
        if (A.load.Bool(characteristic + 4 + 17)) {
          characteristic_ffi["service"] = {};
          characteristic_ffi["service"]["uuid"] = A.load.Ref(characteristic + 4 + 0, undefined);
          if (A.load.Bool(characteristic + 4 + 16)) {
            characteristic_ffi["service"]["isPrimary"] = A.load.Bool(characteristic + 4 + 4);
          }
          characteristic_ffi["service"]["instanceId"] = A.load.Ref(characteristic + 4 + 8, undefined);
          characteristic_ffi["service"]["deviceAddress"] = A.load.Ref(characteristic + 4 + 12, undefined);
        }
        characteristic_ffi["properties"] = A.load.Ref(characteristic + 24, undefined);
        characteristic_ffi["instanceId"] = A.load.Ref(characteristic + 28, undefined);
        characteristic_ffi["value"] = A.load.Ref(characteristic + 32, undefined);

        const _ret = WEBEXT.bluetoothLowEnergy.createCharacteristic(characteristic_ffi, A.H.get<object>(serviceId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CreateDescriptor": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "createDescriptor" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CreateDescriptor": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.createDescriptor);
    },
    "call_CreateDescriptor": (retPtr: Pointer, descriptor: Pointer, characteristicId: heap.Ref<object>): void => {
      const descriptor_ffi = {};

      descriptor_ffi["uuid"] = A.load.Ref(descriptor + 0, undefined);
      if (A.load.Bool(descriptor + 4 + 36)) {
        descriptor_ffi["characteristic"] = {};
        descriptor_ffi["characteristic"]["uuid"] = A.load.Ref(descriptor + 4 + 0, undefined);
        if (A.load.Bool(descriptor + 4 + 4 + 17)) {
          descriptor_ffi["characteristic"]["service"] = {};
          descriptor_ffi["characteristic"]["service"]["uuid"] = A.load.Ref(descriptor + 4 + 4 + 0, undefined);
          if (A.load.Bool(descriptor + 4 + 4 + 16)) {
            descriptor_ffi["characteristic"]["service"]["isPrimary"] = A.load.Bool(descriptor + 4 + 4 + 4);
          }
          descriptor_ffi["characteristic"]["service"]["instanceId"] = A.load.Ref(descriptor + 4 + 4 + 8, undefined);
          descriptor_ffi["characteristic"]["service"]["deviceAddress"] = A.load.Ref(descriptor + 4 + 4 + 12, undefined);
        }
        descriptor_ffi["characteristic"]["properties"] = A.load.Ref(descriptor + 4 + 24, undefined);
        descriptor_ffi["characteristic"]["instanceId"] = A.load.Ref(descriptor + 4 + 28, undefined);
        descriptor_ffi["characteristic"]["value"] = A.load.Ref(descriptor + 4 + 32, undefined);
      }
      descriptor_ffi["permissions"] = A.load.Ref(descriptor + 44, undefined);
      descriptor_ffi["instanceId"] = A.load.Ref(descriptor + 48, undefined);
      descriptor_ffi["value"] = A.load.Ref(descriptor + 52, undefined);

      const _ret = WEBEXT.bluetoothLowEnergy.createDescriptor(descriptor_ffi, A.H.get<object>(characteristicId));
      A.store.Ref(retPtr, _ret);
    },
    "try_CreateDescriptor": (
      retPtr: Pointer,
      errPtr: Pointer,
      descriptor: Pointer,
      characteristicId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const descriptor_ffi = {};

        descriptor_ffi["uuid"] = A.load.Ref(descriptor + 0, undefined);
        if (A.load.Bool(descriptor + 4 + 36)) {
          descriptor_ffi["characteristic"] = {};
          descriptor_ffi["characteristic"]["uuid"] = A.load.Ref(descriptor + 4 + 0, undefined);
          if (A.load.Bool(descriptor + 4 + 4 + 17)) {
            descriptor_ffi["characteristic"]["service"] = {};
            descriptor_ffi["characteristic"]["service"]["uuid"] = A.load.Ref(descriptor + 4 + 4 + 0, undefined);
            if (A.load.Bool(descriptor + 4 + 4 + 16)) {
              descriptor_ffi["characteristic"]["service"]["isPrimary"] = A.load.Bool(descriptor + 4 + 4 + 4);
            }
            descriptor_ffi["characteristic"]["service"]["instanceId"] = A.load.Ref(descriptor + 4 + 4 + 8, undefined);
            descriptor_ffi["characteristic"]["service"]["deviceAddress"] = A.load.Ref(
              descriptor + 4 + 4 + 12,
              undefined
            );
          }
          descriptor_ffi["characteristic"]["properties"] = A.load.Ref(descriptor + 4 + 24, undefined);
          descriptor_ffi["characteristic"]["instanceId"] = A.load.Ref(descriptor + 4 + 28, undefined);
          descriptor_ffi["characteristic"]["value"] = A.load.Ref(descriptor + 4 + 32, undefined);
        }
        descriptor_ffi["permissions"] = A.load.Ref(descriptor + 44, undefined);
        descriptor_ffi["instanceId"] = A.load.Ref(descriptor + 48, undefined);
        descriptor_ffi["value"] = A.load.Ref(descriptor + 52, undefined);

        const _ret = WEBEXT.bluetoothLowEnergy.createDescriptor(descriptor_ffi, A.H.get<object>(characteristicId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CreateService": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "createService" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CreateService": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.createService);
    },
    "call_CreateService": (retPtr: Pointer, service: Pointer): void => {
      const service_ffi = {};

      service_ffi["uuid"] = A.load.Ref(service + 0, undefined);
      if (A.load.Bool(service + 16)) {
        service_ffi["isPrimary"] = A.load.Bool(service + 4);
      }
      service_ffi["instanceId"] = A.load.Ref(service + 8, undefined);
      service_ffi["deviceAddress"] = A.load.Ref(service + 12, undefined);

      const _ret = WEBEXT.bluetoothLowEnergy.createService(service_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_CreateService": (retPtr: Pointer, errPtr: Pointer, service: Pointer): heap.Ref<boolean> => {
      try {
        const service_ffi = {};

        service_ffi["uuid"] = A.load.Ref(service + 0, undefined);
        if (A.load.Bool(service + 16)) {
          service_ffi["isPrimary"] = A.load.Bool(service + 4);
        }
        service_ffi["instanceId"] = A.load.Ref(service + 8, undefined);
        service_ffi["deviceAddress"] = A.load.Ref(service + 12, undefined);

        const _ret = WEBEXT.bluetoothLowEnergy.createService(service_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Disconnect": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "disconnect" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Disconnect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.disconnect);
    },
    "call_Disconnect": (retPtr: Pointer, deviceAddress: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.disconnect(A.H.get<object>(deviceAddress));
      A.store.Ref(retPtr, _ret);
    },
    "try_Disconnect": (retPtr: Pointer, errPtr: Pointer, deviceAddress: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.disconnect(A.H.get<object>(deviceAddress));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCharacteristic": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "getCharacteristic" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCharacteristic": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.getCharacteristic);
    },
    "call_GetCharacteristic": (retPtr: Pointer, characteristicId: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.getCharacteristic(A.H.get<object>(characteristicId));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCharacteristic": (
      retPtr: Pointer,
      errPtr: Pointer,
      characteristicId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.getCharacteristic(A.H.get<object>(characteristicId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCharacteristics": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "getCharacteristics" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCharacteristics": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.getCharacteristics);
    },
    "call_GetCharacteristics": (retPtr: Pointer, serviceId: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.getCharacteristics(A.H.get<object>(serviceId));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCharacteristics": (retPtr: Pointer, errPtr: Pointer, serviceId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.getCharacteristics(A.H.get<object>(serviceId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDescriptor": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "getDescriptor" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDescriptor": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.getDescriptor);
    },
    "call_GetDescriptor": (retPtr: Pointer, descriptorId: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.getDescriptor(A.H.get<object>(descriptorId));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDescriptor": (retPtr: Pointer, errPtr: Pointer, descriptorId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.getDescriptor(A.H.get<object>(descriptorId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDescriptors": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "getDescriptors" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDescriptors": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.getDescriptors);
    },
    "call_GetDescriptors": (retPtr: Pointer, characteristicId: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.getDescriptors(A.H.get<object>(characteristicId));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDescriptors": (retPtr: Pointer, errPtr: Pointer, characteristicId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.getDescriptors(A.H.get<object>(characteristicId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetIncludedServices": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "getIncludedServices" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetIncludedServices": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.getIncludedServices);
    },
    "call_GetIncludedServices": (retPtr: Pointer, serviceId: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.getIncludedServices(A.H.get<object>(serviceId));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetIncludedServices": (retPtr: Pointer, errPtr: Pointer, serviceId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.getIncludedServices(A.H.get<object>(serviceId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetService": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "getService" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetService": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.getService);
    },
    "call_GetService": (retPtr: Pointer, serviceId: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.getService(A.H.get<object>(serviceId));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetService": (retPtr: Pointer, errPtr: Pointer, serviceId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.getService(A.H.get<object>(serviceId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetServices": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "getServices" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetServices": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.getServices);
    },
    "call_GetServices": (retPtr: Pointer, deviceAddress: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.getServices(A.H.get<object>(deviceAddress));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetServices": (retPtr: Pointer, errPtr: Pointer, deviceAddress: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.getServices(A.H.get<object>(deviceAddress));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_NotifyCharacteristicValueChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "notifyCharacteristicValueChanged" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_NotifyCharacteristicValueChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.notifyCharacteristicValueChanged);
    },
    "call_NotifyCharacteristicValueChanged": (
      retPtr: Pointer,
      characteristicId: heap.Ref<object>,
      notification: Pointer
    ): void => {
      const notification_ffi = {};

      notification_ffi["value"] = A.load.Ref(notification + 0, undefined);
      if (A.load.Bool(notification + 5)) {
        notification_ffi["shouldIndicate"] = A.load.Bool(notification + 4);
      }

      const _ret = WEBEXT.bluetoothLowEnergy.notifyCharacteristicValueChanged(
        A.H.get<object>(characteristicId),
        notification_ffi
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_NotifyCharacteristicValueChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      characteristicId: heap.Ref<object>,
      notification: Pointer
    ): heap.Ref<boolean> => {
      try {
        const notification_ffi = {};

        notification_ffi["value"] = A.load.Ref(notification + 0, undefined);
        if (A.load.Bool(notification + 5)) {
          notification_ffi["shouldIndicate"] = A.load.Bool(notification + 4);
        }

        const _ret = WEBEXT.bluetoothLowEnergy.notifyCharacteristicValueChanged(
          A.H.get<object>(characteristicId),
          notification_ffi
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCharacteristicReadRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onCharacteristicReadRequest &&
        "addListener" in WEBEXT?.bluetoothLowEnergy?.onCharacteristicReadRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCharacteristicReadRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.addListener);
    },
    "call_OnCharacteristicReadRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.addListener(A.H.get<object>(callback));
    },
    "try_OnCharacteristicReadRequest": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCharacteristicReadRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onCharacteristicReadRequest &&
        "removeListener" in WEBEXT?.bluetoothLowEnergy?.onCharacteristicReadRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCharacteristicReadRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.removeListener);
    },
    "call_OffCharacteristicReadRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.removeListener(A.H.get<object>(callback));
    },
    "try_OffCharacteristicReadRequest": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCharacteristicReadRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onCharacteristicReadRequest &&
        "hasListener" in WEBEXT?.bluetoothLowEnergy?.onCharacteristicReadRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCharacteristicReadRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.hasListener);
    },
    "call_HasOnCharacteristicReadRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCharacteristicReadRequest": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCharacteristicValueChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onCharacteristicValueChanged &&
        "addListener" in WEBEXT?.bluetoothLowEnergy?.onCharacteristicValueChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCharacteristicValueChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.addListener);
    },
    "call_OnCharacteristicValueChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnCharacteristicValueChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCharacteristicValueChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onCharacteristicValueChanged &&
        "removeListener" in WEBEXT?.bluetoothLowEnergy?.onCharacteristicValueChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCharacteristicValueChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.removeListener);
    },
    "call_OffCharacteristicValueChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffCharacteristicValueChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCharacteristicValueChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onCharacteristicValueChanged &&
        "hasListener" in WEBEXT?.bluetoothLowEnergy?.onCharacteristicValueChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCharacteristicValueChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.hasListener);
    },
    "call_HasOnCharacteristicValueChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCharacteristicValueChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCharacteristicWriteRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onCharacteristicWriteRequest &&
        "addListener" in WEBEXT?.bluetoothLowEnergy?.onCharacteristicWriteRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCharacteristicWriteRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.addListener);
    },
    "call_OnCharacteristicWriteRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.addListener(A.H.get<object>(callback));
    },
    "try_OnCharacteristicWriteRequest": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCharacteristicWriteRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onCharacteristicWriteRequest &&
        "removeListener" in WEBEXT?.bluetoothLowEnergy?.onCharacteristicWriteRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCharacteristicWriteRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.removeListener);
    },
    "call_OffCharacteristicWriteRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.removeListener(A.H.get<object>(callback));
    },
    "try_OffCharacteristicWriteRequest": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCharacteristicWriteRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onCharacteristicWriteRequest &&
        "hasListener" in WEBEXT?.bluetoothLowEnergy?.onCharacteristicWriteRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCharacteristicWriteRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.hasListener);
    },
    "call_HasOnCharacteristicWriteRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCharacteristicWriteRequest": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDescriptorReadRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onDescriptorReadRequest &&
        "addListener" in WEBEXT?.bluetoothLowEnergy?.onDescriptorReadRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDescriptorReadRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.addListener);
    },
    "call_OnDescriptorReadRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.addListener(A.H.get<object>(callback));
    },
    "try_OnDescriptorReadRequest": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDescriptorReadRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onDescriptorReadRequest &&
        "removeListener" in WEBEXT?.bluetoothLowEnergy?.onDescriptorReadRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDescriptorReadRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.removeListener);
    },
    "call_OffDescriptorReadRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.removeListener(A.H.get<object>(callback));
    },
    "try_OffDescriptorReadRequest": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDescriptorReadRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onDescriptorReadRequest &&
        "hasListener" in WEBEXT?.bluetoothLowEnergy?.onDescriptorReadRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDescriptorReadRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.hasListener);
    },
    "call_HasOnDescriptorReadRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDescriptorReadRequest": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDescriptorValueChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onDescriptorValueChanged &&
        "addListener" in WEBEXT?.bluetoothLowEnergy?.onDescriptorValueChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDescriptorValueChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.addListener);
    },
    "call_OnDescriptorValueChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnDescriptorValueChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDescriptorValueChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onDescriptorValueChanged &&
        "removeListener" in WEBEXT?.bluetoothLowEnergy?.onDescriptorValueChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDescriptorValueChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.removeListener);
    },
    "call_OffDescriptorValueChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffDescriptorValueChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDescriptorValueChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onDescriptorValueChanged &&
        "hasListener" in WEBEXT?.bluetoothLowEnergy?.onDescriptorValueChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDescriptorValueChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.hasListener);
    },
    "call_HasOnDescriptorValueChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDescriptorValueChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDescriptorWriteRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onDescriptorWriteRequest &&
        "addListener" in WEBEXT?.bluetoothLowEnergy?.onDescriptorWriteRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDescriptorWriteRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.addListener);
    },
    "call_OnDescriptorWriteRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.addListener(A.H.get<object>(callback));
    },
    "try_OnDescriptorWriteRequest": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDescriptorWriteRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onDescriptorWriteRequest &&
        "removeListener" in WEBEXT?.bluetoothLowEnergy?.onDescriptorWriteRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDescriptorWriteRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.removeListener);
    },
    "call_OffDescriptorWriteRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.removeListener(A.H.get<object>(callback));
    },
    "try_OffDescriptorWriteRequest": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDescriptorWriteRequest": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onDescriptorWriteRequest &&
        "hasListener" in WEBEXT?.bluetoothLowEnergy?.onDescriptorWriteRequest
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDescriptorWriteRequest": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.hasListener);
    },
    "call_HasOnDescriptorWriteRequest": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDescriptorWriteRequest": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnServiceAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy?.onServiceAdded && "addListener" in WEBEXT?.bluetoothLowEnergy?.onServiceAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnServiceAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onServiceAdded.addListener);
    },
    "call_OnServiceAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onServiceAdded.addListener(A.H.get<object>(callback));
    },
    "try_OnServiceAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onServiceAdded.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffServiceAdded": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onServiceAdded &&
        "removeListener" in WEBEXT?.bluetoothLowEnergy?.onServiceAdded
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffServiceAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onServiceAdded.removeListener);
    },
    "call_OffServiceAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onServiceAdded.removeListener(A.H.get<object>(callback));
    },
    "try_OffServiceAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onServiceAdded.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnServiceAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy?.onServiceAdded && "hasListener" in WEBEXT?.bluetoothLowEnergy?.onServiceAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnServiceAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onServiceAdded.hasListener);
    },
    "call_HasOnServiceAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onServiceAdded.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnServiceAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onServiceAdded.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnServiceChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onServiceChanged &&
        "addListener" in WEBEXT?.bluetoothLowEnergy?.onServiceChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnServiceChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onServiceChanged.addListener);
    },
    "call_OnServiceChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onServiceChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnServiceChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onServiceChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffServiceChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onServiceChanged &&
        "removeListener" in WEBEXT?.bluetoothLowEnergy?.onServiceChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffServiceChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onServiceChanged.removeListener);
    },
    "call_OffServiceChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onServiceChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffServiceChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onServiceChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnServiceChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onServiceChanged &&
        "hasListener" in WEBEXT?.bluetoothLowEnergy?.onServiceChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnServiceChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onServiceChanged.hasListener);
    },
    "call_HasOnServiceChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onServiceChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnServiceChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onServiceChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnServiceRemoved": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onServiceRemoved &&
        "addListener" in WEBEXT?.bluetoothLowEnergy?.onServiceRemoved
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnServiceRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onServiceRemoved.addListener);
    },
    "call_OnServiceRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onServiceRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnServiceRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onServiceRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffServiceRemoved": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onServiceRemoved &&
        "removeListener" in WEBEXT?.bluetoothLowEnergy?.onServiceRemoved
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffServiceRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onServiceRemoved.removeListener);
    },
    "call_OffServiceRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onServiceRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffServiceRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onServiceRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnServiceRemoved": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bluetoothLowEnergy?.onServiceRemoved &&
        "hasListener" in WEBEXT?.bluetoothLowEnergy?.onServiceRemoved
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnServiceRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.onServiceRemoved.hasListener);
    },
    "call_HasOnServiceRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.onServiceRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnServiceRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.onServiceRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReadCharacteristicValue": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "readCharacteristicValue" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReadCharacteristicValue": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.readCharacteristicValue);
    },
    "call_ReadCharacteristicValue": (retPtr: Pointer, characteristicId: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.readCharacteristicValue(A.H.get<object>(characteristicId));
      A.store.Ref(retPtr, _ret);
    },
    "try_ReadCharacteristicValue": (
      retPtr: Pointer,
      errPtr: Pointer,
      characteristicId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.readCharacteristicValue(A.H.get<object>(characteristicId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReadDescriptorValue": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "readDescriptorValue" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReadDescriptorValue": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.readDescriptorValue);
    },
    "call_ReadDescriptorValue": (retPtr: Pointer, descriptorId: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.readDescriptorValue(A.H.get<object>(descriptorId));
      A.store.Ref(retPtr, _ret);
    },
    "try_ReadDescriptorValue": (
      retPtr: Pointer,
      errPtr: Pointer,
      descriptorId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.readDescriptorValue(A.H.get<object>(descriptorId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RegisterAdvertisement": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "registerAdvertisement" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RegisterAdvertisement": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.registerAdvertisement);
    },
    "call_RegisterAdvertisement": (retPtr: Pointer, advertisement: Pointer): void => {
      const advertisement_ffi = {};

      advertisement_ffi["type"] = A.load.Enum(advertisement + 0, ["broadcast", "peripheral"]);
      advertisement_ffi["serviceUuids"] = A.load.Ref(advertisement + 4, undefined);
      advertisement_ffi["manufacturerData"] = A.load.Ref(advertisement + 8, undefined);
      advertisement_ffi["solicitUuids"] = A.load.Ref(advertisement + 12, undefined);
      advertisement_ffi["serviceData"] = A.load.Ref(advertisement + 16, undefined);

      const _ret = WEBEXT.bluetoothLowEnergy.registerAdvertisement(advertisement_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RegisterAdvertisement": (retPtr: Pointer, errPtr: Pointer, advertisement: Pointer): heap.Ref<boolean> => {
      try {
        const advertisement_ffi = {};

        advertisement_ffi["type"] = A.load.Enum(advertisement + 0, ["broadcast", "peripheral"]);
        advertisement_ffi["serviceUuids"] = A.load.Ref(advertisement + 4, undefined);
        advertisement_ffi["manufacturerData"] = A.load.Ref(advertisement + 8, undefined);
        advertisement_ffi["solicitUuids"] = A.load.Ref(advertisement + 12, undefined);
        advertisement_ffi["serviceData"] = A.load.Ref(advertisement + 16, undefined);

        const _ret = WEBEXT.bluetoothLowEnergy.registerAdvertisement(advertisement_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RegisterService": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "registerService" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RegisterService": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.registerService);
    },
    "call_RegisterService": (retPtr: Pointer, serviceId: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.registerService(A.H.get<object>(serviceId));
      A.store.Ref(retPtr, _ret);
    },
    "try_RegisterService": (retPtr: Pointer, errPtr: Pointer, serviceId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.registerService(A.H.get<object>(serviceId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveService": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "removeService" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveService": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.removeService);
    },
    "call_RemoveService": (retPtr: Pointer, serviceId: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.removeService(A.H.get<object>(serviceId));
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveService": (retPtr: Pointer, errPtr: Pointer, serviceId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.removeService(A.H.get<object>(serviceId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ResetAdvertising": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "resetAdvertising" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ResetAdvertising": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.resetAdvertising);
    },
    "call_ResetAdvertising": (retPtr: Pointer): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.resetAdvertising();
      A.store.Ref(retPtr, _ret);
    },
    "try_ResetAdvertising": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.resetAdvertising();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendRequestResponse": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "sendRequestResponse" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendRequestResponse": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.sendRequestResponse);
    },
    "call_SendRequestResponse": (retPtr: Pointer, response: Pointer): void => {
      const response_ffi = {};

      if (A.load.Bool(response + 12)) {
        response_ffi["requestId"] = A.load.Int32(response + 0);
      }
      if (A.load.Bool(response + 13)) {
        response_ffi["isError"] = A.load.Bool(response + 4);
      }
      response_ffi["value"] = A.load.Ref(response + 8, undefined);

      const _ret = WEBEXT.bluetoothLowEnergy.sendRequestResponse(response_ffi);
    },
    "try_SendRequestResponse": (retPtr: Pointer, errPtr: Pointer, response: Pointer): heap.Ref<boolean> => {
      try {
        const response_ffi = {};

        if (A.load.Bool(response + 12)) {
          response_ffi["requestId"] = A.load.Int32(response + 0);
        }
        if (A.load.Bool(response + 13)) {
          response_ffi["isError"] = A.load.Bool(response + 4);
        }
        response_ffi["value"] = A.load.Ref(response + 8, undefined);

        const _ret = WEBEXT.bluetoothLowEnergy.sendRequestResponse(response_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetAdvertisingInterval": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "setAdvertisingInterval" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetAdvertisingInterval": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.setAdvertisingInterval);
    },
    "call_SetAdvertisingInterval": (retPtr: Pointer, minInterval: number, maxInterval: number): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.setAdvertisingInterval(minInterval, maxInterval);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetAdvertisingInterval": (
      retPtr: Pointer,
      errPtr: Pointer,
      minInterval: number,
      maxInterval: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.setAdvertisingInterval(minInterval, maxInterval);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartCharacteristicNotifications": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "startCharacteristicNotifications" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartCharacteristicNotifications": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.startCharacteristicNotifications);
    },
    "call_StartCharacteristicNotifications": (
      retPtr: Pointer,
      characteristicId: heap.Ref<object>,
      properties: Pointer
    ): void => {
      const properties_ffi = {};

      if (A.load.Bool(properties + 1)) {
        properties_ffi["persistent"] = A.load.Bool(properties + 0);
      }

      const _ret = WEBEXT.bluetoothLowEnergy.startCharacteristicNotifications(
        A.H.get<object>(characteristicId),
        properties_ffi
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_StartCharacteristicNotifications": (
      retPtr: Pointer,
      errPtr: Pointer,
      characteristicId: heap.Ref<object>,
      properties: Pointer
    ): heap.Ref<boolean> => {
      try {
        const properties_ffi = {};

        if (A.load.Bool(properties + 1)) {
          properties_ffi["persistent"] = A.load.Bool(properties + 0);
        }

        const _ret = WEBEXT.bluetoothLowEnergy.startCharacteristicNotifications(
          A.H.get<object>(characteristicId),
          properties_ffi
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StopCharacteristicNotifications": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "stopCharacteristicNotifications" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StopCharacteristicNotifications": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.stopCharacteristicNotifications);
    },
    "call_StopCharacteristicNotifications": (retPtr: Pointer, characteristicId: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.stopCharacteristicNotifications(A.H.get<object>(characteristicId));
      A.store.Ref(retPtr, _ret);
    },
    "try_StopCharacteristicNotifications": (
      retPtr: Pointer,
      errPtr: Pointer,
      characteristicId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.stopCharacteristicNotifications(A.H.get<object>(characteristicId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UnregisterAdvertisement": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "unregisterAdvertisement" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UnregisterAdvertisement": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.unregisterAdvertisement);
    },
    "call_UnregisterAdvertisement": (retPtr: Pointer, advertisementId: number): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.unregisterAdvertisement(advertisementId);
      A.store.Ref(retPtr, _ret);
    },
    "try_UnregisterAdvertisement": (retPtr: Pointer, errPtr: Pointer, advertisementId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.unregisterAdvertisement(advertisementId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UnregisterService": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "unregisterService" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UnregisterService": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.unregisterService);
    },
    "call_UnregisterService": (retPtr: Pointer, serviceId: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.unregisterService(A.H.get<object>(serviceId));
      A.store.Ref(retPtr, _ret);
    },
    "try_UnregisterService": (retPtr: Pointer, errPtr: Pointer, serviceId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.unregisterService(A.H.get<object>(serviceId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_WriteCharacteristicValue": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "writeCharacteristicValue" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_WriteCharacteristicValue": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.writeCharacteristicValue);
    },
    "call_WriteCharacteristicValue": (
      retPtr: Pointer,
      characteristicId: heap.Ref<object>,
      value: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.writeCharacteristicValue(
        A.H.get<object>(characteristicId),
        A.H.get<object>(value)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_WriteCharacteristicValue": (
      retPtr: Pointer,
      errPtr: Pointer,
      characteristicId: heap.Ref<object>,
      value: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.writeCharacteristicValue(
          A.H.get<object>(characteristicId),
          A.H.get<object>(value)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_WriteDescriptorValue": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothLowEnergy && "writeDescriptorValue" in WEBEXT?.bluetoothLowEnergy) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_WriteDescriptorValue": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothLowEnergy.writeDescriptorValue);
    },
    "call_WriteDescriptorValue": (retPtr: Pointer, descriptorId: heap.Ref<object>, value: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothLowEnergy.writeDescriptorValue(
        A.H.get<object>(descriptorId),
        A.H.get<object>(value)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_WriteDescriptorValue": (
      retPtr: Pointer,
      errPtr: Pointer,
      descriptorId: heap.Ref<object>,
      value: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothLowEnergy.writeDescriptorValue(
          A.H.get<object>(descriptorId),
          A.H.get<object>(value)
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
