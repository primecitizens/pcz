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

//go:wasmimport plat/js/webext/app/runtime constof_ActionType
//go:noescape
func ConstOfActionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/app/runtime store_ActionData
//go:noescape
func ActionDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/app/runtime load_ActionData
//go:noescape
func ActionDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/runtime get_EmbedRequest_EmbedderId
//go:noescape
func GetEmbedRequestEmbedderId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/runtime set_EmbedRequest_EmbedderId
//go:noescape
func SetEmbedRequestEmbedderId(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/app/runtime get_EmbedRequest_Data
//go:noescape
func GetEmbedRequestData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/runtime set_EmbedRequest_Data
//go:noescape
func SetEmbedRequestData(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/app/runtime has_EmbedRequest_Allow
//go:noescape
func HasFuncEmbedRequestAllow(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/runtime func_EmbedRequest_Allow
//go:noescape
func FuncEmbedRequestAllow(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/runtime call_EmbedRequest_Allow
//go:noescape
func CallEmbedRequestAllow(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/app/runtime try_EmbedRequest_Allow
//go:noescape
func TryEmbedRequestAllow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/runtime has_EmbedRequest_Deny
//go:noescape
func HasFuncEmbedRequestDeny(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/runtime func_EmbedRequest_Deny
//go:noescape
func FuncEmbedRequestDeny(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/runtime call_EmbedRequest_Deny
//go:noescape
func CallEmbedRequestDeny(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/runtime try_EmbedRequest_Deny
//go:noescape
func TryEmbedRequestDeny(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/runtime store_LaunchItem
//go:noescape
func LaunchItemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/app/runtime load_LaunchItem
//go:noescape
func LaunchItemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/runtime constof_LaunchSource
//go:noescape
func ConstOfLaunchSource(str js.Ref) uint32

//go:wasmimport plat/js/webext/app/runtime store_LaunchData
//go:noescape
func LaunchDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/app/runtime load_LaunchData
//go:noescape
func LaunchDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/runtime has_OnEmbedRequested
//go:noescape
func HasFuncOnEmbedRequested() js.Ref

//go:wasmimport plat/js/webext/app/runtime func_OnEmbedRequested
//go:noescape
func FuncOnEmbedRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/runtime call_OnEmbedRequested
//go:noescape
func CallOnEmbedRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/runtime try_OnEmbedRequested
//go:noescape
func TryOnEmbedRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/runtime has_OffEmbedRequested
//go:noescape
func HasFuncOffEmbedRequested() js.Ref

//go:wasmimport plat/js/webext/app/runtime func_OffEmbedRequested
//go:noescape
func FuncOffEmbedRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/runtime call_OffEmbedRequested
//go:noescape
func CallOffEmbedRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/runtime try_OffEmbedRequested
//go:noescape
func TryOffEmbedRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/runtime has_HasOnEmbedRequested
//go:noescape
func HasFuncHasOnEmbedRequested() js.Ref

//go:wasmimport plat/js/webext/app/runtime func_HasOnEmbedRequested
//go:noescape
func FuncHasOnEmbedRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/runtime call_HasOnEmbedRequested
//go:noescape
func CallHasOnEmbedRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/runtime try_HasOnEmbedRequested
//go:noescape
func TryHasOnEmbedRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/runtime has_OnLaunched
//go:noescape
func HasFuncOnLaunched() js.Ref

//go:wasmimport plat/js/webext/app/runtime func_OnLaunched
//go:noescape
func FuncOnLaunched(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/runtime call_OnLaunched
//go:noescape
func CallOnLaunched(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/runtime try_OnLaunched
//go:noescape
func TryOnLaunched(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/runtime has_OffLaunched
//go:noescape
func HasFuncOffLaunched() js.Ref

//go:wasmimport plat/js/webext/app/runtime func_OffLaunched
//go:noescape
func FuncOffLaunched(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/runtime call_OffLaunched
//go:noescape
func CallOffLaunched(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/runtime try_OffLaunched
//go:noescape
func TryOffLaunched(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/runtime has_HasOnLaunched
//go:noescape
func HasFuncHasOnLaunched() js.Ref

//go:wasmimport plat/js/webext/app/runtime func_HasOnLaunched
//go:noescape
func FuncHasOnLaunched(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/runtime call_HasOnLaunched
//go:noescape
func CallHasOnLaunched(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/runtime try_HasOnLaunched
//go:noescape
func TryHasOnLaunched(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/runtime has_OnRestarted
//go:noescape
func HasFuncOnRestarted() js.Ref

//go:wasmimport plat/js/webext/app/runtime func_OnRestarted
//go:noescape
func FuncOnRestarted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/runtime call_OnRestarted
//go:noescape
func CallOnRestarted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/runtime try_OnRestarted
//go:noescape
func TryOnRestarted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/runtime has_OffRestarted
//go:noescape
func HasFuncOffRestarted() js.Ref

//go:wasmimport plat/js/webext/app/runtime func_OffRestarted
//go:noescape
func FuncOffRestarted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/runtime call_OffRestarted
//go:noescape
func CallOffRestarted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/runtime try_OffRestarted
//go:noescape
func TryOffRestarted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/runtime has_HasOnRestarted
//go:noescape
func HasFuncHasOnRestarted() js.Ref

//go:wasmimport plat/js/webext/app/runtime func_HasOnRestarted
//go:noescape
func FuncHasOnRestarted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/runtime call_HasOnRestarted
//go:noescape
func CallHasOnRestarted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/runtime try_HasOnRestarted
//go:noescape
func TryHasOnRestarted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
