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

//go:wasmimport plat/js/webext/ttsengine store_AudioBuffer
//go:noescape
func AudioBufferJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/ttsengine load_AudioBuffer
//go:noescape
func AudioBufferJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/ttsengine store_AudioStreamOptions
//go:noescape
func AudioStreamOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/ttsengine load_AudioStreamOptions
//go:noescape
func AudioStreamOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/ttsengine constof_VoiceGender
//go:noescape
func ConstOfVoiceGender(str js.Ref) uint32

//go:wasmimport plat/js/webext/ttsengine store_SpeakOptions
//go:noescape
func SpeakOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/ttsengine load_SpeakOptions
//go:noescape
func SpeakOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/ttsengine has_OnPause
//go:noescape
func HasFuncOnPause() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_OnPause
//go:noescape
func FuncOnPause(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_OnPause
//go:noescape
func CallOnPause(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_OnPause
//go:noescape
func TryOnPause(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_OffPause
//go:noescape
func HasFuncOffPause() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_OffPause
//go:noescape
func FuncOffPause(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_OffPause
//go:noescape
func CallOffPause(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_OffPause
//go:noescape
func TryOffPause(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_HasOnPause
//go:noescape
func HasFuncHasOnPause() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_HasOnPause
//go:noescape
func FuncHasOnPause(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_HasOnPause
//go:noescape
func CallHasOnPause(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_HasOnPause
//go:noescape
func TryHasOnPause(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_OnResume
//go:noescape
func HasFuncOnResume() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_OnResume
//go:noescape
func FuncOnResume(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_OnResume
//go:noescape
func CallOnResume(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_OnResume
//go:noescape
func TryOnResume(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_OffResume
//go:noescape
func HasFuncOffResume() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_OffResume
//go:noescape
func FuncOffResume(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_OffResume
//go:noescape
func CallOffResume(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_OffResume
//go:noescape
func TryOffResume(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_HasOnResume
//go:noescape
func HasFuncHasOnResume() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_HasOnResume
//go:noescape
func FuncHasOnResume(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_HasOnResume
//go:noescape
func CallHasOnResume(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_HasOnResume
//go:noescape
func TryHasOnResume(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_OnSpeak
//go:noescape
func HasFuncOnSpeak() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_OnSpeak
//go:noescape
func FuncOnSpeak(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_OnSpeak
//go:noescape
func CallOnSpeak(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_OnSpeak
//go:noescape
func TryOnSpeak(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_OffSpeak
//go:noescape
func HasFuncOffSpeak() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_OffSpeak
//go:noescape
func FuncOffSpeak(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_OffSpeak
//go:noescape
func CallOffSpeak(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_OffSpeak
//go:noescape
func TryOffSpeak(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_HasOnSpeak
//go:noescape
func HasFuncHasOnSpeak() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_HasOnSpeak
//go:noescape
func FuncHasOnSpeak(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_HasOnSpeak
//go:noescape
func CallHasOnSpeak(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_HasOnSpeak
//go:noescape
func TryHasOnSpeak(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_OnSpeakWithAudioStream
//go:noescape
func HasFuncOnSpeakWithAudioStream() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_OnSpeakWithAudioStream
//go:noescape
func FuncOnSpeakWithAudioStream(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_OnSpeakWithAudioStream
//go:noescape
func CallOnSpeakWithAudioStream(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_OnSpeakWithAudioStream
//go:noescape
func TryOnSpeakWithAudioStream(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_OffSpeakWithAudioStream
//go:noescape
func HasFuncOffSpeakWithAudioStream() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_OffSpeakWithAudioStream
//go:noescape
func FuncOffSpeakWithAudioStream(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_OffSpeakWithAudioStream
//go:noescape
func CallOffSpeakWithAudioStream(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_OffSpeakWithAudioStream
//go:noescape
func TryOffSpeakWithAudioStream(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_HasOnSpeakWithAudioStream
//go:noescape
func HasFuncHasOnSpeakWithAudioStream() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_HasOnSpeakWithAudioStream
//go:noescape
func FuncHasOnSpeakWithAudioStream(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_HasOnSpeakWithAudioStream
//go:noescape
func CallHasOnSpeakWithAudioStream(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_HasOnSpeakWithAudioStream
//go:noescape
func TryHasOnSpeakWithAudioStream(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_OnStop
//go:noescape
func HasFuncOnStop() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_OnStop
//go:noescape
func FuncOnStop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_OnStop
//go:noescape
func CallOnStop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_OnStop
//go:noescape
func TryOnStop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_OffStop
//go:noescape
func HasFuncOffStop() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_OffStop
//go:noescape
func FuncOffStop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_OffStop
//go:noescape
func CallOffStop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_OffStop
//go:noescape
func TryOffStop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_HasOnStop
//go:noescape
func HasFuncHasOnStop() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_HasOnStop
//go:noescape
func FuncHasOnStop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_HasOnStop
//go:noescape
func CallHasOnStop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_HasOnStop
//go:noescape
func TryHasOnStop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_SendTtsAudio
//go:noescape
func HasFuncSendTtsAudio() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_SendTtsAudio
//go:noescape
func FuncSendTtsAudio(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_SendTtsAudio
//go:noescape
func CallSendTtsAudio(
	retPtr unsafe.Pointer,
	requestId float64,
	audio unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine try_SendTtsAudio
//go:noescape
func TrySendTtsAudio(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId float64,
	audio unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_SendTtsEvent
//go:noescape
func HasFuncSendTtsEvent() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_SendTtsEvent
//go:noescape
func FuncSendTtsEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_SendTtsEvent
//go:noescape
func CallSendTtsEvent(
	retPtr unsafe.Pointer,
	requestId float64,
	event unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine try_SendTtsEvent
//go:noescape
func TrySendTtsEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId float64,
	event unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/ttsengine has_UpdateVoices
//go:noescape
func HasFuncUpdateVoices() js.Ref

//go:wasmimport plat/js/webext/ttsengine func_UpdateVoices
//go:noescape
func FuncUpdateVoices(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/ttsengine call_UpdateVoices
//go:noescape
func CallUpdateVoices(
	retPtr unsafe.Pointer,
	voices js.Ref)

//go:wasmimport plat/js/webext/ttsengine try_UpdateVoices
//go:noescape
func TryUpdateVoices(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	voices js.Ref) (ok js.Ref)
