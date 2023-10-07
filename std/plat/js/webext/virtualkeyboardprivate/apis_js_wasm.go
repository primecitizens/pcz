// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package virtualkeyboardprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/virtualkeyboardprivate/bindings"
)

type Bounds struct {
	// Height is "Bounds.height"
	//
	// Required
	Height int64
	// Left is "Bounds.left"
	//
	// Required
	Left int64
	// Top is "Bounds.top"
	//
	// Required
	Top int64
	// Width is "Bounds.width"
	//
	// Required
	Width int64

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

type DisplayFormat uint32

const (
	_ DisplayFormat = iota

	DisplayFormat_TEXT
	DisplayFormat_PNG
	DisplayFormat_HTML
	DisplayFormat_FILE
)

func (DisplayFormat) FromRef(str js.Ref) DisplayFormat {
	return DisplayFormat(bindings.ConstOfDisplayFormat(str))
}

func (x DisplayFormat) String() (string, bool) {
	switch x {
	case DisplayFormat_TEXT:
		return "text", true
	case DisplayFormat_PNG:
		return "png", true
	case DisplayFormat_HTML:
		return "html", true
	case DisplayFormat_FILE:
		return "file", true
	default:
		return "", false
	}
}

type ClipboardItem struct {
	// DisplayFormat is "ClipboardItem.displayFormat"
	//
	// Required
	DisplayFormat DisplayFormat
	// Id is "ClipboardItem.id"
	//
	// Required
	Id js.String
	// ImageData is "ClipboardItem.imageData"
	//
	// Optional
	ImageData js.String
	// TextData is "ClipboardItem.textData"
	//
	// Optional
	TextData js.String
	// TimeCopied is "ClipboardItem.timeCopied"
	//
	// Required
	TimeCopied float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClipboardItem with all fields set.
func (p ClipboardItem) FromRef(ref js.Ref) ClipboardItem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ClipboardItem in the application heap.
func (p ClipboardItem) New() js.Ref {
	return bindings.ClipboardItemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ClipboardItem) UpdateFrom(ref js.Ref) {
	bindings.ClipboardItemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ClipboardItem) Update(ref js.Ref) {
	bindings.ClipboardItemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ClipboardItem) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.ImageData.Ref(),
		p.TextData.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.ImageData = p.ImageData.FromRef(js.Undefined)
	p.TextData = p.TextData.FromRef(js.Undefined)
}

type KeyboardMode uint32

const (
	_ KeyboardMode = iota

	KeyboardMode_FULL_WIDTH
	KeyboardMode_FLOATING
)

func (KeyboardMode) FromRef(str js.Ref) KeyboardMode {
	return KeyboardMode(bindings.ConstOfKeyboardMode(str))
}

func (x KeyboardMode) String() (string, bool) {
	switch x {
	case KeyboardMode_FULL_WIDTH:
		return "FULL_WIDTH", true
	case KeyboardMode_FLOATING:
		return "FLOATING", true
	default:
		return "", false
	}
}

type ContainerBehaviorOptions struct {
	// Bounds is "ContainerBehaviorOptions.bounds"
	//
	// Required
	//
	// NOTE: Bounds.FFI_USE MUST be set to true to get Bounds used.
	Bounds Bounds
	// Mode is "ContainerBehaviorOptions.mode"
	//
	// Required
	Mode KeyboardMode

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContainerBehaviorOptions with all fields set.
func (p ContainerBehaviorOptions) FromRef(ref js.Ref) ContainerBehaviorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ContainerBehaviorOptions in the application heap.
func (p ContainerBehaviorOptions) New() js.Ref {
	return bindings.ContainerBehaviorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ContainerBehaviorOptions) UpdateFrom(ref js.Ref) {
	bindings.ContainerBehaviorOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ContainerBehaviorOptions) Update(ref js.Ref) {
	bindings.ContainerBehaviorOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ContainerBehaviorOptions) FreeMembers(recursive bool) {
	if recursive {
		p.Bounds.FreeMembers(true)
	}
}

type GetClipboardHistoryArgOptions struct {
	// ItemIds is "GetClipboardHistoryArgOptions.itemIds"
	//
	// Optional
	ItemIds js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetClipboardHistoryArgOptions with all fields set.
func (p GetClipboardHistoryArgOptions) FromRef(ref js.Ref) GetClipboardHistoryArgOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetClipboardHistoryArgOptions in the application heap.
func (p GetClipboardHistoryArgOptions) New() js.Ref {
	return bindings.GetClipboardHistoryArgOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetClipboardHistoryArgOptions) UpdateFrom(ref js.Ref) {
	bindings.GetClipboardHistoryArgOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetClipboardHistoryArgOptions) Update(ref js.Ref) {
	bindings.GetClipboardHistoryArgOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetClipboardHistoryArgOptions) FreeMembers(recursive bool) {
	js.Free(
		p.ItemIds.Ref(),
	)
	p.ItemIds = p.ItemIds.FromRef(js.Undefined)
}

type KeyboardConfig struct {
	// A11ymode is "KeyboardConfig.a11ymode"
	//
	// Required
	A11ymode bool
	// Features is "KeyboardConfig.features"
	//
	// Required
	Features js.Array[js.String]
	// Hotrodmode is "KeyboardConfig.hotrodmode"
	//
	// Required
	Hotrodmode bool
	// Layout is "KeyboardConfig.layout"
	//
	// Required
	Layout js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a KeyboardConfig with all fields set.
func (p KeyboardConfig) FromRef(ref js.Ref) KeyboardConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new KeyboardConfig in the application heap.
func (p KeyboardConfig) New() js.Ref {
	return bindings.KeyboardConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *KeyboardConfig) UpdateFrom(ref js.Ref) {
	bindings.KeyboardConfigJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *KeyboardConfig) Update(ref js.Ref) {
	bindings.KeyboardConfigJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *KeyboardConfig) FreeMembers(recursive bool) {
	js.Free(
		p.Features.Ref(),
		p.Layout.Ref(),
	)
	p.Features = p.Features.FromRef(js.Undefined)
	p.Layout = p.Layout.FromRef(js.Undefined)
}

type KeyboardState uint32

const (
	_ KeyboardState = iota

	KeyboardState_ENABLED
	KeyboardState_DISABLED
	KeyboardState_AUTO
)

func (KeyboardState) FromRef(str js.Ref) KeyboardState {
	return KeyboardState(bindings.ConstOfKeyboardState(str))
}

func (x KeyboardState) String() (string, bool) {
	switch x {
	case KeyboardState_ENABLED:
		return "ENABLED", true
	case KeyboardState_DISABLED:
		return "DISABLED", true
	case KeyboardState_AUTO:
		return "AUTO", true
	default:
		return "", false
	}
}

type VirtualKeyboardEventType uint32

const (
	_ VirtualKeyboardEventType = iota

	VirtualKeyboardEventType_KEYUP
	VirtualKeyboardEventType_KEYDOWN
)

func (VirtualKeyboardEventType) FromRef(str js.Ref) VirtualKeyboardEventType {
	return VirtualKeyboardEventType(bindings.ConstOfVirtualKeyboardEventType(str))
}

func (x VirtualKeyboardEventType) String() (string, bool) {
	switch x {
	case VirtualKeyboardEventType_KEYUP:
		return "keyup", true
	case VirtualKeyboardEventType_KEYDOWN:
		return "keydown", true
	default:
		return "", false
	}
}

type VirtualKeyboardEvent struct {
	// CharValue is "VirtualKeyboardEvent.charValue"
	//
	// Required
	CharValue int64
	// KeyCode is "VirtualKeyboardEvent.keyCode"
	//
	// Required
	KeyCode int64
	// KeyName is "VirtualKeyboardEvent.keyName"
	//
	// Required
	KeyName js.String
	// Modifiers is "VirtualKeyboardEvent.modifiers"
	//
	// Optional
	//
	// NOTE: FFI_USE_Modifiers MUST be set to true to make this field effective.
	Modifiers int64
	// Type is "VirtualKeyboardEvent.type"
	//
	// Required
	Type VirtualKeyboardEventType

	FFI_USE_Modifiers bool // for Modifiers.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VirtualKeyboardEvent with all fields set.
func (p VirtualKeyboardEvent) FromRef(ref js.Ref) VirtualKeyboardEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VirtualKeyboardEvent in the application heap.
func (p VirtualKeyboardEvent) New() js.Ref {
	return bindings.VirtualKeyboardEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *VirtualKeyboardEvent) UpdateFrom(ref js.Ref) {
	bindings.VirtualKeyboardEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VirtualKeyboardEvent) Update(ref js.Ref) {
	bindings.VirtualKeyboardEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VirtualKeyboardEvent) FreeMembers(recursive bool) {
	js.Free(
		p.KeyName.Ref(),
	)
	p.KeyName = p.KeyName.FromRef(js.Undefined)
}

// HasFuncDeleteClipboardItem returns true if the function "WEBEXT.virtualKeyboardPrivate.deleteClipboardItem" exists.
func HasFuncDeleteClipboardItem() bool {
	return js.True == bindings.HasFuncDeleteClipboardItem()
}

// FuncDeleteClipboardItem returns the function "WEBEXT.virtualKeyboardPrivate.deleteClipboardItem".
func FuncDeleteClipboardItem() (fn js.Func[func(itemId js.String)]) {
	bindings.FuncDeleteClipboardItem(
		js.Pointer(&fn),
	)
	return
}

// DeleteClipboardItem calls the function "WEBEXT.virtualKeyboardPrivate.deleteClipboardItem" directly.
func DeleteClipboardItem(itemId js.String) (ret js.Void) {
	bindings.CallDeleteClipboardItem(
		js.Pointer(&ret),
		itemId.Ref(),
	)

	return
}

// TryDeleteClipboardItem calls the function "WEBEXT.virtualKeyboardPrivate.deleteClipboardItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDeleteClipboardItem(itemId js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeleteClipboardItem(
		js.Pointer(&ret), js.Pointer(&exception),
		itemId.Ref(),
	)

	return
}

// HasFuncGetClipboardHistory returns true if the function "WEBEXT.virtualKeyboardPrivate.getClipboardHistory" exists.
func HasFuncGetClipboardHistory() bool {
	return js.True == bindings.HasFuncGetClipboardHistory()
}

// FuncGetClipboardHistory returns the function "WEBEXT.virtualKeyboardPrivate.getClipboardHistory".
func FuncGetClipboardHistory() (fn js.Func[func(options GetClipboardHistoryArgOptions) js.Promise[js.Array[ClipboardItem]]]) {
	bindings.FuncGetClipboardHistory(
		js.Pointer(&fn),
	)
	return
}

// GetClipboardHistory calls the function "WEBEXT.virtualKeyboardPrivate.getClipboardHistory" directly.
func GetClipboardHistory(options GetClipboardHistoryArgOptions) (ret js.Promise[js.Array[ClipboardItem]]) {
	bindings.CallGetClipboardHistory(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetClipboardHistory calls the function "WEBEXT.virtualKeyboardPrivate.getClipboardHistory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetClipboardHistory(options GetClipboardHistoryArgOptions) (ret js.Promise[js.Array[ClipboardItem]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetClipboardHistory(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetKeyboardConfig returns true if the function "WEBEXT.virtualKeyboardPrivate.getKeyboardConfig" exists.
func HasFuncGetKeyboardConfig() bool {
	return js.True == bindings.HasFuncGetKeyboardConfig()
}

// FuncGetKeyboardConfig returns the function "WEBEXT.virtualKeyboardPrivate.getKeyboardConfig".
func FuncGetKeyboardConfig() (fn js.Func[func() js.Promise[KeyboardConfig]]) {
	bindings.FuncGetKeyboardConfig(
		js.Pointer(&fn),
	)
	return
}

// GetKeyboardConfig calls the function "WEBEXT.virtualKeyboardPrivate.getKeyboardConfig" directly.
func GetKeyboardConfig() (ret js.Promise[KeyboardConfig]) {
	bindings.CallGetKeyboardConfig(
		js.Pointer(&ret),
	)

	return
}

// TryGetKeyboardConfig calls the function "WEBEXT.virtualKeyboardPrivate.getKeyboardConfig"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetKeyboardConfig() (ret js.Promise[KeyboardConfig], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetKeyboardConfig(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncHideKeyboard returns true if the function "WEBEXT.virtualKeyboardPrivate.hideKeyboard" exists.
func HasFuncHideKeyboard() bool {
	return js.True == bindings.HasFuncHideKeyboard()
}

// FuncHideKeyboard returns the function "WEBEXT.virtualKeyboardPrivate.hideKeyboard".
func FuncHideKeyboard() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncHideKeyboard(
		js.Pointer(&fn),
	)
	return
}

// HideKeyboard calls the function "WEBEXT.virtualKeyboardPrivate.hideKeyboard" directly.
func HideKeyboard() (ret js.Promise[js.Void]) {
	bindings.CallHideKeyboard(
		js.Pointer(&ret),
	)

	return
}

// TryHideKeyboard calls the function "WEBEXT.virtualKeyboardPrivate.hideKeyboard"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHideKeyboard() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHideKeyboard(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInsertText returns true if the function "WEBEXT.virtualKeyboardPrivate.insertText" exists.
func HasFuncInsertText() bool {
	return js.True == bindings.HasFuncInsertText()
}

// FuncInsertText returns the function "WEBEXT.virtualKeyboardPrivate.insertText".
func FuncInsertText() (fn js.Func[func(text js.String) js.Promise[js.Void]]) {
	bindings.FuncInsertText(
		js.Pointer(&fn),
	)
	return
}

// InsertText calls the function "WEBEXT.virtualKeyboardPrivate.insertText" directly.
func InsertText(text js.String) (ret js.Promise[js.Void]) {
	bindings.CallInsertText(
		js.Pointer(&ret),
		text.Ref(),
	)

	return
}

// TryInsertText calls the function "WEBEXT.virtualKeyboardPrivate.insertText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInsertText(text js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInsertText(
		js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
	)

	return
}

// HasFuncKeyboardLoaded returns true if the function "WEBEXT.virtualKeyboardPrivate.keyboardLoaded" exists.
func HasFuncKeyboardLoaded() bool {
	return js.True == bindings.HasFuncKeyboardLoaded()
}

// FuncKeyboardLoaded returns the function "WEBEXT.virtualKeyboardPrivate.keyboardLoaded".
func FuncKeyboardLoaded() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncKeyboardLoaded(
		js.Pointer(&fn),
	)
	return
}

// KeyboardLoaded calls the function "WEBEXT.virtualKeyboardPrivate.keyboardLoaded" directly.
func KeyboardLoaded() (ret js.Promise[js.Void]) {
	bindings.CallKeyboardLoaded(
		js.Pointer(&ret),
	)

	return
}

// TryKeyboardLoaded calls the function "WEBEXT.virtualKeyboardPrivate.keyboardLoaded"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryKeyboardLoaded() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardLoaded(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLockKeyboard returns true if the function "WEBEXT.virtualKeyboardPrivate.lockKeyboard" exists.
func HasFuncLockKeyboard() bool {
	return js.True == bindings.HasFuncLockKeyboard()
}

// FuncLockKeyboard returns the function "WEBEXT.virtualKeyboardPrivate.lockKeyboard".
func FuncLockKeyboard() (fn js.Func[func(lock bool)]) {
	bindings.FuncLockKeyboard(
		js.Pointer(&fn),
	)
	return
}

// LockKeyboard calls the function "WEBEXT.virtualKeyboardPrivate.lockKeyboard" directly.
func LockKeyboard(lock bool) (ret js.Void) {
	bindings.CallLockKeyboard(
		js.Pointer(&ret),
		js.Bool(bool(lock)),
	)

	return
}

// TryLockKeyboard calls the function "WEBEXT.virtualKeyboardPrivate.lockKeyboard"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLockKeyboard(lock bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLockKeyboard(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(lock)),
	)

	return
}

