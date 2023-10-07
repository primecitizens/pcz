import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/webcamprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_AutofocusState": (ref: heap.Ref<string>): number => {
      const idx = ["on", "off"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_PanDirection": (ref: heap.Ref<string>): number => {
      const idx = ["stop", "right", "left"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_Protocol": (ref: heap.Ref<string>): number => {
      const idx = ["visca"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ProtocolConfiguration": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Enum(ptr + 0, -1);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Enum(ptr + 0, ["visca"].indexOf(x["protocol"] as string));
      }
    },
    "load_ProtocolConfiguration": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["protocol"] = A.load.Enum(ptr + 0, ["visca"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Range": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Bool(ptr + 16, "min" in x ? true : false);
        A.store.Float64(ptr + 0, x["min"] === undefined ? 0 : (x["min"] as number));
        A.store.Bool(ptr + 17, "max" in x ? true : false);
        A.store.Float64(ptr + 8, x["max"] === undefined ? 0 : (x["max"] as number));
      }
    },
    "load_Range": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["min"] = A.load.Float64(ptr + 0);
      } else {
        delete x["min"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["max"] = A.load.Float64(ptr + 8);
      } else {
        delete x["max"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_TiltDirection": (ref: heap.Ref<string>): number => {
      const idx = ["stop", "up", "down"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_WebcamConfiguration": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 78, false);
        A.store.Bool(ptr + 72, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 73, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Bool(ptr + 74, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Bool(ptr + 75, false);
        A.store.Float64(ptr + 32, 0);
        A.store.Enum(ptr + 40, -1);
        A.store.Bool(ptr + 76, false);
        A.store.Float64(ptr + 48, 0);
        A.store.Enum(ptr + 56, -1);
        A.store.Bool(ptr + 77, false);
        A.store.Float64(ptr + 64, 0);
      } else {
        A.store.Bool(ptr + 78, true);
        A.store.Bool(ptr + 72, "pan" in x ? true : false);
        A.store.Float64(ptr + 0, x["pan"] === undefined ? 0 : (x["pan"] as number));
        A.store.Bool(ptr + 73, "panSpeed" in x ? true : false);
        A.store.Float64(ptr + 8, x["panSpeed"] === undefined ? 0 : (x["panSpeed"] as number));
        A.store.Enum(ptr + 16, ["stop", "right", "left"].indexOf(x["panDirection"] as string));
        A.store.Bool(ptr + 74, "tilt" in x ? true : false);
        A.store.Float64(ptr + 24, x["tilt"] === undefined ? 0 : (x["tilt"] as number));
        A.store.Bool(ptr + 75, "tiltSpeed" in x ? true : false);
        A.store.Float64(ptr + 32, x["tiltSpeed"] === undefined ? 0 : (x["tiltSpeed"] as number));
        A.store.Enum(ptr + 40, ["stop", "up", "down"].indexOf(x["tiltDirection"] as string));
        A.store.Bool(ptr + 76, "zoom" in x ? true : false);
        A.store.Float64(ptr + 48, x["zoom"] === undefined ? 0 : (x["zoom"] as number));
        A.store.Enum(ptr + 56, ["on", "off"].indexOf(x["autofocusState"] as string));
        A.store.Bool(ptr + 77, "focus" in x ? true : false);
        A.store.Float64(ptr + 64, x["focus"] === undefined ? 0 : (x["focus"] as number));
      }
    },
    "load_WebcamConfiguration": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 72)) {
        x["pan"] = A.load.Float64(ptr + 0);
      } else {
        delete x["pan"];
      }
      if (A.load.Bool(ptr + 73)) {
        x["panSpeed"] = A.load.Float64(ptr + 8);
      } else {
        delete x["panSpeed"];
      }
      x["panDirection"] = A.load.Enum(ptr + 16, ["stop", "right", "left"]);
      if (A.load.Bool(ptr + 74)) {
        x["tilt"] = A.load.Float64(ptr + 24);
      } else {
        delete x["tilt"];
      }
      if (A.load.Bool(ptr + 75)) {
        x["tiltSpeed"] = A.load.Float64(ptr + 32);
      } else {
        delete x["tiltSpeed"];
      }
      x["tiltDirection"] = A.load.Enum(ptr + 40, ["stop", "up", "down"]);
      if (A.load.Bool(ptr + 76)) {
        x["zoom"] = A.load.Float64(ptr + 48);
      } else {
        delete x["zoom"];
      }
      x["autofocusState"] = A.load.Enum(ptr + 56, ["on", "off"]);
      if (A.load.Bool(ptr + 77)) {
        x["focus"] = A.load.Float64(ptr + 64);
      } else {
        delete x["focus"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_WebcamCurrentConfiguration": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 127, false);
        A.store.Bool(ptr + 123, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 124, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 125, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 126, false);
        A.store.Float64(ptr + 24, 0);

        A.store.Bool(ptr + 32 + 18, false);
        A.store.Bool(ptr + 32 + 16, false);
        A.store.Float64(ptr + 32 + 0, 0);
        A.store.Bool(ptr + 32 + 17, false);
        A.store.Float64(ptr + 32 + 8, 0);

        A.store.Bool(ptr + 56 + 18, false);
        A.store.Bool(ptr + 56 + 16, false);
        A.store.Float64(ptr + 56 + 0, 0);
        A.store.Bool(ptr + 56 + 17, false);
        A.store.Float64(ptr + 56 + 8, 0);

        A.store.Bool(ptr + 80 + 18, false);
        A.store.Bool(ptr + 80 + 16, false);
        A.store.Float64(ptr + 80 + 0, 0);
        A.store.Bool(ptr + 80 + 17, false);
        A.store.Float64(ptr + 80 + 8, 0);

        A.store.Bool(ptr + 104 + 18, false);
        A.store.Bool(ptr + 104 + 16, false);
        A.store.Float64(ptr + 104 + 0, 0);
        A.store.Bool(ptr + 104 + 17, false);
        A.store.Float64(ptr + 104 + 8, 0);
      } else {
        A.store.Bool(ptr + 127, true);
        A.store.Bool(ptr + 123, "pan" in x ? true : false);
        A.store.Float64(ptr + 0, x["pan"] === undefined ? 0 : (x["pan"] as number));
        A.store.Bool(ptr + 124, "tilt" in x ? true : false);
        A.store.Float64(ptr + 8, x["tilt"] === undefined ? 0 : (x["tilt"] as number));
        A.store.Bool(ptr + 125, "zoom" in x ? true : false);
        A.store.Float64(ptr + 16, x["zoom"] === undefined ? 0 : (x["zoom"] as number));
        A.store.Bool(ptr + 126, "focus" in x ? true : false);
        A.store.Float64(ptr + 24, x["focus"] === undefined ? 0 : (x["focus"] as number));

        if (typeof x["panRange"] === "undefined") {
          A.store.Bool(ptr + 32 + 18, false);
          A.store.Bool(ptr + 32 + 16, false);
          A.store.Float64(ptr + 32 + 0, 0);
          A.store.Bool(ptr + 32 + 17, false);
          A.store.Float64(ptr + 32 + 8, 0);
        } else {
          A.store.Bool(ptr + 32 + 18, true);
          A.store.Bool(ptr + 32 + 16, "min" in x["panRange"] ? true : false);
          A.store.Float64(ptr + 32 + 0, x["panRange"]["min"] === undefined ? 0 : (x["panRange"]["min"] as number));
          A.store.Bool(ptr + 32 + 17, "max" in x["panRange"] ? true : false);
          A.store.Float64(ptr + 32 + 8, x["panRange"]["max"] === undefined ? 0 : (x["panRange"]["max"] as number));
        }

        if (typeof x["tiltRange"] === "undefined") {
          A.store.Bool(ptr + 56 + 18, false);
          A.store.Bool(ptr + 56 + 16, false);
          A.store.Float64(ptr + 56 + 0, 0);
          A.store.Bool(ptr + 56 + 17, false);
          A.store.Float64(ptr + 56 + 8, 0);
        } else {
          A.store.Bool(ptr + 56 + 18, true);
          A.store.Bool(ptr + 56 + 16, "min" in x["tiltRange"] ? true : false);
          A.store.Float64(ptr + 56 + 0, x["tiltRange"]["min"] === undefined ? 0 : (x["tiltRange"]["min"] as number));
          A.store.Bool(ptr + 56 + 17, "max" in x["tiltRange"] ? true : false);
          A.store.Float64(ptr + 56 + 8, x["tiltRange"]["max"] === undefined ? 0 : (x["tiltRange"]["max"] as number));
        }

        if (typeof x["zoomRange"] === "undefined") {
          A.store.Bool(ptr + 80 + 18, false);
          A.store.Bool(ptr + 80 + 16, false);
          A.store.Float64(ptr + 80 + 0, 0);
          A.store.Bool(ptr + 80 + 17, false);
          A.store.Float64(ptr + 80 + 8, 0);
        } else {
          A.store.Bool(ptr + 80 + 18, true);
          A.store.Bool(ptr + 80 + 16, "min" in x["zoomRange"] ? true : false);
          A.store.Float64(ptr + 80 + 0, x["zoomRange"]["min"] === undefined ? 0 : (x["zoomRange"]["min"] as number));
          A.store.Bool(ptr + 80 + 17, "max" in x["zoomRange"] ? true : false);
          A.store.Float64(ptr + 80 + 8, x["zoomRange"]["max"] === undefined ? 0 : (x["zoomRange"]["max"] as number));
        }

        if (typeof x["focusRange"] === "undefined") {
          A.store.Bool(ptr + 104 + 18, false);
          A.store.Bool(ptr + 104 + 16, false);
          A.store.Float64(ptr + 104 + 0, 0);
          A.store.Bool(ptr + 104 + 17, false);
          A.store.Float64(ptr + 104 + 8, 0);
        } else {
          A.store.Bool(ptr + 104 + 18, true);
          A.store.Bool(ptr + 104 + 16, "min" in x["focusRange"] ? true : false);
          A.store.Float64(ptr + 104 + 0, x["focusRange"]["min"] === undefined ? 0 : (x["focusRange"]["min"] as number));
          A.store.Bool(ptr + 104 + 17, "max" in x["focusRange"] ? true : false);
          A.store.Float64(ptr + 104 + 8, x["focusRange"]["max"] === undefined ? 0 : (x["focusRange"]["max"] as number));
        }
      }
    },
    "load_WebcamCurrentConfiguration": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 123)) {
        x["pan"] = A.load.Float64(ptr + 0);
      } else {
        delete x["pan"];
      }
      if (A.load.Bool(ptr + 124)) {
        x["tilt"] = A.load.Float64(ptr + 8);
      } else {
        delete x["tilt"];
      }
      if (A.load.Bool(ptr + 125)) {
        x["zoom"] = A.load.Float64(ptr + 16);
      } else {
        delete x["zoom"];
      }
      if (A.load.Bool(ptr + 126)) {
        x["focus"] = A.load.Float64(ptr + 24);
      } else {
        delete x["focus"];
      }
      if (A.load.Bool(ptr + 32 + 18)) {
        x["panRange"] = {};
        if (A.load.Bool(ptr + 32 + 16)) {
          x["panRange"]["min"] = A.load.Float64(ptr + 32 + 0);
        } else {
          delete x["panRange"]["min"];
        }
        if (A.load.Bool(ptr + 32 + 17)) {
          x["panRange"]["max"] = A.load.Float64(ptr + 32 + 8);
        } else {
          delete x["panRange"]["max"];
        }
      } else {
        delete x["panRange"];
      }
      if (A.load.Bool(ptr + 56 + 18)) {
        x["tiltRange"] = {};
        if (A.load.Bool(ptr + 56 + 16)) {
          x["tiltRange"]["min"] = A.load.Float64(ptr + 56 + 0);
        } else {
          delete x["tiltRange"]["min"];
        }
        if (A.load.Bool(ptr + 56 + 17)) {
          x["tiltRange"]["max"] = A.load.Float64(ptr + 56 + 8);
        } else {
          delete x["tiltRange"]["max"];
        }
      } else {
        delete x["tiltRange"];
      }
      if (A.load.Bool(ptr + 80 + 18)) {
        x["zoomRange"] = {};
        if (A.load.Bool(ptr + 80 + 16)) {
          x["zoomRange"]["min"] = A.load.Float64(ptr + 80 + 0);
        } else {
          delete x["zoomRange"]["min"];
        }
        if (A.load.Bool(ptr + 80 + 17)) {
          x["zoomRange"]["max"] = A.load.Float64(ptr + 80 + 8);
        } else {
          delete x["zoomRange"]["max"];
        }
      } else {
        delete x["zoomRange"];
      }
      if (A.load.Bool(ptr + 104 + 18)) {
        x["focusRange"] = {};
        if (A.load.Bool(ptr + 104 + 16)) {
          x["focusRange"]["min"] = A.load.Float64(ptr + 104 + 0);
        } else {
          delete x["focusRange"]["min"];
        }
        if (A.load.Bool(ptr + 104 + 17)) {
          x["focusRange"]["max"] = A.load.Float64(ptr + 104 + 8);
        } else {
          delete x["focusRange"]["max"];
        }
      } else {
        delete x["focusRange"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_CloseWebcam": (): heap.Ref<boolean> => {
      if (WEBEXT?.webcamPrivate && "closeWebcam" in WEBEXT?.webcamPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CloseWebcam": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webcamPrivate.closeWebcam);
    },
    "call_CloseWebcam": (retPtr: Pointer, webcamId: heap.Ref<object>): void => {
      const _ret = WEBEXT.webcamPrivate.closeWebcam(A.H.get<object>(webcamId));
    },
    "try_CloseWebcam": (retPtr: Pointer, errPtr: Pointer, webcamId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webcamPrivate.closeWebcam(A.H.get<object>(webcamId));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Get": (): heap.Ref<boolean> => {
      if (WEBEXT?.webcamPrivate && "get" in WEBEXT?.webcamPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Get": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webcamPrivate.get);
    },
    "call_Get": (retPtr: Pointer, webcamId: heap.Ref<object>): void => {
      const _ret = WEBEXT.webcamPrivate.get(A.H.get<object>(webcamId));
      A.store.Ref(retPtr, _ret);
    },
    "try_Get": (retPtr: Pointer, errPtr: Pointer, webcamId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webcamPrivate.get(A.H.get<object>(webcamId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenSerialWebcam": (): heap.Ref<boolean> => {
      if (WEBEXT?.webcamPrivate && "openSerialWebcam" in WEBEXT?.webcamPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenSerialWebcam": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webcamPrivate.openSerialWebcam);
    },
    "call_OpenSerialWebcam": (retPtr: Pointer, path: heap.Ref<object>, protocol: Pointer): void => {
      const protocol_ffi = {};

      protocol_ffi["protocol"] = A.load.Enum(protocol + 0, ["visca"]);

      const _ret = WEBEXT.webcamPrivate.openSerialWebcam(A.H.get<object>(path), protocol_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_OpenSerialWebcam": (
      retPtr: Pointer,
      errPtr: Pointer,
      path: heap.Ref<object>,
      protocol: Pointer
    ): heap.Ref<boolean> => {
      try {
        const protocol_ffi = {};

        protocol_ffi["protocol"] = A.load.Enum(protocol + 0, ["visca"]);

        const _ret = WEBEXT.webcamPrivate.openSerialWebcam(A.H.get<object>(path), protocol_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Reset": (): heap.Ref<boolean> => {
      if (WEBEXT?.webcamPrivate && "reset" in WEBEXT?.webcamPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Reset": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webcamPrivate.reset);
    },
    "call_Reset": (retPtr: Pointer, webcamId: heap.Ref<object>, config: Pointer): void => {
      const config_ffi = {};

      if (A.load.Bool(config + 72)) {
        config_ffi["pan"] = A.load.Float64(config + 0);
      }
      if (A.load.Bool(config + 73)) {
        config_ffi["panSpeed"] = A.load.Float64(config + 8);
      }
      config_ffi["panDirection"] = A.load.Enum(config + 16, ["stop", "right", "left"]);
      if (A.load.Bool(config + 74)) {
        config_ffi["tilt"] = A.load.Float64(config + 24);
      }
      if (A.load.Bool(config + 75)) {
        config_ffi["tiltSpeed"] = A.load.Float64(config + 32);
      }
      config_ffi["tiltDirection"] = A.load.Enum(config + 40, ["stop", "up", "down"]);
      if (A.load.Bool(config + 76)) {
        config_ffi["zoom"] = A.load.Float64(config + 48);
      }
      config_ffi["autofocusState"] = A.load.Enum(config + 56, ["on", "off"]);
      if (A.load.Bool(config + 77)) {
        config_ffi["focus"] = A.load.Float64(config + 64);
      }

      const _ret = WEBEXT.webcamPrivate.reset(A.H.get<object>(webcamId), config_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Reset": (retPtr: Pointer, errPtr: Pointer, webcamId: heap.Ref<object>, config: Pointer): heap.Ref<boolean> => {
      try {
        const config_ffi = {};

        if (A.load.Bool(config + 72)) {
          config_ffi["pan"] = A.load.Float64(config + 0);
        }
        if (A.load.Bool(config + 73)) {
          config_ffi["panSpeed"] = A.load.Float64(config + 8);
        }
        config_ffi["panDirection"] = A.load.Enum(config + 16, ["stop", "right", "left"]);
        if (A.load.Bool(config + 74)) {
          config_ffi["tilt"] = A.load.Float64(config + 24);
        }
        if (A.load.Bool(config + 75)) {
          config_ffi["tiltSpeed"] = A.load.Float64(config + 32);
        }
        config_ffi["tiltDirection"] = A.load.Enum(config + 40, ["stop", "up", "down"]);
        if (A.load.Bool(config + 76)) {
          config_ffi["zoom"] = A.load.Float64(config + 48);
        }
        config_ffi["autofocusState"] = A.load.Enum(config + 56, ["on", "off"]);
        if (A.load.Bool(config + 77)) {
          config_ffi["focus"] = A.load.Float64(config + 64);
        }

        const _ret = WEBEXT.webcamPrivate.reset(A.H.get<object>(webcamId), config_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RestoreCameraPreset": (): heap.Ref<boolean> => {
      if (WEBEXT?.webcamPrivate && "restoreCameraPreset" in WEBEXT?.webcamPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RestoreCameraPreset": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webcamPrivate.restoreCameraPreset);
    },
    "call_RestoreCameraPreset": (retPtr: Pointer, webcamId: heap.Ref<object>, presetNumber: number): void => {
      const _ret = WEBEXT.webcamPrivate.restoreCameraPreset(A.H.get<object>(webcamId), presetNumber);
      A.store.Ref(retPtr, _ret);
    },
    "try_RestoreCameraPreset": (
      retPtr: Pointer,
      errPtr: Pointer,
      webcamId: heap.Ref<object>,
      presetNumber: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webcamPrivate.restoreCameraPreset(A.H.get<object>(webcamId), presetNumber);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Set": (): heap.Ref<boolean> => {
      if (WEBEXT?.webcamPrivate && "set" in WEBEXT?.webcamPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Set": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webcamPrivate.set);
    },
    "call_Set": (retPtr: Pointer, webcamId: heap.Ref<object>, config: Pointer): void => {
      const config_ffi = {};

      if (A.load.Bool(config + 72)) {
        config_ffi["pan"] = A.load.Float64(config + 0);
      }
      if (A.load.Bool(config + 73)) {
        config_ffi["panSpeed"] = A.load.Float64(config + 8);
      }
      config_ffi["panDirection"] = A.load.Enum(config + 16, ["stop", "right", "left"]);
      if (A.load.Bool(config + 74)) {
        config_ffi["tilt"] = A.load.Float64(config + 24);
      }
      if (A.load.Bool(config + 75)) {
        config_ffi["tiltSpeed"] = A.load.Float64(config + 32);
      }
      config_ffi["tiltDirection"] = A.load.Enum(config + 40, ["stop", "up", "down"]);
      if (A.load.Bool(config + 76)) {
        config_ffi["zoom"] = A.load.Float64(config + 48);
      }
      config_ffi["autofocusState"] = A.load.Enum(config + 56, ["on", "off"]);
      if (A.load.Bool(config + 77)) {
        config_ffi["focus"] = A.load.Float64(config + 64);
      }

      const _ret = WEBEXT.webcamPrivate.set(A.H.get<object>(webcamId), config_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Set": (retPtr: Pointer, errPtr: Pointer, webcamId: heap.Ref<object>, config: Pointer): heap.Ref<boolean> => {
      try {
        const config_ffi = {};

        if (A.load.Bool(config + 72)) {
          config_ffi["pan"] = A.load.Float64(config + 0);
        }
        if (A.load.Bool(config + 73)) {
          config_ffi["panSpeed"] = A.load.Float64(config + 8);
        }
        config_ffi["panDirection"] = A.load.Enum(config + 16, ["stop", "right", "left"]);
        if (A.load.Bool(config + 74)) {
          config_ffi["tilt"] = A.load.Float64(config + 24);
        }
        if (A.load.Bool(config + 75)) {
          config_ffi["tiltSpeed"] = A.load.Float64(config + 32);
        }
        config_ffi["tiltDirection"] = A.load.Enum(config + 40, ["stop", "up", "down"]);
        if (A.load.Bool(config + 76)) {
          config_ffi["zoom"] = A.load.Float64(config + 48);
        }
        config_ffi["autofocusState"] = A.load.Enum(config + 56, ["on", "off"]);
        if (A.load.Bool(config + 77)) {
          config_ffi["focus"] = A.load.Float64(config + 64);
        }

        const _ret = WEBEXT.webcamPrivate.set(A.H.get<object>(webcamId), config_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetCameraPreset": (): heap.Ref<boolean> => {
      if (WEBEXT?.webcamPrivate && "setCameraPreset" in WEBEXT?.webcamPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetCameraPreset": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webcamPrivate.setCameraPreset);
    },
    "call_SetCameraPreset": (retPtr: Pointer, webcamId: heap.Ref<object>, presetNumber: number): void => {
      const _ret = WEBEXT.webcamPrivate.setCameraPreset(A.H.get<object>(webcamId), presetNumber);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetCameraPreset": (
      retPtr: Pointer,
      errPtr: Pointer,
      webcamId: heap.Ref<object>,
      presetNumber: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webcamPrivate.setCameraPreset(A.H.get<object>(webcamId), presetNumber);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetHome": (): heap.Ref<boolean> => {
      if (WEBEXT?.webcamPrivate && "setHome" in WEBEXT?.webcamPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetHome": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.webcamPrivate.setHome);
    },
    "call_SetHome": (retPtr: Pointer, webcamId: heap.Ref<object>): void => {
      const _ret = WEBEXT.webcamPrivate.setHome(A.H.get<object>(webcamId));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetHome": (retPtr: Pointer, errPtr: Pointer, webcamId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.webcamPrivate.setHome(A.H.get<object>(webcamId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
