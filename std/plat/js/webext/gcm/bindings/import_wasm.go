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

//go:wasmimport plat/js/webext/gcm get_MAX_MESSAGE_SIZE
//go:noescape
func GetMAX_MESSAGE_SIZE(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/gcm set_MAX_MESSAGE_SIZE
//go:noescape
func SetMAX_MESSAGE_SIZE(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/gcm store_OnMessageArgMessage
//go:noescape
func OnMessageArgMessageJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/gcm load_OnMessageArgMessage
//go:noescape
func OnMessageArgMessageJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/gcm store_OnSendErrorArgError
//go:noescape
func OnSendErrorArgErrorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/gcm load_OnSendErrorArgError
//go:noescape
func OnSendErrorArgErrorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/gcm store_SendArgMessage
//go:noescape
func SendArgMessageJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/gcm load_SendArgMessage
//go:noescape
func SendArgMessageJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/gcm has_OnMessage
//go:noescape
func HasFuncOnMessage() js.Ref

//go:wasmimport plat/js/webext/gcm func_OnMessage
//go:noescape
func FuncOnMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/gcm call_OnMessage
//go:noescape
func CallOnMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/gcm try_OnMessage
//go:noescape
func TryOnMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/gcm has_OffMessage
//go:noescape
func HasFuncOffMessage() js.Ref

//go:wasmimport plat/js/webext/gcm func_OffMessage
//go:noescape
func FuncOffMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/gcm call_OffMessage
//go:noescape
func CallOffMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/gcm try_OffMessage
//go:noescape
func TryOffMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/gcm has_HasOnMessage
//go:noescape
func HasFuncHasOnMessage() js.Ref

//go:wasmimport plat/js/webext/gcm func_HasOnMessage
//go:noescape
func FuncHasOnMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/gcm call_HasOnMessage
//go:noescape
func CallHasOnMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/gcm try_HasOnMessage
//go:noescape
func TryHasOnMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/gcm has_OnMessagesDeleted
//go:noescape
func HasFuncOnMessagesDeleted() js.Ref

//go:wasmimport plat/js/webext/gcm func_OnMessagesDeleted
//go:noescape
func FuncOnMessagesDeleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/gcm call_OnMessagesDeleted
//go:noescape
func CallOnMessagesDeleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/gcm try_OnMessagesDeleted
//go:noescape
func TryOnMessagesDeleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/gcm has_OffMessagesDeleted
//go:noescape
func HasFuncOffMessagesDeleted() js.Ref

//go:wasmimport plat/js/webext/gcm func_OffMessagesDeleted
//go:noescape
func FuncOffMessagesDeleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/gcm call_OffMessagesDeleted
//go:noescape
func CallOffMessagesDeleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/gcm try_OffMessagesDeleted
//go:noescape
func TryOffMessagesDeleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/gcm has_HasOnMessagesDeleted
//go:noescape
func HasFuncHasOnMessagesDeleted() js.Ref

//go:wasmimport plat/js/webext/gcm func_HasOnMessagesDeleted
//go:noescape
func FuncHasOnMessagesDeleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/gcm call_HasOnMessagesDeleted
//go:noescape
func CallHasOnMessagesDeleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/gcm try_HasOnMessagesDeleted
//go:noescape
func TryHasOnMessagesDeleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/gcm has_OnSendError
//go:noescape
func HasFuncOnSendError() js.Ref

//go:wasmimport plat/js/webext/gcm func_OnSendError
//go:noescape
func FuncOnSendError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/gcm call_OnSendError
//go:noescape
func CallOnSendError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/gcm try_OnSendError
//go:noescape
func TryOnSendError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/gcm has_OffSendError
//go:noescape
func HasFuncOffSendError() js.Ref

//go:wasmimport plat/js/webext/gcm func_OffSendError
//go:noescape
func FuncOffSendError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/gcm call_OffSendError
//go:noescape
func CallOffSendError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/gcm try_OffSendError
//go:noescape
func TryOffSendError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/gcm has_HasOnSendError
//go:noescape
func HasFuncHasOnSendError() js.Ref

//go:wasmimport plat/js/webext/gcm func_HasOnSendError
//go:noescape
func FuncHasOnSendError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/gcm call_HasOnSendError
//go:noescape
func CallHasOnSendError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/gcm try_HasOnSendError
//go:noescape
func TryHasOnSendError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/gcm has_Register
//go:noescape
func HasFuncRegister() js.Ref

//go:wasmimport plat/js/webext/gcm func_Register
//go:noescape
func FuncRegister(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/gcm call_Register
//go:noescape
func CallRegister(
	retPtr unsafe.Pointer,
	senderIds js.Ref)

//go:wasmimport plat/js/webext/gcm try_Register
//go:noescape
func TryRegister(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	senderIds js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/gcm has_Send
//go:noescape
func HasFuncSend() js.Ref

//go:wasmimport plat/js/webext/gcm func_Send
//go:noescape
func FuncSend(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/gcm call_Send
//go:noescape
func CallSend(
	retPtr unsafe.Pointer,
	message unsafe.Pointer)

//go:wasmimport plat/js/webext/gcm try_Send
//go:noescape
func TrySend(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/gcm has_Unregister
//go:noescape
func HasFuncUnregister() js.Ref

//go:wasmimport plat/js/webext/gcm func_Unregister
//go:noescape
func FuncUnregister(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/gcm call_Unregister
//go:noescape
func CallUnregister(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/gcm try_Unregister
//go:noescape
func TryUnregister(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
