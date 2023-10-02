// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web/bindings"
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
	IncludeContinuous bool

	FFI_USE_IncludeContinuous bool // for IncludeContinuous.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IsInputPendingOptions with all fields set.
func (p IsInputPendingOptions) FromRef(ref js.Ref) IsInputPendingOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IsInputPendingOptions IsInputPendingOptions [// IsInputPendingOptions] [0x140005bf9a0 0x140005bfa40] 0x1400081ecc0 {0 0}} in the application heap.
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

// IsInputPending calls the method "Scheduling.isInputPending".
//
// The returned bool will be false if there is no such method.
func (this Scheduling) IsInputPending(isInputPendingOptions IsInputPendingOptions) (bool, bool) {
	var _ok bool
	_ret := bindings.CallSchedulingIsInputPending(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&isInputPendingOptions),
	)

	return _ret == js.True, _ok
}

// IsInputPendingFunc returns the method "Scheduling.isInputPending".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Scheduling) IsInputPendingFunc() (fn js.Func[func(isInputPendingOptions IsInputPendingOptions) bool]) {
	return fn.FromRef(
		bindings.SchedulingIsInputPendingFunc(
			this.Ref(),
		),
	)
}

// IsInputPending1 calls the method "Scheduling.isInputPending".
//
// The returned bool will be false if there is no such method.
func (this Scheduling) IsInputPending1() (bool, bool) {
	var _ok bool
	_ret := bindings.CallSchedulingIsInputPending1(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// IsInputPending1Func returns the method "Scheduling.isInputPending".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Scheduling) IsInputPending1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.SchedulingIsInputPending1Func(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this WakeLockSentinel) Released() (bool, bool) {
	var _ok bool
	_ret := bindings.GetWakeLockSentinelReleased(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Type returns the value of property "WakeLockSentinel.type".
//
// The returned bool will be false if there is no such property.
func (this WakeLockSentinel) Type() (WakeLockType, bool) {
	var _ok bool
	_ret := bindings.GetWakeLockSentinelType(
		this.Ref(), js.Pointer(&_ok),
	)
	return WakeLockType(_ret), _ok
}

// Release calls the method "WakeLockSentinel.release".
//
// The returned bool will be false if there is no such method.
func (this WakeLockSentinel) Release() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWakeLockSentinelRelease(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ReleaseFunc returns the method "WakeLockSentinel.release".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WakeLockSentinel) ReleaseFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WakeLockSentinelReleaseFunc(
			this.Ref(),
		),
	)
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

// Request calls the method "WakeLock.request".
//
// The returned bool will be false if there is no such method.
func (this WakeLock) Request(typ WakeLockType) (js.Promise[WakeLockSentinel], bool) {
	var _ok bool
	_ret := bindings.CallWakeLockRequest(
		this.Ref(), js.Pointer(&_ok),
		uint32(typ),
	)

	return js.Promise[WakeLockSentinel]{}.FromRef(_ret), _ok
}

// RequestFunc returns the method "WakeLock.request".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WakeLock) RequestFunc() (fn js.Func[func(typ WakeLockType) js.Promise[WakeLockSentinel]]) {
	return fn.FromRef(
		bindings.WakeLockRequestFunc(
			this.Ref(),
		),
	)
}

// Request1 calls the method "WakeLock.request".
//
// The returned bool will be false if there is no such method.
func (this WakeLock) Request1() (js.Promise[WakeLockSentinel], bool) {
	var _ok bool
	_ret := bindings.CallWakeLockRequest1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[WakeLockSentinel]{}.FromRef(_ret), _ok
}

// Request1Func returns the method "WakeLock.request".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WakeLock) Request1Func() (fn js.Func[func() js.Promise[WakeLockSentinel]]) {
	return fn.FromRef(
		bindings.WakeLockRequest1Func(
			this.Ref(),
		),
	)
}

