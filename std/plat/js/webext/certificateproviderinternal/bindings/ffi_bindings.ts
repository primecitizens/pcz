import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/certificateproviderinternal", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_ReportCertificates": (): heap.Ref<boolean> => {
      if (WEBEXT?.certificateProviderInternal && "reportCertificates" in WEBEXT?.certificateProviderInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportCertificates": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProviderInternal.reportCertificates);
    },
    "call_ReportCertificates": (retPtr: Pointer, requestId: number, certificates: heap.Ref<object>): void => {
      const _ret = WEBEXT.certificateProviderInternal.reportCertificates(requestId, A.H.get<object>(certificates));
      A.store.Ref(retPtr, _ret);
    },
    "try_ReportCertificates": (
      retPtr: Pointer,
      errPtr: Pointer,
      requestId: number,
      certificates: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.certificateProviderInternal.reportCertificates(requestId, A.H.get<object>(certificates));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ReportSignature": (): heap.Ref<boolean> => {
      if (WEBEXT?.certificateProviderInternal && "reportSignature" in WEBEXT?.certificateProviderInternal) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ReportSignature": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.certificateProviderInternal.reportSignature);
    },
    "call_ReportSignature": (retPtr: Pointer, requestId: number, signature: heap.Ref<object>): void => {
      const _ret = WEBEXT.certificateProviderInternal.reportSignature(requestId, A.H.get<object>(signature));
      A.store.Ref(retPtr, _ret);
    },
    "try_ReportSignature": (
      retPtr: Pointer,
      errPtr: Pointer,
      requestId: number,
      signature: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.certificateProviderInternal.reportSignature(requestId, A.H.get<object>(signature));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
