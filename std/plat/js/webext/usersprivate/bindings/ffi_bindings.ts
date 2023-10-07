import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/usersprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_LoginStatusDict": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 4, false);
        A.store.Bool(ptr + 2, false);
        A.store.Bool(ptr + 0, false);
        A.store.Bool(ptr + 3, false);
        A.store.Bool(ptr + 1, false);
      } else {
        A.store.Bool(ptr + 4, true);
        A.store.Bool(ptr + 2, "isLoggedIn" in x ? true : false);
        A.store.Bool(ptr + 0, x["isLoggedIn"] ? true : false);
        A.store.Bool(ptr + 3, "isScreenLocked" in x ? true : false);
        A.store.Bool(ptr + 1, x["isScreenLocked"] ? true : false);
      }
    },
    "load_LoginStatusDict": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 2)) {
        x["isLoggedIn"] = A.load.Bool(ptr + 0);
      } else {
        delete x["isLoggedIn"];
      }
      if (A.load.Bool(ptr + 3)) {
        x["isScreenLocked"] = A.load.Bool(ptr + 1);
      } else {
        delete x["isScreenLocked"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_User": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 16, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 12, false);
        A.store.Bool(ptr + 15, false);
        A.store.Bool(ptr + 13, false);
      } else {
        A.store.Bool(ptr + 16, true);
        A.store.Ref(ptr + 0, x["email"]);
        A.store.Ref(ptr + 4, x["displayEmail"]);
        A.store.Ref(ptr + 8, x["name"]);
        A.store.Bool(ptr + 14, "isOwner" in x ? true : false);
        A.store.Bool(ptr + 12, x["isOwner"] ? true : false);
        A.store.Bool(ptr + 15, "isChild" in x ? true : false);
        A.store.Bool(ptr + 13, x["isChild"] ? true : false);
      }
    },
    "load_User": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["email"] = A.load.Ref(ptr + 0, undefined);
      x["displayEmail"] = A.load.Ref(ptr + 4, undefined);
      x["name"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 14)) {
        x["isOwner"] = A.load.Bool(ptr + 12);
      } else {
        delete x["isOwner"];
      }
      if (A.load.Bool(ptr + 15)) {
        x["isChild"] = A.load.Bool(ptr + 13);
      } else {
        delete x["isChild"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_AddUser": (): heap.Ref<boolean> => {
      if (WEBEXT?.usersPrivate && "addUser" in WEBEXT?.usersPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddUser": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usersPrivate.addUser);
    },
    "call_AddUser": (retPtr: Pointer, email: heap.Ref<object>): void => {
      const _ret = WEBEXT.usersPrivate.addUser(A.H.get<object>(email));
      A.store.Ref(retPtr, _ret);
    },
    "try_AddUser": (retPtr: Pointer, errPtr: Pointer, email: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.usersPrivate.addUser(A.H.get<object>(email));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetCurrentUser": (): heap.Ref<boolean> => {
      if (WEBEXT?.usersPrivate && "getCurrentUser" in WEBEXT?.usersPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetCurrentUser": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usersPrivate.getCurrentUser);
    },
    "call_GetCurrentUser": (retPtr: Pointer): void => {
      const _ret = WEBEXT.usersPrivate.getCurrentUser();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetCurrentUser": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.usersPrivate.getCurrentUser();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetLoginStatus": (): heap.Ref<boolean> => {
      if (WEBEXT?.usersPrivate && "getLoginStatus" in WEBEXT?.usersPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetLoginStatus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usersPrivate.getLoginStatus);
    },
    "call_GetLoginStatus": (retPtr: Pointer): void => {
      const _ret = WEBEXT.usersPrivate.getLoginStatus();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetLoginStatus": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.usersPrivate.getLoginStatus();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetUsers": (): heap.Ref<boolean> => {
      if (WEBEXT?.usersPrivate && "getUsers" in WEBEXT?.usersPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetUsers": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usersPrivate.getUsers);
    },
    "call_GetUsers": (retPtr: Pointer): void => {
      const _ret = WEBEXT.usersPrivate.getUsers();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetUsers": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.usersPrivate.getUsers();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsUserInList": (): heap.Ref<boolean> => {
      if (WEBEXT?.usersPrivate && "isUserInList" in WEBEXT?.usersPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsUserInList": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usersPrivate.isUserInList);
    },
    "call_IsUserInList": (retPtr: Pointer, email: heap.Ref<object>): void => {
      const _ret = WEBEXT.usersPrivate.isUserInList(A.H.get<object>(email));
      A.store.Ref(retPtr, _ret);
    },
    "try_IsUserInList": (retPtr: Pointer, errPtr: Pointer, email: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.usersPrivate.isUserInList(A.H.get<object>(email));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsUserListManaged": (): heap.Ref<boolean> => {
      if (WEBEXT?.usersPrivate && "isUserListManaged" in WEBEXT?.usersPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsUserListManaged": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usersPrivate.isUserListManaged);
    },
    "call_IsUserListManaged": (retPtr: Pointer): void => {
      const _ret = WEBEXT.usersPrivate.isUserListManaged();
      A.store.Ref(retPtr, _ret);
    },
    "try_IsUserListManaged": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.usersPrivate.isUserListManaged();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveUser": (): heap.Ref<boolean> => {
      if (WEBEXT?.usersPrivate && "removeUser" in WEBEXT?.usersPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveUser": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.usersPrivate.removeUser);
    },
    "call_RemoveUser": (retPtr: Pointer, email: heap.Ref<object>): void => {
      const _ret = WEBEXT.usersPrivate.removeUser(A.H.get<object>(email));
      A.store.Ref(retPtr, _ret);
    },
    "try_RemoveUser": (retPtr: Pointer, errPtr: Pointer, email: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.usersPrivate.removeUser(A.H.get<object>(email));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
