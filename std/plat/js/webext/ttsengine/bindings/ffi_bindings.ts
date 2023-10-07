import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/ttsengine", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AudioBuffer": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 19, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 17, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
      } else {
        A.store.Bool(ptr + 19, true);
        A.store.Ref(ptr + 0, x["audioBuffer"]);
        A.store.Bool(ptr + 17, "charIndex" in x ? true : false);
        A.store.Int64(ptr + 8, x["charIndex"] === undefined ? 0 : (x["charIndex"] as number));
        A.store.Bool(ptr + 18, "isLastBuffer" in x ? true : false);
        A.store.Bool(ptr + 16, x["isLastBuffer"] ? true : false);
      }
    },
    "load_AudioBuffer": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["audioBuffer"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 17)) {
        x["charIndex"] = A.load.Int64(ptr + 8);
      } else {
        delete x["charIndex"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["isLastBuffer"] = A.load.Bool(ptr + 16);
      } else {
        delete x["isLastBuffer"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AudioStreamOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Int64(ptr + 0, x["bufferSize"] === undefined ? 0 : (x["bufferSize"] as number));
        A.store.Int64(ptr + 8, x["sampleRate"] === undefined ? 0 : (x["sampleRate"] as number));
      }
    },
    "load_AudioStreamOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["bufferSize"] = A.load.Int64(ptr + 0);
      x["sampleRate"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_VoiceGender": (ref: heap.Ref<string>): number => {
      const idx = ["male", "female"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SpeakOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 43, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 40, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 41, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Ref(ptr + 24, undefined);
        A.store.Bool(ptr + 42, false);
        A.store.Float64(ptr + 32, 0);
      } else {
        A.store.Bool(ptr + 43, true);
        A.store.Enum(ptr + 0, ["male", "female"].indexOf(x["gender"] as string));
        A.store.Ref(ptr + 4, x["lang"]);
        A.store.Bool(ptr + 40, "pitch" in x ? true : false);
        A.store.Float64(ptr + 8, x["pitch"] === undefined ? 0 : (x["pitch"] as number));
        A.store.Bool(ptr + 41, "rate" in x ? true : false);
        A.store.Float64(ptr + 16, x["rate"] === undefined ? 0 : (x["rate"] as number));
        A.store.Ref(ptr + 24, x["voiceName"]);
        A.store.Bool(ptr + 42, "volume" in x ? true : false);
        A.store.Float64(ptr + 32, x["volume"] === undefined ? 0 : (x["volume"] as number));
      }
    },
    "load_SpeakOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["gender"] = A.load.Enum(ptr + 0, ["male", "female"]);
      x["lang"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 40)) {
        x["pitch"] = A.load.Float64(ptr + 8);
      } else {
        delete x["pitch"];
      }
      if (A.load.Bool(ptr + 41)) {
        x["rate"] = A.load.Float64(ptr + 16);
      } else {
        delete x["rate"];
      }
      x["voiceName"] = A.load.Ref(ptr + 24, undefined);
      if (A.load.Bool(ptr + 42)) {
        x["volume"] = A.load.Float64(ptr + 32);
      } else {
        delete x["volume"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_OnPause": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine?.onPause && "addListener" in WEBEXT?.ttsEngine?.onPause) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPause": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.onPause.addListener);
    },
    "call_OnPause": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.onPause.addListener(A.H.get<object>(callback));
    },
    "try_OnPause": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.onPause.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPause": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine?.onPause && "removeListener" in WEBEXT?.ttsEngine?.onPause) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPause": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.onPause.removeListener);
    },
    "call_OffPause": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.onPause.removeListener(A.H.get<object>(callback));
    },
    "try_OffPause": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.onPause.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPause": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine?.onPause && "hasListener" in WEBEXT?.ttsEngine?.onPause) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPause": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.onPause.hasListener);
    },
    "call_HasOnPause": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.onPause.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPause": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.onPause.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnResume": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine?.onResume && "addListener" in WEBEXT?.ttsEngine?.onResume) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnResume": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.onResume.addListener);
    },
    "call_OnResume": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.onResume.addListener(A.H.get<object>(callback));
    },
    "try_OnResume": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.onResume.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffResume": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine?.onResume && "removeListener" in WEBEXT?.ttsEngine?.onResume) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffResume": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.onResume.removeListener);
    },
    "call_OffResume": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.onResume.removeListener(A.H.get<object>(callback));
    },
    "try_OffResume": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.onResume.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnResume": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine?.onResume && "hasListener" in WEBEXT?.ttsEngine?.onResume) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnResume": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.onResume.hasListener);
    },
    "call_HasOnResume": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.onResume.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnResume": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.onResume.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSpeak": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine?.onSpeak && "addListener" in WEBEXT?.ttsEngine?.onSpeak) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSpeak": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.onSpeak.addListener);
    },
    "call_OnSpeak": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.onSpeak.addListener(A.H.get<object>(callback));
    },
    "try_OnSpeak": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.onSpeak.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSpeak": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine?.onSpeak && "removeListener" in WEBEXT?.ttsEngine?.onSpeak) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSpeak": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.onSpeak.removeListener);
    },
    "call_OffSpeak": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.onSpeak.removeListener(A.H.get<object>(callback));
    },
    "try_OffSpeak": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.onSpeak.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSpeak": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine?.onSpeak && "hasListener" in WEBEXT?.ttsEngine?.onSpeak) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSpeak": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.onSpeak.hasListener);
    },
    "call_HasOnSpeak": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.onSpeak.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSpeak": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.onSpeak.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSpeakWithAudioStream": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine?.onSpeakWithAudioStream && "addListener" in WEBEXT?.ttsEngine?.onSpeakWithAudioStream) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSpeakWithAudioStream": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.onSpeakWithAudioStream.addListener);
    },
    "call_OnSpeakWithAudioStream": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.onSpeakWithAudioStream.addListener(A.H.get<object>(callback));
    },
    "try_OnSpeakWithAudioStream": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.onSpeakWithAudioStream.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSpeakWithAudioStream": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine?.onSpeakWithAudioStream && "removeListener" in WEBEXT?.ttsEngine?.onSpeakWithAudioStream) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSpeakWithAudioStream": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.onSpeakWithAudioStream.removeListener);
    },
    "call_OffSpeakWithAudioStream": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.onSpeakWithAudioStream.removeListener(A.H.get<object>(callback));
    },
    "try_OffSpeakWithAudioStream": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.onSpeakWithAudioStream.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSpeakWithAudioStream": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine?.onSpeakWithAudioStream && "hasListener" in WEBEXT?.ttsEngine?.onSpeakWithAudioStream) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSpeakWithAudioStream": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.onSpeakWithAudioStream.hasListener);
    },
    "call_HasOnSpeakWithAudioStream": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.onSpeakWithAudioStream.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSpeakWithAudioStream": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.onSpeakWithAudioStream.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnStop": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine?.onStop && "addListener" in WEBEXT?.ttsEngine?.onStop) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnStop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.onStop.addListener);
    },
    "call_OnStop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.onStop.addListener(A.H.get<object>(callback));
    },
    "try_OnStop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.onStop.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffStop": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine?.onStop && "removeListener" in WEBEXT?.ttsEngine?.onStop) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffStop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.onStop.removeListener);
    },
    "call_OffStop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.onStop.removeListener(A.H.get<object>(callback));
    },
    "try_OffStop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.onStop.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnStop": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine?.onStop && "hasListener" in WEBEXT?.ttsEngine?.onStop) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnStop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.onStop.hasListener);
    },
    "call_HasOnStop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.onStop.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnStop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.onStop.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendTtsAudio": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine && "sendTtsAudio" in WEBEXT?.ttsEngine) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendTtsAudio": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.sendTtsAudio);
    },
    "call_SendTtsAudio": (retPtr: Pointer, requestId: number, audio: Pointer): void => {
      const audio_ffi = {};

      audio_ffi["audioBuffer"] = A.load.Ref(audio + 0, undefined);
      if (A.load.Bool(audio + 17)) {
        audio_ffi["charIndex"] = A.load.Int64(audio + 8);
      }
      if (A.load.Bool(audio + 18)) {
        audio_ffi["isLastBuffer"] = A.load.Bool(audio + 16);
      }

      const _ret = WEBEXT.ttsEngine.sendTtsAudio(requestId, audio_ffi);
    },
    "try_SendTtsAudio": (retPtr: Pointer, errPtr: Pointer, requestId: number, audio: Pointer): heap.Ref<boolean> => {
      try {
        const audio_ffi = {};

        audio_ffi["audioBuffer"] = A.load.Ref(audio + 0, undefined);
        if (A.load.Bool(audio + 17)) {
          audio_ffi["charIndex"] = A.load.Int64(audio + 8);
        }
        if (A.load.Bool(audio + 18)) {
          audio_ffi["isLastBuffer"] = A.load.Bool(audio + 16);
        }

        const _ret = WEBEXT.ttsEngine.sendTtsAudio(requestId, audio_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendTtsEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine && "sendTtsEvent" in WEBEXT?.ttsEngine) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendTtsEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.sendTtsEvent);
    },
    "call_SendTtsEvent": (retPtr: Pointer, requestId: number, event: Pointer): void => {
      const event_ffi = {};

      if (A.load.Bool(event + 36)) {
        event_ffi["charIndex"] = A.load.Int64(event + 0);
      }
      event_ffi["errorMessage"] = A.load.Ref(event + 8, undefined);
      if (A.load.Bool(event + 37)) {
        event_ffi["isFinalEvent"] = A.load.Bool(event + 12);
      }
      if (A.load.Bool(event + 38)) {
        event_ffi["length"] = A.load.Int64(event + 16);
      }
      if (A.load.Bool(event + 39)) {
        event_ffi["srcId"] = A.load.Float64(event + 24);
      }
      event_ffi["type"] = A.load.Enum(event + 32, [
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

      const _ret = WEBEXT.ttsEngine.sendTtsEvent(requestId, event_ffi);
    },
    "try_SendTtsEvent": (retPtr: Pointer, errPtr: Pointer, requestId: number, event: Pointer): heap.Ref<boolean> => {
      try {
        const event_ffi = {};

        if (A.load.Bool(event + 36)) {
          event_ffi["charIndex"] = A.load.Int64(event + 0);
        }
        event_ffi["errorMessage"] = A.load.Ref(event + 8, undefined);
        if (A.load.Bool(event + 37)) {
          event_ffi["isFinalEvent"] = A.load.Bool(event + 12);
        }
        if (A.load.Bool(event + 38)) {
          event_ffi["length"] = A.load.Int64(event + 16);
        }
        if (A.load.Bool(event + 39)) {
          event_ffi["srcId"] = A.load.Float64(event + 24);
        }
        event_ffi["type"] = A.load.Enum(event + 32, [
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

        const _ret = WEBEXT.ttsEngine.sendTtsEvent(requestId, event_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateVoices": (): heap.Ref<boolean> => {
      if (WEBEXT?.ttsEngine && "updateVoices" in WEBEXT?.ttsEngine) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateVoices": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.ttsEngine.updateVoices);
    },
    "call_UpdateVoices": (retPtr: Pointer, voices: heap.Ref<object>): void => {
      const _ret = WEBEXT.ttsEngine.updateVoices(A.H.get<object>(voices));
    },
    "try_UpdateVoices": (retPtr: Pointer, errPtr: Pointer, voices: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.ttsEngine.updateVoices(A.H.get<object>(voices));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
