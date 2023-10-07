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

//go:wasmimport plat/js/webext/imagewriterprivate store_RemovableStorageDevice
//go:noescape
func RemovableStorageDeviceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate load_RemovableStorageDevice
//go:noescape
func RemovableStorageDeviceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate constof_Stage
//go:noescape
func ConstOfStage(str js.Ref) uint32

//go:wasmimport plat/js/webext/imagewriterprivate store_ProgressInfo
//go:noescape
func ProgressInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate load_ProgressInfo
//go:noescape
func ProgressInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate store_UrlWriteOptions
//go:noescape
func UrlWriteOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate load_UrlWriteOptions
//go:noescape
func UrlWriteOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate has_CancelWrite
//go:noescape
func HasFuncCancelWrite() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_CancelWrite
//go:noescape
func FuncCancelWrite(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_CancelWrite
//go:noescape
func CallCancelWrite(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate try_CancelWrite
//go:noescape
func TryCancelWrite(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_DestroyPartitions
//go:noescape
func HasFuncDestroyPartitions() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_DestroyPartitions
//go:noescape
func FuncDestroyPartitions(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_DestroyPartitions
//go:noescape
func CallDestroyPartitions(
	retPtr unsafe.Pointer,
	storageUnitId js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_DestroyPartitions
//go:noescape
func TryDestroyPartitions(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	storageUnitId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_ListRemovableStorageDevices
//go:noescape
func HasFuncListRemovableStorageDevices() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_ListRemovableStorageDevices
//go:noescape
func FuncListRemovableStorageDevices(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_ListRemovableStorageDevices
//go:noescape
func CallListRemovableStorageDevices(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate try_ListRemovableStorageDevices
//go:noescape
func TryListRemovableStorageDevices(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_OnDeviceInserted
//go:noescape
func HasFuncOnDeviceInserted() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_OnDeviceInserted
//go:noescape
func FuncOnDeviceInserted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_OnDeviceInserted
//go:noescape
func CallOnDeviceInserted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_OnDeviceInserted
//go:noescape
func TryOnDeviceInserted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_OffDeviceInserted
//go:noescape
func HasFuncOffDeviceInserted() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_OffDeviceInserted
//go:noescape
func FuncOffDeviceInserted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_OffDeviceInserted
//go:noescape
func CallOffDeviceInserted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_OffDeviceInserted
//go:noescape
func TryOffDeviceInserted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_HasOnDeviceInserted
//go:noescape
func HasFuncHasOnDeviceInserted() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_HasOnDeviceInserted
//go:noescape
func FuncHasOnDeviceInserted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_HasOnDeviceInserted
//go:noescape
func CallHasOnDeviceInserted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_HasOnDeviceInserted
//go:noescape
func TryHasOnDeviceInserted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_OnDeviceRemoved
//go:noescape
func HasFuncOnDeviceRemoved() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_OnDeviceRemoved
//go:noescape
func FuncOnDeviceRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_OnDeviceRemoved
//go:noescape
func CallOnDeviceRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_OnDeviceRemoved
//go:noescape
func TryOnDeviceRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_OffDeviceRemoved
//go:noescape
func HasFuncOffDeviceRemoved() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_OffDeviceRemoved
//go:noescape
func FuncOffDeviceRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_OffDeviceRemoved
//go:noescape
func CallOffDeviceRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_OffDeviceRemoved
//go:noescape
func TryOffDeviceRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_HasOnDeviceRemoved
//go:noescape
func HasFuncHasOnDeviceRemoved() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_HasOnDeviceRemoved
//go:noescape
func FuncHasOnDeviceRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_HasOnDeviceRemoved
//go:noescape
func CallHasOnDeviceRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_HasOnDeviceRemoved
//go:noescape
func TryHasOnDeviceRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_OnWriteComplete
//go:noescape
func HasFuncOnWriteComplete() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_OnWriteComplete
//go:noescape
func FuncOnWriteComplete(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_OnWriteComplete
//go:noescape
func CallOnWriteComplete(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_OnWriteComplete
//go:noescape
func TryOnWriteComplete(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_OffWriteComplete
//go:noescape
func HasFuncOffWriteComplete() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_OffWriteComplete
//go:noescape
func FuncOffWriteComplete(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_OffWriteComplete
//go:noescape
func CallOffWriteComplete(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_OffWriteComplete
//go:noescape
func TryOffWriteComplete(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_HasOnWriteComplete
//go:noescape
func HasFuncHasOnWriteComplete() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_HasOnWriteComplete
//go:noescape
func FuncHasOnWriteComplete(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_HasOnWriteComplete
//go:noescape
func CallHasOnWriteComplete(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_HasOnWriteComplete
//go:noescape
func TryHasOnWriteComplete(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_OnWriteError
//go:noescape
func HasFuncOnWriteError() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_OnWriteError
//go:noescape
func FuncOnWriteError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_OnWriteError
//go:noescape
func CallOnWriteError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_OnWriteError
//go:noescape
func TryOnWriteError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_OffWriteError
//go:noescape
func HasFuncOffWriteError() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_OffWriteError
//go:noescape
func FuncOffWriteError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_OffWriteError
//go:noescape
func CallOffWriteError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_OffWriteError
//go:noescape
func TryOffWriteError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_HasOnWriteError
//go:noescape
func HasFuncHasOnWriteError() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_HasOnWriteError
//go:noescape
func FuncHasOnWriteError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_HasOnWriteError
//go:noescape
func CallHasOnWriteError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_HasOnWriteError
//go:noescape
func TryHasOnWriteError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_OnWriteProgress
//go:noescape
func HasFuncOnWriteProgress() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_OnWriteProgress
//go:noescape
func FuncOnWriteProgress(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_OnWriteProgress
//go:noescape
func CallOnWriteProgress(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_OnWriteProgress
//go:noescape
func TryOnWriteProgress(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_OffWriteProgress
//go:noescape
func HasFuncOffWriteProgress() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_OffWriteProgress
//go:noescape
func FuncOffWriteProgress(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_OffWriteProgress
//go:noescape
func CallOffWriteProgress(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_OffWriteProgress
//go:noescape
func TryOffWriteProgress(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_HasOnWriteProgress
//go:noescape
func HasFuncHasOnWriteProgress() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_HasOnWriteProgress
//go:noescape
func FuncHasOnWriteProgress(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_HasOnWriteProgress
//go:noescape
func CallHasOnWriteProgress(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_HasOnWriteProgress
//go:noescape
func TryHasOnWriteProgress(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_WriteFromFile
//go:noescape
func HasFuncWriteFromFile() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_WriteFromFile
//go:noescape
func FuncWriteFromFile(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_WriteFromFile
//go:noescape
func CallWriteFromFile(
	retPtr unsafe.Pointer,
	storageUnitId js.Ref,
	fileEntry js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate try_WriteFromFile
//go:noescape
func TryWriteFromFile(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	storageUnitId js.Ref,
	fileEntry js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imagewriterprivate has_WriteFromUrl
//go:noescape
func HasFuncWriteFromUrl() js.Ref

//go:wasmimport plat/js/webext/imagewriterprivate func_WriteFromUrl
//go:noescape
func FuncWriteFromUrl(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate call_WriteFromUrl
//go:noescape
func CallWriteFromUrl(
	retPtr unsafe.Pointer,
	storageUnitId js.Ref,
	imageUrl js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/imagewriterprivate try_WriteFromUrl
//go:noescape
func TryWriteFromUrl(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	storageUnitId js.Ref,
	imageUrl js.Ref,
	options unsafe.Pointer) (ok js.Ref)
