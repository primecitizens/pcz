import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/enterprise/deviceattributes", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_GetDeviceAnnotatedLocation": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.enterprise?.deviceAttributes &&
        "getDeviceAnnotatedLocation" in WEBEXT?.enterprise?.deviceAttributes
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDeviceAnnotatedLocation": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.deviceAttributes.getDeviceAnnotatedLocation);
    },
    "call_GetDeviceAnnotatedLocation": (retPtr: Pointer): void => {
      const _ret = WEBEXT.enterprise.deviceAttributes.getDeviceAnnotatedLocation();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDeviceAnnotatedLocation": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.deviceAttributes.getDeviceAnnotatedLocation();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDeviceAssetId": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.deviceAttributes && "getDeviceAssetId" in WEBEXT?.enterprise?.deviceAttributes) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDeviceAssetId": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.deviceAttributes.getDeviceAssetId);
    },
    "call_GetDeviceAssetId": (retPtr: Pointer): void => {
      const _ret = WEBEXT.enterprise.deviceAttributes.getDeviceAssetId();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDeviceAssetId": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.deviceAttributes.getDeviceAssetId();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDeviceHostname": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.deviceAttributes && "getDeviceHostname" in WEBEXT?.enterprise?.deviceAttributes) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDeviceHostname": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.deviceAttributes.getDeviceHostname);
    },
    "call_GetDeviceHostname": (retPtr: Pointer): void => {
      const _ret = WEBEXT.enterprise.deviceAttributes.getDeviceHostname();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDeviceHostname": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.deviceAttributes.getDeviceHostname();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDeviceSerialNumber": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.deviceAttributes && "getDeviceSerialNumber" in WEBEXT?.enterprise?.deviceAttributes) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDeviceSerialNumber": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.deviceAttributes.getDeviceSerialNumber);
    },
    "call_GetDeviceSerialNumber": (retPtr: Pointer): void => {
      const _ret = WEBEXT.enterprise.deviceAttributes.getDeviceSerialNumber();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDeviceSerialNumber": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.deviceAttributes.getDeviceSerialNumber();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDirectoryDeviceId": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.deviceAttributes && "getDirectoryDeviceId" in WEBEXT?.enterprise?.deviceAttributes) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDirectoryDeviceId": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.deviceAttributes.getDirectoryDeviceId);
    },
    "call_GetDirectoryDeviceId": (retPtr: Pointer): void => {
      const _ret = WEBEXT.enterprise.deviceAttributes.getDirectoryDeviceId();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDirectoryDeviceId": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.deviceAttributes.getDirectoryDeviceId();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
