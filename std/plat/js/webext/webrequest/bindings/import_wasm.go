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

//go:wasmimport plat/js/webext/webrequest store_BlockingResponseFieldAuthCredentials
//go:noescape
func BlockingResponseFieldAuthCredentialsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_BlockingResponseFieldAuthCredentials
//go:noescape
func BlockingResponseFieldAuthCredentialsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest store_HttpHeadersElem
//go:noescape
func HttpHeadersElemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_HttpHeadersElem
//go:noescape
func HttpHeadersElemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest store_BlockingResponse
//go:noescape
func BlockingResponseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_BlockingResponse
//go:noescape
func BlockingResponseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest constof_IgnoredActionType
//go:noescape
func ConstOfIgnoredActionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/webrequest get_MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES
//go:noescape
func GetMAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/webrequest set_MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES
//go:noescape
func SetMAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest store_OnActionIgnoredArgDetails
//go:noescape
func OnActionIgnoredArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_OnActionIgnoredArgDetails
//go:noescape
func OnActionIgnoredArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest store_OnAuthRequiredArgDetailsFieldChallenger
//go:noescape
func OnAuthRequiredArgDetailsFieldChallengerJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_OnAuthRequiredArgDetailsFieldChallenger
//go:noescape
func OnAuthRequiredArgDetailsFieldChallengerJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest constof_ResourceType
//go:noescape
func ConstOfResourceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/webrequest store_OnAuthRequiredArgDetails
//go:noescape
func OnAuthRequiredArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_OnAuthRequiredArgDetails
//go:noescape
func OnAuthRequiredArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest constof_OnAuthRequiredOptions
//go:noescape
func ConstOfOnAuthRequiredOptions(str js.Ref) uint32

//go:wasmimport plat/js/webext/webrequest store_OnBeforeRedirectArgDetails
//go:noescape
func OnBeforeRedirectArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_OnBeforeRedirectArgDetails
//go:noescape
func OnBeforeRedirectArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest constof_OnBeforeRedirectOptions
//go:noescape
func ConstOfOnBeforeRedirectOptions(str js.Ref) uint32

//go:wasmimport plat/js/webext/webrequest store_UploadData
//go:noescape
func UploadDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_UploadData
//go:noescape
func UploadDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest store_OnBeforeRequestArgDetailsFieldRequestBody
//go:noescape
func OnBeforeRequestArgDetailsFieldRequestBodyJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_OnBeforeRequestArgDetailsFieldRequestBody
//go:noescape
func OnBeforeRequestArgDetailsFieldRequestBodyJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest store_OnBeforeRequestArgDetails
//go:noescape
func OnBeforeRequestArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_OnBeforeRequestArgDetails
//go:noescape
func OnBeforeRequestArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest constof_OnBeforeRequestOptions
//go:noescape
func ConstOfOnBeforeRequestOptions(str js.Ref) uint32

//go:wasmimport plat/js/webext/webrequest store_OnBeforeSendHeadersArgDetails
//go:noescape
func OnBeforeSendHeadersArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_OnBeforeSendHeadersArgDetails
//go:noescape
func OnBeforeSendHeadersArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest constof_OnBeforeSendHeadersOptions
//go:noescape
func ConstOfOnBeforeSendHeadersOptions(str js.Ref) uint32

//go:wasmimport plat/js/webext/webrequest store_OnCompletedArgDetails
//go:noescape
func OnCompletedArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_OnCompletedArgDetails
//go:noescape
func OnCompletedArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest constof_OnCompletedOptions
//go:noescape
func ConstOfOnCompletedOptions(str js.Ref) uint32

//go:wasmimport plat/js/webext/webrequest store_OnErrorOccurredArgDetails
//go:noescape
func OnErrorOccurredArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_OnErrorOccurredArgDetails
//go:noescape
func OnErrorOccurredArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest constof_OnErrorOccurredOptions
//go:noescape
func ConstOfOnErrorOccurredOptions(str js.Ref) uint32

