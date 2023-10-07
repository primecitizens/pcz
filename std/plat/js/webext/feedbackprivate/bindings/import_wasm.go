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

//go:wasmimport plat/js/webext/feedbackprivate store_AttachedFile
//go:noescape
func AttachedFileJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate load_AttachedFile
//go:noescape
func AttachedFileJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/feedbackprivate constof_FeedbackFlow
//go:noescape
func ConstOfFeedbackFlow(str js.Ref) uint32

//go:wasmimport plat/js/webext/feedbackprivate store_LogsMapEntry
//go:noescape
func LogsMapEntryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate load_LogsMapEntry
//go:noescape
func LogsMapEntryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/feedbackprivate store_FeedbackInfo
//go:noescape
func FeedbackInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate load_FeedbackInfo
//go:noescape
func FeedbackInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/feedbackprivate constof_FeedbackSource
//go:noescape
func ConstOfFeedbackSource(str js.Ref) uint32

//go:wasmimport plat/js/webext/feedbackprivate constof_LandingPageType
//go:noescape
func ConstOfLandingPageType(str js.Ref) uint32

//go:wasmimport plat/js/webext/feedbackprivate constof_LogSource
//go:noescape
func ConstOfLogSource(str js.Ref) uint32

//go:wasmimport plat/js/webext/feedbackprivate store_ReadLogSourceResult
//go:noescape
func ReadLogSourceResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate load_ReadLogSourceResult
//go:noescape
func ReadLogSourceResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/feedbackprivate store_ReadLogSourceParams
//go:noescape
func ReadLogSourceParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate load_ReadLogSourceParams
//go:noescape
func ReadLogSourceParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/feedbackprivate constof_Status
//go:noescape
func ConstOfStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/feedbackprivate store_SendFeedbackResult
//go:noescape
func SendFeedbackResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate load_SendFeedbackResult
//go:noescape
func SendFeedbackResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/feedbackprivate has_GetSystemInformation
//go:noescape
func HasFuncGetSystemInformation() js.Ref

//go:wasmimport plat/js/webext/feedbackprivate func_GetSystemInformation
//go:noescape
func FuncGetSystemInformation(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/feedbackprivate call_GetSystemInformation
//go:noescape
func CallGetSystemInformation(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate try_GetSystemInformation
//go:noescape
func TryGetSystemInformation(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate has_GetUserEmail
//go:noescape
func HasFuncGetUserEmail() js.Ref

//go:wasmimport plat/js/webext/feedbackprivate func_GetUserEmail
//go:noescape
func FuncGetUserEmail(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/feedbackprivate call_GetUserEmail
//go:noescape
func CallGetUserEmail(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate try_GetUserEmail
//go:noescape
func TryGetUserEmail(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate has_OnFeedbackRequested
//go:noescape
func HasFuncOnFeedbackRequested() js.Ref

//go:wasmimport plat/js/webext/feedbackprivate func_OnFeedbackRequested
//go:noescape
func FuncOnFeedbackRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/feedbackprivate call_OnFeedbackRequested
//go:noescape
func CallOnFeedbackRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate try_OnFeedbackRequested
//go:noescape
func TryOnFeedbackRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate has_OffFeedbackRequested
//go:noescape
func HasFuncOffFeedbackRequested() js.Ref

//go:wasmimport plat/js/webext/feedbackprivate func_OffFeedbackRequested
//go:noescape
func FuncOffFeedbackRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/feedbackprivate call_OffFeedbackRequested
//go:noescape
func CallOffFeedbackRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate try_OffFeedbackRequested
//go:noescape
func TryOffFeedbackRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate has_HasOnFeedbackRequested
//go:noescape
func HasFuncHasOnFeedbackRequested() js.Ref

//go:wasmimport plat/js/webext/feedbackprivate func_HasOnFeedbackRequested
//go:noescape
func FuncHasOnFeedbackRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/feedbackprivate call_HasOnFeedbackRequested
//go:noescape
func CallHasOnFeedbackRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate try_HasOnFeedbackRequested
//go:noescape
func TryHasOnFeedbackRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate has_OpenFeedback
//go:noescape
func HasFuncOpenFeedback() js.Ref

//go:wasmimport plat/js/webext/feedbackprivate func_OpenFeedback
//go:noescape
func FuncOpenFeedback(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/feedbackprivate call_OpenFeedback
//go:noescape
func CallOpenFeedback(
	retPtr unsafe.Pointer,
	source uint32)

//go:wasmimport plat/js/webext/feedbackprivate try_OpenFeedback
//go:noescape
func TryOpenFeedback(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate has_ReadLogSource
//go:noescape
func HasFuncReadLogSource() js.Ref

//go:wasmimport plat/js/webext/feedbackprivate func_ReadLogSource
//go:noescape
func FuncReadLogSource(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/feedbackprivate call_ReadLogSource
//go:noescape
func CallReadLogSource(
	retPtr unsafe.Pointer,
	params unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate try_ReadLogSource
//go:noescape
func TryReadLogSource(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	params unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/feedbackprivate has_SendFeedback
//go:noescape
func HasFuncSendFeedback() js.Ref

//go:wasmimport plat/js/webext/feedbackprivate func_SendFeedback
//go:noescape
func FuncSendFeedback(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/feedbackprivate call_SendFeedback
//go:noescape
func CallSendFeedback(
	retPtr unsafe.Pointer,
	feedback unsafe.Pointer,
	loadSystemInfo js.Ref,
	formOpenTime float64)

//go:wasmimport plat/js/webext/feedbackprivate try_SendFeedback
//go:noescape
func TrySendFeedback(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	feedback unsafe.Pointer,
	loadSystemInfo js.Ref,
	formOpenTime float64) (ok js.Ref)
