import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/imageloaderprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_GetArcDocumentsProviderThumbnail": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageLoaderPrivate && "getArcDocumentsProviderThumbnail" in WEBEXT?.imageLoaderPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetArcDocumentsProviderThumbnail": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageLoaderPrivate.getArcDocumentsProviderThumbnail);
    },
    "call_GetArcDocumentsProviderThumbnail": (
      retPtr: Pointer,
      url: heap.Ref<object>,
      widthHint: number,
      heightHint: number,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.imageLoaderPrivate.getArcDocumentsProviderThumbnail(
        A.H.get<object>(url),
        widthHint,
        heightHint,
        A.H.get<object>(callback)
      );
    },
    "try_GetArcDocumentsProviderThumbnail": (
      retPtr: Pointer,
      errPtr: Pointer,
      url: heap.Ref<object>,
      widthHint: number,
      heightHint: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageLoaderPrivate.getArcDocumentsProviderThumbnail(
          A.H.get<object>(url),
          widthHint,
          heightHint,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDriveThumbnail": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageLoaderPrivate && "getDriveThumbnail" in WEBEXT?.imageLoaderPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDriveThumbnail": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageLoaderPrivate.getDriveThumbnail);
    },
    "call_GetDriveThumbnail": (
      retPtr: Pointer,
      url: heap.Ref<object>,
      cropToSquare: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.imageLoaderPrivate.getDriveThumbnail(
        A.H.get<object>(url),
        cropToSquare === A.H.TRUE,
        A.H.get<object>(callback)
      );
    },
    "try_GetDriveThumbnail": (
      retPtr: Pointer,
      errPtr: Pointer,
      url: heap.Ref<object>,
      cropToSquare: heap.Ref<boolean>,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageLoaderPrivate.getDriveThumbnail(
          A.H.get<object>(url),
          cropToSquare === A.H.TRUE,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPdfThumbnail": (): heap.Ref<boolean> => {
      if (WEBEXT?.imageLoaderPrivate && "getPdfThumbnail" in WEBEXT?.imageLoaderPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPdfThumbnail": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.imageLoaderPrivate.getPdfThumbnail);
    },
    "call_GetPdfThumbnail": (
      retPtr: Pointer,
      url: heap.Ref<object>,
      width: number,
      height: number,
      callback: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.imageLoaderPrivate.getPdfThumbnail(
        A.H.get<object>(url),
        width,
        height,
        A.H.get<object>(callback)
      );
    },
    "try_GetPdfThumbnail": (
      retPtr: Pointer,
      errPtr: Pointer,
      url: heap.Ref<object>,
      width: number,
      height: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.imageLoaderPrivate.getPdfThumbnail(
          A.H.get<object>(url),
          width,
          height,
          A.H.get<object>(callback)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
