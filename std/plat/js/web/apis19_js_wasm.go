// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type BluetoothServiceDataFilterInit struct {
	// Service is "BluetoothServiceDataFilterInit.service"
	//
	// Required
	Service BluetoothServiceUUID
	// DataPrefix is "BluetoothServiceDataFilterInit.dataPrefix"
	//
	// Optional
	DataPrefix BufferSource
	// Mask is "BluetoothServiceDataFilterInit.mask"
	//
	// Optional
	Mask BufferSource

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BluetoothServiceDataFilterInit with all fields set.
func (p BluetoothServiceDataFilterInit) FromRef(ref js.Ref) BluetoothServiceDataFilterInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BluetoothServiceDataFilterInit in the application heap.
func (p BluetoothServiceDataFilterInit) New() js.Ref {
	return bindings.BluetoothServiceDataFilterInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BluetoothServiceDataFilterInit) UpdateFrom(ref js.Ref) {
	bindings.BluetoothServiceDataFilterInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BluetoothServiceDataFilterInit) Update(ref js.Ref) {
	bindings.BluetoothServiceDataFilterInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BluetoothServiceDataFilterInit) FreeMembers(recursive bool) {
	js.Free(
		p.Service.Ref(),
		p.DataPrefix.Ref(),
		p.Mask.Ref(),
	)
	p.Service = p.Service.FromRef(js.Undefined)
	p.DataPrefix = p.DataPrefix.FromRef(js.Undefined)
	p.Mask = p.Mask.FromRef(js.Undefined)
}

type BluetoothLEScanFilterInit struct {
	// Services is "BluetoothLEScanFilterInit.services"
	//
	// Optional
	Services js.Array[BluetoothServiceUUID]
	// Name is "BluetoothLEScanFilterInit.name"
	//
	// Optional
	Name js.String
	// NamePrefix is "BluetoothLEScanFilterInit.namePrefix"
	//
	// Optional
	NamePrefix js.String
	// ManufacturerData is "BluetoothLEScanFilterInit.manufacturerData"
	//
	// Optional
	ManufacturerData js.Array[BluetoothManufacturerDataFilterInit]
	// ServiceData is "BluetoothLEScanFilterInit.serviceData"
	//
	// Optional
	ServiceData js.Array[BluetoothServiceDataFilterInit]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BluetoothLEScanFilterInit with all fields set.
func (p BluetoothLEScanFilterInit) FromRef(ref js.Ref) BluetoothLEScanFilterInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BluetoothLEScanFilterInit in the application heap.
func (p BluetoothLEScanFilterInit) New() js.Ref {
	return bindings.BluetoothLEScanFilterInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BluetoothLEScanFilterInit) UpdateFrom(ref js.Ref) {
	bindings.BluetoothLEScanFilterInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BluetoothLEScanFilterInit) Update(ref js.Ref) {
	bindings.BluetoothLEScanFilterInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BluetoothLEScanFilterInit) FreeMembers(recursive bool) {
	js.Free(
		p.Services.Ref(),
		p.Name.Ref(),
		p.NamePrefix.Ref(),
		p.ManufacturerData.Ref(),
		p.ServiceData.Ref(),
	)
	p.Services = p.Services.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.NamePrefix = p.NamePrefix.FromRef(js.Undefined)
	p.ManufacturerData = p.ManufacturerData.FromRef(js.Undefined)
	p.ServiceData = p.ServiceData.FromRef(js.Undefined)
}

type RequestDeviceOptions struct {
	// Filters is "RequestDeviceOptions.filters"
	//
	// Optional
	Filters js.Array[BluetoothLEScanFilterInit]
	// ExclusionFilters is "RequestDeviceOptions.exclusionFilters"
	//
	// Optional
	ExclusionFilters js.Array[BluetoothLEScanFilterInit]
	// OptionalServices is "RequestDeviceOptions.optionalServices"
	//
	// Optional, defaults to [].
	OptionalServices js.Array[BluetoothServiceUUID]
	// OptionalManufacturerData is "RequestDeviceOptions.optionalManufacturerData"
	//
	// Optional, defaults to [].
	OptionalManufacturerData js.Array[uint16]
	// AcceptAllDevices is "RequestDeviceOptions.acceptAllDevices"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_AcceptAllDevices MUST be set to true to make this field effective.
	AcceptAllDevices bool

	FFI_USE_AcceptAllDevices bool // for AcceptAllDevices.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RequestDeviceOptions with all fields set.
func (p RequestDeviceOptions) FromRef(ref js.Ref) RequestDeviceOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RequestDeviceOptions in the application heap.
func (p RequestDeviceOptions) New() js.Ref {
	return bindings.RequestDeviceOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RequestDeviceOptions) UpdateFrom(ref js.Ref) {
	bindings.RequestDeviceOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RequestDeviceOptions) Update(ref js.Ref) {
	bindings.RequestDeviceOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RequestDeviceOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Filters.Ref(),
		p.ExclusionFilters.Ref(),
		p.OptionalServices.Ref(),
		p.OptionalManufacturerData.Ref(),
	)
	p.Filters = p.Filters.FromRef(js.Undefined)
	p.ExclusionFilters = p.ExclusionFilters.FromRef(js.Undefined)
	p.OptionalServices = p.OptionalServices.FromRef(js.Undefined)
	p.OptionalManufacturerData = p.OptionalManufacturerData.FromRef(js.Undefined)
}

type Bluetooth struct {
	EventTarget
}

func (this Bluetooth) Once() Bluetooth {
	this.ref.Once()
	return this
}

func (this Bluetooth) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this Bluetooth) FromRef(ref js.Ref) Bluetooth {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this Bluetooth) Free() {
	this.ref.Free()
}

