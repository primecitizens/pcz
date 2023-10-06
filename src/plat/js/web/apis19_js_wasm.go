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
func (p BluetoothServiceDataFilterInit) UpdateFrom(ref js.Ref) {
	bindings.BluetoothServiceDataFilterInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BluetoothServiceDataFilterInit) Update(ref js.Ref) {
	bindings.BluetoothServiceDataFilterInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p BluetoothLEScanFilterInit) UpdateFrom(ref js.Ref) {
	bindings.BluetoothLEScanFilterInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BluetoothLEScanFilterInit) Update(ref js.Ref) {
	bindings.BluetoothLEScanFilterInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p RequestDeviceOptions) UpdateFrom(ref js.Ref) {
	bindings.RequestDeviceOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RequestDeviceOptions) Update(ref js.Ref) {
	bindings.RequestDeviceOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type Bluetooth struct {
	EventTarget
}

func (this Bluetooth) Once() Bluetooth {
	this.Ref().Once()
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
	this.Ref().Free()
}

// ReferringDevice returns the value of property "Bluetooth.referringDevice".
//
// It returns ok=false if there is no such property.
func (this Bluetooth) ReferringDevice() (ret BluetoothDevice, ok bool) {
	ok = js.True == bindings.GetBluetoothReferringDevice(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetAvailability returns true if the method "Bluetooth.getAvailability" exists.
func (this Bluetooth) HasGetAvailability() bool {
	return js.True == bindings.HasBluetoothGetAvailability(
		this.Ref(),
	)
}

// GetAvailabilityFunc returns the method "Bluetooth.getAvailability".
func (this Bluetooth) GetAvailabilityFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.BluetoothGetAvailabilityFunc(
			this.Ref(),
		),
	)
}

// GetAvailability calls the method "Bluetooth.getAvailability".
func (this Bluetooth) GetAvailability() (ret js.Promise[js.Boolean]) {
	bindings.CallBluetoothGetAvailability(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAvailability calls the method "Bluetooth.getAvailability"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Bluetooth) TryGetAvailability() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothGetAvailability(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetDevices returns true if the method "Bluetooth.getDevices" exists.
func (this Bluetooth) HasGetDevices() bool {
	return js.True == bindings.HasBluetoothGetDevices(
		this.Ref(),
	)
}

// GetDevicesFunc returns the method "Bluetooth.getDevices".
func (this Bluetooth) GetDevicesFunc() (fn js.Func[func() js.Promise[js.Array[BluetoothDevice]]]) {
	return fn.FromRef(
		bindings.BluetoothGetDevicesFunc(
			this.Ref(),
		),
	)
}

// GetDevices calls the method "Bluetooth.getDevices".
func (this Bluetooth) GetDevices() (ret js.Promise[js.Array[BluetoothDevice]]) {
	bindings.CallBluetoothGetDevices(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetDevices calls the method "Bluetooth.getDevices"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Bluetooth) TryGetDevices() (ret js.Promise[js.Array[BluetoothDevice]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothGetDevices(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRequestDevice returns true if the method "Bluetooth.requestDevice" exists.
func (this Bluetooth) HasRequestDevice() bool {
	return js.True == bindings.HasBluetoothRequestDevice(
		this.Ref(),
	)
}

// RequestDeviceFunc returns the method "Bluetooth.requestDevice".
func (this Bluetooth) RequestDeviceFunc() (fn js.Func[func(options RequestDeviceOptions) js.Promise[BluetoothDevice]]) {
	return fn.FromRef(
		bindings.BluetoothRequestDeviceFunc(
			this.Ref(),
		),
	)
}

// RequestDevice calls the method "Bluetooth.requestDevice".
func (this Bluetooth) RequestDevice(options RequestDeviceOptions) (ret js.Promise[BluetoothDevice]) {
	bindings.CallBluetoothRequestDevice(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestDevice calls the method "Bluetooth.requestDevice"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Bluetooth) TryRequestDevice(options RequestDeviceOptions) (ret js.Promise[BluetoothDevice], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRequestDevice(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasRequestDevice1 returns true if the method "Bluetooth.requestDevice" exists.
func (this Bluetooth) HasRequestDevice1() bool {
	return js.True == bindings.HasBluetoothRequestDevice1(
		this.Ref(),
	)
}

// RequestDevice1Func returns the method "Bluetooth.requestDevice".
func (this Bluetooth) RequestDevice1Func() (fn js.Func[func() js.Promise[BluetoothDevice]]) {
	return fn.FromRef(
		bindings.BluetoothRequestDevice1Func(
			this.Ref(),
		),
	)
}

// RequestDevice1 calls the method "Bluetooth.requestDevice".
func (this Bluetooth) RequestDevice1() (ret js.Promise[BluetoothDevice]) {
	bindings.CallBluetoothRequestDevice1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestDevice1 calls the method "Bluetooth.requestDevice"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Bluetooth) TryRequestDevice1() (ret js.Promise[BluetoothDevice], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRequestDevice1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type BluetoothManufacturerDataMap struct {
	ref js.Ref
}

func (this BluetoothManufacturerDataMap) Once() BluetoothManufacturerDataMap {
	this.Ref().Once()
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
	this.Ref().Free()
}

type BluetoothServiceDataMap struct {
	ref js.Ref
}

func (this BluetoothServiceDataMap) Once() BluetoothServiceDataMap {
	this.Ref().Once()
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
	this.Ref().Free()
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
func (p BluetoothAdvertisingEventInit) UpdateFrom(ref js.Ref) {
	bindings.BluetoothAdvertisingEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BluetoothAdvertisingEventInit) Update(ref js.Ref) {
	bindings.BluetoothAdvertisingEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Device returns the value of property "BluetoothAdvertisingEvent.device".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) Device() (ret BluetoothDevice, ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventDevice(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Uuids returns the value of property "BluetoothAdvertisingEvent.uuids".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) Uuids() (ret js.FrozenArray[UUID], ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventUuids(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "BluetoothAdvertisingEvent.name".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Appearance returns the value of property "BluetoothAdvertisingEvent.appearance".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) Appearance() (ret uint16, ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventAppearance(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TxPower returns the value of property "BluetoothAdvertisingEvent.txPower".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) TxPower() (ret int8, ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventTxPower(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Rssi returns the value of property "BluetoothAdvertisingEvent.rssi".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) Rssi() (ret int8, ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventRssi(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ManufacturerData returns the value of property "BluetoothAdvertisingEvent.manufacturerData".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) ManufacturerData() (ret BluetoothManufacturerDataMap, ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventManufacturerData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ServiceData returns the value of property "BluetoothAdvertisingEvent.serviceData".
//
// It returns ok=false if there is no such property.
func (this BluetoothAdvertisingEvent) ServiceData() (ret BluetoothServiceDataMap, ok bool) {
	ok = js.True == bindings.GetBluetoothAdvertisingEventServiceData(
		this.Ref(), js.Pointer(&ret),
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
func (p BluetoothDataFilterInit) UpdateFrom(ref js.Ref) {
	bindings.BluetoothDataFilterInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BluetoothDataFilterInit) Update(ref js.Ref) {
	bindings.BluetoothDataFilterInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p BluetoothPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.BluetoothPermissionDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BluetoothPermissionDescriptor) Update(ref js.Ref) {
	bindings.BluetoothPermissionDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type BluetoothPermissionResult struct {
	PermissionStatus
}

func (this BluetoothPermissionResult) Once() BluetoothPermissionResult {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Devices returns the value of property "BluetoothPermissionResult.devices".
//
// It returns ok=false if there is no such property.
func (this BluetoothPermissionResult) Devices() (ret js.FrozenArray[BluetoothDevice], ok bool) {
	ok = js.True == bindings.GetBluetoothPermissionResultDevices(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDevices sets the value of property "BluetoothPermissionResult.devices" to val.
//
// It returns false if the property cannot be set.
func (this BluetoothPermissionResult) SetDevices(val js.FrozenArray[BluetoothDevice]) bool {
	return js.True == bindings.SetBluetoothPermissionResultDevices(
		this.Ref(),
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
func (p BluetoothPermissionStorage) UpdateFrom(ref js.Ref) {
	bindings.BluetoothPermissionStorageJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BluetoothPermissionStorage) Update(ref js.Ref) {
	bindings.BluetoothPermissionStorageJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type BluetoothUUID struct {
	ref js.Ref
}

func (this BluetoothUUID) Once() BluetoothUUID {
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasGetService returns true if the staticmethod "BluetoothUUID.getService" exists.
func (this BluetoothUUID) HasGetService() bool {
	return js.True == bindings.HasBluetoothUUIDGetService(
		this.Ref(),
	)
}

// GetServiceFunc returns the staticmethod "BluetoothUUID.getService".
func (this BluetoothUUID) GetServiceFunc() (fn js.Func[func(name OneOf_String_Uint32) UUID]) {
	return fn.FromRef(
		bindings.BluetoothUUIDGetServiceFunc(
			this.Ref(),
		),
	)
}

// GetService calls the staticmethod "BluetoothUUID.getService".
func (this BluetoothUUID) GetService(name OneOf_String_Uint32) (ret UUID) {
	bindings.CallBluetoothUUIDGetService(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetService calls the staticmethod "BluetoothUUID.getService"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothUUID) TryGetService(name OneOf_String_Uint32) (ret UUID, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothUUIDGetService(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasGetCharacteristic returns true if the staticmethod "BluetoothUUID.getCharacteristic" exists.
func (this BluetoothUUID) HasGetCharacteristic() bool {
	return js.True == bindings.HasBluetoothUUIDGetCharacteristic(
		this.Ref(),
	)
}

// GetCharacteristicFunc returns the staticmethod "BluetoothUUID.getCharacteristic".
func (this BluetoothUUID) GetCharacteristicFunc() (fn js.Func[func(name OneOf_String_Uint32) UUID]) {
	return fn.FromRef(
		bindings.BluetoothUUIDGetCharacteristicFunc(
			this.Ref(),
		),
	)
}

// GetCharacteristic calls the staticmethod "BluetoothUUID.getCharacteristic".
func (this BluetoothUUID) GetCharacteristic(name OneOf_String_Uint32) (ret UUID) {
	bindings.CallBluetoothUUIDGetCharacteristic(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetCharacteristic calls the staticmethod "BluetoothUUID.getCharacteristic"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothUUID) TryGetCharacteristic(name OneOf_String_Uint32) (ret UUID, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothUUIDGetCharacteristic(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasGetDescriptor returns true if the staticmethod "BluetoothUUID.getDescriptor" exists.
func (this BluetoothUUID) HasGetDescriptor() bool {
	return js.True == bindings.HasBluetoothUUIDGetDescriptor(
		this.Ref(),
	)
}

// GetDescriptorFunc returns the staticmethod "BluetoothUUID.getDescriptor".
func (this BluetoothUUID) GetDescriptorFunc() (fn js.Func[func(name OneOf_String_Uint32) UUID]) {
	return fn.FromRef(
		bindings.BluetoothUUIDGetDescriptorFunc(
			this.Ref(),
		),
	)
}

// GetDescriptor calls the staticmethod "BluetoothUUID.getDescriptor".
func (this BluetoothUUID) GetDescriptor(name OneOf_String_Uint32) (ret UUID) {
	bindings.CallBluetoothUUIDGetDescriptor(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetDescriptor calls the staticmethod "BluetoothUUID.getDescriptor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothUUID) TryGetDescriptor(name OneOf_String_Uint32) (ret UUID, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothUUIDGetDescriptor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasCanonicalUUID returns true if the staticmethod "BluetoothUUID.canonicalUUID" exists.
func (this BluetoothUUID) HasCanonicalUUID() bool {
	return js.True == bindings.HasBluetoothUUIDCanonicalUUID(
		this.Ref(),
	)
}

// CanonicalUUIDFunc returns the staticmethod "BluetoothUUID.canonicalUUID".
func (this BluetoothUUID) CanonicalUUIDFunc() (fn js.Func[func(alias uint32) UUID]) {
	return fn.FromRef(
		bindings.BluetoothUUIDCanonicalUUIDFunc(
			this.Ref(),
		),
	)
}

// CanonicalUUID calls the staticmethod "BluetoothUUID.canonicalUUID".
func (this BluetoothUUID) CanonicalUUID(alias uint32) (ret UUID) {
	bindings.CallBluetoothUUIDCanonicalUUID(
		this.Ref(), js.Pointer(&ret),
		uint32(alias),
	)

	return
}

// TryCanonicalUUID calls the staticmethod "BluetoothUUID.canonicalUUID"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothUUID) TryCanonicalUUID(alias uint32) (ret UUID, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothUUIDCanonicalUUID(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// MinContentSize returns the value of property "IntrinsicSizes.minContentSize".
//
// It returns ok=false if there is no such property.
func (this IntrinsicSizes) MinContentSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetIntrinsicSizesMinContentSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxContentSize returns the value of property "IntrinsicSizes.maxContentSize".
//
// It returns ok=false if there is no such property.
func (this IntrinsicSizes) MaxContentSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetIntrinsicSizesMaxContentSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type LayoutFragment struct {
	ref js.Ref
}

func (this LayoutFragment) Once() LayoutFragment {
	this.Ref().Once()
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
	this.Ref().Free()
}

// InlineSize returns the value of property "LayoutFragment.inlineSize".
//
// It returns ok=false if there is no such property.
func (this LayoutFragment) InlineSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutFragmentInlineSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BlockSize returns the value of property "LayoutFragment.blockSize".
//
// It returns ok=false if there is no such property.
func (this LayoutFragment) BlockSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutFragmentBlockSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// InlineOffset returns the value of property "LayoutFragment.inlineOffset".
//
// It returns ok=false if there is no such property.
func (this LayoutFragment) InlineOffset() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutFragmentInlineOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetInlineOffset sets the value of property "LayoutFragment.inlineOffset" to val.
//
// It returns false if the property cannot be set.
func (this LayoutFragment) SetInlineOffset(val float64) bool {
	return js.True == bindings.SetLayoutFragmentInlineOffset(
		this.Ref(),
		float64(val),
	)
}

// BlockOffset returns the value of property "LayoutFragment.blockOffset".
//
// It returns ok=false if there is no such property.
func (this LayoutFragment) BlockOffset() (ret float64, ok bool) {
	ok = js.True == bindings.GetLayoutFragmentBlockOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBlockOffset sets the value of property "LayoutFragment.blockOffset" to val.
//
// It returns false if the property cannot be set.
func (this LayoutFragment) SetBlockOffset(val float64) bool {
	return js.True == bindings.SetLayoutFragmentBlockOffset(
		this.Ref(),
		float64(val),
	)
}

// Data returns the value of property "LayoutFragment.data".
//
// It returns ok=false if there is no such property.
func (this LayoutFragment) Data() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetLayoutFragmentData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BreakToken returns the value of property "LayoutFragment.breakToken".
//
// It returns ok=false if there is no such property.
func (this LayoutFragment) BreakToken() (ret ChildBreakToken, ok bool) {
	ok = js.True == bindings.GetLayoutFragmentBreakToken(
		this.Ref(), js.Pointer(&ret),
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
func (p LayoutConstraintsOptions) UpdateFrom(ref js.Ref) {
	bindings.LayoutConstraintsOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p LayoutConstraintsOptions) Update(ref js.Ref) {
	bindings.LayoutConstraintsOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type LayoutChild struct {
	ref js.Ref
}

func (this LayoutChild) Once() LayoutChild {
	this.Ref().Once()
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
	this.Ref().Free()
}

// StyleMap returns the value of property "LayoutChild.styleMap".
//
// It returns ok=false if there is no such property.
func (this LayoutChild) StyleMap() (ret StylePropertyMapReadOnly, ok bool) {
	ok = js.True == bindings.GetLayoutChildStyleMap(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasIntrinsicSizes returns true if the method "LayoutChild.intrinsicSizes" exists.
func (this LayoutChild) HasIntrinsicSizes() bool {
	return js.True == bindings.HasLayoutChildIntrinsicSizes(
		this.Ref(),
	)
}

// IntrinsicSizesFunc returns the method "LayoutChild.intrinsicSizes".
func (this LayoutChild) IntrinsicSizesFunc() (fn js.Func[func() js.Promise[IntrinsicSizes]]) {
	return fn.FromRef(
		bindings.LayoutChildIntrinsicSizesFunc(
			this.Ref(),
		),
	)
}

// IntrinsicSizes calls the method "LayoutChild.intrinsicSizes".
func (this LayoutChild) IntrinsicSizes() (ret js.Promise[IntrinsicSizes]) {
	bindings.CallLayoutChildIntrinsicSizes(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryIntrinsicSizes calls the method "LayoutChild.intrinsicSizes"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this LayoutChild) TryIntrinsicSizes() (ret js.Promise[IntrinsicSizes], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLayoutChildIntrinsicSizes(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasLayoutNextFragment returns true if the method "LayoutChild.layoutNextFragment" exists.
func (this LayoutChild) HasLayoutNextFragment() bool {
	return js.True == bindings.HasLayoutChildLayoutNextFragment(
		this.Ref(),
	)
}

// LayoutNextFragmentFunc returns the method "LayoutChild.layoutNextFragment".
func (this LayoutChild) LayoutNextFragmentFunc() (fn js.Func[func(constraints LayoutConstraintsOptions, breakToken ChildBreakToken) js.Promise[LayoutFragment]]) {
	return fn.FromRef(
		bindings.LayoutChildLayoutNextFragmentFunc(
			this.Ref(),
		),
	)
}

// LayoutNextFragment calls the method "LayoutChild.layoutNextFragment".
func (this LayoutChild) LayoutNextFragment(constraints LayoutConstraintsOptions, breakToken ChildBreakToken) (ret js.Promise[LayoutFragment]) {
	bindings.CallLayoutChildLayoutNextFragment(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&constraints),
		breakToken.Ref(),
	)

	return
}

// TryLayoutNextFragment calls the method "LayoutChild.layoutNextFragment"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this LayoutChild) TryLayoutNextFragment(constraints LayoutConstraintsOptions, breakToken ChildBreakToken) (ret js.Promise[LayoutFragment], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLayoutChildLayoutNextFragment(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&constraints),
		breakToken.Ref(),
	)

	return
}

type ChildBreakToken struct {
	ref js.Ref
}

func (this ChildBreakToken) Once() ChildBreakToken {
	this.Ref().Once()
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
	this.Ref().Free()
}

// BreakType returns the value of property "ChildBreakToken.breakType".
//
// It returns ok=false if there is no such property.
func (this ChildBreakToken) BreakType() (ret BreakType, ok bool) {
	ok = js.True == bindings.GetChildBreakTokenBreakType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Child returns the value of property "ChildBreakToken.child".
//
// It returns ok=false if there is no such property.
func (this ChildBreakToken) Child() (ret LayoutChild, ok bool) {
	ok = js.True == bindings.GetChildBreakTokenChild(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type BreakToken struct {
	ref js.Ref
}

func (this BreakToken) Once() BreakToken {
	this.Ref().Once()
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
	this.Ref().Free()
}

// ChildBreakTokens returns the value of property "BreakToken.childBreakTokens".
//
// It returns ok=false if there is no such property.
func (this BreakToken) ChildBreakTokens() (ret js.FrozenArray[ChildBreakToken], ok bool) {
	ok = js.True == bindings.GetBreakTokenChildBreakTokens(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Data returns the value of property "BreakToken.data".
//
// It returns ok=false if there is no such property.
func (this BreakToken) Data() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetBreakTokenData(
		this.Ref(), js.Pointer(&ret),
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
func (p BreakTokenOptions) UpdateFrom(ref js.Ref) {
	bindings.BreakTokenOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BreakTokenOptions) Update(ref js.Ref) {
	bindings.BreakTokenOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Name returns the value of property "BroadcastChannel.name".
//
// It returns ok=false if there is no such property.
func (this BroadcastChannel) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBroadcastChannelName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasPostMessage returns true if the method "BroadcastChannel.postMessage" exists.
func (this BroadcastChannel) HasPostMessage() bool {
	return js.True == bindings.HasBroadcastChannelPostMessage(
		this.Ref(),
	)
}

// PostMessageFunc returns the method "BroadcastChannel.postMessage".
func (this BroadcastChannel) PostMessageFunc() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.BroadcastChannelPostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage calls the method "BroadcastChannel.postMessage".
func (this BroadcastChannel) PostMessage(message js.Any) (ret js.Void) {
	bindings.CallBroadcastChannelPostMessage(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPostMessage calls the method "BroadcastChannel.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BroadcastChannel) TryPostMessage(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBroadcastChannelPostMessage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasClose returns true if the method "BroadcastChannel.close" exists.
func (this BroadcastChannel) HasClose() bool {
	return js.True == bindings.HasBroadcastChannelClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "BroadcastChannel.close".
func (this BroadcastChannel) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.BroadcastChannelCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "BroadcastChannel.close".
func (this BroadcastChannel) Close() (ret js.Void) {
	bindings.CallBroadcastChannelClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "BroadcastChannel.close"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BroadcastChannel) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBroadcastChannelClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type CropTarget struct {
	ref js.Ref
}

func (this CropTarget) Once() CropTarget {
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasFromElement returns true if the staticmethod "CropTarget.fromElement" exists.
func (this CropTarget) HasFromElement() bool {
	return js.True == bindings.HasCropTargetFromElement(
		this.Ref(),
	)
}

// FromElementFunc returns the staticmethod "CropTarget.fromElement".
func (this CropTarget) FromElementFunc() (fn js.Func[func(element Element) js.Promise[CropTarget]]) {
	return fn.FromRef(
		bindings.CropTargetFromElementFunc(
			this.Ref(),
		),
	)
}

// FromElement calls the staticmethod "CropTarget.fromElement".
func (this CropTarget) FromElement(element Element) (ret js.Promise[CropTarget]) {
	bindings.CallCropTargetFromElement(
		this.Ref(), js.Pointer(&ret),
		element.Ref(),
	)

	return
}

// TryFromElement calls the staticmethod "CropTarget.fromElement"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CropTarget) TryFromElement(element Element) (ret js.Promise[CropTarget], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCropTargetFromElement(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
	)

	return
}

type RestrictionTarget struct {
	ref js.Ref
}

func (this RestrictionTarget) Once() RestrictionTarget {
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasFromElement returns true if the staticmethod "RestrictionTarget.fromElement" exists.
func (this RestrictionTarget) HasFromElement() bool {
	return js.True == bindings.HasRestrictionTargetFromElement(
		this.Ref(),
	)
}

// FromElementFunc returns the staticmethod "RestrictionTarget.fromElement".
func (this RestrictionTarget) FromElementFunc() (fn js.Func[func(element Element) js.Promise[RestrictionTarget]]) {
	return fn.FromRef(
		bindings.RestrictionTargetFromElementFunc(
			this.Ref(),
		),
	)
}

// FromElement calls the staticmethod "RestrictionTarget.fromElement".
func (this RestrictionTarget) FromElement(element Element) (ret js.Promise[RestrictionTarget]) {
	bindings.CallRestrictionTargetFromElement(
		this.Ref(), js.Pointer(&ret),
		element.Ref(),
	)

	return
}

// TryFromElement calls the staticmethod "RestrictionTarget.fromElement"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RestrictionTarget) TryFromElement(element Element) (ret js.Promise[RestrictionTarget], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRestrictionTargetFromElement(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
	)

	return
}

type BrowserCaptureMediaStreamTrack struct {
	MediaStreamTrack
}

func (this BrowserCaptureMediaStreamTrack) Once() BrowserCaptureMediaStreamTrack {
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasCropTo returns true if the method "BrowserCaptureMediaStreamTrack.cropTo" exists.
func (this BrowserCaptureMediaStreamTrack) HasCropTo() bool {
	return js.True == bindings.HasBrowserCaptureMediaStreamTrackCropTo(
		this.Ref(),
	)
}

// CropToFunc returns the method "BrowserCaptureMediaStreamTrack.cropTo".
func (this BrowserCaptureMediaStreamTrack) CropToFunc() (fn js.Func[func(cropTarget CropTarget) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BrowserCaptureMediaStreamTrackCropToFunc(
			this.Ref(),
		),
	)
}

// CropTo calls the method "BrowserCaptureMediaStreamTrack.cropTo".
func (this BrowserCaptureMediaStreamTrack) CropTo(cropTarget CropTarget) (ret js.Promise[js.Void]) {
	bindings.CallBrowserCaptureMediaStreamTrackCropTo(
		this.Ref(), js.Pointer(&ret),
		cropTarget.Ref(),
	)

	return
}

// TryCropTo calls the method "BrowserCaptureMediaStreamTrack.cropTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BrowserCaptureMediaStreamTrack) TryCropTo(cropTarget CropTarget) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBrowserCaptureMediaStreamTrackCropTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		cropTarget.Ref(),
	)

	return
}

// HasClone returns true if the method "BrowserCaptureMediaStreamTrack.clone" exists.
func (this BrowserCaptureMediaStreamTrack) HasClone() bool {
	return js.True == bindings.HasBrowserCaptureMediaStreamTrackClone(
		this.Ref(),
	)
}

// CloneFunc returns the method "BrowserCaptureMediaStreamTrack.clone".
func (this BrowserCaptureMediaStreamTrack) CloneFunc() (fn js.Func[func() BrowserCaptureMediaStreamTrack]) {
	return fn.FromRef(
		bindings.BrowserCaptureMediaStreamTrackCloneFunc(
			this.Ref(),
		),
	)
}

// Clone calls the method "BrowserCaptureMediaStreamTrack.clone".
func (this BrowserCaptureMediaStreamTrack) Clone() (ret BrowserCaptureMediaStreamTrack) {
	bindings.CallBrowserCaptureMediaStreamTrackClone(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClone calls the method "BrowserCaptureMediaStreamTrack.clone"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BrowserCaptureMediaStreamTrack) TryClone() (ret BrowserCaptureMediaStreamTrack, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBrowserCaptureMediaStreamTrackClone(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRestrictTo returns true if the method "BrowserCaptureMediaStreamTrack.restrictTo" exists.
func (this BrowserCaptureMediaStreamTrack) HasRestrictTo() bool {
	return js.True == bindings.HasBrowserCaptureMediaStreamTrackRestrictTo(
		this.Ref(),
	)
}

// RestrictToFunc returns the method "BrowserCaptureMediaStreamTrack.restrictTo".
func (this BrowserCaptureMediaStreamTrack) RestrictToFunc() (fn js.Func[func(RestrictionTarget RestrictionTarget) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BrowserCaptureMediaStreamTrackRestrictToFunc(
			this.Ref(),
		),
	)
}

// RestrictTo calls the method "BrowserCaptureMediaStreamTrack.restrictTo".
func (this BrowserCaptureMediaStreamTrack) RestrictTo(RestrictionTarget RestrictionTarget) (ret js.Promise[js.Void]) {
	bindings.CallBrowserCaptureMediaStreamTrackRestrictTo(
		this.Ref(), js.Pointer(&ret),
		RestrictionTarget.Ref(),
	)

	return
}

// TryRestrictTo calls the method "BrowserCaptureMediaStreamTrack.restrictTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BrowserCaptureMediaStreamTrack) TryRestrictTo(RestrictionTarget RestrictionTarget) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBrowserCaptureMediaStreamTrackRestrictTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
func (p QueuingStrategyInit) UpdateFrom(ref js.Ref) {
	bindings.QueuingStrategyInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p QueuingStrategyInit) Update(ref js.Ref) {
	bindings.QueuingStrategyInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
		assert.Throw("invalid", "callback", "invocation")
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// HighWaterMark returns the value of property "ByteLengthQueuingStrategy.highWaterMark".
//
// It returns ok=false if there is no such property.
func (this ByteLengthQueuingStrategy) HighWaterMark() (ret float64, ok bool) {
	ok = js.True == bindings.GetByteLengthQueuingStrategyHighWaterMark(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Size returns the value of property "ByteLengthQueuingStrategy.size".
//
// It returns ok=false if there is no such property.
func (this ByteLengthQueuingStrategy) Size() (ret js.Func[func(arguments ...js.Any) js.Any], ok bool) {
	ok = js.True == bindings.GetByteLengthQueuingStrategySize(
		this.Ref(), js.Pointer(&ret),
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// DocumentURL returns the value of property "CSPViolationReportBody.documentURL".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) DocumentURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyDocumentURL(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Referrer returns the value of property "CSPViolationReportBody.referrer".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) Referrer() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyReferrer(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BlockedURL returns the value of property "CSPViolationReportBody.blockedURL".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) BlockedURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyBlockedURL(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EffectiveDirective returns the value of property "CSPViolationReportBody.effectiveDirective".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) EffectiveDirective() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyEffectiveDirective(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OriginalPolicy returns the value of property "CSPViolationReportBody.originalPolicy".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) OriginalPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyOriginalPolicy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SourceFile returns the value of property "CSPViolationReportBody.sourceFile".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) SourceFile() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodySourceFile(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Sample returns the value of property "CSPViolationReportBody.sample".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) Sample() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodySample(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Disposition returns the value of property "CSPViolationReportBody.disposition".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) Disposition() (ret SecurityPolicyViolationEventDisposition, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyDisposition(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StatusCode returns the value of property "CSPViolationReportBody.statusCode".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) StatusCode() (ret uint16, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyStatusCode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LineNumber returns the value of property "CSPViolationReportBody.lineNumber".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) LineNumber() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyLineNumber(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ColumnNumber returns the value of property "CSPViolationReportBody.columnNumber".
//
// It returns ok=false if there is no such property.
func (this CSPViolationReportBody) ColumnNumber() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCSPViolationReportBodyColumnNumber(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "CSPViolationReportBody.toJSON" exists.
func (this CSPViolationReportBody) HasToJSON() bool {
	return js.True == bindings.HasCSPViolationReportBodyToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "CSPViolationReportBody.toJSON".
func (this CSPViolationReportBody) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.CSPViolationReportBodyToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "CSPViolationReportBody.toJSON".
func (this CSPViolationReportBody) ToJSON() (ret js.Object) {
	bindings.CallCSPViolationReportBodyToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "CSPViolationReportBody.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSPViolationReportBody) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSPViolationReportBodyToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
func (p PropertyDefinition) UpdateFrom(ref js.Ref) {
	bindings.PropertyDefinitionJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PropertyDefinition) Update(ref js.Ref) {
	bindings.PropertyDefinitionJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CSSParserRule struct {
	ref js.Ref
}

func (this CSSParserRule) Once() CSSParserRule {
	this.Ref().Once()
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
	this.Ref().Free()
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
func (p CSSParserOptions) UpdateFrom(ref js.Ref) {
	bindings.CSSParserOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CSSParserOptions) Update(ref js.Ref) {
	bindings.CSSParserOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CSSParserValue struct {
	ref js.Ref
}

func (this CSSParserValue) Once() CSSParserValue {
	this.Ref().Once()
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
	this.Ref().Free()
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Name returns the value of property "CSSParserDeclaration.name".
//
// It returns ok=false if there is no such property.
func (this CSSParserDeclaration) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSParserDeclarationName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Body returns the value of property "CSSParserDeclaration.body".
//
// It returns ok=false if there is no such property.
func (this CSSParserDeclaration) Body() (ret js.FrozenArray[CSSParserValue], ok bool) {
	ok = js.True == bindings.GetCSSParserDeclarationBody(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToString returns true if the method "CSSParserDeclaration.toString" exists.
func (this CSSParserDeclaration) HasToString() bool {
	return js.True == bindings.HasCSSParserDeclarationToString(
		this.Ref(),
	)
}

// ToStringFunc returns the method "CSSParserDeclaration.toString".
func (this CSSParserDeclaration) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CSSParserDeclarationToStringFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "CSSParserDeclaration.toString".
func (this CSSParserDeclaration) ToString() (ret js.String) {
	bindings.CallCSSParserDeclarationToString(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "CSSParserDeclaration.toString"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSParserDeclaration) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParserDeclarationToString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
func (p WorkletOptions) UpdateFrom(ref js.Ref) {
	bindings.WorkletOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WorkletOptions) Update(ref js.Ref) {
	bindings.WorkletOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type Worklet struct {
	ref js.Ref
}

func (this Worklet) Once() Worklet {
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasAddModule returns true if the method "Worklet.addModule" exists.
func (this Worklet) HasAddModule() bool {
	return js.True == bindings.HasWorkletAddModule(
		this.Ref(),
	)
}

// AddModuleFunc returns the method "Worklet.addModule".
func (this Worklet) AddModuleFunc() (fn js.Func[func(moduleURL js.String, options WorkletOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WorkletAddModuleFunc(
			this.Ref(),
		),
	)
}

// AddModule calls the method "Worklet.addModule".
func (this Worklet) AddModule(moduleURL js.String, options WorkletOptions) (ret js.Promise[js.Void]) {
	bindings.CallWorkletAddModule(
		this.Ref(), js.Pointer(&ret),
		moduleURL.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryAddModule calls the method "Worklet.addModule"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Worklet) TryAddModule(moduleURL js.String, options WorkletOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkletAddModule(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		moduleURL.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasAddModule1 returns true if the method "Worklet.addModule" exists.
func (this Worklet) HasAddModule1() bool {
	return js.True == bindings.HasWorkletAddModule1(
		this.Ref(),
	)
}

// AddModule1Func returns the method "Worklet.addModule".
func (this Worklet) AddModule1Func() (fn js.Func[func(moduleURL js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WorkletAddModule1Func(
			this.Ref(),
		),
	)
}

// AddModule1 calls the method "Worklet.addModule".
func (this Worklet) AddModule1(moduleURL js.String) (ret js.Promise[js.Void]) {
	bindings.CallWorkletAddModule1(
		this.Ref(), js.Pointer(&ret),
		moduleURL.Ref(),
	)

	return
}

// TryAddModule1 calls the method "Worklet.addModule"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Worklet) TryAddModule1(moduleURL js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkletAddModule1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		moduleURL.Ref(),
	)

	return
}

type HighlightRegistry struct {
	ref js.Ref
}

func (this HighlightRegistry) Once() HighlightRegistry {
	this.Ref().Once()
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
	this.Ref().Free()
}

type CSS struct{}

// ElementSources returns the value of property "elementSources".
//
// The returned bool will be false if there is no such property.
func (CSS) ElementSources() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetCSSElementSources(
		js.Pointer(&ret),
	)

	return
}

// AnimationWorklet returns the value of property "animationWorklet".
//
// The returned bool will be false if there is no such property.
func (CSS) AnimationWorklet() (ret Worklet, ok bool) {
	ok = js.True == bindings.GetCSSAnimationWorklet(
		js.Pointer(&ret),
	)

	return
}

// PaintWorklet returns the value of property "paintWorklet".
//
// The returned bool will be false if there is no such property.
func (CSS) PaintWorklet() (ret Worklet, ok bool) {
	ok = js.True == bindings.GetCSSPaintWorklet(
		js.Pointer(&ret),
	)

	return
}

// LayoutWorklet returns the value of property "layoutWorklet".
//
// The returned bool will be false if there is no such property.
func (CSS) LayoutWorklet() (ret Worklet, ok bool) {
	ok = js.True == bindings.GetCSSLayoutWorklet(
		js.Pointer(&ret),
	)

	return
}

// Highlights returns the value of property "highlights".
//
// The returned bool will be false if there is no such property.
func (CSS) Highlights() (ret HighlightRegistry, ok bool) {
	ok = js.True == bindings.GetCSSHighlights(
		js.Pointer(&ret),
	)

	return
}

// HasEscape returns ture if the function "CSS.escape" exists.
func (CSS) HasEscape() bool {
	return js.True == bindings.HasCSSEscape()
}

// EscapeFunc returns the function "CSS.escape".
func (CSS) EscapeFunc() (fn js.Func[func(ident js.String) js.String]) {
	return fn.FromRef(
		bindings.CSSEscapeFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryEscape(ident js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSEscape(
		js.Pointer(&ret), js.Pointer(&exception),
		ident.Ref(),
	)
	return
}

// HasRegisterProperty returns ture if the function "CSS.registerProperty" exists.
func (CSS) HasRegisterProperty() bool {
	return js.True == bindings.HasCSSRegisterProperty()
}

// RegisterPropertyFunc returns the function "CSS.registerProperty".
func (CSS) RegisterPropertyFunc() (fn js.Func[func(definition PropertyDefinition)]) {
	return fn.FromRef(
		bindings.CSSRegisterPropertyFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryRegisterProperty(definition PropertyDefinition) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRegisterProperty(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&definition),
	)
	return
}

// HasSupports returns ture if the function "CSS.supports" exists.
func (CSS) HasSupports() bool {
	return js.True == bindings.HasCSSSupports()
}

// SupportsFunc returns the function "CSS.supports".
func (CSS) SupportsFunc() (fn js.Func[func(property js.String, value js.String) bool]) {
	return fn.FromRef(
		bindings.CSSSupportsFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TrySupports(property js.String, value js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSupports(
		js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
		value.Ref(),
	)
	return
}

// HasSupports1 returns ture if the function "CSS.supports" exists.
func (CSS) HasSupports1() bool {
	return js.True == bindings.HasCSSSupports1()
}

// Supports1Func returns the function "CSS.supports".
func (CSS) Supports1Func() (fn js.Func[func(conditionText js.String) bool]) {
	return fn.FromRef(
		bindings.CSSSupports1Func(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TrySupports1(conditionText js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSupports1(
		js.Pointer(&ret), js.Pointer(&exception),
		conditionText.Ref(),
	)
	return
}

// HasNumber returns ture if the function "CSS.number" exists.
func (CSS) HasNumber() bool {
	return js.True == bindings.HasCSSNumber()
}

// NumberFunc returns the function "CSS.number".
func (CSS) NumberFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSNumberFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryNumber(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSNumber(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasPercent returns ture if the function "CSS.percent" exists.
func (CSS) HasPercent() bool {
	return js.True == bindings.HasCSSPercent()
}

// PercentFunc returns the function "CSS.percent".
func (CSS) PercentFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSPercentFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryPercent(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPercent(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasCap returns ture if the function "CSS.cap" exists.
func (CSS) HasCap() bool {
	return js.True == bindings.HasCSSCap()
}

// CapFunc returns the function "CSS.cap".
func (CSS) CapFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCapFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryCap(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCap(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasCh returns ture if the function "CSS.ch" exists.
func (CSS) HasCh() bool {
	return js.True == bindings.HasCSSCh()
}

// ChFunc returns the function "CSS.ch".
func (CSS) ChFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSChFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryCh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasEm returns ture if the function "CSS.em" exists.
func (CSS) HasEm() bool {
	return js.True == bindings.HasCSSEm()
}

// EmFunc returns the function "CSS.em".
func (CSS) EmFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSEmFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryEm(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSEm(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasEx returns ture if the function "CSS.ex" exists.
func (CSS) HasEx() bool {
	return js.True == bindings.HasCSSEx()
}

// ExFunc returns the function "CSS.ex".
func (CSS) ExFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSExFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryEx(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSEx(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasIc returns ture if the function "CSS.ic" exists.
func (CSS) HasIc() bool {
	return js.True == bindings.HasCSSIc()
}

// IcFunc returns the function "CSS.ic".
func (CSS) IcFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSIcFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryIc(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSIc(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasLh returns ture if the function "CSS.lh" exists.
func (CSS) HasLh() bool {
	return js.True == bindings.HasCSSLh()
}

// LhFunc returns the function "CSS.lh".
func (CSS) LhFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSLhFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryLh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSLh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasRcap returns ture if the function "CSS.rcap" exists.
func (CSS) HasRcap() bool {
	return js.True == bindings.HasCSSRcap()
}

// RcapFunc returns the function "CSS.rcap".
func (CSS) RcapFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSRcapFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryRcap(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRcap(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasRch returns ture if the function "CSS.rch" exists.
func (CSS) HasRch() bool {
	return js.True == bindings.HasCSSRch()
}

// RchFunc returns the function "CSS.rch".
func (CSS) RchFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSRchFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryRch(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRch(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasRem returns ture if the function "CSS.rem" exists.
func (CSS) HasRem() bool {
	return js.True == bindings.HasCSSRem()
}

// RemFunc returns the function "CSS.rem".
func (CSS) RemFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSRemFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryRem(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRem(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasRex returns ture if the function "CSS.rex" exists.
func (CSS) HasRex() bool {
	return js.True == bindings.HasCSSRex()
}

// RexFunc returns the function "CSS.rex".
func (CSS) RexFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSRexFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryRex(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRex(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasRic returns ture if the function "CSS.ric" exists.
func (CSS) HasRic() bool {
	return js.True == bindings.HasCSSRic()
}

// RicFunc returns the function "CSS.ric".
func (CSS) RicFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSRicFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryRic(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRic(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasRlh returns ture if the function "CSS.rlh" exists.
func (CSS) HasRlh() bool {
	return js.True == bindings.HasCSSRlh()
}

// RlhFunc returns the function "CSS.rlh".
func (CSS) RlhFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSRlhFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryRlh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRlh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasVw returns ture if the function "CSS.vw" exists.
func (CSS) HasVw() bool {
	return js.True == bindings.HasCSSVw()
}

// VwFunc returns the function "CSS.vw".
func (CSS) VwFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSVwFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryVw(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSVw(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasVh returns ture if the function "CSS.vh" exists.
func (CSS) HasVh() bool {
	return js.True == bindings.HasCSSVh()
}

// VhFunc returns the function "CSS.vh".
func (CSS) VhFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSVhFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryVh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSVh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasVi returns ture if the function "CSS.vi" exists.
func (CSS) HasVi() bool {
	return js.True == bindings.HasCSSVi()
}

// ViFunc returns the function "CSS.vi".
func (CSS) ViFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSViFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryVi(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSVi(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasVb returns ture if the function "CSS.vb" exists.
func (CSS) HasVb() bool {
	return js.True == bindings.HasCSSVb()
}

// VbFunc returns the function "CSS.vb".
func (CSS) VbFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSVbFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryVb(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSVb(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasVmin returns ture if the function "CSS.vmin" exists.
func (CSS) HasVmin() bool {
	return js.True == bindings.HasCSSVmin()
}

// VminFunc returns the function "CSS.vmin".
func (CSS) VminFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSVminFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryVmin(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSVmin(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasVmax returns ture if the function "CSS.vmax" exists.
func (CSS) HasVmax() bool {
	return js.True == bindings.HasCSSVmax()
}

// VmaxFunc returns the function "CSS.vmax".
func (CSS) VmaxFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSVmaxFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryVmax(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSVmax(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasSvw returns ture if the function "CSS.svw" exists.
func (CSS) HasSvw() bool {
	return js.True == bindings.HasCSSSvw()
}

// SvwFunc returns the function "CSS.svw".
func (CSS) SvwFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSSvwFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TrySvw(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSvw(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasSvh returns ture if the function "CSS.svh" exists.
func (CSS) HasSvh() bool {
	return js.True == bindings.HasCSSSvh()
}

// SvhFunc returns the function "CSS.svh".
func (CSS) SvhFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSSvhFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TrySvh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSvh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasSvi returns ture if the function "CSS.svi" exists.
func (CSS) HasSvi() bool {
	return js.True == bindings.HasCSSSvi()
}

// SviFunc returns the function "CSS.svi".
func (CSS) SviFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSSviFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TrySvi(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSvi(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasSvb returns ture if the function "CSS.svb" exists.
func (CSS) HasSvb() bool {
	return js.True == bindings.HasCSSSvb()
}

// SvbFunc returns the function "CSS.svb".
func (CSS) SvbFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSSvbFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TrySvb(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSvb(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasSvmin returns ture if the function "CSS.svmin" exists.
func (CSS) HasSvmin() bool {
	return js.True == bindings.HasCSSSvmin()
}

// SvminFunc returns the function "CSS.svmin".
func (CSS) SvminFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSSvminFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TrySvmin(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSvmin(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasSvmax returns ture if the function "CSS.svmax" exists.
func (CSS) HasSvmax() bool {
	return js.True == bindings.HasCSSSvmax()
}

// SvmaxFunc returns the function "CSS.svmax".
func (CSS) SvmaxFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSSvmaxFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TrySvmax(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSSvmax(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasLvw returns ture if the function "CSS.lvw" exists.
func (CSS) HasLvw() bool {
	return js.True == bindings.HasCSSLvw()
}

// LvwFunc returns the function "CSS.lvw".
func (CSS) LvwFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSLvwFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryLvw(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSLvw(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasLvh returns ture if the function "CSS.lvh" exists.
func (CSS) HasLvh() bool {
	return js.True == bindings.HasCSSLvh()
}

// LvhFunc returns the function "CSS.lvh".
func (CSS) LvhFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSLvhFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryLvh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSLvh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasLvi returns ture if the function "CSS.lvi" exists.
func (CSS) HasLvi() bool {
	return js.True == bindings.HasCSSLvi()
}

// LviFunc returns the function "CSS.lvi".
func (CSS) LviFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSLviFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryLvi(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSLvi(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasLvb returns ture if the function "CSS.lvb" exists.
func (CSS) HasLvb() bool {
	return js.True == bindings.HasCSSLvb()
}

// LvbFunc returns the function "CSS.lvb".
func (CSS) LvbFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSLvbFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryLvb(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSLvb(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasLvmin returns ture if the function "CSS.lvmin" exists.
func (CSS) HasLvmin() bool {
	return js.True == bindings.HasCSSLvmin()
}

// LvminFunc returns the function "CSS.lvmin".
func (CSS) LvminFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSLvminFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryLvmin(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSLvmin(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasLvmax returns ture if the function "CSS.lvmax" exists.
func (CSS) HasLvmax() bool {
	return js.True == bindings.HasCSSLvmax()
}

// LvmaxFunc returns the function "CSS.lvmax".
func (CSS) LvmaxFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSLvmaxFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryLvmax(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSLvmax(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasDvw returns ture if the function "CSS.dvw" exists.
func (CSS) HasDvw() bool {
	return js.True == bindings.HasCSSDvw()
}

// DvwFunc returns the function "CSS.dvw".
func (CSS) DvwFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDvwFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryDvw(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDvw(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasDvh returns ture if the function "CSS.dvh" exists.
func (CSS) HasDvh() bool {
	return js.True == bindings.HasCSSDvh()
}

// DvhFunc returns the function "CSS.dvh".
func (CSS) DvhFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDvhFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryDvh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDvh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasDvi returns ture if the function "CSS.dvi" exists.
func (CSS) HasDvi() bool {
	return js.True == bindings.HasCSSDvi()
}

// DviFunc returns the function "CSS.dvi".
func (CSS) DviFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDviFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryDvi(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDvi(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasDvb returns ture if the function "CSS.dvb" exists.
func (CSS) HasDvb() bool {
	return js.True == bindings.HasCSSDvb()
}

// DvbFunc returns the function "CSS.dvb".
func (CSS) DvbFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDvbFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryDvb(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDvb(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasDvmin returns ture if the function "CSS.dvmin" exists.
func (CSS) HasDvmin() bool {
	return js.True == bindings.HasCSSDvmin()
}

// DvminFunc returns the function "CSS.dvmin".
func (CSS) DvminFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDvminFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryDvmin(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDvmin(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasDvmax returns ture if the function "CSS.dvmax" exists.
func (CSS) HasDvmax() bool {
	return js.True == bindings.HasCSSDvmax()
}

// DvmaxFunc returns the function "CSS.dvmax".
func (CSS) DvmaxFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDvmaxFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryDvmax(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDvmax(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasCqw returns ture if the function "CSS.cqw" exists.
func (CSS) HasCqw() bool {
	return js.True == bindings.HasCSSCqw()
}

// CqwFunc returns the function "CSS.cqw".
func (CSS) CqwFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCqwFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryCqw(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCqw(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasCqh returns ture if the function "CSS.cqh" exists.
func (CSS) HasCqh() bool {
	return js.True == bindings.HasCSSCqh()
}

// CqhFunc returns the function "CSS.cqh".
func (CSS) CqhFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCqhFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryCqh(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCqh(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasCqi returns ture if the function "CSS.cqi" exists.
func (CSS) HasCqi() bool {
	return js.True == bindings.HasCSSCqi()
}

// CqiFunc returns the function "CSS.cqi".
func (CSS) CqiFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCqiFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryCqi(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCqi(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasCqb returns ture if the function "CSS.cqb" exists.
func (CSS) HasCqb() bool {
	return js.True == bindings.HasCSSCqb()
}

// CqbFunc returns the function "CSS.cqb".
func (CSS) CqbFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCqbFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryCqb(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCqb(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasCqmin returns ture if the function "CSS.cqmin" exists.
func (CSS) HasCqmin() bool {
	return js.True == bindings.HasCSSCqmin()
}

// CqminFunc returns the function "CSS.cqmin".
func (CSS) CqminFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCqminFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryCqmin(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCqmin(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasCqmax returns ture if the function "CSS.cqmax" exists.
func (CSS) HasCqmax() bool {
	return js.True == bindings.HasCSSCqmax()
}

// CqmaxFunc returns the function "CSS.cqmax".
func (CSS) CqmaxFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCqmaxFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryCqmax(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCqmax(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasCm returns ture if the function "CSS.cm" exists.
func (CSS) HasCm() bool {
	return js.True == bindings.HasCSSCm()
}

// CmFunc returns the function "CSS.cm".
func (CSS) CmFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCmFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryCm(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSCm(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasMm returns ture if the function "CSS.mm" exists.
func (CSS) HasMm() bool {
	return js.True == bindings.HasCSSMm()
}

// MmFunc returns the function "CSS.mm".
func (CSS) MmFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSMmFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryMm(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSMm(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasQ returns ture if the function "CSS.Q" exists.
func (CSS) HasQ() bool {
	return js.True == bindings.HasCSSQ()
}

// QFunc returns the function "CSS.Q".
func (CSS) QFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSQFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryQ(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSQ(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasIn returns ture if the function "CSS.in" exists.
func (CSS) HasIn() bool {
	return js.True == bindings.HasCSSIn()
}

// InFunc returns the function "CSS.in".
func (CSS) InFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSInFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryIn(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSIn(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasPt returns ture if the function "CSS.pt" exists.
func (CSS) HasPt() bool {
	return js.True == bindings.HasCSSPt()
}

// PtFunc returns the function "CSS.pt".
func (CSS) PtFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSPtFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryPt(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPt(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasPc returns ture if the function "CSS.pc" exists.
func (CSS) HasPc() bool {
	return js.True == bindings.HasCSSPc()
}

// PcFunc returns the function "CSS.pc".
func (CSS) PcFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSPcFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryPc(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPc(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasPx returns ture if the function "CSS.px" exists.
func (CSS) HasPx() bool {
	return js.True == bindings.HasCSSPx()
}

// PxFunc returns the function "CSS.px".
func (CSS) PxFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSPxFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryPx(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPx(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasDeg returns ture if the function "CSS.deg" exists.
func (CSS) HasDeg() bool {
	return js.True == bindings.HasCSSDeg()
}

// DegFunc returns the function "CSS.deg".
func (CSS) DegFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDegFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryDeg(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDeg(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasGrad returns ture if the function "CSS.grad" exists.
func (CSS) HasGrad() bool {
	return js.True == bindings.HasCSSGrad()
}

// GradFunc returns the function "CSS.grad".
func (CSS) GradFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSGradFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryGrad(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSGrad(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasRad returns ture if the function "CSS.rad" exists.
func (CSS) HasRad() bool {
	return js.True == bindings.HasCSSRad()
}

// RadFunc returns the function "CSS.rad".
func (CSS) RadFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSRadFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryRad(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRad(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasTurn returns ture if the function "CSS.turn" exists.
func (CSS) HasTurn() bool {
	return js.True == bindings.HasCSSTurn()
}

// TurnFunc returns the function "CSS.turn".
func (CSS) TurnFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSTurnFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryTurn(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSTurn(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasS returns ture if the function "CSS.s" exists.
func (CSS) HasS() bool {
	return js.True == bindings.HasCSSS()
}

// SFunc returns the function "CSS.s".
func (CSS) SFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSSFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryS(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSS(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasMs returns ture if the function "CSS.ms" exists.
func (CSS) HasMs() bool {
	return js.True == bindings.HasCSSMs()
}

// MsFunc returns the function "CSS.ms".
func (CSS) MsFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSMsFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryMs(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSMs(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasHz returns ture if the function "CSS.Hz" exists.
func (CSS) HasHz() bool {
	return js.True == bindings.HasCSSHz()
}

// HzFunc returns the function "CSS.Hz".
func (CSS) HzFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSHzFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryHz(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSHz(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasKHz returns ture if the function "CSS.kHz" exists.
func (CSS) HasKHz() bool {
	return js.True == bindings.HasCSSKHz()
}

// KHzFunc returns the function "CSS.kHz".
func (CSS) KHzFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSKHzFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryKHz(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSKHz(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasDpi returns ture if the function "CSS.dpi" exists.
func (CSS) HasDpi() bool {
	return js.True == bindings.HasCSSDpi()
}

// DpiFunc returns the function "CSS.dpi".
func (CSS) DpiFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDpiFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryDpi(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDpi(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasDpcm returns ture if the function "CSS.dpcm" exists.
func (CSS) HasDpcm() bool {
	return js.True == bindings.HasCSSDpcm()
}

// DpcmFunc returns the function "CSS.dpcm".
func (CSS) DpcmFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDpcmFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryDpcm(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDpcm(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasDppx returns ture if the function "CSS.dppx" exists.
func (CSS) HasDppx() bool {
	return js.True == bindings.HasCSSDppx()
}

// DppxFunc returns the function "CSS.dppx".
func (CSS) DppxFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDppxFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryDppx(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSDppx(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasFr returns ture if the function "CSS.fr" exists.
func (CSS) HasFr() bool {
	return js.True == bindings.HasCSSFr()
}

// FrFunc returns the function "CSS.fr".
func (CSS) FrFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSFrFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryFr(value float64) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSFr(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)
	return
}

// HasParseStylesheet returns ture if the function "CSS.parseStylesheet" exists.
func (CSS) HasParseStylesheet() bool {
	return js.True == bindings.HasCSSParseStylesheet()
}

// ParseStylesheetFunc returns the function "CSS.parseStylesheet".
func (CSS) ParseStylesheetFunc() (fn js.Func[func(css CSSStringSource, options CSSParserOptions) js.Promise[js.Array[CSSParserRule]]]) {
	return fn.FromRef(
		bindings.CSSParseStylesheetFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryParseStylesheet(css CSSStringSource, options CSSParserOptions) (ret js.Promise[js.Array[CSSParserRule]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseStylesheet(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
		js.Pointer(&options),
	)
	return
}

// HasParseStylesheet1 returns ture if the function "CSS.parseStylesheet" exists.
func (CSS) HasParseStylesheet1() bool {
	return js.True == bindings.HasCSSParseStylesheet1()
}

// ParseStylesheet1Func returns the function "CSS.parseStylesheet".
func (CSS) ParseStylesheet1Func() (fn js.Func[func(css CSSStringSource) js.Promise[js.Array[CSSParserRule]]]) {
	return fn.FromRef(
		bindings.CSSParseStylesheet1Func(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryParseStylesheet1(css CSSStringSource) (ret js.Promise[js.Array[CSSParserRule]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseStylesheet1(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
	)
	return
}

// HasParseRuleList returns ture if the function "CSS.parseRuleList" exists.
func (CSS) HasParseRuleList() bool {
	return js.True == bindings.HasCSSParseRuleList()
}

// ParseRuleListFunc returns the function "CSS.parseRuleList".
func (CSS) ParseRuleListFunc() (fn js.Func[func(css CSSStringSource, options CSSParserOptions) js.Promise[js.Array[CSSParserRule]]]) {
	return fn.FromRef(
		bindings.CSSParseRuleListFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryParseRuleList(css CSSStringSource, options CSSParserOptions) (ret js.Promise[js.Array[CSSParserRule]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseRuleList(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
		js.Pointer(&options),
	)
	return
}

// HasParseRuleList1 returns ture if the function "CSS.parseRuleList" exists.
func (CSS) HasParseRuleList1() bool {
	return js.True == bindings.HasCSSParseRuleList1()
}

// ParseRuleList1Func returns the function "CSS.parseRuleList".
func (CSS) ParseRuleList1Func() (fn js.Func[func(css CSSStringSource) js.Promise[js.Array[CSSParserRule]]]) {
	return fn.FromRef(
		bindings.CSSParseRuleList1Func(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryParseRuleList1(css CSSStringSource) (ret js.Promise[js.Array[CSSParserRule]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseRuleList1(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
	)
	return
}

// HasParseRule returns ture if the function "CSS.parseRule" exists.
func (CSS) HasParseRule() bool {
	return js.True == bindings.HasCSSParseRule()
}

// ParseRuleFunc returns the function "CSS.parseRule".
func (CSS) ParseRuleFunc() (fn js.Func[func(css CSSStringSource, options CSSParserOptions) js.Promise[CSSParserRule]]) {
	return fn.FromRef(
		bindings.CSSParseRuleFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryParseRule(css CSSStringSource, options CSSParserOptions) (ret js.Promise[CSSParserRule], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseRule(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
		js.Pointer(&options),
	)
	return
}

// HasParseRule1 returns ture if the function "CSS.parseRule" exists.
func (CSS) HasParseRule1() bool {
	return js.True == bindings.HasCSSParseRule1()
}

// ParseRule1Func returns the function "CSS.parseRule".
func (CSS) ParseRule1Func() (fn js.Func[func(css CSSStringSource) js.Promise[CSSParserRule]]) {
	return fn.FromRef(
		bindings.CSSParseRule1Func(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryParseRule1(css CSSStringSource) (ret js.Promise[CSSParserRule], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseRule1(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
	)
	return
}

// HasParseDeclarationList returns ture if the function "CSS.parseDeclarationList" exists.
func (CSS) HasParseDeclarationList() bool {
	return js.True == bindings.HasCSSParseDeclarationList()
}

// ParseDeclarationListFunc returns the function "CSS.parseDeclarationList".
func (CSS) ParseDeclarationListFunc() (fn js.Func[func(css CSSStringSource, options CSSParserOptions) js.Promise[js.Array[CSSParserRule]]]) {
	return fn.FromRef(
		bindings.CSSParseDeclarationListFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryParseDeclarationList(css CSSStringSource, options CSSParserOptions) (ret js.Promise[js.Array[CSSParserRule]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseDeclarationList(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
		js.Pointer(&options),
	)
	return
}

// HasParseDeclarationList1 returns ture if the function "CSS.parseDeclarationList" exists.
func (CSS) HasParseDeclarationList1() bool {
	return js.True == bindings.HasCSSParseDeclarationList1()
}

// ParseDeclarationList1Func returns the function "CSS.parseDeclarationList".
func (CSS) ParseDeclarationList1Func() (fn js.Func[func(css CSSStringSource) js.Promise[js.Array[CSSParserRule]]]) {
	return fn.FromRef(
		bindings.CSSParseDeclarationList1Func(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryParseDeclarationList1(css CSSStringSource) (ret js.Promise[js.Array[CSSParserRule]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseDeclarationList1(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
	)
	return
}

// HasParseDeclaration returns ture if the function "CSS.parseDeclaration" exists.
func (CSS) HasParseDeclaration() bool {
	return js.True == bindings.HasCSSParseDeclaration()
}

// ParseDeclarationFunc returns the function "CSS.parseDeclaration".
func (CSS) ParseDeclarationFunc() (fn js.Func[func(css js.String, options CSSParserOptions) CSSParserDeclaration]) {
	return fn.FromRef(
		bindings.CSSParseDeclarationFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryParseDeclaration(css js.String, options CSSParserOptions) (ret CSSParserDeclaration, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseDeclaration(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
		js.Pointer(&options),
	)
	return
}

// HasParseDeclaration1 returns ture if the function "CSS.parseDeclaration" exists.
func (CSS) HasParseDeclaration1() bool {
	return js.True == bindings.HasCSSParseDeclaration1()
}

// ParseDeclaration1Func returns the function "CSS.parseDeclaration".
func (CSS) ParseDeclaration1Func() (fn js.Func[func(css js.String) CSSParserDeclaration]) {
	return fn.FromRef(
		bindings.CSSParseDeclaration1Func(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryParseDeclaration1(css js.String) (ret CSSParserDeclaration, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseDeclaration1(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
	)
	return
}

// HasParseValue returns ture if the function "CSS.parseValue" exists.
func (CSS) HasParseValue() bool {
	return js.True == bindings.HasCSSParseValue()
}

// ParseValueFunc returns the function "CSS.parseValue".
func (CSS) ParseValueFunc() (fn js.Func[func(css js.String) CSSToken]) {
	return fn.FromRef(
		bindings.CSSParseValueFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryParseValue(css js.String) (ret CSSToken, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseValue(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
	)
	return
}

// HasParseValueList returns ture if the function "CSS.parseValueList" exists.
func (CSS) HasParseValueList() bool {
	return js.True == bindings.HasCSSParseValueList()
}

// ParseValueListFunc returns the function "CSS.parseValueList".
func (CSS) ParseValueListFunc() (fn js.Func[func(css js.String) js.Array[CSSToken]]) {
	return fn.FromRef(
		bindings.CSSParseValueListFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
// the catch clause.
func (CSS) TryParseValueList(css js.String) (ret js.Array[CSSToken], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParseValueList(
		js.Pointer(&ret), js.Pointer(&exception),
		css.Ref(),
	)
	return
}

// HasParseCommaValueList returns ture if the function "CSS.parseCommaValueList" exists.
func (CSS) HasParseCommaValueList() bool {
	return js.True == bindings.HasCSSParseCommaValueList()
}

// ParseCommaValueListFunc returns the function "CSS.parseCommaValueList".
func (CSS) ParseCommaValueListFunc() (fn js.Func[func(css js.String) js.Array[js.Array[CSSToken]]]) {
	return fn.FromRef(
		bindings.CSSParseCommaValueListFunc(),
	)
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
// in a try/catch block and returns (_, err, ok = true) when it went though
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// AnimationName returns the value of property "CSSAnimation.animationName".
//
// It returns ok=false if there is no such property.
func (this CSSAnimation) AnimationName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSAnimationAnimationName(
		this.Ref(), js.Pointer(&ret),
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
