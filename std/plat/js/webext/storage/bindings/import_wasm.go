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

//go:wasmimport plat/js/webext/storage constof_AccessLevel
//go:noescape
func ConstOfAccessLevel(str js.Ref) uint32

//go:wasmimport plat/js/webext/storage store_SetAccessLevelArgAccessOptions
//go:noescape
func SetAccessLevelArgAccessOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/storage load_SetAccessLevelArgAccessOptions
//go:noescape
func SetAccessLevelArgAccessOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/storage has_StorageArea_Get
//go:noescape
func HasFuncStorageAreaGet(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/storage func_StorageArea_Get
//go:noescape
func FuncStorageAreaGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/storage call_StorageArea_Get
//go:noescape
func CallStorageAreaGet(
	this js.Ref, retPtr unsafe.Pointer,
	keys js.Ref)

//go:wasmimport plat/js/webext/storage try_StorageArea_Get
//go:noescape
func TryStorageAreaGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	keys js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/storage has_StorageArea_Get1
//go:noescape
func HasFuncStorageAreaGet1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/storage func_StorageArea_Get1
//go:noescape
func FuncStorageAreaGet1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/storage call_StorageArea_Get1
//go:noescape
func CallStorageAreaGet1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/storage try_StorageArea_Get1
//go:noescape
func TryStorageAreaGet1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/storage has_StorageArea_GetBytesInUse
//go:noescape
func HasFuncStorageAreaGetBytesInUse(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/storage func_StorageArea_GetBytesInUse
//go:noescape
func FuncStorageAreaGetBytesInUse(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/storage call_StorageArea_GetBytesInUse
//go:noescape
func CallStorageAreaGetBytesInUse(
	this js.Ref, retPtr unsafe.Pointer,
	keys js.Ref)

//go:wasmimport plat/js/webext/storage try_StorageArea_GetBytesInUse
//go:noescape
func TryStorageAreaGetBytesInUse(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	keys js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/storage has_StorageArea_GetBytesInUse1
//go:noescape
func HasFuncStorageAreaGetBytesInUse1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/storage func_StorageArea_GetBytesInUse1
//go:noescape
func FuncStorageAreaGetBytesInUse1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/storage call_StorageArea_GetBytesInUse1
//go:noescape
func CallStorageAreaGetBytesInUse1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/storage try_StorageArea_GetBytesInUse1
//go:noescape
func TryStorageAreaGetBytesInUse1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/storage has_StorageArea_Set
//go:noescape
func HasFuncStorageAreaSet(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/storage func_StorageArea_Set
//go:noescape
func FuncStorageAreaSet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/storage call_StorageArea_Set
//go:noescape
func CallStorageAreaSet(
	this js.Ref, retPtr unsafe.Pointer,
	items js.Ref)

//go:wasmimport plat/js/webext/storage try_StorageArea_Set
//go:noescape
func TryStorageAreaSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	items js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/storage has_StorageArea_Remove
//go:noescape
func HasFuncStorageAreaRemove(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/storage func_StorageArea_Remove
//go:noescape
func FuncStorageAreaRemove(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/storage call_StorageArea_Remove
//go:noescape
func CallStorageAreaRemove(
	this js.Ref, retPtr unsafe.Pointer,
	keys js.Ref)

//go:wasmimport plat/js/webext/storage try_StorageArea_Remove
//go:noescape
func TryStorageAreaRemove(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	keys js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/storage has_StorageArea_Clear
//go:noescape
func HasFuncStorageAreaClear(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/storage func_StorageArea_Clear
//go:noescape
func FuncStorageAreaClear(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/storage call_StorageArea_Clear
//go:noescape
func CallStorageAreaClear(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/storage try_StorageArea_Clear
//go:noescape
func TryStorageAreaClear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/storage has_StorageArea_SetAccessLevel
//go:noescape
func HasFuncStorageAreaSetAccessLevel(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/storage func_StorageArea_SetAccessLevel
//go:noescape
func FuncStorageAreaSetAccessLevel(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/storage call_StorageArea_SetAccessLevel
//go:noescape
func CallStorageAreaSetAccessLevel(
	this js.Ref, retPtr unsafe.Pointer,
	accessOptions unsafe.Pointer)

//go:wasmimport plat/js/webext/storage try_StorageArea_SetAccessLevel
//go:noescape
func TryStorageAreaSetAccessLevel(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	accessOptions unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/storage store_StorageChange
//go:noescape
func StorageChangeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/storage load_StorageChange
//go:noescape
func StorageChangeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/storage get_Local
//go:noescape
func GetLocal(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/storage set_Local
//go:noescape
func SetLocal(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/storage get_Managed
//go:noescape
func GetManaged(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/storage set_Managed
//go:noescape
func SetManaged(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/storage has_OnChanged
//go:noescape
func HasFuncOnChanged() js.Ref

//go:wasmimport plat/js/webext/storage func_OnChanged
//go:noescape
func FuncOnChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/storage call_OnChanged
//go:noescape
func CallOnChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/storage try_OnChanged
//go:noescape
func TryOnChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/storage has_OffChanged
//go:noescape
func HasFuncOffChanged() js.Ref

//go:wasmimport plat/js/webext/storage func_OffChanged
//go:noescape
func FuncOffChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/storage call_OffChanged
//go:noescape
func CallOffChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/storage try_OffChanged
//go:noescape
func TryOffChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/storage has_HasOnChanged
//go:noescape
func HasFuncHasOnChanged() js.Ref

//go:wasmimport plat/js/webext/storage func_HasOnChanged
//go:noescape
func FuncHasOnChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/storage call_HasOnChanged
//go:noescape
func CallHasOnChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/storage try_HasOnChanged
//go:noescape
func TryHasOnChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/storage get_Session
//go:noescape
func GetSession(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/storage set_Session
//go:noescape
func SetSession(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/storage get_Sync
//go:noescape
func GetSync(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/storage set_Sync
//go:noescape
func SetSync(
	val js.Ref) js.Ref