// ReferringDevice returns the value of property "Bluetooth.referringDevice".
//
// It returns ok=false if there is no such property.
func (this Bluetooth) ReferringDevice() (ret BluetoothDevice, ok bool) {
	ok = js.True == bindings.GetBluetoothReferringDevice(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetAvailability returns true if the method "Bluetooth.getAvailability" exists.
func (this Bluetooth) HasFuncGetAvailability() bool {
	return js.True == bindings.HasFuncBluetoothGetAvailability(
		this.ref,
	)
}

// FuncGetAvailability returns the method "Bluetooth.getAvailability".
func (this Bluetooth) FuncGetAvailability() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncBluetoothGetAvailability(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAvailability calls the method "Bluetooth.getAvailability".
func (this Bluetooth) GetAvailability() (ret js.Promise[js.Boolean]) {
	bindings.CallBluetoothGetAvailability(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetAvailability calls the method "Bluetooth.getAvailability"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Bluetooth) TryGetAvailability() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothGetAvailability(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDevices returns true if the method "Bluetooth.getDevices" exists.
func (this Bluetooth) HasFuncGetDevices() bool {
	return js.True == bindings.HasFuncBluetoothGetDevices(
		this.ref,
	)
}

// FuncGetDevices returns the method "Bluetooth.getDevices".
func (this Bluetooth) FuncGetDevices() (fn js.Func[func() js.Promise[js.Array[BluetoothDevice]]]) {
	bindings.FuncBluetoothGetDevices(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDevices calls the method "Bluetooth.getDevices".
func (this Bluetooth) GetDevices() (ret js.Promise[js.Array[BluetoothDevice]]) {
	bindings.CallBluetoothGetDevices(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetDevices calls the method "Bluetooth.getDevices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Bluetooth) TryGetDevices() (ret js.Promise[js.Array[BluetoothDevice]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothGetDevices(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestDevice returns true if the method "Bluetooth.requestDevice" exists.
func (this Bluetooth) HasFuncRequestDevice() bool {
	return js.True == bindings.HasFuncBluetoothRequestDevice(
		this.ref,
	)
}

// FuncRequestDevice returns the method "Bluetooth.requestDevice".
func (this Bluetooth) FuncRequestDevice() (fn js.Func[func(options RequestDeviceOptions) js.Promise[BluetoothDevice]]) {
	bindings.FuncBluetoothRequestDevice(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestDevice calls the method "Bluetooth.requestDevice".
func (this Bluetooth) RequestDevice(options RequestDeviceOptions) (ret js.Promise[BluetoothDevice]) {
	bindings.CallBluetoothRequestDevice(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestDevice calls the method "Bluetooth.requestDevice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Bluetooth) TryRequestDevice(options RequestDeviceOptions) (ret js.Promise[BluetoothDevice], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRequestDevice(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRequestDevice1 returns true if the method "Bluetooth.requestDevice" exists.
func (this Bluetooth) HasFuncRequestDevice1() bool {
	return js.True == bindings.HasFuncBluetoothRequestDevice1(
		this.ref,
	)
}

// FuncRequestDevice1 returns the method "Bluetooth.requestDevice".
func (this Bluetooth) FuncRequestDevice1() (fn js.Func[func() js.Promise[BluetoothDevice]]) {
	bindings.FuncBluetoothRequestDevice1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestDevice1 calls the method "Bluetooth.requestDevice".
func (this Bluetooth) RequestDevice1() (ret js.Promise[BluetoothDevice]) {
	bindings.CallBluetoothRequestDevice1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestDevice1 calls the method "Bluetooth.requestDevice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Bluetooth) TryRequestDevice1() (ret js.Promise[BluetoothDevice], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRequestDevice1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type BluetoothManufacturerDataMap struct {
	ref js.Ref
}

func (this BluetoothManufacturerDataMap) Once() BluetoothManufacturerDataMap {
	this.ref.Once()
	return this
}

func (this BluetoothManufacturerDataMap) Ref() js.Ref {
	return this.ref
}

func (this BluetoothManufacturerDataMap) FromRef(ref js.Ref) BluetoothManufacturerDataMap {
	this.ref = ref
	return this
}

func (this BluetoothManufacturerDataMap) Free() {
	this.ref.Free()
}

type BluetoothServiceDataMap struct {
	ref js.Ref
}

func (this BluetoothServiceDataMap) Once() BluetoothServiceDataMap {
	this.ref.Once()
	return this
}

func (this BluetoothServiceDataMap) Ref() js.Ref {
	return this.ref
}

func (this BluetoothServiceDataMap) FromRef(ref js.Ref) BluetoothServiceDataMap {
	this.ref = ref
	return this
}

func (this BluetoothServiceDataMap) Free() {
	this.ref.Free()
}

type BluetoothAdvertisingEventInit struct {
	// Device is "BluetoothAdvertisingEventInit.device"
	//
	// Required
	Device BluetoothDevice
	// Uuids is "BluetoothAdvertisingEventInit.uuids"
	//
	// Optional
	Uuids js.Array[OneOf_String_Uint32]
	// Name is "BluetoothAdvertisingEventInit.name"
	//
	// Optional
	Name js.String
	// Appearance is "BluetoothAdvertisingEventInit.appearance"
	//
	// Optional
	//
	// NOTE: FFI_USE_Appearance MUST be set to true to make this field effective.
	Appearance uint16
	// TxPower is "BluetoothAdvertisingEventInit.txPower"
	//
	// Optional
	//
	// NOTE: FFI_USE_TxPower MUST be set to true to make this field effective.
	TxPower int8
	// Rssi is "BluetoothAdvertisingEventInit.rssi"
	//
	// Optional
	//
	// NOTE: FFI_USE_Rssi MUST be set to true to make this field effective.
	Rssi int8
	// ManufacturerData is "BluetoothAdvertisingEventInit.manufacturerData"
	//
	// Optional
	ManufacturerData BluetoothManufacturerDataMap
	// ServiceData is "BluetoothAdvertisingEventInit.serviceData"
	//
	// Optional
	ServiceData BluetoothServiceDataMap
	// Bubbles is "BluetoothAdvertisingEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "BluetoothAdvertisingEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "BluetoothAdvertisingEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Appearance bool // for Appearance.
	FFI_USE_TxPower    bool // for TxPower.
	FFI_USE_Rssi       bool // for Rssi.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BluetoothAdvertisingEventInit with all fields set.
func (p BluetoothAdvertisingEventInit) FromRef(ref js.Ref) BluetoothAdvertisingEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BluetoothAdvertisingEventInit in the application heap.
func (p BluetoothAdvertisingEventInit) New() js.Ref {
	return bindings.BluetoothAdvertisingEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BluetoothAdvertisingEventInit) UpdateFrom(ref js.Ref) {
	bindings.BluetoothAdvertisingEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BluetoothAdvertisingEventInit) Update(ref js.Ref) {
	bindings.BluetoothAdvertisingEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BluetoothAdvertisingEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Device.Ref(),
		p.Uuids.Ref(),
		p.Name.Ref(),
		p.ManufacturerData.Ref(),
		p.ServiceData.Ref(),
	)
	p.Device = p.Device.FromRef(js.Undefined)
	p.Uuids = p.Uuids.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.ManufacturerData = p.ManufacturerData.FromRef(js.Undefined)
	p.ServiceData = p.ServiceData.FromRef(js.Undefined)
}

func NewBluetoothAdvertisingEvent(typ js.String, init BluetoothAdvertisingEventInit) (ret BluetoothAdvertisingEvent) {
	ret.ref = bindings.NewBluetoothAdvertisingEventByBluetoothAdvertisingEvent(
		typ.Ref(),
		js.Pointer(&init))
	return
}

type BluetoothAdvertisingEvent struct {
	Event
}

func (this BluetoothAdvertisingEvent) Once() BluetoothAdvertisingEvent {
	this.ref.Once()
	return this
}

func (this BluetoothAdvertisingEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this BluetoothAdvertisingEvent) FromRef(ref js.Ref) BluetoothAdvertisingEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this BluetoothAdvertisingEvent) Free() {
	this.ref.Free()
}

// Device returns the value of property "BluetoothAdvertisingEvent.device".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) Device() (ret BluetoothDevice, ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventDevice(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Uuids returns the value of property "BluetoothAdvertisingEvent.uuids".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) Uuids() (ret js.FrozenArray[UUID], ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventUuids(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "BluetoothAdvertisingEvent.name".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Appearance returns the value of property "BluetoothAdvertisingEvent.appearance".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) Appearance() (ret uint16, ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventAppearance(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TxPower returns the value of property "BluetoothAdvertisingEvent.txPower".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) TxPower() (ret int8, ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventTxPower(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Rssi returns the value of property "BluetoothAdvertisingEvent.rssi".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) Rssi() (ret int8, ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventRssi(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ManufacturerData returns the value of property "BluetoothAdvertisingEvent.manufacturerData".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) ManufacturerData() (ret BluetoothManufacturerDataMap, ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventManufacturerData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ServiceData returns the value of property "BluetoothAdvertisingEvent.serviceData".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) ServiceData() (ret BluetoothServiceDataMap, ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventServiceData(
		this.ref, js.Pointer(&ret),
	)
	return
}

type BluetoothDataFilterInit struct {
	// DataPrefix is "BluetoothDataFilterInit.dataPrefix"
	//
	// Optional
	DataPrefix BufferSource
	// Mask is "BluetoothDataFilterInit.mask"
	//
	// Optional
	Mask BufferSource

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BluetoothDataFilterInit with all fields set.
func (p BluetoothDataFilterInit) FromRef(ref js.Ref) BluetoothDataFilterInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BluetoothDataFilterInit in the application heap.
func (p BluetoothDataFilterInit) New() js.Ref {
	return bindings.BluetoothDataFilterInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BluetoothDataFilterInit) UpdateFrom(ref js.Ref) {
	bindings.BluetoothDataFilterInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BluetoothDataFilterInit) Update(ref js.Ref) {
	bindings.BluetoothDataFilterInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BluetoothDataFilterInit) FreeMembers(recursive bool) {
	js.Free(
		p.DataPrefix.Ref(),
		p.Mask.Ref(),
	)
	p.DataPrefix = p.DataPrefix.FromRef(js.Undefined)
	p.Mask = p.Mask.FromRef(js.Undefined)
}

type BluetoothPermissionDescriptor struct {
	// DeviceId is "BluetoothPermissionDescriptor.deviceId"
	//
	// Optional
	DeviceId js.String
	// Filters is "BluetoothPermissionDescriptor.filters"
	//
	// Optional
	Filters js.Array[BluetoothLEScanFilterInit]
	// OptionalServices is "BluetoothPermissionDescriptor.optionalServices"
	//
	// Optional, defaults to [].
	OptionalServices js.Array[BluetoothServiceUUID]
	// OptionalManufacturerData is "BluetoothPermissionDescriptor.optionalManufacturerData"
	//
	// Optional, defaults to [].
	OptionalManufacturerData js.Array[uint16]
	// AcceptAllDevices is "BluetoothPermissionDescriptor.acceptAllDevices"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_AcceptAllDevices MUST be set to true to make this field effective.
	AcceptAllDevices bool
	// Name is "BluetoothPermissionDescriptor.name"
	//
	// Required
	Name js.String

	FFI_USE_AcceptAllDevices bool // for AcceptAllDevices.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BluetoothPermissionDescriptor with all fields set.
func (p BluetoothPermissionDescriptor) FromRef(ref js.Ref) BluetoothPermissionDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BluetoothPermissionDescriptor in the application heap.
func (p BluetoothPermissionDescriptor) New() js.Ref {
	return bindings.BluetoothPermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BluetoothPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.BluetoothPermissionDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BluetoothPermissionDescriptor) Update(ref js.Ref) {
	bindings.BluetoothPermissionDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BluetoothPermissionDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.DeviceId.Ref(),
		p.Filters.Ref(),
		p.OptionalServices.Ref(),
		p.OptionalManufacturerData.Ref(),
		p.Name.Ref(),
	)
	p.DeviceId = p.DeviceId.FromRef(js.Undefined)
	p.Filters = p.Filters.FromRef(js.Undefined)
	p.OptionalServices = p.OptionalServices.FromRef(js.Undefined)
	p.OptionalManufacturerData = p.OptionalManufacturerData.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type BluetoothPermissionResult struct {
	PermissionStatus
}

func (this BluetoothPermissionResult) Once() BluetoothPermissionResult {
	this.ref.Once()
	return this
}

func (this BluetoothPermissionResult) Ref() js.Ref {
	return this.PermissionStatus.Ref()
}

func (this BluetoothPermissionResult) FromRef(ref js.Ref) BluetoothPermissionResult {
	this.PermissionStatus = this.PermissionStatus.FromRef(ref)
	return this
}

func (this BluetoothPermissionResult) Free() {
	this.ref.Free()
}

// Devices returns the value of property "BluetoothPermissionResult.devices".
//
// It returns ok=false if there is no such property.
func (this BluetoothPermissionResult) Devices() (ret js.FrozenArray[BluetoothDevice], ok bool) {
	ok = js.True == bindings.GetBluetoothPermissionResultDevices(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDevices sets the value of property "BluetoothPermissionResult.devices" to val.
//
// It returns false if the property cannot be set.
func (this BluetoothPermissionResult) SetDevices(val js.FrozenArray[BluetoothDevice]) bool {
	return js.True == bindings.SetBluetoothPermissionResultDevices(
		this.ref,
		val.Ref(),
	)
}

type BluetoothPermissionStorage struct {
	// AllowedDevices is "BluetoothPermissionStorage.allowedDevices"
	//
	// Required
	AllowedDevices js.Array[AllowedBluetoothDevice]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BluetoothPermissionStorage with all fields set.
func (p BluetoothPermissionStorage) FromRef(ref js.Ref) BluetoothPermissionStorage {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BluetoothPermissionStorage in the application heap.
func (p BluetoothPermissionStorage) New() js.Ref {
	return bindings.BluetoothPermissionStorageJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BluetoothPermissionStorage) UpdateFrom(ref js.Ref) {
	bindings.BluetoothPermissionStorageJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BluetoothPermissionStorage) Update(ref js.Ref) {
	bindings.BluetoothPermissionStorageJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BluetoothPermissionStorage) FreeMembers(recursive bool) {
	js.Free(
		p.AllowedDevices.Ref(),
	)
	p.AllowedDevices = p.AllowedDevices.FromRef(js.Undefined)
}

type BluetoothUUID struct {
	ref js.Ref
}

func (this BluetoothUUID) Once() BluetoothUUID {
	this.ref.Once()
	return this
}

func (this BluetoothUUID) Ref() js.Ref {
	return this.ref
}

func (this BluetoothUUID) FromRef(ref js.Ref) BluetoothUUID {
	this.ref = ref
	return this
}

func (this BluetoothUUID) Free() {
	this.ref.Free()
}

// HasFuncGetService returns true if the static method "BluetoothUUID.getService" exists.
func (this BluetoothUUID) HasFuncGetService() bool {
	return js.True == bindings.HasFuncBluetoothUUIDGetService(
		this.ref,
	)
}

// FuncGetService returns the static method "BluetoothUUID.getService".
func (this BluetoothUUID) FuncGetService() (fn js.Func[func(name OneOf_String_Uint32) UUID]) {
	bindings.FuncBluetoothUUIDGetService(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetService calls the static method "BluetoothUUID.getService".
func (this BluetoothUUID) GetService(name OneOf_String_Uint32) (ret UUID) {
	bindings.CallBluetoothUUIDGetService(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetService calls the static method "BluetoothUUID.getService"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothUUID) TryGetService(name OneOf_String_Uint32) (ret UUID, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothUUIDGetService(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncGetCharacteristic returns true if the static method "BluetoothUUID.getCharacteristic" exists.
func (this BluetoothUUID) HasFuncGetCharacteristic() bool {
	return js.True == bindings.HasFuncBluetoothUUIDGetCharacteristic(
		this.ref,
	)
}

// FuncGetCharacteristic returns the static method "BluetoothUUID.getCharacteristic".
func (this BluetoothUUID) FuncGetCharacteristic() (fn js.Func[func(name OneOf_String_Uint32) UUID]) {
	bindings.FuncBluetoothUUIDGetCharacteristic(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCharacteristic calls the static method "BluetoothUUID.getCharacteristic".
func (this BluetoothUUID) GetCharacteristic(name OneOf_String_Uint32) (ret UUID) {
	bindings.CallBluetoothUUIDGetCharacteristic(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetCharacteristic calls the static method "BluetoothUUID.getCharacteristic"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothUUID) TryGetCharacteristic(name OneOf_String_Uint32) (ret UUID, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothUUIDGetCharacteristic(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncGetDescriptor returns true if the static method "BluetoothUUID.getDescriptor" exists.
func (this BluetoothUUID) HasFuncGetDescriptor() bool {
	return js.True == bindings.HasFuncBluetoothUUIDGetDescriptor(
		this.ref,
	)
}

// FuncGetDescriptor returns the static method "BluetoothUUID.getDescriptor".
func (this BluetoothUUID) FuncGetDescriptor() (fn js.Func[func(name OneOf_String_Uint32) UUID]) {
	bindings.FuncBluetoothUUIDGetDescriptor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDescriptor calls the static method "BluetoothUUID.getDescriptor".
func (this BluetoothUUID) GetDescriptor(name OneOf_String_Uint32) (ret UUID) {
	bindings.CallBluetoothUUIDGetDescriptor(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetDescriptor calls the static method "BluetoothUUID.getDescriptor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothUUID) TryGetDescriptor(name OneOf_String_Uint32) (ret UUID, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothUUIDGetDescriptor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncCanonicalUUID returns true if the static method "BluetoothUUID.canonicalUUID" exists.
func (this BluetoothUUID) HasFuncCanonicalUUID() bool {
	return js.True == bindings.HasFuncBluetoothUUIDCanonicalUUID(
		this.ref,
	)
}

// FuncCanonicalUUID returns the static method "BluetoothUUID.canonicalUUID".
func (this BluetoothUUID) FuncCanonicalUUID() (fn js.Func[func(alias uint32) UUID]) {
	bindings.FuncBluetoothUUIDCanonicalUUID(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CanonicalUUID calls the static method "BluetoothUUID.canonicalUUID".
func (this BluetoothUUID) CanonicalUUID(alias uint32) (ret UUID) {
	bindings.CallBluetoothUUIDCanonicalUUID(
		this.ref, js.Pointer(&ret),
		uint32(alias),
	)

	return
}

// TryCanonicalUUID calls the static method "BluetoothUUID.canonicalUUID"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothUUID) TryCanonicalUUID(alias uint32) (ret UUID, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothUUIDCanonicalUUID(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(alias),
	)

	return
}

type BreakType uint32

const (
	_ BreakType = iota

	BreakType_NONE
	BreakType_LINE
	BreakType_COLUMN
	BreakType_PAGE
	BreakType_REGION
)

func (BreakType) FromRef(str js.Ref) BreakType {
	return BreakType(bindings.ConstOfBreakType(str))
}

func (x BreakType) String() (string, bool) {
	switch x {
	case BreakType_NONE:
		return "none", true
	case BreakType_LINE:
		return "line", true
	case BreakType_COLUMN:
		return "column", true
	case BreakType_PAGE:
		return "page", true
	case BreakType_REGION:
		return "region", true
	default:
		return "", false
	}
}

type IntrinsicSizes struct {
	ref js.Ref
}

func (this IntrinsicSizes) Once() IntrinsicSizes {
	this.ref.Once()
	return this
}

func (this IntrinsicSizes) Ref() js.Ref {
	return this.ref
}

func (this IntrinsicSizes) FromRef(ref js.Ref) IntrinsicSizes {
	this.ref = ref
	return this
}

func (this IntrinsicSizes) Free() {
	this.ref.Free()
}

// MinContentSize returns the value of property "IntrinsicSizes.minContentSize".
//
// It returns ok=false if there is no such property.
func (this IntrinsicSizes) MinContentSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetIntrinsicSizesMinContentSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxContentSize returns the value of property "IntrinsicSizes.maxContentSize".
//
// It returns ok=false if there is no such property.
func (this IntrinsicSizes) MaxContentSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetIntrinsicSizesMaxContentSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

type LayoutFragment struct {
	ref js.Ref
}

func (this LayoutFragment) Once() LayoutFragment {
	this.ref.Once()
	return this
}

func (this LayoutFragment) Ref() js.Ref {
	return this.ref
}

func (this LayoutFragment) FromRef(ref js.Ref) LayoutFragment {
	this.ref = ref
	return this
}

func (this LayoutFragment) Free() {
	this.ref.Free()
}

// InlineSize returns the value of property "LayoutFragment.inlineSize".
//
// It returns ok=false if there is no such property.
func (this LayoutFragment) InlineSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutFragmentInlineSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BlockSize returns the value of property "LayoutFragment.blockSize".
//
// It returns ok=false if there is no such property.
func (this LayoutFragment) BlockSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutFragmentBlockSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InlineOffset returns the value of property "LayoutFragment.inlineOffset".
//
// It returns ok=false if there is no such property.
func (this LayoutFragment) InlineOffset() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutFragmentInlineOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetInlineOffset sets the value of property "LayoutFragment.inlineOffset" to val.
//
// It returns false if the property cannot be set.
func (this LayoutFragment) SetInlineOffset(val float64) bool {
	return js.True == bindings.SetLayoutFragmentInlineOffset(
		this.ref,
		float64(val),
	)
}

// BlockOffset returns the value of property "LayoutFragment.blockOffset".
//
// It returns ok=false if there is no such property.
func (this LayoutFragment) BlockOffset() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutFragmentBlockOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBlockOffset sets the value of property "LayoutFragment.blockOffset" to val.
//
// It returns false if the property cannot be set.
func (this LayoutFragment) SetBlockOffset(val float64) bool {
	return js.True == bindings.SetLayoutFragmentBlockOffset(
		this.ref,
		float64(val),
	)
}

// Data returns the value of property "LayoutFragment.data".
//
// It returns ok=false if there is no such property.
func (this LayoutFragment) Data() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetLayoutFragmentData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BreakToken returns the value of property "LayoutFragment.breakToken".
//
// It returns ok=false if there is no such property.
func (this LayoutFragment) BreakToken() (ret ChildBreakToken, ok bool) {
	ok = js.True == bindings.GetLayoutFragmentBreakToken(
		this.ref, js.Pointer(&ret),
	)
	return
}

type LayoutConstraintsOptions struct {
	// AvailableInlineSize is "LayoutConstraintsOptions.availableInlineSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_AvailableInlineSize MUST be set to true to make this field effective.
	AvailableInlineSize float64
	// AvailableBlockSize is "LayoutConstraintsOptions.availableBlockSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_AvailableBlockSize MUST be set to true to make this field effective.
	AvailableBlockSize float64
	// FixedInlineSize is "LayoutConstraintsOptions.fixedInlineSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_FixedInlineSize MUST be set to true to make this field effective.
	FixedInlineSize float64
	// FixedBlockSize is "LayoutConstraintsOptions.fixedBlockSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_FixedBlockSize MUST be set to true to make this field effective.
	FixedBlockSize float64
	// PercentageInlineSize is "LayoutConstraintsOptions.percentageInlineSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_PercentageInlineSize MUST be set to true to make this field effective.
	PercentageInlineSize float64
	// PercentageBlockSize is "LayoutConstraintsOptions.percentageBlockSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_PercentageBlockSize MUST be set to true to make this field effective.
	PercentageBlockSize float64
	// BlockFragmentationOffset is "LayoutConstraintsOptions.blockFragmentationOffset"
	//
	// Optional
	//
	// NOTE: FFI_USE_BlockFragmentationOffset MUST be set to true to make this field effective.
	BlockFragmentationOffset float64
	// BlockFragmentationType is "LayoutConstraintsOptions.blockFragmentationType"
	//
	// Optional, defaults to "none".
	BlockFragmentationType BlockFragmentationType
	// Data is "LayoutConstraintsOptions.data"
	//
	// Optional
	Data js.Any

	FFI_USE_AvailableInlineSize      bool // for AvailableInlineSize.
	FFI_USE_AvailableBlockSize       bool // for AvailableBlockSize.
	FFI_USE_FixedInlineSize          bool // for FixedInlineSize.
	FFI_USE_FixedBlockSize           bool // for FixedBlockSize.
	FFI_USE_PercentageInlineSize     bool // for PercentageInlineSize.
	FFI_USE_PercentageBlockSize      bool // for PercentageBlockSize.
	FFI_USE_BlockFragmentationOffset bool // for BlockFragmentationOffset.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LayoutConstraintsOptions with all fields set.
func (p LayoutConstraintsOptions) FromRef(ref js.Ref) LayoutConstraintsOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LayoutConstraintsOptions in the application heap.
func (p LayoutConstraintsOptions) New() js.Ref {
	return bindings.LayoutConstraintsOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LayoutConstraintsOptions) UpdateFrom(ref js.Ref) {
	bindings.LayoutConstraintsOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LayoutConstraintsOptions) Update(ref js.Ref) {
	bindings.LayoutConstraintsOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LayoutConstraintsOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
}

type LayoutChild struct {
	ref js.Ref
}

func (this LayoutChild) Once() LayoutChild {
	this.ref.Once()
	return this
}

func (this LayoutChild) Ref() js.Ref {
	return this.ref
}

func (this LayoutChild) FromRef(ref js.Ref) LayoutChild {
	this.ref = ref
	return this
}

func (this LayoutChild) Free() {
	this.ref.Free()
}

// StyleMap returns the value of property "LayoutChild.styleMap".
//
// It returns ok=false if there is no such property.
func (this LayoutChild) StyleMap() (ret StylePropertyMapReadOnly, ok bool) {
	ok = js.True == bindings.GetLayoutChildStyleMap(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncIntrinsicSizes returns true if the method "LayoutChild.intrinsicSizes" exists.
func (this LayoutChild) HasFuncIntrinsicSizes() bool {
	return js.True == bindings.HasFuncLayoutChildIntrinsicSizes(
		this.ref,
	)
}

// FuncIntrinsicSizes returns the method "LayoutChild.intrinsicSizes".
func (this LayoutChild) FuncIntrinsicSizes() (fn js.Func[func() js.Promise[IntrinsicSizes]]) {
	bindings.FuncLayoutChildIntrinsicSizes(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IntrinsicSizes calls the method "LayoutChild.intrinsicSizes".
func (this LayoutChild) IntrinsicSizes() (ret js.Promise[IntrinsicSizes]) {
	bindings.CallLayoutChildIntrinsicSizes(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIntrinsicSizes calls the method "LayoutChild.intrinsicSizes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this LayoutChild) TryIntrinsicSizes() (ret js.Promise[IntrinsicSizes], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLayoutChildIntrinsicSizes(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLayoutNextFragment returns true if the method "LayoutChild.layoutNextFragment" exists.
func (this LayoutChild) HasFuncLayoutNextFragment() bool {
	return js.True == bindings.HasFuncLayoutChildLayoutNextFragment(
		this.ref,
	)
}

// FuncLayoutNextFragment returns the method "LayoutChild.layoutNextFragment".
func (this LayoutChild) FuncLayoutNextFragment() (fn js.Func[func(constraints LayoutConstraintsOptions, breakToken ChildBreakToken) js.Promise[LayoutFragment]]) {
	bindings.FuncLayoutChildLayoutNextFragment(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LayoutNextFragment calls the method "LayoutChild.layoutNextFragment".
func (this LayoutChild) LayoutNextFragment(constraints LayoutConstraintsOptions, breakToken ChildBreakToken) (ret js.Promise[LayoutFragment]) {
	bindings.CallLayoutChildLayoutNextFragment(
		this.ref, js.Pointer(&ret),
		js.Pointer(&constraints),
		breakToken.Ref(),
	)

	return
}

// TryLayoutNextFragment calls the method "LayoutChild.layoutNextFragment"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this LayoutChild) TryLayoutNextFragment(constraints LayoutConstraintsOptions, breakToken ChildBreakToken) (ret js.Promise[LayoutFragment], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLayoutChildLayoutNextFragment(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&constraints),
		breakToken.Ref(),
	)

	return
}

type ChildBreakToken struct {
	ref js.Ref
}

func (this ChildBreakToken) Once() ChildBreakToken {
	this.ref.Once()
	return this
}

func (this ChildBreakToken) Ref() js.Ref {
	return this.ref
}

func (this ChildBreakToken) FromRef(ref js.Ref) ChildBreakToken {
	this.ref = ref
	return this
}

func (this ChildBreakToken) Free() {
	this.ref.Free()
}

// BreakType returns the value of property "ChildBreakToken.breakType".
//
// It returns ok=false if there is no such property.
func (this ChildBreakToken) BreakType() (ret BreakType, ok bool) {
	ok = js.True == bindings.GetChildBreakTokenBreakType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Child returns the value of property "ChildBreakToken.child".
//
// It returns ok=false if there is no such property.
func (this ChildBreakToken) Child() (ret LayoutChild, ok bool) {
	ok = js.True == bindings.GetChildBreakTokenChild(
		this.ref, js.Pointer(&ret),
	)
	return
}

type BreakToken struct {
	ref js.Ref
}

func (this BreakToken) Once() BreakToken {
	this.ref.Once()
	return this
}

func (this BreakToken) Ref() js.Ref {
	return this.ref
}

func (this BreakToken) FromRef(ref js.Ref) BreakToken {
	this.ref = ref
	return this
}

func (this BreakToken) Free() {
	this.ref.Free()
}

// ChildBreakTokens returns the value of property "BreakToken.childBreakTokens".
//
// It returns ok=false if there is no such property.
func (this BreakToken) ChildBreakTokens() (ret js.FrozenArray[ChildBreakToken], ok bool) {
	ok = js.True == bindings.GetBreakTokenChildBreakTokens(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Data returns the value of property "BreakToken.data".
//
// It returns ok=false if there is no such property.
func (this BreakToken) Data() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetBreakTokenData(
		this.ref, js.Pointer(&ret),
	)
	return
}

type BreakTokenOptions struct {
	// ChildBreakTokens is "BreakTokenOptions.childBreakTokens"
	//
	// Optional
	ChildBreakTokens js.Array[ChildBreakToken]
	// Data is "BreakTokenOptions.data"
	//
	// Optional, defaults to null.
	Data js.Any

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BreakTokenOptions with all fields set.
func (p BreakTokenOptions) FromRef(ref js.Ref) BreakTokenOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BreakTokenOptions in the application heap.
func (p BreakTokenOptions) New() js.Ref {
	return bindings.BreakTokenOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BreakTokenOptions) UpdateFrom(ref js.Ref) {
	bindings.BreakTokenOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BreakTokenOptions) Update(ref js.Ref) {
	bindings.BreakTokenOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BreakTokenOptions) FreeMembers(recursive bool) {
	js.Free(
		p.ChildBreakTokens.Ref(),
		p.Data.Ref(),
	)
	p.ChildBreakTokens = p.ChildBreakTokens.FromRef(js.Undefined)
	p.Data = p.Data.FromRef(js.Undefined)
}

func NewBroadcastChannel(name js.String) (ret BroadcastChannel) {
	ret.ref = bindings.NewBroadcastChannelByBroadcastChannel(
		name.Ref())
	return
}

type BroadcastChannel struct {
	EventTarget
}

func (this BroadcastChannel) Once() BroadcastChannel {
	this.ref.Once()
	return this
}

func (this BroadcastChannel) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this BroadcastChannel) FromRef(ref js.Ref) BroadcastChannel {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this BroadcastChannel) Free() {
	this.ref.Free()
}

// Name returns the value of property "BroadcastChannel.name".
//
// It returns ok=false if there is no such property.
func (this BroadcastChannel) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBroadcastChannelName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncPostMessage returns true if the method "BroadcastChannel.postMessage" exists.
func (this BroadcastChannel) HasFuncPostMessage() bool {
	return js.True == bindings.HasFuncBroadcastChannelPostMessage(
		this.ref,
	)
}

// FuncPostMessage returns the method "BroadcastChannel.postMessage".
func (this BroadcastChannel) FuncPostMessage() (fn js.Func[func(message js.Any)]) {
	bindings.FuncBroadcastChannelPostMessage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage calls the method "BroadcastChannel.postMessage".
func (this BroadcastChannel) PostMessage(message js.Any) (ret js.Void) {
	bindings.CallBroadcastChannelPostMessage(
		this.ref, js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPostMessage calls the method "BroadcastChannel.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BroadcastChannel) TryPostMessage(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBroadcastChannelPostMessage(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasFuncClose returns true if the method "BroadcastChannel.close" exists.
func (this BroadcastChannel) HasFuncClose() bool {
	return js.True == bindings.HasFuncBroadcastChannelClose(
		this.ref,
	)
}

// FuncClose returns the method "BroadcastChannel.close".
func (this BroadcastChannel) FuncClose() (fn js.Func[func()]) {
	bindings.FuncBroadcastChannelClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "BroadcastChannel.close".
func (this BroadcastChannel) Close() (ret js.Void) {
	bindings.CallBroadcastChannelClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "BroadcastChannel.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BroadcastChannel) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBroadcastChannelClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type CropTarget struct {
	ref js.Ref
}

func (this CropTarget) Once() CropTarget {
	this.ref.Once()
	return this
}

func (this CropTarget) Ref() js.Ref {
	return this.ref
}

func (this CropTarget) FromRef(ref js.Ref) CropTarget {
	this.ref = ref
	return this
}

func (this CropTarget) Free() {
	this.ref.Free()
}

// HasFuncFromElement returns true if the static method "CropTarget.fromElement" exists.
func (this CropTarget) HasFuncFromElement() bool {
	return js.True == bindings.HasFuncCropTargetFromElement(
		this.ref,
	)
}

// FuncFromElement returns the static method "CropTarget.fromElement".
func (this CropTarget) FuncFromElement() (fn js.Func[func(element Element) js.Promise[CropTarget]]) {
	bindings.FuncCropTargetFromElement(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromElement calls the static method "CropTarget.fromElement".
func (this CropTarget) FromElement(element Element) (ret js.Promise[CropTarget]) {
	bindings.CallCropTargetFromElement(
		this.ref, js.Pointer(&ret),
		element.Ref(),
	)

	return
}

// TryFromElement calls the static method "CropTarget.fromElement"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CropTarget) TryFromElement(element Element) (ret js.Promise[CropTarget], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCropTargetFromElement(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
	)

	return
}

type RestrictionTarget struct {
	ref js.Ref
}

func (this RestrictionTarget) Once() RestrictionTarget {
	this.ref.Once()
	return this
}

func (this RestrictionTarget) Ref() js.Ref {
	return this.ref
}

func (this RestrictionTarget) FromRef(ref js.Ref) RestrictionTarget {
	this.ref = ref
	return this
}

func (this RestrictionTarget) Free() {
	this.ref.Free()
}

// HasFuncFromElement returns true if the static method "RestrictionTarget.fromElement" exists.
func (this RestrictionTarget) HasFuncFromElement() bool {
	return js.True == bindings.HasFuncRestrictionTargetFromElement(
		this.ref,
	)
}

// FuncFromElement returns the static method "RestrictionTarget.fromElement".
func (this RestrictionTarget) FuncFromElement() (fn js.Func[func(element Element) js.Promise[RestrictionTarget]]) {
	bindings.FuncRestrictionTargetFromElement(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromElement calls the static method "RestrictionTarget.fromElement".
func (this RestrictionTarget) FromElement(element Element) (ret js.Promise[RestrictionTarget]) {
	bindings.CallRestrictionTargetFromElement(
		this.ref, js.Pointer(&ret),
		element.Ref(),
	)

	return
}

// TryFromElement calls the static method "RestrictionTarget.fromElement"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RestrictionTarget) TryFromElement(element Element) (ret js.Promise[RestrictionTarget], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRestrictionTargetFromElement(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
	)

	return
}

type BrowserCaptureMediaStreamTrack struct {
	MediaStreamTrack
}

func (this BrowserCaptureMediaStreamTrack) Once() BrowserCaptureMediaStreamTrack {
	this.ref.Once()
	return this
}

func (this BrowserCaptureMediaStreamTrack) Ref() js.Ref {
	return this.MediaStreamTrack.Ref()
}

func (this BrowserCaptureMediaStreamTrack) FromRef(ref js.Ref) BrowserCaptureMediaStreamTrack {
	this.MediaStreamTrack = this.MediaStreamTrack.FromRef(ref)
	return this
}

func (this BrowserCaptureMediaStreamTrack) Free() {
	this.ref.Free()
}

// HasFuncCropTo returns true if the method "BrowserCaptureMediaStreamTrack.cropTo" exists.
func (this BrowserCaptureMediaStreamTrack) HasFuncCropTo() bool {
	return js.True == bindings.HasFuncBrowserCaptureMediaStreamTrackCropTo(
		this.ref,
	)
}

// FuncCropTo returns the method "BrowserCaptureMediaStreamTrack.cropTo".
func (this BrowserCaptureMediaStreamTrack) FuncCropTo() (fn js.Func[func(cropTarget CropTarget) js.Promise[js.Void]]) {
	bindings.FuncBrowserCaptureMediaStreamTrackCropTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CropTo calls the method "BrowserCaptureMediaStreamTrack.cropTo".
func (this BrowserCaptureMediaStreamTrack) CropTo(cropTarget CropTarget) (ret js.Promise[js.Void]) {
	bindings.CallBrowserCaptureMediaStreamTrackCropTo(
		this.ref, js.Pointer(&ret),
		cropTarget.Ref(),
	)

	return
}

// TryCropTo calls the method "BrowserCaptureMediaStreamTrack.cropTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BrowserCaptureMediaStreamTrack) TryCropTo(cropTarget CropTarget) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBrowserCaptureMediaStreamTrackCropTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		cropTarget.Ref(),
	)

	return
}

// HasFuncClone returns true if the method "BrowserCaptureMediaStreamTrack.clone" exists.
func (this BrowserCaptureMediaStreamTrack) HasFuncClone() bool {
	return js.True == bindings.HasFuncBrowserCaptureMediaStreamTrackClone(
		this.ref,
	)
}

// FuncClone returns the method "BrowserCaptureMediaStreamTrack.clone".
func (this BrowserCaptureMediaStreamTrack) FuncClone() (fn js.Func[func() BrowserCaptureMediaStreamTrack]) {
	bindings.FuncBrowserCaptureMediaStreamTrackClone(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clone calls the method "BrowserCaptureMediaStreamTrack.clone".
func (this BrowserCaptureMediaStreamTrack) Clone() (ret BrowserCaptureMediaStreamTrack) {
	bindings.CallBrowserCaptureMediaStreamTrackClone(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClone calls the method "BrowserCaptureMediaStreamTrack.clone"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BrowserCaptureMediaStreamTrack) TryClone() (ret BrowserCaptureMediaStreamTrack, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBrowserCaptureMediaStreamTrackClone(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRestrictTo returns true if the method "BrowserCaptureMediaStreamTrack.restrictTo" exists.
func (this BrowserCaptureMediaStreamTrack) HasFuncRestrictTo() bool {
	return js.True == bindings.HasFuncBrowserCaptureMediaStreamTrackRestrictTo(
		this.ref,
	)
}

// FuncRestrictTo returns the method "BrowserCaptureMediaStreamTrack.restrictTo".
func (this BrowserCaptureMediaStreamTrack) FuncRestrictTo() (fn js.Func[func(RestrictionTarget RestrictionTarget) js.Promise[js.Void]]) {
	bindings.FuncBrowserCaptureMediaStreamTrackRestrictTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RestrictTo calls the method "BrowserCaptureMediaStreamTrack.restrictTo".
func (this BrowserCaptureMediaStreamTrack) RestrictTo(RestrictionTarget RestrictionTarget) (ret js.Promise[js.Void]) {
	bindings.CallBrowserCaptureMediaStreamTrackRestrictTo(
		this.ref, js.Pointer(&ret),
		RestrictionTarget.Ref(),
	)

	return
}

// TryRestrictTo calls the method "BrowserCaptureMediaStreamTrack.restrictTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BrowserCaptureMediaStreamTrack) TryRestrictTo(RestrictionTarget RestrictionTarget) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBrowserCaptureMediaStreamTrackRestrictTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		RestrictionTarget.Ref(),
	)

	return
}

type QueuingStrategyInit struct {
	// HighWaterMark is "QueuingStrategyInit.highWaterMark"
	//
	// Required
	HighWaterMark float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a QueuingStrategyInit with all fields set.
func (p QueuingStrategyInit) FromRef(ref js.Ref) QueuingStrategyInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new QueuingStrategyInit in the application heap.
func (p QueuingStrategyInit) New() js.Ref {
	return bindings.QueuingStrategyInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *QueuingStrategyInit) UpdateFrom(ref js.Ref) {
	bindings.QueuingStrategyInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *QueuingStrategyInit) Update(ref js.Ref) {
	bindings.QueuingStrategyInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *QueuingStrategyInit) FreeMembers(recursive bool) {
}

type FunctionFunc func(this js.Ref, arguments ...js.Any) js.Ref

func (fn FunctionFunc) Register() js.Func[func(arguments ...js.Any) js.Any] {
	return js.RegisterCallback[func(arguments ...js.Any) js.Any](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn FunctionFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Function[T any] struct {
	Fn  func(arg T, this js.Ref, arguments ...js.Any) js.Ref
	Arg T
}

func (cb *Function[T]) Register() js.Func[func(arguments ...js.Any) js.Any] {
	return js.RegisterCallback[func(arguments ...js.Any) js.Any](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *Function[T]) DispatchCallback(
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

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

func NewByteLengthQueuingStrategy(init QueuingStrategyInit) (ret ByteLengthQueuingStrategy) {
	ret.ref = bindings.NewByteLengthQueuingStrategyByByteLengthQueuingStrategy(
		js.Pointer(&init))
	return
}

type ByteLengthQueuingStrategy struct {
	ref js.Ref
}

func (this ByteLengthQueuingStrategy) Once() ByteLengthQueuingStrategy {
	this.ref.Once()
	return this
}

func (this ByteLengthQueuingStrategy) Ref() js.Ref {
	return this.ref
}

func (this ByteLengthQueuingStrategy) FromRef(ref js.Ref) ByteLengthQueuingStrategy {
	this.ref = ref
	return this
}

func (this ByteLengthQueuingStrategy) Free() {
	this.ref.Free()
}

// HighWaterMark returns the value of property "ByteLengthQueuingStrategy.highWaterMark".
//
// It returns ok=false if there is no such property.
func (this ByteLengthQueuingStrategy) HighWaterMark() (ret float64, ok bool) {
	ok = js.True == bindings.GetByteLengthQueuingStrategyHighWaterMark(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Size returns the value of property "ByteLengthQueuingStrategy.size".
//
// It returns ok=false if there is no such property.
func (this ByteLengthQueuingStrategy) Size() (ret js.Func[func(arguments ...js.Any) js.Any], ok bool) {
	ok = js.True == bindings.GetByteLengthQueuingStrategySize(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SecurityPolicyViolationEventDisposition uint32

const (
	_ SecurityPolicyViolationEventDisposition = iota

	SecurityPolicyViolationEventDisposition_ENFORCE
	SecurityPolicyViolationEventDisposition_REPORT
)

func (SecurityPolicyViolationEventDisposition) FromRef(str js.Ref) SecurityPolicyViolationEventDisposition {
	return SecurityPolicyViolationEventDisposition(bindings.ConstOfSecurityPolicyViolationEventDisposition(str))
}

func (x SecurityPolicyViolationEventDisposition) String() (string, bool) {
	switch x {
	case SecurityPolicyViolationEventDisposition_ENFORCE:
		return "enforce", true
	case SecurityPolicyViolationEventDisposition_REPORT:
		return "report", true
	default:
		return "", false
	}
}

type CSPViolationReportBody struct {
	ReportBody
}

func (this CSPViolationReportBody) Once() CSPViolationReportBody {
	this.ref.Once()
	return this
}

func (this CSPViolationReportBody) Ref() js.Ref {
	return this.ReportBody.Ref()
}

func (this CSPViolationReportBody) FromRef(ref js.Ref) CSPViolationReportBody {
	this.ReportBody = this.ReportBody.FromRef(ref)
	return this
}

func (this CSPViolationReportBody) Free() {
	this.ref.Free()
}

// DocumentURL returns the value of property "CSPViolationReportBody.documentURL".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) DocumentURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyDocumentURL(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Referrer returns the value of property "CSPViolationReportBody.referrer".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) Referrer() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyReferrer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BlockedURL returns the value of property "CSPViolationReportBody.blockedURL".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) BlockedURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyBlockedURL(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EffectiveDirective returns the value of property "CSPViolationReportBody.effectiveDirective".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) EffectiveDirective() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyEffectiveDirective(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OriginalPolicy returns the value of property "CSPViolationReportBody.originalPolicy".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) OriginalPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyOriginalPolicy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SourceFile returns the value of property "CSPViolationReportBody.sourceFile".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) SourceFile() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodySourceFile(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Sample returns the value of property "CSPViolationReportBody.sample".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) Sample() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodySample(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Disposition returns the value of property "CSPViolationReportBody.disposition".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) Disposition() (ret SecurityPolicyViolationEventDisposition, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyDisposition(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StatusCode returns the value of property "CSPViolationReportBody.statusCode".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) StatusCode() (ret uint16, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyStatusCode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LineNumber returns the value of property "CSPViolationReportBody.lineNumber".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) LineNumber() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyLineNumber(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ColumnNumber returns the value of property "CSPViolationReportBody.columnNumber".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) ColumnNumber() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyColumnNumber(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncToJSON returns true if the method "CSPViolationReportBody.toJSON" exists.
func (this CSPViolationReportBody) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncCSPViolationReportBodyToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "CSPViolationReportBody.toJSON".
func (this CSPViolationReportBody) FuncToJSON() (fn js.Func[func() js.Object]) {
	bindings.FuncCSPViolationReportBodyToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "CSPViolationReportBody.toJSON".
func (this CSPViolationReportBody) ToJSON() (ret js.Object) {
	bindings.CallCSPViolationReportBodyToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "CSPViolationReportBody.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSPViolationReportBody) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSPViolationReportBodyToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PropertyDefinition struct {
	// Name is "PropertyDefinition.name"
	//
	// Required
	Name js.String
	// Syntax is "PropertyDefinition.syntax"
	//
	// Optional, defaults to "*".
	Syntax js.String
	// Inherits is "PropertyDefinition.inherits"
	//
	// Required
	Inherits bool
	// InitialValue is "PropertyDefinition.initialValue"
	//
	// Optional
	InitialValue js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PropertyDefinition with all fields set.
func (p PropertyDefinition) FromRef(ref js.Ref) PropertyDefinition {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PropertyDefinition in the application heap.
func (p PropertyDefinition) New() js.Ref {
	return bindings.PropertyDefinitionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PropertyDefinition) UpdateFrom(ref js.Ref) {
	bindings.PropertyDefinitionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PropertyDefinition) Update(ref js.Ref) {
	bindings.PropertyDefinitionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PropertyDefinition) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.Syntax.Ref(),
		p.InitialValue.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Syntax = p.Syntax.FromRef(js.Undefined)
	p.InitialValue = p.InitialValue.FromRef(js.Undefined)
}

type CSSParserRule struct {
	ref js.Ref
}

func (this CSSParserRule) Once() CSSParserRule {
	this.ref.Once()
	return this
}

func (this CSSParserRule) Ref() js.Ref {
	return this.ref
}

func (this CSSParserRule) FromRef(ref js.Ref) CSSParserRule {
	this.ref = ref
	return this
}

func (this CSSParserRule) Free() {
	this.ref.Free()
}

type OneOf_String_ReadableStream struct {
	ref js.Ref
}

func (x OneOf_String_ReadableStream) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ReadableStream) Free() {
	x.ref.Free()
}

func (x OneOf_String_ReadableStream) FromRef(ref js.Ref) OneOf_String_ReadableStream {
	return OneOf_String_ReadableStream{
		ref: ref,
	}
}

func (x OneOf_String_ReadableStream) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ReadableStream) ReadableStream() ReadableStream {
	return ReadableStream{}.FromRef(x.ref)
}

type CSSStringSource = OneOf_String_ReadableStream

type CSSParserOptions struct {
	// AtRules is "CSSParserOptions.atRules"
	//
	// Optional
	AtRules js.Object

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CSSParserOptions with all fields set.
func (p CSSParserOptions) FromRef(ref js.Ref) CSSParserOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CSSParserOptions in the application heap.
func (p CSSParserOptions) New() js.Ref {
	return bindings.CSSParserOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CSSParserOptions) UpdateFrom(ref js.Ref) {
	bindings.CSSParserOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CSSParserOptions) Update(ref js.Ref) {
	bindings.CSSParserOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CSSParserOptions) FreeMembers(recursive bool) {
	js.Free(
		p.AtRules.Ref(),
	)
	p.AtRules = p.AtRules.FromRef(js.Undefined)
}

type CSSParserValue struct {
	ref js.Ref
}

func (this CSSParserValue) Once() CSSParserValue {
	this.ref.Once()
	return this
}

func (this CSSParserValue) Ref() js.Ref {
	return this.ref
}

func (this CSSParserValue) FromRef(ref js.Ref) CSSParserValue {
	this.ref = ref
	return this
}

func (this CSSParserValue) Free() {
	this.ref.Free()
}

func NewCSSParserDeclaration(name js.String, body js.Array[CSSParserRule]) (ret CSSParserDeclaration) {
	ret.ref = bindings.NewCSSParserDeclarationByCSSParserDeclaration(
		name.Ref(),
		body.Ref())
	return
}

func NewCSSParserDeclarationByCSSParserDeclaration1(name js.String) (ret CSSParserDeclaration) {
	ret.ref = bindings.NewCSSParserDeclarationByCSSParserDeclaration1(
		name.Ref())
	return
}

type CSSParserDeclaration struct {
	CSSParserRule
}

func (this CSSParserDeclaration) Once() CSSParserDeclaration {
	this.ref.Once()
	return this
}

func (this CSSParserDeclaration) Ref() js.Ref {
	return this.CSSParserRule.Ref()
}

func (this CSSParserDeclaration) FromRef(ref js.Ref) CSSParserDeclaration {
	this.CSSParserRule = this.CSSParserRule.FromRef(ref)
	return this
}

func (this CSSParserDeclaration) Free() {
	this.ref.Free()
}

// Name returns the value of property "CSSParserDeclaration.name".
//
// It returns ok=false if there is no such property.
func (this CSSParserDeclaration) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSParserDeclarationName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Body returns the value of property "CSSParserDeclaration.body".
//
// It returns ok=false if there is no such property.
func (this CSSParserDeclaration) Body() (ret js.FrozenArray[CSSParserValue], ok bool) {
	ok = js.True == bindings.GetCSSParserDeclarationBody(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncToString returns true if the method "CSSParserDeclaration.toString" exists.
func (this CSSParserDeclaration) HasFuncToString() bool {
	return js.True == bindings.HasFuncCSSParserDeclarationToString(
		this.ref,
	)
}

// FuncToString returns the method "CSSParserDeclaration.toString".
func (this CSSParserDeclaration) FuncToString() (fn js.Func[func() js.String]) {
	bindings.FuncCSSParserDeclarationToString(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToString calls the method "CSSParserDeclaration.toString".
func (this CSSParserDeclaration) ToString() (ret js.String) {
	bindings.CallCSSParserDeclarationToString(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "CSSParserDeclaration.toString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSParserDeclaration) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParserDeclarationToString(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OneOf_String_CSSStyleValue_CSSParserValue struct {
	ref js.Ref
}

func (x OneOf_String_CSSStyleValue_CSSParserValue) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_CSSStyleValue_CSSParserValue) Free() {
	x.ref.Free()
}

func (x OneOf_String_CSSStyleValue_CSSParserValue) FromRef(ref js.Ref) OneOf_String_CSSStyleValue_CSSParserValue {
	return OneOf_String_CSSStyleValue_CSSParserValue{
		ref: ref,
	}
}

func (x OneOf_String_CSSStyleValue_CSSParserValue) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_CSSStyleValue_CSSParserValue) CSSStyleValue() CSSStyleValue {
	return CSSStyleValue{}.FromRef(x.ref)
}

func (x OneOf_String_CSSStyleValue_CSSParserValue) CSSParserValue() CSSParserValue {
	return CSSParserValue{}.FromRef(x.ref)
}

type CSSToken = OneOf_String_CSSStyleValue_CSSParserValue

type WorkletOptions struct {
	// Credentials is "WorkletOptions.credentials"
	//
	// Optional, defaults to "same-origin".
	Credentials RequestCredentials

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WorkletOptions with all fields set.
func (p WorkletOptions) FromRef(ref js.Ref) WorkletOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WorkletOptions in the application heap.
func (p WorkletOptions) New() js.Ref {
	return bindings.WorkletOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WorkletOptions) UpdateFrom(ref js.Ref) {
	bindings.WorkletOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WorkletOptions) Update(ref js.Ref) {
	bindings.WorkletOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WorkletOptions) FreeMembers(recursive bool) {
}

type Worklet struct {
	ref js.Ref
}

func (this Worklet) Once() Worklet {
	this.ref.Once()
	return this
}

func (this Worklet) Ref() js.Ref {
	return this.ref
}

func (this Worklet) FromRef(ref js.Ref) Worklet {
	this.ref = ref
	return this
}

func (this Worklet) Free() {
	this.ref.Free()
}

// HasFuncAddModule returns true if the method "Worklet.addModule" exists.
func (this Worklet) HasFuncAddModule() bool {
	return js.True == bindings.HasFuncWorkletAddModule(
		this.ref,
	)
}

// FuncAddModule returns the method "Worklet.addModule".
func (this Worklet) FuncAddModule() (fn js.Func[func(moduleURL js.String, options WorkletOptions) js.Promise[js.Void]]) {
	bindings.FuncWorkletAddModule(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddModule calls the method "Worklet.addModule".
func (this Worklet) AddModule(moduleURL js.String, options WorkletOptions) (ret js.Promise[js.Void]) {
	bindings.CallWorkletAddModule(
		this.ref, js.Pointer(&ret),
		moduleURL.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryAddModule calls the method "Worklet.addModule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Worklet) TryAddModule(moduleURL js.String, options WorkletOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkletAddModule(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		moduleURL.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncAddModule1 returns true if the method "Worklet.addModule" exists.
func (this Worklet) HasFuncAddModule1() bool {
	return js.True == bindings.HasFuncWorkletAddModule1(
		this.ref,
	)
}

// FuncAddModule1 returns the method "Worklet.addModule".
func (this Worklet) FuncAddModule1() (fn js.Func[func(moduleURL js.String) js.Promise[js.Void]]) {
	bindings.FuncWorkletAddModule1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddModule1 calls the method "Worklet.addModule".
func (this Worklet) AddModule1(moduleURL js.String) (ret js.Promise[js.Void]) {
	bindings.CallWorkletAddModule1(
		this.ref, js.Pointer(&ret),
		moduleURL.Ref(),
	)

	return
}

// TryAddModule1 calls the method "Worklet.addModule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Worklet) TryAddModule1(moduleURL js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkletAddModule1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		moduleURL.Ref(),
	)

	return
}

type HighlightRegistry struct {
	ref js.Ref
}

func (this HighlightRegistry) Once() HighlightRegistry {
	this.ref.Once()
	return this
}

func (this HighlightRegistry) Ref() js.Ref {
	return this.ref
}

func (this HighlightRegistry) FromRef(ref js.Ref) HighlightRegistry {
	this.ref = ref
	return this
}

func (this HighlightRegistry) Free() {
	this.ref.Free()
}

type CSS struct{}

// ElementSources returns the value of property "CSS.elementSources".
//
// The returned bool will be false if there is no such property.
func (CSS) ElementSources() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetCSSElementSources(
		js.Pointer(&ret),
	)

	return
}

// AnimationWorklet returns the value of property "CSS.animationWorklet".
//
// The returned bool will be false if there is no such property.
func (CSS) AnimationWorklet() (ret Worklet, ok bool) {
	ok = js.True == bindings.GetCSSAnimationWorklet(
		js.Pointer(&ret),
	)

	return
}

// PaintWorklet returns the value of property "CSS.paintWorklet".
//
// The returned bool will be false if there is no such property.
func (CSS) PaintWorklet() (ret Worklet, ok bool) {
	ok = js.True == bindings.GetCSSPaintWorklet(
		js.Pointer(&ret),
	)

	return
}

// LayoutWorklet returns the value of property "CSS.layoutWorklet".
//
// The returned bool will be false if there is no such property.
func (CSS) LayoutWorklet() (ret Worklet, ok bool) {
	ok = js.True == bindings.GetCSSLayoutWorklet(
		js.Pointer(&ret),
	)

	return
}

// Highlights returns the value of property "CSS.highlights".
//
// The returned bool will be false if there is no such property.
func (CSS) Highlights() (ret HighlightRegistry, ok bool) {
	ok = js.True == bindings.GetCSSHighlights(
		js.Pointer(&ret),
	)

	return
}

// HasFuncEscape returns ture if the function "CSS.escape" exists.
func (CSS) HasFuncEscape() bool {
	return js.True == bindings.HasFuncCSSEscape()
}

// FuncEscape returns the function "CSS.escape".
func (CSS) FuncEscape() (fn js.Func[func(ident js.String) js.String]) {
	bindings.FuncCSSEscape(
		js.Pointer(&fn),
	)
	return
}

// Escape calls the function "CSS.escape".
func (CSS) Escape(ident js.String) (ret js.String) {
	bindings.CallCSSEscape(
		js.Pointer(&ret),
		ident.Ref(),
	)
	return
}

// TryEscape calls the function "CSS.escape"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryEscape(ident js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSEscape(
		js.Pointer(&ret), js.Pointer(&exception),
		ident.Ref(),
	)
	return
}

// HasFuncRegisterProperty returns ture if the function "CSS.registerProperty" exists.
func (CSS) HasFuncRegisterProperty() bool {
	return js.True == bindings.HasFuncCSSRegisterProperty()
}

// FuncRegisterProperty returns the function "CSS.registerProperty".
func (CSS) FuncRegisterProperty() (fn js.Func[func(definition PropertyDefinition)]) {
	bindings.FuncCSSRegisterProperty(
		js.Pointer(&fn),
	)
	return
}

// RegisterProperty calls the function "CSS.registerProperty".
func (CSS) RegisterProperty(definition PropertyDefinition) (ret js.Void) {
	bindings.CallCSSRegisterProperty(
		js.Pointer(&ret),
		js.Pointer(&definition),
	)
	return
}

// TryRegisterProperty calls the function "CSS.registerProperty"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryRegisterProperty(definition PropertyDefinition) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRegisterProperty(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&definition),
	)
	return
}

// HasFuncSupports returns ture if the function "CSS.supports" exists.
func (CSS) HasFuncSupports() bool {
	return js.True == bindings.HasFuncCSSSupports()
}

// FuncSupports returns the function "CSS.supports".
func (CSS) FuncSupports() (fn js.Func[func(property js.String, value js.String) bool]) {
	bindings.FuncCSSSupports(
		js.Pointer(&fn),
	)
	return
}

// Supports calls the function "CSS.supports".
func (CSS) Supports(property js.String, value js.String) (ret bool) {
	bindings.CallCSSSupports(
		js.Pointer(&ret),
		property.Ref(),
		value.Ref(),
	)
	return
}

// TrySupports calls the function "CSS.supports"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TrySupports(property js.String, value js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSupports(
		js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
		value.Ref(),
	)
	return
}

// HasFuncSupports1 returns ture if the function "CSS.supports" exists.
func (CSS) HasFuncSupports1() bool {
	return js.True == bindings.HasFuncCSSSupports1()
}

// FuncSupports1 returns the function "CSS.supports".
func (CSS) FuncSupports1() (fn js.Func[func(conditionText js.String) bool]) {
	bindings.FuncCSSSupports1(
		js.Pointer(&fn),
	)
	return
}

// Supports1 calls the function "CSS.supports".
func (CSS) Supports1(conditionText js.String) (ret bool) {
	bindings.CallCSSSupports1(
		js.Pointer(&ret),
		conditionText.Ref(),
	)
	return
}

// TrySupports1 calls the function "CSS.supports"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TrySupports1(conditionText js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSupports1(
		js.Pointer(&ret), js.Pointer(&exception),
		conditionText.Ref(),
	)
	return
}

// HasFuncNumber returns ture if the function "CSS.number" exists.
func (CSS) HasFuncNumber() bool {
	return js.True == bindings.HasFuncCSSNumber()
}

// FuncNumber returns the function "CSS.number".
func (CSS) FuncNumber() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSNumber(
		js.Pointer(&fn),
	)
	return
}

// Number calls the function "CSS.number".
func (CSS) Number(value float64) (ret CSSUnitValue) {
	bindings.CallCSSNumber(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryNumber calls the function "CSS.number"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryNumber(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSNumber(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncPercent returns ture if the function "CSS.percent" exists.
func (CSS) HasFuncPercent() bool {
	return js.True == bindings.HasFuncCSSPercent()
}

// FuncPercent returns the function "CSS.percent".
func (CSS) FuncPercent() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSPercent(
		js.Pointer(&fn),
	)
	return
}

// Percent calls the function "CSS.percent".
func (CSS) Percent(value float64) (ret CSSUnitValue) {
	bindings.CallCSSPercent(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryPercent calls the function "CSS.percent"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryPercent(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPercent(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncCap returns ture if the function "CSS.cap" exists.
func (CSS) HasFuncCap() bool {
	return js.True == bindings.HasFuncCSSCap()
}

// FuncCap returns the function "CSS.cap".
func (CSS) FuncCap() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSCap(
		js.Pointer(&fn),
	)
	return
}

// Cap calls the function "CSS.cap".
func (CSS) Cap(value float64) (ret CSSUnitValue) {
	bindings.CallCSSCap(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryCap calls the function "CSS.cap"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryCap(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCap(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncCh returns ture if the function "CSS.ch" exists.
func (CSS) HasFuncCh() bool {
	return js.True == bindings.HasFuncCSSCh()
}

// FuncCh returns the function "CSS.ch".
func (CSS) FuncCh() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSCh(
		js.Pointer(&fn),
	)
	return
}

// Ch calls the function "CSS.ch".
func (CSS) Ch(value float64) (ret CSSUnitValue) {
	bindings.CallCSSCh(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryCh calls the function "CSS.ch"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryCh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncEm returns ture if the function "CSS.em" exists.
func (CSS) HasFuncEm() bool {
	return js.True == bindings.HasFuncCSSEm()
}

// FuncEm returns the function "CSS.em".
func (CSS) FuncEm() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSEm(
		js.Pointer(&fn),
	)
	return
}

// Em calls the function "CSS.em".
func (CSS) Em(value float64) (ret CSSUnitValue) {
	bindings.CallCSSEm(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryEm calls the function "CSS.em"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryEm(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSEm(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncEx returns ture if the function "CSS.ex" exists.
func (CSS) HasFuncEx() bool {
	return js.True == bindings.HasFuncCSSEx()
}

// FuncEx returns the function "CSS.ex".
func (CSS) FuncEx() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSEx(
		js.Pointer(&fn),
	)
	return
}

// Ex calls the function "CSS.ex".
func (CSS) Ex(value float64) (ret CSSUnitValue) {
	bindings.CallCSSEx(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryEx calls the function "CSS.ex"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryEx(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSEx(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncIc returns ture if the function "CSS.ic" exists.
func (CSS) HasFuncIc() bool {
	return js.True == bindings.HasFuncCSSIc()
}

// FuncIc returns the function "CSS.ic".
func (CSS) FuncIc() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSIc(
		js.Pointer(&fn),
	)
	return
}

// Ic calls the function "CSS.ic".
func (CSS) Ic(value float64) (ret CSSUnitValue) {
	bindings.CallCSSIc(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryIc calls the function "CSS.ic"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryIc(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSIc(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncLh returns ture if the function "CSS.lh" exists.
func (CSS) HasFuncLh() bool {
	return js.True == bindings.HasFuncCSSLh()
}

// FuncLh returns the function "CSS.lh".
func (CSS) FuncLh() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSLh(
		js.Pointer(&fn),
	)
	return
}

// Lh calls the function "CSS.lh".
func (CSS) Lh(value float64) (ret CSSUnitValue) {
	bindings.CallCSSLh(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryLh calls the function "CSS.lh"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryLh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSLh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncRcap returns ture if the function "CSS.rcap" exists.
func (CSS) HasFuncRcap() bool {
	return js.True == bindings.HasFuncCSSRcap()
}

// FuncRcap returns the function "CSS.rcap".
func (CSS) FuncRcap() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSRcap(
		js.Pointer(&fn),
	)
	return
}

// Rcap calls the function "CSS.rcap".
func (CSS) Rcap(value float64) (ret CSSUnitValue) {
	bindings.CallCSSRcap(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryRcap calls the function "CSS.rcap"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryRcap(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRcap(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncRch returns ture if the function "CSS.rch" exists.
func (CSS) HasFuncRch() bool {
	return js.True == bindings.HasFuncCSSRch()
}

// FuncRch returns the function "CSS.rch".
func (CSS) FuncRch() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSRch(
		js.Pointer(&fn),
	)
	return
}

// Rch calls the function "CSS.rch".
func (CSS) Rch(value float64) (ret CSSUnitValue) {
	bindings.CallCSSRch(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryRch calls the function "CSS.rch"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryRch(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRch(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncRem returns ture if the function "CSS.rem" exists.
func (CSS) HasFuncRem() bool {
	return js.True == bindings.HasFuncCSSRem()
}

// FuncRem returns the function "CSS.rem".
func (CSS) FuncRem() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSRem(
		js.Pointer(&fn),
	)
	return
}

// Rem calls the function "CSS.rem".
func (CSS) Rem(value float64) (ret CSSUnitValue) {
	bindings.CallCSSRem(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryRem calls the function "CSS.rem"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryRem(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRem(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncRex returns ture if the function "CSS.rex" exists.
func (CSS) HasFuncRex() bool {
	return js.True == bindings.HasFuncCSSRex()
}

// FuncRex returns the function "CSS.rex".
func (CSS) FuncRex() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSRex(
		js.Pointer(&fn),
	)
	return
}

// Rex calls the function "CSS.rex".
func (CSS) Rex(value float64) (ret CSSUnitValue) {
	bindings.CallCSSRex(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryRex calls the function "CSS.rex"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryRex(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRex(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncRic returns ture if the function "CSS.ric" exists.
func (CSS) HasFuncRic() bool {
	return js.True == bindings.HasFuncCSSRic()
}

// FuncRic returns the function "CSS.ric".
func (CSS) FuncRic() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSRic(
		js.Pointer(&fn),
	)
	return
}

// Ric calls the function "CSS.ric".
func (CSS) Ric(value float64) (ret CSSUnitValue) {
	bindings.CallCSSRic(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryRic calls the function "CSS.ric"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryRic(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRic(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncRlh returns ture if the function "CSS.rlh" exists.
func (CSS) HasFuncRlh() bool {
	return js.True == bindings.HasFuncCSSRlh()
}

// FuncRlh returns the function "CSS.rlh".
func (CSS) FuncRlh() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSRlh(
		js.Pointer(&fn),
	)
	return
}

// Rlh calls the function "CSS.rlh".
func (CSS) Rlh(value float64) (ret CSSUnitValue) {
	bindings.CallCSSRlh(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryRlh calls the function "CSS.rlh"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryRlh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRlh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncVw returns ture if the function "CSS.vw" exists.
func (CSS) HasFuncVw() bool {
	return js.True == bindings.HasFuncCSSVw()
}

// FuncVw returns the function "CSS.vw".
func (CSS) FuncVw() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSVw(
		js.Pointer(&fn),
	)
	return
}

// Vw calls the function "CSS.vw".
func (CSS) Vw(value float64) (ret CSSUnitValue) {
	bindings.CallCSSVw(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryVw calls the function "CSS.vw"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryVw(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSVw(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncVh returns ture if the function "CSS.vh" exists.
func (CSS) HasFuncVh() bool {
	return js.True == bindings.HasFuncCSSVh()
}

// FuncVh returns the function "CSS.vh".
func (CSS) FuncVh() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSVh(
		js.Pointer(&fn),
	)
	return
}

// Vh calls the function "CSS.vh".
func (CSS) Vh(value float64) (ret CSSUnitValue) {
	bindings.CallCSSVh(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryVh calls the function "CSS.vh"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryVh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSVh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncVi returns ture if the function "CSS.vi" exists.
func (CSS) HasFuncVi() bool {
	return js.True == bindings.HasFuncCSSVi()
}

// FuncVi returns the function "CSS.vi".
func (CSS) FuncVi() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSVi(
		js.Pointer(&fn),
	)
	return
}

// Vi calls the function "CSS.vi".
func (CSS) Vi(value float64) (ret CSSUnitValue) {
	bindings.CallCSSVi(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryVi calls the function "CSS.vi"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryVi(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSVi(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncVb returns ture if the function "CSS.vb" exists.
func (CSS) HasFuncVb() bool {
	return js.True == bindings.HasFuncCSSVb()
}

// FuncVb returns the function "CSS.vb".
func (CSS) FuncVb() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSVb(
		js.Pointer(&fn),
	)
	return
}

// Vb calls the function "CSS.vb".
func (CSS) Vb(value float64) (ret CSSUnitValue) {
	bindings.CallCSSVb(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryVb calls the function "CSS.vb"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryVb(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSVb(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncVmin returns ture if the function "CSS.vmin" exists.
func (CSS) HasFuncVmin() bool {
	return js.True == bindings.HasFuncCSSVmin()
}

// FuncVmin returns the function "CSS.vmin".
func (CSS) FuncVmin() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSVmin(
		js.Pointer(&fn),
	)
	return
}

// Vmin calls the function "CSS.vmin".
func (CSS) Vmin(value float64) (ret CSSUnitValue) {
	bindings.CallCSSVmin(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryVmin calls the function "CSS.vmin"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryVmin(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSVmin(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncVmax returns ture if the function "CSS.vmax" exists.
func (CSS) HasFuncVmax() bool {
	return js.True == bindings.HasFuncCSSVmax()
}

// FuncVmax returns the function "CSS.vmax".
func (CSS) FuncVmax() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSVmax(
		js.Pointer(&fn),
	)
	return
}

// Vmax calls the function "CSS.vmax".
func (CSS) Vmax(value float64) (ret CSSUnitValue) {
	bindings.CallCSSVmax(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryVmax calls the function "CSS.vmax"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryVmax(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSVmax(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncSvw returns ture if the function "CSS.svw" exists.
func (CSS) HasFuncSvw() bool {
	return js.True == bindings.HasFuncCSSSvw()
}

// FuncSvw returns the function "CSS.svw".
func (CSS) FuncSvw() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSSvw(
		js.Pointer(&fn),
	)
	return
}

// Svw calls the function "CSS.svw".
func (CSS) Svw(value float64) (ret CSSUnitValue) {
	bindings.CallCSSSvw(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TrySvw calls the function "CSS.svw"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TrySvw(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSvw(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncSvh returns ture if the function "CSS.svh" exists.
func (CSS) HasFuncSvh() bool {
	return js.True == bindings.HasFuncCSSSvh()
}

// FuncSvh returns the function "CSS.svh".
func (CSS) FuncSvh() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSSvh(
		js.Pointer(&fn),
	)
	return
}

// Svh calls the function "CSS.svh".
func (CSS) Svh(value float64) (ret CSSUnitValue) {
	bindings.CallCSSSvh(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TrySvh calls the function "CSS.svh"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TrySvh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSvh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncSvi returns ture if the function "CSS.svi" exists.
func (CSS) HasFuncSvi() bool {
	return js.True == bindings.HasFuncCSSSvi()
}

// FuncSvi returns the function "CSS.svi".
func (CSS) FuncSvi() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSSvi(
		js.Pointer(&fn),
	)
	return
}

// Svi calls the function "CSS.svi".
func (CSS) Svi(value float64) (ret CSSUnitValue) {
	bindings.CallCSSSvi(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TrySvi calls the function "CSS.svi"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TrySvi(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSvi(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncSvb returns ture if the function "CSS.svb" exists.
func (CSS) HasFuncSvb() bool {
	return js.True == bindings.HasFuncCSSSvb()
}

// FuncSvb returns the function "CSS.svb".
func (CSS) FuncSvb() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSSvb(
		js.Pointer(&fn),
	)
	return
}

// Svb calls the function "CSS.svb".
func (CSS) Svb(value float64) (ret CSSUnitValue) {
	bindings.CallCSSSvb(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TrySvb calls the function "CSS.svb"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TrySvb(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSvb(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncSvmin returns ture if the function "CSS.svmin" exists.
func (CSS) HasFuncSvmin() bool {
	return js.True == bindings.HasFuncCSSSvmin()
}

// FuncSvmin returns the function "CSS.svmin".
func (CSS) FuncSvmin() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSSvmin(
		js.Pointer(&fn),
	)
	return
}

// Svmin calls the function "CSS.svmin".
func (CSS) Svmin(value float64) (ret CSSUnitValue) {
	bindings.CallCSSSvmin(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TrySvmin calls the function "CSS.svmin"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TrySvmin(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSvmin(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncSvmax returns ture if the function "CSS.svmax" exists.
func (CSS) HasFuncSvmax() bool {
	return js.True == bindings.HasFuncCSSSvmax()
}

// FuncSvmax returns the function "CSS.svmax".
func (CSS) FuncSvmax() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSSvmax(
		js.Pointer(&fn),
	)
	return
}

// Svmax calls the function "CSS.svmax".
func (CSS) Svmax(value float64) (ret CSSUnitValue) {
	bindings.CallCSSSvmax(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TrySvmax calls the function "CSS.svmax"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TrySvmax(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSvmax(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncLvw returns ture if the function "CSS.lvw" exists.
func (CSS) HasFuncLvw() bool {
	return js.True == bindings.HasFuncCSSLvw()
}

// FuncLvw returns the function "CSS.lvw".
func (CSS) FuncLvw() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSLvw(
		js.Pointer(&fn),
	)
	return
}

// Lvw calls the function "CSS.lvw".
func (CSS) Lvw(value float64) (ret CSSUnitValue) {
	bindings.CallCSSLvw(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryLvw calls the function "CSS.lvw"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryLvw(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSLvw(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncLvh returns ture if the function "CSS.lvh" exists.
func (CSS) HasFuncLvh() bool {
	return js.True == bindings.HasFuncCSSLvh()
}

// FuncLvh returns the function "CSS.lvh".
func (CSS) FuncLvh() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSLvh(
		js.Pointer(&fn),
	)
	return
}

// Lvh calls the function "CSS.lvh".
func (CSS) Lvh(value float64) (ret CSSUnitValue) {
	bindings.CallCSSLvh(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryLvh calls the function "CSS.lvh"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryLvh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSLvh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncLvi returns ture if the function "CSS.lvi" exists.
func (CSS) HasFuncLvi() bool {
	return js.True == bindings.HasFuncCSSLvi()
}

// FuncLvi returns the function "CSS.lvi".
func (CSS) FuncLvi() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSLvi(
		js.Pointer(&fn),
	)
	return
}

// Lvi calls the function "CSS.lvi".
func (CSS) Lvi(value float64) (ret CSSUnitValue) {
	bindings.CallCSSLvi(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryLvi calls the function "CSS.lvi"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryLvi(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSLvi(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncLvb returns ture if the function "CSS.lvb" exists.
func (CSS) HasFuncLvb() bool {
	return js.True == bindings.HasFuncCSSLvb()
}

// FuncLvb returns the function "CSS.lvb".
func (CSS) FuncLvb() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSLvb(
		js.Pointer(&fn),
	)
	return
}

// Lvb calls the function "CSS.lvb".
func (CSS) Lvb(value float64) (ret CSSUnitValue) {
	bindings.CallCSSLvb(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryLvb calls the function "CSS.lvb"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryLvb(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSLvb(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncLvmin returns ture if the function "CSS.lvmin" exists.
func (CSS) HasFuncLvmin() bool {
	return js.True == bindings.HasFuncCSSLvmin()
}

// FuncLvmin returns the function "CSS.lvmin".
func (CSS) FuncLvmin() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSLvmin(
		js.Pointer(&fn),
	)
	return
}

// Lvmin calls the function "CSS.lvmin".
func (CSS) Lvmin(value float64) (ret CSSUnitValue) {
	bindings.CallCSSLvmin(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryLvmin calls the function "CSS.lvmin"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryLvmin(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSLvmin(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncLvmax returns ture if the function "CSS.lvmax" exists.
func (CSS) HasFuncLvmax() bool {
	return js.True == bindings.HasFuncCSSLvmax()
}

// FuncLvmax returns the function "CSS.lvmax".
func (CSS) FuncLvmax() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSLvmax(
		js.Pointer(&fn),
	)
	return
}

// Lvmax calls the function "CSS.lvmax".
func (CSS) Lvmax(value float64) (ret CSSUnitValue) {
	bindings.CallCSSLvmax(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryLvmax calls the function "CSS.lvmax"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryLvmax(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSLvmax(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncDvw returns ture if the function "CSS.dvw" exists.
func (CSS) HasFuncDvw() bool {
	return js.True == bindings.HasFuncCSSDvw()
}

// FuncDvw returns the function "CSS.dvw".
func (CSS) FuncDvw() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSDvw(
		js.Pointer(&fn),
	)
	return
}

// Dvw calls the function "CSS.dvw".
func (CSS) Dvw(value float64) (ret CSSUnitValue) {
	bindings.CallCSSDvw(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryDvw calls the function "CSS.dvw"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryDvw(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDvw(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncDvh returns ture if the function "CSS.dvh" exists.
func (CSS) HasFuncDvh() bool {
	return js.True == bindings.HasFuncCSSDvh()
}

// FuncDvh returns the function "CSS.dvh".
func (CSS) FuncDvh() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSDvh(
		js.Pointer(&fn),
	)
	return
}

// Dvh calls the function "CSS.dvh".
func (CSS) Dvh(value float64) (ret CSSUnitValue) {
	bindings.CallCSSDvh(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryDvh calls the function "CSS.dvh"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryDvh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDvh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncDvi returns ture if the function "CSS.dvi" exists.
func (CSS) HasFuncDvi() bool {
	return js.True == bindings.HasFuncCSSDvi()
}

// FuncDvi returns the function "CSS.dvi".
func (CSS) FuncDvi() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSDvi(
		js.Pointer(&fn),
	)
	return
}

// Dvi calls the function "CSS.dvi".
func (CSS) Dvi(value float64) (ret CSSUnitValue) {
	bindings.CallCSSDvi(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryDvi calls the function "CSS.dvi"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryDvi(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDvi(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncDvb returns ture if the function "CSS.dvb" exists.
func (CSS) HasFuncDvb() bool {
	return js.True == bindings.HasFuncCSSDvb()
}

// FuncDvb returns the function "CSS.dvb".
func (CSS) FuncDvb() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSDvb(
		js.Pointer(&fn),
	)
	return
}

// Dvb calls the function "CSS.dvb".
func (CSS) Dvb(value float64) (ret CSSUnitValue) {
	bindings.CallCSSDvb(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryDvb calls the function "CSS.dvb"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryDvb(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDvb(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncDvmin returns ture if the function "CSS.dvmin" exists.
func (CSS) HasFuncDvmin() bool {
	return js.True == bindings.HasFuncCSSDvmin()
}

// FuncDvmin returns the function "CSS.dvmin".
func (CSS) FuncDvmin() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSDvmin(
		js.Pointer(&fn),
	)
	return
}

// Dvmin calls the function "CSS.dvmin".
func (CSS) Dvmin(value float64) (ret CSSUnitValue) {
	bindings.CallCSSDvmin(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryDvmin calls the function "CSS.dvmin"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryDvmin(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDvmin(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncDvmax returns ture if the function "CSS.dvmax" exists.
func (CSS) HasFuncDvmax() bool {
	return js.True == bindings.HasFuncCSSDvmax()
}

// FuncDvmax returns the function "CSS.dvmax".
func (CSS) FuncDvmax() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSDvmax(
		js.Pointer(&fn),
	)
	return
}

// Dvmax calls the function "CSS.dvmax".
func (CSS) Dvmax(value float64) (ret CSSUnitValue) {
	bindings.CallCSSDvmax(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryDvmax calls the function "CSS.dvmax"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryDvmax(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDvmax(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncCqw returns ture if the function "CSS.cqw" exists.
func (CSS) HasFuncCqw() bool {
	return js.True == bindings.HasFuncCSSCqw()
}

// FuncCqw returns the function "CSS.cqw".
func (CSS) FuncCqw() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSCqw(
		js.Pointer(&fn),
	)
	return
}

// Cqw calls the function "CSS.cqw".
func (CSS) Cqw(value float64) (ret CSSUnitValue) {
	bindings.CallCSSCqw(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryCqw calls the function "CSS.cqw"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryCqw(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCqw(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncCqh returns ture if the function "CSS.cqh" exists.
func (CSS) HasFuncCqh() bool {
	return js.True == bindings.HasFuncCSSCqh()
}

// FuncCqh returns the function "CSS.cqh".
func (CSS) FuncCqh() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSCqh(
		js.Pointer(&fn),
	)
	return
}

// Cqh calls the function "CSS.cqh".
func (CSS) Cqh(value float64) (ret CSSUnitValue) {
	bindings.CallCSSCqh(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryCqh calls the function "CSS.cqh"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryCqh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCqh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncCqi returns ture if the function "CSS.cqi" exists.
func (CSS) HasFuncCqi() bool {
	return js.True == bindings.HasFuncCSSCqi()
}

// FuncCqi returns the function "CSS.cqi".
func (CSS) FuncCqi() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSCqi(
		js.Pointer(&fn),
	)
	return
}

// Cqi calls the function "CSS.cqi".
func (CSS) Cqi(value float64) (ret CSSUnitValue) {
	bindings.CallCSSCqi(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryCqi calls the function "CSS.cqi"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryCqi(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCqi(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncCqb returns ture if the function "CSS.cqb" exists.
func (CSS) HasFuncCqb() bool {
	return js.True == bindings.HasFuncCSSCqb()
}

// FuncCqb returns the function "CSS.cqb".
func (CSS) FuncCqb() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSCqb(
		js.Pointer(&fn),
	)
	return
}

// Cqb calls the function "CSS.cqb".
func (CSS) Cqb(value float64) (ret CSSUnitValue) {
	bindings.CallCSSCqb(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryCqb calls the function "CSS.cqb"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryCqb(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCqb(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncCqmin returns ture if the function "CSS.cqmin" exists.
func (CSS) HasFuncCqmin() bool {
	return js.True == bindings.HasFuncCSSCqmin()
}

// FuncCqmin returns the function "CSS.cqmin".
func (CSS) FuncCqmin() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSCqmin(
		js.Pointer(&fn),
	)
	return
}

// Cqmin calls the function "CSS.cqmin".
func (CSS) Cqmin(value float64) (ret CSSUnitValue) {
	bindings.CallCSSCqmin(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryCqmin calls the function "CSS.cqmin"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryCqmin(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCqmin(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncCqmax returns ture if the function "CSS.cqmax" exists.
func (CSS) HasFuncCqmax() bool {
	return js.True == bindings.HasFuncCSSCqmax()
}

// FuncCqmax returns the function "CSS.cqmax".
func (CSS) FuncCqmax() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSCqmax(
		js.Pointer(&fn),
	)
	return
}

// Cqmax calls the function "CSS.cqmax".
func (CSS) Cqmax(value float64) (ret CSSUnitValue) {
	bindings.CallCSSCqmax(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryCqmax calls the function "CSS.cqmax"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryCqmax(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCqmax(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncCm returns ture if the function "CSS.cm" exists.
func (CSS) HasFuncCm() bool {
	return js.True == bindings.HasFuncCSSCm()
}

// FuncCm returns the function "CSS.cm".
func (CSS) FuncCm() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSCm(
		js.Pointer(&fn),
	)
	return
}

// Cm calls the function "CSS.cm".
func (CSS) Cm(value float64) (ret CSSUnitValue) {
	bindings.CallCSSCm(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryCm calls the function "CSS.cm"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryCm(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCm(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncMm returns ture if the function "CSS.mm" exists.
func (CSS) HasFuncMm() bool {
	return js.True == bindings.HasFuncCSSMm()
}

// FuncMm returns the function "CSS.mm".
func (CSS) FuncMm() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSMm(
		js.Pointer(&fn),
	)
	return
}

// Mm calls the function "CSS.mm".
func (CSS) Mm(value float64) (ret CSSUnitValue) {
	bindings.CallCSSMm(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryMm calls the function "CSS.mm"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryMm(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSMm(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncQ returns ture if the function "CSS.Q" exists.
func (CSS) HasFuncQ() bool {
	return js.True == bindings.HasFuncCSSQ()
}

// FuncQ returns the function "CSS.Q".
func (CSS) FuncQ() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSQ(
		js.Pointer(&fn),
	)
	return
}

// Q calls the function "CSS.Q".
func (CSS) Q(value float64) (ret CSSUnitValue) {
	bindings.CallCSSQ(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryQ calls the function "CSS.Q"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryQ(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSQ(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncIn returns ture if the function "CSS.in" exists.
func (CSS) HasFuncIn() bool {
	return js.True == bindings.HasFuncCSSIn()
}

// FuncIn returns the function "CSS.in".
func (CSS) FuncIn() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSIn(
		js.Pointer(&fn),
	)
	return
}

// In calls the function "CSS.in".
func (CSS) In(value float64) (ret CSSUnitValue) {
	bindings.CallCSSIn(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryIn calls the function "CSS.in"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryIn(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSIn(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncPt returns ture if the function "CSS.pt" exists.
func (CSS) HasFuncPt() bool {
	return js.True == bindings.HasFuncCSSPt()
}

// FuncPt returns the function "CSS.pt".
func (CSS) FuncPt() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSPt(
		js.Pointer(&fn),
	)
	return
}

// Pt calls the function "CSS.pt".
func (CSS) Pt(value float64) (ret CSSUnitValue) {
	bindings.CallCSSPt(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryPt calls the function "CSS.pt"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryPt(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPt(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncPc returns ture if the function "CSS.pc" exists.
func (CSS) HasFuncPc() bool {
	return js.True == bindings.HasFuncCSSPc()
}

// FuncPc returns the function "CSS.pc".
func (CSS) FuncPc() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSPc(
		js.Pointer(&fn),
	)
	return
}

// Pc calls the function "CSS.pc".
func (CSS) Pc(value float64) (ret CSSUnitValue) {
	bindings.CallCSSPc(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryPc calls the function "CSS.pc"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryPc(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPc(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncPx returns ture if the function "CSS.px" exists.
func (CSS) HasFuncPx() bool {
	return js.True == bindings.HasFuncCSSPx()
}

// FuncPx returns the function "CSS.px".
func (CSS) FuncPx() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSPx(
		js.Pointer(&fn),
	)
	return
}

// Px calls the function "CSS.px".
func (CSS) Px(value float64) (ret CSSUnitValue) {
	bindings.CallCSSPx(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryPx calls the function "CSS.px"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryPx(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPx(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncDeg returns ture if the function "CSS.deg" exists.
func (CSS) HasFuncDeg() bool {
	return js.True == bindings.HasFuncCSSDeg()
}

// FuncDeg returns the function "CSS.deg".
func (CSS) FuncDeg() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSDeg(
		js.Pointer(&fn),
	)
	return
}

// Deg calls the function "CSS.deg".
func (CSS) Deg(value float64) (ret CSSUnitValue) {
	bindings.CallCSSDeg(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryDeg calls the function "CSS.deg"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryDeg(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDeg(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncGrad returns ture if the function "CSS.grad" exists.
func (CSS) HasFuncGrad() bool {
	return js.True == bindings.HasFuncCSSGrad()
}

// FuncGrad returns the function "CSS.grad".
func (CSS) FuncGrad() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSGrad(
		js.Pointer(&fn),
	)
	return
}

// Grad calls the function "CSS.grad".
func (CSS) Grad(value float64) (ret CSSUnitValue) {
	bindings.CallCSSGrad(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryGrad calls the function "CSS.grad"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryGrad(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSGrad(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncRad returns ture if the function "CSS.rad" exists.
func (CSS) HasFuncRad() bool {
	return js.True == bindings.HasFuncCSSRad()
}

// FuncRad returns the function "CSS.rad".
func (CSS) FuncRad() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSRad(
		js.Pointer(&fn),
	)
	return
}

// Rad calls the function "CSS.rad".
func (CSS) Rad(value float64) (ret CSSUnitValue) {
	bindings.CallCSSRad(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryRad calls the function "CSS.rad"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryRad(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRad(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncTurn returns ture if the function "CSS.turn" exists.
func (CSS) HasFuncTurn() bool {
	return js.True == bindings.HasFuncCSSTurn()
}

// FuncTurn returns the function "CSS.turn".
func (CSS) FuncTurn() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSTurn(
		js.Pointer(&fn),
	)
	return
}

// Turn calls the function "CSS.turn".
func (CSS) Turn(value float64) (ret CSSUnitValue) {
	bindings.CallCSSTurn(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryTurn calls the function "CSS.turn"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryTurn(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSTurn(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncS returns ture if the function "CSS.s" exists.
func (CSS) HasFuncS() bool {
	return js.True == bindings.HasFuncCSSS()
}

// FuncS returns the function "CSS.s".
func (CSS) FuncS() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSS(
		js.Pointer(&fn),
	)
	return
}

// S calls the function "CSS.s".
func (CSS) S(value float64) (ret CSSUnitValue) {
	bindings.CallCSSS(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryS calls the function "CSS.s"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryS(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSS(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncMs returns ture if the function "CSS.ms" exists.
func (CSS) HasFuncMs() bool {
	return js.True == bindings.HasFuncCSSMs()
}

// FuncMs returns the function "CSS.ms".
func (CSS) FuncMs() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSMs(
		js.Pointer(&fn),
	)
	return
}

// Ms calls the function "CSS.ms".
func (CSS) Ms(value float64) (ret CSSUnitValue) {
	bindings.CallCSSMs(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryMs calls the function "CSS.ms"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryMs(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSMs(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncHz returns ture if the function "CSS.Hz" exists.
func (CSS) HasFuncHz() bool {
	return js.True == bindings.HasFuncCSSHz()
}

// FuncHz returns the function "CSS.Hz".
func (CSS) FuncHz() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSHz(
		js.Pointer(&fn),
	)
	return
}

// Hz calls the function "CSS.Hz".
func (CSS) Hz(value float64) (ret CSSUnitValue) {
	bindings.CallCSSHz(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryHz calls the function "CSS.Hz"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryHz(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSHz(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncKHz returns ture if the function "CSS.kHz" exists.
func (CSS) HasFuncKHz() bool {
	return js.True == bindings.HasFuncCSSKHz()
}

// FuncKHz returns the function "CSS.kHz".
func (CSS) FuncKHz() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSKHz(
		js.Pointer(&fn),
	)
	return
}

// KHz calls the function "CSS.kHz".
func (CSS) KHz(value float64) (ret CSSUnitValue) {
	bindings.CallCSSKHz(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryKHz calls the function "CSS.kHz"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryKHz(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSKHz(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncDpi returns ture if the function "CSS.dpi" exists.
func (CSS) HasFuncDpi() bool {
	return js.True == bindings.HasFuncCSSDpi()
}

// FuncDpi returns the function "CSS.dpi".
func (CSS) FuncDpi() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSDpi(
		js.Pointer(&fn),
	)
	return
}

// Dpi calls the function "CSS.dpi".
func (CSS) Dpi(value float64) (ret CSSUnitValue) {
	bindings.CallCSSDpi(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryDpi calls the function "CSS.dpi"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryDpi(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDpi(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncDpcm returns ture if the function "CSS.dpcm" exists.
func (CSS) HasFuncDpcm() bool {
	return js.True == bindings.HasFuncCSSDpcm()
}

// FuncDpcm returns the function "CSS.dpcm".
func (CSS) FuncDpcm() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSDpcm(
		js.Pointer(&fn),
	)
	return
}

// Dpcm calls the function "CSS.dpcm".
func (CSS) Dpcm(value float64) (ret CSSUnitValue) {
	bindings.CallCSSDpcm(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryDpcm calls the function "CSS.dpcm"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryDpcm(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDpcm(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncDppx returns ture if the function "CSS.dppx" exists.
func (CSS) HasFuncDppx() bool {
	return js.True == bindings.HasFuncCSSDppx()
}

// FuncDppx returns the function "CSS.dppx".
func (CSS) FuncDppx() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSDppx(
		js.Pointer(&fn),
	)
	return
}

// Dppx calls the function "CSS.dppx".
func (CSS) Dppx(value float64) (ret CSSUnitValue) {
	bindings.CallCSSDppx(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryDppx calls the function "CSS.dppx"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryDppx(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDppx(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncFr returns ture if the function "CSS.fr" exists.
func (CSS) HasFuncFr() bool {
	return js.True == bindings.HasFuncCSSFr()
}

// FuncFr returns the function "CSS.fr".
func (CSS) FuncFr() (fn js.Func[func(value float64) CSSUnitValue]) {
	bindings.FuncCSSFr(
		js.Pointer(&fn),
	)
	return
}

// Fr calls the function "CSS.fr".
func (CSS) Fr(value float64) (ret CSSUnitValue) {
	bindings.CallCSSFr(
		js.Pointer(&ret),
		float64(value),
	)
	return
}

// TryFr calls the function "CSS.fr"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryFr(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSFr(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFuncParseStylesheet returns ture if the function "CSS.parseStylesheet" exists.
func (CSS) HasFuncParseStylesheet() bool {
	return js.True == bindings.HasFuncCSSParseStylesheet()
}

// FuncParseStylesheet returns the function "CSS.parseStylesheet".
func (CSS) FuncParseStylesheet() (fn js.Func[func(css CSSStringSource, options CSSParserOptions) js.Promise[js.Array[CSSParserRule]]]) {
	bindings.FuncCSSParseStylesheet(
		js.Pointer(&fn),
	)
	return
}

// ParseStylesheet calls the function "CSS.parseStylesheet".
func (CSS) ParseStylesheet(css CSSStringSource, options CSSParserOptions) (ret js.Promise[js.Array[CSSParserRule]]) {
	bindings.CallCSSParseStylesheet(
		js.Pointer(&ret),
		css.Ref(),
		js.Pointer(&options),
	)
	return
}

// TryParseStylesheet calls the function "CSS.parseStylesheet"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryParseStylesheet(css CSSStringSource, options CSSParserOptions) (ret js.Promise[js.Array[CSSParserRule]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseStylesheet(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
		js.Pointer(&options),
	)
	return
}

// HasFuncParseStylesheet1 returns ture if the function "CSS.parseStylesheet" exists.
func (CSS) HasFuncParseStylesheet1() bool {
	return js.True == bindings.HasFuncCSSParseStylesheet1()
}

// FuncParseStylesheet1 returns the function "CSS.parseStylesheet".
func (CSS) FuncParseStylesheet1() (fn js.Func[func(css CSSStringSource) js.Promise[js.Array[CSSParserRule]]]) {
	bindings.FuncCSSParseStylesheet1(
		js.Pointer(&fn),
	)
	return
}

// ParseStylesheet1 calls the function "CSS.parseStylesheet".
func (CSS) ParseStylesheet1(css CSSStringSource) (ret js.Promise[js.Array[CSSParserRule]]) {
	bindings.CallCSSParseStylesheet1(
		js.Pointer(&ret),
		css.Ref(),
	)
	return
}

// TryParseStylesheet1 calls the function "CSS.parseStylesheet"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryParseStylesheet1(css CSSStringSource) (ret js.Promise[js.Array[CSSParserRule]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseStylesheet1(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
	)
	return
}

// HasFuncParseRuleList returns ture if the function "CSS.parseRuleList" exists.
func (CSS) HasFuncParseRuleList() bool {
	return js.True == bindings.HasFuncCSSParseRuleList()
}

// FuncParseRuleList returns the function "CSS.parseRuleList".
func (CSS) FuncParseRuleList() (fn js.Func[func(css CSSStringSource, options CSSParserOptions) js.Promise[js.Array[CSSParserRule]]]) {
	bindings.FuncCSSParseRuleList(
		js.Pointer(&fn),
	)
	return
}

// ParseRuleList calls the function "CSS.parseRuleList".
func (CSS) ParseRuleList(css CSSStringSource, options CSSParserOptions) (ret js.Promise[js.Array[CSSParserRule]]) {
	bindings.CallCSSParseRuleList(
		js.Pointer(&ret),
		css.Ref(),
		js.Pointer(&options),
	)
	return
}

// TryParseRuleList calls the function "CSS.parseRuleList"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryParseRuleList(css CSSStringSource, options CSSParserOptions) (ret js.Promise[js.Array[CSSParserRule]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseRuleList(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
		js.Pointer(&options),
	)
	return
}

// HasFuncParseRuleList1 returns ture if the function "CSS.parseRuleList" exists.
func (CSS) HasFuncParseRuleList1() bool {
	return js.True == bindings.HasFuncCSSParseRuleList1()
}

// FuncParseRuleList1 returns the function "CSS.parseRuleList".
func (CSS) FuncParseRuleList1() (fn js.Func[func(css CSSStringSource) js.Promise[js.Array[CSSParserRule]]]) {
	bindings.FuncCSSParseRuleList1(
		js.Pointer(&fn),
	)
	return
}

// ParseRuleList1 calls the function "CSS.parseRuleList".
func (CSS) ParseRuleList1(css CSSStringSource) (ret js.Promise[js.Array[CSSParserRule]]) {
	bindings.CallCSSParseRuleList1(
		js.Pointer(&ret),
		css.Ref(),
	)
	return
}

// TryParseRuleList1 calls the function "CSS.parseRuleList"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryParseRuleList1(css CSSStringSource) (ret js.Promise[js.Array[CSSParserRule]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseRuleList1(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
	)
	return
}

// HasFuncParseRule returns ture if the function "CSS.parseRule" exists.
func (CSS) HasFuncParseRule() bool {
	return js.True == bindings.HasFuncCSSParseRule()
}

// FuncParseRule returns the function "CSS.parseRule".
func (CSS) FuncParseRule() (fn js.Func[func(css CSSStringSource, options CSSParserOptions) js.Promise[CSSParserRule]]) {
	bindings.FuncCSSParseRule(
		js.Pointer(&fn),
	)
	return
}

// ParseRule calls the function "CSS.parseRule".
func (CSS) ParseRule(css CSSStringSource, options CSSParserOptions) (ret js.Promise[CSSParserRule]) {
	bindings.CallCSSParseRule(
		js.Pointer(&ret),
		css.Ref(),
		js.Pointer(&options),
	)
	return
}

// TryParseRule calls the function "CSS.parseRule"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryParseRule(css CSSStringSource, options CSSParserOptions) (ret js.Promise[CSSParserRule], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseRule(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
		js.Pointer(&options),
	)
	return
}

// HasFuncParseRule1 returns ture if the function "CSS.parseRule" exists.
func (CSS) HasFuncParseRule1() bool {
	return js.True == bindings.HasFuncCSSParseRule1()
}

// FuncParseRule1 returns the function "CSS.parseRule".
func (CSS) FuncParseRule1() (fn js.Func[func(css CSSStringSource) js.Promise[CSSParserRule]]) {
	bindings.FuncCSSParseRule1(
		js.Pointer(&fn),
	)
	return
}

// ParseRule1 calls the function "CSS.parseRule".
func (CSS) ParseRule1(css CSSStringSource) (ret js.Promise[CSSParserRule]) {
	bindings.CallCSSParseRule1(
		js.Pointer(&ret),
		css.Ref(),
	)
	return
}

// TryParseRule1 calls the function "CSS.parseRule"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryParseRule1(css CSSStringSource) (ret js.Promise[CSSParserRule], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseRule1(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
	)
	return
}

// HasFuncParseDeclarationList returns ture if the function "CSS.parseDeclarationList" exists.
func (CSS) HasFuncParseDeclarationList() bool {
	return js.True == bindings.HasFuncCSSParseDeclarationList()
}

// FuncParseDeclarationList returns the function "CSS.parseDeclarationList".
func (CSS) FuncParseDeclarationList() (fn js.Func[func(css CSSStringSource, options CSSParserOptions) js.Promise[js.Array[CSSParserRule]]]) {
	bindings.FuncCSSParseDeclarationList(
		js.Pointer(&fn),
	)
	return
}

// ParseDeclarationList calls the function "CSS.parseDeclarationList".
func (CSS) ParseDeclarationList(css CSSStringSource, options CSSParserOptions) (ret js.Promise[js.Array[CSSParserRule]]) {
	bindings.CallCSSParseDeclarationList(
		js.Pointer(&ret),
		css.Ref(),
		js.Pointer(&options),
	)
	return
}

// TryParseDeclarationList calls the function "CSS.parseDeclarationList"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryParseDeclarationList(css CSSStringSource, options CSSParserOptions) (ret js.Promise[js.Array[CSSParserRule]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseDeclarationList(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
		js.Pointer(&options),
	)
	return
}

// HasFuncParseDeclarationList1 returns ture if the function "CSS.parseDeclarationList" exists.
func (CSS) HasFuncParseDeclarationList1() bool {
	return js.True == bindings.HasFuncCSSParseDeclarationList1()
}

// FuncParseDeclarationList1 returns the function "CSS.parseDeclarationList".
func (CSS) FuncParseDeclarationList1() (fn js.Func[func(css CSSStringSource) js.Promise[js.Array[CSSParserRule]]]) {
	bindings.FuncCSSParseDeclarationList1(
		js.Pointer(&fn),
	)
	return
}

// ParseDeclarationList1 calls the function "CSS.parseDeclarationList".
func (CSS) ParseDeclarationList1(css CSSStringSource) (ret js.Promise[js.Array[CSSParserRule]]) {
	bindings.CallCSSParseDeclarationList1(
		js.Pointer(&ret),
		css.Ref(),
	)
	return
}

// TryParseDeclarationList1 calls the function "CSS.parseDeclarationList"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryParseDeclarationList1(css CSSStringSource) (ret js.Promise[js.Array[CSSParserRule]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseDeclarationList1(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
	)
	return
}

// HasFuncParseDeclaration returns ture if the function "CSS.parseDeclaration" exists.
func (CSS) HasFuncParseDeclaration() bool {
	return js.True == bindings.HasFuncCSSParseDeclaration()
}

// FuncParseDeclaration returns the function "CSS.parseDeclaration".
func (CSS) FuncParseDeclaration() (fn js.Func[func(css js.String, options CSSParserOptions) CSSParserDeclaration]) {
	bindings.FuncCSSParseDeclaration(
		js.Pointer(&fn),
	)
	return
}

// ParseDeclaration calls the function "CSS.parseDeclaration".
func (CSS) ParseDeclaration(css js.String, options CSSParserOptions) (ret CSSParserDeclaration) {
	bindings.CallCSSParseDeclaration(
		js.Pointer(&ret),
		css.Ref(),
		js.Pointer(&options),
	)
	return
}

// TryParseDeclaration calls the function "CSS.parseDeclaration"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryParseDeclaration(css js.String, options CSSParserOptions) (ret CSSParserDeclaration, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseDeclaration(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
		js.Pointer(&options),
	)
	return
}

// HasFuncParseDeclaration1 returns ture if the function "CSS.parseDeclaration" exists.
func (CSS) HasFuncParseDeclaration1() bool {
	return js.True == bindings.HasFuncCSSParseDeclaration1()
}

// FuncParseDeclaration1 returns the function "CSS.parseDeclaration".
func (CSS) FuncParseDeclaration1() (fn js.Func[func(css js.String) CSSParserDeclaration]) {
	bindings.FuncCSSParseDeclaration1(
		js.Pointer(&fn),
	)
	return
}

// ParseDeclaration1 calls the function "CSS.parseDeclaration".
func (CSS) ParseDeclaration1(css js.String) (ret CSSParserDeclaration) {
	bindings.CallCSSParseDeclaration1(
		js.Pointer(&ret),
		css.Ref(),
	)
	return
}

// TryParseDeclaration1 calls the function "CSS.parseDeclaration"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryParseDeclaration1(css js.String) (ret CSSParserDeclaration, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseDeclaration1(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
	)
	return
}

// HasFuncParseValue returns ture if the function "CSS.parseValue" exists.
func (CSS) HasFuncParseValue() bool {
	return js.True == bindings.HasFuncCSSParseValue()
}

// FuncParseValue returns the function "CSS.parseValue".
func (CSS) FuncParseValue() (fn js.Func[func(css js.String) CSSToken]) {
	bindings.FuncCSSParseValue(
		js.Pointer(&fn),
	)
	return
}

// ParseValue calls the function "CSS.parseValue".
func (CSS) ParseValue(css js.String) (ret CSSToken) {
	bindings.CallCSSParseValue(
		js.Pointer(&ret),
		css.Ref(),
	)
	return
}

// TryParseValue calls the function "CSS.parseValue"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryParseValue(css js.String) (ret CSSToken, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseValue(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
	)
	return
}

// HasFuncParseValueList returns ture if the function "CSS.parseValueList" exists.
func (CSS) HasFuncParseValueList() bool {
	return js.True == bindings.HasFuncCSSParseValueList()
}

// FuncParseValueList returns the function "CSS.parseValueList".
func (CSS) FuncParseValueList() (fn js.Func[func(css js.String) js.Array[CSSToken]]) {
	bindings.FuncCSSParseValueList(
		js.Pointer(&fn),
	)
	return
}

// ParseValueList calls the function "CSS.parseValueList".
func (CSS) ParseValueList(css js.String) (ret js.Array[CSSToken]) {
	bindings.CallCSSParseValueList(
		js.Pointer(&ret),
		css.Ref(),
	)
	return
}

// TryParseValueList calls the function "CSS.parseValueList"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryParseValueList(css js.String) (ret js.Array[CSSToken], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseValueList(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
	)
	return
}

// HasFuncParseCommaValueList returns ture if the function "CSS.parseCommaValueList" exists.
func (CSS) HasFuncParseCommaValueList() bool {
	return js.True == bindings.HasFuncCSSParseCommaValueList()
}

// FuncParseCommaValueList returns the function "CSS.parseCommaValueList".
func (CSS) FuncParseCommaValueList() (fn js.Func[func(css js.String) js.Array[js.Array[CSSToken]]]) {
	bindings.FuncCSSParseCommaValueList(
		js.Pointer(&fn),
	)
	return
}

// ParseCommaValueList calls the function "CSS.parseCommaValueList".
func (CSS) ParseCommaValueList(css js.String) (ret js.Array[js.Array[CSSToken]]) {
	bindings.CallCSSParseCommaValueList(
		js.Pointer(&ret),
		css.Ref(),
	)
	return
}

// TryParseCommaValueList calls the function "CSS.parseCommaValueList"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (CSS) TryParseCommaValueList(css js.String) (ret js.Array[js.Array[CSSToken]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseCommaValueList(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
	)
	return
}

func NewCSSAnimation(effect AnimationEffect, timeline AnimationTimeline) (ret CSSAnimation) {
	ret.ref = bindings.NewCSSAnimationByCSSAnimation(
		effect.Ref(),
		timeline.Ref())
	return
}

func NewCSSAnimationByCSSAnimation1(effect AnimationEffect) (ret CSSAnimation) {
	ret.ref = bindings.NewCSSAnimationByCSSAnimation1(
		effect.Ref())
	return
}

func NewCSSAnimationByCSSAnimation2() (ret CSSAnimation) {
	ret.ref = bindings.NewCSSAnimationByCSSAnimation2()
	return
}

type CSSAnimation struct {
	Animation
}

func (this CSSAnimation) Once() CSSAnimation {
	this.ref.Once()
	return this
}

func (this CSSAnimation) Ref() js.Ref {
	return this.Animation.Ref()
}

func (this CSSAnimation) FromRef(ref js.Ref) CSSAnimation {
	this.Animation = this.Animation.FromRef(ref)
	return this
}

func (this CSSAnimation) Free() {
	this.ref.Free()
}

// AnimationName returns the value of property "CSSAnimation.animationName".
//
// It returns ok=false if there is no such property.
func (this CSSAnimation) AnimationName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSAnimationAnimationName(
		this.ref, js.Pointer(&ret),
	)
	return
}

type OneOf_String_CSSKeywordValue struct {
	ref js.Ref
}

func (x OneOf_String_CSSKeywordValue) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_CSSKeywordValue) Free() {
	x.ref.Free()
}

func (x OneOf_String_CSSKeywordValue) FromRef(ref js.Ref) OneOf_String_CSSKeywordValue {
	return OneOf_String_CSSKeywordValue{
		ref: ref,
	}
}

func (x OneOf_String_CSSKeywordValue) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_CSSKeywordValue) CSSKeywordValue() CSSKeywordValue {
	return CSSKeywordValue{}.FromRef(x.ref)
}

type CSSKeywordish = OneOf_String_CSSKeywordValue

type OneOf_Float64_CSSNumericValue_String_CSSKeywordValue struct {
	ref js.Ref
}

func (x OneOf_Float64_CSSNumericValue_String_CSSKeywordValue) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Float64_CSSNumericValue_String_CSSKeywordValue) Free() {
	x.ref.Free()
}

func (x OneOf_Float64_CSSNumericValue_String_CSSKeywordValue) FromRef(ref js.Ref) OneOf_Float64_CSSNumericValue_String_CSSKeywordValue {
	return OneOf_Float64_CSSNumericValue_String_CSSKeywordValue{
		ref: ref,
	}
}

func (x OneOf_Float64_CSSNumericValue_String_CSSKeywordValue) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Float64_CSSNumericValue_String_CSSKeywordValue) CSSNumericValue() CSSNumericValue {
	return CSSNumericValue{}.FromRef(x.ref)
}

func (x OneOf_Float64_CSSNumericValue_String_CSSKeywordValue) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_Float64_CSSNumericValue_String_CSSKeywordValue) CSSKeywordValue() CSSKeywordValue {
	return CSSKeywordValue{}.FromRef(x.ref)
}

type CSSColorPercent = OneOf_Float64_CSSNumericValue_String_CSSKeywordValue
