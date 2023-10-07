// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package bluetoothprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/bluetooth"
	"github.com/primecitizens/pcz/std/plat/js/webext/bluetoothprivate/bindings"
)

type ConnectCallbackFunc func(this js.Ref, result ConnectResultType) js.Ref

func (fn ConnectCallbackFunc) Register() js.Func[func(result ConnectResultType)] {
	return js.RegisterCallback[func(result ConnectResultType)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ConnectCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		ConnectResultType(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ConnectCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result ConnectResultType) js.Ref
	Arg T
}

func (cb *ConnectCallback[T]) Register() js.Func[func(result ConnectResultType)] {
	return js.RegisterCallback[func(result ConnectResultType)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ConnectCallback[T]) DispatchCallback(
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

		ConnectResultType(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ConnectResultType uint32

const (
	_ ConnectResultType = iota

	ConnectResultType_ALREADY_CONNECTED
	ConnectResultType_AUTH_CANCELED
	ConnectResultType_AUTH_FAILED
	ConnectResultType_AUTH_REJECTED
	ConnectResultType_AUTH_TIMEOUT
	ConnectResultType_FAILED
	ConnectResultType_IN_PROGRESS
	ConnectResultType_SUCCESS
	ConnectResultType_UNKNOWN_ERROR
	ConnectResultType_UNSUPPORTED_DEVICE
	ConnectResultType_NOT_READY
	ConnectResultType_ALREADY_EXISTS
	ConnectResultType_NOT_CONNECTED
	ConnectResultType_DOES_NOT_EXIST
	ConnectResultType_INVALID_ARGS
)

func (ConnectResultType) FromRef(str js.Ref) ConnectResultType {
	return ConnectResultType(bindings.ConstOfConnectResultType(str))
}

func (x ConnectResultType) String() (string, bool) {
	switch x {
	case ConnectResultType_ALREADY_CONNECTED:
		return "alreadyConnected", true
	case ConnectResultType_AUTH_CANCELED:
		return "authCanceled", true
	case ConnectResultType_AUTH_FAILED:
		return "authFailed", true
	case ConnectResultType_AUTH_REJECTED:
		return "authRejected", true
	case ConnectResultType_AUTH_TIMEOUT:
		return "authTimeout", true
	case ConnectResultType_FAILED:
		return "failed", true
	case ConnectResultType_IN_PROGRESS:
		return "inProgress", true
	case ConnectResultType_SUCCESS:
		return "success", true
	case ConnectResultType_UNKNOWN_ERROR:
		return "unknownError", true
	case ConnectResultType_UNSUPPORTED_DEVICE:
		return "unsupportedDevice", true
	case ConnectResultType_NOT_READY:
		return "notReady", true
	case ConnectResultType_ALREADY_EXISTS:
		return "alreadyExists", true
	case ConnectResultType_NOT_CONNECTED:
		return "notConnected", true
	case ConnectResultType_DOES_NOT_EXIST:
		return "doesNotExist", true
	case ConnectResultType_INVALID_ARGS:
		return "invalidArgs", true
	default:
		return "", false
	}
}

type TransportType uint32

const (
	_ TransportType = iota

	TransportType_LE
	TransportType_BREDR
	TransportType_DUAL
)

func (TransportType) FromRef(str js.Ref) TransportType {
	return TransportType(bindings.ConstOfTransportType(str))
}

func (x TransportType) String() (string, bool) {
	switch x {
	case TransportType_LE:
		return "le", true
	case TransportType_BREDR:
		return "bredr", true
	case TransportType_DUAL:
		return "dual", true
	default:
		return "", false
	}
}

type OneOf_String_ArrayString struct {
	ref js.Ref
}

func (x OneOf_String_ArrayString) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ArrayString) Free() {
	x.ref.Free()
}

func (x OneOf_String_ArrayString) FromRef(ref js.Ref) OneOf_String_ArrayString {
	return OneOf_String_ArrayString{
		ref: ref,
	}
}

func (x OneOf_String_ArrayString) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayString) ArrayString() js.Array[js.String] {
	return js.Array[js.String]{}.FromRef(x.ref)
}

type DiscoveryFilter struct {
	// Transport is "DiscoveryFilter.transport"
	//
	// Optional
	Transport TransportType
	// Uuids is "DiscoveryFilter.uuids"
	//
	// Optional
	Uuids OneOf_String_ArrayString
	// Rssi is "DiscoveryFilter.rssi"
	//
	// Optional
	//
	// NOTE: FFI_USE_Rssi MUST be set to true to make this field effective.
	Rssi int32
	// Pathloss is "DiscoveryFilter.pathloss"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pathloss MUST be set to true to make this field effective.
	Pathloss int32

	FFI_USE_Rssi     bool // for Rssi.
	FFI_USE_Pathloss bool // for Pathloss.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DiscoveryFilter with all fields set.
func (p DiscoveryFilter) FromRef(ref js.Ref) DiscoveryFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DiscoveryFilter in the application heap.
func (p DiscoveryFilter) New() js.Ref {
	return bindings.DiscoveryFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DiscoveryFilter) UpdateFrom(ref js.Ref) {
	bindings.DiscoveryFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DiscoveryFilter) Update(ref js.Ref) {
	bindings.DiscoveryFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DiscoveryFilter) FreeMembers(recursive bool) {
	js.Free(
		p.Uuids.Ref(),
	)
	p.Uuids = p.Uuids.FromRef(js.Undefined)
}

type NewAdapterState struct {
	// Name is "NewAdapterState.name"
	//
	// Optional
	Name js.String
	// Powered is "NewAdapterState.powered"
	//
	// Optional
	//
	// NOTE: FFI_USE_Powered MUST be set to true to make this field effective.
	Powered bool
	// Discoverable is "NewAdapterState.discoverable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Discoverable MUST be set to true to make this field effective.
	Discoverable bool

	FFI_USE_Powered      bool // for Powered.
	FFI_USE_Discoverable bool // for Discoverable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NewAdapterState with all fields set.
func (p NewAdapterState) FromRef(ref js.Ref) NewAdapterState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NewAdapterState in the application heap.
func (p NewAdapterState) New() js.Ref {
	return bindings.NewAdapterStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NewAdapterState) UpdateFrom(ref js.Ref) {
	bindings.NewAdapterStateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NewAdapterState) Update(ref js.Ref) {
	bindings.NewAdapterStateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NewAdapterState) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
}

type PairingEventType uint32

const (
	_ PairingEventType = iota

	PairingEventType_REQUEST_PINCODE
	PairingEventType_DISPLAY_PINCODE
	PairingEventType_REQUEST_PASSKEY
	PairingEventType_DISPLAY_PASSKEY
	PairingEventType_KEYS_ENTERED
	PairingEventType_CONFIRM_PASSKEY
	PairingEventType_REQUEST_AUTHORIZATION
	PairingEventType_COMPLETE
)

func (PairingEventType) FromRef(str js.Ref) PairingEventType {
	return PairingEventType(bindings.ConstOfPairingEventType(str))
}

func (x PairingEventType) String() (string, bool) {
	switch x {
	case PairingEventType_REQUEST_PINCODE:
		return "requestPincode", true
	case PairingEventType_DISPLAY_PINCODE:
		return "displayPincode", true
	case PairingEventType_REQUEST_PASSKEY:
		return "requestPasskey", true
	case PairingEventType_DISPLAY_PASSKEY:
		return "displayPasskey", true
	case PairingEventType_KEYS_ENTERED:
		return "keysEntered", true
	case PairingEventType_CONFIRM_PASSKEY:
		return "confirmPasskey", true
	case PairingEventType_REQUEST_AUTHORIZATION:
		return "requestAuthorization", true
	case PairingEventType_COMPLETE:
		return "complete", true
	default:
		return "", false
	}
}

type PairingEvent struct {
	// Pairing is "PairingEvent.pairing"
	//
	// Optional
	Pairing PairingEventType
	// Device is "PairingEvent.device"
	//
	// Optional
	//
	// NOTE: Device.FFI_USE MUST be set to true to get Device used.
	Device bluetooth.Device
	// Pincode is "PairingEvent.pincode"
	//
	// Optional
	Pincode js.String
	// Passkey is "PairingEvent.passkey"
	//
	// Optional
	//
	// NOTE: FFI_USE_Passkey MUST be set to true to make this field effective.
	Passkey int32
	// EnteredKey is "PairingEvent.enteredKey"
	//
	// Optional
	//
	// NOTE: FFI_USE_EnteredKey MUST be set to true to make this field effective.
	EnteredKey int32

	FFI_USE_Passkey    bool // for Passkey.
	FFI_USE_EnteredKey bool // for EnteredKey.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PairingEvent with all fields set.
func (p PairingEvent) FromRef(ref js.Ref) PairingEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PairingEvent in the application heap.
func (p PairingEvent) New() js.Ref {
	return bindings.PairingEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PairingEvent) UpdateFrom(ref js.Ref) {
	bindings.PairingEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PairingEvent) Update(ref js.Ref) {
	bindings.PairingEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PairingEvent) FreeMembers(recursive bool) {
	js.Free(
		p.Pincode.Ref(),
	)
	p.Pincode = p.Pincode.FromRef(js.Undefined)
	if recursive {
		p.Device.FreeMembers(true)
	}
}

type PairingResponse uint32

const (
	_ PairingResponse = iota

	PairingResponse_CONFIRM
	PairingResponse_REJECT
	PairingResponse_CANCEL
)

func (PairingResponse) FromRef(str js.Ref) PairingResponse {
	return PairingResponse(bindings.ConstOfPairingResponse(str))
}

func (x PairingResponse) String() (string, bool) {
	switch x {
	case PairingResponse_CONFIRM:
		return "confirm", true
	case PairingResponse_REJECT:
		return "reject", true
	case PairingResponse_CANCEL:
		return "cancel", true
	default:
		return "", false
	}
}

type SetPairingResponseOptions struct {
	// Device is "SetPairingResponseOptions.device"
	//
	// Optional
	//
	// NOTE: Device.FFI_USE MUST be set to true to get Device used.
	Device bluetooth.Device
	// Response is "SetPairingResponseOptions.response"
	//
	// Optional
	Response PairingResponse
	// Pincode is "SetPairingResponseOptions.pincode"
	//
	// Optional
	Pincode js.String
	// Passkey is "SetPairingResponseOptions.passkey"
	//
	// Optional
	//
	// NOTE: FFI_USE_Passkey MUST be set to true to make this field effective.
	Passkey int32

	FFI_USE_Passkey bool // for Passkey.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetPairingResponseOptions with all fields set.
func (p SetPairingResponseOptions) FromRef(ref js.Ref) SetPairingResponseOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetPairingResponseOptions in the application heap.
func (p SetPairingResponseOptions) New() js.Ref {
	return bindings.SetPairingResponseOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetPairingResponseOptions) UpdateFrom(ref js.Ref) {
	bindings.SetPairingResponseOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetPairingResponseOptions) Update(ref js.Ref) {
	bindings.SetPairingResponseOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetPairingResponseOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Pincode.Ref(),
	)
	p.Pincode = p.Pincode.FromRef(js.Undefined)
	if recursive {
		p.Device.FreeMembers(true)
	}
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

// HasFuncConnect returns true if the function "WEBEXT.bluetoothPrivate.connect" exists.
func HasFuncConnect() bool {
	return js.True == bindings.HasFuncConnect()
}

// FuncConnect returns the function "WEBEXT.bluetoothPrivate.connect".
func FuncConnect() (fn js.Func[func(deviceAddress js.String) js.Promise[ConnectResultType]]) {
	bindings.FuncConnect(
		js.Pointer(&fn),
	)
	return
}

// Connect calls the function "WEBEXT.bluetoothPrivate.connect" directly.
func Connect(deviceAddress js.String) (ret js.Promise[ConnectResultType]) {
	bindings.CallConnect(
		js.Pointer(&ret),
		deviceAddress.Ref(),
	)

	return
}

// TryConnect calls the function "WEBEXT.bluetoothPrivate.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryConnect(deviceAddress js.String) (ret js.Promise[ConnectResultType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		deviceAddress.Ref(),
	)

	return
}

// HasFuncDisconnectAll returns true if the function "WEBEXT.bluetoothPrivate.disconnectAll" exists.
func HasFuncDisconnectAll() bool {
	return js.True == bindings.HasFuncDisconnectAll()
}

// FuncDisconnectAll returns the function "WEBEXT.bluetoothPrivate.disconnectAll".
func FuncDisconnectAll() (fn js.Func[func(deviceAddress js.String) js.Promise[js.Void]]) {
	bindings.FuncDisconnectAll(
		js.Pointer(&fn),
	)
	return
}

// DisconnectAll calls the function "WEBEXT.bluetoothPrivate.disconnectAll" directly.
func DisconnectAll(deviceAddress js.String) (ret js.Promise[js.Void]) {
	bindings.CallDisconnectAll(
		js.Pointer(&ret),
		deviceAddress.Ref(),
	)

	return
}

// TryDisconnectAll calls the function "WEBEXT.bluetoothPrivate.disconnectAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDisconnectAll(deviceAddress js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDisconnectAll(
		js.Pointer(&ret), js.Pointer(&exception),
		deviceAddress.Ref(),
	)

	return
}

// HasFuncForgetDevice returns true if the function "WEBEXT.bluetoothPrivate.forgetDevice" exists.
func HasFuncForgetDevice() bool {
	return js.True == bindings.HasFuncForgetDevice()
}

// FuncForgetDevice returns the function "WEBEXT.bluetoothPrivate.forgetDevice".
func FuncForgetDevice() (fn js.Func[func(deviceAddress js.String) js.Promise[js.Void]]) {
	bindings.FuncForgetDevice(
		js.Pointer(&fn),
	)
	return
}

// ForgetDevice calls the function "WEBEXT.bluetoothPrivate.forgetDevice" directly.
func ForgetDevice(deviceAddress js.String) (ret js.Promise[js.Void]) {
	bindings.CallForgetDevice(
		js.Pointer(&ret),
		deviceAddress.Ref(),
	)

	return
}

// TryForgetDevice calls the function "WEBEXT.bluetoothPrivate.forgetDevice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryForgetDevice(deviceAddress js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryForgetDevice(
		js.Pointer(&ret), js.Pointer(&exception),
		deviceAddress.Ref(),
	)

	return
}

