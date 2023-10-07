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

//go:wasmimport plat/js/webext/socket store_AcceptInfo
//go:noescape
func AcceptInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/socket load_AcceptInfo
//go:noescape
func AcceptInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/socket store_CreateInfo
//go:noescape
func CreateInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/socket load_CreateInfo
//go:noescape
func CreateInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/socket store_CreateOptions
//go:noescape
func CreateOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/socket load_CreateOptions
//go:noescape
func CreateOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/socket constof_SocketType
//go:noescape
func ConstOfSocketType(str js.Ref) uint32

//go:wasmimport plat/js/webext/socket store_SocketInfo
//go:noescape
func SocketInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/socket load_SocketInfo
//go:noescape
func SocketInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/socket store_NetworkInterface
//go:noescape
func NetworkInterfaceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/socket load_NetworkInterface
//go:noescape
func NetworkInterfaceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/socket store_ReadInfo
//go:noescape
func ReadInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/socket load_ReadInfo
//go:noescape
func ReadInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/socket store_RecvFromInfo
//go:noescape
func RecvFromInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/socket load_RecvFromInfo
//go:noescape
func RecvFromInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/socket store_TLSVersionConstraints
//go:noescape
func TLSVersionConstraintsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/socket load_TLSVersionConstraints
//go:noescape
func TLSVersionConstraintsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/socket store_SecureOptions
//go:noescape
func SecureOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/socket load_SecureOptions
//go:noescape
func SecureOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/socket store_WriteInfo
//go:noescape
func WriteInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/socket load_WriteInfo
//go:noescape
func WriteInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/socket has_Accept
//go:noescape
func HasFuncAccept() js.Ref

