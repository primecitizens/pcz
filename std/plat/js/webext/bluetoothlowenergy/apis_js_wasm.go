// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package bluetoothlowenergy

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/bluetoothlowenergy/bindings"
)

type AdvertisementType uint32

const (
	_ AdvertisementType = iota

	AdvertisementType_BROADCAST
	AdvertisementType_PERIPHERAL
)

func (AdvertisementType) FromRef(str js.Ref) AdvertisementType {
	return AdvertisementType(bindings.ConstOfAdvertisementType(str))
}

func (x AdvertisementType) String() (string, bool) {
	switch x {
	case AdvertisementType_BROADCAST:
		return "broadcast", true
	case AdvertisementType_PERIPHERAL:
		return "peripheral", true
	default:
		return "", false
	}
}

type ManufacturerData struct {
	// Id is "ManufacturerData.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// Data is "ManufacturerData.data"
	//
	// Optional
	Data js.Array[int32]

	FFI_USE_Id bool // for Id.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManufacturerData with all fields set.
func (p ManufacturerData) FromRef(ref js.Ref) ManufacturerData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManufacturerData in the application heap.
func (p ManufacturerData) New() js.Ref {
	return bindings.ManufacturerDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManufacturerData) UpdateFrom(ref js.Ref) {
	bindings.ManufacturerDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManufacturerData) Update(ref js.Ref) {
	bindings.ManufacturerDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManufacturerData) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
}

