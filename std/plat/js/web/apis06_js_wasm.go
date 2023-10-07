// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
	//
	// NOTE: FFI_USE_Q MUST be set to true to make this field effective.
	Q float32
	// Detune is "BiquadFilterOptions.detune"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Detune MUST be set to true to make this field effective.
	Detune float32
	// Frequency is "BiquadFilterOptions.frequency"
	//
	// Optional, defaults to 350.
	//
	// NOTE: FFI_USE_Frequency MUST be set to true to make this field effective.
	Frequency float32
	// Gain is "BiquadFilterOptions.gain"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Gain MUST be set to true to make this field effective.
	Gain float32
	// ChannelCount is "BiquadFilterOptions.channelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
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

// New creates a new BiquadFilterOptions in the application heap.
func (p BiquadFilterOptions) New() js.Ref {
	return bindings.BiquadFilterOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BiquadFilterOptions) UpdateFrom(ref js.Ref) {
	bindings.BiquadFilterOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BiquadFilterOptions) Update(ref js.Ref) {
	bindings.BiquadFilterOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BiquadFilterOptions) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// Value returns the value of property "AudioParam.value".
//
// It returns ok=false if there is no such property.
func (this AudioParam) Value() (ret float32, ok bool) {
	ok = js.True == bindings.GetAudioParamValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "AudioParam.value" to val.
//
// It returns false if the property cannot be set.
func (this AudioParam) SetValue(val float32) bool {
	return js.True == bindings.SetAudioParamValue(
		this.ref,
		float32(val),
	)
}

// AutomationRate returns the value of property "AudioParam.automationRate".
//
// It returns ok=false if there is no such property.
func (this AudioParam) AutomationRate() (ret AutomationRate, ok bool) {
	ok = js.True == bindings.GetAudioParamAutomationRate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAutomationRate sets the value of property "AudioParam.automationRate" to val.
//
// It returns false if the property cannot be set.
func (this AudioParam) SetAutomationRate(val AutomationRate) bool {
	return js.True == bindings.SetAudioParamAutomationRate(
		this.ref,
		uint32(val),
	)
}

// DefaultValue returns the value of property "AudioParam.defaultValue".
//
// It returns ok=false if there is no such property.
func (this AudioParam) DefaultValue() (ret float32, ok bool) {
	ok = js.True == bindings.GetAudioParamDefaultValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MinValue returns the value of property "AudioParam.minValue".
//
// It returns ok=false if there is no such property.
func (this AudioParam) MinValue() (ret float32, ok bool) {
	ok = js.True == bindings.GetAudioParamMinValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxValue returns the value of property "AudioParam.maxValue".
//
// It returns ok=false if there is no such property.
func (this AudioParam) MaxValue() (ret float32, ok bool) {
	ok = js.True == bindings.GetAudioParamMaxValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncSetValueAtTime returns true if the method "AudioParam.setValueAtTime" exists.
func (this AudioParam) HasFuncSetValueAtTime() bool {
	return js.True == bindings.HasFuncAudioParamSetValueAtTime(
		this.ref,
	)
}

// FuncSetValueAtTime returns the method "AudioParam.setValueAtTime".
func (this AudioParam) FuncSetValueAtTime() (fn js.Func[func(value float32, startTime float64) AudioParam]) {
	bindings.FuncAudioParamSetValueAtTime(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetValueAtTime calls the method "AudioParam.setValueAtTime".
func (this AudioParam) SetValueAtTime(value float32, startTime float64) (ret AudioParam) {
	bindings.CallAudioParamSetValueAtTime(
		this.ref, js.Pointer(&ret),
		float32(value),
		float64(startTime),
	)

	return
}

// TrySetValueAtTime calls the method "AudioParam.setValueAtTime"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioParam) TrySetValueAtTime(value float32, startTime float64) (ret AudioParam, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioParamSetValueAtTime(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(value),
		float64(startTime),
	)

	return
}

// HasFuncLinearRampToValueAtTime returns true if the method "AudioParam.linearRampToValueAtTime" exists.
func (this AudioParam) HasFuncLinearRampToValueAtTime() bool {
	return js.True == bindings.HasFuncAudioParamLinearRampToValueAtTime(
		this.ref,
	)
}

// FuncLinearRampToValueAtTime returns the method "AudioParam.linearRampToValueAtTime".
func (this AudioParam) FuncLinearRampToValueAtTime() (fn js.Func[func(value float32, endTime float64) AudioParam]) {
	bindings.FuncAudioParamLinearRampToValueAtTime(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LinearRampToValueAtTime calls the method "AudioParam.linearRampToValueAtTime".
func (this AudioParam) LinearRampToValueAtTime(value float32, endTime float64) (ret AudioParam) {
	bindings.CallAudioParamLinearRampToValueAtTime(
		this.ref, js.Pointer(&ret),
		float32(value),
		float64(endTime),
	)

	return
}

// TryLinearRampToValueAtTime calls the method "AudioParam.linearRampToValueAtTime"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioParam) TryLinearRampToValueAtTime(value float32, endTime float64) (ret AudioParam, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioParamLinearRampToValueAtTime(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(value),
		float64(endTime),
	)

	return
}

// HasFuncExponentialRampToValueAtTime returns true if the method "AudioParam.exponentialRampToValueAtTime" exists.
func (this AudioParam) HasFuncExponentialRampToValueAtTime() bool {
	return js.True == bindings.HasFuncAudioParamExponentialRampToValueAtTime(
		this.ref,
	)
}

// FuncExponentialRampToValueAtTime returns the method "AudioParam.exponentialRampToValueAtTime".
func (this AudioParam) FuncExponentialRampToValueAtTime() (fn js.Func[func(value float32, endTime float64) AudioParam]) {
	bindings.FuncAudioParamExponentialRampToValueAtTime(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ExponentialRampToValueAtTime calls the method "AudioParam.exponentialRampToValueAtTime".
func (this AudioParam) ExponentialRampToValueAtTime(value float32, endTime float64) (ret AudioParam) {
	bindings.CallAudioParamExponentialRampToValueAtTime(
		this.ref, js.Pointer(&ret),
		float32(value),
		float64(endTime),
	)

	return
}

// TryExponentialRampToValueAtTime calls the method "AudioParam.exponentialRampToValueAtTime"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioParam) TryExponentialRampToValueAtTime(value float32, endTime float64) (ret AudioParam, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioParamExponentialRampToValueAtTime(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(value),
		float64(endTime),
	)

	return
}

// HasFuncSetTargetAtTime returns true if the method "AudioParam.setTargetAtTime" exists.
func (this AudioParam) HasFuncSetTargetAtTime() bool {
	return js.True == bindings.HasFuncAudioParamSetTargetAtTime(
		this.ref,
	)
}

// FuncSetTargetAtTime returns the method "AudioParam.setTargetAtTime".
func (this AudioParam) FuncSetTargetAtTime() (fn js.Func[func(target float32, startTime float64, timeConstant float32) AudioParam]) {
	bindings.FuncAudioParamSetTargetAtTime(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetTargetAtTime calls the method "AudioParam.setTargetAtTime".
func (this AudioParam) SetTargetAtTime(target float32, startTime float64, timeConstant float32) (ret AudioParam) {
	bindings.CallAudioParamSetTargetAtTime(
		this.ref, js.Pointer(&ret),
		float32(target),
		float64(startTime),
		float32(timeConstant),
	)

	return
}

// TrySetTargetAtTime calls the method "AudioParam.setTargetAtTime"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioParam) TrySetTargetAtTime(target float32, startTime float64, timeConstant float32) (ret AudioParam, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioParamSetTargetAtTime(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(target),
		float64(startTime),
		float32(timeConstant),
	)

	return
}

// HasFuncSetValueCurveAtTime returns true if the method "AudioParam.setValueCurveAtTime" exists.
func (this AudioParam) HasFuncSetValueCurveAtTime() bool {
	return js.True == bindings.HasFuncAudioParamSetValueCurveAtTime(
		this.ref,
	)
}

// FuncSetValueCurveAtTime returns the method "AudioParam.setValueCurveAtTime".
func (this AudioParam) FuncSetValueCurveAtTime() (fn js.Func[func(values js.Array[float32], startTime float64, duration float64) AudioParam]) {
	bindings.FuncAudioParamSetValueCurveAtTime(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetValueCurveAtTime calls the method "AudioParam.setValueCurveAtTime".
func (this AudioParam) SetValueCurveAtTime(values js.Array[float32], startTime float64, duration float64) (ret AudioParam) {
	bindings.CallAudioParamSetValueCurveAtTime(
		this.ref, js.Pointer(&ret),
		values.Ref(),
		float64(startTime),
		float64(duration),
	)

	return
}

// TrySetValueCurveAtTime calls the method "AudioParam.setValueCurveAtTime"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioParam) TrySetValueCurveAtTime(values js.Array[float32], startTime float64, duration float64) (ret AudioParam, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioParamSetValueCurveAtTime(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		values.Ref(),
		float64(startTime),
		float64(duration),
	)

	return
}

// HasFuncCancelScheduledValues returns true if the method "AudioParam.cancelScheduledValues" exists.
func (this AudioParam) HasFuncCancelScheduledValues() bool {
	return js.True == bindings.HasFuncAudioParamCancelScheduledValues(
		this.ref,
	)
}

// FuncCancelScheduledValues returns the method "AudioParam.cancelScheduledValues".
func (this AudioParam) FuncCancelScheduledValues() (fn js.Func[func(cancelTime float64) AudioParam]) {
	bindings.FuncAudioParamCancelScheduledValues(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CancelScheduledValues calls the method "AudioParam.cancelScheduledValues".
func (this AudioParam) CancelScheduledValues(cancelTime float64) (ret AudioParam) {
	bindings.CallAudioParamCancelScheduledValues(
		this.ref, js.Pointer(&ret),
		float64(cancelTime),
	)

	return
}

// TryCancelScheduledValues calls the method "AudioParam.cancelScheduledValues"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioParam) TryCancelScheduledValues(cancelTime float64) (ret AudioParam, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioParamCancelScheduledValues(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(cancelTime),
	)

	return
}

// HasFuncCancelAndHoldAtTime returns true if the method "AudioParam.cancelAndHoldAtTime" exists.
func (this AudioParam) HasFuncCancelAndHoldAtTime() bool {
	return js.True == bindings.HasFuncAudioParamCancelAndHoldAtTime(
		this.ref,
	)
}

// FuncCancelAndHoldAtTime returns the method "AudioParam.cancelAndHoldAtTime".
func (this AudioParam) FuncCancelAndHoldAtTime() (fn js.Func[func(cancelTime float64) AudioParam]) {
	bindings.FuncAudioParamCancelAndHoldAtTime(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CancelAndHoldAtTime calls the method "AudioParam.cancelAndHoldAtTime".
func (this AudioParam) CancelAndHoldAtTime(cancelTime float64) (ret AudioParam) {
	bindings.CallAudioParamCancelAndHoldAtTime(
		this.ref, js.Pointer(&ret),
		float64(cancelTime),
	)

	return
}

// TryCancelAndHoldAtTime calls the method "AudioParam.cancelAndHoldAtTime"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioParam) TryCancelAndHoldAtTime(cancelTime float64) (ret AudioParam, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioParamCancelAndHoldAtTime(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(cancelTime),
	)

	return
}

func NewBiquadFilterNode(context BaseAudioContext, options BiquadFilterOptions) (ret BiquadFilterNode) {
	ret.ref = bindings.NewBiquadFilterNodeByBiquadFilterNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewBiquadFilterNodeByBiquadFilterNode1(context BaseAudioContext) (ret BiquadFilterNode) {
	ret.ref = bindings.NewBiquadFilterNodeByBiquadFilterNode1(
		context.Ref())
	return
}

type BiquadFilterNode struct {
	AudioNode
}

func (this BiquadFilterNode) Once() BiquadFilterNode {
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "BiquadFilterNode.type".
//
// It returns ok=false if there is no such property.
func (this BiquadFilterNode) Type() (ret BiquadFilterType, ok bool) {
	ok = js.True == bindings.GetBiquadFilterNodeType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "BiquadFilterNode.type" to val.
//
// It returns false if the property cannot be set.
func (this BiquadFilterNode) SetType(val BiquadFilterType) bool {
	return js.True == bindings.SetBiquadFilterNodeType(
		this.ref,
		uint32(val),
	)
}

// Frequency returns the value of property "BiquadFilterNode.frequency".
//
// It returns ok=false if there is no such property.
func (this BiquadFilterNode) Frequency() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetBiquadFilterNodeFrequency(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Detune returns the value of property "BiquadFilterNode.detune".
//
// It returns ok=false if there is no such property.
func (this BiquadFilterNode) Detune() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetBiquadFilterNodeDetune(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Q returns the value of property "BiquadFilterNode.Q".
//
// It returns ok=false if there is no such property.
func (this BiquadFilterNode) Q() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetBiquadFilterNodeQ(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Gain returns the value of property "BiquadFilterNode.gain".
//
// It returns ok=false if there is no such property.
func (this BiquadFilterNode) Gain() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetBiquadFilterNodeGain(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetFrequencyResponse returns true if the method "BiquadFilterNode.getFrequencyResponse" exists.
func (this BiquadFilterNode) HasFuncGetFrequencyResponse() bool {
	return js.True == bindings.HasFuncBiquadFilterNodeGetFrequencyResponse(
		this.ref,
	)
}

// FuncGetFrequencyResponse returns the method "BiquadFilterNode.getFrequencyResponse".
func (this BiquadFilterNode) FuncGetFrequencyResponse() (fn js.Func[func(frequencyHz js.TypedArray[float32], magResponse js.TypedArray[float32], phaseResponse js.TypedArray[float32])]) {
	bindings.FuncBiquadFilterNodeGetFrequencyResponse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFrequencyResponse calls the method "BiquadFilterNode.getFrequencyResponse".
func (this BiquadFilterNode) GetFrequencyResponse(frequencyHz js.TypedArray[float32], magResponse js.TypedArray[float32], phaseResponse js.TypedArray[float32]) (ret js.Void) {
	bindings.CallBiquadFilterNodeGetFrequencyResponse(
		this.ref, js.Pointer(&ret),
		frequencyHz.Ref(),
		magResponse.Ref(),
		phaseResponse.Ref(),
	)

	return
}

// TryGetFrequencyResponse calls the method "BiquadFilterNode.getFrequencyResponse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BiquadFilterNode) TryGetFrequencyResponse(frequencyHz js.TypedArray[float32], magResponse js.TypedArray[float32], phaseResponse js.TypedArray[float32]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBiquadFilterNodeGetFrequencyResponse(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		frequencyHz.Ref(),
		magResponse.Ref(),
		phaseResponse.Ref(),
	)

	return
}

type AudioBufferOptions struct {
	// NumberOfChannels is "AudioBufferOptions.numberOfChannels"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_NumberOfChannels MUST be set to true to make this field effective.
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

// New creates a new AudioBufferOptions in the application heap.
func (p AudioBufferOptions) New() js.Ref {
	return bindings.AudioBufferOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioBufferOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioBufferOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioBufferOptions) Update(ref js.Ref) {
	bindings.AudioBufferOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioBufferOptions) FreeMembers(recursive bool) {
}

func NewAudioBuffer(options AudioBufferOptions) (ret AudioBuffer) {
	ret.ref = bindings.NewAudioBufferByAudioBuffer(
		js.Pointer(&options))
	return
}

type AudioBuffer struct {
	ref js.Ref
}

func (this AudioBuffer) Once() AudioBuffer {
	this.ref.Once()
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
	this.ref.Free()
}

// SampleRate returns the value of property "AudioBuffer.sampleRate".
//
// It returns ok=false if there is no such property.
func (this AudioBuffer) SampleRate() (ret float32, ok bool) {
	ok = js.True == bindings.GetAudioBufferSampleRate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Length returns the value of property "AudioBuffer.length".
//
// It returns ok=false if there is no such property.
func (this AudioBuffer) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioBufferLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Duration returns the value of property "AudioBuffer.duration".
//
// It returns ok=false if there is no such property.
func (this AudioBuffer) Duration() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioBufferDuration(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NumberOfChannels returns the value of property "AudioBuffer.numberOfChannels".
//
// It returns ok=false if there is no such property.
func (this AudioBuffer) NumberOfChannels() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioBufferNumberOfChannels(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetChannelData returns true if the method "AudioBuffer.getChannelData" exists.
func (this AudioBuffer) HasFuncGetChannelData() bool {
	return js.True == bindings.HasFuncAudioBufferGetChannelData(
		this.ref,
	)
}

// FuncGetChannelData returns the method "AudioBuffer.getChannelData".
func (this AudioBuffer) FuncGetChannelData() (fn js.Func[func(channel uint32) js.TypedArray[float32]]) {
	bindings.FuncAudioBufferGetChannelData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetChannelData calls the method "AudioBuffer.getChannelData".
func (this AudioBuffer) GetChannelData(channel uint32) (ret js.TypedArray[float32]) {
	bindings.CallAudioBufferGetChannelData(
		this.ref, js.Pointer(&ret),
		uint32(channel),
	)

	return
}

// TryGetChannelData calls the method "AudioBuffer.getChannelData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioBuffer) TryGetChannelData(channel uint32) (ret js.TypedArray[float32], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioBufferGetChannelData(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(channel),
	)

	return
}

// HasFuncCopyFromChannel returns true if the method "AudioBuffer.copyFromChannel" exists.
func (this AudioBuffer) HasFuncCopyFromChannel() bool {
	return js.True == bindings.HasFuncAudioBufferCopyFromChannel(
		this.ref,
	)
}

// FuncCopyFromChannel returns the method "AudioBuffer.copyFromChannel".
func (this AudioBuffer) FuncCopyFromChannel() (fn js.Func[func(destination js.TypedArray[float32], channelNumber uint32, bufferOffset uint32)]) {
	bindings.FuncAudioBufferCopyFromChannel(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyFromChannel calls the method "AudioBuffer.copyFromChannel".
func (this AudioBuffer) CopyFromChannel(destination js.TypedArray[float32], channelNumber uint32, bufferOffset uint32) (ret js.Void) {
	bindings.CallAudioBufferCopyFromChannel(
		this.ref, js.Pointer(&ret),
		destination.Ref(),
		uint32(channelNumber),
		uint32(bufferOffset),
	)

	return
}

// TryCopyFromChannel calls the method "AudioBuffer.copyFromChannel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioBuffer) TryCopyFromChannel(destination js.TypedArray[float32], channelNumber uint32, bufferOffset uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioBufferCopyFromChannel(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destination.Ref(),
		uint32(channelNumber),
		uint32(bufferOffset),
	)

	return
}

// HasFuncCopyFromChannel1 returns true if the method "AudioBuffer.copyFromChannel" exists.
func (this AudioBuffer) HasFuncCopyFromChannel1() bool {
	return js.True == bindings.HasFuncAudioBufferCopyFromChannel1(
		this.ref,
	)
}

// FuncCopyFromChannel1 returns the method "AudioBuffer.copyFromChannel".
func (this AudioBuffer) FuncCopyFromChannel1() (fn js.Func[func(destination js.TypedArray[float32], channelNumber uint32)]) {
	bindings.FuncAudioBufferCopyFromChannel1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyFromChannel1 calls the method "AudioBuffer.copyFromChannel".
func (this AudioBuffer) CopyFromChannel1(destination js.TypedArray[float32], channelNumber uint32) (ret js.Void) {
	bindings.CallAudioBufferCopyFromChannel1(
		this.ref, js.Pointer(&ret),
		destination.Ref(),
		uint32(channelNumber),
	)

	return
}

// TryCopyFromChannel1 calls the method "AudioBuffer.copyFromChannel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioBuffer) TryCopyFromChannel1(destination js.TypedArray[float32], channelNumber uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioBufferCopyFromChannel1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destination.Ref(),
		uint32(channelNumber),
	)

	return
}

// HasFuncCopyToChannel returns true if the method "AudioBuffer.copyToChannel" exists.
func (this AudioBuffer) HasFuncCopyToChannel() bool {
	return js.True == bindings.HasFuncAudioBufferCopyToChannel(
		this.ref,
	)
}

// FuncCopyToChannel returns the method "AudioBuffer.copyToChannel".
func (this AudioBuffer) FuncCopyToChannel() (fn js.Func[func(source js.TypedArray[float32], channelNumber uint32, bufferOffset uint32)]) {
	bindings.FuncAudioBufferCopyToChannel(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyToChannel calls the method "AudioBuffer.copyToChannel".
func (this AudioBuffer) CopyToChannel(source js.TypedArray[float32], channelNumber uint32, bufferOffset uint32) (ret js.Void) {
	bindings.CallAudioBufferCopyToChannel(
		this.ref, js.Pointer(&ret),
		source.Ref(),
		uint32(channelNumber),
		uint32(bufferOffset),
	)

	return
}

// TryCopyToChannel calls the method "AudioBuffer.copyToChannel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioBuffer) TryCopyToChannel(source js.TypedArray[float32], channelNumber uint32, bufferOffset uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioBufferCopyToChannel(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		source.Ref(),
		uint32(channelNumber),
		uint32(bufferOffset),
	)

	return
}

// HasFuncCopyToChannel1 returns true if the method "AudioBuffer.copyToChannel" exists.
func (this AudioBuffer) HasFuncCopyToChannel1() bool {
	return js.True == bindings.HasFuncAudioBufferCopyToChannel1(
		this.ref,
	)
}

// FuncCopyToChannel1 returns the method "AudioBuffer.copyToChannel".
func (this AudioBuffer) FuncCopyToChannel1() (fn js.Func[func(source js.TypedArray[float32], channelNumber uint32)]) {
	bindings.FuncAudioBufferCopyToChannel1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyToChannel1 calls the method "AudioBuffer.copyToChannel".
func (this AudioBuffer) CopyToChannel1(source js.TypedArray[float32], channelNumber uint32) (ret js.Void) {
	bindings.CallAudioBufferCopyToChannel1(
		this.ref, js.Pointer(&ret),
		source.Ref(),
		uint32(channelNumber),
	)

	return
}

// TryCopyToChannel1 calls the method "AudioBuffer.copyToChannel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioBuffer) TryCopyToChannel1(source js.TypedArray[float32], channelNumber uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioBufferCopyToChannel1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		source.Ref(),
		uint32(channelNumber),
	)

	return
}

type AudioBufferSourceOptions struct {
	// Buffer is "AudioBufferSourceOptions.buffer"
	//
	// Optional
	Buffer AudioBuffer
	// Detune is "AudioBufferSourceOptions.detune"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Detune MUST be set to true to make this field effective.
	Detune float32
	// Loop is "AudioBufferSourceOptions.loop"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Loop MUST be set to true to make this field effective.
	Loop bool
	// LoopEnd is "AudioBufferSourceOptions.loopEnd"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_LoopEnd MUST be set to true to make this field effective.
	LoopEnd float64
	// LoopStart is "AudioBufferSourceOptions.loopStart"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_LoopStart MUST be set to true to make this field effective.
	LoopStart float64
	// PlaybackRate is "AudioBufferSourceOptions.playbackRate"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_PlaybackRate MUST be set to true to make this field effective.
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

// New creates a new AudioBufferSourceOptions in the application heap.
func (p AudioBufferSourceOptions) New() js.Ref {
	return bindings.AudioBufferSourceOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioBufferSourceOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioBufferSourceOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioBufferSourceOptions) Update(ref js.Ref) {
	bindings.AudioBufferSourceOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioBufferSourceOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Buffer.Ref(),
	)
	p.Buffer = p.Buffer.FromRef(js.Undefined)
}

func NewAudioBufferSourceNode(context BaseAudioContext, options AudioBufferSourceOptions) (ret AudioBufferSourceNode) {
	ret.ref = bindings.NewAudioBufferSourceNodeByAudioBufferSourceNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewAudioBufferSourceNodeByAudioBufferSourceNode1(context BaseAudioContext) (ret AudioBufferSourceNode) {
	ret.ref = bindings.NewAudioBufferSourceNodeByAudioBufferSourceNode1(
		context.Ref())
	return
}

type AudioBufferSourceNode struct {
	AudioScheduledSourceNode
}

func (this AudioBufferSourceNode) Once() AudioBufferSourceNode {
	this.ref.Once()
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
	this.ref.Free()
}

// Buffer returns the value of property "AudioBufferSourceNode.buffer".
//
// It returns ok=false if there is no such property.
func (this AudioBufferSourceNode) Buffer() (ret AudioBuffer, ok bool) {
	ok = js.True == bindings.GetAudioBufferSourceNodeBuffer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBuffer sets the value of property "AudioBufferSourceNode.buffer" to val.
//
// It returns false if the property cannot be set.
func (this AudioBufferSourceNode) SetBuffer(val AudioBuffer) bool {
	return js.True == bindings.SetAudioBufferSourceNodeBuffer(
		this.ref,
		val.Ref(),
	)
}

// PlaybackRate returns the value of property "AudioBufferSourceNode.playbackRate".
//
// It returns ok=false if there is no such property.
func (this AudioBufferSourceNode) PlaybackRate() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioBufferSourceNodePlaybackRate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Detune returns the value of property "AudioBufferSourceNode.detune".
//
// It returns ok=false if there is no such property.
func (this AudioBufferSourceNode) Detune() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioBufferSourceNodeDetune(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Loop returns the value of property "AudioBufferSourceNode.loop".
//
// It returns ok=false if there is no such property.
func (this AudioBufferSourceNode) Loop() (ret bool, ok bool) {
	ok = js.True == bindings.GetAudioBufferSourceNodeLoop(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLoop sets the value of property "AudioBufferSourceNode.loop" to val.
//
// It returns false if the property cannot be set.
func (this AudioBufferSourceNode) SetLoop(val bool) bool {
	return js.True == bindings.SetAudioBufferSourceNodeLoop(
		this.ref,
		js.Bool(bool(val)),
	)
}

// LoopStart returns the value of property "AudioBufferSourceNode.loopStart".
//
// It returns ok=false if there is no such property.
func (this AudioBufferSourceNode) LoopStart() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioBufferSourceNodeLoopStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLoopStart sets the value of property "AudioBufferSourceNode.loopStart" to val.
//
// It returns false if the property cannot be set.
func (this AudioBufferSourceNode) SetLoopStart(val float64) bool {
	return js.True == bindings.SetAudioBufferSourceNodeLoopStart(
		this.ref,
		float64(val),
	)
}

// LoopEnd returns the value of property "AudioBufferSourceNode.loopEnd".
//
// It returns ok=false if there is no such property.
func (this AudioBufferSourceNode) LoopEnd() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioBufferSourceNodeLoopEnd(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLoopEnd sets the value of property "AudioBufferSourceNode.loopEnd" to val.
//
// It returns false if the property cannot be set.
func (this AudioBufferSourceNode) SetLoopEnd(val float64) bool {
	return js.True == bindings.SetAudioBufferSourceNodeLoopEnd(
		this.ref,
		float64(val),
	)
}

// HasFuncStart returns true if the method "AudioBufferSourceNode.start" exists.
func (this AudioBufferSourceNode) HasFuncStart() bool {
	return js.True == bindings.HasFuncAudioBufferSourceNodeStart(
		this.ref,
	)
}

// FuncStart returns the method "AudioBufferSourceNode.start".
func (this AudioBufferSourceNode) FuncStart() (fn js.Func[func(when float64, offset float64, duration float64)]) {
	bindings.FuncAudioBufferSourceNodeStart(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start calls the method "AudioBufferSourceNode.start".
func (this AudioBufferSourceNode) Start(when float64, offset float64, duration float64) (ret js.Void) {
	bindings.CallAudioBufferSourceNodeStart(
		this.ref, js.Pointer(&ret),
		float64(when),
		float64(offset),
		float64(duration),
	)

	return
}

// TryStart calls the method "AudioBufferSourceNode.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioBufferSourceNode) TryStart(when float64, offset float64, duration float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioBufferSourceNodeStart(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(when),
		float64(offset),
		float64(duration),
	)

	return
}

// HasFuncStart1 returns true if the method "AudioBufferSourceNode.start" exists.
func (this AudioBufferSourceNode) HasFuncStart1() bool {
	return js.True == bindings.HasFuncAudioBufferSourceNodeStart1(
		this.ref,
	)
}

// FuncStart1 returns the method "AudioBufferSourceNode.start".
func (this AudioBufferSourceNode) FuncStart1() (fn js.Func[func(when float64, offset float64)]) {
	bindings.FuncAudioBufferSourceNodeStart1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start1 calls the method "AudioBufferSourceNode.start".
func (this AudioBufferSourceNode) Start1(when float64, offset float64) (ret js.Void) {
	bindings.CallAudioBufferSourceNodeStart1(
		this.ref, js.Pointer(&ret),
		float64(when),
		float64(offset),
	)

	return
}

// TryStart1 calls the method "AudioBufferSourceNode.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioBufferSourceNode) TryStart1(when float64, offset float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioBufferSourceNodeStart1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(when),
		float64(offset),
	)

	return
}

// HasFuncStart2 returns true if the method "AudioBufferSourceNode.start" exists.
func (this AudioBufferSourceNode) HasFuncStart2() bool {
	return js.True == bindings.HasFuncAudioBufferSourceNodeStart2(
		this.ref,
	)
}

// FuncStart2 returns the method "AudioBufferSourceNode.start".
func (this AudioBufferSourceNode) FuncStart2() (fn js.Func[func(when float64)]) {
	bindings.FuncAudioBufferSourceNodeStart2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start2 calls the method "AudioBufferSourceNode.start".
func (this AudioBufferSourceNode) Start2(when float64) (ret js.Void) {
	bindings.CallAudioBufferSourceNodeStart2(
		this.ref, js.Pointer(&ret),
		float64(when),
	)

	return
}

// TryStart2 calls the method "AudioBufferSourceNode.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioBufferSourceNode) TryStart2(when float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioBufferSourceNodeStart2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(when),
	)

	return
}

// HasFuncStart3 returns true if the method "AudioBufferSourceNode.start" exists.
func (this AudioBufferSourceNode) HasFuncStart3() bool {
	return js.True == bindings.HasFuncAudioBufferSourceNodeStart3(
		this.ref,
	)
}

// FuncStart3 returns the method "AudioBufferSourceNode.start".
func (this AudioBufferSourceNode) FuncStart3() (fn js.Func[func()]) {
	bindings.FuncAudioBufferSourceNodeStart3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start3 calls the method "AudioBufferSourceNode.start".
func (this AudioBufferSourceNode) Start3() (ret js.Void) {
	bindings.CallAudioBufferSourceNodeStart3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStart3 calls the method "AudioBufferSourceNode.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioBufferSourceNode) TryStart3() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioBufferSourceNodeStart3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ChannelMergerOptions struct {
	// NumberOfInputs is "ChannelMergerOptions.numberOfInputs"
	//
	// Optional, defaults to 6.
	//
	// NOTE: FFI_USE_NumberOfInputs MUST be set to true to make this field effective.
	NumberOfInputs uint32
	// ChannelCount is "ChannelMergerOptions.channelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
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

// New creates a new ChannelMergerOptions in the application heap.
func (p ChannelMergerOptions) New() js.Ref {
	return bindings.ChannelMergerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ChannelMergerOptions) UpdateFrom(ref js.Ref) {
	bindings.ChannelMergerOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ChannelMergerOptions) Update(ref js.Ref) {
	bindings.ChannelMergerOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ChannelMergerOptions) FreeMembers(recursive bool) {
}

func NewChannelMergerNode(context BaseAudioContext, options ChannelMergerOptions) (ret ChannelMergerNode) {
	ret.ref = bindings.NewChannelMergerNodeByChannelMergerNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewChannelMergerNodeByChannelMergerNode1(context BaseAudioContext) (ret ChannelMergerNode) {
	ret.ref = bindings.NewChannelMergerNodeByChannelMergerNode1(
		context.Ref())
	return
}

type ChannelMergerNode struct {
	AudioNode
}

func (this ChannelMergerNode) Once() ChannelMergerNode {
	this.ref.Once()
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
	this.ref.Free()
}

type ChannelSplitterOptions struct {
	// NumberOfOutputs is "ChannelSplitterOptions.numberOfOutputs"
	//
	// Optional, defaults to 6.
	//
	// NOTE: FFI_USE_NumberOfOutputs MUST be set to true to make this field effective.
	NumberOfOutputs uint32
	// ChannelCount is "ChannelSplitterOptions.channelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
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

// New creates a new ChannelSplitterOptions in the application heap.
func (p ChannelSplitterOptions) New() js.Ref {
	return bindings.ChannelSplitterOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ChannelSplitterOptions) UpdateFrom(ref js.Ref) {
	bindings.ChannelSplitterOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ChannelSplitterOptions) Update(ref js.Ref) {
	bindings.ChannelSplitterOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ChannelSplitterOptions) FreeMembers(recursive bool) {
}

func NewChannelSplitterNode(context BaseAudioContext, options ChannelSplitterOptions) (ret ChannelSplitterNode) {
	ret.ref = bindings.NewChannelSplitterNodeByChannelSplitterNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewChannelSplitterNodeByChannelSplitterNode1(context BaseAudioContext) (ret ChannelSplitterNode) {
	ret.ref = bindings.NewChannelSplitterNodeByChannelSplitterNode1(
		context.Ref())
	return
}

type ChannelSplitterNode struct {
	AudioNode
}

func (this ChannelSplitterNode) Once() ChannelSplitterNode {
	this.ref.Once()
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
	this.ref.Free()
}

type ConstantSourceOptions struct {
	// Offset is "ConstantSourceOptions.offset"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_Offset MUST be set to true to make this field effective.
	Offset float32

	FFI_USE_Offset bool // for Offset.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConstantSourceOptions with all fields set.
func (p ConstantSourceOptions) FromRef(ref js.Ref) ConstantSourceOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConstantSourceOptions in the application heap.
func (p ConstantSourceOptions) New() js.Ref {
	return bindings.ConstantSourceOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConstantSourceOptions) UpdateFrom(ref js.Ref) {
	bindings.ConstantSourceOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConstantSourceOptions) Update(ref js.Ref) {
	bindings.ConstantSourceOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConstantSourceOptions) FreeMembers(recursive bool) {
}

func NewConstantSourceNode(context BaseAudioContext, options ConstantSourceOptions) (ret ConstantSourceNode) {
	ret.ref = bindings.NewConstantSourceNodeByConstantSourceNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewConstantSourceNodeByConstantSourceNode1(context BaseAudioContext) (ret ConstantSourceNode) {
	ret.ref = bindings.NewConstantSourceNodeByConstantSourceNode1(
		context.Ref())
	return
}

type ConstantSourceNode struct {
	AudioScheduledSourceNode
}

func (this ConstantSourceNode) Once() ConstantSourceNode {
	this.ref.Once()
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
	this.ref.Free()
}

// Offset returns the value of property "ConstantSourceNode.offset".
//
// It returns ok=false if there is no such property.
func (this ConstantSourceNode) Offset() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetConstantSourceNodeOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

type ConvolverOptions struct {
	// Buffer is "ConvolverOptions.buffer"
	//
	// Optional
	Buffer AudioBuffer
	// DisableNormalization is "ConvolverOptions.disableNormalization"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_DisableNormalization MUST be set to true to make this field effective.
	DisableNormalization bool
	// ChannelCount is "ConvolverOptions.channelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
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

// New creates a new ConvolverOptions in the application heap.
func (p ConvolverOptions) New() js.Ref {
	return bindings.ConvolverOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConvolverOptions) UpdateFrom(ref js.Ref) {
	bindings.ConvolverOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConvolverOptions) Update(ref js.Ref) {
	bindings.ConvolverOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConvolverOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Buffer.Ref(),
	)
	p.Buffer = p.Buffer.FromRef(js.Undefined)
}

func NewConvolverNode(context BaseAudioContext, options ConvolverOptions) (ret ConvolverNode) {
	ret.ref = bindings.NewConvolverNodeByConvolverNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewConvolverNodeByConvolverNode1(context BaseAudioContext) (ret ConvolverNode) {
	ret.ref = bindings.NewConvolverNodeByConvolverNode1(
		context.Ref())
	return
}

type ConvolverNode struct {
	AudioNode
}

func (this ConvolverNode) Once() ConvolverNode {
	this.ref.Once()
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
	this.ref.Free()
}

// Buffer returns the value of property "ConvolverNode.buffer".
//
// It returns ok=false if there is no such property.
func (this ConvolverNode) Buffer() (ret AudioBuffer, ok bool) {
	ok = js.True == bindings.GetConvolverNodeBuffer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBuffer sets the value of property "ConvolverNode.buffer" to val.
//
// It returns false if the property cannot be set.
func (this ConvolverNode) SetBuffer(val AudioBuffer) bool {
	return js.True == bindings.SetConvolverNodeBuffer(
		this.ref,
		val.Ref(),
	)
}

// Normalize returns the value of property "ConvolverNode.normalize".
//
// It returns ok=false if there is no such property.
func (this ConvolverNode) Normalize() (ret bool, ok bool) {
	ok = js.True == bindings.GetConvolverNodeNormalize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNormalize sets the value of property "ConvolverNode.normalize" to val.
//
// It returns false if the property cannot be set.
func (this ConvolverNode) SetNormalize(val bool) bool {
	return js.True == bindings.SetConvolverNodeNormalize(
		this.ref,
		js.Bool(bool(val)),
	)
}

type DelayOptions struct {
	// MaxDelayTime is "DelayOptions.maxDelayTime"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_MaxDelayTime MUST be set to true to make this field effective.
	MaxDelayTime float64
	// DelayTime is "DelayOptions.delayTime"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_DelayTime MUST be set to true to make this field effective.
	DelayTime float64
	// ChannelCount is "DelayOptions.channelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
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

// New creates a new DelayOptions in the application heap.
func (p DelayOptions) New() js.Ref {
	return bindings.DelayOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DelayOptions) UpdateFrom(ref js.Ref) {
	bindings.DelayOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DelayOptions) Update(ref js.Ref) {
	bindings.DelayOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DelayOptions) FreeMembers(recursive bool) {
}

func NewDelayNode(context BaseAudioContext, options DelayOptions) (ret DelayNode) {
	ret.ref = bindings.NewDelayNodeByDelayNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewDelayNodeByDelayNode1(context BaseAudioContext) (ret DelayNode) {
	ret.ref = bindings.NewDelayNodeByDelayNode1(
		context.Ref())
	return
}

type DelayNode struct {
	AudioNode
}

func (this DelayNode) Once() DelayNode {
	this.ref.Once()
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
	this.ref.Free()
}

// DelayTime returns the value of property "DelayNode.delayTime".
//
// It returns ok=false if there is no such property.
func (this DelayNode) DelayTime() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetDelayNodeDelayTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

type DynamicsCompressorOptions struct {
	// Attack is "DynamicsCompressorOptions.attack"
	//
	// Optional, defaults to 0.003.
	//
	// NOTE: FFI_USE_Attack MUST be set to true to make this field effective.
	Attack float32
	// Knee is "DynamicsCompressorOptions.knee"
	//
	// Optional, defaults to 30.
	//
	// NOTE: FFI_USE_Knee MUST be set to true to make this field effective.
	Knee float32
	// Ratio is "DynamicsCompressorOptions.ratio"
	//
	// Optional, defaults to 12.
	//
	// NOTE: FFI_USE_Ratio MUST be set to true to make this field effective.
	Ratio float32
	// Release is "DynamicsCompressorOptions.release"
	//
	// Optional, defaults to 0.25.
	//
	// NOTE: FFI_USE_Release MUST be set to true to make this field effective.
	Release float32
	// Threshold is "DynamicsCompressorOptions.threshold"
	//
	// Optional, defaults to -24.
	//
	// NOTE: FFI_USE_Threshold MUST be set to true to make this field effective.
	Threshold float32
	// ChannelCount is "DynamicsCompressorOptions.channelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
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

// New creates a new DynamicsCompressorOptions in the application heap.
func (p DynamicsCompressorOptions) New() js.Ref {
	return bindings.DynamicsCompressorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DynamicsCompressorOptions) UpdateFrom(ref js.Ref) {
	bindings.DynamicsCompressorOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DynamicsCompressorOptions) Update(ref js.Ref) {
	bindings.DynamicsCompressorOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DynamicsCompressorOptions) FreeMembers(recursive bool) {
}

func NewDynamicsCompressorNode(context BaseAudioContext, options DynamicsCompressorOptions) (ret DynamicsCompressorNode) {
	ret.ref = bindings.NewDynamicsCompressorNodeByDynamicsCompressorNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewDynamicsCompressorNodeByDynamicsCompressorNode1(context BaseAudioContext) (ret DynamicsCompressorNode) {
	ret.ref = bindings.NewDynamicsCompressorNodeByDynamicsCompressorNode1(
		context.Ref())
	return
}

type DynamicsCompressorNode struct {
	AudioNode
}

func (this DynamicsCompressorNode) Once() DynamicsCompressorNode {
	this.ref.Once()
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
	this.ref.Free()
}

// Threshold returns the value of property "DynamicsCompressorNode.threshold".
//
// It returns ok=false if there is no such property.
func (this DynamicsCompressorNode) Threshold() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetDynamicsCompressorNodeThreshold(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Knee returns the value of property "DynamicsCompressorNode.knee".
//
// It returns ok=false if there is no such property.
func (this DynamicsCompressorNode) Knee() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetDynamicsCompressorNodeKnee(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Ratio returns the value of property "DynamicsCompressorNode.ratio".
//
// It returns ok=false if there is no such property.
func (this DynamicsCompressorNode) Ratio() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetDynamicsCompressorNodeRatio(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Reduction returns the value of property "DynamicsCompressorNode.reduction".
//
// It returns ok=false if there is no such property.
func (this DynamicsCompressorNode) Reduction() (ret float32, ok bool) {
	ok = js.True == bindings.GetDynamicsCompressorNodeReduction(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Attack returns the value of property "DynamicsCompressorNode.attack".
//
// It returns ok=false if there is no such property.
func (this DynamicsCompressorNode) Attack() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetDynamicsCompressorNodeAttack(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Release returns the value of property "DynamicsCompressorNode.release".
//
// It returns ok=false if there is no such property.
func (this DynamicsCompressorNode) Release() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetDynamicsCompressorNodeRelease(
		this.ref, js.Pointer(&ret),
	)
	return
}

type GainOptions struct {
	// Gain is "GainOptions.gain"
	//
	// Optional, defaults to 1.0.
	//
	// NOTE: FFI_USE_Gain MUST be set to true to make this field effective.
	Gain float32
	// ChannelCount is "GainOptions.channelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
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

// New creates a new GainOptions in the application heap.
func (p GainOptions) New() js.Ref {
	return bindings.GainOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GainOptions) UpdateFrom(ref js.Ref) {
	bindings.GainOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GainOptions) Update(ref js.Ref) {
	bindings.GainOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GainOptions) FreeMembers(recursive bool) {
}

func NewGainNode(context BaseAudioContext, options GainOptions) (ret GainNode) {
	ret.ref = bindings.NewGainNodeByGainNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewGainNodeByGainNode1(context BaseAudioContext) (ret GainNode) {
	ret.ref = bindings.NewGainNodeByGainNode1(
		context.Ref())
	return
}

type GainNode struct {
	AudioNode
}

func (this GainNode) Once() GainNode {
	this.ref.Once()
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
	this.ref.Free()
}

// Gain returns the value of property "GainNode.gain".
//
// It returns ok=false if there is no such property.
func (this GainNode) Gain() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetGainNodeGain(
		this.ref, js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
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

// New creates a new IIRFilterOptions in the application heap.
func (p IIRFilterOptions) New() js.Ref {
	return bindings.IIRFilterOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IIRFilterOptions) UpdateFrom(ref js.Ref) {
	bindings.IIRFilterOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IIRFilterOptions) Update(ref js.Ref) {
	bindings.IIRFilterOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IIRFilterOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Feedforward.Ref(),
		p.Feedback.Ref(),
	)
	p.Feedforward = p.Feedforward.FromRef(js.Undefined)
	p.Feedback = p.Feedback.FromRef(js.Undefined)
}

func NewIIRFilterNode(context BaseAudioContext, options IIRFilterOptions) (ret IIRFilterNode) {
	ret.ref = bindings.NewIIRFilterNodeByIIRFilterNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

type IIRFilterNode struct {
	AudioNode
}

func (this IIRFilterNode) Once() IIRFilterNode {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncGetFrequencyResponse returns true if the method "IIRFilterNode.getFrequencyResponse" exists.
func (this IIRFilterNode) HasFuncGetFrequencyResponse() bool {
	return js.True == bindings.HasFuncIIRFilterNodeGetFrequencyResponse(
		this.ref,
	)
}

// FuncGetFrequencyResponse returns the method "IIRFilterNode.getFrequencyResponse".
func (this IIRFilterNode) FuncGetFrequencyResponse() (fn js.Func[func(frequencyHz js.TypedArray[float32], magResponse js.TypedArray[float32], phaseResponse js.TypedArray[float32])]) {
	bindings.FuncIIRFilterNodeGetFrequencyResponse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFrequencyResponse calls the method "IIRFilterNode.getFrequencyResponse".
func (this IIRFilterNode) GetFrequencyResponse(frequencyHz js.TypedArray[float32], magResponse js.TypedArray[float32], phaseResponse js.TypedArray[float32]) (ret js.Void) {
	bindings.CallIIRFilterNodeGetFrequencyResponse(
		this.ref, js.Pointer(&ret),
		frequencyHz.Ref(),
		magResponse.Ref(),
		phaseResponse.Ref(),
	)

	return
}

// TryGetFrequencyResponse calls the method "IIRFilterNode.getFrequencyResponse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IIRFilterNode) TryGetFrequencyResponse(frequencyHz js.TypedArray[float32], magResponse js.TypedArray[float32], phaseResponse js.TypedArray[float32]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIIRFilterNodeGetFrequencyResponse(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		frequencyHz.Ref(),
		magResponse.Ref(),
		phaseResponse.Ref(),
	)

	return
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
	//
	// NOTE: FFI_USE_DisableNormalization MUST be set to true to make this field effective.
	DisableNormalization bool

	FFI_USE_DisableNormalization bool // for DisableNormalization.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PeriodicWaveOptions with all fields set.
func (p PeriodicWaveOptions) FromRef(ref js.Ref) PeriodicWaveOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PeriodicWaveOptions in the application heap.
func (p PeriodicWaveOptions) New() js.Ref {
	return bindings.PeriodicWaveOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PeriodicWaveOptions) UpdateFrom(ref js.Ref) {
	bindings.PeriodicWaveOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PeriodicWaveOptions) Update(ref js.Ref) {
	bindings.PeriodicWaveOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PeriodicWaveOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Real.Ref(),
		p.Imag.Ref(),
	)
	p.Real = p.Real.FromRef(js.Undefined)
	p.Imag = p.Imag.FromRef(js.Undefined)
}

func NewPeriodicWave(context BaseAudioContext, options PeriodicWaveOptions) (ret PeriodicWave) {
	ret.ref = bindings.NewPeriodicWaveByPeriodicWave(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewPeriodicWaveByPeriodicWave1(context BaseAudioContext) (ret PeriodicWave) {
	ret.ref = bindings.NewPeriodicWaveByPeriodicWave1(
		context.Ref())
	return
}

type PeriodicWave struct {
	ref js.Ref
}

func (this PeriodicWave) Once() PeriodicWave {
	this.ref.Once()
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
	this.ref.Free()
}

type OscillatorOptions struct {
	// Type is "OscillatorOptions.type"
	//
	// Optional, defaults to "sine".
	Type OscillatorType
	// Frequency is "OscillatorOptions.frequency"
	//
	// Optional, defaults to 440.
	//
	// NOTE: FFI_USE_Frequency MUST be set to true to make this field effective.
	Frequency float32
	// Detune is "OscillatorOptions.detune"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Detune MUST be set to true to make this field effective.
	Detune float32
	// PeriodicWave is "OscillatorOptions.periodicWave"
	//
	// Optional
	PeriodicWave PeriodicWave
	// ChannelCount is "OscillatorOptions.channelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
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

// New creates a new OscillatorOptions in the application heap.
func (p OscillatorOptions) New() js.Ref {
	return bindings.OscillatorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OscillatorOptions) UpdateFrom(ref js.Ref) {
	bindings.OscillatorOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OscillatorOptions) Update(ref js.Ref) {
	bindings.OscillatorOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OscillatorOptions) FreeMembers(recursive bool) {
	js.Free(
		p.PeriodicWave.Ref(),
	)
	p.PeriodicWave = p.PeriodicWave.FromRef(js.Undefined)
}

func NewOscillatorNode(context BaseAudioContext, options OscillatorOptions) (ret OscillatorNode) {
	ret.ref = bindings.NewOscillatorNodeByOscillatorNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewOscillatorNodeByOscillatorNode1(context BaseAudioContext) (ret OscillatorNode) {
	ret.ref = bindings.NewOscillatorNodeByOscillatorNode1(
		context.Ref())
	return
}

type OscillatorNode struct {
	AudioScheduledSourceNode
}

func (this OscillatorNode) Once() OscillatorNode {
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "OscillatorNode.type".
//
// It returns ok=false if there is no such property.
func (this OscillatorNode) Type() (ret OscillatorType, ok bool) {
	ok = js.True == bindings.GetOscillatorNodeType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "OscillatorNode.type" to val.
//
// It returns false if the property cannot be set.
func (this OscillatorNode) SetType(val OscillatorType) bool {
	return js.True == bindings.SetOscillatorNodeType(
		this.ref,
		uint32(val),
	)
}

// Frequency returns the value of property "OscillatorNode.frequency".
//
// It returns ok=false if there is no such property.
func (this OscillatorNode) Frequency() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetOscillatorNodeFrequency(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Detune returns the value of property "OscillatorNode.detune".
//
// It returns ok=false if there is no such property.
func (this OscillatorNode) Detune() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetOscillatorNodeDetune(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncSetPeriodicWave returns true if the method "OscillatorNode.setPeriodicWave" exists.
func (this OscillatorNode) HasFuncSetPeriodicWave() bool {
	return js.True == bindings.HasFuncOscillatorNodeSetPeriodicWave(
		this.ref,
	)
}

// FuncSetPeriodicWave returns the method "OscillatorNode.setPeriodicWave".
func (this OscillatorNode) FuncSetPeriodicWave() (fn js.Func[func(periodicWave PeriodicWave)]) {
	bindings.FuncOscillatorNodeSetPeriodicWave(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetPeriodicWave calls the method "OscillatorNode.setPeriodicWave".
func (this OscillatorNode) SetPeriodicWave(periodicWave PeriodicWave) (ret js.Void) {
	bindings.CallOscillatorNodeSetPeriodicWave(
		this.ref, js.Pointer(&ret),
		periodicWave.Ref(),
	)

	return
}

// TrySetPeriodicWave calls the method "OscillatorNode.setPeriodicWave"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this OscillatorNode) TrySetPeriodicWave(periodicWave PeriodicWave) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOscillatorNodeSetPeriodicWave(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		periodicWave.Ref(),
	)

	return
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
	//
	// NOTE: FFI_USE_PositionX MUST be set to true to make this field effective.
	PositionX float32
	// PositionY is "PannerOptions.positionY"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_PositionY MUST be set to true to make this field effective.
	PositionY float32
	// PositionZ is "PannerOptions.positionZ"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_PositionZ MUST be set to true to make this field effective.
	PositionZ float32
	// OrientationX is "PannerOptions.orientationX"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_OrientationX MUST be set to true to make this field effective.
	OrientationX float32
	// OrientationY is "PannerOptions.orientationY"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_OrientationY MUST be set to true to make this field effective.
	OrientationY float32
	// OrientationZ is "PannerOptions.orientationZ"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_OrientationZ MUST be set to true to make this field effective.
	OrientationZ float32
	// RefDistance is "PannerOptions.refDistance"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_RefDistance MUST be set to true to make this field effective.
	RefDistance float64
	// MaxDistance is "PannerOptions.maxDistance"
	//
	// Optional, defaults to 10000.
	//
	// NOTE: FFI_USE_MaxDistance MUST be set to true to make this field effective.
	MaxDistance float64
	// RolloffFactor is "PannerOptions.rolloffFactor"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_RolloffFactor MUST be set to true to make this field effective.
	RolloffFactor float64
	// ConeInnerAngle is "PannerOptions.coneInnerAngle"
	//
	// Optional, defaults to 360.
	//
	// NOTE: FFI_USE_ConeInnerAngle MUST be set to true to make this field effective.
	ConeInnerAngle float64
	// ConeOuterAngle is "PannerOptions.coneOuterAngle"
	//
	// Optional, defaults to 360.
	//
	// NOTE: FFI_USE_ConeOuterAngle MUST be set to true to make this field effective.
	ConeOuterAngle float64
	// ConeOuterGain is "PannerOptions.coneOuterGain"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ConeOuterGain MUST be set to true to make this field effective.
	ConeOuterGain float64
	// ChannelCount is "PannerOptions.channelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
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

// New creates a new PannerOptions in the application heap.
func (p PannerOptions) New() js.Ref {
	return bindings.PannerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PannerOptions) UpdateFrom(ref js.Ref) {
	bindings.PannerOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PannerOptions) Update(ref js.Ref) {
	bindings.PannerOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PannerOptions) FreeMembers(recursive bool) {
}

func NewPannerNode(context BaseAudioContext, options PannerOptions) (ret PannerNode) {
	ret.ref = bindings.NewPannerNodeByPannerNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewPannerNodeByPannerNode1(context BaseAudioContext) (ret PannerNode) {
	ret.ref = bindings.NewPannerNodeByPannerNode1(
		context.Ref())
	return
}

type PannerNode struct {
	AudioNode
}

func (this PannerNode) Once() PannerNode {
	this.ref.Once()
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
	this.ref.Free()
}

// PanningModel returns the value of property "PannerNode.panningModel".
//
// It returns ok=false if there is no such property.
func (this PannerNode) PanningModel() (ret PanningModelType, ok bool) {
	ok = js.True == bindings.GetPannerNodePanningModel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPanningModel sets the value of property "PannerNode.panningModel" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetPanningModel(val PanningModelType) bool {
	return js.True == bindings.SetPannerNodePanningModel(
		this.ref,
		uint32(val),
	)
}

// PositionX returns the value of property "PannerNode.positionX".
//
// It returns ok=false if there is no such property.
func (this PannerNode) PositionX() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetPannerNodePositionX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PositionY returns the value of property "PannerNode.positionY".
//
// It returns ok=false if there is no such property.
func (this PannerNode) PositionY() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetPannerNodePositionY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PositionZ returns the value of property "PannerNode.positionZ".
//
// It returns ok=false if there is no such property.
func (this PannerNode) PositionZ() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetPannerNodePositionZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OrientationX returns the value of property "PannerNode.orientationX".
//
// It returns ok=false if there is no such property.
func (this PannerNode) OrientationX() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetPannerNodeOrientationX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OrientationY returns the value of property "PannerNode.orientationY".
//
// It returns ok=false if there is no such property.
func (this PannerNode) OrientationY() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetPannerNodeOrientationY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OrientationZ returns the value of property "PannerNode.orientationZ".
//
// It returns ok=false if there is no such property.
func (this PannerNode) OrientationZ() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetPannerNodeOrientationZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DistanceModel returns the value of property "PannerNode.distanceModel".
//
// It returns ok=false if there is no such property.
func (this PannerNode) DistanceModel() (ret DistanceModelType, ok bool) {
	ok = js.True == bindings.GetPannerNodeDistanceModel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDistanceModel sets the value of property "PannerNode.distanceModel" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetDistanceModel(val DistanceModelType) bool {
	return js.True == bindings.SetPannerNodeDistanceModel(
		this.ref,
		uint32(val),
	)
}

// RefDistance returns the value of property "PannerNode.refDistance".
//
// It returns ok=false if there is no such property.
func (this PannerNode) RefDistance() (ret float64, ok bool) {
	ok = js.True == bindings.GetPannerNodeRefDistance(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRefDistance sets the value of property "PannerNode.refDistance" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetRefDistance(val float64) bool {
	return js.True == bindings.SetPannerNodeRefDistance(
		this.ref,
		float64(val),
	)
}

// MaxDistance returns the value of property "PannerNode.maxDistance".
//
// It returns ok=false if there is no such property.
func (this PannerNode) MaxDistance() (ret float64, ok bool) {
	ok = js.True == bindings.GetPannerNodeMaxDistance(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMaxDistance sets the value of property "PannerNode.maxDistance" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetMaxDistance(val float64) bool {
	return js.True == bindings.SetPannerNodeMaxDistance(
		this.ref,
		float64(val),
	)
}

// RolloffFactor returns the value of property "PannerNode.rolloffFactor".
//
// It returns ok=false if there is no such property.
func (this PannerNode) RolloffFactor() (ret float64, ok bool) {
	ok = js.True == bindings.GetPannerNodeRolloffFactor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRolloffFactor sets the value of property "PannerNode.rolloffFactor" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetRolloffFactor(val float64) bool {
	return js.True == bindings.SetPannerNodeRolloffFactor(
		this.ref,
		float64(val),
	)
}

// ConeInnerAngle returns the value of property "PannerNode.coneInnerAngle".
//
// It returns ok=false if there is no such property.
func (this PannerNode) ConeInnerAngle() (ret float64, ok bool) {
	ok = js.True == bindings.GetPannerNodeConeInnerAngle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetConeInnerAngle sets the value of property "PannerNode.coneInnerAngle" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetConeInnerAngle(val float64) bool {
	return js.True == bindings.SetPannerNodeConeInnerAngle(
		this.ref,
		float64(val),
	)
}

// ConeOuterAngle returns the value of property "PannerNode.coneOuterAngle".
//
// It returns ok=false if there is no such property.
func (this PannerNode) ConeOuterAngle() (ret float64, ok bool) {
	ok = js.True == bindings.GetPannerNodeConeOuterAngle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetConeOuterAngle sets the value of property "PannerNode.coneOuterAngle" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetConeOuterAngle(val float64) bool {
	return js.True == bindings.SetPannerNodeConeOuterAngle(
		this.ref,
		float64(val),
	)
}

// ConeOuterGain returns the value of property "PannerNode.coneOuterGain".
//
// It returns ok=false if there is no such property.
func (this PannerNode) ConeOuterGain() (ret float64, ok bool) {
	ok = js.True == bindings.GetPannerNodeConeOuterGain(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetConeOuterGain sets the value of property "PannerNode.coneOuterGain" to val.
//
// It returns false if the property cannot be set.
func (this PannerNode) SetConeOuterGain(val float64) bool {
	return js.True == bindings.SetPannerNodeConeOuterGain(
		this.ref,
		float64(val),
	)
}

// HasFuncSetPosition returns true if the method "PannerNode.setPosition" exists.
func (this PannerNode) HasFuncSetPosition() bool {
	return js.True == bindings.HasFuncPannerNodeSetPosition(
		this.ref,
	)
}

// FuncSetPosition returns the method "PannerNode.setPosition".
func (this PannerNode) FuncSetPosition() (fn js.Func[func(x float32, y float32, z float32)]) {
	bindings.FuncPannerNodeSetPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetPosition calls the method "PannerNode.setPosition".
func (this PannerNode) SetPosition(x float32, y float32, z float32) (ret js.Void) {
	bindings.CallPannerNodeSetPosition(
		this.ref, js.Pointer(&ret),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// TrySetPosition calls the method "PannerNode.setPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PannerNode) TrySetPosition(x float32, y float32, z float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPannerNodeSetPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// HasFuncSetOrientation returns true if the method "PannerNode.setOrientation" exists.
func (this PannerNode) HasFuncSetOrientation() bool {
	return js.True == bindings.HasFuncPannerNodeSetOrientation(
		this.ref,
	)
}

// FuncSetOrientation returns the method "PannerNode.setOrientation".
func (this PannerNode) FuncSetOrientation() (fn js.Func[func(x float32, y float32, z float32)]) {
	bindings.FuncPannerNodeSetOrientation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetOrientation calls the method "PannerNode.setOrientation".
func (this PannerNode) SetOrientation(x float32, y float32, z float32) (ret js.Void) {
	bindings.CallPannerNodeSetOrientation(
		this.ref, js.Pointer(&ret),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// TrySetOrientation calls the method "PannerNode.setOrientation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PannerNode) TrySetOrientation(x float32, y float32, z float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPannerNodeSetOrientation(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

type PeriodicWaveConstraints struct {
	// DisableNormalization is "PeriodicWaveConstraints.disableNormalization"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_DisableNormalization MUST be set to true to make this field effective.
	DisableNormalization bool

	FFI_USE_DisableNormalization bool // for DisableNormalization.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PeriodicWaveConstraints with all fields set.
func (p PeriodicWaveConstraints) FromRef(ref js.Ref) PeriodicWaveConstraints {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PeriodicWaveConstraints in the application heap.
func (p PeriodicWaveConstraints) New() js.Ref {
	return bindings.PeriodicWaveConstraintsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PeriodicWaveConstraints) UpdateFrom(ref js.Ref) {
	bindings.PeriodicWaveConstraintsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PeriodicWaveConstraints) Update(ref js.Ref) {
	bindings.PeriodicWaveConstraintsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PeriodicWaveConstraints) FreeMembers(recursive bool) {
}

type ScriptProcessorNode struct {
	AudioNode
}

func (this ScriptProcessorNode) Once() ScriptProcessorNode {
	this.ref.Once()
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
	this.ref.Free()
}

// BufferSize returns the value of property "ScriptProcessorNode.bufferSize".
//
// It returns ok=false if there is no such property.
func (this ScriptProcessorNode) BufferSize() (ret int32, ok bool) {
	ok = js.True == bindings.GetScriptProcessorNodeBufferSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

type StereoPannerOptions struct {
	// Pan is "StereoPannerOptions.pan"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Pan MUST be set to true to make this field effective.
	Pan float32
	// ChannelCount is "StereoPannerOptions.channelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
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

// New creates a new StereoPannerOptions in the application heap.
func (p StereoPannerOptions) New() js.Ref {
	return bindings.StereoPannerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StereoPannerOptions) UpdateFrom(ref js.Ref) {
	bindings.StereoPannerOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StereoPannerOptions) Update(ref js.Ref) {
	bindings.StereoPannerOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StereoPannerOptions) FreeMembers(recursive bool) {
}

func NewStereoPannerNode(context BaseAudioContext, options StereoPannerOptions) (ret StereoPannerNode) {
	ret.ref = bindings.NewStereoPannerNodeByStereoPannerNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewStereoPannerNodeByStereoPannerNode1(context BaseAudioContext) (ret StereoPannerNode) {
	ret.ref = bindings.NewStereoPannerNodeByStereoPannerNode1(
		context.Ref())
	return
}

type StereoPannerNode struct {
	AudioNode
}

func (this StereoPannerNode) Once() StereoPannerNode {
	this.ref.Once()
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
	this.ref.Free()
}

// Pan returns the value of property "StereoPannerNode.pan".
//
// It returns ok=false if there is no such property.
func (this StereoPannerNode) Pan() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetStereoPannerNodePan(
		this.ref, js.Pointer(&ret),
	)
	return
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

type WaveShaperOptions struct {
	// Curve is "WaveShaperOptions.curve"
	//
	// Optional
	Curve js.Array[float32]
	// Oversample is "WaveShaperOptions.oversample"
	//
	// Optional, defaults to "none".
	Oversample OverSampleType
	// ChannelCount is "WaveShaperOptions.channelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
	ChannelCount uint32
	// ChannelCountMode is "WaveShaperOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "WaveShaperOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_ChannelCount bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WaveShaperOptions with all fields set.
func (p WaveShaperOptions) FromRef(ref js.Ref) WaveShaperOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WaveShaperOptions in the application heap.
func (p WaveShaperOptions) New() js.Ref {
	return bindings.WaveShaperOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WaveShaperOptions) UpdateFrom(ref js.Ref) {
	bindings.WaveShaperOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WaveShaperOptions) Update(ref js.Ref) {
	bindings.WaveShaperOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WaveShaperOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Curve.Ref(),
	)
	p.Curve = p.Curve.FromRef(js.Undefined)
}

func NewWaveShaperNode(context BaseAudioContext, options WaveShaperOptions) (ret WaveShaperNode) {
	ret.ref = bindings.NewWaveShaperNodeByWaveShaperNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewWaveShaperNodeByWaveShaperNode1(context BaseAudioContext) (ret WaveShaperNode) {
	ret.ref = bindings.NewWaveShaperNodeByWaveShaperNode1(
		context.Ref())
	return
}

type WaveShaperNode struct {
	AudioNode
}

func (this WaveShaperNode) Once() WaveShaperNode {
	this.ref.Once()
	return this
}

func (this WaveShaperNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this WaveShaperNode) FromRef(ref js.Ref) WaveShaperNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this WaveShaperNode) Free() {
	this.ref.Free()
}

// Curve returns the value of property "WaveShaperNode.curve".
//
// It returns ok=false if there is no such property.
func (this WaveShaperNode) Curve() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetWaveShaperNodeCurve(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCurve sets the value of property "WaveShaperNode.curve" to val.
//
// It returns false if the property cannot be set.
func (this WaveShaperNode) SetCurve(val js.TypedArray[float32]) bool {
	return js.True == bindings.SetWaveShaperNodeCurve(
		this.ref,
		val.Ref(),
	)
}

// Oversample returns the value of property "WaveShaperNode.oversample".
//
// It returns ok=false if there is no such property.
func (this WaveShaperNode) Oversample() (ret OverSampleType, ok bool) {
	ok = js.True == bindings.GetWaveShaperNodeOversample(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetOversample sets the value of property "WaveShaperNode.oversample" to val.
//
// It returns false if the property cannot be set.
func (this WaveShaperNode) SetOversample(val OverSampleType) bool {
	return js.True == bindings.SetWaveShaperNodeOversample(
		this.ref,
		uint32(val),
	)
}