//go:wasmimport plat/js/webext/webrequest store_OnHeadersReceivedArgDetails
//go:noescape
func OnHeadersReceivedArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_OnHeadersReceivedArgDetails
//go:noescape
func OnHeadersReceivedArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest constof_OnHeadersReceivedOptions
//go:noescape
func ConstOfOnHeadersReceivedOptions(str js.Ref) uint32

//go:wasmimport plat/js/webext/webrequest store_OnResponseStartedArgDetails
//go:noescape
func OnResponseStartedArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_OnResponseStartedArgDetails
//go:noescape
func OnResponseStartedArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest constof_OnResponseStartedOptions
//go:noescape
func ConstOfOnResponseStartedOptions(str js.Ref) uint32

//go:wasmimport plat/js/webext/webrequest store_OnSendHeadersArgDetails
//go:noescape
func OnSendHeadersArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_OnSendHeadersArgDetails
//go:noescape
func OnSendHeadersArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest constof_OnSendHeadersOptions
//go:noescape
func ConstOfOnSendHeadersOptions(str js.Ref) uint32

//go:wasmimport plat/js/webext/webrequest store_RequestFilter
//go:noescape
func RequestFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrequest load_RequestFilter
//go:noescape
func RequestFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrequest has_HandlerBehaviorChanged
//go:noescape
func HasFuncHandlerBehaviorChanged() js.Ref

