// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

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

// New creates a new {0x140004cc0e0 USBControlTransferParameters USBControlTransferParameters [// USBControlTransferParameters] [0x140006febe0 0x140006fec80 0x140006fed20 0x140006fedc0 0x140006fee60] 0x14000780090 {0 0}} in the application heap.
func (p USBControlTransferParameters) New() js.Ref {
	return bindings.USBControlTransferParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p USBControlTransferParameters) UpdateFrom(ref js.Ref) {
	bindings.USBControlTransferParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p USBControlTransferParameters) Update(ref js.Ref) {
	bindings.USBControlTransferParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewUSBOutTransferResult(status USBTransferStatus, bytesWritten uint32) USBOutTransferResult {
	return USBOutTransferResult{}.FromRef(
		bindings.NewUSBOutTransferResultByUSBOutTransferResult(
			uint32(status),
			uint32(bytesWritten)),
	)
}

func NewUSBOutTransferResultByUSBOutTransferResult1(status USBTransferStatus) USBOutTransferResult {
	return USBOutTransferResult{}.FromRef(
		bindings.NewUSBOutTransferResultByUSBOutTransferResult1(
			uint32(status)),
	)
}

type USBOutTransferResult struct {
	ref js.Ref
}

func (this USBOutTransferResult) Once() USBOutTransferResult {
	this.Ref().Once()
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
	this.Ref().Free()
}

// BytesWritten returns the value of property "USBOutTransferResult.bytesWritten".
//
// The returned bool will be false if there is no such property.
func (this USBOutTransferResult) BytesWritten() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetUSBOutTransferResultBytesWritten(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Status returns the value of property "USBOutTransferResult.status".
//
// The returned bool will be false if there is no such property.
func (this USBOutTransferResult) Status() (USBTransferStatus, bool) {
	var _ok bool
	_ret := bindings.GetUSBOutTransferResultStatus(
		this.Ref(), js.Pointer(&_ok),
	)
	return USBTransferStatus(_ret), _ok
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

func NewUSBIsochronousInTransferPacket(status USBTransferStatus, data js.DataView) USBIsochronousInTransferPacket {
	return USBIsochronousInTransferPacket{}.FromRef(
		bindings.NewUSBIsochronousInTransferPacketByUSBIsochronousInTransferPacket(
			uint32(status),
			data.Ref()),
	)
}

func NewUSBIsochronousInTransferPacketByUSBIsochronousInTransferPacket1(status USBTransferStatus) USBIsochronousInTransferPacket {
	return USBIsochronousInTransferPacket{}.FromRef(
		bindings.NewUSBIsochronousInTransferPacketByUSBIsochronousInTransferPacket1(
			uint32(status)),
	)
}

type USBIsochronousInTransferPacket struct {
	ref js.Ref
}

func (this USBIsochronousInTransferPacket) Once() USBIsochronousInTransferPacket {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Data returns the value of property "USBIsochronousInTransferPacket.data".
//
// The returned bool will be false if there is no such property.
func (this USBIsochronousInTransferPacket) Data() (js.DataView, bool) {
	var _ok bool
	_ret := bindings.GetUSBIsochronousInTransferPacketData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.DataView{}.FromRef(_ret), _ok
}

// Status returns the value of property "USBIsochronousInTransferPacket.status".
//
// The returned bool will be false if there is no such property.
func (this USBIsochronousInTransferPacket) Status() (USBTransferStatus, bool) {
	var _ok bool
	_ret := bindings.GetUSBIsochronousInTransferPacketStatus(
		this.Ref(), js.Pointer(&_ok),
	)
	return USBTransferStatus(_ret), _ok
}

func NewUSBIsochronousInTransferResult(packets js.Array[USBIsochronousInTransferPacket], data js.DataView) USBIsochronousInTransferResult {
	return USBIsochronousInTransferResult{}.FromRef(
		bindings.NewUSBIsochronousInTransferResultByUSBIsochronousInTransferResult(
			packets.Ref(),
			data.Ref()),
	)
}

func NewUSBIsochronousInTransferResultByUSBIsochronousInTransferResult1(packets js.Array[USBIsochronousInTransferPacket]) USBIsochronousInTransferResult {
	return USBIsochronousInTransferResult{}.FromRef(
		bindings.NewUSBIsochronousInTransferResultByUSBIsochronousInTransferResult1(
			packets.Ref()),
	)
}

type USBIsochronousInTransferResult struct {
	ref js.Ref
}

func (this USBIsochronousInTransferResult) Once() USBIsochronousInTransferResult {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Data returns the value of property "USBIsochronousInTransferResult.data".
//
// The returned bool will be false if there is no such property.
func (this USBIsochronousInTransferResult) Data() (js.DataView, bool) {
	var _ok bool
	_ret := bindings.GetUSBIsochronousInTransferResultData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.DataView{}.FromRef(_ret), _ok
}

// Packets returns the value of property "USBIsochronousInTransferResult.packets".
//
// The returned bool will be false if there is no such property.
func (this USBIsochronousInTransferResult) Packets() (js.FrozenArray[USBIsochronousInTransferPacket], bool) {
	var _ok bool
	_ret := bindings.GetUSBIsochronousInTransferResultPackets(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[USBIsochronousInTransferPacket]{}.FromRef(_ret), _ok
}

func NewUSBIsochronousOutTransferPacket(status USBTransferStatus, bytesWritten uint32) USBIsochronousOutTransferPacket {
	return USBIsochronousOutTransferPacket{}.FromRef(
		bindings.NewUSBIsochronousOutTransferPacketByUSBIsochronousOutTransferPacket(
			uint32(status),
			uint32(bytesWritten)),
	)
}

func NewUSBIsochronousOutTransferPacketByUSBIsochronousOutTransferPacket1(status USBTransferStatus) USBIsochronousOutTransferPacket {
	return USBIsochronousOutTransferPacket{}.FromRef(
		bindings.NewUSBIsochronousOutTransferPacketByUSBIsochronousOutTransferPacket1(
			uint32(status)),
	)
}

type USBIsochronousOutTransferPacket struct {
	ref js.Ref
}

func (this USBIsochronousOutTransferPacket) Once() USBIsochronousOutTransferPacket {
	this.Ref().Once()
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
	this.Ref().Free()
}

// BytesWritten returns the value of property "USBIsochronousOutTransferPacket.bytesWritten".
//
// The returned bool will be false if there is no such property.
func (this USBIsochronousOutTransferPacket) BytesWritten() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetUSBIsochronousOutTransferPacketBytesWritten(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Status returns the value of property "USBIsochronousOutTransferPacket.status".
//
// The returned bool will be false if there is no such property.
func (this USBIsochronousOutTransferPacket) Status() (USBTransferStatus, bool) {
	var _ok bool
	_ret := bindings.GetUSBIsochronousOutTransferPacketStatus(
		this.Ref(), js.Pointer(&_ok),
	)
	return USBTransferStatus(_ret), _ok
}

func NewUSBIsochronousOutTransferResult(packets js.Array[USBIsochronousOutTransferPacket]) USBIsochronousOutTransferResult {
	return USBIsochronousOutTransferResult{}.FromRef(
		bindings.NewUSBIsochronousOutTransferResultByUSBIsochronousOutTransferResult(
			packets.Ref()),
	)
}

type USBIsochronousOutTransferResult struct {
	ref js.Ref
}

func (this USBIsochronousOutTransferResult) Once() USBIsochronousOutTransferResult {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Packets returns the value of property "USBIsochronousOutTransferResult.packets".
//
// The returned bool will be false if there is no such property.
func (this USBIsochronousOutTransferResult) Packets() (js.FrozenArray[USBIsochronousOutTransferPacket], bool) {
	var _ok bool
	_ret := bindings.GetUSBIsochronousOutTransferResultPackets(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[USBIsochronousOutTransferPacket]{}.FromRef(_ret), _ok
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

func NewUSBEndpoint(alternate USBAlternateInterface, endpointNumber uint8, direction USBDirection) USBEndpoint {
	return USBEndpoint{}.FromRef(
		bindings.NewUSBEndpointByUSBEndpoint(
			alternate.Ref(),
			uint32(endpointNumber),
			uint32(direction)),
	)
}

type USBEndpoint struct {
	ref js.Ref
}

func (this USBEndpoint) Once() USBEndpoint {
	this.Ref().Once()
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
	this.Ref().Free()
}

// EndpointNumber returns the value of property "USBEndpoint.endpointNumber".
//
// The returned bool will be false if there is no such property.
func (this USBEndpoint) EndpointNumber() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBEndpointEndpointNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// Direction returns the value of property "USBEndpoint.direction".
//
// The returned bool will be false if there is no such property.
func (this USBEndpoint) Direction() (USBDirection, bool) {
	var _ok bool
	_ret := bindings.GetUSBEndpointDirection(
		this.Ref(), js.Pointer(&_ok),
	)
	return USBDirection(_ret), _ok
}

// Type returns the value of property "USBEndpoint.type".
//
// The returned bool will be false if there is no such property.
func (this USBEndpoint) Type() (USBEndpointType, bool) {
	var _ok bool
	_ret := bindings.GetUSBEndpointType(
		this.Ref(), js.Pointer(&_ok),
	)
	return USBEndpointType(_ret), _ok
}

// PacketSize returns the value of property "USBEndpoint.packetSize".
//
// The returned bool will be false if there is no such property.
func (this USBEndpoint) PacketSize() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetUSBEndpointPacketSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

func NewUSBAlternateInterface(deviceInterface USBInterface, alternateSetting uint8) USBAlternateInterface {
	return USBAlternateInterface{}.FromRef(
		bindings.NewUSBAlternateInterfaceByUSBAlternateInterface(
			deviceInterface.Ref(),
			uint32(alternateSetting)),
	)
}

type USBAlternateInterface struct {
	ref js.Ref
}

func (this USBAlternateInterface) Once() USBAlternateInterface {
	this.Ref().Once()
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
	this.Ref().Free()
}

// AlternateSetting returns the value of property "USBAlternateInterface.alternateSetting".
//
// The returned bool will be false if there is no such property.
func (this USBAlternateInterface) AlternateSetting() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBAlternateInterfaceAlternateSetting(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// InterfaceClass returns the value of property "USBAlternateInterface.interfaceClass".
//
// The returned bool will be false if there is no such property.
func (this USBAlternateInterface) InterfaceClass() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBAlternateInterfaceInterfaceClass(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// InterfaceSubclass returns the value of property "USBAlternateInterface.interfaceSubclass".
//
// The returned bool will be false if there is no such property.
func (this USBAlternateInterface) InterfaceSubclass() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBAlternateInterfaceInterfaceSubclass(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// InterfaceProtocol returns the value of property "USBAlternateInterface.interfaceProtocol".
//
// The returned bool will be false if there is no such property.
func (this USBAlternateInterface) InterfaceProtocol() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBAlternateInterfaceInterfaceProtocol(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// InterfaceName returns the value of property "USBAlternateInterface.interfaceName".
//
// The returned bool will be false if there is no such property.
func (this USBAlternateInterface) InterfaceName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetUSBAlternateInterfaceInterfaceName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Endpoints returns the value of property "USBAlternateInterface.endpoints".
//
// The returned bool will be false if there is no such property.
func (this USBAlternateInterface) Endpoints() (js.FrozenArray[USBEndpoint], bool) {
	var _ok bool
	_ret := bindings.GetUSBAlternateInterfaceEndpoints(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[USBEndpoint]{}.FromRef(_ret), _ok
}

func NewUSBInterface(configuration USBConfiguration, interfaceNumber uint8) USBInterface {
	return USBInterface{}.FromRef(
		bindings.NewUSBInterfaceByUSBInterface(
			configuration.Ref(),
			uint32(interfaceNumber)),
	)
}

type USBInterface struct {
	ref js.Ref
}

func (this USBInterface) Once() USBInterface {
	this.Ref().Once()
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
	this.Ref().Free()
}

// InterfaceNumber returns the value of property "USBInterface.interfaceNumber".
//
// The returned bool will be false if there is no such property.
func (this USBInterface) InterfaceNumber() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBInterfaceInterfaceNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// Alternate returns the value of property "USBInterface.alternate".
//
// The returned bool will be false if there is no such property.
func (this USBInterface) Alternate() (USBAlternateInterface, bool) {
	var _ok bool
	_ret := bindings.GetUSBInterfaceAlternate(
		this.Ref(), js.Pointer(&_ok),
	)
	return USBAlternateInterface{}.FromRef(_ret), _ok
}

// Alternates returns the value of property "USBInterface.alternates".
//
// The returned bool will be false if there is no such property.
func (this USBInterface) Alternates() (js.FrozenArray[USBAlternateInterface], bool) {
	var _ok bool
	_ret := bindings.GetUSBInterfaceAlternates(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[USBAlternateInterface]{}.FromRef(_ret), _ok
}

// Claimed returns the value of property "USBInterface.claimed".
//
// The returned bool will be false if there is no such property.
func (this USBInterface) Claimed() (bool, bool) {
	var _ok bool
	_ret := bindings.GetUSBInterfaceClaimed(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

func NewUSBConfiguration(device USBDevice, configurationValue uint8) USBConfiguration {
	return USBConfiguration{}.FromRef(
		bindings.NewUSBConfigurationByUSBConfiguration(
			device.Ref(),
			uint32(configurationValue)),
	)
}

type USBConfiguration struct {
	ref js.Ref
}

func (this USBConfiguration) Once() USBConfiguration {
	this.Ref().Once()
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
	this.Ref().Free()
}

// ConfigurationValue returns the value of property "USBConfiguration.configurationValue".
//
// The returned bool will be false if there is no such property.
func (this USBConfiguration) ConfigurationValue() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBConfigurationConfigurationValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// ConfigurationName returns the value of property "USBConfiguration.configurationName".
//
// The returned bool will be false if there is no such property.
func (this USBConfiguration) ConfigurationName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetUSBConfigurationConfigurationName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Interfaces returns the value of property "USBConfiguration.interfaces".
//
// The returned bool will be false if there is no such property.
func (this USBConfiguration) Interfaces() (js.FrozenArray[USBInterface], bool) {
	var _ok bool
	_ret := bindings.GetUSBConfigurationInterfaces(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[USBInterface]{}.FromRef(_ret), _ok
}

type USBDevice struct {
	ref js.Ref
}

func (this USBDevice) Once() USBDevice {
	this.Ref().Once()
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
	this.Ref().Free()
}

// UsbVersionMajor returns the value of property "USBDevice.usbVersionMajor".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) UsbVersionMajor() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceUsbVersionMajor(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// UsbVersionMinor returns the value of property "USBDevice.usbVersionMinor".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) UsbVersionMinor() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceUsbVersionMinor(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// UsbVersionSubminor returns the value of property "USBDevice.usbVersionSubminor".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) UsbVersionSubminor() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceUsbVersionSubminor(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// DeviceClass returns the value of property "USBDevice.deviceClass".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) DeviceClass() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceDeviceClass(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// DeviceSubclass returns the value of property "USBDevice.deviceSubclass".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) DeviceSubclass() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceDeviceSubclass(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// DeviceProtocol returns the value of property "USBDevice.deviceProtocol".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) DeviceProtocol() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceDeviceProtocol(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// VendorId returns the value of property "USBDevice.vendorId".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) VendorId() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceVendorId(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// ProductId returns the value of property "USBDevice.productId".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) ProductId() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceProductId(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// DeviceVersionMajor returns the value of property "USBDevice.deviceVersionMajor".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) DeviceVersionMajor() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceDeviceVersionMajor(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// DeviceVersionMinor returns the value of property "USBDevice.deviceVersionMinor".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) DeviceVersionMinor() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceDeviceVersionMinor(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// DeviceVersionSubminor returns the value of property "USBDevice.deviceVersionSubminor".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) DeviceVersionSubminor() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceDeviceVersionSubminor(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// ManufacturerName returns the value of property "USBDevice.manufacturerName".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) ManufacturerName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceManufacturerName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ProductName returns the value of property "USBDevice.productName".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) ProductName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceProductName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SerialNumber returns the value of property "USBDevice.serialNumber".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) SerialNumber() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceSerialNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Configuration returns the value of property "USBDevice.configuration".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) Configuration() (USBConfiguration, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceConfiguration(
		this.Ref(), js.Pointer(&_ok),
	)
	return USBConfiguration{}.FromRef(_ret), _ok
}

// Configurations returns the value of property "USBDevice.configurations".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) Configurations() (js.FrozenArray[USBConfiguration], bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceConfigurations(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[USBConfiguration]{}.FromRef(_ret), _ok
}

// Opened returns the value of property "USBDevice.opened".
//
// The returned bool will be false if there is no such property.
func (this USBDevice) Opened() (bool, bool) {
	var _ok bool
	_ret := bindings.GetUSBDeviceOpened(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Open calls the method "USBDevice.open".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) Open() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceOpen(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// OpenFunc returns the method "USBDevice.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) OpenFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.USBDeviceOpenFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "USBDevice.close".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) Close() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceClose(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CloseFunc returns the method "USBDevice.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) CloseFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.USBDeviceCloseFunc(
			this.Ref(),
		),
	)
}

// Forget calls the method "USBDevice.forget".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) Forget() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceForget(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ForgetFunc returns the method "USBDevice.forget".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) ForgetFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.USBDeviceForgetFunc(
			this.Ref(),
		),
	)
}

// SelectConfiguration calls the method "USBDevice.selectConfiguration".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) SelectConfiguration(configurationValue uint8) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceSelectConfiguration(
		this.Ref(), js.Pointer(&_ok),
		uint32(configurationValue),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SelectConfigurationFunc returns the method "USBDevice.selectConfiguration".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) SelectConfigurationFunc() (fn js.Func[func(configurationValue uint8) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.USBDeviceSelectConfigurationFunc(
			this.Ref(),
		),
	)
}

// ClaimInterface calls the method "USBDevice.claimInterface".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) ClaimInterface(interfaceNumber uint8) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceClaimInterface(
		this.Ref(), js.Pointer(&_ok),
		uint32(interfaceNumber),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ClaimInterfaceFunc returns the method "USBDevice.claimInterface".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) ClaimInterfaceFunc() (fn js.Func[func(interfaceNumber uint8) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.USBDeviceClaimInterfaceFunc(
			this.Ref(),
		),
	)
}

// ReleaseInterface calls the method "USBDevice.releaseInterface".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) ReleaseInterface(interfaceNumber uint8) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceReleaseInterface(
		this.Ref(), js.Pointer(&_ok),
		uint32(interfaceNumber),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ReleaseInterfaceFunc returns the method "USBDevice.releaseInterface".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) ReleaseInterfaceFunc() (fn js.Func[func(interfaceNumber uint8) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.USBDeviceReleaseInterfaceFunc(
			this.Ref(),
		),
	)
}

// SelectAlternateInterface calls the method "USBDevice.selectAlternateInterface".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) SelectAlternateInterface(interfaceNumber uint8, alternateSetting uint8) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceSelectAlternateInterface(
		this.Ref(), js.Pointer(&_ok),
		uint32(interfaceNumber),
		uint32(alternateSetting),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SelectAlternateInterfaceFunc returns the method "USBDevice.selectAlternateInterface".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) SelectAlternateInterfaceFunc() (fn js.Func[func(interfaceNumber uint8, alternateSetting uint8) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.USBDeviceSelectAlternateInterfaceFunc(
			this.Ref(),
		),
	)
}

// ControlTransferIn calls the method "USBDevice.controlTransferIn".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) ControlTransferIn(setup USBControlTransferParameters, length uint16) (js.Promise[USBInTransferResult], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceControlTransferIn(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&setup),
		uint32(length),
	)

	return js.Promise[USBInTransferResult]{}.FromRef(_ret), _ok
}

// ControlTransferInFunc returns the method "USBDevice.controlTransferIn".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) ControlTransferInFunc() (fn js.Func[func(setup USBControlTransferParameters, length uint16) js.Promise[USBInTransferResult]]) {
	return fn.FromRef(
		bindings.USBDeviceControlTransferInFunc(
			this.Ref(),
		),
	)
}

// ControlTransferOut calls the method "USBDevice.controlTransferOut".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) ControlTransferOut(setup USBControlTransferParameters, data BufferSource) (js.Promise[USBOutTransferResult], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceControlTransferOut(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&setup),
		data.Ref(),
	)

	return js.Promise[USBOutTransferResult]{}.FromRef(_ret), _ok
}

// ControlTransferOutFunc returns the method "USBDevice.controlTransferOut".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) ControlTransferOutFunc() (fn js.Func[func(setup USBControlTransferParameters, data BufferSource) js.Promise[USBOutTransferResult]]) {
	return fn.FromRef(
		bindings.USBDeviceControlTransferOutFunc(
			this.Ref(),
		),
	)
}

// ControlTransferOut1 calls the method "USBDevice.controlTransferOut".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) ControlTransferOut1(setup USBControlTransferParameters) (js.Promise[USBOutTransferResult], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceControlTransferOut1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&setup),
	)

	return js.Promise[USBOutTransferResult]{}.FromRef(_ret), _ok
}

