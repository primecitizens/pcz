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

//go:wasmimport plat/js/webext/speechrecognitionprivate constof_SpeechRecognitionType
//go:noescape
func ConstOfSpeechRecognitionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/speechrecognitionprivate store_SpeechRecognitionErrorEvent
//go:noescape
func SpeechRecognitionErrorEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate load_SpeechRecognitionErrorEvent
//go:noescape
func SpeechRecognitionErrorEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate store_SpeechRecognitionResultEvent
//go:noescape
func SpeechRecognitionResultEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate load_SpeechRecognitionResultEvent
//go:noescape
func SpeechRecognitionResultEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate store_SpeechRecognitionStopEvent
//go:noescape
func SpeechRecognitionStopEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate load_SpeechRecognitionStopEvent
//go:noescape
func SpeechRecognitionStopEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate store_StartOptions
//go:noescape
func StartOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate load_StartOptions
//go:noescape
func StartOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate store_StopOptions
//go:noescape
func StopOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate load_StopOptions
//go:noescape
func StopOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate has_OnError
//go:noescape
func HasFuncOnError() js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate func_OnError
//go:noescape
func FuncOnError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/speechrecognitionprivate call_OnError
//go:noescape
func CallOnError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate try_OnError
//go:noescape
func TryOnError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate has_OffError
//go:noescape
func HasFuncOffError() js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate func_OffError
//go:noescape
func FuncOffError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/speechrecognitionprivate call_OffError
//go:noescape
func CallOffError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate try_OffError
//go:noescape
func TryOffError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate has_HasOnError
//go:noescape
func HasFuncHasOnError() js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate func_HasOnError
//go:noescape
func FuncHasOnError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/speechrecognitionprivate call_HasOnError
//go:noescape
func CallHasOnError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate try_HasOnError
//go:noescape
func TryHasOnError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate has_OnResult
//go:noescape
func HasFuncOnResult() js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate func_OnResult
//go:noescape
func FuncOnResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/speechrecognitionprivate call_OnResult
//go:noescape
func CallOnResult(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate try_OnResult
//go:noescape
func TryOnResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate has_OffResult
//go:noescape
func HasFuncOffResult() js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate func_OffResult
//go:noescape
func FuncOffResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/speechrecognitionprivate call_OffResult
//go:noescape
func CallOffResult(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate try_OffResult
//go:noescape
func TryOffResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate has_HasOnResult
//go:noescape
func HasFuncHasOnResult() js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate func_HasOnResult
//go:noescape
func FuncHasOnResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/speechrecognitionprivate call_HasOnResult
//go:noescape
func CallHasOnResult(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate try_HasOnResult
//go:noescape
func TryHasOnResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate has_OnStop
//go:noescape
func HasFuncOnStop() js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate func_OnStop
//go:noescape
func FuncOnStop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/speechrecognitionprivate call_OnStop
//go:noescape
func CallOnStop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate try_OnStop
//go:noescape
func TryOnStop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate has_OffStop
//go:noescape
func HasFuncOffStop() js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate func_OffStop
//go:noescape
func FuncOffStop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/speechrecognitionprivate call_OffStop
//go:noescape
func CallOffStop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate try_OffStop
//go:noescape
func TryOffStop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate has_HasOnStop
//go:noescape
func HasFuncHasOnStop() js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate func_HasOnStop
//go:noescape
func FuncHasOnStop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/speechrecognitionprivate call_HasOnStop
//go:noescape
func CallHasOnStop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate try_HasOnStop
//go:noescape
func TryHasOnStop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate has_Start
//go:noescape
func HasFuncStart() js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate func_Start
//go:noescape
func FuncStart(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/speechrecognitionprivate call_Start
//go:noescape
func CallStart(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/speechrecognitionprivate try_Start
//go:noescape
func TryStart(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/speechrecognitionprivate has_Stop
//go:noescape
func HasFuncStop() js.Ref

//go:wasmimport plat/js/webext/speechrecognitionprivate func_Stop
//go:noescape
func FuncStop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/speechrecognitionprivate call_Stop
//go:noescape
func CallStop(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/speechrecognitionprivate try_Stop
//go:noescape
func TryStop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)
