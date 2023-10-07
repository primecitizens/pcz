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

//go:wasmimport plat/js/webext/processes store_Cache
//go:noescape
func CacheJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/processes load_Cache
//go:noescape
func CacheJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/processes constof_ProcessType
//go:noescape
func ConstOfProcessType(str js.Ref) uint32

//go:wasmimport plat/js/webext/processes store_TaskInfo
//go:noescape
func TaskInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/processes load_TaskInfo
//go:noescape
func TaskInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/processes store_Process
//go:noescape
func ProcessJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/processes load_Process
//go:noescape
func ProcessJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/processes has_GetProcessIdForTab
//go:noescape
func HasFuncGetProcessIdForTab() js.Ref

//go:wasmimport plat/js/webext/processes func_GetProcessIdForTab
//go:noescape
func FuncGetProcessIdForTab(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_GetProcessIdForTab
//go:noescape
func CallGetProcessIdForTab(
	retPtr unsafe.Pointer,
	tabId int32)

//go:wasmimport plat/js/webext/processes try_GetProcessIdForTab
//go:noescape
func TryGetProcessIdForTab(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_GetProcessInfo
//go:noescape
func HasFuncGetProcessInfo() js.Ref

//go:wasmimport plat/js/webext/processes func_GetProcessInfo
//go:noescape
func FuncGetProcessInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_GetProcessInfo
//go:noescape
func CallGetProcessInfo(
	retPtr unsafe.Pointer,
	processIds js.Ref,
	includeMemory js.Ref)

//go:wasmimport plat/js/webext/processes try_GetProcessInfo
//go:noescape
func TryGetProcessInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	processIds js.Ref,
	includeMemory js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_OnCreated
//go:noescape
func HasFuncOnCreated() js.Ref

//go:wasmimport plat/js/webext/processes func_OnCreated
//go:noescape
func FuncOnCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_OnCreated
//go:noescape
func CallOnCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/processes try_OnCreated
//go:noescape
func TryOnCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_OffCreated
//go:noescape
func HasFuncOffCreated() js.Ref

//go:wasmimport plat/js/webext/processes func_OffCreated
//go:noescape
func FuncOffCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_OffCreated
//go:noescape
func CallOffCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/processes try_OffCreated
//go:noescape
func TryOffCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_HasOnCreated
//go:noescape
func HasFuncHasOnCreated() js.Ref

//go:wasmimport plat/js/webext/processes func_HasOnCreated
//go:noescape
func FuncHasOnCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_HasOnCreated
//go:noescape
func CallHasOnCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/processes try_HasOnCreated
//go:noescape
func TryHasOnCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_OnExited
//go:noescape
func HasFuncOnExited() js.Ref

//go:wasmimport plat/js/webext/processes func_OnExited
//go:noescape
func FuncOnExited(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_OnExited
//go:noescape
func CallOnExited(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/processes try_OnExited
//go:noescape
func TryOnExited(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_OffExited
//go:noescape
func HasFuncOffExited() js.Ref

//go:wasmimport plat/js/webext/processes func_OffExited
//go:noescape
func FuncOffExited(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_OffExited
//go:noescape
func CallOffExited(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/processes try_OffExited
//go:noescape
func TryOffExited(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_HasOnExited
//go:noescape
func HasFuncHasOnExited() js.Ref

//go:wasmimport plat/js/webext/processes func_HasOnExited
//go:noescape
func FuncHasOnExited(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_HasOnExited
//go:noescape
func CallHasOnExited(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/processes try_HasOnExited
//go:noescape
func TryHasOnExited(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_OnUnresponsive
//go:noescape
func HasFuncOnUnresponsive() js.Ref

//go:wasmimport plat/js/webext/processes func_OnUnresponsive
//go:noescape
func FuncOnUnresponsive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_OnUnresponsive
//go:noescape
func CallOnUnresponsive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/processes try_OnUnresponsive
//go:noescape
func TryOnUnresponsive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_OffUnresponsive
//go:noescape
func HasFuncOffUnresponsive() js.Ref

//go:wasmimport plat/js/webext/processes func_OffUnresponsive
//go:noescape
func FuncOffUnresponsive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_OffUnresponsive
//go:noescape
func CallOffUnresponsive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/processes try_OffUnresponsive
//go:noescape
func TryOffUnresponsive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_HasOnUnresponsive
//go:noescape
func HasFuncHasOnUnresponsive() js.Ref

//go:wasmimport plat/js/webext/processes func_HasOnUnresponsive
//go:noescape
func FuncHasOnUnresponsive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_HasOnUnresponsive
//go:noescape
func CallHasOnUnresponsive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/processes try_HasOnUnresponsive
//go:noescape
func TryHasOnUnresponsive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_OnUpdated
//go:noescape
func HasFuncOnUpdated() js.Ref

//go:wasmimport plat/js/webext/processes func_OnUpdated
//go:noescape
func FuncOnUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_OnUpdated
//go:noescape
func CallOnUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/processes try_OnUpdated
//go:noescape
func TryOnUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_OffUpdated
//go:noescape
func HasFuncOffUpdated() js.Ref

//go:wasmimport plat/js/webext/processes func_OffUpdated
//go:noescape
func FuncOffUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_OffUpdated
//go:noescape
func CallOffUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/processes try_OffUpdated
//go:noescape
func TryOffUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_HasOnUpdated
//go:noescape
func HasFuncHasOnUpdated() js.Ref

//go:wasmimport plat/js/webext/processes func_HasOnUpdated
//go:noescape
func FuncHasOnUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_HasOnUpdated
//go:noescape
func CallHasOnUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/processes try_HasOnUpdated
//go:noescape
func TryHasOnUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_OnUpdatedWithMemory
//go:noescape
func HasFuncOnUpdatedWithMemory() js.Ref

//go:wasmimport plat/js/webext/processes func_OnUpdatedWithMemory
//go:noescape
func FuncOnUpdatedWithMemory(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_OnUpdatedWithMemory
//go:noescape
func CallOnUpdatedWithMemory(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/processes try_OnUpdatedWithMemory
//go:noescape
func TryOnUpdatedWithMemory(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_OffUpdatedWithMemory
//go:noescape
func HasFuncOffUpdatedWithMemory() js.Ref

//go:wasmimport plat/js/webext/processes func_OffUpdatedWithMemory
//go:noescape
func FuncOffUpdatedWithMemory(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_OffUpdatedWithMemory
//go:noescape
func CallOffUpdatedWithMemory(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/processes try_OffUpdatedWithMemory
//go:noescape
func TryOffUpdatedWithMemory(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_HasOnUpdatedWithMemory
//go:noescape
func HasFuncHasOnUpdatedWithMemory() js.Ref

//go:wasmimport plat/js/webext/processes func_HasOnUpdatedWithMemory
//go:noescape
func FuncHasOnUpdatedWithMemory(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_HasOnUpdatedWithMemory
//go:noescape
func CallHasOnUpdatedWithMemory(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/processes try_HasOnUpdatedWithMemory
//go:noescape
func TryHasOnUpdatedWithMemory(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/processes has_Terminate
//go:noescape
func HasFuncTerminate() js.Ref

//go:wasmimport plat/js/webext/processes func_Terminate
//go:noescape
func FuncTerminate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/processes call_Terminate
//go:noescape
func CallTerminate(
	retPtr unsafe.Pointer,
	processId int32)

//go:wasmimport plat/js/webext/processes try_Terminate
//go:noescape
func TryTerminate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	processId int32) (ok js.Ref)
