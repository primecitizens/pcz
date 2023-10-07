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

//go:wasmimport plat/js/webext/filesystem store_AcceptOption
//go:noescape
func AcceptOptionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystem load_AcceptOption
//go:noescape
func AcceptOptionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystem constof_ChooseEntryType
//go:noescape
func ConstOfChooseEntryType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filesystem store_ChooseEntryOptions
//go:noescape
func ChooseEntryOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystem load_ChooseEntryOptions
//go:noescape
func ChooseEntryOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystem store_Volume
//go:noescape
func VolumeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystem load_Volume
//go:noescape
func VolumeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystem store_RequestFileSystemOptions
//go:noescape
func RequestFileSystemOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystem load_RequestFileSystemOptions
//go:noescape
func RequestFileSystemOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystem store_VolumeListChangedEvent
//go:noescape
func VolumeListChangedEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystem load_VolumeListChangedEvent
//go:noescape
func VolumeListChangedEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystem has_ChooseEntry
//go:noescape
func HasFuncChooseEntry() js.Ref

//go:wasmimport plat/js/webext/filesystem func_ChooseEntry
//go:noescape
func FuncChooseEntry(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystem call_ChooseEntry
//go:noescape
func CallChooseEntry(
	retPtr unsafe.Pointer,
	options unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystem try_ChooseEntry
//go:noescape
func TryChooseEntry(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystem has_GetDisplayPath
//go:noescape
func HasFuncGetDisplayPath() js.Ref

//go:wasmimport plat/js/webext/filesystem func_GetDisplayPath
//go:noescape
func FuncGetDisplayPath(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystem call_GetDisplayPath
//go:noescape
func CallGetDisplayPath(
	retPtr unsafe.Pointer,
	entry js.Ref)

//go:wasmimport plat/js/webext/filesystem try_GetDisplayPath
//go:noescape
func TryGetDisplayPath(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entry js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystem has_GetVolumeList
//go:noescape
func HasFuncGetVolumeList() js.Ref

//go:wasmimport plat/js/webext/filesystem func_GetVolumeList
//go:noescape
func FuncGetVolumeList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystem call_GetVolumeList
//go:noescape
func CallGetVolumeList(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystem try_GetVolumeList
//go:noescape
func TryGetVolumeList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystem has_GetWritableEntry
//go:noescape
func HasFuncGetWritableEntry() js.Ref

//go:wasmimport plat/js/webext/filesystem func_GetWritableEntry
//go:noescape
func FuncGetWritableEntry(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystem call_GetWritableEntry
//go:noescape
func CallGetWritableEntry(
	retPtr unsafe.Pointer,
	entry js.Ref)

//go:wasmimport plat/js/webext/filesystem try_GetWritableEntry
//go:noescape
func TryGetWritableEntry(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entry js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystem has_IsRestorable
//go:noescape
func HasFuncIsRestorable() js.Ref

//go:wasmimport plat/js/webext/filesystem func_IsRestorable
//go:noescape
func FuncIsRestorable(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystem call_IsRestorable
//go:noescape
func CallIsRestorable(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/filesystem try_IsRestorable
//go:noescape
func TryIsRestorable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystem has_IsWritableEntry
//go:noescape
func HasFuncIsWritableEntry() js.Ref

//go:wasmimport plat/js/webext/filesystem func_IsWritableEntry
//go:noescape
func FuncIsWritableEntry(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystem call_IsWritableEntry
//go:noescape
func CallIsWritableEntry(
	retPtr unsafe.Pointer,
	entry js.Ref)

//go:wasmimport plat/js/webext/filesystem try_IsWritableEntry
//go:noescape
func TryIsWritableEntry(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entry js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystem has_OnVolumeListChanged
//go:noescape
func HasFuncOnVolumeListChanged() js.Ref

//go:wasmimport plat/js/webext/filesystem func_OnVolumeListChanged
//go:noescape
func FuncOnVolumeListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystem call_OnVolumeListChanged
//go:noescape
func CallOnVolumeListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystem try_OnVolumeListChanged
//go:noescape
func TryOnVolumeListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystem has_OffVolumeListChanged
//go:noescape
func HasFuncOffVolumeListChanged() js.Ref

//go:wasmimport plat/js/webext/filesystem func_OffVolumeListChanged
//go:noescape
func FuncOffVolumeListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystem call_OffVolumeListChanged
//go:noescape
func CallOffVolumeListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystem try_OffVolumeListChanged
//go:noescape
func TryOffVolumeListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystem has_HasOnVolumeListChanged
//go:noescape
func HasFuncHasOnVolumeListChanged() js.Ref

//go:wasmimport plat/js/webext/filesystem func_HasOnVolumeListChanged
//go:noescape
func FuncHasOnVolumeListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystem call_HasOnVolumeListChanged
//go:noescape
func CallHasOnVolumeListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystem try_HasOnVolumeListChanged
//go:noescape
func TryHasOnVolumeListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystem has_RequestFileSystem
//go:noescape
func HasFuncRequestFileSystem() js.Ref

//go:wasmimport plat/js/webext/filesystem func_RequestFileSystem
//go:noescape
func FuncRequestFileSystem(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystem call_RequestFileSystem
//go:noescape
func CallRequestFileSystem(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystem try_RequestFileSystem
//go:noescape
func TryRequestFileSystem(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystem has_RestoreEntry
//go:noescape
func HasFuncRestoreEntry() js.Ref

//go:wasmimport plat/js/webext/filesystem func_RestoreEntry
//go:noescape
func FuncRestoreEntry(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystem call_RestoreEntry
//go:noescape
func CallRestoreEntry(
	retPtr unsafe.Pointer,
	id js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystem try_RestoreEntry
//go:noescape
func TryRestoreEntry(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystem has_RetainEntry
//go:noescape
func HasFuncRetainEntry() js.Ref

//go:wasmimport plat/js/webext/filesystem func_RetainEntry
//go:noescape
func FuncRetainEntry(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystem call_RetainEntry
//go:noescape
func CallRetainEntry(
	retPtr unsafe.Pointer,
	entry js.Ref)

//go:wasmimport plat/js/webext/filesystem try_RetainEntry
//go:noescape
func TryRetainEntry(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entry js.Ref) (ok js.Ref)
