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
// The returned bool will be false if there is no such property.
func (this Bluetooth) ReferringDevice() (BluetoothDevice, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothReferringDevice(
		this.Ref(), js.Pointer(&_ok),
	)
	return BluetoothDevice{}.FromRef(_ret), _ok
}

// GetAvailability calls the method "Bluetooth.getAvailability".
//
// The returned bool will be false if there is no such method.
func (this Bluetooth) GetAvailability() (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothGetAvailability(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// GetAvailabilityFunc returns the method "Bluetooth.getAvailability".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Bluetooth) GetAvailabilityFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.BluetoothGetAvailabilityFunc(
			this.Ref(),
		),
	)
}

// GetDevices calls the method "Bluetooth.getDevices".
//
// The returned bool will be false if there is no such method.
func (this Bluetooth) GetDevices() (js.Promise[js.Array[BluetoothDevice]], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothGetDevices(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[BluetoothDevice]]{}.FromRef(_ret), _ok
}

// GetDevicesFunc returns the method "Bluetooth.getDevices".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Bluetooth) GetDevicesFunc() (fn js.Func[func() js.Promise[js.Array[BluetoothDevice]]]) {
	return fn.FromRef(
		bindings.BluetoothGetDevicesFunc(
			this.Ref(),
		),
	)
}

// RequestDevice calls the method "Bluetooth.requestDevice".
//
// The returned bool will be false if there is no such method.
func (this Bluetooth) RequestDevice(options RequestDeviceOptions) (js.Promise[BluetoothDevice], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRequestDevice(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[BluetoothDevice]{}.FromRef(_ret), _ok
}

// RequestDeviceFunc returns the method "Bluetooth.requestDevice".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Bluetooth) RequestDeviceFunc() (fn js.Func[func(options RequestDeviceOptions) js.Promise[BluetoothDevice]]) {
	return fn.FromRef(
		bindings.BluetoothRequestDeviceFunc(
			this.Ref(),
		),
	)
}

// RequestDevice1 calls the method "Bluetooth.requestDevice".
//
// The returned bool will be false if there is no such method.
func (this Bluetooth) RequestDevice1() (js.Promise[BluetoothDevice], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRequestDevice1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[BluetoothDevice]{}.FromRef(_ret), _ok
}

// RequestDevice1Func returns the method "Bluetooth.requestDevice".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Bluetooth) RequestDevice1Func() (fn js.Func[func() js.Promise[BluetoothDevice]]) {
	return fn.FromRef(
		bindings.BluetoothRequestDevice1Func(
			this.Ref(),
		),
	)
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

