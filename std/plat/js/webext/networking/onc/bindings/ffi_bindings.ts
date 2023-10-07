import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/networking/onc", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_ActivationStateType": (ref: heap.Ref<string>): number => {
      const idx = ["Activated", "Activating", "NotActivated", "PartiallyActivated"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_CaptivePortalStatus": (ref: heap.Ref<string>): number => {
      const idx = ["Unknown", "Offline", "Online", "Portal", "ProxyAuthRequired"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_FoundNetworkProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Ref(ptr + 0, x["Status"]);
        A.store.Ref(ptr + 4, x["NetworkId"]);
        A.store.Ref(ptr + 8, x["Technology"]);
        A.store.Ref(ptr + 12, x["ShortName"]);
        A.store.Ref(ptr + 16, x["LongName"]);
      }
    },
    "load_FoundNetworkProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["Status"] = A.load.Ref(ptr + 0, undefined);
      x["NetworkId"] = A.load.Ref(ptr + 4, undefined);
      x["Technology"] = A.load.Ref(ptr + 8, undefined);
      x["ShortName"] = A.load.Ref(ptr + 12, undefined);
      x["LongName"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CellularProviderProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["Name"]);
        A.store.Ref(ptr + 4, x["Code"]);
        A.store.Ref(ptr + 8, x["Country"]);
      }
    },
    "load_CellularProviderProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["Name"] = A.load.Ref(ptr + 0, undefined);
      x["Code"] = A.load.Ref(ptr + 4, undefined);
      x["Country"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PaymentPortal": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["Method"]);
        A.store.Ref(ptr + 4, x["PostData"]);
        A.store.Ref(ptr + 8, x["Url"]);
      }
    },
    "load_PaymentPortal": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["Method"] = A.load.Ref(ptr + 0, undefined);
      x["PostData"] = A.load.Ref(ptr + 4, undefined);
      x["Url"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SIMLockStatus": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Ref(ptr + 0, x["LockType"]);
        A.store.Bool(ptr + 12, "LockEnabled" in x ? true : false);
        A.store.Bool(ptr + 4, x["LockEnabled"] ? true : false);
        A.store.Bool(ptr + 13, "RetriesLeft" in x ? true : false);
        A.store.Int32(ptr + 8, x["RetriesLeft"] === undefined ? 0 : (x["RetriesLeft"] as number));
      }
    },
    "load_SIMLockStatus": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["LockType"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["LockEnabled"] = A.load.Bool(ptr + 4);
      } else {
        delete x["LockEnabled"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["RetriesLeft"] = A.load.Int32(ptr + 8);
      } else {
        delete x["RetriesLeft"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CellularProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 127, false);
        A.store.Bool(ptr + 121, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Bool(ptr + 122, false);
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);

        A.store.Bool(ptr + 32 + 12, false);
        A.store.Ref(ptr + 32 + 0, undefined);
        A.store.Ref(ptr + 32 + 4, undefined);
        A.store.Ref(ptr + 32 + 8, undefined);
        A.store.Ref(ptr + 48, undefined);
        A.store.Ref(ptr + 52, undefined);
        A.store.Ref(ptr + 56, undefined);

        A.store.Bool(ptr + 60 + 12, false);
        A.store.Ref(ptr + 60 + 0, undefined);
        A.store.Ref(ptr + 60 + 4, undefined);
        A.store.Ref(ptr + 60 + 8, undefined);
        A.store.Ref(ptr + 76, undefined);
        A.store.Bool(ptr + 123, false);
        A.store.Bool(ptr + 80, false);

        A.store.Bool(ptr + 84 + 12, false);
        A.store.Ref(ptr + 84 + 0, undefined);
        A.store.Ref(ptr + 84 + 4, undefined);
        A.store.Ref(ptr + 84 + 8, undefined);

        A.store.Bool(ptr + 100 + 14, false);
        A.store.Ref(ptr + 100 + 0, undefined);
        A.store.Bool(ptr + 100 + 12, false);
        A.store.Bool(ptr + 100 + 4, false);
        A.store.Bool(ptr + 100 + 13, false);
        A.store.Int32(ptr + 100 + 8, 0);
        A.store.Bool(ptr + 124, false);
        A.store.Bool(ptr + 115, false);
        A.store.Bool(ptr + 125, false);
        A.store.Int32(ptr + 116, 0);
        A.store.Bool(ptr + 126, false);
        A.store.Bool(ptr + 120, false);
      } else {
        A.store.Bool(ptr + 127, true);
        A.store.Bool(ptr + 121, "AutoConnect" in x ? true : false);
        A.store.Bool(ptr + 0, x["AutoConnect"] ? true : false);
        A.store.Ref(ptr + 4, x["ActivationType"]);
        A.store.Enum(
          ptr + 8,
          ["Activated", "Activating", "NotActivated", "PartiallyActivated"].indexOf(x["ActivationState"] as string)
        );
        A.store.Bool(ptr + 122, "AllowRoaming" in x ? true : false);
        A.store.Bool(ptr + 12, x["AllowRoaming"] ? true : false);
        A.store.Ref(ptr + 16, x["Family"]);
        A.store.Ref(ptr + 20, x["FirmwareRevision"]);
        A.store.Ref(ptr + 24, x["FoundNetworks"]);
        A.store.Ref(ptr + 28, x["HardwareRevision"]);

        if (typeof x["HomeProvider"] === "undefined") {
          A.store.Bool(ptr + 32 + 12, false);
          A.store.Ref(ptr + 32 + 0, undefined);
          A.store.Ref(ptr + 32 + 4, undefined);
          A.store.Ref(ptr + 32 + 8, undefined);
        } else {
          A.store.Bool(ptr + 32 + 12, true);
          A.store.Ref(ptr + 32 + 0, x["HomeProvider"]["Name"]);
          A.store.Ref(ptr + 32 + 4, x["HomeProvider"]["Code"]);
          A.store.Ref(ptr + 32 + 8, x["HomeProvider"]["Country"]);
        }
        A.store.Ref(ptr + 48, x["Manufacturer"]);
        A.store.Ref(ptr + 52, x["ModelID"]);
        A.store.Ref(ptr + 56, x["NetworkTechnology"]);

        if (typeof x["PaymentPortal"] === "undefined") {
          A.store.Bool(ptr + 60 + 12, false);
          A.store.Ref(ptr + 60 + 0, undefined);
          A.store.Ref(ptr + 60 + 4, undefined);
          A.store.Ref(ptr + 60 + 8, undefined);
        } else {
          A.store.Bool(ptr + 60 + 12, true);
          A.store.Ref(ptr + 60 + 0, x["PaymentPortal"]["Method"]);
          A.store.Ref(ptr + 60 + 4, x["PaymentPortal"]["PostData"]);
          A.store.Ref(ptr + 60 + 8, x["PaymentPortal"]["Url"]);
        }
        A.store.Ref(ptr + 76, x["RoamingState"]);
        A.store.Bool(ptr + 123, "Scanning" in x ? true : false);
        A.store.Bool(ptr + 80, x["Scanning"] ? true : false);

        if (typeof x["ServingOperator"] === "undefined") {
          A.store.Bool(ptr + 84 + 12, false);
          A.store.Ref(ptr + 84 + 0, undefined);
          A.store.Ref(ptr + 84 + 4, undefined);
          A.store.Ref(ptr + 84 + 8, undefined);
        } else {
          A.store.Bool(ptr + 84 + 12, true);
          A.store.Ref(ptr + 84 + 0, x["ServingOperator"]["Name"]);
          A.store.Ref(ptr + 84 + 4, x["ServingOperator"]["Code"]);
          A.store.Ref(ptr + 84 + 8, x["ServingOperator"]["Country"]);
        }

        if (typeof x["SIMLockStatus"] === "undefined") {
          A.store.Bool(ptr + 100 + 14, false);
          A.store.Ref(ptr + 100 + 0, undefined);
          A.store.Bool(ptr + 100 + 12, false);
          A.store.Bool(ptr + 100 + 4, false);
          A.store.Bool(ptr + 100 + 13, false);
          A.store.Int32(ptr + 100 + 8, 0);
        } else {
          A.store.Bool(ptr + 100 + 14, true);
          A.store.Ref(ptr + 100 + 0, x["SIMLockStatus"]["LockType"]);
          A.store.Bool(ptr + 100 + 12, "LockEnabled" in x["SIMLockStatus"] ? true : false);
          A.store.Bool(ptr + 100 + 4, x["SIMLockStatus"]["LockEnabled"] ? true : false);
          A.store.Bool(ptr + 100 + 13, "RetriesLeft" in x["SIMLockStatus"] ? true : false);
          A.store.Int32(
            ptr + 100 + 8,
            x["SIMLockStatus"]["RetriesLeft"] === undefined ? 0 : (x["SIMLockStatus"]["RetriesLeft"] as number)
          );
        }
        A.store.Bool(ptr + 124, "SIMPresent" in x ? true : false);
        A.store.Bool(ptr + 115, x["SIMPresent"] ? true : false);
        A.store.Bool(ptr + 125, "SignalStrength" in x ? true : false);
        A.store.Int32(ptr + 116, x["SignalStrength"] === undefined ? 0 : (x["SignalStrength"] as number));
        A.store.Bool(ptr + 126, "SupportNetworkScan" in x ? true : false);
        A.store.Bool(ptr + 120, x["SupportNetworkScan"] ? true : false);
      }
    },
    "load_CellularProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 121)) {
        x["AutoConnect"] = A.load.Bool(ptr + 0);
      } else {
        delete x["AutoConnect"];
      }
      x["ActivationType"] = A.load.Ref(ptr + 4, undefined);
      x["ActivationState"] = A.load.Enum(ptr + 8, ["Activated", "Activating", "NotActivated", "PartiallyActivated"]);
      if (A.load.Bool(ptr + 122)) {
        x["AllowRoaming"] = A.load.Bool(ptr + 12);
      } else {
        delete x["AllowRoaming"];
      }
      x["Family"] = A.load.Ref(ptr + 16, undefined);
      x["FirmwareRevision"] = A.load.Ref(ptr + 20, undefined);
      x["FoundNetworks"] = A.load.Ref(ptr + 24, undefined);
      x["HardwareRevision"] = A.load.Ref(ptr + 28, undefined);
      if (A.load.Bool(ptr + 32 + 12)) {
        x["HomeProvider"] = {};
        x["HomeProvider"]["Name"] = A.load.Ref(ptr + 32 + 0, undefined);
        x["HomeProvider"]["Code"] = A.load.Ref(ptr + 32 + 4, undefined);
        x["HomeProvider"]["Country"] = A.load.Ref(ptr + 32 + 8, undefined);
      } else {
        delete x["HomeProvider"];
      }
      x["Manufacturer"] = A.load.Ref(ptr + 48, undefined);
      x["ModelID"] = A.load.Ref(ptr + 52, undefined);
      x["NetworkTechnology"] = A.load.Ref(ptr + 56, undefined);
      if (A.load.Bool(ptr + 60 + 12)) {
        x["PaymentPortal"] = {};
        x["PaymentPortal"]["Method"] = A.load.Ref(ptr + 60 + 0, undefined);
        x["PaymentPortal"]["PostData"] = A.load.Ref(ptr + 60 + 4, undefined);
        x["PaymentPortal"]["Url"] = A.load.Ref(ptr + 60 + 8, undefined);
      } else {
        delete x["PaymentPortal"];
      }
      x["RoamingState"] = A.load.Ref(ptr + 76, undefined);
      if (A.load.Bool(ptr + 123)) {
        x["Scanning"] = A.load.Bool(ptr + 80);
      } else {
        delete x["Scanning"];
      }
      if (A.load.Bool(ptr + 84 + 12)) {
        x["ServingOperator"] = {};
        x["ServingOperator"]["Name"] = A.load.Ref(ptr + 84 + 0, undefined);
        x["ServingOperator"]["Code"] = A.load.Ref(ptr + 84 + 4, undefined);
        x["ServingOperator"]["Country"] = A.load.Ref(ptr + 84 + 8, undefined);
      } else {
        delete x["ServingOperator"];
      }
      if (A.load.Bool(ptr + 100 + 14)) {
        x["SIMLockStatus"] = {};
        x["SIMLockStatus"]["LockType"] = A.load.Ref(ptr + 100 + 0, undefined);
        if (A.load.Bool(ptr + 100 + 12)) {
          x["SIMLockStatus"]["LockEnabled"] = A.load.Bool(ptr + 100 + 4);
        } else {
          delete x["SIMLockStatus"]["LockEnabled"];
        }
        if (A.load.Bool(ptr + 100 + 13)) {
          x["SIMLockStatus"]["RetriesLeft"] = A.load.Int32(ptr + 100 + 8);
        } else {
          delete x["SIMLockStatus"]["RetriesLeft"];
        }
      } else {
        delete x["SIMLockStatus"];
      }
      if (A.load.Bool(ptr + 124)) {
        x["SIMPresent"] = A.load.Bool(ptr + 115);
      } else {
        delete x["SIMPresent"];
      }
      if (A.load.Bool(ptr + 125)) {
        x["SignalStrength"] = A.load.Int32(ptr + 116);
      } else {
        delete x["SignalStrength"];
      }
      if (A.load.Bool(ptr + 126)) {
        x["SupportNetworkScan"] = A.load.Bool(ptr + 120);
      } else {
        delete x["SupportNetworkScan"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CellularStateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 22, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 21, false);
        A.store.Int32(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 22, true);
        A.store.Enum(
          ptr + 0,
          ["Activated", "Activating", "NotActivated", "PartiallyActivated"].indexOf(x["ActivationState"] as string)
        );
        A.store.Ref(ptr + 4, x["NetworkTechnology"]);
        A.store.Ref(ptr + 8, x["RoamingState"]);
        A.store.Bool(ptr + 20, "SIMPresent" in x ? true : false);
        A.store.Bool(ptr + 12, x["SIMPresent"] ? true : false);
        A.store.Bool(ptr + 21, "SignalStrength" in x ? true : false);
        A.store.Int32(ptr + 16, x["SignalStrength"] === undefined ? 0 : (x["SignalStrength"] as number));
      }
    },
    "load_CellularStateProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["ActivationState"] = A.load.Enum(ptr + 0, ["Activated", "Activating", "NotActivated", "PartiallyActivated"]);
      x["NetworkTechnology"] = A.load.Ref(ptr + 4, undefined);
      x["RoamingState"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 20)) {
        x["SIMPresent"] = A.load.Bool(ptr + 12);
      } else {
        delete x["SIMPresent"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["SignalStrength"] = A.load.Int32(ptr + 16);
      } else {
        delete x["SignalStrength"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_IssuerSubjectPattern": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["CommonName"]);
        A.store.Ref(ptr + 4, x["Locality"]);
        A.store.Ref(ptr + 8, x["Organization"]);
        A.store.Ref(ptr + 12, x["OrganizationalUnit"]);
      }
    },
    "load_IssuerSubjectPattern": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["CommonName"] = A.load.Ref(ptr + 0, undefined);
      x["Locality"] = A.load.Ref(ptr + 4, undefined);
      x["Organization"] = A.load.Ref(ptr + 8, undefined);
      x["OrganizationalUnit"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CertificatePattern": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 45, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 4 + 16, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Ref(ptr + 4 + 4, undefined);
        A.store.Ref(ptr + 4 + 8, undefined);
        A.store.Ref(ptr + 4 + 12, undefined);
        A.store.Ref(ptr + 24, undefined);

        A.store.Bool(ptr + 28 + 16, false);
        A.store.Ref(ptr + 28 + 0, undefined);
        A.store.Ref(ptr + 28 + 4, undefined);
        A.store.Ref(ptr + 28 + 8, undefined);
        A.store.Ref(ptr + 28 + 12, undefined);
      } else {
        A.store.Bool(ptr + 45, true);
        A.store.Ref(ptr + 0, x["EnrollmentURI"]);

        if (typeof x["Issuer"] === "undefined") {
          A.store.Bool(ptr + 4 + 16, false);
          A.store.Ref(ptr + 4 + 0, undefined);
          A.store.Ref(ptr + 4 + 4, undefined);
          A.store.Ref(ptr + 4 + 8, undefined);
          A.store.Ref(ptr + 4 + 12, undefined);
        } else {
          A.store.Bool(ptr + 4 + 16, true);
          A.store.Ref(ptr + 4 + 0, x["Issuer"]["CommonName"]);
          A.store.Ref(ptr + 4 + 4, x["Issuer"]["Locality"]);
          A.store.Ref(ptr + 4 + 8, x["Issuer"]["Organization"]);
          A.store.Ref(ptr + 4 + 12, x["Issuer"]["OrganizationalUnit"]);
        }
        A.store.Ref(ptr + 24, x["IssuerCARef"]);

        if (typeof x["Subject"] === "undefined") {
          A.store.Bool(ptr + 28 + 16, false);
          A.store.Ref(ptr + 28 + 0, undefined);
          A.store.Ref(ptr + 28 + 4, undefined);
          A.store.Ref(ptr + 28 + 8, undefined);
          A.store.Ref(ptr + 28 + 12, undefined);
        } else {
          A.store.Bool(ptr + 28 + 16, true);
          A.store.Ref(ptr + 28 + 0, x["Subject"]["CommonName"]);
          A.store.Ref(ptr + 28 + 4, x["Subject"]["Locality"]);
          A.store.Ref(ptr + 28 + 8, x["Subject"]["Organization"]);
          A.store.Ref(ptr + 28 + 12, x["Subject"]["OrganizationalUnit"]);
        }
      }
    },
    "load_CertificatePattern": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["EnrollmentURI"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 4 + 16)) {
        x["Issuer"] = {};
        x["Issuer"]["CommonName"] = A.load.Ref(ptr + 4 + 0, undefined);
        x["Issuer"]["Locality"] = A.load.Ref(ptr + 4 + 4, undefined);
        x["Issuer"]["Organization"] = A.load.Ref(ptr + 4 + 8, undefined);
        x["Issuer"]["OrganizationalUnit"] = A.load.Ref(ptr + 4 + 12, undefined);
      } else {
        delete x["Issuer"];
      }
      x["IssuerCARef"] = A.load.Ref(ptr + 24, undefined);
      if (A.load.Bool(ptr + 28 + 16)) {
        x["Subject"] = {};
        x["Subject"]["CommonName"] = A.load.Ref(ptr + 28 + 0, undefined);
        x["Subject"]["Locality"] = A.load.Ref(ptr + 28 + 4, undefined);
        x["Subject"]["Organization"] = A.load.Ref(ptr + 28 + 8, undefined);
        x["Subject"]["OrganizationalUnit"] = A.load.Ref(ptr + 28 + 12, undefined);
      } else {
        delete x["Subject"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ClientCertificateType": (ref: heap.Ref<string>): number => {
      const idx = ["Ref", "Pattern"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ConnectionStateType": (ref: heap.Ref<string>): number => {
      const idx = ["Connected", "Connecting", "NotConnected"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_DeviceStateType": (ref: heap.Ref<string>): number => {
      const idx = ["Uninitialized", "Disabled", "Enabling", "Enabled", "Prohibited"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_NetworkType": (ref: heap.Ref<string>): number => {
      const idx = ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DeviceStateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 30, false);
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 0, false);

        A.store.Bool(ptr + 4 + 14, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Bool(ptr + 4 + 12, false);
        A.store.Bool(ptr + 4 + 4, false);
        A.store.Bool(ptr + 4 + 13, false);
        A.store.Int32(ptr + 4 + 8, 0);
        A.store.Bool(ptr + 29, false);
        A.store.Bool(ptr + 19, false);
        A.store.Enum(ptr + 20, -1);
        A.store.Enum(ptr + 24, -1);
      } else {
        A.store.Bool(ptr + 30, true);
        A.store.Bool(ptr + 28, "Scanning" in x ? true : false);
        A.store.Bool(ptr + 0, x["Scanning"] ? true : false);

        if (typeof x["SIMLockStatus"] === "undefined") {
          A.store.Bool(ptr + 4 + 14, false);
          A.store.Ref(ptr + 4 + 0, undefined);
          A.store.Bool(ptr + 4 + 12, false);
          A.store.Bool(ptr + 4 + 4, false);
          A.store.Bool(ptr + 4 + 13, false);
          A.store.Int32(ptr + 4 + 8, 0);
        } else {
          A.store.Bool(ptr + 4 + 14, true);
          A.store.Ref(ptr + 4 + 0, x["SIMLockStatus"]["LockType"]);
          A.store.Bool(ptr + 4 + 12, "LockEnabled" in x["SIMLockStatus"] ? true : false);
          A.store.Bool(ptr + 4 + 4, x["SIMLockStatus"]["LockEnabled"] ? true : false);
          A.store.Bool(ptr + 4 + 13, "RetriesLeft" in x["SIMLockStatus"] ? true : false);
          A.store.Int32(
            ptr + 4 + 8,
            x["SIMLockStatus"]["RetriesLeft"] === undefined ? 0 : (x["SIMLockStatus"]["RetriesLeft"] as number)
          );
        }
        A.store.Bool(ptr + 29, "SIMPresent" in x ? true : false);
        A.store.Bool(ptr + 19, x["SIMPresent"] ? true : false);
        A.store.Enum(
          ptr + 20,
          ["Uninitialized", "Disabled", "Enabling", "Enabled", "Prohibited"].indexOf(x["State"] as string)
        );
        A.store.Enum(
          ptr + 24,
          ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"].indexOf(x["Type"] as string)
        );
      }
    },
    "load_DeviceStateProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 28)) {
        x["Scanning"] = A.load.Bool(ptr + 0);
      } else {
        delete x["Scanning"];
      }
      if (A.load.Bool(ptr + 4 + 14)) {
        x["SIMLockStatus"] = {};
        x["SIMLockStatus"]["LockType"] = A.load.Ref(ptr + 4 + 0, undefined);
        if (A.load.Bool(ptr + 4 + 12)) {
          x["SIMLockStatus"]["LockEnabled"] = A.load.Bool(ptr + 4 + 4);
        } else {
          delete x["SIMLockStatus"]["LockEnabled"];
        }
        if (A.load.Bool(ptr + 4 + 13)) {
          x["SIMLockStatus"]["RetriesLeft"] = A.load.Int32(ptr + 4 + 8);
        } else {
          delete x["SIMLockStatus"]["RetriesLeft"];
        }
      } else {
        delete x["SIMLockStatus"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["SIMPresent"] = A.load.Bool(ptr + 19);
      } else {
        delete x["SIMPresent"];
      }
      x["State"] = A.load.Enum(ptr + 20, ["Uninitialized", "Disabled", "Enabling", "Enabled", "Prohibited"]);
      x["Type"] = A.load.Enum(ptr + 24, ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManagedDOMString": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 27, false);
        A.store.Bool(ptr + 25, false);
      } else {
        A.store.Bool(ptr + 28, true);
        A.store.Ref(ptr + 0, x["Active"]);
        A.store.Ref(ptr + 4, x["Effective"]);
        A.store.Ref(ptr + 8, x["UserPolicy"]);
        A.store.Ref(ptr + 12, x["DevicePolicy"]);
        A.store.Ref(ptr + 16, x["UserSetting"]);
        A.store.Ref(ptr + 20, x["SharedSetting"]);
        A.store.Bool(ptr + 26, "UserEditable" in x ? true : false);
        A.store.Bool(ptr + 24, x["UserEditable"] ? true : false);
        A.store.Bool(ptr + 27, "DeviceEditable" in x ? true : false);
        A.store.Bool(ptr + 25, x["DeviceEditable"] ? true : false);
      }
    },
    "load_ManagedDOMString": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["Active"] = A.load.Ref(ptr + 0, undefined);
      x["Effective"] = A.load.Ref(ptr + 4, undefined);
      x["UserPolicy"] = A.load.Ref(ptr + 8, undefined);
      x["DevicePolicy"] = A.load.Ref(ptr + 12, undefined);
      x["UserSetting"] = A.load.Ref(ptr + 16, undefined);
      x["SharedSetting"] = A.load.Ref(ptr + 20, undefined);
      if (A.load.Bool(ptr + 26)) {
        x["UserEditable"] = A.load.Bool(ptr + 24);
      } else {
        delete x["UserEditable"];
      }
      if (A.load.Bool(ptr + 27)) {
        x["DeviceEditable"] = A.load.Bool(ptr + 25);
      } else {
        delete x["DeviceEditable"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_EAPProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 130, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 4 + 45, false);
        A.store.Ref(ptr + 4 + 0, undefined);

        A.store.Bool(ptr + 4 + 4 + 16, false);
        A.store.Ref(ptr + 4 + 4 + 0, undefined);
        A.store.Ref(ptr + 4 + 4 + 4, undefined);
        A.store.Ref(ptr + 4 + 4 + 8, undefined);
        A.store.Ref(ptr + 4 + 4 + 12, undefined);
        A.store.Ref(ptr + 4 + 24, undefined);

        A.store.Bool(ptr + 4 + 28 + 16, false);
        A.store.Ref(ptr + 4 + 28 + 0, undefined);
        A.store.Ref(ptr + 4 + 28 + 4, undefined);
        A.store.Ref(ptr + 4 + 28 + 8, undefined);
        A.store.Ref(ptr + 4 + 28 + 12, undefined);
        A.store.Ref(ptr + 52, undefined);
        A.store.Ref(ptr + 56, undefined);
        A.store.Ref(ptr + 60, undefined);
        A.store.Enum(ptr + 64, -1);
        A.store.Ref(ptr + 68, undefined);
        A.store.Ref(ptr + 72, undefined);
        A.store.Ref(ptr + 76, undefined);
        A.store.Ref(ptr + 80, undefined);
        A.store.Bool(ptr + 127, false);
        A.store.Bool(ptr + 84, false);
        A.store.Ref(ptr + 88, undefined);
        A.store.Ref(ptr + 92, undefined);

        A.store.Bool(ptr + 96 + 28, false);
        A.store.Ref(ptr + 96 + 0, undefined);
        A.store.Ref(ptr + 96 + 4, undefined);
        A.store.Ref(ptr + 96 + 8, undefined);
        A.store.Ref(ptr + 96 + 12, undefined);
        A.store.Ref(ptr + 96 + 16, undefined);
        A.store.Ref(ptr + 96 + 20, undefined);
        A.store.Bool(ptr + 96 + 26, false);
        A.store.Bool(ptr + 96 + 24, false);
        A.store.Bool(ptr + 96 + 27, false);
        A.store.Bool(ptr + 96 + 25, false);
        A.store.Bool(ptr + 128, false);
        A.store.Bool(ptr + 125, false);
        A.store.Bool(ptr + 129, false);
        A.store.Bool(ptr + 126, false);
      } else {
        A.store.Bool(ptr + 130, true);
        A.store.Ref(ptr + 0, x["AnonymousIdentity"]);

        if (typeof x["ClientCertPattern"] === "undefined") {
          A.store.Bool(ptr + 4 + 45, false);
          A.store.Ref(ptr + 4 + 0, undefined);

          A.store.Bool(ptr + 4 + 4 + 16, false);
          A.store.Ref(ptr + 4 + 4 + 0, undefined);
          A.store.Ref(ptr + 4 + 4 + 4, undefined);
          A.store.Ref(ptr + 4 + 4 + 8, undefined);
          A.store.Ref(ptr + 4 + 4 + 12, undefined);
          A.store.Ref(ptr + 4 + 24, undefined);

          A.store.Bool(ptr + 4 + 28 + 16, false);
          A.store.Ref(ptr + 4 + 28 + 0, undefined);
          A.store.Ref(ptr + 4 + 28 + 4, undefined);
          A.store.Ref(ptr + 4 + 28 + 8, undefined);
          A.store.Ref(ptr + 4 + 28 + 12, undefined);
        } else {
          A.store.Bool(ptr + 4 + 45, true);
          A.store.Ref(ptr + 4 + 0, x["ClientCertPattern"]["EnrollmentURI"]);

          if (typeof x["ClientCertPattern"]["Issuer"] === "undefined") {
            A.store.Bool(ptr + 4 + 4 + 16, false);
            A.store.Ref(ptr + 4 + 4 + 0, undefined);
            A.store.Ref(ptr + 4 + 4 + 4, undefined);
            A.store.Ref(ptr + 4 + 4 + 8, undefined);
            A.store.Ref(ptr + 4 + 4 + 12, undefined);
          } else {
            A.store.Bool(ptr + 4 + 4 + 16, true);
            A.store.Ref(ptr + 4 + 4 + 0, x["ClientCertPattern"]["Issuer"]["CommonName"]);
            A.store.Ref(ptr + 4 + 4 + 4, x["ClientCertPattern"]["Issuer"]["Locality"]);
            A.store.Ref(ptr + 4 + 4 + 8, x["ClientCertPattern"]["Issuer"]["Organization"]);
            A.store.Ref(ptr + 4 + 4 + 12, x["ClientCertPattern"]["Issuer"]["OrganizationalUnit"]);
          }
          A.store.Ref(ptr + 4 + 24, x["ClientCertPattern"]["IssuerCARef"]);

          if (typeof x["ClientCertPattern"]["Subject"] === "undefined") {
            A.store.Bool(ptr + 4 + 28 + 16, false);
            A.store.Ref(ptr + 4 + 28 + 0, undefined);
            A.store.Ref(ptr + 4 + 28 + 4, undefined);
            A.store.Ref(ptr + 4 + 28 + 8, undefined);
            A.store.Ref(ptr + 4 + 28 + 12, undefined);
          } else {
            A.store.Bool(ptr + 4 + 28 + 16, true);
            A.store.Ref(ptr + 4 + 28 + 0, x["ClientCertPattern"]["Subject"]["CommonName"]);
            A.store.Ref(ptr + 4 + 28 + 4, x["ClientCertPattern"]["Subject"]["Locality"]);
            A.store.Ref(ptr + 4 + 28 + 8, x["ClientCertPattern"]["Subject"]["Organization"]);
            A.store.Ref(ptr + 4 + 28 + 12, x["ClientCertPattern"]["Subject"]["OrganizationalUnit"]);
          }
        }
        A.store.Ref(ptr + 52, x["ClientCertPKCS11Id"]);
        A.store.Ref(ptr + 56, x["ClientCertProvisioningProfileId"]);
        A.store.Ref(ptr + 60, x["ClientCertRef"]);
        A.store.Enum(ptr + 64, ["Ref", "Pattern"].indexOf(x["ClientCertType"] as string));
        A.store.Ref(ptr + 68, x["Identity"]);
        A.store.Ref(ptr + 72, x["Inner"]);
        A.store.Ref(ptr + 76, x["Outer"]);
        A.store.Ref(ptr + 80, x["Password"]);
        A.store.Bool(ptr + 127, "SaveCredentials" in x ? true : false);
        A.store.Bool(ptr + 84, x["SaveCredentials"] ? true : false);
        A.store.Ref(ptr + 88, x["ServerCAPEMs"]);
        A.store.Ref(ptr + 92, x["ServerCARefs"]);

        if (typeof x["SubjectMatch"] === "undefined") {
          A.store.Bool(ptr + 96 + 28, false);
          A.store.Ref(ptr + 96 + 0, undefined);
          A.store.Ref(ptr + 96 + 4, undefined);
          A.store.Ref(ptr + 96 + 8, undefined);
          A.store.Ref(ptr + 96 + 12, undefined);
          A.store.Ref(ptr + 96 + 16, undefined);
          A.store.Ref(ptr + 96 + 20, undefined);
          A.store.Bool(ptr + 96 + 26, false);
          A.store.Bool(ptr + 96 + 24, false);
          A.store.Bool(ptr + 96 + 27, false);
          A.store.Bool(ptr + 96 + 25, false);
        } else {
          A.store.Bool(ptr + 96 + 28, true);
          A.store.Ref(ptr + 96 + 0, x["SubjectMatch"]["Active"]);
          A.store.Ref(ptr + 96 + 4, x["SubjectMatch"]["Effective"]);
          A.store.Ref(ptr + 96 + 8, x["SubjectMatch"]["UserPolicy"]);
          A.store.Ref(ptr + 96 + 12, x["SubjectMatch"]["DevicePolicy"]);
          A.store.Ref(ptr + 96 + 16, x["SubjectMatch"]["UserSetting"]);
          A.store.Ref(ptr + 96 + 20, x["SubjectMatch"]["SharedSetting"]);
          A.store.Bool(ptr + 96 + 26, "UserEditable" in x["SubjectMatch"] ? true : false);
          A.store.Bool(ptr + 96 + 24, x["SubjectMatch"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 96 + 27, "DeviceEditable" in x["SubjectMatch"] ? true : false);
          A.store.Bool(ptr + 96 + 25, x["SubjectMatch"]["DeviceEditable"] ? true : false);
        }
        A.store.Bool(ptr + 128, "UseProactiveKeyCaching" in x ? true : false);
        A.store.Bool(ptr + 125, x["UseProactiveKeyCaching"] ? true : false);
        A.store.Bool(ptr + 129, "UseSystemCAs" in x ? true : false);
        A.store.Bool(ptr + 126, x["UseSystemCAs"] ? true : false);
      }
    },
    "load_EAPProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["AnonymousIdentity"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 4 + 45)) {
        x["ClientCertPattern"] = {};
        x["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(ptr + 4 + 0, undefined);
        if (A.load.Bool(ptr + 4 + 4 + 16)) {
          x["ClientCertPattern"]["Issuer"] = {};
          x["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(ptr + 4 + 4 + 0, undefined);
          x["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(ptr + 4 + 4 + 4, undefined);
          x["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(ptr + 4 + 4 + 8, undefined);
          x["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(ptr + 4 + 4 + 12, undefined);
        } else {
          delete x["ClientCertPattern"]["Issuer"];
        }
        x["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(ptr + 4 + 24, undefined);
        if (A.load.Bool(ptr + 4 + 28 + 16)) {
          x["ClientCertPattern"]["Subject"] = {};
          x["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(ptr + 4 + 28 + 0, undefined);
          x["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(ptr + 4 + 28 + 4, undefined);
          x["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(ptr + 4 + 28 + 8, undefined);
          x["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(ptr + 4 + 28 + 12, undefined);
        } else {
          delete x["ClientCertPattern"]["Subject"];
        }
      } else {
        delete x["ClientCertPattern"];
      }
      x["ClientCertPKCS11Id"] = A.load.Ref(ptr + 52, undefined);
      x["ClientCertProvisioningProfileId"] = A.load.Ref(ptr + 56, undefined);
      x["ClientCertRef"] = A.load.Ref(ptr + 60, undefined);
      x["ClientCertType"] = A.load.Enum(ptr + 64, ["Ref", "Pattern"]);
      x["Identity"] = A.load.Ref(ptr + 68, undefined);
      x["Inner"] = A.load.Ref(ptr + 72, undefined);
      x["Outer"] = A.load.Ref(ptr + 76, undefined);
      x["Password"] = A.load.Ref(ptr + 80, undefined);
      if (A.load.Bool(ptr + 127)) {
        x["SaveCredentials"] = A.load.Bool(ptr + 84);
      } else {
        delete x["SaveCredentials"];
      }
      x["ServerCAPEMs"] = A.load.Ref(ptr + 88, undefined);
      x["ServerCARefs"] = A.load.Ref(ptr + 92, undefined);
      if (A.load.Bool(ptr + 96 + 28)) {
        x["SubjectMatch"] = {};
        x["SubjectMatch"]["Active"] = A.load.Ref(ptr + 96 + 0, undefined);
        x["SubjectMatch"]["Effective"] = A.load.Ref(ptr + 96 + 4, undefined);
        x["SubjectMatch"]["UserPolicy"] = A.load.Ref(ptr + 96 + 8, undefined);
        x["SubjectMatch"]["DevicePolicy"] = A.load.Ref(ptr + 96 + 12, undefined);
        x["SubjectMatch"]["UserSetting"] = A.load.Ref(ptr + 96 + 16, undefined);
        x["SubjectMatch"]["SharedSetting"] = A.load.Ref(ptr + 96 + 20, undefined);
        if (A.load.Bool(ptr + 96 + 26)) {
          x["SubjectMatch"]["UserEditable"] = A.load.Bool(ptr + 96 + 24);
        } else {
          delete x["SubjectMatch"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 96 + 27)) {
          x["SubjectMatch"]["DeviceEditable"] = A.load.Bool(ptr + 96 + 25);
        } else {
          delete x["SubjectMatch"]["DeviceEditable"];
        }
      } else {
        delete x["SubjectMatch"];
      }
      if (A.load.Bool(ptr + 128)) {
        x["UseProactiveKeyCaching"] = A.load.Bool(ptr + 125);
      } else {
        delete x["UseProactiveKeyCaching"];
      }
      if (A.load.Bool(ptr + 129)) {
        x["UseSystemCAs"] = A.load.Bool(ptr + 126);
      } else {
        delete x["UseSystemCAs"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_EthernetProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 140, false);
        A.store.Bool(ptr + 139, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);

        A.store.Bool(ptr + 8 + 130, false);
        A.store.Ref(ptr + 8 + 0, undefined);

        A.store.Bool(ptr + 8 + 4 + 45, false);
        A.store.Ref(ptr + 8 + 4 + 0, undefined);

        A.store.Bool(ptr + 8 + 4 + 4 + 16, false);
        A.store.Ref(ptr + 8 + 4 + 4 + 0, undefined);
        A.store.Ref(ptr + 8 + 4 + 4 + 4, undefined);
        A.store.Ref(ptr + 8 + 4 + 4 + 8, undefined);
        A.store.Ref(ptr + 8 + 4 + 4 + 12, undefined);
        A.store.Ref(ptr + 8 + 4 + 24, undefined);

        A.store.Bool(ptr + 8 + 4 + 28 + 16, false);
        A.store.Ref(ptr + 8 + 4 + 28 + 0, undefined);
        A.store.Ref(ptr + 8 + 4 + 28 + 4, undefined);
        A.store.Ref(ptr + 8 + 4 + 28 + 8, undefined);
        A.store.Ref(ptr + 8 + 4 + 28 + 12, undefined);
        A.store.Ref(ptr + 8 + 52, undefined);
        A.store.Ref(ptr + 8 + 56, undefined);
        A.store.Ref(ptr + 8 + 60, undefined);
        A.store.Enum(ptr + 8 + 64, -1);
        A.store.Ref(ptr + 8 + 68, undefined);
        A.store.Ref(ptr + 8 + 72, undefined);
        A.store.Ref(ptr + 8 + 76, undefined);
        A.store.Ref(ptr + 8 + 80, undefined);
        A.store.Bool(ptr + 8 + 127, false);
        A.store.Bool(ptr + 8 + 84, false);
        A.store.Ref(ptr + 8 + 88, undefined);
        A.store.Ref(ptr + 8 + 92, undefined);

        A.store.Bool(ptr + 8 + 96 + 28, false);
        A.store.Ref(ptr + 8 + 96 + 0, undefined);
        A.store.Ref(ptr + 8 + 96 + 4, undefined);
        A.store.Ref(ptr + 8 + 96 + 8, undefined);
        A.store.Ref(ptr + 8 + 96 + 12, undefined);
        A.store.Ref(ptr + 8 + 96 + 16, undefined);
        A.store.Ref(ptr + 8 + 96 + 20, undefined);
        A.store.Bool(ptr + 8 + 96 + 26, false);
        A.store.Bool(ptr + 8 + 96 + 24, false);
        A.store.Bool(ptr + 8 + 96 + 27, false);
        A.store.Bool(ptr + 8 + 96 + 25, false);
        A.store.Bool(ptr + 8 + 128, false);
        A.store.Bool(ptr + 8 + 125, false);
        A.store.Bool(ptr + 8 + 129, false);
        A.store.Bool(ptr + 8 + 126, false);
      } else {
        A.store.Bool(ptr + 140, true);
        A.store.Bool(ptr + 139, "AutoConnect" in x ? true : false);
        A.store.Bool(ptr + 0, x["AutoConnect"] ? true : false);
        A.store.Ref(ptr + 4, x["Authentication"]);

        if (typeof x["EAP"] === "undefined") {
          A.store.Bool(ptr + 8 + 130, false);
          A.store.Ref(ptr + 8 + 0, undefined);

          A.store.Bool(ptr + 8 + 4 + 45, false);
          A.store.Ref(ptr + 8 + 4 + 0, undefined);

          A.store.Bool(ptr + 8 + 4 + 4 + 16, false);
          A.store.Ref(ptr + 8 + 4 + 4 + 0, undefined);
          A.store.Ref(ptr + 8 + 4 + 4 + 4, undefined);
          A.store.Ref(ptr + 8 + 4 + 4 + 8, undefined);
          A.store.Ref(ptr + 8 + 4 + 4 + 12, undefined);
          A.store.Ref(ptr + 8 + 4 + 24, undefined);

          A.store.Bool(ptr + 8 + 4 + 28 + 16, false);
          A.store.Ref(ptr + 8 + 4 + 28 + 0, undefined);
          A.store.Ref(ptr + 8 + 4 + 28 + 4, undefined);
          A.store.Ref(ptr + 8 + 4 + 28 + 8, undefined);
          A.store.Ref(ptr + 8 + 4 + 28 + 12, undefined);
          A.store.Ref(ptr + 8 + 52, undefined);
          A.store.Ref(ptr + 8 + 56, undefined);
          A.store.Ref(ptr + 8 + 60, undefined);
          A.store.Enum(ptr + 8 + 64, -1);
          A.store.Ref(ptr + 8 + 68, undefined);
          A.store.Ref(ptr + 8 + 72, undefined);
          A.store.Ref(ptr + 8 + 76, undefined);
          A.store.Ref(ptr + 8 + 80, undefined);
          A.store.Bool(ptr + 8 + 127, false);
          A.store.Bool(ptr + 8 + 84, false);
          A.store.Ref(ptr + 8 + 88, undefined);
          A.store.Ref(ptr + 8 + 92, undefined);

          A.store.Bool(ptr + 8 + 96 + 28, false);
          A.store.Ref(ptr + 8 + 96 + 0, undefined);
          A.store.Ref(ptr + 8 + 96 + 4, undefined);
          A.store.Ref(ptr + 8 + 96 + 8, undefined);
          A.store.Ref(ptr + 8 + 96 + 12, undefined);
          A.store.Ref(ptr + 8 + 96 + 16, undefined);
          A.store.Ref(ptr + 8 + 96 + 20, undefined);
          A.store.Bool(ptr + 8 + 96 + 26, false);
          A.store.Bool(ptr + 8 + 96 + 24, false);
          A.store.Bool(ptr + 8 + 96 + 27, false);
          A.store.Bool(ptr + 8 + 96 + 25, false);
          A.store.Bool(ptr + 8 + 128, false);
          A.store.Bool(ptr + 8 + 125, false);
          A.store.Bool(ptr + 8 + 129, false);
          A.store.Bool(ptr + 8 + 126, false);
        } else {
          A.store.Bool(ptr + 8 + 130, true);
          A.store.Ref(ptr + 8 + 0, x["EAP"]["AnonymousIdentity"]);

          if (typeof x["EAP"]["ClientCertPattern"] === "undefined") {
            A.store.Bool(ptr + 8 + 4 + 45, false);
            A.store.Ref(ptr + 8 + 4 + 0, undefined);

            A.store.Bool(ptr + 8 + 4 + 4 + 16, false);
            A.store.Ref(ptr + 8 + 4 + 4 + 0, undefined);
            A.store.Ref(ptr + 8 + 4 + 4 + 4, undefined);
            A.store.Ref(ptr + 8 + 4 + 4 + 8, undefined);
            A.store.Ref(ptr + 8 + 4 + 4 + 12, undefined);
            A.store.Ref(ptr + 8 + 4 + 24, undefined);

            A.store.Bool(ptr + 8 + 4 + 28 + 16, false);
            A.store.Ref(ptr + 8 + 4 + 28 + 0, undefined);
            A.store.Ref(ptr + 8 + 4 + 28 + 4, undefined);
            A.store.Ref(ptr + 8 + 4 + 28 + 8, undefined);
            A.store.Ref(ptr + 8 + 4 + 28 + 12, undefined);
          } else {
            A.store.Bool(ptr + 8 + 4 + 45, true);
            A.store.Ref(ptr + 8 + 4 + 0, x["EAP"]["ClientCertPattern"]["EnrollmentURI"]);

            if (typeof x["EAP"]["ClientCertPattern"]["Issuer"] === "undefined") {
              A.store.Bool(ptr + 8 + 4 + 4 + 16, false);
              A.store.Ref(ptr + 8 + 4 + 4 + 0, undefined);
              A.store.Ref(ptr + 8 + 4 + 4 + 4, undefined);
              A.store.Ref(ptr + 8 + 4 + 4 + 8, undefined);
              A.store.Ref(ptr + 8 + 4 + 4 + 12, undefined);
            } else {
              A.store.Bool(ptr + 8 + 4 + 4 + 16, true);
              A.store.Ref(ptr + 8 + 4 + 4 + 0, x["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"]);
              A.store.Ref(ptr + 8 + 4 + 4 + 4, x["EAP"]["ClientCertPattern"]["Issuer"]["Locality"]);
              A.store.Ref(ptr + 8 + 4 + 4 + 8, x["EAP"]["ClientCertPattern"]["Issuer"]["Organization"]);
              A.store.Ref(ptr + 8 + 4 + 4 + 12, x["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"]);
            }
            A.store.Ref(ptr + 8 + 4 + 24, x["EAP"]["ClientCertPattern"]["IssuerCARef"]);

            if (typeof x["EAP"]["ClientCertPattern"]["Subject"] === "undefined") {
              A.store.Bool(ptr + 8 + 4 + 28 + 16, false);
              A.store.Ref(ptr + 8 + 4 + 28 + 0, undefined);
              A.store.Ref(ptr + 8 + 4 + 28 + 4, undefined);
              A.store.Ref(ptr + 8 + 4 + 28 + 8, undefined);
              A.store.Ref(ptr + 8 + 4 + 28 + 12, undefined);
            } else {
              A.store.Bool(ptr + 8 + 4 + 28 + 16, true);
              A.store.Ref(ptr + 8 + 4 + 28 + 0, x["EAP"]["ClientCertPattern"]["Subject"]["CommonName"]);
              A.store.Ref(ptr + 8 + 4 + 28 + 4, x["EAP"]["ClientCertPattern"]["Subject"]["Locality"]);
              A.store.Ref(ptr + 8 + 4 + 28 + 8, x["EAP"]["ClientCertPattern"]["Subject"]["Organization"]);
              A.store.Ref(ptr + 8 + 4 + 28 + 12, x["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"]);
            }
          }
          A.store.Ref(ptr + 8 + 52, x["EAP"]["ClientCertPKCS11Id"]);
          A.store.Ref(ptr + 8 + 56, x["EAP"]["ClientCertProvisioningProfileId"]);
          A.store.Ref(ptr + 8 + 60, x["EAP"]["ClientCertRef"]);
          A.store.Enum(ptr + 8 + 64, ["Ref", "Pattern"].indexOf(x["EAP"]["ClientCertType"] as string));
          A.store.Ref(ptr + 8 + 68, x["EAP"]["Identity"]);
          A.store.Ref(ptr + 8 + 72, x["EAP"]["Inner"]);
          A.store.Ref(ptr + 8 + 76, x["EAP"]["Outer"]);
          A.store.Ref(ptr + 8 + 80, x["EAP"]["Password"]);
          A.store.Bool(ptr + 8 + 127, "SaveCredentials" in x["EAP"] ? true : false);
          A.store.Bool(ptr + 8 + 84, x["EAP"]["SaveCredentials"] ? true : false);
          A.store.Ref(ptr + 8 + 88, x["EAP"]["ServerCAPEMs"]);
          A.store.Ref(ptr + 8 + 92, x["EAP"]["ServerCARefs"]);

          if (typeof x["EAP"]["SubjectMatch"] === "undefined") {
            A.store.Bool(ptr + 8 + 96 + 28, false);
            A.store.Ref(ptr + 8 + 96 + 0, undefined);
            A.store.Ref(ptr + 8 + 96 + 4, undefined);
            A.store.Ref(ptr + 8 + 96 + 8, undefined);
            A.store.Ref(ptr + 8 + 96 + 12, undefined);
            A.store.Ref(ptr + 8 + 96 + 16, undefined);
            A.store.Ref(ptr + 8 + 96 + 20, undefined);
            A.store.Bool(ptr + 8 + 96 + 26, false);
            A.store.Bool(ptr + 8 + 96 + 24, false);
            A.store.Bool(ptr + 8 + 96 + 27, false);
            A.store.Bool(ptr + 8 + 96 + 25, false);
          } else {
            A.store.Bool(ptr + 8 + 96 + 28, true);
            A.store.Ref(ptr + 8 + 96 + 0, x["EAP"]["SubjectMatch"]["Active"]);
            A.store.Ref(ptr + 8 + 96 + 4, x["EAP"]["SubjectMatch"]["Effective"]);
            A.store.Ref(ptr + 8 + 96 + 8, x["EAP"]["SubjectMatch"]["UserPolicy"]);
            A.store.Ref(ptr + 8 + 96 + 12, x["EAP"]["SubjectMatch"]["DevicePolicy"]);
            A.store.Ref(ptr + 8 + 96 + 16, x["EAP"]["SubjectMatch"]["UserSetting"]);
            A.store.Ref(ptr + 8 + 96 + 20, x["EAP"]["SubjectMatch"]["SharedSetting"]);
            A.store.Bool(ptr + 8 + 96 + 26, "UserEditable" in x["EAP"]["SubjectMatch"] ? true : false);
            A.store.Bool(ptr + 8 + 96 + 24, x["EAP"]["SubjectMatch"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 8 + 96 + 27, "DeviceEditable" in x["EAP"]["SubjectMatch"] ? true : false);
            A.store.Bool(ptr + 8 + 96 + 25, x["EAP"]["SubjectMatch"]["DeviceEditable"] ? true : false);
          }
          A.store.Bool(ptr + 8 + 128, "UseProactiveKeyCaching" in x["EAP"] ? true : false);
          A.store.Bool(ptr + 8 + 125, x["EAP"]["UseProactiveKeyCaching"] ? true : false);
          A.store.Bool(ptr + 8 + 129, "UseSystemCAs" in x["EAP"] ? true : false);
          A.store.Bool(ptr + 8 + 126, x["EAP"]["UseSystemCAs"] ? true : false);
        }
      }
    },
    "load_EthernetProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 139)) {
        x["AutoConnect"] = A.load.Bool(ptr + 0);
      } else {
        delete x["AutoConnect"];
      }
      x["Authentication"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 8 + 130)) {
        x["EAP"] = {};
        x["EAP"]["AnonymousIdentity"] = A.load.Ref(ptr + 8 + 0, undefined);
        if (A.load.Bool(ptr + 8 + 4 + 45)) {
          x["EAP"]["ClientCertPattern"] = {};
          x["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(ptr + 8 + 4 + 0, undefined);
          if (A.load.Bool(ptr + 8 + 4 + 4 + 16)) {
            x["EAP"]["ClientCertPattern"]["Issuer"] = {};
            x["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(ptr + 8 + 4 + 4 + 0, undefined);
            x["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(ptr + 8 + 4 + 4 + 4, undefined);
            x["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(ptr + 8 + 4 + 4 + 8, undefined);
            x["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(ptr + 8 + 4 + 4 + 12, undefined);
          } else {
            delete x["EAP"]["ClientCertPattern"]["Issuer"];
          }
          x["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(ptr + 8 + 4 + 24, undefined);
          if (A.load.Bool(ptr + 8 + 4 + 28 + 16)) {
            x["EAP"]["ClientCertPattern"]["Subject"] = {};
            x["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(ptr + 8 + 4 + 28 + 0, undefined);
            x["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(ptr + 8 + 4 + 28 + 4, undefined);
            x["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(ptr + 8 + 4 + 28 + 8, undefined);
            x["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
              ptr + 8 + 4 + 28 + 12,
              undefined
            );
          } else {
            delete x["EAP"]["ClientCertPattern"]["Subject"];
          }
        } else {
          delete x["EAP"]["ClientCertPattern"];
        }
        x["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(ptr + 8 + 52, undefined);
        x["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(ptr + 8 + 56, undefined);
        x["EAP"]["ClientCertRef"] = A.load.Ref(ptr + 8 + 60, undefined);
        x["EAP"]["ClientCertType"] = A.load.Enum(ptr + 8 + 64, ["Ref", "Pattern"]);
        x["EAP"]["Identity"] = A.load.Ref(ptr + 8 + 68, undefined);
        x["EAP"]["Inner"] = A.load.Ref(ptr + 8 + 72, undefined);
        x["EAP"]["Outer"] = A.load.Ref(ptr + 8 + 76, undefined);
        x["EAP"]["Password"] = A.load.Ref(ptr + 8 + 80, undefined);
        if (A.load.Bool(ptr + 8 + 127)) {
          x["EAP"]["SaveCredentials"] = A.load.Bool(ptr + 8 + 84);
        } else {
          delete x["EAP"]["SaveCredentials"];
        }
        x["EAP"]["ServerCAPEMs"] = A.load.Ref(ptr + 8 + 88, undefined);
        x["EAP"]["ServerCARefs"] = A.load.Ref(ptr + 8 + 92, undefined);
        if (A.load.Bool(ptr + 8 + 96 + 28)) {
          x["EAP"]["SubjectMatch"] = {};
          x["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(ptr + 8 + 96 + 0, undefined);
          x["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(ptr + 8 + 96 + 4, undefined);
          x["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(ptr + 8 + 96 + 8, undefined);
          x["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(ptr + 8 + 96 + 12, undefined);
          x["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(ptr + 8 + 96 + 16, undefined);
          x["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(ptr + 8 + 96 + 20, undefined);
          if (A.load.Bool(ptr + 8 + 96 + 26)) {
            x["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(ptr + 8 + 96 + 24);
          } else {
            delete x["EAP"]["SubjectMatch"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 8 + 96 + 27)) {
            x["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(ptr + 8 + 96 + 25);
          } else {
            delete x["EAP"]["SubjectMatch"]["DeviceEditable"];
          }
        } else {
          delete x["EAP"]["SubjectMatch"];
        }
        if (A.load.Bool(ptr + 8 + 128)) {
          x["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(ptr + 8 + 125);
        } else {
          delete x["EAP"]["UseProactiveKeyCaching"];
        }
        if (A.load.Bool(ptr + 8 + 129)) {
          x["EAP"]["UseSystemCAs"] = A.load.Bool(ptr + 8 + 126);
        } else {
          delete x["EAP"]["UseSystemCAs"];
        }
      } else {
        delete x["EAP"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_EthernetStateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["Authentication"]);
      }
    },
    "load_EthernetStateProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["Authentication"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GlobalPolicy": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Bool(ptr + 8, "AllowOnlyPolicyNetworksToAutoconnect" in x ? true : false);
        A.store.Bool(ptr + 0, x["AllowOnlyPolicyNetworksToAutoconnect"] ? true : false);
        A.store.Bool(ptr + 9, "AllowOnlyPolicyNetworksToConnect" in x ? true : false);
        A.store.Bool(ptr + 1, x["AllowOnlyPolicyNetworksToConnect"] ? true : false);
        A.store.Bool(ptr + 10, "AllowOnlyPolicyNetworksToConnectIfAvailable" in x ? true : false);
        A.store.Bool(ptr + 2, x["AllowOnlyPolicyNetworksToConnectIfAvailable"] ? true : false);
        A.store.Ref(ptr + 4, x["BlockedHexSSIDs"]);
      }
    },
    "load_GlobalPolicy": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["AllowOnlyPolicyNetworksToAutoconnect"] = A.load.Bool(ptr + 0);
      } else {
        delete x["AllowOnlyPolicyNetworksToAutoconnect"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["AllowOnlyPolicyNetworksToConnect"] = A.load.Bool(ptr + 1);
      } else {
        delete x["AllowOnlyPolicyNetworksToConnect"];
      }
      if (A.load.Bool(ptr + 10)) {
        x["AllowOnlyPolicyNetworksToConnectIfAvailable"] = A.load.Bool(ptr + 2);
      } else {
        delete x["AllowOnlyPolicyNetworksToConnectIfAvailable"];
      }
      x["BlockedHexSSIDs"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManagedBoolean": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 13, false);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Bool(ptr + 14, "Active" in x ? true : false);
        A.store.Bool(ptr + 0, x["Active"] ? true : false);
        A.store.Ref(ptr + 4, x["Effective"]);
        A.store.Bool(ptr + 15, "UserPolicy" in x ? true : false);
        A.store.Bool(ptr + 8, x["UserPolicy"] ? true : false);
        A.store.Bool(ptr + 16, "DevicePolicy" in x ? true : false);
        A.store.Bool(ptr + 9, x["DevicePolicy"] ? true : false);
        A.store.Bool(ptr + 17, "UserSetting" in x ? true : false);
        A.store.Bool(ptr + 10, x["UserSetting"] ? true : false);
        A.store.Bool(ptr + 18, "SharedSetting" in x ? true : false);
        A.store.Bool(ptr + 11, x["SharedSetting"] ? true : false);
        A.store.Bool(ptr + 19, "UserEditable" in x ? true : false);
        A.store.Bool(ptr + 12, x["UserEditable"] ? true : false);
        A.store.Bool(ptr + 20, "DeviceEditable" in x ? true : false);
        A.store.Bool(ptr + 13, x["DeviceEditable"] ? true : false);
      }
    },
    "load_ManagedBoolean": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 14)) {
        x["Active"] = A.load.Bool(ptr + 0);
      } else {
        delete x["Active"];
      }
      x["Effective"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 15)) {
        x["UserPolicy"] = A.load.Bool(ptr + 8);
      } else {
        delete x["UserPolicy"];
      }
      if (A.load.Bool(ptr + 16)) {
        x["DevicePolicy"] = A.load.Bool(ptr + 9);
      } else {
        delete x["DevicePolicy"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["UserSetting"] = A.load.Bool(ptr + 10);
      } else {
        delete x["UserSetting"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["SharedSetting"] = A.load.Bool(ptr + 11);
      } else {
        delete x["SharedSetting"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["UserEditable"] = A.load.Bool(ptr + 12);
      } else {
        delete x["UserEditable"];
      }
      if (A.load.Bool(ptr + 20)) {
        x["DeviceEditable"] = A.load.Bool(ptr + 13);
      } else {
        delete x["DeviceEditable"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManagedCellularProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 134, false);

        A.store.Bool(ptr + 0 + 21, false);
        A.store.Bool(ptr + 0 + 14, false);
        A.store.Bool(ptr + 0 + 0, false);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Bool(ptr + 0 + 15, false);
        A.store.Bool(ptr + 0 + 8, false);
        A.store.Bool(ptr + 0 + 16, false);
        A.store.Bool(ptr + 0 + 9, false);
        A.store.Bool(ptr + 0 + 17, false);
        A.store.Bool(ptr + 0 + 10, false);
        A.store.Bool(ptr + 0 + 18, false);
        A.store.Bool(ptr + 0 + 11, false);
        A.store.Bool(ptr + 0 + 19, false);
        A.store.Bool(ptr + 0 + 12, false);
        A.store.Bool(ptr + 0 + 20, false);
        A.store.Bool(ptr + 0 + 13, false);
        A.store.Ref(ptr + 24, undefined);
        A.store.Enum(ptr + 28, -1);
        A.store.Bool(ptr + 129, false);
        A.store.Bool(ptr + 32, false);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Ref(ptr + 48, undefined);
        A.store.Ref(ptr + 52, undefined);
        A.store.Ref(ptr + 56, undefined);
        A.store.Ref(ptr + 60, undefined);
        A.store.Ref(ptr + 64, undefined);

        A.store.Bool(ptr + 68 + 12, false);
        A.store.Ref(ptr + 68 + 0, undefined);
        A.store.Ref(ptr + 68 + 4, undefined);
        A.store.Ref(ptr + 68 + 8, undefined);
        A.store.Ref(ptr + 84, undefined);
        A.store.Bool(ptr + 130, false);
        A.store.Bool(ptr + 88, false);

        A.store.Bool(ptr + 92 + 12, false);
        A.store.Ref(ptr + 92 + 0, undefined);
        A.store.Ref(ptr + 92 + 4, undefined);
        A.store.Ref(ptr + 92 + 8, undefined);

        A.store.Bool(ptr + 108 + 14, false);
        A.store.Ref(ptr + 108 + 0, undefined);
        A.store.Bool(ptr + 108 + 12, false);
        A.store.Bool(ptr + 108 + 4, false);
        A.store.Bool(ptr + 108 + 13, false);
        A.store.Int32(ptr + 108 + 8, 0);
        A.store.Bool(ptr + 131, false);
        A.store.Bool(ptr + 123, false);
        A.store.Bool(ptr + 132, false);
        A.store.Int32(ptr + 124, 0);
        A.store.Bool(ptr + 133, false);
        A.store.Bool(ptr + 128, false);
      } else {
        A.store.Bool(ptr + 134, true);

        if (typeof x["AutoConnect"] === "undefined") {
          A.store.Bool(ptr + 0 + 21, false);
          A.store.Bool(ptr + 0 + 14, false);
          A.store.Bool(ptr + 0 + 0, false);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Bool(ptr + 0 + 15, false);
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Bool(ptr + 0 + 16, false);
          A.store.Bool(ptr + 0 + 9, false);
          A.store.Bool(ptr + 0 + 17, false);
          A.store.Bool(ptr + 0 + 10, false);
          A.store.Bool(ptr + 0 + 18, false);
          A.store.Bool(ptr + 0 + 11, false);
          A.store.Bool(ptr + 0 + 19, false);
          A.store.Bool(ptr + 0 + 12, false);
          A.store.Bool(ptr + 0 + 20, false);
          A.store.Bool(ptr + 0 + 13, false);
        } else {
          A.store.Bool(ptr + 0 + 21, true);
          A.store.Bool(ptr + 0 + 14, "Active" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 0, x["AutoConnect"]["Active"] ? true : false);
          A.store.Ref(ptr + 0 + 4, x["AutoConnect"]["Effective"]);
          A.store.Bool(ptr + 0 + 15, "UserPolicy" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 8, x["AutoConnect"]["UserPolicy"] ? true : false);
          A.store.Bool(ptr + 0 + 16, "DevicePolicy" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 9, x["AutoConnect"]["DevicePolicy"] ? true : false);
          A.store.Bool(ptr + 0 + 17, "UserSetting" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 10, x["AutoConnect"]["UserSetting"] ? true : false);
          A.store.Bool(ptr + 0 + 18, "SharedSetting" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 11, x["AutoConnect"]["SharedSetting"] ? true : false);
          A.store.Bool(ptr + 0 + 19, "UserEditable" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 12, x["AutoConnect"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 0 + 20, "DeviceEditable" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 13, x["AutoConnect"]["DeviceEditable"] ? true : false);
        }
        A.store.Ref(ptr + 24, x["ActivationType"]);
        A.store.Enum(
          ptr + 28,
          ["Activated", "Activating", "NotActivated", "PartiallyActivated"].indexOf(x["ActivationState"] as string)
        );
        A.store.Bool(ptr + 129, "AllowRoaming" in x ? true : false);
        A.store.Bool(ptr + 32, x["AllowRoaming"] ? true : false);
        A.store.Ref(ptr + 36, x["Family"]);
        A.store.Ref(ptr + 40, x["FirmwareRevision"]);
        A.store.Ref(ptr + 44, x["FoundNetworks"]);
        A.store.Ref(ptr + 48, x["HardwareRevision"]);
        A.store.Ref(ptr + 52, x["HomeProvider"]);
        A.store.Ref(ptr + 56, x["Manufacturer"]);
        A.store.Ref(ptr + 60, x["ModelID"]);
        A.store.Ref(ptr + 64, x["NetworkTechnology"]);

        if (typeof x["PaymentPortal"] === "undefined") {
          A.store.Bool(ptr + 68 + 12, false);
          A.store.Ref(ptr + 68 + 0, undefined);
          A.store.Ref(ptr + 68 + 4, undefined);
          A.store.Ref(ptr + 68 + 8, undefined);
        } else {
          A.store.Bool(ptr + 68 + 12, true);
          A.store.Ref(ptr + 68 + 0, x["PaymentPortal"]["Method"]);
          A.store.Ref(ptr + 68 + 4, x["PaymentPortal"]["PostData"]);
          A.store.Ref(ptr + 68 + 8, x["PaymentPortal"]["Url"]);
        }
        A.store.Ref(ptr + 84, x["RoamingState"]);
        A.store.Bool(ptr + 130, "Scanning" in x ? true : false);
        A.store.Bool(ptr + 88, x["Scanning"] ? true : false);

        if (typeof x["ServingOperator"] === "undefined") {
          A.store.Bool(ptr + 92 + 12, false);
          A.store.Ref(ptr + 92 + 0, undefined);
          A.store.Ref(ptr + 92 + 4, undefined);
          A.store.Ref(ptr + 92 + 8, undefined);
        } else {
          A.store.Bool(ptr + 92 + 12, true);
          A.store.Ref(ptr + 92 + 0, x["ServingOperator"]["Name"]);
          A.store.Ref(ptr + 92 + 4, x["ServingOperator"]["Code"]);
          A.store.Ref(ptr + 92 + 8, x["ServingOperator"]["Country"]);
        }

        if (typeof x["SIMLockStatus"] === "undefined") {
          A.store.Bool(ptr + 108 + 14, false);
          A.store.Ref(ptr + 108 + 0, undefined);
          A.store.Bool(ptr + 108 + 12, false);
          A.store.Bool(ptr + 108 + 4, false);
          A.store.Bool(ptr + 108 + 13, false);
          A.store.Int32(ptr + 108 + 8, 0);
        } else {
          A.store.Bool(ptr + 108 + 14, true);
          A.store.Ref(ptr + 108 + 0, x["SIMLockStatus"]["LockType"]);
          A.store.Bool(ptr + 108 + 12, "LockEnabled" in x["SIMLockStatus"] ? true : false);
          A.store.Bool(ptr + 108 + 4, x["SIMLockStatus"]["LockEnabled"] ? true : false);
          A.store.Bool(ptr + 108 + 13, "RetriesLeft" in x["SIMLockStatus"] ? true : false);
          A.store.Int32(
            ptr + 108 + 8,
            x["SIMLockStatus"]["RetriesLeft"] === undefined ? 0 : (x["SIMLockStatus"]["RetriesLeft"] as number)
          );
        }
        A.store.Bool(ptr + 131, "SIMPresent" in x ? true : false);
        A.store.Bool(ptr + 123, x["SIMPresent"] ? true : false);
        A.store.Bool(ptr + 132, "SignalStrength" in x ? true : false);
        A.store.Int32(ptr + 124, x["SignalStrength"] === undefined ? 0 : (x["SignalStrength"] as number));
        A.store.Bool(ptr + 133, "SupportNetworkScan" in x ? true : false);
        A.store.Bool(ptr + 128, x["SupportNetworkScan"] ? true : false);
      }
    },
    "load_ManagedCellularProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 21)) {
        x["AutoConnect"] = {};
        if (A.load.Bool(ptr + 0 + 14)) {
          x["AutoConnect"]["Active"] = A.load.Bool(ptr + 0 + 0);
        } else {
          delete x["AutoConnect"]["Active"];
        }
        x["AutoConnect"]["Effective"] = A.load.Ref(ptr + 0 + 4, undefined);
        if (A.load.Bool(ptr + 0 + 15)) {
          x["AutoConnect"]["UserPolicy"] = A.load.Bool(ptr + 0 + 8);
        } else {
          delete x["AutoConnect"]["UserPolicy"];
        }
        if (A.load.Bool(ptr + 0 + 16)) {
          x["AutoConnect"]["DevicePolicy"] = A.load.Bool(ptr + 0 + 9);
        } else {
          delete x["AutoConnect"]["DevicePolicy"];
        }
        if (A.load.Bool(ptr + 0 + 17)) {
          x["AutoConnect"]["UserSetting"] = A.load.Bool(ptr + 0 + 10);
        } else {
          delete x["AutoConnect"]["UserSetting"];
        }
        if (A.load.Bool(ptr + 0 + 18)) {
          x["AutoConnect"]["SharedSetting"] = A.load.Bool(ptr + 0 + 11);
        } else {
          delete x["AutoConnect"]["SharedSetting"];
        }
        if (A.load.Bool(ptr + 0 + 19)) {
          x["AutoConnect"]["UserEditable"] = A.load.Bool(ptr + 0 + 12);
        } else {
          delete x["AutoConnect"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 0 + 20)) {
          x["AutoConnect"]["DeviceEditable"] = A.load.Bool(ptr + 0 + 13);
        } else {
          delete x["AutoConnect"]["DeviceEditable"];
        }
      } else {
        delete x["AutoConnect"];
      }
      x["ActivationType"] = A.load.Ref(ptr + 24, undefined);
      x["ActivationState"] = A.load.Enum(ptr + 28, ["Activated", "Activating", "NotActivated", "PartiallyActivated"]);
      if (A.load.Bool(ptr + 129)) {
        x["AllowRoaming"] = A.load.Bool(ptr + 32);
      } else {
        delete x["AllowRoaming"];
      }
      x["Family"] = A.load.Ref(ptr + 36, undefined);
      x["FirmwareRevision"] = A.load.Ref(ptr + 40, undefined);
      x["FoundNetworks"] = A.load.Ref(ptr + 44, undefined);
      x["HardwareRevision"] = A.load.Ref(ptr + 48, undefined);
      x["HomeProvider"] = A.load.Ref(ptr + 52, undefined);
      x["Manufacturer"] = A.load.Ref(ptr + 56, undefined);
      x["ModelID"] = A.load.Ref(ptr + 60, undefined);
      x["NetworkTechnology"] = A.load.Ref(ptr + 64, undefined);
      if (A.load.Bool(ptr + 68 + 12)) {
        x["PaymentPortal"] = {};
        x["PaymentPortal"]["Method"] = A.load.Ref(ptr + 68 + 0, undefined);
        x["PaymentPortal"]["PostData"] = A.load.Ref(ptr + 68 + 4, undefined);
        x["PaymentPortal"]["Url"] = A.load.Ref(ptr + 68 + 8, undefined);
      } else {
        delete x["PaymentPortal"];
      }
      x["RoamingState"] = A.load.Ref(ptr + 84, undefined);
      if (A.load.Bool(ptr + 130)) {
        x["Scanning"] = A.load.Bool(ptr + 88);
      } else {
        delete x["Scanning"];
      }
      if (A.load.Bool(ptr + 92 + 12)) {
        x["ServingOperator"] = {};
        x["ServingOperator"]["Name"] = A.load.Ref(ptr + 92 + 0, undefined);
        x["ServingOperator"]["Code"] = A.load.Ref(ptr + 92 + 4, undefined);
        x["ServingOperator"]["Country"] = A.load.Ref(ptr + 92 + 8, undefined);
      } else {
        delete x["ServingOperator"];
      }
      if (A.load.Bool(ptr + 108 + 14)) {
        x["SIMLockStatus"] = {};
        x["SIMLockStatus"]["LockType"] = A.load.Ref(ptr + 108 + 0, undefined);
        if (A.load.Bool(ptr + 108 + 12)) {
          x["SIMLockStatus"]["LockEnabled"] = A.load.Bool(ptr + 108 + 4);
        } else {
          delete x["SIMLockStatus"]["LockEnabled"];
        }
        if (A.load.Bool(ptr + 108 + 13)) {
          x["SIMLockStatus"]["RetriesLeft"] = A.load.Int32(ptr + 108 + 8);
        } else {
          delete x["SIMLockStatus"]["RetriesLeft"];
        }
      } else {
        delete x["SIMLockStatus"];
      }
      if (A.load.Bool(ptr + 131)) {
        x["SIMPresent"] = A.load.Bool(ptr + 123);
      } else {
        delete x["SIMPresent"];
      }
      if (A.load.Bool(ptr + 132)) {
        x["SignalStrength"] = A.load.Int32(ptr + 124);
      } else {
        delete x["SignalStrength"];
      }
      if (A.load.Bool(ptr + 133)) {
        x["SupportNetworkScan"] = A.load.Bool(ptr + 128);
      } else {
        delete x["SupportNetworkScan"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManagedEthernetProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 53, false);

        A.store.Bool(ptr + 0 + 21, false);
        A.store.Bool(ptr + 0 + 14, false);
        A.store.Bool(ptr + 0 + 0, false);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Bool(ptr + 0 + 15, false);
        A.store.Bool(ptr + 0 + 8, false);
        A.store.Bool(ptr + 0 + 16, false);
        A.store.Bool(ptr + 0 + 9, false);
        A.store.Bool(ptr + 0 + 17, false);
        A.store.Bool(ptr + 0 + 10, false);
        A.store.Bool(ptr + 0 + 18, false);
        A.store.Bool(ptr + 0 + 11, false);
        A.store.Bool(ptr + 0 + 19, false);
        A.store.Bool(ptr + 0 + 12, false);
        A.store.Bool(ptr + 0 + 20, false);
        A.store.Bool(ptr + 0 + 13, false);

        A.store.Bool(ptr + 24 + 28, false);
        A.store.Ref(ptr + 24 + 0, undefined);
        A.store.Ref(ptr + 24 + 4, undefined);
        A.store.Ref(ptr + 24 + 8, undefined);
        A.store.Ref(ptr + 24 + 12, undefined);
        A.store.Ref(ptr + 24 + 16, undefined);
        A.store.Ref(ptr + 24 + 20, undefined);
        A.store.Bool(ptr + 24 + 26, false);
        A.store.Bool(ptr + 24 + 24, false);
        A.store.Bool(ptr + 24 + 27, false);
        A.store.Bool(ptr + 24 + 25, false);
      } else {
        A.store.Bool(ptr + 53, true);

        if (typeof x["AutoConnect"] === "undefined") {
          A.store.Bool(ptr + 0 + 21, false);
          A.store.Bool(ptr + 0 + 14, false);
          A.store.Bool(ptr + 0 + 0, false);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Bool(ptr + 0 + 15, false);
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Bool(ptr + 0 + 16, false);
          A.store.Bool(ptr + 0 + 9, false);
          A.store.Bool(ptr + 0 + 17, false);
          A.store.Bool(ptr + 0 + 10, false);
          A.store.Bool(ptr + 0 + 18, false);
          A.store.Bool(ptr + 0 + 11, false);
          A.store.Bool(ptr + 0 + 19, false);
          A.store.Bool(ptr + 0 + 12, false);
          A.store.Bool(ptr + 0 + 20, false);
          A.store.Bool(ptr + 0 + 13, false);
        } else {
          A.store.Bool(ptr + 0 + 21, true);
          A.store.Bool(ptr + 0 + 14, "Active" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 0, x["AutoConnect"]["Active"] ? true : false);
          A.store.Ref(ptr + 0 + 4, x["AutoConnect"]["Effective"]);
          A.store.Bool(ptr + 0 + 15, "UserPolicy" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 8, x["AutoConnect"]["UserPolicy"] ? true : false);
          A.store.Bool(ptr + 0 + 16, "DevicePolicy" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 9, x["AutoConnect"]["DevicePolicy"] ? true : false);
          A.store.Bool(ptr + 0 + 17, "UserSetting" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 10, x["AutoConnect"]["UserSetting"] ? true : false);
          A.store.Bool(ptr + 0 + 18, "SharedSetting" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 11, x["AutoConnect"]["SharedSetting"] ? true : false);
          A.store.Bool(ptr + 0 + 19, "UserEditable" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 12, x["AutoConnect"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 0 + 20, "DeviceEditable" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 13, x["AutoConnect"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["Authentication"] === "undefined") {
          A.store.Bool(ptr + 24 + 28, false);
          A.store.Ref(ptr + 24 + 0, undefined);
          A.store.Ref(ptr + 24 + 4, undefined);
          A.store.Ref(ptr + 24 + 8, undefined);
          A.store.Ref(ptr + 24 + 12, undefined);
          A.store.Ref(ptr + 24 + 16, undefined);
          A.store.Ref(ptr + 24 + 20, undefined);
          A.store.Bool(ptr + 24 + 26, false);
          A.store.Bool(ptr + 24 + 24, false);
          A.store.Bool(ptr + 24 + 27, false);
          A.store.Bool(ptr + 24 + 25, false);
        } else {
          A.store.Bool(ptr + 24 + 28, true);
          A.store.Ref(ptr + 24 + 0, x["Authentication"]["Active"]);
          A.store.Ref(ptr + 24 + 4, x["Authentication"]["Effective"]);
          A.store.Ref(ptr + 24 + 8, x["Authentication"]["UserPolicy"]);
          A.store.Ref(ptr + 24 + 12, x["Authentication"]["DevicePolicy"]);
          A.store.Ref(ptr + 24 + 16, x["Authentication"]["UserSetting"]);
          A.store.Ref(ptr + 24 + 20, x["Authentication"]["SharedSetting"]);
          A.store.Bool(ptr + 24 + 26, "UserEditable" in x["Authentication"] ? true : false);
          A.store.Bool(ptr + 24 + 24, x["Authentication"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 24 + 27, "DeviceEditable" in x["Authentication"] ? true : false);
          A.store.Bool(ptr + 24 + 25, x["Authentication"]["DeviceEditable"] ? true : false);
        }
      }
    },
    "load_ManagedEthernetProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 21)) {
        x["AutoConnect"] = {};
        if (A.load.Bool(ptr + 0 + 14)) {
          x["AutoConnect"]["Active"] = A.load.Bool(ptr + 0 + 0);
        } else {
          delete x["AutoConnect"]["Active"];
        }
        x["AutoConnect"]["Effective"] = A.load.Ref(ptr + 0 + 4, undefined);
        if (A.load.Bool(ptr + 0 + 15)) {
          x["AutoConnect"]["UserPolicy"] = A.load.Bool(ptr + 0 + 8);
        } else {
          delete x["AutoConnect"]["UserPolicy"];
        }
        if (A.load.Bool(ptr + 0 + 16)) {
          x["AutoConnect"]["DevicePolicy"] = A.load.Bool(ptr + 0 + 9);
        } else {
          delete x["AutoConnect"]["DevicePolicy"];
        }
        if (A.load.Bool(ptr + 0 + 17)) {
          x["AutoConnect"]["UserSetting"] = A.load.Bool(ptr + 0 + 10);
        } else {
          delete x["AutoConnect"]["UserSetting"];
        }
        if (A.load.Bool(ptr + 0 + 18)) {
          x["AutoConnect"]["SharedSetting"] = A.load.Bool(ptr + 0 + 11);
        } else {
          delete x["AutoConnect"]["SharedSetting"];
        }
        if (A.load.Bool(ptr + 0 + 19)) {
          x["AutoConnect"]["UserEditable"] = A.load.Bool(ptr + 0 + 12);
        } else {
          delete x["AutoConnect"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 0 + 20)) {
          x["AutoConnect"]["DeviceEditable"] = A.load.Bool(ptr + 0 + 13);
        } else {
          delete x["AutoConnect"]["DeviceEditable"];
        }
      } else {
        delete x["AutoConnect"];
      }
      if (A.load.Bool(ptr + 24 + 28)) {
        x["Authentication"] = {};
        x["Authentication"]["Active"] = A.load.Ref(ptr + 24 + 0, undefined);
        x["Authentication"]["Effective"] = A.load.Ref(ptr + 24 + 4, undefined);
        x["Authentication"]["UserPolicy"] = A.load.Ref(ptr + 24 + 8, undefined);
        x["Authentication"]["DevicePolicy"] = A.load.Ref(ptr + 24 + 12, undefined);
        x["Authentication"]["UserSetting"] = A.load.Ref(ptr + 24 + 16, undefined);
        x["Authentication"]["SharedSetting"] = A.load.Ref(ptr + 24 + 20, undefined);
        if (A.load.Bool(ptr + 24 + 26)) {
          x["Authentication"]["UserEditable"] = A.load.Bool(ptr + 24 + 24);
        } else {
          delete x["Authentication"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 24 + 27)) {
          x["Authentication"]["DeviceEditable"] = A.load.Bool(ptr + 24 + 25);
        } else {
          delete x["Authentication"]["DeviceEditable"];
        }
      } else {
        delete x["Authentication"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_IPConfigType": (ref: heap.Ref<string>): number => {
      const idx = ["DHCP", "Static"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ManagedIPConfigType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Enum(ptr + 12, -1);
        A.store.Enum(ptr + 16, -1);
        A.store.Enum(ptr + 20, -1);
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 27, false);
        A.store.Bool(ptr + 25, false);
      } else {
        A.store.Bool(ptr + 28, true);
        A.store.Enum(ptr + 0, ["DHCP", "Static"].indexOf(x["Active"] as string));
        A.store.Ref(ptr + 4, x["Effective"]);
        A.store.Enum(ptr + 8, ["DHCP", "Static"].indexOf(x["UserPolicy"] as string));
        A.store.Enum(ptr + 12, ["DHCP", "Static"].indexOf(x["DevicePolicy"] as string));
        A.store.Enum(ptr + 16, ["DHCP", "Static"].indexOf(x["UserSetting"] as string));
        A.store.Enum(ptr + 20, ["DHCP", "Static"].indexOf(x["SharedSetting"] as string));
        A.store.Bool(ptr + 26, "UserEditable" in x ? true : false);
        A.store.Bool(ptr + 24, x["UserEditable"] ? true : false);
        A.store.Bool(ptr + 27, "DeviceEditable" in x ? true : false);
        A.store.Bool(ptr + 25, x["DeviceEditable"] ? true : false);
      }
    },
    "load_ManagedIPConfigType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["Active"] = A.load.Enum(ptr + 0, ["DHCP", "Static"]);
      x["Effective"] = A.load.Ref(ptr + 4, undefined);
      x["UserPolicy"] = A.load.Enum(ptr + 8, ["DHCP", "Static"]);
      x["DevicePolicy"] = A.load.Enum(ptr + 12, ["DHCP", "Static"]);
      x["UserSetting"] = A.load.Enum(ptr + 16, ["DHCP", "Static"]);
      x["SharedSetting"] = A.load.Enum(ptr + 20, ["DHCP", "Static"]);
      if (A.load.Bool(ptr + 26)) {
        x["UserEditable"] = A.load.Bool(ptr + 24);
      } else {
        delete x["UserEditable"];
      }
      if (A.load.Bool(ptr + 27)) {
        x["DeviceEditable"] = A.load.Bool(ptr + 25);
      } else {
        delete x["DeviceEditable"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_IPConfigProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 37, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Bool(ptr + 36, false);
        A.store.Int32(ptr + 24, 0);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
      } else {
        A.store.Bool(ptr + 37, true);
        A.store.Ref(ptr + 0, x["Gateway"]);
        A.store.Ref(ptr + 4, x["IPAddress"]);
        A.store.Ref(ptr + 8, x["ExcludedRoutes"]);
        A.store.Ref(ptr + 12, x["IncludedRoutes"]);
        A.store.Ref(ptr + 16, x["NameServers"]);
        A.store.Ref(ptr + 20, x["SearchDomains"]);
        A.store.Bool(ptr + 36, "RoutingPrefix" in x ? true : false);
        A.store.Int32(ptr + 24, x["RoutingPrefix"] === undefined ? 0 : (x["RoutingPrefix"] as number));
        A.store.Ref(ptr + 28, x["Type"]);
        A.store.Ref(ptr + 32, x["WebProxyAutoDiscoveryUrl"]);
      }
    },
    "load_IPConfigProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["Gateway"] = A.load.Ref(ptr + 0, undefined);
      x["IPAddress"] = A.load.Ref(ptr + 4, undefined);
      x["ExcludedRoutes"] = A.load.Ref(ptr + 8, undefined);
      x["IncludedRoutes"] = A.load.Ref(ptr + 12, undefined);
      x["NameServers"] = A.load.Ref(ptr + 16, undefined);
      x["SearchDomains"] = A.load.Ref(ptr + 20, undefined);
      if (A.load.Bool(ptr + 36)) {
        x["RoutingPrefix"] = A.load.Int32(ptr + 24);
      } else {
        delete x["RoutingPrefix"];
      }
      x["Type"] = A.load.Ref(ptr + 28, undefined);
      x["WebProxyAutoDiscoveryUrl"] = A.load.Ref(ptr + 32, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManagedLong": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 33, false);
        A.store.Bool(ptr + 26, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 27, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 28, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 29, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Bool(ptr + 30, false);
        A.store.Int32(ptr + 20, 0);
        A.store.Bool(ptr + 31, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 25, false);
      } else {
        A.store.Bool(ptr + 33, true);
        A.store.Bool(ptr + 26, "Active" in x ? true : false);
        A.store.Int32(ptr + 0, x["Active"] === undefined ? 0 : (x["Active"] as number));
        A.store.Ref(ptr + 4, x["Effective"]);
        A.store.Bool(ptr + 27, "UserPolicy" in x ? true : false);
        A.store.Int32(ptr + 8, x["UserPolicy"] === undefined ? 0 : (x["UserPolicy"] as number));
        A.store.Bool(ptr + 28, "DevicePolicy" in x ? true : false);
        A.store.Int32(ptr + 12, x["DevicePolicy"] === undefined ? 0 : (x["DevicePolicy"] as number));
        A.store.Bool(ptr + 29, "UserSetting" in x ? true : false);
        A.store.Int32(ptr + 16, x["UserSetting"] === undefined ? 0 : (x["UserSetting"] as number));
        A.store.Bool(ptr + 30, "SharedSetting" in x ? true : false);
        A.store.Int32(ptr + 20, x["SharedSetting"] === undefined ? 0 : (x["SharedSetting"] as number));
        A.store.Bool(ptr + 31, "UserEditable" in x ? true : false);
        A.store.Bool(ptr + 24, x["UserEditable"] ? true : false);
        A.store.Bool(ptr + 32, "DeviceEditable" in x ? true : false);
        A.store.Bool(ptr + 25, x["DeviceEditable"] ? true : false);
      }
    },
    "load_ManagedLong": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 26)) {
        x["Active"] = A.load.Int32(ptr + 0);
      } else {
        delete x["Active"];
      }
      x["Effective"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 27)) {
        x["UserPolicy"] = A.load.Int32(ptr + 8);
      } else {
        delete x["UserPolicy"];
      }
      if (A.load.Bool(ptr + 28)) {
        x["DevicePolicy"] = A.load.Int32(ptr + 12);
      } else {
        delete x["DevicePolicy"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["UserSetting"] = A.load.Int32(ptr + 16);
      } else {
        delete x["UserSetting"];
      }
      if (A.load.Bool(ptr + 30)) {
        x["SharedSetting"] = A.load.Int32(ptr + 20);
      } else {
        delete x["SharedSetting"];
      }
      if (A.load.Bool(ptr + 31)) {
        x["UserEditable"] = A.load.Bool(ptr + 24);
      } else {
        delete x["UserEditable"];
      }
      if (A.load.Bool(ptr + 32)) {
        x["DeviceEditable"] = A.load.Bool(ptr + 25);
      } else {
        delete x["DeviceEditable"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ProxySettingsType": (ref: heap.Ref<string>): number => {
      const idx = ["Direct", "Manual", "PAC", "WPAD"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ManagedProxySettingsType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Enum(ptr + 12, -1);
        A.store.Enum(ptr + 16, -1);
        A.store.Enum(ptr + 20, -1);
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 27, false);
        A.store.Bool(ptr + 25, false);
      } else {
        A.store.Bool(ptr + 28, true);
        A.store.Enum(ptr + 0, ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["Active"] as string));
        A.store.Ref(ptr + 4, x["Effective"]);
        A.store.Enum(ptr + 8, ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["UserPolicy"] as string));
        A.store.Enum(ptr + 12, ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["DevicePolicy"] as string));
        A.store.Enum(ptr + 16, ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["UserSetting"] as string));
        A.store.Enum(ptr + 20, ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["SharedSetting"] as string));
        A.store.Bool(ptr + 26, "UserEditable" in x ? true : false);
        A.store.Bool(ptr + 24, x["UserEditable"] ? true : false);
        A.store.Bool(ptr + 27, "DeviceEditable" in x ? true : false);
        A.store.Bool(ptr + 25, x["DeviceEditable"] ? true : false);
      }
    },
    "load_ManagedProxySettingsType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["Active"] = A.load.Enum(ptr + 0, ["Direct", "Manual", "PAC", "WPAD"]);
      x["Effective"] = A.load.Ref(ptr + 4, undefined);
      x["UserPolicy"] = A.load.Enum(ptr + 8, ["Direct", "Manual", "PAC", "WPAD"]);
      x["DevicePolicy"] = A.load.Enum(ptr + 12, ["Direct", "Manual", "PAC", "WPAD"]);
      x["UserSetting"] = A.load.Enum(ptr + 16, ["Direct", "Manual", "PAC", "WPAD"]);
      x["SharedSetting"] = A.load.Enum(ptr + 20, ["Direct", "Manual", "PAC", "WPAD"]);
      if (A.load.Bool(ptr + 26)) {
        x["UserEditable"] = A.load.Bool(ptr + 24);
      } else {
        delete x["UserEditable"];
      }
      if (A.load.Bool(ptr + 27)) {
        x["DeviceEditable"] = A.load.Bool(ptr + 25);
      } else {
        delete x["DeviceEditable"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManagedProxyLocation": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 66, false);

        A.store.Bool(ptr + 0 + 28, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Ref(ptr + 0 + 8, undefined);
        A.store.Ref(ptr + 0 + 12, undefined);
        A.store.Ref(ptr + 0 + 16, undefined);
        A.store.Ref(ptr + 0 + 20, undefined);
        A.store.Bool(ptr + 0 + 26, false);
        A.store.Bool(ptr + 0 + 24, false);
        A.store.Bool(ptr + 0 + 27, false);
        A.store.Bool(ptr + 0 + 25, false);

        A.store.Bool(ptr + 32 + 33, false);
        A.store.Bool(ptr + 32 + 26, false);
        A.store.Int32(ptr + 32 + 0, 0);
        A.store.Ref(ptr + 32 + 4, undefined);
        A.store.Bool(ptr + 32 + 27, false);
        A.store.Int32(ptr + 32 + 8, 0);
        A.store.Bool(ptr + 32 + 28, false);
        A.store.Int32(ptr + 32 + 12, 0);
        A.store.Bool(ptr + 32 + 29, false);
        A.store.Int32(ptr + 32 + 16, 0);
        A.store.Bool(ptr + 32 + 30, false);
        A.store.Int32(ptr + 32 + 20, 0);
        A.store.Bool(ptr + 32 + 31, false);
        A.store.Bool(ptr + 32 + 24, false);
        A.store.Bool(ptr + 32 + 32, false);
        A.store.Bool(ptr + 32 + 25, false);
      } else {
        A.store.Bool(ptr + 66, true);

        if (typeof x["Host"] === "undefined") {
          A.store.Bool(ptr + 0 + 28, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Ref(ptr + 0 + 8, undefined);
          A.store.Ref(ptr + 0 + 12, undefined);
          A.store.Ref(ptr + 0 + 16, undefined);
          A.store.Ref(ptr + 0 + 20, undefined);
          A.store.Bool(ptr + 0 + 26, false);
          A.store.Bool(ptr + 0 + 24, false);
          A.store.Bool(ptr + 0 + 27, false);
          A.store.Bool(ptr + 0 + 25, false);
        } else {
          A.store.Bool(ptr + 0 + 28, true);
          A.store.Ref(ptr + 0 + 0, x["Host"]["Active"]);
          A.store.Ref(ptr + 0 + 4, x["Host"]["Effective"]);
          A.store.Ref(ptr + 0 + 8, x["Host"]["UserPolicy"]);
          A.store.Ref(ptr + 0 + 12, x["Host"]["DevicePolicy"]);
          A.store.Ref(ptr + 0 + 16, x["Host"]["UserSetting"]);
          A.store.Ref(ptr + 0 + 20, x["Host"]["SharedSetting"]);
          A.store.Bool(ptr + 0 + 26, "UserEditable" in x["Host"] ? true : false);
          A.store.Bool(ptr + 0 + 24, x["Host"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 0 + 27, "DeviceEditable" in x["Host"] ? true : false);
          A.store.Bool(ptr + 0 + 25, x["Host"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["Port"] === "undefined") {
          A.store.Bool(ptr + 32 + 33, false);
          A.store.Bool(ptr + 32 + 26, false);
          A.store.Int32(ptr + 32 + 0, 0);
          A.store.Ref(ptr + 32 + 4, undefined);
          A.store.Bool(ptr + 32 + 27, false);
          A.store.Int32(ptr + 32 + 8, 0);
          A.store.Bool(ptr + 32 + 28, false);
          A.store.Int32(ptr + 32 + 12, 0);
          A.store.Bool(ptr + 32 + 29, false);
          A.store.Int32(ptr + 32 + 16, 0);
          A.store.Bool(ptr + 32 + 30, false);
          A.store.Int32(ptr + 32 + 20, 0);
          A.store.Bool(ptr + 32 + 31, false);
          A.store.Bool(ptr + 32 + 24, false);
          A.store.Bool(ptr + 32 + 32, false);
          A.store.Bool(ptr + 32 + 25, false);
        } else {
          A.store.Bool(ptr + 32 + 33, true);
          A.store.Bool(ptr + 32 + 26, "Active" in x["Port"] ? true : false);
          A.store.Int32(ptr + 32 + 0, x["Port"]["Active"] === undefined ? 0 : (x["Port"]["Active"] as number));
          A.store.Ref(ptr + 32 + 4, x["Port"]["Effective"]);
          A.store.Bool(ptr + 32 + 27, "UserPolicy" in x["Port"] ? true : false);
          A.store.Int32(ptr + 32 + 8, x["Port"]["UserPolicy"] === undefined ? 0 : (x["Port"]["UserPolicy"] as number));
          A.store.Bool(ptr + 32 + 28, "DevicePolicy" in x["Port"] ? true : false);
          A.store.Int32(
            ptr + 32 + 12,
            x["Port"]["DevicePolicy"] === undefined ? 0 : (x["Port"]["DevicePolicy"] as number)
          );
          A.store.Bool(ptr + 32 + 29, "UserSetting" in x["Port"] ? true : false);
          A.store.Int32(
            ptr + 32 + 16,
            x["Port"]["UserSetting"] === undefined ? 0 : (x["Port"]["UserSetting"] as number)
          );
          A.store.Bool(ptr + 32 + 30, "SharedSetting" in x["Port"] ? true : false);
          A.store.Int32(
            ptr + 32 + 20,
            x["Port"]["SharedSetting"] === undefined ? 0 : (x["Port"]["SharedSetting"] as number)
          );
          A.store.Bool(ptr + 32 + 31, "UserEditable" in x["Port"] ? true : false);
          A.store.Bool(ptr + 32 + 24, x["Port"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 32 + 32, "DeviceEditable" in x["Port"] ? true : false);
          A.store.Bool(ptr + 32 + 25, x["Port"]["DeviceEditable"] ? true : false);
        }
      }
    },
    "load_ManagedProxyLocation": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 28)) {
        x["Host"] = {};
        x["Host"]["Active"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["Host"]["Effective"] = A.load.Ref(ptr + 0 + 4, undefined);
        x["Host"]["UserPolicy"] = A.load.Ref(ptr + 0 + 8, undefined);
        x["Host"]["DevicePolicy"] = A.load.Ref(ptr + 0 + 12, undefined);
        x["Host"]["UserSetting"] = A.load.Ref(ptr + 0 + 16, undefined);
        x["Host"]["SharedSetting"] = A.load.Ref(ptr + 0 + 20, undefined);
        if (A.load.Bool(ptr + 0 + 26)) {
          x["Host"]["UserEditable"] = A.load.Bool(ptr + 0 + 24);
        } else {
          delete x["Host"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 0 + 27)) {
          x["Host"]["DeviceEditable"] = A.load.Bool(ptr + 0 + 25);
        } else {
          delete x["Host"]["DeviceEditable"];
        }
      } else {
        delete x["Host"];
      }
      if (A.load.Bool(ptr + 32 + 33)) {
        x["Port"] = {};
        if (A.load.Bool(ptr + 32 + 26)) {
          x["Port"]["Active"] = A.load.Int32(ptr + 32 + 0);
        } else {
          delete x["Port"]["Active"];
        }
        x["Port"]["Effective"] = A.load.Ref(ptr + 32 + 4, undefined);
        if (A.load.Bool(ptr + 32 + 27)) {
          x["Port"]["UserPolicy"] = A.load.Int32(ptr + 32 + 8);
        } else {
          delete x["Port"]["UserPolicy"];
        }
        if (A.load.Bool(ptr + 32 + 28)) {
          x["Port"]["DevicePolicy"] = A.load.Int32(ptr + 32 + 12);
        } else {
          delete x["Port"]["DevicePolicy"];
        }
        if (A.load.Bool(ptr + 32 + 29)) {
          x["Port"]["UserSetting"] = A.load.Int32(ptr + 32 + 16);
        } else {
          delete x["Port"]["UserSetting"];
        }
        if (A.load.Bool(ptr + 32 + 30)) {
          x["Port"]["SharedSetting"] = A.load.Int32(ptr + 32 + 20);
        } else {
          delete x["Port"]["SharedSetting"];
        }
        if (A.load.Bool(ptr + 32 + 31)) {
          x["Port"]["UserEditable"] = A.load.Bool(ptr + 32 + 24);
        } else {
          delete x["Port"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 32 + 32)) {
          x["Port"]["DeviceEditable"] = A.load.Bool(ptr + 32 + 25);
        } else {
          delete x["Port"]["DeviceEditable"];
        }
      } else {
        delete x["Port"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManagedManualProxySettings": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 271, false);

        A.store.Bool(ptr + 0 + 66, false);

        A.store.Bool(ptr + 0 + 0 + 28, false);
        A.store.Ref(ptr + 0 + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 0 + 4, undefined);
        A.store.Ref(ptr + 0 + 0 + 8, undefined);
        A.store.Ref(ptr + 0 + 0 + 12, undefined);
        A.store.Ref(ptr + 0 + 0 + 16, undefined);
        A.store.Ref(ptr + 0 + 0 + 20, undefined);
        A.store.Bool(ptr + 0 + 0 + 26, false);
        A.store.Bool(ptr + 0 + 0 + 24, false);
        A.store.Bool(ptr + 0 + 0 + 27, false);
        A.store.Bool(ptr + 0 + 0 + 25, false);

        A.store.Bool(ptr + 0 + 32 + 33, false);
        A.store.Bool(ptr + 0 + 32 + 26, false);
        A.store.Int32(ptr + 0 + 32 + 0, 0);
        A.store.Ref(ptr + 0 + 32 + 4, undefined);
        A.store.Bool(ptr + 0 + 32 + 27, false);
        A.store.Int32(ptr + 0 + 32 + 8, 0);
        A.store.Bool(ptr + 0 + 32 + 28, false);
        A.store.Int32(ptr + 0 + 32 + 12, 0);
        A.store.Bool(ptr + 0 + 32 + 29, false);
        A.store.Int32(ptr + 0 + 32 + 16, 0);
        A.store.Bool(ptr + 0 + 32 + 30, false);
        A.store.Int32(ptr + 0 + 32 + 20, 0);
        A.store.Bool(ptr + 0 + 32 + 31, false);
        A.store.Bool(ptr + 0 + 32 + 24, false);
        A.store.Bool(ptr + 0 + 32 + 32, false);
        A.store.Bool(ptr + 0 + 32 + 25, false);

        A.store.Bool(ptr + 68 + 66, false);

        A.store.Bool(ptr + 68 + 0 + 28, false);
        A.store.Ref(ptr + 68 + 0 + 0, undefined);
        A.store.Ref(ptr + 68 + 0 + 4, undefined);
        A.store.Ref(ptr + 68 + 0 + 8, undefined);
        A.store.Ref(ptr + 68 + 0 + 12, undefined);
        A.store.Ref(ptr + 68 + 0 + 16, undefined);
        A.store.Ref(ptr + 68 + 0 + 20, undefined);
        A.store.Bool(ptr + 68 + 0 + 26, false);
        A.store.Bool(ptr + 68 + 0 + 24, false);
        A.store.Bool(ptr + 68 + 0 + 27, false);
        A.store.Bool(ptr + 68 + 0 + 25, false);

        A.store.Bool(ptr + 68 + 32 + 33, false);
        A.store.Bool(ptr + 68 + 32 + 26, false);
        A.store.Int32(ptr + 68 + 32 + 0, 0);
        A.store.Ref(ptr + 68 + 32 + 4, undefined);
        A.store.Bool(ptr + 68 + 32 + 27, false);
        A.store.Int32(ptr + 68 + 32 + 8, 0);
        A.store.Bool(ptr + 68 + 32 + 28, false);
        A.store.Int32(ptr + 68 + 32 + 12, 0);
        A.store.Bool(ptr + 68 + 32 + 29, false);
        A.store.Int32(ptr + 68 + 32 + 16, 0);
        A.store.Bool(ptr + 68 + 32 + 30, false);
        A.store.Int32(ptr + 68 + 32 + 20, 0);
        A.store.Bool(ptr + 68 + 32 + 31, false);
        A.store.Bool(ptr + 68 + 32 + 24, false);
        A.store.Bool(ptr + 68 + 32 + 32, false);
        A.store.Bool(ptr + 68 + 32 + 25, false);

        A.store.Bool(ptr + 136 + 66, false);

        A.store.Bool(ptr + 136 + 0 + 28, false);
        A.store.Ref(ptr + 136 + 0 + 0, undefined);
        A.store.Ref(ptr + 136 + 0 + 4, undefined);
        A.store.Ref(ptr + 136 + 0 + 8, undefined);
        A.store.Ref(ptr + 136 + 0 + 12, undefined);
        A.store.Ref(ptr + 136 + 0 + 16, undefined);
        A.store.Ref(ptr + 136 + 0 + 20, undefined);
        A.store.Bool(ptr + 136 + 0 + 26, false);
        A.store.Bool(ptr + 136 + 0 + 24, false);
        A.store.Bool(ptr + 136 + 0 + 27, false);
        A.store.Bool(ptr + 136 + 0 + 25, false);

        A.store.Bool(ptr + 136 + 32 + 33, false);
        A.store.Bool(ptr + 136 + 32 + 26, false);
        A.store.Int32(ptr + 136 + 32 + 0, 0);
        A.store.Ref(ptr + 136 + 32 + 4, undefined);
        A.store.Bool(ptr + 136 + 32 + 27, false);
        A.store.Int32(ptr + 136 + 32 + 8, 0);
        A.store.Bool(ptr + 136 + 32 + 28, false);
        A.store.Int32(ptr + 136 + 32 + 12, 0);
        A.store.Bool(ptr + 136 + 32 + 29, false);
        A.store.Int32(ptr + 136 + 32 + 16, 0);
        A.store.Bool(ptr + 136 + 32 + 30, false);
        A.store.Int32(ptr + 136 + 32 + 20, 0);
        A.store.Bool(ptr + 136 + 32 + 31, false);
        A.store.Bool(ptr + 136 + 32 + 24, false);
        A.store.Bool(ptr + 136 + 32 + 32, false);
        A.store.Bool(ptr + 136 + 32 + 25, false);

        A.store.Bool(ptr + 204 + 66, false);

        A.store.Bool(ptr + 204 + 0 + 28, false);
        A.store.Ref(ptr + 204 + 0 + 0, undefined);
        A.store.Ref(ptr + 204 + 0 + 4, undefined);
        A.store.Ref(ptr + 204 + 0 + 8, undefined);
        A.store.Ref(ptr + 204 + 0 + 12, undefined);
        A.store.Ref(ptr + 204 + 0 + 16, undefined);
        A.store.Ref(ptr + 204 + 0 + 20, undefined);
        A.store.Bool(ptr + 204 + 0 + 26, false);
        A.store.Bool(ptr + 204 + 0 + 24, false);
        A.store.Bool(ptr + 204 + 0 + 27, false);
        A.store.Bool(ptr + 204 + 0 + 25, false);

        A.store.Bool(ptr + 204 + 32 + 33, false);
        A.store.Bool(ptr + 204 + 32 + 26, false);
        A.store.Int32(ptr + 204 + 32 + 0, 0);
        A.store.Ref(ptr + 204 + 32 + 4, undefined);
        A.store.Bool(ptr + 204 + 32 + 27, false);
        A.store.Int32(ptr + 204 + 32 + 8, 0);
        A.store.Bool(ptr + 204 + 32 + 28, false);
        A.store.Int32(ptr + 204 + 32 + 12, 0);
        A.store.Bool(ptr + 204 + 32 + 29, false);
        A.store.Int32(ptr + 204 + 32 + 16, 0);
        A.store.Bool(ptr + 204 + 32 + 30, false);
        A.store.Int32(ptr + 204 + 32 + 20, 0);
        A.store.Bool(ptr + 204 + 32 + 31, false);
        A.store.Bool(ptr + 204 + 32 + 24, false);
        A.store.Bool(ptr + 204 + 32 + 32, false);
        A.store.Bool(ptr + 204 + 32 + 25, false);
      } else {
        A.store.Bool(ptr + 271, true);

        if (typeof x["HTTPProxy"] === "undefined") {
          A.store.Bool(ptr + 0 + 66, false);

          A.store.Bool(ptr + 0 + 0 + 28, false);
          A.store.Ref(ptr + 0 + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 0 + 4, undefined);
          A.store.Ref(ptr + 0 + 0 + 8, undefined);
          A.store.Ref(ptr + 0 + 0 + 12, undefined);
          A.store.Ref(ptr + 0 + 0 + 16, undefined);
          A.store.Ref(ptr + 0 + 0 + 20, undefined);
          A.store.Bool(ptr + 0 + 0 + 26, false);
          A.store.Bool(ptr + 0 + 0 + 24, false);
          A.store.Bool(ptr + 0 + 0 + 27, false);
          A.store.Bool(ptr + 0 + 0 + 25, false);

          A.store.Bool(ptr + 0 + 32 + 33, false);
          A.store.Bool(ptr + 0 + 32 + 26, false);
          A.store.Int32(ptr + 0 + 32 + 0, 0);
          A.store.Ref(ptr + 0 + 32 + 4, undefined);
          A.store.Bool(ptr + 0 + 32 + 27, false);
          A.store.Int32(ptr + 0 + 32 + 8, 0);
          A.store.Bool(ptr + 0 + 32 + 28, false);
          A.store.Int32(ptr + 0 + 32 + 12, 0);
          A.store.Bool(ptr + 0 + 32 + 29, false);
          A.store.Int32(ptr + 0 + 32 + 16, 0);
          A.store.Bool(ptr + 0 + 32 + 30, false);
          A.store.Int32(ptr + 0 + 32 + 20, 0);
          A.store.Bool(ptr + 0 + 32 + 31, false);
          A.store.Bool(ptr + 0 + 32 + 24, false);
          A.store.Bool(ptr + 0 + 32 + 32, false);
          A.store.Bool(ptr + 0 + 32 + 25, false);
        } else {
          A.store.Bool(ptr + 0 + 66, true);

          if (typeof x["HTTPProxy"]["Host"] === "undefined") {
            A.store.Bool(ptr + 0 + 0 + 28, false);
            A.store.Ref(ptr + 0 + 0 + 0, undefined);
            A.store.Ref(ptr + 0 + 0 + 4, undefined);
            A.store.Ref(ptr + 0 + 0 + 8, undefined);
            A.store.Ref(ptr + 0 + 0 + 12, undefined);
            A.store.Ref(ptr + 0 + 0 + 16, undefined);
            A.store.Ref(ptr + 0 + 0 + 20, undefined);
            A.store.Bool(ptr + 0 + 0 + 26, false);
            A.store.Bool(ptr + 0 + 0 + 24, false);
            A.store.Bool(ptr + 0 + 0 + 27, false);
            A.store.Bool(ptr + 0 + 0 + 25, false);
          } else {
            A.store.Bool(ptr + 0 + 0 + 28, true);
            A.store.Ref(ptr + 0 + 0 + 0, x["HTTPProxy"]["Host"]["Active"]);
            A.store.Ref(ptr + 0 + 0 + 4, x["HTTPProxy"]["Host"]["Effective"]);
            A.store.Ref(ptr + 0 + 0 + 8, x["HTTPProxy"]["Host"]["UserPolicy"]);
            A.store.Ref(ptr + 0 + 0 + 12, x["HTTPProxy"]["Host"]["DevicePolicy"]);
            A.store.Ref(ptr + 0 + 0 + 16, x["HTTPProxy"]["Host"]["UserSetting"]);
            A.store.Ref(ptr + 0 + 0 + 20, x["HTTPProxy"]["Host"]["SharedSetting"]);
            A.store.Bool(ptr + 0 + 0 + 26, "UserEditable" in x["HTTPProxy"]["Host"] ? true : false);
            A.store.Bool(ptr + 0 + 0 + 24, x["HTTPProxy"]["Host"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 0 + 0 + 27, "DeviceEditable" in x["HTTPProxy"]["Host"] ? true : false);
            A.store.Bool(ptr + 0 + 0 + 25, x["HTTPProxy"]["Host"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["HTTPProxy"]["Port"] === "undefined") {
            A.store.Bool(ptr + 0 + 32 + 33, false);
            A.store.Bool(ptr + 0 + 32 + 26, false);
            A.store.Int32(ptr + 0 + 32 + 0, 0);
            A.store.Ref(ptr + 0 + 32 + 4, undefined);
            A.store.Bool(ptr + 0 + 32 + 27, false);
            A.store.Int32(ptr + 0 + 32 + 8, 0);
            A.store.Bool(ptr + 0 + 32 + 28, false);
            A.store.Int32(ptr + 0 + 32 + 12, 0);
            A.store.Bool(ptr + 0 + 32 + 29, false);
            A.store.Int32(ptr + 0 + 32 + 16, 0);
            A.store.Bool(ptr + 0 + 32 + 30, false);
            A.store.Int32(ptr + 0 + 32 + 20, 0);
            A.store.Bool(ptr + 0 + 32 + 31, false);
            A.store.Bool(ptr + 0 + 32 + 24, false);
            A.store.Bool(ptr + 0 + 32 + 32, false);
            A.store.Bool(ptr + 0 + 32 + 25, false);
          } else {
            A.store.Bool(ptr + 0 + 32 + 33, true);
            A.store.Bool(ptr + 0 + 32 + 26, "Active" in x["HTTPProxy"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 0 + 32 + 0,
              x["HTTPProxy"]["Port"]["Active"] === undefined ? 0 : (x["HTTPProxy"]["Port"]["Active"] as number)
            );
            A.store.Ref(ptr + 0 + 32 + 4, x["HTTPProxy"]["Port"]["Effective"]);
            A.store.Bool(ptr + 0 + 32 + 27, "UserPolicy" in x["HTTPProxy"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 0 + 32 + 8,
              x["HTTPProxy"]["Port"]["UserPolicy"] === undefined ? 0 : (x["HTTPProxy"]["Port"]["UserPolicy"] as number)
            );
            A.store.Bool(ptr + 0 + 32 + 28, "DevicePolicy" in x["HTTPProxy"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 0 + 32 + 12,
              x["HTTPProxy"]["Port"]["DevicePolicy"] === undefined
                ? 0
                : (x["HTTPProxy"]["Port"]["DevicePolicy"] as number)
            );
            A.store.Bool(ptr + 0 + 32 + 29, "UserSetting" in x["HTTPProxy"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 0 + 32 + 16,
              x["HTTPProxy"]["Port"]["UserSetting"] === undefined
                ? 0
                : (x["HTTPProxy"]["Port"]["UserSetting"] as number)
            );
            A.store.Bool(ptr + 0 + 32 + 30, "SharedSetting" in x["HTTPProxy"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 0 + 32 + 20,
              x["HTTPProxy"]["Port"]["SharedSetting"] === undefined
                ? 0
                : (x["HTTPProxy"]["Port"]["SharedSetting"] as number)
            );
            A.store.Bool(ptr + 0 + 32 + 31, "UserEditable" in x["HTTPProxy"]["Port"] ? true : false);
            A.store.Bool(ptr + 0 + 32 + 24, x["HTTPProxy"]["Port"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 0 + 32 + 32, "DeviceEditable" in x["HTTPProxy"]["Port"] ? true : false);
            A.store.Bool(ptr + 0 + 32 + 25, x["HTTPProxy"]["Port"]["DeviceEditable"] ? true : false);
          }
        }

        if (typeof x["SecureHTTPProxy"] === "undefined") {
          A.store.Bool(ptr + 68 + 66, false);

          A.store.Bool(ptr + 68 + 0 + 28, false);
          A.store.Ref(ptr + 68 + 0 + 0, undefined);
          A.store.Ref(ptr + 68 + 0 + 4, undefined);
          A.store.Ref(ptr + 68 + 0 + 8, undefined);
          A.store.Ref(ptr + 68 + 0 + 12, undefined);
          A.store.Ref(ptr + 68 + 0 + 16, undefined);
          A.store.Ref(ptr + 68 + 0 + 20, undefined);
          A.store.Bool(ptr + 68 + 0 + 26, false);
          A.store.Bool(ptr + 68 + 0 + 24, false);
          A.store.Bool(ptr + 68 + 0 + 27, false);
          A.store.Bool(ptr + 68 + 0 + 25, false);

          A.store.Bool(ptr + 68 + 32 + 33, false);
          A.store.Bool(ptr + 68 + 32 + 26, false);
          A.store.Int32(ptr + 68 + 32 + 0, 0);
          A.store.Ref(ptr + 68 + 32 + 4, undefined);
          A.store.Bool(ptr + 68 + 32 + 27, false);
          A.store.Int32(ptr + 68 + 32 + 8, 0);
          A.store.Bool(ptr + 68 + 32 + 28, false);
          A.store.Int32(ptr + 68 + 32 + 12, 0);
          A.store.Bool(ptr + 68 + 32 + 29, false);
          A.store.Int32(ptr + 68 + 32 + 16, 0);
          A.store.Bool(ptr + 68 + 32 + 30, false);
          A.store.Int32(ptr + 68 + 32 + 20, 0);
          A.store.Bool(ptr + 68 + 32 + 31, false);
          A.store.Bool(ptr + 68 + 32 + 24, false);
          A.store.Bool(ptr + 68 + 32 + 32, false);
          A.store.Bool(ptr + 68 + 32 + 25, false);
        } else {
          A.store.Bool(ptr + 68 + 66, true);

          if (typeof x["SecureHTTPProxy"]["Host"] === "undefined") {
            A.store.Bool(ptr + 68 + 0 + 28, false);
            A.store.Ref(ptr + 68 + 0 + 0, undefined);
            A.store.Ref(ptr + 68 + 0 + 4, undefined);
            A.store.Ref(ptr + 68 + 0 + 8, undefined);
            A.store.Ref(ptr + 68 + 0 + 12, undefined);
            A.store.Ref(ptr + 68 + 0 + 16, undefined);
            A.store.Ref(ptr + 68 + 0 + 20, undefined);
            A.store.Bool(ptr + 68 + 0 + 26, false);
            A.store.Bool(ptr + 68 + 0 + 24, false);
            A.store.Bool(ptr + 68 + 0 + 27, false);
            A.store.Bool(ptr + 68 + 0 + 25, false);
          } else {
            A.store.Bool(ptr + 68 + 0 + 28, true);
            A.store.Ref(ptr + 68 + 0 + 0, x["SecureHTTPProxy"]["Host"]["Active"]);
            A.store.Ref(ptr + 68 + 0 + 4, x["SecureHTTPProxy"]["Host"]["Effective"]);
            A.store.Ref(ptr + 68 + 0 + 8, x["SecureHTTPProxy"]["Host"]["UserPolicy"]);
            A.store.Ref(ptr + 68 + 0 + 12, x["SecureHTTPProxy"]["Host"]["DevicePolicy"]);
            A.store.Ref(ptr + 68 + 0 + 16, x["SecureHTTPProxy"]["Host"]["UserSetting"]);
            A.store.Ref(ptr + 68 + 0 + 20, x["SecureHTTPProxy"]["Host"]["SharedSetting"]);
            A.store.Bool(ptr + 68 + 0 + 26, "UserEditable" in x["SecureHTTPProxy"]["Host"] ? true : false);
            A.store.Bool(ptr + 68 + 0 + 24, x["SecureHTTPProxy"]["Host"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 68 + 0 + 27, "DeviceEditable" in x["SecureHTTPProxy"]["Host"] ? true : false);
            A.store.Bool(ptr + 68 + 0 + 25, x["SecureHTTPProxy"]["Host"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["SecureHTTPProxy"]["Port"] === "undefined") {
            A.store.Bool(ptr + 68 + 32 + 33, false);
            A.store.Bool(ptr + 68 + 32 + 26, false);
            A.store.Int32(ptr + 68 + 32 + 0, 0);
            A.store.Ref(ptr + 68 + 32 + 4, undefined);
            A.store.Bool(ptr + 68 + 32 + 27, false);
            A.store.Int32(ptr + 68 + 32 + 8, 0);
            A.store.Bool(ptr + 68 + 32 + 28, false);
            A.store.Int32(ptr + 68 + 32 + 12, 0);
            A.store.Bool(ptr + 68 + 32 + 29, false);
            A.store.Int32(ptr + 68 + 32 + 16, 0);
            A.store.Bool(ptr + 68 + 32 + 30, false);
            A.store.Int32(ptr + 68 + 32 + 20, 0);
            A.store.Bool(ptr + 68 + 32 + 31, false);
            A.store.Bool(ptr + 68 + 32 + 24, false);
            A.store.Bool(ptr + 68 + 32 + 32, false);
            A.store.Bool(ptr + 68 + 32 + 25, false);
          } else {
            A.store.Bool(ptr + 68 + 32 + 33, true);
            A.store.Bool(ptr + 68 + 32 + 26, "Active" in x["SecureHTTPProxy"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 68 + 32 + 0,
              x["SecureHTTPProxy"]["Port"]["Active"] === undefined
                ? 0
                : (x["SecureHTTPProxy"]["Port"]["Active"] as number)
            );
            A.store.Ref(ptr + 68 + 32 + 4, x["SecureHTTPProxy"]["Port"]["Effective"]);
            A.store.Bool(ptr + 68 + 32 + 27, "UserPolicy" in x["SecureHTTPProxy"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 68 + 32 + 8,
              x["SecureHTTPProxy"]["Port"]["UserPolicy"] === undefined
                ? 0
                : (x["SecureHTTPProxy"]["Port"]["UserPolicy"] as number)
            );
            A.store.Bool(ptr + 68 + 32 + 28, "DevicePolicy" in x["SecureHTTPProxy"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 68 + 32 + 12,
              x["SecureHTTPProxy"]["Port"]["DevicePolicy"] === undefined
                ? 0
                : (x["SecureHTTPProxy"]["Port"]["DevicePolicy"] as number)
            );
            A.store.Bool(ptr + 68 + 32 + 29, "UserSetting" in x["SecureHTTPProxy"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 68 + 32 + 16,
              x["SecureHTTPProxy"]["Port"]["UserSetting"] === undefined
                ? 0
                : (x["SecureHTTPProxy"]["Port"]["UserSetting"] as number)
            );
            A.store.Bool(ptr + 68 + 32 + 30, "SharedSetting" in x["SecureHTTPProxy"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 68 + 32 + 20,
              x["SecureHTTPProxy"]["Port"]["SharedSetting"] === undefined
                ? 0
                : (x["SecureHTTPProxy"]["Port"]["SharedSetting"] as number)
            );
            A.store.Bool(ptr + 68 + 32 + 31, "UserEditable" in x["SecureHTTPProxy"]["Port"] ? true : false);
            A.store.Bool(ptr + 68 + 32 + 24, x["SecureHTTPProxy"]["Port"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 68 + 32 + 32, "DeviceEditable" in x["SecureHTTPProxy"]["Port"] ? true : false);
            A.store.Bool(ptr + 68 + 32 + 25, x["SecureHTTPProxy"]["Port"]["DeviceEditable"] ? true : false);
          }
        }

        if (typeof x["FTPProxy"] === "undefined") {
          A.store.Bool(ptr + 136 + 66, false);

          A.store.Bool(ptr + 136 + 0 + 28, false);
          A.store.Ref(ptr + 136 + 0 + 0, undefined);
          A.store.Ref(ptr + 136 + 0 + 4, undefined);
          A.store.Ref(ptr + 136 + 0 + 8, undefined);
          A.store.Ref(ptr + 136 + 0 + 12, undefined);
          A.store.Ref(ptr + 136 + 0 + 16, undefined);
          A.store.Ref(ptr + 136 + 0 + 20, undefined);
          A.store.Bool(ptr + 136 + 0 + 26, false);
          A.store.Bool(ptr + 136 + 0 + 24, false);
          A.store.Bool(ptr + 136 + 0 + 27, false);
          A.store.Bool(ptr + 136 + 0 + 25, false);

          A.store.Bool(ptr + 136 + 32 + 33, false);
          A.store.Bool(ptr + 136 + 32 + 26, false);
          A.store.Int32(ptr + 136 + 32 + 0, 0);
          A.store.Ref(ptr + 136 + 32 + 4, undefined);
          A.store.Bool(ptr + 136 + 32 + 27, false);
          A.store.Int32(ptr + 136 + 32 + 8, 0);
          A.store.Bool(ptr + 136 + 32 + 28, false);
          A.store.Int32(ptr + 136 + 32 + 12, 0);
          A.store.Bool(ptr + 136 + 32 + 29, false);
          A.store.Int32(ptr + 136 + 32 + 16, 0);
          A.store.Bool(ptr + 136 + 32 + 30, false);
          A.store.Int32(ptr + 136 + 32 + 20, 0);
          A.store.Bool(ptr + 136 + 32 + 31, false);
          A.store.Bool(ptr + 136 + 32 + 24, false);
          A.store.Bool(ptr + 136 + 32 + 32, false);
          A.store.Bool(ptr + 136 + 32 + 25, false);
        } else {
          A.store.Bool(ptr + 136 + 66, true);

          if (typeof x["FTPProxy"]["Host"] === "undefined") {
            A.store.Bool(ptr + 136 + 0 + 28, false);
            A.store.Ref(ptr + 136 + 0 + 0, undefined);
            A.store.Ref(ptr + 136 + 0 + 4, undefined);
            A.store.Ref(ptr + 136 + 0 + 8, undefined);
            A.store.Ref(ptr + 136 + 0 + 12, undefined);
            A.store.Ref(ptr + 136 + 0 + 16, undefined);
            A.store.Ref(ptr + 136 + 0 + 20, undefined);
            A.store.Bool(ptr + 136 + 0 + 26, false);
            A.store.Bool(ptr + 136 + 0 + 24, false);
            A.store.Bool(ptr + 136 + 0 + 27, false);
            A.store.Bool(ptr + 136 + 0 + 25, false);
          } else {
            A.store.Bool(ptr + 136 + 0 + 28, true);
            A.store.Ref(ptr + 136 + 0 + 0, x["FTPProxy"]["Host"]["Active"]);
            A.store.Ref(ptr + 136 + 0 + 4, x["FTPProxy"]["Host"]["Effective"]);
            A.store.Ref(ptr + 136 + 0 + 8, x["FTPProxy"]["Host"]["UserPolicy"]);
            A.store.Ref(ptr + 136 + 0 + 12, x["FTPProxy"]["Host"]["DevicePolicy"]);
            A.store.Ref(ptr + 136 + 0 + 16, x["FTPProxy"]["Host"]["UserSetting"]);
            A.store.Ref(ptr + 136 + 0 + 20, x["FTPProxy"]["Host"]["SharedSetting"]);
            A.store.Bool(ptr + 136 + 0 + 26, "UserEditable" in x["FTPProxy"]["Host"] ? true : false);
            A.store.Bool(ptr + 136 + 0 + 24, x["FTPProxy"]["Host"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 136 + 0 + 27, "DeviceEditable" in x["FTPProxy"]["Host"] ? true : false);
            A.store.Bool(ptr + 136 + 0 + 25, x["FTPProxy"]["Host"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["FTPProxy"]["Port"] === "undefined") {
            A.store.Bool(ptr + 136 + 32 + 33, false);
            A.store.Bool(ptr + 136 + 32 + 26, false);
            A.store.Int32(ptr + 136 + 32 + 0, 0);
            A.store.Ref(ptr + 136 + 32 + 4, undefined);
            A.store.Bool(ptr + 136 + 32 + 27, false);
            A.store.Int32(ptr + 136 + 32 + 8, 0);
            A.store.Bool(ptr + 136 + 32 + 28, false);
            A.store.Int32(ptr + 136 + 32 + 12, 0);
            A.store.Bool(ptr + 136 + 32 + 29, false);
            A.store.Int32(ptr + 136 + 32 + 16, 0);
            A.store.Bool(ptr + 136 + 32 + 30, false);
            A.store.Int32(ptr + 136 + 32 + 20, 0);
            A.store.Bool(ptr + 136 + 32 + 31, false);
            A.store.Bool(ptr + 136 + 32 + 24, false);
            A.store.Bool(ptr + 136 + 32 + 32, false);
            A.store.Bool(ptr + 136 + 32 + 25, false);
          } else {
            A.store.Bool(ptr + 136 + 32 + 33, true);
            A.store.Bool(ptr + 136 + 32 + 26, "Active" in x["FTPProxy"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 136 + 32 + 0,
              x["FTPProxy"]["Port"]["Active"] === undefined ? 0 : (x["FTPProxy"]["Port"]["Active"] as number)
            );
            A.store.Ref(ptr + 136 + 32 + 4, x["FTPProxy"]["Port"]["Effective"]);
            A.store.Bool(ptr + 136 + 32 + 27, "UserPolicy" in x["FTPProxy"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 136 + 32 + 8,
              x["FTPProxy"]["Port"]["UserPolicy"] === undefined ? 0 : (x["FTPProxy"]["Port"]["UserPolicy"] as number)
            );
            A.store.Bool(ptr + 136 + 32 + 28, "DevicePolicy" in x["FTPProxy"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 136 + 32 + 12,
              x["FTPProxy"]["Port"]["DevicePolicy"] === undefined
                ? 0
                : (x["FTPProxy"]["Port"]["DevicePolicy"] as number)
            );
            A.store.Bool(ptr + 136 + 32 + 29, "UserSetting" in x["FTPProxy"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 136 + 32 + 16,
              x["FTPProxy"]["Port"]["UserSetting"] === undefined ? 0 : (x["FTPProxy"]["Port"]["UserSetting"] as number)
            );
            A.store.Bool(ptr + 136 + 32 + 30, "SharedSetting" in x["FTPProxy"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 136 + 32 + 20,
              x["FTPProxy"]["Port"]["SharedSetting"] === undefined
                ? 0
                : (x["FTPProxy"]["Port"]["SharedSetting"] as number)
            );
            A.store.Bool(ptr + 136 + 32 + 31, "UserEditable" in x["FTPProxy"]["Port"] ? true : false);
            A.store.Bool(ptr + 136 + 32 + 24, x["FTPProxy"]["Port"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 136 + 32 + 32, "DeviceEditable" in x["FTPProxy"]["Port"] ? true : false);
            A.store.Bool(ptr + 136 + 32 + 25, x["FTPProxy"]["Port"]["DeviceEditable"] ? true : false);
          }
        }

        if (typeof x["SOCKS"] === "undefined") {
          A.store.Bool(ptr + 204 + 66, false);

          A.store.Bool(ptr + 204 + 0 + 28, false);
          A.store.Ref(ptr + 204 + 0 + 0, undefined);
          A.store.Ref(ptr + 204 + 0 + 4, undefined);
          A.store.Ref(ptr + 204 + 0 + 8, undefined);
          A.store.Ref(ptr + 204 + 0 + 12, undefined);
          A.store.Ref(ptr + 204 + 0 + 16, undefined);
          A.store.Ref(ptr + 204 + 0 + 20, undefined);
          A.store.Bool(ptr + 204 + 0 + 26, false);
          A.store.Bool(ptr + 204 + 0 + 24, false);
          A.store.Bool(ptr + 204 + 0 + 27, false);
          A.store.Bool(ptr + 204 + 0 + 25, false);

          A.store.Bool(ptr + 204 + 32 + 33, false);
          A.store.Bool(ptr + 204 + 32 + 26, false);
          A.store.Int32(ptr + 204 + 32 + 0, 0);
          A.store.Ref(ptr + 204 + 32 + 4, undefined);
          A.store.Bool(ptr + 204 + 32 + 27, false);
          A.store.Int32(ptr + 204 + 32 + 8, 0);
          A.store.Bool(ptr + 204 + 32 + 28, false);
          A.store.Int32(ptr + 204 + 32 + 12, 0);
          A.store.Bool(ptr + 204 + 32 + 29, false);
          A.store.Int32(ptr + 204 + 32 + 16, 0);
          A.store.Bool(ptr + 204 + 32 + 30, false);
          A.store.Int32(ptr + 204 + 32 + 20, 0);
          A.store.Bool(ptr + 204 + 32 + 31, false);
          A.store.Bool(ptr + 204 + 32 + 24, false);
          A.store.Bool(ptr + 204 + 32 + 32, false);
          A.store.Bool(ptr + 204 + 32 + 25, false);
        } else {
          A.store.Bool(ptr + 204 + 66, true);

          if (typeof x["SOCKS"]["Host"] === "undefined") {
            A.store.Bool(ptr + 204 + 0 + 28, false);
            A.store.Ref(ptr + 204 + 0 + 0, undefined);
            A.store.Ref(ptr + 204 + 0 + 4, undefined);
            A.store.Ref(ptr + 204 + 0 + 8, undefined);
            A.store.Ref(ptr + 204 + 0 + 12, undefined);
            A.store.Ref(ptr + 204 + 0 + 16, undefined);
            A.store.Ref(ptr + 204 + 0 + 20, undefined);
            A.store.Bool(ptr + 204 + 0 + 26, false);
            A.store.Bool(ptr + 204 + 0 + 24, false);
            A.store.Bool(ptr + 204 + 0 + 27, false);
            A.store.Bool(ptr + 204 + 0 + 25, false);
          } else {
            A.store.Bool(ptr + 204 + 0 + 28, true);
            A.store.Ref(ptr + 204 + 0 + 0, x["SOCKS"]["Host"]["Active"]);
            A.store.Ref(ptr + 204 + 0 + 4, x["SOCKS"]["Host"]["Effective"]);
            A.store.Ref(ptr + 204 + 0 + 8, x["SOCKS"]["Host"]["UserPolicy"]);
            A.store.Ref(ptr + 204 + 0 + 12, x["SOCKS"]["Host"]["DevicePolicy"]);
            A.store.Ref(ptr + 204 + 0 + 16, x["SOCKS"]["Host"]["UserSetting"]);
            A.store.Ref(ptr + 204 + 0 + 20, x["SOCKS"]["Host"]["SharedSetting"]);
            A.store.Bool(ptr + 204 + 0 + 26, "UserEditable" in x["SOCKS"]["Host"] ? true : false);
            A.store.Bool(ptr + 204 + 0 + 24, x["SOCKS"]["Host"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 204 + 0 + 27, "DeviceEditable" in x["SOCKS"]["Host"] ? true : false);
            A.store.Bool(ptr + 204 + 0 + 25, x["SOCKS"]["Host"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["SOCKS"]["Port"] === "undefined") {
            A.store.Bool(ptr + 204 + 32 + 33, false);
            A.store.Bool(ptr + 204 + 32 + 26, false);
            A.store.Int32(ptr + 204 + 32 + 0, 0);
            A.store.Ref(ptr + 204 + 32 + 4, undefined);
            A.store.Bool(ptr + 204 + 32 + 27, false);
            A.store.Int32(ptr + 204 + 32 + 8, 0);
            A.store.Bool(ptr + 204 + 32 + 28, false);
            A.store.Int32(ptr + 204 + 32 + 12, 0);
            A.store.Bool(ptr + 204 + 32 + 29, false);
            A.store.Int32(ptr + 204 + 32 + 16, 0);
            A.store.Bool(ptr + 204 + 32 + 30, false);
            A.store.Int32(ptr + 204 + 32 + 20, 0);
            A.store.Bool(ptr + 204 + 32 + 31, false);
            A.store.Bool(ptr + 204 + 32 + 24, false);
            A.store.Bool(ptr + 204 + 32 + 32, false);
            A.store.Bool(ptr + 204 + 32 + 25, false);
          } else {
            A.store.Bool(ptr + 204 + 32 + 33, true);
            A.store.Bool(ptr + 204 + 32 + 26, "Active" in x["SOCKS"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 204 + 32 + 0,
              x["SOCKS"]["Port"]["Active"] === undefined ? 0 : (x["SOCKS"]["Port"]["Active"] as number)
            );
            A.store.Ref(ptr + 204 + 32 + 4, x["SOCKS"]["Port"]["Effective"]);
            A.store.Bool(ptr + 204 + 32 + 27, "UserPolicy" in x["SOCKS"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 204 + 32 + 8,
              x["SOCKS"]["Port"]["UserPolicy"] === undefined ? 0 : (x["SOCKS"]["Port"]["UserPolicy"] as number)
            );
            A.store.Bool(ptr + 204 + 32 + 28, "DevicePolicy" in x["SOCKS"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 204 + 32 + 12,
              x["SOCKS"]["Port"]["DevicePolicy"] === undefined ? 0 : (x["SOCKS"]["Port"]["DevicePolicy"] as number)
            );
            A.store.Bool(ptr + 204 + 32 + 29, "UserSetting" in x["SOCKS"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 204 + 32 + 16,
              x["SOCKS"]["Port"]["UserSetting"] === undefined ? 0 : (x["SOCKS"]["Port"]["UserSetting"] as number)
            );
            A.store.Bool(ptr + 204 + 32 + 30, "SharedSetting" in x["SOCKS"]["Port"] ? true : false);
            A.store.Int32(
              ptr + 204 + 32 + 20,
              x["SOCKS"]["Port"]["SharedSetting"] === undefined ? 0 : (x["SOCKS"]["Port"]["SharedSetting"] as number)
            );
            A.store.Bool(ptr + 204 + 32 + 31, "UserEditable" in x["SOCKS"]["Port"] ? true : false);
            A.store.Bool(ptr + 204 + 32 + 24, x["SOCKS"]["Port"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 204 + 32 + 32, "DeviceEditable" in x["SOCKS"]["Port"] ? true : false);
            A.store.Bool(ptr + 204 + 32 + 25, x["SOCKS"]["Port"]["DeviceEditable"] ? true : false);
          }
        }
      }
    },
    "load_ManagedManualProxySettings": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 66)) {
        x["HTTPProxy"] = {};
        if (A.load.Bool(ptr + 0 + 0 + 28)) {
          x["HTTPProxy"]["Host"] = {};
          x["HTTPProxy"]["Host"]["Active"] = A.load.Ref(ptr + 0 + 0 + 0, undefined);
          x["HTTPProxy"]["Host"]["Effective"] = A.load.Ref(ptr + 0 + 0 + 4, undefined);
          x["HTTPProxy"]["Host"]["UserPolicy"] = A.load.Ref(ptr + 0 + 0 + 8, undefined);
          x["HTTPProxy"]["Host"]["DevicePolicy"] = A.load.Ref(ptr + 0 + 0 + 12, undefined);
          x["HTTPProxy"]["Host"]["UserSetting"] = A.load.Ref(ptr + 0 + 0 + 16, undefined);
          x["HTTPProxy"]["Host"]["SharedSetting"] = A.load.Ref(ptr + 0 + 0 + 20, undefined);
          if (A.load.Bool(ptr + 0 + 0 + 26)) {
            x["HTTPProxy"]["Host"]["UserEditable"] = A.load.Bool(ptr + 0 + 0 + 24);
          } else {
            delete x["HTTPProxy"]["Host"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 0 + 0 + 27)) {
            x["HTTPProxy"]["Host"]["DeviceEditable"] = A.load.Bool(ptr + 0 + 0 + 25);
          } else {
            delete x["HTTPProxy"]["Host"]["DeviceEditable"];
          }
        } else {
          delete x["HTTPProxy"]["Host"];
        }
        if (A.load.Bool(ptr + 0 + 32 + 33)) {
          x["HTTPProxy"]["Port"] = {};
          if (A.load.Bool(ptr + 0 + 32 + 26)) {
            x["HTTPProxy"]["Port"]["Active"] = A.load.Int32(ptr + 0 + 32 + 0);
          } else {
            delete x["HTTPProxy"]["Port"]["Active"];
          }
          x["HTTPProxy"]["Port"]["Effective"] = A.load.Ref(ptr + 0 + 32 + 4, undefined);
          if (A.load.Bool(ptr + 0 + 32 + 27)) {
            x["HTTPProxy"]["Port"]["UserPolicy"] = A.load.Int32(ptr + 0 + 32 + 8);
          } else {
            delete x["HTTPProxy"]["Port"]["UserPolicy"];
          }
          if (A.load.Bool(ptr + 0 + 32 + 28)) {
            x["HTTPProxy"]["Port"]["DevicePolicy"] = A.load.Int32(ptr + 0 + 32 + 12);
          } else {
            delete x["HTTPProxy"]["Port"]["DevicePolicy"];
          }
          if (A.load.Bool(ptr + 0 + 32 + 29)) {
            x["HTTPProxy"]["Port"]["UserSetting"] = A.load.Int32(ptr + 0 + 32 + 16);
          } else {
            delete x["HTTPProxy"]["Port"]["UserSetting"];
          }
          if (A.load.Bool(ptr + 0 + 32 + 30)) {
            x["HTTPProxy"]["Port"]["SharedSetting"] = A.load.Int32(ptr + 0 + 32 + 20);
          } else {
            delete x["HTTPProxy"]["Port"]["SharedSetting"];
          }
          if (A.load.Bool(ptr + 0 + 32 + 31)) {
            x["HTTPProxy"]["Port"]["UserEditable"] = A.load.Bool(ptr + 0 + 32 + 24);
          } else {
            delete x["HTTPProxy"]["Port"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 0 + 32 + 32)) {
            x["HTTPProxy"]["Port"]["DeviceEditable"] = A.load.Bool(ptr + 0 + 32 + 25);
          } else {
            delete x["HTTPProxy"]["Port"]["DeviceEditable"];
          }
        } else {
          delete x["HTTPProxy"]["Port"];
        }
      } else {
        delete x["HTTPProxy"];
      }
      if (A.load.Bool(ptr + 68 + 66)) {
        x["SecureHTTPProxy"] = {};
        if (A.load.Bool(ptr + 68 + 0 + 28)) {
          x["SecureHTTPProxy"]["Host"] = {};
          x["SecureHTTPProxy"]["Host"]["Active"] = A.load.Ref(ptr + 68 + 0 + 0, undefined);
          x["SecureHTTPProxy"]["Host"]["Effective"] = A.load.Ref(ptr + 68 + 0 + 4, undefined);
          x["SecureHTTPProxy"]["Host"]["UserPolicy"] = A.load.Ref(ptr + 68 + 0 + 8, undefined);
          x["SecureHTTPProxy"]["Host"]["DevicePolicy"] = A.load.Ref(ptr + 68 + 0 + 12, undefined);
          x["SecureHTTPProxy"]["Host"]["UserSetting"] = A.load.Ref(ptr + 68 + 0 + 16, undefined);
          x["SecureHTTPProxy"]["Host"]["SharedSetting"] = A.load.Ref(ptr + 68 + 0 + 20, undefined);
          if (A.load.Bool(ptr + 68 + 0 + 26)) {
            x["SecureHTTPProxy"]["Host"]["UserEditable"] = A.load.Bool(ptr + 68 + 0 + 24);
          } else {
            delete x["SecureHTTPProxy"]["Host"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 68 + 0 + 27)) {
            x["SecureHTTPProxy"]["Host"]["DeviceEditable"] = A.load.Bool(ptr + 68 + 0 + 25);
          } else {
            delete x["SecureHTTPProxy"]["Host"]["DeviceEditable"];
          }
        } else {
          delete x["SecureHTTPProxy"]["Host"];
        }
        if (A.load.Bool(ptr + 68 + 32 + 33)) {
          x["SecureHTTPProxy"]["Port"] = {};
          if (A.load.Bool(ptr + 68 + 32 + 26)) {
            x["SecureHTTPProxy"]["Port"]["Active"] = A.load.Int32(ptr + 68 + 32 + 0);
          } else {
            delete x["SecureHTTPProxy"]["Port"]["Active"];
          }
          x["SecureHTTPProxy"]["Port"]["Effective"] = A.load.Ref(ptr + 68 + 32 + 4, undefined);
          if (A.load.Bool(ptr + 68 + 32 + 27)) {
            x["SecureHTTPProxy"]["Port"]["UserPolicy"] = A.load.Int32(ptr + 68 + 32 + 8);
          } else {
            delete x["SecureHTTPProxy"]["Port"]["UserPolicy"];
          }
          if (A.load.Bool(ptr + 68 + 32 + 28)) {
            x["SecureHTTPProxy"]["Port"]["DevicePolicy"] = A.load.Int32(ptr + 68 + 32 + 12);
          } else {
            delete x["SecureHTTPProxy"]["Port"]["DevicePolicy"];
          }
          if (A.load.Bool(ptr + 68 + 32 + 29)) {
            x["SecureHTTPProxy"]["Port"]["UserSetting"] = A.load.Int32(ptr + 68 + 32 + 16);
          } else {
            delete x["SecureHTTPProxy"]["Port"]["UserSetting"];
          }
          if (A.load.Bool(ptr + 68 + 32 + 30)) {
            x["SecureHTTPProxy"]["Port"]["SharedSetting"] = A.load.Int32(ptr + 68 + 32 + 20);
          } else {
            delete x["SecureHTTPProxy"]["Port"]["SharedSetting"];
          }
          if (A.load.Bool(ptr + 68 + 32 + 31)) {
            x["SecureHTTPProxy"]["Port"]["UserEditable"] = A.load.Bool(ptr + 68 + 32 + 24);
          } else {
            delete x["SecureHTTPProxy"]["Port"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 68 + 32 + 32)) {
            x["SecureHTTPProxy"]["Port"]["DeviceEditable"] = A.load.Bool(ptr + 68 + 32 + 25);
          } else {
            delete x["SecureHTTPProxy"]["Port"]["DeviceEditable"];
          }
        } else {
          delete x["SecureHTTPProxy"]["Port"];
        }
      } else {
        delete x["SecureHTTPProxy"];
      }
      if (A.load.Bool(ptr + 136 + 66)) {
        x["FTPProxy"] = {};
        if (A.load.Bool(ptr + 136 + 0 + 28)) {
          x["FTPProxy"]["Host"] = {};
          x["FTPProxy"]["Host"]["Active"] = A.load.Ref(ptr + 136 + 0 + 0, undefined);
          x["FTPProxy"]["Host"]["Effective"] = A.load.Ref(ptr + 136 + 0 + 4, undefined);
          x["FTPProxy"]["Host"]["UserPolicy"] = A.load.Ref(ptr + 136 + 0 + 8, undefined);
          x["FTPProxy"]["Host"]["DevicePolicy"] = A.load.Ref(ptr + 136 + 0 + 12, undefined);
          x["FTPProxy"]["Host"]["UserSetting"] = A.load.Ref(ptr + 136 + 0 + 16, undefined);
          x["FTPProxy"]["Host"]["SharedSetting"] = A.load.Ref(ptr + 136 + 0 + 20, undefined);
          if (A.load.Bool(ptr + 136 + 0 + 26)) {
            x["FTPProxy"]["Host"]["UserEditable"] = A.load.Bool(ptr + 136 + 0 + 24);
          } else {
            delete x["FTPProxy"]["Host"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 136 + 0 + 27)) {
            x["FTPProxy"]["Host"]["DeviceEditable"] = A.load.Bool(ptr + 136 + 0 + 25);
          } else {
            delete x["FTPProxy"]["Host"]["DeviceEditable"];
          }
        } else {
          delete x["FTPProxy"]["Host"];
        }
        if (A.load.Bool(ptr + 136 + 32 + 33)) {
          x["FTPProxy"]["Port"] = {};
          if (A.load.Bool(ptr + 136 + 32 + 26)) {
            x["FTPProxy"]["Port"]["Active"] = A.load.Int32(ptr + 136 + 32 + 0);
          } else {
            delete x["FTPProxy"]["Port"]["Active"];
          }
          x["FTPProxy"]["Port"]["Effective"] = A.load.Ref(ptr + 136 + 32 + 4, undefined);
          if (A.load.Bool(ptr + 136 + 32 + 27)) {
            x["FTPProxy"]["Port"]["UserPolicy"] = A.load.Int32(ptr + 136 + 32 + 8);
          } else {
            delete x["FTPProxy"]["Port"]["UserPolicy"];
          }
          if (A.load.Bool(ptr + 136 + 32 + 28)) {
            x["FTPProxy"]["Port"]["DevicePolicy"] = A.load.Int32(ptr + 136 + 32 + 12);
          } else {
            delete x["FTPProxy"]["Port"]["DevicePolicy"];
          }
          if (A.load.Bool(ptr + 136 + 32 + 29)) {
            x["FTPProxy"]["Port"]["UserSetting"] = A.load.Int32(ptr + 136 + 32 + 16);
          } else {
            delete x["FTPProxy"]["Port"]["UserSetting"];
          }
          if (A.load.Bool(ptr + 136 + 32 + 30)) {
            x["FTPProxy"]["Port"]["SharedSetting"] = A.load.Int32(ptr + 136 + 32 + 20);
          } else {
            delete x["FTPProxy"]["Port"]["SharedSetting"];
          }
          if (A.load.Bool(ptr + 136 + 32 + 31)) {
            x["FTPProxy"]["Port"]["UserEditable"] = A.load.Bool(ptr + 136 + 32 + 24);
          } else {
            delete x["FTPProxy"]["Port"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 136 + 32 + 32)) {
            x["FTPProxy"]["Port"]["DeviceEditable"] = A.load.Bool(ptr + 136 + 32 + 25);
          } else {
            delete x["FTPProxy"]["Port"]["DeviceEditable"];
          }
        } else {
          delete x["FTPProxy"]["Port"];
        }
      } else {
        delete x["FTPProxy"];
      }
      if (A.load.Bool(ptr + 204 + 66)) {
        x["SOCKS"] = {};
        if (A.load.Bool(ptr + 204 + 0 + 28)) {
          x["SOCKS"]["Host"] = {};
          x["SOCKS"]["Host"]["Active"] = A.load.Ref(ptr + 204 + 0 + 0, undefined);
          x["SOCKS"]["Host"]["Effective"] = A.load.Ref(ptr + 204 + 0 + 4, undefined);
          x["SOCKS"]["Host"]["UserPolicy"] = A.load.Ref(ptr + 204 + 0 + 8, undefined);
          x["SOCKS"]["Host"]["DevicePolicy"] = A.load.Ref(ptr + 204 + 0 + 12, undefined);
          x["SOCKS"]["Host"]["UserSetting"] = A.load.Ref(ptr + 204 + 0 + 16, undefined);
          x["SOCKS"]["Host"]["SharedSetting"] = A.load.Ref(ptr + 204 + 0 + 20, undefined);
          if (A.load.Bool(ptr + 204 + 0 + 26)) {
            x["SOCKS"]["Host"]["UserEditable"] = A.load.Bool(ptr + 204 + 0 + 24);
          } else {
            delete x["SOCKS"]["Host"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 204 + 0 + 27)) {
            x["SOCKS"]["Host"]["DeviceEditable"] = A.load.Bool(ptr + 204 + 0 + 25);
          } else {
            delete x["SOCKS"]["Host"]["DeviceEditable"];
          }
        } else {
          delete x["SOCKS"]["Host"];
        }
        if (A.load.Bool(ptr + 204 + 32 + 33)) {
          x["SOCKS"]["Port"] = {};
          if (A.load.Bool(ptr + 204 + 32 + 26)) {
            x["SOCKS"]["Port"]["Active"] = A.load.Int32(ptr + 204 + 32 + 0);
          } else {
            delete x["SOCKS"]["Port"]["Active"];
          }
          x["SOCKS"]["Port"]["Effective"] = A.load.Ref(ptr + 204 + 32 + 4, undefined);
          if (A.load.Bool(ptr + 204 + 32 + 27)) {
            x["SOCKS"]["Port"]["UserPolicy"] = A.load.Int32(ptr + 204 + 32 + 8);
          } else {
            delete x["SOCKS"]["Port"]["UserPolicy"];
          }
          if (A.load.Bool(ptr + 204 + 32 + 28)) {
            x["SOCKS"]["Port"]["DevicePolicy"] = A.load.Int32(ptr + 204 + 32 + 12);
          } else {
            delete x["SOCKS"]["Port"]["DevicePolicy"];
          }
          if (A.load.Bool(ptr + 204 + 32 + 29)) {
            x["SOCKS"]["Port"]["UserSetting"] = A.load.Int32(ptr + 204 + 32 + 16);
          } else {
            delete x["SOCKS"]["Port"]["UserSetting"];
          }
          if (A.load.Bool(ptr + 204 + 32 + 30)) {
            x["SOCKS"]["Port"]["SharedSetting"] = A.load.Int32(ptr + 204 + 32 + 20);
          } else {
            delete x["SOCKS"]["Port"]["SharedSetting"];
          }
          if (A.load.Bool(ptr + 204 + 32 + 31)) {
            x["SOCKS"]["Port"]["UserEditable"] = A.load.Bool(ptr + 204 + 32 + 24);
          } else {
            delete x["SOCKS"]["Port"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 204 + 32 + 32)) {
            x["SOCKS"]["Port"]["DeviceEditable"] = A.load.Bool(ptr + 204 + 32 + 25);
          } else {
            delete x["SOCKS"]["Port"]["DeviceEditable"];
          }
        } else {
          delete x["SOCKS"]["Port"];
        }
      } else {
        delete x["SOCKS"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManagedDOMStringList": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 27, false);
        A.store.Bool(ptr + 25, false);
      } else {
        A.store.Bool(ptr + 28, true);
        A.store.Ref(ptr + 0, x["Active"]);
        A.store.Ref(ptr + 4, x["Effective"]);
        A.store.Ref(ptr + 8, x["UserPolicy"]);
        A.store.Ref(ptr + 12, x["DevicePolicy"]);
        A.store.Ref(ptr + 16, x["UserSetting"]);
        A.store.Ref(ptr + 20, x["SharedSetting"]);
        A.store.Bool(ptr + 26, "UserEditable" in x ? true : false);
        A.store.Bool(ptr + 24, x["UserEditable"] ? true : false);
        A.store.Bool(ptr + 27, "DeviceEditable" in x ? true : false);
        A.store.Bool(ptr + 25, x["DeviceEditable"] ? true : false);
      }
    },
    "load_ManagedDOMStringList": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["Active"] = A.load.Ref(ptr + 0, undefined);
      x["Effective"] = A.load.Ref(ptr + 4, undefined);
      x["UserPolicy"] = A.load.Ref(ptr + 8, undefined);
      x["DevicePolicy"] = A.load.Ref(ptr + 12, undefined);
      x["UserSetting"] = A.load.Ref(ptr + 16, undefined);
      x["SharedSetting"] = A.load.Ref(ptr + 20, undefined);
      if (A.load.Bool(ptr + 26)) {
        x["UserEditable"] = A.load.Bool(ptr + 24);
      } else {
        delete x["UserEditable"];
      }
      if (A.load.Bool(ptr + 27)) {
        x["DeviceEditable"] = A.load.Bool(ptr + 25);
      } else {
        delete x["DeviceEditable"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManagedProxySettings": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 365, false);

        A.store.Bool(ptr + 0 + 28, false);
        A.store.Enum(ptr + 0 + 0, -1);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Enum(ptr + 0 + 8, -1);
        A.store.Enum(ptr + 0 + 12, -1);
        A.store.Enum(ptr + 0 + 16, -1);
        A.store.Enum(ptr + 0 + 20, -1);
        A.store.Bool(ptr + 0 + 26, false);
        A.store.Bool(ptr + 0 + 24, false);
        A.store.Bool(ptr + 0 + 27, false);
        A.store.Bool(ptr + 0 + 25, false);

        A.store.Bool(ptr + 32 + 271, false);

        A.store.Bool(ptr + 32 + 0 + 66, false);

        A.store.Bool(ptr + 32 + 0 + 0 + 28, false);
        A.store.Ref(ptr + 32 + 0 + 0 + 0, undefined);
        A.store.Ref(ptr + 32 + 0 + 0 + 4, undefined);
        A.store.Ref(ptr + 32 + 0 + 0 + 8, undefined);
        A.store.Ref(ptr + 32 + 0 + 0 + 12, undefined);
        A.store.Ref(ptr + 32 + 0 + 0 + 16, undefined);
        A.store.Ref(ptr + 32 + 0 + 0 + 20, undefined);
        A.store.Bool(ptr + 32 + 0 + 0 + 26, false);
        A.store.Bool(ptr + 32 + 0 + 0 + 24, false);
        A.store.Bool(ptr + 32 + 0 + 0 + 27, false);
        A.store.Bool(ptr + 32 + 0 + 0 + 25, false);

        A.store.Bool(ptr + 32 + 0 + 32 + 33, false);
        A.store.Bool(ptr + 32 + 0 + 32 + 26, false);
        A.store.Int32(ptr + 32 + 0 + 32 + 0, 0);
        A.store.Ref(ptr + 32 + 0 + 32 + 4, undefined);
        A.store.Bool(ptr + 32 + 0 + 32 + 27, false);
        A.store.Int32(ptr + 32 + 0 + 32 + 8, 0);
        A.store.Bool(ptr + 32 + 0 + 32 + 28, false);
        A.store.Int32(ptr + 32 + 0 + 32 + 12, 0);
        A.store.Bool(ptr + 32 + 0 + 32 + 29, false);
        A.store.Int32(ptr + 32 + 0 + 32 + 16, 0);
        A.store.Bool(ptr + 32 + 0 + 32 + 30, false);
        A.store.Int32(ptr + 32 + 0 + 32 + 20, 0);
        A.store.Bool(ptr + 32 + 0 + 32 + 31, false);
        A.store.Bool(ptr + 32 + 0 + 32 + 24, false);
        A.store.Bool(ptr + 32 + 0 + 32 + 32, false);
        A.store.Bool(ptr + 32 + 0 + 32 + 25, false);

        A.store.Bool(ptr + 32 + 68 + 66, false);

        A.store.Bool(ptr + 32 + 68 + 0 + 28, false);
        A.store.Ref(ptr + 32 + 68 + 0 + 0, undefined);
        A.store.Ref(ptr + 32 + 68 + 0 + 4, undefined);
        A.store.Ref(ptr + 32 + 68 + 0 + 8, undefined);
        A.store.Ref(ptr + 32 + 68 + 0 + 12, undefined);
        A.store.Ref(ptr + 32 + 68 + 0 + 16, undefined);
        A.store.Ref(ptr + 32 + 68 + 0 + 20, undefined);
        A.store.Bool(ptr + 32 + 68 + 0 + 26, false);
        A.store.Bool(ptr + 32 + 68 + 0 + 24, false);
        A.store.Bool(ptr + 32 + 68 + 0 + 27, false);
        A.store.Bool(ptr + 32 + 68 + 0 + 25, false);

        A.store.Bool(ptr + 32 + 68 + 32 + 33, false);
        A.store.Bool(ptr + 32 + 68 + 32 + 26, false);
        A.store.Int32(ptr + 32 + 68 + 32 + 0, 0);
        A.store.Ref(ptr + 32 + 68 + 32 + 4, undefined);
        A.store.Bool(ptr + 32 + 68 + 32 + 27, false);
        A.store.Int32(ptr + 32 + 68 + 32 + 8, 0);
        A.store.Bool(ptr + 32 + 68 + 32 + 28, false);
        A.store.Int32(ptr + 32 + 68 + 32 + 12, 0);
        A.store.Bool(ptr + 32 + 68 + 32 + 29, false);
        A.store.Int32(ptr + 32 + 68 + 32 + 16, 0);
        A.store.Bool(ptr + 32 + 68 + 32 + 30, false);
        A.store.Int32(ptr + 32 + 68 + 32 + 20, 0);
        A.store.Bool(ptr + 32 + 68 + 32 + 31, false);
        A.store.Bool(ptr + 32 + 68 + 32 + 24, false);
        A.store.Bool(ptr + 32 + 68 + 32 + 32, false);
        A.store.Bool(ptr + 32 + 68 + 32 + 25, false);

        A.store.Bool(ptr + 32 + 136 + 66, false);

        A.store.Bool(ptr + 32 + 136 + 0 + 28, false);
        A.store.Ref(ptr + 32 + 136 + 0 + 0, undefined);
        A.store.Ref(ptr + 32 + 136 + 0 + 4, undefined);
        A.store.Ref(ptr + 32 + 136 + 0 + 8, undefined);
        A.store.Ref(ptr + 32 + 136 + 0 + 12, undefined);
        A.store.Ref(ptr + 32 + 136 + 0 + 16, undefined);
        A.store.Ref(ptr + 32 + 136 + 0 + 20, undefined);
        A.store.Bool(ptr + 32 + 136 + 0 + 26, false);
        A.store.Bool(ptr + 32 + 136 + 0 + 24, false);
        A.store.Bool(ptr + 32 + 136 + 0 + 27, false);
        A.store.Bool(ptr + 32 + 136 + 0 + 25, false);

        A.store.Bool(ptr + 32 + 136 + 32 + 33, false);
        A.store.Bool(ptr + 32 + 136 + 32 + 26, false);
        A.store.Int32(ptr + 32 + 136 + 32 + 0, 0);
        A.store.Ref(ptr + 32 + 136 + 32 + 4, undefined);
        A.store.Bool(ptr + 32 + 136 + 32 + 27, false);
        A.store.Int32(ptr + 32 + 136 + 32 + 8, 0);
        A.store.Bool(ptr + 32 + 136 + 32 + 28, false);
        A.store.Int32(ptr + 32 + 136 + 32 + 12, 0);
        A.store.Bool(ptr + 32 + 136 + 32 + 29, false);
        A.store.Int32(ptr + 32 + 136 + 32 + 16, 0);
        A.store.Bool(ptr + 32 + 136 + 32 + 30, false);
        A.store.Int32(ptr + 32 + 136 + 32 + 20, 0);
        A.store.Bool(ptr + 32 + 136 + 32 + 31, false);
        A.store.Bool(ptr + 32 + 136 + 32 + 24, false);
        A.store.Bool(ptr + 32 + 136 + 32 + 32, false);
        A.store.Bool(ptr + 32 + 136 + 32 + 25, false);

        A.store.Bool(ptr + 32 + 204 + 66, false);

        A.store.Bool(ptr + 32 + 204 + 0 + 28, false);
        A.store.Ref(ptr + 32 + 204 + 0 + 0, undefined);
        A.store.Ref(ptr + 32 + 204 + 0 + 4, undefined);
        A.store.Ref(ptr + 32 + 204 + 0 + 8, undefined);
        A.store.Ref(ptr + 32 + 204 + 0 + 12, undefined);
        A.store.Ref(ptr + 32 + 204 + 0 + 16, undefined);
        A.store.Ref(ptr + 32 + 204 + 0 + 20, undefined);
        A.store.Bool(ptr + 32 + 204 + 0 + 26, false);
        A.store.Bool(ptr + 32 + 204 + 0 + 24, false);
        A.store.Bool(ptr + 32 + 204 + 0 + 27, false);
        A.store.Bool(ptr + 32 + 204 + 0 + 25, false);

        A.store.Bool(ptr + 32 + 204 + 32 + 33, false);
        A.store.Bool(ptr + 32 + 204 + 32 + 26, false);
        A.store.Int32(ptr + 32 + 204 + 32 + 0, 0);
        A.store.Ref(ptr + 32 + 204 + 32 + 4, undefined);
        A.store.Bool(ptr + 32 + 204 + 32 + 27, false);
        A.store.Int32(ptr + 32 + 204 + 32 + 8, 0);
        A.store.Bool(ptr + 32 + 204 + 32 + 28, false);
        A.store.Int32(ptr + 32 + 204 + 32 + 12, 0);
        A.store.Bool(ptr + 32 + 204 + 32 + 29, false);
        A.store.Int32(ptr + 32 + 204 + 32 + 16, 0);
        A.store.Bool(ptr + 32 + 204 + 32 + 30, false);
        A.store.Int32(ptr + 32 + 204 + 32 + 20, 0);
        A.store.Bool(ptr + 32 + 204 + 32 + 31, false);
        A.store.Bool(ptr + 32 + 204 + 32 + 24, false);
        A.store.Bool(ptr + 32 + 204 + 32 + 32, false);
        A.store.Bool(ptr + 32 + 204 + 32 + 25, false);

        A.store.Bool(ptr + 304 + 28, false);
        A.store.Ref(ptr + 304 + 0, undefined);
        A.store.Ref(ptr + 304 + 4, undefined);
        A.store.Ref(ptr + 304 + 8, undefined);
        A.store.Ref(ptr + 304 + 12, undefined);
        A.store.Ref(ptr + 304 + 16, undefined);
        A.store.Ref(ptr + 304 + 20, undefined);
        A.store.Bool(ptr + 304 + 26, false);
        A.store.Bool(ptr + 304 + 24, false);
        A.store.Bool(ptr + 304 + 27, false);
        A.store.Bool(ptr + 304 + 25, false);

        A.store.Bool(ptr + 336 + 28, false);
        A.store.Ref(ptr + 336 + 0, undefined);
        A.store.Ref(ptr + 336 + 4, undefined);
        A.store.Ref(ptr + 336 + 8, undefined);
        A.store.Ref(ptr + 336 + 12, undefined);
        A.store.Ref(ptr + 336 + 16, undefined);
        A.store.Ref(ptr + 336 + 20, undefined);
        A.store.Bool(ptr + 336 + 26, false);
        A.store.Bool(ptr + 336 + 24, false);
        A.store.Bool(ptr + 336 + 27, false);
        A.store.Bool(ptr + 336 + 25, false);
      } else {
        A.store.Bool(ptr + 365, true);

        if (typeof x["Type"] === "undefined") {
          A.store.Bool(ptr + 0 + 28, false);
          A.store.Enum(ptr + 0 + 0, -1);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Enum(ptr + 0 + 8, -1);
          A.store.Enum(ptr + 0 + 12, -1);
          A.store.Enum(ptr + 0 + 16, -1);
          A.store.Enum(ptr + 0 + 20, -1);
          A.store.Bool(ptr + 0 + 26, false);
          A.store.Bool(ptr + 0 + 24, false);
          A.store.Bool(ptr + 0 + 27, false);
          A.store.Bool(ptr + 0 + 25, false);
        } else {
          A.store.Bool(ptr + 0 + 28, true);
          A.store.Enum(ptr + 0 + 0, ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["Type"]["Active"] as string));
          A.store.Ref(ptr + 0 + 4, x["Type"]["Effective"]);
          A.store.Enum(ptr + 0 + 8, ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["Type"]["UserPolicy"] as string));
          A.store.Enum(ptr + 0 + 12, ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["Type"]["DevicePolicy"] as string));
          A.store.Enum(ptr + 0 + 16, ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["Type"]["UserSetting"] as string));
          A.store.Enum(ptr + 0 + 20, ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["Type"]["SharedSetting"] as string));
          A.store.Bool(ptr + 0 + 26, "UserEditable" in x["Type"] ? true : false);
          A.store.Bool(ptr + 0 + 24, x["Type"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 0 + 27, "DeviceEditable" in x["Type"] ? true : false);
          A.store.Bool(ptr + 0 + 25, x["Type"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["Manual"] === "undefined") {
          A.store.Bool(ptr + 32 + 271, false);

          A.store.Bool(ptr + 32 + 0 + 66, false);

          A.store.Bool(ptr + 32 + 0 + 0 + 28, false);
          A.store.Ref(ptr + 32 + 0 + 0 + 0, undefined);
          A.store.Ref(ptr + 32 + 0 + 0 + 4, undefined);
          A.store.Ref(ptr + 32 + 0 + 0 + 8, undefined);
          A.store.Ref(ptr + 32 + 0 + 0 + 12, undefined);
          A.store.Ref(ptr + 32 + 0 + 0 + 16, undefined);
          A.store.Ref(ptr + 32 + 0 + 0 + 20, undefined);
          A.store.Bool(ptr + 32 + 0 + 0 + 26, false);
          A.store.Bool(ptr + 32 + 0 + 0 + 24, false);
          A.store.Bool(ptr + 32 + 0 + 0 + 27, false);
          A.store.Bool(ptr + 32 + 0 + 0 + 25, false);

          A.store.Bool(ptr + 32 + 0 + 32 + 33, false);
          A.store.Bool(ptr + 32 + 0 + 32 + 26, false);
          A.store.Int32(ptr + 32 + 0 + 32 + 0, 0);
          A.store.Ref(ptr + 32 + 0 + 32 + 4, undefined);
          A.store.Bool(ptr + 32 + 0 + 32 + 27, false);
          A.store.Int32(ptr + 32 + 0 + 32 + 8, 0);
          A.store.Bool(ptr + 32 + 0 + 32 + 28, false);
          A.store.Int32(ptr + 32 + 0 + 32 + 12, 0);
          A.store.Bool(ptr + 32 + 0 + 32 + 29, false);
          A.store.Int32(ptr + 32 + 0 + 32 + 16, 0);
          A.store.Bool(ptr + 32 + 0 + 32 + 30, false);
          A.store.Int32(ptr + 32 + 0 + 32 + 20, 0);
          A.store.Bool(ptr + 32 + 0 + 32 + 31, false);
          A.store.Bool(ptr + 32 + 0 + 32 + 24, false);
          A.store.Bool(ptr + 32 + 0 + 32 + 32, false);
          A.store.Bool(ptr + 32 + 0 + 32 + 25, false);

          A.store.Bool(ptr + 32 + 68 + 66, false);

          A.store.Bool(ptr + 32 + 68 + 0 + 28, false);
          A.store.Ref(ptr + 32 + 68 + 0 + 0, undefined);
          A.store.Ref(ptr + 32 + 68 + 0 + 4, undefined);
          A.store.Ref(ptr + 32 + 68 + 0 + 8, undefined);
          A.store.Ref(ptr + 32 + 68 + 0 + 12, undefined);
          A.store.Ref(ptr + 32 + 68 + 0 + 16, undefined);
          A.store.Ref(ptr + 32 + 68 + 0 + 20, undefined);
          A.store.Bool(ptr + 32 + 68 + 0 + 26, false);
          A.store.Bool(ptr + 32 + 68 + 0 + 24, false);
          A.store.Bool(ptr + 32 + 68 + 0 + 27, false);
          A.store.Bool(ptr + 32 + 68 + 0 + 25, false);

          A.store.Bool(ptr + 32 + 68 + 32 + 33, false);
          A.store.Bool(ptr + 32 + 68 + 32 + 26, false);
          A.store.Int32(ptr + 32 + 68 + 32 + 0, 0);
          A.store.Ref(ptr + 32 + 68 + 32 + 4, undefined);
          A.store.Bool(ptr + 32 + 68 + 32 + 27, false);
          A.store.Int32(ptr + 32 + 68 + 32 + 8, 0);
          A.store.Bool(ptr + 32 + 68 + 32 + 28, false);
          A.store.Int32(ptr + 32 + 68 + 32 + 12, 0);
          A.store.Bool(ptr + 32 + 68 + 32 + 29, false);
          A.store.Int32(ptr + 32 + 68 + 32 + 16, 0);
          A.store.Bool(ptr + 32 + 68 + 32 + 30, false);
          A.store.Int32(ptr + 32 + 68 + 32 + 20, 0);
          A.store.Bool(ptr + 32 + 68 + 32 + 31, false);
          A.store.Bool(ptr + 32 + 68 + 32 + 24, false);
          A.store.Bool(ptr + 32 + 68 + 32 + 32, false);
          A.store.Bool(ptr + 32 + 68 + 32 + 25, false);

          A.store.Bool(ptr + 32 + 136 + 66, false);

          A.store.Bool(ptr + 32 + 136 + 0 + 28, false);
          A.store.Ref(ptr + 32 + 136 + 0 + 0, undefined);
          A.store.Ref(ptr + 32 + 136 + 0 + 4, undefined);
          A.store.Ref(ptr + 32 + 136 + 0 + 8, undefined);
          A.store.Ref(ptr + 32 + 136 + 0 + 12, undefined);
          A.store.Ref(ptr + 32 + 136 + 0 + 16, undefined);
          A.store.Ref(ptr + 32 + 136 + 0 + 20, undefined);
          A.store.Bool(ptr + 32 + 136 + 0 + 26, false);
          A.store.Bool(ptr + 32 + 136 + 0 + 24, false);
          A.store.Bool(ptr + 32 + 136 + 0 + 27, false);
          A.store.Bool(ptr + 32 + 136 + 0 + 25, false);

          A.store.Bool(ptr + 32 + 136 + 32 + 33, false);
          A.store.Bool(ptr + 32 + 136 + 32 + 26, false);
          A.store.Int32(ptr + 32 + 136 + 32 + 0, 0);
          A.store.Ref(ptr + 32 + 136 + 32 + 4, undefined);
          A.store.Bool(ptr + 32 + 136 + 32 + 27, false);
          A.store.Int32(ptr + 32 + 136 + 32 + 8, 0);
          A.store.Bool(ptr + 32 + 136 + 32 + 28, false);
          A.store.Int32(ptr + 32 + 136 + 32 + 12, 0);
          A.store.Bool(ptr + 32 + 136 + 32 + 29, false);
          A.store.Int32(ptr + 32 + 136 + 32 + 16, 0);
          A.store.Bool(ptr + 32 + 136 + 32 + 30, false);
          A.store.Int32(ptr + 32 + 136 + 32 + 20, 0);
          A.store.Bool(ptr + 32 + 136 + 32 + 31, false);
          A.store.Bool(ptr + 32 + 136 + 32 + 24, false);
          A.store.Bool(ptr + 32 + 136 + 32 + 32, false);
          A.store.Bool(ptr + 32 + 136 + 32 + 25, false);

          A.store.Bool(ptr + 32 + 204 + 66, false);

          A.store.Bool(ptr + 32 + 204 + 0 + 28, false);
          A.store.Ref(ptr + 32 + 204 + 0 + 0, undefined);
          A.store.Ref(ptr + 32 + 204 + 0 + 4, undefined);
          A.store.Ref(ptr + 32 + 204 + 0 + 8, undefined);
          A.store.Ref(ptr + 32 + 204 + 0 + 12, undefined);
          A.store.Ref(ptr + 32 + 204 + 0 + 16, undefined);
          A.store.Ref(ptr + 32 + 204 + 0 + 20, undefined);
          A.store.Bool(ptr + 32 + 204 + 0 + 26, false);
          A.store.Bool(ptr + 32 + 204 + 0 + 24, false);
          A.store.Bool(ptr + 32 + 204 + 0 + 27, false);
          A.store.Bool(ptr + 32 + 204 + 0 + 25, false);

          A.store.Bool(ptr + 32 + 204 + 32 + 33, false);
          A.store.Bool(ptr + 32 + 204 + 32 + 26, false);
          A.store.Int32(ptr + 32 + 204 + 32 + 0, 0);
          A.store.Ref(ptr + 32 + 204 + 32 + 4, undefined);
          A.store.Bool(ptr + 32 + 204 + 32 + 27, false);
          A.store.Int32(ptr + 32 + 204 + 32 + 8, 0);
          A.store.Bool(ptr + 32 + 204 + 32 + 28, false);
          A.store.Int32(ptr + 32 + 204 + 32 + 12, 0);
          A.store.Bool(ptr + 32 + 204 + 32 + 29, false);
          A.store.Int32(ptr + 32 + 204 + 32 + 16, 0);
          A.store.Bool(ptr + 32 + 204 + 32 + 30, false);
          A.store.Int32(ptr + 32 + 204 + 32 + 20, 0);
          A.store.Bool(ptr + 32 + 204 + 32 + 31, false);
          A.store.Bool(ptr + 32 + 204 + 32 + 24, false);
          A.store.Bool(ptr + 32 + 204 + 32 + 32, false);
          A.store.Bool(ptr + 32 + 204 + 32 + 25, false);
        } else {
          A.store.Bool(ptr + 32 + 271, true);

          if (typeof x["Manual"]["HTTPProxy"] === "undefined") {
            A.store.Bool(ptr + 32 + 0 + 66, false);

            A.store.Bool(ptr + 32 + 0 + 0 + 28, false);
            A.store.Ref(ptr + 32 + 0 + 0 + 0, undefined);
            A.store.Ref(ptr + 32 + 0 + 0 + 4, undefined);
            A.store.Ref(ptr + 32 + 0 + 0 + 8, undefined);
            A.store.Ref(ptr + 32 + 0 + 0 + 12, undefined);
            A.store.Ref(ptr + 32 + 0 + 0 + 16, undefined);
            A.store.Ref(ptr + 32 + 0 + 0 + 20, undefined);
            A.store.Bool(ptr + 32 + 0 + 0 + 26, false);
            A.store.Bool(ptr + 32 + 0 + 0 + 24, false);
            A.store.Bool(ptr + 32 + 0 + 0 + 27, false);
            A.store.Bool(ptr + 32 + 0 + 0 + 25, false);

            A.store.Bool(ptr + 32 + 0 + 32 + 33, false);
            A.store.Bool(ptr + 32 + 0 + 32 + 26, false);
            A.store.Int32(ptr + 32 + 0 + 32 + 0, 0);
            A.store.Ref(ptr + 32 + 0 + 32 + 4, undefined);
            A.store.Bool(ptr + 32 + 0 + 32 + 27, false);
            A.store.Int32(ptr + 32 + 0 + 32 + 8, 0);
            A.store.Bool(ptr + 32 + 0 + 32 + 28, false);
            A.store.Int32(ptr + 32 + 0 + 32 + 12, 0);
            A.store.Bool(ptr + 32 + 0 + 32 + 29, false);
            A.store.Int32(ptr + 32 + 0 + 32 + 16, 0);
            A.store.Bool(ptr + 32 + 0 + 32 + 30, false);
            A.store.Int32(ptr + 32 + 0 + 32 + 20, 0);
            A.store.Bool(ptr + 32 + 0 + 32 + 31, false);
            A.store.Bool(ptr + 32 + 0 + 32 + 24, false);
            A.store.Bool(ptr + 32 + 0 + 32 + 32, false);
            A.store.Bool(ptr + 32 + 0 + 32 + 25, false);
          } else {
            A.store.Bool(ptr + 32 + 0 + 66, true);

            if (typeof x["Manual"]["HTTPProxy"]["Host"] === "undefined") {
              A.store.Bool(ptr + 32 + 0 + 0 + 28, false);
              A.store.Ref(ptr + 32 + 0 + 0 + 0, undefined);
              A.store.Ref(ptr + 32 + 0 + 0 + 4, undefined);
              A.store.Ref(ptr + 32 + 0 + 0 + 8, undefined);
              A.store.Ref(ptr + 32 + 0 + 0 + 12, undefined);
              A.store.Ref(ptr + 32 + 0 + 0 + 16, undefined);
              A.store.Ref(ptr + 32 + 0 + 0 + 20, undefined);
              A.store.Bool(ptr + 32 + 0 + 0 + 26, false);
              A.store.Bool(ptr + 32 + 0 + 0 + 24, false);
              A.store.Bool(ptr + 32 + 0 + 0 + 27, false);
              A.store.Bool(ptr + 32 + 0 + 0 + 25, false);
            } else {
              A.store.Bool(ptr + 32 + 0 + 0 + 28, true);
              A.store.Ref(ptr + 32 + 0 + 0 + 0, x["Manual"]["HTTPProxy"]["Host"]["Active"]);
              A.store.Ref(ptr + 32 + 0 + 0 + 4, x["Manual"]["HTTPProxy"]["Host"]["Effective"]);
              A.store.Ref(ptr + 32 + 0 + 0 + 8, x["Manual"]["HTTPProxy"]["Host"]["UserPolicy"]);
              A.store.Ref(ptr + 32 + 0 + 0 + 12, x["Manual"]["HTTPProxy"]["Host"]["DevicePolicy"]);
              A.store.Ref(ptr + 32 + 0 + 0 + 16, x["Manual"]["HTTPProxy"]["Host"]["UserSetting"]);
              A.store.Ref(ptr + 32 + 0 + 0 + 20, x["Manual"]["HTTPProxy"]["Host"]["SharedSetting"]);
              A.store.Bool(ptr + 32 + 0 + 0 + 26, "UserEditable" in x["Manual"]["HTTPProxy"]["Host"] ? true : false);
              A.store.Bool(ptr + 32 + 0 + 0 + 24, x["Manual"]["HTTPProxy"]["Host"]["UserEditable"] ? true : false);
              A.store.Bool(ptr + 32 + 0 + 0 + 27, "DeviceEditable" in x["Manual"]["HTTPProxy"]["Host"] ? true : false);
              A.store.Bool(ptr + 32 + 0 + 0 + 25, x["Manual"]["HTTPProxy"]["Host"]["DeviceEditable"] ? true : false);
            }

            if (typeof x["Manual"]["HTTPProxy"]["Port"] === "undefined") {
              A.store.Bool(ptr + 32 + 0 + 32 + 33, false);
              A.store.Bool(ptr + 32 + 0 + 32 + 26, false);
              A.store.Int32(ptr + 32 + 0 + 32 + 0, 0);
              A.store.Ref(ptr + 32 + 0 + 32 + 4, undefined);
              A.store.Bool(ptr + 32 + 0 + 32 + 27, false);
              A.store.Int32(ptr + 32 + 0 + 32 + 8, 0);
              A.store.Bool(ptr + 32 + 0 + 32 + 28, false);
              A.store.Int32(ptr + 32 + 0 + 32 + 12, 0);
              A.store.Bool(ptr + 32 + 0 + 32 + 29, false);
              A.store.Int32(ptr + 32 + 0 + 32 + 16, 0);
              A.store.Bool(ptr + 32 + 0 + 32 + 30, false);
              A.store.Int32(ptr + 32 + 0 + 32 + 20, 0);
              A.store.Bool(ptr + 32 + 0 + 32 + 31, false);
              A.store.Bool(ptr + 32 + 0 + 32 + 24, false);
              A.store.Bool(ptr + 32 + 0 + 32 + 32, false);
              A.store.Bool(ptr + 32 + 0 + 32 + 25, false);
            } else {
              A.store.Bool(ptr + 32 + 0 + 32 + 33, true);
              A.store.Bool(ptr + 32 + 0 + 32 + 26, "Active" in x["Manual"]["HTTPProxy"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 0 + 32 + 0,
                x["Manual"]["HTTPProxy"]["Port"]["Active"] === undefined
                  ? 0
                  : (x["Manual"]["HTTPProxy"]["Port"]["Active"] as number)
              );
              A.store.Ref(ptr + 32 + 0 + 32 + 4, x["Manual"]["HTTPProxy"]["Port"]["Effective"]);
              A.store.Bool(ptr + 32 + 0 + 32 + 27, "UserPolicy" in x["Manual"]["HTTPProxy"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 0 + 32 + 8,
                x["Manual"]["HTTPProxy"]["Port"]["UserPolicy"] === undefined
                  ? 0
                  : (x["Manual"]["HTTPProxy"]["Port"]["UserPolicy"] as number)
              );
              A.store.Bool(ptr + 32 + 0 + 32 + 28, "DevicePolicy" in x["Manual"]["HTTPProxy"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 0 + 32 + 12,
                x["Manual"]["HTTPProxy"]["Port"]["DevicePolicy"] === undefined
                  ? 0
                  : (x["Manual"]["HTTPProxy"]["Port"]["DevicePolicy"] as number)
              );
              A.store.Bool(ptr + 32 + 0 + 32 + 29, "UserSetting" in x["Manual"]["HTTPProxy"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 0 + 32 + 16,
                x["Manual"]["HTTPProxy"]["Port"]["UserSetting"] === undefined
                  ? 0
                  : (x["Manual"]["HTTPProxy"]["Port"]["UserSetting"] as number)
              );
              A.store.Bool(ptr + 32 + 0 + 32 + 30, "SharedSetting" in x["Manual"]["HTTPProxy"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 0 + 32 + 20,
                x["Manual"]["HTTPProxy"]["Port"]["SharedSetting"] === undefined
                  ? 0
                  : (x["Manual"]["HTTPProxy"]["Port"]["SharedSetting"] as number)
              );
              A.store.Bool(ptr + 32 + 0 + 32 + 31, "UserEditable" in x["Manual"]["HTTPProxy"]["Port"] ? true : false);
              A.store.Bool(ptr + 32 + 0 + 32 + 24, x["Manual"]["HTTPProxy"]["Port"]["UserEditable"] ? true : false);
              A.store.Bool(ptr + 32 + 0 + 32 + 32, "DeviceEditable" in x["Manual"]["HTTPProxy"]["Port"] ? true : false);
              A.store.Bool(ptr + 32 + 0 + 32 + 25, x["Manual"]["HTTPProxy"]["Port"]["DeviceEditable"] ? true : false);
            }
          }

          if (typeof x["Manual"]["SecureHTTPProxy"] === "undefined") {
            A.store.Bool(ptr + 32 + 68 + 66, false);

            A.store.Bool(ptr + 32 + 68 + 0 + 28, false);
            A.store.Ref(ptr + 32 + 68 + 0 + 0, undefined);
            A.store.Ref(ptr + 32 + 68 + 0 + 4, undefined);
            A.store.Ref(ptr + 32 + 68 + 0 + 8, undefined);
            A.store.Ref(ptr + 32 + 68 + 0 + 12, undefined);
            A.store.Ref(ptr + 32 + 68 + 0 + 16, undefined);
            A.store.Ref(ptr + 32 + 68 + 0 + 20, undefined);
            A.store.Bool(ptr + 32 + 68 + 0 + 26, false);
            A.store.Bool(ptr + 32 + 68 + 0 + 24, false);
            A.store.Bool(ptr + 32 + 68 + 0 + 27, false);
            A.store.Bool(ptr + 32 + 68 + 0 + 25, false);

            A.store.Bool(ptr + 32 + 68 + 32 + 33, false);
            A.store.Bool(ptr + 32 + 68 + 32 + 26, false);
            A.store.Int32(ptr + 32 + 68 + 32 + 0, 0);
            A.store.Ref(ptr + 32 + 68 + 32 + 4, undefined);
            A.store.Bool(ptr + 32 + 68 + 32 + 27, false);
            A.store.Int32(ptr + 32 + 68 + 32 + 8, 0);
            A.store.Bool(ptr + 32 + 68 + 32 + 28, false);
            A.store.Int32(ptr + 32 + 68 + 32 + 12, 0);
            A.store.Bool(ptr + 32 + 68 + 32 + 29, false);
            A.store.Int32(ptr + 32 + 68 + 32 + 16, 0);
            A.store.Bool(ptr + 32 + 68 + 32 + 30, false);
            A.store.Int32(ptr + 32 + 68 + 32 + 20, 0);
            A.store.Bool(ptr + 32 + 68 + 32 + 31, false);
            A.store.Bool(ptr + 32 + 68 + 32 + 24, false);
            A.store.Bool(ptr + 32 + 68 + 32 + 32, false);
            A.store.Bool(ptr + 32 + 68 + 32 + 25, false);
          } else {
            A.store.Bool(ptr + 32 + 68 + 66, true);

            if (typeof x["Manual"]["SecureHTTPProxy"]["Host"] === "undefined") {
              A.store.Bool(ptr + 32 + 68 + 0 + 28, false);
              A.store.Ref(ptr + 32 + 68 + 0 + 0, undefined);
              A.store.Ref(ptr + 32 + 68 + 0 + 4, undefined);
              A.store.Ref(ptr + 32 + 68 + 0 + 8, undefined);
              A.store.Ref(ptr + 32 + 68 + 0 + 12, undefined);
              A.store.Ref(ptr + 32 + 68 + 0 + 16, undefined);
              A.store.Ref(ptr + 32 + 68 + 0 + 20, undefined);
              A.store.Bool(ptr + 32 + 68 + 0 + 26, false);
              A.store.Bool(ptr + 32 + 68 + 0 + 24, false);
              A.store.Bool(ptr + 32 + 68 + 0 + 27, false);
              A.store.Bool(ptr + 32 + 68 + 0 + 25, false);
            } else {
              A.store.Bool(ptr + 32 + 68 + 0 + 28, true);
              A.store.Ref(ptr + 32 + 68 + 0 + 0, x["Manual"]["SecureHTTPProxy"]["Host"]["Active"]);
              A.store.Ref(ptr + 32 + 68 + 0 + 4, x["Manual"]["SecureHTTPProxy"]["Host"]["Effective"]);
              A.store.Ref(ptr + 32 + 68 + 0 + 8, x["Manual"]["SecureHTTPProxy"]["Host"]["UserPolicy"]);
              A.store.Ref(ptr + 32 + 68 + 0 + 12, x["Manual"]["SecureHTTPProxy"]["Host"]["DevicePolicy"]);
              A.store.Ref(ptr + 32 + 68 + 0 + 16, x["Manual"]["SecureHTTPProxy"]["Host"]["UserSetting"]);
              A.store.Ref(ptr + 32 + 68 + 0 + 20, x["Manual"]["SecureHTTPProxy"]["Host"]["SharedSetting"]);
              A.store.Bool(
                ptr + 32 + 68 + 0 + 26,
                "UserEditable" in x["Manual"]["SecureHTTPProxy"]["Host"] ? true : false
              );
              A.store.Bool(
                ptr + 32 + 68 + 0 + 24,
                x["Manual"]["SecureHTTPProxy"]["Host"]["UserEditable"] ? true : false
              );
              A.store.Bool(
                ptr + 32 + 68 + 0 + 27,
                "DeviceEditable" in x["Manual"]["SecureHTTPProxy"]["Host"] ? true : false
              );
              A.store.Bool(
                ptr + 32 + 68 + 0 + 25,
                x["Manual"]["SecureHTTPProxy"]["Host"]["DeviceEditable"] ? true : false
              );
            }

            if (typeof x["Manual"]["SecureHTTPProxy"]["Port"] === "undefined") {
              A.store.Bool(ptr + 32 + 68 + 32 + 33, false);
              A.store.Bool(ptr + 32 + 68 + 32 + 26, false);
              A.store.Int32(ptr + 32 + 68 + 32 + 0, 0);
              A.store.Ref(ptr + 32 + 68 + 32 + 4, undefined);
              A.store.Bool(ptr + 32 + 68 + 32 + 27, false);
              A.store.Int32(ptr + 32 + 68 + 32 + 8, 0);
              A.store.Bool(ptr + 32 + 68 + 32 + 28, false);
              A.store.Int32(ptr + 32 + 68 + 32 + 12, 0);
              A.store.Bool(ptr + 32 + 68 + 32 + 29, false);
              A.store.Int32(ptr + 32 + 68 + 32 + 16, 0);
              A.store.Bool(ptr + 32 + 68 + 32 + 30, false);
              A.store.Int32(ptr + 32 + 68 + 32 + 20, 0);
              A.store.Bool(ptr + 32 + 68 + 32 + 31, false);
              A.store.Bool(ptr + 32 + 68 + 32 + 24, false);
              A.store.Bool(ptr + 32 + 68 + 32 + 32, false);
              A.store.Bool(ptr + 32 + 68 + 32 + 25, false);
            } else {
              A.store.Bool(ptr + 32 + 68 + 32 + 33, true);
              A.store.Bool(ptr + 32 + 68 + 32 + 26, "Active" in x["Manual"]["SecureHTTPProxy"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 68 + 32 + 0,
                x["Manual"]["SecureHTTPProxy"]["Port"]["Active"] === undefined
                  ? 0
                  : (x["Manual"]["SecureHTTPProxy"]["Port"]["Active"] as number)
              );
              A.store.Ref(ptr + 32 + 68 + 32 + 4, x["Manual"]["SecureHTTPProxy"]["Port"]["Effective"]);
              A.store.Bool(
                ptr + 32 + 68 + 32 + 27,
                "UserPolicy" in x["Manual"]["SecureHTTPProxy"]["Port"] ? true : false
              );
              A.store.Int32(
                ptr + 32 + 68 + 32 + 8,
                x["Manual"]["SecureHTTPProxy"]["Port"]["UserPolicy"] === undefined
                  ? 0
                  : (x["Manual"]["SecureHTTPProxy"]["Port"]["UserPolicy"] as number)
              );
              A.store.Bool(
                ptr + 32 + 68 + 32 + 28,
                "DevicePolicy" in x["Manual"]["SecureHTTPProxy"]["Port"] ? true : false
              );
              A.store.Int32(
                ptr + 32 + 68 + 32 + 12,
                x["Manual"]["SecureHTTPProxy"]["Port"]["DevicePolicy"] === undefined
                  ? 0
                  : (x["Manual"]["SecureHTTPProxy"]["Port"]["DevicePolicy"] as number)
              );
              A.store.Bool(
                ptr + 32 + 68 + 32 + 29,
                "UserSetting" in x["Manual"]["SecureHTTPProxy"]["Port"] ? true : false
              );
              A.store.Int32(
                ptr + 32 + 68 + 32 + 16,
                x["Manual"]["SecureHTTPProxy"]["Port"]["UserSetting"] === undefined
                  ? 0
                  : (x["Manual"]["SecureHTTPProxy"]["Port"]["UserSetting"] as number)
              );
              A.store.Bool(
                ptr + 32 + 68 + 32 + 30,
                "SharedSetting" in x["Manual"]["SecureHTTPProxy"]["Port"] ? true : false
              );
              A.store.Int32(
                ptr + 32 + 68 + 32 + 20,
                x["Manual"]["SecureHTTPProxy"]["Port"]["SharedSetting"] === undefined
                  ? 0
                  : (x["Manual"]["SecureHTTPProxy"]["Port"]["SharedSetting"] as number)
              );
              A.store.Bool(
                ptr + 32 + 68 + 32 + 31,
                "UserEditable" in x["Manual"]["SecureHTTPProxy"]["Port"] ? true : false
              );
              A.store.Bool(
                ptr + 32 + 68 + 32 + 24,
                x["Manual"]["SecureHTTPProxy"]["Port"]["UserEditable"] ? true : false
              );
              A.store.Bool(
                ptr + 32 + 68 + 32 + 32,
                "DeviceEditable" in x["Manual"]["SecureHTTPProxy"]["Port"] ? true : false
              );
              A.store.Bool(
                ptr + 32 + 68 + 32 + 25,
                x["Manual"]["SecureHTTPProxy"]["Port"]["DeviceEditable"] ? true : false
              );
            }
          }

          if (typeof x["Manual"]["FTPProxy"] === "undefined") {
            A.store.Bool(ptr + 32 + 136 + 66, false);

            A.store.Bool(ptr + 32 + 136 + 0 + 28, false);
            A.store.Ref(ptr + 32 + 136 + 0 + 0, undefined);
            A.store.Ref(ptr + 32 + 136 + 0 + 4, undefined);
            A.store.Ref(ptr + 32 + 136 + 0 + 8, undefined);
            A.store.Ref(ptr + 32 + 136 + 0 + 12, undefined);
            A.store.Ref(ptr + 32 + 136 + 0 + 16, undefined);
            A.store.Ref(ptr + 32 + 136 + 0 + 20, undefined);
            A.store.Bool(ptr + 32 + 136 + 0 + 26, false);
            A.store.Bool(ptr + 32 + 136 + 0 + 24, false);
            A.store.Bool(ptr + 32 + 136 + 0 + 27, false);
            A.store.Bool(ptr + 32 + 136 + 0 + 25, false);

            A.store.Bool(ptr + 32 + 136 + 32 + 33, false);
            A.store.Bool(ptr + 32 + 136 + 32 + 26, false);
            A.store.Int32(ptr + 32 + 136 + 32 + 0, 0);
            A.store.Ref(ptr + 32 + 136 + 32 + 4, undefined);
            A.store.Bool(ptr + 32 + 136 + 32 + 27, false);
            A.store.Int32(ptr + 32 + 136 + 32 + 8, 0);
            A.store.Bool(ptr + 32 + 136 + 32 + 28, false);
            A.store.Int32(ptr + 32 + 136 + 32 + 12, 0);
            A.store.Bool(ptr + 32 + 136 + 32 + 29, false);
            A.store.Int32(ptr + 32 + 136 + 32 + 16, 0);
            A.store.Bool(ptr + 32 + 136 + 32 + 30, false);
            A.store.Int32(ptr + 32 + 136 + 32 + 20, 0);
            A.store.Bool(ptr + 32 + 136 + 32 + 31, false);
            A.store.Bool(ptr + 32 + 136 + 32 + 24, false);
            A.store.Bool(ptr + 32 + 136 + 32 + 32, false);
            A.store.Bool(ptr + 32 + 136 + 32 + 25, false);
          } else {
            A.store.Bool(ptr + 32 + 136 + 66, true);

            if (typeof x["Manual"]["FTPProxy"]["Host"] === "undefined") {
              A.store.Bool(ptr + 32 + 136 + 0 + 28, false);
              A.store.Ref(ptr + 32 + 136 + 0 + 0, undefined);
              A.store.Ref(ptr + 32 + 136 + 0 + 4, undefined);
              A.store.Ref(ptr + 32 + 136 + 0 + 8, undefined);
              A.store.Ref(ptr + 32 + 136 + 0 + 12, undefined);
              A.store.Ref(ptr + 32 + 136 + 0 + 16, undefined);
              A.store.Ref(ptr + 32 + 136 + 0 + 20, undefined);
              A.store.Bool(ptr + 32 + 136 + 0 + 26, false);
              A.store.Bool(ptr + 32 + 136 + 0 + 24, false);
              A.store.Bool(ptr + 32 + 136 + 0 + 27, false);
              A.store.Bool(ptr + 32 + 136 + 0 + 25, false);
            } else {
              A.store.Bool(ptr + 32 + 136 + 0 + 28, true);
              A.store.Ref(ptr + 32 + 136 + 0 + 0, x["Manual"]["FTPProxy"]["Host"]["Active"]);
              A.store.Ref(ptr + 32 + 136 + 0 + 4, x["Manual"]["FTPProxy"]["Host"]["Effective"]);
              A.store.Ref(ptr + 32 + 136 + 0 + 8, x["Manual"]["FTPProxy"]["Host"]["UserPolicy"]);
              A.store.Ref(ptr + 32 + 136 + 0 + 12, x["Manual"]["FTPProxy"]["Host"]["DevicePolicy"]);
              A.store.Ref(ptr + 32 + 136 + 0 + 16, x["Manual"]["FTPProxy"]["Host"]["UserSetting"]);
              A.store.Ref(ptr + 32 + 136 + 0 + 20, x["Manual"]["FTPProxy"]["Host"]["SharedSetting"]);
              A.store.Bool(ptr + 32 + 136 + 0 + 26, "UserEditable" in x["Manual"]["FTPProxy"]["Host"] ? true : false);
              A.store.Bool(ptr + 32 + 136 + 0 + 24, x["Manual"]["FTPProxy"]["Host"]["UserEditable"] ? true : false);
              A.store.Bool(ptr + 32 + 136 + 0 + 27, "DeviceEditable" in x["Manual"]["FTPProxy"]["Host"] ? true : false);
              A.store.Bool(ptr + 32 + 136 + 0 + 25, x["Manual"]["FTPProxy"]["Host"]["DeviceEditable"] ? true : false);
            }

            if (typeof x["Manual"]["FTPProxy"]["Port"] === "undefined") {
              A.store.Bool(ptr + 32 + 136 + 32 + 33, false);
              A.store.Bool(ptr + 32 + 136 + 32 + 26, false);
              A.store.Int32(ptr + 32 + 136 + 32 + 0, 0);
              A.store.Ref(ptr + 32 + 136 + 32 + 4, undefined);
              A.store.Bool(ptr + 32 + 136 + 32 + 27, false);
              A.store.Int32(ptr + 32 + 136 + 32 + 8, 0);
              A.store.Bool(ptr + 32 + 136 + 32 + 28, false);
              A.store.Int32(ptr + 32 + 136 + 32 + 12, 0);
              A.store.Bool(ptr + 32 + 136 + 32 + 29, false);
              A.store.Int32(ptr + 32 + 136 + 32 + 16, 0);
              A.store.Bool(ptr + 32 + 136 + 32 + 30, false);
              A.store.Int32(ptr + 32 + 136 + 32 + 20, 0);
              A.store.Bool(ptr + 32 + 136 + 32 + 31, false);
              A.store.Bool(ptr + 32 + 136 + 32 + 24, false);
              A.store.Bool(ptr + 32 + 136 + 32 + 32, false);
              A.store.Bool(ptr + 32 + 136 + 32 + 25, false);
            } else {
              A.store.Bool(ptr + 32 + 136 + 32 + 33, true);
              A.store.Bool(ptr + 32 + 136 + 32 + 26, "Active" in x["Manual"]["FTPProxy"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 136 + 32 + 0,
                x["Manual"]["FTPProxy"]["Port"]["Active"] === undefined
                  ? 0
                  : (x["Manual"]["FTPProxy"]["Port"]["Active"] as number)
              );
              A.store.Ref(ptr + 32 + 136 + 32 + 4, x["Manual"]["FTPProxy"]["Port"]["Effective"]);
              A.store.Bool(ptr + 32 + 136 + 32 + 27, "UserPolicy" in x["Manual"]["FTPProxy"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 136 + 32 + 8,
                x["Manual"]["FTPProxy"]["Port"]["UserPolicy"] === undefined
                  ? 0
                  : (x["Manual"]["FTPProxy"]["Port"]["UserPolicy"] as number)
              );
              A.store.Bool(ptr + 32 + 136 + 32 + 28, "DevicePolicy" in x["Manual"]["FTPProxy"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 136 + 32 + 12,
                x["Manual"]["FTPProxy"]["Port"]["DevicePolicy"] === undefined
                  ? 0
                  : (x["Manual"]["FTPProxy"]["Port"]["DevicePolicy"] as number)
              );
              A.store.Bool(ptr + 32 + 136 + 32 + 29, "UserSetting" in x["Manual"]["FTPProxy"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 136 + 32 + 16,
                x["Manual"]["FTPProxy"]["Port"]["UserSetting"] === undefined
                  ? 0
                  : (x["Manual"]["FTPProxy"]["Port"]["UserSetting"] as number)
              );
              A.store.Bool(ptr + 32 + 136 + 32 + 30, "SharedSetting" in x["Manual"]["FTPProxy"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 136 + 32 + 20,
                x["Manual"]["FTPProxy"]["Port"]["SharedSetting"] === undefined
                  ? 0
                  : (x["Manual"]["FTPProxy"]["Port"]["SharedSetting"] as number)
              );
              A.store.Bool(ptr + 32 + 136 + 32 + 31, "UserEditable" in x["Manual"]["FTPProxy"]["Port"] ? true : false);
              A.store.Bool(ptr + 32 + 136 + 32 + 24, x["Manual"]["FTPProxy"]["Port"]["UserEditable"] ? true : false);
              A.store.Bool(
                ptr + 32 + 136 + 32 + 32,
                "DeviceEditable" in x["Manual"]["FTPProxy"]["Port"] ? true : false
              );
              A.store.Bool(ptr + 32 + 136 + 32 + 25, x["Manual"]["FTPProxy"]["Port"]["DeviceEditable"] ? true : false);
            }
          }

          if (typeof x["Manual"]["SOCKS"] === "undefined") {
            A.store.Bool(ptr + 32 + 204 + 66, false);

            A.store.Bool(ptr + 32 + 204 + 0 + 28, false);
            A.store.Ref(ptr + 32 + 204 + 0 + 0, undefined);
            A.store.Ref(ptr + 32 + 204 + 0 + 4, undefined);
            A.store.Ref(ptr + 32 + 204 + 0 + 8, undefined);
            A.store.Ref(ptr + 32 + 204 + 0 + 12, undefined);
            A.store.Ref(ptr + 32 + 204 + 0 + 16, undefined);
            A.store.Ref(ptr + 32 + 204 + 0 + 20, undefined);
            A.store.Bool(ptr + 32 + 204 + 0 + 26, false);
            A.store.Bool(ptr + 32 + 204 + 0 + 24, false);
            A.store.Bool(ptr + 32 + 204 + 0 + 27, false);
            A.store.Bool(ptr + 32 + 204 + 0 + 25, false);

            A.store.Bool(ptr + 32 + 204 + 32 + 33, false);
            A.store.Bool(ptr + 32 + 204 + 32 + 26, false);
            A.store.Int32(ptr + 32 + 204 + 32 + 0, 0);
            A.store.Ref(ptr + 32 + 204 + 32 + 4, undefined);
            A.store.Bool(ptr + 32 + 204 + 32 + 27, false);
            A.store.Int32(ptr + 32 + 204 + 32 + 8, 0);
            A.store.Bool(ptr + 32 + 204 + 32 + 28, false);
            A.store.Int32(ptr + 32 + 204 + 32 + 12, 0);
            A.store.Bool(ptr + 32 + 204 + 32 + 29, false);
            A.store.Int32(ptr + 32 + 204 + 32 + 16, 0);
            A.store.Bool(ptr + 32 + 204 + 32 + 30, false);
            A.store.Int32(ptr + 32 + 204 + 32 + 20, 0);
            A.store.Bool(ptr + 32 + 204 + 32 + 31, false);
            A.store.Bool(ptr + 32 + 204 + 32 + 24, false);
            A.store.Bool(ptr + 32 + 204 + 32 + 32, false);
            A.store.Bool(ptr + 32 + 204 + 32 + 25, false);
          } else {
            A.store.Bool(ptr + 32 + 204 + 66, true);

            if (typeof x["Manual"]["SOCKS"]["Host"] === "undefined") {
              A.store.Bool(ptr + 32 + 204 + 0 + 28, false);
              A.store.Ref(ptr + 32 + 204 + 0 + 0, undefined);
              A.store.Ref(ptr + 32 + 204 + 0 + 4, undefined);
              A.store.Ref(ptr + 32 + 204 + 0 + 8, undefined);
              A.store.Ref(ptr + 32 + 204 + 0 + 12, undefined);
              A.store.Ref(ptr + 32 + 204 + 0 + 16, undefined);
              A.store.Ref(ptr + 32 + 204 + 0 + 20, undefined);
              A.store.Bool(ptr + 32 + 204 + 0 + 26, false);
              A.store.Bool(ptr + 32 + 204 + 0 + 24, false);
              A.store.Bool(ptr + 32 + 204 + 0 + 27, false);
              A.store.Bool(ptr + 32 + 204 + 0 + 25, false);
            } else {
              A.store.Bool(ptr + 32 + 204 + 0 + 28, true);
              A.store.Ref(ptr + 32 + 204 + 0 + 0, x["Manual"]["SOCKS"]["Host"]["Active"]);
              A.store.Ref(ptr + 32 + 204 + 0 + 4, x["Manual"]["SOCKS"]["Host"]["Effective"]);
              A.store.Ref(ptr + 32 + 204 + 0 + 8, x["Manual"]["SOCKS"]["Host"]["UserPolicy"]);
              A.store.Ref(ptr + 32 + 204 + 0 + 12, x["Manual"]["SOCKS"]["Host"]["DevicePolicy"]);
              A.store.Ref(ptr + 32 + 204 + 0 + 16, x["Manual"]["SOCKS"]["Host"]["UserSetting"]);
              A.store.Ref(ptr + 32 + 204 + 0 + 20, x["Manual"]["SOCKS"]["Host"]["SharedSetting"]);
              A.store.Bool(ptr + 32 + 204 + 0 + 26, "UserEditable" in x["Manual"]["SOCKS"]["Host"] ? true : false);
              A.store.Bool(ptr + 32 + 204 + 0 + 24, x["Manual"]["SOCKS"]["Host"]["UserEditable"] ? true : false);
              A.store.Bool(ptr + 32 + 204 + 0 + 27, "DeviceEditable" in x["Manual"]["SOCKS"]["Host"] ? true : false);
              A.store.Bool(ptr + 32 + 204 + 0 + 25, x["Manual"]["SOCKS"]["Host"]["DeviceEditable"] ? true : false);
            }

            if (typeof x["Manual"]["SOCKS"]["Port"] === "undefined") {
              A.store.Bool(ptr + 32 + 204 + 32 + 33, false);
              A.store.Bool(ptr + 32 + 204 + 32 + 26, false);
              A.store.Int32(ptr + 32 + 204 + 32 + 0, 0);
              A.store.Ref(ptr + 32 + 204 + 32 + 4, undefined);
              A.store.Bool(ptr + 32 + 204 + 32 + 27, false);
              A.store.Int32(ptr + 32 + 204 + 32 + 8, 0);
              A.store.Bool(ptr + 32 + 204 + 32 + 28, false);
              A.store.Int32(ptr + 32 + 204 + 32 + 12, 0);
              A.store.Bool(ptr + 32 + 204 + 32 + 29, false);
              A.store.Int32(ptr + 32 + 204 + 32 + 16, 0);
              A.store.Bool(ptr + 32 + 204 + 32 + 30, false);
              A.store.Int32(ptr + 32 + 204 + 32 + 20, 0);
              A.store.Bool(ptr + 32 + 204 + 32 + 31, false);
              A.store.Bool(ptr + 32 + 204 + 32 + 24, false);
              A.store.Bool(ptr + 32 + 204 + 32 + 32, false);
              A.store.Bool(ptr + 32 + 204 + 32 + 25, false);
            } else {
              A.store.Bool(ptr + 32 + 204 + 32 + 33, true);
              A.store.Bool(ptr + 32 + 204 + 32 + 26, "Active" in x["Manual"]["SOCKS"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 204 + 32 + 0,
                x["Manual"]["SOCKS"]["Port"]["Active"] === undefined
                  ? 0
                  : (x["Manual"]["SOCKS"]["Port"]["Active"] as number)
              );
              A.store.Ref(ptr + 32 + 204 + 32 + 4, x["Manual"]["SOCKS"]["Port"]["Effective"]);
              A.store.Bool(ptr + 32 + 204 + 32 + 27, "UserPolicy" in x["Manual"]["SOCKS"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 204 + 32 + 8,
                x["Manual"]["SOCKS"]["Port"]["UserPolicy"] === undefined
                  ? 0
                  : (x["Manual"]["SOCKS"]["Port"]["UserPolicy"] as number)
              );
              A.store.Bool(ptr + 32 + 204 + 32 + 28, "DevicePolicy" in x["Manual"]["SOCKS"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 204 + 32 + 12,
                x["Manual"]["SOCKS"]["Port"]["DevicePolicy"] === undefined
                  ? 0
                  : (x["Manual"]["SOCKS"]["Port"]["DevicePolicy"] as number)
              );
              A.store.Bool(ptr + 32 + 204 + 32 + 29, "UserSetting" in x["Manual"]["SOCKS"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 204 + 32 + 16,
                x["Manual"]["SOCKS"]["Port"]["UserSetting"] === undefined
                  ? 0
                  : (x["Manual"]["SOCKS"]["Port"]["UserSetting"] as number)
              );
              A.store.Bool(ptr + 32 + 204 + 32 + 30, "SharedSetting" in x["Manual"]["SOCKS"]["Port"] ? true : false);
              A.store.Int32(
                ptr + 32 + 204 + 32 + 20,
                x["Manual"]["SOCKS"]["Port"]["SharedSetting"] === undefined
                  ? 0
                  : (x["Manual"]["SOCKS"]["Port"]["SharedSetting"] as number)
              );
              A.store.Bool(ptr + 32 + 204 + 32 + 31, "UserEditable" in x["Manual"]["SOCKS"]["Port"] ? true : false);
              A.store.Bool(ptr + 32 + 204 + 32 + 24, x["Manual"]["SOCKS"]["Port"]["UserEditable"] ? true : false);
              A.store.Bool(ptr + 32 + 204 + 32 + 32, "DeviceEditable" in x["Manual"]["SOCKS"]["Port"] ? true : false);
              A.store.Bool(ptr + 32 + 204 + 32 + 25, x["Manual"]["SOCKS"]["Port"]["DeviceEditable"] ? true : false);
            }
          }
        }

        if (typeof x["ExcludeDomains"] === "undefined") {
          A.store.Bool(ptr + 304 + 28, false);
          A.store.Ref(ptr + 304 + 0, undefined);
          A.store.Ref(ptr + 304 + 4, undefined);
          A.store.Ref(ptr + 304 + 8, undefined);
          A.store.Ref(ptr + 304 + 12, undefined);
          A.store.Ref(ptr + 304 + 16, undefined);
          A.store.Ref(ptr + 304 + 20, undefined);
          A.store.Bool(ptr + 304 + 26, false);
          A.store.Bool(ptr + 304 + 24, false);
          A.store.Bool(ptr + 304 + 27, false);
          A.store.Bool(ptr + 304 + 25, false);
        } else {
          A.store.Bool(ptr + 304 + 28, true);
          A.store.Ref(ptr + 304 + 0, x["ExcludeDomains"]["Active"]);
          A.store.Ref(ptr + 304 + 4, x["ExcludeDomains"]["Effective"]);
          A.store.Ref(ptr + 304 + 8, x["ExcludeDomains"]["UserPolicy"]);
          A.store.Ref(ptr + 304 + 12, x["ExcludeDomains"]["DevicePolicy"]);
          A.store.Ref(ptr + 304 + 16, x["ExcludeDomains"]["UserSetting"]);
          A.store.Ref(ptr + 304 + 20, x["ExcludeDomains"]["SharedSetting"]);
          A.store.Bool(ptr + 304 + 26, "UserEditable" in x["ExcludeDomains"] ? true : false);
          A.store.Bool(ptr + 304 + 24, x["ExcludeDomains"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 304 + 27, "DeviceEditable" in x["ExcludeDomains"] ? true : false);
          A.store.Bool(ptr + 304 + 25, x["ExcludeDomains"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["PAC"] === "undefined") {
          A.store.Bool(ptr + 336 + 28, false);
          A.store.Ref(ptr + 336 + 0, undefined);
          A.store.Ref(ptr + 336 + 4, undefined);
          A.store.Ref(ptr + 336 + 8, undefined);
          A.store.Ref(ptr + 336 + 12, undefined);
          A.store.Ref(ptr + 336 + 16, undefined);
          A.store.Ref(ptr + 336 + 20, undefined);
          A.store.Bool(ptr + 336 + 26, false);
          A.store.Bool(ptr + 336 + 24, false);
          A.store.Bool(ptr + 336 + 27, false);
          A.store.Bool(ptr + 336 + 25, false);
        } else {
          A.store.Bool(ptr + 336 + 28, true);
          A.store.Ref(ptr + 336 + 0, x["PAC"]["Active"]);
          A.store.Ref(ptr + 336 + 4, x["PAC"]["Effective"]);
          A.store.Ref(ptr + 336 + 8, x["PAC"]["UserPolicy"]);
          A.store.Ref(ptr + 336 + 12, x["PAC"]["DevicePolicy"]);
          A.store.Ref(ptr + 336 + 16, x["PAC"]["UserSetting"]);
          A.store.Ref(ptr + 336 + 20, x["PAC"]["SharedSetting"]);
          A.store.Bool(ptr + 336 + 26, "UserEditable" in x["PAC"] ? true : false);
          A.store.Bool(ptr + 336 + 24, x["PAC"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 336 + 27, "DeviceEditable" in x["PAC"] ? true : false);
          A.store.Bool(ptr + 336 + 25, x["PAC"]["DeviceEditable"] ? true : false);
        }
      }
    },
    "load_ManagedProxySettings": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 28)) {
        x["Type"] = {};
        x["Type"]["Active"] = A.load.Enum(ptr + 0 + 0, ["Direct", "Manual", "PAC", "WPAD"]);
        x["Type"]["Effective"] = A.load.Ref(ptr + 0 + 4, undefined);
        x["Type"]["UserPolicy"] = A.load.Enum(ptr + 0 + 8, ["Direct", "Manual", "PAC", "WPAD"]);
        x["Type"]["DevicePolicy"] = A.load.Enum(ptr + 0 + 12, ["Direct", "Manual", "PAC", "WPAD"]);
        x["Type"]["UserSetting"] = A.load.Enum(ptr + 0 + 16, ["Direct", "Manual", "PAC", "WPAD"]);
        x["Type"]["SharedSetting"] = A.load.Enum(ptr + 0 + 20, ["Direct", "Manual", "PAC", "WPAD"]);
        if (A.load.Bool(ptr + 0 + 26)) {
          x["Type"]["UserEditable"] = A.load.Bool(ptr + 0 + 24);
        } else {
          delete x["Type"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 0 + 27)) {
          x["Type"]["DeviceEditable"] = A.load.Bool(ptr + 0 + 25);
        } else {
          delete x["Type"]["DeviceEditable"];
        }
      } else {
        delete x["Type"];
      }
      if (A.load.Bool(ptr + 32 + 271)) {
        x["Manual"] = {};
        if (A.load.Bool(ptr + 32 + 0 + 66)) {
          x["Manual"]["HTTPProxy"] = {};
          if (A.load.Bool(ptr + 32 + 0 + 0 + 28)) {
            x["Manual"]["HTTPProxy"]["Host"] = {};
            x["Manual"]["HTTPProxy"]["Host"]["Active"] = A.load.Ref(ptr + 32 + 0 + 0 + 0, undefined);
            x["Manual"]["HTTPProxy"]["Host"]["Effective"] = A.load.Ref(ptr + 32 + 0 + 0 + 4, undefined);
            x["Manual"]["HTTPProxy"]["Host"]["UserPolicy"] = A.load.Ref(ptr + 32 + 0 + 0 + 8, undefined);
            x["Manual"]["HTTPProxy"]["Host"]["DevicePolicy"] = A.load.Ref(ptr + 32 + 0 + 0 + 12, undefined);
            x["Manual"]["HTTPProxy"]["Host"]["UserSetting"] = A.load.Ref(ptr + 32 + 0 + 0 + 16, undefined);
            x["Manual"]["HTTPProxy"]["Host"]["SharedSetting"] = A.load.Ref(ptr + 32 + 0 + 0 + 20, undefined);
            if (A.load.Bool(ptr + 32 + 0 + 0 + 26)) {
              x["Manual"]["HTTPProxy"]["Host"]["UserEditable"] = A.load.Bool(ptr + 32 + 0 + 0 + 24);
            } else {
              delete x["Manual"]["HTTPProxy"]["Host"]["UserEditable"];
            }
            if (A.load.Bool(ptr + 32 + 0 + 0 + 27)) {
              x["Manual"]["HTTPProxy"]["Host"]["DeviceEditable"] = A.load.Bool(ptr + 32 + 0 + 0 + 25);
            } else {
              delete x["Manual"]["HTTPProxy"]["Host"]["DeviceEditable"];
            }
          } else {
            delete x["Manual"]["HTTPProxy"]["Host"];
          }
          if (A.load.Bool(ptr + 32 + 0 + 32 + 33)) {
            x["Manual"]["HTTPProxy"]["Port"] = {};
            if (A.load.Bool(ptr + 32 + 0 + 32 + 26)) {
              x["Manual"]["HTTPProxy"]["Port"]["Active"] = A.load.Int32(ptr + 32 + 0 + 32 + 0);
            } else {
              delete x["Manual"]["HTTPProxy"]["Port"]["Active"];
            }
            x["Manual"]["HTTPProxy"]["Port"]["Effective"] = A.load.Ref(ptr + 32 + 0 + 32 + 4, undefined);
            if (A.load.Bool(ptr + 32 + 0 + 32 + 27)) {
              x["Manual"]["HTTPProxy"]["Port"]["UserPolicy"] = A.load.Int32(ptr + 32 + 0 + 32 + 8);
            } else {
              delete x["Manual"]["HTTPProxy"]["Port"]["UserPolicy"];
            }
            if (A.load.Bool(ptr + 32 + 0 + 32 + 28)) {
              x["Manual"]["HTTPProxy"]["Port"]["DevicePolicy"] = A.load.Int32(ptr + 32 + 0 + 32 + 12);
            } else {
              delete x["Manual"]["HTTPProxy"]["Port"]["DevicePolicy"];
            }
            if (A.load.Bool(ptr + 32 + 0 + 32 + 29)) {
              x["Manual"]["HTTPProxy"]["Port"]["UserSetting"] = A.load.Int32(ptr + 32 + 0 + 32 + 16);
            } else {
              delete x["Manual"]["HTTPProxy"]["Port"]["UserSetting"];
            }
            if (A.load.Bool(ptr + 32 + 0 + 32 + 30)) {
              x["Manual"]["HTTPProxy"]["Port"]["SharedSetting"] = A.load.Int32(ptr + 32 + 0 + 32 + 20);
            } else {
              delete x["Manual"]["HTTPProxy"]["Port"]["SharedSetting"];
            }
            if (A.load.Bool(ptr + 32 + 0 + 32 + 31)) {
              x["Manual"]["HTTPProxy"]["Port"]["UserEditable"] = A.load.Bool(ptr + 32 + 0 + 32 + 24);
            } else {
              delete x["Manual"]["HTTPProxy"]["Port"]["UserEditable"];
            }
            if (A.load.Bool(ptr + 32 + 0 + 32 + 32)) {
              x["Manual"]["HTTPProxy"]["Port"]["DeviceEditable"] = A.load.Bool(ptr + 32 + 0 + 32 + 25);
            } else {
              delete x["Manual"]["HTTPProxy"]["Port"]["DeviceEditable"];
            }
          } else {
            delete x["Manual"]["HTTPProxy"]["Port"];
          }
        } else {
          delete x["Manual"]["HTTPProxy"];
        }
        if (A.load.Bool(ptr + 32 + 68 + 66)) {
          x["Manual"]["SecureHTTPProxy"] = {};
          if (A.load.Bool(ptr + 32 + 68 + 0 + 28)) {
            x["Manual"]["SecureHTTPProxy"]["Host"] = {};
            x["Manual"]["SecureHTTPProxy"]["Host"]["Active"] = A.load.Ref(ptr + 32 + 68 + 0 + 0, undefined);
            x["Manual"]["SecureHTTPProxy"]["Host"]["Effective"] = A.load.Ref(ptr + 32 + 68 + 0 + 4, undefined);
            x["Manual"]["SecureHTTPProxy"]["Host"]["UserPolicy"] = A.load.Ref(ptr + 32 + 68 + 0 + 8, undefined);
            x["Manual"]["SecureHTTPProxy"]["Host"]["DevicePolicy"] = A.load.Ref(ptr + 32 + 68 + 0 + 12, undefined);
            x["Manual"]["SecureHTTPProxy"]["Host"]["UserSetting"] = A.load.Ref(ptr + 32 + 68 + 0 + 16, undefined);
            x["Manual"]["SecureHTTPProxy"]["Host"]["SharedSetting"] = A.load.Ref(ptr + 32 + 68 + 0 + 20, undefined);
            if (A.load.Bool(ptr + 32 + 68 + 0 + 26)) {
              x["Manual"]["SecureHTTPProxy"]["Host"]["UserEditable"] = A.load.Bool(ptr + 32 + 68 + 0 + 24);
            } else {
              delete x["Manual"]["SecureHTTPProxy"]["Host"]["UserEditable"];
            }
            if (A.load.Bool(ptr + 32 + 68 + 0 + 27)) {
              x["Manual"]["SecureHTTPProxy"]["Host"]["DeviceEditable"] = A.load.Bool(ptr + 32 + 68 + 0 + 25);
            } else {
              delete x["Manual"]["SecureHTTPProxy"]["Host"]["DeviceEditable"];
            }
          } else {
            delete x["Manual"]["SecureHTTPProxy"]["Host"];
          }
          if (A.load.Bool(ptr + 32 + 68 + 32 + 33)) {
            x["Manual"]["SecureHTTPProxy"]["Port"] = {};
            if (A.load.Bool(ptr + 32 + 68 + 32 + 26)) {
              x["Manual"]["SecureHTTPProxy"]["Port"]["Active"] = A.load.Int32(ptr + 32 + 68 + 32 + 0);
            } else {
              delete x["Manual"]["SecureHTTPProxy"]["Port"]["Active"];
            }
            x["Manual"]["SecureHTTPProxy"]["Port"]["Effective"] = A.load.Ref(ptr + 32 + 68 + 32 + 4, undefined);
            if (A.load.Bool(ptr + 32 + 68 + 32 + 27)) {
              x["Manual"]["SecureHTTPProxy"]["Port"]["UserPolicy"] = A.load.Int32(ptr + 32 + 68 + 32 + 8);
            } else {
              delete x["Manual"]["SecureHTTPProxy"]["Port"]["UserPolicy"];
            }
            if (A.load.Bool(ptr + 32 + 68 + 32 + 28)) {
              x["Manual"]["SecureHTTPProxy"]["Port"]["DevicePolicy"] = A.load.Int32(ptr + 32 + 68 + 32 + 12);
            } else {
              delete x["Manual"]["SecureHTTPProxy"]["Port"]["DevicePolicy"];
            }
            if (A.load.Bool(ptr + 32 + 68 + 32 + 29)) {
              x["Manual"]["SecureHTTPProxy"]["Port"]["UserSetting"] = A.load.Int32(ptr + 32 + 68 + 32 + 16);
            } else {
              delete x["Manual"]["SecureHTTPProxy"]["Port"]["UserSetting"];
            }
            if (A.load.Bool(ptr + 32 + 68 + 32 + 30)) {
              x["Manual"]["SecureHTTPProxy"]["Port"]["SharedSetting"] = A.load.Int32(ptr + 32 + 68 + 32 + 20);
            } else {
              delete x["Manual"]["SecureHTTPProxy"]["Port"]["SharedSetting"];
            }
            if (A.load.Bool(ptr + 32 + 68 + 32 + 31)) {
              x["Manual"]["SecureHTTPProxy"]["Port"]["UserEditable"] = A.load.Bool(ptr + 32 + 68 + 32 + 24);
            } else {
              delete x["Manual"]["SecureHTTPProxy"]["Port"]["UserEditable"];
            }
            if (A.load.Bool(ptr + 32 + 68 + 32 + 32)) {
              x["Manual"]["SecureHTTPProxy"]["Port"]["DeviceEditable"] = A.load.Bool(ptr + 32 + 68 + 32 + 25);
            } else {
              delete x["Manual"]["SecureHTTPProxy"]["Port"]["DeviceEditable"];
            }
          } else {
            delete x["Manual"]["SecureHTTPProxy"]["Port"];
          }
        } else {
          delete x["Manual"]["SecureHTTPProxy"];
        }
        if (A.load.Bool(ptr + 32 + 136 + 66)) {
          x["Manual"]["FTPProxy"] = {};
          if (A.load.Bool(ptr + 32 + 136 + 0 + 28)) {
            x["Manual"]["FTPProxy"]["Host"] = {};
            x["Manual"]["FTPProxy"]["Host"]["Active"] = A.load.Ref(ptr + 32 + 136 + 0 + 0, undefined);
            x["Manual"]["FTPProxy"]["Host"]["Effective"] = A.load.Ref(ptr + 32 + 136 + 0 + 4, undefined);
            x["Manual"]["FTPProxy"]["Host"]["UserPolicy"] = A.load.Ref(ptr + 32 + 136 + 0 + 8, undefined);
            x["Manual"]["FTPProxy"]["Host"]["DevicePolicy"] = A.load.Ref(ptr + 32 + 136 + 0 + 12, undefined);
            x["Manual"]["FTPProxy"]["Host"]["UserSetting"] = A.load.Ref(ptr + 32 + 136 + 0 + 16, undefined);
            x["Manual"]["FTPProxy"]["Host"]["SharedSetting"] = A.load.Ref(ptr + 32 + 136 + 0 + 20, undefined);
            if (A.load.Bool(ptr + 32 + 136 + 0 + 26)) {
              x["Manual"]["FTPProxy"]["Host"]["UserEditable"] = A.load.Bool(ptr + 32 + 136 + 0 + 24);
            } else {
              delete x["Manual"]["FTPProxy"]["Host"]["UserEditable"];
            }
            if (A.load.Bool(ptr + 32 + 136 + 0 + 27)) {
              x["Manual"]["FTPProxy"]["Host"]["DeviceEditable"] = A.load.Bool(ptr + 32 + 136 + 0 + 25);
            } else {
              delete x["Manual"]["FTPProxy"]["Host"]["DeviceEditable"];
            }
          } else {
            delete x["Manual"]["FTPProxy"]["Host"];
          }
          if (A.load.Bool(ptr + 32 + 136 + 32 + 33)) {
            x["Manual"]["FTPProxy"]["Port"] = {};
            if (A.load.Bool(ptr + 32 + 136 + 32 + 26)) {
              x["Manual"]["FTPProxy"]["Port"]["Active"] = A.load.Int32(ptr + 32 + 136 + 32 + 0);
            } else {
              delete x["Manual"]["FTPProxy"]["Port"]["Active"];
            }
            x["Manual"]["FTPProxy"]["Port"]["Effective"] = A.load.Ref(ptr + 32 + 136 + 32 + 4, undefined);
            if (A.load.Bool(ptr + 32 + 136 + 32 + 27)) {
              x["Manual"]["FTPProxy"]["Port"]["UserPolicy"] = A.load.Int32(ptr + 32 + 136 + 32 + 8);
            } else {
              delete x["Manual"]["FTPProxy"]["Port"]["UserPolicy"];
            }
            if (A.load.Bool(ptr + 32 + 136 + 32 + 28)) {
              x["Manual"]["FTPProxy"]["Port"]["DevicePolicy"] = A.load.Int32(ptr + 32 + 136 + 32 + 12);
            } else {
              delete x["Manual"]["FTPProxy"]["Port"]["DevicePolicy"];
            }
            if (A.load.Bool(ptr + 32 + 136 + 32 + 29)) {
              x["Manual"]["FTPProxy"]["Port"]["UserSetting"] = A.load.Int32(ptr + 32 + 136 + 32 + 16);
            } else {
              delete x["Manual"]["FTPProxy"]["Port"]["UserSetting"];
            }
            if (A.load.Bool(ptr + 32 + 136 + 32 + 30)) {
              x["Manual"]["FTPProxy"]["Port"]["SharedSetting"] = A.load.Int32(ptr + 32 + 136 + 32 + 20);
            } else {
              delete x["Manual"]["FTPProxy"]["Port"]["SharedSetting"];
            }
            if (A.load.Bool(ptr + 32 + 136 + 32 + 31)) {
              x["Manual"]["FTPProxy"]["Port"]["UserEditable"] = A.load.Bool(ptr + 32 + 136 + 32 + 24);
            } else {
              delete x["Manual"]["FTPProxy"]["Port"]["UserEditable"];
            }
            if (A.load.Bool(ptr + 32 + 136 + 32 + 32)) {
              x["Manual"]["FTPProxy"]["Port"]["DeviceEditable"] = A.load.Bool(ptr + 32 + 136 + 32 + 25);
            } else {
              delete x["Manual"]["FTPProxy"]["Port"]["DeviceEditable"];
            }
          } else {
            delete x["Manual"]["FTPProxy"]["Port"];
          }
        } else {
          delete x["Manual"]["FTPProxy"];
        }
        if (A.load.Bool(ptr + 32 + 204 + 66)) {
          x["Manual"]["SOCKS"] = {};
          if (A.load.Bool(ptr + 32 + 204 + 0 + 28)) {
            x["Manual"]["SOCKS"]["Host"] = {};
            x["Manual"]["SOCKS"]["Host"]["Active"] = A.load.Ref(ptr + 32 + 204 + 0 + 0, undefined);
            x["Manual"]["SOCKS"]["Host"]["Effective"] = A.load.Ref(ptr + 32 + 204 + 0 + 4, undefined);
            x["Manual"]["SOCKS"]["Host"]["UserPolicy"] = A.load.Ref(ptr + 32 + 204 + 0 + 8, undefined);
            x["Manual"]["SOCKS"]["Host"]["DevicePolicy"] = A.load.Ref(ptr + 32 + 204 + 0 + 12, undefined);
            x["Manual"]["SOCKS"]["Host"]["UserSetting"] = A.load.Ref(ptr + 32 + 204 + 0 + 16, undefined);
            x["Manual"]["SOCKS"]["Host"]["SharedSetting"] = A.load.Ref(ptr + 32 + 204 + 0 + 20, undefined);
            if (A.load.Bool(ptr + 32 + 204 + 0 + 26)) {
              x["Manual"]["SOCKS"]["Host"]["UserEditable"] = A.load.Bool(ptr + 32 + 204 + 0 + 24);
            } else {
              delete x["Manual"]["SOCKS"]["Host"]["UserEditable"];
            }
            if (A.load.Bool(ptr + 32 + 204 + 0 + 27)) {
              x["Manual"]["SOCKS"]["Host"]["DeviceEditable"] = A.load.Bool(ptr + 32 + 204 + 0 + 25);
            } else {
              delete x["Manual"]["SOCKS"]["Host"]["DeviceEditable"];
            }
          } else {
            delete x["Manual"]["SOCKS"]["Host"];
          }
          if (A.load.Bool(ptr + 32 + 204 + 32 + 33)) {
            x["Manual"]["SOCKS"]["Port"] = {};
            if (A.load.Bool(ptr + 32 + 204 + 32 + 26)) {
              x["Manual"]["SOCKS"]["Port"]["Active"] = A.load.Int32(ptr + 32 + 204 + 32 + 0);
            } else {
              delete x["Manual"]["SOCKS"]["Port"]["Active"];
            }
            x["Manual"]["SOCKS"]["Port"]["Effective"] = A.load.Ref(ptr + 32 + 204 + 32 + 4, undefined);
            if (A.load.Bool(ptr + 32 + 204 + 32 + 27)) {
              x["Manual"]["SOCKS"]["Port"]["UserPolicy"] = A.load.Int32(ptr + 32 + 204 + 32 + 8);
            } else {
              delete x["Manual"]["SOCKS"]["Port"]["UserPolicy"];
            }
            if (A.load.Bool(ptr + 32 + 204 + 32 + 28)) {
              x["Manual"]["SOCKS"]["Port"]["DevicePolicy"] = A.load.Int32(ptr + 32 + 204 + 32 + 12);
            } else {
              delete x["Manual"]["SOCKS"]["Port"]["DevicePolicy"];
            }
            if (A.load.Bool(ptr + 32 + 204 + 32 + 29)) {
              x["Manual"]["SOCKS"]["Port"]["UserSetting"] = A.load.Int32(ptr + 32 + 204 + 32 + 16);
            } else {
              delete x["Manual"]["SOCKS"]["Port"]["UserSetting"];
            }
            if (A.load.Bool(ptr + 32 + 204 + 32 + 30)) {
              x["Manual"]["SOCKS"]["Port"]["SharedSetting"] = A.load.Int32(ptr + 32 + 204 + 32 + 20);
            } else {
              delete x["Manual"]["SOCKS"]["Port"]["SharedSetting"];
            }
            if (A.load.Bool(ptr + 32 + 204 + 32 + 31)) {
              x["Manual"]["SOCKS"]["Port"]["UserEditable"] = A.load.Bool(ptr + 32 + 204 + 32 + 24);
            } else {
              delete x["Manual"]["SOCKS"]["Port"]["UserEditable"];
            }
            if (A.load.Bool(ptr + 32 + 204 + 32 + 32)) {
              x["Manual"]["SOCKS"]["Port"]["DeviceEditable"] = A.load.Bool(ptr + 32 + 204 + 32 + 25);
            } else {
              delete x["Manual"]["SOCKS"]["Port"]["DeviceEditable"];
            }
          } else {
            delete x["Manual"]["SOCKS"]["Port"];
          }
        } else {
          delete x["Manual"]["SOCKS"];
        }
      } else {
        delete x["Manual"];
      }
      if (A.load.Bool(ptr + 304 + 28)) {
        x["ExcludeDomains"] = {};
        x["ExcludeDomains"]["Active"] = A.load.Ref(ptr + 304 + 0, undefined);
        x["ExcludeDomains"]["Effective"] = A.load.Ref(ptr + 304 + 4, undefined);
        x["ExcludeDomains"]["UserPolicy"] = A.load.Ref(ptr + 304 + 8, undefined);
        x["ExcludeDomains"]["DevicePolicy"] = A.load.Ref(ptr + 304 + 12, undefined);
        x["ExcludeDomains"]["UserSetting"] = A.load.Ref(ptr + 304 + 16, undefined);
        x["ExcludeDomains"]["SharedSetting"] = A.load.Ref(ptr + 304 + 20, undefined);
        if (A.load.Bool(ptr + 304 + 26)) {
          x["ExcludeDomains"]["UserEditable"] = A.load.Bool(ptr + 304 + 24);
        } else {
          delete x["ExcludeDomains"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 304 + 27)) {
          x["ExcludeDomains"]["DeviceEditable"] = A.load.Bool(ptr + 304 + 25);
        } else {
          delete x["ExcludeDomains"]["DeviceEditable"];
        }
      } else {
        delete x["ExcludeDomains"];
      }
      if (A.load.Bool(ptr + 336 + 28)) {
        x["PAC"] = {};
        x["PAC"]["Active"] = A.load.Ref(ptr + 336 + 0, undefined);
        x["PAC"]["Effective"] = A.load.Ref(ptr + 336 + 4, undefined);
        x["PAC"]["UserPolicy"] = A.load.Ref(ptr + 336 + 8, undefined);
        x["PAC"]["DevicePolicy"] = A.load.Ref(ptr + 336 + 12, undefined);
        x["PAC"]["UserSetting"] = A.load.Ref(ptr + 336 + 16, undefined);
        x["PAC"]["SharedSetting"] = A.load.Ref(ptr + 336 + 20, undefined);
        if (A.load.Bool(ptr + 336 + 26)) {
          x["PAC"]["UserEditable"] = A.load.Bool(ptr + 336 + 24);
        } else {
          delete x["PAC"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 336 + 27)) {
          x["PAC"]["DeviceEditable"] = A.load.Bool(ptr + 336 + 25);
        } else {
          delete x["PAC"]["DeviceEditable"];
        }
      } else {
        delete x["PAC"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManagedIPConfigProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 193, false);

        A.store.Bool(ptr + 0 + 28, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Ref(ptr + 0 + 8, undefined);
        A.store.Ref(ptr + 0 + 12, undefined);
        A.store.Ref(ptr + 0 + 16, undefined);
        A.store.Ref(ptr + 0 + 20, undefined);
        A.store.Bool(ptr + 0 + 26, false);
        A.store.Bool(ptr + 0 + 24, false);
        A.store.Bool(ptr + 0 + 27, false);
        A.store.Bool(ptr + 0 + 25, false);

        A.store.Bool(ptr + 32 + 28, false);
        A.store.Ref(ptr + 32 + 0, undefined);
        A.store.Ref(ptr + 32 + 4, undefined);
        A.store.Ref(ptr + 32 + 8, undefined);
        A.store.Ref(ptr + 32 + 12, undefined);
        A.store.Ref(ptr + 32 + 16, undefined);
        A.store.Ref(ptr + 32 + 20, undefined);
        A.store.Bool(ptr + 32 + 26, false);
        A.store.Bool(ptr + 32 + 24, false);
        A.store.Bool(ptr + 32 + 27, false);
        A.store.Bool(ptr + 32 + 25, false);

        A.store.Bool(ptr + 64 + 28, false);
        A.store.Ref(ptr + 64 + 0, undefined);
        A.store.Ref(ptr + 64 + 4, undefined);
        A.store.Ref(ptr + 64 + 8, undefined);
        A.store.Ref(ptr + 64 + 12, undefined);
        A.store.Ref(ptr + 64 + 16, undefined);
        A.store.Ref(ptr + 64 + 20, undefined);
        A.store.Bool(ptr + 64 + 26, false);
        A.store.Bool(ptr + 64 + 24, false);
        A.store.Bool(ptr + 64 + 27, false);
        A.store.Bool(ptr + 64 + 25, false);

        A.store.Bool(ptr + 96 + 33, false);
        A.store.Bool(ptr + 96 + 26, false);
        A.store.Int32(ptr + 96 + 0, 0);
        A.store.Ref(ptr + 96 + 4, undefined);
        A.store.Bool(ptr + 96 + 27, false);
        A.store.Int32(ptr + 96 + 8, 0);
        A.store.Bool(ptr + 96 + 28, false);
        A.store.Int32(ptr + 96 + 12, 0);
        A.store.Bool(ptr + 96 + 29, false);
        A.store.Int32(ptr + 96 + 16, 0);
        A.store.Bool(ptr + 96 + 30, false);
        A.store.Int32(ptr + 96 + 20, 0);
        A.store.Bool(ptr + 96 + 31, false);
        A.store.Bool(ptr + 96 + 24, false);
        A.store.Bool(ptr + 96 + 32, false);
        A.store.Bool(ptr + 96 + 25, false);

        A.store.Bool(ptr + 132 + 28, false);
        A.store.Ref(ptr + 132 + 0, undefined);
        A.store.Ref(ptr + 132 + 4, undefined);
        A.store.Ref(ptr + 132 + 8, undefined);
        A.store.Ref(ptr + 132 + 12, undefined);
        A.store.Ref(ptr + 132 + 16, undefined);
        A.store.Ref(ptr + 132 + 20, undefined);
        A.store.Bool(ptr + 132 + 26, false);
        A.store.Bool(ptr + 132 + 24, false);
        A.store.Bool(ptr + 132 + 27, false);
        A.store.Bool(ptr + 132 + 25, false);

        A.store.Bool(ptr + 164 + 28, false);
        A.store.Ref(ptr + 164 + 0, undefined);
        A.store.Ref(ptr + 164 + 4, undefined);
        A.store.Ref(ptr + 164 + 8, undefined);
        A.store.Ref(ptr + 164 + 12, undefined);
        A.store.Ref(ptr + 164 + 16, undefined);
        A.store.Ref(ptr + 164 + 20, undefined);
        A.store.Bool(ptr + 164 + 26, false);
        A.store.Bool(ptr + 164 + 24, false);
        A.store.Bool(ptr + 164 + 27, false);
        A.store.Bool(ptr + 164 + 25, false);
      } else {
        A.store.Bool(ptr + 193, true);

        if (typeof x["Gateway"] === "undefined") {
          A.store.Bool(ptr + 0 + 28, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Ref(ptr + 0 + 8, undefined);
          A.store.Ref(ptr + 0 + 12, undefined);
          A.store.Ref(ptr + 0 + 16, undefined);
          A.store.Ref(ptr + 0 + 20, undefined);
          A.store.Bool(ptr + 0 + 26, false);
          A.store.Bool(ptr + 0 + 24, false);
          A.store.Bool(ptr + 0 + 27, false);
          A.store.Bool(ptr + 0 + 25, false);
        } else {
          A.store.Bool(ptr + 0 + 28, true);
          A.store.Ref(ptr + 0 + 0, x["Gateway"]["Active"]);
          A.store.Ref(ptr + 0 + 4, x["Gateway"]["Effective"]);
          A.store.Ref(ptr + 0 + 8, x["Gateway"]["UserPolicy"]);
          A.store.Ref(ptr + 0 + 12, x["Gateway"]["DevicePolicy"]);
          A.store.Ref(ptr + 0 + 16, x["Gateway"]["UserSetting"]);
          A.store.Ref(ptr + 0 + 20, x["Gateway"]["SharedSetting"]);
          A.store.Bool(ptr + 0 + 26, "UserEditable" in x["Gateway"] ? true : false);
          A.store.Bool(ptr + 0 + 24, x["Gateway"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 0 + 27, "DeviceEditable" in x["Gateway"] ? true : false);
          A.store.Bool(ptr + 0 + 25, x["Gateway"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["IPAddress"] === "undefined") {
          A.store.Bool(ptr + 32 + 28, false);
          A.store.Ref(ptr + 32 + 0, undefined);
          A.store.Ref(ptr + 32 + 4, undefined);
          A.store.Ref(ptr + 32 + 8, undefined);
          A.store.Ref(ptr + 32 + 12, undefined);
          A.store.Ref(ptr + 32 + 16, undefined);
          A.store.Ref(ptr + 32 + 20, undefined);
          A.store.Bool(ptr + 32 + 26, false);
          A.store.Bool(ptr + 32 + 24, false);
          A.store.Bool(ptr + 32 + 27, false);
          A.store.Bool(ptr + 32 + 25, false);
        } else {
          A.store.Bool(ptr + 32 + 28, true);
          A.store.Ref(ptr + 32 + 0, x["IPAddress"]["Active"]);
          A.store.Ref(ptr + 32 + 4, x["IPAddress"]["Effective"]);
          A.store.Ref(ptr + 32 + 8, x["IPAddress"]["UserPolicy"]);
          A.store.Ref(ptr + 32 + 12, x["IPAddress"]["DevicePolicy"]);
          A.store.Ref(ptr + 32 + 16, x["IPAddress"]["UserSetting"]);
          A.store.Ref(ptr + 32 + 20, x["IPAddress"]["SharedSetting"]);
          A.store.Bool(ptr + 32 + 26, "UserEditable" in x["IPAddress"] ? true : false);
          A.store.Bool(ptr + 32 + 24, x["IPAddress"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 32 + 27, "DeviceEditable" in x["IPAddress"] ? true : false);
          A.store.Bool(ptr + 32 + 25, x["IPAddress"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["NameServers"] === "undefined") {
          A.store.Bool(ptr + 64 + 28, false);
          A.store.Ref(ptr + 64 + 0, undefined);
          A.store.Ref(ptr + 64 + 4, undefined);
          A.store.Ref(ptr + 64 + 8, undefined);
          A.store.Ref(ptr + 64 + 12, undefined);
          A.store.Ref(ptr + 64 + 16, undefined);
          A.store.Ref(ptr + 64 + 20, undefined);
          A.store.Bool(ptr + 64 + 26, false);
          A.store.Bool(ptr + 64 + 24, false);
          A.store.Bool(ptr + 64 + 27, false);
          A.store.Bool(ptr + 64 + 25, false);
        } else {
          A.store.Bool(ptr + 64 + 28, true);
          A.store.Ref(ptr + 64 + 0, x["NameServers"]["Active"]);
          A.store.Ref(ptr + 64 + 4, x["NameServers"]["Effective"]);
          A.store.Ref(ptr + 64 + 8, x["NameServers"]["UserPolicy"]);
          A.store.Ref(ptr + 64 + 12, x["NameServers"]["DevicePolicy"]);
          A.store.Ref(ptr + 64 + 16, x["NameServers"]["UserSetting"]);
          A.store.Ref(ptr + 64 + 20, x["NameServers"]["SharedSetting"]);
          A.store.Bool(ptr + 64 + 26, "UserEditable" in x["NameServers"] ? true : false);
          A.store.Bool(ptr + 64 + 24, x["NameServers"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 64 + 27, "DeviceEditable" in x["NameServers"] ? true : false);
          A.store.Bool(ptr + 64 + 25, x["NameServers"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["RoutingPrefix"] === "undefined") {
          A.store.Bool(ptr + 96 + 33, false);
          A.store.Bool(ptr + 96 + 26, false);
          A.store.Int32(ptr + 96 + 0, 0);
          A.store.Ref(ptr + 96 + 4, undefined);
          A.store.Bool(ptr + 96 + 27, false);
          A.store.Int32(ptr + 96 + 8, 0);
          A.store.Bool(ptr + 96 + 28, false);
          A.store.Int32(ptr + 96 + 12, 0);
          A.store.Bool(ptr + 96 + 29, false);
          A.store.Int32(ptr + 96 + 16, 0);
          A.store.Bool(ptr + 96 + 30, false);
          A.store.Int32(ptr + 96 + 20, 0);
          A.store.Bool(ptr + 96 + 31, false);
          A.store.Bool(ptr + 96 + 24, false);
          A.store.Bool(ptr + 96 + 32, false);
          A.store.Bool(ptr + 96 + 25, false);
        } else {
          A.store.Bool(ptr + 96 + 33, true);
          A.store.Bool(ptr + 96 + 26, "Active" in x["RoutingPrefix"] ? true : false);
          A.store.Int32(
            ptr + 96 + 0,
            x["RoutingPrefix"]["Active"] === undefined ? 0 : (x["RoutingPrefix"]["Active"] as number)
          );
          A.store.Ref(ptr + 96 + 4, x["RoutingPrefix"]["Effective"]);
          A.store.Bool(ptr + 96 + 27, "UserPolicy" in x["RoutingPrefix"] ? true : false);
          A.store.Int32(
            ptr + 96 + 8,
            x["RoutingPrefix"]["UserPolicy"] === undefined ? 0 : (x["RoutingPrefix"]["UserPolicy"] as number)
          );
          A.store.Bool(ptr + 96 + 28, "DevicePolicy" in x["RoutingPrefix"] ? true : false);
          A.store.Int32(
            ptr + 96 + 12,
            x["RoutingPrefix"]["DevicePolicy"] === undefined ? 0 : (x["RoutingPrefix"]["DevicePolicy"] as number)
          );
          A.store.Bool(ptr + 96 + 29, "UserSetting" in x["RoutingPrefix"] ? true : false);
          A.store.Int32(
            ptr + 96 + 16,
            x["RoutingPrefix"]["UserSetting"] === undefined ? 0 : (x["RoutingPrefix"]["UserSetting"] as number)
          );
          A.store.Bool(ptr + 96 + 30, "SharedSetting" in x["RoutingPrefix"] ? true : false);
          A.store.Int32(
            ptr + 96 + 20,
            x["RoutingPrefix"]["SharedSetting"] === undefined ? 0 : (x["RoutingPrefix"]["SharedSetting"] as number)
          );
          A.store.Bool(ptr + 96 + 31, "UserEditable" in x["RoutingPrefix"] ? true : false);
          A.store.Bool(ptr + 96 + 24, x["RoutingPrefix"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 96 + 32, "DeviceEditable" in x["RoutingPrefix"] ? true : false);
          A.store.Bool(ptr + 96 + 25, x["RoutingPrefix"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["Type"] === "undefined") {
          A.store.Bool(ptr + 132 + 28, false);
          A.store.Ref(ptr + 132 + 0, undefined);
          A.store.Ref(ptr + 132 + 4, undefined);
          A.store.Ref(ptr + 132 + 8, undefined);
          A.store.Ref(ptr + 132 + 12, undefined);
          A.store.Ref(ptr + 132 + 16, undefined);
          A.store.Ref(ptr + 132 + 20, undefined);
          A.store.Bool(ptr + 132 + 26, false);
          A.store.Bool(ptr + 132 + 24, false);
          A.store.Bool(ptr + 132 + 27, false);
          A.store.Bool(ptr + 132 + 25, false);
        } else {
          A.store.Bool(ptr + 132 + 28, true);
          A.store.Ref(ptr + 132 + 0, x["Type"]["Active"]);
          A.store.Ref(ptr + 132 + 4, x["Type"]["Effective"]);
          A.store.Ref(ptr + 132 + 8, x["Type"]["UserPolicy"]);
          A.store.Ref(ptr + 132 + 12, x["Type"]["DevicePolicy"]);
          A.store.Ref(ptr + 132 + 16, x["Type"]["UserSetting"]);
          A.store.Ref(ptr + 132 + 20, x["Type"]["SharedSetting"]);
          A.store.Bool(ptr + 132 + 26, "UserEditable" in x["Type"] ? true : false);
          A.store.Bool(ptr + 132 + 24, x["Type"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 132 + 27, "DeviceEditable" in x["Type"] ? true : false);
          A.store.Bool(ptr + 132 + 25, x["Type"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["WebProxyAutoDiscoveryUrl"] === "undefined") {
          A.store.Bool(ptr + 164 + 28, false);
          A.store.Ref(ptr + 164 + 0, undefined);
          A.store.Ref(ptr + 164 + 4, undefined);
          A.store.Ref(ptr + 164 + 8, undefined);
          A.store.Ref(ptr + 164 + 12, undefined);
          A.store.Ref(ptr + 164 + 16, undefined);
          A.store.Ref(ptr + 164 + 20, undefined);
          A.store.Bool(ptr + 164 + 26, false);
          A.store.Bool(ptr + 164 + 24, false);
          A.store.Bool(ptr + 164 + 27, false);
          A.store.Bool(ptr + 164 + 25, false);
        } else {
          A.store.Bool(ptr + 164 + 28, true);
          A.store.Ref(ptr + 164 + 0, x["WebProxyAutoDiscoveryUrl"]["Active"]);
          A.store.Ref(ptr + 164 + 4, x["WebProxyAutoDiscoveryUrl"]["Effective"]);
          A.store.Ref(ptr + 164 + 8, x["WebProxyAutoDiscoveryUrl"]["UserPolicy"]);
          A.store.Ref(ptr + 164 + 12, x["WebProxyAutoDiscoveryUrl"]["DevicePolicy"]);
          A.store.Ref(ptr + 164 + 16, x["WebProxyAutoDiscoveryUrl"]["UserSetting"]);
          A.store.Ref(ptr + 164 + 20, x["WebProxyAutoDiscoveryUrl"]["SharedSetting"]);
          A.store.Bool(ptr + 164 + 26, "UserEditable" in x["WebProxyAutoDiscoveryUrl"] ? true : false);
          A.store.Bool(ptr + 164 + 24, x["WebProxyAutoDiscoveryUrl"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 164 + 27, "DeviceEditable" in x["WebProxyAutoDiscoveryUrl"] ? true : false);
          A.store.Bool(ptr + 164 + 25, x["WebProxyAutoDiscoveryUrl"]["DeviceEditable"] ? true : false);
        }
      }
    },
    "load_ManagedIPConfigProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 28)) {
        x["Gateway"] = {};
        x["Gateway"]["Active"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["Gateway"]["Effective"] = A.load.Ref(ptr + 0 + 4, undefined);
        x["Gateway"]["UserPolicy"] = A.load.Ref(ptr + 0 + 8, undefined);
        x["Gateway"]["DevicePolicy"] = A.load.Ref(ptr + 0 + 12, undefined);
        x["Gateway"]["UserSetting"] = A.load.Ref(ptr + 0 + 16, undefined);
        x["Gateway"]["SharedSetting"] = A.load.Ref(ptr + 0 + 20, undefined);
        if (A.load.Bool(ptr + 0 + 26)) {
          x["Gateway"]["UserEditable"] = A.load.Bool(ptr + 0 + 24);
        } else {
          delete x["Gateway"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 0 + 27)) {
          x["Gateway"]["DeviceEditable"] = A.load.Bool(ptr + 0 + 25);
        } else {
          delete x["Gateway"]["DeviceEditable"];
        }
      } else {
        delete x["Gateway"];
      }
      if (A.load.Bool(ptr + 32 + 28)) {
        x["IPAddress"] = {};
        x["IPAddress"]["Active"] = A.load.Ref(ptr + 32 + 0, undefined);
        x["IPAddress"]["Effective"] = A.load.Ref(ptr + 32 + 4, undefined);
        x["IPAddress"]["UserPolicy"] = A.load.Ref(ptr + 32 + 8, undefined);
        x["IPAddress"]["DevicePolicy"] = A.load.Ref(ptr + 32 + 12, undefined);
        x["IPAddress"]["UserSetting"] = A.load.Ref(ptr + 32 + 16, undefined);
        x["IPAddress"]["SharedSetting"] = A.load.Ref(ptr + 32 + 20, undefined);
        if (A.load.Bool(ptr + 32 + 26)) {
          x["IPAddress"]["UserEditable"] = A.load.Bool(ptr + 32 + 24);
        } else {
          delete x["IPAddress"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 32 + 27)) {
          x["IPAddress"]["DeviceEditable"] = A.load.Bool(ptr + 32 + 25);
        } else {
          delete x["IPAddress"]["DeviceEditable"];
        }
      } else {
        delete x["IPAddress"];
      }
      if (A.load.Bool(ptr + 64 + 28)) {
        x["NameServers"] = {};
        x["NameServers"]["Active"] = A.load.Ref(ptr + 64 + 0, undefined);
        x["NameServers"]["Effective"] = A.load.Ref(ptr + 64 + 4, undefined);
        x["NameServers"]["UserPolicy"] = A.load.Ref(ptr + 64 + 8, undefined);
        x["NameServers"]["DevicePolicy"] = A.load.Ref(ptr + 64 + 12, undefined);
        x["NameServers"]["UserSetting"] = A.load.Ref(ptr + 64 + 16, undefined);
        x["NameServers"]["SharedSetting"] = A.load.Ref(ptr + 64 + 20, undefined);
        if (A.load.Bool(ptr + 64 + 26)) {
          x["NameServers"]["UserEditable"] = A.load.Bool(ptr + 64 + 24);
        } else {
          delete x["NameServers"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 64 + 27)) {
          x["NameServers"]["DeviceEditable"] = A.load.Bool(ptr + 64 + 25);
        } else {
          delete x["NameServers"]["DeviceEditable"];
        }
      } else {
        delete x["NameServers"];
      }
      if (A.load.Bool(ptr + 96 + 33)) {
        x["RoutingPrefix"] = {};
        if (A.load.Bool(ptr + 96 + 26)) {
          x["RoutingPrefix"]["Active"] = A.load.Int32(ptr + 96 + 0);
        } else {
          delete x["RoutingPrefix"]["Active"];
        }
        x["RoutingPrefix"]["Effective"] = A.load.Ref(ptr + 96 + 4, undefined);
        if (A.load.Bool(ptr + 96 + 27)) {
          x["RoutingPrefix"]["UserPolicy"] = A.load.Int32(ptr + 96 + 8);
        } else {
          delete x["RoutingPrefix"]["UserPolicy"];
        }
        if (A.load.Bool(ptr + 96 + 28)) {
          x["RoutingPrefix"]["DevicePolicy"] = A.load.Int32(ptr + 96 + 12);
        } else {
          delete x["RoutingPrefix"]["DevicePolicy"];
        }
        if (A.load.Bool(ptr + 96 + 29)) {
          x["RoutingPrefix"]["UserSetting"] = A.load.Int32(ptr + 96 + 16);
        } else {
          delete x["RoutingPrefix"]["UserSetting"];
        }
        if (A.load.Bool(ptr + 96 + 30)) {
          x["RoutingPrefix"]["SharedSetting"] = A.load.Int32(ptr + 96 + 20);
        } else {
          delete x["RoutingPrefix"]["SharedSetting"];
        }
        if (A.load.Bool(ptr + 96 + 31)) {
          x["RoutingPrefix"]["UserEditable"] = A.load.Bool(ptr + 96 + 24);
        } else {
          delete x["RoutingPrefix"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 96 + 32)) {
          x["RoutingPrefix"]["DeviceEditable"] = A.load.Bool(ptr + 96 + 25);
        } else {
          delete x["RoutingPrefix"]["DeviceEditable"];
        }
      } else {
        delete x["RoutingPrefix"];
      }
      if (A.load.Bool(ptr + 132 + 28)) {
        x["Type"] = {};
        x["Type"]["Active"] = A.load.Ref(ptr + 132 + 0, undefined);
        x["Type"]["Effective"] = A.load.Ref(ptr + 132 + 4, undefined);
        x["Type"]["UserPolicy"] = A.load.Ref(ptr + 132 + 8, undefined);
        x["Type"]["DevicePolicy"] = A.load.Ref(ptr + 132 + 12, undefined);
        x["Type"]["UserSetting"] = A.load.Ref(ptr + 132 + 16, undefined);
        x["Type"]["SharedSetting"] = A.load.Ref(ptr + 132 + 20, undefined);
        if (A.load.Bool(ptr + 132 + 26)) {
          x["Type"]["UserEditable"] = A.load.Bool(ptr + 132 + 24);
        } else {
          delete x["Type"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 132 + 27)) {
          x["Type"]["DeviceEditable"] = A.load.Bool(ptr + 132 + 25);
        } else {
          delete x["Type"]["DeviceEditable"];
        }
      } else {
        delete x["Type"];
      }
      if (A.load.Bool(ptr + 164 + 28)) {
        x["WebProxyAutoDiscoveryUrl"] = {};
        x["WebProxyAutoDiscoveryUrl"]["Active"] = A.load.Ref(ptr + 164 + 0, undefined);
        x["WebProxyAutoDiscoveryUrl"]["Effective"] = A.load.Ref(ptr + 164 + 4, undefined);
        x["WebProxyAutoDiscoveryUrl"]["UserPolicy"] = A.load.Ref(ptr + 164 + 8, undefined);
        x["WebProxyAutoDiscoveryUrl"]["DevicePolicy"] = A.load.Ref(ptr + 164 + 12, undefined);
        x["WebProxyAutoDiscoveryUrl"]["UserSetting"] = A.load.Ref(ptr + 164 + 16, undefined);
        x["WebProxyAutoDiscoveryUrl"]["SharedSetting"] = A.load.Ref(ptr + 164 + 20, undefined);
        if (A.load.Bool(ptr + 164 + 26)) {
          x["WebProxyAutoDiscoveryUrl"]["UserEditable"] = A.load.Bool(ptr + 164 + 24);
        } else {
          delete x["WebProxyAutoDiscoveryUrl"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 164 + 27)) {
          x["WebProxyAutoDiscoveryUrl"]["DeviceEditable"] = A.load.Bool(ptr + 164 + 25);
        } else {
          delete x["WebProxyAutoDiscoveryUrl"]["DeviceEditable"];
        }
      } else {
        delete x["WebProxyAutoDiscoveryUrl"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManagedVPNProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 85, false);

        A.store.Bool(ptr + 0 + 21, false);
        A.store.Bool(ptr + 0 + 14, false);
        A.store.Bool(ptr + 0 + 0, false);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Bool(ptr + 0 + 15, false);
        A.store.Bool(ptr + 0 + 8, false);
        A.store.Bool(ptr + 0 + 16, false);
        A.store.Bool(ptr + 0 + 9, false);
        A.store.Bool(ptr + 0 + 17, false);
        A.store.Bool(ptr + 0 + 10, false);
        A.store.Bool(ptr + 0 + 18, false);
        A.store.Bool(ptr + 0 + 11, false);
        A.store.Bool(ptr + 0 + 19, false);
        A.store.Bool(ptr + 0 + 12, false);
        A.store.Bool(ptr + 0 + 20, false);
        A.store.Bool(ptr + 0 + 13, false);

        A.store.Bool(ptr + 24 + 28, false);
        A.store.Ref(ptr + 24 + 0, undefined);
        A.store.Ref(ptr + 24 + 4, undefined);
        A.store.Ref(ptr + 24 + 8, undefined);
        A.store.Ref(ptr + 24 + 12, undefined);
        A.store.Ref(ptr + 24 + 16, undefined);
        A.store.Ref(ptr + 24 + 20, undefined);
        A.store.Bool(ptr + 24 + 26, false);
        A.store.Bool(ptr + 24 + 24, false);
        A.store.Bool(ptr + 24 + 27, false);
        A.store.Bool(ptr + 24 + 25, false);

        A.store.Bool(ptr + 56 + 28, false);
        A.store.Ref(ptr + 56 + 0, undefined);
        A.store.Ref(ptr + 56 + 4, undefined);
        A.store.Ref(ptr + 56 + 8, undefined);
        A.store.Ref(ptr + 56 + 12, undefined);
        A.store.Ref(ptr + 56 + 16, undefined);
        A.store.Ref(ptr + 56 + 20, undefined);
        A.store.Bool(ptr + 56 + 26, false);
        A.store.Bool(ptr + 56 + 24, false);
        A.store.Bool(ptr + 56 + 27, false);
        A.store.Bool(ptr + 56 + 25, false);
      } else {
        A.store.Bool(ptr + 85, true);

        if (typeof x["AutoConnect"] === "undefined") {
          A.store.Bool(ptr + 0 + 21, false);
          A.store.Bool(ptr + 0 + 14, false);
          A.store.Bool(ptr + 0 + 0, false);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Bool(ptr + 0 + 15, false);
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Bool(ptr + 0 + 16, false);
          A.store.Bool(ptr + 0 + 9, false);
          A.store.Bool(ptr + 0 + 17, false);
          A.store.Bool(ptr + 0 + 10, false);
          A.store.Bool(ptr + 0 + 18, false);
          A.store.Bool(ptr + 0 + 11, false);
          A.store.Bool(ptr + 0 + 19, false);
          A.store.Bool(ptr + 0 + 12, false);
          A.store.Bool(ptr + 0 + 20, false);
          A.store.Bool(ptr + 0 + 13, false);
        } else {
          A.store.Bool(ptr + 0 + 21, true);
          A.store.Bool(ptr + 0 + 14, "Active" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 0, x["AutoConnect"]["Active"] ? true : false);
          A.store.Ref(ptr + 0 + 4, x["AutoConnect"]["Effective"]);
          A.store.Bool(ptr + 0 + 15, "UserPolicy" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 8, x["AutoConnect"]["UserPolicy"] ? true : false);
          A.store.Bool(ptr + 0 + 16, "DevicePolicy" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 9, x["AutoConnect"]["DevicePolicy"] ? true : false);
          A.store.Bool(ptr + 0 + 17, "UserSetting" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 10, x["AutoConnect"]["UserSetting"] ? true : false);
          A.store.Bool(ptr + 0 + 18, "SharedSetting" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 11, x["AutoConnect"]["SharedSetting"] ? true : false);
          A.store.Bool(ptr + 0 + 19, "UserEditable" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 12, x["AutoConnect"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 0 + 20, "DeviceEditable" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 0 + 13, x["AutoConnect"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["Host"] === "undefined") {
          A.store.Bool(ptr + 24 + 28, false);
          A.store.Ref(ptr + 24 + 0, undefined);
          A.store.Ref(ptr + 24 + 4, undefined);
          A.store.Ref(ptr + 24 + 8, undefined);
          A.store.Ref(ptr + 24 + 12, undefined);
          A.store.Ref(ptr + 24 + 16, undefined);
          A.store.Ref(ptr + 24 + 20, undefined);
          A.store.Bool(ptr + 24 + 26, false);
          A.store.Bool(ptr + 24 + 24, false);
          A.store.Bool(ptr + 24 + 27, false);
          A.store.Bool(ptr + 24 + 25, false);
        } else {
          A.store.Bool(ptr + 24 + 28, true);
          A.store.Ref(ptr + 24 + 0, x["Host"]["Active"]);
          A.store.Ref(ptr + 24 + 4, x["Host"]["Effective"]);
          A.store.Ref(ptr + 24 + 8, x["Host"]["UserPolicy"]);
          A.store.Ref(ptr + 24 + 12, x["Host"]["DevicePolicy"]);
          A.store.Ref(ptr + 24 + 16, x["Host"]["UserSetting"]);
          A.store.Ref(ptr + 24 + 20, x["Host"]["SharedSetting"]);
          A.store.Bool(ptr + 24 + 26, "UserEditable" in x["Host"] ? true : false);
          A.store.Bool(ptr + 24 + 24, x["Host"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 24 + 27, "DeviceEditable" in x["Host"] ? true : false);
          A.store.Bool(ptr + 24 + 25, x["Host"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["Type"] === "undefined") {
          A.store.Bool(ptr + 56 + 28, false);
          A.store.Ref(ptr + 56 + 0, undefined);
          A.store.Ref(ptr + 56 + 4, undefined);
          A.store.Ref(ptr + 56 + 8, undefined);
          A.store.Ref(ptr + 56 + 12, undefined);
          A.store.Ref(ptr + 56 + 16, undefined);
          A.store.Ref(ptr + 56 + 20, undefined);
          A.store.Bool(ptr + 56 + 26, false);
          A.store.Bool(ptr + 56 + 24, false);
          A.store.Bool(ptr + 56 + 27, false);
          A.store.Bool(ptr + 56 + 25, false);
        } else {
          A.store.Bool(ptr + 56 + 28, true);
          A.store.Ref(ptr + 56 + 0, x["Type"]["Active"]);
          A.store.Ref(ptr + 56 + 4, x["Type"]["Effective"]);
          A.store.Ref(ptr + 56 + 8, x["Type"]["UserPolicy"]);
          A.store.Ref(ptr + 56 + 12, x["Type"]["DevicePolicy"]);
          A.store.Ref(ptr + 56 + 16, x["Type"]["UserSetting"]);
          A.store.Ref(ptr + 56 + 20, x["Type"]["SharedSetting"]);
          A.store.Bool(ptr + 56 + 26, "UserEditable" in x["Type"] ? true : false);
          A.store.Bool(ptr + 56 + 24, x["Type"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 56 + 27, "DeviceEditable" in x["Type"] ? true : false);
          A.store.Bool(ptr + 56 + 25, x["Type"]["DeviceEditable"] ? true : false);
        }
      }
    },
    "load_ManagedVPNProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 21)) {
        x["AutoConnect"] = {};
        if (A.load.Bool(ptr + 0 + 14)) {
          x["AutoConnect"]["Active"] = A.load.Bool(ptr + 0 + 0);
        } else {
          delete x["AutoConnect"]["Active"];
        }
        x["AutoConnect"]["Effective"] = A.load.Ref(ptr + 0 + 4, undefined);
        if (A.load.Bool(ptr + 0 + 15)) {
          x["AutoConnect"]["UserPolicy"] = A.load.Bool(ptr + 0 + 8);
        } else {
          delete x["AutoConnect"]["UserPolicy"];
        }
        if (A.load.Bool(ptr + 0 + 16)) {
          x["AutoConnect"]["DevicePolicy"] = A.load.Bool(ptr + 0 + 9);
        } else {
          delete x["AutoConnect"]["DevicePolicy"];
        }
        if (A.load.Bool(ptr + 0 + 17)) {
          x["AutoConnect"]["UserSetting"] = A.load.Bool(ptr + 0 + 10);
        } else {
          delete x["AutoConnect"]["UserSetting"];
        }
        if (A.load.Bool(ptr + 0 + 18)) {
          x["AutoConnect"]["SharedSetting"] = A.load.Bool(ptr + 0 + 11);
        } else {
          delete x["AutoConnect"]["SharedSetting"];
        }
        if (A.load.Bool(ptr + 0 + 19)) {
          x["AutoConnect"]["UserEditable"] = A.load.Bool(ptr + 0 + 12);
        } else {
          delete x["AutoConnect"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 0 + 20)) {
          x["AutoConnect"]["DeviceEditable"] = A.load.Bool(ptr + 0 + 13);
        } else {
          delete x["AutoConnect"]["DeviceEditable"];
        }
      } else {
        delete x["AutoConnect"];
      }
      if (A.load.Bool(ptr + 24 + 28)) {
        x["Host"] = {};
        x["Host"]["Active"] = A.load.Ref(ptr + 24 + 0, undefined);
        x["Host"]["Effective"] = A.load.Ref(ptr + 24 + 4, undefined);
        x["Host"]["UserPolicy"] = A.load.Ref(ptr + 24 + 8, undefined);
        x["Host"]["DevicePolicy"] = A.load.Ref(ptr + 24 + 12, undefined);
        x["Host"]["UserSetting"] = A.load.Ref(ptr + 24 + 16, undefined);
        x["Host"]["SharedSetting"] = A.load.Ref(ptr + 24 + 20, undefined);
        if (A.load.Bool(ptr + 24 + 26)) {
          x["Host"]["UserEditable"] = A.load.Bool(ptr + 24 + 24);
        } else {
          delete x["Host"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 24 + 27)) {
          x["Host"]["DeviceEditable"] = A.load.Bool(ptr + 24 + 25);
        } else {
          delete x["Host"]["DeviceEditable"];
        }
      } else {
        delete x["Host"];
      }
      if (A.load.Bool(ptr + 56 + 28)) {
        x["Type"] = {};
        x["Type"]["Active"] = A.load.Ref(ptr + 56 + 0, undefined);
        x["Type"]["Effective"] = A.load.Ref(ptr + 56 + 4, undefined);
        x["Type"]["UserPolicy"] = A.load.Ref(ptr + 56 + 8, undefined);
        x["Type"]["DevicePolicy"] = A.load.Ref(ptr + 56 + 12, undefined);
        x["Type"]["UserSetting"] = A.load.Ref(ptr + 56 + 16, undefined);
        x["Type"]["SharedSetting"] = A.load.Ref(ptr + 56 + 20, undefined);
        if (A.load.Bool(ptr + 56 + 26)) {
          x["Type"]["UserEditable"] = A.load.Bool(ptr + 56 + 24);
        } else {
          delete x["Type"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 56 + 27)) {
          x["Type"]["DeviceEditable"] = A.load.Bool(ptr + 56 + 25);
        } else {
          delete x["Type"]["DeviceEditable"];
        }
      } else {
        delete x["Type"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManagedWiFiProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 222, false);

        A.store.Bool(ptr + 0 + 21, false);
        A.store.Bool(ptr + 0 + 14, false);
        A.store.Bool(ptr + 0 + 0, false);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Bool(ptr + 0 + 15, false);
        A.store.Bool(ptr + 0 + 8, false);
        A.store.Bool(ptr + 0 + 16, false);
        A.store.Bool(ptr + 0 + 9, false);
        A.store.Bool(ptr + 0 + 17, false);
        A.store.Bool(ptr + 0 + 10, false);
        A.store.Bool(ptr + 0 + 18, false);
        A.store.Bool(ptr + 0 + 11, false);
        A.store.Bool(ptr + 0 + 19, false);
        A.store.Bool(ptr + 0 + 12, false);
        A.store.Bool(ptr + 0 + 20, false);
        A.store.Bool(ptr + 0 + 13, false);

        A.store.Bool(ptr + 24 + 21, false);
        A.store.Bool(ptr + 24 + 14, false);
        A.store.Bool(ptr + 24 + 0, false);
        A.store.Ref(ptr + 24 + 4, undefined);
        A.store.Bool(ptr + 24 + 15, false);
        A.store.Bool(ptr + 24 + 8, false);
        A.store.Bool(ptr + 24 + 16, false);
        A.store.Bool(ptr + 24 + 9, false);
        A.store.Bool(ptr + 24 + 17, false);
        A.store.Bool(ptr + 24 + 10, false);
        A.store.Bool(ptr + 24 + 18, false);
        A.store.Bool(ptr + 24 + 11, false);
        A.store.Bool(ptr + 24 + 19, false);
        A.store.Bool(ptr + 24 + 12, false);
        A.store.Bool(ptr + 24 + 20, false);
        A.store.Bool(ptr + 24 + 13, false);
        A.store.Ref(ptr + 48, undefined);
        A.store.Bool(ptr + 220, false);
        A.store.Int32(ptr + 52, 0);
        A.store.Ref(ptr + 56, undefined);

        A.store.Bool(ptr + 60 + 28, false);
        A.store.Ref(ptr + 60 + 0, undefined);
        A.store.Ref(ptr + 60 + 4, undefined);
        A.store.Ref(ptr + 60 + 8, undefined);
        A.store.Ref(ptr + 60 + 12, undefined);
        A.store.Ref(ptr + 60 + 16, undefined);
        A.store.Ref(ptr + 60 + 20, undefined);
        A.store.Bool(ptr + 60 + 26, false);
        A.store.Bool(ptr + 60 + 24, false);
        A.store.Bool(ptr + 60 + 27, false);
        A.store.Bool(ptr + 60 + 25, false);

        A.store.Bool(ptr + 92 + 21, false);
        A.store.Bool(ptr + 92 + 14, false);
        A.store.Bool(ptr + 92 + 0, false);
        A.store.Ref(ptr + 92 + 4, undefined);
        A.store.Bool(ptr + 92 + 15, false);
        A.store.Bool(ptr + 92 + 8, false);
        A.store.Bool(ptr + 92 + 16, false);
        A.store.Bool(ptr + 92 + 9, false);
        A.store.Bool(ptr + 92 + 17, false);
        A.store.Bool(ptr + 92 + 10, false);
        A.store.Bool(ptr + 92 + 18, false);
        A.store.Bool(ptr + 92 + 11, false);
        A.store.Bool(ptr + 92 + 19, false);
        A.store.Bool(ptr + 92 + 12, false);
        A.store.Bool(ptr + 92 + 20, false);
        A.store.Bool(ptr + 92 + 13, false);

        A.store.Bool(ptr + 116 + 33, false);
        A.store.Bool(ptr + 116 + 26, false);
        A.store.Int32(ptr + 116 + 0, 0);
        A.store.Ref(ptr + 116 + 4, undefined);
        A.store.Bool(ptr + 116 + 27, false);
        A.store.Int32(ptr + 116 + 8, 0);
        A.store.Bool(ptr + 116 + 28, false);
        A.store.Int32(ptr + 116 + 12, 0);
        A.store.Bool(ptr + 116 + 29, false);
        A.store.Int32(ptr + 116 + 16, 0);
        A.store.Bool(ptr + 116 + 30, false);
        A.store.Int32(ptr + 116 + 20, 0);
        A.store.Bool(ptr + 116 + 31, false);
        A.store.Bool(ptr + 116 + 24, false);
        A.store.Bool(ptr + 116 + 32, false);
        A.store.Bool(ptr + 116 + 25, false);

        A.store.Bool(ptr + 152 + 28, false);
        A.store.Ref(ptr + 152 + 0, undefined);
        A.store.Ref(ptr + 152 + 4, undefined);
        A.store.Ref(ptr + 152 + 8, undefined);
        A.store.Ref(ptr + 152 + 12, undefined);
        A.store.Ref(ptr + 152 + 16, undefined);
        A.store.Ref(ptr + 152 + 20, undefined);
        A.store.Bool(ptr + 152 + 26, false);
        A.store.Bool(ptr + 152 + 24, false);
        A.store.Bool(ptr + 152 + 27, false);
        A.store.Bool(ptr + 152 + 25, false);

        A.store.Bool(ptr + 184 + 28, false);
        A.store.Ref(ptr + 184 + 0, undefined);
        A.store.Ref(ptr + 184 + 4, undefined);
        A.store.Ref(ptr + 184 + 8, undefined);
        A.store.Ref(ptr + 184 + 12, undefined);
        A.store.Ref(ptr + 184 + 16, undefined);
        A.store.Ref(ptr + 184 + 20, undefined);
        A.store.Bool(ptr + 184 + 26, false);
        A.store.Bool(ptr + 184 + 24, false);
        A.store.Bool(ptr + 184 + 27, false);
        A.store.Bool(ptr + 184 + 25, false);
        A.store.Bool(ptr + 221, false);
        A.store.Int32(ptr + 216, 0);
      } else {
        A.store.Bool(ptr + 222, true);

        if (typeof x["AllowGatewayARPPolling"] === "undefined") {
          A.store.Bool(ptr + 0 + 21, false);
          A.store.Bool(ptr + 0 + 14, false);
          A.store.Bool(ptr + 0 + 0, false);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Bool(ptr + 0 + 15, false);
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Bool(ptr + 0 + 16, false);
          A.store.Bool(ptr + 0 + 9, false);
          A.store.Bool(ptr + 0 + 17, false);
          A.store.Bool(ptr + 0 + 10, false);
          A.store.Bool(ptr + 0 + 18, false);
          A.store.Bool(ptr + 0 + 11, false);
          A.store.Bool(ptr + 0 + 19, false);
          A.store.Bool(ptr + 0 + 12, false);
          A.store.Bool(ptr + 0 + 20, false);
          A.store.Bool(ptr + 0 + 13, false);
        } else {
          A.store.Bool(ptr + 0 + 21, true);
          A.store.Bool(ptr + 0 + 14, "Active" in x["AllowGatewayARPPolling"] ? true : false);
          A.store.Bool(ptr + 0 + 0, x["AllowGatewayARPPolling"]["Active"] ? true : false);
          A.store.Ref(ptr + 0 + 4, x["AllowGatewayARPPolling"]["Effective"]);
          A.store.Bool(ptr + 0 + 15, "UserPolicy" in x["AllowGatewayARPPolling"] ? true : false);
          A.store.Bool(ptr + 0 + 8, x["AllowGatewayARPPolling"]["UserPolicy"] ? true : false);
          A.store.Bool(ptr + 0 + 16, "DevicePolicy" in x["AllowGatewayARPPolling"] ? true : false);
          A.store.Bool(ptr + 0 + 9, x["AllowGatewayARPPolling"]["DevicePolicy"] ? true : false);
          A.store.Bool(ptr + 0 + 17, "UserSetting" in x["AllowGatewayARPPolling"] ? true : false);
          A.store.Bool(ptr + 0 + 10, x["AllowGatewayARPPolling"]["UserSetting"] ? true : false);
          A.store.Bool(ptr + 0 + 18, "SharedSetting" in x["AllowGatewayARPPolling"] ? true : false);
          A.store.Bool(ptr + 0 + 11, x["AllowGatewayARPPolling"]["SharedSetting"] ? true : false);
          A.store.Bool(ptr + 0 + 19, "UserEditable" in x["AllowGatewayARPPolling"] ? true : false);
          A.store.Bool(ptr + 0 + 12, x["AllowGatewayARPPolling"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 0 + 20, "DeviceEditable" in x["AllowGatewayARPPolling"] ? true : false);
          A.store.Bool(ptr + 0 + 13, x["AllowGatewayARPPolling"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["AutoConnect"] === "undefined") {
          A.store.Bool(ptr + 24 + 21, false);
          A.store.Bool(ptr + 24 + 14, false);
          A.store.Bool(ptr + 24 + 0, false);
          A.store.Ref(ptr + 24 + 4, undefined);
          A.store.Bool(ptr + 24 + 15, false);
          A.store.Bool(ptr + 24 + 8, false);
          A.store.Bool(ptr + 24 + 16, false);
          A.store.Bool(ptr + 24 + 9, false);
          A.store.Bool(ptr + 24 + 17, false);
          A.store.Bool(ptr + 24 + 10, false);
          A.store.Bool(ptr + 24 + 18, false);
          A.store.Bool(ptr + 24 + 11, false);
          A.store.Bool(ptr + 24 + 19, false);
          A.store.Bool(ptr + 24 + 12, false);
          A.store.Bool(ptr + 24 + 20, false);
          A.store.Bool(ptr + 24 + 13, false);
        } else {
          A.store.Bool(ptr + 24 + 21, true);
          A.store.Bool(ptr + 24 + 14, "Active" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 24 + 0, x["AutoConnect"]["Active"] ? true : false);
          A.store.Ref(ptr + 24 + 4, x["AutoConnect"]["Effective"]);
          A.store.Bool(ptr + 24 + 15, "UserPolicy" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 24 + 8, x["AutoConnect"]["UserPolicy"] ? true : false);
          A.store.Bool(ptr + 24 + 16, "DevicePolicy" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 24 + 9, x["AutoConnect"]["DevicePolicy"] ? true : false);
          A.store.Bool(ptr + 24 + 17, "UserSetting" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 24 + 10, x["AutoConnect"]["UserSetting"] ? true : false);
          A.store.Bool(ptr + 24 + 18, "SharedSetting" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 24 + 11, x["AutoConnect"]["SharedSetting"] ? true : false);
          A.store.Bool(ptr + 24 + 19, "UserEditable" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 24 + 12, x["AutoConnect"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 24 + 20, "DeviceEditable" in x["AutoConnect"] ? true : false);
          A.store.Bool(ptr + 24 + 13, x["AutoConnect"]["DeviceEditable"] ? true : false);
        }
        A.store.Ref(ptr + 48, x["BSSID"]);
        A.store.Bool(ptr + 220, "Frequency" in x ? true : false);
        A.store.Int32(ptr + 52, x["Frequency"] === undefined ? 0 : (x["Frequency"] as number));
        A.store.Ref(ptr + 56, x["FrequencyList"]);

        if (typeof x["HexSSID"] === "undefined") {
          A.store.Bool(ptr + 60 + 28, false);
          A.store.Ref(ptr + 60 + 0, undefined);
          A.store.Ref(ptr + 60 + 4, undefined);
          A.store.Ref(ptr + 60 + 8, undefined);
          A.store.Ref(ptr + 60 + 12, undefined);
          A.store.Ref(ptr + 60 + 16, undefined);
          A.store.Ref(ptr + 60 + 20, undefined);
          A.store.Bool(ptr + 60 + 26, false);
          A.store.Bool(ptr + 60 + 24, false);
          A.store.Bool(ptr + 60 + 27, false);
          A.store.Bool(ptr + 60 + 25, false);
        } else {
          A.store.Bool(ptr + 60 + 28, true);
          A.store.Ref(ptr + 60 + 0, x["HexSSID"]["Active"]);
          A.store.Ref(ptr + 60 + 4, x["HexSSID"]["Effective"]);
          A.store.Ref(ptr + 60 + 8, x["HexSSID"]["UserPolicy"]);
          A.store.Ref(ptr + 60 + 12, x["HexSSID"]["DevicePolicy"]);
          A.store.Ref(ptr + 60 + 16, x["HexSSID"]["UserSetting"]);
          A.store.Ref(ptr + 60 + 20, x["HexSSID"]["SharedSetting"]);
          A.store.Bool(ptr + 60 + 26, "UserEditable" in x["HexSSID"] ? true : false);
          A.store.Bool(ptr + 60 + 24, x["HexSSID"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 60 + 27, "DeviceEditable" in x["HexSSID"] ? true : false);
          A.store.Bool(ptr + 60 + 25, x["HexSSID"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["HiddenSSID"] === "undefined") {
          A.store.Bool(ptr + 92 + 21, false);
          A.store.Bool(ptr + 92 + 14, false);
          A.store.Bool(ptr + 92 + 0, false);
          A.store.Ref(ptr + 92 + 4, undefined);
          A.store.Bool(ptr + 92 + 15, false);
          A.store.Bool(ptr + 92 + 8, false);
          A.store.Bool(ptr + 92 + 16, false);
          A.store.Bool(ptr + 92 + 9, false);
          A.store.Bool(ptr + 92 + 17, false);
          A.store.Bool(ptr + 92 + 10, false);
          A.store.Bool(ptr + 92 + 18, false);
          A.store.Bool(ptr + 92 + 11, false);
          A.store.Bool(ptr + 92 + 19, false);
          A.store.Bool(ptr + 92 + 12, false);
          A.store.Bool(ptr + 92 + 20, false);
          A.store.Bool(ptr + 92 + 13, false);
        } else {
          A.store.Bool(ptr + 92 + 21, true);
          A.store.Bool(ptr + 92 + 14, "Active" in x["HiddenSSID"] ? true : false);
          A.store.Bool(ptr + 92 + 0, x["HiddenSSID"]["Active"] ? true : false);
          A.store.Ref(ptr + 92 + 4, x["HiddenSSID"]["Effective"]);
          A.store.Bool(ptr + 92 + 15, "UserPolicy" in x["HiddenSSID"] ? true : false);
          A.store.Bool(ptr + 92 + 8, x["HiddenSSID"]["UserPolicy"] ? true : false);
          A.store.Bool(ptr + 92 + 16, "DevicePolicy" in x["HiddenSSID"] ? true : false);
          A.store.Bool(ptr + 92 + 9, x["HiddenSSID"]["DevicePolicy"] ? true : false);
          A.store.Bool(ptr + 92 + 17, "UserSetting" in x["HiddenSSID"] ? true : false);
          A.store.Bool(ptr + 92 + 10, x["HiddenSSID"]["UserSetting"] ? true : false);
          A.store.Bool(ptr + 92 + 18, "SharedSetting" in x["HiddenSSID"] ? true : false);
          A.store.Bool(ptr + 92 + 11, x["HiddenSSID"]["SharedSetting"] ? true : false);
          A.store.Bool(ptr + 92 + 19, "UserEditable" in x["HiddenSSID"] ? true : false);
          A.store.Bool(ptr + 92 + 12, x["HiddenSSID"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 92 + 20, "DeviceEditable" in x["HiddenSSID"] ? true : false);
          A.store.Bool(ptr + 92 + 13, x["HiddenSSID"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["RoamThreshold"] === "undefined") {
          A.store.Bool(ptr + 116 + 33, false);
          A.store.Bool(ptr + 116 + 26, false);
          A.store.Int32(ptr + 116 + 0, 0);
          A.store.Ref(ptr + 116 + 4, undefined);
          A.store.Bool(ptr + 116 + 27, false);
          A.store.Int32(ptr + 116 + 8, 0);
          A.store.Bool(ptr + 116 + 28, false);
          A.store.Int32(ptr + 116 + 12, 0);
          A.store.Bool(ptr + 116 + 29, false);
          A.store.Int32(ptr + 116 + 16, 0);
          A.store.Bool(ptr + 116 + 30, false);
          A.store.Int32(ptr + 116 + 20, 0);
          A.store.Bool(ptr + 116 + 31, false);
          A.store.Bool(ptr + 116 + 24, false);
          A.store.Bool(ptr + 116 + 32, false);
          A.store.Bool(ptr + 116 + 25, false);
        } else {
          A.store.Bool(ptr + 116 + 33, true);
          A.store.Bool(ptr + 116 + 26, "Active" in x["RoamThreshold"] ? true : false);
          A.store.Int32(
            ptr + 116 + 0,
            x["RoamThreshold"]["Active"] === undefined ? 0 : (x["RoamThreshold"]["Active"] as number)
          );
          A.store.Ref(ptr + 116 + 4, x["RoamThreshold"]["Effective"]);
          A.store.Bool(ptr + 116 + 27, "UserPolicy" in x["RoamThreshold"] ? true : false);
          A.store.Int32(
            ptr + 116 + 8,
            x["RoamThreshold"]["UserPolicy"] === undefined ? 0 : (x["RoamThreshold"]["UserPolicy"] as number)
          );
          A.store.Bool(ptr + 116 + 28, "DevicePolicy" in x["RoamThreshold"] ? true : false);
          A.store.Int32(
            ptr + 116 + 12,
            x["RoamThreshold"]["DevicePolicy"] === undefined ? 0 : (x["RoamThreshold"]["DevicePolicy"] as number)
          );
          A.store.Bool(ptr + 116 + 29, "UserSetting" in x["RoamThreshold"] ? true : false);
          A.store.Int32(
            ptr + 116 + 16,
            x["RoamThreshold"]["UserSetting"] === undefined ? 0 : (x["RoamThreshold"]["UserSetting"] as number)
          );
          A.store.Bool(ptr + 116 + 30, "SharedSetting" in x["RoamThreshold"] ? true : false);
          A.store.Int32(
            ptr + 116 + 20,
            x["RoamThreshold"]["SharedSetting"] === undefined ? 0 : (x["RoamThreshold"]["SharedSetting"] as number)
          );
          A.store.Bool(ptr + 116 + 31, "UserEditable" in x["RoamThreshold"] ? true : false);
          A.store.Bool(ptr + 116 + 24, x["RoamThreshold"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 116 + 32, "DeviceEditable" in x["RoamThreshold"] ? true : false);
          A.store.Bool(ptr + 116 + 25, x["RoamThreshold"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["SSID"] === "undefined") {
          A.store.Bool(ptr + 152 + 28, false);
          A.store.Ref(ptr + 152 + 0, undefined);
          A.store.Ref(ptr + 152 + 4, undefined);
          A.store.Ref(ptr + 152 + 8, undefined);
          A.store.Ref(ptr + 152 + 12, undefined);
          A.store.Ref(ptr + 152 + 16, undefined);
          A.store.Ref(ptr + 152 + 20, undefined);
          A.store.Bool(ptr + 152 + 26, false);
          A.store.Bool(ptr + 152 + 24, false);
          A.store.Bool(ptr + 152 + 27, false);
          A.store.Bool(ptr + 152 + 25, false);
        } else {
          A.store.Bool(ptr + 152 + 28, true);
          A.store.Ref(ptr + 152 + 0, x["SSID"]["Active"]);
          A.store.Ref(ptr + 152 + 4, x["SSID"]["Effective"]);
          A.store.Ref(ptr + 152 + 8, x["SSID"]["UserPolicy"]);
          A.store.Ref(ptr + 152 + 12, x["SSID"]["DevicePolicy"]);
          A.store.Ref(ptr + 152 + 16, x["SSID"]["UserSetting"]);
          A.store.Ref(ptr + 152 + 20, x["SSID"]["SharedSetting"]);
          A.store.Bool(ptr + 152 + 26, "UserEditable" in x["SSID"] ? true : false);
          A.store.Bool(ptr + 152 + 24, x["SSID"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 152 + 27, "DeviceEditable" in x["SSID"] ? true : false);
          A.store.Bool(ptr + 152 + 25, x["SSID"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["Security"] === "undefined") {
          A.store.Bool(ptr + 184 + 28, false);
          A.store.Ref(ptr + 184 + 0, undefined);
          A.store.Ref(ptr + 184 + 4, undefined);
          A.store.Ref(ptr + 184 + 8, undefined);
          A.store.Ref(ptr + 184 + 12, undefined);
          A.store.Ref(ptr + 184 + 16, undefined);
          A.store.Ref(ptr + 184 + 20, undefined);
          A.store.Bool(ptr + 184 + 26, false);
          A.store.Bool(ptr + 184 + 24, false);
          A.store.Bool(ptr + 184 + 27, false);
          A.store.Bool(ptr + 184 + 25, false);
        } else {
          A.store.Bool(ptr + 184 + 28, true);
          A.store.Ref(ptr + 184 + 0, x["Security"]["Active"]);
          A.store.Ref(ptr + 184 + 4, x["Security"]["Effective"]);
          A.store.Ref(ptr + 184 + 8, x["Security"]["UserPolicy"]);
          A.store.Ref(ptr + 184 + 12, x["Security"]["DevicePolicy"]);
          A.store.Ref(ptr + 184 + 16, x["Security"]["UserSetting"]);
          A.store.Ref(ptr + 184 + 20, x["Security"]["SharedSetting"]);
          A.store.Bool(ptr + 184 + 26, "UserEditable" in x["Security"] ? true : false);
          A.store.Bool(ptr + 184 + 24, x["Security"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 184 + 27, "DeviceEditable" in x["Security"] ? true : false);
          A.store.Bool(ptr + 184 + 25, x["Security"]["DeviceEditable"] ? true : false);
        }
        A.store.Bool(ptr + 221, "SignalStrength" in x ? true : false);
        A.store.Int32(ptr + 216, x["SignalStrength"] === undefined ? 0 : (x["SignalStrength"] as number));
      }
    },
    "load_ManagedWiFiProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 21)) {
        x["AllowGatewayARPPolling"] = {};
        if (A.load.Bool(ptr + 0 + 14)) {
          x["AllowGatewayARPPolling"]["Active"] = A.load.Bool(ptr + 0 + 0);
        } else {
          delete x["AllowGatewayARPPolling"]["Active"];
        }
        x["AllowGatewayARPPolling"]["Effective"] = A.load.Ref(ptr + 0 + 4, undefined);
        if (A.load.Bool(ptr + 0 + 15)) {
          x["AllowGatewayARPPolling"]["UserPolicy"] = A.load.Bool(ptr + 0 + 8);
        } else {
          delete x["AllowGatewayARPPolling"]["UserPolicy"];
        }
        if (A.load.Bool(ptr + 0 + 16)) {
          x["AllowGatewayARPPolling"]["DevicePolicy"] = A.load.Bool(ptr + 0 + 9);
        } else {
          delete x["AllowGatewayARPPolling"]["DevicePolicy"];
        }
        if (A.load.Bool(ptr + 0 + 17)) {
          x["AllowGatewayARPPolling"]["UserSetting"] = A.load.Bool(ptr + 0 + 10);
        } else {
          delete x["AllowGatewayARPPolling"]["UserSetting"];
        }
        if (A.load.Bool(ptr + 0 + 18)) {
          x["AllowGatewayARPPolling"]["SharedSetting"] = A.load.Bool(ptr + 0 + 11);
        } else {
          delete x["AllowGatewayARPPolling"]["SharedSetting"];
        }
        if (A.load.Bool(ptr + 0 + 19)) {
          x["AllowGatewayARPPolling"]["UserEditable"] = A.load.Bool(ptr + 0 + 12);
        } else {
          delete x["AllowGatewayARPPolling"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 0 + 20)) {
          x["AllowGatewayARPPolling"]["DeviceEditable"] = A.load.Bool(ptr + 0 + 13);
        } else {
          delete x["AllowGatewayARPPolling"]["DeviceEditable"];
        }
      } else {
        delete x["AllowGatewayARPPolling"];
      }
      if (A.load.Bool(ptr + 24 + 21)) {
        x["AutoConnect"] = {};
        if (A.load.Bool(ptr + 24 + 14)) {
          x["AutoConnect"]["Active"] = A.load.Bool(ptr + 24 + 0);
        } else {
          delete x["AutoConnect"]["Active"];
        }
        x["AutoConnect"]["Effective"] = A.load.Ref(ptr + 24 + 4, undefined);
        if (A.load.Bool(ptr + 24 + 15)) {
          x["AutoConnect"]["UserPolicy"] = A.load.Bool(ptr + 24 + 8);
        } else {
          delete x["AutoConnect"]["UserPolicy"];
        }
        if (A.load.Bool(ptr + 24 + 16)) {
          x["AutoConnect"]["DevicePolicy"] = A.load.Bool(ptr + 24 + 9);
        } else {
          delete x["AutoConnect"]["DevicePolicy"];
        }
        if (A.load.Bool(ptr + 24 + 17)) {
          x["AutoConnect"]["UserSetting"] = A.load.Bool(ptr + 24 + 10);
        } else {
          delete x["AutoConnect"]["UserSetting"];
        }
        if (A.load.Bool(ptr + 24 + 18)) {
          x["AutoConnect"]["SharedSetting"] = A.load.Bool(ptr + 24 + 11);
        } else {
          delete x["AutoConnect"]["SharedSetting"];
        }
        if (A.load.Bool(ptr + 24 + 19)) {
          x["AutoConnect"]["UserEditable"] = A.load.Bool(ptr + 24 + 12);
        } else {
          delete x["AutoConnect"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 24 + 20)) {
          x["AutoConnect"]["DeviceEditable"] = A.load.Bool(ptr + 24 + 13);
        } else {
          delete x["AutoConnect"]["DeviceEditable"];
        }
      } else {
        delete x["AutoConnect"];
      }
      x["BSSID"] = A.load.Ref(ptr + 48, undefined);
      if (A.load.Bool(ptr + 220)) {
        x["Frequency"] = A.load.Int32(ptr + 52);
      } else {
        delete x["Frequency"];
      }
      x["FrequencyList"] = A.load.Ref(ptr + 56, undefined);
      if (A.load.Bool(ptr + 60 + 28)) {
        x["HexSSID"] = {};
        x["HexSSID"]["Active"] = A.load.Ref(ptr + 60 + 0, undefined);
        x["HexSSID"]["Effective"] = A.load.Ref(ptr + 60 + 4, undefined);
        x["HexSSID"]["UserPolicy"] = A.load.Ref(ptr + 60 + 8, undefined);
        x["HexSSID"]["DevicePolicy"] = A.load.Ref(ptr + 60 + 12, undefined);
        x["HexSSID"]["UserSetting"] = A.load.Ref(ptr + 60 + 16, undefined);
        x["HexSSID"]["SharedSetting"] = A.load.Ref(ptr + 60 + 20, undefined);
        if (A.load.Bool(ptr + 60 + 26)) {
          x["HexSSID"]["UserEditable"] = A.load.Bool(ptr + 60 + 24);
        } else {
          delete x["HexSSID"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 60 + 27)) {
          x["HexSSID"]["DeviceEditable"] = A.load.Bool(ptr + 60 + 25);
        } else {
          delete x["HexSSID"]["DeviceEditable"];
        }
      } else {
        delete x["HexSSID"];
      }
      if (A.load.Bool(ptr + 92 + 21)) {
        x["HiddenSSID"] = {};
        if (A.load.Bool(ptr + 92 + 14)) {
          x["HiddenSSID"]["Active"] = A.load.Bool(ptr + 92 + 0);
        } else {
          delete x["HiddenSSID"]["Active"];
        }
        x["HiddenSSID"]["Effective"] = A.load.Ref(ptr + 92 + 4, undefined);
        if (A.load.Bool(ptr + 92 + 15)) {
          x["HiddenSSID"]["UserPolicy"] = A.load.Bool(ptr + 92 + 8);
        } else {
          delete x["HiddenSSID"]["UserPolicy"];
        }
        if (A.load.Bool(ptr + 92 + 16)) {
          x["HiddenSSID"]["DevicePolicy"] = A.load.Bool(ptr + 92 + 9);
        } else {
          delete x["HiddenSSID"]["DevicePolicy"];
        }
        if (A.load.Bool(ptr + 92 + 17)) {
          x["HiddenSSID"]["UserSetting"] = A.load.Bool(ptr + 92 + 10);
        } else {
          delete x["HiddenSSID"]["UserSetting"];
        }
        if (A.load.Bool(ptr + 92 + 18)) {
          x["HiddenSSID"]["SharedSetting"] = A.load.Bool(ptr + 92 + 11);
        } else {
          delete x["HiddenSSID"]["SharedSetting"];
        }
        if (A.load.Bool(ptr + 92 + 19)) {
          x["HiddenSSID"]["UserEditable"] = A.load.Bool(ptr + 92 + 12);
        } else {
          delete x["HiddenSSID"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 92 + 20)) {
          x["HiddenSSID"]["DeviceEditable"] = A.load.Bool(ptr + 92 + 13);
        } else {
          delete x["HiddenSSID"]["DeviceEditable"];
        }
      } else {
        delete x["HiddenSSID"];
      }
      if (A.load.Bool(ptr + 116 + 33)) {
        x["RoamThreshold"] = {};
        if (A.load.Bool(ptr + 116 + 26)) {
          x["RoamThreshold"]["Active"] = A.load.Int32(ptr + 116 + 0);
        } else {
          delete x["RoamThreshold"]["Active"];
        }
        x["RoamThreshold"]["Effective"] = A.load.Ref(ptr + 116 + 4, undefined);
        if (A.load.Bool(ptr + 116 + 27)) {
          x["RoamThreshold"]["UserPolicy"] = A.load.Int32(ptr + 116 + 8);
        } else {
          delete x["RoamThreshold"]["UserPolicy"];
        }
        if (A.load.Bool(ptr + 116 + 28)) {
          x["RoamThreshold"]["DevicePolicy"] = A.load.Int32(ptr + 116 + 12);
        } else {
          delete x["RoamThreshold"]["DevicePolicy"];
        }
        if (A.load.Bool(ptr + 116 + 29)) {
          x["RoamThreshold"]["UserSetting"] = A.load.Int32(ptr + 116 + 16);
        } else {
          delete x["RoamThreshold"]["UserSetting"];
        }
        if (A.load.Bool(ptr + 116 + 30)) {
          x["RoamThreshold"]["SharedSetting"] = A.load.Int32(ptr + 116 + 20);
        } else {
          delete x["RoamThreshold"]["SharedSetting"];
        }
        if (A.load.Bool(ptr + 116 + 31)) {
          x["RoamThreshold"]["UserEditable"] = A.load.Bool(ptr + 116 + 24);
        } else {
          delete x["RoamThreshold"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 116 + 32)) {
          x["RoamThreshold"]["DeviceEditable"] = A.load.Bool(ptr + 116 + 25);
        } else {
          delete x["RoamThreshold"]["DeviceEditable"];
        }
      } else {
        delete x["RoamThreshold"];
      }
      if (A.load.Bool(ptr + 152 + 28)) {
        x["SSID"] = {};
        x["SSID"]["Active"] = A.load.Ref(ptr + 152 + 0, undefined);
        x["SSID"]["Effective"] = A.load.Ref(ptr + 152 + 4, undefined);
        x["SSID"]["UserPolicy"] = A.load.Ref(ptr + 152 + 8, undefined);
        x["SSID"]["DevicePolicy"] = A.load.Ref(ptr + 152 + 12, undefined);
        x["SSID"]["UserSetting"] = A.load.Ref(ptr + 152 + 16, undefined);
        x["SSID"]["SharedSetting"] = A.load.Ref(ptr + 152 + 20, undefined);
        if (A.load.Bool(ptr + 152 + 26)) {
          x["SSID"]["UserEditable"] = A.load.Bool(ptr + 152 + 24);
        } else {
          delete x["SSID"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 152 + 27)) {
          x["SSID"]["DeviceEditable"] = A.load.Bool(ptr + 152 + 25);
        } else {
          delete x["SSID"]["DeviceEditable"];
        }
      } else {
        delete x["SSID"];
      }
      if (A.load.Bool(ptr + 184 + 28)) {
        x["Security"] = {};
        x["Security"]["Active"] = A.load.Ref(ptr + 184 + 0, undefined);
        x["Security"]["Effective"] = A.load.Ref(ptr + 184 + 4, undefined);
        x["Security"]["UserPolicy"] = A.load.Ref(ptr + 184 + 8, undefined);
        x["Security"]["DevicePolicy"] = A.load.Ref(ptr + 184 + 12, undefined);
        x["Security"]["UserSetting"] = A.load.Ref(ptr + 184 + 16, undefined);
        x["Security"]["SharedSetting"] = A.load.Ref(ptr + 184 + 20, undefined);
        if (A.load.Bool(ptr + 184 + 26)) {
          x["Security"]["UserEditable"] = A.load.Bool(ptr + 184 + 24);
        } else {
          delete x["Security"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 184 + 27)) {
          x["Security"]["DeviceEditable"] = A.load.Bool(ptr + 184 + 25);
        } else {
          delete x["Security"]["DeviceEditable"];
        }
      } else {
        delete x["Security"];
      }
      if (A.load.Bool(ptr + 221)) {
        x["SignalStrength"] = A.load.Int32(ptr + 216);
      } else {
        delete x["SignalStrength"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManagedProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 1293, false);

        A.store.Bool(ptr + 0 + 134, false);

        A.store.Bool(ptr + 0 + 0 + 21, false);
        A.store.Bool(ptr + 0 + 0 + 14, false);
        A.store.Bool(ptr + 0 + 0 + 0, false);
        A.store.Ref(ptr + 0 + 0 + 4, undefined);
        A.store.Bool(ptr + 0 + 0 + 15, false);
        A.store.Bool(ptr + 0 + 0 + 8, false);
        A.store.Bool(ptr + 0 + 0 + 16, false);
        A.store.Bool(ptr + 0 + 0 + 9, false);
        A.store.Bool(ptr + 0 + 0 + 17, false);
        A.store.Bool(ptr + 0 + 0 + 10, false);
        A.store.Bool(ptr + 0 + 0 + 18, false);
        A.store.Bool(ptr + 0 + 0 + 11, false);
        A.store.Bool(ptr + 0 + 0 + 19, false);
        A.store.Bool(ptr + 0 + 0 + 12, false);
        A.store.Bool(ptr + 0 + 0 + 20, false);
        A.store.Bool(ptr + 0 + 0 + 13, false);
        A.store.Ref(ptr + 0 + 24, undefined);
        A.store.Enum(ptr + 0 + 28, -1);
        A.store.Bool(ptr + 0 + 129, false);
        A.store.Bool(ptr + 0 + 32, false);
        A.store.Ref(ptr + 0 + 36, undefined);
        A.store.Ref(ptr + 0 + 40, undefined);
        A.store.Ref(ptr + 0 + 44, undefined);
        A.store.Ref(ptr + 0 + 48, undefined);
        A.store.Ref(ptr + 0 + 52, undefined);
        A.store.Ref(ptr + 0 + 56, undefined);
        A.store.Ref(ptr + 0 + 60, undefined);
        A.store.Ref(ptr + 0 + 64, undefined);

        A.store.Bool(ptr + 0 + 68 + 12, false);
        A.store.Ref(ptr + 0 + 68 + 0, undefined);
        A.store.Ref(ptr + 0 + 68 + 4, undefined);
        A.store.Ref(ptr + 0 + 68 + 8, undefined);
        A.store.Ref(ptr + 0 + 84, undefined);
        A.store.Bool(ptr + 0 + 130, false);
        A.store.Bool(ptr + 0 + 88, false);

        A.store.Bool(ptr + 0 + 92 + 12, false);
        A.store.Ref(ptr + 0 + 92 + 0, undefined);
        A.store.Ref(ptr + 0 + 92 + 4, undefined);
        A.store.Ref(ptr + 0 + 92 + 8, undefined);

        A.store.Bool(ptr + 0 + 108 + 14, false);
        A.store.Ref(ptr + 0 + 108 + 0, undefined);
        A.store.Bool(ptr + 0 + 108 + 12, false);
        A.store.Bool(ptr + 0 + 108 + 4, false);
        A.store.Bool(ptr + 0 + 108 + 13, false);
        A.store.Int32(ptr + 0 + 108 + 8, 0);
        A.store.Bool(ptr + 0 + 131, false);
        A.store.Bool(ptr + 0 + 123, false);
        A.store.Bool(ptr + 0 + 132, false);
        A.store.Int32(ptr + 0 + 124, 0);
        A.store.Bool(ptr + 0 + 133, false);
        A.store.Bool(ptr + 0 + 128, false);
        A.store.Bool(ptr + 1291, false);
        A.store.Bool(ptr + 135, false);
        A.store.Enum(ptr + 136, -1);
        A.store.Ref(ptr + 140, undefined);

        A.store.Bool(ptr + 144 + 53, false);

        A.store.Bool(ptr + 144 + 0 + 21, false);
        A.store.Bool(ptr + 144 + 0 + 14, false);
        A.store.Bool(ptr + 144 + 0 + 0, false);
        A.store.Ref(ptr + 144 + 0 + 4, undefined);
        A.store.Bool(ptr + 144 + 0 + 15, false);
        A.store.Bool(ptr + 144 + 0 + 8, false);
        A.store.Bool(ptr + 144 + 0 + 16, false);
        A.store.Bool(ptr + 144 + 0 + 9, false);
        A.store.Bool(ptr + 144 + 0 + 17, false);
        A.store.Bool(ptr + 144 + 0 + 10, false);
        A.store.Bool(ptr + 144 + 0 + 18, false);
        A.store.Bool(ptr + 144 + 0 + 11, false);
        A.store.Bool(ptr + 144 + 0 + 19, false);
        A.store.Bool(ptr + 144 + 0 + 12, false);
        A.store.Bool(ptr + 144 + 0 + 20, false);
        A.store.Bool(ptr + 144 + 0 + 13, false);

        A.store.Bool(ptr + 144 + 24 + 28, false);
        A.store.Ref(ptr + 144 + 24 + 0, undefined);
        A.store.Ref(ptr + 144 + 24 + 4, undefined);
        A.store.Ref(ptr + 144 + 24 + 8, undefined);
        A.store.Ref(ptr + 144 + 24 + 12, undefined);
        A.store.Ref(ptr + 144 + 24 + 16, undefined);
        A.store.Ref(ptr + 144 + 24 + 20, undefined);
        A.store.Bool(ptr + 144 + 24 + 26, false);
        A.store.Bool(ptr + 144 + 24 + 24, false);
        A.store.Bool(ptr + 144 + 24 + 27, false);
        A.store.Bool(ptr + 144 + 24 + 25, false);
        A.store.Ref(ptr + 200, undefined);

        A.store.Bool(ptr + 204 + 28, false);
        A.store.Enum(ptr + 204 + 0, -1);
        A.store.Ref(ptr + 204 + 4, undefined);
        A.store.Enum(ptr + 204 + 8, -1);
        A.store.Enum(ptr + 204 + 12, -1);
        A.store.Enum(ptr + 204 + 16, -1);
        A.store.Enum(ptr + 204 + 20, -1);
        A.store.Bool(ptr + 204 + 26, false);
        A.store.Bool(ptr + 204 + 24, false);
        A.store.Bool(ptr + 204 + 27, false);
        A.store.Bool(ptr + 204 + 25, false);
        A.store.Ref(ptr + 236, undefined);
        A.store.Ref(ptr + 240, undefined);

        A.store.Bool(ptr + 244 + 21, false);
        A.store.Bool(ptr + 244 + 14, false);
        A.store.Bool(ptr + 244 + 0, false);
        A.store.Ref(ptr + 244 + 4, undefined);
        A.store.Bool(ptr + 244 + 15, false);
        A.store.Bool(ptr + 244 + 8, false);
        A.store.Bool(ptr + 244 + 16, false);
        A.store.Bool(ptr + 244 + 9, false);
        A.store.Bool(ptr + 244 + 17, false);
        A.store.Bool(ptr + 244 + 10, false);
        A.store.Bool(ptr + 244 + 18, false);
        A.store.Bool(ptr + 244 + 11, false);
        A.store.Bool(ptr + 244 + 19, false);
        A.store.Bool(ptr + 244 + 12, false);
        A.store.Bool(ptr + 244 + 20, false);
        A.store.Bool(ptr + 244 + 13, false);

        A.store.Bool(ptr + 268 + 28, false);
        A.store.Ref(ptr + 268 + 0, undefined);
        A.store.Ref(ptr + 268 + 4, undefined);
        A.store.Ref(ptr + 268 + 8, undefined);
        A.store.Ref(ptr + 268 + 12, undefined);
        A.store.Ref(ptr + 268 + 16, undefined);
        A.store.Ref(ptr + 268 + 20, undefined);
        A.store.Bool(ptr + 268 + 26, false);
        A.store.Bool(ptr + 268 + 24, false);
        A.store.Bool(ptr + 268 + 27, false);
        A.store.Bool(ptr + 268 + 25, false);

        A.store.Bool(ptr + 300 + 28, false);
        A.store.Enum(ptr + 300 + 0, -1);
        A.store.Ref(ptr + 300 + 4, undefined);
        A.store.Enum(ptr + 300 + 8, -1);
        A.store.Enum(ptr + 300 + 12, -1);
        A.store.Enum(ptr + 300 + 16, -1);
        A.store.Enum(ptr + 300 + 20, -1);
        A.store.Bool(ptr + 300 + 26, false);
        A.store.Bool(ptr + 300 + 24, false);
        A.store.Bool(ptr + 300 + 27, false);
        A.store.Bool(ptr + 300 + 25, false);

        A.store.Bool(ptr + 332 + 33, false);
        A.store.Bool(ptr + 332 + 26, false);
        A.store.Int32(ptr + 332 + 0, 0);
        A.store.Ref(ptr + 332 + 4, undefined);
        A.store.Bool(ptr + 332 + 27, false);
        A.store.Int32(ptr + 332 + 8, 0);
        A.store.Bool(ptr + 332 + 28, false);
        A.store.Int32(ptr + 332 + 12, 0);
        A.store.Bool(ptr + 332 + 29, false);
        A.store.Int32(ptr + 332 + 16, 0);
        A.store.Bool(ptr + 332 + 30, false);
        A.store.Int32(ptr + 332 + 20, 0);
        A.store.Bool(ptr + 332 + 31, false);
        A.store.Bool(ptr + 332 + 24, false);
        A.store.Bool(ptr + 332 + 32, false);
        A.store.Bool(ptr + 332 + 25, false);

        A.store.Bool(ptr + 368 + 365, false);

        A.store.Bool(ptr + 368 + 0 + 28, false);
        A.store.Enum(ptr + 368 + 0 + 0, -1);
        A.store.Ref(ptr + 368 + 0 + 4, undefined);
        A.store.Enum(ptr + 368 + 0 + 8, -1);
        A.store.Enum(ptr + 368 + 0 + 12, -1);
        A.store.Enum(ptr + 368 + 0 + 16, -1);
        A.store.Enum(ptr + 368 + 0 + 20, -1);
        A.store.Bool(ptr + 368 + 0 + 26, false);
        A.store.Bool(ptr + 368 + 0 + 24, false);
        A.store.Bool(ptr + 368 + 0 + 27, false);
        A.store.Bool(ptr + 368 + 0 + 25, false);

        A.store.Bool(ptr + 368 + 32 + 271, false);

        A.store.Bool(ptr + 368 + 32 + 0 + 66, false);

        A.store.Bool(ptr + 368 + 32 + 0 + 0 + 28, false);
        A.store.Ref(ptr + 368 + 32 + 0 + 0 + 0, undefined);
        A.store.Ref(ptr + 368 + 32 + 0 + 0 + 4, undefined);
        A.store.Ref(ptr + 368 + 32 + 0 + 0 + 8, undefined);
        A.store.Ref(ptr + 368 + 32 + 0 + 0 + 12, undefined);
        A.store.Ref(ptr + 368 + 32 + 0 + 0 + 16, undefined);
        A.store.Ref(ptr + 368 + 32 + 0 + 0 + 20, undefined);
        A.store.Bool(ptr + 368 + 32 + 0 + 0 + 26, false);
        A.store.Bool(ptr + 368 + 32 + 0 + 0 + 24, false);
        A.store.Bool(ptr + 368 + 32 + 0 + 0 + 27, false);
        A.store.Bool(ptr + 368 + 32 + 0 + 0 + 25, false);

        A.store.Bool(ptr + 368 + 32 + 0 + 32 + 33, false);
        A.store.Bool(ptr + 368 + 32 + 0 + 32 + 26, false);
        A.store.Int32(ptr + 368 + 32 + 0 + 32 + 0, 0);
        A.store.Ref(ptr + 368 + 32 + 0 + 32 + 4, undefined);
        A.store.Bool(ptr + 368 + 32 + 0 + 32 + 27, false);
        A.store.Int32(ptr + 368 + 32 + 0 + 32 + 8, 0);
        A.store.Bool(ptr + 368 + 32 + 0 + 32 + 28, false);
        A.store.Int32(ptr + 368 + 32 + 0 + 32 + 12, 0);
        A.store.Bool(ptr + 368 + 32 + 0 + 32 + 29, false);
        A.store.Int32(ptr + 368 + 32 + 0 + 32 + 16, 0);
        A.store.Bool(ptr + 368 + 32 + 0 + 32 + 30, false);
        A.store.Int32(ptr + 368 + 32 + 0 + 32 + 20, 0);
        A.store.Bool(ptr + 368 + 32 + 0 + 32 + 31, false);
        A.store.Bool(ptr + 368 + 32 + 0 + 32 + 24, false);
        A.store.Bool(ptr + 368 + 32 + 0 + 32 + 32, false);
        A.store.Bool(ptr + 368 + 32 + 0 + 32 + 25, false);

        A.store.Bool(ptr + 368 + 32 + 68 + 66, false);

        A.store.Bool(ptr + 368 + 32 + 68 + 0 + 28, false);
        A.store.Ref(ptr + 368 + 32 + 68 + 0 + 0, undefined);
        A.store.Ref(ptr + 368 + 32 + 68 + 0 + 4, undefined);
        A.store.Ref(ptr + 368 + 32 + 68 + 0 + 8, undefined);
        A.store.Ref(ptr + 368 + 32 + 68 + 0 + 12, undefined);
        A.store.Ref(ptr + 368 + 32 + 68 + 0 + 16, undefined);
        A.store.Ref(ptr + 368 + 32 + 68 + 0 + 20, undefined);
        A.store.Bool(ptr + 368 + 32 + 68 + 0 + 26, false);
        A.store.Bool(ptr + 368 + 32 + 68 + 0 + 24, false);
        A.store.Bool(ptr + 368 + 32 + 68 + 0 + 27, false);
        A.store.Bool(ptr + 368 + 32 + 68 + 0 + 25, false);

        A.store.Bool(ptr + 368 + 32 + 68 + 32 + 33, false);
        A.store.Bool(ptr + 368 + 32 + 68 + 32 + 26, false);
        A.store.Int32(ptr + 368 + 32 + 68 + 32 + 0, 0);
        A.store.Ref(ptr + 368 + 32 + 68 + 32 + 4, undefined);
        A.store.Bool(ptr + 368 + 32 + 68 + 32 + 27, false);
        A.store.Int32(ptr + 368 + 32 + 68 + 32 + 8, 0);
        A.store.Bool(ptr + 368 + 32 + 68 + 32 + 28, false);
        A.store.Int32(ptr + 368 + 32 + 68 + 32 + 12, 0);
        A.store.Bool(ptr + 368 + 32 + 68 + 32 + 29, false);
        A.store.Int32(ptr + 368 + 32 + 68 + 32 + 16, 0);
        A.store.Bool(ptr + 368 + 32 + 68 + 32 + 30, false);
        A.store.Int32(ptr + 368 + 32 + 68 + 32 + 20, 0);
        A.store.Bool(ptr + 368 + 32 + 68 + 32 + 31, false);
        A.store.Bool(ptr + 368 + 32 + 68 + 32 + 24, false);
        A.store.Bool(ptr + 368 + 32 + 68 + 32 + 32, false);
        A.store.Bool(ptr + 368 + 32 + 68 + 32 + 25, false);

        A.store.Bool(ptr + 368 + 32 + 136 + 66, false);

        A.store.Bool(ptr + 368 + 32 + 136 + 0 + 28, false);
        A.store.Ref(ptr + 368 + 32 + 136 + 0 + 0, undefined);
        A.store.Ref(ptr + 368 + 32 + 136 + 0 + 4, undefined);
        A.store.Ref(ptr + 368 + 32 + 136 + 0 + 8, undefined);
        A.store.Ref(ptr + 368 + 32 + 136 + 0 + 12, undefined);
        A.store.Ref(ptr + 368 + 32 + 136 + 0 + 16, undefined);
        A.store.Ref(ptr + 368 + 32 + 136 + 0 + 20, undefined);
        A.store.Bool(ptr + 368 + 32 + 136 + 0 + 26, false);
        A.store.Bool(ptr + 368 + 32 + 136 + 0 + 24, false);
        A.store.Bool(ptr + 368 + 32 + 136 + 0 + 27, false);
        A.store.Bool(ptr + 368 + 32 + 136 + 0 + 25, false);

        A.store.Bool(ptr + 368 + 32 + 136 + 32 + 33, false);
        A.store.Bool(ptr + 368 + 32 + 136 + 32 + 26, false);
        A.store.Int32(ptr + 368 + 32 + 136 + 32 + 0, 0);
        A.store.Ref(ptr + 368 + 32 + 136 + 32 + 4, undefined);
        A.store.Bool(ptr + 368 + 32 + 136 + 32 + 27, false);
        A.store.Int32(ptr + 368 + 32 + 136 + 32 + 8, 0);
        A.store.Bool(ptr + 368 + 32 + 136 + 32 + 28, false);
        A.store.Int32(ptr + 368 + 32 + 136 + 32 + 12, 0);
        A.store.Bool(ptr + 368 + 32 + 136 + 32 + 29, false);
        A.store.Int32(ptr + 368 + 32 + 136 + 32 + 16, 0);
        A.store.Bool(ptr + 368 + 32 + 136 + 32 + 30, false);
        A.store.Int32(ptr + 368 + 32 + 136 + 32 + 20, 0);
        A.store.Bool(ptr + 368 + 32 + 136 + 32 + 31, false);
        A.store.Bool(ptr + 368 + 32 + 136 + 32 + 24, false);
        A.store.Bool(ptr + 368 + 32 + 136 + 32 + 32, false);
        A.store.Bool(ptr + 368 + 32 + 136 + 32 + 25, false);

        A.store.Bool(ptr + 368 + 32 + 204 + 66, false);

        A.store.Bool(ptr + 368 + 32 + 204 + 0 + 28, false);
        A.store.Ref(ptr + 368 + 32 + 204 + 0 + 0, undefined);
        A.store.Ref(ptr + 368 + 32 + 204 + 0 + 4, undefined);
        A.store.Ref(ptr + 368 + 32 + 204 + 0 + 8, undefined);
        A.store.Ref(ptr + 368 + 32 + 204 + 0 + 12, undefined);
        A.store.Ref(ptr + 368 + 32 + 204 + 0 + 16, undefined);
        A.store.Ref(ptr + 368 + 32 + 204 + 0 + 20, undefined);
        A.store.Bool(ptr + 368 + 32 + 204 + 0 + 26, false);
        A.store.Bool(ptr + 368 + 32 + 204 + 0 + 24, false);
        A.store.Bool(ptr + 368 + 32 + 204 + 0 + 27, false);
        A.store.Bool(ptr + 368 + 32 + 204 + 0 + 25, false);

        A.store.Bool(ptr + 368 + 32 + 204 + 32 + 33, false);
        A.store.Bool(ptr + 368 + 32 + 204 + 32 + 26, false);
        A.store.Int32(ptr + 368 + 32 + 204 + 32 + 0, 0);
        A.store.Ref(ptr + 368 + 32 + 204 + 32 + 4, undefined);
        A.store.Bool(ptr + 368 + 32 + 204 + 32 + 27, false);
        A.store.Int32(ptr + 368 + 32 + 204 + 32 + 8, 0);
        A.store.Bool(ptr + 368 + 32 + 204 + 32 + 28, false);
        A.store.Int32(ptr + 368 + 32 + 204 + 32 + 12, 0);
        A.store.Bool(ptr + 368 + 32 + 204 + 32 + 29, false);
        A.store.Int32(ptr + 368 + 32 + 204 + 32 + 16, 0);
        A.store.Bool(ptr + 368 + 32 + 204 + 32 + 30, false);
        A.store.Int32(ptr + 368 + 32 + 204 + 32 + 20, 0);
        A.store.Bool(ptr + 368 + 32 + 204 + 32 + 31, false);
        A.store.Bool(ptr + 368 + 32 + 204 + 32 + 24, false);
        A.store.Bool(ptr + 368 + 32 + 204 + 32 + 32, false);
        A.store.Bool(ptr + 368 + 32 + 204 + 32 + 25, false);

        A.store.Bool(ptr + 368 + 304 + 28, false);
        A.store.Ref(ptr + 368 + 304 + 0, undefined);
        A.store.Ref(ptr + 368 + 304 + 4, undefined);
        A.store.Ref(ptr + 368 + 304 + 8, undefined);
        A.store.Ref(ptr + 368 + 304 + 12, undefined);
        A.store.Ref(ptr + 368 + 304 + 16, undefined);
        A.store.Ref(ptr + 368 + 304 + 20, undefined);
        A.store.Bool(ptr + 368 + 304 + 26, false);
        A.store.Bool(ptr + 368 + 304 + 24, false);
        A.store.Bool(ptr + 368 + 304 + 27, false);
        A.store.Bool(ptr + 368 + 304 + 25, false);

        A.store.Bool(ptr + 368 + 336 + 28, false);
        A.store.Ref(ptr + 368 + 336 + 0, undefined);
        A.store.Ref(ptr + 368 + 336 + 4, undefined);
        A.store.Ref(ptr + 368 + 336 + 8, undefined);
        A.store.Ref(ptr + 368 + 336 + 12, undefined);
        A.store.Ref(ptr + 368 + 336 + 16, undefined);
        A.store.Ref(ptr + 368 + 336 + 20, undefined);
        A.store.Bool(ptr + 368 + 336 + 26, false);
        A.store.Bool(ptr + 368 + 336 + 24, false);
        A.store.Bool(ptr + 368 + 336 + 27, false);
        A.store.Bool(ptr + 368 + 336 + 25, false);
        A.store.Bool(ptr + 1292, false);
        A.store.Bool(ptr + 734, false);

        A.store.Bool(ptr + 736 + 193, false);

        A.store.Bool(ptr + 736 + 0 + 28, false);
        A.store.Ref(ptr + 736 + 0 + 0, undefined);
        A.store.Ref(ptr + 736 + 0 + 4, undefined);
        A.store.Ref(ptr + 736 + 0 + 8, undefined);
        A.store.Ref(ptr + 736 + 0 + 12, undefined);
        A.store.Ref(ptr + 736 + 0 + 16, undefined);
        A.store.Ref(ptr + 736 + 0 + 20, undefined);
        A.store.Bool(ptr + 736 + 0 + 26, false);
        A.store.Bool(ptr + 736 + 0 + 24, false);
        A.store.Bool(ptr + 736 + 0 + 27, false);
        A.store.Bool(ptr + 736 + 0 + 25, false);

        A.store.Bool(ptr + 736 + 32 + 28, false);
        A.store.Ref(ptr + 736 + 32 + 0, undefined);
        A.store.Ref(ptr + 736 + 32 + 4, undefined);
        A.store.Ref(ptr + 736 + 32 + 8, undefined);
        A.store.Ref(ptr + 736 + 32 + 12, undefined);
        A.store.Ref(ptr + 736 + 32 + 16, undefined);
        A.store.Ref(ptr + 736 + 32 + 20, undefined);
        A.store.Bool(ptr + 736 + 32 + 26, false);
        A.store.Bool(ptr + 736 + 32 + 24, false);
        A.store.Bool(ptr + 736 + 32 + 27, false);
        A.store.Bool(ptr + 736 + 32 + 25, false);

        A.store.Bool(ptr + 736 + 64 + 28, false);
        A.store.Ref(ptr + 736 + 64 + 0, undefined);
        A.store.Ref(ptr + 736 + 64 + 4, undefined);
        A.store.Ref(ptr + 736 + 64 + 8, undefined);
        A.store.Ref(ptr + 736 + 64 + 12, undefined);
        A.store.Ref(ptr + 736 + 64 + 16, undefined);
        A.store.Ref(ptr + 736 + 64 + 20, undefined);
        A.store.Bool(ptr + 736 + 64 + 26, false);
        A.store.Bool(ptr + 736 + 64 + 24, false);
        A.store.Bool(ptr + 736 + 64 + 27, false);
        A.store.Bool(ptr + 736 + 64 + 25, false);

        A.store.Bool(ptr + 736 + 96 + 33, false);
        A.store.Bool(ptr + 736 + 96 + 26, false);
        A.store.Int32(ptr + 736 + 96 + 0, 0);
        A.store.Ref(ptr + 736 + 96 + 4, undefined);
        A.store.Bool(ptr + 736 + 96 + 27, false);
        A.store.Int32(ptr + 736 + 96 + 8, 0);
        A.store.Bool(ptr + 736 + 96 + 28, false);
        A.store.Int32(ptr + 736 + 96 + 12, 0);
        A.store.Bool(ptr + 736 + 96 + 29, false);
        A.store.Int32(ptr + 736 + 96 + 16, 0);
        A.store.Bool(ptr + 736 + 96 + 30, false);
        A.store.Int32(ptr + 736 + 96 + 20, 0);
        A.store.Bool(ptr + 736 + 96 + 31, false);
        A.store.Bool(ptr + 736 + 96 + 24, false);
        A.store.Bool(ptr + 736 + 96 + 32, false);
        A.store.Bool(ptr + 736 + 96 + 25, false);

        A.store.Bool(ptr + 736 + 132 + 28, false);
        A.store.Ref(ptr + 736 + 132 + 0, undefined);
        A.store.Ref(ptr + 736 + 132 + 4, undefined);
        A.store.Ref(ptr + 736 + 132 + 8, undefined);
        A.store.Ref(ptr + 736 + 132 + 12, undefined);
        A.store.Ref(ptr + 736 + 132 + 16, undefined);
        A.store.Ref(ptr + 736 + 132 + 20, undefined);
        A.store.Bool(ptr + 736 + 132 + 26, false);
        A.store.Bool(ptr + 736 + 132 + 24, false);
        A.store.Bool(ptr + 736 + 132 + 27, false);
        A.store.Bool(ptr + 736 + 132 + 25, false);

        A.store.Bool(ptr + 736 + 164 + 28, false);
        A.store.Ref(ptr + 736 + 164 + 0, undefined);
        A.store.Ref(ptr + 736 + 164 + 4, undefined);
        A.store.Ref(ptr + 736 + 164 + 8, undefined);
        A.store.Ref(ptr + 736 + 164 + 12, undefined);
        A.store.Ref(ptr + 736 + 164 + 16, undefined);
        A.store.Ref(ptr + 736 + 164 + 20, undefined);
        A.store.Bool(ptr + 736 + 164 + 26, false);
        A.store.Bool(ptr + 736 + 164 + 24, false);
        A.store.Bool(ptr + 736 + 164 + 27, false);
        A.store.Bool(ptr + 736 + 164 + 25, false);

        A.store.Bool(ptr + 932 + 37, false);
        A.store.Ref(ptr + 932 + 0, undefined);
        A.store.Ref(ptr + 932 + 4, undefined);
        A.store.Ref(ptr + 932 + 8, undefined);
        A.store.Ref(ptr + 932 + 12, undefined);
        A.store.Ref(ptr + 932 + 16, undefined);
        A.store.Ref(ptr + 932 + 20, undefined);
        A.store.Bool(ptr + 932 + 36, false);
        A.store.Int32(ptr + 932 + 24, 0);
        A.store.Ref(ptr + 932 + 28, undefined);
        A.store.Ref(ptr + 932 + 32, undefined);
        A.store.Ref(ptr + 972, undefined);
        A.store.Enum(ptr + 976, -1);

        A.store.Bool(ptr + 980 + 85, false);

        A.store.Bool(ptr + 980 + 0 + 21, false);
        A.store.Bool(ptr + 980 + 0 + 14, false);
        A.store.Bool(ptr + 980 + 0 + 0, false);
        A.store.Ref(ptr + 980 + 0 + 4, undefined);
        A.store.Bool(ptr + 980 + 0 + 15, false);
        A.store.Bool(ptr + 980 + 0 + 8, false);
        A.store.Bool(ptr + 980 + 0 + 16, false);
        A.store.Bool(ptr + 980 + 0 + 9, false);
        A.store.Bool(ptr + 980 + 0 + 17, false);
        A.store.Bool(ptr + 980 + 0 + 10, false);
        A.store.Bool(ptr + 980 + 0 + 18, false);
        A.store.Bool(ptr + 980 + 0 + 11, false);
        A.store.Bool(ptr + 980 + 0 + 19, false);
        A.store.Bool(ptr + 980 + 0 + 12, false);
        A.store.Bool(ptr + 980 + 0 + 20, false);
        A.store.Bool(ptr + 980 + 0 + 13, false);

        A.store.Bool(ptr + 980 + 24 + 28, false);
        A.store.Ref(ptr + 980 + 24 + 0, undefined);
        A.store.Ref(ptr + 980 + 24 + 4, undefined);
        A.store.Ref(ptr + 980 + 24 + 8, undefined);
        A.store.Ref(ptr + 980 + 24 + 12, undefined);
        A.store.Ref(ptr + 980 + 24 + 16, undefined);
        A.store.Ref(ptr + 980 + 24 + 20, undefined);
        A.store.Bool(ptr + 980 + 24 + 26, false);
        A.store.Bool(ptr + 980 + 24 + 24, false);
        A.store.Bool(ptr + 980 + 24 + 27, false);
        A.store.Bool(ptr + 980 + 24 + 25, false);

        A.store.Bool(ptr + 980 + 56 + 28, false);
        A.store.Ref(ptr + 980 + 56 + 0, undefined);
        A.store.Ref(ptr + 980 + 56 + 4, undefined);
        A.store.Ref(ptr + 980 + 56 + 8, undefined);
        A.store.Ref(ptr + 980 + 56 + 12, undefined);
        A.store.Ref(ptr + 980 + 56 + 16, undefined);
        A.store.Ref(ptr + 980 + 56 + 20, undefined);
        A.store.Bool(ptr + 980 + 56 + 26, false);
        A.store.Bool(ptr + 980 + 56 + 24, false);
        A.store.Bool(ptr + 980 + 56 + 27, false);
        A.store.Bool(ptr + 980 + 56 + 25, false);

        A.store.Bool(ptr + 1068 + 222, false);

        A.store.Bool(ptr + 1068 + 0 + 21, false);
        A.store.Bool(ptr + 1068 + 0 + 14, false);
        A.store.Bool(ptr + 1068 + 0 + 0, false);
        A.store.Ref(ptr + 1068 + 0 + 4, undefined);
        A.store.Bool(ptr + 1068 + 0 + 15, false);
        A.store.Bool(ptr + 1068 + 0 + 8, false);
        A.store.Bool(ptr + 1068 + 0 + 16, false);
        A.store.Bool(ptr + 1068 + 0 + 9, false);
        A.store.Bool(ptr + 1068 + 0 + 17, false);
        A.store.Bool(ptr + 1068 + 0 + 10, false);
        A.store.Bool(ptr + 1068 + 0 + 18, false);
        A.store.Bool(ptr + 1068 + 0 + 11, false);
        A.store.Bool(ptr + 1068 + 0 + 19, false);
        A.store.Bool(ptr + 1068 + 0 + 12, false);
        A.store.Bool(ptr + 1068 + 0 + 20, false);
        A.store.Bool(ptr + 1068 + 0 + 13, false);

        A.store.Bool(ptr + 1068 + 24 + 21, false);
        A.store.Bool(ptr + 1068 + 24 + 14, false);
        A.store.Bool(ptr + 1068 + 24 + 0, false);
        A.store.Ref(ptr + 1068 + 24 + 4, undefined);
        A.store.Bool(ptr + 1068 + 24 + 15, false);
        A.store.Bool(ptr + 1068 + 24 + 8, false);
        A.store.Bool(ptr + 1068 + 24 + 16, false);
        A.store.Bool(ptr + 1068 + 24 + 9, false);
        A.store.Bool(ptr + 1068 + 24 + 17, false);
        A.store.Bool(ptr + 1068 + 24 + 10, false);
        A.store.Bool(ptr + 1068 + 24 + 18, false);
        A.store.Bool(ptr + 1068 + 24 + 11, false);
        A.store.Bool(ptr + 1068 + 24 + 19, false);
        A.store.Bool(ptr + 1068 + 24 + 12, false);
        A.store.Bool(ptr + 1068 + 24 + 20, false);
        A.store.Bool(ptr + 1068 + 24 + 13, false);
        A.store.Ref(ptr + 1068 + 48, undefined);
        A.store.Bool(ptr + 1068 + 220, false);
        A.store.Int32(ptr + 1068 + 52, 0);
        A.store.Ref(ptr + 1068 + 56, undefined);

        A.store.Bool(ptr + 1068 + 60 + 28, false);
        A.store.Ref(ptr + 1068 + 60 + 0, undefined);
        A.store.Ref(ptr + 1068 + 60 + 4, undefined);
        A.store.Ref(ptr + 1068 + 60 + 8, undefined);
        A.store.Ref(ptr + 1068 + 60 + 12, undefined);
        A.store.Ref(ptr + 1068 + 60 + 16, undefined);
        A.store.Ref(ptr + 1068 + 60 + 20, undefined);
        A.store.Bool(ptr + 1068 + 60 + 26, false);
        A.store.Bool(ptr + 1068 + 60 + 24, false);
        A.store.Bool(ptr + 1068 + 60 + 27, false);
        A.store.Bool(ptr + 1068 + 60 + 25, false);

        A.store.Bool(ptr + 1068 + 92 + 21, false);
        A.store.Bool(ptr + 1068 + 92 + 14, false);
        A.store.Bool(ptr + 1068 + 92 + 0, false);
        A.store.Ref(ptr + 1068 + 92 + 4, undefined);
        A.store.Bool(ptr + 1068 + 92 + 15, false);
        A.store.Bool(ptr + 1068 + 92 + 8, false);
        A.store.Bool(ptr + 1068 + 92 + 16, false);
        A.store.Bool(ptr + 1068 + 92 + 9, false);
        A.store.Bool(ptr + 1068 + 92 + 17, false);
        A.store.Bool(ptr + 1068 + 92 + 10, false);
        A.store.Bool(ptr + 1068 + 92 + 18, false);
        A.store.Bool(ptr + 1068 + 92 + 11, false);
        A.store.Bool(ptr + 1068 + 92 + 19, false);
        A.store.Bool(ptr + 1068 + 92 + 12, false);
        A.store.Bool(ptr + 1068 + 92 + 20, false);
        A.store.Bool(ptr + 1068 + 92 + 13, false);

        A.store.Bool(ptr + 1068 + 116 + 33, false);
        A.store.Bool(ptr + 1068 + 116 + 26, false);
        A.store.Int32(ptr + 1068 + 116 + 0, 0);
        A.store.Ref(ptr + 1068 + 116 + 4, undefined);
        A.store.Bool(ptr + 1068 + 116 + 27, false);
        A.store.Int32(ptr + 1068 + 116 + 8, 0);
        A.store.Bool(ptr + 1068 + 116 + 28, false);
        A.store.Int32(ptr + 1068 + 116 + 12, 0);
        A.store.Bool(ptr + 1068 + 116 + 29, false);
        A.store.Int32(ptr + 1068 + 116 + 16, 0);
        A.store.Bool(ptr + 1068 + 116 + 30, false);
        A.store.Int32(ptr + 1068 + 116 + 20, 0);
        A.store.Bool(ptr + 1068 + 116 + 31, false);
        A.store.Bool(ptr + 1068 + 116 + 24, false);
        A.store.Bool(ptr + 1068 + 116 + 32, false);
        A.store.Bool(ptr + 1068 + 116 + 25, false);

        A.store.Bool(ptr + 1068 + 152 + 28, false);
        A.store.Ref(ptr + 1068 + 152 + 0, undefined);
        A.store.Ref(ptr + 1068 + 152 + 4, undefined);
        A.store.Ref(ptr + 1068 + 152 + 8, undefined);
        A.store.Ref(ptr + 1068 + 152 + 12, undefined);
        A.store.Ref(ptr + 1068 + 152 + 16, undefined);
        A.store.Ref(ptr + 1068 + 152 + 20, undefined);
        A.store.Bool(ptr + 1068 + 152 + 26, false);
        A.store.Bool(ptr + 1068 + 152 + 24, false);
        A.store.Bool(ptr + 1068 + 152 + 27, false);
        A.store.Bool(ptr + 1068 + 152 + 25, false);

        A.store.Bool(ptr + 1068 + 184 + 28, false);
        A.store.Ref(ptr + 1068 + 184 + 0, undefined);
        A.store.Ref(ptr + 1068 + 184 + 4, undefined);
        A.store.Ref(ptr + 1068 + 184 + 8, undefined);
        A.store.Ref(ptr + 1068 + 184 + 12, undefined);
        A.store.Ref(ptr + 1068 + 184 + 16, undefined);
        A.store.Ref(ptr + 1068 + 184 + 20, undefined);
        A.store.Bool(ptr + 1068 + 184 + 26, false);
        A.store.Bool(ptr + 1068 + 184 + 24, false);
        A.store.Bool(ptr + 1068 + 184 + 27, false);
        A.store.Bool(ptr + 1068 + 184 + 25, false);
        A.store.Bool(ptr + 1068 + 221, false);
        A.store.Int32(ptr + 1068 + 216, 0);
      } else {
        A.store.Bool(ptr + 1293, true);

        if (typeof x["Cellular"] === "undefined") {
          A.store.Bool(ptr + 0 + 134, false);

          A.store.Bool(ptr + 0 + 0 + 21, false);
          A.store.Bool(ptr + 0 + 0 + 14, false);
          A.store.Bool(ptr + 0 + 0 + 0, false);
          A.store.Ref(ptr + 0 + 0 + 4, undefined);
          A.store.Bool(ptr + 0 + 0 + 15, false);
          A.store.Bool(ptr + 0 + 0 + 8, false);
          A.store.Bool(ptr + 0 + 0 + 16, false);
          A.store.Bool(ptr + 0 + 0 + 9, false);
          A.store.Bool(ptr + 0 + 0 + 17, false);
          A.store.Bool(ptr + 0 + 0 + 10, false);
          A.store.Bool(ptr + 0 + 0 + 18, false);
          A.store.Bool(ptr + 0 + 0 + 11, false);
          A.store.Bool(ptr + 0 + 0 + 19, false);
          A.store.Bool(ptr + 0 + 0 + 12, false);
          A.store.Bool(ptr + 0 + 0 + 20, false);
          A.store.Bool(ptr + 0 + 0 + 13, false);
          A.store.Ref(ptr + 0 + 24, undefined);
          A.store.Enum(ptr + 0 + 28, -1);
          A.store.Bool(ptr + 0 + 129, false);
          A.store.Bool(ptr + 0 + 32, false);
          A.store.Ref(ptr + 0 + 36, undefined);
          A.store.Ref(ptr + 0 + 40, undefined);
          A.store.Ref(ptr + 0 + 44, undefined);
          A.store.Ref(ptr + 0 + 48, undefined);
          A.store.Ref(ptr + 0 + 52, undefined);
          A.store.Ref(ptr + 0 + 56, undefined);
          A.store.Ref(ptr + 0 + 60, undefined);
          A.store.Ref(ptr + 0 + 64, undefined);

          A.store.Bool(ptr + 0 + 68 + 12, false);
          A.store.Ref(ptr + 0 + 68 + 0, undefined);
          A.store.Ref(ptr + 0 + 68 + 4, undefined);
          A.store.Ref(ptr + 0 + 68 + 8, undefined);
          A.store.Ref(ptr + 0 + 84, undefined);
          A.store.Bool(ptr + 0 + 130, false);
          A.store.Bool(ptr + 0 + 88, false);

          A.store.Bool(ptr + 0 + 92 + 12, false);
          A.store.Ref(ptr + 0 + 92 + 0, undefined);
          A.store.Ref(ptr + 0 + 92 + 4, undefined);
          A.store.Ref(ptr + 0 + 92 + 8, undefined);

          A.store.Bool(ptr + 0 + 108 + 14, false);
          A.store.Ref(ptr + 0 + 108 + 0, undefined);
          A.store.Bool(ptr + 0 + 108 + 12, false);
          A.store.Bool(ptr + 0 + 108 + 4, false);
          A.store.Bool(ptr + 0 + 108 + 13, false);
          A.store.Int32(ptr + 0 + 108 + 8, 0);
          A.store.Bool(ptr + 0 + 131, false);
          A.store.Bool(ptr + 0 + 123, false);
          A.store.Bool(ptr + 0 + 132, false);
          A.store.Int32(ptr + 0 + 124, 0);
          A.store.Bool(ptr + 0 + 133, false);
          A.store.Bool(ptr + 0 + 128, false);
        } else {
          A.store.Bool(ptr + 0 + 134, true);

          if (typeof x["Cellular"]["AutoConnect"] === "undefined") {
            A.store.Bool(ptr + 0 + 0 + 21, false);
            A.store.Bool(ptr + 0 + 0 + 14, false);
            A.store.Bool(ptr + 0 + 0 + 0, false);
            A.store.Ref(ptr + 0 + 0 + 4, undefined);
            A.store.Bool(ptr + 0 + 0 + 15, false);
            A.store.Bool(ptr + 0 + 0 + 8, false);
            A.store.Bool(ptr + 0 + 0 + 16, false);
            A.store.Bool(ptr + 0 + 0 + 9, false);
            A.store.Bool(ptr + 0 + 0 + 17, false);
            A.store.Bool(ptr + 0 + 0 + 10, false);
            A.store.Bool(ptr + 0 + 0 + 18, false);
            A.store.Bool(ptr + 0 + 0 + 11, false);
            A.store.Bool(ptr + 0 + 0 + 19, false);
            A.store.Bool(ptr + 0 + 0 + 12, false);
            A.store.Bool(ptr + 0 + 0 + 20, false);
            A.store.Bool(ptr + 0 + 0 + 13, false);
          } else {
            A.store.Bool(ptr + 0 + 0 + 21, true);
            A.store.Bool(ptr + 0 + 0 + 14, "Active" in x["Cellular"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 0 + 0 + 0, x["Cellular"]["AutoConnect"]["Active"] ? true : false);
            A.store.Ref(ptr + 0 + 0 + 4, x["Cellular"]["AutoConnect"]["Effective"]);
            A.store.Bool(ptr + 0 + 0 + 15, "UserPolicy" in x["Cellular"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 0 + 0 + 8, x["Cellular"]["AutoConnect"]["UserPolicy"] ? true : false);
            A.store.Bool(ptr + 0 + 0 + 16, "DevicePolicy" in x["Cellular"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 0 + 0 + 9, x["Cellular"]["AutoConnect"]["DevicePolicy"] ? true : false);
            A.store.Bool(ptr + 0 + 0 + 17, "UserSetting" in x["Cellular"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 0 + 0 + 10, x["Cellular"]["AutoConnect"]["UserSetting"] ? true : false);
            A.store.Bool(ptr + 0 + 0 + 18, "SharedSetting" in x["Cellular"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 0 + 0 + 11, x["Cellular"]["AutoConnect"]["SharedSetting"] ? true : false);
            A.store.Bool(ptr + 0 + 0 + 19, "UserEditable" in x["Cellular"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 0 + 0 + 12, x["Cellular"]["AutoConnect"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 0 + 0 + 20, "DeviceEditable" in x["Cellular"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 0 + 0 + 13, x["Cellular"]["AutoConnect"]["DeviceEditable"] ? true : false);
          }
          A.store.Ref(ptr + 0 + 24, x["Cellular"]["ActivationType"]);
          A.store.Enum(
            ptr + 0 + 28,
            ["Activated", "Activating", "NotActivated", "PartiallyActivated"].indexOf(
              x["Cellular"]["ActivationState"] as string
            )
          );
          A.store.Bool(ptr + 0 + 129, "AllowRoaming" in x["Cellular"] ? true : false);
          A.store.Bool(ptr + 0 + 32, x["Cellular"]["AllowRoaming"] ? true : false);
          A.store.Ref(ptr + 0 + 36, x["Cellular"]["Family"]);
          A.store.Ref(ptr + 0 + 40, x["Cellular"]["FirmwareRevision"]);
          A.store.Ref(ptr + 0 + 44, x["Cellular"]["FoundNetworks"]);
          A.store.Ref(ptr + 0 + 48, x["Cellular"]["HardwareRevision"]);
          A.store.Ref(ptr + 0 + 52, x["Cellular"]["HomeProvider"]);
          A.store.Ref(ptr + 0 + 56, x["Cellular"]["Manufacturer"]);
          A.store.Ref(ptr + 0 + 60, x["Cellular"]["ModelID"]);
          A.store.Ref(ptr + 0 + 64, x["Cellular"]["NetworkTechnology"]);

          if (typeof x["Cellular"]["PaymentPortal"] === "undefined") {
            A.store.Bool(ptr + 0 + 68 + 12, false);
            A.store.Ref(ptr + 0 + 68 + 0, undefined);
            A.store.Ref(ptr + 0 + 68 + 4, undefined);
            A.store.Ref(ptr + 0 + 68 + 8, undefined);
          } else {
            A.store.Bool(ptr + 0 + 68 + 12, true);
            A.store.Ref(ptr + 0 + 68 + 0, x["Cellular"]["PaymentPortal"]["Method"]);
            A.store.Ref(ptr + 0 + 68 + 4, x["Cellular"]["PaymentPortal"]["PostData"]);
            A.store.Ref(ptr + 0 + 68 + 8, x["Cellular"]["PaymentPortal"]["Url"]);
          }
          A.store.Ref(ptr + 0 + 84, x["Cellular"]["RoamingState"]);
          A.store.Bool(ptr + 0 + 130, "Scanning" in x["Cellular"] ? true : false);
          A.store.Bool(ptr + 0 + 88, x["Cellular"]["Scanning"] ? true : false);

          if (typeof x["Cellular"]["ServingOperator"] === "undefined") {
            A.store.Bool(ptr + 0 + 92 + 12, false);
            A.store.Ref(ptr + 0 + 92 + 0, undefined);
            A.store.Ref(ptr + 0 + 92 + 4, undefined);
            A.store.Ref(ptr + 0 + 92 + 8, undefined);
          } else {
            A.store.Bool(ptr + 0 + 92 + 12, true);
            A.store.Ref(ptr + 0 + 92 + 0, x["Cellular"]["ServingOperator"]["Name"]);
            A.store.Ref(ptr + 0 + 92 + 4, x["Cellular"]["ServingOperator"]["Code"]);
            A.store.Ref(ptr + 0 + 92 + 8, x["Cellular"]["ServingOperator"]["Country"]);
          }

          if (typeof x["Cellular"]["SIMLockStatus"] === "undefined") {
            A.store.Bool(ptr + 0 + 108 + 14, false);
            A.store.Ref(ptr + 0 + 108 + 0, undefined);
            A.store.Bool(ptr + 0 + 108 + 12, false);
            A.store.Bool(ptr + 0 + 108 + 4, false);
            A.store.Bool(ptr + 0 + 108 + 13, false);
            A.store.Int32(ptr + 0 + 108 + 8, 0);
          } else {
            A.store.Bool(ptr + 0 + 108 + 14, true);
            A.store.Ref(ptr + 0 + 108 + 0, x["Cellular"]["SIMLockStatus"]["LockType"]);
            A.store.Bool(ptr + 0 + 108 + 12, "LockEnabled" in x["Cellular"]["SIMLockStatus"] ? true : false);
            A.store.Bool(ptr + 0 + 108 + 4, x["Cellular"]["SIMLockStatus"]["LockEnabled"] ? true : false);
            A.store.Bool(ptr + 0 + 108 + 13, "RetriesLeft" in x["Cellular"]["SIMLockStatus"] ? true : false);
            A.store.Int32(
              ptr + 0 + 108 + 8,
              x["Cellular"]["SIMLockStatus"]["RetriesLeft"] === undefined
                ? 0
                : (x["Cellular"]["SIMLockStatus"]["RetriesLeft"] as number)
            );
          }
          A.store.Bool(ptr + 0 + 131, "SIMPresent" in x["Cellular"] ? true : false);
          A.store.Bool(ptr + 0 + 123, x["Cellular"]["SIMPresent"] ? true : false);
          A.store.Bool(ptr + 0 + 132, "SignalStrength" in x["Cellular"] ? true : false);
          A.store.Int32(
            ptr + 0 + 124,
            x["Cellular"]["SignalStrength"] === undefined ? 0 : (x["Cellular"]["SignalStrength"] as number)
          );
          A.store.Bool(ptr + 0 + 133, "SupportNetworkScan" in x["Cellular"] ? true : false);
          A.store.Bool(ptr + 0 + 128, x["Cellular"]["SupportNetworkScan"] ? true : false);
        }
        A.store.Bool(ptr + 1291, "Connectable" in x ? true : false);
        A.store.Bool(ptr + 135, x["Connectable"] ? true : false);
        A.store.Enum(ptr + 136, ["Connected", "Connecting", "NotConnected"].indexOf(x["ConnectionState"] as string));
        A.store.Ref(ptr + 140, x["ErrorState"]);

        if (typeof x["Ethernet"] === "undefined") {
          A.store.Bool(ptr + 144 + 53, false);

          A.store.Bool(ptr + 144 + 0 + 21, false);
          A.store.Bool(ptr + 144 + 0 + 14, false);
          A.store.Bool(ptr + 144 + 0 + 0, false);
          A.store.Ref(ptr + 144 + 0 + 4, undefined);
          A.store.Bool(ptr + 144 + 0 + 15, false);
          A.store.Bool(ptr + 144 + 0 + 8, false);
          A.store.Bool(ptr + 144 + 0 + 16, false);
          A.store.Bool(ptr + 144 + 0 + 9, false);
          A.store.Bool(ptr + 144 + 0 + 17, false);
          A.store.Bool(ptr + 144 + 0 + 10, false);
          A.store.Bool(ptr + 144 + 0 + 18, false);
          A.store.Bool(ptr + 144 + 0 + 11, false);
          A.store.Bool(ptr + 144 + 0 + 19, false);
          A.store.Bool(ptr + 144 + 0 + 12, false);
          A.store.Bool(ptr + 144 + 0 + 20, false);
          A.store.Bool(ptr + 144 + 0 + 13, false);

          A.store.Bool(ptr + 144 + 24 + 28, false);
          A.store.Ref(ptr + 144 + 24 + 0, undefined);
          A.store.Ref(ptr + 144 + 24 + 4, undefined);
          A.store.Ref(ptr + 144 + 24 + 8, undefined);
          A.store.Ref(ptr + 144 + 24 + 12, undefined);
          A.store.Ref(ptr + 144 + 24 + 16, undefined);
          A.store.Ref(ptr + 144 + 24 + 20, undefined);
          A.store.Bool(ptr + 144 + 24 + 26, false);
          A.store.Bool(ptr + 144 + 24 + 24, false);
          A.store.Bool(ptr + 144 + 24 + 27, false);
          A.store.Bool(ptr + 144 + 24 + 25, false);
        } else {
          A.store.Bool(ptr + 144 + 53, true);

          if (typeof x["Ethernet"]["AutoConnect"] === "undefined") {
            A.store.Bool(ptr + 144 + 0 + 21, false);
            A.store.Bool(ptr + 144 + 0 + 14, false);
            A.store.Bool(ptr + 144 + 0 + 0, false);
            A.store.Ref(ptr + 144 + 0 + 4, undefined);
            A.store.Bool(ptr + 144 + 0 + 15, false);
            A.store.Bool(ptr + 144 + 0 + 8, false);
            A.store.Bool(ptr + 144 + 0 + 16, false);
            A.store.Bool(ptr + 144 + 0 + 9, false);
            A.store.Bool(ptr + 144 + 0 + 17, false);
            A.store.Bool(ptr + 144 + 0 + 10, false);
            A.store.Bool(ptr + 144 + 0 + 18, false);
            A.store.Bool(ptr + 144 + 0 + 11, false);
            A.store.Bool(ptr + 144 + 0 + 19, false);
            A.store.Bool(ptr + 144 + 0 + 12, false);
            A.store.Bool(ptr + 144 + 0 + 20, false);
            A.store.Bool(ptr + 144 + 0 + 13, false);
          } else {
            A.store.Bool(ptr + 144 + 0 + 21, true);
            A.store.Bool(ptr + 144 + 0 + 14, "Active" in x["Ethernet"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 144 + 0 + 0, x["Ethernet"]["AutoConnect"]["Active"] ? true : false);
            A.store.Ref(ptr + 144 + 0 + 4, x["Ethernet"]["AutoConnect"]["Effective"]);
            A.store.Bool(ptr + 144 + 0 + 15, "UserPolicy" in x["Ethernet"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 144 + 0 + 8, x["Ethernet"]["AutoConnect"]["UserPolicy"] ? true : false);
            A.store.Bool(ptr + 144 + 0 + 16, "DevicePolicy" in x["Ethernet"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 144 + 0 + 9, x["Ethernet"]["AutoConnect"]["DevicePolicy"] ? true : false);
            A.store.Bool(ptr + 144 + 0 + 17, "UserSetting" in x["Ethernet"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 144 + 0 + 10, x["Ethernet"]["AutoConnect"]["UserSetting"] ? true : false);
            A.store.Bool(ptr + 144 + 0 + 18, "SharedSetting" in x["Ethernet"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 144 + 0 + 11, x["Ethernet"]["AutoConnect"]["SharedSetting"] ? true : false);
            A.store.Bool(ptr + 144 + 0 + 19, "UserEditable" in x["Ethernet"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 144 + 0 + 12, x["Ethernet"]["AutoConnect"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 144 + 0 + 20, "DeviceEditable" in x["Ethernet"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 144 + 0 + 13, x["Ethernet"]["AutoConnect"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["Ethernet"]["Authentication"] === "undefined") {
            A.store.Bool(ptr + 144 + 24 + 28, false);
            A.store.Ref(ptr + 144 + 24 + 0, undefined);
            A.store.Ref(ptr + 144 + 24 + 4, undefined);
            A.store.Ref(ptr + 144 + 24 + 8, undefined);
            A.store.Ref(ptr + 144 + 24 + 12, undefined);
            A.store.Ref(ptr + 144 + 24 + 16, undefined);
            A.store.Ref(ptr + 144 + 24 + 20, undefined);
            A.store.Bool(ptr + 144 + 24 + 26, false);
            A.store.Bool(ptr + 144 + 24 + 24, false);
            A.store.Bool(ptr + 144 + 24 + 27, false);
            A.store.Bool(ptr + 144 + 24 + 25, false);
          } else {
            A.store.Bool(ptr + 144 + 24 + 28, true);
            A.store.Ref(ptr + 144 + 24 + 0, x["Ethernet"]["Authentication"]["Active"]);
            A.store.Ref(ptr + 144 + 24 + 4, x["Ethernet"]["Authentication"]["Effective"]);
            A.store.Ref(ptr + 144 + 24 + 8, x["Ethernet"]["Authentication"]["UserPolicy"]);
            A.store.Ref(ptr + 144 + 24 + 12, x["Ethernet"]["Authentication"]["DevicePolicy"]);
            A.store.Ref(ptr + 144 + 24 + 16, x["Ethernet"]["Authentication"]["UserSetting"]);
            A.store.Ref(ptr + 144 + 24 + 20, x["Ethernet"]["Authentication"]["SharedSetting"]);
            A.store.Bool(ptr + 144 + 24 + 26, "UserEditable" in x["Ethernet"]["Authentication"] ? true : false);
            A.store.Bool(ptr + 144 + 24 + 24, x["Ethernet"]["Authentication"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 144 + 24 + 27, "DeviceEditable" in x["Ethernet"]["Authentication"] ? true : false);
            A.store.Bool(ptr + 144 + 24 + 25, x["Ethernet"]["Authentication"]["DeviceEditable"] ? true : false);
          }
        }
        A.store.Ref(ptr + 200, x["GUID"]);

        if (typeof x["IPAddressConfigType"] === "undefined") {
          A.store.Bool(ptr + 204 + 28, false);
          A.store.Enum(ptr + 204 + 0, -1);
          A.store.Ref(ptr + 204 + 4, undefined);
          A.store.Enum(ptr + 204 + 8, -1);
          A.store.Enum(ptr + 204 + 12, -1);
          A.store.Enum(ptr + 204 + 16, -1);
          A.store.Enum(ptr + 204 + 20, -1);
          A.store.Bool(ptr + 204 + 26, false);
          A.store.Bool(ptr + 204 + 24, false);
          A.store.Bool(ptr + 204 + 27, false);
          A.store.Bool(ptr + 204 + 25, false);
        } else {
          A.store.Bool(ptr + 204 + 28, true);
          A.store.Enum(ptr + 204 + 0, ["DHCP", "Static"].indexOf(x["IPAddressConfigType"]["Active"] as string));
          A.store.Ref(ptr + 204 + 4, x["IPAddressConfigType"]["Effective"]);
          A.store.Enum(ptr + 204 + 8, ["DHCP", "Static"].indexOf(x["IPAddressConfigType"]["UserPolicy"] as string));
          A.store.Enum(ptr + 204 + 12, ["DHCP", "Static"].indexOf(x["IPAddressConfigType"]["DevicePolicy"] as string));
          A.store.Enum(ptr + 204 + 16, ["DHCP", "Static"].indexOf(x["IPAddressConfigType"]["UserSetting"] as string));
          A.store.Enum(ptr + 204 + 20, ["DHCP", "Static"].indexOf(x["IPAddressConfigType"]["SharedSetting"] as string));
          A.store.Bool(ptr + 204 + 26, "UserEditable" in x["IPAddressConfigType"] ? true : false);
          A.store.Bool(ptr + 204 + 24, x["IPAddressConfigType"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 204 + 27, "DeviceEditable" in x["IPAddressConfigType"] ? true : false);
          A.store.Bool(ptr + 204 + 25, x["IPAddressConfigType"]["DeviceEditable"] ? true : false);
        }
        A.store.Ref(ptr + 236, x["IPConfigs"]);
        A.store.Ref(ptr + 240, x["MacAddress"]);

        if (typeof x["Metered"] === "undefined") {
          A.store.Bool(ptr + 244 + 21, false);
          A.store.Bool(ptr + 244 + 14, false);
          A.store.Bool(ptr + 244 + 0, false);
          A.store.Ref(ptr + 244 + 4, undefined);
          A.store.Bool(ptr + 244 + 15, false);
          A.store.Bool(ptr + 244 + 8, false);
          A.store.Bool(ptr + 244 + 16, false);
          A.store.Bool(ptr + 244 + 9, false);
          A.store.Bool(ptr + 244 + 17, false);
          A.store.Bool(ptr + 244 + 10, false);
          A.store.Bool(ptr + 244 + 18, false);
          A.store.Bool(ptr + 244 + 11, false);
          A.store.Bool(ptr + 244 + 19, false);
          A.store.Bool(ptr + 244 + 12, false);
          A.store.Bool(ptr + 244 + 20, false);
          A.store.Bool(ptr + 244 + 13, false);
        } else {
          A.store.Bool(ptr + 244 + 21, true);
          A.store.Bool(ptr + 244 + 14, "Active" in x["Metered"] ? true : false);
          A.store.Bool(ptr + 244 + 0, x["Metered"]["Active"] ? true : false);
          A.store.Ref(ptr + 244 + 4, x["Metered"]["Effective"]);
          A.store.Bool(ptr + 244 + 15, "UserPolicy" in x["Metered"] ? true : false);
          A.store.Bool(ptr + 244 + 8, x["Metered"]["UserPolicy"] ? true : false);
          A.store.Bool(ptr + 244 + 16, "DevicePolicy" in x["Metered"] ? true : false);
          A.store.Bool(ptr + 244 + 9, x["Metered"]["DevicePolicy"] ? true : false);
          A.store.Bool(ptr + 244 + 17, "UserSetting" in x["Metered"] ? true : false);
          A.store.Bool(ptr + 244 + 10, x["Metered"]["UserSetting"] ? true : false);
          A.store.Bool(ptr + 244 + 18, "SharedSetting" in x["Metered"] ? true : false);
          A.store.Bool(ptr + 244 + 11, x["Metered"]["SharedSetting"] ? true : false);
          A.store.Bool(ptr + 244 + 19, "UserEditable" in x["Metered"] ? true : false);
          A.store.Bool(ptr + 244 + 12, x["Metered"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 244 + 20, "DeviceEditable" in x["Metered"] ? true : false);
          A.store.Bool(ptr + 244 + 13, x["Metered"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["Name"] === "undefined") {
          A.store.Bool(ptr + 268 + 28, false);
          A.store.Ref(ptr + 268 + 0, undefined);
          A.store.Ref(ptr + 268 + 4, undefined);
          A.store.Ref(ptr + 268 + 8, undefined);
          A.store.Ref(ptr + 268 + 12, undefined);
          A.store.Ref(ptr + 268 + 16, undefined);
          A.store.Ref(ptr + 268 + 20, undefined);
          A.store.Bool(ptr + 268 + 26, false);
          A.store.Bool(ptr + 268 + 24, false);
          A.store.Bool(ptr + 268 + 27, false);
          A.store.Bool(ptr + 268 + 25, false);
        } else {
          A.store.Bool(ptr + 268 + 28, true);
          A.store.Ref(ptr + 268 + 0, x["Name"]["Active"]);
          A.store.Ref(ptr + 268 + 4, x["Name"]["Effective"]);
          A.store.Ref(ptr + 268 + 8, x["Name"]["UserPolicy"]);
          A.store.Ref(ptr + 268 + 12, x["Name"]["DevicePolicy"]);
          A.store.Ref(ptr + 268 + 16, x["Name"]["UserSetting"]);
          A.store.Ref(ptr + 268 + 20, x["Name"]["SharedSetting"]);
          A.store.Bool(ptr + 268 + 26, "UserEditable" in x["Name"] ? true : false);
          A.store.Bool(ptr + 268 + 24, x["Name"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 268 + 27, "DeviceEditable" in x["Name"] ? true : false);
          A.store.Bool(ptr + 268 + 25, x["Name"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["NameServersConfigType"] === "undefined") {
          A.store.Bool(ptr + 300 + 28, false);
          A.store.Enum(ptr + 300 + 0, -1);
          A.store.Ref(ptr + 300 + 4, undefined);
          A.store.Enum(ptr + 300 + 8, -1);
          A.store.Enum(ptr + 300 + 12, -1);
          A.store.Enum(ptr + 300 + 16, -1);
          A.store.Enum(ptr + 300 + 20, -1);
          A.store.Bool(ptr + 300 + 26, false);
          A.store.Bool(ptr + 300 + 24, false);
          A.store.Bool(ptr + 300 + 27, false);
          A.store.Bool(ptr + 300 + 25, false);
        } else {
          A.store.Bool(ptr + 300 + 28, true);
          A.store.Enum(ptr + 300 + 0, ["DHCP", "Static"].indexOf(x["NameServersConfigType"]["Active"] as string));
          A.store.Ref(ptr + 300 + 4, x["NameServersConfigType"]["Effective"]);
          A.store.Enum(ptr + 300 + 8, ["DHCP", "Static"].indexOf(x["NameServersConfigType"]["UserPolicy"] as string));
          A.store.Enum(
            ptr + 300 + 12,
            ["DHCP", "Static"].indexOf(x["NameServersConfigType"]["DevicePolicy"] as string)
          );
          A.store.Enum(ptr + 300 + 16, ["DHCP", "Static"].indexOf(x["NameServersConfigType"]["UserSetting"] as string));
          A.store.Enum(
            ptr + 300 + 20,
            ["DHCP", "Static"].indexOf(x["NameServersConfigType"]["SharedSetting"] as string)
          );
          A.store.Bool(ptr + 300 + 26, "UserEditable" in x["NameServersConfigType"] ? true : false);
          A.store.Bool(ptr + 300 + 24, x["NameServersConfigType"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 300 + 27, "DeviceEditable" in x["NameServersConfigType"] ? true : false);
          A.store.Bool(ptr + 300 + 25, x["NameServersConfigType"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["Priority"] === "undefined") {
          A.store.Bool(ptr + 332 + 33, false);
          A.store.Bool(ptr + 332 + 26, false);
          A.store.Int32(ptr + 332 + 0, 0);
          A.store.Ref(ptr + 332 + 4, undefined);
          A.store.Bool(ptr + 332 + 27, false);
          A.store.Int32(ptr + 332 + 8, 0);
          A.store.Bool(ptr + 332 + 28, false);
          A.store.Int32(ptr + 332 + 12, 0);
          A.store.Bool(ptr + 332 + 29, false);
          A.store.Int32(ptr + 332 + 16, 0);
          A.store.Bool(ptr + 332 + 30, false);
          A.store.Int32(ptr + 332 + 20, 0);
          A.store.Bool(ptr + 332 + 31, false);
          A.store.Bool(ptr + 332 + 24, false);
          A.store.Bool(ptr + 332 + 32, false);
          A.store.Bool(ptr + 332 + 25, false);
        } else {
          A.store.Bool(ptr + 332 + 33, true);
          A.store.Bool(ptr + 332 + 26, "Active" in x["Priority"] ? true : false);
          A.store.Int32(ptr + 332 + 0, x["Priority"]["Active"] === undefined ? 0 : (x["Priority"]["Active"] as number));
          A.store.Ref(ptr + 332 + 4, x["Priority"]["Effective"]);
          A.store.Bool(ptr + 332 + 27, "UserPolicy" in x["Priority"] ? true : false);
          A.store.Int32(
            ptr + 332 + 8,
            x["Priority"]["UserPolicy"] === undefined ? 0 : (x["Priority"]["UserPolicy"] as number)
          );
          A.store.Bool(ptr + 332 + 28, "DevicePolicy" in x["Priority"] ? true : false);
          A.store.Int32(
            ptr + 332 + 12,
            x["Priority"]["DevicePolicy"] === undefined ? 0 : (x["Priority"]["DevicePolicy"] as number)
          );
          A.store.Bool(ptr + 332 + 29, "UserSetting" in x["Priority"] ? true : false);
          A.store.Int32(
            ptr + 332 + 16,
            x["Priority"]["UserSetting"] === undefined ? 0 : (x["Priority"]["UserSetting"] as number)
          );
          A.store.Bool(ptr + 332 + 30, "SharedSetting" in x["Priority"] ? true : false);
          A.store.Int32(
            ptr + 332 + 20,
            x["Priority"]["SharedSetting"] === undefined ? 0 : (x["Priority"]["SharedSetting"] as number)
          );
          A.store.Bool(ptr + 332 + 31, "UserEditable" in x["Priority"] ? true : false);
          A.store.Bool(ptr + 332 + 24, x["Priority"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 332 + 32, "DeviceEditable" in x["Priority"] ? true : false);
          A.store.Bool(ptr + 332 + 25, x["Priority"]["DeviceEditable"] ? true : false);
        }

        if (typeof x["ProxySettings"] === "undefined") {
          A.store.Bool(ptr + 368 + 365, false);

          A.store.Bool(ptr + 368 + 0 + 28, false);
          A.store.Enum(ptr + 368 + 0 + 0, -1);
          A.store.Ref(ptr + 368 + 0 + 4, undefined);
          A.store.Enum(ptr + 368 + 0 + 8, -1);
          A.store.Enum(ptr + 368 + 0 + 12, -1);
          A.store.Enum(ptr + 368 + 0 + 16, -1);
          A.store.Enum(ptr + 368 + 0 + 20, -1);
          A.store.Bool(ptr + 368 + 0 + 26, false);
          A.store.Bool(ptr + 368 + 0 + 24, false);
          A.store.Bool(ptr + 368 + 0 + 27, false);
          A.store.Bool(ptr + 368 + 0 + 25, false);

          A.store.Bool(ptr + 368 + 32 + 271, false);

          A.store.Bool(ptr + 368 + 32 + 0 + 66, false);

          A.store.Bool(ptr + 368 + 32 + 0 + 0 + 28, false);
          A.store.Ref(ptr + 368 + 32 + 0 + 0 + 0, undefined);
          A.store.Ref(ptr + 368 + 32 + 0 + 0 + 4, undefined);
          A.store.Ref(ptr + 368 + 32 + 0 + 0 + 8, undefined);
          A.store.Ref(ptr + 368 + 32 + 0 + 0 + 12, undefined);
          A.store.Ref(ptr + 368 + 32 + 0 + 0 + 16, undefined);
          A.store.Ref(ptr + 368 + 32 + 0 + 0 + 20, undefined);
          A.store.Bool(ptr + 368 + 32 + 0 + 0 + 26, false);
          A.store.Bool(ptr + 368 + 32 + 0 + 0 + 24, false);
          A.store.Bool(ptr + 368 + 32 + 0 + 0 + 27, false);
          A.store.Bool(ptr + 368 + 32 + 0 + 0 + 25, false);

          A.store.Bool(ptr + 368 + 32 + 0 + 32 + 33, false);
          A.store.Bool(ptr + 368 + 32 + 0 + 32 + 26, false);
          A.store.Int32(ptr + 368 + 32 + 0 + 32 + 0, 0);
          A.store.Ref(ptr + 368 + 32 + 0 + 32 + 4, undefined);
          A.store.Bool(ptr + 368 + 32 + 0 + 32 + 27, false);
          A.store.Int32(ptr + 368 + 32 + 0 + 32 + 8, 0);
          A.store.Bool(ptr + 368 + 32 + 0 + 32 + 28, false);
          A.store.Int32(ptr + 368 + 32 + 0 + 32 + 12, 0);
          A.store.Bool(ptr + 368 + 32 + 0 + 32 + 29, false);
          A.store.Int32(ptr + 368 + 32 + 0 + 32 + 16, 0);
          A.store.Bool(ptr + 368 + 32 + 0 + 32 + 30, false);
          A.store.Int32(ptr + 368 + 32 + 0 + 32 + 20, 0);
          A.store.Bool(ptr + 368 + 32 + 0 + 32 + 31, false);
          A.store.Bool(ptr + 368 + 32 + 0 + 32 + 24, false);
          A.store.Bool(ptr + 368 + 32 + 0 + 32 + 32, false);
          A.store.Bool(ptr + 368 + 32 + 0 + 32 + 25, false);

          A.store.Bool(ptr + 368 + 32 + 68 + 66, false);

          A.store.Bool(ptr + 368 + 32 + 68 + 0 + 28, false);
          A.store.Ref(ptr + 368 + 32 + 68 + 0 + 0, undefined);
          A.store.Ref(ptr + 368 + 32 + 68 + 0 + 4, undefined);
          A.store.Ref(ptr + 368 + 32 + 68 + 0 + 8, undefined);
          A.store.Ref(ptr + 368 + 32 + 68 + 0 + 12, undefined);
          A.store.Ref(ptr + 368 + 32 + 68 + 0 + 16, undefined);
          A.store.Ref(ptr + 368 + 32 + 68 + 0 + 20, undefined);
          A.store.Bool(ptr + 368 + 32 + 68 + 0 + 26, false);
          A.store.Bool(ptr + 368 + 32 + 68 + 0 + 24, false);
          A.store.Bool(ptr + 368 + 32 + 68 + 0 + 27, false);
          A.store.Bool(ptr + 368 + 32 + 68 + 0 + 25, false);

          A.store.Bool(ptr + 368 + 32 + 68 + 32 + 33, false);
          A.store.Bool(ptr + 368 + 32 + 68 + 32 + 26, false);
          A.store.Int32(ptr + 368 + 32 + 68 + 32 + 0, 0);
          A.store.Ref(ptr + 368 + 32 + 68 + 32 + 4, undefined);
          A.store.Bool(ptr + 368 + 32 + 68 + 32 + 27, false);
          A.store.Int32(ptr + 368 + 32 + 68 + 32 + 8, 0);
          A.store.Bool(ptr + 368 + 32 + 68 + 32 + 28, false);
          A.store.Int32(ptr + 368 + 32 + 68 + 32 + 12, 0);
          A.store.Bool(ptr + 368 + 32 + 68 + 32 + 29, false);
          A.store.Int32(ptr + 368 + 32 + 68 + 32 + 16, 0);
          A.store.Bool(ptr + 368 + 32 + 68 + 32 + 30, false);
          A.store.Int32(ptr + 368 + 32 + 68 + 32 + 20, 0);
          A.store.Bool(ptr + 368 + 32 + 68 + 32 + 31, false);
          A.store.Bool(ptr + 368 + 32 + 68 + 32 + 24, false);
          A.store.Bool(ptr + 368 + 32 + 68 + 32 + 32, false);
          A.store.Bool(ptr + 368 + 32 + 68 + 32 + 25, false);

          A.store.Bool(ptr + 368 + 32 + 136 + 66, false);

          A.store.Bool(ptr + 368 + 32 + 136 + 0 + 28, false);
          A.store.Ref(ptr + 368 + 32 + 136 + 0 + 0, undefined);
          A.store.Ref(ptr + 368 + 32 + 136 + 0 + 4, undefined);
          A.store.Ref(ptr + 368 + 32 + 136 + 0 + 8, undefined);
          A.store.Ref(ptr + 368 + 32 + 136 + 0 + 12, undefined);
          A.store.Ref(ptr + 368 + 32 + 136 + 0 + 16, undefined);
          A.store.Ref(ptr + 368 + 32 + 136 + 0 + 20, undefined);
          A.store.Bool(ptr + 368 + 32 + 136 + 0 + 26, false);
          A.store.Bool(ptr + 368 + 32 + 136 + 0 + 24, false);
          A.store.Bool(ptr + 368 + 32 + 136 + 0 + 27, false);
          A.store.Bool(ptr + 368 + 32 + 136 + 0 + 25, false);

          A.store.Bool(ptr + 368 + 32 + 136 + 32 + 33, false);
          A.store.Bool(ptr + 368 + 32 + 136 + 32 + 26, false);
          A.store.Int32(ptr + 368 + 32 + 136 + 32 + 0, 0);
          A.store.Ref(ptr + 368 + 32 + 136 + 32 + 4, undefined);
          A.store.Bool(ptr + 368 + 32 + 136 + 32 + 27, false);
          A.store.Int32(ptr + 368 + 32 + 136 + 32 + 8, 0);
          A.store.Bool(ptr + 368 + 32 + 136 + 32 + 28, false);
          A.store.Int32(ptr + 368 + 32 + 136 + 32 + 12, 0);
          A.store.Bool(ptr + 368 + 32 + 136 + 32 + 29, false);
          A.store.Int32(ptr + 368 + 32 + 136 + 32 + 16, 0);
          A.store.Bool(ptr + 368 + 32 + 136 + 32 + 30, false);
          A.store.Int32(ptr + 368 + 32 + 136 + 32 + 20, 0);
          A.store.Bool(ptr + 368 + 32 + 136 + 32 + 31, false);
          A.store.Bool(ptr + 368 + 32 + 136 + 32 + 24, false);
          A.store.Bool(ptr + 368 + 32 + 136 + 32 + 32, false);
          A.store.Bool(ptr + 368 + 32 + 136 + 32 + 25, false);

          A.store.Bool(ptr + 368 + 32 + 204 + 66, false);

          A.store.Bool(ptr + 368 + 32 + 204 + 0 + 28, false);
          A.store.Ref(ptr + 368 + 32 + 204 + 0 + 0, undefined);
          A.store.Ref(ptr + 368 + 32 + 204 + 0 + 4, undefined);
          A.store.Ref(ptr + 368 + 32 + 204 + 0 + 8, undefined);
          A.store.Ref(ptr + 368 + 32 + 204 + 0 + 12, undefined);
          A.store.Ref(ptr + 368 + 32 + 204 + 0 + 16, undefined);
          A.store.Ref(ptr + 368 + 32 + 204 + 0 + 20, undefined);
          A.store.Bool(ptr + 368 + 32 + 204 + 0 + 26, false);
          A.store.Bool(ptr + 368 + 32 + 204 + 0 + 24, false);
          A.store.Bool(ptr + 368 + 32 + 204 + 0 + 27, false);
          A.store.Bool(ptr + 368 + 32 + 204 + 0 + 25, false);

          A.store.Bool(ptr + 368 + 32 + 204 + 32 + 33, false);
          A.store.Bool(ptr + 368 + 32 + 204 + 32 + 26, false);
          A.store.Int32(ptr + 368 + 32 + 204 + 32 + 0, 0);
          A.store.Ref(ptr + 368 + 32 + 204 + 32 + 4, undefined);
          A.store.Bool(ptr + 368 + 32 + 204 + 32 + 27, false);
          A.store.Int32(ptr + 368 + 32 + 204 + 32 + 8, 0);
          A.store.Bool(ptr + 368 + 32 + 204 + 32 + 28, false);
          A.store.Int32(ptr + 368 + 32 + 204 + 32 + 12, 0);
          A.store.Bool(ptr + 368 + 32 + 204 + 32 + 29, false);
          A.store.Int32(ptr + 368 + 32 + 204 + 32 + 16, 0);
          A.store.Bool(ptr + 368 + 32 + 204 + 32 + 30, false);
          A.store.Int32(ptr + 368 + 32 + 204 + 32 + 20, 0);
          A.store.Bool(ptr + 368 + 32 + 204 + 32 + 31, false);
          A.store.Bool(ptr + 368 + 32 + 204 + 32 + 24, false);
          A.store.Bool(ptr + 368 + 32 + 204 + 32 + 32, false);
          A.store.Bool(ptr + 368 + 32 + 204 + 32 + 25, false);

          A.store.Bool(ptr + 368 + 304 + 28, false);
          A.store.Ref(ptr + 368 + 304 + 0, undefined);
          A.store.Ref(ptr + 368 + 304 + 4, undefined);
          A.store.Ref(ptr + 368 + 304 + 8, undefined);
          A.store.Ref(ptr + 368 + 304 + 12, undefined);
          A.store.Ref(ptr + 368 + 304 + 16, undefined);
          A.store.Ref(ptr + 368 + 304 + 20, undefined);
          A.store.Bool(ptr + 368 + 304 + 26, false);
          A.store.Bool(ptr + 368 + 304 + 24, false);
          A.store.Bool(ptr + 368 + 304 + 27, false);
          A.store.Bool(ptr + 368 + 304 + 25, false);

          A.store.Bool(ptr + 368 + 336 + 28, false);
          A.store.Ref(ptr + 368 + 336 + 0, undefined);
          A.store.Ref(ptr + 368 + 336 + 4, undefined);
          A.store.Ref(ptr + 368 + 336 + 8, undefined);
          A.store.Ref(ptr + 368 + 336 + 12, undefined);
          A.store.Ref(ptr + 368 + 336 + 16, undefined);
          A.store.Ref(ptr + 368 + 336 + 20, undefined);
          A.store.Bool(ptr + 368 + 336 + 26, false);
          A.store.Bool(ptr + 368 + 336 + 24, false);
          A.store.Bool(ptr + 368 + 336 + 27, false);
          A.store.Bool(ptr + 368 + 336 + 25, false);
        } else {
          A.store.Bool(ptr + 368 + 365, true);

          if (typeof x["ProxySettings"]["Type"] === "undefined") {
            A.store.Bool(ptr + 368 + 0 + 28, false);
            A.store.Enum(ptr + 368 + 0 + 0, -1);
            A.store.Ref(ptr + 368 + 0 + 4, undefined);
            A.store.Enum(ptr + 368 + 0 + 8, -1);
            A.store.Enum(ptr + 368 + 0 + 12, -1);
            A.store.Enum(ptr + 368 + 0 + 16, -1);
            A.store.Enum(ptr + 368 + 0 + 20, -1);
            A.store.Bool(ptr + 368 + 0 + 26, false);
            A.store.Bool(ptr + 368 + 0 + 24, false);
            A.store.Bool(ptr + 368 + 0 + 27, false);
            A.store.Bool(ptr + 368 + 0 + 25, false);
          } else {
            A.store.Bool(ptr + 368 + 0 + 28, true);
            A.store.Enum(
              ptr + 368 + 0 + 0,
              ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["ProxySettings"]["Type"]["Active"] as string)
            );
            A.store.Ref(ptr + 368 + 0 + 4, x["ProxySettings"]["Type"]["Effective"]);
            A.store.Enum(
              ptr + 368 + 0 + 8,
              ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["ProxySettings"]["Type"]["UserPolicy"] as string)
            );
            A.store.Enum(
              ptr + 368 + 0 + 12,
              ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["ProxySettings"]["Type"]["DevicePolicy"] as string)
            );
            A.store.Enum(
              ptr + 368 + 0 + 16,
              ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["ProxySettings"]["Type"]["UserSetting"] as string)
            );
            A.store.Enum(
              ptr + 368 + 0 + 20,
              ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["ProxySettings"]["Type"]["SharedSetting"] as string)
            );
            A.store.Bool(ptr + 368 + 0 + 26, "UserEditable" in x["ProxySettings"]["Type"] ? true : false);
            A.store.Bool(ptr + 368 + 0 + 24, x["ProxySettings"]["Type"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 368 + 0 + 27, "DeviceEditable" in x["ProxySettings"]["Type"] ? true : false);
            A.store.Bool(ptr + 368 + 0 + 25, x["ProxySettings"]["Type"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["ProxySettings"]["Manual"] === "undefined") {
            A.store.Bool(ptr + 368 + 32 + 271, false);

            A.store.Bool(ptr + 368 + 32 + 0 + 66, false);

            A.store.Bool(ptr + 368 + 32 + 0 + 0 + 28, false);
            A.store.Ref(ptr + 368 + 32 + 0 + 0 + 0, undefined);
            A.store.Ref(ptr + 368 + 32 + 0 + 0 + 4, undefined);
            A.store.Ref(ptr + 368 + 32 + 0 + 0 + 8, undefined);
            A.store.Ref(ptr + 368 + 32 + 0 + 0 + 12, undefined);
            A.store.Ref(ptr + 368 + 32 + 0 + 0 + 16, undefined);
            A.store.Ref(ptr + 368 + 32 + 0 + 0 + 20, undefined);
            A.store.Bool(ptr + 368 + 32 + 0 + 0 + 26, false);
            A.store.Bool(ptr + 368 + 32 + 0 + 0 + 24, false);
            A.store.Bool(ptr + 368 + 32 + 0 + 0 + 27, false);
            A.store.Bool(ptr + 368 + 32 + 0 + 0 + 25, false);

            A.store.Bool(ptr + 368 + 32 + 0 + 32 + 33, false);
            A.store.Bool(ptr + 368 + 32 + 0 + 32 + 26, false);
            A.store.Int32(ptr + 368 + 32 + 0 + 32 + 0, 0);
            A.store.Ref(ptr + 368 + 32 + 0 + 32 + 4, undefined);
            A.store.Bool(ptr + 368 + 32 + 0 + 32 + 27, false);
            A.store.Int32(ptr + 368 + 32 + 0 + 32 + 8, 0);
            A.store.Bool(ptr + 368 + 32 + 0 + 32 + 28, false);
            A.store.Int32(ptr + 368 + 32 + 0 + 32 + 12, 0);
            A.store.Bool(ptr + 368 + 32 + 0 + 32 + 29, false);
            A.store.Int32(ptr + 368 + 32 + 0 + 32 + 16, 0);
            A.store.Bool(ptr + 368 + 32 + 0 + 32 + 30, false);
            A.store.Int32(ptr + 368 + 32 + 0 + 32 + 20, 0);
            A.store.Bool(ptr + 368 + 32 + 0 + 32 + 31, false);
            A.store.Bool(ptr + 368 + 32 + 0 + 32 + 24, false);
            A.store.Bool(ptr + 368 + 32 + 0 + 32 + 32, false);
            A.store.Bool(ptr + 368 + 32 + 0 + 32 + 25, false);

            A.store.Bool(ptr + 368 + 32 + 68 + 66, false);

            A.store.Bool(ptr + 368 + 32 + 68 + 0 + 28, false);
            A.store.Ref(ptr + 368 + 32 + 68 + 0 + 0, undefined);
            A.store.Ref(ptr + 368 + 32 + 68 + 0 + 4, undefined);
            A.store.Ref(ptr + 368 + 32 + 68 + 0 + 8, undefined);
            A.store.Ref(ptr + 368 + 32 + 68 + 0 + 12, undefined);
            A.store.Ref(ptr + 368 + 32 + 68 + 0 + 16, undefined);
            A.store.Ref(ptr + 368 + 32 + 68 + 0 + 20, undefined);
            A.store.Bool(ptr + 368 + 32 + 68 + 0 + 26, false);
            A.store.Bool(ptr + 368 + 32 + 68 + 0 + 24, false);
            A.store.Bool(ptr + 368 + 32 + 68 + 0 + 27, false);
            A.store.Bool(ptr + 368 + 32 + 68 + 0 + 25, false);

            A.store.Bool(ptr + 368 + 32 + 68 + 32 + 33, false);
            A.store.Bool(ptr + 368 + 32 + 68 + 32 + 26, false);
            A.store.Int32(ptr + 368 + 32 + 68 + 32 + 0, 0);
            A.store.Ref(ptr + 368 + 32 + 68 + 32 + 4, undefined);
            A.store.Bool(ptr + 368 + 32 + 68 + 32 + 27, false);
            A.store.Int32(ptr + 368 + 32 + 68 + 32 + 8, 0);
            A.store.Bool(ptr + 368 + 32 + 68 + 32 + 28, false);
            A.store.Int32(ptr + 368 + 32 + 68 + 32 + 12, 0);
            A.store.Bool(ptr + 368 + 32 + 68 + 32 + 29, false);
            A.store.Int32(ptr + 368 + 32 + 68 + 32 + 16, 0);
            A.store.Bool(ptr + 368 + 32 + 68 + 32 + 30, false);
            A.store.Int32(ptr + 368 + 32 + 68 + 32 + 20, 0);
            A.store.Bool(ptr + 368 + 32 + 68 + 32 + 31, false);
            A.store.Bool(ptr + 368 + 32 + 68 + 32 + 24, false);
            A.store.Bool(ptr + 368 + 32 + 68 + 32 + 32, false);
            A.store.Bool(ptr + 368 + 32 + 68 + 32 + 25, false);

            A.store.Bool(ptr + 368 + 32 + 136 + 66, false);

            A.store.Bool(ptr + 368 + 32 + 136 + 0 + 28, false);
            A.store.Ref(ptr + 368 + 32 + 136 + 0 + 0, undefined);
            A.store.Ref(ptr + 368 + 32 + 136 + 0 + 4, undefined);
            A.store.Ref(ptr + 368 + 32 + 136 + 0 + 8, undefined);
            A.store.Ref(ptr + 368 + 32 + 136 + 0 + 12, undefined);
            A.store.Ref(ptr + 368 + 32 + 136 + 0 + 16, undefined);
            A.store.Ref(ptr + 368 + 32 + 136 + 0 + 20, undefined);
            A.store.Bool(ptr + 368 + 32 + 136 + 0 + 26, false);
            A.store.Bool(ptr + 368 + 32 + 136 + 0 + 24, false);
            A.store.Bool(ptr + 368 + 32 + 136 + 0 + 27, false);
            A.store.Bool(ptr + 368 + 32 + 136 + 0 + 25, false);

            A.store.Bool(ptr + 368 + 32 + 136 + 32 + 33, false);
            A.store.Bool(ptr + 368 + 32 + 136 + 32 + 26, false);
            A.store.Int32(ptr + 368 + 32 + 136 + 32 + 0, 0);
            A.store.Ref(ptr + 368 + 32 + 136 + 32 + 4, undefined);
            A.store.Bool(ptr + 368 + 32 + 136 + 32 + 27, false);
            A.store.Int32(ptr + 368 + 32 + 136 + 32 + 8, 0);
            A.store.Bool(ptr + 368 + 32 + 136 + 32 + 28, false);
            A.store.Int32(ptr + 368 + 32 + 136 + 32 + 12, 0);
            A.store.Bool(ptr + 368 + 32 + 136 + 32 + 29, false);
            A.store.Int32(ptr + 368 + 32 + 136 + 32 + 16, 0);
            A.store.Bool(ptr + 368 + 32 + 136 + 32 + 30, false);
            A.store.Int32(ptr + 368 + 32 + 136 + 32 + 20, 0);
            A.store.Bool(ptr + 368 + 32 + 136 + 32 + 31, false);
            A.store.Bool(ptr + 368 + 32 + 136 + 32 + 24, false);
            A.store.Bool(ptr + 368 + 32 + 136 + 32 + 32, false);
            A.store.Bool(ptr + 368 + 32 + 136 + 32 + 25, false);

            A.store.Bool(ptr + 368 + 32 + 204 + 66, false);

            A.store.Bool(ptr + 368 + 32 + 204 + 0 + 28, false);
            A.store.Ref(ptr + 368 + 32 + 204 + 0 + 0, undefined);
            A.store.Ref(ptr + 368 + 32 + 204 + 0 + 4, undefined);
            A.store.Ref(ptr + 368 + 32 + 204 + 0 + 8, undefined);
            A.store.Ref(ptr + 368 + 32 + 204 + 0 + 12, undefined);
            A.store.Ref(ptr + 368 + 32 + 204 + 0 + 16, undefined);
            A.store.Ref(ptr + 368 + 32 + 204 + 0 + 20, undefined);
            A.store.Bool(ptr + 368 + 32 + 204 + 0 + 26, false);
            A.store.Bool(ptr + 368 + 32 + 204 + 0 + 24, false);
            A.store.Bool(ptr + 368 + 32 + 204 + 0 + 27, false);
            A.store.Bool(ptr + 368 + 32 + 204 + 0 + 25, false);

            A.store.Bool(ptr + 368 + 32 + 204 + 32 + 33, false);
            A.store.Bool(ptr + 368 + 32 + 204 + 32 + 26, false);
            A.store.Int32(ptr + 368 + 32 + 204 + 32 + 0, 0);
            A.store.Ref(ptr + 368 + 32 + 204 + 32 + 4, undefined);
            A.store.Bool(ptr + 368 + 32 + 204 + 32 + 27, false);
            A.store.Int32(ptr + 368 + 32 + 204 + 32 + 8, 0);
            A.store.Bool(ptr + 368 + 32 + 204 + 32 + 28, false);
            A.store.Int32(ptr + 368 + 32 + 204 + 32 + 12, 0);
            A.store.Bool(ptr + 368 + 32 + 204 + 32 + 29, false);
            A.store.Int32(ptr + 368 + 32 + 204 + 32 + 16, 0);
            A.store.Bool(ptr + 368 + 32 + 204 + 32 + 30, false);
            A.store.Int32(ptr + 368 + 32 + 204 + 32 + 20, 0);
            A.store.Bool(ptr + 368 + 32 + 204 + 32 + 31, false);
            A.store.Bool(ptr + 368 + 32 + 204 + 32 + 24, false);
            A.store.Bool(ptr + 368 + 32 + 204 + 32 + 32, false);
            A.store.Bool(ptr + 368 + 32 + 204 + 32 + 25, false);
          } else {
            A.store.Bool(ptr + 368 + 32 + 271, true);

            if (typeof x["ProxySettings"]["Manual"]["HTTPProxy"] === "undefined") {
              A.store.Bool(ptr + 368 + 32 + 0 + 66, false);

              A.store.Bool(ptr + 368 + 32 + 0 + 0 + 28, false);
              A.store.Ref(ptr + 368 + 32 + 0 + 0 + 0, undefined);
              A.store.Ref(ptr + 368 + 32 + 0 + 0 + 4, undefined);
              A.store.Ref(ptr + 368 + 32 + 0 + 0 + 8, undefined);
              A.store.Ref(ptr + 368 + 32 + 0 + 0 + 12, undefined);
              A.store.Ref(ptr + 368 + 32 + 0 + 0 + 16, undefined);
              A.store.Ref(ptr + 368 + 32 + 0 + 0 + 20, undefined);
              A.store.Bool(ptr + 368 + 32 + 0 + 0 + 26, false);
              A.store.Bool(ptr + 368 + 32 + 0 + 0 + 24, false);
              A.store.Bool(ptr + 368 + 32 + 0 + 0 + 27, false);
              A.store.Bool(ptr + 368 + 32 + 0 + 0 + 25, false);

              A.store.Bool(ptr + 368 + 32 + 0 + 32 + 33, false);
              A.store.Bool(ptr + 368 + 32 + 0 + 32 + 26, false);
              A.store.Int32(ptr + 368 + 32 + 0 + 32 + 0, 0);
              A.store.Ref(ptr + 368 + 32 + 0 + 32 + 4, undefined);
              A.store.Bool(ptr + 368 + 32 + 0 + 32 + 27, false);
              A.store.Int32(ptr + 368 + 32 + 0 + 32 + 8, 0);
              A.store.Bool(ptr + 368 + 32 + 0 + 32 + 28, false);
              A.store.Int32(ptr + 368 + 32 + 0 + 32 + 12, 0);
              A.store.Bool(ptr + 368 + 32 + 0 + 32 + 29, false);
              A.store.Int32(ptr + 368 + 32 + 0 + 32 + 16, 0);
              A.store.Bool(ptr + 368 + 32 + 0 + 32 + 30, false);
              A.store.Int32(ptr + 368 + 32 + 0 + 32 + 20, 0);
              A.store.Bool(ptr + 368 + 32 + 0 + 32 + 31, false);
              A.store.Bool(ptr + 368 + 32 + 0 + 32 + 24, false);
              A.store.Bool(ptr + 368 + 32 + 0 + 32 + 32, false);
              A.store.Bool(ptr + 368 + 32 + 0 + 32 + 25, false);
            } else {
              A.store.Bool(ptr + 368 + 32 + 0 + 66, true);

              if (typeof x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"] === "undefined") {
                A.store.Bool(ptr + 368 + 32 + 0 + 0 + 28, false);
                A.store.Ref(ptr + 368 + 32 + 0 + 0 + 0, undefined);
                A.store.Ref(ptr + 368 + 32 + 0 + 0 + 4, undefined);
                A.store.Ref(ptr + 368 + 32 + 0 + 0 + 8, undefined);
                A.store.Ref(ptr + 368 + 32 + 0 + 0 + 12, undefined);
                A.store.Ref(ptr + 368 + 32 + 0 + 0 + 16, undefined);
                A.store.Ref(ptr + 368 + 32 + 0 + 0 + 20, undefined);
                A.store.Bool(ptr + 368 + 32 + 0 + 0 + 26, false);
                A.store.Bool(ptr + 368 + 32 + 0 + 0 + 24, false);
                A.store.Bool(ptr + 368 + 32 + 0 + 0 + 27, false);
                A.store.Bool(ptr + 368 + 32 + 0 + 0 + 25, false);
              } else {
                A.store.Bool(ptr + 368 + 32 + 0 + 0 + 28, true);
                A.store.Ref(ptr + 368 + 32 + 0 + 0 + 0, x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["Active"]);
                A.store.Ref(ptr + 368 + 32 + 0 + 0 + 4, x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["Effective"]);
                A.store.Ref(
                  ptr + 368 + 32 + 0 + 0 + 8,
                  x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["UserPolicy"]
                );
                A.store.Ref(
                  ptr + 368 + 32 + 0 + 0 + 12,
                  x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["DevicePolicy"]
                );
                A.store.Ref(
                  ptr + 368 + 32 + 0 + 0 + 16,
                  x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["UserSetting"]
                );
                A.store.Ref(
                  ptr + 368 + 32 + 0 + 0 + 20,
                  x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["SharedSetting"]
                );
                A.store.Bool(
                  ptr + 368 + 32 + 0 + 0 + 26,
                  "UserEditable" in x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 0 + 0 + 24,
                  x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["UserEditable"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 0 + 0 + 27,
                  "DeviceEditable" in x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 0 + 0 + 25,
                  x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["DeviceEditable"] ? true : false
                );
              }

              if (typeof x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"] === "undefined") {
                A.store.Bool(ptr + 368 + 32 + 0 + 32 + 33, false);
                A.store.Bool(ptr + 368 + 32 + 0 + 32 + 26, false);
                A.store.Int32(ptr + 368 + 32 + 0 + 32 + 0, 0);
                A.store.Ref(ptr + 368 + 32 + 0 + 32 + 4, undefined);
                A.store.Bool(ptr + 368 + 32 + 0 + 32 + 27, false);
                A.store.Int32(ptr + 368 + 32 + 0 + 32 + 8, 0);
                A.store.Bool(ptr + 368 + 32 + 0 + 32 + 28, false);
                A.store.Int32(ptr + 368 + 32 + 0 + 32 + 12, 0);
                A.store.Bool(ptr + 368 + 32 + 0 + 32 + 29, false);
                A.store.Int32(ptr + 368 + 32 + 0 + 32 + 16, 0);
                A.store.Bool(ptr + 368 + 32 + 0 + 32 + 30, false);
                A.store.Int32(ptr + 368 + 32 + 0 + 32 + 20, 0);
                A.store.Bool(ptr + 368 + 32 + 0 + 32 + 31, false);
                A.store.Bool(ptr + 368 + 32 + 0 + 32 + 24, false);
                A.store.Bool(ptr + 368 + 32 + 0 + 32 + 32, false);
                A.store.Bool(ptr + 368 + 32 + 0 + 32 + 25, false);
              } else {
                A.store.Bool(ptr + 368 + 32 + 0 + 32 + 33, true);
                A.store.Bool(
                  ptr + 368 + 32 + 0 + 32 + 26,
                  "Active" in x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 0 + 32 + 0,
                  x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["Active"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["Active"] as number)
                );
                A.store.Ref(
                  ptr + 368 + 32 + 0 + 32 + 4,
                  x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["Effective"]
                );
                A.store.Bool(
                  ptr + 368 + 32 + 0 + 32 + 27,
                  "UserPolicy" in x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 0 + 32 + 8,
                  x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["UserPolicy"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["UserPolicy"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 0 + 32 + 28,
                  "DevicePolicy" in x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 0 + 32 + 12,
                  x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["DevicePolicy"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["DevicePolicy"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 0 + 32 + 29,
                  "UserSetting" in x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 0 + 32 + 16,
                  x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["UserSetting"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["UserSetting"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 0 + 32 + 30,
                  "SharedSetting" in x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 0 + 32 + 20,
                  x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["SharedSetting"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["SharedSetting"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 0 + 32 + 31,
                  "UserEditable" in x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 0 + 32 + 24,
                  x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["UserEditable"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 0 + 32 + 32,
                  "DeviceEditable" in x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 0 + 32 + 25,
                  x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["DeviceEditable"] ? true : false
                );
              }
            }

            if (typeof x["ProxySettings"]["Manual"]["SecureHTTPProxy"] === "undefined") {
              A.store.Bool(ptr + 368 + 32 + 68 + 66, false);

              A.store.Bool(ptr + 368 + 32 + 68 + 0 + 28, false);
              A.store.Ref(ptr + 368 + 32 + 68 + 0 + 0, undefined);
              A.store.Ref(ptr + 368 + 32 + 68 + 0 + 4, undefined);
              A.store.Ref(ptr + 368 + 32 + 68 + 0 + 8, undefined);
              A.store.Ref(ptr + 368 + 32 + 68 + 0 + 12, undefined);
              A.store.Ref(ptr + 368 + 32 + 68 + 0 + 16, undefined);
              A.store.Ref(ptr + 368 + 32 + 68 + 0 + 20, undefined);
              A.store.Bool(ptr + 368 + 32 + 68 + 0 + 26, false);
              A.store.Bool(ptr + 368 + 32 + 68 + 0 + 24, false);
              A.store.Bool(ptr + 368 + 32 + 68 + 0 + 27, false);
              A.store.Bool(ptr + 368 + 32 + 68 + 0 + 25, false);

              A.store.Bool(ptr + 368 + 32 + 68 + 32 + 33, false);
              A.store.Bool(ptr + 368 + 32 + 68 + 32 + 26, false);
              A.store.Int32(ptr + 368 + 32 + 68 + 32 + 0, 0);
              A.store.Ref(ptr + 368 + 32 + 68 + 32 + 4, undefined);
              A.store.Bool(ptr + 368 + 32 + 68 + 32 + 27, false);
              A.store.Int32(ptr + 368 + 32 + 68 + 32 + 8, 0);
              A.store.Bool(ptr + 368 + 32 + 68 + 32 + 28, false);
              A.store.Int32(ptr + 368 + 32 + 68 + 32 + 12, 0);
              A.store.Bool(ptr + 368 + 32 + 68 + 32 + 29, false);
              A.store.Int32(ptr + 368 + 32 + 68 + 32 + 16, 0);
              A.store.Bool(ptr + 368 + 32 + 68 + 32 + 30, false);
              A.store.Int32(ptr + 368 + 32 + 68 + 32 + 20, 0);
              A.store.Bool(ptr + 368 + 32 + 68 + 32 + 31, false);
              A.store.Bool(ptr + 368 + 32 + 68 + 32 + 24, false);
              A.store.Bool(ptr + 368 + 32 + 68 + 32 + 32, false);
              A.store.Bool(ptr + 368 + 32 + 68 + 32 + 25, false);
            } else {
              A.store.Bool(ptr + 368 + 32 + 68 + 66, true);

              if (typeof x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"] === "undefined") {
                A.store.Bool(ptr + 368 + 32 + 68 + 0 + 28, false);
                A.store.Ref(ptr + 368 + 32 + 68 + 0 + 0, undefined);
                A.store.Ref(ptr + 368 + 32 + 68 + 0 + 4, undefined);
                A.store.Ref(ptr + 368 + 32 + 68 + 0 + 8, undefined);
                A.store.Ref(ptr + 368 + 32 + 68 + 0 + 12, undefined);
                A.store.Ref(ptr + 368 + 32 + 68 + 0 + 16, undefined);
                A.store.Ref(ptr + 368 + 32 + 68 + 0 + 20, undefined);
                A.store.Bool(ptr + 368 + 32 + 68 + 0 + 26, false);
                A.store.Bool(ptr + 368 + 32 + 68 + 0 + 24, false);
                A.store.Bool(ptr + 368 + 32 + 68 + 0 + 27, false);
                A.store.Bool(ptr + 368 + 32 + 68 + 0 + 25, false);
              } else {
                A.store.Bool(ptr + 368 + 32 + 68 + 0 + 28, true);
                A.store.Ref(
                  ptr + 368 + 32 + 68 + 0 + 0,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["Active"]
                );
                A.store.Ref(
                  ptr + 368 + 32 + 68 + 0 + 4,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["Effective"]
                );
                A.store.Ref(
                  ptr + 368 + 32 + 68 + 0 + 8,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["UserPolicy"]
                );
                A.store.Ref(
                  ptr + 368 + 32 + 68 + 0 + 12,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["DevicePolicy"]
                );
                A.store.Ref(
                  ptr + 368 + 32 + 68 + 0 + 16,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["UserSetting"]
                );
                A.store.Ref(
                  ptr + 368 + 32 + 68 + 0 + 20,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["SharedSetting"]
                );
                A.store.Bool(
                  ptr + 368 + 32 + 68 + 0 + 26,
                  "UserEditable" in x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 68 + 0 + 24,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["UserEditable"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 68 + 0 + 27,
                  "DeviceEditable" in x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 68 + 0 + 25,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["DeviceEditable"] ? true : false
                );
              }

              if (typeof x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"] === "undefined") {
                A.store.Bool(ptr + 368 + 32 + 68 + 32 + 33, false);
                A.store.Bool(ptr + 368 + 32 + 68 + 32 + 26, false);
                A.store.Int32(ptr + 368 + 32 + 68 + 32 + 0, 0);
                A.store.Ref(ptr + 368 + 32 + 68 + 32 + 4, undefined);
                A.store.Bool(ptr + 368 + 32 + 68 + 32 + 27, false);
                A.store.Int32(ptr + 368 + 32 + 68 + 32 + 8, 0);
                A.store.Bool(ptr + 368 + 32 + 68 + 32 + 28, false);
                A.store.Int32(ptr + 368 + 32 + 68 + 32 + 12, 0);
                A.store.Bool(ptr + 368 + 32 + 68 + 32 + 29, false);
                A.store.Int32(ptr + 368 + 32 + 68 + 32 + 16, 0);
                A.store.Bool(ptr + 368 + 32 + 68 + 32 + 30, false);
                A.store.Int32(ptr + 368 + 32 + 68 + 32 + 20, 0);
                A.store.Bool(ptr + 368 + 32 + 68 + 32 + 31, false);
                A.store.Bool(ptr + 368 + 32 + 68 + 32 + 24, false);
                A.store.Bool(ptr + 368 + 32 + 68 + 32 + 32, false);
                A.store.Bool(ptr + 368 + 32 + 68 + 32 + 25, false);
              } else {
                A.store.Bool(ptr + 368 + 32 + 68 + 32 + 33, true);
                A.store.Bool(
                  ptr + 368 + 32 + 68 + 32 + 26,
                  "Active" in x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 68 + 32 + 0,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["Active"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["Active"] as number)
                );
                A.store.Ref(
                  ptr + 368 + 32 + 68 + 32 + 4,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["Effective"]
                );
                A.store.Bool(
                  ptr + 368 + 32 + 68 + 32 + 27,
                  "UserPolicy" in x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 68 + 32 + 8,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["UserPolicy"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["UserPolicy"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 68 + 32 + 28,
                  "DevicePolicy" in x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 68 + 32 + 12,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["DevicePolicy"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["DevicePolicy"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 68 + 32 + 29,
                  "UserSetting" in x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 68 + 32 + 16,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["UserSetting"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["UserSetting"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 68 + 32 + 30,
                  "SharedSetting" in x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 68 + 32 + 20,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["SharedSetting"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["SharedSetting"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 68 + 32 + 31,
                  "UserEditable" in x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 68 + 32 + 24,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["UserEditable"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 68 + 32 + 32,
                  "DeviceEditable" in x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 68 + 32 + 25,
                  x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["DeviceEditable"] ? true : false
                );
              }
            }

            if (typeof x["ProxySettings"]["Manual"]["FTPProxy"] === "undefined") {
              A.store.Bool(ptr + 368 + 32 + 136 + 66, false);

              A.store.Bool(ptr + 368 + 32 + 136 + 0 + 28, false);
              A.store.Ref(ptr + 368 + 32 + 136 + 0 + 0, undefined);
              A.store.Ref(ptr + 368 + 32 + 136 + 0 + 4, undefined);
              A.store.Ref(ptr + 368 + 32 + 136 + 0 + 8, undefined);
              A.store.Ref(ptr + 368 + 32 + 136 + 0 + 12, undefined);
              A.store.Ref(ptr + 368 + 32 + 136 + 0 + 16, undefined);
              A.store.Ref(ptr + 368 + 32 + 136 + 0 + 20, undefined);
              A.store.Bool(ptr + 368 + 32 + 136 + 0 + 26, false);
              A.store.Bool(ptr + 368 + 32 + 136 + 0 + 24, false);
              A.store.Bool(ptr + 368 + 32 + 136 + 0 + 27, false);
              A.store.Bool(ptr + 368 + 32 + 136 + 0 + 25, false);

              A.store.Bool(ptr + 368 + 32 + 136 + 32 + 33, false);
              A.store.Bool(ptr + 368 + 32 + 136 + 32 + 26, false);
              A.store.Int32(ptr + 368 + 32 + 136 + 32 + 0, 0);
              A.store.Ref(ptr + 368 + 32 + 136 + 32 + 4, undefined);
              A.store.Bool(ptr + 368 + 32 + 136 + 32 + 27, false);
              A.store.Int32(ptr + 368 + 32 + 136 + 32 + 8, 0);
              A.store.Bool(ptr + 368 + 32 + 136 + 32 + 28, false);
              A.store.Int32(ptr + 368 + 32 + 136 + 32 + 12, 0);
              A.store.Bool(ptr + 368 + 32 + 136 + 32 + 29, false);
              A.store.Int32(ptr + 368 + 32 + 136 + 32 + 16, 0);
              A.store.Bool(ptr + 368 + 32 + 136 + 32 + 30, false);
              A.store.Int32(ptr + 368 + 32 + 136 + 32 + 20, 0);
              A.store.Bool(ptr + 368 + 32 + 136 + 32 + 31, false);
              A.store.Bool(ptr + 368 + 32 + 136 + 32 + 24, false);
              A.store.Bool(ptr + 368 + 32 + 136 + 32 + 32, false);
              A.store.Bool(ptr + 368 + 32 + 136 + 32 + 25, false);
            } else {
              A.store.Bool(ptr + 368 + 32 + 136 + 66, true);

              if (typeof x["ProxySettings"]["Manual"]["FTPProxy"]["Host"] === "undefined") {
                A.store.Bool(ptr + 368 + 32 + 136 + 0 + 28, false);
                A.store.Ref(ptr + 368 + 32 + 136 + 0 + 0, undefined);
                A.store.Ref(ptr + 368 + 32 + 136 + 0 + 4, undefined);
                A.store.Ref(ptr + 368 + 32 + 136 + 0 + 8, undefined);
                A.store.Ref(ptr + 368 + 32 + 136 + 0 + 12, undefined);
                A.store.Ref(ptr + 368 + 32 + 136 + 0 + 16, undefined);
                A.store.Ref(ptr + 368 + 32 + 136 + 0 + 20, undefined);
                A.store.Bool(ptr + 368 + 32 + 136 + 0 + 26, false);
                A.store.Bool(ptr + 368 + 32 + 136 + 0 + 24, false);
                A.store.Bool(ptr + 368 + 32 + 136 + 0 + 27, false);
                A.store.Bool(ptr + 368 + 32 + 136 + 0 + 25, false);
              } else {
                A.store.Bool(ptr + 368 + 32 + 136 + 0 + 28, true);
                A.store.Ref(ptr + 368 + 32 + 136 + 0 + 0, x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["Active"]);
                A.store.Ref(
                  ptr + 368 + 32 + 136 + 0 + 4,
                  x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["Effective"]
                );
                A.store.Ref(
                  ptr + 368 + 32 + 136 + 0 + 8,
                  x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["UserPolicy"]
                );
                A.store.Ref(
                  ptr + 368 + 32 + 136 + 0 + 12,
                  x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["DevicePolicy"]
                );
                A.store.Ref(
                  ptr + 368 + 32 + 136 + 0 + 16,
                  x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["UserSetting"]
                );
                A.store.Ref(
                  ptr + 368 + 32 + 136 + 0 + 20,
                  x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["SharedSetting"]
                );
                A.store.Bool(
                  ptr + 368 + 32 + 136 + 0 + 26,
                  "UserEditable" in x["ProxySettings"]["Manual"]["FTPProxy"]["Host"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 136 + 0 + 24,
                  x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["UserEditable"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 136 + 0 + 27,
                  "DeviceEditable" in x["ProxySettings"]["Manual"]["FTPProxy"]["Host"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 136 + 0 + 25,
                  x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["DeviceEditable"] ? true : false
                );
              }

              if (typeof x["ProxySettings"]["Manual"]["FTPProxy"]["Port"] === "undefined") {
                A.store.Bool(ptr + 368 + 32 + 136 + 32 + 33, false);
                A.store.Bool(ptr + 368 + 32 + 136 + 32 + 26, false);
                A.store.Int32(ptr + 368 + 32 + 136 + 32 + 0, 0);
                A.store.Ref(ptr + 368 + 32 + 136 + 32 + 4, undefined);
                A.store.Bool(ptr + 368 + 32 + 136 + 32 + 27, false);
                A.store.Int32(ptr + 368 + 32 + 136 + 32 + 8, 0);
                A.store.Bool(ptr + 368 + 32 + 136 + 32 + 28, false);
                A.store.Int32(ptr + 368 + 32 + 136 + 32 + 12, 0);
                A.store.Bool(ptr + 368 + 32 + 136 + 32 + 29, false);
                A.store.Int32(ptr + 368 + 32 + 136 + 32 + 16, 0);
                A.store.Bool(ptr + 368 + 32 + 136 + 32 + 30, false);
                A.store.Int32(ptr + 368 + 32 + 136 + 32 + 20, 0);
                A.store.Bool(ptr + 368 + 32 + 136 + 32 + 31, false);
                A.store.Bool(ptr + 368 + 32 + 136 + 32 + 24, false);
                A.store.Bool(ptr + 368 + 32 + 136 + 32 + 32, false);
                A.store.Bool(ptr + 368 + 32 + 136 + 32 + 25, false);
              } else {
                A.store.Bool(ptr + 368 + 32 + 136 + 32 + 33, true);
                A.store.Bool(
                  ptr + 368 + 32 + 136 + 32 + 26,
                  "Active" in x["ProxySettings"]["Manual"]["FTPProxy"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 136 + 32 + 0,
                  x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["Active"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["Active"] as number)
                );
                A.store.Ref(
                  ptr + 368 + 32 + 136 + 32 + 4,
                  x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["Effective"]
                );
                A.store.Bool(
                  ptr + 368 + 32 + 136 + 32 + 27,
                  "UserPolicy" in x["ProxySettings"]["Manual"]["FTPProxy"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 136 + 32 + 8,
                  x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["UserPolicy"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["UserPolicy"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 136 + 32 + 28,
                  "DevicePolicy" in x["ProxySettings"]["Manual"]["FTPProxy"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 136 + 32 + 12,
                  x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["DevicePolicy"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["DevicePolicy"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 136 + 32 + 29,
                  "UserSetting" in x["ProxySettings"]["Manual"]["FTPProxy"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 136 + 32 + 16,
                  x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["UserSetting"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["UserSetting"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 136 + 32 + 30,
                  "SharedSetting" in x["ProxySettings"]["Manual"]["FTPProxy"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 136 + 32 + 20,
                  x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["SharedSetting"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["SharedSetting"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 136 + 32 + 31,
                  "UserEditable" in x["ProxySettings"]["Manual"]["FTPProxy"]["Port"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 136 + 32 + 24,
                  x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["UserEditable"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 136 + 32 + 32,
                  "DeviceEditable" in x["ProxySettings"]["Manual"]["FTPProxy"]["Port"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 136 + 32 + 25,
                  x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["DeviceEditable"] ? true : false
                );
              }
            }

            if (typeof x["ProxySettings"]["Manual"]["SOCKS"] === "undefined") {
              A.store.Bool(ptr + 368 + 32 + 204 + 66, false);

              A.store.Bool(ptr + 368 + 32 + 204 + 0 + 28, false);
              A.store.Ref(ptr + 368 + 32 + 204 + 0 + 0, undefined);
              A.store.Ref(ptr + 368 + 32 + 204 + 0 + 4, undefined);
              A.store.Ref(ptr + 368 + 32 + 204 + 0 + 8, undefined);
              A.store.Ref(ptr + 368 + 32 + 204 + 0 + 12, undefined);
              A.store.Ref(ptr + 368 + 32 + 204 + 0 + 16, undefined);
              A.store.Ref(ptr + 368 + 32 + 204 + 0 + 20, undefined);
              A.store.Bool(ptr + 368 + 32 + 204 + 0 + 26, false);
              A.store.Bool(ptr + 368 + 32 + 204 + 0 + 24, false);
              A.store.Bool(ptr + 368 + 32 + 204 + 0 + 27, false);
              A.store.Bool(ptr + 368 + 32 + 204 + 0 + 25, false);

              A.store.Bool(ptr + 368 + 32 + 204 + 32 + 33, false);
              A.store.Bool(ptr + 368 + 32 + 204 + 32 + 26, false);
              A.store.Int32(ptr + 368 + 32 + 204 + 32 + 0, 0);
              A.store.Ref(ptr + 368 + 32 + 204 + 32 + 4, undefined);
              A.store.Bool(ptr + 368 + 32 + 204 + 32 + 27, false);
              A.store.Int32(ptr + 368 + 32 + 204 + 32 + 8, 0);
              A.store.Bool(ptr + 368 + 32 + 204 + 32 + 28, false);
              A.store.Int32(ptr + 368 + 32 + 204 + 32 + 12, 0);
              A.store.Bool(ptr + 368 + 32 + 204 + 32 + 29, false);
              A.store.Int32(ptr + 368 + 32 + 204 + 32 + 16, 0);
              A.store.Bool(ptr + 368 + 32 + 204 + 32 + 30, false);
              A.store.Int32(ptr + 368 + 32 + 204 + 32 + 20, 0);
              A.store.Bool(ptr + 368 + 32 + 204 + 32 + 31, false);
              A.store.Bool(ptr + 368 + 32 + 204 + 32 + 24, false);
              A.store.Bool(ptr + 368 + 32 + 204 + 32 + 32, false);
              A.store.Bool(ptr + 368 + 32 + 204 + 32 + 25, false);
            } else {
              A.store.Bool(ptr + 368 + 32 + 204 + 66, true);

              if (typeof x["ProxySettings"]["Manual"]["SOCKS"]["Host"] === "undefined") {
                A.store.Bool(ptr + 368 + 32 + 204 + 0 + 28, false);
                A.store.Ref(ptr + 368 + 32 + 204 + 0 + 0, undefined);
                A.store.Ref(ptr + 368 + 32 + 204 + 0 + 4, undefined);
                A.store.Ref(ptr + 368 + 32 + 204 + 0 + 8, undefined);
                A.store.Ref(ptr + 368 + 32 + 204 + 0 + 12, undefined);
                A.store.Ref(ptr + 368 + 32 + 204 + 0 + 16, undefined);
                A.store.Ref(ptr + 368 + 32 + 204 + 0 + 20, undefined);
                A.store.Bool(ptr + 368 + 32 + 204 + 0 + 26, false);
                A.store.Bool(ptr + 368 + 32 + 204 + 0 + 24, false);
                A.store.Bool(ptr + 368 + 32 + 204 + 0 + 27, false);
                A.store.Bool(ptr + 368 + 32 + 204 + 0 + 25, false);
              } else {
                A.store.Bool(ptr + 368 + 32 + 204 + 0 + 28, true);
                A.store.Ref(ptr + 368 + 32 + 204 + 0 + 0, x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["Active"]);
                A.store.Ref(ptr + 368 + 32 + 204 + 0 + 4, x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["Effective"]);
                A.store.Ref(ptr + 368 + 32 + 204 + 0 + 8, x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["UserPolicy"]);
                A.store.Ref(
                  ptr + 368 + 32 + 204 + 0 + 12,
                  x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["DevicePolicy"]
                );
                A.store.Ref(
                  ptr + 368 + 32 + 204 + 0 + 16,
                  x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["UserSetting"]
                );
                A.store.Ref(
                  ptr + 368 + 32 + 204 + 0 + 20,
                  x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["SharedSetting"]
                );
                A.store.Bool(
                  ptr + 368 + 32 + 204 + 0 + 26,
                  "UserEditable" in x["ProxySettings"]["Manual"]["SOCKS"]["Host"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 204 + 0 + 24,
                  x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["UserEditable"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 204 + 0 + 27,
                  "DeviceEditable" in x["ProxySettings"]["Manual"]["SOCKS"]["Host"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 204 + 0 + 25,
                  x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["DeviceEditable"] ? true : false
                );
              }

              if (typeof x["ProxySettings"]["Manual"]["SOCKS"]["Port"] === "undefined") {
                A.store.Bool(ptr + 368 + 32 + 204 + 32 + 33, false);
                A.store.Bool(ptr + 368 + 32 + 204 + 32 + 26, false);
                A.store.Int32(ptr + 368 + 32 + 204 + 32 + 0, 0);
                A.store.Ref(ptr + 368 + 32 + 204 + 32 + 4, undefined);
                A.store.Bool(ptr + 368 + 32 + 204 + 32 + 27, false);
                A.store.Int32(ptr + 368 + 32 + 204 + 32 + 8, 0);
                A.store.Bool(ptr + 368 + 32 + 204 + 32 + 28, false);
                A.store.Int32(ptr + 368 + 32 + 204 + 32 + 12, 0);
                A.store.Bool(ptr + 368 + 32 + 204 + 32 + 29, false);
                A.store.Int32(ptr + 368 + 32 + 204 + 32 + 16, 0);
                A.store.Bool(ptr + 368 + 32 + 204 + 32 + 30, false);
                A.store.Int32(ptr + 368 + 32 + 204 + 32 + 20, 0);
                A.store.Bool(ptr + 368 + 32 + 204 + 32 + 31, false);
                A.store.Bool(ptr + 368 + 32 + 204 + 32 + 24, false);
                A.store.Bool(ptr + 368 + 32 + 204 + 32 + 32, false);
                A.store.Bool(ptr + 368 + 32 + 204 + 32 + 25, false);
              } else {
                A.store.Bool(ptr + 368 + 32 + 204 + 32 + 33, true);
                A.store.Bool(
                  ptr + 368 + 32 + 204 + 32 + 26,
                  "Active" in x["ProxySettings"]["Manual"]["SOCKS"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 204 + 32 + 0,
                  x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["Active"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["Active"] as number)
                );
                A.store.Ref(ptr + 368 + 32 + 204 + 32 + 4, x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["Effective"]);
                A.store.Bool(
                  ptr + 368 + 32 + 204 + 32 + 27,
                  "UserPolicy" in x["ProxySettings"]["Manual"]["SOCKS"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 204 + 32 + 8,
                  x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["UserPolicy"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["UserPolicy"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 204 + 32 + 28,
                  "DevicePolicy" in x["ProxySettings"]["Manual"]["SOCKS"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 204 + 32 + 12,
                  x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["DevicePolicy"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["DevicePolicy"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 204 + 32 + 29,
                  "UserSetting" in x["ProxySettings"]["Manual"]["SOCKS"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 204 + 32 + 16,
                  x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["UserSetting"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["UserSetting"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 204 + 32 + 30,
                  "SharedSetting" in x["ProxySettings"]["Manual"]["SOCKS"]["Port"] ? true : false
                );
                A.store.Int32(
                  ptr + 368 + 32 + 204 + 32 + 20,
                  x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["SharedSetting"] === undefined
                    ? 0
                    : (x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["SharedSetting"] as number)
                );
                A.store.Bool(
                  ptr + 368 + 32 + 204 + 32 + 31,
                  "UserEditable" in x["ProxySettings"]["Manual"]["SOCKS"]["Port"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 204 + 32 + 24,
                  x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["UserEditable"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 204 + 32 + 32,
                  "DeviceEditable" in x["ProxySettings"]["Manual"]["SOCKS"]["Port"] ? true : false
                );
                A.store.Bool(
                  ptr + 368 + 32 + 204 + 32 + 25,
                  x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["DeviceEditable"] ? true : false
                );
              }
            }
          }

          if (typeof x["ProxySettings"]["ExcludeDomains"] === "undefined") {
            A.store.Bool(ptr + 368 + 304 + 28, false);
            A.store.Ref(ptr + 368 + 304 + 0, undefined);
            A.store.Ref(ptr + 368 + 304 + 4, undefined);
            A.store.Ref(ptr + 368 + 304 + 8, undefined);
            A.store.Ref(ptr + 368 + 304 + 12, undefined);
            A.store.Ref(ptr + 368 + 304 + 16, undefined);
            A.store.Ref(ptr + 368 + 304 + 20, undefined);
            A.store.Bool(ptr + 368 + 304 + 26, false);
            A.store.Bool(ptr + 368 + 304 + 24, false);
            A.store.Bool(ptr + 368 + 304 + 27, false);
            A.store.Bool(ptr + 368 + 304 + 25, false);
          } else {
            A.store.Bool(ptr + 368 + 304 + 28, true);
            A.store.Ref(ptr + 368 + 304 + 0, x["ProxySettings"]["ExcludeDomains"]["Active"]);
            A.store.Ref(ptr + 368 + 304 + 4, x["ProxySettings"]["ExcludeDomains"]["Effective"]);
            A.store.Ref(ptr + 368 + 304 + 8, x["ProxySettings"]["ExcludeDomains"]["UserPolicy"]);
            A.store.Ref(ptr + 368 + 304 + 12, x["ProxySettings"]["ExcludeDomains"]["DevicePolicy"]);
            A.store.Ref(ptr + 368 + 304 + 16, x["ProxySettings"]["ExcludeDomains"]["UserSetting"]);
            A.store.Ref(ptr + 368 + 304 + 20, x["ProxySettings"]["ExcludeDomains"]["SharedSetting"]);
            A.store.Bool(ptr + 368 + 304 + 26, "UserEditable" in x["ProxySettings"]["ExcludeDomains"] ? true : false);
            A.store.Bool(ptr + 368 + 304 + 24, x["ProxySettings"]["ExcludeDomains"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 368 + 304 + 27, "DeviceEditable" in x["ProxySettings"]["ExcludeDomains"] ? true : false);
            A.store.Bool(ptr + 368 + 304 + 25, x["ProxySettings"]["ExcludeDomains"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["ProxySettings"]["PAC"] === "undefined") {
            A.store.Bool(ptr + 368 + 336 + 28, false);
            A.store.Ref(ptr + 368 + 336 + 0, undefined);
            A.store.Ref(ptr + 368 + 336 + 4, undefined);
            A.store.Ref(ptr + 368 + 336 + 8, undefined);
            A.store.Ref(ptr + 368 + 336 + 12, undefined);
            A.store.Ref(ptr + 368 + 336 + 16, undefined);
            A.store.Ref(ptr + 368 + 336 + 20, undefined);
            A.store.Bool(ptr + 368 + 336 + 26, false);
            A.store.Bool(ptr + 368 + 336 + 24, false);
            A.store.Bool(ptr + 368 + 336 + 27, false);
            A.store.Bool(ptr + 368 + 336 + 25, false);
          } else {
            A.store.Bool(ptr + 368 + 336 + 28, true);
            A.store.Ref(ptr + 368 + 336 + 0, x["ProxySettings"]["PAC"]["Active"]);
            A.store.Ref(ptr + 368 + 336 + 4, x["ProxySettings"]["PAC"]["Effective"]);
            A.store.Ref(ptr + 368 + 336 + 8, x["ProxySettings"]["PAC"]["UserPolicy"]);
            A.store.Ref(ptr + 368 + 336 + 12, x["ProxySettings"]["PAC"]["DevicePolicy"]);
            A.store.Ref(ptr + 368 + 336 + 16, x["ProxySettings"]["PAC"]["UserSetting"]);
            A.store.Ref(ptr + 368 + 336 + 20, x["ProxySettings"]["PAC"]["SharedSetting"]);
            A.store.Bool(ptr + 368 + 336 + 26, "UserEditable" in x["ProxySettings"]["PAC"] ? true : false);
            A.store.Bool(ptr + 368 + 336 + 24, x["ProxySettings"]["PAC"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 368 + 336 + 27, "DeviceEditable" in x["ProxySettings"]["PAC"] ? true : false);
            A.store.Bool(ptr + 368 + 336 + 25, x["ProxySettings"]["PAC"]["DeviceEditable"] ? true : false);
          }
        }
        A.store.Bool(ptr + 1292, "RestrictedConnectivity" in x ? true : false);
        A.store.Bool(ptr + 734, x["RestrictedConnectivity"] ? true : false);

        if (typeof x["StaticIPConfig"] === "undefined") {
          A.store.Bool(ptr + 736 + 193, false);

          A.store.Bool(ptr + 736 + 0 + 28, false);
          A.store.Ref(ptr + 736 + 0 + 0, undefined);
          A.store.Ref(ptr + 736 + 0 + 4, undefined);
          A.store.Ref(ptr + 736 + 0 + 8, undefined);
          A.store.Ref(ptr + 736 + 0 + 12, undefined);
          A.store.Ref(ptr + 736 + 0 + 16, undefined);
          A.store.Ref(ptr + 736 + 0 + 20, undefined);
          A.store.Bool(ptr + 736 + 0 + 26, false);
          A.store.Bool(ptr + 736 + 0 + 24, false);
          A.store.Bool(ptr + 736 + 0 + 27, false);
          A.store.Bool(ptr + 736 + 0 + 25, false);

          A.store.Bool(ptr + 736 + 32 + 28, false);
          A.store.Ref(ptr + 736 + 32 + 0, undefined);
          A.store.Ref(ptr + 736 + 32 + 4, undefined);
          A.store.Ref(ptr + 736 + 32 + 8, undefined);
          A.store.Ref(ptr + 736 + 32 + 12, undefined);
          A.store.Ref(ptr + 736 + 32 + 16, undefined);
          A.store.Ref(ptr + 736 + 32 + 20, undefined);
          A.store.Bool(ptr + 736 + 32 + 26, false);
          A.store.Bool(ptr + 736 + 32 + 24, false);
          A.store.Bool(ptr + 736 + 32 + 27, false);
          A.store.Bool(ptr + 736 + 32 + 25, false);

          A.store.Bool(ptr + 736 + 64 + 28, false);
          A.store.Ref(ptr + 736 + 64 + 0, undefined);
          A.store.Ref(ptr + 736 + 64 + 4, undefined);
          A.store.Ref(ptr + 736 + 64 + 8, undefined);
          A.store.Ref(ptr + 736 + 64 + 12, undefined);
          A.store.Ref(ptr + 736 + 64 + 16, undefined);
          A.store.Ref(ptr + 736 + 64 + 20, undefined);
          A.store.Bool(ptr + 736 + 64 + 26, false);
          A.store.Bool(ptr + 736 + 64 + 24, false);
          A.store.Bool(ptr + 736 + 64 + 27, false);
          A.store.Bool(ptr + 736 + 64 + 25, false);

          A.store.Bool(ptr + 736 + 96 + 33, false);
          A.store.Bool(ptr + 736 + 96 + 26, false);
          A.store.Int32(ptr + 736 + 96 + 0, 0);
          A.store.Ref(ptr + 736 + 96 + 4, undefined);
          A.store.Bool(ptr + 736 + 96 + 27, false);
          A.store.Int32(ptr + 736 + 96 + 8, 0);
          A.store.Bool(ptr + 736 + 96 + 28, false);
          A.store.Int32(ptr + 736 + 96 + 12, 0);
          A.store.Bool(ptr + 736 + 96 + 29, false);
          A.store.Int32(ptr + 736 + 96 + 16, 0);
          A.store.Bool(ptr + 736 + 96 + 30, false);
          A.store.Int32(ptr + 736 + 96 + 20, 0);
          A.store.Bool(ptr + 736 + 96 + 31, false);
          A.store.Bool(ptr + 736 + 96 + 24, false);
          A.store.Bool(ptr + 736 + 96 + 32, false);
          A.store.Bool(ptr + 736 + 96 + 25, false);

          A.store.Bool(ptr + 736 + 132 + 28, false);
          A.store.Ref(ptr + 736 + 132 + 0, undefined);
          A.store.Ref(ptr + 736 + 132 + 4, undefined);
          A.store.Ref(ptr + 736 + 132 + 8, undefined);
          A.store.Ref(ptr + 736 + 132 + 12, undefined);
          A.store.Ref(ptr + 736 + 132 + 16, undefined);
          A.store.Ref(ptr + 736 + 132 + 20, undefined);
          A.store.Bool(ptr + 736 + 132 + 26, false);
          A.store.Bool(ptr + 736 + 132 + 24, false);
          A.store.Bool(ptr + 736 + 132 + 27, false);
          A.store.Bool(ptr + 736 + 132 + 25, false);

          A.store.Bool(ptr + 736 + 164 + 28, false);
          A.store.Ref(ptr + 736 + 164 + 0, undefined);
          A.store.Ref(ptr + 736 + 164 + 4, undefined);
          A.store.Ref(ptr + 736 + 164 + 8, undefined);
          A.store.Ref(ptr + 736 + 164 + 12, undefined);
          A.store.Ref(ptr + 736 + 164 + 16, undefined);
          A.store.Ref(ptr + 736 + 164 + 20, undefined);
          A.store.Bool(ptr + 736 + 164 + 26, false);
          A.store.Bool(ptr + 736 + 164 + 24, false);
          A.store.Bool(ptr + 736 + 164 + 27, false);
          A.store.Bool(ptr + 736 + 164 + 25, false);
        } else {
          A.store.Bool(ptr + 736 + 193, true);

          if (typeof x["StaticIPConfig"]["Gateway"] === "undefined") {
            A.store.Bool(ptr + 736 + 0 + 28, false);
            A.store.Ref(ptr + 736 + 0 + 0, undefined);
            A.store.Ref(ptr + 736 + 0 + 4, undefined);
            A.store.Ref(ptr + 736 + 0 + 8, undefined);
            A.store.Ref(ptr + 736 + 0 + 12, undefined);
            A.store.Ref(ptr + 736 + 0 + 16, undefined);
            A.store.Ref(ptr + 736 + 0 + 20, undefined);
            A.store.Bool(ptr + 736 + 0 + 26, false);
            A.store.Bool(ptr + 736 + 0 + 24, false);
            A.store.Bool(ptr + 736 + 0 + 27, false);
            A.store.Bool(ptr + 736 + 0 + 25, false);
          } else {
            A.store.Bool(ptr + 736 + 0 + 28, true);
            A.store.Ref(ptr + 736 + 0 + 0, x["StaticIPConfig"]["Gateway"]["Active"]);
            A.store.Ref(ptr + 736 + 0 + 4, x["StaticIPConfig"]["Gateway"]["Effective"]);
            A.store.Ref(ptr + 736 + 0 + 8, x["StaticIPConfig"]["Gateway"]["UserPolicy"]);
            A.store.Ref(ptr + 736 + 0 + 12, x["StaticIPConfig"]["Gateway"]["DevicePolicy"]);
            A.store.Ref(ptr + 736 + 0 + 16, x["StaticIPConfig"]["Gateway"]["UserSetting"]);
            A.store.Ref(ptr + 736 + 0 + 20, x["StaticIPConfig"]["Gateway"]["SharedSetting"]);
            A.store.Bool(ptr + 736 + 0 + 26, "UserEditable" in x["StaticIPConfig"]["Gateway"] ? true : false);
            A.store.Bool(ptr + 736 + 0 + 24, x["StaticIPConfig"]["Gateway"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 736 + 0 + 27, "DeviceEditable" in x["StaticIPConfig"]["Gateway"] ? true : false);
            A.store.Bool(ptr + 736 + 0 + 25, x["StaticIPConfig"]["Gateway"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["StaticIPConfig"]["IPAddress"] === "undefined") {
            A.store.Bool(ptr + 736 + 32 + 28, false);
            A.store.Ref(ptr + 736 + 32 + 0, undefined);
            A.store.Ref(ptr + 736 + 32 + 4, undefined);
            A.store.Ref(ptr + 736 + 32 + 8, undefined);
            A.store.Ref(ptr + 736 + 32 + 12, undefined);
            A.store.Ref(ptr + 736 + 32 + 16, undefined);
            A.store.Ref(ptr + 736 + 32 + 20, undefined);
            A.store.Bool(ptr + 736 + 32 + 26, false);
            A.store.Bool(ptr + 736 + 32 + 24, false);
            A.store.Bool(ptr + 736 + 32 + 27, false);
            A.store.Bool(ptr + 736 + 32 + 25, false);
          } else {
            A.store.Bool(ptr + 736 + 32 + 28, true);
            A.store.Ref(ptr + 736 + 32 + 0, x["StaticIPConfig"]["IPAddress"]["Active"]);
            A.store.Ref(ptr + 736 + 32 + 4, x["StaticIPConfig"]["IPAddress"]["Effective"]);
            A.store.Ref(ptr + 736 + 32 + 8, x["StaticIPConfig"]["IPAddress"]["UserPolicy"]);
            A.store.Ref(ptr + 736 + 32 + 12, x["StaticIPConfig"]["IPAddress"]["DevicePolicy"]);
            A.store.Ref(ptr + 736 + 32 + 16, x["StaticIPConfig"]["IPAddress"]["UserSetting"]);
            A.store.Ref(ptr + 736 + 32 + 20, x["StaticIPConfig"]["IPAddress"]["SharedSetting"]);
            A.store.Bool(ptr + 736 + 32 + 26, "UserEditable" in x["StaticIPConfig"]["IPAddress"] ? true : false);
            A.store.Bool(ptr + 736 + 32 + 24, x["StaticIPConfig"]["IPAddress"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 736 + 32 + 27, "DeviceEditable" in x["StaticIPConfig"]["IPAddress"] ? true : false);
            A.store.Bool(ptr + 736 + 32 + 25, x["StaticIPConfig"]["IPAddress"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["StaticIPConfig"]["NameServers"] === "undefined") {
            A.store.Bool(ptr + 736 + 64 + 28, false);
            A.store.Ref(ptr + 736 + 64 + 0, undefined);
            A.store.Ref(ptr + 736 + 64 + 4, undefined);
            A.store.Ref(ptr + 736 + 64 + 8, undefined);
            A.store.Ref(ptr + 736 + 64 + 12, undefined);
            A.store.Ref(ptr + 736 + 64 + 16, undefined);
            A.store.Ref(ptr + 736 + 64 + 20, undefined);
            A.store.Bool(ptr + 736 + 64 + 26, false);
            A.store.Bool(ptr + 736 + 64 + 24, false);
            A.store.Bool(ptr + 736 + 64 + 27, false);
            A.store.Bool(ptr + 736 + 64 + 25, false);
          } else {
            A.store.Bool(ptr + 736 + 64 + 28, true);
            A.store.Ref(ptr + 736 + 64 + 0, x["StaticIPConfig"]["NameServers"]["Active"]);
            A.store.Ref(ptr + 736 + 64 + 4, x["StaticIPConfig"]["NameServers"]["Effective"]);
            A.store.Ref(ptr + 736 + 64 + 8, x["StaticIPConfig"]["NameServers"]["UserPolicy"]);
            A.store.Ref(ptr + 736 + 64 + 12, x["StaticIPConfig"]["NameServers"]["DevicePolicy"]);
            A.store.Ref(ptr + 736 + 64 + 16, x["StaticIPConfig"]["NameServers"]["UserSetting"]);
            A.store.Ref(ptr + 736 + 64 + 20, x["StaticIPConfig"]["NameServers"]["SharedSetting"]);
            A.store.Bool(ptr + 736 + 64 + 26, "UserEditable" in x["StaticIPConfig"]["NameServers"] ? true : false);
            A.store.Bool(ptr + 736 + 64 + 24, x["StaticIPConfig"]["NameServers"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 736 + 64 + 27, "DeviceEditable" in x["StaticIPConfig"]["NameServers"] ? true : false);
            A.store.Bool(ptr + 736 + 64 + 25, x["StaticIPConfig"]["NameServers"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["StaticIPConfig"]["RoutingPrefix"] === "undefined") {
            A.store.Bool(ptr + 736 + 96 + 33, false);
            A.store.Bool(ptr + 736 + 96 + 26, false);
            A.store.Int32(ptr + 736 + 96 + 0, 0);
            A.store.Ref(ptr + 736 + 96 + 4, undefined);
            A.store.Bool(ptr + 736 + 96 + 27, false);
            A.store.Int32(ptr + 736 + 96 + 8, 0);
            A.store.Bool(ptr + 736 + 96 + 28, false);
            A.store.Int32(ptr + 736 + 96 + 12, 0);
            A.store.Bool(ptr + 736 + 96 + 29, false);
            A.store.Int32(ptr + 736 + 96 + 16, 0);
            A.store.Bool(ptr + 736 + 96 + 30, false);
            A.store.Int32(ptr + 736 + 96 + 20, 0);
            A.store.Bool(ptr + 736 + 96 + 31, false);
            A.store.Bool(ptr + 736 + 96 + 24, false);
            A.store.Bool(ptr + 736 + 96 + 32, false);
            A.store.Bool(ptr + 736 + 96 + 25, false);
          } else {
            A.store.Bool(ptr + 736 + 96 + 33, true);
            A.store.Bool(ptr + 736 + 96 + 26, "Active" in x["StaticIPConfig"]["RoutingPrefix"] ? true : false);
            A.store.Int32(
              ptr + 736 + 96 + 0,
              x["StaticIPConfig"]["RoutingPrefix"]["Active"] === undefined
                ? 0
                : (x["StaticIPConfig"]["RoutingPrefix"]["Active"] as number)
            );
            A.store.Ref(ptr + 736 + 96 + 4, x["StaticIPConfig"]["RoutingPrefix"]["Effective"]);
            A.store.Bool(ptr + 736 + 96 + 27, "UserPolicy" in x["StaticIPConfig"]["RoutingPrefix"] ? true : false);
            A.store.Int32(
              ptr + 736 + 96 + 8,
              x["StaticIPConfig"]["RoutingPrefix"]["UserPolicy"] === undefined
                ? 0
                : (x["StaticIPConfig"]["RoutingPrefix"]["UserPolicy"] as number)
            );
            A.store.Bool(ptr + 736 + 96 + 28, "DevicePolicy" in x["StaticIPConfig"]["RoutingPrefix"] ? true : false);
            A.store.Int32(
              ptr + 736 + 96 + 12,
              x["StaticIPConfig"]["RoutingPrefix"]["DevicePolicy"] === undefined
                ? 0
                : (x["StaticIPConfig"]["RoutingPrefix"]["DevicePolicy"] as number)
            );
            A.store.Bool(ptr + 736 + 96 + 29, "UserSetting" in x["StaticIPConfig"]["RoutingPrefix"] ? true : false);
            A.store.Int32(
              ptr + 736 + 96 + 16,
              x["StaticIPConfig"]["RoutingPrefix"]["UserSetting"] === undefined
                ? 0
                : (x["StaticIPConfig"]["RoutingPrefix"]["UserSetting"] as number)
            );
            A.store.Bool(ptr + 736 + 96 + 30, "SharedSetting" in x["StaticIPConfig"]["RoutingPrefix"] ? true : false);
            A.store.Int32(
              ptr + 736 + 96 + 20,
              x["StaticIPConfig"]["RoutingPrefix"]["SharedSetting"] === undefined
                ? 0
                : (x["StaticIPConfig"]["RoutingPrefix"]["SharedSetting"] as number)
            );
            A.store.Bool(ptr + 736 + 96 + 31, "UserEditable" in x["StaticIPConfig"]["RoutingPrefix"] ? true : false);
            A.store.Bool(ptr + 736 + 96 + 24, x["StaticIPConfig"]["RoutingPrefix"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 736 + 96 + 32, "DeviceEditable" in x["StaticIPConfig"]["RoutingPrefix"] ? true : false);
            A.store.Bool(ptr + 736 + 96 + 25, x["StaticIPConfig"]["RoutingPrefix"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["StaticIPConfig"]["Type"] === "undefined") {
            A.store.Bool(ptr + 736 + 132 + 28, false);
            A.store.Ref(ptr + 736 + 132 + 0, undefined);
            A.store.Ref(ptr + 736 + 132 + 4, undefined);
            A.store.Ref(ptr + 736 + 132 + 8, undefined);
            A.store.Ref(ptr + 736 + 132 + 12, undefined);
            A.store.Ref(ptr + 736 + 132 + 16, undefined);
            A.store.Ref(ptr + 736 + 132 + 20, undefined);
            A.store.Bool(ptr + 736 + 132 + 26, false);
            A.store.Bool(ptr + 736 + 132 + 24, false);
            A.store.Bool(ptr + 736 + 132 + 27, false);
            A.store.Bool(ptr + 736 + 132 + 25, false);
          } else {
            A.store.Bool(ptr + 736 + 132 + 28, true);
            A.store.Ref(ptr + 736 + 132 + 0, x["StaticIPConfig"]["Type"]["Active"]);
            A.store.Ref(ptr + 736 + 132 + 4, x["StaticIPConfig"]["Type"]["Effective"]);
            A.store.Ref(ptr + 736 + 132 + 8, x["StaticIPConfig"]["Type"]["UserPolicy"]);
            A.store.Ref(ptr + 736 + 132 + 12, x["StaticIPConfig"]["Type"]["DevicePolicy"]);
            A.store.Ref(ptr + 736 + 132 + 16, x["StaticIPConfig"]["Type"]["UserSetting"]);
            A.store.Ref(ptr + 736 + 132 + 20, x["StaticIPConfig"]["Type"]["SharedSetting"]);
            A.store.Bool(ptr + 736 + 132 + 26, "UserEditable" in x["StaticIPConfig"]["Type"] ? true : false);
            A.store.Bool(ptr + 736 + 132 + 24, x["StaticIPConfig"]["Type"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 736 + 132 + 27, "DeviceEditable" in x["StaticIPConfig"]["Type"] ? true : false);
            A.store.Bool(ptr + 736 + 132 + 25, x["StaticIPConfig"]["Type"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"] === "undefined") {
            A.store.Bool(ptr + 736 + 164 + 28, false);
            A.store.Ref(ptr + 736 + 164 + 0, undefined);
            A.store.Ref(ptr + 736 + 164 + 4, undefined);
            A.store.Ref(ptr + 736 + 164 + 8, undefined);
            A.store.Ref(ptr + 736 + 164 + 12, undefined);
            A.store.Ref(ptr + 736 + 164 + 16, undefined);
            A.store.Ref(ptr + 736 + 164 + 20, undefined);
            A.store.Bool(ptr + 736 + 164 + 26, false);
            A.store.Bool(ptr + 736 + 164 + 24, false);
            A.store.Bool(ptr + 736 + 164 + 27, false);
            A.store.Bool(ptr + 736 + 164 + 25, false);
          } else {
            A.store.Bool(ptr + 736 + 164 + 28, true);
            A.store.Ref(ptr + 736 + 164 + 0, x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["Active"]);
            A.store.Ref(ptr + 736 + 164 + 4, x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["Effective"]);
            A.store.Ref(ptr + 736 + 164 + 8, x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["UserPolicy"]);
            A.store.Ref(ptr + 736 + 164 + 12, x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["DevicePolicy"]);
            A.store.Ref(ptr + 736 + 164 + 16, x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["UserSetting"]);
            A.store.Ref(ptr + 736 + 164 + 20, x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["SharedSetting"]);
            A.store.Bool(
              ptr + 736 + 164 + 26,
              "UserEditable" in x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"] ? true : false
            );
            A.store.Bool(
              ptr + 736 + 164 + 24,
              x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["UserEditable"] ? true : false
            );
            A.store.Bool(
              ptr + 736 + 164 + 27,
              "DeviceEditable" in x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"] ? true : false
            );
            A.store.Bool(
              ptr + 736 + 164 + 25,
              x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["DeviceEditable"] ? true : false
            );
          }
        }

        if (typeof x["SavedIPConfig"] === "undefined") {
          A.store.Bool(ptr + 932 + 37, false);
          A.store.Ref(ptr + 932 + 0, undefined);
          A.store.Ref(ptr + 932 + 4, undefined);
          A.store.Ref(ptr + 932 + 8, undefined);
          A.store.Ref(ptr + 932 + 12, undefined);
          A.store.Ref(ptr + 932 + 16, undefined);
          A.store.Ref(ptr + 932 + 20, undefined);
          A.store.Bool(ptr + 932 + 36, false);
          A.store.Int32(ptr + 932 + 24, 0);
          A.store.Ref(ptr + 932 + 28, undefined);
          A.store.Ref(ptr + 932 + 32, undefined);
        } else {
          A.store.Bool(ptr + 932 + 37, true);
          A.store.Ref(ptr + 932 + 0, x["SavedIPConfig"]["Gateway"]);
          A.store.Ref(ptr + 932 + 4, x["SavedIPConfig"]["IPAddress"]);
          A.store.Ref(ptr + 932 + 8, x["SavedIPConfig"]["ExcludedRoutes"]);
          A.store.Ref(ptr + 932 + 12, x["SavedIPConfig"]["IncludedRoutes"]);
          A.store.Ref(ptr + 932 + 16, x["SavedIPConfig"]["NameServers"]);
          A.store.Ref(ptr + 932 + 20, x["SavedIPConfig"]["SearchDomains"]);
          A.store.Bool(ptr + 932 + 36, "RoutingPrefix" in x["SavedIPConfig"] ? true : false);
          A.store.Int32(
            ptr + 932 + 24,
            x["SavedIPConfig"]["RoutingPrefix"] === undefined ? 0 : (x["SavedIPConfig"]["RoutingPrefix"] as number)
          );
          A.store.Ref(ptr + 932 + 28, x["SavedIPConfig"]["Type"]);
          A.store.Ref(ptr + 932 + 32, x["SavedIPConfig"]["WebProxyAutoDiscoveryUrl"]);
        }
        A.store.Ref(ptr + 972, x["Source"]);
        A.store.Enum(
          ptr + 976,
          ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"].indexOf(x["Type"] as string)
        );

        if (typeof x["VPN"] === "undefined") {
          A.store.Bool(ptr + 980 + 85, false);

          A.store.Bool(ptr + 980 + 0 + 21, false);
          A.store.Bool(ptr + 980 + 0 + 14, false);
          A.store.Bool(ptr + 980 + 0 + 0, false);
          A.store.Ref(ptr + 980 + 0 + 4, undefined);
          A.store.Bool(ptr + 980 + 0 + 15, false);
          A.store.Bool(ptr + 980 + 0 + 8, false);
          A.store.Bool(ptr + 980 + 0 + 16, false);
          A.store.Bool(ptr + 980 + 0 + 9, false);
          A.store.Bool(ptr + 980 + 0 + 17, false);
          A.store.Bool(ptr + 980 + 0 + 10, false);
          A.store.Bool(ptr + 980 + 0 + 18, false);
          A.store.Bool(ptr + 980 + 0 + 11, false);
          A.store.Bool(ptr + 980 + 0 + 19, false);
          A.store.Bool(ptr + 980 + 0 + 12, false);
          A.store.Bool(ptr + 980 + 0 + 20, false);
          A.store.Bool(ptr + 980 + 0 + 13, false);

          A.store.Bool(ptr + 980 + 24 + 28, false);
          A.store.Ref(ptr + 980 + 24 + 0, undefined);
          A.store.Ref(ptr + 980 + 24 + 4, undefined);
          A.store.Ref(ptr + 980 + 24 + 8, undefined);
          A.store.Ref(ptr + 980 + 24 + 12, undefined);
          A.store.Ref(ptr + 980 + 24 + 16, undefined);
          A.store.Ref(ptr + 980 + 24 + 20, undefined);
          A.store.Bool(ptr + 980 + 24 + 26, false);
          A.store.Bool(ptr + 980 + 24 + 24, false);
          A.store.Bool(ptr + 980 + 24 + 27, false);
          A.store.Bool(ptr + 980 + 24 + 25, false);

          A.store.Bool(ptr + 980 + 56 + 28, false);
          A.store.Ref(ptr + 980 + 56 + 0, undefined);
          A.store.Ref(ptr + 980 + 56 + 4, undefined);
          A.store.Ref(ptr + 980 + 56 + 8, undefined);
          A.store.Ref(ptr + 980 + 56 + 12, undefined);
          A.store.Ref(ptr + 980 + 56 + 16, undefined);
          A.store.Ref(ptr + 980 + 56 + 20, undefined);
          A.store.Bool(ptr + 980 + 56 + 26, false);
          A.store.Bool(ptr + 980 + 56 + 24, false);
          A.store.Bool(ptr + 980 + 56 + 27, false);
          A.store.Bool(ptr + 980 + 56 + 25, false);
        } else {
          A.store.Bool(ptr + 980 + 85, true);

          if (typeof x["VPN"]["AutoConnect"] === "undefined") {
            A.store.Bool(ptr + 980 + 0 + 21, false);
            A.store.Bool(ptr + 980 + 0 + 14, false);
            A.store.Bool(ptr + 980 + 0 + 0, false);
            A.store.Ref(ptr + 980 + 0 + 4, undefined);
            A.store.Bool(ptr + 980 + 0 + 15, false);
            A.store.Bool(ptr + 980 + 0 + 8, false);
            A.store.Bool(ptr + 980 + 0 + 16, false);
            A.store.Bool(ptr + 980 + 0 + 9, false);
            A.store.Bool(ptr + 980 + 0 + 17, false);
            A.store.Bool(ptr + 980 + 0 + 10, false);
            A.store.Bool(ptr + 980 + 0 + 18, false);
            A.store.Bool(ptr + 980 + 0 + 11, false);
            A.store.Bool(ptr + 980 + 0 + 19, false);
            A.store.Bool(ptr + 980 + 0 + 12, false);
            A.store.Bool(ptr + 980 + 0 + 20, false);
            A.store.Bool(ptr + 980 + 0 + 13, false);
          } else {
            A.store.Bool(ptr + 980 + 0 + 21, true);
            A.store.Bool(ptr + 980 + 0 + 14, "Active" in x["VPN"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 980 + 0 + 0, x["VPN"]["AutoConnect"]["Active"] ? true : false);
            A.store.Ref(ptr + 980 + 0 + 4, x["VPN"]["AutoConnect"]["Effective"]);
            A.store.Bool(ptr + 980 + 0 + 15, "UserPolicy" in x["VPN"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 980 + 0 + 8, x["VPN"]["AutoConnect"]["UserPolicy"] ? true : false);
            A.store.Bool(ptr + 980 + 0 + 16, "DevicePolicy" in x["VPN"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 980 + 0 + 9, x["VPN"]["AutoConnect"]["DevicePolicy"] ? true : false);
            A.store.Bool(ptr + 980 + 0 + 17, "UserSetting" in x["VPN"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 980 + 0 + 10, x["VPN"]["AutoConnect"]["UserSetting"] ? true : false);
            A.store.Bool(ptr + 980 + 0 + 18, "SharedSetting" in x["VPN"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 980 + 0 + 11, x["VPN"]["AutoConnect"]["SharedSetting"] ? true : false);
            A.store.Bool(ptr + 980 + 0 + 19, "UserEditable" in x["VPN"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 980 + 0 + 12, x["VPN"]["AutoConnect"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 980 + 0 + 20, "DeviceEditable" in x["VPN"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 980 + 0 + 13, x["VPN"]["AutoConnect"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["VPN"]["Host"] === "undefined") {
            A.store.Bool(ptr + 980 + 24 + 28, false);
            A.store.Ref(ptr + 980 + 24 + 0, undefined);
            A.store.Ref(ptr + 980 + 24 + 4, undefined);
            A.store.Ref(ptr + 980 + 24 + 8, undefined);
            A.store.Ref(ptr + 980 + 24 + 12, undefined);
            A.store.Ref(ptr + 980 + 24 + 16, undefined);
            A.store.Ref(ptr + 980 + 24 + 20, undefined);
            A.store.Bool(ptr + 980 + 24 + 26, false);
            A.store.Bool(ptr + 980 + 24 + 24, false);
            A.store.Bool(ptr + 980 + 24 + 27, false);
            A.store.Bool(ptr + 980 + 24 + 25, false);
          } else {
            A.store.Bool(ptr + 980 + 24 + 28, true);
            A.store.Ref(ptr + 980 + 24 + 0, x["VPN"]["Host"]["Active"]);
            A.store.Ref(ptr + 980 + 24 + 4, x["VPN"]["Host"]["Effective"]);
            A.store.Ref(ptr + 980 + 24 + 8, x["VPN"]["Host"]["UserPolicy"]);
            A.store.Ref(ptr + 980 + 24 + 12, x["VPN"]["Host"]["DevicePolicy"]);
            A.store.Ref(ptr + 980 + 24 + 16, x["VPN"]["Host"]["UserSetting"]);
            A.store.Ref(ptr + 980 + 24 + 20, x["VPN"]["Host"]["SharedSetting"]);
            A.store.Bool(ptr + 980 + 24 + 26, "UserEditable" in x["VPN"]["Host"] ? true : false);
            A.store.Bool(ptr + 980 + 24 + 24, x["VPN"]["Host"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 980 + 24 + 27, "DeviceEditable" in x["VPN"]["Host"] ? true : false);
            A.store.Bool(ptr + 980 + 24 + 25, x["VPN"]["Host"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["VPN"]["Type"] === "undefined") {
            A.store.Bool(ptr + 980 + 56 + 28, false);
            A.store.Ref(ptr + 980 + 56 + 0, undefined);
            A.store.Ref(ptr + 980 + 56 + 4, undefined);
            A.store.Ref(ptr + 980 + 56 + 8, undefined);
            A.store.Ref(ptr + 980 + 56 + 12, undefined);
            A.store.Ref(ptr + 980 + 56 + 16, undefined);
            A.store.Ref(ptr + 980 + 56 + 20, undefined);
            A.store.Bool(ptr + 980 + 56 + 26, false);
            A.store.Bool(ptr + 980 + 56 + 24, false);
            A.store.Bool(ptr + 980 + 56 + 27, false);
            A.store.Bool(ptr + 980 + 56 + 25, false);
          } else {
            A.store.Bool(ptr + 980 + 56 + 28, true);
            A.store.Ref(ptr + 980 + 56 + 0, x["VPN"]["Type"]["Active"]);
            A.store.Ref(ptr + 980 + 56 + 4, x["VPN"]["Type"]["Effective"]);
            A.store.Ref(ptr + 980 + 56 + 8, x["VPN"]["Type"]["UserPolicy"]);
            A.store.Ref(ptr + 980 + 56 + 12, x["VPN"]["Type"]["DevicePolicy"]);
            A.store.Ref(ptr + 980 + 56 + 16, x["VPN"]["Type"]["UserSetting"]);
            A.store.Ref(ptr + 980 + 56 + 20, x["VPN"]["Type"]["SharedSetting"]);
            A.store.Bool(ptr + 980 + 56 + 26, "UserEditable" in x["VPN"]["Type"] ? true : false);
            A.store.Bool(ptr + 980 + 56 + 24, x["VPN"]["Type"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 980 + 56 + 27, "DeviceEditable" in x["VPN"]["Type"] ? true : false);
            A.store.Bool(ptr + 980 + 56 + 25, x["VPN"]["Type"]["DeviceEditable"] ? true : false);
          }
        }

        if (typeof x["WiFi"] === "undefined") {
          A.store.Bool(ptr + 1068 + 222, false);

          A.store.Bool(ptr + 1068 + 0 + 21, false);
          A.store.Bool(ptr + 1068 + 0 + 14, false);
          A.store.Bool(ptr + 1068 + 0 + 0, false);
          A.store.Ref(ptr + 1068 + 0 + 4, undefined);
          A.store.Bool(ptr + 1068 + 0 + 15, false);
          A.store.Bool(ptr + 1068 + 0 + 8, false);
          A.store.Bool(ptr + 1068 + 0 + 16, false);
          A.store.Bool(ptr + 1068 + 0 + 9, false);
          A.store.Bool(ptr + 1068 + 0 + 17, false);
          A.store.Bool(ptr + 1068 + 0 + 10, false);
          A.store.Bool(ptr + 1068 + 0 + 18, false);
          A.store.Bool(ptr + 1068 + 0 + 11, false);
          A.store.Bool(ptr + 1068 + 0 + 19, false);
          A.store.Bool(ptr + 1068 + 0 + 12, false);
          A.store.Bool(ptr + 1068 + 0 + 20, false);
          A.store.Bool(ptr + 1068 + 0 + 13, false);

          A.store.Bool(ptr + 1068 + 24 + 21, false);
          A.store.Bool(ptr + 1068 + 24 + 14, false);
          A.store.Bool(ptr + 1068 + 24 + 0, false);
          A.store.Ref(ptr + 1068 + 24 + 4, undefined);
          A.store.Bool(ptr + 1068 + 24 + 15, false);
          A.store.Bool(ptr + 1068 + 24 + 8, false);
          A.store.Bool(ptr + 1068 + 24 + 16, false);
          A.store.Bool(ptr + 1068 + 24 + 9, false);
          A.store.Bool(ptr + 1068 + 24 + 17, false);
          A.store.Bool(ptr + 1068 + 24 + 10, false);
          A.store.Bool(ptr + 1068 + 24 + 18, false);
          A.store.Bool(ptr + 1068 + 24 + 11, false);
          A.store.Bool(ptr + 1068 + 24 + 19, false);
          A.store.Bool(ptr + 1068 + 24 + 12, false);
          A.store.Bool(ptr + 1068 + 24 + 20, false);
          A.store.Bool(ptr + 1068 + 24 + 13, false);
          A.store.Ref(ptr + 1068 + 48, undefined);
          A.store.Bool(ptr + 1068 + 220, false);
          A.store.Int32(ptr + 1068 + 52, 0);
          A.store.Ref(ptr + 1068 + 56, undefined);

          A.store.Bool(ptr + 1068 + 60 + 28, false);
          A.store.Ref(ptr + 1068 + 60 + 0, undefined);
          A.store.Ref(ptr + 1068 + 60 + 4, undefined);
          A.store.Ref(ptr + 1068 + 60 + 8, undefined);
          A.store.Ref(ptr + 1068 + 60 + 12, undefined);
          A.store.Ref(ptr + 1068 + 60 + 16, undefined);
          A.store.Ref(ptr + 1068 + 60 + 20, undefined);
          A.store.Bool(ptr + 1068 + 60 + 26, false);
          A.store.Bool(ptr + 1068 + 60 + 24, false);
          A.store.Bool(ptr + 1068 + 60 + 27, false);
          A.store.Bool(ptr + 1068 + 60 + 25, false);

          A.store.Bool(ptr + 1068 + 92 + 21, false);
          A.store.Bool(ptr + 1068 + 92 + 14, false);
          A.store.Bool(ptr + 1068 + 92 + 0, false);
          A.store.Ref(ptr + 1068 + 92 + 4, undefined);
          A.store.Bool(ptr + 1068 + 92 + 15, false);
          A.store.Bool(ptr + 1068 + 92 + 8, false);
          A.store.Bool(ptr + 1068 + 92 + 16, false);
          A.store.Bool(ptr + 1068 + 92 + 9, false);
          A.store.Bool(ptr + 1068 + 92 + 17, false);
          A.store.Bool(ptr + 1068 + 92 + 10, false);
          A.store.Bool(ptr + 1068 + 92 + 18, false);
          A.store.Bool(ptr + 1068 + 92 + 11, false);
          A.store.Bool(ptr + 1068 + 92 + 19, false);
          A.store.Bool(ptr + 1068 + 92 + 12, false);
          A.store.Bool(ptr + 1068 + 92 + 20, false);
          A.store.Bool(ptr + 1068 + 92 + 13, false);

          A.store.Bool(ptr + 1068 + 116 + 33, false);
          A.store.Bool(ptr + 1068 + 116 + 26, false);
          A.store.Int32(ptr + 1068 + 116 + 0, 0);
          A.store.Ref(ptr + 1068 + 116 + 4, undefined);
          A.store.Bool(ptr + 1068 + 116 + 27, false);
          A.store.Int32(ptr + 1068 + 116 + 8, 0);
          A.store.Bool(ptr + 1068 + 116 + 28, false);
          A.store.Int32(ptr + 1068 + 116 + 12, 0);
          A.store.Bool(ptr + 1068 + 116 + 29, false);
          A.store.Int32(ptr + 1068 + 116 + 16, 0);
          A.store.Bool(ptr + 1068 + 116 + 30, false);
          A.store.Int32(ptr + 1068 + 116 + 20, 0);
          A.store.Bool(ptr + 1068 + 116 + 31, false);
          A.store.Bool(ptr + 1068 + 116 + 24, false);
          A.store.Bool(ptr + 1068 + 116 + 32, false);
          A.store.Bool(ptr + 1068 + 116 + 25, false);

          A.store.Bool(ptr + 1068 + 152 + 28, false);
          A.store.Ref(ptr + 1068 + 152 + 0, undefined);
          A.store.Ref(ptr + 1068 + 152 + 4, undefined);
          A.store.Ref(ptr + 1068 + 152 + 8, undefined);
          A.store.Ref(ptr + 1068 + 152 + 12, undefined);
          A.store.Ref(ptr + 1068 + 152 + 16, undefined);
          A.store.Ref(ptr + 1068 + 152 + 20, undefined);
          A.store.Bool(ptr + 1068 + 152 + 26, false);
          A.store.Bool(ptr + 1068 + 152 + 24, false);
          A.store.Bool(ptr + 1068 + 152 + 27, false);
          A.store.Bool(ptr + 1068 + 152 + 25, false);

          A.store.Bool(ptr + 1068 + 184 + 28, false);
          A.store.Ref(ptr + 1068 + 184 + 0, undefined);
          A.store.Ref(ptr + 1068 + 184 + 4, undefined);
          A.store.Ref(ptr + 1068 + 184 + 8, undefined);
          A.store.Ref(ptr + 1068 + 184 + 12, undefined);
          A.store.Ref(ptr + 1068 + 184 + 16, undefined);
          A.store.Ref(ptr + 1068 + 184 + 20, undefined);
          A.store.Bool(ptr + 1068 + 184 + 26, false);
          A.store.Bool(ptr + 1068 + 184 + 24, false);
          A.store.Bool(ptr + 1068 + 184 + 27, false);
          A.store.Bool(ptr + 1068 + 184 + 25, false);
          A.store.Bool(ptr + 1068 + 221, false);
          A.store.Int32(ptr + 1068 + 216, 0);
        } else {
          A.store.Bool(ptr + 1068 + 222, true);

          if (typeof x["WiFi"]["AllowGatewayARPPolling"] === "undefined") {
            A.store.Bool(ptr + 1068 + 0 + 21, false);
            A.store.Bool(ptr + 1068 + 0 + 14, false);
            A.store.Bool(ptr + 1068 + 0 + 0, false);
            A.store.Ref(ptr + 1068 + 0 + 4, undefined);
            A.store.Bool(ptr + 1068 + 0 + 15, false);
            A.store.Bool(ptr + 1068 + 0 + 8, false);
            A.store.Bool(ptr + 1068 + 0 + 16, false);
            A.store.Bool(ptr + 1068 + 0 + 9, false);
            A.store.Bool(ptr + 1068 + 0 + 17, false);
            A.store.Bool(ptr + 1068 + 0 + 10, false);
            A.store.Bool(ptr + 1068 + 0 + 18, false);
            A.store.Bool(ptr + 1068 + 0 + 11, false);
            A.store.Bool(ptr + 1068 + 0 + 19, false);
            A.store.Bool(ptr + 1068 + 0 + 12, false);
            A.store.Bool(ptr + 1068 + 0 + 20, false);
            A.store.Bool(ptr + 1068 + 0 + 13, false);
          } else {
            A.store.Bool(ptr + 1068 + 0 + 21, true);
            A.store.Bool(ptr + 1068 + 0 + 14, "Active" in x["WiFi"]["AllowGatewayARPPolling"] ? true : false);
            A.store.Bool(ptr + 1068 + 0 + 0, x["WiFi"]["AllowGatewayARPPolling"]["Active"] ? true : false);
            A.store.Ref(ptr + 1068 + 0 + 4, x["WiFi"]["AllowGatewayARPPolling"]["Effective"]);
            A.store.Bool(ptr + 1068 + 0 + 15, "UserPolicy" in x["WiFi"]["AllowGatewayARPPolling"] ? true : false);
            A.store.Bool(ptr + 1068 + 0 + 8, x["WiFi"]["AllowGatewayARPPolling"]["UserPolicy"] ? true : false);
            A.store.Bool(ptr + 1068 + 0 + 16, "DevicePolicy" in x["WiFi"]["AllowGatewayARPPolling"] ? true : false);
            A.store.Bool(ptr + 1068 + 0 + 9, x["WiFi"]["AllowGatewayARPPolling"]["DevicePolicy"] ? true : false);
            A.store.Bool(ptr + 1068 + 0 + 17, "UserSetting" in x["WiFi"]["AllowGatewayARPPolling"] ? true : false);
            A.store.Bool(ptr + 1068 + 0 + 10, x["WiFi"]["AllowGatewayARPPolling"]["UserSetting"] ? true : false);
            A.store.Bool(ptr + 1068 + 0 + 18, "SharedSetting" in x["WiFi"]["AllowGatewayARPPolling"] ? true : false);
            A.store.Bool(ptr + 1068 + 0 + 11, x["WiFi"]["AllowGatewayARPPolling"]["SharedSetting"] ? true : false);
            A.store.Bool(ptr + 1068 + 0 + 19, "UserEditable" in x["WiFi"]["AllowGatewayARPPolling"] ? true : false);
            A.store.Bool(ptr + 1068 + 0 + 12, x["WiFi"]["AllowGatewayARPPolling"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 1068 + 0 + 20, "DeviceEditable" in x["WiFi"]["AllowGatewayARPPolling"] ? true : false);
            A.store.Bool(ptr + 1068 + 0 + 13, x["WiFi"]["AllowGatewayARPPolling"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["WiFi"]["AutoConnect"] === "undefined") {
            A.store.Bool(ptr + 1068 + 24 + 21, false);
            A.store.Bool(ptr + 1068 + 24 + 14, false);
            A.store.Bool(ptr + 1068 + 24 + 0, false);
            A.store.Ref(ptr + 1068 + 24 + 4, undefined);
            A.store.Bool(ptr + 1068 + 24 + 15, false);
            A.store.Bool(ptr + 1068 + 24 + 8, false);
            A.store.Bool(ptr + 1068 + 24 + 16, false);
            A.store.Bool(ptr + 1068 + 24 + 9, false);
            A.store.Bool(ptr + 1068 + 24 + 17, false);
            A.store.Bool(ptr + 1068 + 24 + 10, false);
            A.store.Bool(ptr + 1068 + 24 + 18, false);
            A.store.Bool(ptr + 1068 + 24 + 11, false);
            A.store.Bool(ptr + 1068 + 24 + 19, false);
            A.store.Bool(ptr + 1068 + 24 + 12, false);
            A.store.Bool(ptr + 1068 + 24 + 20, false);
            A.store.Bool(ptr + 1068 + 24 + 13, false);
          } else {
            A.store.Bool(ptr + 1068 + 24 + 21, true);
            A.store.Bool(ptr + 1068 + 24 + 14, "Active" in x["WiFi"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 1068 + 24 + 0, x["WiFi"]["AutoConnect"]["Active"] ? true : false);
            A.store.Ref(ptr + 1068 + 24 + 4, x["WiFi"]["AutoConnect"]["Effective"]);
            A.store.Bool(ptr + 1068 + 24 + 15, "UserPolicy" in x["WiFi"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 1068 + 24 + 8, x["WiFi"]["AutoConnect"]["UserPolicy"] ? true : false);
            A.store.Bool(ptr + 1068 + 24 + 16, "DevicePolicy" in x["WiFi"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 1068 + 24 + 9, x["WiFi"]["AutoConnect"]["DevicePolicy"] ? true : false);
            A.store.Bool(ptr + 1068 + 24 + 17, "UserSetting" in x["WiFi"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 1068 + 24 + 10, x["WiFi"]["AutoConnect"]["UserSetting"] ? true : false);
            A.store.Bool(ptr + 1068 + 24 + 18, "SharedSetting" in x["WiFi"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 1068 + 24 + 11, x["WiFi"]["AutoConnect"]["SharedSetting"] ? true : false);
            A.store.Bool(ptr + 1068 + 24 + 19, "UserEditable" in x["WiFi"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 1068 + 24 + 12, x["WiFi"]["AutoConnect"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 1068 + 24 + 20, "DeviceEditable" in x["WiFi"]["AutoConnect"] ? true : false);
            A.store.Bool(ptr + 1068 + 24 + 13, x["WiFi"]["AutoConnect"]["DeviceEditable"] ? true : false);
          }
          A.store.Ref(ptr + 1068 + 48, x["WiFi"]["BSSID"]);
          A.store.Bool(ptr + 1068 + 220, "Frequency" in x["WiFi"] ? true : false);
          A.store.Int32(ptr + 1068 + 52, x["WiFi"]["Frequency"] === undefined ? 0 : (x["WiFi"]["Frequency"] as number));
          A.store.Ref(ptr + 1068 + 56, x["WiFi"]["FrequencyList"]);

          if (typeof x["WiFi"]["HexSSID"] === "undefined") {
            A.store.Bool(ptr + 1068 + 60 + 28, false);
            A.store.Ref(ptr + 1068 + 60 + 0, undefined);
            A.store.Ref(ptr + 1068 + 60 + 4, undefined);
            A.store.Ref(ptr + 1068 + 60 + 8, undefined);
            A.store.Ref(ptr + 1068 + 60 + 12, undefined);
            A.store.Ref(ptr + 1068 + 60 + 16, undefined);
            A.store.Ref(ptr + 1068 + 60 + 20, undefined);
            A.store.Bool(ptr + 1068 + 60 + 26, false);
            A.store.Bool(ptr + 1068 + 60 + 24, false);
            A.store.Bool(ptr + 1068 + 60 + 27, false);
            A.store.Bool(ptr + 1068 + 60 + 25, false);
          } else {
            A.store.Bool(ptr + 1068 + 60 + 28, true);
            A.store.Ref(ptr + 1068 + 60 + 0, x["WiFi"]["HexSSID"]["Active"]);
            A.store.Ref(ptr + 1068 + 60 + 4, x["WiFi"]["HexSSID"]["Effective"]);
            A.store.Ref(ptr + 1068 + 60 + 8, x["WiFi"]["HexSSID"]["UserPolicy"]);
            A.store.Ref(ptr + 1068 + 60 + 12, x["WiFi"]["HexSSID"]["DevicePolicy"]);
            A.store.Ref(ptr + 1068 + 60 + 16, x["WiFi"]["HexSSID"]["UserSetting"]);
            A.store.Ref(ptr + 1068 + 60 + 20, x["WiFi"]["HexSSID"]["SharedSetting"]);
            A.store.Bool(ptr + 1068 + 60 + 26, "UserEditable" in x["WiFi"]["HexSSID"] ? true : false);
            A.store.Bool(ptr + 1068 + 60 + 24, x["WiFi"]["HexSSID"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 1068 + 60 + 27, "DeviceEditable" in x["WiFi"]["HexSSID"] ? true : false);
            A.store.Bool(ptr + 1068 + 60 + 25, x["WiFi"]["HexSSID"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["WiFi"]["HiddenSSID"] === "undefined") {
            A.store.Bool(ptr + 1068 + 92 + 21, false);
            A.store.Bool(ptr + 1068 + 92 + 14, false);
            A.store.Bool(ptr + 1068 + 92 + 0, false);
            A.store.Ref(ptr + 1068 + 92 + 4, undefined);
            A.store.Bool(ptr + 1068 + 92 + 15, false);
            A.store.Bool(ptr + 1068 + 92 + 8, false);
            A.store.Bool(ptr + 1068 + 92 + 16, false);
            A.store.Bool(ptr + 1068 + 92 + 9, false);
            A.store.Bool(ptr + 1068 + 92 + 17, false);
            A.store.Bool(ptr + 1068 + 92 + 10, false);
            A.store.Bool(ptr + 1068 + 92 + 18, false);
            A.store.Bool(ptr + 1068 + 92 + 11, false);
            A.store.Bool(ptr + 1068 + 92 + 19, false);
            A.store.Bool(ptr + 1068 + 92 + 12, false);
            A.store.Bool(ptr + 1068 + 92 + 20, false);
            A.store.Bool(ptr + 1068 + 92 + 13, false);
          } else {
            A.store.Bool(ptr + 1068 + 92 + 21, true);
            A.store.Bool(ptr + 1068 + 92 + 14, "Active" in x["WiFi"]["HiddenSSID"] ? true : false);
            A.store.Bool(ptr + 1068 + 92 + 0, x["WiFi"]["HiddenSSID"]["Active"] ? true : false);
            A.store.Ref(ptr + 1068 + 92 + 4, x["WiFi"]["HiddenSSID"]["Effective"]);
            A.store.Bool(ptr + 1068 + 92 + 15, "UserPolicy" in x["WiFi"]["HiddenSSID"] ? true : false);
            A.store.Bool(ptr + 1068 + 92 + 8, x["WiFi"]["HiddenSSID"]["UserPolicy"] ? true : false);
            A.store.Bool(ptr + 1068 + 92 + 16, "DevicePolicy" in x["WiFi"]["HiddenSSID"] ? true : false);
            A.store.Bool(ptr + 1068 + 92 + 9, x["WiFi"]["HiddenSSID"]["DevicePolicy"] ? true : false);
            A.store.Bool(ptr + 1068 + 92 + 17, "UserSetting" in x["WiFi"]["HiddenSSID"] ? true : false);
            A.store.Bool(ptr + 1068 + 92 + 10, x["WiFi"]["HiddenSSID"]["UserSetting"] ? true : false);
            A.store.Bool(ptr + 1068 + 92 + 18, "SharedSetting" in x["WiFi"]["HiddenSSID"] ? true : false);
            A.store.Bool(ptr + 1068 + 92 + 11, x["WiFi"]["HiddenSSID"]["SharedSetting"] ? true : false);
            A.store.Bool(ptr + 1068 + 92 + 19, "UserEditable" in x["WiFi"]["HiddenSSID"] ? true : false);
            A.store.Bool(ptr + 1068 + 92 + 12, x["WiFi"]["HiddenSSID"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 1068 + 92 + 20, "DeviceEditable" in x["WiFi"]["HiddenSSID"] ? true : false);
            A.store.Bool(ptr + 1068 + 92 + 13, x["WiFi"]["HiddenSSID"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["WiFi"]["RoamThreshold"] === "undefined") {
            A.store.Bool(ptr + 1068 + 116 + 33, false);
            A.store.Bool(ptr + 1068 + 116 + 26, false);
            A.store.Int32(ptr + 1068 + 116 + 0, 0);
            A.store.Ref(ptr + 1068 + 116 + 4, undefined);
            A.store.Bool(ptr + 1068 + 116 + 27, false);
            A.store.Int32(ptr + 1068 + 116 + 8, 0);
            A.store.Bool(ptr + 1068 + 116 + 28, false);
            A.store.Int32(ptr + 1068 + 116 + 12, 0);
            A.store.Bool(ptr + 1068 + 116 + 29, false);
            A.store.Int32(ptr + 1068 + 116 + 16, 0);
            A.store.Bool(ptr + 1068 + 116 + 30, false);
            A.store.Int32(ptr + 1068 + 116 + 20, 0);
            A.store.Bool(ptr + 1068 + 116 + 31, false);
            A.store.Bool(ptr + 1068 + 116 + 24, false);
            A.store.Bool(ptr + 1068 + 116 + 32, false);
            A.store.Bool(ptr + 1068 + 116 + 25, false);
          } else {
            A.store.Bool(ptr + 1068 + 116 + 33, true);
            A.store.Bool(ptr + 1068 + 116 + 26, "Active" in x["WiFi"]["RoamThreshold"] ? true : false);
            A.store.Int32(
              ptr + 1068 + 116 + 0,
              x["WiFi"]["RoamThreshold"]["Active"] === undefined ? 0 : (x["WiFi"]["RoamThreshold"]["Active"] as number)
            );
            A.store.Ref(ptr + 1068 + 116 + 4, x["WiFi"]["RoamThreshold"]["Effective"]);
            A.store.Bool(ptr + 1068 + 116 + 27, "UserPolicy" in x["WiFi"]["RoamThreshold"] ? true : false);
            A.store.Int32(
              ptr + 1068 + 116 + 8,
              x["WiFi"]["RoamThreshold"]["UserPolicy"] === undefined
                ? 0
                : (x["WiFi"]["RoamThreshold"]["UserPolicy"] as number)
            );
            A.store.Bool(ptr + 1068 + 116 + 28, "DevicePolicy" in x["WiFi"]["RoamThreshold"] ? true : false);
            A.store.Int32(
              ptr + 1068 + 116 + 12,
              x["WiFi"]["RoamThreshold"]["DevicePolicy"] === undefined
                ? 0
                : (x["WiFi"]["RoamThreshold"]["DevicePolicy"] as number)
            );
            A.store.Bool(ptr + 1068 + 116 + 29, "UserSetting" in x["WiFi"]["RoamThreshold"] ? true : false);
            A.store.Int32(
              ptr + 1068 + 116 + 16,
              x["WiFi"]["RoamThreshold"]["UserSetting"] === undefined
                ? 0
                : (x["WiFi"]["RoamThreshold"]["UserSetting"] as number)
            );
            A.store.Bool(ptr + 1068 + 116 + 30, "SharedSetting" in x["WiFi"]["RoamThreshold"] ? true : false);
            A.store.Int32(
              ptr + 1068 + 116 + 20,
              x["WiFi"]["RoamThreshold"]["SharedSetting"] === undefined
                ? 0
                : (x["WiFi"]["RoamThreshold"]["SharedSetting"] as number)
            );
            A.store.Bool(ptr + 1068 + 116 + 31, "UserEditable" in x["WiFi"]["RoamThreshold"] ? true : false);
            A.store.Bool(ptr + 1068 + 116 + 24, x["WiFi"]["RoamThreshold"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 1068 + 116 + 32, "DeviceEditable" in x["WiFi"]["RoamThreshold"] ? true : false);
            A.store.Bool(ptr + 1068 + 116 + 25, x["WiFi"]["RoamThreshold"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["WiFi"]["SSID"] === "undefined") {
            A.store.Bool(ptr + 1068 + 152 + 28, false);
            A.store.Ref(ptr + 1068 + 152 + 0, undefined);
            A.store.Ref(ptr + 1068 + 152 + 4, undefined);
            A.store.Ref(ptr + 1068 + 152 + 8, undefined);
            A.store.Ref(ptr + 1068 + 152 + 12, undefined);
            A.store.Ref(ptr + 1068 + 152 + 16, undefined);
            A.store.Ref(ptr + 1068 + 152 + 20, undefined);
            A.store.Bool(ptr + 1068 + 152 + 26, false);
            A.store.Bool(ptr + 1068 + 152 + 24, false);
            A.store.Bool(ptr + 1068 + 152 + 27, false);
            A.store.Bool(ptr + 1068 + 152 + 25, false);
          } else {
            A.store.Bool(ptr + 1068 + 152 + 28, true);
            A.store.Ref(ptr + 1068 + 152 + 0, x["WiFi"]["SSID"]["Active"]);
            A.store.Ref(ptr + 1068 + 152 + 4, x["WiFi"]["SSID"]["Effective"]);
            A.store.Ref(ptr + 1068 + 152 + 8, x["WiFi"]["SSID"]["UserPolicy"]);
            A.store.Ref(ptr + 1068 + 152 + 12, x["WiFi"]["SSID"]["DevicePolicy"]);
            A.store.Ref(ptr + 1068 + 152 + 16, x["WiFi"]["SSID"]["UserSetting"]);
            A.store.Ref(ptr + 1068 + 152 + 20, x["WiFi"]["SSID"]["SharedSetting"]);
            A.store.Bool(ptr + 1068 + 152 + 26, "UserEditable" in x["WiFi"]["SSID"] ? true : false);
            A.store.Bool(ptr + 1068 + 152 + 24, x["WiFi"]["SSID"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 1068 + 152 + 27, "DeviceEditable" in x["WiFi"]["SSID"] ? true : false);
            A.store.Bool(ptr + 1068 + 152 + 25, x["WiFi"]["SSID"]["DeviceEditable"] ? true : false);
          }

          if (typeof x["WiFi"]["Security"] === "undefined") {
            A.store.Bool(ptr + 1068 + 184 + 28, false);
            A.store.Ref(ptr + 1068 + 184 + 0, undefined);
            A.store.Ref(ptr + 1068 + 184 + 4, undefined);
            A.store.Ref(ptr + 1068 + 184 + 8, undefined);
            A.store.Ref(ptr + 1068 + 184 + 12, undefined);
            A.store.Ref(ptr + 1068 + 184 + 16, undefined);
            A.store.Ref(ptr + 1068 + 184 + 20, undefined);
            A.store.Bool(ptr + 1068 + 184 + 26, false);
            A.store.Bool(ptr + 1068 + 184 + 24, false);
            A.store.Bool(ptr + 1068 + 184 + 27, false);
            A.store.Bool(ptr + 1068 + 184 + 25, false);
          } else {
            A.store.Bool(ptr + 1068 + 184 + 28, true);
            A.store.Ref(ptr + 1068 + 184 + 0, x["WiFi"]["Security"]["Active"]);
            A.store.Ref(ptr + 1068 + 184 + 4, x["WiFi"]["Security"]["Effective"]);
            A.store.Ref(ptr + 1068 + 184 + 8, x["WiFi"]["Security"]["UserPolicy"]);
            A.store.Ref(ptr + 1068 + 184 + 12, x["WiFi"]["Security"]["DevicePolicy"]);
            A.store.Ref(ptr + 1068 + 184 + 16, x["WiFi"]["Security"]["UserSetting"]);
            A.store.Ref(ptr + 1068 + 184 + 20, x["WiFi"]["Security"]["SharedSetting"]);
            A.store.Bool(ptr + 1068 + 184 + 26, "UserEditable" in x["WiFi"]["Security"] ? true : false);
            A.store.Bool(ptr + 1068 + 184 + 24, x["WiFi"]["Security"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 1068 + 184 + 27, "DeviceEditable" in x["WiFi"]["Security"] ? true : false);
            A.store.Bool(ptr + 1068 + 184 + 25, x["WiFi"]["Security"]["DeviceEditable"] ? true : false);
          }
          A.store.Bool(ptr + 1068 + 221, "SignalStrength" in x["WiFi"] ? true : false);
          A.store.Int32(
            ptr + 1068 + 216,
            x["WiFi"]["SignalStrength"] === undefined ? 0 : (x["WiFi"]["SignalStrength"] as number)
          );
        }
      }
    },
    "load_ManagedProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 134)) {
        x["Cellular"] = {};
        if (A.load.Bool(ptr + 0 + 0 + 21)) {
          x["Cellular"]["AutoConnect"] = {};
          if (A.load.Bool(ptr + 0 + 0 + 14)) {
            x["Cellular"]["AutoConnect"]["Active"] = A.load.Bool(ptr + 0 + 0 + 0);
          } else {
            delete x["Cellular"]["AutoConnect"]["Active"];
          }
          x["Cellular"]["AutoConnect"]["Effective"] = A.load.Ref(ptr + 0 + 0 + 4, undefined);
          if (A.load.Bool(ptr + 0 + 0 + 15)) {
            x["Cellular"]["AutoConnect"]["UserPolicy"] = A.load.Bool(ptr + 0 + 0 + 8);
          } else {
            delete x["Cellular"]["AutoConnect"]["UserPolicy"];
          }
          if (A.load.Bool(ptr + 0 + 0 + 16)) {
            x["Cellular"]["AutoConnect"]["DevicePolicy"] = A.load.Bool(ptr + 0 + 0 + 9);
          } else {
            delete x["Cellular"]["AutoConnect"]["DevicePolicy"];
          }
          if (A.load.Bool(ptr + 0 + 0 + 17)) {
            x["Cellular"]["AutoConnect"]["UserSetting"] = A.load.Bool(ptr + 0 + 0 + 10);
          } else {
            delete x["Cellular"]["AutoConnect"]["UserSetting"];
          }
          if (A.load.Bool(ptr + 0 + 0 + 18)) {
            x["Cellular"]["AutoConnect"]["SharedSetting"] = A.load.Bool(ptr + 0 + 0 + 11);
          } else {
            delete x["Cellular"]["AutoConnect"]["SharedSetting"];
          }
          if (A.load.Bool(ptr + 0 + 0 + 19)) {
            x["Cellular"]["AutoConnect"]["UserEditable"] = A.load.Bool(ptr + 0 + 0 + 12);
          } else {
            delete x["Cellular"]["AutoConnect"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 0 + 0 + 20)) {
            x["Cellular"]["AutoConnect"]["DeviceEditable"] = A.load.Bool(ptr + 0 + 0 + 13);
          } else {
            delete x["Cellular"]["AutoConnect"]["DeviceEditable"];
          }
        } else {
          delete x["Cellular"]["AutoConnect"];
        }
        x["Cellular"]["ActivationType"] = A.load.Ref(ptr + 0 + 24, undefined);
        x["Cellular"]["ActivationState"] = A.load.Enum(ptr + 0 + 28, [
          "Activated",
          "Activating",
          "NotActivated",
          "PartiallyActivated",
        ]);
        if (A.load.Bool(ptr + 0 + 129)) {
          x["Cellular"]["AllowRoaming"] = A.load.Bool(ptr + 0 + 32);
        } else {
          delete x["Cellular"]["AllowRoaming"];
        }
        x["Cellular"]["Family"] = A.load.Ref(ptr + 0 + 36, undefined);
        x["Cellular"]["FirmwareRevision"] = A.load.Ref(ptr + 0 + 40, undefined);
        x["Cellular"]["FoundNetworks"] = A.load.Ref(ptr + 0 + 44, undefined);
        x["Cellular"]["HardwareRevision"] = A.load.Ref(ptr + 0 + 48, undefined);
        x["Cellular"]["HomeProvider"] = A.load.Ref(ptr + 0 + 52, undefined);
        x["Cellular"]["Manufacturer"] = A.load.Ref(ptr + 0 + 56, undefined);
        x["Cellular"]["ModelID"] = A.load.Ref(ptr + 0 + 60, undefined);
        x["Cellular"]["NetworkTechnology"] = A.load.Ref(ptr + 0 + 64, undefined);
        if (A.load.Bool(ptr + 0 + 68 + 12)) {
          x["Cellular"]["PaymentPortal"] = {};
          x["Cellular"]["PaymentPortal"]["Method"] = A.load.Ref(ptr + 0 + 68 + 0, undefined);
          x["Cellular"]["PaymentPortal"]["PostData"] = A.load.Ref(ptr + 0 + 68 + 4, undefined);
          x["Cellular"]["PaymentPortal"]["Url"] = A.load.Ref(ptr + 0 + 68 + 8, undefined);
        } else {
          delete x["Cellular"]["PaymentPortal"];
        }
        x["Cellular"]["RoamingState"] = A.load.Ref(ptr + 0 + 84, undefined);
        if (A.load.Bool(ptr + 0 + 130)) {
          x["Cellular"]["Scanning"] = A.load.Bool(ptr + 0 + 88);
        } else {
          delete x["Cellular"]["Scanning"];
        }
        if (A.load.Bool(ptr + 0 + 92 + 12)) {
          x["Cellular"]["ServingOperator"] = {};
          x["Cellular"]["ServingOperator"]["Name"] = A.load.Ref(ptr + 0 + 92 + 0, undefined);
          x["Cellular"]["ServingOperator"]["Code"] = A.load.Ref(ptr + 0 + 92 + 4, undefined);
          x["Cellular"]["ServingOperator"]["Country"] = A.load.Ref(ptr + 0 + 92 + 8, undefined);
        } else {
          delete x["Cellular"]["ServingOperator"];
        }
        if (A.load.Bool(ptr + 0 + 108 + 14)) {
          x["Cellular"]["SIMLockStatus"] = {};
          x["Cellular"]["SIMLockStatus"]["LockType"] = A.load.Ref(ptr + 0 + 108 + 0, undefined);
          if (A.load.Bool(ptr + 0 + 108 + 12)) {
            x["Cellular"]["SIMLockStatus"]["LockEnabled"] = A.load.Bool(ptr + 0 + 108 + 4);
          } else {
            delete x["Cellular"]["SIMLockStatus"]["LockEnabled"];
          }
          if (A.load.Bool(ptr + 0 + 108 + 13)) {
            x["Cellular"]["SIMLockStatus"]["RetriesLeft"] = A.load.Int32(ptr + 0 + 108 + 8);
          } else {
            delete x["Cellular"]["SIMLockStatus"]["RetriesLeft"];
          }
        } else {
          delete x["Cellular"]["SIMLockStatus"];
        }
        if (A.load.Bool(ptr + 0 + 131)) {
          x["Cellular"]["SIMPresent"] = A.load.Bool(ptr + 0 + 123);
        } else {
          delete x["Cellular"]["SIMPresent"];
        }
        if (A.load.Bool(ptr + 0 + 132)) {
          x["Cellular"]["SignalStrength"] = A.load.Int32(ptr + 0 + 124);
        } else {
          delete x["Cellular"]["SignalStrength"];
        }
        if (A.load.Bool(ptr + 0 + 133)) {
          x["Cellular"]["SupportNetworkScan"] = A.load.Bool(ptr + 0 + 128);
        } else {
          delete x["Cellular"]["SupportNetworkScan"];
        }
      } else {
        delete x["Cellular"];
      }
      if (A.load.Bool(ptr + 1291)) {
        x["Connectable"] = A.load.Bool(ptr + 135);
      } else {
        delete x["Connectable"];
      }
      x["ConnectionState"] = A.load.Enum(ptr + 136, ["Connected", "Connecting", "NotConnected"]);
      x["ErrorState"] = A.load.Ref(ptr + 140, undefined);
      if (A.load.Bool(ptr + 144 + 53)) {
        x["Ethernet"] = {};
        if (A.load.Bool(ptr + 144 + 0 + 21)) {
          x["Ethernet"]["AutoConnect"] = {};
          if (A.load.Bool(ptr + 144 + 0 + 14)) {
            x["Ethernet"]["AutoConnect"]["Active"] = A.load.Bool(ptr + 144 + 0 + 0);
          } else {
            delete x["Ethernet"]["AutoConnect"]["Active"];
          }
          x["Ethernet"]["AutoConnect"]["Effective"] = A.load.Ref(ptr + 144 + 0 + 4, undefined);
          if (A.load.Bool(ptr + 144 + 0 + 15)) {
            x["Ethernet"]["AutoConnect"]["UserPolicy"] = A.load.Bool(ptr + 144 + 0 + 8);
          } else {
            delete x["Ethernet"]["AutoConnect"]["UserPolicy"];
          }
          if (A.load.Bool(ptr + 144 + 0 + 16)) {
            x["Ethernet"]["AutoConnect"]["DevicePolicy"] = A.load.Bool(ptr + 144 + 0 + 9);
          } else {
            delete x["Ethernet"]["AutoConnect"]["DevicePolicy"];
          }
          if (A.load.Bool(ptr + 144 + 0 + 17)) {
            x["Ethernet"]["AutoConnect"]["UserSetting"] = A.load.Bool(ptr + 144 + 0 + 10);
          } else {
            delete x["Ethernet"]["AutoConnect"]["UserSetting"];
          }
          if (A.load.Bool(ptr + 144 + 0 + 18)) {
            x["Ethernet"]["AutoConnect"]["SharedSetting"] = A.load.Bool(ptr + 144 + 0 + 11);
          } else {
            delete x["Ethernet"]["AutoConnect"]["SharedSetting"];
          }
          if (A.load.Bool(ptr + 144 + 0 + 19)) {
            x["Ethernet"]["AutoConnect"]["UserEditable"] = A.load.Bool(ptr + 144 + 0 + 12);
          } else {
            delete x["Ethernet"]["AutoConnect"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 144 + 0 + 20)) {
            x["Ethernet"]["AutoConnect"]["DeviceEditable"] = A.load.Bool(ptr + 144 + 0 + 13);
          } else {
            delete x["Ethernet"]["AutoConnect"]["DeviceEditable"];
          }
        } else {
          delete x["Ethernet"]["AutoConnect"];
        }
        if (A.load.Bool(ptr + 144 + 24 + 28)) {
          x["Ethernet"]["Authentication"] = {};
          x["Ethernet"]["Authentication"]["Active"] = A.load.Ref(ptr + 144 + 24 + 0, undefined);
          x["Ethernet"]["Authentication"]["Effective"] = A.load.Ref(ptr + 144 + 24 + 4, undefined);
          x["Ethernet"]["Authentication"]["UserPolicy"] = A.load.Ref(ptr + 144 + 24 + 8, undefined);
          x["Ethernet"]["Authentication"]["DevicePolicy"] = A.load.Ref(ptr + 144 + 24 + 12, undefined);
          x["Ethernet"]["Authentication"]["UserSetting"] = A.load.Ref(ptr + 144 + 24 + 16, undefined);
          x["Ethernet"]["Authentication"]["SharedSetting"] = A.load.Ref(ptr + 144 + 24 + 20, undefined);
          if (A.load.Bool(ptr + 144 + 24 + 26)) {
            x["Ethernet"]["Authentication"]["UserEditable"] = A.load.Bool(ptr + 144 + 24 + 24);
          } else {
            delete x["Ethernet"]["Authentication"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 144 + 24 + 27)) {
            x["Ethernet"]["Authentication"]["DeviceEditable"] = A.load.Bool(ptr + 144 + 24 + 25);
          } else {
            delete x["Ethernet"]["Authentication"]["DeviceEditable"];
          }
        } else {
          delete x["Ethernet"]["Authentication"];
        }
      } else {
        delete x["Ethernet"];
      }
      x["GUID"] = A.load.Ref(ptr + 200, undefined);
      if (A.load.Bool(ptr + 204 + 28)) {
        x["IPAddressConfigType"] = {};
        x["IPAddressConfigType"]["Active"] = A.load.Enum(ptr + 204 + 0, ["DHCP", "Static"]);
        x["IPAddressConfigType"]["Effective"] = A.load.Ref(ptr + 204 + 4, undefined);
        x["IPAddressConfigType"]["UserPolicy"] = A.load.Enum(ptr + 204 + 8, ["DHCP", "Static"]);
        x["IPAddressConfigType"]["DevicePolicy"] = A.load.Enum(ptr + 204 + 12, ["DHCP", "Static"]);
        x["IPAddressConfigType"]["UserSetting"] = A.load.Enum(ptr + 204 + 16, ["DHCP", "Static"]);
        x["IPAddressConfigType"]["SharedSetting"] = A.load.Enum(ptr + 204 + 20, ["DHCP", "Static"]);
        if (A.load.Bool(ptr + 204 + 26)) {
          x["IPAddressConfigType"]["UserEditable"] = A.load.Bool(ptr + 204 + 24);
        } else {
          delete x["IPAddressConfigType"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 204 + 27)) {
          x["IPAddressConfigType"]["DeviceEditable"] = A.load.Bool(ptr + 204 + 25);
        } else {
          delete x["IPAddressConfigType"]["DeviceEditable"];
        }
      } else {
        delete x["IPAddressConfigType"];
      }
      x["IPConfigs"] = A.load.Ref(ptr + 236, undefined);
      x["MacAddress"] = A.load.Ref(ptr + 240, undefined);
      if (A.load.Bool(ptr + 244 + 21)) {
        x["Metered"] = {};
        if (A.load.Bool(ptr + 244 + 14)) {
          x["Metered"]["Active"] = A.load.Bool(ptr + 244 + 0);
        } else {
          delete x["Metered"]["Active"];
        }
        x["Metered"]["Effective"] = A.load.Ref(ptr + 244 + 4, undefined);
        if (A.load.Bool(ptr + 244 + 15)) {
          x["Metered"]["UserPolicy"] = A.load.Bool(ptr + 244 + 8);
        } else {
          delete x["Metered"]["UserPolicy"];
        }
        if (A.load.Bool(ptr + 244 + 16)) {
          x["Metered"]["DevicePolicy"] = A.load.Bool(ptr + 244 + 9);
        } else {
          delete x["Metered"]["DevicePolicy"];
        }
        if (A.load.Bool(ptr + 244 + 17)) {
          x["Metered"]["UserSetting"] = A.load.Bool(ptr + 244 + 10);
        } else {
          delete x["Metered"]["UserSetting"];
        }
        if (A.load.Bool(ptr + 244 + 18)) {
          x["Metered"]["SharedSetting"] = A.load.Bool(ptr + 244 + 11);
        } else {
          delete x["Metered"]["SharedSetting"];
        }
        if (A.load.Bool(ptr + 244 + 19)) {
          x["Metered"]["UserEditable"] = A.load.Bool(ptr + 244 + 12);
        } else {
          delete x["Metered"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 244 + 20)) {
          x["Metered"]["DeviceEditable"] = A.load.Bool(ptr + 244 + 13);
        } else {
          delete x["Metered"]["DeviceEditable"];
        }
      } else {
        delete x["Metered"];
      }
      if (A.load.Bool(ptr + 268 + 28)) {
        x["Name"] = {};
        x["Name"]["Active"] = A.load.Ref(ptr + 268 + 0, undefined);
        x["Name"]["Effective"] = A.load.Ref(ptr + 268 + 4, undefined);
        x["Name"]["UserPolicy"] = A.load.Ref(ptr + 268 + 8, undefined);
        x["Name"]["DevicePolicy"] = A.load.Ref(ptr + 268 + 12, undefined);
        x["Name"]["UserSetting"] = A.load.Ref(ptr + 268 + 16, undefined);
        x["Name"]["SharedSetting"] = A.load.Ref(ptr + 268 + 20, undefined);
        if (A.load.Bool(ptr + 268 + 26)) {
          x["Name"]["UserEditable"] = A.load.Bool(ptr + 268 + 24);
        } else {
          delete x["Name"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 268 + 27)) {
          x["Name"]["DeviceEditable"] = A.load.Bool(ptr + 268 + 25);
        } else {
          delete x["Name"]["DeviceEditable"];
        }
      } else {
        delete x["Name"];
      }
      if (A.load.Bool(ptr + 300 + 28)) {
        x["NameServersConfigType"] = {};
        x["NameServersConfigType"]["Active"] = A.load.Enum(ptr + 300 + 0, ["DHCP", "Static"]);
        x["NameServersConfigType"]["Effective"] = A.load.Ref(ptr + 300 + 4, undefined);
        x["NameServersConfigType"]["UserPolicy"] = A.load.Enum(ptr + 300 + 8, ["DHCP", "Static"]);
        x["NameServersConfigType"]["DevicePolicy"] = A.load.Enum(ptr + 300 + 12, ["DHCP", "Static"]);
        x["NameServersConfigType"]["UserSetting"] = A.load.Enum(ptr + 300 + 16, ["DHCP", "Static"]);
        x["NameServersConfigType"]["SharedSetting"] = A.load.Enum(ptr + 300 + 20, ["DHCP", "Static"]);
        if (A.load.Bool(ptr + 300 + 26)) {
          x["NameServersConfigType"]["UserEditable"] = A.load.Bool(ptr + 300 + 24);
        } else {
          delete x["NameServersConfigType"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 300 + 27)) {
          x["NameServersConfigType"]["DeviceEditable"] = A.load.Bool(ptr + 300 + 25);
        } else {
          delete x["NameServersConfigType"]["DeviceEditable"];
        }
      } else {
        delete x["NameServersConfigType"];
      }
      if (A.load.Bool(ptr + 332 + 33)) {
        x["Priority"] = {};
        if (A.load.Bool(ptr + 332 + 26)) {
          x["Priority"]["Active"] = A.load.Int32(ptr + 332 + 0);
        } else {
          delete x["Priority"]["Active"];
        }
        x["Priority"]["Effective"] = A.load.Ref(ptr + 332 + 4, undefined);
        if (A.load.Bool(ptr + 332 + 27)) {
          x["Priority"]["UserPolicy"] = A.load.Int32(ptr + 332 + 8);
        } else {
          delete x["Priority"]["UserPolicy"];
        }
        if (A.load.Bool(ptr + 332 + 28)) {
          x["Priority"]["DevicePolicy"] = A.load.Int32(ptr + 332 + 12);
        } else {
          delete x["Priority"]["DevicePolicy"];
        }
        if (A.load.Bool(ptr + 332 + 29)) {
          x["Priority"]["UserSetting"] = A.load.Int32(ptr + 332 + 16);
        } else {
          delete x["Priority"]["UserSetting"];
        }
        if (A.load.Bool(ptr + 332 + 30)) {
          x["Priority"]["SharedSetting"] = A.load.Int32(ptr + 332 + 20);
        } else {
          delete x["Priority"]["SharedSetting"];
        }
        if (A.load.Bool(ptr + 332 + 31)) {
          x["Priority"]["UserEditable"] = A.load.Bool(ptr + 332 + 24);
        } else {
          delete x["Priority"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 332 + 32)) {
          x["Priority"]["DeviceEditable"] = A.load.Bool(ptr + 332 + 25);
        } else {
          delete x["Priority"]["DeviceEditable"];
        }
      } else {
        delete x["Priority"];
      }
      if (A.load.Bool(ptr + 368 + 365)) {
        x["ProxySettings"] = {};
        if (A.load.Bool(ptr + 368 + 0 + 28)) {
          x["ProxySettings"]["Type"] = {};
          x["ProxySettings"]["Type"]["Active"] = A.load.Enum(ptr + 368 + 0 + 0, ["Direct", "Manual", "PAC", "WPAD"]);
          x["ProxySettings"]["Type"]["Effective"] = A.load.Ref(ptr + 368 + 0 + 4, undefined);
          x["ProxySettings"]["Type"]["UserPolicy"] = A.load.Enum(ptr + 368 + 0 + 8, [
            "Direct",
            "Manual",
            "PAC",
            "WPAD",
          ]);
          x["ProxySettings"]["Type"]["DevicePolicy"] = A.load.Enum(ptr + 368 + 0 + 12, [
            "Direct",
            "Manual",
            "PAC",
            "WPAD",
          ]);
          x["ProxySettings"]["Type"]["UserSetting"] = A.load.Enum(ptr + 368 + 0 + 16, [
            "Direct",
            "Manual",
            "PAC",
            "WPAD",
          ]);
          x["ProxySettings"]["Type"]["SharedSetting"] = A.load.Enum(ptr + 368 + 0 + 20, [
            "Direct",
            "Manual",
            "PAC",
            "WPAD",
          ]);
          if (A.load.Bool(ptr + 368 + 0 + 26)) {
            x["ProxySettings"]["Type"]["UserEditable"] = A.load.Bool(ptr + 368 + 0 + 24);
          } else {
            delete x["ProxySettings"]["Type"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 368 + 0 + 27)) {
            x["ProxySettings"]["Type"]["DeviceEditable"] = A.load.Bool(ptr + 368 + 0 + 25);
          } else {
            delete x["ProxySettings"]["Type"]["DeviceEditable"];
          }
        } else {
          delete x["ProxySettings"]["Type"];
        }
        if (A.load.Bool(ptr + 368 + 32 + 271)) {
          x["ProxySettings"]["Manual"] = {};
          if (A.load.Bool(ptr + 368 + 32 + 0 + 66)) {
            x["ProxySettings"]["Manual"]["HTTPProxy"] = {};
            if (A.load.Bool(ptr + 368 + 32 + 0 + 0 + 28)) {
              x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"] = {};
              x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["Active"] = A.load.Ref(
                ptr + 368 + 32 + 0 + 0 + 0,
                undefined
              );
              x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["Effective"] = A.load.Ref(
                ptr + 368 + 32 + 0 + 0 + 4,
                undefined
              );
              x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["UserPolicy"] = A.load.Ref(
                ptr + 368 + 32 + 0 + 0 + 8,
                undefined
              );
              x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["DevicePolicy"] = A.load.Ref(
                ptr + 368 + 32 + 0 + 0 + 12,
                undefined
              );
              x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["UserSetting"] = A.load.Ref(
                ptr + 368 + 32 + 0 + 0 + 16,
                undefined
              );
              x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["SharedSetting"] = A.load.Ref(
                ptr + 368 + 32 + 0 + 0 + 20,
                undefined
              );
              if (A.load.Bool(ptr + 368 + 32 + 0 + 0 + 26)) {
                x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["UserEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 0 + 0 + 24
                );
              } else {
                delete x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["UserEditable"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 0 + 0 + 27)) {
                x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["DeviceEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 0 + 0 + 25
                );
              } else {
                delete x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]["DeviceEditable"];
              }
            } else {
              delete x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"];
            }
            if (A.load.Bool(ptr + 368 + 32 + 0 + 32 + 33)) {
              x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"] = {};
              if (A.load.Bool(ptr + 368 + 32 + 0 + 32 + 26)) {
                x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["Active"] = A.load.Int32(ptr + 368 + 32 + 0 + 32 + 0);
              } else {
                delete x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["Active"];
              }
              x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["Effective"] = A.load.Ref(
                ptr + 368 + 32 + 0 + 32 + 4,
                undefined
              );
              if (A.load.Bool(ptr + 368 + 32 + 0 + 32 + 27)) {
                x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["UserPolicy"] = A.load.Int32(
                  ptr + 368 + 32 + 0 + 32 + 8
                );
              } else {
                delete x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["UserPolicy"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 0 + 32 + 28)) {
                x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["DevicePolicy"] = A.load.Int32(
                  ptr + 368 + 32 + 0 + 32 + 12
                );
              } else {
                delete x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["DevicePolicy"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 0 + 32 + 29)) {
                x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["UserSetting"] = A.load.Int32(
                  ptr + 368 + 32 + 0 + 32 + 16
                );
              } else {
                delete x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["UserSetting"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 0 + 32 + 30)) {
                x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["SharedSetting"] = A.load.Int32(
                  ptr + 368 + 32 + 0 + 32 + 20
                );
              } else {
                delete x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["SharedSetting"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 0 + 32 + 31)) {
                x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["UserEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 0 + 32 + 24
                );
              } else {
                delete x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["UserEditable"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 0 + 32 + 32)) {
                x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["DeviceEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 0 + 32 + 25
                );
              } else {
                delete x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"]["DeviceEditable"];
              }
            } else {
              delete x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"];
            }
          } else {
            delete x["ProxySettings"]["Manual"]["HTTPProxy"];
          }
          if (A.load.Bool(ptr + 368 + 32 + 68 + 66)) {
            x["ProxySettings"]["Manual"]["SecureHTTPProxy"] = {};
            if (A.load.Bool(ptr + 368 + 32 + 68 + 0 + 28)) {
              x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"] = {};
              x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["Active"] = A.load.Ref(
                ptr + 368 + 32 + 68 + 0 + 0,
                undefined
              );
              x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["Effective"] = A.load.Ref(
                ptr + 368 + 32 + 68 + 0 + 4,
                undefined
              );
              x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["UserPolicy"] = A.load.Ref(
                ptr + 368 + 32 + 68 + 0 + 8,
                undefined
              );
              x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["DevicePolicy"] = A.load.Ref(
                ptr + 368 + 32 + 68 + 0 + 12,
                undefined
              );
              x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["UserSetting"] = A.load.Ref(
                ptr + 368 + 32 + 68 + 0 + 16,
                undefined
              );
              x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["SharedSetting"] = A.load.Ref(
                ptr + 368 + 32 + 68 + 0 + 20,
                undefined
              );
              if (A.load.Bool(ptr + 368 + 32 + 68 + 0 + 26)) {
                x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["UserEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 68 + 0 + 24
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["UserEditable"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 68 + 0 + 27)) {
                x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["DeviceEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 68 + 0 + 25
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]["DeviceEditable"];
              }
            } else {
              delete x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"];
            }
            if (A.load.Bool(ptr + 368 + 32 + 68 + 32 + 33)) {
              x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"] = {};
              if (A.load.Bool(ptr + 368 + 32 + 68 + 32 + 26)) {
                x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["Active"] = A.load.Int32(
                  ptr + 368 + 32 + 68 + 32 + 0
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["Active"];
              }
              x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["Effective"] = A.load.Ref(
                ptr + 368 + 32 + 68 + 32 + 4,
                undefined
              );
              if (A.load.Bool(ptr + 368 + 32 + 68 + 32 + 27)) {
                x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["UserPolicy"] = A.load.Int32(
                  ptr + 368 + 32 + 68 + 32 + 8
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["UserPolicy"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 68 + 32 + 28)) {
                x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["DevicePolicy"] = A.load.Int32(
                  ptr + 368 + 32 + 68 + 32 + 12
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["DevicePolicy"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 68 + 32 + 29)) {
                x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["UserSetting"] = A.load.Int32(
                  ptr + 368 + 32 + 68 + 32 + 16
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["UserSetting"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 68 + 32 + 30)) {
                x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["SharedSetting"] = A.load.Int32(
                  ptr + 368 + 32 + 68 + 32 + 20
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["SharedSetting"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 68 + 32 + 31)) {
                x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["UserEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 68 + 32 + 24
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["UserEditable"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 68 + 32 + 32)) {
                x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["DeviceEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 68 + 32 + 25
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"]["DeviceEditable"];
              }
            } else {
              delete x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"];
            }
          } else {
            delete x["ProxySettings"]["Manual"]["SecureHTTPProxy"];
          }
          if (A.load.Bool(ptr + 368 + 32 + 136 + 66)) {
            x["ProxySettings"]["Manual"]["FTPProxy"] = {};
            if (A.load.Bool(ptr + 368 + 32 + 136 + 0 + 28)) {
              x["ProxySettings"]["Manual"]["FTPProxy"]["Host"] = {};
              x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["Active"] = A.load.Ref(
                ptr + 368 + 32 + 136 + 0 + 0,
                undefined
              );
              x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["Effective"] = A.load.Ref(
                ptr + 368 + 32 + 136 + 0 + 4,
                undefined
              );
              x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["UserPolicy"] = A.load.Ref(
                ptr + 368 + 32 + 136 + 0 + 8,
                undefined
              );
              x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["DevicePolicy"] = A.load.Ref(
                ptr + 368 + 32 + 136 + 0 + 12,
                undefined
              );
              x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["UserSetting"] = A.load.Ref(
                ptr + 368 + 32 + 136 + 0 + 16,
                undefined
              );
              x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["SharedSetting"] = A.load.Ref(
                ptr + 368 + 32 + 136 + 0 + 20,
                undefined
              );
              if (A.load.Bool(ptr + 368 + 32 + 136 + 0 + 26)) {
                x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["UserEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 136 + 0 + 24
                );
              } else {
                delete x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["UserEditable"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 136 + 0 + 27)) {
                x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["DeviceEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 136 + 0 + 25
                );
              } else {
                delete x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]["DeviceEditable"];
              }
            } else {
              delete x["ProxySettings"]["Manual"]["FTPProxy"]["Host"];
            }
            if (A.load.Bool(ptr + 368 + 32 + 136 + 32 + 33)) {
              x["ProxySettings"]["Manual"]["FTPProxy"]["Port"] = {};
              if (A.load.Bool(ptr + 368 + 32 + 136 + 32 + 26)) {
                x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["Active"] = A.load.Int32(
                  ptr + 368 + 32 + 136 + 32 + 0
                );
              } else {
                delete x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["Active"];
              }
              x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["Effective"] = A.load.Ref(
                ptr + 368 + 32 + 136 + 32 + 4,
                undefined
              );
              if (A.load.Bool(ptr + 368 + 32 + 136 + 32 + 27)) {
                x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["UserPolicy"] = A.load.Int32(
                  ptr + 368 + 32 + 136 + 32 + 8
                );
              } else {
                delete x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["UserPolicy"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 136 + 32 + 28)) {
                x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["DevicePolicy"] = A.load.Int32(
                  ptr + 368 + 32 + 136 + 32 + 12
                );
              } else {
                delete x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["DevicePolicy"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 136 + 32 + 29)) {
                x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["UserSetting"] = A.load.Int32(
                  ptr + 368 + 32 + 136 + 32 + 16
                );
              } else {
                delete x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["UserSetting"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 136 + 32 + 30)) {
                x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["SharedSetting"] = A.load.Int32(
                  ptr + 368 + 32 + 136 + 32 + 20
                );
              } else {
                delete x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["SharedSetting"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 136 + 32 + 31)) {
                x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["UserEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 136 + 32 + 24
                );
              } else {
                delete x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["UserEditable"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 136 + 32 + 32)) {
                x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["DeviceEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 136 + 32 + 25
                );
              } else {
                delete x["ProxySettings"]["Manual"]["FTPProxy"]["Port"]["DeviceEditable"];
              }
            } else {
              delete x["ProxySettings"]["Manual"]["FTPProxy"]["Port"];
            }
          } else {
            delete x["ProxySettings"]["Manual"]["FTPProxy"];
          }
          if (A.load.Bool(ptr + 368 + 32 + 204 + 66)) {
            x["ProxySettings"]["Manual"]["SOCKS"] = {};
            if (A.load.Bool(ptr + 368 + 32 + 204 + 0 + 28)) {
              x["ProxySettings"]["Manual"]["SOCKS"]["Host"] = {};
              x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["Active"] = A.load.Ref(
                ptr + 368 + 32 + 204 + 0 + 0,
                undefined
              );
              x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["Effective"] = A.load.Ref(
                ptr + 368 + 32 + 204 + 0 + 4,
                undefined
              );
              x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["UserPolicy"] = A.load.Ref(
                ptr + 368 + 32 + 204 + 0 + 8,
                undefined
              );
              x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["DevicePolicy"] = A.load.Ref(
                ptr + 368 + 32 + 204 + 0 + 12,
                undefined
              );
              x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["UserSetting"] = A.load.Ref(
                ptr + 368 + 32 + 204 + 0 + 16,
                undefined
              );
              x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["SharedSetting"] = A.load.Ref(
                ptr + 368 + 32 + 204 + 0 + 20,
                undefined
              );
              if (A.load.Bool(ptr + 368 + 32 + 204 + 0 + 26)) {
                x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["UserEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 204 + 0 + 24
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["UserEditable"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 204 + 0 + 27)) {
                x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["DeviceEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 204 + 0 + 25
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SOCKS"]["Host"]["DeviceEditable"];
              }
            } else {
              delete x["ProxySettings"]["Manual"]["SOCKS"]["Host"];
            }
            if (A.load.Bool(ptr + 368 + 32 + 204 + 32 + 33)) {
              x["ProxySettings"]["Manual"]["SOCKS"]["Port"] = {};
              if (A.load.Bool(ptr + 368 + 32 + 204 + 32 + 26)) {
                x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["Active"] = A.load.Int32(ptr + 368 + 32 + 204 + 32 + 0);
              } else {
                delete x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["Active"];
              }
              x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["Effective"] = A.load.Ref(
                ptr + 368 + 32 + 204 + 32 + 4,
                undefined
              );
              if (A.load.Bool(ptr + 368 + 32 + 204 + 32 + 27)) {
                x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["UserPolicy"] = A.load.Int32(
                  ptr + 368 + 32 + 204 + 32 + 8
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["UserPolicy"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 204 + 32 + 28)) {
                x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["DevicePolicy"] = A.load.Int32(
                  ptr + 368 + 32 + 204 + 32 + 12
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["DevicePolicy"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 204 + 32 + 29)) {
                x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["UserSetting"] = A.load.Int32(
                  ptr + 368 + 32 + 204 + 32 + 16
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["UserSetting"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 204 + 32 + 30)) {
                x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["SharedSetting"] = A.load.Int32(
                  ptr + 368 + 32 + 204 + 32 + 20
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["SharedSetting"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 204 + 32 + 31)) {
                x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["UserEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 204 + 32 + 24
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["UserEditable"];
              }
              if (A.load.Bool(ptr + 368 + 32 + 204 + 32 + 32)) {
                x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["DeviceEditable"] = A.load.Bool(
                  ptr + 368 + 32 + 204 + 32 + 25
                );
              } else {
                delete x["ProxySettings"]["Manual"]["SOCKS"]["Port"]["DeviceEditable"];
              }
            } else {
              delete x["ProxySettings"]["Manual"]["SOCKS"]["Port"];
            }
          } else {
            delete x["ProxySettings"]["Manual"]["SOCKS"];
          }
        } else {
          delete x["ProxySettings"]["Manual"];
        }
        if (A.load.Bool(ptr + 368 + 304 + 28)) {
          x["ProxySettings"]["ExcludeDomains"] = {};
          x["ProxySettings"]["ExcludeDomains"]["Active"] = A.load.Ref(ptr + 368 + 304 + 0, undefined);
          x["ProxySettings"]["ExcludeDomains"]["Effective"] = A.load.Ref(ptr + 368 + 304 + 4, undefined);
          x["ProxySettings"]["ExcludeDomains"]["UserPolicy"] = A.load.Ref(ptr + 368 + 304 + 8, undefined);
          x["ProxySettings"]["ExcludeDomains"]["DevicePolicy"] = A.load.Ref(ptr + 368 + 304 + 12, undefined);
          x["ProxySettings"]["ExcludeDomains"]["UserSetting"] = A.load.Ref(ptr + 368 + 304 + 16, undefined);
          x["ProxySettings"]["ExcludeDomains"]["SharedSetting"] = A.load.Ref(ptr + 368 + 304 + 20, undefined);
          if (A.load.Bool(ptr + 368 + 304 + 26)) {
            x["ProxySettings"]["ExcludeDomains"]["UserEditable"] = A.load.Bool(ptr + 368 + 304 + 24);
          } else {
            delete x["ProxySettings"]["ExcludeDomains"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 368 + 304 + 27)) {
            x["ProxySettings"]["ExcludeDomains"]["DeviceEditable"] = A.load.Bool(ptr + 368 + 304 + 25);
          } else {
            delete x["ProxySettings"]["ExcludeDomains"]["DeviceEditable"];
          }
        } else {
          delete x["ProxySettings"]["ExcludeDomains"];
        }
        if (A.load.Bool(ptr + 368 + 336 + 28)) {
          x["ProxySettings"]["PAC"] = {};
          x["ProxySettings"]["PAC"]["Active"] = A.load.Ref(ptr + 368 + 336 + 0, undefined);
          x["ProxySettings"]["PAC"]["Effective"] = A.load.Ref(ptr + 368 + 336 + 4, undefined);
          x["ProxySettings"]["PAC"]["UserPolicy"] = A.load.Ref(ptr + 368 + 336 + 8, undefined);
          x["ProxySettings"]["PAC"]["DevicePolicy"] = A.load.Ref(ptr + 368 + 336 + 12, undefined);
          x["ProxySettings"]["PAC"]["UserSetting"] = A.load.Ref(ptr + 368 + 336 + 16, undefined);
          x["ProxySettings"]["PAC"]["SharedSetting"] = A.load.Ref(ptr + 368 + 336 + 20, undefined);
          if (A.load.Bool(ptr + 368 + 336 + 26)) {
            x["ProxySettings"]["PAC"]["UserEditable"] = A.load.Bool(ptr + 368 + 336 + 24);
          } else {
            delete x["ProxySettings"]["PAC"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 368 + 336 + 27)) {
            x["ProxySettings"]["PAC"]["DeviceEditable"] = A.load.Bool(ptr + 368 + 336 + 25);
          } else {
            delete x["ProxySettings"]["PAC"]["DeviceEditable"];
          }
        } else {
          delete x["ProxySettings"]["PAC"];
        }
      } else {
        delete x["ProxySettings"];
      }
      if (A.load.Bool(ptr + 1292)) {
        x["RestrictedConnectivity"] = A.load.Bool(ptr + 734);
      } else {
        delete x["RestrictedConnectivity"];
      }
      if (A.load.Bool(ptr + 736 + 193)) {
        x["StaticIPConfig"] = {};
        if (A.load.Bool(ptr + 736 + 0 + 28)) {
          x["StaticIPConfig"]["Gateway"] = {};
          x["StaticIPConfig"]["Gateway"]["Active"] = A.load.Ref(ptr + 736 + 0 + 0, undefined);
          x["StaticIPConfig"]["Gateway"]["Effective"] = A.load.Ref(ptr + 736 + 0 + 4, undefined);
          x["StaticIPConfig"]["Gateway"]["UserPolicy"] = A.load.Ref(ptr + 736 + 0 + 8, undefined);
          x["StaticIPConfig"]["Gateway"]["DevicePolicy"] = A.load.Ref(ptr + 736 + 0 + 12, undefined);
          x["StaticIPConfig"]["Gateway"]["UserSetting"] = A.load.Ref(ptr + 736 + 0 + 16, undefined);
          x["StaticIPConfig"]["Gateway"]["SharedSetting"] = A.load.Ref(ptr + 736 + 0 + 20, undefined);
          if (A.load.Bool(ptr + 736 + 0 + 26)) {
            x["StaticIPConfig"]["Gateway"]["UserEditable"] = A.load.Bool(ptr + 736 + 0 + 24);
          } else {
            delete x["StaticIPConfig"]["Gateway"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 736 + 0 + 27)) {
            x["StaticIPConfig"]["Gateway"]["DeviceEditable"] = A.load.Bool(ptr + 736 + 0 + 25);
          } else {
            delete x["StaticIPConfig"]["Gateway"]["DeviceEditable"];
          }
        } else {
          delete x["StaticIPConfig"]["Gateway"];
        }
        if (A.load.Bool(ptr + 736 + 32 + 28)) {
          x["StaticIPConfig"]["IPAddress"] = {};
          x["StaticIPConfig"]["IPAddress"]["Active"] = A.load.Ref(ptr + 736 + 32 + 0, undefined);
          x["StaticIPConfig"]["IPAddress"]["Effective"] = A.load.Ref(ptr + 736 + 32 + 4, undefined);
          x["StaticIPConfig"]["IPAddress"]["UserPolicy"] = A.load.Ref(ptr + 736 + 32 + 8, undefined);
          x["StaticIPConfig"]["IPAddress"]["DevicePolicy"] = A.load.Ref(ptr + 736 + 32 + 12, undefined);
          x["StaticIPConfig"]["IPAddress"]["UserSetting"] = A.load.Ref(ptr + 736 + 32 + 16, undefined);
          x["StaticIPConfig"]["IPAddress"]["SharedSetting"] = A.load.Ref(ptr + 736 + 32 + 20, undefined);
          if (A.load.Bool(ptr + 736 + 32 + 26)) {
            x["StaticIPConfig"]["IPAddress"]["UserEditable"] = A.load.Bool(ptr + 736 + 32 + 24);
          } else {
            delete x["StaticIPConfig"]["IPAddress"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 736 + 32 + 27)) {
            x["StaticIPConfig"]["IPAddress"]["DeviceEditable"] = A.load.Bool(ptr + 736 + 32 + 25);
          } else {
            delete x["StaticIPConfig"]["IPAddress"]["DeviceEditable"];
          }
        } else {
          delete x["StaticIPConfig"]["IPAddress"];
        }
        if (A.load.Bool(ptr + 736 + 64 + 28)) {
          x["StaticIPConfig"]["NameServers"] = {};
          x["StaticIPConfig"]["NameServers"]["Active"] = A.load.Ref(ptr + 736 + 64 + 0, undefined);
          x["StaticIPConfig"]["NameServers"]["Effective"] = A.load.Ref(ptr + 736 + 64 + 4, undefined);
          x["StaticIPConfig"]["NameServers"]["UserPolicy"] = A.load.Ref(ptr + 736 + 64 + 8, undefined);
          x["StaticIPConfig"]["NameServers"]["DevicePolicy"] = A.load.Ref(ptr + 736 + 64 + 12, undefined);
          x["StaticIPConfig"]["NameServers"]["UserSetting"] = A.load.Ref(ptr + 736 + 64 + 16, undefined);
          x["StaticIPConfig"]["NameServers"]["SharedSetting"] = A.load.Ref(ptr + 736 + 64 + 20, undefined);
          if (A.load.Bool(ptr + 736 + 64 + 26)) {
            x["StaticIPConfig"]["NameServers"]["UserEditable"] = A.load.Bool(ptr + 736 + 64 + 24);
          } else {
            delete x["StaticIPConfig"]["NameServers"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 736 + 64 + 27)) {
            x["StaticIPConfig"]["NameServers"]["DeviceEditable"] = A.load.Bool(ptr + 736 + 64 + 25);
          } else {
            delete x["StaticIPConfig"]["NameServers"]["DeviceEditable"];
          }
        } else {
          delete x["StaticIPConfig"]["NameServers"];
        }
        if (A.load.Bool(ptr + 736 + 96 + 33)) {
          x["StaticIPConfig"]["RoutingPrefix"] = {};
          if (A.load.Bool(ptr + 736 + 96 + 26)) {
            x["StaticIPConfig"]["RoutingPrefix"]["Active"] = A.load.Int32(ptr + 736 + 96 + 0);
          } else {
            delete x["StaticIPConfig"]["RoutingPrefix"]["Active"];
          }
          x["StaticIPConfig"]["RoutingPrefix"]["Effective"] = A.load.Ref(ptr + 736 + 96 + 4, undefined);
          if (A.load.Bool(ptr + 736 + 96 + 27)) {
            x["StaticIPConfig"]["RoutingPrefix"]["UserPolicy"] = A.load.Int32(ptr + 736 + 96 + 8);
          } else {
            delete x["StaticIPConfig"]["RoutingPrefix"]["UserPolicy"];
          }
          if (A.load.Bool(ptr + 736 + 96 + 28)) {
            x["StaticIPConfig"]["RoutingPrefix"]["DevicePolicy"] = A.load.Int32(ptr + 736 + 96 + 12);
          } else {
            delete x["StaticIPConfig"]["RoutingPrefix"]["DevicePolicy"];
          }
          if (A.load.Bool(ptr + 736 + 96 + 29)) {
            x["StaticIPConfig"]["RoutingPrefix"]["UserSetting"] = A.load.Int32(ptr + 736 + 96 + 16);
          } else {
            delete x["StaticIPConfig"]["RoutingPrefix"]["UserSetting"];
          }
          if (A.load.Bool(ptr + 736 + 96 + 30)) {
            x["StaticIPConfig"]["RoutingPrefix"]["SharedSetting"] = A.load.Int32(ptr + 736 + 96 + 20);
          } else {
            delete x["StaticIPConfig"]["RoutingPrefix"]["SharedSetting"];
          }
          if (A.load.Bool(ptr + 736 + 96 + 31)) {
            x["StaticIPConfig"]["RoutingPrefix"]["UserEditable"] = A.load.Bool(ptr + 736 + 96 + 24);
          } else {
            delete x["StaticIPConfig"]["RoutingPrefix"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 736 + 96 + 32)) {
            x["StaticIPConfig"]["RoutingPrefix"]["DeviceEditable"] = A.load.Bool(ptr + 736 + 96 + 25);
          } else {
            delete x["StaticIPConfig"]["RoutingPrefix"]["DeviceEditable"];
          }
        } else {
          delete x["StaticIPConfig"]["RoutingPrefix"];
        }
        if (A.load.Bool(ptr + 736 + 132 + 28)) {
          x["StaticIPConfig"]["Type"] = {};
          x["StaticIPConfig"]["Type"]["Active"] = A.load.Ref(ptr + 736 + 132 + 0, undefined);
          x["StaticIPConfig"]["Type"]["Effective"] = A.load.Ref(ptr + 736 + 132 + 4, undefined);
          x["StaticIPConfig"]["Type"]["UserPolicy"] = A.load.Ref(ptr + 736 + 132 + 8, undefined);
          x["StaticIPConfig"]["Type"]["DevicePolicy"] = A.load.Ref(ptr + 736 + 132 + 12, undefined);
          x["StaticIPConfig"]["Type"]["UserSetting"] = A.load.Ref(ptr + 736 + 132 + 16, undefined);
          x["StaticIPConfig"]["Type"]["SharedSetting"] = A.load.Ref(ptr + 736 + 132 + 20, undefined);
          if (A.load.Bool(ptr + 736 + 132 + 26)) {
            x["StaticIPConfig"]["Type"]["UserEditable"] = A.load.Bool(ptr + 736 + 132 + 24);
          } else {
            delete x["StaticIPConfig"]["Type"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 736 + 132 + 27)) {
            x["StaticIPConfig"]["Type"]["DeviceEditable"] = A.load.Bool(ptr + 736 + 132 + 25);
          } else {
            delete x["StaticIPConfig"]["Type"]["DeviceEditable"];
          }
        } else {
          delete x["StaticIPConfig"]["Type"];
        }
        if (A.load.Bool(ptr + 736 + 164 + 28)) {
          x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"] = {};
          x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["Active"] = A.load.Ref(ptr + 736 + 164 + 0, undefined);
          x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["Effective"] = A.load.Ref(ptr + 736 + 164 + 4, undefined);
          x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["UserPolicy"] = A.load.Ref(ptr + 736 + 164 + 8, undefined);
          x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["DevicePolicy"] = A.load.Ref(ptr + 736 + 164 + 12, undefined);
          x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["UserSetting"] = A.load.Ref(ptr + 736 + 164 + 16, undefined);
          x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["SharedSetting"] = A.load.Ref(
            ptr + 736 + 164 + 20,
            undefined
          );
          if (A.load.Bool(ptr + 736 + 164 + 26)) {
            x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["UserEditable"] = A.load.Bool(ptr + 736 + 164 + 24);
          } else {
            delete x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 736 + 164 + 27)) {
            x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["DeviceEditable"] = A.load.Bool(ptr + 736 + 164 + 25);
          } else {
            delete x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]["DeviceEditable"];
          }
        } else {
          delete x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"];
        }
      } else {
        delete x["StaticIPConfig"];
      }
      if (A.load.Bool(ptr + 932 + 37)) {
        x["SavedIPConfig"] = {};
        x["SavedIPConfig"]["Gateway"] = A.load.Ref(ptr + 932 + 0, undefined);
        x["SavedIPConfig"]["IPAddress"] = A.load.Ref(ptr + 932 + 4, undefined);
        x["SavedIPConfig"]["ExcludedRoutes"] = A.load.Ref(ptr + 932 + 8, undefined);
        x["SavedIPConfig"]["IncludedRoutes"] = A.load.Ref(ptr + 932 + 12, undefined);
        x["SavedIPConfig"]["NameServers"] = A.load.Ref(ptr + 932 + 16, undefined);
        x["SavedIPConfig"]["SearchDomains"] = A.load.Ref(ptr + 932 + 20, undefined);
        if (A.load.Bool(ptr + 932 + 36)) {
          x["SavedIPConfig"]["RoutingPrefix"] = A.load.Int32(ptr + 932 + 24);
        } else {
          delete x["SavedIPConfig"]["RoutingPrefix"];
        }
        x["SavedIPConfig"]["Type"] = A.load.Ref(ptr + 932 + 28, undefined);
        x["SavedIPConfig"]["WebProxyAutoDiscoveryUrl"] = A.load.Ref(ptr + 932 + 32, undefined);
      } else {
        delete x["SavedIPConfig"];
      }
      x["Source"] = A.load.Ref(ptr + 972, undefined);
      x["Type"] = A.load.Enum(ptr + 976, ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"]);
      if (A.load.Bool(ptr + 980 + 85)) {
        x["VPN"] = {};
        if (A.load.Bool(ptr + 980 + 0 + 21)) {
          x["VPN"]["AutoConnect"] = {};
          if (A.load.Bool(ptr + 980 + 0 + 14)) {
            x["VPN"]["AutoConnect"]["Active"] = A.load.Bool(ptr + 980 + 0 + 0);
          } else {
            delete x["VPN"]["AutoConnect"]["Active"];
          }
          x["VPN"]["AutoConnect"]["Effective"] = A.load.Ref(ptr + 980 + 0 + 4, undefined);
          if (A.load.Bool(ptr + 980 + 0 + 15)) {
            x["VPN"]["AutoConnect"]["UserPolicy"] = A.load.Bool(ptr + 980 + 0 + 8);
          } else {
            delete x["VPN"]["AutoConnect"]["UserPolicy"];
          }
          if (A.load.Bool(ptr + 980 + 0 + 16)) {
            x["VPN"]["AutoConnect"]["DevicePolicy"] = A.load.Bool(ptr + 980 + 0 + 9);
          } else {
            delete x["VPN"]["AutoConnect"]["DevicePolicy"];
          }
          if (A.load.Bool(ptr + 980 + 0 + 17)) {
            x["VPN"]["AutoConnect"]["UserSetting"] = A.load.Bool(ptr + 980 + 0 + 10);
          } else {
            delete x["VPN"]["AutoConnect"]["UserSetting"];
          }
          if (A.load.Bool(ptr + 980 + 0 + 18)) {
            x["VPN"]["AutoConnect"]["SharedSetting"] = A.load.Bool(ptr + 980 + 0 + 11);
          } else {
            delete x["VPN"]["AutoConnect"]["SharedSetting"];
          }
          if (A.load.Bool(ptr + 980 + 0 + 19)) {
            x["VPN"]["AutoConnect"]["UserEditable"] = A.load.Bool(ptr + 980 + 0 + 12);
          } else {
            delete x["VPN"]["AutoConnect"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 980 + 0 + 20)) {
            x["VPN"]["AutoConnect"]["DeviceEditable"] = A.load.Bool(ptr + 980 + 0 + 13);
          } else {
            delete x["VPN"]["AutoConnect"]["DeviceEditable"];
          }
        } else {
          delete x["VPN"]["AutoConnect"];
        }
        if (A.load.Bool(ptr + 980 + 24 + 28)) {
          x["VPN"]["Host"] = {};
          x["VPN"]["Host"]["Active"] = A.load.Ref(ptr + 980 + 24 + 0, undefined);
          x["VPN"]["Host"]["Effective"] = A.load.Ref(ptr + 980 + 24 + 4, undefined);
          x["VPN"]["Host"]["UserPolicy"] = A.load.Ref(ptr + 980 + 24 + 8, undefined);
          x["VPN"]["Host"]["DevicePolicy"] = A.load.Ref(ptr + 980 + 24 + 12, undefined);
          x["VPN"]["Host"]["UserSetting"] = A.load.Ref(ptr + 980 + 24 + 16, undefined);
          x["VPN"]["Host"]["SharedSetting"] = A.load.Ref(ptr + 980 + 24 + 20, undefined);
          if (A.load.Bool(ptr + 980 + 24 + 26)) {
            x["VPN"]["Host"]["UserEditable"] = A.load.Bool(ptr + 980 + 24 + 24);
          } else {
            delete x["VPN"]["Host"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 980 + 24 + 27)) {
            x["VPN"]["Host"]["DeviceEditable"] = A.load.Bool(ptr + 980 + 24 + 25);
          } else {
            delete x["VPN"]["Host"]["DeviceEditable"];
          }
        } else {
          delete x["VPN"]["Host"];
        }
        if (A.load.Bool(ptr + 980 + 56 + 28)) {
          x["VPN"]["Type"] = {};
          x["VPN"]["Type"]["Active"] = A.load.Ref(ptr + 980 + 56 + 0, undefined);
          x["VPN"]["Type"]["Effective"] = A.load.Ref(ptr + 980 + 56 + 4, undefined);
          x["VPN"]["Type"]["UserPolicy"] = A.load.Ref(ptr + 980 + 56 + 8, undefined);
          x["VPN"]["Type"]["DevicePolicy"] = A.load.Ref(ptr + 980 + 56 + 12, undefined);
          x["VPN"]["Type"]["UserSetting"] = A.load.Ref(ptr + 980 + 56 + 16, undefined);
          x["VPN"]["Type"]["SharedSetting"] = A.load.Ref(ptr + 980 + 56 + 20, undefined);
          if (A.load.Bool(ptr + 980 + 56 + 26)) {
            x["VPN"]["Type"]["UserEditable"] = A.load.Bool(ptr + 980 + 56 + 24);
          } else {
            delete x["VPN"]["Type"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 980 + 56 + 27)) {
            x["VPN"]["Type"]["DeviceEditable"] = A.load.Bool(ptr + 980 + 56 + 25);
          } else {
            delete x["VPN"]["Type"]["DeviceEditable"];
          }
        } else {
          delete x["VPN"]["Type"];
        }
      } else {
        delete x["VPN"];
      }
      if (A.load.Bool(ptr + 1068 + 222)) {
        x["WiFi"] = {};
        if (A.load.Bool(ptr + 1068 + 0 + 21)) {
          x["WiFi"]["AllowGatewayARPPolling"] = {};
          if (A.load.Bool(ptr + 1068 + 0 + 14)) {
            x["WiFi"]["AllowGatewayARPPolling"]["Active"] = A.load.Bool(ptr + 1068 + 0 + 0);
          } else {
            delete x["WiFi"]["AllowGatewayARPPolling"]["Active"];
          }
          x["WiFi"]["AllowGatewayARPPolling"]["Effective"] = A.load.Ref(ptr + 1068 + 0 + 4, undefined);
          if (A.load.Bool(ptr + 1068 + 0 + 15)) {
            x["WiFi"]["AllowGatewayARPPolling"]["UserPolicy"] = A.load.Bool(ptr + 1068 + 0 + 8);
          } else {
            delete x["WiFi"]["AllowGatewayARPPolling"]["UserPolicy"];
          }
          if (A.load.Bool(ptr + 1068 + 0 + 16)) {
            x["WiFi"]["AllowGatewayARPPolling"]["DevicePolicy"] = A.load.Bool(ptr + 1068 + 0 + 9);
          } else {
            delete x["WiFi"]["AllowGatewayARPPolling"]["DevicePolicy"];
          }
          if (A.load.Bool(ptr + 1068 + 0 + 17)) {
            x["WiFi"]["AllowGatewayARPPolling"]["UserSetting"] = A.load.Bool(ptr + 1068 + 0 + 10);
          } else {
            delete x["WiFi"]["AllowGatewayARPPolling"]["UserSetting"];
          }
          if (A.load.Bool(ptr + 1068 + 0 + 18)) {
            x["WiFi"]["AllowGatewayARPPolling"]["SharedSetting"] = A.load.Bool(ptr + 1068 + 0 + 11);
          } else {
            delete x["WiFi"]["AllowGatewayARPPolling"]["SharedSetting"];
          }
          if (A.load.Bool(ptr + 1068 + 0 + 19)) {
            x["WiFi"]["AllowGatewayARPPolling"]["UserEditable"] = A.load.Bool(ptr + 1068 + 0 + 12);
          } else {
            delete x["WiFi"]["AllowGatewayARPPolling"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 1068 + 0 + 20)) {
            x["WiFi"]["AllowGatewayARPPolling"]["DeviceEditable"] = A.load.Bool(ptr + 1068 + 0 + 13);
          } else {
            delete x["WiFi"]["AllowGatewayARPPolling"]["DeviceEditable"];
          }
        } else {
          delete x["WiFi"]["AllowGatewayARPPolling"];
        }
        if (A.load.Bool(ptr + 1068 + 24 + 21)) {
          x["WiFi"]["AutoConnect"] = {};
          if (A.load.Bool(ptr + 1068 + 24 + 14)) {
            x["WiFi"]["AutoConnect"]["Active"] = A.load.Bool(ptr + 1068 + 24 + 0);
          } else {
            delete x["WiFi"]["AutoConnect"]["Active"];
          }
          x["WiFi"]["AutoConnect"]["Effective"] = A.load.Ref(ptr + 1068 + 24 + 4, undefined);
          if (A.load.Bool(ptr + 1068 + 24 + 15)) {
            x["WiFi"]["AutoConnect"]["UserPolicy"] = A.load.Bool(ptr + 1068 + 24 + 8);
          } else {
            delete x["WiFi"]["AutoConnect"]["UserPolicy"];
          }
          if (A.load.Bool(ptr + 1068 + 24 + 16)) {
            x["WiFi"]["AutoConnect"]["DevicePolicy"] = A.load.Bool(ptr + 1068 + 24 + 9);
          } else {
            delete x["WiFi"]["AutoConnect"]["DevicePolicy"];
          }
          if (A.load.Bool(ptr + 1068 + 24 + 17)) {
            x["WiFi"]["AutoConnect"]["UserSetting"] = A.load.Bool(ptr + 1068 + 24 + 10);
          } else {
            delete x["WiFi"]["AutoConnect"]["UserSetting"];
          }
          if (A.load.Bool(ptr + 1068 + 24 + 18)) {
            x["WiFi"]["AutoConnect"]["SharedSetting"] = A.load.Bool(ptr + 1068 + 24 + 11);
          } else {
            delete x["WiFi"]["AutoConnect"]["SharedSetting"];
          }
          if (A.load.Bool(ptr + 1068 + 24 + 19)) {
            x["WiFi"]["AutoConnect"]["UserEditable"] = A.load.Bool(ptr + 1068 + 24 + 12);
          } else {
            delete x["WiFi"]["AutoConnect"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 1068 + 24 + 20)) {
            x["WiFi"]["AutoConnect"]["DeviceEditable"] = A.load.Bool(ptr + 1068 + 24 + 13);
          } else {
            delete x["WiFi"]["AutoConnect"]["DeviceEditable"];
          }
        } else {
          delete x["WiFi"]["AutoConnect"];
        }
        x["WiFi"]["BSSID"] = A.load.Ref(ptr + 1068 + 48, undefined);
        if (A.load.Bool(ptr + 1068 + 220)) {
          x["WiFi"]["Frequency"] = A.load.Int32(ptr + 1068 + 52);
        } else {
          delete x["WiFi"]["Frequency"];
        }
        x["WiFi"]["FrequencyList"] = A.load.Ref(ptr + 1068 + 56, undefined);
        if (A.load.Bool(ptr + 1068 + 60 + 28)) {
          x["WiFi"]["HexSSID"] = {};
          x["WiFi"]["HexSSID"]["Active"] = A.load.Ref(ptr + 1068 + 60 + 0, undefined);
          x["WiFi"]["HexSSID"]["Effective"] = A.load.Ref(ptr + 1068 + 60 + 4, undefined);
          x["WiFi"]["HexSSID"]["UserPolicy"] = A.load.Ref(ptr + 1068 + 60 + 8, undefined);
          x["WiFi"]["HexSSID"]["DevicePolicy"] = A.load.Ref(ptr + 1068 + 60 + 12, undefined);
          x["WiFi"]["HexSSID"]["UserSetting"] = A.load.Ref(ptr + 1068 + 60 + 16, undefined);
          x["WiFi"]["HexSSID"]["SharedSetting"] = A.load.Ref(ptr + 1068 + 60 + 20, undefined);
          if (A.load.Bool(ptr + 1068 + 60 + 26)) {
            x["WiFi"]["HexSSID"]["UserEditable"] = A.load.Bool(ptr + 1068 + 60 + 24);
          } else {
            delete x["WiFi"]["HexSSID"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 1068 + 60 + 27)) {
            x["WiFi"]["HexSSID"]["DeviceEditable"] = A.load.Bool(ptr + 1068 + 60 + 25);
          } else {
            delete x["WiFi"]["HexSSID"]["DeviceEditable"];
          }
        } else {
          delete x["WiFi"]["HexSSID"];
        }
        if (A.load.Bool(ptr + 1068 + 92 + 21)) {
          x["WiFi"]["HiddenSSID"] = {};
          if (A.load.Bool(ptr + 1068 + 92 + 14)) {
            x["WiFi"]["HiddenSSID"]["Active"] = A.load.Bool(ptr + 1068 + 92 + 0);
          } else {
            delete x["WiFi"]["HiddenSSID"]["Active"];
          }
          x["WiFi"]["HiddenSSID"]["Effective"] = A.load.Ref(ptr + 1068 + 92 + 4, undefined);
          if (A.load.Bool(ptr + 1068 + 92 + 15)) {
            x["WiFi"]["HiddenSSID"]["UserPolicy"] = A.load.Bool(ptr + 1068 + 92 + 8);
          } else {
            delete x["WiFi"]["HiddenSSID"]["UserPolicy"];
          }
          if (A.load.Bool(ptr + 1068 + 92 + 16)) {
            x["WiFi"]["HiddenSSID"]["DevicePolicy"] = A.load.Bool(ptr + 1068 + 92 + 9);
          } else {
            delete x["WiFi"]["HiddenSSID"]["DevicePolicy"];
          }
          if (A.load.Bool(ptr + 1068 + 92 + 17)) {
            x["WiFi"]["HiddenSSID"]["UserSetting"] = A.load.Bool(ptr + 1068 + 92 + 10);
          } else {
            delete x["WiFi"]["HiddenSSID"]["UserSetting"];
          }
          if (A.load.Bool(ptr + 1068 + 92 + 18)) {
            x["WiFi"]["HiddenSSID"]["SharedSetting"] = A.load.Bool(ptr + 1068 + 92 + 11);
          } else {
            delete x["WiFi"]["HiddenSSID"]["SharedSetting"];
          }
          if (A.load.Bool(ptr + 1068 + 92 + 19)) {
            x["WiFi"]["HiddenSSID"]["UserEditable"] = A.load.Bool(ptr + 1068 + 92 + 12);
          } else {
            delete x["WiFi"]["HiddenSSID"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 1068 + 92 + 20)) {
            x["WiFi"]["HiddenSSID"]["DeviceEditable"] = A.load.Bool(ptr + 1068 + 92 + 13);
          } else {
            delete x["WiFi"]["HiddenSSID"]["DeviceEditable"];
          }
        } else {
          delete x["WiFi"]["HiddenSSID"];
        }
        if (A.load.Bool(ptr + 1068 + 116 + 33)) {
          x["WiFi"]["RoamThreshold"] = {};
          if (A.load.Bool(ptr + 1068 + 116 + 26)) {
            x["WiFi"]["RoamThreshold"]["Active"] = A.load.Int32(ptr + 1068 + 116 + 0);
          } else {
            delete x["WiFi"]["RoamThreshold"]["Active"];
          }
          x["WiFi"]["RoamThreshold"]["Effective"] = A.load.Ref(ptr + 1068 + 116 + 4, undefined);
          if (A.load.Bool(ptr + 1068 + 116 + 27)) {
            x["WiFi"]["RoamThreshold"]["UserPolicy"] = A.load.Int32(ptr + 1068 + 116 + 8);
          } else {
            delete x["WiFi"]["RoamThreshold"]["UserPolicy"];
          }
          if (A.load.Bool(ptr + 1068 + 116 + 28)) {
            x["WiFi"]["RoamThreshold"]["DevicePolicy"] = A.load.Int32(ptr + 1068 + 116 + 12);
          } else {
            delete x["WiFi"]["RoamThreshold"]["DevicePolicy"];
          }
          if (A.load.Bool(ptr + 1068 + 116 + 29)) {
            x["WiFi"]["RoamThreshold"]["UserSetting"] = A.load.Int32(ptr + 1068 + 116 + 16);
          } else {
            delete x["WiFi"]["RoamThreshold"]["UserSetting"];
          }
          if (A.load.Bool(ptr + 1068 + 116 + 30)) {
            x["WiFi"]["RoamThreshold"]["SharedSetting"] = A.load.Int32(ptr + 1068 + 116 + 20);
          } else {
            delete x["WiFi"]["RoamThreshold"]["SharedSetting"];
          }
          if (A.load.Bool(ptr + 1068 + 116 + 31)) {
            x["WiFi"]["RoamThreshold"]["UserEditable"] = A.load.Bool(ptr + 1068 + 116 + 24);
          } else {
            delete x["WiFi"]["RoamThreshold"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 1068 + 116 + 32)) {
            x["WiFi"]["RoamThreshold"]["DeviceEditable"] = A.load.Bool(ptr + 1068 + 116 + 25);
          } else {
            delete x["WiFi"]["RoamThreshold"]["DeviceEditable"];
          }
        } else {
          delete x["WiFi"]["RoamThreshold"];
        }
        if (A.load.Bool(ptr + 1068 + 152 + 28)) {
          x["WiFi"]["SSID"] = {};
          x["WiFi"]["SSID"]["Active"] = A.load.Ref(ptr + 1068 + 152 + 0, undefined);
          x["WiFi"]["SSID"]["Effective"] = A.load.Ref(ptr + 1068 + 152 + 4, undefined);
          x["WiFi"]["SSID"]["UserPolicy"] = A.load.Ref(ptr + 1068 + 152 + 8, undefined);
          x["WiFi"]["SSID"]["DevicePolicy"] = A.load.Ref(ptr + 1068 + 152 + 12, undefined);
          x["WiFi"]["SSID"]["UserSetting"] = A.load.Ref(ptr + 1068 + 152 + 16, undefined);
          x["WiFi"]["SSID"]["SharedSetting"] = A.load.Ref(ptr + 1068 + 152 + 20, undefined);
          if (A.load.Bool(ptr + 1068 + 152 + 26)) {
            x["WiFi"]["SSID"]["UserEditable"] = A.load.Bool(ptr + 1068 + 152 + 24);
          } else {
            delete x["WiFi"]["SSID"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 1068 + 152 + 27)) {
            x["WiFi"]["SSID"]["DeviceEditable"] = A.load.Bool(ptr + 1068 + 152 + 25);
          } else {
            delete x["WiFi"]["SSID"]["DeviceEditable"];
          }
        } else {
          delete x["WiFi"]["SSID"];
        }
        if (A.load.Bool(ptr + 1068 + 184 + 28)) {
          x["WiFi"]["Security"] = {};
          x["WiFi"]["Security"]["Active"] = A.load.Ref(ptr + 1068 + 184 + 0, undefined);
          x["WiFi"]["Security"]["Effective"] = A.load.Ref(ptr + 1068 + 184 + 4, undefined);
          x["WiFi"]["Security"]["UserPolicy"] = A.load.Ref(ptr + 1068 + 184 + 8, undefined);
          x["WiFi"]["Security"]["DevicePolicy"] = A.load.Ref(ptr + 1068 + 184 + 12, undefined);
          x["WiFi"]["Security"]["UserSetting"] = A.load.Ref(ptr + 1068 + 184 + 16, undefined);
          x["WiFi"]["Security"]["SharedSetting"] = A.load.Ref(ptr + 1068 + 184 + 20, undefined);
          if (A.load.Bool(ptr + 1068 + 184 + 26)) {
            x["WiFi"]["Security"]["UserEditable"] = A.load.Bool(ptr + 1068 + 184 + 24);
          } else {
            delete x["WiFi"]["Security"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 1068 + 184 + 27)) {
            x["WiFi"]["Security"]["DeviceEditable"] = A.load.Bool(ptr + 1068 + 184 + 25);
          } else {
            delete x["WiFi"]["Security"]["DeviceEditable"];
          }
        } else {
          delete x["WiFi"]["Security"];
        }
        if (A.load.Bool(ptr + 1068 + 221)) {
          x["WiFi"]["SignalStrength"] = A.load.Int32(ptr + 1068 + 216);
        } else {
          delete x["WiFi"]["SignalStrength"];
        }
      } else {
        delete x["WiFi"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_VPNStateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["Type"]);
      }
    },
    "load_VPNStateProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["Type"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_WiFiStateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 26, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 24, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 25, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Ref(ptr + 20, undefined);
      } else {
        A.store.Bool(ptr + 26, true);
        A.store.Ref(ptr + 0, x["BSSID"]);
        A.store.Bool(ptr + 24, "Frequency" in x ? true : false);
        A.store.Int32(ptr + 4, x["Frequency"] === undefined ? 0 : (x["Frequency"] as number));
        A.store.Ref(ptr + 8, x["HexSSID"]);
        A.store.Ref(ptr + 12, x["Security"]);
        A.store.Bool(ptr + 25, "SignalStrength" in x ? true : false);
        A.store.Int32(ptr + 16, x["SignalStrength"] === undefined ? 0 : (x["SignalStrength"] as number));
        A.store.Ref(ptr + 20, x["SSID"]);
      }
    },
    "load_WiFiStateProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["BSSID"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 24)) {
        x["Frequency"] = A.load.Int32(ptr + 4);
      } else {
        delete x["Frequency"];
      }
      x["HexSSID"] = A.load.Ref(ptr + 8, undefined);
      x["Security"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 25)) {
        x["SignalStrength"] = A.load.Int32(ptr + 16);
      } else {
        delete x["SignalStrength"];
      }
      x["SSID"] = A.load.Ref(ptr + 20, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_NetworkStateProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 97, false);

        A.store.Bool(ptr + 0 + 22, false);
        A.store.Enum(ptr + 0 + 0, -1);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Ref(ptr + 0 + 8, undefined);
        A.store.Bool(ptr + 0 + 20, false);
        A.store.Bool(ptr + 0 + 12, false);
        A.store.Bool(ptr + 0 + 21, false);
        A.store.Int32(ptr + 0 + 16, 0);
        A.store.Bool(ptr + 95, false);
        A.store.Bool(ptr + 23, false);
        A.store.Enum(ptr + 24, -1);

        A.store.Bool(ptr + 28 + 4, false);
        A.store.Ref(ptr + 28 + 0, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Bool(ptr + 96, false);
        A.store.Int32(ptr + 48, 0);
        A.store.Ref(ptr + 52, undefined);
        A.store.Enum(ptr + 56, -1);

        A.store.Bool(ptr + 60 + 4, false);
        A.store.Ref(ptr + 60 + 0, undefined);

        A.store.Bool(ptr + 68 + 26, false);
        A.store.Ref(ptr + 68 + 0, undefined);
        A.store.Bool(ptr + 68 + 24, false);
        A.store.Int32(ptr + 68 + 4, 0);
        A.store.Ref(ptr + 68 + 8, undefined);
        A.store.Ref(ptr + 68 + 12, undefined);
        A.store.Bool(ptr + 68 + 25, false);
        A.store.Int32(ptr + 68 + 16, 0);
        A.store.Ref(ptr + 68 + 20, undefined);
      } else {
        A.store.Bool(ptr + 97, true);

        if (typeof x["Cellular"] === "undefined") {
          A.store.Bool(ptr + 0 + 22, false);
          A.store.Enum(ptr + 0 + 0, -1);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Ref(ptr + 0 + 8, undefined);
          A.store.Bool(ptr + 0 + 20, false);
          A.store.Bool(ptr + 0 + 12, false);
          A.store.Bool(ptr + 0 + 21, false);
          A.store.Int32(ptr + 0 + 16, 0);
        } else {
          A.store.Bool(ptr + 0 + 22, true);
          A.store.Enum(
            ptr + 0 + 0,
            ["Activated", "Activating", "NotActivated", "PartiallyActivated"].indexOf(
              x["Cellular"]["ActivationState"] as string
            )
          );
          A.store.Ref(ptr + 0 + 4, x["Cellular"]["NetworkTechnology"]);
          A.store.Ref(ptr + 0 + 8, x["Cellular"]["RoamingState"]);
          A.store.Bool(ptr + 0 + 20, "SIMPresent" in x["Cellular"] ? true : false);
          A.store.Bool(ptr + 0 + 12, x["Cellular"]["SIMPresent"] ? true : false);
          A.store.Bool(ptr + 0 + 21, "SignalStrength" in x["Cellular"] ? true : false);
          A.store.Int32(
            ptr + 0 + 16,
            x["Cellular"]["SignalStrength"] === undefined ? 0 : (x["Cellular"]["SignalStrength"] as number)
          );
        }
        A.store.Bool(ptr + 95, "Connectable" in x ? true : false);
        A.store.Bool(ptr + 23, x["Connectable"] ? true : false);
        A.store.Enum(ptr + 24, ["Connected", "Connecting", "NotConnected"].indexOf(x["ConnectionState"] as string));

        if (typeof x["Ethernet"] === "undefined") {
          A.store.Bool(ptr + 28 + 4, false);
          A.store.Ref(ptr + 28 + 0, undefined);
        } else {
          A.store.Bool(ptr + 28 + 4, true);
          A.store.Ref(ptr + 28 + 0, x["Ethernet"]["Authentication"]);
        }
        A.store.Ref(ptr + 36, x["ErrorState"]);
        A.store.Ref(ptr + 40, x["GUID"]);
        A.store.Ref(ptr + 44, x["Name"]);
        A.store.Bool(ptr + 96, "Priority" in x ? true : false);
        A.store.Int32(ptr + 48, x["Priority"] === undefined ? 0 : (x["Priority"] as number));
        A.store.Ref(ptr + 52, x["Source"]);
        A.store.Enum(
          ptr + 56,
          ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"].indexOf(x["Type"] as string)
        );

        if (typeof x["VPN"] === "undefined") {
          A.store.Bool(ptr + 60 + 4, false);
          A.store.Ref(ptr + 60 + 0, undefined);
        } else {
          A.store.Bool(ptr + 60 + 4, true);
          A.store.Ref(ptr + 60 + 0, x["VPN"]["Type"]);
        }

        if (typeof x["WiFi"] === "undefined") {
          A.store.Bool(ptr + 68 + 26, false);
          A.store.Ref(ptr + 68 + 0, undefined);
          A.store.Bool(ptr + 68 + 24, false);
          A.store.Int32(ptr + 68 + 4, 0);
          A.store.Ref(ptr + 68 + 8, undefined);
          A.store.Ref(ptr + 68 + 12, undefined);
          A.store.Bool(ptr + 68 + 25, false);
          A.store.Int32(ptr + 68 + 16, 0);
          A.store.Ref(ptr + 68 + 20, undefined);
        } else {
          A.store.Bool(ptr + 68 + 26, true);
          A.store.Ref(ptr + 68 + 0, x["WiFi"]["BSSID"]);
          A.store.Bool(ptr + 68 + 24, "Frequency" in x["WiFi"] ? true : false);
          A.store.Int32(ptr + 68 + 4, x["WiFi"]["Frequency"] === undefined ? 0 : (x["WiFi"]["Frequency"] as number));
          A.store.Ref(ptr + 68 + 8, x["WiFi"]["HexSSID"]);
          A.store.Ref(ptr + 68 + 12, x["WiFi"]["Security"]);
          A.store.Bool(ptr + 68 + 25, "SignalStrength" in x["WiFi"] ? true : false);
          A.store.Int32(
            ptr + 68 + 16,
            x["WiFi"]["SignalStrength"] === undefined ? 0 : (x["WiFi"]["SignalStrength"] as number)
          );
          A.store.Ref(ptr + 68 + 20, x["WiFi"]["SSID"]);
        }
      }
    },
    "load_NetworkStateProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 22)) {
        x["Cellular"] = {};
        x["Cellular"]["ActivationState"] = A.load.Enum(ptr + 0 + 0, [
          "Activated",
          "Activating",
          "NotActivated",
          "PartiallyActivated",
        ]);
        x["Cellular"]["NetworkTechnology"] = A.load.Ref(ptr + 0 + 4, undefined);
        x["Cellular"]["RoamingState"] = A.load.Ref(ptr + 0 + 8, undefined);
        if (A.load.Bool(ptr + 0 + 20)) {
          x["Cellular"]["SIMPresent"] = A.load.Bool(ptr + 0 + 12);
        } else {
          delete x["Cellular"]["SIMPresent"];
        }
        if (A.load.Bool(ptr + 0 + 21)) {
          x["Cellular"]["SignalStrength"] = A.load.Int32(ptr + 0 + 16);
        } else {
          delete x["Cellular"]["SignalStrength"];
        }
      } else {
        delete x["Cellular"];
      }
      if (A.load.Bool(ptr + 95)) {
        x["Connectable"] = A.load.Bool(ptr + 23);
      } else {
        delete x["Connectable"];
      }
      x["ConnectionState"] = A.load.Enum(ptr + 24, ["Connected", "Connecting", "NotConnected"]);
      if (A.load.Bool(ptr + 28 + 4)) {
        x["Ethernet"] = {};
        x["Ethernet"]["Authentication"] = A.load.Ref(ptr + 28 + 0, undefined);
      } else {
        delete x["Ethernet"];
      }
      x["ErrorState"] = A.load.Ref(ptr + 36, undefined);
      x["GUID"] = A.load.Ref(ptr + 40, undefined);
      x["Name"] = A.load.Ref(ptr + 44, undefined);
      if (A.load.Bool(ptr + 96)) {
        x["Priority"] = A.load.Int32(ptr + 48);
      } else {
        delete x["Priority"];
      }
      x["Source"] = A.load.Ref(ptr + 52, undefined);
      x["Type"] = A.load.Enum(ptr + 56, ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"]);
      if (A.load.Bool(ptr + 60 + 4)) {
        x["VPN"] = {};
        x["VPN"]["Type"] = A.load.Ref(ptr + 60 + 0, undefined);
      } else {
        delete x["VPN"];
      }
      if (A.load.Bool(ptr + 68 + 26)) {
        x["WiFi"] = {};
        x["WiFi"]["BSSID"] = A.load.Ref(ptr + 68 + 0, undefined);
        if (A.load.Bool(ptr + 68 + 24)) {
          x["WiFi"]["Frequency"] = A.load.Int32(ptr + 68 + 4);
        } else {
          delete x["WiFi"]["Frequency"];
        }
        x["WiFi"]["HexSSID"] = A.load.Ref(ptr + 68 + 8, undefined);
        x["WiFi"]["Security"] = A.load.Ref(ptr + 68 + 12, undefined);
        if (A.load.Bool(ptr + 68 + 25)) {
          x["WiFi"]["SignalStrength"] = A.load.Int32(ptr + 68 + 16);
        } else {
          delete x["WiFi"]["SignalStrength"];
        }
        x["WiFi"]["SSID"] = A.load.Ref(ptr + 68 + 20, undefined);
      } else {
        delete x["WiFi"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ProxyLocation": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Ref(ptr + 0, x["Host"]);
        A.store.Bool(ptr + 8, "Port" in x ? true : false);
        A.store.Int32(ptr + 4, x["Port"] === undefined ? 0 : (x["Port"] as number));
      }
    },
    "load_ProxyLocation": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["Host"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 8)) {
        x["Port"] = A.load.Int32(ptr + 4);
      } else {
        delete x["Port"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManualProxySettings": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 46, false);

        A.store.Bool(ptr + 0 + 9, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Bool(ptr + 0 + 8, false);
        A.store.Int32(ptr + 0 + 4, 0);

        A.store.Bool(ptr + 12 + 9, false);
        A.store.Ref(ptr + 12 + 0, undefined);
        A.store.Bool(ptr + 12 + 8, false);
        A.store.Int32(ptr + 12 + 4, 0);

        A.store.Bool(ptr + 24 + 9, false);
        A.store.Ref(ptr + 24 + 0, undefined);
        A.store.Bool(ptr + 24 + 8, false);
        A.store.Int32(ptr + 24 + 4, 0);

        A.store.Bool(ptr + 36 + 9, false);
        A.store.Ref(ptr + 36 + 0, undefined);
        A.store.Bool(ptr + 36 + 8, false);
        A.store.Int32(ptr + 36 + 4, 0);
      } else {
        A.store.Bool(ptr + 46, true);

        if (typeof x["HTTPProxy"] === "undefined") {
          A.store.Bool(ptr + 0 + 9, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Int32(ptr + 0 + 4, 0);
        } else {
          A.store.Bool(ptr + 0 + 9, true);
          A.store.Ref(ptr + 0 + 0, x["HTTPProxy"]["Host"]);
          A.store.Bool(ptr + 0 + 8, "Port" in x["HTTPProxy"] ? true : false);
          A.store.Int32(ptr + 0 + 4, x["HTTPProxy"]["Port"] === undefined ? 0 : (x["HTTPProxy"]["Port"] as number));
        }

        if (typeof x["SecureHTTPProxy"] === "undefined") {
          A.store.Bool(ptr + 12 + 9, false);
          A.store.Ref(ptr + 12 + 0, undefined);
          A.store.Bool(ptr + 12 + 8, false);
          A.store.Int32(ptr + 12 + 4, 0);
        } else {
          A.store.Bool(ptr + 12 + 9, true);
          A.store.Ref(ptr + 12 + 0, x["SecureHTTPProxy"]["Host"]);
          A.store.Bool(ptr + 12 + 8, "Port" in x["SecureHTTPProxy"] ? true : false);
          A.store.Int32(
            ptr + 12 + 4,
            x["SecureHTTPProxy"]["Port"] === undefined ? 0 : (x["SecureHTTPProxy"]["Port"] as number)
          );
        }

        if (typeof x["FTPProxy"] === "undefined") {
          A.store.Bool(ptr + 24 + 9, false);
          A.store.Ref(ptr + 24 + 0, undefined);
          A.store.Bool(ptr + 24 + 8, false);
          A.store.Int32(ptr + 24 + 4, 0);
        } else {
          A.store.Bool(ptr + 24 + 9, true);
          A.store.Ref(ptr + 24 + 0, x["FTPProxy"]["Host"]);
          A.store.Bool(ptr + 24 + 8, "Port" in x["FTPProxy"] ? true : false);
          A.store.Int32(ptr + 24 + 4, x["FTPProxy"]["Port"] === undefined ? 0 : (x["FTPProxy"]["Port"] as number));
        }

        if (typeof x["SOCKS"] === "undefined") {
          A.store.Bool(ptr + 36 + 9, false);
          A.store.Ref(ptr + 36 + 0, undefined);
          A.store.Bool(ptr + 36 + 8, false);
          A.store.Int32(ptr + 36 + 4, 0);
        } else {
          A.store.Bool(ptr + 36 + 9, true);
          A.store.Ref(ptr + 36 + 0, x["SOCKS"]["Host"]);
          A.store.Bool(ptr + 36 + 8, "Port" in x["SOCKS"] ? true : false);
          A.store.Int32(ptr + 36 + 4, x["SOCKS"]["Port"] === undefined ? 0 : (x["SOCKS"]["Port"] as number));
        }
      }
    },
    "load_ManualProxySettings": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 9)) {
        x["HTTPProxy"] = {};
        x["HTTPProxy"]["Host"] = A.load.Ref(ptr + 0 + 0, undefined);
        if (A.load.Bool(ptr + 0 + 8)) {
          x["HTTPProxy"]["Port"] = A.load.Int32(ptr + 0 + 4);
        } else {
          delete x["HTTPProxy"]["Port"];
        }
      } else {
        delete x["HTTPProxy"];
      }
      if (A.load.Bool(ptr + 12 + 9)) {
        x["SecureHTTPProxy"] = {};
        x["SecureHTTPProxy"]["Host"] = A.load.Ref(ptr + 12 + 0, undefined);
        if (A.load.Bool(ptr + 12 + 8)) {
          x["SecureHTTPProxy"]["Port"] = A.load.Int32(ptr + 12 + 4);
        } else {
          delete x["SecureHTTPProxy"]["Port"];
        }
      } else {
        delete x["SecureHTTPProxy"];
      }
      if (A.load.Bool(ptr + 24 + 9)) {
        x["FTPProxy"] = {};
        x["FTPProxy"]["Host"] = A.load.Ref(ptr + 24 + 0, undefined);
        if (A.load.Bool(ptr + 24 + 8)) {
          x["FTPProxy"]["Port"] = A.load.Int32(ptr + 24 + 4);
        } else {
          delete x["FTPProxy"]["Port"];
        }
      } else {
        delete x["FTPProxy"];
      }
      if (A.load.Bool(ptr + 36 + 9)) {
        x["SOCKS"] = {};
        x["SOCKS"]["Host"] = A.load.Ref(ptr + 36 + 0, undefined);
        if (A.load.Bool(ptr + 36 + 8)) {
          x["SOCKS"]["Port"] = A.load.Int32(ptr + 36 + 4);
        } else {
          delete x["SOCKS"]["Port"];
        }
      } else {
        delete x["SOCKS"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ProxySettings": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 60, false);
        A.store.Enum(ptr + 0, -1);

        A.store.Bool(ptr + 4 + 46, false);

        A.store.Bool(ptr + 4 + 0 + 9, false);
        A.store.Ref(ptr + 4 + 0 + 0, undefined);
        A.store.Bool(ptr + 4 + 0 + 8, false);
        A.store.Int32(ptr + 4 + 0 + 4, 0);

        A.store.Bool(ptr + 4 + 12 + 9, false);
        A.store.Ref(ptr + 4 + 12 + 0, undefined);
        A.store.Bool(ptr + 4 + 12 + 8, false);
        A.store.Int32(ptr + 4 + 12 + 4, 0);

        A.store.Bool(ptr + 4 + 24 + 9, false);
        A.store.Ref(ptr + 4 + 24 + 0, undefined);
        A.store.Bool(ptr + 4 + 24 + 8, false);
        A.store.Int32(ptr + 4 + 24 + 4, 0);

        A.store.Bool(ptr + 4 + 36 + 9, false);
        A.store.Ref(ptr + 4 + 36 + 0, undefined);
        A.store.Bool(ptr + 4 + 36 + 8, false);
        A.store.Int32(ptr + 4 + 36 + 4, 0);
        A.store.Ref(ptr + 52, undefined);
        A.store.Ref(ptr + 56, undefined);
      } else {
        A.store.Bool(ptr + 60, true);
        A.store.Enum(ptr + 0, ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["Type"] as string));

        if (typeof x["Manual"] === "undefined") {
          A.store.Bool(ptr + 4 + 46, false);

          A.store.Bool(ptr + 4 + 0 + 9, false);
          A.store.Ref(ptr + 4 + 0 + 0, undefined);
          A.store.Bool(ptr + 4 + 0 + 8, false);
          A.store.Int32(ptr + 4 + 0 + 4, 0);

          A.store.Bool(ptr + 4 + 12 + 9, false);
          A.store.Ref(ptr + 4 + 12 + 0, undefined);
          A.store.Bool(ptr + 4 + 12 + 8, false);
          A.store.Int32(ptr + 4 + 12 + 4, 0);

          A.store.Bool(ptr + 4 + 24 + 9, false);
          A.store.Ref(ptr + 4 + 24 + 0, undefined);
          A.store.Bool(ptr + 4 + 24 + 8, false);
          A.store.Int32(ptr + 4 + 24 + 4, 0);

          A.store.Bool(ptr + 4 + 36 + 9, false);
          A.store.Ref(ptr + 4 + 36 + 0, undefined);
          A.store.Bool(ptr + 4 + 36 + 8, false);
          A.store.Int32(ptr + 4 + 36 + 4, 0);
        } else {
          A.store.Bool(ptr + 4 + 46, true);

          if (typeof x["Manual"]["HTTPProxy"] === "undefined") {
            A.store.Bool(ptr + 4 + 0 + 9, false);
            A.store.Ref(ptr + 4 + 0 + 0, undefined);
            A.store.Bool(ptr + 4 + 0 + 8, false);
            A.store.Int32(ptr + 4 + 0 + 4, 0);
          } else {
            A.store.Bool(ptr + 4 + 0 + 9, true);
            A.store.Ref(ptr + 4 + 0 + 0, x["Manual"]["HTTPProxy"]["Host"]);
            A.store.Bool(ptr + 4 + 0 + 8, "Port" in x["Manual"]["HTTPProxy"] ? true : false);
            A.store.Int32(
              ptr + 4 + 0 + 4,
              x["Manual"]["HTTPProxy"]["Port"] === undefined ? 0 : (x["Manual"]["HTTPProxy"]["Port"] as number)
            );
          }

          if (typeof x["Manual"]["SecureHTTPProxy"] === "undefined") {
            A.store.Bool(ptr + 4 + 12 + 9, false);
            A.store.Ref(ptr + 4 + 12 + 0, undefined);
            A.store.Bool(ptr + 4 + 12 + 8, false);
            A.store.Int32(ptr + 4 + 12 + 4, 0);
          } else {
            A.store.Bool(ptr + 4 + 12 + 9, true);
            A.store.Ref(ptr + 4 + 12 + 0, x["Manual"]["SecureHTTPProxy"]["Host"]);
            A.store.Bool(ptr + 4 + 12 + 8, "Port" in x["Manual"]["SecureHTTPProxy"] ? true : false);
            A.store.Int32(
              ptr + 4 + 12 + 4,
              x["Manual"]["SecureHTTPProxy"]["Port"] === undefined
                ? 0
                : (x["Manual"]["SecureHTTPProxy"]["Port"] as number)
            );
          }

          if (typeof x["Manual"]["FTPProxy"] === "undefined") {
            A.store.Bool(ptr + 4 + 24 + 9, false);
            A.store.Ref(ptr + 4 + 24 + 0, undefined);
            A.store.Bool(ptr + 4 + 24 + 8, false);
            A.store.Int32(ptr + 4 + 24 + 4, 0);
          } else {
            A.store.Bool(ptr + 4 + 24 + 9, true);
            A.store.Ref(ptr + 4 + 24 + 0, x["Manual"]["FTPProxy"]["Host"]);
            A.store.Bool(ptr + 4 + 24 + 8, "Port" in x["Manual"]["FTPProxy"] ? true : false);
            A.store.Int32(
              ptr + 4 + 24 + 4,
              x["Manual"]["FTPProxy"]["Port"] === undefined ? 0 : (x["Manual"]["FTPProxy"]["Port"] as number)
            );
          }

          if (typeof x["Manual"]["SOCKS"] === "undefined") {
            A.store.Bool(ptr + 4 + 36 + 9, false);
            A.store.Ref(ptr + 4 + 36 + 0, undefined);
            A.store.Bool(ptr + 4 + 36 + 8, false);
            A.store.Int32(ptr + 4 + 36 + 4, 0);
          } else {
            A.store.Bool(ptr + 4 + 36 + 9, true);
            A.store.Ref(ptr + 4 + 36 + 0, x["Manual"]["SOCKS"]["Host"]);
            A.store.Bool(ptr + 4 + 36 + 8, "Port" in x["Manual"]["SOCKS"] ? true : false);
            A.store.Int32(
              ptr + 4 + 36 + 4,
              x["Manual"]["SOCKS"]["Port"] === undefined ? 0 : (x["Manual"]["SOCKS"]["Port"] as number)
            );
          }
        }
        A.store.Ref(ptr + 52, x["ExcludeDomains"]);
        A.store.Ref(ptr + 56, x["PAC"]);
      }
    },
    "load_ProxySettings": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["Type"] = A.load.Enum(ptr + 0, ["Direct", "Manual", "PAC", "WPAD"]);
      if (A.load.Bool(ptr + 4 + 46)) {
        x["Manual"] = {};
        if (A.load.Bool(ptr + 4 + 0 + 9)) {
          x["Manual"]["HTTPProxy"] = {};
          x["Manual"]["HTTPProxy"]["Host"] = A.load.Ref(ptr + 4 + 0 + 0, undefined);
          if (A.load.Bool(ptr + 4 + 0 + 8)) {
            x["Manual"]["HTTPProxy"]["Port"] = A.load.Int32(ptr + 4 + 0 + 4);
          } else {
            delete x["Manual"]["HTTPProxy"]["Port"];
          }
        } else {
          delete x["Manual"]["HTTPProxy"];
        }
        if (A.load.Bool(ptr + 4 + 12 + 9)) {
          x["Manual"]["SecureHTTPProxy"] = {};
          x["Manual"]["SecureHTTPProxy"]["Host"] = A.load.Ref(ptr + 4 + 12 + 0, undefined);
          if (A.load.Bool(ptr + 4 + 12 + 8)) {
            x["Manual"]["SecureHTTPProxy"]["Port"] = A.load.Int32(ptr + 4 + 12 + 4);
          } else {
            delete x["Manual"]["SecureHTTPProxy"]["Port"];
          }
        } else {
          delete x["Manual"]["SecureHTTPProxy"];
        }
        if (A.load.Bool(ptr + 4 + 24 + 9)) {
          x["Manual"]["FTPProxy"] = {};
          x["Manual"]["FTPProxy"]["Host"] = A.load.Ref(ptr + 4 + 24 + 0, undefined);
          if (A.load.Bool(ptr + 4 + 24 + 8)) {
            x["Manual"]["FTPProxy"]["Port"] = A.load.Int32(ptr + 4 + 24 + 4);
          } else {
            delete x["Manual"]["FTPProxy"]["Port"];
          }
        } else {
          delete x["Manual"]["FTPProxy"];
        }
        if (A.load.Bool(ptr + 4 + 36 + 9)) {
          x["Manual"]["SOCKS"] = {};
          x["Manual"]["SOCKS"]["Host"] = A.load.Ref(ptr + 4 + 36 + 0, undefined);
          if (A.load.Bool(ptr + 4 + 36 + 8)) {
            x["Manual"]["SOCKS"]["Port"] = A.load.Int32(ptr + 4 + 36 + 4);
          } else {
            delete x["Manual"]["SOCKS"]["Port"];
          }
        } else {
          delete x["Manual"]["SOCKS"];
        }
      } else {
        delete x["Manual"];
      }
      x["ExcludeDomains"] = A.load.Ref(ptr + 52, undefined);
      x["PAC"] = A.load.Ref(ptr + 56, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_VPNProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "AutoConnect" in x ? true : false);
        A.store.Bool(ptr + 0, x["AutoConnect"] ? true : false);
        A.store.Ref(ptr + 4, x["Host"]);
        A.store.Ref(ptr + 8, x["Type"]);
      }
    },
    "load_VPNProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["AutoConnect"] = A.load.Bool(ptr + 0);
      } else {
        delete x["AutoConnect"];
      }
      x["Host"] = A.load.Ref(ptr + 4, undefined);
      x["Type"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_WiFiProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 182, false);
        A.store.Bool(ptr + 176, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 177, false);
        A.store.Bool(ptr + 1, false);
        A.store.Ref(ptr + 4, undefined);

        A.store.Bool(ptr + 8 + 130, false);
        A.store.Ref(ptr + 8 + 0, undefined);

        A.store.Bool(ptr + 8 + 4 + 45, false);
        A.store.Ref(ptr + 8 + 4 + 0, undefined);

        A.store.Bool(ptr + 8 + 4 + 4 + 16, false);
        A.store.Ref(ptr + 8 + 4 + 4 + 0, undefined);
        A.store.Ref(ptr + 8 + 4 + 4 + 4, undefined);
        A.store.Ref(ptr + 8 + 4 + 4 + 8, undefined);
        A.store.Ref(ptr + 8 + 4 + 4 + 12, undefined);
        A.store.Ref(ptr + 8 + 4 + 24, undefined);

        A.store.Bool(ptr + 8 + 4 + 28 + 16, false);
        A.store.Ref(ptr + 8 + 4 + 28 + 0, undefined);
        A.store.Ref(ptr + 8 + 4 + 28 + 4, undefined);
        A.store.Ref(ptr + 8 + 4 + 28 + 8, undefined);
        A.store.Ref(ptr + 8 + 4 + 28 + 12, undefined);
        A.store.Ref(ptr + 8 + 52, undefined);
        A.store.Ref(ptr + 8 + 56, undefined);
        A.store.Ref(ptr + 8 + 60, undefined);
        A.store.Enum(ptr + 8 + 64, -1);
        A.store.Ref(ptr + 8 + 68, undefined);
        A.store.Ref(ptr + 8 + 72, undefined);
        A.store.Ref(ptr + 8 + 76, undefined);
        A.store.Ref(ptr + 8 + 80, undefined);
        A.store.Bool(ptr + 8 + 127, false);
        A.store.Bool(ptr + 8 + 84, false);
        A.store.Ref(ptr + 8 + 88, undefined);
        A.store.Ref(ptr + 8 + 92, undefined);

        A.store.Bool(ptr + 8 + 96 + 28, false);
        A.store.Ref(ptr + 8 + 96 + 0, undefined);
        A.store.Ref(ptr + 8 + 96 + 4, undefined);
        A.store.Ref(ptr + 8 + 96 + 8, undefined);
        A.store.Ref(ptr + 8 + 96 + 12, undefined);
        A.store.Ref(ptr + 8 + 96 + 16, undefined);
        A.store.Ref(ptr + 8 + 96 + 20, undefined);
        A.store.Bool(ptr + 8 + 96 + 26, false);
        A.store.Bool(ptr + 8 + 96 + 24, false);
        A.store.Bool(ptr + 8 + 96 + 27, false);
        A.store.Bool(ptr + 8 + 96 + 25, false);
        A.store.Bool(ptr + 8 + 128, false);
        A.store.Bool(ptr + 8 + 125, false);
        A.store.Bool(ptr + 8 + 129, false);
        A.store.Bool(ptr + 8 + 126, false);
        A.store.Bool(ptr + 178, false);
        A.store.Int32(ptr + 140, 0);
        A.store.Ref(ptr + 144, undefined);
        A.store.Ref(ptr + 148, undefined);
        A.store.Bool(ptr + 179, false);
        A.store.Bool(ptr + 152, false);
        A.store.Ref(ptr + 156, undefined);
        A.store.Bool(ptr + 180, false);
        A.store.Int32(ptr + 160, 0);
        A.store.Ref(ptr + 164, undefined);
        A.store.Ref(ptr + 168, undefined);
        A.store.Bool(ptr + 181, false);
        A.store.Int32(ptr + 172, 0);
      } else {
        A.store.Bool(ptr + 182, true);
        A.store.Bool(ptr + 176, "AllowGatewayARPPolling" in x ? true : false);
        A.store.Bool(ptr + 0, x["AllowGatewayARPPolling"] ? true : false);
        A.store.Bool(ptr + 177, "AutoConnect" in x ? true : false);
        A.store.Bool(ptr + 1, x["AutoConnect"] ? true : false);
        A.store.Ref(ptr + 4, x["BSSID"]);

        if (typeof x["EAP"] === "undefined") {
          A.store.Bool(ptr + 8 + 130, false);
          A.store.Ref(ptr + 8 + 0, undefined);

          A.store.Bool(ptr + 8 + 4 + 45, false);
          A.store.Ref(ptr + 8 + 4 + 0, undefined);

          A.store.Bool(ptr + 8 + 4 + 4 + 16, false);
          A.store.Ref(ptr + 8 + 4 + 4 + 0, undefined);
          A.store.Ref(ptr + 8 + 4 + 4 + 4, undefined);
          A.store.Ref(ptr + 8 + 4 + 4 + 8, undefined);
          A.store.Ref(ptr + 8 + 4 + 4 + 12, undefined);
          A.store.Ref(ptr + 8 + 4 + 24, undefined);

          A.store.Bool(ptr + 8 + 4 + 28 + 16, false);
          A.store.Ref(ptr + 8 + 4 + 28 + 0, undefined);
          A.store.Ref(ptr + 8 + 4 + 28 + 4, undefined);
          A.store.Ref(ptr + 8 + 4 + 28 + 8, undefined);
          A.store.Ref(ptr + 8 + 4 + 28 + 12, undefined);
          A.store.Ref(ptr + 8 + 52, undefined);
          A.store.Ref(ptr + 8 + 56, undefined);
          A.store.Ref(ptr + 8 + 60, undefined);
          A.store.Enum(ptr + 8 + 64, -1);
          A.store.Ref(ptr + 8 + 68, undefined);
          A.store.Ref(ptr + 8 + 72, undefined);
          A.store.Ref(ptr + 8 + 76, undefined);
          A.store.Ref(ptr + 8 + 80, undefined);
          A.store.Bool(ptr + 8 + 127, false);
          A.store.Bool(ptr + 8 + 84, false);
          A.store.Ref(ptr + 8 + 88, undefined);
          A.store.Ref(ptr + 8 + 92, undefined);

          A.store.Bool(ptr + 8 + 96 + 28, false);
          A.store.Ref(ptr + 8 + 96 + 0, undefined);
          A.store.Ref(ptr + 8 + 96 + 4, undefined);
          A.store.Ref(ptr + 8 + 96 + 8, undefined);
          A.store.Ref(ptr + 8 + 96 + 12, undefined);
          A.store.Ref(ptr + 8 + 96 + 16, undefined);
          A.store.Ref(ptr + 8 + 96 + 20, undefined);
          A.store.Bool(ptr + 8 + 96 + 26, false);
          A.store.Bool(ptr + 8 + 96 + 24, false);
          A.store.Bool(ptr + 8 + 96 + 27, false);
          A.store.Bool(ptr + 8 + 96 + 25, false);
          A.store.Bool(ptr + 8 + 128, false);
          A.store.Bool(ptr + 8 + 125, false);
          A.store.Bool(ptr + 8 + 129, false);
          A.store.Bool(ptr + 8 + 126, false);
        } else {
          A.store.Bool(ptr + 8 + 130, true);
          A.store.Ref(ptr + 8 + 0, x["EAP"]["AnonymousIdentity"]);

          if (typeof x["EAP"]["ClientCertPattern"] === "undefined") {
            A.store.Bool(ptr + 8 + 4 + 45, false);
            A.store.Ref(ptr + 8 + 4 + 0, undefined);

            A.store.Bool(ptr + 8 + 4 + 4 + 16, false);
            A.store.Ref(ptr + 8 + 4 + 4 + 0, undefined);
            A.store.Ref(ptr + 8 + 4 + 4 + 4, undefined);
            A.store.Ref(ptr + 8 + 4 + 4 + 8, undefined);
            A.store.Ref(ptr + 8 + 4 + 4 + 12, undefined);
            A.store.Ref(ptr + 8 + 4 + 24, undefined);

            A.store.Bool(ptr + 8 + 4 + 28 + 16, false);
            A.store.Ref(ptr + 8 + 4 + 28 + 0, undefined);
            A.store.Ref(ptr + 8 + 4 + 28 + 4, undefined);
            A.store.Ref(ptr + 8 + 4 + 28 + 8, undefined);
            A.store.Ref(ptr + 8 + 4 + 28 + 12, undefined);
          } else {
            A.store.Bool(ptr + 8 + 4 + 45, true);
            A.store.Ref(ptr + 8 + 4 + 0, x["EAP"]["ClientCertPattern"]["EnrollmentURI"]);

            if (typeof x["EAP"]["ClientCertPattern"]["Issuer"] === "undefined") {
              A.store.Bool(ptr + 8 + 4 + 4 + 16, false);
              A.store.Ref(ptr + 8 + 4 + 4 + 0, undefined);
              A.store.Ref(ptr + 8 + 4 + 4 + 4, undefined);
              A.store.Ref(ptr + 8 + 4 + 4 + 8, undefined);
              A.store.Ref(ptr + 8 + 4 + 4 + 12, undefined);
            } else {
              A.store.Bool(ptr + 8 + 4 + 4 + 16, true);
              A.store.Ref(ptr + 8 + 4 + 4 + 0, x["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"]);
              A.store.Ref(ptr + 8 + 4 + 4 + 4, x["EAP"]["ClientCertPattern"]["Issuer"]["Locality"]);
              A.store.Ref(ptr + 8 + 4 + 4 + 8, x["EAP"]["ClientCertPattern"]["Issuer"]["Organization"]);
              A.store.Ref(ptr + 8 + 4 + 4 + 12, x["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"]);
            }
            A.store.Ref(ptr + 8 + 4 + 24, x["EAP"]["ClientCertPattern"]["IssuerCARef"]);

            if (typeof x["EAP"]["ClientCertPattern"]["Subject"] === "undefined") {
              A.store.Bool(ptr + 8 + 4 + 28 + 16, false);
              A.store.Ref(ptr + 8 + 4 + 28 + 0, undefined);
              A.store.Ref(ptr + 8 + 4 + 28 + 4, undefined);
              A.store.Ref(ptr + 8 + 4 + 28 + 8, undefined);
              A.store.Ref(ptr + 8 + 4 + 28 + 12, undefined);
            } else {
              A.store.Bool(ptr + 8 + 4 + 28 + 16, true);
              A.store.Ref(ptr + 8 + 4 + 28 + 0, x["EAP"]["ClientCertPattern"]["Subject"]["CommonName"]);
              A.store.Ref(ptr + 8 + 4 + 28 + 4, x["EAP"]["ClientCertPattern"]["Subject"]["Locality"]);
              A.store.Ref(ptr + 8 + 4 + 28 + 8, x["EAP"]["ClientCertPattern"]["Subject"]["Organization"]);
              A.store.Ref(ptr + 8 + 4 + 28 + 12, x["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"]);
            }
          }
          A.store.Ref(ptr + 8 + 52, x["EAP"]["ClientCertPKCS11Id"]);
          A.store.Ref(ptr + 8 + 56, x["EAP"]["ClientCertProvisioningProfileId"]);
          A.store.Ref(ptr + 8 + 60, x["EAP"]["ClientCertRef"]);
          A.store.Enum(ptr + 8 + 64, ["Ref", "Pattern"].indexOf(x["EAP"]["ClientCertType"] as string));
          A.store.Ref(ptr + 8 + 68, x["EAP"]["Identity"]);
          A.store.Ref(ptr + 8 + 72, x["EAP"]["Inner"]);
          A.store.Ref(ptr + 8 + 76, x["EAP"]["Outer"]);
          A.store.Ref(ptr + 8 + 80, x["EAP"]["Password"]);
          A.store.Bool(ptr + 8 + 127, "SaveCredentials" in x["EAP"] ? true : false);
          A.store.Bool(ptr + 8 + 84, x["EAP"]["SaveCredentials"] ? true : false);
          A.store.Ref(ptr + 8 + 88, x["EAP"]["ServerCAPEMs"]);
          A.store.Ref(ptr + 8 + 92, x["EAP"]["ServerCARefs"]);

          if (typeof x["EAP"]["SubjectMatch"] === "undefined") {
            A.store.Bool(ptr + 8 + 96 + 28, false);
            A.store.Ref(ptr + 8 + 96 + 0, undefined);
            A.store.Ref(ptr + 8 + 96 + 4, undefined);
            A.store.Ref(ptr + 8 + 96 + 8, undefined);
            A.store.Ref(ptr + 8 + 96 + 12, undefined);
            A.store.Ref(ptr + 8 + 96 + 16, undefined);
            A.store.Ref(ptr + 8 + 96 + 20, undefined);
            A.store.Bool(ptr + 8 + 96 + 26, false);
            A.store.Bool(ptr + 8 + 96 + 24, false);
            A.store.Bool(ptr + 8 + 96 + 27, false);
            A.store.Bool(ptr + 8 + 96 + 25, false);
          } else {
            A.store.Bool(ptr + 8 + 96 + 28, true);
            A.store.Ref(ptr + 8 + 96 + 0, x["EAP"]["SubjectMatch"]["Active"]);
            A.store.Ref(ptr + 8 + 96 + 4, x["EAP"]["SubjectMatch"]["Effective"]);
            A.store.Ref(ptr + 8 + 96 + 8, x["EAP"]["SubjectMatch"]["UserPolicy"]);
            A.store.Ref(ptr + 8 + 96 + 12, x["EAP"]["SubjectMatch"]["DevicePolicy"]);
            A.store.Ref(ptr + 8 + 96 + 16, x["EAP"]["SubjectMatch"]["UserSetting"]);
            A.store.Ref(ptr + 8 + 96 + 20, x["EAP"]["SubjectMatch"]["SharedSetting"]);
            A.store.Bool(ptr + 8 + 96 + 26, "UserEditable" in x["EAP"]["SubjectMatch"] ? true : false);
            A.store.Bool(ptr + 8 + 96 + 24, x["EAP"]["SubjectMatch"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 8 + 96 + 27, "DeviceEditable" in x["EAP"]["SubjectMatch"] ? true : false);
            A.store.Bool(ptr + 8 + 96 + 25, x["EAP"]["SubjectMatch"]["DeviceEditable"] ? true : false);
          }
          A.store.Bool(ptr + 8 + 128, "UseProactiveKeyCaching" in x["EAP"] ? true : false);
          A.store.Bool(ptr + 8 + 125, x["EAP"]["UseProactiveKeyCaching"] ? true : false);
          A.store.Bool(ptr + 8 + 129, "UseSystemCAs" in x["EAP"] ? true : false);
          A.store.Bool(ptr + 8 + 126, x["EAP"]["UseSystemCAs"] ? true : false);
        }
        A.store.Bool(ptr + 178, "Frequency" in x ? true : false);
        A.store.Int32(ptr + 140, x["Frequency"] === undefined ? 0 : (x["Frequency"] as number));
        A.store.Ref(ptr + 144, x["FrequencyList"]);
        A.store.Ref(ptr + 148, x["HexSSID"]);
        A.store.Bool(ptr + 179, "HiddenSSID" in x ? true : false);
        A.store.Bool(ptr + 152, x["HiddenSSID"] ? true : false);
        A.store.Ref(ptr + 156, x["Passphrase"]);
        A.store.Bool(ptr + 180, "RoamThreshold" in x ? true : false);
        A.store.Int32(ptr + 160, x["RoamThreshold"] === undefined ? 0 : (x["RoamThreshold"] as number));
        A.store.Ref(ptr + 164, x["SSID"]);
        A.store.Ref(ptr + 168, x["Security"]);
        A.store.Bool(ptr + 181, "SignalStrength" in x ? true : false);
        A.store.Int32(ptr + 172, x["SignalStrength"] === undefined ? 0 : (x["SignalStrength"] as number));
      }
    },
    "load_WiFiProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 176)) {
        x["AllowGatewayARPPolling"] = A.load.Bool(ptr + 0);
      } else {
        delete x["AllowGatewayARPPolling"];
      }
      if (A.load.Bool(ptr + 177)) {
        x["AutoConnect"] = A.load.Bool(ptr + 1);
      } else {
        delete x["AutoConnect"];
      }
      x["BSSID"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 8 + 130)) {
        x["EAP"] = {};
        x["EAP"]["AnonymousIdentity"] = A.load.Ref(ptr + 8 + 0, undefined);
        if (A.load.Bool(ptr + 8 + 4 + 45)) {
          x["EAP"]["ClientCertPattern"] = {};
          x["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(ptr + 8 + 4 + 0, undefined);
          if (A.load.Bool(ptr + 8 + 4 + 4 + 16)) {
            x["EAP"]["ClientCertPattern"]["Issuer"] = {};
            x["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(ptr + 8 + 4 + 4 + 0, undefined);
            x["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(ptr + 8 + 4 + 4 + 4, undefined);
            x["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(ptr + 8 + 4 + 4 + 8, undefined);
            x["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(ptr + 8 + 4 + 4 + 12, undefined);
          } else {
            delete x["EAP"]["ClientCertPattern"]["Issuer"];
          }
          x["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(ptr + 8 + 4 + 24, undefined);
          if (A.load.Bool(ptr + 8 + 4 + 28 + 16)) {
            x["EAP"]["ClientCertPattern"]["Subject"] = {};
            x["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(ptr + 8 + 4 + 28 + 0, undefined);
            x["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(ptr + 8 + 4 + 28 + 4, undefined);
            x["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(ptr + 8 + 4 + 28 + 8, undefined);
            x["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
              ptr + 8 + 4 + 28 + 12,
              undefined
            );
          } else {
            delete x["EAP"]["ClientCertPattern"]["Subject"];
          }
        } else {
          delete x["EAP"]["ClientCertPattern"];
        }
        x["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(ptr + 8 + 52, undefined);
        x["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(ptr + 8 + 56, undefined);
        x["EAP"]["ClientCertRef"] = A.load.Ref(ptr + 8 + 60, undefined);
        x["EAP"]["ClientCertType"] = A.load.Enum(ptr + 8 + 64, ["Ref", "Pattern"]);
        x["EAP"]["Identity"] = A.load.Ref(ptr + 8 + 68, undefined);
        x["EAP"]["Inner"] = A.load.Ref(ptr + 8 + 72, undefined);
        x["EAP"]["Outer"] = A.load.Ref(ptr + 8 + 76, undefined);
        x["EAP"]["Password"] = A.load.Ref(ptr + 8 + 80, undefined);
        if (A.load.Bool(ptr + 8 + 127)) {
          x["EAP"]["SaveCredentials"] = A.load.Bool(ptr + 8 + 84);
        } else {
          delete x["EAP"]["SaveCredentials"];
        }
        x["EAP"]["ServerCAPEMs"] = A.load.Ref(ptr + 8 + 88, undefined);
        x["EAP"]["ServerCARefs"] = A.load.Ref(ptr + 8 + 92, undefined);
        if (A.load.Bool(ptr + 8 + 96 + 28)) {
          x["EAP"]["SubjectMatch"] = {};
          x["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(ptr + 8 + 96 + 0, undefined);
          x["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(ptr + 8 + 96 + 4, undefined);
          x["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(ptr + 8 + 96 + 8, undefined);
          x["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(ptr + 8 + 96 + 12, undefined);
          x["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(ptr + 8 + 96 + 16, undefined);
          x["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(ptr + 8 + 96 + 20, undefined);
          if (A.load.Bool(ptr + 8 + 96 + 26)) {
            x["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(ptr + 8 + 96 + 24);
          } else {
            delete x["EAP"]["SubjectMatch"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 8 + 96 + 27)) {
            x["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(ptr + 8 + 96 + 25);
          } else {
            delete x["EAP"]["SubjectMatch"]["DeviceEditable"];
          }
        } else {
          delete x["EAP"]["SubjectMatch"];
        }
        if (A.load.Bool(ptr + 8 + 128)) {
          x["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(ptr + 8 + 125);
        } else {
          delete x["EAP"]["UseProactiveKeyCaching"];
        }
        if (A.load.Bool(ptr + 8 + 129)) {
          x["EAP"]["UseSystemCAs"] = A.load.Bool(ptr + 8 + 126);
        } else {
          delete x["EAP"]["UseSystemCAs"];
        }
      } else {
        delete x["EAP"];
      }
      if (A.load.Bool(ptr + 178)) {
        x["Frequency"] = A.load.Int32(ptr + 140);
      } else {
        delete x["Frequency"];
      }
      x["FrequencyList"] = A.load.Ref(ptr + 144, undefined);
      x["HexSSID"] = A.load.Ref(ptr + 148, undefined);
      if (A.load.Bool(ptr + 179)) {
        x["HiddenSSID"] = A.load.Bool(ptr + 152);
      } else {
        delete x["HiddenSSID"];
      }
      x["Passphrase"] = A.load.Ref(ptr + 156, undefined);
      if (A.load.Bool(ptr + 180)) {
        x["RoamThreshold"] = A.load.Int32(ptr + 160);
      } else {
        delete x["RoamThreshold"];
      }
      x["SSID"] = A.load.Ref(ptr + 164, undefined);
      x["Security"] = A.load.Ref(ptr + 168, undefined);
      if (A.load.Bool(ptr + 181)) {
        x["SignalStrength"] = A.load.Int32(ptr + 172);
      } else {
        delete x["SignalStrength"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_NetworkProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 671, false);

        A.store.Bool(ptr + 0 + 127, false);
        A.store.Bool(ptr + 0 + 121, false);
        A.store.Bool(ptr + 0 + 0, false);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Enum(ptr + 0 + 8, -1);
        A.store.Bool(ptr + 0 + 122, false);
        A.store.Bool(ptr + 0 + 12, false);
        A.store.Ref(ptr + 0 + 16, undefined);
        A.store.Ref(ptr + 0 + 20, undefined);
        A.store.Ref(ptr + 0 + 24, undefined);
        A.store.Ref(ptr + 0 + 28, undefined);

        A.store.Bool(ptr + 0 + 32 + 12, false);
        A.store.Ref(ptr + 0 + 32 + 0, undefined);
        A.store.Ref(ptr + 0 + 32 + 4, undefined);
        A.store.Ref(ptr + 0 + 32 + 8, undefined);
        A.store.Ref(ptr + 0 + 48, undefined);
        A.store.Ref(ptr + 0 + 52, undefined);
        A.store.Ref(ptr + 0 + 56, undefined);

        A.store.Bool(ptr + 0 + 60 + 12, false);
        A.store.Ref(ptr + 0 + 60 + 0, undefined);
        A.store.Ref(ptr + 0 + 60 + 4, undefined);
        A.store.Ref(ptr + 0 + 60 + 8, undefined);
        A.store.Ref(ptr + 0 + 76, undefined);
        A.store.Bool(ptr + 0 + 123, false);
        A.store.Bool(ptr + 0 + 80, false);

        A.store.Bool(ptr + 0 + 84 + 12, false);
        A.store.Ref(ptr + 0 + 84 + 0, undefined);
        A.store.Ref(ptr + 0 + 84 + 4, undefined);
        A.store.Ref(ptr + 0 + 84 + 8, undefined);

        A.store.Bool(ptr + 0 + 100 + 14, false);
        A.store.Ref(ptr + 0 + 100 + 0, undefined);
        A.store.Bool(ptr + 0 + 100 + 12, false);
        A.store.Bool(ptr + 0 + 100 + 4, false);
        A.store.Bool(ptr + 0 + 100 + 13, false);
        A.store.Int32(ptr + 0 + 100 + 8, 0);
        A.store.Bool(ptr + 0 + 124, false);
        A.store.Bool(ptr + 0 + 115, false);
        A.store.Bool(ptr + 0 + 125, false);
        A.store.Int32(ptr + 0 + 116, 0);
        A.store.Bool(ptr + 0 + 126, false);
        A.store.Bool(ptr + 0 + 120, false);
        A.store.Bool(ptr + 667, false);
        A.store.Bool(ptr + 128, false);
        A.store.Enum(ptr + 132, -1);
        A.store.Ref(ptr + 136, undefined);

        A.store.Bool(ptr + 140 + 140, false);
        A.store.Bool(ptr + 140 + 139, false);
        A.store.Bool(ptr + 140 + 0, false);
        A.store.Ref(ptr + 140 + 4, undefined);

        A.store.Bool(ptr + 140 + 8 + 130, false);
        A.store.Ref(ptr + 140 + 8 + 0, undefined);

        A.store.Bool(ptr + 140 + 8 + 4 + 45, false);
        A.store.Ref(ptr + 140 + 8 + 4 + 0, undefined);

        A.store.Bool(ptr + 140 + 8 + 4 + 4 + 16, false);
        A.store.Ref(ptr + 140 + 8 + 4 + 4 + 0, undefined);
        A.store.Ref(ptr + 140 + 8 + 4 + 4 + 4, undefined);
        A.store.Ref(ptr + 140 + 8 + 4 + 4 + 8, undefined);
        A.store.Ref(ptr + 140 + 8 + 4 + 4 + 12, undefined);
        A.store.Ref(ptr + 140 + 8 + 4 + 24, undefined);

        A.store.Bool(ptr + 140 + 8 + 4 + 28 + 16, false);
        A.store.Ref(ptr + 140 + 8 + 4 + 28 + 0, undefined);
        A.store.Ref(ptr + 140 + 8 + 4 + 28 + 4, undefined);
        A.store.Ref(ptr + 140 + 8 + 4 + 28 + 8, undefined);
        A.store.Ref(ptr + 140 + 8 + 4 + 28 + 12, undefined);
        A.store.Ref(ptr + 140 + 8 + 52, undefined);
        A.store.Ref(ptr + 140 + 8 + 56, undefined);
        A.store.Ref(ptr + 140 + 8 + 60, undefined);
        A.store.Enum(ptr + 140 + 8 + 64, -1);
        A.store.Ref(ptr + 140 + 8 + 68, undefined);
        A.store.Ref(ptr + 140 + 8 + 72, undefined);
        A.store.Ref(ptr + 140 + 8 + 76, undefined);
        A.store.Ref(ptr + 140 + 8 + 80, undefined);
        A.store.Bool(ptr + 140 + 8 + 127, false);
        A.store.Bool(ptr + 140 + 8 + 84, false);
        A.store.Ref(ptr + 140 + 8 + 88, undefined);
        A.store.Ref(ptr + 140 + 8 + 92, undefined);

        A.store.Bool(ptr + 140 + 8 + 96 + 28, false);
        A.store.Ref(ptr + 140 + 8 + 96 + 0, undefined);
        A.store.Ref(ptr + 140 + 8 + 96 + 4, undefined);
        A.store.Ref(ptr + 140 + 8 + 96 + 8, undefined);
        A.store.Ref(ptr + 140 + 8 + 96 + 12, undefined);
        A.store.Ref(ptr + 140 + 8 + 96 + 16, undefined);
        A.store.Ref(ptr + 140 + 8 + 96 + 20, undefined);
        A.store.Bool(ptr + 140 + 8 + 96 + 26, false);
        A.store.Bool(ptr + 140 + 8 + 96 + 24, false);
        A.store.Bool(ptr + 140 + 8 + 96 + 27, false);
        A.store.Bool(ptr + 140 + 8 + 96 + 25, false);
        A.store.Bool(ptr + 140 + 8 + 128, false);
        A.store.Bool(ptr + 140 + 8 + 125, false);
        A.store.Bool(ptr + 140 + 8 + 129, false);
        A.store.Bool(ptr + 140 + 8 + 126, false);
        A.store.Ref(ptr + 284, undefined);
        A.store.Enum(ptr + 288, -1);
        A.store.Ref(ptr + 292, undefined);
        A.store.Ref(ptr + 296, undefined);
        A.store.Bool(ptr + 668, false);
        A.store.Bool(ptr + 300, false);
        A.store.Ref(ptr + 304, undefined);
        A.store.Enum(ptr + 308, -1);
        A.store.Bool(ptr + 669, false);
        A.store.Int32(ptr + 312, 0);

        A.store.Bool(ptr + 316 + 60, false);
        A.store.Enum(ptr + 316 + 0, -1);

        A.store.Bool(ptr + 316 + 4 + 46, false);

        A.store.Bool(ptr + 316 + 4 + 0 + 9, false);
        A.store.Ref(ptr + 316 + 4 + 0 + 0, undefined);
        A.store.Bool(ptr + 316 + 4 + 0 + 8, false);
        A.store.Int32(ptr + 316 + 4 + 0 + 4, 0);

        A.store.Bool(ptr + 316 + 4 + 12 + 9, false);
        A.store.Ref(ptr + 316 + 4 + 12 + 0, undefined);
        A.store.Bool(ptr + 316 + 4 + 12 + 8, false);
        A.store.Int32(ptr + 316 + 4 + 12 + 4, 0);

        A.store.Bool(ptr + 316 + 4 + 24 + 9, false);
        A.store.Ref(ptr + 316 + 4 + 24 + 0, undefined);
        A.store.Bool(ptr + 316 + 4 + 24 + 8, false);
        A.store.Int32(ptr + 316 + 4 + 24 + 4, 0);

        A.store.Bool(ptr + 316 + 4 + 36 + 9, false);
        A.store.Ref(ptr + 316 + 4 + 36 + 0, undefined);
        A.store.Bool(ptr + 316 + 4 + 36 + 8, false);
        A.store.Int32(ptr + 316 + 4 + 36 + 4, 0);
        A.store.Ref(ptr + 316 + 52, undefined);
        A.store.Ref(ptr + 316 + 56, undefined);
        A.store.Bool(ptr + 670, false);
        A.store.Bool(ptr + 377, false);

        A.store.Bool(ptr + 380 + 37, false);
        A.store.Ref(ptr + 380 + 0, undefined);
        A.store.Ref(ptr + 380 + 4, undefined);
        A.store.Ref(ptr + 380 + 8, undefined);
        A.store.Ref(ptr + 380 + 12, undefined);
        A.store.Ref(ptr + 380 + 16, undefined);
        A.store.Ref(ptr + 380 + 20, undefined);
        A.store.Bool(ptr + 380 + 36, false);
        A.store.Int32(ptr + 380 + 24, 0);
        A.store.Ref(ptr + 380 + 28, undefined);
        A.store.Ref(ptr + 380 + 32, undefined);

        A.store.Bool(ptr + 420 + 37, false);
        A.store.Ref(ptr + 420 + 0, undefined);
        A.store.Ref(ptr + 420 + 4, undefined);
        A.store.Ref(ptr + 420 + 8, undefined);
        A.store.Ref(ptr + 420 + 12, undefined);
        A.store.Ref(ptr + 420 + 16, undefined);
        A.store.Ref(ptr + 420 + 20, undefined);
        A.store.Bool(ptr + 420 + 36, false);
        A.store.Int32(ptr + 420 + 24, 0);
        A.store.Ref(ptr + 420 + 28, undefined);
        A.store.Ref(ptr + 420 + 32, undefined);
        A.store.Ref(ptr + 460, undefined);
        A.store.Enum(ptr + 464, -1);

        A.store.Bool(ptr + 468 + 13, false);
        A.store.Bool(ptr + 468 + 12, false);
        A.store.Bool(ptr + 468 + 0, false);
        A.store.Ref(ptr + 468 + 4, undefined);
        A.store.Ref(ptr + 468 + 8, undefined);

        A.store.Bool(ptr + 484 + 182, false);
        A.store.Bool(ptr + 484 + 176, false);
        A.store.Bool(ptr + 484 + 0, false);
        A.store.Bool(ptr + 484 + 177, false);
        A.store.Bool(ptr + 484 + 1, false);
        A.store.Ref(ptr + 484 + 4, undefined);

        A.store.Bool(ptr + 484 + 8 + 130, false);
        A.store.Ref(ptr + 484 + 8 + 0, undefined);

        A.store.Bool(ptr + 484 + 8 + 4 + 45, false);
        A.store.Ref(ptr + 484 + 8 + 4 + 0, undefined);

        A.store.Bool(ptr + 484 + 8 + 4 + 4 + 16, false);
        A.store.Ref(ptr + 484 + 8 + 4 + 4 + 0, undefined);
        A.store.Ref(ptr + 484 + 8 + 4 + 4 + 4, undefined);
        A.store.Ref(ptr + 484 + 8 + 4 + 4 + 8, undefined);
        A.store.Ref(ptr + 484 + 8 + 4 + 4 + 12, undefined);
        A.store.Ref(ptr + 484 + 8 + 4 + 24, undefined);

        A.store.Bool(ptr + 484 + 8 + 4 + 28 + 16, false);
        A.store.Ref(ptr + 484 + 8 + 4 + 28 + 0, undefined);
        A.store.Ref(ptr + 484 + 8 + 4 + 28 + 4, undefined);
        A.store.Ref(ptr + 484 + 8 + 4 + 28 + 8, undefined);
        A.store.Ref(ptr + 484 + 8 + 4 + 28 + 12, undefined);
        A.store.Ref(ptr + 484 + 8 + 52, undefined);
        A.store.Ref(ptr + 484 + 8 + 56, undefined);
        A.store.Ref(ptr + 484 + 8 + 60, undefined);
        A.store.Enum(ptr + 484 + 8 + 64, -1);
        A.store.Ref(ptr + 484 + 8 + 68, undefined);
        A.store.Ref(ptr + 484 + 8 + 72, undefined);
        A.store.Ref(ptr + 484 + 8 + 76, undefined);
        A.store.Ref(ptr + 484 + 8 + 80, undefined);
        A.store.Bool(ptr + 484 + 8 + 127, false);
        A.store.Bool(ptr + 484 + 8 + 84, false);
        A.store.Ref(ptr + 484 + 8 + 88, undefined);
        A.store.Ref(ptr + 484 + 8 + 92, undefined);

        A.store.Bool(ptr + 484 + 8 + 96 + 28, false);
        A.store.Ref(ptr + 484 + 8 + 96 + 0, undefined);
        A.store.Ref(ptr + 484 + 8 + 96 + 4, undefined);
        A.store.Ref(ptr + 484 + 8 + 96 + 8, undefined);
        A.store.Ref(ptr + 484 + 8 + 96 + 12, undefined);
        A.store.Ref(ptr + 484 + 8 + 96 + 16, undefined);
        A.store.Ref(ptr + 484 + 8 + 96 + 20, undefined);
        A.store.Bool(ptr + 484 + 8 + 96 + 26, false);
        A.store.Bool(ptr + 484 + 8 + 96 + 24, false);
        A.store.Bool(ptr + 484 + 8 + 96 + 27, false);
        A.store.Bool(ptr + 484 + 8 + 96 + 25, false);
        A.store.Bool(ptr + 484 + 8 + 128, false);
        A.store.Bool(ptr + 484 + 8 + 125, false);
        A.store.Bool(ptr + 484 + 8 + 129, false);
        A.store.Bool(ptr + 484 + 8 + 126, false);
        A.store.Bool(ptr + 484 + 178, false);
        A.store.Int32(ptr + 484 + 140, 0);
        A.store.Ref(ptr + 484 + 144, undefined);
        A.store.Ref(ptr + 484 + 148, undefined);
        A.store.Bool(ptr + 484 + 179, false);
        A.store.Bool(ptr + 484 + 152, false);
        A.store.Ref(ptr + 484 + 156, undefined);
        A.store.Bool(ptr + 484 + 180, false);
        A.store.Int32(ptr + 484 + 160, 0);
        A.store.Ref(ptr + 484 + 164, undefined);
        A.store.Ref(ptr + 484 + 168, undefined);
        A.store.Bool(ptr + 484 + 181, false);
        A.store.Int32(ptr + 484 + 172, 0);
      } else {
        A.store.Bool(ptr + 671, true);

        if (typeof x["Cellular"] === "undefined") {
          A.store.Bool(ptr + 0 + 127, false);
          A.store.Bool(ptr + 0 + 121, false);
          A.store.Bool(ptr + 0 + 0, false);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Enum(ptr + 0 + 8, -1);
          A.store.Bool(ptr + 0 + 122, false);
          A.store.Bool(ptr + 0 + 12, false);
          A.store.Ref(ptr + 0 + 16, undefined);
          A.store.Ref(ptr + 0 + 20, undefined);
          A.store.Ref(ptr + 0 + 24, undefined);
          A.store.Ref(ptr + 0 + 28, undefined);

          A.store.Bool(ptr + 0 + 32 + 12, false);
          A.store.Ref(ptr + 0 + 32 + 0, undefined);
          A.store.Ref(ptr + 0 + 32 + 4, undefined);
          A.store.Ref(ptr + 0 + 32 + 8, undefined);
          A.store.Ref(ptr + 0 + 48, undefined);
          A.store.Ref(ptr + 0 + 52, undefined);
          A.store.Ref(ptr + 0 + 56, undefined);

          A.store.Bool(ptr + 0 + 60 + 12, false);
          A.store.Ref(ptr + 0 + 60 + 0, undefined);
          A.store.Ref(ptr + 0 + 60 + 4, undefined);
          A.store.Ref(ptr + 0 + 60 + 8, undefined);
          A.store.Ref(ptr + 0 + 76, undefined);
          A.store.Bool(ptr + 0 + 123, false);
          A.store.Bool(ptr + 0 + 80, false);

          A.store.Bool(ptr + 0 + 84 + 12, false);
          A.store.Ref(ptr + 0 + 84 + 0, undefined);
          A.store.Ref(ptr + 0 + 84 + 4, undefined);
          A.store.Ref(ptr + 0 + 84 + 8, undefined);

          A.store.Bool(ptr + 0 + 100 + 14, false);
          A.store.Ref(ptr + 0 + 100 + 0, undefined);
          A.store.Bool(ptr + 0 + 100 + 12, false);
          A.store.Bool(ptr + 0 + 100 + 4, false);
          A.store.Bool(ptr + 0 + 100 + 13, false);
          A.store.Int32(ptr + 0 + 100 + 8, 0);
          A.store.Bool(ptr + 0 + 124, false);
          A.store.Bool(ptr + 0 + 115, false);
          A.store.Bool(ptr + 0 + 125, false);
          A.store.Int32(ptr + 0 + 116, 0);
          A.store.Bool(ptr + 0 + 126, false);
          A.store.Bool(ptr + 0 + 120, false);
        } else {
          A.store.Bool(ptr + 0 + 127, true);
          A.store.Bool(ptr + 0 + 121, "AutoConnect" in x["Cellular"] ? true : false);
          A.store.Bool(ptr + 0 + 0, x["Cellular"]["AutoConnect"] ? true : false);
          A.store.Ref(ptr + 0 + 4, x["Cellular"]["ActivationType"]);
          A.store.Enum(
            ptr + 0 + 8,
            ["Activated", "Activating", "NotActivated", "PartiallyActivated"].indexOf(
              x["Cellular"]["ActivationState"] as string
            )
          );
          A.store.Bool(ptr + 0 + 122, "AllowRoaming" in x["Cellular"] ? true : false);
          A.store.Bool(ptr + 0 + 12, x["Cellular"]["AllowRoaming"] ? true : false);
          A.store.Ref(ptr + 0 + 16, x["Cellular"]["Family"]);
          A.store.Ref(ptr + 0 + 20, x["Cellular"]["FirmwareRevision"]);
          A.store.Ref(ptr + 0 + 24, x["Cellular"]["FoundNetworks"]);
          A.store.Ref(ptr + 0 + 28, x["Cellular"]["HardwareRevision"]);

          if (typeof x["Cellular"]["HomeProvider"] === "undefined") {
            A.store.Bool(ptr + 0 + 32 + 12, false);
            A.store.Ref(ptr + 0 + 32 + 0, undefined);
            A.store.Ref(ptr + 0 + 32 + 4, undefined);
            A.store.Ref(ptr + 0 + 32 + 8, undefined);
          } else {
            A.store.Bool(ptr + 0 + 32 + 12, true);
            A.store.Ref(ptr + 0 + 32 + 0, x["Cellular"]["HomeProvider"]["Name"]);
            A.store.Ref(ptr + 0 + 32 + 4, x["Cellular"]["HomeProvider"]["Code"]);
            A.store.Ref(ptr + 0 + 32 + 8, x["Cellular"]["HomeProvider"]["Country"]);
          }
          A.store.Ref(ptr + 0 + 48, x["Cellular"]["Manufacturer"]);
          A.store.Ref(ptr + 0 + 52, x["Cellular"]["ModelID"]);
          A.store.Ref(ptr + 0 + 56, x["Cellular"]["NetworkTechnology"]);

          if (typeof x["Cellular"]["PaymentPortal"] === "undefined") {
            A.store.Bool(ptr + 0 + 60 + 12, false);
            A.store.Ref(ptr + 0 + 60 + 0, undefined);
            A.store.Ref(ptr + 0 + 60 + 4, undefined);
            A.store.Ref(ptr + 0 + 60 + 8, undefined);
          } else {
            A.store.Bool(ptr + 0 + 60 + 12, true);
            A.store.Ref(ptr + 0 + 60 + 0, x["Cellular"]["PaymentPortal"]["Method"]);
            A.store.Ref(ptr + 0 + 60 + 4, x["Cellular"]["PaymentPortal"]["PostData"]);
            A.store.Ref(ptr + 0 + 60 + 8, x["Cellular"]["PaymentPortal"]["Url"]);
          }
          A.store.Ref(ptr + 0 + 76, x["Cellular"]["RoamingState"]);
          A.store.Bool(ptr + 0 + 123, "Scanning" in x["Cellular"] ? true : false);
          A.store.Bool(ptr + 0 + 80, x["Cellular"]["Scanning"] ? true : false);

          if (typeof x["Cellular"]["ServingOperator"] === "undefined") {
            A.store.Bool(ptr + 0 + 84 + 12, false);
            A.store.Ref(ptr + 0 + 84 + 0, undefined);
            A.store.Ref(ptr + 0 + 84 + 4, undefined);
            A.store.Ref(ptr + 0 + 84 + 8, undefined);
          } else {
            A.store.Bool(ptr + 0 + 84 + 12, true);
            A.store.Ref(ptr + 0 + 84 + 0, x["Cellular"]["ServingOperator"]["Name"]);
            A.store.Ref(ptr + 0 + 84 + 4, x["Cellular"]["ServingOperator"]["Code"]);
            A.store.Ref(ptr + 0 + 84 + 8, x["Cellular"]["ServingOperator"]["Country"]);
          }

          if (typeof x["Cellular"]["SIMLockStatus"] === "undefined") {
            A.store.Bool(ptr + 0 + 100 + 14, false);
            A.store.Ref(ptr + 0 + 100 + 0, undefined);
            A.store.Bool(ptr + 0 + 100 + 12, false);
            A.store.Bool(ptr + 0 + 100 + 4, false);
            A.store.Bool(ptr + 0 + 100 + 13, false);
            A.store.Int32(ptr + 0 + 100 + 8, 0);
          } else {
            A.store.Bool(ptr + 0 + 100 + 14, true);
            A.store.Ref(ptr + 0 + 100 + 0, x["Cellular"]["SIMLockStatus"]["LockType"]);
            A.store.Bool(ptr + 0 + 100 + 12, "LockEnabled" in x["Cellular"]["SIMLockStatus"] ? true : false);
            A.store.Bool(ptr + 0 + 100 + 4, x["Cellular"]["SIMLockStatus"]["LockEnabled"] ? true : false);
            A.store.Bool(ptr + 0 + 100 + 13, "RetriesLeft" in x["Cellular"]["SIMLockStatus"] ? true : false);
            A.store.Int32(
              ptr + 0 + 100 + 8,
              x["Cellular"]["SIMLockStatus"]["RetriesLeft"] === undefined
                ? 0
                : (x["Cellular"]["SIMLockStatus"]["RetriesLeft"] as number)
            );
          }
          A.store.Bool(ptr + 0 + 124, "SIMPresent" in x["Cellular"] ? true : false);
          A.store.Bool(ptr + 0 + 115, x["Cellular"]["SIMPresent"] ? true : false);
          A.store.Bool(ptr + 0 + 125, "SignalStrength" in x["Cellular"] ? true : false);
          A.store.Int32(
            ptr + 0 + 116,
            x["Cellular"]["SignalStrength"] === undefined ? 0 : (x["Cellular"]["SignalStrength"] as number)
          );
          A.store.Bool(ptr + 0 + 126, "SupportNetworkScan" in x["Cellular"] ? true : false);
          A.store.Bool(ptr + 0 + 120, x["Cellular"]["SupportNetworkScan"] ? true : false);
        }
        A.store.Bool(ptr + 667, "Connectable" in x ? true : false);
        A.store.Bool(ptr + 128, x["Connectable"] ? true : false);
        A.store.Enum(ptr + 132, ["Connected", "Connecting", "NotConnected"].indexOf(x["ConnectionState"] as string));
        A.store.Ref(ptr + 136, x["ErrorState"]);

        if (typeof x["Ethernet"] === "undefined") {
          A.store.Bool(ptr + 140 + 140, false);
          A.store.Bool(ptr + 140 + 139, false);
          A.store.Bool(ptr + 140 + 0, false);
          A.store.Ref(ptr + 140 + 4, undefined);

          A.store.Bool(ptr + 140 + 8 + 130, false);
          A.store.Ref(ptr + 140 + 8 + 0, undefined);

          A.store.Bool(ptr + 140 + 8 + 4 + 45, false);
          A.store.Ref(ptr + 140 + 8 + 4 + 0, undefined);

          A.store.Bool(ptr + 140 + 8 + 4 + 4 + 16, false);
          A.store.Ref(ptr + 140 + 8 + 4 + 4 + 0, undefined);
          A.store.Ref(ptr + 140 + 8 + 4 + 4 + 4, undefined);
          A.store.Ref(ptr + 140 + 8 + 4 + 4 + 8, undefined);
          A.store.Ref(ptr + 140 + 8 + 4 + 4 + 12, undefined);
          A.store.Ref(ptr + 140 + 8 + 4 + 24, undefined);

          A.store.Bool(ptr + 140 + 8 + 4 + 28 + 16, false);
          A.store.Ref(ptr + 140 + 8 + 4 + 28 + 0, undefined);
          A.store.Ref(ptr + 140 + 8 + 4 + 28 + 4, undefined);
          A.store.Ref(ptr + 140 + 8 + 4 + 28 + 8, undefined);
          A.store.Ref(ptr + 140 + 8 + 4 + 28 + 12, undefined);
          A.store.Ref(ptr + 140 + 8 + 52, undefined);
          A.store.Ref(ptr + 140 + 8 + 56, undefined);
          A.store.Ref(ptr + 140 + 8 + 60, undefined);
          A.store.Enum(ptr + 140 + 8 + 64, -1);
          A.store.Ref(ptr + 140 + 8 + 68, undefined);
          A.store.Ref(ptr + 140 + 8 + 72, undefined);
          A.store.Ref(ptr + 140 + 8 + 76, undefined);
          A.store.Ref(ptr + 140 + 8 + 80, undefined);
          A.store.Bool(ptr + 140 + 8 + 127, false);
          A.store.Bool(ptr + 140 + 8 + 84, false);
          A.store.Ref(ptr + 140 + 8 + 88, undefined);
          A.store.Ref(ptr + 140 + 8 + 92, undefined);

          A.store.Bool(ptr + 140 + 8 + 96 + 28, false);
          A.store.Ref(ptr + 140 + 8 + 96 + 0, undefined);
          A.store.Ref(ptr + 140 + 8 + 96 + 4, undefined);
          A.store.Ref(ptr + 140 + 8 + 96 + 8, undefined);
          A.store.Ref(ptr + 140 + 8 + 96 + 12, undefined);
          A.store.Ref(ptr + 140 + 8 + 96 + 16, undefined);
          A.store.Ref(ptr + 140 + 8 + 96 + 20, undefined);
          A.store.Bool(ptr + 140 + 8 + 96 + 26, false);
          A.store.Bool(ptr + 140 + 8 + 96 + 24, false);
          A.store.Bool(ptr + 140 + 8 + 96 + 27, false);
          A.store.Bool(ptr + 140 + 8 + 96 + 25, false);
          A.store.Bool(ptr + 140 + 8 + 128, false);
          A.store.Bool(ptr + 140 + 8 + 125, false);
          A.store.Bool(ptr + 140 + 8 + 129, false);
          A.store.Bool(ptr + 140 + 8 + 126, false);
        } else {
          A.store.Bool(ptr + 140 + 140, true);
          A.store.Bool(ptr + 140 + 139, "AutoConnect" in x["Ethernet"] ? true : false);
          A.store.Bool(ptr + 140 + 0, x["Ethernet"]["AutoConnect"] ? true : false);
          A.store.Ref(ptr + 140 + 4, x["Ethernet"]["Authentication"]);

          if (typeof x["Ethernet"]["EAP"] === "undefined") {
            A.store.Bool(ptr + 140 + 8 + 130, false);
            A.store.Ref(ptr + 140 + 8 + 0, undefined);

            A.store.Bool(ptr + 140 + 8 + 4 + 45, false);
            A.store.Ref(ptr + 140 + 8 + 4 + 0, undefined);

            A.store.Bool(ptr + 140 + 8 + 4 + 4 + 16, false);
            A.store.Ref(ptr + 140 + 8 + 4 + 4 + 0, undefined);
            A.store.Ref(ptr + 140 + 8 + 4 + 4 + 4, undefined);
            A.store.Ref(ptr + 140 + 8 + 4 + 4 + 8, undefined);
            A.store.Ref(ptr + 140 + 8 + 4 + 4 + 12, undefined);
            A.store.Ref(ptr + 140 + 8 + 4 + 24, undefined);

            A.store.Bool(ptr + 140 + 8 + 4 + 28 + 16, false);
            A.store.Ref(ptr + 140 + 8 + 4 + 28 + 0, undefined);
            A.store.Ref(ptr + 140 + 8 + 4 + 28 + 4, undefined);
            A.store.Ref(ptr + 140 + 8 + 4 + 28 + 8, undefined);
            A.store.Ref(ptr + 140 + 8 + 4 + 28 + 12, undefined);
            A.store.Ref(ptr + 140 + 8 + 52, undefined);
            A.store.Ref(ptr + 140 + 8 + 56, undefined);
            A.store.Ref(ptr + 140 + 8 + 60, undefined);
            A.store.Enum(ptr + 140 + 8 + 64, -1);
            A.store.Ref(ptr + 140 + 8 + 68, undefined);
            A.store.Ref(ptr + 140 + 8 + 72, undefined);
            A.store.Ref(ptr + 140 + 8 + 76, undefined);
            A.store.Ref(ptr + 140 + 8 + 80, undefined);
            A.store.Bool(ptr + 140 + 8 + 127, false);
            A.store.Bool(ptr + 140 + 8 + 84, false);
            A.store.Ref(ptr + 140 + 8 + 88, undefined);
            A.store.Ref(ptr + 140 + 8 + 92, undefined);

            A.store.Bool(ptr + 140 + 8 + 96 + 28, false);
            A.store.Ref(ptr + 140 + 8 + 96 + 0, undefined);
            A.store.Ref(ptr + 140 + 8 + 96 + 4, undefined);
            A.store.Ref(ptr + 140 + 8 + 96 + 8, undefined);
            A.store.Ref(ptr + 140 + 8 + 96 + 12, undefined);
            A.store.Ref(ptr + 140 + 8 + 96 + 16, undefined);
            A.store.Ref(ptr + 140 + 8 + 96 + 20, undefined);
            A.store.Bool(ptr + 140 + 8 + 96 + 26, false);
            A.store.Bool(ptr + 140 + 8 + 96 + 24, false);
            A.store.Bool(ptr + 140 + 8 + 96 + 27, false);
            A.store.Bool(ptr + 140 + 8 + 96 + 25, false);
            A.store.Bool(ptr + 140 + 8 + 128, false);
            A.store.Bool(ptr + 140 + 8 + 125, false);
            A.store.Bool(ptr + 140 + 8 + 129, false);
            A.store.Bool(ptr + 140 + 8 + 126, false);
          } else {
            A.store.Bool(ptr + 140 + 8 + 130, true);
            A.store.Ref(ptr + 140 + 8 + 0, x["Ethernet"]["EAP"]["AnonymousIdentity"]);

            if (typeof x["Ethernet"]["EAP"]["ClientCertPattern"] === "undefined") {
              A.store.Bool(ptr + 140 + 8 + 4 + 45, false);
              A.store.Ref(ptr + 140 + 8 + 4 + 0, undefined);

              A.store.Bool(ptr + 140 + 8 + 4 + 4 + 16, false);
              A.store.Ref(ptr + 140 + 8 + 4 + 4 + 0, undefined);
              A.store.Ref(ptr + 140 + 8 + 4 + 4 + 4, undefined);
              A.store.Ref(ptr + 140 + 8 + 4 + 4 + 8, undefined);
              A.store.Ref(ptr + 140 + 8 + 4 + 4 + 12, undefined);
              A.store.Ref(ptr + 140 + 8 + 4 + 24, undefined);

              A.store.Bool(ptr + 140 + 8 + 4 + 28 + 16, false);
              A.store.Ref(ptr + 140 + 8 + 4 + 28 + 0, undefined);
              A.store.Ref(ptr + 140 + 8 + 4 + 28 + 4, undefined);
              A.store.Ref(ptr + 140 + 8 + 4 + 28 + 8, undefined);
              A.store.Ref(ptr + 140 + 8 + 4 + 28 + 12, undefined);
            } else {
              A.store.Bool(ptr + 140 + 8 + 4 + 45, true);
              A.store.Ref(ptr + 140 + 8 + 4 + 0, x["Ethernet"]["EAP"]["ClientCertPattern"]["EnrollmentURI"]);

              if (typeof x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"] === "undefined") {
                A.store.Bool(ptr + 140 + 8 + 4 + 4 + 16, false);
                A.store.Ref(ptr + 140 + 8 + 4 + 4 + 0, undefined);
                A.store.Ref(ptr + 140 + 8 + 4 + 4 + 4, undefined);
                A.store.Ref(ptr + 140 + 8 + 4 + 4 + 8, undefined);
                A.store.Ref(ptr + 140 + 8 + 4 + 4 + 12, undefined);
              } else {
                A.store.Bool(ptr + 140 + 8 + 4 + 4 + 16, true);
                A.store.Ref(
                  ptr + 140 + 8 + 4 + 4 + 0,
                  x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"]
                );
                A.store.Ref(ptr + 140 + 8 + 4 + 4 + 4, x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"]);
                A.store.Ref(
                  ptr + 140 + 8 + 4 + 4 + 8,
                  x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"]
                );
                A.store.Ref(
                  ptr + 140 + 8 + 4 + 4 + 12,
                  x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"]
                );
              }
              A.store.Ref(ptr + 140 + 8 + 4 + 24, x["Ethernet"]["EAP"]["ClientCertPattern"]["IssuerCARef"]);

              if (typeof x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"] === "undefined") {
                A.store.Bool(ptr + 140 + 8 + 4 + 28 + 16, false);
                A.store.Ref(ptr + 140 + 8 + 4 + 28 + 0, undefined);
                A.store.Ref(ptr + 140 + 8 + 4 + 28 + 4, undefined);
                A.store.Ref(ptr + 140 + 8 + 4 + 28 + 8, undefined);
                A.store.Ref(ptr + 140 + 8 + 4 + 28 + 12, undefined);
              } else {
                A.store.Bool(ptr + 140 + 8 + 4 + 28 + 16, true);
                A.store.Ref(
                  ptr + 140 + 8 + 4 + 28 + 0,
                  x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"]
                );
                A.store.Ref(
                  ptr + 140 + 8 + 4 + 28 + 4,
                  x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"]
                );
                A.store.Ref(
                  ptr + 140 + 8 + 4 + 28 + 8,
                  x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"]
                );
                A.store.Ref(
                  ptr + 140 + 8 + 4 + 28 + 12,
                  x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"]
                );
              }
            }
            A.store.Ref(ptr + 140 + 8 + 52, x["Ethernet"]["EAP"]["ClientCertPKCS11Id"]);
            A.store.Ref(ptr + 140 + 8 + 56, x["Ethernet"]["EAP"]["ClientCertProvisioningProfileId"]);
            A.store.Ref(ptr + 140 + 8 + 60, x["Ethernet"]["EAP"]["ClientCertRef"]);
            A.store.Enum(
              ptr + 140 + 8 + 64,
              ["Ref", "Pattern"].indexOf(x["Ethernet"]["EAP"]["ClientCertType"] as string)
            );
            A.store.Ref(ptr + 140 + 8 + 68, x["Ethernet"]["EAP"]["Identity"]);
            A.store.Ref(ptr + 140 + 8 + 72, x["Ethernet"]["EAP"]["Inner"]);
            A.store.Ref(ptr + 140 + 8 + 76, x["Ethernet"]["EAP"]["Outer"]);
            A.store.Ref(ptr + 140 + 8 + 80, x["Ethernet"]["EAP"]["Password"]);
            A.store.Bool(ptr + 140 + 8 + 127, "SaveCredentials" in x["Ethernet"]["EAP"] ? true : false);
            A.store.Bool(ptr + 140 + 8 + 84, x["Ethernet"]["EAP"]["SaveCredentials"] ? true : false);
            A.store.Ref(ptr + 140 + 8 + 88, x["Ethernet"]["EAP"]["ServerCAPEMs"]);
            A.store.Ref(ptr + 140 + 8 + 92, x["Ethernet"]["EAP"]["ServerCARefs"]);

            if (typeof x["Ethernet"]["EAP"]["SubjectMatch"] === "undefined") {
              A.store.Bool(ptr + 140 + 8 + 96 + 28, false);
              A.store.Ref(ptr + 140 + 8 + 96 + 0, undefined);
              A.store.Ref(ptr + 140 + 8 + 96 + 4, undefined);
              A.store.Ref(ptr + 140 + 8 + 96 + 8, undefined);
              A.store.Ref(ptr + 140 + 8 + 96 + 12, undefined);
              A.store.Ref(ptr + 140 + 8 + 96 + 16, undefined);
              A.store.Ref(ptr + 140 + 8 + 96 + 20, undefined);
              A.store.Bool(ptr + 140 + 8 + 96 + 26, false);
              A.store.Bool(ptr + 140 + 8 + 96 + 24, false);
              A.store.Bool(ptr + 140 + 8 + 96 + 27, false);
              A.store.Bool(ptr + 140 + 8 + 96 + 25, false);
            } else {
              A.store.Bool(ptr + 140 + 8 + 96 + 28, true);
              A.store.Ref(ptr + 140 + 8 + 96 + 0, x["Ethernet"]["EAP"]["SubjectMatch"]["Active"]);
              A.store.Ref(ptr + 140 + 8 + 96 + 4, x["Ethernet"]["EAP"]["SubjectMatch"]["Effective"]);
              A.store.Ref(ptr + 140 + 8 + 96 + 8, x["Ethernet"]["EAP"]["SubjectMatch"]["UserPolicy"]);
              A.store.Ref(ptr + 140 + 8 + 96 + 12, x["Ethernet"]["EAP"]["SubjectMatch"]["DevicePolicy"]);
              A.store.Ref(ptr + 140 + 8 + 96 + 16, x["Ethernet"]["EAP"]["SubjectMatch"]["UserSetting"]);
              A.store.Ref(ptr + 140 + 8 + 96 + 20, x["Ethernet"]["EAP"]["SubjectMatch"]["SharedSetting"]);
              A.store.Bool(
                ptr + 140 + 8 + 96 + 26,
                "UserEditable" in x["Ethernet"]["EAP"]["SubjectMatch"] ? true : false
              );
              A.store.Bool(
                ptr + 140 + 8 + 96 + 24,
                x["Ethernet"]["EAP"]["SubjectMatch"]["UserEditable"] ? true : false
              );
              A.store.Bool(
                ptr + 140 + 8 + 96 + 27,
                "DeviceEditable" in x["Ethernet"]["EAP"]["SubjectMatch"] ? true : false
              );
              A.store.Bool(
                ptr + 140 + 8 + 96 + 25,
                x["Ethernet"]["EAP"]["SubjectMatch"]["DeviceEditable"] ? true : false
              );
            }
            A.store.Bool(ptr + 140 + 8 + 128, "UseProactiveKeyCaching" in x["Ethernet"]["EAP"] ? true : false);
            A.store.Bool(ptr + 140 + 8 + 125, x["Ethernet"]["EAP"]["UseProactiveKeyCaching"] ? true : false);
            A.store.Bool(ptr + 140 + 8 + 129, "UseSystemCAs" in x["Ethernet"]["EAP"] ? true : false);
            A.store.Bool(ptr + 140 + 8 + 126, x["Ethernet"]["EAP"]["UseSystemCAs"] ? true : false);
          }
        }
        A.store.Ref(ptr + 284, x["GUID"]);
        A.store.Enum(ptr + 288, ["DHCP", "Static"].indexOf(x["IPAddressConfigType"] as string));
        A.store.Ref(ptr + 292, x["IPConfigs"]);
        A.store.Ref(ptr + 296, x["MacAddress"]);
        A.store.Bool(ptr + 668, "Metered" in x ? true : false);
        A.store.Bool(ptr + 300, x["Metered"] ? true : false);
        A.store.Ref(ptr + 304, x["Name"]);
        A.store.Enum(ptr + 308, ["DHCP", "Static"].indexOf(x["NameServersConfigType"] as string));
        A.store.Bool(ptr + 669, "Priority" in x ? true : false);
        A.store.Int32(ptr + 312, x["Priority"] === undefined ? 0 : (x["Priority"] as number));

        if (typeof x["ProxySettings"] === "undefined") {
          A.store.Bool(ptr + 316 + 60, false);
          A.store.Enum(ptr + 316 + 0, -1);

          A.store.Bool(ptr + 316 + 4 + 46, false);

          A.store.Bool(ptr + 316 + 4 + 0 + 9, false);
          A.store.Ref(ptr + 316 + 4 + 0 + 0, undefined);
          A.store.Bool(ptr + 316 + 4 + 0 + 8, false);
          A.store.Int32(ptr + 316 + 4 + 0 + 4, 0);

          A.store.Bool(ptr + 316 + 4 + 12 + 9, false);
          A.store.Ref(ptr + 316 + 4 + 12 + 0, undefined);
          A.store.Bool(ptr + 316 + 4 + 12 + 8, false);
          A.store.Int32(ptr + 316 + 4 + 12 + 4, 0);

          A.store.Bool(ptr + 316 + 4 + 24 + 9, false);
          A.store.Ref(ptr + 316 + 4 + 24 + 0, undefined);
          A.store.Bool(ptr + 316 + 4 + 24 + 8, false);
          A.store.Int32(ptr + 316 + 4 + 24 + 4, 0);

          A.store.Bool(ptr + 316 + 4 + 36 + 9, false);
          A.store.Ref(ptr + 316 + 4 + 36 + 0, undefined);
          A.store.Bool(ptr + 316 + 4 + 36 + 8, false);
          A.store.Int32(ptr + 316 + 4 + 36 + 4, 0);
          A.store.Ref(ptr + 316 + 52, undefined);
          A.store.Ref(ptr + 316 + 56, undefined);
        } else {
          A.store.Bool(ptr + 316 + 60, true);
          A.store.Enum(
            ptr + 316 + 0,
            ["Direct", "Manual", "PAC", "WPAD"].indexOf(x["ProxySettings"]["Type"] as string)
          );

          if (typeof x["ProxySettings"]["Manual"] === "undefined") {
            A.store.Bool(ptr + 316 + 4 + 46, false);

            A.store.Bool(ptr + 316 + 4 + 0 + 9, false);
            A.store.Ref(ptr + 316 + 4 + 0 + 0, undefined);
            A.store.Bool(ptr + 316 + 4 + 0 + 8, false);
            A.store.Int32(ptr + 316 + 4 + 0 + 4, 0);

            A.store.Bool(ptr + 316 + 4 + 12 + 9, false);
            A.store.Ref(ptr + 316 + 4 + 12 + 0, undefined);
            A.store.Bool(ptr + 316 + 4 + 12 + 8, false);
            A.store.Int32(ptr + 316 + 4 + 12 + 4, 0);

            A.store.Bool(ptr + 316 + 4 + 24 + 9, false);
            A.store.Ref(ptr + 316 + 4 + 24 + 0, undefined);
            A.store.Bool(ptr + 316 + 4 + 24 + 8, false);
            A.store.Int32(ptr + 316 + 4 + 24 + 4, 0);

            A.store.Bool(ptr + 316 + 4 + 36 + 9, false);
            A.store.Ref(ptr + 316 + 4 + 36 + 0, undefined);
            A.store.Bool(ptr + 316 + 4 + 36 + 8, false);
            A.store.Int32(ptr + 316 + 4 + 36 + 4, 0);
          } else {
            A.store.Bool(ptr + 316 + 4 + 46, true);

            if (typeof x["ProxySettings"]["Manual"]["HTTPProxy"] === "undefined") {
              A.store.Bool(ptr + 316 + 4 + 0 + 9, false);
              A.store.Ref(ptr + 316 + 4 + 0 + 0, undefined);
              A.store.Bool(ptr + 316 + 4 + 0 + 8, false);
              A.store.Int32(ptr + 316 + 4 + 0 + 4, 0);
            } else {
              A.store.Bool(ptr + 316 + 4 + 0 + 9, true);
              A.store.Ref(ptr + 316 + 4 + 0 + 0, x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"]);
              A.store.Bool(ptr + 316 + 4 + 0 + 8, "Port" in x["ProxySettings"]["Manual"]["HTTPProxy"] ? true : false);
              A.store.Int32(
                ptr + 316 + 4 + 0 + 4,
                x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"] === undefined
                  ? 0
                  : (x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"] as number)
              );
            }

            if (typeof x["ProxySettings"]["Manual"]["SecureHTTPProxy"] === "undefined") {
              A.store.Bool(ptr + 316 + 4 + 12 + 9, false);
              A.store.Ref(ptr + 316 + 4 + 12 + 0, undefined);
              A.store.Bool(ptr + 316 + 4 + 12 + 8, false);
              A.store.Int32(ptr + 316 + 4 + 12 + 4, 0);
            } else {
              A.store.Bool(ptr + 316 + 4 + 12 + 9, true);
              A.store.Ref(ptr + 316 + 4 + 12 + 0, x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"]);
              A.store.Bool(
                ptr + 316 + 4 + 12 + 8,
                "Port" in x["ProxySettings"]["Manual"]["SecureHTTPProxy"] ? true : false
              );
              A.store.Int32(
                ptr + 316 + 4 + 12 + 4,
                x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"] === undefined
                  ? 0
                  : (x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"] as number)
              );
            }

            if (typeof x["ProxySettings"]["Manual"]["FTPProxy"] === "undefined") {
              A.store.Bool(ptr + 316 + 4 + 24 + 9, false);
              A.store.Ref(ptr + 316 + 4 + 24 + 0, undefined);
              A.store.Bool(ptr + 316 + 4 + 24 + 8, false);
              A.store.Int32(ptr + 316 + 4 + 24 + 4, 0);
            } else {
              A.store.Bool(ptr + 316 + 4 + 24 + 9, true);
              A.store.Ref(ptr + 316 + 4 + 24 + 0, x["ProxySettings"]["Manual"]["FTPProxy"]["Host"]);
              A.store.Bool(ptr + 316 + 4 + 24 + 8, "Port" in x["ProxySettings"]["Manual"]["FTPProxy"] ? true : false);
              A.store.Int32(
                ptr + 316 + 4 + 24 + 4,
                x["ProxySettings"]["Manual"]["FTPProxy"]["Port"] === undefined
                  ? 0
                  : (x["ProxySettings"]["Manual"]["FTPProxy"]["Port"] as number)
              );
            }

            if (typeof x["ProxySettings"]["Manual"]["SOCKS"] === "undefined") {
              A.store.Bool(ptr + 316 + 4 + 36 + 9, false);
              A.store.Ref(ptr + 316 + 4 + 36 + 0, undefined);
              A.store.Bool(ptr + 316 + 4 + 36 + 8, false);
              A.store.Int32(ptr + 316 + 4 + 36 + 4, 0);
            } else {
              A.store.Bool(ptr + 316 + 4 + 36 + 9, true);
              A.store.Ref(ptr + 316 + 4 + 36 + 0, x["ProxySettings"]["Manual"]["SOCKS"]["Host"]);
              A.store.Bool(ptr + 316 + 4 + 36 + 8, "Port" in x["ProxySettings"]["Manual"]["SOCKS"] ? true : false);
              A.store.Int32(
                ptr + 316 + 4 + 36 + 4,
                x["ProxySettings"]["Manual"]["SOCKS"]["Port"] === undefined
                  ? 0
                  : (x["ProxySettings"]["Manual"]["SOCKS"]["Port"] as number)
              );
            }
          }
          A.store.Ref(ptr + 316 + 52, x["ProxySettings"]["ExcludeDomains"]);
          A.store.Ref(ptr + 316 + 56, x["ProxySettings"]["PAC"]);
        }
        A.store.Bool(ptr + 670, "RestrictedConnectivity" in x ? true : false);
        A.store.Bool(ptr + 377, x["RestrictedConnectivity"] ? true : false);

        if (typeof x["StaticIPConfig"] === "undefined") {
          A.store.Bool(ptr + 380 + 37, false);
          A.store.Ref(ptr + 380 + 0, undefined);
          A.store.Ref(ptr + 380 + 4, undefined);
          A.store.Ref(ptr + 380 + 8, undefined);
          A.store.Ref(ptr + 380 + 12, undefined);
          A.store.Ref(ptr + 380 + 16, undefined);
          A.store.Ref(ptr + 380 + 20, undefined);
          A.store.Bool(ptr + 380 + 36, false);
          A.store.Int32(ptr + 380 + 24, 0);
          A.store.Ref(ptr + 380 + 28, undefined);
          A.store.Ref(ptr + 380 + 32, undefined);
        } else {
          A.store.Bool(ptr + 380 + 37, true);
          A.store.Ref(ptr + 380 + 0, x["StaticIPConfig"]["Gateway"]);
          A.store.Ref(ptr + 380 + 4, x["StaticIPConfig"]["IPAddress"]);
          A.store.Ref(ptr + 380 + 8, x["StaticIPConfig"]["ExcludedRoutes"]);
          A.store.Ref(ptr + 380 + 12, x["StaticIPConfig"]["IncludedRoutes"]);
          A.store.Ref(ptr + 380 + 16, x["StaticIPConfig"]["NameServers"]);
          A.store.Ref(ptr + 380 + 20, x["StaticIPConfig"]["SearchDomains"]);
          A.store.Bool(ptr + 380 + 36, "RoutingPrefix" in x["StaticIPConfig"] ? true : false);
          A.store.Int32(
            ptr + 380 + 24,
            x["StaticIPConfig"]["RoutingPrefix"] === undefined ? 0 : (x["StaticIPConfig"]["RoutingPrefix"] as number)
          );
          A.store.Ref(ptr + 380 + 28, x["StaticIPConfig"]["Type"]);
          A.store.Ref(ptr + 380 + 32, x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"]);
        }

        if (typeof x["SavedIPConfig"] === "undefined") {
          A.store.Bool(ptr + 420 + 37, false);
          A.store.Ref(ptr + 420 + 0, undefined);
          A.store.Ref(ptr + 420 + 4, undefined);
          A.store.Ref(ptr + 420 + 8, undefined);
          A.store.Ref(ptr + 420 + 12, undefined);
          A.store.Ref(ptr + 420 + 16, undefined);
          A.store.Ref(ptr + 420 + 20, undefined);
          A.store.Bool(ptr + 420 + 36, false);
          A.store.Int32(ptr + 420 + 24, 0);
          A.store.Ref(ptr + 420 + 28, undefined);
          A.store.Ref(ptr + 420 + 32, undefined);
        } else {
          A.store.Bool(ptr + 420 + 37, true);
          A.store.Ref(ptr + 420 + 0, x["SavedIPConfig"]["Gateway"]);
          A.store.Ref(ptr + 420 + 4, x["SavedIPConfig"]["IPAddress"]);
          A.store.Ref(ptr + 420 + 8, x["SavedIPConfig"]["ExcludedRoutes"]);
          A.store.Ref(ptr + 420 + 12, x["SavedIPConfig"]["IncludedRoutes"]);
          A.store.Ref(ptr + 420 + 16, x["SavedIPConfig"]["NameServers"]);
          A.store.Ref(ptr + 420 + 20, x["SavedIPConfig"]["SearchDomains"]);
          A.store.Bool(ptr + 420 + 36, "RoutingPrefix" in x["SavedIPConfig"] ? true : false);
          A.store.Int32(
            ptr + 420 + 24,
            x["SavedIPConfig"]["RoutingPrefix"] === undefined ? 0 : (x["SavedIPConfig"]["RoutingPrefix"] as number)
          );
          A.store.Ref(ptr + 420 + 28, x["SavedIPConfig"]["Type"]);
          A.store.Ref(ptr + 420 + 32, x["SavedIPConfig"]["WebProxyAutoDiscoveryUrl"]);
        }
        A.store.Ref(ptr + 460, x["Source"]);
        A.store.Enum(
          ptr + 464,
          ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"].indexOf(x["Type"] as string)
        );

        if (typeof x["VPN"] === "undefined") {
          A.store.Bool(ptr + 468 + 13, false);
          A.store.Bool(ptr + 468 + 12, false);
          A.store.Bool(ptr + 468 + 0, false);
          A.store.Ref(ptr + 468 + 4, undefined);
          A.store.Ref(ptr + 468 + 8, undefined);
        } else {
          A.store.Bool(ptr + 468 + 13, true);
          A.store.Bool(ptr + 468 + 12, "AutoConnect" in x["VPN"] ? true : false);
          A.store.Bool(ptr + 468 + 0, x["VPN"]["AutoConnect"] ? true : false);
          A.store.Ref(ptr + 468 + 4, x["VPN"]["Host"]);
          A.store.Ref(ptr + 468 + 8, x["VPN"]["Type"]);
        }

        if (typeof x["WiFi"] === "undefined") {
          A.store.Bool(ptr + 484 + 182, false);
          A.store.Bool(ptr + 484 + 176, false);
          A.store.Bool(ptr + 484 + 0, false);
          A.store.Bool(ptr + 484 + 177, false);
          A.store.Bool(ptr + 484 + 1, false);
          A.store.Ref(ptr + 484 + 4, undefined);

          A.store.Bool(ptr + 484 + 8 + 130, false);
          A.store.Ref(ptr + 484 + 8 + 0, undefined);

          A.store.Bool(ptr + 484 + 8 + 4 + 45, false);
          A.store.Ref(ptr + 484 + 8 + 4 + 0, undefined);

          A.store.Bool(ptr + 484 + 8 + 4 + 4 + 16, false);
          A.store.Ref(ptr + 484 + 8 + 4 + 4 + 0, undefined);
          A.store.Ref(ptr + 484 + 8 + 4 + 4 + 4, undefined);
          A.store.Ref(ptr + 484 + 8 + 4 + 4 + 8, undefined);
          A.store.Ref(ptr + 484 + 8 + 4 + 4 + 12, undefined);
          A.store.Ref(ptr + 484 + 8 + 4 + 24, undefined);

          A.store.Bool(ptr + 484 + 8 + 4 + 28 + 16, false);
          A.store.Ref(ptr + 484 + 8 + 4 + 28 + 0, undefined);
          A.store.Ref(ptr + 484 + 8 + 4 + 28 + 4, undefined);
          A.store.Ref(ptr + 484 + 8 + 4 + 28 + 8, undefined);
          A.store.Ref(ptr + 484 + 8 + 4 + 28 + 12, undefined);
          A.store.Ref(ptr + 484 + 8 + 52, undefined);
          A.store.Ref(ptr + 484 + 8 + 56, undefined);
          A.store.Ref(ptr + 484 + 8 + 60, undefined);
          A.store.Enum(ptr + 484 + 8 + 64, -1);
          A.store.Ref(ptr + 484 + 8 + 68, undefined);
          A.store.Ref(ptr + 484 + 8 + 72, undefined);
          A.store.Ref(ptr + 484 + 8 + 76, undefined);
          A.store.Ref(ptr + 484 + 8 + 80, undefined);
          A.store.Bool(ptr + 484 + 8 + 127, false);
          A.store.Bool(ptr + 484 + 8 + 84, false);
          A.store.Ref(ptr + 484 + 8 + 88, undefined);
          A.store.Ref(ptr + 484 + 8 + 92, undefined);

          A.store.Bool(ptr + 484 + 8 + 96 + 28, false);
          A.store.Ref(ptr + 484 + 8 + 96 + 0, undefined);
          A.store.Ref(ptr + 484 + 8 + 96 + 4, undefined);
          A.store.Ref(ptr + 484 + 8 + 96 + 8, undefined);
          A.store.Ref(ptr + 484 + 8 + 96 + 12, undefined);
          A.store.Ref(ptr + 484 + 8 + 96 + 16, undefined);
          A.store.Ref(ptr + 484 + 8 + 96 + 20, undefined);
          A.store.Bool(ptr + 484 + 8 + 96 + 26, false);
          A.store.Bool(ptr + 484 + 8 + 96 + 24, false);
          A.store.Bool(ptr + 484 + 8 + 96 + 27, false);
          A.store.Bool(ptr + 484 + 8 + 96 + 25, false);
          A.store.Bool(ptr + 484 + 8 + 128, false);
          A.store.Bool(ptr + 484 + 8 + 125, false);
          A.store.Bool(ptr + 484 + 8 + 129, false);
          A.store.Bool(ptr + 484 + 8 + 126, false);
          A.store.Bool(ptr + 484 + 178, false);
          A.store.Int32(ptr + 484 + 140, 0);
          A.store.Ref(ptr + 484 + 144, undefined);
          A.store.Ref(ptr + 484 + 148, undefined);
          A.store.Bool(ptr + 484 + 179, false);
          A.store.Bool(ptr + 484 + 152, false);
          A.store.Ref(ptr + 484 + 156, undefined);
          A.store.Bool(ptr + 484 + 180, false);
          A.store.Int32(ptr + 484 + 160, 0);
          A.store.Ref(ptr + 484 + 164, undefined);
          A.store.Ref(ptr + 484 + 168, undefined);
          A.store.Bool(ptr + 484 + 181, false);
          A.store.Int32(ptr + 484 + 172, 0);
        } else {
          A.store.Bool(ptr + 484 + 182, true);
          A.store.Bool(ptr + 484 + 176, "AllowGatewayARPPolling" in x["WiFi"] ? true : false);
          A.store.Bool(ptr + 484 + 0, x["WiFi"]["AllowGatewayARPPolling"] ? true : false);
          A.store.Bool(ptr + 484 + 177, "AutoConnect" in x["WiFi"] ? true : false);
          A.store.Bool(ptr + 484 + 1, x["WiFi"]["AutoConnect"] ? true : false);
          A.store.Ref(ptr + 484 + 4, x["WiFi"]["BSSID"]);

          if (typeof x["WiFi"]["EAP"] === "undefined") {
            A.store.Bool(ptr + 484 + 8 + 130, false);
            A.store.Ref(ptr + 484 + 8 + 0, undefined);

            A.store.Bool(ptr + 484 + 8 + 4 + 45, false);
            A.store.Ref(ptr + 484 + 8 + 4 + 0, undefined);

            A.store.Bool(ptr + 484 + 8 + 4 + 4 + 16, false);
            A.store.Ref(ptr + 484 + 8 + 4 + 4 + 0, undefined);
            A.store.Ref(ptr + 484 + 8 + 4 + 4 + 4, undefined);
            A.store.Ref(ptr + 484 + 8 + 4 + 4 + 8, undefined);
            A.store.Ref(ptr + 484 + 8 + 4 + 4 + 12, undefined);
            A.store.Ref(ptr + 484 + 8 + 4 + 24, undefined);

            A.store.Bool(ptr + 484 + 8 + 4 + 28 + 16, false);
            A.store.Ref(ptr + 484 + 8 + 4 + 28 + 0, undefined);
            A.store.Ref(ptr + 484 + 8 + 4 + 28 + 4, undefined);
            A.store.Ref(ptr + 484 + 8 + 4 + 28 + 8, undefined);
            A.store.Ref(ptr + 484 + 8 + 4 + 28 + 12, undefined);
            A.store.Ref(ptr + 484 + 8 + 52, undefined);
            A.store.Ref(ptr + 484 + 8 + 56, undefined);
            A.store.Ref(ptr + 484 + 8 + 60, undefined);
            A.store.Enum(ptr + 484 + 8 + 64, -1);
            A.store.Ref(ptr + 484 + 8 + 68, undefined);
            A.store.Ref(ptr + 484 + 8 + 72, undefined);
            A.store.Ref(ptr + 484 + 8 + 76, undefined);
            A.store.Ref(ptr + 484 + 8 + 80, undefined);
            A.store.Bool(ptr + 484 + 8 + 127, false);
            A.store.Bool(ptr + 484 + 8 + 84, false);
            A.store.Ref(ptr + 484 + 8 + 88, undefined);
            A.store.Ref(ptr + 484 + 8 + 92, undefined);

            A.store.Bool(ptr + 484 + 8 + 96 + 28, false);
            A.store.Ref(ptr + 484 + 8 + 96 + 0, undefined);
            A.store.Ref(ptr + 484 + 8 + 96 + 4, undefined);
            A.store.Ref(ptr + 484 + 8 + 96 + 8, undefined);
            A.store.Ref(ptr + 484 + 8 + 96 + 12, undefined);
            A.store.Ref(ptr + 484 + 8 + 96 + 16, undefined);
            A.store.Ref(ptr + 484 + 8 + 96 + 20, undefined);
            A.store.Bool(ptr + 484 + 8 + 96 + 26, false);
            A.store.Bool(ptr + 484 + 8 + 96 + 24, false);
            A.store.Bool(ptr + 484 + 8 + 96 + 27, false);
            A.store.Bool(ptr + 484 + 8 + 96 + 25, false);
            A.store.Bool(ptr + 484 + 8 + 128, false);
            A.store.Bool(ptr + 484 + 8 + 125, false);
            A.store.Bool(ptr + 484 + 8 + 129, false);
            A.store.Bool(ptr + 484 + 8 + 126, false);
          } else {
            A.store.Bool(ptr + 484 + 8 + 130, true);
            A.store.Ref(ptr + 484 + 8 + 0, x["WiFi"]["EAP"]["AnonymousIdentity"]);

            if (typeof x["WiFi"]["EAP"]["ClientCertPattern"] === "undefined") {
              A.store.Bool(ptr + 484 + 8 + 4 + 45, false);
              A.store.Ref(ptr + 484 + 8 + 4 + 0, undefined);

              A.store.Bool(ptr + 484 + 8 + 4 + 4 + 16, false);
              A.store.Ref(ptr + 484 + 8 + 4 + 4 + 0, undefined);
              A.store.Ref(ptr + 484 + 8 + 4 + 4 + 4, undefined);
              A.store.Ref(ptr + 484 + 8 + 4 + 4 + 8, undefined);
              A.store.Ref(ptr + 484 + 8 + 4 + 4 + 12, undefined);
              A.store.Ref(ptr + 484 + 8 + 4 + 24, undefined);

              A.store.Bool(ptr + 484 + 8 + 4 + 28 + 16, false);
              A.store.Ref(ptr + 484 + 8 + 4 + 28 + 0, undefined);
              A.store.Ref(ptr + 484 + 8 + 4 + 28 + 4, undefined);
              A.store.Ref(ptr + 484 + 8 + 4 + 28 + 8, undefined);
              A.store.Ref(ptr + 484 + 8 + 4 + 28 + 12, undefined);
            } else {
              A.store.Bool(ptr + 484 + 8 + 4 + 45, true);
              A.store.Ref(ptr + 484 + 8 + 4 + 0, x["WiFi"]["EAP"]["ClientCertPattern"]["EnrollmentURI"]);

              if (typeof x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"] === "undefined") {
                A.store.Bool(ptr + 484 + 8 + 4 + 4 + 16, false);
                A.store.Ref(ptr + 484 + 8 + 4 + 4 + 0, undefined);
                A.store.Ref(ptr + 484 + 8 + 4 + 4 + 4, undefined);
                A.store.Ref(ptr + 484 + 8 + 4 + 4 + 8, undefined);
                A.store.Ref(ptr + 484 + 8 + 4 + 4 + 12, undefined);
              } else {
                A.store.Bool(ptr + 484 + 8 + 4 + 4 + 16, true);
                A.store.Ref(ptr + 484 + 8 + 4 + 4 + 0, x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"]);
                A.store.Ref(ptr + 484 + 8 + 4 + 4 + 4, x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"]);
                A.store.Ref(ptr + 484 + 8 + 4 + 4 + 8, x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"]);
                A.store.Ref(
                  ptr + 484 + 8 + 4 + 4 + 12,
                  x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"]
                );
              }
              A.store.Ref(ptr + 484 + 8 + 4 + 24, x["WiFi"]["EAP"]["ClientCertPattern"]["IssuerCARef"]);

              if (typeof x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"] === "undefined") {
                A.store.Bool(ptr + 484 + 8 + 4 + 28 + 16, false);
                A.store.Ref(ptr + 484 + 8 + 4 + 28 + 0, undefined);
                A.store.Ref(ptr + 484 + 8 + 4 + 28 + 4, undefined);
                A.store.Ref(ptr + 484 + 8 + 4 + 28 + 8, undefined);
                A.store.Ref(ptr + 484 + 8 + 4 + 28 + 12, undefined);
              } else {
                A.store.Bool(ptr + 484 + 8 + 4 + 28 + 16, true);
                A.store.Ref(ptr + 484 + 8 + 4 + 28 + 0, x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"]);
                A.store.Ref(ptr + 484 + 8 + 4 + 28 + 4, x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"]);
                A.store.Ref(
                  ptr + 484 + 8 + 4 + 28 + 8,
                  x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"]
                );
                A.store.Ref(
                  ptr + 484 + 8 + 4 + 28 + 12,
                  x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"]
                );
              }
            }
            A.store.Ref(ptr + 484 + 8 + 52, x["WiFi"]["EAP"]["ClientCertPKCS11Id"]);
            A.store.Ref(ptr + 484 + 8 + 56, x["WiFi"]["EAP"]["ClientCertProvisioningProfileId"]);
            A.store.Ref(ptr + 484 + 8 + 60, x["WiFi"]["EAP"]["ClientCertRef"]);
            A.store.Enum(ptr + 484 + 8 + 64, ["Ref", "Pattern"].indexOf(x["WiFi"]["EAP"]["ClientCertType"] as string));
            A.store.Ref(ptr + 484 + 8 + 68, x["WiFi"]["EAP"]["Identity"]);
            A.store.Ref(ptr + 484 + 8 + 72, x["WiFi"]["EAP"]["Inner"]);
            A.store.Ref(ptr + 484 + 8 + 76, x["WiFi"]["EAP"]["Outer"]);
            A.store.Ref(ptr + 484 + 8 + 80, x["WiFi"]["EAP"]["Password"]);
            A.store.Bool(ptr + 484 + 8 + 127, "SaveCredentials" in x["WiFi"]["EAP"] ? true : false);
            A.store.Bool(ptr + 484 + 8 + 84, x["WiFi"]["EAP"]["SaveCredentials"] ? true : false);
            A.store.Ref(ptr + 484 + 8 + 88, x["WiFi"]["EAP"]["ServerCAPEMs"]);
            A.store.Ref(ptr + 484 + 8 + 92, x["WiFi"]["EAP"]["ServerCARefs"]);

            if (typeof x["WiFi"]["EAP"]["SubjectMatch"] === "undefined") {
              A.store.Bool(ptr + 484 + 8 + 96 + 28, false);
              A.store.Ref(ptr + 484 + 8 + 96 + 0, undefined);
              A.store.Ref(ptr + 484 + 8 + 96 + 4, undefined);
              A.store.Ref(ptr + 484 + 8 + 96 + 8, undefined);
              A.store.Ref(ptr + 484 + 8 + 96 + 12, undefined);
              A.store.Ref(ptr + 484 + 8 + 96 + 16, undefined);
              A.store.Ref(ptr + 484 + 8 + 96 + 20, undefined);
              A.store.Bool(ptr + 484 + 8 + 96 + 26, false);
              A.store.Bool(ptr + 484 + 8 + 96 + 24, false);
              A.store.Bool(ptr + 484 + 8 + 96 + 27, false);
              A.store.Bool(ptr + 484 + 8 + 96 + 25, false);
            } else {
              A.store.Bool(ptr + 484 + 8 + 96 + 28, true);
              A.store.Ref(ptr + 484 + 8 + 96 + 0, x["WiFi"]["EAP"]["SubjectMatch"]["Active"]);
              A.store.Ref(ptr + 484 + 8 + 96 + 4, x["WiFi"]["EAP"]["SubjectMatch"]["Effective"]);
              A.store.Ref(ptr + 484 + 8 + 96 + 8, x["WiFi"]["EAP"]["SubjectMatch"]["UserPolicy"]);
              A.store.Ref(ptr + 484 + 8 + 96 + 12, x["WiFi"]["EAP"]["SubjectMatch"]["DevicePolicy"]);
              A.store.Ref(ptr + 484 + 8 + 96 + 16, x["WiFi"]["EAP"]["SubjectMatch"]["UserSetting"]);
              A.store.Ref(ptr + 484 + 8 + 96 + 20, x["WiFi"]["EAP"]["SubjectMatch"]["SharedSetting"]);
              A.store.Bool(ptr + 484 + 8 + 96 + 26, "UserEditable" in x["WiFi"]["EAP"]["SubjectMatch"] ? true : false);
              A.store.Bool(ptr + 484 + 8 + 96 + 24, x["WiFi"]["EAP"]["SubjectMatch"]["UserEditable"] ? true : false);
              A.store.Bool(
                ptr + 484 + 8 + 96 + 27,
                "DeviceEditable" in x["WiFi"]["EAP"]["SubjectMatch"] ? true : false
              );
              A.store.Bool(ptr + 484 + 8 + 96 + 25, x["WiFi"]["EAP"]["SubjectMatch"]["DeviceEditable"] ? true : false);
            }
            A.store.Bool(ptr + 484 + 8 + 128, "UseProactiveKeyCaching" in x["WiFi"]["EAP"] ? true : false);
            A.store.Bool(ptr + 484 + 8 + 125, x["WiFi"]["EAP"]["UseProactiveKeyCaching"] ? true : false);
            A.store.Bool(ptr + 484 + 8 + 129, "UseSystemCAs" in x["WiFi"]["EAP"] ? true : false);
            A.store.Bool(ptr + 484 + 8 + 126, x["WiFi"]["EAP"]["UseSystemCAs"] ? true : false);
          }
          A.store.Bool(ptr + 484 + 178, "Frequency" in x["WiFi"] ? true : false);
          A.store.Int32(ptr + 484 + 140, x["WiFi"]["Frequency"] === undefined ? 0 : (x["WiFi"]["Frequency"] as number));
          A.store.Ref(ptr + 484 + 144, x["WiFi"]["FrequencyList"]);
          A.store.Ref(ptr + 484 + 148, x["WiFi"]["HexSSID"]);
          A.store.Bool(ptr + 484 + 179, "HiddenSSID" in x["WiFi"] ? true : false);
          A.store.Bool(ptr + 484 + 152, x["WiFi"]["HiddenSSID"] ? true : false);
          A.store.Ref(ptr + 484 + 156, x["WiFi"]["Passphrase"]);
          A.store.Bool(ptr + 484 + 180, "RoamThreshold" in x["WiFi"] ? true : false);
          A.store.Int32(
            ptr + 484 + 160,
            x["WiFi"]["RoamThreshold"] === undefined ? 0 : (x["WiFi"]["RoamThreshold"] as number)
          );
          A.store.Ref(ptr + 484 + 164, x["WiFi"]["SSID"]);
          A.store.Ref(ptr + 484 + 168, x["WiFi"]["Security"]);
          A.store.Bool(ptr + 484 + 181, "SignalStrength" in x["WiFi"] ? true : false);
          A.store.Int32(
            ptr + 484 + 172,
            x["WiFi"]["SignalStrength"] === undefined ? 0 : (x["WiFi"]["SignalStrength"] as number)
          );
        }
      }
    },
    "load_NetworkProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 127)) {
        x["Cellular"] = {};
        if (A.load.Bool(ptr + 0 + 121)) {
          x["Cellular"]["AutoConnect"] = A.load.Bool(ptr + 0 + 0);
        } else {
          delete x["Cellular"]["AutoConnect"];
        }
        x["Cellular"]["ActivationType"] = A.load.Ref(ptr + 0 + 4, undefined);
        x["Cellular"]["ActivationState"] = A.load.Enum(ptr + 0 + 8, [
          "Activated",
          "Activating",
          "NotActivated",
          "PartiallyActivated",
        ]);
        if (A.load.Bool(ptr + 0 + 122)) {
          x["Cellular"]["AllowRoaming"] = A.load.Bool(ptr + 0 + 12);
        } else {
          delete x["Cellular"]["AllowRoaming"];
        }
        x["Cellular"]["Family"] = A.load.Ref(ptr + 0 + 16, undefined);
        x["Cellular"]["FirmwareRevision"] = A.load.Ref(ptr + 0 + 20, undefined);
        x["Cellular"]["FoundNetworks"] = A.load.Ref(ptr + 0 + 24, undefined);
        x["Cellular"]["HardwareRevision"] = A.load.Ref(ptr + 0 + 28, undefined);
        if (A.load.Bool(ptr + 0 + 32 + 12)) {
          x["Cellular"]["HomeProvider"] = {};
          x["Cellular"]["HomeProvider"]["Name"] = A.load.Ref(ptr + 0 + 32 + 0, undefined);
          x["Cellular"]["HomeProvider"]["Code"] = A.load.Ref(ptr + 0 + 32 + 4, undefined);
          x["Cellular"]["HomeProvider"]["Country"] = A.load.Ref(ptr + 0 + 32 + 8, undefined);
        } else {
          delete x["Cellular"]["HomeProvider"];
        }
        x["Cellular"]["Manufacturer"] = A.load.Ref(ptr + 0 + 48, undefined);
        x["Cellular"]["ModelID"] = A.load.Ref(ptr + 0 + 52, undefined);
        x["Cellular"]["NetworkTechnology"] = A.load.Ref(ptr + 0 + 56, undefined);
        if (A.load.Bool(ptr + 0 + 60 + 12)) {
          x["Cellular"]["PaymentPortal"] = {};
          x["Cellular"]["PaymentPortal"]["Method"] = A.load.Ref(ptr + 0 + 60 + 0, undefined);
          x["Cellular"]["PaymentPortal"]["PostData"] = A.load.Ref(ptr + 0 + 60 + 4, undefined);
          x["Cellular"]["PaymentPortal"]["Url"] = A.load.Ref(ptr + 0 + 60 + 8, undefined);
        } else {
          delete x["Cellular"]["PaymentPortal"];
        }
        x["Cellular"]["RoamingState"] = A.load.Ref(ptr + 0 + 76, undefined);
        if (A.load.Bool(ptr + 0 + 123)) {
          x["Cellular"]["Scanning"] = A.load.Bool(ptr + 0 + 80);
        } else {
          delete x["Cellular"]["Scanning"];
        }
        if (A.load.Bool(ptr + 0 + 84 + 12)) {
          x["Cellular"]["ServingOperator"] = {};
          x["Cellular"]["ServingOperator"]["Name"] = A.load.Ref(ptr + 0 + 84 + 0, undefined);
          x["Cellular"]["ServingOperator"]["Code"] = A.load.Ref(ptr + 0 + 84 + 4, undefined);
          x["Cellular"]["ServingOperator"]["Country"] = A.load.Ref(ptr + 0 + 84 + 8, undefined);
        } else {
          delete x["Cellular"]["ServingOperator"];
        }
        if (A.load.Bool(ptr + 0 + 100 + 14)) {
          x["Cellular"]["SIMLockStatus"] = {};
          x["Cellular"]["SIMLockStatus"]["LockType"] = A.load.Ref(ptr + 0 + 100 + 0, undefined);
          if (A.load.Bool(ptr + 0 + 100 + 12)) {
            x["Cellular"]["SIMLockStatus"]["LockEnabled"] = A.load.Bool(ptr + 0 + 100 + 4);
          } else {
            delete x["Cellular"]["SIMLockStatus"]["LockEnabled"];
          }
          if (A.load.Bool(ptr + 0 + 100 + 13)) {
            x["Cellular"]["SIMLockStatus"]["RetriesLeft"] = A.load.Int32(ptr + 0 + 100 + 8);
          } else {
            delete x["Cellular"]["SIMLockStatus"]["RetriesLeft"];
          }
        } else {
          delete x["Cellular"]["SIMLockStatus"];
        }
        if (A.load.Bool(ptr + 0 + 124)) {
          x["Cellular"]["SIMPresent"] = A.load.Bool(ptr + 0 + 115);
        } else {
          delete x["Cellular"]["SIMPresent"];
        }
        if (A.load.Bool(ptr + 0 + 125)) {
          x["Cellular"]["SignalStrength"] = A.load.Int32(ptr + 0 + 116);
        } else {
          delete x["Cellular"]["SignalStrength"];
        }
        if (A.load.Bool(ptr + 0 + 126)) {
          x["Cellular"]["SupportNetworkScan"] = A.load.Bool(ptr + 0 + 120);
        } else {
          delete x["Cellular"]["SupportNetworkScan"];
        }
      } else {
        delete x["Cellular"];
      }
      if (A.load.Bool(ptr + 667)) {
        x["Connectable"] = A.load.Bool(ptr + 128);
      } else {
        delete x["Connectable"];
      }
      x["ConnectionState"] = A.load.Enum(ptr + 132, ["Connected", "Connecting", "NotConnected"]);
      x["ErrorState"] = A.load.Ref(ptr + 136, undefined);
      if (A.load.Bool(ptr + 140 + 140)) {
        x["Ethernet"] = {};
        if (A.load.Bool(ptr + 140 + 139)) {
          x["Ethernet"]["AutoConnect"] = A.load.Bool(ptr + 140 + 0);
        } else {
          delete x["Ethernet"]["AutoConnect"];
        }
        x["Ethernet"]["Authentication"] = A.load.Ref(ptr + 140 + 4, undefined);
        if (A.load.Bool(ptr + 140 + 8 + 130)) {
          x["Ethernet"]["EAP"] = {};
          x["Ethernet"]["EAP"]["AnonymousIdentity"] = A.load.Ref(ptr + 140 + 8 + 0, undefined);
          if (A.load.Bool(ptr + 140 + 8 + 4 + 45)) {
            x["Ethernet"]["EAP"]["ClientCertPattern"] = {};
            x["Ethernet"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(ptr + 140 + 8 + 4 + 0, undefined);
            if (A.load.Bool(ptr + 140 + 8 + 4 + 4 + 16)) {
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                ptr + 140 + 8 + 4 + 4 + 0,
                undefined
              );
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                ptr + 140 + 8 + 4 + 4 + 4,
                undefined
              );
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                ptr + 140 + 8 + 4 + 4 + 8,
                undefined
              );
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                ptr + 140 + 8 + 4 + 4 + 12,
                undefined
              );
            } else {
              delete x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"];
            }
            x["Ethernet"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(ptr + 140 + 8 + 4 + 24, undefined);
            if (A.load.Bool(ptr + 140 + 8 + 4 + 28 + 16)) {
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"] = {};
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                ptr + 140 + 8 + 4 + 28 + 0,
                undefined
              );
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                ptr + 140 + 8 + 4 + 28 + 4,
                undefined
              );
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                ptr + 140 + 8 + 4 + 28 + 8,
                undefined
              );
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                ptr + 140 + 8 + 4 + 28 + 12,
                undefined
              );
            } else {
              delete x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"];
            }
          } else {
            delete x["Ethernet"]["EAP"]["ClientCertPattern"];
          }
          x["Ethernet"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(ptr + 140 + 8 + 52, undefined);
          x["Ethernet"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(ptr + 140 + 8 + 56, undefined);
          x["Ethernet"]["EAP"]["ClientCertRef"] = A.load.Ref(ptr + 140 + 8 + 60, undefined);
          x["Ethernet"]["EAP"]["ClientCertType"] = A.load.Enum(ptr + 140 + 8 + 64, ["Ref", "Pattern"]);
          x["Ethernet"]["EAP"]["Identity"] = A.load.Ref(ptr + 140 + 8 + 68, undefined);
          x["Ethernet"]["EAP"]["Inner"] = A.load.Ref(ptr + 140 + 8 + 72, undefined);
          x["Ethernet"]["EAP"]["Outer"] = A.load.Ref(ptr + 140 + 8 + 76, undefined);
          x["Ethernet"]["EAP"]["Password"] = A.load.Ref(ptr + 140 + 8 + 80, undefined);
          if (A.load.Bool(ptr + 140 + 8 + 127)) {
            x["Ethernet"]["EAP"]["SaveCredentials"] = A.load.Bool(ptr + 140 + 8 + 84);
          } else {
            delete x["Ethernet"]["EAP"]["SaveCredentials"];
          }
          x["Ethernet"]["EAP"]["ServerCAPEMs"] = A.load.Ref(ptr + 140 + 8 + 88, undefined);
          x["Ethernet"]["EAP"]["ServerCARefs"] = A.load.Ref(ptr + 140 + 8 + 92, undefined);
          if (A.load.Bool(ptr + 140 + 8 + 96 + 28)) {
            x["Ethernet"]["EAP"]["SubjectMatch"] = {};
            x["Ethernet"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(ptr + 140 + 8 + 96 + 0, undefined);
            x["Ethernet"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(ptr + 140 + 8 + 96 + 4, undefined);
            x["Ethernet"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(ptr + 140 + 8 + 96 + 8, undefined);
            x["Ethernet"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(ptr + 140 + 8 + 96 + 12, undefined);
            x["Ethernet"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(ptr + 140 + 8 + 96 + 16, undefined);
            x["Ethernet"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(ptr + 140 + 8 + 96 + 20, undefined);
            if (A.load.Bool(ptr + 140 + 8 + 96 + 26)) {
              x["Ethernet"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(ptr + 140 + 8 + 96 + 24);
            } else {
              delete x["Ethernet"]["EAP"]["SubjectMatch"]["UserEditable"];
            }
            if (A.load.Bool(ptr + 140 + 8 + 96 + 27)) {
              x["Ethernet"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(ptr + 140 + 8 + 96 + 25);
            } else {
              delete x["Ethernet"]["EAP"]["SubjectMatch"]["DeviceEditable"];
            }
          } else {
            delete x["Ethernet"]["EAP"]["SubjectMatch"];
          }
          if (A.load.Bool(ptr + 140 + 8 + 128)) {
            x["Ethernet"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(ptr + 140 + 8 + 125);
          } else {
            delete x["Ethernet"]["EAP"]["UseProactiveKeyCaching"];
          }
          if (A.load.Bool(ptr + 140 + 8 + 129)) {
            x["Ethernet"]["EAP"]["UseSystemCAs"] = A.load.Bool(ptr + 140 + 8 + 126);
          } else {
            delete x["Ethernet"]["EAP"]["UseSystemCAs"];
          }
        } else {
          delete x["Ethernet"]["EAP"];
        }
      } else {
        delete x["Ethernet"];
      }
      x["GUID"] = A.load.Ref(ptr + 284, undefined);
      x["IPAddressConfigType"] = A.load.Enum(ptr + 288, ["DHCP", "Static"]);
      x["IPConfigs"] = A.load.Ref(ptr + 292, undefined);
      x["MacAddress"] = A.load.Ref(ptr + 296, undefined);
      if (A.load.Bool(ptr + 668)) {
        x["Metered"] = A.load.Bool(ptr + 300);
      } else {
        delete x["Metered"];
      }
      x["Name"] = A.load.Ref(ptr + 304, undefined);
      x["NameServersConfigType"] = A.load.Enum(ptr + 308, ["DHCP", "Static"]);
      if (A.load.Bool(ptr + 669)) {
        x["Priority"] = A.load.Int32(ptr + 312);
      } else {
        delete x["Priority"];
      }
      if (A.load.Bool(ptr + 316 + 60)) {
        x["ProxySettings"] = {};
        x["ProxySettings"]["Type"] = A.load.Enum(ptr + 316 + 0, ["Direct", "Manual", "PAC", "WPAD"]);
        if (A.load.Bool(ptr + 316 + 4 + 46)) {
          x["ProxySettings"]["Manual"] = {};
          if (A.load.Bool(ptr + 316 + 4 + 0 + 9)) {
            x["ProxySettings"]["Manual"]["HTTPProxy"] = {};
            x["ProxySettings"]["Manual"]["HTTPProxy"]["Host"] = A.load.Ref(ptr + 316 + 4 + 0 + 0, undefined);
            if (A.load.Bool(ptr + 316 + 4 + 0 + 8)) {
              x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"] = A.load.Int32(ptr + 316 + 4 + 0 + 4);
            } else {
              delete x["ProxySettings"]["Manual"]["HTTPProxy"]["Port"];
            }
          } else {
            delete x["ProxySettings"]["Manual"]["HTTPProxy"];
          }
          if (A.load.Bool(ptr + 316 + 4 + 12 + 9)) {
            x["ProxySettings"]["Manual"]["SecureHTTPProxy"] = {};
            x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Host"] = A.load.Ref(ptr + 316 + 4 + 12 + 0, undefined);
            if (A.load.Bool(ptr + 316 + 4 + 12 + 8)) {
              x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"] = A.load.Int32(ptr + 316 + 4 + 12 + 4);
            } else {
              delete x["ProxySettings"]["Manual"]["SecureHTTPProxy"]["Port"];
            }
          } else {
            delete x["ProxySettings"]["Manual"]["SecureHTTPProxy"];
          }
          if (A.load.Bool(ptr + 316 + 4 + 24 + 9)) {
            x["ProxySettings"]["Manual"]["FTPProxy"] = {};
            x["ProxySettings"]["Manual"]["FTPProxy"]["Host"] = A.load.Ref(ptr + 316 + 4 + 24 + 0, undefined);
            if (A.load.Bool(ptr + 316 + 4 + 24 + 8)) {
              x["ProxySettings"]["Manual"]["FTPProxy"]["Port"] = A.load.Int32(ptr + 316 + 4 + 24 + 4);
            } else {
              delete x["ProxySettings"]["Manual"]["FTPProxy"]["Port"];
            }
          } else {
            delete x["ProxySettings"]["Manual"]["FTPProxy"];
          }
          if (A.load.Bool(ptr + 316 + 4 + 36 + 9)) {
            x["ProxySettings"]["Manual"]["SOCKS"] = {};
            x["ProxySettings"]["Manual"]["SOCKS"]["Host"] = A.load.Ref(ptr + 316 + 4 + 36 + 0, undefined);
            if (A.load.Bool(ptr + 316 + 4 + 36 + 8)) {
              x["ProxySettings"]["Manual"]["SOCKS"]["Port"] = A.load.Int32(ptr + 316 + 4 + 36 + 4);
            } else {
              delete x["ProxySettings"]["Manual"]["SOCKS"]["Port"];
            }
          } else {
            delete x["ProxySettings"]["Manual"]["SOCKS"];
          }
        } else {
          delete x["ProxySettings"]["Manual"];
        }
        x["ProxySettings"]["ExcludeDomains"] = A.load.Ref(ptr + 316 + 52, undefined);
        x["ProxySettings"]["PAC"] = A.load.Ref(ptr + 316 + 56, undefined);
      } else {
        delete x["ProxySettings"];
      }
      if (A.load.Bool(ptr + 670)) {
        x["RestrictedConnectivity"] = A.load.Bool(ptr + 377);
      } else {
        delete x["RestrictedConnectivity"];
      }
      if (A.load.Bool(ptr + 380 + 37)) {
        x["StaticIPConfig"] = {};
        x["StaticIPConfig"]["Gateway"] = A.load.Ref(ptr + 380 + 0, undefined);
        x["StaticIPConfig"]["IPAddress"] = A.load.Ref(ptr + 380 + 4, undefined);
        x["StaticIPConfig"]["ExcludedRoutes"] = A.load.Ref(ptr + 380 + 8, undefined);
        x["StaticIPConfig"]["IncludedRoutes"] = A.load.Ref(ptr + 380 + 12, undefined);
        x["StaticIPConfig"]["NameServers"] = A.load.Ref(ptr + 380 + 16, undefined);
        x["StaticIPConfig"]["SearchDomains"] = A.load.Ref(ptr + 380 + 20, undefined);
        if (A.load.Bool(ptr + 380 + 36)) {
          x["StaticIPConfig"]["RoutingPrefix"] = A.load.Int32(ptr + 380 + 24);
        } else {
          delete x["StaticIPConfig"]["RoutingPrefix"];
        }
        x["StaticIPConfig"]["Type"] = A.load.Ref(ptr + 380 + 28, undefined);
        x["StaticIPConfig"]["WebProxyAutoDiscoveryUrl"] = A.load.Ref(ptr + 380 + 32, undefined);
      } else {
        delete x["StaticIPConfig"];
      }
      if (A.load.Bool(ptr + 420 + 37)) {
        x["SavedIPConfig"] = {};
        x["SavedIPConfig"]["Gateway"] = A.load.Ref(ptr + 420 + 0, undefined);
        x["SavedIPConfig"]["IPAddress"] = A.load.Ref(ptr + 420 + 4, undefined);
        x["SavedIPConfig"]["ExcludedRoutes"] = A.load.Ref(ptr + 420 + 8, undefined);
        x["SavedIPConfig"]["IncludedRoutes"] = A.load.Ref(ptr + 420 + 12, undefined);
        x["SavedIPConfig"]["NameServers"] = A.load.Ref(ptr + 420 + 16, undefined);
        x["SavedIPConfig"]["SearchDomains"] = A.load.Ref(ptr + 420 + 20, undefined);
        if (A.load.Bool(ptr + 420 + 36)) {
          x["SavedIPConfig"]["RoutingPrefix"] = A.load.Int32(ptr + 420 + 24);
        } else {
          delete x["SavedIPConfig"]["RoutingPrefix"];
        }
        x["SavedIPConfig"]["Type"] = A.load.Ref(ptr + 420 + 28, undefined);
        x["SavedIPConfig"]["WebProxyAutoDiscoveryUrl"] = A.load.Ref(ptr + 420 + 32, undefined);
      } else {
        delete x["SavedIPConfig"];
      }
      x["Source"] = A.load.Ref(ptr + 460, undefined);
      x["Type"] = A.load.Enum(ptr + 464, ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"]);
      if (A.load.Bool(ptr + 468 + 13)) {
        x["VPN"] = {};
        if (A.load.Bool(ptr + 468 + 12)) {
          x["VPN"]["AutoConnect"] = A.load.Bool(ptr + 468 + 0);
        } else {
          delete x["VPN"]["AutoConnect"];
        }
        x["VPN"]["Host"] = A.load.Ref(ptr + 468 + 4, undefined);
        x["VPN"]["Type"] = A.load.Ref(ptr + 468 + 8, undefined);
      } else {
        delete x["VPN"];
      }
      if (A.load.Bool(ptr + 484 + 182)) {
        x["WiFi"] = {};
        if (A.load.Bool(ptr + 484 + 176)) {
          x["WiFi"]["AllowGatewayARPPolling"] = A.load.Bool(ptr + 484 + 0);
        } else {
          delete x["WiFi"]["AllowGatewayARPPolling"];
        }
        if (A.load.Bool(ptr + 484 + 177)) {
          x["WiFi"]["AutoConnect"] = A.load.Bool(ptr + 484 + 1);
        } else {
          delete x["WiFi"]["AutoConnect"];
        }
        x["WiFi"]["BSSID"] = A.load.Ref(ptr + 484 + 4, undefined);
        if (A.load.Bool(ptr + 484 + 8 + 130)) {
          x["WiFi"]["EAP"] = {};
          x["WiFi"]["EAP"]["AnonymousIdentity"] = A.load.Ref(ptr + 484 + 8 + 0, undefined);
          if (A.load.Bool(ptr + 484 + 8 + 4 + 45)) {
            x["WiFi"]["EAP"]["ClientCertPattern"] = {};
            x["WiFi"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(ptr + 484 + 8 + 4 + 0, undefined);
            if (A.load.Bool(ptr + 484 + 8 + 4 + 4 + 16)) {
              x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
              x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                ptr + 484 + 8 + 4 + 4 + 0,
                undefined
              );
              x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                ptr + 484 + 8 + 4 + 4 + 4,
                undefined
              );
              x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                ptr + 484 + 8 + 4 + 4 + 8,
                undefined
              );
              x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                ptr + 484 + 8 + 4 + 4 + 12,
                undefined
              );
            } else {
              delete x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"];
            }
            x["WiFi"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(ptr + 484 + 8 + 4 + 24, undefined);
            if (A.load.Bool(ptr + 484 + 8 + 4 + 28 + 16)) {
              x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"] = {};
              x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                ptr + 484 + 8 + 4 + 28 + 0,
                undefined
              );
              x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                ptr + 484 + 8 + 4 + 28 + 4,
                undefined
              );
              x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                ptr + 484 + 8 + 4 + 28 + 8,
                undefined
              );
              x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                ptr + 484 + 8 + 4 + 28 + 12,
                undefined
              );
            } else {
              delete x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"];
            }
          } else {
            delete x["WiFi"]["EAP"]["ClientCertPattern"];
          }
          x["WiFi"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(ptr + 484 + 8 + 52, undefined);
          x["WiFi"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(ptr + 484 + 8 + 56, undefined);
          x["WiFi"]["EAP"]["ClientCertRef"] = A.load.Ref(ptr + 484 + 8 + 60, undefined);
          x["WiFi"]["EAP"]["ClientCertType"] = A.load.Enum(ptr + 484 + 8 + 64, ["Ref", "Pattern"]);
          x["WiFi"]["EAP"]["Identity"] = A.load.Ref(ptr + 484 + 8 + 68, undefined);
          x["WiFi"]["EAP"]["Inner"] = A.load.Ref(ptr + 484 + 8 + 72, undefined);
          x["WiFi"]["EAP"]["Outer"] = A.load.Ref(ptr + 484 + 8 + 76, undefined);
          x["WiFi"]["EAP"]["Password"] = A.load.Ref(ptr + 484 + 8 + 80, undefined);
          if (A.load.Bool(ptr + 484 + 8 + 127)) {
            x["WiFi"]["EAP"]["SaveCredentials"] = A.load.Bool(ptr + 484 + 8 + 84);
          } else {
            delete x["WiFi"]["EAP"]["SaveCredentials"];
          }
          x["WiFi"]["EAP"]["ServerCAPEMs"] = A.load.Ref(ptr + 484 + 8 + 88, undefined);
          x["WiFi"]["EAP"]["ServerCARefs"] = A.load.Ref(ptr + 484 + 8 + 92, undefined);
          if (A.load.Bool(ptr + 484 + 8 + 96 + 28)) {
            x["WiFi"]["EAP"]["SubjectMatch"] = {};
            x["WiFi"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(ptr + 484 + 8 + 96 + 0, undefined);
            x["WiFi"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(ptr + 484 + 8 + 96 + 4, undefined);
            x["WiFi"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(ptr + 484 + 8 + 96 + 8, undefined);
            x["WiFi"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(ptr + 484 + 8 + 96 + 12, undefined);
            x["WiFi"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(ptr + 484 + 8 + 96 + 16, undefined);
            x["WiFi"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(ptr + 484 + 8 + 96 + 20, undefined);
            if (A.load.Bool(ptr + 484 + 8 + 96 + 26)) {
              x["WiFi"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(ptr + 484 + 8 + 96 + 24);
            } else {
              delete x["WiFi"]["EAP"]["SubjectMatch"]["UserEditable"];
            }
            if (A.load.Bool(ptr + 484 + 8 + 96 + 27)) {
              x["WiFi"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(ptr + 484 + 8 + 96 + 25);
            } else {
              delete x["WiFi"]["EAP"]["SubjectMatch"]["DeviceEditable"];
            }
          } else {
            delete x["WiFi"]["EAP"]["SubjectMatch"];
          }
          if (A.load.Bool(ptr + 484 + 8 + 128)) {
            x["WiFi"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(ptr + 484 + 8 + 125);
          } else {
            delete x["WiFi"]["EAP"]["UseProactiveKeyCaching"];
          }
          if (A.load.Bool(ptr + 484 + 8 + 129)) {
            x["WiFi"]["EAP"]["UseSystemCAs"] = A.load.Bool(ptr + 484 + 8 + 126);
          } else {
            delete x["WiFi"]["EAP"]["UseSystemCAs"];
          }
        } else {
          delete x["WiFi"]["EAP"];
        }
        if (A.load.Bool(ptr + 484 + 178)) {
          x["WiFi"]["Frequency"] = A.load.Int32(ptr + 484 + 140);
        } else {
          delete x["WiFi"]["Frequency"];
        }
        x["WiFi"]["FrequencyList"] = A.load.Ref(ptr + 484 + 144, undefined);
        x["WiFi"]["HexSSID"] = A.load.Ref(ptr + 484 + 148, undefined);
        if (A.load.Bool(ptr + 484 + 179)) {
          x["WiFi"]["HiddenSSID"] = A.load.Bool(ptr + 484 + 152);
        } else {
          delete x["WiFi"]["HiddenSSID"];
        }
        x["WiFi"]["Passphrase"] = A.load.Ref(ptr + 484 + 156, undefined);
        if (A.load.Bool(ptr + 484 + 180)) {
          x["WiFi"]["RoamThreshold"] = A.load.Int32(ptr + 484 + 160);
        } else {
          delete x["WiFi"]["RoamThreshold"];
        }
        x["WiFi"]["SSID"] = A.load.Ref(ptr + 484 + 164, undefined);
        x["WiFi"]["Security"] = A.load.Ref(ptr + 484 + 168, undefined);
        if (A.load.Bool(ptr + 484 + 181)) {
          x["WiFi"]["SignalStrength"] = A.load.Int32(ptr + 484 + 172);
        } else {
          delete x["WiFi"]["SignalStrength"];
        }
      } else {
        delete x["WiFi"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ManagedThirdPartyVPNProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 36, false);

        A.store.Bool(ptr + 0 + 28, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Ref(ptr + 0 + 8, undefined);
        A.store.Ref(ptr + 0 + 12, undefined);
        A.store.Ref(ptr + 0 + 16, undefined);
        A.store.Ref(ptr + 0 + 20, undefined);
        A.store.Bool(ptr + 0 + 26, false);
        A.store.Bool(ptr + 0 + 24, false);
        A.store.Bool(ptr + 0 + 27, false);
        A.store.Bool(ptr + 0 + 25, false);
        A.store.Ref(ptr + 32, undefined);
      } else {
        A.store.Bool(ptr + 36, true);

        if (typeof x["ExtensionID"] === "undefined") {
          A.store.Bool(ptr + 0 + 28, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Ref(ptr + 0 + 8, undefined);
          A.store.Ref(ptr + 0 + 12, undefined);
          A.store.Ref(ptr + 0 + 16, undefined);
          A.store.Ref(ptr + 0 + 20, undefined);
          A.store.Bool(ptr + 0 + 26, false);
          A.store.Bool(ptr + 0 + 24, false);
          A.store.Bool(ptr + 0 + 27, false);
          A.store.Bool(ptr + 0 + 25, false);
        } else {
          A.store.Bool(ptr + 0 + 28, true);
          A.store.Ref(ptr + 0 + 0, x["ExtensionID"]["Active"]);
          A.store.Ref(ptr + 0 + 4, x["ExtensionID"]["Effective"]);
          A.store.Ref(ptr + 0 + 8, x["ExtensionID"]["UserPolicy"]);
          A.store.Ref(ptr + 0 + 12, x["ExtensionID"]["DevicePolicy"]);
          A.store.Ref(ptr + 0 + 16, x["ExtensionID"]["UserSetting"]);
          A.store.Ref(ptr + 0 + 20, x["ExtensionID"]["SharedSetting"]);
          A.store.Bool(ptr + 0 + 26, "UserEditable" in x["ExtensionID"] ? true : false);
          A.store.Bool(ptr + 0 + 24, x["ExtensionID"]["UserEditable"] ? true : false);
          A.store.Bool(ptr + 0 + 27, "DeviceEditable" in x["ExtensionID"] ? true : false);
          A.store.Bool(ptr + 0 + 25, x["ExtensionID"]["DeviceEditable"] ? true : false);
        }
        A.store.Ref(ptr + 32, x["ProviderName"]);
      }
    },
    "load_ManagedThirdPartyVPNProperties": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 28)) {
        x["ExtensionID"] = {};
        x["ExtensionID"]["Active"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["ExtensionID"]["Effective"] = A.load.Ref(ptr + 0 + 4, undefined);
        x["ExtensionID"]["UserPolicy"] = A.load.Ref(ptr + 0 + 8, undefined);
        x["ExtensionID"]["DevicePolicy"] = A.load.Ref(ptr + 0 + 12, undefined);
        x["ExtensionID"]["UserSetting"] = A.load.Ref(ptr + 0 + 16, undefined);
        x["ExtensionID"]["SharedSetting"] = A.load.Ref(ptr + 0 + 20, undefined);
        if (A.load.Bool(ptr + 0 + 26)) {
          x["ExtensionID"]["UserEditable"] = A.load.Bool(ptr + 0 + 24);
        } else {
          delete x["ExtensionID"]["UserEditable"];
        }
        if (A.load.Bool(ptr + 0 + 27)) {
          x["ExtensionID"]["DeviceEditable"] = A.load.Bool(ptr + 0 + 25);
        } else {
          delete x["ExtensionID"]["DeviceEditable"];
        }
      } else {
        delete x["ExtensionID"];
      }
      x["ProviderName"] = A.load.Ref(ptr + 32, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_WiMAXProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 136, false);
        A.store.Bool(ptr + 135, false);
        A.store.Bool(ptr + 0, false);

        A.store.Bool(ptr + 4 + 130, false);
        A.store.Ref(ptr + 4 + 0, undefined);

        A.store.Bool(ptr + 4 + 4 + 45, false);
        A.store.Ref(ptr + 4 + 4 + 0, undefined);

        A.store.Bool(ptr + 4 + 4 + 4 + 16, false);
        A.store.Ref(ptr + 4 + 4 + 4 + 0, undefined);
        A.store.Ref(ptr + 4 + 4 + 4 + 4, undefined);
        A.store.Ref(ptr + 4 + 4 + 4 + 8, undefined);
        A.store.Ref(ptr + 4 + 4 + 4 + 12, undefined);
        A.store.Ref(ptr + 4 + 4 + 24, undefined);

        A.store.Bool(ptr + 4 + 4 + 28 + 16, false);
        A.store.Ref(ptr + 4 + 4 + 28 + 0, undefined);
        A.store.Ref(ptr + 4 + 4 + 28 + 4, undefined);
        A.store.Ref(ptr + 4 + 4 + 28 + 8, undefined);
        A.store.Ref(ptr + 4 + 4 + 28 + 12, undefined);
        A.store.Ref(ptr + 4 + 52, undefined);
        A.store.Ref(ptr + 4 + 56, undefined);
        A.store.Ref(ptr + 4 + 60, undefined);
        A.store.Enum(ptr + 4 + 64, -1);
        A.store.Ref(ptr + 4 + 68, undefined);
        A.store.Ref(ptr + 4 + 72, undefined);
        A.store.Ref(ptr + 4 + 76, undefined);
        A.store.Ref(ptr + 4 + 80, undefined);
        A.store.Bool(ptr + 4 + 127, false);
        A.store.Bool(ptr + 4 + 84, false);
        A.store.Ref(ptr + 4 + 88, undefined);
        A.store.Ref(ptr + 4 + 92, undefined);

        A.store.Bool(ptr + 4 + 96 + 28, false);
        A.store.Ref(ptr + 4 + 96 + 0, undefined);
        A.store.Ref(ptr + 4 + 96 + 4, undefined);
        A.store.Ref(ptr + 4 + 96 + 8, undefined);
        A.store.Ref(ptr + 4 + 96 + 12, undefined);
        A.store.Ref(ptr + 4 + 96 + 16, undefined);
        A.store.Ref(ptr + 4 + 96 + 20, undefined);
        A.store.Bool(ptr + 4 + 96 + 26, false);
        A.store.Bool(ptr + 4 + 96 + 24, false);
        A.store.Bool(ptr + 4 + 96 + 27, false);
        A.store.Bool(ptr + 4 + 96 + 25, false);
        A.store.Bool(ptr + 4 + 128, false);
        A.store.Bool(ptr + 4 + 125, false);
        A.store.Bool(ptr + 4 + 129, false);
        A.store.Bool(ptr + 4 + 126, false);
      } else {
        A.store.Bool(ptr + 136, true);
        A.store.Bool(ptr + 135, "AutoConnect" in x ? true : false);
        A.store.Bool(ptr + 0, x["AutoConnect"] ? true : false);

        if (typeof x["EAP"] === "undefined") {
          A.store.Bool(ptr + 4 + 130, false);
          A.store.Ref(ptr + 4 + 0, undefined);

          A.store.Bool(ptr + 4 + 4 + 45, false);
          A.store.Ref(ptr + 4 + 4 + 0, undefined);

          A.store.Bool(ptr + 4 + 4 + 4 + 16, false);
          A.store.Ref(ptr + 4 + 4 + 4 + 0, undefined);
          A.store.Ref(ptr + 4 + 4 + 4 + 4, undefined);
          A.store.Ref(ptr + 4 + 4 + 4 + 8, undefined);
          A.store.Ref(ptr + 4 + 4 + 4 + 12, undefined);
          A.store.Ref(ptr + 4 + 4 + 24, undefined);

          A.store.Bool(ptr + 4 + 4 + 28 + 16, false);
          A.store.Ref(ptr + 4 + 4 + 28 + 0, undefined);
          A.store.Ref(ptr + 4 + 4 + 28 + 4, undefined);
          A.store.Ref(ptr + 4 + 4 + 28 + 8, undefined);
          A.store.Ref(ptr + 4 + 4 + 28 + 12, undefined);
          A.store.Ref(ptr + 4 + 52, undefined);
          A.store.Ref(ptr + 4 + 56, undefined);
          A.store.Ref(ptr + 4 + 60, undefined);
          A.store.Enum(ptr + 4 + 64, -1);
          A.store.Ref(ptr + 4 + 68, undefined);
          A.store.Ref(ptr + 4 + 72, undefined);
          A.store.Ref(ptr + 4 + 76, undefined);
          A.store.Ref(ptr + 4 + 80, undefined);
          A.store.Bool(ptr + 4 + 127, false);
          A.store.Bool(ptr + 4 + 84, false);
          A.store.Ref(ptr + 4 + 88, undefined);
          A.store.Ref(ptr + 4 + 92, undefined);

          A.store.Bool(ptr + 4 + 96 + 28, false);
          A.store.Ref(ptr + 4 + 96 + 0, undefined);
          A.store.Ref(ptr + 4 + 96 + 4, undefined);
          A.store.Ref(ptr + 4 + 96 + 8, undefined);
          A.store.Ref(ptr + 4 + 96 + 12, undefined);
          A.store.Ref(ptr + 4 + 96 + 16, undefined);
          A.store.Ref(ptr + 4 + 96 + 20, undefined);
          A.store.Bool(ptr + 4 + 96 + 26, false);
          A.store.Bool(ptr + 4 + 96 + 24, false);
          A.store.Bool(ptr + 4 + 96 + 27, false);
          A.store.Bool(ptr + 4 + 96 + 25, false);
          A.store.Bool(ptr + 4 + 128, false);
          A.store.Bool(ptr + 4 + 125, false);
          A.store.Bool(ptr + 4 + 129, false);
          A.store.Bool(ptr + 4 + 126, false);
        } else {
          A.store.Bool(ptr + 4 + 130, true);
          A.store.Ref(ptr + 4 + 0, x["EAP"]["AnonymousIdentity"]);

          if (typeof x["EAP"]["ClientCertPattern"] === "undefined") {
            A.store.Bool(ptr + 4 + 4 + 45, false);
            A.store.Ref(ptr + 4 + 4 + 0, undefined);

            A.store.Bool(ptr + 4 + 4 + 4 + 16, false);
            A.store.Ref(ptr + 4 + 4 + 4 + 0, undefined);
            A.store.Ref(ptr + 4 + 4 + 4 + 4, undefined);
            A.store.Ref(ptr + 4 + 4 + 4 + 8, undefined);
            A.store.Ref(ptr + 4 + 4 + 4 + 12, undefined);
            A.store.Ref(ptr + 4 + 4 + 24, undefined);

            A.store.Bool(ptr + 4 + 4 + 28 + 16, false);
            A.store.Ref(ptr + 4 + 4 + 28 + 0, undefined);
            A.store.Ref(ptr + 4 + 4 + 28 + 4, undefined);
            A.store.Ref(ptr + 4 + 4 + 28 + 8, undefined);
            A.store.Ref(ptr + 4 + 4 + 28 + 12, undefined);
          } else {
            A.store.Bool(ptr + 4 + 4 + 45, true);
            A.store.Ref(ptr + 4 + 4 + 0, x["EAP"]["ClientCertPattern"]["EnrollmentURI"]);

            if (typeof x["EAP"]["ClientCertPattern"]["Issuer"] === "undefined") {
              A.store.Bool(ptr + 4 + 4 + 4 + 16, false);
              A.store.Ref(ptr + 4 + 4 + 4 + 0, undefined);
              A.store.Ref(ptr + 4 + 4 + 4 + 4, undefined);
              A.store.Ref(ptr + 4 + 4 + 4 + 8, undefined);
              A.store.Ref(ptr + 4 + 4 + 4 + 12, undefined);
            } else {
              A.store.Bool(ptr + 4 + 4 + 4 + 16, true);
              A.store.Ref(ptr + 4 + 4 + 4 + 0, x["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"]);
              A.store.Ref(ptr + 4 + 4 + 4 + 4, x["EAP"]["ClientCertPattern"]["Issuer"]["Locality"]);
              A.store.Ref(ptr + 4 + 4 + 4 + 8, x["EAP"]["ClientCertPattern"]["Issuer"]["Organization"]);
              A.store.Ref(ptr + 4 + 4 + 4 + 12, x["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"]);
            }
            A.store.Ref(ptr + 4 + 4 + 24, x["EAP"]["ClientCertPattern"]["IssuerCARef"]);

            if (typeof x["EAP"]["ClientCertPattern"]["Subject"] === "undefined") {
              A.store.Bool(ptr + 4 + 4 + 28 + 16, false);
              A.store.Ref(ptr + 4 + 4 + 28 + 0, undefined);
              A.store.Ref(ptr + 4 + 4 + 28 + 4, undefined);
              A.store.Ref(ptr + 4 + 4 + 28 + 8, undefined);
              A.store.Ref(ptr + 4 + 4 + 28 + 12, undefined);
            } else {
              A.store.Bool(ptr + 4 + 4 + 28 + 16, true);
              A.store.Ref(ptr + 4 + 4 + 28 + 0, x["EAP"]["ClientCertPattern"]["Subject"]["CommonName"]);
              A.store.Ref(ptr + 4 + 4 + 28 + 4, x["EAP"]["ClientCertPattern"]["Subject"]["Locality"]);
              A.store.Ref(ptr + 4 + 4 + 28 + 8, x["EAP"]["ClientCertPattern"]["Subject"]["Organization"]);
              A.store.Ref(ptr + 4 + 4 + 28 + 12, x["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"]);
            }
          }
          A.store.Ref(ptr + 4 + 52, x["EAP"]["ClientCertPKCS11Id"]);
          A.store.Ref(ptr + 4 + 56, x["EAP"]["ClientCertProvisioningProfileId"]);
          A.store.Ref(ptr + 4 + 60, x["EAP"]["ClientCertRef"]);
          A.store.Enum(ptr + 4 + 64, ["Ref", "Pattern"].indexOf(x["EAP"]["ClientCertType"] as string));
          A.store.Ref(ptr + 4 + 68, x["EAP"]["Identity"]);
          A.store.Ref(ptr + 4 + 72, x["EAP"]["Inner"]);
          A.store.Ref(ptr + 4 + 76, x["EAP"]["Outer"]);
          A.store.Ref(ptr + 4 + 80, x["EAP"]["Password"]);
          A.store.Bool(ptr + 4 + 127, "SaveCredentials" in x["EAP"] ? true : false);
          A.store.Bool(ptr + 4 + 84, x["EAP"]["SaveCredentials"] ? true : false);
          A.store.Ref(ptr + 4 + 88, x["EAP"]["ServerCAPEMs"]);
          A.store.Ref(ptr + 4 + 92, x["EAP"]["ServerCARefs"]);

          if (typeof x["EAP"]["SubjectMatch"] === "undefined") {
            A.store.Bool(ptr + 4 + 96 + 28, false);
            A.store.Ref(ptr + 4 + 96 + 0, undefined);
            A.store.Ref(ptr + 4 + 96 + 4, undefined);
            A.store.Ref(ptr + 4 + 96 + 8, undefined);
            A.store.Ref(ptr + 4 + 96 + 12, undefined);
            A.store.Ref(ptr + 4 + 96 + 16, undefined);
            A.store.Ref(ptr + 4 + 96 + 20, undefined);
            A.store.Bool(ptr + 4 + 96 + 26, false);
            A.store.Bool(ptr + 4 + 96 + 24, false);
            A.store.Bool(ptr + 4 + 96 + 27, false);
            A.store.Bool(ptr + 4 + 96 + 25, false);
          } else {
            A.store.Bool(ptr + 4 + 96 + 28, true);
            A.store.Ref(ptr + 4 + 96 + 0, x["EAP"]["SubjectMatch"]["Active"]);
            A.store.Ref(ptr + 4 + 96 + 4, x["EAP"]["SubjectMatch"]["Effective"]);
            A.store.Ref(ptr + 4 + 96 + 8, x["EAP"]["SubjectMatch"]["UserPolicy"]);
            A.store.Ref(ptr + 4 + 96 + 12, x["EAP"]["SubjectMatch"]["DevicePolicy"]);
            A.store.Ref(ptr + 4 + 96 + 16, x["EAP"]["SubjectMatch"]["UserSetting"]);
            A.store.Ref(ptr + 4 + 96 + 20, x["EAP"]["SubjectMatch"]["SharedSetting"]);
            A.store.Bool(ptr + 4 + 96 + 26, "UserEditable" in x["EAP"]["SubjectMatch"] ? true : false);
            A.store.Bool(ptr + 4 + 96 + 24, x["EAP"]["SubjectMatch"]["UserEditable"] ? true : false);
            A.store.Bool(ptr + 4 + 96 + 27, "DeviceEditable" in x["EAP"]["SubjectMatch"] ? true : false);
            A.store.Bool(ptr + 4 + 96 + 25, x["EAP"]["SubjectMatch"]["DeviceEditable"] ? true : false);
          }
          A.store.Bool(ptr + 4 + 128, "UseProactiveKeyCaching" in x["EAP"] ? true : false);
          A.store.Bool(ptr + 4 + 125, x["EAP"]["UseProactiveKeyCaching"] ? true : false);
          A.store.Bool(ptr + 4 + 129, "UseSystemCAs" in x["EAP"] ? true : false);
          A.store.Bool(ptr + 4 + 126, x["EAP"]["UseSystemCAs"] ? true : false);
        }
      }
    },
    "load_WiMAXProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 135)) {
        x["AutoConnect"] = A.load.Bool(ptr + 0);
      } else {
        delete x["AutoConnect"];
      }
      if (A.load.Bool(ptr + 4 + 130)) {
        x["EAP"] = {};
        x["EAP"]["AnonymousIdentity"] = A.load.Ref(ptr + 4 + 0, undefined);
        if (A.load.Bool(ptr + 4 + 4 + 45)) {
          x["EAP"]["ClientCertPattern"] = {};
          x["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(ptr + 4 + 4 + 0, undefined);
          if (A.load.Bool(ptr + 4 + 4 + 4 + 16)) {
            x["EAP"]["ClientCertPattern"]["Issuer"] = {};
            x["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(ptr + 4 + 4 + 4 + 0, undefined);
            x["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(ptr + 4 + 4 + 4 + 4, undefined);
            x["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(ptr + 4 + 4 + 4 + 8, undefined);
            x["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(ptr + 4 + 4 + 4 + 12, undefined);
          } else {
            delete x["EAP"]["ClientCertPattern"]["Issuer"];
          }
          x["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(ptr + 4 + 4 + 24, undefined);
          if (A.load.Bool(ptr + 4 + 4 + 28 + 16)) {
            x["EAP"]["ClientCertPattern"]["Subject"] = {};
            x["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(ptr + 4 + 4 + 28 + 0, undefined);
            x["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(ptr + 4 + 4 + 28 + 4, undefined);
            x["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(ptr + 4 + 4 + 28 + 8, undefined);
            x["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
              ptr + 4 + 4 + 28 + 12,
              undefined
            );
          } else {
            delete x["EAP"]["ClientCertPattern"]["Subject"];
          }
        } else {
          delete x["EAP"]["ClientCertPattern"];
        }
        x["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(ptr + 4 + 52, undefined);
        x["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(ptr + 4 + 56, undefined);
        x["EAP"]["ClientCertRef"] = A.load.Ref(ptr + 4 + 60, undefined);
        x["EAP"]["ClientCertType"] = A.load.Enum(ptr + 4 + 64, ["Ref", "Pattern"]);
        x["EAP"]["Identity"] = A.load.Ref(ptr + 4 + 68, undefined);
        x["EAP"]["Inner"] = A.load.Ref(ptr + 4 + 72, undefined);
        x["EAP"]["Outer"] = A.load.Ref(ptr + 4 + 76, undefined);
        x["EAP"]["Password"] = A.load.Ref(ptr + 4 + 80, undefined);
        if (A.load.Bool(ptr + 4 + 127)) {
          x["EAP"]["SaveCredentials"] = A.load.Bool(ptr + 4 + 84);
        } else {
          delete x["EAP"]["SaveCredentials"];
        }
        x["EAP"]["ServerCAPEMs"] = A.load.Ref(ptr + 4 + 88, undefined);
        x["EAP"]["ServerCARefs"] = A.load.Ref(ptr + 4 + 92, undefined);
        if (A.load.Bool(ptr + 4 + 96 + 28)) {
          x["EAP"]["SubjectMatch"] = {};
          x["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(ptr + 4 + 96 + 0, undefined);
          x["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(ptr + 4 + 96 + 4, undefined);
          x["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(ptr + 4 + 96 + 8, undefined);
          x["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(ptr + 4 + 96 + 12, undefined);
          x["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(ptr + 4 + 96 + 16, undefined);
          x["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(ptr + 4 + 96 + 20, undefined);
          if (A.load.Bool(ptr + 4 + 96 + 26)) {
            x["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(ptr + 4 + 96 + 24);
          } else {
            delete x["EAP"]["SubjectMatch"]["UserEditable"];
          }
          if (A.load.Bool(ptr + 4 + 96 + 27)) {
            x["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(ptr + 4 + 96 + 25);
          } else {
            delete x["EAP"]["SubjectMatch"]["DeviceEditable"];
          }
        } else {
          delete x["EAP"]["SubjectMatch"];
        }
        if (A.load.Bool(ptr + 4 + 128)) {
          x["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(ptr + 4 + 125);
        } else {
          delete x["EAP"]["UseProactiveKeyCaching"];
        }
        if (A.load.Bool(ptr + 4 + 129)) {
          x["EAP"]["UseSystemCAs"] = A.load.Bool(ptr + 4 + 126);
        } else {
          delete x["EAP"]["UseSystemCAs"];
        }
      } else {
        delete x["EAP"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_NetworkConfigProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 634, false);

        A.store.Bool(ptr + 0 + 127, false);
        A.store.Bool(ptr + 0 + 121, false);
        A.store.Bool(ptr + 0 + 0, false);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Enum(ptr + 0 + 8, -1);
        A.store.Bool(ptr + 0 + 122, false);
        A.store.Bool(ptr + 0 + 12, false);
        A.store.Ref(ptr + 0 + 16, undefined);
        A.store.Ref(ptr + 0 + 20, undefined);
        A.store.Ref(ptr + 0 + 24, undefined);
        A.store.Ref(ptr + 0 + 28, undefined);

        A.store.Bool(ptr + 0 + 32 + 12, false);
        A.store.Ref(ptr + 0 + 32 + 0, undefined);
        A.store.Ref(ptr + 0 + 32 + 4, undefined);
        A.store.Ref(ptr + 0 + 32 + 8, undefined);
        A.store.Ref(ptr + 0 + 48, undefined);
        A.store.Ref(ptr + 0 + 52, undefined);
        A.store.Ref(ptr + 0 + 56, undefined);

        A.store.Bool(ptr + 0 + 60 + 12, false);
        A.store.Ref(ptr + 0 + 60 + 0, undefined);
        A.store.Ref(ptr + 0 + 60 + 4, undefined);
        A.store.Ref(ptr + 0 + 60 + 8, undefined);
        A.store.Ref(ptr + 0 + 76, undefined);
        A.store.Bool(ptr + 0 + 123, false);
        A.store.Bool(ptr + 0 + 80, false);

        A.store.Bool(ptr + 0 + 84 + 12, false);
        A.store.Ref(ptr + 0 + 84 + 0, undefined);
        A.store.Ref(ptr + 0 + 84 + 4, undefined);
        A.store.Ref(ptr + 0 + 84 + 8, undefined);

        A.store.Bool(ptr + 0 + 100 + 14, false);
        A.store.Ref(ptr + 0 + 100 + 0, undefined);
        A.store.Bool(ptr + 0 + 100 + 12, false);
        A.store.Bool(ptr + 0 + 100 + 4, false);
        A.store.Bool(ptr + 0 + 100 + 13, false);
        A.store.Int32(ptr + 0 + 100 + 8, 0);
        A.store.Bool(ptr + 0 + 124, false);
        A.store.Bool(ptr + 0 + 115, false);
        A.store.Bool(ptr + 0 + 125, false);
        A.store.Int32(ptr + 0 + 116, 0);
        A.store.Bool(ptr + 0 + 126, false);
        A.store.Bool(ptr + 0 + 120, false);

        A.store.Bool(ptr + 128 + 140, false);
        A.store.Bool(ptr + 128 + 139, false);
        A.store.Bool(ptr + 128 + 0, false);
        A.store.Ref(ptr + 128 + 4, undefined);

        A.store.Bool(ptr + 128 + 8 + 130, false);
        A.store.Ref(ptr + 128 + 8 + 0, undefined);

        A.store.Bool(ptr + 128 + 8 + 4 + 45, false);
        A.store.Ref(ptr + 128 + 8 + 4 + 0, undefined);

        A.store.Bool(ptr + 128 + 8 + 4 + 4 + 16, false);
        A.store.Ref(ptr + 128 + 8 + 4 + 4 + 0, undefined);
        A.store.Ref(ptr + 128 + 8 + 4 + 4 + 4, undefined);
        A.store.Ref(ptr + 128 + 8 + 4 + 4 + 8, undefined);
        A.store.Ref(ptr + 128 + 8 + 4 + 4 + 12, undefined);
        A.store.Ref(ptr + 128 + 8 + 4 + 24, undefined);

        A.store.Bool(ptr + 128 + 8 + 4 + 28 + 16, false);
        A.store.Ref(ptr + 128 + 8 + 4 + 28 + 0, undefined);
        A.store.Ref(ptr + 128 + 8 + 4 + 28 + 4, undefined);
        A.store.Ref(ptr + 128 + 8 + 4 + 28 + 8, undefined);
        A.store.Ref(ptr + 128 + 8 + 4 + 28 + 12, undefined);
        A.store.Ref(ptr + 128 + 8 + 52, undefined);
        A.store.Ref(ptr + 128 + 8 + 56, undefined);
        A.store.Ref(ptr + 128 + 8 + 60, undefined);
        A.store.Enum(ptr + 128 + 8 + 64, -1);
        A.store.Ref(ptr + 128 + 8 + 68, undefined);
        A.store.Ref(ptr + 128 + 8 + 72, undefined);
        A.store.Ref(ptr + 128 + 8 + 76, undefined);
        A.store.Ref(ptr + 128 + 8 + 80, undefined);
        A.store.Bool(ptr + 128 + 8 + 127, false);
        A.store.Bool(ptr + 128 + 8 + 84, false);
        A.store.Ref(ptr + 128 + 8 + 88, undefined);
        A.store.Ref(ptr + 128 + 8 + 92, undefined);

        A.store.Bool(ptr + 128 + 8 + 96 + 28, false);
        A.store.Ref(ptr + 128 + 8 + 96 + 0, undefined);
        A.store.Ref(ptr + 128 + 8 + 96 + 4, undefined);
        A.store.Ref(ptr + 128 + 8 + 96 + 8, undefined);
        A.store.Ref(ptr + 128 + 8 + 96 + 12, undefined);
        A.store.Ref(ptr + 128 + 8 + 96 + 16, undefined);
        A.store.Ref(ptr + 128 + 8 + 96 + 20, undefined);
        A.store.Bool(ptr + 128 + 8 + 96 + 26, false);
        A.store.Bool(ptr + 128 + 8 + 96 + 24, false);
        A.store.Bool(ptr + 128 + 8 + 96 + 27, false);
        A.store.Bool(ptr + 128 + 8 + 96 + 25, false);
        A.store.Bool(ptr + 128 + 8 + 128, false);
        A.store.Bool(ptr + 128 + 8 + 125, false);
        A.store.Bool(ptr + 128 + 8 + 129, false);
        A.store.Bool(ptr + 128 + 8 + 126, false);
        A.store.Ref(ptr + 272, undefined);
        A.store.Enum(ptr + 276, -1);
        A.store.Ref(ptr + 280, undefined);
        A.store.Enum(ptr + 284, -1);
        A.store.Bool(ptr + 633, false);
        A.store.Int32(ptr + 288, 0);
        A.store.Enum(ptr + 292, -1);

        A.store.Bool(ptr + 296 + 13, false);
        A.store.Bool(ptr + 296 + 12, false);
        A.store.Bool(ptr + 296 + 0, false);
        A.store.Ref(ptr + 296 + 4, undefined);
        A.store.Ref(ptr + 296 + 8, undefined);

        A.store.Bool(ptr + 312 + 182, false);
        A.store.Bool(ptr + 312 + 176, false);
        A.store.Bool(ptr + 312 + 0, false);
        A.store.Bool(ptr + 312 + 177, false);
        A.store.Bool(ptr + 312 + 1, false);
        A.store.Ref(ptr + 312 + 4, undefined);

        A.store.Bool(ptr + 312 + 8 + 130, false);
        A.store.Ref(ptr + 312 + 8 + 0, undefined);

        A.store.Bool(ptr + 312 + 8 + 4 + 45, false);
        A.store.Ref(ptr + 312 + 8 + 4 + 0, undefined);

        A.store.Bool(ptr + 312 + 8 + 4 + 4 + 16, false);
        A.store.Ref(ptr + 312 + 8 + 4 + 4 + 0, undefined);
        A.store.Ref(ptr + 312 + 8 + 4 + 4 + 4, undefined);
        A.store.Ref(ptr + 312 + 8 + 4 + 4 + 8, undefined);
        A.store.Ref(ptr + 312 + 8 + 4 + 4 + 12, undefined);
        A.store.Ref(ptr + 312 + 8 + 4 + 24, undefined);

        A.store.Bool(ptr + 312 + 8 + 4 + 28 + 16, false);
        A.store.Ref(ptr + 312 + 8 + 4 + 28 + 0, undefined);
        A.store.Ref(ptr + 312 + 8 + 4 + 28 + 4, undefined);
        A.store.Ref(ptr + 312 + 8 + 4 + 28 + 8, undefined);
        A.store.Ref(ptr + 312 + 8 + 4 + 28 + 12, undefined);
        A.store.Ref(ptr + 312 + 8 + 52, undefined);
        A.store.Ref(ptr + 312 + 8 + 56, undefined);
        A.store.Ref(ptr + 312 + 8 + 60, undefined);
        A.store.Enum(ptr + 312 + 8 + 64, -1);
        A.store.Ref(ptr + 312 + 8 + 68, undefined);
        A.store.Ref(ptr + 312 + 8 + 72, undefined);
        A.store.Ref(ptr + 312 + 8 + 76, undefined);
        A.store.Ref(ptr + 312 + 8 + 80, undefined);
        A.store.Bool(ptr + 312 + 8 + 127, false);
        A.store.Bool(ptr + 312 + 8 + 84, false);
        A.store.Ref(ptr + 312 + 8 + 88, undefined);
        A.store.Ref(ptr + 312 + 8 + 92, undefined);

        A.store.Bool(ptr + 312 + 8 + 96 + 28, false);
        A.store.Ref(ptr + 312 + 8 + 96 + 0, undefined);
        A.store.Ref(ptr + 312 + 8 + 96 + 4, undefined);
        A.store.Ref(ptr + 312 + 8 + 96 + 8, undefined);
        A.store.Ref(ptr + 312 + 8 + 96 + 12, undefined);
        A.store.Ref(ptr + 312 + 8 + 96 + 16, undefined);
        A.store.Ref(ptr + 312 + 8 + 96 + 20, undefined);
        A.store.Bool(ptr + 312 + 8 + 96 + 26, false);
        A.store.Bool(ptr + 312 + 8 + 96 + 24, false);
        A.store.Bool(ptr + 312 + 8 + 96 + 27, false);
        A.store.Bool(ptr + 312 + 8 + 96 + 25, false);
        A.store.Bool(ptr + 312 + 8 + 128, false);
        A.store.Bool(ptr + 312 + 8 + 125, false);
        A.store.Bool(ptr + 312 + 8 + 129, false);
        A.store.Bool(ptr + 312 + 8 + 126, false);
        A.store.Bool(ptr + 312 + 178, false);
        A.store.Int32(ptr + 312 + 140, 0);
        A.store.Ref(ptr + 312 + 144, undefined);
        A.store.Ref(ptr + 312 + 148, undefined);
        A.store.Bool(ptr + 312 + 179, false);
        A.store.Bool(ptr + 312 + 152, false);
        A.store.Ref(ptr + 312 + 156, undefined);
        A.store.Bool(ptr + 312 + 180, false);
        A.store.Int32(ptr + 312 + 160, 0);
        A.store.Ref(ptr + 312 + 164, undefined);
        A.store.Ref(ptr + 312 + 168, undefined);
        A.store.Bool(ptr + 312 + 181, false);
        A.store.Int32(ptr + 312 + 172, 0);

        A.store.Bool(ptr + 496 + 136, false);
        A.store.Bool(ptr + 496 + 135, false);
        A.store.Bool(ptr + 496 + 0, false);

        A.store.Bool(ptr + 496 + 4 + 130, false);
        A.store.Ref(ptr + 496 + 4 + 0, undefined);

        A.store.Bool(ptr + 496 + 4 + 4 + 45, false);
        A.store.Ref(ptr + 496 + 4 + 4 + 0, undefined);

        A.store.Bool(ptr + 496 + 4 + 4 + 4 + 16, false);
        A.store.Ref(ptr + 496 + 4 + 4 + 4 + 0, undefined);
        A.store.Ref(ptr + 496 + 4 + 4 + 4 + 4, undefined);
        A.store.Ref(ptr + 496 + 4 + 4 + 4 + 8, undefined);
        A.store.Ref(ptr + 496 + 4 + 4 + 4 + 12, undefined);
        A.store.Ref(ptr + 496 + 4 + 4 + 24, undefined);

        A.store.Bool(ptr + 496 + 4 + 4 + 28 + 16, false);
        A.store.Ref(ptr + 496 + 4 + 4 + 28 + 0, undefined);
        A.store.Ref(ptr + 496 + 4 + 4 + 28 + 4, undefined);
        A.store.Ref(ptr + 496 + 4 + 4 + 28 + 8, undefined);
        A.store.Ref(ptr + 496 + 4 + 4 + 28 + 12, undefined);
        A.store.Ref(ptr + 496 + 4 + 52, undefined);
        A.store.Ref(ptr + 496 + 4 + 56, undefined);
        A.store.Ref(ptr + 496 + 4 + 60, undefined);
        A.store.Enum(ptr + 496 + 4 + 64, -1);
        A.store.Ref(ptr + 496 + 4 + 68, undefined);
        A.store.Ref(ptr + 496 + 4 + 72, undefined);
        A.store.Ref(ptr + 496 + 4 + 76, undefined);
        A.store.Ref(ptr + 496 + 4 + 80, undefined);
        A.store.Bool(ptr + 496 + 4 + 127, false);
        A.store.Bool(ptr + 496 + 4 + 84, false);
        A.store.Ref(ptr + 496 + 4 + 88, undefined);
        A.store.Ref(ptr + 496 + 4 + 92, undefined);

        A.store.Bool(ptr + 496 + 4 + 96 + 28, false);
        A.store.Ref(ptr + 496 + 4 + 96 + 0, undefined);
        A.store.Ref(ptr + 496 + 4 + 96 + 4, undefined);
        A.store.Ref(ptr + 496 + 4 + 96 + 8, undefined);
        A.store.Ref(ptr + 496 + 4 + 96 + 12, undefined);
        A.store.Ref(ptr + 496 + 4 + 96 + 16, undefined);
        A.store.Ref(ptr + 496 + 4 + 96 + 20, undefined);
        A.store.Bool(ptr + 496 + 4 + 96 + 26, false);
        A.store.Bool(ptr + 496 + 4 + 96 + 24, false);
        A.store.Bool(ptr + 496 + 4 + 96 + 27, false);
        A.store.Bool(ptr + 496 + 4 + 96 + 25, false);
        A.store.Bool(ptr + 496 + 4 + 128, false);
        A.store.Bool(ptr + 496 + 4 + 125, false);
        A.store.Bool(ptr + 496 + 4 + 129, false);
        A.store.Bool(ptr + 496 + 4 + 126, false);
      } else {
        A.store.Bool(ptr + 634, true);

        if (typeof x["Cellular"] === "undefined") {
          A.store.Bool(ptr + 0 + 127, false);
          A.store.Bool(ptr + 0 + 121, false);
          A.store.Bool(ptr + 0 + 0, false);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Enum(ptr + 0 + 8, -1);
          A.store.Bool(ptr + 0 + 122, false);
          A.store.Bool(ptr + 0 + 12, false);
          A.store.Ref(ptr + 0 + 16, undefined);
          A.store.Ref(ptr + 0 + 20, undefined);
          A.store.Ref(ptr + 0 + 24, undefined);
          A.store.Ref(ptr + 0 + 28, undefined);

          A.store.Bool(ptr + 0 + 32 + 12, false);
          A.store.Ref(ptr + 0 + 32 + 0, undefined);
          A.store.Ref(ptr + 0 + 32 + 4, undefined);
          A.store.Ref(ptr + 0 + 32 + 8, undefined);
          A.store.Ref(ptr + 0 + 48, undefined);
          A.store.Ref(ptr + 0 + 52, undefined);
          A.store.Ref(ptr + 0 + 56, undefined);

          A.store.Bool(ptr + 0 + 60 + 12, false);
          A.store.Ref(ptr + 0 + 60 + 0, undefined);
          A.store.Ref(ptr + 0 + 60 + 4, undefined);
          A.store.Ref(ptr + 0 + 60 + 8, undefined);
          A.store.Ref(ptr + 0 + 76, undefined);
          A.store.Bool(ptr + 0 + 123, false);
          A.store.Bool(ptr + 0 + 80, false);

          A.store.Bool(ptr + 0 + 84 + 12, false);
          A.store.Ref(ptr + 0 + 84 + 0, undefined);
          A.store.Ref(ptr + 0 + 84 + 4, undefined);
          A.store.Ref(ptr + 0 + 84 + 8, undefined);

          A.store.Bool(ptr + 0 + 100 + 14, false);
          A.store.Ref(ptr + 0 + 100 + 0, undefined);
          A.store.Bool(ptr + 0 + 100 + 12, false);
          A.store.Bool(ptr + 0 + 100 + 4, false);
          A.store.Bool(ptr + 0 + 100 + 13, false);
          A.store.Int32(ptr + 0 + 100 + 8, 0);
          A.store.Bool(ptr + 0 + 124, false);
          A.store.Bool(ptr + 0 + 115, false);
          A.store.Bool(ptr + 0 + 125, false);
          A.store.Int32(ptr + 0 + 116, 0);
          A.store.Bool(ptr + 0 + 126, false);
          A.store.Bool(ptr + 0 + 120, false);
        } else {
          A.store.Bool(ptr + 0 + 127, true);
          A.store.Bool(ptr + 0 + 121, "AutoConnect" in x["Cellular"] ? true : false);
          A.store.Bool(ptr + 0 + 0, x["Cellular"]["AutoConnect"] ? true : false);
          A.store.Ref(ptr + 0 + 4, x["Cellular"]["ActivationType"]);
          A.store.Enum(
            ptr + 0 + 8,
            ["Activated", "Activating", "NotActivated", "PartiallyActivated"].indexOf(
              x["Cellular"]["ActivationState"] as string
            )
          );
          A.store.Bool(ptr + 0 + 122, "AllowRoaming" in x["Cellular"] ? true : false);
          A.store.Bool(ptr + 0 + 12, x["Cellular"]["AllowRoaming"] ? true : false);
          A.store.Ref(ptr + 0 + 16, x["Cellular"]["Family"]);
          A.store.Ref(ptr + 0 + 20, x["Cellular"]["FirmwareRevision"]);
          A.store.Ref(ptr + 0 + 24, x["Cellular"]["FoundNetworks"]);
          A.store.Ref(ptr + 0 + 28, x["Cellular"]["HardwareRevision"]);

          if (typeof x["Cellular"]["HomeProvider"] === "undefined") {
            A.store.Bool(ptr + 0 + 32 + 12, false);
            A.store.Ref(ptr + 0 + 32 + 0, undefined);
            A.store.Ref(ptr + 0 + 32 + 4, undefined);
            A.store.Ref(ptr + 0 + 32 + 8, undefined);
          } else {
            A.store.Bool(ptr + 0 + 32 + 12, true);
            A.store.Ref(ptr + 0 + 32 + 0, x["Cellular"]["HomeProvider"]["Name"]);
            A.store.Ref(ptr + 0 + 32 + 4, x["Cellular"]["HomeProvider"]["Code"]);
            A.store.Ref(ptr + 0 + 32 + 8, x["Cellular"]["HomeProvider"]["Country"]);
          }
          A.store.Ref(ptr + 0 + 48, x["Cellular"]["Manufacturer"]);
          A.store.Ref(ptr + 0 + 52, x["Cellular"]["ModelID"]);
          A.store.Ref(ptr + 0 + 56, x["Cellular"]["NetworkTechnology"]);

          if (typeof x["Cellular"]["PaymentPortal"] === "undefined") {
            A.store.Bool(ptr + 0 + 60 + 12, false);
            A.store.Ref(ptr + 0 + 60 + 0, undefined);
            A.store.Ref(ptr + 0 + 60 + 4, undefined);
            A.store.Ref(ptr + 0 + 60 + 8, undefined);
          } else {
            A.store.Bool(ptr + 0 + 60 + 12, true);
            A.store.Ref(ptr + 0 + 60 + 0, x["Cellular"]["PaymentPortal"]["Method"]);
            A.store.Ref(ptr + 0 + 60 + 4, x["Cellular"]["PaymentPortal"]["PostData"]);
            A.store.Ref(ptr + 0 + 60 + 8, x["Cellular"]["PaymentPortal"]["Url"]);
          }
          A.store.Ref(ptr + 0 + 76, x["Cellular"]["RoamingState"]);
          A.store.Bool(ptr + 0 + 123, "Scanning" in x["Cellular"] ? true : false);
          A.store.Bool(ptr + 0 + 80, x["Cellular"]["Scanning"] ? true : false);

          if (typeof x["Cellular"]["ServingOperator"] === "undefined") {
            A.store.Bool(ptr + 0 + 84 + 12, false);
            A.store.Ref(ptr + 0 + 84 + 0, undefined);
            A.store.Ref(ptr + 0 + 84 + 4, undefined);
            A.store.Ref(ptr + 0 + 84 + 8, undefined);
          } else {
            A.store.Bool(ptr + 0 + 84 + 12, true);
            A.store.Ref(ptr + 0 + 84 + 0, x["Cellular"]["ServingOperator"]["Name"]);
            A.store.Ref(ptr + 0 + 84 + 4, x["Cellular"]["ServingOperator"]["Code"]);
            A.store.Ref(ptr + 0 + 84 + 8, x["Cellular"]["ServingOperator"]["Country"]);
          }

          if (typeof x["Cellular"]["SIMLockStatus"] === "undefined") {
            A.store.Bool(ptr + 0 + 100 + 14, false);
            A.store.Ref(ptr + 0 + 100 + 0, undefined);
            A.store.Bool(ptr + 0 + 100 + 12, false);
            A.store.Bool(ptr + 0 + 100 + 4, false);
            A.store.Bool(ptr + 0 + 100 + 13, false);
            A.store.Int32(ptr + 0 + 100 + 8, 0);
          } else {
            A.store.Bool(ptr + 0 + 100 + 14, true);
            A.store.Ref(ptr + 0 + 100 + 0, x["Cellular"]["SIMLockStatus"]["LockType"]);
            A.store.Bool(ptr + 0 + 100 + 12, "LockEnabled" in x["Cellular"]["SIMLockStatus"] ? true : false);
            A.store.Bool(ptr + 0 + 100 + 4, x["Cellular"]["SIMLockStatus"]["LockEnabled"] ? true : false);
            A.store.Bool(ptr + 0 + 100 + 13, "RetriesLeft" in x["Cellular"]["SIMLockStatus"] ? true : false);
            A.store.Int32(
              ptr + 0 + 100 + 8,
              x["Cellular"]["SIMLockStatus"]["RetriesLeft"] === undefined
                ? 0
                : (x["Cellular"]["SIMLockStatus"]["RetriesLeft"] as number)
            );
          }
          A.store.Bool(ptr + 0 + 124, "SIMPresent" in x["Cellular"] ? true : false);
          A.store.Bool(ptr + 0 + 115, x["Cellular"]["SIMPresent"] ? true : false);
          A.store.Bool(ptr + 0 + 125, "SignalStrength" in x["Cellular"] ? true : false);
          A.store.Int32(
            ptr + 0 + 116,
            x["Cellular"]["SignalStrength"] === undefined ? 0 : (x["Cellular"]["SignalStrength"] as number)
          );
          A.store.Bool(ptr + 0 + 126, "SupportNetworkScan" in x["Cellular"] ? true : false);
          A.store.Bool(ptr + 0 + 120, x["Cellular"]["SupportNetworkScan"] ? true : false);
        }

        if (typeof x["Ethernet"] === "undefined") {
          A.store.Bool(ptr + 128 + 140, false);
          A.store.Bool(ptr + 128 + 139, false);
          A.store.Bool(ptr + 128 + 0, false);
          A.store.Ref(ptr + 128 + 4, undefined);

          A.store.Bool(ptr + 128 + 8 + 130, false);
          A.store.Ref(ptr + 128 + 8 + 0, undefined);

          A.store.Bool(ptr + 128 + 8 + 4 + 45, false);
          A.store.Ref(ptr + 128 + 8 + 4 + 0, undefined);

          A.store.Bool(ptr + 128 + 8 + 4 + 4 + 16, false);
          A.store.Ref(ptr + 128 + 8 + 4 + 4 + 0, undefined);
          A.store.Ref(ptr + 128 + 8 + 4 + 4 + 4, undefined);
          A.store.Ref(ptr + 128 + 8 + 4 + 4 + 8, undefined);
          A.store.Ref(ptr + 128 + 8 + 4 + 4 + 12, undefined);
          A.store.Ref(ptr + 128 + 8 + 4 + 24, undefined);

          A.store.Bool(ptr + 128 + 8 + 4 + 28 + 16, false);
          A.store.Ref(ptr + 128 + 8 + 4 + 28 + 0, undefined);
          A.store.Ref(ptr + 128 + 8 + 4 + 28 + 4, undefined);
          A.store.Ref(ptr + 128 + 8 + 4 + 28 + 8, undefined);
          A.store.Ref(ptr + 128 + 8 + 4 + 28 + 12, undefined);
          A.store.Ref(ptr + 128 + 8 + 52, undefined);
          A.store.Ref(ptr + 128 + 8 + 56, undefined);
          A.store.Ref(ptr + 128 + 8 + 60, undefined);
          A.store.Enum(ptr + 128 + 8 + 64, -1);
          A.store.Ref(ptr + 128 + 8 + 68, undefined);
          A.store.Ref(ptr + 128 + 8 + 72, undefined);
          A.store.Ref(ptr + 128 + 8 + 76, undefined);
          A.store.Ref(ptr + 128 + 8 + 80, undefined);
          A.store.Bool(ptr + 128 + 8 + 127, false);
          A.store.Bool(ptr + 128 + 8 + 84, false);
          A.store.Ref(ptr + 128 + 8 + 88, undefined);
          A.store.Ref(ptr + 128 + 8 + 92, undefined);

          A.store.Bool(ptr + 128 + 8 + 96 + 28, false);
          A.store.Ref(ptr + 128 + 8 + 96 + 0, undefined);
          A.store.Ref(ptr + 128 + 8 + 96 + 4, undefined);
          A.store.Ref(ptr + 128 + 8 + 96 + 8, undefined);
          A.store.Ref(ptr + 128 + 8 + 96 + 12, undefined);
          A.store.Ref(ptr + 128 + 8 + 96 + 16, undefined);
          A.store.Ref(ptr + 128 + 8 + 96 + 20, undefined);
          A.store.Bool(ptr + 128 + 8 + 96 + 26, false);
          A.store.Bool(ptr + 128 + 8 + 96 + 24, false);
          A.store.Bool(ptr + 128 + 8 + 96 + 27, false);
          A.store.Bool(ptr + 128 + 8 + 96 + 25, false);
          A.store.Bool(ptr + 128 + 8 + 128, false);
          A.store.Bool(ptr + 128 + 8 + 125, false);
          A.store.Bool(ptr + 128 + 8 + 129, false);
          A.store.Bool(ptr + 128 + 8 + 126, false);
        } else {
          A.store.Bool(ptr + 128 + 140, true);
          A.store.Bool(ptr + 128 + 139, "AutoConnect" in x["Ethernet"] ? true : false);
          A.store.Bool(ptr + 128 + 0, x["Ethernet"]["AutoConnect"] ? true : false);
          A.store.Ref(ptr + 128 + 4, x["Ethernet"]["Authentication"]);

          if (typeof x["Ethernet"]["EAP"] === "undefined") {
            A.store.Bool(ptr + 128 + 8 + 130, false);
            A.store.Ref(ptr + 128 + 8 + 0, undefined);

            A.store.Bool(ptr + 128 + 8 + 4 + 45, false);
            A.store.Ref(ptr + 128 + 8 + 4 + 0, undefined);

            A.store.Bool(ptr + 128 + 8 + 4 + 4 + 16, false);
            A.store.Ref(ptr + 128 + 8 + 4 + 4 + 0, undefined);
            A.store.Ref(ptr + 128 + 8 + 4 + 4 + 4, undefined);
            A.store.Ref(ptr + 128 + 8 + 4 + 4 + 8, undefined);
            A.store.Ref(ptr + 128 + 8 + 4 + 4 + 12, undefined);
            A.store.Ref(ptr + 128 + 8 + 4 + 24, undefined);

            A.store.Bool(ptr + 128 + 8 + 4 + 28 + 16, false);
            A.store.Ref(ptr + 128 + 8 + 4 + 28 + 0, undefined);
            A.store.Ref(ptr + 128 + 8 + 4 + 28 + 4, undefined);
            A.store.Ref(ptr + 128 + 8 + 4 + 28 + 8, undefined);
            A.store.Ref(ptr + 128 + 8 + 4 + 28 + 12, undefined);
            A.store.Ref(ptr + 128 + 8 + 52, undefined);
            A.store.Ref(ptr + 128 + 8 + 56, undefined);
            A.store.Ref(ptr + 128 + 8 + 60, undefined);
            A.store.Enum(ptr + 128 + 8 + 64, -1);
            A.store.Ref(ptr + 128 + 8 + 68, undefined);
            A.store.Ref(ptr + 128 + 8 + 72, undefined);
            A.store.Ref(ptr + 128 + 8 + 76, undefined);
            A.store.Ref(ptr + 128 + 8 + 80, undefined);
            A.store.Bool(ptr + 128 + 8 + 127, false);
            A.store.Bool(ptr + 128 + 8 + 84, false);
            A.store.Ref(ptr + 128 + 8 + 88, undefined);
            A.store.Ref(ptr + 128 + 8 + 92, undefined);

            A.store.Bool(ptr + 128 + 8 + 96 + 28, false);
            A.store.Ref(ptr + 128 + 8 + 96 + 0, undefined);
            A.store.Ref(ptr + 128 + 8 + 96 + 4, undefined);
            A.store.Ref(ptr + 128 + 8 + 96 + 8, undefined);
            A.store.Ref(ptr + 128 + 8 + 96 + 12, undefined);
            A.store.Ref(ptr + 128 + 8 + 96 + 16, undefined);
            A.store.Ref(ptr + 128 + 8 + 96 + 20, undefined);
            A.store.Bool(ptr + 128 + 8 + 96 + 26, false);
            A.store.Bool(ptr + 128 + 8 + 96 + 24, false);
            A.store.Bool(ptr + 128 + 8 + 96 + 27, false);
            A.store.Bool(ptr + 128 + 8 + 96 + 25, false);
            A.store.Bool(ptr + 128 + 8 + 128, false);
            A.store.Bool(ptr + 128 + 8 + 125, false);
            A.store.Bool(ptr + 128 + 8 + 129, false);
            A.store.Bool(ptr + 128 + 8 + 126, false);
          } else {
            A.store.Bool(ptr + 128 + 8 + 130, true);
            A.store.Ref(ptr + 128 + 8 + 0, x["Ethernet"]["EAP"]["AnonymousIdentity"]);

            if (typeof x["Ethernet"]["EAP"]["ClientCertPattern"] === "undefined") {
              A.store.Bool(ptr + 128 + 8 + 4 + 45, false);
              A.store.Ref(ptr + 128 + 8 + 4 + 0, undefined);

              A.store.Bool(ptr + 128 + 8 + 4 + 4 + 16, false);
              A.store.Ref(ptr + 128 + 8 + 4 + 4 + 0, undefined);
              A.store.Ref(ptr + 128 + 8 + 4 + 4 + 4, undefined);
              A.store.Ref(ptr + 128 + 8 + 4 + 4 + 8, undefined);
              A.store.Ref(ptr + 128 + 8 + 4 + 4 + 12, undefined);
              A.store.Ref(ptr + 128 + 8 + 4 + 24, undefined);

              A.store.Bool(ptr + 128 + 8 + 4 + 28 + 16, false);
              A.store.Ref(ptr + 128 + 8 + 4 + 28 + 0, undefined);
              A.store.Ref(ptr + 128 + 8 + 4 + 28 + 4, undefined);
              A.store.Ref(ptr + 128 + 8 + 4 + 28 + 8, undefined);
              A.store.Ref(ptr + 128 + 8 + 4 + 28 + 12, undefined);
            } else {
              A.store.Bool(ptr + 128 + 8 + 4 + 45, true);
              A.store.Ref(ptr + 128 + 8 + 4 + 0, x["Ethernet"]["EAP"]["ClientCertPattern"]["EnrollmentURI"]);

              if (typeof x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"] === "undefined") {
                A.store.Bool(ptr + 128 + 8 + 4 + 4 + 16, false);
                A.store.Ref(ptr + 128 + 8 + 4 + 4 + 0, undefined);
                A.store.Ref(ptr + 128 + 8 + 4 + 4 + 4, undefined);
                A.store.Ref(ptr + 128 + 8 + 4 + 4 + 8, undefined);
                A.store.Ref(ptr + 128 + 8 + 4 + 4 + 12, undefined);
              } else {
                A.store.Bool(ptr + 128 + 8 + 4 + 4 + 16, true);
                A.store.Ref(
                  ptr + 128 + 8 + 4 + 4 + 0,
                  x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"]
                );
                A.store.Ref(ptr + 128 + 8 + 4 + 4 + 4, x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"]);
                A.store.Ref(
                  ptr + 128 + 8 + 4 + 4 + 8,
                  x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"]
                );
                A.store.Ref(
                  ptr + 128 + 8 + 4 + 4 + 12,
                  x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"]
                );
              }
              A.store.Ref(ptr + 128 + 8 + 4 + 24, x["Ethernet"]["EAP"]["ClientCertPattern"]["IssuerCARef"]);

              if (typeof x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"] === "undefined") {
                A.store.Bool(ptr + 128 + 8 + 4 + 28 + 16, false);
                A.store.Ref(ptr + 128 + 8 + 4 + 28 + 0, undefined);
                A.store.Ref(ptr + 128 + 8 + 4 + 28 + 4, undefined);
                A.store.Ref(ptr + 128 + 8 + 4 + 28 + 8, undefined);
                A.store.Ref(ptr + 128 + 8 + 4 + 28 + 12, undefined);
              } else {
                A.store.Bool(ptr + 128 + 8 + 4 + 28 + 16, true);
                A.store.Ref(
                  ptr + 128 + 8 + 4 + 28 + 0,
                  x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"]
                );
                A.store.Ref(
                  ptr + 128 + 8 + 4 + 28 + 4,
                  x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"]
                );
                A.store.Ref(
                  ptr + 128 + 8 + 4 + 28 + 8,
                  x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"]
                );
                A.store.Ref(
                  ptr + 128 + 8 + 4 + 28 + 12,
                  x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"]
                );
              }
            }
            A.store.Ref(ptr + 128 + 8 + 52, x["Ethernet"]["EAP"]["ClientCertPKCS11Id"]);
            A.store.Ref(ptr + 128 + 8 + 56, x["Ethernet"]["EAP"]["ClientCertProvisioningProfileId"]);
            A.store.Ref(ptr + 128 + 8 + 60, x["Ethernet"]["EAP"]["ClientCertRef"]);
            A.store.Enum(
              ptr + 128 + 8 + 64,
              ["Ref", "Pattern"].indexOf(x["Ethernet"]["EAP"]["ClientCertType"] as string)
            );
            A.store.Ref(ptr + 128 + 8 + 68, x["Ethernet"]["EAP"]["Identity"]);
            A.store.Ref(ptr + 128 + 8 + 72, x["Ethernet"]["EAP"]["Inner"]);
            A.store.Ref(ptr + 128 + 8 + 76, x["Ethernet"]["EAP"]["Outer"]);
            A.store.Ref(ptr + 128 + 8 + 80, x["Ethernet"]["EAP"]["Password"]);
            A.store.Bool(ptr + 128 + 8 + 127, "SaveCredentials" in x["Ethernet"]["EAP"] ? true : false);
            A.store.Bool(ptr + 128 + 8 + 84, x["Ethernet"]["EAP"]["SaveCredentials"] ? true : false);
            A.store.Ref(ptr + 128 + 8 + 88, x["Ethernet"]["EAP"]["ServerCAPEMs"]);
            A.store.Ref(ptr + 128 + 8 + 92, x["Ethernet"]["EAP"]["ServerCARefs"]);

            if (typeof x["Ethernet"]["EAP"]["SubjectMatch"] === "undefined") {
              A.store.Bool(ptr + 128 + 8 + 96 + 28, false);
              A.store.Ref(ptr + 128 + 8 + 96 + 0, undefined);
              A.store.Ref(ptr + 128 + 8 + 96 + 4, undefined);
              A.store.Ref(ptr + 128 + 8 + 96 + 8, undefined);
              A.store.Ref(ptr + 128 + 8 + 96 + 12, undefined);
              A.store.Ref(ptr + 128 + 8 + 96 + 16, undefined);
              A.store.Ref(ptr + 128 + 8 + 96 + 20, undefined);
              A.store.Bool(ptr + 128 + 8 + 96 + 26, false);
              A.store.Bool(ptr + 128 + 8 + 96 + 24, false);
              A.store.Bool(ptr + 128 + 8 + 96 + 27, false);
              A.store.Bool(ptr + 128 + 8 + 96 + 25, false);
            } else {
              A.store.Bool(ptr + 128 + 8 + 96 + 28, true);
              A.store.Ref(ptr + 128 + 8 + 96 + 0, x["Ethernet"]["EAP"]["SubjectMatch"]["Active"]);
              A.store.Ref(ptr + 128 + 8 + 96 + 4, x["Ethernet"]["EAP"]["SubjectMatch"]["Effective"]);
              A.store.Ref(ptr + 128 + 8 + 96 + 8, x["Ethernet"]["EAP"]["SubjectMatch"]["UserPolicy"]);
              A.store.Ref(ptr + 128 + 8 + 96 + 12, x["Ethernet"]["EAP"]["SubjectMatch"]["DevicePolicy"]);
              A.store.Ref(ptr + 128 + 8 + 96 + 16, x["Ethernet"]["EAP"]["SubjectMatch"]["UserSetting"]);
              A.store.Ref(ptr + 128 + 8 + 96 + 20, x["Ethernet"]["EAP"]["SubjectMatch"]["SharedSetting"]);
              A.store.Bool(
                ptr + 128 + 8 + 96 + 26,
                "UserEditable" in x["Ethernet"]["EAP"]["SubjectMatch"] ? true : false
              );
              A.store.Bool(
                ptr + 128 + 8 + 96 + 24,
                x["Ethernet"]["EAP"]["SubjectMatch"]["UserEditable"] ? true : false
              );
              A.store.Bool(
                ptr + 128 + 8 + 96 + 27,
                "DeviceEditable" in x["Ethernet"]["EAP"]["SubjectMatch"] ? true : false
              );
              A.store.Bool(
                ptr + 128 + 8 + 96 + 25,
                x["Ethernet"]["EAP"]["SubjectMatch"]["DeviceEditable"] ? true : false
              );
            }
            A.store.Bool(ptr + 128 + 8 + 128, "UseProactiveKeyCaching" in x["Ethernet"]["EAP"] ? true : false);
            A.store.Bool(ptr + 128 + 8 + 125, x["Ethernet"]["EAP"]["UseProactiveKeyCaching"] ? true : false);
            A.store.Bool(ptr + 128 + 8 + 129, "UseSystemCAs" in x["Ethernet"]["EAP"] ? true : false);
            A.store.Bool(ptr + 128 + 8 + 126, x["Ethernet"]["EAP"]["UseSystemCAs"] ? true : false);
          }
        }
        A.store.Ref(ptr + 272, x["GUID"]);
        A.store.Enum(ptr + 276, ["DHCP", "Static"].indexOf(x["IPAddressConfigType"] as string));
        A.store.Ref(ptr + 280, x["Name"]);
        A.store.Enum(ptr + 284, ["DHCP", "Static"].indexOf(x["NameServersConfigType"] as string));
        A.store.Bool(ptr + 633, "Priority" in x ? true : false);
        A.store.Int32(ptr + 288, x["Priority"] === undefined ? 0 : (x["Priority"] as number));
        A.store.Enum(
          ptr + 292,
          ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"].indexOf(x["Type"] as string)
        );

        if (typeof x["VPN"] === "undefined") {
          A.store.Bool(ptr + 296 + 13, false);
          A.store.Bool(ptr + 296 + 12, false);
          A.store.Bool(ptr + 296 + 0, false);
          A.store.Ref(ptr + 296 + 4, undefined);
          A.store.Ref(ptr + 296 + 8, undefined);
        } else {
          A.store.Bool(ptr + 296 + 13, true);
          A.store.Bool(ptr + 296 + 12, "AutoConnect" in x["VPN"] ? true : false);
          A.store.Bool(ptr + 296 + 0, x["VPN"]["AutoConnect"] ? true : false);
          A.store.Ref(ptr + 296 + 4, x["VPN"]["Host"]);
          A.store.Ref(ptr + 296 + 8, x["VPN"]["Type"]);
        }

        if (typeof x["WiFi"] === "undefined") {
          A.store.Bool(ptr + 312 + 182, false);
          A.store.Bool(ptr + 312 + 176, false);
          A.store.Bool(ptr + 312 + 0, false);
          A.store.Bool(ptr + 312 + 177, false);
          A.store.Bool(ptr + 312 + 1, false);
          A.store.Ref(ptr + 312 + 4, undefined);

          A.store.Bool(ptr + 312 + 8 + 130, false);
          A.store.Ref(ptr + 312 + 8 + 0, undefined);

          A.store.Bool(ptr + 312 + 8 + 4 + 45, false);
          A.store.Ref(ptr + 312 + 8 + 4 + 0, undefined);

          A.store.Bool(ptr + 312 + 8 + 4 + 4 + 16, false);
          A.store.Ref(ptr + 312 + 8 + 4 + 4 + 0, undefined);
          A.store.Ref(ptr + 312 + 8 + 4 + 4 + 4, undefined);
          A.store.Ref(ptr + 312 + 8 + 4 + 4 + 8, undefined);
          A.store.Ref(ptr + 312 + 8 + 4 + 4 + 12, undefined);
          A.store.Ref(ptr + 312 + 8 + 4 + 24, undefined);

          A.store.Bool(ptr + 312 + 8 + 4 + 28 + 16, false);
          A.store.Ref(ptr + 312 + 8 + 4 + 28 + 0, undefined);
          A.store.Ref(ptr + 312 + 8 + 4 + 28 + 4, undefined);
          A.store.Ref(ptr + 312 + 8 + 4 + 28 + 8, undefined);
          A.store.Ref(ptr + 312 + 8 + 4 + 28 + 12, undefined);
          A.store.Ref(ptr + 312 + 8 + 52, undefined);
          A.store.Ref(ptr + 312 + 8 + 56, undefined);
          A.store.Ref(ptr + 312 + 8 + 60, undefined);
          A.store.Enum(ptr + 312 + 8 + 64, -1);
          A.store.Ref(ptr + 312 + 8 + 68, undefined);
          A.store.Ref(ptr + 312 + 8 + 72, undefined);
          A.store.Ref(ptr + 312 + 8 + 76, undefined);
          A.store.Ref(ptr + 312 + 8 + 80, undefined);
          A.store.Bool(ptr + 312 + 8 + 127, false);
          A.store.Bool(ptr + 312 + 8 + 84, false);
          A.store.Ref(ptr + 312 + 8 + 88, undefined);
          A.store.Ref(ptr + 312 + 8 + 92, undefined);

          A.store.Bool(ptr + 312 + 8 + 96 + 28, false);
          A.store.Ref(ptr + 312 + 8 + 96 + 0, undefined);
          A.store.Ref(ptr + 312 + 8 + 96 + 4, undefined);
          A.store.Ref(ptr + 312 + 8 + 96 + 8, undefined);
          A.store.Ref(ptr + 312 + 8 + 96 + 12, undefined);
          A.store.Ref(ptr + 312 + 8 + 96 + 16, undefined);
          A.store.Ref(ptr + 312 + 8 + 96 + 20, undefined);
          A.store.Bool(ptr + 312 + 8 + 96 + 26, false);
          A.store.Bool(ptr + 312 + 8 + 96 + 24, false);
          A.store.Bool(ptr + 312 + 8 + 96 + 27, false);
          A.store.Bool(ptr + 312 + 8 + 96 + 25, false);
          A.store.Bool(ptr + 312 + 8 + 128, false);
          A.store.Bool(ptr + 312 + 8 + 125, false);
          A.store.Bool(ptr + 312 + 8 + 129, false);
          A.store.Bool(ptr + 312 + 8 + 126, false);
          A.store.Bool(ptr + 312 + 178, false);
          A.store.Int32(ptr + 312 + 140, 0);
          A.store.Ref(ptr + 312 + 144, undefined);
          A.store.Ref(ptr + 312 + 148, undefined);
          A.store.Bool(ptr + 312 + 179, false);
          A.store.Bool(ptr + 312 + 152, false);
          A.store.Ref(ptr + 312 + 156, undefined);
          A.store.Bool(ptr + 312 + 180, false);
          A.store.Int32(ptr + 312 + 160, 0);
          A.store.Ref(ptr + 312 + 164, undefined);
          A.store.Ref(ptr + 312 + 168, undefined);
          A.store.Bool(ptr + 312 + 181, false);
          A.store.Int32(ptr + 312 + 172, 0);
        } else {
          A.store.Bool(ptr + 312 + 182, true);
          A.store.Bool(ptr + 312 + 176, "AllowGatewayARPPolling" in x["WiFi"] ? true : false);
          A.store.Bool(ptr + 312 + 0, x["WiFi"]["AllowGatewayARPPolling"] ? true : false);
          A.store.Bool(ptr + 312 + 177, "AutoConnect" in x["WiFi"] ? true : false);
          A.store.Bool(ptr + 312 + 1, x["WiFi"]["AutoConnect"] ? true : false);
          A.store.Ref(ptr + 312 + 4, x["WiFi"]["BSSID"]);

          if (typeof x["WiFi"]["EAP"] === "undefined") {
            A.store.Bool(ptr + 312 + 8 + 130, false);
            A.store.Ref(ptr + 312 + 8 + 0, undefined);

            A.store.Bool(ptr + 312 + 8 + 4 + 45, false);
            A.store.Ref(ptr + 312 + 8 + 4 + 0, undefined);

            A.store.Bool(ptr + 312 + 8 + 4 + 4 + 16, false);
            A.store.Ref(ptr + 312 + 8 + 4 + 4 + 0, undefined);
            A.store.Ref(ptr + 312 + 8 + 4 + 4 + 4, undefined);
            A.store.Ref(ptr + 312 + 8 + 4 + 4 + 8, undefined);
            A.store.Ref(ptr + 312 + 8 + 4 + 4 + 12, undefined);
            A.store.Ref(ptr + 312 + 8 + 4 + 24, undefined);

            A.store.Bool(ptr + 312 + 8 + 4 + 28 + 16, false);
            A.store.Ref(ptr + 312 + 8 + 4 + 28 + 0, undefined);
            A.store.Ref(ptr + 312 + 8 + 4 + 28 + 4, undefined);
            A.store.Ref(ptr + 312 + 8 + 4 + 28 + 8, undefined);
            A.store.Ref(ptr + 312 + 8 + 4 + 28 + 12, undefined);
            A.store.Ref(ptr + 312 + 8 + 52, undefined);
            A.store.Ref(ptr + 312 + 8 + 56, undefined);
            A.store.Ref(ptr + 312 + 8 + 60, undefined);
            A.store.Enum(ptr + 312 + 8 + 64, -1);
            A.store.Ref(ptr + 312 + 8 + 68, undefined);
            A.store.Ref(ptr + 312 + 8 + 72, undefined);
            A.store.Ref(ptr + 312 + 8 + 76, undefined);
            A.store.Ref(ptr + 312 + 8 + 80, undefined);
            A.store.Bool(ptr + 312 + 8 + 127, false);
            A.store.Bool(ptr + 312 + 8 + 84, false);
            A.store.Ref(ptr + 312 + 8 + 88, undefined);
            A.store.Ref(ptr + 312 + 8 + 92, undefined);

            A.store.Bool(ptr + 312 + 8 + 96 + 28, false);
            A.store.Ref(ptr + 312 + 8 + 96 + 0, undefined);
            A.store.Ref(ptr + 312 + 8 + 96 + 4, undefined);
            A.store.Ref(ptr + 312 + 8 + 96 + 8, undefined);
            A.store.Ref(ptr + 312 + 8 + 96 + 12, undefined);
            A.store.Ref(ptr + 312 + 8 + 96 + 16, undefined);
            A.store.Ref(ptr + 312 + 8 + 96 + 20, undefined);
            A.store.Bool(ptr + 312 + 8 + 96 + 26, false);
            A.store.Bool(ptr + 312 + 8 + 96 + 24, false);
            A.store.Bool(ptr + 312 + 8 + 96 + 27, false);
            A.store.Bool(ptr + 312 + 8 + 96 + 25, false);
            A.store.Bool(ptr + 312 + 8 + 128, false);
            A.store.Bool(ptr + 312 + 8 + 125, false);
            A.store.Bool(ptr + 312 + 8 + 129, false);
            A.store.Bool(ptr + 312 + 8 + 126, false);
          } else {
            A.store.Bool(ptr + 312 + 8 + 130, true);
            A.store.Ref(ptr + 312 + 8 + 0, x["WiFi"]["EAP"]["AnonymousIdentity"]);

            if (typeof x["WiFi"]["EAP"]["ClientCertPattern"] === "undefined") {
              A.store.Bool(ptr + 312 + 8 + 4 + 45, false);
              A.store.Ref(ptr + 312 + 8 + 4 + 0, undefined);

              A.store.Bool(ptr + 312 + 8 + 4 + 4 + 16, false);
              A.store.Ref(ptr + 312 + 8 + 4 + 4 + 0, undefined);
              A.store.Ref(ptr + 312 + 8 + 4 + 4 + 4, undefined);
              A.store.Ref(ptr + 312 + 8 + 4 + 4 + 8, undefined);
              A.store.Ref(ptr + 312 + 8 + 4 + 4 + 12, undefined);
              A.store.Ref(ptr + 312 + 8 + 4 + 24, undefined);

              A.store.Bool(ptr + 312 + 8 + 4 + 28 + 16, false);
              A.store.Ref(ptr + 312 + 8 + 4 + 28 + 0, undefined);
              A.store.Ref(ptr + 312 + 8 + 4 + 28 + 4, undefined);
              A.store.Ref(ptr + 312 + 8 + 4 + 28 + 8, undefined);
              A.store.Ref(ptr + 312 + 8 + 4 + 28 + 12, undefined);
            } else {
              A.store.Bool(ptr + 312 + 8 + 4 + 45, true);
              A.store.Ref(ptr + 312 + 8 + 4 + 0, x["WiFi"]["EAP"]["ClientCertPattern"]["EnrollmentURI"]);

              if (typeof x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"] === "undefined") {
                A.store.Bool(ptr + 312 + 8 + 4 + 4 + 16, false);
                A.store.Ref(ptr + 312 + 8 + 4 + 4 + 0, undefined);
                A.store.Ref(ptr + 312 + 8 + 4 + 4 + 4, undefined);
                A.store.Ref(ptr + 312 + 8 + 4 + 4 + 8, undefined);
                A.store.Ref(ptr + 312 + 8 + 4 + 4 + 12, undefined);
              } else {
                A.store.Bool(ptr + 312 + 8 + 4 + 4 + 16, true);
                A.store.Ref(ptr + 312 + 8 + 4 + 4 + 0, x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"]);
                A.store.Ref(ptr + 312 + 8 + 4 + 4 + 4, x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"]);
                A.store.Ref(ptr + 312 + 8 + 4 + 4 + 8, x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"]);
                A.store.Ref(
                  ptr + 312 + 8 + 4 + 4 + 12,
                  x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"]
                );
              }
              A.store.Ref(ptr + 312 + 8 + 4 + 24, x["WiFi"]["EAP"]["ClientCertPattern"]["IssuerCARef"]);

              if (typeof x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"] === "undefined") {
                A.store.Bool(ptr + 312 + 8 + 4 + 28 + 16, false);
                A.store.Ref(ptr + 312 + 8 + 4 + 28 + 0, undefined);
                A.store.Ref(ptr + 312 + 8 + 4 + 28 + 4, undefined);
                A.store.Ref(ptr + 312 + 8 + 4 + 28 + 8, undefined);
                A.store.Ref(ptr + 312 + 8 + 4 + 28 + 12, undefined);
              } else {
                A.store.Bool(ptr + 312 + 8 + 4 + 28 + 16, true);
                A.store.Ref(ptr + 312 + 8 + 4 + 28 + 0, x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"]);
                A.store.Ref(ptr + 312 + 8 + 4 + 28 + 4, x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"]);
                A.store.Ref(
                  ptr + 312 + 8 + 4 + 28 + 8,
                  x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"]
                );
                A.store.Ref(
                  ptr + 312 + 8 + 4 + 28 + 12,
                  x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"]
                );
              }
            }
            A.store.Ref(ptr + 312 + 8 + 52, x["WiFi"]["EAP"]["ClientCertPKCS11Id"]);
            A.store.Ref(ptr + 312 + 8 + 56, x["WiFi"]["EAP"]["ClientCertProvisioningProfileId"]);
            A.store.Ref(ptr + 312 + 8 + 60, x["WiFi"]["EAP"]["ClientCertRef"]);
            A.store.Enum(ptr + 312 + 8 + 64, ["Ref", "Pattern"].indexOf(x["WiFi"]["EAP"]["ClientCertType"] as string));
            A.store.Ref(ptr + 312 + 8 + 68, x["WiFi"]["EAP"]["Identity"]);
            A.store.Ref(ptr + 312 + 8 + 72, x["WiFi"]["EAP"]["Inner"]);
            A.store.Ref(ptr + 312 + 8 + 76, x["WiFi"]["EAP"]["Outer"]);
            A.store.Ref(ptr + 312 + 8 + 80, x["WiFi"]["EAP"]["Password"]);
            A.store.Bool(ptr + 312 + 8 + 127, "SaveCredentials" in x["WiFi"]["EAP"] ? true : false);
            A.store.Bool(ptr + 312 + 8 + 84, x["WiFi"]["EAP"]["SaveCredentials"] ? true : false);
            A.store.Ref(ptr + 312 + 8 + 88, x["WiFi"]["EAP"]["ServerCAPEMs"]);
            A.store.Ref(ptr + 312 + 8 + 92, x["WiFi"]["EAP"]["ServerCARefs"]);

            if (typeof x["WiFi"]["EAP"]["SubjectMatch"] === "undefined") {
              A.store.Bool(ptr + 312 + 8 + 96 + 28, false);
              A.store.Ref(ptr + 312 + 8 + 96 + 0, undefined);
              A.store.Ref(ptr + 312 + 8 + 96 + 4, undefined);
              A.store.Ref(ptr + 312 + 8 + 96 + 8, undefined);
              A.store.Ref(ptr + 312 + 8 + 96 + 12, undefined);
              A.store.Ref(ptr + 312 + 8 + 96 + 16, undefined);
              A.store.Ref(ptr + 312 + 8 + 96 + 20, undefined);
              A.store.Bool(ptr + 312 + 8 + 96 + 26, false);
              A.store.Bool(ptr + 312 + 8 + 96 + 24, false);
              A.store.Bool(ptr + 312 + 8 + 96 + 27, false);
              A.store.Bool(ptr + 312 + 8 + 96 + 25, false);
            } else {
              A.store.Bool(ptr + 312 + 8 + 96 + 28, true);
              A.store.Ref(ptr + 312 + 8 + 96 + 0, x["WiFi"]["EAP"]["SubjectMatch"]["Active"]);
              A.store.Ref(ptr + 312 + 8 + 96 + 4, x["WiFi"]["EAP"]["SubjectMatch"]["Effective"]);
              A.store.Ref(ptr + 312 + 8 + 96 + 8, x["WiFi"]["EAP"]["SubjectMatch"]["UserPolicy"]);
              A.store.Ref(ptr + 312 + 8 + 96 + 12, x["WiFi"]["EAP"]["SubjectMatch"]["DevicePolicy"]);
              A.store.Ref(ptr + 312 + 8 + 96 + 16, x["WiFi"]["EAP"]["SubjectMatch"]["UserSetting"]);
              A.store.Ref(ptr + 312 + 8 + 96 + 20, x["WiFi"]["EAP"]["SubjectMatch"]["SharedSetting"]);
              A.store.Bool(ptr + 312 + 8 + 96 + 26, "UserEditable" in x["WiFi"]["EAP"]["SubjectMatch"] ? true : false);
              A.store.Bool(ptr + 312 + 8 + 96 + 24, x["WiFi"]["EAP"]["SubjectMatch"]["UserEditable"] ? true : false);
              A.store.Bool(
                ptr + 312 + 8 + 96 + 27,
                "DeviceEditable" in x["WiFi"]["EAP"]["SubjectMatch"] ? true : false
              );
              A.store.Bool(ptr + 312 + 8 + 96 + 25, x["WiFi"]["EAP"]["SubjectMatch"]["DeviceEditable"] ? true : false);
            }
            A.store.Bool(ptr + 312 + 8 + 128, "UseProactiveKeyCaching" in x["WiFi"]["EAP"] ? true : false);
            A.store.Bool(ptr + 312 + 8 + 125, x["WiFi"]["EAP"]["UseProactiveKeyCaching"] ? true : false);
            A.store.Bool(ptr + 312 + 8 + 129, "UseSystemCAs" in x["WiFi"]["EAP"] ? true : false);
            A.store.Bool(ptr + 312 + 8 + 126, x["WiFi"]["EAP"]["UseSystemCAs"] ? true : false);
          }
          A.store.Bool(ptr + 312 + 178, "Frequency" in x["WiFi"] ? true : false);
          A.store.Int32(ptr + 312 + 140, x["WiFi"]["Frequency"] === undefined ? 0 : (x["WiFi"]["Frequency"] as number));
          A.store.Ref(ptr + 312 + 144, x["WiFi"]["FrequencyList"]);
          A.store.Ref(ptr + 312 + 148, x["WiFi"]["HexSSID"]);
          A.store.Bool(ptr + 312 + 179, "HiddenSSID" in x["WiFi"] ? true : false);
          A.store.Bool(ptr + 312 + 152, x["WiFi"]["HiddenSSID"] ? true : false);
          A.store.Ref(ptr + 312 + 156, x["WiFi"]["Passphrase"]);
          A.store.Bool(ptr + 312 + 180, "RoamThreshold" in x["WiFi"] ? true : false);
          A.store.Int32(
            ptr + 312 + 160,
            x["WiFi"]["RoamThreshold"] === undefined ? 0 : (x["WiFi"]["RoamThreshold"] as number)
          );
          A.store.Ref(ptr + 312 + 164, x["WiFi"]["SSID"]);
          A.store.Ref(ptr + 312 + 168, x["WiFi"]["Security"]);
          A.store.Bool(ptr + 312 + 181, "SignalStrength" in x["WiFi"] ? true : false);
          A.store.Int32(
            ptr + 312 + 172,
            x["WiFi"]["SignalStrength"] === undefined ? 0 : (x["WiFi"]["SignalStrength"] as number)
          );
        }

        if (typeof x["WiMAX"] === "undefined") {
          A.store.Bool(ptr + 496 + 136, false);
          A.store.Bool(ptr + 496 + 135, false);
          A.store.Bool(ptr + 496 + 0, false);

          A.store.Bool(ptr + 496 + 4 + 130, false);
          A.store.Ref(ptr + 496 + 4 + 0, undefined);

          A.store.Bool(ptr + 496 + 4 + 4 + 45, false);
          A.store.Ref(ptr + 496 + 4 + 4 + 0, undefined);

          A.store.Bool(ptr + 496 + 4 + 4 + 4 + 16, false);
          A.store.Ref(ptr + 496 + 4 + 4 + 4 + 0, undefined);
          A.store.Ref(ptr + 496 + 4 + 4 + 4 + 4, undefined);
          A.store.Ref(ptr + 496 + 4 + 4 + 4 + 8, undefined);
          A.store.Ref(ptr + 496 + 4 + 4 + 4 + 12, undefined);
          A.store.Ref(ptr + 496 + 4 + 4 + 24, undefined);

          A.store.Bool(ptr + 496 + 4 + 4 + 28 + 16, false);
          A.store.Ref(ptr + 496 + 4 + 4 + 28 + 0, undefined);
          A.store.Ref(ptr + 496 + 4 + 4 + 28 + 4, undefined);
          A.store.Ref(ptr + 496 + 4 + 4 + 28 + 8, undefined);
          A.store.Ref(ptr + 496 + 4 + 4 + 28 + 12, undefined);
          A.store.Ref(ptr + 496 + 4 + 52, undefined);
          A.store.Ref(ptr + 496 + 4 + 56, undefined);
          A.store.Ref(ptr + 496 + 4 + 60, undefined);
          A.store.Enum(ptr + 496 + 4 + 64, -1);
          A.store.Ref(ptr + 496 + 4 + 68, undefined);
          A.store.Ref(ptr + 496 + 4 + 72, undefined);
          A.store.Ref(ptr + 496 + 4 + 76, undefined);
          A.store.Ref(ptr + 496 + 4 + 80, undefined);
          A.store.Bool(ptr + 496 + 4 + 127, false);
          A.store.Bool(ptr + 496 + 4 + 84, false);
          A.store.Ref(ptr + 496 + 4 + 88, undefined);
          A.store.Ref(ptr + 496 + 4 + 92, undefined);

          A.store.Bool(ptr + 496 + 4 + 96 + 28, false);
          A.store.Ref(ptr + 496 + 4 + 96 + 0, undefined);
          A.store.Ref(ptr + 496 + 4 + 96 + 4, undefined);
          A.store.Ref(ptr + 496 + 4 + 96 + 8, undefined);
          A.store.Ref(ptr + 496 + 4 + 96 + 12, undefined);
          A.store.Ref(ptr + 496 + 4 + 96 + 16, undefined);
          A.store.Ref(ptr + 496 + 4 + 96 + 20, undefined);
          A.store.Bool(ptr + 496 + 4 + 96 + 26, false);
          A.store.Bool(ptr + 496 + 4 + 96 + 24, false);
          A.store.Bool(ptr + 496 + 4 + 96 + 27, false);
          A.store.Bool(ptr + 496 + 4 + 96 + 25, false);
          A.store.Bool(ptr + 496 + 4 + 128, false);
          A.store.Bool(ptr + 496 + 4 + 125, false);
          A.store.Bool(ptr + 496 + 4 + 129, false);
          A.store.Bool(ptr + 496 + 4 + 126, false);
        } else {
          A.store.Bool(ptr + 496 + 136, true);
          A.store.Bool(ptr + 496 + 135, "AutoConnect" in x["WiMAX"] ? true : false);
          A.store.Bool(ptr + 496 + 0, x["WiMAX"]["AutoConnect"] ? true : false);

          if (typeof x["WiMAX"]["EAP"] === "undefined") {
            A.store.Bool(ptr + 496 + 4 + 130, false);
            A.store.Ref(ptr + 496 + 4 + 0, undefined);

            A.store.Bool(ptr + 496 + 4 + 4 + 45, false);
            A.store.Ref(ptr + 496 + 4 + 4 + 0, undefined);

            A.store.Bool(ptr + 496 + 4 + 4 + 4 + 16, false);
            A.store.Ref(ptr + 496 + 4 + 4 + 4 + 0, undefined);
            A.store.Ref(ptr + 496 + 4 + 4 + 4 + 4, undefined);
            A.store.Ref(ptr + 496 + 4 + 4 + 4 + 8, undefined);
            A.store.Ref(ptr + 496 + 4 + 4 + 4 + 12, undefined);
            A.store.Ref(ptr + 496 + 4 + 4 + 24, undefined);

            A.store.Bool(ptr + 496 + 4 + 4 + 28 + 16, false);
            A.store.Ref(ptr + 496 + 4 + 4 + 28 + 0, undefined);
            A.store.Ref(ptr + 496 + 4 + 4 + 28 + 4, undefined);
            A.store.Ref(ptr + 496 + 4 + 4 + 28 + 8, undefined);
            A.store.Ref(ptr + 496 + 4 + 4 + 28 + 12, undefined);
            A.store.Ref(ptr + 496 + 4 + 52, undefined);
            A.store.Ref(ptr + 496 + 4 + 56, undefined);
            A.store.Ref(ptr + 496 + 4 + 60, undefined);
            A.store.Enum(ptr + 496 + 4 + 64, -1);
            A.store.Ref(ptr + 496 + 4 + 68, undefined);
            A.store.Ref(ptr + 496 + 4 + 72, undefined);
            A.store.Ref(ptr + 496 + 4 + 76, undefined);
            A.store.Ref(ptr + 496 + 4 + 80, undefined);
            A.store.Bool(ptr + 496 + 4 + 127, false);
            A.store.Bool(ptr + 496 + 4 + 84, false);
            A.store.Ref(ptr + 496 + 4 + 88, undefined);
            A.store.Ref(ptr + 496 + 4 + 92, undefined);

            A.store.Bool(ptr + 496 + 4 + 96 + 28, false);
            A.store.Ref(ptr + 496 + 4 + 96 + 0, undefined);
            A.store.Ref(ptr + 496 + 4 + 96 + 4, undefined);
            A.store.Ref(ptr + 496 + 4 + 96 + 8, undefined);
            A.store.Ref(ptr + 496 + 4 + 96 + 12, undefined);
            A.store.Ref(ptr + 496 + 4 + 96 + 16, undefined);
            A.store.Ref(ptr + 496 + 4 + 96 + 20, undefined);
            A.store.Bool(ptr + 496 + 4 + 96 + 26, false);
            A.store.Bool(ptr + 496 + 4 + 96 + 24, false);
            A.store.Bool(ptr + 496 + 4 + 96 + 27, false);
            A.store.Bool(ptr + 496 + 4 + 96 + 25, false);
            A.store.Bool(ptr + 496 + 4 + 128, false);
            A.store.Bool(ptr + 496 + 4 + 125, false);
            A.store.Bool(ptr + 496 + 4 + 129, false);
            A.store.Bool(ptr + 496 + 4 + 126, false);
          } else {
            A.store.Bool(ptr + 496 + 4 + 130, true);
            A.store.Ref(ptr + 496 + 4 + 0, x["WiMAX"]["EAP"]["AnonymousIdentity"]);

            if (typeof x["WiMAX"]["EAP"]["ClientCertPattern"] === "undefined") {
              A.store.Bool(ptr + 496 + 4 + 4 + 45, false);
              A.store.Ref(ptr + 496 + 4 + 4 + 0, undefined);

              A.store.Bool(ptr + 496 + 4 + 4 + 4 + 16, false);
              A.store.Ref(ptr + 496 + 4 + 4 + 4 + 0, undefined);
              A.store.Ref(ptr + 496 + 4 + 4 + 4 + 4, undefined);
              A.store.Ref(ptr + 496 + 4 + 4 + 4 + 8, undefined);
              A.store.Ref(ptr + 496 + 4 + 4 + 4 + 12, undefined);
              A.store.Ref(ptr + 496 + 4 + 4 + 24, undefined);

              A.store.Bool(ptr + 496 + 4 + 4 + 28 + 16, false);
              A.store.Ref(ptr + 496 + 4 + 4 + 28 + 0, undefined);
              A.store.Ref(ptr + 496 + 4 + 4 + 28 + 4, undefined);
              A.store.Ref(ptr + 496 + 4 + 4 + 28 + 8, undefined);
              A.store.Ref(ptr + 496 + 4 + 4 + 28 + 12, undefined);
            } else {
              A.store.Bool(ptr + 496 + 4 + 4 + 45, true);
              A.store.Ref(ptr + 496 + 4 + 4 + 0, x["WiMAX"]["EAP"]["ClientCertPattern"]["EnrollmentURI"]);

              if (typeof x["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"] === "undefined") {
                A.store.Bool(ptr + 496 + 4 + 4 + 4 + 16, false);
                A.store.Ref(ptr + 496 + 4 + 4 + 4 + 0, undefined);
                A.store.Ref(ptr + 496 + 4 + 4 + 4 + 4, undefined);
                A.store.Ref(ptr + 496 + 4 + 4 + 4 + 8, undefined);
                A.store.Ref(ptr + 496 + 4 + 4 + 4 + 12, undefined);
              } else {
                A.store.Bool(ptr + 496 + 4 + 4 + 4 + 16, true);
                A.store.Ref(ptr + 496 + 4 + 4 + 4 + 0, x["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"]);
                A.store.Ref(ptr + 496 + 4 + 4 + 4 + 4, x["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"]);
                A.store.Ref(
                  ptr + 496 + 4 + 4 + 4 + 8,
                  x["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"]
                );
                A.store.Ref(
                  ptr + 496 + 4 + 4 + 4 + 12,
                  x["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"]
                );
              }
              A.store.Ref(ptr + 496 + 4 + 4 + 24, x["WiMAX"]["EAP"]["ClientCertPattern"]["IssuerCARef"]);

              if (typeof x["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"] === "undefined") {
                A.store.Bool(ptr + 496 + 4 + 4 + 28 + 16, false);
                A.store.Ref(ptr + 496 + 4 + 4 + 28 + 0, undefined);
                A.store.Ref(ptr + 496 + 4 + 4 + 28 + 4, undefined);
                A.store.Ref(ptr + 496 + 4 + 4 + 28 + 8, undefined);
                A.store.Ref(ptr + 496 + 4 + 4 + 28 + 12, undefined);
              } else {
                A.store.Bool(ptr + 496 + 4 + 4 + 28 + 16, true);
                A.store.Ref(
                  ptr + 496 + 4 + 4 + 28 + 0,
                  x["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"]
                );
                A.store.Ref(ptr + 496 + 4 + 4 + 28 + 4, x["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"]);
                A.store.Ref(
                  ptr + 496 + 4 + 4 + 28 + 8,
                  x["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"]
                );
                A.store.Ref(
                  ptr + 496 + 4 + 4 + 28 + 12,
                  x["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"]
                );
              }
            }
            A.store.Ref(ptr + 496 + 4 + 52, x["WiMAX"]["EAP"]["ClientCertPKCS11Id"]);
            A.store.Ref(ptr + 496 + 4 + 56, x["WiMAX"]["EAP"]["ClientCertProvisioningProfileId"]);
            A.store.Ref(ptr + 496 + 4 + 60, x["WiMAX"]["EAP"]["ClientCertRef"]);
            A.store.Enum(ptr + 496 + 4 + 64, ["Ref", "Pattern"].indexOf(x["WiMAX"]["EAP"]["ClientCertType"] as string));
            A.store.Ref(ptr + 496 + 4 + 68, x["WiMAX"]["EAP"]["Identity"]);
            A.store.Ref(ptr + 496 + 4 + 72, x["WiMAX"]["EAP"]["Inner"]);
            A.store.Ref(ptr + 496 + 4 + 76, x["WiMAX"]["EAP"]["Outer"]);
            A.store.Ref(ptr + 496 + 4 + 80, x["WiMAX"]["EAP"]["Password"]);
            A.store.Bool(ptr + 496 + 4 + 127, "SaveCredentials" in x["WiMAX"]["EAP"] ? true : false);
            A.store.Bool(ptr + 496 + 4 + 84, x["WiMAX"]["EAP"]["SaveCredentials"] ? true : false);
            A.store.Ref(ptr + 496 + 4 + 88, x["WiMAX"]["EAP"]["ServerCAPEMs"]);
            A.store.Ref(ptr + 496 + 4 + 92, x["WiMAX"]["EAP"]["ServerCARefs"]);

            if (typeof x["WiMAX"]["EAP"]["SubjectMatch"] === "undefined") {
              A.store.Bool(ptr + 496 + 4 + 96 + 28, false);
              A.store.Ref(ptr + 496 + 4 + 96 + 0, undefined);
              A.store.Ref(ptr + 496 + 4 + 96 + 4, undefined);
              A.store.Ref(ptr + 496 + 4 + 96 + 8, undefined);
              A.store.Ref(ptr + 496 + 4 + 96 + 12, undefined);
              A.store.Ref(ptr + 496 + 4 + 96 + 16, undefined);
              A.store.Ref(ptr + 496 + 4 + 96 + 20, undefined);
              A.store.Bool(ptr + 496 + 4 + 96 + 26, false);
              A.store.Bool(ptr + 496 + 4 + 96 + 24, false);
              A.store.Bool(ptr + 496 + 4 + 96 + 27, false);
              A.store.Bool(ptr + 496 + 4 + 96 + 25, false);
            } else {
              A.store.Bool(ptr + 496 + 4 + 96 + 28, true);
              A.store.Ref(ptr + 496 + 4 + 96 + 0, x["WiMAX"]["EAP"]["SubjectMatch"]["Active"]);
              A.store.Ref(ptr + 496 + 4 + 96 + 4, x["WiMAX"]["EAP"]["SubjectMatch"]["Effective"]);
              A.store.Ref(ptr + 496 + 4 + 96 + 8, x["WiMAX"]["EAP"]["SubjectMatch"]["UserPolicy"]);
              A.store.Ref(ptr + 496 + 4 + 96 + 12, x["WiMAX"]["EAP"]["SubjectMatch"]["DevicePolicy"]);
              A.store.Ref(ptr + 496 + 4 + 96 + 16, x["WiMAX"]["EAP"]["SubjectMatch"]["UserSetting"]);
              A.store.Ref(ptr + 496 + 4 + 96 + 20, x["WiMAX"]["EAP"]["SubjectMatch"]["SharedSetting"]);
              A.store.Bool(ptr + 496 + 4 + 96 + 26, "UserEditable" in x["WiMAX"]["EAP"]["SubjectMatch"] ? true : false);
              A.store.Bool(ptr + 496 + 4 + 96 + 24, x["WiMAX"]["EAP"]["SubjectMatch"]["UserEditable"] ? true : false);
              A.store.Bool(
                ptr + 496 + 4 + 96 + 27,
                "DeviceEditable" in x["WiMAX"]["EAP"]["SubjectMatch"] ? true : false
              );
              A.store.Bool(ptr + 496 + 4 + 96 + 25, x["WiMAX"]["EAP"]["SubjectMatch"]["DeviceEditable"] ? true : false);
            }
            A.store.Bool(ptr + 496 + 4 + 128, "UseProactiveKeyCaching" in x["WiMAX"]["EAP"] ? true : false);
            A.store.Bool(ptr + 496 + 4 + 125, x["WiMAX"]["EAP"]["UseProactiveKeyCaching"] ? true : false);
            A.store.Bool(ptr + 496 + 4 + 129, "UseSystemCAs" in x["WiMAX"]["EAP"] ? true : false);
            A.store.Bool(ptr + 496 + 4 + 126, x["WiMAX"]["EAP"]["UseSystemCAs"] ? true : false);
          }
        }
      }
    },
    "load_NetworkConfigProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 127)) {
        x["Cellular"] = {};
        if (A.load.Bool(ptr + 0 + 121)) {
          x["Cellular"]["AutoConnect"] = A.load.Bool(ptr + 0 + 0);
        } else {
          delete x["Cellular"]["AutoConnect"];
        }
        x["Cellular"]["ActivationType"] = A.load.Ref(ptr + 0 + 4, undefined);
        x["Cellular"]["ActivationState"] = A.load.Enum(ptr + 0 + 8, [
          "Activated",
          "Activating",
          "NotActivated",
          "PartiallyActivated",
        ]);
        if (A.load.Bool(ptr + 0 + 122)) {
          x["Cellular"]["AllowRoaming"] = A.load.Bool(ptr + 0 + 12);
        } else {
          delete x["Cellular"]["AllowRoaming"];
        }
        x["Cellular"]["Family"] = A.load.Ref(ptr + 0 + 16, undefined);
        x["Cellular"]["FirmwareRevision"] = A.load.Ref(ptr + 0 + 20, undefined);
        x["Cellular"]["FoundNetworks"] = A.load.Ref(ptr + 0 + 24, undefined);
        x["Cellular"]["HardwareRevision"] = A.load.Ref(ptr + 0 + 28, undefined);
        if (A.load.Bool(ptr + 0 + 32 + 12)) {
          x["Cellular"]["HomeProvider"] = {};
          x["Cellular"]["HomeProvider"]["Name"] = A.load.Ref(ptr + 0 + 32 + 0, undefined);
          x["Cellular"]["HomeProvider"]["Code"] = A.load.Ref(ptr + 0 + 32 + 4, undefined);
          x["Cellular"]["HomeProvider"]["Country"] = A.load.Ref(ptr + 0 + 32 + 8, undefined);
        } else {
          delete x["Cellular"]["HomeProvider"];
        }
        x["Cellular"]["Manufacturer"] = A.load.Ref(ptr + 0 + 48, undefined);
        x["Cellular"]["ModelID"] = A.load.Ref(ptr + 0 + 52, undefined);
        x["Cellular"]["NetworkTechnology"] = A.load.Ref(ptr + 0 + 56, undefined);
        if (A.load.Bool(ptr + 0 + 60 + 12)) {
          x["Cellular"]["PaymentPortal"] = {};
          x["Cellular"]["PaymentPortal"]["Method"] = A.load.Ref(ptr + 0 + 60 + 0, undefined);
          x["Cellular"]["PaymentPortal"]["PostData"] = A.load.Ref(ptr + 0 + 60 + 4, undefined);
          x["Cellular"]["PaymentPortal"]["Url"] = A.load.Ref(ptr + 0 + 60 + 8, undefined);
        } else {
          delete x["Cellular"]["PaymentPortal"];
        }
        x["Cellular"]["RoamingState"] = A.load.Ref(ptr + 0 + 76, undefined);
        if (A.load.Bool(ptr + 0 + 123)) {
          x["Cellular"]["Scanning"] = A.load.Bool(ptr + 0 + 80);
        } else {
          delete x["Cellular"]["Scanning"];
        }
        if (A.load.Bool(ptr + 0 + 84 + 12)) {
          x["Cellular"]["ServingOperator"] = {};
          x["Cellular"]["ServingOperator"]["Name"] = A.load.Ref(ptr + 0 + 84 + 0, undefined);
          x["Cellular"]["ServingOperator"]["Code"] = A.load.Ref(ptr + 0 + 84 + 4, undefined);
          x["Cellular"]["ServingOperator"]["Country"] = A.load.Ref(ptr + 0 + 84 + 8, undefined);
        } else {
          delete x["Cellular"]["ServingOperator"];
        }
        if (A.load.Bool(ptr + 0 + 100 + 14)) {
          x["Cellular"]["SIMLockStatus"] = {};
          x["Cellular"]["SIMLockStatus"]["LockType"] = A.load.Ref(ptr + 0 + 100 + 0, undefined);
          if (A.load.Bool(ptr + 0 + 100 + 12)) {
            x["Cellular"]["SIMLockStatus"]["LockEnabled"] = A.load.Bool(ptr + 0 + 100 + 4);
          } else {
            delete x["Cellular"]["SIMLockStatus"]["LockEnabled"];
          }
          if (A.load.Bool(ptr + 0 + 100 + 13)) {
            x["Cellular"]["SIMLockStatus"]["RetriesLeft"] = A.load.Int32(ptr + 0 + 100 + 8);
          } else {
            delete x["Cellular"]["SIMLockStatus"]["RetriesLeft"];
          }
        } else {
          delete x["Cellular"]["SIMLockStatus"];
        }
        if (A.load.Bool(ptr + 0 + 124)) {
          x["Cellular"]["SIMPresent"] = A.load.Bool(ptr + 0 + 115);
        } else {
          delete x["Cellular"]["SIMPresent"];
        }
        if (A.load.Bool(ptr + 0 + 125)) {
          x["Cellular"]["SignalStrength"] = A.load.Int32(ptr + 0 + 116);
        } else {
          delete x["Cellular"]["SignalStrength"];
        }
        if (A.load.Bool(ptr + 0 + 126)) {
          x["Cellular"]["SupportNetworkScan"] = A.load.Bool(ptr + 0 + 120);
        } else {
          delete x["Cellular"]["SupportNetworkScan"];
        }
      } else {
        delete x["Cellular"];
      }
      if (A.load.Bool(ptr + 128 + 140)) {
        x["Ethernet"] = {};
        if (A.load.Bool(ptr + 128 + 139)) {
          x["Ethernet"]["AutoConnect"] = A.load.Bool(ptr + 128 + 0);
        } else {
          delete x["Ethernet"]["AutoConnect"];
        }
        x["Ethernet"]["Authentication"] = A.load.Ref(ptr + 128 + 4, undefined);
        if (A.load.Bool(ptr + 128 + 8 + 130)) {
          x["Ethernet"]["EAP"] = {};
          x["Ethernet"]["EAP"]["AnonymousIdentity"] = A.load.Ref(ptr + 128 + 8 + 0, undefined);
          if (A.load.Bool(ptr + 128 + 8 + 4 + 45)) {
            x["Ethernet"]["EAP"]["ClientCertPattern"] = {};
            x["Ethernet"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(ptr + 128 + 8 + 4 + 0, undefined);
            if (A.load.Bool(ptr + 128 + 8 + 4 + 4 + 16)) {
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                ptr + 128 + 8 + 4 + 4 + 0,
                undefined
              );
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                ptr + 128 + 8 + 4 + 4 + 4,
                undefined
              );
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                ptr + 128 + 8 + 4 + 4 + 8,
                undefined
              );
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                ptr + 128 + 8 + 4 + 4 + 12,
                undefined
              );
            } else {
              delete x["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"];
            }
            x["Ethernet"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(ptr + 128 + 8 + 4 + 24, undefined);
            if (A.load.Bool(ptr + 128 + 8 + 4 + 28 + 16)) {
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"] = {};
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                ptr + 128 + 8 + 4 + 28 + 0,
                undefined
              );
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                ptr + 128 + 8 + 4 + 28 + 4,
                undefined
              );
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                ptr + 128 + 8 + 4 + 28 + 8,
                undefined
              );
              x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                ptr + 128 + 8 + 4 + 28 + 12,
                undefined
              );
            } else {
              delete x["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"];
            }
          } else {
            delete x["Ethernet"]["EAP"]["ClientCertPattern"];
          }
          x["Ethernet"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(ptr + 128 + 8 + 52, undefined);
          x["Ethernet"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(ptr + 128 + 8 + 56, undefined);
          x["Ethernet"]["EAP"]["ClientCertRef"] = A.load.Ref(ptr + 128 + 8 + 60, undefined);
          x["Ethernet"]["EAP"]["ClientCertType"] = A.load.Enum(ptr + 128 + 8 + 64, ["Ref", "Pattern"]);
          x["Ethernet"]["EAP"]["Identity"] = A.load.Ref(ptr + 128 + 8 + 68, undefined);
          x["Ethernet"]["EAP"]["Inner"] = A.load.Ref(ptr + 128 + 8 + 72, undefined);
          x["Ethernet"]["EAP"]["Outer"] = A.load.Ref(ptr + 128 + 8 + 76, undefined);
          x["Ethernet"]["EAP"]["Password"] = A.load.Ref(ptr + 128 + 8 + 80, undefined);
          if (A.load.Bool(ptr + 128 + 8 + 127)) {
            x["Ethernet"]["EAP"]["SaveCredentials"] = A.load.Bool(ptr + 128 + 8 + 84);
          } else {
            delete x["Ethernet"]["EAP"]["SaveCredentials"];
          }
          x["Ethernet"]["EAP"]["ServerCAPEMs"] = A.load.Ref(ptr + 128 + 8 + 88, undefined);
          x["Ethernet"]["EAP"]["ServerCARefs"] = A.load.Ref(ptr + 128 + 8 + 92, undefined);
          if (A.load.Bool(ptr + 128 + 8 + 96 + 28)) {
            x["Ethernet"]["EAP"]["SubjectMatch"] = {};
            x["Ethernet"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(ptr + 128 + 8 + 96 + 0, undefined);
            x["Ethernet"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(ptr + 128 + 8 + 96 + 4, undefined);
            x["Ethernet"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(ptr + 128 + 8 + 96 + 8, undefined);
            x["Ethernet"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(ptr + 128 + 8 + 96 + 12, undefined);
            x["Ethernet"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(ptr + 128 + 8 + 96 + 16, undefined);
            x["Ethernet"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(ptr + 128 + 8 + 96 + 20, undefined);
            if (A.load.Bool(ptr + 128 + 8 + 96 + 26)) {
              x["Ethernet"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(ptr + 128 + 8 + 96 + 24);
            } else {
              delete x["Ethernet"]["EAP"]["SubjectMatch"]["UserEditable"];
            }
            if (A.load.Bool(ptr + 128 + 8 + 96 + 27)) {
              x["Ethernet"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(ptr + 128 + 8 + 96 + 25);
            } else {
              delete x["Ethernet"]["EAP"]["SubjectMatch"]["DeviceEditable"];
            }
          } else {
            delete x["Ethernet"]["EAP"]["SubjectMatch"];
          }
          if (A.load.Bool(ptr + 128 + 8 + 128)) {
            x["Ethernet"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(ptr + 128 + 8 + 125);
          } else {
            delete x["Ethernet"]["EAP"]["UseProactiveKeyCaching"];
          }
          if (A.load.Bool(ptr + 128 + 8 + 129)) {
            x["Ethernet"]["EAP"]["UseSystemCAs"] = A.load.Bool(ptr + 128 + 8 + 126);
          } else {
            delete x["Ethernet"]["EAP"]["UseSystemCAs"];
          }
        } else {
          delete x["Ethernet"]["EAP"];
        }
      } else {
        delete x["Ethernet"];
      }
      x["GUID"] = A.load.Ref(ptr + 272, undefined);
      x["IPAddressConfigType"] = A.load.Enum(ptr + 276, ["DHCP", "Static"]);
      x["Name"] = A.load.Ref(ptr + 280, undefined);
      x["NameServersConfigType"] = A.load.Enum(ptr + 284, ["DHCP", "Static"]);
      if (A.load.Bool(ptr + 633)) {
        x["Priority"] = A.load.Int32(ptr + 288);
      } else {
        delete x["Priority"];
      }
      x["Type"] = A.load.Enum(ptr + 292, ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"]);
      if (A.load.Bool(ptr + 296 + 13)) {
        x["VPN"] = {};
        if (A.load.Bool(ptr + 296 + 12)) {
          x["VPN"]["AutoConnect"] = A.load.Bool(ptr + 296 + 0);
        } else {
          delete x["VPN"]["AutoConnect"];
        }
        x["VPN"]["Host"] = A.load.Ref(ptr + 296 + 4, undefined);
        x["VPN"]["Type"] = A.load.Ref(ptr + 296 + 8, undefined);
      } else {
        delete x["VPN"];
      }
      if (A.load.Bool(ptr + 312 + 182)) {
        x["WiFi"] = {};
        if (A.load.Bool(ptr + 312 + 176)) {
          x["WiFi"]["AllowGatewayARPPolling"] = A.load.Bool(ptr + 312 + 0);
        } else {
          delete x["WiFi"]["AllowGatewayARPPolling"];
        }
        if (A.load.Bool(ptr + 312 + 177)) {
          x["WiFi"]["AutoConnect"] = A.load.Bool(ptr + 312 + 1);
        } else {
          delete x["WiFi"]["AutoConnect"];
        }
        x["WiFi"]["BSSID"] = A.load.Ref(ptr + 312 + 4, undefined);
        if (A.load.Bool(ptr + 312 + 8 + 130)) {
          x["WiFi"]["EAP"] = {};
          x["WiFi"]["EAP"]["AnonymousIdentity"] = A.load.Ref(ptr + 312 + 8 + 0, undefined);
          if (A.load.Bool(ptr + 312 + 8 + 4 + 45)) {
            x["WiFi"]["EAP"]["ClientCertPattern"] = {};
            x["WiFi"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(ptr + 312 + 8 + 4 + 0, undefined);
            if (A.load.Bool(ptr + 312 + 8 + 4 + 4 + 16)) {
              x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
              x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                ptr + 312 + 8 + 4 + 4 + 0,
                undefined
              );
              x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                ptr + 312 + 8 + 4 + 4 + 4,
                undefined
              );
              x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                ptr + 312 + 8 + 4 + 4 + 8,
                undefined
              );
              x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                ptr + 312 + 8 + 4 + 4 + 12,
                undefined
              );
            } else {
              delete x["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"];
            }
            x["WiFi"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(ptr + 312 + 8 + 4 + 24, undefined);
            if (A.load.Bool(ptr + 312 + 8 + 4 + 28 + 16)) {
              x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"] = {};
              x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                ptr + 312 + 8 + 4 + 28 + 0,
                undefined
              );
              x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                ptr + 312 + 8 + 4 + 28 + 4,
                undefined
              );
              x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                ptr + 312 + 8 + 4 + 28 + 8,
                undefined
              );
              x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                ptr + 312 + 8 + 4 + 28 + 12,
                undefined
              );
            } else {
              delete x["WiFi"]["EAP"]["ClientCertPattern"]["Subject"];
            }
          } else {
            delete x["WiFi"]["EAP"]["ClientCertPattern"];
          }
          x["WiFi"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(ptr + 312 + 8 + 52, undefined);
          x["WiFi"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(ptr + 312 + 8 + 56, undefined);
          x["WiFi"]["EAP"]["ClientCertRef"] = A.load.Ref(ptr + 312 + 8 + 60, undefined);
          x["WiFi"]["EAP"]["ClientCertType"] = A.load.Enum(ptr + 312 + 8 + 64, ["Ref", "Pattern"]);
          x["WiFi"]["EAP"]["Identity"] = A.load.Ref(ptr + 312 + 8 + 68, undefined);
          x["WiFi"]["EAP"]["Inner"] = A.load.Ref(ptr + 312 + 8 + 72, undefined);
          x["WiFi"]["EAP"]["Outer"] = A.load.Ref(ptr + 312 + 8 + 76, undefined);
          x["WiFi"]["EAP"]["Password"] = A.load.Ref(ptr + 312 + 8 + 80, undefined);
          if (A.load.Bool(ptr + 312 + 8 + 127)) {
            x["WiFi"]["EAP"]["SaveCredentials"] = A.load.Bool(ptr + 312 + 8 + 84);
          } else {
            delete x["WiFi"]["EAP"]["SaveCredentials"];
          }
          x["WiFi"]["EAP"]["ServerCAPEMs"] = A.load.Ref(ptr + 312 + 8 + 88, undefined);
          x["WiFi"]["EAP"]["ServerCARefs"] = A.load.Ref(ptr + 312 + 8 + 92, undefined);
          if (A.load.Bool(ptr + 312 + 8 + 96 + 28)) {
            x["WiFi"]["EAP"]["SubjectMatch"] = {};
            x["WiFi"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(ptr + 312 + 8 + 96 + 0, undefined);
            x["WiFi"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(ptr + 312 + 8 + 96 + 4, undefined);
            x["WiFi"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(ptr + 312 + 8 + 96 + 8, undefined);
            x["WiFi"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(ptr + 312 + 8 + 96 + 12, undefined);
            x["WiFi"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(ptr + 312 + 8 + 96 + 16, undefined);
            x["WiFi"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(ptr + 312 + 8 + 96 + 20, undefined);
            if (A.load.Bool(ptr + 312 + 8 + 96 + 26)) {
              x["WiFi"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(ptr + 312 + 8 + 96 + 24);
            } else {
              delete x["WiFi"]["EAP"]["SubjectMatch"]["UserEditable"];
            }
            if (A.load.Bool(ptr + 312 + 8 + 96 + 27)) {
              x["WiFi"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(ptr + 312 + 8 + 96 + 25);
            } else {
              delete x["WiFi"]["EAP"]["SubjectMatch"]["DeviceEditable"];
            }
          } else {
            delete x["WiFi"]["EAP"]["SubjectMatch"];
          }
          if (A.load.Bool(ptr + 312 + 8 + 128)) {
            x["WiFi"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(ptr + 312 + 8 + 125);
          } else {
            delete x["WiFi"]["EAP"]["UseProactiveKeyCaching"];
          }
          if (A.load.Bool(ptr + 312 + 8 + 129)) {
            x["WiFi"]["EAP"]["UseSystemCAs"] = A.load.Bool(ptr + 312 + 8 + 126);
          } else {
            delete x["WiFi"]["EAP"]["UseSystemCAs"];
          }
        } else {
          delete x["WiFi"]["EAP"];
        }
        if (A.load.Bool(ptr + 312 + 178)) {
          x["WiFi"]["Frequency"] = A.load.Int32(ptr + 312 + 140);
        } else {
          delete x["WiFi"]["Frequency"];
        }
        x["WiFi"]["FrequencyList"] = A.load.Ref(ptr + 312 + 144, undefined);
        x["WiFi"]["HexSSID"] = A.load.Ref(ptr + 312 + 148, undefined);
        if (A.load.Bool(ptr + 312 + 179)) {
          x["WiFi"]["HiddenSSID"] = A.load.Bool(ptr + 312 + 152);
        } else {
          delete x["WiFi"]["HiddenSSID"];
        }
        x["WiFi"]["Passphrase"] = A.load.Ref(ptr + 312 + 156, undefined);
        if (A.load.Bool(ptr + 312 + 180)) {
          x["WiFi"]["RoamThreshold"] = A.load.Int32(ptr + 312 + 160);
        } else {
          delete x["WiFi"]["RoamThreshold"];
        }
        x["WiFi"]["SSID"] = A.load.Ref(ptr + 312 + 164, undefined);
        x["WiFi"]["Security"] = A.load.Ref(ptr + 312 + 168, undefined);
        if (A.load.Bool(ptr + 312 + 181)) {
          x["WiFi"]["SignalStrength"] = A.load.Int32(ptr + 312 + 172);
        } else {
          delete x["WiFi"]["SignalStrength"];
        }
      } else {
        delete x["WiFi"];
      }
      if (A.load.Bool(ptr + 496 + 136)) {
        x["WiMAX"] = {};
        if (A.load.Bool(ptr + 496 + 135)) {
          x["WiMAX"]["AutoConnect"] = A.load.Bool(ptr + 496 + 0);
        } else {
          delete x["WiMAX"]["AutoConnect"];
        }
        if (A.load.Bool(ptr + 496 + 4 + 130)) {
          x["WiMAX"]["EAP"] = {};
          x["WiMAX"]["EAP"]["AnonymousIdentity"] = A.load.Ref(ptr + 496 + 4 + 0, undefined);
          if (A.load.Bool(ptr + 496 + 4 + 4 + 45)) {
            x["WiMAX"]["EAP"]["ClientCertPattern"] = {};
            x["WiMAX"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(ptr + 496 + 4 + 4 + 0, undefined);
            if (A.load.Bool(ptr + 496 + 4 + 4 + 4 + 16)) {
              x["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
              x["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                ptr + 496 + 4 + 4 + 4 + 0,
                undefined
              );
              x["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                ptr + 496 + 4 + 4 + 4 + 4,
                undefined
              );
              x["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                ptr + 496 + 4 + 4 + 4 + 8,
                undefined
              );
              x["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                ptr + 496 + 4 + 4 + 4 + 12,
                undefined
              );
            } else {
              delete x["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"];
            }
            x["WiMAX"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(ptr + 496 + 4 + 4 + 24, undefined);
            if (A.load.Bool(ptr + 496 + 4 + 4 + 28 + 16)) {
              x["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"] = {};
              x["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                ptr + 496 + 4 + 4 + 28 + 0,
                undefined
              );
              x["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                ptr + 496 + 4 + 4 + 28 + 4,
                undefined
              );
              x["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                ptr + 496 + 4 + 4 + 28 + 8,
                undefined
              );
              x["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                ptr + 496 + 4 + 4 + 28 + 12,
                undefined
              );
            } else {
              delete x["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"];
            }
          } else {
            delete x["WiMAX"]["EAP"]["ClientCertPattern"];
          }
          x["WiMAX"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(ptr + 496 + 4 + 52, undefined);
          x["WiMAX"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(ptr + 496 + 4 + 56, undefined);
          x["WiMAX"]["EAP"]["ClientCertRef"] = A.load.Ref(ptr + 496 + 4 + 60, undefined);
          x["WiMAX"]["EAP"]["ClientCertType"] = A.load.Enum(ptr + 496 + 4 + 64, ["Ref", "Pattern"]);
          x["WiMAX"]["EAP"]["Identity"] = A.load.Ref(ptr + 496 + 4 + 68, undefined);
          x["WiMAX"]["EAP"]["Inner"] = A.load.Ref(ptr + 496 + 4 + 72, undefined);
          x["WiMAX"]["EAP"]["Outer"] = A.load.Ref(ptr + 496 + 4 + 76, undefined);
          x["WiMAX"]["EAP"]["Password"] = A.load.Ref(ptr + 496 + 4 + 80, undefined);
          if (A.load.Bool(ptr + 496 + 4 + 127)) {
            x["WiMAX"]["EAP"]["SaveCredentials"] = A.load.Bool(ptr + 496 + 4 + 84);
          } else {
            delete x["WiMAX"]["EAP"]["SaveCredentials"];
          }
          x["WiMAX"]["EAP"]["ServerCAPEMs"] = A.load.Ref(ptr + 496 + 4 + 88, undefined);
          x["WiMAX"]["EAP"]["ServerCARefs"] = A.load.Ref(ptr + 496 + 4 + 92, undefined);
          if (A.load.Bool(ptr + 496 + 4 + 96 + 28)) {
            x["WiMAX"]["EAP"]["SubjectMatch"] = {};
            x["WiMAX"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(ptr + 496 + 4 + 96 + 0, undefined);
            x["WiMAX"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(ptr + 496 + 4 + 96 + 4, undefined);
            x["WiMAX"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(ptr + 496 + 4 + 96 + 8, undefined);
            x["WiMAX"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(ptr + 496 + 4 + 96 + 12, undefined);
            x["WiMAX"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(ptr + 496 + 4 + 96 + 16, undefined);
            x["WiMAX"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(ptr + 496 + 4 + 96 + 20, undefined);
            if (A.load.Bool(ptr + 496 + 4 + 96 + 26)) {
              x["WiMAX"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(ptr + 496 + 4 + 96 + 24);
            } else {
              delete x["WiMAX"]["EAP"]["SubjectMatch"]["UserEditable"];
            }
            if (A.load.Bool(ptr + 496 + 4 + 96 + 27)) {
              x["WiMAX"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(ptr + 496 + 4 + 96 + 25);
            } else {
              delete x["WiMAX"]["EAP"]["SubjectMatch"]["DeviceEditable"];
            }
          } else {
            delete x["WiMAX"]["EAP"]["SubjectMatch"];
          }
          if (A.load.Bool(ptr + 496 + 4 + 128)) {
            x["WiMAX"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(ptr + 496 + 4 + 125);
          } else {
            delete x["WiMAX"]["EAP"]["UseProactiveKeyCaching"];
          }
          if (A.load.Bool(ptr + 496 + 4 + 129)) {
            x["WiMAX"]["EAP"]["UseSystemCAs"] = A.load.Bool(ptr + 496 + 4 + 126);
          } else {
            delete x["WiMAX"]["EAP"]["UseSystemCAs"];
          }
        } else {
          delete x["WiMAX"]["EAP"];
        }
      } else {
        delete x["WiMAX"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_NetworkFilter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 15, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 14, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 15, true);
        A.store.Enum(
          ptr + 0,
          ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"].indexOf(x["networkType"] as string)
        );
        A.store.Bool(ptr + 12, "visible" in x ? true : false);
        A.store.Bool(ptr + 4, x["visible"] ? true : false);
        A.store.Bool(ptr + 13, "configured" in x ? true : false);
        A.store.Bool(ptr + 5, x["configured"] ? true : false);
        A.store.Bool(ptr + 14, "limit" in x ? true : false);
        A.store.Int32(ptr + 8, x["limit"] === undefined ? 0 : (x["limit"] as number));
      }
    },
    "load_NetworkFilter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["networkType"] = A.load.Enum(ptr + 0, ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"]);
      if (A.load.Bool(ptr + 12)) {
        x["visible"] = A.load.Bool(ptr + 4);
      } else {
        delete x["visible"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["configured"] = A.load.Bool(ptr + 5);
      } else {
        delete x["configured"];
      }
      if (A.load.Bool(ptr + 14)) {
        x["limit"] = A.load.Int32(ptr + 8);
      } else {
        delete x["limit"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ThirdPartyVPNProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["ExtensionID"]);
        A.store.Ref(ptr + 4, x["ProviderName"]);
      }
    },
    "load_ThirdPartyVPNProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["ExtensionID"] = A.load.Ref(ptr + 0, undefined);
      x["ProviderName"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_CreateNetwork": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc && "createNetwork" in WEBEXT?.networking?.onc) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CreateNetwork": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.createNetwork);
    },
    "call_CreateNetwork": (
      retPtr: Pointer,
      shared: heap.Ref<boolean>,
      properties: Pointer,
      callback: heap.Ref<object>
    ): void => {
      const properties_ffi = {};

      if (A.load.Bool(properties + 0 + 127)) {
        properties_ffi["Cellular"] = {};
        if (A.load.Bool(properties + 0 + 121)) {
          properties_ffi["Cellular"]["AutoConnect"] = A.load.Bool(properties + 0 + 0);
        }
        properties_ffi["Cellular"]["ActivationType"] = A.load.Ref(properties + 0 + 4, undefined);
        properties_ffi["Cellular"]["ActivationState"] = A.load.Enum(properties + 0 + 8, [
          "Activated",
          "Activating",
          "NotActivated",
          "PartiallyActivated",
        ]);
        if (A.load.Bool(properties + 0 + 122)) {
          properties_ffi["Cellular"]["AllowRoaming"] = A.load.Bool(properties + 0 + 12);
        }
        properties_ffi["Cellular"]["Family"] = A.load.Ref(properties + 0 + 16, undefined);
        properties_ffi["Cellular"]["FirmwareRevision"] = A.load.Ref(properties + 0 + 20, undefined);
        properties_ffi["Cellular"]["FoundNetworks"] = A.load.Ref(properties + 0 + 24, undefined);
        properties_ffi["Cellular"]["HardwareRevision"] = A.load.Ref(properties + 0 + 28, undefined);
        if (A.load.Bool(properties + 0 + 32 + 12)) {
          properties_ffi["Cellular"]["HomeProvider"] = {};
          properties_ffi["Cellular"]["HomeProvider"]["Name"] = A.load.Ref(properties + 0 + 32 + 0, undefined);
          properties_ffi["Cellular"]["HomeProvider"]["Code"] = A.load.Ref(properties + 0 + 32 + 4, undefined);
          properties_ffi["Cellular"]["HomeProvider"]["Country"] = A.load.Ref(properties + 0 + 32 + 8, undefined);
        }
        properties_ffi["Cellular"]["Manufacturer"] = A.load.Ref(properties + 0 + 48, undefined);
        properties_ffi["Cellular"]["ModelID"] = A.load.Ref(properties + 0 + 52, undefined);
        properties_ffi["Cellular"]["NetworkTechnology"] = A.load.Ref(properties + 0 + 56, undefined);
        if (A.load.Bool(properties + 0 + 60 + 12)) {
          properties_ffi["Cellular"]["PaymentPortal"] = {};
          properties_ffi["Cellular"]["PaymentPortal"]["Method"] = A.load.Ref(properties + 0 + 60 + 0, undefined);
          properties_ffi["Cellular"]["PaymentPortal"]["PostData"] = A.load.Ref(properties + 0 + 60 + 4, undefined);
          properties_ffi["Cellular"]["PaymentPortal"]["Url"] = A.load.Ref(properties + 0 + 60 + 8, undefined);
        }
        properties_ffi["Cellular"]["RoamingState"] = A.load.Ref(properties + 0 + 76, undefined);
        if (A.load.Bool(properties + 0 + 123)) {
          properties_ffi["Cellular"]["Scanning"] = A.load.Bool(properties + 0 + 80);
        }
        if (A.load.Bool(properties + 0 + 84 + 12)) {
          properties_ffi["Cellular"]["ServingOperator"] = {};
          properties_ffi["Cellular"]["ServingOperator"]["Name"] = A.load.Ref(properties + 0 + 84 + 0, undefined);
          properties_ffi["Cellular"]["ServingOperator"]["Code"] = A.load.Ref(properties + 0 + 84 + 4, undefined);
          properties_ffi["Cellular"]["ServingOperator"]["Country"] = A.load.Ref(properties + 0 + 84 + 8, undefined);
        }
        if (A.load.Bool(properties + 0 + 100 + 14)) {
          properties_ffi["Cellular"]["SIMLockStatus"] = {};
          properties_ffi["Cellular"]["SIMLockStatus"]["LockType"] = A.load.Ref(properties + 0 + 100 + 0, undefined);
          if (A.load.Bool(properties + 0 + 100 + 12)) {
            properties_ffi["Cellular"]["SIMLockStatus"]["LockEnabled"] = A.load.Bool(properties + 0 + 100 + 4);
          }
          if (A.load.Bool(properties + 0 + 100 + 13)) {
            properties_ffi["Cellular"]["SIMLockStatus"]["RetriesLeft"] = A.load.Int32(properties + 0 + 100 + 8);
          }
        }
        if (A.load.Bool(properties + 0 + 124)) {
          properties_ffi["Cellular"]["SIMPresent"] = A.load.Bool(properties + 0 + 115);
        }
        if (A.load.Bool(properties + 0 + 125)) {
          properties_ffi["Cellular"]["SignalStrength"] = A.load.Int32(properties + 0 + 116);
        }
        if (A.load.Bool(properties + 0 + 126)) {
          properties_ffi["Cellular"]["SupportNetworkScan"] = A.load.Bool(properties + 0 + 120);
        }
      }
      if (A.load.Bool(properties + 128 + 140)) {
        properties_ffi["Ethernet"] = {};
        if (A.load.Bool(properties + 128 + 139)) {
          properties_ffi["Ethernet"]["AutoConnect"] = A.load.Bool(properties + 128 + 0);
        }
        properties_ffi["Ethernet"]["Authentication"] = A.load.Ref(properties + 128 + 4, undefined);
        if (A.load.Bool(properties + 128 + 8 + 130)) {
          properties_ffi["Ethernet"]["EAP"] = {};
          properties_ffi["Ethernet"]["EAP"]["AnonymousIdentity"] = A.load.Ref(properties + 128 + 8 + 0, undefined);
          if (A.load.Bool(properties + 128 + 8 + 4 + 45)) {
            properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"] = {};
            properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(
              properties + 128 + 8 + 4 + 0,
              undefined
            );
            if (A.load.Bool(properties + 128 + 8 + 4 + 4 + 16)) {
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                properties + 128 + 8 + 4 + 4 + 0,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                properties + 128 + 8 + 4 + 4 + 4,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                properties + 128 + 8 + 4 + 4 + 8,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                properties + 128 + 8 + 4 + 4 + 12,
                undefined
              );
            }
            properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(
              properties + 128 + 8 + 4 + 24,
              undefined
            );
            if (A.load.Bool(properties + 128 + 8 + 4 + 28 + 16)) {
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"] = {};
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                properties + 128 + 8 + 4 + 28 + 0,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                properties + 128 + 8 + 4 + 28 + 4,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                properties + 128 + 8 + 4 + 28 + 8,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                properties + 128 + 8 + 4 + 28 + 12,
                undefined
              );
            }
          }
          properties_ffi["Ethernet"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(properties + 128 + 8 + 52, undefined);
          properties_ffi["Ethernet"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(
            properties + 128 + 8 + 56,
            undefined
          );
          properties_ffi["Ethernet"]["EAP"]["ClientCertRef"] = A.load.Ref(properties + 128 + 8 + 60, undefined);
          properties_ffi["Ethernet"]["EAP"]["ClientCertType"] = A.load.Enum(properties + 128 + 8 + 64, [
            "Ref",
            "Pattern",
          ]);
          properties_ffi["Ethernet"]["EAP"]["Identity"] = A.load.Ref(properties + 128 + 8 + 68, undefined);
          properties_ffi["Ethernet"]["EAP"]["Inner"] = A.load.Ref(properties + 128 + 8 + 72, undefined);
          properties_ffi["Ethernet"]["EAP"]["Outer"] = A.load.Ref(properties + 128 + 8 + 76, undefined);
          properties_ffi["Ethernet"]["EAP"]["Password"] = A.load.Ref(properties + 128 + 8 + 80, undefined);
          if (A.load.Bool(properties + 128 + 8 + 127)) {
            properties_ffi["Ethernet"]["EAP"]["SaveCredentials"] = A.load.Bool(properties + 128 + 8 + 84);
          }
          properties_ffi["Ethernet"]["EAP"]["ServerCAPEMs"] = A.load.Ref(properties + 128 + 8 + 88, undefined);
          properties_ffi["Ethernet"]["EAP"]["ServerCARefs"] = A.load.Ref(properties + 128 + 8 + 92, undefined);
          if (A.load.Bool(properties + 128 + 8 + 96 + 28)) {
            properties_ffi["Ethernet"]["EAP"]["SubjectMatch"] = {};
            properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(
              properties + 128 + 8 + 96 + 0,
              undefined
            );
            properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(
              properties + 128 + 8 + 96 + 4,
              undefined
            );
            properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(
              properties + 128 + 8 + 96 + 8,
              undefined
            );
            properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(
              properties + 128 + 8 + 96 + 12,
              undefined
            );
            properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(
              properties + 128 + 8 + 96 + 16,
              undefined
            );
            properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(
              properties + 128 + 8 + 96 + 20,
              undefined
            );
            if (A.load.Bool(properties + 128 + 8 + 96 + 26)) {
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(
                properties + 128 + 8 + 96 + 24
              );
            }
            if (A.load.Bool(properties + 128 + 8 + 96 + 27)) {
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(
                properties + 128 + 8 + 96 + 25
              );
            }
          }
          if (A.load.Bool(properties + 128 + 8 + 128)) {
            properties_ffi["Ethernet"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(properties + 128 + 8 + 125);
          }
          if (A.load.Bool(properties + 128 + 8 + 129)) {
            properties_ffi["Ethernet"]["EAP"]["UseSystemCAs"] = A.load.Bool(properties + 128 + 8 + 126);
          }
        }
      }
      properties_ffi["GUID"] = A.load.Ref(properties + 272, undefined);
      properties_ffi["IPAddressConfigType"] = A.load.Enum(properties + 276, ["DHCP", "Static"]);
      properties_ffi["Name"] = A.load.Ref(properties + 280, undefined);
      properties_ffi["NameServersConfigType"] = A.load.Enum(properties + 284, ["DHCP", "Static"]);
      if (A.load.Bool(properties + 633)) {
        properties_ffi["Priority"] = A.load.Int32(properties + 288);
      }
      properties_ffi["Type"] = A.load.Enum(properties + 292, [
        "All",
        "Cellular",
        "Ethernet",
        "Tether",
        "VPN",
        "Wireless",
        "WiFi",
      ]);
      if (A.load.Bool(properties + 296 + 13)) {
        properties_ffi["VPN"] = {};
        if (A.load.Bool(properties + 296 + 12)) {
          properties_ffi["VPN"]["AutoConnect"] = A.load.Bool(properties + 296 + 0);
        }
        properties_ffi["VPN"]["Host"] = A.load.Ref(properties + 296 + 4, undefined);
        properties_ffi["VPN"]["Type"] = A.load.Ref(properties + 296 + 8, undefined);
      }
      if (A.load.Bool(properties + 312 + 182)) {
        properties_ffi["WiFi"] = {};
        if (A.load.Bool(properties + 312 + 176)) {
          properties_ffi["WiFi"]["AllowGatewayARPPolling"] = A.load.Bool(properties + 312 + 0);
        }
        if (A.load.Bool(properties + 312 + 177)) {
          properties_ffi["WiFi"]["AutoConnect"] = A.load.Bool(properties + 312 + 1);
        }
        properties_ffi["WiFi"]["BSSID"] = A.load.Ref(properties + 312 + 4, undefined);
        if (A.load.Bool(properties + 312 + 8 + 130)) {
          properties_ffi["WiFi"]["EAP"] = {};
          properties_ffi["WiFi"]["EAP"]["AnonymousIdentity"] = A.load.Ref(properties + 312 + 8 + 0, undefined);
          if (A.load.Bool(properties + 312 + 8 + 4 + 45)) {
            properties_ffi["WiFi"]["EAP"]["ClientCertPattern"] = {};
            properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(
              properties + 312 + 8 + 4 + 0,
              undefined
            );
            if (A.load.Bool(properties + 312 + 8 + 4 + 4 + 16)) {
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                properties + 312 + 8 + 4 + 4 + 0,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                properties + 312 + 8 + 4 + 4 + 4,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                properties + 312 + 8 + 4 + 4 + 8,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                properties + 312 + 8 + 4 + 4 + 12,
                undefined
              );
            }
            properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(
              properties + 312 + 8 + 4 + 24,
              undefined
            );
            if (A.load.Bool(properties + 312 + 8 + 4 + 28 + 16)) {
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"] = {};
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                properties + 312 + 8 + 4 + 28 + 0,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                properties + 312 + 8 + 4 + 28 + 4,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                properties + 312 + 8 + 4 + 28 + 8,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                properties + 312 + 8 + 4 + 28 + 12,
                undefined
              );
            }
          }
          properties_ffi["WiFi"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(properties + 312 + 8 + 52, undefined);
          properties_ffi["WiFi"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(
            properties + 312 + 8 + 56,
            undefined
          );
          properties_ffi["WiFi"]["EAP"]["ClientCertRef"] = A.load.Ref(properties + 312 + 8 + 60, undefined);
          properties_ffi["WiFi"]["EAP"]["ClientCertType"] = A.load.Enum(properties + 312 + 8 + 64, ["Ref", "Pattern"]);
          properties_ffi["WiFi"]["EAP"]["Identity"] = A.load.Ref(properties + 312 + 8 + 68, undefined);
          properties_ffi["WiFi"]["EAP"]["Inner"] = A.load.Ref(properties + 312 + 8 + 72, undefined);
          properties_ffi["WiFi"]["EAP"]["Outer"] = A.load.Ref(properties + 312 + 8 + 76, undefined);
          properties_ffi["WiFi"]["EAP"]["Password"] = A.load.Ref(properties + 312 + 8 + 80, undefined);
          if (A.load.Bool(properties + 312 + 8 + 127)) {
            properties_ffi["WiFi"]["EAP"]["SaveCredentials"] = A.load.Bool(properties + 312 + 8 + 84);
          }
          properties_ffi["WiFi"]["EAP"]["ServerCAPEMs"] = A.load.Ref(properties + 312 + 8 + 88, undefined);
          properties_ffi["WiFi"]["EAP"]["ServerCARefs"] = A.load.Ref(properties + 312 + 8 + 92, undefined);
          if (A.load.Bool(properties + 312 + 8 + 96 + 28)) {
            properties_ffi["WiFi"]["EAP"]["SubjectMatch"] = {};
            properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(
              properties + 312 + 8 + 96 + 0,
              undefined
            );
            properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(
              properties + 312 + 8 + 96 + 4,
              undefined
            );
            properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(
              properties + 312 + 8 + 96 + 8,
              undefined
            );
            properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(
              properties + 312 + 8 + 96 + 12,
              undefined
            );
            properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(
              properties + 312 + 8 + 96 + 16,
              undefined
            );
            properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(
              properties + 312 + 8 + 96 + 20,
              undefined
            );
            if (A.load.Bool(properties + 312 + 8 + 96 + 26)) {
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(
                properties + 312 + 8 + 96 + 24
              );
            }
            if (A.load.Bool(properties + 312 + 8 + 96 + 27)) {
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(
                properties + 312 + 8 + 96 + 25
              );
            }
          }
          if (A.load.Bool(properties + 312 + 8 + 128)) {
            properties_ffi["WiFi"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(properties + 312 + 8 + 125);
          }
          if (A.load.Bool(properties + 312 + 8 + 129)) {
            properties_ffi["WiFi"]["EAP"]["UseSystemCAs"] = A.load.Bool(properties + 312 + 8 + 126);
          }
        }
        if (A.load.Bool(properties + 312 + 178)) {
          properties_ffi["WiFi"]["Frequency"] = A.load.Int32(properties + 312 + 140);
        }
        properties_ffi["WiFi"]["FrequencyList"] = A.load.Ref(properties + 312 + 144, undefined);
        properties_ffi["WiFi"]["HexSSID"] = A.load.Ref(properties + 312 + 148, undefined);
        if (A.load.Bool(properties + 312 + 179)) {
          properties_ffi["WiFi"]["HiddenSSID"] = A.load.Bool(properties + 312 + 152);
        }
        properties_ffi["WiFi"]["Passphrase"] = A.load.Ref(properties + 312 + 156, undefined);
        if (A.load.Bool(properties + 312 + 180)) {
          properties_ffi["WiFi"]["RoamThreshold"] = A.load.Int32(properties + 312 + 160);
        }
        properties_ffi["WiFi"]["SSID"] = A.load.Ref(properties + 312 + 164, undefined);
        properties_ffi["WiFi"]["Security"] = A.load.Ref(properties + 312 + 168, undefined);
        if (A.load.Bool(properties + 312 + 181)) {
          properties_ffi["WiFi"]["SignalStrength"] = A.load.Int32(properties + 312 + 172);
        }
      }
      if (A.load.Bool(properties + 496 + 136)) {
        properties_ffi["WiMAX"] = {};
        if (A.load.Bool(properties + 496 + 135)) {
          properties_ffi["WiMAX"]["AutoConnect"] = A.load.Bool(properties + 496 + 0);
        }
        if (A.load.Bool(properties + 496 + 4 + 130)) {
          properties_ffi["WiMAX"]["EAP"] = {};
          properties_ffi["WiMAX"]["EAP"]["AnonymousIdentity"] = A.load.Ref(properties + 496 + 4 + 0, undefined);
          if (A.load.Bool(properties + 496 + 4 + 4 + 45)) {
            properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"] = {};
            properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(
              properties + 496 + 4 + 4 + 0,
              undefined
            );
            if (A.load.Bool(properties + 496 + 4 + 4 + 4 + 16)) {
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                properties + 496 + 4 + 4 + 4 + 0,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                properties + 496 + 4 + 4 + 4 + 4,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                properties + 496 + 4 + 4 + 4 + 8,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                properties + 496 + 4 + 4 + 4 + 12,
                undefined
              );
            }
            properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(
              properties + 496 + 4 + 4 + 24,
              undefined
            );
            if (A.load.Bool(properties + 496 + 4 + 4 + 28 + 16)) {
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"] = {};
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                properties + 496 + 4 + 4 + 28 + 0,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                properties + 496 + 4 + 4 + 28 + 4,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                properties + 496 + 4 + 4 + 28 + 8,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                properties + 496 + 4 + 4 + 28 + 12,
                undefined
              );
            }
          }
          properties_ffi["WiMAX"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(properties + 496 + 4 + 52, undefined);
          properties_ffi["WiMAX"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(
            properties + 496 + 4 + 56,
            undefined
          );
          properties_ffi["WiMAX"]["EAP"]["ClientCertRef"] = A.load.Ref(properties + 496 + 4 + 60, undefined);
          properties_ffi["WiMAX"]["EAP"]["ClientCertType"] = A.load.Enum(properties + 496 + 4 + 64, ["Ref", "Pattern"]);
          properties_ffi["WiMAX"]["EAP"]["Identity"] = A.load.Ref(properties + 496 + 4 + 68, undefined);
          properties_ffi["WiMAX"]["EAP"]["Inner"] = A.load.Ref(properties + 496 + 4 + 72, undefined);
          properties_ffi["WiMAX"]["EAP"]["Outer"] = A.load.Ref(properties + 496 + 4 + 76, undefined);
          properties_ffi["WiMAX"]["EAP"]["Password"] = A.load.Ref(properties + 496 + 4 + 80, undefined);
          if (A.load.Bool(properties + 496 + 4 + 127)) {
            properties_ffi["WiMAX"]["EAP"]["SaveCredentials"] = A.load.Bool(properties + 496 + 4 + 84);
          }
          properties_ffi["WiMAX"]["EAP"]["ServerCAPEMs"] = A.load.Ref(properties + 496 + 4 + 88, undefined);
          properties_ffi["WiMAX"]["EAP"]["ServerCARefs"] = A.load.Ref(properties + 496 + 4 + 92, undefined);
          if (A.load.Bool(properties + 496 + 4 + 96 + 28)) {
            properties_ffi["WiMAX"]["EAP"]["SubjectMatch"] = {};
            properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(
              properties + 496 + 4 + 96 + 0,
              undefined
            );
            properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(
              properties + 496 + 4 + 96 + 4,
              undefined
            );
            properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(
              properties + 496 + 4 + 96 + 8,
              undefined
            );
            properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(
              properties + 496 + 4 + 96 + 12,
              undefined
            );
            properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(
              properties + 496 + 4 + 96 + 16,
              undefined
            );
            properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(
              properties + 496 + 4 + 96 + 20,
              undefined
            );
            if (A.load.Bool(properties + 496 + 4 + 96 + 26)) {
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(
                properties + 496 + 4 + 96 + 24
              );
            }
            if (A.load.Bool(properties + 496 + 4 + 96 + 27)) {
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(
                properties + 496 + 4 + 96 + 25
              );
            }
          }
          if (A.load.Bool(properties + 496 + 4 + 128)) {
            properties_ffi["WiMAX"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(properties + 496 + 4 + 125);
          }
          if (A.load.Bool(properties + 496 + 4 + 129)) {
            properties_ffi["WiMAX"]["EAP"]["UseSystemCAs"] = A.load.Bool(properties + 496 + 4 + 126);
          }
        }
      }

      const _ret = WEBEXT.networking.onc.createNetwork(shared === A.H.TRUE, properties_ffi, A.H.get<object>(callback));
    },
    "try_CreateNetwork": (
      retPtr: Pointer,
      errPtr: Pointer,
      shared: heap.Ref<boolean>,
      properties: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const properties_ffi = {};

        if (A.load.Bool(properties + 0 + 127)) {
          properties_ffi["Cellular"] = {};
          if (A.load.Bool(properties + 0 + 121)) {
            properties_ffi["Cellular"]["AutoConnect"] = A.load.Bool(properties + 0 + 0);
          }
          properties_ffi["Cellular"]["ActivationType"] = A.load.Ref(properties + 0 + 4, undefined);
          properties_ffi["Cellular"]["ActivationState"] = A.load.Enum(properties + 0 + 8, [
            "Activated",
            "Activating",
            "NotActivated",
            "PartiallyActivated",
          ]);
          if (A.load.Bool(properties + 0 + 122)) {
            properties_ffi["Cellular"]["AllowRoaming"] = A.load.Bool(properties + 0 + 12);
          }
          properties_ffi["Cellular"]["Family"] = A.load.Ref(properties + 0 + 16, undefined);
          properties_ffi["Cellular"]["FirmwareRevision"] = A.load.Ref(properties + 0 + 20, undefined);
          properties_ffi["Cellular"]["FoundNetworks"] = A.load.Ref(properties + 0 + 24, undefined);
          properties_ffi["Cellular"]["HardwareRevision"] = A.load.Ref(properties + 0 + 28, undefined);
          if (A.load.Bool(properties + 0 + 32 + 12)) {
            properties_ffi["Cellular"]["HomeProvider"] = {};
            properties_ffi["Cellular"]["HomeProvider"]["Name"] = A.load.Ref(properties + 0 + 32 + 0, undefined);
            properties_ffi["Cellular"]["HomeProvider"]["Code"] = A.load.Ref(properties + 0 + 32 + 4, undefined);
            properties_ffi["Cellular"]["HomeProvider"]["Country"] = A.load.Ref(properties + 0 + 32 + 8, undefined);
          }
          properties_ffi["Cellular"]["Manufacturer"] = A.load.Ref(properties + 0 + 48, undefined);
          properties_ffi["Cellular"]["ModelID"] = A.load.Ref(properties + 0 + 52, undefined);
          properties_ffi["Cellular"]["NetworkTechnology"] = A.load.Ref(properties + 0 + 56, undefined);
          if (A.load.Bool(properties + 0 + 60 + 12)) {
            properties_ffi["Cellular"]["PaymentPortal"] = {};
            properties_ffi["Cellular"]["PaymentPortal"]["Method"] = A.load.Ref(properties + 0 + 60 + 0, undefined);
            properties_ffi["Cellular"]["PaymentPortal"]["PostData"] = A.load.Ref(properties + 0 + 60 + 4, undefined);
            properties_ffi["Cellular"]["PaymentPortal"]["Url"] = A.load.Ref(properties + 0 + 60 + 8, undefined);
          }
          properties_ffi["Cellular"]["RoamingState"] = A.load.Ref(properties + 0 + 76, undefined);
          if (A.load.Bool(properties + 0 + 123)) {
            properties_ffi["Cellular"]["Scanning"] = A.load.Bool(properties + 0 + 80);
          }
          if (A.load.Bool(properties + 0 + 84 + 12)) {
            properties_ffi["Cellular"]["ServingOperator"] = {};
            properties_ffi["Cellular"]["ServingOperator"]["Name"] = A.load.Ref(properties + 0 + 84 + 0, undefined);
            properties_ffi["Cellular"]["ServingOperator"]["Code"] = A.load.Ref(properties + 0 + 84 + 4, undefined);
            properties_ffi["Cellular"]["ServingOperator"]["Country"] = A.load.Ref(properties + 0 + 84 + 8, undefined);
          }
          if (A.load.Bool(properties + 0 + 100 + 14)) {
            properties_ffi["Cellular"]["SIMLockStatus"] = {};
            properties_ffi["Cellular"]["SIMLockStatus"]["LockType"] = A.load.Ref(properties + 0 + 100 + 0, undefined);
            if (A.load.Bool(properties + 0 + 100 + 12)) {
              properties_ffi["Cellular"]["SIMLockStatus"]["LockEnabled"] = A.load.Bool(properties + 0 + 100 + 4);
            }
            if (A.load.Bool(properties + 0 + 100 + 13)) {
              properties_ffi["Cellular"]["SIMLockStatus"]["RetriesLeft"] = A.load.Int32(properties + 0 + 100 + 8);
            }
          }
          if (A.load.Bool(properties + 0 + 124)) {
            properties_ffi["Cellular"]["SIMPresent"] = A.load.Bool(properties + 0 + 115);
          }
          if (A.load.Bool(properties + 0 + 125)) {
            properties_ffi["Cellular"]["SignalStrength"] = A.load.Int32(properties + 0 + 116);
          }
          if (A.load.Bool(properties + 0 + 126)) {
            properties_ffi["Cellular"]["SupportNetworkScan"] = A.load.Bool(properties + 0 + 120);
          }
        }
        if (A.load.Bool(properties + 128 + 140)) {
          properties_ffi["Ethernet"] = {};
          if (A.load.Bool(properties + 128 + 139)) {
            properties_ffi["Ethernet"]["AutoConnect"] = A.load.Bool(properties + 128 + 0);
          }
          properties_ffi["Ethernet"]["Authentication"] = A.load.Ref(properties + 128 + 4, undefined);
          if (A.load.Bool(properties + 128 + 8 + 130)) {
            properties_ffi["Ethernet"]["EAP"] = {};
            properties_ffi["Ethernet"]["EAP"]["AnonymousIdentity"] = A.load.Ref(properties + 128 + 8 + 0, undefined);
            if (A.load.Bool(properties + 128 + 8 + 4 + 45)) {
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"] = {};
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(
                properties + 128 + 8 + 4 + 0,
                undefined
              );
              if (A.load.Bool(properties + 128 + 8 + 4 + 4 + 16)) {
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 4 + 0,
                  undefined
                );
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 4 + 4,
                  undefined
                );
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 4 + 8,
                  undefined
                );
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 4 + 12,
                  undefined
                );
              }
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(
                properties + 128 + 8 + 4 + 24,
                undefined
              );
              if (A.load.Bool(properties + 128 + 8 + 4 + 28 + 16)) {
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"] = {};
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 28 + 0,
                  undefined
                );
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 28 + 4,
                  undefined
                );
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 28 + 8,
                  undefined
                );
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 28 + 12,
                  undefined
                );
              }
            }
            properties_ffi["Ethernet"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(properties + 128 + 8 + 52, undefined);
            properties_ffi["Ethernet"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(
              properties + 128 + 8 + 56,
              undefined
            );
            properties_ffi["Ethernet"]["EAP"]["ClientCertRef"] = A.load.Ref(properties + 128 + 8 + 60, undefined);
            properties_ffi["Ethernet"]["EAP"]["ClientCertType"] = A.load.Enum(properties + 128 + 8 + 64, [
              "Ref",
              "Pattern",
            ]);
            properties_ffi["Ethernet"]["EAP"]["Identity"] = A.load.Ref(properties + 128 + 8 + 68, undefined);
            properties_ffi["Ethernet"]["EAP"]["Inner"] = A.load.Ref(properties + 128 + 8 + 72, undefined);
            properties_ffi["Ethernet"]["EAP"]["Outer"] = A.load.Ref(properties + 128 + 8 + 76, undefined);
            properties_ffi["Ethernet"]["EAP"]["Password"] = A.load.Ref(properties + 128 + 8 + 80, undefined);
            if (A.load.Bool(properties + 128 + 8 + 127)) {
              properties_ffi["Ethernet"]["EAP"]["SaveCredentials"] = A.load.Bool(properties + 128 + 8 + 84);
            }
            properties_ffi["Ethernet"]["EAP"]["ServerCAPEMs"] = A.load.Ref(properties + 128 + 8 + 88, undefined);
            properties_ffi["Ethernet"]["EAP"]["ServerCARefs"] = A.load.Ref(properties + 128 + 8 + 92, undefined);
            if (A.load.Bool(properties + 128 + 8 + 96 + 28)) {
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"] = {};
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(
                properties + 128 + 8 + 96 + 0,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(
                properties + 128 + 8 + 96 + 4,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(
                properties + 128 + 8 + 96 + 8,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(
                properties + 128 + 8 + 96 + 12,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(
                properties + 128 + 8 + 96 + 16,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(
                properties + 128 + 8 + 96 + 20,
                undefined
              );
              if (A.load.Bool(properties + 128 + 8 + 96 + 26)) {
                properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(
                  properties + 128 + 8 + 96 + 24
                );
              }
              if (A.load.Bool(properties + 128 + 8 + 96 + 27)) {
                properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(
                  properties + 128 + 8 + 96 + 25
                );
              }
            }
            if (A.load.Bool(properties + 128 + 8 + 128)) {
              properties_ffi["Ethernet"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(properties + 128 + 8 + 125);
            }
            if (A.load.Bool(properties + 128 + 8 + 129)) {
              properties_ffi["Ethernet"]["EAP"]["UseSystemCAs"] = A.load.Bool(properties + 128 + 8 + 126);
            }
          }
        }
        properties_ffi["GUID"] = A.load.Ref(properties + 272, undefined);
        properties_ffi["IPAddressConfigType"] = A.load.Enum(properties + 276, ["DHCP", "Static"]);
        properties_ffi["Name"] = A.load.Ref(properties + 280, undefined);
        properties_ffi["NameServersConfigType"] = A.load.Enum(properties + 284, ["DHCP", "Static"]);
        if (A.load.Bool(properties + 633)) {
          properties_ffi["Priority"] = A.load.Int32(properties + 288);
        }
        properties_ffi["Type"] = A.load.Enum(properties + 292, [
          "All",
          "Cellular",
          "Ethernet",
          "Tether",
          "VPN",
          "Wireless",
          "WiFi",
        ]);
        if (A.load.Bool(properties + 296 + 13)) {
          properties_ffi["VPN"] = {};
          if (A.load.Bool(properties + 296 + 12)) {
            properties_ffi["VPN"]["AutoConnect"] = A.load.Bool(properties + 296 + 0);
          }
          properties_ffi["VPN"]["Host"] = A.load.Ref(properties + 296 + 4, undefined);
          properties_ffi["VPN"]["Type"] = A.load.Ref(properties + 296 + 8, undefined);
        }
        if (A.load.Bool(properties + 312 + 182)) {
          properties_ffi["WiFi"] = {};
          if (A.load.Bool(properties + 312 + 176)) {
            properties_ffi["WiFi"]["AllowGatewayARPPolling"] = A.load.Bool(properties + 312 + 0);
          }
          if (A.load.Bool(properties + 312 + 177)) {
            properties_ffi["WiFi"]["AutoConnect"] = A.load.Bool(properties + 312 + 1);
          }
          properties_ffi["WiFi"]["BSSID"] = A.load.Ref(properties + 312 + 4, undefined);
          if (A.load.Bool(properties + 312 + 8 + 130)) {
            properties_ffi["WiFi"]["EAP"] = {};
            properties_ffi["WiFi"]["EAP"]["AnonymousIdentity"] = A.load.Ref(properties + 312 + 8 + 0, undefined);
            if (A.load.Bool(properties + 312 + 8 + 4 + 45)) {
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"] = {};
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(
                properties + 312 + 8 + 4 + 0,
                undefined
              );
              if (A.load.Bool(properties + 312 + 8 + 4 + 4 + 16)) {
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 4 + 0,
                  undefined
                );
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 4 + 4,
                  undefined
                );
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 4 + 8,
                  undefined
                );
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 4 + 12,
                  undefined
                );
              }
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(
                properties + 312 + 8 + 4 + 24,
                undefined
              );
              if (A.load.Bool(properties + 312 + 8 + 4 + 28 + 16)) {
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"] = {};
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 28 + 0,
                  undefined
                );
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 28 + 4,
                  undefined
                );
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 28 + 8,
                  undefined
                );
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 28 + 12,
                  undefined
                );
              }
            }
            properties_ffi["WiFi"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(properties + 312 + 8 + 52, undefined);
            properties_ffi["WiFi"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(
              properties + 312 + 8 + 56,
              undefined
            );
            properties_ffi["WiFi"]["EAP"]["ClientCertRef"] = A.load.Ref(properties + 312 + 8 + 60, undefined);
            properties_ffi["WiFi"]["EAP"]["ClientCertType"] = A.load.Enum(properties + 312 + 8 + 64, [
              "Ref",
              "Pattern",
            ]);
            properties_ffi["WiFi"]["EAP"]["Identity"] = A.load.Ref(properties + 312 + 8 + 68, undefined);
            properties_ffi["WiFi"]["EAP"]["Inner"] = A.load.Ref(properties + 312 + 8 + 72, undefined);
            properties_ffi["WiFi"]["EAP"]["Outer"] = A.load.Ref(properties + 312 + 8 + 76, undefined);
            properties_ffi["WiFi"]["EAP"]["Password"] = A.load.Ref(properties + 312 + 8 + 80, undefined);
            if (A.load.Bool(properties + 312 + 8 + 127)) {
              properties_ffi["WiFi"]["EAP"]["SaveCredentials"] = A.load.Bool(properties + 312 + 8 + 84);
            }
            properties_ffi["WiFi"]["EAP"]["ServerCAPEMs"] = A.load.Ref(properties + 312 + 8 + 88, undefined);
            properties_ffi["WiFi"]["EAP"]["ServerCARefs"] = A.load.Ref(properties + 312 + 8 + 92, undefined);
            if (A.load.Bool(properties + 312 + 8 + 96 + 28)) {
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"] = {};
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(
                properties + 312 + 8 + 96 + 0,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(
                properties + 312 + 8 + 96 + 4,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(
                properties + 312 + 8 + 96 + 8,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(
                properties + 312 + 8 + 96 + 12,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(
                properties + 312 + 8 + 96 + 16,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(
                properties + 312 + 8 + 96 + 20,
                undefined
              );
              if (A.load.Bool(properties + 312 + 8 + 96 + 26)) {
                properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(
                  properties + 312 + 8 + 96 + 24
                );
              }
              if (A.load.Bool(properties + 312 + 8 + 96 + 27)) {
                properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(
                  properties + 312 + 8 + 96 + 25
                );
              }
            }
            if (A.load.Bool(properties + 312 + 8 + 128)) {
              properties_ffi["WiFi"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(properties + 312 + 8 + 125);
            }
            if (A.load.Bool(properties + 312 + 8 + 129)) {
              properties_ffi["WiFi"]["EAP"]["UseSystemCAs"] = A.load.Bool(properties + 312 + 8 + 126);
            }
          }
          if (A.load.Bool(properties + 312 + 178)) {
            properties_ffi["WiFi"]["Frequency"] = A.load.Int32(properties + 312 + 140);
          }
          properties_ffi["WiFi"]["FrequencyList"] = A.load.Ref(properties + 312 + 144, undefined);
          properties_ffi["WiFi"]["HexSSID"] = A.load.Ref(properties + 312 + 148, undefined);
          if (A.load.Bool(properties + 312 + 179)) {
            properties_ffi["WiFi"]["HiddenSSID"] = A.load.Bool(properties + 312 + 152);
          }
          properties_ffi["WiFi"]["Passphrase"] = A.load.Ref(properties + 312 + 156, undefined);
          if (A.load.Bool(properties + 312 + 180)) {
            properties_ffi["WiFi"]["RoamThreshold"] = A.load.Int32(properties + 312 + 160);
          }
          properties_ffi["WiFi"]["SSID"] = A.load.Ref(properties + 312 + 164, undefined);
          properties_ffi["WiFi"]["Security"] = A.load.Ref(properties + 312 + 168, undefined);
          if (A.load.Bool(properties + 312 + 181)) {
            properties_ffi["WiFi"]["SignalStrength"] = A.load.Int32(properties + 312 + 172);
          }
        }
        if (A.load.Bool(properties + 496 + 136)) {
          properties_ffi["WiMAX"] = {};
          if (A.load.Bool(properties + 496 + 135)) {
            properties_ffi["WiMAX"]["AutoConnect"] = A.load.Bool(properties + 496 + 0);
          }
          if (A.load.Bool(properties + 496 + 4 + 130)) {
            properties_ffi["WiMAX"]["EAP"] = {};
            properties_ffi["WiMAX"]["EAP"]["AnonymousIdentity"] = A.load.Ref(properties + 496 + 4 + 0, undefined);
            if (A.load.Bool(properties + 496 + 4 + 4 + 45)) {
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"] = {};
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(
                properties + 496 + 4 + 4 + 0,
                undefined
              );
              if (A.load.Bool(properties + 496 + 4 + 4 + 4 + 16)) {
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 4 + 0,
                  undefined
                );
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 4 + 4,
                  undefined
                );
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 4 + 8,
                  undefined
                );
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 4 + 12,
                  undefined
                );
              }
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(
                properties + 496 + 4 + 4 + 24,
                undefined
              );
              if (A.load.Bool(properties + 496 + 4 + 4 + 28 + 16)) {
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"] = {};
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 28 + 0,
                  undefined
                );
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 28 + 4,
                  undefined
                );
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 28 + 8,
                  undefined
                );
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 28 + 12,
                  undefined
                );
              }
            }
            properties_ffi["WiMAX"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(properties + 496 + 4 + 52, undefined);
            properties_ffi["WiMAX"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(
              properties + 496 + 4 + 56,
              undefined
            );
            properties_ffi["WiMAX"]["EAP"]["ClientCertRef"] = A.load.Ref(properties + 496 + 4 + 60, undefined);
            properties_ffi["WiMAX"]["EAP"]["ClientCertType"] = A.load.Enum(properties + 496 + 4 + 64, [
              "Ref",
              "Pattern",
            ]);
            properties_ffi["WiMAX"]["EAP"]["Identity"] = A.load.Ref(properties + 496 + 4 + 68, undefined);
            properties_ffi["WiMAX"]["EAP"]["Inner"] = A.load.Ref(properties + 496 + 4 + 72, undefined);
            properties_ffi["WiMAX"]["EAP"]["Outer"] = A.load.Ref(properties + 496 + 4 + 76, undefined);
            properties_ffi["WiMAX"]["EAP"]["Password"] = A.load.Ref(properties + 496 + 4 + 80, undefined);
            if (A.load.Bool(properties + 496 + 4 + 127)) {
              properties_ffi["WiMAX"]["EAP"]["SaveCredentials"] = A.load.Bool(properties + 496 + 4 + 84);
            }
            properties_ffi["WiMAX"]["EAP"]["ServerCAPEMs"] = A.load.Ref(properties + 496 + 4 + 88, undefined);
            properties_ffi["WiMAX"]["EAP"]["ServerCARefs"] = A.load.Ref(properties + 496 + 4 + 92, undefined);
            if (A.load.Bool(properties + 496 + 4 + 96 + 28)) {
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"] = {};
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(
                properties + 496 + 4 + 96 + 0,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(
                properties + 496 + 4 + 96 + 4,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(
                properties + 496 + 4 + 96 + 8,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(
                properties + 496 + 4 + 96 + 12,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(
                properties + 496 + 4 + 96 + 16,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(
                properties + 496 + 4 + 96 + 20,
                undefined
              );
              if (A.load.Bool(properties + 496 + 4 + 96 + 26)) {
                properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(
                  properties + 496 + 4 + 96 + 24
                );
              }
              if (A.load.Bool(properties + 496 + 4 + 96 + 27)) {
                properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(
                  properties + 496 + 4 + 96 + 25
                );
              }
            }
            if (A.load.Bool(properties + 496 + 4 + 128)) {
              properties_ffi["WiMAX"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(properties + 496 + 4 + 125);
            }
            if (A.load.Bool(properties + 496 + 4 + 129)) {
              properties_ffi["WiMAX"]["EAP"]["UseSystemCAs"] = A.load.Bool(properties + 496 + 4 + 126);
            }
          }
        }

        const _ret = WEBEXT.networking.onc.createNetwork(
          shared === A.H.TRUE,
          properties_ffi,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DisableNetworkType": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc && "disableNetworkType" in WEBEXT?.networking?.onc) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DisableNetworkType": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.disableNetworkType);
    },
    "call_DisableNetworkType": (retPtr: Pointer, networkType: number): void => {
      const _ret = WEBEXT.networking.onc.disableNetworkType(
        networkType > 0 && networkType <= 7
          ? ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"][networkType - 1]
          : undefined
      );
    },
    "try_DisableNetworkType": (retPtr: Pointer, errPtr: Pointer, networkType: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.disableNetworkType(
          networkType > 0 && networkType <= 7
            ? ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"][networkType - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_EnableNetworkType": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc && "enableNetworkType" in WEBEXT?.networking?.onc) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EnableNetworkType": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.enableNetworkType);
    },
    "call_EnableNetworkType": (retPtr: Pointer, networkType: number): void => {
      const _ret = WEBEXT.networking.onc.enableNetworkType(
        networkType > 0 && networkType <= 7
          ? ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"][networkType - 1]
          : undefined
      );
    },
    "try_EnableNetworkType": (retPtr: Pointer, errPtr: Pointer, networkType: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.enableNetworkType(
          networkType > 0 && networkType <= 7
            ? ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"][networkType - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ForgetNetwork": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc && "forgetNetwork" in WEBEXT?.networking?.onc) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ForgetNetwork": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.forgetNetwork);
    },
    "call_ForgetNetwork": (retPtr: Pointer, networkGuid: heap.Ref<object>, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.forgetNetwork(A.H.get<object>(networkGuid), A.H.get<object>(callback));
    },
    "try_ForgetNetwork": (
      retPtr: Pointer,
      errPtr: Pointer,
      networkGuid: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.forgetNetwork(A.H.get<object>(networkGuid), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCaptivePortalStatus": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc && "getCaptivePortalStatus" in WEBEXT?.networking?.onc) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCaptivePortalStatus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.getCaptivePortalStatus);
    },
    "call_GetCaptivePortalStatus": (
      retPtr: Pointer,
      networkGuid: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.networking.onc.getCaptivePortalStatus(
        A.H.get<object>(networkGuid),
        A.H.get<object>(callback)
      );
    },
    "try_GetCaptivePortalStatus": (
      retPtr: Pointer,
      errPtr: Pointer,
      networkGuid: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.getCaptivePortalStatus(
          A.H.get<object>(networkGuid),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDeviceStates": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc && "getDeviceStates" in WEBEXT?.networking?.onc) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDeviceStates": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.getDeviceStates);
    },
    "call_GetDeviceStates": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.getDeviceStates(A.H.get<object>(callback));
    },
    "try_GetDeviceStates": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.getDeviceStates(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetGlobalPolicy": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc && "getGlobalPolicy" in WEBEXT?.networking?.onc) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetGlobalPolicy": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.getGlobalPolicy);
    },
    "call_GetGlobalPolicy": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.getGlobalPolicy(A.H.get<object>(callback));
    },
    "try_GetGlobalPolicy": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.getGlobalPolicy(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetManagedProperties": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc && "getManagedProperties" in WEBEXT?.networking?.onc) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetManagedProperties": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.getManagedProperties);
    },
    "call_GetManagedProperties": (retPtr: Pointer, networkGuid: heap.Ref<object>, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.getManagedProperties(A.H.get<object>(networkGuid), A.H.get<object>(callback));
    },
    "try_GetManagedProperties": (
      retPtr: Pointer,
      errPtr: Pointer,
      networkGuid: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.getManagedProperties(
          A.H.get<object>(networkGuid),
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetNetworks": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc && "getNetworks" in WEBEXT?.networking?.onc) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetNetworks": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.getNetworks);
    },
    "call_GetNetworks": (retPtr: Pointer, filter: Pointer, callback: heap.Ref<object>): void => {
      const filter_ffi = {};

      filter_ffi["networkType"] = A.load.Enum(filter + 0, [
        "All",
        "Cellular",
        "Ethernet",
        "Tether",
        "VPN",
        "Wireless",
        "WiFi",
      ]);
      if (A.load.Bool(filter + 12)) {
        filter_ffi["visible"] = A.load.Bool(filter + 4);
      }
      if (A.load.Bool(filter + 13)) {
        filter_ffi["configured"] = A.load.Bool(filter + 5);
      }
      if (A.load.Bool(filter + 14)) {
        filter_ffi["limit"] = A.load.Int32(filter + 8);
      }

      const _ret = WEBEXT.networking.onc.getNetworks(filter_ffi, A.H.get<object>(callback));
    },
    "try_GetNetworks": (
      retPtr: Pointer,
      errPtr: Pointer,
      filter: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const filter_ffi = {};

        filter_ffi["networkType"] = A.load.Enum(filter + 0, [
          "All",
          "Cellular",
          "Ethernet",
          "Tether",
          "VPN",
          "Wireless",
          "WiFi",
        ]);
        if (A.load.Bool(filter + 12)) {
          filter_ffi["visible"] = A.load.Bool(filter + 4);
        }
        if (A.load.Bool(filter + 13)) {
          filter_ffi["configured"] = A.load.Bool(filter + 5);
        }
        if (A.load.Bool(filter + 14)) {
          filter_ffi["limit"] = A.load.Int32(filter + 8);
        }

        const _ret = WEBEXT.networking.onc.getNetworks(filter_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetProperties": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc && "getProperties" in WEBEXT?.networking?.onc) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetProperties": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.getProperties);
    },
    "call_GetProperties": (retPtr: Pointer, networkGuid: heap.Ref<object>, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.getProperties(A.H.get<object>(networkGuid), A.H.get<object>(callback));
    },
    "try_GetProperties": (
      retPtr: Pointer,
      errPtr: Pointer,
      networkGuid: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.getProperties(A.H.get<object>(networkGuid), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetState": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc && "getState" in WEBEXT?.networking?.onc) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.getState);
    },
    "call_GetState": (retPtr: Pointer, networkGuid: heap.Ref<object>, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.getState(A.H.get<object>(networkGuid), A.H.get<object>(callback));
    },
    "try_GetState": (
      retPtr: Pointer,
      errPtr: Pointer,
      networkGuid: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.getState(A.H.get<object>(networkGuid), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeviceStateListChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.networking?.onc?.onDeviceStateListChanged &&
        "addListener" in WEBEXT?.networking?.onc?.onDeviceStateListChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeviceStateListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.onDeviceStateListChanged.addListener);
    },
    "call_OnDeviceStateListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.onDeviceStateListChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnDeviceStateListChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.onDeviceStateListChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeviceStateListChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.networking?.onc?.onDeviceStateListChanged &&
        "removeListener" in WEBEXT?.networking?.onc?.onDeviceStateListChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeviceStateListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.onDeviceStateListChanged.removeListener);
    },
    "call_OffDeviceStateListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.onDeviceStateListChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeviceStateListChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.onDeviceStateListChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeviceStateListChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.networking?.onc?.onDeviceStateListChanged &&
        "hasListener" in WEBEXT?.networking?.onc?.onDeviceStateListChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeviceStateListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.onDeviceStateListChanged.hasListener);
    },
    "call_HasOnDeviceStateListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.onDeviceStateListChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeviceStateListChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.onDeviceStateListChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnNetworkListChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.networking?.onc?.onNetworkListChanged &&
        "addListener" in WEBEXT?.networking?.onc?.onNetworkListChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnNetworkListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.onNetworkListChanged.addListener);
    },
    "call_OnNetworkListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.onNetworkListChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnNetworkListChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.onNetworkListChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffNetworkListChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.networking?.onc?.onNetworkListChanged &&
        "removeListener" in WEBEXT?.networking?.onc?.onNetworkListChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffNetworkListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.onNetworkListChanged.removeListener);
    },
    "call_OffNetworkListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.onNetworkListChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffNetworkListChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.onNetworkListChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnNetworkListChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.networking?.onc?.onNetworkListChanged &&
        "hasListener" in WEBEXT?.networking?.onc?.onNetworkListChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnNetworkListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.onNetworkListChanged.hasListener);
    },
    "call_HasOnNetworkListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.onNetworkListChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnNetworkListChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.onNetworkListChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnNetworksChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc?.onNetworksChanged && "addListener" in WEBEXT?.networking?.onc?.onNetworksChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnNetworksChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.onNetworksChanged.addListener);
    },
    "call_OnNetworksChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.onNetworksChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnNetworksChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.onNetworksChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffNetworksChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.networking?.onc?.onNetworksChanged &&
        "removeListener" in WEBEXT?.networking?.onc?.onNetworksChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffNetworksChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.onNetworksChanged.removeListener);
    },
    "call_OffNetworksChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.onNetworksChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffNetworksChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.onNetworksChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnNetworksChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc?.onNetworksChanged && "hasListener" in WEBEXT?.networking?.onc?.onNetworksChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnNetworksChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.onNetworksChanged.hasListener);
    },
    "call_HasOnNetworksChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.onNetworksChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnNetworksChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.onNetworksChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPortalDetectionCompleted": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.networking?.onc?.onPortalDetectionCompleted &&
        "addListener" in WEBEXT?.networking?.onc?.onPortalDetectionCompleted
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPortalDetectionCompleted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.onPortalDetectionCompleted.addListener);
    },
    "call_OnPortalDetectionCompleted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.onPortalDetectionCompleted.addListener(A.H.get<object>(callback));
    },
    "try_OnPortalDetectionCompleted": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.onPortalDetectionCompleted.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPortalDetectionCompleted": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.networking?.onc?.onPortalDetectionCompleted &&
        "removeListener" in WEBEXT?.networking?.onc?.onPortalDetectionCompleted
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPortalDetectionCompleted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.onPortalDetectionCompleted.removeListener);
    },
    "call_OffPortalDetectionCompleted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.onPortalDetectionCompleted.removeListener(A.H.get<object>(callback));
    },
    "try_OffPortalDetectionCompleted": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.onPortalDetectionCompleted.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPortalDetectionCompleted": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.networking?.onc?.onPortalDetectionCompleted &&
        "hasListener" in WEBEXT?.networking?.onc?.onPortalDetectionCompleted
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPortalDetectionCompleted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.onPortalDetectionCompleted.hasListener);
    },
    "call_HasOnPortalDetectionCompleted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.onPortalDetectionCompleted.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPortalDetectionCompleted": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.onPortalDetectionCompleted.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RequestNetworkScan": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc && "requestNetworkScan" in WEBEXT?.networking?.onc) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RequestNetworkScan": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.requestNetworkScan);
    },
    "call_RequestNetworkScan": (retPtr: Pointer, networkType: number): void => {
      const _ret = WEBEXT.networking.onc.requestNetworkScan(
        networkType > 0 && networkType <= 7
          ? ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"][networkType - 1]
          : undefined
      );
    },
    "try_RequestNetworkScan": (retPtr: Pointer, errPtr: Pointer, networkType: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.requestNetworkScan(
          networkType > 0 && networkType <= 7
            ? ["All", "Cellular", "Ethernet", "Tether", "VPN", "Wireless", "WiFi"][networkType - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetProperties": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc && "setProperties" in WEBEXT?.networking?.onc) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetProperties": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.setProperties);
    },
    "call_SetProperties": (
      retPtr: Pointer,
      networkGuid: heap.Ref<object>,
      properties: Pointer,
      callback: heap.Ref<object>
    ): void => {
      const properties_ffi = {};

      if (A.load.Bool(properties + 0 + 127)) {
        properties_ffi["Cellular"] = {};
        if (A.load.Bool(properties + 0 + 121)) {
          properties_ffi["Cellular"]["AutoConnect"] = A.load.Bool(properties + 0 + 0);
        }
        properties_ffi["Cellular"]["ActivationType"] = A.load.Ref(properties + 0 + 4, undefined);
        properties_ffi["Cellular"]["ActivationState"] = A.load.Enum(properties + 0 + 8, [
          "Activated",
          "Activating",
          "NotActivated",
          "PartiallyActivated",
        ]);
        if (A.load.Bool(properties + 0 + 122)) {
          properties_ffi["Cellular"]["AllowRoaming"] = A.load.Bool(properties + 0 + 12);
        }
        properties_ffi["Cellular"]["Family"] = A.load.Ref(properties + 0 + 16, undefined);
        properties_ffi["Cellular"]["FirmwareRevision"] = A.load.Ref(properties + 0 + 20, undefined);
        properties_ffi["Cellular"]["FoundNetworks"] = A.load.Ref(properties + 0 + 24, undefined);
        properties_ffi["Cellular"]["HardwareRevision"] = A.load.Ref(properties + 0 + 28, undefined);
        if (A.load.Bool(properties + 0 + 32 + 12)) {
          properties_ffi["Cellular"]["HomeProvider"] = {};
          properties_ffi["Cellular"]["HomeProvider"]["Name"] = A.load.Ref(properties + 0 + 32 + 0, undefined);
          properties_ffi["Cellular"]["HomeProvider"]["Code"] = A.load.Ref(properties + 0 + 32 + 4, undefined);
          properties_ffi["Cellular"]["HomeProvider"]["Country"] = A.load.Ref(properties + 0 + 32 + 8, undefined);
        }
        properties_ffi["Cellular"]["Manufacturer"] = A.load.Ref(properties + 0 + 48, undefined);
        properties_ffi["Cellular"]["ModelID"] = A.load.Ref(properties + 0 + 52, undefined);
        properties_ffi["Cellular"]["NetworkTechnology"] = A.load.Ref(properties + 0 + 56, undefined);
        if (A.load.Bool(properties + 0 + 60 + 12)) {
          properties_ffi["Cellular"]["PaymentPortal"] = {};
          properties_ffi["Cellular"]["PaymentPortal"]["Method"] = A.load.Ref(properties + 0 + 60 + 0, undefined);
          properties_ffi["Cellular"]["PaymentPortal"]["PostData"] = A.load.Ref(properties + 0 + 60 + 4, undefined);
          properties_ffi["Cellular"]["PaymentPortal"]["Url"] = A.load.Ref(properties + 0 + 60 + 8, undefined);
        }
        properties_ffi["Cellular"]["RoamingState"] = A.load.Ref(properties + 0 + 76, undefined);
        if (A.load.Bool(properties + 0 + 123)) {
          properties_ffi["Cellular"]["Scanning"] = A.load.Bool(properties + 0 + 80);
        }
        if (A.load.Bool(properties + 0 + 84 + 12)) {
          properties_ffi["Cellular"]["ServingOperator"] = {};
          properties_ffi["Cellular"]["ServingOperator"]["Name"] = A.load.Ref(properties + 0 + 84 + 0, undefined);
          properties_ffi["Cellular"]["ServingOperator"]["Code"] = A.load.Ref(properties + 0 + 84 + 4, undefined);
          properties_ffi["Cellular"]["ServingOperator"]["Country"] = A.load.Ref(properties + 0 + 84 + 8, undefined);
        }
        if (A.load.Bool(properties + 0 + 100 + 14)) {
          properties_ffi["Cellular"]["SIMLockStatus"] = {};
          properties_ffi["Cellular"]["SIMLockStatus"]["LockType"] = A.load.Ref(properties + 0 + 100 + 0, undefined);
          if (A.load.Bool(properties + 0 + 100 + 12)) {
            properties_ffi["Cellular"]["SIMLockStatus"]["LockEnabled"] = A.load.Bool(properties + 0 + 100 + 4);
          }
          if (A.load.Bool(properties + 0 + 100 + 13)) {
            properties_ffi["Cellular"]["SIMLockStatus"]["RetriesLeft"] = A.load.Int32(properties + 0 + 100 + 8);
          }
        }
        if (A.load.Bool(properties + 0 + 124)) {
          properties_ffi["Cellular"]["SIMPresent"] = A.load.Bool(properties + 0 + 115);
        }
        if (A.load.Bool(properties + 0 + 125)) {
          properties_ffi["Cellular"]["SignalStrength"] = A.load.Int32(properties + 0 + 116);
        }
        if (A.load.Bool(properties + 0 + 126)) {
          properties_ffi["Cellular"]["SupportNetworkScan"] = A.load.Bool(properties + 0 + 120);
        }
      }
      if (A.load.Bool(properties + 128 + 140)) {
        properties_ffi["Ethernet"] = {};
        if (A.load.Bool(properties + 128 + 139)) {
          properties_ffi["Ethernet"]["AutoConnect"] = A.load.Bool(properties + 128 + 0);
        }
        properties_ffi["Ethernet"]["Authentication"] = A.load.Ref(properties + 128 + 4, undefined);
        if (A.load.Bool(properties + 128 + 8 + 130)) {
          properties_ffi["Ethernet"]["EAP"] = {};
          properties_ffi["Ethernet"]["EAP"]["AnonymousIdentity"] = A.load.Ref(properties + 128 + 8 + 0, undefined);
          if (A.load.Bool(properties + 128 + 8 + 4 + 45)) {
            properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"] = {};
            properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(
              properties + 128 + 8 + 4 + 0,
              undefined
            );
            if (A.load.Bool(properties + 128 + 8 + 4 + 4 + 16)) {
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                properties + 128 + 8 + 4 + 4 + 0,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                properties + 128 + 8 + 4 + 4 + 4,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                properties + 128 + 8 + 4 + 4 + 8,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                properties + 128 + 8 + 4 + 4 + 12,
                undefined
              );
            }
            properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(
              properties + 128 + 8 + 4 + 24,
              undefined
            );
            if (A.load.Bool(properties + 128 + 8 + 4 + 28 + 16)) {
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"] = {};
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                properties + 128 + 8 + 4 + 28 + 0,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                properties + 128 + 8 + 4 + 28 + 4,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                properties + 128 + 8 + 4 + 28 + 8,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                properties + 128 + 8 + 4 + 28 + 12,
                undefined
              );
            }
          }
          properties_ffi["Ethernet"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(properties + 128 + 8 + 52, undefined);
          properties_ffi["Ethernet"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(
            properties + 128 + 8 + 56,
            undefined
          );
          properties_ffi["Ethernet"]["EAP"]["ClientCertRef"] = A.load.Ref(properties + 128 + 8 + 60, undefined);
          properties_ffi["Ethernet"]["EAP"]["ClientCertType"] = A.load.Enum(properties + 128 + 8 + 64, [
            "Ref",
            "Pattern",
          ]);
          properties_ffi["Ethernet"]["EAP"]["Identity"] = A.load.Ref(properties + 128 + 8 + 68, undefined);
          properties_ffi["Ethernet"]["EAP"]["Inner"] = A.load.Ref(properties + 128 + 8 + 72, undefined);
          properties_ffi["Ethernet"]["EAP"]["Outer"] = A.load.Ref(properties + 128 + 8 + 76, undefined);
          properties_ffi["Ethernet"]["EAP"]["Password"] = A.load.Ref(properties + 128 + 8 + 80, undefined);
          if (A.load.Bool(properties + 128 + 8 + 127)) {
            properties_ffi["Ethernet"]["EAP"]["SaveCredentials"] = A.load.Bool(properties + 128 + 8 + 84);
          }
          properties_ffi["Ethernet"]["EAP"]["ServerCAPEMs"] = A.load.Ref(properties + 128 + 8 + 88, undefined);
          properties_ffi["Ethernet"]["EAP"]["ServerCARefs"] = A.load.Ref(properties + 128 + 8 + 92, undefined);
          if (A.load.Bool(properties + 128 + 8 + 96 + 28)) {
            properties_ffi["Ethernet"]["EAP"]["SubjectMatch"] = {};
            properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(
              properties + 128 + 8 + 96 + 0,
              undefined
            );
            properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(
              properties + 128 + 8 + 96 + 4,
              undefined
            );
            properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(
              properties + 128 + 8 + 96 + 8,
              undefined
            );
            properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(
              properties + 128 + 8 + 96 + 12,
              undefined
            );
            properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(
              properties + 128 + 8 + 96 + 16,
              undefined
            );
            properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(
              properties + 128 + 8 + 96 + 20,
              undefined
            );
            if (A.load.Bool(properties + 128 + 8 + 96 + 26)) {
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(
                properties + 128 + 8 + 96 + 24
              );
            }
            if (A.load.Bool(properties + 128 + 8 + 96 + 27)) {
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(
                properties + 128 + 8 + 96 + 25
              );
            }
          }
          if (A.load.Bool(properties + 128 + 8 + 128)) {
            properties_ffi["Ethernet"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(properties + 128 + 8 + 125);
          }
          if (A.load.Bool(properties + 128 + 8 + 129)) {
            properties_ffi["Ethernet"]["EAP"]["UseSystemCAs"] = A.load.Bool(properties + 128 + 8 + 126);
          }
        }
      }
      properties_ffi["GUID"] = A.load.Ref(properties + 272, undefined);
      properties_ffi["IPAddressConfigType"] = A.load.Enum(properties + 276, ["DHCP", "Static"]);
      properties_ffi["Name"] = A.load.Ref(properties + 280, undefined);
      properties_ffi["NameServersConfigType"] = A.load.Enum(properties + 284, ["DHCP", "Static"]);
      if (A.load.Bool(properties + 633)) {
        properties_ffi["Priority"] = A.load.Int32(properties + 288);
      }
      properties_ffi["Type"] = A.load.Enum(properties + 292, [
        "All",
        "Cellular",
        "Ethernet",
        "Tether",
        "VPN",
        "Wireless",
        "WiFi",
      ]);
      if (A.load.Bool(properties + 296 + 13)) {
        properties_ffi["VPN"] = {};
        if (A.load.Bool(properties + 296 + 12)) {
          properties_ffi["VPN"]["AutoConnect"] = A.load.Bool(properties + 296 + 0);
        }
        properties_ffi["VPN"]["Host"] = A.load.Ref(properties + 296 + 4, undefined);
        properties_ffi["VPN"]["Type"] = A.load.Ref(properties + 296 + 8, undefined);
      }
      if (A.load.Bool(properties + 312 + 182)) {
        properties_ffi["WiFi"] = {};
        if (A.load.Bool(properties + 312 + 176)) {
          properties_ffi["WiFi"]["AllowGatewayARPPolling"] = A.load.Bool(properties + 312 + 0);
        }
        if (A.load.Bool(properties + 312 + 177)) {
          properties_ffi["WiFi"]["AutoConnect"] = A.load.Bool(properties + 312 + 1);
        }
        properties_ffi["WiFi"]["BSSID"] = A.load.Ref(properties + 312 + 4, undefined);
        if (A.load.Bool(properties + 312 + 8 + 130)) {
          properties_ffi["WiFi"]["EAP"] = {};
          properties_ffi["WiFi"]["EAP"]["AnonymousIdentity"] = A.load.Ref(properties + 312 + 8 + 0, undefined);
          if (A.load.Bool(properties + 312 + 8 + 4 + 45)) {
            properties_ffi["WiFi"]["EAP"]["ClientCertPattern"] = {};
            properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(
              properties + 312 + 8 + 4 + 0,
              undefined
            );
            if (A.load.Bool(properties + 312 + 8 + 4 + 4 + 16)) {
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                properties + 312 + 8 + 4 + 4 + 0,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                properties + 312 + 8 + 4 + 4 + 4,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                properties + 312 + 8 + 4 + 4 + 8,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                properties + 312 + 8 + 4 + 4 + 12,
                undefined
              );
            }
            properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(
              properties + 312 + 8 + 4 + 24,
              undefined
            );
            if (A.load.Bool(properties + 312 + 8 + 4 + 28 + 16)) {
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"] = {};
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                properties + 312 + 8 + 4 + 28 + 0,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                properties + 312 + 8 + 4 + 28 + 4,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                properties + 312 + 8 + 4 + 28 + 8,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                properties + 312 + 8 + 4 + 28 + 12,
                undefined
              );
            }
          }
          properties_ffi["WiFi"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(properties + 312 + 8 + 52, undefined);
          properties_ffi["WiFi"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(
            properties + 312 + 8 + 56,
            undefined
          );
          properties_ffi["WiFi"]["EAP"]["ClientCertRef"] = A.load.Ref(properties + 312 + 8 + 60, undefined);
          properties_ffi["WiFi"]["EAP"]["ClientCertType"] = A.load.Enum(properties + 312 + 8 + 64, ["Ref", "Pattern"]);
          properties_ffi["WiFi"]["EAP"]["Identity"] = A.load.Ref(properties + 312 + 8 + 68, undefined);
          properties_ffi["WiFi"]["EAP"]["Inner"] = A.load.Ref(properties + 312 + 8 + 72, undefined);
          properties_ffi["WiFi"]["EAP"]["Outer"] = A.load.Ref(properties + 312 + 8 + 76, undefined);
          properties_ffi["WiFi"]["EAP"]["Password"] = A.load.Ref(properties + 312 + 8 + 80, undefined);
          if (A.load.Bool(properties + 312 + 8 + 127)) {
            properties_ffi["WiFi"]["EAP"]["SaveCredentials"] = A.load.Bool(properties + 312 + 8 + 84);
          }
          properties_ffi["WiFi"]["EAP"]["ServerCAPEMs"] = A.load.Ref(properties + 312 + 8 + 88, undefined);
          properties_ffi["WiFi"]["EAP"]["ServerCARefs"] = A.load.Ref(properties + 312 + 8 + 92, undefined);
          if (A.load.Bool(properties + 312 + 8 + 96 + 28)) {
            properties_ffi["WiFi"]["EAP"]["SubjectMatch"] = {};
            properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(
              properties + 312 + 8 + 96 + 0,
              undefined
            );
            properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(
              properties + 312 + 8 + 96 + 4,
              undefined
            );
            properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(
              properties + 312 + 8 + 96 + 8,
              undefined
            );
            properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(
              properties + 312 + 8 + 96 + 12,
              undefined
            );
            properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(
              properties + 312 + 8 + 96 + 16,
              undefined
            );
            properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(
              properties + 312 + 8 + 96 + 20,
              undefined
            );
            if (A.load.Bool(properties + 312 + 8 + 96 + 26)) {
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(
                properties + 312 + 8 + 96 + 24
              );
            }
            if (A.load.Bool(properties + 312 + 8 + 96 + 27)) {
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(
                properties + 312 + 8 + 96 + 25
              );
            }
          }
          if (A.load.Bool(properties + 312 + 8 + 128)) {
            properties_ffi["WiFi"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(properties + 312 + 8 + 125);
          }
          if (A.load.Bool(properties + 312 + 8 + 129)) {
            properties_ffi["WiFi"]["EAP"]["UseSystemCAs"] = A.load.Bool(properties + 312 + 8 + 126);
          }
        }
        if (A.load.Bool(properties + 312 + 178)) {
          properties_ffi["WiFi"]["Frequency"] = A.load.Int32(properties + 312 + 140);
        }
        properties_ffi["WiFi"]["FrequencyList"] = A.load.Ref(properties + 312 + 144, undefined);
        properties_ffi["WiFi"]["HexSSID"] = A.load.Ref(properties + 312 + 148, undefined);
        if (A.load.Bool(properties + 312 + 179)) {
          properties_ffi["WiFi"]["HiddenSSID"] = A.load.Bool(properties + 312 + 152);
        }
        properties_ffi["WiFi"]["Passphrase"] = A.load.Ref(properties + 312 + 156, undefined);
        if (A.load.Bool(properties + 312 + 180)) {
          properties_ffi["WiFi"]["RoamThreshold"] = A.load.Int32(properties + 312 + 160);
        }
        properties_ffi["WiFi"]["SSID"] = A.load.Ref(properties + 312 + 164, undefined);
        properties_ffi["WiFi"]["Security"] = A.load.Ref(properties + 312 + 168, undefined);
        if (A.load.Bool(properties + 312 + 181)) {
          properties_ffi["WiFi"]["SignalStrength"] = A.load.Int32(properties + 312 + 172);
        }
      }
      if (A.load.Bool(properties + 496 + 136)) {
        properties_ffi["WiMAX"] = {};
        if (A.load.Bool(properties + 496 + 135)) {
          properties_ffi["WiMAX"]["AutoConnect"] = A.load.Bool(properties + 496 + 0);
        }
        if (A.load.Bool(properties + 496 + 4 + 130)) {
          properties_ffi["WiMAX"]["EAP"] = {};
          properties_ffi["WiMAX"]["EAP"]["AnonymousIdentity"] = A.load.Ref(properties + 496 + 4 + 0, undefined);
          if (A.load.Bool(properties + 496 + 4 + 4 + 45)) {
            properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"] = {};
            properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(
              properties + 496 + 4 + 4 + 0,
              undefined
            );
            if (A.load.Bool(properties + 496 + 4 + 4 + 4 + 16)) {
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                properties + 496 + 4 + 4 + 4 + 0,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                properties + 496 + 4 + 4 + 4 + 4,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                properties + 496 + 4 + 4 + 4 + 8,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                properties + 496 + 4 + 4 + 4 + 12,
                undefined
              );
            }
            properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(
              properties + 496 + 4 + 4 + 24,
              undefined
            );
            if (A.load.Bool(properties + 496 + 4 + 4 + 28 + 16)) {
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"] = {};
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                properties + 496 + 4 + 4 + 28 + 0,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                properties + 496 + 4 + 4 + 28 + 4,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                properties + 496 + 4 + 4 + 28 + 8,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                properties + 496 + 4 + 4 + 28 + 12,
                undefined
              );
            }
          }
          properties_ffi["WiMAX"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(properties + 496 + 4 + 52, undefined);
          properties_ffi["WiMAX"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(
            properties + 496 + 4 + 56,
            undefined
          );
          properties_ffi["WiMAX"]["EAP"]["ClientCertRef"] = A.load.Ref(properties + 496 + 4 + 60, undefined);
          properties_ffi["WiMAX"]["EAP"]["ClientCertType"] = A.load.Enum(properties + 496 + 4 + 64, ["Ref", "Pattern"]);
          properties_ffi["WiMAX"]["EAP"]["Identity"] = A.load.Ref(properties + 496 + 4 + 68, undefined);
          properties_ffi["WiMAX"]["EAP"]["Inner"] = A.load.Ref(properties + 496 + 4 + 72, undefined);
          properties_ffi["WiMAX"]["EAP"]["Outer"] = A.load.Ref(properties + 496 + 4 + 76, undefined);
          properties_ffi["WiMAX"]["EAP"]["Password"] = A.load.Ref(properties + 496 + 4 + 80, undefined);
          if (A.load.Bool(properties + 496 + 4 + 127)) {
            properties_ffi["WiMAX"]["EAP"]["SaveCredentials"] = A.load.Bool(properties + 496 + 4 + 84);
          }
          properties_ffi["WiMAX"]["EAP"]["ServerCAPEMs"] = A.load.Ref(properties + 496 + 4 + 88, undefined);
          properties_ffi["WiMAX"]["EAP"]["ServerCARefs"] = A.load.Ref(properties + 496 + 4 + 92, undefined);
          if (A.load.Bool(properties + 496 + 4 + 96 + 28)) {
            properties_ffi["WiMAX"]["EAP"]["SubjectMatch"] = {};
            properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(
              properties + 496 + 4 + 96 + 0,
              undefined
            );
            properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(
              properties + 496 + 4 + 96 + 4,
              undefined
            );
            properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(
              properties + 496 + 4 + 96 + 8,
              undefined
            );
            properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(
              properties + 496 + 4 + 96 + 12,
              undefined
            );
            properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(
              properties + 496 + 4 + 96 + 16,
              undefined
            );
            properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(
              properties + 496 + 4 + 96 + 20,
              undefined
            );
            if (A.load.Bool(properties + 496 + 4 + 96 + 26)) {
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(
                properties + 496 + 4 + 96 + 24
              );
            }
            if (A.load.Bool(properties + 496 + 4 + 96 + 27)) {
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(
                properties + 496 + 4 + 96 + 25
              );
            }
          }
          if (A.load.Bool(properties + 496 + 4 + 128)) {
            properties_ffi["WiMAX"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(properties + 496 + 4 + 125);
          }
          if (A.load.Bool(properties + 496 + 4 + 129)) {
            properties_ffi["WiMAX"]["EAP"]["UseSystemCAs"] = A.load.Bool(properties + 496 + 4 + 126);
          }
        }
      }

      const _ret = WEBEXT.networking.onc.setProperties(
        A.H.get<object>(networkGuid),
        properties_ffi,
        A.H.get<object>(callback)
      );
    },
    "try_SetProperties": (
      retPtr: Pointer,
      errPtr: Pointer,
      networkGuid: heap.Ref<object>,
      properties: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const properties_ffi = {};

        if (A.load.Bool(properties + 0 + 127)) {
          properties_ffi["Cellular"] = {};
          if (A.load.Bool(properties + 0 + 121)) {
            properties_ffi["Cellular"]["AutoConnect"] = A.load.Bool(properties + 0 + 0);
          }
          properties_ffi["Cellular"]["ActivationType"] = A.load.Ref(properties + 0 + 4, undefined);
          properties_ffi["Cellular"]["ActivationState"] = A.load.Enum(properties + 0 + 8, [
            "Activated",
            "Activating",
            "NotActivated",
            "PartiallyActivated",
          ]);
          if (A.load.Bool(properties + 0 + 122)) {
            properties_ffi["Cellular"]["AllowRoaming"] = A.load.Bool(properties + 0 + 12);
          }
          properties_ffi["Cellular"]["Family"] = A.load.Ref(properties + 0 + 16, undefined);
          properties_ffi["Cellular"]["FirmwareRevision"] = A.load.Ref(properties + 0 + 20, undefined);
          properties_ffi["Cellular"]["FoundNetworks"] = A.load.Ref(properties + 0 + 24, undefined);
          properties_ffi["Cellular"]["HardwareRevision"] = A.load.Ref(properties + 0 + 28, undefined);
          if (A.load.Bool(properties + 0 + 32 + 12)) {
            properties_ffi["Cellular"]["HomeProvider"] = {};
            properties_ffi["Cellular"]["HomeProvider"]["Name"] = A.load.Ref(properties + 0 + 32 + 0, undefined);
            properties_ffi["Cellular"]["HomeProvider"]["Code"] = A.load.Ref(properties + 0 + 32 + 4, undefined);
            properties_ffi["Cellular"]["HomeProvider"]["Country"] = A.load.Ref(properties + 0 + 32 + 8, undefined);
          }
          properties_ffi["Cellular"]["Manufacturer"] = A.load.Ref(properties + 0 + 48, undefined);
          properties_ffi["Cellular"]["ModelID"] = A.load.Ref(properties + 0 + 52, undefined);
          properties_ffi["Cellular"]["NetworkTechnology"] = A.load.Ref(properties + 0 + 56, undefined);
          if (A.load.Bool(properties + 0 + 60 + 12)) {
            properties_ffi["Cellular"]["PaymentPortal"] = {};
            properties_ffi["Cellular"]["PaymentPortal"]["Method"] = A.load.Ref(properties + 0 + 60 + 0, undefined);
            properties_ffi["Cellular"]["PaymentPortal"]["PostData"] = A.load.Ref(properties + 0 + 60 + 4, undefined);
            properties_ffi["Cellular"]["PaymentPortal"]["Url"] = A.load.Ref(properties + 0 + 60 + 8, undefined);
          }
          properties_ffi["Cellular"]["RoamingState"] = A.load.Ref(properties + 0 + 76, undefined);
          if (A.load.Bool(properties + 0 + 123)) {
            properties_ffi["Cellular"]["Scanning"] = A.load.Bool(properties + 0 + 80);
          }
          if (A.load.Bool(properties + 0 + 84 + 12)) {
            properties_ffi["Cellular"]["ServingOperator"] = {};
            properties_ffi["Cellular"]["ServingOperator"]["Name"] = A.load.Ref(properties + 0 + 84 + 0, undefined);
            properties_ffi["Cellular"]["ServingOperator"]["Code"] = A.load.Ref(properties + 0 + 84 + 4, undefined);
            properties_ffi["Cellular"]["ServingOperator"]["Country"] = A.load.Ref(properties + 0 + 84 + 8, undefined);
          }
          if (A.load.Bool(properties + 0 + 100 + 14)) {
            properties_ffi["Cellular"]["SIMLockStatus"] = {};
            properties_ffi["Cellular"]["SIMLockStatus"]["LockType"] = A.load.Ref(properties + 0 + 100 + 0, undefined);
            if (A.load.Bool(properties + 0 + 100 + 12)) {
              properties_ffi["Cellular"]["SIMLockStatus"]["LockEnabled"] = A.load.Bool(properties + 0 + 100 + 4);
            }
            if (A.load.Bool(properties + 0 + 100 + 13)) {
              properties_ffi["Cellular"]["SIMLockStatus"]["RetriesLeft"] = A.load.Int32(properties + 0 + 100 + 8);
            }
          }
          if (A.load.Bool(properties + 0 + 124)) {
            properties_ffi["Cellular"]["SIMPresent"] = A.load.Bool(properties + 0 + 115);
          }
          if (A.load.Bool(properties + 0 + 125)) {
            properties_ffi["Cellular"]["SignalStrength"] = A.load.Int32(properties + 0 + 116);
          }
          if (A.load.Bool(properties + 0 + 126)) {
            properties_ffi["Cellular"]["SupportNetworkScan"] = A.load.Bool(properties + 0 + 120);
          }
        }
        if (A.load.Bool(properties + 128 + 140)) {
          properties_ffi["Ethernet"] = {};
          if (A.load.Bool(properties + 128 + 139)) {
            properties_ffi["Ethernet"]["AutoConnect"] = A.load.Bool(properties + 128 + 0);
          }
          properties_ffi["Ethernet"]["Authentication"] = A.load.Ref(properties + 128 + 4, undefined);
          if (A.load.Bool(properties + 128 + 8 + 130)) {
            properties_ffi["Ethernet"]["EAP"] = {};
            properties_ffi["Ethernet"]["EAP"]["AnonymousIdentity"] = A.load.Ref(properties + 128 + 8 + 0, undefined);
            if (A.load.Bool(properties + 128 + 8 + 4 + 45)) {
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"] = {};
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(
                properties + 128 + 8 + 4 + 0,
                undefined
              );
              if (A.load.Bool(properties + 128 + 8 + 4 + 4 + 16)) {
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 4 + 0,
                  undefined
                );
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 4 + 4,
                  undefined
                );
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 4 + 8,
                  undefined
                );
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 4 + 12,
                  undefined
                );
              }
              properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(
                properties + 128 + 8 + 4 + 24,
                undefined
              );
              if (A.load.Bool(properties + 128 + 8 + 4 + 28 + 16)) {
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"] = {};
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 28 + 0,
                  undefined
                );
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 28 + 4,
                  undefined
                );
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 28 + 8,
                  undefined
                );
                properties_ffi["Ethernet"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                  properties + 128 + 8 + 4 + 28 + 12,
                  undefined
                );
              }
            }
            properties_ffi["Ethernet"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(properties + 128 + 8 + 52, undefined);
            properties_ffi["Ethernet"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(
              properties + 128 + 8 + 56,
              undefined
            );
            properties_ffi["Ethernet"]["EAP"]["ClientCertRef"] = A.load.Ref(properties + 128 + 8 + 60, undefined);
            properties_ffi["Ethernet"]["EAP"]["ClientCertType"] = A.load.Enum(properties + 128 + 8 + 64, [
              "Ref",
              "Pattern",
            ]);
            properties_ffi["Ethernet"]["EAP"]["Identity"] = A.load.Ref(properties + 128 + 8 + 68, undefined);
            properties_ffi["Ethernet"]["EAP"]["Inner"] = A.load.Ref(properties + 128 + 8 + 72, undefined);
            properties_ffi["Ethernet"]["EAP"]["Outer"] = A.load.Ref(properties + 128 + 8 + 76, undefined);
            properties_ffi["Ethernet"]["EAP"]["Password"] = A.load.Ref(properties + 128 + 8 + 80, undefined);
            if (A.load.Bool(properties + 128 + 8 + 127)) {
              properties_ffi["Ethernet"]["EAP"]["SaveCredentials"] = A.load.Bool(properties + 128 + 8 + 84);
            }
            properties_ffi["Ethernet"]["EAP"]["ServerCAPEMs"] = A.load.Ref(properties + 128 + 8 + 88, undefined);
            properties_ffi["Ethernet"]["EAP"]["ServerCARefs"] = A.load.Ref(properties + 128 + 8 + 92, undefined);
            if (A.load.Bool(properties + 128 + 8 + 96 + 28)) {
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"] = {};
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(
                properties + 128 + 8 + 96 + 0,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(
                properties + 128 + 8 + 96 + 4,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(
                properties + 128 + 8 + 96 + 8,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(
                properties + 128 + 8 + 96 + 12,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(
                properties + 128 + 8 + 96 + 16,
                undefined
              );
              properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(
                properties + 128 + 8 + 96 + 20,
                undefined
              );
              if (A.load.Bool(properties + 128 + 8 + 96 + 26)) {
                properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(
                  properties + 128 + 8 + 96 + 24
                );
              }
              if (A.load.Bool(properties + 128 + 8 + 96 + 27)) {
                properties_ffi["Ethernet"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(
                  properties + 128 + 8 + 96 + 25
                );
              }
            }
            if (A.load.Bool(properties + 128 + 8 + 128)) {
              properties_ffi["Ethernet"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(properties + 128 + 8 + 125);
            }
            if (A.load.Bool(properties + 128 + 8 + 129)) {
              properties_ffi["Ethernet"]["EAP"]["UseSystemCAs"] = A.load.Bool(properties + 128 + 8 + 126);
            }
          }
        }
        properties_ffi["GUID"] = A.load.Ref(properties + 272, undefined);
        properties_ffi["IPAddressConfigType"] = A.load.Enum(properties + 276, ["DHCP", "Static"]);
        properties_ffi["Name"] = A.load.Ref(properties + 280, undefined);
        properties_ffi["NameServersConfigType"] = A.load.Enum(properties + 284, ["DHCP", "Static"]);
        if (A.load.Bool(properties + 633)) {
          properties_ffi["Priority"] = A.load.Int32(properties + 288);
        }
        properties_ffi["Type"] = A.load.Enum(properties + 292, [
          "All",
          "Cellular",
          "Ethernet",
          "Tether",
          "VPN",
          "Wireless",
          "WiFi",
        ]);
        if (A.load.Bool(properties + 296 + 13)) {
          properties_ffi["VPN"] = {};
          if (A.load.Bool(properties + 296 + 12)) {
            properties_ffi["VPN"]["AutoConnect"] = A.load.Bool(properties + 296 + 0);
          }
          properties_ffi["VPN"]["Host"] = A.load.Ref(properties + 296 + 4, undefined);
          properties_ffi["VPN"]["Type"] = A.load.Ref(properties + 296 + 8, undefined);
        }
        if (A.load.Bool(properties + 312 + 182)) {
          properties_ffi["WiFi"] = {};
          if (A.load.Bool(properties + 312 + 176)) {
            properties_ffi["WiFi"]["AllowGatewayARPPolling"] = A.load.Bool(properties + 312 + 0);
          }
          if (A.load.Bool(properties + 312 + 177)) {
            properties_ffi["WiFi"]["AutoConnect"] = A.load.Bool(properties + 312 + 1);
          }
          properties_ffi["WiFi"]["BSSID"] = A.load.Ref(properties + 312 + 4, undefined);
          if (A.load.Bool(properties + 312 + 8 + 130)) {
            properties_ffi["WiFi"]["EAP"] = {};
            properties_ffi["WiFi"]["EAP"]["AnonymousIdentity"] = A.load.Ref(properties + 312 + 8 + 0, undefined);
            if (A.load.Bool(properties + 312 + 8 + 4 + 45)) {
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"] = {};
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(
                properties + 312 + 8 + 4 + 0,
                undefined
              );
              if (A.load.Bool(properties + 312 + 8 + 4 + 4 + 16)) {
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 4 + 0,
                  undefined
                );
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 4 + 4,
                  undefined
                );
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 4 + 8,
                  undefined
                );
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 4 + 12,
                  undefined
                );
              }
              properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(
                properties + 312 + 8 + 4 + 24,
                undefined
              );
              if (A.load.Bool(properties + 312 + 8 + 4 + 28 + 16)) {
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"] = {};
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 28 + 0,
                  undefined
                );
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 28 + 4,
                  undefined
                );
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 28 + 8,
                  undefined
                );
                properties_ffi["WiFi"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                  properties + 312 + 8 + 4 + 28 + 12,
                  undefined
                );
              }
            }
            properties_ffi["WiFi"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(properties + 312 + 8 + 52, undefined);
            properties_ffi["WiFi"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(
              properties + 312 + 8 + 56,
              undefined
            );
            properties_ffi["WiFi"]["EAP"]["ClientCertRef"] = A.load.Ref(properties + 312 + 8 + 60, undefined);
            properties_ffi["WiFi"]["EAP"]["ClientCertType"] = A.load.Enum(properties + 312 + 8 + 64, [
              "Ref",
              "Pattern",
            ]);
            properties_ffi["WiFi"]["EAP"]["Identity"] = A.load.Ref(properties + 312 + 8 + 68, undefined);
            properties_ffi["WiFi"]["EAP"]["Inner"] = A.load.Ref(properties + 312 + 8 + 72, undefined);
            properties_ffi["WiFi"]["EAP"]["Outer"] = A.load.Ref(properties + 312 + 8 + 76, undefined);
            properties_ffi["WiFi"]["EAP"]["Password"] = A.load.Ref(properties + 312 + 8 + 80, undefined);
            if (A.load.Bool(properties + 312 + 8 + 127)) {
              properties_ffi["WiFi"]["EAP"]["SaveCredentials"] = A.load.Bool(properties + 312 + 8 + 84);
            }
            properties_ffi["WiFi"]["EAP"]["ServerCAPEMs"] = A.load.Ref(properties + 312 + 8 + 88, undefined);
            properties_ffi["WiFi"]["EAP"]["ServerCARefs"] = A.load.Ref(properties + 312 + 8 + 92, undefined);
            if (A.load.Bool(properties + 312 + 8 + 96 + 28)) {
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"] = {};
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(
                properties + 312 + 8 + 96 + 0,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(
                properties + 312 + 8 + 96 + 4,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(
                properties + 312 + 8 + 96 + 8,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(
                properties + 312 + 8 + 96 + 12,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(
                properties + 312 + 8 + 96 + 16,
                undefined
              );
              properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(
                properties + 312 + 8 + 96 + 20,
                undefined
              );
              if (A.load.Bool(properties + 312 + 8 + 96 + 26)) {
                properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(
                  properties + 312 + 8 + 96 + 24
                );
              }
              if (A.load.Bool(properties + 312 + 8 + 96 + 27)) {
                properties_ffi["WiFi"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(
                  properties + 312 + 8 + 96 + 25
                );
              }
            }
            if (A.load.Bool(properties + 312 + 8 + 128)) {
              properties_ffi["WiFi"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(properties + 312 + 8 + 125);
            }
            if (A.load.Bool(properties + 312 + 8 + 129)) {
              properties_ffi["WiFi"]["EAP"]["UseSystemCAs"] = A.load.Bool(properties + 312 + 8 + 126);
            }
          }
          if (A.load.Bool(properties + 312 + 178)) {
            properties_ffi["WiFi"]["Frequency"] = A.load.Int32(properties + 312 + 140);
          }
          properties_ffi["WiFi"]["FrequencyList"] = A.load.Ref(properties + 312 + 144, undefined);
          properties_ffi["WiFi"]["HexSSID"] = A.load.Ref(properties + 312 + 148, undefined);
          if (A.load.Bool(properties + 312 + 179)) {
            properties_ffi["WiFi"]["HiddenSSID"] = A.load.Bool(properties + 312 + 152);
          }
          properties_ffi["WiFi"]["Passphrase"] = A.load.Ref(properties + 312 + 156, undefined);
          if (A.load.Bool(properties + 312 + 180)) {
            properties_ffi["WiFi"]["RoamThreshold"] = A.load.Int32(properties + 312 + 160);
          }
          properties_ffi["WiFi"]["SSID"] = A.load.Ref(properties + 312 + 164, undefined);
          properties_ffi["WiFi"]["Security"] = A.load.Ref(properties + 312 + 168, undefined);
          if (A.load.Bool(properties + 312 + 181)) {
            properties_ffi["WiFi"]["SignalStrength"] = A.load.Int32(properties + 312 + 172);
          }
        }
        if (A.load.Bool(properties + 496 + 136)) {
          properties_ffi["WiMAX"] = {};
          if (A.load.Bool(properties + 496 + 135)) {
            properties_ffi["WiMAX"]["AutoConnect"] = A.load.Bool(properties + 496 + 0);
          }
          if (A.load.Bool(properties + 496 + 4 + 130)) {
            properties_ffi["WiMAX"]["EAP"] = {};
            properties_ffi["WiMAX"]["EAP"]["AnonymousIdentity"] = A.load.Ref(properties + 496 + 4 + 0, undefined);
            if (A.load.Bool(properties + 496 + 4 + 4 + 45)) {
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"] = {};
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["EnrollmentURI"] = A.load.Ref(
                properties + 496 + 4 + 4 + 0,
                undefined
              );
              if (A.load.Bool(properties + 496 + 4 + 4 + 4 + 16)) {
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"] = {};
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["CommonName"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 4 + 0,
                  undefined
                );
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["Locality"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 4 + 4,
                  undefined
                );
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["Organization"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 4 + 8,
                  undefined
                );
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Issuer"]["OrganizationalUnit"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 4 + 12,
                  undefined
                );
              }
              properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["IssuerCARef"] = A.load.Ref(
                properties + 496 + 4 + 4 + 24,
                undefined
              );
              if (A.load.Bool(properties + 496 + 4 + 4 + 28 + 16)) {
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"] = {};
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["CommonName"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 28 + 0,
                  undefined
                );
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["Locality"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 28 + 4,
                  undefined
                );
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["Organization"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 28 + 8,
                  undefined
                );
                properties_ffi["WiMAX"]["EAP"]["ClientCertPattern"]["Subject"]["OrganizationalUnit"] = A.load.Ref(
                  properties + 496 + 4 + 4 + 28 + 12,
                  undefined
                );
              }
            }
            properties_ffi["WiMAX"]["EAP"]["ClientCertPKCS11Id"] = A.load.Ref(properties + 496 + 4 + 52, undefined);
            properties_ffi["WiMAX"]["EAP"]["ClientCertProvisioningProfileId"] = A.load.Ref(
              properties + 496 + 4 + 56,
              undefined
            );
            properties_ffi["WiMAX"]["EAP"]["ClientCertRef"] = A.load.Ref(properties + 496 + 4 + 60, undefined);
            properties_ffi["WiMAX"]["EAP"]["ClientCertType"] = A.load.Enum(properties + 496 + 4 + 64, [
              "Ref",
              "Pattern",
            ]);
            properties_ffi["WiMAX"]["EAP"]["Identity"] = A.load.Ref(properties + 496 + 4 + 68, undefined);
            properties_ffi["WiMAX"]["EAP"]["Inner"] = A.load.Ref(properties + 496 + 4 + 72, undefined);
            properties_ffi["WiMAX"]["EAP"]["Outer"] = A.load.Ref(properties + 496 + 4 + 76, undefined);
            properties_ffi["WiMAX"]["EAP"]["Password"] = A.load.Ref(properties + 496 + 4 + 80, undefined);
            if (A.load.Bool(properties + 496 + 4 + 127)) {
              properties_ffi["WiMAX"]["EAP"]["SaveCredentials"] = A.load.Bool(properties + 496 + 4 + 84);
            }
            properties_ffi["WiMAX"]["EAP"]["ServerCAPEMs"] = A.load.Ref(properties + 496 + 4 + 88, undefined);
            properties_ffi["WiMAX"]["EAP"]["ServerCARefs"] = A.load.Ref(properties + 496 + 4 + 92, undefined);
            if (A.load.Bool(properties + 496 + 4 + 96 + 28)) {
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"] = {};
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["Active"] = A.load.Ref(
                properties + 496 + 4 + 96 + 0,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["Effective"] = A.load.Ref(
                properties + 496 + 4 + 96 + 4,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["UserPolicy"] = A.load.Ref(
                properties + 496 + 4 + 96 + 8,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["DevicePolicy"] = A.load.Ref(
                properties + 496 + 4 + 96 + 12,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["UserSetting"] = A.load.Ref(
                properties + 496 + 4 + 96 + 16,
                undefined
              );
              properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["SharedSetting"] = A.load.Ref(
                properties + 496 + 4 + 96 + 20,
                undefined
              );
              if (A.load.Bool(properties + 496 + 4 + 96 + 26)) {
                properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["UserEditable"] = A.load.Bool(
                  properties + 496 + 4 + 96 + 24
                );
              }
              if (A.load.Bool(properties + 496 + 4 + 96 + 27)) {
                properties_ffi["WiMAX"]["EAP"]["SubjectMatch"]["DeviceEditable"] = A.load.Bool(
                  properties + 496 + 4 + 96 + 25
                );
              }
            }
            if (A.load.Bool(properties + 496 + 4 + 128)) {
              properties_ffi["WiMAX"]["EAP"]["UseProactiveKeyCaching"] = A.load.Bool(properties + 496 + 4 + 125);
            }
            if (A.load.Bool(properties + 496 + 4 + 129)) {
              properties_ffi["WiMAX"]["EAP"]["UseSystemCAs"] = A.load.Bool(properties + 496 + 4 + 126);
            }
          }
        }

        const _ret = WEBEXT.networking.onc.setProperties(
          A.H.get<object>(networkGuid),
          properties_ffi,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartConnect": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc && "startConnect" in WEBEXT?.networking?.onc) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartConnect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.startConnect);
    },
    "call_StartConnect": (retPtr: Pointer, networkGuid: heap.Ref<object>, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.startConnect(A.H.get<object>(networkGuid), A.H.get<object>(callback));
    },
    "try_StartConnect": (
      retPtr: Pointer,
      errPtr: Pointer,
      networkGuid: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.startConnect(A.H.get<object>(networkGuid), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartDisconnect": (): heap.Ref<boolean> => {
      if (WEBEXT?.networking?.onc && "startDisconnect" in WEBEXT?.networking?.onc) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartDisconnect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.networking.onc.startDisconnect);
    },
    "call_StartDisconnect": (retPtr: Pointer, networkGuid: heap.Ref<object>, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.networking.onc.startDisconnect(A.H.get<object>(networkGuid), A.H.get<object>(callback));
    },
    "try_StartDisconnect": (
      retPtr: Pointer,
      errPtr: Pointer,
      networkGuid: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.networking.onc.startDisconnect(A.H.get<object>(networkGuid), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
