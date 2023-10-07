// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package loginscreenui

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/loginscreenui/bindings"
)

type ShowOptions struct {
	// Url is "ShowOptions.url"
	//
	// Optional
	Url js.String
	// UserCanClose is "ShowOptions.userCanClose"
	//
	// Optional
	//
	// NOTE: FFI_USE_UserCanClose MUST be set to true to make this field effective.
	UserCanClose bool

	FFI_USE_UserCanClose bool // for UserCanClose.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ShowOptions with all fields set.
func (p ShowOptions) FromRef(ref js.Ref) ShowOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ShowOptions in the application heap.
func (p ShowOptions) New() js.Ref {
	return bindings.ShowOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ShowOptions) UpdateFrom(ref js.Ref) {
	bindings.ShowOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ShowOptions) Update(ref js.Ref) {
	bindings.ShowOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ShowOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
}

type SimpleCallbackFunc func(this js.Ref) js.Ref

func (fn SimpleCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SimpleCallbackFunc) DispatchCallback(
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

type SimpleCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *SimpleCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SimpleCallback[T]) DispatchCallback(
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

// HasFuncClose returns true if the function "WEBEXT.loginScreenUi.close" exists.
func HasFuncClose() bool {
	return js.True == bindings.HasFuncClose()
}

// FuncClose returns the function "WEBEXT.loginScreenUi.close".
func FuncClose() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncClose(
		js.Pointer(&fn),
	)
	return
}

// Close calls the function "WEBEXT.loginScreenUi.close" directly.
func Close() (ret js.Promise[js.Void]) {
	bindings.CallClose(
		js.Pointer(&ret),
	)

	return
}

// TryClose calls the function "WEBEXT.loginScreenUi.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClose() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClose(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncShow returns true if the function "WEBEXT.loginScreenUi.show" exists.
func HasFuncShow() bool {
	return js.True == bindings.HasFuncShow()
}

// FuncShow returns the function "WEBEXT.loginScreenUi.show".
func FuncShow() (fn js.Func[func(options ShowOptions) js.Promise[js.Void]]) {
	bindings.FuncShow(
		js.Pointer(&fn),
	)
	return
}

// Show calls the function "WEBEXT.loginScreenUi.show" directly.
func Show(options ShowOptions) (ret js.Promise[js.Void]) {
	bindings.CallShow(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryShow calls the function "WEBEXT.loginScreenUi.show"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShow(options ShowOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryShow(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}
