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

//go:wasmimport plat/js/webext/webauthenticationproxy store_CreateRequest
//go:noescape
func CreateRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy load_CreateRequest
//go:noescape
func CreateRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy store_DOMExceptionDetails
//go:noescape
func DOMExceptionDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy load_DOMExceptionDetails
//go:noescape
func DOMExceptionDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy store_CreateResponseDetails
//go:noescape
func CreateResponseDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy load_CreateResponseDetails
//go:noescape
func CreateResponseDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy store_GetRequest
//go:noescape
func GetRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy load_GetRequest
//go:noescape
func GetRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy store_GetResponseDetails
//go:noescape
func GetResponseDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy load_GetResponseDetails
//go:noescape
func GetResponseDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy store_IsUvpaaRequest
//go:noescape
func IsUvpaaRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy load_IsUvpaaRequest
//go:noescape
func IsUvpaaRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy store_IsUvpaaResponseDetails
//go:noescape
func IsUvpaaResponseDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy load_IsUvpaaResponseDetails
//go:noescape
func IsUvpaaResponseDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy has_Attach
//go:noescape
func HasFuncAttach() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_Attach
//go:noescape
func FuncAttach(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_Attach
//go:noescape
func CallAttach(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy try_Attach
//go:noescape
func TryAttach(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_CompleteCreateRequest
//go:noescape
func HasFuncCompleteCreateRequest() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_CompleteCreateRequest
//go:noescape
func FuncCompleteCreateRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_CompleteCreateRequest
//go:noescape
func CallCompleteCreateRequest(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy try_CompleteCreateRequest
//go:noescape
func TryCompleteCreateRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_CompleteGetRequest
//go:noescape
func HasFuncCompleteGetRequest() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_CompleteGetRequest
//go:noescape
func FuncCompleteGetRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_CompleteGetRequest
//go:noescape
func CallCompleteGetRequest(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy try_CompleteGetRequest
//go:noescape
func TryCompleteGetRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_CompleteIsUvpaaRequest
//go:noescape
func HasFuncCompleteIsUvpaaRequest() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_CompleteIsUvpaaRequest
//go:noescape
func FuncCompleteIsUvpaaRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_CompleteIsUvpaaRequest
//go:noescape
func CallCompleteIsUvpaaRequest(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy try_CompleteIsUvpaaRequest
//go:noescape
func TryCompleteIsUvpaaRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_Detach
//go:noescape
func HasFuncDetach() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_Detach
//go:noescape
func FuncDetach(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_Detach
//go:noescape
func CallDetach(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy try_Detach
//go:noescape
func TryDetach(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_OnCreateRequest
//go:noescape
func HasFuncOnCreateRequest() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_OnCreateRequest
//go:noescape
func FuncOnCreateRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_OnCreateRequest
//go:noescape
func CallOnCreateRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy try_OnCreateRequest
//go:noescape
func TryOnCreateRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_OffCreateRequest
//go:noescape
func HasFuncOffCreateRequest() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_OffCreateRequest
//go:noescape
func FuncOffCreateRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_OffCreateRequest
//go:noescape
func CallOffCreateRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy try_OffCreateRequest
//go:noescape
func TryOffCreateRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_HasOnCreateRequest
//go:noescape
func HasFuncHasOnCreateRequest() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_HasOnCreateRequest
//go:noescape
func FuncHasOnCreateRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_HasOnCreateRequest
//go:noescape
func CallHasOnCreateRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy try_HasOnCreateRequest
//go:noescape
func TryHasOnCreateRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_OnGetRequest
//go:noescape
func HasFuncOnGetRequest() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_OnGetRequest
//go:noescape
func FuncOnGetRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_OnGetRequest
//go:noescape
func CallOnGetRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy try_OnGetRequest
//go:noescape
func TryOnGetRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_OffGetRequest
//go:noescape
func HasFuncOffGetRequest() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_OffGetRequest
//go:noescape
func FuncOffGetRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_OffGetRequest
//go:noescape
func CallOffGetRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy try_OffGetRequest
//go:noescape
func TryOffGetRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_HasOnGetRequest
//go:noescape
func HasFuncHasOnGetRequest() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_HasOnGetRequest
//go:noescape
func FuncHasOnGetRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_HasOnGetRequest
//go:noescape
func CallHasOnGetRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy try_HasOnGetRequest
//go:noescape
func TryHasOnGetRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_OnIsUvpaaRequest
//go:noescape
func HasFuncOnIsUvpaaRequest() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_OnIsUvpaaRequest
//go:noescape
func FuncOnIsUvpaaRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_OnIsUvpaaRequest
//go:noescape
func CallOnIsUvpaaRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy try_OnIsUvpaaRequest
//go:noescape
func TryOnIsUvpaaRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_OffIsUvpaaRequest
//go:noescape
func HasFuncOffIsUvpaaRequest() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_OffIsUvpaaRequest
//go:noescape
func FuncOffIsUvpaaRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_OffIsUvpaaRequest
//go:noescape
func CallOffIsUvpaaRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy try_OffIsUvpaaRequest
//go:noescape
func TryOffIsUvpaaRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_HasOnIsUvpaaRequest
//go:noescape
func HasFuncHasOnIsUvpaaRequest() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_HasOnIsUvpaaRequest
//go:noescape
func FuncHasOnIsUvpaaRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_HasOnIsUvpaaRequest
//go:noescape
func CallHasOnIsUvpaaRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy try_HasOnIsUvpaaRequest
//go:noescape
func TryHasOnIsUvpaaRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_OnRemoteSessionStateChange
//go:noescape
func HasFuncOnRemoteSessionStateChange() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_OnRemoteSessionStateChange
//go:noescape
func FuncOnRemoteSessionStateChange(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_OnRemoteSessionStateChange
//go:noescape
func CallOnRemoteSessionStateChange(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy try_OnRemoteSessionStateChange
//go:noescape
func TryOnRemoteSessionStateChange(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_OffRemoteSessionStateChange
//go:noescape
func HasFuncOffRemoteSessionStateChange() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_OffRemoteSessionStateChange
//go:noescape
func FuncOffRemoteSessionStateChange(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_OffRemoteSessionStateChange
//go:noescape
func CallOffRemoteSessionStateChange(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy try_OffRemoteSessionStateChange
//go:noescape
func TryOffRemoteSessionStateChange(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_HasOnRemoteSessionStateChange
//go:noescape
func HasFuncHasOnRemoteSessionStateChange() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_HasOnRemoteSessionStateChange
//go:noescape
func FuncHasOnRemoteSessionStateChange(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_HasOnRemoteSessionStateChange
//go:noescape
func CallHasOnRemoteSessionStateChange(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy try_HasOnRemoteSessionStateChange
//go:noescape
func TryHasOnRemoteSessionStateChange(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_OnRequestCanceled
//go:noescape
func HasFuncOnRequestCanceled() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_OnRequestCanceled
//go:noescape
func FuncOnRequestCanceled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_OnRequestCanceled
//go:noescape
func CallOnRequestCanceled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy try_OnRequestCanceled
//go:noescape
func TryOnRequestCanceled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_OffRequestCanceled
//go:noescape
func HasFuncOffRequestCanceled() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_OffRequestCanceled
//go:noescape
func FuncOffRequestCanceled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_OffRequestCanceled
//go:noescape
func CallOffRequestCanceled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy try_OffRequestCanceled
//go:noescape
func TryOffRequestCanceled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy has_HasOnRequestCanceled
//go:noescape
func HasFuncHasOnRequestCanceled() js.Ref

//go:wasmimport plat/js/webext/webauthenticationproxy func_HasOnRequestCanceled
//go:noescape
func FuncHasOnRequestCanceled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webauthenticationproxy call_HasOnRequestCanceled
//go:noescape
func CallHasOnRequestCanceled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webauthenticationproxy try_HasOnRequestCanceled
//go:noescape
func TryHasOnRequestCanceled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
