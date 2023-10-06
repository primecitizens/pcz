// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

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
func (p IsInputPendingOptions) UpdateFrom(ref js.Ref) {
	bindings.IsInputPendingOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IsInputPendingOptions) Update(ref js.Ref) {
	bindings.IsInputPendingOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type Scheduling struct {
	ref js.Ref
}

func (this Scheduling) Once() Scheduling {
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasIsInputPending returns true if the method "Scheduling.isInputPending" exists.
func (this Scheduling) HasIsInputPending() bool {
	return js.True == bindings.HasSchedulingIsInputPending(
		this.Ref(),
	)
}

// IsInputPendingFunc returns the method "Scheduling.isInputPending".
func (this Scheduling) IsInputPendingFunc() (fn js.Func[func(isInputPendingOptions IsInputPendingOptions) bool]) {
	return fn.FromRef(
		bindings.SchedulingIsInputPendingFunc(
			this.Ref(),
		),
	)
}

// IsInputPending calls the method "Scheduling.isInputPending".
func (this Scheduling) IsInputPending(isInputPendingOptions IsInputPendingOptions) (ret bool) {
	bindings.CallSchedulingIsInputPending(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&isInputPendingOptions),
	)

	return
}

// TryIsInputPending calls the method "Scheduling.isInputPending"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Scheduling) TryIsInputPending(isInputPendingOptions IsInputPendingOptions) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySchedulingIsInputPending(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&isInputPendingOptions),
	)

	return
}

// HasIsInputPending1 returns true if the method "Scheduling.isInputPending" exists.
func (this Scheduling) HasIsInputPending1() bool {
	return js.True == bindings.HasSchedulingIsInputPending1(
		this.Ref(),
	)
}

// IsInputPending1Func returns the method "Scheduling.isInputPending".
func (this Scheduling) IsInputPending1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.SchedulingIsInputPending1Func(
			this.Ref(),
		),
	)
}

// IsInputPending1 calls the method "Scheduling.isInputPending".
func (this Scheduling) IsInputPending1() (ret bool) {
	bindings.CallSchedulingIsInputPending1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryIsInputPending1 calls the method "Scheduling.isInputPending"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Scheduling) TryIsInputPending1() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySchedulingIsInputPending1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Released returns the value of property "WakeLockSentinel.released".