//go:wasmimport plat/js/webext/socket func_Accept
//go:noescape
func FuncAccept(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_Accept
//go:noescape
func CallAccept(
	retPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_Accept
//go:noescape
func TryAccept(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_Bind
//go:noescape
func HasFuncBind() js.Ref

//go:wasmimport plat/js/webext/socket func_Bind
//go:noescape
func FuncBind(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_Bind
//go:noescape
func CallBind(
	retPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	port int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_Bind
//go:noescape
func TryBind(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	port int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_Connect
//go:noescape
func HasFuncConnect() js.Ref

//go:wasmimport plat/js/webext/socket func_Connect
//go:noescape
func FuncConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_Connect
//go:noescape
func CallConnect(
	retPtr unsafe.Pointer,
	socketId int32,
	hostname js.Ref,
	port int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_Connect
//go:noescape
func TryConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	hostname js.Ref,
	port int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_Create
//go:noescape
func HasFuncCreate() js.Ref

//go:wasmimport plat/js/webext/socket func_Create
//go:noescape
func FuncCreate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_Create
//go:noescape
func CallCreate(
	retPtr unsafe.Pointer,
	typ uint32,
	options unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_Create
//go:noescape
func TryCreate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ uint32,
	options unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_Destroy
//go:noescape
func HasFuncDestroy() js.Ref

//go:wasmimport plat/js/webext/socket func_Destroy
//go:noescape
func FuncDestroy(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_Destroy
//go:noescape
func CallDestroy(
	retPtr unsafe.Pointer,
	socketId int32)

//go:wasmimport plat/js/webext/socket try_Destroy
//go:noescape
func TryDestroy(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_Disconnect
//go:noescape
func HasFuncDisconnect() js.Ref

//go:wasmimport plat/js/webext/socket func_Disconnect
//go:noescape
func FuncDisconnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_Disconnect
//go:noescape
func CallDisconnect(
	retPtr unsafe.Pointer,
	socketId int32)

//go:wasmimport plat/js/webext/socket try_Disconnect
//go:noescape
func TryDisconnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_GetInfo
//go:noescape
func HasFuncGetInfo() js.Ref

//go:wasmimport plat/js/webext/socket func_GetInfo
//go:noescape
func FuncGetInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_GetInfo
//go:noescape
func CallGetInfo(
	retPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_GetInfo
//go:noescape
func TryGetInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_GetJoinedGroups
//go:noescape
func HasFuncGetJoinedGroups() js.Ref

//go:wasmimport plat/js/webext/socket func_GetJoinedGroups
//go:noescape
func FuncGetJoinedGroups(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_GetJoinedGroups
//go:noescape
func CallGetJoinedGroups(
	retPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_GetJoinedGroups
//go:noescape
func TryGetJoinedGroups(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_GetNetworkList
//go:noescape
func HasFuncGetNetworkList() js.Ref

//go:wasmimport plat/js/webext/socket func_GetNetworkList
//go:noescape
func FuncGetNetworkList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_GetNetworkList
//go:noescape
func CallGetNetworkList(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_GetNetworkList
//go:noescape
func TryGetNetworkList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_JoinGroup
//go:noescape
func HasFuncJoinGroup() js.Ref

//go:wasmimport plat/js/webext/socket func_JoinGroup
//go:noescape
func FuncJoinGroup(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_JoinGroup
//go:noescape
func CallJoinGroup(
	retPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_JoinGroup
//go:noescape
func TryJoinGroup(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_LeaveGroup
//go:noescape
func HasFuncLeaveGroup() js.Ref

//go:wasmimport plat/js/webext/socket func_LeaveGroup
//go:noescape
func FuncLeaveGroup(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_LeaveGroup
//go:noescape
func CallLeaveGroup(
	retPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_LeaveGroup
//go:noescape
func TryLeaveGroup(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_Listen
//go:noescape
func HasFuncListen() js.Ref

//go:wasmimport plat/js/webext/socket func_Listen
//go:noescape
func FuncListen(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_Listen
//go:noescape
func CallListen(
	retPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	port int32,
	backlog int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_Listen
//go:noescape
func TryListen(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	address js.Ref,
	port int32,
	backlog int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_Read
//go:noescape
func HasFuncRead() js.Ref

//go:wasmimport plat/js/webext/socket func_Read
//go:noescape
func FuncRead(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_Read
//go:noescape
func CallRead(
	retPtr unsafe.Pointer,
	socketId int32,
	bufferSize int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_Read
//go:noescape
func TryRead(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	bufferSize int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_RecvFrom
//go:noescape
func HasFuncRecvFrom() js.Ref

//go:wasmimport plat/js/webext/socket func_RecvFrom
//go:noescape
func FuncRecvFrom(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_RecvFrom
//go:noescape
func CallRecvFrom(
	retPtr unsafe.Pointer,
	socketId int32,
	bufferSize int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_RecvFrom
//go:noescape
func TryRecvFrom(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	bufferSize int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_Secure
//go:noescape
func HasFuncSecure() js.Ref

//go:wasmimport plat/js/webext/socket func_Secure
//go:noescape
func FuncSecure(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_Secure
//go:noescape
func CallSecure(
	retPtr unsafe.Pointer,
	socketId int32,
	options unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_Secure
//go:noescape
func TrySecure(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	options unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_SendTo
//go:noescape
func HasFuncSendTo() js.Ref

//go:wasmimport plat/js/webext/socket func_SendTo
//go:noescape
func FuncSendTo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_SendTo
//go:noescape
func CallSendTo(
	retPtr unsafe.Pointer,
	socketId int32,
	data js.Ref,
	address js.Ref,
	port int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_SendTo
//go:noescape
func TrySendTo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	data js.Ref,
	address js.Ref,
	port int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_SetKeepAlive
//go:noescape
func HasFuncSetKeepAlive() js.Ref

//go:wasmimport plat/js/webext/socket func_SetKeepAlive
//go:noescape
func FuncSetKeepAlive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_SetKeepAlive
//go:noescape
func CallSetKeepAlive(
	retPtr unsafe.Pointer,
	socketId int32,
	enable js.Ref,
	delay int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_SetKeepAlive
//go:noescape
func TrySetKeepAlive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	enable js.Ref,
	delay int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_SetMulticastLoopbackMode
//go:noescape
func HasFuncSetMulticastLoopbackMode() js.Ref

//go:wasmimport plat/js/webext/socket func_SetMulticastLoopbackMode
//go:noescape
func FuncSetMulticastLoopbackMode(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_SetMulticastLoopbackMode
//go:noescape
func CallSetMulticastLoopbackMode(
	retPtr unsafe.Pointer,
	socketId int32,
	enabled js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_SetMulticastLoopbackMode
//go:noescape
func TrySetMulticastLoopbackMode(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	enabled js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_SetMulticastTimeToLive
//go:noescape
func HasFuncSetMulticastTimeToLive() js.Ref

//go:wasmimport plat/js/webext/socket func_SetMulticastTimeToLive
//go:noescape
func FuncSetMulticastTimeToLive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_SetMulticastTimeToLive
//go:noescape
func CallSetMulticastTimeToLive(
	retPtr unsafe.Pointer,
	socketId int32,
	ttl int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_SetMulticastTimeToLive
//go:noescape
func TrySetMulticastTimeToLive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	ttl int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_SetNoDelay
//go:noescape
func HasFuncSetNoDelay() js.Ref

//go:wasmimport plat/js/webext/socket func_SetNoDelay
//go:noescape
func FuncSetNoDelay(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_SetNoDelay
//go:noescape
func CallSetNoDelay(
	retPtr unsafe.Pointer,
	socketId int32,
	noDelay js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_SetNoDelay
//go:noescape
func TrySetNoDelay(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	noDelay js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/socket has_Write
//go:noescape
func HasFuncWrite() js.Ref

//go:wasmimport plat/js/webext/socket func_Write
//go:noescape
func FuncWrite(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/socket call_Write
//go:noescape
func CallWrite(
	retPtr unsafe.Pointer,
	socketId int32,
	data js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/socket try_Write
//go:noescape
func TryWrite(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	socketId int32,
	data js.Ref,
	callback js.Ref) (ok js.Ref)
