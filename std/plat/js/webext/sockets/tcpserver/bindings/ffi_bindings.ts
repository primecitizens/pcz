import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/sockets/tcpserver", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AcceptErrorInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_AcceptErrorInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
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
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 24, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 27, false);
        A.store.Int32(ptr + 20, 0);
      } else {
        A.store.Bool(ptr + 28, true);
        A.store.Bool(ptr + 24, "socketId" in x ? true : false);
        A.store.Int32(ptr + 0, x["socketId"] === undefined ? 0 : (x["socketId"] as number));
        A.store.Bool(ptr + 25, "persistent" in x ? true : false);
        A.store.Bool(ptr + 4, x["persistent"] ? true : false);
        A.store.Ref(ptr + 8, x["name"]);
        A.store.Bool(ptr + 26, "paused" in x ? true : false);
        A.store.Bool(ptr + 12, x["paused"] ? true : false);
        A.store.Ref(ptr + 16, x["localAddress"]);
        A.store.Bool(ptr + 27, "localPort" in x ? true : false);
        A.store.Int32(ptr + 20, x["localPort"] === undefined ? 0 : (x["localPort"] as number));
      }
    },
    "load_SocketInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["socketId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["socketId"];
      }
      if (A.load.Bool(ptr + 25)) {
        x["persistent"] = A.load.Bool(ptr + 4);
      } else {
        delete x["persistent"];
      }
      x["name"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 26)) {
        x["paused"] = A.load.Bool(ptr + 12);
      } else {
        delete x["paused"];
      }
      x["localAddress"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 27)) {
        x["localPort"] = A.load.Int32(ptr + 20);
      } else {
        delete x["localPort"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SocketProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "persistent" in x ? true : false);
        A.store.Bool(ptr + 0, x["persistent"] ? true : false);
        A.store.Ref(ptr + 4, x["name"]);
      }
    },
    "load_SocketProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["persistent"] = A.load.Bool(ptr + 0);
      } else {
        delete x["persistent"];
      }
      x["name"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Close": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcpServer && "close" in WEBEXT?.sockets?.tcpServer) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Close": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcpServer.close);
    },
    "call_Close": (retPtr: Pointer, socketId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcpServer.close(socketId, A.H.get<object>(callback));
    },
    "try_Close": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcpServer.close(socketId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Create": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcpServer && "create" in WEBEXT?.sockets?.tcpServer) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Create": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcpServer.create);
    },
    "call_Create": (retPtr: Pointer, properties: Pointer, callback: heap.Ref<object>): void => {
      const properties_ffi = {};

      if (A.load.Bool(properties + 8)) {
        properties_ffi["persistent"] = A.load.Bool(properties + 0);
      }
      properties_ffi["name"] = A.load.Ref(properties + 4, undefined);

      const _ret = WEBEXT.sockets.tcpServer.create(properties_ffi, A.H.get<object>(callback));
    },
    "try_Create": (
      retPtr: Pointer,
      errPtr: Pointer,
      properties: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const properties_ffi = {};

        if (A.load.Bool(properties + 8)) {
          properties_ffi["persistent"] = A.load.Bool(properties + 0);
        }
        properties_ffi["name"] = A.load.Ref(properties + 4, undefined);

        const _ret = WEBEXT.sockets.tcpServer.create(properties_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Disconnect": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcpServer && "disconnect" in WEBEXT?.sockets?.tcpServer) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Disconnect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcpServer.disconnect);
    },
    "call_Disconnect": (retPtr: Pointer, socketId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcpServer.disconnect(socketId, A.H.get<object>(callback));
    },
    "try_Disconnect": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcpServer.disconnect(socketId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcpServer && "getInfo" in WEBEXT?.sockets?.tcpServer) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcpServer.getInfo);
    },
    "call_GetInfo": (retPtr: Pointer, socketId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcpServer.getInfo(socketId, A.H.get<object>(callback));
    },
    "try_GetInfo": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcpServer.getInfo(socketId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSockets": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcpServer && "getSockets" in WEBEXT?.sockets?.tcpServer) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSockets": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcpServer.getSockets);
    },
    "call_GetSockets": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcpServer.getSockets(A.H.get<object>(callback));
    },
    "try_GetSockets": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcpServer.getSockets(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Listen": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcpServer && "listen" in WEBEXT?.sockets?.tcpServer) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Listen": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcpServer.listen);
    },
    "call_Listen": (
      retPtr: Pointer,
      socketId: number,
      address: heap.Ref<object>,
      port: number,
      backlog: number,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.sockets.tcpServer.listen(
        socketId,
        A.H.get<object>(address),
        port,
        backlog,
        A.H.get<object>(callback)
      );
    },
    "try_Listen": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      address: heap.Ref<object>,
      port: number,
      backlog: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcpServer.listen(
          socketId,
          A.H.get<object>(address),
          port,
          backlog,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAccept": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcpServer?.onAccept && "addListener" in WEBEXT?.sockets?.tcpServer?.onAccept) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAccept": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcpServer.onAccept.addListener);
    },
    "call_OnAccept": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcpServer.onAccept.addListener(A.H.get<object>(callback));
    },
    "try_OnAccept": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcpServer.onAccept.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAccept": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcpServer?.onAccept && "removeListener" in WEBEXT?.sockets?.tcpServer?.onAccept) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAccept": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcpServer.onAccept.removeListener);
    },
    "call_OffAccept": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcpServer.onAccept.removeListener(A.H.get<object>(callback));
    },
    "try_OffAccept": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcpServer.onAccept.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAccept": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcpServer?.onAccept && "hasListener" in WEBEXT?.sockets?.tcpServer?.onAccept) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAccept": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcpServer.onAccept.hasListener);
    },
    "call_HasOnAccept": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcpServer.onAccept.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAccept": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcpServer.onAccept.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAcceptError": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcpServer?.onAcceptError && "addListener" in WEBEXT?.sockets?.tcpServer?.onAcceptError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAcceptError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcpServer.onAcceptError.addListener);
    },
    "call_OnAcceptError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcpServer.onAcceptError.addListener(A.H.get<object>(callback));
    },
    "try_OnAcceptError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcpServer.onAcceptError.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAcceptError": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcpServer?.onAcceptError && "removeListener" in WEBEXT?.sockets?.tcpServer?.onAcceptError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAcceptError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcpServer.onAcceptError.removeListener);
    },
    "call_OffAcceptError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcpServer.onAcceptError.removeListener(A.H.get<object>(callback));
    },
    "try_OffAcceptError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcpServer.onAcceptError.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAcceptError": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcpServer?.onAcceptError && "hasListener" in WEBEXT?.sockets?.tcpServer?.onAcceptError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAcceptError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcpServer.onAcceptError.hasListener);
    },
    "call_HasOnAcceptError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.tcpServer.onAcceptError.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAcceptError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcpServer.onAcceptError.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPaused": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcpServer && "setPaused" in WEBEXT?.sockets?.tcpServer) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPaused": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcpServer.setPaused);
    },
    "call_SetPaused": (
      retPtr: Pointer,
      socketId: number,
      paused: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.sockets.tcpServer.setPaused(socketId, paused === A.H.TRUE, A.H.get<object>(callback));
    },
    "try_SetPaused": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      paused: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.tcpServer.setPaused(socketId, paused === A.H.TRUE, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Update": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.tcpServer && "update" in WEBEXT?.sockets?.tcpServer) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Update": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.tcpServer.update);
    },
    "call_Update": (retPtr: Pointer, socketId: number, properties: Pointer, callback: heap.Ref<object>): void => {
      const properties_ffi = {};

      if (A.load.Bool(properties + 8)) {
        properties_ffi["persistent"] = A.load.Bool(properties + 0);
      }
      properties_ffi["name"] = A.load.Ref(properties + 4, undefined);

      const _ret = WEBEXT.sockets.tcpServer.update(socketId, properties_ffi, A.H.get<object>(callback));
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

        if (A.load.Bool(properties + 8)) {
          properties_ffi["persistent"] = A.load.Bool(properties + 0);
        }
        properties_ffi["name"] = A.load.Ref(properties + 4, undefined);

        const _ret = WEBEXT.sockets.tcpServer.update(socketId, properties_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
