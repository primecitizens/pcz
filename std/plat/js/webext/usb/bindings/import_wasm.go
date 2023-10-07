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

//go:wasmimport plat/js/webext/usb constof_TransferType
//go:noescape
func ConstOfTransferType(str js.Ref) uint32

//go:wasmimport plat/js/webext/usb constof_Direction
//go:noescape
func ConstOfDirection(str js.Ref) uint32

//go:wasmimport plat/js/webext/usb constof_SynchronizationType
//go:noescape
func ConstOfSynchronizationType(str js.Ref) uint32

//go:wasmimport plat/js/webext/usb constof_UsageType
//go:noescape
func ConstOfUsageType(str js.Ref) uint32

//go:wasmimport plat/js/webext/usb store_EndpointDescriptor
//go:noescape
func EndpointDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/usb load_EndpointDescriptor
//go:noescape
func EndpointDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/usb store_InterfaceDescriptor
//go:noescape
func InterfaceDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/usb load_InterfaceDescriptor
//go:noescape
func InterfaceDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/usb store_ConfigDescriptor
//go:noescape
func ConfigDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/usb load_ConfigDescriptor
//go:noescape
func ConfigDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/usb store_ConnectionHandle
//go:noescape
func ConnectionHandleJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/usb load_ConnectionHandle
//go:noescape
func ConnectionHandleJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/usb constof_Recipient
//go:noescape
func ConstOfRecipient(str js.Ref) uint32

//go:wasmimport plat/js/webext/usb constof_RequestType
//go:noescape
func ConstOfRequestType(str js.Ref) uint32

//go:wasmimport plat/js/webext/usb store_ControlTransferInfo
//go:noescape
func ControlTransferInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/usb load_ControlTransferInfo
//go:noescape
func ControlTransferInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/usb store_Device
//go:noescape
func DeviceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/usb load_Device
//go:noescape
func DeviceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/usb store_DeviceFilter
//go:noescape
func DeviceFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/usb load_DeviceFilter
//go:noescape
func DeviceFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/usb store_DevicePromptOptions
//go:noescape
func DevicePromptOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/usb load_DevicePromptOptions
//go:noescape
func DevicePromptOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/usb store_EnumerateDevicesAndRequestAccessOptions
//go:noescape
func EnumerateDevicesAndRequestAccessOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/usb load_EnumerateDevicesAndRequestAccessOptions
//go:noescape
func EnumerateDevicesAndRequestAccessOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/usb store_EnumerateDevicesOptions
//go:noescape
func EnumerateDevicesOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/usb load_EnumerateDevicesOptions
//go:noescape
func EnumerateDevicesOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/usb store_GenericTransferInfo
//go:noescape
func GenericTransferInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/usb load_GenericTransferInfo
//go:noescape
func GenericTransferInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/usb store_IsochronousTransferInfo
//go:noescape
func IsochronousTransferInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/usb load_IsochronousTransferInfo
//go:noescape
func IsochronousTransferInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/usb store_TransferResultInfo
//go:noescape
func TransferResultInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/usb load_TransferResultInfo
//go:noescape
func TransferResultInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/usb has_BulkTransfer
//go:noescape
func HasFuncBulkTransfer() js.Ref

