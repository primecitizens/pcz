// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package brailledisplayprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/brailledisplayprivate/bindings"
)

type DisplayState struct {
	// Available is "DisplayState.available"
	//
	// Optional
	//
	// NOTE: FFI_USE_Available MUST be set to true to make this field effective.
	Available bool
	// TextRowCount is "DisplayState.textRowCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_TextRowCount MUST be set to true to make this field effective.
	TextRowCount int32
	// TextColumnCount is "DisplayState.textColumnCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_TextColumnCount MUST be set to true to make this field effective.
	TextColumnCount int32
	// CellSize is "DisplayState.cellSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_CellSize MUST be set to true to make this field effective.
	CellSize int32

	FFI_USE_Available       bool // for Available.
	FFI_USE_TextRowCount    bool // for TextRowCount.
	FFI_USE_TextColumnCount bool // for TextColumnCount.
	FFI_USE_CellSize        bool // for CellSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DisplayState with all fields set.
func (p DisplayState) FromRef(ref js.Ref) DisplayState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DisplayState in the application heap.
func (p DisplayState) New() js.Ref {
	return bindings.DisplayStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DisplayState) UpdateFrom(ref js.Ref) {
	bindings.DisplayStateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DisplayState) Update(ref js.Ref) {
	bindings.DisplayStateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DisplayState) FreeMembers(recursive bool) {
}

type DisplayStateCallbackFunc func(this js.Ref, result *DisplayState) js.Ref

func (fn DisplayStateCallbackFunc) Register() js.Func[func(result *DisplayState)] {
	return js.RegisterCallback[func(result *DisplayState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DisplayStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DisplayState
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

type DisplayStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *DisplayState) js.Ref
	Arg T
}

func (cb *DisplayStateCallback[T]) Register() js.Func[func(result *DisplayState)] {
	return js.RegisterCallback[func(result *DisplayState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DisplayStateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DisplayState
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

type KeyCommand uint32

const (
	_ KeyCommand = iota

	KeyCommand_LINE_UP
	KeyCommand_LINE_DOWN
	KeyCommand_PAN_LEFT
	KeyCommand_PAN_RIGHT
	KeyCommand_TOP
	KeyCommand_BOTTOM
	KeyCommand_ROUTING
	KeyCommand_SECONDARY_ROUTING
	KeyCommand_DOTS
	KeyCommand_CHORD
	KeyCommand_STANDARD_KEY
)

func (KeyCommand) FromRef(str js.Ref) KeyCommand {
	return KeyCommand(bindings.ConstOfKeyCommand(str))
}

func (x KeyCommand) String() (string, bool) {
	switch x {
	case KeyCommand_LINE_UP:
		return "line_up", true
	case KeyCommand_LINE_DOWN:
		return "line_down", true
	case KeyCommand_PAN_LEFT:
		return "pan_left", true
	case KeyCommand_PAN_RIGHT:
		return "pan_right", true
	case KeyCommand_TOP:
		return "top", true
	case KeyCommand_BOTTOM:
		return "bottom", true
	case KeyCommand_ROUTING:
		return "routing", true
	case KeyCommand_SECONDARY_ROUTING:
		return "secondary_routing", true
	case KeyCommand_DOTS:
		return "dots", true
	case KeyCommand_CHORD:
		return "chord", true
	case KeyCommand_STANDARD_KEY:
		return "standard_key", true
	default:
		return "", false
	}
}

type KeyEvent struct {
	// Command is "KeyEvent.command"
	//
	// Optional
	Command KeyCommand
	// DisplayPosition is "KeyEvent.displayPosition"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayPosition MUST be set to true to make this field effective.
	DisplayPosition int32
	// BrailleDots is "KeyEvent.brailleDots"
	//
	// Optional
	//
	// NOTE: FFI_USE_BrailleDots MUST be set to true to make this field effective.
	BrailleDots int32
	// StandardKeyCode is "KeyEvent.standardKeyCode"
	//
	// Optional
	StandardKeyCode js.String
	// StandardKeyChar is "KeyEvent.standardKeyChar"
	//
	// Optional
	StandardKeyChar js.String
	// SpaceKey is "KeyEvent.spaceKey"
	//
	// Optional
	//
	// NOTE: FFI_USE_SpaceKey MUST be set to true to make this field effective.
	SpaceKey bool
	// AltKey is "KeyEvent.altKey"
	//
	// Optional
	//
	// NOTE: FFI_USE_AltKey MUST be set to true to make this field effective.
	AltKey bool
	// ShiftKey is "KeyEvent.shiftKey"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShiftKey MUST be set to true to make this field effective.
	ShiftKey bool
	// CtrlKey is "KeyEvent.ctrlKey"
	//
	// Optional
	//
	// NOTE: FFI_USE_CtrlKey MUST be set to true to make this field effective.
	CtrlKey bool

	FFI_USE_DisplayPosition bool // for DisplayPosition.
	FFI_USE_BrailleDots     bool // for BrailleDots.
	FFI_USE_SpaceKey        bool // for SpaceKey.
	FFI_USE_AltKey          bool // for AltKey.
	FFI_USE_ShiftKey        bool // for ShiftKey.
	FFI_USE_CtrlKey         bool // for CtrlKey.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a KeyEvent with all fields set.
func (p KeyEvent) FromRef(ref js.Ref) KeyEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new KeyEvent in the application heap.
func (p KeyEvent) New() js.Ref {
	return bindings.KeyEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *KeyEvent) UpdateFrom(ref js.Ref) {
	bindings.KeyEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *KeyEvent) Update(ref js.Ref) {
	bindings.KeyEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *KeyEvent) FreeMembers(recursive bool) {
	js.Free(
		p.StandardKeyCode.Ref(),
		p.StandardKeyChar.Ref(),
	)
	p.StandardKeyCode = p.StandardKeyCode.FromRef(js.Undefined)
	p.StandardKeyChar = p.StandardKeyChar.FromRef(js.Undefined)
}

// HasFuncGetDisplayState returns true if the function "WEBEXT.brailleDisplayPrivate.getDisplayState" exists.
func HasFuncGetDisplayState() bool {
	return js.True == bindings.HasFuncGetDisplayState()
}

// FuncGetDisplayState returns the function "WEBEXT.brailleDisplayPrivate.getDisplayState".
func FuncGetDisplayState() (fn js.Func[func() js.Promise[DisplayState]]) {
	bindings.FuncGetDisplayState(
		js.Pointer(&fn),
	)
	return
}

// GetDisplayState calls the function "WEBEXT.brailleDisplayPrivate.getDisplayState" directly.
func GetDisplayState() (ret js.Promise[DisplayState]) {
	bindings.CallGetDisplayState(
		js.Pointer(&ret),
	)

	return
}

// TryGetDisplayState calls the function "WEBEXT.brailleDisplayPrivate.getDisplayState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDisplayState() (ret js.Promise[DisplayState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDisplayState(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnDisplayStateChangedEventCallbackFunc func(this js.Ref, state *DisplayState) js.Ref

func (fn OnDisplayStateChangedEventCallbackFunc) Register() js.Func[func(state *DisplayState)] {
	return js.RegisterCallback[func(state *DisplayState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDisplayStateChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DisplayState
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

type OnDisplayStateChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, state *DisplayState) js.Ref
	Arg T
}

func (cb *OnDisplayStateChangedEventCallback[T]) Register() js.Func[func(state *DisplayState)] {
	return js.RegisterCallback[func(state *DisplayState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDisplayStateChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DisplayState
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

// HasFuncOnDisplayStateChanged returns true if the function "WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.addListener" exists.
func HasFuncOnDisplayStateChanged() bool {
	return js.True == bindings.HasFuncOnDisplayStateChanged()
}

// FuncOnDisplayStateChanged returns the function "WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.addListener".
func FuncOnDisplayStateChanged() (fn js.Func[func(callback js.Func[func(state *DisplayState)])]) {
	bindings.FuncOnDisplayStateChanged(
		js.Pointer(&fn),
	)
	return
}

// OnDisplayStateChanged calls the function "WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.addListener" directly.
func OnDisplayStateChanged(callback js.Func[func(state *DisplayState)]) (ret js.Void) {
	bindings.CallOnDisplayStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDisplayStateChanged calls the function "WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDisplayStateChanged(callback js.Func[func(state *DisplayState)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDisplayStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDisplayStateChanged returns true if the function "WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.removeListener" exists.
func HasFuncOffDisplayStateChanged() bool {
	return js.True == bindings.HasFuncOffDisplayStateChanged()
}

// FuncOffDisplayStateChanged returns the function "WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.removeListener".
func FuncOffDisplayStateChanged() (fn js.Func[func(callback js.Func[func(state *DisplayState)])]) {
	bindings.FuncOffDisplayStateChanged(
		js.Pointer(&fn),
	)
	return
}

// OffDisplayStateChanged calls the function "WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.removeListener" directly.
func OffDisplayStateChanged(callback js.Func[func(state *DisplayState)]) (ret js.Void) {
	bindings.CallOffDisplayStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDisplayStateChanged calls the function "WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDisplayStateChanged(callback js.Func[func(state *DisplayState)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDisplayStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDisplayStateChanged returns true if the function "WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.hasListener" exists.
func HasFuncHasOnDisplayStateChanged() bool {
	return js.True == bindings.HasFuncHasOnDisplayStateChanged()
}

// FuncHasOnDisplayStateChanged returns the function "WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.hasListener".
func FuncHasOnDisplayStateChanged() (fn js.Func[func(callback js.Func[func(state *DisplayState)]) bool]) {
	bindings.FuncHasOnDisplayStateChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnDisplayStateChanged calls the function "WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.hasListener" directly.
func HasOnDisplayStateChanged(callback js.Func[func(state *DisplayState)]) (ret bool) {
	bindings.CallHasOnDisplayStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDisplayStateChanged calls the function "WEBEXT.brailleDisplayPrivate.onDisplayStateChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDisplayStateChanged(callback js.Func[func(state *DisplayState)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDisplayStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnKeyEventEventCallbackFunc func(this js.Ref, event *KeyEvent) js.Ref

func (fn OnKeyEventEventCallbackFunc) Register() js.Func[func(event *KeyEvent)] {
	return js.RegisterCallback[func(event *KeyEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnKeyEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 KeyEvent
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

type OnKeyEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *KeyEvent) js.Ref
	Arg T
}

func (cb *OnKeyEventEventCallback[T]) Register() js.Func[func(event *KeyEvent)] {
	return js.RegisterCallback[func(event *KeyEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnKeyEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 KeyEvent
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

// HasFuncOnKeyEvent returns true if the function "WEBEXT.brailleDisplayPrivate.onKeyEvent.addListener" exists.
func HasFuncOnKeyEvent() bool {
	return js.True == bindings.HasFuncOnKeyEvent()
}

// FuncOnKeyEvent returns the function "WEBEXT.brailleDisplayPrivate.onKeyEvent.addListener".
func FuncOnKeyEvent() (fn js.Func[func(callback js.Func[func(event *KeyEvent)])]) {
	bindings.FuncOnKeyEvent(
		js.Pointer(&fn),
	)
	return
}

// OnKeyEvent calls the function "WEBEXT.brailleDisplayPrivate.onKeyEvent.addListener" directly.
func OnKeyEvent(callback js.Func[func(event *KeyEvent)]) (ret js.Void) {
	bindings.CallOnKeyEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnKeyEvent calls the function "WEBEXT.brailleDisplayPrivate.onKeyEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnKeyEvent(callback js.Func[func(event *KeyEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnKeyEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffKeyEvent returns true if the function "WEBEXT.brailleDisplayPrivate.onKeyEvent.removeListener" exists.
func HasFuncOffKeyEvent() bool {
	return js.True == bindings.HasFuncOffKeyEvent()
}

// FuncOffKeyEvent returns the function "WEBEXT.brailleDisplayPrivate.onKeyEvent.removeListener".
func FuncOffKeyEvent() (fn js.Func[func(callback js.Func[func(event *KeyEvent)])]) {
	bindings.FuncOffKeyEvent(
		js.Pointer(&fn),
	)
	return
}

// OffKeyEvent calls the function "WEBEXT.brailleDisplayPrivate.onKeyEvent.removeListener" directly.
func OffKeyEvent(callback js.Func[func(event *KeyEvent)]) (ret js.Void) {
	bindings.CallOffKeyEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffKeyEvent calls the function "WEBEXT.brailleDisplayPrivate.onKeyEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffKeyEvent(callback js.Func[func(event *KeyEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffKeyEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnKeyEvent returns true if the function "WEBEXT.brailleDisplayPrivate.onKeyEvent.hasListener" exists.
func HasFuncHasOnKeyEvent() bool {
	return js.True == bindings.HasFuncHasOnKeyEvent()
}

// FuncHasOnKeyEvent returns the function "WEBEXT.brailleDisplayPrivate.onKeyEvent.hasListener".
func FuncHasOnKeyEvent() (fn js.Func[func(callback js.Func[func(event *KeyEvent)]) bool]) {
	bindings.FuncHasOnKeyEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnKeyEvent calls the function "WEBEXT.brailleDisplayPrivate.onKeyEvent.hasListener" directly.
func HasOnKeyEvent(callback js.Func[func(event *KeyEvent)]) (ret bool) {
	bindings.CallHasOnKeyEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnKeyEvent calls the function "WEBEXT.brailleDisplayPrivate.onKeyEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnKeyEvent(callback js.Func[func(event *KeyEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnKeyEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncUpdateBluetoothBrailleDisplayAddress returns true if the function "WEBEXT.brailleDisplayPrivate.updateBluetoothBrailleDisplayAddress" exists.
func HasFuncUpdateBluetoothBrailleDisplayAddress() bool {
	return js.True == bindings.HasFuncUpdateBluetoothBrailleDisplayAddress()
}

// FuncUpdateBluetoothBrailleDisplayAddress returns the function "WEBEXT.brailleDisplayPrivate.updateBluetoothBrailleDisplayAddress".
func FuncUpdateBluetoothBrailleDisplayAddress() (fn js.Func[func(address js.String)]) {
	bindings.FuncUpdateBluetoothBrailleDisplayAddress(
		js.Pointer(&fn),
	)
	return
}

// UpdateBluetoothBrailleDisplayAddress calls the function "WEBEXT.brailleDisplayPrivate.updateBluetoothBrailleDisplayAddress" directly.
func UpdateBluetoothBrailleDisplayAddress(address js.String) (ret js.Void) {
	bindings.CallUpdateBluetoothBrailleDisplayAddress(
		js.Pointer(&ret),
		address.Ref(),
	)

	return
}

// TryUpdateBluetoothBrailleDisplayAddress calls the function "WEBEXT.brailleDisplayPrivate.updateBluetoothBrailleDisplayAddress"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateBluetoothBrailleDisplayAddress(address js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateBluetoothBrailleDisplayAddress(
		js.Pointer(&ret), js.Pointer(&exception),
		address.Ref(),
	)

	return
}

// HasFuncWriteDots returns true if the function "WEBEXT.brailleDisplayPrivate.writeDots" exists.
func HasFuncWriteDots() bool {
	return js.True == bindings.HasFuncWriteDots()
}

// FuncWriteDots returns the function "WEBEXT.brailleDisplayPrivate.writeDots".
func FuncWriteDots() (fn js.Func[func(cells js.ArrayBuffer, columns int32, rows int32)]) {
	bindings.FuncWriteDots(
		js.Pointer(&fn),
	)
	return
}

// WriteDots calls the function "WEBEXT.brailleDisplayPrivate.writeDots" directly.
func WriteDots(cells js.ArrayBuffer, columns int32, rows int32) (ret js.Void) {
	bindings.CallWriteDots(
		js.Pointer(&ret),
		cells.Ref(),
		int32(columns),
		int32(rows),
	)

	return
}

// TryWriteDots calls the function "WEBEXT.brailleDisplayPrivate.writeDots"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryWriteDots(cells js.ArrayBuffer, columns int32, rows int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWriteDots(
		js.Pointer(&ret), js.Pointer(&exception),
		cells.Ref(),
		int32(columns),
		int32(rows),
	)

	return
}
