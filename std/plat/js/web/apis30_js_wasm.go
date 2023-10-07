// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type IsInputPendingOptions struct {
	// IncludeContinuous is "IsInputPendingOptions.includeContinuous"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IncludeContinuous MUST be set to true to make this field effective.
	IncludeContinuous bool

	FFI_USE_IncludeContinuous bool // for IncludeContinuous.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IsInputPendingOptions with all fields set.
func (p IsInputPendingOptions) FromRef(ref js.Ref) IsInputPendingOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IsInputPendingOptions in the application heap.
func (p IsInputPendingOptions) New() js.Ref {
	return bindings.IsInputPendingOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IsInputPendingOptions) UpdateFrom(ref js.Ref) {
	bindings.IsInputPendingOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IsInputPendingOptions) Update(ref js.Ref) {
	bindings.IsInputPendingOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IsInputPendingOptions) FreeMembers(recursive bool) {
}

type Scheduling struct {
	ref js.Ref
}

func (this Scheduling) Once() Scheduling {
	this.ref.Once()
	return this
}

func (this Scheduling) Ref() js.Ref {
	return this.ref
}

func (this Scheduling) FromRef(ref js.Ref) Scheduling {
	this.ref = ref
	return this
}

func (this Scheduling) Free() {
	this.ref.Free()
}

// HasFuncIsInputPending returns true if the method "Scheduling.isInputPending" exists.
func (this Scheduling) HasFuncIsInputPending() bool {
	return js.True == bindings.HasFuncSchedulingIsInputPending(
		this.ref,
	)
}

// FuncIsInputPending returns the method "Scheduling.isInputPending".
func (this Scheduling) FuncIsInputPending() (fn js.Func[func(isInputPendingOptions IsInputPendingOptions) bool]) {
	bindings.FuncSchedulingIsInputPending(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsInputPending calls the method "Scheduling.isInputPending".
func (this Scheduling) IsInputPending(isInputPendingOptions IsInputPendingOptions) (ret bool) {
	bindings.CallSchedulingIsInputPending(
		this.ref, js.Pointer(&ret),
		js.Pointer(&isInputPendingOptions),
	)

	return
}

// TryIsInputPending calls the method "Scheduling.isInputPending"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Scheduling) TryIsInputPending(isInputPendingOptions IsInputPendingOptions) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySchedulingIsInputPending(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&isInputPendingOptions),
	)

	return
}

// HasFuncIsInputPending1 returns true if the method "Scheduling.isInputPending" exists.
func (this Scheduling) HasFuncIsInputPending1() bool {
	return js.True == bindings.HasFuncSchedulingIsInputPending1(
		this.ref,
	)
}

// FuncIsInputPending1 returns the method "Scheduling.isInputPending".
func (this Scheduling) FuncIsInputPending1() (fn js.Func[func() bool]) {
	bindings.FuncSchedulingIsInputPending1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsInputPending1 calls the method "Scheduling.isInputPending".
func (this Scheduling) IsInputPending1() (ret bool) {
	bindings.CallSchedulingIsInputPending1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsInputPending1 calls the method "Scheduling.isInputPending"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Scheduling) TryIsInputPending1() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySchedulingIsInputPending1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type WakeLockType uint32

const (
	_ WakeLockType = iota

	WakeLockType_SCREEN
)

func (WakeLockType) FromRef(str js.Ref) WakeLockType {
	return WakeLockType(bindings.ConstOfWakeLockType(str))
}

func (x WakeLockType) String() (string, bool) {
	switch x {
	case WakeLockType_SCREEN:
		return "screen", true
	default:
		return "", false
	}
}

type WakeLockSentinel struct {
	EventTarget
}

func (this WakeLockSentinel) Once() WakeLockSentinel {
	this.ref.Once()
	return this
}

func (this WakeLockSentinel) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this WakeLockSentinel) FromRef(ref js.Ref) WakeLockSentinel {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this WakeLockSentinel) Free() {
	this.ref.Free()
}

