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

//go:wasmimport plat/js/webext/bluetoothlowenergy constof_AdvertisementType
//go:noescape
func ConstOfAdvertisementType(str js.Ref) uint32

//go:wasmimport plat/js/webext/bluetoothlowenergy store_ManufacturerData
//go:noescape
func ManufacturerDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy load_ManufacturerData
//go:noescape
func ManufacturerDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy store_ServiceData
//go:noescape
func ServiceDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy load_ServiceData
//go:noescape
func ServiceDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy store_Advertisement
//go:noescape
func AdvertisementJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy load_Advertisement
//go:noescape
func AdvertisementJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy store_Service
//go:noescape
func ServiceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy load_Service
//go:noescape
func ServiceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy constof_CharacteristicProperty
//go:noescape
func ConstOfCharacteristicProperty(str js.Ref) uint32

//go:wasmimport plat/js/webext/bluetoothlowenergy store_Characteristic
//go:noescape
func CharacteristicJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy load_Characteristic
//go:noescape
func CharacteristicJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy store_ConnectProperties
//go:noescape
func ConnectPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy load_ConnectProperties
//go:noescape
func ConnectPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy constof_DescriptorPermission
//go:noescape
func ConstOfDescriptorPermission(str js.Ref) uint32

//go:wasmimport plat/js/webext/bluetoothlowenergy store_Descriptor
//go:noescape
func DescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy load_Descriptor
//go:noescape
func DescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy store_Device
//go:noescape
func DeviceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy load_Device
//go:noescape
func DeviceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy store_Notification
//go:noescape
func NotificationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy load_Notification
//go:noescape
func NotificationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy store_NotificationProperties
//go:noescape
func NotificationPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy load_NotificationProperties
//go:noescape
func NotificationPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy store_Request
//go:noescape
func RequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy load_Request
//go:noescape
func RequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy store_Response
//go:noescape
func ResponseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy load_Response
//go:noescape
func ResponseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy has_Connect
//go:noescape
func HasFuncConnect() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_Connect
//go:noescape
func FuncConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_Connect
//go:noescape
func CallConnect(
	retPtr unsafe.Pointer,
	deviceAddress js.Ref,
	properties unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_Connect
//go:noescape
func TryConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deviceAddress js.Ref,
	properties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_CreateCharacteristic
//go:noescape
func HasFuncCreateCharacteristic() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_CreateCharacteristic
//go:noescape
func FuncCreateCharacteristic(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_CreateCharacteristic
//go:noescape
func CallCreateCharacteristic(
	retPtr unsafe.Pointer,
	characteristic unsafe.Pointer,
	serviceId js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_CreateCharacteristic
//go:noescape
func TryCreateCharacteristic(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	characteristic unsafe.Pointer,
	serviceId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_CreateDescriptor
//go:noescape
func HasFuncCreateDescriptor() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_CreateDescriptor
//go:noescape
func FuncCreateDescriptor(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_CreateDescriptor
//go:noescape
func CallCreateDescriptor(
	retPtr unsafe.Pointer,
	descriptor unsafe.Pointer,
	characteristicId js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_CreateDescriptor
//go:noescape
func TryCreateDescriptor(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer,
	characteristicId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_CreateService
//go:noescape
func HasFuncCreateService() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_CreateService
//go:noescape
func FuncCreateService(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_CreateService
//go:noescape
func CallCreateService(
	retPtr unsafe.Pointer,
	service unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_CreateService
//go:noescape
func TryCreateService(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	service unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_Disconnect
//go:noescape
func HasFuncDisconnect() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_Disconnect
//go:noescape
func FuncDisconnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_Disconnect
//go:noescape
func CallDisconnect(
	retPtr unsafe.Pointer,
	deviceAddress js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_Disconnect
//go:noescape
func TryDisconnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deviceAddress js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_GetCharacteristic
//go:noescape
func HasFuncGetCharacteristic() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_GetCharacteristic
//go:noescape
func FuncGetCharacteristic(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_GetCharacteristic
//go:noescape
func CallGetCharacteristic(
	retPtr unsafe.Pointer,
	characteristicId js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_GetCharacteristic
//go:noescape
func TryGetCharacteristic(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	characteristicId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_GetCharacteristics
//go:noescape
func HasFuncGetCharacteristics() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_GetCharacteristics
//go:noescape
func FuncGetCharacteristics(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_GetCharacteristics
//go:noescape
func CallGetCharacteristics(
	retPtr unsafe.Pointer,
	serviceId js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_GetCharacteristics
//go:noescape
func TryGetCharacteristics(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	serviceId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_GetDescriptor
//go:noescape
func HasFuncGetDescriptor() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_GetDescriptor
//go:noescape
func FuncGetDescriptor(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_GetDescriptor
//go:noescape
func CallGetDescriptor(
	retPtr unsafe.Pointer,
	descriptorId js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_GetDescriptor
//go:noescape
func TryGetDescriptor(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptorId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_GetDescriptors
//go:noescape
func HasFuncGetDescriptors() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_GetDescriptors
//go:noescape
func FuncGetDescriptors(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_GetDescriptors
//go:noescape
func CallGetDescriptors(
	retPtr unsafe.Pointer,
	characteristicId js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_GetDescriptors
//go:noescape
func TryGetDescriptors(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	characteristicId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_GetIncludedServices
//go:noescape
func HasFuncGetIncludedServices() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_GetIncludedServices
//go:noescape
func FuncGetIncludedServices(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_GetIncludedServices
//go:noescape
func CallGetIncludedServices(
	retPtr unsafe.Pointer,
	serviceId js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_GetIncludedServices
//go:noescape
func TryGetIncludedServices(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	serviceId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_GetService
//go:noescape
func HasFuncGetService() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_GetService
//go:noescape
func FuncGetService(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_GetService
//go:noescape
func CallGetService(
	retPtr unsafe.Pointer,
	serviceId js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_GetService
//go:noescape
func TryGetService(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	serviceId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_GetServices
//go:noescape
func HasFuncGetServices() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_GetServices
//go:noescape
func FuncGetServices(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_GetServices
//go:noescape
func CallGetServices(
	retPtr unsafe.Pointer,
	deviceAddress js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_GetServices
//go:noescape
func TryGetServices(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deviceAddress js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_NotifyCharacteristicValueChanged
//go:noescape
func HasFuncNotifyCharacteristicValueChanged() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_NotifyCharacteristicValueChanged
//go:noescape
func FuncNotifyCharacteristicValueChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_NotifyCharacteristicValueChanged
//go:noescape
func CallNotifyCharacteristicValueChanged(
	retPtr unsafe.Pointer,
	characteristicId js.Ref,
	notification unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_NotifyCharacteristicValueChanged
//go:noescape
func TryNotifyCharacteristicValueChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	characteristicId js.Ref,
	notification unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OnCharacteristicReadRequest
//go:noescape
func HasFuncOnCharacteristicReadRequest() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OnCharacteristicReadRequest
//go:noescape
func FuncOnCharacteristicReadRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OnCharacteristicReadRequest
//go:noescape
func CallOnCharacteristicReadRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OnCharacteristicReadRequest
//go:noescape
func TryOnCharacteristicReadRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OffCharacteristicReadRequest
//go:noescape
func HasFuncOffCharacteristicReadRequest() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OffCharacteristicReadRequest
//go:noescape
func FuncOffCharacteristicReadRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OffCharacteristicReadRequest
//go:noescape
func CallOffCharacteristicReadRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OffCharacteristicReadRequest
//go:noescape
func TryOffCharacteristicReadRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_HasOnCharacteristicReadRequest
//go:noescape
func HasFuncHasOnCharacteristicReadRequest() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_HasOnCharacteristicReadRequest
//go:noescape
func FuncHasOnCharacteristicReadRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_HasOnCharacteristicReadRequest
//go:noescape
func CallHasOnCharacteristicReadRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_HasOnCharacteristicReadRequest
//go:noescape
func TryHasOnCharacteristicReadRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OnCharacteristicValueChanged
//go:noescape
func HasFuncOnCharacteristicValueChanged() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OnCharacteristicValueChanged
//go:noescape
func FuncOnCharacteristicValueChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OnCharacteristicValueChanged
//go:noescape
func CallOnCharacteristicValueChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OnCharacteristicValueChanged
//go:noescape
func TryOnCharacteristicValueChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OffCharacteristicValueChanged
//go:noescape
func HasFuncOffCharacteristicValueChanged() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OffCharacteristicValueChanged
//go:noescape
func FuncOffCharacteristicValueChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OffCharacteristicValueChanged
//go:noescape
func CallOffCharacteristicValueChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OffCharacteristicValueChanged
//go:noescape
func TryOffCharacteristicValueChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_HasOnCharacteristicValueChanged
//go:noescape
func HasFuncHasOnCharacteristicValueChanged() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_HasOnCharacteristicValueChanged
//go:noescape
func FuncHasOnCharacteristicValueChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_HasOnCharacteristicValueChanged
//go:noescape
func CallHasOnCharacteristicValueChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_HasOnCharacteristicValueChanged
//go:noescape
func TryHasOnCharacteristicValueChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OnCharacteristicWriteRequest
//go:noescape
func HasFuncOnCharacteristicWriteRequest() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OnCharacteristicWriteRequest
//go:noescape
func FuncOnCharacteristicWriteRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OnCharacteristicWriteRequest
//go:noescape
func CallOnCharacteristicWriteRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OnCharacteristicWriteRequest
//go:noescape
func TryOnCharacteristicWriteRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OffCharacteristicWriteRequest
//go:noescape
func HasFuncOffCharacteristicWriteRequest() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OffCharacteristicWriteRequest
//go:noescape
func FuncOffCharacteristicWriteRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OffCharacteristicWriteRequest
//go:noescape
func CallOffCharacteristicWriteRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OffCharacteristicWriteRequest
//go:noescape
func TryOffCharacteristicWriteRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_HasOnCharacteristicWriteRequest
//go:noescape
func HasFuncHasOnCharacteristicWriteRequest() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_HasOnCharacteristicWriteRequest
//go:noescape
func FuncHasOnCharacteristicWriteRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_HasOnCharacteristicWriteRequest
//go:noescape
func CallHasOnCharacteristicWriteRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_HasOnCharacteristicWriteRequest
//go:noescape
func TryHasOnCharacteristicWriteRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OnDescriptorReadRequest
//go:noescape
func HasFuncOnDescriptorReadRequest() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OnDescriptorReadRequest
//go:noescape
func FuncOnDescriptorReadRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OnDescriptorReadRequest
//go:noescape
func CallOnDescriptorReadRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OnDescriptorReadRequest
//go:noescape
func TryOnDescriptorReadRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OffDescriptorReadRequest
//go:noescape
func HasFuncOffDescriptorReadRequest() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OffDescriptorReadRequest
//go:noescape
func FuncOffDescriptorReadRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OffDescriptorReadRequest
//go:noescape
func CallOffDescriptorReadRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OffDescriptorReadRequest
//go:noescape
func TryOffDescriptorReadRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_HasOnDescriptorReadRequest
//go:noescape
func HasFuncHasOnDescriptorReadRequest() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_HasOnDescriptorReadRequest
//go:noescape
func FuncHasOnDescriptorReadRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_HasOnDescriptorReadRequest
//go:noescape
func CallHasOnDescriptorReadRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_HasOnDescriptorReadRequest
//go:noescape
func TryHasOnDescriptorReadRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OnDescriptorValueChanged
//go:noescape
func HasFuncOnDescriptorValueChanged() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OnDescriptorValueChanged
//go:noescape
func FuncOnDescriptorValueChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OnDescriptorValueChanged
//go:noescape
func CallOnDescriptorValueChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OnDescriptorValueChanged
//go:noescape
func TryOnDescriptorValueChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OffDescriptorValueChanged
//go:noescape
func HasFuncOffDescriptorValueChanged() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OffDescriptorValueChanged
//go:noescape
func FuncOffDescriptorValueChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OffDescriptorValueChanged
//go:noescape
func CallOffDescriptorValueChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OffDescriptorValueChanged
//go:noescape
func TryOffDescriptorValueChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_HasOnDescriptorValueChanged
//go:noescape
func HasFuncHasOnDescriptorValueChanged() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_HasOnDescriptorValueChanged
//go:noescape
func FuncHasOnDescriptorValueChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_HasOnDescriptorValueChanged
//go:noescape
func CallHasOnDescriptorValueChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_HasOnDescriptorValueChanged
//go:noescape
func TryHasOnDescriptorValueChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OnDescriptorWriteRequest
//go:noescape
func HasFuncOnDescriptorWriteRequest() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OnDescriptorWriteRequest
//go:noescape
func FuncOnDescriptorWriteRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OnDescriptorWriteRequest
//go:noescape
func CallOnDescriptorWriteRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OnDescriptorWriteRequest
//go:noescape
func TryOnDescriptorWriteRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OffDescriptorWriteRequest
//go:noescape
func HasFuncOffDescriptorWriteRequest() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OffDescriptorWriteRequest
//go:noescape
func FuncOffDescriptorWriteRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OffDescriptorWriteRequest
//go:noescape
func CallOffDescriptorWriteRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OffDescriptorWriteRequest
//go:noescape
func TryOffDescriptorWriteRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_HasOnDescriptorWriteRequest
//go:noescape
func HasFuncHasOnDescriptorWriteRequest() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_HasOnDescriptorWriteRequest
//go:noescape
func FuncHasOnDescriptorWriteRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_HasOnDescriptorWriteRequest
//go:noescape
func CallHasOnDescriptorWriteRequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_HasOnDescriptorWriteRequest
//go:noescape
func TryHasOnDescriptorWriteRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OnServiceAdded
//go:noescape
func HasFuncOnServiceAdded() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OnServiceAdded
//go:noescape
func FuncOnServiceAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OnServiceAdded
//go:noescape
func CallOnServiceAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OnServiceAdded
//go:noescape
func TryOnServiceAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OffServiceAdded
//go:noescape
func HasFuncOffServiceAdded() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OffServiceAdded
//go:noescape
func FuncOffServiceAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OffServiceAdded
//go:noescape
func CallOffServiceAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OffServiceAdded
//go:noescape
func TryOffServiceAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_HasOnServiceAdded
//go:noescape
func HasFuncHasOnServiceAdded() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_HasOnServiceAdded
//go:noescape
func FuncHasOnServiceAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_HasOnServiceAdded
//go:noescape
func CallHasOnServiceAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_HasOnServiceAdded
//go:noescape
func TryHasOnServiceAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OnServiceChanged
//go:noescape
func HasFuncOnServiceChanged() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OnServiceChanged
//go:noescape
func FuncOnServiceChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OnServiceChanged
//go:noescape
func CallOnServiceChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OnServiceChanged
//go:noescape
func TryOnServiceChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OffServiceChanged
//go:noescape
func HasFuncOffServiceChanged() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OffServiceChanged
//go:noescape
func FuncOffServiceChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OffServiceChanged
//go:noescape
func CallOffServiceChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OffServiceChanged
//go:noescape
func TryOffServiceChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_HasOnServiceChanged
//go:noescape
func HasFuncHasOnServiceChanged() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_HasOnServiceChanged
//go:noescape
func FuncHasOnServiceChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_HasOnServiceChanged
//go:noescape
func CallHasOnServiceChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_HasOnServiceChanged
//go:noescape
func TryHasOnServiceChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OnServiceRemoved
//go:noescape
func HasFuncOnServiceRemoved() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OnServiceRemoved
//go:noescape
func FuncOnServiceRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OnServiceRemoved
//go:noescape
func CallOnServiceRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OnServiceRemoved
//go:noescape
func TryOnServiceRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_OffServiceRemoved
//go:noescape
func HasFuncOffServiceRemoved() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_OffServiceRemoved
//go:noescape
func FuncOffServiceRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_OffServiceRemoved
//go:noescape
func CallOffServiceRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_OffServiceRemoved
//go:noescape
func TryOffServiceRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_HasOnServiceRemoved
//go:noescape
func HasFuncHasOnServiceRemoved() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_HasOnServiceRemoved
//go:noescape
func FuncHasOnServiceRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_HasOnServiceRemoved
//go:noescape
func CallHasOnServiceRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_HasOnServiceRemoved
//go:noescape
func TryHasOnServiceRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_ReadCharacteristicValue
//go:noescape
func HasFuncReadCharacteristicValue() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_ReadCharacteristicValue
//go:noescape
func FuncReadCharacteristicValue(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_ReadCharacteristicValue
//go:noescape
func CallReadCharacteristicValue(
	retPtr unsafe.Pointer,
	characteristicId js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_ReadCharacteristicValue
//go:noescape
func TryReadCharacteristicValue(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	characteristicId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_ReadDescriptorValue
//go:noescape
func HasFuncReadDescriptorValue() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_ReadDescriptorValue
//go:noescape
func FuncReadDescriptorValue(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_ReadDescriptorValue
//go:noescape
func CallReadDescriptorValue(
	retPtr unsafe.Pointer,
	descriptorId js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_ReadDescriptorValue
//go:noescape
func TryReadDescriptorValue(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptorId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_RegisterAdvertisement
//go:noescape
func HasFuncRegisterAdvertisement() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_RegisterAdvertisement
//go:noescape
func FuncRegisterAdvertisement(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_RegisterAdvertisement
//go:noescape
func CallRegisterAdvertisement(
	retPtr unsafe.Pointer,
	advertisement unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_RegisterAdvertisement
//go:noescape
func TryRegisterAdvertisement(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	advertisement unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_RegisterService
//go:noescape
func HasFuncRegisterService() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_RegisterService
//go:noescape
func FuncRegisterService(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_RegisterService
//go:noescape
func CallRegisterService(
	retPtr unsafe.Pointer,
	serviceId js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_RegisterService
//go:noescape
func TryRegisterService(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	serviceId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_RemoveService
//go:noescape
func HasFuncRemoveService() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_RemoveService
//go:noescape
func FuncRemoveService(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_RemoveService
//go:noescape
func CallRemoveService(
	retPtr unsafe.Pointer,
	serviceId js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_RemoveService
//go:noescape
func TryRemoveService(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	serviceId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_ResetAdvertising
//go:noescape
func HasFuncResetAdvertising() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_ResetAdvertising
//go:noescape
func FuncResetAdvertising(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_ResetAdvertising
//go:noescape
func CallResetAdvertising(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_ResetAdvertising
//go:noescape
func TryResetAdvertising(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_SendRequestResponse
//go:noescape
func HasFuncSendRequestResponse() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_SendRequestResponse
//go:noescape
func FuncSendRequestResponse(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_SendRequestResponse
//go:noescape
func CallSendRequestResponse(
	retPtr unsafe.Pointer,
	response unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_SendRequestResponse
//go:noescape
func TrySendRequestResponse(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	response unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_SetAdvertisingInterval
//go:noescape
func HasFuncSetAdvertisingInterval() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_SetAdvertisingInterval
//go:noescape
func FuncSetAdvertisingInterval(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_SetAdvertisingInterval
//go:noescape
func CallSetAdvertisingInterval(
	retPtr unsafe.Pointer,
	minInterval int32,
	maxInterval int32)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_SetAdvertisingInterval
//go:noescape
func TrySetAdvertisingInterval(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	minInterval int32,
	maxInterval int32) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_StartCharacteristicNotifications
//go:noescape
func HasFuncStartCharacteristicNotifications() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_StartCharacteristicNotifications
//go:noescape
func FuncStartCharacteristicNotifications(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_StartCharacteristicNotifications
//go:noescape
func CallStartCharacteristicNotifications(
	retPtr unsafe.Pointer,
	characteristicId js.Ref,
	properties unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_StartCharacteristicNotifications
//go:noescape
func TryStartCharacteristicNotifications(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	characteristicId js.Ref,
	properties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_StopCharacteristicNotifications
//go:noescape
func HasFuncStopCharacteristicNotifications() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_StopCharacteristicNotifications
//go:noescape
func FuncStopCharacteristicNotifications(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_StopCharacteristicNotifications
//go:noescape
func CallStopCharacteristicNotifications(
	retPtr unsafe.Pointer,
	characteristicId js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_StopCharacteristicNotifications
//go:noescape
func TryStopCharacteristicNotifications(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	characteristicId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_UnregisterAdvertisement
//go:noescape
func HasFuncUnregisterAdvertisement() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_UnregisterAdvertisement
//go:noescape
func FuncUnregisterAdvertisement(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_UnregisterAdvertisement
//go:noescape
func CallUnregisterAdvertisement(
	retPtr unsafe.Pointer,
	advertisementId int32)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_UnregisterAdvertisement
//go:noescape
func TryUnregisterAdvertisement(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	advertisementId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_UnregisterService
//go:noescape
func HasFuncUnregisterService() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_UnregisterService
//go:noescape
func FuncUnregisterService(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_UnregisterService
//go:noescape
func CallUnregisterService(
	retPtr unsafe.Pointer,
	serviceId js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_UnregisterService
//go:noescape
func TryUnregisterService(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	serviceId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_WriteCharacteristicValue
//go:noescape
func HasFuncWriteCharacteristicValue() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_WriteCharacteristicValue
//go:noescape
func FuncWriteCharacteristicValue(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_WriteCharacteristicValue
//go:noescape
func CallWriteCharacteristicValue(
	retPtr unsafe.Pointer,
	characteristicId js.Ref,
	value js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_WriteCharacteristicValue
//go:noescape
func TryWriteCharacteristicValue(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	characteristicId js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy has_WriteDescriptorValue
//go:noescape
func HasFuncWriteDescriptorValue() js.Ref

//go:wasmimport plat/js/webext/bluetoothlowenergy func_WriteDescriptorValue
//go:noescape
func FuncWriteDescriptorValue(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothlowenergy call_WriteDescriptorValue
//go:noescape
func CallWriteDescriptorValue(
	retPtr unsafe.Pointer,
	descriptorId js.Ref,
	value js.Ref)

//go:wasmimport plat/js/webext/bluetoothlowenergy try_WriteDescriptorValue
//go:noescape
func TryWriteDescriptorValue(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptorId js.Ref,
	value js.Ref) (ok js.Ref)
