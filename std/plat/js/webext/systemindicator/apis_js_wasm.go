// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package systemindicator

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/systemindicator/bindings"
)

type DoneCallbackFunc func(this js.Ref) js.Ref

func (fn DoneCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DoneCallbackFunc) DispatchCallback(
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

type DoneCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *DoneCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DoneCallback[T]) DispatchCallback(
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

type SetIconDetails struct {
	// Path is "SetIconDetails.path"
	//
	// Optional
	Path js.Any
	// ImageData is "SetIconDetails.imageData"
	//
	// Optional
	ImageData js.Any

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetIconDetails with all fields set.
func (p SetIconDetails) FromRef(ref js.Ref) SetIconDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetIconDetails in the application heap.
func (p SetIconDetails) New() js.Ref {
	return bindings.SetIconDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetIconDetails) UpdateFrom(ref js.Ref) {
	bindings.SetIconDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetIconDetails) Update(ref js.Ref) {
	bindings.SetIconDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetIconDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Path.Ref(),
		p.ImageData.Ref(),
	)
	p.Path = p.Path.FromRef(js.Undefined)
	p.ImageData = p.ImageData.FromRef(js.Undefined)
}

// HasFuncDisable returns true if the function "WEBEXT.systemIndicator.disable" exists.
func HasFuncDisable() bool {
	return js.True == bindings.HasFuncDisable()
}

// FuncDisable returns the function "WEBEXT.systemIndicator.disable".
func FuncDisable() (fn js.Func[func()]) {
	bindings.FuncDisable(
		js.Pointer(&fn),
	)
	return
}

// Disable calls the function "WEBEXT.systemIndicator.disable" directly.
func Disable() (ret js.Void) {
	bindings.CallDisable(
		js.Pointer(&ret),
	)

	return
}

// TryDisable calls the function "WEBEXT.systemIndicator.disable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDisable() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDisable(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncEnable returns true if the function "WEBEXT.systemIndicator.enable" exists.
func HasFuncEnable() bool {
	return js.True == bindings.HasFuncEnable()
}

// FuncEnable returns the function "WEBEXT.systemIndicator.enable".
func FuncEnable() (fn js.Func[func()]) {
	bindings.FuncEnable(
		js.Pointer(&fn),
	)
	return
}

// Enable calls the function "WEBEXT.systemIndicator.enable" directly.
func Enable() (ret js.Void) {
	bindings.CallEnable(
		js.Pointer(&ret),
	)

	return
}

// TryEnable calls the function "WEBEXT.systemIndicator.enable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEnable() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEnable(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnClickedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnClickedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnClickedEventCallbackFunc) DispatchCallback(
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

type OnClickedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnClickedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnClickedEventCallback[T]) DispatchCallback(
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

// HasFuncOnClicked returns true if the function "WEBEXT.systemIndicator.onClicked.addListener" exists.
func HasFuncOnClicked() bool {
	return js.True == bindings.HasFuncOnClicked()
}

// FuncOnClicked returns the function "WEBEXT.systemIndicator.onClicked.addListener".
func FuncOnClicked() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnClicked(
		js.Pointer(&fn),
	)
	return
}

// OnClicked calls the function "WEBEXT.systemIndicator.onClicked.addListener" directly.
func OnClicked(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnClicked calls the function "WEBEXT.systemIndicator.onClicked.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnClicked(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffClicked returns true if the function "WEBEXT.systemIndicator.onClicked.removeListener" exists.
func HasFuncOffClicked() bool {
	return js.True == bindings.HasFuncOffClicked()
}

// FuncOffClicked returns the function "WEBEXT.systemIndicator.onClicked.removeListener".
func FuncOffClicked() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffClicked(
		js.Pointer(&fn),
	)
	return
}

// OffClicked calls the function "WEBEXT.systemIndicator.onClicked.removeListener" directly.
func OffClicked(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffClicked calls the function "WEBEXT.systemIndicator.onClicked.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffClicked(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnClicked returns true if the function "WEBEXT.systemIndicator.onClicked.hasListener" exists.
func HasFuncHasOnClicked() bool {
	return js.True == bindings.HasFuncHasOnClicked()
}

// FuncHasOnClicked returns the function "WEBEXT.systemIndicator.onClicked.hasListener".
func FuncHasOnClicked() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnClicked(
		js.Pointer(&fn),
	)
	return
}

// HasOnClicked calls the function "WEBEXT.systemIndicator.onClicked.hasListener" directly.
func HasOnClicked(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnClicked calls the function "WEBEXT.systemIndicator.onClicked.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnClicked(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSetIcon returns true if the function "WEBEXT.systemIndicator.setIcon" exists.
func HasFuncSetIcon() bool {
	return js.True == bindings.HasFuncSetIcon()
}

// FuncSetIcon returns the function "WEBEXT.systemIndicator.setIcon".
func FuncSetIcon() (fn js.Func[func(details SetIconDetails) js.Promise[js.Void]]) {
	bindings.FuncSetIcon(
		js.Pointer(&fn),
	)
	return
}

// SetIcon calls the function "WEBEXT.systemIndicator.setIcon" directly.
func SetIcon(details SetIconDetails) (ret js.Promise[js.Void]) {
	bindings.CallSetIcon(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySetIcon calls the function "WEBEXT.systemIndicator.setIcon"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetIcon(details SetIconDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetIcon(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}
