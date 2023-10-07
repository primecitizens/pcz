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

//go:wasmimport plat/js/web new_DOMException_DOMException
//go:noescape
func NewDOMExceptionByDOMException(
	message js.Ref,
	name js.Ref) js.Ref

//go:wasmimport plat/js/web new_DOMException_DOMException1
//go:noescape
func NewDOMExceptionByDOMException1(
	message js.Ref) js.Ref

//go:wasmimport plat/js/web new_DOMException_DOMException2
//go:noescape
func NewDOMExceptionByDOMException2() js.Ref

//go:wasmimport plat/js/web get_DOMException_Name
//go:noescape
func GetDOMExceptionName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMException_Message
//go:noescape
func GetDOMExceptionMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMException_Code
//go:noescape
func GetDOMExceptionCode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioDestinationNode_MaxChannelCount
//go:noescape
func GetAudioDestinationNodeMaxChannelCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioListener_PositionX
//go:noescape
func GetAudioListenerPositionX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioListener_PositionY
//go:noescape
func GetAudioListenerPositionY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioListener_PositionZ
//go:noescape
func GetAudioListenerPositionZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioListener_ForwardX
//go:noescape
func GetAudioListenerForwardX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioListener_ForwardY
//go:noescape
func GetAudioListenerForwardY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioListener_ForwardZ
//go:noescape
func GetAudioListenerForwardZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioListener_UpX
//go:noescape
func GetAudioListenerUpX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioListener_UpY
//go:noescape
func GetAudioListenerUpY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioListener_UpZ
//go:noescape
func GetAudioListenerUpZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioListener_SetPosition
//go:noescape
func HasFuncAudioListenerSetPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioListener_SetPosition
//go:noescape
func FuncAudioListenerSetPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioListener_SetPosition
//go:noescape
func CallAudioListenerSetPosition(
	this js.Ref, retPtr unsafe.Pointer,
	x float32,
	y float32,
	z float32)

//go:wasmimport plat/js/web try_AudioListener_SetPosition
//go:noescape
func TryAudioListenerSetPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float32,
	y float32,
	z float32) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioListener_SetOrientation
