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

//go:wasmimport plat/js/webext/serial constof_DataBits
//go:noescape
func ConstOfDataBits(str js.Ref) uint32

//go:wasmimport plat/js/webext/serial constof_ParityBit
//go:noescape
func ConstOfParityBit(str js.Ref) uint32

//go:wasmimport plat/js/webext/serial constof_StopBits
//go:noescape
func ConstOfStopBits(str js.Ref) uint32

//go:wasmimport plat/js/webext/serial store_ConnectionInfo
//go:noescape
func ConnectionInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/serial load_ConnectionInfo
//go:noescape
func ConnectionInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/serial store_ConnectionOptions
//go:noescape
func ConnectionOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/serial load_ConnectionOptions
//go:noescape
func ConnectionOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/serial store_DeviceControlSignals
//go:noescape
func DeviceControlSignalsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/serial load_DeviceControlSignals
//go:noescape
func DeviceControlSignalsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/serial store_DeviceInfo
//go:noescape
func DeviceInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/serial load_DeviceInfo
//go:noescape
func DeviceInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/serial store_HostControlSignals
//go:noescape
func HostControlSignalsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/serial load_HostControlSignals
//go:noescape
func HostControlSignalsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/serial constof_ReceiveError
//go:noescape
func ConstOfReceiveError(str js.Ref) uint32

//go:wasmimport plat/js/webext/serial store_ReceiveErrorInfo
//go:noescape
func ReceiveErrorInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/serial load_ReceiveErrorInfo
//go:noescape
func ReceiveErrorInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/serial store_ReceiveInfo
//go:noescape
func ReceiveInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/serial load_ReceiveInfo
//go:noescape
func ReceiveInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/serial constof_SendError
//go:noescape
func ConstOfSendError(str js.Ref) uint32

//go:wasmimport plat/js/webext/serial store_SendInfo
//go:noescape
func SendInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/serial load_SendInfo
//go:noescape
func SendInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/serial has_ClearBreak
//go:noescape
func HasFuncClearBreak() js.Ref

