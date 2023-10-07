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

//go:wasmimport plat/js/webext/windows constof_WindowState
//go:noescape
func ConstOfWindowState(str js.Ref) uint32

//go:wasmimport plat/js/webext/windows constof_CreateType
//go:noescape
func ConstOfCreateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/windows store_CreateArgCreateData
//go:noescape
func CreateArgCreateDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/windows load_CreateArgCreateData
//go:noescape
func CreateArgCreateDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/windows constof_WindowType
//go:noescape
func ConstOfWindowType(str js.Ref) uint32

//go:wasmimport plat/js/webext/windows store_Window
//go:noescape
func WindowJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/windows load_Window
//go:noescape
func WindowJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/windows store_QueryOptions
//go:noescape
func QueryOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/windows load_QueryOptions
//go:noescape
func QueryOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/windows store_UpdateArgUpdateInfo
//go:noescape
func UpdateArgUpdateInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/windows load_UpdateArgUpdateInfo
//go:noescape
func UpdateArgUpdateInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/windows get_WINDOW_ID_CURRENT
//go:noescape
func GetWINDOW_ID_CURRENT(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/windows set_WINDOW_ID_CURRENT
//go:noescape
func SetWINDOW_ID_CURRENT(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/windows get_WINDOW_ID_NONE
//go:noescape
func GetWINDOW_ID_NONE(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/windows set_WINDOW_ID_NONE
//go:noescape
func SetWINDOW_ID_NONE(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/windows has_Create
//go:noescape
func HasFuncCreate() js.Ref

//go:wasmimport plat/js/webext/windows func_Create
//go:noescape
func FuncCreate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_Create
//go:noescape
func CallCreate(
	retPtr unsafe.Pointer,
	createData unsafe.Pointer)

//go:wasmimport plat/js/webext/windows try_Create
//go:noescape
func TryCreate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	createData unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_Get
//go:noescape
func HasFuncGet() js.Ref

//go:wasmimport plat/js/webext/windows func_Get
//go:noescape
func FuncGet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_Get
//go:noescape
func CallGet(
	retPtr unsafe.Pointer,
	windowId float64,
	queryOptions unsafe.Pointer)

//go:wasmimport plat/js/webext/windows try_Get
//go:noescape
func TryGet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	windowId float64,
	queryOptions unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_GetAll
//go:noescape
func HasFuncGetAll() js.Ref

//go:wasmimport plat/js/webext/windows func_GetAll
//go:noescape
func FuncGetAll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_GetAll
//go:noescape
func CallGetAll(
	retPtr unsafe.Pointer,
	queryOptions unsafe.Pointer)

//go:wasmimport plat/js/webext/windows try_GetAll
//go:noescape
func TryGetAll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	queryOptions unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_GetCurrent
//go:noescape
func HasFuncGetCurrent() js.Ref

//go:wasmimport plat/js/webext/windows func_GetCurrent
//go:noescape
func FuncGetCurrent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_GetCurrent
//go:noescape
func CallGetCurrent(
	retPtr unsafe.Pointer,
	queryOptions unsafe.Pointer)

//go:wasmimport plat/js/webext/windows try_GetCurrent
//go:noescape
func TryGetCurrent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	queryOptions unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_GetLastFocused
//go:noescape
func HasFuncGetLastFocused() js.Ref

//go:wasmimport plat/js/webext/windows func_GetLastFocused
//go:noescape
func FuncGetLastFocused(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_GetLastFocused
//go:noescape
func CallGetLastFocused(
	retPtr unsafe.Pointer,
	queryOptions unsafe.Pointer)

//go:wasmimport plat/js/webext/windows try_GetLastFocused
//go:noescape
func TryGetLastFocused(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	queryOptions unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_OnBoundsChanged
//go:noescape
func HasFuncOnBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/windows func_OnBoundsChanged
//go:noescape
func FuncOnBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_OnBoundsChanged
//go:noescape
func CallOnBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/windows try_OnBoundsChanged
//go:noescape
func TryOnBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_OffBoundsChanged
//go:noescape
func HasFuncOffBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/windows func_OffBoundsChanged
//go:noescape
func FuncOffBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_OffBoundsChanged
//go:noescape
func CallOffBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/windows try_OffBoundsChanged
//go:noescape
func TryOffBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_HasOnBoundsChanged
//go:noescape
func HasFuncHasOnBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/windows func_HasOnBoundsChanged
//go:noescape
func FuncHasOnBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_HasOnBoundsChanged
//go:noescape
func CallHasOnBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/windows try_HasOnBoundsChanged
//go:noescape
func TryHasOnBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_OnCreated
//go:noescape
func HasFuncOnCreated() js.Ref

//go:wasmimport plat/js/webext/windows func_OnCreated
//go:noescape
func FuncOnCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_OnCreated
//go:noescape
func CallOnCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/windows try_OnCreated
//go:noescape
func TryOnCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_OffCreated
//go:noescape
func HasFuncOffCreated() js.Ref

//go:wasmimport plat/js/webext/windows func_OffCreated
//go:noescape
func FuncOffCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_OffCreated
//go:noescape
func CallOffCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/windows try_OffCreated
//go:noescape
func TryOffCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_HasOnCreated
//go:noescape
func HasFuncHasOnCreated() js.Ref

//go:wasmimport plat/js/webext/windows func_HasOnCreated
//go:noescape
func FuncHasOnCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_HasOnCreated
//go:noescape
func CallHasOnCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/windows try_HasOnCreated
//go:noescape
func TryHasOnCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_OnFocusChanged
//go:noescape
func HasFuncOnFocusChanged() js.Ref

//go:wasmimport plat/js/webext/windows func_OnFocusChanged
//go:noescape
func FuncOnFocusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_OnFocusChanged
//go:noescape
func CallOnFocusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/windows try_OnFocusChanged
//go:noescape
func TryOnFocusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_OffFocusChanged
//go:noescape
func HasFuncOffFocusChanged() js.Ref

//go:wasmimport plat/js/webext/windows func_OffFocusChanged
//go:noescape
func FuncOffFocusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_OffFocusChanged
//go:noescape
func CallOffFocusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/windows try_OffFocusChanged
//go:noescape
func TryOffFocusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_HasOnFocusChanged
//go:noescape
func HasFuncHasOnFocusChanged() js.Ref

//go:wasmimport plat/js/webext/windows func_HasOnFocusChanged
//go:noescape
func FuncHasOnFocusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_HasOnFocusChanged
//go:noescape
func CallHasOnFocusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/windows try_HasOnFocusChanged
//go:noescape
func TryHasOnFocusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_OnRemoved
//go:noescape
func HasFuncOnRemoved() js.Ref

//go:wasmimport plat/js/webext/windows func_OnRemoved
//go:noescape
func FuncOnRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_OnRemoved
//go:noescape
func CallOnRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/windows try_OnRemoved
//go:noescape
func TryOnRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_OffRemoved
//go:noescape
func HasFuncOffRemoved() js.Ref

//go:wasmimport plat/js/webext/windows func_OffRemoved
//go:noescape
func FuncOffRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_OffRemoved
//go:noescape
func CallOffRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/windows try_OffRemoved
//go:noescape
func TryOffRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_HasOnRemoved
//go:noescape
func HasFuncHasOnRemoved() js.Ref

//go:wasmimport plat/js/webext/windows func_HasOnRemoved
//go:noescape
func FuncHasOnRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_HasOnRemoved
//go:noescape
func CallHasOnRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/windows try_HasOnRemoved
//go:noescape
func TryHasOnRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_Remove
//go:noescape
func HasFuncRemove() js.Ref

//go:wasmimport plat/js/webext/windows func_Remove
//go:noescape
func FuncRemove(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_Remove
//go:noescape
func CallRemove(
	retPtr unsafe.Pointer,
	windowId float64)

//go:wasmimport plat/js/webext/windows try_Remove
//go:noescape
func TryRemove(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	windowId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/windows has_Update
//go:noescape
func HasFuncUpdate() js.Ref

//go:wasmimport plat/js/webext/windows func_Update
//go:noescape
func FuncUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/windows call_Update
//go:noescape
func CallUpdate(
	retPtr unsafe.Pointer,
	windowId float64,
	updateInfo unsafe.Pointer)

//go:wasmimport plat/js/webext/windows try_Update
//go:noescape
func TryUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	windowId float64,
	updateInfo unsafe.Pointer) (ok js.Ref)