// ControlTransferOut1Func returns the method "USBDevice.controlTransferOut".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) ControlTransferOut1Func() (fn js.Func[func(setup USBControlTransferParameters) js.Promise[USBOutTransferResult]]) {
	return fn.FromRef(
		bindings.USBDeviceControlTransferOut1Func(
			this.Ref(),
		),
	)
}

// ClearHalt calls the method "USBDevice.clearHalt".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) ClearHalt(direction USBDirection, endpointNumber uint8) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceClearHalt(
		this.Ref(), js.Pointer(&_ok),
		uint32(direction),
		uint32(endpointNumber),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ClearHaltFunc returns the method "USBDevice.clearHalt".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) ClearHaltFunc() (fn js.Func[func(direction USBDirection, endpointNumber uint8) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.USBDeviceClearHaltFunc(
			this.Ref(),
		),
	)
}

// TransferIn calls the method "USBDevice.transferIn".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) TransferIn(endpointNumber uint8, length uint32) (js.Promise[USBInTransferResult], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceTransferIn(
		this.Ref(), js.Pointer(&_ok),
		uint32(endpointNumber),
		uint32(length),
	)

	return js.Promise[USBInTransferResult]{}.FromRef(_ret), _ok
}

// TransferInFunc returns the method "USBDevice.transferIn".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) TransferInFunc() (fn js.Func[func(endpointNumber uint8, length uint32) js.Promise[USBInTransferResult]]) {
	return fn.FromRef(
		bindings.USBDeviceTransferInFunc(
			this.Ref(),
		),
	)
}