//
// It returns ok=false if there is no such property.
func (this WakeLockSentinel) Released() (ret bool, ok bool) {
	ok = js.True == bindings.GetWakeLockSentinelReleased(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "WakeLockSentinel.type".
//
// It returns ok=false if there is no such property.
func (this WakeLockSentinel) Type() (ret WakeLockType, ok bool) {
	ok = js.True == bindings.GetWakeLockSentinelType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRelease returns true if the method "WakeLockSentinel.release" exists.
func (this WakeLockSentinel) HasRelease() bool {
	return js.True == bindings.HasWakeLockSentinelRelease(
		this.Ref(),
	)
}

// ReleaseFunc returns the method "WakeLockSentinel.release".
func (this WakeLockSentinel) ReleaseFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WakeLockSentinelReleaseFunc(
			this.Ref(),
		),
	)
}

// Release calls the method "WakeLockSentinel.release".
func (this WakeLockSentinel) Release() (ret js.Promise[js.Void]) {
	bindings.CallWakeLockSentinelRelease(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRelease calls the method "WakeLockSentinel.release"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WakeLockSentinel) TryRelease() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWakeLockSentinelRelease(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type WakeLock struct {
	ref js.Ref
}

func (this WakeLock) Once() WakeLock {
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasRequest returns true if the method "WakeLock.request" exists.
func (this WakeLock) HasRequest() bool {
	return js.True == bindings.HasWakeLockRequest(
		this.Ref(),
	)
}

// RequestFunc returns the method "WakeLock.request".
func (this WakeLock) RequestFunc() (fn js.Func[func(typ WakeLockType) js.Promise[WakeLockSentinel]]) {
	return fn.FromRef(
		bindings.WakeLockRequestFunc(
			this.Ref(),
		),
	)
}

// Request calls the method "WakeLock.request".
func (this WakeLock) Request(typ WakeLockType) (ret js.Promise[WakeLockSentinel]) {
	bindings.CallWakeLockRequest(
		this.Ref(), js.Pointer(&ret),
		uint32(typ),
	)

	return
}

// TryRequest calls the method "WakeLock.request"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WakeLock) TryRequest(typ WakeLockType) (ret js.Promise[WakeLockSentinel], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWakeLockRequest(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
	)

	return
}

// HasRequest1 returns true if the method "WakeLock.request" exists.
func (this WakeLock) HasRequest1() bool {
	return js.True == bindings.HasWakeLockRequest1(
		this.Ref(),
	)
}

// Request1Func returns the method "WakeLock.request".
func (this WakeLock) Request1Func() (fn js.Func[func() js.Promise[WakeLockSentinel]]) {
	return fn.FromRef(
		bindings.WakeLockRequest1Func(
			this.Ref(),
		),
	)
}

// Request1 calls the method "WakeLock.request".
func (this WakeLock) Request1() (ret js.Promise[WakeLockSentinel]) {
	bindings.CallWakeLockRequest1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequest1 calls the method "WakeLock.request"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WakeLock) TryRequest1() (ret js.Promise[WakeLockSentinel], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWakeLockRequest1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
func (p PointerEventInit) UpdateFrom(ref js.Ref) {
	bindings.PointerEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PointerEventInit) Update(ref js.Ref) {
	bindings.PointerEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// PointerId returns the value of property "PointerEvent.pointerId".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) PointerId() (ret int32, ok bool) {
	ok = js.True == bindings.GetPointerEventPointerId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "PointerEvent.width".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) Width() (ret float64, ok bool) {
	ok = js.True == bindings.GetPointerEventWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "PointerEvent.height".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) Height() (ret float64, ok bool) {
	ok = js.True == bindings.GetPointerEventHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Pressure returns the value of property "PointerEvent.pressure".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) Pressure() (ret float32, ok bool) {
	ok = js.True == bindings.GetPointerEventPressure(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TangentialPressure returns the value of property "PointerEvent.tangentialPressure".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) TangentialPressure() (ret float32, ok bool) {
	ok = js.True == bindings.GetPointerEventTangentialPressure(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TiltX returns the value of property "PointerEvent.tiltX".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) TiltX() (ret int32, ok bool) {
	ok = js.True == bindings.GetPointerEventTiltX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TiltY returns the value of property "PointerEvent.tiltY".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) TiltY() (ret int32, ok bool) {
	ok = js.True == bindings.GetPointerEventTiltY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Twist returns the value of property "PointerEvent.twist".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) Twist() (ret int32, ok bool) {
	ok = js.True == bindings.GetPointerEventTwist(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AltitudeAngle returns the value of property "PointerEvent.altitudeAngle".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) AltitudeAngle() (ret float64, ok bool) {
	ok = js.True == bindings.GetPointerEventAltitudeAngle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AzimuthAngle returns the value of property "PointerEvent.azimuthAngle".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) AzimuthAngle() (ret float64, ok bool) {
	ok = js.True == bindings.GetPointerEventAzimuthAngle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PointerType returns the value of property "PointerEvent.pointerType".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) PointerType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPointerEventPointerType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsPrimary returns the value of property "PointerEvent.isPrimary".
//
// It returns ok=false if there is no such property.
func (this PointerEvent) IsPrimary() (ret bool, ok bool) {
	ok = js.True == bindings.GetPointerEventIsPrimary(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetCoalescedEvents returns true if the method "PointerEvent.getCoalescedEvents" exists.
func (this PointerEvent) HasGetCoalescedEvents() bool {
	return js.True == bindings.HasPointerEventGetCoalescedEvents(
		this.Ref(),
	)
}

// GetCoalescedEventsFunc returns the method "PointerEvent.getCoalescedEvents".
func (this PointerEvent) GetCoalescedEventsFunc() (fn js.Func[func() js.Array[PointerEvent]]) {
	return fn.FromRef(
		bindings.PointerEventGetCoalescedEventsFunc(
			this.Ref(),
		),
	)
}

// GetCoalescedEvents calls the method "PointerEvent.getCoalescedEvents".
func (this PointerEvent) GetCoalescedEvents() (ret js.Array[PointerEvent]) {
	bindings.CallPointerEventGetCoalescedEvents(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetCoalescedEvents calls the method "PointerEvent.getCoalescedEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PointerEvent) TryGetCoalescedEvents() (ret js.Array[PointerEvent], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPointerEventGetCoalescedEvents(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetPredictedEvents returns true if the method "PointerEvent.getPredictedEvents" exists.
func (this PointerEvent) HasGetPredictedEvents() bool {
	return js.True == bindings.HasPointerEventGetPredictedEvents(
		this.Ref(),
	)
}

// GetPredictedEventsFunc returns the method "PointerEvent.getPredictedEvents".
func (this PointerEvent) GetPredictedEventsFunc() (fn js.Func[func() js.Array[PointerEvent]]) {
	return fn.FromRef(
		bindings.PointerEventGetPredictedEventsFunc(
			this.Ref(),
		),
	)
}

// GetPredictedEvents calls the method "PointerEvent.getPredictedEvents".
func (this PointerEvent) GetPredictedEvents() (ret js.Array[PointerEvent]) {
	bindings.CallPointerEventGetPredictedEvents(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetPredictedEvents calls the method "PointerEvent.getPredictedEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PointerEvent) TryGetPredictedEvents() (ret js.Array[PointerEvent], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPointerEventGetPredictedEvents(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
func (p InkTrailStyle) UpdateFrom(ref js.Ref) {
	bindings.InkTrailStyleJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p InkTrailStyle) Update(ref js.Ref) {
	bindings.InkTrailStyleJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type InkPresenter struct {
	ref js.Ref
}

func (this InkPresenter) Once() InkPresenter {
	this.Ref().Once()
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
	this.Ref().Free()
}

// PresentationArea returns the value of property "InkPresenter.presentationArea".
//
// It returns ok=false if there is no such property.
func (this InkPresenter) PresentationArea() (ret Element, ok bool) {
	ok = js.True == bindings.GetInkPresenterPresentationArea(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ExpectedImprovement returns the value of property "InkPresenter.expectedImprovement".
//
// It returns ok=false if there is no such property.
func (this InkPresenter) ExpectedImprovement() (ret uint32, ok bool) {
	ok = js.True == bindings.GetInkPresenterExpectedImprovement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasUpdateInkTrailStartPoint returns true if the method "InkPresenter.updateInkTrailStartPoint" exists.
func (this InkPresenter) HasUpdateInkTrailStartPoint() bool {
	return js.True == bindings.HasInkPresenterUpdateInkTrailStartPoint(
		this.Ref(),
	)
}

// UpdateInkTrailStartPointFunc returns the method "InkPresenter.updateInkTrailStartPoint".
func (this InkPresenter) UpdateInkTrailStartPointFunc() (fn js.Func[func(event PointerEvent, style InkTrailStyle)]) {
	return fn.FromRef(
		bindings.InkPresenterUpdateInkTrailStartPointFunc(
			this.Ref(),
		),
	)
}

// UpdateInkTrailStartPoint calls the method "InkPresenter.updateInkTrailStartPoint".
func (this InkPresenter) UpdateInkTrailStartPoint(event PointerEvent, style InkTrailStyle) (ret js.Void) {
	bindings.CallInkPresenterUpdateInkTrailStartPoint(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
func (p InkPresenterParam) UpdateFrom(ref js.Ref) {
	bindings.InkPresenterParamJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p InkPresenterParam) Update(ref js.Ref) {
	bindings.InkPresenterParamJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type Ink struct {
	ref js.Ref
}

func (this Ink) Once() Ink {
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasRequestPresenter returns true if the method "Ink.requestPresenter" exists.
func (this Ink) HasRequestPresenter() bool {
	return js.True == bindings.HasInkRequestPresenter(
		this.Ref(),
	)
}

// RequestPresenterFunc returns the method "Ink.requestPresenter".
func (this Ink) RequestPresenterFunc() (fn js.Func[func(param InkPresenterParam) js.Promise[InkPresenter]]) {
	return fn.FromRef(
		bindings.InkRequestPresenterFunc(
			this.Ref(),
		),
	)
}

// RequestPresenter calls the method "Ink.requestPresenter".
func (this Ink) RequestPresenter(param InkPresenterParam) (ret js.Promise[InkPresenter]) {
	bindings.CallInkRequestPresenter(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&param),
	)

	return
}

// TryRequestPresenter calls the method "Ink.requestPresenter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Ink) TryRequestPresenter(param InkPresenterParam) (ret js.Promise[InkPresenter], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInkRequestPresenter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&param),
	)

	return
}

// HasRequestPresenter1 returns true if the method "Ink.requestPresenter" exists.
func (this Ink) HasRequestPresenter1() bool {
	return js.True == bindings.HasInkRequestPresenter1(
		this.Ref(),
	)
}

// RequestPresenter1Func returns the method "Ink.requestPresenter".
func (this Ink) RequestPresenter1Func() (fn js.Func[func() js.Promise[InkPresenter]]) {
	return fn.FromRef(
		bindings.InkRequestPresenter1Func(
			this.Ref(),
		),
	)
}

// RequestPresenter1 calls the method "Ink.requestPresenter".
func (this Ink) RequestPresenter1() (ret js.Promise[InkPresenter]) {
	bindings.CallInkRequestPresenter1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestPresenter1 calls the method "Ink.requestPresenter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Ink) TryRequestPresenter1() (ret js.Promise[InkPresenter], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInkRequestPresenter1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Id returns the value of property "PresentationConnection.id".
//
// It returns ok=false if there is no such property.
func (this PresentationConnection) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPresentationConnectionId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Url returns the value of property "PresentationConnection.url".
//
// It returns ok=false if there is no such property.
func (this PresentationConnection) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPresentationConnectionUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// State returns the value of property "PresentationConnection.state".
//
// It returns ok=false if there is no such property.
func (this PresentationConnection) State() (ret PresentationConnectionState, ok bool) {
	ok = js.True == bindings.GetPresentationConnectionState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BinaryType returns the value of property "PresentationConnection.binaryType".
//
// It returns ok=false if there is no such property.
func (this PresentationConnection) BinaryType() (ret BinaryType, ok bool) {
	ok = js.True == bindings.GetPresentationConnectionBinaryType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBinaryType sets the value of property "PresentationConnection.binaryType" to val.
//
// It returns false if the property cannot be set.
func (this PresentationConnection) SetBinaryType(val BinaryType) bool {
	return js.True == bindings.SetPresentationConnectionBinaryType(
		this.Ref(),
		uint32(val),
	)
}

// HasClose returns true if the method "PresentationConnection.close" exists.
func (this PresentationConnection) HasClose() bool {
	return js.True == bindings.HasPresentationConnectionClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "PresentationConnection.close".
func (this PresentationConnection) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PresentationConnectionCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "PresentationConnection.close".
func (this PresentationConnection) Close() (ret js.Void) {
	bindings.CallPresentationConnectionClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "PresentationConnection.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationConnection) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationConnectionClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasTerminate returns true if the method "PresentationConnection.terminate" exists.
func (this PresentationConnection) HasTerminate() bool {
	return js.True == bindings.HasPresentationConnectionTerminate(
		this.Ref(),
	)
}

// TerminateFunc returns the method "PresentationConnection.terminate".
func (this PresentationConnection) TerminateFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PresentationConnectionTerminateFunc(
			this.Ref(),
		),
	)
}

// Terminate calls the method "PresentationConnection.terminate".
func (this PresentationConnection) Terminate() (ret js.Void) {
	bindings.CallPresentationConnectionTerminate(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTerminate calls the method "PresentationConnection.terminate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationConnection) TryTerminate() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationConnectionTerminate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSend returns true if the method "PresentationConnection.send" exists.
func (this PresentationConnection) HasSend() bool {
	return js.True == bindings.HasPresentationConnectionSend(
		this.Ref(),
	)
}

// SendFunc returns the method "PresentationConnection.send".
func (this PresentationConnection) SendFunc() (fn js.Func[func(message js.String)]) {
	return fn.FromRef(
		bindings.PresentationConnectionSendFunc(
			this.Ref(),
		),
	)
}

// Send calls the method "PresentationConnection.send".
func (this PresentationConnection) Send(message js.String) (ret js.Void) {
	bindings.CallPresentationConnectionSend(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TrySend calls the method "PresentationConnection.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationConnection) TrySend(message js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationConnectionSend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasSend1 returns true if the method "PresentationConnection.send" exists.
func (this PresentationConnection) HasSend1() bool {
	return js.True == bindings.HasPresentationConnectionSend1(
		this.Ref(),
	)
}

// Send1Func returns the method "PresentationConnection.send".
func (this PresentationConnection) Send1Func() (fn js.Func[func(data Blob)]) {
	return fn.FromRef(
		bindings.PresentationConnectionSend1Func(
			this.Ref(),
		),
	)
}

// Send1 calls the method "PresentationConnection.send".
func (this PresentationConnection) Send1(data Blob) (ret js.Void) {
	bindings.CallPresentationConnectionSend1(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySend1 calls the method "PresentationConnection.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationConnection) TrySend1(data Blob) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationConnectionSend1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasSend2 returns true if the method "PresentationConnection.send" exists.
func (this PresentationConnection) HasSend2() bool {
	return js.True == bindings.HasPresentationConnectionSend2(
		this.Ref(),
	)
}

// Send2Func returns the method "PresentationConnection.send".
func (this PresentationConnection) Send2Func() (fn js.Func[func(data js.ArrayBuffer)]) {
	return fn.FromRef(
		bindings.PresentationConnectionSend2Func(
			this.Ref(),
		),
	)
}

// Send2 calls the method "PresentationConnection.send".
func (this PresentationConnection) Send2(data js.ArrayBuffer) (ret js.Void) {
	bindings.CallPresentationConnectionSend2(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySend2 calls the method "PresentationConnection.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationConnection) TrySend2(data js.ArrayBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationConnectionSend2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasSend3 returns true if the method "PresentationConnection.send" exists.
func (this PresentationConnection) HasSend3() bool {
	return js.True == bindings.HasPresentationConnectionSend3(
		this.Ref(),
	)
}

// Send3Func returns the method "PresentationConnection.send".
func (this PresentationConnection) Send3Func() (fn js.Func[func(data js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.PresentationConnectionSend3Func(
			this.Ref(),
		),
	)
}

// Send3 calls the method "PresentationConnection.send".
func (this PresentationConnection) Send3(data js.ArrayBufferView) (ret js.Void) {
	bindings.CallPresentationConnectionSend3(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySend3 calls the method "PresentationConnection.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationConnection) TrySend3(data js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationConnectionSend3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

type PresentationAvailability struct {
	EventTarget
}

func (this PresentationAvailability) Once() PresentationAvailability {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Value returns the value of property "PresentationAvailability.value".
//
// It returns ok=false if there is no such property.
func (this PresentationAvailability) Value() (ret bool, ok bool) {
	ok = js.True == bindings.GetPresentationAvailabilityValue(
		this.Ref(), js.Pointer(&ret),
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasStart returns true if the method "PresentationRequest.start" exists.
func (this PresentationRequest) HasStart() bool {
	return js.True == bindings.HasPresentationRequestStart(
		this.Ref(),
	)
}

// StartFunc returns the method "PresentationRequest.start".
func (this PresentationRequest) StartFunc() (fn js.Func[func() js.Promise[PresentationConnection]]) {
	return fn.FromRef(
		bindings.PresentationRequestStartFunc(
			this.Ref(),
		),
	)
}

// Start calls the method "PresentationRequest.start".
func (this PresentationRequest) Start() (ret js.Promise[PresentationConnection]) {
	bindings.CallPresentationRequestStart(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStart calls the method "PresentationRequest.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationRequest) TryStart() (ret js.Promise[PresentationConnection], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationRequestStart(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReconnect returns true if the method "PresentationRequest.reconnect" exists.
func (this PresentationRequest) HasReconnect() bool {
	return js.True == bindings.HasPresentationRequestReconnect(
		this.Ref(),
	)
}

// ReconnectFunc returns the method "PresentationRequest.reconnect".
func (this PresentationRequest) ReconnectFunc() (fn js.Func[func(presentationId js.String) js.Promise[PresentationConnection]]) {
	return fn.FromRef(
		bindings.PresentationRequestReconnectFunc(
			this.Ref(),
		),
	)
}

// Reconnect calls the method "PresentationRequest.reconnect".
func (this PresentationRequest) Reconnect(presentationId js.String) (ret js.Promise[PresentationConnection]) {
	bindings.CallPresentationRequestReconnect(
		this.Ref(), js.Pointer(&ret),
		presentationId.Ref(),
	)

	return
}

// TryReconnect calls the method "PresentationRequest.reconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationRequest) TryReconnect(presentationId js.String) (ret js.Promise[PresentationConnection], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationRequestReconnect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		presentationId.Ref(),
	)

	return
}

// HasGetAvailability returns true if the method "PresentationRequest.getAvailability" exists.
func (this PresentationRequest) HasGetAvailability() bool {
	return js.True == bindings.HasPresentationRequestGetAvailability(
		this.Ref(),
	)
}

// GetAvailabilityFunc returns the method "PresentationRequest.getAvailability".
func (this PresentationRequest) GetAvailabilityFunc() (fn js.Func[func() js.Promise[PresentationAvailability]]) {
	return fn.FromRef(
		bindings.PresentationRequestGetAvailabilityFunc(
			this.Ref(),
		),
	)
}

// GetAvailability calls the method "PresentationRequest.getAvailability".
func (this PresentationRequest) GetAvailability() (ret js.Promise[PresentationAvailability]) {
	bindings.CallPresentationRequestGetAvailability(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAvailability calls the method "PresentationRequest.getAvailability"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PresentationRequest) TryGetAvailability() (ret js.Promise[PresentationAvailability], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPresentationRequestGetAvailability(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PresentationConnectionList struct {
	EventTarget
}

func (this PresentationConnectionList) Once() PresentationConnectionList {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Connections returns the value of property "PresentationConnectionList.connections".
//
// It returns ok=false if there is no such property.
func (this PresentationConnectionList) Connections() (ret js.FrozenArray[PresentationConnection], ok bool) {
	ok = js.True == bindings.GetPresentationConnectionListConnections(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type PresentationReceiver struct {
	ref js.Ref
}

func (this PresentationReceiver) Once() PresentationReceiver {
	this.Ref().Once()
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
	this.Ref().Free()
}

// ConnectionList returns the value of property "PresentationReceiver.connectionList".
//
// It returns ok=false if there is no such property.
func (this PresentationReceiver) ConnectionList() (ret js.Promise[PresentationConnectionList], ok bool) {
	ok = js.True == bindings.GetPresentationReceiverConnectionList(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type Presentation struct {
	ref js.Ref
}

func (this Presentation) Once() Presentation {
	this.Ref().Once()
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
	this.Ref().Free()
}

// DefaultRequest returns the value of property "Presentation.defaultRequest".
//
// It returns ok=false if there is no such property.
func (this Presentation) DefaultRequest() (ret PresentationRequest, ok bool) {
	ok = js.True == bindings.GetPresentationDefaultRequest(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDefaultRequest sets the value of property "Presentation.defaultRequest" to val.
//
// It returns false if the property cannot be set.
func (this Presentation) SetDefaultRequest(val PresentationRequest) bool {
	return js.True == bindings.SetPresentationDefaultRequest(
		this.Ref(),
		val.Ref(),
	)
}

// Receiver returns the value of property "Presentation.receiver".
//
// It returns ok=false if there is no such property.
func (this Presentation) Receiver() (ret PresentationReceiver, ok bool) {
	ok = js.True == bindings.GetPresentationReceiver(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type VirtualKeyboard struct {
	EventTarget
}

func (this VirtualKeyboard) Once() VirtualKeyboard {
	this.Ref().Once()
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
	this.Ref().Free()
}

// BoundingRect returns the value of property "VirtualKeyboard.boundingRect".
//
// It returns ok=false if there is no such property.
func (this VirtualKeyboard) BoundingRect() (ret DOMRect, ok bool) {
	ok = js.True == bindings.GetVirtualKeyboardBoundingRect(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OverlaysContent returns the value of property "VirtualKeyboard.overlaysContent".
//
// It returns ok=false if there is no such property.
func (this VirtualKeyboard) OverlaysContent() (ret bool, ok bool) {
	ok = js.True == bindings.GetVirtualKeyboardOverlaysContent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetOverlaysContent sets the value of property "VirtualKeyboard.overlaysContent" to val.
//
// It returns false if the property cannot be set.
func (this VirtualKeyboard) SetOverlaysContent(val bool) bool {
	return js.True == bindings.SetVirtualKeyboardOverlaysContent(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// HasShow returns true if the method "VirtualKeyboard.show" exists.
func (this VirtualKeyboard) HasShow() bool {
	return js.True == bindings.HasVirtualKeyboardShow(
		this.Ref(),
	)
}

// ShowFunc returns the method "VirtualKeyboard.show".
func (this VirtualKeyboard) ShowFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.VirtualKeyboardShowFunc(
			this.Ref(),
		),
	)
}

// Show calls the method "VirtualKeyboard.show".
func (this VirtualKeyboard) Show() (ret js.Void) {
	bindings.CallVirtualKeyboardShow(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryShow calls the method "VirtualKeyboard.show"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VirtualKeyboard) TryShow() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVirtualKeyboardShow(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasHide returns true if the method "VirtualKeyboard.hide" exists.
func (this VirtualKeyboard) HasHide() bool {
	return js.True == bindings.HasVirtualKeyboardHide(
		this.Ref(),
	)
}

// HideFunc returns the method "VirtualKeyboard.hide".
func (this VirtualKeyboard) HideFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.VirtualKeyboardHideFunc(
			this.Ref(),
		),
	)
}

// Hide calls the method "VirtualKeyboard.hide".
func (this VirtualKeyboard) Hide() (ret js.Void) {
	bindings.CallVirtualKeyboardHide(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryHide calls the method "VirtualKeyboard.hide"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VirtualKeyboard) TryHide() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVirtualKeyboardHide(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
func (p NavigatorUABrandVersion) UpdateFrom(ref js.Ref) {
	bindings.NavigatorUABrandVersionJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NavigatorUABrandVersion) Update(ref js.Ref) {
	bindings.NavigatorUABrandVersionJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p UADataValues) UpdateFrom(ref js.Ref) {
	bindings.UADataValuesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p UADataValues) Update(ref js.Ref) {
	bindings.UADataValuesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p UALowEntropyJSON) UpdateFrom(ref js.Ref) {
	bindings.UALowEntropyJSONJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p UALowEntropyJSON) Update(ref js.Ref) {
	bindings.UALowEntropyJSONJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type NavigatorUAData struct {
	ref js.Ref
}

func (this NavigatorUAData) Once() NavigatorUAData {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Brands returns the value of property "NavigatorUAData.brands".
//
// It returns ok=false if there is no such property.
func (this NavigatorUAData) Brands() (ret js.FrozenArray[NavigatorUABrandVersion], ok bool) {
	ok = js.True == bindings.GetNavigatorUADataBrands(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Mobile returns the value of property "NavigatorUAData.mobile".
//
// It returns ok=false if there is no such property.
func (this NavigatorUAData) Mobile() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigatorUADataMobile(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Platform returns the value of property "NavigatorUAData.platform".
//
// It returns ok=false if there is no such property.
func (this NavigatorUAData) Platform() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorUADataPlatform(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetHighEntropyValues returns true if the method "NavigatorUAData.getHighEntropyValues" exists.
func (this NavigatorUAData) HasGetHighEntropyValues() bool {
	return js.True == bindings.HasNavigatorUADataGetHighEntropyValues(
		this.Ref(),
	)
}

// GetHighEntropyValuesFunc returns the method "NavigatorUAData.getHighEntropyValues".
func (this NavigatorUAData) GetHighEntropyValuesFunc() (fn js.Func[func(hints js.Array[js.String]) js.Promise[UADataValues]]) {
	return fn.FromRef(
		bindings.NavigatorUADataGetHighEntropyValuesFunc(
			this.Ref(),
		),
	)
}

// GetHighEntropyValues calls the method "NavigatorUAData.getHighEntropyValues".
func (this NavigatorUAData) GetHighEntropyValues(hints js.Array[js.String]) (ret js.Promise[UADataValues]) {
	bindings.CallNavigatorUADataGetHighEntropyValues(
		this.Ref(), js.Pointer(&ret),
		hints.Ref(),
	)

	return
}

// TryGetHighEntropyValues calls the method "NavigatorUAData.getHighEntropyValues"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NavigatorUAData) TryGetHighEntropyValues(hints js.Array[js.String]) (ret js.Promise[UADataValues], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorUADataGetHighEntropyValues(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		hints.Ref(),
	)

	return
}

// HasToJSON returns true if the method "NavigatorUAData.toJSON" exists.
func (this NavigatorUAData) HasToJSON() bool {
	return js.True == bindings.HasNavigatorUADataToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "NavigatorUAData.toJSON".
func (this NavigatorUAData) ToJSONFunc() (fn js.Func[func() UALowEntropyJSON]) {
	return fn.FromRef(
		bindings.NavigatorUADataToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "NavigatorUAData.toJSON".
func (this NavigatorUAData) ToJSON() (ret UALowEntropyJSON) {
	bindings.CallNavigatorUADataToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "NavigatorUAData.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NavigatorUAData) TryToJSON() (ret UALowEntropyJSON, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorUADataToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
func (p StorageEstimate) UpdateFrom(ref js.Ref) {
	bindings.StorageEstimateJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p StorageEstimate) Update(ref js.Ref) {
	bindings.StorageEstimateJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type StorageManager struct {
	ref js.Ref
}

func (this StorageManager) Once() StorageManager {
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasPersisted returns true if the method "StorageManager.persisted" exists.
func (this StorageManager) HasPersisted() bool {
	return js.True == bindings.HasStorageManagerPersisted(
		this.Ref(),
	)
}

// PersistedFunc returns the method "StorageManager.persisted".
func (this StorageManager) PersistedFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.StorageManagerPersistedFunc(
			this.Ref(),
		),
	)
}

// Persisted calls the method "StorageManager.persisted".
func (this StorageManager) Persisted() (ret js.Promise[js.Boolean]) {
	bindings.CallStorageManagerPersisted(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPersisted calls the method "StorageManager.persisted"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageManager) TryPersisted() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageManagerPersisted(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPersist returns true if the method "StorageManager.persist" exists.
func (this StorageManager) HasPersist() bool {
	return js.True == bindings.HasStorageManagerPersist(
		this.Ref(),
	)
}

// PersistFunc returns the method "StorageManager.persist".
func (this StorageManager) PersistFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.StorageManagerPersistFunc(
			this.Ref(),
		),
	)
}

// Persist calls the method "StorageManager.persist".
func (this StorageManager) Persist() (ret js.Promise[js.Boolean]) {
	bindings.CallStorageManagerPersist(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPersist calls the method "StorageManager.persist"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageManager) TryPersist() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageManagerPersist(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasEstimate returns true if the method "StorageManager.estimate" exists.
func (this StorageManager) HasEstimate() bool {
	return js.True == bindings.HasStorageManagerEstimate(
		this.Ref(),
	)
}

// EstimateFunc returns the method "StorageManager.estimate".
func (this StorageManager) EstimateFunc() (fn js.Func[func() js.Promise[StorageEstimate]]) {
	return fn.FromRef(
		bindings.StorageManagerEstimateFunc(
			this.Ref(),
		),
	)
}

// Estimate calls the method "StorageManager.estimate".
func (this StorageManager) Estimate() (ret js.Promise[StorageEstimate]) {
	bindings.CallStorageManagerEstimate(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryEstimate calls the method "StorageManager.estimate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageManager) TryEstimate() (ret js.Promise[StorageEstimate], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageManagerEstimate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetDirectory returns true if the method "StorageManager.getDirectory" exists.
func (this StorageManager) HasGetDirectory() bool {
	return js.True == bindings.HasStorageManagerGetDirectory(
		this.Ref(),
	)
}

// GetDirectoryFunc returns the method "StorageManager.getDirectory".
func (this StorageManager) GetDirectoryFunc() (fn js.Func[func() js.Promise[FileSystemDirectoryHandle]]) {
	return fn.FromRef(
		bindings.StorageManagerGetDirectoryFunc(
			this.Ref(),
		),
	)
}

// GetDirectory calls the method "StorageManager.getDirectory".
func (this StorageManager) GetDirectory() (ret js.Promise[FileSystemDirectoryHandle]) {
	bindings.CallStorageManagerGetDirectory(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetDirectory calls the method "StorageManager.getDirectory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageManager) TryGetDirectory() (ret js.Promise[FileSystemDirectoryHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageManagerGetDirectory(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
	this.Ref().Once()
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
	this.Ref().Free()
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
func (p IDBDatabaseInfo) UpdateFrom(ref js.Ref) {
	bindings.IDBDatabaseInfoJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IDBDatabaseInfo) Update(ref js.Ref) {
	bindings.IDBDatabaseInfoJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IDBFactory struct {
	ref js.Ref
}

func (this IDBFactory) Once() IDBFactory {
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasOpen returns true if the method "IDBFactory.open" exists.
func (this IDBFactory) HasOpen() bool {
	return js.True == bindings.HasIDBFactoryOpen(
		this.Ref(),
	)
}

// OpenFunc returns the method "IDBFactory.open".
func (this IDBFactory) OpenFunc() (fn js.Func[func(name js.String, version uint64) IDBOpenDBRequest]) {
	return fn.FromRef(
		bindings.IDBFactoryOpenFunc(
			this.Ref(),
		),
	)
}

// Open calls the method "IDBFactory.open".
func (this IDBFactory) Open(name js.String, version uint64) (ret IDBOpenDBRequest) {
	bindings.CallIDBFactoryOpen(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		float64(version),
	)

	return
}

// HasOpen1 returns true if the method "IDBFactory.open" exists.
func (this IDBFactory) HasOpen1() bool {
	return js.True == bindings.HasIDBFactoryOpen1(
		this.Ref(),
	)
}

// Open1Func returns the method "IDBFactory.open".
func (this IDBFactory) Open1Func() (fn js.Func[func(name js.String) IDBOpenDBRequest]) {
	return fn.FromRef(
		bindings.IDBFactoryOpen1Func(
			this.Ref(),
		),
	)
}

// Open1 calls the method "IDBFactory.open".
func (this IDBFactory) Open1(name js.String) (ret IDBOpenDBRequest) {
	bindings.CallIDBFactoryOpen1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryOpen1 calls the method "IDBFactory.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBFactory) TryOpen1(name js.String) (ret IDBOpenDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBFactoryOpen1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasDeleteDatabase returns true if the method "IDBFactory.deleteDatabase" exists.
func (this IDBFactory) HasDeleteDatabase() bool {
	return js.True == bindings.HasIDBFactoryDeleteDatabase(
		this.Ref(),
	)
}

// DeleteDatabaseFunc returns the method "IDBFactory.deleteDatabase".
func (this IDBFactory) DeleteDatabaseFunc() (fn js.Func[func(name js.String) IDBOpenDBRequest]) {
	return fn.FromRef(
		bindings.IDBFactoryDeleteDatabaseFunc(
			this.Ref(),
		),
	)
}

// DeleteDatabase calls the method "IDBFactory.deleteDatabase".
func (this IDBFactory) DeleteDatabase(name js.String) (ret IDBOpenDBRequest) {
	bindings.CallIDBFactoryDeleteDatabase(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryDeleteDatabase calls the method "IDBFactory.deleteDatabase"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBFactory) TryDeleteDatabase(name js.String) (ret IDBOpenDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBFactoryDeleteDatabase(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasDatabases returns true if the method "IDBFactory.databases" exists.
func (this IDBFactory) HasDatabases() bool {
	return js.True == bindings.HasIDBFactoryDatabases(
		this.Ref(),
	)
}

// DatabasesFunc returns the method "IDBFactory.databases".
func (this IDBFactory) DatabasesFunc() (fn js.Func[func() js.Promise[js.Array[IDBDatabaseInfo]]]) {
	return fn.FromRef(
		bindings.IDBFactoryDatabasesFunc(
			this.Ref(),
		),
	)
}

// Databases calls the method "IDBFactory.databases".
func (this IDBFactory) Databases() (ret js.Promise[js.Array[IDBDatabaseInfo]]) {
	bindings.CallIDBFactoryDatabases(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDatabases calls the method "IDBFactory.databases"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBFactory) TryDatabases() (ret js.Promise[js.Array[IDBDatabaseInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBFactoryDatabases(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCmp returns true if the method "IDBFactory.cmp" exists.
func (this IDBFactory) HasCmp() bool {
	return js.True == bindings.HasIDBFactoryCmp(
		this.Ref(),
	)
}

// CmpFunc returns the method "IDBFactory.cmp".
func (this IDBFactory) CmpFunc() (fn js.Func[func(first js.Any, second js.Any) int16]) {
	return fn.FromRef(
		bindings.IDBFactoryCmpFunc(
			this.Ref(),
		),
	)
}

// Cmp calls the method "IDBFactory.cmp".
func (this IDBFactory) Cmp(first js.Any, second js.Any) (ret int16) {
	bindings.CallIDBFactoryCmp(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		first.Ref(),
		second.Ref(),
	)

	return
}

type StorageBucket struct {
	ref js.Ref
}

func (this StorageBucket) Once() StorageBucket {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Name returns the value of property "StorageBucket.name".
//
// It returns ok=false if there is no such property.
func (this StorageBucket) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetStorageBucketName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IndexedDB returns the value of property "StorageBucket.indexedDB".
//
// It returns ok=false if there is no such property.
func (this StorageBucket) IndexedDB() (ret IDBFactory, ok bool) {
	ok = js.True == bindings.GetStorageBucketIndexedDB(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Caches returns the value of property "StorageBucket.caches".
//
// It returns ok=false if there is no such property.
func (this StorageBucket) Caches() (ret CacheStorage, ok bool) {
	ok = js.True == bindings.GetStorageBucketCaches(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasPersist returns true if the method "StorageBucket.persist" exists.
func (this StorageBucket) HasPersist() bool {
	return js.True == bindings.HasStorageBucketPersist(
		this.Ref(),
	)
}

// PersistFunc returns the method "StorageBucket.persist".
func (this StorageBucket) PersistFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.StorageBucketPersistFunc(
			this.Ref(),
		),
	)
}

// Persist calls the method "StorageBucket.persist".
func (this StorageBucket) Persist() (ret js.Promise[js.Boolean]) {
	bindings.CallStorageBucketPersist(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPersist calls the method "StorageBucket.persist"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucket) TryPersist() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketPersist(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPersisted returns true if the method "StorageBucket.persisted" exists.
func (this StorageBucket) HasPersisted() bool {
	return js.True == bindings.HasStorageBucketPersisted(
		this.Ref(),
	)
}

// PersistedFunc returns the method "StorageBucket.persisted".
func (this StorageBucket) PersistedFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.StorageBucketPersistedFunc(
			this.Ref(),
		),
	)
}

// Persisted calls the method "StorageBucket.persisted".
func (this StorageBucket) Persisted() (ret js.Promise[js.Boolean]) {
	bindings.CallStorageBucketPersisted(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPersisted calls the method "StorageBucket.persisted"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucket) TryPersisted() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketPersisted(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasEstimate returns true if the method "StorageBucket.estimate" exists.
func (this StorageBucket) HasEstimate() bool {
	return js.True == bindings.HasStorageBucketEstimate(
		this.Ref(),
	)
}

// EstimateFunc returns the method "StorageBucket.estimate".
func (this StorageBucket) EstimateFunc() (fn js.Func[func() js.Promise[StorageEstimate]]) {
	return fn.FromRef(
		bindings.StorageBucketEstimateFunc(
			this.Ref(),
		),
	)
}

// Estimate calls the method "StorageBucket.estimate".
func (this StorageBucket) Estimate() (ret js.Promise[StorageEstimate]) {
	bindings.CallStorageBucketEstimate(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryEstimate calls the method "StorageBucket.estimate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucket) TryEstimate() (ret js.Promise[StorageEstimate], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketEstimate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDurability returns true if the method "StorageBucket.durability" exists.
func (this StorageBucket) HasDurability() bool {
	return js.True == bindings.HasStorageBucketDurability(
		this.Ref(),
	)
}

// DurabilityFunc returns the method "StorageBucket.durability".
func (this StorageBucket) DurabilityFunc() (fn js.Func[func() js.Promise[StorageBucketDurability]]) {
	return fn.FromRef(
		bindings.StorageBucketDurabilityFunc(
			this.Ref(),
		),
	)
}

// Durability calls the method "StorageBucket.durability".
func (this StorageBucket) Durability() (ret js.Promise[StorageBucketDurability]) {
	bindings.CallStorageBucketDurability(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDurability calls the method "StorageBucket.durability"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucket) TryDurability() (ret js.Promise[StorageBucketDurability], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketDurability(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetExpires returns true if the method "StorageBucket.setExpires" exists.
func (this StorageBucket) HasSetExpires() bool {
	return js.True == bindings.HasStorageBucketSetExpires(
		this.Ref(),
	)
}

// SetExpiresFunc returns the method "StorageBucket.setExpires".
func (this StorageBucket) SetExpiresFunc() (fn js.Func[func(expires DOMHighResTimeStamp) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.StorageBucketSetExpiresFunc(
			this.Ref(),
		),
	)
}

// SetExpires calls the method "StorageBucket.setExpires".
func (this StorageBucket) SetExpires(expires DOMHighResTimeStamp) (ret js.Promise[js.Void]) {
	bindings.CallStorageBucketSetExpires(
		this.Ref(), js.Pointer(&ret),
		float64(expires),
	)

	return
}

// TrySetExpires calls the method "StorageBucket.setExpires"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucket) TrySetExpires(expires DOMHighResTimeStamp) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketSetExpires(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(expires),
	)

	return
}

// HasExpires returns true if the method "StorageBucket.expires" exists.
func (this StorageBucket) HasExpires() bool {
	return js.True == bindings.HasStorageBucketExpires(
		this.Ref(),
	)
}

// ExpiresFunc returns the method "StorageBucket.expires".
func (this StorageBucket) ExpiresFunc() (fn js.Func[func() js.Promise[js.Number[DOMHighResTimeStamp]]]) {
	return fn.FromRef(
		bindings.StorageBucketExpiresFunc(
			this.Ref(),
		),
	)
}

// Expires calls the method "StorageBucket.expires".
func (this StorageBucket) Expires() (ret js.Promise[js.Number[DOMHighResTimeStamp]]) {
	bindings.CallStorageBucketExpires(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryExpires calls the method "StorageBucket.expires"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucket) TryExpires() (ret js.Promise[js.Number[DOMHighResTimeStamp]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketExpires(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetDirectory returns true if the method "StorageBucket.getDirectory" exists.
func (this StorageBucket) HasGetDirectory() bool {
	return js.True == bindings.HasStorageBucketGetDirectory(
		this.Ref(),
	)
}

// GetDirectoryFunc returns the method "StorageBucket.getDirectory".
func (this StorageBucket) GetDirectoryFunc() (fn js.Func[func() js.Promise[FileSystemDirectoryHandle]]) {
	return fn.FromRef(
		bindings.StorageBucketGetDirectoryFunc(
			this.Ref(),
		),
	)
}

// GetDirectory calls the method "StorageBucket.getDirectory".
func (this StorageBucket) GetDirectory() (ret js.Promise[FileSystemDirectoryHandle]) {
	bindings.CallStorageBucketGetDirectory(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetDirectory calls the method "StorageBucket.getDirectory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucket) TryGetDirectory() (ret js.Promise[FileSystemDirectoryHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketGetDirectory(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
func (p StorageBucketOptions) UpdateFrom(ref js.Ref) {
	bindings.StorageBucketOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p StorageBucketOptions) Update(ref js.Ref) {
	bindings.StorageBucketOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type StorageBucketManager struct {
	ref js.Ref
}

func (this StorageBucketManager) Once() StorageBucketManager {
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasOpen returns true if the method "StorageBucketManager.open" exists.
func (this StorageBucketManager) HasOpen() bool {
	return js.True == bindings.HasStorageBucketManagerOpen(
		this.Ref(),
	)
}

// OpenFunc returns the method "StorageBucketManager.open".
func (this StorageBucketManager) OpenFunc() (fn js.Func[func(name js.String, options StorageBucketOptions) js.Promise[StorageBucket]]) {
	return fn.FromRef(
		bindings.StorageBucketManagerOpenFunc(
			this.Ref(),
		),
	)
}

// Open calls the method "StorageBucketManager.open".
func (this StorageBucketManager) Open(name js.String, options StorageBucketOptions) (ret js.Promise[StorageBucket]) {
	bindings.CallStorageBucketManagerOpen(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasOpen1 returns true if the method "StorageBucketManager.open" exists.
func (this StorageBucketManager) HasOpen1() bool {
	return js.True == bindings.HasStorageBucketManagerOpen1(
		this.Ref(),
	)
}

// Open1Func returns the method "StorageBucketManager.open".
func (this StorageBucketManager) Open1Func() (fn js.Func[func(name js.String) js.Promise[StorageBucket]]) {
	return fn.FromRef(
		bindings.StorageBucketManagerOpen1Func(
			this.Ref(),
		),
	)
}

// Open1 calls the method "StorageBucketManager.open".
func (this StorageBucketManager) Open1(name js.String) (ret js.Promise[StorageBucket]) {
	bindings.CallStorageBucketManagerOpen1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryOpen1 calls the method "StorageBucketManager.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucketManager) TryOpen1(name js.String) (ret js.Promise[StorageBucket], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketManagerOpen1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasKeys returns true if the method "StorageBucketManager.keys" exists.
func (this StorageBucketManager) HasKeys() bool {
	return js.True == bindings.HasStorageBucketManagerKeys(
		this.Ref(),
	)
}

// KeysFunc returns the method "StorageBucketManager.keys".
func (this StorageBucketManager) KeysFunc() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	return fn.FromRef(
		bindings.StorageBucketManagerKeysFunc(
			this.Ref(),
		),
	)
}

// Keys calls the method "StorageBucketManager.keys".
func (this StorageBucketManager) Keys() (ret js.Promise[js.Array[js.String]]) {
	bindings.CallStorageBucketManagerKeys(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryKeys calls the method "StorageBucketManager.keys"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucketManager) TryKeys() (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketManagerKeys(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDelete returns true if the method "StorageBucketManager.delete" exists.
func (this StorageBucketManager) HasDelete() bool {
	return js.True == bindings.HasStorageBucketManagerDelete(
		this.Ref(),
	)
}

// DeleteFunc returns the method "StorageBucketManager.delete".
func (this StorageBucketManager) DeleteFunc() (fn js.Func[func(name js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.StorageBucketManagerDeleteFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "StorageBucketManager.delete".
func (this StorageBucketManager) Delete(name js.String) (ret js.Promise[js.Void]) {
	bindings.CallStorageBucketManagerDelete(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryDelete calls the method "StorageBucketManager.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageBucketManager) TryDelete(name js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageBucketManagerDelete(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
		assert.Throw("invalid", "callback", "invocation")
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Name returns the value of property "Lock.name".
//
// It returns ok=false if there is no such property.
func (this Lock) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLockName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Mode returns the value of property "Lock.mode".
//
// It returns ok=false if there is no such property.
func (this Lock) Mode() (ret LockMode, ok bool) {
	ok = js.True == bindings.GetLockMode(
		this.Ref(), js.Pointer(&ret),
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
func (p LockOptions) UpdateFrom(ref js.Ref) {
	bindings.LockOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p LockOptions) Update(ref js.Ref) {
	bindings.LockOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p LockInfo) UpdateFrom(ref js.Ref) {
	bindings.LockInfoJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p LockInfo) Update(ref js.Ref) {
	bindings.LockInfoJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p LockManagerSnapshot) UpdateFrom(ref js.Ref) {
	bindings.LockManagerSnapshotJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p LockManagerSnapshot) Update(ref js.Ref) {
	bindings.LockManagerSnapshotJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type LockManager struct {
	ref js.Ref
}

func (this LockManager) Once() LockManager {
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasRequest returns true if the method "LockManager.request" exists.
func (this LockManager) HasRequest() bool {
	return js.True == bindings.HasLockManagerRequest(
		this.Ref(),
	)
}

// RequestFunc returns the method "LockManager.request".
func (this LockManager) RequestFunc() (fn js.Func[func(name js.String, callback js.Func[func(lock Lock) js.Promise[js.Any]]) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.LockManagerRequestFunc(
			this.Ref(),
		),
	)
}

// Request calls the method "LockManager.request".
func (this LockManager) Request(name js.String, callback js.Func[func(lock Lock) js.Promise[js.Any]]) (ret js.Promise[js.Any]) {
	bindings.CallLockManagerRequest(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		callback.Ref(),
	)

	return
}

// HasRequest1 returns true if the method "LockManager.request" exists.
func (this LockManager) HasRequest1() bool {
	return js.True == bindings.HasLockManagerRequest1(
		this.Ref(),
	)
}

// Request1Func returns the method "LockManager.request".
func (this LockManager) Request1Func() (fn js.Func[func(name js.String, options LockOptions, callback js.Func[func(lock Lock) js.Promise[js.Any]]) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.LockManagerRequest1Func(
			this.Ref(),
		),
	)
}

// Request1 calls the method "LockManager.request".
func (this LockManager) Request1(name js.String, options LockOptions, callback js.Func[func(lock Lock) js.Promise[js.Any]]) (ret js.Promise[js.Any]) {
	bindings.CallLockManagerRequest1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// HasQuery returns true if the method "LockManager.query" exists.
func (this LockManager) HasQuery() bool {
	return js.True == bindings.HasLockManagerQuery(
		this.Ref(),
	)
}

// QueryFunc returns the method "LockManager.query".
func (this LockManager) QueryFunc() (fn js.Func[func() js.Promise[LockManagerSnapshot]]) {
	return fn.FromRef(
		bindings.LockManagerQueryFunc(
			this.Ref(),
		),
	)
}

// Query calls the method "LockManager.query".
func (this LockManager) Query() (ret js.Promise[LockManagerSnapshot]) {
	bindings.CallLockManagerQuery(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryQuery calls the method "LockManager.query"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this LockManager) TryQuery() (ret js.Promise[LockManagerSnapshot], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLockManagerQuery(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type MimeType struct {
	ref js.Ref
}

func (this MimeType) Once() MimeType {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Type returns the value of property "MimeType.type".
//
// It returns ok=false if there is no such property.
func (this MimeType) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMimeTypeType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Description returns the value of property "MimeType.description".
//
// It returns ok=false if there is no such property.
func (this MimeType) Description() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMimeTypeDescription(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Suffixes returns the value of property "MimeType.suffixes".
//
// It returns ok=false if there is no such property.
func (this MimeType) Suffixes() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMimeTypeSuffixes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EnabledPlugin returns the value of property "MimeType.enabledPlugin".
//
// It returns ok=false if there is no such property.
func (this MimeType) EnabledPlugin() (ret Plugin, ok bool) {
	ok = js.True == bindings.GetMimeTypeEnabledPlugin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type Plugin struct {
	ref js.Ref
}

func (this Plugin) Once() Plugin {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Name returns the value of property "Plugin.name".
//
// It returns ok=false if there is no such property.
func (this Plugin) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPluginName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Description returns the value of property "Plugin.description".
//
// It returns ok=false if there is no such property.
func (this Plugin) Description() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPluginDescription(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Filename returns the value of property "Plugin.filename".
//
// It returns ok=false if there is no such property.
func (this Plugin) Filename() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPluginFilename(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Length returns the value of property "Plugin.length".
//
// It returns ok=false if there is no such property.
func (this Plugin) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetPluginLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasItem returns true if the method "Plugin.item" exists.
func (this Plugin) HasItem() bool {
	return js.True == bindings.HasPluginItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "Plugin.item".
func (this Plugin) ItemFunc() (fn js.Func[func(index uint32) MimeType]) {
	return fn.FromRef(
		bindings.PluginItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "Plugin.item".
func (this Plugin) Item(index uint32) (ret MimeType) {
	bindings.CallPluginItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "Plugin.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Plugin) TryItem(index uint32) (ret MimeType, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPluginItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasNamedItem returns true if the method "Plugin.namedItem" exists.
func (this Plugin) HasNamedItem() bool {
	return js.True == bindings.HasPluginNamedItem(
		this.Ref(),
	)
}

// NamedItemFunc returns the method "Plugin.namedItem".
func (this Plugin) NamedItemFunc() (fn js.Func[func(name js.String) MimeType]) {
	return fn.FromRef(
		bindings.PluginNamedItemFunc(
			this.Ref(),
		),
	)
}

// NamedItem calls the method "Plugin.namedItem".
func (this Plugin) NamedItem(name js.String) (ret MimeType) {
	bindings.CallPluginNamedItem(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryNamedItem calls the method "Plugin.namedItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Plugin) TryNamedItem(name js.String) (ret MimeType, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPluginNamedItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

type PluginArray struct {
	ref js.Ref
}

func (this PluginArray) Once() PluginArray {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Length returns the value of property "PluginArray.length".
//
// It returns ok=false if there is no such property.
func (this PluginArray) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetPluginArrayLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRefresh returns true if the method "PluginArray.refresh" exists.
func (this PluginArray) HasRefresh() bool {
	return js.True == bindings.HasPluginArrayRefresh(
		this.Ref(),
	)
}

// RefreshFunc returns the method "PluginArray.refresh".
func (this PluginArray) RefreshFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PluginArrayRefreshFunc(
			this.Ref(),
		),
	)
}

// Refresh calls the method "PluginArray.refresh".
func (this PluginArray) Refresh() (ret js.Void) {
	bindings.CallPluginArrayRefresh(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRefresh calls the method "PluginArray.refresh"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PluginArray) TryRefresh() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPluginArrayRefresh(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasItem returns true if the method "PluginArray.item" exists.
func (this PluginArray) HasItem() bool {
	return js.True == bindings.HasPluginArrayItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "PluginArray.item".
func (this PluginArray) ItemFunc() (fn js.Func[func(index uint32) Plugin]) {
	return fn.FromRef(
		bindings.PluginArrayItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "PluginArray.item".
func (this PluginArray) Item(index uint32) (ret Plugin) {
	bindings.CallPluginArrayItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "PluginArray.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PluginArray) TryItem(index uint32) (ret Plugin, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPluginArrayItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasNamedItem returns true if the method "PluginArray.namedItem" exists.
func (this PluginArray) HasNamedItem() bool {
	return js.True == bindings.HasPluginArrayNamedItem(
		this.Ref(),
	)
}

// NamedItemFunc returns the method "PluginArray.namedItem".
func (this PluginArray) NamedItemFunc() (fn js.Func[func(name js.String) Plugin]) {
	return fn.FromRef(
		bindings.PluginArrayNamedItemFunc(
			this.Ref(),
		),
	)
}

// NamedItem calls the method "PluginArray.namedItem".
func (this PluginArray) NamedItem(name js.String) (ret Plugin) {
	bindings.CallPluginArrayNamedItem(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryNamedItem calls the method "PluginArray.namedItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PluginArray) TryNamedItem(name js.String) (ret Plugin, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPluginArrayNamedItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

type MimeTypeArray struct {
	ref js.Ref
}

func (this MimeTypeArray) Once() MimeTypeArray {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Length returns the value of property "MimeTypeArray.length".
//
// It returns ok=false if there is no such property.
func (this MimeTypeArray) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetMimeTypeArrayLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasItem returns true if the method "MimeTypeArray.item" exists.
func (this MimeTypeArray) HasItem() bool {
	return js.True == bindings.HasMimeTypeArrayItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "MimeTypeArray.item".
func (this MimeTypeArray) ItemFunc() (fn js.Func[func(index uint32) MimeType]) {
	return fn.FromRef(
		bindings.MimeTypeArrayItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "MimeTypeArray.item".
func (this MimeTypeArray) Item(index uint32) (ret MimeType) {
	bindings.CallMimeTypeArrayItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "MimeTypeArray.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MimeTypeArray) TryItem(index uint32) (ret MimeType, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMimeTypeArrayItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasNamedItem returns true if the method "MimeTypeArray.namedItem" exists.
func (this MimeTypeArray) HasNamedItem() bool {
	return js.True == bindings.HasMimeTypeArrayNamedItem(
		this.Ref(),
	)
}

// NamedItemFunc returns the method "MimeTypeArray.namedItem".
func (this MimeTypeArray) NamedItemFunc() (fn js.Func[func(name js.String) MimeType]) {
	return fn.FromRef(
		bindings.MimeTypeArrayNamedItemFunc(
			this.Ref(),
		),
	)
}

// NamedItem calls the method "MimeTypeArray.namedItem".
func (this MimeTypeArray) NamedItem(name js.String) (ret MimeType) {
	bindings.CallMimeTypeArrayNamedItem(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryNamedItem calls the method "MimeTypeArray.namedItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MimeTypeArray) TryNamedItem(name js.String) (ret MimeType, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMimeTypeArrayNamedItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
func (p MLComputeResult) UpdateFrom(ref js.Ref) {
	bindings.MLComputeResultJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLComputeResult) Update(ref js.Ref) {
	bindings.MLComputeResultJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MLGraph struct {
	ref js.Ref
}

func (this MLGraph) Once() MLGraph {
	this.Ref().Once()
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
	this.Ref().Free()
}
