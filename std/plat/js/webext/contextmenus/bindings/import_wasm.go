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

//go:wasmimport plat/js/webext/contextmenus get_ACTION_MENU_TOP_LEVEL_LIMIT
//go:noescape
func GetACTION_MENU_TOP_LEVEL_LIMIT(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/contextmenus set_ACTION_MENU_TOP_LEVEL_LIMIT
//go:noescape
func SetACTION_MENU_TOP_LEVEL_LIMIT(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/contextmenus store_OnClickData
//go:noescape
func OnClickDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/contextmenus load_OnClickData
//go:noescape
func OnClickDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/contextmenus constof_ItemType
//go:noescape
func ConstOfItemType(str js.Ref) uint32

//go:wasmimport plat/js/webext/contextmenus constof_ContextType
//go:noescape
func ConstOfContextType(str js.Ref) uint32

//go:wasmimport plat/js/webext/contextmenus store_CreateArgCreateProperties
//go:noescape
func CreateArgCreatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/contextmenus load_CreateArgCreateProperties
//go:noescape
func CreateArgCreatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/contextmenus store_UpdateArgUpdateProperties
//go:noescape
func UpdateArgUpdatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/contextmenus load_UpdateArgUpdateProperties
//go:noescape
func UpdateArgUpdatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/contextmenus has_Create
//go:noescape
func HasFuncCreate() js.Ref

//go:wasmimport plat/js/webext/contextmenus func_Create
//go:noescape
func FuncCreate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/contextmenus call_Create
//go:noescape
func CallCreate(
	retPtr unsafe.Pointer,
	createProperties unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/contextmenus try_Create
//go:noescape
func TryCreate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	createProperties unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/contextmenus has_OnClicked
//go:noescape
func HasFuncOnClicked() js.Ref

//go:wasmimport plat/js/webext/contextmenus func_OnClicked
//go:noescape
func FuncOnClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/contextmenus call_OnClicked
//go:noescape
func CallOnClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/contextmenus try_OnClicked
//go:noescape
func TryOnClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/contextmenus has_OffClicked
//go:noescape
func HasFuncOffClicked() js.Ref

//go:wasmimport plat/js/webext/contextmenus func_OffClicked
//go:noescape
func FuncOffClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/contextmenus call_OffClicked
//go:noescape
func CallOffClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/contextmenus try_OffClicked
//go:noescape
func TryOffClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/contextmenus has_HasOnClicked
//go:noescape
func HasFuncHasOnClicked() js.Ref

//go:wasmimport plat/js/webext/contextmenus func_HasOnClicked
//go:noescape
func FuncHasOnClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/contextmenus call_HasOnClicked
//go:noescape
func CallHasOnClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/contextmenus try_HasOnClicked
//go:noescape
func TryHasOnClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/contextmenus has_Remove
//go:noescape
func HasFuncRemove() js.Ref

//go:wasmimport plat/js/webext/contextmenus func_Remove
//go:noescape
func FuncRemove(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/contextmenus call_Remove
//go:noescape
func CallRemove(
	retPtr unsafe.Pointer,
	menuItemId js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/contextmenus try_Remove
//go:noescape
func TryRemove(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	menuItemId js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/contextmenus has_RemoveAll
//go:noescape
func HasFuncRemoveAll() js.Ref

//go:wasmimport plat/js/webext/contextmenus func_RemoveAll
//go:noescape
func FuncRemoveAll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/contextmenus call_RemoveAll
//go:noescape
func CallRemoveAll(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/contextmenus try_RemoveAll
//go:noescape
func TryRemoveAll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/contextmenus has_Update
//go:noescape
func HasFuncUpdate() js.Ref

//go:wasmimport plat/js/webext/contextmenus func_Update
//go:noescape
func FuncUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/contextmenus call_Update
//go:noescape
func CallUpdate(
	retPtr unsafe.Pointer,
	id js.Ref,
	updateProperties unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/contextmenus try_Update
//go:noescape
func TryUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	updateProperties unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
