import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/smartcardproviderprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_ConnectionState": (ref: heap.Ref<string>): number => {
      const idx = ["ABSENT", "PRESENT", "SWALLOWED", "POWERED", "NEGOTIABLE", "SPECIFIC"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_Disposition": (ref: heap.Ref<string>): number => {
      const idx = ["LEAVE_CARD", "RESET_CARD", "UNPOWER_CARD", "EJECT_CARD"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_Protocol": (ref: heap.Ref<string>): number => {
      const idx = ["UNDEFINED", "T0", "T1", "RAW"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Protocols": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 2, false);
      } else {
        A.store.Bool(ptr + 6, true);
        A.store.Bool(ptr + 3, "t0" in x ? true : false);
        A.store.Bool(ptr + 0, x["t0"] ? true : false);
        A.store.Bool(ptr + 4, "t1" in x ? true : false);
        A.store.Bool(ptr + 1, x["t1"] ? true : false);
        A.store.Bool(ptr + 5, "raw" in x ? true : false);
        A.store.Bool(ptr + 2, x["raw"] ? true : false);
      }
    },
    "load_Protocols": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 3)) {
        x["t0"] = A.load.Bool(ptr + 0);
      } else {
        delete x["t0"];
      }
      if (A.load.Bool(ptr + 4)) {
        x["t1"] = A.load.Bool(ptr + 1);
      } else {
        delete x["t1"];
      }
      if (A.load.Bool(ptr + 5)) {
        x["raw"] = A.load.Bool(ptr + 2);
      } else {
        delete x["raw"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ReaderStateFlags": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 22, false);
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 10, false);
      } else {
        A.store.Bool(ptr + 22, true);
        A.store.Bool(ptr + 11, "unaware" in x ? true : false);
        A.store.Bool(ptr + 0, x["unaware"] ? true : false);
        A.store.Bool(ptr + 12, "ignore" in x ? true : false);
        A.store.Bool(ptr + 1, x["ignore"] ? true : false);
        A.store.Bool(ptr + 13, "changed" in x ? true : false);
        A.store.Bool(ptr + 2, x["changed"] ? true : false);
        A.store.Bool(ptr + 14, "unknown" in x ? true : false);
        A.store.Bool(ptr + 3, x["unknown"] ? true : false);
        A.store.Bool(ptr + 15, "unavailable" in x ? true : false);
        A.store.Bool(ptr + 4, x["unavailable"] ? true : false);
        A.store.Bool(ptr + 16, "empty" in x ? true : false);
        A.store.Bool(ptr + 5, x["empty"] ? true : false);
        A.store.Bool(ptr + 17, "present" in x ? true : false);
        A.store.Bool(ptr + 6, x["present"] ? true : false);
        A.store.Bool(ptr + 18, "exclusive" in x ? true : false);
        A.store.Bool(ptr + 7, x["exclusive"] ? true : false);
        A.store.Bool(ptr + 19, "inuse" in x ? true : false);
        A.store.Bool(ptr + 8, x["inuse"] ? true : false);
        A.store.Bool(ptr + 20, "mute" in x ? true : false);
        A.store.Bool(ptr + 9, x["mute"] ? true : false);
        A.store.Bool(ptr + 21, "unpowered" in x ? true : false);
        A.store.Bool(ptr + 10, x["unpowered"] ? true : false);
      }
    },
    "load_ReaderStateFlags": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 11)) {
        x["unaware"] = A.load.Bool(ptr + 0);
      } else {
        delete x["unaware"];
      }
      if (A.load.Bool(ptr + 12)) {
        x["ignore"] = A.load.Bool(ptr + 1);
      } else {
        delete x["ignore"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["changed"] = A.load.Bool(ptr + 2);
      } else {
        delete x["changed"];
      }
      if (A.load.Bool(ptr + 14)) {
        x["unknown"] = A.load.Bool(ptr + 3);
      } else {
        delete x["unknown"];
      }
      if (A.load.Bool(ptr + 15)) {
        x["unavailable"] = A.load.Bool(ptr + 4);
      } else {
        delete x["unavailable"];
      }
      if (A.load.Bool(ptr + 16)) {
        x["empty"] = A.load.Bool(ptr + 5);
      } else {
        delete x["empty"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["present"] = A.load.Bool(ptr + 6);
      } else {
        delete x["present"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["exclusive"] = A.load.Bool(ptr + 7);
      } else {
        delete x["exclusive"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["inuse"] = A.load.Bool(ptr + 8);
      } else {
        delete x["inuse"];
      }
      if (A.load.Bool(ptr + 20)) {
        x["mute"] = A.load.Bool(ptr + 9);
      } else {
        delete x["mute"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["unpowered"] = A.load.Bool(ptr + 10);
      } else {
        delete x["unpowered"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ReaderStateIn": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 33, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 4 + 22, false);
        A.store.Bool(ptr + 4 + 11, false);
        A.store.Bool(ptr + 4 + 0, false);
        A.store.Bool(ptr + 4 + 12, false);
        A.store.Bool(ptr + 4 + 1, false);
        A.store.Bool(ptr + 4 + 13, false);
        A.store.Bool(ptr + 4 + 2, false);
        A.store.Bool(ptr + 4 + 14, false);
        A.store.Bool(ptr + 4 + 3, false);
        A.store.Bool(ptr + 4 + 15, false);
        A.store.Bool(ptr + 4 + 4, false);
        A.store.Bool(ptr + 4 + 16, false);
        A.store.Bool(ptr + 4 + 5, false);
        A.store.Bool(ptr + 4 + 17, false);
        A.store.Bool(ptr + 4 + 6, false);
        A.store.Bool(ptr + 4 + 18, false);
        A.store.Bool(ptr + 4 + 7, false);
        A.store.Bool(ptr + 4 + 19, false);
        A.store.Bool(ptr + 4 + 8, false);
        A.store.Bool(ptr + 4 + 20, false);
        A.store.Bool(ptr + 4 + 9, false);
        A.store.Bool(ptr + 4 + 21, false);
        A.store.Bool(ptr + 4 + 10, false);
        A.store.Bool(ptr + 32, false);
        A.store.Int32(ptr + 28, 0);
      } else {
        A.store.Bool(ptr + 33, true);
        A.store.Ref(ptr + 0, x["reader"]);

        if (typeof x["currentState"] === "undefined") {
          A.store.Bool(ptr + 4 + 22, false);
          A.store.Bool(ptr + 4 + 11, false);
          A.store.Bool(ptr + 4 + 0, false);
          A.store.Bool(ptr + 4 + 12, false);
          A.store.Bool(ptr + 4 + 1, false);
          A.store.Bool(ptr + 4 + 13, false);
          A.store.Bool(ptr + 4 + 2, false);
          A.store.Bool(ptr + 4 + 14, false);
          A.store.Bool(ptr + 4 + 3, false);
          A.store.Bool(ptr + 4 + 15, false);
          A.store.Bool(ptr + 4 + 4, false);
          A.store.Bool(ptr + 4 + 16, false);
          A.store.Bool(ptr + 4 + 5, false);
          A.store.Bool(ptr + 4 + 17, false);
          A.store.Bool(ptr + 4 + 6, false);
          A.store.Bool(ptr + 4 + 18, false);
          A.store.Bool(ptr + 4 + 7, false);
          A.store.Bool(ptr + 4 + 19, false);
          A.store.Bool(ptr + 4 + 8, false);
          A.store.Bool(ptr + 4 + 20, false);
          A.store.Bool(ptr + 4 + 9, false);
          A.store.Bool(ptr + 4 + 21, false);
          A.store.Bool(ptr + 4 + 10, false);
        } else {
          A.store.Bool(ptr + 4 + 22, true);
          A.store.Bool(ptr + 4 + 11, "unaware" in x["currentState"] ? true : false);
          A.store.Bool(ptr + 4 + 0, x["currentState"]["unaware"] ? true : false);
          A.store.Bool(ptr + 4 + 12, "ignore" in x["currentState"] ? true : false);
          A.store.Bool(ptr + 4 + 1, x["currentState"]["ignore"] ? true : false);
          A.store.Bool(ptr + 4 + 13, "changed" in x["currentState"] ? true : false);
          A.store.Bool(ptr + 4 + 2, x["currentState"]["changed"] ? true : false);
          A.store.Bool(ptr + 4 + 14, "unknown" in x["currentState"] ? true : false);
          A.store.Bool(ptr + 4 + 3, x["currentState"]["unknown"] ? true : false);
          A.store.Bool(ptr + 4 + 15, "unavailable" in x["currentState"] ? true : false);
          A.store.Bool(ptr + 4 + 4, x["currentState"]["unavailable"] ? true : false);
          A.store.Bool(ptr + 4 + 16, "empty" in x["currentState"] ? true : false);
          A.store.Bool(ptr + 4 + 5, x["currentState"]["empty"] ? true : false);
          A.store.Bool(ptr + 4 + 17, "present" in x["currentState"] ? true : false);
          A.store.Bool(ptr + 4 + 6, x["currentState"]["present"] ? true : false);
          A.store.Bool(ptr + 4 + 18, "exclusive" in x["currentState"] ? true : false);
          A.store.Bool(ptr + 4 + 7, x["currentState"]["exclusive"] ? true : false);
          A.store.Bool(ptr + 4 + 19, "inuse" in x["currentState"] ? true : false);
          A.store.Bool(ptr + 4 + 8, x["currentState"]["inuse"] ? true : false);
          A.store.Bool(ptr + 4 + 20, "mute" in x["currentState"] ? true : false);
          A.store.Bool(ptr + 4 + 9, x["currentState"]["mute"] ? true : false);
          A.store.Bool(ptr + 4 + 21, "unpowered" in x["currentState"] ? true : false);
          A.store.Bool(ptr + 4 + 10, x["currentState"]["unpowered"] ? true : false);
        }
        A.store.Bool(ptr + 32, "currentCount" in x ? true : false);
        A.store.Int32(ptr + 28, x["currentCount"] === undefined ? 0 : (x["currentCount"] as number));
      }
    },
    "load_ReaderStateIn": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["reader"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 4 + 22)) {
        x["currentState"] = {};
        if (A.load.Bool(ptr + 4 + 11)) {
          x["currentState"]["unaware"] = A.load.Bool(ptr + 4 + 0);
        } else {
          delete x["currentState"]["unaware"];
        }
        if (A.load.Bool(ptr + 4 + 12)) {
          x["currentState"]["ignore"] = A.load.Bool(ptr + 4 + 1);
        } else {
          delete x["currentState"]["ignore"];
        }
        if (A.load.Bool(ptr + 4 + 13)) {
          x["currentState"]["changed"] = A.load.Bool(ptr + 4 + 2);
        } else {
          delete x["currentState"]["changed"];
        }
        if (A.load.Bool(ptr + 4 + 14)) {
          x["currentState"]["unknown"] = A.load.Bool(ptr + 4 + 3);
        } else {
          delete x["currentState"]["unknown"];
        }
        if (A.load.Bool(ptr + 4 + 15)) {
          x["currentState"]["unavailable"] = A.load.Bool(ptr + 4 + 4);
        } else {
          delete x["currentState"]["unavailable"];
        }
        if (A.load.Bool(ptr + 4 + 16)) {
          x["currentState"]["empty"] = A.load.Bool(ptr + 4 + 5);
        } else {
          delete x["currentState"]["empty"];
        }
        if (A.load.Bool(ptr + 4 + 17)) {
          x["currentState"]["present"] = A.load.Bool(ptr + 4 + 6);
        } else {
          delete x["currentState"]["present"];
        }
        if (A.load.Bool(ptr + 4 + 18)) {
          x["currentState"]["exclusive"] = A.load.Bool(ptr + 4 + 7);
        } else {
          delete x["currentState"]["exclusive"];
        }
        if (A.load.Bool(ptr + 4 + 19)) {
          x["currentState"]["inuse"] = A.load.Bool(ptr + 4 + 8);
        } else {
          delete x["currentState"]["inuse"];
        }
        if (A.load.Bool(ptr + 4 + 20)) {
          x["currentState"]["mute"] = A.load.Bool(ptr + 4 + 9);
        } else {
          delete x["currentState"]["mute"];
        }
        if (A.load.Bool(ptr + 4 + 21)) {
          x["currentState"]["unpowered"] = A.load.Bool(ptr + 4 + 10);
        } else {
          delete x["currentState"]["unpowered"];
        }
      } else {
        delete x["currentState"];
      }
      if (A.load.Bool(ptr + 32)) {
        x["currentCount"] = A.load.Int32(ptr + 28);
      } else {
        delete x["currentCount"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ReaderStateOut": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 37, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 4 + 22, false);
        A.store.Bool(ptr + 4 + 11, false);
        A.store.Bool(ptr + 4 + 0, false);
        A.store.Bool(ptr + 4 + 12, false);
        A.store.Bool(ptr + 4 + 1, false);
        A.store.Bool(ptr + 4 + 13, false);
        A.store.Bool(ptr + 4 + 2, false);
        A.store.Bool(ptr + 4 + 14, false);
        A.store.Bool(ptr + 4 + 3, false);
        A.store.Bool(ptr + 4 + 15, false);
        A.store.Bool(ptr + 4 + 4, false);
        A.store.Bool(ptr + 4 + 16, false);
        A.store.Bool(ptr + 4 + 5, false);
        A.store.Bool(ptr + 4 + 17, false);
        A.store.Bool(ptr + 4 + 6, false);
        A.store.Bool(ptr + 4 + 18, false);
        A.store.Bool(ptr + 4 + 7, false);
        A.store.Bool(ptr + 4 + 19, false);
        A.store.Bool(ptr + 4 + 8, false);
        A.store.Bool(ptr + 4 + 20, false);
        A.store.Bool(ptr + 4 + 9, false);
        A.store.Bool(ptr + 4 + 21, false);
        A.store.Bool(ptr + 4 + 10, false);
        A.store.Bool(ptr + 36, false);
        A.store.Int32(ptr + 28, 0);
        A.store.Ref(ptr + 32, undefined);
      } else {
        A.store.Bool(ptr + 37, true);
        A.store.Ref(ptr + 0, x["reader"]);

        if (typeof x["eventState"] === "undefined") {
          A.store.Bool(ptr + 4 + 22, false);
          A.store.Bool(ptr + 4 + 11, false);
          A.store.Bool(ptr + 4 + 0, false);
          A.store.Bool(ptr + 4 + 12, false);
          A.store.Bool(ptr + 4 + 1, false);
          A.store.Bool(ptr + 4 + 13, false);
          A.store.Bool(ptr + 4 + 2, false);
          A.store.Bool(ptr + 4 + 14, false);
          A.store.Bool(ptr + 4 + 3, false);
          A.store.Bool(ptr + 4 + 15, false);
          A.store.Bool(ptr + 4 + 4, false);
          A.store.Bool(ptr + 4 + 16, false);
          A.store.Bool(ptr + 4 + 5, false);
          A.store.Bool(ptr + 4 + 17, false);
          A.store.Bool(ptr + 4 + 6, false);
          A.store.Bool(ptr + 4 + 18, false);
          A.store.Bool(ptr + 4 + 7, false);
          A.store.Bool(ptr + 4 + 19, false);
          A.store.Bool(ptr + 4 + 8, false);
          A.store.Bool(ptr + 4 + 20, false);
          A.store.Bool(ptr + 4 + 9, false);
          A.store.Bool(ptr + 4 + 21, false);
          A.store.Bool(ptr + 4 + 10, false);
        } else {
          A.store.Bool(ptr + 4 + 22, true);
          A.store.Bool(ptr + 4 + 11, "unaware" in x["eventState"] ? true : false);
          A.store.Bool(ptr + 4 + 0, x["eventState"]["unaware"] ? true : false);
          A.store.Bool(ptr + 4 + 12, "ignore" in x["eventState"] ? true : false);
          A.store.Bool(ptr + 4 + 1, x["eventState"]["ignore"] ? true : false);
          A.store.Bool(ptr + 4 + 13, "changed" in x["eventState"] ? true : false);
          A.store.Bool(ptr + 4 + 2, x["eventState"]["changed"] ? true : false);
          A.store.Bool(ptr + 4 + 14, "unknown" in x["eventState"] ? true : false);
          A.store.Bool(ptr + 4 + 3, x["eventState"]["unknown"] ? true : false);
          A.store.Bool(ptr + 4 + 15, "unavailable" in x["eventState"] ? true : false);
          A.store.Bool(ptr + 4 + 4, x["eventState"]["unavailable"] ? true : false);
          A.store.Bool(ptr + 4 + 16, "empty" in x["eventState"] ? true : false);
          A.store.Bool(ptr + 4 + 5, x["eventState"]["empty"] ? true : false);
          A.store.Bool(ptr + 4 + 17, "present" in x["eventState"] ? true : false);
          A.store.Bool(ptr + 4 + 6, x["eventState"]["present"] ? true : false);
          A.store.Bool(ptr + 4 + 18, "exclusive" in x["eventState"] ? true : false);
          A.store.Bool(ptr + 4 + 7, x["eventState"]["exclusive"] ? true : false);
          A.store.Bool(ptr + 4 + 19, "inuse" in x["eventState"] ? true : false);
          A.store.Bool(ptr + 4 + 8, x["eventState"]["inuse"] ? true : false);
          A.store.Bool(ptr + 4 + 20, "mute" in x["eventState"] ? true : false);
          A.store.Bool(ptr + 4 + 9, x["eventState"]["mute"] ? true : false);
          A.store.Bool(ptr + 4 + 21, "unpowered" in x["eventState"] ? true : false);
          A.store.Bool(ptr + 4 + 10, x["eventState"]["unpowered"] ? true : false);
        }
        A.store.Bool(ptr + 36, "eventCount" in x ? true : false);
        A.store.Int32(ptr + 28, x["eventCount"] === undefined ? 0 : (x["eventCount"] as number));
        A.store.Ref(ptr + 32, x["atr"]);
      }
    },
    "load_ReaderStateOut": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["reader"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 4 + 22)) {
        x["eventState"] = {};
        if (A.load.Bool(ptr + 4 + 11)) {
          x["eventState"]["unaware"] = A.load.Bool(ptr + 4 + 0);
        } else {
          delete x["eventState"]["unaware"];
        }
        if (A.load.Bool(ptr + 4 + 12)) {
          x["eventState"]["ignore"] = A.load.Bool(ptr + 4 + 1);
        } else {
          delete x["eventState"]["ignore"];
        }
        if (A.load.Bool(ptr + 4 + 13)) {
          x["eventState"]["changed"] = A.load.Bool(ptr + 4 + 2);
        } else {
          delete x["eventState"]["changed"];
        }
        if (A.load.Bool(ptr + 4 + 14)) {
          x["eventState"]["unknown"] = A.load.Bool(ptr + 4 + 3);
        } else {
          delete x["eventState"]["unknown"];
        }
        if (A.load.Bool(ptr + 4 + 15)) {
          x["eventState"]["unavailable"] = A.load.Bool(ptr + 4 + 4);
        } else {
          delete x["eventState"]["unavailable"];
        }
        if (A.load.Bool(ptr + 4 + 16)) {
          x["eventState"]["empty"] = A.load.Bool(ptr + 4 + 5);
        } else {
          delete x["eventState"]["empty"];
        }
        if (A.load.Bool(ptr + 4 + 17)) {
          x["eventState"]["present"] = A.load.Bool(ptr + 4 + 6);
        } else {
          delete x["eventState"]["present"];
        }
        if (A.load.Bool(ptr + 4 + 18)) {
          x["eventState"]["exclusive"] = A.load.Bool(ptr + 4 + 7);
        } else {
          delete x["eventState"]["exclusive"];
        }
        if (A.load.Bool(ptr + 4 + 19)) {
          x["eventState"]["inuse"] = A.load.Bool(ptr + 4 + 8);
        } else {
          delete x["eventState"]["inuse"];
        }
        if (A.load.Bool(ptr + 4 + 20)) {
          x["eventState"]["mute"] = A.load.Bool(ptr + 4 + 9);
        } else {
          delete x["eventState"]["mute"];
        }
        if (A.load.Bool(ptr + 4 + 21)) {
          x["eventState"]["unpowered"] = A.load.Bool(ptr + 4 + 10);
        } else {
          delete x["eventState"]["unpowered"];
        }
      } else {
        delete x["eventState"];
      }
      if (A.load.Bool(ptr + 36)) {
        x["eventCount"] = A.load.Int32(ptr + 28);
      } else {
        delete x["eventCount"];
      }
      x["atr"] = A.load.Ref(ptr + 32, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ResultCode": (ref: heap.Ref<string>): number => {
      const idx = [
        "SUCCESS",
        "REMOVED_CARD",
        "RESET_CARD",
        "UNPOWERED_CARD",
        "UNRESPONSIVE_CARD",
        "UNSUPPORTED_CARD",
        "READER_UNAVAILABLE",
        "SHARING_VIOLATION",
        "NOT_TRANSACTED",
        "NO_SMARTCARD",
        "PROTO_MISMATCH",
        "SYSTEM_CANCELLED",
        "NOT_READY",
        "CANCELLED",
        "INSUFFICIENT_BUFFER",
        "INVALID_HANDLE",
        "INVALID_PARAMETER",
        "INVALID_VALUE",
        "NO_MEMORY",
        "TIMEOUT",
        "UNKNOWN_READER",
        "UNSUPPORTED_FEATURE",
        "NO_READERS_AVAILABLE",
        "SERVICE_STOPPED",
        "NO_SERVICE",
        "COMM_ERROR",
        "INTERNAL_ERROR",
        "UNKNOWN_ERROR",
        "SERVER_TOO_BUSY",
        "UNEXPECTED",
        "SHUTDOWN",
        "UNKNOWN",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ShareMode": (ref: heap.Ref<string>): number => {
      const idx = ["SHARED", "EXCLUSIVE", "DIRECT"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Timeout": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "milliseconds" in x ? true : false);
        A.store.Int32(ptr + 0, x["milliseconds"] === undefined ? 0 : (x["milliseconds"] as number));
      }
    },
    "load_Timeout": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["milliseconds"] = A.load.Int32(ptr + 0);
      } else {
        delete x["milliseconds"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_OnBeginTransactionRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onBeginTransactionRequested &&
        "addListener" in WEBEXT?.smartCardProviderPrivate?.onBeginTransactionRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnBeginTransactionRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.addListener);
    },
    "call_OnBeginTransactionRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnBeginTransactionRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffBeginTransactionRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onBeginTransactionRequested &&
        "removeListener" in WEBEXT?.smartCardProviderPrivate?.onBeginTransactionRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffBeginTransactionRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.removeListener);
    },
    "call_OffBeginTransactionRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.removeListener(
        A.H.get<object>(callback)
      );
    },
    "try_OffBeginTransactionRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnBeginTransactionRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onBeginTransactionRequested &&
        "hasListener" in WEBEXT?.smartCardProviderPrivate?.onBeginTransactionRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnBeginTransactionRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.hasListener);
    },
    "call_HasOnBeginTransactionRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnBeginTransactionRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onBeginTransactionRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCancelRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onCancelRequested &&
        "addListener" in WEBEXT?.smartCardProviderPrivate?.onCancelRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCancelRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onCancelRequested.addListener);
    },
    "call_OnCancelRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onCancelRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnCancelRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onCancelRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCancelRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onCancelRequested &&
        "removeListener" in WEBEXT?.smartCardProviderPrivate?.onCancelRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCancelRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onCancelRequested.removeListener);
    },
    "call_OffCancelRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onCancelRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffCancelRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onCancelRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCancelRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onCancelRequested &&
        "hasListener" in WEBEXT?.smartCardProviderPrivate?.onCancelRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCancelRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onCancelRequested.hasListener);
    },
    "call_HasOnCancelRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onCancelRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCancelRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onCancelRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnConnectRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onConnectRequested &&
        "addListener" in WEBEXT?.smartCardProviderPrivate?.onConnectRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnConnectRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onConnectRequested.addListener);
    },
    "call_OnConnectRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onConnectRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnConnectRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onConnectRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffConnectRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onConnectRequested &&
        "removeListener" in WEBEXT?.smartCardProviderPrivate?.onConnectRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffConnectRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onConnectRequested.removeListener);
    },
    "call_OffConnectRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onConnectRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffConnectRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onConnectRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnConnectRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onConnectRequested &&
        "hasListener" in WEBEXT?.smartCardProviderPrivate?.onConnectRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnConnectRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onConnectRequested.hasListener);
    },
    "call_HasOnConnectRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onConnectRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnConnectRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onConnectRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnControlRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onControlRequested &&
        "addListener" in WEBEXT?.smartCardProviderPrivate?.onControlRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnControlRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onControlRequested.addListener);
    },
    "call_OnControlRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onControlRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnControlRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onControlRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffControlRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onControlRequested &&
        "removeListener" in WEBEXT?.smartCardProviderPrivate?.onControlRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffControlRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onControlRequested.removeListener);
    },
    "call_OffControlRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onControlRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffControlRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onControlRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnControlRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onControlRequested &&
        "hasListener" in WEBEXT?.smartCardProviderPrivate?.onControlRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnControlRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onControlRequested.hasListener);
    },
    "call_HasOnControlRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onControlRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnControlRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onControlRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDisconnectRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onDisconnectRequested &&
        "addListener" in WEBEXT?.smartCardProviderPrivate?.onDisconnectRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDisconnectRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onDisconnectRequested.addListener);
    },
    "call_OnDisconnectRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onDisconnectRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnDisconnectRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onDisconnectRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDisconnectRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onDisconnectRequested &&
        "removeListener" in WEBEXT?.smartCardProviderPrivate?.onDisconnectRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDisconnectRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onDisconnectRequested.removeListener);
    },
    "call_OffDisconnectRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onDisconnectRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffDisconnectRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onDisconnectRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDisconnectRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onDisconnectRequested &&
        "hasListener" in WEBEXT?.smartCardProviderPrivate?.onDisconnectRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDisconnectRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onDisconnectRequested.hasListener);
    },
    "call_HasOnDisconnectRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onDisconnectRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDisconnectRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onDisconnectRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnEndTransactionRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onEndTransactionRequested &&
        "addListener" in WEBEXT?.smartCardProviderPrivate?.onEndTransactionRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnEndTransactionRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.addListener);
    },
    "call_OnEndTransactionRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnEndTransactionRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffEndTransactionRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onEndTransactionRequested &&
        "removeListener" in WEBEXT?.smartCardProviderPrivate?.onEndTransactionRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffEndTransactionRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.removeListener);
    },
    "call_OffEndTransactionRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffEndTransactionRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnEndTransactionRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onEndTransactionRequested &&
        "hasListener" in WEBEXT?.smartCardProviderPrivate?.onEndTransactionRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnEndTransactionRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.hasListener);
    },
    "call_HasOnEndTransactionRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnEndTransactionRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onEndTransactionRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnEstablishContextRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onEstablishContextRequested &&
        "addListener" in WEBEXT?.smartCardProviderPrivate?.onEstablishContextRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnEstablishContextRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.addListener);
    },
    "call_OnEstablishContextRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnEstablishContextRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffEstablishContextRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onEstablishContextRequested &&
        "removeListener" in WEBEXT?.smartCardProviderPrivate?.onEstablishContextRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffEstablishContextRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.removeListener);
    },
    "call_OffEstablishContextRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.removeListener(
        A.H.get<object>(callback)
      );
    },
    "try_OffEstablishContextRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnEstablishContextRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onEstablishContextRequested &&
        "hasListener" in WEBEXT?.smartCardProviderPrivate?.onEstablishContextRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnEstablishContextRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.hasListener);
    },
    "call_HasOnEstablishContextRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnEstablishContextRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onEstablishContextRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnGetAttribRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onGetAttribRequested &&
        "addListener" in WEBEXT?.smartCardProviderPrivate?.onGetAttribRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnGetAttribRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onGetAttribRequested.addListener);
    },
    "call_OnGetAttribRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onGetAttribRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnGetAttribRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onGetAttribRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffGetAttribRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onGetAttribRequested &&
        "removeListener" in WEBEXT?.smartCardProviderPrivate?.onGetAttribRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffGetAttribRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onGetAttribRequested.removeListener);
    },
    "call_OffGetAttribRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onGetAttribRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffGetAttribRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onGetAttribRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnGetAttribRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onGetAttribRequested &&
        "hasListener" in WEBEXT?.smartCardProviderPrivate?.onGetAttribRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnGetAttribRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onGetAttribRequested.hasListener);
    },
    "call_HasOnGetAttribRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onGetAttribRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnGetAttribRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onGetAttribRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnGetStatusChangeRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onGetStatusChangeRequested &&
        "addListener" in WEBEXT?.smartCardProviderPrivate?.onGetStatusChangeRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnGetStatusChangeRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.addListener);
    },
    "call_OnGetStatusChangeRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnGetStatusChangeRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffGetStatusChangeRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onGetStatusChangeRequested &&
        "removeListener" in WEBEXT?.smartCardProviderPrivate?.onGetStatusChangeRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffGetStatusChangeRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.removeListener);
    },
    "call_OffGetStatusChangeRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffGetStatusChangeRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnGetStatusChangeRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onGetStatusChangeRequested &&
        "hasListener" in WEBEXT?.smartCardProviderPrivate?.onGetStatusChangeRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnGetStatusChangeRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.hasListener);
    },
    "call_HasOnGetStatusChangeRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnGetStatusChangeRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onGetStatusChangeRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnListReadersRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onListReadersRequested &&
        "addListener" in WEBEXT?.smartCardProviderPrivate?.onListReadersRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnListReadersRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onListReadersRequested.addListener);
    },
    "call_OnListReadersRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onListReadersRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnListReadersRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onListReadersRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffListReadersRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onListReadersRequested &&
        "removeListener" in WEBEXT?.smartCardProviderPrivate?.onListReadersRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffListReadersRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onListReadersRequested.removeListener);
    },
    "call_OffListReadersRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onListReadersRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffListReadersRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onListReadersRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnListReadersRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onListReadersRequested &&
        "hasListener" in WEBEXT?.smartCardProviderPrivate?.onListReadersRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnListReadersRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onListReadersRequested.hasListener);
    },
    "call_HasOnListReadersRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onListReadersRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnListReadersRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onListReadersRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnReleaseContextRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onReleaseContextRequested &&
        "addListener" in WEBEXT?.smartCardProviderPrivate?.onReleaseContextRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnReleaseContextRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.addListener);
    },
    "call_OnReleaseContextRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnReleaseContextRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffReleaseContextRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onReleaseContextRequested &&
        "removeListener" in WEBEXT?.smartCardProviderPrivate?.onReleaseContextRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffReleaseContextRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.removeListener);
    },
    "call_OffReleaseContextRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffReleaseContextRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.removeListener(
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnReleaseContextRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onReleaseContextRequested &&
        "hasListener" in WEBEXT?.smartCardProviderPrivate?.onReleaseContextRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnReleaseContextRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.hasListener);
    },
    "call_HasOnReleaseContextRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnReleaseContextRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onReleaseContextRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSetAttribRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onSetAttribRequested &&
        "addListener" in WEBEXT?.smartCardProviderPrivate?.onSetAttribRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSetAttribRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onSetAttribRequested.addListener);
    },
    "call_OnSetAttribRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onSetAttribRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnSetAttribRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onSetAttribRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSetAttribRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onSetAttribRequested &&
        "removeListener" in WEBEXT?.smartCardProviderPrivate?.onSetAttribRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSetAttribRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onSetAttribRequested.removeListener);
    },
    "call_OffSetAttribRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onSetAttribRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffSetAttribRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onSetAttribRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSetAttribRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onSetAttribRequested &&
        "hasListener" in WEBEXT?.smartCardProviderPrivate?.onSetAttribRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSetAttribRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onSetAttribRequested.hasListener);
    },
    "call_HasOnSetAttribRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onSetAttribRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSetAttribRequested": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onSetAttribRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnStatusRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onStatusRequested &&
        "addListener" in WEBEXT?.smartCardProviderPrivate?.onStatusRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnStatusRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onStatusRequested.addListener);
    },
    "call_OnStatusRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onStatusRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnStatusRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onStatusRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffStatusRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onStatusRequested &&
        "removeListener" in WEBEXT?.smartCardProviderPrivate?.onStatusRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffStatusRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onStatusRequested.removeListener);
    },
    "call_OffStatusRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onStatusRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffStatusRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onStatusRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnStatusRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onStatusRequested &&
        "hasListener" in WEBEXT?.smartCardProviderPrivate?.onStatusRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnStatusRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onStatusRequested.hasListener);
    },
    "call_HasOnStatusRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onStatusRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnStatusRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onStatusRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTransmitRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onTransmitRequested &&
        "addListener" in WEBEXT?.smartCardProviderPrivate?.onTransmitRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTransmitRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onTransmitRequested.addListener);
    },
    "call_OnTransmitRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onTransmitRequested.addListener(A.H.get<object>(callback));
    },
    "try_OnTransmitRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onTransmitRequested.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffTransmitRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onTransmitRequested &&
        "removeListener" in WEBEXT?.smartCardProviderPrivate?.onTransmitRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffTransmitRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onTransmitRequested.removeListener);
    },
    "call_OffTransmitRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onTransmitRequested.removeListener(A.H.get<object>(callback));
    },
    "try_OffTransmitRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onTransmitRequested.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnTransmitRequested": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.smartCardProviderPrivate?.onTransmitRequested &&
        "hasListener" in WEBEXT?.smartCardProviderPrivate?.onTransmitRequested
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnTransmitRequested": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.onTransmitRequested.hasListener);
    },
    "call_HasOnTransmitRequested": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.onTransmitRequested.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnTransmitRequested": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.onTransmitRequested.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReportConnectResult": (): heap.Ref<boolean> => {
      if (WEBEXT?.smartCardProviderPrivate && "reportConnectResult" in WEBEXT?.smartCardProviderPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportConnectResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.reportConnectResult);
    },
    "call_ReportConnectResult": (
      retPtr: Pointer,
      requestId: number,
      scardHandle: number,
      activeProtocol: number,
      resultCode: number
    ): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.reportConnectResult(
        requestId,
        scardHandle,
        activeProtocol > 0 && activeProtocol <= 4 ? ["UNDEFINED", "T0", "T1", "RAW"][activeProtocol - 1] : undefined,
        resultCode > 0 && resultCode <= 32
          ? [
              "SUCCESS",
              "REMOVED_CARD",
              "RESET_CARD",
              "UNPOWERED_CARD",
              "UNRESPONSIVE_CARD",
              "UNSUPPORTED_CARD",
              "READER_UNAVAILABLE",
              "SHARING_VIOLATION",
              "NOT_TRANSACTED",
              "NO_SMARTCARD",
              "PROTO_MISMATCH",
              "SYSTEM_CANCELLED",
              "NOT_READY",
              "CANCELLED",
              "INSUFFICIENT_BUFFER",
              "INVALID_HANDLE",
              "INVALID_PARAMETER",
              "INVALID_VALUE",
              "NO_MEMORY",
              "TIMEOUT",
              "UNKNOWN_READER",
              "UNSUPPORTED_FEATURE",
              "NO_READERS_AVAILABLE",
              "SERVICE_STOPPED",
              "NO_SERVICE",
              "COMM_ERROR",
              "INTERNAL_ERROR",
              "UNKNOWN_ERROR",
              "SERVER_TOO_BUSY",
              "UNEXPECTED",
              "SHUTDOWN",
              "UNKNOWN",
            ][resultCode - 1]
          : undefined
      );
    },
    "try_ReportConnectResult": (
      retPtr: Pointer,
      errPtr: Pointer,
      requestId: number,
      scardHandle: number,
      activeProtocol: number,
      resultCode: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.reportConnectResult(
          requestId,
          scardHandle,
          activeProtocol > 0 && activeProtocol <= 4 ? ["UNDEFINED", "T0", "T1", "RAW"][activeProtocol - 1] : undefined,
          resultCode > 0 && resultCode <= 32
            ? [
                "SUCCESS",
                "REMOVED_CARD",
                "RESET_CARD",
                "UNPOWERED_CARD",
                "UNRESPONSIVE_CARD",
                "UNSUPPORTED_CARD",
                "READER_UNAVAILABLE",
                "SHARING_VIOLATION",
                "NOT_TRANSACTED",
                "NO_SMARTCARD",
                "PROTO_MISMATCH",
                "SYSTEM_CANCELLED",
                "NOT_READY",
                "CANCELLED",
                "INSUFFICIENT_BUFFER",
                "INVALID_HANDLE",
                "INVALID_PARAMETER",
                "INVALID_VALUE",
                "NO_MEMORY",
                "TIMEOUT",
                "UNKNOWN_READER",
                "UNSUPPORTED_FEATURE",
                "NO_READERS_AVAILABLE",
                "SERVICE_STOPPED",
                "NO_SERVICE",
                "COMM_ERROR",
                "INTERNAL_ERROR",
                "UNKNOWN_ERROR",
                "SERVER_TOO_BUSY",
                "UNEXPECTED",
                "SHUTDOWN",
                "UNKNOWN",
              ][resultCode - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReportDataResult": (): heap.Ref<boolean> => {
      if (WEBEXT?.smartCardProviderPrivate && "reportDataResult" in WEBEXT?.smartCardProviderPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportDataResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.reportDataResult);
    },
    "call_ReportDataResult": (retPtr: Pointer, requestId: number, data: heap.Ref<object>, resultCode: number): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.reportDataResult(
        requestId,
        A.H.get<object>(data),
        resultCode > 0 && resultCode <= 32
          ? [
              "SUCCESS",
              "REMOVED_CARD",
              "RESET_CARD",
              "UNPOWERED_CARD",
              "UNRESPONSIVE_CARD",
              "UNSUPPORTED_CARD",
              "READER_UNAVAILABLE",
              "SHARING_VIOLATION",
              "NOT_TRANSACTED",
              "NO_SMARTCARD",
              "PROTO_MISMATCH",
              "SYSTEM_CANCELLED",
              "NOT_READY",
              "CANCELLED",
              "INSUFFICIENT_BUFFER",
              "INVALID_HANDLE",
              "INVALID_PARAMETER",
              "INVALID_VALUE",
              "NO_MEMORY",
              "TIMEOUT",
              "UNKNOWN_READER",
              "UNSUPPORTED_FEATURE",
              "NO_READERS_AVAILABLE",
              "SERVICE_STOPPED",
              "NO_SERVICE",
              "COMM_ERROR",
              "INTERNAL_ERROR",
              "UNKNOWN_ERROR",
              "SERVER_TOO_BUSY",
              "UNEXPECTED",
              "SHUTDOWN",
              "UNKNOWN",
            ][resultCode - 1]
          : undefined
      );
    },
    "try_ReportDataResult": (
      retPtr: Pointer,
      errPtr: Pointer,
      requestId: number,
      data: heap.Ref<object>,
      resultCode: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.reportDataResult(
          requestId,
          A.H.get<object>(data),
          resultCode > 0 && resultCode <= 32
            ? [
                "SUCCESS",
                "REMOVED_CARD",
                "RESET_CARD",
                "UNPOWERED_CARD",
                "UNRESPONSIVE_CARD",
                "UNSUPPORTED_CARD",
                "READER_UNAVAILABLE",
                "SHARING_VIOLATION",
                "NOT_TRANSACTED",
                "NO_SMARTCARD",
                "PROTO_MISMATCH",
                "SYSTEM_CANCELLED",
                "NOT_READY",
                "CANCELLED",
                "INSUFFICIENT_BUFFER",
                "INVALID_HANDLE",
                "INVALID_PARAMETER",
                "INVALID_VALUE",
                "NO_MEMORY",
                "TIMEOUT",
                "UNKNOWN_READER",
                "UNSUPPORTED_FEATURE",
                "NO_READERS_AVAILABLE",
                "SERVICE_STOPPED",
                "NO_SERVICE",
                "COMM_ERROR",
                "INTERNAL_ERROR",
                "UNKNOWN_ERROR",
                "SERVER_TOO_BUSY",
                "UNEXPECTED",
                "SHUTDOWN",
                "UNKNOWN",
              ][resultCode - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReportEstablishContextResult": (): heap.Ref<boolean> => {
      if (WEBEXT?.smartCardProviderPrivate && "reportEstablishContextResult" in WEBEXT?.smartCardProviderPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportEstablishContextResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.reportEstablishContextResult);
    },
    "call_ReportEstablishContextResult": (
      retPtr: Pointer,
      requestId: number,
      scardContext: number,
      resultCode: number
    ): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.reportEstablishContextResult(
        requestId,
        scardContext,
        resultCode > 0 && resultCode <= 32
          ? [
              "SUCCESS",
              "REMOVED_CARD",
              "RESET_CARD",
              "UNPOWERED_CARD",
              "UNRESPONSIVE_CARD",
              "UNSUPPORTED_CARD",
              "READER_UNAVAILABLE",
              "SHARING_VIOLATION",
              "NOT_TRANSACTED",
              "NO_SMARTCARD",
              "PROTO_MISMATCH",
              "SYSTEM_CANCELLED",
              "NOT_READY",
              "CANCELLED",
              "INSUFFICIENT_BUFFER",
              "INVALID_HANDLE",
              "INVALID_PARAMETER",
              "INVALID_VALUE",
              "NO_MEMORY",
              "TIMEOUT",
              "UNKNOWN_READER",
              "UNSUPPORTED_FEATURE",
              "NO_READERS_AVAILABLE",
              "SERVICE_STOPPED",
              "NO_SERVICE",
              "COMM_ERROR",
              "INTERNAL_ERROR",
              "UNKNOWN_ERROR",
              "SERVER_TOO_BUSY",
              "UNEXPECTED",
              "SHUTDOWN",
              "UNKNOWN",
            ][resultCode - 1]
          : undefined
      );
    },
    "try_ReportEstablishContextResult": (
      retPtr: Pointer,
      errPtr: Pointer,
      requestId: number,
      scardContext: number,
      resultCode: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.reportEstablishContextResult(
          requestId,
          scardContext,
          resultCode > 0 && resultCode <= 32
            ? [
                "SUCCESS",
                "REMOVED_CARD",
                "RESET_CARD",
                "UNPOWERED_CARD",
                "UNRESPONSIVE_CARD",
                "UNSUPPORTED_CARD",
                "READER_UNAVAILABLE",
                "SHARING_VIOLATION",
                "NOT_TRANSACTED",
                "NO_SMARTCARD",
                "PROTO_MISMATCH",
                "SYSTEM_CANCELLED",
                "NOT_READY",
                "CANCELLED",
                "INSUFFICIENT_BUFFER",
                "INVALID_HANDLE",
                "INVALID_PARAMETER",
                "INVALID_VALUE",
                "NO_MEMORY",
                "TIMEOUT",
                "UNKNOWN_READER",
                "UNSUPPORTED_FEATURE",
                "NO_READERS_AVAILABLE",
                "SERVICE_STOPPED",
                "NO_SERVICE",
                "COMM_ERROR",
                "INTERNAL_ERROR",
                "UNKNOWN_ERROR",
                "SERVER_TOO_BUSY",
                "UNEXPECTED",
                "SHUTDOWN",
                "UNKNOWN",
              ][resultCode - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReportGetStatusChangeResult": (): heap.Ref<boolean> => {
      if (WEBEXT?.smartCardProviderPrivate && "reportGetStatusChangeResult" in WEBEXT?.smartCardProviderPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportGetStatusChangeResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.reportGetStatusChangeResult);
    },
    "call_ReportGetStatusChangeResult": (
      retPtr: Pointer,
      requestId: number,
      readerStates: heap.Ref<object>,
      resultCode: number
    ): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.reportGetStatusChangeResult(
        requestId,
        A.H.get<object>(readerStates),
        resultCode > 0 && resultCode <= 32
          ? [
              "SUCCESS",
              "REMOVED_CARD",
              "RESET_CARD",
              "UNPOWERED_CARD",
              "UNRESPONSIVE_CARD",
              "UNSUPPORTED_CARD",
              "READER_UNAVAILABLE",
              "SHARING_VIOLATION",
              "NOT_TRANSACTED",
              "NO_SMARTCARD",
              "PROTO_MISMATCH",
              "SYSTEM_CANCELLED",
              "NOT_READY",
              "CANCELLED",
              "INSUFFICIENT_BUFFER",
              "INVALID_HANDLE",
              "INVALID_PARAMETER",
              "INVALID_VALUE",
              "NO_MEMORY",
              "TIMEOUT",
              "UNKNOWN_READER",
              "UNSUPPORTED_FEATURE",
              "NO_READERS_AVAILABLE",
              "SERVICE_STOPPED",
              "NO_SERVICE",
              "COMM_ERROR",
              "INTERNAL_ERROR",
              "UNKNOWN_ERROR",
              "SERVER_TOO_BUSY",
              "UNEXPECTED",
              "SHUTDOWN",
              "UNKNOWN",
            ][resultCode - 1]
          : undefined
      );
    },
    "try_ReportGetStatusChangeResult": (
      retPtr: Pointer,
      errPtr: Pointer,
      requestId: number,
      readerStates: heap.Ref<object>,
      resultCode: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.reportGetStatusChangeResult(
          requestId,
          A.H.get<object>(readerStates),
          resultCode > 0 && resultCode <= 32
            ? [
                "SUCCESS",
                "REMOVED_CARD",
                "RESET_CARD",
                "UNPOWERED_CARD",
                "UNRESPONSIVE_CARD",
                "UNSUPPORTED_CARD",
                "READER_UNAVAILABLE",
                "SHARING_VIOLATION",
                "NOT_TRANSACTED",
                "NO_SMARTCARD",
                "PROTO_MISMATCH",
                "SYSTEM_CANCELLED",
                "NOT_READY",
                "CANCELLED",
                "INSUFFICIENT_BUFFER",
                "INVALID_HANDLE",
                "INVALID_PARAMETER",
                "INVALID_VALUE",
                "NO_MEMORY",
                "TIMEOUT",
                "UNKNOWN_READER",
                "UNSUPPORTED_FEATURE",
                "NO_READERS_AVAILABLE",
                "SERVICE_STOPPED",
                "NO_SERVICE",
                "COMM_ERROR",
                "INTERNAL_ERROR",
                "UNKNOWN_ERROR",
                "SERVER_TOO_BUSY",
                "UNEXPECTED",
                "SHUTDOWN",
                "UNKNOWN",
              ][resultCode - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReportListReadersResult": (): heap.Ref<boolean> => {
      if (WEBEXT?.smartCardProviderPrivate && "reportListReadersResult" in WEBEXT?.smartCardProviderPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportListReadersResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.reportListReadersResult);
    },
    "call_ReportListReadersResult": (
      retPtr: Pointer,
      requestId: number,
      readers: heap.Ref<object>,
      resultCode: number
    ): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.reportListReadersResult(
        requestId,
        A.H.get<object>(readers),
        resultCode > 0 && resultCode <= 32
          ? [
              "SUCCESS",
              "REMOVED_CARD",
              "RESET_CARD",
              "UNPOWERED_CARD",
              "UNRESPONSIVE_CARD",
              "UNSUPPORTED_CARD",
              "READER_UNAVAILABLE",
              "SHARING_VIOLATION",
              "NOT_TRANSACTED",
              "NO_SMARTCARD",
              "PROTO_MISMATCH",
              "SYSTEM_CANCELLED",
              "NOT_READY",
              "CANCELLED",
              "INSUFFICIENT_BUFFER",
              "INVALID_HANDLE",
              "INVALID_PARAMETER",
              "INVALID_VALUE",
              "NO_MEMORY",
              "TIMEOUT",
              "UNKNOWN_READER",
              "UNSUPPORTED_FEATURE",
              "NO_READERS_AVAILABLE",
              "SERVICE_STOPPED",
              "NO_SERVICE",
              "COMM_ERROR",
              "INTERNAL_ERROR",
              "UNKNOWN_ERROR",
              "SERVER_TOO_BUSY",
              "UNEXPECTED",
              "SHUTDOWN",
              "UNKNOWN",
            ][resultCode - 1]
          : undefined
      );
    },
    "try_ReportListReadersResult": (
      retPtr: Pointer,
      errPtr: Pointer,
      requestId: number,
      readers: heap.Ref<object>,
      resultCode: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.reportListReadersResult(
          requestId,
          A.H.get<object>(readers),
          resultCode > 0 && resultCode <= 32
            ? [
                "SUCCESS",
                "REMOVED_CARD",
                "RESET_CARD",
                "UNPOWERED_CARD",
                "UNRESPONSIVE_CARD",
                "UNSUPPORTED_CARD",
                "READER_UNAVAILABLE",
                "SHARING_VIOLATION",
                "NOT_TRANSACTED",
                "NO_SMARTCARD",
                "PROTO_MISMATCH",
                "SYSTEM_CANCELLED",
                "NOT_READY",
                "CANCELLED",
                "INSUFFICIENT_BUFFER",
                "INVALID_HANDLE",
                "INVALID_PARAMETER",
                "INVALID_VALUE",
                "NO_MEMORY",
                "TIMEOUT",
                "UNKNOWN_READER",
                "UNSUPPORTED_FEATURE",
                "NO_READERS_AVAILABLE",
                "SERVICE_STOPPED",
                "NO_SERVICE",
                "COMM_ERROR",
                "INTERNAL_ERROR",
                "UNKNOWN_ERROR",
                "SERVER_TOO_BUSY",
                "UNEXPECTED",
                "SHUTDOWN",
                "UNKNOWN",
              ][resultCode - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReportPlainResult": (): heap.Ref<boolean> => {
      if (WEBEXT?.smartCardProviderPrivate && "reportPlainResult" in WEBEXT?.smartCardProviderPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportPlainResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.reportPlainResult);
    },
    "call_ReportPlainResult": (retPtr: Pointer, requestId: number, resultCode: number): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.reportPlainResult(
        requestId,
        resultCode > 0 && resultCode <= 32
          ? [
              "SUCCESS",
              "REMOVED_CARD",
              "RESET_CARD",
              "UNPOWERED_CARD",
              "UNRESPONSIVE_CARD",
              "UNSUPPORTED_CARD",
              "READER_UNAVAILABLE",
              "SHARING_VIOLATION",
              "NOT_TRANSACTED",
              "NO_SMARTCARD",
              "PROTO_MISMATCH",
              "SYSTEM_CANCELLED",
              "NOT_READY",
              "CANCELLED",
              "INSUFFICIENT_BUFFER",
              "INVALID_HANDLE",
              "INVALID_PARAMETER",
              "INVALID_VALUE",
              "NO_MEMORY",
              "TIMEOUT",
              "UNKNOWN_READER",
              "UNSUPPORTED_FEATURE",
              "NO_READERS_AVAILABLE",
              "SERVICE_STOPPED",
              "NO_SERVICE",
              "COMM_ERROR",
              "INTERNAL_ERROR",
              "UNKNOWN_ERROR",
              "SERVER_TOO_BUSY",
              "UNEXPECTED",
              "SHUTDOWN",
              "UNKNOWN",
            ][resultCode - 1]
          : undefined
      );
    },
    "try_ReportPlainResult": (
      retPtr: Pointer,
      errPtr: Pointer,
      requestId: number,
      resultCode: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.reportPlainResult(
          requestId,
          resultCode > 0 && resultCode <= 32
            ? [
                "SUCCESS",
                "REMOVED_CARD",
                "RESET_CARD",
                "UNPOWERED_CARD",
                "UNRESPONSIVE_CARD",
                "UNSUPPORTED_CARD",
                "READER_UNAVAILABLE",
                "SHARING_VIOLATION",
                "NOT_TRANSACTED",
                "NO_SMARTCARD",
                "PROTO_MISMATCH",
                "SYSTEM_CANCELLED",
                "NOT_READY",
                "CANCELLED",
                "INSUFFICIENT_BUFFER",
                "INVALID_HANDLE",
                "INVALID_PARAMETER",
                "INVALID_VALUE",
                "NO_MEMORY",
                "TIMEOUT",
                "UNKNOWN_READER",
                "UNSUPPORTED_FEATURE",
                "NO_READERS_AVAILABLE",
                "SERVICE_STOPPED",
                "NO_SERVICE",
                "COMM_ERROR",
                "INTERNAL_ERROR",
                "UNKNOWN_ERROR",
                "SERVER_TOO_BUSY",
                "UNEXPECTED",
                "SHUTDOWN",
                "UNKNOWN",
              ][resultCode - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReportReleaseContextResult": (): heap.Ref<boolean> => {
      if (WEBEXT?.smartCardProviderPrivate && "reportReleaseContextResult" in WEBEXT?.smartCardProviderPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportReleaseContextResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.reportReleaseContextResult);
    },
    "call_ReportReleaseContextResult": (retPtr: Pointer, requestId: number, resultCode: number): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.reportReleaseContextResult(
        requestId,
        resultCode > 0 && resultCode <= 32
          ? [
              "SUCCESS",
              "REMOVED_CARD",
              "RESET_CARD",
              "UNPOWERED_CARD",
              "UNRESPONSIVE_CARD",
              "UNSUPPORTED_CARD",
              "READER_UNAVAILABLE",
              "SHARING_VIOLATION",
              "NOT_TRANSACTED",
              "NO_SMARTCARD",
              "PROTO_MISMATCH",
              "SYSTEM_CANCELLED",
              "NOT_READY",
              "CANCELLED",
              "INSUFFICIENT_BUFFER",
              "INVALID_HANDLE",
              "INVALID_PARAMETER",
              "INVALID_VALUE",
              "NO_MEMORY",
              "TIMEOUT",
              "UNKNOWN_READER",
              "UNSUPPORTED_FEATURE",
              "NO_READERS_AVAILABLE",
              "SERVICE_STOPPED",
              "NO_SERVICE",
              "COMM_ERROR",
              "INTERNAL_ERROR",
              "UNKNOWN_ERROR",
              "SERVER_TOO_BUSY",
              "UNEXPECTED",
              "SHUTDOWN",
              "UNKNOWN",
            ][resultCode - 1]
          : undefined
      );
    },
    "try_ReportReleaseContextResult": (
      retPtr: Pointer,
      errPtr: Pointer,
      requestId: number,
      resultCode: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.reportReleaseContextResult(
          requestId,
          resultCode > 0 && resultCode <= 32
            ? [
                "SUCCESS",
                "REMOVED_CARD",
                "RESET_CARD",
                "UNPOWERED_CARD",
                "UNRESPONSIVE_CARD",
                "UNSUPPORTED_CARD",
                "READER_UNAVAILABLE",
                "SHARING_VIOLATION",
                "NOT_TRANSACTED",
                "NO_SMARTCARD",
                "PROTO_MISMATCH",
                "SYSTEM_CANCELLED",
                "NOT_READY",
                "CANCELLED",
                "INSUFFICIENT_BUFFER",
                "INVALID_HANDLE",
                "INVALID_PARAMETER",
                "INVALID_VALUE",
                "NO_MEMORY",
                "TIMEOUT",
                "UNKNOWN_READER",
                "UNSUPPORTED_FEATURE",
                "NO_READERS_AVAILABLE",
                "SERVICE_STOPPED",
                "NO_SERVICE",
                "COMM_ERROR",
                "INTERNAL_ERROR",
                "UNKNOWN_ERROR",
                "SERVER_TOO_BUSY",
                "UNEXPECTED",
                "SHUTDOWN",
                "UNKNOWN",
              ][resultCode - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReportStatusResult": (): heap.Ref<boolean> => {
      if (WEBEXT?.smartCardProviderPrivate && "reportStatusResult" in WEBEXT?.smartCardProviderPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportStatusResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.smartCardProviderPrivate.reportStatusResult);
    },
    "call_ReportStatusResult": (
      retPtr: Pointer,
      requestId: number,
      readerName: heap.Ref<object>,
      state: number,
      protocol: number,
      atr: heap.Ref<object>,
      resultCode: number
    ): void => {
      const _ret = WEBEXT.smartCardProviderPrivate.reportStatusResult(
        requestId,
        A.H.get<object>(readerName),
        state > 0 && state <= 6
          ? ["ABSENT", "PRESENT", "SWALLOWED", "POWERED", "NEGOTIABLE", "SPECIFIC"][state - 1]
          : undefined,
        protocol > 0 && protocol <= 4 ? ["UNDEFINED", "T0", "T1", "RAW"][protocol - 1] : undefined,
        A.H.get<object>(atr),
        resultCode > 0 && resultCode <= 32
          ? [
              "SUCCESS",
              "REMOVED_CARD",
              "RESET_CARD",
              "UNPOWERED_CARD",
              "UNRESPONSIVE_CARD",
              "UNSUPPORTED_CARD",
              "READER_UNAVAILABLE",
              "SHARING_VIOLATION",
              "NOT_TRANSACTED",
              "NO_SMARTCARD",
              "PROTO_MISMATCH",
              "SYSTEM_CANCELLED",
              "NOT_READY",
              "CANCELLED",
              "INSUFFICIENT_BUFFER",
              "INVALID_HANDLE",
              "INVALID_PARAMETER",
              "INVALID_VALUE",
              "NO_MEMORY",
              "TIMEOUT",
              "UNKNOWN_READER",
              "UNSUPPORTED_FEATURE",
              "NO_READERS_AVAILABLE",
              "SERVICE_STOPPED",
              "NO_SERVICE",
              "COMM_ERROR",
              "INTERNAL_ERROR",
              "UNKNOWN_ERROR",
              "SERVER_TOO_BUSY",
              "UNEXPECTED",
              "SHUTDOWN",
              "UNKNOWN",
            ][resultCode - 1]
          : undefined
      );
    },
    "try_ReportStatusResult": (
      retPtr: Pointer,
      errPtr: Pointer,
      requestId: number,
      readerName: heap.Ref<object>,
      state: number,
      protocol: number,
      atr: heap.Ref<object>,
      resultCode: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.smartCardProviderPrivate.reportStatusResult(
          requestId,
          A.H.get<object>(readerName),
          state > 0 && state <= 6
            ? ["ABSENT", "PRESENT", "SWALLOWED", "POWERED", "NEGOTIABLE", "SPECIFIC"][state - 1]
            : undefined,
          protocol > 0 && protocol <= 4 ? ["UNDEFINED", "T0", "T1", "RAW"][protocol - 1] : undefined,
          A.H.get<object>(atr),
          resultCode > 0 && resultCode <= 32
            ? [
                "SUCCESS",
                "REMOVED_CARD",
                "RESET_CARD",
                "UNPOWERED_CARD",
                "UNRESPONSIVE_CARD",
                "UNSUPPORTED_CARD",
                "READER_UNAVAILABLE",
                "SHARING_VIOLATION",
                "NOT_TRANSACTED",
                "NO_SMARTCARD",
                "PROTO_MISMATCH",
                "SYSTEM_CANCELLED",
                "NOT_READY",
                "CANCELLED",
                "INSUFFICIENT_BUFFER",
                "INVALID_HANDLE",
                "INVALID_PARAMETER",
                "INVALID_VALUE",
                "NO_MEMORY",
                "TIMEOUT",
                "UNKNOWN_READER",
                "UNSUPPORTED_FEATURE",
                "NO_READERS_AVAILABLE",
                "SERVICE_STOPPED",
                "NO_SERVICE",
                "COMM_ERROR",
                "INTERNAL_ERROR",
                "UNKNOWN_ERROR",
                "SERVER_TOO_BUSY",
                "UNEXPECTED",
                "SHUTDOWN",
                "UNKNOWN",
              ][resultCode - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