type OnBoundsChangedEventCallbackFunc func(this js.Ref, bounds *Bounds) js.Ref

func (fn OnBoundsChangedEventCallbackFunc) Register() js.Func[func(bounds *Bounds)] {
	return js.RegisterCallback[func(bounds *Bounds)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnBoundsChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Bounds
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

type OnBoundsChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, bounds *Bounds) js.Ref
	Arg T
}

func (cb *OnBoundsChangedEventCallback[T]) Register() js.Func[func(bounds *Bounds)] {
	return js.RegisterCallback[func(bounds *Bounds)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnBoundsChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Bounds
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

// HasFuncOnBoundsChanged returns true if the function "WEBEXT.virtualKeyboardPrivate.onBoundsChanged.addListener" exists.
func HasFuncOnBoundsChanged() bool {
	return js.True == bindings.HasFuncOnBoundsChanged()
}

// FuncOnBoundsChanged returns the function "WEBEXT.virtualKeyboardPrivate.onBoundsChanged.addListener".
func FuncOnBoundsChanged() (fn js.Func[func(callback js.Func[func(bounds *Bounds)])]) {
	bindings.FuncOnBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// OnBoundsChanged calls the function "WEBEXT.virtualKeyboardPrivate.onBoundsChanged.addListener" directly.
func OnBoundsChanged(callback js.Func[func(bounds *Bounds)]) (ret js.Void) {
	bindings.CallOnBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnBoundsChanged calls the function "WEBEXT.virtualKeyboardPrivate.onBoundsChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnBoundsChanged(callback js.Func[func(bounds *Bounds)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffBoundsChanged returns true if the function "WEBEXT.virtualKeyboardPrivate.onBoundsChanged.removeListener" exists.
func HasFuncOffBoundsChanged() bool {
	return js.True == bindings.HasFuncOffBoundsChanged()
}

// FuncOffBoundsChanged returns the function "WEBEXT.virtualKeyboardPrivate.onBoundsChanged.removeListener".
func FuncOffBoundsChanged() (fn js.Func[func(callback js.Func[func(bounds *Bounds)])]) {
	bindings.FuncOffBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// OffBoundsChanged calls the function "WEBEXT.virtualKeyboardPrivate.onBoundsChanged.removeListener" directly.
func OffBoundsChanged(callback js.Func[func(bounds *Bounds)]) (ret js.Void) {
	bindings.CallOffBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffBoundsChanged calls the function "WEBEXT.virtualKeyboardPrivate.onBoundsChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffBoundsChanged(callback js.Func[func(bounds *Bounds)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnBoundsChanged returns true if the function "WEBEXT.virtualKeyboardPrivate.onBoundsChanged.hasListener" exists.
func HasFuncHasOnBoundsChanged() bool {
	return js.True == bindings.HasFuncHasOnBoundsChanged()
}

// FuncHasOnBoundsChanged returns the function "WEBEXT.virtualKeyboardPrivate.onBoundsChanged.hasListener".
func FuncHasOnBoundsChanged() (fn js.Func[func(callback js.Func[func(bounds *Bounds)]) bool]) {
	bindings.FuncHasOnBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnBoundsChanged calls the function "WEBEXT.virtualKeyboardPrivate.onBoundsChanged.hasListener" directly.
func HasOnBoundsChanged(callback js.Func[func(bounds *Bounds)]) (ret bool) {
	bindings.CallHasOnBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnBoundsChanged calls the function "WEBEXT.virtualKeyboardPrivate.onBoundsChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnBoundsChanged(callback js.Func[func(bounds *Bounds)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnClipboardHistoryChangedEventCallbackFunc func(this js.Ref, itemIds js.Array[js.String]) js.Ref

func (fn OnClipboardHistoryChangedEventCallbackFunc) Register() js.Func[func(itemIds js.Array[js.String])] {
	return js.RegisterCallback[func(itemIds js.Array[js.String])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnClipboardHistoryChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnClipboardHistoryChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, itemIds js.Array[js.String]) js.Ref
	Arg T
}

func (cb *OnClipboardHistoryChangedEventCallback[T]) Register() js.Func[func(itemIds js.Array[js.String])] {
	return js.RegisterCallback[func(itemIds js.Array[js.String])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnClipboardHistoryChangedEventCallback[T]) DispatchCallback(
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

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnClipboardHistoryChanged returns true if the function "WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.addListener" exists.
func HasFuncOnClipboardHistoryChanged() bool {
	return js.True == bindings.HasFuncOnClipboardHistoryChanged()
}

// FuncOnClipboardHistoryChanged returns the function "WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.addListener".
func FuncOnClipboardHistoryChanged() (fn js.Func[func(callback js.Func[func(itemIds js.Array[js.String])])]) {
	bindings.FuncOnClipboardHistoryChanged(
		js.Pointer(&fn),
	)
	return
}

// OnClipboardHistoryChanged calls the function "WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.addListener" directly.
func OnClipboardHistoryChanged(callback js.Func[func(itemIds js.Array[js.String])]) (ret js.Void) {
	bindings.CallOnClipboardHistoryChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnClipboardHistoryChanged calls the function "WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnClipboardHistoryChanged(callback js.Func[func(itemIds js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnClipboardHistoryChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffClipboardHistoryChanged returns true if the function "WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.removeListener" exists.
func HasFuncOffClipboardHistoryChanged() bool {
	return js.True == bindings.HasFuncOffClipboardHistoryChanged()
}

// FuncOffClipboardHistoryChanged returns the function "WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.removeListener".
func FuncOffClipboardHistoryChanged() (fn js.Func[func(callback js.Func[func(itemIds js.Array[js.String])])]) {
	bindings.FuncOffClipboardHistoryChanged(
		js.Pointer(&fn),
	)
	return
}

// OffClipboardHistoryChanged calls the function "WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.removeListener" directly.
func OffClipboardHistoryChanged(callback js.Func[func(itemIds js.Array[js.String])]) (ret js.Void) {
	bindings.CallOffClipboardHistoryChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffClipboardHistoryChanged calls the function "WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffClipboardHistoryChanged(callback js.Func[func(itemIds js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffClipboardHistoryChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnClipboardHistoryChanged returns true if the function "WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.hasListener" exists.
func HasFuncHasOnClipboardHistoryChanged() bool {
	return js.True == bindings.HasFuncHasOnClipboardHistoryChanged()
}

// FuncHasOnClipboardHistoryChanged returns the function "WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.hasListener".
func FuncHasOnClipboardHistoryChanged() (fn js.Func[func(callback js.Func[func(itemIds js.Array[js.String])]) bool]) {
	bindings.FuncHasOnClipboardHistoryChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnClipboardHistoryChanged calls the function "WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.hasListener" directly.
func HasOnClipboardHistoryChanged(callback js.Func[func(itemIds js.Array[js.String])]) (ret bool) {
	bindings.CallHasOnClipboardHistoryChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnClipboardHistoryChanged calls the function "WEBEXT.virtualKeyboardPrivate.onClipboardHistoryChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnClipboardHistoryChanged(callback js.Func[func(itemIds js.Array[js.String])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnClipboardHistoryChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnClipboardItemUpdatedEventCallbackFunc func(this js.Ref, clipboardHistoryItem *ClipboardItem) js.Ref

func (fn OnClipboardItemUpdatedEventCallbackFunc) Register() js.Func[func(clipboardHistoryItem *ClipboardItem)] {
	return js.RegisterCallback[func(clipboardHistoryItem *ClipboardItem)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnClipboardItemUpdatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ClipboardItem
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

type OnClipboardItemUpdatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, clipboardHistoryItem *ClipboardItem) js.Ref
	Arg T
}

func (cb *OnClipboardItemUpdatedEventCallback[T]) Register() js.Func[func(clipboardHistoryItem *ClipboardItem)] {
	return js.RegisterCallback[func(clipboardHistoryItem *ClipboardItem)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnClipboardItemUpdatedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ClipboardItem
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

// HasFuncOnClipboardItemUpdated returns true if the function "WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.addListener" exists.
func HasFuncOnClipboardItemUpdated() bool {
	return js.True == bindings.HasFuncOnClipboardItemUpdated()
}

// FuncOnClipboardItemUpdated returns the function "WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.addListener".
func FuncOnClipboardItemUpdated() (fn js.Func[func(callback js.Func[func(clipboardHistoryItem *ClipboardItem)])]) {
	bindings.FuncOnClipboardItemUpdated(
		js.Pointer(&fn),
	)
	return
}

// OnClipboardItemUpdated calls the function "WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.addListener" directly.
func OnClipboardItemUpdated(callback js.Func[func(clipboardHistoryItem *ClipboardItem)]) (ret js.Void) {
	bindings.CallOnClipboardItemUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnClipboardItemUpdated calls the function "WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnClipboardItemUpdated(callback js.Func[func(clipboardHistoryItem *ClipboardItem)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnClipboardItemUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffClipboardItemUpdated returns true if the function "WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.removeListener" exists.
func HasFuncOffClipboardItemUpdated() bool {
	return js.True == bindings.HasFuncOffClipboardItemUpdated()
}

// FuncOffClipboardItemUpdated returns the function "WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.removeListener".
func FuncOffClipboardItemUpdated() (fn js.Func[func(callback js.Func[func(clipboardHistoryItem *ClipboardItem)])]) {
	bindings.FuncOffClipboardItemUpdated(
		js.Pointer(&fn),
	)
	return
}

// OffClipboardItemUpdated calls the function "WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.removeListener" directly.
func OffClipboardItemUpdated(callback js.Func[func(clipboardHistoryItem *ClipboardItem)]) (ret js.Void) {
	bindings.CallOffClipboardItemUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffClipboardItemUpdated calls the function "WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffClipboardItemUpdated(callback js.Func[func(clipboardHistoryItem *ClipboardItem)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffClipboardItemUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnClipboardItemUpdated returns true if the function "WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.hasListener" exists.
func HasFuncHasOnClipboardItemUpdated() bool {
	return js.True == bindings.HasFuncHasOnClipboardItemUpdated()
}

// FuncHasOnClipboardItemUpdated returns the function "WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.hasListener".
func FuncHasOnClipboardItemUpdated() (fn js.Func[func(callback js.Func[func(clipboardHistoryItem *ClipboardItem)]) bool]) {
	bindings.FuncHasOnClipboardItemUpdated(
		js.Pointer(&fn),
	)
	return
}

// HasOnClipboardItemUpdated calls the function "WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.hasListener" directly.
func HasOnClipboardItemUpdated(callback js.Func[func(clipboardHistoryItem *ClipboardItem)]) (ret bool) {
	bindings.CallHasOnClipboardItemUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnClipboardItemUpdated calls the function "WEBEXT.virtualKeyboardPrivate.onClipboardItemUpdated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnClipboardItemUpdated(callback js.Func[func(clipboardHistoryItem *ClipboardItem)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnClipboardItemUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnColorProviderChangedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnColorProviderChangedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnColorProviderChangedEventCallbackFunc) DispatchCallback(
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

type OnColorProviderChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnColorProviderChangedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnColorProviderChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnColorProviderChanged returns true if the function "WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.addListener" exists.
func HasFuncOnColorProviderChanged() bool {
	return js.True == bindings.HasFuncOnColorProviderChanged()
}

// FuncOnColorProviderChanged returns the function "WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.addListener".
func FuncOnColorProviderChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnColorProviderChanged(
		js.Pointer(&fn),
	)
	return
}

// OnColorProviderChanged calls the function "WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.addListener" directly.
func OnColorProviderChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnColorProviderChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnColorProviderChanged calls the function "WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnColorProviderChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnColorProviderChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffColorProviderChanged returns true if the function "WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.removeListener" exists.
func HasFuncOffColorProviderChanged() bool {
	return js.True == bindings.HasFuncOffColorProviderChanged()
}

// FuncOffColorProviderChanged returns the function "WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.removeListener".
func FuncOffColorProviderChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffColorProviderChanged(
		js.Pointer(&fn),
	)
	return
}

// OffColorProviderChanged calls the function "WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.removeListener" directly.
func OffColorProviderChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffColorProviderChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffColorProviderChanged calls the function "WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffColorProviderChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffColorProviderChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnColorProviderChanged returns true if the function "WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.hasListener" exists.
func HasFuncHasOnColorProviderChanged() bool {
	return js.True == bindings.HasFuncHasOnColorProviderChanged()
}

// FuncHasOnColorProviderChanged returns the function "WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.hasListener".
func FuncHasOnColorProviderChanged() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnColorProviderChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnColorProviderChanged calls the function "WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.hasListener" directly.
func HasOnColorProviderChanged(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnColorProviderChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnColorProviderChanged calls the function "WEBEXT.virtualKeyboardPrivate.onColorProviderChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnColorProviderChanged(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnColorProviderChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnKeyboardClosedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnKeyboardClosedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnKeyboardClosedEventCallbackFunc) DispatchCallback(
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

type OnKeyboardClosedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnKeyboardClosedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnKeyboardClosedEventCallback[T]) DispatchCallback(
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

// HasFuncOnKeyboardClosed returns true if the function "WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.addListener" exists.
func HasFuncOnKeyboardClosed() bool {
	return js.True == bindings.HasFuncOnKeyboardClosed()
}

// FuncOnKeyboardClosed returns the function "WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.addListener".
func FuncOnKeyboardClosed() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnKeyboardClosed(
		js.Pointer(&fn),
	)
	return
}

// OnKeyboardClosed calls the function "WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.addListener" directly.
func OnKeyboardClosed(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnKeyboardClosed(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnKeyboardClosed calls the function "WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnKeyboardClosed(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnKeyboardClosed(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffKeyboardClosed returns true if the function "WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.removeListener" exists.
func HasFuncOffKeyboardClosed() bool {
	return js.True == bindings.HasFuncOffKeyboardClosed()
}

// FuncOffKeyboardClosed returns the function "WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.removeListener".
func FuncOffKeyboardClosed() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffKeyboardClosed(
		js.Pointer(&fn),
	)
	return
}

// OffKeyboardClosed calls the function "WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.removeListener" directly.
func OffKeyboardClosed(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffKeyboardClosed(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffKeyboardClosed calls the function "WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffKeyboardClosed(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffKeyboardClosed(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnKeyboardClosed returns true if the function "WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.hasListener" exists.
func HasFuncHasOnKeyboardClosed() bool {
	return js.True == bindings.HasFuncHasOnKeyboardClosed()
}

// FuncHasOnKeyboardClosed returns the function "WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.hasListener".
func FuncHasOnKeyboardClosed() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnKeyboardClosed(
		js.Pointer(&fn),
	)
	return
}

// HasOnKeyboardClosed calls the function "WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.hasListener" directly.
func HasOnKeyboardClosed(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnKeyboardClosed(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnKeyboardClosed calls the function "WEBEXT.virtualKeyboardPrivate.onKeyboardClosed.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnKeyboardClosed(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnKeyboardClosed(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnKeyboardConfigChangedEventCallbackFunc func(this js.Ref, config *KeyboardConfig) js.Ref

func (fn OnKeyboardConfigChangedEventCallbackFunc) Register() js.Func[func(config *KeyboardConfig)] {
	return js.RegisterCallback[func(config *KeyboardConfig)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnKeyboardConfigChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 KeyboardConfig
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

type OnKeyboardConfigChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, config *KeyboardConfig) js.Ref
	Arg T
}

func (cb *OnKeyboardConfigChangedEventCallback[T]) Register() js.Func[func(config *KeyboardConfig)] {
	return js.RegisterCallback[func(config *KeyboardConfig)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnKeyboardConfigChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 KeyboardConfig
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

// HasFuncOnKeyboardConfigChanged returns true if the function "WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.addListener" exists.
func HasFuncOnKeyboardConfigChanged() bool {
	return js.True == bindings.HasFuncOnKeyboardConfigChanged()
}

// FuncOnKeyboardConfigChanged returns the function "WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.addListener".
func FuncOnKeyboardConfigChanged() (fn js.Func[func(callback js.Func[func(config *KeyboardConfig)])]) {
	bindings.FuncOnKeyboardConfigChanged(
		js.Pointer(&fn),
	)
	return
}

// OnKeyboardConfigChanged calls the function "WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.addListener" directly.
func OnKeyboardConfigChanged(callback js.Func[func(config *KeyboardConfig)]) (ret js.Void) {
	bindings.CallOnKeyboardConfigChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnKeyboardConfigChanged calls the function "WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnKeyboardConfigChanged(callback js.Func[func(config *KeyboardConfig)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnKeyboardConfigChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffKeyboardConfigChanged returns true if the function "WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.removeListener" exists.
func HasFuncOffKeyboardConfigChanged() bool {
	return js.True == bindings.HasFuncOffKeyboardConfigChanged()
}

// FuncOffKeyboardConfigChanged returns the function "WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.removeListener".
func FuncOffKeyboardConfigChanged() (fn js.Func[func(callback js.Func[func(config *KeyboardConfig)])]) {
	bindings.FuncOffKeyboardConfigChanged(
		js.Pointer(&fn),
	)
	return
}

// OffKeyboardConfigChanged calls the function "WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.removeListener" directly.
func OffKeyboardConfigChanged(callback js.Func[func(config *KeyboardConfig)]) (ret js.Void) {
	bindings.CallOffKeyboardConfigChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffKeyboardConfigChanged calls the function "WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffKeyboardConfigChanged(callback js.Func[func(config *KeyboardConfig)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffKeyboardConfigChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnKeyboardConfigChanged returns true if the function "WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.hasListener" exists.
func HasFuncHasOnKeyboardConfigChanged() bool {
	return js.True == bindings.HasFuncHasOnKeyboardConfigChanged()
}

// FuncHasOnKeyboardConfigChanged returns the function "WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.hasListener".
func FuncHasOnKeyboardConfigChanged() (fn js.Func[func(callback js.Func[func(config *KeyboardConfig)]) bool]) {
	bindings.FuncHasOnKeyboardConfigChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnKeyboardConfigChanged calls the function "WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.hasListener" directly.
func HasOnKeyboardConfigChanged(callback js.Func[func(config *KeyboardConfig)]) (ret bool) {
	bindings.CallHasOnKeyboardConfigChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnKeyboardConfigChanged calls the function "WEBEXT.virtualKeyboardPrivate.onKeyboardConfigChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnKeyboardConfigChanged(callback js.Func[func(config *KeyboardConfig)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnKeyboardConfigChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOpenSettings returns true if the function "WEBEXT.virtualKeyboardPrivate.openSettings" exists.
func HasFuncOpenSettings() bool {
	return js.True == bindings.HasFuncOpenSettings()
}

// FuncOpenSettings returns the function "WEBEXT.virtualKeyboardPrivate.openSettings".
func FuncOpenSettings() (fn js.Func[func()]) {
	bindings.FuncOpenSettings(
		js.Pointer(&fn),
	)
	return
}

// OpenSettings calls the function "WEBEXT.virtualKeyboardPrivate.openSettings" directly.
func OpenSettings() (ret js.Void) {
	bindings.CallOpenSettings(
		js.Pointer(&ret),
	)

	return
}

// TryOpenSettings calls the function "WEBEXT.virtualKeyboardPrivate.openSettings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenSettings() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenSettings(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncOpenSuggestionSettings returns true if the function "WEBEXT.virtualKeyboardPrivate.openSuggestionSettings" exists.
func HasFuncOpenSuggestionSettings() bool {
	return js.True == bindings.HasFuncOpenSuggestionSettings()
}

// FuncOpenSuggestionSettings returns the function "WEBEXT.virtualKeyboardPrivate.openSuggestionSettings".
func FuncOpenSuggestionSettings() (fn js.Func[func()]) {
	bindings.FuncOpenSuggestionSettings(
		js.Pointer(&fn),
	)
	return
}

// OpenSuggestionSettings calls the function "WEBEXT.virtualKeyboardPrivate.openSuggestionSettings" directly.
func OpenSuggestionSettings() (ret js.Void) {
	bindings.CallOpenSuggestionSettings(
		js.Pointer(&ret),
	)

	return
}

// TryOpenSuggestionSettings calls the function "WEBEXT.virtualKeyboardPrivate.openSuggestionSettings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenSuggestionSettings() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenSuggestionSettings(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPasteClipboardItem returns true if the function "WEBEXT.virtualKeyboardPrivate.pasteClipboardItem" exists.
func HasFuncPasteClipboardItem() bool {
	return js.True == bindings.HasFuncPasteClipboardItem()
}

// FuncPasteClipboardItem returns the function "WEBEXT.virtualKeyboardPrivate.pasteClipboardItem".
func FuncPasteClipboardItem() (fn js.Func[func(itemId js.String)]) {
	bindings.FuncPasteClipboardItem(
		js.Pointer(&fn),
	)
	return
}

// PasteClipboardItem calls the function "WEBEXT.virtualKeyboardPrivate.pasteClipboardItem" directly.
func PasteClipboardItem(itemId js.String) (ret js.Void) {
	bindings.CallPasteClipboardItem(
		js.Pointer(&ret),
		itemId.Ref(),
	)

	return
}

// TryPasteClipboardItem calls the function "WEBEXT.virtualKeyboardPrivate.pasteClipboardItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryPasteClipboardItem(itemId js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPasteClipboardItem(
		js.Pointer(&ret), js.Pointer(&exception),
		itemId.Ref(),
	)

	return
}

// HasFuncSendKeyEvent returns true if the function "WEBEXT.virtualKeyboardPrivate.sendKeyEvent" exists.
func HasFuncSendKeyEvent() bool {
	return js.True == bindings.HasFuncSendKeyEvent()
}

// FuncSendKeyEvent returns the function "WEBEXT.virtualKeyboardPrivate.sendKeyEvent".
func FuncSendKeyEvent() (fn js.Func[func(keyEvent VirtualKeyboardEvent) js.Promise[js.Void]]) {
	bindings.FuncSendKeyEvent(
		js.Pointer(&fn),
	)
	return
}

// SendKeyEvent calls the function "WEBEXT.virtualKeyboardPrivate.sendKeyEvent" directly.
func SendKeyEvent(keyEvent VirtualKeyboardEvent) (ret js.Promise[js.Void]) {
	bindings.CallSendKeyEvent(
		js.Pointer(&ret),
		js.Pointer(&keyEvent),
	)

	return
}

// TrySendKeyEvent calls the function "WEBEXT.virtualKeyboardPrivate.sendKeyEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendKeyEvent(keyEvent VirtualKeyboardEvent) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendKeyEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&keyEvent),
	)

	return
}

// HasFuncSetAreaToRemainOnScreen returns true if the function "WEBEXT.virtualKeyboardPrivate.setAreaToRemainOnScreen" exists.
func HasFuncSetAreaToRemainOnScreen() bool {
	return js.True == bindings.HasFuncSetAreaToRemainOnScreen()
}

// FuncSetAreaToRemainOnScreen returns the function "WEBEXT.virtualKeyboardPrivate.setAreaToRemainOnScreen".
func FuncSetAreaToRemainOnScreen() (fn js.Func[func(bounds Bounds)]) {
	bindings.FuncSetAreaToRemainOnScreen(
		js.Pointer(&fn),
	)
	return
}

// SetAreaToRemainOnScreen calls the function "WEBEXT.virtualKeyboardPrivate.setAreaToRemainOnScreen" directly.
func SetAreaToRemainOnScreen(bounds Bounds) (ret js.Void) {
	bindings.CallSetAreaToRemainOnScreen(
		js.Pointer(&ret),
		js.Pointer(&bounds),
	)

	return
}

// TrySetAreaToRemainOnScreen calls the function "WEBEXT.virtualKeyboardPrivate.setAreaToRemainOnScreen"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetAreaToRemainOnScreen(bounds Bounds) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetAreaToRemainOnScreen(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&bounds),
	)

	return
}

// HasFuncSetContainerBehavior returns true if the function "WEBEXT.virtualKeyboardPrivate.setContainerBehavior" exists.
func HasFuncSetContainerBehavior() bool {
	return js.True == bindings.HasFuncSetContainerBehavior()
}

// FuncSetContainerBehavior returns the function "WEBEXT.virtualKeyboardPrivate.setContainerBehavior".
func FuncSetContainerBehavior() (fn js.Func[func(options ContainerBehaviorOptions) js.Promise[js.Boolean]]) {
	bindings.FuncSetContainerBehavior(
		js.Pointer(&fn),
	)
	return
}

// SetContainerBehavior calls the function "WEBEXT.virtualKeyboardPrivate.setContainerBehavior" directly.
func SetContainerBehavior(options ContainerBehaviorOptions) (ret js.Promise[js.Boolean]) {
	bindings.CallSetContainerBehavior(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TrySetContainerBehavior calls the function "WEBEXT.virtualKeyboardPrivate.setContainerBehavior"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetContainerBehavior(options ContainerBehaviorOptions) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetContainerBehavior(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncSetDraggableArea returns true if the function "WEBEXT.virtualKeyboardPrivate.setDraggableArea" exists.
func HasFuncSetDraggableArea() bool {
	return js.True == bindings.HasFuncSetDraggableArea()
}

// FuncSetDraggableArea returns the function "WEBEXT.virtualKeyboardPrivate.setDraggableArea".
func FuncSetDraggableArea() (fn js.Func[func(bounds Bounds)]) {
	bindings.FuncSetDraggableArea(
		js.Pointer(&fn),
	)
	return
}

// SetDraggableArea calls the function "WEBEXT.virtualKeyboardPrivate.setDraggableArea" directly.
func SetDraggableArea(bounds Bounds) (ret js.Void) {
	bindings.CallSetDraggableArea(
		js.Pointer(&ret),
		js.Pointer(&bounds),
	)

	return
}

// TrySetDraggableArea calls the function "WEBEXT.virtualKeyboardPrivate.setDraggableArea"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetDraggableArea(bounds Bounds) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetDraggableArea(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&bounds),
	)

	return
}

// HasFuncSetHitTestBounds returns true if the function "WEBEXT.virtualKeyboardPrivate.setHitTestBounds" exists.
func HasFuncSetHitTestBounds() bool {
	return js.True == bindings.HasFuncSetHitTestBounds()
}

// FuncSetHitTestBounds returns the function "WEBEXT.virtualKeyboardPrivate.setHitTestBounds".
func FuncSetHitTestBounds() (fn js.Func[func(boundsList js.Array[Bounds])]) {
	bindings.FuncSetHitTestBounds(
		js.Pointer(&fn),
	)
	return
}

// SetHitTestBounds calls the function "WEBEXT.virtualKeyboardPrivate.setHitTestBounds" directly.
func SetHitTestBounds(boundsList js.Array[Bounds]) (ret js.Void) {
	bindings.CallSetHitTestBounds(
		js.Pointer(&ret),
		boundsList.Ref(),
	)

	return
}

// TrySetHitTestBounds calls the function "WEBEXT.virtualKeyboardPrivate.setHitTestBounds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetHitTestBounds(boundsList js.Array[Bounds]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetHitTestBounds(
		js.Pointer(&ret), js.Pointer(&exception),
		boundsList.Ref(),
	)

	return
}

// HasFuncSetHotrodKeyboard returns true if the function "WEBEXT.virtualKeyboardPrivate.setHotrodKeyboard" exists.
func HasFuncSetHotrodKeyboard() bool {
	return js.True == bindings.HasFuncSetHotrodKeyboard()
}

// FuncSetHotrodKeyboard returns the function "WEBEXT.virtualKeyboardPrivate.setHotrodKeyboard".
func FuncSetHotrodKeyboard() (fn js.Func[func(enable bool)]) {
	bindings.FuncSetHotrodKeyboard(
		js.Pointer(&fn),
	)
	return
}

// SetHotrodKeyboard calls the function "WEBEXT.virtualKeyboardPrivate.setHotrodKeyboard" directly.
func SetHotrodKeyboard(enable bool) (ret js.Void) {
	bindings.CallSetHotrodKeyboard(
		js.Pointer(&ret),
		js.Bool(bool(enable)),
	)

	return
}

// TrySetHotrodKeyboard calls the function "WEBEXT.virtualKeyboardPrivate.setHotrodKeyboard"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetHotrodKeyboard(enable bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetHotrodKeyboard(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enable)),
	)

	return
}

// HasFuncSetKeyboardState returns true if the function "WEBEXT.virtualKeyboardPrivate.setKeyboardState" exists.
func HasFuncSetKeyboardState() bool {
	return js.True == bindings.HasFuncSetKeyboardState()
}

// FuncSetKeyboardState returns the function "WEBEXT.virtualKeyboardPrivate.setKeyboardState".
func FuncSetKeyboardState() (fn js.Func[func(state KeyboardState)]) {
	bindings.FuncSetKeyboardState(
		js.Pointer(&fn),
	)
	return
}

// SetKeyboardState calls the function "WEBEXT.virtualKeyboardPrivate.setKeyboardState" directly.
func SetKeyboardState(state KeyboardState) (ret js.Void) {
	bindings.CallSetKeyboardState(
		js.Pointer(&ret),
		uint32(state),
	)

	return
}

// TrySetKeyboardState calls the function "WEBEXT.virtualKeyboardPrivate.setKeyboardState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetKeyboardState(state KeyboardState) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetKeyboardState(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(state),
	)

	return
}

// HasFuncSetOccludedBounds returns true if the function "WEBEXT.virtualKeyboardPrivate.setOccludedBounds" exists.
func HasFuncSetOccludedBounds() bool {
	return js.True == bindings.HasFuncSetOccludedBounds()
}

// FuncSetOccludedBounds returns the function "WEBEXT.virtualKeyboardPrivate.setOccludedBounds".
func FuncSetOccludedBounds() (fn js.Func[func(boundsList js.Array[Bounds])]) {
	bindings.FuncSetOccludedBounds(
		js.Pointer(&fn),
	)
	return
}

// SetOccludedBounds calls the function "WEBEXT.virtualKeyboardPrivate.setOccludedBounds" directly.
func SetOccludedBounds(boundsList js.Array[Bounds]) (ret js.Void) {
	bindings.CallSetOccludedBounds(
		js.Pointer(&ret),
		boundsList.Ref(),
	)

	return
}

// TrySetOccludedBounds calls the function "WEBEXT.virtualKeyboardPrivate.setOccludedBounds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetOccludedBounds(boundsList js.Array[Bounds]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetOccludedBounds(
		js.Pointer(&ret), js.Pointer(&exception),
		boundsList.Ref(),
	)

	return
}

// HasFuncSetWindowBoundsInScreen returns true if the function "WEBEXT.virtualKeyboardPrivate.setWindowBoundsInScreen" exists.
func HasFuncSetWindowBoundsInScreen() bool {
	return js.True == bindings.HasFuncSetWindowBoundsInScreen()
}

// FuncSetWindowBoundsInScreen returns the function "WEBEXT.virtualKeyboardPrivate.setWindowBoundsInScreen".
func FuncSetWindowBoundsInScreen() (fn js.Func[func(bounds Bounds)]) {
	bindings.FuncSetWindowBoundsInScreen(
		js.Pointer(&fn),
	)
	return
}

// SetWindowBoundsInScreen calls the function "WEBEXT.virtualKeyboardPrivate.setWindowBoundsInScreen" directly.
func SetWindowBoundsInScreen(bounds Bounds) (ret js.Void) {
	bindings.CallSetWindowBoundsInScreen(
		js.Pointer(&ret),
		js.Pointer(&bounds),
	)

	return
}

// TrySetWindowBoundsInScreen calls the function "WEBEXT.virtualKeyboardPrivate.setWindowBoundsInScreen"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetWindowBoundsInScreen(bounds Bounds) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetWindowBoundsInScreen(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&bounds),
	)

	return
}
