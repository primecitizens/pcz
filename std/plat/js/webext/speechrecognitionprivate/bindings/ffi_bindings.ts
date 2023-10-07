import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/speechrecognitionprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_SpeechRecognitionType": (ref: heap.Ref<string>): number => {
      const idx = ["onDevice", "network"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SpeechRecognitionErrorEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "clientId" in x ? true : false);
        A.store.Int32(ptr + 0, x["clientId"] === undefined ? 0 : (x["clientId"] as number));
        A.store.Ref(ptr + 4, x["message"]);
      }
    },
    "load_SpeechRecognitionErrorEvent": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["clientId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["clientId"];
      }
      x["message"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SpeechRecognitionResultEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 11, true);
        A.store.Bool(ptr + 9, "clientId" in x ? true : false);
        A.store.Int32(ptr + 0, x["clientId"] === undefined ? 0 : (x["clientId"] as number));
        A.store.Ref(ptr + 4, x["transcript"]);
        A.store.Bool(ptr + 10, "isFinal" in x ? true : false);
        A.store.Bool(ptr + 8, x["isFinal"] ? true : false);
      }
    },
    "load_SpeechRecognitionResultEvent": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 9)) {
        x["clientId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["clientId"];
      }
      x["transcript"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 10)) {
        x["isFinal"] = A.load.Bool(ptr + 8);
      } else {
        delete x["isFinal"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SpeechRecognitionStopEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "clientId" in x ? true : false);
        A.store.Int32(ptr + 0, x["clientId"] === undefined ? 0 : (x["clientId"] as number));
      }
    },
    "load_SpeechRecognitionStopEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["clientId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["clientId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_StartOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 11, true);
        A.store.Bool(ptr + 9, "clientId" in x ? true : false);
        A.store.Int32(ptr + 0, x["clientId"] === undefined ? 0 : (x["clientId"] as number));
        A.store.Ref(ptr + 4, x["locale"]);
        A.store.Bool(ptr + 10, "interimResults" in x ? true : false);
        A.store.Bool(ptr + 8, x["interimResults"] ? true : false);
      }
    },
    "load_StartOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 9)) {
        x["clientId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["clientId"];
      }
      x["locale"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 10)) {
        x["interimResults"] = A.load.Bool(ptr + 8);
      } else {
        delete x["interimResults"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_StopOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "clientId" in x ? true : false);
        A.store.Int32(ptr + 0, x["clientId"] === undefined ? 0 : (x["clientId"] as number));
      }
    },
    "load_StopOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["clientId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["clientId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_OnError": (): heap.Ref<boolean> => {
      if (WEBEXT?.speechRecognitionPrivate?.onError && "addListener" in WEBEXT?.speechRecognitionPrivate?.onError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.speechRecognitionPrivate.onError.addListener);
    },
    "call_OnError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.speechRecognitionPrivate.onError.addListener(A.H.get<object>(callback));
    },
    "try_OnError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.speechRecognitionPrivate.onError.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffError": (): heap.Ref<boolean> => {
      if (WEBEXT?.speechRecognitionPrivate?.onError && "removeListener" in WEBEXT?.speechRecognitionPrivate?.onError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.speechRecognitionPrivate.onError.removeListener);
    },
    "call_OffError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.speechRecognitionPrivate.onError.removeListener(A.H.get<object>(callback));
    },
    "try_OffError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.speechRecognitionPrivate.onError.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnError": (): heap.Ref<boolean> => {
      if (WEBEXT?.speechRecognitionPrivate?.onError && "hasListener" in WEBEXT?.speechRecognitionPrivate?.onError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.speechRecognitionPrivate.onError.hasListener);
    },
    "call_HasOnError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.speechRecognitionPrivate.onError.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.speechRecognitionPrivate.onError.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnResult": (): heap.Ref<boolean> => {
      if (WEBEXT?.speechRecognitionPrivate?.onResult && "addListener" in WEBEXT?.speechRecognitionPrivate?.onResult) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.speechRecognitionPrivate.onResult.addListener);
    },
    "call_OnResult": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.speechRecognitionPrivate.onResult.addListener(A.H.get<object>(callback));
    },
    "try_OnResult": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.speechRecognitionPrivate.onResult.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffResult": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.speechRecognitionPrivate?.onResult &&
        "removeListener" in WEBEXT?.speechRecognitionPrivate?.onResult
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.speechRecognitionPrivate.onResult.removeListener);
    },
    "call_OffResult": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.speechRecognitionPrivate.onResult.removeListener(A.H.get<object>(callback));
    },
    "try_OffResult": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.speechRecognitionPrivate.onResult.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnResult": (): heap.Ref<boolean> => {
      if (WEBEXT?.speechRecognitionPrivate?.onResult && "hasListener" in WEBEXT?.speechRecognitionPrivate?.onResult) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.speechRecognitionPrivate.onResult.hasListener);
    },
    "call_HasOnResult": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.speechRecognitionPrivate.onResult.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnResult": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.speechRecognitionPrivate.onResult.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnStop": (): heap.Ref<boolean> => {
      if (WEBEXT?.speechRecognitionPrivate?.onStop && "addListener" in WEBEXT?.speechRecognitionPrivate?.onStop) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnStop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.speechRecognitionPrivate.onStop.addListener);
    },
    "call_OnStop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.speechRecognitionPrivate.onStop.addListener(A.H.get<object>(callback));
    },
    "try_OnStop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.speechRecognitionPrivate.onStop.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffStop": (): heap.Ref<boolean> => {
      if (WEBEXT?.speechRecognitionPrivate?.onStop && "removeListener" in WEBEXT?.speechRecognitionPrivate?.onStop) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffStop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.speechRecognitionPrivate.onStop.removeListener);
    },
    "call_OffStop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.speechRecognitionPrivate.onStop.removeListener(A.H.get<object>(callback));
    },
    "try_OffStop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.speechRecognitionPrivate.onStop.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnStop": (): heap.Ref<boolean> => {
      if (WEBEXT?.speechRecognitionPrivate?.onStop && "hasListener" in WEBEXT?.speechRecognitionPrivate?.onStop) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnStop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.speechRecognitionPrivate.onStop.hasListener);
    },
    "call_HasOnStop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.speechRecognitionPrivate.onStop.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnStop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.speechRecognitionPrivate.onStop.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Start": (): heap.Ref<boolean> => {
      if (WEBEXT?.speechRecognitionPrivate && "start" in WEBEXT?.speechRecognitionPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Start": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.speechRecognitionPrivate.start);
    },
    "call_Start": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 9)) {
        options_ffi["clientId"] = A.load.Int32(options + 0);
      }
      options_ffi["locale"] = A.load.Ref(options + 4, undefined);
      if (A.load.Bool(options + 10)) {
        options_ffi["interimResults"] = A.load.Bool(options + 8);
      }

      const _ret = WEBEXT.speechRecognitionPrivate.start(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Start": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 9)) {
          options_ffi["clientId"] = A.load.Int32(options + 0);
        }
        options_ffi["locale"] = A.load.Ref(options + 4, undefined);
        if (A.load.Bool(options + 10)) {
          options_ffi["interimResults"] = A.load.Bool(options + 8);
        }

        const _ret = WEBEXT.speechRecognitionPrivate.start(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Stop": (): heap.Ref<boolean> => {
      if (WEBEXT?.speechRecognitionPrivate && "stop" in WEBEXT?.speechRecognitionPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Stop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.speechRecognitionPrivate.stop);
    },
    "call_Stop": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 4)) {
        options_ffi["clientId"] = A.load.Int32(options + 0);
      }

      const _ret = WEBEXT.speechRecognitionPrivate.stop(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Stop": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 4)) {
          options_ffi["clientId"] = A.load.Int32(options + 0);
        }

        const _ret = WEBEXT.speechRecognitionPrivate.stop(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
