import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/bluetoothsocket", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_AcceptError": (ref: heap.Ref<string>): number => {
      const idx = ["system_error", "not_listening"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_AcceptErrorInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "socketId" in x ? true : false);
        A.store.Int32(ptr + 0, x["socketId"] === undefined ? 0 : (x["socketId"] as number));
        A.store.Ref(ptr + 4, x["errorMessage"]);
        A.store.Enum(ptr + 8, ["system_error", "not_listening"].indexOf(x["error"] as string));
      }
    },
    "load_AcceptErrorInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["socketId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["socketId"];
      }
      x["errorMessage"] = A.load.Ref(ptr + 4, undefined);
      x["error"] = A.load.Enum(ptr + 8, ["system_error", "not_listening"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AcceptInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "socketId" in x ? true : false);
        A.store.Int32(ptr + 0, x["socketId"] === undefined ? 0 : (x["socketId"] as number));
        A.store.Bool(ptr + 9, "clientSocketId" in x ? true : false);
        A.store.Int32(ptr + 4, x["clientSocketId"] === undefined ? 0 : (x["clientSocketId"] as number));
      }
    },
    "load_AcceptInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["socketId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["socketId"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["clientSocketId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["clientSocketId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CreateInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "socketId" in x ? true : false);
        A.store.Int32(ptr + 0, x["socketId"] === undefined ? 0 : (x["socketId"] as number));
      }
    },
    "load_CreateInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["socketId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["socketId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SocketInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 33, false);
        A.store.Bool(ptr + 28, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 29, false);
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 30, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 31, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 33, true);
        A.store.Bool(ptr + 28, "socketId" in x ? true : false);
        A.store.Int32(ptr + 0, x["socketId"] === undefined ? 0 : (x["socketId"] as number));
        A.store.Bool(ptr + 29, "persistent" in x ? true : false);
        A.store.Bool(ptr + 4, x["persistent"] ? true : false);
        A.store.Ref(ptr + 8, x["name"]);
        A.store.Bool(ptr + 30, "bufferSize" in x ? true : false);
        A.store.Int32(ptr + 12, x["bufferSize"] === undefined ? 0 : (x["bufferSize"] as number));
        A.store.Bool(ptr + 31, "paused" in x ? true : false);
        A.store.Bool(ptr + 16, x["paused"] ? true : false);
        A.store.Bool(ptr + 32, "connected" in x ? true : false);
        A.store.Bool(ptr + 17, x["connected"] ? true : false);
        A.store.Ref(ptr + 20, x["address"]);
        A.store.Ref(ptr + 24, x["uuid"]);
      }
    },
    "load_SocketInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 28)) {
        x["socketId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["socketId"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["persistent"] = A.load.Bool(ptr + 4);
      } else {
        delete x["persistent"];
      }
      x["name"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 30)) {
        x["bufferSize"] = A.load.Int32(ptr + 12);
      } else {
        delete x["bufferSize"];
      }
      if (A.load.Bool(ptr + 31)) {
        x["paused"] = A.load.Bool(ptr + 16);
      } else {
        delete x["paused"];
      }
      if (A.load.Bool(ptr + 32)) {
        x["connected"] = A.load.Bool(ptr + 17);
      } else {
        delete x["connected"];
      }
      x["address"] = A.load.Ref(ptr + 20, undefined);
      x["uuid"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ListenOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 14, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 15, true);
        A.store.Bool(ptr + 12, "channel" in x ? true : false);
        A.store.Int32(ptr + 0, x["channel"] === undefined ? 0 : (x["channel"] as number));
        A.store.Bool(ptr + 13, "psm" in x ? true : false);
        A.store.Int32(ptr + 4, x["psm"] === undefined ? 0 : (x["psm"] as number));
        A.store.Bool(ptr + 14, "backlog" in x ? true : false);
        A.store.Int32(ptr + 8, x["backlog"] === undefined ? 0 : (x["backlog"] as number));
      }
    },
    "load_ListenOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["channel"] = A.load.Int32(ptr + 0);
      } else {
        delete x["channel"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["psm"] = A.load.Int32(ptr + 4);
      } else {
        delete x["psm"];
      }
      if (A.load.Bool(ptr + 14)) {
        x["backlog"] = A.load.Int32(ptr + 8);
      } else {
        delete x["backlog"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ReceiveError": (ref: heap.Ref<string>): number => {
      const idx = ["disconnected", "system_error", "not_connected"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ReceiveErrorInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Bool(ptr + 12, "socketId" in x ? true : false);
        A.store.Int32(ptr + 0, x["socketId"] === undefined ? 0 : (x["socketId"] as number));
        A.store.Ref(ptr + 4, x["errorMessage"]);
        A.store.Enum(ptr + 8, ["disconnected", "system_error", "not_connected"].indexOf(x["error"] as string));
      }
    },
    "load_ReceiveErrorInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["socketId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["socketId"];
      }
      x["errorMessage"] = A.load.Ref(ptr + 4, undefined);
      x["error"] = A.load.Enum(ptr + 8, ["disconnected", "system_error", "not_connected"]);
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
        A.store.Bool(ptr + 8, "socketId" in x ? true : false);
        A.store.Int32(ptr + 0, x["socketId"] === undefined ? 0 : (x["socketId"] as number));
        A.store.Ref(ptr + 4, x["data"]);
      }
    },
    "load_ReceiveInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["socketId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["socketId"];
      }
      x["data"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SocketProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Bool(ptr + 12, "persistent" in x ? true : false);
        A.store.Bool(ptr + 0, x["persistent"] ? true : false);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Bool(ptr + 13, "bufferSize" in x ? true : false);
        A.store.Int32(ptr + 8, x["bufferSize"] === undefined ? 0 : (x["bufferSize"] as number));
      }
    },
    "load_SocketProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["persistent"] = A.load.Bool(ptr + 0);
      } else {
        delete x["persistent"];
      }
      x["name"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 13)) {
        x["bufferSize"] = A.load.Int32(ptr + 8);
      } else {
        delete x["bufferSize"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Close": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket && "close" in WEBEXT?.bluetoothSocket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Close": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.close);
    },
    "call_Close": (retPtr: Pointer, socketId: number): void => {
      const _ret = WEBEXT.bluetoothSocket.close(socketId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Close": (retPtr: Pointer, errPtr: Pointer, socketId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.close(socketId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Connect": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket && "connect" in WEBEXT?.bluetoothSocket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Connect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.connect);
    },
    "call_Connect": (retPtr: Pointer, socketId: number, address: heap.Ref<object>, uuid: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothSocket.connect(socketId, A.H.get<object>(address), A.H.get<object>(uuid));
      A.store.Ref(retPtr, _ret);
    },
    "try_Connect": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      address: heap.Ref<object>,
      uuid: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.connect(socketId, A.H.get<object>(address), A.H.get<object>(uuid));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Create": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket && "create" in WEBEXT?.bluetoothSocket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Create": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.create);
    },
    "call_Create": (retPtr: Pointer, properties: Pointer): void => {
      const properties_ffi = {};

      if (A.load.Bool(properties + 12)) {
        properties_ffi["persistent"] = A.load.Bool(properties + 0);
      }
      properties_ffi["name"] = A.load.Ref(properties + 4, undefined);
      if (A.load.Bool(properties + 13)) {
        properties_ffi["bufferSize"] = A.load.Int32(properties + 8);
      }

      const _ret = WEBEXT.bluetoothSocket.create(properties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Create": (retPtr: Pointer, errPtr: Pointer, properties: Pointer): heap.Ref<boolean> => {
      try {
        const properties_ffi = {};

        if (A.load.Bool(properties + 12)) {
          properties_ffi["persistent"] = A.load.Bool(properties + 0);
        }
        properties_ffi["name"] = A.load.Ref(properties + 4, undefined);
        if (A.load.Bool(properties + 13)) {
          properties_ffi["bufferSize"] = A.load.Int32(properties + 8);
        }

        const _ret = WEBEXT.bluetoothSocket.create(properties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Disconnect": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket && "disconnect" in WEBEXT?.bluetoothSocket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Disconnect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.disconnect);
    },
    "call_Disconnect": (retPtr: Pointer, socketId: number): void => {
      const _ret = WEBEXT.bluetoothSocket.disconnect(socketId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Disconnect": (retPtr: Pointer, errPtr: Pointer, socketId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.disconnect(socketId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket && "getInfo" in WEBEXT?.bluetoothSocket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.getInfo);
    },
    "call_GetInfo": (retPtr: Pointer, socketId: number): void => {
      const _ret = WEBEXT.bluetoothSocket.getInfo(socketId);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetInfo": (retPtr: Pointer, errPtr: Pointer, socketId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.getInfo(socketId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSockets": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket && "getSockets" in WEBEXT?.bluetoothSocket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSockets": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.getSockets);
    },
    "call_GetSockets": (retPtr: Pointer): void => {
      const _ret = WEBEXT.bluetoothSocket.getSockets();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSockets": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.getSockets();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ListenUsingL2cap": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket && "listenUsingL2cap" in WEBEXT?.bluetoothSocket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ListenUsingL2cap": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.listenUsingL2cap);
    },
    "call_ListenUsingL2cap": (retPtr: Pointer, socketId: number, uuid: heap.Ref<object>, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 12)) {
        options_ffi["channel"] = A.load.Int32(options + 0);
      }
      if (A.load.Bool(options + 13)) {
        options_ffi["psm"] = A.load.Int32(options + 4);
      }
      if (A.load.Bool(options + 14)) {
        options_ffi["backlog"] = A.load.Int32(options + 8);
      }

      const _ret = WEBEXT.bluetoothSocket.listenUsingL2cap(socketId, A.H.get<object>(uuid), options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ListenUsingL2cap": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      uuid: heap.Ref<object>,
      options: Pointer
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 12)) {
          options_ffi["channel"] = A.load.Int32(options + 0);
        }
        if (A.load.Bool(options + 13)) {
          options_ffi["psm"] = A.load.Int32(options + 4);
        }
        if (A.load.Bool(options + 14)) {
          options_ffi["backlog"] = A.load.Int32(options + 8);
        }

        const _ret = WEBEXT.bluetoothSocket.listenUsingL2cap(socketId, A.H.get<object>(uuid), options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ListenUsingRfcomm": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket && "listenUsingRfcomm" in WEBEXT?.bluetoothSocket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ListenUsingRfcomm": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.listenUsingRfcomm);
    },
    "call_ListenUsingRfcomm": (retPtr: Pointer, socketId: number, uuid: heap.Ref<object>, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 12)) {
        options_ffi["channel"] = A.load.Int32(options + 0);
      }
      if (A.load.Bool(options + 13)) {
        options_ffi["psm"] = A.load.Int32(options + 4);
      }
      if (A.load.Bool(options + 14)) {
        options_ffi["backlog"] = A.load.Int32(options + 8);
      }

      const _ret = WEBEXT.bluetoothSocket.listenUsingRfcomm(socketId, A.H.get<object>(uuid), options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ListenUsingRfcomm": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      uuid: heap.Ref<object>,
      options: Pointer
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 12)) {
          options_ffi["channel"] = A.load.Int32(options + 0);
        }
        if (A.load.Bool(options + 13)) {
          options_ffi["psm"] = A.load.Int32(options + 4);
        }
        if (A.load.Bool(options + 14)) {
          options_ffi["backlog"] = A.load.Int32(options + 8);
        }

        const _ret = WEBEXT.bluetoothSocket.listenUsingRfcomm(socketId, A.H.get<object>(uuid), options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAccept": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket?.onAccept && "addListener" in WEBEXT?.bluetoothSocket?.onAccept) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAccept": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.onAccept.addListener);
    },
    "call_OnAccept": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothSocket.onAccept.addListener(A.H.get<object>(callback));
    },
    "try_OnAccept": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.onAccept.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAccept": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket?.onAccept && "removeListener" in WEBEXT?.bluetoothSocket?.onAccept) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAccept": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.onAccept.removeListener);
    },
    "call_OffAccept": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothSocket.onAccept.removeListener(A.H.get<object>(callback));
    },
    "try_OffAccept": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.onAccept.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAccept": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket?.onAccept && "hasListener" in WEBEXT?.bluetoothSocket?.onAccept) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAccept": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.onAccept.hasListener);
    },
    "call_HasOnAccept": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothSocket.onAccept.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAccept": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.onAccept.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAcceptError": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket?.onAcceptError && "addListener" in WEBEXT?.bluetoothSocket?.onAcceptError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAcceptError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.onAcceptError.addListener);
    },
    "call_OnAcceptError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothSocket.onAcceptError.addListener(A.H.get<object>(callback));
    },
    "try_OnAcceptError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.onAcceptError.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAcceptError": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket?.onAcceptError && "removeListener" in WEBEXT?.bluetoothSocket?.onAcceptError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAcceptError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.onAcceptError.removeListener);
    },
    "call_OffAcceptError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothSocket.onAcceptError.removeListener(A.H.get<object>(callback));
    },
    "try_OffAcceptError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.onAcceptError.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAcceptError": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket?.onAcceptError && "hasListener" in WEBEXT?.bluetoothSocket?.onAcceptError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAcceptError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.onAcceptError.hasListener);
    },
    "call_HasOnAcceptError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothSocket.onAcceptError.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAcceptError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.onAcceptError.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnReceive": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket?.onReceive && "addListener" in WEBEXT?.bluetoothSocket?.onReceive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnReceive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.onReceive.addListener);
    },
    "call_OnReceive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothSocket.onReceive.addListener(A.H.get<object>(callback));
    },
    "try_OnReceive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.onReceive.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffReceive": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket?.onReceive && "removeListener" in WEBEXT?.bluetoothSocket?.onReceive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffReceive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.onReceive.removeListener);
    },
    "call_OffReceive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothSocket.onReceive.removeListener(A.H.get<object>(callback));
    },
    "try_OffReceive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.onReceive.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnReceive": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket?.onReceive && "hasListener" in WEBEXT?.bluetoothSocket?.onReceive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnReceive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.onReceive.hasListener);
    },
    "call_HasOnReceive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothSocket.onReceive.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnReceive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.onReceive.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnReceiveError": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket?.onReceiveError && "addListener" in WEBEXT?.bluetoothSocket?.onReceiveError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnReceiveError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.onReceiveError.addListener);
    },
    "call_OnReceiveError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothSocket.onReceiveError.addListener(A.H.get<object>(callback));
    },
    "try_OnReceiveError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.onReceiveError.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffReceiveError": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket?.onReceiveError && "removeListener" in WEBEXT?.bluetoothSocket?.onReceiveError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffReceiveError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.onReceiveError.removeListener);
    },
    "call_OffReceiveError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothSocket.onReceiveError.removeListener(A.H.get<object>(callback));
    },
    "try_OffReceiveError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.onReceiveError.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnReceiveError": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket?.onReceiveError && "hasListener" in WEBEXT?.bluetoothSocket?.onReceiveError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnReceiveError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.onReceiveError.hasListener);
    },
    "call_HasOnReceiveError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothSocket.onReceiveError.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnReceiveError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.onReceiveError.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Send": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket && "send" in WEBEXT?.bluetoothSocket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Send": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.send);
    },
    "call_Send": (retPtr: Pointer, socketId: number, data: heap.Ref<object>): void => {
      const _ret = WEBEXT.bluetoothSocket.send(socketId, A.H.get<object>(data));
      A.store.Ref(retPtr, _ret);
    },
    "try_Send": (retPtr: Pointer, errPtr: Pointer, socketId: number, data: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.send(socketId, A.H.get<object>(data));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPaused": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket && "setPaused" in WEBEXT?.bluetoothSocket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPaused": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.setPaused);
    },
    "call_SetPaused": (retPtr: Pointer, socketId: number, paused: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.bluetoothSocket.setPaused(socketId, paused === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetPaused": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      paused: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bluetoothSocket.setPaused(socketId, paused === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Update": (): heap.Ref<boolean> => {
      if (WEBEXT?.bluetoothSocket && "update" in WEBEXT?.bluetoothSocket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Update": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bluetoothSocket.update);
    },
    "call_Update": (retPtr: Pointer, socketId: number, properties: Pointer): void => {
      const properties_ffi = {};

      if (A.load.Bool(properties + 12)) {
        properties_ffi["persistent"] = A.load.Bool(properties + 0);
      }
      properties_ffi["name"] = A.load.Ref(properties + 4, undefined);
      if (A.load.Bool(properties + 13)) {
        properties_ffi["bufferSize"] = A.load.Int32(properties + 8);
      }

      const _ret = WEBEXT.bluetoothSocket.update(socketId, properties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Update": (retPtr: Pointer, errPtr: Pointer, socketId: number, properties: Pointer): heap.Ref<boolean> => {
      try {
        const properties_ffi = {};

        if (A.load.Bool(properties + 12)) {
          properties_ffi["persistent"] = A.load.Bool(properties + 0);
        }
        properties_ffi["name"] = A.load.Ref(properties + 4, undefined);
        if (A.load.Bool(properties + 13)) {
          properties_ffi["bufferSize"] = A.load.Int32(properties + 8);
        }

        const _ret = WEBEXT.bluetoothSocket.update(socketId, properties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
