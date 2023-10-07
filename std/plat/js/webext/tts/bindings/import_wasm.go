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

//go:wasmimport plat/js/webext/tts constof_EventType
//go:noescape
func ConstOfEventType(str js.Ref) uint32

//go:wasmimport plat/js/webext/tts store_TtsEvent
//go:noescape
func TtsEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tts load_TtsEvent
//go:noescape
func TtsEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tts constof_VoiceGender
//go:noescape
func ConstOfVoiceGender(str js.Ref) uint32

//go:wasmimport plat/js/webext/tts store_TtsOptions
//go:noescape
func TtsOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tts load_TtsOptions
//go:noescape
func TtsOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tts store_TtsVoice
//go:noescape
func TtsVoiceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tts load_TtsVoice
//go:noescape
func TtsVoiceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tts has_GetVoices
//go:noescape
func HasFuncGetVoices() js.Ref

//go:wasmimport plat/js/webext/tts func_GetVoices
//go:noescape
func FuncGetVoices(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tts call_GetVoices
//go:noescape
func CallGetVoices(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/tts try_GetVoices
//go:noescape
func TryGetVoices(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tts has_IsSpeaking
//go:noescape
func HasFuncIsSpeaking() js.Ref

//go:wasmimport plat/js/webext/tts func_IsSpeaking
//go:noescape
func FuncIsSpeaking(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tts call_IsSpeaking
//go:noescape
func CallIsSpeaking(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/tts try_IsSpeaking
//go:noescape
func TryIsSpeaking(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tts has_OnEvent
//go:noescape
func HasFuncOnEvent() js.Ref

//go:wasmimport plat/js/webext/tts func_OnEvent
//go:noescape
func FuncOnEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tts call_OnEvent
//go:noescape
func CallOnEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tts try_OnEvent
//go:noescape
func TryOnEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tts has_OffEvent
//go:noescape
func HasFuncOffEvent() js.Ref

//go:wasmimport plat/js/webext/tts func_OffEvent
//go:noescape
func FuncOffEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tts call_OffEvent
//go:noescape
func CallOffEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tts try_OffEvent
//go:noescape
func TryOffEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tts has_HasOnEvent
//go:noescape
func HasFuncHasOnEvent() js.Ref

//go:wasmimport plat/js/webext/tts func_HasOnEvent
//go:noescape
func FuncHasOnEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tts call_HasOnEvent
//go:noescape
func CallHasOnEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tts try_HasOnEvent
//go:noescape
func TryHasOnEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tts has_Pause
//go:noescape
func HasFuncPause() js.Ref

//go:wasmimport plat/js/webext/tts func_Pause
//go:noescape
func FuncPause(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tts call_Pause
//go:noescape
func CallPause(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/tts try_Pause
//go:noescape
func TryPause(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tts has_Resume
//go:noescape
func HasFuncResume() js.Ref

//go:wasmimport plat/js/webext/tts func_Resume
//go:noescape
func FuncResume(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tts call_Resume
//go:noescape
func CallResume(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/tts try_Resume
//go:noescape
func TryResume(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tts has_Speak
//go:noescape
func HasFuncSpeak() js.Ref

//go:wasmimport plat/js/webext/tts func_Speak
//go:noescape
func FuncSpeak(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tts call_Speak
//go:noescape
func CallSpeak(
	retPtr unsafe.Pointer,
	utterance js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/tts try_Speak
//go:noescape
func TrySpeak(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	utterance js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tts has_Stop
//go:noescape
func HasFuncStop() js.Ref

//go:wasmimport plat/js/webext/tts func_Stop
//go:noescape
func FuncStop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tts call_Stop
//go:noescape
func CallStop(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/tts try_Stop
//go:noescape
func TryStop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
