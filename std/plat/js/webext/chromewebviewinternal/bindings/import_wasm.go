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

//go:wasmimport plat/js/webext/chromewebviewinternal store_ContextMenuItem
//go:noescape
func ContextMenuItemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal load_ContextMenuItem
//go:noescape
func ContextMenuItemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromewebviewinternal store_ContextMenusCreateArgCreateProperties
//go:noescape
func ContextMenusCreateArgCreatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal load_ContextMenusCreateArgCreateProperties
//go:noescape
func ContextMenusCreateArgCreatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromewebviewinternal store_ContextMenusUpdateArgUpdateProperties
//go:noescape
func ContextMenusUpdateArgUpdatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal load_ContextMenusUpdateArgUpdateProperties
//go:noescape
func ContextMenusUpdateArgUpdatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromewebviewinternal store_OnShowArgEvent
//go:noescape
func OnShowArgEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal load_OnShowArgEvent
//go:noescape
func OnShowArgEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromewebviewinternal has_ContextMenusCreate
//go:noescape
func HasFuncContextMenusCreate() js.Ref

//go:wasmimport plat/js/webext/chromewebviewinternal func_ContextMenusCreate
//go:noescape
func FuncContextMenusCreate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromewebviewinternal call_ContextMenusCreate
//go:noescape
func CallContextMenusCreate(
	retPtr unsafe.Pointer,
	instanceId float64,
	createProperties unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal try_ContextMenusCreate
//go:noescape
func TryContextMenusCreate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	createProperties unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal has_ContextMenusRemove
//go:noescape
func HasFuncContextMenusRemove() js.Ref

//go:wasmimport plat/js/webext/chromewebviewinternal func_ContextMenusRemove
//go:noescape
func FuncContextMenusRemove(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromewebviewinternal call_ContextMenusRemove
//go:noescape
func CallContextMenusRemove(
	retPtr unsafe.Pointer,
	instanceId float64,
	menuItemId js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal try_ContextMenusRemove
//go:noescape
func TryContextMenusRemove(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	menuItemId js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal has_ContextMenusRemoveAll
//go:noescape
func HasFuncContextMenusRemoveAll() js.Ref

//go:wasmimport plat/js/webext/chromewebviewinternal func_ContextMenusRemoveAll
//go:noescape
func FuncContextMenusRemoveAll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromewebviewinternal call_ContextMenusRemoveAll
//go:noescape
func CallContextMenusRemoveAll(
	retPtr unsafe.Pointer,
	instanceId float64,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal try_ContextMenusRemoveAll
//go:noescape
func TryContextMenusRemoveAll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal has_ContextMenusUpdate
//go:noescape
func HasFuncContextMenusUpdate() js.Ref

//go:wasmimport plat/js/webext/chromewebviewinternal func_ContextMenusUpdate
//go:noescape
func FuncContextMenusUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromewebviewinternal call_ContextMenusUpdate
//go:noescape
func CallContextMenusUpdate(
	retPtr unsafe.Pointer,
	instanceId float64,
	id js.Ref,
	updateProperties unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal try_ContextMenusUpdate
//go:noescape
func TryContextMenusUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	id js.Ref,
	updateProperties unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal has_OnClicked
//go:noescape
func HasFuncOnClicked() js.Ref

//go:wasmimport plat/js/webext/chromewebviewinternal func_OnClicked
//go:noescape
func FuncOnClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromewebviewinternal call_OnClicked
//go:noescape
func CallOnClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal try_OnClicked
//go:noescape
func TryOnClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal has_OffClicked
//go:noescape
func HasFuncOffClicked() js.Ref

//go:wasmimport plat/js/webext/chromewebviewinternal func_OffClicked
//go:noescape
func FuncOffClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromewebviewinternal call_OffClicked
//go:noescape
func CallOffClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal try_OffClicked
//go:noescape
func TryOffClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal has_HasOnClicked
//go:noescape
func HasFuncHasOnClicked() js.Ref

//go:wasmimport plat/js/webext/chromewebviewinternal func_HasOnClicked
//go:noescape
func FuncHasOnClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromewebviewinternal call_HasOnClicked
//go:noescape
func CallHasOnClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal try_HasOnClicked
//go:noescape
func TryHasOnClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal has_OnShow
//go:noescape
func HasFuncOnShow() js.Ref

//go:wasmimport plat/js/webext/chromewebviewinternal func_OnShow
//go:noescape
func FuncOnShow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromewebviewinternal call_OnShow
//go:noescape
func CallOnShow(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal try_OnShow
//go:noescape
func TryOnShow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal has_OffShow
//go:noescape
func HasFuncOffShow() js.Ref

//go:wasmimport plat/js/webext/chromewebviewinternal func_OffShow
//go:noescape
func FuncOffShow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromewebviewinternal call_OffShow
//go:noescape
func CallOffShow(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal try_OffShow
//go:noescape
func TryOffShow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal has_HasOnShow
//go:noescape
func HasFuncHasOnShow() js.Ref

//go:wasmimport plat/js/webext/chromewebviewinternal func_HasOnShow
//go:noescape
func FuncHasOnShow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromewebviewinternal call_HasOnShow
//go:noescape
func CallHasOnShow(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal try_HasOnShow
//go:noescape
func TryHasOnShow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal has_ShowContextMenu
//go:noescape
func HasFuncShowContextMenu() js.Ref

//go:wasmimport plat/js/webext/chromewebviewinternal func_ShowContextMenu
//go:noescape
func FuncShowContextMenu(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromewebviewinternal call_ShowContextMenu
//go:noescape
func CallShowContextMenu(
	retPtr unsafe.Pointer,
	instanceId float64,
	requestId float64,
	itemsToShow js.Ref)

//go:wasmimport plat/js/webext/chromewebviewinternal try_ShowContextMenu
//go:noescape
func TryShowContextMenu(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	requestId float64,
	itemsToShow js.Ref) (ok js.Ref)
