// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package idle

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/idle/bindings"
)

type IdleState uint32

const (
	_ IdleState = iota

	IdleState_ACTIVE
	IdleState_IDLE
	IdleState_LOCKED
)

func (IdleState) FromRef(str js.Ref) IdleState {
	return IdleState(bindings.ConstOfIdleState(str))
}

func (x IdleState) String() (string, bool) {
	switch x {
	case IdleState_ACTIVE:
		return "active", true
	case IdleState_IDLE:
		return "idle", true
	case IdleState_LOCKED:
		return "locked", true
	default:
		return "", false
	}
}

// HasFuncGetAutoLockDelay returns true if the function "WEBEXT.idle.getAutoLockDelay" exists.
func HasFuncGetAutoLockDelay() bool {
	return js.True == bindings.HasFuncGetAutoLockDelay()
}

// FuncGetAutoLockDelay returns the function "WEBEXT.idle.getAutoLockDelay".
func FuncGetAutoLockDelay() (fn js.Func[func() js.Promise[js.BigInt[int64]]]) {
	bindings.FuncGetAutoLockDelay(
		js.Pointer(&fn),
	)
	return
}

// GetAutoLockDelay calls the function "WEBEXT.idle.getAutoLockDelay" directly.
func GetAutoLockDelay() (ret js.Promise[js.BigInt[int64]]) {
	bindings.CallGetAutoLockDelay(
		js.Pointer(&ret),
	)

	return
}

// TryGetAutoLockDelay calls the function "WEBEXT.idle.getAutoLockDelay"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAutoLockDelay() (ret js.Promise[js.BigInt[int64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAutoLockDelay(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnStateChangedEventCallbackFunc func(this js.Ref, newState IdleState) js.Ref

func (fn OnStateChangedEventCallbackFunc) Register() js.Func[func(newState IdleState)] {
	return js.RegisterCallback[func(newState IdleState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnStateChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		IdleState(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnStateChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, newState IdleState) js.Ref
	Arg T
}

func (cb *OnStateChangedEventCallback[T]) Register() js.Func[func(newState IdleState)] {
	return js.RegisterCallback[func(newState IdleState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnStateChangedEventCallback[T]) DispatchCallback(
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

		IdleState(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnStateChanged returns true if the function "WEBEXT.idle.onStateChanged.addListener" exists.
func HasFuncOnStateChanged() bool {
	return js.True == bindings.HasFuncOnStateChanged()
}

// FuncOnStateChanged returns the function "WEBEXT.idle.onStateChanged.addListener".
func FuncOnStateChanged() (fn js.Func[func(callback js.Func[func(newState IdleState)])]) {
	bindings.FuncOnStateChanged(
		js.Pointer(&fn),
	)
	return
}

// OnStateChanged calls the function "WEBEXT.idle.onStateChanged.addListener" directly.
func OnStateChanged(callback js.Func[func(newState IdleState)]) (ret js.Void) {
	bindings.CallOnStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnStateChanged calls the function "WEBEXT.idle.onStateChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnStateChanged(callback js.Func[func(newState IdleState)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffStateChanged returns true if the function "WEBEXT.idle.onStateChanged.removeListener" exists.
func HasFuncOffStateChanged() bool {
	return js.True == bindings.HasFuncOffStateChanged()
}

// FuncOffStateChanged returns the function "WEBEXT.idle.onStateChanged.removeListener".
func FuncOffStateChanged() (fn js.Func[func(callback js.Func[func(newState IdleState)])]) {
	bindings.FuncOffStateChanged(
		js.Pointer(&fn),
	)
	return
}

// OffStateChanged calls the function "WEBEXT.idle.onStateChanged.removeListener" directly.
func OffStateChanged(callback js.Func[func(newState IdleState)]) (ret js.Void) {
	bindings.CallOffStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffStateChanged calls the function "WEBEXT.idle.onStateChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffStateChanged(callback js.Func[func(newState IdleState)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnStateChanged returns true if the function "WEBEXT.idle.onStateChanged.hasListener" exists.
func HasFuncHasOnStateChanged() bool {
	return js.True == bindings.HasFuncHasOnStateChanged()
}

// FuncHasOnStateChanged returns the function "WEBEXT.idle.onStateChanged.hasListener".
func FuncHasOnStateChanged() (fn js.Func[func(callback js.Func[func(newState IdleState)]) bool]) {
	bindings.FuncHasOnStateChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnStateChanged calls the function "WEBEXT.idle.onStateChanged.hasListener" directly.
func HasOnStateChanged(callback js.Func[func(newState IdleState)]) (ret bool) {
	bindings.CallHasOnStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnStateChanged calls the function "WEBEXT.idle.onStateChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnStateChanged(callback js.Func[func(newState IdleState)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncQueryState returns true if the function "WEBEXT.idle.queryState" exists.
func HasFuncQueryState() bool {
	return js.True == bindings.HasFuncQueryState()
}

// FuncQueryState returns the function "WEBEXT.idle.queryState".
func FuncQueryState() (fn js.Func[func(detectionIntervalInSeconds int64) js.Promise[IdleState]]) {
	bindings.FuncQueryState(
		js.Pointer(&fn),
	)
	return
}

// QueryState calls the function "WEBEXT.idle.queryState" directly.
func QueryState(detectionIntervalInSeconds int64) (ret js.Promise[IdleState]) {
	bindings.CallQueryState(
		js.Pointer(&ret),
		float64(detectionIntervalInSeconds),
	)

	return
}

// TryQueryState calls the function "WEBEXT.idle.queryState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryQueryState(detectionIntervalInSeconds int64) (ret js.Promise[IdleState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryQueryState(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(detectionIntervalInSeconds),
	)

	return
}

// HasFuncSetDetectionInterval returns true if the function "WEBEXT.idle.setDetectionInterval" exists.
func HasFuncSetDetectionInterval() bool {
	return js.True == bindings.HasFuncSetDetectionInterval()
}

// FuncSetDetectionInterval returns the function "WEBEXT.idle.setDetectionInterval".
func FuncSetDetectionInterval() (fn js.Func[func(intervalInSeconds int64)]) {
	bindings.FuncSetDetectionInterval(
		js.Pointer(&fn),
	)
	return
}

// SetDetectionInterval calls the function "WEBEXT.idle.setDetectionInterval" directly.
func SetDetectionInterval(intervalInSeconds int64) (ret js.Void) {
	bindings.CallSetDetectionInterval(
		js.Pointer(&ret),
		float64(intervalInSeconds),
	)

	return
}

// TrySetDetectionInterval calls the function "WEBEXT.idle.setDetectionInterval"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetDetectionInterval(intervalInSeconds int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetDetectionInterval(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(intervalInSeconds),
	)

	return
}
