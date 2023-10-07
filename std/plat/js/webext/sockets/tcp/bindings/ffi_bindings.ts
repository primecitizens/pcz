import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/sockets/tcp", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
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
    "constof_DnsQueryType": (ref: heap.Ref<string>): number => {
      const idx = ["any", "ipv4", "ipv6"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SocketInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 43, false);
        A.store.Bool(ptr + 36, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 37, false);
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 38, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 39, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 40, false);
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 20, undefined);
        A.store.Bool(ptr + 41, false);
        A.store.Int32(ptr + 24, 0);
        A.store.Ref(ptr + 28, undefined);
        A.store.Bool(ptr + 42, false);
        A.store.Int32(ptr + 32, 0);
      } else {
        A.store.Bool(ptr + 43, true);
        A.store.Bool(ptr + 36, "socketId" in x ? true : false);
        A.store.Int32(ptr + 0, x["socketId"] === undefined ? 0 : (x["socketId"] as number));
        A.store.Bool(ptr + 37, "persistent" in x ? true : false);
        A.store.Bool(ptr + 4, x["persistent"] ? true : false);
        A.store.Ref(ptr + 8, x["name"]);
        A.store.Bool(ptr + 38, "bufferSize" in x ? true : false);
        A.store.Int32(ptr + 12, x["bufferSize"] === undefined ? 0 : (x["bufferSize"] as number));
        A.store.Bool(ptr + 39, "paused" in x ? true : false);
        A.store.Bool(ptr + 16, x["paused"] ? true : false);
        A.store.Bool(ptr + 40, "connected" in x ? true : false);
        A.store.Bool(ptr + 17, x["connected"] ? true : false);
        A.store.Ref(ptr + 20, x["localAddress"]);
        A.store.Bool(ptr + 41, "localPort" in x ? true : false);
        A.store.Int32(ptr + 24, x["localPort"] === undefined ? 0 : (x["localPort"] as number));
        A.store.Ref(ptr + 28, x["peerAddress"]);
        A.store.Bool(ptr + 42, "peerPort" in x ? true : false);
        A.store.Int32(ptr + 32, x["peerPort"] === undefined ? 0 : (x["peerPort"] as number));
      }
    },
    "load_SocketInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 36)) {
        x["socketId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["socketId"];
      }
      if (A.load.Bool(ptr + 37)) {
        x["persistent"] = A.load.Bool(ptr + 4);
      } else {
        delete x["persistent"];
      }
      x["name"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 38)) {
        x["bufferSize"] = A.load.Int32(ptr + 12);
      } else {
        delete x["bufferSize"];
      }
      if (A.load.Bool(ptr + 39)) {
        x["paused"] = A.load.Bool(ptr + 16);
      } else {
        delete x["paused"];
      }
      if (A.load.Bool(ptr + 40)) {
        x["connected"] = A.load.Bool(ptr + 17);
      } else {
        delete x["connected"];
      }
      x["localAddress"] = A.load.Ref(ptr + 20, undefined);
      if (A.load.Bool(ptr + 41)) {
        x["localPort"] = A.load.Int32(ptr + 24);
      } else {
        delete x["localPort"];
      }
      x["peerAddress"] = A.load.Ref(ptr + 28, undefined);
      if (A.load.Bool(ptr + 42)) {
        x["peerPort"] = A.load.Int32(ptr + 32);
      } else {
        delete x["peerPort"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ReceiveErrorInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Bool(ptr + 9, "resultCode" in x ? true : false);
        A.store.Int32(ptr + 4, x["resultCode"] === undefined ? 0 : (x["resultCode"] as number));
      }
    },
    "load_ReceiveErrorInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["socketId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["socketId"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["resultCode"] = A.load.Int32(ptr + 4);
      } else {
        delete x["resultCode"];
      }
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

    "store_TLSVersionConstraints": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["min"]);
        A.store.Ref(ptr + 4, x["max"]);
      }
    },
    "load_TLSVersionConstraints": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["min"] = A.load.Ref(ptr + 0, undefined);
      x["max"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SecureOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);

        A.store.Bool(ptr + 0 + 8, false);
        A.store.Ref(ptr + 0 + 0, undefined);
        A.store.Ref(ptr + 0 + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);

        if (typeof x["tlsVersion"] === "undefined") {
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Ref(ptr + 0 + 0, undefined);
          A.store.Ref(ptr + 0 + 4, undefined);
        } else {
          A.store.Bool(ptr + 0 + 8, true);
          A.store.Ref(ptr + 0 + 0, x["tlsVersion"]["min"]);
          A.store.Ref(ptr + 0 + 4, x["tlsVersion"]["max"]);
        }
      }
    },
    "load_SecureOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 8)) {
        x["tlsVersion"] = {};
        x["tlsVersion"]["min"] = A.load.Ref(ptr + 0 + 0, undefined);
        x["tlsVersion"]["max"] = A.load.Ref(ptr + 0 + 4, undefined);
      } else {
        delete x["tlsVersion"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SendInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "resultCode" in x ? true : false);
        A.store.Int32(ptr + 0, x["resultCode"] === undefined ? 0 : (x["resultCode"] as number));
        A.store.Bool(ptr + 9, "bytesSent" in x ? true : false);
        A.store.Int32(ptr + 4, x["bytesSent"] === undefined ? 0 : (x["bytesSent"] as number));
      }
    },
    "load_SendInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["resultCode"] = A.load.Int32(ptr + 0);
      } else {
        delete x["resultCode"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["bytesSent"] = A.load.Int32(ptr + 4);
      } else {
        delete x["bytesSent"];
      }
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
      if (WEBEXT?.sockets?.tcp && "close" in WEBEXT?.sockets?.tcp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Close": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.close);
    },
    "call_Close": (retPtr: Pointer, socketId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcp.close(socketId, A.H.get<object>(callback));
    },
    "try_Close": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcp.close(socketId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Connect": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp && "connect" in WEBEXT?.sockets?.tcp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Connect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.connect);
    },
    "call_Connect": (
      retPtr: Pointer,
      socketId: number,
      peerAddress: heap.Ref<object>,
      peerPort: number,
      dnsQueryType: number,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.sockets.tcp.connect(
        socketId,
        A.H.get<object>(peerAddress),
        peerPort,
        dnsQueryType > 0 && dnsQueryType <= 3 ? ["any", "ipv4", "ipv6"][dnsQueryType - 1] : undefined,
        A.H.get<object>(callback)
      );
    },
    "try_Connect": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      peerAddress: heap.Ref<object>,
      peerPort: number,
      dnsQueryType: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcp.connect(
          socketId,
          A.H.get<object>(peerAddress),
          peerPort,
          dnsQueryType > 0 && dnsQueryType <= 3 ? ["any", "ipv4", "ipv6"][dnsQueryType - 1] : undefined,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Create": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp && "create" in WEBEXT?.sockets?.tcp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Create": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.create);
    },
    "call_Create": (retPtr: Pointer, properties: Pointer, callback: heap.Ref<object>): void => {
      const properties_ffi = {};

      if (A.load.Bool(properties + 12)) {
        properties_ffi["persistent"] = A.load.Bool(properties + 0);
      }
      properties_ffi["name"] = A.load.Ref(properties + 4, undefined);
      if (A.load.Bool(properties + 13)) {
        properties_ffi["bufferSize"] = A.load.Int32(properties + 8);
      }

      const _ret = WEBEXT.sockets.tcp.create(properties_ffi, A.H.get<object>(callback));
    },
    "try_Create": (
      retPtr: Pointer,
      errPtr: Pointer,
      properties: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const properties_ffi = {};

        if (A.load.Bool(properties + 12)) {
          properties_ffi["persistent"] = A.load.Bool(properties + 0);
        }
        properties_ffi["name"] = A.load.Ref(properties + 4, undefined);
        if (A.load.Bool(properties + 13)) {
          properties_ffi["bufferSize"] = A.load.Int32(properties + 8);
        }

        const _ret = WEBEXT.sockets.tcp.create(properties_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Disconnect": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp && "disconnect" in WEBEXT?.sockets?.tcp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Disconnect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.disconnect);
    },
    "call_Disconnect": (retPtr: Pointer, socketId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcp.disconnect(socketId, A.H.get<object>(callback));
    },
    "try_Disconnect": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcp.disconnect(socketId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp && "getInfo" in WEBEXT?.sockets?.tcp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.getInfo);
    },
    "call_GetInfo": (retPtr: Pointer, socketId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcp.getInfo(socketId, A.H.get<object>(callback));
    },
    "try_GetInfo": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcp.getInfo(socketId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSockets": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp && "getSockets" in WEBEXT?.sockets?.tcp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSockets": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.getSockets);
    },
    "call_GetSockets": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcp.getSockets(A.H.get<object>(callback));
    },
    "try_GetSockets": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcp.getSockets(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnReceive": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp?.onReceive && "addListener" in WEBEXT?.sockets?.tcp?.onReceive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnReceive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.onReceive.addListener);
    },
    "call_OnReceive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcp.onReceive.addListener(A.H.get<object>(callback));
    },
    "try_OnReceive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcp.onReceive.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffReceive": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp?.onReceive && "removeListener" in WEBEXT?.sockets?.tcp?.onReceive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffReceive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.onReceive.removeListener);
    },
    "call_OffReceive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcp.onReceive.removeListener(A.H.get<object>(callback));
    },
    "try_OffReceive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcp.onReceive.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnReceive": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp?.onReceive && "hasListener" in WEBEXT?.sockets?.tcp?.onReceive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnReceive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.onReceive.hasListener);
    },
    "call_HasOnReceive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcp.onReceive.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnReceive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcp.onReceive.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnReceiveError": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp?.onReceiveError && "addListener" in WEBEXT?.sockets?.tcp?.onReceiveError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnReceiveError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.onReceiveError.addListener);
    },
    "call_OnReceiveError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcp.onReceiveError.addListener(A.H.get<object>(callback));
    },
    "try_OnReceiveError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcp.onReceiveError.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffReceiveError": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp?.onReceiveError && "removeListener" in WEBEXT?.sockets?.tcp?.onReceiveError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffReceiveError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.onReceiveError.removeListener);
    },
    "call_OffReceiveError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcp.onReceiveError.removeListener(A.H.get<object>(callback));
    },
    "try_OffReceiveError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcp.onReceiveError.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnReceiveError": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp?.onReceiveError && "hasListener" in WEBEXT?.sockets?.tcp?.onReceiveError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnReceiveError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.onReceiveError.hasListener);
    },
    "call_HasOnReceiveError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcp.onReceiveError.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnReceiveError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcp.onReceiveError.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Secure": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp && "secure" in WEBEXT?.sockets?.tcp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Secure": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.secure);
    },
    "call_Secure": (retPtr: Pointer, socketId: number, options: Pointer, callback: heap.Ref<object>): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 0 + 8)) {
        options_ffi["tlsVersion"] = {};
        options_ffi["tlsVersion"]["min"] = A.load.Ref(options + 0 + 0, undefined);
        options_ffi["tlsVersion"]["max"] = A.load.Ref(options + 0 + 4, undefined);
      }

      const _ret = WEBEXT.sockets.tcp.secure(socketId, options_ffi, A.H.get<object>(callback));
    },
    "try_Secure": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      options: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 0 + 8)) {
          options_ffi["tlsVersion"] = {};
          options_ffi["tlsVersion"]["min"] = A.load.Ref(options + 0 + 0, undefined);
          options_ffi["tlsVersion"]["max"] = A.load.Ref(options + 0 + 4, undefined);
        }

        const _ret = WEBEXT.sockets.tcp.secure(socketId, options_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Send": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp && "send" in WEBEXT?.sockets?.tcp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Send": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.send);
    },
    "call_Send": (retPtr: Pointer, socketId: number, data: heap.Ref<object>, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcp.send(socketId, A.H.get<object>(data), A.H.get<object>(callback));
    },
    "try_Send": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      data: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcp.send(socketId, A.H.get<object>(data), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetKeepAlive": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp && "setKeepAlive" in WEBEXT?.sockets?.tcp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetKeepAlive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.setKeepAlive);
    },
    "call_SetKeepAlive": (
      retPtr: Pointer,
      socketId: number,
      enable: heap.Ref<boolean>,
      delay: number,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.sockets.tcp.setKeepAlive(socketId, enable === A.H.TRUE, delay, A.H.get<object>(callback));
    },
    "try_SetKeepAlive": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      enable: heap.Ref<boolean>,
      delay: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcp.setKeepAlive(socketId, enable === A.H.TRUE, delay, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetNoDelay": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp && "setNoDelay" in WEBEXT?.sockets?.tcp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetNoDelay": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.setNoDelay);
    },
    "call_SetNoDelay": (
      retPtr: Pointer,
      socketId: number,
      noDelay: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.sockets.tcp.setNoDelay(socketId, noDelay === A.H.TRUE, A.H.get<object>(callback));
    },
    "try_SetNoDelay": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      noDelay: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcp.setNoDelay(socketId, noDelay === A.H.TRUE, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPaused": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp && "setPaused" in WEBEXT?.sockets?.tcp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPaused": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.setPaused);
    },
    "call_SetPaused": (
      retPtr: Pointer,
      socketId: number,
      paused: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.sockets.tcp.setPaused(socketId, paused === A.H.TRUE, A.H.get<object>(callback));
    },
    "try_SetPaused": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      paused: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcp.setPaused(socketId, paused === A.H.TRUE, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Update": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcp && "update" in WEBEXT?.sockets?.tcp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Update": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcp.update);
    },
    "call_Update": (retPtr: Pointer, socketId: number, properties: Pointer, callback: heap.Ref<object>): void => {
      const properties_ffi = {};

      if (A.load.Bool(properties + 12)) {
        properties_ffi["persistent"] = A.load.Bool(properties + 0);
      }
      properties_ffi["name"] = A.load.Ref(properties + 4, undefined);
      if (A.load.Bool(properties + 13)) {
        properties_ffi["bufferSize"] = A.load.Int32(properties + 8);
      }

      const _ret = WEBEXT.sockets.tcp.update(socketId, properties_ffi, A.H.get<object>(callback));
    },
    "try_Update": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      properties: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const properties_ffi = {};

        if (A.load.Bool(properties + 12)) {
          properties_ffi["persistent"] = A.load.Bool(properties + 0);
        }
        properties_ffi["name"] = A.load.Ref(properties + 4, undefined);
        if (A.load.Bool(properties + 13)) {
          properties_ffi["bufferSize"] = A.load.Int32(properties + 8);
        }

        const _ret = WEBEXT.sockets.tcp.update(socketId, properties_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
