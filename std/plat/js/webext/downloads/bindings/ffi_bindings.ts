import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/downloads", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_BooleanDelta": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 1, false);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Bool(ptr + 2, "previous" in x ? true : false);
        A.store.Bool(ptr + 0, x["previous"] ? true : false);
        A.store.Bool(ptr + 3, "current" in x ? true : false);
        A.store.Bool(ptr + 1, x["current"] ? true : false);
      }
    },
    "load_BooleanDelta": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 2)) {
        x["previous"] = A.load.Bool(ptr + 0);
      } else {
        delete x["previous"];
      }
      if (A.load.Bool(ptr + 3)) {
        x["current"] = A.load.Bool(ptr + 1);
      } else {
        delete x["current"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DangerType": (ref: heap.Ref<string>): number => {
      const idx = [
        "file",
        "url",
        "content",
        "uncommon",
        "host",
        "unwanted",
        "safe",
        "accepted",
        "allowlistedByPolicy",
        "asyncScanning",
        "passwordProtected",
        "blockedTooLarge",
        "sensitiveContentWarning",
        "sensitiveContentBlock",
        "unsupportedFileType",
        "deepScannedFailed",
        "deepScannedSafe",
        "deepScannedOpenedDangerous",
        "promptForScaning",
        "accountCompromise",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DoubleDelta": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Bool(ptr + 16, "previous" in x ? true : false);
        A.store.Float64(ptr + 0, x["previous"] === undefined ? 0 : (x["previous"] as number));
        A.store.Bool(ptr + 17, "current" in x ? true : false);
        A.store.Float64(ptr + 8, x["current"] === undefined ? 0 : (x["current"] as number));
      }
    },
    "load_DoubleDelta": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["previous"] = A.load.Float64(ptr + 0);
      } else {
        delete x["previous"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["current"] = A.load.Float64(ptr + 8);
      } else {
        delete x["current"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_StringDelta": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["previous"]);
        A.store.Ref(ptr + 4, x["current"]);
      }
    },
    "load_StringDelta": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["previous"] = A.load.Ref(ptr + 0, undefined);
      x["current"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DownloadDelta": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 169, false);
        A.store.Bool(ptr + 168, false);
        A.store.Int32(ptr + 0, 0);

        A.store.Bool(ptr + 4 + 8, false);
        A.store.Ref(ptr + 4 + 0, undefined);
        A.store.Ref(ptr + 4 + 4, undefined);

        A.store.Bool(ptr + 16 + 8, false);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 4, undefined);

        A.store.Bool(ptr + 28 + 8, false);
        A.store.Ref(ptr + 28 + 0, undefined);
        A.store.Ref(ptr + 28 + 4, undefined);

        A.store.Bool(ptr + 40 + 8, false);
        A.store.Ref(ptr + 40 + 0, undefined);
        A.store.Ref(ptr + 40 + 4, undefined);

        A.store.Bool(ptr + 52 + 8, false);
        A.store.Ref(ptr + 52 + 0, undefined);
        A.store.Ref(ptr + 52 + 4, undefined);

        A.store.Bool(ptr + 64 + 8, false);
        A.store.Ref(ptr + 64 + 0, undefined);
        A.store.Ref(ptr + 64 + 4, undefined);

        A.store.Bool(ptr + 76 + 8, false);
        A.store.Ref(ptr + 76 + 0, undefined);
        A.store.Ref(ptr + 76 + 4, undefined);

        A.store.Bool(ptr + 88 + 8, false);
        A.store.Ref(ptr + 88 + 0, undefined);
        A.store.Ref(ptr + 88 + 4, undefined);

        A.store.Bool(ptr + 97 + 4, false);
        A.store.Bool(ptr + 97 + 2, false);
        A.store.Bool(ptr + 97 + 0, false);
        A.store.Bool(ptr + 97 + 3, false);
        A.store.Bool(ptr + 97 + 1, false);

        A.store.Bool(ptr + 102 + 4, false);
        A.store.Bool(ptr + 102 + 2, false);
        A.store.Bool(ptr + 102 + 0, false);
        A.store.Bool(ptr + 102 + 3, false);
        A.store.Bool(ptr + 102 + 1, false);

        A.store.Bool(ptr + 108 + 8, false);
        A.store.Ref(ptr + 108 + 0, undefined);
        A.store.Ref(ptr + 108 + 4, undefined);

        A.store.Bool(ptr + 120 + 18, false);
        A.store.Bool(ptr + 120 + 16, false);
        A.store.Float64(ptr + 120 + 0, 0);
        A.store.Bool(ptr + 120 + 17, false);
        A.store.Float64(ptr + 120 + 8, 0);

        A.store.Bool(ptr + 144 + 18, false);
        A.store.Bool(ptr + 144 + 16, false);
        A.store.Float64(ptr + 144 + 0, 0);
        A.store.Bool(ptr + 144 + 17, false);
        A.store.Float64(ptr + 144 + 8, 0);

        A.store.Bool(ptr + 163 + 4, false);
        A.store.Bool(ptr + 163 + 2, false);
        A.store.Bool(ptr + 163 + 0, false);
        A.store.Bool(ptr + 163 + 3, false);
        A.store.Bool(ptr + 163 + 1, false);
      } else {
        A.store.Bool(ptr + 169, true);
        A.store.Bool(ptr + 168, "id" in x ? true : false);
        A.store.Int32(ptr + 0, x["id"] === undefined ? 0 : (x["id"] as number));

        if (typeof x["url"] === "undefined") {
          A.store.Bool(ptr + 4 + 8, false);
          A.store.Ref(ptr + 4 + 0, undefined);
          A.store.Ref(ptr + 4 + 4, undefined);
        } else {
          A.store.Bool(ptr + 4 + 8, true);
          A.store.Ref(ptr + 4 + 0, x["url"]["previous"]);
          A.store.Ref(ptr + 4 + 4, x["url"]["current"]);
        }

        if (typeof x["finalUrl"] === "undefined") {
          A.store.Bool(ptr + 16 + 8, false);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 4, undefined);
        } else {
          A.store.Bool(ptr + 16 + 8, true);
          A.store.Ref(ptr + 16 + 0, x["finalUrl"]["previous"]);
          A.store.Ref(ptr + 16 + 4, x["finalUrl"]["current"]);
        }

        if (typeof x["filename"] === "undefined") {
          A.store.Bool(ptr + 28 + 8, false);
          A.store.Ref(ptr + 28 + 0, undefined);
          A.store.Ref(ptr + 28 + 4, undefined);
        } else {
          A.store.Bool(ptr + 28 + 8, true);
          A.store.Ref(ptr + 28 + 0, x["filename"]["previous"]);
          A.store.Ref(ptr + 28 + 4, x["filename"]["current"]);
        }

        if (typeof x["danger"] === "undefined") {
          A.store.Bool(ptr + 40 + 8, false);
          A.store.Ref(ptr + 40 + 0, undefined);
          A.store.Ref(ptr + 40 + 4, undefined);
        } else {
          A.store.Bool(ptr + 40 + 8, true);
          A.store.Ref(ptr + 40 + 0, x["danger"]["previous"]);
          A.store.Ref(ptr + 40 + 4, x["danger"]["current"]);
        }

        if (typeof x["mime"] === "undefined") {
          A.store.Bool(ptr + 52 + 8, false);
          A.store.Ref(ptr + 52 + 0, undefined);
          A.store.Ref(ptr + 52 + 4, undefined);
        } else {
          A.store.Bool(ptr + 52 + 8, true);
          A.store.Ref(ptr + 52 + 0, x["mime"]["previous"]);
          A.store.Ref(ptr + 52 + 4, x["mime"]["current"]);
        }

        if (typeof x["startTime"] === "undefined") {
          A.store.Bool(ptr + 64 + 8, false);
          A.store.Ref(ptr + 64 + 0, undefined);
          A.store.Ref(ptr + 64 + 4, undefined);
        } else {
          A.store.Bool(ptr + 64 + 8, true);
          A.store.Ref(ptr + 64 + 0, x["startTime"]["previous"]);
          A.store.Ref(ptr + 64 + 4, x["startTime"]["current"]);
        }

        if (typeof x["endTime"] === "undefined") {
          A.store.Bool(ptr + 76 + 8, false);
          A.store.Ref(ptr + 76 + 0, undefined);
          A.store.Ref(ptr + 76 + 4, undefined);
        } else {
          A.store.Bool(ptr + 76 + 8, true);
          A.store.Ref(ptr + 76 + 0, x["endTime"]["previous"]);
          A.store.Ref(ptr + 76 + 4, x["endTime"]["current"]);
        }

        if (typeof x["state"] === "undefined") {
          A.store.Bool(ptr + 88 + 8, false);
          A.store.Ref(ptr + 88 + 0, undefined);
          A.store.Ref(ptr + 88 + 4, undefined);
        } else {
          A.store.Bool(ptr + 88 + 8, true);
          A.store.Ref(ptr + 88 + 0, x["state"]["previous"]);
          A.store.Ref(ptr + 88 + 4, x["state"]["current"]);
        }

        if (typeof x["canResume"] === "undefined") {
          A.store.Bool(ptr + 97 + 4, false);
          A.store.Bool(ptr + 97 + 2, false);
          A.store.Bool(ptr + 97 + 0, false);
          A.store.Bool(ptr + 97 + 3, false);
          A.store.Bool(ptr + 97 + 1, false);
        } else {
          A.store.Bool(ptr + 97 + 4, true);
          A.store.Bool(ptr + 97 + 2, "previous" in x["canResume"] ? true : false);
          A.store.Bool(ptr + 97 + 0, x["canResume"]["previous"] ? true : false);
          A.store.Bool(ptr + 97 + 3, "current" in x["canResume"] ? true : false);
          A.store.Bool(ptr + 97 + 1, x["canResume"]["current"] ? true : false);
        }

        if (typeof x["paused"] === "undefined") {
          A.store.Bool(ptr + 102 + 4, false);
          A.store.Bool(ptr + 102 + 2, false);
          A.store.Bool(ptr + 102 + 0, false);
          A.store.Bool(ptr + 102 + 3, false);
          A.store.Bool(ptr + 102 + 1, false);
        } else {
          A.store.Bool(ptr + 102 + 4, true);
          A.store.Bool(ptr + 102 + 2, "previous" in x["paused"] ? true : false);
          A.store.Bool(ptr + 102 + 0, x["paused"]["previous"] ? true : false);
          A.store.Bool(ptr + 102 + 3, "current" in x["paused"] ? true : false);
          A.store.Bool(ptr + 102 + 1, x["paused"]["current"] ? true : false);
        }

        if (typeof x["error"] === "undefined") {
          A.store.Bool(ptr + 108 + 8, false);
          A.store.Ref(ptr + 108 + 0, undefined);
          A.store.Ref(ptr + 108 + 4, undefined);
        } else {
          A.store.Bool(ptr + 108 + 8, true);
          A.store.Ref(ptr + 108 + 0, x["error"]["previous"]);
          A.store.Ref(ptr + 108 + 4, x["error"]["current"]);
        }

        if (typeof x["totalBytes"] === "undefined") {
          A.store.Bool(ptr + 120 + 18, false);
          A.store.Bool(ptr + 120 + 16, false);
          A.store.Float64(ptr + 120 + 0, 0);
          A.store.Bool(ptr + 120 + 17, false);
          A.store.Float64(ptr + 120 + 8, 0);
        } else {
          A.store.Bool(ptr + 120 + 18, true);
          A.store.Bool(ptr + 120 + 16, "previous" in x["totalBytes"] ? true : false);
          A.store.Float64(
            ptr + 120 + 0,
            x["totalBytes"]["previous"] === undefined ? 0 : (x["totalBytes"]["previous"] as number)
          );
          A.store.Bool(ptr + 120 + 17, "current" in x["totalBytes"] ? true : false);
          A.store.Float64(
            ptr + 120 + 8,
            x["totalBytes"]["current"] === undefined ? 0 : (x["totalBytes"]["current"] as number)
          );
        }

        if (typeof x["fileSize"] === "undefined") {
          A.store.Bool(ptr + 144 + 18, false);
          A.store.Bool(ptr + 144 + 16, false);
          A.store.Float64(ptr + 144 + 0, 0);
          A.store.Bool(ptr + 144 + 17, false);
          A.store.Float64(ptr + 144 + 8, 0);
        } else {
          A.store.Bool(ptr + 144 + 18, true);
          A.store.Bool(ptr + 144 + 16, "previous" in x["fileSize"] ? true : false);
          A.store.Float64(
            ptr + 144 + 0,
            x["fileSize"]["previous"] === undefined ? 0 : (x["fileSize"]["previous"] as number)
          );
          A.store.Bool(ptr + 144 + 17, "current" in x["fileSize"] ? true : false);
          A.store.Float64(
            ptr + 144 + 8,
            x["fileSize"]["current"] === undefined ? 0 : (x["fileSize"]["current"] as number)
          );
        }

        if (typeof x["exists"] === "undefined") {
          A.store.Bool(ptr + 163 + 4, false);
          A.store.Bool(ptr + 163 + 2, false);
          A.store.Bool(ptr + 163 + 0, false);
          A.store.Bool(ptr + 163 + 3, false);
          A.store.Bool(ptr + 163 + 1, false);
        } else {
          A.store.Bool(ptr + 163 + 4, true);
          A.store.Bool(ptr + 163 + 2, "previous" in x["exists"] ? true : false);
          A.store.Bool(ptr + 163 + 0, x["exists"]["previous"] ? true : false);
          A.store.Bool(ptr + 163 + 3, "current" in x["exists"] ? true : false);
          A.store.Bool(ptr + 163 + 1, x["exists"]["current"] ? true : false);
        }
      }
    },
    "load_DownloadDelta": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 168)) {
        x["id"] = A.load.Int32(ptr + 0);
      } else {
        delete x["id"];
      }
      if (A.load.Bool(ptr + 4 + 8)) {
        x["url"] = {};
        x["url"]["previous"] = A.load.Ref(ptr + 4 + 0, undefined);
        x["url"]["current"] = A.load.Ref(ptr + 4 + 4, undefined);
      } else {
        delete x["url"];
      }
      if (A.load.Bool(ptr + 16 + 8)) {
        x["finalUrl"] = {};
        x["finalUrl"]["previous"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["finalUrl"]["current"] = A.load.Ref(ptr + 16 + 4, undefined);
      } else {
        delete x["finalUrl"];
      }
      if (A.load.Bool(ptr + 28 + 8)) {
        x["filename"] = {};
        x["filename"]["previous"] = A.load.Ref(ptr + 28 + 0, undefined);
        x["filename"]["current"] = A.load.Ref(ptr + 28 + 4, undefined);
      } else {
        delete x["filename"];
      }
      if (A.load.Bool(ptr + 40 + 8)) {
        x["danger"] = {};
        x["danger"]["previous"] = A.load.Ref(ptr + 40 + 0, undefined);
        x["danger"]["current"] = A.load.Ref(ptr + 40 + 4, undefined);
      } else {
        delete x["danger"];
      }
      if (A.load.Bool(ptr + 52 + 8)) {
        x["mime"] = {};
        x["mime"]["previous"] = A.load.Ref(ptr + 52 + 0, undefined);
        x["mime"]["current"] = A.load.Ref(ptr + 52 + 4, undefined);
      } else {
        delete x["mime"];
      }
      if (A.load.Bool(ptr + 64 + 8)) {
        x["startTime"] = {};
        x["startTime"]["previous"] = A.load.Ref(ptr + 64 + 0, undefined);
        x["startTime"]["current"] = A.load.Ref(ptr + 64 + 4, undefined);
      } else {
        delete x["startTime"];
      }
      if (A.load.Bool(ptr + 76 + 8)) {
        x["endTime"] = {};
        x["endTime"]["previous"] = A.load.Ref(ptr + 76 + 0, undefined);
        x["endTime"]["current"] = A.load.Ref(ptr + 76 + 4, undefined);
      } else {
        delete x["endTime"];
      }
      if (A.load.Bool(ptr + 88 + 8)) {
        x["state"] = {};
        x["state"]["previous"] = A.load.Ref(ptr + 88 + 0, undefined);
        x["state"]["current"] = A.load.Ref(ptr + 88 + 4, undefined);
      } else {
        delete x["state"];
      }
      if (A.load.Bool(ptr + 97 + 4)) {
        x["canResume"] = {};
        if (A.load.Bool(ptr + 97 + 2)) {
          x["canResume"]["previous"] = A.load.Bool(ptr + 97 + 0);
        } else {
          delete x["canResume"]["previous"];
        }
        if (A.load.Bool(ptr + 97 + 3)) {
          x["canResume"]["current"] = A.load.Bool(ptr + 97 + 1);
        } else {
          delete x["canResume"]["current"];
        }
      } else {
        delete x["canResume"];
      }
      if (A.load.Bool(ptr + 102 + 4)) {
        x["paused"] = {};
        if (A.load.Bool(ptr + 102 + 2)) {
          x["paused"]["previous"] = A.load.Bool(ptr + 102 + 0);
        } else {
          delete x["paused"]["previous"];
        }
        if (A.load.Bool(ptr + 102 + 3)) {
          x["paused"]["current"] = A.load.Bool(ptr + 102 + 1);
        } else {
          delete x["paused"]["current"];
        }
      } else {
        delete x["paused"];
      }
      if (A.load.Bool(ptr + 108 + 8)) {
        x["error"] = {};
        x["error"]["previous"] = A.load.Ref(ptr + 108 + 0, undefined);
        x["error"]["current"] = A.load.Ref(ptr + 108 + 4, undefined);
      } else {
        delete x["error"];
      }
      if (A.load.Bool(ptr + 120 + 18)) {
        x["totalBytes"] = {};
        if (A.load.Bool(ptr + 120 + 16)) {
          x["totalBytes"]["previous"] = A.load.Float64(ptr + 120 + 0);
        } else {
          delete x["totalBytes"]["previous"];
        }
        if (A.load.Bool(ptr + 120 + 17)) {
          x["totalBytes"]["current"] = A.load.Float64(ptr + 120 + 8);
        } else {
          delete x["totalBytes"]["current"];
        }
      } else {
        delete x["totalBytes"];
      }
      if (A.load.Bool(ptr + 144 + 18)) {
        x["fileSize"] = {};
        if (A.load.Bool(ptr + 144 + 16)) {
          x["fileSize"]["previous"] = A.load.Float64(ptr + 144 + 0);
        } else {
          delete x["fileSize"]["previous"];
        }
        if (A.load.Bool(ptr + 144 + 17)) {
          x["fileSize"]["current"] = A.load.Float64(ptr + 144 + 8);
        } else {
          delete x["fileSize"]["current"];
        }
      } else {
        delete x["fileSize"];
      }
      if (A.load.Bool(ptr + 163 + 4)) {
        x["exists"] = {};
        if (A.load.Bool(ptr + 163 + 2)) {
          x["exists"]["previous"] = A.load.Bool(ptr + 163 + 0);
        } else {
          delete x["exists"]["previous"];
        }
        if (A.load.Bool(ptr + 163 + 3)) {
          x["exists"]["current"] = A.load.Bool(ptr + 163 + 1);
        } else {
          delete x["exists"]["current"];
        }
      } else {
        delete x["exists"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_State": (ref: heap.Ref<string>): number => {
      const idx = ["in_progress", "interrupted", "complete"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_InterruptReason": (ref: heap.Ref<string>): number => {
      const idx = [
        "FILE_FAILED",
        "FILE_ACCESS_DENIED",
        "FILE_NO_SPACE",
        "FILE_NAME_TOO_LONG",
        "FILE_TOO_LARGE",
        "FILE_VIRUS_INFECTED",
        "FILE_TRANSIENT_ERROR",
        "FILE_BLOCKED",
        "FILE_SECURITY_CHECK_FAILED",
        "FILE_TOO_SHORT",
        "FILE_HASH_MISMATCH",
        "FILE_SAME_AS_SOURCE",
        "NETWORK_FAILED",
        "NETWORK_TIMEOUT",
        "NETWORK_DISCONNECTED",
        "NETWORK_SERVER_DOWN",
        "NETWORK_INVALID_REQUEST",
        "SERVER_FAILED",
        "SERVER_NO_RANGE",
        "SERVER_BAD_CONTENT",
        "SERVER_UNAUTHORIZED",
        "SERVER_CERT_PROBLEM",
        "SERVER_FORBIDDEN",
        "SERVER_UNREACHABLE",
        "SERVER_CONTENT_LENGTH_MISMATCH",
        "SERVER_CROSS_ORIGIN_REDIRECT",
        "USER_CANCELED",
        "USER_SHUTDOWN",
        "CRASH",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DownloadItem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 100, false);
        A.store.Bool(ptr + 92, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 93, false);
        A.store.Bool(ptr + 20, false);
        A.store.Enum(ptr + 24, -1);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Ref(ptr + 40, undefined);
        A.store.Enum(ptr + 44, -1);
        A.store.Bool(ptr + 94, false);
        A.store.Bool(ptr + 48, false);
        A.store.Bool(ptr + 95, false);
        A.store.Bool(ptr + 49, false);
        A.store.Enum(ptr + 52, -1);
        A.store.Bool(ptr + 96, false);
        A.store.Float64(ptr + 56, 0);
        A.store.Bool(ptr + 97, false);
        A.store.Float64(ptr + 64, 0);
        A.store.Bool(ptr + 98, false);
        A.store.Float64(ptr + 72, 0);
        A.store.Bool(ptr + 99, false);
        A.store.Bool(ptr + 80, false);
        A.store.Ref(ptr + 84, undefined);
        A.store.Ref(ptr + 88, undefined);
      } else {
        A.store.Bool(ptr + 100, true);
        A.store.Bool(ptr + 92, "id" in x ? true : false);
        A.store.Int32(ptr + 0, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Ref(ptr + 4, x["url"]);
        A.store.Ref(ptr + 8, x["finalUrl"]);
        A.store.Ref(ptr + 12, x["referrer"]);
        A.store.Ref(ptr + 16, x["filename"]);
        A.store.Bool(ptr + 93, "incognito" in x ? true : false);
        A.store.Bool(ptr + 20, x["incognito"] ? true : false);
        A.store.Enum(
          ptr + 24,
          [
            "file",
            "url",
            "content",
            "uncommon",
            "host",
            "unwanted",
            "safe",
            "accepted",
            "allowlistedByPolicy",
            "asyncScanning",
            "passwordProtected",
            "blockedTooLarge",
            "sensitiveContentWarning",
            "sensitiveContentBlock",
            "unsupportedFileType",
            "deepScannedFailed",
            "deepScannedSafe",
            "deepScannedOpenedDangerous",
            "promptForScaning",
            "accountCompromise",
          ].indexOf(x["danger"] as string)
        );
        A.store.Ref(ptr + 28, x["mime"]);
        A.store.Ref(ptr + 32, x["startTime"]);
        A.store.Ref(ptr + 36, x["endTime"]);
        A.store.Ref(ptr + 40, x["estimatedEndTime"]);
        A.store.Enum(ptr + 44, ["in_progress", "interrupted", "complete"].indexOf(x["state"] as string));
        A.store.Bool(ptr + 94, "paused" in x ? true : false);
        A.store.Bool(ptr + 48, x["paused"] ? true : false);
        A.store.Bool(ptr + 95, "canResume" in x ? true : false);
        A.store.Bool(ptr + 49, x["canResume"] ? true : false);
        A.store.Enum(
          ptr + 52,
          [
            "FILE_FAILED",
            "FILE_ACCESS_DENIED",
            "FILE_NO_SPACE",
            "FILE_NAME_TOO_LONG",
            "FILE_TOO_LARGE",
            "FILE_VIRUS_INFECTED",
            "FILE_TRANSIENT_ERROR",
            "FILE_BLOCKED",
            "FILE_SECURITY_CHECK_FAILED",
            "FILE_TOO_SHORT",
            "FILE_HASH_MISMATCH",
            "FILE_SAME_AS_SOURCE",
            "NETWORK_FAILED",
            "NETWORK_TIMEOUT",
            "NETWORK_DISCONNECTED",
            "NETWORK_SERVER_DOWN",
            "NETWORK_INVALID_REQUEST",
            "SERVER_FAILED",
            "SERVER_NO_RANGE",
            "SERVER_BAD_CONTENT",
            "SERVER_UNAUTHORIZED",
            "SERVER_CERT_PROBLEM",
            "SERVER_FORBIDDEN",
            "SERVER_UNREACHABLE",
            "SERVER_CONTENT_LENGTH_MISMATCH",
            "SERVER_CROSS_ORIGIN_REDIRECT",
            "USER_CANCELED",
            "USER_SHUTDOWN",
            "CRASH",
          ].indexOf(x["error"] as string)
        );
        A.store.Bool(ptr + 96, "bytesReceived" in x ? true : false);
        A.store.Float64(ptr + 56, x["bytesReceived"] === undefined ? 0 : (x["bytesReceived"] as number));
        A.store.Bool(ptr + 97, "totalBytes" in x ? true : false);
        A.store.Float64(ptr + 64, x["totalBytes"] === undefined ? 0 : (x["totalBytes"] as number));
        A.store.Bool(ptr + 98, "fileSize" in x ? true : false);
        A.store.Float64(ptr + 72, x["fileSize"] === undefined ? 0 : (x["fileSize"] as number));
        A.store.Bool(ptr + 99, "exists" in x ? true : false);
        A.store.Bool(ptr + 80, x["exists"] ? true : false);
        A.store.Ref(ptr + 84, x["byExtensionId"]);
        A.store.Ref(ptr + 88, x["byExtensionName"]);
      }
    },
    "load_DownloadItem": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 92)) {
        x["id"] = A.load.Int32(ptr + 0);
      } else {
        delete x["id"];
      }
      x["url"] = A.load.Ref(ptr + 4, undefined);
      x["finalUrl"] = A.load.Ref(ptr + 8, undefined);
      x["referrer"] = A.load.Ref(ptr + 12, undefined);
      x["filename"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 93)) {
        x["incognito"] = A.load.Bool(ptr + 20);
      } else {
        delete x["incognito"];
      }
      x["danger"] = A.load.Enum(ptr + 24, [
        "file",
        "url",
        "content",
        "uncommon",
        "host",
        "unwanted",
        "safe",
        "accepted",
        "allowlistedByPolicy",
        "asyncScanning",
        "passwordProtected",
        "blockedTooLarge",
        "sensitiveContentWarning",
        "sensitiveContentBlock",
        "unsupportedFileType",
        "deepScannedFailed",
        "deepScannedSafe",
        "deepScannedOpenedDangerous",
        "promptForScaning",
        "accountCompromise",
      ]);
      x["mime"] = A.load.Ref(ptr + 28, undefined);
      x["startTime"] = A.load.Ref(ptr + 32, undefined);
      x["endTime"] = A.load.Ref(ptr + 36, undefined);
      x["estimatedEndTime"] = A.load.Ref(ptr + 40, undefined);
      x["state"] = A.load.Enum(ptr + 44, ["in_progress", "interrupted", "complete"]);
      if (A.load.Bool(ptr + 94)) {
        x["paused"] = A.load.Bool(ptr + 48);
      } else {
        delete x["paused"];
      }
      if (A.load.Bool(ptr + 95)) {
        x["canResume"] = A.load.Bool(ptr + 49);
      } else {
        delete x["canResume"];
      }
      x["error"] = A.load.Enum(ptr + 52, [
        "FILE_FAILED",
        "FILE_ACCESS_DENIED",
        "FILE_NO_SPACE",
        "FILE_NAME_TOO_LONG",
        "FILE_TOO_LARGE",
        "FILE_VIRUS_INFECTED",
        "FILE_TRANSIENT_ERROR",
        "FILE_BLOCKED",
        "FILE_SECURITY_CHECK_FAILED",
        "FILE_TOO_SHORT",
        "FILE_HASH_MISMATCH",
        "FILE_SAME_AS_SOURCE",
        "NETWORK_FAILED",
        "NETWORK_TIMEOUT",
        "NETWORK_DISCONNECTED",
        "NETWORK_SERVER_DOWN",
        "NETWORK_INVALID_REQUEST",
        "SERVER_FAILED",
        "SERVER_NO_RANGE",
        "SERVER_BAD_CONTENT",
        "SERVER_UNAUTHORIZED",
        "SERVER_CERT_PROBLEM",
        "SERVER_FORBIDDEN",
        "SERVER_UNREACHABLE",
        "SERVER_CONTENT_LENGTH_MISMATCH",
        "SERVER_CROSS_ORIGIN_REDIRECT",
        "USER_CANCELED",
        "USER_SHUTDOWN",
        "CRASH",
      ]);
      if (A.load.Bool(ptr + 96)) {
        x["bytesReceived"] = A.load.Float64(ptr + 56);
      } else {
        delete x["bytesReceived"];
      }
      if (A.load.Bool(ptr + 97)) {
        x["totalBytes"] = A.load.Float64(ptr + 64);
      } else {
        delete x["totalBytes"];
      }
      if (A.load.Bool(ptr + 98)) {
        x["fileSize"] = A.load.Float64(ptr + 72);
      } else {
        delete x["fileSize"];
      }
      if (A.load.Bool(ptr + 99)) {
        x["exists"] = A.load.Bool(ptr + 80);
      } else {
        delete x["exists"];
      }
      x["byExtensionId"] = A.load.Ref(ptr + 84, undefined);
      x["byExtensionName"] = A.load.Ref(ptr + 88, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_FilenameConflictAction": (ref: heap.Ref<string>): number => {
      const idx = ["uniquify", "overwrite", "prompt"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_HttpMethod": (ref: heap.Ref<string>): number => {
      const idx = ["GET", "POST"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_HeaderNameValuePair": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["value"]);
      }
    },
    "load_HeaderNameValuePair": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["value"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DownloadOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 29, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 12, false);
        A.store.Enum(ptr + 16, -1);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 29, true);
        A.store.Ref(ptr + 0, x["url"]);
        A.store.Ref(ptr + 4, x["filename"]);
        A.store.Enum(ptr + 8, ["uniquify", "overwrite", "prompt"].indexOf(x["conflictAction"] as string));
        A.store.Bool(ptr + 28, "saveAs" in x ? true : false);
        A.store.Bool(ptr + 12, x["saveAs"] ? true : false);
        A.store.Enum(ptr + 16, ["GET", "POST"].indexOf(x["method"] as string));
        A.store.Ref(ptr + 20, x["headers"]);
        A.store.Ref(ptr + 24, x["body"]);
      }
    },
    "load_DownloadOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["url"] = A.load.Ref(ptr + 0, undefined);
      x["filename"] = A.load.Ref(ptr + 4, undefined);
      x["conflictAction"] = A.load.Enum(ptr + 8, ["uniquify", "overwrite", "prompt"]);
      if (A.load.Bool(ptr + 28)) {
        x["saveAs"] = A.load.Bool(ptr + 12);
      } else {
        delete x["saveAs"];
      }
      x["method"] = A.load.Enum(ptr + 16, ["GET", "POST"]);
      x["headers"] = A.load.Ref(ptr + 20, undefined);
      x["body"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DownloadQuery": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 138, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 129, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Bool(ptr + 130, false);
        A.store.Float64(ptr + 32, 0);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Ref(ptr + 48, undefined);
        A.store.Bool(ptr + 131, false);
        A.store.Int32(ptr + 52, 0);
        A.store.Ref(ptr + 56, undefined);
        A.store.Bool(ptr + 132, false);
        A.store.Int32(ptr + 60, 0);
        A.store.Ref(ptr + 64, undefined);
        A.store.Ref(ptr + 68, undefined);
        A.store.Ref(ptr + 72, undefined);
        A.store.Enum(ptr + 76, -1);
        A.store.Ref(ptr + 80, undefined);
        A.store.Ref(ptr + 84, undefined);
        A.store.Ref(ptr + 88, undefined);
        A.store.Enum(ptr + 92, -1);
        A.store.Bool(ptr + 133, false);
        A.store.Bool(ptr + 96, false);
        A.store.Enum(ptr + 100, -1);
        A.store.Bool(ptr + 134, false);
        A.store.Float64(ptr + 104, 0);
        A.store.Bool(ptr + 135, false);
        A.store.Float64(ptr + 112, 0);
        A.store.Bool(ptr + 136, false);
        A.store.Float64(ptr + 120, 0);
        A.store.Bool(ptr + 137, false);
        A.store.Bool(ptr + 128, false);
      } else {
        A.store.Bool(ptr + 138, true);
        A.store.Ref(ptr + 0, x["query"]);
        A.store.Ref(ptr + 4, x["startedBefore"]);
        A.store.Ref(ptr + 8, x["startedAfter"]);
        A.store.Ref(ptr + 12, x["endedBefore"]);
        A.store.Ref(ptr + 16, x["endedAfter"]);
        A.store.Bool(ptr + 129, "totalBytesGreater" in x ? true : false);
        A.store.Float64(ptr + 24, x["totalBytesGreater"] === undefined ? 0 : (x["totalBytesGreater"] as number));
        A.store.Bool(ptr + 130, "totalBytesLess" in x ? true : false);
        A.store.Float64(ptr + 32, x["totalBytesLess"] === undefined ? 0 : (x["totalBytesLess"] as number));
        A.store.Ref(ptr + 40, x["filenameRegex"]);
        A.store.Ref(ptr + 44, x["urlRegex"]);
        A.store.Ref(ptr + 48, x["finalUrlRegex"]);
        A.store.Bool(ptr + 131, "limit" in x ? true : false);
        A.store.Int32(ptr + 52, x["limit"] === undefined ? 0 : (x["limit"] as number));
        A.store.Ref(ptr + 56, x["orderBy"]);
        A.store.Bool(ptr + 132, "id" in x ? true : false);
        A.store.Int32(ptr + 60, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Ref(ptr + 64, x["url"]);
        A.store.Ref(ptr + 68, x["finalUrl"]);
        A.store.Ref(ptr + 72, x["filename"]);
        A.store.Enum(
          ptr + 76,
          [
            "file",
            "url",
            "content",
            "uncommon",
            "host",
            "unwanted",
            "safe",
            "accepted",
            "allowlistedByPolicy",
            "asyncScanning",
            "passwordProtected",
            "blockedTooLarge",
            "sensitiveContentWarning",
            "sensitiveContentBlock",
            "unsupportedFileType",
            "deepScannedFailed",
            "deepScannedSafe",
            "deepScannedOpenedDangerous",
            "promptForScaning",
            "accountCompromise",
          ].indexOf(x["danger"] as string)
        );
        A.store.Ref(ptr + 80, x["mime"]);
        A.store.Ref(ptr + 84, x["startTime"]);
        A.store.Ref(ptr + 88, x["endTime"]);
        A.store.Enum(ptr + 92, ["in_progress", "interrupted", "complete"].indexOf(x["state"] as string));
        A.store.Bool(ptr + 133, "paused" in x ? true : false);
        A.store.Bool(ptr + 96, x["paused"] ? true : false);
        A.store.Enum(
          ptr + 100,
          [
            "FILE_FAILED",
            "FILE_ACCESS_DENIED",
            "FILE_NO_SPACE",
            "FILE_NAME_TOO_LONG",
            "FILE_TOO_LARGE",
            "FILE_VIRUS_INFECTED",
            "FILE_TRANSIENT_ERROR",
            "FILE_BLOCKED",
            "FILE_SECURITY_CHECK_FAILED",
            "FILE_TOO_SHORT",
            "FILE_HASH_MISMATCH",
            "FILE_SAME_AS_SOURCE",
            "NETWORK_FAILED",
            "NETWORK_TIMEOUT",
            "NETWORK_DISCONNECTED",
            "NETWORK_SERVER_DOWN",
            "NETWORK_INVALID_REQUEST",
            "SERVER_FAILED",
            "SERVER_NO_RANGE",
            "SERVER_BAD_CONTENT",
            "SERVER_UNAUTHORIZED",
            "SERVER_CERT_PROBLEM",
            "SERVER_FORBIDDEN",
            "SERVER_UNREACHABLE",
            "SERVER_CONTENT_LENGTH_MISMATCH",
            "SERVER_CROSS_ORIGIN_REDIRECT",
            "USER_CANCELED",
            "USER_SHUTDOWN",
            "CRASH",
          ].indexOf(x["error"] as string)
        );
        A.store.Bool(ptr + 134, "bytesReceived" in x ? true : false);
        A.store.Float64(ptr + 104, x["bytesReceived"] === undefined ? 0 : (x["bytesReceived"] as number));
        A.store.Bool(ptr + 135, "totalBytes" in x ? true : false);
        A.store.Float64(ptr + 112, x["totalBytes"] === undefined ? 0 : (x["totalBytes"] as number));
        A.store.Bool(ptr + 136, "fileSize" in x ? true : false);
        A.store.Float64(ptr + 120, x["fileSize"] === undefined ? 0 : (x["fileSize"] as number));
        A.store.Bool(ptr + 137, "exists" in x ? true : false);
        A.store.Bool(ptr + 128, x["exists"] ? true : false);
      }
    },
    "load_DownloadQuery": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["query"] = A.load.Ref(ptr + 0, undefined);
      x["startedBefore"] = A.load.Ref(ptr + 4, undefined);
      x["startedAfter"] = A.load.Ref(ptr + 8, undefined);
      x["endedBefore"] = A.load.Ref(ptr + 12, undefined);
      x["endedAfter"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 129)) {
        x["totalBytesGreater"] = A.load.Float64(ptr + 24);
      } else {
        delete x["totalBytesGreater"];
      }
      if (A.load.Bool(ptr + 130)) {
        x["totalBytesLess"] = A.load.Float64(ptr + 32);
      } else {
        delete x["totalBytesLess"];
      }
      x["filenameRegex"] = A.load.Ref(ptr + 40, undefined);
      x["urlRegex"] = A.load.Ref(ptr + 44, undefined);
      x["finalUrlRegex"] = A.load.Ref(ptr + 48, undefined);
      if (A.load.Bool(ptr + 131)) {
        x["limit"] = A.load.Int32(ptr + 52);
      } else {
        delete x["limit"];
      }
      x["orderBy"] = A.load.Ref(ptr + 56, undefined);
      if (A.load.Bool(ptr + 132)) {
        x["id"] = A.load.Int32(ptr + 60);
      } else {
        delete x["id"];
      }
      x["url"] = A.load.Ref(ptr + 64, undefined);
      x["finalUrl"] = A.load.Ref(ptr + 68, undefined);
      x["filename"] = A.load.Ref(ptr + 72, undefined);
      x["danger"] = A.load.Enum(ptr + 76, [
        "file",
        "url",
        "content",
        "uncommon",
        "host",
        "unwanted",
        "safe",
        "accepted",
        "allowlistedByPolicy",
        "asyncScanning",
        "passwordProtected",
        "blockedTooLarge",
        "sensitiveContentWarning",
        "sensitiveContentBlock",
        "unsupportedFileType",
        "deepScannedFailed",
        "deepScannedSafe",
        "deepScannedOpenedDangerous",
        "promptForScaning",
        "accountCompromise",
      ]);
      x["mime"] = A.load.Ref(ptr + 80, undefined);
      x["startTime"] = A.load.Ref(ptr + 84, undefined);
      x["endTime"] = A.load.Ref(ptr + 88, undefined);
      x["state"] = A.load.Enum(ptr + 92, ["in_progress", "interrupted", "complete"]);
      if (A.load.Bool(ptr + 133)) {
        x["paused"] = A.load.Bool(ptr + 96);
      } else {
        delete x["paused"];
      }
      x["error"] = A.load.Enum(ptr + 100, [
        "FILE_FAILED",
        "FILE_ACCESS_DENIED",
        "FILE_NO_SPACE",
        "FILE_NAME_TOO_LONG",
        "FILE_TOO_LARGE",
        "FILE_VIRUS_INFECTED",
        "FILE_TRANSIENT_ERROR",
        "FILE_BLOCKED",
        "FILE_SECURITY_CHECK_FAILED",
        "FILE_TOO_SHORT",
        "FILE_HASH_MISMATCH",
        "FILE_SAME_AS_SOURCE",
        "NETWORK_FAILED",
        "NETWORK_TIMEOUT",
        "NETWORK_DISCONNECTED",
        "NETWORK_SERVER_DOWN",
        "NETWORK_INVALID_REQUEST",
        "SERVER_FAILED",
        "SERVER_NO_RANGE",
        "SERVER_BAD_CONTENT",
        "SERVER_UNAUTHORIZED",
        "SERVER_CERT_PROBLEM",
        "SERVER_FORBIDDEN",
        "SERVER_UNREACHABLE",
        "SERVER_CONTENT_LENGTH_MISMATCH",
        "SERVER_CROSS_ORIGIN_REDIRECT",
        "USER_CANCELED",
        "USER_SHUTDOWN",
        "CRASH",
      ]);
      if (A.load.Bool(ptr + 134)) {
        x["bytesReceived"] = A.load.Float64(ptr + 104);
      } else {
        delete x["bytesReceived"];
      }
      if (A.load.Bool(ptr + 135)) {
        x["totalBytes"] = A.load.Float64(ptr + 112);
      } else {
        delete x["totalBytes"];
      }
      if (A.load.Bool(ptr + 136)) {
        x["fileSize"] = A.load.Float64(ptr + 120);
      } else {
        delete x["fileSize"];
      }
      if (A.load.Bool(ptr + 137)) {
        x["exists"] = A.load.Bool(ptr + 128);
      } else {
        delete x["exists"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_FilenameSuggestion": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["filename"]);
        A.store.Enum(ptr + 4, ["uniquify", "overwrite", "prompt"].indexOf(x["conflictAction"] as string));
      }
    },
    "load_FilenameSuggestion": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["filename"] = A.load.Ref(ptr + 0, undefined);
      x["conflictAction"] = A.load.Enum(ptr + 4, ["uniquify", "overwrite", "prompt"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetFileIconOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "size" in x ? true : false);
        A.store.Int32(ptr + 0, x["size"] === undefined ? 0 : (x["size"] as number));
      }
    },
    "load_GetFileIconOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["size"] = A.load.Int32(ptr + 0);
      } else {
        delete x["size"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UiOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 1, "enabled" in x ? true : false);
        A.store.Bool(ptr + 0, x["enabled"] ? true : false);
      }
    },
    "load_UiOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 1)) {
        x["enabled"] = A.load.Bool(ptr + 0);
      } else {
        delete x["enabled"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_AcceptDanger": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads && "acceptDanger" in WEBEXT?.downloads) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AcceptDanger": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.acceptDanger);
    },
    "call_AcceptDanger": (retPtr: Pointer, downloadId: number): void => {
      const _ret = WEBEXT.downloads.acceptDanger(downloadId);
      A.store.Ref(retPtr, _ret);
    },
    "try_AcceptDanger": (retPtr: Pointer, errPtr: Pointer, downloadId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.acceptDanger(downloadId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Cancel": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads && "cancel" in WEBEXT?.downloads) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Cancel": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.cancel);
    },
    "call_Cancel": (retPtr: Pointer, downloadId: number): void => {
      const _ret = WEBEXT.downloads.cancel(downloadId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Cancel": (retPtr: Pointer, errPtr: Pointer, downloadId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.cancel(downloadId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Download": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads && "download" in WEBEXT?.downloads) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Download": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.download);
    },
    "call_Download": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["url"] = A.load.Ref(options + 0, undefined);
      options_ffi["filename"] = A.load.Ref(options + 4, undefined);
      options_ffi["conflictAction"] = A.load.Enum(options + 8, ["uniquify", "overwrite", "prompt"]);
      if (A.load.Bool(options + 28)) {
        options_ffi["saveAs"] = A.load.Bool(options + 12);
      }
      options_ffi["method"] = A.load.Enum(options + 16, ["GET", "POST"]);
      options_ffi["headers"] = A.load.Ref(options + 20, undefined);
      options_ffi["body"] = A.load.Ref(options + 24, undefined);

      const _ret = WEBEXT.downloads.download(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Download": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["url"] = A.load.Ref(options + 0, undefined);
        options_ffi["filename"] = A.load.Ref(options + 4, undefined);
        options_ffi["conflictAction"] = A.load.Enum(options + 8, ["uniquify", "overwrite", "prompt"]);
        if (A.load.Bool(options + 28)) {
          options_ffi["saveAs"] = A.load.Bool(options + 12);
        }
        options_ffi["method"] = A.load.Enum(options + 16, ["GET", "POST"]);
        options_ffi["headers"] = A.load.Ref(options + 20, undefined);
        options_ffi["body"] = A.load.Ref(options + 24, undefined);

        const _ret = WEBEXT.downloads.download(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Erase": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads && "erase" in WEBEXT?.downloads) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Erase": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.erase);
    },
    "call_Erase": (retPtr: Pointer, query: Pointer): void => {
      const query_ffi = {};

      query_ffi["query"] = A.load.Ref(query + 0, undefined);
      query_ffi["startedBefore"] = A.load.Ref(query + 4, undefined);
      query_ffi["startedAfter"] = A.load.Ref(query + 8, undefined);
      query_ffi["endedBefore"] = A.load.Ref(query + 12, undefined);
      query_ffi["endedAfter"] = A.load.Ref(query + 16, undefined);
      if (A.load.Bool(query + 129)) {
        query_ffi["totalBytesGreater"] = A.load.Float64(query + 24);
      }
      if (A.load.Bool(query + 130)) {
        query_ffi["totalBytesLess"] = A.load.Float64(query + 32);
      }
      query_ffi["filenameRegex"] = A.load.Ref(query + 40, undefined);
      query_ffi["urlRegex"] = A.load.Ref(query + 44, undefined);
      query_ffi["finalUrlRegex"] = A.load.Ref(query + 48, undefined);
      if (A.load.Bool(query + 131)) {
        query_ffi["limit"] = A.load.Int32(query + 52);
      }
      query_ffi["orderBy"] = A.load.Ref(query + 56, undefined);
      if (A.load.Bool(query + 132)) {
        query_ffi["id"] = A.load.Int32(query + 60);
      }
      query_ffi["url"] = A.load.Ref(query + 64, undefined);
      query_ffi["finalUrl"] = A.load.Ref(query + 68, undefined);
      query_ffi["filename"] = A.load.Ref(query + 72, undefined);
      query_ffi["danger"] = A.load.Enum(query + 76, [
        "file",
        "url",
        "content",
        "uncommon",
        "host",
        "unwanted",
        "safe",
        "accepted",
        "allowlistedByPolicy",
        "asyncScanning",
        "passwordProtected",
        "blockedTooLarge",
        "sensitiveContentWarning",
        "sensitiveContentBlock",
        "unsupportedFileType",
        "deepScannedFailed",
        "deepScannedSafe",
        "deepScannedOpenedDangerous",
        "promptForScaning",
        "accountCompromise",
      ]);
      query_ffi["mime"] = A.load.Ref(query + 80, undefined);
      query_ffi["startTime"] = A.load.Ref(query + 84, undefined);
      query_ffi["endTime"] = A.load.Ref(query + 88, undefined);
      query_ffi["state"] = A.load.Enum(query + 92, ["in_progress", "interrupted", "complete"]);
      if (A.load.Bool(query + 133)) {
        query_ffi["paused"] = A.load.Bool(query + 96);
      }
      query_ffi["error"] = A.load.Enum(query + 100, [
        "FILE_FAILED",
        "FILE_ACCESS_DENIED",
        "FILE_NO_SPACE",
        "FILE_NAME_TOO_LONG",
        "FILE_TOO_LARGE",
        "FILE_VIRUS_INFECTED",
        "FILE_TRANSIENT_ERROR",
        "FILE_BLOCKED",
        "FILE_SECURITY_CHECK_FAILED",
        "FILE_TOO_SHORT",
        "FILE_HASH_MISMATCH",
        "FILE_SAME_AS_SOURCE",
        "NETWORK_FAILED",
        "NETWORK_TIMEOUT",
        "NETWORK_DISCONNECTED",
        "NETWORK_SERVER_DOWN",
        "NETWORK_INVALID_REQUEST",
        "SERVER_FAILED",
        "SERVER_NO_RANGE",
        "SERVER_BAD_CONTENT",
        "SERVER_UNAUTHORIZED",
        "SERVER_CERT_PROBLEM",
        "SERVER_FORBIDDEN",
        "SERVER_UNREACHABLE",
        "SERVER_CONTENT_LENGTH_MISMATCH",
        "SERVER_CROSS_ORIGIN_REDIRECT",
        "USER_CANCELED",
        "USER_SHUTDOWN",
        "CRASH",
      ]);
      if (A.load.Bool(query + 134)) {
        query_ffi["bytesReceived"] = A.load.Float64(query + 104);
      }
      if (A.load.Bool(query + 135)) {
        query_ffi["totalBytes"] = A.load.Float64(query + 112);
      }
      if (A.load.Bool(query + 136)) {
        query_ffi["fileSize"] = A.load.Float64(query + 120);
      }
      if (A.load.Bool(query + 137)) {
        query_ffi["exists"] = A.load.Bool(query + 128);
      }

      const _ret = WEBEXT.downloads.erase(query_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Erase": (retPtr: Pointer, errPtr: Pointer, query: Pointer): heap.Ref<boolean> => {
      try {
        const query_ffi = {};

        query_ffi["query"] = A.load.Ref(query + 0, undefined);
        query_ffi["startedBefore"] = A.load.Ref(query + 4, undefined);
        query_ffi["startedAfter"] = A.load.Ref(query + 8, undefined);
        query_ffi["endedBefore"] = A.load.Ref(query + 12, undefined);
        query_ffi["endedAfter"] = A.load.Ref(query + 16, undefined);
        if (A.load.Bool(query + 129)) {
          query_ffi["totalBytesGreater"] = A.load.Float64(query + 24);
        }
        if (A.load.Bool(query + 130)) {
          query_ffi["totalBytesLess"] = A.load.Float64(query + 32);
        }
        query_ffi["filenameRegex"] = A.load.Ref(query + 40, undefined);
        query_ffi["urlRegex"] = A.load.Ref(query + 44, undefined);
        query_ffi["finalUrlRegex"] = A.load.Ref(query + 48, undefined);
        if (A.load.Bool(query + 131)) {
          query_ffi["limit"] = A.load.Int32(query + 52);
        }
        query_ffi["orderBy"] = A.load.Ref(query + 56, undefined);
        if (A.load.Bool(query + 132)) {
          query_ffi["id"] = A.load.Int32(query + 60);
        }
        query_ffi["url"] = A.load.Ref(query + 64, undefined);
        query_ffi["finalUrl"] = A.load.Ref(query + 68, undefined);
        query_ffi["filename"] = A.load.Ref(query + 72, undefined);
        query_ffi["danger"] = A.load.Enum(query + 76, [
          "file",
          "url",
          "content",
          "uncommon",
          "host",
          "unwanted",
          "safe",
          "accepted",
          "allowlistedByPolicy",
          "asyncScanning",
          "passwordProtected",
          "blockedTooLarge",
          "sensitiveContentWarning",
          "sensitiveContentBlock",
          "unsupportedFileType",
          "deepScannedFailed",
          "deepScannedSafe",
          "deepScannedOpenedDangerous",
          "promptForScaning",
          "accountCompromise",
        ]);
        query_ffi["mime"] = A.load.Ref(query + 80, undefined);
        query_ffi["startTime"] = A.load.Ref(query + 84, undefined);
        query_ffi["endTime"] = A.load.Ref(query + 88, undefined);
        query_ffi["state"] = A.load.Enum(query + 92, ["in_progress", "interrupted", "complete"]);
        if (A.load.Bool(query + 133)) {
          query_ffi["paused"] = A.load.Bool(query + 96);
        }
        query_ffi["error"] = A.load.Enum(query + 100, [
          "FILE_FAILED",
          "FILE_ACCESS_DENIED",
          "FILE_NO_SPACE",
          "FILE_NAME_TOO_LONG",
          "FILE_TOO_LARGE",
          "FILE_VIRUS_INFECTED",
          "FILE_TRANSIENT_ERROR",
          "FILE_BLOCKED",
          "FILE_SECURITY_CHECK_FAILED",
          "FILE_TOO_SHORT",
          "FILE_HASH_MISMATCH",
          "FILE_SAME_AS_SOURCE",
          "NETWORK_FAILED",
          "NETWORK_TIMEOUT",
          "NETWORK_DISCONNECTED",
          "NETWORK_SERVER_DOWN",
          "NETWORK_INVALID_REQUEST",
          "SERVER_FAILED",
          "SERVER_NO_RANGE",
          "SERVER_BAD_CONTENT",
          "SERVER_UNAUTHORIZED",
          "SERVER_CERT_PROBLEM",
          "SERVER_FORBIDDEN",
          "SERVER_UNREACHABLE",
          "SERVER_CONTENT_LENGTH_MISMATCH",
          "SERVER_CROSS_ORIGIN_REDIRECT",
          "USER_CANCELED",
          "USER_SHUTDOWN",
          "CRASH",
        ]);
        if (A.load.Bool(query + 134)) {
          query_ffi["bytesReceived"] = A.load.Float64(query + 104);
        }
        if (A.load.Bool(query + 135)) {
          query_ffi["totalBytes"] = A.load.Float64(query + 112);
        }
        if (A.load.Bool(query + 136)) {
          query_ffi["fileSize"] = A.load.Float64(query + 120);
        }
        if (A.load.Bool(query + 137)) {
          query_ffi["exists"] = A.load.Bool(query + 128);
        }

        const _ret = WEBEXT.downloads.erase(query_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetFileIcon": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads && "getFileIcon" in WEBEXT?.downloads) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetFileIcon": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.getFileIcon);
    },
    "call_GetFileIcon": (retPtr: Pointer, downloadId: number, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 4)) {
        options_ffi["size"] = A.load.Int32(options + 0);
      }

      const _ret = WEBEXT.downloads.getFileIcon(downloadId, options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetFileIcon": (retPtr: Pointer, errPtr: Pointer, downloadId: number, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 4)) {
          options_ffi["size"] = A.load.Int32(options + 0);
        }

        const _ret = WEBEXT.downloads.getFileIcon(downloadId, options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads?.onChanged && "addListener" in WEBEXT?.downloads?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.onChanged.addListener);
    },
    "call_OnChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.downloads.onChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.onChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads?.onChanged && "removeListener" in WEBEXT?.downloads?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.onChanged.removeListener);
    },
    "call_OffChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.downloads.onChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.onChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads?.onChanged && "hasListener" in WEBEXT?.downloads?.onChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.onChanged.hasListener);
    },
    "call_HasOnChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.downloads.onChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.onChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads?.onCreated && "addListener" in WEBEXT?.downloads?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.onCreated.addListener);
    },
    "call_OnCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.downloads.onCreated.addListener(A.H.get<object>(callback));
    },
    "try_OnCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.onCreated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads?.onCreated && "removeListener" in WEBEXT?.downloads?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.onCreated.removeListener);
    },
    "call_OffCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.downloads.onCreated.removeListener(A.H.get<object>(callback));
    },
    "try_OffCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.onCreated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads?.onCreated && "hasListener" in WEBEXT?.downloads?.onCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.onCreated.hasListener);
    },
    "call_HasOnCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.downloads.onCreated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.onCreated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeterminingFilename": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads?.onDeterminingFilename && "addListener" in WEBEXT?.downloads?.onDeterminingFilename) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeterminingFilename": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.onDeterminingFilename.addListener);
    },
    "call_OnDeterminingFilename": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.downloads.onDeterminingFilename.addListener(A.H.get<object>(callback));
    },
    "try_OnDeterminingFilename": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.onDeterminingFilename.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeterminingFilename": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads?.onDeterminingFilename && "removeListener" in WEBEXT?.downloads?.onDeterminingFilename) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeterminingFilename": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.onDeterminingFilename.removeListener);
    },
    "call_OffDeterminingFilename": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.downloads.onDeterminingFilename.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeterminingFilename": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.onDeterminingFilename.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeterminingFilename": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads?.onDeterminingFilename && "hasListener" in WEBEXT?.downloads?.onDeterminingFilename) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeterminingFilename": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.onDeterminingFilename.hasListener);
    },
    "call_HasOnDeterminingFilename": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.downloads.onDeterminingFilename.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeterminingFilename": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.onDeterminingFilename.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnErased": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads?.onErased && "addListener" in WEBEXT?.downloads?.onErased) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnErased": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.onErased.addListener);
    },
    "call_OnErased": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.downloads.onErased.addListener(A.H.get<object>(callback));
    },
    "try_OnErased": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.onErased.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffErased": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads?.onErased && "removeListener" in WEBEXT?.downloads?.onErased) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffErased": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.onErased.removeListener);
    },
    "call_OffErased": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.downloads.onErased.removeListener(A.H.get<object>(callback));
    },
    "try_OffErased": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.onErased.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnErased": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads?.onErased && "hasListener" in WEBEXT?.downloads?.onErased) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnErased": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.onErased.hasListener);
    },
    "call_HasOnErased": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.downloads.onErased.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnErased": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.onErased.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Open": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads && "open" in WEBEXT?.downloads) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Open": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.open);
    },
    "call_Open": (retPtr: Pointer, downloadId: number): void => {
      const _ret = WEBEXT.downloads.open(downloadId);
    },
    "try_Open": (retPtr: Pointer, errPtr: Pointer, downloadId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.open(downloadId);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Pause": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads && "pause" in WEBEXT?.downloads) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Pause": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.pause);
    },
    "call_Pause": (retPtr: Pointer, downloadId: number): void => {
      const _ret = WEBEXT.downloads.pause(downloadId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Pause": (retPtr: Pointer, errPtr: Pointer, downloadId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.pause(downloadId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveFile": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads && "removeFile" in WEBEXT?.downloads) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveFile": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.removeFile);
    },
    "call_RemoveFile": (retPtr: Pointer, downloadId: number): void => {
      const _ret = WEBEXT.downloads.removeFile(downloadId);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveFile": (retPtr: Pointer, errPtr: Pointer, downloadId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.removeFile(downloadId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Resume": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads && "resume" in WEBEXT?.downloads) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Resume": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.resume);
    },
    "call_Resume": (retPtr: Pointer, downloadId: number): void => {
      const _ret = WEBEXT.downloads.resume(downloadId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Resume": (retPtr: Pointer, errPtr: Pointer, downloadId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.resume(downloadId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Search": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads && "search" in WEBEXT?.downloads) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Search": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.search);
    },
    "call_Search": (retPtr: Pointer, query: Pointer): void => {
      const query_ffi = {};

      query_ffi["query"] = A.load.Ref(query + 0, undefined);
      query_ffi["startedBefore"] = A.load.Ref(query + 4, undefined);
      query_ffi["startedAfter"] = A.load.Ref(query + 8, undefined);
      query_ffi["endedBefore"] = A.load.Ref(query + 12, undefined);
      query_ffi["endedAfter"] = A.load.Ref(query + 16, undefined);
      if (A.load.Bool(query + 129)) {
        query_ffi["totalBytesGreater"] = A.load.Float64(query + 24);
      }
      if (A.load.Bool(query + 130)) {
        query_ffi["totalBytesLess"] = A.load.Float64(query + 32);
      }
      query_ffi["filenameRegex"] = A.load.Ref(query + 40, undefined);
      query_ffi["urlRegex"] = A.load.Ref(query + 44, undefined);
      query_ffi["finalUrlRegex"] = A.load.Ref(query + 48, undefined);
      if (A.load.Bool(query + 131)) {
        query_ffi["limit"] = A.load.Int32(query + 52);
      }
      query_ffi["orderBy"] = A.load.Ref(query + 56, undefined);
      if (A.load.Bool(query + 132)) {
        query_ffi["id"] = A.load.Int32(query + 60);
      }
      query_ffi["url"] = A.load.Ref(query + 64, undefined);
      query_ffi["finalUrl"] = A.load.Ref(query + 68, undefined);
      query_ffi["filename"] = A.load.Ref(query + 72, undefined);
      query_ffi["danger"] = A.load.Enum(query + 76, [
        "file",
        "url",
        "content",
        "uncommon",
        "host",
        "unwanted",
        "safe",
        "accepted",
        "allowlistedByPolicy",
        "asyncScanning",
        "passwordProtected",
        "blockedTooLarge",
        "sensitiveContentWarning",
        "sensitiveContentBlock",
        "unsupportedFileType",
        "deepScannedFailed",
        "deepScannedSafe",
        "deepScannedOpenedDangerous",
        "promptForScaning",
        "accountCompromise",
      ]);
      query_ffi["mime"] = A.load.Ref(query + 80, undefined);
      query_ffi["startTime"] = A.load.Ref(query + 84, undefined);
      query_ffi["endTime"] = A.load.Ref(query + 88, undefined);
      query_ffi["state"] = A.load.Enum(query + 92, ["in_progress", "interrupted", "complete"]);
      if (A.load.Bool(query + 133)) {
        query_ffi["paused"] = A.load.Bool(query + 96);
      }
      query_ffi["error"] = A.load.Enum(query + 100, [
        "FILE_FAILED",
        "FILE_ACCESS_DENIED",
        "FILE_NO_SPACE",
        "FILE_NAME_TOO_LONG",
        "FILE_TOO_LARGE",
        "FILE_VIRUS_INFECTED",
        "FILE_TRANSIENT_ERROR",
        "FILE_BLOCKED",
        "FILE_SECURITY_CHECK_FAILED",
        "FILE_TOO_SHORT",
        "FILE_HASH_MISMATCH",
        "FILE_SAME_AS_SOURCE",
        "NETWORK_FAILED",
        "NETWORK_TIMEOUT",
        "NETWORK_DISCONNECTED",
        "NETWORK_SERVER_DOWN",
        "NETWORK_INVALID_REQUEST",
        "SERVER_FAILED",
        "SERVER_NO_RANGE",
        "SERVER_BAD_CONTENT",
        "SERVER_UNAUTHORIZED",
        "SERVER_CERT_PROBLEM",
        "SERVER_FORBIDDEN",
        "SERVER_UNREACHABLE",
        "SERVER_CONTENT_LENGTH_MISMATCH",
        "SERVER_CROSS_ORIGIN_REDIRECT",
        "USER_CANCELED",
        "USER_SHUTDOWN",
        "CRASH",
      ]);
      if (A.load.Bool(query + 134)) {
        query_ffi["bytesReceived"] = A.load.Float64(query + 104);
      }
      if (A.load.Bool(query + 135)) {
        query_ffi["totalBytes"] = A.load.Float64(query + 112);
      }
      if (A.load.Bool(query + 136)) {
        query_ffi["fileSize"] = A.load.Float64(query + 120);
      }
      if (A.load.Bool(query + 137)) {
        query_ffi["exists"] = A.load.Bool(query + 128);
      }

      const _ret = WEBEXT.downloads.search(query_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Search": (retPtr: Pointer, errPtr: Pointer, query: Pointer): heap.Ref<boolean> => {
      try {
        const query_ffi = {};

        query_ffi["query"] = A.load.Ref(query + 0, undefined);
        query_ffi["startedBefore"] = A.load.Ref(query + 4, undefined);
        query_ffi["startedAfter"] = A.load.Ref(query + 8, undefined);
        query_ffi["endedBefore"] = A.load.Ref(query + 12, undefined);
        query_ffi["endedAfter"] = A.load.Ref(query + 16, undefined);
        if (A.load.Bool(query + 129)) {
          query_ffi["totalBytesGreater"] = A.load.Float64(query + 24);
        }
        if (A.load.Bool(query + 130)) {
          query_ffi["totalBytesLess"] = A.load.Float64(query + 32);
        }
        query_ffi["filenameRegex"] = A.load.Ref(query + 40, undefined);
        query_ffi["urlRegex"] = A.load.Ref(query + 44, undefined);
        query_ffi["finalUrlRegex"] = A.load.Ref(query + 48, undefined);
        if (A.load.Bool(query + 131)) {
          query_ffi["limit"] = A.load.Int32(query + 52);
        }
        query_ffi["orderBy"] = A.load.Ref(query + 56, undefined);
        if (A.load.Bool(query + 132)) {
          query_ffi["id"] = A.load.Int32(query + 60);
        }
        query_ffi["url"] = A.load.Ref(query + 64, undefined);
        query_ffi["finalUrl"] = A.load.Ref(query + 68, undefined);
        query_ffi["filename"] = A.load.Ref(query + 72, undefined);
        query_ffi["danger"] = A.load.Enum(query + 76, [
          "file",
          "url",
          "content",
          "uncommon",
          "host",
          "unwanted",
          "safe",
          "accepted",
          "allowlistedByPolicy",
          "asyncScanning",
          "passwordProtected",
          "blockedTooLarge",
          "sensitiveContentWarning",
          "sensitiveContentBlock",
          "unsupportedFileType",
          "deepScannedFailed",
          "deepScannedSafe",
          "deepScannedOpenedDangerous",
          "promptForScaning",
          "accountCompromise",
        ]);
        query_ffi["mime"] = A.load.Ref(query + 80, undefined);
        query_ffi["startTime"] = A.load.Ref(query + 84, undefined);
        query_ffi["endTime"] = A.load.Ref(query + 88, undefined);
        query_ffi["state"] = A.load.Enum(query + 92, ["in_progress", "interrupted", "complete"]);
        if (A.load.Bool(query + 133)) {
          query_ffi["paused"] = A.load.Bool(query + 96);
        }
        query_ffi["error"] = A.load.Enum(query + 100, [
          "FILE_FAILED",
          "FILE_ACCESS_DENIED",
          "FILE_NO_SPACE",
          "FILE_NAME_TOO_LONG",
          "FILE_TOO_LARGE",
          "FILE_VIRUS_INFECTED",
          "FILE_TRANSIENT_ERROR",
          "FILE_BLOCKED",
          "FILE_SECURITY_CHECK_FAILED",
          "FILE_TOO_SHORT",
          "FILE_HASH_MISMATCH",
          "FILE_SAME_AS_SOURCE",
          "NETWORK_FAILED",
          "NETWORK_TIMEOUT",
          "NETWORK_DISCONNECTED",
          "NETWORK_SERVER_DOWN",
          "NETWORK_INVALID_REQUEST",
          "SERVER_FAILED",
          "SERVER_NO_RANGE",
          "SERVER_BAD_CONTENT",
          "SERVER_UNAUTHORIZED",
          "SERVER_CERT_PROBLEM",
          "SERVER_FORBIDDEN",
          "SERVER_UNREACHABLE",
          "SERVER_CONTENT_LENGTH_MISMATCH",
          "SERVER_CROSS_ORIGIN_REDIRECT",
          "USER_CANCELED",
          "USER_SHUTDOWN",
          "CRASH",
        ]);
        if (A.load.Bool(query + 134)) {
          query_ffi["bytesReceived"] = A.load.Float64(query + 104);
        }
        if (A.load.Bool(query + 135)) {
          query_ffi["totalBytes"] = A.load.Float64(query + 112);
        }
        if (A.load.Bool(query + 136)) {
          query_ffi["fileSize"] = A.load.Float64(query + 120);
        }
        if (A.load.Bool(query + 137)) {
          query_ffi["exists"] = A.load.Bool(query + 128);
        }

        const _ret = WEBEXT.downloads.search(query_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetShelfEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads && "setShelfEnabled" in WEBEXT?.downloads) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetShelfEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.setShelfEnabled);
    },
    "call_SetShelfEnabled": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.downloads.setShelfEnabled(enabled === A.H.TRUE);
    },
    "try_SetShelfEnabled": (retPtr: Pointer, errPtr: Pointer, enabled: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.setShelfEnabled(enabled === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetUiOptions": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads && "setUiOptions" in WEBEXT?.downloads) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetUiOptions": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.setUiOptions);
    },
    "call_SetUiOptions": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 1)) {
        options_ffi["enabled"] = A.load.Bool(options + 0);
      }

      const _ret = WEBEXT.downloads.setUiOptions(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetUiOptions": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 1)) {
          options_ffi["enabled"] = A.load.Bool(options + 0);
        }

        const _ret = WEBEXT.downloads.setUiOptions(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Show": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads && "show" in WEBEXT?.downloads) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Show": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.show);
    },
    "call_Show": (retPtr: Pointer, downloadId: number): void => {
      const _ret = WEBEXT.downloads.show(downloadId);
    },
    "try_Show": (retPtr: Pointer, errPtr: Pointer, downloadId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.show(downloadId);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ShowDefaultFolder": (): heap.Ref<boolean> => {
      if (WEBEXT?.downloads && "showDefaultFolder" in WEBEXT?.downloads) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ShowDefaultFolder": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.downloads.showDefaultFolder);
    },
    "call_ShowDefaultFolder": (retPtr: Pointer): void => {
      const _ret = WEBEXT.downloads.showDefaultFolder();
    },
    "try_ShowDefaultFolder": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.downloads.showDefaultFolder();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
