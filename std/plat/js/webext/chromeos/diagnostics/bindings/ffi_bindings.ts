import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/chromeos/diagnostics", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_AcPowerStatus": (ref: heap.Ref<string>): number => {
      const idx = ["connected", "disconnected"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_CancelRoutineRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["uuid"]);
      }
    },
    "load_CancelRoutineRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["uuid"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CreateMemoryRoutineResponse": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["uuid"]);
      }
    },
    "load_CreateMemoryRoutineResponse": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["uuid"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DiskReadRoutineType": (ref: heap.Ref<string>): number => {
      const idx = ["linear", "random"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ExceptionReason": (ref: heap.Ref<string>): number => {
      const idx = ["unknown", "unexpected", "unsupported", "app_ui_closed"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ExceptionInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["uuid"]);
        A.store.Enum(ptr + 4, ["unknown", "unexpected", "unsupported", "app_ui_closed"].indexOf(x["reason"] as string));
        A.store.Ref(ptr + 8, x["debugMessage"]);
      }
    },
    "load_ExceptionInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["uuid"] = A.load.Ref(ptr + 0, undefined);
      x["reason"] = A.load.Enum(ptr + 4, ["unknown", "unexpected", "unsupported", "app_ui_closed"]);
      x["debugMessage"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RoutineType": (ref: heap.Ref<string>): number => {
      const idx = [
        "ac_power",
        "battery_capacity",
        "battery_charge",
        "battery_discharge",
        "battery_health",
        "cpu_cache",
        "cpu_floating_point_accuracy",
        "cpu_prime_search",
        "cpu_stress",
        "disk_read",
        "dns_resolution",
        "memory",
        "nvme_wear_level",
        "smartctl_check",
        "lan_connectivity",
        "signal_strength",
        "dns_resolver_present",
        "gateway_can_be_pinged",
        "sensitive_sensor",
        "nvme_self_test",
        "fingerprint_alive",
        "smartctl_check_with_percentage_used",
        "emmc_lifetime",
        "bluetooth_power",
        "ufs_lifetime",
        "power_button",
        "audio_driver",
        "bluetooth_discovery",
        "bluetooth_scanning",
        "bluetooth_pairing",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_GetAvailableRoutinesResponse": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["routines"]);
      }
    },
    "load_GetAvailableRoutinesResponse": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["routines"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RoutineStatus": (ref: heap.Ref<string>): number => {
      const idx = [
        "unknown",
        "ready",
        "running",
        "waiting_user_action",
        "passed",
        "failed",
        "error",
        "cancelled",
        "failed_to_start",
        "removed",
        "cancelling",
        "unsupported",
        "not_run",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_UserMessageType": (ref: heap.Ref<string>): number => {
      const idx = ["unknown", "unplug_ac_power", "plug_in_ac_power", "press_power_button"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_GetRoutineUpdateResponse": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 20, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Ref(ptr + 12, undefined);
        A.store.Enum(ptr + 16, -1);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Bool(ptr + 20, "progress_percent" in x ? true : false);
        A.store.Int32(ptr + 0, x["progress_percent"] === undefined ? 0 : (x["progress_percent"] as number));
        A.store.Ref(ptr + 4, x["output"]);
        A.store.Enum(
          ptr + 8,
          [
            "unknown",
            "ready",
            "running",
            "waiting_user_action",
            "passed",
            "failed",
            "error",
            "cancelled",
            "failed_to_start",
            "removed",
            "cancelling",
            "unsupported",
            "not_run",
          ].indexOf(x["status"] as string)
        );
        A.store.Ref(ptr + 12, x["status_message"]);
        A.store.Enum(
          ptr + 16,
          ["unknown", "unplug_ac_power", "plug_in_ac_power", "press_power_button"].indexOf(x["user_message"] as string)
        );
      }
    },
    "load_GetRoutineUpdateResponse": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 20)) {
        x["progress_percent"] = A.load.Int32(ptr + 0);
      } else {
        delete x["progress_percent"];
      }
      x["output"] = A.load.Ref(ptr + 4, undefined);
      x["status"] = A.load.Enum(ptr + 8, [
        "unknown",
        "ready",
        "running",
        "waiting_user_action",
        "passed",
        "failed",
        "error",
        "cancelled",
        "failed_to_start",
        "removed",
        "cancelling",
        "unsupported",
        "not_run",
      ]);
      x["status_message"] = A.load.Ref(ptr + 12, undefined);
      x["user_message"] = A.load.Enum(ptr + 16, [
        "unknown",
        "unplug_ac_power",
        "plug_in_ac_power",
        "press_power_button",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RoutineCommandType": (ref: heap.Ref<string>): number => {
      const idx = ["cancel", "remove", "resume", "status"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_GetRoutineUpdateRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "id" in x ? true : false);
        A.store.Int32(ptr + 0, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Enum(ptr + 4, ["cancel", "remove", "resume", "status"].indexOf(x["command"] as string));
      }
    },
    "load_GetRoutineUpdateRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["id"] = A.load.Int32(ptr + 0);
      } else {
        delete x["id"];
      }
      x["command"] = A.load.Enum(ptr + 4, ["cancel", "remove", "resume", "status"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_MemtesterTestItemEnum": (ref: heap.Ref<string>): number => {
      const idx = [
        "unknown",
        "stuck_address",
        "compare_and",
        "compare_div",
        "compare_mul",
        "compare_or",
        "compare_sub",
        "compare_xor",
        "sequential_increment",
        "bit_flip",
        "bit_spread",
        "block_sequential",
        "checkerboard",
        "random_value",
        "solid_bits",
        "walking_ones",
        "walking_zeroes",
        "eight_bit_writes",
        "sixteen_bit_writes",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_MemtesterResult": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["passed_items"]);
        A.store.Ref(ptr + 4, x["failed_items"]);
      }
    },
    "load_MemtesterResult": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["passed_items"] = A.load.Ref(ptr + 0, undefined);
      x["failed_items"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MemoryRoutineFinishedInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 27, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 26, false);
        A.store.Float64(ptr + 8, 0);

        A.store.Bool(ptr + 16 + 8, false);
        A.store.Ref(ptr + 16 + 0, undefined);
        A.store.Ref(ptr + 16 + 4, undefined);
      } else {
        A.store.Bool(ptr + 27, true);
        A.store.Ref(ptr + 0, x["uuid"]);
        A.store.Bool(ptr + 25, "has_passed" in x ? true : false);
        A.store.Bool(ptr + 4, x["has_passed"] ? true : false);
        A.store.Bool(ptr + 26, "bytesTested" in x ? true : false);
        A.store.Float64(ptr + 8, x["bytesTested"] === undefined ? 0 : (x["bytesTested"] as number));

        if (typeof x["result"] === "undefined") {
          A.store.Bool(ptr + 16 + 8, false);
          A.store.Ref(ptr + 16 + 0, undefined);
          A.store.Ref(ptr + 16 + 4, undefined);
        } else {
          A.store.Bool(ptr + 16 + 8, true);
          A.store.Ref(ptr + 16 + 0, x["result"]["passed_items"]);
          A.store.Ref(ptr + 16 + 4, x["result"]["failed_items"]);
        }
      }
    },
    "load_MemoryRoutineFinishedInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["uuid"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 25)) {
        x["has_passed"] = A.load.Bool(ptr + 4);
      } else {
        delete x["has_passed"];
      }
      if (A.load.Bool(ptr + 26)) {
        x["bytesTested"] = A.load.Float64(ptr + 8);
      } else {
        delete x["bytesTested"];
      }
      if (A.load.Bool(ptr + 16 + 8)) {
        x["result"] = {};
        x["result"]["passed_items"] = A.load.Ref(ptr + 16 + 0, undefined);
        x["result"]["failed_items"] = A.load.Ref(ptr + 16 + 4, undefined);
      } else {
        delete x["result"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_NvmeSelfTestType": (ref: heap.Ref<string>): number => {
      const idx = ["short_test", "long_test"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RoutineInitializedInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["uuid"]);
      }
    },
    "load_RoutineInitializedInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["uuid"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RoutineRunningInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Ref(ptr + 0, x["uuid"]);
        A.store.Bool(ptr + 8, "percentage" in x ? true : false);
        A.store.Int32(ptr + 4, x["percentage"] === undefined ? 0 : (x["percentage"] as number));
      }
    },
    "load_RoutineRunningInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["uuid"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 8)) {
        x["percentage"] = A.load.Int32(ptr + 4);
      } else {
        delete x["percentage"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RoutineSupportStatus": (ref: heap.Ref<string>): number => {
      const idx = ["supported", "unsupported"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RoutineSupportStatusInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["supported", "unsupported"].indexOf(x["status"] as string));
      }
    },
    "load_RoutineSupportStatusInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["status"] = A.load.Enum(ptr + 0, ["supported", "unsupported"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RoutineWaitingReason": (ref: heap.Ref<string>): number => {
      const idx = ["waiting_to_be_scheduled", "waiting_user_input"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_RoutineWaitingInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Enum(ptr + 8, -1);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["uuid"]);
        A.store.Bool(ptr + 16, "percentage" in x ? true : false);
        A.store.Int32(ptr + 4, x["percentage"] === undefined ? 0 : (x["percentage"] as number));
        A.store.Enum(ptr + 8, ["waiting_to_be_scheduled", "waiting_user_input"].indexOf(x["reason"] as string));
        A.store.Ref(ptr + 12, x["message"]);
      }
    },
    "load_RoutineWaitingInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["uuid"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["percentage"] = A.load.Int32(ptr + 4);
      } else {
        delete x["percentage"];
      }
      x["reason"] = A.load.Enum(ptr + 8, ["waiting_to_be_scheduled", "waiting_user_input"]);
      x["message"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RunAcPowerRoutineRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["connected", "disconnected"].indexOf(x["expected_status"] as string));
        A.store.Ref(ptr + 4, x["expected_power_type"]);
      }
    },
    "load_RunAcPowerRoutineRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["expected_status"] = A.load.Enum(ptr + 0, ["connected", "disconnected"]);
      x["expected_power_type"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RunBatteryChargeRoutineRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "length_seconds" in x ? true : false);
        A.store.Int32(ptr + 0, x["length_seconds"] === undefined ? 0 : (x["length_seconds"] as number));
        A.store.Bool(ptr + 9, "minimum_charge_percent_required" in x ? true : false);
        A.store.Int32(
          ptr + 4,
          x["minimum_charge_percent_required"] === undefined ? 0 : (x["minimum_charge_percent_required"] as number)
        );
      }
    },
    "load_RunBatteryChargeRoutineRequest": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["length_seconds"] = A.load.Int32(ptr + 0);
      } else {
        delete x["length_seconds"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["minimum_charge_percent_required"] = A.load.Int32(ptr + 4);
      } else {
        delete x["minimum_charge_percent_required"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RunBatteryDischargeRoutineRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "length_seconds" in x ? true : false);
        A.store.Int32(ptr + 0, x["length_seconds"] === undefined ? 0 : (x["length_seconds"] as number));
        A.store.Bool(ptr + 9, "maximum_discharge_percent_allowed" in x ? true : false);
        A.store.Int32(
          ptr + 4,
          x["maximum_discharge_percent_allowed"] === undefined ? 0 : (x["maximum_discharge_percent_allowed"] as number)
        );
      }
    },
    "load_RunBatteryDischargeRoutineRequest": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["length_seconds"] = A.load.Int32(ptr + 0);
      } else {
        delete x["length_seconds"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["maximum_discharge_percent_allowed"] = A.load.Int32(ptr + 4);
      } else {
        delete x["maximum_discharge_percent_allowed"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RunBluetoothPairingRoutineRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["peripheral_id"]);
      }
    },
    "load_RunBluetoothPairingRoutineRequest": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["peripheral_id"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RunBluetoothScanningRoutineRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "length_seconds" in x ? true : false);
        A.store.Int32(ptr + 0, x["length_seconds"] === undefined ? 0 : (x["length_seconds"] as number));
      }
    },
    "load_RunBluetoothScanningRoutineRequest": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["length_seconds"] = A.load.Int32(ptr + 0);
      } else {
        delete x["length_seconds"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RunCpuRoutineRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "length_seconds" in x ? true : false);
        A.store.Int32(ptr + 0, x["length_seconds"] === undefined ? 0 : (x["length_seconds"] as number));
      }
    },
    "load_RunCpuRoutineRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["length_seconds"] = A.load.Int32(ptr + 0);
      } else {
        delete x["length_seconds"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RunDiskReadRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Enum(ptr + 0, ["linear", "random"].indexOf(x["type"] as string));
        A.store.Bool(ptr + 12, "length_seconds" in x ? true : false);
        A.store.Int32(ptr + 4, x["length_seconds"] === undefined ? 0 : (x["length_seconds"] as number));
        A.store.Bool(ptr + 13, "file_size_mb" in x ? true : false);
        A.store.Int32(ptr + 8, x["file_size_mb"] === undefined ? 0 : (x["file_size_mb"] as number));
      }
    },
    "load_RunDiskReadRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, ["linear", "random"]);
      if (A.load.Bool(ptr + 12)) {
        x["length_seconds"] = A.load.Int32(ptr + 4);
      } else {
        delete x["length_seconds"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["file_size_mb"] = A.load.Int32(ptr + 8);
      } else {
        delete x["file_size_mb"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RunMemoryRoutineArguments": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "maxTestingMemKib" in x ? true : false);
        A.store.Int32(ptr + 0, x["maxTestingMemKib"] === undefined ? 0 : (x["maxTestingMemKib"] as number));
      }
    },
    "load_RunMemoryRoutineArguments": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["maxTestingMemKib"] = A.load.Int32(ptr + 0);
      } else {
        delete x["maxTestingMemKib"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RunNvmeSelfTestRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["short_test", "long_test"].indexOf(x["test_type"] as string));
      }
    },
    "load_RunNvmeSelfTestRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["test_type"] = A.load.Enum(ptr + 0, ["short_test", "long_test"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RunNvmeWearLevelRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "wear_level_threshold" in x ? true : false);
        A.store.Int32(ptr + 0, x["wear_level_threshold"] === undefined ? 0 : (x["wear_level_threshold"] as number));
      }
    },
    "load_RunNvmeWearLevelRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["wear_level_threshold"] = A.load.Int32(ptr + 0);
      } else {
        delete x["wear_level_threshold"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RunPowerButtonRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "timeout_seconds" in x ? true : false);
        A.store.Int32(ptr + 0, x["timeout_seconds"] === undefined ? 0 : (x["timeout_seconds"] as number));
      }
    },
    "load_RunPowerButtonRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["timeout_seconds"] = A.load.Int32(ptr + 0);
      } else {
        delete x["timeout_seconds"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RunRoutineResponse": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "id" in x ? true : false);
        A.store.Int32(ptr + 0, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Enum(
          ptr + 4,
          [
            "unknown",
            "ready",
            "running",
            "waiting_user_action",
            "passed",
            "failed",
            "error",
            "cancelled",
            "failed_to_start",
            "removed",
            "cancelling",
            "unsupported",
            "not_run",
          ].indexOf(x["status"] as string)
        );
      }
    },
    "load_RunRoutineResponse": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["id"] = A.load.Int32(ptr + 0);
      } else {
        delete x["id"];
      }
      x["status"] = A.load.Enum(ptr + 4, [
        "unknown",
        "ready",
        "running",
        "waiting_user_action",
        "passed",
        "failed",
        "error",
        "cancelled",
        "failed_to_start",
        "removed",
        "cancelling",
        "unsupported",
        "not_run",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RunSmartctlCheckRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "percentage_used_threshold" in x ? true : false);
        A.store.Int32(
          ptr + 0,
          x["percentage_used_threshold"] === undefined ? 0 : (x["percentage_used_threshold"] as number)
        );
      }
    },
    "load_RunSmartctlCheckRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["percentage_used_threshold"] = A.load.Int32(ptr + 0);
      } else {
        delete x["percentage_used_threshold"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_StartRoutineRequest": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["uuid"]);
      }
    },
    "load_StartRoutineRequest": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["uuid"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_CancelRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "cancelRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CancelRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.cancelRoutine);
    },
    "call_CancelRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      request_ffi["uuid"] = A.load.Ref(request + 0, undefined);

      const _ret = WEBEXT.os.diagnostics.cancelRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_CancelRoutine": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        request_ffi["uuid"] = A.load.Ref(request + 0, undefined);

        const _ret = WEBEXT.os.diagnostics.cancelRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CreateMemoryRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "createMemoryRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CreateMemoryRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.createMemoryRoutine);
    },
    "call_CreateMemoryRoutine": (retPtr: Pointer, args: Pointer): void => {
      const args_ffi = {};

      if (A.load.Bool(args + 4)) {
        args_ffi["maxTestingMemKib"] = A.load.Int32(args + 0);
      }

      const _ret = WEBEXT.os.diagnostics.createMemoryRoutine(args_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_CreateMemoryRoutine": (retPtr: Pointer, errPtr: Pointer, args: Pointer): heap.Ref<boolean> => {
      try {
        const args_ffi = {};

        if (A.load.Bool(args + 4)) {
          args_ffi["maxTestingMemKib"] = A.load.Int32(args + 0);
        }

        const _ret = WEBEXT.os.diagnostics.createMemoryRoutine(args_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAvailableRoutines": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "getAvailableRoutines" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAvailableRoutines": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.getAvailableRoutines);
    },
    "call_GetAvailableRoutines": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.getAvailableRoutines();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAvailableRoutines": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.getAvailableRoutines();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetRoutineUpdate": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "getRoutineUpdate" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetRoutineUpdate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.getRoutineUpdate);
    },
    "call_GetRoutineUpdate": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 8)) {
        request_ffi["id"] = A.load.Int32(request + 0);
      }
      request_ffi["command"] = A.load.Enum(request + 4, ["cancel", "remove", "resume", "status"]);

      const _ret = WEBEXT.os.diagnostics.getRoutineUpdate(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetRoutineUpdate": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 8)) {
          request_ffi["id"] = A.load.Int32(request + 0);
        }
        request_ffi["command"] = A.load.Enum(request + 4, ["cancel", "remove", "resume", "status"]);

        const _ret = WEBEXT.os.diagnostics.getRoutineUpdate(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsMemoryRoutineArgumentSupported": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "isMemoryRoutineArgumentSupported" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsMemoryRoutineArgumentSupported": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.isMemoryRoutineArgumentSupported);
    },
    "call_IsMemoryRoutineArgumentSupported": (retPtr: Pointer, args: Pointer): void => {
      const args_ffi = {};

      if (A.load.Bool(args + 4)) {
        args_ffi["maxTestingMemKib"] = A.load.Int32(args + 0);
      }

      const _ret = WEBEXT.os.diagnostics.isMemoryRoutineArgumentSupported(args_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_IsMemoryRoutineArgumentSupported": (retPtr: Pointer, errPtr: Pointer, args: Pointer): heap.Ref<boolean> => {
      try {
        const args_ffi = {};

        if (A.load.Bool(args + 4)) {
          args_ffi["maxTestingMemKib"] = A.load.Int32(args + 0);
        }

        const _ret = WEBEXT.os.diagnostics.isMemoryRoutineArgumentSupported(args_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMemoryRoutineFinished": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.diagnostics?.onMemoryRoutineFinished &&
        "addListener" in WEBEXT?.os?.diagnostics?.onMemoryRoutineFinished
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMemoryRoutineFinished": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.onMemoryRoutineFinished.addListener);
    },
    "call_OnMemoryRoutineFinished": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.diagnostics.onMemoryRoutineFinished.addListener(A.H.get<object>(callback));
    },
    "try_OnMemoryRoutineFinished": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.onMemoryRoutineFinished.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMemoryRoutineFinished": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.diagnostics?.onMemoryRoutineFinished &&
        "removeListener" in WEBEXT?.os?.diagnostics?.onMemoryRoutineFinished
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMemoryRoutineFinished": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.onMemoryRoutineFinished.removeListener);
    },
    "call_OffMemoryRoutineFinished": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.diagnostics.onMemoryRoutineFinished.removeListener(A.H.get<object>(callback));
    },
    "try_OffMemoryRoutineFinished": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.onMemoryRoutineFinished.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMemoryRoutineFinished": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.diagnostics?.onMemoryRoutineFinished &&
        "hasListener" in WEBEXT?.os?.diagnostics?.onMemoryRoutineFinished
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMemoryRoutineFinished": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.onMemoryRoutineFinished.hasListener);
    },
    "call_HasOnMemoryRoutineFinished": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.diagnostics.onMemoryRoutineFinished.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMemoryRoutineFinished": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.onMemoryRoutineFinished.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRoutineException": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics?.onRoutineException && "addListener" in WEBEXT?.os?.diagnostics?.onRoutineException) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRoutineException": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.onRoutineException.addListener);
    },
    "call_OnRoutineException": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.diagnostics.onRoutineException.addListener(A.H.get<object>(callback));
    },
    "try_OnRoutineException": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.onRoutineException.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRoutineException": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.diagnostics?.onRoutineException &&
        "removeListener" in WEBEXT?.os?.diagnostics?.onRoutineException
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRoutineException": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.onRoutineException.removeListener);
    },
    "call_OffRoutineException": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.diagnostics.onRoutineException.removeListener(A.H.get<object>(callback));
    },
    "try_OffRoutineException": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.onRoutineException.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRoutineException": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics?.onRoutineException && "hasListener" in WEBEXT?.os?.diagnostics?.onRoutineException) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRoutineException": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.onRoutineException.hasListener);
    },
    "call_HasOnRoutineException": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.diagnostics.onRoutineException.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRoutineException": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.onRoutineException.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRoutineInitialized": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.diagnostics?.onRoutineInitialized &&
        "addListener" in WEBEXT?.os?.diagnostics?.onRoutineInitialized
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRoutineInitialized": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.onRoutineInitialized.addListener);
    },
    "call_OnRoutineInitialized": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.diagnostics.onRoutineInitialized.addListener(A.H.get<object>(callback));
    },
    "try_OnRoutineInitialized": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.onRoutineInitialized.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRoutineInitialized": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.diagnostics?.onRoutineInitialized &&
        "removeListener" in WEBEXT?.os?.diagnostics?.onRoutineInitialized
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRoutineInitialized": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.onRoutineInitialized.removeListener);
    },
    "call_OffRoutineInitialized": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.diagnostics.onRoutineInitialized.removeListener(A.H.get<object>(callback));
    },
    "try_OffRoutineInitialized": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.onRoutineInitialized.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRoutineInitialized": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.os?.diagnostics?.onRoutineInitialized &&
        "hasListener" in WEBEXT?.os?.diagnostics?.onRoutineInitialized
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRoutineInitialized": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.onRoutineInitialized.hasListener);
    },
    "call_HasOnRoutineInitialized": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.diagnostics.onRoutineInitialized.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRoutineInitialized": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.onRoutineInitialized.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRoutineRunning": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics?.onRoutineRunning && "addListener" in WEBEXT?.os?.diagnostics?.onRoutineRunning) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRoutineRunning": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.onRoutineRunning.addListener);
    },
    "call_OnRoutineRunning": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.diagnostics.onRoutineRunning.addListener(A.H.get<object>(callback));
    },
    "try_OnRoutineRunning": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.onRoutineRunning.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRoutineRunning": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics?.onRoutineRunning && "removeListener" in WEBEXT?.os?.diagnostics?.onRoutineRunning) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRoutineRunning": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.onRoutineRunning.removeListener);
    },
    "call_OffRoutineRunning": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.diagnostics.onRoutineRunning.removeListener(A.H.get<object>(callback));
    },
    "try_OffRoutineRunning": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.onRoutineRunning.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRoutineRunning": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics?.onRoutineRunning && "hasListener" in WEBEXT?.os?.diagnostics?.onRoutineRunning) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRoutineRunning": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.onRoutineRunning.hasListener);
    },
    "call_HasOnRoutineRunning": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.diagnostics.onRoutineRunning.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRoutineRunning": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.onRoutineRunning.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnRoutineWaiting": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics?.onRoutineWaiting && "addListener" in WEBEXT?.os?.diagnostics?.onRoutineWaiting) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnRoutineWaiting": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.onRoutineWaiting.addListener);
    },
    "call_OnRoutineWaiting": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.diagnostics.onRoutineWaiting.addListener(A.H.get<object>(callback));
    },
    "try_OnRoutineWaiting": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.onRoutineWaiting.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffRoutineWaiting": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics?.onRoutineWaiting && "removeListener" in WEBEXT?.os?.diagnostics?.onRoutineWaiting) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffRoutineWaiting": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.onRoutineWaiting.removeListener);
    },
    "call_OffRoutineWaiting": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.diagnostics.onRoutineWaiting.removeListener(A.H.get<object>(callback));
    },
    "try_OffRoutineWaiting": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.onRoutineWaiting.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnRoutineWaiting": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics?.onRoutineWaiting && "hasListener" in WEBEXT?.os?.diagnostics?.onRoutineWaiting) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnRoutineWaiting": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.onRoutineWaiting.hasListener);
    },
    "call_HasOnRoutineWaiting": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.os.diagnostics.onRoutineWaiting.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnRoutineWaiting": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.onRoutineWaiting.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunAcPowerRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runAcPowerRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunAcPowerRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runAcPowerRoutine);
    },
    "call_RunAcPowerRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      request_ffi["expected_status"] = A.load.Enum(request + 0, ["connected", "disconnected"]);
      request_ffi["expected_power_type"] = A.load.Ref(request + 4, undefined);

      const _ret = WEBEXT.os.diagnostics.runAcPowerRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RunAcPowerRoutine": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        request_ffi["expected_status"] = A.load.Enum(request + 0, ["connected", "disconnected"]);
        request_ffi["expected_power_type"] = A.load.Ref(request + 4, undefined);

        const _ret = WEBEXT.os.diagnostics.runAcPowerRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunAudioDriverRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runAudioDriverRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunAudioDriverRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runAudioDriverRoutine);
    },
    "call_RunAudioDriverRoutine": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.runAudioDriverRoutine();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunAudioDriverRoutine": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.runAudioDriverRoutine();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunBatteryCapacityRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runBatteryCapacityRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunBatteryCapacityRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runBatteryCapacityRoutine);
    },
    "call_RunBatteryCapacityRoutine": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.runBatteryCapacityRoutine();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunBatteryCapacityRoutine": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.runBatteryCapacityRoutine();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunBatteryChargeRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runBatteryChargeRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunBatteryChargeRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runBatteryChargeRoutine);
    },
    "call_RunBatteryChargeRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 8)) {
        request_ffi["length_seconds"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 9)) {
        request_ffi["minimum_charge_percent_required"] = A.load.Int32(request + 4);
      }

      const _ret = WEBEXT.os.diagnostics.runBatteryChargeRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RunBatteryChargeRoutine": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 8)) {
          request_ffi["length_seconds"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 9)) {
          request_ffi["minimum_charge_percent_required"] = A.load.Int32(request + 4);
        }

        const _ret = WEBEXT.os.diagnostics.runBatteryChargeRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunBatteryDischargeRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runBatteryDischargeRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunBatteryDischargeRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runBatteryDischargeRoutine);
    },
    "call_RunBatteryDischargeRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 8)) {
        request_ffi["length_seconds"] = A.load.Int32(request + 0);
      }
      if (A.load.Bool(request + 9)) {
        request_ffi["maximum_discharge_percent_allowed"] = A.load.Int32(request + 4);
      }

      const _ret = WEBEXT.os.diagnostics.runBatteryDischargeRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RunBatteryDischargeRoutine": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 8)) {
          request_ffi["length_seconds"] = A.load.Int32(request + 0);
        }
        if (A.load.Bool(request + 9)) {
          request_ffi["maximum_discharge_percent_allowed"] = A.load.Int32(request + 4);
        }

        const _ret = WEBEXT.os.diagnostics.runBatteryDischargeRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunBatteryHealthRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runBatteryHealthRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunBatteryHealthRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runBatteryHealthRoutine);
    },
    "call_RunBatteryHealthRoutine": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.runBatteryHealthRoutine();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunBatteryHealthRoutine": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.runBatteryHealthRoutine();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunBluetoothDiscoveryRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runBluetoothDiscoveryRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunBluetoothDiscoveryRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runBluetoothDiscoveryRoutine);
    },
    "call_RunBluetoothDiscoveryRoutine": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.runBluetoothDiscoveryRoutine();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunBluetoothDiscoveryRoutine": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.runBluetoothDiscoveryRoutine();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunBluetoothPairingRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runBluetoothPairingRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunBluetoothPairingRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runBluetoothPairingRoutine);
    },
    "call_RunBluetoothPairingRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      request_ffi["peripheral_id"] = A.load.Ref(request + 0, undefined);

      const _ret = WEBEXT.os.diagnostics.runBluetoothPairingRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RunBluetoothPairingRoutine": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        request_ffi["peripheral_id"] = A.load.Ref(request + 0, undefined);

        const _ret = WEBEXT.os.diagnostics.runBluetoothPairingRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunBluetoothPowerRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runBluetoothPowerRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunBluetoothPowerRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runBluetoothPowerRoutine);
    },
    "call_RunBluetoothPowerRoutine": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.runBluetoothPowerRoutine();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunBluetoothPowerRoutine": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.runBluetoothPowerRoutine();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunBluetoothScanningRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runBluetoothScanningRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunBluetoothScanningRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runBluetoothScanningRoutine);
    },
    "call_RunBluetoothScanningRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 4)) {
        request_ffi["length_seconds"] = A.load.Int32(request + 0);
      }

      const _ret = WEBEXT.os.diagnostics.runBluetoothScanningRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RunBluetoothScanningRoutine": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 4)) {
          request_ffi["length_seconds"] = A.load.Int32(request + 0);
        }

        const _ret = WEBEXT.os.diagnostics.runBluetoothScanningRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunCpuCacheRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runCpuCacheRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunCpuCacheRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runCpuCacheRoutine);
    },
    "call_RunCpuCacheRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 4)) {
        request_ffi["length_seconds"] = A.load.Int32(request + 0);
      }

      const _ret = WEBEXT.os.diagnostics.runCpuCacheRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RunCpuCacheRoutine": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 4)) {
          request_ffi["length_seconds"] = A.load.Int32(request + 0);
        }

        const _ret = WEBEXT.os.diagnostics.runCpuCacheRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunCpuFloatingPointAccuracyRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runCpuFloatingPointAccuracyRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunCpuFloatingPointAccuracyRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runCpuFloatingPointAccuracyRoutine);
    },
    "call_RunCpuFloatingPointAccuracyRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 4)) {
        request_ffi["length_seconds"] = A.load.Int32(request + 0);
      }

      const _ret = WEBEXT.os.diagnostics.runCpuFloatingPointAccuracyRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RunCpuFloatingPointAccuracyRoutine": (
      retPtr: Pointer,
      errPtr: Pointer,
      request: Pointer
    ): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 4)) {
          request_ffi["length_seconds"] = A.load.Int32(request + 0);
        }

        const _ret = WEBEXT.os.diagnostics.runCpuFloatingPointAccuracyRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunCpuPrimeSearchRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runCpuPrimeSearchRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunCpuPrimeSearchRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runCpuPrimeSearchRoutine);
    },
    "call_RunCpuPrimeSearchRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 4)) {
        request_ffi["length_seconds"] = A.load.Int32(request + 0);
      }

      const _ret = WEBEXT.os.diagnostics.runCpuPrimeSearchRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RunCpuPrimeSearchRoutine": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 4)) {
          request_ffi["length_seconds"] = A.load.Int32(request + 0);
        }

        const _ret = WEBEXT.os.diagnostics.runCpuPrimeSearchRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunCpuStressRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runCpuStressRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunCpuStressRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runCpuStressRoutine);
    },
    "call_RunCpuStressRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 4)) {
        request_ffi["length_seconds"] = A.load.Int32(request + 0);
      }

      const _ret = WEBEXT.os.diagnostics.runCpuStressRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RunCpuStressRoutine": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 4)) {
          request_ffi["length_seconds"] = A.load.Int32(request + 0);
        }

        const _ret = WEBEXT.os.diagnostics.runCpuStressRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunDiskReadRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runDiskReadRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunDiskReadRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runDiskReadRoutine);
    },
    "call_RunDiskReadRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      request_ffi["type"] = A.load.Enum(request + 0, ["linear", "random"]);
      if (A.load.Bool(request + 12)) {
        request_ffi["length_seconds"] = A.load.Int32(request + 4);
      }
      if (A.load.Bool(request + 13)) {
        request_ffi["file_size_mb"] = A.load.Int32(request + 8);
      }

      const _ret = WEBEXT.os.diagnostics.runDiskReadRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RunDiskReadRoutine": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        request_ffi["type"] = A.load.Enum(request + 0, ["linear", "random"]);
        if (A.load.Bool(request + 12)) {
          request_ffi["length_seconds"] = A.load.Int32(request + 4);
        }
        if (A.load.Bool(request + 13)) {
          request_ffi["file_size_mb"] = A.load.Int32(request + 8);
        }

        const _ret = WEBEXT.os.diagnostics.runDiskReadRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunDnsResolutionRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runDnsResolutionRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunDnsResolutionRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runDnsResolutionRoutine);
    },
    "call_RunDnsResolutionRoutine": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.runDnsResolutionRoutine();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunDnsResolutionRoutine": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.runDnsResolutionRoutine();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunDnsResolverPresentRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runDnsResolverPresentRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunDnsResolverPresentRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runDnsResolverPresentRoutine);
    },
    "call_RunDnsResolverPresentRoutine": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.runDnsResolverPresentRoutine();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunDnsResolverPresentRoutine": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.runDnsResolverPresentRoutine();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunEmmcLifetimeRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runEmmcLifetimeRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunEmmcLifetimeRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runEmmcLifetimeRoutine);
    },
    "call_RunEmmcLifetimeRoutine": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.runEmmcLifetimeRoutine();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunEmmcLifetimeRoutine": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.runEmmcLifetimeRoutine();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunFingerprintAliveRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runFingerprintAliveRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunFingerprintAliveRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runFingerprintAliveRoutine);
    },
    "call_RunFingerprintAliveRoutine": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.runFingerprintAliveRoutine();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunFingerprintAliveRoutine": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.runFingerprintAliveRoutine();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunGatewayCanBePingedRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runGatewayCanBePingedRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunGatewayCanBePingedRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runGatewayCanBePingedRoutine);
    },
    "call_RunGatewayCanBePingedRoutine": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.runGatewayCanBePingedRoutine();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunGatewayCanBePingedRoutine": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.runGatewayCanBePingedRoutine();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunLanConnectivityRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runLanConnectivityRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunLanConnectivityRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runLanConnectivityRoutine);
    },
    "call_RunLanConnectivityRoutine": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.runLanConnectivityRoutine();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunLanConnectivityRoutine": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.runLanConnectivityRoutine();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunMemoryRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runMemoryRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunMemoryRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runMemoryRoutine);
    },
    "call_RunMemoryRoutine": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.runMemoryRoutine();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunMemoryRoutine": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.runMemoryRoutine();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunNvmeSelfTestRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runNvmeSelfTestRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunNvmeSelfTestRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runNvmeSelfTestRoutine);
    },
    "call_RunNvmeSelfTestRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      request_ffi["test_type"] = A.load.Enum(request + 0, ["short_test", "long_test"]);

      const _ret = WEBEXT.os.diagnostics.runNvmeSelfTestRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RunNvmeSelfTestRoutine": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        request_ffi["test_type"] = A.load.Enum(request + 0, ["short_test", "long_test"]);

        const _ret = WEBEXT.os.diagnostics.runNvmeSelfTestRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunNvmeWearLevelRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runNvmeWearLevelRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunNvmeWearLevelRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runNvmeWearLevelRoutine);
    },
    "call_RunNvmeWearLevelRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 4)) {
        request_ffi["wear_level_threshold"] = A.load.Int32(request + 0);
      }

      const _ret = WEBEXT.os.diagnostics.runNvmeWearLevelRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RunNvmeWearLevelRoutine": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 4)) {
          request_ffi["wear_level_threshold"] = A.load.Int32(request + 0);
        }

        const _ret = WEBEXT.os.diagnostics.runNvmeWearLevelRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunPowerButtonRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runPowerButtonRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunPowerButtonRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runPowerButtonRoutine);
    },
    "call_RunPowerButtonRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 4)) {
        request_ffi["timeout_seconds"] = A.load.Int32(request + 0);
      }

      const _ret = WEBEXT.os.diagnostics.runPowerButtonRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RunPowerButtonRoutine": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 4)) {
          request_ffi["timeout_seconds"] = A.load.Int32(request + 0);
        }

        const _ret = WEBEXT.os.diagnostics.runPowerButtonRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunSensitiveSensorRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runSensitiveSensorRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunSensitiveSensorRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runSensitiveSensorRoutine);
    },
    "call_RunSensitiveSensorRoutine": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.runSensitiveSensorRoutine();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunSensitiveSensorRoutine": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.runSensitiveSensorRoutine();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunSignalStrengthRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runSignalStrengthRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunSignalStrengthRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runSignalStrengthRoutine);
    },
    "call_RunSignalStrengthRoutine": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.runSignalStrengthRoutine();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunSignalStrengthRoutine": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.runSignalStrengthRoutine();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunSmartctlCheckRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runSmartctlCheckRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunSmartctlCheckRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runSmartctlCheckRoutine);
    },
    "call_RunSmartctlCheckRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      if (A.load.Bool(request + 4)) {
        request_ffi["percentage_used_threshold"] = A.load.Int32(request + 0);
      }

      const _ret = WEBEXT.os.diagnostics.runSmartctlCheckRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RunSmartctlCheckRoutine": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        if (A.load.Bool(request + 4)) {
          request_ffi["percentage_used_threshold"] = A.load.Int32(request + 0);
        }

        const _ret = WEBEXT.os.diagnostics.runSmartctlCheckRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunUfsLifetimeRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "runUfsLifetimeRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunUfsLifetimeRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.runUfsLifetimeRoutine);
    },
    "call_RunUfsLifetimeRoutine": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.diagnostics.runUfsLifetimeRoutine();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunUfsLifetimeRoutine": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.diagnostics.runUfsLifetimeRoutine();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartRoutine": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.diagnostics && "startRoutine" in WEBEXT?.os?.diagnostics) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartRoutine": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.diagnostics.startRoutine);
    },
    "call_StartRoutine": (retPtr: Pointer, request: Pointer): void => {
      const request_ffi = {};

      request_ffi["uuid"] = A.load.Ref(request + 0, undefined);

      const _ret = WEBEXT.os.diagnostics.startRoutine(request_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_StartRoutine": (retPtr: Pointer, errPtr: Pointer, request: Pointer): heap.Ref<boolean> => {
      try {
        const request_ffi = {};

        request_ffi["uuid"] = A.load.Ref(request + 0, undefined);

        const _ret = WEBEXT.os.diagnostics.startRoutine(request_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
