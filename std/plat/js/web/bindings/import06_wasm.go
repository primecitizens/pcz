// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

type (
	_ unsafe.Pointer
	_ js.Ref
)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_AudioParam_Value
//go:noescape
func SetAudioParamValue(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_AudioParam_AutomationRate
//go:noescape
func GetAudioParamAutomationRate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_AudioParam_AutomationRate
//go:noescape
func SetAudioParamAutomationRate(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_AudioParam_DefaultValue
//go:noescape
func GetAudioParamDefaultValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioParam_MinValue
//go:noescape
func GetAudioParamMinValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioParam_MaxValue
//go:noescape
func GetAudioParamMaxValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioParam_SetValueAtTime
//go:noescape
func HasFuncAudioParamSetValueAtTime(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioParam_SetValueAtTime
//go:noescape
func FuncAudioParamSetValueAtTime(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioParam_SetValueAtTime
//go:noescape
func CallAudioParamSetValueAtTime(
	this js.Ref, retPtr unsafe.Pointer,
	value float32,
	startTime float64)

//go:wasmimport plat/js/web try_AudioParam_SetValueAtTime
//go:noescape
func TryAudioParamSetValueAtTime(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float32,
	startTime float64) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioParam_LinearRampToValueAtTime
//go:noescape
func HasFuncAudioParamLinearRampToValueAtTime(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioParam_LinearRampToValueAtTime
//go:noescape
func FuncAudioParamLinearRampToValueAtTime(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioParam_LinearRampToValueAtTime
//go:noescape
func CallAudioParamLinearRampToValueAtTime(
	this js.Ref, retPtr unsafe.Pointer,
	value float32,
	endTime float64)

//go:wasmimport plat/js/web try_AudioParam_LinearRampToValueAtTime
//go:noescape
func TryAudioParamLinearRampToValueAtTime(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float32,
	endTime float64) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioParam_ExponentialRampToValueAtTime
//go:noescape
func HasFuncAudioParamExponentialRampToValueAtTime(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioParam_ExponentialRampToValueAtTime
//go:noescape
func FuncAudioParamExponentialRampToValueAtTime(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioParam_ExponentialRampToValueAtTime
//go:noescape
func CallAudioParamExponentialRampToValueAtTime(
	this js.Ref, retPtr unsafe.Pointer,
	value float32,
	endTime float64)

//go:wasmimport plat/js/web try_AudioParam_ExponentialRampToValueAtTime
//go:noescape
func TryAudioParamExponentialRampToValueAtTime(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float32,
	endTime float64) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioParam_SetTargetAtTime
//go:noescape
func HasFuncAudioParamSetTargetAtTime(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioParam_SetTargetAtTime
//go:noescape
func FuncAudioParamSetTargetAtTime(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioParam_SetTargetAtTime
//go:noescape
func CallAudioParamSetTargetAtTime(
	this js.Ref, retPtr unsafe.Pointer,
	target float32,
	startTime float64,
	timeConstant float32)

//go:wasmimport plat/js/web try_AudioParam_SetTargetAtTime
//go:noescape
func TryAudioParamSetTargetAtTime(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	target float32,
	startTime float64,
	timeConstant float32) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioParam_SetValueCurveAtTime
//go:noescape
func HasFuncAudioParamSetValueCurveAtTime(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioParam_SetValueCurveAtTime
//go:noescape
func FuncAudioParamSetValueCurveAtTime(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioParam_SetValueCurveAtTime
//go:noescape
func CallAudioParamSetValueCurveAtTime(
	this js.Ref, retPtr unsafe.Pointer,
	values js.Ref,
	startTime float64,
	duration float64)

//go:wasmimport plat/js/web try_AudioParam_SetValueCurveAtTime
//go:noescape
func TryAudioParamSetValueCurveAtTime(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	values js.Ref,
	startTime float64,
	duration float64) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioParam_CancelScheduledValues
//go:noescape
func HasFuncAudioParamCancelScheduledValues(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioParam_CancelScheduledValues
//go:noescape
func FuncAudioParamCancelScheduledValues(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioParam_CancelScheduledValues
//go:noescape
func CallAudioParamCancelScheduledValues(
	this js.Ref, retPtr unsafe.Pointer,
	cancelTime float64)

//go:wasmimport plat/js/web try_AudioParam_CancelScheduledValues
//go:noescape
func TryAudioParamCancelScheduledValues(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cancelTime float64) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioParam_CancelAndHoldAtTime
//go:noescape
func HasFuncAudioParamCancelAndHoldAtTime(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioParam_CancelAndHoldAtTime
//go:noescape
func FuncAudioParamCancelAndHoldAtTime(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioParam_CancelAndHoldAtTime
//go:noescape
func CallAudioParamCancelAndHoldAtTime(
	this js.Ref, retPtr unsafe.Pointer,
	cancelTime float64)

//go:wasmimport plat/js/web try_AudioParam_CancelAndHoldAtTime
//go:noescape
func TryAudioParamCancelAndHoldAtTime(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cancelTime float64) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_BiquadFilterNode_Type
//go:noescape
func SetBiquadFilterNodeType(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_BiquadFilterNode_Frequency
//go:noescape
func GetBiquadFilterNodeFrequency(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BiquadFilterNode_Detune
//go:noescape
func GetBiquadFilterNodeDetune(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BiquadFilterNode_Q
//go:noescape
func GetBiquadFilterNodeQ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BiquadFilterNode_Gain
//go:noescape
func GetBiquadFilterNodeGain(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BiquadFilterNode_GetFrequencyResponse
//go:noescape
func HasFuncBiquadFilterNodeGetFrequencyResponse(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BiquadFilterNode_GetFrequencyResponse
//go:noescape
func FuncBiquadFilterNodeGetFrequencyResponse(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BiquadFilterNode_GetFrequencyResponse
//go:noescape
func CallBiquadFilterNodeGetFrequencyResponse(
	this js.Ref, retPtr unsafe.Pointer,
	frequencyHz js.Ref,
	magResponse js.Ref,
	phaseResponse js.Ref)

//go:wasmimport plat/js/web try_BiquadFilterNode_GetFrequencyResponse
//go:noescape
func TryBiquadFilterNodeGetFrequencyResponse(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	frequencyHz js.Ref,
	magResponse js.Ref,
	phaseResponse js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioBuffer_Length
//go:noescape
func GetAudioBufferLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioBuffer_Duration
//go:noescape
func GetAudioBufferDuration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioBuffer_NumberOfChannels
//go:noescape
func GetAudioBufferNumberOfChannels(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioBuffer_GetChannelData
//go:noescape
func HasFuncAudioBufferGetChannelData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioBuffer_GetChannelData
//go:noescape
func FuncAudioBufferGetChannelData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioBuffer_GetChannelData
//go:noescape
func CallAudioBufferGetChannelData(
	this js.Ref, retPtr unsafe.Pointer,
	channel uint32)

//go:wasmimport plat/js/web try_AudioBuffer_GetChannelData
//go:noescape
func TryAudioBufferGetChannelData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	channel uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioBuffer_CopyFromChannel
//go:noescape
func HasFuncAudioBufferCopyFromChannel(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioBuffer_CopyFromChannel
//go:noescape
func FuncAudioBufferCopyFromChannel(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioBuffer_CopyFromChannel
//go:noescape
func CallAudioBufferCopyFromChannel(
	this js.Ref, retPtr unsafe.Pointer,
	destination js.Ref,
	channelNumber uint32,
	bufferOffset uint32)

//go:wasmimport plat/js/web try_AudioBuffer_CopyFromChannel
//go:noescape
func TryAudioBufferCopyFromChannel(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destination js.Ref,
	channelNumber uint32,
	bufferOffset uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioBuffer_CopyFromChannel1
//go:noescape
func HasFuncAudioBufferCopyFromChannel1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioBuffer_CopyFromChannel1
//go:noescape
func FuncAudioBufferCopyFromChannel1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioBuffer_CopyFromChannel1
//go:noescape
func CallAudioBufferCopyFromChannel1(
	this js.Ref, retPtr unsafe.Pointer,
	destination js.Ref,
	channelNumber uint32)

//go:wasmimport plat/js/web try_AudioBuffer_CopyFromChannel1
//go:noescape
func TryAudioBufferCopyFromChannel1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destination js.Ref,
	channelNumber uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioBuffer_CopyToChannel
//go:noescape
func HasFuncAudioBufferCopyToChannel(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioBuffer_CopyToChannel
//go:noescape
func FuncAudioBufferCopyToChannel(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioBuffer_CopyToChannel
//go:noescape
func CallAudioBufferCopyToChannel(
	this js.Ref, retPtr unsafe.Pointer,
	source js.Ref,
	channelNumber uint32,
	bufferOffset uint32)

//go:wasmimport plat/js/web try_AudioBuffer_CopyToChannel
//go:noescape
func TryAudioBufferCopyToChannel(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source js.Ref,
	channelNumber uint32,
	bufferOffset uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioBuffer_CopyToChannel1
//go:noescape
func HasFuncAudioBufferCopyToChannel1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioBuffer_CopyToChannel1
//go:noescape
func FuncAudioBufferCopyToChannel1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioBuffer_CopyToChannel1
//go:noescape
func CallAudioBufferCopyToChannel1(
	this js.Ref, retPtr unsafe.Pointer,
	source js.Ref,
	channelNumber uint32)

//go:wasmimport plat/js/web try_AudioBuffer_CopyToChannel1
//go:noescape
func TryAudioBufferCopyToChannel1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source js.Ref,
	channelNumber uint32) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_AudioBufferSourceNode_Buffer
//go:noescape
func SetAudioBufferSourceNodeBuffer(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_AudioBufferSourceNode_PlaybackRate
//go:noescape
func GetAudioBufferSourceNodePlaybackRate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioBufferSourceNode_Detune
//go:noescape
func GetAudioBufferSourceNodeDetune(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioBufferSourceNode_Loop
//go:noescape
func GetAudioBufferSourceNodeLoop(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_AudioBufferSourceNode_Loop
//go:noescape
func SetAudioBufferSourceNodeLoop(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_AudioBufferSourceNode_LoopStart
//go:noescape
func GetAudioBufferSourceNodeLoopStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_AudioBufferSourceNode_LoopStart
//go:noescape
func SetAudioBufferSourceNodeLoopStart(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_AudioBufferSourceNode_LoopEnd
//go:noescape
func GetAudioBufferSourceNodeLoopEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_AudioBufferSourceNode_LoopEnd
//go:noescape
func SetAudioBufferSourceNodeLoopEnd(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web has_AudioBufferSourceNode_Start
//go:noescape
func HasFuncAudioBufferSourceNodeStart(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioBufferSourceNode_Start
//go:noescape
func FuncAudioBufferSourceNodeStart(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioBufferSourceNode_Start
//go:noescape
func CallAudioBufferSourceNodeStart(
	this js.Ref, retPtr unsafe.Pointer,
	when float64,
	offset float64,
	duration float64)

//go:wasmimport plat/js/web try_AudioBufferSourceNode_Start
//go:noescape
func TryAudioBufferSourceNodeStart(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	when float64,
	offset float64,
	duration float64) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioBufferSourceNode_Start1
//go:noescape
func HasFuncAudioBufferSourceNodeStart1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioBufferSourceNode_Start1
//go:noescape
func FuncAudioBufferSourceNodeStart1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioBufferSourceNode_Start1
//go:noescape
func CallAudioBufferSourceNodeStart1(
	this js.Ref, retPtr unsafe.Pointer,
	when float64,
	offset float64)

//go:wasmimport plat/js/web try_AudioBufferSourceNode_Start1
//go:noescape
func TryAudioBufferSourceNodeStart1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	when float64,
	offset float64) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioBufferSourceNode_Start2
//go:noescape
func HasFuncAudioBufferSourceNodeStart2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioBufferSourceNode_Start2
//go:noescape
func FuncAudioBufferSourceNodeStart2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioBufferSourceNode_Start2
//go:noescape
func CallAudioBufferSourceNodeStart2(
	this js.Ref, retPtr unsafe.Pointer,
	when float64)

//go:wasmimport plat/js/web try_AudioBufferSourceNode_Start2
//go:noescape
func TryAudioBufferSourceNodeStart2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	when float64) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioBufferSourceNode_Start3
//go:noescape
func HasFuncAudioBufferSourceNodeStart3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioBufferSourceNode_Start3
//go:noescape
func FuncAudioBufferSourceNodeStart3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioBufferSourceNode_Start3
//go:noescape
func CallAudioBufferSourceNodeStart3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioBufferSourceNode_Start3
//go:noescape
func TryAudioBufferSourceNodeStart3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ConvolverNode_Buffer
//go:noescape
func SetConvolverNodeBuffer(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ConvolverNode_Normalize
//go:noescape
func GetConvolverNodeNormalize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DynamicsCompressorNode_Knee
//go:noescape
func GetDynamicsCompressorNodeKnee(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DynamicsCompressorNode_Ratio
//go:noescape
func GetDynamicsCompressorNodeRatio(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DynamicsCompressorNode_Reduction
//go:noescape
func GetDynamicsCompressorNodeReduction(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DynamicsCompressorNode_Attack
//go:noescape
func GetDynamicsCompressorNodeAttack(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DynamicsCompressorNode_Release
//go:noescape
func GetDynamicsCompressorNodeRelease(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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

//go:wasmimport plat/js/web has_IIRFilterNode_GetFrequencyResponse
//go:noescape
func HasFuncIIRFilterNodeGetFrequencyResponse(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IIRFilterNode_GetFrequencyResponse
//go:noescape
func FuncIIRFilterNodeGetFrequencyResponse(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IIRFilterNode_GetFrequencyResponse
//go:noescape
func CallIIRFilterNodeGetFrequencyResponse(
	this js.Ref, retPtr unsafe.Pointer,
	frequencyHz js.Ref,
	magResponse js.Ref,
	phaseResponse js.Ref)

//go:wasmimport plat/js/web try_IIRFilterNode_GetFrequencyResponse
//go:noescape
func TryIIRFilterNodeGetFrequencyResponse(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	frequencyHz js.Ref,
	magResponse js.Ref,
	phaseResponse js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OscillatorNode_Type
//go:noescape
func SetOscillatorNodeType(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_OscillatorNode_Frequency
//go:noescape
func GetOscillatorNodeFrequency(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_OscillatorNode_Detune
//go:noescape
func GetOscillatorNodeDetune(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OscillatorNode_SetPeriodicWave
//go:noescape
func HasFuncOscillatorNodeSetPeriodicWave(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OscillatorNode_SetPeriodicWave
//go:noescape
func FuncOscillatorNodeSetPeriodicWave(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OscillatorNode_SetPeriodicWave
//go:noescape
func CallOscillatorNodeSetPeriodicWave(
	this js.Ref, retPtr unsafe.Pointer,
	periodicWave js.Ref)

//go:wasmimport plat/js/web try_OscillatorNode_SetPeriodicWave
//go:noescape
func TryOscillatorNodeSetPeriodicWave(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	periodicWave js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_PannerNode_PanningModel
//go:noescape
func SetPannerNodePanningModel(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_PositionX
//go:noescape
func GetPannerNodePositionX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PannerNode_PositionY
//go:noescape
func GetPannerNodePositionY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PannerNode_PositionZ
//go:noescape
func GetPannerNodePositionZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PannerNode_OrientationX
//go:noescape
func GetPannerNodeOrientationX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PannerNode_OrientationY
//go:noescape
func GetPannerNodeOrientationY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PannerNode_OrientationZ
//go:noescape
func GetPannerNodeOrientationZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PannerNode_DistanceModel
//go:noescape
func GetPannerNodeDistanceModel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_PannerNode_DistanceModel
//go:noescape
func SetPannerNodeDistanceModel(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_RefDistance
//go:noescape
func GetPannerNodeRefDistance(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_PannerNode_RefDistance
//go:noescape
func SetPannerNodeRefDistance(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_MaxDistance
//go:noescape
func GetPannerNodeMaxDistance(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_PannerNode_MaxDistance
//go:noescape
func SetPannerNodeMaxDistance(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_RolloffFactor
//go:noescape
func GetPannerNodeRolloffFactor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_PannerNode_RolloffFactor
//go:noescape
func SetPannerNodeRolloffFactor(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_ConeInnerAngle
//go:noescape
func GetPannerNodeConeInnerAngle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_PannerNode_ConeInnerAngle
//go:noescape
func SetPannerNodeConeInnerAngle(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_ConeOuterAngle
//go:noescape
func GetPannerNodeConeOuterAngle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_PannerNode_ConeOuterAngle
//go:noescape
func SetPannerNodeConeOuterAngle(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PannerNode_ConeOuterGain
//go:noescape
func GetPannerNodeConeOuterGain(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_PannerNode_ConeOuterGain
//go:noescape
func SetPannerNodeConeOuterGain(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web has_PannerNode_SetPosition
//go:noescape
func HasFuncPannerNodeSetPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PannerNode_SetPosition
//go:noescape
func FuncPannerNodeSetPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_PannerNode_SetPosition
//go:noescape
func CallPannerNodeSetPosition(
	this js.Ref, retPtr unsafe.Pointer,
	x float32,
	y float32,
	z float32)

//go:wasmimport plat/js/web try_PannerNode_SetPosition
//go:noescape
func TryPannerNodeSetPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float32,
	y float32,
	z float32) (ok js.Ref)

//go:wasmimport plat/js/web has_PannerNode_SetOrientation
//go:noescape
func HasFuncPannerNodeSetOrientation(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PannerNode_SetOrientation
//go:noescape
func FuncPannerNodeSetOrientation(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_PannerNode_SetOrientation
//go:noescape
func CallPannerNodeSetOrientation(
	this js.Ref, retPtr unsafe.Pointer,
	x float32,
	y float32,
	z float32)

//go:wasmimport plat/js/web try_PannerNode_SetOrientation
//go:noescape
func TryPannerNodeSetOrientation(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float32,
	y float32,
	z float32) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_OverSampleType
//go:noescape
func ConstOfOverSampleType(str js.Ref) uint32

//go:wasmimport plat/js/web store_WaveShaperOptions
//go:noescape
func WaveShaperOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WaveShaperOptions
//go:noescape
func WaveShaperOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_WaveShaperNode_WaveShaperNode
//go:noescape
func NewWaveShaperNodeByWaveShaperNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_WaveShaperNode_WaveShaperNode1
//go:noescape
func NewWaveShaperNodeByWaveShaperNode1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_WaveShaperNode_Curve
//go:noescape
func GetWaveShaperNodeCurve(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_WaveShaperNode_Curve
//go:noescape
func SetWaveShaperNodeCurve(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_WaveShaperNode_Oversample
//go:noescape
func GetWaveShaperNodeOversample(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_WaveShaperNode_Oversample
//go:noescape
func SetWaveShaperNodeOversample(
	this js.Ref,
	val uint32,
) js.Ref
