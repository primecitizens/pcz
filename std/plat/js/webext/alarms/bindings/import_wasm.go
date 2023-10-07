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

//go:wasmimport plat/js/webext/alarms store_Alarm
//go:noescape
func AlarmJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/alarms load_Alarm
//go:noescape
func AlarmJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/alarms store_AlarmCreateInfo
//go:noescape
func AlarmCreateInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/alarms load_AlarmCreateInfo
//go:noescape
func AlarmCreateInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/alarms has_Clear
//go:noescape
func HasFuncClear() js.Ref

//go:wasmimport plat/js/webext/alarms func_Clear
//go:noescape
func FuncClear(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/alarms call_Clear
//go:noescape
func CallClear(
	retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/webext/alarms try_Clear
//go:noescape
func TryClear(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/alarms has_ClearAll
//go:noescape
func HasFuncClearAll() js.Ref

//go:wasmimport plat/js/webext/alarms func_ClearAll
//go:noescape
func FuncClearAll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/alarms call_ClearAll
//go:noescape
func CallClearAll(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/alarms try_ClearAll
//go:noescape
func TryClearAll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/alarms has_Create
//go:noescape
func HasFuncCreate() js.Ref

//go:wasmimport plat/js/webext/alarms func_Create
//go:noescape
func FuncCreate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/alarms call_Create
//go:noescape
func CallCreate(
	retPtr unsafe.Pointer,
	name js.Ref,
	alarmInfo unsafe.Pointer)

//go:wasmimport plat/js/webext/alarms try_Create
//go:noescape
func TryCreate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	alarmInfo unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/alarms has_Get
//go:noescape
func HasFuncGet() js.Ref

//go:wasmimport plat/js/webext/alarms func_Get
//go:noescape
func FuncGet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/alarms call_Get
//go:noescape
func CallGet(
	retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/webext/alarms try_Get
//go:noescape
func TryGet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/alarms has_GetAll
//go:noescape
func HasFuncGetAll() js.Ref

//go:wasmimport plat/js/webext/alarms func_GetAll
//go:noescape
func FuncGetAll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/alarms call_GetAll
//go:noescape
func CallGetAll(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/alarms try_GetAll
//go:noescape
func TryGetAll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/alarms has_OnAlarm
//go:noescape
func HasFuncOnAlarm() js.Ref

//go:wasmimport plat/js/webext/alarms func_OnAlarm
//go:noescape
func FuncOnAlarm(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/alarms call_OnAlarm
//go:noescape
func CallOnAlarm(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/alarms try_OnAlarm
//go:noescape
func TryOnAlarm(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/alarms has_OffAlarm
//go:noescape
func HasFuncOffAlarm() js.Ref

//go:wasmimport plat/js/webext/alarms func_OffAlarm
//go:noescape
func FuncOffAlarm(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/alarms call_OffAlarm
//go:noescape
func CallOffAlarm(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/alarms try_OffAlarm
//go:noescape
func TryOffAlarm(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/alarms has_HasOnAlarm
//go:noescape
func HasFuncHasOnAlarm() js.Ref

//go:wasmimport plat/js/webext/alarms func_HasOnAlarm
//go:noescape
func FuncHasOnAlarm(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/alarms call_HasOnAlarm
//go:noescape
func CallHasOnAlarm(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/alarms try_HasOnAlarm
//go:noescape
func TryHasOnAlarm(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
