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

//go:wasmimport plat/js/webext/instanceid store_DeleteTokenArgDeleteTokenParams
//go:noescape
func DeleteTokenArgDeleteTokenParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/instanceid load_DeleteTokenArgDeleteTokenParams
//go:noescape
func DeleteTokenArgDeleteTokenParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/instanceid store_GetTokenArgGetTokenParams
//go:noescape
func GetTokenArgGetTokenParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/instanceid load_GetTokenArgGetTokenParams
//go:noescape
func GetTokenArgGetTokenParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/instanceid has_DeleteID
//go:noescape
func HasFuncDeleteID() js.Ref

//go:wasmimport plat/js/webext/instanceid func_DeleteID
//go:noescape
func FuncDeleteID(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/instanceid call_DeleteID
//go:noescape
func CallDeleteID(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/instanceid try_DeleteID
//go:noescape
func TryDeleteID(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/instanceid has_DeleteToken
//go:noescape
func HasFuncDeleteToken() js.Ref

//go:wasmimport plat/js/webext/instanceid func_DeleteToken
//go:noescape
func FuncDeleteToken(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/instanceid call_DeleteToken
//go:noescape
func CallDeleteToken(
	retPtr unsafe.Pointer,
	deleteTokenParams unsafe.Pointer)

//go:wasmimport plat/js/webext/instanceid try_DeleteToken
//go:noescape
func TryDeleteToken(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deleteTokenParams unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/instanceid has_GetCreationTime
//go:noescape
func HasFuncGetCreationTime() js.Ref

//go:wasmimport plat/js/webext/instanceid func_GetCreationTime
//go:noescape
func FuncGetCreationTime(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/instanceid call_GetCreationTime
//go:noescape
func CallGetCreationTime(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/instanceid try_GetCreationTime
//go:noescape
func TryGetCreationTime(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/instanceid has_GetID
//go:noescape
func HasFuncGetID() js.Ref

//go:wasmimport plat/js/webext/instanceid func_GetID
//go:noescape
func FuncGetID(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/instanceid call_GetID
//go:noescape
func CallGetID(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/instanceid try_GetID
//go:noescape
func TryGetID(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/instanceid has_GetToken
//go:noescape
func HasFuncGetToken() js.Ref

//go:wasmimport plat/js/webext/instanceid func_GetToken
//go:noescape
func FuncGetToken(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/instanceid call_GetToken
//go:noescape
func CallGetToken(
	retPtr unsafe.Pointer,
	getTokenParams unsafe.Pointer)

//go:wasmimport plat/js/webext/instanceid try_GetToken
//go:noescape
func TryGetToken(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	getTokenParams unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/instanceid has_OnTokenRefresh
//go:noescape
func HasFuncOnTokenRefresh() js.Ref

//go:wasmimport plat/js/webext/instanceid func_OnTokenRefresh
//go:noescape
func FuncOnTokenRefresh(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/instanceid call_OnTokenRefresh
//go:noescape
func CallOnTokenRefresh(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/instanceid try_OnTokenRefresh
//go:noescape
func TryOnTokenRefresh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/instanceid has_OffTokenRefresh
//go:noescape
func HasFuncOffTokenRefresh() js.Ref

//go:wasmimport plat/js/webext/instanceid func_OffTokenRefresh
//go:noescape
func FuncOffTokenRefresh(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/instanceid call_OffTokenRefresh
//go:noescape
func CallOffTokenRefresh(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/instanceid try_OffTokenRefresh
//go:noescape
func TryOffTokenRefresh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/instanceid has_HasOnTokenRefresh
//go:noescape
func HasFuncHasOnTokenRefresh() js.Ref

//go:wasmimport plat/js/webext/instanceid func_HasOnTokenRefresh
//go:noescape
func FuncHasOnTokenRefresh(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/instanceid call_HasOnTokenRefresh
//go:noescape
func CallHasOnTokenRefresh(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/instanceid try_HasOnTokenRefresh
//go:noescape
func TryHasOnTokenRefresh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
