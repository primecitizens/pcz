import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/hid", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_HidConnectInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
        A.store.Int32(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Bool(ptr + 4, "connectionId" in x ? true : false);
        A.store.Int32(ptr + 0, x["connectionId"] === undefined ? 0 : (x["connectionId"] as number));
      }
    },
    "load_HidConnectInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 4)) {
        x["connectionId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["connectionId"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DeviceFilter": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 18, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 19, false);
        A.store.Int32(ptr + 12, 0);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Bool(ptr + 16, "vendorId" in x ? true : false);
        A.store.Int32(ptr + 0, x["vendorId"] === undefined ? 0 : (x["vendorId"] as number));
        A.store.Bool(ptr + 17, "productId" in x ? true : false);
        A.store.Int32(ptr + 4, x["productId"] === undefined ? 0 : (x["productId"] as number));
        A.store.Bool(ptr + 18, "usagePage" in x ? true : false);
        A.store.Int32(ptr + 8, x["usagePage"] === undefined ? 0 : (x["usagePage"] as number));
        A.store.Bool(ptr + 19, "usage" in x ? true : false);
        A.store.Int32(ptr + 12, x["usage"] === undefined ? 0 : (x["usage"] as number));
      }
    },
    "load_DeviceFilter": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["vendorId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["vendorId"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["productId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["productId"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["usagePage"] = A.load.Int32(ptr + 8);
      } else {
        delete x["usagePage"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["usage"] = A.load.Int32(ptr + 12);
      } else {
        delete x["usage"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_HidCollectionInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Bool(ptr + 12, "usagePage" in x ? true : false);
        A.store.Int32(ptr + 0, x["usagePage"] === undefined ? 0 : (x["usagePage"] as number));
        A.store.Bool(ptr + 13, "usage" in x ? true : false);
        A.store.Int32(ptr + 4, x["usage"] === undefined ? 0 : (x["usage"] as number));
        A.store.Ref(ptr + 8, x["reportIds"]);
      }
    },
    "load_HidCollectionInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["usagePage"] = A.load.Int32(ptr + 0);
      } else {
        delete x["usagePage"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["usage"] = A.load.Int32(ptr + 4);
      } else {
        delete x["usage"];
      }
      x["reportIds"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_HidDeviceInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 46, false);
        A.store.Bool(ptr + 40, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 41, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 42, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Bool(ptr + 43, false);
        A.store.Int32(ptr + 24, 0);
        A.store.Bool(ptr + 44, false);
        A.store.Int32(ptr + 28, 0);
        A.store.Bool(ptr + 45, false);
        A.store.Int32(ptr + 32, 0);
        A.store.Ref(ptr + 36, undefined);
      } else {
        A.store.Bool(ptr + 46, true);
        A.store.Bool(ptr + 40, "deviceId" in x ? true : false);
        A.store.Int32(ptr + 0, x["deviceId"] === undefined ? 0 : (x["deviceId"] as number));
        A.store.Bool(ptr + 41, "vendorId" in x ? true : false);
        A.store.Int32(ptr + 4, x["vendorId"] === undefined ? 0 : (x["vendorId"] as number));
        A.store.Bool(ptr + 42, "productId" in x ? true : false);
        A.store.Int32(ptr + 8, x["productId"] === undefined ? 0 : (x["productId"] as number));
        A.store.Ref(ptr + 12, x["productName"]);
        A.store.Ref(ptr + 16, x["serialNumber"]);
        A.store.Ref(ptr + 20, x["collections"]);
        A.store.Bool(ptr + 43, "maxInputReportSize" in x ? true : false);
        A.store.Int32(ptr + 24, x["maxInputReportSize"] === undefined ? 0 : (x["maxInputReportSize"] as number));
        A.store.Bool(ptr + 44, "maxOutputReportSize" in x ? true : false);
        A.store.Int32(ptr + 28, x["maxOutputReportSize"] === undefined ? 0 : (x["maxOutputReportSize"] as number));
        A.store.Bool(ptr + 45, "maxFeatureReportSize" in x ? true : false);
        A.store.Int32(ptr + 32, x["maxFeatureReportSize"] === undefined ? 0 : (x["maxFeatureReportSize"] as number));
        A.store.Ref(ptr + 36, x["reportDescriptor"]);
      }
    },
    "load_HidDeviceInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 40)) {
        x["deviceId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["deviceId"];
      }
      if (A.load.Bool(ptr + 41)) {
        x["vendorId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["vendorId"];
      }
      if (A.load.Bool(ptr + 42)) {
        x["productId"] = A.load.Int32(ptr + 8);
      } else {
        delete x["productId"];
      }
      x["productName"] = A.load.Ref(ptr + 12, undefined);
      x["serialNumber"] = A.load.Ref(ptr + 16, undefined);
      x["collections"] = A.load.Ref(ptr + 20, undefined);
      if (A.load.Bool(ptr + 43)) {
        x["maxInputReportSize"] = A.load.Int32(ptr + 24);
      } else {
        delete x["maxInputReportSize"];
      }
      if (A.load.Bool(ptr + 44)) {
        x["maxOutputReportSize"] = A.load.Int32(ptr + 28);
      } else {
        delete x["maxOutputReportSize"];
      }
      if (A.load.Bool(ptr + 45)) {
        x["maxFeatureReportSize"] = A.load.Int32(ptr + 32);
      } else {
        delete x["maxFeatureReportSize"];
      }
      x["reportDescriptor"] = A.load.Ref(ptr + 36, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetDevicesOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Bool(ptr + 12, "vendorId" in x ? true : false);
        A.store.Int32(ptr + 0, x["vendorId"] === undefined ? 0 : (x["vendorId"] as number));
        A.store.Bool(ptr + 13, "productId" in x ? true : false);
        A.store.Int32(ptr + 4, x["productId"] === undefined ? 0 : (x["productId"] as number));
        A.store.Ref(ptr + 8, x["filters"]);
      }
    },
    "load_GetDevicesOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["vendorId"] = A.load.Int32(ptr + 0);
      } else {
        delete x["vendorId"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["productId"] = A.load.Int32(ptr + 4);
      } else {
        delete x["productId"];
      }
      x["filters"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_Connect": (): heap.Ref<boolean> => {
      if (WEBEXT?.hid && "connect" in WEBEXT?.hid) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Connect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.hid.connect);
    },
    "call_Connect": (retPtr: Pointer, deviceId: number): void => {
      const _ret = WEBEXT.hid.connect(deviceId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Connect": (retPtr: Pointer, errPtr: Pointer, deviceId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.hid.connect(deviceId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Disconnect": (): heap.Ref<boolean> => {
      if (WEBEXT?.hid && "disconnect" in WEBEXT?.hid) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Disconnect": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.hid.disconnect);
    },
    "call_Disconnect": (retPtr: Pointer, connectionId: number): void => {
      const _ret = WEBEXT.hid.disconnect(connectionId);
      A.store.Ref(retPtr, _ret);
    },
    "try_Disconnect": (retPtr: Pointer, errPtr: Pointer, connectionId: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.hid.disconnect(connectionId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDevices": (): heap.Ref<boolean> => {
      if (WEBEXT?.hid && "getDevices" in WEBEXT?.hid) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDevices": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.hid.getDevices);
    },
    "call_GetDevices": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 12)) {
        options_ffi["vendorId"] = A.load.Int32(options + 0);
      }
      if (A.load.Bool(options + 13)) {
        options_ffi["productId"] = A.load.Int32(options + 4);
      }
      options_ffi["filters"] = A.load.Ref(options + 8, undefined);

      const _ret = WEBEXT.hid.getDevices(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDevices": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 12)) {
          options_ffi["vendorId"] = A.load.Int32(options + 0);
        }
        if (A.load.Bool(options + 13)) {
          options_ffi["productId"] = A.load.Int32(options + 4);
        }
        options_ffi["filters"] = A.load.Ref(options + 8, undefined);

        const _ret = WEBEXT.hid.getDevices(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeviceAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.hid?.onDeviceAdded && "addListener" in WEBEXT?.hid?.onDeviceAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeviceAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.hid.onDeviceAdded.addListener);
    },
    "call_OnDeviceAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.hid.onDeviceAdded.addListener(A.H.get<object>(callback));
    },
    "try_OnDeviceAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.hid.onDeviceAdded.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeviceAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.hid?.onDeviceAdded && "removeListener" in WEBEXT?.hid?.onDeviceAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeviceAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.hid.onDeviceAdded.removeListener);
    },
    "call_OffDeviceAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.hid.onDeviceAdded.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeviceAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.hid.onDeviceAdded.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeviceAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.hid?.onDeviceAdded && "hasListener" in WEBEXT?.hid?.onDeviceAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeviceAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.hid.onDeviceAdded.hasListener);
    },
    "call_HasOnDeviceAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.hid.onDeviceAdded.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeviceAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.hid.onDeviceAdded.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeviceRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.hid?.onDeviceRemoved && "addListener" in WEBEXT?.hid?.onDeviceRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeviceRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.hid.onDeviceRemoved.addListener);
    },
    "call_OnDeviceRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.hid.onDeviceRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnDeviceRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.hid.onDeviceRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeviceRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.hid?.onDeviceRemoved && "removeListener" in WEBEXT?.hid?.onDeviceRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeviceRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.hid.onDeviceRemoved.removeListener);
    },
    "call_OffDeviceRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.hid.onDeviceRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeviceRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.hid.onDeviceRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeviceRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.hid?.onDeviceRemoved && "hasListener" in WEBEXT?.hid?.onDeviceRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeviceRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.hid.onDeviceRemoved.hasListener);
    },
    "call_HasOnDeviceRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.hid.onDeviceRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeviceRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.hid.onDeviceRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Receive": (): heap.Ref<boolean> => {
      if (WEBEXT?.hid && "receive" in WEBEXT?.hid) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Receive": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.hid.receive);
    },
    "call_Receive": (retPtr: Pointer, connectionId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.hid.receive(connectionId, A.H.get<object>(callback));
    },
    "try_Receive": (
      retPtr: Pointer,
      errPtr: Pointer,
      connectionId: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.hid.receive(connectionId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReceiveFeatureReport": (): heap.Ref<boolean> => {
      if (WEBEXT?.hid && "receiveFeatureReport" in WEBEXT?.hid) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReceiveFeatureReport": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.hid.receiveFeatureReport);
    },
    "call_ReceiveFeatureReport": (retPtr: Pointer, connectionId: number, reportId: number): void => {
      const _ret = WEBEXT.hid.receiveFeatureReport(connectionId, reportId);
      A.store.Ref(retPtr, _ret);
    },
    "try_ReceiveFeatureReport": (
      retPtr: Pointer,
      errPtr: Pointer,
      connectionId: number,
      reportId: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.hid.receiveFeatureReport(connectionId, reportId);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Send": (): heap.Ref<boolean> => {
      if (WEBEXT?.hid && "send" in WEBEXT?.hid) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Send": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.hid.send);
    },
    "call_Send": (retPtr: Pointer, connectionId: number, reportId: number, data: heap.Ref<object>): void => {
      const _ret = WEBEXT.hid.send(connectionId, reportId, A.H.get<object>(data));
      A.store.Ref(retPtr, _ret);
    },
    "try_Send": (
      retPtr: Pointer,
      errPtr: Pointer,
      connectionId: number,
      reportId: number,
      data: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.hid.send(connectionId, reportId, A.H.get<object>(data));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendFeatureReport": (): heap.Ref<boolean> => {
      if (WEBEXT?.hid && "sendFeatureReport" in WEBEXT?.hid) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendFeatureReport": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.hid.sendFeatureReport);
    },
    "call_SendFeatureReport": (
      retPtr: Pointer,
      connectionId: number,
      reportId: number,
      data: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.hid.sendFeatureReport(connectionId, reportId, A.H.get<object>(data));
      A.store.Ref(retPtr, _ret);
    },
    "try_SendFeatureReport": (
      retPtr: Pointer,
      errPtr: Pointer,
      connectionId: number,
      reportId: number,
      data: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.hid.sendFeatureReport(connectionId, reportId, A.H.get<object>(data));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
