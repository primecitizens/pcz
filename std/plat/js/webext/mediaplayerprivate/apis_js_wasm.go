// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package mediaplayerprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/mediaplayerprivate/bindings"
)

type OnNextTrackEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnNextTrackEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnNextTrackEventCallbackFunc) DispatchCallback(
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

type OnNextTrackEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnNextTrackEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnNextTrackEventCallback[T]) DispatchCallback(
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

// HasFuncOnNextTrack returns true if the function "WEBEXT.mediaPlayerPrivate.onNextTrack.addListener" exists.
func HasFuncOnNextTrack() bool {
	return js.True == bindings.HasFuncOnNextTrack()
}

// FuncOnNextTrack returns the function "WEBEXT.mediaPlayerPrivate.onNextTrack.addListener".
func FuncOnNextTrack() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnNextTrack(
		js.Pointer(&fn),
	)
	return
}

// OnNextTrack calls the function "WEBEXT.mediaPlayerPrivate.onNextTrack.addListener" directly.
func OnNextTrack(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnNextTrack(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnNextTrack calls the function "WEBEXT.mediaPlayerPrivate.onNextTrack.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnNextTrack(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnNextTrack(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffNextTrack returns true if the function "WEBEXT.mediaPlayerPrivate.onNextTrack.removeListener" exists.
func HasFuncOffNextTrack() bool {
	return js.True == bindings.HasFuncOffNextTrack()
}

// FuncOffNextTrack returns the function "WEBEXT.mediaPlayerPrivate.onNextTrack.removeListener".
func FuncOffNextTrack() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffNextTrack(
		js.Pointer(&fn),
	)
	return
}

// OffNextTrack calls the function "WEBEXT.mediaPlayerPrivate.onNextTrack.removeListener" directly.
func OffNextTrack(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffNextTrack(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffNextTrack calls the function "WEBEXT.mediaPlayerPrivate.onNextTrack.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffNextTrack(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffNextTrack(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnNextTrack returns true if the function "WEBEXT.mediaPlayerPrivate.onNextTrack.hasListener" exists.
func HasFuncHasOnNextTrack() bool {
	return js.True == bindings.HasFuncHasOnNextTrack()
}

// FuncHasOnNextTrack returns the function "WEBEXT.mediaPlayerPrivate.onNextTrack.hasListener".
func FuncHasOnNextTrack() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnNextTrack(
		js.Pointer(&fn),
	)
	return
}

// HasOnNextTrack calls the function "WEBEXT.mediaPlayerPrivate.onNextTrack.hasListener" directly.
func HasOnNextTrack(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnNextTrack(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnNextTrack calls the function "WEBEXT.mediaPlayerPrivate.onNextTrack.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnNextTrack(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnNextTrack(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPrevTrackEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnPrevTrackEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPrevTrackEventCallbackFunc) DispatchCallback(
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

type OnPrevTrackEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnPrevTrackEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPrevTrackEventCallback[T]) DispatchCallback(
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

// HasFuncOnPrevTrack returns true if the function "WEBEXT.mediaPlayerPrivate.onPrevTrack.addListener" exists.
func HasFuncOnPrevTrack() bool {
	return js.True == bindings.HasFuncOnPrevTrack()
}

// FuncOnPrevTrack returns the function "WEBEXT.mediaPlayerPrivate.onPrevTrack.addListener".
func FuncOnPrevTrack() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnPrevTrack(
		js.Pointer(&fn),
	)
	return
}

// OnPrevTrack calls the function "WEBEXT.mediaPlayerPrivate.onPrevTrack.addListener" directly.
func OnPrevTrack(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnPrevTrack(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPrevTrack calls the function "WEBEXT.mediaPlayerPrivate.onPrevTrack.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPrevTrack(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPrevTrack(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPrevTrack returns true if the function "WEBEXT.mediaPlayerPrivate.onPrevTrack.removeListener" exists.
func HasFuncOffPrevTrack() bool {
	return js.True == bindings.HasFuncOffPrevTrack()
}

// FuncOffPrevTrack returns the function "WEBEXT.mediaPlayerPrivate.onPrevTrack.removeListener".
func FuncOffPrevTrack() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffPrevTrack(
		js.Pointer(&fn),
	)
	return
}

// OffPrevTrack calls the function "WEBEXT.mediaPlayerPrivate.onPrevTrack.removeListener" directly.
func OffPrevTrack(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffPrevTrack(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPrevTrack calls the function "WEBEXT.mediaPlayerPrivate.onPrevTrack.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPrevTrack(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPrevTrack(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPrevTrack returns true if the function "WEBEXT.mediaPlayerPrivate.onPrevTrack.hasListener" exists.
func HasFuncHasOnPrevTrack() bool {
	return js.True == bindings.HasFuncHasOnPrevTrack()
}

// FuncHasOnPrevTrack returns the function "WEBEXT.mediaPlayerPrivate.onPrevTrack.hasListener".
func FuncHasOnPrevTrack() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnPrevTrack(
		js.Pointer(&fn),
	)
	return
}

// HasOnPrevTrack calls the function "WEBEXT.mediaPlayerPrivate.onPrevTrack.hasListener" directly.
func HasOnPrevTrack(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnPrevTrack(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPrevTrack calls the function "WEBEXT.mediaPlayerPrivate.onPrevTrack.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPrevTrack(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPrevTrack(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnTogglePlayStateEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnTogglePlayStateEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnTogglePlayStateEventCallbackFunc) DispatchCallback(
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

type OnTogglePlayStateEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnTogglePlayStateEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnTogglePlayStateEventCallback[T]) DispatchCallback(
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

// HasFuncOnTogglePlayState returns true if the function "WEBEXT.mediaPlayerPrivate.onTogglePlayState.addListener" exists.
func HasFuncOnTogglePlayState() bool {
	return js.True == bindings.HasFuncOnTogglePlayState()
}

// FuncOnTogglePlayState returns the function "WEBEXT.mediaPlayerPrivate.onTogglePlayState.addListener".
func FuncOnTogglePlayState() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnTogglePlayState(
		js.Pointer(&fn),
	)
	return
}

// OnTogglePlayState calls the function "WEBEXT.mediaPlayerPrivate.onTogglePlayState.addListener" directly.
func OnTogglePlayState(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnTogglePlayState(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnTogglePlayState calls the function "WEBEXT.mediaPlayerPrivate.onTogglePlayState.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTogglePlayState(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTogglePlayState(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffTogglePlayState returns true if the function "WEBEXT.mediaPlayerPrivate.onTogglePlayState.removeListener" exists.
func HasFuncOffTogglePlayState() bool {
	return js.True == bindings.HasFuncOffTogglePlayState()
}

// FuncOffTogglePlayState returns the function "WEBEXT.mediaPlayerPrivate.onTogglePlayState.removeListener".
func FuncOffTogglePlayState() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffTogglePlayState(
		js.Pointer(&fn),
	)
	return
}

// OffTogglePlayState calls the function "WEBEXT.mediaPlayerPrivate.onTogglePlayState.removeListener" directly.
func OffTogglePlayState(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffTogglePlayState(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffTogglePlayState calls the function "WEBEXT.mediaPlayerPrivate.onTogglePlayState.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffTogglePlayState(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffTogglePlayState(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnTogglePlayState returns true if the function "WEBEXT.mediaPlayerPrivate.onTogglePlayState.hasListener" exists.
func HasFuncHasOnTogglePlayState() bool {
	return js.True == bindings.HasFuncHasOnTogglePlayState()
}

// FuncHasOnTogglePlayState returns the function "WEBEXT.mediaPlayerPrivate.onTogglePlayState.hasListener".
func FuncHasOnTogglePlayState() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnTogglePlayState(
		js.Pointer(&fn),
	)
	return
}

// HasOnTogglePlayState calls the function "WEBEXT.mediaPlayerPrivate.onTogglePlayState.hasListener" directly.
func HasOnTogglePlayState(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnTogglePlayState(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnTogglePlayState calls the function "WEBEXT.mediaPlayerPrivate.onTogglePlayState.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnTogglePlayState(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnTogglePlayState(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}
