// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package bluetooth

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/bluetooth/bindings"
)

type AdapterState struct {
	// Address is "AdapterState.address"
	//
	// Optional
	Address js.String
	// Name is "AdapterState.name"
	//
	// Optional
	Name js.String
	// Powered is "AdapterState.powered"
	//
	// Optional
	//
	// NOTE: FFI_USE_Powered MUST be set to true to make this field effective.
	Powered bool
	// Available is "AdapterState.available"
	//
	// Optional
	//
	// NOTE: FFI_USE_Available MUST be set to true to make this field effective.
	Available bool
	// Discovering is "AdapterState.discovering"
	//
	// Optional
	//
	// NOTE: FFI_USE_Discovering MUST be set to true to make this field effective.
	Discovering bool

	FFI_USE_Powered     bool // for Powered.
	FFI_USE_Available   bool // for Available.
	FFI_USE_Discovering bool // for Discovering.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AdapterState with all fields set.
func (p AdapterState) FromRef(ref js.Ref) AdapterState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AdapterState in the application heap.
func (p AdapterState) New() js.Ref {
	return bindings.AdapterStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AdapterState) UpdateFrom(ref js.Ref) {
	bindings.AdapterStateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AdapterState) Update(ref js.Ref) {
	bindings.AdapterStateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AdapterState) FreeMembers(recursive bool) {
	js.Free(
		p.Address.Ref(),
		p.Name.Ref(),
	)
	p.Address = p.Address.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type AdapterStateCallbackFunc func(this js.Ref, adapterInfo *AdapterState) js.Ref

func (fn AdapterStateCallbackFunc) Register() js.Func[func(adapterInfo *AdapterState)] {
	return js.RegisterCallback[func(adapterInfo *AdapterState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AdapterStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AdapterState
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

type AdapterStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, adapterInfo *AdapterState) js.Ref
	Arg T
}

func (cb *AdapterStateCallback[T]) Register() js.Func[func(adapterInfo *AdapterState)] {
	return js.RegisterCallback[func(adapterInfo *AdapterState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AdapterStateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AdapterState
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

type FilterType uint32

const (
	_ FilterType = iota

	FilterType_ALL
	FilterType_KNOWN
)

func (FilterType) FromRef(str js.Ref) FilterType {
	return FilterType(bindings.ConstOfFilterType(str))
}

func (x FilterType) String() (string, bool) {
	switch x {
	case FilterType_ALL:
		return "all", true
	case FilterType_KNOWN:
		return "known", true
	default:
		return "", false
	}
}

type BluetoothFilter struct {
	// FilterType is "BluetoothFilter.filterType"
	//
	// Optional
	FilterType FilterType
	// Limit is "BluetoothFilter.limit"
	//
	// Optional
	//
	// NOTE: FFI_USE_Limit MUST be set to true to make this field effective.
	Limit int32

	FFI_USE_Limit bool // for Limit.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BluetoothFilter with all fields set.
func (p BluetoothFilter) FromRef(ref js.Ref) BluetoothFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BluetoothFilter in the application heap.
func (p BluetoothFilter) New() js.Ref {
	return bindings.BluetoothFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BluetoothFilter) UpdateFrom(ref js.Ref) {
	bindings.BluetoothFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BluetoothFilter) Update(ref js.Ref) {
	bindings.BluetoothFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BluetoothFilter) FreeMembers(recursive bool) {
}

type VendorIdSource uint32

const (
	_ VendorIdSource = iota

	VendorIdSource_BLUETOOTH
	VendorIdSource_USB
)

func (VendorIdSource) FromRef(str js.Ref) VendorIdSource {
	return VendorIdSource(bindings.ConstOfVendorIdSource(str))
}

func (x VendorIdSource) String() (string, bool) {
	switch x {
	case VendorIdSource_BLUETOOTH:
		return "bluetooth", true
	case VendorIdSource_USB:
		return "usb", true
	default:
		return "", false
	}
}

type DeviceType uint32

const (
	_ DeviceType = iota

	DeviceType_COMPUTER
	DeviceType_PHONE
	DeviceType_MODEM
	DeviceType_AUDIO
	DeviceType_CAR_AUDIO
	DeviceType_VIDEO
	DeviceType_PERIPHERAL
	DeviceType_JOYSTICK
	DeviceType_GAMEPAD
	DeviceType_KEYBOARD
	DeviceType_MOUSE
	DeviceType_TABLET
	DeviceType_KEYBOARD_MOUSE_COMBO
)

func (DeviceType) FromRef(str js.Ref) DeviceType {
	return DeviceType(bindings.ConstOfDeviceType(str))
}

func (x DeviceType) String() (string, bool) {
	switch x {
	case DeviceType_COMPUTER:
		return "computer", true
	case DeviceType_PHONE:
		return "phone", true
	case DeviceType_MODEM:
		return "modem", true
	case DeviceType_AUDIO:
		return "audio", true
	case DeviceType_CAR_AUDIO:
		return "carAudio", true
	case DeviceType_VIDEO:
		return "video", true
	case DeviceType_PERIPHERAL:
		return "peripheral", true
	case DeviceType_JOYSTICK:
		return "joystick", true
	case DeviceType_GAMEPAD:
		return "gamepad", true
	case DeviceType_KEYBOARD:
		return "keyboard", true
	case DeviceType_MOUSE:
		return "mouse", true
	case DeviceType_TABLET:
		return "tablet", true
	case DeviceType_KEYBOARD_MOUSE_COMBO:
		return "keyboardMouseCombo", true
	default:
		return "", false
	}
}

type Transport uint32

const (
	_ Transport = iota

	Transport_INVALID
	Transport_CLASSIC
	Transport_LE
	Transport_DUAL
)

func (Transport) FromRef(str js.Ref) Transport {
	return Transport(bindings.ConstOfTransport(str))
}

func (x Transport) String() (string, bool) {
	switch x {
	case Transport_INVALID:
		return "invalid", true
	case Transport_CLASSIC:
		return "classic", true
	case Transport_LE:
		return "le", true
	case Transport_DUAL:
		return "dual", true
	default:
		return "", false
	}
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
	// VendorIdSource is "Device.vendorIdSource"
	//
	// Optional
	VendorIdSource VendorIdSource
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
	// DeviceId is "Device.deviceId"
	//
	// Optional
	//
	// NOTE: FFI_USE_DeviceId MUST be set to true to make this field effective.
	DeviceId int32
	// Type is "Device.type"
	//
	// Optional
	Type DeviceType
	// Paired is "Device.paired"
	//
	// Optional
	//
	// NOTE: FFI_USE_Paired MUST be set to true to make this field effective.
	Paired bool
	// Connected is "Device.connected"
	//
	// Optional
	//
	// NOTE: FFI_USE_Connected MUST be set to true to make this field effective.
	Connected bool
	// Connecting is "Device.connecting"
	//
	// Optional
	//
	// NOTE: FFI_USE_Connecting MUST be set to true to make this field effective.
	Connecting bool
	// Connectable is "Device.connectable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Connectable MUST be set to true to make this field effective.
	Connectable bool
	// Uuids is "Device.uuids"
	//
	// Optional
	Uuids js.Array[js.String]
	// InquiryRssi is "Device.inquiryRssi"
	//
	// Optional
	//
	// NOTE: FFI_USE_InquiryRssi MUST be set to true to make this field effective.
	InquiryRssi int32
	// InquiryTxPower is "Device.inquiryTxPower"
	//
	// Optional
	//
	// NOTE: FFI_USE_InquiryTxPower MUST be set to true to make this field effective.
	InquiryTxPower int32
	// Transport is "Device.transport"
	//
	// Optional
	Transport Transport
	// BatteryPercentage is "Device.batteryPercentage"
	//
	// Optional
	//
	// NOTE: FFI_USE_BatteryPercentage MUST be set to true to make this field effective.
	BatteryPercentage int32

	FFI_USE_DeviceClass       bool // for DeviceClass.
	FFI_USE_VendorId          bool // for VendorId.
	FFI_USE_ProductId         bool // for ProductId.
	FFI_USE_DeviceId          bool // for DeviceId.
	FFI_USE_Paired            bool // for Paired.
	FFI_USE_Connected         bool // for Connected.
	FFI_USE_Connecting        bool // for Connecting.
	FFI_USE_Connectable       bool // for Connectable.
	FFI_USE_InquiryRssi       bool // for InquiryRssi.
	FFI_USE_InquiryTxPower    bool // for InquiryTxPower.
	FFI_USE_BatteryPercentage bool // for BatteryPercentage.

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
		p.Uuids.Ref(),
	)
	p.Address = p.Address.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Uuids = p.Uuids.FromRef(js.Undefined)
}

type GetDeviceCallbackFunc func(this js.Ref, deviceInfo *Device) js.Ref

func (fn GetDeviceCallbackFunc) Register() js.Func[func(deviceInfo *Device)] {
	return js.RegisterCallback[func(deviceInfo *Device)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDeviceCallbackFunc) DispatchCallback(
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

type GetDeviceCallback[T any] struct {
	Fn  func(arg T, this js.Ref, deviceInfo *Device) js.Ref
	Arg T
}

func (cb *GetDeviceCallback[T]) Register() js.Func[func(deviceInfo *Device)] {
	return js.RegisterCallback[func(deviceInfo *Device)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDeviceCallback[T]) DispatchCallback(
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

type GetDevicesCallbackFunc func(this js.Ref, deviceInfos js.Array[Device]) js.Ref

func (fn GetDevicesCallbackFunc) Register() js.Func[func(deviceInfos js.Array[Device])] {
	return js.RegisterCallback[func(deviceInfos js.Array[Device])](
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
	Fn  func(arg T, this js.Ref, deviceInfos js.Array[Device]) js.Ref
	Arg T
}

func (cb *GetDevicesCallback[T]) Register() js.Func[func(deviceInfos js.Array[Device])] {
	return js.RegisterCallback[func(deviceInfos js.Array[Device])](
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

type StartDiscoveryCallbackFunc func(this js.Ref) js.Ref

func (fn StartDiscoveryCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn StartDiscoveryCallbackFunc) DispatchCallback(
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

type StartDiscoveryCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *StartDiscoveryCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *StartDiscoveryCallback[T]) DispatchCallback(
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

type StopDiscoveryCallbackFunc func(this js.Ref) js.Ref

func (fn StopDiscoveryCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn StopDiscoveryCallbackFunc) DispatchCallback(
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

type StopDiscoveryCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *StopDiscoveryCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *StopDiscoveryCallback[T]) DispatchCallback(
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

// HasFuncGetAdapterState returns true if the function "WEBEXT.bluetooth.getAdapterState" exists.
func HasFuncGetAdapterState() bool {
	return js.True == bindings.HasFuncGetAdapterState()
}

// FuncGetAdapterState returns the function "WEBEXT.bluetooth.getAdapterState".
func FuncGetAdapterState() (fn js.Func[func() js.Promise[AdapterState]]) {
	bindings.FuncGetAdapterState(
		js.Pointer(&fn),
	)
	return
}

// GetAdapterState calls the function "WEBEXT.bluetooth.getAdapterState" directly.
func GetAdapterState() (ret js.Promise[AdapterState]) {
	bindings.CallGetAdapterState(
		js.Pointer(&ret),
	)

	return
}

// TryGetAdapterState calls the function "WEBEXT.bluetooth.getAdapterState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAdapterState() (ret js.Promise[AdapterState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAdapterState(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDevice returns true if the function "WEBEXT.bluetooth.getDevice" exists.
func HasFuncGetDevice() bool {
	return js.True == bindings.HasFuncGetDevice()
}

// FuncGetDevice returns the function "WEBEXT.bluetooth.getDevice".
func FuncGetDevice() (fn js.Func[func(deviceAddress js.String) js.Promise[Device]]) {
	bindings.FuncGetDevice(
		js.Pointer(&fn),
	)
	return
}

// GetDevice calls the function "WEBEXT.bluetooth.getDevice" directly.
func GetDevice(deviceAddress js.String) (ret js.Promise[Device]) {
	bindings.CallGetDevice(
		js.Pointer(&ret),
		deviceAddress.Ref(),
	)

	return
}

// TryGetDevice calls the function "WEBEXT.bluetooth.getDevice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDevice(deviceAddress js.String) (ret js.Promise[Device], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDevice(
		js.Pointer(&ret), js.Pointer(&exception),
		deviceAddress.Ref(),
	)

	return
}

// HasFuncGetDevices returns true if the function "WEBEXT.bluetooth.getDevices" exists.
func HasFuncGetDevices() bool {
	return js.True == bindings.HasFuncGetDevices()
}

// FuncGetDevices returns the function "WEBEXT.bluetooth.getDevices".
func FuncGetDevices() (fn js.Func[func(filter BluetoothFilter) js.Promise[js.Array[Device]]]) {
	bindings.FuncGetDevices(
		js.Pointer(&fn),
	)
	return
}

// GetDevices calls the function "WEBEXT.bluetooth.getDevices" directly.
func GetDevices(filter BluetoothFilter) (ret js.Promise[js.Array[Device]]) {
	bindings.CallGetDevices(
		js.Pointer(&ret),
		js.Pointer(&filter),
	)

	return
}

// TryGetDevices calls the function "WEBEXT.bluetooth.getDevices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDevices(filter BluetoothFilter) (ret js.Promise[js.Array[Device]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDevices(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&filter),
	)

	return
}

type OnAdapterStateChangedEventCallbackFunc func(this js.Ref, state *AdapterState) js.Ref

func (fn OnAdapterStateChangedEventCallbackFunc) Register() js.Func[func(state *AdapterState)] {
	return js.RegisterCallback[func(state *AdapterState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAdapterStateChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AdapterState
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

type OnAdapterStateChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, state *AdapterState) js.Ref
	Arg T
}

func (cb *OnAdapterStateChangedEventCallback[T]) Register() js.Func[func(state *AdapterState)] {
	return js.RegisterCallback[func(state *AdapterState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAdapterStateChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AdapterState
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

// HasFuncOnAdapterStateChanged returns true if the function "WEBEXT.bluetooth.onAdapterStateChanged.addListener" exists.
func HasFuncOnAdapterStateChanged() bool {
	return js.True == bindings.HasFuncOnAdapterStateChanged()
}

// FuncOnAdapterStateChanged returns the function "WEBEXT.bluetooth.onAdapterStateChanged.addListener".
func FuncOnAdapterStateChanged() (fn js.Func[func(callback js.Func[func(state *AdapterState)])]) {
	bindings.FuncOnAdapterStateChanged(
		js.Pointer(&fn),
	)
	return
}

// OnAdapterStateChanged calls the function "WEBEXT.bluetooth.onAdapterStateChanged.addListener" directly.
func OnAdapterStateChanged(callback js.Func[func(state *AdapterState)]) (ret js.Void) {
	bindings.CallOnAdapterStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAdapterStateChanged calls the function "WEBEXT.bluetooth.onAdapterStateChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAdapterStateChanged(callback js.Func[func(state *AdapterState)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAdapterStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAdapterStateChanged returns true if the function "WEBEXT.bluetooth.onAdapterStateChanged.removeListener" exists.
func HasFuncOffAdapterStateChanged() bool {
	return js.True == bindings.HasFuncOffAdapterStateChanged()
}

// FuncOffAdapterStateChanged returns the function "WEBEXT.bluetooth.onAdapterStateChanged.removeListener".
func FuncOffAdapterStateChanged() (fn js.Func[func(callback js.Func[func(state *AdapterState)])]) {
	bindings.FuncOffAdapterStateChanged(
		js.Pointer(&fn),
	)
	return
}

// OffAdapterStateChanged calls the function "WEBEXT.bluetooth.onAdapterStateChanged.removeListener" directly.
func OffAdapterStateChanged(callback js.Func[func(state *AdapterState)]) (ret js.Void) {
	bindings.CallOffAdapterStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAdapterStateChanged calls the function "WEBEXT.bluetooth.onAdapterStateChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAdapterStateChanged(callback js.Func[func(state *AdapterState)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAdapterStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAdapterStateChanged returns true if the function "WEBEXT.bluetooth.onAdapterStateChanged.hasListener" exists.
func HasFuncHasOnAdapterStateChanged() bool {
	return js.True == bindings.HasFuncHasOnAdapterStateChanged()
}

// FuncHasOnAdapterStateChanged returns the function "WEBEXT.bluetooth.onAdapterStateChanged.hasListener".
func FuncHasOnAdapterStateChanged() (fn js.Func[func(callback js.Func[func(state *AdapterState)]) bool]) {
	bindings.FuncHasOnAdapterStateChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnAdapterStateChanged calls the function "WEBEXT.bluetooth.onAdapterStateChanged.hasListener" directly.
func HasOnAdapterStateChanged(callback js.Func[func(state *AdapterState)]) (ret bool) {
	bindings.CallHasOnAdapterStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAdapterStateChanged calls the function "WEBEXT.bluetooth.onAdapterStateChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAdapterStateChanged(callback js.Func[func(state *AdapterState)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAdapterStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
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

// HasFuncOnDeviceAdded returns true if the function "WEBEXT.bluetooth.onDeviceAdded.addListener" exists.
func HasFuncOnDeviceAdded() bool {
	return js.True == bindings.HasFuncOnDeviceAdded()
}

// FuncOnDeviceAdded returns the function "WEBEXT.bluetooth.onDeviceAdded.addListener".
func FuncOnDeviceAdded() (fn js.Func[func(callback js.Func[func(device *Device)])]) {
	bindings.FuncOnDeviceAdded(
		js.Pointer(&fn),
	)
	return
}

// OnDeviceAdded calls the function "WEBEXT.bluetooth.onDeviceAdded.addListener" directly.
func OnDeviceAdded(callback js.Func[func(device *Device)]) (ret js.Void) {
	bindings.CallOnDeviceAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeviceAdded calls the function "WEBEXT.bluetooth.onDeviceAdded.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeviceAdded(callback js.Func[func(device *Device)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeviceAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeviceAdded returns true if the function "WEBEXT.bluetooth.onDeviceAdded.removeListener" exists.
func HasFuncOffDeviceAdded() bool {
	return js.True == bindings.HasFuncOffDeviceAdded()
}

// FuncOffDeviceAdded returns the function "WEBEXT.bluetooth.onDeviceAdded.removeListener".
func FuncOffDeviceAdded() (fn js.Func[func(callback js.Func[func(device *Device)])]) {
	bindings.FuncOffDeviceAdded(
		js.Pointer(&fn),
	)
	return
}

// OffDeviceAdded calls the function "WEBEXT.bluetooth.onDeviceAdded.removeListener" directly.
func OffDeviceAdded(callback js.Func[func(device *Device)]) (ret js.Void) {
	bindings.CallOffDeviceAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeviceAdded calls the function "WEBEXT.bluetooth.onDeviceAdded.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeviceAdded(callback js.Func[func(device *Device)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeviceAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeviceAdded returns true if the function "WEBEXT.bluetooth.onDeviceAdded.hasListener" exists.
func HasFuncHasOnDeviceAdded() bool {
	return js.True == bindings.HasFuncHasOnDeviceAdded()
}

// FuncHasOnDeviceAdded returns the function "WEBEXT.bluetooth.onDeviceAdded.hasListener".
func FuncHasOnDeviceAdded() (fn js.Func[func(callback js.Func[func(device *Device)]) bool]) {
	bindings.FuncHasOnDeviceAdded(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeviceAdded calls the function "WEBEXT.bluetooth.onDeviceAdded.hasListener" directly.
func HasOnDeviceAdded(callback js.Func[func(device *Device)]) (ret bool) {
	bindings.CallHasOnDeviceAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeviceAdded calls the function "WEBEXT.bluetooth.onDeviceAdded.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeviceAdded(callback js.Func[func(device *Device)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeviceAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDeviceChangedEventCallbackFunc func(this js.Ref, device *Device) js.Ref

func (fn OnDeviceChangedEventCallbackFunc) Register() js.Func[func(device *Device)] {
	return js.RegisterCallback[func(device *Device)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeviceChangedEventCallbackFunc) DispatchCallback(
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

type OnDeviceChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, device *Device) js.Ref
	Arg T
}

func (cb *OnDeviceChangedEventCallback[T]) Register() js.Func[func(device *Device)] {
	return js.RegisterCallback[func(device *Device)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeviceChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnDeviceChanged returns true if the function "WEBEXT.bluetooth.onDeviceChanged.addListener" exists.
func HasFuncOnDeviceChanged() bool {
	return js.True == bindings.HasFuncOnDeviceChanged()
}

// FuncOnDeviceChanged returns the function "WEBEXT.bluetooth.onDeviceChanged.addListener".
func FuncOnDeviceChanged() (fn js.Func[func(callback js.Func[func(device *Device)])]) {
	bindings.FuncOnDeviceChanged(
		js.Pointer(&fn),
	)
	return
}

// OnDeviceChanged calls the function "WEBEXT.bluetooth.onDeviceChanged.addListener" directly.
func OnDeviceChanged(callback js.Func[func(device *Device)]) (ret js.Void) {
	bindings.CallOnDeviceChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeviceChanged calls the function "WEBEXT.bluetooth.onDeviceChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeviceChanged(callback js.Func[func(device *Device)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeviceChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeviceChanged returns true if the function "WEBEXT.bluetooth.onDeviceChanged.removeListener" exists.
func HasFuncOffDeviceChanged() bool {
	return js.True == bindings.HasFuncOffDeviceChanged()
}

// FuncOffDeviceChanged returns the function "WEBEXT.bluetooth.onDeviceChanged.removeListener".
func FuncOffDeviceChanged() (fn js.Func[func(callback js.Func[func(device *Device)])]) {
	bindings.FuncOffDeviceChanged(
		js.Pointer(&fn),
	)
	return
}

// OffDeviceChanged calls the function "WEBEXT.bluetooth.onDeviceChanged.removeListener" directly.
func OffDeviceChanged(callback js.Func[func(device *Device)]) (ret js.Void) {
	bindings.CallOffDeviceChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeviceChanged calls the function "WEBEXT.bluetooth.onDeviceChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeviceChanged(callback js.Func[func(device *Device)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeviceChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeviceChanged returns true if the function "WEBEXT.bluetooth.onDeviceChanged.hasListener" exists.
func HasFuncHasOnDeviceChanged() bool {
	return js.True == bindings.HasFuncHasOnDeviceChanged()
}

// FuncHasOnDeviceChanged returns the function "WEBEXT.bluetooth.onDeviceChanged.hasListener".
func FuncHasOnDeviceChanged() (fn js.Func[func(callback js.Func[func(device *Device)]) bool]) {
	bindings.FuncHasOnDeviceChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeviceChanged calls the function "WEBEXT.bluetooth.onDeviceChanged.hasListener" directly.
func HasOnDeviceChanged(callback js.Func[func(device *Device)]) (ret bool) {
	bindings.CallHasOnDeviceChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeviceChanged calls the function "WEBEXT.bluetooth.onDeviceChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeviceChanged(callback js.Func[func(device *Device)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeviceChanged(
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

// HasFuncOnDeviceRemoved returns true if the function "WEBEXT.bluetooth.onDeviceRemoved.addListener" exists.
func HasFuncOnDeviceRemoved() bool {
	return js.True == bindings.HasFuncOnDeviceRemoved()
}

// FuncOnDeviceRemoved returns the function "WEBEXT.bluetooth.onDeviceRemoved.addListener".
func FuncOnDeviceRemoved() (fn js.Func[func(callback js.Func[func(device *Device)])]) {
	bindings.FuncOnDeviceRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnDeviceRemoved calls the function "WEBEXT.bluetooth.onDeviceRemoved.addListener" directly.
func OnDeviceRemoved(callback js.Func[func(device *Device)]) (ret js.Void) {
	bindings.CallOnDeviceRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeviceRemoved calls the function "WEBEXT.bluetooth.onDeviceRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeviceRemoved(callback js.Func[func(device *Device)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeviceRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeviceRemoved returns true if the function "WEBEXT.bluetooth.onDeviceRemoved.removeListener" exists.
func HasFuncOffDeviceRemoved() bool {
	return js.True == bindings.HasFuncOffDeviceRemoved()
}

// FuncOffDeviceRemoved returns the function "WEBEXT.bluetooth.onDeviceRemoved.removeListener".
func FuncOffDeviceRemoved() (fn js.Func[func(callback js.Func[func(device *Device)])]) {
	bindings.FuncOffDeviceRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffDeviceRemoved calls the function "WEBEXT.bluetooth.onDeviceRemoved.removeListener" directly.
func OffDeviceRemoved(callback js.Func[func(device *Device)]) (ret js.Void) {
	bindings.CallOffDeviceRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeviceRemoved calls the function "WEBEXT.bluetooth.onDeviceRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeviceRemoved(callback js.Func[func(device *Device)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeviceRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeviceRemoved returns true if the function "WEBEXT.bluetooth.onDeviceRemoved.hasListener" exists.
func HasFuncHasOnDeviceRemoved() bool {
	return js.True == bindings.HasFuncHasOnDeviceRemoved()
}

// FuncHasOnDeviceRemoved returns the function "WEBEXT.bluetooth.onDeviceRemoved.hasListener".
func FuncHasOnDeviceRemoved() (fn js.Func[func(callback js.Func[func(device *Device)]) bool]) {
	bindings.FuncHasOnDeviceRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeviceRemoved calls the function "WEBEXT.bluetooth.onDeviceRemoved.hasListener" directly.
func HasOnDeviceRemoved(callback js.Func[func(device *Device)]) (ret bool) {
	bindings.CallHasOnDeviceRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeviceRemoved calls the function "WEBEXT.bluetooth.onDeviceRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeviceRemoved(callback js.Func[func(device *Device)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeviceRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncStartDiscovery returns true if the function "WEBEXT.bluetooth.startDiscovery" exists.
func HasFuncStartDiscovery() bool {
	return js.True == bindings.HasFuncStartDiscovery()
}

// FuncStartDiscovery returns the function "WEBEXT.bluetooth.startDiscovery".
func FuncStartDiscovery() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncStartDiscovery(
		js.Pointer(&fn),
	)
	return
}

// StartDiscovery calls the function "WEBEXT.bluetooth.startDiscovery" directly.
func StartDiscovery() (ret js.Promise[js.Void]) {
	bindings.CallStartDiscovery(
		js.Pointer(&ret),
	)

	return
}

// TryStartDiscovery calls the function "WEBEXT.bluetooth.startDiscovery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartDiscovery() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartDiscovery(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStopDiscovery returns true if the function "WEBEXT.bluetooth.stopDiscovery" exists.
func HasFuncStopDiscovery() bool {
	return js.True == bindings.HasFuncStopDiscovery()
}

// FuncStopDiscovery returns the function "WEBEXT.bluetooth.stopDiscovery".
func FuncStopDiscovery() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncStopDiscovery(
		js.Pointer(&fn),
	)
	return
}

// StopDiscovery calls the function "WEBEXT.bluetooth.stopDiscovery" directly.
func StopDiscovery() (ret js.Promise[js.Void]) {
	bindings.CallStopDiscovery(
		js.Pointer(&ret),
	)

	return
}

// TryStopDiscovery calls the function "WEBEXT.bluetooth.stopDiscovery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStopDiscovery() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStopDiscovery(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
