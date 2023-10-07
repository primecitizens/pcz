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

//go:wasmimport plat/js/webext/lockscreen/data store_DataItemInfo
//go:noescape
func DataItemInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/lockscreen/data load_DataItemInfo
//go:noescape
func DataItemInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/lockscreen/data store_DataItemsAvailableEvent
//go:noescape
func DataItemsAvailableEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/lockscreen/data load_DataItemsAvailableEvent
//go:noescape
func DataItemsAvailableEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/lockscreen/data has_Create
//go:noescape
func HasFuncCreate() js.Ref

//go:wasmimport plat/js/webext/lockscreen/data func_Create
//go:noescape
func FuncCreate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/lockscreen/data call_Create
//go:noescape
func CallCreate(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/lockscreen/data try_Create
//go:noescape
func TryCreate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/lockscreen/data has_Delete
//go:noescape
func HasFuncDelete() js.Ref

//go:wasmimport plat/js/webext/lockscreen/data func_Delete
//go:noescape
func FuncDelete(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/lockscreen/data call_Delete
//go:noescape
func CallDelete(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/lockscreen/data try_Delete
//go:noescape
func TryDelete(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/lockscreen/data has_GetAll
//go:noescape
func HasFuncGetAll() js.Ref

//go:wasmimport plat/js/webext/lockscreen/data func_GetAll
//go:noescape
func FuncGetAll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/lockscreen/data call_GetAll
//go:noescape
func CallGetAll(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/lockscreen/data try_GetAll
//go:noescape
func TryGetAll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/lockscreen/data has_GetContent
//go:noescape
func HasFuncGetContent() js.Ref

//go:wasmimport plat/js/webext/lockscreen/data func_GetContent
//go:noescape
func FuncGetContent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/lockscreen/data call_GetContent
//go:noescape
func CallGetContent(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/lockscreen/data try_GetContent
//go:noescape
func TryGetContent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/lockscreen/data has_OnDataItemsAvailable
//go:noescape
func HasFuncOnDataItemsAvailable() js.Ref

//go:wasmimport plat/js/webext/lockscreen/data func_OnDataItemsAvailable
//go:noescape
func FuncOnDataItemsAvailable(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/lockscreen/data call_OnDataItemsAvailable
//go:noescape
func CallOnDataItemsAvailable(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/lockscreen/data try_OnDataItemsAvailable
//go:noescape
func TryOnDataItemsAvailable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/lockscreen/data has_OffDataItemsAvailable
//go:noescape
func HasFuncOffDataItemsAvailable() js.Ref

//go:wasmimport plat/js/webext/lockscreen/data func_OffDataItemsAvailable
//go:noescape
func FuncOffDataItemsAvailable(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/lockscreen/data call_OffDataItemsAvailable
//go:noescape
func CallOffDataItemsAvailable(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/lockscreen/data try_OffDataItemsAvailable
//go:noescape
func TryOffDataItemsAvailable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/lockscreen/data has_HasOnDataItemsAvailable
//go:noescape
func HasFuncHasOnDataItemsAvailable() js.Ref

//go:wasmimport plat/js/webext/lockscreen/data func_HasOnDataItemsAvailable
//go:noescape
func FuncHasOnDataItemsAvailable(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/lockscreen/data call_HasOnDataItemsAvailable
//go:noescape
func CallHasOnDataItemsAvailable(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/lockscreen/data try_HasOnDataItemsAvailable
//go:noescape
func TryHasOnDataItemsAvailable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/lockscreen/data has_SetContent
//go:noescape
func HasFuncSetContent() js.Ref

//go:wasmimport plat/js/webext/lockscreen/data func_SetContent
//go:noescape
func FuncSetContent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/lockscreen/data call_SetContent
//go:noescape
func CallSetContent(
	retPtr unsafe.Pointer,
	id js.Ref,
	data js.Ref)

//go:wasmimport plat/js/webext/lockscreen/data try_SetContent
//go:noescape
func TrySetContent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	data js.Ref) (ok js.Ref)
