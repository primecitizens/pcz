import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/test", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_AssertPromiseRejectsArgExpectedMessageChoice1": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_AssertPromiseRejectsArgExpectedMessageChoice1": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AssertPromiseRejectsArgPromise": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_AssertPromiseRejectsArgPromise": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AssertPromiseRejectsReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_AssertPromiseRejectsReturnType": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_AssertThrowsArgMessageChoice1": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_AssertThrowsArgMessageChoice1": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetConfigReturnTypeFieldFtpServer": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Int64(ptr + 0, x["port"] === undefined ? 0 : (x["port"] as number));
      }
    },
    "load_GetConfigReturnTypeFieldFtpServer": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["port"] = A.load.Int64(ptr + 0);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetConfigReturnTypeFieldLoginStatus": (ptr: Pointer, ref: heap.Ref<any>) => {
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
    "load_GetConfigReturnTypeFieldLoginStatus": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
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

    "store_GetConfigReturnTypeFieldTestServer": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Int64(ptr + 0, 0);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Int64(ptr + 0, x["port"] === undefined ? 0 : (x["port"] as number));
      }
    },
    "load_GetConfigReturnTypeFieldTestServer": (
      ptr: Pointer,
      create: heap.Ref<boolean>,
      ref: heap.Ref<any>
    ): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["port"] = A.load.Int64(ptr + 0);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_GetConfigReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 66, false);
        A.store.Ref(ptr + 0, undefined);

        A.store.Bool(ptr + 8 + 8, false);
        A.store.Int64(ptr + 8 + 0, 0);

        A.store.Bool(ptr + 17 + 4, false);
        A.store.Bool(ptr + 17 + 2, false);
        A.store.Bool(ptr + 17 + 0, false);
        A.store.Bool(ptr + 17 + 3, false);
        A.store.Bool(ptr + 17 + 1, false);
        A.store.Ref(ptr + 24, undefined);

        A.store.Bool(ptr + 32 + 8, false);
        A.store.Int64(ptr + 32 + 0, 0);
        A.store.Bool(ptr + 64, false);
        A.store.Int64(ptr + 48, 0);
        A.store.Bool(ptr + 65, false);
        A.store.Int64(ptr + 56, 0);
      } else {
        A.store.Bool(ptr + 66, true);
        A.store.Ref(ptr + 0, x["customArg"]);

        if (typeof x["ftpServer"] === "undefined") {
          A.store.Bool(ptr + 8 + 8, false);
          A.store.Int64(ptr + 8 + 0, 0);
        } else {
          A.store.Bool(ptr + 8 + 8, true);
          A.store.Int64(ptr + 8 + 0, x["ftpServer"]["port"] === undefined ? 0 : (x["ftpServer"]["port"] as number));
        }

        if (typeof x["loginStatus"] === "undefined") {
          A.store.Bool(ptr + 17 + 4, false);
          A.store.Bool(ptr + 17 + 2, false);
          A.store.Bool(ptr + 17 + 0, false);
          A.store.Bool(ptr + 17 + 3, false);
          A.store.Bool(ptr + 17 + 1, false);
        } else {
          A.store.Bool(ptr + 17 + 4, true);
          A.store.Bool(ptr + 17 + 2, "isLoggedIn" in x["loginStatus"] ? true : false);
          A.store.Bool(ptr + 17 + 0, x["loginStatus"]["isLoggedIn"] ? true : false);
          A.store.Bool(ptr + 17 + 3, "isScreenLocked" in x["loginStatus"] ? true : false);
          A.store.Bool(ptr + 17 + 1, x["loginStatus"]["isScreenLocked"] ? true : false);
        }
        A.store.Ref(ptr + 24, x["testDataDirectory"]);

        if (typeof x["testServer"] === "undefined") {
          A.store.Bool(ptr + 32 + 8, false);
          A.store.Int64(ptr + 32 + 0, 0);
        } else {
          A.store.Bool(ptr + 32 + 8, true);
          A.store.Int64(ptr + 32 + 0, x["testServer"]["port"] === undefined ? 0 : (x["testServer"]["port"] as number));
        }
        A.store.Bool(ptr + 64, "testWebSocketPort" in x ? true : false);
        A.store.Int64(ptr + 48, x["testWebSocketPort"] === undefined ? 0 : (x["testWebSocketPort"] as number));
        A.store.Bool(ptr + 65, "testWebTransportPort" in x ? true : false);
        A.store.Int64(ptr + 56, x["testWebTransportPort"] === undefined ? 0 : (x["testWebTransportPort"] as number));
      }
    },
    "load_GetConfigReturnType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["customArg"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 8 + 8)) {
        x["ftpServer"] = {};
        x["ftpServer"]["port"] = A.load.Int64(ptr + 8 + 0);
      } else {
        delete x["ftpServer"];
      }
      if (A.load.Bool(ptr + 17 + 4)) {
        x["loginStatus"] = {};
        if (A.load.Bool(ptr + 17 + 2)) {
          x["loginStatus"]["isLoggedIn"] = A.load.Bool(ptr + 17 + 0);
        } else {
          delete x["loginStatus"]["isLoggedIn"];
        }
        if (A.load.Bool(ptr + 17 + 3)) {
          x["loginStatus"]["isScreenLocked"] = A.load.Bool(ptr + 17 + 1);
        } else {
          delete x["loginStatus"]["isScreenLocked"];
        }
      } else {
        delete x["loginStatus"];
      }
      x["testDataDirectory"] = A.load.Ref(ptr + 24, undefined);
      if (A.load.Bool(ptr + 32 + 8)) {
        x["testServer"] = {};
        x["testServer"]["port"] = A.load.Int64(ptr + 32 + 0);
      } else {
        delete x["testServer"];
      }
      if (A.load.Bool(ptr + 64)) {
        x["testWebSocketPort"] = A.load.Int64(ptr + 48);
      } else {
        delete x["testWebSocketPort"];
      }
      if (A.load.Bool(ptr + 65)) {
        x["testWebTransportPort"] = A.load.Int64(ptr + 56);
      } else {
        delete x["testWebTransportPort"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_LoadScriptReturnType": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 0, false);
      } else {
        A.store.Bool(ptr + 0, true);
      }
    },
    "load_LoadScriptReturnType": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "store_OnMessageArgInfo": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 5, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 4, false);
      } else {
        A.store.Bool(ptr + 5, true);
        A.store.Ref(ptr + 0, x["data"]);
        A.store.Bool(ptr + 4, x["lastMessage"] ? true : false);
      }
    },
    "load_OnMessageArgInfo": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["data"] = A.load.Ref(ptr + 0, undefined);
      x["lastMessage"] = A.load.Bool(ptr + 4);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "has_AssertEq": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "assertEq" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AssertEq": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.assertEq);
    },
    "call_AssertEq": (
      retPtr: Pointer,
      expected: heap.Ref<object>,
      actual: heap.Ref<object>,
      message: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.test.assertEq(A.H.get<object>(expected), A.H.get<object>(actual), A.H.get<object>(message));
    },
    "try_AssertEq": (
      retPtr: Pointer,
      errPtr: Pointer,
      expected: heap.Ref<object>,
      actual: heap.Ref<object>,
      message: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.assertEq(A.H.get<object>(expected), A.H.get<object>(actual), A.H.get<object>(message));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AssertFalse": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "assertFalse" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AssertFalse": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.assertFalse);
    },
    "call_AssertFalse": (retPtr: Pointer, test: heap.Ref<any>, message: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.assertFalse(A.H.get<any>(test), A.H.get<object>(message));
    },
    "try_AssertFalse": (
      retPtr: Pointer,
      errPtr: Pointer,
      test: heap.Ref<any>,
      message: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.assertFalse(A.H.get<any>(test), A.H.get<object>(message));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AssertLastError": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "assertLastError" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AssertLastError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.assertLastError);
    },
    "call_AssertLastError": (retPtr: Pointer, expectedError: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.assertLastError(A.H.get<object>(expectedError));
    },
    "try_AssertLastError": (retPtr: Pointer, errPtr: Pointer, expectedError: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.assertLastError(A.H.get<object>(expectedError));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AssertNe": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "assertNe" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AssertNe": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.assertNe);
    },
    "call_AssertNe": (
      retPtr: Pointer,
      expected: heap.Ref<object>,
      actual: heap.Ref<object>,
      message: heap.Ref<object>
    ): void => {
      const _ret = WEBEXT.test.assertNe(A.H.get<object>(expected), A.H.get<object>(actual), A.H.get<object>(message));
    },
    "try_AssertNe": (
      retPtr: Pointer,
      errPtr: Pointer,
      expected: heap.Ref<object>,
      actual: heap.Ref<object>,
      message: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.assertNe(A.H.get<object>(expected), A.H.get<object>(actual), A.H.get<object>(message));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AssertNoLastError": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "assertNoLastError" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AssertNoLastError": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.assertNoLastError);
    },
    "call_AssertNoLastError": (retPtr: Pointer): void => {
      const _ret = WEBEXT.test.assertNoLastError();
    },
    "try_AssertNoLastError": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.assertNoLastError();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AssertPromiseRejects": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "assertPromiseRejects" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AssertPromiseRejects": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.assertPromiseRejects);
    },
    "call_AssertPromiseRejects": (retPtr: Pointer, promise: Pointer, expectedMessage: heap.Ref<any>): void => {
      const promise_ffi = {};

      const _ret = WEBEXT.test.assertPromiseRejects(promise_ffi, A.H.get<any>(expectedMessage));
      if (typeof _ret === "undefined") {
        A.store.Bool(retPtr + 0, false);
      } else {
        A.store.Bool(retPtr + 0, true);
      }
    },
    "try_AssertPromiseRejects": (
      retPtr: Pointer,
      errPtr: Pointer,
      promise: Pointer,
      expectedMessage: heap.Ref<any>
    ): heap.Ref<boolean> => {
      try {
        const promise_ffi = {};

        const _ret = WEBEXT.test.assertPromiseRejects(promise_ffi, A.H.get<any>(expectedMessage));
        if (typeof _ret === "undefined") {
          A.store.Bool(retPtr + 0, false);
        } else {
          A.store.Bool(retPtr + 0, true);
        }
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AssertThrows": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "assertThrows" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AssertThrows": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.assertThrows);
    },
    "call_AssertThrows": (
      retPtr: Pointer,
      fn: heap.Ref<object>,
      self: heap.Ref<object>,
      args: heap.Ref<object>,
      message: heap.Ref<any>
    ): void => {
      const _ret = WEBEXT.test.assertThrows(
        A.H.get<object>(fn),
        A.H.get<object>(self),
        A.H.get<object>(args),
        A.H.get<any>(message)
      );
    },
    "try_AssertThrows": (
      retPtr: Pointer,
      errPtr: Pointer,
      fn: heap.Ref<object>,
      self: heap.Ref<object>,
      args: heap.Ref<object>,
      message: heap.Ref<any>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.assertThrows(
          A.H.get<object>(fn),
          A.H.get<object>(self),
          A.H.get<object>(args),
          A.H.get<any>(message)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AssertTrue": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "assertTrue" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AssertTrue": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.assertTrue);
    },
    "call_AssertTrue": (retPtr: Pointer, test: heap.Ref<any>, message: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.assertTrue(A.H.get<any>(test), A.H.get<object>(message));
    },
    "try_AssertTrue": (
      retPtr: Pointer,
      errPtr: Pointer,
      test: heap.Ref<any>,
      message: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.assertTrue(A.H.get<any>(test), A.H.get<object>(message));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Callback": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "callback" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Callback": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.callback);
    },
    "call_Callback": (retPtr: Pointer, func: heap.Ref<object>, expectedError: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.callback(A.H.get<object>(func), A.H.get<object>(expectedError));
    },
    "try_Callback": (
      retPtr: Pointer,
      errPtr: Pointer,
      func: heap.Ref<object>,
      expectedError: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.callback(A.H.get<object>(func), A.H.get<object>(expectedError));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CallbackAdded": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "callbackAdded" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CallbackAdded": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.callbackAdded);
    },
    "call_CallbackAdded": (retPtr: Pointer): void => {
      const _ret = WEBEXT.test.callbackAdded();
    },
    "try_CallbackAdded": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.callbackAdded();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CallbackFail": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "callbackFail" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CallbackFail": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.callbackFail);
    },
    "call_CallbackFail": (retPtr: Pointer, expectedError: heap.Ref<object>, func: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.callbackFail(A.H.get<object>(expectedError), A.H.get<object>(func));
    },
    "try_CallbackFail": (
      retPtr: Pointer,
      errPtr: Pointer,
      expectedError: heap.Ref<object>,
      func: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.callbackFail(A.H.get<object>(expectedError), A.H.get<object>(func));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CallbackPass": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "callbackPass" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CallbackPass": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.callbackPass);
    },
    "call_CallbackPass": (retPtr: Pointer, func: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.callbackPass(A.H.get<object>(func));
    },
    "try_CallbackPass": (retPtr: Pointer, errPtr: Pointer, func: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.callbackPass(A.H.get<object>(func));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_CheckDeepEq": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "checkDeepEq" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_CheckDeepEq": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.checkDeepEq);
    },
    "call_CheckDeepEq": (retPtr: Pointer, expected: heap.Ref<object>, actual: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.checkDeepEq(A.H.get<object>(expected), A.H.get<object>(actual));
    },
    "try_CheckDeepEq": (
      retPtr: Pointer,
      errPtr: Pointer,
      expected: heap.Ref<object>,
      actual: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.checkDeepEq(A.H.get<object>(expected), A.H.get<object>(actual));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Fail": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "fail" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Fail": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.fail);
    },
    "call_Fail": (retPtr: Pointer, message: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.fail(A.H.get<object>(message));
    },
    "try_Fail": (retPtr: Pointer, errPtr: Pointer, message: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.fail(A.H.get<object>(message));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetApiDefinitions": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "getApiDefinitions" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetApiDefinitions": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.getApiDefinitions);
    },
    "call_GetApiDefinitions": (retPtr: Pointer, apiNames: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.getApiDefinitions(A.H.get<object>(apiNames));
    },
    "try_GetApiDefinitions": (retPtr: Pointer, errPtr: Pointer, apiNames: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.getApiDefinitions(A.H.get<object>(apiNames));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetApiFeatures": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "getApiFeatures" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetApiFeatures": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.getApiFeatures);
    },
    "call_GetApiFeatures": (retPtr: Pointer): void => {
      const _ret = WEBEXT.test.getApiFeatures();
    },
    "try_GetApiFeatures": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.getApiFeatures();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetConfig": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "getConfig" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetConfig": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.getConfig);
    },
    "call_GetConfig": (retPtr: Pointer): void => {
      const _ret = WEBEXT.test.getConfig();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetConfig": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.getConfig();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetModuleSystem": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "getModuleSystem" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetModuleSystem": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.getModuleSystem);
    },
    "call_GetModuleSystem": (retPtr: Pointer, context: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.getModuleSystem(A.H.get<object>(context));
      A.store.Ref(retPtr, _ret);
    },
    "try_GetModuleSystem": (retPtr: Pointer, errPtr: Pointer, context: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.getModuleSystem(A.H.get<object>(context));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetWakeEventPage": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "getWakeEventPage" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetWakeEventPage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.getWakeEventPage);
    },
    "call_GetWakeEventPage": (retPtr: Pointer): void => {
      const _ret = WEBEXT.test.getWakeEventPage();
      A.store.Ref(retPtr, _ret);
    },
    "try_GetWakeEventPage": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.getWakeEventPage();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_IsProcessingUserGesture": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "isProcessingUserGesture" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_IsProcessingUserGesture": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.isProcessingUserGesture);
    },
    "call_IsProcessingUserGesture": (retPtr: Pointer): void => {
      const _ret = WEBEXT.test.isProcessingUserGesture();
    },
    "try_IsProcessingUserGesture": (retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.isProcessingUserGesture();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ListenForever": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "listenForever" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ListenForever": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.listenForever);
    },
    "call_ListenForever": (retPtr: Pointer, event: heap.Ref<object>, func: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.listenForever(A.H.get<object>(event), A.H.get<object>(func));
    },
    "try_ListenForever": (
      retPtr: Pointer,
      errPtr: Pointer,
      event: heap.Ref<object>,
      func: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.listenForever(A.H.get<object>(event), A.H.get<object>(func));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ListenOnce": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "listenOnce" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ListenOnce": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.listenOnce);
    },
    "call_ListenOnce": (retPtr: Pointer, event: heap.Ref<object>, func: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.listenOnce(A.H.get<object>(event), A.H.get<object>(func));
    },
    "try_ListenOnce": (
      retPtr: Pointer,
      errPtr: Pointer,
      event: heap.Ref<object>,
      func: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.listenOnce(A.H.get<object>(event), A.H.get<object>(func));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_LoadScript": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "loadScript" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_LoadScript": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.loadScript);
    },
    "call_LoadScript": (retPtr: Pointer, scriptUrl: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.loadScript(A.H.get<object>(scriptUrl));
      if (typeof _ret === "undefined") {
        A.store.Bool(retPtr + 0, false);
      } else {
        A.store.Bool(retPtr + 0, true);
      }
    },
    "try_LoadScript": (retPtr: Pointer, errPtr: Pointer, scriptUrl: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.loadScript(A.H.get<object>(scriptUrl));
        if (typeof _ret === "undefined") {
          A.store.Bool(retPtr + 0, false);
        } else {
          A.store.Bool(retPtr + 0, true);
        }
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Log": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "log" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Log": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.log);
    },
    "call_Log": (retPtr: Pointer, message: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.log(A.H.get<object>(message));
    },
    "try_Log": (retPtr: Pointer, errPtr: Pointer, message: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.log(A.H.get<object>(message));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_NotifyFail": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "notifyFail" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_NotifyFail": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.notifyFail);
    },
    "call_NotifyFail": (retPtr: Pointer, message: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.notifyFail(A.H.get<object>(message));
    },
    "try_NotifyFail": (retPtr: Pointer, errPtr: Pointer, message: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.notifyFail(A.H.get<object>(message));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_NotifyPass": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "notifyPass" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_NotifyPass": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.notifyPass);
    },
    "call_NotifyPass": (retPtr: Pointer, message: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.notifyPass(A.H.get<object>(message));
    },
    "try_NotifyPass": (retPtr: Pointer, errPtr: Pointer, message: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.notifyPass(A.H.get<object>(message));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OnMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.test?.onMessage && "addListener" in WEBEXT?.test?.onMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OnMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.onMessage.addListener);
    },
    "call_OnMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.onMessage.addListener(A.H.get<object>(callback));
    },
    "try_OnMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.onMessage.addListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OffMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.test?.onMessage && "removeListener" in WEBEXT?.test?.onMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OffMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.onMessage.removeListener);
    },
    "call_OffMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.onMessage.removeListener(A.H.get<object>(callback));
    },
    "try_OffMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.onMessage.removeListener(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_HasOnMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.test?.onMessage && "hasListener" in WEBEXT?.test?.onMessage) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_HasOnMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.onMessage.hasListener);
    },
    "call_HasOnMessage": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.onMessage.hasListener(A.H.get<object>(callback));
      A.store.Bool(retPtr, _ret);
    },
    "try_HasOnMessage": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.onMessage.hasListener(A.H.get<object>(callback));
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_OpenFileUrl": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "openFileUrl" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_OpenFileUrl": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.openFileUrl);
    },
    "call_OpenFileUrl": (retPtr: Pointer, url: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.openFileUrl(A.H.get<object>(url));
    },
    "try_OpenFileUrl": (retPtr: Pointer, errPtr: Pointer, url: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.openFileUrl(A.H.get<object>(url));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunTests": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "runTests" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunTests": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.runTests);
    },
    "call_RunTests": (retPtr: Pointer, tests: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.runTests(A.H.get<object>(tests));
    },
    "try_RunTests": (retPtr: Pointer, errPtr: Pointer, tests: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.runTests(A.H.get<object>(tests));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RunWithUserGesture": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "runWithUserGesture" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RunWithUserGesture": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.runWithUserGesture);
    },
    "call_RunWithUserGesture": (retPtr: Pointer, functionToRun: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.runWithUserGesture(A.H.get<object>(functionToRun));
    },
    "try_RunWithUserGesture": (
      retPtr: Pointer,
      errPtr: Pointer,
      functionToRun: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.runWithUserGesture(A.H.get<object>(functionToRun));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendMessage": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "sendMessage" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendMessage": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.sendMessage);
    },
    "call_SendMessage": (retPtr: Pointer, message: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.sendMessage(A.H.get<object>(message));
      A.store.Ref(retPtr, _ret);
    },
    "try_SendMessage": (retPtr: Pointer, errPtr: Pointer, message: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.sendMessage(A.H.get<object>(message));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SendScriptResult": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "sendScriptResult" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SendScriptResult": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.sendScriptResult);
    },
    "call_SendScriptResult": (retPtr: Pointer, result: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.sendScriptResult(A.H.get<object>(result));
      A.store.Ref(retPtr, _ret);
    },
    "try_SendScriptResult": (retPtr: Pointer, errPtr: Pointer, result: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.sendScriptResult(A.H.get<object>(result));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetExceptionHandler": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "setExceptionHandler" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetExceptionHandler": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.setExceptionHandler);
    },
    "call_SetExceptionHandler": (retPtr: Pointer, handler: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.setExceptionHandler(A.H.get<object>(handler));
    },
    "try_SetExceptionHandler": (retPtr: Pointer, errPtr: Pointer, handler: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.setExceptionHandler(A.H.get<object>(handler));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_Succeed": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "succeed" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_Succeed": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.succeed);
    },
    "call_Succeed": (retPtr: Pointer, message: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.succeed(A.H.get<object>(message));
    },
    "try_Succeed": (retPtr: Pointer, errPtr: Pointer, message: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.succeed(A.H.get<object>(message));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_WaitForRoundTrip": (): heap.Ref<boolean> => {
      if (WEBEXT?.test && "waitForRoundTrip" in WEBEXT?.test) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_WaitForRoundTrip": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.test.waitForRoundTrip);
    },
    "call_WaitForRoundTrip": (retPtr: Pointer, message: heap.Ref<object>): void => {
      const _ret = WEBEXT.test.waitForRoundTrip(A.H.get<object>(message));
      A.store.Ref(retPtr, _ret);
    },
    "try_WaitForRoundTrip": (retPtr: Pointer, errPtr: Pointer, message: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.test.waitForRoundTrip(A.H.get<object>(message));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