//go:wasmimport plat/js/webext/serial func_ClearBreak
//go:noescape
func FuncClearBreak(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_ClearBreak
//go:noescape
func CallClearBreak(
	retPtr unsafe.Pointer,
	connectionId int32)

//go:wasmimport plat/js/webext/serial try_ClearBreak
//go:noescape
func TryClearBreak(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	connectionId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_Connect
//go:noescape
func HasFuncConnect() js.Ref

//go:wasmimport plat/js/webext/serial func_Connect
//go:noescape
func FuncConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_Connect
//go:noescape
func CallConnect(
	retPtr unsafe.Pointer,
	path js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/serial try_Connect
//go:noescape
func TryConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_Disconnect
//go:noescape
func HasFuncDisconnect() js.Ref

//go:wasmimport plat/js/webext/serial func_Disconnect
//go:noescape
func FuncDisconnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_Disconnect
//go:noescape
func CallDisconnect(
	retPtr unsafe.Pointer,
	connectionId int32)

//go:wasmimport plat/js/webext/serial try_Disconnect
//go:noescape
func TryDisconnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	connectionId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_Flush
//go:noescape
func HasFuncFlush() js.Ref

//go:wasmimport plat/js/webext/serial func_Flush
//go:noescape
func FuncFlush(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_Flush
//go:noescape
func CallFlush(
	retPtr unsafe.Pointer,
	connectionId int32)

//go:wasmimport plat/js/webext/serial try_Flush
//go:noescape
func TryFlush(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	connectionId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_GetConnections
//go:noescape
func HasFuncGetConnections() js.Ref

//go:wasmimport plat/js/webext/serial func_GetConnections
//go:noescape
func FuncGetConnections(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_GetConnections
//go:noescape
func CallGetConnections(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/serial try_GetConnections
//go:noescape
func TryGetConnections(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_GetControlSignals
//go:noescape
func HasFuncGetControlSignals() js.Ref

//go:wasmimport plat/js/webext/serial func_GetControlSignals
//go:noescape
func FuncGetControlSignals(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_GetControlSignals
//go:noescape
func CallGetControlSignals(
	retPtr unsafe.Pointer,
	connectionId int32)

//go:wasmimport plat/js/webext/serial try_GetControlSignals
//go:noescape
func TryGetControlSignals(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	connectionId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_GetDevices
//go:noescape
func HasFuncGetDevices() js.Ref

//go:wasmimport plat/js/webext/serial func_GetDevices
//go:noescape
func FuncGetDevices(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_GetDevices
//go:noescape
func CallGetDevices(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/serial try_GetDevices
//go:noescape
func TryGetDevices(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_GetInfo
//go:noescape
func HasFuncGetInfo() js.Ref

//go:wasmimport plat/js/webext/serial func_GetInfo
//go:noescape
func FuncGetInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_GetInfo
//go:noescape
func CallGetInfo(
	retPtr unsafe.Pointer,
	connectionId int32)

//go:wasmimport plat/js/webext/serial try_GetInfo
//go:noescape
func TryGetInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	connectionId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_OnReceive
//go:noescape
func HasFuncOnReceive() js.Ref

//go:wasmimport plat/js/webext/serial func_OnReceive
//go:noescape
func FuncOnReceive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_OnReceive
//go:noescape
func CallOnReceive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/serial try_OnReceive
//go:noescape
func TryOnReceive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_OffReceive
//go:noescape
func HasFuncOffReceive() js.Ref

//go:wasmimport plat/js/webext/serial func_OffReceive
//go:noescape
func FuncOffReceive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_OffReceive
//go:noescape
func CallOffReceive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/serial try_OffReceive
//go:noescape
func TryOffReceive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_HasOnReceive
//go:noescape
func HasFuncHasOnReceive() js.Ref

//go:wasmimport plat/js/webext/serial func_HasOnReceive
//go:noescape
func FuncHasOnReceive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_HasOnReceive
//go:noescape
func CallHasOnReceive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/serial try_HasOnReceive
//go:noescape
func TryHasOnReceive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_OnReceiveError
//go:noescape
func HasFuncOnReceiveError() js.Ref

//go:wasmimport plat/js/webext/serial func_OnReceiveError
//go:noescape
func FuncOnReceiveError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_OnReceiveError
//go:noescape
func CallOnReceiveError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/serial try_OnReceiveError
//go:noescape
func TryOnReceiveError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_OffReceiveError
//go:noescape
func HasFuncOffReceiveError() js.Ref

//go:wasmimport plat/js/webext/serial func_OffReceiveError
//go:noescape
func FuncOffReceiveError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_OffReceiveError
//go:noescape
func CallOffReceiveError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/serial try_OffReceiveError
//go:noescape
func TryOffReceiveError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_HasOnReceiveError
//go:noescape
func HasFuncHasOnReceiveError() js.Ref

//go:wasmimport plat/js/webext/serial func_HasOnReceiveError
//go:noescape
func FuncHasOnReceiveError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_HasOnReceiveError
//go:noescape
func CallHasOnReceiveError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/serial try_HasOnReceiveError
//go:noescape
func TryHasOnReceiveError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_Send
//go:noescape
func HasFuncSend() js.Ref

//go:wasmimport plat/js/webext/serial func_Send
//go:noescape
func FuncSend(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_Send
//go:noescape
func CallSend(
	retPtr unsafe.Pointer,
	connectionId int32,
	data js.Ref)

//go:wasmimport plat/js/webext/serial try_Send
//go:noescape
func TrySend(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	connectionId int32,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_SetBreak
//go:noescape
func HasFuncSetBreak() js.Ref

//go:wasmimport plat/js/webext/serial func_SetBreak
//go:noescape
func FuncSetBreak(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_SetBreak
//go:noescape
func CallSetBreak(
	retPtr unsafe.Pointer,
	connectionId int32)

//go:wasmimport plat/js/webext/serial try_SetBreak
//go:noescape
func TrySetBreak(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	connectionId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_SetControlSignals
//go:noescape
func HasFuncSetControlSignals() js.Ref

//go:wasmimport plat/js/webext/serial func_SetControlSignals
//go:noescape
func FuncSetControlSignals(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_SetControlSignals
//go:noescape
func CallSetControlSignals(
	retPtr unsafe.Pointer,
	connectionId int32,
	signals unsafe.Pointer)

//go:wasmimport plat/js/webext/serial try_SetControlSignals
//go:noescape
func TrySetControlSignals(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	connectionId int32,
	signals unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_SetPaused
//go:noescape
func HasFuncSetPaused() js.Ref

//go:wasmimport plat/js/webext/serial func_SetPaused
//go:noescape
func FuncSetPaused(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_SetPaused
//go:noescape
func CallSetPaused(
	retPtr unsafe.Pointer,
	connectionId int32,
	paused js.Ref)

//go:wasmimport plat/js/webext/serial try_SetPaused
//go:noescape
func TrySetPaused(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	connectionId int32,
	paused js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/serial has_Update
//go:noescape
func HasFuncUpdate() js.Ref

//go:wasmimport plat/js/webext/serial func_Update
//go:noescape
func FuncUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/serial call_Update
//go:noescape
func CallUpdate(
	retPtr unsafe.Pointer,
	connectionId int32,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/serial try_Update
//go:noescape
func TryUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	connectionId int32,
	options unsafe.Pointer) (ok js.Ref)
