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

//go:wasmimport plat/js/webext/bluetooth store_AdapterState
//go:noescape
func AdapterStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetooth load_AdapterState
//go:noescape
func AdapterStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetooth constof_FilterType
//go:noescape
func ConstOfFilterType(str js.Ref) uint32

//go:wasmimport plat/js/webext/bluetooth store_BluetoothFilter
//go:noescape
func BluetoothFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetooth load_BluetoothFilter
//go:noescape
func BluetoothFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetooth constof_VendorIdSource
//go:noescape
func ConstOfVendorIdSource(str js.Ref) uint32

//go:wasmimport plat/js/webext/bluetooth constof_DeviceType
//go:noescape
func ConstOfDeviceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/bluetooth constof_Transport
//go:noescape
func ConstOfTransport(str js.Ref) uint32

//go:wasmimport plat/js/webext/bluetooth store_Device
//go:noescape
func DeviceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetooth load_Device
//go:noescape
func DeviceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetooth has_GetAdapterState
//go:noescape
func HasFuncGetAdapterState() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_GetAdapterState
//go:noescape
func FuncGetAdapterState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_GetAdapterState
//go:noescape
func CallGetAdapterState(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth try_GetAdapterState
//go:noescape
func TryGetAdapterState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_GetDevice
//go:noescape
func HasFuncGetDevice() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_GetDevice
//go:noescape
func FuncGetDevice(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_GetDevice
//go:noescape
func CallGetDevice(
	retPtr unsafe.Pointer,
	deviceAddress js.Ref)

//go:wasmimport plat/js/webext/bluetooth try_GetDevice
//go:noescape
func TryGetDevice(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deviceAddress js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_GetDevices
//go:noescape
func HasFuncGetDevices() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_GetDevices
//go:noescape
func FuncGetDevices(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_GetDevices
//go:noescape
func CallGetDevices(
	retPtr unsafe.Pointer,
	filter unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth try_GetDevices
//go:noescape
func TryGetDevices(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_OnAdapterStateChanged
//go:noescape
func HasFuncOnAdapterStateChanged() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_OnAdapterStateChanged
//go:noescape
func FuncOnAdapterStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_OnAdapterStateChanged
//go:noescape
func CallOnAdapterStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetooth try_OnAdapterStateChanged
//go:noescape
func TryOnAdapterStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_OffAdapterStateChanged
//go:noescape
func HasFuncOffAdapterStateChanged() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_OffAdapterStateChanged
//go:noescape
func FuncOffAdapterStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_OffAdapterStateChanged
//go:noescape
func CallOffAdapterStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetooth try_OffAdapterStateChanged
//go:noescape
func TryOffAdapterStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_HasOnAdapterStateChanged
//go:noescape
func HasFuncHasOnAdapterStateChanged() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_HasOnAdapterStateChanged
//go:noescape
func FuncHasOnAdapterStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_HasOnAdapterStateChanged
//go:noescape
func CallHasOnAdapterStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetooth try_HasOnAdapterStateChanged
//go:noescape
func TryHasOnAdapterStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_OnDeviceAdded
//go:noescape
func HasFuncOnDeviceAdded() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_OnDeviceAdded
//go:noescape
func FuncOnDeviceAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_OnDeviceAdded
//go:noescape
func CallOnDeviceAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetooth try_OnDeviceAdded
//go:noescape
func TryOnDeviceAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_OffDeviceAdded
//go:noescape
func HasFuncOffDeviceAdded() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_OffDeviceAdded
//go:noescape
func FuncOffDeviceAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_OffDeviceAdded
//go:noescape
func CallOffDeviceAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetooth try_OffDeviceAdded
//go:noescape
func TryOffDeviceAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_HasOnDeviceAdded
//go:noescape
func HasFuncHasOnDeviceAdded() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_HasOnDeviceAdded
//go:noescape
func FuncHasOnDeviceAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_HasOnDeviceAdded
//go:noescape
func CallHasOnDeviceAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetooth try_HasOnDeviceAdded
//go:noescape
func TryHasOnDeviceAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_OnDeviceChanged
//go:noescape
func HasFuncOnDeviceChanged() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_OnDeviceChanged
//go:noescape
func FuncOnDeviceChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_OnDeviceChanged
//go:noescape
func CallOnDeviceChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetooth try_OnDeviceChanged
//go:noescape
func TryOnDeviceChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_OffDeviceChanged
//go:noescape
func HasFuncOffDeviceChanged() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_OffDeviceChanged
//go:noescape
func FuncOffDeviceChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_OffDeviceChanged
//go:noescape
func CallOffDeviceChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetooth try_OffDeviceChanged
//go:noescape
func TryOffDeviceChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_HasOnDeviceChanged
//go:noescape
func HasFuncHasOnDeviceChanged() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_HasOnDeviceChanged
//go:noescape
func FuncHasOnDeviceChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_HasOnDeviceChanged
//go:noescape
func CallHasOnDeviceChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetooth try_HasOnDeviceChanged
//go:noescape
func TryHasOnDeviceChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_OnDeviceRemoved
//go:noescape
func HasFuncOnDeviceRemoved() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_OnDeviceRemoved
//go:noescape
func FuncOnDeviceRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_OnDeviceRemoved
//go:noescape
func CallOnDeviceRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetooth try_OnDeviceRemoved
//go:noescape
func TryOnDeviceRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_OffDeviceRemoved
//go:noescape
func HasFuncOffDeviceRemoved() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_OffDeviceRemoved
//go:noescape
func FuncOffDeviceRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_OffDeviceRemoved
//go:noescape
func CallOffDeviceRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetooth try_OffDeviceRemoved
//go:noescape
func TryOffDeviceRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_HasOnDeviceRemoved
//go:noescape
func HasFuncHasOnDeviceRemoved() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_HasOnDeviceRemoved
//go:noescape
func FuncHasOnDeviceRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_HasOnDeviceRemoved
//go:noescape
func CallHasOnDeviceRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetooth try_HasOnDeviceRemoved
//go:noescape
func TryHasOnDeviceRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_StartDiscovery
//go:noescape
func HasFuncStartDiscovery() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_StartDiscovery
//go:noescape
func FuncStartDiscovery(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_StartDiscovery
//go:noescape
func CallStartDiscovery(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth try_StartDiscovery
//go:noescape
func TryStartDiscovery(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetooth has_StopDiscovery
//go:noescape
func HasFuncStopDiscovery() js.Ref

//go:wasmimport plat/js/webext/bluetooth func_StopDiscovery
//go:noescape
func FuncStopDiscovery(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth call_StopDiscovery
//go:noescape
func CallStopDiscovery(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetooth try_StopDiscovery
//go:noescape
func TryStopDiscovery(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