// TransferOut calls the method "USBDevice.transferOut".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) TransferOut(endpointNumber uint8, data BufferSource) (js.Promise[USBOutTransferResult], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceTransferOut(
		this.Ref(), js.Pointer(&_ok),
		uint32(endpointNumber),
		data.Ref(),
	)

	return js.Promise[USBOutTransferResult]{}.FromRef(_ret), _ok
}

// TransferOutFunc returns the method "USBDevice.transferOut".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) TransferOutFunc() (fn js.Func[func(endpointNumber uint8, data BufferSource) js.Promise[USBOutTransferResult]]) {
	return fn.FromRef(
		bindings.USBDeviceTransferOutFunc(
			this.Ref(),
		),
	)
}

// IsochronousTransferIn calls the method "USBDevice.isochronousTransferIn".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) IsochronousTransferIn(endpointNumber uint8, packetLengths js.Array[uint32]) (js.Promise[USBIsochronousInTransferResult], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceIsochronousTransferIn(
		this.Ref(), js.Pointer(&_ok),
		uint32(endpointNumber),
		packetLengths.Ref(),
	)

	return js.Promise[USBIsochronousInTransferResult]{}.FromRef(_ret), _ok
}

// IsochronousTransferInFunc returns the method "USBDevice.isochronousTransferIn".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) IsochronousTransferInFunc() (fn js.Func[func(endpointNumber uint8, packetLengths js.Array[uint32]) js.Promise[USBIsochronousInTransferResult]]) {
	return fn.FromRef(
		bindings.USBDeviceIsochronousTransferInFunc(
			this.Ref(),
		),
	)
}