type PointerEventInit struct {
	// PointerId is "PointerEventInit.pointerId"
	//
	// Optional, defaults to 0.
	PointerId int32
	// Width is "PointerEventInit.width"
	//
	// Optional, defaults to 1.
	Width float64
	// Height is "PointerEventInit.height"
	//
	// Optional, defaults to 1.
	Height float64
	// Pressure is "PointerEventInit.pressure"
	//
	// Optional, defaults to 0.
	Pressure float32
	// TangentialPressure is "PointerEventInit.tangentialPressure"
	//
	// Optional, defaults to 0.
	TangentialPressure float32
	// TiltX is "PointerEventInit.tiltX"
	//
	// Optional
	TiltX int32
	// TiltY is "PointerEventInit.tiltY"
	//
	// Optional
	TiltY int32
	// Twist is "PointerEventInit.twist"
	//
	// Optional, defaults to 0.
	Twist int32
	// AltitudeAngle is "PointerEventInit.altitudeAngle"
	//
	// Optional
	AltitudeAngle float64
	// AzimuthAngle is "PointerEventInit.azimuthAngle"
	//
	// Optional
	AzimuthAngle float64
	// PointerType is "PointerEventInit.pointerType"
	//
	// Optional, defaults to "".
	PointerType js.String
	// IsPrimary is "PointerEventInit.isPrimary"
	//
	// Optional, defaults to false.
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
	MovementX float64
	// MovementY is "PointerEventInit.movementY"
	//
	// Optional, defaults to 0.
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

// New creates a new {0x140004cc0e0 PointerEventInit PointerEventInit [// PointerEventInit] [0x140005bfea0 0x140005e1220 0x1400069c8c0 0x1400069ca00 0x1400069d0e0 0x140007c60a0 0x140007c61e0 0x140007c6320 0x140007c6460 0x140007c65a0 0x140007c66e0 0x140007c6780 0x140007c68c0 0x140007c6960 0x140007c6a00 0x140007c6b40 0x140005e00a0 0x140005f1900 0x1400069c960 0x1400069d040 0x140007c6000 0x140007c6140 0x140007c6280 0x140007c63c0 0x140007c6500 0x140007c6640 0x140007c6820 0x140007c6aa0 0x140007c6be0] 0x1400081ed08 {0 0}} in the application heap.
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

func NewPointerEvent(typ js.String, eventInitDict PointerEventInit) PointerEvent {
	return PointerEvent{}.FromRef(
		bindings.NewPointerEventByPointerEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewPointerEventByPointerEvent1(typ js.String) PointerEvent {
	return PointerEvent{}.FromRef(
		bindings.NewPointerEventByPointerEvent1(
			typ.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this PointerEvent) PointerId() (int32, bool) {
	var _ok bool
	_ret := bindings.GetPointerEventPointerId(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Width returns the value of property "PointerEvent.width".
//
// The returned bool will be false if there is no such property.
func (this PointerEvent) Width() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPointerEventWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Height returns the value of property "PointerEvent.height".
//
// The returned bool will be false if there is no such property.
func (this PointerEvent) Height() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPointerEventHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Pressure returns the value of property "PointerEvent.pressure".
//
// The returned bool will be false if there is no such property.
func (this PointerEvent) Pressure() (float32, bool) {
	var _ok bool
	_ret := bindings.GetPointerEventPressure(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// TangentialPressure returns the value of property "PointerEvent.tangentialPressure".
//
// The returned bool will be false if there is no such property.
func (this PointerEvent) TangentialPressure() (float32, bool) {
	var _ok bool
	_ret := bindings.GetPointerEventTangentialPressure(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// TiltX returns the value of property "PointerEvent.tiltX".
//
// The returned bool will be false if there is no such property.
func (this PointerEvent) TiltX() (int32, bool) {
	var _ok bool
	_ret := bindings.GetPointerEventTiltX(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// TiltY returns the value of property "PointerEvent.tiltY".
//
// The returned bool will be false if there is no such property.
func (this PointerEvent) TiltY() (int32, bool) {
	var _ok bool
	_ret := bindings.GetPointerEventTiltY(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Twist returns the value of property "PointerEvent.twist".
//
// The returned bool will be false if there is no such property.
func (this PointerEvent) Twist() (int32, bool) {
	var _ok bool
	_ret := bindings.GetPointerEventTwist(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// AltitudeAngle returns the value of property "PointerEvent.altitudeAngle".
//
// The returned bool will be false if there is no such property.
func (this PointerEvent) AltitudeAngle() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPointerEventAltitudeAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// AzimuthAngle returns the value of property "PointerEvent.azimuthAngle".
//
// The returned bool will be false if there is no such property.
func (this PointerEvent) AzimuthAngle() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPointerEventAzimuthAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// PointerType returns the value of property "PointerEvent.pointerType".
//
// The returned bool will be false if there is no such property.
func (this PointerEvent) PointerType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPointerEventPointerType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// IsPrimary returns the value of property "PointerEvent.isPrimary".
//
// The returned bool will be false if there is no such property.
func (this PointerEvent) IsPrimary() (bool, bool) {
	var _ok bool
	_ret := bindings.GetPointerEventIsPrimary(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// GetCoalescedEvents calls the method "PointerEvent.getCoalescedEvents".
//
// The returned bool will be false if there is no such method.
func (this PointerEvent) GetCoalescedEvents() (js.Array[PointerEvent], bool) {
	var _ok bool
	_ret := bindings.CallPointerEventGetCoalescedEvents(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[PointerEvent]{}.FromRef(_ret), _ok
}

// GetCoalescedEventsFunc returns the method "PointerEvent.getCoalescedEvents".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PointerEvent) GetCoalescedEventsFunc() (fn js.Func[func() js.Array[PointerEvent]]) {
	return fn.FromRef(
		bindings.PointerEventGetCoalescedEventsFunc(
			this.Ref(),
		),
	)
}

// GetPredictedEvents calls the method "PointerEvent.getPredictedEvents".
//
// The returned bool will be false if there is no such method.
func (this PointerEvent) GetPredictedEvents() (js.Array[PointerEvent], bool) {
	var _ok bool
	_ret := bindings.CallPointerEventGetPredictedEvents(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[PointerEvent]{}.FromRef(_ret), _ok
}

// GetPredictedEventsFunc returns the method "PointerEvent.getPredictedEvents".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PointerEvent) GetPredictedEventsFunc() (fn js.Func[func() js.Array[PointerEvent]]) {
	return fn.FromRef(
		bindings.PointerEventGetPredictedEventsFunc(
			this.Ref(),
		),
	)
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

// New creates a new {0x140004cc0e0 InkTrailStyle InkTrailStyle [// InkTrailStyle] [0x140007c6d20 0x140007c6dc0] 0x1400081edb0 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this InkPresenter) PresentationArea() (Element, bool) {
	var _ok bool
	_ret := bindings.GetInkPresenterPresentationArea(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// ExpectedImprovement returns the value of property "InkPresenter.expectedImprovement".
//
// The returned bool will be false if there is no such property.
func (this InkPresenter) ExpectedImprovement() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetInkPresenterExpectedImprovement(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// UpdateInkTrailStartPoint calls the method "InkPresenter.updateInkTrailStartPoint".
//
// The returned bool will be false if there is no such method.
func (this InkPresenter) UpdateInkTrailStartPoint(event PointerEvent, style InkTrailStyle) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallInkPresenterUpdateInkTrailStartPoint(
		this.Ref(), js.Pointer(&_ok),
		event.Ref(),
		js.Pointer(&style),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UpdateInkTrailStartPointFunc returns the method "InkPresenter.updateInkTrailStartPoint".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this InkPresenter) UpdateInkTrailStartPointFunc() (fn js.Func[func(event PointerEvent, style InkTrailStyle)]) {
	return fn.FromRef(
		bindings.InkPresenterUpdateInkTrailStartPointFunc(
			this.Ref(),
		),
	)
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

// New creates a new {0x140004cc0e0 InkPresenterParam InkPresenterParam [// InkPresenterParam] [0x140007c6f00] 0x1400081edf8 {0 0}} in the application heap.
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

// RequestPresenter calls the method "Ink.requestPresenter".
//
// The returned bool will be false if there is no such method.
func (this Ink) RequestPresenter(param InkPresenterParam) (js.Promise[InkPresenter], bool) {
	var _ok bool
	_ret := bindings.CallInkRequestPresenter(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&param),
	)

	return js.Promise[InkPresenter]{}.FromRef(_ret), _ok
}

// RequestPresenterFunc returns the method "Ink.requestPresenter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Ink) RequestPresenterFunc() (fn js.Func[func(param InkPresenterParam) js.Promise[InkPresenter]]) {
	return fn.FromRef(
		bindings.InkRequestPresenterFunc(
			this.Ref(),
		),
	)
}

// RequestPresenter1 calls the method "Ink.requestPresenter".
//
// The returned bool will be false if there is no such method.
func (this Ink) RequestPresenter1() (js.Promise[InkPresenter], bool) {
	var _ok bool
	_ret := bindings.CallInkRequestPresenter1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[InkPresenter]{}.FromRef(_ret), _ok
}

// RequestPresenter1Func returns the method "Ink.requestPresenter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Ink) RequestPresenter1Func() (fn js.Func[func() js.Promise[InkPresenter]]) {
	return fn.FromRef(
		bindings.InkRequestPresenter1Func(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this PresentationConnection) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPresentationConnectionId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Url returns the value of property "PresentationConnection.url".
//
// The returned bool will be false if there is no such property.
func (this PresentationConnection) Url() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPresentationConnectionUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// State returns the value of property "PresentationConnection.state".
//
// The returned bool will be false if there is no such property.
func (this PresentationConnection) State() (PresentationConnectionState, bool) {
	var _ok bool
	_ret := bindings.GetPresentationConnectionState(
		this.Ref(), js.Pointer(&_ok),
	)
	return PresentationConnectionState(_ret), _ok
}

// BinaryType returns the value of property "PresentationConnection.binaryType".
//
// The returned bool will be false if there is no such property.
func (this PresentationConnection) BinaryType() (BinaryType, bool) {
	var _ok bool
	_ret := bindings.GetPresentationConnectionBinaryType(
		this.Ref(), js.Pointer(&_ok),
	)
	return BinaryType(_ret), _ok
}

// BinaryType sets the value of property "PresentationConnection.binaryType" to val.
//
// It returns false if the property cannot be set.
func (this PresentationConnection) SetBinaryType(val BinaryType) bool {
	return js.True == bindings.SetPresentationConnectionBinaryType(
		this.Ref(),
		uint32(val),
	)
}

// Close calls the method "PresentationConnection.close".
//
// The returned bool will be false if there is no such method.
func (this PresentationConnection) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPresentationConnectionClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "PresentationConnection.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PresentationConnection) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PresentationConnectionCloseFunc(
			this.Ref(),
		),
	)
}

// Terminate calls the method "PresentationConnection.terminate".
//
// The returned bool will be false if there is no such method.
func (this PresentationConnection) Terminate() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPresentationConnectionTerminate(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TerminateFunc returns the method "PresentationConnection.terminate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PresentationConnection) TerminateFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PresentationConnectionTerminateFunc(
			this.Ref(),
		),
	)
}

// Send calls the method "PresentationConnection.send".
//
// The returned bool will be false if there is no such method.
func (this PresentationConnection) Send(message js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPresentationConnectionSend(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SendFunc returns the method "PresentationConnection.send".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PresentationConnection) SendFunc() (fn js.Func[func(message js.String)]) {
	return fn.FromRef(
		bindings.PresentationConnectionSendFunc(
			this.Ref(),
		),
	)
}

// Send1 calls the method "PresentationConnection.send".
//
// The returned bool will be false if there is no such method.
func (this PresentationConnection) Send1(data Blob) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPresentationConnectionSend1(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Send1Func returns the method "PresentationConnection.send".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PresentationConnection) Send1Func() (fn js.Func[func(data Blob)]) {
	return fn.FromRef(
		bindings.PresentationConnectionSend1Func(
			this.Ref(),
		),
	)
}

// Send2 calls the method "PresentationConnection.send".
//
// The returned bool will be false if there is no such method.
func (this PresentationConnection) Send2(data js.ArrayBuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPresentationConnectionSend2(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Send2Func returns the method "PresentationConnection.send".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PresentationConnection) Send2Func() (fn js.Func[func(data js.ArrayBuffer)]) {
	return fn.FromRef(
		bindings.PresentationConnectionSend2Func(
			this.Ref(),
		),
	)
}

// Send3 calls the method "PresentationConnection.send".
//
// The returned bool will be false if there is no such method.
func (this PresentationConnection) Send3(data ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPresentationConnectionSend3(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Send3Func returns the method "PresentationConnection.send".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PresentationConnection) Send3Func() (fn js.Func[func(data ArrayBufferView)]) {
	return fn.FromRef(
		bindings.PresentationConnectionSend3Func(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this PresentationAvailability) Value() (bool, bool) {
	var _ok bool
	_ret := bindings.GetPresentationAvailabilityValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

func NewPresentationRequest(url js.String) PresentationRequest {
	return PresentationRequest{}.FromRef(
		bindings.NewPresentationRequestByPresentationRequest(
			url.Ref()),
	)
}

func NewPresentationRequestByPresentationRequest1(urls js.Array[js.String]) PresentationRequest {
	return PresentationRequest{}.FromRef(
		bindings.NewPresentationRequestByPresentationRequest1(
			urls.Ref()),
	)
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

// Start calls the method "PresentationRequest.start".
//
// The returned bool will be false if there is no such method.
func (this PresentationRequest) Start() (js.Promise[PresentationConnection], bool) {
	var _ok bool
	_ret := bindings.CallPresentationRequestStart(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[PresentationConnection]{}.FromRef(_ret), _ok
}

// StartFunc returns the method "PresentationRequest.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PresentationRequest) StartFunc() (fn js.Func[func() js.Promise[PresentationConnection]]) {
	return fn.FromRef(
		bindings.PresentationRequestStartFunc(
			this.Ref(),
		),
	)
}

// Reconnect calls the method "PresentationRequest.reconnect".
//
// The returned bool will be false if there is no such method.
func (this PresentationRequest) Reconnect(presentationId js.String) (js.Promise[PresentationConnection], bool) {
	var _ok bool
	_ret := bindings.CallPresentationRequestReconnect(
		this.Ref(), js.Pointer(&_ok),
		presentationId.Ref(),
	)

	return js.Promise[PresentationConnection]{}.FromRef(_ret), _ok
}

// ReconnectFunc returns the method "PresentationRequest.reconnect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PresentationRequest) ReconnectFunc() (fn js.Func[func(presentationId js.String) js.Promise[PresentationConnection]]) {
	return fn.FromRef(
		bindings.PresentationRequestReconnectFunc(
			this.Ref(),
		),
	)
}

// GetAvailability calls the method "PresentationRequest.getAvailability".
//
// The returned bool will be false if there is no such method.
func (this PresentationRequest) GetAvailability() (js.Promise[PresentationAvailability], bool) {
	var _ok bool
	_ret := bindings.CallPresentationRequestGetAvailability(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[PresentationAvailability]{}.FromRef(_ret), _ok
}

// GetAvailabilityFunc returns the method "PresentationRequest.getAvailability".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PresentationRequest) GetAvailabilityFunc() (fn js.Func[func() js.Promise[PresentationAvailability]]) {
	return fn.FromRef(
		bindings.PresentationRequestGetAvailabilityFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this PresentationConnectionList) Connections() (js.FrozenArray[PresentationConnection], bool) {
	var _ok bool
	_ret := bindings.GetPresentationConnectionListConnections(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[PresentationConnection]{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this PresentationReceiver) ConnectionList() (js.Promise[PresentationConnectionList], bool) {
	var _ok bool
	_ret := bindings.GetPresentationReceiverConnectionList(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[PresentationConnectionList]{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this Presentation) DefaultRequest() (PresentationRequest, bool) {
	var _ok bool
	_ret := bindings.GetPresentationDefaultRequest(
		this.Ref(), js.Pointer(&_ok),
	)
	return PresentationRequest{}.FromRef(_ret), _ok
}

// DefaultRequest sets the value of property "Presentation.defaultRequest" to val.
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
// The returned bool will be false if there is no such property.
func (this Presentation) Receiver() (PresentationReceiver, bool) {
	var _ok bool
	_ret := bindings.GetPresentationReceiver(
		this.Ref(), js.Pointer(&_ok),
	)
	return PresentationReceiver{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this VirtualKeyboard) BoundingRect() (DOMRect, bool) {
	var _ok bool
	_ret := bindings.GetVirtualKeyboardBoundingRect(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMRect{}.FromRef(_ret), _ok
}

// OverlaysContent returns the value of property "VirtualKeyboard.overlaysContent".
//
// The returned bool will be false if there is no such property.
func (this VirtualKeyboard) OverlaysContent() (bool, bool) {
	var _ok bool
	_ret := bindings.GetVirtualKeyboardOverlaysContent(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// OverlaysContent sets the value of property "VirtualKeyboard.overlaysContent" to val.
//
// It returns false if the property cannot be set.
func (this VirtualKeyboard) SetOverlaysContent(val bool) bool {
	return js.True == bindings.SetVirtualKeyboardOverlaysContent(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Show calls the method "VirtualKeyboard.show".
//
// The returned bool will be false if there is no such method.
func (this VirtualKeyboard) Show() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallVirtualKeyboardShow(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ShowFunc returns the method "VirtualKeyboard.show".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VirtualKeyboard) ShowFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.VirtualKeyboardShowFunc(
			this.Ref(),
		),
	)
}

// Hide calls the method "VirtualKeyboard.hide".
//
// The returned bool will be false if there is no such method.
func (this VirtualKeyboard) Hide() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallVirtualKeyboardHide(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// HideFunc returns the method "VirtualKeyboard.hide".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VirtualKeyboard) HideFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.VirtualKeyboardHideFunc(
			this.Ref(),
		),
	)
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

// New creates a new {0x140004cc0e0 NavigatorUABrandVersion NavigatorUABrandVersion [// NavigatorUABrandVersion] [0x140007c72c0 0x140007c7360] 0x1400081eee8 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 UADataValues UADataValues [// UADataValues] [0x140007c7180 0x140007c7220 0x140007c7400 0x140007c74a0 0x140007c7540 0x140007c75e0 0x140007c7680 0x140007c77c0 0x140007c7860 0x140007c7900 0x140007c79a0 0x140007c7720 0x140007c7a40] 0x1400081eed0 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 UALowEntropyJSON UALowEntropyJSON [// UALowEntropyJSON] [0x140007c7ae0 0x140007c7b80 0x140007c7cc0 0x140007c7c20] 0x1400081ef18 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this NavigatorUAData) Brands() (js.FrozenArray[NavigatorUABrandVersion], bool) {
	var _ok bool
	_ret := bindings.GetNavigatorUADataBrands(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[NavigatorUABrandVersion]{}.FromRef(_ret), _ok
}

// Mobile returns the value of property "NavigatorUAData.mobile".
//
// The returned bool will be false if there is no such property.
func (this NavigatorUAData) Mobile() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorUADataMobile(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Platform returns the value of property "NavigatorUAData.platform".
//
// The returned bool will be false if there is no such property.
func (this NavigatorUAData) Platform() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorUADataPlatform(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// GetHighEntropyValues calls the method "NavigatorUAData.getHighEntropyValues".
//
// The returned bool will be false if there is no such method.
func (this NavigatorUAData) GetHighEntropyValues(hints js.Array[js.String]) (js.Promise[UADataValues], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorUADataGetHighEntropyValues(
		this.Ref(), js.Pointer(&_ok),
		hints.Ref(),
	)

	return js.Promise[UADataValues]{}.FromRef(_ret), _ok
}

// GetHighEntropyValuesFunc returns the method "NavigatorUAData.getHighEntropyValues".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NavigatorUAData) GetHighEntropyValuesFunc() (fn js.Func[func(hints js.Array[js.String]) js.Promise[UADataValues]]) {
	return fn.FromRef(
		bindings.NavigatorUADataGetHighEntropyValuesFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "NavigatorUAData.toJSON".
//
// The returned bool will be false if there is no such method.
func (this NavigatorUAData) ToJSON() (UALowEntropyJSON, bool) {
	var _ret UALowEntropyJSON
	_ok := js.True == bindings.CallNavigatorUADataToJSON(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// ToJSONFunc returns the method "NavigatorUAData.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NavigatorUAData) ToJSONFunc() (fn js.Func[func() UALowEntropyJSON]) {
	return fn.FromRef(
		bindings.NavigatorUADataToJSONFunc(
			this.Ref(),
		),
	)
}

type StorageEstimate struct {
	// Usage is "StorageEstimate.usage"
	//
	// Optional
	Usage uint64
	// Quota is "StorageEstimate.quota"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 StorageEstimate StorageEstimate [// StorageEstimate] [0x140007c7e00 0x140007c7f40 0x140007c7ea0 0x140007d2000] 0x1400081ef48 {0 0}} in the application heap.
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

// Persisted calls the method "StorageManager.persisted".
//
// The returned bool will be false if there is no such method.
func (this StorageManager) Persisted() (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallStorageManagerPersisted(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// PersistedFunc returns the method "StorageManager.persisted".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageManager) PersistedFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.StorageManagerPersistedFunc(
			this.Ref(),
		),
	)
}

// Persist calls the method "StorageManager.persist".
//
// The returned bool will be false if there is no such method.
func (this StorageManager) Persist() (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallStorageManagerPersist(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// PersistFunc returns the method "StorageManager.persist".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageManager) PersistFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.StorageManagerPersistFunc(
			this.Ref(),
		),
	)
}

// Estimate calls the method "StorageManager.estimate".
//
// The returned bool will be false if there is no such method.
func (this StorageManager) Estimate() (js.Promise[StorageEstimate], bool) {
	var _ok bool
	_ret := bindings.CallStorageManagerEstimate(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[StorageEstimate]{}.FromRef(_ret), _ok
}

// EstimateFunc returns the method "StorageManager.estimate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageManager) EstimateFunc() (fn js.Func[func() js.Promise[StorageEstimate]]) {
	return fn.FromRef(
		bindings.StorageManagerEstimateFunc(
			this.Ref(),
		),
	)
}

// GetDirectory calls the method "StorageManager.getDirectory".
//
// The returned bool will be false if there is no such method.
func (this StorageManager) GetDirectory() (js.Promise[FileSystemDirectoryHandle], bool) {
	var _ok bool
	_ret := bindings.CallStorageManagerGetDirectory(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[FileSystemDirectoryHandle]{}.FromRef(_ret), _ok
}

// GetDirectoryFunc returns the method "StorageManager.getDirectory".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageManager) GetDirectoryFunc() (fn js.Func[func() js.Promise[FileSystemDirectoryHandle]]) {
	return fn.FromRef(
		bindings.StorageManagerGetDirectoryFunc(
			this.Ref(),
		),
	)
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
	Version uint64

	FFI_USE_Version bool // for Version.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IDBDatabaseInfo with all fields set.
func (p IDBDatabaseInfo) FromRef(ref js.Ref) IDBDatabaseInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IDBDatabaseInfo IDBDatabaseInfo [// IDBDatabaseInfo] [0x140007d20a0 0x140007d2140 0x140007d21e0] 0x1400081efa8 {0 0}} in the application heap.
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

// Open calls the method "IDBFactory.open".
//
// The returned bool will be false if there is no such method.
func (this IDBFactory) Open(name js.String, version uint64) (IDBOpenDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBFactoryOpen(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		float64(version),
	)

	return IDBOpenDBRequest{}.FromRef(_ret), _ok
}

// OpenFunc returns the method "IDBFactory.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBFactory) OpenFunc() (fn js.Func[func(name js.String, version uint64) IDBOpenDBRequest]) {
	return fn.FromRef(
		bindings.IDBFactoryOpenFunc(
			this.Ref(),
		),
	)
}

// Open1 calls the method "IDBFactory.open".
//
// The returned bool will be false if there is no such method.
func (this IDBFactory) Open1(name js.String) (IDBOpenDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBFactoryOpen1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return IDBOpenDBRequest{}.FromRef(_ret), _ok
}

// Open1Func returns the method "IDBFactory.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBFactory) Open1Func() (fn js.Func[func(name js.String) IDBOpenDBRequest]) {
	return fn.FromRef(
		bindings.IDBFactoryOpen1Func(
			this.Ref(),
		),
	)
}

// DeleteDatabase calls the method "IDBFactory.deleteDatabase".
//
// The returned bool will be false if there is no such method.
func (this IDBFactory) DeleteDatabase(name js.String) (IDBOpenDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBFactoryDeleteDatabase(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return IDBOpenDBRequest{}.FromRef(_ret), _ok
}

// DeleteDatabaseFunc returns the method "IDBFactory.deleteDatabase".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBFactory) DeleteDatabaseFunc() (fn js.Func[func(name js.String) IDBOpenDBRequest]) {
	return fn.FromRef(
		bindings.IDBFactoryDeleteDatabaseFunc(
			this.Ref(),
		),
	)
}

// Databases calls the method "IDBFactory.databases".
//
// The returned bool will be false if there is no such method.
func (this IDBFactory) Databases() (js.Promise[js.Array[IDBDatabaseInfo]], bool) {
	var _ok bool
	_ret := bindings.CallIDBFactoryDatabases(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[IDBDatabaseInfo]]{}.FromRef(_ret), _ok
}

// DatabasesFunc returns the method "IDBFactory.databases".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBFactory) DatabasesFunc() (fn js.Func[func() js.Promise[js.Array[IDBDatabaseInfo]]]) {
	return fn.FromRef(
		bindings.IDBFactoryDatabasesFunc(
			this.Ref(),
		),
	)
}

// Cmp calls the method "IDBFactory.cmp".
//
// The returned bool will be false if there is no such method.
func (this IDBFactory) Cmp(first js.Any, second js.Any) (int16, bool) {
	var _ok bool
	_ret := bindings.CallIDBFactoryCmp(
		this.Ref(), js.Pointer(&_ok),
		first.Ref(),
		second.Ref(),
	)

	return int16(_ret), _ok
}

// CmpFunc returns the method "IDBFactory.cmp".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBFactory) CmpFunc() (fn js.Func[func(first js.Any, second js.Any) int16]) {
	return fn.FromRef(
		bindings.IDBFactoryCmpFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this StorageBucket) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetStorageBucketName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// IndexedDB returns the value of property "StorageBucket.indexedDB".
//
// The returned bool will be false if there is no such property.
func (this StorageBucket) IndexedDB() (IDBFactory, bool) {
	var _ok bool
	_ret := bindings.GetStorageBucketIndexedDB(
		this.Ref(), js.Pointer(&_ok),
	)
	return IDBFactory{}.FromRef(_ret), _ok
}

// Caches returns the value of property "StorageBucket.caches".
//
// The returned bool will be false if there is no such property.
func (this StorageBucket) Caches() (CacheStorage, bool) {
	var _ok bool
	_ret := bindings.GetStorageBucketCaches(
		this.Ref(), js.Pointer(&_ok),
	)
	return CacheStorage{}.FromRef(_ret), _ok
}

// Persist calls the method "StorageBucket.persist".
//
// The returned bool will be false if there is no such method.
func (this StorageBucket) Persist() (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallStorageBucketPersist(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// PersistFunc returns the method "StorageBucket.persist".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageBucket) PersistFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.StorageBucketPersistFunc(
			this.Ref(),
		),
	)
}

// Persisted calls the method "StorageBucket.persisted".
//
// The returned bool will be false if there is no such method.
func (this StorageBucket) Persisted() (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallStorageBucketPersisted(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// PersistedFunc returns the method "StorageBucket.persisted".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageBucket) PersistedFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.StorageBucketPersistedFunc(
			this.Ref(),
		),
	)
}

// Estimate calls the method "StorageBucket.estimate".
//
// The returned bool will be false if there is no such method.
func (this StorageBucket) Estimate() (js.Promise[StorageEstimate], bool) {
	var _ok bool
	_ret := bindings.CallStorageBucketEstimate(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[StorageEstimate]{}.FromRef(_ret), _ok
}

// EstimateFunc returns the method "StorageBucket.estimate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageBucket) EstimateFunc() (fn js.Func[func() js.Promise[StorageEstimate]]) {
	return fn.FromRef(
		bindings.StorageBucketEstimateFunc(
			this.Ref(),
		),
	)
}

// Durability calls the method "StorageBucket.durability".
//
// The returned bool will be false if there is no such method.
func (this StorageBucket) Durability() (js.Promise[StorageBucketDurability], bool) {
	var _ok bool
	_ret := bindings.CallStorageBucketDurability(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[StorageBucketDurability]{}.FromRef(_ret), _ok
}

// DurabilityFunc returns the method "StorageBucket.durability".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageBucket) DurabilityFunc() (fn js.Func[func() js.Promise[StorageBucketDurability]]) {
	return fn.FromRef(
		bindings.StorageBucketDurabilityFunc(
			this.Ref(),
		),
	)
}

// SetExpires calls the method "StorageBucket.setExpires".
//
// The returned bool will be false if there is no such method.
func (this StorageBucket) SetExpires(expires DOMHighResTimeStamp) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallStorageBucketSetExpires(
		this.Ref(), js.Pointer(&_ok),
		float64(expires),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetExpiresFunc returns the method "StorageBucket.setExpires".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageBucket) SetExpiresFunc() (fn js.Func[func(expires DOMHighResTimeStamp) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.StorageBucketSetExpiresFunc(
			this.Ref(),
		),
	)
}

// Expires calls the method "StorageBucket.expires".
//
// The returned bool will be false if there is no such method.
func (this StorageBucket) Expires() (js.Promise[js.Number[DOMHighResTimeStamp]], bool) {
	var _ok bool
	_ret := bindings.CallStorageBucketExpires(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Number[DOMHighResTimeStamp]]{}.FromRef(_ret), _ok
}

// ExpiresFunc returns the method "StorageBucket.expires".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageBucket) ExpiresFunc() (fn js.Func[func() js.Promise[js.Number[DOMHighResTimeStamp]]]) {
	return fn.FromRef(
		bindings.StorageBucketExpiresFunc(
			this.Ref(),
		),
	)
}

// GetDirectory calls the method "StorageBucket.getDirectory".
//
// The returned bool will be false if there is no such method.
func (this StorageBucket) GetDirectory() (js.Promise[FileSystemDirectoryHandle], bool) {
	var _ok bool
	_ret := bindings.CallStorageBucketGetDirectory(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[FileSystemDirectoryHandle]{}.FromRef(_ret), _ok
}

// GetDirectoryFunc returns the method "StorageBucket.getDirectory".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageBucket) GetDirectoryFunc() (fn js.Func[func() js.Promise[FileSystemDirectoryHandle]]) {
	return fn.FromRef(
		bindings.StorageBucketGetDirectoryFunc(
			this.Ref(),
		),
	)
}

type StorageBucketOptions struct {
	// Persisted is "StorageBucketOptions.persisted"
	//
	// Optional, defaults to null.
	Persisted bool
	// Durability is "StorageBucketOptions.durability"
	//
	// Optional, defaults to null.
	Durability StorageBucketDurability
	// Quota is "StorageBucketOptions.quota"
	//
	// Optional, defaults to null.
	Quota uint64
	// Expires is "StorageBucketOptions.expires"
	//
	// Optional, defaults to null.
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

// New creates a new {0x140004cc0e0 StorageBucketOptions StorageBucketOptions [// StorageBucketOptions] [0x140007d2320 0x140007d2460 0x140007d2500 0x140007d2640 0x140007d23c0 0x140007d25a0 0x140007d26e0] 0x1400081f020 {0 0}} in the application heap.
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

// Open calls the method "StorageBucketManager.open".
//
// The returned bool will be false if there is no such method.
func (this StorageBucketManager) Open(name js.String, options StorageBucketOptions) (js.Promise[StorageBucket], bool) {
	var _ok bool
	_ret := bindings.CallStorageBucketManagerOpen(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[StorageBucket]{}.FromRef(_ret), _ok
}

// OpenFunc returns the method "StorageBucketManager.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageBucketManager) OpenFunc() (fn js.Func[func(name js.String, options StorageBucketOptions) js.Promise[StorageBucket]]) {
	return fn.FromRef(
		bindings.StorageBucketManagerOpenFunc(
			this.Ref(),
		),
	)
}

// Open1 calls the method "StorageBucketManager.open".
//
// The returned bool will be false if there is no such method.
func (this StorageBucketManager) Open1(name js.String) (js.Promise[StorageBucket], bool) {
	var _ok bool
	_ret := bindings.CallStorageBucketManagerOpen1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.Promise[StorageBucket]{}.FromRef(_ret), _ok
}

// Open1Func returns the method "StorageBucketManager.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageBucketManager) Open1Func() (fn js.Func[func(name js.String) js.Promise[StorageBucket]]) {
	return fn.FromRef(
		bindings.StorageBucketManagerOpen1Func(
			this.Ref(),
		),
	)
}

// Keys calls the method "StorageBucketManager.keys".
//
// The returned bool will be false if there is no such method.
func (this StorageBucketManager) Keys() (js.Promise[js.Array[js.String]], bool) {
	var _ok bool
	_ret := bindings.CallStorageBucketManagerKeys(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[js.String]]{}.FromRef(_ret), _ok
}

// KeysFunc returns the method "StorageBucketManager.keys".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageBucketManager) KeysFunc() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	return fn.FromRef(
		bindings.StorageBucketManagerKeysFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "StorageBucketManager.delete".
//
// The returned bool will be false if there is no such method.
func (this StorageBucketManager) Delete(name js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallStorageBucketManagerDelete(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// DeleteFunc returns the method "StorageBucketManager.delete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageBucketManager) DeleteFunc() (fn js.Func[func(name js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.StorageBucketManagerDeleteFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this Lock) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetLockName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Mode returns the value of property "Lock.mode".
//
// The returned bool will be false if there is no such property.
func (this Lock) Mode() (LockMode, bool) {
	var _ok bool
	_ret := bindings.GetLockMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return LockMode(_ret), _ok
}

type LockOptions struct {
	// Mode is "LockOptions.mode"
	//
	// Optional, defaults to "exclusive".
	Mode LockMode
	// IfAvailable is "LockOptions.ifAvailable"
	//
	// Optional, defaults to false.
	IfAvailable bool
	// Steal is "LockOptions.steal"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 LockOptions LockOptions [// LockOptions] [0x140007d2820 0x140007d28c0 0x140007d2a00 0x140007d2b40 0x140007d2960 0x140007d2aa0] 0x1400081f068 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 LockInfo LockInfo [// LockInfo] [0x140007d2be0 0x140007d2c80 0x140007d2d20] 0x1400081f0f8 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 LockManagerSnapshot LockManagerSnapshot [// LockManagerSnapshot] [0x140007d2dc0 0x140007d2e60] 0x1400081f0c8 {0 0}} in the application heap.
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

// Request calls the method "LockManager.request".
//
// The returned bool will be false if there is no such method.
func (this LockManager) Request(name js.String, callback js.Func[func(lock Lock) js.Promise[js.Any]]) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallLockManagerRequest(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		callback.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// RequestFunc returns the method "LockManager.request".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this LockManager) RequestFunc() (fn js.Func[func(name js.String, callback js.Func[func(lock Lock) js.Promise[js.Any]]) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.LockManagerRequestFunc(
			this.Ref(),
		),
	)
}

// Request1 calls the method "LockManager.request".
//
// The returned bool will be false if there is no such method.
func (this LockManager) Request1(name js.String, options LockOptions, callback js.Func[func(lock Lock) js.Promise[js.Any]]) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallLockManagerRequest1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		js.Pointer(&options),
		callback.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// Request1Func returns the method "LockManager.request".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this LockManager) Request1Func() (fn js.Func[func(name js.String, options LockOptions, callback js.Func[func(lock Lock) js.Promise[js.Any]]) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.LockManagerRequest1Func(
			this.Ref(),
		),
	)
}

// Query calls the method "LockManager.query".
//
// The returned bool will be false if there is no such method.
func (this LockManager) Query() (js.Promise[LockManagerSnapshot], bool) {
	var _ok bool
	_ret := bindings.CallLockManagerQuery(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[LockManagerSnapshot]{}.FromRef(_ret), _ok
}

// QueryFunc returns the method "LockManager.query".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this LockManager) QueryFunc() (fn js.Func[func() js.Promise[LockManagerSnapshot]]) {
	return fn.FromRef(
		bindings.LockManagerQueryFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this MimeType) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMimeTypeType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Description returns the value of property "MimeType.description".
//
// The returned bool will be false if there is no such property.
func (this MimeType) Description() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMimeTypeDescription(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Suffixes returns the value of property "MimeType.suffixes".
//
// The returned bool will be false if there is no such property.
func (this MimeType) Suffixes() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMimeTypeSuffixes(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// EnabledPlugin returns the value of property "MimeType.enabledPlugin".
//
// The returned bool will be false if there is no such property.
func (this MimeType) EnabledPlugin() (Plugin, bool) {
	var _ok bool
	_ret := bindings.GetMimeTypeEnabledPlugin(
		this.Ref(), js.Pointer(&_ok),
	)
	return Plugin{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this Plugin) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPluginName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Description returns the value of property "Plugin.description".
//
// The returned bool will be false if there is no such property.
func (this Plugin) Description() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPluginDescription(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Filename returns the value of property "Plugin.filename".
//
// The returned bool will be false if there is no such property.
func (this Plugin) Filename() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPluginFilename(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Length returns the value of property "Plugin.length".
//
// The returned bool will be false if there is no such property.
func (this Plugin) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetPluginLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Item calls the method "Plugin.item".
//
// The returned bool will be false if there is no such method.
func (this Plugin) Item(index uint32) (MimeType, bool) {
	var _ok bool
	_ret := bindings.CallPluginItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return MimeType{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "Plugin.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Plugin) ItemFunc() (fn js.Func[func(index uint32) MimeType]) {
	return fn.FromRef(
		bindings.PluginItemFunc(
			this.Ref(),
		),
	)
}

// NamedItem calls the method "Plugin.namedItem".
//
// The returned bool will be false if there is no such method.
func (this Plugin) NamedItem(name js.String) (MimeType, bool) {
	var _ok bool
	_ret := bindings.CallPluginNamedItem(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return MimeType{}.FromRef(_ret), _ok
}

// NamedItemFunc returns the method "Plugin.namedItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Plugin) NamedItemFunc() (fn js.Func[func(name js.String) MimeType]) {
	return fn.FromRef(
		bindings.PluginNamedItemFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this PluginArray) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetPluginArrayLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Refresh calls the method "PluginArray.refresh".
//
// The returned bool will be false if there is no such method.
func (this PluginArray) Refresh() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPluginArrayRefresh(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RefreshFunc returns the method "PluginArray.refresh".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PluginArray) RefreshFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PluginArrayRefreshFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "PluginArray.item".
//
// The returned bool will be false if there is no such method.
func (this PluginArray) Item(index uint32) (Plugin, bool) {
	var _ok bool
	_ret := bindings.CallPluginArrayItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return Plugin{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "PluginArray.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PluginArray) ItemFunc() (fn js.Func[func(index uint32) Plugin]) {
	return fn.FromRef(
		bindings.PluginArrayItemFunc(
			this.Ref(),
		),
	)
}

// NamedItem calls the method "PluginArray.namedItem".
//
// The returned bool will be false if there is no such method.
func (this PluginArray) NamedItem(name js.String) (Plugin, bool) {
	var _ok bool
	_ret := bindings.CallPluginArrayNamedItem(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return Plugin{}.FromRef(_ret), _ok
}

// NamedItemFunc returns the method "PluginArray.namedItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PluginArray) NamedItemFunc() (fn js.Func[func(name js.String) Plugin]) {
	return fn.FromRef(
		bindings.PluginArrayNamedItemFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this MimeTypeArray) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetMimeTypeArrayLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Item calls the method "MimeTypeArray.item".
//
// The returned bool will be false if there is no such method.
func (this MimeTypeArray) Item(index uint32) (MimeType, bool) {
	var _ok bool
	_ret := bindings.CallMimeTypeArrayItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return MimeType{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "MimeTypeArray.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MimeTypeArray) ItemFunc() (fn js.Func[func(index uint32) MimeType]) {
	return fn.FromRef(
		bindings.MimeTypeArrayItemFunc(
			this.Ref(),
		),
	)
}

// NamedItem calls the method "MimeTypeArray.namedItem".
//
// The returned bool will be false if there is no such method.
func (this MimeTypeArray) NamedItem(name js.String) (MimeType, bool) {
	var _ok bool
	_ret := bindings.CallMimeTypeArrayNamedItem(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return MimeType{}.FromRef(_ret), _ok
}

// NamedItemFunc returns the method "MimeTypeArray.namedItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MimeTypeArray) NamedItemFunc() (fn js.Func[func(name js.String) MimeType]) {
	return fn.FromRef(
		bindings.MimeTypeArrayNamedItemFunc(
			this.Ref(),
		),
	)
}

type MLNamedArrayBufferViews = js.Record[ArrayBufferView]

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

// New creates a new {0x140004cc0e0 MLComputeResult MLComputeResult [// MLComputeResult] [0x140007d3040 0x140007d30e0] 0x1400081f1a0 {0 0}} in the application heap.
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
