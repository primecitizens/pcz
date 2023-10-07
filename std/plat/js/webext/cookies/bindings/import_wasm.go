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

//go:wasmimport plat/js/webext/cookies constof_SameSiteStatus
//go:noescape
func ConstOfSameSiteStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/cookies store_Cookie
//go:noescape
func CookieJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/cookies load_Cookie
//go:noescape
func CookieJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/cookies store_CookieDetails
//go:noescape
func CookieDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/cookies load_CookieDetails
//go:noescape
func CookieDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/cookies store_CookieStore
//go:noescape
func CookieStoreJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/cookies load_CookieStore
//go:noescape
func CookieStoreJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/cookies store_GetAllArgDetails
//go:noescape
func GetAllArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/cookies load_GetAllArgDetails
//go:noescape
func GetAllArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/cookies constof_OnChangedCause
//go:noescape
func ConstOfOnChangedCause(str js.Ref) uint32

//go:wasmimport plat/js/webext/cookies store_OnChangedArgChangeInfo
//go:noescape
func OnChangedArgChangeInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/cookies load_OnChangedArgChangeInfo
//go:noescape
func OnChangedArgChangeInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/cookies store_RemoveReturnType
//go:noescape
func RemoveReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/cookies load_RemoveReturnType
//go:noescape
func RemoveReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/cookies store_SetArgDetails
//go:noescape
func SetArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/cookies load_SetArgDetails
//go:noescape
func SetArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/cookies has_Get
//go:noescape
func HasFuncGet() js.Ref

//go:wasmimport plat/js/webext/cookies func_Get
//go:noescape
func FuncGet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/cookies call_Get
//go:noescape
func CallGet(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/cookies try_Get
//go:noescape
func TryGet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/cookies has_GetAll
//go:noescape
func HasFuncGetAll() js.Ref

//go:wasmimport plat/js/webext/cookies func_GetAll
//go:noescape
func FuncGetAll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/cookies call_GetAll
//go:noescape
func CallGetAll(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/cookies try_GetAll
//go:noescape
func TryGetAll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/cookies has_GetAllCookieStores
//go:noescape
func HasFuncGetAllCookieStores() js.Ref

//go:wasmimport plat/js/webext/cookies func_GetAllCookieStores
//go:noescape
func FuncGetAllCookieStores(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/cookies call_GetAllCookieStores
//go:noescape
func CallGetAllCookieStores(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/cookies try_GetAllCookieStores
//go:noescape
func TryGetAllCookieStores(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/cookies has_OnChanged
//go:noescape
func HasFuncOnChanged() js.Ref

//go:wasmimport plat/js/webext/cookies func_OnChanged
//go:noescape
func FuncOnChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/cookies call_OnChanged
//go:noescape
func CallOnChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/cookies try_OnChanged
//go:noescape
func TryOnChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/cookies has_OffChanged
//go:noescape
func HasFuncOffChanged() js.Ref

//go:wasmimport plat/js/webext/cookies func_OffChanged
//go:noescape
func FuncOffChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/cookies call_OffChanged
//go:noescape
func CallOffChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/cookies try_OffChanged
//go:noescape
func TryOffChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/cookies has_HasOnChanged
//go:noescape
func HasFuncHasOnChanged() js.Ref

//go:wasmimport plat/js/webext/cookies func_HasOnChanged
//go:noescape
func FuncHasOnChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/cookies call_HasOnChanged
//go:noescape
func CallHasOnChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/cookies try_HasOnChanged
//go:noescape
func TryHasOnChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/cookies has_Remove
//go:noescape
func HasFuncRemove() js.Ref

//go:wasmimport plat/js/webext/cookies func_Remove
//go:noescape
func FuncRemove(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/cookies call_Remove
//go:noescape
func CallRemove(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/cookies try_Remove
//go:noescape
func TryRemove(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/cookies has_Set
//go:noescape
func HasFuncSet() js.Ref

//go:wasmimport plat/js/webext/cookies func_Set
//go:noescape
func FuncSet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/cookies call_Set
//go:noescape
func CallSet(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/cookies try_Set
//go:noescape
func TrySet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)
