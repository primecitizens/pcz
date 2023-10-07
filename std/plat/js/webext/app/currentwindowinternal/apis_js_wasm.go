// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package currentwindowinternal

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/app/currentwindowinternal/bindings"
)

type Bounds struct {
	// Left is "Bounds.left"
	//
	// Optional
	//
	// NOTE: FFI_USE_Left MUST be set to true to make this field effective.
	Left int32
	// Top is "Bounds.top"
	//
	// Optional
	//
	// NOTE: FFI_USE_Top MUST be set to true to make this field effective.
	Top int32
	// Width is "Bounds.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int32
	// Height is "Bounds.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int32

	FFI_USE_Left   bool // for Left.
	FFI_USE_Top    bool // for Top.
	FFI_USE_Width  bool // for Width.
	FFI_USE_Height bool // for Height.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Bounds with all fields set.
func (p Bounds) FromRef(ref js.Ref) Bounds {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Bounds in the application heap.
func (p Bounds) New() js.Ref {
	return bindings.BoundsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Bounds) UpdateFrom(ref js.Ref) {
	bindings.BoundsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Bounds) Update(ref js.Ref) {
	bindings.BoundsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Bounds) FreeMembers(recursive bool) {
}

type RegionRect struct {
	// Left is "RegionRect.left"
	//
	// Optional
	//
	// NOTE: FFI_USE_Left MUST be set to true to make this field effective.
	Left int32
	// Top is "RegionRect.top"
	//
	// Optional
	//
	// NOTE: FFI_USE_Top MUST be set to true to make this field effective.
	Top int32
	// Width is "RegionRect.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int32
	// Height is "RegionRect.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int32

	FFI_USE_Left   bool // for Left.
	FFI_USE_Top    bool // for Top.
	FFI_USE_Width  bool // for Width.
	FFI_USE_Height bool // for Height.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RegionRect with all fields set.
func (p RegionRect) FromRef(ref js.Ref) RegionRect {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RegionRect in the application heap.
func (p RegionRect) New() js.Ref {
	return bindings.RegionRectJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RegionRect) UpdateFrom(ref js.Ref) {
	bindings.RegionRectJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RegionRect) Update(ref js.Ref) {
	bindings.RegionRectJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RegionRect) FreeMembers(recursive bool) {
}

type Region struct {
	// Rects is "Region.rects"
	//
	// Optional
	Rects js.Array[RegionRect]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Region with all fields set.
func (p Region) FromRef(ref js.Ref) Region {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Region in the application heap.
func (p Region) New() js.Ref {
	return bindings.RegionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Region) UpdateFrom(ref js.Ref) {
	bindings.RegionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Region) Update(ref js.Ref) {
	bindings.RegionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Region) FreeMembers(recursive bool) {
	js.Free(
		p.Rects.Ref(),
	)
	p.Rects = p.Rects.FromRef(js.Undefined)
}

type SizeConstraints struct {
	// MinWidth is "SizeConstraints.minWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_MinWidth MUST be set to true to make this field effective.
	MinWidth int32
	// MinHeight is "SizeConstraints.minHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_MinHeight MUST be set to true to make this field effective.
	MinHeight int32
	// MaxWidth is "SizeConstraints.maxWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxWidth MUST be set to true to make this field effective.
	MaxWidth int32
	// MaxHeight is "SizeConstraints.maxHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxHeight MUST be set to true to make this field effective.
	MaxHeight int32

	FFI_USE_MinWidth  bool // for MinWidth.
	FFI_USE_MinHeight bool // for MinHeight.
	FFI_USE_MaxWidth  bool // for MaxWidth.
	FFI_USE_MaxHeight bool // for MaxHeight.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SizeConstraints with all fields set.
func (p SizeConstraints) FromRef(ref js.Ref) SizeConstraints {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SizeConstraints in the application heap.
func (p SizeConstraints) New() js.Ref {
	return bindings.SizeConstraintsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SizeConstraints) UpdateFrom(ref js.Ref) {
	bindings.SizeConstraintsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SizeConstraints) Update(ref js.Ref) {
	bindings.SizeConstraintsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SizeConstraints) FreeMembers(recursive bool) {
}

// HasFuncClearAttention returns true if the function "WEBEXT.app.currentWindowInternal.clearAttention" exists.
func HasFuncClearAttention() bool {
	return js.True == bindings.HasFuncClearAttention()
}

// FuncClearAttention returns the function "WEBEXT.app.currentWindowInternal.clearAttention".
func FuncClearAttention() (fn js.Func[func()]) {
	bindings.FuncClearAttention(
		js.Pointer(&fn),
	)
	return
}

// ClearAttention calls the function "WEBEXT.app.currentWindowInternal.clearAttention" directly.
func ClearAttention() (ret js.Void) {
	bindings.CallClearAttention(
		js.Pointer(&ret),
	)

	return
}

// TryClearAttention calls the function "WEBEXT.app.currentWindowInternal.clearAttention"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClearAttention() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryClearAttention(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDrawAttention returns true if the function "WEBEXT.app.currentWindowInternal.drawAttention" exists.
func HasFuncDrawAttention() bool {
	return js.True == bindings.HasFuncDrawAttention()
}

// FuncDrawAttention returns the function "WEBEXT.app.currentWindowInternal.drawAttention".
func FuncDrawAttention() (fn js.Func[func()]) {
	bindings.FuncDrawAttention(
		js.Pointer(&fn),
	)
	return
}

// DrawAttention calls the function "WEBEXT.app.currentWindowInternal.drawAttention" directly.
func DrawAttention() (ret js.Void) {
	bindings.CallDrawAttention(
		js.Pointer(&ret),
	)

	return
}

// TryDrawAttention calls the function "WEBEXT.app.currentWindowInternal.drawAttention"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDrawAttention() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDrawAttention(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFocus returns true if the function "WEBEXT.app.currentWindowInternal.focus" exists.
func HasFuncFocus() bool {
	return js.True == bindings.HasFuncFocus()
}

// FuncFocus returns the function "WEBEXT.app.currentWindowInternal.focus".
func FuncFocus() (fn js.Func[func()]) {
	bindings.FuncFocus(
		js.Pointer(&fn),
	)
	return
}

// Focus calls the function "WEBEXT.app.currentWindowInternal.focus" directly.
func Focus() (ret js.Void) {
	bindings.CallFocus(
		js.Pointer(&ret),
	)

	return
}

// TryFocus calls the function "WEBEXT.app.currentWindowInternal.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryFocus() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFocus(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFullscreen returns true if the function "WEBEXT.app.currentWindowInternal.fullscreen" exists.
func HasFuncFullscreen() bool {
	return js.True == bindings.HasFuncFullscreen()
}

// FuncFullscreen returns the function "WEBEXT.app.currentWindowInternal.fullscreen".
func FuncFullscreen() (fn js.Func[func()]) {
	bindings.FuncFullscreen(
		js.Pointer(&fn),
	)
	return
}

// Fullscreen calls the function "WEBEXT.app.currentWindowInternal.fullscreen" directly.
func Fullscreen() (ret js.Void) {
	bindings.CallFullscreen(
		js.Pointer(&ret),
	)

	return
}

// TryFullscreen calls the function "WEBEXT.app.currentWindowInternal.fullscreen"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryFullscreen() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFullscreen(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncHide returns true if the function "WEBEXT.app.currentWindowInternal.hide" exists.
func HasFuncHide() bool {
	return js.True == bindings.HasFuncHide()
}

// FuncHide returns the function "WEBEXT.app.currentWindowInternal.hide".
func FuncHide() (fn js.Func[func()]) {
	bindings.FuncHide(
		js.Pointer(&fn),
	)
	return
}

// Hide calls the function "WEBEXT.app.currentWindowInternal.hide" directly.
func Hide() (ret js.Void) {
	bindings.CallHide(
		js.Pointer(&ret),
	)

	return
}

// TryHide calls the function "WEBEXT.app.currentWindowInternal.hide"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHide() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHide(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMaximize returns true if the function "WEBEXT.app.currentWindowInternal.maximize" exists.
func HasFuncMaximize() bool {
	return js.True == bindings.HasFuncMaximize()
}

// FuncMaximize returns the function "WEBEXT.app.currentWindowInternal.maximize".
func FuncMaximize() (fn js.Func[func()]) {
	bindings.FuncMaximize(
		js.Pointer(&fn),
	)
	return
}

// Maximize calls the function "WEBEXT.app.currentWindowInternal.maximize" directly.
func Maximize() (ret js.Void) {
	bindings.CallMaximize(
		js.Pointer(&ret),
	)

	return
}

// TryMaximize calls the function "WEBEXT.app.currentWindowInternal.maximize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMaximize() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMaximize(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMinimize returns true if the function "WEBEXT.app.currentWindowInternal.minimize" exists.
func HasFuncMinimize() bool {
	return js.True == bindings.HasFuncMinimize()
}

// FuncMinimize returns the function "WEBEXT.app.currentWindowInternal.minimize".
func FuncMinimize() (fn js.Func[func()]) {
	bindings.FuncMinimize(
		js.Pointer(&fn),
	)
	return
}

// Minimize calls the function "WEBEXT.app.currentWindowInternal.minimize" directly.
func Minimize() (ret js.Void) {
	bindings.CallMinimize(
		js.Pointer(&ret),
	)

	return
}

// TryMinimize calls the function "WEBEXT.app.currentWindowInternal.minimize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMinimize() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMinimize(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnAlphaEnabledChangedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnAlphaEnabledChangedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAlphaEnabledChangedEventCallbackFunc) DispatchCallback(
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

type OnAlphaEnabledChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnAlphaEnabledChangedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAlphaEnabledChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnAlphaEnabledChanged returns true if the function "WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.addListener" exists.
func HasFuncOnAlphaEnabledChanged() bool {
	return js.True == bindings.HasFuncOnAlphaEnabledChanged()
}

// FuncOnAlphaEnabledChanged returns the function "WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.addListener".
func FuncOnAlphaEnabledChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnAlphaEnabledChanged(
		js.Pointer(&fn),
	)
	return
}

// OnAlphaEnabledChanged calls the function "WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.addListener" directly.
func OnAlphaEnabledChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnAlphaEnabledChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAlphaEnabledChanged calls the function "WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAlphaEnabledChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAlphaEnabledChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAlphaEnabledChanged returns true if the function "WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.removeListener" exists.
func HasFuncOffAlphaEnabledChanged() bool {
	return js.True == bindings.HasFuncOffAlphaEnabledChanged()
}

// FuncOffAlphaEnabledChanged returns the function "WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.removeListener".
func FuncOffAlphaEnabledChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffAlphaEnabledChanged(
		js.Pointer(&fn),
	)
	return
}

// OffAlphaEnabledChanged calls the function "WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.removeListener" directly.
func OffAlphaEnabledChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffAlphaEnabledChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAlphaEnabledChanged calls the function "WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAlphaEnabledChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAlphaEnabledChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAlphaEnabledChanged returns true if the function "WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.hasListener" exists.
func HasFuncHasOnAlphaEnabledChanged() bool {
	return js.True == bindings.HasFuncHasOnAlphaEnabledChanged()
}

// FuncHasOnAlphaEnabledChanged returns the function "WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.hasListener".
func FuncHasOnAlphaEnabledChanged() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnAlphaEnabledChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnAlphaEnabledChanged calls the function "WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.hasListener" directly.
func HasOnAlphaEnabledChanged(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnAlphaEnabledChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAlphaEnabledChanged calls the function "WEBEXT.app.currentWindowInternal.onAlphaEnabledChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAlphaEnabledChanged(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAlphaEnabledChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnBoundsChangedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnBoundsChangedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnBoundsChangedEventCallbackFunc) DispatchCallback(
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

type OnBoundsChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnBoundsChangedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnBoundsChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnBoundsChanged returns true if the function "WEBEXT.app.currentWindowInternal.onBoundsChanged.addListener" exists.
func HasFuncOnBoundsChanged() bool {
	return js.True == bindings.HasFuncOnBoundsChanged()
}

// FuncOnBoundsChanged returns the function "WEBEXT.app.currentWindowInternal.onBoundsChanged.addListener".
func FuncOnBoundsChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// OnBoundsChanged calls the function "WEBEXT.app.currentWindowInternal.onBoundsChanged.addListener" directly.
func OnBoundsChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnBoundsChanged calls the function "WEBEXT.app.currentWindowInternal.onBoundsChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnBoundsChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffBoundsChanged returns true if the function "WEBEXT.app.currentWindowInternal.onBoundsChanged.removeListener" exists.
func HasFuncOffBoundsChanged() bool {
	return js.True == bindings.HasFuncOffBoundsChanged()
}

// FuncOffBoundsChanged returns the function "WEBEXT.app.currentWindowInternal.onBoundsChanged.removeListener".
func FuncOffBoundsChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// OffBoundsChanged calls the function "WEBEXT.app.currentWindowInternal.onBoundsChanged.removeListener" directly.
func OffBoundsChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffBoundsChanged calls the function "WEBEXT.app.currentWindowInternal.onBoundsChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffBoundsChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnBoundsChanged returns true if the function "WEBEXT.app.currentWindowInternal.onBoundsChanged.hasListener" exists.
func HasFuncHasOnBoundsChanged() bool {
	return js.True == bindings.HasFuncHasOnBoundsChanged()
}

// FuncHasOnBoundsChanged returns the function "WEBEXT.app.currentWindowInternal.onBoundsChanged.hasListener".
func FuncHasOnBoundsChanged() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnBoundsChanged calls the function "WEBEXT.app.currentWindowInternal.onBoundsChanged.hasListener" directly.
func HasOnBoundsChanged(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnBoundsChanged calls the function "WEBEXT.app.currentWindowInternal.onBoundsChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnBoundsChanged(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnClosedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnClosedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnClosedEventCallbackFunc) DispatchCallback(
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

type OnClosedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnClosedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnClosedEventCallback[T]) DispatchCallback(
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

// HasFuncOnClosed returns true if the function "WEBEXT.app.currentWindowInternal.onClosed.addListener" exists.
func HasFuncOnClosed() bool {
	return js.True == bindings.HasFuncOnClosed()
}

// FuncOnClosed returns the function "WEBEXT.app.currentWindowInternal.onClosed.addListener".
func FuncOnClosed() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnClosed(
		js.Pointer(&fn),
	)
	return
}

// OnClosed calls the function "WEBEXT.app.currentWindowInternal.onClosed.addListener" directly.
func OnClosed(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnClosed(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnClosed calls the function "WEBEXT.app.currentWindowInternal.onClosed.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnClosed(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnClosed(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffClosed returns true if the function "WEBEXT.app.currentWindowInternal.onClosed.removeListener" exists.
func HasFuncOffClosed() bool {
	return js.True == bindings.HasFuncOffClosed()
}

// FuncOffClosed returns the function "WEBEXT.app.currentWindowInternal.onClosed.removeListener".
func FuncOffClosed() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffClosed(
		js.Pointer(&fn),
	)
	return
}

// OffClosed calls the function "WEBEXT.app.currentWindowInternal.onClosed.removeListener" directly.
func OffClosed(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffClosed(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffClosed calls the function "WEBEXT.app.currentWindowInternal.onClosed.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffClosed(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffClosed(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnClosed returns true if the function "WEBEXT.app.currentWindowInternal.onClosed.hasListener" exists.
func HasFuncHasOnClosed() bool {
	return js.True == bindings.HasFuncHasOnClosed()
}

// FuncHasOnClosed returns the function "WEBEXT.app.currentWindowInternal.onClosed.hasListener".
func FuncHasOnClosed() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnClosed(
		js.Pointer(&fn),
	)
	return
}

// HasOnClosed calls the function "WEBEXT.app.currentWindowInternal.onClosed.hasListener" directly.
func HasOnClosed(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnClosed(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnClosed calls the function "WEBEXT.app.currentWindowInternal.onClosed.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnClosed(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnClosed(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnFullscreenedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnFullscreenedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnFullscreenedEventCallbackFunc) DispatchCallback(
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

type OnFullscreenedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnFullscreenedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnFullscreenedEventCallback[T]) DispatchCallback(
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

// HasFuncOnFullscreened returns true if the function "WEBEXT.app.currentWindowInternal.onFullscreened.addListener" exists.
func HasFuncOnFullscreened() bool {
	return js.True == bindings.HasFuncOnFullscreened()
}

// FuncOnFullscreened returns the function "WEBEXT.app.currentWindowInternal.onFullscreened.addListener".
func FuncOnFullscreened() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnFullscreened(
		js.Pointer(&fn),
	)
	return
}

// OnFullscreened calls the function "WEBEXT.app.currentWindowInternal.onFullscreened.addListener" directly.
func OnFullscreened(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnFullscreened(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnFullscreened calls the function "WEBEXT.app.currentWindowInternal.onFullscreened.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnFullscreened(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnFullscreened(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffFullscreened returns true if the function "WEBEXT.app.currentWindowInternal.onFullscreened.removeListener" exists.
func HasFuncOffFullscreened() bool {
	return js.True == bindings.HasFuncOffFullscreened()
}

// FuncOffFullscreened returns the function "WEBEXT.app.currentWindowInternal.onFullscreened.removeListener".
func FuncOffFullscreened() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffFullscreened(
		js.Pointer(&fn),
	)
	return
}

// OffFullscreened calls the function "WEBEXT.app.currentWindowInternal.onFullscreened.removeListener" directly.
func OffFullscreened(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffFullscreened(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffFullscreened calls the function "WEBEXT.app.currentWindowInternal.onFullscreened.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffFullscreened(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffFullscreened(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnFullscreened returns true if the function "WEBEXT.app.currentWindowInternal.onFullscreened.hasListener" exists.
func HasFuncHasOnFullscreened() bool {
	return js.True == bindings.HasFuncHasOnFullscreened()
}

// FuncHasOnFullscreened returns the function "WEBEXT.app.currentWindowInternal.onFullscreened.hasListener".
func FuncHasOnFullscreened() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnFullscreened(
		js.Pointer(&fn),
	)
	return
}

// HasOnFullscreened calls the function "WEBEXT.app.currentWindowInternal.onFullscreened.hasListener" directly.
func HasOnFullscreened(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnFullscreened(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnFullscreened calls the function "WEBEXT.app.currentWindowInternal.onFullscreened.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnFullscreened(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnFullscreened(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMaximizedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnMaximizedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMaximizedEventCallbackFunc) DispatchCallback(
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

type OnMaximizedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnMaximizedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMaximizedEventCallback[T]) DispatchCallback(
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

// HasFuncOnMaximized returns true if the function "WEBEXT.app.currentWindowInternal.onMaximized.addListener" exists.
func HasFuncOnMaximized() bool {
	return js.True == bindings.HasFuncOnMaximized()
}

// FuncOnMaximized returns the function "WEBEXT.app.currentWindowInternal.onMaximized.addListener".
func FuncOnMaximized() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnMaximized(
		js.Pointer(&fn),
	)
	return
}

// OnMaximized calls the function "WEBEXT.app.currentWindowInternal.onMaximized.addListener" directly.
func OnMaximized(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnMaximized(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMaximized calls the function "WEBEXT.app.currentWindowInternal.onMaximized.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMaximized(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMaximized(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMaximized returns true if the function "WEBEXT.app.currentWindowInternal.onMaximized.removeListener" exists.
func HasFuncOffMaximized() bool {
	return js.True == bindings.HasFuncOffMaximized()
}

// FuncOffMaximized returns the function "WEBEXT.app.currentWindowInternal.onMaximized.removeListener".
func FuncOffMaximized() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffMaximized(
		js.Pointer(&fn),
	)
	return
}

// OffMaximized calls the function "WEBEXT.app.currentWindowInternal.onMaximized.removeListener" directly.
func OffMaximized(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffMaximized(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMaximized calls the function "WEBEXT.app.currentWindowInternal.onMaximized.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMaximized(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMaximized(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMaximized returns true if the function "WEBEXT.app.currentWindowInternal.onMaximized.hasListener" exists.
func HasFuncHasOnMaximized() bool {
	return js.True == bindings.HasFuncHasOnMaximized()
}

// FuncHasOnMaximized returns the function "WEBEXT.app.currentWindowInternal.onMaximized.hasListener".
func FuncHasOnMaximized() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnMaximized(
		js.Pointer(&fn),
	)
	return
}

// HasOnMaximized calls the function "WEBEXT.app.currentWindowInternal.onMaximized.hasListener" directly.
func HasOnMaximized(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnMaximized(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMaximized calls the function "WEBEXT.app.currentWindowInternal.onMaximized.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMaximized(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMaximized(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMinimizedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnMinimizedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMinimizedEventCallbackFunc) DispatchCallback(
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

type OnMinimizedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnMinimizedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMinimizedEventCallback[T]) DispatchCallback(
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

// HasFuncOnMinimized returns true if the function "WEBEXT.app.currentWindowInternal.onMinimized.addListener" exists.
func HasFuncOnMinimized() bool {
	return js.True == bindings.HasFuncOnMinimized()
}

// FuncOnMinimized returns the function "WEBEXT.app.currentWindowInternal.onMinimized.addListener".
func FuncOnMinimized() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnMinimized(
		js.Pointer(&fn),
	)
	return
}

// OnMinimized calls the function "WEBEXT.app.currentWindowInternal.onMinimized.addListener" directly.
func OnMinimized(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnMinimized(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMinimized calls the function "WEBEXT.app.currentWindowInternal.onMinimized.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMinimized(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMinimized(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMinimized returns true if the function "WEBEXT.app.currentWindowInternal.onMinimized.removeListener" exists.
func HasFuncOffMinimized() bool {
	return js.True == bindings.HasFuncOffMinimized()
}

// FuncOffMinimized returns the function "WEBEXT.app.currentWindowInternal.onMinimized.removeListener".
func FuncOffMinimized() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffMinimized(
		js.Pointer(&fn),
	)
	return
}

// OffMinimized calls the function "WEBEXT.app.currentWindowInternal.onMinimized.removeListener" directly.
func OffMinimized(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffMinimized(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMinimized calls the function "WEBEXT.app.currentWindowInternal.onMinimized.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMinimized(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMinimized(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMinimized returns true if the function "WEBEXT.app.currentWindowInternal.onMinimized.hasListener" exists.
func HasFuncHasOnMinimized() bool {
	return js.True == bindings.HasFuncHasOnMinimized()
}

// FuncHasOnMinimized returns the function "WEBEXT.app.currentWindowInternal.onMinimized.hasListener".
func FuncHasOnMinimized() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnMinimized(
		js.Pointer(&fn),
	)
	return
}

// HasOnMinimized calls the function "WEBEXT.app.currentWindowInternal.onMinimized.hasListener" directly.
func HasOnMinimized(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnMinimized(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMinimized calls the function "WEBEXT.app.currentWindowInternal.onMinimized.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMinimized(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMinimized(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRestoredEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnRestoredEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRestoredEventCallbackFunc) DispatchCallback(
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

type OnRestoredEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnRestoredEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRestoredEventCallback[T]) DispatchCallback(
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

// HasFuncOnRestored returns true if the function "WEBEXT.app.currentWindowInternal.onRestored.addListener" exists.
func HasFuncOnRestored() bool {
	return js.True == bindings.HasFuncOnRestored()
}

// FuncOnRestored returns the function "WEBEXT.app.currentWindowInternal.onRestored.addListener".
func FuncOnRestored() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnRestored(
		js.Pointer(&fn),
	)
	return
}

// OnRestored calls the function "WEBEXT.app.currentWindowInternal.onRestored.addListener" directly.
func OnRestored(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnRestored(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRestored calls the function "WEBEXT.app.currentWindowInternal.onRestored.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRestored(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRestored(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRestored returns true if the function "WEBEXT.app.currentWindowInternal.onRestored.removeListener" exists.
func HasFuncOffRestored() bool {
	return js.True == bindings.HasFuncOffRestored()
}

// FuncOffRestored returns the function "WEBEXT.app.currentWindowInternal.onRestored.removeListener".
func FuncOffRestored() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffRestored(
		js.Pointer(&fn),
	)
	return
}

// OffRestored calls the function "WEBEXT.app.currentWindowInternal.onRestored.removeListener" directly.
func OffRestored(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffRestored(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRestored calls the function "WEBEXT.app.currentWindowInternal.onRestored.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRestored(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRestored(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRestored returns true if the function "WEBEXT.app.currentWindowInternal.onRestored.hasListener" exists.
func HasFuncHasOnRestored() bool {
	return js.True == bindings.HasFuncHasOnRestored()
}

// FuncHasOnRestored returns the function "WEBEXT.app.currentWindowInternal.onRestored.hasListener".
func FuncHasOnRestored() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnRestored(
		js.Pointer(&fn),
	)
	return
}

// HasOnRestored calls the function "WEBEXT.app.currentWindowInternal.onRestored.hasListener" directly.
func HasOnRestored(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnRestored(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRestored calls the function "WEBEXT.app.currentWindowInternal.onRestored.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRestored(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRestored(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRestore returns true if the function "WEBEXT.app.currentWindowInternal.restore" exists.
func HasFuncRestore() bool {
	return js.True == bindings.HasFuncRestore()
}

// FuncRestore returns the function "WEBEXT.app.currentWindowInternal.restore".
func FuncRestore() (fn js.Func[func()]) {
	bindings.FuncRestore(
		js.Pointer(&fn),
	)
	return
}

// Restore calls the function "WEBEXT.app.currentWindowInternal.restore" directly.
func Restore() (ret js.Void) {
	bindings.CallRestore(
		js.Pointer(&ret),
	)

	return
}

// TryRestore calls the function "WEBEXT.app.currentWindowInternal.restore"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRestore() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRestore(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetActivateOnPointer returns true if the function "WEBEXT.app.currentWindowInternal.setActivateOnPointer" exists.
func HasFuncSetActivateOnPointer() bool {
	return js.True == bindings.HasFuncSetActivateOnPointer()
}

// FuncSetActivateOnPointer returns the function "WEBEXT.app.currentWindowInternal.setActivateOnPointer".
func FuncSetActivateOnPointer() (fn js.Func[func(activate_on_pointer bool)]) {
	bindings.FuncSetActivateOnPointer(
		js.Pointer(&fn),
	)
	return
}

// SetActivateOnPointer calls the function "WEBEXT.app.currentWindowInternal.setActivateOnPointer" directly.
func SetActivateOnPointer(activate_on_pointer bool) (ret js.Void) {
	bindings.CallSetActivateOnPointer(
		js.Pointer(&ret),
		js.Bool(bool(activate_on_pointer)),
	)

	return
}

// TrySetActivateOnPointer calls the function "WEBEXT.app.currentWindowInternal.setActivateOnPointer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetActivateOnPointer(activate_on_pointer bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetActivateOnPointer(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(activate_on_pointer)),
	)

	return
}

// HasFuncSetAlwaysOnTop returns true if the function "WEBEXT.app.currentWindowInternal.setAlwaysOnTop" exists.
func HasFuncSetAlwaysOnTop() bool {
	return js.True == bindings.HasFuncSetAlwaysOnTop()
}

// FuncSetAlwaysOnTop returns the function "WEBEXT.app.currentWindowInternal.setAlwaysOnTop".
func FuncSetAlwaysOnTop() (fn js.Func[func(always_on_top bool)]) {
	bindings.FuncSetAlwaysOnTop(
		js.Pointer(&fn),
	)
	return
}

// SetAlwaysOnTop calls the function "WEBEXT.app.currentWindowInternal.setAlwaysOnTop" directly.
func SetAlwaysOnTop(always_on_top bool) (ret js.Void) {
	bindings.CallSetAlwaysOnTop(
		js.Pointer(&ret),
		js.Bool(bool(always_on_top)),
	)

	return
}

// TrySetAlwaysOnTop calls the function "WEBEXT.app.currentWindowInternal.setAlwaysOnTop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetAlwaysOnTop(always_on_top bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetAlwaysOnTop(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(always_on_top)),
	)

	return
}

// HasFuncSetBounds returns true if the function "WEBEXT.app.currentWindowInternal.setBounds" exists.
func HasFuncSetBounds() bool {
	return js.True == bindings.HasFuncSetBounds()
}

// FuncSetBounds returns the function "WEBEXT.app.currentWindowInternal.setBounds".
func FuncSetBounds() (fn js.Func[func(boundsType js.String, bounds Bounds)]) {
	bindings.FuncSetBounds(
		js.Pointer(&fn),
	)
	return
}

// SetBounds calls the function "WEBEXT.app.currentWindowInternal.setBounds" directly.
func SetBounds(boundsType js.String, bounds Bounds) (ret js.Void) {
	bindings.CallSetBounds(
		js.Pointer(&ret),
		boundsType.Ref(),
		js.Pointer(&bounds),
	)

	return
}

// TrySetBounds calls the function "WEBEXT.app.currentWindowInternal.setBounds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetBounds(boundsType js.String, bounds Bounds) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetBounds(
		js.Pointer(&ret), js.Pointer(&exception),
		boundsType.Ref(),
		js.Pointer(&bounds),
	)

	return
}

// HasFuncSetIcon returns true if the function "WEBEXT.app.currentWindowInternal.setIcon" exists.
func HasFuncSetIcon() bool {
	return js.True == bindings.HasFuncSetIcon()
}

// FuncSetIcon returns the function "WEBEXT.app.currentWindowInternal.setIcon".
func FuncSetIcon() (fn js.Func[func(icon_url js.String)]) {
	bindings.FuncSetIcon(
		js.Pointer(&fn),
	)
	return
}

// SetIcon calls the function "WEBEXT.app.currentWindowInternal.setIcon" directly.
func SetIcon(icon_url js.String) (ret js.Void) {
	bindings.CallSetIcon(
		js.Pointer(&ret),
		icon_url.Ref(),
	)

	return
}

// TrySetIcon calls the function "WEBEXT.app.currentWindowInternal.setIcon"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetIcon(icon_url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetIcon(
		js.Pointer(&ret), js.Pointer(&exception),
		icon_url.Ref(),
	)

	return
}

// HasFuncSetShape returns true if the function "WEBEXT.app.currentWindowInternal.setShape" exists.
func HasFuncSetShape() bool {
	return js.True == bindings.HasFuncSetShape()
}

// FuncSetShape returns the function "WEBEXT.app.currentWindowInternal.setShape".
func FuncSetShape() (fn js.Func[func(region Region)]) {
	bindings.FuncSetShape(
		js.Pointer(&fn),
	)
	return
}

// SetShape calls the function "WEBEXT.app.currentWindowInternal.setShape" directly.
func SetShape(region Region) (ret js.Void) {
	bindings.CallSetShape(
		js.Pointer(&ret),
		js.Pointer(&region),
	)

	return
}

// TrySetShape calls the function "WEBEXT.app.currentWindowInternal.setShape"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetShape(region Region) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetShape(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&region),
	)

	return
}

// HasFuncSetSizeConstraints returns true if the function "WEBEXT.app.currentWindowInternal.setSizeConstraints" exists.
func HasFuncSetSizeConstraints() bool {
	return js.True == bindings.HasFuncSetSizeConstraints()
}

// FuncSetSizeConstraints returns the function "WEBEXT.app.currentWindowInternal.setSizeConstraints".
func FuncSetSizeConstraints() (fn js.Func[func(boundsType js.String, constraints SizeConstraints)]) {
	bindings.FuncSetSizeConstraints(
		js.Pointer(&fn),
	)
	return
}

// SetSizeConstraints calls the function "WEBEXT.app.currentWindowInternal.setSizeConstraints" directly.
func SetSizeConstraints(boundsType js.String, constraints SizeConstraints) (ret js.Void) {
	bindings.CallSetSizeConstraints(
		js.Pointer(&ret),
		boundsType.Ref(),
		js.Pointer(&constraints),
	)

	return
}

// TrySetSizeConstraints calls the function "WEBEXT.app.currentWindowInternal.setSizeConstraints"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetSizeConstraints(boundsType js.String, constraints SizeConstraints) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetSizeConstraints(
		js.Pointer(&ret), js.Pointer(&exception),
		boundsType.Ref(),
		js.Pointer(&constraints),
	)

	return
}

// HasFuncSetVisibleOnAllWorkspaces returns true if the function "WEBEXT.app.currentWindowInternal.setVisibleOnAllWorkspaces" exists.
func HasFuncSetVisibleOnAllWorkspaces() bool {
	return js.True == bindings.HasFuncSetVisibleOnAllWorkspaces()
}

// FuncSetVisibleOnAllWorkspaces returns the function "WEBEXT.app.currentWindowInternal.setVisibleOnAllWorkspaces".
func FuncSetVisibleOnAllWorkspaces() (fn js.Func[func(always_visible bool)]) {
	bindings.FuncSetVisibleOnAllWorkspaces(
		js.Pointer(&fn),
	)
	return
}

// SetVisibleOnAllWorkspaces calls the function "WEBEXT.app.currentWindowInternal.setVisibleOnAllWorkspaces" directly.
func SetVisibleOnAllWorkspaces(always_visible bool) (ret js.Void) {
	bindings.CallSetVisibleOnAllWorkspaces(
		js.Pointer(&ret),
		js.Bool(bool(always_visible)),
	)

	return
}

// TrySetVisibleOnAllWorkspaces calls the function "WEBEXT.app.currentWindowInternal.setVisibleOnAllWorkspaces"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetVisibleOnAllWorkspaces(always_visible bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetVisibleOnAllWorkspaces(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(always_visible)),
	)

	return
}

// HasFuncShow returns true if the function "WEBEXT.app.currentWindowInternal.show" exists.
func HasFuncShow() bool {
	return js.True == bindings.HasFuncShow()
}

// FuncShow returns the function "WEBEXT.app.currentWindowInternal.show".
func FuncShow() (fn js.Func[func(focused bool)]) {
	bindings.FuncShow(
		js.Pointer(&fn),
	)
	return
}

// Show calls the function "WEBEXT.app.currentWindowInternal.show" directly.
func Show(focused bool) (ret js.Void) {
	bindings.CallShow(
		js.Pointer(&ret),
		js.Bool(bool(focused)),
	)

	return
}

// TryShow calls the function "WEBEXT.app.currentWindowInternal.show"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShow(focused bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryShow(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(focused)),
	)

	return
}
