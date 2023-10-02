// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package bindings

import (
	"unsafe"

	"github.com/primecitizens/std/ffi/js"
)

func _() {
	var (
		_ js.Void
		_ unsafe.Pointer
	)
}

//go:wasmimport plat/js/web new_AmbientLightSensor_AmbientLightSensor
//go:noescape

func NewAmbientLightSensorByAmbientLightSensor(
	sensorOptions unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_AmbientLightSensor_AmbientLightSensor1
//go:noescape

func NewAmbientLightSensorByAmbientLightSensor1() js.Ref

//go:wasmimport plat/js/web get_AmbientLightSensor_Illuminance
//go:noescape

func GetAmbientLightSensorIlluminance(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web constof_BiquadFilterType
//go:noescape

func ConstOfBiquadFilterType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_ChannelCountMode
//go:noescape

func ConstOfChannelCountMode(str js.Ref) uint32

//go:wasmimport plat/js/web constof_ChannelInterpretation
//go:noescape

func ConstOfChannelInterpretation(str js.Ref) uint32

//go:wasmimport plat/js/web store_BiquadFilterOptions
//go:noescape

func BiquadFilterOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BiquadFilterOptions
//go:noescape

func BiquadFilterOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_AutomationRate
//go:noescape

func ConstOfAutomationRate(str js.Ref) uint32

//go:wasmimport plat/js/web get_AudioParam_Value
//go:noescape

func GetAudioParamValue(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_AudioParam_Value
//go:noescape

func SetAudioParamValue(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_AudioParam_AutomationRate
//go:noescape

func GetAudioParamAutomationRate(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_AudioParam_AutomationRate
//go:noescape

func SetAudioParamAutomationRate(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_AudioParam_DefaultValue
//go:noescape

func GetAudioParamDefaultValue(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_AudioParam_MinValue
//go:noescape

func GetAudioParamMinValue(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_AudioParam_MaxValue
//go:noescape

func GetAudioParamMaxValue(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web call_AudioParam_SetValueAtTime
//go:noescape

func CallAudioParamSetValueAtTime(
	this js.Ref,
	ptr unsafe.Pointer,

	value float32,
	startTime float64,
) js.Ref

//go:wasmimport plat/js/web func_AudioParam_SetValueAtTime
//go:noescape

func AudioParamSetValueAtTimeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioParam_LinearRampToValueAtTime
//go:noescape

func CallAudioParamLinearRampToValueAtTime(
	this js.Ref,
	ptr unsafe.Pointer,

	value float32,
	endTime float64,
) js.Ref

//go:wasmimport plat/js/web func_AudioParam_LinearRampToValueAtTime
//go:noescape

func AudioParamLinearRampToValueAtTimeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioParam_ExponentialRampToValueAtTime
//go:noescape

func CallAudioParamExponentialRampToValueAtTime(
	this js.Ref,
	ptr unsafe.Pointer,

	value float32,
	endTime float64,
) js.Ref

//go:wasmimport plat/js/web func_AudioParam_ExponentialRampToValueAtTime
//go:noescape

func AudioParamExponentialRampToValueAtTimeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioParam_SetTargetAtTime
//go:noescape

func CallAudioParamSetTargetAtTime(
	this js.Ref,
	ptr unsafe.Pointer,

	target float32,
	startTime float64,
	timeConstant float32,
) js.Ref

//go:wasmimport plat/js/web func_AudioParam_SetTargetAtTime
//go:noescape

func AudioParamSetTargetAtTimeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioParam_SetValueCurveAtTime
//go:noescape

func CallAudioParamSetValueCurveAtTime(
	this js.Ref,
	ptr unsafe.Pointer,

	values js.Ref,
	startTime float64,
	duration float64,
) js.Ref

//go:wasmimport plat/js/web func_AudioParam_SetValueCurveAtTime
//go:noescape

func AudioParamSetValueCurveAtTimeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioParam_CancelScheduledValues
//go:noescape

func CallAudioParamCancelScheduledValues(
	this js.Ref,
	ptr unsafe.Pointer,

	cancelTime float64,
) js.Ref

//go:wasmimport plat/js/web func_AudioParam_CancelScheduledValues
//go:noescape

func AudioParamCancelScheduledValuesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioParam_CancelAndHoldAtTime
//go:noescape

func CallAudioParamCancelAndHoldAtTime(
	this js.Ref,
	ptr unsafe.Pointer,

	cancelTime float64,
) js.Ref

//go:wasmimport plat/js/web func_AudioParam_CancelAndHoldAtTime
//go:noescape

func AudioParamCancelAndHoldAtTimeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_BiquadFilterNode_BiquadFilterNode
//go:noescape

func NewBiquadFilterNodeByBiquadFilterNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_BiquadFilterNode_BiquadFilterNode1
//go:noescape

func NewBiquadFilterNodeByBiquadFilterNode1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_BiquadFilterNode_Type
//go:noescape

func GetBiquadFilterNodeType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_BiquadFilterNode_Type
//go:noescape

func SetBiquadFilterNodeType(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_BiquadFilterNode_Frequency
//go:noescape

func GetBiquadFilterNodeFrequency(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_BiquadFilterNode_Detune
//go:noescape

func GetBiquadFilterNodeDetune(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_BiquadFilterNode_Q
//go:noescape

func GetBiquadFilterNodeQ(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_BiquadFilterNode_Gain
//go:noescape

func GetBiquadFilterNodeGain(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_BiquadFilterNode_GetFrequencyResponse
//go:noescape

func CallBiquadFilterNodeGetFrequencyResponse(
	this js.Ref,
	ptr unsafe.Pointer,

	frequencyHz js.Ref,
	magResponse js.Ref,
	phaseResponse js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_BiquadFilterNode_GetFrequencyResponse
//go:noescape

func BiquadFilterNodeGetFrequencyResponseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_AudioBufferOptions
//go:noescape

func AudioBufferOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioBufferOptions
//go:noescape

func AudioBufferOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_AudioBuffer_AudioBuffer
//go:noescape

func NewAudioBufferByAudioBuffer(
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_AudioBuffer_SampleRate
//go:noescape

func GetAudioBufferSampleRate(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_AudioBuffer_Length
//go:noescape

func GetAudioBufferLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_AudioBuffer_Duration
//go:noescape

func GetAudioBufferDuration(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_AudioBuffer_NumberOfChannels
//go:noescape

func GetAudioBufferNumberOfChannels(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_AudioBuffer_GetChannelData
//go:noescape

func CallAudioBufferGetChannelData(
	this js.Ref,
	ptr unsafe.Pointer,

	channel uint32,
) js.Ref

//go:wasmimport plat/js/web func_AudioBuffer_GetChannelData
//go:noescape

func AudioBufferGetChannelDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioBuffer_CopyFromChannel
//go:noescape

func CallAudioBufferCopyFromChannel(
	this js.Ref,
	ptr unsafe.Pointer,

	destination js.Ref,
	channelNumber uint32,
	bufferOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_AudioBuffer_CopyFromChannel
//go:noescape

func AudioBufferCopyFromChannelFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioBuffer_CopyFromChannel1
//go:noescape

func CallAudioBufferCopyFromChannel1(
	this js.Ref,
	ptr unsafe.Pointer,

	destination js.Ref,
	channelNumber uint32,
) js.Ref

//go:wasmimport plat/js/web func_AudioBuffer_CopyFromChannel1
//go:noescape

func AudioBufferCopyFromChannel1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioBuffer_CopyToChannel
//go:noescape

func CallAudioBufferCopyToChannel(
	this js.Ref,
	ptr unsafe.Pointer,

	source js.Ref,
	channelNumber uint32,
	bufferOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_AudioBuffer_CopyToChannel
//go:noescape

func AudioBufferCopyToChannelFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioBuffer_CopyToChannel1
//go:noescape

func CallAudioBufferCopyToChannel1(
	this js.Ref,
	ptr unsafe.Pointer,

	source js.Ref,
	channelNumber uint32,
) js.Ref

//go:wasmimport plat/js/web func_AudioBuffer_CopyToChannel1
//go:noescape

func AudioBufferCopyToChannel1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_AudioBufferSourceOptions
//go:noescape

func AudioBufferSourceOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioBufferSourceOptions
//go:noescape

func AudioBufferSourceOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_AudioBufferSourceNode_AudioBufferSourceNode
//go:noescape

func NewAudioBufferSourceNodeByAudioBufferSourceNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_AudioBufferSourceNode_AudioBufferSourceNode1
//go:noescape

func NewAudioBufferSourceNodeByAudioBufferSourceNode1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_AudioBufferSourceNode_Buffer
//go:noescape

func GetAudioBufferSourceNodeBuffer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_AudioBufferSourceNode_Buffer
//go:noescape

func SetAudioBufferSourceNodeBuffer(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_AudioBufferSourceNode_PlaybackRate
//go:noescape

func GetAudioBufferSourceNodePlaybackRate(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_AudioBufferSourceNode_Detune
//go:noescape

func GetAudioBufferSourceNodeDetune(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_AudioBufferSourceNode_Loop
//go:noescape

func GetAudioBufferSourceNodeLoop(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_AudioBufferSourceNode_Loop
//go:noescape

func SetAudioBufferSourceNodeLoop(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_AudioBufferSourceNode_LoopStart
//go:noescape

func GetAudioBufferSourceNodeLoopStart(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_AudioBufferSourceNode_LoopStart
//go:noescape

func SetAudioBufferSourceNodeLoopStart(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_AudioBufferSourceNode_LoopEnd
//go:noescape

func GetAudioBufferSourceNodeLoopEnd(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_AudioBufferSourceNode_LoopEnd
//go:noescape

func SetAudioBufferSourceNodeLoopEnd(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web call_AudioBufferSourceNode_Start
//go:noescape

func CallAudioBufferSourceNodeStart(
	this js.Ref,
	ptr unsafe.Pointer,

	when float64,
	offset float64,
	duration float64,
) js.Ref

//go:wasmimport plat/js/web func_AudioBufferSourceNode_Start
//go:noescape

func AudioBufferSourceNodeStartFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioBufferSourceNode_Start1
//go:noescape

func CallAudioBufferSourceNodeStart1(
	this js.Ref,
	ptr unsafe.Pointer,

	when float64,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_AudioBufferSourceNode_Start1
//go:noescape

func AudioBufferSourceNodeStart1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioBufferSourceNode_Start2
//go:noescape

func CallAudioBufferSourceNodeStart2(
	this js.Ref,
	ptr unsafe.Pointer,

	when float64,
) js.Ref

//go:wasmimport plat/js/web func_AudioBufferSourceNode_Start2
//go:noescape

func AudioBufferSourceNodeStart2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioBufferSourceNode_Start3
//go:noescape

func CallAudioBufferSourceNodeStart3(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AudioBufferSourceNode_Start3
//go:noescape

func AudioBufferSourceNodeStart3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_ChannelMergerOptions
//go:noescape

func ChannelMergerOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ChannelMergerOptions
//go:noescape

func ChannelMergerOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ChannelMergerNode_ChannelMergerNode
//go:noescape

func NewChannelMergerNodeByChannelMergerNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ChannelMergerNode_ChannelMergerNode1
//go:noescape

func NewChannelMergerNodeByChannelMergerNode1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web store_ChannelSplitterOptions
//go:noescape

func ChannelSplitterOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ChannelSplitterOptions
//go:noescape

func ChannelSplitterOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ChannelSplitterNode_ChannelSplitterNode
//go:noescape

func NewChannelSplitterNodeByChannelSplitterNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ChannelSplitterNode_ChannelSplitterNode1
//go:noescape

func NewChannelSplitterNodeByChannelSplitterNode1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web store_ConstantSourceOptions
//go:noescape

func ConstantSourceOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ConstantSourceOptions
//go:noescape

func ConstantSourceOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ConstantSourceNode_ConstantSourceNode
//go:noescape

func NewConstantSourceNodeByConstantSourceNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ConstantSourceNode_ConstantSourceNode1
//go:noescape

func NewConstantSourceNodeByConstantSourceNode1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_ConstantSourceNode_Offset
//go:noescape

func GetConstantSourceNodeOffset(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_ConvolverOptions
//go:noescape

func ConvolverOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ConvolverOptions
//go:noescape

func ConvolverOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ConvolverNode_ConvolverNode
//go:noescape

func NewConvolverNodeByConvolverNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ConvolverNode_ConvolverNode1
//go:noescape

func NewConvolverNodeByConvolverNode1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_ConvolverNode_Buffer
//go:noescape

func GetConvolverNodeBuffer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ConvolverNode_Buffer
//go:noescape

func SetConvolverNodeBuffer(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ConvolverNode_Normalize
//go:noescape

func GetConvolverNodeNormalize(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ConvolverNode_Normalize
//go:noescape

func SetConvolverNodeNormalize(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_DelayOptions
//go:noescape

func DelayOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DelayOptions
//go:noescape

func DelayOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_DelayNode_DelayNode
//go:noescape

func NewDelayNodeByDelayNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_DelayNode_DelayNode1
//go:noescape

func NewDelayNodeByDelayNode1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_DelayNode_DelayTime
//go:noescape

func GetDelayNodeDelayTime(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_DynamicsCompressorOptions
//go:noescape

func DynamicsCompressorOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DynamicsCompressorOptions
//go:noescape

func DynamicsCompressorOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_DynamicsCompressorNode_DynamicsCompressorNode
//go:noescape

func NewDynamicsCompressorNodeByDynamicsCompressorNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_DynamicsCompressorNode_DynamicsCompressorNode1
//go:noescape

func NewDynamicsCompressorNodeByDynamicsCompressorNode1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_DynamicsCompressorNode_Threshold
//go:noescape

func GetDynamicsCompressorNodeThreshold(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_DynamicsCompressorNode_Knee
//go:noescape

func GetDynamicsCompressorNodeKnee(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_DynamicsCompressorNode_Ratio
//go:noescape

func GetDynamicsCompressorNodeRatio(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_DynamicsCompressorNode_Reduction
//go:noescape

func GetDynamicsCompressorNodeReduction(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_DynamicsCompressorNode_Attack
//go:noescape

func GetDynamicsCompressorNodeAttack(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_DynamicsCompressorNode_Release
//go:noescape

func GetDynamicsCompressorNodeRelease(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_GainOptions
//go:noescape

func GainOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GainOptions
//go:noescape

func GainOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_GainNode_GainNode
//go:noescape

func NewGainNodeByGainNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_GainNode_GainNode1
//go:noescape

func NewGainNodeByGainNode1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_GainNode_Gain
//go:noescape

func GetGainNodeGain(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_IIRFilterOptions
//go:noescape

func IIRFilterOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IIRFilterOptions
//go:noescape

func IIRFilterOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_IIRFilterNode_IIRFilterNode
//go:noescape

func NewIIRFilterNodeByIIRFilterNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web call_IIRFilterNode_GetFrequencyResponse
//go:noescape

func CallIIRFilterNodeGetFrequencyResponse(
	this js.Ref,
	ptr unsafe.Pointer,

	frequencyHz js.Ref,
	magResponse js.Ref,
	phaseResponse js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_IIRFilterNode_GetFrequencyResponse
//go:noescape

func IIRFilterNodeGetFrequencyResponseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_OscillatorType
//go:noescape

func ConstOfOscillatorType(str js.Ref) uint32

//go:wasmimport plat/js/web store_PeriodicWaveOptions
//go:noescape

func PeriodicWaveOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PeriodicWaveOptions
//go:noescape

func PeriodicWaveOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PeriodicWave_PeriodicWave
//go:noescape

func NewPeriodicWaveByPeriodicWave(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_PeriodicWave_PeriodicWave1
//go:noescape

func NewPeriodicWaveByPeriodicWave1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web store_OscillatorOptions
//go:noescape

func OscillatorOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_OscillatorOptions
//go:noescape

func OscillatorOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_OscillatorNode_OscillatorNode
//go:noescape

func NewOscillatorNodeByOscillatorNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_OscillatorNode_OscillatorNode1
//go:noescape

func NewOscillatorNodeByOscillatorNode1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_OscillatorNode_Type
//go:noescape

func GetOscillatorNodeType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_OscillatorNode_Type
//go:noescape

func SetOscillatorNodeType(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_OscillatorNode_Frequency
//go:noescape

func GetOscillatorNodeFrequency(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_OscillatorNode_Detune
//go:noescape

func GetOscillatorNodeDetune(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_OscillatorNode_SetPeriodicWave
//go:noescape

func CallOscillatorNodeSetPeriodicWave(
	this js.Ref,
	ptr unsafe.Pointer,

	periodicWave js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_OscillatorNode_SetPeriodicWave
//go:noescape

func OscillatorNodeSetPeriodicWaveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_PanningModelType
//go:noescape

func ConstOfPanningModelType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_DistanceModelType
//go:noescape

func ConstOfDistanceModelType(str js.Ref) uint32

//go:wasmimport plat/js/web store_PannerOptions
//go:noescape

func PannerOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PannerOptions
//go:noescape

func PannerOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PannerNode_PannerNode
//go:noescape

func NewPannerNodeByPannerNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_PannerNode_PannerNode1
//go:noescape

func NewPannerNodeByPannerNode1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_PannerNode_PanningModel
//go:noescape

func GetPannerNodePanningModel(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_PannerNode_PanningModel
//go:noescape

func SetPannerNodePanningModel(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_PositionX
//go:noescape

func GetPannerNodePositionX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_PositionY
//go:noescape

func GetPannerNodePositionY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_PositionZ
//go:noescape

func GetPannerNodePositionZ(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_OrientationX
//go:noescape

func GetPannerNodeOrientationX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_OrientationY
//go:noescape

func GetPannerNodeOrientationY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_OrientationZ
//go:noescape

func GetPannerNodeOrientationZ(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_DistanceModel
//go:noescape

func GetPannerNodeDistanceModel(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_PannerNode_DistanceModel
//go:noescape

func SetPannerNodeDistanceModel(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_RefDistance
//go:noescape

func GetPannerNodeRefDistance(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_PannerNode_RefDistance
//go:noescape

func SetPannerNodeRefDistance(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_MaxDistance
//go:noescape

func GetPannerNodeMaxDistance(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_PannerNode_MaxDistance
//go:noescape

func SetPannerNodeMaxDistance(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_RolloffFactor
//go:noescape

func GetPannerNodeRolloffFactor(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_PannerNode_RolloffFactor
//go:noescape

func SetPannerNodeRolloffFactor(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_ConeInnerAngle
//go:noescape

func GetPannerNodeConeInnerAngle(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_PannerNode_ConeInnerAngle
//go:noescape

func SetPannerNodeConeInnerAngle(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_ConeOuterAngle
//go:noescape

func GetPannerNodeConeOuterAngle(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_PannerNode_ConeOuterAngle
//go:noescape

func SetPannerNodeConeOuterAngle(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_ConeOuterGain
//go:noescape

func GetPannerNodeConeOuterGain(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_PannerNode_ConeOuterGain
//go:noescape

func SetPannerNodeConeOuterGain(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web call_PannerNode_SetPosition
//go:noescape

func CallPannerNodeSetPosition(
	this js.Ref,
	ptr unsafe.Pointer,

	x float32,
	y float32,
	z float32,
) js.Ref

//go:wasmimport plat/js/web func_PannerNode_SetPosition
//go:noescape

func PannerNodeSetPositionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PannerNode_SetOrientation
//go:noescape

func CallPannerNodeSetOrientation(
	this js.Ref,
	ptr unsafe.Pointer,

	x float32,
	y float32,
	z float32,
) js.Ref

//go:wasmimport plat/js/web func_PannerNode_SetOrientation
//go:noescape

func PannerNodeSetOrientationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_PeriodicWaveConstraints
//go:noescape

func PeriodicWaveConstraintsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PeriodicWaveConstraints
//go:noescape

func PeriodicWaveConstraintsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_ScriptProcessorNode_BufferSize
//go:noescape

func GetScriptProcessorNodeBufferSize(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web store_StereoPannerOptions
//go:noescape

func StereoPannerOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_StereoPannerOptions
//go:noescape

func StereoPannerOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_StereoPannerNode_StereoPannerNode
//go:noescape

func NewStereoPannerNodeByStereoPannerNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_StereoPannerNode_StereoPannerNode1
//go:noescape

func NewStereoPannerNodeByStereoPannerNode1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_StereoPannerNode_Pan
//go:noescape

func GetStereoPannerNodePan(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web constof_OverSampleType
//go:noescape

func ConstOfOverSampleType(str js.Ref) uint32
