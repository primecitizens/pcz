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

//go:wasmimport plat/js/webext/debugger store_Debuggee
//go:noescape
func DebuggeeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/debugger load_Debuggee
//go:noescape
func DebuggeeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/debugger constof_DetachReason
//go:noescape
func ConstOfDetachReason(str js.Ref) uint32

//go:wasmimport plat/js/webext/debugger constof_TargetInfoType
//go:noescape
func ConstOfTargetInfoType(str js.Ref) uint32

//go:wasmimport plat/js/webext/debugger store_TargetInfo
//go:noescape
func TargetInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/debugger load_TargetInfo
//go:noescape
func TargetInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/debugger has_Attach
//go:noescape
func HasFuncAttach() js.Ref

//go:wasmimport plat/js/webext/debugger func_Attach
//go:noescape
func FuncAttach(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/debugger call_Attach
//go:noescape
func CallAttach(
	retPtr unsafe.Pointer,
	target unsafe.Pointer,
	requiredVersion js.Ref)

//go:wasmimport plat/js/webext/debugger try_Attach
//go:noescape
func TryAttach(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	target unsafe.Pointer,
	requiredVersion js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/debugger has_Detach
//go:noescape
func HasFuncDetach() js.Ref

//go:wasmimport plat/js/webext/debugger func_Detach
//go:noescape
func FuncDetach(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/debugger call_Detach
//go:noescape
func CallDetach(
	retPtr unsafe.Pointer,
	target unsafe.Pointer)

//go:wasmimport plat/js/webext/debugger try_Detach
//go:noescape
func TryDetach(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	target unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/debugger has_GetTargets
//go:noescape
func HasFuncGetTargets() js.Ref

//go:wasmimport plat/js/webext/debugger func_GetTargets
//go:noescape
func FuncGetTargets(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/debugger call_GetTargets
//go:noescape
func CallGetTargets(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/debugger try_GetTargets
//go:noescape
func TryGetTargets(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/debugger has_OnDetach
//go:noescape
func HasFuncOnDetach() js.Ref

//go:wasmimport plat/js/webext/debugger func_OnDetach
//go:noescape
func FuncOnDetach(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/debugger call_OnDetach
//go:noescape
func CallOnDetach(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/debugger try_OnDetach
//go:noescape
func TryOnDetach(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/debugger has_OffDetach
//go:noescape
func HasFuncOffDetach() js.Ref

//go:wasmimport plat/js/webext/debugger func_OffDetach
//go:noescape
func FuncOffDetach(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/debugger call_OffDetach
//go:noescape
func CallOffDetach(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/debugger try_OffDetach
//go:noescape
func TryOffDetach(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/debugger has_HasOnDetach
//go:noescape
func HasFuncHasOnDetach() js.Ref

//go:wasmimport plat/js/webext/debugger func_HasOnDetach
//go:noescape
func FuncHasOnDetach(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/debugger call_HasOnDetach
//go:noescape
func CallHasOnDetach(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/debugger try_HasOnDetach
//go:noescape
func TryHasOnDetach(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/debugger has_OnEvent
//go:noescape
func HasFuncOnEvent() js.Ref

//go:wasmimport plat/js/webext/debugger func_OnEvent
//go:noescape
func FuncOnEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/debugger call_OnEvent
//go:noescape
func CallOnEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/debugger try_OnEvent
//go:noescape
func TryOnEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/debugger has_OffEvent
//go:noescape
func HasFuncOffEvent() js.Ref

//go:wasmimport plat/js/webext/debugger func_OffEvent
//go:noescape
func FuncOffEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/debugger call_OffEvent
//go:noescape
func CallOffEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/debugger try_OffEvent
//go:noescape
func TryOffEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/debugger has_HasOnEvent
//go:noescape
func HasFuncHasOnEvent() js.Ref

//go:wasmimport plat/js/webext/debugger func_HasOnEvent
//go:noescape
func FuncHasOnEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/debugger call_HasOnEvent
//go:noescape
func CallHasOnEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/debugger try_HasOnEvent
//go:noescape
func TryHasOnEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/debugger has_SendCommand
//go:noescape
func HasFuncSendCommand() js.Ref

//go:wasmimport plat/js/webext/debugger func_SendCommand
//go:noescape
func FuncSendCommand(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/debugger call_SendCommand
//go:noescape
func CallSendCommand(
	retPtr unsafe.Pointer,
	target unsafe.Pointer,
	method js.Ref,
	commandParams js.Ref)

//go:wasmimport plat/js/webext/debugger try_SendCommand
//go:noescape
func TrySendCommand(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	target unsafe.Pointer,
	method js.Ref,
	commandParams js.Ref) (ok js.Ref)