func NewBluetoothAdvertisingEvent(typ js.String, init BluetoothAdvertisingEventInit) BluetoothAdvertisingEvent {
	return BluetoothAdvertisingEvent{}.FromRef(
		bindings.NewBluetoothAdvertisingEventByBluetoothAdvertisingEvent(
			typ.Ref(),
			js.Pointer(&init)),
	)
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
// The returned bool will be false if there is no such property.
func (this BluetoothAdvertisingEvent) Device() (BluetoothDevice, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothAdvertisingEventDevice(
		this.Ref(), js.Pointer(&_ok),
	)
	return BluetoothDevice{}.FromRef(_ret), _ok
}

// Uuids returns the value of property "BluetoothAdvertisingEvent.uuids".
//
// The returned bool will be false if there is no such property.
func (this BluetoothAdvertisingEvent) Uuids() (js.FrozenArray[UUID], bool) {
	var _ok bool
	_ret := bindings.GetBluetoothAdvertisingEventUuids(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[UUID]{}.FromRef(_ret), _ok
}

// Name returns the value of property "BluetoothAdvertisingEvent.name".
//
// The returned bool will be false if there is no such property.
func (this BluetoothAdvertisingEvent) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothAdvertisingEventName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Appearance returns the value of property "BluetoothAdvertisingEvent.appearance".
//
// The returned bool will be false if there is no such property.
func (this BluetoothAdvertisingEvent) Appearance() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothAdvertisingEventAppearance(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// TxPower returns the value of property "BluetoothAdvertisingEvent.txPower".
//
// The returned bool will be false if there is no such property.
func (this BluetoothAdvertisingEvent) TxPower() (int8, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothAdvertisingEventTxPower(
		this.Ref(), js.Pointer(&_ok),
	)
	return int8(_ret), _ok
}

// Rssi returns the value of property "BluetoothAdvertisingEvent.rssi".
//
// The returned bool will be false if there is no such property.
func (this BluetoothAdvertisingEvent) Rssi() (int8, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothAdvertisingEventRssi(
		this.Ref(), js.Pointer(&_ok),
	)
	return int8(_ret), _ok
}

// ManufacturerData returns the value of property "BluetoothAdvertisingEvent.manufacturerData".
//
// The returned bool will be false if there is no such property.
func (this BluetoothAdvertisingEvent) ManufacturerData() (BluetoothManufacturerDataMap, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothAdvertisingEventManufacturerData(
		this.Ref(), js.Pointer(&_ok),
	)
	return BluetoothManufacturerDataMap{}.FromRef(_ret), _ok
}

// ServiceData returns the value of property "BluetoothAdvertisingEvent.serviceData".
//
// The returned bool will be false if there is no such property.
func (this BluetoothAdvertisingEvent) ServiceData() (BluetoothServiceDataMap, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothAdvertisingEventServiceData(
		this.Ref(), js.Pointer(&_ok),
	)
	return BluetoothServiceDataMap{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this BluetoothPermissionResult) Devices() (js.FrozenArray[BluetoothDevice], bool) {
	var _ok bool
	_ret := bindings.GetBluetoothPermissionResultDevices(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[BluetoothDevice]{}.FromRef(_ret), _ok
}

// Devices sets the value of property "BluetoothPermissionResult.devices" to val.
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

// GetService calls the staticmethod "BluetoothUUID.getService".
//
// The returned bool will be false if there is no such method.
func (this BluetoothUUID) GetService(name OneOf_String_Uint32) (UUID, bool) {
	var _ok bool
	_ret := bindings.CallBluetoothUUIDGetService(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return UUID{}.FromRef(_ret), _ok
}

// GetServiceFunc returns the staticmethod "BluetoothUUID.getService".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothUUID) GetServiceFunc() (fn js.Func[func(name OneOf_String_Uint32) UUID]) {
	return fn.FromRef(
		bindings.BluetoothUUIDGetServiceFunc(
			this.Ref(),
		),
	)
}

// GetCharacteristic calls the staticmethod "BluetoothUUID.getCharacteristic".
//
// The returned bool will be false if there is no such method.
func (this BluetoothUUID) GetCharacteristic(name OneOf_String_Uint32) (UUID, bool) {
	var _ok bool
	_ret := bindings.CallBluetoothUUIDGetCharacteristic(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return UUID{}.FromRef(_ret), _ok
}

// GetCharacteristicFunc returns the staticmethod "BluetoothUUID.getCharacteristic".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothUUID) GetCharacteristicFunc() (fn js.Func[func(name OneOf_String_Uint32) UUID]) {
	return fn.FromRef(
		bindings.BluetoothUUIDGetCharacteristicFunc(
			this.Ref(),
		),
	)
}

// GetDescriptor calls the staticmethod "BluetoothUUID.getDescriptor".
//
// The returned bool will be false if there is no such method.
func (this BluetoothUUID) GetDescriptor(name OneOf_String_Uint32) (UUID, bool) {
	var _ok bool
	_ret := bindings.CallBluetoothUUIDGetDescriptor(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return UUID{}.FromRef(_ret), _ok
}

// GetDescriptorFunc returns the staticmethod "BluetoothUUID.getDescriptor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothUUID) GetDescriptorFunc() (fn js.Func[func(name OneOf_String_Uint32) UUID]) {
	return fn.FromRef(
		bindings.BluetoothUUIDGetDescriptorFunc(
			this.Ref(),
		),
	)
}

// CanonicalUUID calls the staticmethod "BluetoothUUID.canonicalUUID".
//
// The returned bool will be false if there is no such method.
func (this BluetoothUUID) CanonicalUUID(alias uint32) (UUID, bool) {
	var _ok bool
	_ret := bindings.CallBluetoothUUIDCanonicalUUID(
		this.Ref(), js.Pointer(&_ok),
		uint32(alias),
	)

	return UUID{}.FromRef(_ret), _ok
}

// CanonicalUUIDFunc returns the staticmethod "BluetoothUUID.canonicalUUID".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothUUID) CanonicalUUIDFunc() (fn js.Func[func(alias uint32) UUID]) {
	return fn.FromRef(
		bindings.BluetoothUUIDCanonicalUUIDFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this IntrinsicSizes) MinContentSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetIntrinsicSizesMinContentSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// MaxContentSize returns the value of property "IntrinsicSizes.maxContentSize".
//
// The returned bool will be false if there is no such property.
func (this IntrinsicSizes) MaxContentSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetIntrinsicSizesMaxContentSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this LayoutFragment) InlineSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutFragmentInlineSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// BlockSize returns the value of property "LayoutFragment.blockSize".
//
// The returned bool will be false if there is no such property.
func (this LayoutFragment) BlockSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutFragmentBlockSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// InlineOffset returns the value of property "LayoutFragment.inlineOffset".
//
// The returned bool will be false if there is no such property.
func (this LayoutFragment) InlineOffset() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutFragmentInlineOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// InlineOffset sets the value of property "LayoutFragment.inlineOffset" to val.
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
// The returned bool will be false if there is no such property.
func (this LayoutFragment) BlockOffset() (float64, bool) {
	var _ok bool
	_ret := bindings.GetLayoutFragmentBlockOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// BlockOffset sets the value of property "LayoutFragment.blockOffset" to val.
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
// The returned bool will be false if there is no such property.
func (this LayoutFragment) Data() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetLayoutFragmentData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// BreakToken returns the value of property "LayoutFragment.breakToken".
//
// The returned bool will be false if there is no such property.
func (this LayoutFragment) BreakToken() (ChildBreakToken, bool) {
	var _ok bool
	_ret := bindings.GetLayoutFragmentBreakToken(
		this.Ref(), js.Pointer(&_ok),
	)
	return ChildBreakToken{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this LayoutChild) StyleMap() (StylePropertyMapReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetLayoutChildStyleMap(
		this.Ref(), js.Pointer(&_ok),
	)
	return StylePropertyMapReadOnly{}.FromRef(_ret), _ok
}

// IntrinsicSizes calls the method "LayoutChild.intrinsicSizes".
//
// The returned bool will be false if there is no such method.
func (this LayoutChild) IntrinsicSizes() (js.Promise[IntrinsicSizes], bool) {
	var _ok bool
	_ret := bindings.CallLayoutChildIntrinsicSizes(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[IntrinsicSizes]{}.FromRef(_ret), _ok
}

// IntrinsicSizesFunc returns the method "LayoutChild.intrinsicSizes".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this LayoutChild) IntrinsicSizesFunc() (fn js.Func[func() js.Promise[IntrinsicSizes]]) {
	return fn.FromRef(
		bindings.LayoutChildIntrinsicSizesFunc(
			this.Ref(),
		),
	)
}

// LayoutNextFragment calls the method "LayoutChild.layoutNextFragment".
//
// The returned bool will be false if there is no such method.
func (this LayoutChild) LayoutNextFragment(constraints LayoutConstraintsOptions, breakToken ChildBreakToken) (js.Promise[LayoutFragment], bool) {
	var _ok bool
	_ret := bindings.CallLayoutChildLayoutNextFragment(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&constraints),
		breakToken.Ref(),
	)

	return js.Promise[LayoutFragment]{}.FromRef(_ret), _ok
}

// LayoutNextFragmentFunc returns the method "LayoutChild.layoutNextFragment".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this LayoutChild) LayoutNextFragmentFunc() (fn js.Func[func(constraints LayoutConstraintsOptions, breakToken ChildBreakToken) js.Promise[LayoutFragment]]) {
	return fn.FromRef(
		bindings.LayoutChildLayoutNextFragmentFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this ChildBreakToken) BreakType() (BreakType, bool) {
	var _ok bool
	_ret := bindings.GetChildBreakTokenBreakType(
		this.Ref(), js.Pointer(&_ok),
	)
	return BreakType(_ret), _ok
}

// Child returns the value of property "ChildBreakToken.child".
//
// The returned bool will be false if there is no such property.
func (this ChildBreakToken) Child() (LayoutChild, bool) {
	var _ok bool
	_ret := bindings.GetChildBreakTokenChild(
		this.Ref(), js.Pointer(&_ok),
	)
	return LayoutChild{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this BreakToken) ChildBreakTokens() (js.FrozenArray[ChildBreakToken], bool) {
	var _ok bool
	_ret := bindings.GetBreakTokenChildBreakTokens(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[ChildBreakToken]{}.FromRef(_ret), _ok
}

// Data returns the value of property "BreakToken.data".
//
// The returned bool will be false if there is no such property.
func (this BreakToken) Data() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetBreakTokenData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
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

func NewBroadcastChannel(name js.String) BroadcastChannel {
	return BroadcastChannel{}.FromRef(
		bindings.NewBroadcastChannelByBroadcastChannel(
			name.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this BroadcastChannel) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetBroadcastChannelName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// PostMessage calls the method "BroadcastChannel.postMessage".
//
// The returned bool will be false if there is no such method.
func (this BroadcastChannel) PostMessage(message js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallBroadcastChannelPostMessage(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessageFunc returns the method "BroadcastChannel.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BroadcastChannel) PostMessageFunc() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.BroadcastChannelPostMessageFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "BroadcastChannel.close".
//
// The returned bool will be false if there is no such method.
func (this BroadcastChannel) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallBroadcastChannelClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "BroadcastChannel.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BroadcastChannel) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.BroadcastChannelCloseFunc(
			this.Ref(),
		),
	)
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

// FromElement calls the staticmethod "CropTarget.fromElement".
//
// The returned bool will be false if there is no such method.
func (this CropTarget) FromElement(element Element) (js.Promise[CropTarget], bool) {
	var _ok bool
	_ret := bindings.CallCropTargetFromElement(
		this.Ref(), js.Pointer(&_ok),
		element.Ref(),
	)

	return js.Promise[CropTarget]{}.FromRef(_ret), _ok
}

// FromElementFunc returns the staticmethod "CropTarget.fromElement".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CropTarget) FromElementFunc() (fn js.Func[func(element Element) js.Promise[CropTarget]]) {
	return fn.FromRef(
		bindings.CropTargetFromElementFunc(
			this.Ref(),
		),
	)
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

// FromElement calls the staticmethod "RestrictionTarget.fromElement".
//
// The returned bool will be false if there is no such method.
func (this RestrictionTarget) FromElement(element Element) (js.Promise[RestrictionTarget], bool) {
	var _ok bool
	_ret := bindings.CallRestrictionTargetFromElement(
		this.Ref(), js.Pointer(&_ok),
		element.Ref(),
	)

	return js.Promise[RestrictionTarget]{}.FromRef(_ret), _ok
}

// FromElementFunc returns the staticmethod "RestrictionTarget.fromElement".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RestrictionTarget) FromElementFunc() (fn js.Func[func(element Element) js.Promise[RestrictionTarget]]) {
	return fn.FromRef(
		bindings.RestrictionTargetFromElementFunc(
			this.Ref(),
		),
	)
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

// CropTo calls the method "BrowserCaptureMediaStreamTrack.cropTo".
//
// The returned bool will be false if there is no such method.
func (this BrowserCaptureMediaStreamTrack) CropTo(cropTarget CropTarget) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallBrowserCaptureMediaStreamTrackCropTo(
		this.Ref(), js.Pointer(&_ok),
		cropTarget.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CropToFunc returns the method "BrowserCaptureMediaStreamTrack.cropTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BrowserCaptureMediaStreamTrack) CropToFunc() (fn js.Func[func(cropTarget CropTarget) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BrowserCaptureMediaStreamTrackCropToFunc(
			this.Ref(),
		),
	)
}

// Clone calls the method "BrowserCaptureMediaStreamTrack.clone".
//
// The returned bool will be false if there is no such method.
func (this BrowserCaptureMediaStreamTrack) Clone() (BrowserCaptureMediaStreamTrack, bool) {
	var _ok bool
	_ret := bindings.CallBrowserCaptureMediaStreamTrackClone(
		this.Ref(), js.Pointer(&_ok),
	)

	return BrowserCaptureMediaStreamTrack{}.FromRef(_ret), _ok
}

// CloneFunc returns the method "BrowserCaptureMediaStreamTrack.clone".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BrowserCaptureMediaStreamTrack) CloneFunc() (fn js.Func[func() BrowserCaptureMediaStreamTrack]) {
	return fn.FromRef(
		bindings.BrowserCaptureMediaStreamTrackCloneFunc(
			this.Ref(),
		),
	)
}

// RestrictTo calls the method "BrowserCaptureMediaStreamTrack.restrictTo".
//
// The returned bool will be false if there is no such method.
func (this BrowserCaptureMediaStreamTrack) RestrictTo(RestrictionTarget RestrictionTarget) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallBrowserCaptureMediaStreamTrackRestrictTo(
		this.Ref(), js.Pointer(&_ok),
		RestrictionTarget.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// RestrictToFunc returns the method "BrowserCaptureMediaStreamTrack.restrictTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BrowserCaptureMediaStreamTrack) RestrictToFunc() (fn js.Func[func(RestrictionTarget RestrictionTarget) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BrowserCaptureMediaStreamTrackRestrictToFunc(
			this.Ref(),
		),
	)
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

func NewByteLengthQueuingStrategy(init QueuingStrategyInit) ByteLengthQueuingStrategy {
	return ByteLengthQueuingStrategy{}.FromRef(
		bindings.NewByteLengthQueuingStrategyByByteLengthQueuingStrategy(
			js.Pointer(&init)),
	)
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
// The returned bool will be false if there is no such property.
func (this ByteLengthQueuingStrategy) HighWaterMark() (float64, bool) {
	var _ok bool
	_ret := bindings.GetByteLengthQueuingStrategyHighWaterMark(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Size returns the value of property "ByteLengthQueuingStrategy.size".
//
// The returned bool will be false if there is no such property.
func (this ByteLengthQueuingStrategy) Size() (js.Func[func(arguments ...js.Any) js.Any], bool) {
	var _ok bool
	_ret := bindings.GetByteLengthQueuingStrategySize(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Func[func(arguments ...js.Any) js.Any]{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this CSPViolationReportBody) DocumentURL() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSPViolationReportBodyDocumentURL(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Referrer returns the value of property "CSPViolationReportBody.referrer".
//
// The returned bool will be false if there is no such property.
func (this CSPViolationReportBody) Referrer() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSPViolationReportBodyReferrer(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// BlockedURL returns the value of property "CSPViolationReportBody.blockedURL".
//
// The returned bool will be false if there is no such property.
func (this CSPViolationReportBody) BlockedURL() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSPViolationReportBodyBlockedURL(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// EffectiveDirective returns the value of property "CSPViolationReportBody.effectiveDirective".
//
// The returned bool will be false if there is no such property.
func (this CSPViolationReportBody) EffectiveDirective() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSPViolationReportBodyEffectiveDirective(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// OriginalPolicy returns the value of property "CSPViolationReportBody.originalPolicy".
//
// The returned bool will be false if there is no such property.
func (this CSPViolationReportBody) OriginalPolicy() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSPViolationReportBodyOriginalPolicy(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SourceFile returns the value of property "CSPViolationReportBody.sourceFile".
//
// The returned bool will be false if there is no such property.
func (this CSPViolationReportBody) SourceFile() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSPViolationReportBodySourceFile(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Sample returns the value of property "CSPViolationReportBody.sample".
//
// The returned bool will be false if there is no such property.
func (this CSPViolationReportBody) Sample() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSPViolationReportBodySample(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Disposition returns the value of property "CSPViolationReportBody.disposition".
//
// The returned bool will be false if there is no such property.
func (this CSPViolationReportBody) Disposition() (SecurityPolicyViolationEventDisposition, bool) {
	var _ok bool
	_ret := bindings.GetCSPViolationReportBodyDisposition(
		this.Ref(), js.Pointer(&_ok),
	)
	return SecurityPolicyViolationEventDisposition(_ret), _ok
}

// StatusCode returns the value of property "CSPViolationReportBody.statusCode".
//
// The returned bool will be false if there is no such property.
func (this CSPViolationReportBody) StatusCode() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetCSPViolationReportBodyStatusCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// LineNumber returns the value of property "CSPViolationReportBody.lineNumber".
//
// The returned bool will be false if there is no such property.
func (this CSPViolationReportBody) LineNumber() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetCSPViolationReportBodyLineNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ColumnNumber returns the value of property "CSPViolationReportBody.columnNumber".
//
// The returned bool will be false if there is no such property.
func (this CSPViolationReportBody) ColumnNumber() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetCSPViolationReportBodyColumnNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ToJSON calls the method "CSPViolationReportBody.toJSON".
//
// The returned bool will be false if there is no such method.
func (this CSPViolationReportBody) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallCSPViolationReportBodyToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "CSPViolationReportBody.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSPViolationReportBody) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.CSPViolationReportBodyToJSONFunc(
			this.Ref(),
		),
	)
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

func NewCSSParserDeclaration(name js.String, body js.Array[CSSParserRule]) CSSParserDeclaration {
	return CSSParserDeclaration{}.FromRef(
		bindings.NewCSSParserDeclarationByCSSParserDeclaration(
			name.Ref(),
			body.Ref()),
	)
}

func NewCSSParserDeclarationByCSSParserDeclaration1(name js.String) CSSParserDeclaration {
	return CSSParserDeclaration{}.FromRef(
		bindings.NewCSSParserDeclarationByCSSParserDeclaration1(
			name.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSParserDeclaration) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSParserDeclarationName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Body returns the value of property "CSSParserDeclaration.body".
//
// The returned bool will be false if there is no such property.
func (this CSSParserDeclaration) Body() (js.FrozenArray[CSSParserValue], bool) {
	var _ok bool
	_ret := bindings.GetCSSParserDeclarationBody(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[CSSParserValue]{}.FromRef(_ret), _ok
}

// ToString calls the method "CSSParserDeclaration.toString".
//
// The returned bool will be false if there is no such method.
func (this CSSParserDeclaration) ToString() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallCSSParserDeclarationToString(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToStringFunc returns the method "CSSParserDeclaration.toString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSParserDeclaration) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CSSParserDeclarationToStringFunc(
			this.Ref(),
		),
	)
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

// AddModule calls the method "Worklet.addModule".
//
// The returned bool will be false if there is no such method.
func (this Worklet) AddModule(moduleURL js.String, options WorkletOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWorkletAddModule(
		this.Ref(), js.Pointer(&_ok),
		moduleURL.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// AddModuleFunc returns the method "Worklet.addModule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Worklet) AddModuleFunc() (fn js.Func[func(moduleURL js.String, options WorkletOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WorkletAddModuleFunc(
			this.Ref(),
		),
	)
}

// AddModule1 calls the method "Worklet.addModule".
//
// The returned bool will be false if there is no such method.
func (this Worklet) AddModule1(moduleURL js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWorkletAddModule1(
		this.Ref(), js.Pointer(&_ok),
		moduleURL.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// AddModule1Func returns the method "Worklet.addModule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Worklet) AddModule1Func() (fn js.Func[func(moduleURL js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WorkletAddModule1Func(
			this.Ref(),
		),
	)
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
func (CSS) ElementSources() (js.Any, bool) {
	var ok bool
	_ret := bindings.GetCSSElementSources(
		js.Pointer(&ok),
	)
	return js.Any{}.FromRef(_ret), ok
}

// AnimationWorklet returns the value of property "animationWorklet".
//
// The returned bool will be false if there is no such property.
func (CSS) AnimationWorklet() (Worklet, bool) {
	var ok bool
	_ret := bindings.GetCSSAnimationWorklet(
		js.Pointer(&ok),
	)
	return Worklet{}.FromRef(_ret), ok
}

// PaintWorklet returns the value of property "paintWorklet".
//
// The returned bool will be false if there is no such property.
func (CSS) PaintWorklet() (Worklet, bool) {
	var ok bool
	_ret := bindings.GetCSSPaintWorklet(
		js.Pointer(&ok),
	)
	return Worklet{}.FromRef(_ret), ok
}

// LayoutWorklet returns the value of property "layoutWorklet".
//
// The returned bool will be false if there is no such property.
func (CSS) LayoutWorklet() (Worklet, bool) {
	var ok bool
	_ret := bindings.GetCSSLayoutWorklet(
		js.Pointer(&ok),
	)
	return Worklet{}.FromRef(_ret), ok
}

// Highlights returns the value of property "highlights".
//
// The returned bool will be false if there is no such property.
func (CSS) Highlights() (HighlightRegistry, bool) {
	var ok bool
	_ret := bindings.GetCSSHighlights(
		js.Pointer(&ok),
	)
	return HighlightRegistry{}.FromRef(_ret), ok
}

// Escape calls the function "CSS.escape".
//
// The returned bool will be false if there is no such method.
func (CSS) Escape(ident js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallCSSEscape(
		js.Pointer(&_ok),
		ident.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// EscapeFunc returns the function "CSS.escape".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) EscapeFunc() (fn js.Func[func(ident js.String) js.String]) {
	return fn.FromRef(
		bindings.CSSEscapeFunc(),
	)
}

// RegisterProperty calls the function "CSS.registerProperty".
//
// The returned bool will be false if there is no such method.
func (CSS) RegisterProperty(definition PropertyDefinition) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCSSRegisterProperty(
		js.Pointer(&_ok),
		js.Pointer(&definition),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RegisterPropertyFunc returns the function "CSS.registerProperty".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) RegisterPropertyFunc() (fn js.Func[func(definition PropertyDefinition)]) {
	return fn.FromRef(
		bindings.CSSRegisterPropertyFunc(),
	)
}

// Supports calls the function "CSS.supports".
//
// The returned bool will be false if there is no such method.
func (CSS) Supports(property js.String, value js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallCSSSupports(
		js.Pointer(&_ok),
		property.Ref(),
		value.Ref(),
	)

	return _ret == js.True, _ok
}

// SupportsFunc returns the function "CSS.supports".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) SupportsFunc() (fn js.Func[func(property js.String, value js.String) bool]) {
	return fn.FromRef(
		bindings.CSSSupportsFunc(),
	)
}

// Supports1 calls the function "CSS.supports".
//
// The returned bool will be false if there is no such method.
func (CSS) Supports1(conditionText js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallCSSSupports1(
		js.Pointer(&_ok),
		conditionText.Ref(),
	)

	return _ret == js.True, _ok
}

// Supports1Func returns the function "CSS.supports".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) Supports1Func() (fn js.Func[func(conditionText js.String) bool]) {
	return fn.FromRef(
		bindings.CSSSupports1Func(),
	)
}

// Number calls the function "CSS.number".
//
// The returned bool will be false if there is no such method.
func (CSS) Number(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSNumber(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// NumberFunc returns the function "CSS.number".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) NumberFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSNumberFunc(),
	)
}

// Percent calls the function "CSS.percent".
//
// The returned bool will be false if there is no such method.
func (CSS) Percent(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSPercent(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// PercentFunc returns the function "CSS.percent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) PercentFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSPercentFunc(),
	)
}

// Cap calls the function "CSS.cap".
//
// The returned bool will be false if there is no such method.
func (CSS) Cap(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSCap(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// CapFunc returns the function "CSS.cap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) CapFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCapFunc(),
	)
}

// Ch calls the function "CSS.ch".
//
// The returned bool will be false if there is no such method.
func (CSS) Ch(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSCh(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// ChFunc returns the function "CSS.ch".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ChFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSChFunc(),
	)
}

// Em calls the function "CSS.em".
//
// The returned bool will be false if there is no such method.
func (CSS) Em(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSEm(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// EmFunc returns the function "CSS.em".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) EmFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSEmFunc(),
	)
}

// Ex calls the function "CSS.ex".
//
// The returned bool will be false if there is no such method.
func (CSS) Ex(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSEx(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// ExFunc returns the function "CSS.ex".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ExFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSExFunc(),
	)
}

// Ic calls the function "CSS.ic".
//
// The returned bool will be false if there is no such method.
func (CSS) Ic(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSIc(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// IcFunc returns the function "CSS.ic".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) IcFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSIcFunc(),
	)
}

// Lh calls the function "CSS.lh".
//
// The returned bool will be false if there is no such method.
func (CSS) Lh(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSLh(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// LhFunc returns the function "CSS.lh".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) LhFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSLhFunc(),
	)
}

// Rcap calls the function "CSS.rcap".
//
// The returned bool will be false if there is no such method.
func (CSS) Rcap(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSRcap(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// RcapFunc returns the function "CSS.rcap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) RcapFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSRcapFunc(),
	)
}

// Rch calls the function "CSS.rch".
//
// The returned bool will be false if there is no such method.
func (CSS) Rch(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSRch(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// RchFunc returns the function "CSS.rch".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) RchFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSRchFunc(),
	)
}

// Rem calls the function "CSS.rem".
//
// The returned bool will be false if there is no such method.
func (CSS) Rem(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSRem(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// RemFunc returns the function "CSS.rem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) RemFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSRemFunc(),
	)
}

// Rex calls the function "CSS.rex".
//
// The returned bool will be false if there is no such method.
func (CSS) Rex(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSRex(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// RexFunc returns the function "CSS.rex".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) RexFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSRexFunc(),
	)
}

// Ric calls the function "CSS.ric".
//
// The returned bool will be false if there is no such method.
func (CSS) Ric(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSRic(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// RicFunc returns the function "CSS.ric".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) RicFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSRicFunc(),
	)
}

// Rlh calls the function "CSS.rlh".
//
// The returned bool will be false if there is no such method.
func (CSS) Rlh(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSRlh(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// RlhFunc returns the function "CSS.rlh".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) RlhFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSRlhFunc(),
	)
}

// Vw calls the function "CSS.vw".
//
// The returned bool will be false if there is no such method.
func (CSS) Vw(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSVw(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// VwFunc returns the function "CSS.vw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) VwFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSVwFunc(),
	)
}

// Vh calls the function "CSS.vh".
//
// The returned bool will be false if there is no such method.
func (CSS) Vh(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSVh(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// VhFunc returns the function "CSS.vh".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) VhFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSVhFunc(),
	)
}

// Vi calls the function "CSS.vi".
//
// The returned bool will be false if there is no such method.
func (CSS) Vi(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSVi(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// ViFunc returns the function "CSS.vi".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ViFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSViFunc(),
	)
}

// Vb calls the function "CSS.vb".
//
// The returned bool will be false if there is no such method.
func (CSS) Vb(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSVb(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// VbFunc returns the function "CSS.vb".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) VbFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSVbFunc(),
	)
}

// Vmin calls the function "CSS.vmin".
//
// The returned bool will be false if there is no such method.
func (CSS) Vmin(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSVmin(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// VminFunc returns the function "CSS.vmin".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) VminFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSVminFunc(),
	)
}

// Vmax calls the function "CSS.vmax".
//
// The returned bool will be false if there is no such method.
func (CSS) Vmax(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSVmax(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// VmaxFunc returns the function "CSS.vmax".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) VmaxFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSVmaxFunc(),
	)
}

// Svw calls the function "CSS.svw".
//
// The returned bool will be false if there is no such method.
func (CSS) Svw(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSSvw(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// SvwFunc returns the function "CSS.svw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) SvwFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSSvwFunc(),
	)
}

// Svh calls the function "CSS.svh".
//
// The returned bool will be false if there is no such method.
func (CSS) Svh(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSSvh(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// SvhFunc returns the function "CSS.svh".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) SvhFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSSvhFunc(),
	)
}

// Svi calls the function "CSS.svi".
//
// The returned bool will be false if there is no such method.
func (CSS) Svi(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSSvi(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// SviFunc returns the function "CSS.svi".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) SviFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSSviFunc(),
	)
}

// Svb calls the function "CSS.svb".
//
// The returned bool will be false if there is no such method.
func (CSS) Svb(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSSvb(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// SvbFunc returns the function "CSS.svb".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) SvbFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSSvbFunc(),
	)
}

// Svmin calls the function "CSS.svmin".
//
// The returned bool will be false if there is no such method.
func (CSS) Svmin(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSSvmin(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// SvminFunc returns the function "CSS.svmin".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) SvminFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSSvminFunc(),
	)
}

// Svmax calls the function "CSS.svmax".
//
// The returned bool will be false if there is no such method.
func (CSS) Svmax(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSSvmax(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// SvmaxFunc returns the function "CSS.svmax".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) SvmaxFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSSvmaxFunc(),
	)
}

// Lvw calls the function "CSS.lvw".
//
// The returned bool will be false if there is no such method.
func (CSS) Lvw(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSLvw(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// LvwFunc returns the function "CSS.lvw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) LvwFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSLvwFunc(),
	)
}

// Lvh calls the function "CSS.lvh".
//
// The returned bool will be false if there is no such method.
func (CSS) Lvh(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSLvh(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// LvhFunc returns the function "CSS.lvh".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) LvhFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSLvhFunc(),
	)
}

// Lvi calls the function "CSS.lvi".
//
// The returned bool will be false if there is no such method.
func (CSS) Lvi(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSLvi(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// LviFunc returns the function "CSS.lvi".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) LviFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSLviFunc(),
	)
}

// Lvb calls the function "CSS.lvb".
//
// The returned bool will be false if there is no such method.
func (CSS) Lvb(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSLvb(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// LvbFunc returns the function "CSS.lvb".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) LvbFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSLvbFunc(),
	)
}

// Lvmin calls the function "CSS.lvmin".
//
// The returned bool will be false if there is no such method.
func (CSS) Lvmin(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSLvmin(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// LvminFunc returns the function "CSS.lvmin".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) LvminFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSLvminFunc(),
	)
}

// Lvmax calls the function "CSS.lvmax".
//
// The returned bool will be false if there is no such method.
func (CSS) Lvmax(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSLvmax(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// LvmaxFunc returns the function "CSS.lvmax".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) LvmaxFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSLvmaxFunc(),
	)
}

// Dvw calls the function "CSS.dvw".
//
// The returned bool will be false if there is no such method.
func (CSS) Dvw(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSDvw(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// DvwFunc returns the function "CSS.dvw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) DvwFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDvwFunc(),
	)
}

// Dvh calls the function "CSS.dvh".
//
// The returned bool will be false if there is no such method.
func (CSS) Dvh(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSDvh(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// DvhFunc returns the function "CSS.dvh".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) DvhFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDvhFunc(),
	)
}

// Dvi calls the function "CSS.dvi".
//
// The returned bool will be false if there is no such method.
func (CSS) Dvi(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSDvi(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// DviFunc returns the function "CSS.dvi".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) DviFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDviFunc(),
	)
}

// Dvb calls the function "CSS.dvb".
//
// The returned bool will be false if there is no such method.
func (CSS) Dvb(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSDvb(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// DvbFunc returns the function "CSS.dvb".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) DvbFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDvbFunc(),
	)
}

// Dvmin calls the function "CSS.dvmin".
//
// The returned bool will be false if there is no such method.
func (CSS) Dvmin(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSDvmin(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// DvminFunc returns the function "CSS.dvmin".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) DvminFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDvminFunc(),
	)
}

// Dvmax calls the function "CSS.dvmax".
//
// The returned bool will be false if there is no such method.
func (CSS) Dvmax(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSDvmax(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// DvmaxFunc returns the function "CSS.dvmax".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) DvmaxFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDvmaxFunc(),
	)
}

// Cqw calls the function "CSS.cqw".
//
// The returned bool will be false if there is no such method.
func (CSS) Cqw(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSCqw(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// CqwFunc returns the function "CSS.cqw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) CqwFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCqwFunc(),
	)
}

// Cqh calls the function "CSS.cqh".
//
// The returned bool will be false if there is no such method.
func (CSS) Cqh(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSCqh(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// CqhFunc returns the function "CSS.cqh".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) CqhFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCqhFunc(),
	)
}

// Cqi calls the function "CSS.cqi".
//
// The returned bool will be false if there is no such method.
func (CSS) Cqi(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSCqi(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// CqiFunc returns the function "CSS.cqi".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) CqiFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCqiFunc(),
	)
}

// Cqb calls the function "CSS.cqb".
//
// The returned bool will be false if there is no such method.
func (CSS) Cqb(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSCqb(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// CqbFunc returns the function "CSS.cqb".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) CqbFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCqbFunc(),
	)
}

// Cqmin calls the function "CSS.cqmin".
//
// The returned bool will be false if there is no such method.
func (CSS) Cqmin(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSCqmin(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// CqminFunc returns the function "CSS.cqmin".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) CqminFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCqminFunc(),
	)
}

// Cqmax calls the function "CSS.cqmax".
//
// The returned bool will be false if there is no such method.
func (CSS) Cqmax(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSCqmax(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// CqmaxFunc returns the function "CSS.cqmax".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) CqmaxFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCqmaxFunc(),
	)
}

// Cm calls the function "CSS.cm".
//
// The returned bool will be false if there is no such method.
func (CSS) Cm(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSCm(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// CmFunc returns the function "CSS.cm".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) CmFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSCmFunc(),
	)
}

// Mm calls the function "CSS.mm".
//
// The returned bool will be false if there is no such method.
func (CSS) Mm(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSMm(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// MmFunc returns the function "CSS.mm".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) MmFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSMmFunc(),
	)
}

// Q calls the function "CSS.Q".
//
// The returned bool will be false if there is no such method.
func (CSS) Q(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSQ(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// QFunc returns the function "CSS.Q".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) QFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSQFunc(),
	)
}

// In calls the function "CSS.in".
//
// The returned bool will be false if there is no such method.
func (CSS) In(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSIn(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// InFunc returns the function "CSS.in".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) InFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSInFunc(),
	)
}

// Pt calls the function "CSS.pt".
//
// The returned bool will be false if there is no such method.
func (CSS) Pt(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSPt(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// PtFunc returns the function "CSS.pt".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) PtFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSPtFunc(),
	)
}

// Pc calls the function "CSS.pc".
//
// The returned bool will be false if there is no such method.
func (CSS) Pc(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSPc(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// PcFunc returns the function "CSS.pc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) PcFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSPcFunc(),
	)
}

// Px calls the function "CSS.px".
//
// The returned bool will be false if there is no such method.
func (CSS) Px(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSPx(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// PxFunc returns the function "CSS.px".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) PxFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSPxFunc(),
	)
}

// Deg calls the function "CSS.deg".
//
// The returned bool will be false if there is no such method.
func (CSS) Deg(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSDeg(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// DegFunc returns the function "CSS.deg".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) DegFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDegFunc(),
	)
}

// Grad calls the function "CSS.grad".
//
// The returned bool will be false if there is no such method.
func (CSS) Grad(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSGrad(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// GradFunc returns the function "CSS.grad".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) GradFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSGradFunc(),
	)
}

// Rad calls the function "CSS.rad".
//
// The returned bool will be false if there is no such method.
func (CSS) Rad(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSRad(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// RadFunc returns the function "CSS.rad".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) RadFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSRadFunc(),
	)
}

// Turn calls the function "CSS.turn".
//
// The returned bool will be false if there is no such method.
func (CSS) Turn(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSTurn(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// TurnFunc returns the function "CSS.turn".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) TurnFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSTurnFunc(),
	)
}

// S calls the function "CSS.s".
//
// The returned bool will be false if there is no such method.
func (CSS) S(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSS(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// SFunc returns the function "CSS.s".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) SFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSSFunc(),
	)
}

// Ms calls the function "CSS.ms".
//
// The returned bool will be false if there is no such method.
func (CSS) Ms(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSMs(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// MsFunc returns the function "CSS.ms".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) MsFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSMsFunc(),
	)
}

// Hz calls the function "CSS.Hz".
//
// The returned bool will be false if there is no such method.
func (CSS) Hz(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSHz(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// HzFunc returns the function "CSS.Hz".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) HzFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSHzFunc(),
	)
}

// KHz calls the function "CSS.kHz".
//
// The returned bool will be false if there is no such method.
func (CSS) KHz(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSKHz(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// KHzFunc returns the function "CSS.kHz".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) KHzFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSKHzFunc(),
	)
}

// Dpi calls the function "CSS.dpi".
//
// The returned bool will be false if there is no such method.
func (CSS) Dpi(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSDpi(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// DpiFunc returns the function "CSS.dpi".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) DpiFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDpiFunc(),
	)
}

// Dpcm calls the function "CSS.dpcm".
//
// The returned bool will be false if there is no such method.
func (CSS) Dpcm(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSDpcm(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// DpcmFunc returns the function "CSS.dpcm".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) DpcmFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDpcmFunc(),
	)
}

// Dppx calls the function "CSS.dppx".
//
// The returned bool will be false if there is no such method.
func (CSS) Dppx(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSDppx(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// DppxFunc returns the function "CSS.dppx".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) DppxFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSDppxFunc(),
	)
}

// Fr calls the function "CSS.fr".
//
// The returned bool will be false if there is no such method.
func (CSS) Fr(value float64) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSFr(
		js.Pointer(&_ok),
		float64(value),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// FrFunc returns the function "CSS.fr".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) FrFunc() (fn js.Func[func(value float64) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSFrFunc(),
	)
}

// ParseStylesheet calls the function "CSS.parseStylesheet".
//
// The returned bool will be false if there is no such method.
func (CSS) ParseStylesheet(css CSSStringSource, options CSSParserOptions) (js.Promise[js.Array[CSSParserRule]], bool) {
	var _ok bool
	_ret := bindings.CallCSSParseStylesheet(
		js.Pointer(&_ok),
		css.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Array[CSSParserRule]]{}.FromRef(_ret), _ok
}

// ParseStylesheetFunc returns the function "CSS.parseStylesheet".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ParseStylesheetFunc() (fn js.Func[func(css CSSStringSource, options CSSParserOptions) js.Promise[js.Array[CSSParserRule]]]) {
	return fn.FromRef(
		bindings.CSSParseStylesheetFunc(),
	)
}

// ParseStylesheet1 calls the function "CSS.parseStylesheet".
//
// The returned bool will be false if there is no such method.
func (CSS) ParseStylesheet1(css CSSStringSource) (js.Promise[js.Array[CSSParserRule]], bool) {
	var _ok bool
	_ret := bindings.CallCSSParseStylesheet1(
		js.Pointer(&_ok),
		css.Ref(),
	)

	return js.Promise[js.Array[CSSParserRule]]{}.FromRef(_ret), _ok
}

// ParseStylesheet1Func returns the function "CSS.parseStylesheet".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ParseStylesheet1Func() (fn js.Func[func(css CSSStringSource) js.Promise[js.Array[CSSParserRule]]]) {
	return fn.FromRef(
		bindings.CSSParseStylesheet1Func(),
	)
}

// ParseRuleList calls the function "CSS.parseRuleList".
//
// The returned bool will be false if there is no such method.
func (CSS) ParseRuleList(css CSSStringSource, options CSSParserOptions) (js.Promise[js.Array[CSSParserRule]], bool) {
	var _ok bool
	_ret := bindings.CallCSSParseRuleList(
		js.Pointer(&_ok),
		css.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Array[CSSParserRule]]{}.FromRef(_ret), _ok
}

// ParseRuleListFunc returns the function "CSS.parseRuleList".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ParseRuleListFunc() (fn js.Func[func(css CSSStringSource, options CSSParserOptions) js.Promise[js.Array[CSSParserRule]]]) {
	return fn.FromRef(
		bindings.CSSParseRuleListFunc(),
	)
}

// ParseRuleList1 calls the function "CSS.parseRuleList".
//
// The returned bool will be false if there is no such method.
func (CSS) ParseRuleList1(css CSSStringSource) (js.Promise[js.Array[CSSParserRule]], bool) {
	var _ok bool
	_ret := bindings.CallCSSParseRuleList1(
		js.Pointer(&_ok),
		css.Ref(),
	)

	return js.Promise[js.Array[CSSParserRule]]{}.FromRef(_ret), _ok
}

// ParseRuleList1Func returns the function "CSS.parseRuleList".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ParseRuleList1Func() (fn js.Func[func(css CSSStringSource) js.Promise[js.Array[CSSParserRule]]]) {
	return fn.FromRef(
		bindings.CSSParseRuleList1Func(),
	)
}

// ParseRule calls the function "CSS.parseRule".
//
// The returned bool will be false if there is no such method.
func (CSS) ParseRule(css CSSStringSource, options CSSParserOptions) (js.Promise[CSSParserRule], bool) {
	var _ok bool
	_ret := bindings.CallCSSParseRule(
		js.Pointer(&_ok),
		css.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[CSSParserRule]{}.FromRef(_ret), _ok
}

// ParseRuleFunc returns the function "CSS.parseRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ParseRuleFunc() (fn js.Func[func(css CSSStringSource, options CSSParserOptions) js.Promise[CSSParserRule]]) {
	return fn.FromRef(
		bindings.CSSParseRuleFunc(),
	)
}

// ParseRule1 calls the function "CSS.parseRule".
//
// The returned bool will be false if there is no such method.
func (CSS) ParseRule1(css CSSStringSource) (js.Promise[CSSParserRule], bool) {
	var _ok bool
	_ret := bindings.CallCSSParseRule1(
		js.Pointer(&_ok),
		css.Ref(),
	)

	return js.Promise[CSSParserRule]{}.FromRef(_ret), _ok
}

// ParseRule1Func returns the function "CSS.parseRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ParseRule1Func() (fn js.Func[func(css CSSStringSource) js.Promise[CSSParserRule]]) {
	return fn.FromRef(
		bindings.CSSParseRule1Func(),
	)
}

// ParseDeclarationList calls the function "CSS.parseDeclarationList".
//
// The returned bool will be false if there is no such method.
func (CSS) ParseDeclarationList(css CSSStringSource, options CSSParserOptions) (js.Promise[js.Array[CSSParserRule]], bool) {
	var _ok bool
	_ret := bindings.CallCSSParseDeclarationList(
		js.Pointer(&_ok),
		css.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Array[CSSParserRule]]{}.FromRef(_ret), _ok
}

// ParseDeclarationListFunc returns the function "CSS.parseDeclarationList".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ParseDeclarationListFunc() (fn js.Func[func(css CSSStringSource, options CSSParserOptions) js.Promise[js.Array[CSSParserRule]]]) {
	return fn.FromRef(
		bindings.CSSParseDeclarationListFunc(),
	)
}

// ParseDeclarationList1 calls the function "CSS.parseDeclarationList".
//
// The returned bool will be false if there is no such method.
func (CSS) ParseDeclarationList1(css CSSStringSource) (js.Promise[js.Array[CSSParserRule]], bool) {
	var _ok bool
	_ret := bindings.CallCSSParseDeclarationList1(
		js.Pointer(&_ok),
		css.Ref(),
	)

	return js.Promise[js.Array[CSSParserRule]]{}.FromRef(_ret), _ok
}

// ParseDeclarationList1Func returns the function "CSS.parseDeclarationList".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ParseDeclarationList1Func() (fn js.Func[func(css CSSStringSource) js.Promise[js.Array[CSSParserRule]]]) {
	return fn.FromRef(
		bindings.CSSParseDeclarationList1Func(),
	)
}

// ParseDeclaration calls the function "CSS.parseDeclaration".
//
// The returned bool will be false if there is no such method.
func (CSS) ParseDeclaration(css js.String, options CSSParserOptions) (CSSParserDeclaration, bool) {
	var _ok bool
	_ret := bindings.CallCSSParseDeclaration(
		js.Pointer(&_ok),
		css.Ref(),
		js.Pointer(&options),
	)

	return CSSParserDeclaration{}.FromRef(_ret), _ok
}

// ParseDeclarationFunc returns the function "CSS.parseDeclaration".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ParseDeclarationFunc() (fn js.Func[func(css js.String, options CSSParserOptions) CSSParserDeclaration]) {
	return fn.FromRef(
		bindings.CSSParseDeclarationFunc(),
	)
}

// ParseDeclaration1 calls the function "CSS.parseDeclaration".
//
// The returned bool will be false if there is no such method.
func (CSS) ParseDeclaration1(css js.String) (CSSParserDeclaration, bool) {
	var _ok bool
	_ret := bindings.CallCSSParseDeclaration1(
		js.Pointer(&_ok),
		css.Ref(),
	)

	return CSSParserDeclaration{}.FromRef(_ret), _ok
}

// ParseDeclaration1Func returns the function "CSS.parseDeclaration".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ParseDeclaration1Func() (fn js.Func[func(css js.String) CSSParserDeclaration]) {
	return fn.FromRef(
		bindings.CSSParseDeclaration1Func(),
	)
}

// ParseValue calls the function "CSS.parseValue".
//
// The returned bool will be false if there is no such method.
func (CSS) ParseValue(css js.String) (CSSToken, bool) {
	var _ok bool
	_ret := bindings.CallCSSParseValue(
		js.Pointer(&_ok),
		css.Ref(),
	)

	return CSSToken{}.FromRef(_ret), _ok
}

// ParseValueFunc returns the function "CSS.parseValue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ParseValueFunc() (fn js.Func[func(css js.String) CSSToken]) {
	return fn.FromRef(
		bindings.CSSParseValueFunc(),
	)
}

// ParseValueList calls the function "CSS.parseValueList".
//
// The returned bool will be false if there is no such method.
func (CSS) ParseValueList(css js.String) (js.Array[CSSToken], bool) {
	var _ok bool
	_ret := bindings.CallCSSParseValueList(
		js.Pointer(&_ok),
		css.Ref(),
	)

	return js.Array[CSSToken]{}.FromRef(_ret), _ok
}

// ParseValueListFunc returns the function "CSS.parseValueList".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ParseValueListFunc() (fn js.Func[func(css js.String) js.Array[CSSToken]]) {
	return fn.FromRef(
		bindings.CSSParseValueListFunc(),
	)
}

// ParseCommaValueList calls the function "CSS.parseCommaValueList".
//
// The returned bool will be false if there is no such method.
func (CSS) ParseCommaValueList(css js.String) (js.Array[js.Array[CSSToken]], bool) {
	var _ok bool
	_ret := bindings.CallCSSParseCommaValueList(
		js.Pointer(&_ok),
		css.Ref(),
	)

	return js.Array[js.Array[CSSToken]]{}.FromRef(_ret), _ok
}

// ParseCommaValueListFunc returns the function "CSS.parseCommaValueList".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSS) ParseCommaValueListFunc() (fn js.Func[func(css js.String) js.Array[js.Array[CSSToken]]]) {
	return fn.FromRef(
		bindings.CSSParseCommaValueListFunc(),
	)
}

func NewCSSAnimation(effect AnimationEffect, timeline AnimationTimeline) CSSAnimation {
	return CSSAnimation{}.FromRef(
		bindings.NewCSSAnimationByCSSAnimation(
			effect.Ref(),
			timeline.Ref()),
	)
}

func NewCSSAnimationByCSSAnimation1(effect AnimationEffect) CSSAnimation {
	return CSSAnimation{}.FromRef(
		bindings.NewCSSAnimationByCSSAnimation1(
			effect.Ref()),
	)
}

func NewCSSAnimationByCSSAnimation2() CSSAnimation {
	return CSSAnimation{}.FromRef(
		bindings.NewCSSAnimationByCSSAnimation2(),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSAnimation) AnimationName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSAnimationAnimationName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
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
