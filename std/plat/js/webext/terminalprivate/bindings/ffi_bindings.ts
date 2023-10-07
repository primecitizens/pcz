import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/terminalprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_GetOSInfoReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 1, false);
      } else {
        A.store.Bool(ptr + 2, true);
        A.store.Bool(ptr + 0, x["alternative_emulator"] ? true : false);
        A.store.Bool(ptr + 1, x["tast"] ? true : false);
      }
    },
    "load_GetOSInfoReturnType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["alternative_emulator"] = A.load.Bool(ptr + 0);
      x["tast"] = A.load.Bool(ptr + 1);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OpenWindowArgData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Bool(ptr + 0, false);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "asTab" in x ? true : false);
        A.store.Bool(ptr + 0, x["asTab"] ? true : false);
        A.store.Ref(ptr + 4, x["url"]);
      }
    },
    "load_OpenWindowArgData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["asTab"] = A.load.Bool(ptr + 0);
      } else {
        delete x["asTab"];
      }
      x["url"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_OutputType": (ref: heap.Ref<string>): number => {
      const idx = ["stdout", "stderr", "exit"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "has_AckOutput": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate && "ackOutput" in WEBEXT?.terminalPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AckOutput": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.ackOutput);
    },
    "call_AckOutput": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.terminalPrivate.ackOutput(A.H.get<object>(id));
    },
    "try_AckOutput": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.ackOutput(A.H.get<object>(id));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CloseTerminalProcess": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate && "closeTerminalProcess" in WEBEXT?.terminalPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CloseTerminalProcess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.closeTerminalProcess);
    },
    "call_CloseTerminalProcess": (retPtr: Pointer, id: heap.Ref<object>): void => {
      const _ret = WEBEXT.terminalPrivate.closeTerminalProcess(A.H.get<object>(id));
      A.store.Ref(retPtr, _ret);
    },
    "try_CloseTerminalProcess": (retPtr: Pointer, errPtr: Pointer, id: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.closeTerminalProcess(A.H.get<object>(id));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetOSInfo": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate && "getOSInfo" in WEBEXT?.terminalPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetOSInfo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.getOSInfo);
    },
    "call_GetOSInfo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.terminalPrivate.getOSInfo();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetOSInfo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.getOSInfo();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetPrefs": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate && "getPrefs" in WEBEXT?.terminalPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetPrefs": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.getPrefs);
    },
    "call_GetPrefs": (retPtr: Pointer, paths: heap.Ref<object>): void => {
      const _ret = WEBEXT.terminalPrivate.getPrefs(A.H.get<object>(paths));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetPrefs": (retPtr: Pointer, errPtr: Pointer, paths: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.getPrefs(A.H.get<object>(paths));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnPrefChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate?.onPrefChanged && "addListener" in WEBEXT?.terminalPrivate?.onPrefChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnPrefChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.onPrefChanged.addListener);
    },
    "call_OnPrefChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.terminalPrivate.onPrefChanged.addListener(A.H.get<object>(callback));
    },
    "try_OnPrefChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.onPrefChanged.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffPrefChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate?.onPrefChanged && "removeListener" in WEBEXT?.terminalPrivate?.onPrefChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffPrefChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.onPrefChanged.removeListener);
    },
    "call_OffPrefChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.terminalPrivate.onPrefChanged.removeListener(A.H.get<object>(callback));
    },
    "try_OffPrefChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.onPrefChanged.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnPrefChanged": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate?.onPrefChanged && "hasListener" in WEBEXT?.terminalPrivate?.onPrefChanged) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnPrefChanged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.onPrefChanged.hasListener);
    },
    "call_HasOnPrefChanged": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.terminalPrivate.onPrefChanged.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnPrefChanged": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.onPrefChanged.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnProcessOutput": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate?.onProcessOutput && "addListener" in WEBEXT?.terminalPrivate?.onProcessOutput) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnProcessOutput": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.onProcessOutput.addListener);
    },
    "call_OnProcessOutput": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.terminalPrivate.onProcessOutput.addListener(A.H.get<object>(callback));
    },
    "try_OnProcessOutput": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.onProcessOutput.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffProcessOutput": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate?.onProcessOutput && "removeListener" in WEBEXT?.terminalPrivate?.onProcessOutput) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffProcessOutput": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.onProcessOutput.removeListener);
    },
    "call_OffProcessOutput": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.terminalPrivate.onProcessOutput.removeListener(A.H.get<object>(callback));
    },
    "try_OffProcessOutput": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.onProcessOutput.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnProcessOutput": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate?.onProcessOutput && "hasListener" in WEBEXT?.terminalPrivate?.onProcessOutput) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnProcessOutput": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.onProcessOutput.hasListener);
    },
    "call_HasOnProcessOutput": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.terminalPrivate.onProcessOutput.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnProcessOutput": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.onProcessOutput.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnTerminalResize": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate && "onTerminalResize" in WEBEXT?.terminalPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnTerminalResize": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.onTerminalResize);
    },
    "call_OnTerminalResize": (retPtr: Pointer, id: heap.Ref<object>, width: number, height: number): void => {
      const _ret = WEBEXT.terminalPrivate.onTerminalResize(A.H.get<object>(id), width, height);
      A.store.Ref(retPtr, _ret);
    },
    "try_OnTerminalResize": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<object>,
      width: number,
      height: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.onTerminalResize(A.H.get<object>(id), width, height);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenOptionsPage": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate && "openOptionsPage" in WEBEXT?.terminalPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenOptionsPage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.openOptionsPage);
    },
    "call_OpenOptionsPage": (retPtr: Pointer): void => {
      const _ret = WEBEXT.terminalPrivate.openOptionsPage();
      A.store.Ref(retPtr, _ret);
    },
    "try_OpenOptionsPage": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.openOptionsPage();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenSettingsSubpage": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate && "openSettingsSubpage" in WEBEXT?.terminalPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenSettingsSubpage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.openSettingsSubpage);
    },
    "call_OpenSettingsSubpage": (retPtr: Pointer, subpage: heap.Ref<object>): void => {
      const _ret = WEBEXT.terminalPrivate.openSettingsSubpage(A.H.get<object>(subpage));
      A.store.Ref(retPtr, _ret);
    },
    "try_OpenSettingsSubpage": (retPtr: Pointer, errPtr: Pointer, subpage: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.openSettingsSubpage(A.H.get<object>(subpage));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenTerminalProcess": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate && "openTerminalProcess" in WEBEXT?.terminalPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenTerminalProcess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.openTerminalProcess);
    },
    "call_OpenTerminalProcess": (retPtr: Pointer, processName: heap.Ref<object>, args: heap.Ref<object>): void => {
      const _ret = WEBEXT.terminalPrivate.openTerminalProcess(A.H.get<object>(processName), A.H.get<object>(args));
      A.store.Ref(retPtr, _ret);
    },
    "try_OpenTerminalProcess": (
      retPtr: Pointer,
      errPtr: Pointer,
      processName: heap.Ref<object>,
      args: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.openTerminalProcess(A.H.get<object>(processName), A.H.get<object>(args));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenVmshellProcess": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate && "openVmshellProcess" in WEBEXT?.terminalPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenVmshellProcess": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.openVmshellProcess);
    },
    "call_OpenVmshellProcess": (retPtr: Pointer, args: heap.Ref<object>): void => {
      const _ret = WEBEXT.terminalPrivate.openVmshellProcess(A.H.get<object>(args));
      A.store.Ref(retPtr, _ret);
    },
    "try_OpenVmshellProcess": (retPtr: Pointer, errPtr: Pointer, args: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.openVmshellProcess(A.H.get<object>(args));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenWindow": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate && "openWindow" in WEBEXT?.terminalPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenWindow": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.openWindow);
    },
    "call_OpenWindow": (retPtr: Pointer, data: Pointer): void => {
      const data_ffi = {};

      if (A.load.Bool(data + 8)) {
        data_ffi["asTab"] = A.load.Bool(data + 0);
      }
      data_ffi["url"] = A.load.Ref(data + 4, undefined);

      const _ret = WEBEXT.terminalPrivate.openWindow(data_ffi);
    },
    "try_OpenWindow": (retPtr: Pointer, errPtr: Pointer, data: Pointer): heap.Ref<boolean> => {
      try {
        const data_ffi = {};

        if (A.load.Bool(data + 8)) {
          data_ffi["asTab"] = A.load.Bool(data + 0);
        }
        data_ffi["url"] = A.load.Ref(data + 4, undefined);

        const _ret = WEBEXT.terminalPrivate.openWindow(data_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendInput": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate && "sendInput" in WEBEXT?.terminalPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendInput": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.sendInput);
    },
    "call_SendInput": (retPtr: Pointer, id: heap.Ref<object>, input: heap.Ref<object>): void => {
      const _ret = WEBEXT.terminalPrivate.sendInput(A.H.get<object>(id), A.H.get<object>(input));
      A.store.Ref(retPtr, _ret);
    },
    "try_SendInput": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<object>,
      input: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.sendInput(A.H.get<object>(id), A.H.get<object>(input));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetPrefs": (): heap.Ref<boolean> => {
      if (WEBEXT?.terminalPrivate && "setPrefs" in WEBEXT?.terminalPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetPrefs": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.terminalPrivate.setPrefs);
    },
    "call_SetPrefs": (retPtr: Pointer, prefs: heap.Ref<object>): void => {
      const _ret = WEBEXT.terminalPrivate.setPrefs(A.H.get<object>(prefs));
      A.store.Ref(retPtr, _ret);
    },
    "try_SetPrefs": (retPtr: Pointer, errPtr: Pointer, prefs: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.terminalPrivate.setPrefs(A.H.get<object>(prefs));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
