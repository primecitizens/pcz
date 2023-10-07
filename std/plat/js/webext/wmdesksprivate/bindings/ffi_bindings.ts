import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/wmdesksprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Desk": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["deskUuid"]);
        A.store.Ref(ptr + 4, x["deskName"]);
      }
    },
    "load_Desk": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["deskUuid"] = A.load.Ref(ptr + 0, undefined);
      x["deskName"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_SavedDeskType": (ref: heap.Ref<string>): number => {
      const idx = ["kTemplate", "kSaveAndRecall", "kUnknown"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SavedDesk": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Ref(ptr + 0, x["savedDeskUuid"]);
        A.store.Ref(ptr + 4, x["savedDeskName"]);
        A.store.Enum(ptr + 8, ["kTemplate", "kSaveAndRecall", "kUnknown"].indexOf(x["savedDeskType"] as string));
      }
    },
    "load_SavedDesk": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["savedDeskUuid"] = A.load.Ref(ptr + 0, undefined);
      x["savedDeskName"] = A.load.Ref(ptr + 4, undefined);
      x["savedDeskType"] = A.load.Enum(ptr + 8, ["kTemplate", "kSaveAndRecall", "kUnknown"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_LaunchOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Ref(ptr + 0, undefined);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Ref(ptr + 0, x["deskName"]);
      }
    },
    "load_LaunchOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["deskName"] = A.load.Ref(ptr + 0, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_OnDeskAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate?.OnDeskAdded && "addListener" in WEBEXT?.wmDesksPrivate?.OnDeskAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeskAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.OnDeskAdded.addListener);
    },
    "call_OnDeskAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.wmDesksPrivate.OnDeskAdded.addListener(A.H.get<object>(callback));
    },
    "try_OnDeskAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.OnDeskAdded.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeskAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate?.OnDeskAdded && "removeListener" in WEBEXT?.wmDesksPrivate?.OnDeskAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeskAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.OnDeskAdded.removeListener);
    },
    "call_OffDeskAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.wmDesksPrivate.OnDeskAdded.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeskAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.OnDeskAdded.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeskAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate?.OnDeskAdded && "hasListener" in WEBEXT?.wmDesksPrivate?.OnDeskAdded) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeskAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.OnDeskAdded.hasListener);
    },
    "call_HasOnDeskAdded": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.wmDesksPrivate.OnDeskAdded.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeskAdded": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.OnDeskAdded.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeskRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate?.OnDeskRemoved && "addListener" in WEBEXT?.wmDesksPrivate?.OnDeskRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeskRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.OnDeskRemoved.addListener);
    },
    "call_OnDeskRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.wmDesksPrivate.OnDeskRemoved.addListener(A.H.get<object>(callback));
    },
    "try_OnDeskRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.OnDeskRemoved.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeskRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate?.OnDeskRemoved && "removeListener" in WEBEXT?.wmDesksPrivate?.OnDeskRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeskRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.OnDeskRemoved.removeListener);
    },
    "call_OffDeskRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.wmDesksPrivate.OnDeskRemoved.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeskRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.OnDeskRemoved.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeskRemoved": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate?.OnDeskRemoved && "hasListener" in WEBEXT?.wmDesksPrivate?.OnDeskRemoved) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeskRemoved": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.OnDeskRemoved.hasListener);
    },
    "call_HasOnDeskRemoved": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.wmDesksPrivate.OnDeskRemoved.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeskRemoved": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.OnDeskRemoved.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDeskSwitched": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate?.OnDeskSwitched && "addListener" in WEBEXT?.wmDesksPrivate?.OnDeskSwitched) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDeskSwitched": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.OnDeskSwitched.addListener);
    },
    "call_OnDeskSwitched": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.wmDesksPrivate.OnDeskSwitched.addListener(A.H.get<object>(callback));
    },
    "try_OnDeskSwitched": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.OnDeskSwitched.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDeskSwitched": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate?.OnDeskSwitched && "removeListener" in WEBEXT?.wmDesksPrivate?.OnDeskSwitched) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDeskSwitched": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.OnDeskSwitched.removeListener);
    },
    "call_OffDeskSwitched": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.wmDesksPrivate.OnDeskSwitched.removeListener(A.H.get<object>(callback));
    },
    "try_OffDeskSwitched": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.OnDeskSwitched.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDeskSwitched": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate?.OnDeskSwitched && "hasListener" in WEBEXT?.wmDesksPrivate?.OnDeskSwitched) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDeskSwitched": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.OnDeskSwitched.hasListener);
    },
    "call_HasOnDeskSwitched": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.wmDesksPrivate.OnDeskSwitched.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDeskSwitched": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.OnDeskSwitched.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },

    "store_RemoveDeskOptions": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 1, false);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Bool(ptr + 2, "combineDesks" in x ? true : false);
        A.store.Bool(ptr + 0, x["combineDesks"] ? true : false);
        A.store.Bool(ptr + 3, "allowUndo" in x ? true : false);
        A.store.Bool(ptr + 1, x["allowUndo"] ? true : false);
      }
    },
    "load_RemoveDeskOptions": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 2)) {
        x["combineDesks"] = A.load.Bool(ptr + 0);
      } else {
        delete x["combineDesks"];
      }
      if (A.load.Bool(ptr + 3)) {
        x["allowUndo"] = A.load.Bool(ptr + 1);
      } else {
        delete x["allowUndo"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_WindowProperties": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 1, false);
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 1, "allDesks" in x ? true : false);
        A.store.Bool(ptr + 0, x["allDesks"] ? true : false);
      }
    },
    "load_WindowProperties": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 1)) {
        x["allDesks"] = A.load.Bool(ptr + 0);
      } else {
        delete x["allDesks"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_DeleteSavedDesk": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate && "deleteSavedDesk" in WEBEXT?.wmDesksPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_DeleteSavedDesk": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.deleteSavedDesk);
    },
    "call_DeleteSavedDesk": (retPtr: Pointer, savedDeskUuid: heap.Ref<object>): void => {
      const _ret = WEBEXT.wmDesksPrivate.deleteSavedDesk(A.H.get<object>(savedDeskUuid));
      A.store.Ref(retPtr, _ret);
    },
    "try_DeleteSavedDesk": (retPtr: Pointer, errPtr: Pointer, savedDeskUuid: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.deleteSavedDesk(A.H.get<object>(savedDeskUuid));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetActiveDesk": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate && "getActiveDesk" in WEBEXT?.wmDesksPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetActiveDesk": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.getActiveDesk);
    },
    "call_GetActiveDesk": (retPtr: Pointer): void => {
      const _ret = WEBEXT.wmDesksPrivate.getActiveDesk();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetActiveDesk": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.getActiveDesk();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAllDesks": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate && "getAllDesks" in WEBEXT?.wmDesksPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAllDesks": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.getAllDesks);
    },
    "call_GetAllDesks": (retPtr: Pointer): void => {
      const _ret = WEBEXT.wmDesksPrivate.getAllDesks();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetAllDesks": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.getAllDesks();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDeskByID": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate && "getDeskByID" in WEBEXT?.wmDesksPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDeskByID": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.getDeskByID);
    },
    "call_GetDeskByID": (retPtr: Pointer, deskUuid: heap.Ref<object>): void => {
      const _ret = WEBEXT.wmDesksPrivate.getDeskByID(A.H.get<object>(deskUuid));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDeskByID": (retPtr: Pointer, errPtr: Pointer, deskUuid: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.getDeskByID(A.H.get<object>(deskUuid));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDeskTemplateJson": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate && "getDeskTemplateJson" in WEBEXT?.wmDesksPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDeskTemplateJson": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.getDeskTemplateJson);
    },
    "call_GetDeskTemplateJson": (retPtr: Pointer, templateUuid: heap.Ref<object>): void => {
      const _ret = WEBEXT.wmDesksPrivate.getDeskTemplateJson(A.H.get<object>(templateUuid));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetDeskTemplateJson": (
      retPtr: Pointer,
      errPtr: Pointer,
      templateUuid: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.getDeskTemplateJson(A.H.get<object>(templateUuid));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSavedDesks": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate && "getSavedDesks" in WEBEXT?.wmDesksPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSavedDesks": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.getSavedDesks);
    },
    "call_GetSavedDesks": (retPtr: Pointer): void => {
      const _ret = WEBEXT.wmDesksPrivate.getSavedDesks();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSavedDesks": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.getSavedDesks();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LaunchDesk": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate && "launchDesk" in WEBEXT?.wmDesksPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LaunchDesk": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.launchDesk);
    },
    "call_LaunchDesk": (retPtr: Pointer, launchOptions: Pointer): void => {
      const launchOptions_ffi = {};

      launchOptions_ffi["deskName"] = A.load.Ref(launchOptions + 0, undefined);

      const _ret = WEBEXT.wmDesksPrivate.launchDesk(launchOptions_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_LaunchDesk": (retPtr: Pointer, errPtr: Pointer, launchOptions: Pointer): heap.Ref<boolean> => {
      try {
        const launchOptions_ffi = {};

        launchOptions_ffi["deskName"] = A.load.Ref(launchOptions + 0, undefined);

        const _ret = WEBEXT.wmDesksPrivate.launchDesk(launchOptions_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RecallSavedDesk": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate && "recallSavedDesk" in WEBEXT?.wmDesksPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RecallSavedDesk": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.recallSavedDesk);
    },
    "call_RecallSavedDesk": (retPtr: Pointer, savedDeskUuid: heap.Ref<object>): void => {
      const _ret = WEBEXT.wmDesksPrivate.recallSavedDesk(A.H.get<object>(savedDeskUuid));
      A.store.Ref(retPtr, _ret);
    },
    "try_RecallSavedDesk": (retPtr: Pointer, errPtr: Pointer, savedDeskUuid: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.recallSavedDesk(A.H.get<object>(savedDeskUuid));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveDesk": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate && "removeDesk" in WEBEXT?.wmDesksPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveDesk": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.removeDesk);
    },
    "call_RemoveDesk": (retPtr: Pointer, deskId: heap.Ref<object>, removeDeskOptions: Pointer): void => {
      const removeDeskOptions_ffi = {};

      if (A.load.Bool(removeDeskOptions + 2)) {
        removeDeskOptions_ffi["combineDesks"] = A.load.Bool(removeDeskOptions + 0);
      }
      if (A.load.Bool(removeDeskOptions + 3)) {
        removeDeskOptions_ffi["allowUndo"] = A.load.Bool(removeDeskOptions + 1);
      }

      const _ret = WEBEXT.wmDesksPrivate.removeDesk(A.H.get<object>(deskId), removeDeskOptions_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveDesk": (
      retPtr: Pointer,
      errPtr: Pointer,
      deskId: heap.Ref<object>,
      removeDeskOptions: Pointer
    ): heap.Ref<boolean> => {
      try {
        const removeDeskOptions_ffi = {};

        if (A.load.Bool(removeDeskOptions + 2)) {
          removeDeskOptions_ffi["combineDesks"] = A.load.Bool(removeDeskOptions + 0);
        }
        if (A.load.Bool(removeDeskOptions + 3)) {
          removeDeskOptions_ffi["allowUndo"] = A.load.Bool(removeDeskOptions + 1);
        }

        const _ret = WEBEXT.wmDesksPrivate.removeDesk(A.H.get<object>(deskId), removeDeskOptions_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SaveActiveDesk": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate && "saveActiveDesk" in WEBEXT?.wmDesksPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SaveActiveDesk": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.saveActiveDesk);
    },
    "call_SaveActiveDesk": (retPtr: Pointer): void => {
      const _ret = WEBEXT.wmDesksPrivate.saveActiveDesk();
      A.store.Ref(retPtr, _ret);
    },
    "try_SaveActiveDesk": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.saveActiveDesk();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetWindowProperties": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate && "setWindowProperties" in WEBEXT?.wmDesksPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetWindowProperties": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.setWindowProperties);
    },
    "call_SetWindowProperties": (retPtr: Pointer, windowId: number, windowProperties: Pointer): void => {
      const windowProperties_ffi = {};

      if (A.load.Bool(windowProperties + 1)) {
        windowProperties_ffi["allDesks"] = A.load.Bool(windowProperties + 0);
      }

      const _ret = WEBEXT.wmDesksPrivate.setWindowProperties(windowId, windowProperties_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_SetWindowProperties": (
      retPtr: Pointer,
      errPtr: Pointer,
      windowId: number,
      windowProperties: Pointer
    ): heap.Ref<boolean> => {
      try {
        const windowProperties_ffi = {};

        if (A.load.Bool(windowProperties + 1)) {
          windowProperties_ffi["allDesks"] = A.load.Bool(windowProperties + 0);
        }

        const _ret = WEBEXT.wmDesksPrivate.setWindowProperties(windowId, windowProperties_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SwitchDesk": (): heap.Ref<boolean> => {
      if (WEBEXT?.wmDesksPrivate && "switchDesk" in WEBEXT?.wmDesksPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SwitchDesk": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.wmDesksPrivate.switchDesk);
    },
    "call_SwitchDesk": (retPtr: Pointer, deskUuid: heap.Ref<object>): void => {
      const _ret = WEBEXT.wmDesksPrivate.switchDesk(A.H.get<object>(deskUuid));
      A.store.Ref(retPtr, _ret);
    },
    "try_SwitchDesk": (retPtr: Pointer, errPtr: Pointer, deskUuid: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.wmDesksPrivate.switchDesk(A.H.get<object>(deskUuid));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
