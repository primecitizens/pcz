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

//go:wasmimport plat/js/webext/vpnprovider store_Parameters
//go:noescape
func ParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/vpnprovider load_Parameters
//go:noescape
func ParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/vpnprovider constof_PlatformMessage
//go:noescape
func ConstOfPlatformMessage(str js.Ref) uint32

//go:wasmimport plat/js/webext/vpnprovider constof_UIEvent
//go:noescape
func ConstOfUIEvent(str js.Ref) uint32

//go:wasmimport plat/js/webext/vpnprovider constof_VpnConnectionState
//go:noescape
func ConstOfVpnConnectionState(str js.Ref) uint32

//go:wasmimport plat/js/webext/vpnprovider has_CreateConfig
//go:noescape
func HasFuncCreateConfig() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_CreateConfig
//go:noescape
func FuncCreateConfig(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_CreateConfig
//go:noescape
func CallCreateConfig(
	retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_CreateConfig
//go:noescape
func TryCreateConfig(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_DestroyConfig
//go:noescape
func HasFuncDestroyConfig() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_DestroyConfig
//go:noescape
func FuncDestroyConfig(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_DestroyConfig
//go:noescape
func CallDestroyConfig(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_DestroyConfig
//go:noescape
func TryDestroyConfig(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_NotifyConnectionStateChanged
//go:noescape
func HasFuncNotifyConnectionStateChanged() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_NotifyConnectionStateChanged
//go:noescape
func FuncNotifyConnectionStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_NotifyConnectionStateChanged
//go:noescape
func CallNotifyConnectionStateChanged(
	retPtr unsafe.Pointer,
	state uint32)

//go:wasmimport plat/js/webext/vpnprovider try_NotifyConnectionStateChanged
//go:noescape
func TryNotifyConnectionStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	state uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_OnConfigCreated
//go:noescape
func HasFuncOnConfigCreated() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_OnConfigCreated
//go:noescape
func FuncOnConfigCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_OnConfigCreated
//go:noescape
func CallOnConfigCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_OnConfigCreated
//go:noescape
func TryOnConfigCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_OffConfigCreated
//go:noescape
func HasFuncOffConfigCreated() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_OffConfigCreated
//go:noescape
func FuncOffConfigCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_OffConfigCreated
//go:noescape
func CallOffConfigCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_OffConfigCreated
//go:noescape
func TryOffConfigCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_HasOnConfigCreated
//go:noescape
func HasFuncHasOnConfigCreated() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_HasOnConfigCreated
//go:noescape
func FuncHasOnConfigCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_HasOnConfigCreated
//go:noescape
func CallHasOnConfigCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_HasOnConfigCreated
//go:noescape
func TryHasOnConfigCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_OnConfigRemoved
//go:noescape
func HasFuncOnConfigRemoved() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_OnConfigRemoved
//go:noescape
func FuncOnConfigRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_OnConfigRemoved
//go:noescape
func CallOnConfigRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_OnConfigRemoved
//go:noescape
func TryOnConfigRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_OffConfigRemoved
//go:noescape
func HasFuncOffConfigRemoved() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_OffConfigRemoved
//go:noescape
func FuncOffConfigRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_OffConfigRemoved
//go:noescape
func CallOffConfigRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_OffConfigRemoved
//go:noescape
func TryOffConfigRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_HasOnConfigRemoved
//go:noescape
func HasFuncHasOnConfigRemoved() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_HasOnConfigRemoved
//go:noescape
func FuncHasOnConfigRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_HasOnConfigRemoved
//go:noescape
func CallHasOnConfigRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_HasOnConfigRemoved
//go:noescape
func TryHasOnConfigRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_OnPacketReceived
//go:noescape
func HasFuncOnPacketReceived() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_OnPacketReceived
//go:noescape
func FuncOnPacketReceived(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_OnPacketReceived
//go:noescape
func CallOnPacketReceived(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_OnPacketReceived
//go:noescape
func TryOnPacketReceived(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_OffPacketReceived
//go:noescape
func HasFuncOffPacketReceived() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_OffPacketReceived
//go:noescape
func FuncOffPacketReceived(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_OffPacketReceived
//go:noescape
func CallOffPacketReceived(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_OffPacketReceived
//go:noescape
func TryOffPacketReceived(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_HasOnPacketReceived
//go:noescape
func HasFuncHasOnPacketReceived() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_HasOnPacketReceived
//go:noescape
func FuncHasOnPacketReceived(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_HasOnPacketReceived
//go:noescape
func CallHasOnPacketReceived(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_HasOnPacketReceived
//go:noescape
func TryHasOnPacketReceived(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_OnPlatformMessage
//go:noescape
func HasFuncOnPlatformMessage() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_OnPlatformMessage
//go:noescape
func FuncOnPlatformMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_OnPlatformMessage
//go:noescape
func CallOnPlatformMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_OnPlatformMessage
//go:noescape
func TryOnPlatformMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_OffPlatformMessage
//go:noescape
func HasFuncOffPlatformMessage() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_OffPlatformMessage
//go:noescape
func FuncOffPlatformMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_OffPlatformMessage
//go:noescape
func CallOffPlatformMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_OffPlatformMessage
//go:noescape
func TryOffPlatformMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_HasOnPlatformMessage
//go:noescape
func HasFuncHasOnPlatformMessage() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_HasOnPlatformMessage
//go:noescape
func FuncHasOnPlatformMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_HasOnPlatformMessage
//go:noescape
func CallHasOnPlatformMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_HasOnPlatformMessage
//go:noescape
func TryHasOnPlatformMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_OnUIEvent
//go:noescape
func HasFuncOnUIEvent() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_OnUIEvent
//go:noescape
func FuncOnUIEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_OnUIEvent
//go:noescape
func CallOnUIEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_OnUIEvent
//go:noescape
func TryOnUIEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_OffUIEvent
//go:noescape
func HasFuncOffUIEvent() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_OffUIEvent
//go:noescape
func FuncOffUIEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_OffUIEvent
//go:noescape
func CallOffUIEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_OffUIEvent
//go:noescape
func TryOffUIEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_HasOnUIEvent
//go:noescape
func HasFuncHasOnUIEvent() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_HasOnUIEvent
//go:noescape
func FuncHasOnUIEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_HasOnUIEvent
//go:noescape
func CallHasOnUIEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_HasOnUIEvent
//go:noescape
func TryHasOnUIEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_SendPacket
//go:noescape
func HasFuncSendPacket() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_SendPacket
//go:noescape
func FuncSendPacket(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_SendPacket
//go:noescape
func CallSendPacket(
	retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/webext/vpnprovider try_SendPacket
//go:noescape
func TrySendPacket(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/vpnprovider has_SetParameters
//go:noescape
func HasFuncSetParameters() js.Ref

//go:wasmimport plat/js/webext/vpnprovider func_SetParameters
//go:noescape
func FuncSetParameters(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider call_SetParameters
//go:noescape
func CallSetParameters(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/vpnprovider try_SetParameters
//go:noescape
func TrySetParameters(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)
