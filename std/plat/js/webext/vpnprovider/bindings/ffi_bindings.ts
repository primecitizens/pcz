import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/vpnprovider", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Parameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 32, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
      } else {
        A.store.Bool(ptr + 32, true);
        A.store.Ref(ptr + 0, x["address"]);
        A.store.Ref(ptr + 4, x["broadcastAddress"]);
        A.store.Ref(ptr + 8, x["mtu"]);
        A.store.Ref(ptr + 12, x["exclusionList"]);
        A.store.Ref(ptr + 16, x["inclusionList"]);
        A.store.Ref(ptr + 20, x["domainSearch"]);
        A.store.Ref(ptr + 24, x["dnsServers"]);
        A.store.Ref(ptr + 28, x["reconnect"]);
      }
    },
    "load_Parameters": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["address"] = A.load.Ref(ptr + 0, undefined);
      x["broadcastAddress"] = A.load.Ref(ptr + 4, undefined);
      x["mtu"] = A.load.Ref(ptr + 8, undefined);
      x["exclusionList"] = A.load.Ref(ptr + 12, undefined);
      x["inclusionList"] = A.load.Ref(ptr + 16, undefined);
      x["domainSearch"] = A.load.Ref(ptr + 20, undefined);
      x["dnsServers"] = A.load.Ref(ptr + 24, undefined);
      x["reconnect"] = A.load.Ref(ptr + 28, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PlatformMessage": (ref: heap.Ref<string>): number => {
      const idx = [
        "connected",
        "disconnected",
        "error",
        "linkDown",
        "linkUp",
        "linkChanged",
        "suspend",
        "resume",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_UIEvent": (ref: heap.Ref<string>): number => {
      const idx = ["showAddDialog", "showConfigureDialog"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_VpnConnectionState": (ref: heap.Ref<string>): number => {
      const idx = ["connected", "failure"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_CreateConfig": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider && "createConfig" in WEBEXT?.vpnProvider) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CreateConfig": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.createConfig);
    },
    "call_CreateConfig": (retPtr: Pointer, name: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.createConfig(A.H.get<object>(name));
      A.store.Ref(retPtr, _ret);
    },
    "try_CreateConfig": (retPtr: Pointer, errPtr: Pointer, name: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.createConfig(A.H.get<object>(name));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DestroyConfig": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider && "destroyConfig" in WEBEXT?.vpnProvider) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DestroyConfig": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.destroyConfig);
    },
    "call_DestroyConfig": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.destroyConfig(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_DestroyConfig": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.destroyConfig(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_NotifyConnectionStateChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider && "notifyConnectionStateChanged" in WEBEXT?.vpnProvider) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_NotifyConnectionStateChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.notifyConnectionStateChanged);
    },
    "call_NotifyConnectionStateChanged": (retPtr: Pointer, state: number): void => {
      const _ret = WEBEXT.vpnProvider.notifyConnectionStateChanged(
        state > 0 && state <= 2 ? ["connected", "failure"][state - 1] : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_NotifyConnectionStateChanged": (retPtr: Pointer, errPtr: Pointer, state: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.notifyConnectionStateChanged(
          state > 0 && state <= 2 ? ["connected", "failure"][state - 1] : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnConfigCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider?.onConfigCreated && "addListener" in WEBEXT?.vpnProvider?.onConfigCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnConfigCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.onConfigCreated.addListener);
    },
    "call_OnConfigCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.onConfigCreated.addListener(A.H.get<object>(callback));
    },
    "try_OnConfigCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.onConfigCreated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffConfigCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider?.onConfigCreated && "removeListener" in WEBEXT?.vpnProvider?.onConfigCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffConfigCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.onConfigCreated.removeListener);
    },
    "call_OffConfigCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.onConfigCreated.removeListener(A.H.get<object>(callback));
    },
    "try_OffConfigCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.onConfigCreated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnConfigCreated": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider?.onConfigCreated && "hasListener" in WEBEXT?.vpnProvider?.onConfigCreated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnConfigCreated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.onConfigCreated.hasListener);
    },
    "call_HasOnConfigCreated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.onConfigCreated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnConfigCreated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.onConfigCreated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnConfigRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider?.onConfigRemoved && "addListener" in WEBEXT?.vpnProvider?.onConfigRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnConfigRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.onConfigRemoved.addListener);
    },
    "call_OnConfigRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.onConfigRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnConfigRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.onConfigRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffConfigRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider?.onConfigRemoved && "removeListener" in WEBEXT?.vpnProvider?.onConfigRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffConfigRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.onConfigRemoved.removeListener);
    },
    "call_OffConfigRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.onConfigRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffConfigRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.onConfigRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnConfigRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider?.onConfigRemoved && "hasListener" in WEBEXT?.vpnProvider?.onConfigRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnConfigRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.onConfigRemoved.hasListener);
    },
    "call_HasOnConfigRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.onConfigRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnConfigRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.onConfigRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPacketReceived": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider?.onPacketReceived && "addListener" in WEBEXT?.vpnProvider?.onPacketReceived) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPacketReceived": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.onPacketReceived.addListener);
    },
    "call_OnPacketReceived": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.onPacketReceived.addListener(A.H.get<object>(callback));
    },
    "try_OnPacketReceived": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.onPacketReceived.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPacketReceived": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider?.onPacketReceived && "removeListener" in WEBEXT?.vpnProvider?.onPacketReceived) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPacketReceived": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.onPacketReceived.removeListener);
    },
    "call_OffPacketReceived": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.onPacketReceived.removeListener(A.H.get<object>(callback));
    },
    "try_OffPacketReceived": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.onPacketReceived.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPacketReceived": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider?.onPacketReceived && "hasListener" in WEBEXT?.vpnProvider?.onPacketReceived) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPacketReceived": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.onPacketReceived.hasListener);
    },
    "call_HasOnPacketReceived": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.onPacketReceived.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPacketReceived": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.onPacketReceived.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPlatformMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider?.onPlatformMessage && "addListener" in WEBEXT?.vpnProvider?.onPlatformMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPlatformMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.onPlatformMessage.addListener);
    },
    "call_OnPlatformMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.onPlatformMessage.addListener(A.H.get<object>(callback));
    },
    "try_OnPlatformMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.onPlatformMessage.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPlatformMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider?.onPlatformMessage && "removeListener" in WEBEXT?.vpnProvider?.onPlatformMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPlatformMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.onPlatformMessage.removeListener);
    },
    "call_OffPlatformMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.onPlatformMessage.removeListener(A.H.get<object>(callback));
    },
    "try_OffPlatformMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.onPlatformMessage.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPlatformMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider?.onPlatformMessage && "hasListener" in WEBEXT?.vpnProvider?.onPlatformMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPlatformMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.onPlatformMessage.hasListener);
    },
    "call_HasOnPlatformMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.onPlatformMessage.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPlatformMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.onPlatformMessage.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnUIEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider?.onUIEvent && "addListener" in WEBEXT?.vpnProvider?.onUIEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnUIEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.onUIEvent.addListener);
    },
    "call_OnUIEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.onUIEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnUIEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.onUIEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffUIEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider?.onUIEvent && "removeListener" in WEBEXT?.vpnProvider?.onUIEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffUIEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.onUIEvent.removeListener);
    },
    "call_OffUIEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.onUIEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffUIEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.onUIEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnUIEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider?.onUIEvent && "hasListener" in WEBEXT?.vpnProvider?.onUIEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnUIEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.onUIEvent.hasListener);
    },
    "call_HasOnUIEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.onUIEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnUIEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.onUIEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendPacket": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider && "sendPacket" in WEBEXT?.vpnProvider) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendPacket": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.sendPacket);
    },
    "call_SendPacket": (retPtr: Pointer, data: heap.Ref<object>): void => {
      const _ret = WEBEXT.vpnProvider.sendPacket(A.H.get<object>(data));
      A.store.Ref(retPtr, _ret);
    },
    "try_SendPacket": (retPtr: Pointer, errPtr: Pointer, data: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.vpnProvider.sendPacket(A.H.get<object>(data));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetParameters": (): heap.Ref<boolean> => {
      if (WEBEXT?.vpnProvider && "setParameters" in WEBEXT?.vpnProvider) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetParameters": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.vpnProvider.setParameters);
    },
    "call_SetParameters": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["address"] = A.load.Ref(parameters + 0, undefined);
      parameters_ffi["broadcastAddress"] = A.load.Ref(parameters + 4, undefined);
      parameters_ffi["mtu"] = A.load.Ref(parameters + 8, undefined);
      parameters_ffi["exclusionList"] = A.load.Ref(parameters + 12, undefined);
      parameters_ffi["inclusionList"] = A.load.Ref(parameters + 16, undefined);
      parameters_ffi["domainSearch"] = A.load.Ref(parameters + 20, undefined);
      parameters_ffi["dnsServers"] = A.load.Ref(parameters + 24, undefined);
      parameters_ffi["reconnect"] = A.load.Ref(parameters + 28, undefined);

      const _ret = WEBEXT.vpnProvider.setParameters(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetParameters": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["address"] = A.load.Ref(parameters + 0, undefined);
        parameters_ffi["broadcastAddress"] = A.load.Ref(parameters + 4, undefined);
        parameters_ffi["mtu"] = A.load.Ref(parameters + 8, undefined);
        parameters_ffi["exclusionList"] = A.load.Ref(parameters + 12, undefined);
        parameters_ffi["inclusionList"] = A.load.Ref(parameters + 16, undefined);
        parameters_ffi["domainSearch"] = A.load.Ref(parameters + 20, undefined);
        parameters_ffi["dnsServers"] = A.load.Ref(parameters + 24, undefined);
        parameters_ffi["reconnect"] = A.load.Ref(parameters + 28, undefined);

        const _ret = WEBEXT.vpnProvider.setParameters(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
