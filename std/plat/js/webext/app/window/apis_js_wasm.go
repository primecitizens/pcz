// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package window

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/app/runtime"
	"github.com/primecitizens/pcz/std/plat/js/webext/app/window/bindings"
)

type ContentBounds struct {
	// Left is "ContentBounds.left"
	//
	// Optional
	//
	// NOTE: FFI_USE_Left MUST be set to true to make this field effective.
	Left int32
	// Top is "ContentBounds.top"
	//
	// Optional
	//
	// NOTE: FFI_USE_Top MUST be set to true to make this field effective.
	Top int32
	// Width is "ContentBounds.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int32
	// Height is "ContentBounds.height"
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

// FromRef calls UpdateFrom and returns a ContentBounds with all fields set.
func (p ContentBounds) FromRef(ref js.Ref) ContentBounds {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ContentBounds in the application heap.
func (p ContentBounds) New() js.Ref {
	return bindings.ContentBoundsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ContentBounds) UpdateFrom(ref js.Ref) {
	bindings.ContentBoundsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ContentBounds) Update(ref js.Ref) {
	bindings.ContentBoundsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ContentBounds) FreeMembers(recursive bool) {
}

type Bounds struct {
	ref js.Ref
}

func (this Bounds) Once() Bounds {
	this.ref.Once()
	return this
}

func (this Bounds) Ref() js.Ref {
	return this.ref
}

func (this Bounds) FromRef(ref js.Ref) Bounds {
	this.ref = ref
	return this
}

func (this Bounds) Free() {
	this.ref.Free()
}

// Left returns the value of property "Bounds.left".
//
// It returns ok=false if there is no such property.
func (this Bounds) Left() (ret int32, ok bool) {
	ok = js.True == bindings.GetBoundsLeft(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLeft sets the value of property "Bounds.left" to val.
//
// It returns false if the property cannot be set.
func (this Bounds) SetLeft(val int32) bool {
	return js.True == bindings.SetBoundsLeft(
		this.ref,
		int32(val),
	)
}

// Top returns the value of property "Bounds.top".
//
// It returns ok=false if there is no such property.
func (this Bounds) Top() (ret int32, ok bool) {
	ok = js.True == bindings.GetBoundsTop(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTop sets the value of property "Bounds.top" to val.
//
// It returns false if the property cannot be set.
func (this Bounds) SetTop(val int32) bool {
	return js.True == bindings.SetBoundsTop(
		this.ref,
		int32(val),
	)
}

// Width returns the value of property "Bounds.width".
//
// It returns ok=false if there is no such property.
func (this Bounds) Width() (ret int32, ok bool) {
	ok = js.True == bindings.GetBoundsWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "Bounds.width" to val.
//
// It returns false if the property cannot be set.
func (this Bounds) SetWidth(val int32) bool {
	return js.True == bindings.SetBoundsWidth(
		this.ref,
		int32(val),
	)
}

// Height returns the value of property "Bounds.height".
//
// It returns ok=false if there is no such property.
func (this Bounds) Height() (ret int32, ok bool) {
	ok = js.True == bindings.GetBoundsHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "Bounds.height" to val.
//
// It returns false if the property cannot be set.
func (this Bounds) SetHeight(val int32) bool {
	return js.True == bindings.SetBoundsHeight(
		this.ref,
		int32(val),
	)
}

// MinWidth returns the value of property "Bounds.minWidth".
//
// It returns ok=false if there is no such property.
func (this Bounds) MinWidth() (ret int32, ok bool) {
	ok = js.True == bindings.GetBoundsMinWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMinWidth sets the value of property "Bounds.minWidth" to val.
//
// It returns false if the property cannot be set.
func (this Bounds) SetMinWidth(val int32) bool {
	return js.True == bindings.SetBoundsMinWidth(
		this.ref,
		int32(val),
	)
}

// MinHeight returns the value of property "Bounds.minHeight".
//
// It returns ok=false if there is no such property.
func (this Bounds) MinHeight() (ret int32, ok bool) {
	ok = js.True == bindings.GetBoundsMinHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMinHeight sets the value of property "Bounds.minHeight" to val.
//
// It returns false if the property cannot be set.
func (this Bounds) SetMinHeight(val int32) bool {
	return js.True == bindings.SetBoundsMinHeight(
		this.ref,
		int32(val),
	)
}

// MaxWidth returns the value of property "Bounds.maxWidth".
//
// It returns ok=false if there is no such property.
func (this Bounds) MaxWidth() (ret int32, ok bool) {
	ok = js.True == bindings.GetBoundsMaxWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMaxWidth sets the value of property "Bounds.maxWidth" to val.
//
// It returns false if the property cannot be set.
func (this Bounds) SetMaxWidth(val int32) bool {
	return js.True == bindings.SetBoundsMaxWidth(
		this.ref,
		int32(val),
	)
}

// MaxHeight returns the value of property "Bounds.maxHeight".
//
// It returns ok=false if there is no such property.
func (this Bounds) MaxHeight() (ret int32, ok bool) {
	ok = js.True == bindings.GetBoundsMaxHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMaxHeight sets the value of property "Bounds.maxHeight" to val.
//
// It returns false if the property cannot be set.
func (this Bounds) SetMaxHeight(val int32) bool {
	return js.True == bindings.SetBoundsMaxHeight(
		this.ref,
		int32(val),
	)
}

// HasFuncSetPosition returns true if the method "Bounds.setPosition" exists.
func (this Bounds) HasFuncSetPosition() bool {
	return js.True == bindings.HasFuncBoundsSetPosition(
		this.ref,
	)
}

// FuncSetPosition returns the method "Bounds.setPosition".
func (this Bounds) FuncSetPosition() (fn js.Func[func(left int32, top int32)]) {
	bindings.FuncBoundsSetPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetPosition calls the method "Bounds.setPosition".
func (this Bounds) SetPosition(left int32, top int32) (ret js.Void) {
	bindings.CallBoundsSetPosition(
		this.ref, js.Pointer(&ret),
		int32(left),
		int32(top),
	)

	return
}

// TrySetPosition calls the method "Bounds.setPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Bounds) TrySetPosition(left int32, top int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBoundsSetPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(left),
		int32(top),
	)

	return
}

// HasFuncSetSize returns true if the method "Bounds.setSize" exists.
func (this Bounds) HasFuncSetSize() bool {
	return js.True == bindings.HasFuncBoundsSetSize(
		this.ref,
	)
}

// FuncSetSize returns the method "Bounds.setSize".
func (this Bounds) FuncSetSize() (fn js.Func[func(width int32, height int32)]) {
	bindings.FuncBoundsSetSize(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetSize calls the method "Bounds.setSize".
func (this Bounds) SetSize(width int32, height int32) (ret js.Void) {
	bindings.CallBoundsSetSize(
		this.ref, js.Pointer(&ret),
		int32(width),
		int32(height),
	)

	return
}

// TrySetSize calls the method "Bounds.setSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Bounds) TrySetSize(width int32, height int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBoundsSetSize(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(width),
		int32(height),
	)

	return
}

// HasFuncSetMinimumSize returns true if the method "Bounds.setMinimumSize" exists.
func (this Bounds) HasFuncSetMinimumSize() bool {
	return js.True == bindings.HasFuncBoundsSetMinimumSize(
		this.ref,
	)
}

// FuncSetMinimumSize returns the method "Bounds.setMinimumSize".
func (this Bounds) FuncSetMinimumSize() (fn js.Func[func(minWidth int32, minHeight int32)]) {
	bindings.FuncBoundsSetMinimumSize(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetMinimumSize calls the method "Bounds.setMinimumSize".
func (this Bounds) SetMinimumSize(minWidth int32, minHeight int32) (ret js.Void) {
	bindings.CallBoundsSetMinimumSize(
		this.ref, js.Pointer(&ret),
		int32(minWidth),
		int32(minHeight),
	)

	return
}

// TrySetMinimumSize calls the method "Bounds.setMinimumSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Bounds) TrySetMinimumSize(minWidth int32, minHeight int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBoundsSetMinimumSize(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(minWidth),
		int32(minHeight),
	)

	return
}

// HasFuncSetMaximumSize returns true if the method "Bounds.setMaximumSize" exists.
func (this Bounds) HasFuncSetMaximumSize() bool {
	return js.True == bindings.HasFuncBoundsSetMaximumSize(
		this.ref,
	)
}

// FuncSetMaximumSize returns the method "Bounds.setMaximumSize".
func (this Bounds) FuncSetMaximumSize() (fn js.Func[func(maxWidth int32, maxHeight int32)]) {
	bindings.FuncBoundsSetMaximumSize(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetMaximumSize calls the method "Bounds.setMaximumSize".
func (this Bounds) SetMaximumSize(maxWidth int32, maxHeight int32) (ret js.Void) {
	bindings.CallBoundsSetMaximumSize(
		this.ref, js.Pointer(&ret),
		int32(maxWidth),
		int32(maxHeight),
	)

	return
}

// TrySetMaximumSize calls the method "Bounds.setMaximumSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Bounds) TrySetMaximumSize(maxWidth int32, maxHeight int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBoundsSetMaximumSize(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(maxWidth),
		int32(maxHeight),
	)

	return
}

type AppWindow struct {
	ref js.Ref
}

func (this AppWindow) Once() AppWindow {
	this.ref.Once()
	return this
}

func (this AppWindow) Ref() js.Ref {
	return this.ref
}

func (this AppWindow) FromRef(ref js.Ref) AppWindow {
	this.ref = ref
	return this
}

func (this AppWindow) Free() {
	this.ref.Free()
}

// HasFrameColor returns the value of property "AppWindow.hasFrameColor".
//
// It returns ok=false if there is no such property.
func (this AppWindow) HasFrameColor() (ret bool, ok bool) {
	ok = js.True == bindings.GetAppWindowHasFrameColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHasFrameColor sets the value of property "AppWindow.hasFrameColor" to val.
//
// It returns false if the property cannot be set.
func (this AppWindow) SetHasFrameColor(val bool) bool {
	return js.True == bindings.SetAppWindowHasFrameColor(
		this.ref,
		js.Bool(bool(val)),
	)
}

// ActiveFrameColor returns the value of property "AppWindow.activeFrameColor".
//
// It returns ok=false if there is no such property.
func (this AppWindow) ActiveFrameColor() (ret int32, ok bool) {
	ok = js.True == bindings.GetAppWindowActiveFrameColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetActiveFrameColor sets the value of property "AppWindow.activeFrameColor" to val.
//
// It returns false if the property cannot be set.
func (this AppWindow) SetActiveFrameColor(val int32) bool {
	return js.True == bindings.SetAppWindowActiveFrameColor(
		this.ref,
		int32(val),
	)
}

// InactiveFrameColor returns the value of property "AppWindow.inactiveFrameColor".
//
// It returns ok=false if there is no such property.
func (this AppWindow) InactiveFrameColor() (ret int32, ok bool) {
	ok = js.True == bindings.GetAppWindowInactiveFrameColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetInactiveFrameColor sets the value of property "AppWindow.inactiveFrameColor" to val.
//
// It returns false if the property cannot be set.
func (this AppWindow) SetInactiveFrameColor(val int32) bool {
	return js.True == bindings.SetAppWindowInactiveFrameColor(
		this.ref,
		int32(val),
	)
}

// ContentWindow returns the value of property "AppWindow.contentWindow".
//
// It returns ok=false if there is no such property.
func (this AppWindow) ContentWindow() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetAppWindowContentWindow(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetContentWindow sets the value of property "AppWindow.contentWindow" to val.
//
// It returns false if the property cannot be set.
func (this AppWindow) SetContentWindow(val js.Object) bool {
	return js.True == bindings.SetAppWindowContentWindow(
		this.ref,
		val.Ref(),
	)
}

// Id returns the value of property "AppWindow.id".
//
// It returns ok=false if there is no such property.
func (this AppWindow) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAppWindowId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetId sets the value of property "AppWindow.id" to val.
//
// It returns false if the property cannot be set.
func (this AppWindow) SetId(val js.String) bool {
	return js.True == bindings.SetAppWindowId(
		this.ref,
		val.Ref(),
	)
}

// InnerBounds returns the value of property "AppWindow.innerBounds".
//
// It returns ok=false if there is no such property.
func (this AppWindow) InnerBounds() (ret Bounds, ok bool) {
	ok = js.True == bindings.GetAppWindowInnerBounds(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetInnerBounds sets the value of property "AppWindow.innerBounds" to val.
//
// It returns false if the property cannot be set.
func (this AppWindow) SetInnerBounds(val Bounds) bool {
	return js.True == bindings.SetAppWindowInnerBounds(
		this.ref,
		val.Ref(),
	)
}

// OuterBounds returns the value of property "AppWindow.outerBounds".
//
// It returns ok=false if there is no such property.
func (this AppWindow) OuterBounds() (ret Bounds, ok bool) {
	ok = js.True == bindings.GetAppWindowOuterBounds(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetOuterBounds sets the value of property "AppWindow.outerBounds" to val.
//
// It returns false if the property cannot be set.
func (this AppWindow) SetOuterBounds(val Bounds) bool {
	return js.True == bindings.SetAppWindowOuterBounds(
		this.ref,
		val.Ref(),
	)
}

// HasFuncFocus returns true if the method "AppWindow.focus" exists.
func (this AppWindow) HasFuncFocus() bool {
	return js.True == bindings.HasFuncAppWindowFocus(
		this.ref,
	)
}

// FuncFocus returns the method "AppWindow.focus".
func (this AppWindow) FuncFocus() (fn js.Func[func()]) {
	bindings.FuncAppWindowFocus(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Focus calls the method "AppWindow.focus".
func (this AppWindow) Focus() (ret js.Void) {
	bindings.CallAppWindowFocus(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFocus calls the method "AppWindow.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryFocus() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowFocus(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFullscreen returns true if the method "AppWindow.fullscreen" exists.
func (this AppWindow) HasFuncFullscreen() bool {
	return js.True == bindings.HasFuncAppWindowFullscreen(
		this.ref,
	)
}

// FuncFullscreen returns the method "AppWindow.fullscreen".
func (this AppWindow) FuncFullscreen() (fn js.Func[func()]) {
	bindings.FuncAppWindowFullscreen(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Fullscreen calls the method "AppWindow.fullscreen".
func (this AppWindow) Fullscreen() (ret js.Void) {
	bindings.CallAppWindowFullscreen(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFullscreen calls the method "AppWindow.fullscreen"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryFullscreen() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowFullscreen(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsFullscreen returns true if the method "AppWindow.isFullscreen" exists.
func (this AppWindow) HasFuncIsFullscreen() bool {
	return js.True == bindings.HasFuncAppWindowIsFullscreen(
		this.ref,
	)
}

// FuncIsFullscreen returns the method "AppWindow.isFullscreen".
func (this AppWindow) FuncIsFullscreen() (fn js.Func[func() bool]) {
	bindings.FuncAppWindowIsFullscreen(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsFullscreen calls the method "AppWindow.isFullscreen".
func (this AppWindow) IsFullscreen() (ret bool) {
	bindings.CallAppWindowIsFullscreen(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsFullscreen calls the method "AppWindow.isFullscreen"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryIsFullscreen() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowIsFullscreen(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMinimize returns true if the method "AppWindow.minimize" exists.
func (this AppWindow) HasFuncMinimize() bool {
	return js.True == bindings.HasFuncAppWindowMinimize(
		this.ref,
	)
}

// FuncMinimize returns the method "AppWindow.minimize".
func (this AppWindow) FuncMinimize() (fn js.Func[func()]) {
	bindings.FuncAppWindowMinimize(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Minimize calls the method "AppWindow.minimize".
func (this AppWindow) Minimize() (ret js.Void) {
	bindings.CallAppWindowMinimize(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMinimize calls the method "AppWindow.minimize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryMinimize() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowMinimize(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsMinimized returns true if the method "AppWindow.isMinimized" exists.
func (this AppWindow) HasFuncIsMinimized() bool {
	return js.True == bindings.HasFuncAppWindowIsMinimized(
		this.ref,
	)
}

// FuncIsMinimized returns the method "AppWindow.isMinimized".
func (this AppWindow) FuncIsMinimized() (fn js.Func[func() bool]) {
	bindings.FuncAppWindowIsMinimized(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsMinimized calls the method "AppWindow.isMinimized".
func (this AppWindow) IsMinimized() (ret bool) {
	bindings.CallAppWindowIsMinimized(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsMinimized calls the method "AppWindow.isMinimized"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryIsMinimized() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowIsMinimized(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMaximize returns true if the method "AppWindow.maximize" exists.
func (this AppWindow) HasFuncMaximize() bool {
	return js.True == bindings.HasFuncAppWindowMaximize(
		this.ref,
	)
}

// FuncMaximize returns the method "AppWindow.maximize".
func (this AppWindow) FuncMaximize() (fn js.Func[func()]) {
	bindings.FuncAppWindowMaximize(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Maximize calls the method "AppWindow.maximize".
func (this AppWindow) Maximize() (ret js.Void) {
	bindings.CallAppWindowMaximize(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMaximize calls the method "AppWindow.maximize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryMaximize() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowMaximize(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsMaximized returns true if the method "AppWindow.isMaximized" exists.
func (this AppWindow) HasFuncIsMaximized() bool {
	return js.True == bindings.HasFuncAppWindowIsMaximized(
		this.ref,
	)
}

// FuncIsMaximized returns the method "AppWindow.isMaximized".
func (this AppWindow) FuncIsMaximized() (fn js.Func[func() bool]) {
	bindings.FuncAppWindowIsMaximized(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsMaximized calls the method "AppWindow.isMaximized".
func (this AppWindow) IsMaximized() (ret bool) {
	bindings.CallAppWindowIsMaximized(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsMaximized calls the method "AppWindow.isMaximized"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryIsMaximized() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowIsMaximized(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRestore returns true if the method "AppWindow.restore" exists.
func (this AppWindow) HasFuncRestore() bool {
	return js.True == bindings.HasFuncAppWindowRestore(
		this.ref,
	)
}

// FuncRestore returns the method "AppWindow.restore".
func (this AppWindow) FuncRestore() (fn js.Func[func()]) {
	bindings.FuncAppWindowRestore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Restore calls the method "AppWindow.restore".
func (this AppWindow) Restore() (ret js.Void) {
	bindings.CallAppWindowRestore(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRestore calls the method "AppWindow.restore"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryRestore() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowRestore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveTo returns true if the method "AppWindow.moveTo" exists.
func (this AppWindow) HasFuncMoveTo() bool {
	return js.True == bindings.HasFuncAppWindowMoveTo(
		this.ref,
	)
}

// FuncMoveTo returns the method "AppWindow.moveTo".
func (this AppWindow) FuncMoveTo() (fn js.Func[func(left int32, top int32)]) {
	bindings.FuncAppWindowMoveTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveTo calls the method "AppWindow.moveTo".
func (this AppWindow) MoveTo(left int32, top int32) (ret js.Void) {
	bindings.CallAppWindowMoveTo(
		this.ref, js.Pointer(&ret),
		int32(left),
		int32(top),
	)

	return
}

// TryMoveTo calls the method "AppWindow.moveTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryMoveTo(left int32, top int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowMoveTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(left),
		int32(top),
	)

	return
}

// HasFuncResizeTo returns true if the method "AppWindow.resizeTo" exists.
func (this AppWindow) HasFuncResizeTo() bool {
	return js.True == bindings.HasFuncAppWindowResizeTo(
		this.ref,
	)
}

// FuncResizeTo returns the method "AppWindow.resizeTo".
func (this AppWindow) FuncResizeTo() (fn js.Func[func(width int32, height int32)]) {
	bindings.FuncAppWindowResizeTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ResizeTo calls the method "AppWindow.resizeTo".
func (this AppWindow) ResizeTo(width int32, height int32) (ret js.Void) {
	bindings.CallAppWindowResizeTo(
		this.ref, js.Pointer(&ret),
		int32(width),
		int32(height),
	)

	return
}

// TryResizeTo calls the method "AppWindow.resizeTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryResizeTo(width int32, height int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowResizeTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(width),
		int32(height),
	)

	return
}

// HasFuncDrawAttention returns true if the method "AppWindow.drawAttention" exists.
func (this AppWindow) HasFuncDrawAttention() bool {
	return js.True == bindings.HasFuncAppWindowDrawAttention(
		this.ref,
	)
}

// FuncDrawAttention returns the method "AppWindow.drawAttention".
func (this AppWindow) FuncDrawAttention() (fn js.Func[func()]) {
	bindings.FuncAppWindowDrawAttention(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawAttention calls the method "AppWindow.drawAttention".
func (this AppWindow) DrawAttention() (ret js.Void) {
	bindings.CallAppWindowDrawAttention(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDrawAttention calls the method "AppWindow.drawAttention"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryDrawAttention() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowDrawAttention(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClearAttention returns true if the method "AppWindow.clearAttention" exists.
func (this AppWindow) HasFuncClearAttention() bool {
	return js.True == bindings.HasFuncAppWindowClearAttention(
		this.ref,
	)
}

// FuncClearAttention returns the method "AppWindow.clearAttention".
func (this AppWindow) FuncClearAttention() (fn js.Func[func()]) {
	bindings.FuncAppWindowClearAttention(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearAttention calls the method "AppWindow.clearAttention".
func (this AppWindow) ClearAttention() (ret js.Void) {
	bindings.CallAppWindowClearAttention(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClearAttention calls the method "AppWindow.clearAttention"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryClearAttention() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowClearAttention(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClose returns true if the method "AppWindow.close" exists.
func (this AppWindow) HasFuncClose() bool {
	return js.True == bindings.HasFuncAppWindowClose(
		this.ref,
	)
}

// FuncClose returns the method "AppWindow.close".
func (this AppWindow) FuncClose() (fn js.Func[func()]) {
	bindings.FuncAppWindowClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "AppWindow.close".
func (this AppWindow) Close() (ret js.Void) {
	bindings.CallAppWindowClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "AppWindow.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncShow returns true if the method "AppWindow.show" exists.
func (this AppWindow) HasFuncShow() bool {
	return js.True == bindings.HasFuncAppWindowShow(
		this.ref,
	)
}

// FuncShow returns the method "AppWindow.show".
func (this AppWindow) FuncShow() (fn js.Func[func(focused bool)]) {
	bindings.FuncAppWindowShow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Show calls the method "AppWindow.show".
func (this AppWindow) Show(focused bool) (ret js.Void) {
	bindings.CallAppWindowShow(
		this.ref, js.Pointer(&ret),
		js.Bool(bool(focused)),
	)

	return
}

// TryShow calls the method "AppWindow.show"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryShow(focused bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowShow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(focused)),
	)

	return
}

// HasFuncShow1 returns true if the method "AppWindow.show" exists.
func (this AppWindow) HasFuncShow1() bool {
	return js.True == bindings.HasFuncAppWindowShow1(
		this.ref,
	)
}

// FuncShow1 returns the method "AppWindow.show".
func (this AppWindow) FuncShow1() (fn js.Func[func()]) {
	bindings.FuncAppWindowShow1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Show1 calls the method "AppWindow.show".
func (this AppWindow) Show1() (ret js.Void) {
	bindings.CallAppWindowShow1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryShow1 calls the method "AppWindow.show"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryShow1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowShow1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncHide returns true if the method "AppWindow.hide" exists.
func (this AppWindow) HasFuncHide() bool {
	return js.True == bindings.HasFuncAppWindowHide(
		this.ref,
	)
}

// FuncHide returns the method "AppWindow.hide".
func (this AppWindow) FuncHide() (fn js.Func[func()]) {
	bindings.FuncAppWindowHide(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Hide calls the method "AppWindow.hide".
func (this AppWindow) Hide() (ret js.Void) {
	bindings.CallAppWindowHide(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryHide calls the method "AppWindow.hide"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryHide() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowHide(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetBounds returns true if the method "AppWindow.getBounds" exists.
func (this AppWindow) HasFuncGetBounds() bool {
	return js.True == bindings.HasFuncAppWindowGetBounds(
		this.ref,
	)
}

// FuncGetBounds returns the method "AppWindow.getBounds".
func (this AppWindow) FuncGetBounds() (fn js.Func[func() ContentBounds]) {
	bindings.FuncAppWindowGetBounds(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBounds calls the method "AppWindow.getBounds".
func (this AppWindow) GetBounds() (ret ContentBounds) {
	bindings.CallAppWindowGetBounds(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetBounds calls the method "AppWindow.getBounds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryGetBounds() (ret ContentBounds, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowGetBounds(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetBounds returns true if the method "AppWindow.setBounds" exists.
func (this AppWindow) HasFuncSetBounds() bool {
	return js.True == bindings.HasFuncAppWindowSetBounds(
		this.ref,
	)
}

// FuncSetBounds returns the method "AppWindow.setBounds".
func (this AppWindow) FuncSetBounds() (fn js.Func[func(bounds ContentBounds)]) {
	bindings.FuncAppWindowSetBounds(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetBounds calls the method "AppWindow.setBounds".
func (this AppWindow) SetBounds(bounds ContentBounds) (ret js.Void) {
	bindings.CallAppWindowSetBounds(
		this.ref, js.Pointer(&ret),
		js.Pointer(&bounds),
	)

	return
}

// TrySetBounds calls the method "AppWindow.setBounds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TrySetBounds(bounds ContentBounds) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowSetBounds(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&bounds),
	)

	return
}

// HasFuncSetIcon returns true if the method "AppWindow.setIcon" exists.
func (this AppWindow) HasFuncSetIcon() bool {
	return js.True == bindings.HasFuncAppWindowSetIcon(
		this.ref,
	)
}

// FuncSetIcon returns the method "AppWindow.setIcon".
func (this AppWindow) FuncSetIcon() (fn js.Func[func(iconUrl js.String)]) {
	bindings.FuncAppWindowSetIcon(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetIcon calls the method "AppWindow.setIcon".
func (this AppWindow) SetIcon(iconUrl js.String) (ret js.Void) {
	bindings.CallAppWindowSetIcon(
		this.ref, js.Pointer(&ret),
		iconUrl.Ref(),
	)

	return
}

// TrySetIcon calls the method "AppWindow.setIcon"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TrySetIcon(iconUrl js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowSetIcon(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		iconUrl.Ref(),
	)

	return
}

// HasFuncIsAlwaysOnTop returns true if the method "AppWindow.isAlwaysOnTop" exists.
func (this AppWindow) HasFuncIsAlwaysOnTop() bool {
	return js.True == bindings.HasFuncAppWindowIsAlwaysOnTop(
		this.ref,
	)
}

// FuncIsAlwaysOnTop returns the method "AppWindow.isAlwaysOnTop".
func (this AppWindow) FuncIsAlwaysOnTop() (fn js.Func[func() bool]) {
	bindings.FuncAppWindowIsAlwaysOnTop(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsAlwaysOnTop calls the method "AppWindow.isAlwaysOnTop".
func (this AppWindow) IsAlwaysOnTop() (ret bool) {
	bindings.CallAppWindowIsAlwaysOnTop(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsAlwaysOnTop calls the method "AppWindow.isAlwaysOnTop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryIsAlwaysOnTop() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowIsAlwaysOnTop(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetAlwaysOnTop returns true if the method "AppWindow.setAlwaysOnTop" exists.
func (this AppWindow) HasFuncSetAlwaysOnTop() bool {
	return js.True == bindings.HasFuncAppWindowSetAlwaysOnTop(
		this.ref,
	)
}

// FuncSetAlwaysOnTop returns the method "AppWindow.setAlwaysOnTop".
func (this AppWindow) FuncSetAlwaysOnTop() (fn js.Func[func(alwaysOnTop bool)]) {
	bindings.FuncAppWindowSetAlwaysOnTop(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetAlwaysOnTop calls the method "AppWindow.setAlwaysOnTop".
func (this AppWindow) SetAlwaysOnTop(alwaysOnTop bool) (ret js.Void) {
	bindings.CallAppWindowSetAlwaysOnTop(
		this.ref, js.Pointer(&ret),
		js.Bool(bool(alwaysOnTop)),
	)

	return
}

// TrySetAlwaysOnTop calls the method "AppWindow.setAlwaysOnTop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TrySetAlwaysOnTop(alwaysOnTop bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowSetAlwaysOnTop(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(alwaysOnTop)),
	)

	return
}

// HasFuncAlphaEnabled returns true if the method "AppWindow.alphaEnabled" exists.
func (this AppWindow) HasFuncAlphaEnabled() bool {
	return js.True == bindings.HasFuncAppWindowAlphaEnabled(
		this.ref,
	)
}

// FuncAlphaEnabled returns the method "AppWindow.alphaEnabled".
func (this AppWindow) FuncAlphaEnabled() (fn js.Func[func() bool]) {
	bindings.FuncAppWindowAlphaEnabled(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AlphaEnabled calls the method "AppWindow.alphaEnabled".
func (this AppWindow) AlphaEnabled() (ret bool) {
	bindings.CallAppWindowAlphaEnabled(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAlphaEnabled calls the method "AppWindow.alphaEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TryAlphaEnabled() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowAlphaEnabled(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetVisibleOnAllWorkspaces returns true if the method "AppWindow.setVisibleOnAllWorkspaces" exists.
func (this AppWindow) HasFuncSetVisibleOnAllWorkspaces() bool {
	return js.True == bindings.HasFuncAppWindowSetVisibleOnAllWorkspaces(
		this.ref,
	)
}

// FuncSetVisibleOnAllWorkspaces returns the method "AppWindow.setVisibleOnAllWorkspaces".
func (this AppWindow) FuncSetVisibleOnAllWorkspaces() (fn js.Func[func(alwaysVisible bool)]) {
	bindings.FuncAppWindowSetVisibleOnAllWorkspaces(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetVisibleOnAllWorkspaces calls the method "AppWindow.setVisibleOnAllWorkspaces".
func (this AppWindow) SetVisibleOnAllWorkspaces(alwaysVisible bool) (ret js.Void) {
	bindings.CallAppWindowSetVisibleOnAllWorkspaces(
		this.ref, js.Pointer(&ret),
		js.Bool(bool(alwaysVisible)),
	)

	return
}

// TrySetVisibleOnAllWorkspaces calls the method "AppWindow.setVisibleOnAllWorkspaces"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AppWindow) TrySetVisibleOnAllWorkspaces(alwaysVisible bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAppWindowSetVisibleOnAllWorkspaces(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(alwaysVisible)),
	)

	return
}

type BoundsSpecification struct {
	// Left is "BoundsSpecification.left"
	//
	// Optional
	//
	// NOTE: FFI_USE_Left MUST be set to true to make this field effective.
	Left int32
	// Top is "BoundsSpecification.top"
	//
	// Optional
	//
	// NOTE: FFI_USE_Top MUST be set to true to make this field effective.
	Top int32
	// Width is "BoundsSpecification.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int32
	// Height is "BoundsSpecification.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int32
	// MinWidth is "BoundsSpecification.minWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_MinWidth MUST be set to true to make this field effective.
	MinWidth int32
	// MinHeight is "BoundsSpecification.minHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_MinHeight MUST be set to true to make this field effective.
	MinHeight int32
	// MaxWidth is "BoundsSpecification.maxWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxWidth MUST be set to true to make this field effective.
	MaxWidth int32
	// MaxHeight is "BoundsSpecification.maxHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxHeight MUST be set to true to make this field effective.
	MaxHeight int32

	FFI_USE_Left      bool // for Left.
	FFI_USE_Top       bool // for Top.
	FFI_USE_Width     bool // for Width.
	FFI_USE_Height    bool // for Height.
	FFI_USE_MinWidth  bool // for MinWidth.
	FFI_USE_MinHeight bool // for MinHeight.
	FFI_USE_MaxWidth  bool // for MaxWidth.
	FFI_USE_MaxHeight bool // for MaxHeight.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BoundsSpecification with all fields set.
func (p BoundsSpecification) FromRef(ref js.Ref) BoundsSpecification {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BoundsSpecification in the application heap.
func (p BoundsSpecification) New() js.Ref {
	return bindings.BoundsSpecificationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BoundsSpecification) UpdateFrom(ref js.Ref) {
	bindings.BoundsSpecificationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BoundsSpecification) Update(ref js.Ref) {
	bindings.BoundsSpecificationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BoundsSpecification) FreeMembers(recursive bool) {
}

type CreateWindowCallbackFunc func(this js.Ref, createdWindow js.Object) js.Ref

func (fn CreateWindowCallbackFunc) Register() js.Func[func(createdWindow js.Object)] {
	return js.RegisterCallback[func(createdWindow js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CreateWindowCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateWindowCallback[T any] struct {
	Fn  func(arg T, this js.Ref, createdWindow js.Object) js.Ref
	Arg T
}

func (cb *CreateWindowCallback[T]) Register() js.Func[func(createdWindow js.Object)] {
	return js.RegisterCallback[func(createdWindow js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CreateWindowCallback[T]) DispatchCallback(
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

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type WindowType uint32

const (
	_ WindowType = iota

	WindowType_SHELL
	WindowType_PANEL
)

func (WindowType) FromRef(str js.Ref) WindowType {
	return WindowType(bindings.ConstOfWindowType(str))
}

func (x WindowType) String() (string, bool) {
	switch x {
	case WindowType_SHELL:
		return "shell", true
	case WindowType_PANEL:
		return "panel", true
	default:
		return "", false
	}
}

type FrameOptions struct {
	// Type is "FrameOptions.type"
	//
	// Optional
	Type js.String
	// Color is "FrameOptions.color"
	//
	// Optional
	Color js.String
	// ActiveColor is "FrameOptions.activeColor"
	//
	// Optional
	ActiveColor js.String
	// InactiveColor is "FrameOptions.inactiveColor"
	//
	// Optional
	InactiveColor js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FrameOptions with all fields set.
func (p FrameOptions) FromRef(ref js.Ref) FrameOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FrameOptions in the application heap.
func (p FrameOptions) New() js.Ref {
	return bindings.FrameOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FrameOptions) UpdateFrom(ref js.Ref) {
	bindings.FrameOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FrameOptions) Update(ref js.Ref) {
	bindings.FrameOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FrameOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Type.Ref(),
		p.Color.Ref(),
		p.ActiveColor.Ref(),
		p.InactiveColor.Ref(),
	)
	p.Type = p.Type.FromRef(js.Undefined)
	p.Color = p.Color.FromRef(js.Undefined)
	p.ActiveColor = p.ActiveColor.FromRef(js.Undefined)
	p.InactiveColor = p.InactiveColor.FromRef(js.Undefined)
}

type OneOf_String_FrameOptions struct {
	ref js.Ref
}

func (x OneOf_String_FrameOptions) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_FrameOptions) Free() {
	x.ref.Free()
}

func (x OneOf_String_FrameOptions) FromRef(ref js.Ref) OneOf_String_FrameOptions {
	return OneOf_String_FrameOptions{
		ref: ref,
	}
}

func (x OneOf_String_FrameOptions) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_FrameOptions) FrameOptions() FrameOptions {
	var ret FrameOptions
	ret.UpdateFrom(x.ref)
	return ret
}

type State uint32

const (
	_ State = iota

	State_NORMAL
	State_FULLSCREEN
	State_MAXIMIZED
	State_MINIMIZED
)

func (State) FromRef(str js.Ref) State {
	return State(bindings.ConstOfState(str))
}

func (x State) String() (string, bool) {
	switch x {
	case State_NORMAL:
		return "normal", true
	case State_FULLSCREEN:
		return "fullscreen", true
	case State_MAXIMIZED:
		return "maximized", true
	case State_MINIMIZED:
		return "minimized", true
	default:
		return "", false
	}
}

type CreateWindowOptions struct {
	// Id is "CreateWindowOptions.id"
	//
	// Optional
	Id js.String
	// InnerBounds is "CreateWindowOptions.innerBounds"
	//
	// Optional
	//
	// NOTE: InnerBounds.FFI_USE MUST be set to true to get InnerBounds used.
	InnerBounds BoundsSpecification
	// OuterBounds is "CreateWindowOptions.outerBounds"
	//
	// Optional
	//
	// NOTE: OuterBounds.FFI_USE MUST be set to true to get OuterBounds used.
	OuterBounds BoundsSpecification
	// DefaultWidth is "CreateWindowOptions.defaultWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_DefaultWidth MUST be set to true to make this field effective.
	DefaultWidth int32
	// DefaultHeight is "CreateWindowOptions.defaultHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_DefaultHeight MUST be set to true to make this field effective.
	DefaultHeight int32
	// DefaultLeft is "CreateWindowOptions.defaultLeft"
	//
	// Optional
	//
	// NOTE: FFI_USE_DefaultLeft MUST be set to true to make this field effective.
	DefaultLeft int32
	// DefaultTop is "CreateWindowOptions.defaultTop"
	//
	// Optional
	//
	// NOTE: FFI_USE_DefaultTop MUST be set to true to make this field effective.
	DefaultTop int32
	// Width is "CreateWindowOptions.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int32
	// Height is "CreateWindowOptions.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int32
	// Left is "CreateWindowOptions.left"
	//
	// Optional
	//
	// NOTE: FFI_USE_Left MUST be set to true to make this field effective.
	Left int32
	// Top is "CreateWindowOptions.top"
	//
	// Optional
	//
	// NOTE: FFI_USE_Top MUST be set to true to make this field effective.
	Top int32
	// MinWidth is "CreateWindowOptions.minWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_MinWidth MUST be set to true to make this field effective.
	MinWidth int32
	// MinHeight is "CreateWindowOptions.minHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_MinHeight MUST be set to true to make this field effective.
	MinHeight int32
	// MaxWidth is "CreateWindowOptions.maxWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxWidth MUST be set to true to make this field effective.
	MaxWidth int32
	// MaxHeight is "CreateWindowOptions.maxHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxHeight MUST be set to true to make this field effective.
	MaxHeight int32
	// Type is "CreateWindowOptions.type"
	//
	// Optional
	Type WindowType
	// Ime is "CreateWindowOptions.ime"
	//
	// Optional
	//
	// NOTE: FFI_USE_Ime MUST be set to true to make this field effective.
	Ime bool
	// ShowInShelf is "CreateWindowOptions.showInShelf"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShowInShelf MUST be set to true to make this field effective.
	ShowInShelf bool
	// Icon is "CreateWindowOptions.icon"
	//
	// Optional
	Icon js.String
	// Frame is "CreateWindowOptions.frame"
	//
	// Optional
	Frame OneOf_String_FrameOptions
	// Bounds is "CreateWindowOptions.bounds"
	//
	// Optional
	//
	// NOTE: Bounds.FFI_USE MUST be set to true to get Bounds used.
	Bounds ContentBounds
	// AlphaEnabled is "CreateWindowOptions.alphaEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_AlphaEnabled MUST be set to true to make this field effective.
	AlphaEnabled bool
	// State is "CreateWindowOptions.state"
	//
	// Optional
	State State
	// Hidden is "CreateWindowOptions.hidden"
	//
	// Optional
	//
	// NOTE: FFI_USE_Hidden MUST be set to true to make this field effective.
	Hidden bool
	// Resizable is "CreateWindowOptions.resizable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Resizable MUST be set to true to make this field effective.
	Resizable bool
	// Singleton is "CreateWindowOptions.singleton"
	//
	// Optional
	//
	// NOTE: FFI_USE_Singleton MUST be set to true to make this field effective.
	Singleton bool
	// AlwaysOnTop is "CreateWindowOptions.alwaysOnTop"
	//
	// Optional
	//
	// NOTE: FFI_USE_AlwaysOnTop MUST be set to true to make this field effective.
	AlwaysOnTop bool
	// Focused is "CreateWindowOptions.focused"
	//
	// Optional
	//
	// NOTE: FFI_USE_Focused MUST be set to true to make this field effective.
	Focused bool
	// VisibleOnAllWorkspaces is "CreateWindowOptions.visibleOnAllWorkspaces"
	//
	// Optional
	//
	// NOTE: FFI_USE_VisibleOnAllWorkspaces MUST be set to true to make this field effective.
	VisibleOnAllWorkspaces bool
	// LockScreenAction is "CreateWindowOptions.lockScreenAction"
	//
	// Optional
	LockScreenAction runtime.ActionType

	FFI_USE_DefaultWidth           bool // for DefaultWidth.
	FFI_USE_DefaultHeight          bool // for DefaultHeight.
	FFI_USE_DefaultLeft            bool // for DefaultLeft.
	FFI_USE_DefaultTop             bool // for DefaultTop.
	FFI_USE_Width                  bool // for Width.
	FFI_USE_Height                 bool // for Height.
	FFI_USE_Left                   bool // for Left.
	FFI_USE_Top                    bool // for Top.
	FFI_USE_MinWidth               bool // for MinWidth.
	FFI_USE_MinHeight              bool // for MinHeight.
	FFI_USE_MaxWidth               bool // for MaxWidth.
	FFI_USE_MaxHeight              bool // for MaxHeight.
	FFI_USE_Ime                    bool // for Ime.
	FFI_USE_ShowInShelf            bool // for ShowInShelf.
	FFI_USE_AlphaEnabled           bool // for AlphaEnabled.
	FFI_USE_Hidden                 bool // for Hidden.
	FFI_USE_Resizable              bool // for Resizable.
	FFI_USE_Singleton              bool // for Singleton.
	FFI_USE_AlwaysOnTop            bool // for AlwaysOnTop.
	FFI_USE_Focused                bool // for Focused.
	FFI_USE_VisibleOnAllWorkspaces bool // for VisibleOnAllWorkspaces.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CreateWindowOptions with all fields set.
func (p CreateWindowOptions) FromRef(ref js.Ref) CreateWindowOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CreateWindowOptions in the application heap.
func (p CreateWindowOptions) New() js.Ref {
	return bindings.CreateWindowOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CreateWindowOptions) UpdateFrom(ref js.Ref) {
	bindings.CreateWindowOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CreateWindowOptions) Update(ref js.Ref) {
	bindings.CreateWindowOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CreateWindowOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Icon.Ref(),
		p.Frame.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Icon = p.Icon.FromRef(js.Undefined)
	p.Frame = p.Frame.FromRef(js.Undefined)
	if recursive {
		p.InnerBounds.FreeMembers(true)
		p.OuterBounds.FreeMembers(true)
		p.Bounds.FreeMembers(true)
	}
}

// HasFuncCanSetVisibleOnAllWorkspaces returns true if the function "WEBEXT.app.window.canSetVisibleOnAllWorkspaces" exists.
func HasFuncCanSetVisibleOnAllWorkspaces() bool {
	return js.True == bindings.HasFuncCanSetVisibleOnAllWorkspaces()
}

// FuncCanSetVisibleOnAllWorkspaces returns the function "WEBEXT.app.window.canSetVisibleOnAllWorkspaces".
func FuncCanSetVisibleOnAllWorkspaces() (fn js.Func[func() bool]) {
	bindings.FuncCanSetVisibleOnAllWorkspaces(
		js.Pointer(&fn),
	)
	return
}

// CanSetVisibleOnAllWorkspaces calls the function "WEBEXT.app.window.canSetVisibleOnAllWorkspaces" directly.
func CanSetVisibleOnAllWorkspaces() (ret bool) {
	bindings.CallCanSetVisibleOnAllWorkspaces(
		js.Pointer(&ret),
	)

	return
}

// TryCanSetVisibleOnAllWorkspaces calls the function "WEBEXT.app.window.canSetVisibleOnAllWorkspaces"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCanSetVisibleOnAllWorkspaces() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanSetVisibleOnAllWorkspaces(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreate returns true if the function "WEBEXT.app.window.create" exists.
func HasFuncCreate() bool {
	return js.True == bindings.HasFuncCreate()
}

// FuncCreate returns the function "WEBEXT.app.window.create".
func FuncCreate() (fn js.Func[func(url js.String, options CreateWindowOptions) js.Promise[js.Object]]) {
	bindings.FuncCreate(
		js.Pointer(&fn),
	)
	return
}

// Create calls the function "WEBEXT.app.window.create" directly.
func Create(url js.String, options CreateWindowOptions) (ret js.Promise[js.Object]) {
	bindings.CallCreate(
		js.Pointer(&ret),
		url.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryCreate calls the function "WEBEXT.app.window.create"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreate(url js.String, options CreateWindowOptions) (ret js.Promise[js.Object], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreate(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncCurrent returns true if the function "WEBEXT.app.window.current" exists.
func HasFuncCurrent() bool {
	return js.True == bindings.HasFuncCurrent()
}

// FuncCurrent returns the function "WEBEXT.app.window.current".
func FuncCurrent() (fn js.Func[func() AppWindow]) {
	bindings.FuncCurrent(
		js.Pointer(&fn),
	)
	return
}

// Current calls the function "WEBEXT.app.window.current" directly.
func Current() (ret AppWindow) {
	bindings.CallCurrent(
		js.Pointer(&ret),
	)

	return
}

// TryCurrent calls the function "WEBEXT.app.window.current"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCurrent() (ret AppWindow, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCurrent(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGet returns true if the function "WEBEXT.app.window.get" exists.
func HasFuncGet() bool {
	return js.True == bindings.HasFuncGet()
}

// FuncGet returns the function "WEBEXT.app.window.get".
func FuncGet() (fn js.Func[func(id js.String) AppWindow]) {
	bindings.FuncGet(
		js.Pointer(&fn),
	)
	return
}

// Get calls the function "WEBEXT.app.window.get" directly.
func Get(id js.String) (ret AppWindow) {
	bindings.CallGet(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGet calls the function "WEBEXT.app.window.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGet(id js.String) (ret AppWindow, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGet(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncGetAll returns true if the function "WEBEXT.app.window.getAll" exists.
func HasFuncGetAll() bool {
	return js.True == bindings.HasFuncGetAll()
}

// FuncGetAll returns the function "WEBEXT.app.window.getAll".
func FuncGetAll() (fn js.Func[func() js.Array[AppWindow]]) {
	bindings.FuncGetAll(
		js.Pointer(&fn),
	)
	return
}

// GetAll calls the function "WEBEXT.app.window.getAll" directly.
func GetAll() (ret js.Array[AppWindow]) {
	bindings.CallGetAll(
		js.Pointer(&ret),
	)

	return
}

// TryGetAll calls the function "WEBEXT.app.window.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAll() (ret js.Array[AppWindow], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAll(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInitializeAppWindow returns true if the function "WEBEXT.app.window.initializeAppWindow" exists.
func HasFuncInitializeAppWindow() bool {
	return js.True == bindings.HasFuncInitializeAppWindow()
}

// FuncInitializeAppWindow returns the function "WEBEXT.app.window.initializeAppWindow".
func FuncInitializeAppWindow() (fn js.Func[func(state js.Object)]) {
	bindings.FuncInitializeAppWindow(
		js.Pointer(&fn),
	)
	return
}

// InitializeAppWindow calls the function "WEBEXT.app.window.initializeAppWindow" directly.
func InitializeAppWindow(state js.Object) (ret js.Void) {
	bindings.CallInitializeAppWindow(
		js.Pointer(&ret),
		state.Ref(),
	)

	return
}

// TryInitializeAppWindow calls the function "WEBEXT.app.window.initializeAppWindow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInitializeAppWindow(state js.Object) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInitializeAppWindow(
		js.Pointer(&ret), js.Pointer(&exception),
		state.Ref(),
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

// HasFuncOnAlphaEnabledChanged returns true if the function "WEBEXT.app.window.onAlphaEnabledChanged.addListener" exists.
func HasFuncOnAlphaEnabledChanged() bool {
	return js.True == bindings.HasFuncOnAlphaEnabledChanged()
}

// FuncOnAlphaEnabledChanged returns the function "WEBEXT.app.window.onAlphaEnabledChanged.addListener".
func FuncOnAlphaEnabledChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnAlphaEnabledChanged(
		js.Pointer(&fn),
	)
	return
}

// OnAlphaEnabledChanged calls the function "WEBEXT.app.window.onAlphaEnabledChanged.addListener" directly.
func OnAlphaEnabledChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnAlphaEnabledChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAlphaEnabledChanged calls the function "WEBEXT.app.window.onAlphaEnabledChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAlphaEnabledChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAlphaEnabledChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAlphaEnabledChanged returns true if the function "WEBEXT.app.window.onAlphaEnabledChanged.removeListener" exists.
func HasFuncOffAlphaEnabledChanged() bool {
	return js.True == bindings.HasFuncOffAlphaEnabledChanged()
}

// FuncOffAlphaEnabledChanged returns the function "WEBEXT.app.window.onAlphaEnabledChanged.removeListener".
func FuncOffAlphaEnabledChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffAlphaEnabledChanged(
		js.Pointer(&fn),
	)
	return
}

// OffAlphaEnabledChanged calls the function "WEBEXT.app.window.onAlphaEnabledChanged.removeListener" directly.
func OffAlphaEnabledChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffAlphaEnabledChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAlphaEnabledChanged calls the function "WEBEXT.app.window.onAlphaEnabledChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAlphaEnabledChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAlphaEnabledChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAlphaEnabledChanged returns true if the function "WEBEXT.app.window.onAlphaEnabledChanged.hasListener" exists.
func HasFuncHasOnAlphaEnabledChanged() bool {
	return js.True == bindings.HasFuncHasOnAlphaEnabledChanged()
}

// FuncHasOnAlphaEnabledChanged returns the function "WEBEXT.app.window.onAlphaEnabledChanged.hasListener".
func FuncHasOnAlphaEnabledChanged() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnAlphaEnabledChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnAlphaEnabledChanged calls the function "WEBEXT.app.window.onAlphaEnabledChanged.hasListener" directly.
func HasOnAlphaEnabledChanged(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnAlphaEnabledChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAlphaEnabledChanged calls the function "WEBEXT.app.window.onAlphaEnabledChanged.hasListener"
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

// HasFuncOnBoundsChanged returns true if the function "WEBEXT.app.window.onBoundsChanged.addListener" exists.
func HasFuncOnBoundsChanged() bool {
	return js.True == bindings.HasFuncOnBoundsChanged()
}

// FuncOnBoundsChanged returns the function "WEBEXT.app.window.onBoundsChanged.addListener".
func FuncOnBoundsChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// OnBoundsChanged calls the function "WEBEXT.app.window.onBoundsChanged.addListener" directly.
func OnBoundsChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnBoundsChanged calls the function "WEBEXT.app.window.onBoundsChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnBoundsChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffBoundsChanged returns true if the function "WEBEXT.app.window.onBoundsChanged.removeListener" exists.
func HasFuncOffBoundsChanged() bool {
	return js.True == bindings.HasFuncOffBoundsChanged()
}

// FuncOffBoundsChanged returns the function "WEBEXT.app.window.onBoundsChanged.removeListener".
func FuncOffBoundsChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// OffBoundsChanged calls the function "WEBEXT.app.window.onBoundsChanged.removeListener" directly.
func OffBoundsChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffBoundsChanged calls the function "WEBEXT.app.window.onBoundsChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffBoundsChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnBoundsChanged returns true if the function "WEBEXT.app.window.onBoundsChanged.hasListener" exists.
func HasFuncHasOnBoundsChanged() bool {
	return js.True == bindings.HasFuncHasOnBoundsChanged()
}

// FuncHasOnBoundsChanged returns the function "WEBEXT.app.window.onBoundsChanged.hasListener".
func FuncHasOnBoundsChanged() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnBoundsChanged calls the function "WEBEXT.app.window.onBoundsChanged.hasListener" directly.
func HasOnBoundsChanged(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnBoundsChanged calls the function "WEBEXT.app.window.onBoundsChanged.hasListener"
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

// HasFuncOnClosed returns true if the function "WEBEXT.app.window.onClosed.addListener" exists.
func HasFuncOnClosed() bool {
	return js.True == bindings.HasFuncOnClosed()
}

// FuncOnClosed returns the function "WEBEXT.app.window.onClosed.addListener".
func FuncOnClosed() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnClosed(
		js.Pointer(&fn),
	)
	return
}

// OnClosed calls the function "WEBEXT.app.window.onClosed.addListener" directly.
func OnClosed(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnClosed(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnClosed calls the function "WEBEXT.app.window.onClosed.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnClosed(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnClosed(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffClosed returns true if the function "WEBEXT.app.window.onClosed.removeListener" exists.
func HasFuncOffClosed() bool {
	return js.True == bindings.HasFuncOffClosed()
}

// FuncOffClosed returns the function "WEBEXT.app.window.onClosed.removeListener".
func FuncOffClosed() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffClosed(
		js.Pointer(&fn),
	)
	return
}

// OffClosed calls the function "WEBEXT.app.window.onClosed.removeListener" directly.
func OffClosed(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffClosed(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffClosed calls the function "WEBEXT.app.window.onClosed.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffClosed(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffClosed(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnClosed returns true if the function "WEBEXT.app.window.onClosed.hasListener" exists.
func HasFuncHasOnClosed() bool {
	return js.True == bindings.HasFuncHasOnClosed()
}

// FuncHasOnClosed returns the function "WEBEXT.app.window.onClosed.hasListener".
func FuncHasOnClosed() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnClosed(
		js.Pointer(&fn),
	)
	return
}

// HasOnClosed calls the function "WEBEXT.app.window.onClosed.hasListener" directly.
func HasOnClosed(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnClosed(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnClosed calls the function "WEBEXT.app.window.onClosed.hasListener"
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

// HasFuncOnFullscreened returns true if the function "WEBEXT.app.window.onFullscreened.addListener" exists.
func HasFuncOnFullscreened() bool {
	return js.True == bindings.HasFuncOnFullscreened()
}

// FuncOnFullscreened returns the function "WEBEXT.app.window.onFullscreened.addListener".
func FuncOnFullscreened() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnFullscreened(
		js.Pointer(&fn),
	)
	return
}

// OnFullscreened calls the function "WEBEXT.app.window.onFullscreened.addListener" directly.
func OnFullscreened(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnFullscreened(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnFullscreened calls the function "WEBEXT.app.window.onFullscreened.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnFullscreened(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnFullscreened(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffFullscreened returns true if the function "WEBEXT.app.window.onFullscreened.removeListener" exists.
func HasFuncOffFullscreened() bool {
	return js.True == bindings.HasFuncOffFullscreened()
}

// FuncOffFullscreened returns the function "WEBEXT.app.window.onFullscreened.removeListener".
func FuncOffFullscreened() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffFullscreened(
		js.Pointer(&fn),
	)
	return
}

// OffFullscreened calls the function "WEBEXT.app.window.onFullscreened.removeListener" directly.
func OffFullscreened(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffFullscreened(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffFullscreened calls the function "WEBEXT.app.window.onFullscreened.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffFullscreened(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffFullscreened(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnFullscreened returns true if the function "WEBEXT.app.window.onFullscreened.hasListener" exists.
func HasFuncHasOnFullscreened() bool {
	return js.True == bindings.HasFuncHasOnFullscreened()
}

// FuncHasOnFullscreened returns the function "WEBEXT.app.window.onFullscreened.hasListener".
func FuncHasOnFullscreened() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnFullscreened(
		js.Pointer(&fn),
	)
	return
}

// HasOnFullscreened calls the function "WEBEXT.app.window.onFullscreened.hasListener" directly.
func HasOnFullscreened(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnFullscreened(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnFullscreened calls the function "WEBEXT.app.window.onFullscreened.hasListener"
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

// HasFuncOnMaximized returns true if the function "WEBEXT.app.window.onMaximized.addListener" exists.
func HasFuncOnMaximized() bool {
	return js.True == bindings.HasFuncOnMaximized()
}

// FuncOnMaximized returns the function "WEBEXT.app.window.onMaximized.addListener".
func FuncOnMaximized() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnMaximized(
		js.Pointer(&fn),
	)
	return
}

// OnMaximized calls the function "WEBEXT.app.window.onMaximized.addListener" directly.
func OnMaximized(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnMaximized(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMaximized calls the function "WEBEXT.app.window.onMaximized.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMaximized(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMaximized(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMaximized returns true if the function "WEBEXT.app.window.onMaximized.removeListener" exists.
func HasFuncOffMaximized() bool {
	return js.True == bindings.HasFuncOffMaximized()
}

// FuncOffMaximized returns the function "WEBEXT.app.window.onMaximized.removeListener".
func FuncOffMaximized() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffMaximized(
		js.Pointer(&fn),
	)
	return
}

// OffMaximized calls the function "WEBEXT.app.window.onMaximized.removeListener" directly.
func OffMaximized(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffMaximized(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMaximized calls the function "WEBEXT.app.window.onMaximized.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMaximized(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMaximized(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMaximized returns true if the function "WEBEXT.app.window.onMaximized.hasListener" exists.
func HasFuncHasOnMaximized() bool {
	return js.True == bindings.HasFuncHasOnMaximized()
}

// FuncHasOnMaximized returns the function "WEBEXT.app.window.onMaximized.hasListener".
func FuncHasOnMaximized() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnMaximized(
		js.Pointer(&fn),
	)
	return
}

// HasOnMaximized calls the function "WEBEXT.app.window.onMaximized.hasListener" directly.
func HasOnMaximized(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnMaximized(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMaximized calls the function "WEBEXT.app.window.onMaximized.hasListener"
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

// HasFuncOnMinimized returns true if the function "WEBEXT.app.window.onMinimized.addListener" exists.
func HasFuncOnMinimized() bool {
	return js.True == bindings.HasFuncOnMinimized()
}

// FuncOnMinimized returns the function "WEBEXT.app.window.onMinimized.addListener".
func FuncOnMinimized() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnMinimized(
		js.Pointer(&fn),
	)
	return
}

// OnMinimized calls the function "WEBEXT.app.window.onMinimized.addListener" directly.
func OnMinimized(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnMinimized(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMinimized calls the function "WEBEXT.app.window.onMinimized.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMinimized(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMinimized(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMinimized returns true if the function "WEBEXT.app.window.onMinimized.removeListener" exists.
func HasFuncOffMinimized() bool {
	return js.True == bindings.HasFuncOffMinimized()
}

// FuncOffMinimized returns the function "WEBEXT.app.window.onMinimized.removeListener".
func FuncOffMinimized() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffMinimized(
		js.Pointer(&fn),
	)
	return
}

// OffMinimized calls the function "WEBEXT.app.window.onMinimized.removeListener" directly.
func OffMinimized(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffMinimized(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMinimized calls the function "WEBEXT.app.window.onMinimized.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMinimized(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMinimized(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMinimized returns true if the function "WEBEXT.app.window.onMinimized.hasListener" exists.
func HasFuncHasOnMinimized() bool {
	return js.True == bindings.HasFuncHasOnMinimized()
}

// FuncHasOnMinimized returns the function "WEBEXT.app.window.onMinimized.hasListener".
func FuncHasOnMinimized() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnMinimized(
		js.Pointer(&fn),
	)
	return
}

// HasOnMinimized calls the function "WEBEXT.app.window.onMinimized.hasListener" directly.
func HasOnMinimized(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnMinimized(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMinimized calls the function "WEBEXT.app.window.onMinimized.hasListener"
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

// HasFuncOnRestored returns true if the function "WEBEXT.app.window.onRestored.addListener" exists.
func HasFuncOnRestored() bool {
	return js.True == bindings.HasFuncOnRestored()
}

// FuncOnRestored returns the function "WEBEXT.app.window.onRestored.addListener".
func FuncOnRestored() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnRestored(
		js.Pointer(&fn),
	)
	return
}

// OnRestored calls the function "WEBEXT.app.window.onRestored.addListener" directly.
func OnRestored(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnRestored(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRestored calls the function "WEBEXT.app.window.onRestored.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRestored(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRestored(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRestored returns true if the function "WEBEXT.app.window.onRestored.removeListener" exists.
func HasFuncOffRestored() bool {
	return js.True == bindings.HasFuncOffRestored()
}

// FuncOffRestored returns the function "WEBEXT.app.window.onRestored.removeListener".
func FuncOffRestored() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffRestored(
		js.Pointer(&fn),
	)
	return
}

// OffRestored calls the function "WEBEXT.app.window.onRestored.removeListener" directly.
func OffRestored(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffRestored(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRestored calls the function "WEBEXT.app.window.onRestored.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRestored(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRestored(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRestored returns true if the function "WEBEXT.app.window.onRestored.hasListener" exists.
func HasFuncHasOnRestored() bool {
	return js.True == bindings.HasFuncHasOnRestored()
}

// FuncHasOnRestored returns the function "WEBEXT.app.window.onRestored.hasListener".
func FuncHasOnRestored() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnRestored(
		js.Pointer(&fn),
	)
	return
}

// HasOnRestored calls the function "WEBEXT.app.window.onRestored.hasListener" directly.
func HasOnRestored(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnRestored(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRestored calls the function "WEBEXT.app.window.onRestored.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRestored(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRestored(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}
