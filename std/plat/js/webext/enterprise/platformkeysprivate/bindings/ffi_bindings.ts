import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/enterprise/platformkeysprivate", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "has_ChallengeMachineKey": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.platformKeysPrivate && "challengeMachineKey" in WEBEXT?.enterprise?.platformKeysPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ChallengeMachineKey": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.platformKeysPrivate.challengeMachineKey);
    },
    "call_ChallengeMachineKey": (retPtr: Pointer, challenge: heap.Ref<object>): void => {
      const _ret = WEBEXT.enterprise.platformKeysPrivate.challengeMachineKey(A.H.get<object>(challenge));
      A.store.Ref(retPtr, _ret);
    },
    "try_ChallengeMachineKey": (retPtr: Pointer, errPtr: Pointer, challenge: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.platformKeysPrivate.challengeMachineKey(A.H.get<object>(challenge));
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_ChallengeUserKey": (): heap.Ref<boolean> => {
      if (WEBEXT?.enterprise?.platformKeysPrivate && "challengeUserKey" in WEBEXT?.enterprise?.platformKeysPrivate) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_ChallengeUserKey": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.enterprise.platformKeysPrivate.challengeUserKey);
    },
    "call_ChallengeUserKey": (retPtr: Pointer, challenge: heap.Ref<object>, registerKey: heap.Ref<boolean>): void => {
      const _ret = WEBEXT.enterprise.platformKeysPrivate.challengeUserKey(
        A.H.get<object>(challenge),
        registerKey === A.H.TRUE
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_ChallengeUserKey": (
      retPtr: Pointer,
      errPtr: Pointer,
      challenge: heap.Ref<object>,
      registerKey: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.enterprise.platformKeysPrivate.challengeUserKey(
          A.H.get<object>(challenge),
          registerKey === A.H.TRUE
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
