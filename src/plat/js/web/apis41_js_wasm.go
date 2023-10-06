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
// It returns ok=false if there is no such property.
func (this LargestContentfulPaint) RenderTime() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetLargestContentfulPaintRenderTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LoadTime returns the value of property "LargestContentfulPaint.loadTime".
//
// It returns ok=false if there is no such property.
func (this LargestContentfulPaint) LoadTime() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetLargestContentfulPaintLoadTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Size returns the value of property "LargestContentfulPaint.size".
//
// It returns ok=false if there is no such property.
func (this LargestContentfulPaint) Size() (ret uint32, ok bool) {
	ok = js.True == bindings.GetLargestContentfulPaintSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "LargestContentfulPaint.id".
//
// It returns ok=false if there is no such property.
func (this LargestContentfulPaint) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLargestContentfulPaintId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Url returns the value of property "LargestContentfulPaint.url".
//
// It returns ok=false if there is no such property.
func (this LargestContentfulPaint) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLargestContentfulPaintUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Element returns the value of property "LargestContentfulPaint.element".
//
// It returns ok=false if there is no such property.
func (this LargestContentfulPaint) Element() (ret Element, ok bool) {
	ok = js.True == bindings.GetLargestContentfulPaintElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "LargestContentfulPaint.toJSON" exists.
func (this LargestContentfulPaint) HasToJSON() bool {
	return js.True == bindings.HasLargestContentfulPaintToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "LargestContentfulPaint.toJSON".
func (this LargestContentfulPaint) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.LargestContentfulPaintToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "LargestContentfulPaint.toJSON".
func (this LargestContentfulPaint) ToJSON() (ret js.Object) {
	bindings.CallLargestContentfulPaintToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "LargestContentfulPaint.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this LargestContentfulPaint) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLargestContentfulPaintToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this LayoutConstraints) AvailableInlineSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutConstraintsAvailableInlineSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AvailableBlockSize returns the value of property "LayoutConstraints.availableBlockSize".
//
// It returns ok=false if there is no such property.
func (this LayoutConstraints) AvailableBlockSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutConstraintsAvailableBlockSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FixedInlineSize returns the value of property "LayoutConstraints.fixedInlineSize".
//
// It returns ok=false if there is no such property.
func (this LayoutConstraints) FixedInlineSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutConstraintsFixedInlineSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FixedBlockSize returns the value of property "LayoutConstraints.fixedBlockSize".
//
// It returns ok=false if there is no such property.
func (this LayoutConstraints) FixedBlockSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutConstraintsFixedBlockSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PercentageInlineSize returns the value of property "LayoutConstraints.percentageInlineSize".
//
// It returns ok=false if there is no such property.
func (this LayoutConstraints) PercentageInlineSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutConstraintsPercentageInlineSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PercentageBlockSize returns the value of property "LayoutConstraints.percentageBlockSize".
//
// It returns ok=false if there is no such property.
func (this LayoutConstraints) PercentageBlockSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutConstraintsPercentageBlockSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BlockFragmentationOffset returns the value of property "LayoutConstraints.blockFragmentationOffset".
//
// It returns ok=false if there is no such property.
func (this LayoutConstraints) BlockFragmentationOffset() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutConstraintsBlockFragmentationOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BlockFragmentationType returns the value of property "LayoutConstraints.blockFragmentationType".
//
// It returns ok=false if there is no such property.
func (this LayoutConstraints) BlockFragmentationType() (ret BlockFragmentationType, ok bool) {
	ok = js.True == bindings.GetLayoutConstraintsBlockFragmentationType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Data returns the value of property "LayoutConstraints.data".
//
// It returns ok=false if there is no such property.
func (this LayoutConstraints) Data() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetLayoutConstraintsData(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this LayoutEdges) InlineStart() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutEdgesInlineStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// InlineEnd returns the value of property "LayoutEdges.inlineEnd".
//
// It returns ok=false if there is no such property.
func (this LayoutEdges) InlineEnd() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutEdgesInlineEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BlockStart returns the value of property "LayoutEdges.blockStart".
//
// It returns ok=false if there is no such property.
func (this LayoutEdges) BlockStart() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutEdgesBlockStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BlockEnd returns the value of property "LayoutEdges.blockEnd".
//
// It returns ok=false if there is no such property.
func (this LayoutEdges) BlockEnd() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutEdgesBlockEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Inline returns the value of property "LayoutEdges.inline".
//
// It returns ok=false if there is no such property.
func (this LayoutEdges) Inline() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutEdgesInline(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Block returns the value of property "LayoutEdges.block".
//
// It returns ok=false if there is no such property.
func (this LayoutEdges) Block() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutEdgesBlock(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// New creates a new LayoutOptions in the application heap.
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
// It returns ok=false if there is no such property.
func (this LayoutShiftAttribution) Node() (ret Node, ok bool) {
	ok = js.True == bindings.GetLayoutShiftAttributionNode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PreviousRect returns the value of property "LayoutShiftAttribution.previousRect".
//
// It returns ok=false if there is no such property.
func (this LayoutShiftAttribution) PreviousRect() (ret DOMRectReadOnly, ok bool) {
	ok = js.True == bindings.GetLayoutShiftAttributionPreviousRect(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CurrentRect returns the value of property "LayoutShiftAttribution.currentRect".
//
// It returns ok=false if there is no such property.
func (this LayoutShiftAttribution) CurrentRect() (ret DOMRectReadOnly, ok bool) {
	ok = js.True == bindings.GetLayoutShiftAttributionCurrentRect(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this LayoutShift) Value() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutShiftValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HadRecentInput returns the value of property "LayoutShift.hadRecentInput".
//
// It returns ok=false if there is no such property.
func (this LayoutShift) HadRecentInput() (ret bool, ok bool) {
	ok = js.True == bindings.GetLayoutShiftHadRecentInput(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LastInputTime returns the value of property "LayoutShift.lastInputTime".
//
// It returns ok=false if there is no such property.
func (this LayoutShift) LastInputTime() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetLayoutShiftLastInputTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Sources returns the value of property "LayoutShift.sources".
//
// It returns ok=false if there is no such property.
func (this LayoutShift) Sources() (ret js.FrozenArray[LayoutShiftAttribution], ok bool) {
	ok = js.True == bindings.GetLayoutShiftSources(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "LayoutShift.toJSON" exists.
func (this LayoutShift) HasToJSON() bool {
	return js.True == bindings.HasLayoutShiftToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "LayoutShift.toJSON".
func (this LayoutShift) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.LayoutShiftToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "LayoutShift.toJSON".
func (this LayoutShift) ToJSON() (ret js.Object) {
	bindings.CallLayoutShiftToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "LayoutShift.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this LayoutShift) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLayoutShiftToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// HasRegisterLayout returns true if the method "LayoutWorkletGlobalScope.registerLayout" exists.
func (this LayoutWorkletGlobalScope) HasRegisterLayout() bool {
	return js.True == bindings.HasLayoutWorkletGlobalScopeRegisterLayout(
		this.Ref(),
	)
}

// RegisterLayoutFunc returns the method "LayoutWorkletGlobalScope.registerLayout".
func (this LayoutWorkletGlobalScope) RegisterLayoutFunc() (fn js.Func[func(name js.String, layoutCtor js.Func[func()])]) {
	return fn.FromRef(
		bindings.LayoutWorkletGlobalScopeRegisterLayoutFunc(
			this.Ref(),
		),
	)
}

// RegisterLayout calls the method "LayoutWorkletGlobalScope.registerLayout".
func (this LayoutWorkletGlobalScope) RegisterLayout(name js.String, layoutCtor js.Func[func()]) (ret js.Void) {
	bindings.CallLayoutWorkletGlobalScopeRegisterLayout(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		layoutCtor.Ref(),
	)

	return
}

// TryRegisterLayout calls the method "LayoutWorkletGlobalScope.registerLayout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this LayoutWorkletGlobalScope) TryRegisterLayout(name js.String, layoutCtor js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLayoutWorkletGlobalScopeRegisterLayout(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		layoutCtor.Ref(),
	)

	return
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

// New creates a new LinearAccelerationReadingValues in the application heap.
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

func NewLinearAccelerationSensor(options AccelerometerSensorOptions) (ret LinearAccelerationSensor) {
	ret.ref = bindings.NewLinearAccelerationSensorByLinearAccelerationSensor(
		js.Pointer(&options))
	return
}

func NewLinearAccelerationSensorByLinearAccelerationSensor1() (ret LinearAccelerationSensor) {
	ret.ref = bindings.NewLinearAccelerationSensorByLinearAccelerationSensor1()
	return
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
// It returns ok=false if there is no such property.
func (this MIDIPort) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMIDIPortId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Manufacturer returns the value of property "MIDIPort.manufacturer".
//
// It returns ok=false if there is no such property.
func (this MIDIPort) Manufacturer() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMIDIPortManufacturer(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "MIDIPort.name".
//
// It returns ok=false if there is no such property.
func (this MIDIPort) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMIDIPortName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "MIDIPort.type".
//
// It returns ok=false if there is no such property.
func (this MIDIPort) Type() (ret MIDIPortType, ok bool) {
	ok = js.True == bindings.GetMIDIPortType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Version returns the value of property "MIDIPort.version".
//
// It returns ok=false if there is no such property.
func (this MIDIPort) Version() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMIDIPortVersion(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// State returns the value of property "MIDIPort.state".
//
// It returns ok=false if there is no such property.
func (this MIDIPort) State() (ret MIDIPortDeviceState, ok bool) {
	ok = js.True == bindings.GetMIDIPortState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Connection returns the value of property "MIDIPort.connection".
//
// It returns ok=false if there is no such property.
func (this MIDIPort) Connection() (ret MIDIPortConnectionState, ok bool) {
	ok = js.True == bindings.GetMIDIPortConnection(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasOpen returns true if the method "MIDIPort.open" exists.
func (this MIDIPort) HasOpen() bool {
	return js.True == bindings.HasMIDIPortOpen(
		this.Ref(),
	)
}

// OpenFunc returns the method "MIDIPort.open".
func (this MIDIPort) OpenFunc() (fn js.Func[func() js.Promise[MIDIPort]]) {
	return fn.FromRef(
		bindings.MIDIPortOpenFunc(
			this.Ref(),
		),
	)
}

// Open calls the method "MIDIPort.open".
func (this MIDIPort) Open() (ret js.Promise[MIDIPort]) {
	bindings.CallMIDIPortOpen(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryOpen calls the method "MIDIPort.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MIDIPort) TryOpen() (ret js.Promise[MIDIPort], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMIDIPortOpen(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "MIDIPort.close" exists.
func (this MIDIPort) HasClose() bool {
	return js.True == bindings.HasMIDIPortClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "MIDIPort.close".
func (this MIDIPort) CloseFunc() (fn js.Func[func() js.Promise[MIDIPort]]) {
	return fn.FromRef(
		bindings.MIDIPortCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "MIDIPort.close".
func (this MIDIPort) Close() (ret js.Promise[MIDIPort]) {
	bindings.CallMIDIPortClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "MIDIPort.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MIDIPort) TryClose() (ret js.Promise[MIDIPort], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMIDIPortClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type MIDIConnectionEventInit struct {
	// Port is "MIDIConnectionEventInit.port"
	//
	// Optional
	Port MIDIPort
	// Bubbles is "MIDIConnectionEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "MIDIConnectionEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "MIDIConnectionEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new MIDIConnectionEventInit in the application heap.
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

func NewMIDIConnectionEvent(typ js.String, eventInitDict MIDIConnectionEventInit) (ret MIDIConnectionEvent) {
	ret.ref = bindings.NewMIDIConnectionEventByMIDIConnectionEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewMIDIConnectionEventByMIDIConnectionEvent1(typ js.String) (ret MIDIConnectionEvent) {
	ret.ref = bindings.NewMIDIConnectionEventByMIDIConnectionEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this MIDIConnectionEvent) Port() (ret MIDIPort, ok bool) {
	ok = js.True == bindings.GetMIDIConnectionEventPort(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "MIDIMessageEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "MIDIMessageEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new MIDIMessageEventInit in the application heap.
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

func NewMIDIMessageEvent(typ js.String, eventInitDict MIDIMessageEventInit) (ret MIDIMessageEvent) {
	ret.ref = bindings.NewMIDIMessageEventByMIDIMessageEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewMIDIMessageEventByMIDIMessageEvent1(typ js.String) (ret MIDIMessageEvent) {
	ret.ref = bindings.NewMIDIMessageEventByMIDIMessageEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this MIDIMessageEvent) Data() (ret js.TypedArray[uint8], ok bool) {
	ok = js.True == bindings.GetMIDIMessageEventData(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// HasSend returns true if the method "MIDIOutput.send" exists.
func (this MIDIOutput) HasSend() bool {
	return js.True == bindings.HasMIDIOutputSend(
		this.Ref(),
	)
}

// SendFunc returns the method "MIDIOutput.send".
func (this MIDIOutput) SendFunc() (fn js.Func[func(data js.Array[uint8], timestamp DOMHighResTimeStamp)]) {
	return fn.FromRef(
		bindings.MIDIOutputSendFunc(
			this.Ref(),
		),
	)
}

// Send calls the method "MIDIOutput.send".
func (this MIDIOutput) Send(data js.Array[uint8], timestamp DOMHighResTimeStamp) (ret js.Void) {
	bindings.CallMIDIOutputSend(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
		float64(timestamp),
	)

	return
}

// TrySend calls the method "MIDIOutput.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MIDIOutput) TrySend(data js.Array[uint8], timestamp DOMHighResTimeStamp) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMIDIOutputSend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
		float64(timestamp),
	)

	return
}

// HasSend1 returns true if the method "MIDIOutput.send" exists.
func (this MIDIOutput) HasSend1() bool {
	return js.True == bindings.HasMIDIOutputSend1(
		this.Ref(),
	)
}

// Send1Func returns the method "MIDIOutput.send".
func (this MIDIOutput) Send1Func() (fn js.Func[func(data js.Array[uint8])]) {
	return fn.FromRef(
		bindings.MIDIOutputSend1Func(
			this.Ref(),
		),
	)
}

// Send1 calls the method "MIDIOutput.send".
func (this MIDIOutput) Send1(data js.Array[uint8]) (ret js.Void) {
	bindings.CallMIDIOutputSend1(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySend1 calls the method "MIDIOutput.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MIDIOutput) TrySend1(data js.Array[uint8]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMIDIOutputSend1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasClear returns true if the method "MIDIOutput.clear" exists.
func (this MIDIOutput) HasClear() bool {
	return js.True == bindings.HasMIDIOutputClear(
		this.Ref(),
	)
}

// ClearFunc returns the method "MIDIOutput.clear".
func (this MIDIOutput) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MIDIOutputClearFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "MIDIOutput.clear".
func (this MIDIOutput) Clear() (ret js.Void) {
	bindings.CallMIDIOutputClear(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "MIDIOutput.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MIDIOutput) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMIDIOutputClear(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
	//
	// NOTE: FFI_USE_Axis MUST be set to true to make this field effective.
	Axis uint32
	// Epsilon is "MLBatchNormalizationOptions.epsilon"
	//
	// Optional, defaults to 1e-5.
	//
	// NOTE: FFI_USE_Epsilon MUST be set to true to make this field effective.
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

// New creates a new MLBatchNormalizationOptions in the application heap.
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
	//
	// NOTE: FFI_USE_Offset MUST be set to true to make this field effective.
	Offset uint64
	// Size is "MLBufferResourceView.size"
	//
	// Optional
	//
	// NOTE: FFI_USE_Size MUST be set to true to make this field effective.
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

// New creates a new MLBufferResourceView in the application heap.
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
	//
	// NOTE: FFI_USE_MinValue MUST be set to true to make this field effective.
	MinValue float32
	// MaxValue is "MLClampOptions.maxValue"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxValue MUST be set to true to make this field effective.
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

// New creates a new MLClampOptions in the application heap.
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
	//
	// NOTE: FFI_USE_Groups MUST be set to true to make this field effective.
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

// New creates a new MLConv2dOptions in the application heap.
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
	//
	// NOTE: FFI_USE_Groups MUST be set to true to make this field effective.
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

// New creates a new MLConvTranspose2dOptions in the application heap.
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
	//
	// NOTE: FFI_USE_Alpha MUST be set to true to make this field effective.
	Alpha float32

	FFI_USE_Alpha bool // for Alpha.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLEluOptions with all fields set.
func (p MLEluOptions) FromRef(ref js.Ref) MLEluOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLEluOptions in the application heap.
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
	//
	// NOTE: FFI_USE_Alpha MUST be set to true to make this field effective.
	Alpha float32
	// Beta is "MLGemmOptions.beta"
	//
	// Optional, defaults to 1.0.
	//
	// NOTE: FFI_USE_Beta MUST be set to true to make this field effective.
	Beta float32
	// ATranspose is "MLGemmOptions.aTranspose"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ATranspose MUST be set to true to make this field effective.
	ATranspose bool
	// BTranspose is "MLGemmOptions.bTranspose"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_BTranspose MUST be set to true to make this field effective.
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

// New creates a new MLGemmOptions in the application heap.
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
