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

type LargestContentfulPaint struct {
	PerformanceEntry
}

func (this LargestContentfulPaint) Once() LargestContentfulPaint {
	this.Ref().Once()
	return this
}

func (this LargestContentfulPaint) Ref() js.Ref {
	return this.PerformanceEntry.Ref()
}

func (this LargestContentfulPaint) FromRef(ref js.Ref) LargestContentfulPaint {
	this.PerformanceEntry = this.PerformanceEntry.FromRef(ref)
	return this
}

func (this LargestContentfulPaint) Free() {
	this.Ref().Free()
}

// RenderTime returns the value of property "LargestContentfulPaint.renderTime".
//
// The returned bool will be false if there is no such property.
func (this LargestContentfulPaint) RenderTime() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetLargestContentfulPaintRenderTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// LoadTime returns the value of property "LargestContentfulPaint.loadTime".
//
// The returned bool will be false if there is no such property.
func (this LargestContentfulPaint) LoadTime() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetLargestContentfulPaintLoadTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// Size returns the value of property "LargestContentfulPaint.size".
//
// The returned bool will be false if there is no such property.
func (this LargestContentfulPaint) Size() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetLargestContentfulPaintSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Id returns the value of property "LargestContentfulPaint.id".
//
// The returned bool will be false if there is no such property.
func (this LargestContentfulPaint) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetLargestContentfulPaintId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Url returns the value of property "LargestContentfulPaint.url".
//
// The returned bool will be false if there is no such property.
func (this LargestContentfulPaint) Url() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetLargestContentfulPaintUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Element returns the value of property "LargestContentfulPaint.element".
//
// The returned bool will be false if there is no such property.
func (this LargestContentfulPaint) Element() (Element, bool) {
	var _ok bool
	_ret := bindings.GetLargestContentfulPaintElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// ToJSON calls the method "LargestContentfulPaint.toJSON".
//
// The returned bool will be false if there is no such method.
func (this LargestContentfulPaint) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallLargestContentfulPaintToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "LargestContentfulPaint.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this LargestContentfulPaint) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.LargestContentfulPaintToJSONFunc(
			this.Ref(),
		),
	)
}

type LatencyMode uint32

const (
	_ LatencyMode = iota

	LatencyMode_QUALITY
	LatencyMode_REALTIME
)

func (LatencyMode) FromRef(str js.Ref) LatencyMode {
	return LatencyMode(bindings.ConstOfLatencyMode(str))
}

func (x LatencyMode) String() (string, bool) {
	switch x {
	case LatencyMode_QUALITY:
		return "quality", true
	case LatencyMode_REALTIME:
		return "realtime", true
	default:
		return "", false
	}
}

type LayoutConstraints struct {
	ref js.Ref
}

func (this LayoutConstraints) Once() LayoutConstraints {
	this.Ref().Once()
	return this
}

func (this LayoutConstraints) Ref() js.Ref {
	return this.ref
}

func (this LayoutConstraints) FromRef(ref js.Ref) LayoutConstraints {
	this.ref = ref
	return this
}

func (this LayoutConstraints) Free() {
	this.Ref().Free()
}

// AvailableInlineSize returns the value of property "LayoutConstraints.availableInlineSize".
//
// The returned bool will be false if there is no such property.
func (this LayoutConstraints) AvailableInlineSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutConstraintsAvailableInlineSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// AvailableBlockSize returns the value of property "LayoutConstraints.availableBlockSize".
//
// The returned bool will be false if there is no such property.
func (this LayoutConstraints) AvailableBlockSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutConstraintsAvailableBlockSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// FixedInlineSize returns the value of property "LayoutConstraints.fixedInlineSize".
//
// The returned bool will be false if there is no such property.
func (this LayoutConstraints) FixedInlineSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutConstraintsFixedInlineSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// FixedBlockSize returns the value of property "LayoutConstraints.fixedBlockSize".
//
// The returned bool will be false if there is no such property.
func (this LayoutConstraints) FixedBlockSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutConstraintsFixedBlockSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// PercentageInlineSize returns the value of property "LayoutConstraints.percentageInlineSize".
//
// The returned bool will be false if there is no such property.
func (this LayoutConstraints) PercentageInlineSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutConstraintsPercentageInlineSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// PercentageBlockSize returns the value of property "LayoutConstraints.percentageBlockSize".
//
// The returned bool will be false if there is no such property.
func (this LayoutConstraints) PercentageBlockSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutConstraintsPercentageBlockSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// BlockFragmentationOffset returns the value of property "LayoutConstraints.blockFragmentationOffset".
//
// The returned bool will be false if there is no such property.
func (this LayoutConstraints) BlockFragmentationOffset() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutConstraintsBlockFragmentationOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// BlockFragmentationType returns the value of property "LayoutConstraints.blockFragmentationType".
//
// The returned bool will be false if there is no such property.
func (this LayoutConstraints) BlockFragmentationType() (BlockFragmentationType, bool) {
	var _ok bool
	_ret := bindings.GetLayoutConstraintsBlockFragmentationType(
		this.Ref(), js.Pointer(&_ok),
	)
	return BlockFragmentationType(_ret), _ok
}

// Data returns the value of property "LayoutConstraints.data".
//
// The returned bool will be false if there is no such property.
func (this LayoutConstraints) Data() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetLayoutConstraintsData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

type LayoutEdges struct {
	ref js.Ref
}

