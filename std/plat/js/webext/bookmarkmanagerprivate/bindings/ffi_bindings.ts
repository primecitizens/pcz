import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/bookmarkmanagerprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_BookmarkNodeDataElement": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Ref(ptr + 12, undefined);
        A.store.Ref(ptr + 16, undefined);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Ref(ptr + 0, x["children"]);
        A.store.Ref(ptr + 4, x["id"]);
        A.store.Ref(ptr + 8, x["parentId"]);
        A.store.Ref(ptr + 12, x["title"]);
        A.store.Ref(ptr + 16, x["url"]);
      }
    },
    "load_BookmarkNodeDataElement": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["children"] = A.load.Ref(ptr + 0, undefined);
      x["id"] = A.load.Ref(ptr + 4, undefined);
      x["parentId"] = A.load.Ref(ptr + 8, undefined);
      x["title"] = A.load.Ref(ptr + 12, undefined);
      x["url"] = A.load.Ref(ptr + 16, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_BookmarkNodeData": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Ref(ptr + 0, x["elements"]);
        A.store.Bool(ptr + 4, x["sameProfile"] ? true : false);
      }
    },
    "load_BookmarkNodeData": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["elements"] = A.load.Ref(ptr + 0, undefined);
      x["sameProfile"] = A.load.Bool(ptr + 4);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_CanPaste": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate && "canPaste" in WEBEXT?.bookmarkManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CanPaste": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.canPaste);
    },
    "call_CanPaste": (retPtr: Pointer, parentId: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.canPaste(A.H.get<object>(parentId));
      A.store.Ref(retPtr, _ret);
    },
    "try_CanPaste": (retPtr: Pointer, errPtr: Pointer, parentId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.canPaste(A.H.get<object>(parentId));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Copy": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate && "copy" in WEBEXT?.bookmarkManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Copy": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.copy);
    },
    "call_Copy": (retPtr: Pointer, idList: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.copy(A.H.get<object>(idList));
      A.store.Ref(retPtr, _ret);
    },
    "try_Copy": (retPtr: Pointer, errPtr: Pointer, idList: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.copy(A.H.get<object>(idList));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Cut": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate && "cut" in WEBEXT?.bookmarkManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Cut": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.cut);
    },
    "call_Cut": (retPtr: Pointer, idList: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.cut(A.H.get<object>(idList));
      A.store.Ref(retPtr, _ret);
    },
    "try_Cut": (retPtr: Pointer, errPtr: Pointer, idList: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.cut(A.H.get<object>(idList));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Drop": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate && "drop" in WEBEXT?.bookmarkManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Drop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.drop);
    },
    "call_Drop": (retPtr: Pointer, parentId: heap.Ref<object>, index: number): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.drop(A.H.get<object>(parentId), index);
      A.store.Ref(retPtr, _ret);
    },
    "try_Drop": (retPtr: Pointer, errPtr: Pointer, parentId: heap.Ref<object>, index: number): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.drop(A.H.get<object>(parentId), index);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Export": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate && "export" in WEBEXT?.bookmarkManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Export": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.export);
    },
    "call_Export": (retPtr: Pointer): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.export();
      A.store.Ref(retPtr, _ret);
    },
    "try_Export": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.export();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetSubtree": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate && "getSubtree" in WEBEXT?.bookmarkManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetSubtree": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.getSubtree);
    },
    "call_GetSubtree": (retPtr: Pointer, id: heap.Ref<object>, foldersOnly: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.getSubtree(A.H.get<object>(id), foldersOnly === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_GetSubtree": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<object>,
      foldersOnly: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.getSubtree(A.H.get<object>(id), foldersOnly === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Import": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate && "import" in WEBEXT?.bookmarkManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Import": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.import);
    },
    "call_Import": (retPtr: Pointer): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.import();
      A.store.Ref(retPtr, _ret);
    },
    "try_Import": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.import();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDragEnter": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate?.onDragEnter && "addListener" in WEBEXT?.bookmarkManagerPrivate?.onDragEnter) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDragEnter": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.onDragEnter.addListener);
    },
    "call_OnDragEnter": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.onDragEnter.addListener(A.H.get<object>(callback));
    },
    "try_OnDragEnter": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.onDragEnter.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDragEnter": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bookmarkManagerPrivate?.onDragEnter &&
        "removeListener" in WEBEXT?.bookmarkManagerPrivate?.onDragEnter
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDragEnter": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.onDragEnter.removeListener);
    },
    "call_OffDragEnter": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.onDragEnter.removeListener(A.H.get<object>(callback));
    },
    "try_OffDragEnter": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.onDragEnter.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDragEnter": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate?.onDragEnter && "hasListener" in WEBEXT?.bookmarkManagerPrivate?.onDragEnter) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDragEnter": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.onDragEnter.hasListener);
    },
    "call_HasOnDragEnter": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.onDragEnter.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDragEnter": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.onDragEnter.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDragLeave": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate?.onDragLeave && "addListener" in WEBEXT?.bookmarkManagerPrivate?.onDragLeave) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDragLeave": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.onDragLeave.addListener);
    },
    "call_OnDragLeave": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.onDragLeave.addListener(A.H.get<object>(callback));
    },
    "try_OnDragLeave": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.onDragLeave.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDragLeave": (): heap.Ref<boolean> => {
      if (
        WEBEXT?.bookmarkManagerPrivate?.onDragLeave &&
        "removeListener" in WEBEXT?.bookmarkManagerPrivate?.onDragLeave
      ) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDragLeave": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.onDragLeave.removeListener);
    },
    "call_OffDragLeave": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.onDragLeave.removeListener(A.H.get<object>(callback));
    },
    "try_OffDragLeave": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.onDragLeave.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDragLeave": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate?.onDragLeave && "hasListener" in WEBEXT?.bookmarkManagerPrivate?.onDragLeave) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDragLeave": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.onDragLeave.hasListener);
    },
    "call_HasOnDragLeave": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.onDragLeave.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDragLeave": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.onDragLeave.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnDrop": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate?.onDrop && "addListener" in WEBEXT?.bookmarkManagerPrivate?.onDrop) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnDrop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.onDrop.addListener);
    },
    "call_OnDrop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.onDrop.addListener(A.H.get<object>(callback));
    },
    "try_OnDrop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.onDrop.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffDrop": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate?.onDrop && "removeListener" in WEBEXT?.bookmarkManagerPrivate?.onDrop) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffDrop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.onDrop.removeListener);
    },
    "call_OffDrop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.onDrop.removeListener(A.H.get<object>(callback));
    },
    "try_OffDrop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.onDrop.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnDrop": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate?.onDrop && "hasListener" in WEBEXT?.bookmarkManagerPrivate?.onDrop) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnDrop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.onDrop.hasListener);
    },
    "call_HasOnDrop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.onDrop.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnDrop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.onDrop.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenInNewTab": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate && "openInNewTab" in WEBEXT?.bookmarkManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenInNewTab": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.openInNewTab);
    },
    "call_OpenInNewTab": (retPtr: Pointer, id: heap.Ref<object>, active: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.openInNewTab(A.H.get<object>(id), active === A.H.TRUE);
    },
    "try_OpenInNewTab": (
      retPtr: Pointer,
      errPtr: Pointer,
      id: heap.Ref<object>,
      active: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.openInNewTab(A.H.get<object>(id), active === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenInNewWindow": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate && "openInNewWindow" in WEBEXT?.bookmarkManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenInNewWindow": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.openInNewWindow);
    },
    "call_OpenInNewWindow": (retPtr: Pointer, idList: heap.Ref<object>, incognito: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.openInNewWindow(A.H.get<object>(idList), incognito === A.H.TRUE);
    },
    "try_OpenInNewWindow": (
      retPtr: Pointer,
      errPtr: Pointer,
      idList: heap.Ref<object>,
      incognito: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.openInNewWindow(A.H.get<object>(idList), incognito === A.H.TRUE);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Paste": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate && "paste" in WEBEXT?.bookmarkManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Paste": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.paste);
    },
    "call_Paste": (retPtr: Pointer, parentId: heap.Ref<object>, selectedIdList: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.paste(A.H.get<object>(parentId), A.H.get<object>(selectedIdList));
      A.store.Ref(retPtr, _ret);
    },
    "try_Paste": (
      retPtr: Pointer,
      errPtr: Pointer,
      parentId: heap.Ref<object>,
      selectedIdList: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.paste(A.H.get<object>(parentId), A.H.get<object>(selectedIdList));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Redo": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate && "redo" in WEBEXT?.bookmarkManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Redo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.redo);
    },
    "call_Redo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.redo();
    },
    "try_Redo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.redo();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveTrees": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate && "removeTrees" in WEBEXT?.bookmarkManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveTrees": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.removeTrees);
    },
    "call_RemoveTrees": (retPtr: Pointer, idList: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.removeTrees(A.H.get<object>(idList));
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveTrees": (retPtr: Pointer, errPtr: Pointer, idList: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.removeTrees(A.H.get<object>(idList));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SortChildren": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate && "sortChildren" in WEBEXT?.bookmarkManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SortChildren": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.sortChildren);
    },
    "call_SortChildren": (retPtr: Pointer, parentId: heap.Ref<object>): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.sortChildren(A.H.get<object>(parentId));
    },
    "try_SortChildren": (retPtr: Pointer, errPtr: Pointer, parentId: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.sortChildren(A.H.get<object>(parentId));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_StartDrag": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate && "startDrag" in WEBEXT?.bookmarkManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_StartDrag": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.startDrag);
    },
    "call_StartDrag": (
      retPtr: Pointer,
      idList: heap.Ref<object>,
      dragNodeIndex: number,
      isFromTouch: heap.Ref<boolean>,
      x: number,
      y: number
    ): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.startDrag(
        A.H.get<object>(idList),
        dragNodeIndex,
        isFromTouch === A.H.TRUE,
        x,
        y
      );
    },
    "try_StartDrag": (
      retPtr: Pointer,
      errPtr: Pointer,
      idList: heap.Ref<object>,
      dragNodeIndex: number,
      isFromTouch: heap.Ref<boolean>,
      x: number,
      y: number
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.startDrag(
          A.H.get<object>(idList),
          dragNodeIndex,
          isFromTouch === A.H.TRUE,
          x,
          y
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Undo": (): heap.Ref<boolean> => {
      if (WEBEXT?.bookmarkManagerPrivate && "undo" in WEBEXT?.bookmarkManagerPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Undo": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.bookmarkManagerPrivate.undo);
    },
    "call_Undo": (retPtr: Pointer): void => {
      const _ret = WEBEXT.bookmarkManagerPrivate.undo();
    },
    "try_Undo": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.bookmarkManagerPrivate.undo();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
