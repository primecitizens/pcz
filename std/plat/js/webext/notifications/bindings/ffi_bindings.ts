import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/notifications", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_NotificationBitmap": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Bool(ptr + 12, "width" in x ? true : false);
        A.store.Int32(ptr + 0, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Bool(ptr + 13, "height" in x ? true : false);
        A.store.Int32(ptr + 4, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Ref(ptr + 8, x["data"]);
      }
    },
    "load_NotificationBitmap": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["width"] = A.load.Int32(ptr + 0);
      } else {
        delete x["width"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["height"] = A.load.Int32(ptr + 4);
      } else {
        delete x["height"];
      }
      x["data"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_NotificationButton": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 23, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);

        A.store.Bool(ptr + 8 + 14, false);
        A.store.Bool(ptr + 8 + 12, false);
        A.store.Int32(ptr + 8 + 0, 0);
        A.store.Bool(ptr + 8 + 13, false);
        A.store.Int32(ptr + 8 + 4, 0);
        A.store.Ref(ptr + 8 + 8, undefined);
      } else {
        A.store.Bool(ptr + 23, true);
        A.store.Ref(ptr + 0, x["title"]);
        A.store.Ref(ptr + 4, x["iconUrl"]);

        if (typeof x["iconBitmap"] === "undefined") {
          A.store.Bool(ptr + 8 + 14, false);
          A.store.Bool(ptr + 8 + 12, false);
          A.store.Int32(ptr + 8 + 0, 0);
          A.store.Bool(ptr + 8 + 13, false);
          A.store.Int32(ptr + 8 + 4, 0);
          A.store.Ref(ptr + 8 + 8, undefined);
        } else {
          A.store.Bool(ptr + 8 + 14, true);
          A.store.Bool(ptr + 8 + 12, "width" in x["iconBitmap"] ? true : false);
          A.store.Int32(ptr + 8 + 0, x["iconBitmap"]["width"] === undefined ? 0 : (x["iconBitmap"]["width"] as number));
          A.store.Bool(ptr + 8 + 13, "height" in x["iconBitmap"] ? true : false);
          A.store.Int32(
            ptr + 8 + 4,
            x["iconBitmap"]["height"] === undefined ? 0 : (x["iconBitmap"]["height"] as number)
          );
          A.store.Ref(ptr + 8 + 8, x["iconBitmap"]["data"]);
        }
      }
    },
    "load_NotificationButton": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["title"] = A.load.Ref(ptr + 0, undefined);
      x["iconUrl"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 8 + 14)) {
        x["iconBitmap"] = {};
        if (A.load.Bool(ptr + 8 + 12)) {
          x["iconBitmap"]["width"] = A.load.Int32(ptr + 8 + 0);
        } else {
          delete x["iconBitmap"]["width"];
        }
        if (A.load.Bool(ptr + 8 + 13)) {
          x["iconBitmap"]["height"] = A.load.Int32(ptr + 8 + 4);
        } else {
          delete x["iconBitmap"]["height"];
        }
        x["iconBitmap"]["data"] = A.load.Ref(ptr + 8 + 8, undefined);
      } else {
        delete x["iconBitmap"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_NotificationItem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["title"]);
        A.store.Ref(ptr + 4, x["message"]);
      }
    },
    "load_NotificationItem": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["title"] = A.load.Ref(ptr + 0, undefined);
      x["message"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_TemplateType": (ref: heap.Ref<string>): number => {
      const idx = ["basic", "image", "list", "progress"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_NotificationOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 117, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);

        A.store.Bool(ptr + 8 + 14, false);
        A.store.Bool(ptr + 8 + 12, false);
        A.store.Int32(ptr + 8 + 0, 0);
        A.store.Bool(ptr + 8 + 13, false);
        A.store.Int32(ptr + 8 + 4, 0);
        A.store.Ref(ptr + 8 + 8, undefined);
        A.store.Ref(ptr + 24, undefined);

        A.store.Bool(ptr + 28 + 14, false);
        A.store.Bool(ptr + 28 + 12, false);
        A.store.Int32(ptr + 28 + 0, 0);
        A.store.Bool(ptr + 28 + 13, false);
        A.store.Int32(ptr + 28 + 4, 0);
        A.store.Ref(ptr + 28 + 8, undefined);
        A.store.Ref(ptr + 44, undefined);
        A.store.Ref(ptr + 48, undefined);
        A.store.Ref(ptr + 52, undefined);
        A.store.Bool(ptr + 111, false);
        A.store.Int32(ptr + 56, 0);
        A.store.Bool(ptr + 112, false);
        A.store.Float64(ptr + 64, 0);
        A.store.Ref(ptr + 72, undefined);
        A.store.Ref(ptr + 76, undefined);
        A.store.Ref(ptr + 80, undefined);

        A.store.Bool(ptr + 84 + 14, false);
        A.store.Bool(ptr + 84 + 12, false);
        A.store.Int32(ptr + 84 + 0, 0);
        A.store.Bool(ptr + 84 + 13, false);
        A.store.Int32(ptr + 84 + 4, 0);
        A.store.Ref(ptr + 84 + 8, undefined);
        A.store.Ref(ptr + 100, undefined);
        A.store.Bool(ptr + 113, false);
        A.store.Int32(ptr + 104, 0);
        A.store.Bool(ptr + 114, false);
        A.store.Bool(ptr + 108, false);
        A.store.Bool(ptr + 115, false);
        A.store.Bool(ptr + 109, false);
        A.store.Bool(ptr + 116, false);
        A.store.Bool(ptr + 110, false);
      } else {
        A.store.Bool(ptr + 117, true);
        A.store.Enum(ptr + 0, ["basic", "image", "list", "progress"].indexOf(x["type"] as string));
        A.store.Ref(ptr + 4, x["iconUrl"]);

        if (typeof x["iconBitmap"] === "undefined") {
          A.store.Bool(ptr + 8 + 14, false);
          A.store.Bool(ptr + 8 + 12, false);
          A.store.Int32(ptr + 8 + 0, 0);
          A.store.Bool(ptr + 8 + 13, false);
          A.store.Int32(ptr + 8 + 4, 0);
          A.store.Ref(ptr + 8 + 8, undefined);
        } else {
          A.store.Bool(ptr + 8 + 14, true);
          A.store.Bool(ptr + 8 + 12, "width" in x["iconBitmap"] ? true : false);
          A.store.Int32(ptr + 8 + 0, x["iconBitmap"]["width"] === undefined ? 0 : (x["iconBitmap"]["width"] as number));
          A.store.Bool(ptr + 8 + 13, "height" in x["iconBitmap"] ? true : false);
          A.store.Int32(
            ptr + 8 + 4,
            x["iconBitmap"]["height"] === undefined ? 0 : (x["iconBitmap"]["height"] as number)
          );
          A.store.Ref(ptr + 8 + 8, x["iconBitmap"]["data"]);
        }
        A.store.Ref(ptr + 24, x["appIconMaskUrl"]);

        if (typeof x["appIconMaskBitmap"] === "undefined") {
          A.store.Bool(ptr + 28 + 14, false);
          A.store.Bool(ptr + 28 + 12, false);
          A.store.Int32(ptr + 28 + 0, 0);
          A.store.Bool(ptr + 28 + 13, false);
          A.store.Int32(ptr + 28 + 4, 0);
          A.store.Ref(ptr + 28 + 8, undefined);
        } else {
          A.store.Bool(ptr + 28 + 14, true);
          A.store.Bool(ptr + 28 + 12, "width" in x["appIconMaskBitmap"] ? true : false);
          A.store.Int32(
            ptr + 28 + 0,
            x["appIconMaskBitmap"]["width"] === undefined ? 0 : (x["appIconMaskBitmap"]["width"] as number)
          );
          A.store.Bool(ptr + 28 + 13, "height" in x["appIconMaskBitmap"] ? true : false);
          A.store.Int32(
            ptr + 28 + 4,
            x["appIconMaskBitmap"]["height"] === undefined ? 0 : (x["appIconMaskBitmap"]["height"] as number)
          );
          A.store.Ref(ptr + 28 + 8, x["appIconMaskBitmap"]["data"]);
        }
        A.store.Ref(ptr + 44, x["title"]);
        A.store.Ref(ptr + 48, x["message"]);
        A.store.Ref(ptr + 52, x["contextMessage"]);
        A.store.Bool(ptr + 111, "priority" in x ? true : false);
        A.store.Int32(ptr + 56, x["priority"] === undefined ? 0 : (x["priority"] as number));
        A.store.Bool(ptr + 112, "eventTime" in x ? true : false);
        A.store.Float64(ptr + 64, x["eventTime"] === undefined ? 0 : (x["eventTime"] as number));
        A.store.Ref(ptr + 72, x["buttons"]);
        A.store.Ref(ptr + 76, x["expandedMessage"]);
        A.store.Ref(ptr + 80, x["imageUrl"]);

        if (typeof x["imageBitmap"] === "undefined") {
          A.store.Bool(ptr + 84 + 14, false);
          A.store.Bool(ptr + 84 + 12, false);
          A.store.Int32(ptr + 84 + 0, 0);
          A.store.Bool(ptr + 84 + 13, false);
          A.store.Int32(ptr + 84 + 4, 0);
          A.store.Ref(ptr + 84 + 8, undefined);
        } else {
          A.store.Bool(ptr + 84 + 14, true);
          A.store.Bool(ptr + 84 + 12, "width" in x["imageBitmap"] ? true : false);
          A.store.Int32(
            ptr + 84 + 0,
            x["imageBitmap"]["width"] === undefined ? 0 : (x["imageBitmap"]["width"] as number)
          );
          A.store.Bool(ptr + 84 + 13, "height" in x["imageBitmap"] ? true : false);
          A.store.Int32(
            ptr + 84 + 4,
            x["imageBitmap"]["height"] === undefined ? 0 : (x["imageBitmap"]["height"] as number)
          );
          A.store.Ref(ptr + 84 + 8, x["imageBitmap"]["data"]);
        }
        A.store.Ref(ptr + 100, x["items"]);
        A.store.Bool(ptr + 113, "progress" in x ? true : false);
        A.store.Int32(ptr + 104, x["progress"] === undefined ? 0 : (x["progress"] as number));
        A.store.Bool(ptr + 114, "isClickable" in x ? true : false);
        A.store.Bool(ptr + 108, x["isClickable"] ? true : false);
        A.store.Bool(ptr + 115, "requireInteraction" in x ? true : false);
        A.store.Bool(ptr + 109, x["requireInteraction"] ? true : false);
        A.store.Bool(ptr + 116, "silent" in x ? true : false);
        A.store.Bool(ptr + 110, x["silent"] ? true : false);
      }
    },
    "load_NotificationOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["type"] = A.load.Enum(ptr + 0, ["basic", "image", "list", "progress"]);
      x["iconUrl"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 8 + 14)) {
        x["iconBitmap"] = {};
        if (A.load.Bool(ptr + 8 + 12)) {
          x["iconBitmap"]["width"] = A.load.Int32(ptr + 8 + 0);
        } else {
          delete x["iconBitmap"]["width"];
        }
        if (A.load.Bool(ptr + 8 + 13)) {
          x["iconBitmap"]["height"] = A.load.Int32(ptr + 8 + 4);
        } else {
          delete x["iconBitmap"]["height"];
        }
        x["iconBitmap"]["data"] = A.load.Ref(ptr + 8 + 8, undefined);
      } else {
        delete x["iconBitmap"];
      }
      x["appIconMaskUrl"] = A.load.Ref(ptr + 24, undefined);
      if (A.load.Bool(ptr + 28 + 14)) {
        x["appIconMaskBitmap"] = {};
        if (A.load.Bool(ptr + 28 + 12)) {
          x["appIconMaskBitmap"]["width"] = A.load.Int32(ptr + 28 + 0);
        } else {
          delete x["appIconMaskBitmap"]["width"];
        }
        if (A.load.Bool(ptr + 28 + 13)) {
          x["appIconMaskBitmap"]["height"] = A.load.Int32(ptr + 28 + 4);
        } else {
          delete x["appIconMaskBitmap"]["height"];
        }
        x["appIconMaskBitmap"]["data"] = A.load.Ref(ptr + 28 + 8, undefined);
      } else {
        delete x["appIconMaskBitmap"];
      }
      x["title"] = A.load.Ref(ptr + 44, undefined);
      x["message"] = A.load.Ref(ptr + 48, undefined);
      x["contextMessage"] = A.load.Ref(ptr + 52, undefined);
      if (A.load.Bool(ptr + 111)) {
        x["priority"] = A.load.Int32(ptr + 56);
      } else {
        delete x["priority"];
      }
      if (A.load.Bool(ptr + 112)) {
        x["eventTime"] = A.load.Float64(ptr + 64);
      } else {
        delete x["eventTime"];
      }
      x["buttons"] = A.load.Ref(ptr + 72, undefined);
      x["expandedMessage"] = A.load.Ref(ptr + 76, undefined);
      x["imageUrl"] = A.load.Ref(ptr + 80, undefined);
      if (A.load.Bool(ptr + 84 + 14)) {
        x["imageBitmap"] = {};
        if (A.load.Bool(ptr + 84 + 12)) {
          x["imageBitmap"]["width"] = A.load.Int32(ptr + 84 + 0);
        } else {
          delete x["imageBitmap"]["width"];
        }
        if (A.load.Bool(ptr + 84 + 13)) {
          x["imageBitmap"]["height"] = A.load.Int32(ptr + 84 + 4);
        } else {
          delete x["imageBitmap"]["height"];
        }
        x["imageBitmap"]["data"] = A.load.Ref(ptr + 84 + 8, undefined);
      } else {
        delete x["imageBitmap"];
      }
      x["items"] = A.load.Ref(ptr + 100, undefined);
      if (A.load.Bool(ptr + 113)) {
        x["progress"] = A.load.Int32(ptr + 104);
      } else {
        delete x["progress"];
      }
      if (A.load.Bool(ptr + 114)) {
        x["isClickable"] = A.load.Bool(ptr + 108);
      } else {
        delete x["isClickable"];
      }
      if (A.load.Bool(ptr + 115)) {
        x["requireInteraction"] = A.load.Bool(ptr + 109);
      } else {
        delete x["requireInteraction"];
      }
      if (A.load.Bool(ptr + 116)) {
        x["silent"] = A.load.Bool(ptr + 110);
      } else {
        delete x["silent"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_PermissionLevel": (ref: heap.Ref<string>): number => {
      const idx = ["granted", "denied"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_Clear": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications && "clear" in WEBEXT?.notifications) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Clear": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.clear);
    },
    "call_Clear": (retPtr: Pointer, notificationId: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.clear(A.H.get<object>(notificationId));
      A.store.Ref(retPtr, _ret);
    },
    "try_Clear": (retPtr: Pointer, errPtr: Pointer, notificationId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.clear(A.H.get<object>(notificationId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Create": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications && "create" in WEBEXT?.notifications) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Create": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.create);
    },
    "call_Create": (retPtr: Pointer, notificationId: heap.Ref<object>, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["type"] = A.load.Enum(options + 0, ["basic", "image", "list", "progress"]);
      options_ffi["iconUrl"] = A.load.Ref(options + 4, undefined);
      if (A.load.Bool(options + 8 + 14)) {
        options_ffi["iconBitmap"] = {};
        if (A.load.Bool(options + 8 + 12)) {
          options_ffi["iconBitmap"]["width"] = A.load.Int32(options + 8 + 0);
        }
        if (A.load.Bool(options + 8 + 13)) {
          options_ffi["iconBitmap"]["height"] = A.load.Int32(options + 8 + 4);
        }
        options_ffi["iconBitmap"]["data"] = A.load.Ref(options + 8 + 8, undefined);
      }
      options_ffi["appIconMaskUrl"] = A.load.Ref(options + 24, undefined);
      if (A.load.Bool(options + 28 + 14)) {
        options_ffi["appIconMaskBitmap"] = {};
        if (A.load.Bool(options + 28 + 12)) {
          options_ffi["appIconMaskBitmap"]["width"] = A.load.Int32(options + 28 + 0);
        }
        if (A.load.Bool(options + 28 + 13)) {
          options_ffi["appIconMaskBitmap"]["height"] = A.load.Int32(options + 28 + 4);
        }
        options_ffi["appIconMaskBitmap"]["data"] = A.load.Ref(options + 28 + 8, undefined);
      }
      options_ffi["title"] = A.load.Ref(options + 44, undefined);
      options_ffi["message"] = A.load.Ref(options + 48, undefined);
      options_ffi["contextMessage"] = A.load.Ref(options + 52, undefined);
      if (A.load.Bool(options + 111)) {
        options_ffi["priority"] = A.load.Int32(options + 56);
      }
      if (A.load.Bool(options + 112)) {
        options_ffi["eventTime"] = A.load.Float64(options + 64);
      }
      options_ffi["buttons"] = A.load.Ref(options + 72, undefined);
      options_ffi["expandedMessage"] = A.load.Ref(options + 76, undefined);
      options_ffi["imageUrl"] = A.load.Ref(options + 80, undefined);
      if (A.load.Bool(options + 84 + 14)) {
        options_ffi["imageBitmap"] = {};
        if (A.load.Bool(options + 84 + 12)) {
          options_ffi["imageBitmap"]["width"] = A.load.Int32(options + 84 + 0);
        }
        if (A.load.Bool(options + 84 + 13)) {
          options_ffi["imageBitmap"]["height"] = A.load.Int32(options + 84 + 4);
        }
        options_ffi["imageBitmap"]["data"] = A.load.Ref(options + 84 + 8, undefined);
      }
      options_ffi["items"] = A.load.Ref(options + 100, undefined);
      if (A.load.Bool(options + 113)) {
        options_ffi["progress"] = A.load.Int32(options + 104);
      }
      if (A.load.Bool(options + 114)) {
        options_ffi["isClickable"] = A.load.Bool(options + 108);
      }
      if (A.load.Bool(options + 115)) {
        options_ffi["requireInteraction"] = A.load.Bool(options + 109);
      }
      if (A.load.Bool(options + 116)) {
        options_ffi["silent"] = A.load.Bool(options + 110);
      }

      const _ret = WEBEXT.notifications.create(A.H.get<object>(notificationId), options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Create": (
      retPtr: Pointer,
      errPtr: Pointer,
      notificationId: heap.Ref<object>,
      options: Pointer
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["type"] = A.load.Enum(options + 0, ["basic", "image", "list", "progress"]);
        options_ffi["iconUrl"] = A.load.Ref(options + 4, undefined);
        if (A.load.Bool(options + 8 + 14)) {
          options_ffi["iconBitmap"] = {};
          if (A.load.Bool(options + 8 + 12)) {
            options_ffi["iconBitmap"]["width"] = A.load.Int32(options + 8 + 0);
          }
          if (A.load.Bool(options + 8 + 13)) {
            options_ffi["iconBitmap"]["height"] = A.load.Int32(options + 8 + 4);
          }
          options_ffi["iconBitmap"]["data"] = A.load.Ref(options + 8 + 8, undefined);
        }
        options_ffi["appIconMaskUrl"] = A.load.Ref(options + 24, undefined);
        if (A.load.Bool(options + 28 + 14)) {
          options_ffi["appIconMaskBitmap"] = {};
          if (A.load.Bool(options + 28 + 12)) {
            options_ffi["appIconMaskBitmap"]["width"] = A.load.Int32(options + 28 + 0);
          }
          if (A.load.Bool(options + 28 + 13)) {
            options_ffi["appIconMaskBitmap"]["height"] = A.load.Int32(options + 28 + 4);
          }
          options_ffi["appIconMaskBitmap"]["data"] = A.load.Ref(options + 28 + 8, undefined);
        }
        options_ffi["title"] = A.load.Ref(options + 44, undefined);
        options_ffi["message"] = A.load.Ref(options + 48, undefined);
        options_ffi["contextMessage"] = A.load.Ref(options + 52, undefined);
        if (A.load.Bool(options + 111)) {
          options_ffi["priority"] = A.load.Int32(options + 56);
        }
        if (A.load.Bool(options + 112)) {
          options_ffi["eventTime"] = A.load.Float64(options + 64);
        }
        options_ffi["buttons"] = A.load.Ref(options + 72, undefined);
        options_ffi["expandedMessage"] = A.load.Ref(options + 76, undefined);
        options_ffi["imageUrl"] = A.load.Ref(options + 80, undefined);
        if (A.load.Bool(options + 84 + 14)) {
          options_ffi["imageBitmap"] = {};
          if (A.load.Bool(options + 84 + 12)) {
            options_ffi["imageBitmap"]["width"] = A.load.Int32(options + 84 + 0);
          }
          if (A.load.Bool(options + 84 + 13)) {
            options_ffi["imageBitmap"]["height"] = A.load.Int32(options + 84 + 4);
          }
          options_ffi["imageBitmap"]["data"] = A.load.Ref(options + 84 + 8, undefined);
        }
        options_ffi["items"] = A.load.Ref(options + 100, undefined);
        if (A.load.Bool(options + 113)) {
          options_ffi["progress"] = A.load.Int32(options + 104);
        }
        if (A.load.Bool(options + 114)) {
          options_ffi["isClickable"] = A.load.Bool(options + 108);
        }
        if (A.load.Bool(options + 115)) {
          options_ffi["requireInteraction"] = A.load.Bool(options + 109);
        }
        if (A.load.Bool(options + 116)) {
          options_ffi["silent"] = A.load.Bool(options + 110);
        }

        const _ret = WEBEXT.notifications.create(A.H.get<object>(notificationId), options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAll": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications && "getAll" in WEBEXT?.notifications) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.getAll);
    },
    "call_GetAll": (retPtr: Pointer): void => {
      const _ret = WEBEXT.notifications.getAll();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAll": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.getAll();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPermissionLevel": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications && "getPermissionLevel" in WEBEXT?.notifications) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPermissionLevel": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.getPermissionLevel);
    },
    "call_GetPermissionLevel": (retPtr: Pointer): void => {
      const _ret = WEBEXT.notifications.getPermissionLevel();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPermissionLevel": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.getPermissionLevel();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnButtonClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications?.onButtonClicked && "addListener" in WEBEXT?.notifications?.onButtonClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnButtonClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.onButtonClicked.addListener);
    },
    "call_OnButtonClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.onButtonClicked.addListener(A.H.get<object>(callback));
    },
    "try_OnButtonClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.onButtonClicked.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffButtonClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications?.onButtonClicked && "removeListener" in WEBEXT?.notifications?.onButtonClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffButtonClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.onButtonClicked.removeListener);
    },
    "call_OffButtonClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.onButtonClicked.removeListener(A.H.get<object>(callback));
    },
    "try_OffButtonClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.onButtonClicked.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnButtonClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications?.onButtonClicked && "hasListener" in WEBEXT?.notifications?.onButtonClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnButtonClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.onButtonClicked.hasListener);
    },
    "call_HasOnButtonClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.onButtonClicked.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnButtonClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.onButtonClicked.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications?.onClicked && "addListener" in WEBEXT?.notifications?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.onClicked.addListener);
    },
    "call_OnClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.onClicked.addListener(A.H.get<object>(callback));
    },
    "try_OnClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.onClicked.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications?.onClicked && "removeListener" in WEBEXT?.notifications?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.onClicked.removeListener);
    },
    "call_OffClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.onClicked.removeListener(A.H.get<object>(callback));
    },
    "try_OffClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.onClicked.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications?.onClicked && "hasListener" in WEBEXT?.notifications?.onClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.onClicked.hasListener);
    },
    "call_HasOnClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.onClicked.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.onClicked.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnClosed": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications?.onClosed && "addListener" in WEBEXT?.notifications?.onClosed) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClosed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.onClosed.addListener);
    },
    "call_OnClosed": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.onClosed.addListener(A.H.get<object>(callback));
    },
    "try_OnClosed": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.onClosed.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClosed": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications?.onClosed && "removeListener" in WEBEXT?.notifications?.onClosed) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClosed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.onClosed.removeListener);
    },
    "call_OffClosed": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.onClosed.removeListener(A.H.get<object>(callback));
    },
    "try_OffClosed": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.onClosed.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClosed": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications?.onClosed && "hasListener" in WEBEXT?.notifications?.onClosed) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClosed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.onClosed.hasListener);
    },
    "call_HasOnClosed": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.onClosed.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClosed": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.onClosed.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPermissionLevelChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.notifications?.onPermissionLevelChanged &&
        "addListener" in WEBEXT?.notifications?.onPermissionLevelChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPermissionLevelChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.onPermissionLevelChanged.addListener);
    },
    "call_OnPermissionLevelChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.onPermissionLevelChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnPermissionLevelChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.onPermissionLevelChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPermissionLevelChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.notifications?.onPermissionLevelChanged &&
        "removeListener" in WEBEXT?.notifications?.onPermissionLevelChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPermissionLevelChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.onPermissionLevelChanged.removeListener);
    },
    "call_OffPermissionLevelChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.onPermissionLevelChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffPermissionLevelChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.onPermissionLevelChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPermissionLevelChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.notifications?.onPermissionLevelChanged &&
        "hasListener" in WEBEXT?.notifications?.onPermissionLevelChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPermissionLevelChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.onPermissionLevelChanged.hasListener);
    },
    "call_HasOnPermissionLevelChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.onPermissionLevelChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPermissionLevelChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.onPermissionLevelChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnShowSettings": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications?.onShowSettings && "addListener" in WEBEXT?.notifications?.onShowSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnShowSettings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.onShowSettings.addListener);
    },
    "call_OnShowSettings": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.onShowSettings.addListener(A.H.get<object>(callback));
    },
    "try_OnShowSettings": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.onShowSettings.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffShowSettings": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications?.onShowSettings && "removeListener" in WEBEXT?.notifications?.onShowSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffShowSettings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.onShowSettings.removeListener);
    },
    "call_OffShowSettings": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.onShowSettings.removeListener(A.H.get<object>(callback));
    },
    "try_OffShowSettings": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.onShowSettings.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnShowSettings": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications?.onShowSettings && "hasListener" in WEBEXT?.notifications?.onShowSettings) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnShowSettings": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.onShowSettings.hasListener);
    },
    "call_HasOnShowSettings": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.notifications.onShowSettings.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnShowSettings": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.notifications.onShowSettings.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Update": (): heap.Ref<boolean> => {
      if (WEBEXT?.notifications && "update" in WEBEXT?.notifications) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Update": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.notifications.update);
    },
    "call_Update": (retPtr: Pointer, notificationId: heap.Ref<object>, options: Pointer): void => {
      const options_ffi = {};

      options_ffi["type"] = A.load.Enum(options + 0, ["basic", "image", "list", "progress"]);
      options_ffi["iconUrl"] = A.load.Ref(options + 4, undefined);
      if (A.load.Bool(options + 8 + 14)) {
        options_ffi["iconBitmap"] = {};
        if (A.load.Bool(options + 8 + 12)) {
          options_ffi["iconBitmap"]["width"] = A.load.Int32(options + 8 + 0);
        }
        if (A.load.Bool(options + 8 + 13)) {
          options_ffi["iconBitmap"]["height"] = A.load.Int32(options + 8 + 4);
        }
        options_ffi["iconBitmap"]["data"] = A.load.Ref(options + 8 + 8, undefined);
      }
      options_ffi["appIconMaskUrl"] = A.load.Ref(options + 24, undefined);
      if (A.load.Bool(options + 28 + 14)) {
        options_ffi["appIconMaskBitmap"] = {};
        if (A.load.Bool(options + 28 + 12)) {
          options_ffi["appIconMaskBitmap"]["width"] = A.load.Int32(options + 28 + 0);
        }
        if (A.load.Bool(options + 28 + 13)) {
          options_ffi["appIconMaskBitmap"]["height"] = A.load.Int32(options + 28 + 4);
        }
        options_ffi["appIconMaskBitmap"]["data"] = A.load.Ref(options + 28 + 8, undefined);
      }
      options_ffi["title"] = A.load.Ref(options + 44, undefined);
      options_ffi["message"] = A.load.Ref(options + 48, undefined);
      options_ffi["contextMessage"] = A.load.Ref(options + 52, undefined);
      if (A.load.Bool(options + 111)) {
        options_ffi["priority"] = A.load.Int32(options + 56);
      }
      if (A.load.Bool(options + 112)) {
        options_ffi["eventTime"] = A.load.Float64(options + 64);
      }
      options_ffi["buttons"] = A.load.Ref(options + 72, undefined);
      options_ffi["expandedMessage"] = A.load.Ref(options + 76, undefined);
      options_ffi["imageUrl"] = A.load.Ref(options + 80, undefined);
      if (A.load.Bool(options + 84 + 14)) {
        options_ffi["imageBitmap"] = {};
        if (A.load.Bool(options + 84 + 12)) {
          options_ffi["imageBitmap"]["width"] = A.load.Int32(options + 84 + 0);
        }
        if (A.load.Bool(options + 84 + 13)) {
          options_ffi["imageBitmap"]["height"] = A.load.Int32(options + 84 + 4);
        }
        options_ffi["imageBitmap"]["data"] = A.load.Ref(options + 84 + 8, undefined);
      }
      options_ffi["items"] = A.load.Ref(options + 100, undefined);
      if (A.load.Bool(options + 113)) {
        options_ffi["progress"] = A.load.Int32(options + 104);
      }
      if (A.load.Bool(options + 114)) {
        options_ffi["isClickable"] = A.load.Bool(options + 108);
      }
      if (A.load.Bool(options + 115)) {
        options_ffi["requireInteraction"] = A.load.Bool(options + 109);
      }
      if (A.load.Bool(options + 116)) {
        options_ffi["silent"] = A.load.Bool(options + 110);
      }

      const _ret = WEBEXT.notifications.update(A.H.get<object>(notificationId), options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_Update": (
      retPtr: Pointer,
      errPtr: Pointer,
      notificationId: heap.Ref<object>,
      options: Pointer
    ): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        options_ffi["type"] = A.load.Enum(options + 0, ["basic", "image", "list", "progress"]);
        options_ffi["iconUrl"] = A.load.Ref(options + 4, undefined);
        if (A.load.Bool(options + 8 + 14)) {
          options_ffi["iconBitmap"] = {};
          if (A.load.Bool(options + 8 + 12)) {
            options_ffi["iconBitmap"]["width"] = A.load.Int32(options + 8 + 0);
          }
          if (A.load.Bool(options + 8 + 13)) {
            options_ffi["iconBitmap"]["height"] = A.load.Int32(options + 8 + 4);
          }
          options_ffi["iconBitmap"]["data"] = A.load.Ref(options + 8 + 8, undefined);
        }
        options_ffi["appIconMaskUrl"] = A.load.Ref(options + 24, undefined);
        if (A.load.Bool(options + 28 + 14)) {
          options_ffi["appIconMaskBitmap"] = {};
          if (A.load.Bool(options + 28 + 12)) {
            options_ffi["appIconMaskBitmap"]["width"] = A.load.Int32(options + 28 + 0);
          }
          if (A.load.Bool(options + 28 + 13)) {
            options_ffi["appIconMaskBitmap"]["height"] = A.load.Int32(options + 28 + 4);
          }
          options_ffi["appIconMaskBitmap"]["data"] = A.load.Ref(options + 28 + 8, undefined);
        }
        options_ffi["title"] = A.load.Ref(options + 44, undefined);
        options_ffi["message"] = A.load.Ref(options + 48, undefined);
        options_ffi["contextMessage"] = A.load.Ref(options + 52, undefined);
        if (A.load.Bool(options + 111)) {
          options_ffi["priority"] = A.load.Int32(options + 56);
        }
        if (A.load.Bool(options + 112)) {
          options_ffi["eventTime"] = A.load.Float64(options + 64);
        }
        options_ffi["buttons"] = A.load.Ref(options + 72, undefined);
        options_ffi["expandedMessage"] = A.load.Ref(options + 76, undefined);
        options_ffi["imageUrl"] = A.load.Ref(options + 80, undefined);
        if (A.load.Bool(options + 84 + 14)) {
          options_ffi["imageBitmap"] = {};
          if (A.load.Bool(options + 84 + 12)) {
            options_ffi["imageBitmap"]["width"] = A.load.Int32(options + 84 + 0);
          }
          if (A.load.Bool(options + 84 + 13)) {
            options_ffi["imageBitmap"]["height"] = A.load.Int32(options + 84 + 4);
          }
          options_ffi["imageBitmap"]["data"] = A.load.Ref(options + 84 + 8, undefined);
        }
        options_ffi["items"] = A.load.Ref(options + 100, undefined);
        if (A.load.Bool(options + 113)) {
          options_ffi["progress"] = A.load.Int32(options + 104);
        }
        if (A.load.Bool(options + 114)) {
          options_ffi["isClickable"] = A.load.Bool(options + 108);
        }
        if (A.load.Bool(options + 115)) {
          options_ffi["requireInteraction"] = A.load.Bool(options + 109);
        }
        if (A.load.Bool(options + 116)) {
          options_ffi["silent"] = A.load.Bool(options + 110);
        }

        const _ret = WEBEXT.notifications.update(A.H.get<object>(notificationId), options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
