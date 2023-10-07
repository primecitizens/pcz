// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package extensionoptionsinternal

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/extensionoptionsinternal/bindings"
)

type PreferredSizeChangedOptions struct {
	// Width is "PreferredSizeChangedOptions.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width float64
	// Height is "PreferredSizeChangedOptions.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height float64

	FFI_USE_Width  bool // for Width.
	FFI_USE_Height bool // for Height.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PreferredSizeChangedOptions with all fields set.
func (p PreferredSizeChangedOptions) FromRef(ref js.Ref) PreferredSizeChangedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PreferredSizeChangedOptions in the application heap.
func (p PreferredSizeChangedOptions) New() js.Ref {
	return bindings.PreferredSizeChangedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PreferredSizeChangedOptions) UpdateFrom(ref js.Ref) {
	bindings.PreferredSizeChangedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PreferredSizeChangedOptions) Update(ref js.Ref) {
	bindings.PreferredSizeChangedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PreferredSizeChangedOptions) FreeMembers(recursive bool) {
}

type SizeChangedOptions struct {
	// OldWidth is "SizeChangedOptions.oldWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_OldWidth MUST be set to true to make this field effective.
	OldWidth int32
	// OldHeight is "SizeChangedOptions.oldHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_OldHeight MUST be set to true to make this field effective.
	OldHeight int32
	// NewWidth is "SizeChangedOptions.newWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_NewWidth MUST be set to true to make this field effective.
	NewWidth int32
	// NewHeight is "SizeChangedOptions.newHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_NewHeight MUST be set to true to make this field effective.
	NewHeight int32

	FFI_USE_OldWidth  bool // for OldWidth.
	FFI_USE_OldHeight bool // for OldHeight.
	FFI_USE_NewWidth  bool // for NewWidth.
	FFI_USE_NewHeight bool // for NewHeight.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SizeChangedOptions with all fields set.
func (p SizeChangedOptions) FromRef(ref js.Ref) SizeChangedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SizeChangedOptions in the application heap.
func (p SizeChangedOptions) New() js.Ref {
	return bindings.SizeChangedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SizeChangedOptions) UpdateFrom(ref js.Ref) {
	bindings.SizeChangedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SizeChangedOptions) Update(ref js.Ref) {
	bindings.SizeChangedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SizeChangedOptions) FreeMembers(recursive bool) {
}

type OnCloseEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnCloseEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCloseEventCallbackFunc) DispatchCallback(
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

type OnCloseEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnCloseEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCloseEventCallback[T]) DispatchCallback(
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

// HasFuncOnClose returns true if the function "WEBEXT.extensionOptionsInternal.onClose.addListener" exists.
func HasFuncOnClose() bool {
	return js.True == bindings.HasFuncOnClose()
}

// FuncOnClose returns the function "WEBEXT.extensionOptionsInternal.onClose.addListener".
func FuncOnClose() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnClose(
		js.Pointer(&fn),
	)
	return
}

// OnClose calls the function "WEBEXT.extensionOptionsInternal.onClose.addListener" directly.
func OnClose(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnClose(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnClose calls the function "WEBEXT.extensionOptionsInternal.onClose.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnClose(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnClose(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffClose returns true if the function "WEBEXT.extensionOptionsInternal.onClose.removeListener" exists.
func HasFuncOffClose() bool {
	return js.True == bindings.HasFuncOffClose()
}

// FuncOffClose returns the function "WEBEXT.extensionOptionsInternal.onClose.removeListener".
func FuncOffClose() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffClose(
		js.Pointer(&fn),
	)
	return
}

// OffClose calls the function "WEBEXT.extensionOptionsInternal.onClose.removeListener" directly.
func OffClose(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffClose(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffClose calls the function "WEBEXT.extensionOptionsInternal.onClose.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffClose(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffClose(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnClose returns true if the function "WEBEXT.extensionOptionsInternal.onClose.hasListener" exists.
func HasFuncHasOnClose() bool {
	return js.True == bindings.HasFuncHasOnClose()
}

// FuncHasOnClose returns the function "WEBEXT.extensionOptionsInternal.onClose.hasListener".
func FuncHasOnClose() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnClose(
		js.Pointer(&fn),
	)
	return
}

// HasOnClose calls the function "WEBEXT.extensionOptionsInternal.onClose.hasListener" directly.
func HasOnClose(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnClose(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnClose calls the function "WEBEXT.extensionOptionsInternal.onClose.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnClose(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnClose(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnLoadEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnLoadEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnLoadEventCallbackFunc) DispatchCallback(
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

type OnLoadEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnLoadEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnLoadEventCallback[T]) DispatchCallback(
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

// HasFuncOnLoad returns true if the function "WEBEXT.extensionOptionsInternal.onLoad.addListener" exists.
func HasFuncOnLoad() bool {
	return js.True == bindings.HasFuncOnLoad()
}

// FuncOnLoad returns the function "WEBEXT.extensionOptionsInternal.onLoad.addListener".
func FuncOnLoad() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnLoad(
		js.Pointer(&fn),
	)
	return
}

// OnLoad calls the function "WEBEXT.extensionOptionsInternal.onLoad.addListener" directly.
func OnLoad(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnLoad(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnLoad calls the function "WEBEXT.extensionOptionsInternal.onLoad.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnLoad(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnLoad(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffLoad returns true if the function "WEBEXT.extensionOptionsInternal.onLoad.removeListener" exists.
func HasFuncOffLoad() bool {
	return js.True == bindings.HasFuncOffLoad()
}

// FuncOffLoad returns the function "WEBEXT.extensionOptionsInternal.onLoad.removeListener".
func FuncOffLoad() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffLoad(
		js.Pointer(&fn),
	)
	return
}

// OffLoad calls the function "WEBEXT.extensionOptionsInternal.onLoad.removeListener" directly.
func OffLoad(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffLoad(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffLoad calls the function "WEBEXT.extensionOptionsInternal.onLoad.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffLoad(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffLoad(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnLoad returns true if the function "WEBEXT.extensionOptionsInternal.onLoad.hasListener" exists.
func HasFuncHasOnLoad() bool {
	return js.True == bindings.HasFuncHasOnLoad()
}

// FuncHasOnLoad returns the function "WEBEXT.extensionOptionsInternal.onLoad.hasListener".
func FuncHasOnLoad() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnLoad(
		js.Pointer(&fn),
	)
	return
}

// HasOnLoad calls the function "WEBEXT.extensionOptionsInternal.onLoad.hasListener" directly.
func HasOnLoad(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnLoad(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnLoad calls the function "WEBEXT.extensionOptionsInternal.onLoad.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnLoad(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnLoad(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPreferredSizeChangedEventCallbackFunc func(this js.Ref, options *PreferredSizeChangedOptions) js.Ref

func (fn OnPreferredSizeChangedEventCallbackFunc) Register() js.Func[func(options *PreferredSizeChangedOptions)] {
	return js.RegisterCallback[func(options *PreferredSizeChangedOptions)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPreferredSizeChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PreferredSizeChangedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnPreferredSizeChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, options *PreferredSizeChangedOptions) js.Ref
	Arg T
}

func (cb *OnPreferredSizeChangedEventCallback[T]) Register() js.Func[func(options *PreferredSizeChangedOptions)] {
	return js.RegisterCallback[func(options *PreferredSizeChangedOptions)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPreferredSizeChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PreferredSizeChangedOptions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnPreferredSizeChanged returns true if the function "WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.addListener" exists.
func HasFuncOnPreferredSizeChanged() bool {
	return js.True == bindings.HasFuncOnPreferredSizeChanged()
}

// FuncOnPreferredSizeChanged returns the function "WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.addListener".
func FuncOnPreferredSizeChanged() (fn js.Func[func(callback js.Func[func(options *PreferredSizeChangedOptions)])]) {
	bindings.FuncOnPreferredSizeChanged(
		js.Pointer(&fn),
	)
	return
}

// OnPreferredSizeChanged calls the function "WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.addListener" directly.
func OnPreferredSizeChanged(callback js.Func[func(options *PreferredSizeChangedOptions)]) (ret js.Void) {
	bindings.CallOnPreferredSizeChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPreferredSizeChanged calls the function "WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPreferredSizeChanged(callback js.Func[func(options *PreferredSizeChangedOptions)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPreferredSizeChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPreferredSizeChanged returns true if the function "WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.removeListener" exists.
func HasFuncOffPreferredSizeChanged() bool {
	return js.True == bindings.HasFuncOffPreferredSizeChanged()
}

// FuncOffPreferredSizeChanged returns the function "WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.removeListener".
func FuncOffPreferredSizeChanged() (fn js.Func[func(callback js.Func[func(options *PreferredSizeChangedOptions)])]) {
	bindings.FuncOffPreferredSizeChanged(
		js.Pointer(&fn),
	)
	return
}

// OffPreferredSizeChanged calls the function "WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.removeListener" directly.
func OffPreferredSizeChanged(callback js.Func[func(options *PreferredSizeChangedOptions)]) (ret js.Void) {
	bindings.CallOffPreferredSizeChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPreferredSizeChanged calls the function "WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPreferredSizeChanged(callback js.Func[func(options *PreferredSizeChangedOptions)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPreferredSizeChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPreferredSizeChanged returns true if the function "WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.hasListener" exists.
func HasFuncHasOnPreferredSizeChanged() bool {
	return js.True == bindings.HasFuncHasOnPreferredSizeChanged()
}

// FuncHasOnPreferredSizeChanged returns the function "WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.hasListener".
func FuncHasOnPreferredSizeChanged() (fn js.Func[func(callback js.Func[func(options *PreferredSizeChangedOptions)]) bool]) {
	bindings.FuncHasOnPreferredSizeChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnPreferredSizeChanged calls the function "WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.hasListener" directly.
func HasOnPreferredSizeChanged(callback js.Func[func(options *PreferredSizeChangedOptions)]) (ret bool) {
	bindings.CallHasOnPreferredSizeChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPreferredSizeChanged calls the function "WEBEXT.extensionOptionsInternal.onPreferredSizeChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPreferredSizeChanged(callback js.Func[func(options *PreferredSizeChangedOptions)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPreferredSizeChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}
