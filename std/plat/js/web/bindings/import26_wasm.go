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

//go:wasmimport plat/js/web constof_USBRecipient
//go:noescape
func ConstOfUSBRecipient(str js.Ref) uint32

//go:wasmimport plat/js/web store_USBControlTransferParameters
//go:noescape
func USBControlTransferParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_USBControlTransferParameters
//go:noescape
func USBControlTransferParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_USBOutTransferResult_USBOutTransferResult
//go:noescape
func NewUSBOutTransferResultByUSBOutTransferResult(
	status uint32,
	bytesWritten uint32) js.Ref

//go:wasmimport plat/js/web new_USBOutTransferResult_USBOutTransferResult1
//go:noescape
func NewUSBOutTransferResultByUSBOutTransferResult1(
	status uint32) js.Ref

//go:wasmimport plat/js/web get_USBOutTransferResult_BytesWritten
//go:noescape
func GetUSBOutTransferResultBytesWritten(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBOutTransferResult_Status
//go:noescape
func GetUSBOutTransferResultStatus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_USBDirection
//go:noescape
func ConstOfUSBDirection(str js.Ref) uint32

//go:wasmimport plat/js/web new_USBIsochronousInTransferPacket_USBIsochronousInTransferPacket
//go:noescape
func NewUSBIsochronousInTransferPacketByUSBIsochronousInTransferPacket(
	status uint32,
	data js.Ref) js.Ref

//go:wasmimport plat/js/web new_USBIsochronousInTransferPacket_USBIsochronousInTransferPacket1
//go:noescape
func NewUSBIsochronousInTransferPacketByUSBIsochronousInTransferPacket1(
	status uint32) js.Ref

//go:wasmimport plat/js/web get_USBIsochronousInTransferPacket_Data
//go:noescape
func GetUSBIsochronousInTransferPacketData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBIsochronousInTransferPacket_Status
//go:noescape
func GetUSBIsochronousInTransferPacketStatus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_USBIsochronousInTransferResult_USBIsochronousInTransferResult
//go:noescape
func NewUSBIsochronousInTransferResultByUSBIsochronousInTransferResult(
	packets js.Ref,
	data js.Ref) js.Ref

//go:wasmimport plat/js/web new_USBIsochronousInTransferResult_USBIsochronousInTransferResult1
//go:noescape
func NewUSBIsochronousInTransferResultByUSBIsochronousInTransferResult1(
	packets js.Ref) js.Ref

//go:wasmimport plat/js/web get_USBIsochronousInTransferResult_Data
//go:noescape
func GetUSBIsochronousInTransferResultData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBIsochronousInTransferResult_Packets
//go:noescape
func GetUSBIsochronousInTransferResultPackets(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_USBIsochronousOutTransferPacket_USBIsochronousOutTransferPacket
//go:noescape
func NewUSBIsochronousOutTransferPacketByUSBIsochronousOutTransferPacket(
	status uint32,
	bytesWritten uint32) js.Ref

//go:wasmimport plat/js/web new_USBIsochronousOutTransferPacket_USBIsochronousOutTransferPacket1
//go:noescape
func NewUSBIsochronousOutTransferPacketByUSBIsochronousOutTransferPacket1(
	status uint32) js.Ref

//go:wasmimport plat/js/web get_USBIsochronousOutTransferPacket_BytesWritten
//go:noescape
func GetUSBIsochronousOutTransferPacketBytesWritten(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBIsochronousOutTransferPacket_Status
//go:noescape
func GetUSBIsochronousOutTransferPacketStatus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_USBIsochronousOutTransferResult_USBIsochronousOutTransferResult
//go:noescape
func NewUSBIsochronousOutTransferResultByUSBIsochronousOutTransferResult(
	packets js.Ref) js.Ref

//go:wasmimport plat/js/web get_USBIsochronousOutTransferResult_Packets
//go:noescape
func GetUSBIsochronousOutTransferResultPackets(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_USBEndpointType
//go:noescape
func ConstOfUSBEndpointType(str js.Ref) uint32

//go:wasmimport plat/js/web new_USBEndpoint_USBEndpoint
//go:noescape
func NewUSBEndpointByUSBEndpoint(
	alternate js.Ref,
	endpointNumber uint32,
	direction uint32) js.Ref

//go:wasmimport plat/js/web get_USBEndpoint_EndpointNumber
//go:noescape
func GetUSBEndpointEndpointNumber(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBEndpoint_Direction
//go:noescape
func GetUSBEndpointDirection(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBEndpoint_Type
//go:noescape
func GetUSBEndpointType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBEndpoint_PacketSize
//go:noescape
func GetUSBEndpointPacketSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_USBAlternateInterface_USBAlternateInterface
//go:noescape
func NewUSBAlternateInterfaceByUSBAlternateInterface(
	deviceInterface js.Ref,
	alternateSetting uint32) js.Ref

//go:wasmimport plat/js/web get_USBAlternateInterface_AlternateSetting
//go:noescape
func GetUSBAlternateInterfaceAlternateSetting(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBAlternateInterface_InterfaceClass
//go:noescape
func GetUSBAlternateInterfaceInterfaceClass(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBAlternateInterface_InterfaceSubclass
//go:noescape
func GetUSBAlternateInterfaceInterfaceSubclass(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBAlternateInterface_InterfaceProtocol
//go:noescape
func GetUSBAlternateInterfaceInterfaceProtocol(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBAlternateInterface_InterfaceName
//go:noescape
func GetUSBAlternateInterfaceInterfaceName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBAlternateInterface_Endpoints
//go:noescape
func GetUSBAlternateInterfaceEndpoints(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_USBInterface_USBInterface
//go:noescape
func NewUSBInterfaceByUSBInterface(
	configuration js.Ref,
	interfaceNumber uint32) js.Ref

//go:wasmimport plat/js/web get_USBInterface_InterfaceNumber
//go:noescape
func GetUSBInterfaceInterfaceNumber(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBInterface_Alternate
//go:noescape
func GetUSBInterfaceAlternate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBInterface_Alternates
//go:noescape
func GetUSBInterfaceAlternates(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBInterface_Claimed
//go:noescape
func GetUSBInterfaceClaimed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_USBConfiguration_USBConfiguration
//go:noescape
func NewUSBConfigurationByUSBConfiguration(
	device js.Ref,
	configurationValue uint32) js.Ref

//go:wasmimport plat/js/web get_USBConfiguration_ConfigurationValue
//go:noescape
func GetUSBConfigurationConfigurationValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBConfiguration_ConfigurationName
//go:noescape
func GetUSBConfigurationConfigurationName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBConfiguration_Interfaces
//go:noescape
func GetUSBConfigurationInterfaces(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_UsbVersionMajor
//go:noescape
func GetUSBDeviceUsbVersionMajor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_UsbVersionMinor
//go:noescape
func GetUSBDeviceUsbVersionMinor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_UsbVersionSubminor
//go:noescape
func GetUSBDeviceUsbVersionSubminor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_DeviceClass
//go:noescape
func GetUSBDeviceDeviceClass(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_DeviceSubclass
//go:noescape
func GetUSBDeviceDeviceSubclass(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_DeviceProtocol
//go:noescape
func GetUSBDeviceDeviceProtocol(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_VendorId
//go:noescape
func GetUSBDeviceVendorId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_ProductId
//go:noescape
func GetUSBDeviceProductId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_DeviceVersionMajor
//go:noescape
func GetUSBDeviceDeviceVersionMajor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_DeviceVersionMinor
//go:noescape
func GetUSBDeviceDeviceVersionMinor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_DeviceVersionSubminor
//go:noescape
func GetUSBDeviceDeviceVersionSubminor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_ManufacturerName
//go:noescape
func GetUSBDeviceManufacturerName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_ProductName
//go:noescape
func GetUSBDeviceProductName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_SerialNumber
//go:noescape
func GetUSBDeviceSerialNumber(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_Configuration
//go:noescape
func GetUSBDeviceConfiguration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_Configurations
//go:noescape
func GetUSBDeviceConfigurations(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBDevice_Opened
//go:noescape
func GetUSBDeviceOpened(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_Open
//go:noescape
func HasFuncUSBDeviceOpen(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_Open
//go:noescape
func FuncUSBDeviceOpen(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_Open
//go:noescape
func CallUSBDeviceOpen(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_USBDevice_Open
//go:noescape
func TryUSBDeviceOpen(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_Close
//go:noescape
func HasFuncUSBDeviceClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_Close
//go:noescape
func FuncUSBDeviceClose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_Close
//go:noescape
func CallUSBDeviceClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_USBDevice_Close
//go:noescape
func TryUSBDeviceClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_Forget
//go:noescape
func HasFuncUSBDeviceForget(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_Forget
//go:noescape
func FuncUSBDeviceForget(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_Forget
//go:noescape
func CallUSBDeviceForget(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_USBDevice_Forget
//go:noescape
func TryUSBDeviceForget(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_SelectConfiguration
//go:noescape
func HasFuncUSBDeviceSelectConfiguration(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_SelectConfiguration
//go:noescape
func FuncUSBDeviceSelectConfiguration(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_SelectConfiguration
//go:noescape
func CallUSBDeviceSelectConfiguration(
	this js.Ref, retPtr unsafe.Pointer,
	configurationValue uint32)

//go:wasmimport plat/js/web try_USBDevice_SelectConfiguration
//go:noescape
func TryUSBDeviceSelectConfiguration(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	configurationValue uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_ClaimInterface
//go:noescape
func HasFuncUSBDeviceClaimInterface(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_ClaimInterface
//go:noescape
func FuncUSBDeviceClaimInterface(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_ClaimInterface
//go:noescape
func CallUSBDeviceClaimInterface(
	this js.Ref, retPtr unsafe.Pointer,
	interfaceNumber uint32)

//go:wasmimport plat/js/web try_USBDevice_ClaimInterface
//go:noescape
func TryUSBDeviceClaimInterface(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	interfaceNumber uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_ReleaseInterface
//go:noescape
func HasFuncUSBDeviceReleaseInterface(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_ReleaseInterface
//go:noescape
func FuncUSBDeviceReleaseInterface(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_ReleaseInterface
//go:noescape
func CallUSBDeviceReleaseInterface(
	this js.Ref, retPtr unsafe.Pointer,
	interfaceNumber uint32)

//go:wasmimport plat/js/web try_USBDevice_ReleaseInterface
//go:noescape
func TryUSBDeviceReleaseInterface(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	interfaceNumber uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_SelectAlternateInterface
//go:noescape
func HasFuncUSBDeviceSelectAlternateInterface(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_SelectAlternateInterface
//go:noescape
func FuncUSBDeviceSelectAlternateInterface(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_SelectAlternateInterface
//go:noescape
func CallUSBDeviceSelectAlternateInterface(
	this js.Ref, retPtr unsafe.Pointer,
	interfaceNumber uint32,
	alternateSetting uint32)

//go:wasmimport plat/js/web try_USBDevice_SelectAlternateInterface
//go:noescape
func TryUSBDeviceSelectAlternateInterface(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	interfaceNumber uint32,
	alternateSetting uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_ControlTransferIn
//go:noescape
func HasFuncUSBDeviceControlTransferIn(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_ControlTransferIn
//go:noescape
func FuncUSBDeviceControlTransferIn(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_ControlTransferIn
//go:noescape
func CallUSBDeviceControlTransferIn(
	this js.Ref, retPtr unsafe.Pointer,
	setup unsafe.Pointer,
	length uint32)

//go:wasmimport plat/js/web try_USBDevice_ControlTransferIn
//go:noescape
func TryUSBDeviceControlTransferIn(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	setup unsafe.Pointer,
	length uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_ControlTransferOut
//go:noescape
func HasFuncUSBDeviceControlTransferOut(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_ControlTransferOut
//go:noescape
func FuncUSBDeviceControlTransferOut(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_ControlTransferOut
//go:noescape
func CallUSBDeviceControlTransferOut(
	this js.Ref, retPtr unsafe.Pointer,
	setup unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_USBDevice_ControlTransferOut
//go:noescape
func TryUSBDeviceControlTransferOut(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	setup unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_ControlTransferOut1
//go:noescape
func HasFuncUSBDeviceControlTransferOut1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_ControlTransferOut1
//go:noescape
func FuncUSBDeviceControlTransferOut1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_ControlTransferOut1
//go:noescape
func CallUSBDeviceControlTransferOut1(
	this js.Ref, retPtr unsafe.Pointer,
	setup unsafe.Pointer)

//go:wasmimport plat/js/web try_USBDevice_ControlTransferOut1
//go:noescape
func TryUSBDeviceControlTransferOut1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	setup unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_ClearHalt
//go:noescape
func HasFuncUSBDeviceClearHalt(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_ClearHalt
//go:noescape
func FuncUSBDeviceClearHalt(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_ClearHalt
//go:noescape
func CallUSBDeviceClearHalt(
	this js.Ref, retPtr unsafe.Pointer,
	direction uint32,
	endpointNumber uint32)

//go:wasmimport plat/js/web try_USBDevice_ClearHalt
//go:noescape
func TryUSBDeviceClearHalt(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	direction uint32,
	endpointNumber uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_TransferIn
//go:noescape
func HasFuncUSBDeviceTransferIn(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_TransferIn
//go:noescape
func FuncUSBDeviceTransferIn(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_TransferIn
//go:noescape
func CallUSBDeviceTransferIn(
	this js.Ref, retPtr unsafe.Pointer,
	endpointNumber uint32,
	length uint32)

//go:wasmimport plat/js/web try_USBDevice_TransferIn
//go:noescape
func TryUSBDeviceTransferIn(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	endpointNumber uint32,
	length uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_TransferOut
//go:noescape
func HasFuncUSBDeviceTransferOut(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_TransferOut
//go:noescape
func FuncUSBDeviceTransferOut(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_TransferOut
//go:noescape
func CallUSBDeviceTransferOut(
	this js.Ref, retPtr unsafe.Pointer,
	endpointNumber uint32,
	data js.Ref)

//go:wasmimport plat/js/web try_USBDevice_TransferOut
//go:noescape
func TryUSBDeviceTransferOut(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	endpointNumber uint32,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_IsochronousTransferIn
//go:noescape
func HasFuncUSBDeviceIsochronousTransferIn(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_IsochronousTransferIn
//go:noescape
func FuncUSBDeviceIsochronousTransferIn(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_IsochronousTransferIn
//go:noescape
func CallUSBDeviceIsochronousTransferIn(
	this js.Ref, retPtr unsafe.Pointer,
	endpointNumber uint32,
	packetLengths js.Ref)

//go:wasmimport plat/js/web try_USBDevice_IsochronousTransferIn
//go:noescape
func TryUSBDeviceIsochronousTransferIn(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	endpointNumber uint32,
	packetLengths js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_IsochronousTransferOut
//go:noescape
func HasFuncUSBDeviceIsochronousTransferOut(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_IsochronousTransferOut
//go:noescape
func FuncUSBDeviceIsochronousTransferOut(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_IsochronousTransferOut
//go:noescape
func CallUSBDeviceIsochronousTransferOut(
	this js.Ref, retPtr unsafe.Pointer,
	endpointNumber uint32,
	data js.Ref,
	packetLengths js.Ref)

//go:wasmimport plat/js/web try_USBDevice_IsochronousTransferOut
//go:noescape
func TryUSBDeviceIsochronousTransferOut(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	endpointNumber uint32,
	data js.Ref,
	packetLengths js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_USBDevice_Reset
//go:noescape
func HasFuncUSBDeviceReset(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USBDevice_Reset
//go:noescape
func FuncUSBDeviceReset(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USBDevice_Reset
//go:noescape
func CallUSBDeviceReset(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_USBDevice_Reset
//go:noescape
func TryUSBDeviceReset(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_USBDeviceFilter
//go:noescape
func USBDeviceFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_USBDeviceFilter
//go:noescape
func USBDeviceFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_USBDeviceRequestOptions
//go:noescape
func USBDeviceRequestOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_USBDeviceRequestOptions
//go:noescape
func USBDeviceRequestOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_USB_GetDevices
//go:noescape
func HasFuncUSBGetDevices(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USB_GetDevices
//go:noescape
func FuncUSBGetDevices(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USB_GetDevices
//go:noescape
func CallUSBGetDevices(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_USB_GetDevices
//go:noescape
func TryUSBGetDevices(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_USB_RequestDevice
//go:noescape
func HasFuncUSBRequestDevice(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_USB_RequestDevice
//go:noescape
func FuncUSBRequestDevice(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_USB_RequestDevice
//go:noescape
func CallUSBRequestDevice(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_USB_RequestDevice
//go:noescape
func TryUSBRequestDevice(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_EpubReadingSystem_HasFeature
//go:noescape
func HasFuncEpubReadingSystemHasFeature(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EpubReadingSystem_HasFeature
//go:noescape
func FuncEpubReadingSystemHasFeature(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EpubReadingSystem_HasFeature
//go:noescape
func CallEpubReadingSystemHasFeature(
	this js.Ref, retPtr unsafe.Pointer,
	feature js.Ref,
	version js.Ref)

//go:wasmimport plat/js/web try_EpubReadingSystem_HasFeature
//go:noescape
func TryEpubReadingSystemHasFeature(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	feature js.Ref,
	version js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_EpubReadingSystem_HasFeature1
//go:noescape
func HasFuncEpubReadingSystemHasFeature1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EpubReadingSystem_HasFeature1
//go:noescape
func FuncEpubReadingSystemHasFeature1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EpubReadingSystem_HasFeature1
//go:noescape
func CallEpubReadingSystemHasFeature1(
	this js.Ref, retPtr unsafe.Pointer,
	feature js.Ref)

//go:wasmimport plat/js/web try_EpubReadingSystem_HasFeature1
//go:noescape
func TryEpubReadingSystemHasFeature1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	feature js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_XRSessionMode
//go:noescape
func ConstOfXRSessionMode(str js.Ref) uint32

//go:wasmimport plat/js/web store_XRWebGLLayerInit
//go:noescape
func XRWebGLLayerInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRWebGLLayerInit
//go:noescape
func XRWebGLLayerInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_XRViewport_X
//go:noescape
func GetXRViewportX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRViewport_Y
//go:noescape
func GetXRViewportY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRViewport_Width
//go:noescape
func GetXRViewportWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRViewport_Height
//go:noescape
func GetXRViewportHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_XREye
//go:noescape
func ConstOfXREye(str js.Ref) uint32

//go:wasmimport plat/js/web new_XRRigidTransform_XRRigidTransform
//go:noescape
func NewXRRigidTransformByXRRigidTransform(
	position unsafe.Pointer,
	orientation unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_XRRigidTransform_XRRigidTransform1
//go:noescape
func NewXRRigidTransformByXRRigidTransform1(
	position unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_XRRigidTransform_XRRigidTransform2
//go:noescape
func NewXRRigidTransformByXRRigidTransform2() js.Ref

//go:wasmimport plat/js/web get_XRRigidTransform_Position
//go:noescape
func GetXRRigidTransformPosition(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRRigidTransform_Orientation
//go:noescape
func GetXRRigidTransformOrientation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRRigidTransform_Matrix
//go:noescape
func GetXRRigidTransformMatrix(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRRigidTransform_Inverse
//go:noescape
func GetXRRigidTransformInverse(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRCamera_Width
//go:noescape
func GetXRCameraWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRCamera_Height
//go:noescape
func GetXRCameraHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRView_Eye
//go:noescape
func GetXRViewEye(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRView_ProjectionMatrix
//go:noescape
func GetXRViewProjectionMatrix(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRView_Transform
//go:noescape
func GetXRViewTransform(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRView_RecommendedViewportScale
//go:noescape
func GetXRViewRecommendedViewportScale(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRView_IsFirstPersonObserver
//go:noescape
func GetXRViewIsFirstPersonObserver(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRView_Camera
//go:noescape
func GetXRViewCamera(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRView_RequestViewportScale
//go:noescape
func HasFuncXRViewRequestViewportScale(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRView_RequestViewportScale
//go:noescape
func FuncXRViewRequestViewportScale(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRView_RequestViewportScale
//go:noescape
func CallXRViewRequestViewportScale(
	this js.Ref, retPtr unsafe.Pointer,
	scale float64)

//go:wasmimport plat/js/web try_XRView_RequestViewportScale
//go:noescape
func TryXRViewRequestViewportScale(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scale float64) (ok js.Ref)

//go:wasmimport plat/js/web new_XRWebGLLayer_XRWebGLLayer
//go:noescape
func NewXRWebGLLayerByXRWebGLLayer(
	session js.Ref,
	context js.Ref,
	layerInit unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_XRWebGLLayer_XRWebGLLayer1
//go:noescape
func NewXRWebGLLayerByXRWebGLLayer1(
	session js.Ref,
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_XRWebGLLayer_Antialias
//go:noescape
func GetXRWebGLLayerAntialias(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLLayer_IgnoreDepthValues
//go:noescape
func GetXRWebGLLayerIgnoreDepthValues(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLLayer_FixedFoveation
//go:noescape
func GetXRWebGLLayerFixedFoveation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XRWebGLLayer_FixedFoveation
//go:noescape
func SetXRWebGLLayerFixedFoveation(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XRWebGLLayer_Framebuffer
//go:noescape
func GetXRWebGLLayerFramebuffer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLLayer_FramebufferWidth
//go:noescape
func GetXRWebGLLayerFramebufferWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLLayer_FramebufferHeight
//go:noescape
func GetXRWebGLLayerFramebufferHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLLayer_GetViewport
//go:noescape
func HasFuncXRWebGLLayerGetViewport(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLLayer_GetViewport
//go:noescape
func FuncXRWebGLLayerGetViewport(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLLayer_GetViewport
//go:noescape
func CallXRWebGLLayerGetViewport(
	this js.Ref, retPtr unsafe.Pointer,
	view js.Ref)

//go:wasmimport plat/js/web try_XRWebGLLayer_GetViewport
//go:noescape
func TryXRWebGLLayerGetViewport(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	view js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLLayer_GetNativeFramebufferScaleFactor
//go:noescape
func HasFuncXRWebGLLayerGetNativeFramebufferScaleFactor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLLayer_GetNativeFramebufferScaleFactor
//go:noescape
func FuncXRWebGLLayerGetNativeFramebufferScaleFactor(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLLayer_GetNativeFramebufferScaleFactor
//go:noescape
func CallXRWebGLLayerGetNativeFramebufferScaleFactor(
	this js.Ref, retPtr unsafe.Pointer,
	session js.Ref)

//go:wasmimport plat/js/web try_XRWebGLLayer_GetNativeFramebufferScaleFactor
//go:noescape
func TryXRWebGLLayerGetNativeFramebufferScaleFactor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	session js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_XRRenderStateInit
//go:noescape
func XRRenderStateInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRRenderStateInit
//go:noescape
func XRRenderStateInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_XRReferenceSpace_GetOffsetReferenceSpace
//go:noescape
func HasFuncXRReferenceSpaceGetOffsetReferenceSpace(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRReferenceSpace_GetOffsetReferenceSpace
//go:noescape
func FuncXRReferenceSpaceGetOffsetReferenceSpace(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRReferenceSpace_GetOffsetReferenceSpace
//go:noescape
func CallXRReferenceSpaceGetOffsetReferenceSpace(
	this js.Ref, retPtr unsafe.Pointer,
	originOffset js.Ref)

//go:wasmimport plat/js/web try_XRReferenceSpace_GetOffsetReferenceSpace
//go:noescape
func TryXRReferenceSpaceGetOffsetReferenceSpace(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	originOffset js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_XRReferenceSpaceType
//go:noescape
func ConstOfXRReferenceSpaceType(str js.Ref) uint32

//go:wasmimport plat/js/web get_XRViewerPose_Views
//go:noescape
func GetXRViewerPoseViews(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRPose_Transform
//go:noescape
func GetXRPoseTransform(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRPose_LinearVelocity
//go:noescape
func GetXRPoseLinearVelocity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRPose_AngularVelocity
//go:noescape
func GetXRPoseAngularVelocity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRPose_EmulatedPosition
//go:noescape
func GetXRPoseEmulatedPosition(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRAnchor_AnchorSpace
//go:noescape
func GetXRAnchorAnchorSpace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRAnchor_RequestPersistentHandle
//go:noescape
func HasFuncXRAnchorRequestPersistentHandle(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRAnchor_RequestPersistentHandle
//go:noescape
func FuncXRAnchorRequestPersistentHandle(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRAnchor_RequestPersistentHandle
//go:noescape
func CallXRAnchorRequestPersistentHandle(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XRAnchor_RequestPersistentHandle
//go:noescape
func TryXRAnchorRequestPersistentHandle(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRAnchor_Delete
//go:noescape
func HasFuncXRAnchorDelete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRAnchor_Delete
//go:noescape
func FuncXRAnchorDelete(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRAnchor_Delete
//go:noescape
func CallXRAnchorDelete(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XRAnchor_Delete
//go:noescape
func TryXRAnchorDelete(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRLightEstimate_SphericalHarmonicsCoefficients
//go:noescape
func GetXRLightEstimateSphericalHarmonicsCoefficients(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRLightEstimate_PrimaryLightDirection
//go:noescape
func GetXRLightEstimatePrimaryLightDirection(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRLightEstimate_PrimaryLightIntensity
//go:noescape
func GetXRLightEstimatePrimaryLightIntensity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRLightProbe_ProbeSpace
//go:noescape
func GetXRLightProbeProbeSpace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRCPUDepthInformation_Data
//go:noescape
func GetXRCPUDepthInformationData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRCPUDepthInformation_GetDepthInMeters
//go:noescape
func HasFuncXRCPUDepthInformationGetDepthInMeters(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRCPUDepthInformation_GetDepthInMeters
//go:noescape
func FuncXRCPUDepthInformationGetDepthInMeters(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRCPUDepthInformation_GetDepthInMeters
//go:noescape
func CallXRCPUDepthInformationGetDepthInMeters(
	this js.Ref, retPtr unsafe.Pointer,
	x float32,
	y float32)

//go:wasmimport plat/js/web try_XRCPUDepthInformation_GetDepthInMeters
//go:noescape
func TryXRCPUDepthInformationGetDepthInMeters(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float32,
	y float32) (ok js.Ref)

//go:wasmimport plat/js/web get_XRJointPose_Radius
//go:noescape
func GetXRJointPoseRadius(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_XRHandJoint
//go:noescape
func ConstOfXRHandJoint(str js.Ref) uint32

//go:wasmimport plat/js/web get_XRJointSpace_JointName
//go:noescape
func GetXRJointSpaceJointName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRHitTestResult_GetPose
//go:noescape
func HasFuncXRHitTestResultGetPose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRHitTestResult_GetPose
//go:noescape
func FuncXRHitTestResultGetPose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRHitTestResult_GetPose
//go:noescape
func CallXRHitTestResultGetPose(
	this js.Ref, retPtr unsafe.Pointer,
	baseSpace js.Ref)

//go:wasmimport plat/js/web try_XRHitTestResult_GetPose
//go:noescape
func TryXRHitTestResultGetPose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	baseSpace js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRHitTestResult_CreateAnchor
//go:noescape
func HasFuncXRHitTestResultCreateAnchor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRHitTestResult_CreateAnchor
//go:noescape
func FuncXRHitTestResultCreateAnchor(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRHitTestResult_CreateAnchor
//go:noescape
func CallXRHitTestResultCreateAnchor(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XRHitTestResult_CreateAnchor
//go:noescape
func TryXRHitTestResultCreateAnchor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
