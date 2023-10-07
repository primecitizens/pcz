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

//go:wasmimport plat/js/webext/webrtcloggingprivate store_MetaDataEntry
//go:noescape
func MetaDataEntryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate load_MetaDataEntry
//go:noescape
func MetaDataEntryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate store_RecordingInfo
//go:noescape
func RecordingInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate load_RecordingInfo
//go:noescape
func RecordingInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate store_RequestInfo
//go:noescape
func RequestInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate load_RequestInfo
//go:noescape
func RequestInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate store_StartEventLoggingResult
//go:noescape
func StartEventLoggingResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate load_StartEventLoggingResult
//go:noescape
func StartEventLoggingResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate store_UploadResult
//go:noescape
func UploadResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate load_UploadResult
//go:noescape
func UploadResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate has_Discard
//go:noescape
func HasFuncDiscard() js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate func_Discard
//go:noescape
func FuncDiscard(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcloggingprivate call_Discard
//go:noescape
func CallDiscard(
	retPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate try_Discard
//go:noescape
func TryDiscard(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate has_GetLogsDirectory
//go:noescape
func HasFuncGetLogsDirectory() js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate func_GetLogsDirectory
//go:noescape
func FuncGetLogsDirectory(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcloggingprivate call_GetLogsDirectory
//go:noescape
func CallGetLogsDirectory(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate try_GetLogsDirectory
//go:noescape
func TryGetLogsDirectory(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate has_SetMetaData
//go:noescape
func HasFuncSetMetaData() js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate func_SetMetaData
//go:noescape
func FuncSetMetaData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcloggingprivate call_SetMetaData
//go:noescape
func CallSetMetaData(
	retPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	metaData js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate try_SetMetaData
//go:noescape
func TrySetMetaData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	metaData js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate has_SetUploadOnRenderClose
//go:noescape
func HasFuncSetUploadOnRenderClose() js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate func_SetUploadOnRenderClose
//go:noescape
func FuncSetUploadOnRenderClose(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcloggingprivate call_SetUploadOnRenderClose
//go:noescape
func CallSetUploadOnRenderClose(
	retPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	shouldUpload js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate try_SetUploadOnRenderClose
//go:noescape
func TrySetUploadOnRenderClose(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	shouldUpload js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate has_Start
//go:noescape
func HasFuncStart() js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate func_Start
//go:noescape
func FuncStart(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcloggingprivate call_Start
//go:noescape
func CallStart(
	retPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate try_Start
//go:noescape
func TryStart(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate has_StartAudioDebugRecordings
//go:noescape
func HasFuncStartAudioDebugRecordings() js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate func_StartAudioDebugRecordings
//go:noescape
func FuncStartAudioDebugRecordings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcloggingprivate call_StartAudioDebugRecordings
//go:noescape
func CallStartAudioDebugRecordings(
	retPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	seconds int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate try_StartAudioDebugRecordings
//go:noescape
func TryStartAudioDebugRecordings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	seconds int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate has_StartEventLogging
//go:noescape
func HasFuncStartEventLogging() js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate func_StartEventLogging
//go:noescape
func FuncStartEventLogging(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcloggingprivate call_StartEventLogging
//go:noescape
func CallStartEventLogging(
	retPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	sessionId js.Ref,
	maxLogSizeBytes int32,
	outputPeriodMs int32,
	webAppId int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate try_StartEventLogging
//go:noescape
func TryStartEventLogging(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	sessionId js.Ref,
	maxLogSizeBytes int32,
	outputPeriodMs int32,
	webAppId int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate has_StartRtpDump
//go:noescape
func HasFuncStartRtpDump() js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate func_StartRtpDump
//go:noescape
func FuncStartRtpDump(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcloggingprivate call_StartRtpDump
//go:noescape
func CallStartRtpDump(
	retPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	incoming js.Ref,
	outgoing js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate try_StartRtpDump
//go:noescape
func TryStartRtpDump(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	incoming js.Ref,
	outgoing js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate has_Stop
//go:noescape
func HasFuncStop() js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate func_Stop
//go:noescape
func FuncStop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcloggingprivate call_Stop
//go:noescape
func CallStop(
	retPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate try_Stop
//go:noescape
func TryStop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate has_StopAudioDebugRecordings
//go:noescape
func HasFuncStopAudioDebugRecordings() js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate func_StopAudioDebugRecordings
//go:noescape
func FuncStopAudioDebugRecordings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcloggingprivate call_StopAudioDebugRecordings
//go:noescape
func CallStopAudioDebugRecordings(
	retPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate try_StopAudioDebugRecordings
//go:noescape
func TryStopAudioDebugRecordings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate has_StopRtpDump
//go:noescape
func HasFuncStopRtpDump() js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate func_StopRtpDump
//go:noescape
func FuncStopRtpDump(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcloggingprivate call_StopRtpDump
//go:noescape
func CallStopRtpDump(
	retPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	incoming js.Ref,
	outgoing js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate try_StopRtpDump
//go:noescape
func TryStopRtpDump(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	incoming js.Ref,
	outgoing js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate has_Store
//go:noescape
func HasFuncStore() js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate func_Store
//go:noescape
func FuncStore(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcloggingprivate call_Store
//go:noescape
func CallStore(
	retPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	logId js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate try_Store
//go:noescape
func TryStore(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	logId js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate has_Upload
//go:noescape
func HasFuncUpload() js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate func_Upload
//go:noescape
func FuncUpload(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcloggingprivate call_Upload
//go:noescape
func CallUpload(
	retPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate try_Upload
//go:noescape
func TryUpload(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate has_UploadStored
//go:noescape
func HasFuncUploadStored() js.Ref

//go:wasmimport plat/js/webext/webrtcloggingprivate func_UploadStored
//go:noescape
func FuncUploadStored(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcloggingprivate call_UploadStored
//go:noescape
func CallUploadStored(
	retPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	logId js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcloggingprivate try_UploadStored
//go:noescape
func TryUploadStored(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer,
	securityOrigin js.Ref,
	logId js.Ref,
	callback js.Ref) (ok js.Ref)
