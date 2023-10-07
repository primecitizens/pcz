import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/socket", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
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
        A.store.Bool(ptr + 8, "resultCode" in x ? true : false);
        A.store.Int32(ptr + 0, x["resultCode"] === undefined ? 0 : (x["resultCode"] as number));
        A.store.Bool(ptr + 9, "socketId" in x ? true : false);
        A.store.Int32(ptr + 4, x["socketId"] === undefined ? 0 : (x["socketId"] as number));
      }
    },
    "load_AcceptInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["resultCode"] = A.load.Int32(ptr + 0);
      } else {
        delete x["resultCode"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["socketId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["socketId"];
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

    "store_CreateOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_CreateOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_SocketType": (ref: heap.Ref<string>): number => {
      const idx = ["tcp", "udp"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SocketInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 27, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 25, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 26, false);
        A.store.Int32(ptr + 20, 0);
      } else {
        A.store.Bool(ptr + 27, true);
        A.store.Enum(ptr + 0, ["tcp", "udp"].indexOf(x["socketType"] as string));
        A.store.Bool(ptr + 24, "connected" in x ? true : false);
        A.store.Bool(ptr + 4, x["connected"] ? true : false);
        A.store.Ref(ptr + 8, x["peerAddress"]);
        A.store.Bool(ptr + 25, "peerPort" in x ? true : false);
        A.store.Int32(ptr + 12, x["peerPort"] === undefined ? 0 : (x["peerPort"] as number));
        A.store.Ref(ptr + 16, x["localAddress"]);
        A.store.Bool(ptr + 26, "localPort" in x ? true : false);
        A.store.Int32(ptr + 20, x["localPort"] === undefined ? 0 : (x["localPort"] as number));
      }
    },
    "load_SocketInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["socketType"] = A.load.Enum(ptr + 0, ["tcp", "udp"]);
      if (A.load.Bool(ptr + 24)) {
        x["connected"] = A.load.Bool(ptr + 4);
      } else {
        delete x["connected"];
      }
      x["peerAddress"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 25)) {
        x["peerPort"] = A.load.Int32(ptr + 12);
      } else {
        delete x["peerPort"];
      }
      x["localAddress"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 26)) {
        x["localPort"] = A.load.Int32(ptr + 20);
      } else {
        delete x["localPort"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_NetworkInterface": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["address"]);
        A.store.Bool(ptr + 12, "prefixLength" in x ? true : false);
        A.store.Int32(ptr + 8, x["prefixLength"] === undefined ? 0 : (x["prefixLength"] as number));
      }
    },
    "load_NetworkInterface": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["address"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["prefixLength"] = A.load.Int32(ptr + 8);
      } else {
        delete x["prefixLength"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ReadInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "resultCode" in x ? true : false);
        A.store.Int32(ptr + 0, x["resultCode"] === undefined ? 0 : (x["resultCode"] as number));
        A.store.Ref(ptr + 4, x["data"]);
      }
    },
    "load_ReadInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["resultCode"] = A.load.Int32(ptr + 0);
      } else {
        delete x["resultCode"];
      }
      x["data"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_RecvFromInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 12, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Bool(ptr + 16, "resultCode" in x ? true : false);
        A.store.Int32(ptr + 0, x["resultCode"] === undefined ? 0 : (x["resultCode"] as number));
        A.store.Ref(ptr + 4, x["data"]);
        A.store.Ref(ptr + 8, x["address"]);
        A.store.Bool(ptr + 17, "port" in x ? true : false);
        A.store.Int32(ptr + 12, x["port"] === undefined ? 0 : (x["port"] as number));
      }
    },
    "load_RecvFromInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["resultCode"] = A.load.Int32(ptr + 0);
      } else {
        delete x["resultCode"];
      }
      x["data"] = A.load.Ref(ptr + 4, undefined);
      x["address"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 17)) {
        x["port"] = A.load.Int32(ptr + 12);
      } else {
        delete x["port"];
      }
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

    "store_WriteInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "bytesWritten" in x ? true : false);
        A.store.Int32(ptr + 0, x["bytesWritten"] === undefined ? 0 : (x["bytesWritten"] as number));
      }
    },
    "load_WriteInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["bytesWritten"] = A.load.Int32(ptr + 0);
      } else {
        delete x["bytesWritten"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Accept": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "accept" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Accept": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.accept);
    },
    "call_Accept": (retPtr: Pointer, socketId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.socket.accept(socketId, A.H.get<object>(callback));
    },
    "try_Accept": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.accept(socketId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Bind": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "bind" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Bind": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.bind);
    },
    "call_Bind": (
      retPtr: Pointer,
      socketId: number,
      address: heap.Ref<object>,
      port: number,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.socket.bind(socketId, A.H.get<object>(address), port, A.H.get<object>(callback));
    },
    "try_Bind": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      address: heap.Ref<object>,
      port: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.bind(socketId, A.H.get<object>(address), port, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Connect": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "connect" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Connect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.connect);
    },
    "call_Connect": (
      retPtr: Pointer,
      socketId: number,
      hostname: heap.Ref<object>,
      port: number,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.socket.connect(socketId, A.H.get<object>(hostname), port, A.H.get<object>(callback));
    },
    "try_Connect": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      hostname: heap.Ref<object>,
      port: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.connect(socketId, A.H.get<object>(hostname), port, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Create": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "create" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Create": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.create);
    },
    "call_Create": (retPtr: Pointer, type: number, options: Pointer, callback: heap.Ref<object>): void => {
      const options_ffi = {};

      const _ret = WEBEXT.socket.create(
        type > 0 && type <= 2 ? ["tcp", "udp"][type - 1] : undefined,
        options_ffi,
        A.H.get<object>(callback)
      );
    },
    "try_Create": (
      retPtr: Pointer,
      errPtr: Pointer,
      type: number,
      options: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        const _ret = WEBEXT.socket.create(
          type > 0 && type <= 2 ? ["tcp", "udp"][type - 1] : undefined,
          options_ffi,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Destroy": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "destroy" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Destroy": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.destroy);
    },
    "call_Destroy": (retPtr: Pointer, socketId: number): void => {
      const _ret = WEBEXT.socket.destroy(socketId);
    },
    "try_Destroy": (retPtr: Pointer, errPtr: Pointer, socketId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.destroy(socketId);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Disconnect": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "disconnect" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Disconnect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.disconnect);
    },
    "call_Disconnect": (retPtr: Pointer, socketId: number): void => {
      const _ret = WEBEXT.socket.disconnect(socketId);
    },
    "try_Disconnect": (retPtr: Pointer, errPtr: Pointer, socketId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.disconnect(socketId);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "getInfo" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.getInfo);
    },
    "call_GetInfo": (retPtr: Pointer, socketId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.socket.getInfo(socketId, A.H.get<object>(callback));
    },
    "try_GetInfo": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.getInfo(socketId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetJoinedGroups": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "getJoinedGroups" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetJoinedGroups": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.getJoinedGroups);
    },
    "call_GetJoinedGroups": (retPtr: Pointer, socketId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.socket.getJoinedGroups(socketId, A.H.get<object>(callback));
    },
    "try_GetJoinedGroups": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.getJoinedGroups(socketId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetNetworkList": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "getNetworkList" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetNetworkList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.getNetworkList);
    },
    "call_GetNetworkList": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.socket.getNetworkList(A.H.get<object>(callback));
    },
    "try_GetNetworkList": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.getNetworkList(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_JoinGroup": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "joinGroup" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_JoinGroup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.joinGroup);
    },
    "call_JoinGroup": (
      retPtr: Pointer,
      socketId: number,
      address: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.socket.joinGroup(socketId, A.H.get<object>(address), A.H.get<object>(callback));
    },
    "try_JoinGroup": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      address: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.joinGroup(socketId, A.H.get<object>(address), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LeaveGroup": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "leaveGroup" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LeaveGroup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.leaveGroup);
    },
    "call_LeaveGroup": (
      retPtr: Pointer,
      socketId: number,
      address: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.socket.leaveGroup(socketId, A.H.get<object>(address), A.H.get<object>(callback));
    },
    "try_LeaveGroup": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      address: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.leaveGroup(socketId, A.H.get<object>(address), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Listen": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "listen" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Listen": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.listen);
    },
    "call_Listen": (
      retPtr: Pointer,
      socketId: number,
      address: heap.Ref<object>,
      port: number,
      backlog: number,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.socket.listen(socketId, A.H.get<object>(address), port, backlog, A.H.get<object>(callback));
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
        const _ret = WEBEXT.socket.listen(socketId, A.H.get<object>(address), port, backlog, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Read": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "read" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Read": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.read);
    },
    "call_Read": (retPtr: Pointer, socketId: number, bufferSize: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.socket.read(socketId, bufferSize, A.H.get<object>(callback));
    },
    "try_Read": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      bufferSize: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.read(socketId, bufferSize, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecvFrom": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "recvFrom" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecvFrom": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.recvFrom);
    },
    "call_RecvFrom": (retPtr: Pointer, socketId: number, bufferSize: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.socket.recvFrom(socketId, bufferSize, A.H.get<object>(callback));
    },
    "try_RecvFrom": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      bufferSize: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.recvFrom(socketId, bufferSize, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Secure": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "secure" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Secure": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.secure);
    },
    "call_Secure": (retPtr: Pointer, socketId: number, options: Pointer, callback: heap.Ref<object>): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 0 + 8)) {
        options_ffi["tlsVersion"] = {};
        options_ffi["tlsVersion"]["min"] = A.load.Ref(options + 0 + 0, undefined);
        options_ffi["tlsVersion"]["max"] = A.load.Ref(options + 0 + 4, undefined);
      }

      const _ret = WEBEXT.socket.secure(socketId, options_ffi, A.H.get<object>(callback));
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

        const _ret = WEBEXT.socket.secure(socketId, options_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendTo": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "sendTo" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendTo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.sendTo);
    },
    "call_SendTo": (
      retPtr: Pointer,
      socketId: number,
      data: heap.Ref<object>,
      address: heap.Ref<object>,
      port: number,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.socket.sendTo(
        socketId,
        A.H.get<object>(data),
        A.H.get<object>(address),
        port,
        A.H.get<object>(callback)
      );
    },
    "try_SendTo": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      data: heap.Ref<object>,
      address: heap.Ref<object>,
      port: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.sendTo(
          socketId,
          A.H.get<object>(data),
          A.H.get<object>(address),
          port,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetKeepAlive": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "setKeepAlive" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetKeepAlive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.setKeepAlive);
    },
    "call_SetKeepAlive": (
      retPtr: Pointer,
      socketId: number,
      enable: heap.Ref<boolean>,
      delay: number,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.socket.setKeepAlive(socketId, enable === A.H.TRUE, delay, A.H.get<object>(callback));
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
        const _ret = WEBEXT.socket.setKeepAlive(socketId, enable === A.H.TRUE, delay, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetMulticastLoopbackMode": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "setMulticastLoopbackMode" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetMulticastLoopbackMode": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.setMulticastLoopbackMode);
    },
    "call_SetMulticastLoopbackMode": (
      retPtr: Pointer,
      socketId: number,
      enabled: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.socket.setMulticastLoopbackMode(socketId, enabled === A.H.TRUE, A.H.get<object>(callback));
    },
    "try_SetMulticastLoopbackMode": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      enabled: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.setMulticastLoopbackMode(socketId, enabled === A.H.TRUE, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetMulticastTimeToLive": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "setMulticastTimeToLive" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetMulticastTimeToLive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.setMulticastTimeToLive);
    },
    "call_SetMulticastTimeToLive": (
      retPtr: Pointer,
      socketId: number,
      ttl: number,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.socket.setMulticastTimeToLive(socketId, ttl, A.H.get<object>(callback));
    },
    "try_SetMulticastTimeToLive": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      ttl: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.setMulticastTimeToLive(socketId, ttl, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetNoDelay": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "setNoDelay" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetNoDelay": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.setNoDelay);
    },
    "call_SetNoDelay": (
      retPtr: Pointer,
      socketId: number,
      noDelay: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.socket.setNoDelay(socketId, noDelay === A.H.TRUE, A.H.get<object>(callback));
    },
    "try_SetNoDelay": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      noDelay: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.setNoDelay(socketId, noDelay === A.H.TRUE, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Write": (): heap.Ref<boolean> => {
      if (WEBEXT?.socket && "write" in WEBEXT?.socket) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Write": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.socket.write);
    },
    "call_Write": (retPtr: Pointer, socketId: number, data: heap.Ref<object>, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.socket.write(socketId, A.H.get<object>(data), A.H.get<object>(callback));
    },
    "try_Write": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      data: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.socket.write(socketId, A.H.get<object>(data), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
