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

//go:wasmimport plat/js/webext/sockets/udp store_CreateInfo
//go:noescape
func CreateInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/udp load_CreateInfo
//go:noescape
func CreateInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/udp constof_DnsQueryType
//go:noescape
func ConstOfDnsQueryType(str js.Ref) uint32

//go:wasmimport plat/js/webext/sockets/udp store_SocketInfo
//go:noescape
func SocketInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/udp load_SocketInfo
//go:noescape
func SocketInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/udp store_ReceiveErrorInfo
//go:noescape
func ReceiveErrorInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/udp load_ReceiveErrorInfo
//go:noescape
func ReceiveErrorInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/udp store_ReceiveInfo
//go:noescape
func ReceiveInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/udp load_ReceiveInfo
//go:noescape
func ReceiveInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/udp store_SendInfo
//go:noescape
func SendInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/udp load_SendInfo
//go:noescape
func SendInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/udp store_SocketProperties
//go:noescape
func SocketPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sockets/udp load_SocketProperties
//go:noescape
func SocketPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sockets/udp has_Bind
//go:noescape
func HasFuncBind() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_Bind
//go:noescape
func FuncBind(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_Bind
//go:noescape
func CallBind(
	retPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	port int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_Bind
//go:noescape
func TryBind(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	port int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_Close
//go:noescape
func HasFuncClose() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_Close
//go:noescape
func FuncClose(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_Close
//go:noescape
func CallClose(
	retPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_Close
//go:noescape
func TryClose(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_Create
//go:noescape
func HasFuncCreate() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_Create
//go:noescape
func FuncCreate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_Create
//go:noescape
func CallCreate(
	retPtr unsafe.Pointer,
	properties unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_Create
//go:noescape
func TryCreate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	properties unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_GetInfo
//go:noescape
func HasFuncGetInfo() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_GetInfo
//go:noescape
func FuncGetInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_GetInfo
//go:noescape
func CallGetInfo(
	retPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_GetInfo
//go:noescape
func TryGetInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_GetJoinedGroups
//go:noescape
func HasFuncGetJoinedGroups() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_GetJoinedGroups
//go:noescape
func FuncGetJoinedGroups(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_GetJoinedGroups
//go:noescape
func CallGetJoinedGroups(
	retPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_GetJoinedGroups
//go:noescape
func TryGetJoinedGroups(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_GetSockets
//go:noescape
func HasFuncGetSockets() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_GetSockets
//go:noescape
func FuncGetSockets(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_GetSockets
//go:noescape
func CallGetSockets(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_GetSockets
//go:noescape
func TryGetSockets(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_JoinGroup
//go:noescape
func HasFuncJoinGroup() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_JoinGroup
//go:noescape
func FuncJoinGroup(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_JoinGroup
//go:noescape
func CallJoinGroup(
	retPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_JoinGroup
//go:noescape
func TryJoinGroup(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_LeaveGroup
//go:noescape
func HasFuncLeaveGroup() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_LeaveGroup
//go:noescape
func FuncLeaveGroup(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_LeaveGroup
//go:noescape
func CallLeaveGroup(
	retPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_LeaveGroup
//go:noescape
func TryLeaveGroup(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_OnReceive
//go:noescape
func HasFuncOnReceive() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_OnReceive
//go:noescape
func FuncOnReceive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_OnReceive
//go:noescape
func CallOnReceive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_OnReceive
//go:noescape
func TryOnReceive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_OffReceive
//go:noescape
func HasFuncOffReceive() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_OffReceive
//go:noescape
func FuncOffReceive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_OffReceive
//go:noescape
func CallOffReceive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_OffReceive
//go:noescape
func TryOffReceive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_HasOnReceive
//go:noescape
func HasFuncHasOnReceive() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_HasOnReceive
//go:noescape
func FuncHasOnReceive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_HasOnReceive
//go:noescape
func CallHasOnReceive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_HasOnReceive
//go:noescape
func TryHasOnReceive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_OnReceiveError
//go:noescape
func HasFuncOnReceiveError() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_OnReceiveError
//go:noescape
func FuncOnReceiveError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_OnReceiveError
//go:noescape
func CallOnReceiveError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_OnReceiveError
//go:noescape
func TryOnReceiveError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_OffReceiveError
//go:noescape
func HasFuncOffReceiveError() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_OffReceiveError
//go:noescape
func FuncOffReceiveError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_OffReceiveError
//go:noescape
func CallOffReceiveError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_OffReceiveError
//go:noescape
func TryOffReceiveError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_HasOnReceiveError
//go:noescape
func HasFuncHasOnReceiveError() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_HasOnReceiveError
//go:noescape
func FuncHasOnReceiveError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_HasOnReceiveError
//go:noescape
func CallHasOnReceiveError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_HasOnReceiveError
//go:noescape
func TryHasOnReceiveError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_Send
//go:noescape
func HasFuncSend() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_Send
//go:noescape
func FuncSend(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_Send
//go:noescape
func CallSend(
	retPtr unsafe.Pointer,
	socketId int32,
	data js.Ref,
	address js.Ref,
	port int32,
	dnsQueryType uint32,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_Send
//go:noescape
func TrySend(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	data js.Ref,
	address js.Ref,
	port int32,
	dnsQueryType uint32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_SetBroadcast
//go:noescape
func HasFuncSetBroadcast() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_SetBroadcast
//go:noescape
func FuncSetBroadcast(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_SetBroadcast
//go:noescape
func CallSetBroadcast(
	retPtr unsafe.Pointer,
	socketId int32,
	enabled js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_SetBroadcast
//go:noescape
func TrySetBroadcast(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	enabled js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_SetMulticastLoopbackMode
//go:noescape
func HasFuncSetMulticastLoopbackMode() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_SetMulticastLoopbackMode
//go:noescape
func FuncSetMulticastLoopbackMode(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_SetMulticastLoopbackMode
//go:noescape
func CallSetMulticastLoopbackMode(
	retPtr unsafe.Pointer,
	socketId int32,
	enabled js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_SetMulticastLoopbackMode
//go:noescape
func TrySetMulticastLoopbackMode(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	enabled js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_SetMulticastTimeToLive
//go:noescape
func HasFuncSetMulticastTimeToLive() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_SetMulticastTimeToLive
//go:noescape
func FuncSetMulticastTimeToLive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_SetMulticastTimeToLive
//go:noescape
func CallSetMulticastTimeToLive(
	retPtr unsafe.Pointer,
	socketId int32,
	ttl int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_SetMulticastTimeToLive
//go:noescape
func TrySetMulticastTimeToLive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	ttl int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_SetPaused
//go:noescape
func HasFuncSetPaused() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_SetPaused
//go:noescape
func FuncSetPaused(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_SetPaused
//go:noescape
func CallSetPaused(
	retPtr unsafe.Pointer,
	socketId int32,
	paused js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_SetPaused
//go:noescape
func TrySetPaused(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	paused js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sockets/udp has_Update
//go:noescape
func HasFuncUpdate() js.Ref

//go:wasmimport plat/js/webext/sockets/udp func_Update
//go:noescape
func FuncUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sockets/udp call_Update
//go:noescape
func CallUpdate(
	retPtr unsafe.Pointer,
	socketId int32,
	properties unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sockets/udp try_Update
//go:noescape
func TryUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	properties unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
