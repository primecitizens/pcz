// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package loginstate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/loginstate/bindings"
)

type ProfileType uint32

const (
	_ ProfileType = iota

	ProfileType_SIGNIN_PROFILE
	ProfileType_USER_PROFILE
)

func (ProfileType) FromRef(str js.Ref) ProfileType {
	return ProfileType(bindings.ConstOfProfileType(str))
}

func (x ProfileType) String() (string, bool) {
	switch x {
	case ProfileType_SIGNIN_PROFILE:
		return "SIGNIN_PROFILE", true
	case ProfileType_USER_PROFILE:
		return "USER_PROFILE", true
	default:
		return "", false
	}
}

type ProfileTypeCallbackFunc func(this js.Ref, result ProfileType) js.Ref

func (fn ProfileTypeCallbackFunc) Register() js.Func[func(result ProfileType)] {
	return js.RegisterCallback[func(result ProfileType)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ProfileTypeCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		ProfileType(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ProfileTypeCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result ProfileType) js.Ref
	Arg T
}

func (cb *ProfileTypeCallback[T]) Register() js.Func[func(result ProfileType)] {
	return js.RegisterCallback[func(result ProfileType)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ProfileTypeCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		ProfileType(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SessionState uint32

const (
	_ SessionState = iota

	SessionState_UNKNOWN
	SessionState_IN_OOBE_SCREEN
	SessionState_IN_LOGIN_SCREEN
	SessionState_IN_SESSION
	SessionState_IN_LOCK_SCREEN
	SessionState_IN_RMA_SCREEN
)

func (SessionState) FromRef(str js.Ref) SessionState {
	return SessionState(bindings.ConstOfSessionState(str))
}

func (x SessionState) String() (string, bool) {
	switch x {
	case SessionState_UNKNOWN:
		return "UNKNOWN", true
	case SessionState_IN_OOBE_SCREEN:
		return "IN_OOBE_SCREEN", true
	case SessionState_IN_LOGIN_SCREEN:
		return "IN_LOGIN_SCREEN", true
	case SessionState_IN_SESSION:
		return "IN_SESSION", true
	case SessionState_IN_LOCK_SCREEN:
		return "IN_LOCK_SCREEN", true
	case SessionState_IN_RMA_SCREEN:
		return "IN_RMA_SCREEN", true
	default:
		return "", false
	}
}

type SessionStateCallbackFunc func(this js.Ref, result SessionState) js.Ref

func (fn SessionStateCallbackFunc) Register() js.Func[func(result SessionState)] {
	return js.RegisterCallback[func(result SessionState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SessionStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		SessionState(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SessionStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result SessionState) js.Ref
	Arg T
}

func (cb *SessionStateCallback[T]) Register() js.Func[func(result SessionState)] {
	return js.RegisterCallback[func(result SessionState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SessionStateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		SessionState(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncGetProfileType returns true if the function "WEBEXT.loginState.getProfileType" exists.
func HasFuncGetProfileType() bool {
	return js.True == bindings.HasFuncGetProfileType()
}

// FuncGetProfileType returns the function "WEBEXT.loginState.getProfileType".
func FuncGetProfileType() (fn js.Func[func() js.Promise[ProfileType]]) {
	bindings.FuncGetProfileType(
		js.Pointer(&fn),
	)
	return
}

// GetProfileType calls the function "WEBEXT.loginState.getProfileType" directly.
func GetProfileType() (ret js.Promise[ProfileType]) {
	bindings.CallGetProfileType(
		js.Pointer(&ret),
	)

	return
}

// TryGetProfileType calls the function "WEBEXT.loginState.getProfileType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetProfileType() (ret js.Promise[ProfileType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetProfileType(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetSessionState returns true if the function "WEBEXT.loginState.getSessionState" exists.
func HasFuncGetSessionState() bool {
	return js.True == bindings.HasFuncGetSessionState()
}

// FuncGetSessionState returns the function "WEBEXT.loginState.getSessionState".
func FuncGetSessionState() (fn js.Func[func() js.Promise[SessionState]]) {
	bindings.FuncGetSessionState(
		js.Pointer(&fn),
	)
	return
}

// GetSessionState calls the function "WEBEXT.loginState.getSessionState" directly.
func GetSessionState() (ret js.Promise[SessionState]) {
	bindings.CallGetSessionState(
		js.Pointer(&ret),
	)

	return
}

// TryGetSessionState calls the function "WEBEXT.loginState.getSessionState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSessionState() (ret js.Promise[SessionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSessionState(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnSessionStateChangedEventCallbackFunc func(this js.Ref, sessionState SessionState) js.Ref

func (fn OnSessionStateChangedEventCallbackFunc) Register() js.Func[func(sessionState SessionState)] {
	return js.RegisterCallback[func(sessionState SessionState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSessionStateChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		SessionState(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnSessionStateChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, sessionState SessionState) js.Ref
	Arg T
}

func (cb *OnSessionStateChangedEventCallback[T]) Register() js.Func[func(sessionState SessionState)] {
	return js.RegisterCallback[func(sessionState SessionState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSessionStateChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		SessionState(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnSessionStateChanged returns true if the function "WEBEXT.loginState.onSessionStateChanged.addListener" exists.
func HasFuncOnSessionStateChanged() bool {
	return js.True == bindings.HasFuncOnSessionStateChanged()
}

// FuncOnSessionStateChanged returns the function "WEBEXT.loginState.onSessionStateChanged.addListener".
func FuncOnSessionStateChanged() (fn js.Func[func(callback js.Func[func(sessionState SessionState)])]) {
	bindings.FuncOnSessionStateChanged(
		js.Pointer(&fn),
	)
	return
}

// OnSessionStateChanged calls the function "WEBEXT.loginState.onSessionStateChanged.addListener" directly.
func OnSessionStateChanged(callback js.Func[func(sessionState SessionState)]) (ret js.Void) {
	bindings.CallOnSessionStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSessionStateChanged calls the function "WEBEXT.loginState.onSessionStateChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSessionStateChanged(callback js.Func[func(sessionState SessionState)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSessionStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSessionStateChanged returns true if the function "WEBEXT.loginState.onSessionStateChanged.removeListener" exists.
func HasFuncOffSessionStateChanged() bool {
	return js.True == bindings.HasFuncOffSessionStateChanged()
}

// FuncOffSessionStateChanged returns the function "WEBEXT.loginState.onSessionStateChanged.removeListener".
func FuncOffSessionStateChanged() (fn js.Func[func(callback js.Func[func(sessionState SessionState)])]) {
	bindings.FuncOffSessionStateChanged(
		js.Pointer(&fn),
	)
	return
}

// OffSessionStateChanged calls the function "WEBEXT.loginState.onSessionStateChanged.removeListener" directly.
func OffSessionStateChanged(callback js.Func[func(sessionState SessionState)]) (ret js.Void) {
	bindings.CallOffSessionStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSessionStateChanged calls the function "WEBEXT.loginState.onSessionStateChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSessionStateChanged(callback js.Func[func(sessionState SessionState)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSessionStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSessionStateChanged returns true if the function "WEBEXT.loginState.onSessionStateChanged.hasListener" exists.
func HasFuncHasOnSessionStateChanged() bool {
	return js.True == bindings.HasFuncHasOnSessionStateChanged()
}

// FuncHasOnSessionStateChanged returns the function "WEBEXT.loginState.onSessionStateChanged.hasListener".
func FuncHasOnSessionStateChanged() (fn js.Func[func(callback js.Func[func(sessionState SessionState)]) bool]) {
	bindings.FuncHasOnSessionStateChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnSessionStateChanged calls the function "WEBEXT.loginState.onSessionStateChanged.hasListener" directly.
func HasOnSessionStateChanged(callback js.Func[func(sessionState SessionState)]) (ret bool) {
	bindings.CallHasOnSessionStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSessionStateChanged calls the function "WEBEXT.loginState.onSessionStateChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSessionStateChanged(callback js.Func[func(sessionState SessionState)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSessionStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}
