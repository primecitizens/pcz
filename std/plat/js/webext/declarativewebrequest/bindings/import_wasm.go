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

//go:wasmimport plat/js/webext/declarativewebrequest store_RequestCookie
//go:noescape
func RequestCookieJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_RequestCookie
//go:noescape
func RequestCookieJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_AddRequestCookieInstanceType
//go:noescape
func ConstOfAddRequestCookieInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_AddRequestCookie
//go:noescape
func AddRequestCookieJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_AddRequestCookie
//go:noescape
func AddRequestCookieJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest store_ResponseCookie
//go:noescape
func ResponseCookieJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_ResponseCookie
//go:noescape
func ResponseCookieJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_AddResponseCookieInstanceType
//go:noescape
func ConstOfAddResponseCookieInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_AddResponseCookie
//go:noescape
func AddResponseCookieJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_AddResponseCookie
//go:noescape
func AddResponseCookieJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_AddResponseHeaderInstanceType
//go:noescape
func ConstOfAddResponseHeaderInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_AddResponseHeader
//go:noescape
func AddResponseHeaderJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_AddResponseHeader
//go:noescape
func AddResponseHeaderJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_CancelRequestInstanceType
//go:noescape
func ConstOfCancelRequestInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_CancelRequest
//go:noescape
func CancelRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_CancelRequest
//go:noescape
func CancelRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_EditRequestCookieInstanceType
//go:noescape
func ConstOfEditRequestCookieInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_EditRequestCookie
//go:noescape
func EditRequestCookieJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_EditRequestCookie
//go:noescape
func EditRequestCookieJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest store_FilterResponseCookie
//go:noescape
func FilterResponseCookieJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_FilterResponseCookie
//go:noescape
func FilterResponseCookieJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_EditResponseCookieInstanceType
//go:noescape
func ConstOfEditResponseCookieInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_EditResponseCookie
//go:noescape
func EditResponseCookieJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_EditResponseCookie
//go:noescape
func EditResponseCookieJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest store_HeaderFilter
//go:noescape
func HeaderFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_HeaderFilter
//go:noescape
func HeaderFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_IgnoreRulesInstanceType
//go:noescape
func ConstOfIgnoreRulesInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_IgnoreRules
//go:noescape
func IgnoreRulesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_IgnoreRules
//go:noescape
func IgnoreRulesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_Stage
//go:noescape
func ConstOfStage(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_OnMessageArgDetails
//go:noescape
func OnMessageArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_OnMessageArgDetails
//go:noescape
func OnMessageArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_RedirectByRegExInstanceType
//go:noescape
func ConstOfRedirectByRegExInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_RedirectByRegEx
//go:noescape
func RedirectByRegExJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_RedirectByRegEx
//go:noescape
func RedirectByRegExJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_RedirectRequestInstanceType
//go:noescape
func ConstOfRedirectRequestInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_RedirectRequest
//go:noescape
func RedirectRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_RedirectRequest
//go:noescape
func RedirectRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_RedirectToEmptyDocumentInstanceType
//go:noescape
func ConstOfRedirectToEmptyDocumentInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_RedirectToEmptyDocument
//go:noescape
func RedirectToEmptyDocumentJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_RedirectToEmptyDocument
//go:noescape
func RedirectToEmptyDocumentJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_RedirectToTransparentImageInstanceType
//go:noescape
func ConstOfRedirectToTransparentImageInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_RedirectToTransparentImage
//go:noescape
func RedirectToTransparentImageJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_RedirectToTransparentImage
//go:noescape
func RedirectToTransparentImageJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_RemoveRequestCookieInstanceType
//go:noescape
func ConstOfRemoveRequestCookieInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_RemoveRequestCookie
//go:noescape
func RemoveRequestCookieJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_RemoveRequestCookie
//go:noescape
func RemoveRequestCookieJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_RemoveRequestHeaderInstanceType
//go:noescape
func ConstOfRemoveRequestHeaderInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_RemoveRequestHeader
//go:noescape
func RemoveRequestHeaderJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_RemoveRequestHeader
//go:noescape
func RemoveRequestHeaderJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_RemoveResponseCookieInstanceType
//go:noescape
func ConstOfRemoveResponseCookieInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_RemoveResponseCookie
//go:noescape
func RemoveResponseCookieJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_RemoveResponseCookie
//go:noescape
func RemoveResponseCookieJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_RemoveResponseHeaderInstanceType
//go:noescape
func ConstOfRemoveResponseHeaderInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_RemoveResponseHeader
//go:noescape
func RemoveResponseHeaderJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_RemoveResponseHeader
//go:noescape
func RemoveResponseHeaderJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_RequestMatcherInstanceType
//go:noescape
func ConstOfRequestMatcherInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_RequestMatcher
//go:noescape
func RequestMatcherJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_RequestMatcher
//go:noescape
func RequestMatcherJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_SendMessageToExtensionInstanceType
//go:noescape
func ConstOfSendMessageToExtensionInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_SendMessageToExtension
//go:noescape
func SendMessageToExtensionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_SendMessageToExtension
//go:noescape
func SendMessageToExtensionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest constof_SetRequestHeaderInstanceType
//go:noescape
func ConstOfSetRequestHeaderInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativewebrequest store_SetRequestHeader
//go:noescape
func SetRequestHeaderJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest load_SetRequestHeader
//go:noescape
func SetRequestHeaderJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest has_OnMessage
//go:noescape
func HasFuncOnMessage() js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest func_OnMessage
//go:noescape
func FuncOnMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativewebrequest call_OnMessage
//go:noescape
func CallOnMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest try_OnMessage
//go:noescape
func TryOnMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest has_OffMessage
//go:noescape
func HasFuncOffMessage() js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest func_OffMessage
//go:noescape
func FuncOffMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativewebrequest call_OffMessage
//go:noescape
func CallOffMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest try_OffMessage
//go:noescape
func TryOffMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest has_HasOnMessage
//go:noescape
func HasFuncHasOnMessage() js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest func_HasOnMessage
//go:noescape
func FuncHasOnMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativewebrequest call_HasOnMessage
//go:noescape
func CallHasOnMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest try_HasOnMessage
//go:noescape
func TryHasOnMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest has_OnRequest
//go:noescape
func HasFuncOnRequest() js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest func_OnRequest
//go:noescape
func FuncOnRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativewebrequest call_OnRequest
//go:noescape
func CallOnRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest try_OnRequest
//go:noescape
func TryOnRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest has_OffRequest
//go:noescape
func HasFuncOffRequest() js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest func_OffRequest
//go:noescape
func FuncOffRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativewebrequest call_OffRequest
//go:noescape
func CallOffRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest try_OffRequest
//go:noescape
func TryOffRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest has_HasOnRequest
//go:noescape
func HasFuncHasOnRequest() js.Ref

//go:wasmimport plat/js/webext/declarativewebrequest func_HasOnRequest
//go:noescape
func FuncHasOnRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativewebrequest call_HasOnRequest
//go:noescape
func CallHasOnRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/declarativewebrequest try_HasOnRequest
//go:noescape
func TryHasOnRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
