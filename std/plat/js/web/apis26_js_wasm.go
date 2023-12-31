// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

const (
	_ USBRecipient = iota

	USBRecipient_DEVICE
	USBRecipient_INTERFACE
	USBRecipient_ENDPOINT
	USBRecipient_OTHER
)

func (USBRecipient) FromRef(str js.Ref) USBRecipient {
	return USBRecipient(bindings.ConstOfUSBRecipient(str))
}

func (x USBRecipient) String() (string, bool) {
	switch x {
	case USBRecipient_DEVICE:
		return "device", true
	case USBRecipient_INTERFACE:
		return "interface", true
	case USBRecipient_ENDPOINT:
		return "endpoint", true
	case USBRecipient_OTHER:
		return "other", true
	default:
		return "", false
	}
}

type USBControlTransferParameters struct {
	// RequestType is "USBControlTransferParameters.requestType"
	//
	// Required
	RequestType USBRequestType
	// Recipient is "USBControlTransferParameters.recipient"
	//
	// Required
	Recipient USBRecipient
	// Request is "USBControlTransferParameters.request"
	//
	// Required
	Request uint8
	// Value is "USBControlTransferParameters.value"
	//
	// Required
	Value uint16
	// Index is "USBControlTransferParameters.index"
	//
	// Required
	Index uint16

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a USBControlTransferParameters with all fields set.
func (p USBControlTransferParameters) FromRef(ref js.Ref) USBControlTransferParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new USBControlTransferParameters in the application heap.
func (p USBControlTransferParameters) New() js.Ref {
	return bindings.USBControlTransferParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *USBControlTransferParameters) UpdateFrom(ref js.Ref) {
	bindings.USBControlTransferParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *USBControlTransferParameters) Update(ref js.Ref) {
	bindings.USBControlTransferParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *USBControlTransferParameters) FreeMembers(recursive bool) {
}

func NewUSBOutTransferResult(status USBTransferStatus, bytesWritten uint32) (ret USBOutTransferResult) {
	ret.ref = bindings.NewUSBOutTransferResultByUSBOutTransferResult(
		uint32(status),
		uint32(bytesWritten))
	return
}

func NewUSBOutTransferResultByUSBOutTransferResult1(status USBTransferStatus) (ret USBOutTransferResult) {
	ret.ref = bindings.NewUSBOutTransferResultByUSBOutTransferResult1(
		uint32(status))
	return
}

type USBOutTransferResult struct {
	ref js.Ref
}

func (this USBOutTransferResult) Once() USBOutTransferResult {
	this.ref.Once()
	return this
}

func (this USBOutTransferResult) Ref() js.Ref {
	return this.ref
}

func (this USBOutTransferResult) FromRef(ref js.Ref) USBOutTransferResult {
	this.ref = ref
	return this
}

func (this USBOutTransferResult) Free() {
	this.ref.Free()
}

// BytesWritten returns the value of property "USBOutTransferResult.bytesWritten".
//
// It returns ok=false if there is no such property.
func (this USBOutTransferResult) BytesWritten() (ret uint32, ok bool) {
	ok = js.True == bindings.GetUSBOutTransferResultBytesWritten(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Status returns the value of property "USBOutTransferResult.status".
//
// It returns ok=false if there is no such property.
func (this USBOutTransferResult) Status() (ret USBTransferStatus, ok bool) {
	ok = js.True == bindings.GetUSBOutTransferResultStatus(
		this.ref, js.Pointer(&ret),
	)
	return
}

type USBDirection uint32

const (
	_ USBDirection = iota

	USBDirection_IN
	USBDirection_OUT
)

func (USBDirection) FromRef(str js.Ref) USBDirection {
	return USBDirection(bindings.ConstOfUSBDirection(str))
}

func (x USBDirection) String() (string, bool) {
	switch x {
	case USBDirection_IN:
		return "in", true
	case USBDirection_OUT:
		return "out", true
	default:
		return "", false
	}
}

func NewUSBIsochronousInTransferPacket(status USBTransferStatus, data js.DataView) (ret USBIsochronousInTransferPacket) {
	ret.ref = bindings.NewUSBIsochronousInTransferPacketByUSBIsochronousInTransferPacket(
		uint32(status),
		data.Ref())
	return
}

func NewUSBIsochronousInTransferPacketByUSBIsochronousInTransferPacket1(status USBTransferStatus) (ret USBIsochronousInTransferPacket) {
	ret.ref = bindings.NewUSBIsochronousInTransferPacketByUSBIsochronousInTransferPacket1(
		uint32(status))
	return
}

type USBIsochronousInTransferPacket struct {
	ref js.Ref
}

func (this USBIsochronousInTransferPacket) Once() USBIsochronousInTransferPacket {
	this.ref.Once()
	return this
}

func (this USBIsochronousInTransferPacket) Ref() js.Ref {
	return this.ref
}

func (this USBIsochronousInTransferPacket) FromRef(ref js.Ref) USBIsochronousInTransferPacket {
	this.ref = ref
	return this
}

func (this USBIsochronousInTransferPacket) Free() {
	this.ref.Free()
}

// Data returns the value of property "USBIsochronousInTransferPacket.data".
//
// It returns ok=false if there is no such property.
func (this USBIsochronousInTransferPacket) Data() (ret js.DataView, ok bool) {
	ok = js.True == bindings.GetUSBIsochronousInTransferPacketData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Status returns the value of property "USBIsochronousInTransferPacket.status".
//
// It returns ok=false if there is no such property.
func (this USBIsochronousInTransferPacket) Status() (ret USBTransferStatus, ok bool) {
	ok = js.True == bindings.GetUSBIsochronousInTransferPacketStatus(
		this.ref, js.Pointer(&ret),
	)
	return
}

func NewUSBIsochronousInTransferResult(packets js.Array[USBIsochronousInTransferPacket], data js.DataView) (ret USBIsochronousInTransferResult) {
	ret.ref = bindings.NewUSBIsochronousInTransferResultByUSBIsochronousInTransferResult(
		packets.Ref(),
		data.Ref())
	return
}

func NewUSBIsochronousInTransferResultByUSBIsochronousInTransferResult1(packets js.Array[USBIsochronousInTransferPacket]) (ret USBIsochronousInTransferResult) {
	ret.ref = bindings.NewUSBIsochronousInTransferResultByUSBIsochronousInTransferResult1(
		packets.Ref())
	return
}

type USBIsochronousInTransferResult struct {
	ref js.Ref
}

func (this USBIsochronousInTransferResult) Once() USBIsochronousInTransferResult {
	this.ref.Once()
	return this
}

func (this USBIsochronousInTransferResult) Ref() js.Ref {
	return this.ref
}

func (this USBIsochronousInTransferResult) FromRef(ref js.Ref) USBIsochronousInTransferResult {
	this.ref = ref
	return this
}

func (this USBIsochronousInTransferResult) Free() {
	this.ref.Free()
}

// Data returns the value of property "USBIsochronousInTransferResult.data".
//
// It returns ok=false if there is no such property.
func (this USBIsochronousInTransferResult) Data() (ret js.DataView, ok bool) {
	ok = js.True == bindings.GetUSBIsochronousInTransferResultData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Packets returns the value of property "USBIsochronousInTransferResult.packets".
//
// It returns ok=false if there is no such property.
func (this USBIsochronousInTransferResult) Packets() (ret js.FrozenArray[USBIsochronousInTransferPacket], ok bool) {
	ok = js.True == bindings.GetUSBIsochronousInTransferResultPackets(
		this.ref, js.Pointer(&ret),
	)
	return
}

func NewUSBIsochronousOutTransferPacket(status USBTransferStatus, bytesWritten uint32) (ret USBIsochronousOutTransferPacket) {
	ret.ref = bindings.NewUSBIsochronousOutTransferPacketByUSBIsochronousOutTransferPacket(
		uint32(status),
		uint32(bytesWritten))
	return
}

func NewUSBIsochronousOutTransferPacketByUSBIsochronousOutTransferPacket1(status USBTransferStatus) (ret USBIsochronousOutTransferPacket) {
	ret.ref = bindings.NewUSBIsochronousOutTransferPacketByUSBIsochronousOutTransferPacket1(
		uint32(status))
	return
}

type USBIsochronousOutTransferPacket struct {
	ref js.Ref
}

func (this USBIsochronousOutTransferPacket) Once() USBIsochronousOutTransferPacket {
	this.ref.Once()
	return this
}

func (this USBIsochronousOutTransferPacket) Ref() js.Ref {
	return this.ref
}

func (this USBIsochronousOutTransferPacket) FromRef(ref js.Ref) USBIsochronousOutTransferPacket {
	this.ref = ref
	return this
}

func (this USBIsochronousOutTransferPacket) Free() {
	this.ref.Free()
}

// BytesWritten returns the value of property "USBIsochronousOutTransferPacket.bytesWritten".
//
// It returns ok=false if there is no such property.
func (this USBIsochronousOutTransferPacket) BytesWritten() (ret uint32, ok bool) {
	ok = js.True == bindings.GetUSBIsochronousOutTransferPacketBytesWritten(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Status returns the value of property "USBIsochronousOutTransferPacket.status".
//
// It returns ok=false if there is no such property.
func (this USBIsochronousOutTransferPacket) Status() (ret USBTransferStatus, ok bool) {
	ok = js.True == bindings.GetUSBIsochronousOutTransferPacketStatus(
		this.ref, js.Pointer(&ret),
	)
	return
}

func NewUSBIsochronousOutTransferResult(packets js.Array[USBIsochronousOutTransferPacket]) (ret USBIsochronousOutTransferResult) {
	ret.ref = bindings.NewUSBIsochronousOutTransferResultByUSBIsochronousOutTransferResult(
		packets.Ref())
	return
}

type USBIsochronousOutTransferResult struct {
	ref js.Ref
}

func (this USBIsochronousOutTransferResult) Once() USBIsochronousOutTransferResult {
	this.ref.Once()
	return this
}

func (this USBIsochronousOutTransferResult) Ref() js.Ref {
	return this.ref
}

func (this USBIsochronousOutTransferResult) FromRef(ref js.Ref) USBIsochronousOutTransferResult {
	this.ref = ref
	return this
}

func (this USBIsochronousOutTransferResult) Free() {
	this.ref.Free()
}

// Packets returns the value of property "USBIsochronousOutTransferResult.packets".
//
// It returns ok=false if there is no such property.
func (this USBIsochronousOutTransferResult) Packets() (ret js.FrozenArray[USBIsochronousOutTransferPacket], ok bool) {
	ok = js.True == bindings.GetUSBIsochronousOutTransferResultPackets(
		this.ref, js.Pointer(&ret),
	)
	return
}

type USBEndpointType uint32

const (
	_ USBEndpointType = iota

	USBEndpointType_BULK
	USBEndpointType_INTERRUPT
	USBEndpointType_ISOCHRONOUS
)

func (USBEndpointType) FromRef(str js.Ref) USBEndpointType {
	return USBEndpointType(bindings.ConstOfUSBEndpointType(str))
}

func (x USBEndpointType) String() (string, bool) {
	switch x {
	case USBEndpointType_BULK:
		return "bulk", true
	case USBEndpointType_INTERRUPT:
		return "interrupt", true
	case USBEndpointType_ISOCHRONOUS:
		return "isochronous", true
	default:
		return "", false
	}
}

func NewUSBEndpoint(alternate USBAlternateInterface, endpointNumber uint8, direction USBDirection) (ret USBEndpoint) {
	ret.ref = bindings.NewUSBEndpointByUSBEndpoint(
		alternate.Ref(),
		uint32(endpointNumber),
		uint32(direction))
	return
}

type USBEndpoint struct {
	ref js.Ref
}

func (this USBEndpoint) Once() USBEndpoint {
	this.ref.Once()
	return this
}

func (this USBEndpoint) Ref() js.Ref {
	return this.ref
}

func (this USBEndpoint) FromRef(ref js.Ref) USBEndpoint {
	this.ref = ref
	return this
}

func (this USBEndpoint) Free() {
	this.ref.Free()
}

// EndpointNumber returns the value of property "USBEndpoint.endpointNumber".
//
// It returns ok=false if there is no such property.
func (this USBEndpoint) EndpointNumber() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBEndpointEndpointNumber(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Direction returns the value of property "USBEndpoint.direction".
//
// It returns ok=false if there is no such property.
func (this USBEndpoint) Direction() (ret USBDirection, ok bool) {
	ok = js.True == bindings.GetUSBEndpointDirection(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "USBEndpoint.type".
//
// It returns ok=false if there is no such property.
func (this USBEndpoint) Type() (ret USBEndpointType, ok bool) {
	ok = js.True == bindings.GetUSBEndpointType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PacketSize returns the value of property "USBEndpoint.packetSize".
//
// It returns ok=false if there is no such property.
func (this USBEndpoint) PacketSize() (ret uint32, ok bool) {
	ok = js.True == bindings.GetUSBEndpointPacketSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

func NewUSBAlternateInterface(deviceInterface USBInterface, alternateSetting uint8) (ret USBAlternateInterface) {
	ret.ref = bindings.NewUSBAlternateInterfaceByUSBAlternateInterface(
		deviceInterface.Ref(),
		uint32(alternateSetting))
	return
}

type USBAlternateInterface struct {
	ref js.Ref
}

func (this USBAlternateInterface) Once() USBAlternateInterface {
	this.ref.Once()
	return this
}

func (this USBAlternateInterface) Ref() js.Ref {
	return this.ref
}

func (this USBAlternateInterface) FromRef(ref js.Ref) USBAlternateInterface {
	this.ref = ref
	return this
}

func (this USBAlternateInterface) Free() {
	this.ref.Free()
}

// AlternateSetting returns the value of property "USBAlternateInterface.alternateSetting".
//
// It returns ok=false if there is no such property.
func (this USBAlternateInterface) AlternateSetting() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBAlternateInterfaceAlternateSetting(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InterfaceClass returns the value of property "USBAlternateInterface.interfaceClass".
//
// It returns ok=false if there is no such property.
func (this USBAlternateInterface) InterfaceClass() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBAlternateInterfaceInterfaceClass(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InterfaceSubclass returns the value of property "USBAlternateInterface.interfaceSubclass".
//
// It returns ok=false if there is no such property.
func (this USBAlternateInterface) InterfaceSubclass() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBAlternateInterfaceInterfaceSubclass(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InterfaceProtocol returns the value of property "USBAlternateInterface.interfaceProtocol".
//
// It returns ok=false if there is no such property.
func (this USBAlternateInterface) InterfaceProtocol() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBAlternateInterfaceInterfaceProtocol(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InterfaceName returns the value of property "USBAlternateInterface.interfaceName".
//
// It returns ok=false if there is no such property.
func (this USBAlternateInterface) InterfaceName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetUSBAlternateInterfaceInterfaceName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Endpoints returns the value of property "USBAlternateInterface.endpoints".
//
// It returns ok=false if there is no such property.
func (this USBAlternateInterface) Endpoints() (ret js.FrozenArray[USBEndpoint], ok bool) {
	ok = js.True == bindings.GetUSBAlternateInterfaceEndpoints(
		this.ref, js.Pointer(&ret),
	)
	return
}

func NewUSBInterface(configuration USBConfiguration, interfaceNumber uint8) (ret USBInterface) {
	ret.ref = bindings.NewUSBInterfaceByUSBInterface(
		configuration.Ref(),
		uint32(interfaceNumber))
	return
}

type USBInterface struct {
	ref js.Ref
}

func (this USBInterface) Once() USBInterface {
	this.ref.Once()
	return this
}

func (this USBInterface) Ref() js.Ref {
	return this.ref
}

func (this USBInterface) FromRef(ref js.Ref) USBInterface {
	this.ref = ref
	return this
}

func (this USBInterface) Free() {
	this.ref.Free()
}

// InterfaceNumber returns the value of property "USBInterface.interfaceNumber".
//
// It returns ok=false if there is no such property.
func (this USBInterface) InterfaceNumber() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBInterfaceInterfaceNumber(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Alternate returns the value of property "USBInterface.alternate".
//
// It returns ok=false if there is no such property.
func (this USBInterface) Alternate() (ret USBAlternateInterface, ok bool) {
	ok = js.True == bindings.GetUSBInterfaceAlternate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Alternates returns the value of property "USBInterface.alternates".
//
// It returns ok=false if there is no such property.
func (this USBInterface) Alternates() (ret js.FrozenArray[USBAlternateInterface], ok bool) {
	ok = js.True == bindings.GetUSBInterfaceAlternates(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Claimed returns the value of property "USBInterface.claimed".
//
// It returns ok=false if there is no such property.
func (this USBInterface) Claimed() (ret bool, ok bool) {
	ok = js.True == bindings.GetUSBInterfaceClaimed(
		this.ref, js.Pointer(&ret),
	)
	return
}

func NewUSBConfiguration(device USBDevice, configurationValue uint8) (ret USBConfiguration) {
	ret.ref = bindings.NewUSBConfigurationByUSBConfiguration(
		device.Ref(),
		uint32(configurationValue))
	return
}

type USBConfiguration struct {
	ref js.Ref
}

func (this USBConfiguration) Once() USBConfiguration {
	this.ref.Once()
	return this
}

func (this USBConfiguration) Ref() js.Ref {
	return this.ref
}

func (this USBConfiguration) FromRef(ref js.Ref) USBConfiguration {
	this.ref = ref
	return this
}

func (this USBConfiguration) Free() {
	this.ref.Free()
}

// ConfigurationValue returns the value of property "USBConfiguration.configurationValue".
//
// It returns ok=false if there is no such property.
func (this USBConfiguration) ConfigurationValue() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBConfigurationConfigurationValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ConfigurationName returns the value of property "USBConfiguration.configurationName".
//
// It returns ok=false if there is no such property.
func (this USBConfiguration) ConfigurationName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetUSBConfigurationConfigurationName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Interfaces returns the value of property "USBConfiguration.interfaces".
//
// It returns ok=false if there is no such property.
func (this USBConfiguration) Interfaces() (ret js.FrozenArray[USBInterface], ok bool) {
	ok = js.True == bindings.GetUSBConfigurationInterfaces(
		this.ref, js.Pointer(&ret),
	)
	return
}

type USBDevice struct {
	ref js.Ref
}

func (this USBDevice) Once() USBDevice {
	this.ref.Once()
	return this
}

func (this USBDevice) Ref() js.Ref {
	return this.ref
}

func (this USBDevice) FromRef(ref js.Ref) USBDevice {
	this.ref = ref
	return this
}

func (this USBDevice) Free() {
	this.ref.Free()
}

// UsbVersionMajor returns the value of property "USBDevice.usbVersionMajor".
//
// It returns ok=false if there is no such property.
func (this USBDevice) UsbVersionMajor() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBDeviceUsbVersionMajor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UsbVersionMinor returns the value of property "USBDevice.usbVersionMinor".
//
// It returns ok=false if there is no such property.
func (this USBDevice) UsbVersionMinor() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBDeviceUsbVersionMinor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UsbVersionSubminor returns the value of property "USBDevice.usbVersionSubminor".
//
// It returns ok=false if there is no such property.
func (this USBDevice) UsbVersionSubminor() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBDeviceUsbVersionSubminor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DeviceClass returns the value of property "USBDevice.deviceClass".
//
// It returns ok=false if there is no such property.
func (this USBDevice) DeviceClass() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBDeviceDeviceClass(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DeviceSubclass returns the value of property "USBDevice.deviceSubclass".
//
// It returns ok=false if there is no such property.
func (this USBDevice) DeviceSubclass() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBDeviceDeviceSubclass(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DeviceProtocol returns the value of property "USBDevice.deviceProtocol".
//
// It returns ok=false if there is no such property.
func (this USBDevice) DeviceProtocol() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBDeviceDeviceProtocol(
		this.ref, js.Pointer(&ret),
	)
	return
}

// VendorId returns the value of property "USBDevice.vendorId".
//
// It returns ok=false if there is no such property.
func (this USBDevice) VendorId() (ret uint16, ok bool) {
	ok = js.True == bindings.GetUSBDeviceVendorId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ProductId returns the value of property "USBDevice.productId".
//
// It returns ok=false if there is no such property.
func (this USBDevice) ProductId() (ret uint16, ok bool) {
	ok = js.True == bindings.GetUSBDeviceProductId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DeviceVersionMajor returns the value of property "USBDevice.deviceVersionMajor".
//
// It returns ok=false if there is no such property.
func (this USBDevice) DeviceVersionMajor() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBDeviceDeviceVersionMajor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DeviceVersionMinor returns the value of property "USBDevice.deviceVersionMinor".
//
// It returns ok=false if there is no such property.
func (this USBDevice) DeviceVersionMinor() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBDeviceDeviceVersionMinor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DeviceVersionSubminor returns the value of property "USBDevice.deviceVersionSubminor".
//
// It returns ok=false if there is no such property.
func (this USBDevice) DeviceVersionSubminor() (ret uint8, ok bool) {
	ok = js.True == bindings.GetUSBDeviceDeviceVersionSubminor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ManufacturerName returns the value of property "USBDevice.manufacturerName".
//
// It returns ok=false if there is no such property.
func (this USBDevice) ManufacturerName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetUSBDeviceManufacturerName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ProductName returns the value of property "USBDevice.productName".
//
// It returns ok=false if there is no such property.
func (this USBDevice) ProductName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetUSBDeviceProductName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SerialNumber returns the value of property "USBDevice.serialNumber".
//
// It returns ok=false if there is no such property.
func (this USBDevice) SerialNumber() (ret js.String, ok bool) {
	ok = js.True == bindings.GetUSBDeviceSerialNumber(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Configuration returns the value of property "USBDevice.configuration".
//
// It returns ok=false if there is no such property.
func (this USBDevice) Configuration() (ret USBConfiguration, ok bool) {
	ok = js.True == bindings.GetUSBDeviceConfiguration(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Configurations returns the value of property "USBDevice.configurations".
//
// It returns ok=false if there is no such property.
func (this USBDevice) Configurations() (ret js.FrozenArray[USBConfiguration], ok bool) {
	ok = js.True == bindings.GetUSBDeviceConfigurations(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Opened returns the value of property "USBDevice.opened".
//
// It returns ok=false if there is no such property.
func (this USBDevice) Opened() (ret bool, ok bool) {
	ok = js.True == bindings.GetUSBDeviceOpened(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncOpen returns true if the method "USBDevice.open" exists.
func (this USBDevice) HasFuncOpen() bool {
	return js.True == bindings.HasFuncUSBDeviceOpen(
		this.ref,
	)
}

// FuncOpen returns the method "USBDevice.open".
func (this USBDevice) FuncOpen() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncUSBDeviceOpen(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open calls the method "USBDevice.open".
func (this USBDevice) Open() (ret js.Promise[js.Void]) {
	bindings.CallUSBDeviceOpen(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryOpen calls the method "USBDevice.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TryOpen() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceOpen(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClose returns true if the method "USBDevice.close" exists.
func (this USBDevice) HasFuncClose() bool {
	return js.True == bindings.HasFuncUSBDeviceClose(
		this.ref,
	)
}

// FuncClose returns the method "USBDevice.close".
func (this USBDevice) FuncClose() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncUSBDeviceClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "USBDevice.close".
func (this USBDevice) Close() (ret js.Promise[js.Void]) {
	bindings.CallUSBDeviceClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "USBDevice.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TryClose() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncForget returns true if the method "USBDevice.forget" exists.
func (this USBDevice) HasFuncForget() bool {
	return js.True == bindings.HasFuncUSBDeviceForget(
		this.ref,
	)
}

// FuncForget returns the method "USBDevice.forget".
func (this USBDevice) FuncForget() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncUSBDeviceForget(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Forget calls the method "USBDevice.forget".
func (this USBDevice) Forget() (ret js.Promise[js.Void]) {
	bindings.CallUSBDeviceForget(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryForget calls the method "USBDevice.forget"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TryForget() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceForget(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSelectConfiguration returns true if the method "USBDevice.selectConfiguration" exists.
func (this USBDevice) HasFuncSelectConfiguration() bool {
	return js.True == bindings.HasFuncUSBDeviceSelectConfiguration(
		this.ref,
	)
}

// FuncSelectConfiguration returns the method "USBDevice.selectConfiguration".
func (this USBDevice) FuncSelectConfiguration() (fn js.Func[func(configurationValue uint8) js.Promise[js.Void]]) {
	bindings.FuncUSBDeviceSelectConfiguration(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SelectConfiguration calls the method "USBDevice.selectConfiguration".
func (this USBDevice) SelectConfiguration(configurationValue uint8) (ret js.Promise[js.Void]) {
	bindings.CallUSBDeviceSelectConfiguration(
		this.ref, js.Pointer(&ret),
		uint32(configurationValue),
	)

	return
}

// TrySelectConfiguration calls the method "USBDevice.selectConfiguration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TrySelectConfiguration(configurationValue uint8) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceSelectConfiguration(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(configurationValue),
	)

	return
}

// HasFuncClaimInterface returns true if the method "USBDevice.claimInterface" exists.
func (this USBDevice) HasFuncClaimInterface() bool {
	return js.True == bindings.HasFuncUSBDeviceClaimInterface(
		this.ref,
	)
}

// FuncClaimInterface returns the method "USBDevice.claimInterface".
func (this USBDevice) FuncClaimInterface() (fn js.Func[func(interfaceNumber uint8) js.Promise[js.Void]]) {
	bindings.FuncUSBDeviceClaimInterface(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClaimInterface calls the method "USBDevice.claimInterface".
func (this USBDevice) ClaimInterface(interfaceNumber uint8) (ret js.Promise[js.Void]) {
	bindings.CallUSBDeviceClaimInterface(
		this.ref, js.Pointer(&ret),
		uint32(interfaceNumber),
	)

	return
}

// TryClaimInterface calls the method "USBDevice.claimInterface"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TryClaimInterface(interfaceNumber uint8) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceClaimInterface(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(interfaceNumber),
	)

	return
}

// HasFuncReleaseInterface returns true if the method "USBDevice.releaseInterface" exists.
func (this USBDevice) HasFuncReleaseInterface() bool {
	return js.True == bindings.HasFuncUSBDeviceReleaseInterface(
		this.ref,
	)
}

// FuncReleaseInterface returns the method "USBDevice.releaseInterface".
func (this USBDevice) FuncReleaseInterface() (fn js.Func[func(interfaceNumber uint8) js.Promise[js.Void]]) {
	bindings.FuncUSBDeviceReleaseInterface(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReleaseInterface calls the method "USBDevice.releaseInterface".
func (this USBDevice) ReleaseInterface(interfaceNumber uint8) (ret js.Promise[js.Void]) {
	bindings.CallUSBDeviceReleaseInterface(
		this.ref, js.Pointer(&ret),
		uint32(interfaceNumber),
	)

	return
}

// TryReleaseInterface calls the method "USBDevice.releaseInterface"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TryReleaseInterface(interfaceNumber uint8) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceReleaseInterface(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(interfaceNumber),
	)

	return
}

// HasFuncSelectAlternateInterface returns true if the method "USBDevice.selectAlternateInterface" exists.
func (this USBDevice) HasFuncSelectAlternateInterface() bool {
	return js.True == bindings.HasFuncUSBDeviceSelectAlternateInterface(
		this.ref,
	)
}

// FuncSelectAlternateInterface returns the method "USBDevice.selectAlternateInterface".
func (this USBDevice) FuncSelectAlternateInterface() (fn js.Func[func(interfaceNumber uint8, alternateSetting uint8) js.Promise[js.Void]]) {
	bindings.FuncUSBDeviceSelectAlternateInterface(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SelectAlternateInterface calls the method "USBDevice.selectAlternateInterface".
func (this USBDevice) SelectAlternateInterface(interfaceNumber uint8, alternateSetting uint8) (ret js.Promise[js.Void]) {
	bindings.CallUSBDeviceSelectAlternateInterface(
		this.ref, js.Pointer(&ret),
		uint32(interfaceNumber),
		uint32(alternateSetting),
	)

	return
}

// TrySelectAlternateInterface calls the method "USBDevice.selectAlternateInterface"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TrySelectAlternateInterface(interfaceNumber uint8, alternateSetting uint8) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceSelectAlternateInterface(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(interfaceNumber),
		uint32(alternateSetting),
	)

	return
}

// HasFuncControlTransferIn returns true if the method "USBDevice.controlTransferIn" exists.
func (this USBDevice) HasFuncControlTransferIn() bool {
	return js.True == bindings.HasFuncUSBDeviceControlTransferIn(
		this.ref,
	)
}

// FuncControlTransferIn returns the method "USBDevice.controlTransferIn".
func (this USBDevice) FuncControlTransferIn() (fn js.Func[func(setup USBControlTransferParameters, length uint16) js.Promise[USBInTransferResult]]) {
	bindings.FuncUSBDeviceControlTransferIn(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ControlTransferIn calls the method "USBDevice.controlTransferIn".
func (this USBDevice) ControlTransferIn(setup USBControlTransferParameters, length uint16) (ret js.Promise[USBInTransferResult]) {
	bindings.CallUSBDeviceControlTransferIn(
		this.ref, js.Pointer(&ret),
		js.Pointer(&setup),
		uint32(length),
	)

	return
}

// TryControlTransferIn calls the method "USBDevice.controlTransferIn"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TryControlTransferIn(setup USBControlTransferParameters, length uint16) (ret js.Promise[USBInTransferResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceControlTransferIn(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&setup),
		uint32(length),
	)

	return
}

// HasFuncControlTransferOut returns true if the method "USBDevice.controlTransferOut" exists.
func (this USBDevice) HasFuncControlTransferOut() bool {
	return js.True == bindings.HasFuncUSBDeviceControlTransferOut(
		this.ref,
	)
}

// FuncControlTransferOut returns the method "USBDevice.controlTransferOut".
func (this USBDevice) FuncControlTransferOut() (fn js.Func[func(setup USBControlTransferParameters, data BufferSource) js.Promise[USBOutTransferResult]]) {
	bindings.FuncUSBDeviceControlTransferOut(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ControlTransferOut calls the method "USBDevice.controlTransferOut".
func (this USBDevice) ControlTransferOut(setup USBControlTransferParameters, data BufferSource) (ret js.Promise[USBOutTransferResult]) {
	bindings.CallUSBDeviceControlTransferOut(
		this.ref, js.Pointer(&ret),
		js.Pointer(&setup),
		data.Ref(),
	)

	return
}

// TryControlTransferOut calls the method "USBDevice.controlTransferOut"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TryControlTransferOut(setup USBControlTransferParameters, data BufferSource) (ret js.Promise[USBOutTransferResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceControlTransferOut(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&setup),
		data.Ref(),
	)

	return
}

// HasFuncControlTransferOut1 returns true if the method "USBDevice.controlTransferOut" exists.
func (this USBDevice) HasFuncControlTransferOut1() bool {
	return js.True == bindings.HasFuncUSBDeviceControlTransferOut1(
		this.ref,
	)
}

// FuncControlTransferOut1 returns the method "USBDevice.controlTransferOut".
func (this USBDevice) FuncControlTransferOut1() (fn js.Func[func(setup USBControlTransferParameters) js.Promise[USBOutTransferResult]]) {
	bindings.FuncUSBDeviceControlTransferOut1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ControlTransferOut1 calls the method "USBDevice.controlTransferOut".
func (this USBDevice) ControlTransferOut1(setup USBControlTransferParameters) (ret js.Promise[USBOutTransferResult]) {
	bindings.CallUSBDeviceControlTransferOut1(
		this.ref, js.Pointer(&ret),
		js.Pointer(&setup),
	)

	return
}

// TryControlTransferOut1 calls the method "USBDevice.controlTransferOut"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TryControlTransferOut1(setup USBControlTransferParameters) (ret js.Promise[USBOutTransferResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceControlTransferOut1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&setup),
	)

	return
}

// HasFuncClearHalt returns true if the method "USBDevice.clearHalt" exists.
func (this USBDevice) HasFuncClearHalt() bool {
	return js.True == bindings.HasFuncUSBDeviceClearHalt(
		this.ref,
	)
}

// FuncClearHalt returns the method "USBDevice.clearHalt".
func (this USBDevice) FuncClearHalt() (fn js.Func[func(direction USBDirection, endpointNumber uint8) js.Promise[js.Void]]) {
	bindings.FuncUSBDeviceClearHalt(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearHalt calls the method "USBDevice.clearHalt".
func (this USBDevice) ClearHalt(direction USBDirection, endpointNumber uint8) (ret js.Promise[js.Void]) {
	bindings.CallUSBDeviceClearHalt(
		this.ref, js.Pointer(&ret),
		uint32(direction),
		uint32(endpointNumber),
	)

	return
}

// TryClearHalt calls the method "USBDevice.clearHalt"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TryClearHalt(direction USBDirection, endpointNumber uint8) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceClearHalt(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(direction),
		uint32(endpointNumber),
	)

	return
}

// HasFuncTransferIn returns true if the method "USBDevice.transferIn" exists.
func (this USBDevice) HasFuncTransferIn() bool {
	return js.True == bindings.HasFuncUSBDeviceTransferIn(
		this.ref,
	)
}

// FuncTransferIn returns the method "USBDevice.transferIn".
func (this USBDevice) FuncTransferIn() (fn js.Func[func(endpointNumber uint8, length uint32) js.Promise[USBInTransferResult]]) {
	bindings.FuncUSBDeviceTransferIn(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TransferIn calls the method "USBDevice.transferIn".
func (this USBDevice) TransferIn(endpointNumber uint8, length uint32) (ret js.Promise[USBInTransferResult]) {
	bindings.CallUSBDeviceTransferIn(
		this.ref, js.Pointer(&ret),
		uint32(endpointNumber),
		uint32(length),
	)

	return
}

// TryTransferIn calls the method "USBDevice.transferIn"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TryTransferIn(endpointNumber uint8, length uint32) (ret js.Promise[USBInTransferResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceTransferIn(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(endpointNumber),
		uint32(length),
	)

	return
}

// HasFuncTransferOut returns true if the method "USBDevice.transferOut" exists.
func (this USBDevice) HasFuncTransferOut() bool {
	return js.True == bindings.HasFuncUSBDeviceTransferOut(
		this.ref,
	)
}

// FuncTransferOut returns the method "USBDevice.transferOut".
func (this USBDevice) FuncTransferOut() (fn js.Func[func(endpointNumber uint8, data BufferSource) js.Promise[USBOutTransferResult]]) {
	bindings.FuncUSBDeviceTransferOut(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TransferOut calls the method "USBDevice.transferOut".
func (this USBDevice) TransferOut(endpointNumber uint8, data BufferSource) (ret js.Promise[USBOutTransferResult]) {
	bindings.CallUSBDeviceTransferOut(
		this.ref, js.Pointer(&ret),
		uint32(endpointNumber),
		data.Ref(),
	)

	return
}

// TryTransferOut calls the method "USBDevice.transferOut"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TryTransferOut(endpointNumber uint8, data BufferSource) (ret js.Promise[USBOutTransferResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceTransferOut(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(endpointNumber),
		data.Ref(),
	)

	return
}

// HasFuncIsochronousTransferIn returns true if the method "USBDevice.isochronousTransferIn" exists.
func (this USBDevice) HasFuncIsochronousTransferIn() bool {
	return js.True == bindings.HasFuncUSBDeviceIsochronousTransferIn(
		this.ref,
	)
}

// FuncIsochronousTransferIn returns the method "USBDevice.isochronousTransferIn".
func (this USBDevice) FuncIsochronousTransferIn() (fn js.Func[func(endpointNumber uint8, packetLengths js.Array[uint32]) js.Promise[USBIsochronousInTransferResult]]) {
	bindings.FuncUSBDeviceIsochronousTransferIn(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsochronousTransferIn calls the method "USBDevice.isochronousTransferIn".
func (this USBDevice) IsochronousTransferIn(endpointNumber uint8, packetLengths js.Array[uint32]) (ret js.Promise[USBIsochronousInTransferResult]) {
	bindings.CallUSBDeviceIsochronousTransferIn(
		this.ref, js.Pointer(&ret),
		uint32(endpointNumber),
		packetLengths.Ref(),
	)

	return
}

// TryIsochronousTransferIn calls the method "USBDevice.isochronousTransferIn"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TryIsochronousTransferIn(endpointNumber uint8, packetLengths js.Array[uint32]) (ret js.Promise[USBIsochronousInTransferResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceIsochronousTransferIn(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(endpointNumber),
		packetLengths.Ref(),
	)

	return
}

// HasFuncIsochronousTransferOut returns true if the method "USBDevice.isochronousTransferOut" exists.
func (this USBDevice) HasFuncIsochronousTransferOut() bool {
	return js.True == bindings.HasFuncUSBDeviceIsochronousTransferOut(
		this.ref,
	)
}

// FuncIsochronousTransferOut returns the method "USBDevice.isochronousTransferOut".
func (this USBDevice) FuncIsochronousTransferOut() (fn js.Func[func(endpointNumber uint8, data BufferSource, packetLengths js.Array[uint32]) js.Promise[USBIsochronousOutTransferResult]]) {
	bindings.FuncUSBDeviceIsochronousTransferOut(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsochronousTransferOut calls the method "USBDevice.isochronousTransferOut".
func (this USBDevice) IsochronousTransferOut(endpointNumber uint8, data BufferSource, packetLengths js.Array[uint32]) (ret js.Promise[USBIsochronousOutTransferResult]) {
	bindings.CallUSBDeviceIsochronousTransferOut(
		this.ref, js.Pointer(&ret),
		uint32(endpointNumber),
		data.Ref(),
		packetLengths.Ref(),
	)

	return
}

// TryIsochronousTransferOut calls the method "USBDevice.isochronousTransferOut"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TryIsochronousTransferOut(endpointNumber uint8, data BufferSource, packetLengths js.Array[uint32]) (ret js.Promise[USBIsochronousOutTransferResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceIsochronousTransferOut(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(endpointNumber),
		data.Ref(),
		packetLengths.Ref(),
	)

	return
}

// HasFuncReset returns true if the method "USBDevice.reset" exists.
func (this USBDevice) HasFuncReset() bool {
	return js.True == bindings.HasFuncUSBDeviceReset(
		this.ref,
	)
}

// FuncReset returns the method "USBDevice.reset".
func (this USBDevice) FuncReset() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncUSBDeviceReset(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Reset calls the method "USBDevice.reset".
func (this USBDevice) Reset() (ret js.Promise[js.Void]) {
	bindings.CallUSBDeviceReset(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "USBDevice.reset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USBDevice) TryReset() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBDeviceReset(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type USBDeviceFilter struct {
	// VendorId is "USBDeviceFilter.vendorId"
	//
	// Optional
	//
	// NOTE: FFI_USE_VendorId MUST be set to true to make this field effective.
	VendorId uint16
	// ProductId is "USBDeviceFilter.productId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProductId MUST be set to true to make this field effective.
	ProductId uint16
	// ClassCode is "USBDeviceFilter.classCode"
	//
	// Optional
	//
	// NOTE: FFI_USE_ClassCode MUST be set to true to make this field effective.
	ClassCode uint8
	// SubclassCode is "USBDeviceFilter.subclassCode"
	//
	// Optional
	//
	// NOTE: FFI_USE_SubclassCode MUST be set to true to make this field effective.
	SubclassCode uint8
	// ProtocolCode is "USBDeviceFilter.protocolCode"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProtocolCode MUST be set to true to make this field effective.
	ProtocolCode uint8
	// SerialNumber is "USBDeviceFilter.serialNumber"
	//
	// Optional
	SerialNumber js.String

	FFI_USE_VendorId     bool // for VendorId.
	FFI_USE_ProductId    bool // for ProductId.
	FFI_USE_ClassCode    bool // for ClassCode.
	FFI_USE_SubclassCode bool // for SubclassCode.
	FFI_USE_ProtocolCode bool // for ProtocolCode.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a USBDeviceFilter with all fields set.
func (p USBDeviceFilter) FromRef(ref js.Ref) USBDeviceFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new USBDeviceFilter in the application heap.
func (p USBDeviceFilter) New() js.Ref {
	return bindings.USBDeviceFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *USBDeviceFilter) UpdateFrom(ref js.Ref) {
	bindings.USBDeviceFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *USBDeviceFilter) Update(ref js.Ref) {
	bindings.USBDeviceFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *USBDeviceFilter) FreeMembers(recursive bool) {
	js.Free(
		p.SerialNumber.Ref(),
	)
	p.SerialNumber = p.SerialNumber.FromRef(js.Undefined)
}

type USBDeviceRequestOptions struct {
	// Filters is "USBDeviceRequestOptions.filters"
	//
	// Required
	Filters js.Array[USBDeviceFilter]
	// ExclusionFilters is "USBDeviceRequestOptions.exclusionFilters"
	//
	// Optional, defaults to [].
	ExclusionFilters js.Array[USBDeviceFilter]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a USBDeviceRequestOptions with all fields set.
func (p USBDeviceRequestOptions) FromRef(ref js.Ref) USBDeviceRequestOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new USBDeviceRequestOptions in the application heap.
func (p USBDeviceRequestOptions) New() js.Ref {
	return bindings.USBDeviceRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *USBDeviceRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.USBDeviceRequestOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *USBDeviceRequestOptions) Update(ref js.Ref) {
	bindings.USBDeviceRequestOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *USBDeviceRequestOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Filters.Ref(),
		p.ExclusionFilters.Ref(),
	)
	p.Filters = p.Filters.FromRef(js.Undefined)
	p.ExclusionFilters = p.ExclusionFilters.FromRef(js.Undefined)
}

type USB struct {
	EventTarget
}

func (this USB) Once() USB {
	this.ref.Once()
	return this
}

func (this USB) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this USB) FromRef(ref js.Ref) USB {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this USB) Free() {
	this.ref.Free()
}

// HasFuncGetDevices returns true if the method "USB.getDevices" exists.
func (this USB) HasFuncGetDevices() bool {
	return js.True == bindings.HasFuncUSBGetDevices(
		this.ref,
	)
}

// FuncGetDevices returns the method "USB.getDevices".
func (this USB) FuncGetDevices() (fn js.Func[func() js.Promise[js.Array[USBDevice]]]) {
	bindings.FuncUSBGetDevices(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDevices calls the method "USB.getDevices".
func (this USB) GetDevices() (ret js.Promise[js.Array[USBDevice]]) {
	bindings.CallUSBGetDevices(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetDevices calls the method "USB.getDevices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USB) TryGetDevices() (ret js.Promise[js.Array[USBDevice]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBGetDevices(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestDevice returns true if the method "USB.requestDevice" exists.
func (this USB) HasFuncRequestDevice() bool {
	return js.True == bindings.HasFuncUSBRequestDevice(
		this.ref,
	)
}

// FuncRequestDevice returns the method "USB.requestDevice".
func (this USB) FuncRequestDevice() (fn js.Func[func(options USBDeviceRequestOptions) js.Promise[USBDevice]]) {
	bindings.FuncUSBRequestDevice(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestDevice calls the method "USB.requestDevice".
func (this USB) RequestDevice(options USBDeviceRequestOptions) (ret js.Promise[USBDevice]) {
	bindings.CallUSBRequestDevice(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestDevice calls the method "USB.requestDevice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this USB) TryRequestDevice(options USBDeviceRequestOptions) (ret js.Promise[USBDevice], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUSBRequestDevice(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

type EpubReadingSystem struct {
	ref js.Ref
}

func (this EpubReadingSystem) Once() EpubReadingSystem {
	this.ref.Once()
	return this
}

func (this EpubReadingSystem) Ref() js.Ref {
	return this.ref
}

func (this EpubReadingSystem) FromRef(ref js.Ref) EpubReadingSystem {
	this.ref = ref
	return this
}

func (this EpubReadingSystem) Free() {
	this.ref.Free()
}

// HasFuncHasFeature returns true if the method "EpubReadingSystem.hasFeature" exists.
func (this EpubReadingSystem) HasFuncHasFeature() bool {
	return js.True == bindings.HasFuncEpubReadingSystemHasFeature(
		this.ref,
	)
}

// FuncHasFeature returns the method "EpubReadingSystem.hasFeature".
func (this EpubReadingSystem) FuncHasFeature() (fn js.Func[func(feature js.String, version js.String) bool]) {
	bindings.FuncEpubReadingSystemHasFeature(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HasFeature calls the method "EpubReadingSystem.hasFeature".
func (this EpubReadingSystem) HasFeature(feature js.String, version js.String) (ret bool) {
	bindings.CallEpubReadingSystemHasFeature(
		this.ref, js.Pointer(&ret),
		feature.Ref(),
		version.Ref(),
	)

	return
}

// TryHasFeature calls the method "EpubReadingSystem.hasFeature"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EpubReadingSystem) TryHasFeature(feature js.String, version js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEpubReadingSystemHasFeature(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		feature.Ref(),
		version.Ref(),
	)

	return
}

// HasFuncHasFeature1 returns true if the method "EpubReadingSystem.hasFeature" exists.
func (this EpubReadingSystem) HasFuncHasFeature1() bool {
	return js.True == bindings.HasFuncEpubReadingSystemHasFeature1(
		this.ref,
	)
}

// FuncHasFeature1 returns the method "EpubReadingSystem.hasFeature".
func (this EpubReadingSystem) FuncHasFeature1() (fn js.Func[func(feature js.String) bool]) {
	bindings.FuncEpubReadingSystemHasFeature1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HasFeature1 calls the method "EpubReadingSystem.hasFeature".
func (this EpubReadingSystem) HasFeature1(feature js.String) (ret bool) {
	bindings.CallEpubReadingSystemHasFeature1(
		this.ref, js.Pointer(&ret),
		feature.Ref(),
	)

	return
}

// TryHasFeature1 calls the method "EpubReadingSystem.hasFeature"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EpubReadingSystem) TryHasFeature1(feature js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEpubReadingSystemHasFeature1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		feature.Ref(),
	)

	return
}

type XRSessionMode uint32

const (
	_ XRSessionMode = iota

	XRSessionMode_INLINE
	XRSessionMode_IMMERSIVE_VR
	XRSessionMode_IMMERSIVE_AR
)

func (XRSessionMode) FromRef(str js.Ref) XRSessionMode {
	return XRSessionMode(bindings.ConstOfXRSessionMode(str))
}

func (x XRSessionMode) String() (string, bool) {
	switch x {
	case XRSessionMode_INLINE:
		return "inline", true
	case XRSessionMode_IMMERSIVE_VR:
		return "immersive-vr", true
	case XRSessionMode_IMMERSIVE_AR:
		return "immersive-ar", true
	default:
		return "", false
	}
}

type OneOf_WebGLRenderingContext_WebGL2RenderingContext struct {
	ref js.Ref
}

func (x OneOf_WebGLRenderingContext_WebGL2RenderingContext) Ref() js.Ref {
	return x.ref
}

func (x OneOf_WebGLRenderingContext_WebGL2RenderingContext) Free() {
	x.ref.Free()
}

func (x OneOf_WebGLRenderingContext_WebGL2RenderingContext) FromRef(ref js.Ref) OneOf_WebGLRenderingContext_WebGL2RenderingContext {
	return OneOf_WebGLRenderingContext_WebGL2RenderingContext{
		ref: ref,
	}
}

func (x OneOf_WebGLRenderingContext_WebGL2RenderingContext) WebGLRenderingContext() WebGLRenderingContext {
	return WebGLRenderingContext{}.FromRef(x.ref)
}

func (x OneOf_WebGLRenderingContext_WebGL2RenderingContext) WebGL2RenderingContext() WebGL2RenderingContext {
	return WebGL2RenderingContext{}.FromRef(x.ref)
}

type XRWebGLRenderingContext = OneOf_WebGLRenderingContext_WebGL2RenderingContext

type XRWebGLLayerInit struct {
	// Antialias is "XRWebGLLayerInit.antialias"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Antialias MUST be set to true to make this field effective.
	Antialias bool
	// Depth is "XRWebGLLayerInit.depth"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Depth MUST be set to true to make this field effective.
	Depth bool
	// Stencil is "XRWebGLLayerInit.stencil"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Stencil MUST be set to true to make this field effective.
	Stencil bool
	// Alpha is "XRWebGLLayerInit.alpha"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Alpha MUST be set to true to make this field effective.
	Alpha bool
	// IgnoreDepthValues is "XRWebGLLayerInit.ignoreDepthValues"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IgnoreDepthValues MUST be set to true to make this field effective.
	IgnoreDepthValues bool
	// FramebufferScaleFactor is "XRWebGLLayerInit.framebufferScaleFactor"
	//
	// Optional, defaults to 1.0.
	//
	// NOTE: FFI_USE_FramebufferScaleFactor MUST be set to true to make this field effective.
	FramebufferScaleFactor float64

	FFI_USE_Antialias              bool // for Antialias.
	FFI_USE_Depth                  bool // for Depth.
	FFI_USE_Stencil                bool // for Stencil.
	FFI_USE_Alpha                  bool // for Alpha.
	FFI_USE_IgnoreDepthValues      bool // for IgnoreDepthValues.
	FFI_USE_FramebufferScaleFactor bool // for FramebufferScaleFactor.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRWebGLLayerInit with all fields set.
func (p XRWebGLLayerInit) FromRef(ref js.Ref) XRWebGLLayerInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRWebGLLayerInit in the application heap.
func (p XRWebGLLayerInit) New() js.Ref {
	return bindings.XRWebGLLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRWebGLLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRWebGLLayerInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRWebGLLayerInit) Update(ref js.Ref) {
	bindings.XRWebGLLayerInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRWebGLLayerInit) FreeMembers(recursive bool) {
}

type XRViewport struct {
	ref js.Ref
}

func (this XRViewport) Once() XRViewport {
	this.ref.Once()
	return this
}

func (this XRViewport) Ref() js.Ref {
	return this.ref
}

func (this XRViewport) FromRef(ref js.Ref) XRViewport {
	this.ref = ref
	return this
}

func (this XRViewport) Free() {
	this.ref.Free()
}

// X returns the value of property "XRViewport.x".
//
// It returns ok=false if there is no such property.
func (this XRViewport) X() (ret int32, ok bool) {
	ok = js.True == bindings.GetXRViewportX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "XRViewport.y".
//
// It returns ok=false if there is no such property.
func (this XRViewport) Y() (ret int32, ok bool) {
	ok = js.True == bindings.GetXRViewportY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "XRViewport.width".
//
// It returns ok=false if there is no such property.
func (this XRViewport) Width() (ret int32, ok bool) {
	ok = js.True == bindings.GetXRViewportWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "XRViewport.height".
//
// It returns ok=false if there is no such property.
func (this XRViewport) Height() (ret int32, ok bool) {
	ok = js.True == bindings.GetXRViewportHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XREye uint32

const (
	_ XREye = iota

	XREye_NONE
	XREye_LEFT
	XREye_RIGHT
)

func (XREye) FromRef(str js.Ref) XREye {
	return XREye(bindings.ConstOfXREye(str))
}

func (x XREye) String() (string, bool) {
	switch x {
	case XREye_NONE:
		return "none", true
	case XREye_LEFT:
		return "left", true
	case XREye_RIGHT:
		return "right", true
	default:
		return "", false
	}
}

func NewXRRigidTransform(position DOMPointInit, orientation DOMPointInit) (ret XRRigidTransform) {
	ret.ref = bindings.NewXRRigidTransformByXRRigidTransform(
		js.Pointer(&position),
		js.Pointer(&orientation))
	return
}

func NewXRRigidTransformByXRRigidTransform1(position DOMPointInit) (ret XRRigidTransform) {
	ret.ref = bindings.NewXRRigidTransformByXRRigidTransform1(
		js.Pointer(&position))
	return
}

func NewXRRigidTransformByXRRigidTransform2() (ret XRRigidTransform) {
	ret.ref = bindings.NewXRRigidTransformByXRRigidTransform2()
	return
}

type XRRigidTransform struct {
	ref js.Ref
}

func (this XRRigidTransform) Once() XRRigidTransform {
	this.ref.Once()
	return this
}

func (this XRRigidTransform) Ref() js.Ref {
	return this.ref
}

func (this XRRigidTransform) FromRef(ref js.Ref) XRRigidTransform {
	this.ref = ref
	return this
}

func (this XRRigidTransform) Free() {
	this.ref.Free()
}

// Position returns the value of property "XRRigidTransform.position".
//
// It returns ok=false if there is no such property.
func (this XRRigidTransform) Position() (ret DOMPointReadOnly, ok bool) {
	ok = js.True == bindings.GetXRRigidTransformPosition(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Orientation returns the value of property "XRRigidTransform.orientation".
//
// It returns ok=false if there is no such property.
func (this XRRigidTransform) Orientation() (ret DOMPointReadOnly, ok bool) {
	ok = js.True == bindings.GetXRRigidTransformOrientation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Matrix returns the value of property "XRRigidTransform.matrix".
//
// It returns ok=false if there is no such property.
func (this XRRigidTransform) Matrix() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetXRRigidTransformMatrix(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Inverse returns the value of property "XRRigidTransform.inverse".
//
// It returns ok=false if there is no such property.
func (this XRRigidTransform) Inverse() (ret XRRigidTransform, ok bool) {
	ok = js.True == bindings.GetXRRigidTransformInverse(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRCamera struct {
	ref js.Ref
}

func (this XRCamera) Once() XRCamera {
	this.ref.Once()
	return this
}

func (this XRCamera) Ref() js.Ref {
	return this.ref
}

func (this XRCamera) FromRef(ref js.Ref) XRCamera {
	this.ref = ref
	return this
}

func (this XRCamera) Free() {
	this.ref.Free()
}

// Width returns the value of property "XRCamera.width".
//
// It returns ok=false if there is no such property.
func (this XRCamera) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRCameraWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "XRCamera.height".
//
// It returns ok=false if there is no such property.
func (this XRCamera) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRCameraHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRView struct {
	ref js.Ref
}

func (this XRView) Once() XRView {
	this.ref.Once()
	return this
}

func (this XRView) Ref() js.Ref {
	return this.ref
}

func (this XRView) FromRef(ref js.Ref) XRView {
	this.ref = ref
	return this
}

func (this XRView) Free() {
	this.ref.Free()
}

// Eye returns the value of property "XRView.eye".
//
// It returns ok=false if there is no such property.
func (this XRView) Eye() (ret XREye, ok bool) {
	ok = js.True == bindings.GetXRViewEye(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ProjectionMatrix returns the value of property "XRView.projectionMatrix".
//
// It returns ok=false if there is no such property.
func (this XRView) ProjectionMatrix() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetXRViewProjectionMatrix(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Transform returns the value of property "XRView.transform".
//
// It returns ok=false if there is no such property.
func (this XRView) Transform() (ret XRRigidTransform, ok bool) {
	ok = js.True == bindings.GetXRViewTransform(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RecommendedViewportScale returns the value of property "XRView.recommendedViewportScale".
//
// It returns ok=false if there is no such property.
func (this XRView) RecommendedViewportScale() (ret float64, ok bool) {
	ok = js.True == bindings.GetXRViewRecommendedViewportScale(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsFirstPersonObserver returns the value of property "XRView.isFirstPersonObserver".
//
// It returns ok=false if there is no such property.
func (this XRView) IsFirstPersonObserver() (ret bool, ok bool) {
	ok = js.True == bindings.GetXRViewIsFirstPersonObserver(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Camera returns the value of property "XRView.camera".
//
// It returns ok=false if there is no such property.
func (this XRView) Camera() (ret XRCamera, ok bool) {
	ok = js.True == bindings.GetXRViewCamera(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncRequestViewportScale returns true if the method "XRView.requestViewportScale" exists.
func (this XRView) HasFuncRequestViewportScale() bool {
	return js.True == bindings.HasFuncXRViewRequestViewportScale(
		this.ref,
	)
}

// FuncRequestViewportScale returns the method "XRView.requestViewportScale".
func (this XRView) FuncRequestViewportScale() (fn js.Func[func(scale float64)]) {
	bindings.FuncXRViewRequestViewportScale(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestViewportScale calls the method "XRView.requestViewportScale".
func (this XRView) RequestViewportScale(scale float64) (ret js.Void) {
	bindings.CallXRViewRequestViewportScale(
		this.ref, js.Pointer(&ret),
		float64(scale),
	)

	return
}

// TryRequestViewportScale calls the method "XRView.requestViewportScale"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRView) TryRequestViewportScale(scale float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRViewRequestViewportScale(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
	)

	return
}

func NewXRWebGLLayer(session XRSession, context XRWebGLRenderingContext, layerInit XRWebGLLayerInit) (ret XRWebGLLayer) {
	ret.ref = bindings.NewXRWebGLLayerByXRWebGLLayer(
		session.Ref(),
		context.Ref(),
		js.Pointer(&layerInit))
	return
}

func NewXRWebGLLayerByXRWebGLLayer1(session XRSession, context XRWebGLRenderingContext) (ret XRWebGLLayer) {
	ret.ref = bindings.NewXRWebGLLayerByXRWebGLLayer1(
		session.Ref(),
		context.Ref())
	return
}

type XRWebGLLayer struct {
	XRLayer
}

func (this XRWebGLLayer) Once() XRWebGLLayer {
	this.ref.Once()
	return this
}

func (this XRWebGLLayer) Ref() js.Ref {
	return this.XRLayer.Ref()
}

func (this XRWebGLLayer) FromRef(ref js.Ref) XRWebGLLayer {
	this.XRLayer = this.XRLayer.FromRef(ref)
	return this
}

func (this XRWebGLLayer) Free() {
	this.ref.Free()
}

// Antialias returns the value of property "XRWebGLLayer.antialias".
//
// It returns ok=false if there is no such property.
func (this XRWebGLLayer) Antialias() (ret bool, ok bool) {
	ok = js.True == bindings.GetXRWebGLLayerAntialias(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IgnoreDepthValues returns the value of property "XRWebGLLayer.ignoreDepthValues".
//
// It returns ok=false if there is no such property.
func (this XRWebGLLayer) IgnoreDepthValues() (ret bool, ok bool) {
	ok = js.True == bindings.GetXRWebGLLayerIgnoreDepthValues(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FixedFoveation returns the value of property "XRWebGLLayer.fixedFoveation".
//
// It returns ok=false if there is no such property.
func (this XRWebGLLayer) FixedFoveation() (ret float32, ok bool) {
	ok = js.True == bindings.GetXRWebGLLayerFixedFoveation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFixedFoveation sets the value of property "XRWebGLLayer.fixedFoveation" to val.
//
// It returns false if the property cannot be set.
func (this XRWebGLLayer) SetFixedFoveation(val float32) bool {
	return js.True == bindings.SetXRWebGLLayerFixedFoveation(
		this.ref,
		float32(val),
	)
}

// Framebuffer returns the value of property "XRWebGLLayer.framebuffer".
//
// It returns ok=false if there is no such property.
func (this XRWebGLLayer) Framebuffer() (ret WebGLFramebuffer, ok bool) {
	ok = js.True == bindings.GetXRWebGLLayerFramebuffer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FramebufferWidth returns the value of property "XRWebGLLayer.framebufferWidth".
//
// It returns ok=false if there is no such property.
func (this XRWebGLLayer) FramebufferWidth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRWebGLLayerFramebufferWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FramebufferHeight returns the value of property "XRWebGLLayer.framebufferHeight".
//
// It returns ok=false if there is no such property.
func (this XRWebGLLayer) FramebufferHeight() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRWebGLLayerFramebufferHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetViewport returns true if the method "XRWebGLLayer.getViewport" exists.
func (this XRWebGLLayer) HasFuncGetViewport() bool {
	return js.True == bindings.HasFuncXRWebGLLayerGetViewport(
		this.ref,
	)
}

// FuncGetViewport returns the method "XRWebGLLayer.getViewport".
func (this XRWebGLLayer) FuncGetViewport() (fn js.Func[func(view XRView) XRViewport]) {
	bindings.FuncXRWebGLLayerGetViewport(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetViewport calls the method "XRWebGLLayer.getViewport".
func (this XRWebGLLayer) GetViewport(view XRView) (ret XRViewport) {
	bindings.CallXRWebGLLayerGetViewport(
		this.ref, js.Pointer(&ret),
		view.Ref(),
	)

	return
}

// TryGetViewport calls the method "XRWebGLLayer.getViewport"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLLayer) TryGetViewport(view XRView) (ret XRViewport, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLLayerGetViewport(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		view.Ref(),
	)

	return
}

// HasFuncGetNativeFramebufferScaleFactor returns true if the static method "XRWebGLLayer.getNativeFramebufferScaleFactor" exists.
func (this XRWebGLLayer) HasFuncGetNativeFramebufferScaleFactor() bool {
	return js.True == bindings.HasFuncXRWebGLLayerGetNativeFramebufferScaleFactor(
		this.ref,
	)
}

// FuncGetNativeFramebufferScaleFactor returns the static method "XRWebGLLayer.getNativeFramebufferScaleFactor".
func (this XRWebGLLayer) FuncGetNativeFramebufferScaleFactor() (fn js.Func[func(session XRSession) float64]) {
	bindings.FuncXRWebGLLayerGetNativeFramebufferScaleFactor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetNativeFramebufferScaleFactor calls the static method "XRWebGLLayer.getNativeFramebufferScaleFactor".
func (this XRWebGLLayer) GetNativeFramebufferScaleFactor(session XRSession) (ret float64) {
	bindings.CallXRWebGLLayerGetNativeFramebufferScaleFactor(
		this.ref, js.Pointer(&ret),
		session.Ref(),
	)

	return
}

// TryGetNativeFramebufferScaleFactor calls the static method "XRWebGLLayer.getNativeFramebufferScaleFactor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLLayer) TryGetNativeFramebufferScaleFactor(session XRSession) (ret float64, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLLayerGetNativeFramebufferScaleFactor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		session.Ref(),
	)

	return
}

type XRLayer struct {
	EventTarget
}

func (this XRLayer) Once() XRLayer {
	this.ref.Once()
	return this
}

func (this XRLayer) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this XRLayer) FromRef(ref js.Ref) XRLayer {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this XRLayer) Free() {
	this.ref.Free()
}

type XRRenderStateInit struct {
	// DepthNear is "XRRenderStateInit.depthNear"
	//
	// Optional
	//
	// NOTE: FFI_USE_DepthNear MUST be set to true to make this field effective.
	DepthNear float64
	// DepthFar is "XRRenderStateInit.depthFar"
	//
	// Optional
	//
	// NOTE: FFI_USE_DepthFar MUST be set to true to make this field effective.
	DepthFar float64
	// InlineVerticalFieldOfView is "XRRenderStateInit.inlineVerticalFieldOfView"
	//
	// Optional
	//
	// NOTE: FFI_USE_InlineVerticalFieldOfView MUST be set to true to make this field effective.
	InlineVerticalFieldOfView float64
	// BaseLayer is "XRRenderStateInit.baseLayer"
	//
	// Optional
	BaseLayer XRWebGLLayer
	// Layers is "XRRenderStateInit.layers"
	//
	// Optional
	Layers js.Array[XRLayer]

	FFI_USE_DepthNear                 bool // for DepthNear.
	FFI_USE_DepthFar                  bool // for DepthFar.
	FFI_USE_InlineVerticalFieldOfView bool // for InlineVerticalFieldOfView.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRRenderStateInit with all fields set.
func (p XRRenderStateInit) FromRef(ref js.Ref) XRRenderStateInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRRenderStateInit in the application heap.
func (p XRRenderStateInit) New() js.Ref {
	return bindings.XRRenderStateInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRRenderStateInit) UpdateFrom(ref js.Ref) {
	bindings.XRRenderStateInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRRenderStateInit) Update(ref js.Ref) {
	bindings.XRRenderStateInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRRenderStateInit) FreeMembers(recursive bool) {
	js.Free(
		p.BaseLayer.Ref(),
		p.Layers.Ref(),
	)
	p.BaseLayer = p.BaseLayer.FromRef(js.Undefined)
	p.Layers = p.Layers.FromRef(js.Undefined)
}

type XRReferenceSpace struct {
	XRSpace
}

func (this XRReferenceSpace) Once() XRReferenceSpace {
	this.ref.Once()
	return this
}

func (this XRReferenceSpace) Ref() js.Ref {
	return this.XRSpace.Ref()
}

func (this XRReferenceSpace) FromRef(ref js.Ref) XRReferenceSpace {
	this.XRSpace = this.XRSpace.FromRef(ref)
	return this
}

func (this XRReferenceSpace) Free() {
	this.ref.Free()
}

// HasFuncGetOffsetReferenceSpace returns true if the method "XRReferenceSpace.getOffsetReferenceSpace" exists.
func (this XRReferenceSpace) HasFuncGetOffsetReferenceSpace() bool {
	return js.True == bindings.HasFuncXRReferenceSpaceGetOffsetReferenceSpace(
		this.ref,
	)
}

// FuncGetOffsetReferenceSpace returns the method "XRReferenceSpace.getOffsetReferenceSpace".
func (this XRReferenceSpace) FuncGetOffsetReferenceSpace() (fn js.Func[func(originOffset XRRigidTransform) XRReferenceSpace]) {
	bindings.FuncXRReferenceSpaceGetOffsetReferenceSpace(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetOffsetReferenceSpace calls the method "XRReferenceSpace.getOffsetReferenceSpace".
func (this XRReferenceSpace) GetOffsetReferenceSpace(originOffset XRRigidTransform) (ret XRReferenceSpace) {
	bindings.CallXRReferenceSpaceGetOffsetReferenceSpace(
		this.ref, js.Pointer(&ret),
		originOffset.Ref(),
	)

	return
}

// TryGetOffsetReferenceSpace calls the method "XRReferenceSpace.getOffsetReferenceSpace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRReferenceSpace) TryGetOffsetReferenceSpace(originOffset XRRigidTransform) (ret XRReferenceSpace, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRReferenceSpaceGetOffsetReferenceSpace(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		originOffset.Ref(),
	)

	return
}

type XRReferenceSpaceType uint32

const (
	_ XRReferenceSpaceType = iota

	XRReferenceSpaceType_VIEWER
	XRReferenceSpaceType_LOCAL
	XRReferenceSpaceType_LOCAL_FLOOR
	XRReferenceSpaceType_BOUNDED_FLOOR
	XRReferenceSpaceType_UNBOUNDED
)

func (XRReferenceSpaceType) FromRef(str js.Ref) XRReferenceSpaceType {
	return XRReferenceSpaceType(bindings.ConstOfXRReferenceSpaceType(str))
}

func (x XRReferenceSpaceType) String() (string, bool) {
	switch x {
	case XRReferenceSpaceType_VIEWER:
		return "viewer", true
	case XRReferenceSpaceType_LOCAL:
		return "local", true
	case XRReferenceSpaceType_LOCAL_FLOOR:
		return "local-floor", true
	case XRReferenceSpaceType_BOUNDED_FLOOR:
		return "bounded-floor", true
	case XRReferenceSpaceType_UNBOUNDED:
		return "unbounded", true
	default:
		return "", false
	}
}

type XRFrameRequestCallbackFunc func(this js.Ref, time DOMHighResTimeStamp, frame XRFrame) js.Ref

func (fn XRFrameRequestCallbackFunc) Register() js.Func[func(time DOMHighResTimeStamp, frame XRFrame)] {
	return js.RegisterCallback[func(time DOMHighResTimeStamp, frame XRFrame)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn XRFrameRequestCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[DOMHighResTimeStamp]{}.FromRef(args[0+1]).Get(),
		XRFrame{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type XRFrameRequestCallback[T any] struct {
	Fn  func(arg T, this js.Ref, time DOMHighResTimeStamp, frame XRFrame) js.Ref
	Arg T
}

func (cb *XRFrameRequestCallback[T]) Register() js.Func[func(time DOMHighResTimeStamp, frame XRFrame)] {
	return js.RegisterCallback[func(time DOMHighResTimeStamp, frame XRFrame)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *XRFrameRequestCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[DOMHighResTimeStamp]{}.FromRef(args[0+1]).Get(),
		XRFrame{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type XRViewerPose struct {
	XRPose
}

func (this XRViewerPose) Once() XRViewerPose {
	this.ref.Once()
	return this
}

func (this XRViewerPose) Ref() js.Ref {
	return this.XRPose.Ref()
}

func (this XRViewerPose) FromRef(ref js.Ref) XRViewerPose {
	this.XRPose = this.XRPose.FromRef(ref)
	return this
}

func (this XRViewerPose) Free() {
	this.ref.Free()
}

// Views returns the value of property "XRViewerPose.views".
//
// It returns ok=false if there is no such property.
func (this XRViewerPose) Views() (ret js.FrozenArray[XRView], ok bool) {
	ok = js.True == bindings.GetXRViewerPoseViews(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRPose struct {
	ref js.Ref
}

func (this XRPose) Once() XRPose {
	this.ref.Once()
	return this
}

func (this XRPose) Ref() js.Ref {
	return this.ref
}

func (this XRPose) FromRef(ref js.Ref) XRPose {
	this.ref = ref
	return this
}

func (this XRPose) Free() {
	this.ref.Free()
}

// Transform returns the value of property "XRPose.transform".
//
// It returns ok=false if there is no such property.
func (this XRPose) Transform() (ret XRRigidTransform, ok bool) {
	ok = js.True == bindings.GetXRPoseTransform(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LinearVelocity returns the value of property "XRPose.linearVelocity".
//
// It returns ok=false if there is no such property.
func (this XRPose) LinearVelocity() (ret DOMPointReadOnly, ok bool) {
	ok = js.True == bindings.GetXRPoseLinearVelocity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AngularVelocity returns the value of property "XRPose.angularVelocity".
//
// It returns ok=false if there is no such property.
func (this XRPose) AngularVelocity() (ret DOMPointReadOnly, ok bool) {
	ok = js.True == bindings.GetXRPoseAngularVelocity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EmulatedPosition returns the value of property "XRPose.emulatedPosition".
//
// It returns ok=false if there is no such property.
func (this XRPose) EmulatedPosition() (ret bool, ok bool) {
	ok = js.True == bindings.GetXRPoseEmulatedPosition(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRSpace struct {
	EventTarget
}

func (this XRSpace) Once() XRSpace {
	this.ref.Once()
	return this
}

func (this XRSpace) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this XRSpace) FromRef(ref js.Ref) XRSpace {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this XRSpace) Free() {
	this.ref.Free()
}

type XRAnchor struct {
	ref js.Ref
}

func (this XRAnchor) Once() XRAnchor {
	this.ref.Once()
	return this
}

func (this XRAnchor) Ref() js.Ref {
	return this.ref
}

func (this XRAnchor) FromRef(ref js.Ref) XRAnchor {
	this.ref = ref
	return this
}

func (this XRAnchor) Free() {
	this.ref.Free()
}

// AnchorSpace returns the value of property "XRAnchor.anchorSpace".
//
// It returns ok=false if there is no such property.
func (this XRAnchor) AnchorSpace() (ret XRSpace, ok bool) {
	ok = js.True == bindings.GetXRAnchorAnchorSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncRequestPersistentHandle returns true if the method "XRAnchor.requestPersistentHandle" exists.
func (this XRAnchor) HasFuncRequestPersistentHandle() bool {
	return js.True == bindings.HasFuncXRAnchorRequestPersistentHandle(
		this.ref,
	)
}

// FuncRequestPersistentHandle returns the method "XRAnchor.requestPersistentHandle".
func (this XRAnchor) FuncRequestPersistentHandle() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncXRAnchorRequestPersistentHandle(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestPersistentHandle calls the method "XRAnchor.requestPersistentHandle".
func (this XRAnchor) RequestPersistentHandle() (ret js.Promise[js.String]) {
	bindings.CallXRAnchorRequestPersistentHandle(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestPersistentHandle calls the method "XRAnchor.requestPersistentHandle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRAnchor) TryRequestPersistentHandle() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRAnchorRequestPersistentHandle(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDelete returns true if the method "XRAnchor.delete" exists.
func (this XRAnchor) HasFuncDelete() bool {
	return js.True == bindings.HasFuncXRAnchorDelete(
		this.ref,
	)
}

// FuncDelete returns the method "XRAnchor.delete".
func (this XRAnchor) FuncDelete() (fn js.Func[func()]) {
	bindings.FuncXRAnchorDelete(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Delete calls the method "XRAnchor.delete".
func (this XRAnchor) Delete() (ret js.Void) {
	bindings.CallXRAnchorDelete(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDelete calls the method "XRAnchor.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRAnchor) TryDelete() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRAnchorDelete(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type XRLightEstimate struct {
	ref js.Ref
}

func (this XRLightEstimate) Once() XRLightEstimate {
	this.ref.Once()
	return this
}

func (this XRLightEstimate) Ref() js.Ref {
	return this.ref
}

func (this XRLightEstimate) FromRef(ref js.Ref) XRLightEstimate {
	this.ref = ref
	return this
}

func (this XRLightEstimate) Free() {
	this.ref.Free()
}

// SphericalHarmonicsCoefficients returns the value of property "XRLightEstimate.sphericalHarmonicsCoefficients".
//
// It returns ok=false if there is no such property.
func (this XRLightEstimate) SphericalHarmonicsCoefficients() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetXRLightEstimateSphericalHarmonicsCoefficients(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PrimaryLightDirection returns the value of property "XRLightEstimate.primaryLightDirection".
//
// It returns ok=false if there is no such property.
func (this XRLightEstimate) PrimaryLightDirection() (ret DOMPointReadOnly, ok bool) {
	ok = js.True == bindings.GetXRLightEstimatePrimaryLightDirection(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PrimaryLightIntensity returns the value of property "XRLightEstimate.primaryLightIntensity".
//
// It returns ok=false if there is no such property.
func (this XRLightEstimate) PrimaryLightIntensity() (ret DOMPointReadOnly, ok bool) {
	ok = js.True == bindings.GetXRLightEstimatePrimaryLightIntensity(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRLightProbe struct {
	EventTarget
}

func (this XRLightProbe) Once() XRLightProbe {
	this.ref.Once()
	return this
}

func (this XRLightProbe) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this XRLightProbe) FromRef(ref js.Ref) XRLightProbe {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this XRLightProbe) Free() {
	this.ref.Free()
}

// ProbeSpace returns the value of property "XRLightProbe.probeSpace".
//
// It returns ok=false if there is no such property.
func (this XRLightProbe) ProbeSpace() (ret XRSpace, ok bool) {
	ok = js.True == bindings.GetXRLightProbeProbeSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRCPUDepthInformation struct {
	XRDepthInformation
}

func (this XRCPUDepthInformation) Once() XRCPUDepthInformation {
	this.ref.Once()
	return this
}

func (this XRCPUDepthInformation) Ref() js.Ref {
	return this.XRDepthInformation.Ref()
}

func (this XRCPUDepthInformation) FromRef(ref js.Ref) XRCPUDepthInformation {
	this.XRDepthInformation = this.XRDepthInformation.FromRef(ref)
	return this
}

func (this XRCPUDepthInformation) Free() {
	this.ref.Free()
}

// Data returns the value of property "XRCPUDepthInformation.data".
//
// It returns ok=false if there is no such property.
func (this XRCPUDepthInformation) Data() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetXRCPUDepthInformationData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetDepthInMeters returns true if the method "XRCPUDepthInformation.getDepthInMeters" exists.
func (this XRCPUDepthInformation) HasFuncGetDepthInMeters() bool {
	return js.True == bindings.HasFuncXRCPUDepthInformationGetDepthInMeters(
		this.ref,
	)
}

// FuncGetDepthInMeters returns the method "XRCPUDepthInformation.getDepthInMeters".
func (this XRCPUDepthInformation) FuncGetDepthInMeters() (fn js.Func[func(x float32, y float32) float32]) {
	bindings.FuncXRCPUDepthInformationGetDepthInMeters(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDepthInMeters calls the method "XRCPUDepthInformation.getDepthInMeters".
func (this XRCPUDepthInformation) GetDepthInMeters(x float32, y float32) (ret float32) {
	bindings.CallXRCPUDepthInformationGetDepthInMeters(
		this.ref, js.Pointer(&ret),
		float32(x),
		float32(y),
	)

	return
}

// TryGetDepthInMeters calls the method "XRCPUDepthInformation.getDepthInMeters"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRCPUDepthInformation) TryGetDepthInMeters(x float32, y float32) (ret float32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRCPUDepthInformationGetDepthInMeters(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(x),
		float32(y),
	)

	return
}

type XRJointPose struct {
	XRPose
}

func (this XRJointPose) Once() XRJointPose {
	this.ref.Once()
	return this
}

func (this XRJointPose) Ref() js.Ref {
	return this.XRPose.Ref()
}

func (this XRJointPose) FromRef(ref js.Ref) XRJointPose {
	this.XRPose = this.XRPose.FromRef(ref)
	return this
}

func (this XRJointPose) Free() {
	this.ref.Free()
}

// Radius returns the value of property "XRJointPose.radius".
//
// It returns ok=false if there is no such property.
func (this XRJointPose) Radius() (ret float32, ok bool) {
	ok = js.True == bindings.GetXRJointPoseRadius(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRHandJoint uint32

const (
	_ XRHandJoint = iota

	XRHandJoint_WRIST
	XRHandJoint_THUMB_METACARPAL
	XRHandJoint_THUMB_PHALANX_PROXIMAL
	XRHandJoint_THUMB_PHALANX_DISTAL
	XRHandJoint_THUMB_TIP
	XRHandJoint_INDEX_FINGER_METACARPAL
	XRHandJoint_INDEX_FINGER_PHALANX_PROXIMAL
	XRHandJoint_INDEX_FINGER_PHALANX_INTERMEDIATE
	XRHandJoint_INDEX_FINGER_PHALANX_DISTAL
	XRHandJoint_INDEX_FINGER_TIP
	XRHandJoint_MIDDLE_FINGER_METACARPAL
	XRHandJoint_MIDDLE_FINGER_PHALANX_PROXIMAL
	XRHandJoint_MIDDLE_FINGER_PHALANX_INTERMEDIATE
	XRHandJoint_MIDDLE_FINGER_PHALANX_DISTAL
	XRHandJoint_MIDDLE_FINGER_TIP
	XRHandJoint_RING_FINGER_METACARPAL
	XRHandJoint_RING_FINGER_PHALANX_PROXIMAL
	XRHandJoint_RING_FINGER_PHALANX_INTERMEDIATE
	XRHandJoint_RING_FINGER_PHALANX_DISTAL
	XRHandJoint_RING_FINGER_TIP
	XRHandJoint_PINKY_FINGER_METACARPAL
	XRHandJoint_PINKY_FINGER_PHALANX_PROXIMAL
	XRHandJoint_PINKY_FINGER_PHALANX_INTERMEDIATE
	XRHandJoint_PINKY_FINGER_PHALANX_DISTAL
	XRHandJoint_PINKY_FINGER_TIP
)

func (XRHandJoint) FromRef(str js.Ref) XRHandJoint {
	return XRHandJoint(bindings.ConstOfXRHandJoint(str))
}

func (x XRHandJoint) String() (string, bool) {
	switch x {
	case XRHandJoint_WRIST:
		return "wrist", true
	case XRHandJoint_THUMB_METACARPAL:
		return "thumb-metacarpal", true
	case XRHandJoint_THUMB_PHALANX_PROXIMAL:
		return "thumb-phalanx-proximal", true
	case XRHandJoint_THUMB_PHALANX_DISTAL:
		return "thumb-phalanx-distal", true
	case XRHandJoint_THUMB_TIP:
		return "thumb-tip", true
	case XRHandJoint_INDEX_FINGER_METACARPAL:
		return "index-finger-metacarpal", true
	case XRHandJoint_INDEX_FINGER_PHALANX_PROXIMAL:
		return "index-finger-phalanx-proximal", true
	case XRHandJoint_INDEX_FINGER_PHALANX_INTERMEDIATE:
		return "index-finger-phalanx-intermediate", true
	case XRHandJoint_INDEX_FINGER_PHALANX_DISTAL:
		return "index-finger-phalanx-distal", true
	case XRHandJoint_INDEX_FINGER_TIP:
		return "index-finger-tip", true
	case XRHandJoint_MIDDLE_FINGER_METACARPAL:
		return "middle-finger-metacarpal", true
	case XRHandJoint_MIDDLE_FINGER_PHALANX_PROXIMAL:
		return "middle-finger-phalanx-proximal", true
	case XRHandJoint_MIDDLE_FINGER_PHALANX_INTERMEDIATE:
		return "middle-finger-phalanx-intermediate", true
	case XRHandJoint_MIDDLE_FINGER_PHALANX_DISTAL:
		return "middle-finger-phalanx-distal", true
	case XRHandJoint_MIDDLE_FINGER_TIP:
		return "middle-finger-tip", true
	case XRHandJoint_RING_FINGER_METACARPAL:
		return "ring-finger-metacarpal", true
	case XRHandJoint_RING_FINGER_PHALANX_PROXIMAL:
		return "ring-finger-phalanx-proximal", true
	case XRHandJoint_RING_FINGER_PHALANX_INTERMEDIATE:
		return "ring-finger-phalanx-intermediate", true
	case XRHandJoint_RING_FINGER_PHALANX_DISTAL:
		return "ring-finger-phalanx-distal", true
	case XRHandJoint_RING_FINGER_TIP:
		return "ring-finger-tip", true
	case XRHandJoint_PINKY_FINGER_METACARPAL:
		return "pinky-finger-metacarpal", true
	case XRHandJoint_PINKY_FINGER_PHALANX_PROXIMAL:
		return "pinky-finger-phalanx-proximal", true
	case XRHandJoint_PINKY_FINGER_PHALANX_INTERMEDIATE:
		return "pinky-finger-phalanx-intermediate", true
	case XRHandJoint_PINKY_FINGER_PHALANX_DISTAL:
		return "pinky-finger-phalanx-distal", true
	case XRHandJoint_PINKY_FINGER_TIP:
		return "pinky-finger-tip", true
	default:
		return "", false
	}
}

type XRJointSpace struct {
	XRSpace
}

func (this XRJointSpace) Once() XRJointSpace {
	this.ref.Once()
	return this
}

func (this XRJointSpace) Ref() js.Ref {
	return this.XRSpace.Ref()
}

func (this XRJointSpace) FromRef(ref js.Ref) XRJointSpace {
	this.XRSpace = this.XRSpace.FromRef(ref)
	return this
}

func (this XRJointSpace) Free() {
	this.ref.Free()
}

// JointName returns the value of property "XRJointSpace.jointName".
//
// It returns ok=false if there is no such property.
func (this XRJointSpace) JointName() (ret XRHandJoint, ok bool) {
	ok = js.True == bindings.GetXRJointSpaceJointName(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRHitTestResult struct {
	ref js.Ref
}

func (this XRHitTestResult) Once() XRHitTestResult {
	this.ref.Once()
	return this
}

func (this XRHitTestResult) Ref() js.Ref {
	return this.ref
}

func (this XRHitTestResult) FromRef(ref js.Ref) XRHitTestResult {
	this.ref = ref
	return this
}

func (this XRHitTestResult) Free() {
	this.ref.Free()
}

// HasFuncGetPose returns true if the method "XRHitTestResult.getPose" exists.
func (this XRHitTestResult) HasFuncGetPose() bool {
	return js.True == bindings.HasFuncXRHitTestResultGetPose(
		this.ref,
	)
}

// FuncGetPose returns the method "XRHitTestResult.getPose".
func (this XRHitTestResult) FuncGetPose() (fn js.Func[func(baseSpace XRSpace) XRPose]) {
	bindings.FuncXRHitTestResultGetPose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetPose calls the method "XRHitTestResult.getPose".
func (this XRHitTestResult) GetPose(baseSpace XRSpace) (ret XRPose) {
	bindings.CallXRHitTestResultGetPose(
		this.ref, js.Pointer(&ret),
		baseSpace.Ref(),
	)

	return
}

// TryGetPose calls the method "XRHitTestResult.getPose"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRHitTestResult) TryGetPose(baseSpace XRSpace) (ret XRPose, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRHitTestResultGetPose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		baseSpace.Ref(),
	)

	return
}

// HasFuncCreateAnchor returns true if the method "XRHitTestResult.createAnchor" exists.
func (this XRHitTestResult) HasFuncCreateAnchor() bool {
	return js.True == bindings.HasFuncXRHitTestResultCreateAnchor(
		this.ref,
	)
}

// FuncCreateAnchor returns the method "XRHitTestResult.createAnchor".
func (this XRHitTestResult) FuncCreateAnchor() (fn js.Func[func() js.Promise[XRAnchor]]) {
	bindings.FuncXRHitTestResultCreateAnchor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateAnchor calls the method "XRHitTestResult.createAnchor".
func (this XRHitTestResult) CreateAnchor() (ret js.Promise[XRAnchor]) {
	bindings.CallXRHitTestResultCreateAnchor(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateAnchor calls the method "XRHitTestResult.createAnchor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRHitTestResult) TryCreateAnchor() (ret js.Promise[XRAnchor], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRHitTestResultCreateAnchor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
