import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/tts", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_EventType": (ref: heap.Ref<string>): number => {
      const idx = [
        "start",
        "end",
        "word",
        "sentence",
        "marker",
        "interrupted",
        "cancelled",
        "error",
        "pause",
        "resume",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_TtsEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 40, false);
        A.store.Bool(ptr + 36, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 37, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 38, false);
        A.store.Int64(ptr + 16, 0);
        A.store.Bool(ptr + 39, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Enum(ptr + 32, -1);
      } else {
        A.store.Bool(ptr + 40, true);
        A.store.Bool(ptr + 36, "charIndex" in x ? true : false);
        A.store.Int64(ptr + 0, x["charIndex"] === undefined ? 0 : (x["charIndex"] as number));
        A.store.Ref(ptr + 8, x["errorMessage"]);
        A.store.Bool(ptr + 37, "isFinalEvent" in x ? true : false);
        A.store.Bool(ptr + 12, x["isFinalEvent"] ? true : false);
        A.store.Bool(ptr + 38, "length" in x ? true : false);
        A.store.Int64(ptr + 16, x["length"] === undefined ? 0 : (x["length"] as number));
        A.store.Bool(ptr + 39, "srcId" in x ? true : false);
        A.store.Float64(ptr + 24, x["srcId"] === undefined ? 0 : (x["srcId"] as number));
        A.store.Enum(
          ptr + 32,
          [
            "start",
            "end",
            "word",
            "sentence",
            "marker",
            "interrupted",
            "cancelled",
            "error",
            "pause",
            "resume",
          ].indexOf(x["type"] as string)
        );
      }
    },
    "load_TtsEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 36)) {
        x["charIndex"] = A.load.Int64(ptr + 0);
      } else {
        delete x["charIndex"];
      }
      x["errorMessage"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 37)) {
        x["isFinalEvent"] = A.load.Bool(ptr + 12);
      } else {
        delete x["isFinalEvent"];
      }
      if (A.load.Bool(ptr + 38)) {
        x["length"] = A.load.Int64(ptr + 16);
      } else {
        delete x["length"];
      }
      if (A.load.Bool(ptr + 39)) {
        x["srcId"] = A.load.Float64(ptr + 24);
      } else {
        delete x["srcId"];
      }
      x["type"] = A.load.Enum(ptr + 32, [
        "start",
        "end",
        "word",
        "sentence",
        "marker",
        "interrupted",
        "cancelled",
        "error",
        "pause",
        "resume",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_VoiceGender": (ref: heap.Ref<string>): number => {
      const idx = ["male", "female"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_TtsOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 60, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 56, false);
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 8, undefined);
        A.store.Enum(ptr + 12, -1);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Bool(ptr + 57, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Bool(ptr + 58, false);
        A.store.Float64(ptr + 32, 0);
        A.store.Ref(ptr + 40, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Bool(ptr + 59, false);
        A.store.Float64(ptr + 48, 0);
      } else {
        A.store.Bool(ptr + 60, true);
        A.store.Ref(ptr + 0, x["desiredEventTypes"]);
        A.store.Bool(ptr + 56, "enqueue" in x ? true : false);
        A.store.Bool(ptr + 4, x["enqueue"] ? true : false);
        A.store.Ref(ptr + 8, x["extensionId"]);
        A.store.Enum(ptr + 12, ["male", "female"].indexOf(x["gender"] as string));
        A.store.Ref(ptr + 16, x["lang"]);
        A.store.Ref(ptr + 20, x["onEvent"]);
        A.store.Bool(ptr + 57, "pitch" in x ? true : false);
        A.store.Float64(ptr + 24, x["pitch"] === undefined ? 0 : (x["pitch"] as number));
        A.store.Bool(ptr + 58, "rate" in x ? true : false);
        A.store.Float64(ptr + 32, x["rate"] === undefined ? 0 : (x["rate"] as number));
        A.store.Ref(ptr + 40, x["requiredEventTypes"]);
        A.store.Ref(ptr + 44, x["voiceName"]);
        A.store.Bool(ptr + 59, "volume" in x ? true : false);
        A.store.Float64(ptr + 48, x["volume"] === undefined ? 0 : (x["volume"] as number));
      }
    },
    "load_TtsOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["desiredEventTypes"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 56)) {
        x["enqueue"] = A.load.Bool(ptr + 4);
      } else {
        delete x["enqueue"];
      }
      x["extensionId"] = A.load.Ref(ptr + 8, undefined);
      x["gender"] = A.load.Enum(ptr + 12, ["male", "female"]);
      x["lang"] = A.load.Ref(ptr + 16, undefined);
      x["onEvent"] = A.load.Ref(ptr + 20, undefined);
      if (A.load.Bool(ptr + 57)) {
        x["pitch"] = A.load.Float64(ptr + 24);
      } else {
        delete x["pitch"];
      }
      if (A.load.Bool(ptr + 58)) {
        x["rate"] = A.load.Float64(ptr + 32);
      } else {
        delete x["rate"];
      }
      x["requiredEventTypes"] = A.load.Ref(ptr + 40, undefined);
      x["voiceName"] = A.load.Ref(ptr + 44, undefined);
      if (A.load.Bool(ptr + 59)) {
        x["volume"] = A.load.Float64(ptr + 48);
      } else {
        delete x["volume"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TtsVoice": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 20, undefined);
      } else {
        A.store.Bool(ptr + 25, true);
        A.store.Ref(ptr + 0, x["eventTypes"]);
        A.store.Ref(ptr + 4, x["extensionId"]);
        A.store.Enum(ptr + 8, ["male", "female"].indexOf(x["gender"] as string));
        A.store.Ref(ptr + 12, x["lang"]);
        A.store.Bool(ptr + 24, "remote" in x ? true : false);
        A.store.Bool(ptr + 16, x["remote"] ? true : false);
        A.store.Ref(ptr + 20, x["voiceName"]);
      }
    },
    "load_TtsVoice": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["eventTypes"] = A.load.Ref(ptr + 0, undefined);
      x["extensionId"] = A.load.Ref(ptr + 4, undefined);
      x["gender"] = A.load.Enum(ptr + 8, ["male", "female"]);
      x["lang"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 24)) {
        x["remote"] = A.load.Bool(ptr + 16);
      } else {
        delete x["remote"];
      }
      x["voiceName"] = A.load.Ref(ptr + 20, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetVoices": (): heap.Ref<boolean> => {
      if (WEBEXT?.tts && "getVoices" in WEBEXT?.tts) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetVoices": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tts.getVoices);
    },
    "call_GetVoices": (retPtr: Pointer): void => {
      const _ret = WEBEXT.tts.getVoices();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetVoices": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tts.getVoices();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsSpeaking": (): heap.Ref<boolean> => {
      if (WEBEXT?.tts && "isSpeaking" in WEBEXT?.tts) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsSpeaking": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tts.isSpeaking);
    },
    "call_IsSpeaking": (retPtr: Pointer): void => {
      const _ret = WEBEXT.tts.isSpeaking();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsSpeaking": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tts.isSpeaking();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.tts?.onEvent && "addListener" in WEBEXT?.tts?.onEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tts.onEvent.addListener);
    },
    "call_OnEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tts.onEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tts.onEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.tts?.onEvent && "removeListener" in WEBEXT?.tts?.onEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tts.onEvent.removeListener);
    },
    "call_OffEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tts.onEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tts.onEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.tts?.onEvent && "hasListener" in WEBEXT?.tts?.onEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tts.onEvent.hasListener);
    },
    "call_HasOnEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.tts.onEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tts.onEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Pause": (): heap.Ref<boolean> => {
      if (WEBEXT?.tts && "pause" in WEBEXT?.tts) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Pause": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tts.pause);
    },
    "call_Pause": (retPtr: Pointer): void => {
      const _ret = WEBEXT.tts.pause();
    },
    "try_Pause": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tts.pause();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Resume": (): heap.Ref<boolean> => {
      if (WEBEXT?.tts && "resume" in WEBEXT?.tts) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Resume": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tts.resume);
    },
    "call_Resume": (retPtr: Pointer): void => {
      const _ret = WEBEXT.tts.resume();
    },
    "try_Resume": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tts.resume();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Speak": (): heap.Ref<boolean> => {
      if (WEBEXT?.tts && "speak" in WEBEXT?.tts) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Speak": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tts.speak);
    },
    "call_Speak": (retPtr: Pointer, utterance: heap.Ref<object>, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["desiredEventTypes"] = A.load.Ref(options + 0, undefined);
      if (A.load.Bool(options + 56)) {
        options_ffi["enqueue"] = A.load.Bool(options + 4);
      }
      options_ffi["extensionId"] = A.load.Ref(options + 8, undefined);
      options_ffi["gender"] = A.load.Enum(options + 12, ["male", "female"]);
      options_ffi["lang"] = A.load.Ref(options + 16, undefined);
      options_ffi["onEvent"] = A.load.Ref(options + 20, undefined);
      if (A.load.Bool(options + 57)) {
        options_ffi["pitch"] = A.load.Float64(options + 24);
      }
      if (A.load.Bool(options + 58)) {
        options_ffi["rate"] = A.load.Float64(options + 32);
      }
      options_ffi["requiredEventTypes"] = A.load.Ref(options + 40, undefined);
      options_ffi["voiceName"] = A.load.Ref(options + 44, undefined);
      if (A.load.Bool(options + 59)) {
        options_ffi["volume"] = A.load.Float64(options + 48);
      }

      const _ret = WEBEXT.tts.speak(A.H.get<object>(utterance), options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Speak": (
      retPtr: Pointer,
      errPtr: Pointer,
      utterance: heap.Ref<object>,
      options: Pointer
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["desiredEventTypes"] = A.load.Ref(options + 0, undefined);
        if (A.load.Bool(options + 56)) {
          options_ffi["enqueue"] = A.load.Bool(options + 4);
        }
        options_ffi["extensionId"] = A.load.Ref(options + 8, undefined);
        options_ffi["gender"] = A.load.Enum(options + 12, ["male", "female"]);
        options_ffi["lang"] = A.load.Ref(options + 16, undefined);
        options_ffi["onEvent"] = A.load.Ref(options + 20, undefined);
        if (A.load.Bool(options + 57)) {
          options_ffi["pitch"] = A.load.Float64(options + 24);
        }
        if (A.load.Bool(options + 58)) {
          options_ffi["rate"] = A.load.Float64(options + 32);
        }
        options_ffi["requiredEventTypes"] = A.load.Ref(options + 40, undefined);
        options_ffi["voiceName"] = A.load.Ref(options + 44, undefined);
        if (A.load.Bool(options + 59)) {
          options_ffi["volume"] = A.load.Float64(options + 48);
        }

        const _ret = WEBEXT.tts.speak(A.H.get<object>(utterance), options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Stop": (): heap.Ref<boolean> => {
      if (WEBEXT?.tts && "stop" in WEBEXT?.tts) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Stop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.tts.stop);
    },
    "call_Stop": (retPtr: Pointer): void => {
      const _ret = WEBEXT.tts.stop();
    },
    "try_Stop": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.tts.stop();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
