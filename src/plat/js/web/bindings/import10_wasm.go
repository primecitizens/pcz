// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

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

//go:wasmimport plat/js/web store_AudioEncoderInit
//go:noescape
func AudioEncoderInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioEncoderInit
//go:noescape
func AudioEncoderInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_BitrateMode
//go:noescape
func ConstOfBitrateMode(str js.Ref) uint32

//go:wasmimport plat/js/web store_FlacEncoderConfig
//go:noescape
func FlacEncoderConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FlacEncoderConfig
//go:noescape
func FlacEncoderConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_OpusBitstreamFormat
//go:noescape
func ConstOfOpusBitstreamFormat(str js.Ref) uint32

//go:wasmimport plat/js/web store_OpusEncoderConfig
//go:noescape
func OpusEncoderConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_OpusEncoderConfig
//go:noescape
func OpusEncoderConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AudioEncoderConfig
//go:noescape
func AudioEncoderConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioEncoderConfig
//go:noescape
func AudioEncoderConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AudioEncoderSupport
//go:noescape
func AudioEncoderSupportJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioEncoderSupport
//go:noescape
func AudioEncoderSupportJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_AudioEncoder_AudioEncoder
//go:noescape
func NewAudioEncoderByAudioEncoder(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_AudioEncoder_State
//go:noescape
func GetAudioEncoderState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioEncoder_EncodeQueueSize
//go:noescape
func GetAudioEncoderEncodeQueueSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioEncoder_Configure
//go:noescape
func HasAudioEncoderConfigure(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioEncoder_Configure
//go:noescape
func AudioEncoderConfigureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioEncoder_Configure
//go:noescape
func CallAudioEncoderConfigure(
	this js.Ref, retPtr unsafe.Pointer,
	config unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioEncoder_Configure
//go:noescape
func TryAudioEncoderConfigure(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	config unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioEncoder_Encode
//go:noescape
func HasAudioEncoderEncode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioEncoder_Encode
//go:noescape
func AudioEncoderEncodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioEncoder_Encode
//go:noescape
func CallAudioEncoderEncode(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_AudioEncoder_Encode
//go:noescape
func TryAudioEncoderEncode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioEncoder_Flush
//go:noescape
func HasAudioEncoderFlush(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioEncoder_Flush
//go:noescape
func AudioEncoderFlushFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioEncoder_Flush
//go:noescape
func CallAudioEncoderFlush(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioEncoder_Flush
//go:noescape
func TryAudioEncoderFlush(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioEncoder_Reset
//go:noescape
func HasAudioEncoderReset(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioEncoder_Reset
//go:noescape
func AudioEncoderResetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioEncoder_Reset
//go:noescape
func CallAudioEncoderReset(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioEncoder_Reset
//go:noescape
func TryAudioEncoderReset(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioEncoder_Close
//go:noescape
func HasAudioEncoderClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioEncoder_Close
//go:noescape
func AudioEncoderCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioEncoder_Close
//go:noescape
func CallAudioEncoderClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioEncoder_Close
//go:noescape
func TryAudioEncoderClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioEncoder_IsConfigSupported
//go:noescape
func HasAudioEncoderIsConfigSupported(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioEncoder_IsConfigSupported
//go:noescape
func AudioEncoderIsConfigSupportedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioEncoder_IsConfigSupported
//go:noescape
func CallAudioEncoderIsConfigSupported(
	this js.Ref, retPtr unsafe.Pointer,
	config unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioEncoder_IsConfigSupported
//go:noescape
func TryAudioEncoderIsConfigSupported(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	config unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioNode_Context
//go:noescape
func GetAudioNodeContext(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioNode_NumberOfInputs
//go:noescape
func GetAudioNodeNumberOfInputs(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioNode_NumberOfOutputs
//go:noescape
func GetAudioNodeNumberOfOutputs(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioNode_ChannelCount
//go:noescape
func GetAudioNodeChannelCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_AudioNode_ChannelCount
//go:noescape
func SetAudioNodeChannelCount(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_AudioNode_ChannelCountMode
//go:noescape
func GetAudioNodeChannelCountMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_AudioNode_ChannelCountMode
//go:noescape
func SetAudioNodeChannelCountMode(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_AudioNode_ChannelInterpretation
//go:noescape
func GetAudioNodeChannelInterpretation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_AudioNode_ChannelInterpretation
//go:noescape
func SetAudioNodeChannelInterpretation(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web has_AudioNode_Connect
//go:noescape
func HasAudioNodeConnect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioNode_Connect
//go:noescape
func AudioNodeConnectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioNode_Connect
//go:noescape
func CallAudioNodeConnect(
	this js.Ref, retPtr unsafe.Pointer,
	destinationNode js.Ref,
	output uint32,
	input uint32)

//go:wasmimport plat/js/web try_AudioNode_Connect
//go:noescape
func TryAudioNodeConnect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destinationNode js.Ref,
	output uint32,
	input uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioNode_Connect1
//go:noescape
func HasAudioNodeConnect1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioNode_Connect1
//go:noescape
func AudioNodeConnect1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioNode_Connect1
//go:noescape
func CallAudioNodeConnect1(
	this js.Ref, retPtr unsafe.Pointer,
	destinationNode js.Ref,
	output uint32)

//go:wasmimport plat/js/web try_AudioNode_Connect1
//go:noescape
func TryAudioNodeConnect1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destinationNode js.Ref,
	output uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioNode_Connect2
//go:noescape
func HasAudioNodeConnect2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioNode_Connect2
//go:noescape
func AudioNodeConnect2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioNode_Connect2
//go:noescape
func CallAudioNodeConnect2(
	this js.Ref, retPtr unsafe.Pointer,
	destinationNode js.Ref)

//go:wasmimport plat/js/web try_AudioNode_Connect2
//go:noescape
func TryAudioNodeConnect2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destinationNode js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioNode_Connect3
//go:noescape
func HasAudioNodeConnect3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioNode_Connect3
//go:noescape
func AudioNodeConnect3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioNode_Connect3
//go:noescape
func CallAudioNodeConnect3(
	this js.Ref, retPtr unsafe.Pointer,
	destinationParam js.Ref,
	output uint32)

//go:wasmimport plat/js/web try_AudioNode_Connect3
//go:noescape
func TryAudioNodeConnect3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destinationParam js.Ref,
	output uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioNode_Connect4
//go:noescape
func HasAudioNodeConnect4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioNode_Connect4
//go:noescape
func AudioNodeConnect4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioNode_Connect4
//go:noescape
func CallAudioNodeConnect4(
	this js.Ref, retPtr unsafe.Pointer,
	destinationParam js.Ref)

//go:wasmimport plat/js/web try_AudioNode_Connect4
//go:noescape
func TryAudioNodeConnect4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destinationParam js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioNode_Disconnect
//go:noescape
func HasAudioNodeDisconnect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioNode_Disconnect
//go:noescape
func AudioNodeDisconnectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioNode_Disconnect
//go:noescape
func CallAudioNodeDisconnect(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioNode_Disconnect
//go:noescape
func TryAudioNodeDisconnect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioNode_Disconnect1
//go:noescape
func HasAudioNodeDisconnect1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioNode_Disconnect1
//go:noescape
func AudioNodeDisconnect1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioNode_Disconnect1
//go:noescape
func CallAudioNodeDisconnect1(
	this js.Ref, retPtr unsafe.Pointer,
	output uint32)

//go:wasmimport plat/js/web try_AudioNode_Disconnect1
//go:noescape
func TryAudioNodeDisconnect1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	output uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioNode_Disconnect2
//go:noescape
func HasAudioNodeDisconnect2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioNode_Disconnect2
//go:noescape
func AudioNodeDisconnect2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioNode_Disconnect2
//go:noescape
func CallAudioNodeDisconnect2(
	this js.Ref, retPtr unsafe.Pointer,
	destinationNode js.Ref)

//go:wasmimport plat/js/web try_AudioNode_Disconnect2
//go:noescape
func TryAudioNodeDisconnect2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destinationNode js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioNode_Disconnect3
//go:noescape
func HasAudioNodeDisconnect3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioNode_Disconnect3
//go:noescape
func AudioNodeDisconnect3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioNode_Disconnect3
//go:noescape
func CallAudioNodeDisconnect3(
	this js.Ref, retPtr unsafe.Pointer,
	destinationNode js.Ref,
	output uint32)

//go:wasmimport plat/js/web try_AudioNode_Disconnect3
//go:noescape
func TryAudioNodeDisconnect3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destinationNode js.Ref,
	output uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioNode_Disconnect4
//go:noescape
func HasAudioNodeDisconnect4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioNode_Disconnect4
//go:noescape
func AudioNodeDisconnect4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioNode_Disconnect4
//go:noescape
func CallAudioNodeDisconnect4(
	this js.Ref, retPtr unsafe.Pointer,
	destinationNode js.Ref,
	output uint32,
	input uint32)

//go:wasmimport plat/js/web try_AudioNode_Disconnect4
//go:noescape
func TryAudioNodeDisconnect4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destinationNode js.Ref,
	output uint32,
	input uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioNode_Disconnect5
//go:noescape
func HasAudioNodeDisconnect5(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioNode_Disconnect5
//go:noescape
func AudioNodeDisconnect5Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioNode_Disconnect5
//go:noescape
func CallAudioNodeDisconnect5(
	this js.Ref, retPtr unsafe.Pointer,
	destinationParam js.Ref)

//go:wasmimport plat/js/web try_AudioNode_Disconnect5
//go:noescape
func TryAudioNodeDisconnect5(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destinationParam js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioNode_Disconnect6
//go:noescape
func HasAudioNodeDisconnect6(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioNode_Disconnect6
//go:noescape
func AudioNodeDisconnect6Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioNode_Disconnect6
//go:noescape
func CallAudioNodeDisconnect6(
	this js.Ref, retPtr unsafe.Pointer,
	destinationParam js.Ref,
	output uint32)

//go:wasmimport plat/js/web try_AudioNode_Disconnect6
//go:noescape
func TryAudioNodeDisconnect6(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destinationParam js.Ref,
	output uint32) (ok js.Ref)

//go:wasmimport plat/js/web store_AudioOutputOptions
//go:noescape
func AudioOutputOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioOutputOptions
//go:noescape
func AudioOutputOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AudioParamDescriptor
//go:noescape
func AudioParamDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioParamDescriptor
//go:noescape
func AudioParamDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AudioProcessingEventInit
//go:noescape
func AudioProcessingEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioProcessingEventInit
//go:noescape
func AudioProcessingEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_AudioProcessingEvent_AudioProcessingEvent
//go:noescape
func NewAudioProcessingEventByAudioProcessingEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_AudioProcessingEvent_PlaybackTime
//go:noescape
func GetAudioProcessingEventPlaybackTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioProcessingEvent_InputBuffer
//go:noescape
func GetAudioProcessingEventInputBuffer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioProcessingEvent_OutputBuffer
//go:noescape
func GetAudioProcessingEventOutputBuffer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_AudioRenderCapacityEventInit
//go:noescape
func AudioRenderCapacityEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioRenderCapacityEventInit
//go:noescape
func AudioRenderCapacityEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_AudioRenderCapacityEvent_AudioRenderCapacityEvent
//go:noescape
func NewAudioRenderCapacityEventByAudioRenderCapacityEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_AudioRenderCapacityEvent_AudioRenderCapacityEvent1
//go:noescape
func NewAudioRenderCapacityEventByAudioRenderCapacityEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_AudioRenderCapacityEvent_Timestamp
//go:noescape
func GetAudioRenderCapacityEventTimestamp(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioRenderCapacityEvent_AverageLoad
//go:noescape
func GetAudioRenderCapacityEventAverageLoad(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioRenderCapacityEvent_PeakLoad
//go:noescape
func GetAudioRenderCapacityEventPeakLoad(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioRenderCapacityEvent_UnderrunRatio
//go:noescape
func GetAudioRenderCapacityEventUnderrunRatio(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioScheduledSourceNode_Start
//go:noescape
func HasAudioScheduledSourceNodeStart(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioScheduledSourceNode_Start
//go:noescape
func AudioScheduledSourceNodeStartFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioScheduledSourceNode_Start
//go:noescape
func CallAudioScheduledSourceNodeStart(
	this js.Ref, retPtr unsafe.Pointer,
	when float64)

//go:wasmimport plat/js/web try_AudioScheduledSourceNode_Start
//go:noescape
func TryAudioScheduledSourceNodeStart(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	when float64) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioScheduledSourceNode_Start1
//go:noescape
func HasAudioScheduledSourceNodeStart1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioScheduledSourceNode_Start1
//go:noescape
func AudioScheduledSourceNodeStart1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioScheduledSourceNode_Start1
//go:noescape
func CallAudioScheduledSourceNodeStart1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioScheduledSourceNode_Start1
//go:noescape
func TryAudioScheduledSourceNodeStart1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioScheduledSourceNode_Stop
//go:noescape
func HasAudioScheduledSourceNodeStop(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioScheduledSourceNode_Stop
//go:noescape
func AudioScheduledSourceNodeStopFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioScheduledSourceNode_Stop
//go:noescape
func CallAudioScheduledSourceNodeStop(
	this js.Ref, retPtr unsafe.Pointer,
	when float64)

//go:wasmimport plat/js/web try_AudioScheduledSourceNode_Stop
//go:noescape
func TryAudioScheduledSourceNodeStop(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	when float64) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioScheduledSourceNode_Stop1
//go:noescape
func HasAudioScheduledSourceNodeStop1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioScheduledSourceNode_Stop1
//go:noescape
func AudioScheduledSourceNodeStop1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioScheduledSourceNode_Stop1
//go:noescape
func CallAudioScheduledSourceNodeStop1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioScheduledSourceNode_Stop1
//go:noescape
func TryAudioScheduledSourceNodeStop1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioWorkletProcessor_Port
//go:noescape
func GetAudioWorkletProcessorPort(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioWorkletGlobalScope_CurrentFrame
//go:noescape
func GetAudioWorkletGlobalScopeCurrentFrame(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioWorkletGlobalScope_CurrentTime
//go:noescape
func GetAudioWorkletGlobalScopeCurrentTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioWorkletGlobalScope_SampleRate
//go:noescape
func GetAudioWorkletGlobalScopeSampleRate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioWorkletGlobalScope_Port
//go:noescape
func GetAudioWorkletGlobalScopePort(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioWorkletGlobalScope_RegisterProcessor
//go:noescape
func HasAudioWorkletGlobalScopeRegisterProcessor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioWorkletGlobalScope_RegisterProcessor
//go:noescape
func AudioWorkletGlobalScopeRegisterProcessorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioWorkletGlobalScope_RegisterProcessor
//go:noescape
func CallAudioWorkletGlobalScopeRegisterProcessor(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	processorCtor js.Ref)

//go:wasmimport plat/js/web try_AudioWorkletGlobalScope_RegisterProcessor
//go:noescape
func TryAudioWorkletGlobalScopeRegisterProcessor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	processorCtor js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_AudioWorkletNodeOptions
//go:noescape
func AudioWorkletNodeOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioWorkletNodeOptions
//go:noescape
func AudioWorkletNodeOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_AudioWorkletNode_AudioWorkletNode
//go:noescape
func NewAudioWorkletNodeByAudioWorkletNode(
	context js.Ref,
	name js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_AudioWorkletNode_AudioWorkletNode1
//go:noescape
func NewAudioWorkletNodeByAudioWorkletNode1(
	context js.Ref,
	name js.Ref) js.Ref

//go:wasmimport plat/js/web get_AudioWorkletNode_Parameters
//go:noescape
func GetAudioWorkletNodeParameters(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioWorkletNode_Port
//go:noescape
func GetAudioWorkletNodePort(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_AuthenticationExtensionsPRFValues
//go:noescape
func AuthenticationExtensionsPRFValuesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticationExtensionsPRFValues
//go:noescape
func AuthenticationExtensionsPRFValuesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuthenticationExtensionsPRFInputs
//go:noescape
func AuthenticationExtensionsPRFInputsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticationExtensionsPRFInputs
//go:noescape
func AuthenticationExtensionsPRFInputsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PaymentCurrencyAmount
//go:noescape
func PaymentCurrencyAmountJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentCurrencyAmount
//go:noescape
func PaymentCurrencyAmountJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PaymentCredentialInstrument
//go:noescape
func PaymentCredentialInstrumentJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentCredentialInstrument
//go:noescape
func PaymentCredentialInstrumentJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuthenticationExtensionsPaymentInputs
//go:noescape
func AuthenticationExtensionsPaymentInputsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticationExtensionsPaymentInputs
//go:noescape
func AuthenticationExtensionsPaymentInputsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_HMACGetSecretInput
//go:noescape
func HMACGetSecretInputJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_HMACGetSecretInput
//go:noescape
func HMACGetSecretInputJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuthenticationExtensionsLargeBlobInputs
//go:noescape
func AuthenticationExtensionsLargeBlobInputsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticationExtensionsLargeBlobInputs
//go:noescape
func AuthenticationExtensionsLargeBlobInputsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuthenticationExtensionsDevicePublicKeyInputs
//go:noescape
func AuthenticationExtensionsDevicePublicKeyInputsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticationExtensionsDevicePublicKeyInputs
//go:noescape
func AuthenticationExtensionsDevicePublicKeyInputsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuthenticationExtensionsClientInputs
//go:noescape
func AuthenticationExtensionsClientInputsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticationExtensionsClientInputs
//go:noescape
func AuthenticationExtensionsClientInputsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuthenticationExtensionsClientInputsJSON
//go:noescape
func AuthenticationExtensionsClientInputsJSONJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticationExtensionsClientInputsJSON
//go:noescape
func AuthenticationExtensionsClientInputsJSONJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuthenticationExtensionsDevicePublicKeyOutputs
//go:noescape
func AuthenticationExtensionsDevicePublicKeyOutputsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticationExtensionsDevicePublicKeyOutputs
//go:noescape
func AuthenticationExtensionsDevicePublicKeyOutputsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_CredentialPropertiesOutput
//go:noescape
func CredentialPropertiesOutputJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CredentialPropertiesOutput
//go:noescape
func CredentialPropertiesOutputJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuthenticationExtensionsLargeBlobOutputs
//go:noescape
func AuthenticationExtensionsLargeBlobOutputsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticationExtensionsLargeBlobOutputs
//go:noescape
func AuthenticationExtensionsLargeBlobOutputsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_HMACGetSecretOutput
//go:noescape
func HMACGetSecretOutputJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_HMACGetSecretOutput
//go:noescape
func HMACGetSecretOutputJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuthenticationExtensionsPRFOutputs
//go:noescape
func AuthenticationExtensionsPRFOutputsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticationExtensionsPRFOutputs
//go:noescape
func AuthenticationExtensionsPRFOutputsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuthenticationExtensionsClientOutputs
//go:noescape
func AuthenticationExtensionsClientOutputsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticationExtensionsClientOutputs
//go:noescape
func AuthenticationExtensionsClientOutputsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuthenticationExtensionsClientOutputsJSON
//go:noescape
func AuthenticationExtensionsClientOutputsJSONJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticationExtensionsClientOutputsJSON
//go:noescape
func AuthenticationExtensionsClientOutputsJSONJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuthenticatorAssertionResponseJSON
//go:noescape
func AuthenticatorAssertionResponseJSONJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticatorAssertionResponseJSON
//go:noescape
func AuthenticatorAssertionResponseJSONJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuthenticationResponseJSON
//go:noescape
func AuthenticationResponseJSONJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticationResponseJSON
//go:noescape
func AuthenticationResponseJSONJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_AuthenticatorAssertionResponse_AuthenticatorData
//go:noescape
func GetAuthenticatorAssertionResponseAuthenticatorData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AuthenticatorAssertionResponse_Signature
//go:noescape
func GetAuthenticatorAssertionResponseSignature(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AuthenticatorAssertionResponse_UserHandle
//go:noescape
func GetAuthenticatorAssertionResponseUserHandle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AuthenticatorAssertionResponse_AttestationObject
//go:noescape
func GetAuthenticatorAssertionResponseAttestationObject(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_AuthenticatorAttachment
//go:noescape
func ConstOfAuthenticatorAttachment(str js.Ref) uint32
