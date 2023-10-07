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

//go:wasmimport plat/js/webext/extension constof_ViewType
//go:noescape
func ConstOfViewType(str js.Ref) uint32

//go:wasmimport plat/js/webext/extension store_GetViewsArgFetchProperties
//go:noescape
func GetViewsArgFetchPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extension load_GetViewsArgFetchProperties
//go:noescape
func GetViewsArgFetchPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extension store_LastErrorProperty
//go:noescape
func LastErrorPropertyJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extension load_LastErrorProperty
//go:noescape
func LastErrorPropertyJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extension has_GetBackgroundPage
//go:noescape
func HasFuncGetBackgroundPage() js.Ref

//go:wasmimport plat/js/webext/extension func_GetBackgroundPage
//go:noescape
func FuncGetBackgroundPage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extension call_GetBackgroundPage
//go:noescape
func CallGetBackgroundPage(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/extension try_GetBackgroundPage
//go:noescape
func TryGetBackgroundPage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/extension has_GetExtensionTabs
//go:noescape
func HasFuncGetExtensionTabs() js.Ref

//go:wasmimport plat/js/webext/extension func_GetExtensionTabs
//go:noescape
func FuncGetExtensionTabs(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extension call_GetExtensionTabs
//go:noescape
func CallGetExtensionTabs(
	retPtr unsafe.Pointer,
	windowId float64)

//go:wasmimport plat/js/webext/extension try_GetExtensionTabs
//go:noescape
func TryGetExtensionTabs(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	windowId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/extension has_GetURL
//go:noescape
func HasFuncGetURL() js.Ref

//go:wasmimport plat/js/webext/extension func_GetURL
//go:noescape
func FuncGetURL(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extension call_GetURL
//go:noescape
func CallGetURL(
	retPtr unsafe.Pointer,
	path js.Ref)

//go:wasmimport plat/js/webext/extension try_GetURL
//go:noescape
func TryGetURL(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extension has_GetViews
//go:noescape
func HasFuncGetViews() js.Ref

//go:wasmimport plat/js/webext/extension func_GetViews
//go:noescape
func FuncGetViews(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extension call_GetViews
//go:noescape
func CallGetViews(
	retPtr unsafe.Pointer,
	fetchProperties unsafe.Pointer)

//go:wasmimport plat/js/webext/extension try_GetViews
//go:noescape
func TryGetViews(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fetchProperties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/extension get_InIncognitoContext
//go:noescape
func GetInIncognitoContext(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/extension set_InIncognitoContext
//go:noescape
func SetInIncognitoContext(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/extension has_IsAllowedFileSchemeAccess
//go:noescape
func HasFuncIsAllowedFileSchemeAccess() js.Ref

//go:wasmimport plat/js/webext/extension func_IsAllowedFileSchemeAccess
//go:noescape
func FuncIsAllowedFileSchemeAccess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extension call_IsAllowedFileSchemeAccess
//go:noescape
func CallIsAllowedFileSchemeAccess(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/extension try_IsAllowedFileSchemeAccess
//go:noescape
func TryIsAllowedFileSchemeAccess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/extension has_IsAllowedIncognitoAccess
//go:noescape
func HasFuncIsAllowedIncognitoAccess() js.Ref

//go:wasmimport plat/js/webext/extension func_IsAllowedIncognitoAccess
//go:noescape
func FuncIsAllowedIncognitoAccess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extension call_IsAllowedIncognitoAccess
//go:noescape
func CallIsAllowedIncognitoAccess(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/extension try_IsAllowedIncognitoAccess
//go:noescape
func TryIsAllowedIncognitoAccess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/extension get_LastError
//go:noescape
func GetLastError(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/extension set_LastError
//go:noescape
func SetLastError(
	val unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/extension has_OnRequest
//go:noescape
func HasFuncOnRequest() js.Ref

//go:wasmimport plat/js/webext/extension func_OnRequest
//go:noescape
func FuncOnRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extension call_OnRequest
//go:noescape
func CallOnRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/extension try_OnRequest
//go:noescape
func TryOnRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extension has_OffRequest
//go:noescape
func HasFuncOffRequest() js.Ref

//go:wasmimport plat/js/webext/extension func_OffRequest
//go:noescape
func FuncOffRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extension call_OffRequest
//go:noescape
func CallOffRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/extension try_OffRequest
//go:noescape
func TryOffRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extension has_HasOnRequest
//go:noescape
func HasFuncHasOnRequest() js.Ref

//go:wasmimport plat/js/webext/extension func_HasOnRequest
//go:noescape
func FuncHasOnRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extension call_HasOnRequest
//go:noescape
func CallHasOnRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/extension try_HasOnRequest
//go:noescape
func TryHasOnRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extension has_OnRequestExternal
//go:noescape
func HasFuncOnRequestExternal() js.Ref

//go:wasmimport plat/js/webext/extension func_OnRequestExternal
//go:noescape
func FuncOnRequestExternal(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extension call_OnRequestExternal
//go:noescape
func CallOnRequestExternal(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/extension try_OnRequestExternal
//go:noescape
func TryOnRequestExternal(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extension has_OffRequestExternal
//go:noescape
func HasFuncOffRequestExternal() js.Ref

//go:wasmimport plat/js/webext/extension func_OffRequestExternal
//go:noescape
func FuncOffRequestExternal(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extension call_OffRequestExternal
//go:noescape
func CallOffRequestExternal(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/extension try_OffRequestExternal
//go:noescape
func TryOffRequestExternal(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extension has_HasOnRequestExternal
//go:noescape
func HasFuncHasOnRequestExternal() js.Ref

//go:wasmimport plat/js/webext/extension func_HasOnRequestExternal
//go:noescape
func FuncHasOnRequestExternal(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extension call_HasOnRequestExternal
//go:noescape
func CallHasOnRequestExternal(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/extension try_HasOnRequestExternal
//go:noescape
func TryHasOnRequestExternal(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extension has_SendRequest
//go:noescape
func HasFuncSendRequest() js.Ref

//go:wasmimport plat/js/webext/extension func_SendRequest
//go:noescape
func FuncSendRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extension call_SendRequest
//go:noescape
func CallSendRequest(
	retPtr unsafe.Pointer,
	extensionId js.Ref,
	request js.Ref)

//go:wasmimport plat/js/webext/extension try_SendRequest
//go:noescape
func TrySendRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	extensionId js.Ref,
	request js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extension has_SetUpdateUrlData
//go:noescape
func HasFuncSetUpdateUrlData() js.Ref

//go:wasmimport plat/js/webext/extension func_SetUpdateUrlData
//go:noescape
func FuncSetUpdateUrlData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extension call_SetUpdateUrlData
//go:noescape
func CallSetUpdateUrlData(
	retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/webext/extension try_SetUpdateUrlData
//go:noescape
func TrySetUpdateUrlData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)