//go:wasmimport plat/js/webext/webrequest func_HandlerBehaviorChanged
//go:noescape
func FuncHandlerBehaviorChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_HandlerBehaviorChanged
//go:noescape
func CallHandlerBehaviorChanged(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest try_HandlerBehaviorChanged
//go:noescape
func TryHandlerBehaviorChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OnActionIgnored
//go:noescape
func HasFuncOnActionIgnored() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OnActionIgnored
//go:noescape
func FuncOnActionIgnored(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OnActionIgnored
//go:noescape
func CallOnActionIgnored(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OnActionIgnored
//go:noescape
func TryOnActionIgnored(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OffActionIgnored
//go:noescape
func HasFuncOffActionIgnored() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OffActionIgnored
//go:noescape
func FuncOffActionIgnored(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OffActionIgnored
//go:noescape
func CallOffActionIgnored(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OffActionIgnored
//go:noescape
func TryOffActionIgnored(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_HasOnActionIgnored
//go:noescape
func HasFuncHasOnActionIgnored() js.Ref

//go:wasmimport plat/js/webext/webrequest func_HasOnActionIgnored
//go:noescape
func FuncHasOnActionIgnored(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_HasOnActionIgnored
//go:noescape
func CallHasOnActionIgnored(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_HasOnActionIgnored
//go:noescape
func TryHasOnActionIgnored(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OnAuthRequired
//go:noescape
func HasFuncOnAuthRequired() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OnAuthRequired
//go:noescape
func FuncOnAuthRequired(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OnAuthRequired
//go:noescape
func CallOnAuthRequired(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OnAuthRequired
//go:noescape
func TryOnAuthRequired(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OffAuthRequired
//go:noescape
func HasFuncOffAuthRequired() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OffAuthRequired
//go:noescape
func FuncOffAuthRequired(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OffAuthRequired
//go:noescape
func CallOffAuthRequired(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OffAuthRequired
//go:noescape
func TryOffAuthRequired(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_HasOnAuthRequired
//go:noescape
func HasFuncHasOnAuthRequired() js.Ref

//go:wasmimport plat/js/webext/webrequest func_HasOnAuthRequired
//go:noescape
func FuncHasOnAuthRequired(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_HasOnAuthRequired
//go:noescape
func CallHasOnAuthRequired(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_HasOnAuthRequired
//go:noescape
func TryHasOnAuthRequired(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OnBeforeRedirect
//go:noescape
func HasFuncOnBeforeRedirect() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OnBeforeRedirect
//go:noescape
func FuncOnBeforeRedirect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OnBeforeRedirect
//go:noescape
func CallOnBeforeRedirect(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OnBeforeRedirect
//go:noescape
func TryOnBeforeRedirect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OffBeforeRedirect
//go:noescape
func HasFuncOffBeforeRedirect() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OffBeforeRedirect
//go:noescape
func FuncOffBeforeRedirect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OffBeforeRedirect
//go:noescape
func CallOffBeforeRedirect(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OffBeforeRedirect
//go:noescape
func TryOffBeforeRedirect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_HasOnBeforeRedirect
//go:noescape
func HasFuncHasOnBeforeRedirect() js.Ref

//go:wasmimport plat/js/webext/webrequest func_HasOnBeforeRedirect
//go:noescape
func FuncHasOnBeforeRedirect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_HasOnBeforeRedirect
//go:noescape
func CallHasOnBeforeRedirect(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_HasOnBeforeRedirect
//go:noescape
func TryHasOnBeforeRedirect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OnBeforeRequest
//go:noescape
func HasFuncOnBeforeRequest() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OnBeforeRequest
//go:noescape
func FuncOnBeforeRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OnBeforeRequest
//go:noescape
func CallOnBeforeRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OnBeforeRequest
//go:noescape
func TryOnBeforeRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OffBeforeRequest
//go:noescape
func HasFuncOffBeforeRequest() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OffBeforeRequest
//go:noescape
func FuncOffBeforeRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OffBeforeRequest
//go:noescape
func CallOffBeforeRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OffBeforeRequest
//go:noescape
func TryOffBeforeRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_HasOnBeforeRequest
//go:noescape
func HasFuncHasOnBeforeRequest() js.Ref

//go:wasmimport plat/js/webext/webrequest func_HasOnBeforeRequest
//go:noescape
func FuncHasOnBeforeRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_HasOnBeforeRequest
//go:noescape
func CallHasOnBeforeRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_HasOnBeforeRequest
//go:noescape
func TryHasOnBeforeRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OnBeforeSendHeaders
//go:noescape
func HasFuncOnBeforeSendHeaders() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OnBeforeSendHeaders
//go:noescape
func FuncOnBeforeSendHeaders(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OnBeforeSendHeaders
//go:noescape
func CallOnBeforeSendHeaders(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OnBeforeSendHeaders
//go:noescape
func TryOnBeforeSendHeaders(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OffBeforeSendHeaders
//go:noescape
func HasFuncOffBeforeSendHeaders() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OffBeforeSendHeaders
//go:noescape
func FuncOffBeforeSendHeaders(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OffBeforeSendHeaders
//go:noescape
func CallOffBeforeSendHeaders(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OffBeforeSendHeaders
//go:noescape
func TryOffBeforeSendHeaders(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_HasOnBeforeSendHeaders
//go:noescape
func HasFuncHasOnBeforeSendHeaders() js.Ref

//go:wasmimport plat/js/webext/webrequest func_HasOnBeforeSendHeaders
//go:noescape
func FuncHasOnBeforeSendHeaders(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_HasOnBeforeSendHeaders
//go:noescape
func CallHasOnBeforeSendHeaders(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_HasOnBeforeSendHeaders
//go:noescape
func TryHasOnBeforeSendHeaders(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OnCompleted
//go:noescape
func HasFuncOnCompleted() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OnCompleted
//go:noescape
func FuncOnCompleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OnCompleted
//go:noescape
func CallOnCompleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OnCompleted
//go:noescape
func TryOnCompleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OffCompleted
//go:noescape
func HasFuncOffCompleted() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OffCompleted
//go:noescape
func FuncOffCompleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OffCompleted
//go:noescape
func CallOffCompleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OffCompleted
//go:noescape
func TryOffCompleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_HasOnCompleted
//go:noescape
func HasFuncHasOnCompleted() js.Ref

//go:wasmimport plat/js/webext/webrequest func_HasOnCompleted
//go:noescape
func FuncHasOnCompleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_HasOnCompleted
//go:noescape
func CallHasOnCompleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_HasOnCompleted
//go:noescape
func TryHasOnCompleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OnErrorOccurred
//go:noescape
func HasFuncOnErrorOccurred() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OnErrorOccurred
//go:noescape
func FuncOnErrorOccurred(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OnErrorOccurred
//go:noescape
func CallOnErrorOccurred(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OnErrorOccurred
//go:noescape
func TryOnErrorOccurred(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OffErrorOccurred
//go:noescape
func HasFuncOffErrorOccurred() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OffErrorOccurred
//go:noescape
func FuncOffErrorOccurred(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OffErrorOccurred
//go:noescape
func CallOffErrorOccurred(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OffErrorOccurred
//go:noescape
func TryOffErrorOccurred(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_HasOnErrorOccurred
//go:noescape
func HasFuncHasOnErrorOccurred() js.Ref

//go:wasmimport plat/js/webext/webrequest func_HasOnErrorOccurred
//go:noescape
func FuncHasOnErrorOccurred(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_HasOnErrorOccurred
//go:noescape
func CallHasOnErrorOccurred(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_HasOnErrorOccurred
//go:noescape
func TryHasOnErrorOccurred(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OnHeadersReceived
//go:noescape
func HasFuncOnHeadersReceived() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OnHeadersReceived
//go:noescape
func FuncOnHeadersReceived(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OnHeadersReceived
//go:noescape
func CallOnHeadersReceived(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OnHeadersReceived
//go:noescape
func TryOnHeadersReceived(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OffHeadersReceived
//go:noescape
func HasFuncOffHeadersReceived() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OffHeadersReceived
//go:noescape
func FuncOffHeadersReceived(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OffHeadersReceived
//go:noescape
func CallOffHeadersReceived(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OffHeadersReceived
//go:noescape
func TryOffHeadersReceived(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_HasOnHeadersReceived
//go:noescape
func HasFuncHasOnHeadersReceived() js.Ref

//go:wasmimport plat/js/webext/webrequest func_HasOnHeadersReceived
//go:noescape
func FuncHasOnHeadersReceived(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_HasOnHeadersReceived
//go:noescape
func CallHasOnHeadersReceived(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_HasOnHeadersReceived
//go:noescape
func TryHasOnHeadersReceived(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OnResponseStarted
//go:noescape
func HasFuncOnResponseStarted() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OnResponseStarted
//go:noescape
func FuncOnResponseStarted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OnResponseStarted
//go:noescape
func CallOnResponseStarted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OnResponseStarted
//go:noescape
func TryOnResponseStarted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OffResponseStarted
//go:noescape
func HasFuncOffResponseStarted() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OffResponseStarted
//go:noescape
func FuncOffResponseStarted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OffResponseStarted
//go:noescape
func CallOffResponseStarted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OffResponseStarted
//go:noescape
func TryOffResponseStarted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_HasOnResponseStarted
//go:noescape
func HasFuncHasOnResponseStarted() js.Ref

//go:wasmimport plat/js/webext/webrequest func_HasOnResponseStarted
//go:noescape
func FuncHasOnResponseStarted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_HasOnResponseStarted
//go:noescape
func CallHasOnResponseStarted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_HasOnResponseStarted
//go:noescape
func TryHasOnResponseStarted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OnSendHeaders
//go:noescape
func HasFuncOnSendHeaders() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OnSendHeaders
//go:noescape
func FuncOnSendHeaders(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OnSendHeaders
//go:noescape
func CallOnSendHeaders(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OnSendHeaders
//go:noescape
func TryOnSendHeaders(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_OffSendHeaders
//go:noescape
func HasFuncOffSendHeaders() js.Ref

//go:wasmimport plat/js/webext/webrequest func_OffSendHeaders
//go:noescape
func FuncOffSendHeaders(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_OffSendHeaders
//go:noescape
func CallOffSendHeaders(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_OffSendHeaders
//go:noescape
func TryOffSendHeaders(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequest has_HasOnSendHeaders
//go:noescape
func HasFuncHasOnSendHeaders() js.Ref

//go:wasmimport plat/js/webext/webrequest func_HasOnSendHeaders
//go:noescape
func FuncHasOnSendHeaders(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequest call_HasOnSendHeaders
//go:noescape
func CallHasOnSendHeaders(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrequest try_HasOnSendHeaders
//go:noescape
func TryHasOnSendHeaders(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