func (this LayoutEdges) Once() LayoutEdges {
	this.Ref().Once()
	return this
}

func (this LayoutEdges) Ref() js.Ref {
	return this.ref
}

func (this LayoutEdges) FromRef(ref js.Ref) LayoutEdges {
	this.ref = ref
	return this
}

func (this LayoutEdges) Free() {
	this.Ref().Free()
}

// InlineStart returns the value of property "LayoutEdges.inlineStart".
//
// The returned bool will be false if there is no such property.
func (this LayoutEdges) InlineStart() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutEdgesInlineStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// InlineEnd returns the value of property "LayoutEdges.inlineEnd".
//
// The returned bool will be false if there is no such property.
func (this LayoutEdges) InlineEnd() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutEdgesInlineEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// BlockStart returns the value of property "LayoutEdges.blockStart".
//
// The returned bool will be false if there is no such property.
func (this LayoutEdges) BlockStart() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutEdgesBlockStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// BlockEnd returns the value of property "LayoutEdges.blockEnd".
//
// The returned bool will be false if there is no such property.
func (this LayoutEdges) BlockEnd() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutEdgesBlockEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Inline returns the value of property "LayoutEdges.inline".
//
// The returned bool will be false if there is no such property.
func (this LayoutEdges) Inline() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutEdgesInline(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Block returns the value of property "LayoutEdges.block".
//
// The returned bool will be false if there is no such property.
func (this LayoutEdges) Block() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutEdgesBlock(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

type LayoutSizingMode uint32

const (
	_ LayoutSizingMode = iota

	LayoutSizingMode_BLOCK_LIKE
	LayoutSizingMode_MANUAL
)

func (LayoutSizingMode) FromRef(str js.Ref) LayoutSizingMode {
	return LayoutSizingMode(bindings.ConstOfLayoutSizingMode(str))
}

func (x LayoutSizingMode) String() (string, bool) {
	switch x {
	case LayoutSizingMode_BLOCK_LIKE:
		return "block-like", true
	case LayoutSizingMode_MANUAL:
		return "manual", true
	default:
		return "", false
	}
}

type LayoutOptions struct {
	// ChildDisplay is "LayoutOptions.childDisplay"
	//
	// Optional, defaults to "block".
	ChildDisplay ChildDisplayType
	// Sizing is "LayoutOptions.sizing"
	//
	// Optional, defaults to "block-like".
	Sizing LayoutSizingMode

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LayoutOptions with all fields set.
func (p LayoutOptions) FromRef(ref js.Ref) LayoutOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 LayoutOptions LayoutOptions [// LayoutOptions] [0x1400088dea0 0x1400088df40] 0x140008be060 {0 0}} in the application heap.
func (p LayoutOptions) New() js.Ref {
	return bindings.LayoutOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p LayoutOptions) UpdateFrom(ref js.Ref) {
	bindings.LayoutOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p LayoutOptions) Update(ref js.Ref) {
	bindings.LayoutOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type LayoutShiftAttribution struct {
	ref js.Ref
}

func (this LayoutShiftAttribution) Once() LayoutShiftAttribution {
	this.Ref().Once()
	return this
}

func (this LayoutShiftAttribution) Ref() js.Ref {
	return this.ref
}

func (this LayoutShiftAttribution) FromRef(ref js.Ref) LayoutShiftAttribution {
	this.ref = ref
	return this
}

func (this LayoutShiftAttribution) Free() {
	this.Ref().Free()
}

// Node returns the value of property "LayoutShiftAttribution.node".
//
// The returned bool will be false if there is no such property.
func (this LayoutShiftAttribution) Node() (Node, bool) {
	var _ok bool
	_ret := bindings.GetLayoutShiftAttributionNode(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// PreviousRect returns the value of property "LayoutShiftAttribution.previousRect".
//
// The returned bool will be false if there is no such property.
func (this LayoutShiftAttribution) PreviousRect() (DOMRectReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetLayoutShiftAttributionPreviousRect(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMRectReadOnly{}.FromRef(_ret), _ok
}

// CurrentRect returns the value of property "LayoutShiftAttribution.currentRect".
//
// The returned bool will be false if there is no such property.
func (this LayoutShiftAttribution) CurrentRect() (DOMRectReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetLayoutShiftAttributionCurrentRect(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMRectReadOnly{}.FromRef(_ret), _ok
}

type LayoutShift struct {
	PerformanceEntry
}

func (this LayoutShift) Once() LayoutShift {
	this.Ref().Once()
	return this
}

func (this LayoutShift) Ref() js.Ref {
	return this.PerformanceEntry.Ref()
}

func (this LayoutShift) FromRef(ref js.Ref) LayoutShift {
	this.PerformanceEntry = this.PerformanceEntry.FromRef(ref)
	return this
}

func (this LayoutShift) Free() {
	this.Ref().Free()
}

// Value returns the value of property "LayoutShift.value".
//
// The returned bool will be false if there is no such property.
func (this LayoutShift) Value() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutShiftValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// HadRecentInput returns the value of property "LayoutShift.hadRecentInput".
//
// The returned bool will be false if there is no such property.
func (this LayoutShift) HadRecentInput() (bool, bool) {
	var _ok bool
	_ret := bindings.GetLayoutShiftHadRecentInput(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// LastInputTime returns the value of property "LayoutShift.lastInputTime".
//
// The returned bool will be false if there is no such property.
func (this LayoutShift) LastInputTime() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetLayoutShiftLastInputTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// Sources returns the value of property "LayoutShift.sources".
//
// The returned bool will be false if there is no such property.
func (this LayoutShift) Sources() (js.FrozenArray[LayoutShiftAttribution], bool) {
	var _ok bool
	_ret := bindings.GetLayoutShiftSources(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[LayoutShiftAttribution]{}.FromRef(_ret), _ok
}

// ToJSON calls the method "LayoutShift.toJSON".
//
// The returned bool will be false if there is no such method.
func (this LayoutShift) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallLayoutShiftToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "LayoutShift.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this LayoutShift) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.LayoutShiftToJSONFunc(
			this.Ref(),
		),
	)
}

type LayoutWorkletGlobalScope struct {
	WorkletGlobalScope
}

func (this LayoutWorkletGlobalScope) Once() LayoutWorkletGlobalScope {
	this.Ref().Once()
	return this
}

func (this LayoutWorkletGlobalScope) Ref() js.Ref {
	return this.WorkletGlobalScope.Ref()
}

func (this LayoutWorkletGlobalScope) FromRef(ref js.Ref) LayoutWorkletGlobalScope {
	this.WorkletGlobalScope = this.WorkletGlobalScope.FromRef(ref)
	return this
}

func (this LayoutWorkletGlobalScope) Free() {
	this.Ref().Free()
}

// RegisterLayout calls the method "LayoutWorkletGlobalScope.registerLayout".
//
// The returned bool will be false if there is no such method.
func (this LayoutWorkletGlobalScope) RegisterLayout(name js.String, layoutCtor js.Func[func()]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallLayoutWorkletGlobalScopeRegisterLayout(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		layoutCtor.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RegisterLayoutFunc returns the method "LayoutWorkletGlobalScope.registerLayout".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this LayoutWorkletGlobalScope) RegisterLayoutFunc() (fn js.Func[func(name js.String, layoutCtor js.Func[func()])]) {
	return fn.FromRef(
		bindings.LayoutWorkletGlobalScopeRegisterLayoutFunc(
			this.Ref(),
		),
	)
}

type LineAlignSetting uint32

const (
	_ LineAlignSetting = iota

	LineAlignSetting_START
	LineAlignSetting_CENTER
	LineAlignSetting_END
)

func (LineAlignSetting) FromRef(str js.Ref) LineAlignSetting {
	return LineAlignSetting(bindings.ConstOfLineAlignSetting(str))
}

func (x LineAlignSetting) String() (string, bool) {
	switch x {
	case LineAlignSetting_START:
		return "start", true
	case LineAlignSetting_CENTER:
		return "center", true
	case LineAlignSetting_END:
		return "end", true
	default:
		return "", false
	}
}

type OneOf_Float64_AutoKeyword struct {
	ref js.Ref
}

func (x OneOf_Float64_AutoKeyword) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Float64_AutoKeyword) Free() {
	x.ref.Free()
}

func (x OneOf_Float64_AutoKeyword) FromRef(ref js.Ref) OneOf_Float64_AutoKeyword {
	return OneOf_Float64_AutoKeyword{
		ref: ref,
	}
}

func (x OneOf_Float64_AutoKeyword) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Float64_AutoKeyword) AutoKeyword() AutoKeyword {
	return AutoKeyword(0).FromRef(x.ref)
}

type LineAndPositionSetting = OneOf_Float64_AutoKeyword

type LinearAccelerationReadingValues struct {
	// X is "LinearAccelerationReadingValues.x"
	//
	// Required
	X float64
	// Y is "LinearAccelerationReadingValues.y"
	//
	// Required
	Y float64
	// Z is "LinearAccelerationReadingValues.z"
	//
	// Required
	Z float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LinearAccelerationReadingValues with all fields set.
func (p LinearAccelerationReadingValues) FromRef(ref js.Ref) LinearAccelerationReadingValues {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 LinearAccelerationReadingValues LinearAccelerationReadingValues [// LinearAccelerationReadingValues] [0x140008c2140 0x140008c21e0 0x140008c2280] 0x140008be078 {0 0}} in the application heap.
func (p LinearAccelerationReadingValues) New() js.Ref {
	return bindings.LinearAccelerationReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p LinearAccelerationReadingValues) UpdateFrom(ref js.Ref) {
	bindings.LinearAccelerationReadingValuesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p LinearAccelerationReadingValues) Update(ref js.Ref) {
	bindings.LinearAccelerationReadingValuesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewLinearAccelerationSensor(options AccelerometerSensorOptions) LinearAccelerationSensor {
	return LinearAccelerationSensor{}.FromRef(
		bindings.NewLinearAccelerationSensorByLinearAccelerationSensor(
			js.Pointer(&options)),
	)
}

func NewLinearAccelerationSensorByLinearAccelerationSensor1() LinearAccelerationSensor {
	return LinearAccelerationSensor{}.FromRef(
		bindings.NewLinearAccelerationSensorByLinearAccelerationSensor1(),
	)
}

type LinearAccelerationSensor struct {
	Accelerometer
}

func (this LinearAccelerationSensor) Once() LinearAccelerationSensor {
	this.Ref().Once()
	return this
}

func (this LinearAccelerationSensor) Ref() js.Ref {
	return this.Accelerometer.Ref()
}

func (this LinearAccelerationSensor) FromRef(ref js.Ref) LinearAccelerationSensor {
	this.Accelerometer = this.Accelerometer.FromRef(ref)
	return this
}

func (this LinearAccelerationSensor) Free() {
	this.Ref().Free()
}

type MIDIPortType uint32

const (
	_ MIDIPortType = iota

	MIDIPortType_INPUT
	MIDIPortType_OUTPUT
)

func (MIDIPortType) FromRef(str js.Ref) MIDIPortType {
	return MIDIPortType(bindings.ConstOfMIDIPortType(str))
}

func (x MIDIPortType) String() (string, bool) {
	switch x {
	case MIDIPortType_INPUT:
		return "input", true
	case MIDIPortType_OUTPUT:
		return "output", true
	default:
		return "", false
	}
}

type MIDIPortDeviceState uint32

const (
	_ MIDIPortDeviceState = iota

	MIDIPortDeviceState_DISCONNECTED
	MIDIPortDeviceState_CONNECTED
)

func (MIDIPortDeviceState) FromRef(str js.Ref) MIDIPortDeviceState {
	return MIDIPortDeviceState(bindings.ConstOfMIDIPortDeviceState(str))
}

func (x MIDIPortDeviceState) String() (string, bool) {
	switch x {
	case MIDIPortDeviceState_DISCONNECTED:
		return "disconnected", true
	case MIDIPortDeviceState_CONNECTED:
		return "connected", true
	default:
		return "", false
	}
}

type MIDIPortConnectionState uint32

const (
	_ MIDIPortConnectionState = iota

	MIDIPortConnectionState_OPEN
	MIDIPortConnectionState_CLOSED
	MIDIPortConnectionState_PENDING
)

func (MIDIPortConnectionState) FromRef(str js.Ref) MIDIPortConnectionState {
	return MIDIPortConnectionState(bindings.ConstOfMIDIPortConnectionState(str))
}

func (x MIDIPortConnectionState) String() (string, bool) {
	switch x {
	case MIDIPortConnectionState_OPEN:
		return "open", true
	case MIDIPortConnectionState_CLOSED:
		return "closed", true
	case MIDIPortConnectionState_PENDING:
		return "pending", true
	default:
		return "", false
	}
}

type MIDIPort struct {
	EventTarget
}

func (this MIDIPort) Once() MIDIPort {
	this.Ref().Once()
	return this
}

func (this MIDIPort) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this MIDIPort) FromRef(ref js.Ref) MIDIPort {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this MIDIPort) Free() {
	this.Ref().Free()
}

// Id returns the value of property "MIDIPort.id".
//
// The returned bool will be false if there is no such property.
func (this MIDIPort) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMIDIPortId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Manufacturer returns the value of property "MIDIPort.manufacturer".
//
// The returned bool will be false if there is no such property.
func (this MIDIPort) Manufacturer() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMIDIPortManufacturer(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name returns the value of property "MIDIPort.name".
//
// The returned bool will be false if there is no such property.
func (this MIDIPort) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMIDIPortName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type returns the value of property "MIDIPort.type".
//
// The returned bool will be false if there is no such property.
func (this MIDIPort) Type() (MIDIPortType, bool) {
	var _ok bool
	_ret := bindings.GetMIDIPortType(
		this.Ref(), js.Pointer(&_ok),
	)
	return MIDIPortType(_ret), _ok
}

// Version returns the value of property "MIDIPort.version".
//
// The returned bool will be false if there is no such property.
func (this MIDIPort) Version() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMIDIPortVersion(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// State returns the value of property "MIDIPort.state".
//
// The returned bool will be false if there is no such property.
func (this MIDIPort) State() (MIDIPortDeviceState, bool) {
	var _ok bool
	_ret := bindings.GetMIDIPortState(
		this.Ref(), js.Pointer(&_ok),
	)
	return MIDIPortDeviceState(_ret), _ok
}

// Connection returns the value of property "MIDIPort.connection".
//
// The returned bool will be false if there is no such property.
func (this MIDIPort) Connection() (MIDIPortConnectionState, bool) {
	var _ok bool
	_ret := bindings.GetMIDIPortConnection(
		this.Ref(), js.Pointer(&_ok),
	)
	return MIDIPortConnectionState(_ret), _ok
}

// Open calls the method "MIDIPort.open".
//
// The returned bool will be false if there is no such method.
func (this MIDIPort) Open() (js.Promise[MIDIPort], bool) {
	var _ok bool
	_ret := bindings.CallMIDIPortOpen(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[MIDIPort]{}.FromRef(_ret), _ok
}

// OpenFunc returns the method "MIDIPort.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MIDIPort) OpenFunc() (fn js.Func[func() js.Promise[MIDIPort]]) {
	return fn.FromRef(
		bindings.MIDIPortOpenFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "MIDIPort.close".
//
// The returned bool will be false if there is no such method.
func (this MIDIPort) Close() (js.Promise[MIDIPort], bool) {
	var _ok bool
	_ret := bindings.CallMIDIPortClose(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[MIDIPort]{}.FromRef(_ret), _ok
}

// CloseFunc returns the method "MIDIPort.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MIDIPort) CloseFunc() (fn js.Func[func() js.Promise[MIDIPort]]) {
	return fn.FromRef(
		bindings.MIDIPortCloseFunc(
			this.Ref(),
		),
	)
}

type MIDIConnectionEventInit struct {
	// Port is "MIDIConnectionEventInit.port"
	//
	// Optional
	Port MIDIPort
	// Bubbles is "MIDIConnectionEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "MIDIConnectionEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "MIDIConnectionEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MIDIConnectionEventInit with all fields set.
func (p MIDIConnectionEventInit) FromRef(ref js.Ref) MIDIConnectionEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MIDIConnectionEventInit MIDIConnectionEventInit [// MIDIConnectionEventInit] [0x140008c23c0 0x140008c2460 0x140008c25a0 0x140008c26e0 0x140008c2500 0x140008c2640 0x140008c2780] 0x140008be108 {0 0}} in the application heap.
func (p MIDIConnectionEventInit) New() js.Ref {
	return bindings.MIDIConnectionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MIDIConnectionEventInit) UpdateFrom(ref js.Ref) {
	bindings.MIDIConnectionEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MIDIConnectionEventInit) Update(ref js.Ref) {
	bindings.MIDIConnectionEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewMIDIConnectionEvent(typ js.String, eventInitDict MIDIConnectionEventInit) MIDIConnectionEvent {
	return MIDIConnectionEvent{}.FromRef(
		bindings.NewMIDIConnectionEventByMIDIConnectionEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewMIDIConnectionEventByMIDIConnectionEvent1(typ js.String) MIDIConnectionEvent {
	return MIDIConnectionEvent{}.FromRef(
		bindings.NewMIDIConnectionEventByMIDIConnectionEvent1(
			typ.Ref()),
	)
}

type MIDIConnectionEvent struct {
	Event
}

func (this MIDIConnectionEvent) Once() MIDIConnectionEvent {
	this.Ref().Once()
	return this
}

func (this MIDIConnectionEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this MIDIConnectionEvent) FromRef(ref js.Ref) MIDIConnectionEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this MIDIConnectionEvent) Free() {
	this.Ref().Free()
}

// Port returns the value of property "MIDIConnectionEvent.port".
//
// The returned bool will be false if there is no such property.
func (this MIDIConnectionEvent) Port() (MIDIPort, bool) {
	var _ok bool
	_ret := bindings.GetMIDIConnectionEventPort(
		this.Ref(), js.Pointer(&_ok),
	)
	return MIDIPort{}.FromRef(_ret), _ok
}

type MIDIInput struct {
	MIDIPort
}

func (this MIDIInput) Once() MIDIInput {
	this.Ref().Once()
	return this
}

func (this MIDIInput) Ref() js.Ref {
	return this.MIDIPort.Ref()
}

func (this MIDIInput) FromRef(ref js.Ref) MIDIInput {
	this.MIDIPort = this.MIDIPort.FromRef(ref)
	return this
}

func (this MIDIInput) Free() {
	this.Ref().Free()
}

type MIDIMessageEventInit struct {
	// Data is "MIDIMessageEventInit.data"
	//
	// Optional
	Data js.TypedArray[uint8]
	// Bubbles is "MIDIMessageEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "MIDIMessageEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "MIDIMessageEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MIDIMessageEventInit with all fields set.
func (p MIDIMessageEventInit) FromRef(ref js.Ref) MIDIMessageEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MIDIMessageEventInit MIDIMessageEventInit [// MIDIMessageEventInit] [0x140008c2820 0x140008c28c0 0x140008c2a00 0x140008c2b40 0x140008c2960 0x140008c2aa0 0x140008c2be0] 0x140008be1c8 {0 0}} in the application heap.
func (p MIDIMessageEventInit) New() js.Ref {
	return bindings.MIDIMessageEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MIDIMessageEventInit) UpdateFrom(ref js.Ref) {
	bindings.MIDIMessageEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MIDIMessageEventInit) Update(ref js.Ref) {
	bindings.MIDIMessageEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewMIDIMessageEvent(typ js.String, eventInitDict MIDIMessageEventInit) MIDIMessageEvent {
	return MIDIMessageEvent{}.FromRef(
		bindings.NewMIDIMessageEventByMIDIMessageEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewMIDIMessageEventByMIDIMessageEvent1(typ js.String) MIDIMessageEvent {
	return MIDIMessageEvent{}.FromRef(
		bindings.NewMIDIMessageEventByMIDIMessageEvent1(
			typ.Ref()),
	)
}

type MIDIMessageEvent struct {
	Event
}

func (this MIDIMessageEvent) Once() MIDIMessageEvent {
	this.Ref().Once()
	return this
}

func (this MIDIMessageEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this MIDIMessageEvent) FromRef(ref js.Ref) MIDIMessageEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this MIDIMessageEvent) Free() {
	this.Ref().Free()
}

// Data returns the value of property "MIDIMessageEvent.data".
//
// The returned bool will be false if there is no such property.
func (this MIDIMessageEvent) Data() (js.TypedArray[uint8], bool) {
	var _ok bool
	_ret := bindings.GetMIDIMessageEventData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[uint8]{}.FromRef(_ret), _ok
}

type MIDIOutput struct {
	MIDIPort
}

func (this MIDIOutput) Once() MIDIOutput {
	this.Ref().Once()
	return this
}

func (this MIDIOutput) Ref() js.Ref {
	return this.MIDIPort.Ref()
}

func (this MIDIOutput) FromRef(ref js.Ref) MIDIOutput {
	this.MIDIPort = this.MIDIPort.FromRef(ref)
	return this
}

func (this MIDIOutput) Free() {
	this.Ref().Free()
}

// Send calls the method "MIDIOutput.send".
//
// The returned bool will be false if there is no such method.
func (this MIDIOutput) Send(data js.Array[uint8], timestamp DOMHighResTimeStamp) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMIDIOutputSend(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
		float64(timestamp),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SendFunc returns the method "MIDIOutput.send".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MIDIOutput) SendFunc() (fn js.Func[func(data js.Array[uint8], timestamp DOMHighResTimeStamp)]) {
	return fn.FromRef(
		bindings.MIDIOutputSendFunc(
			this.Ref(),
		),
	)
}

// Send1 calls the method "MIDIOutput.send".
//
// The returned bool will be false if there is no such method.
func (this MIDIOutput) Send1(data js.Array[uint8]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMIDIOutputSend1(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Send1Func returns the method "MIDIOutput.send".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MIDIOutput) Send1Func() (fn js.Func[func(data js.Array[uint8])]) {
	return fn.FromRef(
		bindings.MIDIOutputSend1Func(
			this.Ref(),
		),
	)
}

// Clear calls the method "MIDIOutput.clear".
//
// The returned bool will be false if there is no such method.
func (this MIDIOutput) Clear() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMIDIOutputClear(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearFunc returns the method "MIDIOutput.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MIDIOutput) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MIDIOutputClearFunc(
			this.Ref(),
		),
	)
}

type MLActivation struct {
	ref js.Ref
}

func (this MLActivation) Once() MLActivation {
	this.Ref().Once()
	return this
}

func (this MLActivation) Ref() js.Ref {
	return this.ref
}

func (this MLActivation) FromRef(ref js.Ref) MLActivation {
	this.ref = ref
	return this
}

func (this MLActivation) Free() {
	this.Ref().Free()
}

type MLAutoPad uint32

const (
	_ MLAutoPad = iota

	MLAutoPad_EXPLICIT
	MLAutoPad_SAME_UPPER
	MLAutoPad_SAME_LOWER
)

func (MLAutoPad) FromRef(str js.Ref) MLAutoPad {
	return MLAutoPad(bindings.ConstOfMLAutoPad(str))
}

func (x MLAutoPad) String() (string, bool) {
	switch x {
	case MLAutoPad_EXPLICIT:
		return "explicit", true
	case MLAutoPad_SAME_UPPER:
		return "same-upper", true
	case MLAutoPad_SAME_LOWER:
		return "same-lower", true
	default:
		return "", false
	}
}

type MLOperand struct {
	ref js.Ref
}

func (this MLOperand) Once() MLOperand {
	this.Ref().Once()
	return this
}

func (this MLOperand) Ref() js.Ref {
	return this.ref
}

func (this MLOperand) FromRef(ref js.Ref) MLOperand {
	this.ref = ref
	return this
}

func (this MLOperand) Free() {
	this.Ref().Free()
}

type MLBatchNormalizationOptions struct {
	// Scale is "MLBatchNormalizationOptions.scale"
	//
	// Optional
	Scale MLOperand
	// Bias is "MLBatchNormalizationOptions.bias"
	//
	// Optional
	Bias MLOperand
	// Axis is "MLBatchNormalizationOptions.axis"
	//
	// Optional, defaults to 1.
	Axis uint32
	// Epsilon is "MLBatchNormalizationOptions.epsilon"
	//
	// Optional, defaults to 1e-5.
	Epsilon float32
	// Activation is "MLBatchNormalizationOptions.activation"
	//
	// Optional
	Activation MLActivation

	FFI_USE_Axis    bool // for Axis.
	FFI_USE_Epsilon bool // for Epsilon.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLBatchNormalizationOptions with all fields set.
func (p MLBatchNormalizationOptions) FromRef(ref js.Ref) MLBatchNormalizationOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MLBatchNormalizationOptions MLBatchNormalizationOptions [// MLBatchNormalizationOptions] [0x140008c2c80 0x140008c2d20 0x140008c2dc0 0x140008c2f00 0x140008c3040 0x140008c2e60 0x140008c2fa0] 0x140008be258 {0 0}} in the application heap.
func (p MLBatchNormalizationOptions) New() js.Ref {
	return bindings.MLBatchNormalizationOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MLBatchNormalizationOptions) UpdateFrom(ref js.Ref) {
	bindings.MLBatchNormalizationOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLBatchNormalizationOptions) Update(ref js.Ref) {
	bindings.MLBatchNormalizationOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MLBufferResourceView struct {
	// Resource is "MLBufferResourceView.resource"
	//
	// Required
	Resource GPUBuffer
	// Offset is "MLBufferResourceView.offset"
	//
	// Optional, defaults to 0.
	Offset uint64
	// Size is "MLBufferResourceView.size"
	//
	// Optional
	Size uint64

	FFI_USE_Offset bool // for Offset.
	FFI_USE_Size   bool // for Size.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLBufferResourceView with all fields set.
func (p MLBufferResourceView) FromRef(ref js.Ref) MLBufferResourceView {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MLBufferResourceView MLBufferResourceView [// MLBufferResourceView] [0x140008c30e0 0x140008c3180 0x140008c32c0 0x140008c3220 0x140008c3360] 0x140008be2b8 {0 0}} in the application heap.
func (p MLBufferResourceView) New() js.Ref {
	return bindings.MLBufferResourceViewJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MLBufferResourceView) UpdateFrom(ref js.Ref) {
	bindings.MLBufferResourceViewJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLBufferResourceView) Update(ref js.Ref) {
	bindings.MLBufferResourceViewJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView struct {
	ref js.Ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView) FromRef(ref js.Ref) OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView {
	return OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView{
		ref: ref,
	}
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView) MLBufferResourceView() MLBufferResourceView {
	var ret MLBufferResourceView
	ret.UpdateFrom(x.ref)
	return ret
}

type MLBufferView = OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_MLBufferResourceView

type MLClampOptions struct {
	// MinValue is "MLClampOptions.minValue"
	//
	// Optional
	MinValue float32
	// MaxValue is "MLClampOptions.maxValue"
	//
	// Optional
	MaxValue float32

	FFI_USE_MinValue bool // for MinValue.
	FFI_USE_MaxValue bool // for MaxValue.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLClampOptions with all fields set.
func (p MLClampOptions) FromRef(ref js.Ref) MLClampOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MLClampOptions MLClampOptions [// MLClampOptions] [0x140008c3400 0x140008c3540 0x140008c34a0 0x140008c35e0] 0x140008be378 {0 0}} in the application heap.
func (p MLClampOptions) New() js.Ref {
	return bindings.MLClampOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MLClampOptions) UpdateFrom(ref js.Ref) {
	bindings.MLClampOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLClampOptions) Update(ref js.Ref) {
	bindings.MLClampOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MLConv2dFilterOperandLayout uint32

const (
	_ MLConv2dFilterOperandLayout = iota

	MLConv2dFilterOperandLayout_OIHW
	MLConv2dFilterOperandLayout_HWIO
	MLConv2dFilterOperandLayout_OHWI
	MLConv2dFilterOperandLayout_IHWO
)

func (MLConv2dFilterOperandLayout) FromRef(str js.Ref) MLConv2dFilterOperandLayout {
	return MLConv2dFilterOperandLayout(bindings.ConstOfMLConv2dFilterOperandLayout(str))
}

func (x MLConv2dFilterOperandLayout) String() (string, bool) {
	switch x {
	case MLConv2dFilterOperandLayout_OIHW:
		return "oihw", true
	case MLConv2dFilterOperandLayout_HWIO:
		return "hwio", true
	case MLConv2dFilterOperandLayout_OHWI:
		return "ohwi", true
	case MLConv2dFilterOperandLayout_IHWO:
		return "ihwo", true
	default:
		return "", false
	}
}

type MLInputOperandLayout uint32

const (
	_ MLInputOperandLayout = iota

	MLInputOperandLayout_NCHW
	MLInputOperandLayout_NHWC
)

func (MLInputOperandLayout) FromRef(str js.Ref) MLInputOperandLayout {
	return MLInputOperandLayout(bindings.ConstOfMLInputOperandLayout(str))
}

func (x MLInputOperandLayout) String() (string, bool) {
	switch x {
	case MLInputOperandLayout_NCHW:
		return "nchw", true
	case MLInputOperandLayout_NHWC:
		return "nhwc", true
	default:
		return "", false
	}
}

type MLConv2dOptions struct {
	// Padding is "MLConv2dOptions.padding"
	//
	// Optional
	Padding js.Array[uint32]
	// Strides is "MLConv2dOptions.strides"
	//
	// Optional
	Strides js.Array[uint32]
	// Dilations is "MLConv2dOptions.dilations"
	//
	// Optional
	Dilations js.Array[uint32]
	// AutoPad is "MLConv2dOptions.autoPad"
	//
	// Optional, defaults to "explicit".
	AutoPad MLAutoPad
	// Groups is "MLConv2dOptions.groups"
	//
	// Optional, defaults to 1.
	Groups uint32
	// InputLayout is "MLConv2dOptions.inputLayout"
	//
	// Optional, defaults to "nchw".
	InputLayout MLInputOperandLayout
	// FilterLayout is "MLConv2dOptions.filterLayout"
	//
	// Optional, defaults to "oihw".
	FilterLayout MLConv2dFilterOperandLayout
	// Bias is "MLConv2dOptions.bias"
	//
	// Optional
	Bias MLOperand
	// Activation is "MLConv2dOptions.activation"
	//
	// Optional
	Activation MLActivation

	FFI_USE_Groups bool // for Groups.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLConv2dOptions with all fields set.
func (p MLConv2dOptions) FromRef(ref js.Ref) MLConv2dOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MLConv2dOptions MLConv2dOptions [// MLConv2dOptions] [0x140008c3680 0x140008c3720 0x140008c37c0 0x140008c3860 0x140008c3900 0x140008c3a40 0x140008c3ae0 0x140008c3b80 0x140008c3c20 0x140008c39a0] 0x140008be3a8 {0 0}} in the application heap.
func (p MLConv2dOptions) New() js.Ref {
	return bindings.MLConv2dOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MLConv2dOptions) UpdateFrom(ref js.Ref) {
	bindings.MLConv2dOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLConv2dOptions) Update(ref js.Ref) {
	bindings.MLConv2dOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MLConvTranspose2dFilterOperandLayout uint32

const (
	_ MLConvTranspose2dFilterOperandLayout = iota

	MLConvTranspose2dFilterOperandLayout_IOHW
	MLConvTranspose2dFilterOperandLayout_HWOI
	MLConvTranspose2dFilterOperandLayout_OHWI
)

func (MLConvTranspose2dFilterOperandLayout) FromRef(str js.Ref) MLConvTranspose2dFilterOperandLayout {
	return MLConvTranspose2dFilterOperandLayout(bindings.ConstOfMLConvTranspose2dFilterOperandLayout(str))
}

func (x MLConvTranspose2dFilterOperandLayout) String() (string, bool) {
	switch x {
	case MLConvTranspose2dFilterOperandLayout_IOHW:
		return "iohw", true
	case MLConvTranspose2dFilterOperandLayout_HWOI:
		return "hwoi", true
	case MLConvTranspose2dFilterOperandLayout_OHWI:
		return "ohwi", true
	default:
		return "", false
	}
}

type MLConvTranspose2dOptions struct {
	// Padding is "MLConvTranspose2dOptions.padding"
	//
	// Optional
	Padding js.Array[uint32]
	// Strides is "MLConvTranspose2dOptions.strides"
	//
	// Optional
	Strides js.Array[uint32]
	// Dilations is "MLConvTranspose2dOptions.dilations"
	//
	// Optional
	Dilations js.Array[uint32]
	// OutputPadding is "MLConvTranspose2dOptions.outputPadding"
	//
	// Optional
	OutputPadding js.Array[uint32]
	// OutputSizes is "MLConvTranspose2dOptions.outputSizes"
	//
	// Optional
	OutputSizes js.Array[uint32]
	// AutoPad is "MLConvTranspose2dOptions.autoPad"
	//
	// Optional, defaults to "explicit".
	AutoPad MLAutoPad
	// Groups is "MLConvTranspose2dOptions.groups"
	//
	// Optional, defaults to 1.
	Groups uint32
	// InputLayout is "MLConvTranspose2dOptions.inputLayout"
	//
	// Optional, defaults to "nchw".
	InputLayout MLInputOperandLayout
	// FilterLayout is "MLConvTranspose2dOptions.filterLayout"
	//
	// Optional, defaults to "iohw".
	FilterLayout MLConvTranspose2dFilterOperandLayout
	// Bias is "MLConvTranspose2dOptions.bias"
	//
	// Optional
	Bias MLOperand
	// Activation is "MLConvTranspose2dOptions.activation"
	//
	// Optional
	Activation MLActivation

	FFI_USE_Groups bool // for Groups.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLConvTranspose2dOptions with all fields set.
func (p MLConvTranspose2dOptions) FromRef(ref js.Ref) MLConvTranspose2dOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MLConvTranspose2dOptions MLConvTranspose2dOptions [// MLConvTranspose2dOptions] [0x140008c3cc0 0x140008c3d60 0x140008c3e00 0x140008c3ea0 0x140008c3f40 0x140008d0000 0x140008d00a0 0x140008d01e0 0x140008d0280 0x140008d0320 0x140008d03c0 0x140008d0140] 0x140008be3d8 {0 0}} in the application heap.
func (p MLConvTranspose2dOptions) New() js.Ref {
	return bindings.MLConvTranspose2dOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MLConvTranspose2dOptions) UpdateFrom(ref js.Ref) {
	bindings.MLConvTranspose2dOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLConvTranspose2dOptions) Update(ref js.Ref) {
	bindings.MLConvTranspose2dOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MLEluOptions struct {
	// Alpha is "MLEluOptions.alpha"
	//
	// Optional, defaults to 1.
	Alpha float32

	FFI_USE_Alpha bool // for Alpha.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLEluOptions with all fields set.
func (p MLEluOptions) FromRef(ref js.Ref) MLEluOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MLEluOptions MLEluOptions [// MLEluOptions] [0x140008d0460 0x140008d0500] 0x140008be408 {0 0}} in the application heap.
func (p MLEluOptions) New() js.Ref {
	return bindings.MLEluOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MLEluOptions) UpdateFrom(ref js.Ref) {
	bindings.MLEluOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLEluOptions) Update(ref js.Ref) {
	bindings.MLEluOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MLGemmOptions struct {
	// C is "MLGemmOptions.c"
	//
	// Optional
	C MLOperand
	// Alpha is "MLGemmOptions.alpha"
	//
	// Optional, defaults to 1.0.
	Alpha float32
	// Beta is "MLGemmOptions.beta"
	//
	// Optional, defaults to 1.0.
	Beta float32
	// ATranspose is "MLGemmOptions.aTranspose"
	//
	// Optional, defaults to false.
	ATranspose bool
	// BTranspose is "MLGemmOptions.bTranspose"
	//
	// Optional, defaults to false.
	BTranspose bool

	FFI_USE_Alpha      bool // for Alpha.
	FFI_USE_Beta       bool // for Beta.
	FFI_USE_ATranspose bool // for ATranspose.
	FFI_USE_BTranspose bool // for BTranspose.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLGemmOptions with all fields set.
func (p MLGemmOptions) FromRef(ref js.Ref) MLGemmOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MLGemmOptions MLGemmOptions [// MLGemmOptions] [0x140008d05a0 0x140008d0640 0x140008d0780 0x140008d08c0 0x140008d0a00 0x140008d06e0 0x140008d0820 0x140008d0960 0x140008d0aa0] 0x140008be420 {0 0}} in the application heap.
func (p MLGemmOptions) New() js.Ref {
	return bindings.MLGemmOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MLGemmOptions) UpdateFrom(ref js.Ref) {
	bindings.MLGemmOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLGemmOptions) Update(ref js.Ref) {
	bindings.MLGemmOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MLOperandType uint32
