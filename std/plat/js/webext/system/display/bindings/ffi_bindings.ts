import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/system/display", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_ActiveState": (ref: heap.Ref<string>): number => {
      const idx = ["active", "inactive"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Bounds": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Bool(ptr + 16, "left" in x ? true : false);
        A.store.Int32(ptr + 0, x["left"] === undefined ? 0 : (x["left"] as number));
        A.store.Bool(ptr + 17, "top" in x ? true : false);
        A.store.Int32(ptr + 4, x["top"] === undefined ? 0 : (x["top"] as number));
        A.store.Bool(ptr + 18, "width" in x ? true : false);
        A.store.Int32(ptr + 8, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Bool(ptr + 19, "height" in x ? true : false);
        A.store.Int32(ptr + 12, x["height"] === undefined ? 0 : (x["height"] as number));
      }
    },
    "load_Bounds": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["left"] = A.load.Int32(ptr + 0);
      } else {
        delete x["left"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["top"] = A.load.Int32(ptr + 4);
      } else {
        delete x["top"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["width"] = A.load.Int32(ptr + 8);
      } else {
        delete x["width"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["height"] = A.load.Int32(ptr + 12);
      } else {
        delete x["height"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Edid": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Ref(ptr + 0, x["manufacturerId"]);
        A.store.Ref(ptr + 4, x["productId"]);
        A.store.Bool(ptr + 12, "yearOfManufacture" in x ? true : false);
        A.store.Int32(ptr + 8, x["yearOfManufacture"] === undefined ? 0 : (x["yearOfManufacture"] as number));
      }
    },
    "load_Edid": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["manufacturerId"] = A.load.Ref(ptr + 0, undefined);
      x["productId"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["yearOfManufacture"] = A.load.Int32(ptr + 8);
      } else {
        delete x["yearOfManufacture"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Insets": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Bool(ptr + 16, "left" in x ? true : false);
        A.store.Int32(ptr + 0, x["left"] === undefined ? 0 : (x["left"] as number));
        A.store.Bool(ptr + 17, "top" in x ? true : false);
        A.store.Int32(ptr + 4, x["top"] === undefined ? 0 : (x["top"] as number));
        A.store.Bool(ptr + 18, "right" in x ? true : false);
        A.store.Int32(ptr + 8, x["right"] === undefined ? 0 : (x["right"] as number));
        A.store.Bool(ptr + 19, "bottom" in x ? true : false);
        A.store.Int32(ptr + 12, x["bottom"] === undefined ? 0 : (x["bottom"] as number));
      }
    },
    "load_Insets": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["left"] = A.load.Int32(ptr + 0);
      } else {
        delete x["left"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["top"] = A.load.Int32(ptr + 4);
      } else {
        delete x["top"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["right"] = A.load.Int32(ptr + 8);
      } else {
        delete x["right"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["bottom"] = A.load.Int32(ptr + 12);
      } else {
        delete x["bottom"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DisplayMode": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 53, false);
        A.store.Bool(ptr + 43, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 44, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 45, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 46, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 47, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 48, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Bool(ptr + 49, false);
        A.store.Float64(ptr + 32, 0);
        A.store.Bool(ptr + 50, false);
        A.store.Bool(ptr + 40, false);
        A.store.Bool(ptr + 51, false);
        A.store.Bool(ptr + 41, false);
        A.store.Bool(ptr + 52, false);
        A.store.Bool(ptr + 42, false);
      } else {
        A.store.Bool(ptr + 53, true);
        A.store.Bool(ptr + 43, "width" in x ? true : false);
        A.store.Int32(ptr + 0, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Bool(ptr + 44, "height" in x ? true : false);
        A.store.Int32(ptr + 4, x["height"] === undefined ? 0 : (x["height"] as number));
        A.store.Bool(ptr + 45, "widthInNativePixels" in x ? true : false);
        A.store.Int32(ptr + 8, x["widthInNativePixels"] === undefined ? 0 : (x["widthInNativePixels"] as number));
        A.store.Bool(ptr + 46, "heightInNativePixels" in x ? true : false);
        A.store.Int32(ptr + 12, x["heightInNativePixels"] === undefined ? 0 : (x["heightInNativePixels"] as number));
        A.store.Bool(ptr + 47, "uiScale" in x ? true : false);
        A.store.Float64(ptr + 16, x["uiScale"] === undefined ? 0 : (x["uiScale"] as number));
        A.store.Bool(ptr + 48, "deviceScaleFactor" in x ? true : false);
        A.store.Float64(ptr + 24, x["deviceScaleFactor"] === undefined ? 0 : (x["deviceScaleFactor"] as number));
        A.store.Bool(ptr + 49, "refreshRate" in x ? true : false);
        A.store.Float64(ptr + 32, x["refreshRate"] === undefined ? 0 : (x["refreshRate"] as number));
        A.store.Bool(ptr + 50, "isNative" in x ? true : false);
        A.store.Bool(ptr + 40, x["isNative"] ? true : false);
        A.store.Bool(ptr + 51, "isSelected" in x ? true : false);
        A.store.Bool(ptr + 41, x["isSelected"] ? true : false);
        A.store.Bool(ptr + 52, "isInterlaced" in x ? true : false);
        A.store.Bool(ptr + 42, x["isInterlaced"] ? true : false);
      }
    },
    "load_DisplayMode": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 43)) {
        x["width"] = A.load.Int32(ptr + 0);
      } else {
        delete x["width"];
      }
      if (A.load.Bool(ptr + 44)) {
        x["height"] = A.load.Int32(ptr + 4);
      } else {
        delete x["height"];
      }
      if (A.load.Bool(ptr + 45)) {
        x["widthInNativePixels"] = A.load.Int32(ptr + 8);
      } else {
        delete x["widthInNativePixels"];
      }
      if (A.load.Bool(ptr + 46)) {
        x["heightInNativePixels"] = A.load.Int32(ptr + 12);
      } else {
        delete x["heightInNativePixels"];
      }
      if (A.load.Bool(ptr + 47)) {
        x["uiScale"] = A.load.Float64(ptr + 16);
      } else {
        delete x["uiScale"];
      }
      if (A.load.Bool(ptr + 48)) {
        x["deviceScaleFactor"] = A.load.Float64(ptr + 24);
      } else {
        delete x["deviceScaleFactor"];
      }
      if (A.load.Bool(ptr + 49)) {
        x["refreshRate"] = A.load.Float64(ptr + 32);
      } else {
        delete x["refreshRate"];
      }
      if (A.load.Bool(ptr + 50)) {
        x["isNative"] = A.load.Bool(ptr + 40);
      } else {
        delete x["isNative"];
      }
      if (A.load.Bool(ptr + 51)) {
        x["isSelected"] = A.load.Bool(ptr + 41);
      } else {
        delete x["isSelected"];
      }
      if (A.load.Bool(ptr + 52)) {
        x["isInterlaced"] = A.load.Bool(ptr + 42);
      } else {
        delete x["isInterlaced"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DisplayUnitInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 171, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);

        A.store.Bool(ptr + 8 + 13, false);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Ref(ptr + 8 + 4, undefined);
        A.store.Bool(ptr + 8 + 12, false);
        A.store.Int32(ptr + 8 + 8, 0);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Bool(ptr + 160, false);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 161, false);
        A.store.Bool(ptr + 33, false);
        A.store.Bool(ptr + 162, false);
        A.store.Bool(ptr + 34, false);
        A.store.Enum(ptr + 36, -1);
        A.store.Bool(ptr + 163, false);
        A.store.Bool(ptr + 40, false);
        A.store.Bool(ptr + 164, false);
        A.store.Bool(ptr + 41, false);
        A.store.Bool(ptr + 165, false);
        A.store.Float64(ptr + 48, 0);
        A.store.Bool(ptr + 166, false);
        A.store.Float64(ptr + 56, 0);
        A.store.Bool(ptr + 167, false);
        A.store.Int32(ptr + 64, 0);

        A.store.Bool(ptr + 68 + 20, false);
        A.store.Bool(ptr + 68 + 16, false);
        A.store.Int32(ptr + 68 + 0, 0);
        A.store.Bool(ptr + 68 + 17, false);
        A.store.Int32(ptr + 68 + 4, 0);
        A.store.Bool(ptr + 68 + 18, false);
        A.store.Int32(ptr + 68 + 8, 0);
        A.store.Bool(ptr + 68 + 19, false);
        A.store.Int32(ptr + 68 + 12, 0);

        A.store.Bool(ptr + 92 + 20, false);
        A.store.Bool(ptr + 92 + 16, false);
        A.store.Int32(ptr + 92 + 0, 0);
        A.store.Bool(ptr + 92 + 17, false);
        A.store.Int32(ptr + 92 + 4, 0);
        A.store.Bool(ptr + 92 + 18, false);
        A.store.Int32(ptr + 92 + 8, 0);
        A.store.Bool(ptr + 92 + 19, false);
        A.store.Int32(ptr + 92 + 12, 0);

        A.store.Bool(ptr + 116 + 20, false);
        A.store.Bool(ptr + 116 + 16, false);
        A.store.Int32(ptr + 116 + 0, 0);
        A.store.Bool(ptr + 116 + 17, false);
        A.store.Int32(ptr + 116 + 4, 0);
        A.store.Bool(ptr + 116 + 18, false);
        A.store.Int32(ptr + 116 + 8, 0);
        A.store.Bool(ptr + 116 + 19, false);
        A.store.Int32(ptr + 116 + 12, 0);
        A.store.Ref(ptr + 140, undefined);
        A.store.Bool(ptr + 168, false);
        A.store.Bool(ptr + 144, false);
        A.store.Bool(ptr + 169, false);
        A.store.Bool(ptr + 145, false);
        A.store.Ref(ptr + 148, undefined);
        A.store.Bool(ptr + 170, false);
        A.store.Float64(ptr + 152, 0);
      } else {
        A.store.Bool(ptr + 171, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["name"]);

        if (typeof x["edid"] === "undefined") {
          A.store.Bool(ptr + 8 + 13, false);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Ref(ptr + 8 + 4, undefined);
          A.store.Bool(ptr + 8 + 12, false);
          A.store.Int32(ptr + 8 + 8, 0);
        } else {
          A.store.Bool(ptr + 8 + 13, true);
          A.store.Ref(ptr + 8 + 0, x["edid"]["manufacturerId"]);
          A.store.Ref(ptr + 8 + 4, x["edid"]["productId"]);
          A.store.Bool(ptr + 8 + 12, "yearOfManufacture" in x["edid"] ? true : false);
          A.store.Int32(
            ptr + 8 + 8,
            x["edid"]["yearOfManufacture"] === undefined ? 0 : (x["edid"]["yearOfManufacture"] as number)
          );
        }
        A.store.Ref(ptr + 24, x["mirroringSourceId"]);
        A.store.Ref(ptr + 28, x["mirroringDestinationIds"]);
        A.store.Bool(ptr + 160, "isPrimary" in x ? true : false);
        A.store.Bool(ptr + 32, x["isPrimary"] ? true : false);
        A.store.Bool(ptr + 161, "isInternal" in x ? true : false);
        A.store.Bool(ptr + 33, x["isInternal"] ? true : false);
        A.store.Bool(ptr + 162, "isEnabled" in x ? true : false);
        A.store.Bool(ptr + 34, x["isEnabled"] ? true : false);
        A.store.Enum(ptr + 36, ["active", "inactive"].indexOf(x["activeState"] as string));
        A.store.Bool(ptr + 163, "isUnified" in x ? true : false);
        A.store.Bool(ptr + 40, x["isUnified"] ? true : false);
        A.store.Bool(ptr + 164, "isAutoRotationAllowed" in x ? true : false);
        A.store.Bool(ptr + 41, x["isAutoRotationAllowed"] ? true : false);
        A.store.Bool(ptr + 165, "dpiX" in x ? true : false);
        A.store.Float64(ptr + 48, x["dpiX"] === undefined ? 0 : (x["dpiX"] as number));
        A.store.Bool(ptr + 166, "dpiY" in x ? true : false);
        A.store.Float64(ptr + 56, x["dpiY"] === undefined ? 0 : (x["dpiY"] as number));
        A.store.Bool(ptr + 167, "rotation" in x ? true : false);
        A.store.Int32(ptr + 64, x["rotation"] === undefined ? 0 : (x["rotation"] as number));

        if (typeof x["bounds"] === "undefined") {
          A.store.Bool(ptr + 68 + 20, false);
          A.store.Bool(ptr + 68 + 16, false);
          A.store.Int32(ptr + 68 + 0, 0);
          A.store.Bool(ptr + 68 + 17, false);
          A.store.Int32(ptr + 68 + 4, 0);
          A.store.Bool(ptr + 68 + 18, false);
          A.store.Int32(ptr + 68 + 8, 0);
          A.store.Bool(ptr + 68 + 19, false);
          A.store.Int32(ptr + 68 + 12, 0);
        } else {
          A.store.Bool(ptr + 68 + 20, true);
          A.store.Bool(ptr + 68 + 16, "left" in x["bounds"] ? true : false);
          A.store.Int32(ptr + 68 + 0, x["bounds"]["left"] === undefined ? 0 : (x["bounds"]["left"] as number));
          A.store.Bool(ptr + 68 + 17, "top" in x["bounds"] ? true : false);
          A.store.Int32(ptr + 68 + 4, x["bounds"]["top"] === undefined ? 0 : (x["bounds"]["top"] as number));
          A.store.Bool(ptr + 68 + 18, "width" in x["bounds"] ? true : false);
          A.store.Int32(ptr + 68 + 8, x["bounds"]["width"] === undefined ? 0 : (x["bounds"]["width"] as number));
          A.store.Bool(ptr + 68 + 19, "height" in x["bounds"] ? true : false);
          A.store.Int32(ptr + 68 + 12, x["bounds"]["height"] === undefined ? 0 : (x["bounds"]["height"] as number));
        }

        if (typeof x["overscan"] === "undefined") {
          A.store.Bool(ptr + 92 + 20, false);
          A.store.Bool(ptr + 92 + 16, false);
          A.store.Int32(ptr + 92 + 0, 0);
          A.store.Bool(ptr + 92 + 17, false);
          A.store.Int32(ptr + 92 + 4, 0);
          A.store.Bool(ptr + 92 + 18, false);
          A.store.Int32(ptr + 92 + 8, 0);
          A.store.Bool(ptr + 92 + 19, false);
          A.store.Int32(ptr + 92 + 12, 0);
        } else {
          A.store.Bool(ptr + 92 + 20, true);
          A.store.Bool(ptr + 92 + 16, "left" in x["overscan"] ? true : false);
          A.store.Int32(ptr + 92 + 0, x["overscan"]["left"] === undefined ? 0 : (x["overscan"]["left"] as number));
          A.store.Bool(ptr + 92 + 17, "top" in x["overscan"] ? true : false);
          A.store.Int32(ptr + 92 + 4, x["overscan"]["top"] === undefined ? 0 : (x["overscan"]["top"] as number));
          A.store.Bool(ptr + 92 + 18, "right" in x["overscan"] ? true : false);
          A.store.Int32(ptr + 92 + 8, x["overscan"]["right"] === undefined ? 0 : (x["overscan"]["right"] as number));
          A.store.Bool(ptr + 92 + 19, "bottom" in x["overscan"] ? true : false);
          A.store.Int32(ptr + 92 + 12, x["overscan"]["bottom"] === undefined ? 0 : (x["overscan"]["bottom"] as number));
        }

        if (typeof x["workArea"] === "undefined") {
          A.store.Bool(ptr + 116 + 20, false);
          A.store.Bool(ptr + 116 + 16, false);
          A.store.Int32(ptr + 116 + 0, 0);
          A.store.Bool(ptr + 116 + 17, false);
          A.store.Int32(ptr + 116 + 4, 0);
          A.store.Bool(ptr + 116 + 18, false);
          A.store.Int32(ptr + 116 + 8, 0);
          A.store.Bool(ptr + 116 + 19, false);
          A.store.Int32(ptr + 116 + 12, 0);
        } else {
          A.store.Bool(ptr + 116 + 20, true);
          A.store.Bool(ptr + 116 + 16, "left" in x["workArea"] ? true : false);
          A.store.Int32(ptr + 116 + 0, x["workArea"]["left"] === undefined ? 0 : (x["workArea"]["left"] as number));
          A.store.Bool(ptr + 116 + 17, "top" in x["workArea"] ? true : false);
          A.store.Int32(ptr + 116 + 4, x["workArea"]["top"] === undefined ? 0 : (x["workArea"]["top"] as number));
          A.store.Bool(ptr + 116 + 18, "width" in x["workArea"] ? true : false);
          A.store.Int32(ptr + 116 + 8, x["workArea"]["width"] === undefined ? 0 : (x["workArea"]["width"] as number));
          A.store.Bool(ptr + 116 + 19, "height" in x["workArea"] ? true : false);
          A.store.Int32(
            ptr + 116 + 12,
            x["workArea"]["height"] === undefined ? 0 : (x["workArea"]["height"] as number)
          );
        }
        A.store.Ref(ptr + 140, x["modes"]);
        A.store.Bool(ptr + 168, "hasTouchSupport" in x ? true : false);
        A.store.Bool(ptr + 144, x["hasTouchSupport"] ? true : false);
        A.store.Bool(ptr + 169, "hasAccelerometerSupport" in x ? true : false);
        A.store.Bool(ptr + 145, x["hasAccelerometerSupport"] ? true : false);
        A.store.Ref(ptr + 148, x["availableDisplayZoomFactors"]);
        A.store.Bool(ptr + 170, "displayZoomFactor" in x ? true : false);
        A.store.Float64(ptr + 152, x["displayZoomFactor"] === undefined ? 0 : (x["displayZoomFactor"] as number));
      }
    },
    "load_DisplayUnitInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 8 + 13)) {
        x["edid"] = {};
        x["edid"]["manufacturerId"] = A.load.Ref(ptr + 8 + 0, undefined);
        x["edid"]["productId"] = A.load.Ref(ptr + 8 + 4, undefined);
        if (A.load.Bool(ptr + 8 + 12)) {
          x["edid"]["yearOfManufacture"] = A.load.Int32(ptr + 8 + 8);
        } else {
          delete x["edid"]["yearOfManufacture"];
        }
      } else {
        delete x["edid"];
      }
      x["mirroringSourceId"] = A.load.Ref(ptr + 24, undefined);
      x["mirroringDestinationIds"] = A.load.Ref(ptr + 28, undefined);
      if (A.load.Bool(ptr + 160)) {
        x["isPrimary"] = A.load.Bool(ptr + 32);
      } else {
        delete x["isPrimary"];
      }
      if (A.load.Bool(ptr + 161)) {
        x["isInternal"] = A.load.Bool(ptr + 33);
      } else {
        delete x["isInternal"];
      }
      if (A.load.Bool(ptr + 162)) {
        x["isEnabled"] = A.load.Bool(ptr + 34);
      } else {
        delete x["isEnabled"];
      }
      x["activeState"] = A.load.Enum(ptr + 36, ["active", "inactive"]);
      if (A.load.Bool(ptr + 163)) {
        x["isUnified"] = A.load.Bool(ptr + 40);
      } else {
        delete x["isUnified"];
      }
      if (A.load.Bool(ptr + 164)) {
        x["isAutoRotationAllowed"] = A.load.Bool(ptr + 41);
      } else {
        delete x["isAutoRotationAllowed"];
      }
      if (A.load.Bool(ptr + 165)) {
        x["dpiX"] = A.load.Float64(ptr + 48);
      } else {
        delete x["dpiX"];
      }
      if (A.load.Bool(ptr + 166)) {
        x["dpiY"] = A.load.Float64(ptr + 56);
      } else {
        delete x["dpiY"];
      }
      if (A.load.Bool(ptr + 167)) {
        x["rotation"] = A.load.Int32(ptr + 64);
      } else {
        delete x["rotation"];
      }
      if (A.load.Bool(ptr + 68 + 20)) {
        x["bounds"] = {};
        if (A.load.Bool(ptr + 68 + 16)) {
          x["bounds"]["left"] = A.load.Int32(ptr + 68 + 0);
        } else {
          delete x["bounds"]["left"];
        }
        if (A.load.Bool(ptr + 68 + 17)) {
          x["bounds"]["top"] = A.load.Int32(ptr + 68 + 4);
        } else {
          delete x["bounds"]["top"];
        }
        if (A.load.Bool(ptr + 68 + 18)) {
          x["bounds"]["width"] = A.load.Int32(ptr + 68 + 8);
        } else {
          delete x["bounds"]["width"];
        }
        if (A.load.Bool(ptr + 68 + 19)) {
          x["bounds"]["height"] = A.load.Int32(ptr + 68 + 12);
        } else {
          delete x["bounds"]["height"];
        }
      } else {
        delete x["bounds"];
      }
      if (A.load.Bool(ptr + 92 + 20)) {
        x["overscan"] = {};
        if (A.load.Bool(ptr + 92 + 16)) {
          x["overscan"]["left"] = A.load.Int32(ptr + 92 + 0);
        } else {
          delete x["overscan"]["left"];
        }
        if (A.load.Bool(ptr + 92 + 17)) {
          x["overscan"]["top"] = A.load.Int32(ptr + 92 + 4);
        } else {
          delete x["overscan"]["top"];
        }
        if (A.load.Bool(ptr + 92 + 18)) {
          x["overscan"]["right"] = A.load.Int32(ptr + 92 + 8);
        } else {
          delete x["overscan"]["right"];
        }
        if (A.load.Bool(ptr + 92 + 19)) {
          x["overscan"]["bottom"] = A.load.Int32(ptr + 92 + 12);
        } else {
          delete x["overscan"]["bottom"];
        }
      } else {
        delete x["overscan"];
      }
      if (A.load.Bool(ptr + 116 + 20)) {
        x["workArea"] = {};
        if (A.load.Bool(ptr + 116 + 16)) {
          x["workArea"]["left"] = A.load.Int32(ptr + 116 + 0);
        } else {
          delete x["workArea"]["left"];
        }
        if (A.load.Bool(ptr + 116 + 17)) {
          x["workArea"]["top"] = A.load.Int32(ptr + 116 + 4);
        } else {
          delete x["workArea"]["top"];
        }
        if (A.load.Bool(ptr + 116 + 18)) {
          x["workArea"]["width"] = A.load.Int32(ptr + 116 + 8);
        } else {
          delete x["workArea"]["width"];
        }
        if (A.load.Bool(ptr + 116 + 19)) {
          x["workArea"]["height"] = A.load.Int32(ptr + 116 + 12);
        } else {
          delete x["workArea"]["height"];
        }
      } else {
        delete x["workArea"];
      }
      x["modes"] = A.load.Ref(ptr + 140, undefined);
      if (A.load.Bool(ptr + 168)) {
        x["hasTouchSupport"] = A.load.Bool(ptr + 144);
      } else {
        delete x["hasTouchSupport"];
      }
      if (A.load.Bool(ptr + 169)) {
        x["hasAccelerometerSupport"] = A.load.Bool(ptr + 145);
      } else {
        delete x["hasAccelerometerSupport"];
      }
      x["availableDisplayZoomFactors"] = A.load.Ref(ptr + 148, undefined);
      if (A.load.Bool(ptr + 170)) {
        x["displayZoomFactor"] = A.load.Float64(ptr + 152);
      } else {
        delete x["displayZoomFactor"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_LayoutPosition": (ref: heap.Ref<string>): number => {
      const idx = ["top", "right", "bottom", "left"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_DisplayLayout": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 12, 0);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["parentId"]);
        A.store.Enum(ptr + 8, ["top", "right", "bottom", "left"].indexOf(x["position"] as string));
        A.store.Bool(ptr + 16, "offset" in x ? true : false);
        A.store.Int32(ptr + 12, x["offset"] === undefined ? 0 : (x["offset"] as number));
      }
    },
    "load_DisplayLayout": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["parentId"] = A.load.Ref(ptr + 4, undefined);
      x["position"] = A.load.Enum(ptr + 8, ["top", "right", "bottom", "left"]);
      if (A.load.Bool(ptr + 16)) {
        x["offset"] = A.load.Int32(ptr + 12);
      } else {
        delete x["offset"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DisplayProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 118, false);
        A.store.Bool(ptr + 112, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 113, false);
        A.store.Bool(ptr + 8, false);

        A.store.Bool(ptr + 12 + 20, false);
        A.store.Bool(ptr + 12 + 16, false);
        A.store.Int32(ptr + 12 + 0, 0);
        A.store.Bool(ptr + 12 + 17, false);
        A.store.Int32(ptr + 12 + 4, 0);
        A.store.Bool(ptr + 12 + 18, false);
        A.store.Int32(ptr + 12 + 8, 0);
        A.store.Bool(ptr + 12 + 19, false);
        A.store.Int32(ptr + 12 + 12, 0);
        A.store.Bool(ptr + 114, false);
        A.store.Int32(ptr + 36, 0);
        A.store.Bool(ptr + 115, false);
        A.store.Int32(ptr + 40, 0);
        A.store.Bool(ptr + 116, false);
        A.store.Int32(ptr + 44, 0);

        A.store.Bool(ptr + 48 + 53, false);
        A.store.Bool(ptr + 48 + 43, false);
        A.store.Int32(ptr + 48 + 0, 0);
        A.store.Bool(ptr + 48 + 44, false);
        A.store.Int32(ptr + 48 + 4, 0);
        A.store.Bool(ptr + 48 + 45, false);
        A.store.Int32(ptr + 48 + 8, 0);
        A.store.Bool(ptr + 48 + 46, false);
        A.store.Int32(ptr + 48 + 12, 0);
        A.store.Bool(ptr + 48 + 47, false);
        A.store.Float64(ptr + 48 + 16, 0);
        A.store.Bool(ptr + 48 + 48, false);
        A.store.Float64(ptr + 48 + 24, 0);
        A.store.Bool(ptr + 48 + 49, false);
        A.store.Float64(ptr + 48 + 32, 0);
        A.store.Bool(ptr + 48 + 50, false);
        A.store.Bool(ptr + 48 + 40, false);
        A.store.Bool(ptr + 48 + 51, false);
        A.store.Bool(ptr + 48 + 41, false);
        A.store.Bool(ptr + 48 + 52, false);
        A.store.Bool(ptr + 48 + 42, false);
        A.store.Bool(ptr + 117, false);
        A.store.Float64(ptr + 104, 0);
      } else {
        A.store.Bool(ptr + 118, true);
        A.store.Bool(ptr + 112, "isUnified" in x ? true : false);
        A.store.Bool(ptr + 0, x["isUnified"] ? true : false);
        A.store.Ref(ptr + 4, x["mirroringSourceId"]);
        A.store.Bool(ptr + 113, "isPrimary" in x ? true : false);
        A.store.Bool(ptr + 8, x["isPrimary"] ? true : false);

        if (typeof x["overscan"] === "undefined") {
          A.store.Bool(ptr + 12 + 20, false);
          A.store.Bool(ptr + 12 + 16, false);
          A.store.Int32(ptr + 12 + 0, 0);
          A.store.Bool(ptr + 12 + 17, false);
          A.store.Int32(ptr + 12 + 4, 0);
          A.store.Bool(ptr + 12 + 18, false);
          A.store.Int32(ptr + 12 + 8, 0);
          A.store.Bool(ptr + 12 + 19, false);
          A.store.Int32(ptr + 12 + 12, 0);
        } else {
          A.store.Bool(ptr + 12 + 20, true);
          A.store.Bool(ptr + 12 + 16, "left" in x["overscan"] ? true : false);
          A.store.Int32(ptr + 12 + 0, x["overscan"]["left"] === undefined ? 0 : (x["overscan"]["left"] as number));
          A.store.Bool(ptr + 12 + 17, "top" in x["overscan"] ? true : false);
          A.store.Int32(ptr + 12 + 4, x["overscan"]["top"] === undefined ? 0 : (x["overscan"]["top"] as number));
          A.store.Bool(ptr + 12 + 18, "right" in x["overscan"] ? true : false);
          A.store.Int32(ptr + 12 + 8, x["overscan"]["right"] === undefined ? 0 : (x["overscan"]["right"] as number));
          A.store.Bool(ptr + 12 + 19, "bottom" in x["overscan"] ? true : false);
          A.store.Int32(ptr + 12 + 12, x["overscan"]["bottom"] === undefined ? 0 : (x["overscan"]["bottom"] as number));
        }
        A.store.Bool(ptr + 114, "rotation" in x ? true : false);
        A.store.Int32(ptr + 36, x["rotation"] === undefined ? 0 : (x["rotation"] as number));
        A.store.Bool(ptr + 115, "boundsOriginX" in x ? true : false);
        A.store.Int32(ptr + 40, x["boundsOriginX"] === undefined ? 0 : (x["boundsOriginX"] as number));
        A.store.Bool(ptr + 116, "boundsOriginY" in x ? true : false);
        A.store.Int32(ptr + 44, x["boundsOriginY"] === undefined ? 0 : (x["boundsOriginY"] as number));

        if (typeof x["displayMode"] === "undefined") {
          A.store.Bool(ptr + 48 + 53, false);
          A.store.Bool(ptr + 48 + 43, false);
          A.store.Int32(ptr + 48 + 0, 0);
          A.store.Bool(ptr + 48 + 44, false);
          A.store.Int32(ptr + 48 + 4, 0);
          A.store.Bool(ptr + 48 + 45, false);
          A.store.Int32(ptr + 48 + 8, 0);
          A.store.Bool(ptr + 48 + 46, false);
          A.store.Int32(ptr + 48 + 12, 0);
          A.store.Bool(ptr + 48 + 47, false);
          A.store.Float64(ptr + 48 + 16, 0);
          A.store.Bool(ptr + 48 + 48, false);
          A.store.Float64(ptr + 48 + 24, 0);
          A.store.Bool(ptr + 48 + 49, false);
          A.store.Float64(ptr + 48 + 32, 0);
          A.store.Bool(ptr + 48 + 50, false);
          A.store.Bool(ptr + 48 + 40, false);
          A.store.Bool(ptr + 48 + 51, false);
          A.store.Bool(ptr + 48 + 41, false);
          A.store.Bool(ptr + 48 + 52, false);
          A.store.Bool(ptr + 48 + 42, false);
        } else {
          A.store.Bool(ptr + 48 + 53, true);
          A.store.Bool(ptr + 48 + 43, "width" in x["displayMode"] ? true : false);
          A.store.Int32(
            ptr + 48 + 0,
            x["displayMode"]["width"] === undefined ? 0 : (x["displayMode"]["width"] as number)
          );
          A.store.Bool(ptr + 48 + 44, "height" in x["displayMode"] ? true : false);
          A.store.Int32(
            ptr + 48 + 4,
            x["displayMode"]["height"] === undefined ? 0 : (x["displayMode"]["height"] as number)
          );
          A.store.Bool(ptr + 48 + 45, "widthInNativePixels" in x["displayMode"] ? true : false);
          A.store.Int32(
            ptr + 48 + 8,
            x["displayMode"]["widthInNativePixels"] === undefined
              ? 0
              : (x["displayMode"]["widthInNativePixels"] as number)
          );
          A.store.Bool(ptr + 48 + 46, "heightInNativePixels" in x["displayMode"] ? true : false);
          A.store.Int32(
            ptr + 48 + 12,
            x["displayMode"]["heightInNativePixels"] === undefined
              ? 0
              : (x["displayMode"]["heightInNativePixels"] as number)
          );
          A.store.Bool(ptr + 48 + 47, "uiScale" in x["displayMode"] ? true : false);
          A.store.Float64(
            ptr + 48 + 16,
            x["displayMode"]["uiScale"] === undefined ? 0 : (x["displayMode"]["uiScale"] as number)
          );
          A.store.Bool(ptr + 48 + 48, "deviceScaleFactor" in x["displayMode"] ? true : false);
          A.store.Float64(
            ptr + 48 + 24,
            x["displayMode"]["deviceScaleFactor"] === undefined ? 0 : (x["displayMode"]["deviceScaleFactor"] as number)
          );
          A.store.Bool(ptr + 48 + 49, "refreshRate" in x["displayMode"] ? true : false);
          A.store.Float64(
            ptr + 48 + 32,
            x["displayMode"]["refreshRate"] === undefined ? 0 : (x["displayMode"]["refreshRate"] as number)
          );
          A.store.Bool(ptr + 48 + 50, "isNative" in x["displayMode"] ? true : false);
          A.store.Bool(ptr + 48 + 40, x["displayMode"]["isNative"] ? true : false);
          A.store.Bool(ptr + 48 + 51, "isSelected" in x["displayMode"] ? true : false);
          A.store.Bool(ptr + 48 + 41, x["displayMode"]["isSelected"] ? true : false);
          A.store.Bool(ptr + 48 + 52, "isInterlaced" in x["displayMode"] ? true : false);
          A.store.Bool(ptr + 48 + 42, x["displayMode"]["isInterlaced"] ? true : false);
        }
        A.store.Bool(ptr + 117, "displayZoomFactor" in x ? true : false);
        A.store.Float64(ptr + 104, x["displayZoomFactor"] === undefined ? 0 : (x["displayZoomFactor"] as number));
      }
    },
    "load_DisplayProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 112)) {
        x["isUnified"] = A.load.Bool(ptr + 0);
      } else {
        delete x["isUnified"];
      }
      x["mirroringSourceId"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 113)) {
        x["isPrimary"] = A.load.Bool(ptr + 8);
      } else {
        delete x["isPrimary"];
      }
      if (A.load.Bool(ptr + 12 + 20)) {
        x["overscan"] = {};
        if (A.load.Bool(ptr + 12 + 16)) {
          x["overscan"]["left"] = A.load.Int32(ptr + 12 + 0);
        } else {
          delete x["overscan"]["left"];
        }
        if (A.load.Bool(ptr + 12 + 17)) {
          x["overscan"]["top"] = A.load.Int32(ptr + 12 + 4);
        } else {
          delete x["overscan"]["top"];
        }
        if (A.load.Bool(ptr + 12 + 18)) {
          x["overscan"]["right"] = A.load.Int32(ptr + 12 + 8);
        } else {
          delete x["overscan"]["right"];
        }
        if (A.load.Bool(ptr + 12 + 19)) {
          x["overscan"]["bottom"] = A.load.Int32(ptr + 12 + 12);
        } else {
          delete x["overscan"]["bottom"];
        }
      } else {
        delete x["overscan"];
      }
      if (A.load.Bool(ptr + 114)) {
        x["rotation"] = A.load.Int32(ptr + 36);
      } else {
        delete x["rotation"];
      }
      if (A.load.Bool(ptr + 115)) {
        x["boundsOriginX"] = A.load.Int32(ptr + 40);
      } else {
        delete x["boundsOriginX"];
      }
      if (A.load.Bool(ptr + 116)) {
        x["boundsOriginY"] = A.load.Int32(ptr + 44);
      } else {
        delete x["boundsOriginY"];
      }
      if (A.load.Bool(ptr + 48 + 53)) {
        x["displayMode"] = {};
        if (A.load.Bool(ptr + 48 + 43)) {
          x["displayMode"]["width"] = A.load.Int32(ptr + 48 + 0);
        } else {
          delete x["displayMode"]["width"];
        }
        if (A.load.Bool(ptr + 48 + 44)) {
          x["displayMode"]["height"] = A.load.Int32(ptr + 48 + 4);
        } else {
          delete x["displayMode"]["height"];
        }
        if (A.load.Bool(ptr + 48 + 45)) {
          x["displayMode"]["widthInNativePixels"] = A.load.Int32(ptr + 48 + 8);
        } else {
          delete x["displayMode"]["widthInNativePixels"];
        }
        if (A.load.Bool(ptr + 48 + 46)) {
          x["displayMode"]["heightInNativePixels"] = A.load.Int32(ptr + 48 + 12);
        } else {
          delete x["displayMode"]["heightInNativePixels"];
        }
        if (A.load.Bool(ptr + 48 + 47)) {
          x["displayMode"]["uiScale"] = A.load.Float64(ptr + 48 + 16);
        } else {
          delete x["displayMode"]["uiScale"];
        }
        if (A.load.Bool(ptr + 48 + 48)) {
          x["displayMode"]["deviceScaleFactor"] = A.load.Float64(ptr + 48 + 24);
        } else {
          delete x["displayMode"]["deviceScaleFactor"];
        }
        if (A.load.Bool(ptr + 48 + 49)) {
          x["displayMode"]["refreshRate"] = A.load.Float64(ptr + 48 + 32);
        } else {
          delete x["displayMode"]["refreshRate"];
        }
        if (A.load.Bool(ptr + 48 + 50)) {
          x["displayMode"]["isNative"] = A.load.Bool(ptr + 48 + 40);
        } else {
          delete x["displayMode"]["isNative"];
        }
        if (A.load.Bool(ptr + 48 + 51)) {
          x["displayMode"]["isSelected"] = A.load.Bool(ptr + 48 + 41);
        } else {
          delete x["displayMode"]["isSelected"];
        }
        if (A.load.Bool(ptr + 48 + 52)) {
          x["displayMode"]["isInterlaced"] = A.load.Bool(ptr + 48 + 42);
        } else {
          delete x["displayMode"]["isInterlaced"];
        }
      } else {
        delete x["displayMode"];
      }
      if (A.load.Bool(ptr + 117)) {
        x["displayZoomFactor"] = A.load.Float64(ptr + 104);
      } else {
        delete x["displayZoomFactor"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetInfoFlags": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 1, "singleUnified" in x ? true : false);
        A.store.Bool(ptr + 0, x["singleUnified"] ? true : false);
      }
    },
    "load_GetInfoFlags": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 1)) {
        x["singleUnified"] = A.load.Bool(ptr + 0);
      } else {
        delete x["singleUnified"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_MirrorMode": (ref: heap.Ref<string>): number => {
      const idx = ["off", "normal", "mixed"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_MirrorModeInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Enum(ptr + 0, ["off", "normal", "mixed"].indexOf(x["mode"] as string));
        A.store.Ref(ptr + 4, x["mirroringSourceId"]);
        A.store.Ref(ptr + 8, x["mirroringDestinationIds"]);
      }
    },
    "load_MirrorModeInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["mode"] = A.load.Enum(ptr + 0, ["off", "normal", "mixed"]);
      x["mirroringSourceId"] = A.load.Ref(ptr + 4, undefined);
      x["mirroringDestinationIds"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Point": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 9, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 10, true);
        A.store.Bool(ptr + 8, "x" in x ? true : false);
        A.store.Int32(ptr + 0, x["x"] === undefined ? 0 : (x["x"] as number));
        A.store.Bool(ptr + 9, "y" in x ? true : false);
        A.store.Int32(ptr + 4, x["y"] === undefined ? 0 : (x["y"] as number));
      }
    },
    "load_Point": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["x"] = A.load.Int32(ptr + 0);
      } else {
        delete x["x"];
      }
      if (A.load.Bool(ptr + 9)) {
        x["y"] = A.load.Int32(ptr + 4);
      } else {
        delete x["y"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TouchCalibrationPair": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 23, false);

        A.store.Bool(ptr + 0 + 10, false);
        A.store.Bool(ptr + 0 + 8, false);
        A.store.Int32(ptr + 0 + 0, 0);
        A.store.Bool(ptr + 0 + 9, false);
        A.store.Int32(ptr + 0 + 4, 0);

        A.store.Bool(ptr + 12 + 10, false);
        A.store.Bool(ptr + 12 + 8, false);
        A.store.Int32(ptr + 12 + 0, 0);
        A.store.Bool(ptr + 12 + 9, false);
        A.store.Int32(ptr + 12 + 4, 0);
      } else {
        A.store.Bool(ptr + 23, true);

        if (typeof x["displayPoint"] === "undefined") {
          A.store.Bool(ptr + 0 + 10, false);
          A.store.Bool(ptr + 0 + 8, false);
          A.store.Int32(ptr + 0 + 0, 0);
          A.store.Bool(ptr + 0 + 9, false);
          A.store.Int32(ptr + 0 + 4, 0);
        } else {
          A.store.Bool(ptr + 0 + 10, true);
          A.store.Bool(ptr + 0 + 8, "x" in x["displayPoint"] ? true : false);
          A.store.Int32(ptr + 0 + 0, x["displayPoint"]["x"] === undefined ? 0 : (x["displayPoint"]["x"] as number));
          A.store.Bool(ptr + 0 + 9, "y" in x["displayPoint"] ? true : false);
          A.store.Int32(ptr + 0 + 4, x["displayPoint"]["y"] === undefined ? 0 : (x["displayPoint"]["y"] as number));
        }

        if (typeof x["touchPoint"] === "undefined") {
          A.store.Bool(ptr + 12 + 10, false);
          A.store.Bool(ptr + 12 + 8, false);
          A.store.Int32(ptr + 12 + 0, 0);
          A.store.Bool(ptr + 12 + 9, false);
          A.store.Int32(ptr + 12 + 4, 0);
        } else {
          A.store.Bool(ptr + 12 + 10, true);
          A.store.Bool(ptr + 12 + 8, "x" in x["touchPoint"] ? true : false);
          A.store.Int32(ptr + 12 + 0, x["touchPoint"]["x"] === undefined ? 0 : (x["touchPoint"]["x"] as number));
          A.store.Bool(ptr + 12 + 9, "y" in x["touchPoint"] ? true : false);
          A.store.Int32(ptr + 12 + 4, x["touchPoint"]["y"] === undefined ? 0 : (x["touchPoint"]["y"] as number));
        }
      }
    },
    "load_TouchCalibrationPair": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 10)) {
        x["displayPoint"] = {};
        if (A.load.Bool(ptr + 0 + 8)) {
          x["displayPoint"]["x"] = A.load.Int32(ptr + 0 + 0);
        } else {
          delete x["displayPoint"]["x"];
        }
        if (A.load.Bool(ptr + 0 + 9)) {
          x["displayPoint"]["y"] = A.load.Int32(ptr + 0 + 4);
        } else {
          delete x["displayPoint"]["y"];
        }
      } else {
        delete x["displayPoint"];
      }
      if (A.load.Bool(ptr + 12 + 10)) {
        x["touchPoint"] = {};
        if (A.load.Bool(ptr + 12 + 8)) {
          x["touchPoint"]["x"] = A.load.Int32(ptr + 12 + 0);
        } else {
          delete x["touchPoint"]["x"];
        }
        if (A.load.Bool(ptr + 12 + 9)) {
          x["touchPoint"]["y"] = A.load.Int32(ptr + 12 + 4);
        } else {
          delete x["touchPoint"]["y"];
        }
      } else {
        delete x["touchPoint"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_TouchCalibrationPairQuad": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 96, false);

        A.store.Bool(ptr + 0 + 23, false);

        A.store.Bool(ptr + 0 + 0 + 10, false);
        A.store.Bool(ptr + 0 + 0 + 8, false);
        A.store.Int32(ptr + 0 + 0 + 0, 0);
        A.store.Bool(ptr + 0 + 0 + 9, false);
        A.store.Int32(ptr + 0 + 0 + 4, 0);

        A.store.Bool(ptr + 0 + 12 + 10, false);
        A.store.Bool(ptr + 0 + 12 + 8, false);
        A.store.Int32(ptr + 0 + 12 + 0, 0);
        A.store.Bool(ptr + 0 + 12 + 9, false);
        A.store.Int32(ptr + 0 + 12 + 4, 0);

        A.store.Bool(ptr + 24 + 23, false);

        A.store.Bool(ptr + 24 + 0 + 10, false);
        A.store.Bool(ptr + 24 + 0 + 8, false);
        A.store.Int32(ptr + 24 + 0 + 0, 0);
        A.store.Bool(ptr + 24 + 0 + 9, false);
        A.store.Int32(ptr + 24 + 0 + 4, 0);

        A.store.Bool(ptr + 24 + 12 + 10, false);
        A.store.Bool(ptr + 24 + 12 + 8, false);
        A.store.Int32(ptr + 24 + 12 + 0, 0);
        A.store.Bool(ptr + 24 + 12 + 9, false);
        A.store.Int32(ptr + 24 + 12 + 4, 0);

        A.store.Bool(ptr + 48 + 23, false);

        A.store.Bool(ptr + 48 + 0 + 10, false);
        A.store.Bool(ptr + 48 + 0 + 8, false);
        A.store.Int32(ptr + 48 + 0 + 0, 0);
        A.store.Bool(ptr + 48 + 0 + 9, false);
        A.store.Int32(ptr + 48 + 0 + 4, 0);

        A.store.Bool(ptr + 48 + 12 + 10, false);
        A.store.Bool(ptr + 48 + 12 + 8, false);
        A.store.Int32(ptr + 48 + 12 + 0, 0);
        A.store.Bool(ptr + 48 + 12 + 9, false);
        A.store.Int32(ptr + 48 + 12 + 4, 0);

        A.store.Bool(ptr + 72 + 23, false);

        A.store.Bool(ptr + 72 + 0 + 10, false);
        A.store.Bool(ptr + 72 + 0 + 8, false);
        A.store.Int32(ptr + 72 + 0 + 0, 0);
        A.store.Bool(ptr + 72 + 0 + 9, false);
        A.store.Int32(ptr + 72 + 0 + 4, 0);

        A.store.Bool(ptr + 72 + 12 + 10, false);
        A.store.Bool(ptr + 72 + 12 + 8, false);
        A.store.Int32(ptr + 72 + 12 + 0, 0);
        A.store.Bool(ptr + 72 + 12 + 9, false);
        A.store.Int32(ptr + 72 + 12 + 4, 0);
      } else {
        A.store.Bool(ptr + 96, true);

        if (typeof x["pair1"] === "undefined") {
          A.store.Bool(ptr + 0 + 23, false);

          A.store.Bool(ptr + 0 + 0 + 10, false);
          A.store.Bool(ptr + 0 + 0 + 8, false);
          A.store.Int32(ptr + 0 + 0 + 0, 0);
          A.store.Bool(ptr + 0 + 0 + 9, false);
          A.store.Int32(ptr + 0 + 0 + 4, 0);

          A.store.Bool(ptr + 0 + 12 + 10, false);
          A.store.Bool(ptr + 0 + 12 + 8, false);
          A.store.Int32(ptr + 0 + 12 + 0, 0);
          A.store.Bool(ptr + 0 + 12 + 9, false);
          A.store.Int32(ptr + 0 + 12 + 4, 0);
        } else {
          A.store.Bool(ptr + 0 + 23, true);

          if (typeof x["pair1"]["displayPoint"] === "undefined") {
            A.store.Bool(ptr + 0 + 0 + 10, false);
            A.store.Bool(ptr + 0 + 0 + 8, false);
            A.store.Int32(ptr + 0 + 0 + 0, 0);
            A.store.Bool(ptr + 0 + 0 + 9, false);
            A.store.Int32(ptr + 0 + 0 + 4, 0);
          } else {
            A.store.Bool(ptr + 0 + 0 + 10, true);
            A.store.Bool(ptr + 0 + 0 + 8, "x" in x["pair1"]["displayPoint"] ? true : false);
            A.store.Int32(
              ptr + 0 + 0 + 0,
              x["pair1"]["displayPoint"]["x"] === undefined ? 0 : (x["pair1"]["displayPoint"]["x"] as number)
            );
            A.store.Bool(ptr + 0 + 0 + 9, "y" in x["pair1"]["displayPoint"] ? true : false);
            A.store.Int32(
              ptr + 0 + 0 + 4,
              x["pair1"]["displayPoint"]["y"] === undefined ? 0 : (x["pair1"]["displayPoint"]["y"] as number)
            );
          }

          if (typeof x["pair1"]["touchPoint"] === "undefined") {
            A.store.Bool(ptr + 0 + 12 + 10, false);
            A.store.Bool(ptr + 0 + 12 + 8, false);
            A.store.Int32(ptr + 0 + 12 + 0, 0);
            A.store.Bool(ptr + 0 + 12 + 9, false);
            A.store.Int32(ptr + 0 + 12 + 4, 0);
          } else {
            A.store.Bool(ptr + 0 + 12 + 10, true);
            A.store.Bool(ptr + 0 + 12 + 8, "x" in x["pair1"]["touchPoint"] ? true : false);
            A.store.Int32(
              ptr + 0 + 12 + 0,
              x["pair1"]["touchPoint"]["x"] === undefined ? 0 : (x["pair1"]["touchPoint"]["x"] as number)
            );
            A.store.Bool(ptr + 0 + 12 + 9, "y" in x["pair1"]["touchPoint"] ? true : false);
            A.store.Int32(
              ptr + 0 + 12 + 4,
              x["pair1"]["touchPoint"]["y"] === undefined ? 0 : (x["pair1"]["touchPoint"]["y"] as number)
            );
          }
        }

        if (typeof x["pair2"] === "undefined") {
          A.store.Bool(ptr + 24 + 23, false);

          A.store.Bool(ptr + 24 + 0 + 10, false);
          A.store.Bool(ptr + 24 + 0 + 8, false);
          A.store.Int32(ptr + 24 + 0 + 0, 0);
          A.store.Bool(ptr + 24 + 0 + 9, false);
          A.store.Int32(ptr + 24 + 0 + 4, 0);

          A.store.Bool(ptr + 24 + 12 + 10, false);
          A.store.Bool(ptr + 24 + 12 + 8, false);
          A.store.Int32(ptr + 24 + 12 + 0, 0);
          A.store.Bool(ptr + 24 + 12 + 9, false);
          A.store.Int32(ptr + 24 + 12 + 4, 0);
        } else {
          A.store.Bool(ptr + 24 + 23, true);

          if (typeof x["pair2"]["displayPoint"] === "undefined") {
            A.store.Bool(ptr + 24 + 0 + 10, false);
            A.store.Bool(ptr + 24 + 0 + 8, false);
            A.store.Int32(ptr + 24 + 0 + 0, 0);
            A.store.Bool(ptr + 24 + 0 + 9, false);
            A.store.Int32(ptr + 24 + 0 + 4, 0);
          } else {
            A.store.Bool(ptr + 24 + 0 + 10, true);
            A.store.Bool(ptr + 24 + 0 + 8, "x" in x["pair2"]["displayPoint"] ? true : false);
            A.store.Int32(
              ptr + 24 + 0 + 0,
              x["pair2"]["displayPoint"]["x"] === undefined ? 0 : (x["pair2"]["displayPoint"]["x"] as number)
            );
            A.store.Bool(ptr + 24 + 0 + 9, "y" in x["pair2"]["displayPoint"] ? true : false);
            A.store.Int32(
              ptr + 24 + 0 + 4,
              x["pair2"]["displayPoint"]["y"] === undefined ? 0 : (x["pair2"]["displayPoint"]["y"] as number)
            );
          }

          if (typeof x["pair2"]["touchPoint"] === "undefined") {
            A.store.Bool(ptr + 24 + 12 + 10, false);
            A.store.Bool(ptr + 24 + 12 + 8, false);
            A.store.Int32(ptr + 24 + 12 + 0, 0);
            A.store.Bool(ptr + 24 + 12 + 9, false);
            A.store.Int32(ptr + 24 + 12 + 4, 0);
          } else {
            A.store.Bool(ptr + 24 + 12 + 10, true);
            A.store.Bool(ptr + 24 + 12 + 8, "x" in x["pair2"]["touchPoint"] ? true : false);
            A.store.Int32(
              ptr + 24 + 12 + 0,
              x["pair2"]["touchPoint"]["x"] === undefined ? 0 : (x["pair2"]["touchPoint"]["x"] as number)
            );
            A.store.Bool(ptr + 24 + 12 + 9, "y" in x["pair2"]["touchPoint"] ? true : false);
            A.store.Int32(
              ptr + 24 + 12 + 4,
              x["pair2"]["touchPoint"]["y"] === undefined ? 0 : (x["pair2"]["touchPoint"]["y"] as number)
            );
          }
        }

        if (typeof x["pair3"] === "undefined") {
          A.store.Bool(ptr + 48 + 23, false);

          A.store.Bool(ptr + 48 + 0 + 10, false);
          A.store.Bool(ptr + 48 + 0 + 8, false);
          A.store.Int32(ptr + 48 + 0 + 0, 0);
          A.store.Bool(ptr + 48 + 0 + 9, false);
          A.store.Int32(ptr + 48 + 0 + 4, 0);

          A.store.Bool(ptr + 48 + 12 + 10, false);
          A.store.Bool(ptr + 48 + 12 + 8, false);
          A.store.Int32(ptr + 48 + 12 + 0, 0);
          A.store.Bool(ptr + 48 + 12 + 9, false);
          A.store.Int32(ptr + 48 + 12 + 4, 0);
        } else {
          A.store.Bool(ptr + 48 + 23, true);

          if (typeof x["pair3"]["displayPoint"] === "undefined") {
            A.store.Bool(ptr + 48 + 0 + 10, false);
            A.store.Bool(ptr + 48 + 0 + 8, false);
            A.store.Int32(ptr + 48 + 0 + 0, 0);
            A.store.Bool(ptr + 48 + 0 + 9, false);
            A.store.Int32(ptr + 48 + 0 + 4, 0);
          } else {
            A.store.Bool(ptr + 48 + 0 + 10, true);
            A.store.Bool(ptr + 48 + 0 + 8, "x" in x["pair3"]["displayPoint"] ? true : false);
            A.store.Int32(
              ptr + 48 + 0 + 0,
              x["pair3"]["displayPoint"]["x"] === undefined ? 0 : (x["pair3"]["displayPoint"]["x"] as number)
            );
            A.store.Bool(ptr + 48 + 0 + 9, "y" in x["pair3"]["displayPoint"] ? true : false);
            A.store.Int32(
              ptr + 48 + 0 + 4,
              x["pair3"]["displayPoint"]["y"] === undefined ? 0 : (x["pair3"]["displayPoint"]["y"] as number)
            );
          }

          if (typeof x["pair3"]["touchPoint"] === "undefined") {
            A.store.Bool(ptr + 48 + 12 + 10, false);
            A.store.Bool(ptr + 48 + 12 + 8, false);
            A.store.Int32(ptr + 48 + 12 + 0, 0);
            A.store.Bool(ptr + 48 + 12 + 9, false);
            A.store.Int32(ptr + 48 + 12 + 4, 0);
          } else {
            A.store.Bool(ptr + 48 + 12 + 10, true);
            A.store.Bool(ptr + 48 + 12 + 8, "x" in x["pair3"]["touchPoint"] ? true : false);
            A.store.Int32(
              ptr + 48 + 12 + 0,
              x["pair3"]["touchPoint"]["x"] === undefined ? 0 : (x["pair3"]["touchPoint"]["x"] as number)
            );
            A.store.Bool(ptr + 48 + 12 + 9, "y" in x["pair3"]["touchPoint"] ? true : false);
            A.store.Int32(
              ptr + 48 + 12 + 4,
              x["pair3"]["touchPoint"]["y"] === undefined ? 0 : (x["pair3"]["touchPoint"]["y"] as number)
            );
          }
        }

        if (typeof x["pair4"] === "undefined") {
          A.store.Bool(ptr + 72 + 23, false);

          A.store.Bool(ptr + 72 + 0 + 10, false);
          A.store.Bool(ptr + 72 + 0 + 8, false);
          A.store.Int32(ptr + 72 + 0 + 0, 0);
          A.store.Bool(ptr + 72 + 0 + 9, false);
          A.store.Int32(ptr + 72 + 0 + 4, 0);

          A.store.Bool(ptr + 72 + 12 + 10, false);
          A.store.Bool(ptr + 72 + 12 + 8, false);
          A.store.Int32(ptr + 72 + 12 + 0, 0);
          A.store.Bool(ptr + 72 + 12 + 9, false);
          A.store.Int32(ptr + 72 + 12 + 4, 0);
        } else {
          A.store.Bool(ptr + 72 + 23, true);

          if (typeof x["pair4"]["displayPoint"] === "undefined") {
            A.store.Bool(ptr + 72 + 0 + 10, false);
            A.store.Bool(ptr + 72 + 0 + 8, false);
            A.store.Int32(ptr + 72 + 0 + 0, 0);
            A.store.Bool(ptr + 72 + 0 + 9, false);
            A.store.Int32(ptr + 72 + 0 + 4, 0);
          } else {
            A.store.Bool(ptr + 72 + 0 + 10, true);
            A.store.Bool(ptr + 72 + 0 + 8, "x" in x["pair4"]["displayPoint"] ? true : false);
            A.store.Int32(
              ptr + 72 + 0 + 0,
              x["pair4"]["displayPoint"]["x"] === undefined ? 0 : (x["pair4"]["displayPoint"]["x"] as number)
            );
            A.store.Bool(ptr + 72 + 0 + 9, "y" in x["pair4"]["displayPoint"] ? true : false);
            A.store.Int32(
              ptr + 72 + 0 + 4,
              x["pair4"]["displayPoint"]["y"] === undefined ? 0 : (x["pair4"]["displayPoint"]["y"] as number)
            );
          }

          if (typeof x["pair4"]["touchPoint"] === "undefined") {
            A.store.Bool(ptr + 72 + 12 + 10, false);
            A.store.Bool(ptr + 72 + 12 + 8, false);
            A.store.Int32(ptr + 72 + 12 + 0, 0);
            A.store.Bool(ptr + 72 + 12 + 9, false);
            A.store.Int32(ptr + 72 + 12 + 4, 0);
          } else {
            A.store.Bool(ptr + 72 + 12 + 10, true);
            A.store.Bool(ptr + 72 + 12 + 8, "x" in x["pair4"]["touchPoint"] ? true : false);
            A.store.Int32(
              ptr + 72 + 12 + 0,
              x["pair4"]["touchPoint"]["x"] === undefined ? 0 : (x["pair4"]["touchPoint"]["x"] as number)
            );
            A.store.Bool(ptr + 72 + 12 + 9, "y" in x["pair4"]["touchPoint"] ? true : false);
            A.store.Int32(
              ptr + 72 + 12 + 4,
              x["pair4"]["touchPoint"]["y"] === undefined ? 0 : (x["pair4"]["touchPoint"]["y"] as number)
            );
          }
        }
      }
    },
    "load_TouchCalibrationPairQuad": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 23)) {
        x["pair1"] = {};
        if (A.load.Bool(ptr + 0 + 0 + 10)) {
          x["pair1"]["displayPoint"] = {};
          if (A.load.Bool(ptr + 0 + 0 + 8)) {
            x["pair1"]["displayPoint"]["x"] = A.load.Int32(ptr + 0 + 0 + 0);
          } else {
            delete x["pair1"]["displayPoint"]["x"];
          }
          if (A.load.Bool(ptr + 0 + 0 + 9)) {
            x["pair1"]["displayPoint"]["y"] = A.load.Int32(ptr + 0 + 0 + 4);
          } else {
            delete x["pair1"]["displayPoint"]["y"];
          }
        } else {
          delete x["pair1"]["displayPoint"];
        }
        if (A.load.Bool(ptr + 0 + 12 + 10)) {
          x["pair1"]["touchPoint"] = {};
          if (A.load.Bool(ptr + 0 + 12 + 8)) {
            x["pair1"]["touchPoint"]["x"] = A.load.Int32(ptr + 0 + 12 + 0);
          } else {
            delete x["pair1"]["touchPoint"]["x"];
          }
          if (A.load.Bool(ptr + 0 + 12 + 9)) {
            x["pair1"]["touchPoint"]["y"] = A.load.Int32(ptr + 0 + 12 + 4);
          } else {
            delete x["pair1"]["touchPoint"]["y"];
          }
        } else {
          delete x["pair1"]["touchPoint"];
        }
      } else {
        delete x["pair1"];
      }
      if (A.load.Bool(ptr + 24 + 23)) {
        x["pair2"] = {};
        if (A.load.Bool(ptr + 24 + 0 + 10)) {
          x["pair2"]["displayPoint"] = {};
          if (A.load.Bool(ptr + 24 + 0 + 8)) {
            x["pair2"]["displayPoint"]["x"] = A.load.Int32(ptr + 24 + 0 + 0);
          } else {
            delete x["pair2"]["displayPoint"]["x"];
          }
          if (A.load.Bool(ptr + 24 + 0 + 9)) {
            x["pair2"]["displayPoint"]["y"] = A.load.Int32(ptr + 24 + 0 + 4);
          } else {
            delete x["pair2"]["displayPoint"]["y"];
          }
        } else {
          delete x["pair2"]["displayPoint"];
        }
        if (A.load.Bool(ptr + 24 + 12 + 10)) {
          x["pair2"]["touchPoint"] = {};
          if (A.load.Bool(ptr + 24 + 12 + 8)) {
            x["pair2"]["touchPoint"]["x"] = A.load.Int32(ptr + 24 + 12 + 0);
          } else {
            delete x["pair2"]["touchPoint"]["x"];
          }
          if (A.load.Bool(ptr + 24 + 12 + 9)) {
            x["pair2"]["touchPoint"]["y"] = A.load.Int32(ptr + 24 + 12 + 4);
          } else {
            delete x["pair2"]["touchPoint"]["y"];
          }
        } else {
          delete x["pair2"]["touchPoint"];
        }
      } else {
        delete x["pair2"];
      }
      if (A.load.Bool(ptr + 48 + 23)) {
        x["pair3"] = {};
        if (A.load.Bool(ptr + 48 + 0 + 10)) {
          x["pair3"]["displayPoint"] = {};
          if (A.load.Bool(ptr + 48 + 0 + 8)) {
            x["pair3"]["displayPoint"]["x"] = A.load.Int32(ptr + 48 + 0 + 0);
          } else {
            delete x["pair3"]["displayPoint"]["x"];
          }
          if (A.load.Bool(ptr + 48 + 0 + 9)) {
            x["pair3"]["displayPoint"]["y"] = A.load.Int32(ptr + 48 + 0 + 4);
          } else {
            delete x["pair3"]["displayPoint"]["y"];
          }
        } else {
          delete x["pair3"]["displayPoint"];
        }
        if (A.load.Bool(ptr + 48 + 12 + 10)) {
          x["pair3"]["touchPoint"] = {};
          if (A.load.Bool(ptr + 48 + 12 + 8)) {
            x["pair3"]["touchPoint"]["x"] = A.load.Int32(ptr + 48 + 12 + 0);
          } else {
            delete x["pair3"]["touchPoint"]["x"];
          }
          if (A.load.Bool(ptr + 48 + 12 + 9)) {
            x["pair3"]["touchPoint"]["y"] = A.load.Int32(ptr + 48 + 12 + 4);
          } else {
            delete x["pair3"]["touchPoint"]["y"];
          }
        } else {
          delete x["pair3"]["touchPoint"];
        }
      } else {
        delete x["pair3"];
      }
      if (A.load.Bool(ptr + 72 + 23)) {
        x["pair4"] = {};
        if (A.load.Bool(ptr + 72 + 0 + 10)) {
          x["pair4"]["displayPoint"] = {};
          if (A.load.Bool(ptr + 72 + 0 + 8)) {
            x["pair4"]["displayPoint"]["x"] = A.load.Int32(ptr + 72 + 0 + 0);
          } else {
            delete x["pair4"]["displayPoint"]["x"];
          }
          if (A.load.Bool(ptr + 72 + 0 + 9)) {
            x["pair4"]["displayPoint"]["y"] = A.load.Int32(ptr + 72 + 0 + 4);
          } else {
            delete x["pair4"]["displayPoint"]["y"];
          }
        } else {
          delete x["pair4"]["displayPoint"];
        }
        if (A.load.Bool(ptr + 72 + 12 + 10)) {
          x["pair4"]["touchPoint"] = {};
          if (A.load.Bool(ptr + 72 + 12 + 8)) {
            x["pair4"]["touchPoint"]["x"] = A.load.Int32(ptr + 72 + 12 + 0);
          } else {
            delete x["pair4"]["touchPoint"]["x"];
          }
          if (A.load.Bool(ptr + 72 + 12 + 9)) {
            x["pair4"]["touchPoint"]["y"] = A.load.Int32(ptr + 72 + 12 + 4);
          } else {
            delete x["pair4"]["touchPoint"]["y"];
          }
        } else {
          delete x["pair4"]["touchPoint"];
        }
      } else {
        delete x["pair4"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_ClearTouchCalibration": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display && "clearTouchCalibration" in WEBEXT?.system?.display) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ClearTouchCalibration": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.clearTouchCalibration);
    },
    "call_ClearTouchCalibration": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.display.clearTouchCalibration(A.H.get<object>(id));
    },
    "try_ClearTouchCalibration": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.display.clearTouchCalibration(A.H.get<object>(id));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CompleteCustomTouchCalibration": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display && "completeCustomTouchCalibration" in WEBEXT?.system?.display) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CompleteCustomTouchCalibration": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.completeCustomTouchCalibration);
    },
    "call_CompleteCustomTouchCalibration": (retPtr: Pointer, pairs: Pointer, bounds: Pointer): void => {
      const pairs_ffi = {};

      if (A.load.Bool(pairs + 0 + 23)) {
        pairs_ffi["pair1"] = {};
        if (A.load.Bool(pairs + 0 + 0 + 10)) {
          pairs_ffi["pair1"]["displayPoint"] = {};
          if (A.load.Bool(pairs + 0 + 0 + 8)) {
            pairs_ffi["pair1"]["displayPoint"]["x"] = A.load.Int32(pairs + 0 + 0 + 0);
          }
          if (A.load.Bool(pairs + 0 + 0 + 9)) {
            pairs_ffi["pair1"]["displayPoint"]["y"] = A.load.Int32(pairs + 0 + 0 + 4);
          }
        }
        if (A.load.Bool(pairs + 0 + 12 + 10)) {
          pairs_ffi["pair1"]["touchPoint"] = {};
          if (A.load.Bool(pairs + 0 + 12 + 8)) {
            pairs_ffi["pair1"]["touchPoint"]["x"] = A.load.Int32(pairs + 0 + 12 + 0);
          }
          if (A.load.Bool(pairs + 0 + 12 + 9)) {
            pairs_ffi["pair1"]["touchPoint"]["y"] = A.load.Int32(pairs + 0 + 12 + 4);
          }
        }
      }
      if (A.load.Bool(pairs + 24 + 23)) {
        pairs_ffi["pair2"] = {};
        if (A.load.Bool(pairs + 24 + 0 + 10)) {
          pairs_ffi["pair2"]["displayPoint"] = {};
          if (A.load.Bool(pairs + 24 + 0 + 8)) {
            pairs_ffi["pair2"]["displayPoint"]["x"] = A.load.Int32(pairs + 24 + 0 + 0);
          }
          if (A.load.Bool(pairs + 24 + 0 + 9)) {
            pairs_ffi["pair2"]["displayPoint"]["y"] = A.load.Int32(pairs + 24 + 0 + 4);
          }
        }
        if (A.load.Bool(pairs + 24 + 12 + 10)) {
          pairs_ffi["pair2"]["touchPoint"] = {};
          if (A.load.Bool(pairs + 24 + 12 + 8)) {
            pairs_ffi["pair2"]["touchPoint"]["x"] = A.load.Int32(pairs + 24 + 12 + 0);
          }
          if (A.load.Bool(pairs + 24 + 12 + 9)) {
            pairs_ffi["pair2"]["touchPoint"]["y"] = A.load.Int32(pairs + 24 + 12 + 4);
          }
        }
      }
      if (A.load.Bool(pairs + 48 + 23)) {
        pairs_ffi["pair3"] = {};
        if (A.load.Bool(pairs + 48 + 0 + 10)) {
          pairs_ffi["pair3"]["displayPoint"] = {};
          if (A.load.Bool(pairs + 48 + 0 + 8)) {
            pairs_ffi["pair3"]["displayPoint"]["x"] = A.load.Int32(pairs + 48 + 0 + 0);
          }
          if (A.load.Bool(pairs + 48 + 0 + 9)) {
            pairs_ffi["pair3"]["displayPoint"]["y"] = A.load.Int32(pairs + 48 + 0 + 4);
          }
        }
        if (A.load.Bool(pairs + 48 + 12 + 10)) {
          pairs_ffi["pair3"]["touchPoint"] = {};
          if (A.load.Bool(pairs + 48 + 12 + 8)) {
            pairs_ffi["pair3"]["touchPoint"]["x"] = A.load.Int32(pairs + 48 + 12 + 0);
          }
          if (A.load.Bool(pairs + 48 + 12 + 9)) {
            pairs_ffi["pair3"]["touchPoint"]["y"] = A.load.Int32(pairs + 48 + 12 + 4);
          }
        }
      }
      if (A.load.Bool(pairs + 72 + 23)) {
        pairs_ffi["pair4"] = {};
        if (A.load.Bool(pairs + 72 + 0 + 10)) {
          pairs_ffi["pair4"]["displayPoint"] = {};
          if (A.load.Bool(pairs + 72 + 0 + 8)) {
            pairs_ffi["pair4"]["displayPoint"]["x"] = A.load.Int32(pairs + 72 + 0 + 0);
          }
          if (A.load.Bool(pairs + 72 + 0 + 9)) {
            pairs_ffi["pair4"]["displayPoint"]["y"] = A.load.Int32(pairs + 72 + 0 + 4);
          }
        }
        if (A.load.Bool(pairs + 72 + 12 + 10)) {
          pairs_ffi["pair4"]["touchPoint"] = {};
          if (A.load.Bool(pairs + 72 + 12 + 8)) {
            pairs_ffi["pair4"]["touchPoint"]["x"] = A.load.Int32(pairs + 72 + 12 + 0);
          }
          if (A.load.Bool(pairs + 72 + 12 + 9)) {
            pairs_ffi["pair4"]["touchPoint"]["y"] = A.load.Int32(pairs + 72 + 12 + 4);
          }
        }
      }

      const bounds_ffi = {};

      if (A.load.Bool(bounds + 16)) {
        bounds_ffi["left"] = A.load.Int32(bounds + 0);
      }
      if (A.load.Bool(bounds + 17)) {
        bounds_ffi["top"] = A.load.Int32(bounds + 4);
      }
      if (A.load.Bool(bounds + 18)) {
        bounds_ffi["width"] = A.load.Int32(bounds + 8);
      }
      if (A.load.Bool(bounds + 19)) {
        bounds_ffi["height"] = A.load.Int32(bounds + 12);
      }

      const _ret = WEBEXT.system.display.completeCustomTouchCalibration(pairs_ffi, bounds_ffi);
    },
    "try_CompleteCustomTouchCalibration": (
      retPtr: Pointer,
      errPtr: Pointer,
      pairs: Pointer,
      bounds: Pointer
    ): heap.Ref<boolean> => {
      try {
        const pairs_ffi = {};

        if (A.load.Bool(pairs + 0 + 23)) {
          pairs_ffi["pair1"] = {};
          if (A.load.Bool(pairs + 0 + 0 + 10)) {
            pairs_ffi["pair1"]["displayPoint"] = {};
            if (A.load.Bool(pairs + 0 + 0 + 8)) {
              pairs_ffi["pair1"]["displayPoint"]["x"] = A.load.Int32(pairs + 0 + 0 + 0);
            }
            if (A.load.Bool(pairs + 0 + 0 + 9)) {
              pairs_ffi["pair1"]["displayPoint"]["y"] = A.load.Int32(pairs + 0 + 0 + 4);
            }
          }
          if (A.load.Bool(pairs + 0 + 12 + 10)) {
            pairs_ffi["pair1"]["touchPoint"] = {};
            if (A.load.Bool(pairs + 0 + 12 + 8)) {
              pairs_ffi["pair1"]["touchPoint"]["x"] = A.load.Int32(pairs + 0 + 12 + 0);
            }
            if (A.load.Bool(pairs + 0 + 12 + 9)) {
              pairs_ffi["pair1"]["touchPoint"]["y"] = A.load.Int32(pairs + 0 + 12 + 4);
            }
          }
        }
        if (A.load.Bool(pairs + 24 + 23)) {
          pairs_ffi["pair2"] = {};
          if (A.load.Bool(pairs + 24 + 0 + 10)) {
            pairs_ffi["pair2"]["displayPoint"] = {};
            if (A.load.Bool(pairs + 24 + 0 + 8)) {
              pairs_ffi["pair2"]["displayPoint"]["x"] = A.load.Int32(pairs + 24 + 0 + 0);
            }
            if (A.load.Bool(pairs + 24 + 0 + 9)) {
              pairs_ffi["pair2"]["displayPoint"]["y"] = A.load.Int32(pairs + 24 + 0 + 4);
            }
          }
          if (A.load.Bool(pairs + 24 + 12 + 10)) {
            pairs_ffi["pair2"]["touchPoint"] = {};
            if (A.load.Bool(pairs + 24 + 12 + 8)) {
              pairs_ffi["pair2"]["touchPoint"]["x"] = A.load.Int32(pairs + 24 + 12 + 0);
            }
            if (A.load.Bool(pairs + 24 + 12 + 9)) {
              pairs_ffi["pair2"]["touchPoint"]["y"] = A.load.Int32(pairs + 24 + 12 + 4);
            }
          }
        }
        if (A.load.Bool(pairs + 48 + 23)) {
          pairs_ffi["pair3"] = {};
          if (A.load.Bool(pairs + 48 + 0 + 10)) {
            pairs_ffi["pair3"]["displayPoint"] = {};
            if (A.load.Bool(pairs + 48 + 0 + 8)) {
              pairs_ffi["pair3"]["displayPoint"]["x"] = A.load.Int32(pairs + 48 + 0 + 0);
            }
            if (A.load.Bool(pairs + 48 + 0 + 9)) {
              pairs_ffi["pair3"]["displayPoint"]["y"] = A.load.Int32(pairs + 48 + 0 + 4);
            }
          }
          if (A.load.Bool(pairs + 48 + 12 + 10)) {
            pairs_ffi["pair3"]["touchPoint"] = {};
            if (A.load.Bool(pairs + 48 + 12 + 8)) {
              pairs_ffi["pair3"]["touchPoint"]["x"] = A.load.Int32(pairs + 48 + 12 + 0);
            }
            if (A.load.Bool(pairs + 48 + 12 + 9)) {
              pairs_ffi["pair3"]["touchPoint"]["y"] = A.load.Int32(pairs + 48 + 12 + 4);
            }
          }
        }
        if (A.load.Bool(pairs + 72 + 23)) {
          pairs_ffi["pair4"] = {};
          if (A.load.Bool(pairs + 72 + 0 + 10)) {
            pairs_ffi["pair4"]["displayPoint"] = {};
            if (A.load.Bool(pairs + 72 + 0 + 8)) {
              pairs_ffi["pair4"]["displayPoint"]["x"] = A.load.Int32(pairs + 72 + 0 + 0);
            }
            if (A.load.Bool(pairs + 72 + 0 + 9)) {
              pairs_ffi["pair4"]["displayPoint"]["y"] = A.load.Int32(pairs + 72 + 0 + 4);
            }
          }
          if (A.load.Bool(pairs + 72 + 12 + 10)) {
            pairs_ffi["pair4"]["touchPoint"] = {};
            if (A.load.Bool(pairs + 72 + 12 + 8)) {
              pairs_ffi["pair4"]["touchPoint"]["x"] = A.load.Int32(pairs + 72 + 12 + 0);
            }
            if (A.load.Bool(pairs + 72 + 12 + 9)) {
              pairs_ffi["pair4"]["touchPoint"]["y"] = A.load.Int32(pairs + 72 + 12 + 4);
            }
          }
        }

        const bounds_ffi = {};

        if (A.load.Bool(bounds + 16)) {
          bounds_ffi["left"] = A.load.Int32(bounds + 0);
        }
        if (A.load.Bool(bounds + 17)) {
          bounds_ffi["top"] = A.load.Int32(bounds + 4);
        }
        if (A.load.Bool(bounds + 18)) {
          bounds_ffi["width"] = A.load.Int32(bounds + 8);
        }
        if (A.load.Bool(bounds + 19)) {
          bounds_ffi["height"] = A.load.Int32(bounds + 12);
        }

        const _ret = WEBEXT.system.display.completeCustomTouchCalibration(pairs_ffi, bounds_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_EnableUnifiedDesktop": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display && "enableUnifiedDesktop" in WEBEXT?.system?.display) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EnableUnifiedDesktop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.enableUnifiedDesktop);
    },
    "call_EnableUnifiedDesktop": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.system.display.enableUnifiedDesktop(enabled === A.H.TRUE);
    },
    "try_EnableUnifiedDesktop": (retPtr: Pointer, errPtr: Pointer, enabled: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.display.enableUnifiedDesktop(enabled === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDisplayLayout": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display && "getDisplayLayout" in WEBEXT?.system?.display) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDisplayLayout": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.getDisplayLayout);
    },
    "call_GetDisplayLayout": (retPtr: Pointer): void => {
      const _ret = WEBEXT.system.display.getDisplayLayout();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDisplayLayout": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.display.getDisplayLayout();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display && "getInfo" in WEBEXT?.system?.display) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.getInfo);
    },
    "call_GetInfo": (retPtr: Pointer, flags: Pointer): void => {
      const flags_ffi = {};

      if (A.load.Bool(flags + 1)) {
        flags_ffi["singleUnified"] = A.load.Bool(flags + 0);
      }

      const _ret = WEBEXT.system.display.getInfo(flags_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetInfo": (retPtr: Pointer, errPtr: Pointer, flags: Pointer): heap.Ref<boolean> => {
      try {
        const flags_ffi = {};

        if (A.load.Bool(flags + 1)) {
          flags_ffi["singleUnified"] = A.load.Bool(flags + 0);
        }

        const _ret = WEBEXT.system.display.getInfo(flags_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDisplayChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display?.onDisplayChanged && "addListener" in WEBEXT?.system?.display?.onDisplayChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDisplayChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.onDisplayChanged.addListener);
    },
    "call_OnDisplayChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.display.onDisplayChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnDisplayChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.display.onDisplayChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDisplayChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display?.onDisplayChanged && "removeListener" in WEBEXT?.system?.display?.onDisplayChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDisplayChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.onDisplayChanged.removeListener);
    },
    "call_OffDisplayChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.display.onDisplayChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffDisplayChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.display.onDisplayChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDisplayChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display?.onDisplayChanged && "hasListener" in WEBEXT?.system?.display?.onDisplayChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDisplayChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.onDisplayChanged.hasListener);
    },
    "call_HasOnDisplayChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.display.onDisplayChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDisplayChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.display.onDisplayChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OverscanCalibrationAdjust": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display && "overscanCalibrationAdjust" in WEBEXT?.system?.display) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OverscanCalibrationAdjust": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.overscanCalibrationAdjust);
    },
    "call_OverscanCalibrationAdjust": (retPtr: Pointer, id: heap.Ref<object>, delta: Pointer): void => {
      const delta_ffi = {};

      if (A.load.Bool(delta + 16)) {
        delta_ffi["left"] = A.load.Int32(delta + 0);
      }
      if (A.load.Bool(delta + 17)) {
        delta_ffi["top"] = A.load.Int32(delta + 4);
      }
      if (A.load.Bool(delta + 18)) {
        delta_ffi["right"] = A.load.Int32(delta + 8);
      }
      if (A.load.Bool(delta + 19)) {
        delta_ffi["bottom"] = A.load.Int32(delta + 12);
      }

      const _ret = WEBEXT.system.display.overscanCalibrationAdjust(A.H.get<object>(id), delta_ffi);
    },
    "try_OverscanCalibrationAdjust": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<object>,
      delta: Pointer
    ): heap.Ref<boolean> => {
      try {
        const delta_ffi = {};

        if (A.load.Bool(delta + 16)) {
          delta_ffi["left"] = A.load.Int32(delta + 0);
        }
        if (A.load.Bool(delta + 17)) {
          delta_ffi["top"] = A.load.Int32(delta + 4);
        }
        if (A.load.Bool(delta + 18)) {
          delta_ffi["right"] = A.load.Int32(delta + 8);
        }
        if (A.load.Bool(delta + 19)) {
          delta_ffi["bottom"] = A.load.Int32(delta + 12);
        }

        const _ret = WEBEXT.system.display.overscanCalibrationAdjust(A.H.get<object>(id), delta_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OverscanCalibrationComplete": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display && "overscanCalibrationComplete" in WEBEXT?.system?.display) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OverscanCalibrationComplete": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.overscanCalibrationComplete);
    },
    "call_OverscanCalibrationComplete": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.display.overscanCalibrationComplete(A.H.get<object>(id));
    },
    "try_OverscanCalibrationComplete": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.display.overscanCalibrationComplete(A.H.get<object>(id));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OverscanCalibrationReset": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display && "overscanCalibrationReset" in WEBEXT?.system?.display) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OverscanCalibrationReset": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.overscanCalibrationReset);
    },
    "call_OverscanCalibrationReset": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.display.overscanCalibrationReset(A.H.get<object>(id));
    },
    "try_OverscanCalibrationReset": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.display.overscanCalibrationReset(A.H.get<object>(id));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OverscanCalibrationStart": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display && "overscanCalibrationStart" in WEBEXT?.system?.display) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OverscanCalibrationStart": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.overscanCalibrationStart);
    },
    "call_OverscanCalibrationStart": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.display.overscanCalibrationStart(A.H.get<object>(id));
    },
    "try_OverscanCalibrationStart": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.display.overscanCalibrationStart(A.H.get<object>(id));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetDisplayLayout": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display && "setDisplayLayout" in WEBEXT?.system?.display) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetDisplayLayout": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.setDisplayLayout);
    },
    "call_SetDisplayLayout": (retPtr: Pointer, layouts: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.display.setDisplayLayout(A.H.get<object>(layouts));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetDisplayLayout": (retPtr: Pointer, errPtr: Pointer, layouts: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.display.setDisplayLayout(A.H.get<object>(layouts));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetDisplayProperties": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display && "setDisplayProperties" in WEBEXT?.system?.display) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetDisplayProperties": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.setDisplayProperties);
    },
    "call_SetDisplayProperties": (retPtr: Pointer, id: heap.Ref<object>, info: Pointer): void => {
      const info_ffi = {};

      if (A.load.Bool(info + 112)) {
        info_ffi["isUnified"] = A.load.Bool(info + 0);
      }
      info_ffi["mirroringSourceId"] = A.load.Ref(info + 4, undefined);
      if (A.load.Bool(info + 113)) {
        info_ffi["isPrimary"] = A.load.Bool(info + 8);
      }
      if (A.load.Bool(info + 12 + 20)) {
        info_ffi["overscan"] = {};
        if (A.load.Bool(info + 12 + 16)) {
          info_ffi["overscan"]["left"] = A.load.Int32(info + 12 + 0);
        }
        if (A.load.Bool(info + 12 + 17)) {
          info_ffi["overscan"]["top"] = A.load.Int32(info + 12 + 4);
        }
        if (A.load.Bool(info + 12 + 18)) {
          info_ffi["overscan"]["right"] = A.load.Int32(info + 12 + 8);
        }
        if (A.load.Bool(info + 12 + 19)) {
          info_ffi["overscan"]["bottom"] = A.load.Int32(info + 12 + 12);
        }
      }
      if (A.load.Bool(info + 114)) {
        info_ffi["rotation"] = A.load.Int32(info + 36);
      }
      if (A.load.Bool(info + 115)) {
        info_ffi["boundsOriginX"] = A.load.Int32(info + 40);
      }
      if (A.load.Bool(info + 116)) {
        info_ffi["boundsOriginY"] = A.load.Int32(info + 44);
      }
      if (A.load.Bool(info + 48 + 53)) {
        info_ffi["displayMode"] = {};
        if (A.load.Bool(info + 48 + 43)) {
          info_ffi["displayMode"]["width"] = A.load.Int32(info + 48 + 0);
        }
        if (A.load.Bool(info + 48 + 44)) {
          info_ffi["displayMode"]["height"] = A.load.Int32(info + 48 + 4);
        }
        if (A.load.Bool(info + 48 + 45)) {
          info_ffi["displayMode"]["widthInNativePixels"] = A.load.Int32(info + 48 + 8);
        }
        if (A.load.Bool(info + 48 + 46)) {
          info_ffi["displayMode"]["heightInNativePixels"] = A.load.Int32(info + 48 + 12);
        }
        if (A.load.Bool(info + 48 + 47)) {
          info_ffi["displayMode"]["uiScale"] = A.load.Float64(info + 48 + 16);
        }
        if (A.load.Bool(info + 48 + 48)) {
          info_ffi["displayMode"]["deviceScaleFactor"] = A.load.Float64(info + 48 + 24);
        }
        if (A.load.Bool(info + 48 + 49)) {
          info_ffi["displayMode"]["refreshRate"] = A.load.Float64(info + 48 + 32);
        }
        if (A.load.Bool(info + 48 + 50)) {
          info_ffi["displayMode"]["isNative"] = A.load.Bool(info + 48 + 40);
        }
        if (A.load.Bool(info + 48 + 51)) {
          info_ffi["displayMode"]["isSelected"] = A.load.Bool(info + 48 + 41);
        }
        if (A.load.Bool(info + 48 + 52)) {
          info_ffi["displayMode"]["isInterlaced"] = A.load.Bool(info + 48 + 42);
        }
      }
      if (A.load.Bool(info + 117)) {
        info_ffi["displayZoomFactor"] = A.load.Float64(info + 104);
      }

      const _ret = WEBEXT.system.display.setDisplayProperties(A.H.get<object>(id), info_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetDisplayProperties": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<object>,
      info: Pointer
    ): heap.Ref<boolean> => {
      try {
        const info_ffi = {};

        if (A.load.Bool(info + 112)) {
          info_ffi["isUnified"] = A.load.Bool(info + 0);
        }
        info_ffi["mirroringSourceId"] = A.load.Ref(info + 4, undefined);
        if (A.load.Bool(info + 113)) {
          info_ffi["isPrimary"] = A.load.Bool(info + 8);
        }
        if (A.load.Bool(info + 12 + 20)) {
          info_ffi["overscan"] = {};
          if (A.load.Bool(info + 12 + 16)) {
            info_ffi["overscan"]["left"] = A.load.Int32(info + 12 + 0);
          }
          if (A.load.Bool(info + 12 + 17)) {
            info_ffi["overscan"]["top"] = A.load.Int32(info + 12 + 4);
          }
          if (A.load.Bool(info + 12 + 18)) {
            info_ffi["overscan"]["right"] = A.load.Int32(info + 12 + 8);
          }
          if (A.load.Bool(info + 12 + 19)) {
            info_ffi["overscan"]["bottom"] = A.load.Int32(info + 12 + 12);
          }
        }
        if (A.load.Bool(info + 114)) {
          info_ffi["rotation"] = A.load.Int32(info + 36);
        }
        if (A.load.Bool(info + 115)) {
          info_ffi["boundsOriginX"] = A.load.Int32(info + 40);
        }
        if (A.load.Bool(info + 116)) {
          info_ffi["boundsOriginY"] = A.load.Int32(info + 44);
        }
        if (A.load.Bool(info + 48 + 53)) {
          info_ffi["displayMode"] = {};
          if (A.load.Bool(info + 48 + 43)) {
            info_ffi["displayMode"]["width"] = A.load.Int32(info + 48 + 0);
          }
          if (A.load.Bool(info + 48 + 44)) {
            info_ffi["displayMode"]["height"] = A.load.Int32(info + 48 + 4);
          }
          if (A.load.Bool(info + 48 + 45)) {
            info_ffi["displayMode"]["widthInNativePixels"] = A.load.Int32(info + 48 + 8);
          }
          if (A.load.Bool(info + 48 + 46)) {
            info_ffi["displayMode"]["heightInNativePixels"] = A.load.Int32(info + 48 + 12);
          }
          if (A.load.Bool(info + 48 + 47)) {
            info_ffi["displayMode"]["uiScale"] = A.load.Float64(info + 48 + 16);
          }
          if (A.load.Bool(info + 48 + 48)) {
            info_ffi["displayMode"]["deviceScaleFactor"] = A.load.Float64(info + 48 + 24);
          }
          if (A.load.Bool(info + 48 + 49)) {
            info_ffi["displayMode"]["refreshRate"] = A.load.Float64(info + 48 + 32);
          }
          if (A.load.Bool(info + 48 + 50)) {
            info_ffi["displayMode"]["isNative"] = A.load.Bool(info + 48 + 40);
          }
          if (A.load.Bool(info + 48 + 51)) {
            info_ffi["displayMode"]["isSelected"] = A.load.Bool(info + 48 + 41);
          }
          if (A.load.Bool(info + 48 + 52)) {
            info_ffi["displayMode"]["isInterlaced"] = A.load.Bool(info + 48 + 42);
          }
        }
        if (A.load.Bool(info + 117)) {
          info_ffi["displayZoomFactor"] = A.load.Float64(info + 104);
        }

        const _ret = WEBEXT.system.display.setDisplayProperties(A.H.get<object>(id), info_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetMirrorMode": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display && "setMirrorMode" in WEBEXT?.system?.display) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetMirrorMode": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.setMirrorMode);
    },
    "call_SetMirrorMode": (retPtr: Pointer, info: Pointer): void => {
      const info_ffi = {};

      info_ffi["mode"] = A.load.Enum(info + 0, ["off", "normal", "mixed"]);
      info_ffi["mirroringSourceId"] = A.load.Ref(info + 4, undefined);
      info_ffi["mirroringDestinationIds"] = A.load.Ref(info + 8, undefined);

      const _ret = WEBEXT.system.display.setMirrorMode(info_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetMirrorMode": (retPtr: Pointer, errPtr: Pointer, info: Pointer): heap.Ref<boolean> => {
      try {
        const info_ffi = {};

        info_ffi["mode"] = A.load.Enum(info + 0, ["off", "normal", "mixed"]);
        info_ffi["mirroringSourceId"] = A.load.Ref(info + 4, undefined);
        info_ffi["mirroringDestinationIds"] = A.load.Ref(info + 8, undefined);

        const _ret = WEBEXT.system.display.setMirrorMode(info_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ShowNativeTouchCalibration": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display && "showNativeTouchCalibration" in WEBEXT?.system?.display) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ShowNativeTouchCalibration": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.showNativeTouchCalibration);
    },
    "call_ShowNativeTouchCalibration": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.display.showNativeTouchCalibration(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_ShowNativeTouchCalibration": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.display.showNativeTouchCalibration(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartCustomTouchCalibration": (): heap.Ref<boolean> => {
      if (WEBEXT?.system?.display && "startCustomTouchCalibration" in WEBEXT?.system?.display) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartCustomTouchCalibration": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.system.display.startCustomTouchCalibration);
    },
    "call_StartCustomTouchCalibration": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.system.display.startCustomTouchCalibration(A.H.get<object>(id));
    },
    "try_StartCustomTouchCalibration": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.system.display.startCustomTouchCalibration(A.H.get<object>(id));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
