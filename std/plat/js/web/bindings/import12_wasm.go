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

//go:wasmimport plat/js/web constof_RequestDestination
//go:noescape
func ConstOfRequestDestination(str js.Ref) uint32

//go:wasmimport plat/js/web new_Request_Request
//go:noescape
func NewRequestByRequest(
	input js.Ref,
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_Request_Request1
//go:noescape
func NewRequestByRequest1(
	input js.Ref) js.Ref

//go:wasmimport plat/js/web get_Request_Method
//go:noescape
func GetRequestMethod(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_Url
//go:noescape
func GetRequestUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_Headers
//go:noescape
func GetRequestHeaders(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_Destination
//go:noescape
func GetRequestDestination(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_Referrer
//go:noescape
func GetRequestReferrer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_ReferrerPolicy
//go:noescape
func GetRequestReferrerPolicy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_Mode
//go:noescape
func GetRequestMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_Credentials
//go:noescape
func GetRequestCredentials(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_Cache
//go:noescape
func GetRequestCache(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_Redirect
//go:noescape
func GetRequestRedirect(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_Integrity
//go:noescape
func GetRequestIntegrity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_Keepalive
//go:noescape
func GetRequestKeepalive(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_IsReloadNavigation
//go:noescape
func GetRequestIsReloadNavigation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_IsHistoryNavigation
//go:noescape
func GetRequestIsHistoryNavigation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_Signal
//go:noescape
func GetRequestSignal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_Duplex
//go:noescape
func GetRequestDuplex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_Body
//go:noescape
func GetRequestBody(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Request_BodyUsed
//go:noescape
func GetRequestBodyUsed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Request_Clone
//go:noescape
func HasFuncRequestClone(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Request_Clone
//go:noescape
func FuncRequestClone(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Request_Clone
//go:noescape
func CallRequestClone(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Request_Clone
//go:noescape
func TryRequestClone(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Request_ArrayBuffer
//go:noescape
func HasFuncRequestArrayBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Request_ArrayBuffer
//go:noescape
func FuncRequestArrayBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Request_ArrayBuffer
//go:noescape
func CallRequestArrayBuffer(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Request_ArrayBuffer
//go:noescape
func TryRequestArrayBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Request_Blob
//go:noescape
func HasFuncRequestBlob(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Request_Blob
//go:noescape
func FuncRequestBlob(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Request_Blob
//go:noescape
func CallRequestBlob(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Request_Blob
//go:noescape
func TryRequestBlob(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Request_FormData
//go:noescape
func HasFuncRequestFormData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Request_FormData
//go:noescape
func FuncRequestFormData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Request_FormData
//go:noescape
func CallRequestFormData(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Request_FormData
//go:noescape
func TryRequestFormData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Request_Json
//go:noescape
func HasFuncRequestJson(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Request_Json
//go:noescape
func FuncRequestJson(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Request_Json
//go:noescape
func CallRequestJson(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Request_Json
//go:noescape
func TryRequestJson(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Request_Text
//go:noescape
func HasFuncRequestText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Request_Text
//go:noescape
func FuncRequestText(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Request_Text
//go:noescape
func CallRequestText(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Request_Text
//go:noescape
func TryRequestText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ResponseInit
//go:noescape
func ResponseInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ResponseInit
//go:noescape
func ResponseInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_ResponseType
//go:noescape
func ConstOfResponseType(str js.Ref) uint32

//go:wasmimport plat/js/web new_Response_Response
//go:noescape
func NewResponseByResponse(
	body js.Ref,
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_Response_Response1
//go:noescape
func NewResponseByResponse1(
	body js.Ref) js.Ref

//go:wasmimport plat/js/web new_Response_Response2
//go:noescape
func NewResponseByResponse2() js.Ref

//go:wasmimport plat/js/web get_Response_Type
//go:noescape
func GetResponseType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Response_Url
//go:noescape
func GetResponseUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Response_Redirected
//go:noescape
func GetResponseRedirected(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Response_Status
//go:noescape
func GetResponseStatus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Response_Ok
//go:noescape
func GetResponseOk(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Response_StatusText
//go:noescape
func GetResponseStatusText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Response_Headers
//go:noescape
func GetResponseHeaders(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Response_Body
//go:noescape
func GetResponseBody(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Response_BodyUsed
//go:noescape
func GetResponseBodyUsed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Response_Error
//go:noescape
func HasFuncResponseError(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Response_Error
//go:noescape
func FuncResponseError(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Response_Error
//go:noescape
func CallResponseError(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Response_Error
//go:noescape
func TryResponseError(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Response_Redirect
//go:noescape
func HasFuncResponseRedirect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Response_Redirect
//go:noescape
func FuncResponseRedirect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Response_Redirect
//go:noescape
func CallResponseRedirect(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref,
	status uint32)

//go:wasmimport plat/js/web try_Response_Redirect
//go:noescape
func TryResponseRedirect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref,
	status uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Response_Redirect1
//go:noescape
func HasFuncResponseRedirect1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Response_Redirect1
//go:noescape
func FuncResponseRedirect1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Response_Redirect1
//go:noescape
func CallResponseRedirect1(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/web try_Response_Redirect1
//go:noescape
func TryResponseRedirect1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Response_Json
//go:noescape
func HasFuncResponseJson(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Response_Json
//go:noescape
func FuncResponseJson(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Response_Json
//go:noescape
func CallResponseJson(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref,
	init unsafe.Pointer)

//go:wasmimport plat/js/web try_Response_Json
//go:noescape
func TryResponseJson(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref,
	init unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Response_Json1
//go:noescape
func HasFuncResponseJson1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Response_Json1
//go:noescape
func FuncResponseJson1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Response_Json1
//go:noescape
func CallResponseJson1(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_Response_Json1
//go:noescape
func TryResponseJson1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Response_Clone
//go:noescape
func HasFuncResponseClone(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Response_Clone
//go:noescape
func FuncResponseClone(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Response_Clone
//go:noescape
func CallResponseClone(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Response_Clone
//go:noescape
func TryResponseClone(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Response_ArrayBuffer
//go:noescape
func HasFuncResponseArrayBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Response_ArrayBuffer
//go:noescape
func FuncResponseArrayBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Response_ArrayBuffer
//go:noescape
func CallResponseArrayBuffer(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Response_ArrayBuffer
//go:noescape
func TryResponseArrayBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Response_Blob
//go:noescape
func HasFuncResponseBlob(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Response_Blob
//go:noescape
func FuncResponseBlob(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Response_Blob
//go:noescape
func CallResponseBlob(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Response_Blob
//go:noescape
func TryResponseBlob(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Response_FormData
//go:noescape
func HasFuncResponseFormData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Response_FormData
//go:noescape
func FuncResponseFormData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Response_FormData
//go:noescape
func CallResponseFormData(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Response_FormData
//go:noescape
func TryResponseFormData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Response_Json2
//go:noescape
func HasFuncResponseJson2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Response_Json2
//go:noescape
func FuncResponseJson2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Response_Json2
//go:noescape
func CallResponseJson2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Response_Json2
//go:noescape
func TryResponseJson2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Response_Text
//go:noescape
func HasFuncResponseText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Response_Text
//go:noescape
func FuncResponseText(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Response_Text
//go:noescape
func CallResponseText(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Response_Text
//go:noescape
func TryResponseText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BackgroundFetchRecord_Request
//go:noescape
func GetBackgroundFetchRecordRequest(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BackgroundFetchRecord_ResponseReady
//go:noescape
func GetBackgroundFetchRecordResponseReady(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_CacheQueryOptions
//go:noescape
func CacheQueryOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CacheQueryOptions
//go:noescape
func CacheQueryOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_BackgroundFetchResult
//go:noescape
func ConstOfBackgroundFetchResult(str js.Ref) uint32

//go:wasmimport plat/js/web constof_BackgroundFetchFailureReason
//go:noescape
func ConstOfBackgroundFetchFailureReason(str js.Ref) uint32

//go:wasmimport plat/js/web get_BackgroundFetchRegistration_Id
//go:noescape
func GetBackgroundFetchRegistrationId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BackgroundFetchRegistration_UploadTotal
//go:noescape
func GetBackgroundFetchRegistrationUploadTotal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BackgroundFetchRegistration_Uploaded
//go:noescape
func GetBackgroundFetchRegistrationUploaded(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BackgroundFetchRegistration_DownloadTotal
//go:noescape
func GetBackgroundFetchRegistrationDownloadTotal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BackgroundFetchRegistration_Downloaded
//go:noescape
func GetBackgroundFetchRegistrationDownloaded(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BackgroundFetchRegistration_Result
//go:noescape
func GetBackgroundFetchRegistrationResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BackgroundFetchRegistration_FailureReason
//go:noescape
func GetBackgroundFetchRegistrationFailureReason(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BackgroundFetchRegistration_RecordsAvailable
//go:noescape
func GetBackgroundFetchRegistrationRecordsAvailable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BackgroundFetchRegistration_Abort
//go:noescape
func HasFuncBackgroundFetchRegistrationAbort(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BackgroundFetchRegistration_Abort
//go:noescape
func FuncBackgroundFetchRegistrationAbort(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BackgroundFetchRegistration_Abort
//go:noescape
func CallBackgroundFetchRegistrationAbort(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BackgroundFetchRegistration_Abort
//go:noescape
func TryBackgroundFetchRegistrationAbort(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BackgroundFetchRegistration_Match
//go:noescape
func HasFuncBackgroundFetchRegistrationMatch(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BackgroundFetchRegistration_Match
//go:noescape
func FuncBackgroundFetchRegistrationMatch(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BackgroundFetchRegistration_Match
//go:noescape
func CallBackgroundFetchRegistrationMatch(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_BackgroundFetchRegistration_Match
//go:noescape
func TryBackgroundFetchRegistrationMatch(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BackgroundFetchRegistration_Match1
//go:noescape
func HasFuncBackgroundFetchRegistrationMatch1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BackgroundFetchRegistration_Match1
//go:noescape
func FuncBackgroundFetchRegistrationMatch1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BackgroundFetchRegistration_Match1
//go:noescape
func CallBackgroundFetchRegistrationMatch1(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref)

//go:wasmimport plat/js/web try_BackgroundFetchRegistration_Match1
//go:noescape
func TryBackgroundFetchRegistrationMatch1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BackgroundFetchRegistration_MatchAll
//go:noescape
func HasFuncBackgroundFetchRegistrationMatchAll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BackgroundFetchRegistration_MatchAll
//go:noescape
func FuncBackgroundFetchRegistrationMatchAll(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BackgroundFetchRegistration_MatchAll
//go:noescape
func CallBackgroundFetchRegistrationMatchAll(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_BackgroundFetchRegistration_MatchAll
//go:noescape
func TryBackgroundFetchRegistrationMatchAll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BackgroundFetchRegistration_MatchAll1
//go:noescape
func HasFuncBackgroundFetchRegistrationMatchAll1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BackgroundFetchRegistration_MatchAll1
//go:noescape
func FuncBackgroundFetchRegistrationMatchAll1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BackgroundFetchRegistration_MatchAll1
//go:noescape
func CallBackgroundFetchRegistrationMatchAll1(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref)

//go:wasmimport plat/js/web try_BackgroundFetchRegistration_MatchAll1
//go:noescape
func TryBackgroundFetchRegistrationMatchAll1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BackgroundFetchRegistration_MatchAll2
//go:noescape
func HasFuncBackgroundFetchRegistrationMatchAll2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BackgroundFetchRegistration_MatchAll2
//go:noescape
func FuncBackgroundFetchRegistrationMatchAll2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BackgroundFetchRegistration_MatchAll2
//go:noescape
func CallBackgroundFetchRegistrationMatchAll2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BackgroundFetchRegistration_MatchAll2
//go:noescape
func TryBackgroundFetchRegistrationMatchAll2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_BackgroundFetchEventInit
//go:noescape
func BackgroundFetchEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BackgroundFetchEventInit
//go:noescape
func BackgroundFetchEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_BackgroundFetchEvent_BackgroundFetchEvent
//go:noescape
func NewBackgroundFetchEventByBackgroundFetchEvent(
	typ js.Ref,
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_BackgroundFetchEvent_Registration
//go:noescape
func GetBackgroundFetchEventRegistration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ImageResource
//go:noescape
func ImageResourceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ImageResource
//go:noescape
func ImageResourceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_BackgroundFetchOptions
//go:noescape
func BackgroundFetchOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BackgroundFetchOptions
//go:noescape
func BackgroundFetchOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_BackgroundFetchManager_Fetch
//go:noescape
func HasFuncBackgroundFetchManagerFetch(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BackgroundFetchManager_Fetch
//go:noescape
func FuncBackgroundFetchManagerFetch(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BackgroundFetchManager_Fetch
//go:noescape
func CallBackgroundFetchManagerFetch(
	this js.Ref, retPtr unsafe.Pointer,
	id js.Ref,
	requests js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_BackgroundFetchManager_Fetch
//go:noescape
func TryBackgroundFetchManagerFetch(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	requests js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BackgroundFetchManager_Fetch1
//go:noescape
func HasFuncBackgroundFetchManagerFetch1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BackgroundFetchManager_Fetch1
//go:noescape
func FuncBackgroundFetchManagerFetch1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BackgroundFetchManager_Fetch1
//go:noescape
func CallBackgroundFetchManagerFetch1(
	this js.Ref, retPtr unsafe.Pointer,
	id js.Ref,
	requests js.Ref)

//go:wasmimport plat/js/web try_BackgroundFetchManager_Fetch1
//go:noescape
func TryBackgroundFetchManagerFetch1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	requests js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BackgroundFetchManager_Get
//go:noescape
func HasFuncBackgroundFetchManagerGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BackgroundFetchManager_Get
//go:noescape
func FuncBackgroundFetchManagerGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BackgroundFetchManager_Get
//go:noescape
func CallBackgroundFetchManagerGet(
	this js.Ref, retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/web try_BackgroundFetchManager_Get
//go:noescape
func TryBackgroundFetchManagerGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BackgroundFetchManager_GetIds
//go:noescape
func HasFuncBackgroundFetchManagerGetIds(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BackgroundFetchManager_GetIds
//go:noescape
func FuncBackgroundFetchManagerGetIds(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BackgroundFetchManager_GetIds
//go:noescape
func CallBackgroundFetchManagerGetIds(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BackgroundFetchManager_GetIds
//go:noescape
func TryBackgroundFetchManagerGetIds(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_BackgroundFetchUIOptions
//go:noescape
func BackgroundFetchUIOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BackgroundFetchUIOptions
//go:noescape
func BackgroundFetchUIOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_BackgroundFetchUpdateUIEvent_BackgroundFetchUpdateUIEvent
//go:noescape
func NewBackgroundFetchUpdateUIEventByBackgroundFetchUpdateUIEvent(
	typ js.Ref,
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_BackgroundFetchUpdateUIEvent_UpdateUI
//go:noescape
func HasFuncBackgroundFetchUpdateUIEventUpdateUI(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BackgroundFetchUpdateUIEvent_UpdateUI
//go:noescape
func FuncBackgroundFetchUpdateUIEventUpdateUI(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BackgroundFetchUpdateUIEvent_UpdateUI
//go:noescape
func CallBackgroundFetchUpdateUIEventUpdateUI(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_BackgroundFetchUpdateUIEvent_UpdateUI
//go:noescape
func TryBackgroundFetchUpdateUIEventUpdateUI(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BackgroundFetchUpdateUIEvent_UpdateUI1
//go:noescape
func HasFuncBackgroundFetchUpdateUIEventUpdateUI1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BackgroundFetchUpdateUIEvent_UpdateUI1
//go:noescape
func FuncBackgroundFetchUpdateUIEventUpdateUI1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BackgroundFetchUpdateUIEvent_UpdateUI1
//go:noescape
func CallBackgroundFetchUpdateUIEventUpdateUI1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BackgroundFetchUpdateUIEvent_UpdateUI1
//go:noescape
func TryBackgroundFetchUpdateUIEventUpdateUI1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_BackgroundSyncOptions
//go:noescape
func BackgroundSyncOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BackgroundSyncOptions
//go:noescape
func BackgroundSyncOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_BarProp_Visible
//go:noescape
func GetBarPropVisible(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_BarcodeFormat
//go:noescape
func ConstOfBarcodeFormat(str js.Ref) uint32

//go:wasmimport plat/js/web store_BarcodeDetectorOptions
//go:noescape
func BarcodeDetectorOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BarcodeDetectorOptions
//go:noescape
func BarcodeDetectorOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_DetectedBarcode
//go:noescape
func DetectedBarcodeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DetectedBarcode
//go:noescape
func DetectedBarcodeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_Alt
//go:noescape
func GetHTMLImageElementAlt(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_Alt
//go:noescape
func SetHTMLImageElementAlt(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_Src
//go:noescape
func GetHTMLImageElementSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_Src
//go:noescape
func SetHTMLImageElementSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_Srcset
//go:noescape
func GetHTMLImageElementSrcset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_Srcset
//go:noescape
func SetHTMLImageElementSrcset(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_Sizes
//go:noescape
func GetHTMLImageElementSizes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_Sizes
//go:noescape
func SetHTMLImageElementSizes(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_CrossOrigin
//go:noescape
func GetHTMLImageElementCrossOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_CrossOrigin
//go:noescape
func SetHTMLImageElementCrossOrigin(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_UseMap
//go:noescape
func GetHTMLImageElementUseMap(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_UseMap
//go:noescape
func SetHTMLImageElementUseMap(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_IsMap
//go:noescape
func GetHTMLImageElementIsMap(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_IsMap
//go:noescape
func SetHTMLImageElementIsMap(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_Width
//go:noescape
func GetHTMLImageElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_Width
//go:noescape
func SetHTMLImageElementWidth(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_Height
//go:noescape
func GetHTMLImageElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_Height
//go:noescape
func SetHTMLImageElementHeight(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_NaturalWidth
//go:noescape
func GetHTMLImageElementNaturalWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLImageElement_NaturalHeight
//go:noescape
func GetHTMLImageElementNaturalHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLImageElement_Complete
//go:noescape
func GetHTMLImageElementComplete(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLImageElement_CurrentSrc
//go:noescape
func GetHTMLImageElementCurrentSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLImageElement_ReferrerPolicy
//go:noescape
func GetHTMLImageElementReferrerPolicy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_ReferrerPolicy
//go:noescape
func SetHTMLImageElementReferrerPolicy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_Decoding
//go:noescape
func GetHTMLImageElementDecoding(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_Decoding
//go:noescape
func SetHTMLImageElementDecoding(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_Loading
//go:noescape
func GetHTMLImageElementLoading(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_Loading
//go:noescape
func SetHTMLImageElementLoading(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_FetchPriority
//go:noescape
func GetHTMLImageElementFetchPriority(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_FetchPriority
//go:noescape
func SetHTMLImageElementFetchPriority(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_X
//go:noescape
func GetHTMLImageElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLImageElement_Y
//go:noescape
func GetHTMLImageElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLImageElement_Name
//go:noescape
func GetHTMLImageElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_Name
//go:noescape
func SetHTMLImageElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_Lowsrc
//go:noescape
func GetHTMLImageElementLowsrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_Lowsrc
//go:noescape
func SetHTMLImageElementLowsrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_Align
//go:noescape
func GetHTMLImageElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_Align
//go:noescape
func SetHTMLImageElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_Hspace
//go:noescape
func GetHTMLImageElementHspace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_Hspace
//go:noescape
func SetHTMLImageElementHspace(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_Vspace
//go:noescape
func GetHTMLImageElementVspace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_Vspace
//go:noescape
func SetHTMLImageElementVspace(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_LongDesc
//go:noescape
func GetHTMLImageElementLongDesc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_LongDesc
//go:noescape
func SetHTMLImageElementLongDesc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_Border
//go:noescape
func GetHTMLImageElementBorder(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_Border
//go:noescape
func SetHTMLImageElementBorder(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLImageElement_AttributionSrc
//go:noescape
func GetHTMLImageElementAttributionSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLImageElement_AttributionSrc
//go:noescape
func SetHTMLImageElementAttributionSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLImageElement_Decode
//go:noescape
func HasFuncHTMLImageElementDecode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLImageElement_Decode
//go:noescape
func FuncHTMLImageElementDecode(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLImageElement_Decode
//go:noescape
func CallHTMLImageElementDecode(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLImageElement_Decode
//go:noescape
func TryHTMLImageElementDecode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGImageElement_X
//go:noescape
func GetSVGImageElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGImageElement_Y
//go:noescape
func GetSVGImageElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGImageElement_Width
//go:noescape
func GetSVGImageElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGImageElement_Height
//go:noescape
func GetSVGImageElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGImageElement_PreserveAspectRatio
//go:noescape
func GetSVGImageElementPreserveAspectRatio(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGImageElement_CrossOrigin
//go:noescape
func GetSVGImageElementCrossOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGImageElement_CrossOrigin
//go:noescape
func SetSVGImageElementCrossOrigin(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGImageElement_Href
//go:noescape
func GetSVGImageElementHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_VideoFrameCallbackMetadata
//go:noescape
func VideoFrameCallbackMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoFrameCallbackMetadata
//go:noescape
func VideoFrameCallbackMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_PictureInPictureWindow_Width
//go:noescape
func GetPictureInPictureWindowWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PictureInPictureWindow_Height
//go:noescape
func GetPictureInPictureWindowHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VideoPlaybackQuality_CreationTime
//go:noescape
func GetVideoPlaybackQualityCreationTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VideoPlaybackQuality_DroppedVideoFrames
//go:noescape
func GetVideoPlaybackQualityDroppedVideoFrames(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VideoPlaybackQuality_TotalVideoFrames
//go:noescape
func GetVideoPlaybackQualityTotalVideoFrames(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VideoPlaybackQuality_CorruptedVideoFrames
//go:noescape
func GetVideoPlaybackQualityCorruptedVideoFrames(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLVideoElement_Width
//go:noescape
func GetHTMLVideoElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLVideoElement_Width
//go:noescape
func SetHTMLVideoElementWidth(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLVideoElement_Height
//go:noescape
func GetHTMLVideoElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLVideoElement_Height
//go:noescape
func SetHTMLVideoElementHeight(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLVideoElement_VideoWidth
//go:noescape
func GetHTMLVideoElementVideoWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLVideoElement_VideoHeight
//go:noescape
func GetHTMLVideoElementVideoHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLVideoElement_Poster
//go:noescape
func GetHTMLVideoElementPoster(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLVideoElement_Poster
//go:noescape
func SetHTMLVideoElementPoster(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLVideoElement_PlaysInline
//go:noescape
func GetHTMLVideoElementPlaysInline(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLVideoElement_PlaysInline
//go:noescape
func SetHTMLVideoElementPlaysInline(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLVideoElement_DisablePictureInPicture
//go:noescape
func GetHTMLVideoElementDisablePictureInPicture(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLVideoElement_DisablePictureInPicture
//go:noescape
func SetHTMLVideoElementDisablePictureInPicture(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLVideoElement_RequestVideoFrameCallback
//go:noescape
func HasFuncHTMLVideoElementRequestVideoFrameCallback(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLVideoElement_RequestVideoFrameCallback
//go:noescape
func FuncHTMLVideoElementRequestVideoFrameCallback(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLVideoElement_RequestVideoFrameCallback
//go:noescape
func CallHTMLVideoElementRequestVideoFrameCallback(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/web try_HTMLVideoElement_RequestVideoFrameCallback
//go:noescape
func TryHTMLVideoElementRequestVideoFrameCallback(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLVideoElement_CancelVideoFrameCallback
//go:noescape
func HasFuncHTMLVideoElementCancelVideoFrameCallback(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLVideoElement_CancelVideoFrameCallback
//go:noescape
func FuncHTMLVideoElementCancelVideoFrameCallback(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLVideoElement_CancelVideoFrameCallback
//go:noescape
func CallHTMLVideoElementCancelVideoFrameCallback(
	this js.Ref, retPtr unsafe.Pointer,
	handle uint32)

//go:wasmimport plat/js/web try_HTMLVideoElement_CancelVideoFrameCallback
//go:noescape
func TryHTMLVideoElementCancelVideoFrameCallback(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLVideoElement_RequestPictureInPicture
//go:noescape
func HasFuncHTMLVideoElementRequestPictureInPicture(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLVideoElement_RequestPictureInPicture
//go:noescape
func FuncHTMLVideoElementRequestPictureInPicture(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLVideoElement_RequestPictureInPicture
//go:noescape
func CallHTMLVideoElementRequestPictureInPicture(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLVideoElement_RequestPictureInPicture
//go:noescape
func TryHTMLVideoElementRequestPictureInPicture(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLVideoElement_GetVideoPlaybackQuality
//go:noescape
func HasFuncHTMLVideoElementGetVideoPlaybackQuality(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLVideoElement_GetVideoPlaybackQuality
//go:noescape
func FuncHTMLVideoElementGetVideoPlaybackQuality(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLVideoElement_GetVideoPlaybackQuality
//go:noescape
func CallHTMLVideoElementGetVideoPlaybackQuality(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLVideoElement_GetVideoPlaybackQuality
//go:noescape
func TryHTMLVideoElementGetVideoPlaybackQuality(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_PredefinedColorSpace
//go:noescape
func ConstOfPredefinedColorSpace(str js.Ref) uint32

//go:wasmimport plat/js/web store_CanvasRenderingContext2DSettings
//go:noescape
func CanvasRenderingContext2DSettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CanvasRenderingContext2DSettings
//go:noescape
func CanvasRenderingContext2DSettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_CanvasGradient_AddColorStop
//go:noescape
func HasFuncCanvasGradientAddColorStop(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasGradient_AddColorStop
//go:noescape
func FuncCanvasGradientAddColorStop(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CanvasGradient_AddColorStop
//go:noescape
func CallCanvasGradientAddColorStop(
	this js.Ref, retPtr unsafe.Pointer,
	offset float64,
	color js.Ref)

//go:wasmimport plat/js/web try_CanvasGradient_AddColorStop
//go:noescape
func TryCanvasGradientAddColorStop(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	offset float64,
	color js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasPattern_SetTransform
//go:noescape
func HasFuncCanvasPatternSetTransform(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasPattern_SetTransform
//go:noescape
func FuncCanvasPatternSetTransform(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CanvasPattern_SetTransform
//go:noescape
func CallCanvasPatternSetTransform(
	this js.Ref, retPtr unsafe.Pointer,
	transform unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasPattern_SetTransform
//go:noescape
func TryCanvasPatternSetTransform(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	transform unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasPattern_SetTransform1
//go:noescape
func HasFuncCanvasPatternSetTransform1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasPattern_SetTransform1
//go:noescape
func FuncCanvasPatternSetTransform1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CanvasPattern_SetTransform1
//go:noescape
func CallCanvasPatternSetTransform1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasPattern_SetTransform1
//go:noescape
func TryCanvasPatternSetTransform1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ImageBitmap_Width
//go:noescape
func GetImageBitmapWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ImageBitmap_Height
//go:noescape
func GetImageBitmapHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ImageBitmap_Close
//go:noescape
func HasFuncImageBitmapClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ImageBitmap_Close
//go:noescape
func FuncImageBitmapClose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ImageBitmap_Close
//go:noescape
func CallImageBitmapClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ImageBitmap_Close
//go:noescape
func TryImageBitmapClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_CanvasFillRule
//go:noescape
func ConstOfCanvasFillRule(str js.Ref) uint32

//go:wasmimport plat/js/web new_Path2D_Path2D
//go:noescape
func NewPath2DByPath2D(
	path js.Ref) js.Ref

//go:wasmimport plat/js/web new_Path2D_Path2D1
//go:noescape
func NewPath2DByPath2D1() js.Ref

//go:wasmimport plat/js/web has_Path2D_AddPath
//go:noescape
func HasFuncPath2DAddPath(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Path2D_AddPath
//go:noescape
func FuncPath2DAddPath(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Path2D_AddPath
//go:noescape
func CallPath2DAddPath(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	transform unsafe.Pointer)

//go:wasmimport plat/js/web try_Path2D_AddPath
//go:noescape
func TryPath2DAddPath(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	transform unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Path2D_AddPath1
//go:noescape
func HasFuncPath2DAddPath1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Path2D_AddPath1
//go:noescape
func FuncPath2DAddPath1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Path2D_AddPath1
//go:noescape
func CallPath2DAddPath1(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref)

//go:wasmimport plat/js/web try_Path2D_AddPath1
//go:noescape
func TryPath2DAddPath1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Path2D_ClosePath
//go:noescape
func HasFuncPath2DClosePath(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Path2D_ClosePath
//go:noescape
func FuncPath2DClosePath(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Path2D_ClosePath
//go:noescape
func CallPath2DClosePath(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Path2D_ClosePath
//go:noescape
func TryPath2DClosePath(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Path2D_MoveTo
//go:noescape
func HasFuncPath2DMoveTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Path2D_MoveTo
//go:noescape
func FuncPath2DMoveTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Path2D_MoveTo
//go:noescape
func CallPath2DMoveTo(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_Path2D_MoveTo
//go:noescape
func TryPath2DMoveTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Path2D_LineTo
//go:noescape
func HasFuncPath2DLineTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Path2D_LineTo
//go:noescape
func FuncPath2DLineTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Path2D_LineTo
//go:noescape
func CallPath2DLineTo(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_Path2D_LineTo
//go:noescape
func TryPath2DLineTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Path2D_QuadraticCurveTo
//go:noescape
func HasFuncPath2DQuadraticCurveTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Path2D_QuadraticCurveTo
//go:noescape
func FuncPath2DQuadraticCurveTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Path2D_QuadraticCurveTo
//go:noescape
func CallPath2DQuadraticCurveTo(
	this js.Ref, retPtr unsafe.Pointer,
	cpx float64,
	cpy float64,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_Path2D_QuadraticCurveTo
//go:noescape
func TryPath2DQuadraticCurveTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cpx float64,
	cpy float64,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Path2D_BezierCurveTo
//go:noescape
func HasFuncPath2DBezierCurveTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Path2D_BezierCurveTo
//go:noescape
func FuncPath2DBezierCurveTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Path2D_BezierCurveTo
//go:noescape
func CallPath2DBezierCurveTo(
	this js.Ref, retPtr unsafe.Pointer,
	cp1x float64,
	cp1y float64,
	cp2x float64,
	cp2y float64,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_Path2D_BezierCurveTo
//go:noescape
func TryPath2DBezierCurveTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cp1x float64,
	cp1y float64,
	cp2x float64,
	cp2y float64,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Path2D_ArcTo
//go:noescape
func HasFuncPath2DArcTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Path2D_ArcTo
//go:noescape
func FuncPath2DArcTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Path2D_ArcTo
//go:noescape
func CallPath2DArcTo(
	this js.Ref, retPtr unsafe.Pointer,
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
	radius float64)

//go:wasmimport plat/js/web try_Path2D_ArcTo
//go:noescape
func TryPath2DArcTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
	radius float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Path2D_Rect
//go:noescape
func HasFuncPath2DRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Path2D_Rect
//go:noescape
func FuncPath2DRect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Path2D_Rect
//go:noescape
func CallPath2DRect(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64)

//go:wasmimport plat/js/web try_Path2D_Rect
//go:noescape
func TryPath2DRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Path2D_RoundRect
//go:noescape
func HasFuncPath2DRoundRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Path2D_RoundRect
//go:noescape
func FuncPath2DRoundRect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Path2D_RoundRect
//go:noescape
func CallPath2DRoundRect(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64,
	radii js.Ref)

//go:wasmimport plat/js/web try_Path2D_RoundRect
//go:noescape
func TryPath2DRoundRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64,
	radii js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Path2D_RoundRect1
//go:noescape
func HasFuncPath2DRoundRect1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Path2D_RoundRect1
//go:noescape
func FuncPath2DRoundRect1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Path2D_RoundRect1
//go:noescape
func CallPath2DRoundRect1(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64)

//go:wasmimport plat/js/web try_Path2D_RoundRect1
//go:noescape
func TryPath2DRoundRect1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Path2D_Arc
//go:noescape
func HasFuncPath2DArc(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Path2D_Arc
//go:noescape
func FuncPath2DArc(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Path2D_Arc
//go:noescape
func CallPath2DArc(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	radius float64,
	startAngle float64,
	endAngle float64,
	counterclockwise js.Ref)

//go:wasmimport plat/js/web try_Path2D_Arc
//go:noescape
func TryPath2DArc(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	radius float64,
	startAngle float64,
	endAngle float64,
	counterclockwise js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Path2D_Arc1
//go:noescape
func HasFuncPath2DArc1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Path2D_Arc1
//go:noescape
func FuncPath2DArc1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Path2D_Arc1
//go:noescape
func CallPath2DArc1(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	radius float64,
	startAngle float64,
	endAngle float64)

//go:wasmimport plat/js/web try_Path2D_Arc1
//go:noescape
func TryPath2DArc1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	radius float64,
	startAngle float64,
	endAngle float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Path2D_Ellipse
//go:noescape
func HasFuncPath2DEllipse(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Path2D_Ellipse
//go:noescape
func FuncPath2DEllipse(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Path2D_Ellipse
//go:noescape
func CallPath2DEllipse(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	radiusX float64,
	radiusY float64,
	rotation float64,
	startAngle float64,
	endAngle float64,
	counterclockwise js.Ref)

//go:wasmimport plat/js/web try_Path2D_Ellipse
//go:noescape
func TryPath2DEllipse(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	radiusX float64,
	radiusY float64,
	rotation float64,
	startAngle float64,
	endAngle float64,
	counterclockwise js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Path2D_Ellipse1
//go:noescape
func HasFuncPath2DEllipse1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Path2D_Ellipse1
//go:noescape
func FuncPath2DEllipse1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Path2D_Ellipse1
//go:noescape
func CallPath2DEllipse1(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	radiusX float64,
	radiusY float64,
	rotation float64,
	startAngle float64,
	endAngle float64)

//go:wasmimport plat/js/web try_Path2D_Ellipse1
//go:noescape
func TryPath2DEllipse1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	radiusX float64,
	radiusY float64,
	rotation float64,
	startAngle float64,
	endAngle float64) (ok js.Ref)

//go:wasmimport plat/js/web get_TextMetrics_Width
//go:noescape
func GetTextMetricsWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextMetrics_ActualBoundingBoxLeft
//go:noescape
func GetTextMetricsActualBoundingBoxLeft(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextMetrics_ActualBoundingBoxRight
//go:noescape
func GetTextMetricsActualBoundingBoxRight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextMetrics_FontBoundingBoxAscent
//go:noescape
func GetTextMetricsFontBoundingBoxAscent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextMetrics_FontBoundingBoxDescent
//go:noescape
func GetTextMetricsFontBoundingBoxDescent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextMetrics_ActualBoundingBoxAscent
//go:noescape
func GetTextMetricsActualBoundingBoxAscent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextMetrics_ActualBoundingBoxDescent
//go:noescape
func GetTextMetricsActualBoundingBoxDescent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextMetrics_EmHeightAscent
//go:noescape
func GetTextMetricsEmHeightAscent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextMetrics_EmHeightDescent
//go:noescape
func GetTextMetricsEmHeightDescent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextMetrics_HangingBaseline
//go:noescape
func GetTextMetricsHangingBaseline(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextMetrics_AlphabeticBaseline
//go:noescape
func GetTextMetricsAlphabeticBaseline(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextMetrics_IdeographicBaseline
//go:noescape
func GetTextMetricsIdeographicBaseline(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ImageDataSettings
//go:noescape
func ImageDataSettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ImageDataSettings
//go:noescape
func ImageDataSettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ImageData_ImageData
//go:noescape
func NewImageDataByImageData(
	sw uint32,
	sh uint32,
	settings unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ImageData_ImageData1
//go:noescape
func NewImageDataByImageData1(
	sw uint32,
	sh uint32) js.Ref

//go:wasmimport plat/js/web new_ImageData_ImageData2
//go:noescape
func NewImageDataByImageData2(
	data js.Ref,
	sw uint32,
	sh uint32,
	settings unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ImageData_ImageData3
//go:noescape
func NewImageDataByImageData3(
	data js.Ref,
	sw uint32,
	sh uint32) js.Ref

//go:wasmimport plat/js/web new_ImageData_ImageData4
//go:noescape
func NewImageDataByImageData4(
	data js.Ref,
	sw uint32) js.Ref

//go:wasmimport plat/js/web get_ImageData_Width
//go:noescape
func GetImageDataWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ImageData_Height
//go:noescape
func GetImageDataHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ImageData_Data
//go:noescape
func GetImageDataData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ImageData_ColorSpace
//go:noescape
func GetImageDataColorSpace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)
