import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/enterprise/reportingprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_AntiVirusProductState": (ref: heap.Ref<string>): number => {
      const idx = ["ON", "OFF", "SNOOZED", "EXPIRED"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_AntiVirusSignal": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["displayName"]);
        A.store.Ref(ptr + 4, x["productId"]);
        A.store.Enum(ptr + 8, ["ON", "OFF", "SNOOZED", "EXPIRED"].indexOf(x["state"] as string));
      }
    },
    "load_AntiVirusSignal": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["displayName"] = A.load.Ref(ptr + 0, undefined);
      x["productId"] = A.load.Ref(ptr + 4, undefined);
      x["state"] = A.load.Enum(ptr + 8, ["ON", "OFF", "SNOOZED", "EXPIRED"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_CertificateStatus": (ref: heap.Ref<string>): number => {
      const idx = ["OK", "POLICY_UNSET"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Certificate": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["OK", "POLICY_UNSET"].indexOf(x["status"] as string));
        A.store.Ref(ptr + 4, x["encodedCertificate"]);
      }
    },
    "load_Certificate": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["status"] = A.load.Enum(ptr + 0, ["OK", "POLICY_UNSET"]);
      x["encodedCertificate"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RealtimeUrlCheckMode": (ref: heap.Ref<string>): number => {
      const idx = ["DISABLED", "ENABLED_MAIN_FRAME"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SafeBrowsingLevel": (ref: heap.Ref<string>): number => {
      const idx = ["DISABLED", "STANDARD", "ENHANCED"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PasswordProtectionTrigger": (ref: heap.Ref<string>): number => {
      const idx = ["PASSWORD_PROTECTION_OFF", "PASSWORD_REUSE", "PHISHING_REUSE", "POLICY_UNSET"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SettingValue": (ref: heap.Ref<string>): number => {
      const idx = ["UNKNOWN", "DISABLED", "ENABLED"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ContextInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 68, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Enum(ptr + 24, -1);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Enum(ptr + 36, -1);
        A.store.Bool(ptr + 64, false);
        A.store.Bool(ptr + 40, false);
        A.store.Bool(ptr + 65, false);
        A.store.Bool(ptr + 41, false);
        A.store.Enum(ptr + 44, -1);
        A.store.Bool(ptr + 66, false);
        A.store.Bool(ptr + 48, false);
        A.store.Bool(ptr + 67, false);
        A.store.Bool(ptr + 49, false);
        A.store.Enum(ptr + 52, -1);
        A.store.Ref(ptr + 56, undefined);
        A.store.Ref(ptr + 60, undefined);
      } else {
        A.store.Bool(ptr + 68, true);
        A.store.Ref(ptr + 0, x["browserAffiliationIds"]);
        A.store.Ref(ptr + 4, x["profileAffiliationIds"]);
        A.store.Ref(ptr + 8, x["onFileAttachedProviders"]);
        A.store.Ref(ptr + 12, x["onFileDownloadedProviders"]);
        A.store.Ref(ptr + 16, x["onBulkDataEntryProviders"]);
        A.store.Ref(ptr + 20, x["onPrintProviders"]);
        A.store.Enum(ptr + 24, ["DISABLED", "ENABLED_MAIN_FRAME"].indexOf(x["realtimeUrlCheckMode"] as string));
        A.store.Ref(ptr + 28, x["onSecurityEventProviders"]);
        A.store.Ref(ptr + 32, x["browserVersion"]);
        A.store.Enum(
          ptr + 36,
          ["DISABLED", "STANDARD", "ENHANCED"].indexOf(x["safeBrowsingProtectionLevel"] as string)
        );
        A.store.Bool(ptr + 64, "siteIsolationEnabled" in x ? true : false);
        A.store.Bool(ptr + 40, x["siteIsolationEnabled"] ? true : false);
        A.store.Bool(ptr + 65, "builtInDnsClientEnabled" in x ? true : false);
        A.store.Bool(ptr + 41, x["builtInDnsClientEnabled"] ? true : false);
        A.store.Enum(
          ptr + 44,
          ["PASSWORD_PROTECTION_OFF", "PASSWORD_REUSE", "PHISHING_REUSE", "POLICY_UNSET"].indexOf(
            x["passwordProtectionWarningTrigger"] as string
          )
        );
        A.store.Bool(ptr + 66, "chromeRemoteDesktopAppBlocked" in x ? true : false);
        A.store.Bool(ptr + 48, x["chromeRemoteDesktopAppBlocked"] ? true : false);
        A.store.Bool(ptr + 67, "thirdPartyBlockingEnabled" in x ? true : false);
        A.store.Bool(ptr + 49, x["thirdPartyBlockingEnabled"] ? true : false);
        A.store.Enum(ptr + 52, ["UNKNOWN", "DISABLED", "ENABLED"].indexOf(x["osFirewall"] as string));
        A.store.Ref(ptr + 56, x["systemDnsServers"]);
        A.store.Ref(ptr + 60, x["enterpriseProfileId"]);
      }
    },
    "load_ContextInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["browserAffiliationIds"] = A.load.Ref(ptr + 0, undefined);
      x["profileAffiliationIds"] = A.load.Ref(ptr + 4, undefined);
      x["onFileAttachedProviders"] = A.load.Ref(ptr + 8, undefined);
      x["onFileDownloadedProviders"] = A.load.Ref(ptr + 12, undefined);
      x["onBulkDataEntryProviders"] = A.load.Ref(ptr + 16, undefined);
      x["onPrintProviders"] = A.load.Ref(ptr + 20, undefined);
      x["realtimeUrlCheckMode"] = A.load.Enum(ptr + 24, ["DISABLED", "ENABLED_MAIN_FRAME"]);
      x["onSecurityEventProviders"] = A.load.Ref(ptr + 28, undefined);
      x["browserVersion"] = A.load.Ref(ptr + 32, undefined);
      x["safeBrowsingProtectionLevel"] = A.load.Enum(ptr + 36, ["DISABLED", "STANDARD", "ENHANCED"]);
      if (A.load.Bool(ptr + 64)) {
        x["siteIsolationEnabled"] = A.load.Bool(ptr + 40);
      } else {
        delete x["siteIsolationEnabled"];
      }
      if (A.load.Bool(ptr + 65)) {
        x["builtInDnsClientEnabled"] = A.load.Bool(ptr + 41);
      } else {
        delete x["builtInDnsClientEnabled"];
      }
      x["passwordProtectionWarningTrigger"] = A.load.Enum(ptr + 44, [
        "PASSWORD_PROTECTION_OFF",
        "PASSWORD_REUSE",
        "PHISHING_REUSE",
        "POLICY_UNSET",
      ]);
      if (A.load.Bool(ptr + 66)) {
        x["chromeRemoteDesktopAppBlocked"] = A.load.Bool(ptr + 48);
      } else {
        delete x["chromeRemoteDesktopAppBlocked"];
      }
      if (A.load.Bool(ptr + 67)) {
        x["thirdPartyBlockingEnabled"] = A.load.Bool(ptr + 49);
      } else {
        delete x["thirdPartyBlockingEnabled"];
      }
      x["osFirewall"] = A.load.Enum(ptr + 52, ["UNKNOWN", "DISABLED", "ENABLED"]);
      x["systemDnsServers"] = A.load.Ref(ptr + 56, undefined);
      x["enterpriseProfileId"] = A.load.Ref(ptr + 60, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DeviceInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 48, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Enum(ptr + 20, -1);
        A.store.Enum(ptr + 24, -1);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
        A.store.Enum(ptr + 44, -1);
      } else {
        A.store.Bool(ptr + 48, true);
        A.store.Ref(ptr + 0, x["osName"]);
        A.store.Ref(ptr + 4, x["osVersion"]);
        A.store.Ref(ptr + 8, x["deviceHostName"]);
        A.store.Ref(ptr + 12, x["deviceModel"]);
        A.store.Ref(ptr + 16, x["serialNumber"]);
        A.store.Enum(ptr + 20, ["UNKNOWN", "DISABLED", "ENABLED"].indexOf(x["screenLockSecured"] as string));
        A.store.Enum(ptr + 24, ["UNKNOWN", "DISABLED", "ENABLED"].indexOf(x["diskEncrypted"] as string));
        A.store.Ref(ptr + 28, x["macAddresses"]);
        A.store.Ref(ptr + 32, x["windowsMachineDomain"]);
        A.store.Ref(ptr + 36, x["windowsUserDomain"]);
        A.store.Ref(ptr + 40, x["securityPatchLevel"]);
        A.store.Enum(ptr + 44, ["UNKNOWN", "DISABLED", "ENABLED"].indexOf(x["secureBootEnabled"] as string));
      }
    },
    "load_DeviceInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["osName"] = A.load.Ref(ptr + 0, undefined);
      x["osVersion"] = A.load.Ref(ptr + 4, undefined);
      x["deviceHostName"] = A.load.Ref(ptr + 8, undefined);
      x["deviceModel"] = A.load.Ref(ptr + 12, undefined);
      x["serialNumber"] = A.load.Ref(ptr + 16, undefined);
      x["screenLockSecured"] = A.load.Enum(ptr + 20, ["UNKNOWN", "DISABLED", "ENABLED"]);
      x["diskEncrypted"] = A.load.Enum(ptr + 24, ["UNKNOWN", "DISABLED", "ENABLED"]);
      x["macAddresses"] = A.load.Ref(ptr + 28, undefined);
      x["windowsMachineDomain"] = A.load.Ref(ptr + 32, undefined);
      x["windowsUserDomain"] = A.load.Ref(ptr + 36, undefined);
      x["securityPatchLevel"] = A.load.Ref(ptr + 40, undefined);
      x["secureBootEnabled"] = A.load.Enum(ptr + 44, ["UNKNOWN", "DISABLED", "ENABLED"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_EventType": (ref: heap.Ref<string>): number => {
      const idx = ["DEVICE", "USER"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_EnqueueRecordRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Ref(ptr + 0, x["recordData"]);
        A.store.Bool(ptr + 12, "priority" in x ? true : false);
        A.store.Int32(ptr + 4, x["priority"] === undefined ? 0 : (x["priority"] as number));
        A.store.Enum(ptr + 8, ["DEVICE", "USER"].indexOf(x["eventType"] as string));
      }
    },
    "load_EnqueueRecordRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["recordData"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["priority"] = A.load.Int32(ptr + 4);
      } else {
        delete x["priority"];
      }
      x["eventType"] = A.load.Enum(ptr + 8, ["DEVICE", "USER"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PresenceValue": (ref: heap.Ref<string>): number => {
      const idx = ["UNSPECIFIED", "ACCESS_DENIED", "NOT_FOUND", "FOUND"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_GetFileSystemInfoResponse": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 29, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 29, true);
        A.store.Ref(ptr + 0, x["path"]);
        A.store.Enum(ptr + 4, ["UNSPECIFIED", "ACCESS_DENIED", "NOT_FOUND", "FOUND"].indexOf(x["presence"] as string));
        A.store.Ref(ptr + 8, x["sha256Hash"]);
        A.store.Bool(ptr + 28, "isRunning" in x ? true : false);
        A.store.Bool(ptr + 12, x["isRunning"] ? true : false);
        A.store.Ref(ptr + 16, x["publicKeysHashes"]);
        A.store.Ref(ptr + 20, x["productName"]);
        A.store.Ref(ptr + 24, x["version"]);
      }
    },
    "load_GetFileSystemInfoResponse": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["path"] = A.load.Ref(ptr + 0, undefined);
      x["presence"] = A.load.Enum(ptr + 4, ["UNSPECIFIED", "ACCESS_DENIED", "NOT_FOUND", "FOUND"]);
      x["sha256Hash"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 28)) {
        x["isRunning"] = A.load.Bool(ptr + 12);
      } else {
        delete x["isRunning"];
      }
      x["publicKeysHashes"] = A.load.Ref(ptr + 16, undefined);
      x["productName"] = A.load.Ref(ptr + 20, undefined);
      x["version"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetFileSystemInfoOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Ref(ptr + 0, x["path"]);
        A.store.Bool(ptr + 6, "computeSha256" in x ? true : false);
        A.store.Bool(ptr + 4, x["computeSha256"] ? true : false);
        A.store.Bool(ptr + 7, "computeExecutableMetadata" in x ? true : false);
        A.store.Bool(ptr + 5, x["computeExecutableMetadata"] ? true : false);
      }
    },
    "load_GetFileSystemInfoOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["path"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 6)) {
        x["computeSha256"] = A.load.Bool(ptr + 4);
      } else {
        delete x["computeSha256"];
      }
      if (A.load.Bool(ptr + 7)) {
        x["computeExecutableMetadata"] = A.load.Bool(ptr + 5);
      } else {
        delete x["computeExecutableMetadata"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UserContext": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["userId"]);
      }
    },
    "load_UserContext": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["userId"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetFileSystemInfoRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);

        A.store.Bool(ptr + 0 + 4, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);

        if (typeof x["userContext"] === "undefined") {
          A.store.Bool(ptr + 0 + 4, false);
          A.store.Ref(ptr + 0 + 0, undefined);
        } else {
          A.store.Bool(ptr + 0 + 4, true);
          A.store.Ref(ptr + 0 + 0, x["userContext"]["userId"]);
        }
        A.store.Ref(ptr + 8, x["options"]);
      }
    },
    "load_GetFileSystemInfoRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 4)) {
        x["userContext"] = {};
        x["userContext"]["userId"] = A.load.Ref(ptr + 0 + 0, undefined);
      } else {
        delete x["userContext"];
      }
      x["options"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RegistryHive": (ref: heap.Ref<string>): number => {
      const idx = ["HKEY_CLASSES_ROOT", "HKEY_LOCAL_MACHINE", "HKEY_CURRENT_USER"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_GetSettingsOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 12, -1);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["path"]);
        A.store.Ref(ptr + 4, x["key"]);
        A.store.Bool(ptr + 16, "getValue" in x ? true : false);
        A.store.Bool(ptr + 8, x["getValue"] ? true : false);
        A.store.Enum(
          ptr + 12,
          ["HKEY_CLASSES_ROOT", "HKEY_LOCAL_MACHINE", "HKEY_CURRENT_USER"].indexOf(x["hive"] as string)
        );
      }
    },
    "load_GetSettingsOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["path"] = A.load.Ref(ptr + 0, undefined);
      x["key"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["getValue"] = A.load.Bool(ptr + 8);
      } else {
        delete x["getValue"];
      }
      x["hive"] = A.load.Enum(ptr + 12, ["HKEY_CLASSES_ROOT", "HKEY_LOCAL_MACHINE", "HKEY_CURRENT_USER"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetSettingsRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);

        A.store.Bool(ptr + 0 + 4, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);

        if (typeof x["userContext"] === "undefined") {
          A.store.Bool(ptr + 0 + 4, false);
          A.store.Ref(ptr + 0 + 0, undefined);
        } else {
          A.store.Bool(ptr + 0 + 4, true);
          A.store.Ref(ptr + 0 + 0, x["userContext"]["userId"]);
        }
        A.store.Ref(ptr + 8, x["options"]);
      }
    },
    "load_GetSettingsRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 4)) {
        x["userContext"] = {};
        x["userContext"]["userId"] = A.load.Ref(ptr + 0 + 0, undefined);
      } else {
        delete x["userContext"];
      }
      x["options"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetSettingsResponse": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Enum(ptr + 12, -1);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Ref(ptr + 0, x["path"]);
        A.store.Ref(ptr + 4, x["key"]);
        A.store.Enum(
          ptr + 8,
          ["HKEY_CLASSES_ROOT", "HKEY_LOCAL_MACHINE", "HKEY_CURRENT_USER"].indexOf(x["hive"] as string)
        );
        A.store.Enum(ptr + 12, ["UNSPECIFIED", "ACCESS_DENIED", "NOT_FOUND", "FOUND"].indexOf(x["presence"] as string));
        A.store.Ref(ptr + 16, x["value"]);
      }
    },
    "load_GetSettingsResponse": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["path"] = A.load.Ref(ptr + 0, undefined);
      x["key"] = A.load.Ref(ptr + 4, undefined);
      x["hive"] = A.load.Enum(ptr + 8, ["HKEY_CLASSES_ROOT", "HKEY_LOCAL_MACHINE", "HKEY_CURRENT_USER"]);
      x["presence"] = A.load.Enum(ptr + 12, ["UNSPECIFIED", "ACCESS_DENIED", "NOT_FOUND", "FOUND"]);
      x["value"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_HotfixSignal": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["hotfixId"]);
      }
    },
    "load_HotfixSignal": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["hotfixId"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_EnqueueRecord": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.reportingPrivate && "enqueueRecord" in WEBEXT?.enterprise?.reportingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EnqueueRecord": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.reportingPrivate.enqueueRecord);
    },
    "call_EnqueueRecord": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      request_ffi["recordData"] = A.load.Ref(request + 0, undefined);
      if (A.load.Bool(request + 12)) {
        request_ffi["priority"] = A.load.Int32(request + 4);
      }
      request_ffi["eventType"] = A.load.Enum(request + 8, ["DEVICE", "USER"]);

      const _ret = WEBEXT.enterprise.reportingPrivate.enqueueRecord(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_EnqueueRecord": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        request_ffi["recordData"] = A.load.Ref(request + 0, undefined);
        if (A.load.Bool(request + 12)) {
          request_ffi["priority"] = A.load.Int32(request + 4);
        }
        request_ffi["eventType"] = A.load.Enum(request + 8, ["DEVICE", "USER"]);

        const _ret = WEBEXT.enterprise.reportingPrivate.enqueueRecord(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAvInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.reportingPrivate && "getAvInfo" in WEBEXT?.enterprise?.reportingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAvInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.reportingPrivate.getAvInfo);
    },
    "call_GetAvInfo": (retPtr: Pointer, userContext: Pointer): void => {
      const userContext_ffi = {};

      userContext_ffi["userId"] = A.load.Ref(userContext + 0, undefined);

      const _ret = WEBEXT.enterprise.reportingPrivate.getAvInfo(userContext_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAvInfo": (retPtr: Pointer, errPtr: Pointer, userContext: Pointer): heap.Ref<boolean> => {
      try {
        const userContext_ffi = {};

        userContext_ffi["userId"] = A.load.Ref(userContext + 0, undefined);

        const _ret = WEBEXT.enterprise.reportingPrivate.getAvInfo(userContext_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCertificate": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.reportingPrivate && "getCertificate" in WEBEXT?.enterprise?.reportingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCertificate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.reportingPrivate.getCertificate);
    },
    "call_GetCertificate": (retPtr: Pointer, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.enterprise.reportingPrivate.getCertificate(A.H.get<object>(url));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCertificate": (retPtr: Pointer, errPtr: Pointer, url: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.reportingPrivate.getCertificate(A.H.get<object>(url));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetContextInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.reportingPrivate && "getContextInfo" in WEBEXT?.enterprise?.reportingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetContextInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.reportingPrivate.getContextInfo);
    },
    "call_GetContextInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.enterprise.reportingPrivate.getContextInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetContextInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.reportingPrivate.getContextInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDeviceData": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.reportingPrivate && "getDeviceData" in WEBEXT?.enterprise?.reportingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDeviceData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.reportingPrivate.getDeviceData);
    },
    "call_GetDeviceData": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.enterprise.reportingPrivate.getDeviceData(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDeviceData": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.reportingPrivate.getDeviceData(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDeviceId": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.reportingPrivate && "getDeviceId" in WEBEXT?.enterprise?.reportingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDeviceId": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.reportingPrivate.getDeviceId);
    },
    "call_GetDeviceId": (retPtr: Pointer): void => {
      const _ret = WEBEXT.enterprise.reportingPrivate.getDeviceId();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDeviceId": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.reportingPrivate.getDeviceId();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDeviceInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.reportingPrivate && "getDeviceInfo" in WEBEXT?.enterprise?.reportingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDeviceInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.reportingPrivate.getDeviceInfo);
    },
    "call_GetDeviceInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.enterprise.reportingPrivate.getDeviceInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDeviceInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.reportingPrivate.getDeviceInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetFileSystemInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.reportingPrivate && "getFileSystemInfo" in WEBEXT?.enterprise?.reportingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetFileSystemInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.reportingPrivate.getFileSystemInfo);
    },
    "call_GetFileSystemInfo": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 0 + 4)) {
        request_ffi["userContext"] = {};
        request_ffi["userContext"]["userId"] = A.load.Ref(request + 0 + 0, undefined);
      }
      request_ffi["options"] = A.load.Ref(request + 8, undefined);

      const _ret = WEBEXT.enterprise.reportingPrivate.getFileSystemInfo(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetFileSystemInfo": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 0 + 4)) {
          request_ffi["userContext"] = {};
          request_ffi["userContext"]["userId"] = A.load.Ref(request + 0 + 0, undefined);
        }
        request_ffi["options"] = A.load.Ref(request + 8, undefined);

        const _ret = WEBEXT.enterprise.reportingPrivate.getFileSystemInfo(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetHotfixes": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.reportingPrivate && "getHotfixes" in WEBEXT?.enterprise?.reportingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetHotfixes": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.reportingPrivate.getHotfixes);
    },
    "call_GetHotfixes": (retPtr: Pointer, userContext: Pointer): void => {
      const userContext_ffi = {};

      userContext_ffi["userId"] = A.load.Ref(userContext + 0, undefined);

      const _ret = WEBEXT.enterprise.reportingPrivate.getHotfixes(userContext_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetHotfixes": (retPtr: Pointer, errPtr: Pointer, userContext: Pointer): heap.Ref<boolean> => {
      try {
        const userContext_ffi = {};

        userContext_ffi["userId"] = A.load.Ref(userContext + 0, undefined);

        const _ret = WEBEXT.enterprise.reportingPrivate.getHotfixes(userContext_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPersistentSecret": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.reportingPrivate && "getPersistentSecret" in WEBEXT?.enterprise?.reportingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPersistentSecret": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.reportingPrivate.getPersistentSecret);
    },
    "call_GetPersistentSecret": (retPtr: Pointer, resetSecret: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.enterprise.reportingPrivate.getPersistentSecret(resetSecret === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPersistentSecret": (
      retPtr: Pointer,
      errPtr: Pointer,
      resetSecret: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.reportingPrivate.getPersistentSecret(resetSecret === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSettings": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.reportingPrivate && "getSettings" in WEBEXT?.enterprise?.reportingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSettings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.reportingPrivate.getSettings);
    },
    "call_GetSettings": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 0 + 4)) {
        request_ffi["userContext"] = {};
        request_ffi["userContext"]["userId"] = A.load.Ref(request + 0 + 0, undefined);
      }
      request_ffi["options"] = A.load.Ref(request + 8, undefined);

      const _ret = WEBEXT.enterprise.reportingPrivate.getSettings(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSettings": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 0 + 4)) {
          request_ffi["userContext"] = {};
          request_ffi["userContext"]["userId"] = A.load.Ref(request + 0 + 0, undefined);
        }
        request_ffi["options"] = A.load.Ref(request + 8, undefined);

        const _ret = WEBEXT.enterprise.reportingPrivate.getSettings(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetDeviceData": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.reportingPrivate && "setDeviceData" in WEBEXT?.enterprise?.reportingPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetDeviceData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.reportingPrivate.setDeviceData);
    },
    "call_SetDeviceData": (retPtr: Pointer, id: heap.Ref<object>, data: heap.Ref<object>): void => {
      const _ret = WEBEXT.enterprise.reportingPrivate.setDeviceData(A.H.get<object>(id), A.H.get<object>(data));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetDeviceData": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<object>,
      data: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.reportingPrivate.setDeviceData(A.H.get<object>(id), A.H.get<object>(data));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