//go:noescape
func HasFuncAudioListenerSetOrientation(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioListener_SetOrientation
//go:noescape
func FuncAudioListenerSetOrientation(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioListener_SetOrientation
//go:noescape
func CallAudioListenerSetOrientation(
	this js.Ref, retPtr unsafe.Pointer,
	x float32,
	y float32,
	z float32,
	xUp float32,
	yUp float32,
	zUp float32)

//go:wasmimport plat/js/web try_AudioListener_SetOrientation
//go:noescape
func TryAudioListenerSetOrientation(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float32,
	y float32,
	z float32,
	xUp float32,
	yUp float32,
	zUp float32) (ok js.Ref)

//go:wasmimport plat/js/web constof_AudioContextState
//go:noescape
func ConstOfAudioContextState(str js.Ref) uint32

//go:wasmimport plat/js/web store_StructuredSerializeOptions
//go:noescape
func StructuredSerializeOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_StructuredSerializeOptions
//go:noescape
func StructuredSerializeOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_MessagePort_PostMessage
//go:noescape
func HasFuncMessagePortPostMessage(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MessagePort_PostMessage
//go:noescape
func FuncMessagePortPostMessage(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MessagePort_PostMessage
//go:noescape
func CallMessagePortPostMessage(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	transfer js.Ref)

//go:wasmimport plat/js/web try_MessagePort_PostMessage
//go:noescape
func TryMessagePortPostMessage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	transfer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MessagePort_PostMessage1
//go:noescape
func HasFuncMessagePortPostMessage1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MessagePort_PostMessage1
//go:noescape
func FuncMessagePortPostMessage1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MessagePort_PostMessage1
//go:noescape
func CallMessagePortPostMessage1(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MessagePort_PostMessage1
//go:noescape
func TryMessagePortPostMessage1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MessagePort_PostMessage2
//go:noescape
func HasFuncMessagePortPostMessage2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MessagePort_PostMessage2
//go:noescape
func FuncMessagePortPostMessage2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MessagePort_PostMessage2
//go:noescape
func CallMessagePortPostMessage2(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/web try_MessagePort_PostMessage2
//go:noescape
func TryMessagePortPostMessage2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MessagePort_Start
//go:noescape
func HasFuncMessagePortStart(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MessagePort_Start
//go:noescape
func FuncMessagePortStart(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MessagePort_Start
//go:noescape
func CallMessagePortStart(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MessagePort_Start
//go:noescape
func TryMessagePortStart(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MessagePort_Close
//go:noescape
func HasFuncMessagePortClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MessagePort_Close
//go:noescape
func FuncMessagePortClose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MessagePort_Close
//go:noescape
func CallMessagePortClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MessagePort_Close
//go:noescape
func TryMessagePortClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioWorklet_Port
//go:noescape
func GetAudioWorkletPort(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BaseAudioContext_Destination
//go:noescape
func GetBaseAudioContextDestination(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BaseAudioContext_SampleRate
//go:noescape
func GetBaseAudioContextSampleRate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BaseAudioContext_CurrentTime
//go:noescape
func GetBaseAudioContextCurrentTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BaseAudioContext_Listener
//go:noescape
func GetBaseAudioContextListener(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BaseAudioContext_State
//go:noescape
func GetBaseAudioContextState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BaseAudioContext_AudioWorklet
//go:noescape
func GetBaseAudioContextAudioWorklet(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateAnalyser
//go:noescape
func HasFuncBaseAudioContextCreateAnalyser(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateAnalyser
//go:noescape
func FuncBaseAudioContextCreateAnalyser(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateAnalyser
//go:noescape
func CallBaseAudioContextCreateAnalyser(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateAnalyser
//go:noescape
func TryBaseAudioContextCreateAnalyser(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateBiquadFilter
//go:noescape
func HasFuncBaseAudioContextCreateBiquadFilter(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateBiquadFilter
//go:noescape
func FuncBaseAudioContextCreateBiquadFilter(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateBiquadFilter
//go:noescape
func CallBaseAudioContextCreateBiquadFilter(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateBiquadFilter
//go:noescape
func TryBaseAudioContextCreateBiquadFilter(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateBuffer
//go:noescape
func HasFuncBaseAudioContextCreateBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateBuffer
//go:noescape
func FuncBaseAudioContextCreateBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateBuffer
//go:noescape
func CallBaseAudioContextCreateBuffer(
	this js.Ref, retPtr unsafe.Pointer,
	numberOfChannels uint32,
	length uint32,
	sampleRate float32)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateBuffer
//go:noescape
func TryBaseAudioContextCreateBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	numberOfChannels uint32,
	length uint32,
	sampleRate float32) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateBufferSource
//go:noescape
func HasFuncBaseAudioContextCreateBufferSource(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateBufferSource
//go:noescape
func FuncBaseAudioContextCreateBufferSource(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateBufferSource
//go:noescape
func CallBaseAudioContextCreateBufferSource(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateBufferSource
//go:noescape
func TryBaseAudioContextCreateBufferSource(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateChannelMerger
//go:noescape
func HasFuncBaseAudioContextCreateChannelMerger(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateChannelMerger
//go:noescape
func FuncBaseAudioContextCreateChannelMerger(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateChannelMerger
//go:noescape
func CallBaseAudioContextCreateChannelMerger(
	this js.Ref, retPtr unsafe.Pointer,
	numberOfInputs uint32)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateChannelMerger
//go:noescape
func TryBaseAudioContextCreateChannelMerger(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	numberOfInputs uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateChannelMerger1
//go:noescape
func HasFuncBaseAudioContextCreateChannelMerger1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateChannelMerger1
//go:noescape
func FuncBaseAudioContextCreateChannelMerger1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateChannelMerger1
//go:noescape
func CallBaseAudioContextCreateChannelMerger1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateChannelMerger1
//go:noescape
func TryBaseAudioContextCreateChannelMerger1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateChannelSplitter
//go:noescape
func HasFuncBaseAudioContextCreateChannelSplitter(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateChannelSplitter
//go:noescape
func FuncBaseAudioContextCreateChannelSplitter(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateChannelSplitter
//go:noescape
func CallBaseAudioContextCreateChannelSplitter(
	this js.Ref, retPtr unsafe.Pointer,
	numberOfOutputs uint32)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateChannelSplitter
//go:noescape
func TryBaseAudioContextCreateChannelSplitter(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	numberOfOutputs uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateChannelSplitter1
//go:noescape
func HasFuncBaseAudioContextCreateChannelSplitter1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateChannelSplitter1
//go:noescape
func FuncBaseAudioContextCreateChannelSplitter1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateChannelSplitter1
//go:noescape
func CallBaseAudioContextCreateChannelSplitter1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateChannelSplitter1
//go:noescape
func TryBaseAudioContextCreateChannelSplitter1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateConstantSource
//go:noescape
func HasFuncBaseAudioContextCreateConstantSource(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateConstantSource
//go:noescape
func FuncBaseAudioContextCreateConstantSource(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateConstantSource
//go:noescape
func CallBaseAudioContextCreateConstantSource(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateConstantSource
//go:noescape
func TryBaseAudioContextCreateConstantSource(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateConvolver
//go:noescape
func HasFuncBaseAudioContextCreateConvolver(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateConvolver
//go:noescape
func FuncBaseAudioContextCreateConvolver(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateConvolver
//go:noescape
func CallBaseAudioContextCreateConvolver(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateConvolver
//go:noescape
func TryBaseAudioContextCreateConvolver(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateDelay
//go:noescape
func HasFuncBaseAudioContextCreateDelay(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateDelay
//go:noescape
func FuncBaseAudioContextCreateDelay(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateDelay
//go:noescape
func CallBaseAudioContextCreateDelay(
	this js.Ref, retPtr unsafe.Pointer,
	maxDelayTime float64)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateDelay
//go:noescape
func TryBaseAudioContextCreateDelay(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	maxDelayTime float64) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateDelay1
//go:noescape
func HasFuncBaseAudioContextCreateDelay1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateDelay1
//go:noescape
func FuncBaseAudioContextCreateDelay1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateDelay1
//go:noescape
func CallBaseAudioContextCreateDelay1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateDelay1
//go:noescape
func TryBaseAudioContextCreateDelay1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateDynamicsCompressor
//go:noescape
func HasFuncBaseAudioContextCreateDynamicsCompressor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateDynamicsCompressor
//go:noescape
func FuncBaseAudioContextCreateDynamicsCompressor(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateDynamicsCompressor
//go:noescape
func CallBaseAudioContextCreateDynamicsCompressor(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateDynamicsCompressor
//go:noescape
func TryBaseAudioContextCreateDynamicsCompressor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateGain
//go:noescape
func HasFuncBaseAudioContextCreateGain(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateGain
//go:noescape
func FuncBaseAudioContextCreateGain(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateGain
//go:noescape
func CallBaseAudioContextCreateGain(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateGain
//go:noescape
func TryBaseAudioContextCreateGain(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateIIRFilter
//go:noescape
func HasFuncBaseAudioContextCreateIIRFilter(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateIIRFilter
//go:noescape
func FuncBaseAudioContextCreateIIRFilter(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateIIRFilter
//go:noescape
func CallBaseAudioContextCreateIIRFilter(
	this js.Ref, retPtr unsafe.Pointer,
	feedforward js.Ref,
	feedback js.Ref)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateIIRFilter
//go:noescape
func TryBaseAudioContextCreateIIRFilter(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	feedforward js.Ref,
	feedback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateOscillator
//go:noescape
func HasFuncBaseAudioContextCreateOscillator(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateOscillator
//go:noescape
func FuncBaseAudioContextCreateOscillator(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateOscillator
//go:noescape
func CallBaseAudioContextCreateOscillator(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateOscillator
//go:noescape
func TryBaseAudioContextCreateOscillator(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreatePanner
//go:noescape
func HasFuncBaseAudioContextCreatePanner(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreatePanner
//go:noescape
func FuncBaseAudioContextCreatePanner(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreatePanner
//go:noescape
func CallBaseAudioContextCreatePanner(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreatePanner
//go:noescape
func TryBaseAudioContextCreatePanner(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreatePeriodicWave
//go:noescape
func HasFuncBaseAudioContextCreatePeriodicWave(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreatePeriodicWave
//go:noescape
func FuncBaseAudioContextCreatePeriodicWave(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreatePeriodicWave
//go:noescape
func CallBaseAudioContextCreatePeriodicWave(
	this js.Ref, retPtr unsafe.Pointer,
	real js.Ref,
	imag js.Ref,
	constraints unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreatePeriodicWave
//go:noescape
func TryBaseAudioContextCreatePeriodicWave(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	real js.Ref,
	imag js.Ref,
	constraints unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreatePeriodicWave1
//go:noescape
func HasFuncBaseAudioContextCreatePeriodicWave1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreatePeriodicWave1
//go:noescape
func FuncBaseAudioContextCreatePeriodicWave1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreatePeriodicWave1
//go:noescape
func CallBaseAudioContextCreatePeriodicWave1(
	this js.Ref, retPtr unsafe.Pointer,
	real js.Ref,
	imag js.Ref)

//go:wasmimport plat/js/web try_BaseAudioContext_CreatePeriodicWave1
//go:noescape
func TryBaseAudioContextCreatePeriodicWave1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	real js.Ref,
	imag js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateScriptProcessor
//go:noescape
func HasFuncBaseAudioContextCreateScriptProcessor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateScriptProcessor
//go:noescape
func FuncBaseAudioContextCreateScriptProcessor(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateScriptProcessor
//go:noescape
func CallBaseAudioContextCreateScriptProcessor(
	this js.Ref, retPtr unsafe.Pointer,
	bufferSize uint32,
	numberOfInputChannels uint32,
	numberOfOutputChannels uint32)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateScriptProcessor
//go:noescape
func TryBaseAudioContextCreateScriptProcessor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bufferSize uint32,
	numberOfInputChannels uint32,
	numberOfOutputChannels uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateScriptProcessor1
//go:noescape
func HasFuncBaseAudioContextCreateScriptProcessor1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateScriptProcessor1
//go:noescape
func FuncBaseAudioContextCreateScriptProcessor1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateScriptProcessor1
//go:noescape
func CallBaseAudioContextCreateScriptProcessor1(
	this js.Ref, retPtr unsafe.Pointer,
	bufferSize uint32,
	numberOfInputChannels uint32)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateScriptProcessor1
//go:noescape
func TryBaseAudioContextCreateScriptProcessor1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bufferSize uint32,
	numberOfInputChannels uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateScriptProcessor2
//go:noescape
func HasFuncBaseAudioContextCreateScriptProcessor2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateScriptProcessor2
//go:noescape
func FuncBaseAudioContextCreateScriptProcessor2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateScriptProcessor2
//go:noescape
func CallBaseAudioContextCreateScriptProcessor2(
	this js.Ref, retPtr unsafe.Pointer,
	bufferSize uint32)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateScriptProcessor2
//go:noescape
func TryBaseAudioContextCreateScriptProcessor2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bufferSize uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateScriptProcessor3
//go:noescape
func HasFuncBaseAudioContextCreateScriptProcessor3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateScriptProcessor3
//go:noescape
func FuncBaseAudioContextCreateScriptProcessor3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateScriptProcessor3
//go:noescape
func CallBaseAudioContextCreateScriptProcessor3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateScriptProcessor3
//go:noescape
func TryBaseAudioContextCreateScriptProcessor3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateStereoPanner
//go:noescape
func HasFuncBaseAudioContextCreateStereoPanner(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateStereoPanner
//go:noescape
func FuncBaseAudioContextCreateStereoPanner(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateStereoPanner
//go:noescape
func CallBaseAudioContextCreateStereoPanner(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateStereoPanner
//go:noescape
func TryBaseAudioContextCreateStereoPanner(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_CreateWaveShaper
//go:noescape
func HasFuncBaseAudioContextCreateWaveShaper(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_CreateWaveShaper
//go:noescape
func FuncBaseAudioContextCreateWaveShaper(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_CreateWaveShaper
//go:noescape
func CallBaseAudioContextCreateWaveShaper(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BaseAudioContext_CreateWaveShaper
//go:noescape
func TryBaseAudioContextCreateWaveShaper(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_DecodeAudioData
//go:noescape
func HasFuncBaseAudioContextDecodeAudioData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_DecodeAudioData
//go:noescape
func FuncBaseAudioContextDecodeAudioData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_DecodeAudioData
//go:noescape
func CallBaseAudioContextDecodeAudioData(
	this js.Ref, retPtr unsafe.Pointer,
	audioData js.Ref,
	successCallback js.Ref,
	errorCallback js.Ref)

//go:wasmimport plat/js/web try_BaseAudioContext_DecodeAudioData
//go:noescape
func TryBaseAudioContextDecodeAudioData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	audioData js.Ref,
	successCallback js.Ref,
	errorCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_DecodeAudioData1
//go:noescape
func HasFuncBaseAudioContextDecodeAudioData1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_DecodeAudioData1
//go:noescape
func FuncBaseAudioContextDecodeAudioData1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_DecodeAudioData1
//go:noescape
func CallBaseAudioContextDecodeAudioData1(
	this js.Ref, retPtr unsafe.Pointer,
	audioData js.Ref,
	successCallback js.Ref)

//go:wasmimport plat/js/web try_BaseAudioContext_DecodeAudioData1
//go:noescape
func TryBaseAudioContextDecodeAudioData1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	audioData js.Ref,
	successCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BaseAudioContext_DecodeAudioData2
//go:noescape
func HasFuncBaseAudioContextDecodeAudioData2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BaseAudioContext_DecodeAudioData2
//go:noescape
func FuncBaseAudioContextDecodeAudioData2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BaseAudioContext_DecodeAudioData2
//go:noescape
func CallBaseAudioContextDecodeAudioData2(
	this js.Ref, retPtr unsafe.Pointer,
	audioData js.Ref)

//go:wasmimport plat/js/web try_BaseAudioContext_DecodeAudioData2
//go:noescape
func TryBaseAudioContextDecodeAudioData2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	audioData js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_AnalyserOptions
//go:noescape
func AnalyserOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AnalyserOptions
//go:noescape
func AnalyserOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_AnalyserNode_AnalyserNode
//go:noescape
func NewAnalyserNodeByAnalyserNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_AnalyserNode_AnalyserNode1
//go:noescape
func NewAnalyserNodeByAnalyserNode1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_AnalyserNode_FftSize
//go:noescape
func GetAnalyserNodeFftSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_AnalyserNode_FftSize
//go:noescape
func SetAnalyserNodeFftSize(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_AnalyserNode_FrequencyBinCount
//go:noescape
func GetAnalyserNodeFrequencyBinCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AnalyserNode_MinDecibels
//go:noescape
func GetAnalyserNodeMinDecibels(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_AnalyserNode_MinDecibels
//go:noescape
func SetAnalyserNodeMinDecibels(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_AnalyserNode_MaxDecibels
//go:noescape
func GetAnalyserNodeMaxDecibels(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_AnalyserNode_MaxDecibels
//go:noescape
func SetAnalyserNodeMaxDecibels(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_AnalyserNode_SmoothingTimeConstant
//go:noescape
func GetAnalyserNodeSmoothingTimeConstant(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_AnalyserNode_SmoothingTimeConstant
//go:noescape
func SetAnalyserNodeSmoothingTimeConstant(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web has_AnalyserNode_GetFloatFrequencyData
//go:noescape
func HasFuncAnalyserNodeGetFloatFrequencyData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AnalyserNode_GetFloatFrequencyData
//go:noescape
func FuncAnalyserNodeGetFloatFrequencyData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AnalyserNode_GetFloatFrequencyData
//go:noescape
func CallAnalyserNodeGetFloatFrequencyData(
	this js.Ref, retPtr unsafe.Pointer,
	array js.Ref)

//go:wasmimport plat/js/web try_AnalyserNode_GetFloatFrequencyData
//go:noescape
func TryAnalyserNodeGetFloatFrequencyData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	array js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AnalyserNode_GetByteFrequencyData
//go:noescape
func HasFuncAnalyserNodeGetByteFrequencyData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AnalyserNode_GetByteFrequencyData
//go:noescape
func FuncAnalyserNodeGetByteFrequencyData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AnalyserNode_GetByteFrequencyData
//go:noescape
func CallAnalyserNodeGetByteFrequencyData(
	this js.Ref, retPtr unsafe.Pointer,
	array js.Ref)

//go:wasmimport plat/js/web try_AnalyserNode_GetByteFrequencyData
//go:noescape
func TryAnalyserNodeGetByteFrequencyData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	array js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AnalyserNode_GetFloatTimeDomainData
//go:noescape
func HasFuncAnalyserNodeGetFloatTimeDomainData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AnalyserNode_GetFloatTimeDomainData
//go:noescape
func FuncAnalyserNodeGetFloatTimeDomainData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AnalyserNode_GetFloatTimeDomainData
//go:noescape
func CallAnalyserNodeGetFloatTimeDomainData(
	this js.Ref, retPtr unsafe.Pointer,
	array js.Ref)

//go:wasmimport plat/js/web try_AnalyserNode_GetFloatTimeDomainData
//go:noescape
func TryAnalyserNodeGetFloatTimeDomainData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	array js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AnalyserNode_GetByteTimeDomainData
//go:noescape
func HasFuncAnalyserNodeGetByteTimeDomainData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AnalyserNode_GetByteTimeDomainData
//go:noescape
func FuncAnalyserNodeGetByteTimeDomainData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AnalyserNode_GetByteTimeDomainData
//go:noescape
func CallAnalyserNodeGetByteTimeDomainData(
	this js.Ref, retPtr unsafe.Pointer,
	array js.Ref)

//go:wasmimport plat/js/web try_AnalyserNode_GetByteTimeDomainData
//go:noescape
func TryAnalyserNodeGetByteTimeDomainData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	array js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_AnimationEventInit
//go:noescape
func AnimationEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AnimationEventInit
//go:noescape
func AnimationEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_AnimationEvent_AnimationEvent
//go:noescape
func NewAnimationEventByAnimationEvent(
	typ js.Ref,
	animationEventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_AnimationEvent_AnimationEvent1
//go:noescape
func NewAnimationEventByAnimationEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_AnimationEvent_AnimationName
//go:noescape
func GetAnimationEventAnimationName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AnimationEvent_ElapsedTime
//go:noescape
func GetAnimationEventElapsedTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AnimationEvent_PseudoElement
//go:noescape
func GetAnimationEventPseudoElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_AnimationPlaybackEventInit
//go:noescape
func AnimationPlaybackEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AnimationPlaybackEventInit
//go:noescape
func AnimationPlaybackEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_AnimationPlaybackEvent_AnimationPlaybackEvent
//go:noescape
func NewAnimationPlaybackEventByAnimationPlaybackEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_AnimationPlaybackEvent_AnimationPlaybackEvent1
//go:noescape
func NewAnimationPlaybackEventByAnimationPlaybackEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_AnimationPlaybackEvent_CurrentTime
//go:noescape
func GetAnimationPlaybackEventCurrentTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AnimationPlaybackEvent_TimelineTime
//go:noescape
func GetAnimationPlaybackEventTimelineTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AnimationWorkletGlobalScope_RegisterAnimator
//go:noescape
func HasFuncAnimationWorkletGlobalScopeRegisterAnimator(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AnimationWorkletGlobalScope_RegisterAnimator
//go:noescape
func FuncAnimationWorkletGlobalScopeRegisterAnimator(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AnimationWorkletGlobalScope_RegisterAnimator
//go:noescape
func CallAnimationWorkletGlobalScopeRegisterAnimator(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	animatorCtor js.Ref)

//go:wasmimport plat/js/web try_AnimationWorkletGlobalScope_RegisterAnimator
//go:noescape
func TryAnimationWorkletGlobalScopeRegisterAnimator(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	animatorCtor js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_AppBannerPromptOutcome
//go:noescape
func ConstOfAppBannerPromptOutcome(str js.Ref) uint32

//go:wasmimport plat/js/web constof_AppendMode
//go:noescape
func ConstOfAppendMode(str js.Ref) uint32

//go:wasmimport plat/js/web constof_AttestationConveyancePreference
//go:noescape
func ConstOfAttestationConveyancePreference(str js.Ref) uint32

//go:wasmimport plat/js/web store_AttributionReportingRequestOptions
//go:noescape
func AttributionReportingRequestOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AttributionReportingRequestOptions
//go:noescape
func AttributionReportingRequestOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuctionAd
//go:noescape
func AuctionAdJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuctionAd
//go:noescape
func AuctionAdJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuctionAdConfig
//go:noescape
func AuctionAdConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuctionAdConfig
//go:noescape
func AuctionAdConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuctionAdInterestGroup
//go:noescape
func AuctionAdInterestGroupJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuctionAdInterestGroup
//go:noescape
func AuctionAdInterestGroupJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuctionAdInterestGroupKey
//go:noescape
func AuctionAdInterestGroupKeyJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuctionAdInterestGroupKey
//go:noescape
func AuctionAdInterestGroupKeyJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AudioConfiguration
//go:noescape
func AudioConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioConfiguration
//go:noescape
func AudioConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_AudioContextLatencyCategory
//go:noescape
func ConstOfAudioContextLatencyCategory(str js.Ref) uint32

//go:wasmimport plat/js/web constof_AudioSinkType
//go:noescape
func ConstOfAudioSinkType(str js.Ref) uint32

//go:wasmimport plat/js/web store_AudioSinkOptions
//go:noescape
func AudioSinkOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioSinkOptions
//go:noescape
func AudioSinkOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AudioContextOptions
//go:noescape
func AudioContextOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioContextOptions
//go:noescape
func AudioContextOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AudioTimestamp
//go:noescape
func AudioTimestampJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioTimestamp
//go:noescape
func AudioTimestampJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_CanPlayTypeResult
//go:noescape
func ConstOfCanPlayTypeResult(str js.Ref) uint32

//go:wasmimport plat/js/web get_TextTrackCue_Track
//go:noescape
func GetTextTrackCueTrack(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextTrackCue_Id
//go:noescape
func GetTextTrackCueId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_TextTrackCue_Id
//go:noescape
func SetTextTrackCueId(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_TextTrackCue_StartTime
//go:noescape
func GetTextTrackCueStartTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_TextTrackCue_StartTime
//go:noescape
func SetTextTrackCueStartTime(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_TextTrackCue_EndTime
//go:noescape
func GetTextTrackCueEndTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_TextTrackCue_EndTime
//go:noescape
func SetTextTrackCueEndTime(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_TextTrackCue_PauseOnExit
//go:noescape
func GetTextTrackCuePauseOnExit(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_TextTrackCue_PauseOnExit
//go:noescape
func SetTextTrackCuePauseOnExit(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web constof_TextTrackKind
//go:noescape
func ConstOfTextTrackKind(str js.Ref) uint32
