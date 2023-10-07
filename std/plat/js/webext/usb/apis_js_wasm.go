// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package usb

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/usb/bindings"
)

type CloseDeviceCallbackFunc func(this js.Ref) js.Ref

func (fn CloseDeviceCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CloseDeviceCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CloseDeviceCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *CloseDeviceCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CloseDeviceCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TransferType uint32

const (
	_ TransferType = iota

	TransferType_CONTROL
	TransferType_INTERRUPT
	TransferType_ISOCHRONOUS
	TransferType_BULK
)

func (TransferType) FromRef(str js.Ref) TransferType {
	return TransferType(bindings.ConstOfTransferType(str))
}

func (x TransferType) String() (string, bool) {
	switch x {
	case TransferType_CONTROL:
		return "control", true
	case TransferType_INTERRUPT:
		return "interrupt", true
	case TransferType_ISOCHRONOUS:
		return "isochronous", true
	case TransferType_BULK:
		return "bulk", true
	default:
		return "", false
	}
}

type Direction uint32

const (
	_ Direction = iota

	Direction_IN
	Direction_OUT
)

func (Direction) FromRef(str js.Ref) Direction {
	return Direction(bindings.ConstOfDirection(str))
}

func (x Direction) String() (string, bool) {
	switch x {
	case Direction_IN:
		return "in", true
	case Direction_OUT:
		return "out", true
	default:
		return "", false
	}
}

type SynchronizationType uint32

const (
	_ SynchronizationType = iota

	SynchronizationType_ASYNCHRONOUS
	SynchronizationType_ADAPTIVE
	SynchronizationType_SYNCHRONOUS
)

func (SynchronizationType) FromRef(str js.Ref) SynchronizationType {
	return SynchronizationType(bindings.ConstOfSynchronizationType(str))
}

func (x SynchronizationType) String() (string, bool) {
	switch x {
	case SynchronizationType_ASYNCHRONOUS:
		return "asynchronous", true
	case SynchronizationType_ADAPTIVE:
		return "adaptive", true
	case SynchronizationType_SYNCHRONOUS:
		return "synchronous", true
	default:
		return "", false
	}
}

type UsageType uint32

const (
	_ UsageType = iota

	UsageType_DATA
	UsageType_FEEDBACK
	UsageType_EXPLICIT_FEEDBACK
	UsageType_PERIODIC
	UsageType_NOTIFICATION
)

func (UsageType) FromRef(str js.Ref) UsageType {
	return UsageType(bindings.ConstOfUsageType(str))
}

func (x UsageType) String() (string, bool) {
	switch x {
	case UsageType_DATA:
		return "data", true
	case UsageType_FEEDBACK:
		return "feedback", true
	case UsageType_EXPLICIT_FEEDBACK:
		return "explicitFeedback", true
	case UsageType_PERIODIC:
		return "periodic", true
	case UsageType_NOTIFICATION:
		return "notification", true
	default:
		return "", false
	}
}

type EndpointDescriptor struct {
	// Address is "EndpointDescriptor.address"
	//
	// Optional
	//
	// NOTE: FFI_USE_Address MUST be set to true to make this field effective.
	Address int32
	// Type is "EndpointDescriptor.type"
	//
	// Optional
	Type TransferType
	// Direction is "EndpointDescriptor.direction"
	//
	// Optional
	Direction Direction
	// MaximumPacketSize is "EndpointDescriptor.maximumPacketSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaximumPacketSize MUST be set to true to make this field effective.
	MaximumPacketSize int32
	// Synchronization is "EndpointDescriptor.synchronization"
	//
	// Optional
	Synchronization SynchronizationType
	// Usage is "EndpointDescriptor.usage"
	//
	// Optional
	Usage UsageType
	// PollingInterval is "EndpointDescriptor.pollingInterval"
	//
	// Optional
	//
	// NOTE: FFI_USE_PollingInterval MUST be set to true to make this field effective.
	PollingInterval int32
	// ExtraData is "EndpointDescriptor.extra_data"
	//
	// Optional
	ExtraData js.ArrayBuffer

	FFI_USE_Address           bool // for Address.
	FFI_USE_MaximumPacketSize bool // for MaximumPacketSize.
	FFI_USE_PollingInterval   bool // for PollingInterval.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EndpointDescriptor with all fields set.
func (p EndpointDescriptor) FromRef(ref js.Ref) EndpointDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EndpointDescriptor in the application heap.
func (p EndpointDescriptor) New() js.Ref {
	return bindings.EndpointDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EndpointDescriptor) UpdateFrom(ref js.Ref) {
	bindings.EndpointDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EndpointDescriptor) Update(ref js.Ref) {
	bindings.EndpointDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EndpointDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.ExtraData.Ref(),
	)
	p.ExtraData = p.ExtraData.FromRef(js.Undefined)
}

type InterfaceDescriptor struct {
	// InterfaceNumber is "InterfaceDescriptor.interfaceNumber"
	//
	// Optional
	//
	// NOTE: FFI_USE_InterfaceNumber MUST be set to true to make this field effective.
	InterfaceNumber int32
	// AlternateSetting is "InterfaceDescriptor.alternateSetting"
	//
	// Optional
	//
	// NOTE: FFI_USE_AlternateSetting MUST be set to true to make this field effective.
	AlternateSetting int32
	// InterfaceClass is "InterfaceDescriptor.interfaceClass"
	//
	// Optional
	//
	// NOTE: FFI_USE_InterfaceClass MUST be set to true to make this field effective.
	InterfaceClass int32
	// InterfaceSubclass is "InterfaceDescriptor.interfaceSubclass"
	//
	// Optional
	//
	// NOTE: FFI_USE_InterfaceSubclass MUST be set to true to make this field effective.
	InterfaceSubclass int32
	// InterfaceProtocol is "InterfaceDescriptor.interfaceProtocol"
	//
	// Optional
	//
	// NOTE: FFI_USE_InterfaceProtocol MUST be set to true to make this field effective.
	InterfaceProtocol int32
	// Description is "InterfaceDescriptor.description"
	//
	// Optional
	Description js.String
	// Endpoints is "InterfaceDescriptor.endpoints"
	//
	// Optional
	Endpoints js.Array[EndpointDescriptor]
	// ExtraData is "InterfaceDescriptor.extra_data"
	//
	// Optional
	ExtraData js.ArrayBuffer

	FFI_USE_InterfaceNumber   bool // for InterfaceNumber.
	FFI_USE_AlternateSetting  bool // for AlternateSetting.
	FFI_USE_InterfaceClass    bool // for InterfaceClass.
	FFI_USE_InterfaceSubclass bool // for InterfaceSubclass.
	FFI_USE_InterfaceProtocol bool // for InterfaceProtocol.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InterfaceDescriptor with all fields set.
func (p InterfaceDescriptor) FromRef(ref js.Ref) InterfaceDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InterfaceDescriptor in the application heap.
func (p InterfaceDescriptor) New() js.Ref {
	return bindings.InterfaceDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InterfaceDescriptor) UpdateFrom(ref js.Ref) {
	bindings.InterfaceDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InterfaceDescriptor) Update(ref js.Ref) {
	bindings.InterfaceDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InterfaceDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Description.Ref(),
		p.Endpoints.Ref(),
		p.ExtraData.Ref(),
	)
	p.Description = p.Description.FromRef(js.Undefined)
	p.Endpoints = p.Endpoints.FromRef(js.Undefined)
	p.ExtraData = p.ExtraData.FromRef(js.Undefined)
}

type ConfigDescriptor struct {
	// Active is "ConfigDescriptor.active"
	//
	// Optional
	//
	// NOTE: FFI_USE_Active MUST be set to true to make this field effective.
	Active bool
	// ConfigurationValue is "ConfigDescriptor.configurationValue"
	//
	// Optional
	//
	// NOTE: FFI_USE_ConfigurationValue MUST be set to true to make this field effective.
	ConfigurationValue int32
	// Description is "ConfigDescriptor.description"
	//
	// Optional
	Description js.String
	// SelfPowered is "ConfigDescriptor.selfPowered"
	//
	// Optional
	//
	// NOTE: FFI_USE_SelfPowered MUST be set to true to make this field effective.
	SelfPowered bool
	// RemoteWakeup is "ConfigDescriptor.remoteWakeup"
	//
	// Optional
	//
	// NOTE: FFI_USE_RemoteWakeup MUST be set to true to make this field effective.
	RemoteWakeup bool
	// MaxPower is "ConfigDescriptor.maxPower"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxPower MUST be set to true to make this field effective.
	MaxPower int32
	// Interfaces is "ConfigDescriptor.interfaces"
	//
	// Optional
	Interfaces js.Array[InterfaceDescriptor]
	// ExtraData is "ConfigDescriptor.extra_data"
	//
	// Optional
	ExtraData js.ArrayBuffer

	FFI_USE_Active             bool // for Active.
	FFI_USE_ConfigurationValue bool // for ConfigurationValue.
	FFI_USE_SelfPowered        bool // for SelfPowered.
	FFI_USE_RemoteWakeup       bool // for RemoteWakeup.
	FFI_USE_MaxPower           bool // for MaxPower.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConfigDescriptor with all fields set.
func (p ConfigDescriptor) FromRef(ref js.Ref) ConfigDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConfigDescriptor in the application heap.
func (p ConfigDescriptor) New() js.Ref {
	return bindings.ConfigDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConfigDescriptor) UpdateFrom(ref js.Ref) {
	bindings.ConfigDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConfigDescriptor) Update(ref js.Ref) {
	bindings.ConfigDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConfigDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Description.Ref(),
		p.Interfaces.Ref(),
		p.ExtraData.Ref(),
	)
	p.Description = p.Description.FromRef(js.Undefined)
	p.Interfaces = p.Interfaces.FromRef(js.Undefined)
	p.ExtraData = p.ExtraData.FromRef(js.Undefined)
}

type ConnectionHandle struct {
	// Handle is "ConnectionHandle.handle"
	//
	// Optional
	//
	// NOTE: FFI_USE_Handle MUST be set to true to make this field effective.
	Handle int32
	// VendorId is "ConnectionHandle.vendorId"
	//
	// Optional
	//
	// NOTE: FFI_USE_VendorId MUST be set to true to make this field effective.
	VendorId int32
	// ProductId is "ConnectionHandle.productId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProductId MUST be set to true to make this field effective.
	ProductId int32

	FFI_USE_Handle    bool // for Handle.
	FFI_USE_VendorId  bool // for VendorId.
	FFI_USE_ProductId bool // for ProductId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConnectionHandle with all fields set.
func (p ConnectionHandle) FromRef(ref js.Ref) ConnectionHandle {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConnectionHandle in the application heap.
func (p ConnectionHandle) New() js.Ref {
	return bindings.ConnectionHandleJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConnectionHandle) UpdateFrom(ref js.Ref) {
	bindings.ConnectionHandleJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConnectionHandle) Update(ref js.Ref) {
	bindings.ConnectionHandleJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConnectionHandle) FreeMembers(recursive bool) {
}

type Recipient uint32

const (
	_ Recipient = iota

	Recipient_DEVICE
	Recipient__INTERFACE
	Recipient_ENDPOINT
	Recipient_OTHER
)

func (Recipient) FromRef(str js.Ref) Recipient {
	return Recipient(bindings.ConstOfRecipient(str))
}

func (x Recipient) String() (string, bool) {
	switch x {
	case Recipient_DEVICE:
		return "device", true
	case Recipient__INTERFACE:
		return "_interface", true
	case Recipient_ENDPOINT:
		return "endpoint", true
	case Recipient_OTHER:
		return "other", true
	default:
		return "", false
	}
}

type RequestType uint32

const (
	_ RequestType = iota

	RequestType_STANDARD
	RequestType_CLASS
	RequestType_VENDOR
	RequestType_RESERVED
)

func (RequestType) FromRef(str js.Ref) RequestType {
	return RequestType(bindings.ConstOfRequestType(str))
}

func (x RequestType) String() (string, bool) {
	switch x {
	case RequestType_STANDARD:
		return "standard", true
	case RequestType_CLASS:
		return "class", true
	case RequestType_VENDOR:
		return "vendor", true
	case RequestType_RESERVED:
		return "reserved", true
	default:
		return "", false
	}
}

type ControlTransferInfo struct {
	// Direction is "ControlTransferInfo.direction"
	//
	// Optional
	Direction Direction
	// Recipient is "ControlTransferInfo.recipient"
	//
	// Optional
	Recipient Recipient
	// RequestType is "ControlTransferInfo.requestType"
	//
	// Optional
	RequestType RequestType
	// Request is "ControlTransferInfo.request"
	//
	// Optional
	//
	// NOTE: FFI_USE_Request MUST be set to true to make this field effective.
	Request int32
	// Value is "ControlTransferInfo.value"
	//
	// Optional
	//
	// NOTE: FFI_USE_Value MUST be set to true to make this field effective.
	Value int32
	// Index is "ControlTransferInfo.index"
	//
	// Optional
	//
	// NOTE: FFI_USE_Index MUST be set to true to make this field effective.
	Index int32
	// Length is "ControlTransferInfo.length"
	//
	// Optional
	//
	// NOTE: FFI_USE_Length MUST be set to true to make this field effective.
	Length int32
	// Data is "ControlTransferInfo.data"
	//
	// Optional
	Data js.ArrayBuffer
	// Timeout is "ControlTransferInfo.timeout"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timeout MUST be set to true to make this field effective.
	Timeout int32

	FFI_USE_Request bool // for Request.
	FFI_USE_Value   bool // for Value.
	FFI_USE_Index   bool // for Index.
	FFI_USE_Length  bool // for Length.
	FFI_USE_Timeout bool // for Timeout.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ControlTransferInfo with all fields set.
func (p ControlTransferInfo) FromRef(ref js.Ref) ControlTransferInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ControlTransferInfo in the application heap.
func (p ControlTransferInfo) New() js.Ref {
	return bindings.ControlTransferInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ControlTransferInfo) UpdateFrom(ref js.Ref) {
	bindings.ControlTransferInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ControlTransferInfo) Update(ref js.Ref) {
	bindings.ControlTransferInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ControlTransferInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
}

type Device struct {
	// Device is "Device.device"
	//
	// Optional
	//
	// NOTE: FFI_USE_Device MUST be set to true to make this field effective.
	Device int32
	// VendorId is "Device.vendorId"
	//
	// Optional
	//
	// NOTE: FFI_USE_VendorId MUST be set to true to make this field effective.
	VendorId int32
	// ProductId is "Device.productId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProductId MUST be set to true to make this field effective.
	ProductId int32
	// Version is "Device.version"
	//
	// Optional
	//
	// NOTE: FFI_USE_Version MUST be set to true to make this field effective.
	Version int32
	// ProductName is "Device.productName"
	//
	// Optional
	ProductName js.String
	// ManufacturerName is "Device.manufacturerName"
	//
	// Optional
	ManufacturerName js.String
	// SerialNumber is "Device.serialNumber"
	//
	// Optional
	SerialNumber js.String

	FFI_USE_Device    bool // for Device.
	FFI_USE_VendorId  bool // for VendorId.
	FFI_USE_ProductId bool // for ProductId.
	FFI_USE_Version   bool // for Version.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Device with all fields set.
func (p Device) FromRef(ref js.Ref) Device {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Device in the application heap.
func (p Device) New() js.Ref {
	return bindings.DeviceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Device) UpdateFrom(ref js.Ref) {
	bindings.DeviceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Device) Update(ref js.Ref) {
	bindings.DeviceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Device) FreeMembers(recursive bool) {
	js.Free(
		p.ProductName.Ref(),
		p.ManufacturerName.Ref(),
		p.SerialNumber.Ref(),
	)
	p.ProductName = p.ProductName.FromRef(js.Undefined)
	p.ManufacturerName = p.ManufacturerName.FromRef(js.Undefined)
	p.SerialNumber = p.SerialNumber.FromRef(js.Undefined)
}

type DeviceFilter struct {
	// VendorId is "DeviceFilter.vendorId"
	//
	// Optional
	//
	// NOTE: FFI_USE_VendorId MUST be set to true to make this field effective.
	VendorId int32
	// ProductId is "DeviceFilter.productId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProductId MUST be set to true to make this field effective.
	ProductId int32
	// InterfaceClass is "DeviceFilter.interfaceClass"
	//
	// Optional
	//
	// NOTE: FFI_USE_InterfaceClass MUST be set to true to make this field effective.
	InterfaceClass int32
	// InterfaceSubclass is "DeviceFilter.interfaceSubclass"
	//
	// Optional
	//
	// NOTE: FFI_USE_InterfaceSubclass MUST be set to true to make this field effective.
	InterfaceSubclass int32
	// InterfaceProtocol is "DeviceFilter.interfaceProtocol"
	//
	// Optional
	//
	// NOTE: FFI_USE_InterfaceProtocol MUST be set to true to make this field effective.
	InterfaceProtocol int32

	FFI_USE_VendorId          bool // for VendorId.
	FFI_USE_ProductId         bool // for ProductId.
	FFI_USE_InterfaceClass    bool // for InterfaceClass.
	FFI_USE_InterfaceSubclass bool // for InterfaceSubclass.
	FFI_USE_InterfaceProtocol bool // for InterfaceProtocol.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeviceFilter with all fields set.
func (p DeviceFilter) FromRef(ref js.Ref) DeviceFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeviceFilter in the application heap.
func (p DeviceFilter) New() js.Ref {
	return bindings.DeviceFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeviceFilter) UpdateFrom(ref js.Ref) {
	bindings.DeviceFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeviceFilter) Update(ref js.Ref) {
	bindings.DeviceFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeviceFilter) FreeMembers(recursive bool) {
}

type DevicePromptOptions struct {
	// Multiple is "DevicePromptOptions.multiple"
	//
	// Optional
	//
	// NOTE: FFI_USE_Multiple MUST be set to true to make this field effective.
	Multiple bool
	// Filters is "DevicePromptOptions.filters"
	//
	// Optional
	Filters js.Array[DeviceFilter]

	FFI_USE_Multiple bool // for Multiple.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DevicePromptOptions with all fields set.
func (p DevicePromptOptions) FromRef(ref js.Ref) DevicePromptOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DevicePromptOptions in the application heap.
func (p DevicePromptOptions) New() js.Ref {
	return bindings.DevicePromptOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DevicePromptOptions) UpdateFrom(ref js.Ref) {
	bindings.DevicePromptOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DevicePromptOptions) Update(ref js.Ref) {
	bindings.DevicePromptOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DevicePromptOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Filters.Ref(),
	)
	p.Filters = p.Filters.FromRef(js.Undefined)
}

type EnumerateDevicesAndRequestAccessOptions struct {
	// VendorId is "EnumerateDevicesAndRequestAccessOptions.vendorId"
	//
	// Optional
	//
	// NOTE: FFI_USE_VendorId MUST be set to true to make this field effective.
	VendorId int32
	// ProductId is "EnumerateDevicesAndRequestAccessOptions.productId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProductId MUST be set to true to make this field effective.
	ProductId int32
	// InterfaceId is "EnumerateDevicesAndRequestAccessOptions.interfaceId"
	//
	// Optional
	//
	// NOTE: FFI_USE_InterfaceId MUST be set to true to make this field effective.
	InterfaceId int32

	FFI_USE_VendorId    bool // for VendorId.
	FFI_USE_ProductId   bool // for ProductId.
	FFI_USE_InterfaceId bool // for InterfaceId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EnumerateDevicesAndRequestAccessOptions with all fields set.
func (p EnumerateDevicesAndRequestAccessOptions) FromRef(ref js.Ref) EnumerateDevicesAndRequestAccessOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EnumerateDevicesAndRequestAccessOptions in the application heap.
func (p EnumerateDevicesAndRequestAccessOptions) New() js.Ref {
	return bindings.EnumerateDevicesAndRequestAccessOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EnumerateDevicesAndRequestAccessOptions) UpdateFrom(ref js.Ref) {
	bindings.EnumerateDevicesAndRequestAccessOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EnumerateDevicesAndRequestAccessOptions) Update(ref js.Ref) {
	bindings.EnumerateDevicesAndRequestAccessOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EnumerateDevicesAndRequestAccessOptions) FreeMembers(recursive bool) {
}

type EnumerateDevicesOptions struct {
	// VendorId is "EnumerateDevicesOptions.vendorId"
	//
	// Optional
	//
	// NOTE: FFI_USE_VendorId MUST be set to true to make this field effective.
	VendorId int32
	// ProductId is "EnumerateDevicesOptions.productId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProductId MUST be set to true to make this field effective.
	ProductId int32
	// Filters is "EnumerateDevicesOptions.filters"
	//
	// Optional
	Filters js.Array[DeviceFilter]

	FFI_USE_VendorId  bool // for VendorId.
	FFI_USE_ProductId bool // for ProductId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EnumerateDevicesOptions with all fields set.
func (p EnumerateDevicesOptions) FromRef(ref js.Ref) EnumerateDevicesOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EnumerateDevicesOptions in the application heap.
func (p EnumerateDevicesOptions) New() js.Ref {
	return bindings.EnumerateDevicesOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EnumerateDevicesOptions) UpdateFrom(ref js.Ref) {
	bindings.EnumerateDevicesOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EnumerateDevicesOptions) Update(ref js.Ref) {
	bindings.EnumerateDevicesOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EnumerateDevicesOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Filters.Ref(),
	)
	p.Filters = p.Filters.FromRef(js.Undefined)
}

type FindDevicesCallbackFunc func(this js.Ref, handles js.Array[ConnectionHandle]) js.Ref

func (fn FindDevicesCallbackFunc) Register() js.Func[func(handles js.Array[ConnectionHandle])] {
	return js.RegisterCallback[func(handles js.Array[ConnectionHandle])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn FindDevicesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ConnectionHandle]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FindDevicesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, handles js.Array[ConnectionHandle]) js.Ref
	Arg T
}

func (cb *FindDevicesCallback[T]) Register() js.Func[func(handles js.Array[ConnectionHandle])] {
	return js.RegisterCallback[func(handles js.Array[ConnectionHandle])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *FindDevicesCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Array[ConnectionHandle]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GenericTransferInfo struct {
	// Direction is "GenericTransferInfo.direction"
	//
	// Optional
	Direction Direction
	// Endpoint is "GenericTransferInfo.endpoint"
	//
	// Optional
	//
	// NOTE: FFI_USE_Endpoint MUST be set to true to make this field effective.
	Endpoint int32
	// Length is "GenericTransferInfo.length"
	//
	// Optional
	//
	// NOTE: FFI_USE_Length MUST be set to true to make this field effective.
	Length int32
	// Data is "GenericTransferInfo.data"
	//
	// Optional
	Data js.ArrayBuffer
	// Timeout is "GenericTransferInfo.timeout"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timeout MUST be set to true to make this field effective.
	Timeout int32

	FFI_USE_Endpoint bool // for Endpoint.
	FFI_USE_Length   bool // for Length.
	FFI_USE_Timeout  bool // for Timeout.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GenericTransferInfo with all fields set.
func (p GenericTransferInfo) FromRef(ref js.Ref) GenericTransferInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GenericTransferInfo in the application heap.
func (p GenericTransferInfo) New() js.Ref {
	return bindings.GenericTransferInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GenericTransferInfo) UpdateFrom(ref js.Ref) {
	bindings.GenericTransferInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GenericTransferInfo) Update(ref js.Ref) {
	bindings.GenericTransferInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GenericTransferInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
}

type GetConfigurationCallbackFunc func(this js.Ref, config *ConfigDescriptor) js.Ref

func (fn GetConfigurationCallbackFunc) Register() js.Func[func(config *ConfigDescriptor)] {
	return js.RegisterCallback[func(config *ConfigDescriptor)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetConfigurationCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ConfigDescriptor
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetConfigurationCallback[T any] struct {
	Fn  func(arg T, this js.Ref, config *ConfigDescriptor) js.Ref
	Arg T
}

func (cb *GetConfigurationCallback[T]) Register() js.Func[func(config *ConfigDescriptor)] {
	return js.RegisterCallback[func(config *ConfigDescriptor)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetConfigurationCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ConfigDescriptor
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetConfigurationsCallbackFunc func(this js.Ref, configs js.Array[ConfigDescriptor]) js.Ref

func (fn GetConfigurationsCallbackFunc) Register() js.Func[func(configs js.Array[ConfigDescriptor])] {
	return js.RegisterCallback[func(configs js.Array[ConfigDescriptor])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetConfigurationsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ConfigDescriptor]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetConfigurationsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, configs js.Array[ConfigDescriptor]) js.Ref
	Arg T
}

func (cb *GetConfigurationsCallback[T]) Register() js.Func[func(configs js.Array[ConfigDescriptor])] {
	return js.RegisterCallback[func(configs js.Array[ConfigDescriptor])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetConfigurationsCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Array[ConfigDescriptor]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDevicesCallbackFunc func(this js.Ref, devices js.Array[Device]) js.Ref

func (fn GetDevicesCallbackFunc) Register() js.Func[func(devices js.Array[Device])] {
	return js.RegisterCallback[func(devices js.Array[Device])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDevicesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Device]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDevicesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, devices js.Array[Device]) js.Ref
	Arg T
}

func (cb *GetDevicesCallback[T]) Register() js.Func[func(devices js.Array[Device])] {
	return js.RegisterCallback[func(devices js.Array[Device])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDevicesCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Array[Device]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type IsochronousTransferInfo struct {
	// TransferInfo is "IsochronousTransferInfo.transferInfo"
	//
	// Optional
	//
	// NOTE: TransferInfo.FFI_USE MUST be set to true to get TransferInfo used.
	TransferInfo GenericTransferInfo
	// Packets is "IsochronousTransferInfo.packets"
	//
	// Optional
	//
	// NOTE: FFI_USE_Packets MUST be set to true to make this field effective.
	Packets int32
	// PacketLength is "IsochronousTransferInfo.packetLength"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketLength MUST be set to true to make this field effective.
	PacketLength int32

	FFI_USE_Packets      bool // for Packets.
	FFI_USE_PacketLength bool // for PacketLength.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IsochronousTransferInfo with all fields set.
func (p IsochronousTransferInfo) FromRef(ref js.Ref) IsochronousTransferInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IsochronousTransferInfo in the application heap.
func (p IsochronousTransferInfo) New() js.Ref {
	return bindings.IsochronousTransferInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IsochronousTransferInfo) UpdateFrom(ref js.Ref) {
	bindings.IsochronousTransferInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IsochronousTransferInfo) Update(ref js.Ref) {
	bindings.IsochronousTransferInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IsochronousTransferInfo) FreeMembers(recursive bool) {
	if recursive {
		p.TransferInfo.FreeMembers(true)
	}
}

type ListInterfacesCallbackFunc func(this js.Ref, descriptors js.Array[InterfaceDescriptor]) js.Ref

func (fn ListInterfacesCallbackFunc) Register() js.Func[func(descriptors js.Array[InterfaceDescriptor])] {
	return js.RegisterCallback[func(descriptors js.Array[InterfaceDescriptor])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ListInterfacesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[InterfaceDescriptor]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ListInterfacesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, descriptors js.Array[InterfaceDescriptor]) js.Ref
	Arg T
}

func (cb *ListInterfacesCallback[T]) Register() js.Func[func(descriptors js.Array[InterfaceDescriptor])] {
	return js.RegisterCallback[func(descriptors js.Array[InterfaceDescriptor])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ListInterfacesCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Array[InterfaceDescriptor]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OpenDeviceCallbackFunc func(this js.Ref, handle *ConnectionHandle) js.Ref

func (fn OpenDeviceCallbackFunc) Register() js.Func[func(handle *ConnectionHandle)] {
	return js.RegisterCallback[func(handle *ConnectionHandle)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OpenDeviceCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ConnectionHandle
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OpenDeviceCallback[T any] struct {
	Fn  func(arg T, this js.Ref, handle *ConnectionHandle) js.Ref
	Arg T
}

func (cb *OpenDeviceCallback[T]) Register() js.Func[func(handle *ConnectionHandle)] {
	return js.RegisterCallback[func(handle *ConnectionHandle)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OpenDeviceCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ConnectionHandle
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RequestAccessCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn RequestAccessCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RequestAccessCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RequestAccessCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *RequestAccessCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RequestAccessCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ResetDeviceCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn ResetDeviceCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ResetDeviceCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ResetDeviceCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *ResetDeviceCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ResetDeviceCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TransferCallbackFunc func(this js.Ref, info *TransferResultInfo) js.Ref

func (fn TransferCallbackFunc) Register() js.Func[func(info *TransferResultInfo)] {
	return js.RegisterCallback[func(info *TransferResultInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn TransferCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TransferResultInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TransferCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *TransferResultInfo) js.Ref
	Arg T
}

func (cb *TransferCallback[T]) Register() js.Func[func(info *TransferResultInfo)] {
	return js.RegisterCallback[func(info *TransferResultInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *TransferCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TransferResultInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TransferResultInfo struct {
	// ResultCode is "TransferResultInfo.resultCode"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResultCode MUST be set to true to make this field effective.
	ResultCode int32
	// Data is "TransferResultInfo.data"
	//
	// Optional
	Data js.ArrayBuffer

	FFI_USE_ResultCode bool // for ResultCode.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TransferResultInfo with all fields set.
func (p TransferResultInfo) FromRef(ref js.Ref) TransferResultInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TransferResultInfo in the application heap.
func (p TransferResultInfo) New() js.Ref {
	return bindings.TransferResultInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TransferResultInfo) UpdateFrom(ref js.Ref) {
	bindings.TransferResultInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TransferResultInfo) Update(ref js.Ref) {
	bindings.TransferResultInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TransferResultInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
}

type VoidCallbackFunc func(this js.Ref) js.Ref

func (fn VoidCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VoidCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type VoidCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *VoidCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VoidCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncBulkTransfer returns true if the function "WEBEXT.usb.bulkTransfer" exists.
func HasFuncBulkTransfer() bool {
	return js.True == bindings.HasFuncBulkTransfer()
}

// FuncBulkTransfer returns the function "WEBEXT.usb.bulkTransfer".
func FuncBulkTransfer() (fn js.Func[func(handle ConnectionHandle, transferInfo GenericTransferInfo) js.Promise[TransferResultInfo]]) {
	bindings.FuncBulkTransfer(
		js.Pointer(&fn),
	)
	return
}

// BulkTransfer calls the function "WEBEXT.usb.bulkTransfer" directly.
func BulkTransfer(handle ConnectionHandle, transferInfo GenericTransferInfo) (ret js.Promise[TransferResultInfo]) {
	bindings.CallBulkTransfer(
		js.Pointer(&ret),
		js.Pointer(&handle),
		js.Pointer(&transferInfo),
	)

	return
}

// TryBulkTransfer calls the function "WEBEXT.usb.bulkTransfer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryBulkTransfer(handle ConnectionHandle, transferInfo GenericTransferInfo) (ret js.Promise[TransferResultInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBulkTransfer(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&handle),
		js.Pointer(&transferInfo),
	)

	return
}

// HasFuncClaimInterface returns true if the function "WEBEXT.usb.claimInterface" exists.
func HasFuncClaimInterface() bool {
	return js.True == bindings.HasFuncClaimInterface()
}

// FuncClaimInterface returns the function "WEBEXT.usb.claimInterface".
func FuncClaimInterface() (fn js.Func[func(handle ConnectionHandle, interfaceNumber int32) js.Promise[js.Void]]) {
	bindings.FuncClaimInterface(
		js.Pointer(&fn),
	)
	return
}

// ClaimInterface calls the function "WEBEXT.usb.claimInterface" directly.
func ClaimInterface(handle ConnectionHandle, interfaceNumber int32) (ret js.Promise[js.Void]) {
	bindings.CallClaimInterface(
		js.Pointer(&ret),
		js.Pointer(&handle),
		int32(interfaceNumber),
	)

	return
}

// TryClaimInterface calls the function "WEBEXT.usb.claimInterface"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClaimInterface(handle ConnectionHandle, interfaceNumber int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClaimInterface(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&handle),
		int32(interfaceNumber),
	)

	return
}

// HasFuncCloseDevice returns true if the function "WEBEXT.usb.closeDevice" exists.
func HasFuncCloseDevice() bool {
	return js.True == bindings.HasFuncCloseDevice()
}

// FuncCloseDevice returns the function "WEBEXT.usb.closeDevice".
func FuncCloseDevice() (fn js.Func[func(handle ConnectionHandle) js.Promise[js.Void]]) {
	bindings.FuncCloseDevice(
		js.Pointer(&fn),
	)
	return
}

// CloseDevice calls the function "WEBEXT.usb.closeDevice" directly.
func CloseDevice(handle ConnectionHandle) (ret js.Promise[js.Void]) {
	bindings.CallCloseDevice(
		js.Pointer(&ret),
		js.Pointer(&handle),
	)

	return
}

// TryCloseDevice calls the function "WEBEXT.usb.closeDevice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCloseDevice(handle ConnectionHandle) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCloseDevice(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&handle),
	)

	return
}

// HasFuncControlTransfer returns true if the function "WEBEXT.usb.controlTransfer" exists.
func HasFuncControlTransfer() bool {
	return js.True == bindings.HasFuncControlTransfer()
}

// FuncControlTransfer returns the function "WEBEXT.usb.controlTransfer".
func FuncControlTransfer() (fn js.Func[func(handle ConnectionHandle, transferInfo ControlTransferInfo) js.Promise[TransferResultInfo]]) {
	bindings.FuncControlTransfer(
		js.Pointer(&fn),
	)
	return
}

// ControlTransfer calls the function "WEBEXT.usb.controlTransfer" directly.
func ControlTransfer(handle ConnectionHandle, transferInfo ControlTransferInfo) (ret js.Promise[TransferResultInfo]) {
	bindings.CallControlTransfer(
		js.Pointer(&ret),
		js.Pointer(&handle),
		js.Pointer(&transferInfo),
	)

	return
}

// TryControlTransfer calls the function "WEBEXT.usb.controlTransfer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryControlTransfer(handle ConnectionHandle, transferInfo ControlTransferInfo) (ret js.Promise[TransferResultInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryControlTransfer(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&handle),
		js.Pointer(&transferInfo),
	)

	return
}

// HasFuncFindDevices returns true if the function "WEBEXT.usb.findDevices" exists.
func HasFuncFindDevices() bool {
	return js.True == bindings.HasFuncFindDevices()
}

// FuncFindDevices returns the function "WEBEXT.usb.findDevices".
func FuncFindDevices() (fn js.Func[func(options EnumerateDevicesAndRequestAccessOptions) js.Promise[js.Array[ConnectionHandle]]]) {
	bindings.FuncFindDevices(
		js.Pointer(&fn),
	)
	return
}

// FindDevices calls the function "WEBEXT.usb.findDevices" directly.
func FindDevices(options EnumerateDevicesAndRequestAccessOptions) (ret js.Promise[js.Array[ConnectionHandle]]) {
	bindings.CallFindDevices(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryFindDevices calls the function "WEBEXT.usb.findDevices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryFindDevices(options EnumerateDevicesAndRequestAccessOptions) (ret js.Promise[js.Array[ConnectionHandle]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFindDevices(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetConfiguration returns true if the function "WEBEXT.usb.getConfiguration" exists.
func HasFuncGetConfiguration() bool {
	return js.True == bindings.HasFuncGetConfiguration()
}

// FuncGetConfiguration returns the function "WEBEXT.usb.getConfiguration".
func FuncGetConfiguration() (fn js.Func[func(handle ConnectionHandle) js.Promise[ConfigDescriptor]]) {
	bindings.FuncGetConfiguration(
		js.Pointer(&fn),
	)
	return
}

// GetConfiguration calls the function "WEBEXT.usb.getConfiguration" directly.
func GetConfiguration(handle ConnectionHandle) (ret js.Promise[ConfigDescriptor]) {
	bindings.CallGetConfiguration(
		js.Pointer(&ret),
		js.Pointer(&handle),
	)

	return
}

// TryGetConfiguration calls the function "WEBEXT.usb.getConfiguration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetConfiguration(handle ConnectionHandle) (ret js.Promise[ConfigDescriptor], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetConfiguration(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&handle),
	)

	return
}

// HasFuncGetConfigurations returns true if the function "WEBEXT.usb.getConfigurations" exists.
func HasFuncGetConfigurations() bool {
	return js.True == bindings.HasFuncGetConfigurations()
}

// FuncGetConfigurations returns the function "WEBEXT.usb.getConfigurations".
func FuncGetConfigurations() (fn js.Func[func(device Device) js.Promise[js.Array[ConfigDescriptor]]]) {
	bindings.FuncGetConfigurations(
		js.Pointer(&fn),
	)
	return
}

// GetConfigurations calls the function "WEBEXT.usb.getConfigurations" directly.
func GetConfigurations(device Device) (ret js.Promise[js.Array[ConfigDescriptor]]) {
	bindings.CallGetConfigurations(
		js.Pointer(&ret),
		js.Pointer(&device),
	)

	return
}

// TryGetConfigurations calls the function "WEBEXT.usb.getConfigurations"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetConfigurations(device Device) (ret js.Promise[js.Array[ConfigDescriptor]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetConfigurations(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&device),
	)

	return
}

// HasFuncGetDevices returns true if the function "WEBEXT.usb.getDevices" exists.
func HasFuncGetDevices() bool {
	return js.True == bindings.HasFuncGetDevices()
}

// FuncGetDevices returns the function "WEBEXT.usb.getDevices".
func FuncGetDevices() (fn js.Func[func(options EnumerateDevicesOptions) js.Promise[js.Array[Device]]]) {
	bindings.FuncGetDevices(
		js.Pointer(&fn),
	)
	return
}

// GetDevices calls the function "WEBEXT.usb.getDevices" directly.
func GetDevices(options EnumerateDevicesOptions) (ret js.Promise[js.Array[Device]]) {
	bindings.CallGetDevices(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetDevices calls the function "WEBEXT.usb.getDevices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDevices(options EnumerateDevicesOptions) (ret js.Promise[js.Array[Device]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDevices(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetUserSelectedDevices returns true if the function "WEBEXT.usb.getUserSelectedDevices" exists.
func HasFuncGetUserSelectedDevices() bool {
	return js.True == bindings.HasFuncGetUserSelectedDevices()
}

// FuncGetUserSelectedDevices returns the function "WEBEXT.usb.getUserSelectedDevices".
func FuncGetUserSelectedDevices() (fn js.Func[func(options DevicePromptOptions) js.Promise[js.Array[Device]]]) {
	bindings.FuncGetUserSelectedDevices(
		js.Pointer(&fn),
	)
	return
}

// GetUserSelectedDevices calls the function "WEBEXT.usb.getUserSelectedDevices" directly.
func GetUserSelectedDevices(options DevicePromptOptions) (ret js.Promise[js.Array[Device]]) {
	bindings.CallGetUserSelectedDevices(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetUserSelectedDevices calls the function "WEBEXT.usb.getUserSelectedDevices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetUserSelectedDevices(options DevicePromptOptions) (ret js.Promise[js.Array[Device]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetUserSelectedDevices(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncInterruptTransfer returns true if the function "WEBEXT.usb.interruptTransfer" exists.
func HasFuncInterruptTransfer() bool {
	return js.True == bindings.HasFuncInterruptTransfer()
}

// FuncInterruptTransfer returns the function "WEBEXT.usb.interruptTransfer".
func FuncInterruptTransfer() (fn js.Func[func(handle ConnectionHandle, transferInfo GenericTransferInfo) js.Promise[TransferResultInfo]]) {
	bindings.FuncInterruptTransfer(
		js.Pointer(&fn),
	)
	return
}

// InterruptTransfer calls the function "WEBEXT.usb.interruptTransfer" directly.
func InterruptTransfer(handle ConnectionHandle, transferInfo GenericTransferInfo) (ret js.Promise[TransferResultInfo]) {
	bindings.CallInterruptTransfer(
		js.Pointer(&ret),
		js.Pointer(&handle),
		js.Pointer(&transferInfo),
	)

	return
}

// TryInterruptTransfer calls the function "WEBEXT.usb.interruptTransfer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInterruptTransfer(handle ConnectionHandle, transferInfo GenericTransferInfo) (ret js.Promise[TransferResultInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInterruptTransfer(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&handle),
		js.Pointer(&transferInfo),
	)

	return
}

// HasFuncIsochronousTransfer returns true if the function "WEBEXT.usb.isochronousTransfer" exists.
func HasFuncIsochronousTransfer() bool {
	return js.True == bindings.HasFuncIsochronousTransfer()
}

// FuncIsochronousTransfer returns the function "WEBEXT.usb.isochronousTransfer".
func FuncIsochronousTransfer() (fn js.Func[func(handle ConnectionHandle, transferInfo IsochronousTransferInfo) js.Promise[TransferResultInfo]]) {
	bindings.FuncIsochronousTransfer(
		js.Pointer(&fn),
	)
	return
}

// IsochronousTransfer calls the function "WEBEXT.usb.isochronousTransfer" directly.
func IsochronousTransfer(handle ConnectionHandle, transferInfo IsochronousTransferInfo) (ret js.Promise[TransferResultInfo]) {
	bindings.CallIsochronousTransfer(
		js.Pointer(&ret),
		js.Pointer(&handle),
		js.Pointer(&transferInfo),
	)

	return
}

// TryIsochronousTransfer calls the function "WEBEXT.usb.isochronousTransfer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsochronousTransfer(handle ConnectionHandle, transferInfo IsochronousTransferInfo) (ret js.Promise[TransferResultInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsochronousTransfer(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&handle),
		js.Pointer(&transferInfo),
	)

	return
}

// HasFuncListInterfaces returns true if the function "WEBEXT.usb.listInterfaces" exists.
func HasFuncListInterfaces() bool {
	return js.True == bindings.HasFuncListInterfaces()
}

// FuncListInterfaces returns the function "WEBEXT.usb.listInterfaces".
func FuncListInterfaces() (fn js.Func[func(handle ConnectionHandle) js.Promise[js.Array[InterfaceDescriptor]]]) {
	bindings.FuncListInterfaces(
		js.Pointer(&fn),
	)
	return
}

// ListInterfaces calls the function "WEBEXT.usb.listInterfaces" directly.
func ListInterfaces(handle ConnectionHandle) (ret js.Promise[js.Array[InterfaceDescriptor]]) {
	bindings.CallListInterfaces(
		js.Pointer(&ret),
		js.Pointer(&handle),
	)

	return
}

// TryListInterfaces calls the function "WEBEXT.usb.listInterfaces"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryListInterfaces(handle ConnectionHandle) (ret js.Promise[js.Array[InterfaceDescriptor]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryListInterfaces(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&handle),
	)

	return
}

type OnDeviceAddedEventCallbackFunc func(this js.Ref, device *Device) js.Ref

func (fn OnDeviceAddedEventCallbackFunc) Register() js.Func[func(device *Device)] {
	return js.RegisterCallback[func(device *Device)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeviceAddedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Device
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDeviceAddedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, device *Device) js.Ref
	Arg T
}

func (cb *OnDeviceAddedEventCallback[T]) Register() js.Func[func(device *Device)] {
	return js.RegisterCallback[func(device *Device)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeviceAddedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Device
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDeviceAdded returns true if the function "WEBEXT.usb.onDeviceAdded.addListener" exists.
func HasFuncOnDeviceAdded() bool {
	return js.True == bindings.HasFuncOnDeviceAdded()
}

// FuncOnDeviceAdded returns the function "WEBEXT.usb.onDeviceAdded.addListener".
func FuncOnDeviceAdded() (fn js.Func[func(callback js.Func[func(device *Device)])]) {
	bindings.FuncOnDeviceAdded(
		js.Pointer(&fn),
	)
	return
}

// OnDeviceAdded calls the function "WEBEXT.usb.onDeviceAdded.addListener" directly.
func OnDeviceAdded(callback js.Func[func(device *Device)]) (ret js.Void) {
	bindings.CallOnDeviceAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeviceAdded calls the function "WEBEXT.usb.onDeviceAdded.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeviceAdded(callback js.Func[func(device *Device)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeviceAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeviceAdded returns true if the function "WEBEXT.usb.onDeviceAdded.removeListener" exists.
func HasFuncOffDeviceAdded() bool {
	return js.True == bindings.HasFuncOffDeviceAdded()
}

// FuncOffDeviceAdded returns the function "WEBEXT.usb.onDeviceAdded.removeListener".
func FuncOffDeviceAdded() (fn js.Func[func(callback js.Func[func(device *Device)])]) {
	bindings.FuncOffDeviceAdded(
		js.Pointer(&fn),
	)
	return
}

// OffDeviceAdded calls the function "WEBEXT.usb.onDeviceAdded.removeListener" directly.
func OffDeviceAdded(callback js.Func[func(device *Device)]) (ret js.Void) {
	bindings.CallOffDeviceAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeviceAdded calls the function "WEBEXT.usb.onDeviceAdded.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeviceAdded(callback js.Func[func(device *Device)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeviceAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeviceAdded returns true if the function "WEBEXT.usb.onDeviceAdded.hasListener" exists.
func HasFuncHasOnDeviceAdded() bool {
	return js.True == bindings.HasFuncHasOnDeviceAdded()
}

// FuncHasOnDeviceAdded returns the function "WEBEXT.usb.onDeviceAdded.hasListener".
func FuncHasOnDeviceAdded() (fn js.Func[func(callback js.Func[func(device *Device)]) bool]) {
	bindings.FuncHasOnDeviceAdded(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeviceAdded calls the function "WEBEXT.usb.onDeviceAdded.hasListener" directly.
func HasOnDeviceAdded(callback js.Func[func(device *Device)]) (ret bool) {
	bindings.CallHasOnDeviceAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeviceAdded calls the function "WEBEXT.usb.onDeviceAdded.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeviceAdded(callback js.Func[func(device *Device)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeviceAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDeviceRemovedEventCallbackFunc func(this js.Ref, device *Device) js.Ref

func (fn OnDeviceRemovedEventCallbackFunc) Register() js.Func[func(device *Device)] {
	return js.RegisterCallback[func(device *Device)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeviceRemovedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Device
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDeviceRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, device *Device) js.Ref
	Arg T
}

func (cb *OnDeviceRemovedEventCallback[T]) Register() js.Func[func(device *Device)] {
	return js.RegisterCallback[func(device *Device)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeviceRemovedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Device
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDeviceRemoved returns true if the function "WEBEXT.usb.onDeviceRemoved.addListener" exists.
func HasFuncOnDeviceRemoved() bool {
	return js.True == bindings.HasFuncOnDeviceRemoved()
}

// FuncOnDeviceRemoved returns the function "WEBEXT.usb.onDeviceRemoved.addListener".
func FuncOnDeviceRemoved() (fn js.Func[func(callback js.Func[func(device *Device)])]) {
	bindings.FuncOnDeviceRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnDeviceRemoved calls the function "WEBEXT.usb.onDeviceRemoved.addListener" directly.
func OnDeviceRemoved(callback js.Func[func(device *Device)]) (ret js.Void) {
	bindings.CallOnDeviceRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeviceRemoved calls the function "WEBEXT.usb.onDeviceRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeviceRemoved(callback js.Func[func(device *Device)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeviceRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeviceRemoved returns true if the function "WEBEXT.usb.onDeviceRemoved.removeListener" exists.
func HasFuncOffDeviceRemoved() bool {
	return js.True == bindings.HasFuncOffDeviceRemoved()
}

// FuncOffDeviceRemoved returns the function "WEBEXT.usb.onDeviceRemoved.removeListener".
func FuncOffDeviceRemoved() (fn js.Func[func(callback js.Func[func(device *Device)])]) {
	bindings.FuncOffDeviceRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffDeviceRemoved calls the function "WEBEXT.usb.onDeviceRemoved.removeListener" directly.
func OffDeviceRemoved(callback js.Func[func(device *Device)]) (ret js.Void) {
	bindings.CallOffDeviceRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeviceRemoved calls the function "WEBEXT.usb.onDeviceRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeviceRemoved(callback js.Func[func(device *Device)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeviceRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeviceRemoved returns true if the function "WEBEXT.usb.onDeviceRemoved.hasListener" exists.
func HasFuncHasOnDeviceRemoved() bool {
	return js.True == bindings.HasFuncHasOnDeviceRemoved()
}

// FuncHasOnDeviceRemoved returns the function "WEBEXT.usb.onDeviceRemoved.hasListener".
func FuncHasOnDeviceRemoved() (fn js.Func[func(callback js.Func[func(device *Device)]) bool]) {
	bindings.FuncHasOnDeviceRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeviceRemoved calls the function "WEBEXT.usb.onDeviceRemoved.hasListener" directly.
func HasOnDeviceRemoved(callback js.Func[func(device *Device)]) (ret bool) {
	bindings.CallHasOnDeviceRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeviceRemoved calls the function "WEBEXT.usb.onDeviceRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeviceRemoved(callback js.Func[func(device *Device)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeviceRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOpenDevice returns true if the function "WEBEXT.usb.openDevice" exists.
func HasFuncOpenDevice() bool {
	return js.True == bindings.HasFuncOpenDevice()
}

// FuncOpenDevice returns the function "WEBEXT.usb.openDevice".
func FuncOpenDevice() (fn js.Func[func(device Device) js.Promise[ConnectionHandle]]) {
	bindings.FuncOpenDevice(
		js.Pointer(&fn),
	)
	return
}

// OpenDevice calls the function "WEBEXT.usb.openDevice" directly.
func OpenDevice(device Device) (ret js.Promise[ConnectionHandle]) {
	bindings.CallOpenDevice(
		js.Pointer(&ret),
		js.Pointer(&device),
	)

	return
}

// TryOpenDevice calls the function "WEBEXT.usb.openDevice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenDevice(device Device) (ret js.Promise[ConnectionHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenDevice(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&device),
	)

	return
}

// HasFuncReleaseInterface returns true if the function "WEBEXT.usb.releaseInterface" exists.
func HasFuncReleaseInterface() bool {
	return js.True == bindings.HasFuncReleaseInterface()
}

// FuncReleaseInterface returns the function "WEBEXT.usb.releaseInterface".
func FuncReleaseInterface() (fn js.Func[func(handle ConnectionHandle, interfaceNumber int32) js.Promise[js.Void]]) {
	bindings.FuncReleaseInterface(
		js.Pointer(&fn),
	)
	return
}

// ReleaseInterface calls the function "WEBEXT.usb.releaseInterface" directly.
func ReleaseInterface(handle ConnectionHandle, interfaceNumber int32) (ret js.Promise[js.Void]) {
	bindings.CallReleaseInterface(
		js.Pointer(&ret),
		js.Pointer(&handle),
		int32(interfaceNumber),
	)

	return
}

// TryReleaseInterface calls the function "WEBEXT.usb.releaseInterface"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReleaseInterface(handle ConnectionHandle, interfaceNumber int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReleaseInterface(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&handle),
		int32(interfaceNumber),
	)

	return
}

// HasFuncRequestAccess returns true if the function "WEBEXT.usb.requestAccess" exists.
func HasFuncRequestAccess() bool {
	return js.True == bindings.HasFuncRequestAccess()
}

// FuncRequestAccess returns the function "WEBEXT.usb.requestAccess".
func FuncRequestAccess() (fn js.Func[func(device Device, interfaceId int32) js.Promise[js.Boolean]]) {
	bindings.FuncRequestAccess(
		js.Pointer(&fn),
	)
	return
}

// RequestAccess calls the function "WEBEXT.usb.requestAccess" directly.
func RequestAccess(device Device, interfaceId int32) (ret js.Promise[js.Boolean]) {
	bindings.CallRequestAccess(
		js.Pointer(&ret),
		js.Pointer(&device),
		int32(interfaceId),
	)

	return
}

// TryRequestAccess calls the function "WEBEXT.usb.requestAccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRequestAccess(device Device, interfaceId int32) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestAccess(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&device),
		int32(interfaceId),
	)

	return
}

// HasFuncResetDevice returns true if the function "WEBEXT.usb.resetDevice" exists.
func HasFuncResetDevice() bool {
	return js.True == bindings.HasFuncResetDevice()
}

// FuncResetDevice returns the function "WEBEXT.usb.resetDevice".
func FuncResetDevice() (fn js.Func[func(handle ConnectionHandle) js.Promise[js.Boolean]]) {
	bindings.FuncResetDevice(
		js.Pointer(&fn),
	)
	return
}

// ResetDevice calls the function "WEBEXT.usb.resetDevice" directly.
func ResetDevice(handle ConnectionHandle) (ret js.Promise[js.Boolean]) {
	bindings.CallResetDevice(
		js.Pointer(&ret),
		js.Pointer(&handle),
	)

	return
}

// TryResetDevice calls the function "WEBEXT.usb.resetDevice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryResetDevice(handle ConnectionHandle) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResetDevice(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&handle),
	)

	return
}

// HasFuncSetConfiguration returns true if the function "WEBEXT.usb.setConfiguration" exists.
func HasFuncSetConfiguration() bool {
	return js.True == bindings.HasFuncSetConfiguration()
}

// FuncSetConfiguration returns the function "WEBEXT.usb.setConfiguration".
func FuncSetConfiguration() (fn js.Func[func(handle ConnectionHandle, configurationValue int32) js.Promise[js.Void]]) {
	bindings.FuncSetConfiguration(
		js.Pointer(&fn),
	)
	return
}

// SetConfiguration calls the function "WEBEXT.usb.setConfiguration" directly.
func SetConfiguration(handle ConnectionHandle, configurationValue int32) (ret js.Promise[js.Void]) {
	bindings.CallSetConfiguration(
		js.Pointer(&ret),
		js.Pointer(&handle),
		int32(configurationValue),
	)

	return
}

// TrySetConfiguration calls the function "WEBEXT.usb.setConfiguration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetConfiguration(handle ConnectionHandle, configurationValue int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetConfiguration(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&handle),
		int32(configurationValue),
	)

	return
}

// HasFuncSetInterfaceAlternateSetting returns true if the function "WEBEXT.usb.setInterfaceAlternateSetting" exists.
func HasFuncSetInterfaceAlternateSetting() bool {
	return js.True == bindings.HasFuncSetInterfaceAlternateSetting()
}

// FuncSetInterfaceAlternateSetting returns the function "WEBEXT.usb.setInterfaceAlternateSetting".
func FuncSetInterfaceAlternateSetting() (fn js.Func[func(handle ConnectionHandle, interfaceNumber int32, alternateSetting int32) js.Promise[js.Void]]) {
	bindings.FuncSetInterfaceAlternateSetting(
		js.Pointer(&fn),
	)
	return
}

// SetInterfaceAlternateSetting calls the function "WEBEXT.usb.setInterfaceAlternateSetting" directly.
func SetInterfaceAlternateSetting(handle ConnectionHandle, interfaceNumber int32, alternateSetting int32) (ret js.Promise[js.Void]) {
	bindings.CallSetInterfaceAlternateSetting(
		js.Pointer(&ret),
		js.Pointer(&handle),
		int32(interfaceNumber),
		int32(alternateSetting),
	)

	return
}

// TrySetInterfaceAlternateSetting calls the function "WEBEXT.usb.setInterfaceAlternateSetting"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetInterfaceAlternateSetting(handle ConnectionHandle, interfaceNumber int32, alternateSetting int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetInterfaceAlternateSetting(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&handle),
		int32(interfaceNumber),
		int32(alternateSetting),
	)

	return
}