type OnDeviceAddressChangedEventCallbackFunc func(this js.Ref, device *bluetooth.Device, oldAddress js.String) js.Ref

func (fn OnDeviceAddressChangedEventCallbackFunc) Register() js.Func[func(device *bluetooth.Device, oldAddress js.String)] {
	return js.RegisterCallback[func(device *bluetooth.Device, oldAddress js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeviceAddressChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 bluetooth.Device
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

type OnDeviceAddressChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, device *bluetooth.Device, oldAddress js.String) js.Ref
	Arg T
}

func (cb *OnDeviceAddressChangedEventCallback[T]) Register() js.Func[func(device *bluetooth.Device, oldAddress js.String)] {
	return js.RegisterCallback[func(device *bluetooth.Device, oldAddress js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeviceAddressChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 bluetooth.Device
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

// HasFuncOnDeviceAddressChanged returns true if the function "WEBEXT.bluetoothPrivate.onDeviceAddressChanged.addListener" exists.
func HasFuncOnDeviceAddressChanged() bool {
	return js.True == bindings.HasFuncOnDeviceAddressChanged()
}

// FuncOnDeviceAddressChanged returns the function "WEBEXT.bluetoothPrivate.onDeviceAddressChanged.addListener".
func FuncOnDeviceAddressChanged() (fn js.Func[func(callback js.Func[func(device *bluetooth.Device, oldAddress js.String)])]) {
	bindings.FuncOnDeviceAddressChanged(
		js.Pointer(&fn),
	)
	return
}

// OnDeviceAddressChanged calls the function "WEBEXT.bluetoothPrivate.onDeviceAddressChanged.addListener" directly.
func OnDeviceAddressChanged(callback js.Func[func(device *bluetooth.Device, oldAddress js.String)]) (ret js.Void) {
	bindings.CallOnDeviceAddressChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeviceAddressChanged calls the function "WEBEXT.bluetoothPrivate.onDeviceAddressChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeviceAddressChanged(callback js.Func[func(device *bluetooth.Device, oldAddress js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeviceAddressChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeviceAddressChanged returns true if the function "WEBEXT.bluetoothPrivate.onDeviceAddressChanged.removeListener" exists.
func HasFuncOffDeviceAddressChanged() bool {
	return js.True == bindings.HasFuncOffDeviceAddressChanged()
}

// FuncOffDeviceAddressChanged returns the function "WEBEXT.bluetoothPrivate.onDeviceAddressChanged.removeListener".
func FuncOffDeviceAddressChanged() (fn js.Func[func(callback js.Func[func(device *bluetooth.Device, oldAddress js.String)])]) {
	bindings.FuncOffDeviceAddressChanged(
		js.Pointer(&fn),
	)
	return
}

// OffDeviceAddressChanged calls the function "WEBEXT.bluetoothPrivate.onDeviceAddressChanged.removeListener" directly.
func OffDeviceAddressChanged(callback js.Func[func(device *bluetooth.Device, oldAddress js.String)]) (ret js.Void) {
	bindings.CallOffDeviceAddressChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeviceAddressChanged calls the function "WEBEXT.bluetoothPrivate.onDeviceAddressChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeviceAddressChanged(callback js.Func[func(device *bluetooth.Device, oldAddress js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeviceAddressChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeviceAddressChanged returns true if the function "WEBEXT.bluetoothPrivate.onDeviceAddressChanged.hasListener" exists.
func HasFuncHasOnDeviceAddressChanged() bool {
	return js.True == bindings.HasFuncHasOnDeviceAddressChanged()
}

// FuncHasOnDeviceAddressChanged returns the function "WEBEXT.bluetoothPrivate.onDeviceAddressChanged.hasListener".
func FuncHasOnDeviceAddressChanged() (fn js.Func[func(callback js.Func[func(device *bluetooth.Device, oldAddress js.String)]) bool]) {
	bindings.FuncHasOnDeviceAddressChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeviceAddressChanged calls the function "WEBEXT.bluetoothPrivate.onDeviceAddressChanged.hasListener" directly.
func HasOnDeviceAddressChanged(callback js.Func[func(device *bluetooth.Device, oldAddress js.String)]) (ret bool) {
	bindings.CallHasOnDeviceAddressChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeviceAddressChanged calls the function "WEBEXT.bluetoothPrivate.onDeviceAddressChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeviceAddressChanged(callback js.Func[func(device *bluetooth.Device, oldAddress js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeviceAddressChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPairingEventCallbackFunc func(this js.Ref, pairingEvent *PairingEvent) js.Ref

func (fn OnPairingEventCallbackFunc) Register() js.Func[func(pairingEvent *PairingEvent)] {
	return js.RegisterCallback[func(pairingEvent *PairingEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPairingEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PairingEvent
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

type OnPairingEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, pairingEvent *PairingEvent) js.Ref
	Arg T
}

func (cb *OnPairingEventCallback[T]) Register() js.Func[func(pairingEvent *PairingEvent)] {
	return js.RegisterCallback[func(pairingEvent *PairingEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPairingEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PairingEvent
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

// HasFuncOnPairing returns true if the function "WEBEXT.bluetoothPrivate.onPairing.addListener" exists.
func HasFuncOnPairing() bool {
	return js.True == bindings.HasFuncOnPairing()
}

// FuncOnPairing returns the function "WEBEXT.bluetoothPrivate.onPairing.addListener".
func FuncOnPairing() (fn js.Func[func(callback js.Func[func(pairingEvent *PairingEvent)])]) {
	bindings.FuncOnPairing(
		js.Pointer(&fn),
	)
	return
}

// OnPairing calls the function "WEBEXT.bluetoothPrivate.onPairing.addListener" directly.
func OnPairing(callback js.Func[func(pairingEvent *PairingEvent)]) (ret js.Void) {
	bindings.CallOnPairing(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPairing calls the function "WEBEXT.bluetoothPrivate.onPairing.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPairing(callback js.Func[func(pairingEvent *PairingEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPairing(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPairing returns true if the function "WEBEXT.bluetoothPrivate.onPairing.removeListener" exists.
func HasFuncOffPairing() bool {
	return js.True == bindings.HasFuncOffPairing()
}

// FuncOffPairing returns the function "WEBEXT.bluetoothPrivate.onPairing.removeListener".
func FuncOffPairing() (fn js.Func[func(callback js.Func[func(pairingEvent *PairingEvent)])]) {
	bindings.FuncOffPairing(
		js.Pointer(&fn),
	)
	return
}

// OffPairing calls the function "WEBEXT.bluetoothPrivate.onPairing.removeListener" directly.
func OffPairing(callback js.Func[func(pairingEvent *PairingEvent)]) (ret js.Void) {
	bindings.CallOffPairing(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPairing calls the function "WEBEXT.bluetoothPrivate.onPairing.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPairing(callback js.Func[func(pairingEvent *PairingEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPairing(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPairing returns true if the function "WEBEXT.bluetoothPrivate.onPairing.hasListener" exists.
func HasFuncHasOnPairing() bool {
	return js.True == bindings.HasFuncHasOnPairing()
}

// FuncHasOnPairing returns the function "WEBEXT.bluetoothPrivate.onPairing.hasListener".
func FuncHasOnPairing() (fn js.Func[func(callback js.Func[func(pairingEvent *PairingEvent)]) bool]) {
	bindings.FuncHasOnPairing(
		js.Pointer(&fn),
	)
	return
}

// HasOnPairing calls the function "WEBEXT.bluetoothPrivate.onPairing.hasListener" directly.
func HasOnPairing(callback js.Func[func(pairingEvent *PairingEvent)]) (ret bool) {
	bindings.CallHasOnPairing(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPairing calls the function "WEBEXT.bluetoothPrivate.onPairing.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPairing(callback js.Func[func(pairingEvent *PairingEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPairing(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncPair returns true if the function "WEBEXT.bluetoothPrivate.pair" exists.
func HasFuncPair() bool {
	return js.True == bindings.HasFuncPair()
}

// FuncPair returns the function "WEBEXT.bluetoothPrivate.pair".
func FuncPair() (fn js.Func[func(deviceAddress js.String) js.Promise[js.Void]]) {
	bindings.FuncPair(
		js.Pointer(&fn),
	)
	return
}

// Pair calls the function "WEBEXT.bluetoothPrivate.pair" directly.
func Pair(deviceAddress js.String) (ret js.Promise[js.Void]) {
	bindings.CallPair(
		js.Pointer(&ret),
		deviceAddress.Ref(),
	)

	return
}

// TryPair calls the function "WEBEXT.bluetoothPrivate.pair"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryPair(deviceAddress js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPair(
		js.Pointer(&ret), js.Pointer(&exception),
		deviceAddress.Ref(),
	)

	return
}

// HasFuncRecordDeviceSelection returns true if the function "WEBEXT.bluetoothPrivate.recordDeviceSelection" exists.
func HasFuncRecordDeviceSelection() bool {
	return js.True == bindings.HasFuncRecordDeviceSelection()
}

// FuncRecordDeviceSelection returns the function "WEBEXT.bluetoothPrivate.recordDeviceSelection".
func FuncRecordDeviceSelection() (fn js.Func[func(selectionDurationMs int32, wasPaired bool, transport bluetooth.Transport)]) {
	bindings.FuncRecordDeviceSelection(
		js.Pointer(&fn),
	)
	return
}

// RecordDeviceSelection calls the function "WEBEXT.bluetoothPrivate.recordDeviceSelection" directly.
func RecordDeviceSelection(selectionDurationMs int32, wasPaired bool, transport bluetooth.Transport) (ret js.Void) {
	bindings.CallRecordDeviceSelection(
		js.Pointer(&ret),
		int32(selectionDurationMs),
		js.Bool(bool(wasPaired)),
		uint32(transport),
	)

	return
}

// TryRecordDeviceSelection calls the function "WEBEXT.bluetoothPrivate.recordDeviceSelection"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordDeviceSelection(selectionDurationMs int32, wasPaired bool, transport bluetooth.Transport) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordDeviceSelection(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(selectionDurationMs),
		js.Bool(bool(wasPaired)),
		uint32(transport),
	)

	return
}

// HasFuncRecordPairing returns true if the function "WEBEXT.bluetoothPrivate.recordPairing" exists.
func HasFuncRecordPairing() bool {
	return js.True == bindings.HasFuncRecordPairing()
}

// FuncRecordPairing returns the function "WEBEXT.bluetoothPrivate.recordPairing".
func FuncRecordPairing() (fn js.Func[func(transport bluetooth.Transport, pairingDurationMs int32, result ConnectResultType)]) {
	bindings.FuncRecordPairing(
		js.Pointer(&fn),
	)
	return
}

// RecordPairing calls the function "WEBEXT.bluetoothPrivate.recordPairing" directly.
func RecordPairing(transport bluetooth.Transport, pairingDurationMs int32, result ConnectResultType) (ret js.Void) {
	bindings.CallRecordPairing(
		js.Pointer(&ret),
		uint32(transport),
		int32(pairingDurationMs),
		uint32(result),
	)

	return
}

// TryRecordPairing calls the function "WEBEXT.bluetoothPrivate.recordPairing"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordPairing(transport bluetooth.Transport, pairingDurationMs int32, result ConnectResultType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordPairing(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(transport),
		int32(pairingDurationMs),
		uint32(result),
	)

	return
}

// HasFuncRecordReconnection returns true if the function "WEBEXT.bluetoothPrivate.recordReconnection" exists.
func HasFuncRecordReconnection() bool {
	return js.True == bindings.HasFuncRecordReconnection()
}

// FuncRecordReconnection returns the function "WEBEXT.bluetoothPrivate.recordReconnection".
func FuncRecordReconnection() (fn js.Func[func(result ConnectResultType)]) {
	bindings.FuncRecordReconnection(
		js.Pointer(&fn),
	)
	return
}

// RecordReconnection calls the function "WEBEXT.bluetoothPrivate.recordReconnection" directly.
func RecordReconnection(result ConnectResultType) (ret js.Void) {
	bindings.CallRecordReconnection(
		js.Pointer(&ret),
		uint32(result),
	)

	return
}

// TryRecordReconnection calls the function "WEBEXT.bluetoothPrivate.recordReconnection"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordReconnection(result ConnectResultType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordReconnection(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(result),
	)

	return
}

// HasFuncSetAdapterState returns true if the function "WEBEXT.bluetoothPrivate.setAdapterState" exists.
func HasFuncSetAdapterState() bool {
	return js.True == bindings.HasFuncSetAdapterState()
}

// FuncSetAdapterState returns the function "WEBEXT.bluetoothPrivate.setAdapterState".
func FuncSetAdapterState() (fn js.Func[func(adapterState NewAdapterState) js.Promise[js.Void]]) {
	bindings.FuncSetAdapterState(
		js.Pointer(&fn),
	)
	return
}

// SetAdapterState calls the function "WEBEXT.bluetoothPrivate.setAdapterState" directly.
func SetAdapterState(adapterState NewAdapterState) (ret js.Promise[js.Void]) {
	bindings.CallSetAdapterState(
		js.Pointer(&ret),
		js.Pointer(&adapterState),
	)

	return
}

// TrySetAdapterState calls the function "WEBEXT.bluetoothPrivate.setAdapterState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetAdapterState(adapterState NewAdapterState) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetAdapterState(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&adapterState),
	)

	return
}

// HasFuncSetDiscoveryFilter returns true if the function "WEBEXT.bluetoothPrivate.setDiscoveryFilter" exists.
func HasFuncSetDiscoveryFilter() bool {
	return js.True == bindings.HasFuncSetDiscoveryFilter()
}

// FuncSetDiscoveryFilter returns the function "WEBEXT.bluetoothPrivate.setDiscoveryFilter".
func FuncSetDiscoveryFilter() (fn js.Func[func(discoveryFilter DiscoveryFilter) js.Promise[js.Void]]) {
	bindings.FuncSetDiscoveryFilter(
		js.Pointer(&fn),
	)
	return
}

// SetDiscoveryFilter calls the function "WEBEXT.bluetoothPrivate.setDiscoveryFilter" directly.
func SetDiscoveryFilter(discoveryFilter DiscoveryFilter) (ret js.Promise[js.Void]) {
	bindings.CallSetDiscoveryFilter(
		js.Pointer(&ret),
		js.Pointer(&discoveryFilter),
	)

	return
}

// TrySetDiscoveryFilter calls the function "WEBEXT.bluetoothPrivate.setDiscoveryFilter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetDiscoveryFilter(discoveryFilter DiscoveryFilter) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetDiscoveryFilter(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&discoveryFilter),
	)

	return
}

// HasFuncSetPairingResponse returns true if the function "WEBEXT.bluetoothPrivate.setPairingResponse" exists.
func HasFuncSetPairingResponse() bool {
	return js.True == bindings.HasFuncSetPairingResponse()
}

// FuncSetPairingResponse returns the function "WEBEXT.bluetoothPrivate.setPairingResponse".
func FuncSetPairingResponse() (fn js.Func[func(options SetPairingResponseOptions) js.Promise[js.Void]]) {
	bindings.FuncSetPairingResponse(
		js.Pointer(&fn),
	)
	return
}

// SetPairingResponse calls the function "WEBEXT.bluetoothPrivate.setPairingResponse" directly.
func SetPairingResponse(options SetPairingResponseOptions) (ret js.Promise[js.Void]) {
	bindings.CallSetPairingResponse(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TrySetPairingResponse calls the function "WEBEXT.bluetoothPrivate.setPairingResponse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPairingResponse(options SetPairingResponseOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPairingResponse(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}