type ServiceData struct {
	// Uuid is "ServiceData.uuid"
	//
	// Optional
	Uuid js.String
	// Data is "ServiceData.data"
	//
	// Optional
	Data js.Array[int32]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ServiceData with all fields set.
func (p ServiceData) FromRef(ref js.Ref) ServiceData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ServiceData in the application heap.
func (p ServiceData) New() js.Ref {
	return bindings.ServiceDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ServiceData) UpdateFrom(ref js.Ref) {
	bindings.ServiceDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ServiceData) Update(ref js.Ref) {
	bindings.ServiceDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ServiceData) FreeMembers(recursive bool) {
	js.Free(
		p.Uuid.Ref(),
		p.Data.Ref(),
	)
	p.Uuid = p.Uuid.FromRef(js.Undefined)
	p.Data = p.Data.FromRef(js.Undefined)
}

type Advertisement struct {
	// Type is "Advertisement.type"
	//
	// Optional
	Type AdvertisementType
	// ServiceUuids is "Advertisement.serviceUuids"
	//
	// Optional
	ServiceUuids js.Array[js.String]
	// ManufacturerData is "Advertisement.manufacturerData"
	//
	// Optional
	ManufacturerData js.Array[ManufacturerData]
	// SolicitUuids is "Advertisement.solicitUuids"
	//
	// Optional
	SolicitUuids js.Array[js.String]
	// ServiceData is "Advertisement.serviceData"
	//
	// Optional
	ServiceData js.Array[ServiceData]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Advertisement with all fields set.
func (p Advertisement) FromRef(ref js.Ref) Advertisement {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Advertisement in the application heap.
func (p Advertisement) New() js.Ref {
	return bindings.AdvertisementJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Advertisement) UpdateFrom(ref js.Ref) {
	bindings.AdvertisementJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Advertisement) Update(ref js.Ref) {
	bindings.AdvertisementJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Advertisement) FreeMembers(recursive bool) {
	js.Free(
		p.ServiceUuids.Ref(),
		p.ManufacturerData.Ref(),
		p.SolicitUuids.Ref(),
		p.ServiceData.Ref(),
	)
	p.ServiceUuids = p.ServiceUuids.FromRef(js.Undefined)
	p.ManufacturerData = p.ManufacturerData.FromRef(js.Undefined)
	p.SolicitUuids = p.SolicitUuids.FromRef(js.Undefined)
	p.ServiceData = p.ServiceData.FromRef(js.Undefined)
}

type Service struct {
	// Uuid is "Service.uuid"
	//
	// Optional
	Uuid js.String
	// IsPrimary is "Service.isPrimary"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsPrimary MUST be set to true to make this field effective.
	IsPrimary bool
	// InstanceId is "Service.instanceId"
	//
	// Optional
	InstanceId js.String
	// DeviceAddress is "Service.deviceAddress"
	//
	// Optional
	DeviceAddress js.String

	FFI_USE_IsPrimary bool // for IsPrimary.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Service with all fields set.
func (p Service) FromRef(ref js.Ref) Service {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Service in the application heap.
func (p Service) New() js.Ref {
	return bindings.ServiceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Service) UpdateFrom(ref js.Ref) {
	bindings.ServiceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Service) Update(ref js.Ref) {
	bindings.ServiceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Service) FreeMembers(recursive bool) {
	js.Free(
		p.Uuid.Ref(),
		p.InstanceId.Ref(),
		p.DeviceAddress.Ref(),
	)
	p.Uuid = p.Uuid.FromRef(js.Undefined)
	p.InstanceId = p.InstanceId.FromRef(js.Undefined)
	p.DeviceAddress = p.DeviceAddress.FromRef(js.Undefined)
}

type CharacteristicProperty uint32

const (
	_ CharacteristicProperty = iota

	CharacteristicProperty_BROADCAST
	CharacteristicProperty_READ
	CharacteristicProperty_WRITE_WITHOUT_RESPONSE
	CharacteristicProperty_WRITE
	CharacteristicProperty_NOTIFY
	CharacteristicProperty_INDICATE
	CharacteristicProperty_AUTHENTICATED_SIGNED_WRITES
	CharacteristicProperty_EXTENDED_PROPERTIES
	CharacteristicProperty_RELIABLE_WRITE
	CharacteristicProperty_WRITABLE_AUXILIARIES
	CharacteristicProperty_ENCRYPT_READ
	CharacteristicProperty_ENCRYPT_WRITE
	CharacteristicProperty_ENCRYPT_AUTHENTICATED_READ
	CharacteristicProperty_ENCRYPT_AUTHENTICATED_WRITE
)

func (CharacteristicProperty) FromRef(str js.Ref) CharacteristicProperty {
	return CharacteristicProperty(bindings.ConstOfCharacteristicProperty(str))
}

func (x CharacteristicProperty) String() (string, bool) {
	switch x {
	case CharacteristicProperty_BROADCAST:
		return "broadcast", true
	case CharacteristicProperty_READ:
		return "read", true
	case CharacteristicProperty_WRITE_WITHOUT_RESPONSE:
		return "writeWithoutResponse", true
	case CharacteristicProperty_WRITE:
		return "write", true
	case CharacteristicProperty_NOTIFY:
		return "notify", true
	case CharacteristicProperty_INDICATE:
		return "indicate", true
	case CharacteristicProperty_AUTHENTICATED_SIGNED_WRITES:
		return "authenticatedSignedWrites", true
	case CharacteristicProperty_EXTENDED_PROPERTIES:
		return "extendedProperties", true
	case CharacteristicProperty_RELIABLE_WRITE:
		return "reliableWrite", true
	case CharacteristicProperty_WRITABLE_AUXILIARIES:
		return "writableAuxiliaries", true
	case CharacteristicProperty_ENCRYPT_READ:
		return "encryptRead", true
	case CharacteristicProperty_ENCRYPT_WRITE:
		return "encryptWrite", true
	case CharacteristicProperty_ENCRYPT_AUTHENTICATED_READ:
		return "encryptAuthenticatedRead", true
	case CharacteristicProperty_ENCRYPT_AUTHENTICATED_WRITE:
		return "encryptAuthenticatedWrite", true
	default:
		return "", false
	}
}

type Characteristic struct {
	// Uuid is "Characteristic.uuid"
	//
	// Optional
	Uuid js.String
	// Service is "Characteristic.service"
	//
	// Optional
	//
	// NOTE: Service.FFI_USE MUST be set to true to get Service used.
	Service Service
	// Properties is "Characteristic.properties"
	//
	// Optional
	Properties js.Array[CharacteristicProperty]
	// InstanceId is "Characteristic.instanceId"
	//
	// Optional
	InstanceId js.String
	// Value is "Characteristic.value"
	//
	// Optional
	Value js.ArrayBuffer

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Characteristic with all fields set.
func (p Characteristic) FromRef(ref js.Ref) Characteristic {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Characteristic in the application heap.
func (p Characteristic) New() js.Ref {
	return bindings.CharacteristicJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Characteristic) UpdateFrom(ref js.Ref) {
	bindings.CharacteristicJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Characteristic) Update(ref js.Ref) {
	bindings.CharacteristicJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Characteristic) FreeMembers(recursive bool) {
	js.Free(
		p.Uuid.Ref(),
		p.Properties.Ref(),
		p.InstanceId.Ref(),
		p.Value.Ref(),
	)
	p.Uuid = p.Uuid.FromRef(js.Undefined)
	p.Properties = p.Properties.FromRef(js.Undefined)
	p.InstanceId = p.InstanceId.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
	if recursive {
		p.Service.FreeMembers(true)
	}
}

type CharacteristicCallbackFunc func(this js.Ref, result *Characteristic) js.Ref

func (fn CharacteristicCallbackFunc) Register() js.Func[func(result *Characteristic)] {
	return js.RegisterCallback[func(result *Characteristic)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CharacteristicCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Characteristic
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

type CharacteristicCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *Characteristic) js.Ref
	Arg T
}

func (cb *CharacteristicCallback[T]) Register() js.Func[func(result *Characteristic)] {
	return js.RegisterCallback[func(result *Characteristic)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CharacteristicCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Characteristic
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

type CharacteristicsCallbackFunc func(this js.Ref, result js.Array[Characteristic]) js.Ref

func (fn CharacteristicsCallbackFunc) Register() js.Func[func(result js.Array[Characteristic])] {
	return js.RegisterCallback[func(result js.Array[Characteristic])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CharacteristicsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Characteristic]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CharacteristicsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.Array[Characteristic]) js.Ref
	Arg T
}

func (cb *CharacteristicsCallback[T]) Register() js.Func[func(result js.Array[Characteristic])] {
	return js.RegisterCallback[func(result js.Array[Characteristic])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CharacteristicsCallback[T]) DispatchCallback(
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

		js.Array[Characteristic]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ConnectProperties struct {
	// Persistent is "ConnectProperties.persistent"
	//
	// Optional
	//
	// NOTE: FFI_USE_Persistent MUST be set to true to make this field effective.
	Persistent bool

	FFI_USE_Persistent bool // for Persistent.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConnectProperties with all fields set.
func (p ConnectProperties) FromRef(ref js.Ref) ConnectProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConnectProperties in the application heap.
func (p ConnectProperties) New() js.Ref {
	return bindings.ConnectPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConnectProperties) UpdateFrom(ref js.Ref) {
	bindings.ConnectPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConnectProperties) Update(ref js.Ref) {
	bindings.ConnectPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConnectProperties) FreeMembers(recursive bool) {
}

type CreateCharacteristicCallbackFunc func(this js.Ref, characteristicId js.String) js.Ref

func (fn CreateCharacteristicCallbackFunc) Register() js.Func[func(characteristicId js.String)] {
	return js.RegisterCallback[func(characteristicId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CreateCharacteristicCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateCharacteristicCallback[T any] struct {
	Fn  func(arg T, this js.Ref, characteristicId js.String) js.Ref
	Arg T
}

func (cb *CreateCharacteristicCallback[T]) Register() js.Func[func(characteristicId js.String)] {
	return js.RegisterCallback[func(characteristicId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CreateCharacteristicCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateDescriptorCallbackFunc func(this js.Ref, descriptorId js.String) js.Ref

func (fn CreateDescriptorCallbackFunc) Register() js.Func[func(descriptorId js.String)] {
	return js.RegisterCallback[func(descriptorId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CreateDescriptorCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateDescriptorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, descriptorId js.String) js.Ref
	Arg T
}

func (cb *CreateDescriptorCallback[T]) Register() js.Func[func(descriptorId js.String)] {
	return js.RegisterCallback[func(descriptorId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CreateDescriptorCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateServiceCallbackFunc func(this js.Ref, serviceId js.String) js.Ref

func (fn CreateServiceCallbackFunc) Register() js.Func[func(serviceId js.String)] {
	return js.RegisterCallback[func(serviceId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CreateServiceCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateServiceCallback[T any] struct {
	Fn  func(arg T, this js.Ref, serviceId js.String) js.Ref
	Arg T
}

func (cb *CreateServiceCallback[T]) Register() js.Func[func(serviceId js.String)] {
	return js.RegisterCallback[func(serviceId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CreateServiceCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DescriptorPermission uint32

const (
	_ DescriptorPermission = iota

	DescriptorPermission_READ
	DescriptorPermission_WRITE
	DescriptorPermission_ENCRYPTED_READ
	DescriptorPermission_ENCRYPTED_WRITE
	DescriptorPermission_ENCRYPTED_AUTHENTICATED_READ
	DescriptorPermission_ENCRYPTED_AUTHENTICATED_WRITE
)

func (DescriptorPermission) FromRef(str js.Ref) DescriptorPermission {
	return DescriptorPermission(bindings.ConstOfDescriptorPermission(str))
}

func (x DescriptorPermission) String() (string, bool) {
	switch x {
	case DescriptorPermission_READ:
		return "read", true
	case DescriptorPermission_WRITE:
		return "write", true
	case DescriptorPermission_ENCRYPTED_READ:
		return "encryptedRead", true
	case DescriptorPermission_ENCRYPTED_WRITE:
		return "encryptedWrite", true
	case DescriptorPermission_ENCRYPTED_AUTHENTICATED_READ:
		return "encryptedAuthenticatedRead", true
	case DescriptorPermission_ENCRYPTED_AUTHENTICATED_WRITE:
		return "encryptedAuthenticatedWrite", true
	default:
		return "", false
	}
}

type Descriptor struct {
	// Uuid is "Descriptor.uuid"
	//
	// Optional
	Uuid js.String
	// Characteristic is "Descriptor.characteristic"
	//
	// Optional
	//
	// NOTE: Characteristic.FFI_USE MUST be set to true to get Characteristic used.
	Characteristic Characteristic
	// Permissions is "Descriptor.permissions"
	//
	// Optional
	Permissions js.Array[DescriptorPermission]
	// InstanceId is "Descriptor.instanceId"
	//
	// Optional
	InstanceId js.String
	// Value is "Descriptor.value"
	//
	// Optional
	Value js.ArrayBuffer

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Descriptor with all fields set.
func (p Descriptor) FromRef(ref js.Ref) Descriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Descriptor in the application heap.
func (p Descriptor) New() js.Ref {
	return bindings.DescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Descriptor) UpdateFrom(ref js.Ref) {
	bindings.DescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Descriptor) Update(ref js.Ref) {
	bindings.DescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Descriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Uuid.Ref(),
		p.Permissions.Ref(),
		p.InstanceId.Ref(),
		p.Value.Ref(),
	)
	p.Uuid = p.Uuid.FromRef(js.Undefined)
	p.Permissions = p.Permissions.FromRef(js.Undefined)
	p.InstanceId = p.InstanceId.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
	if recursive {
		p.Characteristic.FreeMembers(true)
	}
}

type DescriptorCallbackFunc func(this js.Ref, result *Descriptor) js.Ref

func (fn DescriptorCallbackFunc) Register() js.Func[func(result *Descriptor)] {
	return js.RegisterCallback[func(result *Descriptor)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DescriptorCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Descriptor
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

type DescriptorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *Descriptor) js.Ref
	Arg T
}

func (cb *DescriptorCallback[T]) Register() js.Func[func(result *Descriptor)] {
	return js.RegisterCallback[func(result *Descriptor)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DescriptorCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Descriptor
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

type DescriptorsCallbackFunc func(this js.Ref, result js.Array[Descriptor]) js.Ref

func (fn DescriptorsCallbackFunc) Register() js.Func[func(result js.Array[Descriptor])] {
	return js.RegisterCallback[func(result js.Array[Descriptor])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DescriptorsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Descriptor]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DescriptorsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.Array[Descriptor]) js.Ref
	Arg T
}

func (cb *DescriptorsCallback[T]) Register() js.Func[func(result js.Array[Descriptor])] {
	return js.RegisterCallback[func(result js.Array[Descriptor])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DescriptorsCallback[T]) DispatchCallback(
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

		js.Array[Descriptor]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Device struct {
	// Address is "Device.address"
	//
	// Optional
	Address js.String
	// Name is "Device.name"
	//
	// Optional
	Name js.String
	// DeviceClass is "Device.deviceClass"
	//
	// Optional
	//
	// NOTE: FFI_USE_DeviceClass MUST be set to true to make this field effective.
	DeviceClass int32

	FFI_USE_DeviceClass bool // for DeviceClass.

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
		p.Address.Ref(),
		p.Name.Ref(),
	)
	p.Address = p.Address.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type Notification struct {
	// Value is "Notification.value"
	//
	// Optional
	Value js.ArrayBuffer
	// ShouldIndicate is "Notification.shouldIndicate"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShouldIndicate MUST be set to true to make this field effective.
	ShouldIndicate bool

	FFI_USE_ShouldIndicate bool // for ShouldIndicate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Notification with all fields set.
func (p Notification) FromRef(ref js.Ref) Notification {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Notification in the application heap.
func (p Notification) New() js.Ref {
	return bindings.NotificationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Notification) UpdateFrom(ref js.Ref) {
	bindings.NotificationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Notification) Update(ref js.Ref) {
	bindings.NotificationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Notification) FreeMembers(recursive bool) {
	js.Free(
		p.Value.Ref(),
	)
	p.Value = p.Value.FromRef(js.Undefined)
}

type NotificationProperties struct {
	// Persistent is "NotificationProperties.persistent"
	//
	// Optional
	//
	// NOTE: FFI_USE_Persistent MUST be set to true to make this field effective.
	Persistent bool

	FFI_USE_Persistent bool // for Persistent.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NotificationProperties with all fields set.
func (p NotificationProperties) FromRef(ref js.Ref) NotificationProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NotificationProperties in the application heap.
func (p NotificationProperties) New() js.Ref {
	return bindings.NotificationPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NotificationProperties) UpdateFrom(ref js.Ref) {
	bindings.NotificationPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NotificationProperties) Update(ref js.Ref) {
	bindings.NotificationPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NotificationProperties) FreeMembers(recursive bool) {
}

type RegisterAdvertisementCallbackFunc func(this js.Ref, advertisementId int32) js.Ref

func (fn RegisterAdvertisementCallbackFunc) Register() js.Func[func(advertisementId int32)] {
	return js.RegisterCallback[func(advertisementId int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RegisterAdvertisementCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RegisterAdvertisementCallback[T any] struct {
	Fn  func(arg T, this js.Ref, advertisementId int32) js.Ref
	Arg T
}

func (cb *RegisterAdvertisementCallback[T]) Register() js.Func[func(advertisementId int32)] {
	return js.RegisterCallback[func(advertisementId int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RegisterAdvertisementCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Request struct {
	// RequestId is "Request.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// Device is "Request.device"
	//
	// Optional
	//
	// NOTE: Device.FFI_USE MUST be set to true to get Device used.
	Device Device
	// Value is "Request.value"
	//
	// Optional
	Value js.ArrayBuffer

	FFI_USE_RequestId bool // for RequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Request with all fields set.
func (p Request) FromRef(ref js.Ref) Request {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Request in the application heap.
func (p Request) New() js.Ref {
	return bindings.RequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Request) UpdateFrom(ref js.Ref) {
	bindings.RequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Request) Update(ref js.Ref) {
	bindings.RequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Request) FreeMembers(recursive bool) {
	js.Free(
		p.Value.Ref(),
	)
	p.Value = p.Value.FromRef(js.Undefined)
	if recursive {
		p.Device.FreeMembers(true)
	}
}

type Response struct {
	// RequestId is "Response.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// IsError is "Response.isError"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsError MUST be set to true to make this field effective.
	IsError bool
	// Value is "Response.value"
	//
	// Optional
	Value js.ArrayBuffer

	FFI_USE_RequestId bool // for RequestId.
	FFI_USE_IsError   bool // for IsError.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Response with all fields set.
func (p Response) FromRef(ref js.Ref) Response {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Response in the application heap.
func (p Response) New() js.Ref {
	return bindings.ResponseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Response) UpdateFrom(ref js.Ref) {
	bindings.ResponseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Response) Update(ref js.Ref) {
	bindings.ResponseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Response) FreeMembers(recursive bool) {
	js.Free(
		p.Value.Ref(),
	)
	p.Value = p.Value.FromRef(js.Undefined)
}

type ResultCallbackFunc func(this js.Ref) js.Ref

func (fn ResultCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ResultCallbackFunc) DispatchCallback(
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

type ResultCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ResultCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ResultCallback[T]) DispatchCallback(
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

type ServiceCallbackFunc func(this js.Ref, result *Service) js.Ref

func (fn ServiceCallbackFunc) Register() js.Func[func(result *Service)] {
	return js.RegisterCallback[func(result *Service)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ServiceCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Service
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

type ServiceCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *Service) js.Ref
	Arg T
}

func (cb *ServiceCallback[T]) Register() js.Func[func(result *Service)] {
	return js.RegisterCallback[func(result *Service)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ServiceCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Service
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

type ServicesCallbackFunc func(this js.Ref, result js.Array[Service]) js.Ref

func (fn ServicesCallbackFunc) Register() js.Func[func(result js.Array[Service])] {
	return js.RegisterCallback[func(result js.Array[Service])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ServicesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Service]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ServicesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.Array[Service]) js.Ref
	Arg T
}

func (cb *ServicesCallback[T]) Register() js.Func[func(result js.Array[Service])] {
	return js.RegisterCallback[func(result js.Array[Service])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ServicesCallback[T]) DispatchCallback(
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

		js.Array[Service]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncConnect returns true if the function "WEBEXT.bluetoothLowEnergy.connect" exists.
func HasFuncConnect() bool {
	return js.True == bindings.HasFuncConnect()
}

// FuncConnect returns the function "WEBEXT.bluetoothLowEnergy.connect".
func FuncConnect() (fn js.Func[func(deviceAddress js.String, properties ConnectProperties) js.Promise[js.Void]]) {
	bindings.FuncConnect(
		js.Pointer(&fn),
	)
	return
}

// Connect calls the function "WEBEXT.bluetoothLowEnergy.connect" directly.
func Connect(deviceAddress js.String, properties ConnectProperties) (ret js.Promise[js.Void]) {
	bindings.CallConnect(
		js.Pointer(&ret),
		deviceAddress.Ref(),
		js.Pointer(&properties),
	)

	return
}

// TryConnect calls the function "WEBEXT.bluetoothLowEnergy.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryConnect(deviceAddress js.String, properties ConnectProperties) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		deviceAddress.Ref(),
		js.Pointer(&properties),
	)

	return
}

// HasFuncCreateCharacteristic returns true if the function "WEBEXT.bluetoothLowEnergy.createCharacteristic" exists.
func HasFuncCreateCharacteristic() bool {
	return js.True == bindings.HasFuncCreateCharacteristic()
}

// FuncCreateCharacteristic returns the function "WEBEXT.bluetoothLowEnergy.createCharacteristic".
func FuncCreateCharacteristic() (fn js.Func[func(characteristic Characteristic, serviceId js.String) js.Promise[js.String]]) {
	bindings.FuncCreateCharacteristic(
		js.Pointer(&fn),
	)
	return
}

// CreateCharacteristic calls the function "WEBEXT.bluetoothLowEnergy.createCharacteristic" directly.
func CreateCharacteristic(characteristic Characteristic, serviceId js.String) (ret js.Promise[js.String]) {
	bindings.CallCreateCharacteristic(
		js.Pointer(&ret),
		js.Pointer(&characteristic),
		serviceId.Ref(),
	)

	return
}

// TryCreateCharacteristic calls the function "WEBEXT.bluetoothLowEnergy.createCharacteristic"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreateCharacteristic(characteristic Characteristic, serviceId js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreateCharacteristic(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&characteristic),
		serviceId.Ref(),
	)

	return
}

// HasFuncCreateDescriptor returns true if the function "WEBEXT.bluetoothLowEnergy.createDescriptor" exists.
func HasFuncCreateDescriptor() bool {
	return js.True == bindings.HasFuncCreateDescriptor()
}

// FuncCreateDescriptor returns the function "WEBEXT.bluetoothLowEnergy.createDescriptor".
func FuncCreateDescriptor() (fn js.Func[func(descriptor Descriptor, characteristicId js.String) js.Promise[js.String]]) {
	bindings.FuncCreateDescriptor(
		js.Pointer(&fn),
	)
	return
}

// CreateDescriptor calls the function "WEBEXT.bluetoothLowEnergy.createDescriptor" directly.
func CreateDescriptor(descriptor Descriptor, characteristicId js.String) (ret js.Promise[js.String]) {
	bindings.CallCreateDescriptor(
		js.Pointer(&ret),
		js.Pointer(&descriptor),
		characteristicId.Ref(),
	)

	return
}

// TryCreateDescriptor calls the function "WEBEXT.bluetoothLowEnergy.createDescriptor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreateDescriptor(descriptor Descriptor, characteristicId js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreateDescriptor(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
		characteristicId.Ref(),
	)

	return
}

// HasFuncCreateService returns true if the function "WEBEXT.bluetoothLowEnergy.createService" exists.
func HasFuncCreateService() bool {
	return js.True == bindings.HasFuncCreateService()
}

// FuncCreateService returns the function "WEBEXT.bluetoothLowEnergy.createService".
func FuncCreateService() (fn js.Func[func(service Service) js.Promise[js.String]]) {
	bindings.FuncCreateService(
		js.Pointer(&fn),
	)
	return
}

// CreateService calls the function "WEBEXT.bluetoothLowEnergy.createService" directly.
func CreateService(service Service) (ret js.Promise[js.String]) {
	bindings.CallCreateService(
		js.Pointer(&ret),
		js.Pointer(&service),
	)

	return
}

// TryCreateService calls the function "WEBEXT.bluetoothLowEnergy.createService"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreateService(service Service) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreateService(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&service),
	)

	return
}

// HasFuncDisconnect returns true if the function "WEBEXT.bluetoothLowEnergy.disconnect" exists.
func HasFuncDisconnect() bool {
	return js.True == bindings.HasFuncDisconnect()
}

// FuncDisconnect returns the function "WEBEXT.bluetoothLowEnergy.disconnect".
func FuncDisconnect() (fn js.Func[func(deviceAddress js.String) js.Promise[js.Void]]) {
	bindings.FuncDisconnect(
		js.Pointer(&fn),
	)
	return
}

// Disconnect calls the function "WEBEXT.bluetoothLowEnergy.disconnect" directly.
func Disconnect(deviceAddress js.String) (ret js.Promise[js.Void]) {
	bindings.CallDisconnect(
		js.Pointer(&ret),
		deviceAddress.Ref(),
	)

	return
}

// TryDisconnect calls the function "WEBEXT.bluetoothLowEnergy.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDisconnect(deviceAddress js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDisconnect(
		js.Pointer(&ret), js.Pointer(&exception),
		deviceAddress.Ref(),
	)

	return
}

// HasFuncGetCharacteristic returns true if the function "WEBEXT.bluetoothLowEnergy.getCharacteristic" exists.
func HasFuncGetCharacteristic() bool {
	return js.True == bindings.HasFuncGetCharacteristic()
}

// FuncGetCharacteristic returns the function "WEBEXT.bluetoothLowEnergy.getCharacteristic".
func FuncGetCharacteristic() (fn js.Func[func(characteristicId js.String) js.Promise[Characteristic]]) {
	bindings.FuncGetCharacteristic(
		js.Pointer(&fn),
	)
	return
}

// GetCharacteristic calls the function "WEBEXT.bluetoothLowEnergy.getCharacteristic" directly.
func GetCharacteristic(characteristicId js.String) (ret js.Promise[Characteristic]) {
	bindings.CallGetCharacteristic(
		js.Pointer(&ret),
		characteristicId.Ref(),
	)

	return
}

// TryGetCharacteristic calls the function "WEBEXT.bluetoothLowEnergy.getCharacteristic"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCharacteristic(characteristicId js.String) (ret js.Promise[Characteristic], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCharacteristic(
		js.Pointer(&ret), js.Pointer(&exception),
		characteristicId.Ref(),
	)

	return
}

// HasFuncGetCharacteristics returns true if the function "WEBEXT.bluetoothLowEnergy.getCharacteristics" exists.
func HasFuncGetCharacteristics() bool {
	return js.True == bindings.HasFuncGetCharacteristics()
}

// FuncGetCharacteristics returns the function "WEBEXT.bluetoothLowEnergy.getCharacteristics".
func FuncGetCharacteristics() (fn js.Func[func(serviceId js.String) js.Promise[js.Array[Characteristic]]]) {
	bindings.FuncGetCharacteristics(
		js.Pointer(&fn),
	)
	return
}

// GetCharacteristics calls the function "WEBEXT.bluetoothLowEnergy.getCharacteristics" directly.
func GetCharacteristics(serviceId js.String) (ret js.Promise[js.Array[Characteristic]]) {
	bindings.CallGetCharacteristics(
		js.Pointer(&ret),
		serviceId.Ref(),
	)

	return
}

// TryGetCharacteristics calls the function "WEBEXT.bluetoothLowEnergy.getCharacteristics"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCharacteristics(serviceId js.String) (ret js.Promise[js.Array[Characteristic]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCharacteristics(
		js.Pointer(&ret), js.Pointer(&exception),
		serviceId.Ref(),
	)

	return
}

// HasFuncGetDescriptor returns true if the function "WEBEXT.bluetoothLowEnergy.getDescriptor" exists.
func HasFuncGetDescriptor() bool {
	return js.True == bindings.HasFuncGetDescriptor()
}

// FuncGetDescriptor returns the function "WEBEXT.bluetoothLowEnergy.getDescriptor".
func FuncGetDescriptor() (fn js.Func[func(descriptorId js.String) js.Promise[Descriptor]]) {
	bindings.FuncGetDescriptor(
		js.Pointer(&fn),
	)
	return
}

// GetDescriptor calls the function "WEBEXT.bluetoothLowEnergy.getDescriptor" directly.
func GetDescriptor(descriptorId js.String) (ret js.Promise[Descriptor]) {
	bindings.CallGetDescriptor(
		js.Pointer(&ret),
		descriptorId.Ref(),
	)

	return
}

// TryGetDescriptor calls the function "WEBEXT.bluetoothLowEnergy.getDescriptor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDescriptor(descriptorId js.String) (ret js.Promise[Descriptor], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDescriptor(
		js.Pointer(&ret), js.Pointer(&exception),
		descriptorId.Ref(),
	)

	return
}

// HasFuncGetDescriptors returns true if the function "WEBEXT.bluetoothLowEnergy.getDescriptors" exists.
func HasFuncGetDescriptors() bool {
	return js.True == bindings.HasFuncGetDescriptors()
}

// FuncGetDescriptors returns the function "WEBEXT.bluetoothLowEnergy.getDescriptors".
func FuncGetDescriptors() (fn js.Func[func(characteristicId js.String) js.Promise[js.Array[Descriptor]]]) {
	bindings.FuncGetDescriptors(
		js.Pointer(&fn),
	)
	return
}

// GetDescriptors calls the function "WEBEXT.bluetoothLowEnergy.getDescriptors" directly.
func GetDescriptors(characteristicId js.String) (ret js.Promise[js.Array[Descriptor]]) {
	bindings.CallGetDescriptors(
		js.Pointer(&ret),
		characteristicId.Ref(),
	)

	return
}

// TryGetDescriptors calls the function "WEBEXT.bluetoothLowEnergy.getDescriptors"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDescriptors(characteristicId js.String) (ret js.Promise[js.Array[Descriptor]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDescriptors(
		js.Pointer(&ret), js.Pointer(&exception),
		characteristicId.Ref(),
	)

	return
}

// HasFuncGetIncludedServices returns true if the function "WEBEXT.bluetoothLowEnergy.getIncludedServices" exists.
func HasFuncGetIncludedServices() bool {
	return js.True == bindings.HasFuncGetIncludedServices()
}

// FuncGetIncludedServices returns the function "WEBEXT.bluetoothLowEnergy.getIncludedServices".
func FuncGetIncludedServices() (fn js.Func[func(serviceId js.String) js.Promise[js.Array[Service]]]) {
	bindings.FuncGetIncludedServices(
		js.Pointer(&fn),
	)
	return
}

// GetIncludedServices calls the function "WEBEXT.bluetoothLowEnergy.getIncludedServices" directly.
func GetIncludedServices(serviceId js.String) (ret js.Promise[js.Array[Service]]) {
	bindings.CallGetIncludedServices(
		js.Pointer(&ret),
		serviceId.Ref(),
	)

	return
}

// TryGetIncludedServices calls the function "WEBEXT.bluetoothLowEnergy.getIncludedServices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetIncludedServices(serviceId js.String) (ret js.Promise[js.Array[Service]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetIncludedServices(
		js.Pointer(&ret), js.Pointer(&exception),
		serviceId.Ref(),
	)

	return
}

// HasFuncGetService returns true if the function "WEBEXT.bluetoothLowEnergy.getService" exists.
func HasFuncGetService() bool {
	return js.True == bindings.HasFuncGetService()
}

// FuncGetService returns the function "WEBEXT.bluetoothLowEnergy.getService".
func FuncGetService() (fn js.Func[func(serviceId js.String) js.Promise[Service]]) {
	bindings.FuncGetService(
		js.Pointer(&fn),
	)
	return
}

// GetService calls the function "WEBEXT.bluetoothLowEnergy.getService" directly.
func GetService(serviceId js.String) (ret js.Promise[Service]) {
	bindings.CallGetService(
		js.Pointer(&ret),
		serviceId.Ref(),
	)

	return
}

// TryGetService calls the function "WEBEXT.bluetoothLowEnergy.getService"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetService(serviceId js.String) (ret js.Promise[Service], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetService(
		js.Pointer(&ret), js.Pointer(&exception),
		serviceId.Ref(),
	)

	return
}

// HasFuncGetServices returns true if the function "WEBEXT.bluetoothLowEnergy.getServices" exists.
func HasFuncGetServices() bool {
	return js.True == bindings.HasFuncGetServices()
}

// FuncGetServices returns the function "WEBEXT.bluetoothLowEnergy.getServices".
func FuncGetServices() (fn js.Func[func(deviceAddress js.String) js.Promise[js.Array[Service]]]) {
	bindings.FuncGetServices(
		js.Pointer(&fn),
	)
	return
}

// GetServices calls the function "WEBEXT.bluetoothLowEnergy.getServices" directly.
func GetServices(deviceAddress js.String) (ret js.Promise[js.Array[Service]]) {
	bindings.CallGetServices(
		js.Pointer(&ret),
		deviceAddress.Ref(),
	)

	return
}

// TryGetServices calls the function "WEBEXT.bluetoothLowEnergy.getServices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetServices(deviceAddress js.String) (ret js.Promise[js.Array[Service]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetServices(
		js.Pointer(&ret), js.Pointer(&exception),
		deviceAddress.Ref(),
	)

	return
}

// HasFuncNotifyCharacteristicValueChanged returns true if the function "WEBEXT.bluetoothLowEnergy.notifyCharacteristicValueChanged" exists.
func HasFuncNotifyCharacteristicValueChanged() bool {
	return js.True == bindings.HasFuncNotifyCharacteristicValueChanged()
}

// FuncNotifyCharacteristicValueChanged returns the function "WEBEXT.bluetoothLowEnergy.notifyCharacteristicValueChanged".
func FuncNotifyCharacteristicValueChanged() (fn js.Func[func(characteristicId js.String, notification Notification) js.Promise[js.Void]]) {
	bindings.FuncNotifyCharacteristicValueChanged(
		js.Pointer(&fn),
	)
	return
}

// NotifyCharacteristicValueChanged calls the function "WEBEXT.bluetoothLowEnergy.notifyCharacteristicValueChanged" directly.
func NotifyCharacteristicValueChanged(characteristicId js.String, notification Notification) (ret js.Promise[js.Void]) {
	bindings.CallNotifyCharacteristicValueChanged(
		js.Pointer(&ret),
		characteristicId.Ref(),
		js.Pointer(&notification),
	)

	return
}

// TryNotifyCharacteristicValueChanged calls the function "WEBEXT.bluetoothLowEnergy.notifyCharacteristicValueChanged"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryNotifyCharacteristicValueChanged(characteristicId js.String, notification Notification) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNotifyCharacteristicValueChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		characteristicId.Ref(),
		js.Pointer(&notification),
	)

	return
}

type OnCharacteristicReadRequestEventCallbackFunc func(this js.Ref, request *Request, characteristicId js.String) js.Ref

func (fn OnCharacteristicReadRequestEventCallbackFunc) Register() js.Func[func(request *Request, characteristicId js.String)] {
	return js.RegisterCallback[func(request *Request, characteristicId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCharacteristicReadRequestEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Request
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnCharacteristicReadRequestEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, request *Request, characteristicId js.String) js.Ref
	Arg T
}

func (cb *OnCharacteristicReadRequestEventCallback[T]) Register() js.Func[func(request *Request, characteristicId js.String)] {
	return js.RegisterCallback[func(request *Request, characteristicId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCharacteristicReadRequestEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Request
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnCharacteristicReadRequest returns true if the function "WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.addListener" exists.
func HasFuncOnCharacteristicReadRequest() bool {
	return js.True == bindings.HasFuncOnCharacteristicReadRequest()
}

// FuncOnCharacteristicReadRequest returns the function "WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.addListener".
func FuncOnCharacteristicReadRequest() (fn js.Func[func(callback js.Func[func(request *Request, characteristicId js.String)])]) {
	bindings.FuncOnCharacteristicReadRequest(
		js.Pointer(&fn),
	)
	return
}

// OnCharacteristicReadRequest calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.addListener" directly.
func OnCharacteristicReadRequest(callback js.Func[func(request *Request, characteristicId js.String)]) (ret js.Void) {
	bindings.CallOnCharacteristicReadRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCharacteristicReadRequest calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCharacteristicReadRequest(callback js.Func[func(request *Request, characteristicId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCharacteristicReadRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCharacteristicReadRequest returns true if the function "WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.removeListener" exists.
func HasFuncOffCharacteristicReadRequest() bool {
	return js.True == bindings.HasFuncOffCharacteristicReadRequest()
}

// FuncOffCharacteristicReadRequest returns the function "WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.removeListener".
func FuncOffCharacteristicReadRequest() (fn js.Func[func(callback js.Func[func(request *Request, characteristicId js.String)])]) {
	bindings.FuncOffCharacteristicReadRequest(
		js.Pointer(&fn),
	)
	return
}

// OffCharacteristicReadRequest calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.removeListener" directly.
func OffCharacteristicReadRequest(callback js.Func[func(request *Request, characteristicId js.String)]) (ret js.Void) {
	bindings.CallOffCharacteristicReadRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCharacteristicReadRequest calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCharacteristicReadRequest(callback js.Func[func(request *Request, characteristicId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCharacteristicReadRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCharacteristicReadRequest returns true if the function "WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.hasListener" exists.
func HasFuncHasOnCharacteristicReadRequest() bool {
	return js.True == bindings.HasFuncHasOnCharacteristicReadRequest()
}

// FuncHasOnCharacteristicReadRequest returns the function "WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.hasListener".
func FuncHasOnCharacteristicReadRequest() (fn js.Func[func(callback js.Func[func(request *Request, characteristicId js.String)]) bool]) {
	bindings.FuncHasOnCharacteristicReadRequest(
		js.Pointer(&fn),
	)
	return
}

// HasOnCharacteristicReadRequest calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.hasListener" directly.
func HasOnCharacteristicReadRequest(callback js.Func[func(request *Request, characteristicId js.String)]) (ret bool) {
	bindings.CallHasOnCharacteristicReadRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCharacteristicReadRequest calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicReadRequest.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCharacteristicReadRequest(callback js.Func[func(request *Request, characteristicId js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCharacteristicReadRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCharacteristicValueChangedEventCallbackFunc func(this js.Ref, characteristic *Characteristic) js.Ref

func (fn OnCharacteristicValueChangedEventCallbackFunc) Register() js.Func[func(characteristic *Characteristic)] {
	return js.RegisterCallback[func(characteristic *Characteristic)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCharacteristicValueChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Characteristic
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

type OnCharacteristicValueChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, characteristic *Characteristic) js.Ref
	Arg T
}

func (cb *OnCharacteristicValueChangedEventCallback[T]) Register() js.Func[func(characteristic *Characteristic)] {
	return js.RegisterCallback[func(characteristic *Characteristic)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCharacteristicValueChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Characteristic
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

// HasFuncOnCharacteristicValueChanged returns true if the function "WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.addListener" exists.
func HasFuncOnCharacteristicValueChanged() bool {
	return js.True == bindings.HasFuncOnCharacteristicValueChanged()
}

// FuncOnCharacteristicValueChanged returns the function "WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.addListener".
func FuncOnCharacteristicValueChanged() (fn js.Func[func(callback js.Func[func(characteristic *Characteristic)])]) {
	bindings.FuncOnCharacteristicValueChanged(
		js.Pointer(&fn),
	)
	return
}

// OnCharacteristicValueChanged calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.addListener" directly.
func OnCharacteristicValueChanged(callback js.Func[func(characteristic *Characteristic)]) (ret js.Void) {
	bindings.CallOnCharacteristicValueChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCharacteristicValueChanged calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCharacteristicValueChanged(callback js.Func[func(characteristic *Characteristic)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCharacteristicValueChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCharacteristicValueChanged returns true if the function "WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.removeListener" exists.
func HasFuncOffCharacteristicValueChanged() bool {
	return js.True == bindings.HasFuncOffCharacteristicValueChanged()
}

// FuncOffCharacteristicValueChanged returns the function "WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.removeListener".
func FuncOffCharacteristicValueChanged() (fn js.Func[func(callback js.Func[func(characteristic *Characteristic)])]) {
	bindings.FuncOffCharacteristicValueChanged(
		js.Pointer(&fn),
	)
	return
}

// OffCharacteristicValueChanged calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.removeListener" directly.
func OffCharacteristicValueChanged(callback js.Func[func(characteristic *Characteristic)]) (ret js.Void) {
	bindings.CallOffCharacteristicValueChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCharacteristicValueChanged calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCharacteristicValueChanged(callback js.Func[func(characteristic *Characteristic)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCharacteristicValueChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCharacteristicValueChanged returns true if the function "WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.hasListener" exists.
func HasFuncHasOnCharacteristicValueChanged() bool {
	return js.True == bindings.HasFuncHasOnCharacteristicValueChanged()
}

// FuncHasOnCharacteristicValueChanged returns the function "WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.hasListener".
func FuncHasOnCharacteristicValueChanged() (fn js.Func[func(callback js.Func[func(characteristic *Characteristic)]) bool]) {
	bindings.FuncHasOnCharacteristicValueChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnCharacteristicValueChanged calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.hasListener" directly.
func HasOnCharacteristicValueChanged(callback js.Func[func(characteristic *Characteristic)]) (ret bool) {
	bindings.CallHasOnCharacteristicValueChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCharacteristicValueChanged calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicValueChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCharacteristicValueChanged(callback js.Func[func(characteristic *Characteristic)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCharacteristicValueChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCharacteristicWriteRequestEventCallbackFunc func(this js.Ref, request *Request, characteristicId js.String) js.Ref

func (fn OnCharacteristicWriteRequestEventCallbackFunc) Register() js.Func[func(request *Request, characteristicId js.String)] {
	return js.RegisterCallback[func(request *Request, characteristicId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCharacteristicWriteRequestEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Request
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnCharacteristicWriteRequestEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, request *Request, characteristicId js.String) js.Ref
	Arg T
}

func (cb *OnCharacteristicWriteRequestEventCallback[T]) Register() js.Func[func(request *Request, characteristicId js.String)] {
	return js.RegisterCallback[func(request *Request, characteristicId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCharacteristicWriteRequestEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Request
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnCharacteristicWriteRequest returns true if the function "WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.addListener" exists.
func HasFuncOnCharacteristicWriteRequest() bool {
	return js.True == bindings.HasFuncOnCharacteristicWriteRequest()
}

// FuncOnCharacteristicWriteRequest returns the function "WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.addListener".
func FuncOnCharacteristicWriteRequest() (fn js.Func[func(callback js.Func[func(request *Request, characteristicId js.String)])]) {
	bindings.FuncOnCharacteristicWriteRequest(
		js.Pointer(&fn),
	)
	return
}

// OnCharacteristicWriteRequest calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.addListener" directly.
func OnCharacteristicWriteRequest(callback js.Func[func(request *Request, characteristicId js.String)]) (ret js.Void) {
	bindings.CallOnCharacteristicWriteRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCharacteristicWriteRequest calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCharacteristicWriteRequest(callback js.Func[func(request *Request, characteristicId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCharacteristicWriteRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCharacteristicWriteRequest returns true if the function "WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.removeListener" exists.
func HasFuncOffCharacteristicWriteRequest() bool {
	return js.True == bindings.HasFuncOffCharacteristicWriteRequest()
}

// FuncOffCharacteristicWriteRequest returns the function "WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.removeListener".
func FuncOffCharacteristicWriteRequest() (fn js.Func[func(callback js.Func[func(request *Request, characteristicId js.String)])]) {
	bindings.FuncOffCharacteristicWriteRequest(
		js.Pointer(&fn),
	)
	return
}

// OffCharacteristicWriteRequest calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.removeListener" directly.
func OffCharacteristicWriteRequest(callback js.Func[func(request *Request, characteristicId js.String)]) (ret js.Void) {
	bindings.CallOffCharacteristicWriteRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCharacteristicWriteRequest calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCharacteristicWriteRequest(callback js.Func[func(request *Request, characteristicId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCharacteristicWriteRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCharacteristicWriteRequest returns true if the function "WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.hasListener" exists.
func HasFuncHasOnCharacteristicWriteRequest() bool {
	return js.True == bindings.HasFuncHasOnCharacteristicWriteRequest()
}

// FuncHasOnCharacteristicWriteRequest returns the function "WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.hasListener".
func FuncHasOnCharacteristicWriteRequest() (fn js.Func[func(callback js.Func[func(request *Request, characteristicId js.String)]) bool]) {
	bindings.FuncHasOnCharacteristicWriteRequest(
		js.Pointer(&fn),
	)
	return
}

// HasOnCharacteristicWriteRequest calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.hasListener" directly.
func HasOnCharacteristicWriteRequest(callback js.Func[func(request *Request, characteristicId js.String)]) (ret bool) {
	bindings.CallHasOnCharacteristicWriteRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCharacteristicWriteRequest calls the function "WEBEXT.bluetoothLowEnergy.onCharacteristicWriteRequest.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCharacteristicWriteRequest(callback js.Func[func(request *Request, characteristicId js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCharacteristicWriteRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDescriptorReadRequestEventCallbackFunc func(this js.Ref, request *Request, descriptorId js.String) js.Ref

func (fn OnDescriptorReadRequestEventCallbackFunc) Register() js.Func[func(request *Request, descriptorId js.String)] {
	return js.RegisterCallback[func(request *Request, descriptorId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDescriptorReadRequestEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Request
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDescriptorReadRequestEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, request *Request, descriptorId js.String) js.Ref
	Arg T
}

func (cb *OnDescriptorReadRequestEventCallback[T]) Register() js.Func[func(request *Request, descriptorId js.String)] {
	return js.RegisterCallback[func(request *Request, descriptorId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDescriptorReadRequestEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Request
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDescriptorReadRequest returns true if the function "WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.addListener" exists.
func HasFuncOnDescriptorReadRequest() bool {
	return js.True == bindings.HasFuncOnDescriptorReadRequest()
}

// FuncOnDescriptorReadRequest returns the function "WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.addListener".
func FuncOnDescriptorReadRequest() (fn js.Func[func(callback js.Func[func(request *Request, descriptorId js.String)])]) {
	bindings.FuncOnDescriptorReadRequest(
		js.Pointer(&fn),
	)
	return
}

// OnDescriptorReadRequest calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.addListener" directly.
func OnDescriptorReadRequest(callback js.Func[func(request *Request, descriptorId js.String)]) (ret js.Void) {
	bindings.CallOnDescriptorReadRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDescriptorReadRequest calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDescriptorReadRequest(callback js.Func[func(request *Request, descriptorId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDescriptorReadRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDescriptorReadRequest returns true if the function "WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.removeListener" exists.
func HasFuncOffDescriptorReadRequest() bool {
	return js.True == bindings.HasFuncOffDescriptorReadRequest()
}

// FuncOffDescriptorReadRequest returns the function "WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.removeListener".
func FuncOffDescriptorReadRequest() (fn js.Func[func(callback js.Func[func(request *Request, descriptorId js.String)])]) {
	bindings.FuncOffDescriptorReadRequest(
		js.Pointer(&fn),
	)
	return
}

// OffDescriptorReadRequest calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.removeListener" directly.
func OffDescriptorReadRequest(callback js.Func[func(request *Request, descriptorId js.String)]) (ret js.Void) {
	bindings.CallOffDescriptorReadRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDescriptorReadRequest calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDescriptorReadRequest(callback js.Func[func(request *Request, descriptorId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDescriptorReadRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDescriptorReadRequest returns true if the function "WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.hasListener" exists.
func HasFuncHasOnDescriptorReadRequest() bool {
	return js.True == bindings.HasFuncHasOnDescriptorReadRequest()
}

// FuncHasOnDescriptorReadRequest returns the function "WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.hasListener".
func FuncHasOnDescriptorReadRequest() (fn js.Func[func(callback js.Func[func(request *Request, descriptorId js.String)]) bool]) {
	bindings.FuncHasOnDescriptorReadRequest(
		js.Pointer(&fn),
	)
	return
}

// HasOnDescriptorReadRequest calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.hasListener" directly.
func HasOnDescriptorReadRequest(callback js.Func[func(request *Request, descriptorId js.String)]) (ret bool) {
	bindings.CallHasOnDescriptorReadRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDescriptorReadRequest calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorReadRequest.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDescriptorReadRequest(callback js.Func[func(request *Request, descriptorId js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDescriptorReadRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDescriptorValueChangedEventCallbackFunc func(this js.Ref, descriptor *Descriptor) js.Ref

func (fn OnDescriptorValueChangedEventCallbackFunc) Register() js.Func[func(descriptor *Descriptor)] {
	return js.RegisterCallback[func(descriptor *Descriptor)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDescriptorValueChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Descriptor
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

type OnDescriptorValueChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, descriptor *Descriptor) js.Ref
	Arg T
}

func (cb *OnDescriptorValueChangedEventCallback[T]) Register() js.Func[func(descriptor *Descriptor)] {
	return js.RegisterCallback[func(descriptor *Descriptor)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDescriptorValueChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Descriptor
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

// HasFuncOnDescriptorValueChanged returns true if the function "WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.addListener" exists.
func HasFuncOnDescriptorValueChanged() bool {
	return js.True == bindings.HasFuncOnDescriptorValueChanged()
}

// FuncOnDescriptorValueChanged returns the function "WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.addListener".
func FuncOnDescriptorValueChanged() (fn js.Func[func(callback js.Func[func(descriptor *Descriptor)])]) {
	bindings.FuncOnDescriptorValueChanged(
		js.Pointer(&fn),
	)
	return
}

// OnDescriptorValueChanged calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.addListener" directly.
func OnDescriptorValueChanged(callback js.Func[func(descriptor *Descriptor)]) (ret js.Void) {
	bindings.CallOnDescriptorValueChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDescriptorValueChanged calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDescriptorValueChanged(callback js.Func[func(descriptor *Descriptor)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDescriptorValueChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDescriptorValueChanged returns true if the function "WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.removeListener" exists.
func HasFuncOffDescriptorValueChanged() bool {
	return js.True == bindings.HasFuncOffDescriptorValueChanged()
}

// FuncOffDescriptorValueChanged returns the function "WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.removeListener".
func FuncOffDescriptorValueChanged() (fn js.Func[func(callback js.Func[func(descriptor *Descriptor)])]) {
	bindings.FuncOffDescriptorValueChanged(
		js.Pointer(&fn),
	)
	return
}

// OffDescriptorValueChanged calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.removeListener" directly.
func OffDescriptorValueChanged(callback js.Func[func(descriptor *Descriptor)]) (ret js.Void) {
	bindings.CallOffDescriptorValueChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDescriptorValueChanged calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDescriptorValueChanged(callback js.Func[func(descriptor *Descriptor)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDescriptorValueChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDescriptorValueChanged returns true if the function "WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.hasListener" exists.
func HasFuncHasOnDescriptorValueChanged() bool {
	return js.True == bindings.HasFuncHasOnDescriptorValueChanged()
}

// FuncHasOnDescriptorValueChanged returns the function "WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.hasListener".
func FuncHasOnDescriptorValueChanged() (fn js.Func[func(callback js.Func[func(descriptor *Descriptor)]) bool]) {
	bindings.FuncHasOnDescriptorValueChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnDescriptorValueChanged calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.hasListener" directly.
func HasOnDescriptorValueChanged(callback js.Func[func(descriptor *Descriptor)]) (ret bool) {
	bindings.CallHasOnDescriptorValueChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDescriptorValueChanged calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorValueChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDescriptorValueChanged(callback js.Func[func(descriptor *Descriptor)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDescriptorValueChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDescriptorWriteRequestEventCallbackFunc func(this js.Ref, request *Request, descriptorId js.String) js.Ref

func (fn OnDescriptorWriteRequestEventCallbackFunc) Register() js.Func[func(request *Request, descriptorId js.String)] {
	return js.RegisterCallback[func(request *Request, descriptorId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDescriptorWriteRequestEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Request
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDescriptorWriteRequestEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, request *Request, descriptorId js.String) js.Ref
	Arg T
}

func (cb *OnDescriptorWriteRequestEventCallback[T]) Register() js.Func[func(request *Request, descriptorId js.String)] {
	return js.RegisterCallback[func(request *Request, descriptorId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDescriptorWriteRequestEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Request
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDescriptorWriteRequest returns true if the function "WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.addListener" exists.
func HasFuncOnDescriptorWriteRequest() bool {
	return js.True == bindings.HasFuncOnDescriptorWriteRequest()
}

// FuncOnDescriptorWriteRequest returns the function "WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.addListener".
func FuncOnDescriptorWriteRequest() (fn js.Func[func(callback js.Func[func(request *Request, descriptorId js.String)])]) {
	bindings.FuncOnDescriptorWriteRequest(
		js.Pointer(&fn),
	)
	return
}

// OnDescriptorWriteRequest calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.addListener" directly.
func OnDescriptorWriteRequest(callback js.Func[func(request *Request, descriptorId js.String)]) (ret js.Void) {
	bindings.CallOnDescriptorWriteRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDescriptorWriteRequest calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDescriptorWriteRequest(callback js.Func[func(request *Request, descriptorId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDescriptorWriteRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDescriptorWriteRequest returns true if the function "WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.removeListener" exists.
func HasFuncOffDescriptorWriteRequest() bool {
	return js.True == bindings.HasFuncOffDescriptorWriteRequest()
}

// FuncOffDescriptorWriteRequest returns the function "WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.removeListener".
func FuncOffDescriptorWriteRequest() (fn js.Func[func(callback js.Func[func(request *Request, descriptorId js.String)])]) {
	bindings.FuncOffDescriptorWriteRequest(
		js.Pointer(&fn),
	)
	return
}

// OffDescriptorWriteRequest calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.removeListener" directly.
func OffDescriptorWriteRequest(callback js.Func[func(request *Request, descriptorId js.String)]) (ret js.Void) {
	bindings.CallOffDescriptorWriteRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDescriptorWriteRequest calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDescriptorWriteRequest(callback js.Func[func(request *Request, descriptorId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDescriptorWriteRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDescriptorWriteRequest returns true if the function "WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.hasListener" exists.
func HasFuncHasOnDescriptorWriteRequest() bool {
	return js.True == bindings.HasFuncHasOnDescriptorWriteRequest()
}

// FuncHasOnDescriptorWriteRequest returns the function "WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.hasListener".
func FuncHasOnDescriptorWriteRequest() (fn js.Func[func(callback js.Func[func(request *Request, descriptorId js.String)]) bool]) {
	bindings.FuncHasOnDescriptorWriteRequest(
		js.Pointer(&fn),
	)
	return
}

// HasOnDescriptorWriteRequest calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.hasListener" directly.
func HasOnDescriptorWriteRequest(callback js.Func[func(request *Request, descriptorId js.String)]) (ret bool) {
	bindings.CallHasOnDescriptorWriteRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDescriptorWriteRequest calls the function "WEBEXT.bluetoothLowEnergy.onDescriptorWriteRequest.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDescriptorWriteRequest(callback js.Func[func(request *Request, descriptorId js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDescriptorWriteRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnServiceAddedEventCallbackFunc func(this js.Ref, service *Service) js.Ref

func (fn OnServiceAddedEventCallbackFunc) Register() js.Func[func(service *Service)] {
	return js.RegisterCallback[func(service *Service)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnServiceAddedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Service
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

type OnServiceAddedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, service *Service) js.Ref
	Arg T
}

func (cb *OnServiceAddedEventCallback[T]) Register() js.Func[func(service *Service)] {
	return js.RegisterCallback[func(service *Service)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnServiceAddedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Service
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

// HasFuncOnServiceAdded returns true if the function "WEBEXT.bluetoothLowEnergy.onServiceAdded.addListener" exists.
func HasFuncOnServiceAdded() bool {
	return js.True == bindings.HasFuncOnServiceAdded()
}

// FuncOnServiceAdded returns the function "WEBEXT.bluetoothLowEnergy.onServiceAdded.addListener".
func FuncOnServiceAdded() (fn js.Func[func(callback js.Func[func(service *Service)])]) {
	bindings.FuncOnServiceAdded(
		js.Pointer(&fn),
	)
	return
}

// OnServiceAdded calls the function "WEBEXT.bluetoothLowEnergy.onServiceAdded.addListener" directly.
func OnServiceAdded(callback js.Func[func(service *Service)]) (ret js.Void) {
	bindings.CallOnServiceAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnServiceAdded calls the function "WEBEXT.bluetoothLowEnergy.onServiceAdded.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnServiceAdded(callback js.Func[func(service *Service)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnServiceAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffServiceAdded returns true if the function "WEBEXT.bluetoothLowEnergy.onServiceAdded.removeListener" exists.
func HasFuncOffServiceAdded() bool {
	return js.True == bindings.HasFuncOffServiceAdded()
}

// FuncOffServiceAdded returns the function "WEBEXT.bluetoothLowEnergy.onServiceAdded.removeListener".
func FuncOffServiceAdded() (fn js.Func[func(callback js.Func[func(service *Service)])]) {
	bindings.FuncOffServiceAdded(
		js.Pointer(&fn),
	)
	return
}

// OffServiceAdded calls the function "WEBEXT.bluetoothLowEnergy.onServiceAdded.removeListener" directly.
func OffServiceAdded(callback js.Func[func(service *Service)]) (ret js.Void) {
	bindings.CallOffServiceAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffServiceAdded calls the function "WEBEXT.bluetoothLowEnergy.onServiceAdded.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffServiceAdded(callback js.Func[func(service *Service)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffServiceAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnServiceAdded returns true if the function "WEBEXT.bluetoothLowEnergy.onServiceAdded.hasListener" exists.
func HasFuncHasOnServiceAdded() bool {
	return js.True == bindings.HasFuncHasOnServiceAdded()
}

// FuncHasOnServiceAdded returns the function "WEBEXT.bluetoothLowEnergy.onServiceAdded.hasListener".
func FuncHasOnServiceAdded() (fn js.Func[func(callback js.Func[func(service *Service)]) bool]) {
	bindings.FuncHasOnServiceAdded(
		js.Pointer(&fn),
	)
	return
}

// HasOnServiceAdded calls the function "WEBEXT.bluetoothLowEnergy.onServiceAdded.hasListener" directly.
func HasOnServiceAdded(callback js.Func[func(service *Service)]) (ret bool) {
	bindings.CallHasOnServiceAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnServiceAdded calls the function "WEBEXT.bluetoothLowEnergy.onServiceAdded.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnServiceAdded(callback js.Func[func(service *Service)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnServiceAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnServiceChangedEventCallbackFunc func(this js.Ref, service *Service) js.Ref

func (fn OnServiceChangedEventCallbackFunc) Register() js.Func[func(service *Service)] {
	return js.RegisterCallback[func(service *Service)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnServiceChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Service
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

type OnServiceChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, service *Service) js.Ref
	Arg T
}

func (cb *OnServiceChangedEventCallback[T]) Register() js.Func[func(service *Service)] {
	return js.RegisterCallback[func(service *Service)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnServiceChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Service
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

// HasFuncOnServiceChanged returns true if the function "WEBEXT.bluetoothLowEnergy.onServiceChanged.addListener" exists.
func HasFuncOnServiceChanged() bool {
	return js.True == bindings.HasFuncOnServiceChanged()
}

// FuncOnServiceChanged returns the function "WEBEXT.bluetoothLowEnergy.onServiceChanged.addListener".
func FuncOnServiceChanged() (fn js.Func[func(callback js.Func[func(service *Service)])]) {
	bindings.FuncOnServiceChanged(
		js.Pointer(&fn),
	)
	return
}

// OnServiceChanged calls the function "WEBEXT.bluetoothLowEnergy.onServiceChanged.addListener" directly.
func OnServiceChanged(callback js.Func[func(service *Service)]) (ret js.Void) {
	bindings.CallOnServiceChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnServiceChanged calls the function "WEBEXT.bluetoothLowEnergy.onServiceChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnServiceChanged(callback js.Func[func(service *Service)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnServiceChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffServiceChanged returns true if the function "WEBEXT.bluetoothLowEnergy.onServiceChanged.removeListener" exists.
func HasFuncOffServiceChanged() bool {
	return js.True == bindings.HasFuncOffServiceChanged()
}

// FuncOffServiceChanged returns the function "WEBEXT.bluetoothLowEnergy.onServiceChanged.removeListener".
func FuncOffServiceChanged() (fn js.Func[func(callback js.Func[func(service *Service)])]) {
	bindings.FuncOffServiceChanged(
		js.Pointer(&fn),
	)
	return
}

// OffServiceChanged calls the function "WEBEXT.bluetoothLowEnergy.onServiceChanged.removeListener" directly.
func OffServiceChanged(callback js.Func[func(service *Service)]) (ret js.Void) {
	bindings.CallOffServiceChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffServiceChanged calls the function "WEBEXT.bluetoothLowEnergy.onServiceChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffServiceChanged(callback js.Func[func(service *Service)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffServiceChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnServiceChanged returns true if the function "WEBEXT.bluetoothLowEnergy.onServiceChanged.hasListener" exists.
func HasFuncHasOnServiceChanged() bool {
	return js.True == bindings.HasFuncHasOnServiceChanged()
}

// FuncHasOnServiceChanged returns the function "WEBEXT.bluetoothLowEnergy.onServiceChanged.hasListener".
func FuncHasOnServiceChanged() (fn js.Func[func(callback js.Func[func(service *Service)]) bool]) {
	bindings.FuncHasOnServiceChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnServiceChanged calls the function "WEBEXT.bluetoothLowEnergy.onServiceChanged.hasListener" directly.
func HasOnServiceChanged(callback js.Func[func(service *Service)]) (ret bool) {
	bindings.CallHasOnServiceChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnServiceChanged calls the function "WEBEXT.bluetoothLowEnergy.onServiceChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnServiceChanged(callback js.Func[func(service *Service)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnServiceChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnServiceRemovedEventCallbackFunc func(this js.Ref, service *Service) js.Ref

func (fn OnServiceRemovedEventCallbackFunc) Register() js.Func[func(service *Service)] {
	return js.RegisterCallback[func(service *Service)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnServiceRemovedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Service
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

type OnServiceRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, service *Service) js.Ref
	Arg T
}

func (cb *OnServiceRemovedEventCallback[T]) Register() js.Func[func(service *Service)] {
	return js.RegisterCallback[func(service *Service)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnServiceRemovedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Service
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

// HasFuncOnServiceRemoved returns true if the function "WEBEXT.bluetoothLowEnergy.onServiceRemoved.addListener" exists.
func HasFuncOnServiceRemoved() bool {
	return js.True == bindings.HasFuncOnServiceRemoved()
}

// FuncOnServiceRemoved returns the function "WEBEXT.bluetoothLowEnergy.onServiceRemoved.addListener".
func FuncOnServiceRemoved() (fn js.Func[func(callback js.Func[func(service *Service)])]) {
	bindings.FuncOnServiceRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnServiceRemoved calls the function "WEBEXT.bluetoothLowEnergy.onServiceRemoved.addListener" directly.
func OnServiceRemoved(callback js.Func[func(service *Service)]) (ret js.Void) {
	bindings.CallOnServiceRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnServiceRemoved calls the function "WEBEXT.bluetoothLowEnergy.onServiceRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnServiceRemoved(callback js.Func[func(service *Service)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnServiceRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffServiceRemoved returns true if the function "WEBEXT.bluetoothLowEnergy.onServiceRemoved.removeListener" exists.
func HasFuncOffServiceRemoved() bool {
	return js.True == bindings.HasFuncOffServiceRemoved()
}

// FuncOffServiceRemoved returns the function "WEBEXT.bluetoothLowEnergy.onServiceRemoved.removeListener".
func FuncOffServiceRemoved() (fn js.Func[func(callback js.Func[func(service *Service)])]) {
	bindings.FuncOffServiceRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffServiceRemoved calls the function "WEBEXT.bluetoothLowEnergy.onServiceRemoved.removeListener" directly.
func OffServiceRemoved(callback js.Func[func(service *Service)]) (ret js.Void) {
	bindings.CallOffServiceRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffServiceRemoved calls the function "WEBEXT.bluetoothLowEnergy.onServiceRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffServiceRemoved(callback js.Func[func(service *Service)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffServiceRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnServiceRemoved returns true if the function "WEBEXT.bluetoothLowEnergy.onServiceRemoved.hasListener" exists.
func HasFuncHasOnServiceRemoved() bool {
	return js.True == bindings.HasFuncHasOnServiceRemoved()
}

// FuncHasOnServiceRemoved returns the function "WEBEXT.bluetoothLowEnergy.onServiceRemoved.hasListener".
func FuncHasOnServiceRemoved() (fn js.Func[func(callback js.Func[func(service *Service)]) bool]) {
	bindings.FuncHasOnServiceRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnServiceRemoved calls the function "WEBEXT.bluetoothLowEnergy.onServiceRemoved.hasListener" directly.
func HasOnServiceRemoved(callback js.Func[func(service *Service)]) (ret bool) {
	bindings.CallHasOnServiceRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnServiceRemoved calls the function "WEBEXT.bluetoothLowEnergy.onServiceRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnServiceRemoved(callback js.Func[func(service *Service)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnServiceRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncReadCharacteristicValue returns true if the function "WEBEXT.bluetoothLowEnergy.readCharacteristicValue" exists.
func HasFuncReadCharacteristicValue() bool {
	return js.True == bindings.HasFuncReadCharacteristicValue()
}

// FuncReadCharacteristicValue returns the function "WEBEXT.bluetoothLowEnergy.readCharacteristicValue".
func FuncReadCharacteristicValue() (fn js.Func[func(characteristicId js.String) js.Promise[Characteristic]]) {
	bindings.FuncReadCharacteristicValue(
		js.Pointer(&fn),
	)
	return
}

// ReadCharacteristicValue calls the function "WEBEXT.bluetoothLowEnergy.readCharacteristicValue" directly.
func ReadCharacteristicValue(characteristicId js.String) (ret js.Promise[Characteristic]) {
	bindings.CallReadCharacteristicValue(
		js.Pointer(&ret),
		characteristicId.Ref(),
	)

	return
}

// TryReadCharacteristicValue calls the function "WEBEXT.bluetoothLowEnergy.readCharacteristicValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReadCharacteristicValue(characteristicId js.String) (ret js.Promise[Characteristic], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadCharacteristicValue(
		js.Pointer(&ret), js.Pointer(&exception),
		characteristicId.Ref(),
	)

	return
}

// HasFuncReadDescriptorValue returns true if the function "WEBEXT.bluetoothLowEnergy.readDescriptorValue" exists.
func HasFuncReadDescriptorValue() bool {
	return js.True == bindings.HasFuncReadDescriptorValue()
}

// FuncReadDescriptorValue returns the function "WEBEXT.bluetoothLowEnergy.readDescriptorValue".
func FuncReadDescriptorValue() (fn js.Func[func(descriptorId js.String) js.Promise[Descriptor]]) {
	bindings.FuncReadDescriptorValue(
		js.Pointer(&fn),
	)
	return
}

// ReadDescriptorValue calls the function "WEBEXT.bluetoothLowEnergy.readDescriptorValue" directly.
func ReadDescriptorValue(descriptorId js.String) (ret js.Promise[Descriptor]) {
	bindings.CallReadDescriptorValue(
		js.Pointer(&ret),
		descriptorId.Ref(),
	)

	return
}

// TryReadDescriptorValue calls the function "WEBEXT.bluetoothLowEnergy.readDescriptorValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReadDescriptorValue(descriptorId js.String) (ret js.Promise[Descriptor], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadDescriptorValue(
		js.Pointer(&ret), js.Pointer(&exception),
		descriptorId.Ref(),
	)

	return
}

// HasFuncRegisterAdvertisement returns true if the function "WEBEXT.bluetoothLowEnergy.registerAdvertisement" exists.
func HasFuncRegisterAdvertisement() bool {
	return js.True == bindings.HasFuncRegisterAdvertisement()
}

// FuncRegisterAdvertisement returns the function "WEBEXT.bluetoothLowEnergy.registerAdvertisement".
func FuncRegisterAdvertisement() (fn js.Func[func(advertisement Advertisement) js.Promise[js.Number[int32]]]) {
	bindings.FuncRegisterAdvertisement(
		js.Pointer(&fn),
	)
	return
}

// RegisterAdvertisement calls the function "WEBEXT.bluetoothLowEnergy.registerAdvertisement" directly.
func RegisterAdvertisement(advertisement Advertisement) (ret js.Promise[js.Number[int32]]) {
	bindings.CallRegisterAdvertisement(
		js.Pointer(&ret),
		js.Pointer(&advertisement),
	)

	return
}

// TryRegisterAdvertisement calls the function "WEBEXT.bluetoothLowEnergy.registerAdvertisement"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRegisterAdvertisement(advertisement Advertisement) (ret js.Promise[js.Number[int32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRegisterAdvertisement(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&advertisement),
	)

	return
}

// HasFuncRegisterService returns true if the function "WEBEXT.bluetoothLowEnergy.registerService" exists.
func HasFuncRegisterService() bool {
	return js.True == bindings.HasFuncRegisterService()
}

// FuncRegisterService returns the function "WEBEXT.bluetoothLowEnergy.registerService".
func FuncRegisterService() (fn js.Func[func(serviceId js.String) js.Promise[js.Void]]) {
	bindings.FuncRegisterService(
		js.Pointer(&fn),
	)
	return
}

// RegisterService calls the function "WEBEXT.bluetoothLowEnergy.registerService" directly.
func RegisterService(serviceId js.String) (ret js.Promise[js.Void]) {
	bindings.CallRegisterService(
		js.Pointer(&ret),
		serviceId.Ref(),
	)

	return
}

// TryRegisterService calls the function "WEBEXT.bluetoothLowEnergy.registerService"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRegisterService(serviceId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRegisterService(
		js.Pointer(&ret), js.Pointer(&exception),
		serviceId.Ref(),
	)

	return
}

// HasFuncRemoveService returns true if the function "WEBEXT.bluetoothLowEnergy.removeService" exists.
func HasFuncRemoveService() bool {
	return js.True == bindings.HasFuncRemoveService()
}

// FuncRemoveService returns the function "WEBEXT.bluetoothLowEnergy.removeService".
func FuncRemoveService() (fn js.Func[func(serviceId js.String) js.Promise[js.Void]]) {
	bindings.FuncRemoveService(
		js.Pointer(&fn),
	)
	return
}

// RemoveService calls the function "WEBEXT.bluetoothLowEnergy.removeService" directly.
func RemoveService(serviceId js.String) (ret js.Promise[js.Void]) {
	bindings.CallRemoveService(
		js.Pointer(&ret),
		serviceId.Ref(),
	)

	return
}

// TryRemoveService calls the function "WEBEXT.bluetoothLowEnergy.removeService"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveService(serviceId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveService(
		js.Pointer(&ret), js.Pointer(&exception),
		serviceId.Ref(),
	)

	return
}

// HasFuncResetAdvertising returns true if the function "WEBEXT.bluetoothLowEnergy.resetAdvertising" exists.
func HasFuncResetAdvertising() bool {
	return js.True == bindings.HasFuncResetAdvertising()
}

// FuncResetAdvertising returns the function "WEBEXT.bluetoothLowEnergy.resetAdvertising".
func FuncResetAdvertising() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncResetAdvertising(
		js.Pointer(&fn),
	)
	return
}

// ResetAdvertising calls the function "WEBEXT.bluetoothLowEnergy.resetAdvertising" directly.
func ResetAdvertising() (ret js.Promise[js.Void]) {
	bindings.CallResetAdvertising(
		js.Pointer(&ret),
	)

	return
}

// TryResetAdvertising calls the function "WEBEXT.bluetoothLowEnergy.resetAdvertising"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryResetAdvertising() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResetAdvertising(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSendRequestResponse returns true if the function "WEBEXT.bluetoothLowEnergy.sendRequestResponse" exists.
func HasFuncSendRequestResponse() bool {
	return js.True == bindings.HasFuncSendRequestResponse()
}

// FuncSendRequestResponse returns the function "WEBEXT.bluetoothLowEnergy.sendRequestResponse".
func FuncSendRequestResponse() (fn js.Func[func(response Response)]) {
	bindings.FuncSendRequestResponse(
		js.Pointer(&fn),
	)
	return
}

// SendRequestResponse calls the function "WEBEXT.bluetoothLowEnergy.sendRequestResponse" directly.
func SendRequestResponse(response Response) (ret js.Void) {
	bindings.CallSendRequestResponse(
		js.Pointer(&ret),
		js.Pointer(&response),
	)

	return
}

// TrySendRequestResponse calls the function "WEBEXT.bluetoothLowEnergy.sendRequestResponse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendRequestResponse(response Response) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendRequestResponse(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&response),
	)

	return
}

// HasFuncSetAdvertisingInterval returns true if the function "WEBEXT.bluetoothLowEnergy.setAdvertisingInterval" exists.
func HasFuncSetAdvertisingInterval() bool {
	return js.True == bindings.HasFuncSetAdvertisingInterval()
}

// FuncSetAdvertisingInterval returns the function "WEBEXT.bluetoothLowEnergy.setAdvertisingInterval".
func FuncSetAdvertisingInterval() (fn js.Func[func(minInterval int32, maxInterval int32) js.Promise[js.Void]]) {
	bindings.FuncSetAdvertisingInterval(
		js.Pointer(&fn),
	)
	return
}

// SetAdvertisingInterval calls the function "WEBEXT.bluetoothLowEnergy.setAdvertisingInterval" directly.
func SetAdvertisingInterval(minInterval int32, maxInterval int32) (ret js.Promise[js.Void]) {
	bindings.CallSetAdvertisingInterval(
		js.Pointer(&ret),
		int32(minInterval),
		int32(maxInterval),
	)

	return
}

// TrySetAdvertisingInterval calls the function "WEBEXT.bluetoothLowEnergy.setAdvertisingInterval"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetAdvertisingInterval(minInterval int32, maxInterval int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetAdvertisingInterval(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(minInterval),
		int32(maxInterval),
	)

	return
}

// HasFuncStartCharacteristicNotifications returns true if the function "WEBEXT.bluetoothLowEnergy.startCharacteristicNotifications" exists.
func HasFuncStartCharacteristicNotifications() bool {
	return js.True == bindings.HasFuncStartCharacteristicNotifications()
}

// FuncStartCharacteristicNotifications returns the function "WEBEXT.bluetoothLowEnergy.startCharacteristicNotifications".
func FuncStartCharacteristicNotifications() (fn js.Func[func(characteristicId js.String, properties NotificationProperties) js.Promise[js.Void]]) {
	bindings.FuncStartCharacteristicNotifications(
		js.Pointer(&fn),
	)
	return
}

// StartCharacteristicNotifications calls the function "WEBEXT.bluetoothLowEnergy.startCharacteristicNotifications" directly.
func StartCharacteristicNotifications(characteristicId js.String, properties NotificationProperties) (ret js.Promise[js.Void]) {
	bindings.CallStartCharacteristicNotifications(
		js.Pointer(&ret),
		characteristicId.Ref(),
		js.Pointer(&properties),
	)

	return
}

// TryStartCharacteristicNotifications calls the function "WEBEXT.bluetoothLowEnergy.startCharacteristicNotifications"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartCharacteristicNotifications(characteristicId js.String, properties NotificationProperties) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartCharacteristicNotifications(
		js.Pointer(&ret), js.Pointer(&exception),
		characteristicId.Ref(),
		js.Pointer(&properties),
	)

	return
}

// HasFuncStopCharacteristicNotifications returns true if the function "WEBEXT.bluetoothLowEnergy.stopCharacteristicNotifications" exists.
func HasFuncStopCharacteristicNotifications() bool {
	return js.True == bindings.HasFuncStopCharacteristicNotifications()
}

// FuncStopCharacteristicNotifications returns the function "WEBEXT.bluetoothLowEnergy.stopCharacteristicNotifications".
func FuncStopCharacteristicNotifications() (fn js.Func[func(characteristicId js.String) js.Promise[js.Void]]) {
	bindings.FuncStopCharacteristicNotifications(
		js.Pointer(&fn),
	)
	return
}

// StopCharacteristicNotifications calls the function "WEBEXT.bluetoothLowEnergy.stopCharacteristicNotifications" directly.
func StopCharacteristicNotifications(characteristicId js.String) (ret js.Promise[js.Void]) {
	bindings.CallStopCharacteristicNotifications(
		js.Pointer(&ret),
		characteristicId.Ref(),
	)

	return
}

// TryStopCharacteristicNotifications calls the function "WEBEXT.bluetoothLowEnergy.stopCharacteristicNotifications"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStopCharacteristicNotifications(characteristicId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStopCharacteristicNotifications(
		js.Pointer(&ret), js.Pointer(&exception),
		characteristicId.Ref(),
	)

	return
}

// HasFuncUnregisterAdvertisement returns true if the function "WEBEXT.bluetoothLowEnergy.unregisterAdvertisement" exists.
func HasFuncUnregisterAdvertisement() bool {
	return js.True == bindings.HasFuncUnregisterAdvertisement()
}

// FuncUnregisterAdvertisement returns the function "WEBEXT.bluetoothLowEnergy.unregisterAdvertisement".
func FuncUnregisterAdvertisement() (fn js.Func[func(advertisementId int32) js.Promise[js.Void]]) {
	bindings.FuncUnregisterAdvertisement(
		js.Pointer(&fn),
	)
	return
}

// UnregisterAdvertisement calls the function "WEBEXT.bluetoothLowEnergy.unregisterAdvertisement" directly.
func UnregisterAdvertisement(advertisementId int32) (ret js.Promise[js.Void]) {
	bindings.CallUnregisterAdvertisement(
		js.Pointer(&ret),
		int32(advertisementId),
	)

	return
}

// TryUnregisterAdvertisement calls the function "WEBEXT.bluetoothLowEnergy.unregisterAdvertisement"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUnregisterAdvertisement(advertisementId int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUnregisterAdvertisement(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(advertisementId),
	)

	return
}

// HasFuncUnregisterService returns true if the function "WEBEXT.bluetoothLowEnergy.unregisterService" exists.
func HasFuncUnregisterService() bool {
	return js.True == bindings.HasFuncUnregisterService()
}

// FuncUnregisterService returns the function "WEBEXT.bluetoothLowEnergy.unregisterService".
func FuncUnregisterService() (fn js.Func[func(serviceId js.String) js.Promise[js.Void]]) {
	bindings.FuncUnregisterService(
		js.Pointer(&fn),
	)
	return
}

// UnregisterService calls the function "WEBEXT.bluetoothLowEnergy.unregisterService" directly.
func UnregisterService(serviceId js.String) (ret js.Promise[js.Void]) {
	bindings.CallUnregisterService(
		js.Pointer(&ret),
		serviceId.Ref(),
	)

	return
}

// TryUnregisterService calls the function "WEBEXT.bluetoothLowEnergy.unregisterService"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUnregisterService(serviceId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUnregisterService(
		js.Pointer(&ret), js.Pointer(&exception),
		serviceId.Ref(),
	)

	return
}

// HasFuncWriteCharacteristicValue returns true if the function "WEBEXT.bluetoothLowEnergy.writeCharacteristicValue" exists.
func HasFuncWriteCharacteristicValue() bool {
	return js.True == bindings.HasFuncWriteCharacteristicValue()
}

// FuncWriteCharacteristicValue returns the function "WEBEXT.bluetoothLowEnergy.writeCharacteristicValue".
func FuncWriteCharacteristicValue() (fn js.Func[func(characteristicId js.String, value js.ArrayBuffer) js.Promise[js.Void]]) {
	bindings.FuncWriteCharacteristicValue(
		js.Pointer(&fn),
	)
	return
}

// WriteCharacteristicValue calls the function "WEBEXT.bluetoothLowEnergy.writeCharacteristicValue" directly.
func WriteCharacteristicValue(characteristicId js.String, value js.ArrayBuffer) (ret js.Promise[js.Void]) {
	bindings.CallWriteCharacteristicValue(
		js.Pointer(&ret),
		characteristicId.Ref(),
		value.Ref(),
	)

	return
}

// TryWriteCharacteristicValue calls the function "WEBEXT.bluetoothLowEnergy.writeCharacteristicValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryWriteCharacteristicValue(characteristicId js.String, value js.ArrayBuffer) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWriteCharacteristicValue(
		js.Pointer(&ret), js.Pointer(&exception),
		characteristicId.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncWriteDescriptorValue returns true if the function "WEBEXT.bluetoothLowEnergy.writeDescriptorValue" exists.
func HasFuncWriteDescriptorValue() bool {
	return js.True == bindings.HasFuncWriteDescriptorValue()
}

// FuncWriteDescriptorValue returns the function "WEBEXT.bluetoothLowEnergy.writeDescriptorValue".
func FuncWriteDescriptorValue() (fn js.Func[func(descriptorId js.String, value js.ArrayBuffer) js.Promise[js.Void]]) {
	bindings.FuncWriteDescriptorValue(
		js.Pointer(&fn),
	)
	return
}

// WriteDescriptorValue calls the function "WEBEXT.bluetoothLowEnergy.writeDescriptorValue" directly.
func WriteDescriptorValue(descriptorId js.String, value js.ArrayBuffer) (ret js.Promise[js.Void]) {
	bindings.CallWriteDescriptorValue(
		js.Pointer(&ret),
		descriptorId.Ref(),
		value.Ref(),
	)

	return
}

// TryWriteDescriptorValue calls the function "WEBEXT.bluetoothLowEnergy.writeDescriptorValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryWriteDescriptorValue(descriptorId js.String, value js.ArrayBuffer) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWriteDescriptorValue(
		js.Pointer(&ret), js.Pointer(&exception),
		descriptorId.Ref(),
		value.Ref(),
	)

	return
}
