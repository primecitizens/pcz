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

//go:wasmimport plat/js/webext/sockets/tcpserver store_AcceptErrorInfo
//go:noescape
func AcceptErrorInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver load_AcceptErrorInfo
//go:noescape
func AcceptErrorInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver store_AcceptInfo
//go:noescape
func AcceptInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver load_AcceptInfo
//go:noescape
func AcceptInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver store_CreateInfo
//go:noescape
func CreateInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver load_CreateInfo
//go:noescape
func CreateInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver store_SocketInfo
//go:noescape
func SocketInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver load_SocketInfo
//go:noescape
func SocketInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver store_SocketProperties
//go:noescape
func SocketPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver load_SocketProperties
//go:noescape
func SocketPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver has_Close
//go:noescape
func HasFuncClose() js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver func_Close
//go:noescape
func FuncClose(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcpserver call_Close
//go:noescape
func CallClose(
	retPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver try_Close
//go:noescape
func TryClose(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver has_Create
//go:noescape
func HasFuncCreate() js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver func_Create
//go:noescape
func FuncCreate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcpserver call_Create
//go:noescape
func CallCreate(
	retPtr unsafe.Pointer,
	properties unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver try_Create
//go:noescape
func TryCreate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	properties unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver has_Disconnect
//go:noescape
func HasFuncDisconnect() js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver func_Disconnect
//go:noescape
func FuncDisconnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcpserver call_Disconnect
//go:noescape
func CallDisconnect(
	retPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver try_Disconnect
//go:noescape
func TryDisconnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver has_GetInfo
//go:noescape
func HasFuncGetInfo() js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver func_GetInfo
//go:noescape
func FuncGetInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcpserver call_GetInfo
//go:noescape
func CallGetInfo(
	retPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver try_GetInfo
//go:noescape
func TryGetInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver has_GetSockets
//go:noescape
func HasFuncGetSockets() js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver func_GetSockets
//go:noescape
func FuncGetSockets(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcpserver call_GetSockets
//go:noescape
func CallGetSockets(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver try_GetSockets
//go:noescape
func TryGetSockets(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver has_Listen
//go:noescape
func HasFuncListen() js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver func_Listen
//go:noescape
func FuncListen(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcpserver call_Listen
//go:noescape
func CallListen(
	retPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	port int32,
	backlog int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver try_Listen
//go:noescape
func TryListen(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	port int32,
	backlog int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver has_OnAccept
//go:noescape
func HasFuncOnAccept() js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver func_OnAccept
//go:noescape
func FuncOnAccept(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcpserver call_OnAccept
//go:noescape
func CallOnAccept(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver try_OnAccept
//go:noescape
func TryOnAccept(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver has_OffAccept
//go:noescape
func HasFuncOffAccept() js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver func_OffAccept
//go:noescape
func FuncOffAccept(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcpserver call_OffAccept
//go:noescape
func CallOffAccept(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver try_OffAccept
//go:noescape
func TryOffAccept(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver has_HasOnAccept
//go:noescape
func HasFuncHasOnAccept() js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver func_HasOnAccept
//go:noescape
func FuncHasOnAccept(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcpserver call_HasOnAccept
//go:noescape
func CallHasOnAccept(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver try_HasOnAccept
//go:noescape
func TryHasOnAccept(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver has_OnAcceptError
//go:noescape
func HasFuncOnAcceptError() js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver func_OnAcceptError
//go:noescape
func FuncOnAcceptError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcpserver call_OnAcceptError
//go:noescape
func CallOnAcceptError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver try_OnAcceptError
//go:noescape
func TryOnAcceptError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver has_OffAcceptError
//go:noescape
func HasFuncOffAcceptError() js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver func_OffAcceptError
//go:noescape
func FuncOffAcceptError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcpserver call_OffAcceptError
//go:noescape
func CallOffAcceptError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver try_OffAcceptError
//go:noescape
func TryOffAcceptError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver has_HasOnAcceptError
//go:noescape
func HasFuncHasOnAcceptError() js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver func_HasOnAcceptError
//go:noescape
func FuncHasOnAcceptError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcpserver call_HasOnAcceptError
//go:noescape
func CallHasOnAcceptError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver try_HasOnAcceptError
//go:noescape
func TryHasOnAcceptError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver has_SetPaused
//go:noescape
func HasFuncSetPaused() js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver func_SetPaused
//go:noescape
func FuncSetPaused(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcpserver call_SetPaused
//go:noescape
func CallSetPaused(
	retPtr unsafe.Pointer,
	socketId int32,
	paused js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver try_SetPaused
//go:noescape
func TrySetPaused(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	paused js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver has_Update
//go:noescape
func HasFuncUpdate() js.Ref

//go:wasmimport plat/js/webext/sockets/tcpserver func_Update
//go:noescape
func FuncUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcpserver call_Update
//go:noescape
func CallUpdate(
	retPtr unsafe.Pointer,
	socketId int32,
	properties unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcpserver try_Update
//go:noescape
func TryUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	properties unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
