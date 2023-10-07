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

//go:wasmimport plat/js/webext/hid store_HidConnectInfo
//go:noescape
func HidConnectInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/hid load_HidConnectInfo
//go:noescape
func HidConnectInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/hid store_DeviceFilter
//go:noescape
func DeviceFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/hid load_DeviceFilter
//go:noescape
func DeviceFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/hid store_HidCollectionInfo
//go:noescape
func HidCollectionInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/hid load_HidCollectionInfo
//go:noescape
func HidCollectionInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/hid store_HidDeviceInfo
//go:noescape
func HidDeviceInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/hid load_HidDeviceInfo
//go:noescape
func HidDeviceInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/hid store_GetDevicesOptions
//go:noescape
func GetDevicesOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/hid load_GetDevicesOptions
//go:noescape
func GetDevicesOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/hid has_Connect
//go:noescape
func HasFuncConnect() js.Ref

//go:wasmimport plat/js/webext/hid func_Connect
//go:noescape
func FuncConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/hid call_Connect
//go:noescape
func CallConnect(
	retPtr unsafe.Pointer,
	deviceId int32)

//go:wasmimport plat/js/webext/hid try_Connect
//go:noescape
func TryConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deviceId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/hid has_Disconnect
//go:noescape
func HasFuncDisconnect() js.Ref

//go:wasmimport plat/js/webext/hid func_Disconnect
//go:noescape
func FuncDisconnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/hid call_Disconnect
//go:noescape
func CallDisconnect(
	retPtr unsafe.Pointer,
	connectionId int32)

//go:wasmimport plat/js/webext/hid try_Disconnect
//go:noescape
func TryDisconnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	connectionId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/hid has_GetDevices
//go:noescape
func HasFuncGetDevices() js.Ref

//go:wasmimport plat/js/webext/hid func_GetDevices
//go:noescape
func FuncGetDevices(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/hid call_GetDevices
//go:noescape
func CallGetDevices(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/hid try_GetDevices
//go:noescape
func TryGetDevices(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/hid has_OnDeviceAdded
//go:noescape
func HasFuncOnDeviceAdded() js.Ref

//go:wasmimport plat/js/webext/hid func_OnDeviceAdded
//go:noescape
func FuncOnDeviceAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/hid call_OnDeviceAdded
//go:noescape
func CallOnDeviceAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/hid try_OnDeviceAdded
//go:noescape
func TryOnDeviceAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/hid has_OffDeviceAdded
//go:noescape
func HasFuncOffDeviceAdded() js.Ref

//go:wasmimport plat/js/webext/hid func_OffDeviceAdded
//go:noescape
func FuncOffDeviceAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/hid call_OffDeviceAdded
//go:noescape
func CallOffDeviceAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/hid try_OffDeviceAdded
//go:noescape
func TryOffDeviceAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/hid has_HasOnDeviceAdded
//go:noescape
func HasFuncHasOnDeviceAdded() js.Ref

//go:wasmimport plat/js/webext/hid func_HasOnDeviceAdded
//go:noescape
func FuncHasOnDeviceAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/hid call_HasOnDeviceAdded
//go:noescape
func CallHasOnDeviceAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/hid try_HasOnDeviceAdded
//go:noescape
func TryHasOnDeviceAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/hid has_OnDeviceRemoved
//go:noescape
func HasFuncOnDeviceRemoved() js.Ref

//go:wasmimport plat/js/webext/hid func_OnDeviceRemoved
//go:noescape
func FuncOnDeviceRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/hid call_OnDeviceRemoved
//go:noescape
func CallOnDeviceRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/hid try_OnDeviceRemoved
//go:noescape
func TryOnDeviceRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/hid has_OffDeviceRemoved
//go:noescape
func HasFuncOffDeviceRemoved() js.Ref

//go:wasmimport plat/js/webext/hid func_OffDeviceRemoved
//go:noescape
func FuncOffDeviceRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/hid call_OffDeviceRemoved
//go:noescape
func CallOffDeviceRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/hid try_OffDeviceRemoved
//go:noescape
func TryOffDeviceRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/hid has_HasOnDeviceRemoved
//go:noescape
func HasFuncHasOnDeviceRemoved() js.Ref

//go:wasmimport plat/js/webext/hid func_HasOnDeviceRemoved
//go:noescape
func FuncHasOnDeviceRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/hid call_HasOnDeviceRemoved
//go:noescape
func CallHasOnDeviceRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/hid try_HasOnDeviceRemoved
//go:noescape
func TryHasOnDeviceRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/hid has_Receive
//go:noescape
func HasFuncReceive() js.Ref

//go:wasmimport plat/js/webext/hid func_Receive
//go:noescape
func FuncReceive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/hid call_Receive
//go:noescape
func CallReceive(
	retPtr unsafe.Pointer,
	connectionId int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/hid try_Receive
//go:noescape
func TryReceive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	connectionId int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/hid has_ReceiveFeatureReport
//go:noescape
func HasFuncReceiveFeatureReport() js.Ref

//go:wasmimport plat/js/webext/hid func_ReceiveFeatureReport
//go:noescape
func FuncReceiveFeatureReport(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/hid call_ReceiveFeatureReport
//go:noescape
func CallReceiveFeatureReport(
	retPtr unsafe.Pointer,
	connectionId int32,
	reportId int32)

//go:wasmimport plat/js/webext/hid try_ReceiveFeatureReport
//go:noescape
func TryReceiveFeatureReport(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	connectionId int32,
	reportId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/hid has_Send
//go:noescape
func HasFuncSend() js.Ref

//go:wasmimport plat/js/webext/hid func_Send
//go:noescape
func FuncSend(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/hid call_Send
//go:noescape
func CallSend(
	retPtr unsafe.Pointer,
	connectionId int32,
	reportId int32,
	data js.Ref)

//go:wasmimport plat/js/webext/hid try_Send
//go:noescape
func TrySend(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	connectionId int32,
	reportId int32,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/hid has_SendFeatureReport
//go:noescape
func HasFuncSendFeatureReport() js.Ref

//go:wasmimport plat/js/webext/hid func_SendFeatureReport
//go:noescape
func FuncSendFeatureReport(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/hid call_SendFeatureReport
//go:noescape
func CallSendFeatureReport(
	retPtr unsafe.Pointer,
	connectionId int32,
	reportId int32,
	data js.Ref)

//go:wasmimport plat/js/webext/hid try_SendFeatureReport
//go:noescape
func TrySendFeatureReport(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	connectionId int32,
	reportId int32,
	data js.Ref) (ok js.Ref)
