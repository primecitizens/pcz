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

//go:wasmimport plat/js/webext/notifications store_NotificationBitmap
//go:noescape
func NotificationBitmapJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/notifications load_NotificationBitmap
//go:noescape
func NotificationBitmapJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/notifications store_NotificationButton
//go:noescape
func NotificationButtonJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/notifications load_NotificationButton
//go:noescape
func NotificationButtonJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/notifications store_NotificationItem
//go:noescape
func NotificationItemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/notifications load_NotificationItem
//go:noescape
func NotificationItemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/notifications constof_TemplateType
//go:noescape
func ConstOfTemplateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/notifications store_NotificationOptions
//go:noescape
func NotificationOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/notifications load_NotificationOptions
//go:noescape
func NotificationOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/notifications constof_PermissionLevel
//go:noescape
func ConstOfPermissionLevel(str js.Ref) uint32

//go:wasmimport plat/js/webext/notifications has_Clear
//go:noescape
func HasFuncClear() js.Ref

//go:wasmimport plat/js/webext/notifications func_Clear
//go:noescape
func FuncClear(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_Clear
//go:noescape
func CallClear(
	retPtr unsafe.Pointer,
	notificationId js.Ref)

//go:wasmimport plat/js/webext/notifications try_Clear
//go:noescape
func TryClear(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	notificationId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_Create
//go:noescape
func HasFuncCreate() js.Ref

//go:wasmimport plat/js/webext/notifications func_Create
//go:noescape
func FuncCreate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_Create
//go:noescape
func CallCreate(
	retPtr unsafe.Pointer,
	notificationId js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications try_Create
//go:noescape
func TryCreate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	notificationId js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_GetAll
//go:noescape
func HasFuncGetAll() js.Ref

//go:wasmimport plat/js/webext/notifications func_GetAll
//go:noescape
func FuncGetAll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_GetAll
//go:noescape
func CallGetAll(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications try_GetAll
//go:noescape
func TryGetAll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_GetPermissionLevel
//go:noescape
func HasFuncGetPermissionLevel() js.Ref

//go:wasmimport plat/js/webext/notifications func_GetPermissionLevel
//go:noescape
func FuncGetPermissionLevel(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_GetPermissionLevel
//go:noescape
func CallGetPermissionLevel(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications try_GetPermissionLevel
//go:noescape
func TryGetPermissionLevel(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_OnButtonClicked
//go:noescape
func HasFuncOnButtonClicked() js.Ref

//go:wasmimport plat/js/webext/notifications func_OnButtonClicked
//go:noescape
func FuncOnButtonClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_OnButtonClicked
//go:noescape
func CallOnButtonClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/notifications try_OnButtonClicked
//go:noescape
func TryOnButtonClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_OffButtonClicked
//go:noescape
func HasFuncOffButtonClicked() js.Ref

//go:wasmimport plat/js/webext/notifications func_OffButtonClicked
//go:noescape
func FuncOffButtonClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_OffButtonClicked
//go:noescape
func CallOffButtonClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/notifications try_OffButtonClicked
//go:noescape
func TryOffButtonClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_HasOnButtonClicked
//go:noescape
func HasFuncHasOnButtonClicked() js.Ref

//go:wasmimport plat/js/webext/notifications func_HasOnButtonClicked
//go:noescape
func FuncHasOnButtonClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_HasOnButtonClicked
//go:noescape
func CallHasOnButtonClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/notifications try_HasOnButtonClicked
//go:noescape
func TryHasOnButtonClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_OnClicked
//go:noescape
func HasFuncOnClicked() js.Ref

//go:wasmimport plat/js/webext/notifications func_OnClicked
//go:noescape
func FuncOnClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_OnClicked
//go:noescape
func CallOnClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/notifications try_OnClicked
//go:noescape
func TryOnClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_OffClicked
//go:noescape
func HasFuncOffClicked() js.Ref

//go:wasmimport plat/js/webext/notifications func_OffClicked
//go:noescape
func FuncOffClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_OffClicked
//go:noescape
func CallOffClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/notifications try_OffClicked
//go:noescape
func TryOffClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_HasOnClicked
//go:noescape
func HasFuncHasOnClicked() js.Ref

//go:wasmimport plat/js/webext/notifications func_HasOnClicked
//go:noescape
func FuncHasOnClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_HasOnClicked
//go:noescape
func CallHasOnClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/notifications try_HasOnClicked
//go:noescape
func TryHasOnClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_OnClosed
//go:noescape
func HasFuncOnClosed() js.Ref

//go:wasmimport plat/js/webext/notifications func_OnClosed
//go:noescape
func FuncOnClosed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_OnClosed
//go:noescape
func CallOnClosed(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/notifications try_OnClosed
//go:noescape
func TryOnClosed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_OffClosed
//go:noescape
func HasFuncOffClosed() js.Ref

//go:wasmimport plat/js/webext/notifications func_OffClosed
//go:noescape
func FuncOffClosed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_OffClosed
//go:noescape
func CallOffClosed(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/notifications try_OffClosed
//go:noescape
func TryOffClosed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_HasOnClosed
//go:noescape
func HasFuncHasOnClosed() js.Ref

//go:wasmimport plat/js/webext/notifications func_HasOnClosed
//go:noescape
func FuncHasOnClosed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_HasOnClosed
//go:noescape
func CallHasOnClosed(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/notifications try_HasOnClosed
//go:noescape
func TryHasOnClosed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_OnPermissionLevelChanged
//go:noescape
func HasFuncOnPermissionLevelChanged() js.Ref

//go:wasmimport plat/js/webext/notifications func_OnPermissionLevelChanged
//go:noescape
func FuncOnPermissionLevelChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_OnPermissionLevelChanged
//go:noescape
func CallOnPermissionLevelChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/notifications try_OnPermissionLevelChanged
//go:noescape
func TryOnPermissionLevelChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_OffPermissionLevelChanged
//go:noescape
func HasFuncOffPermissionLevelChanged() js.Ref

//go:wasmimport plat/js/webext/notifications func_OffPermissionLevelChanged
//go:noescape
func FuncOffPermissionLevelChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_OffPermissionLevelChanged
//go:noescape
func CallOffPermissionLevelChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/notifications try_OffPermissionLevelChanged
//go:noescape
func TryOffPermissionLevelChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_HasOnPermissionLevelChanged
//go:noescape
func HasFuncHasOnPermissionLevelChanged() js.Ref

//go:wasmimport plat/js/webext/notifications func_HasOnPermissionLevelChanged
//go:noescape
func FuncHasOnPermissionLevelChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_HasOnPermissionLevelChanged
//go:noescape
func CallHasOnPermissionLevelChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/notifications try_HasOnPermissionLevelChanged
//go:noescape
func TryHasOnPermissionLevelChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_OnShowSettings
//go:noescape
func HasFuncOnShowSettings() js.Ref

//go:wasmimport plat/js/webext/notifications func_OnShowSettings
//go:noescape
func FuncOnShowSettings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_OnShowSettings
//go:noescape
func CallOnShowSettings(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/notifications try_OnShowSettings
//go:noescape
func TryOnShowSettings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_OffShowSettings
//go:noescape
func HasFuncOffShowSettings() js.Ref

//go:wasmimport plat/js/webext/notifications func_OffShowSettings
//go:noescape
func FuncOffShowSettings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_OffShowSettings
//go:noescape
func CallOffShowSettings(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/notifications try_OffShowSettings
//go:noescape
func TryOffShowSettings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_HasOnShowSettings
//go:noescape
func HasFuncHasOnShowSettings() js.Ref

//go:wasmimport plat/js/webext/notifications func_HasOnShowSettings
//go:noescape
func FuncHasOnShowSettings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_HasOnShowSettings
//go:noescape
func CallHasOnShowSettings(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/notifications try_HasOnShowSettings
//go:noescape
func TryHasOnShowSettings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/notifications has_Update
//go:noescape
func HasFuncUpdate() js.Ref

//go:wasmimport plat/js/webext/notifications func_Update
//go:noescape
func FuncUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications call_Update
//go:noescape
func CallUpdate(
	retPtr unsafe.Pointer,
	notificationId js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/notifications try_Update
//go:noescape
func TryUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	notificationId js.Ref,
	options unsafe.Pointer) (ok js.Ref)
