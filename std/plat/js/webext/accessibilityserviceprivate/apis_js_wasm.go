// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package accessibilityserviceprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/accessibilityserviceprivate/bindings"
)

type VoidCallbackFunc func(this js.Ref) js.Ref

func (fn VoidCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VoidCallbackFunc) DispatchCallback(
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

type VoidCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *VoidCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VoidCallback[T]) DispatchCallback(
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

// HasFuncSpeakSelectedText returns true if the function "WEBEXT.accessibilityServicePrivate.speakSelectedText" exists.
func HasFuncSpeakSelectedText() bool {
	return js.True == bindings.HasFuncSpeakSelectedText()
}

// FuncSpeakSelectedText returns the function "WEBEXT.accessibilityServicePrivate.speakSelectedText".
func FuncSpeakSelectedText() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncSpeakSelectedText(
		js.Pointer(&fn),
	)
	return
}

// SpeakSelectedText calls the function "WEBEXT.accessibilityServicePrivate.speakSelectedText" directly.
func SpeakSelectedText() (ret js.Promise[js.Void]) {
	bindings.CallSpeakSelectedText(
		js.Pointer(&ret),
	)

	return
}

// TrySpeakSelectedText calls the function "WEBEXT.accessibilityServicePrivate.speakSelectedText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySpeakSelectedText() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeakSelectedText(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
