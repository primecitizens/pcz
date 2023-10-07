import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/mediaperceptionprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AudioSpectrogram": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["values"]);
      }
    },
    "load_AudioSpectrogram": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["values"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AudioHumanPresenceDetection": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 22, false);
        A.store.Bool(ptr + 21, false);
        A.store.Float64(ptr + 0, 0);

        A.store.Bool(ptr + 8 + 4, false);
        A.store.Ref(ptr + 8 + 0, undefined);

        A.store.Bool(ptr + 16 + 4, false);
        A.store.Ref(ptr + 16 + 0, undefined);
      } else {
        A.store.Bool(ptr + 22, true);
        A.store.Bool(ptr + 21, "humanPresenceLikelihood" in x ? true : false);
        A.store.Float64(
          ptr + 0,
          x["humanPresenceLikelihood"] === undefined ? 0 : (x["humanPresenceLikelihood"] as number)
        );

        if (typeof x["noiseSpectrogram"] === "undefined") {
          A.store.Bool(ptr + 8 + 4, false);
          A.store.Ref(ptr + 8 + 0, undefined);
        } else {
          A.store.Bool(ptr + 8 + 4, true);
          A.store.Ref(ptr + 8 + 0, x["noiseSpectrogram"]["values"]);
        }

        if (typeof x["frameSpectrogram"] === "undefined") {
          A.store.Bool(ptr + 16 + 4, false);
          A.store.Ref(ptr + 16 + 0, undefined);
        } else {
          A.store.Bool(ptr + 16 + 4, true);
          A.store.Ref(ptr + 16 + 0, x["frameSpectrogram"]["values"]);
        }
      }
    },
    "load_AudioHumanPresenceDetection": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 21)) {
        x["humanPresenceLikelihood"] = A.load.Float64(ptr + 0);
      } else {
        delete x["humanPresenceLikelihood"];
      }
      if (A.load.Bool(ptr + 8 + 4)) {
        x["noiseSpectrogram"] = {};
        x["noiseSpectrogram"]["values"] = A.load.Ref(ptr + 8 + 0, undefined);
      } else {
        delete x["noiseSpectrogram"];
      }
      if (A.load.Bool(ptr + 16 + 4)) {
        x["frameSpectrogram"] = {};
        x["frameSpectrogram"]["values"] = A.load.Ref(ptr + 16 + 0, undefined);
      } else {
        delete x["frameSpectrogram"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AudioLocalization": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "azimuthRadians" in x ? true : false);
        A.store.Float64(ptr + 0, x["azimuthRadians"] === undefined ? 0 : (x["azimuthRadians"] as number));
        A.store.Ref(ptr + 8, x["azimuthScores"]);
      }
    },
    "load_AudioLocalization": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["azimuthRadians"] = A.load.Float64(ptr + 0);
      } else {
        delete x["azimuthRadians"];
      }
      x["azimuthScores"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_HotwordType": (ref: heap.Ref<string>): number => {
      const idx = ["UNKNOWN_TYPE", "OK_GOOGLE"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Hotword": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 37, false);
        A.store.Bool(ptr + 32, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Enum(ptr + 4, -1);
        A.store.Bool(ptr + 33, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 34, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 35, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Bool(ptr + 36, false);
        A.store.Float64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 37, true);
        A.store.Bool(ptr + 32, "id" in x ? true : false);
        A.store.Int32(ptr + 0, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Enum(ptr + 4, ["UNKNOWN_TYPE", "OK_GOOGLE"].indexOf(x["type"] as string));
        A.store.Bool(ptr + 33, "frameId" in x ? true : false);
        A.store.Int32(ptr + 8, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Bool(ptr + 34, "startTimestampMs" in x ? true : false);
        A.store.Int32(ptr + 12, x["startTimestampMs"] === undefined ? 0 : (x["startTimestampMs"] as number));
        A.store.Bool(ptr + 35, "endTimestampMs" in x ? true : false);
        A.store.Int32(ptr + 16, x["endTimestampMs"] === undefined ? 0 : (x["endTimestampMs"] as number));
        A.store.Bool(ptr + 36, "confidence" in x ? true : false);
        A.store.Float64(ptr + 24, x["confidence"] === undefined ? 0 : (x["confidence"] as number));
      }
    },
    "load_Hotword": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 32)) {
        x["id"] = A.load.Int32(ptr + 0);
      } else {
        delete x["id"];
      }
      x["type"] = A.load.Enum(ptr + 4, ["UNKNOWN_TYPE", "OK_GOOGLE"]);
      if (A.load.Bool(ptr + 33)) {
        x["frameId"] = A.load.Int32(ptr + 8);
      } else {
        delete x["frameId"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["startTimestampMs"] = A.load.Int32(ptr + 12);
      } else {
        delete x["startTimestampMs"];
      }
      if (A.load.Bool(ptr + 35)) {
        x["endTimestampMs"] = A.load.Int32(ptr + 16);
      } else {
        delete x["endTimestampMs"];
      }
      if (A.load.Bool(ptr + 36)) {
        x["confidence"] = A.load.Float64(ptr + 24);
      } else {
        delete x["confidence"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_HotwordDetection": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["hotwords"]);
      }
    },
    "load_HotwordDetection": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["hotwords"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AudioPerception": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 54, false);
        A.store.Bool(ptr + 53, false);
        A.store.Float64(ptr + 0, 0);

        A.store.Bool(ptr + 8 + 13, false);
        A.store.Bool(ptr + 8 + 12, false);
        A.store.Float64(ptr + 8 + 0, 0);
        A.store.Ref(ptr + 8 + 8, undefined);

        A.store.Bool(ptr + 24 + 22, false);
        A.store.Bool(ptr + 24 + 21, false);
        A.store.Float64(ptr + 24 + 0, 0);

        A.store.Bool(ptr + 24 + 8 + 4, false);
        A.store.Ref(ptr + 24 + 8 + 0, undefined);

        A.store.Bool(ptr + 24 + 16 + 4, false);
        A.store.Ref(ptr + 24 + 16 + 0, undefined);

        A.store.Bool(ptr + 48 + 4, false);
        A.store.Ref(ptr + 48 + 0, undefined);
      } else {
        A.store.Bool(ptr + 54, true);
        A.store.Bool(ptr + 53, "timestampUs" in x ? true : false);
        A.store.Float64(ptr + 0, x["timestampUs"] === undefined ? 0 : (x["timestampUs"] as number));

        if (typeof x["audioLocalization"] === "undefined") {
          A.store.Bool(ptr + 8 + 13, false);
          A.store.Bool(ptr + 8 + 12, false);
          A.store.Float64(ptr + 8 + 0, 0);
          A.store.Ref(ptr + 8 + 8, undefined);
        } else {
          A.store.Bool(ptr + 8 + 13, true);
          A.store.Bool(ptr + 8 + 12, "azimuthRadians" in x["audioLocalization"] ? true : false);
          A.store.Float64(
            ptr + 8 + 0,
            x["audioLocalization"]["azimuthRadians"] === undefined
              ? 0
              : (x["audioLocalization"]["azimuthRadians"] as number)
          );
          A.store.Ref(ptr + 8 + 8, x["audioLocalization"]["azimuthScores"]);
        }

        if (typeof x["audioHumanPresenceDetection"] === "undefined") {
          A.store.Bool(ptr + 24 + 22, false);
          A.store.Bool(ptr + 24 + 21, false);
          A.store.Float64(ptr + 24 + 0, 0);

          A.store.Bool(ptr + 24 + 8 + 4, false);
          A.store.Ref(ptr + 24 + 8 + 0, undefined);

          A.store.Bool(ptr + 24 + 16 + 4, false);
          A.store.Ref(ptr + 24 + 16 + 0, undefined);
        } else {
          A.store.Bool(ptr + 24 + 22, true);
          A.store.Bool(ptr + 24 + 21, "humanPresenceLikelihood" in x["audioHumanPresenceDetection"] ? true : false);
          A.store.Float64(
            ptr + 24 + 0,
            x["audioHumanPresenceDetection"]["humanPresenceLikelihood"] === undefined
              ? 0
              : (x["audioHumanPresenceDetection"]["humanPresenceLikelihood"] as number)
          );

          if (typeof x["audioHumanPresenceDetection"]["noiseSpectrogram"] === "undefined") {
            A.store.Bool(ptr + 24 + 8 + 4, false);
            A.store.Ref(ptr + 24 + 8 + 0, undefined);
          } else {
            A.store.Bool(ptr + 24 + 8 + 4, true);
            A.store.Ref(ptr + 24 + 8 + 0, x["audioHumanPresenceDetection"]["noiseSpectrogram"]["values"]);
          }

          if (typeof x["audioHumanPresenceDetection"]["frameSpectrogram"] === "undefined") {
            A.store.Bool(ptr + 24 + 16 + 4, false);
            A.store.Ref(ptr + 24 + 16 + 0, undefined);
          } else {
            A.store.Bool(ptr + 24 + 16 + 4, true);
            A.store.Ref(ptr + 24 + 16 + 0, x["audioHumanPresenceDetection"]["frameSpectrogram"]["values"]);
          }
        }

        if (typeof x["hotwordDetection"] === "undefined") {
          A.store.Bool(ptr + 48 + 4, false);
          A.store.Ref(ptr + 48 + 0, undefined);
        } else {
          A.store.Bool(ptr + 48 + 4, true);
          A.store.Ref(ptr + 48 + 0, x["hotwordDetection"]["hotwords"]);
        }
      }
    },
    "load_AudioPerception": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 53)) {
        x["timestampUs"] = A.load.Float64(ptr + 0);
      } else {
        delete x["timestampUs"];
      }
      if (A.load.Bool(ptr + 8 + 13)) {
        x["audioLocalization"] = {};
        if (A.load.Bool(ptr + 8 + 12)) {
          x["audioLocalization"]["azimuthRadians"] = A.load.Float64(ptr + 8 + 0);
        } else {
          delete x["audioLocalization"]["azimuthRadians"];
        }
        x["audioLocalization"]["azimuthScores"] = A.load.Ref(ptr + 8 + 8, undefined);
      } else {
        delete x["audioLocalization"];
      }
      if (A.load.Bool(ptr + 24 + 22)) {
        x["audioHumanPresenceDetection"] = {};
        if (A.load.Bool(ptr + 24 + 21)) {
          x["audioHumanPresenceDetection"]["humanPresenceLikelihood"] = A.load.Float64(ptr + 24 + 0);
        } else {
          delete x["audioHumanPresenceDetection"]["humanPresenceLikelihood"];
        }
        if (A.load.Bool(ptr + 24 + 8 + 4)) {
          x["audioHumanPresenceDetection"]["noiseSpectrogram"] = {};
          x["audioHumanPresenceDetection"]["noiseSpectrogram"]["values"] = A.load.Ref(ptr + 24 + 8 + 0, undefined);
        } else {
          delete x["audioHumanPresenceDetection"]["noiseSpectrogram"];
        }
        if (A.load.Bool(ptr + 24 + 16 + 4)) {
          x["audioHumanPresenceDetection"]["frameSpectrogram"] = {};
          x["audioHumanPresenceDetection"]["frameSpectrogram"]["values"] = A.load.Ref(ptr + 24 + 16 + 0, undefined);
        } else {
          delete x["audioHumanPresenceDetection"]["frameSpectrogram"];
        }
      } else {
        delete x["audioHumanPresenceDetection"];
      }
      if (A.load.Bool(ptr + 48 + 4)) {
        x["hotwordDetection"] = {};
        x["hotwordDetection"]["hotwords"] = A.load.Ref(ptr + 48 + 0, undefined);
      } else {
        delete x["hotwordDetection"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AudioVisualHumanPresenceDetection": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Float64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "humanPresenceLikelihood" in x ? true : false);
        A.store.Float64(
          ptr + 0,
          x["humanPresenceLikelihood"] === undefined ? 0 : (x["humanPresenceLikelihood"] as number)
        );
      }
    },
    "load_AudioVisualHumanPresenceDetection": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["humanPresenceLikelihood"] = A.load.Float64(ptr + 0);
      } else {
        delete x["humanPresenceLikelihood"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AudioVisualPerception": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 18, false);
        A.store.Float64(ptr + 0, 0);

        A.store.Bool(ptr + 8 + 9, false);
        A.store.Bool(ptr + 8 + 8, false);
        A.store.Float64(ptr + 8 + 0, 0);
      } else {
        A.store.Bool(ptr + 19, true);
        A.store.Bool(ptr + 18, "timestampUs" in x ? true : false);
        A.store.Float64(ptr + 0, x["timestampUs"] === undefined ? 0 : (x["timestampUs"] as number));

        if (typeof x["audioVisualHumanPresenceDetection"] === "undefined") {
          A.store.Bool(ptr + 8 + 9, false);
          A.store.Bool(ptr + 8 + 8, false);
          A.store.Float64(ptr + 8 + 0, 0);
        } else {
          A.store.Bool(ptr + 8 + 9, true);
          A.store.Bool(ptr + 8 + 8, "humanPresenceLikelihood" in x["audioVisualHumanPresenceDetection"] ? true : false);
          A.store.Float64(
            ptr + 8 + 0,
            x["audioVisualHumanPresenceDetection"]["humanPresenceLikelihood"] === undefined
              ? 0
              : (x["audioVisualHumanPresenceDetection"]["humanPresenceLikelihood"] as number)
          );
        }
      }
    },
    "load_AudioVisualPerception": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 18)) {
        x["timestampUs"] = A.load.Float64(ptr + 0);
      } else {
        delete x["timestampUs"];
      }
      if (A.load.Bool(ptr + 8 + 9)) {
        x["audioVisualHumanPresenceDetection"] = {};
        if (A.load.Bool(ptr + 8 + 8)) {
          x["audioVisualHumanPresenceDetection"]["humanPresenceLikelihood"] = A.load.Float64(ptr + 8 + 0);
        } else {
          delete x["audioVisualHumanPresenceDetection"]["humanPresenceLikelihood"];
        }
      } else {
        delete x["audioVisualHumanPresenceDetection"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Point": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Bool(ptr + 16, "x" in x ? true : false);
        A.store.Float64(ptr + 0, x["x"] === undefined ? 0 : (x["x"] as number));
        A.store.Bool(ptr + 17, "y" in x ? true : false);
        A.store.Float64(ptr + 8, x["y"] === undefined ? 0 : (x["y"] as number));
      }
    },
    "load_Point": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["x"] = A.load.Float64(ptr + 0);
      } else {
        delete x["x"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["y"] = A.load.Float64(ptr + 8);
      } else {
        delete x["y"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_BoundingBox": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 52, false);
        A.store.Bool(ptr + 51, false);
        A.store.Bool(ptr + 0, false);

        A.store.Bool(ptr + 8 + 18, false);
        A.store.Bool(ptr + 8 + 16, false);
        A.store.Float64(ptr + 8 + 0, 0);
        A.store.Bool(ptr + 8 + 17, false);
        A.store.Float64(ptr + 8 + 8, 0);

        A.store.Bool(ptr + 32 + 18, false);
        A.store.Bool(ptr + 32 + 16, false);
        A.store.Float64(ptr + 32 + 0, 0);
        A.store.Bool(ptr + 32 + 17, false);
        A.store.Float64(ptr + 32 + 8, 0);
      } else {
        A.store.Bool(ptr + 52, true);
        A.store.Bool(ptr + 51, "normalized" in x ? true : false);
        A.store.Bool(ptr + 0, x["normalized"] ? true : false);

        if (typeof x["topLeft"] === "undefined") {
          A.store.Bool(ptr + 8 + 18, false);
          A.store.Bool(ptr + 8 + 16, false);
          A.store.Float64(ptr + 8 + 0, 0);
          A.store.Bool(ptr + 8 + 17, false);
          A.store.Float64(ptr + 8 + 8, 0);
        } else {
          A.store.Bool(ptr + 8 + 18, true);
          A.store.Bool(ptr + 8 + 16, "x" in x["topLeft"] ? true : false);
          A.store.Float64(ptr + 8 + 0, x["topLeft"]["x"] === undefined ? 0 : (x["topLeft"]["x"] as number));
          A.store.Bool(ptr + 8 + 17, "y" in x["topLeft"] ? true : false);
          A.store.Float64(ptr + 8 + 8, x["topLeft"]["y"] === undefined ? 0 : (x["topLeft"]["y"] as number));
        }

        if (typeof x["bottomRight"] === "undefined") {
          A.store.Bool(ptr + 32 + 18, false);
          A.store.Bool(ptr + 32 + 16, false);
          A.store.Float64(ptr + 32 + 0, 0);
          A.store.Bool(ptr + 32 + 17, false);
          A.store.Float64(ptr + 32 + 8, 0);
        } else {
          A.store.Bool(ptr + 32 + 18, true);
          A.store.Bool(ptr + 32 + 16, "x" in x["bottomRight"] ? true : false);
          A.store.Float64(ptr + 32 + 0, x["bottomRight"]["x"] === undefined ? 0 : (x["bottomRight"]["x"] as number));
          A.store.Bool(ptr + 32 + 17, "y" in x["bottomRight"] ? true : false);
          A.store.Float64(ptr + 32 + 8, x["bottomRight"]["y"] === undefined ? 0 : (x["bottomRight"]["y"] as number));
        }
      }
    },
    "load_BoundingBox": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 51)) {
        x["normalized"] = A.load.Bool(ptr + 0);
      } else {
        delete x["normalized"];
      }
      if (A.load.Bool(ptr + 8 + 18)) {
        x["topLeft"] = {};
        if (A.load.Bool(ptr + 8 + 16)) {
          x["topLeft"]["x"] = A.load.Float64(ptr + 8 + 0);
        } else {
          delete x["topLeft"]["x"];
        }
        if (A.load.Bool(ptr + 8 + 17)) {
          x["topLeft"]["y"] = A.load.Float64(ptr + 8 + 8);
        } else {
          delete x["topLeft"]["y"];
        }
      } else {
        delete x["topLeft"];
      }
      if (A.load.Bool(ptr + 32 + 18)) {
        x["bottomRight"] = {};
        if (A.load.Bool(ptr + 32 + 16)) {
          x["bottomRight"]["x"] = A.load.Float64(ptr + 32 + 0);
        } else {
          delete x["bottomRight"]["x"];
        }
        if (A.load.Bool(ptr + 32 + 17)) {
          x["bottomRight"]["y"] = A.load.Float64(ptr + 32 + 8);
        } else {
          delete x["bottomRight"]["y"];
        }
      } else {
        delete x["bottomRight"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ComponentType": (ref: heap.Ref<string>): number => {
      const idx = ["LIGHT", "FULL"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Component": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["LIGHT", "FULL"].indexOf(x["type"] as string));
      }
    },
    "load_Component": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, ["LIGHT", "FULL"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ComponentInstallationError": (ref: heap.Ref<string>): number => {
      const idx = [
        "UNKNOWN_COMPONENT",
        "INSTALL_FAILURE",
        "MOUNT_FAILURE",
        "COMPATIBILITY_CHECK_FAILED",
        "NOT_FOUND",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ComponentStatus": (ref: heap.Ref<string>): number => {
      const idx = ["UNKNOWN", "INSTALLED", "FAILED_TO_INSTALL"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ComponentState": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Enum(ptr + 0, ["UNKNOWN", "INSTALLED", "FAILED_TO_INSTALL"].indexOf(x["status"] as string));
        A.store.Ref(ptr + 4, x["version"]);
        A.store.Enum(
          ptr + 8,
          ["UNKNOWN_COMPONENT", "INSTALL_FAILURE", "MOUNT_FAILURE", "COMPATIBILITY_CHECK_FAILED", "NOT_FOUND"].indexOf(
            x["installationErrorCode"] as string
          )
        );
      }
    },
    "load_ComponentState": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["status"] = A.load.Enum(ptr + 0, ["UNKNOWN", "INSTALLED", "FAILED_TO_INSTALL"]);
      x["version"] = A.load.Ref(ptr + 4, undefined);
      x["installationErrorCode"] = A.load.Enum(ptr + 8, [
        "UNKNOWN_COMPONENT",
        "INSTALL_FAILURE",
        "MOUNT_FAILURE",
        "COMPATIBILITY_CHECK_FAILED",
        "NOT_FOUND",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ServiceError": (ref: heap.Ref<string>): number => {
      const idx = [
        "SERVICE_UNREACHABLE",
        "SERVICE_NOT_RUNNING",
        "SERVICE_BUSY_LAUNCHING",
        "SERVICE_NOT_INSTALLED",
        "MOJO_CONNECTION_FAILURE",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_EntityType": (ref: heap.Ref<string>): number => {
      const idx = ["UNSPECIFIED", "FACE", "PERSON", "MOTION_REGION", "LABELED_REGION"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_DistanceUnits": (ref: heap.Ref<string>): number => {
      const idx = ["UNSPECIFIED", "METERS", "PIXELS"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Distance": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Enum(ptr + 0, ["UNSPECIFIED", "METERS", "PIXELS"].indexOf(x["units"] as string));
        A.store.Bool(ptr + 16, "magnitude" in x ? true : false);
        A.store.Float64(ptr + 8, x["magnitude"] === undefined ? 0 : (x["magnitude"] as number));
      }
    },
    "load_Distance": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["units"] = A.load.Enum(ptr + 0, ["UNSPECIFIED", "METERS", "PIXELS"]);
      if (A.load.Bool(ptr + 16)) {
        x["magnitude"] = A.load.Float64(ptr + 8);
      } else {
        delete x["magnitude"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Entity": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 100, false);
        A.store.Bool(ptr + 98, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Enum(ptr + 4, -1);
        A.store.Ref(ptr + 8, undefined);

        A.store.Bool(ptr + 16 + 52, false);
        A.store.Bool(ptr + 16 + 51, false);
        A.store.Bool(ptr + 16 + 0, false);

        A.store.Bool(ptr + 16 + 8 + 18, false);
        A.store.Bool(ptr + 16 + 8 + 16, false);
        A.store.Float64(ptr + 16 + 8 + 0, 0);
        A.store.Bool(ptr + 16 + 8 + 17, false);
        A.store.Float64(ptr + 16 + 8 + 8, 0);

        A.store.Bool(ptr + 16 + 32 + 18, false);
        A.store.Bool(ptr + 16 + 32 + 16, false);
        A.store.Float64(ptr + 16 + 32 + 0, 0);
        A.store.Bool(ptr + 16 + 32 + 17, false);
        A.store.Float64(ptr + 16 + 32 + 8, 0);
        A.store.Bool(ptr + 99, false);
        A.store.Float64(ptr + 72, 0);

        A.store.Bool(ptr + 80 + 17, false);
        A.store.Enum(ptr + 80 + 0, -1);
        A.store.Bool(ptr + 80 + 16, false);
        A.store.Float64(ptr + 80 + 8, 0);
      } else {
        A.store.Bool(ptr + 100, true);
        A.store.Bool(ptr + 98, "id" in x ? true : false);
        A.store.Int32(ptr + 0, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Enum(
          ptr + 4,
          ["UNSPECIFIED", "FACE", "PERSON", "MOTION_REGION", "LABELED_REGION"].indexOf(x["type"] as string)
        );
        A.store.Ref(ptr + 8, x["entityLabel"]);

        if (typeof x["boundingBox"] === "undefined") {
          A.store.Bool(ptr + 16 + 52, false);
          A.store.Bool(ptr + 16 + 51, false);
          A.store.Bool(ptr + 16 + 0, false);

          A.store.Bool(ptr + 16 + 8 + 18, false);
          A.store.Bool(ptr + 16 + 8 + 16, false);
          A.store.Float64(ptr + 16 + 8 + 0, 0);
          A.store.Bool(ptr + 16 + 8 + 17, false);
          A.store.Float64(ptr + 16 + 8 + 8, 0);

          A.store.Bool(ptr + 16 + 32 + 18, false);
          A.store.Bool(ptr + 16 + 32 + 16, false);
          A.store.Float64(ptr + 16 + 32 + 0, 0);
          A.store.Bool(ptr + 16 + 32 + 17, false);
          A.store.Float64(ptr + 16 + 32 + 8, 0);
        } else {
          A.store.Bool(ptr + 16 + 52, true);
          A.store.Bool(ptr + 16 + 51, "normalized" in x["boundingBox"] ? true : false);
          A.store.Bool(ptr + 16 + 0, x["boundingBox"]["normalized"] ? true : false);

          if (typeof x["boundingBox"]["topLeft"] === "undefined") {
            A.store.Bool(ptr + 16 + 8 + 18, false);
            A.store.Bool(ptr + 16 + 8 + 16, false);
            A.store.Float64(ptr + 16 + 8 + 0, 0);
            A.store.Bool(ptr + 16 + 8 + 17, false);
            A.store.Float64(ptr + 16 + 8 + 8, 0);
          } else {
            A.store.Bool(ptr + 16 + 8 + 18, true);
            A.store.Bool(ptr + 16 + 8 + 16, "x" in x["boundingBox"]["topLeft"] ? true : false);
            A.store.Float64(
              ptr + 16 + 8 + 0,
              x["boundingBox"]["topLeft"]["x"] === undefined ? 0 : (x["boundingBox"]["topLeft"]["x"] as number)
            );
            A.store.Bool(ptr + 16 + 8 + 17, "y" in x["boundingBox"]["topLeft"] ? true : false);
            A.store.Float64(
              ptr + 16 + 8 + 8,
              x["boundingBox"]["topLeft"]["y"] === undefined ? 0 : (x["boundingBox"]["topLeft"]["y"] as number)
            );
          }

          if (typeof x["boundingBox"]["bottomRight"] === "undefined") {
            A.store.Bool(ptr + 16 + 32 + 18, false);
            A.store.Bool(ptr + 16 + 32 + 16, false);
            A.store.Float64(ptr + 16 + 32 + 0, 0);
            A.store.Bool(ptr + 16 + 32 + 17, false);
            A.store.Float64(ptr + 16 + 32 + 8, 0);
          } else {
            A.store.Bool(ptr + 16 + 32 + 18, true);
            A.store.Bool(ptr + 16 + 32 + 16, "x" in x["boundingBox"]["bottomRight"] ? true : false);
            A.store.Float64(
              ptr + 16 + 32 + 0,
              x["boundingBox"]["bottomRight"]["x"] === undefined ? 0 : (x["boundingBox"]["bottomRight"]["x"] as number)
            );
            A.store.Bool(ptr + 16 + 32 + 17, "y" in x["boundingBox"]["bottomRight"] ? true : false);
            A.store.Float64(
              ptr + 16 + 32 + 8,
              x["boundingBox"]["bottomRight"]["y"] === undefined ? 0 : (x["boundingBox"]["bottomRight"]["y"] as number)
            );
          }
        }
        A.store.Bool(ptr + 99, "confidence" in x ? true : false);
        A.store.Float64(ptr + 72, x["confidence"] === undefined ? 0 : (x["confidence"] as number));

        if (typeof x["depth"] === "undefined") {
          A.store.Bool(ptr + 80 + 17, false);
          A.store.Enum(ptr + 80 + 0, -1);
          A.store.Bool(ptr + 80 + 16, false);
          A.store.Float64(ptr + 80 + 8, 0);
        } else {
          A.store.Bool(ptr + 80 + 17, true);
          A.store.Enum(ptr + 80 + 0, ["UNSPECIFIED", "METERS", "PIXELS"].indexOf(x["depth"]["units"] as string));
          A.store.Bool(ptr + 80 + 16, "magnitude" in x["depth"] ? true : false);
          A.store.Float64(
            ptr + 80 + 8,
            x["depth"]["magnitude"] === undefined ? 0 : (x["depth"]["magnitude"] as number)
          );
        }
      }
    },
    "load_Entity": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 98)) {
        x["id"] = A.load.Int32(ptr + 0);
      } else {
        delete x["id"];
      }
      x["type"] = A.load.Enum(ptr + 4, ["UNSPECIFIED", "FACE", "PERSON", "MOTION_REGION", "LABELED_REGION"]);
      x["entityLabel"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 16 + 52)) {
        x["boundingBox"] = {};
        if (A.load.Bool(ptr + 16 + 51)) {
          x["boundingBox"]["normalized"] = A.load.Bool(ptr + 16 + 0);
        } else {
          delete x["boundingBox"]["normalized"];
        }
        if (A.load.Bool(ptr + 16 + 8 + 18)) {
          x["boundingBox"]["topLeft"] = {};
          if (A.load.Bool(ptr + 16 + 8 + 16)) {
            x["boundingBox"]["topLeft"]["x"] = A.load.Float64(ptr + 16 + 8 + 0);
          } else {
            delete x["boundingBox"]["topLeft"]["x"];
          }
          if (A.load.Bool(ptr + 16 + 8 + 17)) {
            x["boundingBox"]["topLeft"]["y"] = A.load.Float64(ptr + 16 + 8 + 8);
          } else {
            delete x["boundingBox"]["topLeft"]["y"];
          }
        } else {
          delete x["boundingBox"]["topLeft"];
        }
        if (A.load.Bool(ptr + 16 + 32 + 18)) {
          x["boundingBox"]["bottomRight"] = {};
          if (A.load.Bool(ptr + 16 + 32 + 16)) {
            x["boundingBox"]["bottomRight"]["x"] = A.load.Float64(ptr + 16 + 32 + 0);
          } else {
            delete x["boundingBox"]["bottomRight"]["x"];
          }
          if (A.load.Bool(ptr + 16 + 32 + 17)) {
            x["boundingBox"]["bottomRight"]["y"] = A.load.Float64(ptr + 16 + 32 + 8);
          } else {
            delete x["boundingBox"]["bottomRight"]["y"];
          }
        } else {
          delete x["boundingBox"]["bottomRight"];
        }
      } else {
        delete x["boundingBox"];
      }
      if (A.load.Bool(ptr + 99)) {
        x["confidence"] = A.load.Float64(ptr + 72);
      } else {
        delete x["confidence"];
      }
      if (A.load.Bool(ptr + 80 + 17)) {
        x["depth"] = {};
        x["depth"]["units"] = A.load.Enum(ptr + 80 + 0, ["UNSPECIFIED", "METERS", "PIXELS"]);
        if (A.load.Bool(ptr + 80 + 16)) {
          x["depth"]["magnitude"] = A.load.Float64(ptr + 80 + 8);
        } else {
          delete x["depth"]["magnitude"];
        }
      } else {
        delete x["depth"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PacketLatency": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Ref(ptr + 0, x["packetLabel"]);
        A.store.Bool(ptr + 8, "latencyUsec" in x ? true : false);
        A.store.Int32(ptr + 4, x["latencyUsec"] === undefined ? 0 : (x["latencyUsec"] as number));
      }
    },
    "load_PacketLatency": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["packetLabel"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 8)) {
        x["latencyUsec"] = A.load.Int32(ptr + 4);
      } else {
        delete x["latencyUsec"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_LightCondition": (ref: heap.Ref<string>): number => {
      const idx = ["UNSPECIFIED", "NO_CHANGE", "TURNED_ON", "TURNED_OFF", "DIMMER", "BRIGHTER", "BLACK_FRAME"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },

    "store_VideoHumanPresenceDetection": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 35, false);
        A.store.Bool(ptr + 32, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 33, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Bool(ptr + 34, false);
        A.store.Float64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 35, true);
        A.store.Bool(ptr + 32, "humanPresenceLikelihood" in x ? true : false);
        A.store.Float64(
          ptr + 0,
          x["humanPresenceLikelihood"] === undefined ? 0 : (x["humanPresenceLikelihood"] as number)
        );
        A.store.Bool(ptr + 33, "motionDetectedLikelihood" in x ? true : false);
        A.store.Float64(
          ptr + 8,
          x["motionDetectedLikelihood"] === undefined ? 0 : (x["motionDetectedLikelihood"] as number)
        );
        A.store.Enum(
          ptr + 16,
          ["UNSPECIFIED", "NO_CHANGE", "TURNED_ON", "TURNED_OFF", "DIMMER", "BRIGHTER", "BLACK_FRAME"].indexOf(
            x["lightCondition"] as string
          )
        );
        A.store.Bool(ptr + 34, "lightConditionLikelihood" in x ? true : false);
        A.store.Float64(
          ptr + 24,
          x["lightConditionLikelihood"] === undefined ? 0 : (x["lightConditionLikelihood"] as number)
        );
      }
    },
    "load_VideoHumanPresenceDetection": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 32)) {
        x["humanPresenceLikelihood"] = A.load.Float64(ptr + 0);
      } else {
        delete x["humanPresenceLikelihood"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["motionDetectedLikelihood"] = A.load.Float64(ptr + 8);
      } else {
        delete x["motionDetectedLikelihood"];
      }
      x["lightCondition"] = A.load.Enum(ptr + 16, [
        "UNSPECIFIED",
        "NO_CHANGE",
        "TURNED_ON",
        "TURNED_OFF",
        "DIMMER",
        "BRIGHTER",
        "BLACK_FRAME",
      ]);
      if (A.load.Bool(ptr + 34)) {
        x["lightConditionLikelihood"] = A.load.Float64(ptr + 24);
      } else {
        delete x["lightConditionLikelihood"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_FramePerceptionType": (ref: heap.Ref<string>): number => {
      const idx = ["UNKNOWN_TYPE", "FACE_DETECTION", "PERSON_DETECTION", "MOTION_DETECTION"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_FramePerception": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 76, false);
        A.store.Bool(ptr + 72, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 73, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 74, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 75, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);

        A.store.Bool(ptr + 32 + 35, false);
        A.store.Bool(ptr + 32 + 32, false);
        A.store.Float64(ptr + 32 + 0, 0);
        A.store.Bool(ptr + 32 + 33, false);
        A.store.Float64(ptr + 32 + 8, 0);
        A.store.Enum(ptr + 32 + 16, -1);
        A.store.Bool(ptr + 32 + 34, false);
        A.store.Float64(ptr + 32 + 24, 0);
        A.store.Ref(ptr + 68, undefined);
      } else {
        A.store.Bool(ptr + 76, true);
        A.store.Bool(ptr + 72, "frameId" in x ? true : false);
        A.store.Int32(ptr + 0, x["frameId"] === undefined ? 0 : (x["frameId"] as number));
        A.store.Bool(ptr + 73, "frameWidthInPx" in x ? true : false);
        A.store.Int32(ptr + 4, x["frameWidthInPx"] === undefined ? 0 : (x["frameWidthInPx"] as number));
        A.store.Bool(ptr + 74, "frameHeightInPx" in x ? true : false);
        A.store.Int32(ptr + 8, x["frameHeightInPx"] === undefined ? 0 : (x["frameHeightInPx"] as number));
        A.store.Bool(ptr + 75, "timestamp" in x ? true : false);
        A.store.Float64(ptr + 16, x["timestamp"] === undefined ? 0 : (x["timestamp"] as number));
        A.store.Ref(ptr + 24, x["entities"]);
        A.store.Ref(ptr + 28, x["packetLatency"]);

        if (typeof x["videoHumanPresenceDetection"] === "undefined") {
          A.store.Bool(ptr + 32 + 35, false);
          A.store.Bool(ptr + 32 + 32, false);
          A.store.Float64(ptr + 32 + 0, 0);
          A.store.Bool(ptr + 32 + 33, false);
          A.store.Float64(ptr + 32 + 8, 0);
          A.store.Enum(ptr + 32 + 16, -1);
          A.store.Bool(ptr + 32 + 34, false);
          A.store.Float64(ptr + 32 + 24, 0);
        } else {
          A.store.Bool(ptr + 32 + 35, true);
          A.store.Bool(ptr + 32 + 32, "humanPresenceLikelihood" in x["videoHumanPresenceDetection"] ? true : false);
          A.store.Float64(
            ptr + 32 + 0,
            x["videoHumanPresenceDetection"]["humanPresenceLikelihood"] === undefined
              ? 0
              : (x["videoHumanPresenceDetection"]["humanPresenceLikelihood"] as number)
          );
          A.store.Bool(ptr + 32 + 33, "motionDetectedLikelihood" in x["videoHumanPresenceDetection"] ? true : false);
          A.store.Float64(
            ptr + 32 + 8,
            x["videoHumanPresenceDetection"]["motionDetectedLikelihood"] === undefined
              ? 0
              : (x["videoHumanPresenceDetection"]["motionDetectedLikelihood"] as number)
          );
          A.store.Enum(
            ptr + 32 + 16,
            ["UNSPECIFIED", "NO_CHANGE", "TURNED_ON", "TURNED_OFF", "DIMMER", "BRIGHTER", "BLACK_FRAME"].indexOf(
              x["videoHumanPresenceDetection"]["lightCondition"] as string
            )
          );
          A.store.Bool(ptr + 32 + 34, "lightConditionLikelihood" in x["videoHumanPresenceDetection"] ? true : false);
          A.store.Float64(
            ptr + 32 + 24,
            x["videoHumanPresenceDetection"]["lightConditionLikelihood"] === undefined
              ? 0
              : (x["videoHumanPresenceDetection"]["lightConditionLikelihood"] as number)
          );
        }
        A.store.Ref(ptr + 68, x["framePerceptionTypes"]);
      }
    },
    "load_FramePerception": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 72)) {
        x["frameId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["frameId"];
      }
      if (A.load.Bool(ptr + 73)) {
        x["frameWidthInPx"] = A.load.Int32(ptr + 4);
      } else {
        delete x["frameWidthInPx"];
      }
      if (A.load.Bool(ptr + 74)) {
        x["frameHeightInPx"] = A.load.Int32(ptr + 8);
      } else {
        delete x["frameHeightInPx"];
      }
      if (A.load.Bool(ptr + 75)) {
        x["timestamp"] = A.load.Float64(ptr + 16);
      } else {
        delete x["timestamp"];
      }
      x["entities"] = A.load.Ref(ptr + 24, undefined);
      x["packetLatency"] = A.load.Ref(ptr + 28, undefined);
      if (A.load.Bool(ptr + 32 + 35)) {
        x["videoHumanPresenceDetection"] = {};
        if (A.load.Bool(ptr + 32 + 32)) {
          x["videoHumanPresenceDetection"]["humanPresenceLikelihood"] = A.load.Float64(ptr + 32 + 0);
        } else {
          delete x["videoHumanPresenceDetection"]["humanPresenceLikelihood"];
        }
        if (A.load.Bool(ptr + 32 + 33)) {
          x["videoHumanPresenceDetection"]["motionDetectedLikelihood"] = A.load.Float64(ptr + 32 + 8);
        } else {
          delete x["videoHumanPresenceDetection"]["motionDetectedLikelihood"];
        }
        x["videoHumanPresenceDetection"]["lightCondition"] = A.load.Enum(ptr + 32 + 16, [
          "UNSPECIFIED",
          "NO_CHANGE",
          "TURNED_ON",
          "TURNED_OFF",
          "DIMMER",
          "BRIGHTER",
          "BLACK_FRAME",
        ]);
        if (A.load.Bool(ptr + 32 + 34)) {
          x["videoHumanPresenceDetection"]["lightConditionLikelihood"] = A.load.Float64(ptr + 32 + 24);
        } else {
          delete x["videoHumanPresenceDetection"]["lightConditionLikelihood"];
        }
      } else {
        delete x["videoHumanPresenceDetection"];
      }
      x["framePerceptionTypes"] = A.load.Ref(ptr + 68, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ImageFormat": (ref: heap.Ref<string>): number => {
      const idx = ["RAW", "PNG", "JPEG"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ImageFrame": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 23, false);
        A.store.Bool(ptr + 20, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 21, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Enum(ptr + 8, -1);
        A.store.Bool(ptr + 22, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 23, true);
        A.store.Bool(ptr + 20, "width" in x ? true : false);
        A.store.Int32(ptr + 0, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Bool(ptr + 21, "height" in x ? true : false);
        A.store.Int32(ptr + 4, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Enum(ptr + 8, ["RAW", "PNG", "JPEG"].indexOf(x["format"] as string));
        A.store.Bool(ptr + 22, "dataLength" in x ? true : false);
        A.store.Int32(ptr + 12, x["dataLength"] === undefined ? 0 : (x["dataLength"] as number));
        A.store.Ref(ptr + 16, x["frame"]);
      }
    },
    "load_ImageFrame": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 20)) {
        x["width"] = A.load.Int32(ptr + 0);
      } else {
        delete x["width"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["height"] = A.load.Int32(ptr + 4);
      } else {
        delete x["height"];
      }
      x["format"] = A.load.Enum(ptr + 8, ["RAW", "PNG", "JPEG"]);
      if (A.load.Bool(ptr + 22)) {
        x["dataLength"] = A.load.Int32(ptr + 12);
      } else {
        delete x["dataLength"];
      }
      x["frame"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Metadata": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["visualExperienceControllerVersion"]);
      }
    },
    "load_Metadata": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["visualExperienceControllerVersion"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_PerceptionSample": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 185, false);

        A.store.Bool(ptr + 0 + 76, false);
        A.store.Bool(ptr + 0 + 72, false);
        A.store.Int32(ptr + 0 + 0, 0);
        A.store.Bool(ptr + 0 + 73, false);
        A.store.Int32(ptr + 0 + 4, 0);
        A.store.Bool(ptr + 0 + 74, false);
        A.store.Int32(ptr + 0 + 8, 0);
        A.store.Bool(ptr + 0 + 75, false);
        A.store.Float64(ptr + 0 + 16, 0);
        A.store.Ref(ptr + 0 + 24, undefined);
        A.store.Ref(ptr + 0 + 28, undefined);

        A.store.Bool(ptr + 0 + 32 + 35, false);
        A.store.Bool(ptr + 0 + 32 + 32, false);
        A.store.Float64(ptr + 0 + 32 + 0, 0);
        A.store.Bool(ptr + 0 + 32 + 33, false);
        A.store.Float64(ptr + 0 + 32 + 8, 0);
        A.store.Enum(ptr + 0 + 32 + 16, -1);
        A.store.Bool(ptr + 0 + 32 + 34, false);
        A.store.Float64(ptr + 0 + 32 + 24, 0);
        A.store.Ref(ptr + 0 + 68, undefined);

        A.store.Bool(ptr + 80 + 23, false);
        A.store.Bool(ptr + 80 + 20, false);
        A.store.Int32(ptr + 80 + 0, 0);
        A.store.Bool(ptr + 80 + 21, false);
        A.store.Int32(ptr + 80 + 4, 0);
        A.store.Enum(ptr + 80 + 8, -1);
        A.store.Bool(ptr + 80 + 22, false);
        A.store.Int32(ptr + 80 + 12, 0);
        A.store.Ref(ptr + 80 + 16, undefined);

        A.store.Bool(ptr + 104 + 54, false);
        A.store.Bool(ptr + 104 + 53, false);
        A.store.Float64(ptr + 104 + 0, 0);

        A.store.Bool(ptr + 104 + 8 + 13, false);
        A.store.Bool(ptr + 104 + 8 + 12, false);
        A.store.Float64(ptr + 104 + 8 + 0, 0);
        A.store.Ref(ptr + 104 + 8 + 8, undefined);

        A.store.Bool(ptr + 104 + 24 + 22, false);
        A.store.Bool(ptr + 104 + 24 + 21, false);
        A.store.Float64(ptr + 104 + 24 + 0, 0);

        A.store.Bool(ptr + 104 + 24 + 8 + 4, false);
        A.store.Ref(ptr + 104 + 24 + 8 + 0, undefined);

        A.store.Bool(ptr + 104 + 24 + 16 + 4, false);
        A.store.Ref(ptr + 104 + 24 + 16 + 0, undefined);

        A.store.Bool(ptr + 104 + 48 + 4, false);
        A.store.Ref(ptr + 104 + 48 + 0, undefined);

        A.store.Bool(ptr + 160 + 19, false);
        A.store.Bool(ptr + 160 + 18, false);
        A.store.Float64(ptr + 160 + 0, 0);

        A.store.Bool(ptr + 160 + 8 + 9, false);
        A.store.Bool(ptr + 160 + 8 + 8, false);
        A.store.Float64(ptr + 160 + 8 + 0, 0);

        A.store.Bool(ptr + 180 + 4, false);
        A.store.Ref(ptr + 180 + 0, undefined);
      } else {
        A.store.Bool(ptr + 185, true);

        if (typeof x["framePerception"] === "undefined") {
          A.store.Bool(ptr + 0 + 76, false);
          A.store.Bool(ptr + 0 + 72, false);
          A.store.Int32(ptr + 0 + 0, 0);
          A.store.Bool(ptr + 0 + 73, false);
          A.store.Int32(ptr + 0 + 4, 0);
          A.store.Bool(ptr + 0 + 74, false);
          A.store.Int32(ptr + 0 + 8, 0);
          A.store.Bool(ptr + 0 + 75, false);
          A.store.Float64(ptr + 0 + 16, 0);
          A.store.Ref(ptr + 0 + 24, undefined);
          A.store.Ref(ptr + 0 + 28, undefined);

          A.store.Bool(ptr + 0 + 32 + 35, false);
          A.store.Bool(ptr + 0 + 32 + 32, false);
          A.store.Float64(ptr + 0 + 32 + 0, 0);
          A.store.Bool(ptr + 0 + 32 + 33, false);
          A.store.Float64(ptr + 0 + 32 + 8, 0);
          A.store.Enum(ptr + 0 + 32 + 16, -1);
          A.store.Bool(ptr + 0 + 32 + 34, false);
          A.store.Float64(ptr + 0 + 32 + 24, 0);
          A.store.Ref(ptr + 0 + 68, undefined);
        } else {
          A.store.Bool(ptr + 0 + 76, true);
          A.store.Bool(ptr + 0 + 72, "frameId" in x["framePerception"] ? true : false);
          A.store.Int32(
            ptr + 0 + 0,
            x["framePerception"]["frameId"] === undefined ? 0 : (x["framePerception"]["frameId"] as number)
          );
          A.store.Bool(ptr + 0 + 73, "frameWidthInPx" in x["framePerception"] ? true : false);
          A.store.Int32(
            ptr + 0 + 4,
            x["framePerception"]["frameWidthInPx"] === undefined
              ? 0
              : (x["framePerception"]["frameWidthInPx"] as number)
          );
          A.store.Bool(ptr + 0 + 74, "frameHeightInPx" in x["framePerception"] ? true : false);
          A.store.Int32(
            ptr + 0 + 8,
            x["framePerception"]["frameHeightInPx"] === undefined
              ? 0
              : (x["framePerception"]["frameHeightInPx"] as number)
          );
          A.store.Bool(ptr + 0 + 75, "timestamp" in x["framePerception"] ? true : false);
          A.store.Float64(
            ptr + 0 + 16,
            x["framePerception"]["timestamp"] === undefined ? 0 : (x["framePerception"]["timestamp"] as number)
          );
          A.store.Ref(ptr + 0 + 24, x["framePerception"]["entities"]);
          A.store.Ref(ptr + 0 + 28, x["framePerception"]["packetLatency"]);

          if (typeof x["framePerception"]["videoHumanPresenceDetection"] === "undefined") {
            A.store.Bool(ptr + 0 + 32 + 35, false);
            A.store.Bool(ptr + 0 + 32 + 32, false);
            A.store.Float64(ptr + 0 + 32 + 0, 0);
            A.store.Bool(ptr + 0 + 32 + 33, false);
            A.store.Float64(ptr + 0 + 32 + 8, 0);
            A.store.Enum(ptr + 0 + 32 + 16, -1);
            A.store.Bool(ptr + 0 + 32 + 34, false);
            A.store.Float64(ptr + 0 + 32 + 24, 0);
          } else {
            A.store.Bool(ptr + 0 + 32 + 35, true);
            A.store.Bool(
              ptr + 0 + 32 + 32,
              "humanPresenceLikelihood" in x["framePerception"]["videoHumanPresenceDetection"] ? true : false
            );
            A.store.Float64(
              ptr + 0 + 32 + 0,
              x["framePerception"]["videoHumanPresenceDetection"]["humanPresenceLikelihood"] === undefined
                ? 0
                : (x["framePerception"]["videoHumanPresenceDetection"]["humanPresenceLikelihood"] as number)
            );
            A.store.Bool(
              ptr + 0 + 32 + 33,
              "motionDetectedLikelihood" in x["framePerception"]["videoHumanPresenceDetection"] ? true : false
            );
            A.store.Float64(
              ptr + 0 + 32 + 8,
              x["framePerception"]["videoHumanPresenceDetection"]["motionDetectedLikelihood"] === undefined
                ? 0
                : (x["framePerception"]["videoHumanPresenceDetection"]["motionDetectedLikelihood"] as number)
            );
            A.store.Enum(
              ptr + 0 + 32 + 16,
              ["UNSPECIFIED", "NO_CHANGE", "TURNED_ON", "TURNED_OFF", "DIMMER", "BRIGHTER", "BLACK_FRAME"].indexOf(
                x["framePerception"]["videoHumanPresenceDetection"]["lightCondition"] as string
              )
            );
            A.store.Bool(
              ptr + 0 + 32 + 34,
              "lightConditionLikelihood" in x["framePerception"]["videoHumanPresenceDetection"] ? true : false
            );
            A.store.Float64(
              ptr + 0 + 32 + 24,
              x["framePerception"]["videoHumanPresenceDetection"]["lightConditionLikelihood"] === undefined
                ? 0
                : (x["framePerception"]["videoHumanPresenceDetection"]["lightConditionLikelihood"] as number)
            );
          }
          A.store.Ref(ptr + 0 + 68, x["framePerception"]["framePerceptionTypes"]);
        }

        if (typeof x["imageFrame"] === "undefined") {
          A.store.Bool(ptr + 80 + 23, false);
          A.store.Bool(ptr + 80 + 20, false);
          A.store.Int32(ptr + 80 + 0, 0);
          A.store.Bool(ptr + 80 + 21, false);
          A.store.Int32(ptr + 80 + 4, 0);
          A.store.Enum(ptr + 80 + 8, -1);
          A.store.Bool(ptr + 80 + 22, false);
          A.store.Int32(ptr + 80 + 12, 0);
          A.store.Ref(ptr + 80 + 16, undefined);
        } else {
          A.store.Bool(ptr + 80 + 23, true);
          A.store.Bool(ptr + 80 + 20, "width" in x["imageFrame"] ? true : false);
          A.store.Int32(
            ptr + 80 + 0,
            x["imageFrame"]["width"] === undefined ? 0 : (x["imageFrame"]["width"] as number)
          );
          A.store.Bool(ptr + 80 + 21, "height" in x["imageFrame"] ? true : false);
          A.store.Int32(
            ptr + 80 + 4,
            x["imageFrame"]["height"] === undefined ? 0 : (x["imageFrame"]["height"] as number)
          );
          A.store.Enum(ptr + 80 + 8, ["RAW", "PNG", "JPEG"].indexOf(x["imageFrame"]["format"] as string));
          A.store.Bool(ptr + 80 + 22, "dataLength" in x["imageFrame"] ? true : false);
          A.store.Int32(
            ptr + 80 + 12,
            x["imageFrame"]["dataLength"] === undefined ? 0 : (x["imageFrame"]["dataLength"] as number)
          );
          A.store.Ref(ptr + 80 + 16, x["imageFrame"]["frame"]);
        }

        if (typeof x["audioPerception"] === "undefined") {
          A.store.Bool(ptr + 104 + 54, false);
          A.store.Bool(ptr + 104 + 53, false);
          A.store.Float64(ptr + 104 + 0, 0);

          A.store.Bool(ptr + 104 + 8 + 13, false);
          A.store.Bool(ptr + 104 + 8 + 12, false);
          A.store.Float64(ptr + 104 + 8 + 0, 0);
          A.store.Ref(ptr + 104 + 8 + 8, undefined);

          A.store.Bool(ptr + 104 + 24 + 22, false);
          A.store.Bool(ptr + 104 + 24 + 21, false);
          A.store.Float64(ptr + 104 + 24 + 0, 0);

          A.store.Bool(ptr + 104 + 24 + 8 + 4, false);
          A.store.Ref(ptr + 104 + 24 + 8 + 0, undefined);

          A.store.Bool(ptr + 104 + 24 + 16 + 4, false);
          A.store.Ref(ptr + 104 + 24 + 16 + 0, undefined);

          A.store.Bool(ptr + 104 + 48 + 4, false);
          A.store.Ref(ptr + 104 + 48 + 0, undefined);
        } else {
          A.store.Bool(ptr + 104 + 54, true);
          A.store.Bool(ptr + 104 + 53, "timestampUs" in x["audioPerception"] ? true : false);
          A.store.Float64(
            ptr + 104 + 0,
            x["audioPerception"]["timestampUs"] === undefined ? 0 : (x["audioPerception"]["timestampUs"] as number)
          );

          if (typeof x["audioPerception"]["audioLocalization"] === "undefined") {
            A.store.Bool(ptr + 104 + 8 + 13, false);
            A.store.Bool(ptr + 104 + 8 + 12, false);
            A.store.Float64(ptr + 104 + 8 + 0, 0);
            A.store.Ref(ptr + 104 + 8 + 8, undefined);
          } else {
            A.store.Bool(ptr + 104 + 8 + 13, true);
            A.store.Bool(
              ptr + 104 + 8 + 12,
              "azimuthRadians" in x["audioPerception"]["audioLocalization"] ? true : false
            );
            A.store.Float64(
              ptr + 104 + 8 + 0,
              x["audioPerception"]["audioLocalization"]["azimuthRadians"] === undefined
                ? 0
                : (x["audioPerception"]["audioLocalization"]["azimuthRadians"] as number)
            );
            A.store.Ref(ptr + 104 + 8 + 8, x["audioPerception"]["audioLocalization"]["azimuthScores"]);
          }

          if (typeof x["audioPerception"]["audioHumanPresenceDetection"] === "undefined") {
            A.store.Bool(ptr + 104 + 24 + 22, false);
            A.store.Bool(ptr + 104 + 24 + 21, false);
            A.store.Float64(ptr + 104 + 24 + 0, 0);

            A.store.Bool(ptr + 104 + 24 + 8 + 4, false);
            A.store.Ref(ptr + 104 + 24 + 8 + 0, undefined);

            A.store.Bool(ptr + 104 + 24 + 16 + 4, false);
            A.store.Ref(ptr + 104 + 24 + 16 + 0, undefined);
          } else {
            A.store.Bool(ptr + 104 + 24 + 22, true);
            A.store.Bool(
              ptr + 104 + 24 + 21,
              "humanPresenceLikelihood" in x["audioPerception"]["audioHumanPresenceDetection"] ? true : false
            );
            A.store.Float64(
              ptr + 104 + 24 + 0,
              x["audioPerception"]["audioHumanPresenceDetection"]["humanPresenceLikelihood"] === undefined
                ? 0
                : (x["audioPerception"]["audioHumanPresenceDetection"]["humanPresenceLikelihood"] as number)
            );

            if (typeof x["audioPerception"]["audioHumanPresenceDetection"]["noiseSpectrogram"] === "undefined") {
              A.store.Bool(ptr + 104 + 24 + 8 + 4, false);
              A.store.Ref(ptr + 104 + 24 + 8 + 0, undefined);
            } else {
              A.store.Bool(ptr + 104 + 24 + 8 + 4, true);
              A.store.Ref(
                ptr + 104 + 24 + 8 + 0,
                x["audioPerception"]["audioHumanPresenceDetection"]["noiseSpectrogram"]["values"]
              );
            }

            if (typeof x["audioPerception"]["audioHumanPresenceDetection"]["frameSpectrogram"] === "undefined") {
              A.store.Bool(ptr + 104 + 24 + 16 + 4, false);
              A.store.Ref(ptr + 104 + 24 + 16 + 0, undefined);
            } else {
              A.store.Bool(ptr + 104 + 24 + 16 + 4, true);
              A.store.Ref(
                ptr + 104 + 24 + 16 + 0,
                x["audioPerception"]["audioHumanPresenceDetection"]["frameSpectrogram"]["values"]
              );
            }
          }

          if (typeof x["audioPerception"]["hotwordDetection"] === "undefined") {
            A.store.Bool(ptr + 104 + 48 + 4, false);
            A.store.Ref(ptr + 104 + 48 + 0, undefined);
          } else {
            A.store.Bool(ptr + 104 + 48 + 4, true);
            A.store.Ref(ptr + 104 + 48 + 0, x["audioPerception"]["hotwordDetection"]["hotwords"]);
          }
        }

        if (typeof x["audioVisualPerception"] === "undefined") {
          A.store.Bool(ptr + 160 + 19, false);
          A.store.Bool(ptr + 160 + 18, false);
          A.store.Float64(ptr + 160 + 0, 0);

          A.store.Bool(ptr + 160 + 8 + 9, false);
          A.store.Bool(ptr + 160 + 8 + 8, false);
          A.store.Float64(ptr + 160 + 8 + 0, 0);
        } else {
          A.store.Bool(ptr + 160 + 19, true);
          A.store.Bool(ptr + 160 + 18, "timestampUs" in x["audioVisualPerception"] ? true : false);
          A.store.Float64(
            ptr + 160 + 0,
            x["audioVisualPerception"]["timestampUs"] === undefined
              ? 0
              : (x["audioVisualPerception"]["timestampUs"] as number)
          );

          if (typeof x["audioVisualPerception"]["audioVisualHumanPresenceDetection"] === "undefined") {
            A.store.Bool(ptr + 160 + 8 + 9, false);
            A.store.Bool(ptr + 160 + 8 + 8, false);
            A.store.Float64(ptr + 160 + 8 + 0, 0);
          } else {
            A.store.Bool(ptr + 160 + 8 + 9, true);
            A.store.Bool(
              ptr + 160 + 8 + 8,
              "humanPresenceLikelihood" in x["audioVisualPerception"]["audioVisualHumanPresenceDetection"]
                ? true
                : false
            );
            A.store.Float64(
              ptr + 160 + 8 + 0,
              x["audioVisualPerception"]["audioVisualHumanPresenceDetection"]["humanPresenceLikelihood"] === undefined
                ? 0
                : (x["audioVisualPerception"]["audioVisualHumanPresenceDetection"]["humanPresenceLikelihood"] as number)
            );
          }
        }

        if (typeof x["metadata"] === "undefined") {
          A.store.Bool(ptr + 180 + 4, false);
          A.store.Ref(ptr + 180 + 0, undefined);
        } else {
          A.store.Bool(ptr + 180 + 4, true);
          A.store.Ref(ptr + 180 + 0, x["metadata"]["visualExperienceControllerVersion"]);
        }
      }
    },
    "load_PerceptionSample": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 76)) {
        x["framePerception"] = {};
        if (A.load.Bool(ptr + 0 + 72)) {
          x["framePerception"]["frameId"] = A.load.Int32(ptr + 0 + 0);
        } else {
          delete x["framePerception"]["frameId"];
        }
        if (A.load.Bool(ptr + 0 + 73)) {
          x["framePerception"]["frameWidthInPx"] = A.load.Int32(ptr + 0 + 4);
        } else {
          delete x["framePerception"]["frameWidthInPx"];
        }
        if (A.load.Bool(ptr + 0 + 74)) {
          x["framePerception"]["frameHeightInPx"] = A.load.Int32(ptr + 0 + 8);
        } else {
          delete x["framePerception"]["frameHeightInPx"];
        }
        if (A.load.Bool(ptr + 0 + 75)) {
          x["framePerception"]["timestamp"] = A.load.Float64(ptr + 0 + 16);
        } else {
          delete x["framePerception"]["timestamp"];
        }
        x["framePerception"]["entities"] = A.load.Ref(ptr + 0 + 24, undefined);
        x["framePerception"]["packetLatency"] = A.load.Ref(ptr + 0 + 28, undefined);
        if (A.load.Bool(ptr + 0 + 32 + 35)) {
          x["framePerception"]["videoHumanPresenceDetection"] = {};
          if (A.load.Bool(ptr + 0 + 32 + 32)) {
            x["framePerception"]["videoHumanPresenceDetection"]["humanPresenceLikelihood"] = A.load.Float64(
              ptr + 0 + 32 + 0
            );
          } else {
            delete x["framePerception"]["videoHumanPresenceDetection"]["humanPresenceLikelihood"];
          }
          if (A.load.Bool(ptr + 0 + 32 + 33)) {
            x["framePerception"]["videoHumanPresenceDetection"]["motionDetectedLikelihood"] = A.load.Float64(
              ptr + 0 + 32 + 8
            );
          } else {
            delete x["framePerception"]["videoHumanPresenceDetection"]["motionDetectedLikelihood"];
          }
          x["framePerception"]["videoHumanPresenceDetection"]["lightCondition"] = A.load.Enum(ptr + 0 + 32 + 16, [
            "UNSPECIFIED",
            "NO_CHANGE",
            "TURNED_ON",
            "TURNED_OFF",
            "DIMMER",
            "BRIGHTER",
            "BLACK_FRAME",
          ]);
          if (A.load.Bool(ptr + 0 + 32 + 34)) {
            x["framePerception"]["videoHumanPresenceDetection"]["lightConditionLikelihood"] = A.load.Float64(
              ptr + 0 + 32 + 24
            );
          } else {
            delete x["framePerception"]["videoHumanPresenceDetection"]["lightConditionLikelihood"];
          }
        } else {
          delete x["framePerception"]["videoHumanPresenceDetection"];
        }
        x["framePerception"]["framePerceptionTypes"] = A.load.Ref(ptr + 0 + 68, undefined);
      } else {
        delete x["framePerception"];
      }
      if (A.load.Bool(ptr + 80 + 23)) {
        x["imageFrame"] = {};
        if (A.load.Bool(ptr + 80 + 20)) {
          x["imageFrame"]["width"] = A.load.Int32(ptr + 80 + 0);
        } else {
          delete x["imageFrame"]["width"];
        }
        if (A.load.Bool(ptr + 80 + 21)) {
          x["imageFrame"]["height"] = A.load.Int32(ptr + 80 + 4);
        } else {
          delete x["imageFrame"]["height"];
        }
        x["imageFrame"]["format"] = A.load.Enum(ptr + 80 + 8, ["RAW", "PNG", "JPEG"]);
        if (A.load.Bool(ptr + 80 + 22)) {
          x["imageFrame"]["dataLength"] = A.load.Int32(ptr + 80 + 12);
        } else {
          delete x["imageFrame"]["dataLength"];
        }
        x["imageFrame"]["frame"] = A.load.Ref(ptr + 80 + 16, undefined);
      } else {
        delete x["imageFrame"];
      }
      if (A.load.Bool(ptr + 104 + 54)) {
        x["audioPerception"] = {};
        if (A.load.Bool(ptr + 104 + 53)) {
          x["audioPerception"]["timestampUs"] = A.load.Float64(ptr + 104 + 0);
        } else {
          delete x["audioPerception"]["timestampUs"];
        }
        if (A.load.Bool(ptr + 104 + 8 + 13)) {
          x["audioPerception"]["audioLocalization"] = {};
          if (A.load.Bool(ptr + 104 + 8 + 12)) {
            x["audioPerception"]["audioLocalization"]["azimuthRadians"] = A.load.Float64(ptr + 104 + 8 + 0);
          } else {
            delete x["audioPerception"]["audioLocalization"]["azimuthRadians"];
          }
          x["audioPerception"]["audioLocalization"]["azimuthScores"] = A.load.Ref(ptr + 104 + 8 + 8, undefined);
        } else {
          delete x["audioPerception"]["audioLocalization"];
        }
        if (A.load.Bool(ptr + 104 + 24 + 22)) {
          x["audioPerception"]["audioHumanPresenceDetection"] = {};
          if (A.load.Bool(ptr + 104 + 24 + 21)) {
            x["audioPerception"]["audioHumanPresenceDetection"]["humanPresenceLikelihood"] = A.load.Float64(
              ptr + 104 + 24 + 0
            );
          } else {
            delete x["audioPerception"]["audioHumanPresenceDetection"]["humanPresenceLikelihood"];
          }
          if (A.load.Bool(ptr + 104 + 24 + 8 + 4)) {
            x["audioPerception"]["audioHumanPresenceDetection"]["noiseSpectrogram"] = {};
            x["audioPerception"]["audioHumanPresenceDetection"]["noiseSpectrogram"]["values"] = A.load.Ref(
              ptr + 104 + 24 + 8 + 0,
              undefined
            );
          } else {
            delete x["audioPerception"]["audioHumanPresenceDetection"]["noiseSpectrogram"];
          }
          if (A.load.Bool(ptr + 104 + 24 + 16 + 4)) {
            x["audioPerception"]["audioHumanPresenceDetection"]["frameSpectrogram"] = {};
            x["audioPerception"]["audioHumanPresenceDetection"]["frameSpectrogram"]["values"] = A.load.Ref(
              ptr + 104 + 24 + 16 + 0,
              undefined
            );
          } else {
            delete x["audioPerception"]["audioHumanPresenceDetection"]["frameSpectrogram"];
          }
        } else {
          delete x["audioPerception"]["audioHumanPresenceDetection"];
        }
        if (A.load.Bool(ptr + 104 + 48 + 4)) {
          x["audioPerception"]["hotwordDetection"] = {};
          x["audioPerception"]["hotwordDetection"]["hotwords"] = A.load.Ref(ptr + 104 + 48 + 0, undefined);
        } else {
          delete x["audioPerception"]["hotwordDetection"];
        }
      } else {
        delete x["audioPerception"];
      }
      if (A.load.Bool(ptr + 160 + 19)) {
        x["audioVisualPerception"] = {};
        if (A.load.Bool(ptr + 160 + 18)) {
          x["audioVisualPerception"]["timestampUs"] = A.load.Float64(ptr + 160 + 0);
        } else {
          delete x["audioVisualPerception"]["timestampUs"];
        }
        if (A.load.Bool(ptr + 160 + 8 + 9)) {
          x["audioVisualPerception"]["audioVisualHumanPresenceDetection"] = {};
          if (A.load.Bool(ptr + 160 + 8 + 8)) {
            x["audioVisualPerception"]["audioVisualHumanPresenceDetection"]["humanPresenceLikelihood"] = A.load.Float64(
              ptr + 160 + 8 + 0
            );
          } else {
            delete x["audioVisualPerception"]["audioVisualHumanPresenceDetection"]["humanPresenceLikelihood"];
          }
        } else {
          delete x["audioVisualPerception"]["audioVisualHumanPresenceDetection"];
        }
      } else {
        delete x["audioVisualPerception"];
      }
      if (A.load.Bool(ptr + 180 + 4)) {
        x["metadata"] = {};
        x["metadata"]["visualExperienceControllerVersion"] = A.load.Ref(ptr + 180 + 0, undefined);
      } else {
        delete x["metadata"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Diagnostics": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(
          ptr + 0,
          [
            "SERVICE_UNREACHABLE",
            "SERVICE_NOT_RUNNING",
            "SERVICE_BUSY_LAUNCHING",
            "SERVICE_NOT_INSTALLED",
            "MOJO_CONNECTION_FAILURE",
          ].indexOf(x["serviceError"] as string)
        );
        A.store.Ref(ptr + 4, x["perceptionSamples"]);
      }
    },
    "load_Diagnostics": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["serviceError"] = A.load.Enum(ptr + 0, [
        "SERVICE_UNREACHABLE",
        "SERVICE_NOT_RUNNING",
        "SERVICE_BUSY_LAUNCHING",
        "SERVICE_NOT_INSTALLED",
        "MOJO_CONNECTION_FAILURE",
      ]);
      x["perceptionSamples"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_Feature": (ref: heap.Ref<string>): number => {
      const idx = [
        "AUTOZOOM",
        "HOTWORD_DETECTION",
        "OCCUPANCY_DETECTION",
        "EDGE_EMBEDDINGS",
        "SOFTWARE_CROPPING",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_MediaPerception": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 25, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);

        A.store.Bool(ptr + 20 + 4, false);
        A.store.Ref(ptr + 20 + 0, undefined);
      } else {
        A.store.Bool(ptr + 26, true);
        A.store.Bool(ptr + 25, "timestamp" in x ? true : false);
        A.store.Float64(ptr + 0, x["timestamp"] === undefined ? 0 : (x["timestamp"] as number));
        A.store.Ref(ptr + 8, x["framePerceptions"]);
        A.store.Ref(ptr + 12, x["audioPerceptions"]);
        A.store.Ref(ptr + 16, x["audioVisualPerceptions"]);

        if (typeof x["metadata"] === "undefined") {
          A.store.Bool(ptr + 20 + 4, false);
          A.store.Ref(ptr + 20 + 0, undefined);
        } else {
          A.store.Bool(ptr + 20 + 4, true);
          A.store.Ref(ptr + 20 + 0, x["metadata"]["visualExperienceControllerVersion"]);
        }
      }
    },
    "load_MediaPerception": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 25)) {
        x["timestamp"] = A.load.Float64(ptr + 0);
      } else {
        delete x["timestamp"];
      }
      x["framePerceptions"] = A.load.Ref(ptr + 8, undefined);
      x["audioPerceptions"] = A.load.Ref(ptr + 12, undefined);
      x["audioVisualPerceptions"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 20 + 4)) {
        x["metadata"] = {};
        x["metadata"]["visualExperienceControllerVersion"] = A.load.Ref(ptr + 20 + 0, undefined);
      } else {
        delete x["metadata"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_NamedTemplateArgument": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_NamedTemplateArgument": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["value"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ProcessStatus": (ref: heap.Ref<string>): number => {
      const idx = ["UNKNOWN", "STARTED", "STOPPED", "SERVICE_ERROR"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ProcessState": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["UNKNOWN", "STARTED", "STOPPED", "SERVICE_ERROR"].indexOf(x["status"] as string));
        A.store.Enum(
          ptr + 4,
          [
            "SERVICE_UNREACHABLE",
            "SERVICE_NOT_RUNNING",
            "SERVICE_BUSY_LAUNCHING",
            "SERVICE_NOT_INSTALLED",
            "MOJO_CONNECTION_FAILURE",
          ].indexOf(x["serviceError"] as string)
        );
      }
    },
    "load_ProcessState": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["status"] = A.load.Enum(ptr + 0, ["UNKNOWN", "STARTED", "STOPPED", "SERVICE_ERROR"]);
      x["serviceError"] = A.load.Enum(ptr + 4, [
        "SERVICE_UNREACHABLE",
        "SERVICE_NOT_RUNNING",
        "SERVICE_BUSY_LAUNCHING",
        "SERVICE_NOT_INSTALLED",
        "MOJO_CONNECTION_FAILURE",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_Status": (ref: heap.Ref<string>): number => {
      const idx = [
        "UNINITIALIZED",
        "STARTED",
        "RUNNING",
        "SUSPENDED",
        "RESTARTING",
        "STOPPED",
        "SERVICE_ERROR",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_VideoStreamParam": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 19, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 18, false);
        A.store.Int32(ptr + 12, 0);
      } else {
        A.store.Bool(ptr + 19, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Bool(ptr + 16, "width" in x ? true : false);
        A.store.Int32(ptr + 4, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Bool(ptr + 17, "height" in x ? true : false);
        A.store.Int32(ptr + 8, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Bool(ptr + 18, "frameRate" in x ? true : false);
        A.store.Int32(ptr + 12, x["frameRate"] === undefined ? 0 : (x["frameRate"] as number));
      }
    },
    "load_VideoStreamParam": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["width"] = A.load.Int32(ptr + 4);
      } else {
        delete x["width"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["height"] = A.load.Int32(ptr + 8);
      } else {
        delete x["height"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["frameRate"] = A.load.Int32(ptr + 12);
      } else {
        delete x["frameRate"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Whiteboard": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 105, false);

        A.store.Bool(ptr + 0 + 18, false);
        A.store.Bool(ptr + 0 + 16, false);
        A.store.Float64(ptr + 0 + 0, 0);
        A.store.Bool(ptr + 0 + 17, false);
        A.store.Float64(ptr + 0 + 8, 0);

        A.store.Bool(ptr + 24 + 18, false);
        A.store.Bool(ptr + 24 + 16, false);
        A.store.Float64(ptr + 24 + 0, 0);
        A.store.Bool(ptr + 24 + 17, false);
        A.store.Float64(ptr + 24 + 8, 0);

        A.store.Bool(ptr + 48 + 18, false);
        A.store.Bool(ptr + 48 + 16, false);
        A.store.Float64(ptr + 48 + 0, 0);
        A.store.Bool(ptr + 48 + 17, false);
        A.store.Float64(ptr + 48 + 8, 0);

        A.store.Bool(ptr + 72 + 18, false);
        A.store.Bool(ptr + 72 + 16, false);
        A.store.Float64(ptr + 72 + 0, 0);
        A.store.Bool(ptr + 72 + 17, false);
        A.store.Float64(ptr + 72 + 8, 0);
        A.store.Bool(ptr + 104, false);
        A.store.Float64(ptr + 96, 0);
      } else {
        A.store.Bool(ptr + 105, true);

        if (typeof x["topLeft"] === "undefined") {
          A.store.Bool(ptr + 0 + 18, false);
          A.store.Bool(ptr + 0 + 16, false);
          A.store.Float64(ptr + 0 + 0, 0);
          A.store.Bool(ptr + 0 + 17, false);
          A.store.Float64(ptr + 0 + 8, 0);
        } else {
          A.store.Bool(ptr + 0 + 18, true);
          A.store.Bool(ptr + 0 + 16, "x" in x["topLeft"] ? true : false);
          A.store.Float64(ptr + 0 + 0, x["topLeft"]["x"] === undefined ? 0 : (x["topLeft"]["x"] as number));
          A.store.Bool(ptr + 0 + 17, "y" in x["topLeft"] ? true : false);
          A.store.Float64(ptr + 0 + 8, x["topLeft"]["y"] === undefined ? 0 : (x["topLeft"]["y"] as number));
        }

        if (typeof x["topRight"] === "undefined") {
          A.store.Bool(ptr + 24 + 18, false);
          A.store.Bool(ptr + 24 + 16, false);
          A.store.Float64(ptr + 24 + 0, 0);
          A.store.Bool(ptr + 24 + 17, false);
          A.store.Float64(ptr + 24 + 8, 0);
        } else {
          A.store.Bool(ptr + 24 + 18, true);
          A.store.Bool(ptr + 24 + 16, "x" in x["topRight"] ? true : false);
          A.store.Float64(ptr + 24 + 0, x["topRight"]["x"] === undefined ? 0 : (x["topRight"]["x"] as number));
          A.store.Bool(ptr + 24 + 17, "y" in x["topRight"] ? true : false);
          A.store.Float64(ptr + 24 + 8, x["topRight"]["y"] === undefined ? 0 : (x["topRight"]["y"] as number));
        }

        if (typeof x["bottomLeft"] === "undefined") {
          A.store.Bool(ptr + 48 + 18, false);
          A.store.Bool(ptr + 48 + 16, false);
          A.store.Float64(ptr + 48 + 0, 0);
          A.store.Bool(ptr + 48 + 17, false);
          A.store.Float64(ptr + 48 + 8, 0);
        } else {
          A.store.Bool(ptr + 48 + 18, true);
          A.store.Bool(ptr + 48 + 16, "x" in x["bottomLeft"] ? true : false);
          A.store.Float64(ptr + 48 + 0, x["bottomLeft"]["x"] === undefined ? 0 : (x["bottomLeft"]["x"] as number));
          A.store.Bool(ptr + 48 + 17, "y" in x["bottomLeft"] ? true : false);
          A.store.Float64(ptr + 48 + 8, x["bottomLeft"]["y"] === undefined ? 0 : (x["bottomLeft"]["y"] as number));
        }

        if (typeof x["bottomRight"] === "undefined") {
          A.store.Bool(ptr + 72 + 18, false);
          A.store.Bool(ptr + 72 + 16, false);
          A.store.Float64(ptr + 72 + 0, 0);
          A.store.Bool(ptr + 72 + 17, false);
          A.store.Float64(ptr + 72 + 8, 0);
        } else {
          A.store.Bool(ptr + 72 + 18, true);
          A.store.Bool(ptr + 72 + 16, "x" in x["bottomRight"] ? true : false);
          A.store.Float64(ptr + 72 + 0, x["bottomRight"]["x"] === undefined ? 0 : (x["bottomRight"]["x"] as number));
          A.store.Bool(ptr + 72 + 17, "y" in x["bottomRight"] ? true : false);
          A.store.Float64(ptr + 72 + 8, x["bottomRight"]["y"] === undefined ? 0 : (x["bottomRight"]["y"] as number));
        }
        A.store.Bool(ptr + 104, "aspectRatio" in x ? true : false);
        A.store.Float64(ptr + 96, x["aspectRatio"] === undefined ? 0 : (x["aspectRatio"] as number));
      }
    },
    "load_Whiteboard": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 18)) {
        x["topLeft"] = {};
        if (A.load.Bool(ptr + 0 + 16)) {
          x["topLeft"]["x"] = A.load.Float64(ptr + 0 + 0);
        } else {
          delete x["topLeft"]["x"];
        }
        if (A.load.Bool(ptr + 0 + 17)) {
          x["topLeft"]["y"] = A.load.Float64(ptr + 0 + 8);
        } else {
          delete x["topLeft"]["y"];
        }
      } else {
        delete x["topLeft"];
      }
      if (A.load.Bool(ptr + 24 + 18)) {
        x["topRight"] = {};
        if (A.load.Bool(ptr + 24 + 16)) {
          x["topRight"]["x"] = A.load.Float64(ptr + 24 + 0);
        } else {
          delete x["topRight"]["x"];
        }
        if (A.load.Bool(ptr + 24 + 17)) {
          x["topRight"]["y"] = A.load.Float64(ptr + 24 + 8);
        } else {
          delete x["topRight"]["y"];
        }
      } else {
        delete x["topRight"];
      }
      if (A.load.Bool(ptr + 48 + 18)) {
        x["bottomLeft"] = {};
        if (A.load.Bool(ptr + 48 + 16)) {
          x["bottomLeft"]["x"] = A.load.Float64(ptr + 48 + 0);
        } else {
          delete x["bottomLeft"]["x"];
        }
        if (A.load.Bool(ptr + 48 + 17)) {
          x["bottomLeft"]["y"] = A.load.Float64(ptr + 48 + 8);
        } else {
          delete x["bottomLeft"]["y"];
        }
      } else {
        delete x["bottomLeft"];
      }
      if (A.load.Bool(ptr + 72 + 18)) {
        x["bottomRight"] = {};
        if (A.load.Bool(ptr + 72 + 16)) {
          x["bottomRight"]["x"] = A.load.Float64(ptr + 72 + 0);
        } else {
          delete x["bottomRight"]["x"];
        }
        if (A.load.Bool(ptr + 72 + 17)) {
          x["bottomRight"]["y"] = A.load.Float64(ptr + 72 + 8);
        } else {
          delete x["bottomRight"]["y"];
        }
      } else {
        delete x["bottomRight"];
      }
      if (A.load.Bool(ptr + 104)) {
        x["aspectRatio"] = A.load.Float64(ptr + 96);
      } else {
        delete x["aspectRatio"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_State": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 140, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);

        A.store.Bool(ptr + 24 + 105, false);

        A.store.Bool(ptr + 24 + 0 + 18, false);
        A.store.Bool(ptr + 24 + 0 + 16, false);
        A.store.Float64(ptr + 24 + 0 + 0, 0);
        A.store.Bool(ptr + 24 + 0 + 17, false);
        A.store.Float64(ptr + 24 + 0 + 8, 0);

        A.store.Bool(ptr + 24 + 24 + 18, false);
        A.store.Bool(ptr + 24 + 24 + 16, false);
        A.store.Float64(ptr + 24 + 24 + 0, 0);
        A.store.Bool(ptr + 24 + 24 + 17, false);
        A.store.Float64(ptr + 24 + 24 + 8, 0);

        A.store.Bool(ptr + 24 + 48 + 18, false);
        A.store.Bool(ptr + 24 + 48 + 16, false);
        A.store.Float64(ptr + 24 + 48 + 0, 0);
        A.store.Bool(ptr + 24 + 48 + 17, false);
        A.store.Float64(ptr + 24 + 48 + 8, 0);

        A.store.Bool(ptr + 24 + 72 + 18, false);
        A.store.Bool(ptr + 24 + 72 + 16, false);
        A.store.Float64(ptr + 24 + 72 + 0, 0);
        A.store.Bool(ptr + 24 + 72 + 17, false);
        A.store.Float64(ptr + 24 + 72 + 8, 0);
        A.store.Bool(ptr + 24 + 104, false);
        A.store.Float64(ptr + 24 + 96, 0);
        A.store.Ref(ptr + 132, undefined);
        A.store.Ref(ptr + 136, undefined);
      } else {
        A.store.Bool(ptr + 140, true);
        A.store.Enum(
          ptr + 0,
          ["UNINITIALIZED", "STARTED", "RUNNING", "SUSPENDED", "RESTARTING", "STOPPED", "SERVICE_ERROR"].indexOf(
            x["status"] as string
          )
        );
        A.store.Ref(ptr + 4, x["deviceContext"]);
        A.store.Enum(
          ptr + 8,
          [
            "SERVICE_UNREACHABLE",
            "SERVICE_NOT_RUNNING",
            "SERVICE_BUSY_LAUNCHING",
            "SERVICE_NOT_INSTALLED",
            "MOJO_CONNECTION_FAILURE",
          ].indexOf(x["serviceError"] as string)
        );
        A.store.Ref(ptr + 12, x["videoStreamParam"]);
        A.store.Ref(ptr + 16, x["configuration"]);

        if (typeof x["whiteboard"] === "undefined") {
          A.store.Bool(ptr + 24 + 105, false);

          A.store.Bool(ptr + 24 + 0 + 18, false);
          A.store.Bool(ptr + 24 + 0 + 16, false);
          A.store.Float64(ptr + 24 + 0 + 0, 0);
          A.store.Bool(ptr + 24 + 0 + 17, false);
          A.store.Float64(ptr + 24 + 0 + 8, 0);

          A.store.Bool(ptr + 24 + 24 + 18, false);
          A.store.Bool(ptr + 24 + 24 + 16, false);
          A.store.Float64(ptr + 24 + 24 + 0, 0);
          A.store.Bool(ptr + 24 + 24 + 17, false);
          A.store.Float64(ptr + 24 + 24 + 8, 0);

          A.store.Bool(ptr + 24 + 48 + 18, false);
          A.store.Bool(ptr + 24 + 48 + 16, false);
          A.store.Float64(ptr + 24 + 48 + 0, 0);
          A.store.Bool(ptr + 24 + 48 + 17, false);
          A.store.Float64(ptr + 24 + 48 + 8, 0);

          A.store.Bool(ptr + 24 + 72 + 18, false);
          A.store.Bool(ptr + 24 + 72 + 16, false);
          A.store.Float64(ptr + 24 + 72 + 0, 0);
          A.store.Bool(ptr + 24 + 72 + 17, false);
          A.store.Float64(ptr + 24 + 72 + 8, 0);
          A.store.Bool(ptr + 24 + 104, false);
          A.store.Float64(ptr + 24 + 96, 0);
        } else {
          A.store.Bool(ptr + 24 + 105, true);

          if (typeof x["whiteboard"]["topLeft"] === "undefined") {
            A.store.Bool(ptr + 24 + 0 + 18, false);
            A.store.Bool(ptr + 24 + 0 + 16, false);
            A.store.Float64(ptr + 24 + 0 + 0, 0);
            A.store.Bool(ptr + 24 + 0 + 17, false);
            A.store.Float64(ptr + 24 + 0 + 8, 0);
          } else {
            A.store.Bool(ptr + 24 + 0 + 18, true);
            A.store.Bool(ptr + 24 + 0 + 16, "x" in x["whiteboard"]["topLeft"] ? true : false);
            A.store.Float64(
              ptr + 24 + 0 + 0,
              x["whiteboard"]["topLeft"]["x"] === undefined ? 0 : (x["whiteboard"]["topLeft"]["x"] as number)
            );
            A.store.Bool(ptr + 24 + 0 + 17, "y" in x["whiteboard"]["topLeft"] ? true : false);
            A.store.Float64(
              ptr + 24 + 0 + 8,
              x["whiteboard"]["topLeft"]["y"] === undefined ? 0 : (x["whiteboard"]["topLeft"]["y"] as number)
            );
          }

          if (typeof x["whiteboard"]["topRight"] === "undefined") {
            A.store.Bool(ptr + 24 + 24 + 18, false);
            A.store.Bool(ptr + 24 + 24 + 16, false);
            A.store.Float64(ptr + 24 + 24 + 0, 0);
            A.store.Bool(ptr + 24 + 24 + 17, false);
            A.store.Float64(ptr + 24 + 24 + 8, 0);
          } else {
            A.store.Bool(ptr + 24 + 24 + 18, true);
            A.store.Bool(ptr + 24 + 24 + 16, "x" in x["whiteboard"]["topRight"] ? true : false);
            A.store.Float64(
              ptr + 24 + 24 + 0,
              x["whiteboard"]["topRight"]["x"] === undefined ? 0 : (x["whiteboard"]["topRight"]["x"] as number)
            );
            A.store.Bool(ptr + 24 + 24 + 17, "y" in x["whiteboard"]["topRight"] ? true : false);
            A.store.Float64(
              ptr + 24 + 24 + 8,
              x["whiteboard"]["topRight"]["y"] === undefined ? 0 : (x["whiteboard"]["topRight"]["y"] as number)
            );
          }

          if (typeof x["whiteboard"]["bottomLeft"] === "undefined") {
            A.store.Bool(ptr + 24 + 48 + 18, false);
            A.store.Bool(ptr + 24 + 48 + 16, false);
            A.store.Float64(ptr + 24 + 48 + 0, 0);
            A.store.Bool(ptr + 24 + 48 + 17, false);
            A.store.Float64(ptr + 24 + 48 + 8, 0);
          } else {
            A.store.Bool(ptr + 24 + 48 + 18, true);
            A.store.Bool(ptr + 24 + 48 + 16, "x" in x["whiteboard"]["bottomLeft"] ? true : false);
            A.store.Float64(
              ptr + 24 + 48 + 0,
              x["whiteboard"]["bottomLeft"]["x"] === undefined ? 0 : (x["whiteboard"]["bottomLeft"]["x"] as number)
            );
            A.store.Bool(ptr + 24 + 48 + 17, "y" in x["whiteboard"]["bottomLeft"] ? true : false);
            A.store.Float64(
              ptr + 24 + 48 + 8,
              x["whiteboard"]["bottomLeft"]["y"] === undefined ? 0 : (x["whiteboard"]["bottomLeft"]["y"] as number)
            );
          }

          if (typeof x["whiteboard"]["bottomRight"] === "undefined") {
            A.store.Bool(ptr + 24 + 72 + 18, false);
            A.store.Bool(ptr + 24 + 72 + 16, false);
            A.store.Float64(ptr + 24 + 72 + 0, 0);
            A.store.Bool(ptr + 24 + 72 + 17, false);
            A.store.Float64(ptr + 24 + 72 + 8, 0);
          } else {
            A.store.Bool(ptr + 24 + 72 + 18, true);
            A.store.Bool(ptr + 24 + 72 + 16, "x" in x["whiteboard"]["bottomRight"] ? true : false);
            A.store.Float64(
              ptr + 24 + 72 + 0,
              x["whiteboard"]["bottomRight"]["x"] === undefined ? 0 : (x["whiteboard"]["bottomRight"]["x"] as number)
            );
            A.store.Bool(ptr + 24 + 72 + 17, "y" in x["whiteboard"]["bottomRight"] ? true : false);
            A.store.Float64(
              ptr + 24 + 72 + 8,
              x["whiteboard"]["bottomRight"]["y"] === undefined ? 0 : (x["whiteboard"]["bottomRight"]["y"] as number)
            );
          }
          A.store.Bool(ptr + 24 + 104, "aspectRatio" in x["whiteboard"] ? true : false);
          A.store.Float64(
            ptr + 24 + 96,
            x["whiteboard"]["aspectRatio"] === undefined ? 0 : (x["whiteboard"]["aspectRatio"] as number)
          );
        }
        A.store.Ref(ptr + 132, x["features"]);
        A.store.Ref(ptr + 136, x["namedTemplateArguments"]);
      }
    },
    "load_State": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["status"] = A.load.Enum(ptr + 0, [
        "UNINITIALIZED",
        "STARTED",
        "RUNNING",
        "SUSPENDED",
        "RESTARTING",
        "STOPPED",
        "SERVICE_ERROR",
      ]);
      x["deviceContext"] = A.load.Ref(ptr + 4, undefined);
      x["serviceError"] = A.load.Enum(ptr + 8, [
        "SERVICE_UNREACHABLE",
        "SERVICE_NOT_RUNNING",
        "SERVICE_BUSY_LAUNCHING",
        "SERVICE_NOT_INSTALLED",
        "MOJO_CONNECTION_FAILURE",
      ]);
      x["videoStreamParam"] = A.load.Ref(ptr + 12, undefined);
      x["configuration"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 24 + 105)) {
        x["whiteboard"] = {};
        if (A.load.Bool(ptr + 24 + 0 + 18)) {
          x["whiteboard"]["topLeft"] = {};
          if (A.load.Bool(ptr + 24 + 0 + 16)) {
            x["whiteboard"]["topLeft"]["x"] = A.load.Float64(ptr + 24 + 0 + 0);
          } else {
            delete x["whiteboard"]["topLeft"]["x"];
          }
          if (A.load.Bool(ptr + 24 + 0 + 17)) {
            x["whiteboard"]["topLeft"]["y"] = A.load.Float64(ptr + 24 + 0 + 8);
          } else {
            delete x["whiteboard"]["topLeft"]["y"];
          }
        } else {
          delete x["whiteboard"]["topLeft"];
        }
        if (A.load.Bool(ptr + 24 + 24 + 18)) {
          x["whiteboard"]["topRight"] = {};
          if (A.load.Bool(ptr + 24 + 24 + 16)) {
            x["whiteboard"]["topRight"]["x"] = A.load.Float64(ptr + 24 + 24 + 0);
          } else {
            delete x["whiteboard"]["topRight"]["x"];
          }
          if (A.load.Bool(ptr + 24 + 24 + 17)) {
            x["whiteboard"]["topRight"]["y"] = A.load.Float64(ptr + 24 + 24 + 8);
          } else {
            delete x["whiteboard"]["topRight"]["y"];
          }
        } else {
          delete x["whiteboard"]["topRight"];
        }
        if (A.load.Bool(ptr + 24 + 48 + 18)) {
          x["whiteboard"]["bottomLeft"] = {};
          if (A.load.Bool(ptr + 24 + 48 + 16)) {
            x["whiteboard"]["bottomLeft"]["x"] = A.load.Float64(ptr + 24 + 48 + 0);
          } else {
            delete x["whiteboard"]["bottomLeft"]["x"];
          }
          if (A.load.Bool(ptr + 24 + 48 + 17)) {
            x["whiteboard"]["bottomLeft"]["y"] = A.load.Float64(ptr + 24 + 48 + 8);
          } else {
            delete x["whiteboard"]["bottomLeft"]["y"];
          }
        } else {
          delete x["whiteboard"]["bottomLeft"];
        }
        if (A.load.Bool(ptr + 24 + 72 + 18)) {
          x["whiteboard"]["bottomRight"] = {};
          if (A.load.Bool(ptr + 24 + 72 + 16)) {
            x["whiteboard"]["bottomRight"]["x"] = A.load.Float64(ptr + 24 + 72 + 0);
          } else {
            delete x["whiteboard"]["bottomRight"]["x"];
          }
          if (A.load.Bool(ptr + 24 + 72 + 17)) {
            x["whiteboard"]["bottomRight"]["y"] = A.load.Float64(ptr + 24 + 72 + 8);
          } else {
            delete x["whiteboard"]["bottomRight"]["y"];
          }
        } else {
          delete x["whiteboard"]["bottomRight"];
        }
        if (A.load.Bool(ptr + 24 + 104)) {
          x["whiteboard"]["aspectRatio"] = A.load.Float64(ptr + 24 + 96);
        } else {
          delete x["whiteboard"]["aspectRatio"];
        }
      } else {
        delete x["whiteboard"];
      }
      x["features"] = A.load.Ref(ptr + 132, undefined);
      x["namedTemplateArguments"] = A.load.Ref(ptr + 136, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetDiagnostics": (): heap.Ref<boolean> => {
      if (WEBEXT?.mediaPerceptionPrivate && "getDiagnostics" in WEBEXT?.mediaPerceptionPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDiagnostics": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPerceptionPrivate.getDiagnostics);
    },
    "call_GetDiagnostics": (retPtr: Pointer): void => {
      const _ret = WEBEXT.mediaPerceptionPrivate.getDiagnostics();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDiagnostics": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mediaPerceptionPrivate.getDiagnostics();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetState": (): heap.Ref<boolean> => {
      if (WEBEXT?.mediaPerceptionPrivate && "getState" in WEBEXT?.mediaPerceptionPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPerceptionPrivate.getState);
    },
    "call_GetState": (retPtr: Pointer): void => {
      const _ret = WEBEXT.mediaPerceptionPrivate.getState();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetState": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mediaPerceptionPrivate.getState();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMediaPerception": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.mediaPerceptionPrivate?.onMediaPerception &&
        "addListener" in WEBEXT?.mediaPerceptionPrivate?.onMediaPerception
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMediaPerception": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPerceptionPrivate.onMediaPerception.addListener);
    },
    "call_OnMediaPerception": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mediaPerceptionPrivate.onMediaPerception.addListener(A.H.get<object>(callback));
    },
    "try_OnMediaPerception": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mediaPerceptionPrivate.onMediaPerception.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMediaPerception": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.mediaPerceptionPrivate?.onMediaPerception &&
        "removeListener" in WEBEXT?.mediaPerceptionPrivate?.onMediaPerception
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMediaPerception": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPerceptionPrivate.onMediaPerception.removeListener);
    },
    "call_OffMediaPerception": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mediaPerceptionPrivate.onMediaPerception.removeListener(A.H.get<object>(callback));
    },
    "try_OffMediaPerception": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mediaPerceptionPrivate.onMediaPerception.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMediaPerception": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.mediaPerceptionPrivate?.onMediaPerception &&
        "hasListener" in WEBEXT?.mediaPerceptionPrivate?.onMediaPerception
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMediaPerception": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPerceptionPrivate.onMediaPerception.hasListener);
    },
    "call_HasOnMediaPerception": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.mediaPerceptionPrivate.onMediaPerception.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMediaPerception": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.mediaPerceptionPrivate.onMediaPerception.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetAnalyticsComponent": (): heap.Ref<boolean> => {
      if (WEBEXT?.mediaPerceptionPrivate && "setAnalyticsComponent" in WEBEXT?.mediaPerceptionPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetAnalyticsComponent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPerceptionPrivate.setAnalyticsComponent);
    },
    "call_SetAnalyticsComponent": (retPtr: Pointer, component: Pointer): void => {
      const component_ffi = {};

      component_ffi["type"] = A.load.Enum(component + 0, ["LIGHT", "FULL"]);

      const _ret = WEBEXT.mediaPerceptionPrivate.setAnalyticsComponent(component_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetAnalyticsComponent": (retPtr: Pointer, errPtr: Pointer, component: Pointer): heap.Ref<boolean> => {
      try {
        const component_ffi = {};

        component_ffi["type"] = A.load.Enum(component + 0, ["LIGHT", "FULL"]);

        const _ret = WEBEXT.mediaPerceptionPrivate.setAnalyticsComponent(component_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetComponentProcessState": (): heap.Ref<boolean> => {
      if (WEBEXT?.mediaPerceptionPrivate && "setComponentProcessState" in WEBEXT?.mediaPerceptionPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetComponentProcessState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPerceptionPrivate.setComponentProcessState);
    },
    "call_SetComponentProcessState": (retPtr: Pointer, processState: Pointer): void => {
      const processState_ffi = {};

      processState_ffi["status"] = A.load.Enum(processState + 0, ["UNKNOWN", "STARTED", "STOPPED", "SERVICE_ERROR"]);
      processState_ffi["serviceError"] = A.load.Enum(processState + 4, [
        "SERVICE_UNREACHABLE",
        "SERVICE_NOT_RUNNING",
        "SERVICE_BUSY_LAUNCHING",
        "SERVICE_NOT_INSTALLED",
        "MOJO_CONNECTION_FAILURE",
      ]);

      const _ret = WEBEXT.mediaPerceptionPrivate.setComponentProcessState(processState_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetComponentProcessState": (retPtr: Pointer, errPtr: Pointer, processState: Pointer): heap.Ref<boolean> => {
      try {
        const processState_ffi = {};

        processState_ffi["status"] = A.load.Enum(processState + 0, ["UNKNOWN", "STARTED", "STOPPED", "SERVICE_ERROR"]);
        processState_ffi["serviceError"] = A.load.Enum(processState + 4, [
          "SERVICE_UNREACHABLE",
          "SERVICE_NOT_RUNNING",
          "SERVICE_BUSY_LAUNCHING",
          "SERVICE_NOT_INSTALLED",
          "MOJO_CONNECTION_FAILURE",
        ]);

        const _ret = WEBEXT.mediaPerceptionPrivate.setComponentProcessState(processState_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetState": (): heap.Ref<boolean> => {
      if (WEBEXT?.mediaPerceptionPrivate && "setState" in WEBEXT?.mediaPerceptionPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.mediaPerceptionPrivate.setState);
    },
    "call_SetState": (retPtr: Pointer, state: Pointer): void => {
      const state_ffi = {};

      state_ffi["status"] = A.load.Enum(state + 0, [
        "UNINITIALIZED",
        "STARTED",
        "RUNNING",
        "SUSPENDED",
        "RESTARTING",
        "STOPPED",
        "SERVICE_ERROR",
      ]);
      state_ffi["deviceContext"] = A.load.Ref(state + 4, undefined);
      state_ffi["serviceError"] = A.load.Enum(state + 8, [
        "SERVICE_UNREACHABLE",
        "SERVICE_NOT_RUNNING",
        "SERVICE_BUSY_LAUNCHING",
        "SERVICE_NOT_INSTALLED",
        "MOJO_CONNECTION_FAILURE",
      ]);
      state_ffi["videoStreamParam"] = A.load.Ref(state + 12, undefined);
      state_ffi["configuration"] = A.load.Ref(state + 16, undefined);
      if (A.load.Bool(state + 24 + 105)) {
        state_ffi["whiteboard"] = {};
        if (A.load.Bool(state + 24 + 0 + 18)) {
          state_ffi["whiteboard"]["topLeft"] = {};
          if (A.load.Bool(state + 24 + 0 + 16)) {
            state_ffi["whiteboard"]["topLeft"]["x"] = A.load.Float64(state + 24 + 0 + 0);
          }
          if (A.load.Bool(state + 24 + 0 + 17)) {
            state_ffi["whiteboard"]["topLeft"]["y"] = A.load.Float64(state + 24 + 0 + 8);
          }
        }
        if (A.load.Bool(state + 24 + 24 + 18)) {
          state_ffi["whiteboard"]["topRight"] = {};
          if (A.load.Bool(state + 24 + 24 + 16)) {
            state_ffi["whiteboard"]["topRight"]["x"] = A.load.Float64(state + 24 + 24 + 0);
          }
          if (A.load.Bool(state + 24 + 24 + 17)) {
            state_ffi["whiteboard"]["topRight"]["y"] = A.load.Float64(state + 24 + 24 + 8);
          }
        }
        if (A.load.Bool(state + 24 + 48 + 18)) {
          state_ffi["whiteboard"]["bottomLeft"] = {};
          if (A.load.Bool(state + 24 + 48 + 16)) {
            state_ffi["whiteboard"]["bottomLeft"]["x"] = A.load.Float64(state + 24 + 48 + 0);
          }
          if (A.load.Bool(state + 24 + 48 + 17)) {
            state_ffi["whiteboard"]["bottomLeft"]["y"] = A.load.Float64(state + 24 + 48 + 8);
          }
        }
        if (A.load.Bool(state + 24 + 72 + 18)) {
          state_ffi["whiteboard"]["bottomRight"] = {};
          if (A.load.Bool(state + 24 + 72 + 16)) {
            state_ffi["whiteboard"]["bottomRight"]["x"] = A.load.Float64(state + 24 + 72 + 0);
          }
          if (A.load.Bool(state + 24 + 72 + 17)) {
            state_ffi["whiteboard"]["bottomRight"]["y"] = A.load.Float64(state + 24 + 72 + 8);
          }
        }
        if (A.load.Bool(state + 24 + 104)) {
          state_ffi["whiteboard"]["aspectRatio"] = A.load.Float64(state + 24 + 96);
        }
      }
      state_ffi["features"] = A.load.Ref(state + 132, undefined);
      state_ffi["namedTemplateArguments"] = A.load.Ref(state + 136, undefined);

      const _ret = WEBEXT.mediaPerceptionPrivate.setState(state_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetState": (retPtr: Pointer, errPtr: Pointer, state: Pointer): heap.Ref<boolean> => {
      try {
        const state_ffi = {};

        state_ffi["status"] = A.load.Enum(state + 0, [
          "UNINITIALIZED",
          "STARTED",
          "RUNNING",
          "SUSPENDED",
          "RESTARTING",
          "STOPPED",
          "SERVICE_ERROR",
        ]);
        state_ffi["deviceContext"] = A.load.Ref(state + 4, undefined);
        state_ffi["serviceError"] = A.load.Enum(state + 8, [
          "SERVICE_UNREACHABLE",
          "SERVICE_NOT_RUNNING",
          "SERVICE_BUSY_LAUNCHING",
          "SERVICE_NOT_INSTALLED",
          "MOJO_CONNECTION_FAILURE",
        ]);
        state_ffi["videoStreamParam"] = A.load.Ref(state + 12, undefined);
        state_ffi["configuration"] = A.load.Ref(state + 16, undefined);
        if (A.load.Bool(state + 24 + 105)) {
          state_ffi["whiteboard"] = {};
          if (A.load.Bool(state + 24 + 0 + 18)) {
            state_ffi["whiteboard"]["topLeft"] = {};
            if (A.load.Bool(state + 24 + 0 + 16)) {
              state_ffi["whiteboard"]["topLeft"]["x"] = A.load.Float64(state + 24 + 0 + 0);
            }
            if (A.load.Bool(state + 24 + 0 + 17)) {
              state_ffi["whiteboard"]["topLeft"]["y"] = A.load.Float64(state + 24 + 0 + 8);
            }
          }
          if (A.load.Bool(state + 24 + 24 + 18)) {
            state_ffi["whiteboard"]["topRight"] = {};
            if (A.load.Bool(state + 24 + 24 + 16)) {
              state_ffi["whiteboard"]["topRight"]["x"] = A.load.Float64(state + 24 + 24 + 0);
            }
            if (A.load.Bool(state + 24 + 24 + 17)) {
              state_ffi["whiteboard"]["topRight"]["y"] = A.load.Float64(state + 24 + 24 + 8);
            }
          }
          if (A.load.Bool(state + 24 + 48 + 18)) {
            state_ffi["whiteboard"]["bottomLeft"] = {};
            if (A.load.Bool(state + 24 + 48 + 16)) {
              state_ffi["whiteboard"]["bottomLeft"]["x"] = A.load.Float64(state + 24 + 48 + 0);
            }
            if (A.load.Bool(state + 24 + 48 + 17)) {
              state_ffi["whiteboard"]["bottomLeft"]["y"] = A.load.Float64(state + 24 + 48 + 8);
            }
          }
          if (A.load.Bool(state + 24 + 72 + 18)) {
            state_ffi["whiteboard"]["bottomRight"] = {};
            if (A.load.Bool(state + 24 + 72 + 16)) {
              state_ffi["whiteboard"]["bottomRight"]["x"] = A.load.Float64(state + 24 + 72 + 0);
            }
            if (A.load.Bool(state + 24 + 72 + 17)) {
              state_ffi["whiteboard"]["bottomRight"]["y"] = A.load.Float64(state + 24 + 72 + 8);
            }
          }
          if (A.load.Bool(state + 24 + 104)) {
            state_ffi["whiteboard"]["aspectRatio"] = A.load.Float64(state + 24 + 96);
          }
        }
        state_ffi["features"] = A.load.Ref(state + 132, undefined);
        state_ffi["namedTemplateArguments"] = A.load.Ref(state + 136, undefined);

        const _ret = WEBEXT.mediaPerceptionPrivate.setState(state_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
