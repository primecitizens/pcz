import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/autotestprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Accelerator": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 10, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 11, false);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 13, false);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Ref(ptr + 0, x["keyCode"]);
        A.store.Bool(ptr + 9, "shift" in x ? true : false);
        A.store.Bool(ptr + 4, x["shift"] ? true : false);
        A.store.Bool(ptr + 10, "control" in x ? true : false);
        A.store.Bool(ptr + 5, x["control"] ? true : false);
        A.store.Bool(ptr + 11, "alt" in x ? true : false);
        A.store.Bool(ptr + 6, x["alt"] ? true : false);
        A.store.Bool(ptr + 12, "search" in x ? true : false);
        A.store.Bool(ptr + 7, x["search"] ? true : false);
        A.store.Bool(ptr + 13, "pressed" in x ? true : false);
        A.store.Bool(ptr + 8, x["pressed"] ? true : false);
      }
    },
    "load_Accelerator": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["keyCode"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 9)) {
        x["shift"] = A.load.Bool(ptr + 4);
      } else {
        delete x["shift"];
      }
      if (A.load.Bool(ptr + 10)) {
        x["control"] = A.load.Bool(ptr + 5);
      } else {
        delete x["control"];
      }
      if (A.load.Bool(ptr + 11)) {
        x["alt"] = A.load.Bool(ptr + 6);
      } else {
        delete x["alt"];
      }
      if (A.load.Bool(ptr + 12)) {
        x["search"] = A.load.Bool(ptr + 7);
      } else {
        delete x["search"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["pressed"] = A.load.Bool(ptr + 8);
      } else {
        delete x["pressed"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_AppType": (ref: heap.Ref<string>): number => {
      const idx = [
        "Arc",
        "BuiltIn",
        "Crostini",
        "Extension",
        "Web",
        "MacOS",
        "PluginVm",
        "StandaloneBrowser",
        "Remote",
        "Borealis",
        "Bruschetta",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_AppInstallSource": (ref: heap.Ref<string>): number => {
      const idx = [
        "Unknown",
        "System",
        "Policy",
        "Oem",
        "Default",
        "Sync",
        "User",
        "SubApp",
        "Kiosk",
        "CommandLine",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_AppReadiness": (ref: heap.Ref<string>): number => {
      const idx = [
        "Ready",
        "DisabledByBlacklist",
        "DisabledByPolicy",
        "DisabledByUser",
        "Terminated",
        "UninstalledByUser",
        "Removed",
        "UninstalledByMigration",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_App": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 36, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Enum(ptr + 16, -1);
        A.store.Enum(ptr + 20, -1);
        A.store.Enum(ptr + 24, -1);
        A.store.Ref(ptr + 28, undefined);
        A.store.Bool(ptr + 34, false);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 35, false);
        A.store.Bool(ptr + 33, false);
      } else {
        A.store.Bool(ptr + 36, true);
        A.store.Ref(ptr + 0, x["appId"]);
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Ref(ptr + 8, x["shortName"]);
        A.store.Ref(ptr + 12, x["publisherId"]);
        A.store.Enum(
          ptr + 16,
          [
            "Arc",
            "BuiltIn",
            "Crostini",
            "Extension",
            "Web",
            "MacOS",
            "PluginVm",
            "StandaloneBrowser",
            "Remote",
            "Borealis",
            "Bruschetta",
          ].indexOf(x["type"] as string)
        );
        A.store.Enum(
          ptr + 20,
          ["Unknown", "System", "Policy", "Oem", "Default", "Sync", "User", "SubApp", "Kiosk", "CommandLine"].indexOf(
            x["installSource"] as string
          )
        );
        A.store.Enum(
          ptr + 24,
          [
            "Ready",
            "DisabledByBlacklist",
            "DisabledByPolicy",
            "DisabledByUser",
            "Terminated",
            "UninstalledByUser",
            "Removed",
            "UninstalledByMigration",
          ].indexOf(x["readiness"] as string)
        );
        A.store.Ref(ptr + 28, x["additionalSearchTerms"]);
        A.store.Bool(ptr + 34, "showInLauncher" in x ? true : false);
        A.store.Bool(ptr + 32, x["showInLauncher"] ? true : false);
        A.store.Bool(ptr + 35, "showInSearch" in x ? true : false);
        A.store.Bool(ptr + 33, x["showInSearch"] ? true : false);
      }
    },
    "load_App": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["appId"] = A.load.Ref(ptr + 0, undefined);
      x["name"] = A.load.Ref(ptr + 4, undefined);
      x["shortName"] = A.load.Ref(ptr + 8, undefined);
      x["publisherId"] = A.load.Ref(ptr + 12, undefined);
      x["type"] = A.load.Enum(ptr + 16, [
        "Arc",
        "BuiltIn",
        "Crostini",
        "Extension",
        "Web",
        "MacOS",
        "PluginVm",
        "StandaloneBrowser",
        "Remote",
        "Borealis",
        "Bruschetta",
      ]);
      x["installSource"] = A.load.Enum(ptr + 20, [
        "Unknown",
        "System",
        "Policy",
        "Oem",
        "Default",
        "Sync",
        "User",
        "SubApp",
        "Kiosk",
        "CommandLine",
      ]);
      x["readiness"] = A.load.Enum(ptr + 24, [
        "Ready",
        "DisabledByBlacklist",
        "DisabledByPolicy",
        "DisabledByUser",
        "Terminated",
        "UninstalledByUser",
        "Removed",
        "UninstalledByMigration",
      ]);
      x["additionalSearchTerms"] = A.load.Ref(ptr + 28, undefined);
      if (A.load.Bool(ptr + 34)) {
        x["showInLauncher"] = A.load.Bool(ptr + 32);
      } else {
        delete x["showInLauncher"];
      }
      if (A.load.Bool(ptr + 35)) {
        x["showInSearch"] = A.load.Bool(ptr + 33);
      } else {
        delete x["showInSearch"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_AppWindowType": (ref: heap.Ref<string>): number => {
      const idx = ["Browser", "ChromeApp", "ArcApp", "CrostiniApp", "SystemApp", "ExtensionApp", "Lacros"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_WindowStateType": (ref: heap.Ref<string>): number => {
      const idx = [
        "Normal",
        "Minimized",
        "Maximized",
        "Fullscreen",
        "PrimarySnapped",
        "SecondarySnapped",
        "PIP",
        "Floated",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Bounds": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 36, false);
        A.store.Bool(ptr + 32, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 33, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 34, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 35, false);
        A.store.Float64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 36, true);
        A.store.Bool(ptr + 32, "left" in x ? true : false);
        A.store.Float64(ptr + 0, x["left"] === undefined ? 0 : (x["left"] as number));
        A.store.Bool(ptr + 33, "top" in x ? true : false);
        A.store.Float64(ptr + 8, x["top"] === undefined ? 0 : (x["top"] as number));
        A.store.Bool(ptr + 34, "width" in x ? true : false);
        A.store.Float64(ptr + 16, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Bool(ptr + 35, "height" in x ? true : false);
        A.store.Float64(ptr + 24, x["height"] === undefined ? 0 : (x["height"] as number));
      }
    },
    "load_Bounds": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 32)) {
        x["left"] = A.load.Float64(ptr + 0);
      } else {
        delete x["left"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["top"] = A.load.Float64(ptr + 8);
      } else {
        delete x["top"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["width"] = A.load.Float64(ptr + 16);
      } else {
        delete x["width"];
      }
      if (A.load.Bool(ptr + 35)) {
        x["height"] = A.load.Float64(ptr + 24);
      } else {
        delete x["height"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_FrameMode": (ref: heap.Ref<string>): number => {
      const idx = ["Normal", "Immersive"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OverviewInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 39, false);

        A.store.Bool(ptr + 0 + 36, false);
        A.store.Bool(ptr + 0 + 32, false);
        A.store.Float64(ptr + 0 + 0, 0);
        A.store.Bool(ptr + 0 + 33, false);
        A.store.Float64(ptr + 0 + 8, 0);
        A.store.Bool(ptr + 0 + 34, false);
        A.store.Float64(ptr + 0 + 16, 0);
        A.store.Bool(ptr + 0 + 35, false);
        A.store.Float64(ptr + 0 + 24, 0);
        A.store.Bool(ptr + 38, false);
        A.store.Bool(ptr + 37, false);
      } else {
        A.store.Bool(ptr + 39, true);

        if (typeof x["bounds"] === "undefined") {
          A.store.Bool(ptr + 0 + 36, false);
          A.store.Bool(ptr + 0 + 32, false);
          A.store.Float64(ptr + 0 + 0, 0);
          A.store.Bool(ptr + 0 + 33, false);
          A.store.Float64(ptr + 0 + 8, 0);
          A.store.Bool(ptr + 0 + 34, false);
          A.store.Float64(ptr + 0 + 16, 0);
          A.store.Bool(ptr + 0 + 35, false);
          A.store.Float64(ptr + 0 + 24, 0);
        } else {
          A.store.Bool(ptr + 0 + 36, true);
          A.store.Bool(ptr + 0 + 32, "left" in x["bounds"] ? true : false);
          A.store.Float64(ptr + 0 + 0, x["bounds"]["left"] === undefined ? 0 : (x["bounds"]["left"] as number));
          A.store.Bool(ptr + 0 + 33, "top" in x["bounds"] ? true : false);
          A.store.Float64(ptr + 0 + 8, x["bounds"]["top"] === undefined ? 0 : (x["bounds"]["top"] as number));
          A.store.Bool(ptr + 0 + 34, "width" in x["bounds"] ? true : false);
          A.store.Float64(ptr + 0 + 16, x["bounds"]["width"] === undefined ? 0 : (x["bounds"]["width"] as number));
          A.store.Bool(ptr + 0 + 35, "height" in x["bounds"] ? true : false);
          A.store.Float64(ptr + 0 + 24, x["bounds"]["height"] === undefined ? 0 : (x["bounds"]["height"] as number));
        }
        A.store.Bool(ptr + 38, "isDragged" in x ? true : false);
        A.store.Bool(ptr + 37, x["isDragged"] ? true : false);
      }
    },
    "load_OverviewInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 36)) {
        x["bounds"] = {};
        if (A.load.Bool(ptr + 0 + 32)) {
          x["bounds"]["left"] = A.load.Float64(ptr + 0 + 0);
        } else {
          delete x["bounds"]["left"];
        }
        if (A.load.Bool(ptr + 0 + 33)) {
          x["bounds"]["top"] = A.load.Float64(ptr + 0 + 8);
        } else {
          delete x["bounds"]["top"];
        }
        if (A.load.Bool(ptr + 0 + 34)) {
          x["bounds"]["width"] = A.load.Float64(ptr + 0 + 16);
        } else {
          delete x["bounds"]["width"];
        }
        if (A.load.Bool(ptr + 0 + 35)) {
          x["bounds"]["height"] = A.load.Float64(ptr + 0 + 24);
        } else {
          delete x["bounds"]["height"];
        }
      } else {
        delete x["bounds"];
      }
      if (A.load.Bool(ptr + 38)) {
        x["isDragged"] = A.load.Bool(ptr + 37);
      } else {
        delete x["isDragged"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AppWindowInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 207, false);
        A.store.Bool(ptr + 192, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
        A.store.Enum(ptr + 12, -1);

        A.store.Bool(ptr + 16 + 36, false);
        A.store.Bool(ptr + 16 + 32, false);
        A.store.Float64(ptr + 16 + 0, 0);
        A.store.Bool(ptr + 16 + 33, false);
        A.store.Float64(ptr + 16 + 8, 0);
        A.store.Bool(ptr + 16 + 34, false);
        A.store.Float64(ptr + 16 + 16, 0);
        A.store.Bool(ptr + 16 + 35, false);
        A.store.Float64(ptr + 16 + 24, 0);
        A.store.Ref(ptr + 56, undefined);
        A.store.Bool(ptr + 193, false);
        A.store.Bool(ptr + 60, false);
        A.store.Bool(ptr + 194, false);
        A.store.Bool(ptr + 61, false);
        A.store.Ref(ptr + 64, undefined);
        A.store.Bool(ptr + 195, false);
        A.store.Bool(ptr + 68, false);

        A.store.Bool(ptr + 72 + 36, false);
        A.store.Bool(ptr + 72 + 32, false);
        A.store.Float64(ptr + 72 + 0, 0);
        A.store.Bool(ptr + 72 + 33, false);
        A.store.Float64(ptr + 72 + 8, 0);
        A.store.Bool(ptr + 72 + 34, false);
        A.store.Float64(ptr + 72 + 16, 0);
        A.store.Bool(ptr + 72 + 35, false);
        A.store.Float64(ptr + 72 + 24, 0);
        A.store.Bool(ptr + 196, false);
        A.store.Bool(ptr + 109, false);
        A.store.Bool(ptr + 197, false);
        A.store.Bool(ptr + 110, false);
        A.store.Bool(ptr + 198, false);
        A.store.Bool(ptr + 111, false);
        A.store.Bool(ptr + 199, false);
        A.store.Bool(ptr + 112, false);
        A.store.Bool(ptr + 200, false);
        A.store.Bool(ptr + 113, false);
        A.store.Bool(ptr + 201, false);
        A.store.Bool(ptr + 114, false);
        A.store.Bool(ptr + 202, false);
        A.store.Int32(ptr + 116, 0);
        A.store.Enum(ptr + 120, -1);
        A.store.Bool(ptr + 203, false);
        A.store.Bool(ptr + 124, false);
        A.store.Bool(ptr + 204, false);
        A.store.Int32(ptr + 128, 0);
        A.store.Bool(ptr + 205, false);
        A.store.Int32(ptr + 132, 0);
        A.store.Bool(ptr + 206, false);
        A.store.Int32(ptr + 136, 0);
        A.store.Ref(ptr + 140, undefined);

        A.store.Bool(ptr + 144 + 39, false);

        A.store.Bool(ptr + 144 + 0 + 36, false);
        A.store.Bool(ptr + 144 + 0 + 32, false);
        A.store.Float64(ptr + 144 + 0 + 0, 0);
        A.store.Bool(ptr + 144 + 0 + 33, false);
        A.store.Float64(ptr + 144 + 0 + 8, 0);
        A.store.Bool(ptr + 144 + 0 + 34, false);
        A.store.Float64(ptr + 144 + 0 + 16, 0);
        A.store.Bool(ptr + 144 + 0 + 35, false);
        A.store.Float64(ptr + 144 + 0 + 24, 0);
        A.store.Bool(ptr + 144 + 38, false);
        A.store.Bool(ptr + 144 + 37, false);
        A.store.Ref(ptr + 184, undefined);
        A.store.Ref(ptr + 188, undefined);
      } else {
        A.store.Bool(ptr + 207, true);
        A.store.Bool(ptr + 192, "id" in x ? true : false);
        A.store.Int32(ptr + 0, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Ref(ptr + 4, x["name"]);
        A.store.Enum(
          ptr + 8,
          ["Browser", "ChromeApp", "ArcApp", "CrostiniApp", "SystemApp", "ExtensionApp", "Lacros"].indexOf(
            x["windowType"] as string
          )
        );
        A.store.Enum(
          ptr + 12,
          [
            "Normal",
            "Minimized",
            "Maximized",
            "Fullscreen",
            "PrimarySnapped",
            "SecondarySnapped",
            "PIP",
            "Floated",
          ].indexOf(x["stateType"] as string)
        );

        if (typeof x["boundsInRoot"] === "undefined") {
          A.store.Bool(ptr + 16 + 36, false);
          A.store.Bool(ptr + 16 + 32, false);
          A.store.Float64(ptr + 16 + 0, 0);
          A.store.Bool(ptr + 16 + 33, false);
          A.store.Float64(ptr + 16 + 8, 0);
          A.store.Bool(ptr + 16 + 34, false);
          A.store.Float64(ptr + 16 + 16, 0);
          A.store.Bool(ptr + 16 + 35, false);
          A.store.Float64(ptr + 16 + 24, 0);
        } else {
          A.store.Bool(ptr + 16 + 36, true);
          A.store.Bool(ptr + 16 + 32, "left" in x["boundsInRoot"] ? true : false);
          A.store.Float64(
            ptr + 16 + 0,
            x["boundsInRoot"]["left"] === undefined ? 0 : (x["boundsInRoot"]["left"] as number)
          );
          A.store.Bool(ptr + 16 + 33, "top" in x["boundsInRoot"] ? true : false);
          A.store.Float64(
            ptr + 16 + 8,
            x["boundsInRoot"]["top"] === undefined ? 0 : (x["boundsInRoot"]["top"] as number)
          );
          A.store.Bool(ptr + 16 + 34, "width" in x["boundsInRoot"] ? true : false);
          A.store.Float64(
            ptr + 16 + 16,
            x["boundsInRoot"]["width"] === undefined ? 0 : (x["boundsInRoot"]["width"] as number)
          );
          A.store.Bool(ptr + 16 + 35, "height" in x["boundsInRoot"] ? true : false);
          A.store.Float64(
            ptr + 16 + 24,
            x["boundsInRoot"]["height"] === undefined ? 0 : (x["boundsInRoot"]["height"] as number)
          );
        }
        A.store.Ref(ptr + 56, x["displayId"]);
        A.store.Bool(ptr + 193, "isVisible" in x ? true : false);
        A.store.Bool(ptr + 60, x["isVisible"] ? true : false);
        A.store.Bool(ptr + 194, "canFocus" in x ? true : false);
        A.store.Bool(ptr + 61, x["canFocus"] ? true : false);
        A.store.Ref(ptr + 64, x["title"]);
        A.store.Bool(ptr + 195, "isAnimating" in x ? true : false);
        A.store.Bool(ptr + 68, x["isAnimating"] ? true : false);

        if (typeof x["targetBounds"] === "undefined") {
          A.store.Bool(ptr + 72 + 36, false);
          A.store.Bool(ptr + 72 + 32, false);
          A.store.Float64(ptr + 72 + 0, 0);
          A.store.Bool(ptr + 72 + 33, false);
          A.store.Float64(ptr + 72 + 8, 0);
          A.store.Bool(ptr + 72 + 34, false);
          A.store.Float64(ptr + 72 + 16, 0);
          A.store.Bool(ptr + 72 + 35, false);
          A.store.Float64(ptr + 72 + 24, 0);
        } else {
          A.store.Bool(ptr + 72 + 36, true);
          A.store.Bool(ptr + 72 + 32, "left" in x["targetBounds"] ? true : false);
          A.store.Float64(
            ptr + 72 + 0,
            x["targetBounds"]["left"] === undefined ? 0 : (x["targetBounds"]["left"] as number)
          );
          A.store.Bool(ptr + 72 + 33, "top" in x["targetBounds"] ? true : false);
          A.store.Float64(
            ptr + 72 + 8,
            x["targetBounds"]["top"] === undefined ? 0 : (x["targetBounds"]["top"] as number)
          );
          A.store.Bool(ptr + 72 + 34, "width" in x["targetBounds"] ? true : false);
          A.store.Float64(
            ptr + 72 + 16,
            x["targetBounds"]["width"] === undefined ? 0 : (x["targetBounds"]["width"] as number)
          );
          A.store.Bool(ptr + 72 + 35, "height" in x["targetBounds"] ? true : false);
          A.store.Float64(
            ptr + 72 + 24,
            x["targetBounds"]["height"] === undefined ? 0 : (x["targetBounds"]["height"] as number)
          );
        }
        A.store.Bool(ptr + 196, "targetVisibility" in x ? true : false);
        A.store.Bool(ptr + 109, x["targetVisibility"] ? true : false);
        A.store.Bool(ptr + 197, "isActive" in x ? true : false);
        A.store.Bool(ptr + 110, x["isActive"] ? true : false);
        A.store.Bool(ptr + 198, "hasFocus" in x ? true : false);
        A.store.Bool(ptr + 111, x["hasFocus"] ? true : false);
        A.store.Bool(ptr + 199, "onActiveDesk" in x ? true : false);
        A.store.Bool(ptr + 112, x["onActiveDesk"] ? true : false);
        A.store.Bool(ptr + 200, "hasCapture" in x ? true : false);
        A.store.Bool(ptr + 113, x["hasCapture"] ? true : false);
        A.store.Bool(ptr + 201, "canResize" in x ? true : false);
        A.store.Bool(ptr + 114, x["canResize"] ? true : false);
        A.store.Bool(ptr + 202, "stackingOrder" in x ? true : false);
        A.store.Int32(ptr + 116, x["stackingOrder"] === undefined ? 0 : (x["stackingOrder"] as number));
        A.store.Enum(ptr + 120, ["Normal", "Immersive"].indexOf(x["frameMode"] as string));
        A.store.Bool(ptr + 203, "isFrameVisible" in x ? true : false);
        A.store.Bool(ptr + 124, x["isFrameVisible"] ? true : false);
        A.store.Bool(ptr + 204, "captionHeight" in x ? true : false);
        A.store.Int32(ptr + 128, x["captionHeight"] === undefined ? 0 : (x["captionHeight"] as number));
        A.store.Bool(ptr + 205, "captionButtonEnabledStatus" in x ? true : false);
        A.store.Int32(
          ptr + 132,
          x["captionButtonEnabledStatus"] === undefined ? 0 : (x["captionButtonEnabledStatus"] as number)
        );
        A.store.Bool(ptr + 206, "captionButtonVisibleStatus" in x ? true : false);
        A.store.Int32(
          ptr + 136,
          x["captionButtonVisibleStatus"] === undefined ? 0 : (x["captionButtonVisibleStatus"] as number)
        );
        A.store.Ref(ptr + 140, x["arcPackageName"]);

        if (typeof x["overviewInfo"] === "undefined") {
          A.store.Bool(ptr + 144 + 39, false);

          A.store.Bool(ptr + 144 + 0 + 36, false);
          A.store.Bool(ptr + 144 + 0 + 32, false);
          A.store.Float64(ptr + 144 + 0 + 0, 0);
          A.store.Bool(ptr + 144 + 0 + 33, false);
          A.store.Float64(ptr + 144 + 0 + 8, 0);
          A.store.Bool(ptr + 144 + 0 + 34, false);
          A.store.Float64(ptr + 144 + 0 + 16, 0);
          A.store.Bool(ptr + 144 + 0 + 35, false);
          A.store.Float64(ptr + 144 + 0 + 24, 0);
          A.store.Bool(ptr + 144 + 38, false);
          A.store.Bool(ptr + 144 + 37, false);
        } else {
          A.store.Bool(ptr + 144 + 39, true);

          if (typeof x["overviewInfo"]["bounds"] === "undefined") {
            A.store.Bool(ptr + 144 + 0 + 36, false);
            A.store.Bool(ptr + 144 + 0 + 32, false);
            A.store.Float64(ptr + 144 + 0 + 0, 0);
            A.store.Bool(ptr + 144 + 0 + 33, false);
            A.store.Float64(ptr + 144 + 0 + 8, 0);
            A.store.Bool(ptr + 144 + 0 + 34, false);
            A.store.Float64(ptr + 144 + 0 + 16, 0);
            A.store.Bool(ptr + 144 + 0 + 35, false);
            A.store.Float64(ptr + 144 + 0 + 24, 0);
          } else {
            A.store.Bool(ptr + 144 + 0 + 36, true);
            A.store.Bool(ptr + 144 + 0 + 32, "left" in x["overviewInfo"]["bounds"] ? true : false);
            A.store.Float64(
              ptr + 144 + 0 + 0,
              x["overviewInfo"]["bounds"]["left"] === undefined ? 0 : (x["overviewInfo"]["bounds"]["left"] as number)
            );
            A.store.Bool(ptr + 144 + 0 + 33, "top" in x["overviewInfo"]["bounds"] ? true : false);
            A.store.Float64(
              ptr + 144 + 0 + 8,
              x["overviewInfo"]["bounds"]["top"] === undefined ? 0 : (x["overviewInfo"]["bounds"]["top"] as number)
            );
            A.store.Bool(ptr + 144 + 0 + 34, "width" in x["overviewInfo"]["bounds"] ? true : false);
            A.store.Float64(
              ptr + 144 + 0 + 16,
              x["overviewInfo"]["bounds"]["width"] === undefined ? 0 : (x["overviewInfo"]["bounds"]["width"] as number)
            );
            A.store.Bool(ptr + 144 + 0 + 35, "height" in x["overviewInfo"]["bounds"] ? true : false);
            A.store.Float64(
              ptr + 144 + 0 + 24,
              x["overviewInfo"]["bounds"]["height"] === undefined
                ? 0
                : (x["overviewInfo"]["bounds"]["height"] as number)
            );
          }
          A.store.Bool(ptr + 144 + 38, "isDragged" in x["overviewInfo"] ? true : false);
          A.store.Bool(ptr + 144 + 37, x["overviewInfo"]["isDragged"] ? true : false);
        }
        A.store.Ref(ptr + 184, x["fullRestoreWindowAppId"]);
        A.store.Ref(ptr + 188, x["appId"]);
      }
    },
    "load_AppWindowInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 192)) {
        x["id"] = A.load.Int32(ptr + 0);
      } else {
        delete x["id"];
      }
      x["name"] = A.load.Ref(ptr + 4, undefined);
      x["windowType"] = A.load.Enum(ptr + 8, [
        "Browser",
        "ChromeApp",
        "ArcApp",
        "CrostiniApp",
        "SystemApp",
        "ExtensionApp",
        "Lacros",
      ]);
      x["stateType"] = A.load.Enum(ptr + 12, [
        "Normal",
        "Minimized",
        "Maximized",
        "Fullscreen",
        "PrimarySnapped",
        "SecondarySnapped",
        "PIP",
        "Floated",
      ]);
      if (A.load.Bool(ptr + 16 + 36)) {
        x["boundsInRoot"] = {};
        if (A.load.Bool(ptr + 16 + 32)) {
          x["boundsInRoot"]["left"] = A.load.Float64(ptr + 16 + 0);
        } else {
          delete x["boundsInRoot"]["left"];
        }
        if (A.load.Bool(ptr + 16 + 33)) {
          x["boundsInRoot"]["top"] = A.load.Float64(ptr + 16 + 8);
        } else {
          delete x["boundsInRoot"]["top"];
        }
        if (A.load.Bool(ptr + 16 + 34)) {
          x["boundsInRoot"]["width"] = A.load.Float64(ptr + 16 + 16);
        } else {
          delete x["boundsInRoot"]["width"];
        }
        if (A.load.Bool(ptr + 16 + 35)) {
          x["boundsInRoot"]["height"] = A.load.Float64(ptr + 16 + 24);
        } else {
          delete x["boundsInRoot"]["height"];
        }
      } else {
        delete x["boundsInRoot"];
      }
      x["displayId"] = A.load.Ref(ptr + 56, undefined);
      if (A.load.Bool(ptr + 193)) {
        x["isVisible"] = A.load.Bool(ptr + 60);
      } else {
        delete x["isVisible"];
      }
      if (A.load.Bool(ptr + 194)) {
        x["canFocus"] = A.load.Bool(ptr + 61);
      } else {
        delete x["canFocus"];
      }
      x["title"] = A.load.Ref(ptr + 64, undefined);
      if (A.load.Bool(ptr + 195)) {
        x["isAnimating"] = A.load.Bool(ptr + 68);
      } else {
        delete x["isAnimating"];
      }
      if (A.load.Bool(ptr + 72 + 36)) {
        x["targetBounds"] = {};
        if (A.load.Bool(ptr + 72 + 32)) {
          x["targetBounds"]["left"] = A.load.Float64(ptr + 72 + 0);
        } else {
          delete x["targetBounds"]["left"];
        }
        if (A.load.Bool(ptr + 72 + 33)) {
          x["targetBounds"]["top"] = A.load.Float64(ptr + 72 + 8);
        } else {
          delete x["targetBounds"]["top"];
        }
        if (A.load.Bool(ptr + 72 + 34)) {
          x["targetBounds"]["width"] = A.load.Float64(ptr + 72 + 16);
        } else {
          delete x["targetBounds"]["width"];
        }
        if (A.load.Bool(ptr + 72 + 35)) {
          x["targetBounds"]["height"] = A.load.Float64(ptr + 72 + 24);
        } else {
          delete x["targetBounds"]["height"];
        }
      } else {
        delete x["targetBounds"];
      }
      if (A.load.Bool(ptr + 196)) {
        x["targetVisibility"] = A.load.Bool(ptr + 109);
      } else {
        delete x["targetVisibility"];
      }
      if (A.load.Bool(ptr + 197)) {
        x["isActive"] = A.load.Bool(ptr + 110);
      } else {
        delete x["isActive"];
      }
      if (A.load.Bool(ptr + 198)) {
        x["hasFocus"] = A.load.Bool(ptr + 111);
      } else {
        delete x["hasFocus"];
      }
      if (A.load.Bool(ptr + 199)) {
        x["onActiveDesk"] = A.load.Bool(ptr + 112);
      } else {
        delete x["onActiveDesk"];
      }
      if (A.load.Bool(ptr + 200)) {
        x["hasCapture"] = A.load.Bool(ptr + 113);
      } else {
        delete x["hasCapture"];
      }
      if (A.load.Bool(ptr + 201)) {
        x["canResize"] = A.load.Bool(ptr + 114);
      } else {
        delete x["canResize"];
      }
      if (A.load.Bool(ptr + 202)) {
        x["stackingOrder"] = A.load.Int32(ptr + 116);
      } else {
        delete x["stackingOrder"];
      }
      x["frameMode"] = A.load.Enum(ptr + 120, ["Normal", "Immersive"]);
      if (A.load.Bool(ptr + 203)) {
        x["isFrameVisible"] = A.load.Bool(ptr + 124);
      } else {
        delete x["isFrameVisible"];
      }
      if (A.load.Bool(ptr + 204)) {
        x["captionHeight"] = A.load.Int32(ptr + 128);
      } else {
        delete x["captionHeight"];
      }
      if (A.load.Bool(ptr + 205)) {
        x["captionButtonEnabledStatus"] = A.load.Int32(ptr + 132);
      } else {
        delete x["captionButtonEnabledStatus"];
      }
      if (A.load.Bool(ptr + 206)) {
        x["captionButtonVisibleStatus"] = A.load.Int32(ptr + 136);
      } else {
        delete x["captionButtonVisibleStatus"];
      }
      x["arcPackageName"] = A.load.Ref(ptr + 140, undefined);
      if (A.load.Bool(ptr + 144 + 39)) {
        x["overviewInfo"] = {};
        if (A.load.Bool(ptr + 144 + 0 + 36)) {
          x["overviewInfo"]["bounds"] = {};
          if (A.load.Bool(ptr + 144 + 0 + 32)) {
            x["overviewInfo"]["bounds"]["left"] = A.load.Float64(ptr + 144 + 0 + 0);
          } else {
            delete x["overviewInfo"]["bounds"]["left"];
          }
          if (A.load.Bool(ptr + 144 + 0 + 33)) {
            x["overviewInfo"]["bounds"]["top"] = A.load.Float64(ptr + 144 + 0 + 8);
          } else {
            delete x["overviewInfo"]["bounds"]["top"];
          }
          if (A.load.Bool(ptr + 144 + 0 + 34)) {
            x["overviewInfo"]["bounds"]["width"] = A.load.Float64(ptr + 144 + 0 + 16);
          } else {
            delete x["overviewInfo"]["bounds"]["width"];
          }
          if (A.load.Bool(ptr + 144 + 0 + 35)) {
            x["overviewInfo"]["bounds"]["height"] = A.load.Float64(ptr + 144 + 0 + 24);
          } else {
            delete x["overviewInfo"]["bounds"]["height"];
          }
        } else {
          delete x["overviewInfo"]["bounds"];
        }
        if (A.load.Bool(ptr + 144 + 38)) {
          x["overviewInfo"]["isDragged"] = A.load.Bool(ptr + 144 + 37);
        } else {
          delete x["overviewInfo"]["isDragged"];
        }
      } else {
        delete x["overviewInfo"];
      }
      x["fullRestoreWindowAppId"] = A.load.Ref(ptr + 184, undefined);
      x["appId"] = A.load.Ref(ptr + 188, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ArcAppDict": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 56, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 47, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Bool(ptr + 48, false);
        A.store.Float64(ptr + 32, 0);
        A.store.Bool(ptr + 49, false);
        A.store.Bool(ptr + 40, false);
        A.store.Bool(ptr + 50, false);
        A.store.Bool(ptr + 41, false);
        A.store.Bool(ptr + 51, false);
        A.store.Bool(ptr + 42, false);
        A.store.Bool(ptr + 52, false);
        A.store.Bool(ptr + 43, false);
        A.store.Bool(ptr + 53, false);
        A.store.Bool(ptr + 44, false);
        A.store.Bool(ptr + 54, false);
        A.store.Bool(ptr + 45, false);
        A.store.Bool(ptr + 55, false);
        A.store.Bool(ptr + 46, false);
      } else {
        A.store.Bool(ptr + 56, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Ref(ptr + 4, x["packageName"]);
        A.store.Ref(ptr + 8, x["activity"]);
        A.store.Ref(ptr + 12, x["intentUri"]);
        A.store.Ref(ptr + 16, x["iconResourceId"]);
        A.store.Bool(ptr + 47, "lastLaunchTime" in x ? true : false);
        A.store.Float64(ptr + 24, x["lastLaunchTime"] === undefined ? 0 : (x["lastLaunchTime"] as number));
        A.store.Bool(ptr + 48, "installTime" in x ? true : false);
        A.store.Float64(ptr + 32, x["installTime"] === undefined ? 0 : (x["installTime"] as number));
        A.store.Bool(ptr + 49, "sticky" in x ? true : false);
        A.store.Bool(ptr + 40, x["sticky"] ? true : false);
        A.store.Bool(ptr + 50, "notificationsEnabled" in x ? true : false);
        A.store.Bool(ptr + 41, x["notificationsEnabled"] ? true : false);
        A.store.Bool(ptr + 51, "ready" in x ? true : false);
        A.store.Bool(ptr + 42, x["ready"] ? true : false);
        A.store.Bool(ptr + 52, "suspended" in x ? true : false);
        A.store.Bool(ptr + 43, x["suspended"] ? true : false);
        A.store.Bool(ptr + 53, "showInLauncher" in x ? true : false);
        A.store.Bool(ptr + 44, x["showInLauncher"] ? true : false);
        A.store.Bool(ptr + 54, "shortcut" in x ? true : false);
        A.store.Bool(ptr + 45, x["shortcut"] ? true : false);
        A.store.Bool(ptr + 55, "launchable" in x ? true : false);
        A.store.Bool(ptr + 46, x["launchable"] ? true : false);
      }
    },
    "load_ArcAppDict": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      x["packageName"] = A.load.Ref(ptr + 4, undefined);
      x["activity"] = A.load.Ref(ptr + 8, undefined);
      x["intentUri"] = A.load.Ref(ptr + 12, undefined);
      x["iconResourceId"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 47)) {
        x["lastLaunchTime"] = A.load.Float64(ptr + 24);
      } else {
        delete x["lastLaunchTime"];
      }
      if (A.load.Bool(ptr + 48)) {
        x["installTime"] = A.load.Float64(ptr + 32);
      } else {
        delete x["installTime"];
      }
      if (A.load.Bool(ptr + 49)) {
        x["sticky"] = A.load.Bool(ptr + 40);
      } else {
        delete x["sticky"];
      }
      if (A.load.Bool(ptr + 50)) {
        x["notificationsEnabled"] = A.load.Bool(ptr + 41);
      } else {
        delete x["notificationsEnabled"];
      }
      if (A.load.Bool(ptr + 51)) {
        x["ready"] = A.load.Bool(ptr + 42);
      } else {
        delete x["ready"];
      }
      if (A.load.Bool(ptr + 52)) {
        x["suspended"] = A.load.Bool(ptr + 43);
      } else {
        delete x["suspended"];
      }
      if (A.load.Bool(ptr + 53)) {
        x["showInLauncher"] = A.load.Bool(ptr + 44);
      } else {
        delete x["showInLauncher"];
      }
      if (A.load.Bool(ptr + 54)) {
        x["shortcut"] = A.load.Bool(ptr + 45);
      } else {
        delete x["shortcut"];
      }
      if (A.load.Bool(ptr + 55)) {
        x["launchable"] = A.load.Bool(ptr + 46);
      } else {
        delete x["launchable"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ArcAppKillsDict": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 63, false);
        A.store.Bool(ptr + 56, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 57, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 58, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 59, false);
        A.store.Float64(ptr + 24, 0);
        A.store.Bool(ptr + 60, false);
        A.store.Float64(ptr + 32, 0);
        A.store.Bool(ptr + 61, false);
        A.store.Float64(ptr + 40, 0);
        A.store.Bool(ptr + 62, false);
        A.store.Float64(ptr + 48, 0);
      } else {
        A.store.Bool(ptr + 63, true);
        A.store.Bool(ptr + 56, "oom" in x ? true : false);
        A.store.Float64(ptr + 0, x["oom"] === undefined ? 0 : (x["oom"] as number));
        A.store.Bool(ptr + 57, "lmkdForeground" in x ? true : false);
        A.store.Float64(ptr + 8, x["lmkdForeground"] === undefined ? 0 : (x["lmkdForeground"] as number));
        A.store.Bool(ptr + 58, "lmkdPerceptible" in x ? true : false);
        A.store.Float64(ptr + 16, x["lmkdPerceptible"] === undefined ? 0 : (x["lmkdPerceptible"] as number));
        A.store.Bool(ptr + 59, "lmkdCached" in x ? true : false);
        A.store.Float64(ptr + 24, x["lmkdCached"] === undefined ? 0 : (x["lmkdCached"] as number));
        A.store.Bool(ptr + 60, "pressureForeground" in x ? true : false);
        A.store.Float64(ptr + 32, x["pressureForeground"] === undefined ? 0 : (x["pressureForeground"] as number));
        A.store.Bool(ptr + 61, "pressurePerceptible" in x ? true : false);
        A.store.Float64(ptr + 40, x["pressurePerceptible"] === undefined ? 0 : (x["pressurePerceptible"] as number));
        A.store.Bool(ptr + 62, "pressureCached" in x ? true : false);
        A.store.Float64(ptr + 48, x["pressureCached"] === undefined ? 0 : (x["pressureCached"] as number));
      }
    },
    "load_ArcAppKillsDict": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 56)) {
        x["oom"] = A.load.Float64(ptr + 0);
      } else {
        delete x["oom"];
      }
      if (A.load.Bool(ptr + 57)) {
        x["lmkdForeground"] = A.load.Float64(ptr + 8);
      } else {
        delete x["lmkdForeground"];
      }
      if (A.load.Bool(ptr + 58)) {
        x["lmkdPerceptible"] = A.load.Float64(ptr + 16);
      } else {
        delete x["lmkdPerceptible"];
      }
      if (A.load.Bool(ptr + 59)) {
        x["lmkdCached"] = A.load.Float64(ptr + 24);
      } else {
        delete x["lmkdCached"];
      }
      if (A.load.Bool(ptr + 60)) {
        x["pressureForeground"] = A.load.Float64(ptr + 32);
      } else {
        delete x["pressureForeground"];
      }
      if (A.load.Bool(ptr + 61)) {
        x["pressurePerceptible"] = A.load.Float64(ptr + 40);
      } else {
        delete x["pressurePerceptible"];
      }
      if (A.load.Bool(ptr + 62)) {
        x["pressureCached"] = A.load.Float64(ptr + 48);
      } else {
        delete x["pressureCached"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ArcAppTracingInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 36, false);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 33, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 34, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 35, false);
        A.store.Float64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 36, true);
        A.store.Bool(ptr + 32, "success" in x ? true : false);
        A.store.Bool(ptr + 0, x["success"] ? true : false);
        A.store.Bool(ptr + 33, "fps" in x ? true : false);
        A.store.Float64(ptr + 8, x["fps"] === undefined ? 0 : (x["fps"] as number));
        A.store.Bool(ptr + 34, "commitDeviation" in x ? true : false);
        A.store.Float64(ptr + 16, x["commitDeviation"] === undefined ? 0 : (x["commitDeviation"] as number));
        A.store.Bool(ptr + 35, "renderQuality" in x ? true : false);
        A.store.Float64(ptr + 24, x["renderQuality"] === undefined ? 0 : (x["renderQuality"] as number));
      }
    },
    "load_ArcAppTracingInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 32)) {
        x["success"] = A.load.Bool(ptr + 0);
      } else {
        delete x["success"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["fps"] = A.load.Float64(ptr + 8);
      } else {
        delete x["fps"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["commitDeviation"] = A.load.Float64(ptr + 16);
      } else {
        delete x["commitDeviation"];
      }
      if (A.load.Bool(ptr + 35)) {
        x["renderQuality"] = A.load.Float64(ptr + 24);
      } else {
        delete x["renderQuality"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ArcPackageDict": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 30, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 26, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 27, false);
        A.store.Float64(ptr + 16, 0);
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 29, false);
        A.store.Bool(ptr + 25, false);
      } else {
        A.store.Bool(ptr + 30, true);
        A.store.Ref(ptr + 0, x["packageName"]);
        A.store.Bool(ptr + 26, "packageVersion" in x ? true : false);
        A.store.Int32(ptr + 4, x["packageVersion"] === undefined ? 0 : (x["packageVersion"] as number));
        A.store.Ref(ptr + 8, x["lastBackupAndroidId"]);
        A.store.Bool(ptr + 27, "lastBackupTime" in x ? true : false);
        A.store.Float64(ptr + 16, x["lastBackupTime"] === undefined ? 0 : (x["lastBackupTime"] as number));
        A.store.Bool(ptr + 28, "shouldSync" in x ? true : false);
        A.store.Bool(ptr + 24, x["shouldSync"] ? true : false);
        A.store.Bool(ptr + 29, "vpnProvider" in x ? true : false);
        A.store.Bool(ptr + 25, x["vpnProvider"] ? true : false);
      }
    },
    "load_ArcPackageDict": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["packageName"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 26)) {
        x["packageVersion"] = A.load.Int32(ptr + 4);
      } else {
        delete x["packageVersion"];
      }
      x["lastBackupAndroidId"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 27)) {
        x["lastBackupTime"] = A.load.Float64(ptr + 16);
      } else {
        delete x["lastBackupTime"];
      }
      if (A.load.Bool(ptr + 28)) {
        x["shouldSync"] = A.load.Bool(ptr + 24);
      } else {
        delete x["shouldSync"];
      }
      if (A.load.Bool(ptr + 29)) {
        x["vpnProvider"] = A.load.Bool(ptr + 25);
      } else {
        delete x["vpnProvider"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ArcState": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 26, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 27, false);
        A.store.Float64(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 28, true);
        A.store.Bool(ptr + 24, "provisioned" in x ? true : false);
        A.store.Bool(ptr + 0, x["provisioned"] ? true : false);
        A.store.Bool(ptr + 25, "tosNeeded" in x ? true : false);
        A.store.Bool(ptr + 1, x["tosNeeded"] ? true : false);
        A.store.Bool(ptr + 26, "preStartTime" in x ? true : false);
        A.store.Float64(ptr + 8, x["preStartTime"] === undefined ? 0 : (x["preStartTime"] as number));
        A.store.Bool(ptr + 27, "startTime" in x ? true : false);
        A.store.Float64(ptr + 16, x["startTime"] === undefined ? 0 : (x["startTime"] as number));
      }
    },
    "load_ArcState": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 24)) {
        x["provisioned"] = A.load.Bool(ptr + 0);
      } else {
        delete x["provisioned"];
      }
      if (A.load.Bool(ptr + 25)) {
        x["tosNeeded"] = A.load.Bool(ptr + 1);
      } else {
        delete x["tosNeeded"];
      }
      if (A.load.Bool(ptr + 26)) {
        x["preStartTime"] = A.load.Float64(ptr + 8);
      } else {
        delete x["preStartTime"];
      }
      if (A.load.Bool(ptr + 27)) {
        x["startTime"] = A.load.Float64(ptr + 16);
      } else {
        delete x["startTime"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AssistantQueryResponse": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["text"]);
        A.store.Ref(ptr + 4, x["htmlResponse"]);
        A.store.Ref(ptr + 8, x["openUrl"]);
        A.store.Ref(ptr + 12, x["openAppResponse"]);
      }
    },
    "load_AssistantQueryResponse": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["text"] = A.load.Ref(ptr + 0, undefined);
      x["htmlResponse"] = A.load.Ref(ptr + 4, undefined);
      x["openUrl"] = A.load.Ref(ptr + 8, undefined);
      x["openAppResponse"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AssistantQueryStatus": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);

        A.store.Bool(ptr + 8 + 16, false);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Ref(ptr + 8 + 4, undefined);
        A.store.Ref(ptr + 8 + 8, undefined);
        A.store.Ref(ptr + 8 + 12, undefined);
      } else {
        A.store.Bool(ptr + 26, true);
        A.store.Bool(ptr + 25, "isMicOpen" in x ? true : false);
        A.store.Bool(ptr + 0, x["isMicOpen"] ? true : false);
        A.store.Ref(ptr + 4, x["queryText"]);

        if (typeof x["queryResponse"] === "undefined") {
          A.store.Bool(ptr + 8 + 16, false);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Ref(ptr + 8 + 4, undefined);
          A.store.Ref(ptr + 8 + 8, undefined);
          A.store.Ref(ptr + 8 + 12, undefined);
        } else {
          A.store.Bool(ptr + 8 + 16, true);
          A.store.Ref(ptr + 8 + 0, x["queryResponse"]["text"]);
          A.store.Ref(ptr + 8 + 4, x["queryResponse"]["htmlResponse"]);
          A.store.Ref(ptr + 8 + 8, x["queryResponse"]["openUrl"]);
          A.store.Ref(ptr + 8 + 12, x["queryResponse"]["openAppResponse"]);
        }
      }
    },
    "load_AssistantQueryStatus": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 25)) {
        x["isMicOpen"] = A.load.Bool(ptr + 0);
      } else {
        delete x["isMicOpen"];
      }
      x["queryText"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 8 + 16)) {
        x["queryResponse"] = {};
        x["queryResponse"]["text"] = A.load.Ref(ptr + 8 + 0, undefined);
        x["queryResponse"]["htmlResponse"] = A.load.Ref(ptr + 8 + 4, undefined);
        x["queryResponse"]["openUrl"] = A.load.Ref(ptr + 8 + 8, undefined);
        x["queryResponse"]["openAppResponse"] = A.load.Ref(ptr + 8 + 12, undefined);
      } else {
        delete x["queryResponse"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CryptohomeRecoveryDataDict": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["reauthProofToken"]);
        A.store.Ref(ptr + 4, x["refreshToken"]);
      }
    },
    "load_CryptohomeRecoveryDataDict": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["reauthProofToken"] = A.load.Ref(ptr + 0, undefined);
      x["refreshToken"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DesksInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 19, true);
        A.store.Bool(ptr + 16, "activeDeskIndex" in x ? true : false);
        A.store.Int32(ptr + 0, x["activeDeskIndex"] === undefined ? 0 : (x["activeDeskIndex"] as number));
        A.store.Bool(ptr + 17, "numDesks" in x ? true : false);
        A.store.Int32(ptr + 4, x["numDesks"] === undefined ? 0 : (x["numDesks"] as number));
        A.store.Bool(ptr + 18, "isAnimating" in x ? true : false);
        A.store.Bool(ptr + 8, x["isAnimating"] ? true : false);
        A.store.Ref(ptr + 12, x["deskContainers"]);
      }
    },
    "load_DesksInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["activeDeskIndex"] = A.load.Int32(ptr + 0);
      } else {
        delete x["activeDeskIndex"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["numDesks"] = A.load.Int32(ptr + 4);
      } else {
        delete x["numDesks"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["isAnimating"] = A.load.Bool(ptr + 8);
      } else {
        delete x["isAnimating"];
      }
      x["deskContainers"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DisplaySmoothnessData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 18, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 19, true);
        A.store.Bool(ptr + 16, "framesExpected" in x ? true : false);
        A.store.Int32(ptr + 0, x["framesExpected"] === undefined ? 0 : (x["framesExpected"] as number));
        A.store.Bool(ptr + 17, "framesProduced" in x ? true : false);
        A.store.Int32(ptr + 4, x["framesProduced"] === undefined ? 0 : (x["framesProduced"] as number));
        A.store.Bool(ptr + 18, "jankCount" in x ? true : false);
        A.store.Int32(ptr + 8, x["jankCount"] === undefined ? 0 : (x["jankCount"] as number));
        A.store.Ref(ptr + 12, x["throughput"]);
      }
    },
    "load_DisplaySmoothnessData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["framesExpected"] = A.load.Int32(ptr + 0);
      } else {
        delete x["framesExpected"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["framesProduced"] = A.load.Int32(ptr + 4);
      } else {
        delete x["framesProduced"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["jankCount"] = A.load.Int32(ptr + 8);
      } else {
        delete x["jankCount"];
      }
      x["throughput"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ExtensionInfoDict": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 52, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Ref(ptr + 28, undefined);
        A.store.Ref(ptr + 32, undefined);
        A.store.Ref(ptr + 36, undefined);
        A.store.Bool(ptr + 46, false);
        A.store.Bool(ptr + 40, false);
        A.store.Bool(ptr + 47, false);
        A.store.Bool(ptr + 41, false);
        A.store.Bool(ptr + 48, false);
        A.store.Bool(ptr + 42, false);
        A.store.Bool(ptr + 49, false);
        A.store.Bool(ptr + 43, false);
        A.store.Bool(ptr + 50, false);
        A.store.Bool(ptr + 44, false);
        A.store.Bool(ptr + 51, false);
        A.store.Bool(ptr + 45, false);
      } else {
        A.store.Bool(ptr + 52, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["version"]);
        A.store.Ref(ptr + 8, x["name"]);
        A.store.Ref(ptr + 12, x["publicKey"]);
        A.store.Ref(ptr + 16, x["description"]);
        A.store.Ref(ptr + 20, x["backgroundUrl"]);
        A.store.Ref(ptr + 24, x["optionsUrl"]);
        A.store.Ref(ptr + 28, x["hostPermissions"]);
        A.store.Ref(ptr + 32, x["effectiveHostPermissions"]);
        A.store.Ref(ptr + 36, x["apiPermissions"]);
        A.store.Bool(ptr + 46, "isComponent" in x ? true : false);
        A.store.Bool(ptr + 40, x["isComponent"] ? true : false);
        A.store.Bool(ptr + 47, "isInternal" in x ? true : false);
        A.store.Bool(ptr + 41, x["isInternal"] ? true : false);
        A.store.Bool(ptr + 48, "isUserInstalled" in x ? true : false);
        A.store.Bool(ptr + 42, x["isUserInstalled"] ? true : false);
        A.store.Bool(ptr + 49, "isEnabled" in x ? true : false);
        A.store.Bool(ptr + 43, x["isEnabled"] ? true : false);
        A.store.Bool(ptr + 50, "allowedInIncognito" in x ? true : false);
        A.store.Bool(ptr + 44, x["allowedInIncognito"] ? true : false);
        A.store.Bool(ptr + 51, "hasPageAction" in x ? true : false);
        A.store.Bool(ptr + 45, x["hasPageAction"] ? true : false);
      }
    },
    "load_ExtensionInfoDict": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["version"] = A.load.Ref(ptr + 4, undefined);
      x["name"] = A.load.Ref(ptr + 8, undefined);
      x["publicKey"] = A.load.Ref(ptr + 12, undefined);
      x["description"] = A.load.Ref(ptr + 16, undefined);
      x["backgroundUrl"] = A.load.Ref(ptr + 20, undefined);
      x["optionsUrl"] = A.load.Ref(ptr + 24, undefined);
      x["hostPermissions"] = A.load.Ref(ptr + 28, undefined);
      x["effectiveHostPermissions"] = A.load.Ref(ptr + 32, undefined);
      x["apiPermissions"] = A.load.Ref(ptr + 36, undefined);
      if (A.load.Bool(ptr + 46)) {
        x["isComponent"] = A.load.Bool(ptr + 40);
      } else {
        delete x["isComponent"];
      }
      if (A.load.Bool(ptr + 47)) {
        x["isInternal"] = A.load.Bool(ptr + 41);
      } else {
        delete x["isInternal"];
      }
      if (A.load.Bool(ptr + 48)) {
        x["isUserInstalled"] = A.load.Bool(ptr + 42);
      } else {
        delete x["isUserInstalled"];
      }
      if (A.load.Bool(ptr + 49)) {
        x["isEnabled"] = A.load.Bool(ptr + 43);
      } else {
        delete x["isEnabled"];
      }
      if (A.load.Bool(ptr + 50)) {
        x["allowedInIncognito"] = A.load.Bool(ptr + 44);
      } else {
        delete x["allowedInIncognito"];
      }
      if (A.load.Bool(ptr + 51)) {
        x["hasPageAction"] = A.load.Bool(ptr + 45);
      } else {
        delete x["hasPageAction"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ExtensionsInfoArray": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["extensions"]);
      }
    },
    "load_ExtensionsInfoArray": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["extensions"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_FrameCountingPerSinkData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Ref(ptr + 0, x["sinkType"]);
        A.store.Bool(ptr + 16, "isRoot" in x ? true : false);
        A.store.Bool(ptr + 4, x["isRoot"] ? true : false);
        A.store.Ref(ptr + 8, x["debugLabel"]);
        A.store.Ref(ptr + 12, x["presentedFrames"]);
      }
    },
    "load_FrameCountingPerSinkData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["sinkType"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["isRoot"] = A.load.Bool(ptr + 4);
      } else {
        delete x["isRoot"];
      }
      x["debugLabel"] = A.load.Ref(ptr + 8, undefined);
      x["presentedFrames"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetAccessTokenData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["accessToken"]);
        A.store.Ref(ptr + 4, x["expirationTimeUnixMs"]);
      }
    },
    "load_GetAccessTokenData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["accessToken"] = A.load.Ref(ptr + 0, undefined);
      x["expirationTimeUnixMs"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetAccessTokenParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 13, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 13, true);
        A.store.Ref(ptr + 0, x["email"]);
        A.store.Ref(ptr + 4, x["scopes"]);
        A.store.Bool(ptr + 12, "timeoutMs" in x ? true : false);
        A.store.Int32(ptr + 8, x["timeoutMs"] === undefined ? 0 : (x["timeoutMs"] as number));
      }
    },
    "load_GetAccessTokenParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["email"] = A.load.Ref(ptr + 0, undefined);
      x["scopes"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 12)) {
        x["timeoutMs"] = A.load.Int32(ptr + 8);
      } else {
        delete x["timeoutMs"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetCurrentInputMethodDescriptorData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["keyboardLayout"]);
      }
    },
    "load_GetCurrentInputMethodDescriptorData": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["keyboardLayout"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_LacrosState": (ref: heap.Ref<string>): number => {
      const idx = [
        "NotInitialized",
        "Reloading",
        "Mounting",
        "Unavailable",
        "Stopped",
        "CreatingLogFile",
        "PreLaunched",
        "Starting",
        "Running",
        "Terminating",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_LacrosMode": (ref: heap.Ref<string>): number => {
      const idx = ["Disabled", "Only"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_LacrosInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 17, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 8, undefined);
        A.store.Enum(ptr + 12, -1);
      } else {
        A.store.Bool(ptr + 17, true);
        A.store.Enum(
          ptr + 0,
          [
            "NotInitialized",
            "Reloading",
            "Mounting",
            "Unavailable",
            "Stopped",
            "CreatingLogFile",
            "PreLaunched",
            "Starting",
            "Running",
            "Terminating",
          ].indexOf(x["state"] as string)
        );
        A.store.Bool(ptr + 16, "isKeepAlive" in x ? true : false);
        A.store.Bool(ptr + 4, x["isKeepAlive"] ? true : false);
        A.store.Ref(ptr + 8, x["lacrosPath"]);
        A.store.Enum(ptr + 12, ["Disabled", "Only"].indexOf(x["mode"] as string));
      }
    },
    "load_LacrosInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["state"] = A.load.Enum(ptr + 0, [
        "NotInitialized",
        "Reloading",
        "Mounting",
        "Unavailable",
        "Stopped",
        "CreatingLogFile",
        "PreLaunched",
        "Starting",
        "Running",
        "Terminating",
      ]);
      if (A.load.Bool(ptr + 16)) {
        x["isKeepAlive"] = A.load.Bool(ptr + 4);
      } else {
        delete x["isKeepAlive"];
      }
      x["lacrosPath"] = A.load.Ref(ptr + 8, undefined);
      x["mode"] = A.load.Enum(ptr + 12, ["Disabled", "Only"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_LauncherSearchBoxState": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["ghostText"]);
      }
    },
    "load_LauncherSearchBoxState": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["ghostText"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_LoginEventRecorderData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 4, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Ref(ptr + 0, x["name"]);
        A.store.Bool(ptr + 8, "microsecnods_since_unix_epoch" in x ? true : false);
        A.store.Int32(
          ptr + 4,
          x["microsecnods_since_unix_epoch"] === undefined ? 0 : (x["microsecnods_since_unix_epoch"] as number)
        );
      }
    },
    "load_LoginEventRecorderData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["name"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 8)) {
        x["microsecnods_since_unix_epoch"] = A.load.Int32(ptr + 4);
      } else {
        delete x["microsecnods_since_unix_epoch"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SystemWebApp": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["internalName"]);
        A.store.Ref(ptr + 4, x["url"]);
        A.store.Ref(ptr + 8, x["name"]);
        A.store.Ref(ptr + 12, x["startUrl"]);
      }
    },
    "load_SystemWebApp": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["internalName"] = A.load.Ref(ptr + 0, undefined);
      x["url"] = A.load.Ref(ptr + 4, undefined);
      x["name"] = A.load.Ref(ptr + 8, undefined);
      x["startUrl"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ScrollableShelfInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 116, false);
        A.store.Bool(ptr + 109, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 110, false);
        A.store.Float64(ptr + 8, 0);
        A.store.Bool(ptr + 111, false);
        A.store.Float64(ptr + 16, 0);

        A.store.Bool(ptr + 24 + 36, false);
        A.store.Bool(ptr + 24 + 32, false);
        A.store.Float64(ptr + 24 + 0, 0);
        A.store.Bool(ptr + 24 + 33, false);
        A.store.Float64(ptr + 24 + 8, 0);
        A.store.Bool(ptr + 24 + 34, false);
        A.store.Float64(ptr + 24 + 16, 0);
        A.store.Bool(ptr + 24 + 35, false);
        A.store.Float64(ptr + 24 + 24, 0);

        A.store.Bool(ptr + 64 + 36, false);
        A.store.Bool(ptr + 64 + 32, false);
        A.store.Float64(ptr + 64 + 0, 0);
        A.store.Bool(ptr + 64 + 33, false);
        A.store.Float64(ptr + 64 + 8, 0);
        A.store.Bool(ptr + 64 + 34, false);
        A.store.Float64(ptr + 64 + 16, 0);
        A.store.Bool(ptr + 64 + 35, false);
        A.store.Float64(ptr + 64 + 24, 0);
        A.store.Bool(ptr + 112, false);
        A.store.Bool(ptr + 101, false);
        A.store.Bool(ptr + 113, false);
        A.store.Bool(ptr + 102, false);
        A.store.Bool(ptr + 114, false);
        A.store.Bool(ptr + 103, false);
        A.store.Ref(ptr + 104, undefined);
        A.store.Bool(ptr + 115, false);
        A.store.Bool(ptr + 108, false);
      } else {
        A.store.Bool(ptr + 116, true);
        A.store.Bool(ptr + 109, "mainAxisOffset" in x ? true : false);
        A.store.Float64(ptr + 0, x["mainAxisOffset"] === undefined ? 0 : (x["mainAxisOffset"] as number));
        A.store.Bool(ptr + 110, "pageOffset" in x ? true : false);
        A.store.Float64(ptr + 8, x["pageOffset"] === undefined ? 0 : (x["pageOffset"] as number));
        A.store.Bool(ptr + 111, "targetMainAxisOffset" in x ? true : false);
        A.store.Float64(ptr + 16, x["targetMainAxisOffset"] === undefined ? 0 : (x["targetMainAxisOffset"] as number));

        if (typeof x["leftArrowBounds"] === "undefined") {
          A.store.Bool(ptr + 24 + 36, false);
          A.store.Bool(ptr + 24 + 32, false);
          A.store.Float64(ptr + 24 + 0, 0);
          A.store.Bool(ptr + 24 + 33, false);
          A.store.Float64(ptr + 24 + 8, 0);
          A.store.Bool(ptr + 24 + 34, false);
          A.store.Float64(ptr + 24 + 16, 0);
          A.store.Bool(ptr + 24 + 35, false);
          A.store.Float64(ptr + 24 + 24, 0);
        } else {
          A.store.Bool(ptr + 24 + 36, true);
          A.store.Bool(ptr + 24 + 32, "left" in x["leftArrowBounds"] ? true : false);
          A.store.Float64(
            ptr + 24 + 0,
            x["leftArrowBounds"]["left"] === undefined ? 0 : (x["leftArrowBounds"]["left"] as number)
          );
          A.store.Bool(ptr + 24 + 33, "top" in x["leftArrowBounds"] ? true : false);
          A.store.Float64(
            ptr + 24 + 8,
            x["leftArrowBounds"]["top"] === undefined ? 0 : (x["leftArrowBounds"]["top"] as number)
          );
          A.store.Bool(ptr + 24 + 34, "width" in x["leftArrowBounds"] ? true : false);
          A.store.Float64(
            ptr + 24 + 16,
            x["leftArrowBounds"]["width"] === undefined ? 0 : (x["leftArrowBounds"]["width"] as number)
          );
          A.store.Bool(ptr + 24 + 35, "height" in x["leftArrowBounds"] ? true : false);
          A.store.Float64(
            ptr + 24 + 24,
            x["leftArrowBounds"]["height"] === undefined ? 0 : (x["leftArrowBounds"]["height"] as number)
          );
        }

        if (typeof x["rightArrowBounds"] === "undefined") {
          A.store.Bool(ptr + 64 + 36, false);
          A.store.Bool(ptr + 64 + 32, false);
          A.store.Float64(ptr + 64 + 0, 0);
          A.store.Bool(ptr + 64 + 33, false);
          A.store.Float64(ptr + 64 + 8, 0);
          A.store.Bool(ptr + 64 + 34, false);
          A.store.Float64(ptr + 64 + 16, 0);
          A.store.Bool(ptr + 64 + 35, false);
          A.store.Float64(ptr + 64 + 24, 0);
        } else {
          A.store.Bool(ptr + 64 + 36, true);
          A.store.Bool(ptr + 64 + 32, "left" in x["rightArrowBounds"] ? true : false);
          A.store.Float64(
            ptr + 64 + 0,
            x["rightArrowBounds"]["left"] === undefined ? 0 : (x["rightArrowBounds"]["left"] as number)
          );
          A.store.Bool(ptr + 64 + 33, "top" in x["rightArrowBounds"] ? true : false);
          A.store.Float64(
            ptr + 64 + 8,
            x["rightArrowBounds"]["top"] === undefined ? 0 : (x["rightArrowBounds"]["top"] as number)
          );
          A.store.Bool(ptr + 64 + 34, "width" in x["rightArrowBounds"] ? true : false);
          A.store.Float64(
            ptr + 64 + 16,
            x["rightArrowBounds"]["width"] === undefined ? 0 : (x["rightArrowBounds"]["width"] as number)
          );
          A.store.Bool(ptr + 64 + 35, "height" in x["rightArrowBounds"] ? true : false);
          A.store.Float64(
            ptr + 64 + 24,
            x["rightArrowBounds"]["height"] === undefined ? 0 : (x["rightArrowBounds"]["height"] as number)
          );
        }
        A.store.Bool(ptr + 112, "isAnimating" in x ? true : false);
        A.store.Bool(ptr + 101, x["isAnimating"] ? true : false);
        A.store.Bool(ptr + 113, "iconsUnderAnimation" in x ? true : false);
        A.store.Bool(ptr + 102, x["iconsUnderAnimation"] ? true : false);
        A.store.Bool(ptr + 114, "isOverflow" in x ? true : false);
        A.store.Bool(ptr + 103, x["isOverflow"] ? true : false);
        A.store.Ref(ptr + 104, x["iconsBoundsInScreen"]);
        A.store.Bool(ptr + 115, "isShelfWidgetAnimating" in x ? true : false);
        A.store.Bool(ptr + 108, x["isShelfWidgetAnimating"] ? true : false);
      }
    },
    "load_ScrollableShelfInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 109)) {
        x["mainAxisOffset"] = A.load.Float64(ptr + 0);
      } else {
        delete x["mainAxisOffset"];
      }
      if (A.load.Bool(ptr + 110)) {
        x["pageOffset"] = A.load.Float64(ptr + 8);
      } else {
        delete x["pageOffset"];
      }
      if (A.load.Bool(ptr + 111)) {
        x["targetMainAxisOffset"] = A.load.Float64(ptr + 16);
      } else {
        delete x["targetMainAxisOffset"];
      }
      if (A.load.Bool(ptr + 24 + 36)) {
        x["leftArrowBounds"] = {};
        if (A.load.Bool(ptr + 24 + 32)) {
          x["leftArrowBounds"]["left"] = A.load.Float64(ptr + 24 + 0);
        } else {
          delete x["leftArrowBounds"]["left"];
        }
        if (A.load.Bool(ptr + 24 + 33)) {
          x["leftArrowBounds"]["top"] = A.load.Float64(ptr + 24 + 8);
        } else {
          delete x["leftArrowBounds"]["top"];
        }
        if (A.load.Bool(ptr + 24 + 34)) {
          x["leftArrowBounds"]["width"] = A.load.Float64(ptr + 24 + 16);
        } else {
          delete x["leftArrowBounds"]["width"];
        }
        if (A.load.Bool(ptr + 24 + 35)) {
          x["leftArrowBounds"]["height"] = A.load.Float64(ptr + 24 + 24);
        } else {
          delete x["leftArrowBounds"]["height"];
        }
      } else {
        delete x["leftArrowBounds"];
      }
      if (A.load.Bool(ptr + 64 + 36)) {
        x["rightArrowBounds"] = {};
        if (A.load.Bool(ptr + 64 + 32)) {
          x["rightArrowBounds"]["left"] = A.load.Float64(ptr + 64 + 0);
        } else {
          delete x["rightArrowBounds"]["left"];
        }
        if (A.load.Bool(ptr + 64 + 33)) {
          x["rightArrowBounds"]["top"] = A.load.Float64(ptr + 64 + 8);
        } else {
          delete x["rightArrowBounds"]["top"];
        }
        if (A.load.Bool(ptr + 64 + 34)) {
          x["rightArrowBounds"]["width"] = A.load.Float64(ptr + 64 + 16);
        } else {
          delete x["rightArrowBounds"]["width"];
        }
        if (A.load.Bool(ptr + 64 + 35)) {
          x["rightArrowBounds"]["height"] = A.load.Float64(ptr + 64 + 24);
        } else {
          delete x["rightArrowBounds"]["height"];
        }
      } else {
        delete x["rightArrowBounds"];
      }
      if (A.load.Bool(ptr + 112)) {
        x["isAnimating"] = A.load.Bool(ptr + 101);
      } else {
        delete x["isAnimating"];
      }
      if (A.load.Bool(ptr + 113)) {
        x["iconsUnderAnimation"] = A.load.Bool(ptr + 102);
      } else {
        delete x["iconsUnderAnimation"];
      }
      if (A.load.Bool(ptr + 114)) {
        x["isOverflow"] = A.load.Bool(ptr + 103);
      } else {
        delete x["isOverflow"];
      }
      x["iconsBoundsInScreen"] = A.load.Ref(ptr + 104, undefined);
      if (A.load.Bool(ptr + 115)) {
        x["isShelfWidgetAnimating"] = A.load.Bool(ptr + 108);
      } else {
        delete x["isShelfWidgetAnimating"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ShelfAlignmentType": (ref: heap.Ref<string>): number => {
      const idx = ["Bottom", "Left", "Right"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ShelfItemType": (ref: heap.Ref<string>): number => {
      const idx = ["PinnedApp", "BrowserShortcut", "App", "UnpinnedBrowserShortcut", "Dialog"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ShelfItemStatus": (ref: heap.Ref<string>): number => {
      const idx = ["Closed", "Running", "Attention"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ShelfItem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Enum(ptr + 12, -1);
        A.store.Enum(ptr + 16, -1);
        A.store.Bool(ptr + 24, false);
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 26, false);
        A.store.Bool(ptr + 22, false);
        A.store.Bool(ptr + 27, false);
        A.store.Bool(ptr + 23, false);
      } else {
        A.store.Bool(ptr + 28, true);
        A.store.Ref(ptr + 0, x["appId"]);
        A.store.Ref(ptr + 4, x["launchId"]);
        A.store.Ref(ptr + 8, x["title"]);
        A.store.Enum(
          ptr + 12,
          ["PinnedApp", "BrowserShortcut", "App", "UnpinnedBrowserShortcut", "Dialog"].indexOf(x["type"] as string)
        );
        A.store.Enum(ptr + 16, ["Closed", "Running", "Attention"].indexOf(x["status"] as string));
        A.store.Bool(ptr + 24, "showsTooltip" in x ? true : false);
        A.store.Bool(ptr + 20, x["showsTooltip"] ? true : false);
        A.store.Bool(ptr + 25, "pinnedByPolicy" in x ? true : false);
        A.store.Bool(ptr + 21, x["pinnedByPolicy"] ? true : false);
        A.store.Bool(ptr + 26, "pinStateForcedByType" in x ? true : false);
        A.store.Bool(ptr + 22, x["pinStateForcedByType"] ? true : false);
        A.store.Bool(ptr + 27, "hasNotification" in x ? true : false);
        A.store.Bool(ptr + 23, x["hasNotification"] ? true : false);
      }
    },
    "load_ShelfItem": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["appId"] = A.load.Ref(ptr + 0, undefined);
      x["launchId"] = A.load.Ref(ptr + 4, undefined);
      x["title"] = A.load.Ref(ptr + 8, undefined);
      x["type"] = A.load.Enum(ptr + 12, ["PinnedApp", "BrowserShortcut", "App", "UnpinnedBrowserShortcut", "Dialog"]);
      x["status"] = A.load.Enum(ptr + 16, ["Closed", "Running", "Attention"]);
      if (A.load.Bool(ptr + 24)) {
        x["showsTooltip"] = A.load.Bool(ptr + 20);
      } else {
        delete x["showsTooltip"];
      }
      if (A.load.Bool(ptr + 25)) {
        x["pinnedByPolicy"] = A.load.Bool(ptr + 21);
      } else {
        delete x["pinnedByPolicy"];
      }
      if (A.load.Bool(ptr + 26)) {
        x["pinStateForcedByType"] = A.load.Bool(ptr + 22);
      } else {
        delete x["pinStateForcedByType"];
      }
      if (A.load.Bool(ptr + 27)) {
        x["hasNotification"] = A.load.Bool(ptr + 23);
      } else {
        delete x["hasNotification"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Location": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 16, false);
        A.store.Float64(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Float64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Bool(ptr + 16, "x" in x ? true : false);
        A.store.Float64(ptr + 0, x["x"] === undefined ? 0 : (x["x"] as number));
        A.store.Bool(ptr + 17, "y" in x ? true : false);
        A.store.Float64(ptr + 8, x["y"] === undefined ? 0 : (x["y"] as number));
      }
    },
    "load_Location": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["x"] = A.load.Float64(ptr + 0);
      } else {
        delete x["x"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["y"] = A.load.Float64(ptr + 8);
      } else {
        delete x["y"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_HotseatSwipeDescriptor": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 43, false);

        A.store.Bool(ptr + 0 + 18, false);
        A.store.Bool(ptr + 0 + 16, false);
        A.store.Float64(ptr + 0 + 0, 0);
        A.store.Bool(ptr + 0 + 17, false);
        A.store.Float64(ptr + 0 + 8, 0);

        A.store.Bool(ptr + 24 + 18, false);
        A.store.Bool(ptr + 24 + 16, false);
        A.store.Float64(ptr + 24 + 0, 0);
        A.store.Bool(ptr + 24 + 17, false);
        A.store.Float64(ptr + 24 + 8, 0);
      } else {
        A.store.Bool(ptr + 43, true);

        if (typeof x["swipeStartLocation"] === "undefined") {
          A.store.Bool(ptr + 0 + 18, false);
          A.store.Bool(ptr + 0 + 16, false);
          A.store.Float64(ptr + 0 + 0, 0);
          A.store.Bool(ptr + 0 + 17, false);
          A.store.Float64(ptr + 0 + 8, 0);
        } else {
          A.store.Bool(ptr + 0 + 18, true);
          A.store.Bool(ptr + 0 + 16, "x" in x["swipeStartLocation"] ? true : false);
          A.store.Float64(
            ptr + 0 + 0,
            x["swipeStartLocation"]["x"] === undefined ? 0 : (x["swipeStartLocation"]["x"] as number)
          );
          A.store.Bool(ptr + 0 + 17, "y" in x["swipeStartLocation"] ? true : false);
          A.store.Float64(
            ptr + 0 + 8,
            x["swipeStartLocation"]["y"] === undefined ? 0 : (x["swipeStartLocation"]["y"] as number)
          );
        }

        if (typeof x["swipeEndLocation"] === "undefined") {
          A.store.Bool(ptr + 24 + 18, false);
          A.store.Bool(ptr + 24 + 16, false);
          A.store.Float64(ptr + 24 + 0, 0);
          A.store.Bool(ptr + 24 + 17, false);
          A.store.Float64(ptr + 24 + 8, 0);
        } else {
          A.store.Bool(ptr + 24 + 18, true);
          A.store.Bool(ptr + 24 + 16, "x" in x["swipeEndLocation"] ? true : false);
          A.store.Float64(
            ptr + 24 + 0,
            x["swipeEndLocation"]["x"] === undefined ? 0 : (x["swipeEndLocation"]["x"] as number)
          );
          A.store.Bool(ptr + 24 + 17, "y" in x["swipeEndLocation"] ? true : false);
          A.store.Float64(
            ptr + 24 + 8,
            x["swipeEndLocation"]["y"] === undefined ? 0 : (x["swipeEndLocation"]["y"] as number)
          );
        }
      }
    },
    "load_HotseatSwipeDescriptor": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 18)) {
        x["swipeStartLocation"] = {};
        if (A.load.Bool(ptr + 0 + 16)) {
          x["swipeStartLocation"]["x"] = A.load.Float64(ptr + 0 + 0);
        } else {
          delete x["swipeStartLocation"]["x"];
        }
        if (A.load.Bool(ptr + 0 + 17)) {
          x["swipeStartLocation"]["y"] = A.load.Float64(ptr + 0 + 8);
        } else {
          delete x["swipeStartLocation"]["y"];
        }
      } else {
        delete x["swipeStartLocation"];
      }
      if (A.load.Bool(ptr + 24 + 18)) {
        x["swipeEndLocation"] = {};
        if (A.load.Bool(ptr + 24 + 16)) {
          x["swipeEndLocation"]["x"] = A.load.Float64(ptr + 24 + 0);
        } else {
          delete x["swipeEndLocation"]["x"];
        }
        if (A.load.Bool(ptr + 24 + 17)) {
          x["swipeEndLocation"]["y"] = A.load.Float64(ptr + 24 + 8);
        } else {
          delete x["swipeEndLocation"]["y"];
        }
      } else {
        delete x["swipeEndLocation"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_HotseatState": (ref: heap.Ref<string>): number => {
      const idx = ["Hidden", "ShownClamShell", "ShownHomeLauncher", "Extended"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_HotseatInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 52, false);

        A.store.Bool(ptr + 0 + 43, false);

        A.store.Bool(ptr + 0 + 0 + 18, false);
        A.store.Bool(ptr + 0 + 0 + 16, false);
        A.store.Float64(ptr + 0 + 0 + 0, 0);
        A.store.Bool(ptr + 0 + 0 + 17, false);
        A.store.Float64(ptr + 0 + 0 + 8, 0);

        A.store.Bool(ptr + 0 + 24 + 18, false);
        A.store.Bool(ptr + 0 + 24 + 16, false);
        A.store.Float64(ptr + 0 + 24 + 0, 0);
        A.store.Bool(ptr + 0 + 24 + 17, false);
        A.store.Float64(ptr + 0 + 24 + 8, 0);
        A.store.Enum(ptr + 44, -1);
        A.store.Bool(ptr + 50, false);
        A.store.Bool(ptr + 48, false);
        A.store.Bool(ptr + 51, false);
        A.store.Bool(ptr + 49, false);
      } else {
        A.store.Bool(ptr + 52, true);

        if (typeof x["swipeUp"] === "undefined") {
          A.store.Bool(ptr + 0 + 43, false);

          A.store.Bool(ptr + 0 + 0 + 18, false);
          A.store.Bool(ptr + 0 + 0 + 16, false);
          A.store.Float64(ptr + 0 + 0 + 0, 0);
          A.store.Bool(ptr + 0 + 0 + 17, false);
          A.store.Float64(ptr + 0 + 0 + 8, 0);

          A.store.Bool(ptr + 0 + 24 + 18, false);
          A.store.Bool(ptr + 0 + 24 + 16, false);
          A.store.Float64(ptr + 0 + 24 + 0, 0);
          A.store.Bool(ptr + 0 + 24 + 17, false);
          A.store.Float64(ptr + 0 + 24 + 8, 0);
        } else {
          A.store.Bool(ptr + 0 + 43, true);

          if (typeof x["swipeUp"]["swipeStartLocation"] === "undefined") {
            A.store.Bool(ptr + 0 + 0 + 18, false);
            A.store.Bool(ptr + 0 + 0 + 16, false);
            A.store.Float64(ptr + 0 + 0 + 0, 0);
            A.store.Bool(ptr + 0 + 0 + 17, false);
            A.store.Float64(ptr + 0 + 0 + 8, 0);
          } else {
            A.store.Bool(ptr + 0 + 0 + 18, true);
            A.store.Bool(ptr + 0 + 0 + 16, "x" in x["swipeUp"]["swipeStartLocation"] ? true : false);
            A.store.Float64(
              ptr + 0 + 0 + 0,
              x["swipeUp"]["swipeStartLocation"]["x"] === undefined
                ? 0
                : (x["swipeUp"]["swipeStartLocation"]["x"] as number)
            );
            A.store.Bool(ptr + 0 + 0 + 17, "y" in x["swipeUp"]["swipeStartLocation"] ? true : false);
            A.store.Float64(
              ptr + 0 + 0 + 8,
              x["swipeUp"]["swipeStartLocation"]["y"] === undefined
                ? 0
                : (x["swipeUp"]["swipeStartLocation"]["y"] as number)
            );
          }

          if (typeof x["swipeUp"]["swipeEndLocation"] === "undefined") {
            A.store.Bool(ptr + 0 + 24 + 18, false);
            A.store.Bool(ptr + 0 + 24 + 16, false);
            A.store.Float64(ptr + 0 + 24 + 0, 0);
            A.store.Bool(ptr + 0 + 24 + 17, false);
            A.store.Float64(ptr + 0 + 24 + 8, 0);
          } else {
            A.store.Bool(ptr + 0 + 24 + 18, true);
            A.store.Bool(ptr + 0 + 24 + 16, "x" in x["swipeUp"]["swipeEndLocation"] ? true : false);
            A.store.Float64(
              ptr + 0 + 24 + 0,
              x["swipeUp"]["swipeEndLocation"]["x"] === undefined
                ? 0
                : (x["swipeUp"]["swipeEndLocation"]["x"] as number)
            );
            A.store.Bool(ptr + 0 + 24 + 17, "y" in x["swipeUp"]["swipeEndLocation"] ? true : false);
            A.store.Float64(
              ptr + 0 + 24 + 8,
              x["swipeUp"]["swipeEndLocation"]["y"] === undefined
                ? 0
                : (x["swipeUp"]["swipeEndLocation"]["y"] as number)
            );
          }
        }
        A.store.Enum(
          ptr + 44,
          ["Hidden", "ShownClamShell", "ShownHomeLauncher", "Extended"].indexOf(x["state"] as string)
        );
        A.store.Bool(ptr + 50, "isAnimating" in x ? true : false);
        A.store.Bool(ptr + 48, x["isAnimating"] ? true : false);
        A.store.Bool(ptr + 51, "isAutoHidden" in x ? true : false);
        A.store.Bool(ptr + 49, x["isAutoHidden"] ? true : false);
      }
    },
    "load_HotseatInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 43)) {
        x["swipeUp"] = {};
        if (A.load.Bool(ptr + 0 + 0 + 18)) {
          x["swipeUp"]["swipeStartLocation"] = {};
          if (A.load.Bool(ptr + 0 + 0 + 16)) {
            x["swipeUp"]["swipeStartLocation"]["x"] = A.load.Float64(ptr + 0 + 0 + 0);
          } else {
            delete x["swipeUp"]["swipeStartLocation"]["x"];
          }
          if (A.load.Bool(ptr + 0 + 0 + 17)) {
            x["swipeUp"]["swipeStartLocation"]["y"] = A.load.Float64(ptr + 0 + 0 + 8);
          } else {
            delete x["swipeUp"]["swipeStartLocation"]["y"];
          }
        } else {
          delete x["swipeUp"]["swipeStartLocation"];
        }
        if (A.load.Bool(ptr + 0 + 24 + 18)) {
          x["swipeUp"]["swipeEndLocation"] = {};
          if (A.load.Bool(ptr + 0 + 24 + 16)) {
            x["swipeUp"]["swipeEndLocation"]["x"] = A.load.Float64(ptr + 0 + 24 + 0);
          } else {
            delete x["swipeUp"]["swipeEndLocation"]["x"];
          }
          if (A.load.Bool(ptr + 0 + 24 + 17)) {
            x["swipeUp"]["swipeEndLocation"]["y"] = A.load.Float64(ptr + 0 + 24 + 8);
          } else {
            delete x["swipeUp"]["swipeEndLocation"]["y"];
          }
        } else {
          delete x["swipeUp"]["swipeEndLocation"];
        }
      } else {
        delete x["swipeUp"];
      }
      x["state"] = A.load.Enum(ptr + 44, ["Hidden", "ShownClamShell", "ShownHomeLauncher", "Extended"]);
      if (A.load.Bool(ptr + 50)) {
        x["isAnimating"] = A.load.Bool(ptr + 48);
      } else {
        delete x["isAnimating"];
      }
      if (A.load.Bool(ptr + 51)) {
        x["isAutoHidden"] = A.load.Bool(ptr + 49);
      } else {
        delete x["isAutoHidden"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ShelfUIInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 173, false);

        A.store.Bool(ptr + 0 + 52, false);

        A.store.Bool(ptr + 0 + 0 + 43, false);

        A.store.Bool(ptr + 0 + 0 + 0 + 18, false);
        A.store.Bool(ptr + 0 + 0 + 0 + 16, false);
        A.store.Float64(ptr + 0 + 0 + 0 + 0, 0);
        A.store.Bool(ptr + 0 + 0 + 0 + 17, false);
        A.store.Float64(ptr + 0 + 0 + 0 + 8, 0);

        A.store.Bool(ptr + 0 + 0 + 24 + 18, false);
        A.store.Bool(ptr + 0 + 0 + 24 + 16, false);
        A.store.Float64(ptr + 0 + 0 + 24 + 0, 0);
        A.store.Bool(ptr + 0 + 0 + 24 + 17, false);
        A.store.Float64(ptr + 0 + 0 + 24 + 8, 0);
        A.store.Enum(ptr + 0 + 44, -1);
        A.store.Bool(ptr + 0 + 50, false);
        A.store.Bool(ptr + 0 + 48, false);
        A.store.Bool(ptr + 0 + 51, false);
        A.store.Bool(ptr + 0 + 49, false);

        A.store.Bool(ptr + 56 + 116, false);
        A.store.Bool(ptr + 56 + 109, false);
        A.store.Float64(ptr + 56 + 0, 0);
        A.store.Bool(ptr + 56 + 110, false);
        A.store.Float64(ptr + 56 + 8, 0);
        A.store.Bool(ptr + 56 + 111, false);
        A.store.Float64(ptr + 56 + 16, 0);

        A.store.Bool(ptr + 56 + 24 + 36, false);
        A.store.Bool(ptr + 56 + 24 + 32, false);
        A.store.Float64(ptr + 56 + 24 + 0, 0);
        A.store.Bool(ptr + 56 + 24 + 33, false);
        A.store.Float64(ptr + 56 + 24 + 8, 0);
        A.store.Bool(ptr + 56 + 24 + 34, false);
        A.store.Float64(ptr + 56 + 24 + 16, 0);
        A.store.Bool(ptr + 56 + 24 + 35, false);
        A.store.Float64(ptr + 56 + 24 + 24, 0);

        A.store.Bool(ptr + 56 + 64 + 36, false);
        A.store.Bool(ptr + 56 + 64 + 32, false);
        A.store.Float64(ptr + 56 + 64 + 0, 0);
        A.store.Bool(ptr + 56 + 64 + 33, false);
        A.store.Float64(ptr + 56 + 64 + 8, 0);
        A.store.Bool(ptr + 56 + 64 + 34, false);
        A.store.Float64(ptr + 56 + 64 + 16, 0);
        A.store.Bool(ptr + 56 + 64 + 35, false);
        A.store.Float64(ptr + 56 + 64 + 24, 0);
        A.store.Bool(ptr + 56 + 112, false);
        A.store.Bool(ptr + 56 + 101, false);
        A.store.Bool(ptr + 56 + 113, false);
        A.store.Bool(ptr + 56 + 102, false);
        A.store.Bool(ptr + 56 + 114, false);
        A.store.Bool(ptr + 56 + 103, false);
        A.store.Ref(ptr + 56 + 104, undefined);
        A.store.Bool(ptr + 56 + 115, false);
        A.store.Bool(ptr + 56 + 108, false);
      } else {
        A.store.Bool(ptr + 173, true);

        if (typeof x["hotseatInfo"] === "undefined") {
          A.store.Bool(ptr + 0 + 52, false);

          A.store.Bool(ptr + 0 + 0 + 43, false);

          A.store.Bool(ptr + 0 + 0 + 0 + 18, false);
          A.store.Bool(ptr + 0 + 0 + 0 + 16, false);
          A.store.Float64(ptr + 0 + 0 + 0 + 0, 0);
          A.store.Bool(ptr + 0 + 0 + 0 + 17, false);
          A.store.Float64(ptr + 0 + 0 + 0 + 8, 0);

          A.store.Bool(ptr + 0 + 0 + 24 + 18, false);
          A.store.Bool(ptr + 0 + 0 + 24 + 16, false);
          A.store.Float64(ptr + 0 + 0 + 24 + 0, 0);
          A.store.Bool(ptr + 0 + 0 + 24 + 17, false);
          A.store.Float64(ptr + 0 + 0 + 24 + 8, 0);
          A.store.Enum(ptr + 0 + 44, -1);
          A.store.Bool(ptr + 0 + 50, false);
          A.store.Bool(ptr + 0 + 48, false);
          A.store.Bool(ptr + 0 + 51, false);
          A.store.Bool(ptr + 0 + 49, false);
        } else {
          A.store.Bool(ptr + 0 + 52, true);

          if (typeof x["hotseatInfo"]["swipeUp"] === "undefined") {
            A.store.Bool(ptr + 0 + 0 + 43, false);

            A.store.Bool(ptr + 0 + 0 + 0 + 18, false);
            A.store.Bool(ptr + 0 + 0 + 0 + 16, false);
            A.store.Float64(ptr + 0 + 0 + 0 + 0, 0);
            A.store.Bool(ptr + 0 + 0 + 0 + 17, false);
            A.store.Float64(ptr + 0 + 0 + 0 + 8, 0);

            A.store.Bool(ptr + 0 + 0 + 24 + 18, false);
            A.store.Bool(ptr + 0 + 0 + 24 + 16, false);
            A.store.Float64(ptr + 0 + 0 + 24 + 0, 0);
            A.store.Bool(ptr + 0 + 0 + 24 + 17, false);
            A.store.Float64(ptr + 0 + 0 + 24 + 8, 0);
          } else {
            A.store.Bool(ptr + 0 + 0 + 43, true);

            if (typeof x["hotseatInfo"]["swipeUp"]["swipeStartLocation"] === "undefined") {
              A.store.Bool(ptr + 0 + 0 + 0 + 18, false);
              A.store.Bool(ptr + 0 + 0 + 0 + 16, false);
              A.store.Float64(ptr + 0 + 0 + 0 + 0, 0);
              A.store.Bool(ptr + 0 + 0 + 0 + 17, false);
              A.store.Float64(ptr + 0 + 0 + 0 + 8, 0);
            } else {
              A.store.Bool(ptr + 0 + 0 + 0 + 18, true);
              A.store.Bool(
                ptr + 0 + 0 + 0 + 16,
                "x" in x["hotseatInfo"]["swipeUp"]["swipeStartLocation"] ? true : false
              );
              A.store.Float64(
                ptr + 0 + 0 + 0 + 0,
                x["hotseatInfo"]["swipeUp"]["swipeStartLocation"]["x"] === undefined
                  ? 0
                  : (x["hotseatInfo"]["swipeUp"]["swipeStartLocation"]["x"] as number)
              );
              A.store.Bool(
                ptr + 0 + 0 + 0 + 17,
                "y" in x["hotseatInfo"]["swipeUp"]["swipeStartLocation"] ? true : false
              );
              A.store.Float64(
                ptr + 0 + 0 + 0 + 8,
                x["hotseatInfo"]["swipeUp"]["swipeStartLocation"]["y"] === undefined
                  ? 0
                  : (x["hotseatInfo"]["swipeUp"]["swipeStartLocation"]["y"] as number)
              );
            }

            if (typeof x["hotseatInfo"]["swipeUp"]["swipeEndLocation"] === "undefined") {
              A.store.Bool(ptr + 0 + 0 + 24 + 18, false);
              A.store.Bool(ptr + 0 + 0 + 24 + 16, false);
              A.store.Float64(ptr + 0 + 0 + 24 + 0, 0);
              A.store.Bool(ptr + 0 + 0 + 24 + 17, false);
              A.store.Float64(ptr + 0 + 0 + 24 + 8, 0);
            } else {
              A.store.Bool(ptr + 0 + 0 + 24 + 18, true);
              A.store.Bool(
                ptr + 0 + 0 + 24 + 16,
                "x" in x["hotseatInfo"]["swipeUp"]["swipeEndLocation"] ? true : false
              );
              A.store.Float64(
                ptr + 0 + 0 + 24 + 0,
                x["hotseatInfo"]["swipeUp"]["swipeEndLocation"]["x"] === undefined
                  ? 0
                  : (x["hotseatInfo"]["swipeUp"]["swipeEndLocation"]["x"] as number)
              );
              A.store.Bool(
                ptr + 0 + 0 + 24 + 17,
                "y" in x["hotseatInfo"]["swipeUp"]["swipeEndLocation"] ? true : false
              );
              A.store.Float64(
                ptr + 0 + 0 + 24 + 8,
                x["hotseatInfo"]["swipeUp"]["swipeEndLocation"]["y"] === undefined
                  ? 0
                  : (x["hotseatInfo"]["swipeUp"]["swipeEndLocation"]["y"] as number)
              );
            }
          }
          A.store.Enum(
            ptr + 0 + 44,
            ["Hidden", "ShownClamShell", "ShownHomeLauncher", "Extended"].indexOf(x["hotseatInfo"]["state"] as string)
          );
          A.store.Bool(ptr + 0 + 50, "isAnimating" in x["hotseatInfo"] ? true : false);
          A.store.Bool(ptr + 0 + 48, x["hotseatInfo"]["isAnimating"] ? true : false);
          A.store.Bool(ptr + 0 + 51, "isAutoHidden" in x["hotseatInfo"] ? true : false);
          A.store.Bool(ptr + 0 + 49, x["hotseatInfo"]["isAutoHidden"] ? true : false);
        }

        if (typeof x["scrollableShelfInfo"] === "undefined") {
          A.store.Bool(ptr + 56 + 116, false);
          A.store.Bool(ptr + 56 + 109, false);
          A.store.Float64(ptr + 56 + 0, 0);
          A.store.Bool(ptr + 56 + 110, false);
          A.store.Float64(ptr + 56 + 8, 0);
          A.store.Bool(ptr + 56 + 111, false);
          A.store.Float64(ptr + 56 + 16, 0);

          A.store.Bool(ptr + 56 + 24 + 36, false);
          A.store.Bool(ptr + 56 + 24 + 32, false);
          A.store.Float64(ptr + 56 + 24 + 0, 0);
          A.store.Bool(ptr + 56 + 24 + 33, false);
          A.store.Float64(ptr + 56 + 24 + 8, 0);
          A.store.Bool(ptr + 56 + 24 + 34, false);
          A.store.Float64(ptr + 56 + 24 + 16, 0);
          A.store.Bool(ptr + 56 + 24 + 35, false);
          A.store.Float64(ptr + 56 + 24 + 24, 0);

          A.store.Bool(ptr + 56 + 64 + 36, false);
          A.store.Bool(ptr + 56 + 64 + 32, false);
          A.store.Float64(ptr + 56 + 64 + 0, 0);
          A.store.Bool(ptr + 56 + 64 + 33, false);
          A.store.Float64(ptr + 56 + 64 + 8, 0);
          A.store.Bool(ptr + 56 + 64 + 34, false);
          A.store.Float64(ptr + 56 + 64 + 16, 0);
          A.store.Bool(ptr + 56 + 64 + 35, false);
          A.store.Float64(ptr + 56 + 64 + 24, 0);
          A.store.Bool(ptr + 56 + 112, false);
          A.store.Bool(ptr + 56 + 101, false);
          A.store.Bool(ptr + 56 + 113, false);
          A.store.Bool(ptr + 56 + 102, false);
          A.store.Bool(ptr + 56 + 114, false);
          A.store.Bool(ptr + 56 + 103, false);
          A.store.Ref(ptr + 56 + 104, undefined);
          A.store.Bool(ptr + 56 + 115, false);
          A.store.Bool(ptr + 56 + 108, false);
        } else {
          A.store.Bool(ptr + 56 + 116, true);
          A.store.Bool(ptr + 56 + 109, "mainAxisOffset" in x["scrollableShelfInfo"] ? true : false);
          A.store.Float64(
            ptr + 56 + 0,
            x["scrollableShelfInfo"]["mainAxisOffset"] === undefined
              ? 0
              : (x["scrollableShelfInfo"]["mainAxisOffset"] as number)
          );
          A.store.Bool(ptr + 56 + 110, "pageOffset" in x["scrollableShelfInfo"] ? true : false);
          A.store.Float64(
            ptr + 56 + 8,
            x["scrollableShelfInfo"]["pageOffset"] === undefined
              ? 0
              : (x["scrollableShelfInfo"]["pageOffset"] as number)
          );
          A.store.Bool(ptr + 56 + 111, "targetMainAxisOffset" in x["scrollableShelfInfo"] ? true : false);
          A.store.Float64(
            ptr + 56 + 16,
            x["scrollableShelfInfo"]["targetMainAxisOffset"] === undefined
              ? 0
              : (x["scrollableShelfInfo"]["targetMainAxisOffset"] as number)
          );

          if (typeof x["scrollableShelfInfo"]["leftArrowBounds"] === "undefined") {
            A.store.Bool(ptr + 56 + 24 + 36, false);
            A.store.Bool(ptr + 56 + 24 + 32, false);
            A.store.Float64(ptr + 56 + 24 + 0, 0);
            A.store.Bool(ptr + 56 + 24 + 33, false);
            A.store.Float64(ptr + 56 + 24 + 8, 0);
            A.store.Bool(ptr + 56 + 24 + 34, false);
            A.store.Float64(ptr + 56 + 24 + 16, 0);
            A.store.Bool(ptr + 56 + 24 + 35, false);
            A.store.Float64(ptr + 56 + 24 + 24, 0);
          } else {
            A.store.Bool(ptr + 56 + 24 + 36, true);
            A.store.Bool(ptr + 56 + 24 + 32, "left" in x["scrollableShelfInfo"]["leftArrowBounds"] ? true : false);
            A.store.Float64(
              ptr + 56 + 24 + 0,
              x["scrollableShelfInfo"]["leftArrowBounds"]["left"] === undefined
                ? 0
                : (x["scrollableShelfInfo"]["leftArrowBounds"]["left"] as number)
            );
            A.store.Bool(ptr + 56 + 24 + 33, "top" in x["scrollableShelfInfo"]["leftArrowBounds"] ? true : false);
            A.store.Float64(
              ptr + 56 + 24 + 8,
              x["scrollableShelfInfo"]["leftArrowBounds"]["top"] === undefined
                ? 0
                : (x["scrollableShelfInfo"]["leftArrowBounds"]["top"] as number)
            );
            A.store.Bool(ptr + 56 + 24 + 34, "width" in x["scrollableShelfInfo"]["leftArrowBounds"] ? true : false);
            A.store.Float64(
              ptr + 56 + 24 + 16,
              x["scrollableShelfInfo"]["leftArrowBounds"]["width"] === undefined
                ? 0
                : (x["scrollableShelfInfo"]["leftArrowBounds"]["width"] as number)
            );
            A.store.Bool(ptr + 56 + 24 + 35, "height" in x["scrollableShelfInfo"]["leftArrowBounds"] ? true : false);
            A.store.Float64(
              ptr + 56 + 24 + 24,
              x["scrollableShelfInfo"]["leftArrowBounds"]["height"] === undefined
                ? 0
                : (x["scrollableShelfInfo"]["leftArrowBounds"]["height"] as number)
            );
          }

          if (typeof x["scrollableShelfInfo"]["rightArrowBounds"] === "undefined") {
            A.store.Bool(ptr + 56 + 64 + 36, false);
            A.store.Bool(ptr + 56 + 64 + 32, false);
            A.store.Float64(ptr + 56 + 64 + 0, 0);
            A.store.Bool(ptr + 56 + 64 + 33, false);
            A.store.Float64(ptr + 56 + 64 + 8, 0);
            A.store.Bool(ptr + 56 + 64 + 34, false);
            A.store.Float64(ptr + 56 + 64 + 16, 0);
            A.store.Bool(ptr + 56 + 64 + 35, false);
            A.store.Float64(ptr + 56 + 64 + 24, 0);
          } else {
            A.store.Bool(ptr + 56 + 64 + 36, true);
            A.store.Bool(ptr + 56 + 64 + 32, "left" in x["scrollableShelfInfo"]["rightArrowBounds"] ? true : false);
            A.store.Float64(
              ptr + 56 + 64 + 0,
              x["scrollableShelfInfo"]["rightArrowBounds"]["left"] === undefined
                ? 0
                : (x["scrollableShelfInfo"]["rightArrowBounds"]["left"] as number)
            );
            A.store.Bool(ptr + 56 + 64 + 33, "top" in x["scrollableShelfInfo"]["rightArrowBounds"] ? true : false);
            A.store.Float64(
              ptr + 56 + 64 + 8,
              x["scrollableShelfInfo"]["rightArrowBounds"]["top"] === undefined
                ? 0
                : (x["scrollableShelfInfo"]["rightArrowBounds"]["top"] as number)
            );
            A.store.Bool(ptr + 56 + 64 + 34, "width" in x["scrollableShelfInfo"]["rightArrowBounds"] ? true : false);
            A.store.Float64(
              ptr + 56 + 64 + 16,
              x["scrollableShelfInfo"]["rightArrowBounds"]["width"] === undefined
                ? 0
                : (x["scrollableShelfInfo"]["rightArrowBounds"]["width"] as number)
            );
            A.store.Bool(ptr + 56 + 64 + 35, "height" in x["scrollableShelfInfo"]["rightArrowBounds"] ? true : false);
            A.store.Float64(
              ptr + 56 + 64 + 24,
              x["scrollableShelfInfo"]["rightArrowBounds"]["height"] === undefined
                ? 0
                : (x["scrollableShelfInfo"]["rightArrowBounds"]["height"] as number)
            );
          }
          A.store.Bool(ptr + 56 + 112, "isAnimating" in x["scrollableShelfInfo"] ? true : false);
          A.store.Bool(ptr + 56 + 101, x["scrollableShelfInfo"]["isAnimating"] ? true : false);
          A.store.Bool(ptr + 56 + 113, "iconsUnderAnimation" in x["scrollableShelfInfo"] ? true : false);
          A.store.Bool(ptr + 56 + 102, x["scrollableShelfInfo"]["iconsUnderAnimation"] ? true : false);
          A.store.Bool(ptr + 56 + 114, "isOverflow" in x["scrollableShelfInfo"] ? true : false);
          A.store.Bool(ptr + 56 + 103, x["scrollableShelfInfo"]["isOverflow"] ? true : false);
          A.store.Ref(ptr + 56 + 104, x["scrollableShelfInfo"]["iconsBoundsInScreen"]);
          A.store.Bool(ptr + 56 + 115, "isShelfWidgetAnimating" in x["scrollableShelfInfo"] ? true : false);
          A.store.Bool(ptr + 56 + 108, x["scrollableShelfInfo"]["isShelfWidgetAnimating"] ? true : false);
        }
      }
    },
    "load_ShelfUIInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 52)) {
        x["hotseatInfo"] = {};
        if (A.load.Bool(ptr + 0 + 0 + 43)) {
          x["hotseatInfo"]["swipeUp"] = {};
          if (A.load.Bool(ptr + 0 + 0 + 0 + 18)) {
            x["hotseatInfo"]["swipeUp"]["swipeStartLocation"] = {};
            if (A.load.Bool(ptr + 0 + 0 + 0 + 16)) {
              x["hotseatInfo"]["swipeUp"]["swipeStartLocation"]["x"] = A.load.Float64(ptr + 0 + 0 + 0 + 0);
            } else {
              delete x["hotseatInfo"]["swipeUp"]["swipeStartLocation"]["x"];
            }
            if (A.load.Bool(ptr + 0 + 0 + 0 + 17)) {
              x["hotseatInfo"]["swipeUp"]["swipeStartLocation"]["y"] = A.load.Float64(ptr + 0 + 0 + 0 + 8);
            } else {
              delete x["hotseatInfo"]["swipeUp"]["swipeStartLocation"]["y"];
            }
          } else {
            delete x["hotseatInfo"]["swipeUp"]["swipeStartLocation"];
          }
          if (A.load.Bool(ptr + 0 + 0 + 24 + 18)) {
            x["hotseatInfo"]["swipeUp"]["swipeEndLocation"] = {};
            if (A.load.Bool(ptr + 0 + 0 + 24 + 16)) {
              x["hotseatInfo"]["swipeUp"]["swipeEndLocation"]["x"] = A.load.Float64(ptr + 0 + 0 + 24 + 0);
            } else {
              delete x["hotseatInfo"]["swipeUp"]["swipeEndLocation"]["x"];
            }
            if (A.load.Bool(ptr + 0 + 0 + 24 + 17)) {
              x["hotseatInfo"]["swipeUp"]["swipeEndLocation"]["y"] = A.load.Float64(ptr + 0 + 0 + 24 + 8);
            } else {
              delete x["hotseatInfo"]["swipeUp"]["swipeEndLocation"]["y"];
            }
          } else {
            delete x["hotseatInfo"]["swipeUp"]["swipeEndLocation"];
          }
        } else {
          delete x["hotseatInfo"]["swipeUp"];
        }
        x["hotseatInfo"]["state"] = A.load.Enum(ptr + 0 + 44, [
          "Hidden",
          "ShownClamShell",
          "ShownHomeLauncher",
          "Extended",
        ]);
        if (A.load.Bool(ptr + 0 + 50)) {
          x["hotseatInfo"]["isAnimating"] = A.load.Bool(ptr + 0 + 48);
        } else {
          delete x["hotseatInfo"]["isAnimating"];
        }
        if (A.load.Bool(ptr + 0 + 51)) {
          x["hotseatInfo"]["isAutoHidden"] = A.load.Bool(ptr + 0 + 49);
        } else {
          delete x["hotseatInfo"]["isAutoHidden"];
        }
      } else {
        delete x["hotseatInfo"];
      }
      if (A.load.Bool(ptr + 56 + 116)) {
        x["scrollableShelfInfo"] = {};
        if (A.load.Bool(ptr + 56 + 109)) {
          x["scrollableShelfInfo"]["mainAxisOffset"] = A.load.Float64(ptr + 56 + 0);
        } else {
          delete x["scrollableShelfInfo"]["mainAxisOffset"];
        }
        if (A.load.Bool(ptr + 56 + 110)) {
          x["scrollableShelfInfo"]["pageOffset"] = A.load.Float64(ptr + 56 + 8);
        } else {
          delete x["scrollableShelfInfo"]["pageOffset"];
        }
        if (A.load.Bool(ptr + 56 + 111)) {
          x["scrollableShelfInfo"]["targetMainAxisOffset"] = A.load.Float64(ptr + 56 + 16);
        } else {
          delete x["scrollableShelfInfo"]["targetMainAxisOffset"];
        }
        if (A.load.Bool(ptr + 56 + 24 + 36)) {
          x["scrollableShelfInfo"]["leftArrowBounds"] = {};
          if (A.load.Bool(ptr + 56 + 24 + 32)) {
            x["scrollableShelfInfo"]["leftArrowBounds"]["left"] = A.load.Float64(ptr + 56 + 24 + 0);
          } else {
            delete x["scrollableShelfInfo"]["leftArrowBounds"]["left"];
          }
          if (A.load.Bool(ptr + 56 + 24 + 33)) {
            x["scrollableShelfInfo"]["leftArrowBounds"]["top"] = A.load.Float64(ptr + 56 + 24 + 8);
          } else {
            delete x["scrollableShelfInfo"]["leftArrowBounds"]["top"];
          }
          if (A.load.Bool(ptr + 56 + 24 + 34)) {
            x["scrollableShelfInfo"]["leftArrowBounds"]["width"] = A.load.Float64(ptr + 56 + 24 + 16);
          } else {
            delete x["scrollableShelfInfo"]["leftArrowBounds"]["width"];
          }
          if (A.load.Bool(ptr + 56 + 24 + 35)) {
            x["scrollableShelfInfo"]["leftArrowBounds"]["height"] = A.load.Float64(ptr + 56 + 24 + 24);
          } else {
            delete x["scrollableShelfInfo"]["leftArrowBounds"]["height"];
          }
        } else {
          delete x["scrollableShelfInfo"]["leftArrowBounds"];
        }
        if (A.load.Bool(ptr + 56 + 64 + 36)) {
          x["scrollableShelfInfo"]["rightArrowBounds"] = {};
          if (A.load.Bool(ptr + 56 + 64 + 32)) {
            x["scrollableShelfInfo"]["rightArrowBounds"]["left"] = A.load.Float64(ptr + 56 + 64 + 0);
          } else {
            delete x["scrollableShelfInfo"]["rightArrowBounds"]["left"];
          }
          if (A.load.Bool(ptr + 56 + 64 + 33)) {
            x["scrollableShelfInfo"]["rightArrowBounds"]["top"] = A.load.Float64(ptr + 56 + 64 + 8);
          } else {
            delete x["scrollableShelfInfo"]["rightArrowBounds"]["top"];
          }
          if (A.load.Bool(ptr + 56 + 64 + 34)) {
            x["scrollableShelfInfo"]["rightArrowBounds"]["width"] = A.load.Float64(ptr + 56 + 64 + 16);
          } else {
            delete x["scrollableShelfInfo"]["rightArrowBounds"]["width"];
          }
          if (A.load.Bool(ptr + 56 + 64 + 35)) {
            x["scrollableShelfInfo"]["rightArrowBounds"]["height"] = A.load.Float64(ptr + 56 + 64 + 24);
          } else {
            delete x["scrollableShelfInfo"]["rightArrowBounds"]["height"];
          }
        } else {
          delete x["scrollableShelfInfo"]["rightArrowBounds"];
        }
        if (A.load.Bool(ptr + 56 + 112)) {
          x["scrollableShelfInfo"]["isAnimating"] = A.load.Bool(ptr + 56 + 101);
        } else {
          delete x["scrollableShelfInfo"]["isAnimating"];
        }
        if (A.load.Bool(ptr + 56 + 113)) {
          x["scrollableShelfInfo"]["iconsUnderAnimation"] = A.load.Bool(ptr + 56 + 102);
        } else {
          delete x["scrollableShelfInfo"]["iconsUnderAnimation"];
        }
        if (A.load.Bool(ptr + 56 + 114)) {
          x["scrollableShelfInfo"]["isOverflow"] = A.load.Bool(ptr + 56 + 103);
        } else {
          delete x["scrollableShelfInfo"]["isOverflow"];
        }
        x["scrollableShelfInfo"]["iconsBoundsInScreen"] = A.load.Ref(ptr + 56 + 104, undefined);
        if (A.load.Bool(ptr + 56 + 115)) {
          x["scrollableShelfInfo"]["isShelfWidgetAnimating"] = A.load.Bool(ptr + 56 + 108);
        } else {
          delete x["scrollableShelfInfo"]["isShelfWidgetAnimating"];
        }
      } else {
        delete x["scrollableShelfInfo"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ThroughputTrackerAnimationData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 25, false);
        A.store.Bool(ptr + 20, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 21, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 22, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 23, false);
        A.store.Int32(ptr + 12, 0);
        A.store.Bool(ptr + 24, false);
        A.store.Int32(ptr + 16, 0);
      } else {
        A.store.Bool(ptr + 25, true);
        A.store.Bool(ptr + 20, "startOffsetMs" in x ? true : false);
        A.store.Int32(ptr + 0, x["startOffsetMs"] === undefined ? 0 : (x["startOffsetMs"] as number));
        A.store.Bool(ptr + 21, "stopOffsetMs" in x ? true : false);
        A.store.Int32(ptr + 4, x["stopOffsetMs"] === undefined ? 0 : (x["stopOffsetMs"] as number));
        A.store.Bool(ptr + 22, "framesExpected" in x ? true : false);
        A.store.Int32(ptr + 8, x["framesExpected"] === undefined ? 0 : (x["framesExpected"] as number));
        A.store.Bool(ptr + 23, "framesProduced" in x ? true : false);
        A.store.Int32(ptr + 12, x["framesProduced"] === undefined ? 0 : (x["framesProduced"] as number));
        A.store.Bool(ptr + 24, "jankCount" in x ? true : false);
        A.store.Int32(ptr + 16, x["jankCount"] === undefined ? 0 : (x["jankCount"] as number));
      }
    },
    "load_ThroughputTrackerAnimationData": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 20)) {
        x["startOffsetMs"] = A.load.Int32(ptr + 0);
      } else {
        delete x["startOffsetMs"];
      }
      if (A.load.Bool(ptr + 21)) {
        x["stopOffsetMs"] = A.load.Int32(ptr + 4);
      } else {
        delete x["stopOffsetMs"];
      }
      if (A.load.Bool(ptr + 22)) {
        x["framesExpected"] = A.load.Int32(ptr + 8);
      } else {
        delete x["framesExpected"];
      }
      if (A.load.Bool(ptr + 23)) {
        x["framesProduced"] = A.load.Int32(ptr + 12);
      } else {
        delete x["framesProduced"];
      }
      if (A.load.Bool(ptr + 24)) {
        x["jankCount"] = A.load.Int32(ptr + 16);
      } else {
        delete x["jankCount"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_LauncherStateType": (ref: heap.Ref<string>): number => {
      const idx = ["Closed", "FullscreenAllApps", "FullscreenSearch"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_LoginStatusDict": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 39, false);
        A.store.Bool(ptr + 29, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 30, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 31, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 32, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 33, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 34, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 35, false);
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 36, false);
        A.store.Bool(ptr + 7, false);
        A.store.Bool(ptr + 37, false);
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
        A.store.Bool(ptr + 38, false);
        A.store.Bool(ptr + 28, false);
      } else {
        A.store.Bool(ptr + 39, true);
        A.store.Bool(ptr + 29, "isLoggedIn" in x ? true : false);
        A.store.Bool(ptr + 0, x["isLoggedIn"] ? true : false);
        A.store.Bool(ptr + 30, "isOwner" in x ? true : false);
        A.store.Bool(ptr + 1, x["isOwner"] ? true : false);
        A.store.Bool(ptr + 31, "isScreenLocked" in x ? true : false);
        A.store.Bool(ptr + 2, x["isScreenLocked"] ? true : false);
        A.store.Bool(ptr + 32, "isLockscreenWallpaperAnimating" in x ? true : false);
        A.store.Bool(ptr + 3, x["isLockscreenWallpaperAnimating"] ? true : false);
        A.store.Bool(ptr + 33, "isReadyForPassword" in x ? true : false);
        A.store.Bool(ptr + 4, x["isReadyForPassword"] ? true : false);
        A.store.Bool(ptr + 34, "areAllUserImagesLoaded" in x ? true : false);
        A.store.Bool(ptr + 5, x["areAllUserImagesLoaded"] ? true : false);
        A.store.Bool(ptr + 35, "isRegularUser" in x ? true : false);
        A.store.Bool(ptr + 6, x["isRegularUser"] ? true : false);
        A.store.Bool(ptr + 36, "isGuest" in x ? true : false);
        A.store.Bool(ptr + 7, x["isGuest"] ? true : false);
        A.store.Bool(ptr + 37, "isKiosk" in x ? true : false);
        A.store.Bool(ptr + 8, x["isKiosk"] ? true : false);
        A.store.Ref(ptr + 12, x["email"]);
        A.store.Ref(ptr + 16, x["displayEmail"]);
        A.store.Ref(ptr + 20, x["displayName"]);
        A.store.Ref(ptr + 24, x["userImage"]);
        A.store.Bool(ptr + 38, "hasValidOauth2Token" in x ? true : false);
        A.store.Bool(ptr + 28, x["hasValidOauth2Token"] ? true : false);
      }
    },
    "load_LoginStatusDict": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 29)) {
        x["isLoggedIn"] = A.load.Bool(ptr + 0);
      } else {
        delete x["isLoggedIn"];
      }
      if (A.load.Bool(ptr + 30)) {
        x["isOwner"] = A.load.Bool(ptr + 1);
      } else {
        delete x["isOwner"];
      }
      if (A.load.Bool(ptr + 31)) {
        x["isScreenLocked"] = A.load.Bool(ptr + 2);
      } else {
        delete x["isScreenLocked"];
      }
      if (A.load.Bool(ptr + 32)) {
        x["isLockscreenWallpaperAnimating"] = A.load.Bool(ptr + 3);
      } else {
        delete x["isLockscreenWallpaperAnimating"];
      }
      if (A.load.Bool(ptr + 33)) {
        x["isReadyForPassword"] = A.load.Bool(ptr + 4);
      } else {
        delete x["isReadyForPassword"];
      }
      if (A.load.Bool(ptr + 34)) {
        x["areAllUserImagesLoaded"] = A.load.Bool(ptr + 5);
      } else {
        delete x["areAllUserImagesLoaded"];
      }
      if (A.load.Bool(ptr + 35)) {
        x["isRegularUser"] = A.load.Bool(ptr + 6);
      } else {
        delete x["isRegularUser"];
      }
      if (A.load.Bool(ptr + 36)) {
        x["isGuest"] = A.load.Bool(ptr + 7);
      } else {
        delete x["isGuest"];
      }
      if (A.load.Bool(ptr + 37)) {
        x["isKiosk"] = A.load.Bool(ptr + 8);
      } else {
        delete x["isKiosk"];
      }
      x["email"] = A.load.Ref(ptr + 12, undefined);
      x["displayEmail"] = A.load.Ref(ptr + 16, undefined);
      x["displayName"] = A.load.Ref(ptr + 20, undefined);
      x["userImage"] = A.load.Ref(ptr + 24, undefined);
      if (A.load.Bool(ptr + 38)) {
        x["hasValidOauth2Token"] = A.load.Bool(ptr + 28);
      } else {
        delete x["hasValidOauth2Token"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MakeFuseboxTempDirData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["fuseboxFilePath"]);
        A.store.Ref(ptr + 4, x["underlyingFilePath"]);
      }
    },
    "load_MakeFuseboxTempDirData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["fuseboxFilePath"] = A.load.Ref(ptr + 0, undefined);
      x["underlyingFilePath"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_MouseButton": (ref: heap.Ref<string>): number => {
      const idx = ["Left", "Middle", "Right", "Back", "Forward"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Notification": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 26, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Bool(ptr + 24, false);
        A.store.Int32(ptr + 16, 0);
        A.store.Bool(ptr + 25, false);
        A.store.Int32(ptr + 20, 0);
      } else {
        A.store.Bool(ptr + 26, true);
        A.store.Ref(ptr + 0, x["id"]);
        A.store.Ref(ptr + 4, x["type"]);
        A.store.Ref(ptr + 8, x["title"]);
        A.store.Ref(ptr + 12, x["message"]);
        A.store.Bool(ptr + 24, "priority" in x ? true : false);
        A.store.Int32(ptr + 16, x["priority"] === undefined ? 0 : (x["priority"] as number));
        A.store.Bool(ptr + 25, "progress" in x ? true : false);
        A.store.Int32(ptr + 20, x["progress"] === undefined ? 0 : (x["progress"] as number));
      }
    },
    "load_Notification": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["id"] = A.load.Ref(ptr + 0, undefined);
      x["type"] = A.load.Ref(ptr + 4, undefined);
      x["title"] = A.load.Ref(ptr + 8, undefined);
      x["message"] = A.load.Ref(ptr + 12, undefined);
      if (A.load.Bool(ptr + 24)) {
        x["priority"] = A.load.Int32(ptr + 16);
      } else {
        delete x["priority"];
      }
      if (A.load.Bool(ptr + 25)) {
        x["progress"] = A.load.Int32(ptr + 20);
      } else {
        delete x["progress"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OverviewStateType": (ref: heap.Ref<string>): number => {
      const idx = ["Shown", "Hidden"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_PlayStoreState": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 6, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 2, false);
      } else {
        A.store.Bool(ptr + 6, true);
        A.store.Bool(ptr + 3, "allowed" in x ? true : false);
        A.store.Bool(ptr + 0, x["allowed"] ? true : false);
        A.store.Bool(ptr + 4, "enabled" in x ? true : false);
        A.store.Bool(ptr + 1, x["enabled"] ? true : false);
        A.store.Bool(ptr + 5, "managed" in x ? true : false);
        A.store.Bool(ptr + 2, x["managed"] ? true : false);
      }
    },
    "load_PlayStoreState": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 3)) {
        x["allowed"] = A.load.Bool(ptr + 0);
      } else {
        delete x["allowed"];
      }
      if (A.load.Bool(ptr + 4)) {
        x["enabled"] = A.load.Bool(ptr + 1);
      } else {
        delete x["enabled"];
      }
      if (A.load.Bool(ptr + 5)) {
        x["managed"] = A.load.Bool(ptr + 2);
      } else {
        delete x["managed"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_Printer": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Ref(ptr + 20, undefined);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 28, true);
        A.store.Ref(ptr + 0, x["printerName"]);
        A.store.Ref(ptr + 4, x["printerId"]);
        A.store.Ref(ptr + 8, x["printerType"]);
        A.store.Ref(ptr + 12, x["printerDesc"]);
        A.store.Ref(ptr + 16, x["printerMakeAndModel"]);
        A.store.Ref(ptr + 20, x["printerUri"]);
        A.store.Ref(ptr + 24, x["printerPpd"]);
      }
    },
    "load_Printer": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["printerName"] = A.load.Ref(ptr + 0, undefined);
      x["printerId"] = A.load.Ref(ptr + 4, undefined);
      x["printerType"] = A.load.Ref(ptr + 8, undefined);
      x["printerDesc"] = A.load.Ref(ptr + 12, undefined);
      x["printerMakeAndModel"] = A.load.Ref(ptr + 16, undefined);
      x["printerUri"] = A.load.Ref(ptr + 20, undefined);
      x["printerPpd"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ResetHoldingSpaceOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 1, "markTimeOfFirstAdd" in x ? true : false);
        A.store.Bool(ptr + 0, x["markTimeOfFirstAdd"] ? true : false);
      }
    },
    "load_ResetHoldingSpaceOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 1)) {
        x["markTimeOfFirstAdd"] = A.load.Bool(ptr + 0);
      } else {
        delete x["markTimeOfFirstAdd"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_RotationType": (ref: heap.Ref<string>): number => {
      const idx = ["RotateAny", "Rotate0", "Rotate90", "Rotate180", "Rotate270"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ScrollableShelfState": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Float64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "scrollDistance" in x ? true : false);
        A.store.Float64(ptr + 0, x["scrollDistance"] === undefined ? 0 : (x["scrollDistance"] as number));
      }
    },
    "load_ScrollableShelfState": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["scrollDistance"] = A.load.Float64(ptr + 0);
      } else {
        delete x["scrollDistance"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetWindowBoundsResult": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 44, false);

        A.store.Bool(ptr + 0 + 36, false);
        A.store.Bool(ptr + 0 + 32, false);
        A.store.Float64(ptr + 0 + 0, 0);
        A.store.Bool(ptr + 0 + 33, false);
        A.store.Float64(ptr + 0 + 8, 0);
        A.store.Bool(ptr + 0 + 34, false);
        A.store.Float64(ptr + 0 + 16, 0);
        A.store.Bool(ptr + 0 + 35, false);
        A.store.Float64(ptr + 0 + 24, 0);
        A.store.Ref(ptr + 40, undefined);
      } else {
        A.store.Bool(ptr + 44, true);

        if (typeof x["bounds"] === "undefined") {
          A.store.Bool(ptr + 0 + 36, false);
          A.store.Bool(ptr + 0 + 32, false);
          A.store.Float64(ptr + 0 + 0, 0);
          A.store.Bool(ptr + 0 + 33, false);
          A.store.Float64(ptr + 0 + 8, 0);
          A.store.Bool(ptr + 0 + 34, false);
          A.store.Float64(ptr + 0 + 16, 0);
          A.store.Bool(ptr + 0 + 35, false);
          A.store.Float64(ptr + 0 + 24, 0);
        } else {
          A.store.Bool(ptr + 0 + 36, true);
          A.store.Bool(ptr + 0 + 32, "left" in x["bounds"] ? true : false);
          A.store.Float64(ptr + 0 + 0, x["bounds"]["left"] === undefined ? 0 : (x["bounds"]["left"] as number));
          A.store.Bool(ptr + 0 + 33, "top" in x["bounds"] ? true : false);
          A.store.Float64(ptr + 0 + 8, x["bounds"]["top"] === undefined ? 0 : (x["bounds"]["top"] as number));
          A.store.Bool(ptr + 0 + 34, "width" in x["bounds"] ? true : false);
          A.store.Float64(ptr + 0 + 16, x["bounds"]["width"] === undefined ? 0 : (x["bounds"]["width"] as number));
          A.store.Bool(ptr + 0 + 35, "height" in x["bounds"] ? true : false);
          A.store.Float64(ptr + 0 + 24, x["bounds"]["height"] === undefined ? 0 : (x["bounds"]["height"] as number));
        }
        A.store.Ref(ptr + 40, x["displayId"]);
      }
    },
    "load_SetWindowBoundsResult": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 0 + 36)) {
        x["bounds"] = {};
        if (A.load.Bool(ptr + 0 + 32)) {
          x["bounds"]["left"] = A.load.Float64(ptr + 0 + 0);
        } else {
          delete x["bounds"]["left"];
        }
        if (A.load.Bool(ptr + 0 + 33)) {
          x["bounds"]["top"] = A.load.Float64(ptr + 0 + 8);
        } else {
          delete x["bounds"]["top"];
        }
        if (A.load.Bool(ptr + 0 + 34)) {
          x["bounds"]["width"] = A.load.Float64(ptr + 0 + 16);
        } else {
          delete x["bounds"]["width"];
        }
        if (A.load.Bool(ptr + 0 + 35)) {
          x["bounds"]["height"] = A.load.Float64(ptr + 0 + 24);
        } else {
          delete x["bounds"]["height"];
        }
      } else {
        delete x["bounds"];
      }
      x["displayId"] = A.load.Ref(ptr + 40, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ShelfIconPinUpdateParam": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 6, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 6, true);
        A.store.Ref(ptr + 0, x["appId"]);
        A.store.Bool(ptr + 5, "pinned" in x ? true : false);
        A.store.Bool(ptr + 4, x["pinned"] ? true : false);
      }
    },
    "load_ShelfIconPinUpdateParam": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["appId"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 5)) {
        x["pinned"] = A.load.Bool(ptr + 4);
      } else {
        delete x["pinned"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_ShelfState": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Float64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "scrollDistance" in x ? true : false);
        A.store.Float64(ptr + 0, x["scrollDistance"] === undefined ? 0 : (x["scrollDistance"] as number));
      }
    },
    "load_ShelfState": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["scrollDistance"] = A.load.Float64(ptr + 0);
      } else {
        delete x["scrollDistance"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ThemeStyle": (ref: heap.Ref<string>): number => {
      const idx = ["TonalSpot", "Vibrant", "Expressive", "Spritz", "Rainbow", "FruitSalad"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_WMEventType": (ref: heap.Ref<string>): number => {
      const idx = [
        "WMEventNormal",
        "WMEventMaximize",
        "WMEventMinimize",
        "WMEventFullscreen",
        "WMEventSnapPrimary",
        "WMEventSnapSecondary",
        "WMEventFloat",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_WindowStateChangeDict": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 6, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 5, false);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 6, true);
        A.store.Enum(
          ptr + 0,
          [
            "WMEventNormal",
            "WMEventMaximize",
            "WMEventMinimize",
            "WMEventFullscreen",
            "WMEventSnapPrimary",
            "WMEventSnapSecondary",
            "WMEventFloat",
          ].indexOf(x["eventType"] as string)
        );
        A.store.Bool(ptr + 5, "failIfNoChange" in x ? true : false);
        A.store.Bool(ptr + 4, x["failIfNoChange"] ? true : false);
      }
    },
    "load_WindowStateChangeDict": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["eventType"] = A.load.Enum(ptr + 0, [
        "WMEventNormal",
        "WMEventMaximize",
        "WMEventMinimize",
        "WMEventFullscreen",
        "WMEventSnapPrimary",
        "WMEventSnapSecondary",
        "WMEventFloat",
      ]);
      if (A.load.Bool(ptr + 5)) {
        x["failIfNoChange"] = A.load.Bool(ptr + 4);
      } else {
        delete x["failIfNoChange"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_ActivateAccelerator": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "activateAccelerator" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ActivateAccelerator": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.activateAccelerator);
    },
    "call_ActivateAccelerator": (retPtr: Pointer, accelerator: Pointer): void => {
      const accelerator_ffi = {};

      accelerator_ffi["keyCode"] = A.load.Ref(accelerator + 0, undefined);
      if (A.load.Bool(accelerator + 9)) {
        accelerator_ffi["shift"] = A.load.Bool(accelerator + 4);
      }
      if (A.load.Bool(accelerator + 10)) {
        accelerator_ffi["control"] = A.load.Bool(accelerator + 5);
      }
      if (A.load.Bool(accelerator + 11)) {
        accelerator_ffi["alt"] = A.load.Bool(accelerator + 6);
      }
      if (A.load.Bool(accelerator + 12)) {
        accelerator_ffi["search"] = A.load.Bool(accelerator + 7);
      }
      if (A.load.Bool(accelerator + 13)) {
        accelerator_ffi["pressed"] = A.load.Bool(accelerator + 8);
      }

      const _ret = WEBEXT.autotestPrivate.activateAccelerator(accelerator_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ActivateAccelerator": (retPtr: Pointer, errPtr: Pointer, accelerator: Pointer): heap.Ref<boolean> => {
      try {
        const accelerator_ffi = {};

        accelerator_ffi["keyCode"] = A.load.Ref(accelerator + 0, undefined);
        if (A.load.Bool(accelerator + 9)) {
          accelerator_ffi["shift"] = A.load.Bool(accelerator + 4);
        }
        if (A.load.Bool(accelerator + 10)) {
          accelerator_ffi["control"] = A.load.Bool(accelerator + 5);
        }
        if (A.load.Bool(accelerator + 11)) {
          accelerator_ffi["alt"] = A.load.Bool(accelerator + 6);
        }
        if (A.load.Bool(accelerator + 12)) {
          accelerator_ffi["search"] = A.load.Bool(accelerator + 7);
        }
        if (A.load.Bool(accelerator + 13)) {
          accelerator_ffi["pressed"] = A.load.Bool(accelerator + 8);
        }

        const _ret = WEBEXT.autotestPrivate.activateAccelerator(accelerator_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ActivateAdjacentDesksToTargetIndex": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "activateAdjacentDesksToTargetIndex" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ActivateAdjacentDesksToTargetIndex": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.activateAdjacentDesksToTargetIndex);
    },
    "call_ActivateAdjacentDesksToTargetIndex": (retPtr: Pointer, index: number): void => {
      const _ret = WEBEXT.autotestPrivate.activateAdjacentDesksToTargetIndex(index);
      A.store.Ref(retPtr, _ret);
    },
    "try_ActivateAdjacentDesksToTargetIndex": (retPtr: Pointer, errPtr: Pointer, index: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.activateAdjacentDesksToTargetIndex(index);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ActivateAppWindow": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "activateAppWindow" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ActivateAppWindow": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.activateAppWindow);
    },
    "call_ActivateAppWindow": (retPtr: Pointer, id: number): void => {
      const _ret = WEBEXT.autotestPrivate.activateAppWindow(id);
      A.store.Ref(retPtr, _ret);
    },
    "try_ActivateAppWindow": (retPtr: Pointer, errPtr: Pointer, id: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.activateAppWindow(id);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ActivateDeskAtIndex": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "activateDeskAtIndex" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ActivateDeskAtIndex": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.activateDeskAtIndex);
    },
    "call_ActivateDeskAtIndex": (retPtr: Pointer, index: number): void => {
      const _ret = WEBEXT.autotestPrivate.activateDeskAtIndex(index);
      A.store.Ref(retPtr, _ret);
    },
    "try_ActivateDeskAtIndex": (retPtr: Pointer, errPtr: Pointer, index: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.activateDeskAtIndex(index);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AddLoginEventForTesting": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "addLoginEventForTesting" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddLoginEventForTesting": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.addLoginEventForTesting);
    },
    "call_AddLoginEventForTesting": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.addLoginEventForTesting();
      A.store.Ref(retPtr, _ret);
    },
    "try_AddLoginEventForTesting": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.addLoginEventForTesting();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ArcAppTracingStart": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "arcAppTracingStart" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ArcAppTracingStart": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.arcAppTracingStart);
    },
    "call_ArcAppTracingStart": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.arcAppTracingStart();
      A.store.Ref(retPtr, _ret);
    },
    "try_ArcAppTracingStart": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.arcAppTracingStart();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ArcAppTracingStopAndAnalyze": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "arcAppTracingStopAndAnalyze" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ArcAppTracingStopAndAnalyze": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.arcAppTracingStopAndAnalyze);
    },
    "call_ArcAppTracingStopAndAnalyze": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.arcAppTracingStopAndAnalyze();
      A.store.Ref(retPtr, _ret);
    },
    "try_ArcAppTracingStopAndAnalyze": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.arcAppTracingStopAndAnalyze();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_BootstrapMachineLearningService": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "bootstrapMachineLearningService" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_BootstrapMachineLearningService": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.bootstrapMachineLearningService);
    },
    "call_BootstrapMachineLearningService": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.bootstrapMachineLearningService();
      A.store.Ref(retPtr, _ret);
    },
    "try_BootstrapMachineLearningService": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.bootstrapMachineLearningService();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CloseApp": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "closeApp" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CloseApp": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.closeApp);
    },
    "call_CloseApp": (retPtr: Pointer, appId: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.closeApp(A.H.get<object>(appId));
      A.store.Ref(retPtr, _ret);
    },
    "try_CloseApp": (retPtr: Pointer, errPtr: Pointer, appId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.closeApp(A.H.get<object>(appId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CloseAppWindow": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "closeAppWindow" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CloseAppWindow": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.closeAppWindow);
    },
    "call_CloseAppWindow": (retPtr: Pointer, id: number): void => {
      const _ret = WEBEXT.autotestPrivate.closeAppWindow(id);
      A.store.Ref(retPtr, _ret);
    },
    "try_CloseAppWindow": (retPtr: Pointer, errPtr: Pointer, id: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.closeAppWindow(id);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CouldAllowCrostini": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "couldAllowCrostini" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CouldAllowCrostini": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.couldAllowCrostini);
    },
    "call_CouldAllowCrostini": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.couldAllowCrostini();
      A.store.Ref(retPtr, _ret);
    },
    "try_CouldAllowCrostini": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.couldAllowCrostini();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CreateNewDesk": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "createNewDesk" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CreateNewDesk": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.createNewDesk);
    },
    "call_CreateNewDesk": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.createNewDesk();
      A.store.Ref(retPtr, _ret);
    },
    "try_CreateNewDesk": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.createNewDesk();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DisableAutomation": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "disableAutomation" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DisableAutomation": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.disableAutomation);
    },
    "call_DisableAutomation": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.disableAutomation();
      A.store.Ref(retPtr, _ret);
    },
    "try_DisableAutomation": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.disableAutomation();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DisableSwitchAccessDialog": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "disableSwitchAccessDialog" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DisableSwitchAccessDialog": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.disableSwitchAccessDialog);
    },
    "call_DisableSwitchAccessDialog": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.disableSwitchAccessDialog();
    },
    "try_DisableSwitchAccessDialog": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.disableSwitchAccessDialog();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_EnableAssistantAndWaitForReady": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "enableAssistantAndWaitForReady" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_EnableAssistantAndWaitForReady": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.enableAssistantAndWaitForReady);
    },
    "call_EnableAssistantAndWaitForReady": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.enableAssistantAndWaitForReady();
      A.store.Ref(retPtr, _ret);
    },
    "try_EnableAssistantAndWaitForReady": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.enableAssistantAndWaitForReady();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ExportCrostini": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "exportCrostini" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ExportCrostini": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.exportCrostini);
    },
    "call_ExportCrostini": (retPtr: Pointer, path: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.exportCrostini(A.H.get<object>(path));
      A.store.Ref(retPtr, _ret);
    },
    "try_ExportCrostini": (retPtr: Pointer, errPtr: Pointer, path: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.exportCrostini(A.H.get<object>(path));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ForceAutoThemeMode": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "forceAutoThemeMode" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ForceAutoThemeMode": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.forceAutoThemeMode);
    },
    "call_ForceAutoThemeMode": (retPtr: Pointer, darkModeEnabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.forceAutoThemeMode(darkModeEnabled === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_ForceAutoThemeMode": (
      retPtr: Pointer,
      errPtr: Pointer,
      darkModeEnabled: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.forceAutoThemeMode(darkModeEnabled === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAccessToken": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getAccessToken" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAccessToken": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getAccessToken);
    },
    "call_GetAccessToken": (retPtr: Pointer, accessTokenParams: Pointer): void => {
      const accessTokenParams_ffi = {};

      accessTokenParams_ffi["email"] = A.load.Ref(accessTokenParams + 0, undefined);
      accessTokenParams_ffi["scopes"] = A.load.Ref(accessTokenParams + 4, undefined);
      if (A.load.Bool(accessTokenParams + 12)) {
        accessTokenParams_ffi["timeoutMs"] = A.load.Int32(accessTokenParams + 8);
      }

      const _ret = WEBEXT.autotestPrivate.getAccessToken(accessTokenParams_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAccessToken": (retPtr: Pointer, errPtr: Pointer, accessTokenParams: Pointer): heap.Ref<boolean> => {
      try {
        const accessTokenParams_ffi = {};

        accessTokenParams_ffi["email"] = A.load.Ref(accessTokenParams + 0, undefined);
        accessTokenParams_ffi["scopes"] = A.load.Ref(accessTokenParams + 4, undefined);
        if (A.load.Bool(accessTokenParams + 12)) {
          accessTokenParams_ffi["timeoutMs"] = A.load.Int32(accessTokenParams + 8);
        }

        const _ret = WEBEXT.autotestPrivate.getAccessToken(accessTokenParams_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAllEnterprisePolicies": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getAllEnterprisePolicies" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAllEnterprisePolicies": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getAllEnterprisePolicies);
    },
    "call_GetAllEnterprisePolicies": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getAllEnterprisePolicies();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAllEnterprisePolicies": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getAllEnterprisePolicies();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAllInstalledApps": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getAllInstalledApps" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAllInstalledApps": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getAllInstalledApps);
    },
    "call_GetAllInstalledApps": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getAllInstalledApps();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAllInstalledApps": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getAllInstalledApps();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAppWindowList": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getAppWindowList" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAppWindowList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getAppWindowList);
    },
    "call_GetAppWindowList": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getAppWindowList();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAppWindowList": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getAppWindowList();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetArcApp": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getArcApp" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetArcApp": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getArcApp);
    },
    "call_GetArcApp": (retPtr: Pointer, appId: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.getArcApp(A.H.get<object>(appId));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetArcApp": (retPtr: Pointer, errPtr: Pointer, appId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getArcApp(A.H.get<object>(appId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetArcAppKills": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getArcAppKills" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetArcAppKills": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getArcAppKills);
    },
    "call_GetArcAppKills": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getArcAppKills();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetArcAppKills": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getArcAppKills();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetArcPackage": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getArcPackage" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetArcPackage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getArcPackage);
    },
    "call_GetArcPackage": (retPtr: Pointer, packageName: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.getArcPackage(A.H.get<object>(packageName));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetArcPackage": (retPtr: Pointer, errPtr: Pointer, packageName: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getArcPackage(A.H.get<object>(packageName));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetArcStartTime": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getArcStartTime" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetArcStartTime": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getArcStartTime);
    },
    "call_GetArcStartTime": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getArcStartTime();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetArcStartTime": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getArcStartTime();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetArcState": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getArcState" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetArcState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getArcState);
    },
    "call_GetArcState": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getArcState();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetArcState": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getArcState();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetClipboardTextData": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getClipboardTextData" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetClipboardTextData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getClipboardTextData);
    },
    "call_GetClipboardTextData": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getClipboardTextData();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetClipboardTextData": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getClipboardTextData();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCryptohomeRecoveryData": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getCryptohomeRecoveryData" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCryptohomeRecoveryData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getCryptohomeRecoveryData);
    },
    "call_GetCryptohomeRecoveryData": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getCryptohomeRecoveryData();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCryptohomeRecoveryData": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getCryptohomeRecoveryData();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCurrentInputMethodDescriptor": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getCurrentInputMethodDescriptor" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCurrentInputMethodDescriptor": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getCurrentInputMethodDescriptor);
    },
    "call_GetCurrentInputMethodDescriptor": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getCurrentInputMethodDescriptor();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCurrentInputMethodDescriptor": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getCurrentInputMethodDescriptor();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDefaultPinnedAppIds": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getDefaultPinnedAppIds" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDefaultPinnedAppIds": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getDefaultPinnedAppIds);
    },
    "call_GetDefaultPinnedAppIds": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getDefaultPinnedAppIds();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDefaultPinnedAppIds": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getDefaultPinnedAppIds();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDeskCount": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getDeskCount" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDeskCount": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getDeskCount);
    },
    "call_GetDeskCount": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getDeskCount();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDeskCount": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getDeskCount();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDesksInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getDesksInfo" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDesksInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getDesksInfo);
    },
    "call_GetDesksInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getDesksInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDesksInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getDesksInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDisplaySmoothness": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getDisplaySmoothness" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDisplaySmoothness": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getDisplaySmoothness);
    },
    "call_GetDisplaySmoothness": (retPtr: Pointer, displayId: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.getDisplaySmoothness(A.H.get<object>(displayId));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDisplaySmoothness": (retPtr: Pointer, errPtr: Pointer, displayId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getDisplaySmoothness(A.H.get<object>(displayId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetExtensionsInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getExtensionsInfo" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetExtensionsInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getExtensionsInfo);
    },
    "call_GetExtensionsInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getExtensionsInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetExtensionsInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getExtensionsInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetLacrosInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getLacrosInfo" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetLacrosInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getLacrosInfo);
    },
    "call_GetLacrosInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getLacrosInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetLacrosInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getLacrosInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetLauncherSearchBoxState": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getLauncherSearchBoxState" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetLauncherSearchBoxState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getLauncherSearchBoxState);
    },
    "call_GetLauncherSearchBoxState": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getLauncherSearchBoxState();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetLauncherSearchBoxState": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getLauncherSearchBoxState();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetLoginEventRecorderLoginEvents": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getLoginEventRecorderLoginEvents" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetLoginEventRecorderLoginEvents": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getLoginEventRecorderLoginEvents);
    },
    "call_GetLoginEventRecorderLoginEvents": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getLoginEventRecorderLoginEvents();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetLoginEventRecorderLoginEvents": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getLoginEventRecorderLoginEvents();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPlayStoreState": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getPlayStoreState" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPlayStoreState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getPlayStoreState);
    },
    "call_GetPlayStoreState": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getPlayStoreState();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPlayStoreState": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getPlayStoreState();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPrimaryDisplayScaleFactor": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getPrimaryDisplayScaleFactor" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPrimaryDisplayScaleFactor": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getPrimaryDisplayScaleFactor);
    },
    "call_GetPrimaryDisplayScaleFactor": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getPrimaryDisplayScaleFactor();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPrimaryDisplayScaleFactor": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getPrimaryDisplayScaleFactor();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPrinterList": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getPrinterList" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPrinterList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getPrinterList);
    },
    "call_GetPrinterList": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getPrinterList();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPrinterList": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getPrinterList();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetRegisteredSystemWebApps": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getRegisteredSystemWebApps" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetRegisteredSystemWebApps": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getRegisteredSystemWebApps);
    },
    "call_GetRegisteredSystemWebApps": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getRegisteredSystemWebApps();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetRegisteredSystemWebApps": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getRegisteredSystemWebApps();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetScrollableShelfInfoForState": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getScrollableShelfInfoForState" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetScrollableShelfInfoForState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getScrollableShelfInfoForState);
    },
    "call_GetScrollableShelfInfoForState": (retPtr: Pointer, state: Pointer): void => {
      const state_ffi = {};

      if (A.load.Bool(state + 8)) {
        state_ffi["scrollDistance"] = A.load.Float64(state + 0);
      }

      const _ret = WEBEXT.autotestPrivate.getScrollableShelfInfoForState(state_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetScrollableShelfInfoForState": (retPtr: Pointer, errPtr: Pointer, state: Pointer): heap.Ref<boolean> => {
      try {
        const state_ffi = {};

        if (A.load.Bool(state + 8)) {
          state_ffi["scrollDistance"] = A.load.Float64(state + 0);
        }

        const _ret = WEBEXT.autotestPrivate.getScrollableShelfInfoForState(state_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetShelfAlignment": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getShelfAlignment" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetShelfAlignment": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getShelfAlignment);
    },
    "call_GetShelfAlignment": (retPtr: Pointer, displayId: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.getShelfAlignment(A.H.get<object>(displayId));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetShelfAlignment": (retPtr: Pointer, errPtr: Pointer, displayId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getShelfAlignment(A.H.get<object>(displayId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetShelfAutoHideBehavior": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getShelfAutoHideBehavior" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetShelfAutoHideBehavior": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getShelfAutoHideBehavior);
    },
    "call_GetShelfAutoHideBehavior": (retPtr: Pointer, displayId: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.getShelfAutoHideBehavior(A.H.get<object>(displayId));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetShelfAutoHideBehavior": (
      retPtr: Pointer,
      errPtr: Pointer,
      displayId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getShelfAutoHideBehavior(A.H.get<object>(displayId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetShelfItems": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getShelfItems" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetShelfItems": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getShelfItems);
    },
    "call_GetShelfItems": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getShelfItems();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetShelfItems": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getShelfItems();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetShelfUIInfoForState": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getShelfUIInfoForState" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetShelfUIInfoForState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getShelfUIInfoForState);
    },
    "call_GetShelfUIInfoForState": (retPtr: Pointer, state: Pointer): void => {
      const state_ffi = {};

      if (A.load.Bool(state + 8)) {
        state_ffi["scrollDistance"] = A.load.Float64(state + 0);
      }

      const _ret = WEBEXT.autotestPrivate.getShelfUIInfoForState(state_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetShelfUIInfoForState": (retPtr: Pointer, errPtr: Pointer, state: Pointer): heap.Ref<boolean> => {
      try {
        const state_ffi = {};

        if (A.load.Bool(state + 8)) {
          state_ffi["scrollDistance"] = A.load.Float64(state + 0);
        }

        const _ret = WEBEXT.autotestPrivate.getShelfUIInfoForState(state_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetThroughputTrackerData": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getThroughputTrackerData" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetThroughputTrackerData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getThroughputTrackerData);
    },
    "call_GetThroughputTrackerData": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getThroughputTrackerData();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetThroughputTrackerData": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getThroughputTrackerData();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetVisibleNotifications": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "getVisibleNotifications" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetVisibleNotifications": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.getVisibleNotifications);
    },
    "call_GetVisibleNotifications": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.getVisibleNotifications();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetVisibleNotifications": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.getVisibleNotifications();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ImportCrostini": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "importCrostini" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ImportCrostini": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.importCrostini);
    },
    "call_ImportCrostini": (retPtr: Pointer, path: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.importCrostini(A.H.get<object>(path));
      A.store.Ref(retPtr, _ret);
    },
    "try_ImportCrostini": (retPtr: Pointer, errPtr: Pointer, path: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.importCrostini(A.H.get<object>(path));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InitializeEvents": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "initializeEvents" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InitializeEvents": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.initializeEvents);
    },
    "call_InitializeEvents": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.initializeEvents();
    },
    "try_InitializeEvents": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.initializeEvents();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InstallBorealis": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "installBorealis" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InstallBorealis": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.installBorealis);
    },
    "call_InstallBorealis": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.installBorealis();
      A.store.Ref(retPtr, _ret);
    },
    "try_InstallBorealis": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.installBorealis();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InstallBruschetta": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "installBruschetta" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InstallBruschetta": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.installBruschetta);
    },
    "call_InstallBruschetta": (retPtr: Pointer, vm_name: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.installBruschetta(A.H.get<object>(vm_name));
      A.store.Ref(retPtr, _ret);
    },
    "try_InstallBruschetta": (retPtr: Pointer, errPtr: Pointer, vm_name: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.installBruschetta(A.H.get<object>(vm_name));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_InstallPWAForCurrentURL": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "installPWAForCurrentURL" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_InstallPWAForCurrentURL": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.installPWAForCurrentURL);
    },
    "call_InstallPWAForCurrentURL": (retPtr: Pointer, timeout_ms: number): void => {
      const _ret = WEBEXT.autotestPrivate.installPWAForCurrentURL(timeout_ms);
      A.store.Ref(retPtr, _ret);
    },
    "try_InstallPWAForCurrentURL": (retPtr: Pointer, errPtr: Pointer, timeout_ms: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.installPWAForCurrentURL(timeout_ms);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsAppShown": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "isAppShown" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsAppShown": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.isAppShown);
    },
    "call_IsAppShown": (retPtr: Pointer, appId: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.isAppShown(A.H.get<object>(appId));
      A.store.Ref(retPtr, _ret);
    },
    "try_IsAppShown": (retPtr: Pointer, errPtr: Pointer, appId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.isAppShown(A.H.get<object>(appId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsArcPackageListInitialRefreshed": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "isArcPackageListInitialRefreshed" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsArcPackageListInitialRefreshed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.isArcPackageListInitialRefreshed);
    },
    "call_IsArcPackageListInitialRefreshed": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.isArcPackageListInitialRefreshed();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsArcPackageListInitialRefreshed": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.isArcPackageListInitialRefreshed();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsArcProvisioned": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "isArcProvisioned" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsArcProvisioned": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.isArcProvisioned);
    },
    "call_IsArcProvisioned": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.isArcProvisioned();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsArcProvisioned": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.isArcProvisioned();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsFeatureEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "isFeatureEnabled" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsFeatureEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.isFeatureEnabled);
    },
    "call_IsFeatureEnabled": (retPtr: Pointer, feature_name: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.isFeatureEnabled(A.H.get<object>(feature_name));
      A.store.Ref(retPtr, _ret);
    },
    "try_IsFeatureEnabled": (retPtr: Pointer, errPtr: Pointer, feature_name: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.isFeatureEnabled(A.H.get<object>(feature_name));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsInputMethodReadyForTesting": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "isInputMethodReadyForTesting" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsInputMethodReadyForTesting": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.isInputMethodReadyForTesting);
    },
    "call_IsInputMethodReadyForTesting": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.isInputMethodReadyForTesting();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsInputMethodReadyForTesting": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.isInputMethodReadyForTesting();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsSystemWebAppOpen": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "isSystemWebAppOpen" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsSystemWebAppOpen": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.isSystemWebAppOpen);
    },
    "call_IsSystemWebAppOpen": (retPtr: Pointer, appId: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.isSystemWebAppOpen(A.H.get<object>(appId));
      A.store.Ref(retPtr, _ret);
    },
    "try_IsSystemWebAppOpen": (retPtr: Pointer, errPtr: Pointer, appId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.isSystemWebAppOpen(A.H.get<object>(appId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsTabletModeEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "isTabletModeEnabled" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsTabletModeEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.isTabletModeEnabled);
    },
    "call_IsTabletModeEnabled": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.isTabletModeEnabled();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsTabletModeEnabled": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.isTabletModeEnabled();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LaunchApp": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "launchApp" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LaunchApp": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.launchApp);
    },
    "call_LaunchApp": (retPtr: Pointer, appId: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.launchApp(A.H.get<object>(appId));
      A.store.Ref(retPtr, _ret);
    },
    "try_LaunchApp": (retPtr: Pointer, errPtr: Pointer, appId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.launchApp(A.H.get<object>(appId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LaunchFilesAppToPath": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "launchFilesAppToPath" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LaunchFilesAppToPath": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.launchFilesAppToPath);
    },
    "call_LaunchFilesAppToPath": (retPtr: Pointer, absolutePath: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.launchFilesAppToPath(A.H.get<object>(absolutePath));
      A.store.Ref(retPtr, _ret);
    },
    "try_LaunchFilesAppToPath": (
      retPtr: Pointer,
      errPtr: Pointer,
      absolutePath: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.launchFilesAppToPath(A.H.get<object>(absolutePath));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LaunchSystemWebApp": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "launchSystemWebApp" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LaunchSystemWebApp": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.launchSystemWebApp);
    },
    "call_LaunchSystemWebApp": (retPtr: Pointer, appName: heap.Ref<object>, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.launchSystemWebApp(A.H.get<object>(appName), A.H.get<object>(url));
      A.store.Ref(retPtr, _ret);
    },
    "try_LaunchSystemWebApp": (
      retPtr: Pointer,
      errPtr: Pointer,
      appName: heap.Ref<object>,
      url: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.launchSystemWebApp(A.H.get<object>(appName), A.H.get<object>(url));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LoadSmartDimComponent": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "loadSmartDimComponent" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LoadSmartDimComponent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.loadSmartDimComponent);
    },
    "call_LoadSmartDimComponent": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.loadSmartDimComponent();
      A.store.Ref(retPtr, _ret);
    },
    "try_LoadSmartDimComponent": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.loadSmartDimComponent();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LockScreen": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "lockScreen" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LockScreen": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.lockScreen);
    },
    "call_LockScreen": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.lockScreen();
    },
    "try_LockScreen": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.lockScreen();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LoginStatus": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "loginStatus" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LoginStatus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.loginStatus);
    },
    "call_LoginStatus": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.loginStatus();
      A.store.Ref(retPtr, _ret);
    },
    "try_LoginStatus": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.loginStatus();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Logout": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "logout" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Logout": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.logout);
    },
    "call_Logout": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.logout();
    },
    "try_Logout": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.logout();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_MakeFuseboxTempDir": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "makeFuseboxTempDir" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MakeFuseboxTempDir": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.makeFuseboxTempDir);
    },
    "call_MakeFuseboxTempDir": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.makeFuseboxTempDir();
      A.store.Ref(retPtr, _ret);
    },
    "try_MakeFuseboxTempDir": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.makeFuseboxTempDir();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_MouseClick": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "mouseClick" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MouseClick": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.mouseClick);
    },
    "call_MouseClick": (retPtr: Pointer, button: number): void => {
      const _ret = WEBEXT.autotestPrivate.mouseClick(
        button > 0 && button <= 5 ? ["Left", "Middle", "Right", "Back", "Forward"][button - 1] : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_MouseClick": (retPtr: Pointer, errPtr: Pointer, button: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.mouseClick(
          button > 0 && button <= 5 ? ["Left", "Middle", "Right", "Back", "Forward"][button - 1] : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_MouseMove": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "mouseMove" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MouseMove": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.mouseMove);
    },
    "call_MouseMove": (retPtr: Pointer, location: Pointer, duration_in_ms: number): void => {
      const location_ffi = {};

      if (A.load.Bool(location + 16)) {
        location_ffi["x"] = A.load.Float64(location + 0);
      }
      if (A.load.Bool(location + 17)) {
        location_ffi["y"] = A.load.Float64(location + 8);
      }

      const _ret = WEBEXT.autotestPrivate.mouseMove(location_ffi, duration_in_ms);
      A.store.Ref(retPtr, _ret);
    },
    "try_MouseMove": (
      retPtr: Pointer,
      errPtr: Pointer,
      location: Pointer,
      duration_in_ms: number
    ): heap.Ref<boolean> => {
      try {
        const location_ffi = {};

        if (A.load.Bool(location + 16)) {
          location_ffi["x"] = A.load.Float64(location + 0);
        }
        if (A.load.Bool(location + 17)) {
          location_ffi["y"] = A.load.Float64(location + 8);
        }

        const _ret = WEBEXT.autotestPrivate.mouseMove(location_ffi, duration_in_ms);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_MousePress": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "mousePress" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MousePress": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.mousePress);
    },
    "call_MousePress": (retPtr: Pointer, button: number): void => {
      const _ret = WEBEXT.autotestPrivate.mousePress(
        button > 0 && button <= 5 ? ["Left", "Middle", "Right", "Back", "Forward"][button - 1] : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_MousePress": (retPtr: Pointer, errPtr: Pointer, button: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.mousePress(
          button > 0 && button <= 5 ? ["Left", "Middle", "Right", "Back", "Forward"][button - 1] : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_MouseRelease": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "mouseRelease" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_MouseRelease": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.mouseRelease);
    },
    "call_MouseRelease": (retPtr: Pointer, button: number): void => {
      const _ret = WEBEXT.autotestPrivate.mouseRelease(
        button > 0 && button <= 5 ? ["Left", "Middle", "Right", "Back", "Forward"][button - 1] : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_MouseRelease": (retPtr: Pointer, errPtr: Pointer, button: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.mouseRelease(
          button > 0 && button <= 5 ? ["Left", "Middle", "Right", "Back", "Forward"][button - 1] : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnClipboardDataChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.autotestPrivate?.onClipboardDataChanged &&
        "addListener" in WEBEXT?.autotestPrivate?.onClipboardDataChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnClipboardDataChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.onClipboardDataChanged.addListener);
    },
    "call_OnClipboardDataChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.onClipboardDataChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnClipboardDataChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.onClipboardDataChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffClipboardDataChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.autotestPrivate?.onClipboardDataChanged &&
        "removeListener" in WEBEXT?.autotestPrivate?.onClipboardDataChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffClipboardDataChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.onClipboardDataChanged.removeListener);
    },
    "call_OffClipboardDataChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.onClipboardDataChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffClipboardDataChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.onClipboardDataChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnClipboardDataChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.autotestPrivate?.onClipboardDataChanged &&
        "hasListener" in WEBEXT?.autotestPrivate?.onClipboardDataChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnClipboardDataChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.onClipboardDataChanged.hasListener);
    },
    "call_HasOnClipboardDataChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.onClipboardDataChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnClipboardDataChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.onClipboardDataChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_PinShelfIcon": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "pinShelfIcon" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_PinShelfIcon": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.pinShelfIcon);
    },
    "call_PinShelfIcon": (retPtr: Pointer, appId: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.pinShelfIcon(A.H.get<object>(appId));
      A.store.Ref(retPtr, _ret);
    },
    "try_PinShelfIcon": (retPtr: Pointer, errPtr: Pointer, appId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.pinShelfIcon(A.H.get<object>(appId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RefreshEnterprisePolicies": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "refreshEnterprisePolicies" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RefreshEnterprisePolicies": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.refreshEnterprisePolicies);
    },
    "call_RefreshEnterprisePolicies": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.refreshEnterprisePolicies();
      A.store.Ref(retPtr, _ret);
    },
    "try_RefreshEnterprisePolicies": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.refreshEnterprisePolicies();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RefreshRemoteCommands": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "refreshRemoteCommands" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RefreshRemoteCommands": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.refreshRemoteCommands);
    },
    "call_RefreshRemoteCommands": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.refreshRemoteCommands();
      A.store.Ref(retPtr, _ret);
    },
    "try_RefreshRemoteCommands": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.refreshRemoteCommands();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RegisterComponent": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "registerComponent" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RegisterComponent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.registerComponent);
    },
    "call_RegisterComponent": (retPtr: Pointer, name: heap.Ref<object>, path: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.registerComponent(A.H.get<object>(name), A.H.get<object>(path));
    },
    "try_RegisterComponent": (
      retPtr: Pointer,
      errPtr: Pointer,
      name: heap.Ref<object>,
      path: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.registerComponent(A.H.get<object>(name), A.H.get<object>(path));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveActiveDesk": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "removeActiveDesk" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveActiveDesk": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.removeActiveDesk);
    },
    "call_RemoveActiveDesk": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.removeActiveDesk();
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveActiveDesk": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.removeActiveDesk();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveAllNotifications": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "removeAllNotifications" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveAllNotifications": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.removeAllNotifications);
    },
    "call_RemoveAllNotifications": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.removeAllNotifications();
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveAllNotifications": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.removeAllNotifications();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveBruschetta": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "removeBruschetta" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveBruschetta": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.removeBruschetta);
    },
    "call_RemoveBruschetta": (retPtr: Pointer, vm_name: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.removeBruschetta(A.H.get<object>(vm_name));
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveBruschetta": (retPtr: Pointer, errPtr: Pointer, vm_name: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.removeBruschetta(A.H.get<object>(vm_name));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveComponentExtension": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "removeComponentExtension" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveComponentExtension": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.removeComponentExtension);
    },
    "call_RemoveComponentExtension": (retPtr: Pointer, extensionId: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.removeComponentExtension(A.H.get<object>(extensionId));
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveComponentExtension": (
      retPtr: Pointer,
      errPtr: Pointer,
      extensionId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.removeComponentExtension(A.H.get<object>(extensionId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveFuseboxTempDir": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "removeFuseboxTempDir" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveFuseboxTempDir": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.removeFuseboxTempDir);
    },
    "call_RemoveFuseboxTempDir": (retPtr: Pointer, fuseboxFilePath: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.removeFuseboxTempDir(A.H.get<object>(fuseboxFilePath));
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveFuseboxTempDir": (
      retPtr: Pointer,
      errPtr: Pointer,
      fuseboxFilePath: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.removeFuseboxTempDir(A.H.get<object>(fuseboxFilePath));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemovePrinter": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "removePrinter" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemovePrinter": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.removePrinter);
    },
    "call_RemovePrinter": (retPtr: Pointer, printerId: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.removePrinter(A.H.get<object>(printerId));
    },
    "try_RemovePrinter": (retPtr: Pointer, errPtr: Pointer, printerId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.removePrinter(A.H.get<object>(printerId));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ResetHoldingSpace": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "resetHoldingSpace" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ResetHoldingSpace": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.resetHoldingSpace);
    },
    "call_ResetHoldingSpace": (retPtr: Pointer, options: Pointer): void => {
      const options_ffi = {};

      if (A.load.Bool(options + 1)) {
        options_ffi["markTimeOfFirstAdd"] = A.load.Bool(options + 0);
      }

      const _ret = WEBEXT.autotestPrivate.resetHoldingSpace(options_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ResetHoldingSpace": (retPtr: Pointer, errPtr: Pointer, options: Pointer): heap.Ref<boolean> => {
      try {
        const options_ffi = {};

        if (A.load.Bool(options + 1)) {
          options_ffi["markTimeOfFirstAdd"] = A.load.Bool(options + 0);
        }

        const _ret = WEBEXT.autotestPrivate.resetHoldingSpace(options_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Restart": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "restart" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Restart": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.restart);
    },
    "call_Restart": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.restart();
    },
    "try_Restart": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.restart();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunCrostiniInstaller": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "runCrostiniInstaller" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunCrostiniInstaller": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.runCrostiniInstaller);
    },
    "call_RunCrostiniInstaller": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.runCrostiniInstaller();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunCrostiniInstaller": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.runCrostiniInstaller();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunCrostiniUninstaller": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "runCrostiniUninstaller" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunCrostiniUninstaller": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.runCrostiniUninstaller);
    },
    "call_RunCrostiniUninstaller": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.runCrostiniUninstaller();
      A.store.Ref(retPtr, _ret);
    },
    "try_RunCrostiniUninstaller": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.runCrostiniUninstaller();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendArcOverlayColor": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "sendArcOverlayColor" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendArcOverlayColor": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.sendArcOverlayColor);
    },
    "call_SendArcOverlayColor": (retPtr: Pointer, color: number, theme: number): void => {
      const _ret = WEBEXT.autotestPrivate.sendArcOverlayColor(
        color,
        theme > 0 && theme <= 6
          ? ["TonalSpot", "Vibrant", "Expressive", "Spritz", "Rainbow", "FruitSalad"][theme - 1]
          : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_SendArcOverlayColor": (retPtr: Pointer, errPtr: Pointer, color: number, theme: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.sendArcOverlayColor(
          color,
          theme > 0 && theme <= 6
            ? ["TonalSpot", "Vibrant", "Expressive", "Spritz", "Rainbow", "FruitSalad"][theme - 1]
            : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendAssistantTextQuery": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "sendAssistantTextQuery" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendAssistantTextQuery": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.sendAssistantTextQuery);
    },
    "call_SendAssistantTextQuery": (retPtr: Pointer, query: heap.Ref<object>, timeout_ms: number): void => {
      const _ret = WEBEXT.autotestPrivate.sendAssistantTextQuery(A.H.get<object>(query), timeout_ms);
      A.store.Ref(retPtr, _ret);
    },
    "try_SendAssistantTextQuery": (
      retPtr: Pointer,
      errPtr: Pointer,
      query: heap.Ref<object>,
      timeout_ms: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.sendAssistantTextQuery(A.H.get<object>(query), timeout_ms);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetAllowedPref": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setAllowedPref" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetAllowedPref": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setAllowedPref);
    },
    "call_SetAllowedPref": (retPtr: Pointer, pref_name: heap.Ref<object>, value: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.setAllowedPref(A.H.get<object>(pref_name), A.H.get<object>(value));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetAllowedPref": (
      retPtr: Pointer,
      errPtr: Pointer,
      pref_name: heap.Ref<object>,
      value: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setAllowedPref(A.H.get<object>(pref_name), A.H.get<object>(value));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetAppWindowState": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setAppWindowState" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetAppWindowState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setAppWindowState);
    },
    "call_SetAppWindowState": (retPtr: Pointer, id: number, change: Pointer, wait: heap.Ref<boolean>): void => {
      const change_ffi = {};

      change_ffi["eventType"] = A.load.Enum(change + 0, [
        "WMEventNormal",
        "WMEventMaximize",
        "WMEventMinimize",
        "WMEventFullscreen",
        "WMEventSnapPrimary",
        "WMEventSnapSecondary",
        "WMEventFloat",
      ]);
      if (A.load.Bool(change + 5)) {
        change_ffi["failIfNoChange"] = A.load.Bool(change + 4);
      }

      const _ret = WEBEXT.autotestPrivate.setAppWindowState(id, change_ffi, wait === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetAppWindowState": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: number,
      change: Pointer,
      wait: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const change_ffi = {};

        change_ffi["eventType"] = A.load.Enum(change + 0, [
          "WMEventNormal",
          "WMEventMaximize",
          "WMEventMinimize",
          "WMEventFullscreen",
          "WMEventSnapPrimary",
          "WMEventSnapSecondary",
          "WMEventFloat",
        ]);
        if (A.load.Bool(change + 5)) {
          change_ffi["failIfNoChange"] = A.load.Bool(change + 4);
        }

        const _ret = WEBEXT.autotestPrivate.setAppWindowState(id, change_ffi, wait === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetArcAppWindowFocus": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setArcAppWindowFocus" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetArcAppWindowFocus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setArcAppWindowFocus);
    },
    "call_SetArcAppWindowFocus": (retPtr: Pointer, packageName: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.setArcAppWindowFocus(A.H.get<object>(packageName));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetArcAppWindowFocus": (
      retPtr: Pointer,
      errPtr: Pointer,
      packageName: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setArcAppWindowFocus(A.H.get<object>(packageName));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetArcInteractiveState": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setArcInteractiveState" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetArcInteractiveState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setArcInteractiveState);
    },
    "call_SetArcInteractiveState": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.setArcInteractiveState(enabled === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetArcInteractiveState": (retPtr: Pointer, errPtr: Pointer, enabled: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setArcInteractiveState(enabled === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetArcTouchMode": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setArcTouchMode" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetArcTouchMode": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setArcTouchMode);
    },
    "call_SetArcTouchMode": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.setArcTouchMode(enabled === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetArcTouchMode": (retPtr: Pointer, errPtr: Pointer, enabled: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setArcTouchMode(enabled === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetAssistantEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setAssistantEnabled" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetAssistantEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setAssistantEnabled);
    },
    "call_SetAssistantEnabled": (retPtr: Pointer, enabled: heap.Ref<boolean>, timeout_ms: number): void => {
      const _ret = WEBEXT.autotestPrivate.setAssistantEnabled(enabled === A.H.TRUE, timeout_ms);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetAssistantEnabled": (
      retPtr: Pointer,
      errPtr: Pointer,
      enabled: heap.Ref<boolean>,
      timeout_ms: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setAssistantEnabled(enabled === A.H.TRUE, timeout_ms);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetClipboardTextData": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setClipboardTextData" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetClipboardTextData": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setClipboardTextData);
    },
    "call_SetClipboardTextData": (retPtr: Pointer, data: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.setClipboardTextData(A.H.get<object>(data));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetClipboardTextData": (retPtr: Pointer, errPtr: Pointer, data: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setClipboardTextData(A.H.get<object>(data));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetCrostiniAppScaled": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setCrostiniAppScaled" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetCrostiniAppScaled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setCrostiniAppScaled);
    },
    "call_SetCrostiniAppScaled": (retPtr: Pointer, appId: heap.Ref<object>, scaled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.setCrostiniAppScaled(A.H.get<object>(appId), scaled === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetCrostiniAppScaled": (
      retPtr: Pointer,
      errPtr: Pointer,
      appId: heap.Ref<object>,
      scaled: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setCrostiniAppScaled(A.H.get<object>(appId), scaled === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetCrostiniEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setCrostiniEnabled" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetCrostiniEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setCrostiniEnabled);
    },
    "call_SetCrostiniEnabled": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.setCrostiniEnabled(enabled === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetCrostiniEnabled": (retPtr: Pointer, errPtr: Pointer, enabled: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setCrostiniEnabled(enabled === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetMetricsEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setMetricsEnabled" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetMetricsEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setMetricsEnabled);
    },
    "call_SetMetricsEnabled": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.setMetricsEnabled(enabled === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetMetricsEnabled": (retPtr: Pointer, errPtr: Pointer, enabled: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setMetricsEnabled(enabled === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetMouseReverseScroll": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setMouseReverseScroll" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetMouseReverseScroll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setMouseReverseScroll);
    },
    "call_SetMouseReverseScroll": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.setMouseReverseScroll(enabled === A.H.TRUE);
    },
    "try_SetMouseReverseScroll": (retPtr: Pointer, errPtr: Pointer, enabled: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setMouseReverseScroll(enabled === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetMouseSensitivity": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setMouseSensitivity" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetMouseSensitivity": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setMouseSensitivity);
    },
    "call_SetMouseSensitivity": (retPtr: Pointer, value: number): void => {
      const _ret = WEBEXT.autotestPrivate.setMouseSensitivity(value);
    },
    "try_SetMouseSensitivity": (retPtr: Pointer, errPtr: Pointer, value: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setMouseSensitivity(value);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetNaturalScroll": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setNaturalScroll" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetNaturalScroll": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setNaturalScroll);
    },
    "call_SetNaturalScroll": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.setNaturalScroll(enabled === A.H.TRUE);
    },
    "try_SetNaturalScroll": (retPtr: Pointer, errPtr: Pointer, enabled: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setNaturalScroll(enabled === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetOverviewModeState": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setOverviewModeState" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetOverviewModeState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setOverviewModeState);
    },
    "call_SetOverviewModeState": (retPtr: Pointer, start: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.setOverviewModeState(start === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetOverviewModeState": (retPtr: Pointer, errPtr: Pointer, start: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setOverviewModeState(start === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPlayStoreEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setPlayStoreEnabled" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPlayStoreEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setPlayStoreEnabled);
    },
    "call_SetPlayStoreEnabled": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.setPlayStoreEnabled(enabled === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetPlayStoreEnabled": (retPtr: Pointer, errPtr: Pointer, enabled: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setPlayStoreEnabled(enabled === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPluginVMPolicy": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setPluginVMPolicy" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPluginVMPolicy": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setPluginVMPolicy);
    },
    "call_SetPluginVMPolicy": (
      retPtr: Pointer,
      imageUrl: heap.Ref<object>,
      imageHash: heap.Ref<object>,
      licenseKey: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.autotestPrivate.setPluginVMPolicy(
        A.H.get<object>(imageUrl),
        A.H.get<object>(imageHash),
        A.H.get<object>(licenseKey)
      );
    },
    "try_SetPluginVMPolicy": (
      retPtr: Pointer,
      errPtr: Pointer,
      imageUrl: heap.Ref<object>,
      imageHash: heap.Ref<object>,
      licenseKey: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setPluginVMPolicy(
          A.H.get<object>(imageUrl),
          A.H.get<object>(imageHash),
          A.H.get<object>(licenseKey)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPrimaryButtonRight": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setPrimaryButtonRight" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPrimaryButtonRight": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setPrimaryButtonRight);
    },
    "call_SetPrimaryButtonRight": (retPtr: Pointer, right: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.setPrimaryButtonRight(right === A.H.TRUE);
    },
    "try_SetPrimaryButtonRight": (retPtr: Pointer, errPtr: Pointer, right: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setPrimaryButtonRight(right === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetShelfAlignment": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setShelfAlignment" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetShelfAlignment": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setShelfAlignment);
    },
    "call_SetShelfAlignment": (retPtr: Pointer, displayId: heap.Ref<object>, alignment: number): void => {
      const _ret = WEBEXT.autotestPrivate.setShelfAlignment(
        A.H.get<object>(displayId),
        alignment > 0 && alignment <= 3 ? ["Bottom", "Left", "Right"][alignment - 1] : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_SetShelfAlignment": (
      retPtr: Pointer,
      errPtr: Pointer,
      displayId: heap.Ref<object>,
      alignment: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setShelfAlignment(
          A.H.get<object>(displayId),
          alignment > 0 && alignment <= 3 ? ["Bottom", "Left", "Right"][alignment - 1] : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetShelfAutoHideBehavior": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setShelfAutoHideBehavior" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetShelfAutoHideBehavior": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setShelfAutoHideBehavior);
    },
    "call_SetShelfAutoHideBehavior": (
      retPtr: Pointer,
      displayId: heap.Ref<object>,
      behavior: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.autotestPrivate.setShelfAutoHideBehavior(
        A.H.get<object>(displayId),
        A.H.get<object>(behavior)
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_SetShelfAutoHideBehavior": (
      retPtr: Pointer,
      errPtr: Pointer,
      displayId: heap.Ref<object>,
      behavior: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setShelfAutoHideBehavior(
          A.H.get<object>(displayId),
          A.H.get<object>(behavior)
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetShelfIconPin": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setShelfIconPin" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetShelfIconPin": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setShelfIconPin);
    },
    "call_SetShelfIconPin": (retPtr: Pointer, updateParams: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.setShelfIconPin(A.H.get<object>(updateParams));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetShelfIconPin": (retPtr: Pointer, errPtr: Pointer, updateParams: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setShelfIconPin(A.H.get<object>(updateParams));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetTabletModeEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setTabletModeEnabled" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetTabletModeEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setTabletModeEnabled);
    },
    "call_SetTabletModeEnabled": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.setTabletModeEnabled(enabled === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetTabletModeEnabled": (retPtr: Pointer, errPtr: Pointer, enabled: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setTabletModeEnabled(enabled === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetTapDragging": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setTapDragging" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetTapDragging": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setTapDragging);
    },
    "call_SetTapDragging": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.setTapDragging(enabled === A.H.TRUE);
    },
    "try_SetTapDragging": (retPtr: Pointer, errPtr: Pointer, enabled: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setTapDragging(enabled === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetTapToClick": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setTapToClick" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetTapToClick": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setTapToClick);
    },
    "call_SetTapToClick": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.setTapToClick(enabled === A.H.TRUE);
    },
    "try_SetTapToClick": (retPtr: Pointer, errPtr: Pointer, enabled: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setTapToClick(enabled === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetThreeFingerClick": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setThreeFingerClick" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetThreeFingerClick": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setThreeFingerClick);
    },
    "call_SetThreeFingerClick": (retPtr: Pointer, enabled: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.setThreeFingerClick(enabled === A.H.TRUE);
    },
    "try_SetThreeFingerClick": (retPtr: Pointer, errPtr: Pointer, enabled: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setThreeFingerClick(enabled === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetTouchpadSensitivity": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setTouchpadSensitivity" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetTouchpadSensitivity": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setTouchpadSensitivity);
    },
    "call_SetTouchpadSensitivity": (retPtr: Pointer, value: number): void => {
      const _ret = WEBEXT.autotestPrivate.setTouchpadSensitivity(value);
    },
    "try_SetTouchpadSensitivity": (retPtr: Pointer, errPtr: Pointer, value: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setTouchpadSensitivity(value);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetWhitelistedPref": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setWhitelistedPref" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetWhitelistedPref": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setWhitelistedPref);
    },
    "call_SetWhitelistedPref": (retPtr: Pointer, pref_name: heap.Ref<object>, value: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.setWhitelistedPref(A.H.get<object>(pref_name), A.H.get<object>(value));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetWhitelistedPref": (
      retPtr: Pointer,
      errPtr: Pointer,
      pref_name: heap.Ref<object>,
      value: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.setWhitelistedPref(A.H.get<object>(pref_name), A.H.get<object>(value));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetWindowBounds": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "setWindowBounds" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetWindowBounds": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.setWindowBounds);
    },
    "call_SetWindowBounds": (retPtr: Pointer, id: number, bounds: Pointer, displayId: heap.Ref<object>): void => {
      const bounds_ffi = {};

      if (A.load.Bool(bounds + 32)) {
        bounds_ffi["left"] = A.load.Float64(bounds + 0);
      }
      if (A.load.Bool(bounds + 33)) {
        bounds_ffi["top"] = A.load.Float64(bounds + 8);
      }
      if (A.load.Bool(bounds + 34)) {
        bounds_ffi["width"] = A.load.Float64(bounds + 16);
      }
      if (A.load.Bool(bounds + 35)) {
        bounds_ffi["height"] = A.load.Float64(bounds + 24);
      }

      const _ret = WEBEXT.autotestPrivate.setWindowBounds(id, bounds_ffi, A.H.get<object>(displayId));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetWindowBounds": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: number,
      bounds: Pointer,
      displayId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const bounds_ffi = {};

        if (A.load.Bool(bounds + 32)) {
          bounds_ffi["left"] = A.load.Float64(bounds + 0);
        }
        if (A.load.Bool(bounds + 33)) {
          bounds_ffi["top"] = A.load.Float64(bounds + 8);
        }
        if (A.load.Bool(bounds + 34)) {
          bounds_ffi["width"] = A.load.Float64(bounds + 16);
        }
        if (A.load.Bool(bounds + 35)) {
          bounds_ffi["height"] = A.load.Float64(bounds + 24);
        }

        const _ret = WEBEXT.autotestPrivate.setWindowBounds(id, bounds_ffi, A.H.get<object>(displayId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ShowPluginVMInstaller": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "showPluginVMInstaller" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ShowPluginVMInstaller": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.showPluginVMInstaller);
    },
    "call_ShowPluginVMInstaller": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.showPluginVMInstaller();
    },
    "try_ShowPluginVMInstaller": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.showPluginVMInstaller();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ShowVirtualKeyboardIfEnabled": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "showVirtualKeyboardIfEnabled" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ShowVirtualKeyboardIfEnabled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.showVirtualKeyboardIfEnabled);
    },
    "call_ShowVirtualKeyboardIfEnabled": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.showVirtualKeyboardIfEnabled();
    },
    "try_ShowVirtualKeyboardIfEnabled": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.showVirtualKeyboardIfEnabled();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Shutdown": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "shutdown" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Shutdown": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.shutdown);
    },
    "call_Shutdown": (retPtr: Pointer, force: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.autotestPrivate.shutdown(force === A.H.TRUE);
    },
    "try_Shutdown": (retPtr: Pointer, errPtr: Pointer, force: heap.Ref<boolean>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.shutdown(force === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SimulateAsanMemoryBug": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "simulateAsanMemoryBug" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SimulateAsanMemoryBug": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.simulateAsanMemoryBug);
    },
    "call_SimulateAsanMemoryBug": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.simulateAsanMemoryBug();
    },
    "try_SimulateAsanMemoryBug": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.simulateAsanMemoryBug();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartArc": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "startArc" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartArc": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.startArc);
    },
    "call_StartArc": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.startArc();
      A.store.Ref(retPtr, _ret);
    },
    "try_StartArc": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.startArc();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartFrameCounting": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "startFrameCounting" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartFrameCounting": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.startFrameCounting);
    },
    "call_StartFrameCounting": (retPtr: Pointer, bucketSizeInSeconds: number): void => {
      const _ret = WEBEXT.autotestPrivate.startFrameCounting(bucketSizeInSeconds);
      A.store.Ref(retPtr, _ret);
    },
    "try_StartFrameCounting": (retPtr: Pointer, errPtr: Pointer, bucketSizeInSeconds: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.startFrameCounting(bucketSizeInSeconds);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartLoginEventRecorderDataCollection": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "startLoginEventRecorderDataCollection" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartLoginEventRecorderDataCollection": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.startLoginEventRecorderDataCollection);
    },
    "call_StartLoginEventRecorderDataCollection": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.startLoginEventRecorderDataCollection();
      A.store.Ref(retPtr, _ret);
    },
    "try_StartLoginEventRecorderDataCollection": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.startLoginEventRecorderDataCollection();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartSmoothnessTracking": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "startSmoothnessTracking" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartSmoothnessTracking": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.startSmoothnessTracking);
    },
    "call_StartSmoothnessTracking": (
      retPtr: Pointer,
      displayId: heap.Ref<object>,
      throughputIntervalMs: number
    ): void => {
      const _ret = WEBEXT.autotestPrivate.startSmoothnessTracking(A.H.get<object>(displayId), throughputIntervalMs);
      A.store.Ref(retPtr, _ret);
    },
    "try_StartSmoothnessTracking": (
      retPtr: Pointer,
      errPtr: Pointer,
      displayId: heap.Ref<object>,
      throughputIntervalMs: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.startSmoothnessTracking(A.H.get<object>(displayId), throughputIntervalMs);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartThroughputTrackerDataCollection": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "startThroughputTrackerDataCollection" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartThroughputTrackerDataCollection": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.startThroughputTrackerDataCollection);
    },
    "call_StartThroughputTrackerDataCollection": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.startThroughputTrackerDataCollection();
      A.store.Ref(retPtr, _ret);
    },
    "try_StartThroughputTrackerDataCollection": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.startThroughputTrackerDataCollection();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StopArc": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "stopArc" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StopArc": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.stopArc);
    },
    "call_StopArc": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.stopArc();
      A.store.Ref(retPtr, _ret);
    },
    "try_StopArc": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.stopArc();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StopFrameCounting": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "stopFrameCounting" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StopFrameCounting": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.stopFrameCounting);
    },
    "call_StopFrameCounting": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.stopFrameCounting();
      A.store.Ref(retPtr, _ret);
    },
    "try_StopFrameCounting": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.stopFrameCounting();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StopSmoothnessTracking": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "stopSmoothnessTracking" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StopSmoothnessTracking": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.stopSmoothnessTracking);
    },
    "call_StopSmoothnessTracking": (retPtr: Pointer, displayId: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.stopSmoothnessTracking(A.H.get<object>(displayId));
      A.store.Ref(retPtr, _ret);
    },
    "try_StopSmoothnessTracking": (
      retPtr: Pointer,
      errPtr: Pointer,
      displayId: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.stopSmoothnessTracking(A.H.get<object>(displayId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StopThroughputTrackerDataCollection": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "stopThroughputTrackerDataCollection" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StopThroughputTrackerDataCollection": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.stopThroughputTrackerDataCollection);
    },
    "call_StopThroughputTrackerDataCollection": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.stopThroughputTrackerDataCollection();
      A.store.Ref(retPtr, _ret);
    },
    "try_StopThroughputTrackerDataCollection": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.stopThroughputTrackerDataCollection();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SwapWindowsInSplitView": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "swapWindowsInSplitView" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SwapWindowsInSplitView": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.swapWindowsInSplitView);
    },
    "call_SwapWindowsInSplitView": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.swapWindowsInSplitView();
      A.store.Ref(retPtr, _ret);
    },
    "try_SwapWindowsInSplitView": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.swapWindowsInSplitView();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_TakeScreenshot": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "takeScreenshot" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_TakeScreenshot": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.takeScreenshot);
    },
    "call_TakeScreenshot": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.takeScreenshot();
      A.store.Ref(retPtr, _ret);
    },
    "try_TakeScreenshot": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.takeScreenshot();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_TakeScreenshotForDisplay": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "takeScreenshotForDisplay" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_TakeScreenshotForDisplay": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.takeScreenshotForDisplay);
    },
    "call_TakeScreenshotForDisplay": (retPtr: Pointer, display_id: heap.Ref<object>): void => {
      const _ret = WEBEXT.autotestPrivate.takeScreenshotForDisplay(A.H.get<object>(display_id));
      A.store.Ref(retPtr, _ret);
    },
    "try_TakeScreenshotForDisplay": (
      retPtr: Pointer,
      errPtr: Pointer,
      display_id: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.takeScreenshotForDisplay(A.H.get<object>(display_id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdatePrinter": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "updatePrinter" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdatePrinter": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.updatePrinter);
    },
    "call_UpdatePrinter": (retPtr: Pointer, printer: Pointer): void => {
      const printer_ffi = {};

      printer_ffi["printerName"] = A.load.Ref(printer + 0, undefined);
      printer_ffi["printerId"] = A.load.Ref(printer + 4, undefined);
      printer_ffi["printerType"] = A.load.Ref(printer + 8, undefined);
      printer_ffi["printerDesc"] = A.load.Ref(printer + 12, undefined);
      printer_ffi["printerMakeAndModel"] = A.load.Ref(printer + 16, undefined);
      printer_ffi["printerUri"] = A.load.Ref(printer + 20, undefined);
      printer_ffi["printerPpd"] = A.load.Ref(printer + 24, undefined);

      const _ret = WEBEXT.autotestPrivate.updatePrinter(printer_ffi);
    },
    "try_UpdatePrinter": (retPtr: Pointer, errPtr: Pointer, printer: Pointer): heap.Ref<boolean> => {
      try {
        const printer_ffi = {};

        printer_ffi["printerName"] = A.load.Ref(printer + 0, undefined);
        printer_ffi["printerId"] = A.load.Ref(printer + 4, undefined);
        printer_ffi["printerType"] = A.load.Ref(printer + 8, undefined);
        printer_ffi["printerDesc"] = A.load.Ref(printer + 12, undefined);
        printer_ffi["printerMakeAndModel"] = A.load.Ref(printer + 16, undefined);
        printer_ffi["printerUri"] = A.load.Ref(printer + 20, undefined);
        printer_ffi["printerPpd"] = A.load.Ref(printer + 24, undefined);

        const _ret = WEBEXT.autotestPrivate.updatePrinter(printer_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_WaitForAmbientPhotoAnimation": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "waitForAmbientPhotoAnimation" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_WaitForAmbientPhotoAnimation": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.waitForAmbientPhotoAnimation);
    },
    "call_WaitForAmbientPhotoAnimation": (retPtr: Pointer, numCompletions: number, timeout: number): void => {
      const _ret = WEBEXT.autotestPrivate.waitForAmbientPhotoAnimation(numCompletions, timeout);
      A.store.Ref(retPtr, _ret);
    },
    "try_WaitForAmbientPhotoAnimation": (
      retPtr: Pointer,
      errPtr: Pointer,
      numCompletions: number,
      timeout: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.waitForAmbientPhotoAnimation(numCompletions, timeout);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_WaitForAmbientVideo": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "waitForAmbientVideo" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_WaitForAmbientVideo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.waitForAmbientVideo);
    },
    "call_WaitForAmbientVideo": (retPtr: Pointer, timeout: number): void => {
      const _ret = WEBEXT.autotestPrivate.waitForAmbientVideo(timeout);
      A.store.Ref(retPtr, _ret);
    },
    "try_WaitForAmbientVideo": (retPtr: Pointer, errPtr: Pointer, timeout: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.waitForAmbientVideo(timeout);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_WaitForAssistantQueryStatus": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "waitForAssistantQueryStatus" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_WaitForAssistantQueryStatus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.waitForAssistantQueryStatus);
    },
    "call_WaitForAssistantQueryStatus": (retPtr: Pointer, timeout_s: number): void => {
      const _ret = WEBEXT.autotestPrivate.waitForAssistantQueryStatus(timeout_s);
      A.store.Ref(retPtr, _ret);
    },
    "try_WaitForAssistantQueryStatus": (retPtr: Pointer, errPtr: Pointer, timeout_s: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.waitForAssistantQueryStatus(timeout_s);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_WaitForDisplayRotation": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "waitForDisplayRotation" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_WaitForDisplayRotation": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.waitForDisplayRotation);
    },
    "call_WaitForDisplayRotation": (retPtr: Pointer, displayId: heap.Ref<object>, rotation: number): void => {
      const _ret = WEBEXT.autotestPrivate.waitForDisplayRotation(
        A.H.get<object>(displayId),
        rotation > 0 && rotation <= 5
          ? ["RotateAny", "Rotate0", "Rotate90", "Rotate180", "Rotate270"][rotation - 1]
          : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_WaitForDisplayRotation": (
      retPtr: Pointer,
      errPtr: Pointer,
      displayId: heap.Ref<object>,
      rotation: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.waitForDisplayRotation(
          A.H.get<object>(displayId),
          rotation > 0 && rotation <= 5
            ? ["RotateAny", "Rotate0", "Rotate90", "Rotate180", "Rotate270"][rotation - 1]
            : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_WaitForLauncherState": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "waitForLauncherState" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_WaitForLauncherState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.waitForLauncherState);
    },
    "call_WaitForLauncherState": (retPtr: Pointer, launcherState: number): void => {
      const _ret = WEBEXT.autotestPrivate.waitForLauncherState(
        launcherState > 0 && launcherState <= 3
          ? ["Closed", "FullscreenAllApps", "FullscreenSearch"][launcherState - 1]
          : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_WaitForLauncherState": (retPtr: Pointer, errPtr: Pointer, launcherState: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.waitForLauncherState(
          launcherState > 0 && launcherState <= 3
            ? ["Closed", "FullscreenAllApps", "FullscreenSearch"][launcherState - 1]
            : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_WaitForOverviewState": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "waitForOverviewState" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_WaitForOverviewState": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.waitForOverviewState);
    },
    "call_WaitForOverviewState": (retPtr: Pointer, overviewState: number): void => {
      const _ret = WEBEXT.autotestPrivate.waitForOverviewState(
        overviewState > 0 && overviewState <= 2 ? ["Shown", "Hidden"][overviewState - 1] : undefined
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_WaitForOverviewState": (retPtr: Pointer, errPtr: Pointer, overviewState: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.waitForOverviewState(
          overviewState > 0 && overviewState <= 2 ? ["Shown", "Hidden"][overviewState - 1] : undefined
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_WaitForSystemWebAppsInstall": (): heap.Ref<boolean> => {
      if (WEBEXT?.autotestPrivate && "waitForSystemWebAppsInstall" in WEBEXT?.autotestPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_WaitForSystemWebAppsInstall": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.autotestPrivate.waitForSystemWebAppsInstall);
    },
    "call_WaitForSystemWebAppsInstall": (retPtr: Pointer): void => {
      const _ret = WEBEXT.autotestPrivate.waitForSystemWebAppsInstall();
      A.store.Ref(retPtr, _ret);
    },
    "try_WaitForSystemWebAppsInstall": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.autotestPrivate.waitForSystemWebAppsInstall();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
