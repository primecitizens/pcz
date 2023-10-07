import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/input/ime", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_AssistiveWindowButton": (ref: heap.Ref<string>): number => {
      const idx = ["undo", "addToDictionary"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_AssistiveWindowType": (ref: heap.Ref<string>): number => {
      const idx = ["undo"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_AssistiveWindowProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Bool(ptr + 8, false);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Ref(ptr + 0, x["announceString"]);
        A.store.Enum(ptr + 4, ["undo"].indexOf(x["type"] as string));
        A.store.Bool(ptr + 8, x["visible"] ? true : false);
      }
    },
    "load_AssistiveWindowProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["announceString"] = A.load.Ref(ptr + 0, undefined);
      x["type"] = A.load.Enum(ptr + 4, ["undo"]);
      x["visible"] = A.load.Bool(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_AutoCapitalizeType": (ref: heap.Ref<string>): number => {
      const idx = ["characters", "words", "sentences"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_ClearCompositionArgParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Int64(ptr + 0, x["contextID"] === undefined ? 0 : (x["contextID"] as number));
      }
    },
    "load_ClearCompositionArgParameters": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["contextID"] = A.load.Int64(ptr + 0);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_CommitTextArgParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Int64(ptr + 0, x["contextID"] === undefined ? 0 : (x["contextID"] as number));
        A.store.Ref(ptr + 8, x["text"]);
      }
    },
    "load_CommitTextArgParameters": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["contextID"] = A.load.Int64(ptr + 0);
      x["text"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DeleteSurroundingTextArgParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 32, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Int64(ptr + 16, 0);
        A.store.Int64(ptr + 24, 0);
      } else {
        A.store.Bool(ptr + 32, true);
        A.store.Int64(ptr + 0, x["contextID"] === undefined ? 0 : (x["contextID"] as number));
        A.store.Ref(ptr + 8, x["engineID"]);
        A.store.Int64(ptr + 16, x["length"] === undefined ? 0 : (x["length"] as number));
        A.store.Int64(ptr + 24, x["offset"] === undefined ? 0 : (x["offset"] as number));
      }
    },
    "load_DeleteSurroundingTextArgParameters": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["contextID"] = A.load.Int64(ptr + 0);
      x["engineID"] = A.load.Ref(ptr + 8, undefined);
      x["length"] = A.load.Int64(ptr + 16);
      x["offset"] = A.load.Int64(ptr + 24);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_InputContextType": (ref: heap.Ref<string>): number => {
      const idx = ["text", "search", "tel", "url", "email", "number", "password", "null"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_InputContext": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 24, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 5, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 17, false);
        A.store.Enum(ptr + 20, -1);
      } else {
        A.store.Bool(ptr + 24, true);
        A.store.Enum(ptr + 0, ["characters", "words", "sentences"].indexOf(x["autoCapitalize"] as string));
        A.store.Bool(ptr + 4, x["autoComplete"] ? true : false);
        A.store.Bool(ptr + 5, x["autoCorrect"] ? true : false);
        A.store.Int64(ptr + 8, x["contextID"] === undefined ? 0 : (x["contextID"] as number));
        A.store.Bool(ptr + 16, x["shouldDoLearning"] ? true : false);
        A.store.Bool(ptr + 17, x["spellCheck"] ? true : false);
        A.store.Enum(
          ptr + 20,
          ["text", "search", "tel", "url", "email", "number", "password", "null"].indexOf(x["type"] as string)
        );
      }
    },
    "load_InputContext": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["autoCapitalize"] = A.load.Enum(ptr + 0, ["characters", "words", "sentences"]);
      x["autoComplete"] = A.load.Bool(ptr + 4);
      x["autoCorrect"] = A.load.Bool(ptr + 5);
      x["contextID"] = A.load.Int64(ptr + 8);
      x["shouldDoLearning"] = A.load.Bool(ptr + 16);
      x["spellCheck"] = A.load.Bool(ptr + 17);
      x["type"] = A.load.Enum(ptr + 20, ["text", "search", "tel", "url", "email", "number", "password", "null"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_KeyboardEventType": (ref: heap.Ref<string>): number => {
      const idx = ["keyup", "keydown"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_KeyboardEvent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 50, false);
        A.store.Bool(ptr + 44, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 45, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 46, false);
        A.store.Bool(ptr + 2, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Bool(ptr + 47, false);
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 48, false);
        A.store.Int64(ptr + 24, 0);
        A.store.Ref(ptr + 32, undefined);
        A.store.Bool(ptr + 49, false);
        A.store.Bool(ptr + 36, false);
        A.store.Enum(ptr + 40, -1);
      } else {
        A.store.Bool(ptr + 50, true);
        A.store.Bool(ptr + 44, "altKey" in x ? true : false);
        A.store.Bool(ptr + 0, x["altKey"] ? true : false);
        A.store.Bool(ptr + 45, "altgrKey" in x ? true : false);
        A.store.Bool(ptr + 1, x["altgrKey"] ? true : false);
        A.store.Bool(ptr + 46, "capsLock" in x ? true : false);
        A.store.Bool(ptr + 2, x["capsLock"] ? true : false);
        A.store.Ref(ptr + 4, x["code"]);
        A.store.Bool(ptr + 47, "ctrlKey" in x ? true : false);
        A.store.Bool(ptr + 8, x["ctrlKey"] ? true : false);
        A.store.Ref(ptr + 12, x["extensionId"]);
        A.store.Ref(ptr + 16, x["key"]);
        A.store.Bool(ptr + 48, "keyCode" in x ? true : false);
        A.store.Int64(ptr + 24, x["keyCode"] === undefined ? 0 : (x["keyCode"] as number));
        A.store.Ref(ptr + 32, x["requestId"]);
        A.store.Bool(ptr + 49, "shiftKey" in x ? true : false);
        A.store.Bool(ptr + 36, x["shiftKey"] ? true : false);
        A.store.Enum(ptr + 40, ["keyup", "keydown"].indexOf(x["type"] as string));
      }
    },
    "load_KeyboardEvent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 44)) {
        x["altKey"] = A.load.Bool(ptr + 0);
      } else {
        delete x["altKey"];
      }
      if (A.load.Bool(ptr + 45)) {
        x["altgrKey"] = A.load.Bool(ptr + 1);
      } else {
        delete x["altgrKey"];
      }
      if (A.load.Bool(ptr + 46)) {
        x["capsLock"] = A.load.Bool(ptr + 2);
      } else {
        delete x["capsLock"];
      }
      x["code"] = A.load.Ref(ptr + 4, undefined);
      if (A.load.Bool(ptr + 47)) {
        x["ctrlKey"] = A.load.Bool(ptr + 8);
      } else {
        delete x["ctrlKey"];
      }
      x["extensionId"] = A.load.Ref(ptr + 12, undefined);
      x["key"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 48)) {
        x["keyCode"] = A.load.Int64(ptr + 24);
      } else {
        delete x["keyCode"];
      }
      x["requestId"] = A.load.Ref(ptr + 32, undefined);
      if (A.load.Bool(ptr + 49)) {
        x["shiftKey"] = A.load.Bool(ptr + 36);
      } else {
        delete x["shiftKey"];
      }
      x["type"] = A.load.Enum(ptr + 40, ["keyup", "keydown"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_MenuItemStyle": (ref: heap.Ref<string>): number => {
      const idx = ["check", "radio", "separator"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_MenuItem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 17, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 18, false);
        A.store.Bool(ptr + 1, false);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Enum(ptr + 12, -1);
        A.store.Bool(ptr + 19, false);
        A.store.Bool(ptr + 16, false);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Bool(ptr + 17, "checked" in x ? true : false);
        A.store.Bool(ptr + 0, x["checked"] ? true : false);
        A.store.Bool(ptr + 18, "enabled" in x ? true : false);
        A.store.Bool(ptr + 1, x["enabled"] ? true : false);
        A.store.Ref(ptr + 4, x["id"]);
        A.store.Ref(ptr + 8, x["label"]);
        A.store.Enum(ptr + 12, ["check", "radio", "separator"].indexOf(x["style"] as string));
        A.store.Bool(ptr + 19, "visible" in x ? true : false);
        A.store.Bool(ptr + 16, x["visible"] ? true : false);
      }
    },
    "load_MenuItem": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 17)) {
        x["checked"] = A.load.Bool(ptr + 0);
      } else {
        delete x["checked"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["enabled"] = A.load.Bool(ptr + 1);
      } else {
        delete x["enabled"];
      }
      x["id"] = A.load.Ref(ptr + 4, undefined);
      x["label"] = A.load.Ref(ptr + 8, undefined);
      x["style"] = A.load.Enum(ptr + 12, ["check", "radio", "separator"]);
      if (A.load.Bool(ptr + 19)) {
        x["visible"] = A.load.Bool(ptr + 16);
      } else {
        delete x["visible"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_MenuParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["engineID"]);
        A.store.Ref(ptr + 4, x["items"]);
      }
    },
    "load_MenuParameters": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["engineID"] = A.load.Ref(ptr + 0, undefined);
      x["items"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_MouseButton": (ref: heap.Ref<string>): number => {
      const idx = ["left", "middle", "right"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_OnAssistiveWindowButtonClickedArgDetails": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Enum(ptr + 0, ["undo", "addToDictionary"].indexOf(x["buttonID"] as string));
        A.store.Enum(ptr + 4, ["undo"].indexOf(x["windowType"] as string));
      }
    },
    "load_OnAssistiveWindowButtonClickedArgDetails": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["buttonID"] = A.load.Enum(ptr + 0, ["undo", "addToDictionary"]);
      x["windowType"] = A.load.Enum(ptr + 4, ["undo"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnSurroundingTextChangedArgSurroundingInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 28, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
        A.store.Int64(ptr + 16, 0);
        A.store.Ref(ptr + 24, undefined);
      } else {
        A.store.Bool(ptr + 28, true);
        A.store.Int64(ptr + 0, x["anchor"] === undefined ? 0 : (x["anchor"] as number));
        A.store.Int64(ptr + 8, x["focus"] === undefined ? 0 : (x["focus"] as number));
        A.store.Int64(ptr + 16, x["offset"] === undefined ? 0 : (x["offset"] as number));
        A.store.Ref(ptr + 24, x["text"]);
      }
    },
    "load_OnSurroundingTextChangedArgSurroundingInfo": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["anchor"] = A.load.Int64(ptr + 0);
      x["focus"] = A.load.Int64(ptr + 8);
      x["offset"] = A.load.Int64(ptr + 16);
      x["text"] = A.load.Ref(ptr + 24, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_ScreenType": (ref: heap.Ref<string>): number => {
      const idx = ["normal", "login", "lock", "secondary-login"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SendKeyEventsArgParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Int64(ptr + 0, x["contextID"] === undefined ? 0 : (x["contextID"] as number));
        A.store.Ref(ptr + 8, x["keyData"]);
      }
    },
    "load_SendKeyEventsArgParameters": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["contextID"] = A.load.Int64(ptr + 0);
      x["keyData"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetAssistiveWindowButtonHighlightedArgParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 24, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
        A.store.Int64(ptr + 8, 0);
        A.store.Bool(ptr + 16, false);
        A.store.Enum(ptr + 20, -1);
      } else {
        A.store.Bool(ptr + 24, true);
        A.store.Ref(ptr + 0, x["announceString"]);
        A.store.Enum(ptr + 4, ["undo", "addToDictionary"].indexOf(x["buttonID"] as string));
        A.store.Int64(ptr + 8, x["contextID"] === undefined ? 0 : (x["contextID"] as number));
        A.store.Bool(ptr + 16, x["highlighted"] ? true : false);
        A.store.Enum(ptr + 20, ["undo"].indexOf(x["windowType"] as string));
      }
    },
    "load_SetAssistiveWindowButtonHighlightedArgParameters": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["announceString"] = A.load.Ref(ptr + 0, undefined);
      x["buttonID"] = A.load.Enum(ptr + 4, ["undo", "addToDictionary"]);
      x["contextID"] = A.load.Int64(ptr + 8);
      x["highlighted"] = A.load.Bool(ptr + 16);
      x["windowType"] = A.load.Enum(ptr + 20, ["undo"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetAssistiveWindowPropertiesArgParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Int64(ptr + 0, 0);

        A.store.Bool(ptr + 8 + 9, false);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Enum(ptr + 8 + 4, -1);
        A.store.Bool(ptr + 8 + 8, false);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Int64(ptr + 0, x["contextID"] === undefined ? 0 : (x["contextID"] as number));

        if (typeof x["properties"] === "undefined") {
          A.store.Bool(ptr + 8 + 9, false);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Enum(ptr + 8 + 4, -1);
          A.store.Bool(ptr + 8 + 8, false);
        } else {
          A.store.Bool(ptr + 8 + 9, true);
          A.store.Ref(ptr + 8 + 0, x["properties"]["announceString"]);
          A.store.Enum(ptr + 8 + 4, ["undo"].indexOf(x["properties"]["type"] as string));
          A.store.Bool(ptr + 8 + 8, x["properties"]["visible"] ? true : false);
        }
      }
    },
    "load_SetAssistiveWindowPropertiesArgParameters": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["contextID"] = A.load.Int64(ptr + 0);
      if (A.load.Bool(ptr + 8 + 9)) {
        x["properties"] = {};
        x["properties"]["announceString"] = A.load.Ref(ptr + 8 + 0, undefined);
        x["properties"]["type"] = A.load.Enum(ptr + 8 + 4, ["undo"]);
        x["properties"]["visible"] = A.load.Bool(ptr + 8 + 8);
      } else {
        delete x["properties"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_WindowPosition": (ref: heap.Ref<string>): number => {
      const idx = ["cursor", "composition"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SetCandidateWindowPropertiesArgParametersFieldProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 55, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 48, false);
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 49, false);
        A.store.Int64(ptr + 8, 0);
        A.store.Bool(ptr + 50, false);
        A.store.Bool(ptr + 16, false);
        A.store.Bool(ptr + 51, false);
        A.store.Int64(ptr + 24, 0);
        A.store.Bool(ptr + 52, false);
        A.store.Int64(ptr + 32, 0);
        A.store.Bool(ptr + 53, false);
        A.store.Bool(ptr + 40, false);
        A.store.Bool(ptr + 54, false);
        A.store.Bool(ptr + 41, false);
        A.store.Enum(ptr + 44, -1);
      } else {
        A.store.Bool(ptr + 55, true);
        A.store.Ref(ptr + 0, x["auxiliaryText"]);
        A.store.Bool(ptr + 48, "auxiliaryTextVisible" in x ? true : false);
        A.store.Bool(ptr + 4, x["auxiliaryTextVisible"] ? true : false);
        A.store.Bool(ptr + 49, "currentCandidateIndex" in x ? true : false);
        A.store.Int64(ptr + 8, x["currentCandidateIndex"] === undefined ? 0 : (x["currentCandidateIndex"] as number));
        A.store.Bool(ptr + 50, "cursorVisible" in x ? true : false);
        A.store.Bool(ptr + 16, x["cursorVisible"] ? true : false);
        A.store.Bool(ptr + 51, "pageSize" in x ? true : false);
        A.store.Int64(ptr + 24, x["pageSize"] === undefined ? 0 : (x["pageSize"] as number));
        A.store.Bool(ptr + 52, "totalCandidates" in x ? true : false);
        A.store.Int64(ptr + 32, x["totalCandidates"] === undefined ? 0 : (x["totalCandidates"] as number));
        A.store.Bool(ptr + 53, "vertical" in x ? true : false);
        A.store.Bool(ptr + 40, x["vertical"] ? true : false);
        A.store.Bool(ptr + 54, "visible" in x ? true : false);
        A.store.Bool(ptr + 41, x["visible"] ? true : false);
        A.store.Enum(ptr + 44, ["cursor", "composition"].indexOf(x["windowPosition"] as string));
      }
    },
    "load_SetCandidateWindowPropertiesArgParametersFieldProperties": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["auxiliaryText"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 48)) {
        x["auxiliaryTextVisible"] = A.load.Bool(ptr + 4);
      } else {
        delete x["auxiliaryTextVisible"];
      }
      if (A.load.Bool(ptr + 49)) {
        x["currentCandidateIndex"] = A.load.Int64(ptr + 8);
      } else {
        delete x["currentCandidateIndex"];
      }
      if (A.load.Bool(ptr + 50)) {
        x["cursorVisible"] = A.load.Bool(ptr + 16);
      } else {
        delete x["cursorVisible"];
      }
      if (A.load.Bool(ptr + 51)) {
        x["pageSize"] = A.load.Int64(ptr + 24);
      } else {
        delete x["pageSize"];
      }
      if (A.load.Bool(ptr + 52)) {
        x["totalCandidates"] = A.load.Int64(ptr + 32);
      } else {
        delete x["totalCandidates"];
      }
      if (A.load.Bool(ptr + 53)) {
        x["vertical"] = A.load.Bool(ptr + 40);
      } else {
        delete x["vertical"];
      }
      if (A.load.Bool(ptr + 54)) {
        x["visible"] = A.load.Bool(ptr + 41);
      } else {
        delete x["visible"];
      }
      x["windowPosition"] = A.load.Enum(ptr + 44, ["cursor", "composition"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetCandidateWindowPropertiesArgParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 64, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 8 + 55, false);
        A.store.Ref(ptr + 8 + 0, undefined);
        A.store.Bool(ptr + 8 + 48, false);
        A.store.Bool(ptr + 8 + 4, false);
        A.store.Bool(ptr + 8 + 49, false);
        A.store.Int64(ptr + 8 + 8, 0);
        A.store.Bool(ptr + 8 + 50, false);
        A.store.Bool(ptr + 8 + 16, false);
        A.store.Bool(ptr + 8 + 51, false);
        A.store.Int64(ptr + 8 + 24, 0);
        A.store.Bool(ptr + 8 + 52, false);
        A.store.Int64(ptr + 8 + 32, 0);
        A.store.Bool(ptr + 8 + 53, false);
        A.store.Bool(ptr + 8 + 40, false);
        A.store.Bool(ptr + 8 + 54, false);
        A.store.Bool(ptr + 8 + 41, false);
        A.store.Enum(ptr + 8 + 44, -1);
      } else {
        A.store.Bool(ptr + 64, true);
        A.store.Ref(ptr + 0, x["engineID"]);

        if (typeof x["properties"] === "undefined") {
          A.store.Bool(ptr + 8 + 55, false);
          A.store.Ref(ptr + 8 + 0, undefined);
          A.store.Bool(ptr + 8 + 48, false);
          A.store.Bool(ptr + 8 + 4, false);
          A.store.Bool(ptr + 8 + 49, false);
          A.store.Int64(ptr + 8 + 8, 0);
          A.store.Bool(ptr + 8 + 50, false);
          A.store.Bool(ptr + 8 + 16, false);
          A.store.Bool(ptr + 8 + 51, false);
          A.store.Int64(ptr + 8 + 24, 0);
          A.store.Bool(ptr + 8 + 52, false);
          A.store.Int64(ptr + 8 + 32, 0);
          A.store.Bool(ptr + 8 + 53, false);
          A.store.Bool(ptr + 8 + 40, false);
          A.store.Bool(ptr + 8 + 54, false);
          A.store.Bool(ptr + 8 + 41, false);
          A.store.Enum(ptr + 8 + 44, -1);
        } else {
          A.store.Bool(ptr + 8 + 55, true);
          A.store.Ref(ptr + 8 + 0, x["properties"]["auxiliaryText"]);
          A.store.Bool(ptr + 8 + 48, "auxiliaryTextVisible" in x["properties"] ? true : false);
          A.store.Bool(ptr + 8 + 4, x["properties"]["auxiliaryTextVisible"] ? true : false);
          A.store.Bool(ptr + 8 + 49, "currentCandidateIndex" in x["properties"] ? true : false);
          A.store.Int64(
            ptr + 8 + 8,
            x["properties"]["currentCandidateIndex"] === undefined
              ? 0
              : (x["properties"]["currentCandidateIndex"] as number)
          );
          A.store.Bool(ptr + 8 + 50, "cursorVisible" in x["properties"] ? true : false);
          A.store.Bool(ptr + 8 + 16, x["properties"]["cursorVisible"] ? true : false);
          A.store.Bool(ptr + 8 + 51, "pageSize" in x["properties"] ? true : false);
          A.store.Int64(
            ptr + 8 + 24,
            x["properties"]["pageSize"] === undefined ? 0 : (x["properties"]["pageSize"] as number)
          );
          A.store.Bool(ptr + 8 + 52, "totalCandidates" in x["properties"] ? true : false);
          A.store.Int64(
            ptr + 8 + 32,
            x["properties"]["totalCandidates"] === undefined ? 0 : (x["properties"]["totalCandidates"] as number)
          );
          A.store.Bool(ptr + 8 + 53, "vertical" in x["properties"] ? true : false);
          A.store.Bool(ptr + 8 + 40, x["properties"]["vertical"] ? true : false);
          A.store.Bool(ptr + 8 + 54, "visible" in x["properties"] ? true : false);
          A.store.Bool(ptr + 8 + 41, x["properties"]["visible"] ? true : false);
          A.store.Enum(ptr + 8 + 44, ["cursor", "composition"].indexOf(x["properties"]["windowPosition"] as string));
        }
      }
    },
    "load_SetCandidateWindowPropertiesArgParameters": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["engineID"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 8 + 55)) {
        x["properties"] = {};
        x["properties"]["auxiliaryText"] = A.load.Ref(ptr + 8 + 0, undefined);
        if (A.load.Bool(ptr + 8 + 48)) {
          x["properties"]["auxiliaryTextVisible"] = A.load.Bool(ptr + 8 + 4);
        } else {
          delete x["properties"]["auxiliaryTextVisible"];
        }
        if (A.load.Bool(ptr + 8 + 49)) {
          x["properties"]["currentCandidateIndex"] = A.load.Int64(ptr + 8 + 8);
        } else {
          delete x["properties"]["currentCandidateIndex"];
        }
        if (A.load.Bool(ptr + 8 + 50)) {
          x["properties"]["cursorVisible"] = A.load.Bool(ptr + 8 + 16);
        } else {
          delete x["properties"]["cursorVisible"];
        }
        if (A.load.Bool(ptr + 8 + 51)) {
          x["properties"]["pageSize"] = A.load.Int64(ptr + 8 + 24);
        } else {
          delete x["properties"]["pageSize"];
        }
        if (A.load.Bool(ptr + 8 + 52)) {
          x["properties"]["totalCandidates"] = A.load.Int64(ptr + 8 + 32);
        } else {
          delete x["properties"]["totalCandidates"];
        }
        if (A.load.Bool(ptr + 8 + 53)) {
          x["properties"]["vertical"] = A.load.Bool(ptr + 8 + 40);
        } else {
          delete x["properties"]["vertical"];
        }
        if (A.load.Bool(ptr + 8 + 54)) {
          x["properties"]["visible"] = A.load.Bool(ptr + 8 + 41);
        } else {
          delete x["properties"]["visible"];
        }
        x["properties"]["windowPosition"] = A.load.Enum(ptr + 8 + 44, ["cursor", "composition"]);
      } else {
        delete x["properties"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetCandidatesArgParametersFieldCandidatesElemFieldUsage": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["body"]);
        A.store.Ref(ptr + 4, x["title"]);
      }
    },
    "load_SetCandidatesArgParametersFieldCandidatesElemFieldUsage": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["body"] = A.load.Ref(ptr + 0, undefined);
      x["title"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetCandidatesArgParametersFieldCandidatesElem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 42, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Int64(ptr + 8, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 41, false);
        A.store.Int64(ptr + 24, 0);

        A.store.Bool(ptr + 32 + 8, false);
        A.store.Ref(ptr + 32 + 0, undefined);
        A.store.Ref(ptr + 32 + 4, undefined);
      } else {
        A.store.Bool(ptr + 42, true);
        A.store.Ref(ptr + 0, x["annotation"]);
        A.store.Ref(ptr + 4, x["candidate"]);
        A.store.Int64(ptr + 8, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Ref(ptr + 16, x["label"]);
        A.store.Bool(ptr + 41, "parentId" in x ? true : false);
        A.store.Int64(ptr + 24, x["parentId"] === undefined ? 0 : (x["parentId"] as number));

        if (typeof x["usage"] === "undefined") {
          A.store.Bool(ptr + 32 + 8, false);
          A.store.Ref(ptr + 32 + 0, undefined);
          A.store.Ref(ptr + 32 + 4, undefined);
        } else {
          A.store.Bool(ptr + 32 + 8, true);
          A.store.Ref(ptr + 32 + 0, x["usage"]["body"]);
          A.store.Ref(ptr + 32 + 4, x["usage"]["title"]);
        }
      }
    },
    "load_SetCandidatesArgParametersFieldCandidatesElem": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["annotation"] = A.load.Ref(ptr + 0, undefined);
      x["candidate"] = A.load.Ref(ptr + 4, undefined);
      x["id"] = A.load.Int64(ptr + 8);
      x["label"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 41)) {
        x["parentId"] = A.load.Int64(ptr + 24);
      } else {
        delete x["parentId"];
      }
      if (A.load.Bool(ptr + 32 + 8)) {
        x["usage"] = {};
        x["usage"]["body"] = A.load.Ref(ptr + 32 + 0, undefined);
        x["usage"]["title"] = A.load.Ref(ptr + 32 + 4, undefined);
      } else {
        delete x["usage"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetCandidatesArgParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["candidates"]);
        A.store.Int64(ptr + 8, x["contextID"] === undefined ? 0 : (x["contextID"] as number));
      }
    },
    "load_SetCandidatesArgParameters": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["candidates"] = A.load.Ref(ptr + 0, undefined);
      x["contextID"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_UnderlineStyle": (ref: heap.Ref<string>): number => {
      const idx = ["underline", "doubleUnderline", "noUnderline"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SetCompositionArgParametersFieldSegmentsElem": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Int64(ptr + 0, x["end"] === undefined ? 0 : (x["end"] as number));
        A.store.Int64(ptr + 8, x["start"] === undefined ? 0 : (x["start"] as number));
        A.store.Enum(ptr + 16, ["underline", "doubleUnderline", "noUnderline"].indexOf(x["style"] as string));
      }
    },
    "load_SetCompositionArgParametersFieldSegmentsElem": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["end"] = A.load.Int64(ptr + 0);
      x["start"] = A.load.Int64(ptr + 8);
      x["style"] = A.load.Enum(ptr + 16, ["underline", "doubleUnderline", "noUnderline"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetCompositionArgParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 46, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
        A.store.Ref(ptr + 16, undefined);
        A.store.Bool(ptr + 44, false);
        A.store.Int64(ptr + 24, 0);
        A.store.Bool(ptr + 45, false);
        A.store.Int64(ptr + 32, 0);
        A.store.Ref(ptr + 40, undefined);
      } else {
        A.store.Bool(ptr + 46, true);
        A.store.Int64(ptr + 0, x["contextID"] === undefined ? 0 : (x["contextID"] as number));
        A.store.Int64(ptr + 8, x["cursor"] === undefined ? 0 : (x["cursor"] as number));
        A.store.Ref(ptr + 16, x["segments"]);
        A.store.Bool(ptr + 44, "selectionEnd" in x ? true : false);
        A.store.Int64(ptr + 24, x["selectionEnd"] === undefined ? 0 : (x["selectionEnd"] as number));
        A.store.Bool(ptr + 45, "selectionStart" in x ? true : false);
        A.store.Int64(ptr + 32, x["selectionStart"] === undefined ? 0 : (x["selectionStart"] as number));
        A.store.Ref(ptr + 40, x["text"]);
      }
    },
    "load_SetCompositionArgParameters": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["contextID"] = A.load.Int64(ptr + 0);
      x["cursor"] = A.load.Int64(ptr + 8);
      x["segments"] = A.load.Ref(ptr + 16, undefined);
      if (A.load.Bool(ptr + 44)) {
        x["selectionEnd"] = A.load.Int64(ptr + 24);
      } else {
        delete x["selectionEnd"];
      }
      if (A.load.Bool(ptr + 45)) {
        x["selectionStart"] = A.load.Int64(ptr + 32);
      } else {
        delete x["selectionStart"];
      }
      x["text"] = A.load.Ref(ptr + 40, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SetCursorPositionArgParameters": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Int64(ptr + 0, x["candidateID"] === undefined ? 0 : (x["candidateID"] as number));
        A.store.Int64(ptr + 8, x["contextID"] === undefined ? 0 : (x["contextID"] as number));
      }
    },
    "load_SetCursorPositionArgParameters": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["candidateID"] = A.load.Int64(ptr + 0);
      x["contextID"] = A.load.Int64(ptr + 8);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_ClearComposition": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime && "clearComposition" in WEBEXT?.input?.ime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ClearComposition": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.clearComposition);
    },
    "call_ClearComposition": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["contextID"] = A.load.Int64(parameters + 0);

      const _ret = WEBEXT.input.ime.clearComposition(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_ClearComposition": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["contextID"] = A.load.Int64(parameters + 0);

        const _ret = WEBEXT.input.ime.clearComposition(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CommitText": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime && "commitText" in WEBEXT?.input?.ime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CommitText": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.commitText);
    },
    "call_CommitText": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["contextID"] = A.load.Int64(parameters + 0);
      parameters_ffi["text"] = A.load.Ref(parameters + 8, undefined);

      const _ret = WEBEXT.input.ime.commitText(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_CommitText": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["contextID"] = A.load.Int64(parameters + 0);
        parameters_ffi["text"] = A.load.Ref(parameters + 8, undefined);

        const _ret = WEBEXT.input.ime.commitText(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_DeleteSurroundingText": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime && "deleteSurroundingText" in WEBEXT?.input?.ime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DeleteSurroundingText": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.deleteSurroundingText);
    },
    "call_DeleteSurroundingText": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["contextID"] = A.load.Int64(parameters + 0);
      parameters_ffi["engineID"] = A.load.Ref(parameters + 8, undefined);
      parameters_ffi["length"] = A.load.Int64(parameters + 16);
      parameters_ffi["offset"] = A.load.Int64(parameters + 24);

      const _ret = WEBEXT.input.ime.deleteSurroundingText(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_DeleteSurroundingText": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["contextID"] = A.load.Int64(parameters + 0);
        parameters_ffi["engineID"] = A.load.Ref(parameters + 8, undefined);
        parameters_ffi["length"] = A.load.Int64(parameters + 16);
        parameters_ffi["offset"] = A.load.Int64(parameters + 24);

        const _ret = WEBEXT.input.ime.deleteSurroundingText(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HideInputView": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime && "hideInputView" in WEBEXT?.input?.ime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HideInputView": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.hideInputView);
    },
    "call_HideInputView": (retPtr: Pointer): void => {
      const _ret = WEBEXT.input.ime.hideInputView();
    },
    "try_HideInputView": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.hideInputView();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_KeyEventHandled": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime && "keyEventHandled" in WEBEXT?.input?.ime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_KeyEventHandled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.keyEventHandled);
    },
    "call_KeyEventHandled": (retPtr: Pointer, requestId: heap.Ref<object>, response: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.input.ime.keyEventHandled(A.H.get<object>(requestId), response === A.H.TRUE);
    },
    "try_KeyEventHandled": (
      retPtr: Pointer,
      errPtr: Pointer,
      requestId: heap.Ref<object>,
      response: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.keyEventHandled(A.H.get<object>(requestId), response === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnActivate": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onActivate && "addListener" in WEBEXT?.input?.ime?.onActivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnActivate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onActivate.addListener);
    },
    "call_OnActivate": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onActivate.addListener(A.H.get<object>(callback));
    },
    "try_OnActivate": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onActivate.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffActivate": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onActivate && "removeListener" in WEBEXT?.input?.ime?.onActivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffActivate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onActivate.removeListener);
    },
    "call_OffActivate": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onActivate.removeListener(A.H.get<object>(callback));
    },
    "try_OffActivate": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onActivate.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnActivate": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onActivate && "hasListener" in WEBEXT?.input?.ime?.onActivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnActivate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onActivate.hasListener);
    },
    "call_HasOnActivate": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onActivate.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnActivate": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onActivate.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnAssistiveWindowButtonClicked": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.input?.ime?.onAssistiveWindowButtonClicked &&
        "addListener" in WEBEXT?.input?.ime?.onAssistiveWindowButtonClicked
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnAssistiveWindowButtonClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onAssistiveWindowButtonClicked.addListener);
    },
    "call_OnAssistiveWindowButtonClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onAssistiveWindowButtonClicked.addListener(A.H.get<object>(callback));
    },
    "try_OnAssistiveWindowButtonClicked": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onAssistiveWindowButtonClicked.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffAssistiveWindowButtonClicked": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.input?.ime?.onAssistiveWindowButtonClicked &&
        "removeListener" in WEBEXT?.input?.ime?.onAssistiveWindowButtonClicked
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffAssistiveWindowButtonClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onAssistiveWindowButtonClicked.removeListener);
    },
    "call_OffAssistiveWindowButtonClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onAssistiveWindowButtonClicked.removeListener(A.H.get<object>(callback));
    },
    "try_OffAssistiveWindowButtonClicked": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onAssistiveWindowButtonClicked.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnAssistiveWindowButtonClicked": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.input?.ime?.onAssistiveWindowButtonClicked &&
        "hasListener" in WEBEXT?.input?.ime?.onAssistiveWindowButtonClicked
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnAssistiveWindowButtonClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onAssistiveWindowButtonClicked.hasListener);
    },
    "call_HasOnAssistiveWindowButtonClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onAssistiveWindowButtonClicked.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnAssistiveWindowButtonClicked": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onAssistiveWindowButtonClicked.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnBlur": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onBlur && "addListener" in WEBEXT?.input?.ime?.onBlur) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnBlur": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onBlur.addListener);
    },
    "call_OnBlur": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onBlur.addListener(A.H.get<object>(callback));
    },
    "try_OnBlur": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onBlur.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffBlur": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onBlur && "removeListener" in WEBEXT?.input?.ime?.onBlur) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffBlur": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onBlur.removeListener);
    },
    "call_OffBlur": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onBlur.removeListener(A.H.get<object>(callback));
    },
    "try_OffBlur": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onBlur.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnBlur": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onBlur && "hasListener" in WEBEXT?.input?.ime?.onBlur) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnBlur": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onBlur.hasListener);
    },
    "call_HasOnBlur": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onBlur.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnBlur": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onBlur.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnCandidateClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onCandidateClicked && "addListener" in WEBEXT?.input?.ime?.onCandidateClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnCandidateClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onCandidateClicked.addListener);
    },
    "call_OnCandidateClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onCandidateClicked.addListener(A.H.get<object>(callback));
    },
    "try_OnCandidateClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onCandidateClicked.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffCandidateClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onCandidateClicked && "removeListener" in WEBEXT?.input?.ime?.onCandidateClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffCandidateClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onCandidateClicked.removeListener);
    },
    "call_OffCandidateClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onCandidateClicked.removeListener(A.H.get<object>(callback));
    },
    "try_OffCandidateClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onCandidateClicked.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnCandidateClicked": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onCandidateClicked && "hasListener" in WEBEXT?.input?.ime?.onCandidateClicked) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnCandidateClicked": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onCandidateClicked.hasListener);
    },
    "call_HasOnCandidateClicked": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onCandidateClicked.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnCandidateClicked": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onCandidateClicked.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeactivated": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onDeactivated && "addListener" in WEBEXT?.input?.ime?.onDeactivated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeactivated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onDeactivated.addListener);
    },
    "call_OnDeactivated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onDeactivated.addListener(A.H.get<object>(callback));
    },
    "try_OnDeactivated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onDeactivated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeactivated": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onDeactivated && "removeListener" in WEBEXT?.input?.ime?.onDeactivated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeactivated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onDeactivated.removeListener);
    },
    "call_OffDeactivated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onDeactivated.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeactivated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onDeactivated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeactivated": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onDeactivated && "hasListener" in WEBEXT?.input?.ime?.onDeactivated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeactivated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onDeactivated.hasListener);
    },
    "call_HasOnDeactivated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onDeactivated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeactivated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onDeactivated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnFocus": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onFocus && "addListener" in WEBEXT?.input?.ime?.onFocus) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnFocus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onFocus.addListener);
    },
    "call_OnFocus": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onFocus.addListener(A.H.get<object>(callback));
    },
    "try_OnFocus": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onFocus.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffFocus": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onFocus && "removeListener" in WEBEXT?.input?.ime?.onFocus) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffFocus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onFocus.removeListener);
    },
    "call_OffFocus": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onFocus.removeListener(A.H.get<object>(callback));
    },
    "try_OffFocus": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onFocus.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnFocus": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onFocus && "hasListener" in WEBEXT?.input?.ime?.onFocus) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnFocus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onFocus.hasListener);
    },
    "call_HasOnFocus": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onFocus.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnFocus": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onFocus.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnInputContextUpdate": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onInputContextUpdate && "addListener" in WEBEXT?.input?.ime?.onInputContextUpdate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnInputContextUpdate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onInputContextUpdate.addListener);
    },
    "call_OnInputContextUpdate": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onInputContextUpdate.addListener(A.H.get<object>(callback));
    },
    "try_OnInputContextUpdate": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onInputContextUpdate.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffInputContextUpdate": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onInputContextUpdate && "removeListener" in WEBEXT?.input?.ime?.onInputContextUpdate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffInputContextUpdate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onInputContextUpdate.removeListener);
    },
    "call_OffInputContextUpdate": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onInputContextUpdate.removeListener(A.H.get<object>(callback));
    },
    "try_OffInputContextUpdate": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onInputContextUpdate.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnInputContextUpdate": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onInputContextUpdate && "hasListener" in WEBEXT?.input?.ime?.onInputContextUpdate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnInputContextUpdate": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onInputContextUpdate.hasListener);
    },
    "call_HasOnInputContextUpdate": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onInputContextUpdate.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnInputContextUpdate": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onInputContextUpdate.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnKeyEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onKeyEvent && "addListener" in WEBEXT?.input?.ime?.onKeyEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnKeyEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onKeyEvent.addListener);
    },
    "call_OnKeyEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onKeyEvent.addListener(A.H.get<object>(callback));
    },
    "try_OnKeyEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onKeyEvent.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffKeyEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onKeyEvent && "removeListener" in WEBEXT?.input?.ime?.onKeyEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffKeyEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onKeyEvent.removeListener);
    },
    "call_OffKeyEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onKeyEvent.removeListener(A.H.get<object>(callback));
    },
    "try_OffKeyEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onKeyEvent.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnKeyEvent": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onKeyEvent && "hasListener" in WEBEXT?.input?.ime?.onKeyEvent) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnKeyEvent": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onKeyEvent.hasListener);
    },
    "call_HasOnKeyEvent": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onKeyEvent.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnKeyEvent": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onKeyEvent.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMenuItemActivated": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onMenuItemActivated && "addListener" in WEBEXT?.input?.ime?.onMenuItemActivated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMenuItemActivated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onMenuItemActivated.addListener);
    },
    "call_OnMenuItemActivated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onMenuItemActivated.addListener(A.H.get<object>(callback));
    },
    "try_OnMenuItemActivated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onMenuItemActivated.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMenuItemActivated": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onMenuItemActivated && "removeListener" in WEBEXT?.input?.ime?.onMenuItemActivated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMenuItemActivated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onMenuItemActivated.removeListener);
    },
    "call_OffMenuItemActivated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onMenuItemActivated.removeListener(A.H.get<object>(callback));
    },
    "try_OffMenuItemActivated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onMenuItemActivated.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMenuItemActivated": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onMenuItemActivated && "hasListener" in WEBEXT?.input?.ime?.onMenuItemActivated) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMenuItemActivated": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onMenuItemActivated.hasListener);
    },
    "call_HasOnMenuItemActivated": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onMenuItemActivated.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMenuItemActivated": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onMenuItemActivated.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnReset": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onReset && "addListener" in WEBEXT?.input?.ime?.onReset) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnReset": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onReset.addListener);
    },
    "call_OnReset": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onReset.addListener(A.H.get<object>(callback));
    },
    "try_OnReset": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onReset.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffReset": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onReset && "removeListener" in WEBEXT?.input?.ime?.onReset) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffReset": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onReset.removeListener);
    },
    "call_OffReset": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onReset.removeListener(A.H.get<object>(callback));
    },
    "try_OffReset": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onReset.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnReset": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime?.onReset && "hasListener" in WEBEXT?.input?.ime?.onReset) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnReset": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onReset.hasListener);
    },
    "call_HasOnReset": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onReset.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnReset": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onReset.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnSurroundingTextChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.input?.ime?.onSurroundingTextChanged &&
        "addListener" in WEBEXT?.input?.ime?.onSurroundingTextChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnSurroundingTextChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onSurroundingTextChanged.addListener);
    },
    "call_OnSurroundingTextChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onSurroundingTextChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnSurroundingTextChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onSurroundingTextChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffSurroundingTextChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.input?.ime?.onSurroundingTextChanged &&
        "removeListener" in WEBEXT?.input?.ime?.onSurroundingTextChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffSurroundingTextChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onSurroundingTextChanged.removeListener);
    },
    "call_OffSurroundingTextChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onSurroundingTextChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffSurroundingTextChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onSurroundingTextChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnSurroundingTextChanged": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.input?.ime?.onSurroundingTextChanged &&
        "hasListener" in WEBEXT?.input?.ime?.onSurroundingTextChanged
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnSurroundingTextChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.onSurroundingTextChanged.hasListener);
    },
    "call_HasOnSurroundingTextChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.input.ime.onSurroundingTextChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnSurroundingTextChanged": (
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.input.ime.onSurroundingTextChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendKeyEvents": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime && "sendKeyEvents" in WEBEXT?.input?.ime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendKeyEvents": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.sendKeyEvents);
    },
    "call_SendKeyEvents": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["contextID"] = A.load.Int64(parameters + 0);
      parameters_ffi["keyData"] = A.load.Ref(parameters + 8, undefined);

      const _ret = WEBEXT.input.ime.sendKeyEvents(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SendKeyEvents": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["contextID"] = A.load.Int64(parameters + 0);
        parameters_ffi["keyData"] = A.load.Ref(parameters + 8, undefined);

        const _ret = WEBEXT.input.ime.sendKeyEvents(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetAssistiveWindowButtonHighlighted": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime && "setAssistiveWindowButtonHighlighted" in WEBEXT?.input?.ime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetAssistiveWindowButtonHighlighted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.setAssistiveWindowButtonHighlighted);
    },
    "call_SetAssistiveWindowButtonHighlighted": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["announceString"] = A.load.Ref(parameters + 0, undefined);
      parameters_ffi["buttonID"] = A.load.Enum(parameters + 4, ["undo", "addToDictionary"]);
      parameters_ffi["contextID"] = A.load.Int64(parameters + 8);
      parameters_ffi["highlighted"] = A.load.Bool(parameters + 16);
      parameters_ffi["windowType"] = A.load.Enum(parameters + 20, ["undo"]);

      const _ret = WEBEXT.input.ime.setAssistiveWindowButtonHighlighted(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetAssistiveWindowButtonHighlighted": (
      retPtr: Pointer,
      errPtr: Pointer,
      parameters: Pointer
    ): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["announceString"] = A.load.Ref(parameters + 0, undefined);
        parameters_ffi["buttonID"] = A.load.Enum(parameters + 4, ["undo", "addToDictionary"]);
        parameters_ffi["contextID"] = A.load.Int64(parameters + 8);
        parameters_ffi["highlighted"] = A.load.Bool(parameters + 16);
        parameters_ffi["windowType"] = A.load.Enum(parameters + 20, ["undo"]);

        const _ret = WEBEXT.input.ime.setAssistiveWindowButtonHighlighted(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetAssistiveWindowProperties": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime && "setAssistiveWindowProperties" in WEBEXT?.input?.ime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetAssistiveWindowProperties": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.setAssistiveWindowProperties);
    },
    "call_SetAssistiveWindowProperties": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["contextID"] = A.load.Int64(parameters + 0);
      if (A.load.Bool(parameters + 8 + 9)) {
        parameters_ffi["properties"] = {};
        parameters_ffi["properties"]["announceString"] = A.load.Ref(parameters + 8 + 0, undefined);
        parameters_ffi["properties"]["type"] = A.load.Enum(parameters + 8 + 4, ["undo"]);
        parameters_ffi["properties"]["visible"] = A.load.Bool(parameters + 8 + 8);
      }

      const _ret = WEBEXT.input.ime.setAssistiveWindowProperties(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetAssistiveWindowProperties": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["contextID"] = A.load.Int64(parameters + 0);
        if (A.load.Bool(parameters + 8 + 9)) {
          parameters_ffi["properties"] = {};
          parameters_ffi["properties"]["announceString"] = A.load.Ref(parameters + 8 + 0, undefined);
          parameters_ffi["properties"]["type"] = A.load.Enum(parameters + 8 + 4, ["undo"]);
          parameters_ffi["properties"]["visible"] = A.load.Bool(parameters + 8 + 8);
        }

        const _ret = WEBEXT.input.ime.setAssistiveWindowProperties(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetCandidateWindowProperties": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime && "setCandidateWindowProperties" in WEBEXT?.input?.ime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetCandidateWindowProperties": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.setCandidateWindowProperties);
    },
    "call_SetCandidateWindowProperties": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["engineID"] = A.load.Ref(parameters + 0, undefined);
      if (A.load.Bool(parameters + 8 + 55)) {
        parameters_ffi["properties"] = {};
        parameters_ffi["properties"]["auxiliaryText"] = A.load.Ref(parameters + 8 + 0, undefined);
        if (A.load.Bool(parameters + 8 + 48)) {
          parameters_ffi["properties"]["auxiliaryTextVisible"] = A.load.Bool(parameters + 8 + 4);
        }
        if (A.load.Bool(parameters + 8 + 49)) {
          parameters_ffi["properties"]["currentCandidateIndex"] = A.load.Int64(parameters + 8 + 8);
        }
        if (A.load.Bool(parameters + 8 + 50)) {
          parameters_ffi["properties"]["cursorVisible"] = A.load.Bool(parameters + 8 + 16);
        }
        if (A.load.Bool(parameters + 8 + 51)) {
          parameters_ffi["properties"]["pageSize"] = A.load.Int64(parameters + 8 + 24);
        }
        if (A.load.Bool(parameters + 8 + 52)) {
          parameters_ffi["properties"]["totalCandidates"] = A.load.Int64(parameters + 8 + 32);
        }
        if (A.load.Bool(parameters + 8 + 53)) {
          parameters_ffi["properties"]["vertical"] = A.load.Bool(parameters + 8 + 40);
        }
        if (A.load.Bool(parameters + 8 + 54)) {
          parameters_ffi["properties"]["visible"] = A.load.Bool(parameters + 8 + 41);
        }
        parameters_ffi["properties"]["windowPosition"] = A.load.Enum(parameters + 8 + 44, ["cursor", "composition"]);
      }

      const _ret = WEBEXT.input.ime.setCandidateWindowProperties(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetCandidateWindowProperties": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["engineID"] = A.load.Ref(parameters + 0, undefined);
        if (A.load.Bool(parameters + 8 + 55)) {
          parameters_ffi["properties"] = {};
          parameters_ffi["properties"]["auxiliaryText"] = A.load.Ref(parameters + 8 + 0, undefined);
          if (A.load.Bool(parameters + 8 + 48)) {
            parameters_ffi["properties"]["auxiliaryTextVisible"] = A.load.Bool(parameters + 8 + 4);
          }
          if (A.load.Bool(parameters + 8 + 49)) {
            parameters_ffi["properties"]["currentCandidateIndex"] = A.load.Int64(parameters + 8 + 8);
          }
          if (A.load.Bool(parameters + 8 + 50)) {
            parameters_ffi["properties"]["cursorVisible"] = A.load.Bool(parameters + 8 + 16);
          }
          if (A.load.Bool(parameters + 8 + 51)) {
            parameters_ffi["properties"]["pageSize"] = A.load.Int64(parameters + 8 + 24);
          }
          if (A.load.Bool(parameters + 8 + 52)) {
            parameters_ffi["properties"]["totalCandidates"] = A.load.Int64(parameters + 8 + 32);
          }
          if (A.load.Bool(parameters + 8 + 53)) {
            parameters_ffi["properties"]["vertical"] = A.load.Bool(parameters + 8 + 40);
          }
          if (A.load.Bool(parameters + 8 + 54)) {
            parameters_ffi["properties"]["visible"] = A.load.Bool(parameters + 8 + 41);
          }
          parameters_ffi["properties"]["windowPosition"] = A.load.Enum(parameters + 8 + 44, ["cursor", "composition"]);
        }

        const _ret = WEBEXT.input.ime.setCandidateWindowProperties(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetCandidates": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime && "setCandidates" in WEBEXT?.input?.ime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetCandidates": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.setCandidates);
    },
    "call_SetCandidates": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["candidates"] = A.load.Ref(parameters + 0, undefined);
      parameters_ffi["contextID"] = A.load.Int64(parameters + 8);

      const _ret = WEBEXT.input.ime.setCandidates(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetCandidates": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["candidates"] = A.load.Ref(parameters + 0, undefined);
        parameters_ffi["contextID"] = A.load.Int64(parameters + 8);

        const _ret = WEBEXT.input.ime.setCandidates(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetComposition": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime && "setComposition" in WEBEXT?.input?.ime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetComposition": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.setComposition);
    },
    "call_SetComposition": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["contextID"] = A.load.Int64(parameters + 0);
      parameters_ffi["cursor"] = A.load.Int64(parameters + 8);
      parameters_ffi["segments"] = A.load.Ref(parameters + 16, undefined);
      if (A.load.Bool(parameters + 44)) {
        parameters_ffi["selectionEnd"] = A.load.Int64(parameters + 24);
      }
      if (A.load.Bool(parameters + 45)) {
        parameters_ffi["selectionStart"] = A.load.Int64(parameters + 32);
      }
      parameters_ffi["text"] = A.load.Ref(parameters + 40, undefined);

      const _ret = WEBEXT.input.ime.setComposition(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetComposition": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["contextID"] = A.load.Int64(parameters + 0);
        parameters_ffi["cursor"] = A.load.Int64(parameters + 8);
        parameters_ffi["segments"] = A.load.Ref(parameters + 16, undefined);
        if (A.load.Bool(parameters + 44)) {
          parameters_ffi["selectionEnd"] = A.load.Int64(parameters + 24);
        }
        if (A.load.Bool(parameters + 45)) {
          parameters_ffi["selectionStart"] = A.load.Int64(parameters + 32);
        }
        parameters_ffi["text"] = A.load.Ref(parameters + 40, undefined);

        const _ret = WEBEXT.input.ime.setComposition(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetCursorPosition": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime && "setCursorPosition" in WEBEXT?.input?.ime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetCursorPosition": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.setCursorPosition);
    },
    "call_SetCursorPosition": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["candidateID"] = A.load.Int64(parameters + 0);
      parameters_ffi["contextID"] = A.load.Int64(parameters + 8);

      const _ret = WEBEXT.input.ime.setCursorPosition(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetCursorPosition": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["candidateID"] = A.load.Int64(parameters + 0);
        parameters_ffi["contextID"] = A.load.Int64(parameters + 8);

        const _ret = WEBEXT.input.ime.setCursorPosition(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetMenuItems": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime && "setMenuItems" in WEBEXT?.input?.ime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetMenuItems": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.setMenuItems);
    },
    "call_SetMenuItems": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["engineID"] = A.load.Ref(parameters + 0, undefined);
      parameters_ffi["items"] = A.load.Ref(parameters + 4, undefined);

      const _ret = WEBEXT.input.ime.setMenuItems(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetMenuItems": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["engineID"] = A.load.Ref(parameters + 0, undefined);
        parameters_ffi["items"] = A.load.Ref(parameters + 4, undefined);

        const _ret = WEBEXT.input.ime.setMenuItems(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_UpdateMenuItems": (): heap.Ref<boolean> => {
      if (WEBEXT?.input?.ime && "updateMenuItems" in WEBEXT?.input?.ime) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_UpdateMenuItems": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.input.ime.updateMenuItems);
    },
    "call_UpdateMenuItems": (retPtr: Pointer, parameters: Pointer): void => {
      const parameters_ffi = {};

      parameters_ffi["engineID"] = A.load.Ref(parameters + 0, undefined);
      parameters_ffi["items"] = A.load.Ref(parameters + 4, undefined);

      const _ret = WEBEXT.input.ime.updateMenuItems(parameters_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_UpdateMenuItems": (retPtr: Pointer, errPtr: Pointer, parameters: Pointer): heap.Ref<boolean> => {
      try {
        const parameters_ffi = {};

        parameters_ffi["engineID"] = A.load.Ref(parameters + 0, undefined);
        parameters_ffi["items"] = A.load.Ref(parameters + 4, undefined);

        const _ret = WEBEXT.input.ime.updateMenuItems(parameters_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
