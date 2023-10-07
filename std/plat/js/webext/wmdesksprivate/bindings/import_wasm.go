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

//go:wasmimport plat/js/webext/wmdesksprivate store_Desk
//go:noescape
func DeskJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate load_Desk
//go:noescape
func DeskJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate constof_SavedDeskType
//go:noescape
func ConstOfSavedDeskType(str js.Ref) uint32

//go:wasmimport plat/js/webext/wmdesksprivate store_SavedDesk
//go:noescape
func SavedDeskJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate load_SavedDesk
//go:noescape
func SavedDeskJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate store_LaunchOptions
//go:noescape
func LaunchOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate load_LaunchOptions
//go:noescape
func LaunchOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate has_OnDeskAdded
//go:noescape
func HasFuncOnDeskAdded() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_OnDeskAdded
//go:noescape
func FuncOnDeskAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_OnDeskAdded
//go:noescape
func CallOnDeskAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate try_OnDeskAdded
//go:noescape
func TryOnDeskAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_OffDeskAdded
//go:noescape
func HasFuncOffDeskAdded() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_OffDeskAdded
//go:noescape
func FuncOffDeskAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_OffDeskAdded
//go:noescape
func CallOffDeskAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate try_OffDeskAdded
//go:noescape
func TryOffDeskAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_HasOnDeskAdded
//go:noescape
func HasFuncHasOnDeskAdded() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_HasOnDeskAdded
//go:noescape
func FuncHasOnDeskAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_HasOnDeskAdded
//go:noescape
func CallHasOnDeskAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate try_HasOnDeskAdded
//go:noescape
func TryHasOnDeskAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_OnDeskRemoved
//go:noescape
func HasFuncOnDeskRemoved() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_OnDeskRemoved
//go:noescape
func FuncOnDeskRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_OnDeskRemoved
//go:noescape
func CallOnDeskRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate try_OnDeskRemoved
//go:noescape
func TryOnDeskRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_OffDeskRemoved
//go:noescape
func HasFuncOffDeskRemoved() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_OffDeskRemoved
//go:noescape
func FuncOffDeskRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_OffDeskRemoved
//go:noescape
func CallOffDeskRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate try_OffDeskRemoved
//go:noescape
func TryOffDeskRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_HasOnDeskRemoved
//go:noescape
func HasFuncHasOnDeskRemoved() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_HasOnDeskRemoved
//go:noescape
func FuncHasOnDeskRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_HasOnDeskRemoved
//go:noescape
func CallHasOnDeskRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate try_HasOnDeskRemoved
//go:noescape
func TryHasOnDeskRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_OnDeskSwitched
//go:noescape
func HasFuncOnDeskSwitched() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_OnDeskSwitched
//go:noescape
func FuncOnDeskSwitched(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_OnDeskSwitched
//go:noescape
func CallOnDeskSwitched(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate try_OnDeskSwitched
//go:noescape
func TryOnDeskSwitched(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_OffDeskSwitched
//go:noescape
func HasFuncOffDeskSwitched() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_OffDeskSwitched
//go:noescape
func FuncOffDeskSwitched(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_OffDeskSwitched
//go:noescape
func CallOffDeskSwitched(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate try_OffDeskSwitched
//go:noescape
func TryOffDeskSwitched(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_HasOnDeskSwitched
//go:noescape
func HasFuncHasOnDeskSwitched() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_HasOnDeskSwitched
//go:noescape
func FuncHasOnDeskSwitched(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_HasOnDeskSwitched
//go:noescape
func CallHasOnDeskSwitched(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate try_HasOnDeskSwitched
//go:noescape
func TryHasOnDeskSwitched(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate store_RemoveDeskOptions
//go:noescape
func RemoveDeskOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate load_RemoveDeskOptions
//go:noescape
func RemoveDeskOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate store_WindowProperties
//go:noescape
func WindowPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate load_WindowProperties
//go:noescape
func WindowPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate has_DeleteSavedDesk
//go:noescape
func HasFuncDeleteSavedDesk() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_DeleteSavedDesk
//go:noescape
func FuncDeleteSavedDesk(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_DeleteSavedDesk
//go:noescape
func CallDeleteSavedDesk(
	retPtr unsafe.Pointer,
	savedDeskUuid js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate try_DeleteSavedDesk
//go:noescape
func TryDeleteSavedDesk(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	savedDeskUuid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_GetActiveDesk
//go:noescape
func HasFuncGetActiveDesk() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_GetActiveDesk
//go:noescape
func FuncGetActiveDesk(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_GetActiveDesk
//go:noescape
func CallGetActiveDesk(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate try_GetActiveDesk
//go:noescape
func TryGetActiveDesk(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_GetAllDesks
//go:noescape
func HasFuncGetAllDesks() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_GetAllDesks
//go:noescape
func FuncGetAllDesks(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_GetAllDesks
//go:noescape
func CallGetAllDesks(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate try_GetAllDesks
//go:noescape
func TryGetAllDesks(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_GetDeskByID
//go:noescape
func HasFuncGetDeskByID() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_GetDeskByID
//go:noescape
func FuncGetDeskByID(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_GetDeskByID
//go:noescape
func CallGetDeskByID(
	retPtr unsafe.Pointer,
	deskUuid js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate try_GetDeskByID
//go:noescape
func TryGetDeskByID(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deskUuid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_GetDeskTemplateJson
//go:noescape
func HasFuncGetDeskTemplateJson() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_GetDeskTemplateJson
//go:noescape
func FuncGetDeskTemplateJson(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_GetDeskTemplateJson
//go:noescape
func CallGetDeskTemplateJson(
	retPtr unsafe.Pointer,
	templateUuid js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate try_GetDeskTemplateJson
//go:noescape
func TryGetDeskTemplateJson(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	templateUuid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_GetSavedDesks
//go:noescape
func HasFuncGetSavedDesks() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_GetSavedDesks
//go:noescape
func FuncGetSavedDesks(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_GetSavedDesks
//go:noescape
func CallGetSavedDesks(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate try_GetSavedDesks
//go:noescape
func TryGetSavedDesks(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_LaunchDesk
//go:noescape
func HasFuncLaunchDesk() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_LaunchDesk
//go:noescape
func FuncLaunchDesk(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_LaunchDesk
//go:noescape
func CallLaunchDesk(
	retPtr unsafe.Pointer,
	launchOptions unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate try_LaunchDesk
//go:noescape
func TryLaunchDesk(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	launchOptions unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_RecallSavedDesk
//go:noescape
func HasFuncRecallSavedDesk() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_RecallSavedDesk
//go:noescape
func FuncRecallSavedDesk(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_RecallSavedDesk
//go:noescape
func CallRecallSavedDesk(
	retPtr unsafe.Pointer,
	savedDeskUuid js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate try_RecallSavedDesk
//go:noescape
func TryRecallSavedDesk(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	savedDeskUuid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_RemoveDesk
//go:noescape
func HasFuncRemoveDesk() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_RemoveDesk
//go:noescape
func FuncRemoveDesk(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_RemoveDesk
//go:noescape
func CallRemoveDesk(
	retPtr unsafe.Pointer,
	deskId js.Ref,
	removeDeskOptions unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate try_RemoveDesk
//go:noescape
func TryRemoveDesk(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deskId js.Ref,
	removeDeskOptions unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_SaveActiveDesk
//go:noescape
func HasFuncSaveActiveDesk() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_SaveActiveDesk
//go:noescape
func FuncSaveActiveDesk(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_SaveActiveDesk
//go:noescape
func CallSaveActiveDesk(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate try_SaveActiveDesk
//go:noescape
func TrySaveActiveDesk(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_SetWindowProperties
//go:noescape
func HasFuncSetWindowProperties() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_SetWindowProperties
//go:noescape
func FuncSetWindowProperties(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_SetWindowProperties
//go:noescape
func CallSetWindowProperties(
	retPtr unsafe.Pointer,
	windowId int32,
	windowProperties unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate try_SetWindowProperties
//go:noescape
func TrySetWindowProperties(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	windowId int32,
	windowProperties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate has_SwitchDesk
//go:noescape
func HasFuncSwitchDesk() js.Ref

//go:wasmimport plat/js/webext/wmdesksprivate func_SwitchDesk
//go:noescape
func FuncSwitchDesk(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/wmdesksprivate call_SwitchDesk
//go:noescape
func CallSwitchDesk(
	retPtr unsafe.Pointer,
	deskUuid js.Ref)

//go:wasmimport plat/js/webext/wmdesksprivate try_SwitchDesk
//go:noescape
func TrySwitchDesk(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deskUuid js.Ref) (ok js.Ref)
