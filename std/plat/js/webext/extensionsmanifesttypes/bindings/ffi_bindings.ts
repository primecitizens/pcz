import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/extensionsmanifesttypes", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AutomationChoice1": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 1, false);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "desktop" in x ? true : false);
        A.store.Bool(ptr + 0, x["desktop"] ? true : false);
        A.store.Bool(ptr + 9, "interact" in x ? true : false);
        A.store.Bool(ptr + 1, x["interact"] ? true : false);
        A.store.Ref(ptr + 4, x["matches"]);
      }
    },
    "load_AutomationChoice1": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["desktop"] = A.load.Bool(ptr + 0);
      } else {
        delete x["desktop"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["interact"] = A.load.Bool(ptr + 1);
      } else {
        delete x["interact"];
      }
      x["matches"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Bluetooth": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 2, false);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 11, true);
        A.store.Bool(ptr + 8, "low_energy" in x ? true : false);
        A.store.Bool(ptr + 0, x["low_energy"] ? true : false);
        A.store.Bool(ptr + 9, "peripheral" in x ? true : false);
        A.store.Bool(ptr + 1, x["peripheral"] ? true : false);
        A.store.Bool(ptr + 10, "socket" in x ? true : false);
        A.store.Bool(ptr + 2, x["socket"] ? true : false);
        A.store.Ref(ptr + 4, x["uuids"]);
      }
    },
    "load_Bluetooth": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["low_energy"] = A.load.Bool(ptr + 0);
      } else {
        delete x["low_energy"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["peripheral"] = A.load.Bool(ptr + 1);
      } else {
        delete x["peripheral"];
      }
      if (A.load.Bool(ptr + 10)) {
        x["socket"] = A.load.Bool(ptr + 2);
      } else {
        delete x["socket"];
      }
      x["uuids"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ContentCapabilities": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["matches"]);
        A.store.Ref(ptr + 4, x["permissions"]);
      }
    },
    "load_ContentCapabilities": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["matches"] = A.load.Ref(ptr + 0, undefined);
      x["permissions"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ExternallyConnectable": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "accepts_tls_channel_id" in x ? true : false);
        A.store.Bool(ptr + 0, x["accepts_tls_channel_id"] ? true : false);
        A.store.Ref(ptr + 4, x["ids"]);
        A.store.Ref(ptr + 8, x["matches"]);
      }
    },
    "load_ExternallyConnectable": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["accepts_tls_channel_id"] = A.load.Bool(ptr + 0);
      } else {
        delete x["accepts_tls_channel_id"];
      }
      x["ids"] = A.load.Ref(ptr + 4, undefined);
      x["matches"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_KioskSecondaryAppsElem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "enabled_on_launch" in x ? true : false);
        A.store.Bool(ptr + 0, x["enabled_on_launch"] ? true : false);
        A.store.Ref(ptr + 4, x["id"]);
      }
    },
    "load_KioskSecondaryAppsElem": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["enabled_on_launch"] = A.load.Bool(ptr + 0);
      } else {
        delete x["enabled_on_launch"];
      }
      x["id"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OptionsUI": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 1, false);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "chrome_style" in x ? true : false);
        A.store.Bool(ptr + 0, x["chrome_style"] ? true : false);
        A.store.Bool(ptr + 9, "open_in_tab" in x ? true : false);
        A.store.Bool(ptr + 1, x["open_in_tab"] ? true : false);
        A.store.Ref(ptr + 4, x["page"]);
      }
    },
    "load_OptionsUI": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["chrome_style"] = A.load.Bool(ptr + 0);
      } else {
        delete x["chrome_style"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["open_in_tab"] = A.load.Bool(ptr + 1);
      } else {
        delete x["open_in_tab"];
      }
      x["page"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SocketsFieldTcp": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["connect"]);
      }
    },
    "load_SocketsFieldTcp": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["connect"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SocketsFieldTcpServer": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["listen"]);
      }
    },
    "load_SocketsFieldTcpServer": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["listen"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SocketsFieldUdp": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["bind"]);
        A.store.Ref(ptr + 4, x["multicastMembership"]);
        A.store.Ref(ptr + 8, x["send"]);
      }
    },
    "load_SocketsFieldUdp": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["bind"] = A.load.Ref(ptr + 0, undefined);
      x["multicastMembership"] = A.load.Ref(ptr + 4, undefined);
      x["send"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Sockets": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 29, false);

        A.store.Bool(ptr + 0 + 4, false);
        A.store.Ref(ptr + 0 + 0, undefined);

        A.store.Bool(ptr + 8 + 4, false);
        A.store.Ref(ptr + 8 + 0, undefined);

        A.store.Bool(ptr + 16 + 12, false);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 4, undefined);
        A.store.Ref(ptr + 16 + 8, undefined);
      } else {
        A.store.Bool(ptr + 29, true);

        if (typeof x["tcp"] === "undefined") {
          A.store.Bool(ptr + 0 + 4, false);
          A.store.Ref(ptr + 0 + 0, undefined);
        } else {
          A.store.Bool(ptr + 0 + 4, true);
          A.store.Ref(ptr + 0 + 0, x["tcp"]["connect"]);
        }

        if (typeof x["tcpServer"] === "undefined") {
          A.store.Bool(ptr + 8 + 4, false);
          A.store.Ref(ptr + 8 + 0, undefined);
        } else {
          A.store.Bool(ptr + 8 + 4, true);
          A.store.Ref(ptr + 8 + 0, x["tcpServer"]["listen"]);
        }

        if (typeof x["udp"] === "undefined") {
          A.store.Bool(ptr + 16 + 12, false);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 4, undefined);
          A.store.Ref(ptr + 16 + 8, undefined);
        } else {
          A.store.Bool(ptr + 16 + 12, true);
          A.store.Ref(ptr + 16 + 0, x["udp"]["bind"]);
          A.store.Ref(ptr + 16 + 4, x["udp"]["multicastMembership"]);
          A.store.Ref(ptr + 16 + 8, x["udp"]["send"]);
        }
      }
    },
    "load_Sockets": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 4)) {
        x["tcp"] = {};
        x["tcp"]["connect"] = A.load.Ref(ptr + 0 + 0, undefined);
      } else {
        delete x["tcp"];
      }
      if (A.load.Bool(ptr + 8 + 4)) {
        x["tcpServer"] = {};
        x["tcpServer"]["listen"] = A.load.Ref(ptr + 8 + 0, undefined);
      } else {
        delete x["tcpServer"];
      }
      if (A.load.Bool(ptr + 16 + 12)) {
        x["udp"] = {};
        x["udp"]["bind"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["udp"]["multicastMembership"] = A.load.Ref(ptr + 16 + 4, undefined);
        x["udp"]["send"] = A.load.Ref(ptr + 16 + 8, undefined);
      } else {
        delete x["udp"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UsbPrintersFieldFiltersElem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 44, false);
        A.store.Bool(ptr + 40, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Bool(ptr + 41, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Bool(ptr + 42, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Bool(ptr + 43, false);
        A.store.Int64(ptr + 24, 0);
        A.store.Int64(ptr + 32, 0);
      } else {
        A.store.Bool(ptr + 44, true);
        A.store.Bool(ptr + 40, "interfaceClass" in x ? true : false);
        A.store.Int64(ptr + 0, x["interfaceClass"] === undefined ? 0 : (x["interfaceClass"] as number));
        A.store.Bool(ptr + 41, "interfaceProtocol" in x ? true : false);
        A.store.Int64(ptr + 8, x["interfaceProtocol"] === undefined ? 0 : (x["interfaceProtocol"] as number));
        A.store.Bool(ptr + 42, "interfaceSubclass" in x ? true : false);
        A.store.Int64(ptr + 16, x["interfaceSubclass"] === undefined ? 0 : (x["interfaceSubclass"] as number));
        A.store.Bool(ptr + 43, "productId" in x ? true : false);
        A.store.Int64(ptr + 24, x["productId"] === undefined ? 0 : (x["productId"] as number));
        A.store.Int64(ptr + 32, x["vendorId"] === undefined ? 0 : (x["vendorId"] as number));
      }
    },
    "load_UsbPrintersFieldFiltersElem": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 40)) {
        x["interfaceClass"] = A.load.Int64(ptr + 0);
      } else {
        delete x["interfaceClass"];
      }
      if (A.load.Bool(ptr + 41)) {
        x["interfaceProtocol"] = A.load.Int64(ptr + 8);
      } else {
        delete x["interfaceProtocol"];
      }
      if (A.load.Bool(ptr + 42)) {
        x["interfaceSubclass"] = A.load.Int64(ptr + 16);
      } else {
        delete x["interfaceSubclass"];
      }
      if (A.load.Bool(ptr + 43)) {
        x["productId"] = A.load.Int64(ptr + 24);
      } else {
        delete x["productId"];
      }
      x["vendorId"] = A.load.Int64(ptr + 32);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UsbPrinters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["filters"]);
      }
    },
    "load_UsbPrinters": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["filters"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
  };
});
