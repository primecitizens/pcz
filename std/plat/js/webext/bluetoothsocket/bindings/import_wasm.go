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

//go:wasmimport plat/js/webext/bluetoothsocket constof_AcceptError
//go:noescape
func ConstOfAcceptError(str js.Ref) uint32

//go:wasmimport plat/js/webext/bluetoothsocket store_AcceptErrorInfo
//go:noescape
func AcceptErrorInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket load_AcceptErrorInfo
//go:noescape
func AcceptErrorInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket store_AcceptInfo
//go:noescape
func AcceptInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket load_AcceptInfo
//go:noescape
func AcceptInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket store_CreateInfo
//go:noescape
func CreateInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket load_CreateInfo
//go:noescape
func CreateInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket store_SocketInfo
//go:noescape
func SocketInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket load_SocketInfo
//go:noescape
func SocketInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket store_ListenOptions
//go:noescape
func ListenOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket load_ListenOptions
//go:noescape
func ListenOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket constof_ReceiveError
//go:noescape
func ConstOfReceiveError(str js.Ref) uint32

//go:wasmimport plat/js/webext/bluetoothsocket store_ReceiveErrorInfo
//go:noescape
func ReceiveErrorInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket load_ReceiveErrorInfo
//go:noescape
func ReceiveErrorInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket store_ReceiveInfo
//go:noescape
func ReceiveInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket load_ReceiveInfo
//go:noescape
func ReceiveInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket store_SocketProperties
//go:noescape
func SocketPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket load_SocketProperties
//go:noescape
func SocketPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket has_Close
//go:noescape
func HasFuncClose() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_Close
//go:noescape
func FuncClose(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_Close
//go:noescape
func CallClose(
	retPtr unsafe.Pointer,
	socketId int32)

//go:wasmimport plat/js/webext/bluetoothsocket try_Close
//go:noescape
func TryClose(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_Connect
//go:noescape
func HasFuncConnect() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_Connect
//go:noescape
func FuncConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_Connect
//go:noescape
func CallConnect(
	retPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	uuid js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket try_Connect
//go:noescape
func TryConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	uuid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_Create
//go:noescape
func HasFuncCreate() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_Create
//go:noescape
func FuncCreate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_Create
//go:noescape
func CallCreate(
	retPtr unsafe.Pointer,
	properties unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket try_Create
//go:noescape
func TryCreate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	properties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_Disconnect
//go:noescape
func HasFuncDisconnect() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_Disconnect
//go:noescape
func FuncDisconnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_Disconnect
//go:noescape
func CallDisconnect(
	retPtr unsafe.Pointer,
	socketId int32)

//go:wasmimport plat/js/webext/bluetoothsocket try_Disconnect
//go:noescape
func TryDisconnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_GetInfo
//go:noescape
func HasFuncGetInfo() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_GetInfo
//go:noescape
func FuncGetInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_GetInfo
//go:noescape
func CallGetInfo(
	retPtr unsafe.Pointer,
	socketId int32)

//go:wasmimport plat/js/webext/bluetoothsocket try_GetInfo
//go:noescape
func TryGetInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_GetSockets
//go:noescape
func HasFuncGetSockets() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_GetSockets
//go:noescape
func FuncGetSockets(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_GetSockets
//go:noescape
func CallGetSockets(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket try_GetSockets
//go:noescape
func TryGetSockets(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_ListenUsingL2cap
//go:noescape
func HasFuncListenUsingL2cap() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_ListenUsingL2cap
//go:noescape
func FuncListenUsingL2cap(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_ListenUsingL2cap
//go:noescape
func CallListenUsingL2cap(
	retPtr unsafe.Pointer,
	socketId int32,
	uuid js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket try_ListenUsingL2cap
//go:noescape
func TryListenUsingL2cap(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	uuid js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_ListenUsingRfcomm
//go:noescape
func HasFuncListenUsingRfcomm() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_ListenUsingRfcomm
//go:noescape
func FuncListenUsingRfcomm(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_ListenUsingRfcomm
//go:noescape
func CallListenUsingRfcomm(
	retPtr unsafe.Pointer,
	socketId int32,
	uuid js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket try_ListenUsingRfcomm
//go:noescape
func TryListenUsingRfcomm(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	uuid js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_OnAccept
//go:noescape
func HasFuncOnAccept() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_OnAccept
//go:noescape
func FuncOnAccept(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_OnAccept
//go:noescape
func CallOnAccept(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket try_OnAccept
//go:noescape
func TryOnAccept(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_OffAccept
//go:noescape
func HasFuncOffAccept() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_OffAccept
//go:noescape
func FuncOffAccept(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_OffAccept
//go:noescape
func CallOffAccept(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket try_OffAccept
//go:noescape
func TryOffAccept(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_HasOnAccept
//go:noescape
func HasFuncHasOnAccept() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_HasOnAccept
//go:noescape
func FuncHasOnAccept(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_HasOnAccept
//go:noescape
func CallHasOnAccept(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket try_HasOnAccept
//go:noescape
func TryHasOnAccept(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_OnAcceptError
//go:noescape
func HasFuncOnAcceptError() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_OnAcceptError
//go:noescape
func FuncOnAcceptError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_OnAcceptError
//go:noescape
func CallOnAcceptError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket try_OnAcceptError
//go:noescape
func TryOnAcceptError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_OffAcceptError
//go:noescape
func HasFuncOffAcceptError() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_OffAcceptError
//go:noescape
func FuncOffAcceptError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_OffAcceptError
//go:noescape
func CallOffAcceptError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket try_OffAcceptError
//go:noescape
func TryOffAcceptError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_HasOnAcceptError
//go:noescape
func HasFuncHasOnAcceptError() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_HasOnAcceptError
//go:noescape
func FuncHasOnAcceptError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_HasOnAcceptError
//go:noescape
func CallHasOnAcceptError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket try_HasOnAcceptError
//go:noescape
func TryHasOnAcceptError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_OnReceive
//go:noescape
func HasFuncOnReceive() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_OnReceive
//go:noescape
func FuncOnReceive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_OnReceive
//go:noescape
func CallOnReceive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket try_OnReceive
//go:noescape
func TryOnReceive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_OffReceive
//go:noescape
func HasFuncOffReceive() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_OffReceive
//go:noescape
func FuncOffReceive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_OffReceive
//go:noescape
func CallOffReceive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket try_OffReceive
//go:noescape
func TryOffReceive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_HasOnReceive
//go:noescape
func HasFuncHasOnReceive() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_HasOnReceive
//go:noescape
func FuncHasOnReceive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_HasOnReceive
//go:noescape
func CallHasOnReceive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket try_HasOnReceive
//go:noescape
func TryHasOnReceive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_OnReceiveError
//go:noescape
func HasFuncOnReceiveError() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_OnReceiveError
//go:noescape
func FuncOnReceiveError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_OnReceiveError
//go:noescape
func CallOnReceiveError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket try_OnReceiveError
//go:noescape
func TryOnReceiveError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_OffReceiveError
//go:noescape
func HasFuncOffReceiveError() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_OffReceiveError
//go:noescape
func FuncOffReceiveError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_OffReceiveError
//go:noescape
func CallOffReceiveError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket try_OffReceiveError
//go:noescape
func TryOffReceiveError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_HasOnReceiveError
//go:noescape
func HasFuncHasOnReceiveError() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_HasOnReceiveError
//go:noescape
func FuncHasOnReceiveError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_HasOnReceiveError
//go:noescape
func CallHasOnReceiveError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket try_HasOnReceiveError
//go:noescape
func TryHasOnReceiveError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_Send
//go:noescape
func HasFuncSend() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_Send
//go:noescape
func FuncSend(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_Send
//go:noescape
func CallSend(
	retPtr unsafe.Pointer,
	socketId int32,
	data js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket try_Send
//go:noescape
func TrySend(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_SetPaused
//go:noescape
func HasFuncSetPaused() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_SetPaused
//go:noescape
func FuncSetPaused(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_SetPaused
//go:noescape
func CallSetPaused(
	retPtr unsafe.Pointer,
	socketId int32,
	paused js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket try_SetPaused
//go:noescape
func TrySetPaused(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	paused js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothsocket has_Update
//go:noescape
func HasFuncUpdate() js.Ref

//go:wasmimport plat/js/webext/bluetoothsocket func_Update
//go:noescape
func FuncUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket call_Update
//go:noescape
func CallUpdate(
	retPtr unsafe.Pointer,
	socketId int32,
	properties unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothsocket try_Update
//go:noescape
func TryUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	properties unsafe.Pointer) (ok js.Ref)
