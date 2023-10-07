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

//go:wasmimport plat/js/webext/readinglist store_AddEntryOptions
//go:noescape
func AddEntryOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/readinglist load_AddEntryOptions
//go:noescape
func AddEntryOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/readinglist store_ReadingListEntry
//go:noescape
func ReadingListEntryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/readinglist load_ReadingListEntry
//go:noescape
func ReadingListEntryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/readinglist store_QueryInfo
//go:noescape
func QueryInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/readinglist load_QueryInfo
//go:noescape
func QueryInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/readinglist store_RemoveOptions
//go:noescape
func RemoveOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/readinglist load_RemoveOptions
//go:noescape
func RemoveOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/readinglist store_UpdateEntryOptions
//go:noescape
func UpdateEntryOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/readinglist load_UpdateEntryOptions
//go:noescape
func UpdateEntryOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/readinglist has_AddEntry
//go:noescape
func HasFuncAddEntry() js.Ref

//go:wasmimport plat/js/webext/readinglist func_AddEntry
//go:noescape
func FuncAddEntry(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist call_AddEntry
//go:noescape
func CallAddEntry(
	retPtr unsafe.Pointer,
	entry unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist try_AddEntry
//go:noescape
func TryAddEntry(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entry unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/readinglist has_OnEntryAdded
//go:noescape
func HasFuncOnEntryAdded() js.Ref

//go:wasmimport plat/js/webext/readinglist func_OnEntryAdded
//go:noescape
func FuncOnEntryAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist call_OnEntryAdded
//go:noescape
func CallOnEntryAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/readinglist try_OnEntryAdded
//go:noescape
func TryOnEntryAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/readinglist has_OffEntryAdded
//go:noescape
func HasFuncOffEntryAdded() js.Ref

//go:wasmimport plat/js/webext/readinglist func_OffEntryAdded
//go:noescape
func FuncOffEntryAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist call_OffEntryAdded
//go:noescape
func CallOffEntryAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/readinglist try_OffEntryAdded
//go:noescape
func TryOffEntryAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/readinglist has_HasOnEntryAdded
//go:noescape
func HasFuncHasOnEntryAdded() js.Ref

//go:wasmimport plat/js/webext/readinglist func_HasOnEntryAdded
//go:noescape
func FuncHasOnEntryAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist call_HasOnEntryAdded
//go:noescape
func CallHasOnEntryAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/readinglist try_HasOnEntryAdded
//go:noescape
func TryHasOnEntryAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/readinglist has_OnEntryRemoved
//go:noescape
func HasFuncOnEntryRemoved() js.Ref

//go:wasmimport plat/js/webext/readinglist func_OnEntryRemoved
//go:noescape
func FuncOnEntryRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist call_OnEntryRemoved
//go:noescape
func CallOnEntryRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/readinglist try_OnEntryRemoved
//go:noescape
func TryOnEntryRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/readinglist has_OffEntryRemoved
//go:noescape
func HasFuncOffEntryRemoved() js.Ref

//go:wasmimport plat/js/webext/readinglist func_OffEntryRemoved
//go:noescape
func FuncOffEntryRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist call_OffEntryRemoved
//go:noescape
func CallOffEntryRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/readinglist try_OffEntryRemoved
//go:noescape
func TryOffEntryRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/readinglist has_HasOnEntryRemoved
//go:noescape
func HasFuncHasOnEntryRemoved() js.Ref

//go:wasmimport plat/js/webext/readinglist func_HasOnEntryRemoved
//go:noescape
func FuncHasOnEntryRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist call_HasOnEntryRemoved
//go:noescape
func CallHasOnEntryRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/readinglist try_HasOnEntryRemoved
//go:noescape
func TryHasOnEntryRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/readinglist has_OnEntryUpdated
//go:noescape
func HasFuncOnEntryUpdated() js.Ref

//go:wasmimport plat/js/webext/readinglist func_OnEntryUpdated
//go:noescape
func FuncOnEntryUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist call_OnEntryUpdated
//go:noescape
func CallOnEntryUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/readinglist try_OnEntryUpdated
//go:noescape
func TryOnEntryUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/readinglist has_OffEntryUpdated
//go:noescape
func HasFuncOffEntryUpdated() js.Ref

//go:wasmimport plat/js/webext/readinglist func_OffEntryUpdated
//go:noescape
func FuncOffEntryUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist call_OffEntryUpdated
//go:noescape
func CallOffEntryUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/readinglist try_OffEntryUpdated
//go:noescape
func TryOffEntryUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/readinglist has_HasOnEntryUpdated
//go:noescape
func HasFuncHasOnEntryUpdated() js.Ref

//go:wasmimport plat/js/webext/readinglist func_HasOnEntryUpdated
//go:noescape
func FuncHasOnEntryUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist call_HasOnEntryUpdated
//go:noescape
func CallHasOnEntryUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/readinglist try_HasOnEntryUpdated
//go:noescape
func TryHasOnEntryUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/readinglist has_Query
//go:noescape
func HasFuncQuery() js.Ref

//go:wasmimport plat/js/webext/readinglist func_Query
//go:noescape
func FuncQuery(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist call_Query
//go:noescape
func CallQuery(
	retPtr unsafe.Pointer,
	info unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist try_Query
//go:noescape
func TryQuery(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	info unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/readinglist has_RemoveEntry
//go:noescape
func HasFuncRemoveEntry() js.Ref

//go:wasmimport plat/js/webext/readinglist func_RemoveEntry
//go:noescape
func FuncRemoveEntry(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist call_RemoveEntry
//go:noescape
func CallRemoveEntry(
	retPtr unsafe.Pointer,
	info unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist try_RemoveEntry
//go:noescape
func TryRemoveEntry(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	info unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/readinglist has_UpdateEntry
//go:noescape
func HasFuncUpdateEntry() js.Ref

//go:wasmimport plat/js/webext/readinglist func_UpdateEntry
//go:noescape
func FuncUpdateEntry(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist call_UpdateEntry
//go:noescape
func CallUpdateEntry(
	retPtr unsafe.Pointer,
	info unsafe.Pointer)

//go:wasmimport plat/js/webext/readinglist try_UpdateEntry
//go:noescape
func TryUpdateEntry(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	info unsafe.Pointer) (ok js.Ref)
