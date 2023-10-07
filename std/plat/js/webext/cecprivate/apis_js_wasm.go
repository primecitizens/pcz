// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package cecprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/cecprivate/bindings"
)

type ChangePowerStateCallbackFunc func(this js.Ref) js.Ref

func (fn ChangePowerStateCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ChangePowerStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ChangePowerStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ChangePowerStateCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ChangePowerStateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DisplayCecPowerState uint32

const (
	_ DisplayCecPowerState = iota

	DisplayCecPowerState_ERROR
	DisplayCecPowerState_ADAPTER_NOT_CONFIGURED
	DisplayCecPowerState_NO_DEVICE
	DisplayCecPowerState_ON
	DisplayCecPowerState_STANDBY
	DisplayCecPowerState_TRANSITIONING_TO_ON
	DisplayCecPowerState_TRANSITIONING_TO_STANDBY
	DisplayCecPowerState_UNKNOWN
)

func (DisplayCecPowerState) FromRef(str js.Ref) DisplayCecPowerState {
	return DisplayCecPowerState(bindings.ConstOfDisplayCecPowerState(str))
}

func (x DisplayCecPowerState) String() (string, bool) {
	switch x {
	case DisplayCecPowerState_ERROR:
		return "error", true
	case DisplayCecPowerState_ADAPTER_NOT_CONFIGURED:
		return "adapterNotConfigured", true
	case DisplayCecPowerState_NO_DEVICE:
		return "noDevice", true
	case DisplayCecPowerState_ON:
		return "on", true
	case DisplayCecPowerState_STANDBY:
		return "standby", true
	case DisplayCecPowerState_TRANSITIONING_TO_ON:
		return "transitioningToOn", true
	case DisplayCecPowerState_TRANSITIONING_TO_STANDBY:
		return "transitioningToStandby", true
	case DisplayCecPowerState_UNKNOWN:
		return "unknown", true
	default:
		return "", false
	}
}

type DisplayCecPowerStateCallbackFunc func(this js.Ref, powerStates js.Array[DisplayCecPowerState]) js.Ref

func (fn DisplayCecPowerStateCallbackFunc) Register() js.Func[func(powerStates js.Array[DisplayCecPowerState])] {
	return js.RegisterCallback[func(powerStates js.Array[DisplayCecPowerState])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DisplayCecPowerStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[DisplayCecPowerState]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DisplayCecPowerStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, powerStates js.Array[DisplayCecPowerState]) js.Ref
	Arg T
}

func (cb *DisplayCecPowerStateCallback[T]) Register() js.Func[func(powerStates js.Array[DisplayCecPowerState])] {
	return js.RegisterCallback[func(powerStates js.Array[DisplayCecPowerState])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DisplayCecPowerStateCallback[T]) DispatchCallback(
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

		js.Array[DisplayCecPowerState]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncQueryDisplayCecPowerState returns true if the function "WEBEXT.cecPrivate.queryDisplayCecPowerState" exists.
func HasFuncQueryDisplayCecPowerState() bool {
	return js.True == bindings.HasFuncQueryDisplayCecPowerState()
}

// FuncQueryDisplayCecPowerState returns the function "WEBEXT.cecPrivate.queryDisplayCecPowerState".
func FuncQueryDisplayCecPowerState() (fn js.Func[func() js.Promise[js.Array[DisplayCecPowerState]]]) {
	bindings.FuncQueryDisplayCecPowerState(
		js.Pointer(&fn),
	)
	return
}

// QueryDisplayCecPowerState calls the function "WEBEXT.cecPrivate.queryDisplayCecPowerState" directly.
func QueryDisplayCecPowerState() (ret js.Promise[js.Array[DisplayCecPowerState]]) {
	bindings.CallQueryDisplayCecPowerState(
		js.Pointer(&ret),
	)

	return
}

// TryQueryDisplayCecPowerState calls the function "WEBEXT.cecPrivate.queryDisplayCecPowerState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryQueryDisplayCecPowerState() (ret js.Promise[js.Array[DisplayCecPowerState]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryQueryDisplayCecPowerState(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSendStandBy returns true if the function "WEBEXT.cecPrivate.sendStandBy" exists.
func HasFuncSendStandBy() bool {
	return js.True == bindings.HasFuncSendStandBy()
}

// FuncSendStandBy returns the function "WEBEXT.cecPrivate.sendStandBy".
func FuncSendStandBy() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncSendStandBy(
		js.Pointer(&fn),
	)
	return
}

// SendStandBy calls the function "WEBEXT.cecPrivate.sendStandBy" directly.
func SendStandBy() (ret js.Promise[js.Void]) {
	bindings.CallSendStandBy(
		js.Pointer(&ret),
	)

	return
}

// TrySendStandBy calls the function "WEBEXT.cecPrivate.sendStandBy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendStandBy() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendStandBy(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSendWakeUp returns true if the function "WEBEXT.cecPrivate.sendWakeUp" exists.
func HasFuncSendWakeUp() bool {
	return js.True == bindings.HasFuncSendWakeUp()
}

// FuncSendWakeUp returns the function "WEBEXT.cecPrivate.sendWakeUp".
func FuncSendWakeUp() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncSendWakeUp(
		js.Pointer(&fn),
	)
	return
}

// SendWakeUp calls the function "WEBEXT.cecPrivate.sendWakeUp" directly.
func SendWakeUp() (ret js.Promise[js.Void]) {
	bindings.CallSendWakeUp(
		js.Pointer(&ret),
	)

	return
}

// TrySendWakeUp calls the function "WEBEXT.cecPrivate.sendWakeUp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendWakeUp() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendWakeUp(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