// IsochronousTransferOut calls the method "USBDevice.isochronousTransferOut".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) IsochronousTransferOut(endpointNumber uint8, data BufferSource, packetLengths js.Array[uint32]) (js.Promise[USBIsochronousOutTransferResult], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceIsochronousTransferOut(
		this.Ref(), js.Pointer(&_ok),
		uint32(endpointNumber),
		data.Ref(),
		packetLengths.Ref(),
	)

	return js.Promise[USBIsochronousOutTransferResult]{}.FromRef(_ret), _ok
}

// IsochronousTransferOutFunc returns the method "USBDevice.isochronousTransferOut".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) IsochronousTransferOutFunc() (fn js.Func[func(endpointNumber uint8, data BufferSource, packetLengths js.Array[uint32]) js.Promise[USBIsochronousOutTransferResult]]) {
	return fn.FromRef(
		bindings.USBDeviceIsochronousTransferOutFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "USBDevice.reset".
//
// The returned bool will be false if there is no such method.
func (this USBDevice) Reset() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallUSBDeviceReset(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ResetFunc returns the method "USBDevice.reset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USBDevice) ResetFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.USBDeviceResetFunc(
			this.Ref(),
		),
	)
}

type USBDeviceFilter struct {
	// VendorId is "USBDeviceFilter.vendorId"
	//
	// Optional
	VendorId uint16
	// ProductId is "USBDeviceFilter.productId"
	//
	// Optional
	ProductId uint16
	// ClassCode is "USBDeviceFilter.classCode"
	//
	// Optional
	ClassCode uint8
	// SubclassCode is "USBDeviceFilter.subclassCode"
	//
	// Optional
	SubclassCode uint8
	// ProtocolCode is "USBDeviceFilter.protocolCode"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 USBDeviceFilter USBDeviceFilter [// USBDeviceFilter] [0x140006ff4a0 0x140006ff5e0 0x140006ff720 0x140006ff860 0x140006ff9a0 0x140006ffae0 0x140006ff540 0x140006ff680 0x140006ff7c0 0x140006ff900 0x140006ffa40] 0x14000780138 {0 0}} in the application heap.
func (p USBDeviceFilter) New() js.Ref {
	return bindings.USBDeviceFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p USBDeviceFilter) UpdateFrom(ref js.Ref) {
	bindings.USBDeviceFilterJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p USBDeviceFilter) Update(ref js.Ref) {
	bindings.USBDeviceFilterJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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

// New creates a new {0x140004cc0e0 USBDeviceRequestOptions USBDeviceRequestOptions [// USBDeviceRequestOptions] [0x140006ffb80 0x140006ffc20] 0x14000780108 {0 0}} in the application heap.
func (p USBDeviceRequestOptions) New() js.Ref {
	return bindings.USBDeviceRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p USBDeviceRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.USBDeviceRequestOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p USBDeviceRequestOptions) Update(ref js.Ref) {
	bindings.USBDeviceRequestOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type USB struct {
	EventTarget
}

func (this USB) Once() USB {
	this.Ref().Once()
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
	this.Ref().Free()
}

// GetDevices calls the method "USB.getDevices".
//
// The returned bool will be false if there is no such method.
func (this USB) GetDevices() (js.Promise[js.Array[USBDevice]], bool) {
	var _ok bool
	_ret := bindings.CallUSBGetDevices(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[USBDevice]]{}.FromRef(_ret), _ok
}

// GetDevicesFunc returns the method "USB.getDevices".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USB) GetDevicesFunc() (fn js.Func[func() js.Promise[js.Array[USBDevice]]]) {
	return fn.FromRef(
		bindings.USBGetDevicesFunc(
			this.Ref(),
		),
	)
}

// RequestDevice calls the method "USB.requestDevice".
//
// The returned bool will be false if there is no such method.
func (this USB) RequestDevice(options USBDeviceRequestOptions) (js.Promise[USBDevice], bool) {
	var _ok bool
	_ret := bindings.CallUSBRequestDevice(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[USBDevice]{}.FromRef(_ret), _ok
}

// RequestDeviceFunc returns the method "USB.requestDevice".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this USB) RequestDeviceFunc() (fn js.Func[func(options USBDeviceRequestOptions) js.Promise[USBDevice]]) {
	return fn.FromRef(
		bindings.USBRequestDeviceFunc(
			this.Ref(),
		),
	)
}

type EpubReadingSystem struct {
	ref js.Ref
}

func (this EpubReadingSystem) Once() EpubReadingSystem {
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasFeature calls the method "EpubReadingSystem.hasFeature".
//
// The returned bool will be false if there is no such method.
func (this EpubReadingSystem) HasFeature(feature js.String, version js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallEpubReadingSystemHasFeature(
		this.Ref(), js.Pointer(&_ok),
		feature.Ref(),
		version.Ref(),
	)

	return _ret == js.True, _ok
}

// HasFeatureFunc returns the method "EpubReadingSystem.hasFeature".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EpubReadingSystem) HasFeatureFunc() (fn js.Func[func(feature js.String, version js.String) bool]) {
	return fn.FromRef(
		bindings.EpubReadingSystemHasFeatureFunc(
			this.Ref(),
		),
	)
}

// HasFeature1 calls the method "EpubReadingSystem.hasFeature".
//
// The returned bool will be false if there is no such method.
func (this EpubReadingSystem) HasFeature1(feature js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallEpubReadingSystemHasFeature1(
		this.Ref(), js.Pointer(&_ok),
		feature.Ref(),
	)

	return _ret == js.True, _ok
}

// HasFeature1Func returns the method "EpubReadingSystem.hasFeature".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EpubReadingSystem) HasFeature1Func() (fn js.Func[func(feature js.String) bool]) {
	return fn.FromRef(
		bindings.EpubReadingSystemHasFeature1Func(
			this.Ref(),
		),
	)
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
	Antialias bool
	// Depth is "XRWebGLLayerInit.depth"
	//
	// Optional, defaults to true.
	Depth bool
	// Stencil is "XRWebGLLayerInit.stencil"
	//
	// Optional, defaults to false.
	Stencil bool
	// Alpha is "XRWebGLLayerInit.alpha"
	//
	// Optional, defaults to true.
	Alpha bool
	// IgnoreDepthValues is "XRWebGLLayerInit.ignoreDepthValues"
	//
	// Optional, defaults to false.
	IgnoreDepthValues bool
	// FramebufferScaleFactor is "XRWebGLLayerInit.framebufferScaleFactor"
	//
	// Optional, defaults to 1.0.
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

// New creates a new {0x140004cc0e0 XRWebGLLayerInit XRWebGLLayerInit [// XRWebGLLayerInit] [0x140007a40a0 0x140007a41e0 0x140007a4320 0x140007a4460 0x140007a45a0 0x140007a46e0 0x140007a4140 0x140007a4280 0x140007a43c0 0x140007a4500 0x140007a4640 0x140007a4780] 0x14000780270 {0 0}} in the application heap.
func (p XRWebGLLayerInit) New() js.Ref {
	return bindings.XRWebGLLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRWebGLLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRWebGLLayerInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRWebGLLayerInit) Update(ref js.Ref) {
	bindings.XRWebGLLayerInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type XRViewport struct {
	ref js.Ref
}

func (this XRViewport) Once() XRViewport {
	this.Ref().Once()
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
	this.Ref().Free()
}

// X returns the value of property "XRViewport.x".
//
// The returned bool will be false if there is no such property.
func (this XRViewport) X() (int32, bool) {
	var _ok bool
	_ret := bindings.GetXRViewportX(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Y returns the value of property "XRViewport.y".
//
// The returned bool will be false if there is no such property.
func (this XRViewport) Y() (int32, bool) {
	var _ok bool
	_ret := bindings.GetXRViewportY(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Width returns the value of property "XRViewport.width".
//
// The returned bool will be false if there is no such property.
func (this XRViewport) Width() (int32, bool) {
	var _ok bool
	_ret := bindings.GetXRViewportWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Height returns the value of property "XRViewport.height".
//
// The returned bool will be false if there is no such property.
func (this XRViewport) Height() (int32, bool) {
	var _ok bool
	_ret := bindings.GetXRViewportHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
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

func NewXRRigidTransform(position DOMPointInit, orientation DOMPointInit) XRRigidTransform {
	return XRRigidTransform{}.FromRef(
		bindings.NewXRRigidTransformByXRRigidTransform(
			js.Pointer(&position),
			js.Pointer(&orientation)),
	)
}

func NewXRRigidTransformByXRRigidTransform1(position DOMPointInit) XRRigidTransform {
	return XRRigidTransform{}.FromRef(
		bindings.NewXRRigidTransformByXRRigidTransform1(
			js.Pointer(&position)),
	)
}

func NewXRRigidTransformByXRRigidTransform2() XRRigidTransform {
	return XRRigidTransform{}.FromRef(
		bindings.NewXRRigidTransformByXRRigidTransform2(),
	)
}

type XRRigidTransform struct {
	ref js.Ref
}

func (this XRRigidTransform) Once() XRRigidTransform {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Position returns the value of property "XRRigidTransform.position".
//
// The returned bool will be false if there is no such property.
func (this XRRigidTransform) Position() (DOMPointReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetXRRigidTransformPosition(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMPointReadOnly{}.FromRef(_ret), _ok
}

// Orientation returns the value of property "XRRigidTransform.orientation".
//
// The returned bool will be false if there is no such property.
func (this XRRigidTransform) Orientation() (DOMPointReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetXRRigidTransformOrientation(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMPointReadOnly{}.FromRef(_ret), _ok
}

// Matrix returns the value of property "XRRigidTransform.matrix".
//
// The returned bool will be false if there is no such property.
func (this XRRigidTransform) Matrix() (js.TypedArray[float32], bool) {
	var _ok bool
	_ret := bindings.GetXRRigidTransformMatrix(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[float32]{}.FromRef(_ret), _ok
}

// Inverse returns the value of property "XRRigidTransform.inverse".
//
// The returned bool will be false if there is no such property.
func (this XRRigidTransform) Inverse() (XRRigidTransform, bool) {
	var _ok bool
	_ret := bindings.GetXRRigidTransformInverse(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRRigidTransform{}.FromRef(_ret), _ok
}

type XRCamera struct {
	ref js.Ref
}

func (this XRCamera) Once() XRCamera {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Width returns the value of property "XRCamera.width".
//
// The returned bool will be false if there is no such property.
func (this XRCamera) Width() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRCameraWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Height returns the value of property "XRCamera.height".
//
// The returned bool will be false if there is no such property.
func (this XRCamera) Height() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRCameraHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

type XRView struct {
	ref js.Ref
}

func (this XRView) Once() XRView {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Eye returns the value of property "XRView.eye".
//
// The returned bool will be false if there is no such property.
func (this XRView) Eye() (XREye, bool) {
	var _ok bool
	_ret := bindings.GetXRViewEye(
		this.Ref(), js.Pointer(&_ok),
	)
	return XREye(_ret), _ok
}

// ProjectionMatrix returns the value of property "XRView.projectionMatrix".
//
// The returned bool will be false if there is no such property.
func (this XRView) ProjectionMatrix() (js.TypedArray[float32], bool) {
	var _ok bool
	_ret := bindings.GetXRViewProjectionMatrix(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[float32]{}.FromRef(_ret), _ok
}

// Transform returns the value of property "XRView.transform".
//
// The returned bool will be false if there is no such property.
func (this XRView) Transform() (XRRigidTransform, bool) {
	var _ok bool
	_ret := bindings.GetXRViewTransform(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRRigidTransform{}.FromRef(_ret), _ok
}

// RecommendedViewportScale returns the value of property "XRView.recommendedViewportScale".
//
// The returned bool will be false if there is no such property.
func (this XRView) RecommendedViewportScale() (float64, bool) {
	var _ok bool
	_ret := bindings.GetXRViewRecommendedViewportScale(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// IsFirstPersonObserver returns the value of property "XRView.isFirstPersonObserver".
//
// The returned bool will be false if there is no such property.
func (this XRView) IsFirstPersonObserver() (bool, bool) {
	var _ok bool
	_ret := bindings.GetXRViewIsFirstPersonObserver(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Camera returns the value of property "XRView.camera".
//
// The returned bool will be false if there is no such property.
func (this XRView) Camera() (XRCamera, bool) {
	var _ok bool
	_ret := bindings.GetXRViewCamera(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRCamera{}.FromRef(_ret), _ok
}

// RequestViewportScale calls the method "XRView.requestViewportScale".
//
// The returned bool will be false if there is no such method.
func (this XRView) RequestViewportScale(scale float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXRViewRequestViewportScale(
		this.Ref(), js.Pointer(&_ok),
		float64(scale),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RequestViewportScaleFunc returns the method "XRView.requestViewportScale".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRView) RequestViewportScaleFunc() (fn js.Func[func(scale float64)]) {
	return fn.FromRef(
		bindings.XRViewRequestViewportScaleFunc(
			this.Ref(),
		),
	)
}

func NewXRWebGLLayer(session XRSession, context XRWebGLRenderingContext, layerInit XRWebGLLayerInit) XRWebGLLayer {
	return XRWebGLLayer{}.FromRef(
		bindings.NewXRWebGLLayerByXRWebGLLayer(
			session.Ref(),
			context.Ref(),
			js.Pointer(&layerInit)),
	)
}

func NewXRWebGLLayerByXRWebGLLayer1(session XRSession, context XRWebGLRenderingContext) XRWebGLLayer {
	return XRWebGLLayer{}.FromRef(
		bindings.NewXRWebGLLayerByXRWebGLLayer1(
			session.Ref(),
			context.Ref()),
	)
}

type XRWebGLLayer struct {
	XRLayer
}

func (this XRWebGLLayer) Once() XRWebGLLayer {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Antialias returns the value of property "XRWebGLLayer.antialias".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLLayer) Antialias() (bool, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLLayerAntialias(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// IgnoreDepthValues returns the value of property "XRWebGLLayer.ignoreDepthValues".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLLayer) IgnoreDepthValues() (bool, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLLayerIgnoreDepthValues(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// FixedFoveation returns the value of property "XRWebGLLayer.fixedFoveation".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLLayer) FixedFoveation() (float32, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLLayerFixedFoveation(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// FixedFoveation sets the value of property "XRWebGLLayer.fixedFoveation" to val.
//
// It returns false if the property cannot be set.
func (this XRWebGLLayer) SetFixedFoveation(val float32) bool {
	return js.True == bindings.SetXRWebGLLayerFixedFoveation(
		this.Ref(),
		float32(val),
	)
}

// Framebuffer returns the value of property "XRWebGLLayer.framebuffer".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLLayer) Framebuffer() (WebGLFramebuffer, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLLayerFramebuffer(
		this.Ref(), js.Pointer(&_ok),
	)
	return WebGLFramebuffer{}.FromRef(_ret), _ok
}

// FramebufferWidth returns the value of property "XRWebGLLayer.framebufferWidth".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLLayer) FramebufferWidth() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLLayerFramebufferWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// FramebufferHeight returns the value of property "XRWebGLLayer.framebufferHeight".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLLayer) FramebufferHeight() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLLayerFramebufferHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// GetViewport calls the method "XRWebGLLayer.getViewport".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLLayer) GetViewport(view XRView) (XRViewport, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLLayerGetViewport(
		this.Ref(), js.Pointer(&_ok),
		view.Ref(),
	)

	return XRViewport{}.FromRef(_ret), _ok
}

// GetViewportFunc returns the method "XRWebGLLayer.getViewport".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLLayer) GetViewportFunc() (fn js.Func[func(view XRView) XRViewport]) {
	return fn.FromRef(
		bindings.XRWebGLLayerGetViewportFunc(
			this.Ref(),
		),
	)
}

// GetNativeFramebufferScaleFactor calls the staticmethod "XRWebGLLayer.getNativeFramebufferScaleFactor".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLLayer) GetNativeFramebufferScaleFactor(session XRSession) (float64, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLLayerGetNativeFramebufferScaleFactor(
		this.Ref(), js.Pointer(&_ok),
		session.Ref(),
	)

	return float64(_ret), _ok
}

// GetNativeFramebufferScaleFactorFunc returns the staticmethod "XRWebGLLayer.getNativeFramebufferScaleFactor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLLayer) GetNativeFramebufferScaleFactorFunc() (fn js.Func[func(session XRSession) float64]) {
	return fn.FromRef(
		bindings.XRWebGLLayerGetNativeFramebufferScaleFactorFunc(
			this.Ref(),
		),
	)
}

type XRLayer struct {
	EventTarget
}

func (this XRLayer) Once() XRLayer {
	this.Ref().Once()
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
	this.Ref().Free()
}

type XRRenderStateInit struct {
	// DepthNear is "XRRenderStateInit.depthNear"
	//
	// Optional
	DepthNear float64
	// DepthFar is "XRRenderStateInit.depthFar"
	//
	// Optional
	DepthFar float64
	// InlineVerticalFieldOfView is "XRRenderStateInit.inlineVerticalFieldOfView"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 XRRenderStateInit XRRenderStateInit [// XRRenderStateInit] [0x140006ffcc0 0x140006ffe00 0x140006fff40 0x140007a4b40 0x140007a4be0 0x140006ffd60 0x140006ffea0 0x140007a4000] 0x14000780210 {0 0}} in the application heap.
func (p XRRenderStateInit) New() js.Ref {
	return bindings.XRRenderStateInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRRenderStateInit) UpdateFrom(ref js.Ref) {
	bindings.XRRenderStateInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRRenderStateInit) Update(ref js.Ref) {
	bindings.XRRenderStateInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type XRReferenceSpace struct {
	XRSpace
}

func (this XRReferenceSpace) Once() XRReferenceSpace {
	this.Ref().Once()
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
	this.Ref().Free()
}

// GetOffsetReferenceSpace calls the method "XRReferenceSpace.getOffsetReferenceSpace".
//
// The returned bool will be false if there is no such method.
func (this XRReferenceSpace) GetOffsetReferenceSpace(originOffset XRRigidTransform) (XRReferenceSpace, bool) {
	var _ok bool
	_ret := bindings.CallXRReferenceSpaceGetOffsetReferenceSpace(
		this.Ref(), js.Pointer(&_ok),
		originOffset.Ref(),
	)

	return XRReferenceSpace{}.FromRef(_ret), _ok
}

// GetOffsetReferenceSpaceFunc returns the method "XRReferenceSpace.getOffsetReferenceSpace".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRReferenceSpace) GetOffsetReferenceSpaceFunc() (fn js.Func[func(originOffset XRRigidTransform) XRReferenceSpace]) {
	return fn.FromRef(
		bindings.XRReferenceSpaceGetOffsetReferenceSpaceFunc(
			this.Ref(),
		),
	)
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
		assert.Throw("invalid", "callback", "invocation")
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Views returns the value of property "XRViewerPose.views".
//
// The returned bool will be false if there is no such property.
func (this XRViewerPose) Views() (js.FrozenArray[XRView], bool) {
	var _ok bool
	_ret := bindings.GetXRViewerPoseViews(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[XRView]{}.FromRef(_ret), _ok
}

type XRPose struct {
	ref js.Ref
}

func (this XRPose) Once() XRPose {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Transform returns the value of property "XRPose.transform".
//
// The returned bool will be false if there is no such property.
func (this XRPose) Transform() (XRRigidTransform, bool) {
	var _ok bool
	_ret := bindings.GetXRPoseTransform(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRRigidTransform{}.FromRef(_ret), _ok
}

// LinearVelocity returns the value of property "XRPose.linearVelocity".
//
// The returned bool will be false if there is no such property.
func (this XRPose) LinearVelocity() (DOMPointReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetXRPoseLinearVelocity(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMPointReadOnly{}.FromRef(_ret), _ok
}

// AngularVelocity returns the value of property "XRPose.angularVelocity".
//
// The returned bool will be false if there is no such property.
func (this XRPose) AngularVelocity() (DOMPointReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetXRPoseAngularVelocity(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMPointReadOnly{}.FromRef(_ret), _ok
}

// EmulatedPosition returns the value of property "XRPose.emulatedPosition".
//
// The returned bool will be false if there is no such property.
func (this XRPose) EmulatedPosition() (bool, bool) {
	var _ok bool
	_ret := bindings.GetXRPoseEmulatedPosition(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

type XRSpace struct {
	EventTarget
}

func (this XRSpace) Once() XRSpace {
	this.Ref().Once()
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
	this.Ref().Free()
}

type XRAnchor struct {
	ref js.Ref
}

func (this XRAnchor) Once() XRAnchor {
	this.Ref().Once()
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
	this.Ref().Free()
}

// AnchorSpace returns the value of property "XRAnchor.anchorSpace".
//
// The returned bool will be false if there is no such property.
func (this XRAnchor) AnchorSpace() (XRSpace, bool) {
	var _ok bool
	_ret := bindings.GetXRAnchorAnchorSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRSpace{}.FromRef(_ret), _ok
}

// RequestPersistentHandle calls the method "XRAnchor.requestPersistentHandle".
//
// The returned bool will be false if there is no such method.
func (this XRAnchor) RequestPersistentHandle() (js.Promise[js.String], bool) {
	var _ok bool
	_ret := bindings.CallXRAnchorRequestPersistentHandle(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.String]{}.FromRef(_ret), _ok
}

// RequestPersistentHandleFunc returns the method "XRAnchor.requestPersistentHandle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRAnchor) RequestPersistentHandleFunc() (fn js.Func[func() js.Promise[js.String]]) {
	return fn.FromRef(
		bindings.XRAnchorRequestPersistentHandleFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "XRAnchor.delete".
//
// The returned bool will be false if there is no such method.
func (this XRAnchor) Delete() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXRAnchorDelete(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteFunc returns the method "XRAnchor.delete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRAnchor) DeleteFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.XRAnchorDeleteFunc(
			this.Ref(),
		),
	)
}

type XRLightEstimate struct {
	ref js.Ref
}

func (this XRLightEstimate) Once() XRLightEstimate {
	this.Ref().Once()
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
	this.Ref().Free()
}

// SphericalHarmonicsCoefficients returns the value of property "XRLightEstimate.sphericalHarmonicsCoefficients".
//
// The returned bool will be false if there is no such property.
func (this XRLightEstimate) SphericalHarmonicsCoefficients() (js.TypedArray[float32], bool) {
	var _ok bool
	_ret := bindings.GetXRLightEstimateSphericalHarmonicsCoefficients(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[float32]{}.FromRef(_ret), _ok
}

// PrimaryLightDirection returns the value of property "XRLightEstimate.primaryLightDirection".
//
// The returned bool will be false if there is no such property.
func (this XRLightEstimate) PrimaryLightDirection() (DOMPointReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetXRLightEstimatePrimaryLightDirection(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMPointReadOnly{}.FromRef(_ret), _ok
}

// PrimaryLightIntensity returns the value of property "XRLightEstimate.primaryLightIntensity".
//
// The returned bool will be false if there is no such property.
func (this XRLightEstimate) PrimaryLightIntensity() (DOMPointReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetXRLightEstimatePrimaryLightIntensity(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMPointReadOnly{}.FromRef(_ret), _ok
}

type XRLightProbe struct {
	EventTarget
}

func (this XRLightProbe) Once() XRLightProbe {
	this.Ref().Once()
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
	this.Ref().Free()
}

// ProbeSpace returns the value of property "XRLightProbe.probeSpace".
//
// The returned bool will be false if there is no such property.
func (this XRLightProbe) ProbeSpace() (XRSpace, bool) {
	var _ok bool
	_ret := bindings.GetXRLightProbeProbeSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRSpace{}.FromRef(_ret), _ok
}

type XRCPUDepthInformation struct {
	XRDepthInformation
}

func (this XRCPUDepthInformation) Once() XRCPUDepthInformation {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Data returns the value of property "XRCPUDepthInformation.data".
//
// The returned bool will be false if there is no such property.
func (this XRCPUDepthInformation) Data() (js.ArrayBuffer, bool) {
	var _ok bool
	_ret := bindings.GetXRCPUDepthInformationData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.ArrayBuffer{}.FromRef(_ret), _ok
}

// GetDepthInMeters calls the method "XRCPUDepthInformation.getDepthInMeters".
//
// The returned bool will be false if there is no such method.
func (this XRCPUDepthInformation) GetDepthInMeters(x float32, y float32) (float32, bool) {
	var _ok bool
	_ret := bindings.CallXRCPUDepthInformationGetDepthInMeters(
		this.Ref(), js.Pointer(&_ok),
		float32(x),
		float32(y),
	)

	return float32(_ret), _ok
}

// GetDepthInMetersFunc returns the method "XRCPUDepthInformation.getDepthInMeters".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRCPUDepthInformation) GetDepthInMetersFunc() (fn js.Func[func(x float32, y float32) float32]) {
	return fn.FromRef(
		bindings.XRCPUDepthInformationGetDepthInMetersFunc(
			this.Ref(),
		),
	)
}

type XRJointPose struct {
	XRPose
}

func (this XRJointPose) Once() XRJointPose {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Radius returns the value of property "XRJointPose.radius".
//
// The returned bool will be false if there is no such property.
func (this XRJointPose) Radius() (float32, bool) {
	var _ok bool
	_ret := bindings.GetXRJointPoseRadius(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// JointName returns the value of property "XRJointSpace.jointName".
//
// The returned bool will be false if there is no such property.
func (this XRJointSpace) JointName() (XRHandJoint, bool) {
	var _ok bool
	_ret := bindings.GetXRJointSpaceJointName(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRHandJoint(_ret), _ok
}

type XRHitTestResult struct {
	ref js.Ref
}

func (this XRHitTestResult) Once() XRHitTestResult {
	this.Ref().Once()
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
	this.Ref().Free()
}

// GetPose calls the method "XRHitTestResult.getPose".
//
// The returned bool will be false if there is no such method.
func (this XRHitTestResult) GetPose(baseSpace XRSpace) (XRPose, bool) {
	var _ok bool
	_ret := bindings.CallXRHitTestResultGetPose(
		this.Ref(), js.Pointer(&_ok),
		baseSpace.Ref(),
	)

	return XRPose{}.FromRef(_ret), _ok
}

// GetPoseFunc returns the method "XRHitTestResult.getPose".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRHitTestResult) GetPoseFunc() (fn js.Func[func(baseSpace XRSpace) XRPose]) {
	return fn.FromRef(
		bindings.XRHitTestResultGetPoseFunc(
			this.Ref(),
		),
	)
}

// CreateAnchor calls the method "XRHitTestResult.createAnchor".
//
// The returned bool will be false if there is no such method.
func (this XRHitTestResult) CreateAnchor() (js.Promise[XRAnchor], bool) {
	var _ok bool
	_ret := bindings.CallXRHitTestResultCreateAnchor(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[XRAnchor]{}.FromRef(_ret), _ok
}

// CreateAnchorFunc returns the method "XRHitTestResult.createAnchor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRHitTestResult) CreateAnchorFunc() (fn js.Func[func() js.Promise[XRAnchor]]) {
	return fn.FromRef(
		bindings.XRHitTestResultCreateAnchorFunc(
			this.Ref(),
		),
	)
}
