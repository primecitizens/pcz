import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/omnibox", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "constof_DescriptionStyleType": (ref: heap.Ref<string>): number => {
      const idx = ["url", "match", "dim"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_MatchClassification": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 21, false);
        A.store.Bool(ptr + 20, false);
        A.store.Int64(ptr + 0, 0);
        A.store.Int64(ptr + 8, 0);
        A.store.Enum(ptr + 16, -1);
      } else {
        A.store.Bool(ptr + 21, true);
        A.store.Bool(ptr + 20, "length" in x ? true : false);
        A.store.Int64(ptr + 0, x["length"] === undefined ? 0 : (x["length"] as number));
        A.store.Int64(ptr + 8, x["offset"] === undefined ? 0 : (x["offset"] as number));
        A.store.Enum(ptr + 16, ["url", "match", "dim"].indexOf(x["type"] as string));
      }
    },
    "load_MatchClassification": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 20)) {
        x["length"] = A.load.Int64(ptr + 0);
      } else {
        delete x["length"];
      }
      x["offset"] = A.load.Int64(ptr + 8);
      x["type"] = A.load.Enum(ptr + 16, ["url", "match", "dim"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_DefaultSuggestResult": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["description"]);
        A.store.Ref(ptr + 4, x["descriptionStyles"]);
      }
    },
    "load_DefaultSuggestResult": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["description"] = A.load.Ref(ptr + 0, undefined);
      x["descriptionStyles"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_SuggestResult": (ptr: Pointer, ref: heap.Ref<any>) => {
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
        A.store.Ref(ptr + 0, x["content"]);
        A.store.Bool(ptr + 16, "deletable" in x ? true : false);
        A.store.Bool(ptr + 4, x["deletable"] ? true : false);
        A.store.Ref(ptr + 8, x["description"]);
        A.store.Ref(ptr + 12, x["descriptionStyles"]);
      }
    },
    "load_SuggestResult": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["content"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["deletable"] = A.load.Bool(ptr + 4);
      } else {
        delete x["deletable"];
      }
      x["description"] = A.load.Ref(ptr + 8, undefined);
      x["descriptionStyles"] = A.load.Ref(ptr + 12, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OnInputEnteredDisposition": (ref: heap.Ref<string>): number => {
      const idx = ["currentTab", "newForegroundTab", "newBackgroundTab"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_OnDeleteSuggestion": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox?.onDeleteSuggestion && "addListener" in WEBEXT?.omnibox?.onDeleteSuggestion) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeleteSuggestion": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.onDeleteSuggestion.addListener);
    },
    "call_OnDeleteSuggestion": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.onDeleteSuggestion.addListener(A.H.get<object>(callback));
    },
    "try_OnDeleteSuggestion": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.onDeleteSuggestion.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeleteSuggestion": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox?.onDeleteSuggestion && "removeListener" in WEBEXT?.omnibox?.onDeleteSuggestion) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeleteSuggestion": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.onDeleteSuggestion.removeListener);
    },
    "call_OffDeleteSuggestion": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.onDeleteSuggestion.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeleteSuggestion": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.onDeleteSuggestion.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeleteSuggestion": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox?.onDeleteSuggestion && "hasListener" in WEBEXT?.omnibox?.onDeleteSuggestion) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeleteSuggestion": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.onDeleteSuggestion.hasListener);
    },
    "call_HasOnDeleteSuggestion": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.onDeleteSuggestion.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeleteSuggestion": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.onDeleteSuggestion.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnInputCancelled": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox?.onInputCancelled && "addListener" in WEBEXT?.omnibox?.onInputCancelled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnInputCancelled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.onInputCancelled.addListener);
    },
    "call_OnInputCancelled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.onInputCancelled.addListener(A.H.get<object>(callback));
    },
    "try_OnInputCancelled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.onInputCancelled.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffInputCancelled": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox?.onInputCancelled && "removeListener" in WEBEXT?.omnibox?.onInputCancelled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffInputCancelled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.onInputCancelled.removeListener);
    },
    "call_OffInputCancelled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.onInputCancelled.removeListener(A.H.get<object>(callback));
    },
    "try_OffInputCancelled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.onInputCancelled.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnInputCancelled": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox?.onInputCancelled && "hasListener" in WEBEXT?.omnibox?.onInputCancelled) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnInputCancelled": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.onInputCancelled.hasListener);
    },
    "call_HasOnInputCancelled": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.onInputCancelled.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnInputCancelled": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.onInputCancelled.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnInputChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox?.onInputChanged && "addListener" in WEBEXT?.omnibox?.onInputChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnInputChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.onInputChanged.addListener);
    },
    "call_OnInputChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.onInputChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnInputChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.onInputChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffInputChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox?.onInputChanged && "removeListener" in WEBEXT?.omnibox?.onInputChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffInputChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.onInputChanged.removeListener);
    },
    "call_OffInputChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.onInputChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffInputChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.onInputChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnInputChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox?.onInputChanged && "hasListener" in WEBEXT?.omnibox?.onInputChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnInputChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.onInputChanged.hasListener);
    },
    "call_HasOnInputChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.onInputChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnInputChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.onInputChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnInputEntered": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox?.onInputEntered && "addListener" in WEBEXT?.omnibox?.onInputEntered) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnInputEntered": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.onInputEntered.addListener);
    },
    "call_OnInputEntered": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.onInputEntered.addListener(A.H.get<object>(callback));
    },
    "try_OnInputEntered": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.onInputEntered.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffInputEntered": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox?.onInputEntered && "removeListener" in WEBEXT?.omnibox?.onInputEntered) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffInputEntered": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.onInputEntered.removeListener);
    },
    "call_OffInputEntered": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.onInputEntered.removeListener(A.H.get<object>(callback));
    },
    "try_OffInputEntered": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.onInputEntered.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnInputEntered": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox?.onInputEntered && "hasListener" in WEBEXT?.omnibox?.onInputEntered) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnInputEntered": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.onInputEntered.hasListener);
    },
    "call_HasOnInputEntered": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.onInputEntered.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnInputEntered": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.onInputEntered.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnInputStarted": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox?.onInputStarted && "addListener" in WEBEXT?.omnibox?.onInputStarted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnInputStarted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.onInputStarted.addListener);
    },
    "call_OnInputStarted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.onInputStarted.addListener(A.H.get<object>(callback));
    },
    "try_OnInputStarted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.onInputStarted.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffInputStarted": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox?.onInputStarted && "removeListener" in WEBEXT?.omnibox?.onInputStarted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffInputStarted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.onInputStarted.removeListener);
    },
    "call_OffInputStarted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.onInputStarted.removeListener(A.H.get<object>(callback));
    },
    "try_OffInputStarted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.onInputStarted.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnInputStarted": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox?.onInputStarted && "hasListener" in WEBEXT?.omnibox?.onInputStarted) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnInputStarted": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.onInputStarted.hasListener);
    },
    "call_HasOnInputStarted": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.onInputStarted.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnInputStarted": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.onInputStarted.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendSuggestions": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox && "sendSuggestions" in WEBEXT?.omnibox) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendSuggestions": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.sendSuggestions);
    },
    "call_SendSuggestions": (retPtr: Pointer, requestId: number, suggestResults: heap.Ref<object>): void => {
      const _ret = WEBEXT.omnibox.sendSuggestions(requestId, A.H.get<object>(suggestResults));
    },
    "try_SendSuggestions": (
      retPtr: Pointer,
      errPtr: Pointer,
      requestId: number,
      suggestResults: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.omnibox.sendSuggestions(requestId, A.H.get<object>(suggestResults));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetDefaultSuggestion": (): heap.Ref<boolean> => {
      if (WEBEXT?.omnibox && "setDefaultSuggestion" in WEBEXT?.omnibox) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetDefaultSuggestion": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.omnibox.setDefaultSuggestion);
    },
    "call_SetDefaultSuggestion": (retPtr: Pointer, suggestion: Pointer): void => {
      const suggestion_ffi = {};

      suggestion_ffi["description"] = A.load.Ref(suggestion + 0, undefined);
      suggestion_ffi["descriptionStyles"] = A.load.Ref(suggestion + 4, undefined);

      const _ret = WEBEXT.omnibox.setDefaultSuggestion(suggestion_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetDefaultSuggestion": (retPtr: Pointer, errPtr: Pointer, suggestion: Pointer): heap.Ref<boolean> => {
      try {
        const suggestion_ffi = {};

        suggestion_ffi["description"] = A.load.Ref(suggestion + 0, undefined);
        suggestion_ffi["descriptionStyles"] = A.load.Ref(suggestion + 4, undefined);

        const _ret = WEBEXT.omnibox.setDefaultSuggestion(suggestion_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