//go:wasmimport plat/js/webext/usb func_BulkTransfer
//go:noescape
func FuncBulkTransfer(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_BulkTransfer
//go:noescape
func CallBulkTransfer(
	retPtr unsafe.Pointer,
	handle unsafe.Pointer,
	transferInfo unsafe.Pointer)

//go:wasmimport plat/js/webext/usb try_BulkTransfer
//go:noescape
func TryBulkTransfer(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle unsafe.Pointer,
	transferInfo unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_ClaimInterface
//go:noescape
func HasFuncClaimInterface() js.Ref

//go:wasmimport plat/js/webext/usb func_ClaimInterface
//go:noescape
func FuncClaimInterface(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_ClaimInterface
//go:noescape
func CallClaimInterface(
	retPtr unsafe.Pointer,
	handle unsafe.Pointer,
	interfaceNumber int32)

//go:wasmimport plat/js/webext/usb try_ClaimInterface
//go:noescape
func TryClaimInterface(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle unsafe.Pointer,
	interfaceNumber int32) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_CloseDevice
//go:noescape
func HasFuncCloseDevice() js.Ref

//go:wasmimport plat/js/webext/usb func_CloseDevice
//go:noescape
func FuncCloseDevice(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_CloseDevice
//go:noescape
func CallCloseDevice(
	retPtr unsafe.Pointer,
	handle unsafe.Pointer)

//go:wasmimport plat/js/webext/usb try_CloseDevice
//go:noescape
func TryCloseDevice(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_ControlTransfer
//go:noescape
func HasFuncControlTransfer() js.Ref

//go:wasmimport plat/js/webext/usb func_ControlTransfer
//go:noescape
func FuncControlTransfer(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_ControlTransfer
//go:noescape
func CallControlTransfer(
	retPtr unsafe.Pointer,
	handle unsafe.Pointer,
	transferInfo unsafe.Pointer)

//go:wasmimport plat/js/webext/usb try_ControlTransfer
//go:noescape
func TryControlTransfer(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle unsafe.Pointer,
	transferInfo unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_FindDevices
//go:noescape
func HasFuncFindDevices() js.Ref

//go:wasmimport plat/js/webext/usb func_FindDevices
//go:noescape
func FuncFindDevices(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_FindDevices
//go:noescape
func CallFindDevices(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/usb try_FindDevices
//go:noescape
func TryFindDevices(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_GetConfiguration
//go:noescape
func HasFuncGetConfiguration() js.Ref

//go:wasmimport plat/js/webext/usb func_GetConfiguration
//go:noescape
func FuncGetConfiguration(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_GetConfiguration
//go:noescape
func CallGetConfiguration(
	retPtr unsafe.Pointer,
	handle unsafe.Pointer)

//go:wasmimport plat/js/webext/usb try_GetConfiguration
//go:noescape
func TryGetConfiguration(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_GetConfigurations
//go:noescape
func HasFuncGetConfigurations() js.Ref

//go:wasmimport plat/js/webext/usb func_GetConfigurations
//go:noescape
func FuncGetConfigurations(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_GetConfigurations
//go:noescape
func CallGetConfigurations(
	retPtr unsafe.Pointer,
	device unsafe.Pointer)

//go:wasmimport plat/js/webext/usb try_GetConfigurations
//go:noescape
func TryGetConfigurations(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	device unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_GetDevices
//go:noescape
func HasFuncGetDevices() js.Ref

//go:wasmimport plat/js/webext/usb func_GetDevices
//go:noescape
func FuncGetDevices(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_GetDevices
//go:noescape
func CallGetDevices(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/usb try_GetDevices
//go:noescape
func TryGetDevices(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_GetUserSelectedDevices
//go:noescape
func HasFuncGetUserSelectedDevices() js.Ref

//go:wasmimport plat/js/webext/usb func_GetUserSelectedDevices
//go:noescape
func FuncGetUserSelectedDevices(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_GetUserSelectedDevices
//go:noescape
func CallGetUserSelectedDevices(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/usb try_GetUserSelectedDevices
//go:noescape
func TryGetUserSelectedDevices(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_InterruptTransfer
//go:noescape
func HasFuncInterruptTransfer() js.Ref

//go:wasmimport plat/js/webext/usb func_InterruptTransfer
//go:noescape
func FuncInterruptTransfer(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_InterruptTransfer
//go:noescape
func CallInterruptTransfer(
	retPtr unsafe.Pointer,
	handle unsafe.Pointer,
	transferInfo unsafe.Pointer)

//go:wasmimport plat/js/webext/usb try_InterruptTransfer
//go:noescape
func TryInterruptTransfer(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle unsafe.Pointer,
	transferInfo unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_IsochronousTransfer
//go:noescape
func HasFuncIsochronousTransfer() js.Ref

//go:wasmimport plat/js/webext/usb func_IsochronousTransfer
//go:noescape
func FuncIsochronousTransfer(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_IsochronousTransfer
//go:noescape
func CallIsochronousTransfer(
	retPtr unsafe.Pointer,
	handle unsafe.Pointer,
	transferInfo unsafe.Pointer)

//go:wasmimport plat/js/webext/usb try_IsochronousTransfer
//go:noescape
func TryIsochronousTransfer(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle unsafe.Pointer,
	transferInfo unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_ListInterfaces
//go:noescape
func HasFuncListInterfaces() js.Ref

//go:wasmimport plat/js/webext/usb func_ListInterfaces
//go:noescape
func FuncListInterfaces(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_ListInterfaces
//go:noescape
func CallListInterfaces(
	retPtr unsafe.Pointer,
	handle unsafe.Pointer)

//go:wasmimport plat/js/webext/usb try_ListInterfaces
//go:noescape
func TryListInterfaces(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_OnDeviceAdded
//go:noescape
func HasFuncOnDeviceAdded() js.Ref

//go:wasmimport plat/js/webext/usb func_OnDeviceAdded
//go:noescape
func FuncOnDeviceAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_OnDeviceAdded
//go:noescape
func CallOnDeviceAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/usb try_OnDeviceAdded
//go:noescape
func TryOnDeviceAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_OffDeviceAdded
//go:noescape
func HasFuncOffDeviceAdded() js.Ref

//go:wasmimport plat/js/webext/usb func_OffDeviceAdded
//go:noescape
func FuncOffDeviceAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_OffDeviceAdded
//go:noescape
func CallOffDeviceAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/usb try_OffDeviceAdded
//go:noescape
func TryOffDeviceAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_HasOnDeviceAdded
//go:noescape
func HasFuncHasOnDeviceAdded() js.Ref

//go:wasmimport plat/js/webext/usb func_HasOnDeviceAdded
//go:noescape
func FuncHasOnDeviceAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_HasOnDeviceAdded
//go:noescape
func CallHasOnDeviceAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/usb try_HasOnDeviceAdded
//go:noescape
func TryHasOnDeviceAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_OnDeviceRemoved
//go:noescape
func HasFuncOnDeviceRemoved() js.Ref

//go:wasmimport plat/js/webext/usb func_OnDeviceRemoved
//go:noescape
func FuncOnDeviceRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_OnDeviceRemoved
//go:noescape
func CallOnDeviceRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/usb try_OnDeviceRemoved
//go:noescape
func TryOnDeviceRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_OffDeviceRemoved
//go:noescape
func HasFuncOffDeviceRemoved() js.Ref

//go:wasmimport plat/js/webext/usb func_OffDeviceRemoved
//go:noescape
func FuncOffDeviceRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_OffDeviceRemoved
//go:noescape
func CallOffDeviceRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/usb try_OffDeviceRemoved
//go:noescape
func TryOffDeviceRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_HasOnDeviceRemoved
//go:noescape
func HasFuncHasOnDeviceRemoved() js.Ref

//go:wasmimport plat/js/webext/usb func_HasOnDeviceRemoved
//go:noescape
func FuncHasOnDeviceRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_HasOnDeviceRemoved
//go:noescape
func CallHasOnDeviceRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/usb try_HasOnDeviceRemoved
//go:noescape
func TryHasOnDeviceRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_OpenDevice
//go:noescape
func HasFuncOpenDevice() js.Ref

//go:wasmimport plat/js/webext/usb func_OpenDevice
//go:noescape
func FuncOpenDevice(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_OpenDevice
//go:noescape
func CallOpenDevice(
	retPtr unsafe.Pointer,
	device unsafe.Pointer)

//go:wasmimport plat/js/webext/usb try_OpenDevice
//go:noescape
func TryOpenDevice(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	device unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_ReleaseInterface
//go:noescape
func HasFuncReleaseInterface() js.Ref

//go:wasmimport plat/js/webext/usb func_ReleaseInterface
//go:noescape
func FuncReleaseInterface(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_ReleaseInterface
//go:noescape
func CallReleaseInterface(
	retPtr unsafe.Pointer,
	handle unsafe.Pointer,
	interfaceNumber int32)

//go:wasmimport plat/js/webext/usb try_ReleaseInterface
//go:noescape
func TryReleaseInterface(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle unsafe.Pointer,
	interfaceNumber int32) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_RequestAccess
//go:noescape
func HasFuncRequestAccess() js.Ref

//go:wasmimport plat/js/webext/usb func_RequestAccess
//go:noescape
func FuncRequestAccess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_RequestAccess
//go:noescape
func CallRequestAccess(
	retPtr unsafe.Pointer,
	device unsafe.Pointer,
	interfaceId int32)

//go:wasmimport plat/js/webext/usb try_RequestAccess
//go:noescape
func TryRequestAccess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	device unsafe.Pointer,
	interfaceId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_ResetDevice
//go:noescape
func HasFuncResetDevice() js.Ref

//go:wasmimport plat/js/webext/usb func_ResetDevice
//go:noescape
func FuncResetDevice(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_ResetDevice
//go:noescape
func CallResetDevice(
	retPtr unsafe.Pointer,
	handle unsafe.Pointer)

//go:wasmimport plat/js/webext/usb try_ResetDevice
//go:noescape
func TryResetDevice(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_SetConfiguration
//go:noescape
func HasFuncSetConfiguration() js.Ref

//go:wasmimport plat/js/webext/usb func_SetConfiguration
//go:noescape
func FuncSetConfiguration(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_SetConfiguration
//go:noescape
func CallSetConfiguration(
	retPtr unsafe.Pointer,
	handle unsafe.Pointer,
	configurationValue int32)

//go:wasmimport plat/js/webext/usb try_SetConfiguration
//go:noescape
func TrySetConfiguration(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle unsafe.Pointer,
	configurationValue int32) (ok js.Ref)

//go:wasmimport plat/js/webext/usb has_SetInterfaceAlternateSetting
//go:noescape
func HasFuncSetInterfaceAlternateSetting() js.Ref

//go:wasmimport plat/js/webext/usb func_SetInterfaceAlternateSetting
//go:noescape
func FuncSetInterfaceAlternateSetting(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usb call_SetInterfaceAlternateSetting
//go:noescape
func CallSetInterfaceAlternateSetting(
	retPtr unsafe.Pointer,
	handle unsafe.Pointer,
	interfaceNumber int32,
	alternateSetting int32)

//go:wasmimport plat/js/webext/usb try_SetInterfaceAlternateSetting
//go:noescape
func TrySetInterfaceAlternateSetting(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle unsafe.Pointer,
	interfaceNumber int32,
	alternateSetting int32) (ok js.Ref)
