import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/sockets/udp", (A: Application) => {
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
        A.store.Ref(ptr + 20, undefined);
        A.store.Bool(ptr + 32, false);
        A.store.Int32(ptr + 24, 0);
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
        A.store.Ref(ptr + 20, x["localAddress"]);
        A.store.Bool(ptr + 32, "localPort" in x ? true : false);
        A.store.Int32(ptr + 24, x["localPort"] === undefined ? 0 : (x["localPort"] as number));
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
      x["localAddress"] = A.load.Ref(ptr + 20, undefined);
      if (A.load.Bool(ptr + 32)) {
        x["localPort"] = A.load.Int32(ptr + 24);
      } else {
        delete x["localPort"];
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
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 12, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Bool(ptr + 16, "socketId" in x ? true : false);
        A.store.Int32(ptr + 0, x["socketId"] === undefined ? 0 : (x["socketId"] as number));
        A.store.Ref(ptr + 4, x["data"]);
        A.store.Ref(ptr + 8, x["remoteAddress"]);
        A.store.Bool(ptr + 17, "remotePort" in x ? true : false);
        A.store.Int32(ptr + 12, x["remotePort"] === undefined ? 0 : (x["remotePort"] as number));
      }
    },
    "load_ReceiveInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["socketId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["socketId"];
      }
      x["data"] = A.load.Ref(ptr + 4, undefined);
      x["remoteAddress"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 17)) {
        x["remotePort"] = A.load.Int32(ptr + 12);
      } else {
        delete x["remotePort"];
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
    "has_Bind": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp && "bind" in WEBEXT?.sockets?.udp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Bind": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.bind);
    },
    "call_Bind": (
      retPtr: Pointer,
      socketId: number,
      address: heap.Ref<object>,
      port: number,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.sockets.udp.bind(socketId, A.H.get<object>(address), port, A.H.get<object>(callback));
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
        const _ret = WEBEXT.sockets.udp.bind(socketId, A.H.get<object>(address), port, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Close": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp && "close" in WEBEXT?.sockets?.udp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Close": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.close);
    },
    "call_Close": (retPtr: Pointer, socketId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.udp.close(socketId, A.H.get<object>(callback));
    },
    "try_Close": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.close(socketId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Create": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp && "create" in WEBEXT?.sockets?.udp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Create": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.create);
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

      const _ret = WEBEXT.sockets.udp.create(properties_ffi, A.H.get<object>(callback));
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

        const _ret = WEBEXT.sockets.udp.create(properties_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp && "getInfo" in WEBEXT?.sockets?.udp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.getInfo);
    },
    "call_GetInfo": (retPtr: Pointer, socketId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.udp.getInfo(socketId, A.H.get<object>(callback));
    },
    "try_GetInfo": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.getInfo(socketId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetJoinedGroups": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp && "getJoinedGroups" in WEBEXT?.sockets?.udp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetJoinedGroups": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.getJoinedGroups);
    },
    "call_GetJoinedGroups": (retPtr: Pointer, socketId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.udp.getJoinedGroups(socketId, A.H.get<object>(callback));
    },
    "try_GetJoinedGroups": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.getJoinedGroups(socketId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSockets": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp && "getSockets" in WEBEXT?.sockets?.udp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSockets": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.getSockets);
    },
    "call_GetSockets": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.udp.getSockets(A.H.get<object>(callback));
    },
    "try_GetSockets": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.getSockets(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_JoinGroup": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp && "joinGroup" in WEBEXT?.sockets?.udp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_JoinGroup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.joinGroup);
    },
    "call_JoinGroup": (
      retPtr: Pointer,
      socketId: number,
      address: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.sockets.udp.joinGroup(socketId, A.H.get<object>(address), A.H.get<object>(callback));
    },
    "try_JoinGroup": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      address: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.joinGroup(socketId, A.H.get<object>(address), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LeaveGroup": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp && "leaveGroup" in WEBEXT?.sockets?.udp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LeaveGroup": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.leaveGroup);
    },
    "call_LeaveGroup": (
      retPtr: Pointer,
      socketId: number,
      address: heap.Ref<object>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.sockets.udp.leaveGroup(socketId, A.H.get<object>(address), A.H.get<object>(callback));
    },
    "try_LeaveGroup": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      address: heap.Ref<object>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.leaveGroup(socketId, A.H.get<object>(address), A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnReceive": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp?.onReceive && "addListener" in WEBEXT?.sockets?.udp?.onReceive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnReceive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.onReceive.addListener);
    },
    "call_OnReceive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.udp.onReceive.addListener(A.H.get<object>(callback));
    },
    "try_OnReceive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.onReceive.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffReceive": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp?.onReceive && "removeListener" in WEBEXT?.sockets?.udp?.onReceive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffReceive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.onReceive.removeListener);
    },
    "call_OffReceive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.udp.onReceive.removeListener(A.H.get<object>(callback));
    },
    "try_OffReceive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.onReceive.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnReceive": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp?.onReceive && "hasListener" in WEBEXT?.sockets?.udp?.onReceive) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnReceive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.onReceive.hasListener);
    },
    "call_HasOnReceive": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.udp.onReceive.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnReceive": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.onReceive.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnReceiveError": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp?.onReceiveError && "addListener" in WEBEXT?.sockets?.udp?.onReceiveError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnReceiveError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.onReceiveError.addListener);
    },
    "call_OnReceiveError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.udp.onReceiveError.addListener(A.H.get<object>(callback));
    },
    "try_OnReceiveError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.onReceiveError.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffReceiveError": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp?.onReceiveError && "removeListener" in WEBEXT?.sockets?.udp?.onReceiveError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffReceiveError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.onReceiveError.removeListener);
    },
    "call_OffReceiveError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.udp.onReceiveError.removeListener(A.H.get<object>(callback));
    },
    "try_OffReceiveError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.onReceiveError.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnReceiveError": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp?.onReceiveError && "hasListener" in WEBEXT?.sockets?.udp?.onReceiveError) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnReceiveError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.onReceiveError.hasListener);
    },
    "call_HasOnReceiveError": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.sockets.udp.onReceiveError.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnReceiveError": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.onReceiveError.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Send": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp && "send" in WEBEXT?.sockets?.udp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Send": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.send);
    },
    "call_Send": (
      retPtr: Pointer,
      socketId: number,
      data: heap.Ref<object>,
      address: heap.Ref<object>,
      port: number,
      dnsQueryType: number,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.sockets.udp.send(
        socketId,
        A.H.get<object>(data),
        A.H.get<object>(address),
        port,
        dnsQueryType > 0 && dnsQueryType <= 3 ? ["any", "ipv4", "ipv6"][dnsQueryType - 1] : undefined,
        A.H.get<object>(callback)
      );
    },
    "try_Send": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      data: heap.Ref<object>,
      address: heap.Ref<object>,
      port: number,
      dnsQueryType: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.send(
          socketId,
          A.H.get<object>(data),
          A.H.get<object>(address),
          port,
          dnsQueryType > 0 && dnsQueryType <= 3 ? ["any", "ipv4", "ipv6"][dnsQueryType - 1] : undefined,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetBroadcast": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp && "setBroadcast" in WEBEXT?.sockets?.udp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetBroadcast": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.setBroadcast);
    },
    "call_SetBroadcast": (
      retPtr: Pointer,
      socketId: number,
      enabled: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.sockets.udp.setBroadcast(socketId, enabled === A.H.TRUE, A.H.get<object>(callback));
    },
    "try_SetBroadcast": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      enabled: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.setBroadcast(socketId, enabled === A.H.TRUE, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetMulticastLoopbackMode": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp && "setMulticastLoopbackMode" in WEBEXT?.sockets?.udp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetMulticastLoopbackMode": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.setMulticastLoopbackMode);
    },
    "call_SetMulticastLoopbackMode": (
      retPtr: Pointer,
      socketId: number,
      enabled: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.sockets.udp.setMulticastLoopbackMode(
        socketId,
        enabled === A.H.TRUE,
        A.H.get<object>(callback)
      );
    },
    "try_SetMulticastLoopbackMode": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      enabled: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.setMulticastLoopbackMode(
          socketId,
          enabled === A.H.TRUE,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetMulticastTimeToLive": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp && "setMulticastTimeToLive" in WEBEXT?.sockets?.udp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetMulticastTimeToLive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.setMulticastTimeToLive);
    },
    "call_SetMulticastTimeToLive": (
      retPtr: Pointer,
      socketId: number,
      ttl: number,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.sockets.udp.setMulticastTimeToLive(socketId, ttl, A.H.get<object>(callback));
    },
    "try_SetMulticastTimeToLive": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      ttl: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.setMulticastTimeToLive(socketId, ttl, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPaused": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp && "setPaused" in WEBEXT?.sockets?.udp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPaused": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.setPaused);
    },
    "call_SetPaused": (
      retPtr: Pointer,
      socketId: number,
      paused: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.sockets.udp.setPaused(socketId, paused === A.H.TRUE, A.H.get<object>(callback));
    },
    "try_SetPaused": (
      retPtr: Pointer,
      errPtr: Pointer,
      socketId: number,
      paused: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.sockets.udp.setPaused(socketId, paused === A.H.TRUE, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Update": (): heap.Ref<boolean> => {
      if (WEBEXT?.sockets?.udp && "update" in WEBEXT?.sockets?.udp) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Update": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.sockets.udp.update);
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

      const _ret = WEBEXT.sockets.udp.update(socketId, properties_ffi, A.H.get<object>(callback));
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

        const _ret = WEBEXT.sockets.udp.update(socketId, properties_ffi, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
