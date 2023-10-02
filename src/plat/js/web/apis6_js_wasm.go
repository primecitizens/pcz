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

func NewAmbientLightSensor(sensorOptions SensorOptions) AmbientLightSensor {
	return AmbientLightSensor{}.FromRef(
		bindings.NewAmbientLightSensorByAmbientLightSensor(
			js.Pointer(&sensorOptions)),
	)
}

func NewAmbientLightSensorByAmbientLightSensor1() AmbientLightSensor {
	return AmbientLightSensor{}.FromRef(
		bindings.NewAmbientLightSensorByAmbientLightSensor1(),
	)
}

type AmbientLightSensor struct {
	Sensor
}

func (this AmbientLightSensor) Once() AmbientLightSensor {
	this.Ref().Once()
	return this
}

func (this AmbientLightSensor) Ref() js.Ref {
	return this.Sensor.Ref()
}

func (this AmbientLightSensor) FromRef(ref js.Ref) AmbientLightSensor {
	this.Sensor = this.Sensor.FromRef(ref)
	return this
}

func (this AmbientLightSensor) Free() {
	this.Ref().Free()
}

// Illuminance returns the value of property "AmbientLightSensor.illuminance".
//
// The returned bool will be false if there is no such property.
func (this AmbientLightSensor) Illuminance() (float64, bool) {
	var _ok bool
	_ret := bindings.GetAmbientLightSensorIlluminance(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

type BiquadFilterType uint32

const (
	_ BiquadFilterType = iota

	BiquadFilterType_LOWPASS
	BiquadFilterType_HIGHPASS
	BiquadFilterType_BANDPASS
	BiquadFilterType_LOWSHELF
	BiquadFilterType_HIGHSHELF
	BiquadFilterType_PEAKING
	BiquadFilterType_NOTCH
	BiquadFilterType_ALLPASS
)

func (BiquadFilterType) FromRef(str js.Ref) BiquadFilterType {
	return BiquadFilterType(bindings.ConstOfBiquadFilterType(str))
}

func (x BiquadFilterType) String() (string, bool) {
	switch x {
	case BiquadFilterType_LOWPASS:
		return "lowpass", true
	case BiquadFilterType_HIGHPASS:
		return "highpass", true
	case BiquadFilterType_BANDPASS:
		return "bandpass", true
	case BiquadFilterType_LOWSHELF:
		return "lowshelf", true
	case BiquadFilterType_HIGHSHELF:
		return "highshelf", true
	case BiquadFilterType_PEAKING:
		return "peaking", true
	case BiquadFilterType_NOTCH:
		return "notch", true
	case BiquadFilterType_ALLPASS:
		return "allpass", true
	default:
		return "", false
	}
}

type ChannelCountMode uint32

const (
	_ ChannelCountMode = iota

	ChannelCountMode_MAX
	ChannelCountMode_CLAMPED_MAX
	ChannelCountMode_EXPLICIT
)

func (ChannelCountMode) FromRef(str js.Ref) ChannelCountMode {
	return ChannelCountMode(bindings.ConstOfChannelCountMode(str))
}

func (x ChannelCountMode) String() (string, bool) {
	switch x {
	case ChannelCountMode_MAX:
		return "max", true
	case ChannelCountMode_CLAMPED_MAX:
		return "clamped-max", true
	case ChannelCountMode_EXPLICIT:
		return "explicit", true
	default:
		return "", false
	}
}

type ChannelInterpretation uint32

const (
	_ ChannelInterpretation = iota

	ChannelInterpretation_SPEAKERS
	ChannelInterpretation_DISCRETE
)

func (ChannelInterpretation) FromRef(str js.Ref) ChannelInterpretation {
	return ChannelInterpretation(bindings.ConstOfChannelInterpretation(str))
}

func (x ChannelInterpretation) String() (string, bool) {
	switch x {
	case ChannelInterpretation_SPEAKERS:
		return "speakers", true
	case ChannelInterpretation_DISCRETE:
		return "discrete", true
	default:
		return "", false
	}
}

type BiquadFilterOptions struct {
	// Type is "BiquadFilterOptions.type"
	//
	// Optional, defaults to "lowpass".
	Type BiquadFilterType
	// Q is "BiquadFilterOptions.Q"
	//
	// Optional, defaults to 1.
	Q float32
	// Detune is "BiquadFilterOptions.detune"
	//
	// Optional, defaults to 0.
	Detune float32
	// Frequency is "BiquadFilterOptions.frequency"
	//
	// Optional, defaults to 350.
	Frequency float32
	// Gain is "BiquadFilterOptions.gain"
	//
	// Optional, defaults to 0.
	Gain float32
	// ChannelCount is "BiquadFilterOptions.channelCount"
	//
	// Optional
	ChannelCount uint32
	// ChannelCountMode is "BiquadFilterOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "BiquadFilterOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_Q            bool // for Q.
	FFI_USE_Detune       bool // for Detune.
	FFI_USE_Frequency    bool // for Frequency.
	FFI_USE_Gain         bool // for Gain.
	FFI_USE_ChannelCount bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BiquadFilterOptions with all fields set.
func (p BiquadFilterOptions) FromRef(ref js.Ref) BiquadFilterOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 BiquadFilterOptions BiquadFilterOptions [// BiquadFilterOptions] [0x14000a368c0 0x14000a36960 0x14000a36aa0 0x14000a36be0 0x14000a36d20 0x14000a36e60 0x14000a36fa0 0x14000a37040 0x14000a36a00 0x14000a36b40 0x14000a36c80 0x14000a36dc0 0x14000a36f00] 0x14001af31e8 {0 0}} in the application heap.
func (p BiquadFilterOptions) New() js.Ref {
	return bindings.BiquadFilterOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p BiquadFilterOptions) UpdateFrom(ref js.Ref) {
	bindings.BiquadFilterOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BiquadFilterOptions) Update(ref js.Ref) {
	bindings.BiquadFilterOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AutomationRate uint32

const (
	_ AutomationRate = iota

	AutomationRate_A_RATE
	AutomationRate_K_RATE
)

func (AutomationRate) FromRef(str js.Ref) AutomationRate {
	return AutomationRate(bindings.ConstOfAutomationRate(str))
}

func (x AutomationRate) String() (string, bool) {
	switch x {
	case AutomationRate_A_RATE:
		return "a-rate", true
	case AutomationRate_K_RATE:
		return "k-rate", true
	default:
		return "", false
	}
}

type AudioParam struct {
	ref js.Ref
}

func (this AudioParam) Once() AudioParam {
	this.Ref().Once()
	return this
}

func (this AudioParam) Ref() js.Ref {
	return this.ref
}

func (this AudioParam) FromRef(ref js.Ref) AudioParam {
	this.ref = ref
	return this
}

func (this AudioParam) Free() {
	this.Ref().Free()
}

// Value returns the value of property "AudioParam.value".
//
// The returned bool will be false if there is no such property.
func (this AudioParam) Value() (float32, bool) {
	var _ok bool
	_ret := bindings.GetAudioParamValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Value sets the value of property "AudioParam.value" to val.
//
// It returns false if the property cannot be set.
func (this AudioParam) SetValue(val float32) bool {
	return js.True == bindings.SetAudioParamValue(
		this.Ref(),
		float32(val),
	)
}

// AutomationRate returns the value of property "AudioParam.automationRate".
//
// The returned bool will be false if there is no such property.
func (this AudioParam) AutomationRate() (AutomationRate, bool) {
	var _ok bool
	_ret := bindings.GetAudioParamAutomationRate(
		this.Ref(), js.Pointer(&_ok),
	)
	return AutomationRate(_ret), _ok
}

// AutomationRate sets the value of property "AudioParam.automationRate" to val.
//
// It returns false if the property cannot be set.
func (this AudioParam) SetAutomationRate(val AutomationRate) bool {
	return js.True == bindings.SetAudioParamAutomationRate(
		this.Ref(),
		uint32(val),
	)
}

// DefaultValue returns the value of property "AudioParam.defaultValue".
//
// The returned bool will be false if there is no such property.
func (this AudioParam) DefaultValue() (float32, bool) {
	var _ok bool
	_ret := bindings.GetAudioParamDefaultValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// MinValue returns the value of property "AudioParam.minValue".
//
// The returned bool will be false if there is no such property.
func (this AudioParam) MinValue() (float32, bool) {
	var _ok bool
	_ret := bindings.GetAudioParamMinValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// MaxValue returns the value of property "AudioParam.maxValue".
//
// The returned bool will be false if there is no such property.
func (this AudioParam) MaxValue() (float32, bool) {
	var _ok bool
	_ret := bindings.GetAudioParamMaxValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// SetValueAtTime calls the method "AudioParam.setValueAtTime".
//
// The returned bool will be false if there is no such method.
func (this AudioParam) SetValueAtTime(value float32, startTime float64) (AudioParam, bool) {
	var _ok bool
	_ret := bindings.CallAudioParamSetValueAtTime(
		this.Ref(), js.Pointer(&_ok),
		float32(value),
		float64(startTime),
	)

	return AudioParam{}.FromRef(_ret), _ok
}

// SetValueAtTimeFunc returns the method "AudioParam.setValueAtTime".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioParam) SetValueAtTimeFunc() (fn js.Func[func(value float32, startTime float64) AudioParam]) {
	return fn.FromRef(
		bindings.AudioParamSetValueAtTimeFunc(
			this.Ref(),
		),
	)
}

// LinearRampToValueAtTime calls the method "AudioParam.linearRampToValueAtTime".
//
// The returned bool will be false if there is no such method.
func (this AudioParam) LinearRampToValueAtTime(value float32, endTime float64) (AudioParam, bool) {
	var _ok bool
	_ret := bindings.CallAudioParamLinearRampToValueAtTime(
		this.Ref(), js.Pointer(&_ok),
		float32(value),
		float64(endTime),
	)

	return AudioParam{}.FromRef(_ret), _ok
}

// LinearRampToValueAtTimeFunc returns the method "AudioParam.linearRampToValueAtTime".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioParam) LinearRampToValueAtTimeFunc() (fn js.Func[func(value float32, endTime float64) AudioParam]) {
	return fn.FromRef(
		bindings.AudioParamLinearRampToValueAtTimeFunc(
			this.Ref(),
		),
	)
}

// ExponentialRampToValueAtTime calls the method "AudioParam.exponentialRampToValueAtTime".
//
// The returned bool will be false if there is no such method.
func (this AudioParam) ExponentialRampToValueAtTime(value float32, endTime float64) (AudioParam, bool) {
	var _ok bool
	_ret := bindings.CallAudioParamExponentialRampToValueAtTime(
		this.Ref(), js.Pointer(&_ok),
		float32(value),
		float64(endTime),
	)

	return AudioParam{}.FromRef(_ret), _ok
}

// ExponentialRampToValueAtTimeFunc returns the method "AudioParam.exponentialRampToValueAtTime".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioParam) ExponentialRampToValueAtTimeFunc() (fn js.Func[func(value float32, endTime float64) AudioParam]) {
	return fn.FromRef(
		bindings.AudioParamExponentialRampToValueAtTimeFunc(
			this.Ref(),
		),
	)
}

// SetTargetAtTime calls the method "AudioParam.setTargetAtTime".
//
// The returned bool will be false if there is no such method.
func (this AudioParam) SetTargetAtTime(target float32, startTime float64, timeConstant float32) (AudioParam, bool) {
	var _ok bool
	_ret := bindings.CallAudioParamSetTargetAtTime(
		this.Ref(), js.Pointer(&_ok),
		float32(target),
		float64(startTime),
		float32(timeConstant),
	)

	return AudioParam{}.FromRef(_ret), _ok
}

// SetTargetAtTimeFunc returns the method "AudioParam.setTargetAtTime".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioParam) SetTargetAtTimeFunc() (fn js.Func[func(target float32, startTime float64, timeConstant float32) AudioParam]) {
	return fn.FromRef(
		bindings.AudioParamSetTargetAtTimeFunc(
			this.Ref(),
		),
	)
}

// SetValueCurveAtTime calls the method "AudioParam.setValueCurveAtTime".
//
// The returned bool will be false if there is no such method.
func (this AudioParam) SetValueCurveAtTime(values js.Array[float32], startTime float64, duration float64) (AudioParam, bool) {
	var _ok bool
	_ret := bindings.CallAudioParamSetValueCurveAtTime(
		this.Ref(), js.Pointer(&_ok),
		values.Ref(),
		float64(startTime),
		float64(duration),
	)

	return AudioParam{}.FromRef(_ret), _ok
}

// SetValueCurveAtTimeFunc returns the method "AudioParam.setValueCurveAtTime".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioParam) SetValueCurveAtTimeFunc() (fn js.Func[func(values js.Array[float32], startTime float64, duration float64) AudioParam]) {
	return fn.FromRef(
		bindings.AudioParamSetValueCurveAtTimeFunc(
			this.Ref(),
		),
	)
}

// CancelScheduledValues calls the method "AudioParam.cancelScheduledValues".
//
// The returned bool will be false if there is no such method.
func (this AudioParam) CancelScheduledValues(cancelTime float64) (AudioParam, bool) {
	var _ok bool
	_ret := bindings.CallAudioParamCancelScheduledValues(
		this.Ref(), js.Pointer(&_ok),
		float64(cancelTime),
	)

	return AudioParam{}.FromRef(_ret), _ok
}

// CancelScheduledValuesFunc returns the method "AudioParam.cancelScheduledValues".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioParam) CancelScheduledValuesFunc() (fn js.Func[func(cancelTime float64) AudioParam]) {
	return fn.FromRef(
		bindings.AudioParamCancelScheduledValuesFunc(
			this.Ref(),
		),
	)
}

// CancelAndHoldAtTime calls the method "AudioParam.cancelAndHoldAtTime".
//
// The returned bool will be false if there is no such method.
func (this AudioParam) CancelAndHoldAtTime(cancelTime float64) (AudioParam, bool) {
	var _ok bool
	_ret := bindings.CallAudioParamCancelAndHoldAtTime(
		this.Ref(), js.Pointer(&_ok),
		float64(cancelTime),
	)

	return AudioParam{}.FromRef(_ret), _ok
}

// CancelAndHoldAtTimeFunc returns the method "AudioParam.cancelAndHoldAtTime".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioParam) CancelAndHoldAtTimeFunc() (fn js.Func[func(cancelTime float64) AudioParam]) {
	return fn.FromRef(
		bindings.AudioParamCancelAndHoldAtTimeFunc(
			this.Ref(),
		),
	)
}

func NewBiquadFilterNode(context BaseAudioContext, options BiquadFilterOptions) BiquadFilterNode {
	return BiquadFilterNode{}.FromRef(
		bindings.NewBiquadFilterNodeByBiquadFilterNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewBiquadFilterNodeByBiquadFilterNode1(context BaseAudioContext) BiquadFilterNode {
	return BiquadFilterNode{}.FromRef(
		bindings.NewBiquadFilterNodeByBiquadFilterNode1(
			context.Ref()),
	)
}

type BiquadFilterNode struct {
	AudioNode
}

func (this BiquadFilterNode) Once() BiquadFilterNode {
	this.Ref().Once()
	return this
}

func (this BiquadFilterNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this BiquadFilterNode) FromRef(ref js.Ref) BiquadFilterNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this BiquadFilterNode) Free() {
	this.Ref().Free()
}

// Type returns the value of property "BiquadFilterNode.type".
//
// The returned bool will be false if there is no such property.
func (this BiquadFilterNode) Type() (BiquadFilterType, bool) {
	var _ok bool
	_ret := bindings.GetBiquadFilterNodeType(
		this.Ref(), js.Pointer(&_ok),
	)
	return BiquadFilterType(_ret), _ok
}

// Type sets the value of property "BiquadFilterNode.type" to val.
//
// It returns false if the property cannot be set.
func (this BiquadFilterNode) SetType(val BiquadFilterType) bool {
	return js.True == bindings.SetBiquadFilterNodeType(
		this.Ref(),
		uint32(val),
	)
}

// Frequency returns the value of property "BiquadFilterNode.frequency".
//
// The returned bool will be false if there is no such property.
func (this BiquadFilterNode) Frequency() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetBiquadFilterNodeFrequency(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// Detune returns the value of property "BiquadFilterNode.detune".
//
// The returned bool will be false if there is no such property.
func (this BiquadFilterNode) Detune() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetBiquadFilterNodeDetune(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// Q returns the value of property "BiquadFilterNode.Q".
//
// The returned bool will be false if there is no such property.
func (this BiquadFilterNode) Q() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetBiquadFilterNodeQ(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// Gain returns the value of property "BiquadFilterNode.gain".
//
// The returned bool will be false if there is no such property.
func (this BiquadFilterNode) Gain() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetBiquadFilterNodeGain(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// GetFrequencyResponse calls the method "BiquadFilterNode.getFrequencyResponse".
//
// The returned bool will be false if there is no such method.
func (this BiquadFilterNode) GetFrequencyResponse(frequencyHz js.TypedArray[float32], magResponse js.TypedArray[float32], phaseResponse js.TypedArray[float32]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallBiquadFilterNodeGetFrequencyResponse(
		this.Ref(), js.Pointer(&_ok),
		frequencyHz.Ref(),
		magResponse.Ref(),
		phaseResponse.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetFrequencyResponseFunc returns the method "BiquadFilterNode.getFrequencyResponse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BiquadFilterNode) GetFrequencyResponseFunc() (fn js.Func[func(frequencyHz js.TypedArray[float32], magResponse js.TypedArray[float32], phaseResponse js.TypedArray[float32])]) {
	return fn.FromRef(
		bindings.BiquadFilterNodeGetFrequencyResponseFunc(
			this.Ref(),
		),
	)
}

type AudioBufferOptions struct {
	// NumberOfChannels is "AudioBufferOptions.numberOfChannels"
	//
	// Optional, defaults to 1.
	NumberOfChannels uint32
	// Length is "AudioBufferOptions.length"
	//
	// Required
	Length uint32
	// SampleRate is "AudioBufferOptions.sampleRate"
	//
	// Required
	SampleRate float32

	FFI_USE_NumberOfChannels bool // for NumberOfChannels.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioBufferOptions with all fields set.
func (p AudioBufferOptions) FromRef(ref js.Ref) AudioBufferOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AudioBufferOptions AudioBufferOptions [// AudioBufferOptions] [0x14000a37220 0x14000a37360 0x14000a37400 0x14000a372c0] 0x14001af32c0 {0 0}} in the application heap.
func (p AudioBufferOptions) New() js.Ref {
	return bindings.AudioBufferOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AudioBufferOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioBufferOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioBufferOptions) Update(ref js.Ref) {
	bindings.AudioBufferOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewAudioBuffer(options AudioBufferOptions) AudioBuffer {
	return AudioBuffer{}.FromRef(
		bindings.NewAudioBufferByAudioBuffer(
			js.Pointer(&options)),
	)
}

type AudioBuffer struct {
	ref js.Ref
}

func (this AudioBuffer) Once() AudioBuffer {
	this.Ref().Once()
	return this
}

func (this AudioBuffer) Ref() js.Ref {
	return this.ref
}

func (this AudioBuffer) FromRef(ref js.Ref) AudioBuffer {
	this.ref = ref
	return this
}

func (this AudioBuffer) Free() {
	this.Ref().Free()
}

// SampleRate returns the value of property "AudioBuffer.sampleRate".
//
// The returned bool will be false if there is no such property.
func (this AudioBuffer) SampleRate() (float32, bool) {
	var _ok bool
	_ret := bindings.GetAudioBufferSampleRate(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Length returns the value of property "AudioBuffer.length".
//
// The returned bool will be false if there is no such property.
func (this AudioBuffer) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetAudioBufferLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Duration returns the value of property "AudioBuffer.duration".
//
// The returned bool will be false if there is no such property.
func (this AudioBuffer) Duration() (float64, bool) {
	var _ok bool
	_ret := bindings.GetAudioBufferDuration(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// NumberOfChannels returns the value of property "AudioBuffer.numberOfChannels".
//
// The returned bool will be false if there is no such property.
func (this AudioBuffer) NumberOfChannels() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetAudioBufferNumberOfChannels(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// GetChannelData calls the method "AudioBuffer.getChannelData".
//
// The returned bool will be false if there is no such method.
func (this AudioBuffer) GetChannelData(channel uint32) (js.TypedArray[float32], bool) {
	var _ok bool
	_ret := bindings.CallAudioBufferGetChannelData(
		this.Ref(), js.Pointer(&_ok),
		uint32(channel),
	)

	return js.TypedArray[float32]{}.FromRef(_ret), _ok
}

// GetChannelDataFunc returns the method "AudioBuffer.getChannelData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioBuffer) GetChannelDataFunc() (fn js.Func[func(channel uint32) js.TypedArray[float32]]) {
	return fn.FromRef(
		bindings.AudioBufferGetChannelDataFunc(
			this.Ref(),
		),
	)
}

// CopyFromChannel calls the method "AudioBuffer.copyFromChannel".
//
// The returned bool will be false if there is no such method.
func (this AudioBuffer) CopyFromChannel(destination js.TypedArray[float32], channelNumber uint32, bufferOffset uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioBufferCopyFromChannel(
		this.Ref(), js.Pointer(&_ok),
		destination.Ref(),
		uint32(channelNumber),
		uint32(bufferOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyFromChannelFunc returns the method "AudioBuffer.copyFromChannel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioBuffer) CopyFromChannelFunc() (fn js.Func[func(destination js.TypedArray[float32], channelNumber uint32, bufferOffset uint32)]) {
	return fn.FromRef(
		bindings.AudioBufferCopyFromChannelFunc(
			this.Ref(),
		),
	)
}

// CopyFromChannel1 calls the method "AudioBuffer.copyFromChannel".
//
// The returned bool will be false if there is no such method.
func (this AudioBuffer) CopyFromChannel1(destination js.TypedArray[float32], channelNumber uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioBufferCopyFromChannel1(
		this.Ref(), js.Pointer(&_ok),
		destination.Ref(),
		uint32(channelNumber),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyFromChannel1Func returns the method "AudioBuffer.copyFromChannel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioBuffer) CopyFromChannel1Func() (fn js.Func[func(destination js.TypedArray[float32], channelNumber uint32)]) {
	return fn.FromRef(
		bindings.AudioBufferCopyFromChannel1Func(
			this.Ref(),
		),
	)
}

// CopyToChannel calls the method "AudioBuffer.copyToChannel".
//
// The returned bool will be false if there is no such method.
func (this AudioBuffer) CopyToChannel(source js.TypedArray[float32], channelNumber uint32, bufferOffset uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioBufferCopyToChannel(
		this.Ref(), js.Pointer(&_ok),
		source.Ref(),
		uint32(channelNumber),
		uint32(bufferOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyToChannelFunc returns the method "AudioBuffer.copyToChannel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioBuffer) CopyToChannelFunc() (fn js.Func[func(source js.TypedArray[float32], channelNumber uint32, bufferOffset uint32)]) {
	return fn.FromRef(
		bindings.AudioBufferCopyToChannelFunc(
			this.Ref(),
		),
	)
}

// CopyToChannel1 calls the method "AudioBuffer.copyToChannel".
//
// The returned bool will be false if there is no such method.
func (this AudioBuffer) CopyToChannel1(source js.TypedArray[float32], channelNumber uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioBufferCopyToChannel1(
		this.Ref(), js.Pointer(&_ok),
		source.Ref(),
		uint32(channelNumber),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyToChannel1Func returns the method "AudioBuffer.copyToChannel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioBuffer) CopyToChannel1Func() (fn js.Func[func(source js.TypedArray[float32], channelNumber uint32)]) {
	return fn.FromRef(
		bindings.AudioBufferCopyToChannel1Func(
			this.Ref(),
		),
	)
}

type AudioBufferSourceOptions struct {
	// Buffer is "AudioBufferSourceOptions.buffer"
	//
	// Optional
	Buffer AudioBuffer
	// Detune is "AudioBufferSourceOptions.detune"
	//
	// Optional, defaults to 0.
	Detune float32
	// Loop is "AudioBufferSourceOptions.loop"
	//
	// Optional, defaults to false.
	Loop bool
	// LoopEnd is "AudioBufferSourceOptions.loopEnd"
	//
	// Optional, defaults to 0.
	LoopEnd float64
	// LoopStart is "AudioBufferSourceOptions.loopStart"
	//
	// Optional, defaults to 0.
	LoopStart float64
	// PlaybackRate is "AudioBufferSourceOptions.playbackRate"
	//
	// Optional, defaults to 1.
	PlaybackRate float32

	FFI_USE_Detune       bool // for Detune.
	FFI_USE_Loop         bool // for Loop.
	FFI_USE_LoopEnd      bool // for LoopEnd.
	FFI_USE_LoopStart    bool // for LoopStart.
	FFI_USE_PlaybackRate bool // for PlaybackRate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioBufferSourceOptions with all fields set.
func (p AudioBufferSourceOptions) FromRef(ref js.Ref) AudioBufferSourceOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AudioBufferSourceOptions AudioBufferSourceOptions [// AudioBufferSourceOptions] [0x14000a375e0 0x14000a37680 0x14000a377c0 0x14000a37900 0x14000a37a40 0x14000a37b80 0x14000a37720 0x14000a37860 0x14000a379a0 0x14000a37ae0 0x14000a37c20] 0x140010ae0d8 {0 0}} in the application heap.
func (p AudioBufferSourceOptions) New() js.Ref {
	return bindings.AudioBufferSourceOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AudioBufferSourceOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioBufferSourceOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioBufferSourceOptions) Update(ref js.Ref) {
	bindings.AudioBufferSourceOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewAudioBufferSourceNode(context BaseAudioContext, options AudioBufferSourceOptions) AudioBufferSourceNode {
	return AudioBufferSourceNode{}.FromRef(
		bindings.NewAudioBufferSourceNodeByAudioBufferSourceNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewAudioBufferSourceNodeByAudioBufferSourceNode1(context BaseAudioContext) AudioBufferSourceNode {
	return AudioBufferSourceNode{}.FromRef(
		bindings.NewAudioBufferSourceNodeByAudioBufferSourceNode1(
			context.Ref()),
	)
}

type AudioBufferSourceNode struct {
	AudioScheduledSourceNode
}

func (this AudioBufferSourceNode) Once() AudioBufferSourceNode {
	this.Ref().Once()
	return this
}

func (this AudioBufferSourceNode) Ref() js.Ref {
	return this.AudioScheduledSourceNode.Ref()
}

func (this AudioBufferSourceNode) FromRef(ref js.Ref) AudioBufferSourceNode {
	this.AudioScheduledSourceNode = this.AudioScheduledSourceNode.FromRef(ref)
	return this
}

func (this AudioBufferSourceNode) Free() {
	this.Ref().Free()
}

// Buffer returns the value of property "AudioBufferSourceNode.buffer".
//
// The returned bool will be false if there is no such property.
func (this AudioBufferSourceNode) Buffer() (AudioBuffer, bool) {
	var _ok bool
	_ret := bindings.GetAudioBufferSourceNodeBuffer(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioBuffer{}.FromRef(_ret), _ok
}

// Buffer sets the value of property "AudioBufferSourceNode.buffer" to val.
//
// It returns false if the property cannot be set.
func (this AudioBufferSourceNode) SetBuffer(val AudioBuffer) bool {
	return js.True == bindings.SetAudioBufferSourceNodeBuffer(
		this.Ref(),
		val.Ref(),
	)
}

// PlaybackRate returns the value of property "AudioBufferSourceNode.playbackRate".
//
// The returned bool will be false if there is no such property.
func (this AudioBufferSourceNode) PlaybackRate() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetAudioBufferSourceNodePlaybackRate(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// Detune returns the value of property "AudioBufferSourceNode.detune".
//
// The returned bool will be false if there is no such property.
func (this AudioBufferSourceNode) Detune() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetAudioBufferSourceNodeDetune(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// Loop returns the value of property "AudioBufferSourceNode.loop".
//
// The returned bool will be false if there is no such property.
func (this AudioBufferSourceNode) Loop() (bool, bool) {
	var _ok bool
	_ret := bindings.GetAudioBufferSourceNodeLoop(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Loop sets the value of property "AudioBufferSourceNode.loop" to val.
//
// It returns false if the property cannot be set.
func (this AudioBufferSourceNode) SetLoop(val bool) bool {
	return js.True == bindings.SetAudioBufferSourceNodeLoop(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// LoopStart returns the value of property "AudioBufferSourceNode.loopStart".
//
// The returned bool will be false if there is no such property.
func (this AudioBufferSourceNode) LoopStart() (float64, bool) {
	var _ok bool
	_ret := bindings.GetAudioBufferSourceNodeLoopStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// LoopStart sets the value of property "AudioBufferSourceNode.loopStart" to val.
//
// It returns false if the property cannot be set.
func (this AudioBufferSourceNode) SetLoopStart(val float64) bool {
	return js.True == bindings.SetAudioBufferSourceNodeLoopStart(
		this.Ref(),
		float64(val),
	)
}

// LoopEnd returns the value of property "AudioBufferSourceNode.loopEnd".
//
// The returned bool will be false if there is no such property.
func (this AudioBufferSourceNode) LoopEnd() (float64, bool) {
	var _ok bool
	_ret := bindings.GetAudioBufferSourceNodeLoopEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// LoopEnd sets the value of property "AudioBufferSourceNode.loopEnd" to val.
//
// It returns false if the property cannot be set.
func (this AudioBufferSourceNode) SetLoopEnd(val float64) bool {
	return js.True == bindings.SetAudioBufferSourceNodeLoopEnd(
		this.Ref(),
		float64(val),
	)
}

// Start calls the method "AudioBufferSourceNode.start".
//
// The returned bool will be false if there is no such method.
func (this AudioBufferSourceNode) Start(when float64, offset float64, duration float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioBufferSourceNodeStart(
		this.Ref(), js.Pointer(&_ok),
		float64(when),
		float64(offset),
		float64(duration),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StartFunc returns the method "AudioBufferSourceNode.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioBufferSourceNode) StartFunc() (fn js.Func[func(when float64, offset float64, duration float64)]) {
	return fn.FromRef(
		bindings.AudioBufferSourceNodeStartFunc(
			this.Ref(),
		),
	)
}

// Start1 calls the method "AudioBufferSourceNode.start".
//
// The returned bool will be false if there is no such method.
func (this AudioBufferSourceNode) Start1(when float64, offset float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioBufferSourceNodeStart1(
		this.Ref(), js.Pointer(&_ok),
		float64(when),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Start1Func returns the method "AudioBufferSourceNode.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioBufferSourceNode) Start1Func() (fn js.Func[func(when float64, offset float64)]) {
	return fn.FromRef(
		bindings.AudioBufferSourceNodeStart1Func(
			this.Ref(),
		),
	)
}

// Start2 calls the method "AudioBufferSourceNode.start".
//
// The returned bool will be false if there is no such method.
func (this AudioBufferSourceNode) Start2(when float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioBufferSourceNodeStart2(
		this.Ref(), js.Pointer(&_ok),
		float64(when),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Start2Func returns the method "AudioBufferSourceNode.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioBufferSourceNode) Start2Func() (fn js.Func[func(when float64)]) {
	return fn.FromRef(
		bindings.AudioBufferSourceNodeStart2Func(
			this.Ref(),
		),
	)
}

// Start3 calls the method "AudioBufferSourceNode.start".
//
// The returned bool will be false if there is no such method.
func (this AudioBufferSourceNode) Start3() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioBufferSourceNodeStart3(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Start3Func returns the method "AudioBufferSourceNode.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioBufferSourceNode) Start3Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioBufferSourceNodeStart3Func(
			this.Ref(),
		),
	)
}

type ChannelMergerOptions struct {
	// NumberOfInputs is "ChannelMergerOptions.numberOfInputs"
	//
	// Optional, defaults to 6.
	NumberOfInputs uint32
	// ChannelCount is "ChannelMergerOptions.channelCount"
	//
	// Optional
	ChannelCount uint32
	// ChannelCountMode is "ChannelMergerOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "ChannelMergerOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_NumberOfInputs bool // for NumberOfInputs.
	FFI_USE_ChannelCount   bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ChannelMergerOptions with all fields set.
func (p ChannelMergerOptions) FromRef(ref js.Ref) ChannelMergerOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ChannelMergerOptions ChannelMergerOptions [// ChannelMergerOptions] [0x14000a37d60 0x14000a37ea0 0x140014ba000 0x140014ba0a0 0x14000a37e00 0x14000a37f40] 0x140010ae1f8 {0 0}} in the application heap.
func (p ChannelMergerOptions) New() js.Ref {
	return bindings.ChannelMergerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ChannelMergerOptions) UpdateFrom(ref js.Ref) {
	bindings.ChannelMergerOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ChannelMergerOptions) Update(ref js.Ref) {
	bindings.ChannelMergerOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewChannelMergerNode(context BaseAudioContext, options ChannelMergerOptions) ChannelMergerNode {
	return ChannelMergerNode{}.FromRef(
		bindings.NewChannelMergerNodeByChannelMergerNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewChannelMergerNodeByChannelMergerNode1(context BaseAudioContext) ChannelMergerNode {
	return ChannelMergerNode{}.FromRef(
		bindings.NewChannelMergerNodeByChannelMergerNode1(
			context.Ref()),
	)
}

type ChannelMergerNode struct {
	AudioNode
}

func (this ChannelMergerNode) Once() ChannelMergerNode {
	this.Ref().Once()
	return this
}

func (this ChannelMergerNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this ChannelMergerNode) FromRef(ref js.Ref) ChannelMergerNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this ChannelMergerNode) Free() {
	this.Ref().Free()
}

type ChannelSplitterOptions struct {
	// NumberOfOutputs is "ChannelSplitterOptions.numberOfOutputs"
	//
	// Optional, defaults to 6.
	NumberOfOutputs uint32
	// ChannelCount is "ChannelSplitterOptions.channelCount"
	//
	// Optional
	ChannelCount uint32
	// ChannelCountMode is "ChannelSplitterOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "ChannelSplitterOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_NumberOfOutputs bool // for NumberOfOutputs.
	FFI_USE_ChannelCount    bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ChannelSplitterOptions with all fields set.
func (p ChannelSplitterOptions) FromRef(ref js.Ref) ChannelSplitterOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ChannelSplitterOptions ChannelSplitterOptions [// ChannelSplitterOptions] [0x140014ba140 0x140014ba280 0x140014ba3c0 0x140014ba460 0x140014ba1e0 0x140014ba320] 0x140010aeb28 {0 0}} in the application heap.
func (p ChannelSplitterOptions) New() js.Ref {
	return bindings.ChannelSplitterOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ChannelSplitterOptions) UpdateFrom(ref js.Ref) {
	bindings.ChannelSplitterOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ChannelSplitterOptions) Update(ref js.Ref) {
	bindings.ChannelSplitterOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewChannelSplitterNode(context BaseAudioContext, options ChannelSplitterOptions) ChannelSplitterNode {
	return ChannelSplitterNode{}.FromRef(
		bindings.NewChannelSplitterNodeByChannelSplitterNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewChannelSplitterNodeByChannelSplitterNode1(context BaseAudioContext) ChannelSplitterNode {
	return ChannelSplitterNode{}.FromRef(
		bindings.NewChannelSplitterNodeByChannelSplitterNode1(
			context.Ref()),
	)
}

type ChannelSplitterNode struct {
	AudioNode
}

func (this ChannelSplitterNode) Once() ChannelSplitterNode {
	this.Ref().Once()
	return this
}

func (this ChannelSplitterNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this ChannelSplitterNode) FromRef(ref js.Ref) ChannelSplitterNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this ChannelSplitterNode) Free() {
	this.Ref().Free()
}

type ConstantSourceOptions struct {
	// Offset is "ConstantSourceOptions.offset"
	//
	// Optional, defaults to 1.
	Offset float32

	FFI_USE_Offset bool // for Offset.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConstantSourceOptions with all fields set.
func (p ConstantSourceOptions) FromRef(ref js.Ref) ConstantSourceOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ConstantSourceOptions ConstantSourceOptions [// ConstantSourceOptions] [0x140014ba500 0x140014ba5a0] 0x140010af0b0 {0 0}} in the application heap.
func (p ConstantSourceOptions) New() js.Ref {
	return bindings.ConstantSourceOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ConstantSourceOptions) UpdateFrom(ref js.Ref) {
	bindings.ConstantSourceOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ConstantSourceOptions) Update(ref js.Ref) {
	bindings.ConstantSourceOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewConstantSourceNode(context BaseAudioContext, options ConstantSourceOptions) ConstantSourceNode {
	return ConstantSourceNode{}.FromRef(
		bindings.NewConstantSourceNodeByConstantSourceNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewConstantSourceNodeByConstantSourceNode1(context BaseAudioContext) ConstantSourceNode {
	return ConstantSourceNode{}.FromRef(
		bindings.NewConstantSourceNodeByConstantSourceNode1(
			context.Ref()),
	)
}

type ConstantSourceNode struct {
	AudioScheduledSourceNode
}

func (this ConstantSourceNode) Once() ConstantSourceNode {
	this.Ref().Once()
	return this
}

func (this ConstantSourceNode) Ref() js.Ref {
	return this.AudioScheduledSourceNode.Ref()
}

func (this ConstantSourceNode) FromRef(ref js.Ref) ConstantSourceNode {
	this.AudioScheduledSourceNode = this.AudioScheduledSourceNode.FromRef(ref)
	return this
}

func (this ConstantSourceNode) Free() {
	this.Ref().Free()
}

// Offset returns the value of property "ConstantSourceNode.offset".
//
// The returned bool will be false if there is no such property.
func (this ConstantSourceNode) Offset() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetConstantSourceNodeOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

type ConvolverOptions struct {
	// Buffer is "ConvolverOptions.buffer"
	//
	// Optional
	Buffer AudioBuffer
	// DisableNormalization is "ConvolverOptions.disableNormalization"
	//
	// Optional, defaults to false.
	DisableNormalization bool
	// ChannelCount is "ConvolverOptions.channelCount"
	//
	// Optional
	ChannelCount uint32
	// ChannelCountMode is "ConvolverOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "ConvolverOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_DisableNormalization bool // for DisableNormalization.
	FFI_USE_ChannelCount         bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConvolverOptions with all fields set.
func (p ConvolverOptions) FromRef(ref js.Ref) ConvolverOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ConvolverOptions ConvolverOptions [// ConvolverOptions] [0x140014ba640 0x140014ba6e0 0x140014ba820 0x140014ba960 0x140014baa00 0x140014ba780 0x140014ba8c0] 0x140010af110 {0 0}} in the application heap.
func (p ConvolverOptions) New() js.Ref {
	return bindings.ConvolverOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ConvolverOptions) UpdateFrom(ref js.Ref) {
	bindings.ConvolverOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ConvolverOptions) Update(ref js.Ref) {
	bindings.ConvolverOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewConvolverNode(context BaseAudioContext, options ConvolverOptions) ConvolverNode {
	return ConvolverNode{}.FromRef(
		bindings.NewConvolverNodeByConvolverNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewConvolverNodeByConvolverNode1(context BaseAudioContext) ConvolverNode {
	return ConvolverNode{}.FromRef(
		bindings.NewConvolverNodeByConvolverNode1(
			context.Ref()),
	)
}

type ConvolverNode struct {
	AudioNode
}

func (this ConvolverNode) Once() ConvolverNode {
	this.Ref().Once()
	return this
}

func (this ConvolverNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this ConvolverNode) FromRef(ref js.Ref) ConvolverNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this ConvolverNode) Free() {
	this.Ref().Free()
}

// Buffer returns the value of property "ConvolverNode.buffer".
//
// The returned bool will be false if there is no such property.
func (this ConvolverNode) Buffer() (AudioBuffer, bool) {
	var _ok bool
	_ret := bindings.GetConvolverNodeBuffer(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioBuffer{}.FromRef(_ret), _ok
}

// Buffer sets the value of property "ConvolverNode.buffer" to val.
//
// It returns false if the property cannot be set.
func (this ConvolverNode) SetBuffer(val AudioBuffer) bool {
	return js.True == bindings.SetConvolverNodeBuffer(
		this.Ref(),
		val.Ref(),
	)
}

// Normalize returns the value of property "ConvolverNode.normalize".
//
// The returned bool will be false if there is no such property.
func (this ConvolverNode) Normalize() (bool, bool) {
	var _ok bool
	_ret := bindings.GetConvolverNodeNormalize(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Normalize sets the value of property "ConvolverNode.normalize" to val.
//
// It returns false if the property cannot be set.
func (this ConvolverNode) SetNormalize(val bool) bool {
	return js.True == bindings.SetConvolverNodeNormalize(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

type DelayOptions struct {
	// MaxDelayTime is "DelayOptions.maxDelayTime"
	//
	// Optional, defaults to 1.
	MaxDelayTime float64
	// DelayTime is "DelayOptions.delayTime"
	//
	// Optional, defaults to 0.
	DelayTime float64
	// ChannelCount is "DelayOptions.channelCount"
	//
	// Optional
	ChannelCount uint32
	// ChannelCountMode is "DelayOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "DelayOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_MaxDelayTime bool // for MaxDelayTime.
	FFI_USE_DelayTime    bool // for DelayTime.
	FFI_USE_ChannelCount bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DelayOptions with all fields set.
func (p DelayOptions) FromRef(ref js.Ref) DelayOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 DelayOptions DelayOptions [// DelayOptions] [0x140014bab40 0x140014bac80 0x140014badc0 0x140014baf00 0x140014bafa0 0x140014babe0 0x140014bad20 0x140014bae60] 0x140010af170 {0 0}} in the application heap.
func (p DelayOptions) New() js.Ref {
	return bindings.DelayOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DelayOptions) UpdateFrom(ref js.Ref) {
	bindings.DelayOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DelayOptions) Update(ref js.Ref) {
	bindings.DelayOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewDelayNode(context BaseAudioContext, options DelayOptions) DelayNode {
	return DelayNode{}.FromRef(
		bindings.NewDelayNodeByDelayNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewDelayNodeByDelayNode1(context BaseAudioContext) DelayNode {
	return DelayNode{}.FromRef(
		bindings.NewDelayNodeByDelayNode1(
			context.Ref()),
	)
}

type DelayNode struct {
	AudioNode
}

func (this DelayNode) Once() DelayNode {
	this.Ref().Once()
	return this
}

func (this DelayNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this DelayNode) FromRef(ref js.Ref) DelayNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this DelayNode) Free() {
	this.Ref().Free()
}

// DelayTime returns the value of property "DelayNode.delayTime".
//
// The returned bool will be false if there is no such property.
func (this DelayNode) DelayTime() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetDelayNodeDelayTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

type DynamicsCompressorOptions struct {
	// Attack is "DynamicsCompressorOptions.attack"
	//
	// Optional, defaults to 0.003.
	Attack float32
	// Knee is "DynamicsCompressorOptions.knee"
	//
	// Optional, defaults to 30.
	Knee float32
	// Ratio is "DynamicsCompressorOptions.ratio"
	//
	// Optional, defaults to 12.
	Ratio float32
	// Release is "DynamicsCompressorOptions.release"
	//
	// Optional, defaults to 0.25.
	Release float32
	// Threshold is "DynamicsCompressorOptions.threshold"
	//
	// Optional, defaults to -24.
	Threshold float32
	// ChannelCount is "DynamicsCompressorOptions.channelCount"
	//
	// Optional
	ChannelCount uint32
	// ChannelCountMode is "DynamicsCompressorOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "DynamicsCompressorOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_Attack       bool // for Attack.
	FFI_USE_Knee         bool // for Knee.
	FFI_USE_Ratio        bool // for Ratio.
	FFI_USE_Release      bool // for Release.
	FFI_USE_Threshold    bool // for Threshold.
	FFI_USE_ChannelCount bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DynamicsCompressorOptions with all fields set.
func (p DynamicsCompressorOptions) FromRef(ref js.Ref) DynamicsCompressorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 DynamicsCompressorOptions DynamicsCompressorOptions [// DynamicsCompressorOptions] [0x140014bb040 0x140014bb180 0x140014bb2c0 0x140014bb400 0x140014bb540 0x140014bb680 0x140014bb7c0 0x140014bb860 0x140014bb0e0 0x140014bb220 0x140014bb360 0x140014bb4a0 0x140014bb5e0 0x140014bb720] 0x140010af2f0 {0 0}} in the application heap.
func (p DynamicsCompressorOptions) New() js.Ref {
	return bindings.DynamicsCompressorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DynamicsCompressorOptions) UpdateFrom(ref js.Ref) {
	bindings.DynamicsCompressorOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DynamicsCompressorOptions) Update(ref js.Ref) {
	bindings.DynamicsCompressorOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewDynamicsCompressorNode(context BaseAudioContext, options DynamicsCompressorOptions) DynamicsCompressorNode {
	return DynamicsCompressorNode{}.FromRef(
		bindings.NewDynamicsCompressorNodeByDynamicsCompressorNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewDynamicsCompressorNodeByDynamicsCompressorNode1(context BaseAudioContext) DynamicsCompressorNode {
	return DynamicsCompressorNode{}.FromRef(
		bindings.NewDynamicsCompressorNodeByDynamicsCompressorNode1(
			context.Ref()),
	)
}

type DynamicsCompressorNode struct {
	AudioNode
}

func (this DynamicsCompressorNode) Once() DynamicsCompressorNode {
	this.Ref().Once()
	return this
}

func (this DynamicsCompressorNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this DynamicsCompressorNode) FromRef(ref js.Ref) DynamicsCompressorNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this DynamicsCompressorNode) Free() {
	this.Ref().Free()
}

// Threshold returns the value of property "DynamicsCompressorNode.threshold".
//
// The returned bool will be false if there is no such property.
func (this DynamicsCompressorNode) Threshold() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetDynamicsCompressorNodeThreshold(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// Knee returns the value of property "DynamicsCompressorNode.knee".
//
// The returned bool will be false if there is no such property.
func (this DynamicsCompressorNode) Knee() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetDynamicsCompressorNodeKnee(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// Ratio returns the value of property "DynamicsCompressorNode.ratio".
//
// The returned bool will be false if there is no such property.
func (this DynamicsCompressorNode) Ratio() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetDynamicsCompressorNodeRatio(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// Reduction returns the value of property "DynamicsCompressorNode.reduction".
//
// The returned bool will be false if there is no such property.
func (this DynamicsCompressorNode) Reduction() (float32, bool) {
	var _ok bool
	_ret := bindings.GetDynamicsCompressorNodeReduction(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Attack returns the value of property "DynamicsCompressorNode.attack".
//
// The returned bool will be false if there is no such property.
func (this DynamicsCompressorNode) Attack() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetDynamicsCompressorNodeAttack(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// Release returns the value of property "DynamicsCompressorNode.release".
//
// The returned bool will be false if there is no such property.
func (this DynamicsCompressorNode) Release() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetDynamicsCompressorNodeRelease(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

type GainOptions struct {
	// Gain is "GainOptions.gain"
	//
	// Optional, defaults to 1.0.
	Gain float32
	// ChannelCount is "GainOptions.channelCount"
	//
	// Optional
	ChannelCount uint32
	// ChannelCountMode is "GainOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "GainOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_Gain         bool // for Gain.
	FFI_USE_ChannelCount bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GainOptions with all fields set.
func (p GainOptions) FromRef(ref js.Ref) GainOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GainOptions GainOptions [// GainOptions] [0x140014bb9a0 0x140014bbae0 0x140014bbc20 0x140014bbd60 0x140014bba40 0x140014bbb80] 0x140010af380 {0 0}} in the application heap.
func (p GainOptions) New() js.Ref {
	return bindings.GainOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GainOptions) UpdateFrom(ref js.Ref) {
	bindings.GainOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GainOptions) Update(ref js.Ref) {
	bindings.GainOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewGainNode(context BaseAudioContext, options GainOptions) GainNode {
	return GainNode{}.FromRef(
		bindings.NewGainNodeByGainNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewGainNodeByGainNode1(context BaseAudioContext) GainNode {
	return GainNode{}.FromRef(
		bindings.NewGainNodeByGainNode1(
			context.Ref()),
	)
}

type GainNode struct {
	AudioNode
}

func (this GainNode) Once() GainNode {
	this.Ref().Once()
	return this
}

func (this GainNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this GainNode) FromRef(ref js.Ref) GainNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this GainNode) Free() {
	this.Ref().Free()
}

// Gain returns the value of property "GainNode.gain".
//
// The returned bool will be false if there is no such property.
func (this GainNode) Gain() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetGainNodeGain(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

type IIRFilterOptions struct {
	// Feedforward is "IIRFilterOptions.feedforward"
	//
	// Required
	Feedforward js.Array[float64]
	// Feedback is "IIRFilterOptions.feedback"
	//
	// Required
	Feedback js.Array[float64]
	// ChannelCount is "IIRFilterOptions.channelCount"
	//
	// Optional
	ChannelCount uint32
	// ChannelCountMode is "IIRFilterOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "IIRFilterOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_ChannelCount bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IIRFilterOptions with all fields set.
func (p IIRFilterOptions) FromRef(ref js.Ref) IIRFilterOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IIRFilterOptions IIRFilterOptions [// IIRFilterOptions] [0x140014bbe00 0x140014bbea0 0x140014bbf40 0x140000fc0a0 0x140000fc140 0x140000fc000] 0x140010af3e0 {0 0}} in the application heap.
func (p IIRFilterOptions) New() js.Ref {
	return bindings.IIRFilterOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IIRFilterOptions) UpdateFrom(ref js.Ref) {
	bindings.IIRFilterOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IIRFilterOptions) Update(ref js.Ref) {
	bindings.IIRFilterOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewIIRFilterNode(context BaseAudioContext, options IIRFilterOptions) IIRFilterNode {
	return IIRFilterNode{}.FromRef(
		bindings.NewIIRFilterNodeByIIRFilterNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

type IIRFilterNode struct {
	AudioNode
}

func (this IIRFilterNode) Once() IIRFilterNode {
	this.Ref().Once()
	return this
}

func (this IIRFilterNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this IIRFilterNode) FromRef(ref js.Ref) IIRFilterNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this IIRFilterNode) Free() {
	this.Ref().Free()
}

// GetFrequencyResponse calls the method "IIRFilterNode.getFrequencyResponse".
//
// The returned bool will be false if there is no such method.
func (this IIRFilterNode) GetFrequencyResponse(frequencyHz js.TypedArray[float32], magResponse js.TypedArray[float32], phaseResponse js.TypedArray[float32]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallIIRFilterNodeGetFrequencyResponse(
		this.Ref(), js.Pointer(&_ok),
		frequencyHz.Ref(),
		magResponse.Ref(),
		phaseResponse.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetFrequencyResponseFunc returns the method "IIRFilterNode.getFrequencyResponse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IIRFilterNode) GetFrequencyResponseFunc() (fn js.Func[func(frequencyHz js.TypedArray[float32], magResponse js.TypedArray[float32], phaseResponse js.TypedArray[float32])]) {
	return fn.FromRef(
		bindings.IIRFilterNodeGetFrequencyResponseFunc(
			this.Ref(),
		),
	)
}

type OscillatorType uint32

const (
	_ OscillatorType = iota

	OscillatorType_SINE
	OscillatorType_SQUARE
	OscillatorType_SAWTOOTH
	OscillatorType_TRIANGLE
	OscillatorType_CUSTOM
)

func (OscillatorType) FromRef(str js.Ref) OscillatorType {
	return OscillatorType(bindings.ConstOfOscillatorType(str))
}

func (x OscillatorType) String() (string, bool) {
	switch x {
	case OscillatorType_SINE:
		return "sine", true
	case OscillatorType_SQUARE:
		return "square", true
	case OscillatorType_SAWTOOTH:
		return "sawtooth", true
	case OscillatorType_TRIANGLE:
		return "triangle", true
	case OscillatorType_CUSTOM:
		return "custom", true
	default:
		return "", false
	}
}

type PeriodicWaveOptions struct {
	// Real is "PeriodicWaveOptions.real"
	//
	// Optional
	Real js.Array[float32]
	// Imag is "PeriodicWaveOptions.imag"
	//
	// Optional
	Imag js.Array[float32]
	// DisableNormalization is "PeriodicWaveOptions.disableNormalization"
	//
	// Optional, defaults to false.
	DisableNormalization bool

	FFI_USE_DisableNormalization bool // for DisableNormalization.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PeriodicWaveOptions with all fields set.
func (p PeriodicWaveOptions) FromRef(ref js.Ref) PeriodicWaveOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PeriodicWaveOptions PeriodicWaveOptions [// PeriodicWaveOptions] [0x140000fc500 0x140000fc5a0 0x140000fc640 0x140000fc6e0] 0x140010af4b8 {0 0}} in the application heap.
func (p PeriodicWaveOptions) New() js.Ref {
	return bindings.PeriodicWaveOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PeriodicWaveOptions) UpdateFrom(ref js.Ref) {
	bindings.PeriodicWaveOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PeriodicWaveOptions) Update(ref js.Ref) {
	bindings.PeriodicWaveOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPeriodicWave(context BaseAudioContext, options PeriodicWaveOptions) PeriodicWave {
	return PeriodicWave{}.FromRef(
		bindings.NewPeriodicWaveByPeriodicWave(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewPeriodicWaveByPeriodicWave1(context BaseAudioContext) PeriodicWave {
	return PeriodicWave{}.FromRef(
		bindings.NewPeriodicWaveByPeriodicWave1(
			context.Ref()),
	)
}

type PeriodicWave struct {
	ref js.Ref
}

func (this PeriodicWave) Once() PeriodicWave {
	this.Ref().Once()
	return this
}

func (this PeriodicWave) Ref() js.Ref {
	return this.ref
}

func (this PeriodicWave) FromRef(ref js.Ref) PeriodicWave {
	this.ref = ref
	return this
}

func (this PeriodicWave) Free() {
	this.Ref().Free()
}

type OscillatorOptions struct {
	// Type is "OscillatorOptions.type"
	//
	// Optional, defaults to "sine".
	Type OscillatorType
	// Frequency is "OscillatorOptions.frequency"
	//
	// Optional, defaults to 440.
	Frequency float32
	// Detune is "OscillatorOptions.detune"
	//
	// Optional, defaults to 0.
	Detune float32
	// PeriodicWave is "OscillatorOptions.periodicWave"
	//
	// Optional
	PeriodicWave PeriodicWave
	// ChannelCount is "OscillatorOptions.channelCount"
	//
	// Optional
	ChannelCount uint32
	// ChannelCountMode is "OscillatorOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "OscillatorOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_Frequency    bool // for Frequency.
	FFI_USE_Detune       bool // for Detune.
	FFI_USE_ChannelCount bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OscillatorOptions with all fields set.
func (p OscillatorOptions) FromRef(ref js.Ref) OscillatorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 OscillatorOptions OscillatorOptions [// OscillatorOptions] [0x140000fc1e0 0x140000fc280 0x140000fc3c0 0x140000fc780 0x140000fc820 0x140000fc960 0x140000fca00 0x140000fc320 0x140000fc460 0x140000fc8c0] 0x140010af440 {0 0}} in the application heap.
func (p OscillatorOptions) New() js.Ref {
	return bindings.OscillatorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p OscillatorOptions) UpdateFrom(ref js.Ref) {
	bindings.OscillatorOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p OscillatorOptions) Update(ref js.Ref) {
	bindings.OscillatorOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewOscillatorNode(context BaseAudioContext, options OscillatorOptions) OscillatorNode {
	return OscillatorNode{}.FromRef(
		bindings.NewOscillatorNodeByOscillatorNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewOscillatorNodeByOscillatorNode1(context BaseAudioContext) OscillatorNode {
	return OscillatorNode{}.FromRef(
		bindings.NewOscillatorNodeByOscillatorNode1(
			context.Ref()),
	)
}

type OscillatorNode struct {
	AudioScheduledSourceNode
}

func (this OscillatorNode) Once() OscillatorNode {
	this.Ref().Once()
	return this
}

func (this OscillatorNode) Ref() js.Ref {
	return this.AudioScheduledSourceNode.Ref()
}

func (this OscillatorNode) FromRef(ref js.Ref) OscillatorNode {
	this.AudioScheduledSourceNode = this.AudioScheduledSourceNode.FromRef(ref)
	return this
}

func (this OscillatorNode) Free() {
	this.Ref().Free()
}

// Type returns the value of property "OscillatorNode.type".
//
// The returned bool will be false if there is no such property.
func (this OscillatorNode) Type() (OscillatorType, bool) {
	var _ok bool
	_ret := bindings.GetOscillatorNodeType(
		this.Ref(), js.Pointer(&_ok),
	)
	return OscillatorType(_ret), _ok
}

// Type sets the value of property "OscillatorNode.type" to val.
//
// It returns false if the property cannot be set.
func (this OscillatorNode) SetType(val OscillatorType) bool {
	return js.True == bindings.SetOscillatorNodeType(
		this.Ref(),
		uint32(val),
	)
}

// Frequency returns the value of property "OscillatorNode.frequency".
//
// The returned bool will be false if there is no such property.
func (this OscillatorNode) Frequency() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetOscillatorNodeFrequency(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// Detune returns the value of property "OscillatorNode.detune".
//
// The returned bool will be false if there is no such property.
func (this OscillatorNode) Detune() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetOscillatorNodeDetune(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// SetPeriodicWave calls the method "OscillatorNode.setPeriodicWave".
//
// The returned bool will be false if there is no such method.
func (this OscillatorNode) SetPeriodicWave(periodicWave PeriodicWave) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOscillatorNodeSetPeriodicWave(
		this.Ref(), js.Pointer(&_ok),
		periodicWave.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPeriodicWaveFunc returns the method "OscillatorNode.setPeriodicWave".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OscillatorNode) SetPeriodicWaveFunc() (fn js.Func[func(periodicWave PeriodicWave)]) {
	return fn.FromRef(
		bindings.OscillatorNodeSetPeriodicWaveFunc(
			this.Ref(),
		),
	)
}

type PanningModelType uint32

const (
	_ PanningModelType = iota

	PanningModelType_EQUALPOWER
	PanningModelType_HRTF
)

func (PanningModelType) FromRef(str js.Ref) PanningModelType {
	return PanningModelType(bindings.ConstOfPanningModelType(str))
}

func (x PanningModelType) String() (string, bool) {
	switch x {
	case PanningModelType_EQUALPOWER:
		return "equalpower", true
	case PanningModelType_HRTF:
		return "HRTF", true
	default:
		return "", false
	}
}

type DistanceModelType uint32

const (
	_ DistanceModelType = iota

	DistanceModelType_LINEAR
	DistanceModelType_INVERSE
	DistanceModelType_EXPONENTIAL
)

func (DistanceModelType) FromRef(str js.Ref) DistanceModelType {
	return DistanceModelType(bindings.ConstOfDistanceModelType(str))
}

func (x DistanceModelType) String() (string, bool) {
	switch x {
	case DistanceModelType_LINEAR:
		return "linear", true
	case DistanceModelType_INVERSE:
		return "inverse", true
	case DistanceModelType_EXPONENTIAL:
		return "exponential", true
	default:
		return "", false
	}
}

type PannerOptions struct {
	// PanningModel is "PannerOptions.panningModel"
	//
	// Optional, defaults to "equalpower".
	PanningModel PanningModelType
	// DistanceModel is "PannerOptions.distanceModel"
	//
	// Optional, defaults to "inverse".
	DistanceModel DistanceModelType
	// PositionX is "PannerOptions.positionX"
	//
	// Optional, defaults to 0.
	PositionX float32
	// PositionY is "PannerOptions.positionY"
	//
	// Optional, defaults to 0.
	PositionY float32
	// PositionZ is "PannerOptions.positionZ"
	//
	// Optional, defaults to 0.
	PositionZ float32
	// OrientationX is "PannerOptions.orientationX"
	//
	// Optional, defaults to 1.
	OrientationX float32
	// OrientationY is "PannerOptions.orientationY"
	//
	// Optional, defaults to 0.
	OrientationY float32
	// OrientationZ is "PannerOptions.orientationZ"
	//
	// Optional, defaults to 0.
	OrientationZ float32
	// RefDistance is "PannerOptions.refDistance"
	//
	// Optional, defaults to 1.
	RefDistance float64
	// MaxDistance is "PannerOptions.maxDistance"
	//
	// Optional, defaults to 10000.
	MaxDistance float64
	// RolloffFactor is "PannerOptions.rolloffFactor"
	//
	// Optional, defaults to 1.
	RolloffFactor float64
	// ConeInnerAngle is "PannerOptions.coneInnerAngle"
	//
	// Optional, defaults to 360.
	ConeInnerAngle float64
	// ConeOuterAngle is "PannerOptions.coneOuterAngle"
	//
	// Optional, defaults to 360.
	ConeOuterAngle float64
	// ConeOuterGain is "PannerOptions.coneOuterGain"
	//
	// Optional, defaults to 0.
	ConeOuterGain float64
	// ChannelCount is "PannerOptions.channelCount"
	//
	// Optional
	ChannelCount uint32
	// ChannelCountMode is "PannerOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "PannerOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_PositionX      bool // for PositionX.
	FFI_USE_PositionY      bool // for PositionY.
	FFI_USE_PositionZ      bool // for PositionZ.
	FFI_USE_OrientationX   bool // for OrientationX.
	FFI_USE_OrientationY   bool // for OrientationY.
	FFI_USE_OrientationZ   bool // for OrientationZ.
	FFI_USE_RefDistance    bool // for RefDistance.
	FFI_USE_MaxDistance    bool // for MaxDistance.
	FFI_USE_RolloffFactor  bool // for RolloffFactor.
	FFI_USE_ConeInnerAngle bool // for ConeInnerAngle.
	FFI_USE_ConeOuterAngle bool // for ConeOuterAngle.
	FFI_USE_ConeOuterGain  bool // for ConeOuterGain.
	FFI_USE_ChannelCount   bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PannerOptions with all fields set.
func (p PannerOptions) FromRef(ref js.Ref) PannerOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PannerOptions PannerOptions [// PannerOptions] [0x140000fcb40 0x140000fcbe0 0x140000fcc80 0x140000fcdc0 0x140000fcf00 0x140000fd040 0x140000fd180 0x140000fd2c0 0x140000fd400 0x140000fd540 0x140000fd680 0x140000fd7c0 0x140000fd900 0x140000fda40 0x140000fdb80 0x140000fdcc0 0x140000fdd60 0x140000fcd20 0x140000fce60 0x140000fcfa0 0x140000fd0e0 0x140000fd220 0x140000fd360 0x140000fd4a0 0x140000fd5e0 0x140000fd720 0x140000fd860 0x140000fd9a0 0x140000fdae0 0x140000fdc20] 0x140010af530 {0 0}} in the application heap.
func (p PannerOptions) New() js.Ref {
	return bindings.PannerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PannerOptions) UpdateFrom(ref js.Ref) {
	bindings.PannerOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PannerOptions) Update(ref js.Ref) {
	bindings.PannerOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPannerNode(context BaseAudioContext, options PannerOptions) PannerNode {
	return PannerNode{}.FromRef(
		bindings.NewPannerNodeByPannerNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewPannerNodeByPannerNode1(context BaseAudioContext) PannerNode {
	return PannerNode{}.FromRef(
		bindings.NewPannerNodeByPannerNode1(
			context.Ref()),
	)
}

type PannerNode struct {
	AudioNode
}

func (this PannerNode) Once() PannerNode {
	this.Ref().Once()
	return this
}

func (this PannerNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this PannerNode) FromRef(ref js.Ref) PannerNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this PannerNode) Free() {
	this.Ref().Free()
}

// PanningModel returns the value of property "PannerNode.panningModel".
//
// The returned bool will be false if there is no such property.
func (this PannerNode) PanningModel() (PanningModelType, bool) {
	var _ok bool
	_ret := bindings.GetPannerNodePanningModel(
		this.Ref(), js.Pointer(&_ok),
	)
	return PanningModelType(_ret), _ok
}

// PanningModel sets the value of property "PannerNode.panningModel" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetPanningModel(val PanningModelType) bool {
	return js.True == bindings.SetPannerNodePanningModel(
		this.Ref(),
		uint32(val),
	)
}

// PositionX returns the value of property "PannerNode.positionX".
//
// The returned bool will be false if there is no such property.
func (this PannerNode) PositionX() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetPannerNodePositionX(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// PositionY returns the value of property "PannerNode.positionY".
//
// The returned bool will be false if there is no such property.
func (this PannerNode) PositionY() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetPannerNodePositionY(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// PositionZ returns the value of property "PannerNode.positionZ".
//
// The returned bool will be false if there is no such property.
func (this PannerNode) PositionZ() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetPannerNodePositionZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// OrientationX returns the value of property "PannerNode.orientationX".
//
// The returned bool will be false if there is no such property.
func (this PannerNode) OrientationX() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetPannerNodeOrientationX(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// OrientationY returns the value of property "PannerNode.orientationY".
//
// The returned bool will be false if there is no such property.
func (this PannerNode) OrientationY() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetPannerNodeOrientationY(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// OrientationZ returns the value of property "PannerNode.orientationZ".
//
// The returned bool will be false if there is no such property.
func (this PannerNode) OrientationZ() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetPannerNodeOrientationZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// DistanceModel returns the value of property "PannerNode.distanceModel".
//
// The returned bool will be false if there is no such property.
func (this PannerNode) DistanceModel() (DistanceModelType, bool) {
	var _ok bool
	_ret := bindings.GetPannerNodeDistanceModel(
		this.Ref(), js.Pointer(&_ok),
	)
	return DistanceModelType(_ret), _ok
}

// DistanceModel sets the value of property "PannerNode.distanceModel" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetDistanceModel(val DistanceModelType) bool {
	return js.True == bindings.SetPannerNodeDistanceModel(
		this.Ref(),
		uint32(val),
	)
}

// RefDistance returns the value of property "PannerNode.refDistance".
//
// The returned bool will be false if there is no such property.
func (this PannerNode) RefDistance() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPannerNodeRefDistance(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// RefDistance sets the value of property "PannerNode.refDistance" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetRefDistance(val float64) bool {
	return js.True == bindings.SetPannerNodeRefDistance(
		this.Ref(),
		float64(val),
	)
}

// MaxDistance returns the value of property "PannerNode.maxDistance".
//
// The returned bool will be false if there is no such property.
func (this PannerNode) MaxDistance() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPannerNodeMaxDistance(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// MaxDistance sets the value of property "PannerNode.maxDistance" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetMaxDistance(val float64) bool {
	return js.True == bindings.SetPannerNodeMaxDistance(
		this.Ref(),
		float64(val),
	)
}

// RolloffFactor returns the value of property "PannerNode.rolloffFactor".
//
// The returned bool will be false if there is no such property.
func (this PannerNode) RolloffFactor() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPannerNodeRolloffFactor(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// RolloffFactor sets the value of property "PannerNode.rolloffFactor" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetRolloffFactor(val float64) bool {
	return js.True == bindings.SetPannerNodeRolloffFactor(
		this.Ref(),
		float64(val),
	)
}

// ConeInnerAngle returns the value of property "PannerNode.coneInnerAngle".
//
// The returned bool will be false if there is no such property.
func (this PannerNode) ConeInnerAngle() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPannerNodeConeInnerAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ConeInnerAngle sets the value of property "PannerNode.coneInnerAngle" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetConeInnerAngle(val float64) bool {
	return js.True == bindings.SetPannerNodeConeInnerAngle(
		this.Ref(),
		float64(val),
	)
}

// ConeOuterAngle returns the value of property "PannerNode.coneOuterAngle".
//
// The returned bool will be false if there is no such property.
func (this PannerNode) ConeOuterAngle() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPannerNodeConeOuterAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ConeOuterAngle sets the value of property "PannerNode.coneOuterAngle" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetConeOuterAngle(val float64) bool {
	return js.True == bindings.SetPannerNodeConeOuterAngle(
		this.Ref(),
		float64(val),
	)
}

// ConeOuterGain returns the value of property "PannerNode.coneOuterGain".
//
// The returned bool will be false if there is no such property.
func (this PannerNode) ConeOuterGain() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPannerNodeConeOuterGain(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ConeOuterGain sets the value of property "PannerNode.coneOuterGain" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetConeOuterGain(val float64) bool {
	return js.True == bindings.SetPannerNodeConeOuterGain(
		this.Ref(),
		float64(val),
	)
}

// SetPosition calls the method "PannerNode.setPosition".
//
// The returned bool will be false if there is no such method.
func (this PannerNode) SetPosition(x float32, y float32, z float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPannerNodeSetPosition(
		this.Ref(), js.Pointer(&_ok),
		float32(x),
		float32(y),
		float32(z),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPositionFunc returns the method "PannerNode.setPosition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PannerNode) SetPositionFunc() (fn js.Func[func(x float32, y float32, z float32)]) {
	return fn.FromRef(
		bindings.PannerNodeSetPositionFunc(
			this.Ref(),
		),
	)
}

// SetOrientation calls the method "PannerNode.setOrientation".
//
// The returned bool will be false if there is no such method.
func (this PannerNode) SetOrientation(x float32, y float32, z float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPannerNodeSetOrientation(
		this.Ref(), js.Pointer(&_ok),
		float32(x),
		float32(y),
		float32(z),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetOrientationFunc returns the method "PannerNode.setOrientation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PannerNode) SetOrientationFunc() (fn js.Func[func(x float32, y float32, z float32)]) {
	return fn.FromRef(
		bindings.PannerNodeSetOrientationFunc(
			this.Ref(),
		),
	)
}

type PeriodicWaveConstraints struct {
	// DisableNormalization is "PeriodicWaveConstraints.disableNormalization"
	//
	// Optional, defaults to false.
	DisableNormalization bool

	FFI_USE_DisableNormalization bool // for DisableNormalization.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PeriodicWaveConstraints with all fields set.
func (p PeriodicWaveConstraints) FromRef(ref js.Ref) PeriodicWaveConstraints {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PeriodicWaveConstraints PeriodicWaveConstraints [// PeriodicWaveConstraints] [0x140000fdea0 0x140000fdf40] 0x14000aa3020 {0 0}} in the application heap.
func (p PeriodicWaveConstraints) New() js.Ref {
	return bindings.PeriodicWaveConstraintsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PeriodicWaveConstraints) UpdateFrom(ref js.Ref) {
	bindings.PeriodicWaveConstraintsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PeriodicWaveConstraints) Update(ref js.Ref) {
	bindings.PeriodicWaveConstraintsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ScriptProcessorNode struct {
	AudioNode
}

func (this ScriptProcessorNode) Once() ScriptProcessorNode {
	this.Ref().Once()
	return this
}

func (this ScriptProcessorNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this ScriptProcessorNode) FromRef(ref js.Ref) ScriptProcessorNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this ScriptProcessorNode) Free() {
	this.Ref().Free()
}

// BufferSize returns the value of property "ScriptProcessorNode.bufferSize".
//
// The returned bool will be false if there is no such property.
func (this ScriptProcessorNode) BufferSize() (int32, bool) {
	var _ok bool
	_ret := bindings.GetScriptProcessorNodeBufferSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

type StereoPannerOptions struct {
	// Pan is "StereoPannerOptions.pan"
	//
	// Optional, defaults to 0.
	Pan float32
	// ChannelCount is "StereoPannerOptions.channelCount"
	//
	// Optional
	ChannelCount uint32
	// ChannelCountMode is "StereoPannerOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "StereoPannerOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_Pan          bool // for Pan.
	FFI_USE_ChannelCount bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StereoPannerOptions with all fields set.
func (p StereoPannerOptions) FromRef(ref js.Ref) StereoPannerOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 StereoPannerOptions StereoPannerOptions [// StereoPannerOptions] [0x14000182000 0x14000182140 0x14000182280 0x14000182320 0x140001820a0 0x140001821e0] 0x14000aa3200 {0 0}} in the application heap.
func (p StereoPannerOptions) New() js.Ref {
	return bindings.StereoPannerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p StereoPannerOptions) UpdateFrom(ref js.Ref) {
	bindings.StereoPannerOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p StereoPannerOptions) Update(ref js.Ref) {
	bindings.StereoPannerOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewStereoPannerNode(context BaseAudioContext, options StereoPannerOptions) StereoPannerNode {
	return StereoPannerNode{}.FromRef(
		bindings.NewStereoPannerNodeByStereoPannerNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewStereoPannerNodeByStereoPannerNode1(context BaseAudioContext) StereoPannerNode {
	return StereoPannerNode{}.FromRef(
		bindings.NewStereoPannerNodeByStereoPannerNode1(
			context.Ref()),
	)
}

type StereoPannerNode struct {
	AudioNode
}

func (this StereoPannerNode) Once() StereoPannerNode {
	this.Ref().Once()
	return this
}

func (this StereoPannerNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this StereoPannerNode) FromRef(ref js.Ref) StereoPannerNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this StereoPannerNode) Free() {
	this.Ref().Free()
}

// Pan returns the value of property "StereoPannerNode.pan".
//
// The returned bool will be false if there is no such property.
func (this StereoPannerNode) Pan() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetStereoPannerNodePan(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

type OverSampleType uint32

const (
	_ OverSampleType = iota

	OverSampleType_NONE
	OverSampleType_2X
	OverSampleType_4X
)

func (OverSampleType) FromRef(str js.Ref) OverSampleType {
	return OverSampleType(bindings.ConstOfOverSampleType(str))
}

func (x OverSampleType) String() (string, bool) {
	switch x {
	case OverSampleType_NONE:
		return "none", true
	case OverSampleType_2X:
		return "2x", true
	case OverSampleType_4X:
		return "4x", true
	default:
		return "", false
	}
}
