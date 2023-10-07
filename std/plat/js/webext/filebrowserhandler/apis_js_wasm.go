// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package filebrowserhandler

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/filebrowserhandler/bindings"
)

type FileHandlerExecuteEventDetails struct {
	// Entries is "FileHandlerExecuteEventDetails.entries"
	//
	// Required
	Entries js.Array[js.Any]
	// TabId is "FileHandlerExecuteEventDetails.tab_id"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int64

	FFI_USE_TabId bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileHandlerExecuteEventDetails with all fields set.
func (p FileHandlerExecuteEventDetails) FromRef(ref js.Ref) FileHandlerExecuteEventDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileHandlerExecuteEventDetails in the application heap.
func (p FileHandlerExecuteEventDetails) New() js.Ref {
	return bindings.FileHandlerExecuteEventDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FileHandlerExecuteEventDetails) UpdateFrom(ref js.Ref) {
	bindings.FileHandlerExecuteEventDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileHandlerExecuteEventDetails) Update(ref js.Ref) {
	bindings.FileHandlerExecuteEventDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileHandlerExecuteEventDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Entries.Ref(),
	)
	p.Entries = p.Entries.FromRef(js.Undefined)
}

type OnExecuteEventCallbackFunc func(this js.Ref, id js.String, details *FileHandlerExecuteEventDetails) js.Ref

func (fn OnExecuteEventCallbackFunc) Register() js.Func[func(id js.String, details *FileHandlerExecuteEventDetails)] {
	return js.RegisterCallback[func(id js.String, details *FileHandlerExecuteEventDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnExecuteEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 FileHandlerExecuteEventDetails
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnExecuteEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, id js.String, details *FileHandlerExecuteEventDetails) js.Ref
	Arg T
}

func (cb *OnExecuteEventCallback[T]) Register() js.Func[func(id js.String, details *FileHandlerExecuteEventDetails)] {
	return js.RegisterCallback[func(id js.String, details *FileHandlerExecuteEventDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnExecuteEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 FileHandlerExecuteEventDetails
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnExecute returns true if the function "WEBEXT.fileBrowserHandler.onExecute.addListener" exists.
func HasFuncOnExecute() bool {
	return js.True == bindings.HasFuncOnExecute()
}

// FuncOnExecute returns the function "WEBEXT.fileBrowserHandler.onExecute.addListener".
func FuncOnExecute() (fn js.Func[func(callback js.Func[func(id js.String, details *FileHandlerExecuteEventDetails)])]) {
	bindings.FuncOnExecute(
		js.Pointer(&fn),
	)
	return
}

// OnExecute calls the function "WEBEXT.fileBrowserHandler.onExecute.addListener" directly.
func OnExecute(callback js.Func[func(id js.String, details *FileHandlerExecuteEventDetails)]) (ret js.Void) {
	bindings.CallOnExecute(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnExecute calls the function "WEBEXT.fileBrowserHandler.onExecute.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnExecute(callback js.Func[func(id js.String, details *FileHandlerExecuteEventDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnExecute(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffExecute returns true if the function "WEBEXT.fileBrowserHandler.onExecute.removeListener" exists.
func HasFuncOffExecute() bool {
	return js.True == bindings.HasFuncOffExecute()
}

// FuncOffExecute returns the function "WEBEXT.fileBrowserHandler.onExecute.removeListener".
func FuncOffExecute() (fn js.Func[func(callback js.Func[func(id js.String, details *FileHandlerExecuteEventDetails)])]) {
	bindings.FuncOffExecute(
		js.Pointer(&fn),
	)
	return
}

// OffExecute calls the function "WEBEXT.fileBrowserHandler.onExecute.removeListener" directly.
func OffExecute(callback js.Func[func(id js.String, details *FileHandlerExecuteEventDetails)]) (ret js.Void) {
	bindings.CallOffExecute(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffExecute calls the function "WEBEXT.fileBrowserHandler.onExecute.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffExecute(callback js.Func[func(id js.String, details *FileHandlerExecuteEventDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffExecute(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnExecute returns true if the function "WEBEXT.fileBrowserHandler.onExecute.hasListener" exists.
func HasFuncHasOnExecute() bool {
	return js.True == bindings.HasFuncHasOnExecute()
}

// FuncHasOnExecute returns the function "WEBEXT.fileBrowserHandler.onExecute.hasListener".
func FuncHasOnExecute() (fn js.Func[func(callback js.Func[func(id js.String, details *FileHandlerExecuteEventDetails)]) bool]) {
	bindings.FuncHasOnExecute(
		js.Pointer(&fn),
	)
	return
}

// HasOnExecute calls the function "WEBEXT.fileBrowserHandler.onExecute.hasListener" directly.
func HasOnExecute(callback js.Func[func(id js.String, details *FileHandlerExecuteEventDetails)]) (ret bool) {
	bindings.CallHasOnExecute(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnExecute calls the function "WEBEXT.fileBrowserHandler.onExecute.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnExecute(callback js.Func[func(id js.String, details *FileHandlerExecuteEventDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnExecute(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}
