import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/chromeos/telemetry", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AudioOutputNodeInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 27, false);
        A.store.Bool(ptr + 24, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 26, false);
        A.store.Int32(ptr + 20, 0);
      } else {
        A.store.Bool(ptr + 27, true);
        A.store.Bool(ptr + 24, "id" in x ? true : false);
        A.store.Float64(ptr + 0, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Ref(ptr + 8, x["name"]);
        A.store.Ref(ptr + 12, x["deviceName"]);
        A.store.Bool(ptr + 25, "active" in x ? true : false);
        A.store.Bool(ptr + 16, x["active"] ? true : false);
        A.store.Bool(ptr + 26, "nodeVolume" in x ? true : false);
        A.store.Int32(ptr + 20, x["nodeVolume"] === undefined ? 0 : (x["nodeVolume"] as number));
      }
    },
    "load_AudioOutputNodeInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["id"] = A.load.Float64(ptr + 0);
      } else {
        delete x["id"];
      }
      x["name"] = A.load.Ref(ptr + 8, undefined);
      x["deviceName"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 25)) {
        x["active"] = A.load.Bool(ptr + 16);
      } else {
        delete x["active"];
      }
      if (A.load.Bool(ptr + 26)) {
        x["nodeVolume"] = A.load.Int32(ptr + 20);
      } else {
        delete x["nodeVolume"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AudioInputNodeInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 27, false);
        A.store.Bool(ptr + 24, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 26, false);
        A.store.Int32(ptr + 20, 0);
      } else {
        A.store.Bool(ptr + 27, true);
        A.store.Bool(ptr + 24, "id" in x ? true : false);
        A.store.Float64(ptr + 0, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Ref(ptr + 8, x["name"]);
        A.store.Ref(ptr + 12, x["deviceName"]);
        A.store.Bool(ptr + 25, "active" in x ? true : false);
        A.store.Bool(ptr + 16, x["active"] ? true : false);
        A.store.Bool(ptr + 26, "nodeGain" in x ? true : false);
        A.store.Int32(ptr + 20, x["nodeGain"] === undefined ? 0 : (x["nodeGain"] as number));
      }
    },
    "load_AudioInputNodeInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["id"] = A.load.Float64(ptr + 0);
      } else {
        delete x["id"];
      }
      x["name"] = A.load.Ref(ptr + 8, undefined);
      x["deviceName"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 25)) {
        x["active"] = A.load.Bool(ptr + 16);
      } else {
        delete x["active"];
      }
      if (A.load.Bool(ptr + 26)) {
        x["nodeGain"] = A.load.Int32(ptr + 20);
      } else {
        delete x["nodeGain"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AudioInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 22, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 23, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 24, true);
        A.store.Bool(ptr + 20, "outputMute" in x ? true : false);
        A.store.Bool(ptr + 0, x["outputMute"] ? true : false);
        A.store.Bool(ptr + 21, "inputMute" in x ? true : false);
        A.store.Bool(ptr + 1, x["inputMute"] ? true : false);
        A.store.Bool(ptr + 22, "underruns" in x ? true : false);
        A.store.Int32(ptr + 4, x["underruns"] === undefined ? 0 : (x["underruns"] as number));
        A.store.Bool(ptr + 23, "severeUnderruns" in x ? true : false);
        A.store.Int32(ptr + 8, x["severeUnderruns"] === undefined ? 0 : (x["severeUnderruns"] as number));
        A.store.Ref(ptr + 12, x["outputNodes"]);
        A.store.Ref(ptr + 16, x["inputNodes"]);
      }
    },
    "load_AudioInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 20)) {
        x["outputMute"] = A.load.Bool(ptr + 0);
      } else {
        delete x["outputMute"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["inputMute"] = A.load.Bool(ptr + 1);
      } else {
        delete x["inputMute"];
      }
      if (A.load.Bool(ptr + 22)) {
        x["underruns"] = A.load.Int32(ptr + 4);
      } else {
        delete x["underruns"];
      }
      if (A.load.Bool(ptr + 23)) {
        x["severeUnderruns"] = A.load.Int32(ptr + 8);
      } else {
        delete x["severeUnderruns"];
      }
      x["outputNodes"] = A.load.Ref(ptr + 12, undefined);
      x["inputNodes"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_BatteryInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 104, false);
        A.store.Bool(ptr + 96, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 97, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Bool(ptr + 98, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Bool(ptr + 99, false);
        A.store.Float64(ptr + 32, 0);
        A.store.Bool(ptr + 100, false);
        A.store.Float64(ptr + 40, 0);
        A.store.Ref(ptr + 48, undefined);
        A.store.Bool(ptr + 101, false);
        A.store.Float64(ptr + 56, 0);
        A.store.Bool(ptr + 102, false);
        A.store.Float64(ptr + 64, 0);
        A.store.Ref(ptr + 72, undefined);
        A.store.Ref(ptr + 76, undefined);
        A.store.Ref(ptr + 80, undefined);
        A.store.Bool(ptr + 103, false);
        A.store.Float64(ptr + 88, 0);
      } else {
        A.store.Bool(ptr + 104, true);
        A.store.Bool(ptr + 96, "cycleCount" in x ? true : false);
        A.store.Float64(ptr + 0, x["cycleCount"] === undefined ? 0 : (x["cycleCount"] as number));
        A.store.Bool(ptr + 97, "voltageNow" in x ? true : false);
        A.store.Float64(ptr + 8, x["voltageNow"] === undefined ? 0 : (x["voltageNow"] as number));
        A.store.Ref(ptr + 16, x["vendor"]);
        A.store.Ref(ptr + 20, x["serialNumber"]);
        A.store.Bool(ptr + 98, "chargeFullDesign" in x ? true : false);
        A.store.Float64(ptr + 24, x["chargeFullDesign"] === undefined ? 0 : (x["chargeFullDesign"] as number));
        A.store.Bool(ptr + 99, "chargeFull" in x ? true : false);
        A.store.Float64(ptr + 32, x["chargeFull"] === undefined ? 0 : (x["chargeFull"] as number));
        A.store.Bool(ptr + 100, "voltageMinDesign" in x ? true : false);
        A.store.Float64(ptr + 40, x["voltageMinDesign"] === undefined ? 0 : (x["voltageMinDesign"] as number));
        A.store.Ref(ptr + 48, x["modelName"]);
        A.store.Bool(ptr + 101, "chargeNow" in x ? true : false);
        A.store.Float64(ptr + 56, x["chargeNow"] === undefined ? 0 : (x["chargeNow"] as number));
        A.store.Bool(ptr + 102, "currentNow" in x ? true : false);
        A.store.Float64(ptr + 64, x["currentNow"] === undefined ? 0 : (x["currentNow"] as number));
        A.store.Ref(ptr + 72, x["technology"]);
        A.store.Ref(ptr + 76, x["status"]);
        A.store.Ref(ptr + 80, x["manufactureDate"]);
        A.store.Bool(ptr + 103, "temperature" in x ? true : false);
        A.store.Float64(ptr + 88, x["temperature"] === undefined ? 0 : (x["temperature"] as number));
      }
    },
    "load_BatteryInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 96)) {
        x["cycleCount"] = A.load.Float64(ptr + 0);
      } else {
        delete x["cycleCount"];
      }
      if (A.load.Bool(ptr + 97)) {
        x["voltageNow"] = A.load.Float64(ptr + 8);
      } else {
        delete x["voltageNow"];
      }
      x["vendor"] = A.load.Ref(ptr + 16, undefined);
      x["serialNumber"] = A.load.Ref(ptr + 20, undefined);
      if (A.load.Bool(ptr + 98)) {
        x["chargeFullDesign"] = A.load.Float64(ptr + 24);
      } else {
        delete x["chargeFullDesign"];
      }
      if (A.load.Bool(ptr + 99)) {
        x["chargeFull"] = A.load.Float64(ptr + 32);
      } else {
        delete x["chargeFull"];
      }
      if (A.load.Bool(ptr + 100)) {
        x["voltageMinDesign"] = A.load.Float64(ptr + 40);
      } else {
        delete x["voltageMinDesign"];
      }
      x["modelName"] = A.load.Ref(ptr + 48, undefined);
      if (A.load.Bool(ptr + 101)) {
        x["chargeNow"] = A.load.Float64(ptr + 56);
      } else {
        delete x["chargeNow"];
      }
      if (A.load.Bool(ptr + 102)) {
        x["currentNow"] = A.load.Float64(ptr + 64);
      } else {
        delete x["currentNow"];
      }
      x["technology"] = A.load.Ref(ptr + 72, undefined);
      x["status"] = A.load.Ref(ptr + 76, undefined);
      x["manufactureDate"] = A.load.Ref(ptr + 80, undefined);
      if (A.load.Bool(ptr + 103)) {
        x["temperature"] = A.load.Float64(ptr + 88);
      } else {
        delete x["temperature"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_CpuArchitectureEnum": (ref: heap.Ref<string>): number => {
      const idx = ["unknown", "x86_64", "aarch64", "armv7l"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_CpuCStateInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Bool(ptr + 16, "timeInStateSinceLastBootUs" in x ? true : false);
        A.store.Float64(
          ptr + 8,
          x["timeInStateSinceLastBootUs"] === undefined ? 0 : (x["timeInStateSinceLastBootUs"] as number)
        );
      }
    },
    "load_CpuCStateInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["timeInStateSinceLastBootUs"] = A.load.Float64(ptr + 8);
      } else {
        delete x["timeInStateSinceLastBootUs"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_LogicalCpuInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Float64(ptr + 16, 0);
        A.store.Ref(ptr + 24, undefined);
        A.store.Bool(ptr + 36, false);
        A.store.Int32(ptr + 28, 0);
      } else {
        A.store.Bool(ptr + 37, true);
        A.store.Bool(ptr + 32, "maxClockSpeedKhz" in x ? true : false);
        A.store.Int32(ptr + 0, x["maxClockSpeedKhz"] === undefined ? 0 : (x["maxClockSpeedKhz"] as number));
        A.store.Bool(ptr + 33, "scalingMaxFrequencyKhz" in x ? true : false);
        A.store.Int32(ptr + 4, x["scalingMaxFrequencyKhz"] === undefined ? 0 : (x["scalingMaxFrequencyKhz"] as number));
        A.store.Bool(ptr + 34, "scalingCurrentFrequencyKhz" in x ? true : false);
        A.store.Int32(
          ptr + 8,
          x["scalingCurrentFrequencyKhz"] === undefined ? 0 : (x["scalingCurrentFrequencyKhz"] as number)
        );
        A.store.Bool(ptr + 35, "idleTimeMs" in x ? true : false);
        A.store.Float64(ptr + 16, x["idleTimeMs"] === undefined ? 0 : (x["idleTimeMs"] as number));
        A.store.Ref(ptr + 24, x["cStates"]);
        A.store.Bool(ptr + 36, "coreId" in x ? true : false);
        A.store.Int32(ptr + 28, x["coreId"] === undefined ? 0 : (x["coreId"] as number));
      }
    },
    "load_LogicalCpuInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 32)) {
        x["maxClockSpeedKhz"] = A.load.Int32(ptr + 0);
      } else {
        delete x["maxClockSpeedKhz"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["scalingMaxFrequencyKhz"] = A.load.Int32(ptr + 4);
      } else {
        delete x["scalingMaxFrequencyKhz"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["scalingCurrentFrequencyKhz"] = A.load.Int32(ptr + 8);
      } else {
        delete x["scalingCurrentFrequencyKhz"];
      }
      if (A.load.Bool(ptr + 35)) {
        x["idleTimeMs"] = A.load.Float64(ptr + 16);
      } else {
        delete x["idleTimeMs"];
      }
      x["cStates"] = A.load.Ref(ptr + 24, undefined);
      if (A.load.Bool(ptr + 36)) {
        x["coreId"] = A.load.Int32(ptr + 28);
      } else {
        delete x["coreId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PhysicalCpuInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["modelName"]);
        A.store.Ref(ptr + 4, x["logicalCpus"]);
      }
    },
    "load_PhysicalCpuInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["modelName"] = A.load.Ref(ptr + 0, undefined);
      x["logicalCpus"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CpuInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "numTotalThreads" in x ? true : false);
        A.store.Int32(ptr + 0, x["numTotalThreads"] === undefined ? 0 : (x["numTotalThreads"] as number));
        A.store.Enum(ptr + 4, ["unknown", "x86_64", "aarch64", "armv7l"].indexOf(x["architecture"] as string));
        A.store.Ref(ptr + 8, x["physicalCpus"]);
      }
    },
    "load_CpuInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["numTotalThreads"] = A.load.Int32(ptr + 0);
      } else {
        delete x["numTotalThreads"];
      }
      x["architecture"] = A.load.Enum(ptr + 4, ["unknown", "x86_64", "aarch64", "armv7l"]);
      x["physicalCpus"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DisplayInputType": (ref: heap.Ref<string>): number => {
      const idx = ["unknown", "digital", "analog"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_EmbeddedDisplayInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 75, false);
        A.store.Bool(ptr + 64, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 65, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 66, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 67, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 68, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 69, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Bool(ptr + 70, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Ref(ptr + 32, undefined);
        A.store.Bool(ptr + 71, false);
        A.store.Int32(ptr + 36, 0);
        A.store.Bool(ptr + 72, false);
        A.store.Int32(ptr + 40, 0);
        A.store.Bool(ptr + 73, false);
        A.store.Int32(ptr + 44, 0);
        A.store.Bool(ptr + 74, false);
        A.store.Int32(ptr + 48, 0);
        A.store.Ref(ptr + 52, undefined);
        A.store.Enum(ptr + 56, -1);
        A.store.Ref(ptr + 60, undefined);
      } else {
        A.store.Bool(ptr + 75, true);
        A.store.Bool(ptr + 64, "privacyScreenSupported" in x ? true : false);
        A.store.Bool(ptr + 0, x["privacyScreenSupported"] ? true : false);
        A.store.Bool(ptr + 65, "privacyScreenEnabled" in x ? true : false);
        A.store.Bool(ptr + 1, x["privacyScreenEnabled"] ? true : false);
        A.store.Bool(ptr + 66, "displayWidth" in x ? true : false);
        A.store.Int32(ptr + 4, x["displayWidth"] === undefined ? 0 : (x["displayWidth"] as number));
        A.store.Bool(ptr + 67, "displayHeight" in x ? true : false);
        A.store.Int32(ptr + 8, x["displayHeight"] === undefined ? 0 : (x["displayHeight"] as number));
        A.store.Bool(ptr + 68, "resolutionHorizontal" in x ? true : false);
        A.store.Int32(ptr + 12, x["resolutionHorizontal"] === undefined ? 0 : (x["resolutionHorizontal"] as number));
        A.store.Bool(ptr + 69, "resolutionVertical" in x ? true : false);
        A.store.Int32(ptr + 16, x["resolutionVertical"] === undefined ? 0 : (x["resolutionVertical"] as number));
        A.store.Bool(ptr + 70, "refreshRate" in x ? true : false);
        A.store.Float64(ptr + 24, x["refreshRate"] === undefined ? 0 : (x["refreshRate"] as number));
        A.store.Ref(ptr + 32, x["manufacturer"]);
        A.store.Bool(ptr + 71, "modelId" in x ? true : false);
        A.store.Int32(ptr + 36, x["modelId"] === undefined ? 0 : (x["modelId"] as number));
        A.store.Bool(ptr + 72, "serialNumber" in x ? true : false);
        A.store.Int32(ptr + 40, x["serialNumber"] === undefined ? 0 : (x["serialNumber"] as number));
        A.store.Bool(ptr + 73, "manufactureWeek" in x ? true : false);
        A.store.Int32(ptr + 44, x["manufactureWeek"] === undefined ? 0 : (x["manufactureWeek"] as number));
        A.store.Bool(ptr + 74, "manufactureYear" in x ? true : false);
        A.store.Int32(ptr + 48, x["manufactureYear"] === undefined ? 0 : (x["manufactureYear"] as number));
        A.store.Ref(ptr + 52, x["edidVersion"]);
        A.store.Enum(ptr + 56, ["unknown", "digital", "analog"].indexOf(x["inputType"] as string));
        A.store.Ref(ptr + 60, x["displayName"]);
      }
    },
    "load_EmbeddedDisplayInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 64)) {
        x["privacyScreenSupported"] = A.load.Bool(ptr + 0);
      } else {
        delete x["privacyScreenSupported"];
      }
      if (A.load.Bool(ptr + 65)) {
        x["privacyScreenEnabled"] = A.load.Bool(ptr + 1);
      } else {
        delete x["privacyScreenEnabled"];
      }
      if (A.load.Bool(ptr + 66)) {
        x["displayWidth"] = A.load.Int32(ptr + 4);
      } else {
        delete x["displayWidth"];
      }
      if (A.load.Bool(ptr + 67)) {
        x["displayHeight"] = A.load.Int32(ptr + 8);
      } else {
        delete x["displayHeight"];
      }
      if (A.load.Bool(ptr + 68)) {
        x["resolutionHorizontal"] = A.load.Int32(ptr + 12);
      } else {
        delete x["resolutionHorizontal"];
      }
      if (A.load.Bool(ptr + 69)) {
        x["resolutionVertical"] = A.load.Int32(ptr + 16);
      } else {
        delete x["resolutionVertical"];
      }
      if (A.load.Bool(ptr + 70)) {
        x["refreshRate"] = A.load.Float64(ptr + 24);
      } else {
        delete x["refreshRate"];
      }
      x["manufacturer"] = A.load.Ref(ptr + 32, undefined);
      if (A.load.Bool(ptr + 71)) {
        x["modelId"] = A.load.Int32(ptr + 36);
      } else {
        delete x["modelId"];
      }
      if (A.load.Bool(ptr + 72)) {
        x["serialNumber"] = A.load.Int32(ptr + 40);
      } else {
        delete x["serialNumber"];
      }
      if (A.load.Bool(ptr + 73)) {
        x["manufactureWeek"] = A.load.Int32(ptr + 44);
      } else {
        delete x["manufactureWeek"];
      }
      if (A.load.Bool(ptr + 74)) {
        x["manufactureYear"] = A.load.Int32(ptr + 48);
      } else {
        delete x["manufactureYear"];
      }
      x["edidVersion"] = A.load.Ref(ptr + 52, undefined);
      x["inputType"] = A.load.Enum(ptr + 56, ["unknown", "digital", "analog"]);
      x["displayName"] = A.load.Ref(ptr + 60, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ExternalDisplayInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 65, false);
        A.store.Bool(ptr + 56, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 57, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 58, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 59, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 60, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Ref(ptr + 24, undefined);
        A.store.Bool(ptr + 61, false);
        A.store.Int32(ptr + 28, 0);
        A.store.Bool(ptr + 62, false);
        A.store.Int32(ptr + 32, 0);
        A.store.Bool(ptr + 63, false);
        A.store.Int32(ptr + 36, 0);
        A.store.Bool(ptr + 64, false);
        A.store.Int32(ptr + 40, 0);
        A.store.Ref(ptr + 44, undefined);
        A.store.Enum(ptr + 48, -1);
        A.store.Ref(ptr + 52, undefined);
      } else {
        A.store.Bool(ptr + 65, true);
        A.store.Bool(ptr + 56, "displayWidth" in x ? true : false);
        A.store.Int32(ptr + 0, x["displayWidth"] === undefined ? 0 : (x["displayWidth"] as number));
        A.store.Bool(ptr + 57, "displayHeight" in x ? true : false);
        A.store.Int32(ptr + 4, x["displayHeight"] === undefined ? 0 : (x["displayHeight"] as number));
        A.store.Bool(ptr + 58, "resolutionHorizontal" in x ? true : false);
        A.store.Int32(ptr + 8, x["resolutionHorizontal"] === undefined ? 0 : (x["resolutionHorizontal"] as number));
        A.store.Bool(ptr + 59, "resolutionVertical" in x ? true : false);
        A.store.Int32(ptr + 12, x["resolutionVertical"] === undefined ? 0 : (x["resolutionVertical"] as number));
        A.store.Bool(ptr + 60, "refreshRate" in x ? true : false);
        A.store.Float64(ptr + 16, x["refreshRate"] === undefined ? 0 : (x["refreshRate"] as number));
        A.store.Ref(ptr + 24, x["manufacturer"]);
        A.store.Bool(ptr + 61, "modelId" in x ? true : false);
        A.store.Int32(ptr + 28, x["modelId"] === undefined ? 0 : (x["modelId"] as number));
        A.store.Bool(ptr + 62, "serialNumber" in x ? true : false);
        A.store.Int32(ptr + 32, x["serialNumber"] === undefined ? 0 : (x["serialNumber"] as number));
        A.store.Bool(ptr + 63, "manufactureWeek" in x ? true : false);
        A.store.Int32(ptr + 36, x["manufactureWeek"] === undefined ? 0 : (x["manufactureWeek"] as number));
        A.store.Bool(ptr + 64, "manufactureYear" in x ? true : false);
        A.store.Int32(ptr + 40, x["manufactureYear"] === undefined ? 0 : (x["manufactureYear"] as number));
        A.store.Ref(ptr + 44, x["edidVersion"]);
        A.store.Enum(ptr + 48, ["unknown", "digital", "analog"].indexOf(x["inputType"] as string));
        A.store.Ref(ptr + 52, x["displayName"]);
      }
    },
    "load_ExternalDisplayInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 56)) {
        x["displayWidth"] = A.load.Int32(ptr + 0);
      } else {
        delete x["displayWidth"];
      }
      if (A.load.Bool(ptr + 57)) {
        x["displayHeight"] = A.load.Int32(ptr + 4);
      } else {
        delete x["displayHeight"];
      }
      if (A.load.Bool(ptr + 58)) {
        x["resolutionHorizontal"] = A.load.Int32(ptr + 8);
      } else {
        delete x["resolutionHorizontal"];
      }
      if (A.load.Bool(ptr + 59)) {
        x["resolutionVertical"] = A.load.Int32(ptr + 12);
      } else {
        delete x["resolutionVertical"];
      }
      if (A.load.Bool(ptr + 60)) {
        x["refreshRate"] = A.load.Float64(ptr + 16);
      } else {
        delete x["refreshRate"];
      }
      x["manufacturer"] = A.load.Ref(ptr + 24, undefined);
      if (A.load.Bool(ptr + 61)) {
        x["modelId"] = A.load.Int32(ptr + 28);
      } else {
        delete x["modelId"];
      }
      if (A.load.Bool(ptr + 62)) {
        x["serialNumber"] = A.load.Int32(ptr + 32);
      } else {
        delete x["serialNumber"];
      }
      if (A.load.Bool(ptr + 63)) {
        x["manufactureWeek"] = A.load.Int32(ptr + 36);
      } else {
        delete x["manufactureWeek"];
      }
      if (A.load.Bool(ptr + 64)) {
        x["manufactureYear"] = A.load.Int32(ptr + 40);
      } else {
        delete x["manufactureYear"];
      }
      x["edidVersion"] = A.load.Ref(ptr + 44, undefined);
      x["inputType"] = A.load.Enum(ptr + 48, ["unknown", "digital", "analog"]);
      x["displayName"] = A.load.Ref(ptr + 52, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DisplayInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 80, false);

        A.store.Bool(ptr + 0 + 75, false);
        A.store.Bool(ptr + 0 + 64, false);
        A.store.Bool(ptr + 0 + 0, false);
        A.store.Bool(ptr + 0 + 65, false);
        A.store.Bool(ptr + 0 + 1, false);
        A.store.Bool(ptr + 0 + 66, false);
        A.store.Int32(ptr + 0 + 4, 0);
        A.store.Bool(ptr + 0 + 67, false);
        A.store.Int32(ptr + 0 + 8, 0);
        A.store.Bool(ptr + 0 + 68, false);
        A.store.Int32(ptr + 0 + 12, 0);
        A.store.Bool(ptr + 0 + 69, false);
        A.store.Int32(ptr + 0 + 16, 0);
        A.store.Bool(ptr + 0 + 70, false);
        A.store.Float64(ptr + 0 + 24, 0);
        A.store.Ref(ptr + 0 + 32, undefined);
        A.store.Bool(ptr + 0 + 71, false);
        A.store.Int32(ptr + 0 + 36, 0);
        A.store.Bool(ptr + 0 + 72, false);
        A.store.Int32(ptr + 0 + 40, 0);
        A.store.Bool(ptr + 0 + 73, false);
        A.store.Int32(ptr + 0 + 44, 0);
        A.store.Bool(ptr + 0 + 74, false);
        A.store.Int32(ptr + 0 + 48, 0);
        A.store.Ref(ptr + 0 + 52, undefined);
        A.store.Enum(ptr + 0 + 56, -1);
        A.store.Ref(ptr + 0 + 60, undefined);
        A.store.Ref(ptr + 76, undefined);
      } else {
        A.store.Bool(ptr + 80, true);

        if (typeof x["embeddedDisplay"] === "undefined") {
          A.store.Bool(ptr + 0 + 75, false);
          A.store.Bool(ptr + 0 + 64, false);
          A.store.Bool(ptr + 0 + 0, false);
          A.store.Bool(ptr + 0 + 65, false);
          A.store.Bool(ptr + 0 + 1, false);
          A.store.Bool(ptr + 0 + 66, false);
          A.store.Int32(ptr + 0 + 4, 0);
          A.store.Bool(ptr + 0 + 67, false);
          A.store.Int32(ptr + 0 + 8, 0);
          A.store.Bool(ptr + 0 + 68, false);
          A.store.Int32(ptr + 0 + 12, 0);
          A.store.Bool(ptr + 0 + 69, false);
          A.store.Int32(ptr + 0 + 16, 0);
          A.store.Bool(ptr + 0 + 70, false);
          A.store.Float64(ptr + 0 + 24, 0);
          A.store.Ref(ptr + 0 + 32, undefined);
          A.store.Bool(ptr + 0 + 71, false);
          A.store.Int32(ptr + 0 + 36, 0);
          A.store.Bool(ptr + 0 + 72, false);
          A.store.Int32(ptr + 0 + 40, 0);
          A.store.Bool(ptr + 0 + 73, false);
          A.store.Int32(ptr + 0 + 44, 0);
          A.store.Bool(ptr + 0 + 74, false);
          A.store.Int32(ptr + 0 + 48, 0);
          A.store.Ref(ptr + 0 + 52, undefined);
          A.store.Enum(ptr + 0 + 56, -1);
          A.store.Ref(ptr + 0 + 60, undefined);
        } else {
          A.store.Bool(ptr + 0 + 75, true);
          A.store.Bool(ptr + 0 + 64, "privacyScreenSupported" in x["embeddedDisplay"] ? true : false);
          A.store.Bool(ptr + 0 + 0, x["embeddedDisplay"]["privacyScreenSupported"] ? true : false);
          A.store.Bool(ptr + 0 + 65, "privacyScreenEnabled" in x["embeddedDisplay"] ? true : false);
          A.store.Bool(ptr + 0 + 1, x["embeddedDisplay"]["privacyScreenEnabled"] ? true : false);
          A.store.Bool(ptr + 0 + 66, "displayWidth" in x["embeddedDisplay"] ? true : false);
          A.store.Int32(
            ptr + 0 + 4,
            x["embeddedDisplay"]["displayWidth"] === undefined ? 0 : (x["embeddedDisplay"]["displayWidth"] as number)
          );
          A.store.Bool(ptr + 0 + 67, "displayHeight" in x["embeddedDisplay"] ? true : false);
          A.store.Int32(
            ptr + 0 + 8,
            x["embeddedDisplay"]["displayHeight"] === undefined ? 0 : (x["embeddedDisplay"]["displayHeight"] as number)
          );
          A.store.Bool(ptr + 0 + 68, "resolutionHorizontal" in x["embeddedDisplay"] ? true : false);
          A.store.Int32(
            ptr + 0 + 12,
            x["embeddedDisplay"]["resolutionHorizontal"] === undefined
              ? 0
              : (x["embeddedDisplay"]["resolutionHorizontal"] as number)
          );
          A.store.Bool(ptr + 0 + 69, "resolutionVertical" in x["embeddedDisplay"] ? true : false);
          A.store.Int32(
            ptr + 0 + 16,
            x["embeddedDisplay"]["resolutionVertical"] === undefined
              ? 0
              : (x["embeddedDisplay"]["resolutionVertical"] as number)
          );
          A.store.Bool(ptr + 0 + 70, "refreshRate" in x["embeddedDisplay"] ? true : false);
          A.store.Float64(
            ptr + 0 + 24,
            x["embeddedDisplay"]["refreshRate"] === undefined ? 0 : (x["embeddedDisplay"]["refreshRate"] as number)
          );
          A.store.Ref(ptr + 0 + 32, x["embeddedDisplay"]["manufacturer"]);
          A.store.Bool(ptr + 0 + 71, "modelId" in x["embeddedDisplay"] ? true : false);
          A.store.Int32(
            ptr + 0 + 36,
            x["embeddedDisplay"]["modelId"] === undefined ? 0 : (x["embeddedDisplay"]["modelId"] as number)
          );
          A.store.Bool(ptr + 0 + 72, "serialNumber" in x["embeddedDisplay"] ? true : false);
          A.store.Int32(
            ptr + 0 + 40,
            x["embeddedDisplay"]["serialNumber"] === undefined ? 0 : (x["embeddedDisplay"]["serialNumber"] as number)
          );
          A.store.Bool(ptr + 0 + 73, "manufactureWeek" in x["embeddedDisplay"] ? true : false);
          A.store.Int32(
            ptr + 0 + 44,
            x["embeddedDisplay"]["manufactureWeek"] === undefined
              ? 0
              : (x["embeddedDisplay"]["manufactureWeek"] as number)
          );
          A.store.Bool(ptr + 0 + 74, "manufactureYear" in x["embeddedDisplay"] ? true : false);
          A.store.Int32(
            ptr + 0 + 48,
            x["embeddedDisplay"]["manufactureYear"] === undefined
              ? 0
              : (x["embeddedDisplay"]["manufactureYear"] as number)
          );
          A.store.Ref(ptr + 0 + 52, x["embeddedDisplay"]["edidVersion"]);
          A.store.Enum(
            ptr + 0 + 56,
            ["unknown", "digital", "analog"].indexOf(x["embeddedDisplay"]["inputType"] as string)
          );
          A.store.Ref(ptr + 0 + 60, x["embeddedDisplay"]["displayName"]);
        }
        A.store.Ref(ptr + 76, x["externalDisplays"]);
      }
    },
    "load_DisplayInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 75)) {
        x["embeddedDisplay"] = {};
        if (A.load.Bool(ptr + 0 + 64)) {
          x["embeddedDisplay"]["privacyScreenSupported"] = A.load.Bool(ptr + 0 + 0);
        } else {
          delete x["embeddedDisplay"]["privacyScreenSupported"];
        }
        if (A.load.Bool(ptr + 0 + 65)) {
          x["embeddedDisplay"]["privacyScreenEnabled"] = A.load.Bool(ptr + 0 + 1);
        } else {
          delete x["embeddedDisplay"]["privacyScreenEnabled"];
        }
        if (A.load.Bool(ptr + 0 + 66)) {
          x["embeddedDisplay"]["displayWidth"] = A.load.Int32(ptr + 0 + 4);
        } else {
          delete x["embeddedDisplay"]["displayWidth"];
        }
        if (A.load.Bool(ptr + 0 + 67)) {
          x["embeddedDisplay"]["displayHeight"] = A.load.Int32(ptr + 0 + 8);
        } else {
          delete x["embeddedDisplay"]["displayHeight"];
        }
        if (A.load.Bool(ptr + 0 + 68)) {
          x["embeddedDisplay"]["resolutionHorizontal"] = A.load.Int32(ptr + 0 + 12);
        } else {
          delete x["embeddedDisplay"]["resolutionHorizontal"];
        }
        if (A.load.Bool(ptr + 0 + 69)) {
          x["embeddedDisplay"]["resolutionVertical"] = A.load.Int32(ptr + 0 + 16);
        } else {
          delete x["embeddedDisplay"]["resolutionVertical"];
        }
        if (A.load.Bool(ptr + 0 + 70)) {
          x["embeddedDisplay"]["refreshRate"] = A.load.Float64(ptr + 0 + 24);
        } else {
          delete x["embeddedDisplay"]["refreshRate"];
        }
        x["embeddedDisplay"]["manufacturer"] = A.load.Ref(ptr + 0 + 32, undefined);
        if (A.load.Bool(ptr + 0 + 71)) {
          x["embeddedDisplay"]["modelId"] = A.load.Int32(ptr + 0 + 36);
        } else {
          delete x["embeddedDisplay"]["modelId"];
        }
        if (A.load.Bool(ptr + 0 + 72)) {
          x["embeddedDisplay"]["serialNumber"] = A.load.Int32(ptr + 0 + 40);
        } else {
          delete x["embeddedDisplay"]["serialNumber"];
        }
        if (A.load.Bool(ptr + 0 + 73)) {
          x["embeddedDisplay"]["manufactureWeek"] = A.load.Int32(ptr + 0 + 44);
        } else {
          delete x["embeddedDisplay"]["manufactureWeek"];
        }
        if (A.load.Bool(ptr + 0 + 74)) {
          x["embeddedDisplay"]["manufactureYear"] = A.load.Int32(ptr + 0 + 48);
        } else {
          delete x["embeddedDisplay"]["manufactureYear"];
        }
        x["embeddedDisplay"]["edidVersion"] = A.load.Ref(ptr + 0 + 52, undefined);
        x["embeddedDisplay"]["inputType"] = A.load.Enum(ptr + 0 + 56, ["unknown", "digital", "analog"]);
        x["embeddedDisplay"]["displayName"] = A.load.Ref(ptr + 0 + 60, undefined);
      } else {
        delete x["embeddedDisplay"];
      }
      x["externalDisplays"] = A.load.Ref(ptr + 76, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_FwupdVersionFormat": (ref: heap.Ref<string>): number => {
      const idx = [
        "plain",
        "number",
        "pair",
        "triplet",
        "quad",
        "bcd",
        "intelMe",
        "intelMe2",
        "surfaceLegacy",
        "surface",
        "dellBios",
        "hex",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_FwupdFirmwareVersionInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["version"]);
        A.store.Enum(
          ptr + 4,
          [
            "plain",
            "number",
            "pair",
            "triplet",
            "quad",
            "bcd",
            "intelMe",
            "intelMe2",
            "surfaceLegacy",
            "surface",
            "dellBios",
            "hex",
          ].indexOf(x["version_format"] as string)
        );
      }
    },
    "load_FwupdFirmwareVersionInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["version"] = A.load.Ref(ptr + 0, undefined);
      x["version_format"] = A.load.Enum(ptr + 4, [
        "plain",
        "number",
        "pair",
        "triplet",
        "quad",
        "bcd",
        "intelMe",
        "intelMe2",
        "surfaceLegacy",
        "surface",
        "dellBios",
        "hex",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_NetworkType": (ref: heap.Ref<string>): number => {
      const idx = ["cellular", "ethernet", "tether", "vpn", "wifi"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_NetworkState": (ref: heap.Ref<string>): number => {
      const idx = [
        "uninitialized",
        "disabled",
        "prohibited",
        "not_connected",
        "connecting",
        "portal",
        "connected",
        "online",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_NetworkInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 33, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 32, false);
        A.store.Float64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 33, true);
        A.store.Enum(ptr + 0, ["cellular", "ethernet", "tether", "vpn", "wifi"].indexOf(x["type"] as string));
        A.store.Enum(
          ptr + 4,
          [
            "uninitialized",
            "disabled",
            "prohibited",
            "not_connected",
            "connecting",
            "portal",
            "connected",
            "online",
          ].indexOf(x["state"] as string)
        );
        A.store.Ref(ptr + 8, x["macAddress"]);
        A.store.Ref(ptr + 12, x["ipv4Address"]);
        A.store.Ref(ptr + 16, x["ipv6Addresses"]);
        A.store.Bool(ptr + 32, "signalStrength" in x ? true : false);
        A.store.Float64(ptr + 24, x["signalStrength"] === undefined ? 0 : (x["signalStrength"] as number));
      }
    },
    "load_NetworkInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, ["cellular", "ethernet", "tether", "vpn", "wifi"]);
      x["state"] = A.load.Enum(ptr + 4, [
        "uninitialized",
        "disabled",
        "prohibited",
        "not_connected",
        "connecting",
        "portal",
        "connected",
        "online",
      ]);
      x["macAddress"] = A.load.Ref(ptr + 8, undefined);
      x["ipv4Address"] = A.load.Ref(ptr + 12, undefined);
      x["ipv6Addresses"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 32)) {
        x["signalStrength"] = A.load.Float64(ptr + 24);
      } else {
        delete x["signalStrength"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_InternetConnectivityInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["networks"]);
      }
    },
    "load_InternetConnectivityInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["networks"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MarketingInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["marketingName"]);
      }
    },
    "load_MarketingInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["marketingName"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MemoryInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 24, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 25, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 26, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 27, false);
        A.store.Float64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 28, true);
        A.store.Bool(ptr + 24, "totalMemoryKiB" in x ? true : false);
        A.store.Int32(ptr + 0, x["totalMemoryKiB"] === undefined ? 0 : (x["totalMemoryKiB"] as number));
        A.store.Bool(ptr + 25, "freeMemoryKiB" in x ? true : false);
        A.store.Int32(ptr + 4, x["freeMemoryKiB"] === undefined ? 0 : (x["freeMemoryKiB"] as number));
        A.store.Bool(ptr + 26, "availableMemoryKiB" in x ? true : false);
        A.store.Int32(ptr + 8, x["availableMemoryKiB"] === undefined ? 0 : (x["availableMemoryKiB"] as number));
        A.store.Bool(ptr + 27, "pageFaultsSinceLastBoot" in x ? true : false);
        A.store.Float64(
          ptr + 16,
          x["pageFaultsSinceLastBoot"] === undefined ? 0 : (x["pageFaultsSinceLastBoot"] as number)
        );
      }
    },
    "load_MemoryInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["totalMemoryKiB"] = A.load.Int32(ptr + 0);
      } else {
        delete x["totalMemoryKiB"];
      }
      if (A.load.Bool(ptr + 25)) {
        x["freeMemoryKiB"] = A.load.Int32(ptr + 4);
      } else {
        delete x["freeMemoryKiB"];
      }
      if (A.load.Bool(ptr + 26)) {
        x["availableMemoryKiB"] = A.load.Int32(ptr + 8);
      } else {
        delete x["availableMemoryKiB"];
      }
      if (A.load.Bool(ptr + 27)) {
        x["pageFaultsSinceLastBoot"] = A.load.Float64(ptr + 16);
      } else {
        delete x["pageFaultsSinceLastBoot"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_NonRemovableBlockDeviceInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["type"]);
        A.store.Bool(ptr + 16, "size" in x ? true : false);
        A.store.Float64(ptr + 8, x["size"] === undefined ? 0 : (x["size"] as number));
      }
    },
    "load_NonRemovableBlockDeviceInfo": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["type"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["size"] = A.load.Float64(ptr + 8);
      } else {
        delete x["size"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_NonRemovableBlockDeviceInfoResponse": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["deviceInfos"]);
      }
    },
    "load_NonRemovableBlockDeviceInfoResponse": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["deviceInfos"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OemData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["oemData"]);
      }
    },
    "load_OemData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["oemData"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OsVersionInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["releaseMilestone"]);
        A.store.Ref(ptr + 4, x["buildNumber"]);
        A.store.Ref(ptr + 8, x["patchNumber"]);
        A.store.Ref(ptr + 12, x["releaseChannel"]);
      }
    },
    "load_OsVersionInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["releaseMilestone"] = A.load.Ref(ptr + 0, undefined);
      x["buildNumber"] = A.load.Ref(ptr + 4, undefined);
      x["patchNumber"] = A.load.Ref(ptr + 8, undefined);
      x["releaseChannel"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_StatefulPartitionInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Bool(ptr + 16, "availableSpace" in x ? true : false);
        A.store.Float64(ptr + 0, x["availableSpace"] === undefined ? 0 : (x["availableSpace"] as number));
        A.store.Bool(ptr + 17, "totalSpace" in x ? true : false);
        A.store.Float64(ptr + 8, x["totalSpace"] === undefined ? 0 : (x["totalSpace"] as number));
      }
    },
    "load_StatefulPartitionInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["availableSpace"] = A.load.Float64(ptr + 0);
      } else {
        delete x["availableSpace"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["totalSpace"] = A.load.Float64(ptr + 8);
      } else {
        delete x["totalSpace"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TpmDictionaryAttack": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 19, false);
        A.store.Int32(ptr + 12, 0);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Bool(ptr + 16, "counter" in x ? true : false);
        A.store.Int32(ptr + 0, x["counter"] === undefined ? 0 : (x["counter"] as number));
        A.store.Bool(ptr + 17, "threshold" in x ? true : false);
        A.store.Int32(ptr + 4, x["threshold"] === undefined ? 0 : (x["threshold"] as number));
        A.store.Bool(ptr + 18, "lockoutInEffect" in x ? true : false);
        A.store.Bool(ptr + 8, x["lockoutInEffect"] ? true : false);
        A.store.Bool(ptr + 19, "lockoutSecondsRemaining" in x ? true : false);
        A.store.Int32(
          ptr + 12,
          x["lockoutSecondsRemaining"] === undefined ? 0 : (x["lockoutSecondsRemaining"] as number)
        );
      }
    },
    "load_TpmDictionaryAttack": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["counter"] = A.load.Int32(ptr + 0);
      } else {
        delete x["counter"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["threshold"] = A.load.Int32(ptr + 4);
      } else {
        delete x["threshold"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["lockoutInEffect"] = A.load.Bool(ptr + 8);
      } else {
        delete x["lockoutInEffect"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["lockoutSecondsRemaining"] = A.load.Int32(ptr + 12);
      } else {
        delete x["lockoutSecondsRemaining"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_TpmGSCVersion": (ref: heap.Ref<string>): number => {
      const idx = ["not_gsc", "cr50", "ti50"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_TpmVersion": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 41, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 36, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 37, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 38, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Bool(ptr + 39, false);
        A.store.Int32(ptr + 20, 0);
        A.store.Bool(ptr + 40, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Ref(ptr + 32, undefined);
      } else {
        A.store.Bool(ptr + 41, true);
        A.store.Enum(ptr + 0, ["not_gsc", "cr50", "ti50"].indexOf(x["gscVersion"] as string));
        A.store.Bool(ptr + 36, "family" in x ? true : false);
        A.store.Int32(ptr + 4, x["family"] === undefined ? 0 : (x["family"] as number));
        A.store.Bool(ptr + 37, "specLevel" in x ? true : false);
        A.store.Float64(ptr + 8, x["specLevel"] === undefined ? 0 : (x["specLevel"] as number));
        A.store.Bool(ptr + 38, "manufacturer" in x ? true : false);
        A.store.Int32(ptr + 16, x["manufacturer"] === undefined ? 0 : (x["manufacturer"] as number));
        A.store.Bool(ptr + 39, "tpmModel" in x ? true : false);
        A.store.Int32(ptr + 20, x["tpmModel"] === undefined ? 0 : (x["tpmModel"] as number));
        A.store.Bool(ptr + 40, "firmwareVersion" in x ? true : false);
        A.store.Float64(ptr + 24, x["firmwareVersion"] === undefined ? 0 : (x["firmwareVersion"] as number));
        A.store.Ref(ptr + 32, x["vendorSpecific"]);
      }
    },
    "load_TpmVersion": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["gscVersion"] = A.load.Enum(ptr + 0, ["not_gsc", "cr50", "ti50"]);
      if (A.load.Bool(ptr + 36)) {
        x["family"] = A.load.Int32(ptr + 4);
      } else {
        delete x["family"];
      }
      if (A.load.Bool(ptr + 37)) {
        x["specLevel"] = A.load.Float64(ptr + 8);
      } else {
        delete x["specLevel"];
      }
      if (A.load.Bool(ptr + 38)) {
        x["manufacturer"] = A.load.Int32(ptr + 16);
      } else {
        delete x["manufacturer"];
      }
      if (A.load.Bool(ptr + 39)) {
        x["tpmModel"] = A.load.Int32(ptr + 20);
      } else {
        delete x["tpmModel"];
      }
      if (A.load.Bool(ptr + 40)) {
        x["firmwareVersion"] = A.load.Float64(ptr + 24);
      } else {
        delete x["firmwareVersion"];
      }
      x["vendorSpecific"] = A.load.Ref(ptr + 32, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TpmStatus": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Bool(ptr + 3, "enabled" in x ? true : false);
        A.store.Bool(ptr + 0, x["enabled"] ? true : false);
        A.store.Bool(ptr + 4, "owned" in x ? true : false);
        A.store.Bool(ptr + 1, x["owned"] ? true : false);
        A.store.Bool(ptr + 5, "ownerPasswordIsPresent" in x ? true : false);
        A.store.Bool(ptr + 2, x["ownerPasswordIsPresent"] ? true : false);
      }
    },
    "load_TpmStatus": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 3)) {
        x["enabled"] = A.load.Bool(ptr + 0);
      } else {
        delete x["enabled"];
      }
      if (A.load.Bool(ptr + 4)) {
        x["owned"] = A.load.Bool(ptr + 1);
      } else {
        delete x["owned"];
      }
      if (A.load.Bool(ptr + 5)) {
        x["ownerPasswordIsPresent"] = A.load.Bool(ptr + 2);
      } else {
        delete x["ownerPasswordIsPresent"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TpmInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 73, false);

        A.store.Bool(ptr + 0 + 41, false);
        A.store.Enum(ptr + 0 + 0, -1);
        A.store.Bool(ptr + 0 + 36, false);
        A.store.Int32(ptr + 0 + 4, 0);
        A.store.Bool(ptr + 0 + 37, false);
        A.store.Float64(ptr + 0 + 8, 0);
        A.store.Bool(ptr + 0 + 38, false);
        A.store.Int32(ptr + 0 + 16, 0);
        A.store.Bool(ptr + 0 + 39, false);
        A.store.Int32(ptr + 0 + 20, 0);
        A.store.Bool(ptr + 0 + 40, false);
        A.store.Float64(ptr + 0 + 24, 0);
        A.store.Ref(ptr + 0 + 32, undefined);

        A.store.Bool(ptr + 42 + 6, false);
        A.store.Bool(ptr + 42 + 3, false);
        A.store.Bool(ptr + 42 + 0, false);
        A.store.Bool(ptr + 42 + 4, false);
        A.store.Bool(ptr + 42 + 1, false);
        A.store.Bool(ptr + 42 + 5, false);
        A.store.Bool(ptr + 42 + 2, false);

        A.store.Bool(ptr + 52 + 20, false);
        A.store.Bool(ptr + 52 + 16, false);
        A.store.Int32(ptr + 52 + 0, 0);
        A.store.Bool(ptr + 52 + 17, false);
        A.store.Int32(ptr + 52 + 4, 0);
        A.store.Bool(ptr + 52 + 18, false);
        A.store.Bool(ptr + 52 + 8, false);
        A.store.Bool(ptr + 52 + 19, false);
        A.store.Int32(ptr + 52 + 12, 0);
      } else {
        A.store.Bool(ptr + 73, true);

        if (typeof x["version"] === "undefined") {
          A.store.Bool(ptr + 0 + 41, false);
          A.store.Enum(ptr + 0 + 0, -1);
          A.store.Bool(ptr + 0 + 36, false);
          A.store.Int32(ptr + 0 + 4, 0);
          A.store.Bool(ptr + 0 + 37, false);
          A.store.Float64(ptr + 0 + 8, 0);
          A.store.Bool(ptr + 0 + 38, false);
          A.store.Int32(ptr + 0 + 16, 0);
          A.store.Bool(ptr + 0 + 39, false);
          A.store.Int32(ptr + 0 + 20, 0);
          A.store.Bool(ptr + 0 + 40, false);
          A.store.Float64(ptr + 0 + 24, 0);
          A.store.Ref(ptr + 0 + 32, undefined);
        } else {
          A.store.Bool(ptr + 0 + 41, true);
          A.store.Enum(ptr + 0 + 0, ["not_gsc", "cr50", "ti50"].indexOf(x["version"]["gscVersion"] as string));
          A.store.Bool(ptr + 0 + 36, "family" in x["version"] ? true : false);
          A.store.Int32(ptr + 0 + 4, x["version"]["family"] === undefined ? 0 : (x["version"]["family"] as number));
          A.store.Bool(ptr + 0 + 37, "specLevel" in x["version"] ? true : false);
          A.store.Float64(
            ptr + 0 + 8,
            x["version"]["specLevel"] === undefined ? 0 : (x["version"]["specLevel"] as number)
          );
          A.store.Bool(ptr + 0 + 38, "manufacturer" in x["version"] ? true : false);
          A.store.Int32(
            ptr + 0 + 16,
            x["version"]["manufacturer"] === undefined ? 0 : (x["version"]["manufacturer"] as number)
          );
          A.store.Bool(ptr + 0 + 39, "tpmModel" in x["version"] ? true : false);
          A.store.Int32(
            ptr + 0 + 20,
            x["version"]["tpmModel"] === undefined ? 0 : (x["version"]["tpmModel"] as number)
          );
          A.store.Bool(ptr + 0 + 40, "firmwareVersion" in x["version"] ? true : false);
          A.store.Float64(
            ptr + 0 + 24,
            x["version"]["firmwareVersion"] === undefined ? 0 : (x["version"]["firmwareVersion"] as number)
          );
          A.store.Ref(ptr + 0 + 32, x["version"]["vendorSpecific"]);
        }

        if (typeof x["status"] === "undefined") {
          A.store.Bool(ptr + 42 + 6, false);
          A.store.Bool(ptr + 42 + 3, false);
          A.store.Bool(ptr + 42 + 0, false);
          A.store.Bool(ptr + 42 + 4, false);
          A.store.Bool(ptr + 42 + 1, false);
          A.store.Bool(ptr + 42 + 5, false);
          A.store.Bool(ptr + 42 + 2, false);
        } else {
          A.store.Bool(ptr + 42 + 6, true);
          A.store.Bool(ptr + 42 + 3, "enabled" in x["status"] ? true : false);
          A.store.Bool(ptr + 42 + 0, x["status"]["enabled"] ? true : false);
          A.store.Bool(ptr + 42 + 4, "owned" in x["status"] ? true : false);
          A.store.Bool(ptr + 42 + 1, x["status"]["owned"] ? true : false);
          A.store.Bool(ptr + 42 + 5, "ownerPasswordIsPresent" in x["status"] ? true : false);
          A.store.Bool(ptr + 42 + 2, x["status"]["ownerPasswordIsPresent"] ? true : false);
        }

        if (typeof x["dictionaryAttack"] === "undefined") {
          A.store.Bool(ptr + 52 + 20, false);
          A.store.Bool(ptr + 52 + 16, false);
          A.store.Int32(ptr + 52 + 0, 0);
          A.store.Bool(ptr + 52 + 17, false);
          A.store.Int32(ptr + 52 + 4, 0);
          A.store.Bool(ptr + 52 + 18, false);
          A.store.Bool(ptr + 52 + 8, false);
          A.store.Bool(ptr + 52 + 19, false);
          A.store.Int32(ptr + 52 + 12, 0);
        } else {
          A.store.Bool(ptr + 52 + 20, true);
          A.store.Bool(ptr + 52 + 16, "counter" in x["dictionaryAttack"] ? true : false);
          A.store.Int32(
            ptr + 52 + 0,
            x["dictionaryAttack"]["counter"] === undefined ? 0 : (x["dictionaryAttack"]["counter"] as number)
          );
          A.store.Bool(ptr + 52 + 17, "threshold" in x["dictionaryAttack"] ? true : false);
          A.store.Int32(
            ptr + 52 + 4,
            x["dictionaryAttack"]["threshold"] === undefined ? 0 : (x["dictionaryAttack"]["threshold"] as number)
          );
          A.store.Bool(ptr + 52 + 18, "lockoutInEffect" in x["dictionaryAttack"] ? true : false);
          A.store.Bool(ptr + 52 + 8, x["dictionaryAttack"]["lockoutInEffect"] ? true : false);
          A.store.Bool(ptr + 52 + 19, "lockoutSecondsRemaining" in x["dictionaryAttack"] ? true : false);
          A.store.Int32(
            ptr + 52 + 12,
            x["dictionaryAttack"]["lockoutSecondsRemaining"] === undefined
              ? 0
              : (x["dictionaryAttack"]["lockoutSecondsRemaining"] as number)
          );
        }
      }
    },
    "load_TpmInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 41)) {
        x["version"] = {};
        x["version"]["gscVersion"] = A.load.Enum(ptr + 0 + 0, ["not_gsc", "cr50", "ti50"]);
        if (A.load.Bool(ptr + 0 + 36)) {
          x["version"]["family"] = A.load.Int32(ptr + 0 + 4);
        } else {
          delete x["version"]["family"];
        }
        if (A.load.Bool(ptr + 0 + 37)) {
          x["version"]["specLevel"] = A.load.Float64(ptr + 0 + 8);
        } else {
          delete x["version"]["specLevel"];
        }
        if (A.load.Bool(ptr + 0 + 38)) {
          x["version"]["manufacturer"] = A.load.Int32(ptr + 0 + 16);
        } else {
          delete x["version"]["manufacturer"];
        }
        if (A.load.Bool(ptr + 0 + 39)) {
          x["version"]["tpmModel"] = A.load.Int32(ptr + 0 + 20);
        } else {
          delete x["version"]["tpmModel"];
        }
        if (A.load.Bool(ptr + 0 + 40)) {
          x["version"]["firmwareVersion"] = A.load.Float64(ptr + 0 + 24);
        } else {
          delete x["version"]["firmwareVersion"];
        }
        x["version"]["vendorSpecific"] = A.load.Ref(ptr + 0 + 32, undefined);
      } else {
        delete x["version"];
      }
      if (A.load.Bool(ptr + 42 + 6)) {
        x["status"] = {};
        if (A.load.Bool(ptr + 42 + 3)) {
          x["status"]["enabled"] = A.load.Bool(ptr + 42 + 0);
        } else {
          delete x["status"]["enabled"];
        }
        if (A.load.Bool(ptr + 42 + 4)) {
          x["status"]["owned"] = A.load.Bool(ptr + 42 + 1);
        } else {
          delete x["status"]["owned"];
        }
        if (A.load.Bool(ptr + 42 + 5)) {
          x["status"]["ownerPasswordIsPresent"] = A.load.Bool(ptr + 42 + 2);
        } else {
          delete x["status"]["ownerPasswordIsPresent"];
        }
      } else {
        delete x["status"];
      }
      if (A.load.Bool(ptr + 52 + 20)) {
        x["dictionaryAttack"] = {};
        if (A.load.Bool(ptr + 52 + 16)) {
          x["dictionaryAttack"]["counter"] = A.load.Int32(ptr + 52 + 0);
        } else {
          delete x["dictionaryAttack"]["counter"];
        }
        if (A.load.Bool(ptr + 52 + 17)) {
          x["dictionaryAttack"]["threshold"] = A.load.Int32(ptr + 52 + 4);
        } else {
          delete x["dictionaryAttack"]["threshold"];
        }
        if (A.load.Bool(ptr + 52 + 18)) {
          x["dictionaryAttack"]["lockoutInEffect"] = A.load.Bool(ptr + 52 + 8);
        } else {
          delete x["dictionaryAttack"]["lockoutInEffect"];
        }
        if (A.load.Bool(ptr + 52 + 19)) {
          x["dictionaryAttack"]["lockoutSecondsRemaining"] = A.load.Int32(ptr + 52 + 12);
        } else {
          delete x["dictionaryAttack"]["lockoutSecondsRemaining"];
        }
      } else {
        delete x["dictionaryAttack"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UsbBusInterfaceInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 40, false);
        A.store.Bool(ptr + 36, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 37, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 38, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 39, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Ref(ptr + 32, undefined);
      } else {
        A.store.Bool(ptr + 40, true);
        A.store.Bool(ptr + 36, "interfaceNumber" in x ? true : false);
        A.store.Float64(ptr + 0, x["interfaceNumber"] === undefined ? 0 : (x["interfaceNumber"] as number));
        A.store.Bool(ptr + 37, "classId" in x ? true : false);
        A.store.Float64(ptr + 8, x["classId"] === undefined ? 0 : (x["classId"] as number));
        A.store.Bool(ptr + 38, "subclassId" in x ? true : false);
        A.store.Float64(ptr + 16, x["subclassId"] === undefined ? 0 : (x["subclassId"] as number));
        A.store.Bool(ptr + 39, "protocolId" in x ? true : false);
        A.store.Float64(ptr + 24, x["protocolId"] === undefined ? 0 : (x["protocolId"] as number));
        A.store.Ref(ptr + 32, x["driver"]);
      }
    },
    "load_UsbBusInterfaceInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 36)) {
        x["interfaceNumber"] = A.load.Float64(ptr + 0);
      } else {
        delete x["interfaceNumber"];
      }
      if (A.load.Bool(ptr + 37)) {
        x["classId"] = A.load.Float64(ptr + 8);
      } else {
        delete x["classId"];
      }
      if (A.load.Bool(ptr + 38)) {
        x["subclassId"] = A.load.Float64(ptr + 16);
      } else {
        delete x["subclassId"];
      }
      if (A.load.Bool(ptr + 39)) {
        x["protocolId"] = A.load.Float64(ptr + 24);
      } else {
        delete x["protocolId"];
      }
      x["driver"] = A.load.Ref(ptr + 32, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_UsbVersion": (ref: heap.Ref<string>): number => {
      const idx = ["unknown", "usb1", "usb2", "usb3"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_UsbSpecSpeed": (ref: heap.Ref<string>): number => {
      const idx = ["unknown", "n1_5Mbps", "n12Mbps", "n480Mbps", "n5Gbps", "n10Gbps", "n20Gbps"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_UsbBusInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 69, false);
        A.store.Bool(ptr + 64, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 65, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 66, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 67, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Bool(ptr + 68, false);
        A.store.Float64(ptr + 32, 0);
        A.store.Ref(ptr + 40, undefined);

        A.store.Bool(ptr + 44 + 8, false);
        A.store.Ref(ptr + 44 + 0, undefined);
        A.store.Enum(ptr + 44 + 4, -1);
        A.store.Enum(ptr + 56, -1);
        A.store.Enum(ptr + 60, -1);
      } else {
        A.store.Bool(ptr + 69, true);
        A.store.Bool(ptr + 64, "classId" in x ? true : false);
        A.store.Float64(ptr + 0, x["classId"] === undefined ? 0 : (x["classId"] as number));
        A.store.Bool(ptr + 65, "subclassId" in x ? true : false);
        A.store.Float64(ptr + 8, x["subclassId"] === undefined ? 0 : (x["subclassId"] as number));
        A.store.Bool(ptr + 66, "protocolId" in x ? true : false);
        A.store.Float64(ptr + 16, x["protocolId"] === undefined ? 0 : (x["protocolId"] as number));
        A.store.Bool(ptr + 67, "vendorId" in x ? true : false);
        A.store.Float64(ptr + 24, x["vendorId"] === undefined ? 0 : (x["vendorId"] as number));
        A.store.Bool(ptr + 68, "productId" in x ? true : false);
        A.store.Float64(ptr + 32, x["productId"] === undefined ? 0 : (x["productId"] as number));
        A.store.Ref(ptr + 40, x["interfaces"]);

        if (typeof x["fwupdFirmwareVersionInfo"] === "undefined") {
          A.store.Bool(ptr + 44 + 8, false);
          A.store.Ref(ptr + 44 + 0, undefined);
          A.store.Enum(ptr + 44 + 4, -1);
        } else {
          A.store.Bool(ptr + 44 + 8, true);
          A.store.Ref(ptr + 44 + 0, x["fwupdFirmwareVersionInfo"]["version"]);
          A.store.Enum(
            ptr + 44 + 4,
            [
              "plain",
              "number",
              "pair",
              "triplet",
              "quad",
              "bcd",
              "intelMe",
              "intelMe2",
              "surfaceLegacy",
              "surface",
              "dellBios",
              "hex",
            ].indexOf(x["fwupdFirmwareVersionInfo"]["version_format"] as string)
          );
        }
        A.store.Enum(ptr + 56, ["unknown", "usb1", "usb2", "usb3"].indexOf(x["version"] as string));
        A.store.Enum(
          ptr + 60,
          ["unknown", "n1_5Mbps", "n12Mbps", "n480Mbps", "n5Gbps", "n10Gbps", "n20Gbps"].indexOf(
            x["spec_speed"] as string
          )
        );
      }
    },
    "load_UsbBusInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 64)) {
        x["classId"] = A.load.Float64(ptr + 0);
      } else {
        delete x["classId"];
      }
      if (A.load.Bool(ptr + 65)) {
        x["subclassId"] = A.load.Float64(ptr + 8);
      } else {
        delete x["subclassId"];
      }
      if (A.load.Bool(ptr + 66)) {
        x["protocolId"] = A.load.Float64(ptr + 16);
      } else {
        delete x["protocolId"];
      }
      if (A.load.Bool(ptr + 67)) {
        x["vendorId"] = A.load.Float64(ptr + 24);
      } else {
        delete x["vendorId"];
      }
      if (A.load.Bool(ptr + 68)) {
        x["productId"] = A.load.Float64(ptr + 32);
      } else {
        delete x["productId"];
      }
      x["interfaces"] = A.load.Ref(ptr + 40, undefined);
      if (A.load.Bool(ptr + 44 + 8)) {
        x["fwupdFirmwareVersionInfo"] = {};
        x["fwupdFirmwareVersionInfo"]["version"] = A.load.Ref(ptr + 44 + 0, undefined);
        x["fwupdFirmwareVersionInfo"]["version_format"] = A.load.Enum(ptr + 44 + 4, [
          "plain",
          "number",
          "pair",
          "triplet",
          "quad",
          "bcd",
          "intelMe",
          "intelMe2",
          "surfaceLegacy",
          "surface",
          "dellBios",
          "hex",
        ]);
      } else {
        delete x["fwupdFirmwareVersionInfo"];
      }
      x["version"] = A.load.Enum(ptr + 56, ["unknown", "usb1", "usb2", "usb3"]);
      x["spec_speed"] = A.load.Enum(ptr + 60, [
        "unknown",
        "n1_5Mbps",
        "n12Mbps",
        "n480Mbps",
        "n5Gbps",
        "n10Gbps",
        "n20Gbps",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_UsbBusDevices": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["devices"]);
      }
    },
    "load_UsbBusDevices": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["devices"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_VpdInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["activateDate"]);
        A.store.Ref(ptr + 4, x["modelName"]);
        A.store.Ref(ptr + 8, x["serialNumber"]);
        A.store.Ref(ptr + 12, x["skuNumber"]);
      }
    },
    "load_VpdInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["activateDate"] = A.load.Ref(ptr + 0, undefined);
      x["modelName"] = A.load.Ref(ptr + 4, undefined);
      x["serialNumber"] = A.load.Ref(ptr + 8, undefined);
      x["skuNumber"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetAudioInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.telemetry && "getAudioInfo" in WEBEXT?.os?.telemetry) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAudioInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.telemetry.getAudioInfo);
    },
    "call_GetAudioInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.telemetry.getAudioInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAudioInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.telemetry.getAudioInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetBatteryInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.telemetry && "getBatteryInfo" in WEBEXT?.os?.telemetry) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetBatteryInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.telemetry.getBatteryInfo);
    },
    "call_GetBatteryInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.telemetry.getBatteryInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetBatteryInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.telemetry.getBatteryInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCpuInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.telemetry && "getCpuInfo" in WEBEXT?.os?.telemetry) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCpuInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.telemetry.getCpuInfo);
    },
    "call_GetCpuInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.telemetry.getCpuInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCpuInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.telemetry.getCpuInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDisplayInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.telemetry && "getDisplayInfo" in WEBEXT?.os?.telemetry) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDisplayInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.telemetry.getDisplayInfo);
    },
    "call_GetDisplayInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.telemetry.getDisplayInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDisplayInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.telemetry.getDisplayInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetInternetConnectivityInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.telemetry && "getInternetConnectivityInfo" in WEBEXT?.os?.telemetry) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetInternetConnectivityInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.telemetry.getInternetConnectivityInfo);
    },
    "call_GetInternetConnectivityInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.telemetry.getInternetConnectivityInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetInternetConnectivityInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.telemetry.getInternetConnectivityInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetMarketingInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.telemetry && "getMarketingInfo" in WEBEXT?.os?.telemetry) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetMarketingInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.telemetry.getMarketingInfo);
    },
    "call_GetMarketingInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.telemetry.getMarketingInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetMarketingInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.telemetry.getMarketingInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetMemoryInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.telemetry && "getMemoryInfo" in WEBEXT?.os?.telemetry) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetMemoryInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.telemetry.getMemoryInfo);
    },
    "call_GetMemoryInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.telemetry.getMemoryInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetMemoryInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.telemetry.getMemoryInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetNonRemovableBlockDevicesInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.telemetry && "getNonRemovableBlockDevicesInfo" in WEBEXT?.os?.telemetry) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetNonRemovableBlockDevicesInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.telemetry.getNonRemovableBlockDevicesInfo);
    },
    "call_GetNonRemovableBlockDevicesInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.telemetry.getNonRemovableBlockDevicesInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetNonRemovableBlockDevicesInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.telemetry.getNonRemovableBlockDevicesInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetOemData": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.telemetry && "getOemData" in WEBEXT?.os?.telemetry) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetOemData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.telemetry.getOemData);
    },
    "call_GetOemData": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.telemetry.getOemData();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetOemData": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.telemetry.getOemData();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetOsVersionInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.telemetry && "getOsVersionInfo" in WEBEXT?.os?.telemetry) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetOsVersionInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.telemetry.getOsVersionInfo);
    },
    "call_GetOsVersionInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.telemetry.getOsVersionInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetOsVersionInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.telemetry.getOsVersionInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetStatefulPartitionInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.telemetry && "getStatefulPartitionInfo" in WEBEXT?.os?.telemetry) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetStatefulPartitionInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.telemetry.getStatefulPartitionInfo);
    },
    "call_GetStatefulPartitionInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.telemetry.getStatefulPartitionInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetStatefulPartitionInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.telemetry.getStatefulPartitionInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetTpmInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.telemetry && "getTpmInfo" in WEBEXT?.os?.telemetry) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetTpmInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.telemetry.getTpmInfo);
    },
    "call_GetTpmInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.telemetry.getTpmInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetTpmInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.telemetry.getTpmInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetUsbBusInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.telemetry && "getUsbBusInfo" in WEBEXT?.os?.telemetry) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetUsbBusInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.telemetry.getUsbBusInfo);
    },
    "call_GetUsbBusInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.telemetry.getUsbBusInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetUsbBusInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.telemetry.getUsbBusInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetVpdInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.os?.telemetry && "getVpdInfo" in WEBEXT?.os?.telemetry) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetVpdInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.os.telemetry.getVpdInfo);
    },
    "call_GetVpdInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.os.telemetry.getVpdInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetVpdInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.os.telemetry.getVpdInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
