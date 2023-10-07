import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/audio", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_StreamType": (ref: heap.Ref<string>): number => {
      const idx = ["INPUT", "OUTPUT"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_DeviceType": (ref: heap.Ref<string>): number => {
      const idx = [
        "HEADPHONE",
        "MIC",
        "USB",
        "BLUETOOTH",
        "HDMI",
        "INTERNAL_SPEAKER",
        "INTERNAL_MIC",
        "FRONT_MIC",
        "REAR_MIC",
        "KEYBOARD_MIC",
        "HOTWORD",
        "LINEOUT",
        "POST_MIX_LOOPBACK",
        "POST_DSP_LOOPBACK",
        "ALSA_LOOPBACK",
        "OTHER",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_AudioDeviceInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 34, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Enum(ptr + 8, -1);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 33, false);
        A.store.Int32(ptr + 24, 0);
        A.store.Ref(ptr + 28, undefined);
      } else {
        A.store.Bool(ptr + 34, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Enum(ptr + 4, ["INPUT", "OUTPUT"].indexOf(x["streamType"] as string));
        A.store.Enum(
          ptr + 8,
          [
            "HEADPHONE",
            "MIC",
            "USB",
            "BLUETOOTH",
            "HDMI",
            "INTERNAL_SPEAKER",
            "INTERNAL_MIC",
            "FRONT_MIC",
            "REAR_MIC",
            "KEYBOARD_MIC",
            "HOTWORD",
            "LINEOUT",
            "POST_MIX_LOOPBACK",
            "POST_DSP_LOOPBACK",
            "ALSA_LOOPBACK",
            "OTHER",
          ].indexOf(x["deviceType"] as string)
        );
        A.store.Ref(ptr + 12, x["displayName"]);
        A.store.Ref(ptr + 16, x["deviceName"]);
        A.store.Bool(ptr + 32, "isActive" in x ? true : false);
        A.store.Bool(ptr + 20, x["isActive"] ? true : false);
        A.store.Bool(ptr + 33, "level" in x ? true : false);
        A.store.Int32(ptr + 24, x["level"] === undefined ? 0 : (x["level"] as number));
        A.store.Ref(ptr + 28, x["stableDeviceId"]);
      }
    },
    "load_AudioDeviceInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["streamType"] = A.load.Enum(ptr + 4, ["INPUT", "OUTPUT"]);
      x["deviceType"] = A.load.Enum(ptr + 8, [
        "HEADPHONE",
        "MIC",
        "USB",
        "BLUETOOTH",
        "HDMI",
        "INTERNAL_SPEAKER",
        "INTERNAL_MIC",
        "FRONT_MIC",
        "REAR_MIC",
        "KEYBOARD_MIC",
        "HOTWORD",
        "LINEOUT",
        "POST_MIX_LOOPBACK",
        "POST_DSP_LOOPBACK",
        "ALSA_LOOPBACK",
        "OTHER",
      ]);
      x["displayName"] = A.load.Ref(ptr + 12, undefined);
      x["deviceName"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 32)) {
        x["isActive"] = A.load.Bool(ptr + 20);
      } else {
        delete x["isActive"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["level"] = A.load.Int32(ptr + 24);
      } else {
        delete x["level"];
      }
      x["stableDeviceId"] = A.load.Ref(ptr + 28, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DeviceFilter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 6, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 6, true);
        A.store.Ref(ptr + 0, x["streamTypes"]);
        A.store.Bool(ptr + 5, "isActive" in x ? true : false);
        A.store.Bool(ptr + 4, x["isActive"] ? true : false);
      }
    },
    "load_DeviceFilter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["streamTypes"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 5)) {
        x["isActive"] = A.load.Bool(ptr + 4);
      } else {
        delete x["isActive"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DeviceIdLists": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["input"]);
        A.store.Ref(ptr + 4, x["output"]);
      }
    },
    "load_DeviceIdLists": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["input"] = A.load.Ref(ptr + 0, undefined);
      x["output"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DeviceProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "level" in x ? true : false);
        A.store.Int32(ptr + 0, x["level"] === undefined ? 0 : (x["level"] as number));
      }
    },
    "load_DeviceProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["level"] = A.load.Int32(ptr + 0);
      } else {
        delete x["level"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_LevelChangedEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Ref(ptr + 0, x["deviceId"]);
        A.store.Bool(ptr + 8, "level" in x ? true : false);
        A.store.Int32(ptr + 4, x["level"] === undefined ? 0 : (x["level"] as number));
      }
    },
    "load_LevelChangedEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["deviceId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 8)) {
        x["level"] = A.load.Int32(ptr + 4);
      } else {
        delete x["level"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MuteChangedEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 6, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 6, true);
        A.store.Enum(ptr + 0, ["INPUT", "OUTPUT"].indexOf(x["streamType"] as string));
        A.store.Bool(ptr + 5, "isMuted" in x ? true : false);
        A.store.Bool(ptr + 4, x["isMuted"] ? true : false);
      }
    },
    "load_MuteChangedEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["streamType"] = A.load.Enum(ptr + 0, ["INPUT", "OUTPUT"]);
      if (A.load.Bool(ptr + 5)) {
        x["isMuted"] = A.load.Bool(ptr + 4);
      } else {
        delete x["isMuted"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_GetDevices": (): heap.Ref<boolean> => {
      if (WEBEXT?.audio && "getDevices" in WEBEXT?.audio) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDevices": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.audio.getDevices);
    },
    "call_GetDevices": (retPtr: Pointer, filter: Pointer): void => {
      const filter_ffi = {};

      filter_ffi["streamTypes"] = A.load.Ref(filter + 0, undefined);
      if (A.load.Bool(filter + 5)) {
        filter_ffi["isActive"] = A.load.Bool(filter + 4);
      }

      const _ret = WEBEXT.audio.getDevices(filter_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDevices": (retPtr: Pointer, errPtr: Pointer, filter: Pointer): heap.Ref<boolean> => {
      try {
        const filter_ffi = {};

        filter_ffi["streamTypes"] = A.load.Ref(filter + 0, undefined);
        if (A.load.Bool(filter + 5)) {
          filter_ffi["isActive"] = A.load.Bool(filter + 4);
        }

        const _ret = WEBEXT.audio.getDevices(filter_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetMute": (): heap.Ref<boolean> => {
      if (WEBEXT?.audio && "getMute" in WEBEXT?.audio) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetMute": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.audio.getMute);
    },
    "call_GetMute": (retPtr: Pointer, streamType: number): void => {
      const _ret = WEBEXT.audio.getMute(
        streamType > 0 && streamType <= 2 ? ["INPUT", "OUTPUT"][streamType - 1] : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_GetMute": (retPtr: Pointer, errPtr: Pointer, streamType: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.audio.getMute(
          streamType > 0 && streamType <= 2 ? ["INPUT", "OUTPUT"][streamType - 1] : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeviceListChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.audio?.onDeviceListChanged && "addListener" in WEBEXT?.audio?.onDeviceListChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeviceListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.audio.onDeviceListChanged.addListener);
    },
    "call_OnDeviceListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.audio.onDeviceListChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnDeviceListChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.audio.onDeviceListChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeviceListChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.audio?.onDeviceListChanged && "removeListener" in WEBEXT?.audio?.onDeviceListChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeviceListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.audio.onDeviceListChanged.removeListener);
    },
    "call_OffDeviceListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.audio.onDeviceListChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeviceListChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.audio.onDeviceListChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeviceListChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.audio?.onDeviceListChanged && "hasListener" in WEBEXT?.audio?.onDeviceListChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeviceListChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.audio.onDeviceListChanged.hasListener);
    },
    "call_HasOnDeviceListChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.audio.onDeviceListChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeviceListChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.audio.onDeviceListChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnLevelChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.audio?.onLevelChanged && "addListener" in WEBEXT?.audio?.onLevelChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnLevelChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.audio.onLevelChanged.addListener);
    },
    "call_OnLevelChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.audio.onLevelChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnLevelChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.audio.onLevelChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffLevelChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.audio?.onLevelChanged && "removeListener" in WEBEXT?.audio?.onLevelChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffLevelChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.audio.onLevelChanged.removeListener);
    },
    "call_OffLevelChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.audio.onLevelChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffLevelChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.audio.onLevelChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnLevelChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.audio?.onLevelChanged && "hasListener" in WEBEXT?.audio?.onLevelChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnLevelChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.audio.onLevelChanged.hasListener);
    },
    "call_HasOnLevelChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.audio.onLevelChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnLevelChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.audio.onLevelChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMuteChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.audio?.onMuteChanged && "addListener" in WEBEXT?.audio?.onMuteChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMuteChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.audio.onMuteChanged.addListener);
    },
    "call_OnMuteChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.audio.onMuteChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnMuteChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.audio.onMuteChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMuteChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.audio?.onMuteChanged && "removeListener" in WEBEXT?.audio?.onMuteChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMuteChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.audio.onMuteChanged.removeListener);
    },
    "call_OffMuteChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.audio.onMuteChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffMuteChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.audio.onMuteChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMuteChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.audio?.onMuteChanged && "hasListener" in WEBEXT?.audio?.onMuteChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMuteChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.audio.onMuteChanged.hasListener);
    },
    "call_HasOnMuteChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.audio.onMuteChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMuteChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.audio.onMuteChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetActiveDevices": (): heap.Ref<boolean> => {
      if (WEBEXT?.audio && "setActiveDevices" in WEBEXT?.audio) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetActiveDevices": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.audio.setActiveDevices);
    },
    "call_SetActiveDevices": (retPtr: Pointer, ids: Pointer): void => {
      const ids_ffi = {};

      ids_ffi["input"] = A.load.Ref(ids + 0, undefined);
      ids_ffi["output"] = A.load.Ref(ids + 4, undefined);

      const _ret = WEBEXT.audio.setActiveDevices(ids_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetActiveDevices": (retPtr: Pointer, errPtr: Pointer, ids: Pointer): heap.Ref<boolean> => {
      try {
        const ids_ffi = {};

        ids_ffi["input"] = A.load.Ref(ids + 0, undefined);
        ids_ffi["output"] = A.load.Ref(ids + 4, undefined);

        const _ret = WEBEXT.audio.setActiveDevices(ids_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetMute": (): heap.Ref<boolean> => {
      if (WEBEXT?.audio && "setMute" in WEBEXT?.audio) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetMute": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.audio.setMute);
    },
    "call_SetMute": (retPtr: Pointer, streamType: number, isMuted: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.audio.setMute(
        streamType > 0 && streamType <= 2 ? ["INPUT", "OUTPUT"][streamType - 1] : undefined,
        isMuted === A.H.TRUE
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_SetMute": (
      retPtr: Pointer,
      errPtr: Pointer,
      streamType: number,
      isMuted: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.audio.setMute(
          streamType > 0 && streamType <= 2 ? ["INPUT", "OUTPUT"][streamType - 1] : undefined,
          isMuted === A.H.TRUE
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetProperties": (): heap.Ref<boolean> => {
      if (WEBEXT?.audio && "setProperties" in WEBEXT?.audio) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetProperties": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.audio.setProperties);
    },
    "call_SetProperties": (retPtr: Pointer, id: heap.Ref<object>, properties: Pointer): void => {
      const properties_ffi = {};

      if (A.load.Bool(properties + 4)) {
        properties_ffi["level"] = A.load.Int32(properties + 0);
      }

      const _ret = WEBEXT.audio.setProperties(A.H.get<object>(id), properties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetProperties": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<object>,
      properties: Pointer
    ): heap.Ref<boolean> => {
      try {
        const properties_ffi = {};

        if (A.load.Bool(properties + 4)) {
          properties_ffi["level"] = A.load.Int32(properties + 0);
        }

        const _ret = WEBEXT.audio.setProperties(A.H.get<object>(id), properties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
