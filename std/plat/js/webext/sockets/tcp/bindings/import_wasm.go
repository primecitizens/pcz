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

//go:wasmimport plat/js/webext/sockets/tcp store_CreateInfo
//go:noescape
func CreateInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp load_CreateInfo
//go:noescape
func CreateInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/tcp constof_DnsQueryType
//go:noescape
func ConstOfDnsQueryType(str js.Ref) uint32

//go:wasmimport plat/js/webext/sockets/tcp store_SocketInfo
//go:noescape
func SocketInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp load_SocketInfo
//go:noescape
func SocketInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/tcp store_ReceiveErrorInfo
//go:noescape
func ReceiveErrorInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp load_ReceiveErrorInfo
//go:noescape
func ReceiveErrorInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/tcp store_ReceiveInfo
//go:noescape
func ReceiveInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp load_ReceiveInfo
//go:noescape
func ReceiveInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/tcp store_TLSVersionConstraints
//go:noescape
func TLSVersionConstraintsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp load_TLSVersionConstraints
//go:noescape
func TLSVersionConstraintsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/tcp store_SecureOptions
//go:noescape
func SecureOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp load_SecureOptions
//go:noescape
func SecureOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/tcp store_SendInfo
//go:noescape
func SendInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp load_SendInfo
//go:noescape
func SendInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/tcp store_SocketProperties
//go:noescape
func SocketPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp load_SocketProperties
//go:noescape
func SocketPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/tcp has_Close
//go:noescape
func HasFuncClose() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_Close
//go:noescape
func FuncClose(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_Close
//go:noescape
func CallClose(
	retPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_Close
//go:noescape
func TryClose(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_Connect
//go:noescape
func HasFuncConnect() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_Connect
//go:noescape
func FuncConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_Connect
//go:noescape
func CallConnect(
	retPtr unsafe.Pointer,
	socketId int32,
	peerAddress js.Ref,
	peerPort int32,
	dnsQueryType uint32,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_Connect
//go:noescape
func TryConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	peerAddress js.Ref,
	peerPort int32,
	dnsQueryType uint32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_Create
//go:noescape
func HasFuncCreate() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_Create
//go:noescape
func FuncCreate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_Create
//go:noescape
func CallCreate(
	retPtr unsafe.Pointer,
	properties unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_Create
//go:noescape
func TryCreate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	properties unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_Disconnect
//go:noescape
func HasFuncDisconnect() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_Disconnect
//go:noescape
func FuncDisconnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_Disconnect
//go:noescape
func CallDisconnect(
	retPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_Disconnect
//go:noescape
func TryDisconnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_GetInfo
//go:noescape
func HasFuncGetInfo() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_GetInfo
//go:noescape
func FuncGetInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_GetInfo
//go:noescape
func CallGetInfo(
	retPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_GetInfo
//go:noescape
func TryGetInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_GetSockets
//go:noescape
func HasFuncGetSockets() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_GetSockets
//go:noescape
func FuncGetSockets(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_GetSockets
//go:noescape
func CallGetSockets(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_GetSockets
//go:noescape
func TryGetSockets(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_OnReceive
//go:noescape
func HasFuncOnReceive() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_OnReceive
//go:noescape
func FuncOnReceive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_OnReceive
//go:noescape
func CallOnReceive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_OnReceive
//go:noescape
func TryOnReceive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_OffReceive
//go:noescape
func HasFuncOffReceive() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_OffReceive
//go:noescape
func FuncOffReceive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_OffReceive
//go:noescape
func CallOffReceive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_OffReceive
//go:noescape
func TryOffReceive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_HasOnReceive
//go:noescape
func HasFuncHasOnReceive() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_HasOnReceive
//go:noescape
func FuncHasOnReceive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_HasOnReceive
//go:noescape
func CallHasOnReceive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_HasOnReceive
//go:noescape
func TryHasOnReceive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_OnReceiveError
//go:noescape
func HasFuncOnReceiveError() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_OnReceiveError
//go:noescape
func FuncOnReceiveError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_OnReceiveError
//go:noescape
func CallOnReceiveError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_OnReceiveError
//go:noescape
func TryOnReceiveError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_OffReceiveError
//go:noescape
func HasFuncOffReceiveError() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_OffReceiveError
//go:noescape
func FuncOffReceiveError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_OffReceiveError
//go:noescape
func CallOffReceiveError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_OffReceiveError
//go:noescape
func TryOffReceiveError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_HasOnReceiveError
//go:noescape
func HasFuncHasOnReceiveError() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_HasOnReceiveError
//go:noescape
func FuncHasOnReceiveError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_HasOnReceiveError
//go:noescape
func CallHasOnReceiveError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_HasOnReceiveError
//go:noescape
func TryHasOnReceiveError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_Secure
//go:noescape
func HasFuncSecure() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_Secure
//go:noescape
func FuncSecure(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_Secure
//go:noescape
func CallSecure(
	retPtr unsafe.Pointer,
	socketId int32,
	options unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_Secure
//go:noescape
func TrySecure(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	options unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_Send
//go:noescape
func HasFuncSend() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_Send
//go:noescape
func FuncSend(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_Send
//go:noescape
func CallSend(
	retPtr unsafe.Pointer,
	socketId int32,
	data js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_Send
//go:noescape
func TrySend(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	data js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_SetKeepAlive
//go:noescape
func HasFuncSetKeepAlive() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_SetKeepAlive
//go:noescape
func FuncSetKeepAlive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_SetKeepAlive
//go:noescape
func CallSetKeepAlive(
	retPtr unsafe.Pointer,
	socketId int32,
	enable js.Ref,
	delay int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_SetKeepAlive
//go:noescape
func TrySetKeepAlive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	enable js.Ref,
	delay int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_SetNoDelay
//go:noescape
func HasFuncSetNoDelay() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_SetNoDelay
//go:noescape
func FuncSetNoDelay(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_SetNoDelay
//go:noescape
func CallSetNoDelay(
	retPtr unsafe.Pointer,
	socketId int32,
	noDelay js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_SetNoDelay
//go:noescape
func TrySetNoDelay(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	noDelay js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_SetPaused
//go:noescape
func HasFuncSetPaused() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_SetPaused
//go:noescape
func FuncSetPaused(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_SetPaused
//go:noescape
func CallSetPaused(
	retPtr unsafe.Pointer,
	socketId int32,
	paused js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_SetPaused
//go:noescape
func TrySetPaused(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	paused js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp has_Update
//go:noescape
func HasFuncUpdate() js.Ref

//go:wasmimport plat/js/webext/sockets/tcp func_Update
//go:noescape
func FuncUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/tcp call_Update
//go:noescape
func CallUpdate(
	retPtr unsafe.Pointer,
	socketId int32,
	properties unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/tcp try_Update
//go:noescape
func TryUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	properties unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
