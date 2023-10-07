import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/serial", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_DataBits": (ref: heap.Ref<string>): number => {
      const idx = ["seven", "eight"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ParityBit": (ref: heap.Ref<string>): number => {
      const idx = ["no", "odd", "even"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_StopBits": (ref: heap.Ref<string>): number => {
      const idx = ["one", "two"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ConnectionInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 49, false);
        A.store.Bool(ptr + 41, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 42, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 43, false);
        A.store.Bool(ptr + 5, false);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 44, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 45, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Bool(ptr + 46, false);
        A.store.Int32(ptr + 20, 0);
        A.store.Bool(ptr + 47, false);
        A.store.Int32(ptr + 24, 0);
        A.store.Enum(ptr + 28, -1);
        A.store.Enum(ptr + 32, -1);
        A.store.Enum(ptr + 36, -1);
        A.store.Bool(ptr + 48, false);
        A.store.Bool(ptr + 40, false);
      } else {
        A.store.Bool(ptr + 49, true);
        A.store.Bool(ptr + 41, "connectionId" in x ? true : false);
        A.store.Int32(ptr + 0, x["connectionId"] === undefined ? 0 : (x["connectionId"] as number));
        A.store.Bool(ptr + 42, "paused" in x ? true : false);
        A.store.Bool(ptr + 4, x["paused"] ? true : false);
        A.store.Bool(ptr + 43, "persistent" in x ? true : false);
        A.store.Bool(ptr + 5, x["persistent"] ? true : false);
        A.store.Ref(ptr + 8, x["name"]);
        A.store.Bool(ptr + 44, "bufferSize" in x ? true : false);
        A.store.Int32(ptr + 12, x["bufferSize"] === undefined ? 0 : (x["bufferSize"] as number));
        A.store.Bool(ptr + 45, "receiveTimeout" in x ? true : false);
        A.store.Int32(ptr + 16, x["receiveTimeout"] === undefined ? 0 : (x["receiveTimeout"] as number));
        A.store.Bool(ptr + 46, "sendTimeout" in x ? true : false);
        A.store.Int32(ptr + 20, x["sendTimeout"] === undefined ? 0 : (x["sendTimeout"] as number));
        A.store.Bool(ptr + 47, "bitrate" in x ? true : false);
        A.store.Int32(ptr + 24, x["bitrate"] === undefined ? 0 : (x["bitrate"] as number));
        A.store.Enum(ptr + 28, ["seven", "eight"].indexOf(x["dataBits"] as string));
        A.store.Enum(ptr + 32, ["no", "odd", "even"].indexOf(x["parityBit"] as string));
        A.store.Enum(ptr + 36, ["one", "two"].indexOf(x["stopBits"] as string));
        A.store.Bool(ptr + 48, "ctsFlowControl" in x ? true : false);
        A.store.Bool(ptr + 40, x["ctsFlowControl"] ? true : false);
      }
    },
    "load_ConnectionInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 41)) {
        x["connectionId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["connectionId"];
      }
      if (A.load.Bool(ptr + 42)) {
        x["paused"] = A.load.Bool(ptr + 4);
      } else {
        delete x["paused"];
      }
      if (A.load.Bool(ptr + 43)) {
        x["persistent"] = A.load.Bool(ptr + 5);
      } else {
        delete x["persistent"];
      }
      x["name"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 44)) {
        x["bufferSize"] = A.load.Int32(ptr + 12);
      } else {
        delete x["bufferSize"];
      }
      if (A.load.Bool(ptr + 45)) {
        x["receiveTimeout"] = A.load.Int32(ptr + 16);
      } else {
        delete x["receiveTimeout"];
      }
      if (A.load.Bool(ptr + 46)) {
        x["sendTimeout"] = A.load.Int32(ptr + 20);
      } else {
        delete x["sendTimeout"];
      }
      if (A.load.Bool(ptr + 47)) {
        x["bitrate"] = A.load.Int32(ptr + 24);
      } else {
        delete x["bitrate"];
      }
      x["dataBits"] = A.load.Enum(ptr + 28, ["seven", "eight"]);
      x["parityBit"] = A.load.Enum(ptr + 32, ["no", "odd", "even"]);
      x["stopBits"] = A.load.Enum(ptr + 36, ["one", "two"]);
      if (A.load.Bool(ptr + 48)) {
        x["ctsFlowControl"] = A.load.Bool(ptr + 40);
      } else {
        delete x["ctsFlowControl"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ConnectionOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 46, false);
        A.store.Bool(ptr + 40, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 41, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 42, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Enum(ptr + 16, -1);
        A.store.Enum(ptr + 20, -1);
        A.store.Enum(ptr + 24, -1);
        A.store.Bool(ptr + 43, false);
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 44, false);
        A.store.Int32(ptr + 32, 0);
        A.store.Bool(ptr + 45, false);
        A.store.Int32(ptr + 36, 0);
      } else {
        A.store.Bool(ptr + 46, true);
        A.store.Bool(ptr + 40, "persistent" in x ? true : false);
        A.store.Bool(ptr + 0, x["persistent"] ? true : false);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Bool(ptr + 41, "bufferSize" in x ? true : false);
        A.store.Int32(ptr + 8, x["bufferSize"] === undefined ? 0 : (x["bufferSize"] as number));
        A.store.Bool(ptr + 42, "bitrate" in x ? true : false);
        A.store.Int32(ptr + 12, x["bitrate"] === undefined ? 0 : (x["bitrate"] as number));
        A.store.Enum(ptr + 16, ["seven", "eight"].indexOf(x["dataBits"] as string));
        A.store.Enum(ptr + 20, ["no", "odd", "even"].indexOf(x["parityBit"] as string));
        A.store.Enum(ptr + 24, ["one", "two"].indexOf(x["stopBits"] as string));
        A.store.Bool(ptr + 43, "ctsFlowControl" in x ? true : false);
        A.store.Bool(ptr + 28, x["ctsFlowControl"] ? true : false);
        A.store.Bool(ptr + 44, "receiveTimeout" in x ? true : false);
        A.store.Int32(ptr + 32, x["receiveTimeout"] === undefined ? 0 : (x["receiveTimeout"] as number));
        A.store.Bool(ptr + 45, "sendTimeout" in x ? true : false);
        A.store.Int32(ptr + 36, x["sendTimeout"] === undefined ? 0 : (x["sendTimeout"] as number));
      }
    },
    "load_ConnectionOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 40)) {
        x["persistent"] = A.load.Bool(ptr + 0);
      } else {
        delete x["persistent"];
      }
      x["name"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 41)) {
        x["bufferSize"] = A.load.Int32(ptr + 8);
      } else {
        delete x["bufferSize"];
      }
      if (A.load.Bool(ptr + 42)) {
        x["bitrate"] = A.load.Int32(ptr + 12);
      } else {
        delete x["bitrate"];
      }
      x["dataBits"] = A.load.Enum(ptr + 16, ["seven", "eight"]);
      x["parityBit"] = A.load.Enum(ptr + 20, ["no", "odd", "even"]);
      x["stopBits"] = A.load.Enum(ptr + 24, ["one", "two"]);
      if (A.load.Bool(ptr + 43)) {
        x["ctsFlowControl"] = A.load.Bool(ptr + 28);
      } else {
        delete x["ctsFlowControl"];
      }
      if (A.load.Bool(ptr + 44)) {
        x["receiveTimeout"] = A.load.Int32(ptr + 32);
      } else {
        delete x["receiveTimeout"];
      }
      if (A.load.Bool(ptr + 45)) {
        x["sendTimeout"] = A.load.Int32(ptr + 36);
      } else {
        delete x["sendTimeout"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DeviceControlSignals": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 3, false);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Bool(ptr + 4, "dcd" in x ? true : false);
        A.store.Bool(ptr + 0, x["dcd"] ? true : false);
        A.store.Bool(ptr + 5, "cts" in x ? true : false);
        A.store.Bool(ptr + 1, x["cts"] ? true : false);
        A.store.Bool(ptr + 6, "ri" in x ? true : false);
        A.store.Bool(ptr + 2, x["ri"] ? true : false);
        A.store.Bool(ptr + 7, "dsr" in x ? true : false);
        A.store.Bool(ptr + 3, x["dsr"] ? true : false);
      }
    },
    "load_DeviceControlSignals": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["dcd"] = A.load.Bool(ptr + 0);
      } else {
        delete x["dcd"];
      }
      if (A.load.Bool(ptr + 5)) {
        x["cts"] = A.load.Bool(ptr + 1);
      } else {
        delete x["cts"];
      }
      if (A.load.Bool(ptr + 6)) {
        x["ri"] = A.load.Bool(ptr + 2);
      } else {
        delete x["ri"];
      }
      if (A.load.Bool(ptr + 7)) {
        x["dsr"] = A.load.Bool(ptr + 3);
      } else {
        delete x["dsr"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DeviceInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Ref(ptr + 0, x["path"]);
        A.store.Bool(ptr + 16, "vendorId" in x ? true : false);
        A.store.Int32(ptr + 4, x["vendorId"] === undefined ? 0 : (x["vendorId"] as number));
        A.store.Bool(ptr + 17, "productId" in x ? true : false);
        A.store.Int32(ptr + 8, x["productId"] === undefined ? 0 : (x["productId"] as number));
        A.store.Ref(ptr + 12, x["displayName"]);
      }
    },
    "load_DeviceInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["path"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["vendorId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["vendorId"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["productId"] = A.load.Int32(ptr + 8);
      } else {
        delete x["productId"];
      }
      x["displayName"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_HostControlSignals": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 1, false);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Bool(ptr + 2, "dtr" in x ? true : false);
        A.store.Bool(ptr + 0, x["dtr"] ? true : false);
        A.store.Bool(ptr + 3, "rts" in x ? true : false);
        A.store.Bool(ptr + 1, x["rts"] ? true : false);
      }
    },
    "load_HostControlSignals": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 2)) {
        x["dtr"] = A.load.Bool(ptr + 0);
      } else {
        delete x["dtr"];
      }
      if (A.load.Bool(ptr + 3)) {
        x["rts"] = A.load.Bool(ptr + 1);
      } else {
        delete x["rts"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ReceiveError": (ref: heap.Ref<string>): number => {
      const idx = [
        "disconnected",
        "timeout",
        "device_lost",
        "break",
        "frame_error",
        "overrun",
        "buffer_overflow",
        "parity_error",
        "system_error",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ReceiveErrorInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "connectionId" in x ? true : false);
        A.store.Int32(ptr + 0, x["connectionId"] === undefined ? 0 : (x["connectionId"] as number));
        A.store.Enum(
          ptr + 4,
          [
            "disconnected",
            "timeout",
            "device_lost",
            "break",
            "frame_error",
            "overrun",
            "buffer_overflow",
            "parity_error",
            "system_error",
          ].indexOf(x["error"] as string)
        );
      }
    },
    "load_ReceiveErrorInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["connectionId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["connectionId"];
      }
      x["error"] = A.load.Enum(ptr + 4, [
        "disconnected",
        "timeout",
        "device_lost",
        "break",
        "frame_error",
        "overrun",
        "buffer_overflow",
        "parity_error",
        "system_error",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ReceiveInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "connectionId" in x ? true : false);
        A.store.Int32(ptr + 0, x["connectionId"] === undefined ? 0 : (x["connectionId"] as number));
        A.store.Ref(ptr + 4, x["data"]);
      }
    },
    "load_ReceiveInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["connectionId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["connectionId"];
      }
      x["data"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_SendError": (ref: heap.Ref<string>): number => {
      const idx = ["disconnected", "pending", "timeout", "system_error"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SendInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "bytesSent" in x ? true : false);
        A.store.Int32(ptr + 0, x["bytesSent"] === undefined ? 0 : (x["bytesSent"] as number));
        A.store.Enum(ptr + 4, ["disconnected", "pending", "timeout", "system_error"].indexOf(x["error"] as string));
      }
    },
    "load_SendInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["bytesSent"] = A.load.Int32(ptr + 0);
      } else {
        delete x["bytesSent"];
      }
      x["error"] = A.load.Enum(ptr + 4, ["disconnected", "pending", "timeout", "system_error"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_ClearBreak": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial && "clearBreak" in WEBEXT?.serial) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ClearBreak": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.clearBreak);
    },
    "call_ClearBreak": (retPtr: Pointer, connectionId: number): void => {
      const _ret = WEBEXT.serial.clearBreak(connectionId);
      A.store.Ref(retPtr, _ret);
    },
    "try_ClearBreak": (retPtr: Pointer, errPtr: Pointer, connectionId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.clearBreak(connectionId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Connect": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial && "connect" in WEBEXT?.serial) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Connect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.connect);
    },
    "call_Connect": (retPtr: Pointer, path: heap.Ref<object>, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 40)) {
        options_ffi["persistent"] = A.load.Bool(options + 0);
      }
      options_ffi["name"] = A.load.Ref(options + 4, undefined);
      if (A.load.Bool(options + 41)) {
        options_ffi["bufferSize"] = A.load.Int32(options + 8);
      }
      if (A.load.Bool(options + 42)) {
        options_ffi["bitrate"] = A.load.Int32(options + 12);
      }
      options_ffi["dataBits"] = A.load.Enum(options + 16, ["seven", "eight"]);
      options_ffi["parityBit"] = A.load.Enum(options + 20, ["no", "odd", "even"]);
      options_ffi["stopBits"] = A.load.Enum(options + 24, ["one", "two"]);
      if (A.load.Bool(options + 43)) {
        options_ffi["ctsFlowControl"] = A.load.Bool(options + 28);
      }
      if (A.load.Bool(options + 44)) {
        options_ffi["receiveTimeout"] = A.load.Int32(options + 32);
      }
      if (A.load.Bool(options + 45)) {
        options_ffi["sendTimeout"] = A.load.Int32(options + 36);
      }

      const _ret = WEBEXT.serial.connect(A.H.get<object>(path), options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Connect": (retPtr: Pointer, errPtr: Pointer, path: heap.Ref<object>, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 40)) {
          options_ffi["persistent"] = A.load.Bool(options + 0);
        }
        options_ffi["name"] = A.load.Ref(options + 4, undefined);
        if (A.load.Bool(options + 41)) {
          options_ffi["bufferSize"] = A.load.Int32(options + 8);
        }
        if (A.load.Bool(options + 42)) {
          options_ffi["bitrate"] = A.load.Int32(options + 12);
        }
        options_ffi["dataBits"] = A.load.Enum(options + 16, ["seven", "eight"]);
        options_ffi["parityBit"] = A.load.Enum(options + 20, ["no", "odd", "even"]);
        options_ffi["stopBits"] = A.load.Enum(options + 24, ["one", "two"]);
        if (A.load.Bool(options + 43)) {
          options_ffi["ctsFlowControl"] = A.load.Bool(options + 28);
        }
        if (A.load.Bool(options + 44)) {
          options_ffi["receiveTimeout"] = A.load.Int32(options + 32);
        }
        if (A.load.Bool(options + 45)) {
          options_ffi["sendTimeout"] = A.load.Int32(options + 36);
        }

        const _ret = WEBEXT.serial.connect(A.H.get<object>(path), options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Disconnect": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial && "disconnect" in WEBEXT?.serial) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Disconnect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.disconnect);
    },
    "call_Disconnect": (retPtr: Pointer, connectionId: number): void => {
      const _ret = WEBEXT.serial.disconnect(connectionId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Disconnect": (retPtr: Pointer, errPtr: Pointer, connectionId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.disconnect(connectionId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Flush": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial && "flush" in WEBEXT?.serial) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Flush": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.flush);
    },
    "call_Flush": (retPtr: Pointer, connectionId: number): void => {
      const _ret = WEBEXT.serial.flush(connectionId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Flush": (retPtr: Pointer, errPtr: Pointer, connectionId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.flush(connectionId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetConnections": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial && "getConnections" in WEBEXT?.serial) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetConnections": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.getConnections);
    },
    "call_GetConnections": (retPtr: Pointer): void => {
      const _ret = WEBEXT.serial.getConnections();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetConnections": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.getConnections();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetControlSignals": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial && "getControlSignals" in WEBEXT?.serial) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetControlSignals": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.getControlSignals);
    },
    "call_GetControlSignals": (retPtr: Pointer, connectionId: number): void => {
      const _ret = WEBEXT.serial.getControlSignals(connectionId);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetControlSignals": (retPtr: Pointer, errPtr: Pointer, connectionId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.getControlSignals(connectionId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDevices": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial && "getDevices" in WEBEXT?.serial) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDevices": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.getDevices);
    },
    "call_GetDevices": (retPtr: Pointer): void => {
      const _ret = WEBEXT.serial.getDevices();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDevices": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.getDevices();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial && "getInfo" in WEBEXT?.serial) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.getInfo);
    },
    "call_GetInfo": (retPtr: Pointer, connectionId: number): void => {
      const _ret = WEBEXT.serial.getInfo(connectionId);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetInfo": (retPtr: Pointer, errPtr: Pointer, connectionId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.getInfo(connectionId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnReceive": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial?.onReceive && "addListener" in WEBEXT?.serial?.onReceive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnReceive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.onReceive.addListener);
    },
    "call_OnReceive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.serial.onReceive.addListener(A.H.get<object>(callback));
    },
    "try_OnReceive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.onReceive.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffReceive": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial?.onReceive && "removeListener" in WEBEXT?.serial?.onReceive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffReceive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.onReceive.removeListener);
    },
    "call_OffReceive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.serial.onReceive.removeListener(A.H.get<object>(callback));
    },
    "try_OffReceive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.onReceive.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnReceive": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial?.onReceive && "hasListener" in WEBEXT?.serial?.onReceive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnReceive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.onReceive.hasListener);
    },
    "call_HasOnReceive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.serial.onReceive.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnReceive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.onReceive.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnReceiveError": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial?.onReceiveError && "addListener" in WEBEXT?.serial?.onReceiveError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnReceiveError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.onReceiveError.addListener);
    },
    "call_OnReceiveError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.serial.onReceiveError.addListener(A.H.get<object>(callback));
    },
    "try_OnReceiveError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.onReceiveError.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffReceiveError": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial?.onReceiveError && "removeListener" in WEBEXT?.serial?.onReceiveError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffReceiveError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.onReceiveError.removeListener);
    },
    "call_OffReceiveError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.serial.onReceiveError.removeListener(A.H.get<object>(callback));
    },
    "try_OffReceiveError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.onReceiveError.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnReceiveError": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial?.onReceiveError && "hasListener" in WEBEXT?.serial?.onReceiveError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnReceiveError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.onReceiveError.hasListener);
    },
    "call_HasOnReceiveError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.serial.onReceiveError.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnReceiveError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.onReceiveError.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Send": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial && "send" in WEBEXT?.serial) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Send": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.send);
    },
    "call_Send": (retPtr: Pointer, connectionId: number, data: heap.Ref<object>): void => {
      const _ret = WEBEXT.serial.send(connectionId, A.H.get<object>(data));
      A.store.Ref(retPtr, _ret);
    },
    "try_Send": (retPtr: Pointer, errPtr: Pointer, connectionId: number, data: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.send(connectionId, A.H.get<object>(data));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetBreak": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial && "setBreak" in WEBEXT?.serial) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetBreak": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.setBreak);
    },
    "call_SetBreak": (retPtr: Pointer, connectionId: number): void => {
      const _ret = WEBEXT.serial.setBreak(connectionId);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetBreak": (retPtr: Pointer, errPtr: Pointer, connectionId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.setBreak(connectionId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetControlSignals": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial && "setControlSignals" in WEBEXT?.serial) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetControlSignals": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.setControlSignals);
    },
    "call_SetControlSignals": (retPtr: Pointer, connectionId: number, signals: Pointer): void => {
      const signals_ffi = {};

      if (A.load.Bool(signals + 2)) {
        signals_ffi["dtr"] = A.load.Bool(signals + 0);
      }
      if (A.load.Bool(signals + 3)) {
        signals_ffi["rts"] = A.load.Bool(signals + 1);
      }

      const _ret = WEBEXT.serial.setControlSignals(connectionId, signals_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetControlSignals": (
      retPtr: Pointer,
      errPtr: Pointer,
      connectionId: number,
      signals: Pointer
    ): heap.Ref<boolean> => {
      try {
        const signals_ffi = {};

        if (A.load.Bool(signals + 2)) {
          signals_ffi["dtr"] = A.load.Bool(signals + 0);
        }
        if (A.load.Bool(signals + 3)) {
          signals_ffi["rts"] = A.load.Bool(signals + 1);
        }

        const _ret = WEBEXT.serial.setControlSignals(connectionId, signals_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPaused": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial && "setPaused" in WEBEXT?.serial) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPaused": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.setPaused);
    },
    "call_SetPaused": (retPtr: Pointer, connectionId: number, paused: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.serial.setPaused(connectionId, paused === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetPaused": (
      retPtr: Pointer,
      errPtr: Pointer,
      connectionId: number,
      paused: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.serial.setPaused(connectionId, paused === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Update": (): heap.Ref<boolean> => {
      if (WEBEXT?.serial && "update" in WEBEXT?.serial) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Update": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.serial.update);
    },
    "call_Update": (retPtr: Pointer, connectionId: number, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 40)) {
        options_ffi["persistent"] = A.load.Bool(options + 0);
      }
      options_ffi["name"] = A.load.Ref(options + 4, undefined);
      if (A.load.Bool(options + 41)) {
        options_ffi["bufferSize"] = A.load.Int32(options + 8);
      }
      if (A.load.Bool(options + 42)) {
        options_ffi["bitrate"] = A.load.Int32(options + 12);
      }
      options_ffi["dataBits"] = A.load.Enum(options + 16, ["seven", "eight"]);
      options_ffi["parityBit"] = A.load.Enum(options + 20, ["no", "odd", "even"]);
      options_ffi["stopBits"] = A.load.Enum(options + 24, ["one", "two"]);
      if (A.load.Bool(options + 43)) {
        options_ffi["ctsFlowControl"] = A.load.Bool(options + 28);
      }
      if (A.load.Bool(options + 44)) {
        options_ffi["receiveTimeout"] = A.load.Int32(options + 32);
      }
      if (A.load.Bool(options + 45)) {
        options_ffi["sendTimeout"] = A.load.Int32(options + 36);
      }

      const _ret = WEBEXT.serial.update(connectionId, options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Update": (retPtr: Pointer, errPtr: Pointer, connectionId: number, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 40)) {
          options_ffi["persistent"] = A.load.Bool(options + 0);
        }
        options_ffi["name"] = A.load.Ref(options + 4, undefined);
        if (A.load.Bool(options + 41)) {
          options_ffi["bufferSize"] = A.load.Int32(options + 8);
        }
        if (A.load.Bool(options + 42)) {
          options_ffi["bitrate"] = A.load.Int32(options + 12);
        }
        options_ffi["dataBits"] = A.load.Enum(options + 16, ["seven", "eight"]);
        options_ffi["parityBit"] = A.load.Enum(options + 20, ["no", "odd", "even"]);
        options_ffi["stopBits"] = A.load.Enum(options + 24, ["one", "two"]);
        if (A.load.Bool(options + 43)) {
          options_ffi["ctsFlowControl"] = A.load.Bool(options + 28);
        }
        if (A.load.Bool(options + 44)) {
          options_ffi["receiveTimeout"] = A.load.Int32(options + 32);
        }
        if (A.load.Bool(options + 45)) {
          options_ffi["sendTimeout"] = A.load.Int32(options + 36);
        }

        const _ret = WEBEXT.serial.update(connectionId, options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