// Released returns the value of property "WakeLockSentinel.released".
//
// It returns ok=false if there is no such property.
func (this WakeLockSentinel) Released() (ret bool, ok bool) {
	ok = js.True == bindings.GetWakeLockSentinelReleased(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "WakeLockSentinel.type".
//
// It returns ok=false if there is no such property.
func (this WakeLockSentinel) Type() (ret WakeLockType, ok bool) {
	ok = js.True == bindings.GetWakeLockSentinelType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncRelease returns true if the method "WakeLockSentinel.release" exists.
func (this WakeLockSentinel) HasFuncRelease() bool {
	return js.True == bindings.HasFuncWakeLockSentinelRelease(
		this.ref,
	)
}

// FuncRelease returns the method "WakeLockSentinel.release".
func (this WakeLockSentinel) FuncRelease() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncWakeLockSentinelRelease(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Release calls the method "WakeLockSentinel.release".
func (this WakeLockSentinel) Release() (ret js.Promise[js.Void]) {
	bindings.CallWakeLockSentinelRelease(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRelease calls the method "WakeLockSentinel.release"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WakeLockSentinel) TryRelease() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWakeLockSentinelRelease(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type WakeLock struct {
	ref js.Ref
}

func (this WakeLock) Once() WakeLock {
	this.ref.Once()
	return this
}

func (this WakeLock) Ref() js.Ref {
	return this.ref
}

func (this WakeLock) FromRef(ref js.Ref) WakeLock {
	this.ref = ref
	return this
}

func (this WakeLock) Free() {
	this.ref.Free()
}

// HasFuncRequest returns true if the method "WakeLock.request" exists.
func (this WakeLock) HasFuncRequest() bool {
	return js.True == bindings.HasFuncWakeLockRequest(
		this.ref,
	)
}

// FuncRequest returns the method "WakeLock.request".
func (this WakeLock) FuncRequest() (fn js.Func[func(typ WakeLockType) js.Promise[WakeLockSentinel]]) {
	bindings.FuncWakeLockRequest(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Request calls the method "WakeLock.request".
func (this WakeLock) Request(typ WakeLockType) (ret js.Promise[WakeLockSentinel]) {
	bindings.CallWakeLockRequest(
		this.ref, js.Pointer(&ret),
		uint32(typ),
	)

	return
}

// TryRequest calls the method "WakeLock.request"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WakeLock) TryRequest(typ WakeLockType) (ret js.Promise[WakeLockSentinel], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWakeLockRequest(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
	)

	return
}

// HasFuncRequest1 returns true if the method "WakeLock.request" exists.
func (this WakeLock) HasFuncRequest1() bool {
	return js.True == bindings.HasFuncWakeLockRequest1(
		this.ref,
	)
}

// FuncRequest1 returns the method "WakeLock.request".
func (this WakeLock) FuncRequest1() (fn js.Func[func() js.Promise[WakeLockSentinel]]) {
	bindings.FuncWakeLockRequest1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Request1 calls the method "WakeLock.request".
func (this WakeLock) Request1() (ret js.Promise[WakeLockSentinel]) {
	bindings.CallWakeLockRequest1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequest1 calls the method "WakeLock.request"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WakeLock) TryRequest1() (ret js.Promise[WakeLockSentinel], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWakeLockRequest1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PointerEventInit struct {
	// PointerId is "PointerEventInit.pointerId"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_PointerId MUST be set to true to make this field effective.
	PointerId int32
	// Width is "PointerEventInit.width"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width float64
	// Height is "PointerEventInit.height"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height float64
	// Pressure is "PointerEventInit.pressure"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Pressure MUST be set to true to make this field effective.
	Pressure float32
	// TangentialPressure is "PointerEventInit.tangentialPressure"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_TangentialPressure MUST be set to true to make this field effective.
	TangentialPressure float32
	// TiltX is "PointerEventInit.tiltX"
	//
	// Optional
	//
	// NOTE: FFI_USE_TiltX MUST be set to true to make this field effective.
	TiltX int32
	// TiltY is "PointerEventInit.tiltY"
	//
	// Optional
	//
	// NOTE: FFI_USE_TiltY MUST be set to true to make this field effective.
	TiltY int32
	// Twist is "PointerEventInit.twist"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Twist MUST be set to true to make this field effective.
	Twist int32
	// AltitudeAngle is "PointerEventInit.altitudeAngle"
	//
	// Optional
	//
	// NOTE: FFI_USE_AltitudeAngle MUST be set to true to make this field effective.
	AltitudeAngle float64
	// AzimuthAngle is "PointerEventInit.azimuthAngle"
	//
	// Optional
	//
	// NOTE: FFI_USE_AzimuthAngle MUST be set to true to make this field effective.
	AzimuthAngle float64
	// PointerType is "PointerEventInit.pointerType"
	//
	// Optional, defaults to "".
	PointerType js.String
	// IsPrimary is "PointerEventInit.isPrimary"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IsPrimary MUST be set to true to make this field effective.
	IsPrimary bool
	// CoalescedEvents is "PointerEventInit.coalescedEvents"
	//
	// Optional, defaults to [].
	CoalescedEvents js.Array[PointerEvent]
	// PredictedEvents is "PointerEventInit.predictedEvents"
	//
	// Optional, defaults to [].
	PredictedEvents js.Array[PointerEvent]
	// MovementX is "PointerEventInit.movementX"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_MovementX MUST be set to true to make this field effective.
	MovementX float64
	// MovementY is "PointerEventInit.movementY"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_MovementY MUST be set to true to make this field effective.
	MovementY float64

	FFI_USE_PointerId          bool // for PointerId.
	FFI_USE_Width              bool // for Width.
	FFI_USE_Height             bool // for Height.
	FFI_USE_Pressure           bool // for Pressure.
	FFI_USE_TangentialPressure bool // for TangentialPressure.
	FFI_USE_TiltX              bool // for TiltX.
	FFI_USE_TiltY              bool // for TiltY.
	FFI_USE_Twist              bool // for Twist.
	FFI_USE_AltitudeAngle      bool // for AltitudeAngle.
	FFI_USE_AzimuthAngle       bool // for AzimuthAngle.
	FFI_USE_IsPrimary          bool // for IsPrimary.
	FFI_USE_MovementX          bool // for MovementX.
	FFI_USE_MovementY          bool // for MovementY.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PointerEventInit with all fields set.
func (p PointerEventInit) FromRef(ref js.Ref) PointerEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PointerEventInit in the application heap.
func (p PointerEventInit) New() js.Ref {
	return bindings.PointerEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PointerEventInit) UpdateFrom(ref js.Ref) {
	bindings.PointerEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PointerEventInit) Update(ref js.Ref) {
	bindings.PointerEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PointerEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.PointerType.Ref(),
		p.CoalescedEvents.Ref(),
		p.PredictedEvents.Ref(),
	)
	p.PointerType = p.PointerType.FromRef(js.Undefined)
	p.CoalescedEvents = p.CoalescedEvents.FromRef(js.Undefined)
	p.PredictedEvents = p.PredictedEvents.FromRef(js.Undefined)
}

func NewPointerEvent(typ js.String, eventInitDict PointerEventInit) (ret PointerEvent) {
	ret.ref = bindings.NewPointerEventByPointerEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewPointerEventByPointerEvent1(typ js.String) (ret PointerEvent) {
	ret.ref = bindings.NewPointerEventByPointerEvent1(
		typ.Ref())
	return
}

type PointerEvent struct {
	MouseEvent
}

func (this PointerEvent) Once() PointerEvent {
	this.ref.Once()
	return this
}

func (this PointerEvent) Ref() js.Ref {
	return this.MouseEvent.Ref()
}

func (this PointerEvent) FromRef(ref js.Ref) PointerEvent {
	this.MouseEvent = this.MouseEvent.FromRef(ref)
	return this
}

func (this PointerEvent) Free() {
	this.ref.Free()
}

// PointerId returns the value of property "PointerEvent.pointerId".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) PointerId() (ret int32, ok bool) {
	ok = js.True == bindings.GetPointerEventPointerId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "PointerEvent.width".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) Width() (ret float64, ok bool) {
	ok = js.True == bindings.GetPointerEventWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "PointerEvent.height".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) Height() (ret float64, ok bool) {
	ok = js.True == bindings.GetPointerEventHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Pressure returns the value of property "PointerEvent.pressure".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) Pressure() (ret float32, ok bool) {
	ok = js.True == bindings.GetPointerEventPressure(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TangentialPressure returns the value of property "PointerEvent.tangentialPressure".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) TangentialPressure() (ret float32, ok bool) {
	ok = js.True == bindings.GetPointerEventTangentialPressure(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TiltX returns the value of property "PointerEvent.tiltX".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) TiltX() (ret int32, ok bool) {
	ok = js.True == bindings.GetPointerEventTiltX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TiltY returns the value of property "PointerEvent.tiltY".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) TiltY() (ret int32, ok bool) {
	ok = js.True == bindings.GetPointerEventTiltY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Twist returns the value of property "PointerEvent.twist".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) Twist() (ret int32, ok bool) {
	ok = js.True == bindings.GetPointerEventTwist(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AltitudeAngle returns the value of property "PointerEvent.altitudeAngle".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) AltitudeAngle() (ret float64, ok bool) {
	ok = js.True == bindings.GetPointerEventAltitudeAngle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AzimuthAngle returns the value of property "PointerEvent.azimuthAngle".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) AzimuthAngle() (ret float64, ok bool) {
	ok = js.True == bindings.GetPointerEventAzimuthAngle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PointerType returns the value of property "PointerEvent.pointerType".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) PointerType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPointerEventPointerType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsPrimary returns the value of property "PointerEvent.isPrimary".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) IsPrimary() (ret bool, ok bool) {
	ok = js.True == bindings.GetPointerEventIsPrimary(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetCoalescedEvents returns true if the method "PointerEvent.getCoalescedEvents" exists.
func (this PointerEvent) HasFuncGetCoalescedEvents() bool {
	return js.True == bindings.HasFuncPointerEventGetCoalescedEvents(
		this.ref,
	)
}

// FuncGetCoalescedEvents returns the method "PointerEvent.getCoalescedEvents".
func (this PointerEvent) FuncGetCoalescedEvents() (fn js.Func[func() js.Array[PointerEvent]]) {
	bindings.FuncPointerEventGetCoalescedEvents(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCoalescedEvents calls the method "PointerEvent.getCoalescedEvents".
func (this PointerEvent) GetCoalescedEvents() (ret js.Array[PointerEvent]) {
	bindings.CallPointerEventGetCoalescedEvents(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetCoalescedEvents calls the method "PointerEvent.getCoalescedEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PointerEvent) TryGetCoalescedEvents() (ret js.Array[PointerEvent], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPointerEventGetCoalescedEvents(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetPredictedEvents returns true if the method "PointerEvent.getPredictedEvents" exists.
func (this PointerEvent) HasFuncGetPredictedEvents() bool {
	return js.True == bindings.HasFuncPointerEventGetPredictedEvents(
		this.ref,
	)
}

// FuncGetPredictedEvents returns the method "PointerEvent.getPredictedEvents".
func (this PointerEvent) FuncGetPredictedEvents() (fn js.Func[func() js.Array[PointerEvent]]) {
	bindings.FuncPointerEventGetPredictedEvents(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetPredictedEvents calls the method "PointerEvent.getPredictedEvents".
func (this PointerEvent) GetPredictedEvents() (ret js.Array[PointerEvent]) {
	bindings.CallPointerEventGetPredictedEvents(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetPredictedEvents calls the method "PointerEvent.getPredictedEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PointerEvent) TryGetPredictedEvents() (ret js.Array[PointerEvent], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPointerEventGetPredictedEvents(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type InkTrailStyle struct {
	// Color is "InkTrailStyle.color"
	//
	// Required
	Color js.String
	// Diameter is "InkTrailStyle.diameter"
	//
	// Required
	Diameter float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InkTrailStyle with all fields set.
func (p InkTrailStyle) FromRef(ref js.Ref) InkTrailStyle {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InkTrailStyle in the application heap.
func (p InkTrailStyle) New() js.Ref {
	return bindings.InkTrailStyleJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InkTrailStyle) UpdateFrom(ref js.Ref) {
	bindings.InkTrailStyleJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InkTrailStyle) Update(ref js.Ref) {
	bindings.InkTrailStyleJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InkTrailStyle) FreeMembers(recursive bool) {
	js.Free(
		p.Color.Ref(),
	)
	p.Color = p.Color.FromRef(js.Undefined)
}

type InkPresenter struct {
	ref js.Ref
}

func (this InkPresenter) Once() InkPresenter {
	this.ref.Once()
	return this
}

func (this InkPresenter) Ref() js.Ref {
	return this.ref
}

func (this InkPresenter) FromRef(ref js.Ref) InkPresenter {
	this.ref = ref
	return this
}

func (this InkPresenter) Free() {
	this.ref.Free()
}

// PresentationArea returns the value of property "InkPresenter.presentationArea".
//
// It returns ok=false if there is no such property.
func (this InkPresenter) PresentationArea() (ret Element, ok bool) {
	ok = js.True == bindings.GetInkPresenterPresentationArea(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ExpectedImprovement returns the value of property "InkPresenter.expectedImprovement".
//
// It returns ok=false if there is no such property.
func (this InkPresenter) ExpectedImprovement() (ret uint32, ok bool) {
	ok = js.True == bindings.GetInkPresenterExpectedImprovement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncUpdateInkTrailStartPoint returns true if the method "InkPresenter.updateInkTrailStartPoint" exists.
func (this InkPresenter) HasFuncUpdateInkTrailStartPoint() bool {
	return js.True == bindings.HasFuncInkPresenterUpdateInkTrailStartPoint(
		this.ref,
	)
}

// FuncUpdateInkTrailStartPoint returns the method "InkPresenter.updateInkTrailStartPoint".
func (this InkPresenter) FuncUpdateInkTrailStartPoint() (fn js.Func[func(event PointerEvent, style InkTrailStyle)]) {
	bindings.FuncInkPresenterUpdateInkTrailStartPoint(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdateInkTrailStartPoint calls the method "InkPresenter.updateInkTrailStartPoint".
func (this InkPresenter) UpdateInkTrailStartPoint(event PointerEvent, style InkTrailStyle) (ret js.Void) {
	bindings.CallInkPresenterUpdateInkTrailStartPoint(
		this.ref, js.Pointer(&ret),
		event.Ref(),
		js.Pointer(&style),
	)

	return
}

// TryUpdateInkTrailStartPoint calls the method "InkPresenter.updateInkTrailStartPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this InkPresenter) TryUpdateInkTrailStartPoint(event PointerEvent, style InkTrailStyle) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInkPresenterUpdateInkTrailStartPoint(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		event.Ref(),
		js.Pointer(&style),
	)

	return
}

type InkPresenterParam struct {
	// PresentationArea is "InkPresenterParam.presentationArea"
	//
	// Optional, defaults to null.
	PresentationArea Element

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InkPresenterParam with all fields set.
func (p InkPresenterParam) FromRef(ref js.Ref) InkPresenterParam {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InkPresenterParam in the application heap.
func (p InkPresenterParam) New() js.Ref {
	return bindings.InkPresenterParamJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InkPresenterParam) UpdateFrom(ref js.Ref) {
	bindings.InkPresenterParamJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InkPresenterParam) Update(ref js.Ref) {
	bindings.InkPresenterParamJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InkPresenterParam) FreeMembers(recursive bool) {
	js.Free(
		p.PresentationArea.Ref(),
	)
	p.PresentationArea = p.PresentationArea.FromRef(js.Undefined)
}

type Ink struct {
	ref js.Ref
}

func (this Ink) Once() Ink {
	this.ref.Once()
	return this
}

func (this Ink) Ref() js.Ref {
	return this.ref
}

func (this Ink) FromRef(ref js.Ref) Ink {
	this.ref = ref
	return this
}

func (this Ink) Free() {
	this.ref.Free()
}

// HasFuncRequestPresenter returns true if the method "Ink.requestPresenter" exists.
func (this Ink) HasFuncRequestPresenter() bool {
	return js.True == bindings.HasFuncInkRequestPresenter(
		this.ref,
	)
}

// FuncRequestPresenter returns the method "Ink.requestPresenter".
func (this Ink) FuncRequestPresenter() (fn js.Func[func(param InkPresenterParam) js.Promise[InkPresenter]]) {
	bindings.FuncInkRequestPresenter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestPresenter calls the method "Ink.requestPresenter".
func (this Ink) RequestPresenter(param InkPresenterParam) (ret js.Promise[InkPresenter]) {
	bindings.CallInkRequestPresenter(
		this.ref, js.Pointer(&ret),
		js.Pointer(&param),
	)

	return
}

// TryRequestPresenter calls the method "Ink.requestPresenter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Ink) TryRequestPresenter(param InkPresenterParam) (ret js.Promise[InkPresenter], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInkRequestPresenter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&param),
	)

	return
}

// HasFuncRequestPresenter1 returns true if the method "Ink.requestPresenter" exists.
func (this Ink) HasFuncRequestPresenter1() bool {
	return js.True == bindings.HasFuncInkRequestPresenter1(
		this.ref,
	)
}

// FuncRequestPresenter1 returns the method "Ink.requestPresenter".
func (this Ink) FuncRequestPresenter1() (fn js.Func[func() js.Promise[InkPresenter]]) {
	bindings.FuncInkRequestPresenter1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestPresenter1 calls the method "Ink.requestPresenter".
func (this Ink) RequestPresenter1() (ret js.Promise[InkPresenter]) {
	bindings.CallInkRequestPresenter1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestPresenter1 calls the method "Ink.requestPresenter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Ink) TryRequestPresenter1() (ret js.Promise[InkPresenter], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInkRequestPresenter1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PresentationConnectionState uint32

const (
	_ PresentationConnectionState = iota

	PresentationConnectionState_CONNECTING
	PresentationConnectionState_CONNECTED
	PresentationConnectionState_CLOSED
	PresentationConnectionState_TERMINATED
)

func (PresentationConnectionState) FromRef(str js.Ref) PresentationConnectionState {
	return PresentationConnectionState(bindings.ConstOfPresentationConnectionState(str))
}

func (x PresentationConnectionState) String() (string, bool) {
	switch x {
	case PresentationConnectionState_CONNECTING:
		return "connecting", true
	case PresentationConnectionState_CONNECTED:
		return "connected", true
	case PresentationConnectionState_CLOSED:
		return "closed", true
	case PresentationConnectionState_TERMINATED:
		return "terminated", true
	default:
		return "", false
	}
}

type PresentationConnection struct {
	EventTarget
}

func (this PresentationConnection) Once() PresentationConnection {
	this.ref.Once()
	return this
}

func (this PresentationConnection) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this PresentationConnection) FromRef(ref js.Ref) PresentationConnection {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this PresentationConnection) Free() {
	this.ref.Free()
}

// Id returns the value of property "PresentationConnection.id".
//
// It returns ok=false if there is no such property.
func (this PresentationConnection) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPresentationConnectionId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Url returns the value of property "PresentationConnection.url".
//
// It returns ok=false if there is no such property.
func (this PresentationConnection) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPresentationConnectionUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// State returns the value of property "PresentationConnection.state".
//
// It returns ok=false if there is no such property.
func (this PresentationConnection) State() (ret PresentationConnectionState, ok bool) {
	ok = js.True == bindings.GetPresentationConnectionState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BinaryType returns the value of property "PresentationConnection.binaryType".
//
// It returns ok=false if there is no such property.
func (this PresentationConnection) BinaryType() (ret BinaryType, ok bool) {
	ok = js.True == bindings.GetPresentationConnectionBinaryType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBinaryType sets the value of property "PresentationConnection.binaryType" to val.
//
// It returns false if the property cannot be set.
func (this PresentationConnection) SetBinaryType(val BinaryType) bool {
	return js.True == bindings.SetPresentationConnectionBinaryType(
		this.ref,
		uint32(val),
	)
}

// HasFuncClose returns true if the method "PresentationConnection.close" exists.
func (this PresentationConnection) HasFuncClose() bool {
	return js.True == bindings.HasFuncPresentationConnectionClose(
		this.ref,
	)
}

// FuncClose returns the method "PresentationConnection.close".
func (this PresentationConnection) FuncClose() (fn js.Func[func()]) {
	bindings.FuncPresentationConnectionClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "PresentationConnection.close".
func (this PresentationConnection) Close() (ret js.Void) {
	bindings.CallPresentationConnectionClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "PresentationConnection.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationConnection) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationConnectionClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncTerminate returns true if the method "PresentationConnection.terminate" exists.
func (this PresentationConnection) HasFuncTerminate() bool {
	return js.True == bindings.HasFuncPresentationConnectionTerminate(
		this.ref,
	)
}

// FuncTerminate returns the method "PresentationConnection.terminate".
func (this PresentationConnection) FuncTerminate() (fn js.Func[func()]) {
	bindings.FuncPresentationConnectionTerminate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Terminate calls the method "PresentationConnection.terminate".
func (this PresentationConnection) Terminate() (ret js.Void) {
	bindings.CallPresentationConnectionTerminate(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTerminate calls the method "PresentationConnection.terminate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationConnection) TryTerminate() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationConnectionTerminate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSend returns true if the method "PresentationConnection.send" exists.
func (this PresentationConnection) HasFuncSend() bool {
	return js.True == bindings.HasFuncPresentationConnectionSend(
		this.ref,
	)
}

// FuncSend returns the method "PresentationConnection.send".
func (this PresentationConnection) FuncSend() (fn js.Func[func(message js.String)]) {
	bindings.FuncPresentationConnectionSend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Send calls the method "PresentationConnection.send".
func (this PresentationConnection) Send(message js.String) (ret js.Void) {
	bindings.CallPresentationConnectionSend(
		this.ref, js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TrySend calls the method "PresentationConnection.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationConnection) TrySend(message js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationConnectionSend(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasFuncSend1 returns true if the method "PresentationConnection.send" exists.
func (this PresentationConnection) HasFuncSend1() bool {
	return js.True == bindings.HasFuncPresentationConnectionSend1(
		this.ref,
	)
}

// FuncSend1 returns the method "PresentationConnection.send".
func (this PresentationConnection) FuncSend1() (fn js.Func[func(data Blob)]) {
	bindings.FuncPresentationConnectionSend1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Send1 calls the method "PresentationConnection.send".
func (this PresentationConnection) Send1(data Blob) (ret js.Void) {
	bindings.CallPresentationConnectionSend1(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySend1 calls the method "PresentationConnection.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationConnection) TrySend1(data Blob) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationConnectionSend1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncSend2 returns true if the method "PresentationConnection.send" exists.
func (this PresentationConnection) HasFuncSend2() bool {
	return js.True == bindings.HasFuncPresentationConnectionSend2(
		this.ref,
	)
}

// FuncSend2 returns the method "PresentationConnection.send".
func (this PresentationConnection) FuncSend2() (fn js.Func[func(data js.ArrayBuffer)]) {
	bindings.FuncPresentationConnectionSend2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Send2 calls the method "PresentationConnection.send".
func (this PresentationConnection) Send2(data js.ArrayBuffer) (ret js.Void) {
	bindings.CallPresentationConnectionSend2(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySend2 calls the method "PresentationConnection.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationConnection) TrySend2(data js.ArrayBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationConnectionSend2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncSend3 returns true if the method "PresentationConnection.send" exists.
func (this PresentationConnection) HasFuncSend3() bool {
	return js.True == bindings.HasFuncPresentationConnectionSend3(
		this.ref,
	)
}

// FuncSend3 returns the method "PresentationConnection.send".
func (this PresentationConnection) FuncSend3() (fn js.Func[func(data js.ArrayBufferView)]) {
	bindings.FuncPresentationConnectionSend3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Send3 calls the method "PresentationConnection.send".
func (this PresentationConnection) Send3(data js.ArrayBufferView) (ret js.Void) {
	bindings.CallPresentationConnectionSend3(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySend3 calls the method "PresentationConnection.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationConnection) TrySend3(data js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationConnectionSend3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

type PresentationAvailability struct {
	EventTarget
}

func (this PresentationAvailability) Once() PresentationAvailability {
	this.ref.Once()
	return this
}

func (this PresentationAvailability) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this PresentationAvailability) FromRef(ref js.Ref) PresentationAvailability {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this PresentationAvailability) Free() {
	this.ref.Free()
}

// Value returns the value of property "PresentationAvailability.value".
//
// It returns ok=false if there is no such property.
func (this PresentationAvailability) Value() (ret bool, ok bool) {
	ok = js.True == bindings.GetPresentationAvailabilityValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

func NewPresentationRequest(url js.String) (ret PresentationRequest) {
	ret.ref = bindings.NewPresentationRequestByPresentationRequest(
		url.Ref())
	return
}

func NewPresentationRequestByPresentationRequest1(urls js.Array[js.String]) (ret PresentationRequest) {
	ret.ref = bindings.NewPresentationRequestByPresentationRequest1(
		urls.Ref())
	return
}

type PresentationRequest struct {
	EventTarget
}

func (this PresentationRequest) Once() PresentationRequest {
	this.ref.Once()
	return this
}

func (this PresentationRequest) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this PresentationRequest) FromRef(ref js.Ref) PresentationRequest {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this PresentationRequest) Free() {
	this.ref.Free()
}

// HasFuncStart returns true if the method "PresentationRequest.start" exists.
func (this PresentationRequest) HasFuncStart() bool {
	return js.True == bindings.HasFuncPresentationRequestStart(
		this.ref,
	)
}

// FuncStart returns the method "PresentationRequest.start".
func (this PresentationRequest) FuncStart() (fn js.Func[func() js.Promise[PresentationConnection]]) {
	bindings.FuncPresentationRequestStart(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start calls the method "PresentationRequest.start".
func (this PresentationRequest) Start() (ret js.Promise[PresentationConnection]) {
	bindings.CallPresentationRequestStart(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStart calls the method "PresentationRequest.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationRequest) TryStart() (ret js.Promise[PresentationConnection], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationRequestStart(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReconnect returns true if the method "PresentationRequest.reconnect" exists.
func (this PresentationRequest) HasFuncReconnect() bool {
	return js.True == bindings.HasFuncPresentationRequestReconnect(
		this.ref,
	)
}

// FuncReconnect returns the method "PresentationRequest.reconnect".
func (this PresentationRequest) FuncReconnect() (fn js.Func[func(presentationId js.String) js.Promise[PresentationConnection]]) {
	bindings.FuncPresentationRequestReconnect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Reconnect calls the method "PresentationRequest.reconnect".
func (this PresentationRequest) Reconnect(presentationId js.String) (ret js.Promise[PresentationConnection]) {
	bindings.CallPresentationRequestReconnect(
		this.ref, js.Pointer(&ret),
		presentationId.Ref(),
	)

	return
}

// TryReconnect calls the method "PresentationRequest.reconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationRequest) TryReconnect(presentationId js.String) (ret js.Promise[PresentationConnection], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationRequestReconnect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		presentationId.Ref(),
	)

	return
}

// HasFuncGetAvailability returns true if the method "PresentationRequest.getAvailability" exists.
func (this PresentationRequest) HasFuncGetAvailability() bool {
	return js.True == bindings.HasFuncPresentationRequestGetAvailability(
		this.ref,
	)
}

// FuncGetAvailability returns the method "PresentationRequest.getAvailability".
func (this PresentationRequest) FuncGetAvailability() (fn js.Func[func() js.Promise[PresentationAvailability]]) {
	bindings.FuncPresentationRequestGetAvailability(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAvailability calls the method "PresentationRequest.getAvailability".
func (this PresentationRequest) GetAvailability() (ret js.Promise[PresentationAvailability]) {
	bindings.CallPresentationRequestGetAvailability(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetAvailability calls the method "PresentationRequest.getAvailability"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationRequest) TryGetAvailability() (ret js.Promise[PresentationAvailability], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationRequestGetAvailability(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PresentationConnectionList struct {
	EventTarget
}

func (this PresentationConnectionList) Once() PresentationConnectionList {
	this.ref.Once()
	return this
}

func (this PresentationConnectionList) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this PresentationConnectionList) FromRef(ref js.Ref) PresentationConnectionList {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this PresentationConnectionList) Free() {
	this.ref.Free()
}

// Connections returns the value of property "PresentationConnectionList.connections".
//
// It returns ok=false if there is no such property.
func (this PresentationConnectionList) Connections() (ret js.FrozenArray[PresentationConnection], ok bool) {
	ok = js.True == bindings.GetPresentationConnectionListConnections(
		this.ref, js.Pointer(&ret),
	)
	return
}

type PresentationReceiver struct {
	ref js.Ref
}

func (this PresentationReceiver) Once() PresentationReceiver {
	this.ref.Once()
	return this
}

func (this PresentationReceiver) Ref() js.Ref {
	return this.ref
}

func (this PresentationReceiver) FromRef(ref js.Ref) PresentationReceiver {
	this.ref = ref
	return this
}

func (this PresentationReceiver) Free() {
	this.ref.Free()
}

// ConnectionList returns the value of property "PresentationReceiver.connectionList".
//
// It returns ok=false if there is no such property.
func (this PresentationReceiver) ConnectionList() (ret js.Promise[PresentationConnectionList], ok bool) {
	ok = js.True == bindings.GetPresentationReceiverConnectionList(
		this.ref, js.Pointer(&ret),
	)
	return
}

type Presentation struct {
	ref js.Ref
}

func (this Presentation) Once() Presentation {
	this.ref.Once()
	return this
}

func (this Presentation) Ref() js.Ref {
	return this.ref
}

func (this Presentation) FromRef(ref js.Ref) Presentation {
	this.ref = ref
	return this
}

func (this Presentation) Free() {
	this.ref.Free()
}

// DefaultRequest returns the value of property "Presentation.defaultRequest".
//
// It returns ok=false if there is no such property.
func (this Presentation) DefaultRequest() (ret PresentationRequest, ok bool) {
	ok = js.True == bindings.GetPresentationDefaultRequest(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDefaultRequest sets the value of property "Presentation.defaultRequest" to val.
//
// It returns false if the property cannot be set.
func (this Presentation) SetDefaultRequest(val PresentationRequest) bool {
	return js.True == bindings.SetPresentationDefaultRequest(
		this.ref,
		val.Ref(),
	)
}

// Receiver returns the value of property "Presentation.receiver".
//
// It returns ok=false if there is no such property.
func (this Presentation) Receiver() (ret PresentationReceiver, ok bool) {
	ok = js.True == bindings.GetPresentationReceiver(
		this.ref, js.Pointer(&ret),
	)
	return
}

type VirtualKeyboard struct {
	EventTarget
}

func (this VirtualKeyboard) Once() VirtualKeyboard {
	this.ref.Once()
	return this
}

func (this VirtualKeyboard) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this VirtualKeyboard) FromRef(ref js.Ref) VirtualKeyboard {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this VirtualKeyboard) Free() {
	this.ref.Free()
}

// BoundingRect returns the value of property "VirtualKeyboard.boundingRect".
//
// It returns ok=false if there is no such property.
func (this VirtualKeyboard) BoundingRect() (ret DOMRect, ok bool) {
	ok = js.True == bindings.GetVirtualKeyboardBoundingRect(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OverlaysContent returns the value of property "VirtualKeyboard.overlaysContent".
//
// It returns ok=false if there is no such property.
func (this VirtualKeyboard) OverlaysContent() (ret bool, ok bool) {
	ok = js.True == bindings.GetVirtualKeyboardOverlaysContent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetOverlaysContent sets the value of property "VirtualKeyboard.overlaysContent" to val.
//
// It returns false if the property cannot be set.
func (this VirtualKeyboard) SetOverlaysContent(val bool) bool {
	return js.True == bindings.SetVirtualKeyboardOverlaysContent(
		this.ref,
		js.Bool(bool(val)),
	)
}

// HasFuncShow returns true if the method "VirtualKeyboard.show" exists.
func (this VirtualKeyboard) HasFuncShow() bool {
	return js.True == bindings.HasFuncVirtualKeyboardShow(
		this.ref,
	)
}

// FuncShow returns the method "VirtualKeyboard.show".
func (this VirtualKeyboard) FuncShow() (fn js.Func[func()]) {
	bindings.FuncVirtualKeyboardShow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Show calls the method "VirtualKeyboard.show".
func (this VirtualKeyboard) Show() (ret js.Void) {
	bindings.CallVirtualKeyboardShow(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryShow calls the method "VirtualKeyboard.show"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VirtualKeyboard) TryShow() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVirtualKeyboardShow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncHide returns true if the method "VirtualKeyboard.hide" exists.
func (this VirtualKeyboard) HasFuncHide() bool {
	return js.True == bindings.HasFuncVirtualKeyboardHide(
		this.ref,
	)
}

// FuncHide returns the method "VirtualKeyboard.hide".
func (this VirtualKeyboard) FuncHide() (fn js.Func[func()]) {
	bindings.FuncVirtualKeyboardHide(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Hide calls the method "VirtualKeyboard.hide".
func (this VirtualKeyboard) Hide() (ret js.Void) {
	bindings.CallVirtualKeyboardHide(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryHide calls the method "VirtualKeyboard.hide"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VirtualKeyboard) TryHide() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVirtualKeyboardHide(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type NavigatorUABrandVersion struct {
	// Brand is "NavigatorUABrandVersion.brand"
	//
	// Optional
	Brand js.String
	// Version is "NavigatorUABrandVersion.version"
	//
	// Optional
	Version js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NavigatorUABrandVersion with all fields set.
func (p NavigatorUABrandVersion) FromRef(ref js.Ref) NavigatorUABrandVersion {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NavigatorUABrandVersion in the application heap.
func (p NavigatorUABrandVersion) New() js.Ref {
	return bindings.NavigatorUABrandVersionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NavigatorUABrandVersion) UpdateFrom(ref js.Ref) {
	bindings.NavigatorUABrandVersionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NavigatorUABrandVersion) Update(ref js.Ref) {
	bindings.NavigatorUABrandVersionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NavigatorUABrandVersion) FreeMembers(recursive bool) {
	js.Free(
		p.Brand.Ref(),
		p.Version.Ref(),
	)
	p.Brand = p.Brand.FromRef(js.Undefined)
	p.Version = p.Version.FromRef(js.Undefined)
}

type UADataValues struct {
	// Architecture is "UADataValues.architecture"
	//
	// Optional
	Architecture js.String
	// Bitness is "UADataValues.bitness"
	//
	// Optional
	Bitness js.String
	// Brands is "UADataValues.brands"
	//
	// Optional
	Brands js.Array[NavigatorUABrandVersion]
	// FormFactor is "UADataValues.formFactor"
	//
	// Optional
	FormFactor js.String
	// FullVersionList is "UADataValues.fullVersionList"
	//
	// Optional
	FullVersionList js.Array[NavigatorUABrandVersion]
	// Model is "UADataValues.model"
	//
	// Optional
	Model js.String
	// Mobile is "UADataValues.mobile"
	//
	// Optional
	//
	// NOTE: FFI_USE_Mobile MUST be set to true to make this field effective.
	Mobile bool
	// Platform is "UADataValues.platform"
	//
	// Optional
	Platform js.String
	// PlatformVersion is "UADataValues.platformVersion"
	//
	// Optional
	PlatformVersion js.String
	// UaFullVersion is "UADataValues.uaFullVersion"
	//
	// Optional
	UaFullVersion js.String
	// Wow64 is "UADataValues.wow64"
	//
	// Optional
	//
	// NOTE: FFI_USE_Wow64 MUST be set to true to make this field effective.
	Wow64 bool

	FFI_USE_Mobile bool // for Mobile.
	FFI_USE_Wow64  bool // for Wow64.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UADataValues with all fields set.
func (p UADataValues) FromRef(ref js.Ref) UADataValues {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UADataValues in the application heap.
func (p UADataValues) New() js.Ref {
	return bindings.UADataValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UADataValues) UpdateFrom(ref js.Ref) {
	bindings.UADataValuesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UADataValues) Update(ref js.Ref) {
	bindings.UADataValuesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UADataValues) FreeMembers(recursive bool) {
	js.Free(
		p.Architecture.Ref(),
		p.Bitness.Ref(),
		p.Brands.Ref(),
		p.FormFactor.Ref(),
		p.FullVersionList.Ref(),
		p.Model.Ref(),
		p.Platform.Ref(),
		p.PlatformVersion.Ref(),
		p.UaFullVersion.Ref(),
	)
	p.Architecture = p.Architecture.FromRef(js.Undefined)
	p.Bitness = p.Bitness.FromRef(js.Undefined)
	p.Brands = p.Brands.FromRef(js.Undefined)
	p.FormFactor = p.FormFactor.FromRef(js.Undefined)
	p.FullVersionList = p.FullVersionList.FromRef(js.Undefined)
	p.Model = p.Model.FromRef(js.Undefined)
	p.Platform = p.Platform.FromRef(js.Undefined)
	p.PlatformVersion = p.PlatformVersion.FromRef(js.Undefined)
	p.UaFullVersion = p.UaFullVersion.FromRef(js.Undefined)
}

type UALowEntropyJSON struct {
	// Brands is "UALowEntropyJSON.brands"
	//
	// Optional
	Brands js.Array[NavigatorUABrandVersion]
	// Mobile is "UALowEntropyJSON.mobile"
	//
	// Optional
	//
	// NOTE: FFI_USE_Mobile MUST be set to true to make this field effective.
	Mobile bool
	// Platform is "UALowEntropyJSON.platform"
	//
	// Optional
	Platform js.String

	FFI_USE_Mobile bool // for Mobile.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UALowEntropyJSON with all fields set.
func (p UALowEntropyJSON) FromRef(ref js.Ref) UALowEntropyJSON {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UALowEntropyJSON in the application heap.
func (p UALowEntropyJSON) New() js.Ref {
	return bindings.UALowEntropyJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UALowEntropyJSON) UpdateFrom(ref js.Ref) {
	bindings.UALowEntropyJSONJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UALowEntropyJSON) Update(ref js.Ref) {
	bindings.UALowEntropyJSONJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UALowEntropyJSON) FreeMembers(recursive bool) {
	js.Free(
		p.Brands.Ref(),
		p.Platform.Ref(),
	)
	p.Brands = p.Brands.FromRef(js.Undefined)
	p.Platform = p.Platform.FromRef(js.Undefined)
}

type NavigatorUAData struct {
	ref js.Ref
}

func (this NavigatorUAData) Once() NavigatorUAData {
	this.ref.Once()
	return this
}

func (this NavigatorUAData) Ref() js.Ref {
	return this.ref
}

func (this NavigatorUAData) FromRef(ref js.Ref) NavigatorUAData {
	this.ref = ref
	return this
}

func (this NavigatorUAData) Free() {
	this.ref.Free()
}

// Brands returns the value of property "NavigatorUAData.brands".
//
// It returns ok=false if there is no such property.
func (this NavigatorUAData) Brands() (ret js.FrozenArray[NavigatorUABrandVersion], ok bool) {
	ok = js.True == bindings.GetNavigatorUADataBrands(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Mobile returns the value of property "NavigatorUAData.mobile".
//
// It returns ok=false if there is no such property.
func (this NavigatorUAData) Mobile() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigatorUADataMobile(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Platform returns the value of property "NavigatorUAData.platform".
//
// It returns ok=false if there is no such property.
func (this NavigatorUAData) Platform() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorUADataPlatform(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetHighEntropyValues returns true if the method "NavigatorUAData.getHighEntropyValues" exists.
func (this NavigatorUAData) HasFuncGetHighEntropyValues() bool {
	return js.True == bindings.HasFuncNavigatorUADataGetHighEntropyValues(
		this.ref,
	)
}

// FuncGetHighEntropyValues returns the method "NavigatorUAData.getHighEntropyValues".
func (this NavigatorUAData) FuncGetHighEntropyValues() (fn js.Func[func(hints js.Array[js.String]) js.Promise[UADataValues]]) {
	bindings.FuncNavigatorUADataGetHighEntropyValues(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetHighEntropyValues calls the method "NavigatorUAData.getHighEntropyValues".
func (this NavigatorUAData) GetHighEntropyValues(hints js.Array[js.String]) (ret js.Promise[UADataValues]) {
	bindings.CallNavigatorUADataGetHighEntropyValues(
		this.ref, js.Pointer(&ret),
		hints.Ref(),
	)

	return
}

// TryGetHighEntropyValues calls the method "NavigatorUAData.getHighEntropyValues"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NavigatorUAData) TryGetHighEntropyValues(hints js.Array[js.String]) (ret js.Promise[UADataValues], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorUADataGetHighEntropyValues(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		hints.Ref(),
	)

	return
}

// HasFuncToJSON returns true if the method "NavigatorUAData.toJSON" exists.
func (this NavigatorUAData) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncNavigatorUADataToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "NavigatorUAData.toJSON".
func (this NavigatorUAData) FuncToJSON() (fn js.Func[func() UALowEntropyJSON]) {
	bindings.FuncNavigatorUADataToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "NavigatorUAData.toJSON".
func (this NavigatorUAData) ToJSON() (ret UALowEntropyJSON) {
	bindings.CallNavigatorUADataToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "NavigatorUAData.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NavigatorUAData) TryToJSON() (ret UALowEntropyJSON, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorUADataToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type StorageEstimate struct {
	// Usage is "StorageEstimate.usage"
	//
	// Optional
	//
	// NOTE: FFI_USE_Usage MUST be set to true to make this field effective.
	Usage uint64
	// Quota is "StorageEstimate.quota"
	//
	// Optional
	//
	// NOTE: FFI_USE_Quota MUST be set to true to make this field effective.
	Quota uint64

	FFI_USE_Usage bool // for Usage.
	FFI_USE_Quota bool // for Quota.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StorageEstimate with all fields set.
func (p StorageEstimate) FromRef(ref js.Ref) StorageEstimate {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StorageEstimate in the application heap.
func (p StorageEstimate) New() js.Ref {
	return bindings.StorageEstimateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StorageEstimate) UpdateFrom(ref js.Ref) {
	bindings.StorageEstimateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StorageEstimate) Update(ref js.Ref) {
	bindings.StorageEstimateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StorageEstimate) FreeMembers(recursive bool) {
}

type StorageManager struct {
	ref js.Ref
}

func (this StorageManager) Once() StorageManager {
	this.ref.Once()
	return this
}

func (this StorageManager) Ref() js.Ref {
	return this.ref
}

func (this StorageManager) FromRef(ref js.Ref) StorageManager {
	this.ref = ref
	return this
}

func (this StorageManager) Free() {
	this.ref.Free()
}

// HasFuncPersisted returns true if the method "StorageManager.persisted" exists.
func (this StorageManager) HasFuncPersisted() bool {
	return js.True == bindings.HasFuncStorageManagerPersisted(
		this.ref,
	)
}

// FuncPersisted returns the method "StorageManager.persisted".
func (this StorageManager) FuncPersisted() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncStorageManagerPersisted(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Persisted calls the method "StorageManager.persisted".
func (this StorageManager) Persisted() (ret js.Promise[js.Boolean]) {
	bindings.CallStorageManagerPersisted(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPersisted calls the method "StorageManager.persisted"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageManager) TryPersisted() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageManagerPersisted(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPersist returns true if the method "StorageManager.persist" exists.
func (this StorageManager) HasFuncPersist() bool {
	return js.True == bindings.HasFuncStorageManagerPersist(
		this.ref,
	)
}

// FuncPersist returns the method "StorageManager.persist".
func (this StorageManager) FuncPersist() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncStorageManagerPersist(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Persist calls the method "StorageManager.persist".
func (this StorageManager) Persist() (ret js.Promise[js.Boolean]) {
	bindings.CallStorageManagerPersist(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPersist calls the method "StorageManager.persist"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageManager) TryPersist() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageManagerPersist(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncEstimate returns true if the method "StorageManager.estimate" exists.
func (this StorageManager) HasFuncEstimate() bool {
	return js.True == bindings.HasFuncStorageManagerEstimate(
		this.ref,
	)
}

// FuncEstimate returns the method "StorageManager.estimate".
func (this StorageManager) FuncEstimate() (fn js.Func[func() js.Promise[StorageEstimate]]) {
	bindings.FuncStorageManagerEstimate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Estimate calls the method "StorageManager.estimate".
func (this StorageManager) Estimate() (ret js.Promise[StorageEstimate]) {
	bindings.CallStorageManagerEstimate(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryEstimate calls the method "StorageManager.estimate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageManager) TryEstimate() (ret js.Promise[StorageEstimate], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageManagerEstimate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDirectory returns true if the method "StorageManager.getDirectory" exists.
func (this StorageManager) HasFuncGetDirectory() bool {
	return js.True == bindings.HasFuncStorageManagerGetDirectory(
		this.ref,
	)
}

// FuncGetDirectory returns the method "StorageManager.getDirectory".
func (this StorageManager) FuncGetDirectory() (fn js.Func[func() js.Promise[FileSystemDirectoryHandle]]) {
	bindings.FuncStorageManagerGetDirectory(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDirectory calls the method "StorageManager.getDirectory".
func (this StorageManager) GetDirectory() (ret js.Promise[FileSystemDirectoryHandle]) {
	bindings.CallStorageManagerGetDirectory(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetDirectory calls the method "StorageManager.getDirectory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageManager) TryGetDirectory() (ret js.Promise[FileSystemDirectoryHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageManagerGetDirectory(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type StorageBucketDurability uint32

const (
	_ StorageBucketDurability = iota

	StorageBucketDurability_STRICT
	StorageBucketDurability_RELAXED
)

func (StorageBucketDurability) FromRef(str js.Ref) StorageBucketDurability {
	return StorageBucketDurability(bindings.ConstOfStorageBucketDurability(str))
}

func (x StorageBucketDurability) String() (string, bool) {
	switch x {
	case StorageBucketDurability_STRICT:
		return "strict", true
	case StorageBucketDurability_RELAXED:
		return "relaxed", true
	default:
		return "", false
	}
}

type IDBOpenDBRequest struct {
	IDBRequest
}

func (this IDBOpenDBRequest) Once() IDBOpenDBRequest {
	this.ref.Once()
	return this
}

func (this IDBOpenDBRequest) Ref() js.Ref {
	return this.IDBRequest.Ref()
}

func (this IDBOpenDBRequest) FromRef(ref js.Ref) IDBOpenDBRequest {
	this.IDBRequest = this.IDBRequest.FromRef(ref)
	return this
}

func (this IDBOpenDBRequest) Free() {
	this.ref.Free()
}

type IDBDatabaseInfo struct {
	// Name is "IDBDatabaseInfo.name"
	//
	// Optional
	Name js.String
	// Version is "IDBDatabaseInfo.version"
	//
	// Optional
	//
	// NOTE: FFI_USE_Version MUST be set to true to make this field effective.
	Version uint64

	FFI_USE_Version bool // for Version.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IDBDatabaseInfo with all fields set.
func (p IDBDatabaseInfo) FromRef(ref js.Ref) IDBDatabaseInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IDBDatabaseInfo in the application heap.
func (p IDBDatabaseInfo) New() js.Ref {
	return bindings.IDBDatabaseInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IDBDatabaseInfo) UpdateFrom(ref js.Ref) {
	bindings.IDBDatabaseInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IDBDatabaseInfo) Update(ref js.Ref) {
	bindings.IDBDatabaseInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IDBDatabaseInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
}

type IDBFactory struct {
	ref js.Ref
}

func (this IDBFactory) Once() IDBFactory {
	this.ref.Once()
	return this
}

func (this IDBFactory) Ref() js.Ref {
	return this.ref
}

func (this IDBFactory) FromRef(ref js.Ref) IDBFactory {
	this.ref = ref
	return this
}

func (this IDBFactory) Free() {
	this.ref.Free()
}

// HasFuncOpen returns true if the method "IDBFactory.open" exists.
func (this IDBFactory) HasFuncOpen() bool {
	return js.True == bindings.HasFuncIDBFactoryOpen(
		this.ref,
	)
}

// FuncOpen returns the method "IDBFactory.open".
func (this IDBFactory) FuncOpen() (fn js.Func[func(name js.String, version uint64) IDBOpenDBRequest]) {
	bindings.FuncIDBFactoryOpen(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open calls the method "IDBFactory.open".
func (this IDBFactory) Open(name js.String, version uint64) (ret IDBOpenDBRequest) {
	bindings.CallIDBFactoryOpen(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		float64(version),
	)

	return
}

// TryOpen calls the method "IDBFactory.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBFactory) TryOpen(name js.String, version uint64) (ret IDBOpenDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBFactoryOpen(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		float64(version),
	)

	return
}

// HasFuncOpen1 returns true if the method "IDBFactory.open" exists.
func (this IDBFactory) HasFuncOpen1() bool {
	return js.True == bindings.HasFuncIDBFactoryOpen1(
		this.ref,
	)
}

// FuncOpen1 returns the method "IDBFactory.open".
func (this IDBFactory) FuncOpen1() (fn js.Func[func(name js.String) IDBOpenDBRequest]) {
	bindings.FuncIDBFactoryOpen1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open1 calls the method "IDBFactory.open".
func (this IDBFactory) Open1(name js.String) (ret IDBOpenDBRequest) {
	bindings.CallIDBFactoryOpen1(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryOpen1 calls the method "IDBFactory.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBFactory) TryOpen1(name js.String) (ret IDBOpenDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBFactoryOpen1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncDeleteDatabase returns true if the method "IDBFactory.deleteDatabase" exists.
func (this IDBFactory) HasFuncDeleteDatabase() bool {
	return js.True == bindings.HasFuncIDBFactoryDeleteDatabase(
		this.ref,
	)
}

// FuncDeleteDatabase returns the method "IDBFactory.deleteDatabase".
func (this IDBFactory) FuncDeleteDatabase() (fn js.Func[func(name js.String) IDBOpenDBRequest]) {
	bindings.FuncIDBFactoryDeleteDatabase(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteDatabase calls the method "IDBFactory.deleteDatabase".
func (this IDBFactory) DeleteDatabase(name js.String) (ret IDBOpenDBRequest) {
	bindings.CallIDBFactoryDeleteDatabase(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryDeleteDatabase calls the method "IDBFactory.deleteDatabase"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBFactory) TryDeleteDatabase(name js.String) (ret IDBOpenDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBFactoryDeleteDatabase(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncDatabases returns true if the method "IDBFactory.databases" exists.
func (this IDBFactory) HasFuncDatabases() bool {
	return js.True == bindings.HasFuncIDBFactoryDatabases(
		this.ref,
	)
}

// FuncDatabases returns the method "IDBFactory.databases".
func (this IDBFactory) FuncDatabases() (fn js.Func[func() js.Promise[js.Array[IDBDatabaseInfo]]]) {
	bindings.FuncIDBFactoryDatabases(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Databases calls the method "IDBFactory.databases".
func (this IDBFactory) Databases() (ret js.Promise[js.Array[IDBDatabaseInfo]]) {
	bindings.CallIDBFactoryDatabases(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDatabases calls the method "IDBFactory.databases"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBFactory) TryDatabases() (ret js.Promise[js.Array[IDBDatabaseInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBFactoryDatabases(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCmp returns true if the method "IDBFactory.cmp" exists.
func (this IDBFactory) HasFuncCmp() bool {
	return js.True == bindings.HasFuncIDBFactoryCmp(
		this.ref,
	)
}

// FuncCmp returns the method "IDBFactory.cmp".
func (this IDBFactory) FuncCmp() (fn js.Func[func(first js.Any, second js.Any) int16]) {
	bindings.FuncIDBFactoryCmp(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Cmp calls the method "IDBFactory.cmp".
func (this IDBFactory) Cmp(first js.Any, second js.Any) (ret int16) {
	bindings.CallIDBFactoryCmp(
		this.ref, js.Pointer(&ret),
		first.Ref(),
		second.Ref(),
	)

	return
}

// TryCmp calls the method "IDBFactory.cmp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBFactory) TryCmp(first js.Any, second js.Any) (ret int16, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBFactoryCmp(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		first.Ref(),
		second.Ref(),
	)

	return
}

type StorageBucket struct {
	ref js.Ref
}

func (this StorageBucket) Once() StorageBucket {
	this.ref.Once()
	return this
}

func (this StorageBucket) Ref() js.Ref {
	return this.ref
}

func (this StorageBucket) FromRef(ref js.Ref) StorageBucket {
	this.ref = ref
	return this
}

func (this StorageBucket) Free() {
	this.ref.Free()
}

// Name returns the value of property "StorageBucket.name".
//
// It returns ok=false if there is no such property.
func (this StorageBucket) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetStorageBucketName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IndexedDB returns the value of property "StorageBucket.indexedDB".
//
// It returns ok=false if there is no such property.
func (this StorageBucket) IndexedDB() (ret IDBFactory, ok bool) {
	ok = js.True == bindings.GetStorageBucketIndexedDB(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Caches returns the value of property "StorageBucket.caches".
//
// It returns ok=false if there is no such property.
func (this StorageBucket) Caches() (ret CacheStorage, ok bool) {
	ok = js.True == bindings.GetStorageBucketCaches(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncPersist returns true if the method "StorageBucket.persist" exists.
func (this StorageBucket) HasFuncPersist() bool {
	return js.True == bindings.HasFuncStorageBucketPersist(
		this.ref,
	)
}

// FuncPersist returns the method "StorageBucket.persist".
func (this StorageBucket) FuncPersist() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncStorageBucketPersist(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Persist calls the method "StorageBucket.persist".
func (this StorageBucket) Persist() (ret js.Promise[js.Boolean]) {
	bindings.CallStorageBucketPersist(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPersist calls the method "StorageBucket.persist"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucket) TryPersist() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketPersist(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPersisted returns true if the method "StorageBucket.persisted" exists.
func (this StorageBucket) HasFuncPersisted() bool {
	return js.True == bindings.HasFuncStorageBucketPersisted(
		this.ref,
	)
}

// FuncPersisted returns the method "StorageBucket.persisted".
func (this StorageBucket) FuncPersisted() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncStorageBucketPersisted(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Persisted calls the method "StorageBucket.persisted".
func (this StorageBucket) Persisted() (ret js.Promise[js.Boolean]) {
	bindings.CallStorageBucketPersisted(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPersisted calls the method "StorageBucket.persisted"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucket) TryPersisted() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketPersisted(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncEstimate returns true if the method "StorageBucket.estimate" exists.
func (this StorageBucket) HasFuncEstimate() bool {
	return js.True == bindings.HasFuncStorageBucketEstimate(
		this.ref,
	)
}

// FuncEstimate returns the method "StorageBucket.estimate".
func (this StorageBucket) FuncEstimate() (fn js.Func[func() js.Promise[StorageEstimate]]) {
	bindings.FuncStorageBucketEstimate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Estimate calls the method "StorageBucket.estimate".
func (this StorageBucket) Estimate() (ret js.Promise[StorageEstimate]) {
	bindings.CallStorageBucketEstimate(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryEstimate calls the method "StorageBucket.estimate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucket) TryEstimate() (ret js.Promise[StorageEstimate], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketEstimate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDurability returns true if the method "StorageBucket.durability" exists.
func (this StorageBucket) HasFuncDurability() bool {
	return js.True == bindings.HasFuncStorageBucketDurability(
		this.ref,
	)
}

// FuncDurability returns the method "StorageBucket.durability".
func (this StorageBucket) FuncDurability() (fn js.Func[func() js.Promise[StorageBucketDurability]]) {
	bindings.FuncStorageBucketDurability(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Durability calls the method "StorageBucket.durability".
func (this StorageBucket) Durability() (ret js.Promise[StorageBucketDurability]) {
	bindings.CallStorageBucketDurability(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDurability calls the method "StorageBucket.durability"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucket) TryDurability() (ret js.Promise[StorageBucketDurability], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketDurability(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetExpires returns true if the method "StorageBucket.setExpires" exists.
func (this StorageBucket) HasFuncSetExpires() bool {
	return js.True == bindings.HasFuncStorageBucketSetExpires(
		this.ref,
	)
}

// FuncSetExpires returns the method "StorageBucket.setExpires".
func (this StorageBucket) FuncSetExpires() (fn js.Func[func(expires DOMHighResTimeStamp) js.Promise[js.Void]]) {
	bindings.FuncStorageBucketSetExpires(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetExpires calls the method "StorageBucket.setExpires".
func (this StorageBucket) SetExpires(expires DOMHighResTimeStamp) (ret js.Promise[js.Void]) {
	bindings.CallStorageBucketSetExpires(
		this.ref, js.Pointer(&ret),
		float64(expires),
	)

	return
}

// TrySetExpires calls the method "StorageBucket.setExpires"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucket) TrySetExpires(expires DOMHighResTimeStamp) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketSetExpires(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(expires),
	)

	return
}

// HasFuncExpires returns true if the method "StorageBucket.expires" exists.
func (this StorageBucket) HasFuncExpires() bool {
	return js.True == bindings.HasFuncStorageBucketExpires(
		this.ref,
	)
}

// FuncExpires returns the method "StorageBucket.expires".
func (this StorageBucket) FuncExpires() (fn js.Func[func() js.Promise[js.Number[DOMHighResTimeStamp]]]) {
	bindings.FuncStorageBucketExpires(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Expires calls the method "StorageBucket.expires".
func (this StorageBucket) Expires() (ret js.Promise[js.Number[DOMHighResTimeStamp]]) {
	bindings.CallStorageBucketExpires(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryExpires calls the method "StorageBucket.expires"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucket) TryExpires() (ret js.Promise[js.Number[DOMHighResTimeStamp]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketExpires(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDirectory returns true if the method "StorageBucket.getDirectory" exists.
func (this StorageBucket) HasFuncGetDirectory() bool {
	return js.True == bindings.HasFuncStorageBucketGetDirectory(
		this.ref,
	)
}

// FuncGetDirectory returns the method "StorageBucket.getDirectory".
func (this StorageBucket) FuncGetDirectory() (fn js.Func[func() js.Promise[FileSystemDirectoryHandle]]) {
	bindings.FuncStorageBucketGetDirectory(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDirectory calls the method "StorageBucket.getDirectory".
func (this StorageBucket) GetDirectory() (ret js.Promise[FileSystemDirectoryHandle]) {
	bindings.CallStorageBucketGetDirectory(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetDirectory calls the method "StorageBucket.getDirectory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucket) TryGetDirectory() (ret js.Promise[FileSystemDirectoryHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketGetDirectory(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type StorageBucketOptions struct {
	// Persisted is "StorageBucketOptions.persisted"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_Persisted MUST be set to true to make this field effective.
	Persisted bool
	// Durability is "StorageBucketOptions.durability"
	//
	// Optional, defaults to null.
	Durability StorageBucketDurability
	// Quota is "StorageBucketOptions.quota"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_Quota MUST be set to true to make this field effective.
	Quota uint64
	// Expires is "StorageBucketOptions.expires"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_Expires MUST be set to true to make this field effective.
	Expires DOMHighResTimeStamp

	FFI_USE_Persisted bool // for Persisted.
	FFI_USE_Quota     bool // for Quota.
	FFI_USE_Expires   bool // for Expires.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StorageBucketOptions with all fields set.
func (p StorageBucketOptions) FromRef(ref js.Ref) StorageBucketOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StorageBucketOptions in the application heap.
func (p StorageBucketOptions) New() js.Ref {
	return bindings.StorageBucketOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StorageBucketOptions) UpdateFrom(ref js.Ref) {
	bindings.StorageBucketOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StorageBucketOptions) Update(ref js.Ref) {
	bindings.StorageBucketOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StorageBucketOptions) FreeMembers(recursive bool) {
}

type StorageBucketManager struct {
	ref js.Ref
}

func (this StorageBucketManager) Once() StorageBucketManager {
	this.ref.Once()
	return this
}

func (this StorageBucketManager) Ref() js.Ref {
	return this.ref
}

func (this StorageBucketManager) FromRef(ref js.Ref) StorageBucketManager {
	this.ref = ref
	return this
}

func (this StorageBucketManager) Free() {
	this.ref.Free()
}

// HasFuncOpen returns true if the method "StorageBucketManager.open" exists.
func (this StorageBucketManager) HasFuncOpen() bool {
	return js.True == bindings.HasFuncStorageBucketManagerOpen(
		this.ref,
	)
}

// FuncOpen returns the method "StorageBucketManager.open".
func (this StorageBucketManager) FuncOpen() (fn js.Func[func(name js.String, options StorageBucketOptions) js.Promise[StorageBucket]]) {
	bindings.FuncStorageBucketManagerOpen(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open calls the method "StorageBucketManager.open".
func (this StorageBucketManager) Open(name js.String, options StorageBucketOptions) (ret js.Promise[StorageBucket]) {
	bindings.CallStorageBucketManagerOpen(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryOpen calls the method "StorageBucketManager.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucketManager) TryOpen(name js.String, options StorageBucketOptions) (ret js.Promise[StorageBucket], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketManagerOpen(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncOpen1 returns true if the method "StorageBucketManager.open" exists.
func (this StorageBucketManager) HasFuncOpen1() bool {
	return js.True == bindings.HasFuncStorageBucketManagerOpen1(
		this.ref,
	)
}

// FuncOpen1 returns the method "StorageBucketManager.open".
func (this StorageBucketManager) FuncOpen1() (fn js.Func[func(name js.String) js.Promise[StorageBucket]]) {
	bindings.FuncStorageBucketManagerOpen1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open1 calls the method "StorageBucketManager.open".
func (this StorageBucketManager) Open1(name js.String) (ret js.Promise[StorageBucket]) {
	bindings.CallStorageBucketManagerOpen1(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryOpen1 calls the method "StorageBucketManager.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucketManager) TryOpen1(name js.String) (ret js.Promise[StorageBucket], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketManagerOpen1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncKeys returns true if the method "StorageBucketManager.keys" exists.
func (this StorageBucketManager) HasFuncKeys() bool {
	return js.True == bindings.HasFuncStorageBucketManagerKeys(
		this.ref,
	)
}

// FuncKeys returns the method "StorageBucketManager.keys".
func (this StorageBucketManager) FuncKeys() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	bindings.FuncStorageBucketManagerKeys(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Keys calls the method "StorageBucketManager.keys".
func (this StorageBucketManager) Keys() (ret js.Promise[js.Array[js.String]]) {
	bindings.CallStorageBucketManagerKeys(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryKeys calls the method "StorageBucketManager.keys"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucketManager) TryKeys() (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketManagerKeys(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDelete returns true if the method "StorageBucketManager.delete" exists.
func (this StorageBucketManager) HasFuncDelete() bool {
	return js.True == bindings.HasFuncStorageBucketManagerDelete(
		this.ref,
	)
}

// FuncDelete returns the method "StorageBucketManager.delete".
func (this StorageBucketManager) FuncDelete() (fn js.Func[func(name js.String) js.Promise[js.Void]]) {
	bindings.FuncStorageBucketManagerDelete(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Delete calls the method "StorageBucketManager.delete".
func (this StorageBucketManager) Delete(name js.String) (ret js.Promise[js.Void]) {
	bindings.CallStorageBucketManagerDelete(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryDelete calls the method "StorageBucketManager.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucketManager) TryDelete(name js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketManagerDelete(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

type LockGrantedCallbackFunc func(this js.Ref, lock Lock) js.Ref

func (fn LockGrantedCallbackFunc) Register() js.Func[func(lock Lock) js.Promise[js.Any]] {
	return js.RegisterCallback[func(lock Lock) js.Promise[js.Any]](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn LockGrantedCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		Lock{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LockGrantedCallback[T any] struct {
	Fn  func(arg T, this js.Ref, lock Lock) js.Ref
	Arg T
}

func (cb *LockGrantedCallback[T]) Register() js.Func[func(lock Lock) js.Promise[js.Any]] {
	return js.RegisterCallback[func(lock Lock) js.Promise[js.Any]](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *LockGrantedCallback[T]) DispatchCallback(
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

		Lock{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LockMode uint32

const (
	_ LockMode = iota

	LockMode_SHARED
	LockMode_EXCLUSIVE
)

func (LockMode) FromRef(str js.Ref) LockMode {
	return LockMode(bindings.ConstOfLockMode(str))
}

func (x LockMode) String() (string, bool) {
	switch x {
	case LockMode_SHARED:
		return "shared", true
	case LockMode_EXCLUSIVE:
		return "exclusive", true
	default:
		return "", false
	}
}

type Lock struct {
	ref js.Ref
}

func (this Lock) Once() Lock {
	this.ref.Once()
	return this
}

func (this Lock) Ref() js.Ref {
	return this.ref
}

func (this Lock) FromRef(ref js.Ref) Lock {
	this.ref = ref
	return this
}

func (this Lock) Free() {
	this.ref.Free()
}

// Name returns the value of property "Lock.name".
//
// It returns ok=false if there is no such property.
func (this Lock) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLockName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Mode returns the value of property "Lock.mode".
//
// It returns ok=false if there is no such property.
func (this Lock) Mode() (ret LockMode, ok bool) {
	ok = js.True == bindings.GetLockMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

type LockOptions struct {
	// Mode is "LockOptions.mode"
	//
	// Optional, defaults to "exclusive".
	Mode LockMode
	// IfAvailable is "LockOptions.ifAvailable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IfAvailable MUST be set to true to make this field effective.
	IfAvailable bool
	// Steal is "LockOptions.steal"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Steal MUST be set to true to make this field effective.
	Steal bool
	// Signal is "LockOptions.signal"
	//
	// Optional
	Signal AbortSignal

	FFI_USE_IfAvailable bool // for IfAvailable.
	FFI_USE_Steal       bool // for Steal.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LockOptions with all fields set.
func (p LockOptions) FromRef(ref js.Ref) LockOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LockOptions in the application heap.
func (p LockOptions) New() js.Ref {
	return bindings.LockOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LockOptions) UpdateFrom(ref js.Ref) {
	bindings.LockOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LockOptions) Update(ref js.Ref) {
	bindings.LockOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LockOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Signal.Ref(),
	)
	p.Signal = p.Signal.FromRef(js.Undefined)
}

type LockInfo struct {
	// Name is "LockInfo.name"
	//
	// Optional
	Name js.String
	// Mode is "LockInfo.mode"
	//
	// Optional
	Mode LockMode
	// ClientId is "LockInfo.clientId"
	//
	// Optional
	ClientId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LockInfo with all fields set.
func (p LockInfo) FromRef(ref js.Ref) LockInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LockInfo in the application heap.
func (p LockInfo) New() js.Ref {
	return bindings.LockInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LockInfo) UpdateFrom(ref js.Ref) {
	bindings.LockInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LockInfo) Update(ref js.Ref) {
	bindings.LockInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LockInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.ClientId.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.ClientId = p.ClientId.FromRef(js.Undefined)
}

type LockManagerSnapshot struct {
	// Held is "LockManagerSnapshot.held"
	//
	// Optional
	Held js.Array[LockInfo]
	// Pending is "LockManagerSnapshot.pending"
	//
	// Optional
	Pending js.Array[LockInfo]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LockManagerSnapshot with all fields set.
func (p LockManagerSnapshot) FromRef(ref js.Ref) LockManagerSnapshot {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LockManagerSnapshot in the application heap.
func (p LockManagerSnapshot) New() js.Ref {
	return bindings.LockManagerSnapshotJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LockManagerSnapshot) UpdateFrom(ref js.Ref) {
	bindings.LockManagerSnapshotJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LockManagerSnapshot) Update(ref js.Ref) {
	bindings.LockManagerSnapshotJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LockManagerSnapshot) FreeMembers(recursive bool) {
	js.Free(
		p.Held.Ref(),
		p.Pending.Ref(),
	)
	p.Held = p.Held.FromRef(js.Undefined)
	p.Pending = p.Pending.FromRef(js.Undefined)
}

type LockManager struct {
	ref js.Ref
}

func (this LockManager) Once() LockManager {
	this.ref.Once()
	return this
}

func (this LockManager) Ref() js.Ref {
	return this.ref
}

func (this LockManager) FromRef(ref js.Ref) LockManager {
	this.ref = ref
	return this
}

func (this LockManager) Free() {
	this.ref.Free()
}

// HasFuncRequest returns true if the method "LockManager.request" exists.
func (this LockManager) HasFuncRequest() bool {
	return js.True == bindings.HasFuncLockManagerRequest(
		this.ref,
	)
}

// FuncRequest returns the method "LockManager.request".
func (this LockManager) FuncRequest() (fn js.Func[func(name js.String, callback js.Func[func(lock Lock) js.Promise[js.Any]]) js.Promise[js.Any]]) {
	bindings.FuncLockManagerRequest(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Request calls the method "LockManager.request".
func (this LockManager) Request(name js.String, callback js.Func[func(lock Lock) js.Promise[js.Any]]) (ret js.Promise[js.Any]) {
	bindings.CallLockManagerRequest(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		callback.Ref(),
	)

	return
}

// TryRequest calls the method "LockManager.request"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this LockManager) TryRequest(name js.String, callback js.Func[func(lock Lock) js.Promise[js.Any]]) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLockManagerRequest(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncRequest1 returns true if the method "LockManager.request" exists.
func (this LockManager) HasFuncRequest1() bool {
	return js.True == bindings.HasFuncLockManagerRequest1(
		this.ref,
	)
}

// FuncRequest1 returns the method "LockManager.request".
func (this LockManager) FuncRequest1() (fn js.Func[func(name js.String, options LockOptions, callback js.Func[func(lock Lock) js.Promise[js.Any]]) js.Promise[js.Any]]) {
	bindings.FuncLockManagerRequest1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Request1 calls the method "LockManager.request".
func (this LockManager) Request1(name js.String, options LockOptions, callback js.Func[func(lock Lock) js.Promise[js.Any]]) (ret js.Promise[js.Any]) {
	bindings.CallLockManagerRequest1(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// TryRequest1 calls the method "LockManager.request"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this LockManager) TryRequest1(name js.String, options LockOptions, callback js.Func[func(lock Lock) js.Promise[js.Any]]) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLockManagerRequest1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// HasFuncQuery returns true if the method "LockManager.query" exists.
func (this LockManager) HasFuncQuery() bool {
	return js.True == bindings.HasFuncLockManagerQuery(
		this.ref,
	)
}

// FuncQuery returns the method "LockManager.query".
func (this LockManager) FuncQuery() (fn js.Func[func() js.Promise[LockManagerSnapshot]]) {
	bindings.FuncLockManagerQuery(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Query calls the method "LockManager.query".
func (this LockManager) Query() (ret js.Promise[LockManagerSnapshot]) {
	bindings.CallLockManagerQuery(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryQuery calls the method "LockManager.query"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this LockManager) TryQuery() (ret js.Promise[LockManagerSnapshot], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLockManagerQuery(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type MimeType struct {
	ref js.Ref
}

func (this MimeType) Once() MimeType {
	this.ref.Once()
	return this
}

func (this MimeType) Ref() js.Ref {
	return this.ref
}

func (this MimeType) FromRef(ref js.Ref) MimeType {
	this.ref = ref
	return this
}

func (this MimeType) Free() {
	this.ref.Free()
}

// Type returns the value of property "MimeType.type".
//
// It returns ok=false if there is no such property.
func (this MimeType) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMimeTypeType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Description returns the value of property "MimeType.description".
//
// It returns ok=false if there is no such property.
func (this MimeType) Description() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMimeTypeDescription(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Suffixes returns the value of property "MimeType.suffixes".
//
// It returns ok=false if there is no such property.
func (this MimeType) Suffixes() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMimeTypeSuffixes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EnabledPlugin returns the value of property "MimeType.enabledPlugin".
//
// It returns ok=false if there is no such property.
func (this MimeType) EnabledPlugin() (ret Plugin, ok bool) {
	ok = js.True == bindings.GetMimeTypeEnabledPlugin(
		this.ref, js.Pointer(&ret),
	)
	return
}

type Plugin struct {
	ref js.Ref
}

func (this Plugin) Once() Plugin {
	this.ref.Once()
	return this
}

func (this Plugin) Ref() js.Ref {
	return this.ref
}

func (this Plugin) FromRef(ref js.Ref) Plugin {
	this.ref = ref
	return this
}

func (this Plugin) Free() {
	this.ref.Free()
}

// Name returns the value of property "Plugin.name".
//
// It returns ok=false if there is no such property.
func (this Plugin) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPluginName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Description returns the value of property "Plugin.description".
//
// It returns ok=false if there is no such property.
func (this Plugin) Description() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPluginDescription(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Filename returns the value of property "Plugin.filename".
//
// It returns ok=false if there is no such property.
func (this Plugin) Filename() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPluginFilename(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Length returns the value of property "Plugin.length".
//
// It returns ok=false if there is no such property.
func (this Plugin) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetPluginLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "Plugin.item" exists.
func (this Plugin) HasFuncItem() bool {
	return js.True == bindings.HasFuncPluginItem(
		this.ref,
	)
}

// FuncItem returns the method "Plugin.item".
func (this Plugin) FuncItem() (fn js.Func[func(index uint32) MimeType]) {
	bindings.FuncPluginItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "Plugin.item".
func (this Plugin) Item(index uint32) (ret MimeType) {
	bindings.CallPluginItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "Plugin.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Plugin) TryItem(index uint32) (ret MimeType, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPluginItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncNamedItem returns true if the method "Plugin.namedItem" exists.
func (this Plugin) HasFuncNamedItem() bool {
	return js.True == bindings.HasFuncPluginNamedItem(
		this.ref,
	)
}

// FuncNamedItem returns the method "Plugin.namedItem".
func (this Plugin) FuncNamedItem() (fn js.Func[func(name js.String) MimeType]) {
	bindings.FuncPluginNamedItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// NamedItem calls the method "Plugin.namedItem".
func (this Plugin) NamedItem(name js.String) (ret MimeType) {
	bindings.CallPluginNamedItem(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryNamedItem calls the method "Plugin.namedItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Plugin) TryNamedItem(name js.String) (ret MimeType, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPluginNamedItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

type PluginArray struct {
	ref js.Ref
}

func (this PluginArray) Once() PluginArray {
	this.ref.Once()
	return this
}

func (this PluginArray) Ref() js.Ref {
	return this.ref
}

func (this PluginArray) FromRef(ref js.Ref) PluginArray {
	this.ref = ref
	return this
}

func (this PluginArray) Free() {
	this.ref.Free()
}

// Length returns the value of property "PluginArray.length".
//
// It returns ok=false if there is no such property.
func (this PluginArray) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetPluginArrayLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncRefresh returns true if the method "PluginArray.refresh" exists.
func (this PluginArray) HasFuncRefresh() bool {
	return js.True == bindings.HasFuncPluginArrayRefresh(
		this.ref,
	)
}

// FuncRefresh returns the method "PluginArray.refresh".
func (this PluginArray) FuncRefresh() (fn js.Func[func()]) {
	bindings.FuncPluginArrayRefresh(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Refresh calls the method "PluginArray.refresh".
func (this PluginArray) Refresh() (ret js.Void) {
	bindings.CallPluginArrayRefresh(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRefresh calls the method "PluginArray.refresh"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PluginArray) TryRefresh() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPluginArrayRefresh(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncItem returns true if the method "PluginArray.item" exists.
func (this PluginArray) HasFuncItem() bool {
	return js.True == bindings.HasFuncPluginArrayItem(
		this.ref,
	)
}

// FuncItem returns the method "PluginArray.item".
func (this PluginArray) FuncItem() (fn js.Func[func(index uint32) Plugin]) {
	bindings.FuncPluginArrayItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "PluginArray.item".
func (this PluginArray) Item(index uint32) (ret Plugin) {
	bindings.CallPluginArrayItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "PluginArray.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PluginArray) TryItem(index uint32) (ret Plugin, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPluginArrayItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncNamedItem returns true if the method "PluginArray.namedItem" exists.
func (this PluginArray) HasFuncNamedItem() bool {
	return js.True == bindings.HasFuncPluginArrayNamedItem(
		this.ref,
	)
}

// FuncNamedItem returns the method "PluginArray.namedItem".
func (this PluginArray) FuncNamedItem() (fn js.Func[func(name js.String) Plugin]) {
	bindings.FuncPluginArrayNamedItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// NamedItem calls the method "PluginArray.namedItem".
func (this PluginArray) NamedItem(name js.String) (ret Plugin) {
	bindings.CallPluginArrayNamedItem(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryNamedItem calls the method "PluginArray.namedItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PluginArray) TryNamedItem(name js.String) (ret Plugin, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPluginArrayNamedItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

type MimeTypeArray struct {
	ref js.Ref
}

func (this MimeTypeArray) Once() MimeTypeArray {
	this.ref.Once()
	return this
}

func (this MimeTypeArray) Ref() js.Ref {
	return this.ref
}

func (this MimeTypeArray) FromRef(ref js.Ref) MimeTypeArray {
	this.ref = ref
	return this
}

func (this MimeTypeArray) Free() {
	this.ref.Free()
}

// Length returns the value of property "MimeTypeArray.length".
//
// It returns ok=false if there is no such property.
func (this MimeTypeArray) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetMimeTypeArrayLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "MimeTypeArray.item" exists.
func (this MimeTypeArray) HasFuncItem() bool {
	return js.True == bindings.HasFuncMimeTypeArrayItem(
		this.ref,
	)
}

// FuncItem returns the method "MimeTypeArray.item".
func (this MimeTypeArray) FuncItem() (fn js.Func[func(index uint32) MimeType]) {
	bindings.FuncMimeTypeArrayItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "MimeTypeArray.item".
func (this MimeTypeArray) Item(index uint32) (ret MimeType) {
	bindings.CallMimeTypeArrayItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "MimeTypeArray.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MimeTypeArray) TryItem(index uint32) (ret MimeType, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMimeTypeArrayItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncNamedItem returns true if the method "MimeTypeArray.namedItem" exists.
func (this MimeTypeArray) HasFuncNamedItem() bool {
	return js.True == bindings.HasFuncMimeTypeArrayNamedItem(
		this.ref,
	)
}

// FuncNamedItem returns the method "MimeTypeArray.namedItem".
func (this MimeTypeArray) FuncNamedItem() (fn js.Func[func(name js.String) MimeType]) {
	bindings.FuncMimeTypeArrayNamedItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// NamedItem calls the method "MimeTypeArray.namedItem".
func (this MimeTypeArray) NamedItem(name js.String) (ret MimeType) {
	bindings.CallMimeTypeArrayNamedItem(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryNamedItem calls the method "MimeTypeArray.namedItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MimeTypeArray) TryNamedItem(name js.String) (ret MimeType, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMimeTypeArrayNamedItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

type MLNamedArrayBufferViews = js.Record[js.ArrayBufferView]

type MLComputeResult struct {
	// Inputs is "MLComputeResult.inputs"
	//
	// Optional
	Inputs MLNamedArrayBufferViews
	// Outputs is "MLComputeResult.outputs"
	//
	// Optional
	Outputs MLNamedArrayBufferViews

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLComputeResult with all fields set.
func (p MLComputeResult) FromRef(ref js.Ref) MLComputeResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLComputeResult in the application heap.
func (p MLComputeResult) New() js.Ref {
	return bindings.MLComputeResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLComputeResult) UpdateFrom(ref js.Ref) {
	bindings.MLComputeResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLComputeResult) Update(ref js.Ref) {
	bindings.MLComputeResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLComputeResult) FreeMembers(recursive bool) {
	js.Free(
		p.Inputs.Ref(),
		p.Outputs.Ref(),
	)
	p.Inputs = p.Inputs.FromRef(js.Undefined)
	p.Outputs = p.Outputs.FromRef(js.Undefined)
}

type MLGraph struct {
	ref js.Ref
}

func (this MLGraph) Once() MLGraph {
	this.ref.Once()
	return this
}

func (this MLGraph) Ref() js.Ref {
	return this.ref
}

func (this MLGraph) FromRef(ref js.Ref) MLGraph {
	this.ref = ref
	return this
}

func (this MLGraph) Free() {
	this.ref.Free()
}
