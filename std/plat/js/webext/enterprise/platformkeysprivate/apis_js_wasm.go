// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package platformkeysprivate

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/enterprise/platformkeysprivate/bindings"
)

// HasFuncChallengeMachineKey returns true if the function "WEBEXT.enterprise.platformKeysPrivate.challengeMachineKey" exists.
func HasFuncChallengeMachineKey() bool {
	return js.True == bindings.HasFuncChallengeMachineKey()
}

// FuncChallengeMachineKey returns the function "WEBEXT.enterprise.platformKeysPrivate.challengeMachineKey".
func FuncChallengeMachineKey() (fn js.Func[func(challenge js.String) js.Promise[js.String]]) {
	bindings.FuncChallengeMachineKey(
		js.Pointer(&fn),
	)
	return
}

// ChallengeMachineKey calls the function "WEBEXT.enterprise.platformKeysPrivate.challengeMachineKey" directly.
func ChallengeMachineKey(challenge js.String) (ret js.Promise[js.String]) {
	bindings.CallChallengeMachineKey(
		js.Pointer(&ret),
		challenge.Ref(),
	)

	return
}

// TryChallengeMachineKey calls the function "WEBEXT.enterprise.platformKeysPrivate.challengeMachineKey"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryChallengeMachineKey(challenge js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryChallengeMachineKey(
		js.Pointer(&ret), js.Pointer(&exception),
		challenge.Ref(),
	)

	return
}

// HasFuncChallengeUserKey returns true if the function "WEBEXT.enterprise.platformKeysPrivate.challengeUserKey" exists.
func HasFuncChallengeUserKey() bool {
	return js.True == bindings.HasFuncChallengeUserKey()
}

// FuncChallengeUserKey returns the function "WEBEXT.enterprise.platformKeysPrivate.challengeUserKey".
func FuncChallengeUserKey() (fn js.Func[func(challenge js.String, registerKey bool) js.Promise[js.String]]) {
	bindings.FuncChallengeUserKey(
		js.Pointer(&fn),
	)
	return
}

// ChallengeUserKey calls the function "WEBEXT.enterprise.platformKeysPrivate.challengeUserKey" directly.
func ChallengeUserKey(challenge js.String, registerKey bool) (ret js.Promise[js.String]) {
	bindings.CallChallengeUserKey(
		js.Pointer(&ret),
		challenge.Ref(),
		js.Bool(bool(registerKey)),
	)

	return
}

// TryChallengeUserKey calls the function "WEBEXT.enterprise.platformKeysPrivate.challengeUserKey"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryChallengeUserKey(challenge js.String, registerKey bool) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryChallengeUserKey(
		js.Pointer(&ret), js.Pointer(&exception),
		challenge.Ref(),
		js.Bool(bool(registerKey)),
	)

	return
}
