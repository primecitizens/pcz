import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/passwordsprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AddPasswordOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 16, false);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Ref(ptr + 0, x["url"]);
        A.store.Ref(ptr + 4, x["username"]);
        A.store.Ref(ptr + 8, x["password"]);
        A.store.Ref(ptr + 12, x["note"]);
        A.store.Bool(ptr + 17, "useAccountStore" in x ? true : false);
        A.store.Bool(ptr + 16, x["useAccountStore"] ? true : false);
      }
    },
    "load_AddPasswordOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      x["username"] = A.load.Ref(ptr + 4, undefined);
      x["password"] = A.load.Ref(ptr + 8, undefined);
      x["note"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 17)) {
        x["useAccountStore"] = A.load.Bool(ptr + 16);
      } else {
        delete x["useAccountStore"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_CompromiseType": (ref: heap.Ref<string>): number => {
      const idx = ["LEAKED", "PHISHED", "REUSED", "WEAK"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_CompromisedInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 17, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
      } else {
        A.store.Bool(ptr + 19, true);
        A.store.Bool(ptr + 17, "compromiseTime" in x ? true : false);
        A.store.Float64(ptr + 0, x["compromiseTime"] === undefined ? 0 : (x["compromiseTime"] as number));
        A.store.Ref(ptr + 8, x["elapsedTimeSinceCompromise"]);
        A.store.Ref(ptr + 12, x["compromiseTypes"]);
        A.store.Bool(ptr + 18, "isMuted" in x ? true : false);
        A.store.Bool(ptr + 16, x["isMuted"] ? true : false);
      }
    },
    "load_CompromisedInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 17)) {
        x["compromiseTime"] = A.load.Float64(ptr + 0);
      } else {
        delete x["compromiseTime"];
      }
      x["elapsedTimeSinceCompromise"] = A.load.Ref(ptr + 8, undefined);
      x["compromiseTypes"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 18)) {
        x["isMuted"] = A.load.Bool(ptr + 16);
      } else {
        delete x["isMuted"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DomainInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["url"]);
        A.store.Ref(ptr + 8, x["signonRealm"]);
      }
    },
    "load_DomainInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["url"] = A.load.Ref(ptr + 4, undefined);
      x["signonRealm"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PasswordStoreSet": (ref: heap.Ref<string>): number => {
      const idx = ["DEVICE", "ACCOUNT", "DEVICE_AND_ACCOUNT"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PasswordUiEntry": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 62, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 60, false);
        A.store.Int32(ptr + 20, 0);
        A.store.Enum(ptr + 24, -1);
        A.store.Bool(ptr + 61, false);
        A.store.Bool(ptr + 28, false);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);

        A.store.Bool(ptr + 40 + 19, false);
        A.store.Bool(ptr + 40 + 17, false);
        A.store.Float64(ptr + 40 + 0, 0);
        A.store.Ref(ptr + 40 + 8, undefined);
        A.store.Ref(ptr + 40 + 12, undefined);
        A.store.Bool(ptr + 40 + 18, false);
        A.store.Bool(ptr + 40 + 16, false);
      } else {
        A.store.Bool(ptr + 62, true);
        A.store.Ref(ptr + 0, x["affiliatedDomains"]);
        A.store.Ref(ptr + 4, x["username"]);
        A.store.Ref(ptr + 8, x["displayName"]);
        A.store.Ref(ptr + 12, x["password"]);
        A.store.Ref(ptr + 16, x["federationText"]);
        A.store.Bool(ptr + 60, "id" in x ? true : false);
        A.store.Int32(ptr + 20, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Enum(ptr + 24, ["DEVICE", "ACCOUNT", "DEVICE_AND_ACCOUNT"].indexOf(x["storedIn"] as string));
        A.store.Bool(ptr + 61, "isPasskey" in x ? true : false);
        A.store.Bool(ptr + 28, x["isPasskey"] ? true : false);
        A.store.Ref(ptr + 32, x["note"]);
        A.store.Ref(ptr + 36, x["changePasswordUrl"]);

        if (typeof x["compromisedInfo"] === "undefined") {
          A.store.Bool(ptr + 40 + 19, false);
          A.store.Bool(ptr + 40 + 17, false);
          A.store.Float64(ptr + 40 + 0, 0);
          A.store.Ref(ptr + 40 + 8, undefined);
          A.store.Ref(ptr + 40 + 12, undefined);
          A.store.Bool(ptr + 40 + 18, false);
          A.store.Bool(ptr + 40 + 16, false);
        } else {
          A.store.Bool(ptr + 40 + 19, true);
          A.store.Bool(ptr + 40 + 17, "compromiseTime" in x["compromisedInfo"] ? true : false);
          A.store.Float64(
            ptr + 40 + 0,
            x["compromisedInfo"]["compromiseTime"] === undefined
              ? 0
              : (x["compromisedInfo"]["compromiseTime"] as number)
          );
          A.store.Ref(ptr + 40 + 8, x["compromisedInfo"]["elapsedTimeSinceCompromise"]);
          A.store.Ref(ptr + 40 + 12, x["compromisedInfo"]["compromiseTypes"]);
          A.store.Bool(ptr + 40 + 18, "isMuted" in x["compromisedInfo"] ? true : false);
          A.store.Bool(ptr + 40 + 16, x["compromisedInfo"]["isMuted"] ? true : false);
        }
      }
    },
    "load_PasswordUiEntry": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["affiliatedDomains"] = A.load.Ref(ptr + 0, undefined);
      x["username"] = A.load.Ref(ptr + 4, undefined);
      x["displayName"] = A.load.Ref(ptr + 8, undefined);
      x["password"] = A.load.Ref(ptr + 12, undefined);
      x["federationText"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 60)) {
        x["id"] = A.load.Int32(ptr + 20);
      } else {
        delete x["id"];
      }
      x["storedIn"] = A.load.Enum(ptr + 24, ["DEVICE", "ACCOUNT", "DEVICE_AND_ACCOUNT"]);
      if (A.load.Bool(ptr + 61)) {
        x["isPasskey"] = A.load.Bool(ptr + 28);
      } else {
        delete x["isPasskey"];
      }
      x["note"] = A.load.Ref(ptr + 32, undefined);
      x["changePasswordUrl"] = A.load.Ref(ptr + 36, undefined);
      if (A.load.Bool(ptr + 40 + 19)) {
        x["compromisedInfo"] = {};
        if (A.load.Bool(ptr + 40 + 17)) {
          x["compromisedInfo"]["compromiseTime"] = A.load.Float64(ptr + 40 + 0);
        } else {
          delete x["compromisedInfo"]["compromiseTime"];
        }
        x["compromisedInfo"]["elapsedTimeSinceCompromise"] = A.load.Ref(ptr + 40 + 8, undefined);
        x["compromisedInfo"]["compromiseTypes"] = A.load.Ref(ptr + 40 + 12, undefined);
        if (A.load.Bool(ptr + 40 + 18)) {
          x["compromisedInfo"]["isMuted"] = A.load.Bool(ptr + 40 + 16);
        } else {
          delete x["compromisedInfo"]["isMuted"];
        }
      } else {
        delete x["compromisedInfo"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CredentialGroup": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["iconUrl"]);
        A.store.Ref(ptr + 8, x["entries"]);
      }
    },
    "load_CredentialGroup": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["iconUrl"] = A.load.Ref(ptr + 4, undefined);
      x["entries"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PasswordUiEntryList": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["entries"]);
      }
    },
    "load_PasswordUiEntryList": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["entries"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UrlCollection": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["signonRealm"]);
        A.store.Ref(ptr + 4, x["shown"]);
        A.store.Ref(ptr + 8, x["link"]);
      }
    },
    "load_UrlCollection": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["signonRealm"] = A.load.Ref(ptr + 0, undefined);
      x["shown"] = A.load.Ref(ptr + 4, undefined);
      x["link"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ExceptionEntry": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);

        A.store.Bool(ptr + 0 + 12, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
        A.store.Ref(ptr + 0 + 8, undefined);
        A.store.Bool(ptr + 20, false);
        A.store.Int32(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 21, true);

        if (typeof x["urls"] === "undefined") {
          A.store.Bool(ptr + 0 + 12, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
          A.store.Ref(ptr + 0 + 8, undefined);
        } else {
          A.store.Bool(ptr + 0 + 12, true);
          A.store.Ref(ptr + 0 + 0, x["urls"]["signonRealm"]);
          A.store.Ref(ptr + 0 + 4, x["urls"]["shown"]);
          A.store.Ref(ptr + 0 + 8, x["urls"]["link"]);
        }
        A.store.Bool(ptr + 20, "id" in x ? true : false);
        A.store.Int32(ptr + 16, x["id"] === undefined ? 0 : (x["id"] as number));
      }
    },
    "load_ExceptionEntry": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 12)) {
        x["urls"] = {};
        x["urls"]["signonRealm"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["urls"]["shown"] = A.load.Ref(ptr + 0 + 4, undefined);
        x["urls"]["link"] = A.load.Ref(ptr + 0 + 8, undefined);
      } else {
        delete x["urls"];
      }
      if (A.load.Bool(ptr + 20)) {
        x["id"] = A.load.Int32(ptr + 16);
      } else {
        delete x["id"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ExportProgressStatus": (ref: heap.Ref<string>): number => {
      const idx = ["NOT_STARTED", "IN_PROGRESS", "SUCCEEDED", "FAILED_CANCELLED", "FAILED_WRITE_FAILED"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_FamilyFetchStatus": (ref: heap.Ref<string>): number => {
      const idx = ["UNKNOWN_ERROR", "NO_MEMBERS", "SUCCESS"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PublicKey": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Ref(ptr + 0, x["value"]);
        A.store.Bool(ptr + 8, "version" in x ? true : false);
        A.store.Int32(ptr + 4, x["version"] === undefined ? 0 : (x["version"] as number));
      }
    },
    "load_PublicKey": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["value"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 8)) {
        x["version"] = A.load.Int32(ptr + 4);
      } else {
        delete x["version"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RecipientInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 31, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 30, false);
        A.store.Bool(ptr + 16, false);

        A.store.Bool(ptr + 20 + 9, false);
        A.store.Ref(ptr + 20 + 0, undefined);
        A.store.Bool(ptr + 20 + 8, false);
        A.store.Int32(ptr + 20 + 4, 0);
      } else {
        A.store.Bool(ptr + 31, true);
        A.store.Ref(ptr + 0, x["userId"]);
        A.store.Ref(ptr + 4, x["email"]);
        A.store.Ref(ptr + 8, x["displayName"]);
        A.store.Ref(ptr + 12, x["profileImageUrl"]);
        A.store.Bool(ptr + 30, "isEligible" in x ? true : false);
        A.store.Bool(ptr + 16, x["isEligible"] ? true : false);

        if (typeof x["publicKey"] === "undefined") {
          A.store.Bool(ptr + 20 + 9, false);
          A.store.Ref(ptr + 20 + 0, undefined);
          A.store.Bool(ptr + 20 + 8, false);
          A.store.Int32(ptr + 20 + 4, 0);
        } else {
          A.store.Bool(ptr + 20 + 9, true);
          A.store.Ref(ptr + 20 + 0, x["publicKey"]["value"]);
          A.store.Bool(ptr + 20 + 8, "version" in x["publicKey"] ? true : false);
          A.store.Int32(
            ptr + 20 + 4,
            x["publicKey"]["version"] === undefined ? 0 : (x["publicKey"]["version"] as number)
          );
        }
      }
    },
    "load_RecipientInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["userId"] = A.load.Ref(ptr + 0, undefined);
      x["email"] = A.load.Ref(ptr + 4, undefined);
      x["displayName"] = A.load.Ref(ptr + 8, undefined);
      x["profileImageUrl"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 30)) {
        x["isEligible"] = A.load.Bool(ptr + 16);
      } else {
        delete x["isEligible"];
      }
      if (A.load.Bool(ptr + 20 + 9)) {
        x["publicKey"] = {};
        x["publicKey"]["value"] = A.load.Ref(ptr + 20 + 0, undefined);
        if (A.load.Bool(ptr + 20 + 8)) {
          x["publicKey"]["version"] = A.load.Int32(ptr + 20 + 4);
        } else {
          delete x["publicKey"]["version"];
        }
      } else {
        delete x["publicKey"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_FamilyFetchResults": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["UNKNOWN_ERROR", "NO_MEMBERS", "SUCCESS"].indexOf(x["status"] as string));
        A.store.Ref(ptr + 4, x["familyMembers"]);
      }
    },
    "load_FamilyFetchResults": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["status"] = A.load.Enum(ptr + 0, ["UNKNOWN_ERROR", "NO_MEMBERS", "SUCCESS"]);
      x["familyMembers"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ImportEntryStatus": (ref: heap.Ref<string>): number => {
      const idx = [
        "UNKNOWN_ERROR",
        "MISSING_PASSWORD",
        "MISSING_URL",
        "INVALID_URL",
        "NON_ASCII_URL",
        "LONG_URL",
        "LONG_PASSWORD",
        "LONG_USERNAME",
        "CONFLICT_PROFILE",
        "CONFLICT_ACCOUNT",
        "LONG_NOTE",
        "LONG_CONCATENATED_NOTE",
        "VALID",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ImportEntry": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 20, false);
        A.store.Int32(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Enum(
          ptr + 0,
          [
            "UNKNOWN_ERROR",
            "MISSING_PASSWORD",
            "MISSING_URL",
            "INVALID_URL",
            "NON_ASCII_URL",
            "LONG_URL",
            "LONG_PASSWORD",
            "LONG_USERNAME",
            "CONFLICT_PROFILE",
            "CONFLICT_ACCOUNT",
            "LONG_NOTE",
            "LONG_CONCATENATED_NOTE",
            "VALID",
          ].indexOf(x["status"] as string)
        );
        A.store.Ref(ptr + 4, x["url"]);
        A.store.Ref(ptr + 8, x["username"]);
        A.store.Ref(ptr + 12, x["password"]);
        A.store.Bool(ptr + 20, "id" in x ? true : false);
        A.store.Int32(ptr + 16, x["id"] === undefined ? 0 : (x["id"] as number));
      }
    },
    "load_ImportEntry": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["status"] = A.load.Enum(ptr + 0, [
        "UNKNOWN_ERROR",
        "MISSING_PASSWORD",
        "MISSING_URL",
        "INVALID_URL",
        "NON_ASCII_URL",
        "LONG_URL",
        "LONG_PASSWORD",
        "LONG_USERNAME",
        "CONFLICT_PROFILE",
        "CONFLICT_ACCOUNT",
        "LONG_NOTE",
        "LONG_CONCATENATED_NOTE",
        "VALID",
      ]);
      x["url"] = A.load.Ref(ptr + 4, undefined);
      x["username"] = A.load.Ref(ptr + 8, undefined);
      x["password"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 20)) {
        x["id"] = A.load.Int32(ptr + 16);
      } else {
        delete x["id"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ImportResultsStatus": (ref: heap.Ref<string>): number => {
      const idx = [
        "UNKNOWN_ERROR",
        "SUCCESS",
        "IO_ERROR",
        "BAD_FORMAT",
        "DISMISSED",
        "MAX_FILE_SIZE",
        "IMPORT_ALREADY_ACTIVE",
        "NUM_PASSWORDS_EXCEEDED",
        "CONFLICTS",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ImportResults": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Enum(
          ptr + 0,
          [
            "UNKNOWN_ERROR",
            "SUCCESS",
            "IO_ERROR",
            "BAD_FORMAT",
            "DISMISSED",
            "MAX_FILE_SIZE",
            "IMPORT_ALREADY_ACTIVE",
            "NUM_PASSWORDS_EXCEEDED",
            "CONFLICTS",
          ].indexOf(x["status"] as string)
        );
        A.store.Bool(ptr + 16, "numberImported" in x ? true : false);
        A.store.Int32(ptr + 4, x["numberImported"] === undefined ? 0 : (x["numberImported"] as number));
        A.store.Ref(ptr + 8, x["displayedEntries"]);
        A.store.Ref(ptr + 12, x["fileName"]);
      }
    },
    "load_ImportResults": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["status"] = A.load.Enum(ptr + 0, [
        "UNKNOWN_ERROR",
        "SUCCESS",
        "IO_ERROR",
        "BAD_FORMAT",
        "DISMISSED",
        "MAX_FILE_SIZE",
        "IMPORT_ALREADY_ACTIVE",
        "NUM_PASSWORDS_EXCEEDED",
        "CONFLICTS",
      ]);
      if (A.load.Bool(ptr + 16)) {
        x["numberImported"] = A.load.Int32(ptr + 4);
      } else {
        delete x["numberImported"];
      }
      x["displayedEntries"] = A.load.Ref(ptr + 8, undefined);
      x["fileName"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PasswordCheckState": (ref: heap.Ref<string>): number => {
      const idx = [
        "IDLE",
        "RUNNING",
        "CANCELED",
        "OFFLINE",
        "SIGNED_OUT",
        "NO_PASSWORDS",
        "QUOTA_LIMIT",
        "OTHER_ERROR",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PasswordCheckStatus": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 23, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 20, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 21, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 22, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 23, true);
        A.store.Enum(
          ptr + 0,
          [
            "IDLE",
            "RUNNING",
            "CANCELED",
            "OFFLINE",
            "SIGNED_OUT",
            "NO_PASSWORDS",
            "QUOTA_LIMIT",
            "OTHER_ERROR",
          ].indexOf(x["state"] as string)
        );
        A.store.Bool(ptr + 20, "totalNumberOfPasswords" in x ? true : false);
        A.store.Int32(ptr + 4, x["totalNumberOfPasswords"] === undefined ? 0 : (x["totalNumberOfPasswords"] as number));
        A.store.Bool(ptr + 21, "alreadyProcessed" in x ? true : false);
        A.store.Int32(ptr + 8, x["alreadyProcessed"] === undefined ? 0 : (x["alreadyProcessed"] as number));
        A.store.Bool(ptr + 22, "remainingInQueue" in x ? true : false);
        A.store.Int32(ptr + 12, x["remainingInQueue"] === undefined ? 0 : (x["remainingInQueue"] as number));
        A.store.Ref(ptr + 16, x["elapsedTimeSinceLastCheck"]);
      }
    },
    "load_PasswordCheckStatus": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["state"] = A.load.Enum(ptr + 0, [
        "IDLE",
        "RUNNING",
        "CANCELED",
        "OFFLINE",
        "SIGNED_OUT",
        "NO_PASSWORDS",
        "QUOTA_LIMIT",
        "OTHER_ERROR",
      ]);
      if (A.load.Bool(ptr + 20)) {
        x["totalNumberOfPasswords"] = A.load.Int32(ptr + 4);
      } else {
        delete x["totalNumberOfPasswords"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["alreadyProcessed"] = A.load.Int32(ptr + 8);
      } else {
        delete x["alreadyProcessed"];
      }
      if (A.load.Bool(ptr + 22)) {
        x["remainingInQueue"] = A.load.Int32(ptr + 12);
      } else {
        delete x["remainingInQueue"];
      }
      x["elapsedTimeSinceLastCheck"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PasswordExportProgress": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Enum(
          ptr + 0,
          ["NOT_STARTED", "IN_PROGRESS", "SUCCEEDED", "FAILED_CANCELLED", "FAILED_WRITE_FAILED"].indexOf(
            x["status"] as string
          )
        );
        A.store.Ref(ptr + 4, x["filePath"]);
        A.store.Ref(ptr + 8, x["folderName"]);
      }
    },
    "load_PasswordExportProgress": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["status"] = A.load.Enum(ptr + 0, [
        "NOT_STARTED",
        "IN_PROGRESS",
        "SUCCEEDED",
        "FAILED_CANCELLED",
        "FAILED_WRITE_FAILED",
      ]);
      x["filePath"] = A.load.Ref(ptr + 4, undefined);
      x["folderName"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PlaintextReason": (ref: heap.Ref<string>): number => {
      const idx = ["VIEW", "COPY", "EDIT"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_AddPassword": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "addPassword" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddPassword": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.addPassword);
    },
    "call_AddPassword": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["url"] = A.load.Ref(options + 0, undefined);
      options_ffi["username"] = A.load.Ref(options + 4, undefined);
      options_ffi["password"] = A.load.Ref(options + 8, undefined);
      options_ffi["note"] = A.load.Ref(options + 12, undefined);
      if (A.load.Bool(options + 17)) {
        options_ffi["useAccountStore"] = A.load.Bool(options + 16);
      }

      const _ret = WEBEXT.passwordsPrivate.addPassword(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_AddPassword": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["url"] = A.load.Ref(options + 0, undefined);
        options_ffi["username"] = A.load.Ref(options + 4, undefined);
        options_ffi["password"] = A.load.Ref(options + 8, undefined);
        options_ffi["note"] = A.load.Ref(options + 12, undefined);
        if (A.load.Bool(options + 17)) {
          options_ffi["useAccountStore"] = A.load.Bool(options + 16);
        }

        const _ret = WEBEXT.passwordsPrivate.addPassword(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ChangeCredential": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "changeCredential" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ChangeCredential": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.changeCredential);
    },
    "call_ChangeCredential": (retPtr: Pointer, credential: Pointer): void => {
      const credential_ffi = {};

      credential_ffi["affiliatedDomains"] = A.load.Ref(credential + 0, undefined);
      credential_ffi["username"] = A.load.Ref(credential + 4, undefined);
      credential_ffi["displayName"] = A.load.Ref(credential + 8, undefined);
      credential_ffi["password"] = A.load.Ref(credential + 12, undefined);
      credential_ffi["federationText"] = A.load.Ref(credential + 16, undefined);
      if (A.load.Bool(credential + 60)) {
        credential_ffi["id"] = A.load.Int32(credential + 20);
      }
      credential_ffi["storedIn"] = A.load.Enum(credential + 24, ["DEVICE", "ACCOUNT", "DEVICE_AND_ACCOUNT"]);
      if (A.load.Bool(credential + 61)) {
        credential_ffi["isPasskey"] = A.load.Bool(credential + 28);
      }
      credential_ffi["note"] = A.load.Ref(credential + 32, undefined);
      credential_ffi["changePasswordUrl"] = A.load.Ref(credential + 36, undefined);
      if (A.load.Bool(credential + 40 + 19)) {
        credential_ffi["compromisedInfo"] = {};
        if (A.load.Bool(credential + 40 + 17)) {
          credential_ffi["compromisedInfo"]["compromiseTime"] = A.load.Float64(credential + 40 + 0);
        }
        credential_ffi["compromisedInfo"]["elapsedTimeSinceCompromise"] = A.load.Ref(credential + 40 + 8, undefined);
        credential_ffi["compromisedInfo"]["compromiseTypes"] = A.load.Ref(credential + 40 + 12, undefined);
        if (A.load.Bool(credential + 40 + 18)) {
          credential_ffi["compromisedInfo"]["isMuted"] = A.load.Bool(credential + 40 + 16);
        }
      }

      const _ret = WEBEXT.passwordsPrivate.changeCredential(credential_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ChangeCredential": (retPtr: Pointer, errPtr: Pointer, credential: Pointer): heap.Ref<boolean> => {
      try {
        const credential_ffi = {};

        credential_ffi["affiliatedDomains"] = A.load.Ref(credential + 0, undefined);
        credential_ffi["username"] = A.load.Ref(credential + 4, undefined);
        credential_ffi["displayName"] = A.load.Ref(credential + 8, undefined);
        credential_ffi["password"] = A.load.Ref(credential + 12, undefined);
        credential_ffi["federationText"] = A.load.Ref(credential + 16, undefined);
        if (A.load.Bool(credential + 60)) {
          credential_ffi["id"] = A.load.Int32(credential + 20);
        }
        credential_ffi["storedIn"] = A.load.Enum(credential + 24, ["DEVICE", "ACCOUNT", "DEVICE_AND_ACCOUNT"]);
        if (A.load.Bool(credential + 61)) {
          credential_ffi["isPasskey"] = A.load.Bool(credential + 28);
        }
        credential_ffi["note"] = A.load.Ref(credential + 32, undefined);
        credential_ffi["changePasswordUrl"] = A.load.Ref(credential + 36, undefined);
        if (A.load.Bool(credential + 40 + 19)) {
          credential_ffi["compromisedInfo"] = {};
          if (A.load.Bool(credential + 40 + 17)) {
            credential_ffi["compromisedInfo"]["compromiseTime"] = A.load.Float64(credential + 40 + 0);
          }
          credential_ffi["compromisedInfo"]["elapsedTimeSinceCompromise"] = A.load.Ref(credential + 40 + 8, undefined);
          credential_ffi["compromisedInfo"]["compromiseTypes"] = A.load.Ref(credential + 40 + 12, undefined);
          if (A.load.Bool(credential + 40 + 18)) {
            credential_ffi["compromisedInfo"]["isMuted"] = A.load.Bool(credential + 40 + 16);
          }
        }

        const _ret = WEBEXT.passwordsPrivate.changeCredential(credential_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ContinueImport": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "continueImport" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ContinueImport": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.continueImport);
    },
    "call_ContinueImport": (retPtr: Pointer, selectedIds: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.continueImport(A.H.get<object>(selectedIds));
      A.store.Ref(retPtr, _ret);
    },
    "try_ContinueImport": (retPtr: Pointer, errPtr: Pointer, selectedIds: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.continueImport(A.H.get<object>(selectedIds));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ExportPasswords": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "exportPasswords" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ExportPasswords": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.exportPasswords);
    },
    "call_ExportPasswords": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.exportPasswords();
      A.store.Ref(retPtr, _ret);
    },
    "try_ExportPasswords": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.exportPasswords();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ExtendAuthValidity": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "extendAuthValidity" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ExtendAuthValidity": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.extendAuthValidity);
    },
    "call_ExtendAuthValidity": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.extendAuthValidity();
      A.store.Ref(retPtr, _ret);
    },
    "try_ExtendAuthValidity": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.extendAuthValidity();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_FetchFamilyMembers": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "fetchFamilyMembers" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_FetchFamilyMembers": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.fetchFamilyMembers);
    },
    "call_FetchFamilyMembers": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.fetchFamilyMembers();
      A.store.Ref(retPtr, _ret);
    },
    "try_FetchFamilyMembers": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.fetchFamilyMembers();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCredentialGroups": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "getCredentialGroups" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCredentialGroups": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.getCredentialGroups);
    },
    "call_GetCredentialGroups": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.getCredentialGroups();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCredentialGroups": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.getCredentialGroups();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCredentialsWithReusedPassword": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "getCredentialsWithReusedPassword" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCredentialsWithReusedPassword": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.getCredentialsWithReusedPassword);
    },
    "call_GetCredentialsWithReusedPassword": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.getCredentialsWithReusedPassword();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCredentialsWithReusedPassword": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.getCredentialsWithReusedPassword();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetInsecureCredentials": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "getInsecureCredentials" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetInsecureCredentials": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.getInsecureCredentials);
    },
    "call_GetInsecureCredentials": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.getInsecureCredentials();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetInsecureCredentials": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.getInsecureCredentials();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPasswordCheckStatus": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "getPasswordCheckStatus" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPasswordCheckStatus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.getPasswordCheckStatus);
    },
    "call_GetPasswordCheckStatus": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.getPasswordCheckStatus();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPasswordCheckStatus": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.getPasswordCheckStatus();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPasswordExceptionList": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "getPasswordExceptionList" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPasswordExceptionList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.getPasswordExceptionList);
    },
    "call_GetPasswordExceptionList": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.getPasswordExceptionList();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPasswordExceptionList": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.getPasswordExceptionList();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSavedPasswordList": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "getSavedPasswordList" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSavedPasswordList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.getSavedPasswordList);
    },
    "call_GetSavedPasswordList": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.getSavedPasswordList();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSavedPasswordList": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.getSavedPasswordList();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetUrlCollection": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "getUrlCollection" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetUrlCollection": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.getUrlCollection);
    },
    "call_GetUrlCollection": (retPtr: Pointer, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.getUrlCollection(A.H.get<object>(url));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetUrlCollection": (retPtr: Pointer, errPtr: Pointer, url: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.getUrlCollection(A.H.get<object>(url));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ImportPasswords": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "importPasswords" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ImportPasswords": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.importPasswords);
    },
    "call_ImportPasswords": (retPtr: Pointer, toStore: number): void => {
      const _ret = WEBEXT.passwordsPrivate.importPasswords(
        toStore > 0 && toStore <= 3 ? ["DEVICE", "ACCOUNT", "DEVICE_AND_ACCOUNT"][toStore - 1] : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_ImportPasswords": (retPtr: Pointer, errPtr: Pointer, toStore: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.importPasswords(
          toStore > 0 && toStore <= 3 ? ["DEVICE", "ACCOUNT", "DEVICE_AND_ACCOUNT"][toStore - 1] : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsAccountStoreDefault": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "isAccountStoreDefault" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsAccountStoreDefault": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.isAccountStoreDefault);
    },
    "call_IsAccountStoreDefault": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.isAccountStoreDefault();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsAccountStoreDefault": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.isAccountStoreDefault();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsOptedInForAccountStorage": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "isOptedInForAccountStorage" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsOptedInForAccountStorage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.isOptedInForAccountStorage);
    },
    "call_IsOptedInForAccountStorage": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.isOptedInForAccountStorage();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsOptedInForAccountStorage": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.isOptedInForAccountStorage();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_MovePasswordsToAccount": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "movePasswordsToAccount" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MovePasswordsToAccount": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.movePasswordsToAccount);
    },
    "call_MovePasswordsToAccount": (retPtr: Pointer, ids: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.movePasswordsToAccount(A.H.get<object>(ids));
    },
    "try_MovePasswordsToAccount": (retPtr: Pointer, errPtr: Pointer, ids: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.movePasswordsToAccount(A.H.get<object>(ids));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_MuteInsecureCredential": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "muteInsecureCredential" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MuteInsecureCredential": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.muteInsecureCredential);
    },
    "call_MuteInsecureCredential": (retPtr: Pointer, credential: Pointer): void => {
      const credential_ffi = {};

      credential_ffi["affiliatedDomains"] = A.load.Ref(credential + 0, undefined);
      credential_ffi["username"] = A.load.Ref(credential + 4, undefined);
      credential_ffi["displayName"] = A.load.Ref(credential + 8, undefined);
      credential_ffi["password"] = A.load.Ref(credential + 12, undefined);
      credential_ffi["federationText"] = A.load.Ref(credential + 16, undefined);
      if (A.load.Bool(credential + 60)) {
        credential_ffi["id"] = A.load.Int32(credential + 20);
      }
      credential_ffi["storedIn"] = A.load.Enum(credential + 24, ["DEVICE", "ACCOUNT", "DEVICE_AND_ACCOUNT"]);
      if (A.load.Bool(credential + 61)) {
        credential_ffi["isPasskey"] = A.load.Bool(credential + 28);
      }
      credential_ffi["note"] = A.load.Ref(credential + 32, undefined);
      credential_ffi["changePasswordUrl"] = A.load.Ref(credential + 36, undefined);
      if (A.load.Bool(credential + 40 + 19)) {
        credential_ffi["compromisedInfo"] = {};
        if (A.load.Bool(credential + 40 + 17)) {
          credential_ffi["compromisedInfo"]["compromiseTime"] = A.load.Float64(credential + 40 + 0);
        }
        credential_ffi["compromisedInfo"]["elapsedTimeSinceCompromise"] = A.load.Ref(credential + 40 + 8, undefined);
        credential_ffi["compromisedInfo"]["compromiseTypes"] = A.load.Ref(credential + 40 + 12, undefined);
        if (A.load.Bool(credential + 40 + 18)) {
          credential_ffi["compromisedInfo"]["isMuted"] = A.load.Bool(credential + 40 + 16);
        }
      }

      const _ret = WEBEXT.passwordsPrivate.muteInsecureCredential(credential_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_MuteInsecureCredential": (retPtr: Pointer, errPtr: Pointer, credential: Pointer): heap.Ref<boolean> => {
      try {
        const credential_ffi = {};

        credential_ffi["affiliatedDomains"] = A.load.Ref(credential + 0, undefined);
        credential_ffi["username"] = A.load.Ref(credential + 4, undefined);
        credential_ffi["displayName"] = A.load.Ref(credential + 8, undefined);
        credential_ffi["password"] = A.load.Ref(credential + 12, undefined);
        credential_ffi["federationText"] = A.load.Ref(credential + 16, undefined);
        if (A.load.Bool(credential + 60)) {
          credential_ffi["id"] = A.load.Int32(credential + 20);
        }
        credential_ffi["storedIn"] = A.load.Enum(credential + 24, ["DEVICE", "ACCOUNT", "DEVICE_AND_ACCOUNT"]);
        if (A.load.Bool(credential + 61)) {
          credential_ffi["isPasskey"] = A.load.Bool(credential + 28);
        }
        credential_ffi["note"] = A.load.Ref(credential + 32, undefined);
        credential_ffi["changePasswordUrl"] = A.load.Ref(credential + 36, undefined);
        if (A.load.Bool(credential + 40 + 19)) {
          credential_ffi["compromisedInfo"] = {};
          if (A.load.Bool(credential + 40 + 17)) {
            credential_ffi["compromisedInfo"]["compromiseTime"] = A.load.Float64(credential + 40 + 0);
          }
          credential_ffi["compromisedInfo"]["elapsedTimeSinceCompromise"] = A.load.Ref(credential + 40 + 8, undefined);
          credential_ffi["compromisedInfo"]["compromiseTypes"] = A.load.Ref(credential + 40 + 12, undefined);
          if (A.load.Bool(credential + 40 + 18)) {
            credential_ffi["compromisedInfo"]["isMuted"] = A.load.Bool(credential + 40 + 16);
          }
        }

        const _ret = WEBEXT.passwordsPrivate.muteInsecureCredential(credential_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAccountStorageOptInStateChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onAccountStorageOptInStateChanged &&
        "addListener" in WEBEXT?.passwordsPrivate?.onAccountStorageOptInStateChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAccountStorageOptInStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.addListener);
    },
    "call_OnAccountStorageOptInStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnAccountStorageOptInStateChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAccountStorageOptInStateChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onAccountStorageOptInStateChanged &&
        "removeListener" in WEBEXT?.passwordsPrivate?.onAccountStorageOptInStateChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAccountStorageOptInStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.removeListener);
    },
    "call_OffAccountStorageOptInStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffAccountStorageOptInStateChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAccountStorageOptInStateChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onAccountStorageOptInStateChanged &&
        "hasListener" in WEBEXT?.passwordsPrivate?.onAccountStorageOptInStateChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAccountStorageOptInStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.hasListener);
    },
    "call_HasOnAccountStorageOptInStateChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAccountStorageOptInStateChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnInsecureCredentialsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onInsecureCredentialsChanged &&
        "addListener" in WEBEXT?.passwordsPrivate?.onInsecureCredentialsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnInsecureCredentialsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.addListener);
    },
    "call_OnInsecureCredentialsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnInsecureCredentialsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffInsecureCredentialsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onInsecureCredentialsChanged &&
        "removeListener" in WEBEXT?.passwordsPrivate?.onInsecureCredentialsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffInsecureCredentialsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.removeListener);
    },
    "call_OffInsecureCredentialsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffInsecureCredentialsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnInsecureCredentialsChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onInsecureCredentialsChanged &&
        "hasListener" in WEBEXT?.passwordsPrivate?.onInsecureCredentialsChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnInsecureCredentialsChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.hasListener);
    },
    "call_HasOnInsecureCredentialsChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnInsecureCredentialsChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPasswordCheckStatusChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onPasswordCheckStatusChanged &&
        "addListener" in WEBEXT?.passwordsPrivate?.onPasswordCheckStatusChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPasswordCheckStatusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.addListener);
    },
    "call_OnPasswordCheckStatusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnPasswordCheckStatusChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPasswordCheckStatusChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onPasswordCheckStatusChanged &&
        "removeListener" in WEBEXT?.passwordsPrivate?.onPasswordCheckStatusChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPasswordCheckStatusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.removeListener);
    },
    "call_OffPasswordCheckStatusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffPasswordCheckStatusChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPasswordCheckStatusChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onPasswordCheckStatusChanged &&
        "hasListener" in WEBEXT?.passwordsPrivate?.onPasswordCheckStatusChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPasswordCheckStatusChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.hasListener);
    },
    "call_HasOnPasswordCheckStatusChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPasswordCheckStatusChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPasswordExceptionsListChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onPasswordExceptionsListChanged &&
        "addListener" in WEBEXT?.passwordsPrivate?.onPasswordExceptionsListChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPasswordExceptionsListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.addListener);
    },
    "call_OnPasswordExceptionsListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnPasswordExceptionsListChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPasswordExceptionsListChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onPasswordExceptionsListChanged &&
        "removeListener" in WEBEXT?.passwordsPrivate?.onPasswordExceptionsListChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPasswordExceptionsListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.removeListener);
    },
    "call_OffPasswordExceptionsListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffPasswordExceptionsListChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPasswordExceptionsListChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onPasswordExceptionsListChanged &&
        "hasListener" in WEBEXT?.passwordsPrivate?.onPasswordExceptionsListChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPasswordExceptionsListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.hasListener);
    },
    "call_HasOnPasswordExceptionsListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPasswordExceptionsListChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPasswordManagerAuthTimeout": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onPasswordManagerAuthTimeout &&
        "addListener" in WEBEXT?.passwordsPrivate?.onPasswordManagerAuthTimeout
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPasswordManagerAuthTimeout": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.addListener);
    },
    "call_OnPasswordManagerAuthTimeout": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.addListener(A.H.get<object>(callback));
    },
    "try_OnPasswordManagerAuthTimeout": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPasswordManagerAuthTimeout": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onPasswordManagerAuthTimeout &&
        "removeListener" in WEBEXT?.passwordsPrivate?.onPasswordManagerAuthTimeout
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPasswordManagerAuthTimeout": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.removeListener);
    },
    "call_OffPasswordManagerAuthTimeout": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.removeListener(A.H.get<object>(callback));
    },
    "try_OffPasswordManagerAuthTimeout": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPasswordManagerAuthTimeout": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onPasswordManagerAuthTimeout &&
        "hasListener" in WEBEXT?.passwordsPrivate?.onPasswordManagerAuthTimeout
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPasswordManagerAuthTimeout": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.hasListener);
    },
    "call_HasOnPasswordManagerAuthTimeout": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPasswordManagerAuthTimeout": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPasswordsFileExportProgress": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onPasswordsFileExportProgress &&
        "addListener" in WEBEXT?.passwordsPrivate?.onPasswordsFileExportProgress
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPasswordsFileExportProgress": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.addListener);
    },
    "call_OnPasswordsFileExportProgress": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.addListener(A.H.get<object>(callback));
    },
    "try_OnPasswordsFileExportProgress": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPasswordsFileExportProgress": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onPasswordsFileExportProgress &&
        "removeListener" in WEBEXT?.passwordsPrivate?.onPasswordsFileExportProgress
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPasswordsFileExportProgress": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.removeListener);
    },
    "call_OffPasswordsFileExportProgress": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.removeListener(A.H.get<object>(callback));
    },
    "try_OffPasswordsFileExportProgress": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPasswordsFileExportProgress": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onPasswordsFileExportProgress &&
        "hasListener" in WEBEXT?.passwordsPrivate?.onPasswordsFileExportProgress
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPasswordsFileExportProgress": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.hasListener);
    },
    "call_HasOnPasswordsFileExportProgress": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPasswordsFileExportProgress": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSavedPasswordsListChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onSavedPasswordsListChanged &&
        "addListener" in WEBEXT?.passwordsPrivate?.onSavedPasswordsListChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSavedPasswordsListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.addListener);
    },
    "call_OnSavedPasswordsListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnSavedPasswordsListChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSavedPasswordsListChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onSavedPasswordsListChanged &&
        "removeListener" in WEBEXT?.passwordsPrivate?.onSavedPasswordsListChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSavedPasswordsListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.removeListener);
    },
    "call_OffSavedPasswordsListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffSavedPasswordsListChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSavedPasswordsListChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.passwordsPrivate?.onSavedPasswordsListChanged &&
        "hasListener" in WEBEXT?.passwordsPrivate?.onSavedPasswordsListChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSavedPasswordsListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.hasListener);
    },
    "call_HasOnSavedPasswordsListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSavedPasswordsListChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OptInForAccountStorage": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "optInForAccountStorage" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OptInForAccountStorage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.optInForAccountStorage);
    },
    "call_OptInForAccountStorage": (retPtr: Pointer, optIn: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.passwordsPrivate.optInForAccountStorage(optIn === A.H.TRUE);
    },
    "try_OptInForAccountStorage": (retPtr: Pointer, errPtr: Pointer, optIn: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.optInForAccountStorage(optIn === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecordPasswordsPageAccessInSettings": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "recordPasswordsPageAccessInSettings" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecordPasswordsPageAccessInSettings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.recordPasswordsPageAccessInSettings);
    },
    "call_RecordPasswordsPageAccessInSettings": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.recordPasswordsPageAccessInSettings();
    },
    "try_RecordPasswordsPageAccessInSettings": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.recordPasswordsPageAccessInSettings();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveCredential": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "removeCredential" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveCredential": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.removeCredential);
    },
    "call_RemoveCredential": (retPtr: Pointer, id: number, fromStores: number): void => {
      const _ret = WEBEXT.passwordsPrivate.removeCredential(
        id,
        fromStores > 0 && fromStores <= 3 ? ["DEVICE", "ACCOUNT", "DEVICE_AND_ACCOUNT"][fromStores - 1] : undefined
      );
    },
    "try_RemoveCredential": (retPtr: Pointer, errPtr: Pointer, id: number, fromStores: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.removeCredential(
          id,
          fromStores > 0 && fromStores <= 3 ? ["DEVICE", "ACCOUNT", "DEVICE_AND_ACCOUNT"][fromStores - 1] : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemovePasswordException": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "removePasswordException" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemovePasswordException": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.removePasswordException);
    },
    "call_RemovePasswordException": (retPtr: Pointer, id: number): void => {
      const _ret = WEBEXT.passwordsPrivate.removePasswordException(id);
    },
    "try_RemovePasswordException": (retPtr: Pointer, errPtr: Pointer, id: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.removePasswordException(id);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RequestCredentialsDetails": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "requestCredentialsDetails" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RequestCredentialsDetails": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.requestCredentialsDetails);
    },
    "call_RequestCredentialsDetails": (retPtr: Pointer, ids: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.requestCredentialsDetails(A.H.get<object>(ids));
      A.store.Ref(retPtr, _ret);
    },
    "try_RequestCredentialsDetails": (retPtr: Pointer, errPtr: Pointer, ids: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.requestCredentialsDetails(A.H.get<object>(ids));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RequestExportProgressStatus": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "requestExportProgressStatus" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RequestExportProgressStatus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.requestExportProgressStatus);
    },
    "call_RequestExportProgressStatus": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.requestExportProgressStatus();
      A.store.Ref(retPtr, _ret);
    },
    "try_RequestExportProgressStatus": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.requestExportProgressStatus();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RequestPlaintextPassword": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "requestPlaintextPassword" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RequestPlaintextPassword": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.requestPlaintextPassword);
    },
    "call_RequestPlaintextPassword": (retPtr: Pointer, id: number, reason: number): void => {
      const _ret = WEBEXT.passwordsPrivate.requestPlaintextPassword(
        id,
        reason > 0 && reason <= 3 ? ["VIEW", "COPY", "EDIT"][reason - 1] : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_RequestPlaintextPassword": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: number,
      reason: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.requestPlaintextPassword(
          id,
          reason > 0 && reason <= 3 ? ["VIEW", "COPY", "EDIT"][reason - 1] : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ResetImporter": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "resetImporter" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ResetImporter": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.resetImporter);
    },
    "call_ResetImporter": (retPtr: Pointer, deleteFile: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.passwordsPrivate.resetImporter(deleteFile === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_ResetImporter": (retPtr: Pointer, errPtr: Pointer, deleteFile: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.resetImporter(deleteFile === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SharePassword": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "sharePassword" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SharePassword": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.sharePassword);
    },
    "call_SharePassword": (retPtr: Pointer, id: number, recipients: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.sharePassword(id, A.H.get<object>(recipients));
      A.store.Ref(retPtr, _ret);
    },
    "try_SharePassword": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: number,
      recipients: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.sharePassword(id, A.H.get<object>(recipients));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ShowAddShortcutDialog": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "showAddShortcutDialog" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ShowAddShortcutDialog": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.showAddShortcutDialog);
    },
    "call_ShowAddShortcutDialog": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.showAddShortcutDialog();
    },
    "try_ShowAddShortcutDialog": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.showAddShortcutDialog();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ShowExportedFileInShell": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "showExportedFileInShell" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ShowExportedFileInShell": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.showExportedFileInShell);
    },
    "call_ShowExportedFileInShell": (retPtr: Pointer, file_path: heap.Ref<object>): void => {
      const _ret = WEBEXT.passwordsPrivate.showExportedFileInShell(A.H.get<object>(file_path));
    },
    "try_ShowExportedFileInShell": (
      retPtr: Pointer,
      errPtr: Pointer,
      file_path: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.showExportedFileInShell(A.H.get<object>(file_path));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartPasswordCheck": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "startPasswordCheck" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartPasswordCheck": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.startPasswordCheck);
    },
    "call_StartPasswordCheck": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.startPasswordCheck();
      A.store.Ref(retPtr, _ret);
    },
    "try_StartPasswordCheck": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.startPasswordCheck();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SwitchBiometricAuthBeforeFillingState": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "switchBiometricAuthBeforeFillingState" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SwitchBiometricAuthBeforeFillingState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.switchBiometricAuthBeforeFillingState);
    },
    "call_SwitchBiometricAuthBeforeFillingState": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.switchBiometricAuthBeforeFillingState();
    },
    "try_SwitchBiometricAuthBeforeFillingState": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.switchBiometricAuthBeforeFillingState();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UndoRemoveSavedPasswordOrException": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "undoRemoveSavedPasswordOrException" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UndoRemoveSavedPasswordOrException": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.undoRemoveSavedPasswordOrException);
    },
    "call_UndoRemoveSavedPasswordOrException": (retPtr: Pointer): void => {
      const _ret = WEBEXT.passwordsPrivate.undoRemoveSavedPasswordOrException();
    },
    "try_UndoRemoveSavedPasswordOrException": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.passwordsPrivate.undoRemoveSavedPasswordOrException();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UnmuteInsecureCredential": (): heap.Ref<boolean> => {
      if (WEBEXT?.passwordsPrivate && "unmuteInsecureCredential" in WEBEXT?.passwordsPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UnmuteInsecureCredential": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.passwordsPrivate.unmuteInsecureCredential);
    },
    "call_UnmuteInsecureCredential": (retPtr: Pointer, credential: Pointer): void => {
      const credential_ffi = {};

      credential_ffi["affiliatedDomains"] = A.load.Ref(credential + 0, undefined);
      credential_ffi["username"] = A.load.Ref(credential + 4, undefined);
      credential_ffi["displayName"] = A.load.Ref(credential + 8, undefined);
      credential_ffi["password"] = A.load.Ref(credential + 12, undefined);
      credential_ffi["federationText"] = A.load.Ref(credential + 16, undefined);
      if (A.load.Bool(credential + 60)) {
        credential_ffi["id"] = A.load.Int32(credential + 20);
      }
      credential_ffi["storedIn"] = A.load.Enum(credential + 24, ["DEVICE", "ACCOUNT", "DEVICE_AND_ACCOUNT"]);
      if (A.load.Bool(credential + 61)) {
        credential_ffi["isPasskey"] = A.load.Bool(credential + 28);
      }
      credential_ffi["note"] = A.load.Ref(credential + 32, undefined);
      credential_ffi["changePasswordUrl"] = A.load.Ref(credential + 36, undefined);
      if (A.load.Bool(credential + 40 + 19)) {
        credential_ffi["compromisedInfo"] = {};
        if (A.load.Bool(credential + 40 + 17)) {
          credential_ffi["compromisedInfo"]["compromiseTime"] = A.load.Float64(credential + 40 + 0);
        }
        credential_ffi["compromisedInfo"]["elapsedTimeSinceCompromise"] = A.load.Ref(credential + 40 + 8, undefined);
        credential_ffi["compromisedInfo"]["compromiseTypes"] = A.load.Ref(credential + 40 + 12, undefined);
        if (A.load.Bool(credential + 40 + 18)) {
          credential_ffi["compromisedInfo"]["isMuted"] = A.load.Bool(credential + 40 + 16);
        }
      }

      const _ret = WEBEXT.passwordsPrivate.unmuteInsecureCredential(credential_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_UnmuteInsecureCredential": (retPtr: Pointer, errPtr: Pointer, credential: Pointer): heap.Ref<boolean> => {
      try {
        const credential_ffi = {};

        credential_ffi["affiliatedDomains"] = A.load.Ref(credential + 0, undefined);
        credential_ffi["username"] = A.load.Ref(credential + 4, undefined);
        credential_ffi["displayName"] = A.load.Ref(credential + 8, undefined);
        credential_ffi["password"] = A.load.Ref(credential + 12, undefined);
        credential_ffi["federationText"] = A.load.Ref(credential + 16, undefined);
        if (A.load.Bool(credential + 60)) {
          credential_ffi["id"] = A.load.Int32(credential + 20);
        }
        credential_ffi["storedIn"] = A.load.Enum(credential + 24, ["DEVICE", "ACCOUNT", "DEVICE_AND_ACCOUNT"]);
        if (A.load.Bool(credential + 61)) {
          credential_ffi["isPasskey"] = A.load.Bool(credential + 28);
        }
        credential_ffi["note"] = A.load.Ref(credential + 32, undefined);
        credential_ffi["changePasswordUrl"] = A.load.Ref(credential + 36, undefined);
        if (A.load.Bool(credential + 40 + 19)) {
          credential_ffi["compromisedInfo"] = {};
          if (A.load.Bool(credential + 40 + 17)) {
            credential_ffi["compromisedInfo"]["compromiseTime"] = A.load.Float64(credential + 40 + 0);
          }
          credential_ffi["compromisedInfo"]["elapsedTimeSinceCompromise"] = A.load.Ref(credential + 40 + 8, undefined);
          credential_ffi["compromisedInfo"]["compromiseTypes"] = A.load.Ref(credential + 40 + 12, undefined);
          if (A.load.Bool(credential + 40 + 18)) {
            credential_ffi["compromisedInfo"]["isMuted"] = A.load.Bool(credential + 40 + 16);
          }
        }

        const _ret = WEBEXT.passwordsPrivate.unmuteInsecureCredential(credential_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
