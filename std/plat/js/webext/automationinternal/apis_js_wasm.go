// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package automationinternal

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/automationinternal/bindings"
)

type AXEventParams struct {
	// TreeID is "AXEventParams.treeID"
	//
	// Optional
	TreeID js.String
	// TargetID is "AXEventParams.targetID"
	//
	// Optional
	//
	// NOTE: FFI_USE_TargetID MUST be set to true to make this field effective.
	TargetID int32
	// EventType is "AXEventParams.eventType"
	//
	// Optional
	EventType js.String
	// EventFrom is "AXEventParams.eventFrom"
	//
	// Optional
	EventFrom js.String
	// MouseX is "AXEventParams.mouseX"
	//
	// Optional
	//
	// NOTE: FFI_USE_MouseX MUST be set to true to make this field effective.
	MouseX float64
	// MouseY is "AXEventParams.mouseY"
	//
	// Optional
	//
	// NOTE: FFI_USE_MouseY MUST be set to true to make this field effective.
	MouseY float64
	// ActionRequestID is "AXEventParams.actionRequestID"
	//
	// Optional
	//
	// NOTE: FFI_USE_ActionRequestID MUST be set to true to make this field effective.
	ActionRequestID int32

	FFI_USE_TargetID        bool // for TargetID.
	FFI_USE_MouseX          bool // for MouseX.
	FFI_USE_MouseY          bool // for MouseY.
	FFI_USE_ActionRequestID bool // for ActionRequestID.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AXEventParams with all fields set.
func (p AXEventParams) FromRef(ref js.Ref) AXEventParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AXEventParams in the application heap.
func (p AXEventParams) New() js.Ref {
	return bindings.AXEventParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AXEventParams) UpdateFrom(ref js.Ref) {
	bindings.AXEventParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AXEventParams) Update(ref js.Ref) {
	bindings.AXEventParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AXEventParams) FreeMembers(recursive bool) {
	js.Free(
		p.TreeID.Ref(),
		p.EventType.Ref(),
		p.EventFrom.Ref(),
	)
	p.TreeID = p.TreeID.FromRef(js.Undefined)
	p.EventType = p.EventType.FromRef(js.Undefined)
	p.EventFrom = p.EventFrom.FromRef(js.Undefined)
}

type AXTextLocationParams struct {
	// TreeID is "AXTextLocationParams.treeID"
	//
	// Optional
	TreeID js.String
	// NodeID is "AXTextLocationParams.nodeID"
	//
	// Optional
	//
	// NOTE: FFI_USE_NodeID MUST be set to true to make this field effective.
	NodeID int32
	// Result is "AXTextLocationParams.result"
	//
	// Optional
	//
	// NOTE: FFI_USE_Result MUST be set to true to make this field effective.
	Result bool
	// Left is "AXTextLocationParams.left"
	//
	// Optional
	//
	// NOTE: FFI_USE_Left MUST be set to true to make this field effective.
	Left int32
	// Top is "AXTextLocationParams.top"
	//
	// Optional
	//
	// NOTE: FFI_USE_Top MUST be set to true to make this field effective.
	Top int32
	// Width is "AXTextLocationParams.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int32
	// Height is "AXTextLocationParams.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int32
	// RequestID is "AXTextLocationParams.requestID"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestID MUST be set to true to make this field effective.
	RequestID int32

	FFI_USE_NodeID    bool // for NodeID.
	FFI_USE_Result    bool // for Result.
	FFI_USE_Left      bool // for Left.
	FFI_USE_Top       bool // for Top.
	FFI_USE_Width     bool // for Width.
	FFI_USE_Height    bool // for Height.
	FFI_USE_RequestID bool // for RequestID.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AXTextLocationParams with all fields set.
func (p AXTextLocationParams) FromRef(ref js.Ref) AXTextLocationParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AXTextLocationParams in the application heap.
func (p AXTextLocationParams) New() js.Ref {
	return bindings.AXTextLocationParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AXTextLocationParams) UpdateFrom(ref js.Ref) {
	bindings.AXTextLocationParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AXTextLocationParams) Update(ref js.Ref) {
	bindings.AXTextLocationParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AXTextLocationParams) FreeMembers(recursive bool) {
	js.Free(
		p.TreeID.Ref(),
	)
	p.TreeID = p.TreeID.FromRef(js.Undefined)
}

type DisableDesktopCallbackFunc func(this js.Ref) js.Ref

func (fn DisableDesktopCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DisableDesktopCallbackFunc) DispatchCallback(
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

type DisableDesktopCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *DisableDesktopCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DisableDesktopCallback[T]) DispatchCallback(
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

type EnableDesktopCallbackFunc func(this js.Ref, tree_id js.String) js.Ref

func (fn EnableDesktopCallbackFunc) Register() js.Func[func(tree_id js.String)] {
	return js.RegisterCallback[func(tree_id js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EnableDesktopCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EnableDesktopCallback[T any] struct {
	Fn  func(arg T, this js.Ref, tree_id js.String) js.Ref
	Arg T
}

func (cb *EnableDesktopCallback[T]) Register() js.Func[func(tree_id js.String)] {
	return js.RegisterCallback[func(tree_id js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EnableDesktopCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetImageDataParams struct {
	// MaxWidth is "GetImageDataParams.maxWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxWidth MUST be set to true to make this field effective.
	MaxWidth int32
	// MaxHeight is "GetImageDataParams.maxHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxHeight MUST be set to true to make this field effective.
	MaxHeight int32

	FFI_USE_MaxWidth  bool // for MaxWidth.
	FFI_USE_MaxHeight bool // for MaxHeight.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetImageDataParams with all fields set.
func (p GetImageDataParams) FromRef(ref js.Ref) GetImageDataParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetImageDataParams in the application heap.
func (p GetImageDataParams) New() js.Ref {
	return bindings.GetImageDataParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetImageDataParams) UpdateFrom(ref js.Ref) {
	bindings.GetImageDataParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetImageDataParams) Update(ref js.Ref) {
	bindings.GetImageDataParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetImageDataParams) FreeMembers(recursive bool) {
}

type GetTextLocationDataParams struct {
	// StartIndex is "GetTextLocationDataParams.startIndex"
	//
	// Optional
	//
	// NOTE: FFI_USE_StartIndex MUST be set to true to make this field effective.
	StartIndex int32
	// EndIndex is "GetTextLocationDataParams.endIndex"
	//
	// Optional
	//
	// NOTE: FFI_USE_EndIndex MUST be set to true to make this field effective.
	EndIndex int32

	FFI_USE_StartIndex bool // for StartIndex.
	FFI_USE_EndIndex   bool // for EndIndex.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetTextLocationDataParams with all fields set.
func (p GetTextLocationDataParams) FromRef(ref js.Ref) GetTextLocationDataParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetTextLocationDataParams in the application heap.
func (p GetTextLocationDataParams) New() js.Ref {
	return bindings.GetTextLocationDataParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetTextLocationDataParams) UpdateFrom(ref js.Ref) {
	bindings.GetTextLocationDataParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetTextLocationDataParams) Update(ref js.Ref) {
	bindings.GetTextLocationDataParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetTextLocationDataParams) FreeMembers(recursive bool) {
}

type HitTestParams struct {
	// X is "HitTestParams.x"
	//
	// Optional
	//
	// NOTE: FFI_USE_X MUST be set to true to make this field effective.
	X int32
	// Y is "HitTestParams.y"
	//
	// Optional
	//
	// NOTE: FFI_USE_Y MUST be set to true to make this field effective.
	Y int32
	// EventToFire is "HitTestParams.eventToFire"
	//
	// Optional
	EventToFire js.String

	FFI_USE_X bool // for X.
	FFI_USE_Y bool // for Y.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HitTestParams with all fields set.
func (p HitTestParams) FromRef(ref js.Ref) HitTestParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HitTestParams in the application heap.
func (p HitTestParams) New() js.Ref {
	return bindings.HitTestParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HitTestParams) UpdateFrom(ref js.Ref) {
	bindings.HitTestParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HitTestParams) Update(ref js.Ref) {
	bindings.HitTestParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HitTestParams) FreeMembers(recursive bool) {
	js.Free(
		p.EventToFire.Ref(),
	)
	p.EventToFire = p.EventToFire.FromRef(js.Undefined)
}

type PerformActionRequiredParams struct {
	// TreeID is "PerformActionRequiredParams.treeID"
	//
	// Optional
	TreeID js.String
	// AutomationNodeID is "PerformActionRequiredParams.automationNodeID"
	//
	// Optional
	//
	// NOTE: FFI_USE_AutomationNodeID MUST be set to true to make this field effective.
	AutomationNodeID int32
	// ActionType is "PerformActionRequiredParams.actionType"
	//
	// Optional
	ActionType js.String
	// RequestID is "PerformActionRequiredParams.requestID"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestID MUST be set to true to make this field effective.
	RequestID int32

	FFI_USE_AutomationNodeID bool // for AutomationNodeID.
	FFI_USE_RequestID        bool // for RequestID.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PerformActionRequiredParams with all fields set.
func (p PerformActionRequiredParams) FromRef(ref js.Ref) PerformActionRequiredParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PerformActionRequiredParams in the application heap.
func (p PerformActionRequiredParams) New() js.Ref {
	return bindings.PerformActionRequiredParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PerformActionRequiredParams) UpdateFrom(ref js.Ref) {
	bindings.PerformActionRequiredParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PerformActionRequiredParams) Update(ref js.Ref) {
	bindings.PerformActionRequiredParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PerformActionRequiredParams) FreeMembers(recursive bool) {
	js.Free(
		p.TreeID.Ref(),
		p.ActionType.Ref(),
	)
	p.TreeID = p.TreeID.FromRef(js.Undefined)
	p.ActionType = p.ActionType.FromRef(js.Undefined)
}

type PerformCustomActionParams struct {
	// CustomActionID is "PerformCustomActionParams.customActionID"
	//
	// Optional
	//
	// NOTE: FFI_USE_CustomActionID MUST be set to true to make this field effective.
	CustomActionID int32

	FFI_USE_CustomActionID bool // for CustomActionID.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PerformCustomActionParams with all fields set.
func (p PerformCustomActionParams) FromRef(ref js.Ref) PerformCustomActionParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PerformCustomActionParams in the application heap.
func (p PerformCustomActionParams) New() js.Ref {
	return bindings.PerformCustomActionParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PerformCustomActionParams) UpdateFrom(ref js.Ref) {
	bindings.PerformCustomActionParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PerformCustomActionParams) Update(ref js.Ref) {
	bindings.PerformCustomActionParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PerformCustomActionParams) FreeMembers(recursive bool) {
}

type ReplaceSelectedTextParams struct {
	// Value is "ReplaceSelectedTextParams.value"
	//
	// Optional
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReplaceSelectedTextParams with all fields set.
func (p ReplaceSelectedTextParams) FromRef(ref js.Ref) ReplaceSelectedTextParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReplaceSelectedTextParams in the application heap.
func (p ReplaceSelectedTextParams) New() js.Ref {
	return bindings.ReplaceSelectedTextParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReplaceSelectedTextParams) UpdateFrom(ref js.Ref) {
	bindings.ReplaceSelectedTextParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReplaceSelectedTextParams) Update(ref js.Ref) {
	bindings.ReplaceSelectedTextParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReplaceSelectedTextParams) FreeMembers(recursive bool) {
	js.Free(
		p.Value.Ref(),
	)
	p.Value = p.Value.FromRef(js.Undefined)
}

type ScrollToPointParams struct {
	// X is "ScrollToPointParams.x"
	//
	// Optional
	//
	// NOTE: FFI_USE_X MUST be set to true to make this field effective.
	X int32
	// Y is "ScrollToPointParams.y"
	//
	// Optional
	//
	// NOTE: FFI_USE_Y MUST be set to true to make this field effective.
	Y int32

	FFI_USE_X bool // for X.
	FFI_USE_Y bool // for Y.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScrollToPointParams with all fields set.
func (p ScrollToPointParams) FromRef(ref js.Ref) ScrollToPointParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ScrollToPointParams in the application heap.
func (p ScrollToPointParams) New() js.Ref {
	return bindings.ScrollToPointParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ScrollToPointParams) UpdateFrom(ref js.Ref) {
	bindings.ScrollToPointParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ScrollToPointParams) Update(ref js.Ref) {
	bindings.ScrollToPointParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ScrollToPointParams) FreeMembers(recursive bool) {
}

type ScrollToPositionAtRowColumnParams struct {
	// Row is "ScrollToPositionAtRowColumnParams.row"
	//
	// Optional
	//
	// NOTE: FFI_USE_Row MUST be set to true to make this field effective.
	Row int32
	// Column is "ScrollToPositionAtRowColumnParams.column"
	//
	// Optional
	//
	// NOTE: FFI_USE_Column MUST be set to true to make this field effective.
	Column int32

	FFI_USE_Row    bool // for Row.
	FFI_USE_Column bool // for Column.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScrollToPositionAtRowColumnParams with all fields set.
func (p ScrollToPositionAtRowColumnParams) FromRef(ref js.Ref) ScrollToPositionAtRowColumnParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ScrollToPositionAtRowColumnParams in the application heap.
func (p ScrollToPositionAtRowColumnParams) New() js.Ref {
	return bindings.ScrollToPositionAtRowColumnParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ScrollToPositionAtRowColumnParams) UpdateFrom(ref js.Ref) {
	bindings.ScrollToPositionAtRowColumnParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ScrollToPositionAtRowColumnParams) Update(ref js.Ref) {
	bindings.ScrollToPositionAtRowColumnParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ScrollToPositionAtRowColumnParams) FreeMembers(recursive bool) {
}

type SetScrollOffsetParams struct {
	// X is "SetScrollOffsetParams.x"
	//
	// Optional
	//
	// NOTE: FFI_USE_X MUST be set to true to make this field effective.
	X int32
	// Y is "SetScrollOffsetParams.y"
	//
	// Optional
	//
	// NOTE: FFI_USE_Y MUST be set to true to make this field effective.
	Y int32

	FFI_USE_X bool // for X.
	FFI_USE_Y bool // for Y.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetScrollOffsetParams with all fields set.
func (p SetScrollOffsetParams) FromRef(ref js.Ref) SetScrollOffsetParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetScrollOffsetParams in the application heap.
func (p SetScrollOffsetParams) New() js.Ref {
	return bindings.SetScrollOffsetParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetScrollOffsetParams) UpdateFrom(ref js.Ref) {
	bindings.SetScrollOffsetParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetScrollOffsetParams) Update(ref js.Ref) {
	bindings.SetScrollOffsetParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetScrollOffsetParams) FreeMembers(recursive bool) {
}

type SetSelectionParams struct {
	// FocusNodeID is "SetSelectionParams.focusNodeID"
	//
	// Optional
	//
	// NOTE: FFI_USE_FocusNodeID MUST be set to true to make this field effective.
	FocusNodeID int32
	// AnchorOffset is "SetSelectionParams.anchorOffset"
	//
	// Optional
	//
	// NOTE: FFI_USE_AnchorOffset MUST be set to true to make this field effective.
	AnchorOffset int32
	// FocusOffset is "SetSelectionParams.focusOffset"
	//
	// Optional
	//
	// NOTE: FFI_USE_FocusOffset MUST be set to true to make this field effective.
	FocusOffset int32

	FFI_USE_FocusNodeID  bool // for FocusNodeID.
	FFI_USE_AnchorOffset bool // for AnchorOffset.
	FFI_USE_FocusOffset  bool // for FocusOffset.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetSelectionParams with all fields set.
func (p SetSelectionParams) FromRef(ref js.Ref) SetSelectionParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetSelectionParams in the application heap.
func (p SetSelectionParams) New() js.Ref {
	return bindings.SetSelectionParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetSelectionParams) UpdateFrom(ref js.Ref) {
	bindings.SetSelectionParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetSelectionParams) Update(ref js.Ref) {
	bindings.SetSelectionParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetSelectionParams) FreeMembers(recursive bool) {
}

type SetValueParams struct {
	// Value is "SetValueParams.value"
	//
	// Optional
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetValueParams with all fields set.
func (p SetValueParams) FromRef(ref js.Ref) SetValueParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetValueParams in the application heap.
func (p SetValueParams) New() js.Ref {
	return bindings.SetValueParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetValueParams) UpdateFrom(ref js.Ref) {
	bindings.SetValueParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetValueParams) Update(ref js.Ref) {
	bindings.SetValueParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetValueParams) FreeMembers(recursive bool) {
	js.Free(
		p.Value.Ref(),
	)
	p.Value = p.Value.FromRef(js.Undefined)
}

// HasFuncDisableDesktop returns true if the function "WEBEXT.automationInternal.disableDesktop" exists.
func HasFuncDisableDesktop() bool {
	return js.True == bindings.HasFuncDisableDesktop()
}

// FuncDisableDesktop returns the function "WEBEXT.automationInternal.disableDesktop".
func FuncDisableDesktop() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncDisableDesktop(
		js.Pointer(&fn),
	)
	return
}

// DisableDesktop calls the function "WEBEXT.automationInternal.disableDesktop" directly.
func DisableDesktop(callback js.Func[func()]) (ret js.Void) {
	bindings.CallDisableDesktop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryDisableDesktop calls the function "WEBEXT.automationInternal.disableDesktop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDisableDesktop(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDisableDesktop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncEnableDesktop returns true if the function "WEBEXT.automationInternal.enableDesktop" exists.
func HasFuncEnableDesktop() bool {
	return js.True == bindings.HasFuncEnableDesktop()
}

// FuncEnableDesktop returns the function "WEBEXT.automationInternal.enableDesktop".
func FuncEnableDesktop() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncEnableDesktop(
		js.Pointer(&fn),
	)
	return
}

// EnableDesktop calls the function "WEBEXT.automationInternal.enableDesktop" directly.
func EnableDesktop() (ret js.Promise[js.String]) {
	bindings.CallEnableDesktop(
		js.Pointer(&ret),
	)

	return
}

// TryEnableDesktop calls the function "WEBEXT.automationInternal.enableDesktop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEnableDesktop() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryEnableDesktop(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncEnableTree returns true if the function "WEBEXT.automationInternal.enableTree" exists.
func HasFuncEnableTree() bool {
	return js.True == bindings.HasFuncEnableTree()
}

// FuncEnableTree returns the function "WEBEXT.automationInternal.enableTree".
func FuncEnableTree() (fn js.Func[func(tree_id js.String)]) {
	bindings.FuncEnableTree(
		js.Pointer(&fn),
	)
	return
}

// EnableTree calls the function "WEBEXT.automationInternal.enableTree" directly.
func EnableTree(tree_id js.String) (ret js.Void) {
	bindings.CallEnableTree(
		js.Pointer(&ret),
		tree_id.Ref(),
	)

	return
}

// TryEnableTree calls the function "WEBEXT.automationInternal.enableTree"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEnableTree(tree_id js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEnableTree(
		js.Pointer(&ret), js.Pointer(&exception),
		tree_id.Ref(),
	)

	return
}

type OnAccessibilityEventEventCallbackFunc func(this js.Ref, update *AXEventParams) js.Ref

func (fn OnAccessibilityEventEventCallbackFunc) Register() js.Func[func(update *AXEventParams)] {
	return js.RegisterCallback[func(update *AXEventParams)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAccessibilityEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AXEventParams
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

type OnAccessibilityEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, update *AXEventParams) js.Ref
	Arg T
}

func (cb *OnAccessibilityEventEventCallback[T]) Register() js.Func[func(update *AXEventParams)] {
	return js.RegisterCallback[func(update *AXEventParams)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAccessibilityEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AXEventParams
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

// HasFuncOnAccessibilityEvent returns true if the function "WEBEXT.automationInternal.onAccessibilityEvent.addListener" exists.
func HasFuncOnAccessibilityEvent() bool {
	return js.True == bindings.HasFuncOnAccessibilityEvent()
}

// FuncOnAccessibilityEvent returns the function "WEBEXT.automationInternal.onAccessibilityEvent.addListener".
func FuncOnAccessibilityEvent() (fn js.Func[func(callback js.Func[func(update *AXEventParams)])]) {
	bindings.FuncOnAccessibilityEvent(
		js.Pointer(&fn),
	)
	return
}

// OnAccessibilityEvent calls the function "WEBEXT.automationInternal.onAccessibilityEvent.addListener" directly.
func OnAccessibilityEvent(callback js.Func[func(update *AXEventParams)]) (ret js.Void) {
	bindings.CallOnAccessibilityEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAccessibilityEvent calls the function "WEBEXT.automationInternal.onAccessibilityEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAccessibilityEvent(callback js.Func[func(update *AXEventParams)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAccessibilityEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAccessibilityEvent returns true if the function "WEBEXT.automationInternal.onAccessibilityEvent.removeListener" exists.
func HasFuncOffAccessibilityEvent() bool {
	return js.True == bindings.HasFuncOffAccessibilityEvent()
}

// FuncOffAccessibilityEvent returns the function "WEBEXT.automationInternal.onAccessibilityEvent.removeListener".
func FuncOffAccessibilityEvent() (fn js.Func[func(callback js.Func[func(update *AXEventParams)])]) {
	bindings.FuncOffAccessibilityEvent(
		js.Pointer(&fn),
	)
	return
}

// OffAccessibilityEvent calls the function "WEBEXT.automationInternal.onAccessibilityEvent.removeListener" directly.
func OffAccessibilityEvent(callback js.Func[func(update *AXEventParams)]) (ret js.Void) {
	bindings.CallOffAccessibilityEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAccessibilityEvent calls the function "WEBEXT.automationInternal.onAccessibilityEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAccessibilityEvent(callback js.Func[func(update *AXEventParams)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAccessibilityEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAccessibilityEvent returns true if the function "WEBEXT.automationInternal.onAccessibilityEvent.hasListener" exists.
func HasFuncHasOnAccessibilityEvent() bool {
	return js.True == bindings.HasFuncHasOnAccessibilityEvent()
}

// FuncHasOnAccessibilityEvent returns the function "WEBEXT.automationInternal.onAccessibilityEvent.hasListener".
func FuncHasOnAccessibilityEvent() (fn js.Func[func(callback js.Func[func(update *AXEventParams)]) bool]) {
	bindings.FuncHasOnAccessibilityEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnAccessibilityEvent calls the function "WEBEXT.automationInternal.onAccessibilityEvent.hasListener" directly.
func HasOnAccessibilityEvent(callback js.Func[func(update *AXEventParams)]) (ret bool) {
	bindings.CallHasOnAccessibilityEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAccessibilityEvent calls the function "WEBEXT.automationInternal.onAccessibilityEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAccessibilityEvent(callback js.Func[func(update *AXEventParams)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAccessibilityEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnAccessibilityTreeDestroyedEventCallbackFunc func(this js.Ref, treeID js.String) js.Ref

func (fn OnAccessibilityTreeDestroyedEventCallbackFunc) Register() js.Func[func(treeID js.String)] {
	return js.RegisterCallback[func(treeID js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAccessibilityTreeDestroyedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnAccessibilityTreeDestroyedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, treeID js.String) js.Ref
	Arg T
}

func (cb *OnAccessibilityTreeDestroyedEventCallback[T]) Register() js.Func[func(treeID js.String)] {
	return js.RegisterCallback[func(treeID js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAccessibilityTreeDestroyedEventCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnAccessibilityTreeDestroyed returns true if the function "WEBEXT.automationInternal.onAccessibilityTreeDestroyed.addListener" exists.
func HasFuncOnAccessibilityTreeDestroyed() bool {
	return js.True == bindings.HasFuncOnAccessibilityTreeDestroyed()
}

// FuncOnAccessibilityTreeDestroyed returns the function "WEBEXT.automationInternal.onAccessibilityTreeDestroyed.addListener".
func FuncOnAccessibilityTreeDestroyed() (fn js.Func[func(callback js.Func[func(treeID js.String)])]) {
	bindings.FuncOnAccessibilityTreeDestroyed(
		js.Pointer(&fn),
	)
	return
}

// OnAccessibilityTreeDestroyed calls the function "WEBEXT.automationInternal.onAccessibilityTreeDestroyed.addListener" directly.
func OnAccessibilityTreeDestroyed(callback js.Func[func(treeID js.String)]) (ret js.Void) {
	bindings.CallOnAccessibilityTreeDestroyed(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAccessibilityTreeDestroyed calls the function "WEBEXT.automationInternal.onAccessibilityTreeDestroyed.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAccessibilityTreeDestroyed(callback js.Func[func(treeID js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAccessibilityTreeDestroyed(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAccessibilityTreeDestroyed returns true if the function "WEBEXT.automationInternal.onAccessibilityTreeDestroyed.removeListener" exists.
func HasFuncOffAccessibilityTreeDestroyed() bool {
	return js.True == bindings.HasFuncOffAccessibilityTreeDestroyed()
}

// FuncOffAccessibilityTreeDestroyed returns the function "WEBEXT.automationInternal.onAccessibilityTreeDestroyed.removeListener".
func FuncOffAccessibilityTreeDestroyed() (fn js.Func[func(callback js.Func[func(treeID js.String)])]) {
	bindings.FuncOffAccessibilityTreeDestroyed(
		js.Pointer(&fn),
	)
	return
}

// OffAccessibilityTreeDestroyed calls the function "WEBEXT.automationInternal.onAccessibilityTreeDestroyed.removeListener" directly.
func OffAccessibilityTreeDestroyed(callback js.Func[func(treeID js.String)]) (ret js.Void) {
	bindings.CallOffAccessibilityTreeDestroyed(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAccessibilityTreeDestroyed calls the function "WEBEXT.automationInternal.onAccessibilityTreeDestroyed.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAccessibilityTreeDestroyed(callback js.Func[func(treeID js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAccessibilityTreeDestroyed(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAccessibilityTreeDestroyed returns true if the function "WEBEXT.automationInternal.onAccessibilityTreeDestroyed.hasListener" exists.
func HasFuncHasOnAccessibilityTreeDestroyed() bool {
	return js.True == bindings.HasFuncHasOnAccessibilityTreeDestroyed()
}

// FuncHasOnAccessibilityTreeDestroyed returns the function "WEBEXT.automationInternal.onAccessibilityTreeDestroyed.hasListener".
func FuncHasOnAccessibilityTreeDestroyed() (fn js.Func[func(callback js.Func[func(treeID js.String)]) bool]) {
	bindings.FuncHasOnAccessibilityTreeDestroyed(
		js.Pointer(&fn),
	)
	return
}

// HasOnAccessibilityTreeDestroyed calls the function "WEBEXT.automationInternal.onAccessibilityTreeDestroyed.hasListener" directly.
func HasOnAccessibilityTreeDestroyed(callback js.Func[func(treeID js.String)]) (ret bool) {
	bindings.CallHasOnAccessibilityTreeDestroyed(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAccessibilityTreeDestroyed calls the function "WEBEXT.automationInternal.onAccessibilityTreeDestroyed.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAccessibilityTreeDestroyed(callback js.Func[func(treeID js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAccessibilityTreeDestroyed(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnAccessibilityTreeSerializationErrorEventCallbackFunc func(this js.Ref, treeID js.String) js.Ref

func (fn OnAccessibilityTreeSerializationErrorEventCallbackFunc) Register() js.Func[func(treeID js.String)] {
	return js.RegisterCallback[func(treeID js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAccessibilityTreeSerializationErrorEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnAccessibilityTreeSerializationErrorEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, treeID js.String) js.Ref
	Arg T
}

func (cb *OnAccessibilityTreeSerializationErrorEventCallback[T]) Register() js.Func[func(treeID js.String)] {
	return js.RegisterCallback[func(treeID js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAccessibilityTreeSerializationErrorEventCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnAccessibilityTreeSerializationError returns true if the function "WEBEXT.automationInternal.onAccessibilityTreeSerializationError.addListener" exists.
func HasFuncOnAccessibilityTreeSerializationError() bool {
	return js.True == bindings.HasFuncOnAccessibilityTreeSerializationError()
}

// FuncOnAccessibilityTreeSerializationError returns the function "WEBEXT.automationInternal.onAccessibilityTreeSerializationError.addListener".
func FuncOnAccessibilityTreeSerializationError() (fn js.Func[func(callback js.Func[func(treeID js.String)])]) {
	bindings.FuncOnAccessibilityTreeSerializationError(
		js.Pointer(&fn),
	)
	return
}

// OnAccessibilityTreeSerializationError calls the function "WEBEXT.automationInternal.onAccessibilityTreeSerializationError.addListener" directly.
func OnAccessibilityTreeSerializationError(callback js.Func[func(treeID js.String)]) (ret js.Void) {
	bindings.CallOnAccessibilityTreeSerializationError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAccessibilityTreeSerializationError calls the function "WEBEXT.automationInternal.onAccessibilityTreeSerializationError.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAccessibilityTreeSerializationError(callback js.Func[func(treeID js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAccessibilityTreeSerializationError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAccessibilityTreeSerializationError returns true if the function "WEBEXT.automationInternal.onAccessibilityTreeSerializationError.removeListener" exists.
func HasFuncOffAccessibilityTreeSerializationError() bool {
	return js.True == bindings.HasFuncOffAccessibilityTreeSerializationError()
}

// FuncOffAccessibilityTreeSerializationError returns the function "WEBEXT.automationInternal.onAccessibilityTreeSerializationError.removeListener".
func FuncOffAccessibilityTreeSerializationError() (fn js.Func[func(callback js.Func[func(treeID js.String)])]) {
	bindings.FuncOffAccessibilityTreeSerializationError(
		js.Pointer(&fn),
	)
	return
}

// OffAccessibilityTreeSerializationError calls the function "WEBEXT.automationInternal.onAccessibilityTreeSerializationError.removeListener" directly.
func OffAccessibilityTreeSerializationError(callback js.Func[func(treeID js.String)]) (ret js.Void) {
	bindings.CallOffAccessibilityTreeSerializationError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAccessibilityTreeSerializationError calls the function "WEBEXT.automationInternal.onAccessibilityTreeSerializationError.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAccessibilityTreeSerializationError(callback js.Func[func(treeID js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAccessibilityTreeSerializationError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAccessibilityTreeSerializationError returns true if the function "WEBEXT.automationInternal.onAccessibilityTreeSerializationError.hasListener" exists.
func HasFuncHasOnAccessibilityTreeSerializationError() bool {
	return js.True == bindings.HasFuncHasOnAccessibilityTreeSerializationError()
}

// FuncHasOnAccessibilityTreeSerializationError returns the function "WEBEXT.automationInternal.onAccessibilityTreeSerializationError.hasListener".
func FuncHasOnAccessibilityTreeSerializationError() (fn js.Func[func(callback js.Func[func(treeID js.String)]) bool]) {
	bindings.FuncHasOnAccessibilityTreeSerializationError(
		js.Pointer(&fn),
	)
	return
}

// HasOnAccessibilityTreeSerializationError calls the function "WEBEXT.automationInternal.onAccessibilityTreeSerializationError.hasListener" directly.
func HasOnAccessibilityTreeSerializationError(callback js.Func[func(treeID js.String)]) (ret bool) {
	bindings.CallHasOnAccessibilityTreeSerializationError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAccessibilityTreeSerializationError calls the function "WEBEXT.automationInternal.onAccessibilityTreeSerializationError.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAccessibilityTreeSerializationError(callback js.Func[func(treeID js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAccessibilityTreeSerializationError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnActionResultEventCallbackFunc func(this js.Ref, treeID js.String, requestID int32, result bool) js.Ref

func (fn OnActionResultEventCallbackFunc) Register() js.Func[func(treeID js.String, requestID int32, result bool)] {
	return js.RegisterCallback[func(treeID js.String, requestID int32, result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnActionResultEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		args[2+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnActionResultEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, treeID js.String, requestID int32, result bool) js.Ref
	Arg T
}

func (cb *OnActionResultEventCallback[T]) Register() js.Func[func(treeID js.String, requestID int32, result bool)] {
	return js.RegisterCallback[func(treeID js.String, requestID int32, result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnActionResultEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		args[2+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnActionResult returns true if the function "WEBEXT.automationInternal.onActionResult.addListener" exists.
func HasFuncOnActionResult() bool {
	return js.True == bindings.HasFuncOnActionResult()
}

// FuncOnActionResult returns the function "WEBEXT.automationInternal.onActionResult.addListener".
func FuncOnActionResult() (fn js.Func[func(callback js.Func[func(treeID js.String, requestID int32, result bool)])]) {
	bindings.FuncOnActionResult(
		js.Pointer(&fn),
	)
	return
}

// OnActionResult calls the function "WEBEXT.automationInternal.onActionResult.addListener" directly.
func OnActionResult(callback js.Func[func(treeID js.String, requestID int32, result bool)]) (ret js.Void) {
	bindings.CallOnActionResult(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnActionResult calls the function "WEBEXT.automationInternal.onActionResult.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnActionResult(callback js.Func[func(treeID js.String, requestID int32, result bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnActionResult(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffActionResult returns true if the function "WEBEXT.automationInternal.onActionResult.removeListener" exists.
func HasFuncOffActionResult() bool {
	return js.True == bindings.HasFuncOffActionResult()
}

// FuncOffActionResult returns the function "WEBEXT.automationInternal.onActionResult.removeListener".
func FuncOffActionResult() (fn js.Func[func(callback js.Func[func(treeID js.String, requestID int32, result bool)])]) {
	bindings.FuncOffActionResult(
		js.Pointer(&fn),
	)
	return
}

// OffActionResult calls the function "WEBEXT.automationInternal.onActionResult.removeListener" directly.
func OffActionResult(callback js.Func[func(treeID js.String, requestID int32, result bool)]) (ret js.Void) {
	bindings.CallOffActionResult(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffActionResult calls the function "WEBEXT.automationInternal.onActionResult.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffActionResult(callback js.Func[func(treeID js.String, requestID int32, result bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffActionResult(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnActionResult returns true if the function "WEBEXT.automationInternal.onActionResult.hasListener" exists.
func HasFuncHasOnActionResult() bool {
	return js.True == bindings.HasFuncHasOnActionResult()
}

// FuncHasOnActionResult returns the function "WEBEXT.automationInternal.onActionResult.hasListener".
func FuncHasOnActionResult() (fn js.Func[func(callback js.Func[func(treeID js.String, requestID int32, result bool)]) bool]) {
	bindings.FuncHasOnActionResult(
		js.Pointer(&fn),
	)
	return
}

// HasOnActionResult calls the function "WEBEXT.automationInternal.onActionResult.hasListener" directly.
func HasOnActionResult(callback js.Func[func(treeID js.String, requestID int32, result bool)]) (ret bool) {
	bindings.CallHasOnActionResult(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnActionResult calls the function "WEBEXT.automationInternal.onActionResult.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnActionResult(callback js.Func[func(treeID js.String, requestID int32, result bool)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnActionResult(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnAllAutomationEventListenersRemovedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnAllAutomationEventListenersRemovedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAllAutomationEventListenersRemovedEventCallbackFunc) DispatchCallback(
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

type OnAllAutomationEventListenersRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnAllAutomationEventListenersRemovedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAllAutomationEventListenersRemovedEventCallback[T]) DispatchCallback(
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

// HasFuncOnAllAutomationEventListenersRemoved returns true if the function "WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.addListener" exists.
func HasFuncOnAllAutomationEventListenersRemoved() bool {
	return js.True == bindings.HasFuncOnAllAutomationEventListenersRemoved()
}

// FuncOnAllAutomationEventListenersRemoved returns the function "WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.addListener".
func FuncOnAllAutomationEventListenersRemoved() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnAllAutomationEventListenersRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnAllAutomationEventListenersRemoved calls the function "WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.addListener" directly.
func OnAllAutomationEventListenersRemoved(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnAllAutomationEventListenersRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAllAutomationEventListenersRemoved calls the function "WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAllAutomationEventListenersRemoved(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAllAutomationEventListenersRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAllAutomationEventListenersRemoved returns true if the function "WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.removeListener" exists.
func HasFuncOffAllAutomationEventListenersRemoved() bool {
	return js.True == bindings.HasFuncOffAllAutomationEventListenersRemoved()
}

// FuncOffAllAutomationEventListenersRemoved returns the function "WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.removeListener".
func FuncOffAllAutomationEventListenersRemoved() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffAllAutomationEventListenersRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffAllAutomationEventListenersRemoved calls the function "WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.removeListener" directly.
func OffAllAutomationEventListenersRemoved(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffAllAutomationEventListenersRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAllAutomationEventListenersRemoved calls the function "WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAllAutomationEventListenersRemoved(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAllAutomationEventListenersRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAllAutomationEventListenersRemoved returns true if the function "WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.hasListener" exists.
func HasFuncHasOnAllAutomationEventListenersRemoved() bool {
	return js.True == bindings.HasFuncHasOnAllAutomationEventListenersRemoved()
}

// FuncHasOnAllAutomationEventListenersRemoved returns the function "WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.hasListener".
func FuncHasOnAllAutomationEventListenersRemoved() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnAllAutomationEventListenersRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnAllAutomationEventListenersRemoved calls the function "WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.hasListener" directly.
func HasOnAllAutomationEventListenersRemoved(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnAllAutomationEventListenersRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAllAutomationEventListenersRemoved calls the function "WEBEXT.automationInternal.onAllAutomationEventListenersRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAllAutomationEventListenersRemoved(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAllAutomationEventListenersRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnChildTreeIDEventCallbackFunc func(this js.Ref, treeID js.String) js.Ref

func (fn OnChildTreeIDEventCallbackFunc) Register() js.Func[func(treeID js.String)] {
	return js.RegisterCallback[func(treeID js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnChildTreeIDEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnChildTreeIDEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, treeID js.String) js.Ref
	Arg T
}

func (cb *OnChildTreeIDEventCallback[T]) Register() js.Func[func(treeID js.String)] {
	return js.RegisterCallback[func(treeID js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnChildTreeIDEventCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnChildTreeID returns true if the function "WEBEXT.automationInternal.onChildTreeID.addListener" exists.
func HasFuncOnChildTreeID() bool {
	return js.True == bindings.HasFuncOnChildTreeID()
}

// FuncOnChildTreeID returns the function "WEBEXT.automationInternal.onChildTreeID.addListener".
func FuncOnChildTreeID() (fn js.Func[func(callback js.Func[func(treeID js.String)])]) {
	bindings.FuncOnChildTreeID(
		js.Pointer(&fn),
	)
	return
}

// OnChildTreeID calls the function "WEBEXT.automationInternal.onChildTreeID.addListener" directly.
func OnChildTreeID(callback js.Func[func(treeID js.String)]) (ret js.Void) {
	bindings.CallOnChildTreeID(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnChildTreeID calls the function "WEBEXT.automationInternal.onChildTreeID.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnChildTreeID(callback js.Func[func(treeID js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnChildTreeID(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffChildTreeID returns true if the function "WEBEXT.automationInternal.onChildTreeID.removeListener" exists.
func HasFuncOffChildTreeID() bool {
	return js.True == bindings.HasFuncOffChildTreeID()
}

// FuncOffChildTreeID returns the function "WEBEXT.automationInternal.onChildTreeID.removeListener".
func FuncOffChildTreeID() (fn js.Func[func(callback js.Func[func(treeID js.String)])]) {
	bindings.FuncOffChildTreeID(
		js.Pointer(&fn),
	)
	return
}

// OffChildTreeID calls the function "WEBEXT.automationInternal.onChildTreeID.removeListener" directly.
func OffChildTreeID(callback js.Func[func(treeID js.String)]) (ret js.Void) {
	bindings.CallOffChildTreeID(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffChildTreeID calls the function "WEBEXT.automationInternal.onChildTreeID.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffChildTreeID(callback js.Func[func(treeID js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffChildTreeID(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnChildTreeID returns true if the function "WEBEXT.automationInternal.onChildTreeID.hasListener" exists.
func HasFuncHasOnChildTreeID() bool {
	return js.True == bindings.HasFuncHasOnChildTreeID()
}

// FuncHasOnChildTreeID returns the function "WEBEXT.automationInternal.onChildTreeID.hasListener".
func FuncHasOnChildTreeID() (fn js.Func[func(callback js.Func[func(treeID js.String)]) bool]) {
	bindings.FuncHasOnChildTreeID(
		js.Pointer(&fn),
	)
	return
}

// HasOnChildTreeID calls the function "WEBEXT.automationInternal.onChildTreeID.hasListener" directly.
func HasOnChildTreeID(callback js.Func[func(treeID js.String)]) (ret bool) {
	bindings.CallHasOnChildTreeID(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnChildTreeID calls the function "WEBEXT.automationInternal.onChildTreeID.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnChildTreeID(callback js.Func[func(treeID js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnChildTreeID(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnGetTextLocationResultEventCallbackFunc func(this js.Ref, params *AXTextLocationParams) js.Ref

func (fn OnGetTextLocationResultEventCallbackFunc) Register() js.Func[func(params *AXTextLocationParams)] {
	return js.RegisterCallback[func(params *AXTextLocationParams)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnGetTextLocationResultEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AXTextLocationParams
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

type OnGetTextLocationResultEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, params *AXTextLocationParams) js.Ref
	Arg T
}

func (cb *OnGetTextLocationResultEventCallback[T]) Register() js.Func[func(params *AXTextLocationParams)] {
	return js.RegisterCallback[func(params *AXTextLocationParams)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnGetTextLocationResultEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AXTextLocationParams
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

// HasFuncOnGetTextLocationResult returns true if the function "WEBEXT.automationInternal.onGetTextLocationResult.addListener" exists.
func HasFuncOnGetTextLocationResult() bool {
	return js.True == bindings.HasFuncOnGetTextLocationResult()
}

// FuncOnGetTextLocationResult returns the function "WEBEXT.automationInternal.onGetTextLocationResult.addListener".
func FuncOnGetTextLocationResult() (fn js.Func[func(callback js.Func[func(params *AXTextLocationParams)])]) {
	bindings.FuncOnGetTextLocationResult(
		js.Pointer(&fn),
	)
	return
}

// OnGetTextLocationResult calls the function "WEBEXT.automationInternal.onGetTextLocationResult.addListener" directly.
func OnGetTextLocationResult(callback js.Func[func(params *AXTextLocationParams)]) (ret js.Void) {
	bindings.CallOnGetTextLocationResult(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnGetTextLocationResult calls the function "WEBEXT.automationInternal.onGetTextLocationResult.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnGetTextLocationResult(callback js.Func[func(params *AXTextLocationParams)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnGetTextLocationResult(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffGetTextLocationResult returns true if the function "WEBEXT.automationInternal.onGetTextLocationResult.removeListener" exists.
func HasFuncOffGetTextLocationResult() bool {
	return js.True == bindings.HasFuncOffGetTextLocationResult()
}

// FuncOffGetTextLocationResult returns the function "WEBEXT.automationInternal.onGetTextLocationResult.removeListener".
func FuncOffGetTextLocationResult() (fn js.Func[func(callback js.Func[func(params *AXTextLocationParams)])]) {
	bindings.FuncOffGetTextLocationResult(
		js.Pointer(&fn),
	)
	return
}

// OffGetTextLocationResult calls the function "WEBEXT.automationInternal.onGetTextLocationResult.removeListener" directly.
func OffGetTextLocationResult(callback js.Func[func(params *AXTextLocationParams)]) (ret js.Void) {
	bindings.CallOffGetTextLocationResult(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffGetTextLocationResult calls the function "WEBEXT.automationInternal.onGetTextLocationResult.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffGetTextLocationResult(callback js.Func[func(params *AXTextLocationParams)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffGetTextLocationResult(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnGetTextLocationResult returns true if the function "WEBEXT.automationInternal.onGetTextLocationResult.hasListener" exists.
func HasFuncHasOnGetTextLocationResult() bool {
	return js.True == bindings.HasFuncHasOnGetTextLocationResult()
}

// FuncHasOnGetTextLocationResult returns the function "WEBEXT.automationInternal.onGetTextLocationResult.hasListener".
func FuncHasOnGetTextLocationResult() (fn js.Func[func(callback js.Func[func(params *AXTextLocationParams)]) bool]) {
	bindings.FuncHasOnGetTextLocationResult(
		js.Pointer(&fn),
	)
	return
}

// HasOnGetTextLocationResult calls the function "WEBEXT.automationInternal.onGetTextLocationResult.hasListener" directly.
func HasOnGetTextLocationResult(callback js.Func[func(params *AXTextLocationParams)]) (ret bool) {
	bindings.CallHasOnGetTextLocationResult(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnGetTextLocationResult calls the function "WEBEXT.automationInternal.onGetTextLocationResult.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnGetTextLocationResult(callback js.Func[func(params *AXTextLocationParams)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnGetTextLocationResult(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnNodesRemovedEventCallbackFunc func(this js.Ref, treeID js.String, nodeIDs js.Array[int32]) js.Ref

func (fn OnNodesRemovedEventCallbackFunc) Register() js.Func[func(treeID js.String, nodeIDs js.Array[int32])] {
	return js.RegisterCallback[func(treeID js.String, nodeIDs js.Array[int32])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnNodesRemovedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.Array[int32]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnNodesRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, treeID js.String, nodeIDs js.Array[int32]) js.Ref
	Arg T
}

func (cb *OnNodesRemovedEventCallback[T]) Register() js.Func[func(treeID js.String, nodeIDs js.Array[int32])] {
	return js.RegisterCallback[func(treeID js.String, nodeIDs js.Array[int32])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnNodesRemovedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.Array[int32]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnNodesRemoved returns true if the function "WEBEXT.automationInternal.onNodesRemoved.addListener" exists.
func HasFuncOnNodesRemoved() bool {
	return js.True == bindings.HasFuncOnNodesRemoved()
}

// FuncOnNodesRemoved returns the function "WEBEXT.automationInternal.onNodesRemoved.addListener".
func FuncOnNodesRemoved() (fn js.Func[func(callback js.Func[func(treeID js.String, nodeIDs js.Array[int32])])]) {
	bindings.FuncOnNodesRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnNodesRemoved calls the function "WEBEXT.automationInternal.onNodesRemoved.addListener" directly.
func OnNodesRemoved(callback js.Func[func(treeID js.String, nodeIDs js.Array[int32])]) (ret js.Void) {
	bindings.CallOnNodesRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnNodesRemoved calls the function "WEBEXT.automationInternal.onNodesRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnNodesRemoved(callback js.Func[func(treeID js.String, nodeIDs js.Array[int32])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnNodesRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffNodesRemoved returns true if the function "WEBEXT.automationInternal.onNodesRemoved.removeListener" exists.
func HasFuncOffNodesRemoved() bool {
	return js.True == bindings.HasFuncOffNodesRemoved()
}

// FuncOffNodesRemoved returns the function "WEBEXT.automationInternal.onNodesRemoved.removeListener".
func FuncOffNodesRemoved() (fn js.Func[func(callback js.Func[func(treeID js.String, nodeIDs js.Array[int32])])]) {
	bindings.FuncOffNodesRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffNodesRemoved calls the function "WEBEXT.automationInternal.onNodesRemoved.removeListener" directly.
func OffNodesRemoved(callback js.Func[func(treeID js.String, nodeIDs js.Array[int32])]) (ret js.Void) {
	bindings.CallOffNodesRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffNodesRemoved calls the function "WEBEXT.automationInternal.onNodesRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffNodesRemoved(callback js.Func[func(treeID js.String, nodeIDs js.Array[int32])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffNodesRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnNodesRemoved returns true if the function "WEBEXT.automationInternal.onNodesRemoved.hasListener" exists.
func HasFuncHasOnNodesRemoved() bool {
	return js.True == bindings.HasFuncHasOnNodesRemoved()
}

// FuncHasOnNodesRemoved returns the function "WEBEXT.automationInternal.onNodesRemoved.hasListener".
func FuncHasOnNodesRemoved() (fn js.Func[func(callback js.Func[func(treeID js.String, nodeIDs js.Array[int32])]) bool]) {
	bindings.FuncHasOnNodesRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnNodesRemoved calls the function "WEBEXT.automationInternal.onNodesRemoved.hasListener" directly.
func HasOnNodesRemoved(callback js.Func[func(treeID js.String, nodeIDs js.Array[int32])]) (ret bool) {
	bindings.CallHasOnNodesRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnNodesRemoved calls the function "WEBEXT.automationInternal.onNodesRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnNodesRemoved(callback js.Func[func(treeID js.String, nodeIDs js.Array[int32])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnNodesRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnTreeChangeEventCallbackFunc func(this js.Ref, observerID int32, treeID js.String, nodeID int32, changeType js.String) js.Ref

func (fn OnTreeChangeEventCallbackFunc) Register() js.Func[func(observerID int32, treeID js.String, nodeID int32, changeType js.String)] {
	return js.RegisterCallback[func(observerID int32, treeID js.String, nodeID int32, changeType js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnTreeChangeEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.String{}.FromRef(args[1+1]),
		js.Number[int32]{}.FromRef(args[2+1]).Get(),
		js.String{}.FromRef(args[3+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnTreeChangeEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, observerID int32, treeID js.String, nodeID int32, changeType js.String) js.Ref
	Arg T
}

func (cb *OnTreeChangeEventCallback[T]) Register() js.Func[func(observerID int32, treeID js.String, nodeID int32, changeType js.String)] {
	return js.RegisterCallback[func(observerID int32, treeID js.String, nodeID int32, changeType js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnTreeChangeEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.String{}.FromRef(args[1+1]),
		js.Number[int32]{}.FromRef(args[2+1]).Get(),
		js.String{}.FromRef(args[3+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnTreeChange returns true if the function "WEBEXT.automationInternal.onTreeChange.addListener" exists.
func HasFuncOnTreeChange() bool {
	return js.True == bindings.HasFuncOnTreeChange()
}

// FuncOnTreeChange returns the function "WEBEXT.automationInternal.onTreeChange.addListener".
func FuncOnTreeChange() (fn js.Func[func(callback js.Func[func(observerID int32, treeID js.String, nodeID int32, changeType js.String)])]) {
	bindings.FuncOnTreeChange(
		js.Pointer(&fn),
	)
	return
}

// OnTreeChange calls the function "WEBEXT.automationInternal.onTreeChange.addListener" directly.
func OnTreeChange(callback js.Func[func(observerID int32, treeID js.String, nodeID int32, changeType js.String)]) (ret js.Void) {
	bindings.CallOnTreeChange(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnTreeChange calls the function "WEBEXT.automationInternal.onTreeChange.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTreeChange(callback js.Func[func(observerID int32, treeID js.String, nodeID int32, changeType js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTreeChange(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffTreeChange returns true if the function "WEBEXT.automationInternal.onTreeChange.removeListener" exists.
func HasFuncOffTreeChange() bool {
	return js.True == bindings.HasFuncOffTreeChange()
}

// FuncOffTreeChange returns the function "WEBEXT.automationInternal.onTreeChange.removeListener".
func FuncOffTreeChange() (fn js.Func[func(callback js.Func[func(observerID int32, treeID js.String, nodeID int32, changeType js.String)])]) {
	bindings.FuncOffTreeChange(
		js.Pointer(&fn),
	)
	return
}

// OffTreeChange calls the function "WEBEXT.automationInternal.onTreeChange.removeListener" directly.
func OffTreeChange(callback js.Func[func(observerID int32, treeID js.String, nodeID int32, changeType js.String)]) (ret js.Void) {
	bindings.CallOffTreeChange(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffTreeChange calls the function "WEBEXT.automationInternal.onTreeChange.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffTreeChange(callback js.Func[func(observerID int32, treeID js.String, nodeID int32, changeType js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffTreeChange(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnTreeChange returns true if the function "WEBEXT.automationInternal.onTreeChange.hasListener" exists.
func HasFuncHasOnTreeChange() bool {
	return js.True == bindings.HasFuncHasOnTreeChange()
}

// FuncHasOnTreeChange returns the function "WEBEXT.automationInternal.onTreeChange.hasListener".
func FuncHasOnTreeChange() (fn js.Func[func(callback js.Func[func(observerID int32, treeID js.String, nodeID int32, changeType js.String)]) bool]) {
	bindings.FuncHasOnTreeChange(
		js.Pointer(&fn),
	)
	return
}

// HasOnTreeChange calls the function "WEBEXT.automationInternal.onTreeChange.hasListener" directly.
func HasOnTreeChange(callback js.Func[func(observerID int32, treeID js.String, nodeID int32, changeType js.String)]) (ret bool) {
	bindings.CallHasOnTreeChange(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnTreeChange calls the function "WEBEXT.automationInternal.onTreeChange.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnTreeChange(callback js.Func[func(observerID int32, treeID js.String, nodeID int32, changeType js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnTreeChange(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncPerformAction returns true if the function "WEBEXT.automationInternal.performAction" exists.
func HasFuncPerformAction() bool {
	return js.True == bindings.HasFuncPerformAction()
}

// FuncPerformAction returns the function "WEBEXT.automationInternal.performAction".
func FuncPerformAction() (fn js.Func[func(args PerformActionRequiredParams, opt_args js.Object)]) {
	bindings.FuncPerformAction(
		js.Pointer(&fn),
	)
	return
}

// PerformAction calls the function "WEBEXT.automationInternal.performAction" directly.
func PerformAction(args PerformActionRequiredParams, opt_args js.Object) (ret js.Void) {
	bindings.CallPerformAction(
		js.Pointer(&ret),
		js.Pointer(&args),
		opt_args.Ref(),
	)

	return
}

// TryPerformAction calls the function "WEBEXT.automationInternal.performAction"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryPerformAction(args PerformActionRequiredParams, opt_args js.Object) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformAction(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&args),
		opt_args.Ref(),
	)

	return
}
