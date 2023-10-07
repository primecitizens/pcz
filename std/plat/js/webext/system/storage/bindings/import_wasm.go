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

//go:wasmimport plat/js/webext/system/storage constof_EjectDeviceResultCode
//go:noescape
func ConstOfEjectDeviceResultCode(str js.Ref) uint32

//go:wasmimport plat/js/webext/system/storage store_StorageAvailableCapacityInfo
//go:noescape
func StorageAvailableCapacityInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/storage load_StorageAvailableCapacityInfo
//go:noescape
func StorageAvailableCapacityInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/storage constof_StorageUnitType
//go:noescape
func ConstOfStorageUnitType(str js.Ref) uint32

//go:wasmimport plat/js/webext/system/storage store_StorageUnitInfo
//go:noescape
func StorageUnitInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/storage load_StorageUnitInfo
//go:noescape
func StorageUnitInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/storage has_EjectDevice
//go:noescape
func HasFuncEjectDevice() js.Ref

//go:wasmimport plat/js/webext/system/storage func_EjectDevice
//go:noescape
func FuncEjectDevice(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/storage call_EjectDevice
//go:noescape
func CallEjectDevice(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/system/storage try_EjectDevice
//go:noescape
func TryEjectDevice(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/storage has_GetAvailableCapacity
//go:noescape
func HasFuncGetAvailableCapacity() js.Ref

//go:wasmimport plat/js/webext/system/storage func_GetAvailableCapacity
//go:noescape
func FuncGetAvailableCapacity(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/storage call_GetAvailableCapacity
//go:noescape
func CallGetAvailableCapacity(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/system/storage try_GetAvailableCapacity
//go:noescape
func TryGetAvailableCapacity(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/storage has_GetInfo
//go:noescape
func HasFuncGetInfo() js.Ref

//go:wasmimport plat/js/webext/system/storage func_GetInfo
//go:noescape
func FuncGetInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/storage call_GetInfo
//go:noescape
func CallGetInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/system/storage try_GetInfo
//go:noescape
func TryGetInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/system/storage has_OnAttached
//go:noescape
func HasFuncOnAttached() js.Ref

//go:wasmimport plat/js/webext/system/storage func_OnAttached
//go:noescape
func FuncOnAttached(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/storage call_OnAttached
//go:noescape
func CallOnAttached(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/system/storage try_OnAttached
//go:noescape
func TryOnAttached(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/storage has_OffAttached
//go:noescape
func HasFuncOffAttached() js.Ref

//go:wasmimport plat/js/webext/system/storage func_OffAttached
//go:noescape
func FuncOffAttached(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/storage call_OffAttached
//go:noescape
func CallOffAttached(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/system/storage try_OffAttached
//go:noescape
func TryOffAttached(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/storage has_HasOnAttached
//go:noescape
func HasFuncHasOnAttached() js.Ref

//go:wasmimport plat/js/webext/system/storage func_HasOnAttached
//go:noescape
func FuncHasOnAttached(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/storage call_HasOnAttached
//go:noescape
func CallHasOnAttached(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/system/storage try_HasOnAttached
//go:noescape
func TryHasOnAttached(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/storage has_OnDetached
//go:noescape
func HasFuncOnDetached() js.Ref

//go:wasmimport plat/js/webext/system/storage func_OnDetached
//go:noescape
func FuncOnDetached(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/storage call_OnDetached
//go:noescape
func CallOnDetached(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/system/storage try_OnDetached
//go:noescape
func TryOnDetached(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/storage has_OffDetached
//go:noescape
func HasFuncOffDetached() js.Ref

//go:wasmimport plat/js/webext/system/storage func_OffDetached
//go:noescape
func FuncOffDetached(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/storage call_OffDetached
//go:noescape
func CallOffDetached(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/system/storage try_OffDetached
//go:noescape
func TryOffDetached(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/storage has_HasOnDetached
//go:noescape
func HasFuncHasOnDetached() js.Ref

//go:wasmimport plat/js/webext/system/storage func_HasOnDetached
//go:noescape
func FuncHasOnDetached(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/storage call_HasOnDetached
//go:noescape
func CallHasOnDetached(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/system/storage try_HasOnDetached
//go:noescape
func TryHasOnDetached(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
