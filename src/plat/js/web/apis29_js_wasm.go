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

type FlowControlType uint32

const (
	_ FlowControlType = iota

	FlowControlType_NONE
	FlowControlType_HARDWARE
)

func (FlowControlType) FromRef(str js.Ref) FlowControlType {
	return FlowControlType(bindings.ConstOfFlowControlType(str))
}

func (x FlowControlType) String() (string, bool) {
	switch x {
	case FlowControlType_NONE:
		return "none", true
	case FlowControlType_HARDWARE:
		return "hardware", true
	default:
		return "", false
	}
}

type SerialOptions struct {
	// BaudRate is "SerialOptions.baudRate"
	//
	// Required
	BaudRate uint32
	// DataBits is "SerialOptions.dataBits"
	//
	// Optional, defaults to 8.
	//
	// NOTE: FFI_USE_DataBits MUST be set to true to make this field effective.
	DataBits uint8
	// StopBits is "SerialOptions.stopBits"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_StopBits MUST be set to true to make this field effective.
	StopBits uint8
	// Parity is "SerialOptions.parity"
	//
	// Optional, defaults to "none".
	Parity ParityType
	// BufferSize is "SerialOptions.bufferSize"
	//
	// Optional, defaults to 255.
	//
	// NOTE: FFI_USE_BufferSize MUST be set to true to make this field effective.
	BufferSize uint32
	// FlowControl is "SerialOptions.flowControl"
	//
	// Optional, defaults to "none".
	FlowControl FlowControlType

	FFI_USE_DataBits   bool // for DataBits.
	FFI_USE_StopBits   bool // for StopBits.
	FFI_USE_BufferSize bool // for BufferSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SerialOptions with all fields set.
func (p SerialOptions) FromRef(ref js.Ref) SerialOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SerialOptions in the application heap.
func (p SerialOptions) New() js.Ref {
	return bindings.SerialOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SerialOptions) UpdateFrom(ref js.Ref) {
	bindings.SerialOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SerialOptions) Update(ref js.Ref) {
	bindings.SerialOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type SerialOutputSignals struct {
	// DataTerminalReady is "SerialOutputSignals.dataTerminalReady"
	//
	// Optional
	//
	// NOTE: FFI_USE_DataTerminalReady MUST be set to true to make this field effective.
	DataTerminalReady bool
	// RequestToSend is "SerialOutputSignals.requestToSend"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestToSend MUST be set to true to make this field effective.
	RequestToSend bool
	// Break is "SerialOutputSignals.break"
	//
	// Optional
	//
	// NOTE: FFI_USE_Break MUST be set to true to make this field effective.
	Break bool

	FFI_USE_DataTerminalReady bool // for DataTerminalReady.
	FFI_USE_RequestToSend     bool // for RequestToSend.
	FFI_USE_Break             bool // for Break.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SerialOutputSignals with all fields set.
func (p SerialOutputSignals) FromRef(ref js.Ref) SerialOutputSignals {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SerialOutputSignals in the application heap.
func (p SerialOutputSignals) New() js.Ref {
	return bindings.SerialOutputSignalsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SerialOutputSignals) UpdateFrom(ref js.Ref) {
	bindings.SerialOutputSignalsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SerialOutputSignals) Update(ref js.Ref) {
	bindings.SerialOutputSignalsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type SerialInputSignals struct {
	// DataCarrierDetect is "SerialInputSignals.dataCarrierDetect"
	//
	// Required
	DataCarrierDetect bool
	// ClearToSend is "SerialInputSignals.clearToSend"
	//
	// Required
	ClearToSend bool
	// RingIndicator is "SerialInputSignals.ringIndicator"
	//
	// Required
	RingIndicator bool
	// DataSetReady is "SerialInputSignals.dataSetReady"
	//
	// Required
	DataSetReady bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SerialInputSignals with all fields set.
func (p SerialInputSignals) FromRef(ref js.Ref) SerialInputSignals {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SerialInputSignals in the application heap.
func (p SerialInputSignals) New() js.Ref {
	return bindings.SerialInputSignalsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SerialInputSignals) UpdateFrom(ref js.Ref) {
	bindings.SerialInputSignalsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SerialInputSignals) Update(ref js.Ref) {
	bindings.SerialInputSignalsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type SerialPort struct {
	EventTarget
}

func (this SerialPort) Once() SerialPort {
	this.Ref().Once()
	return this
}

func (this SerialPort) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this SerialPort) FromRef(ref js.Ref) SerialPort {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this SerialPort) Free() {
	this.Ref().Free()
}

// Readable returns the value of property "SerialPort.readable".
//
// It returns ok=false if there is no such property.
func (this SerialPort) Readable() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetSerialPortReadable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Writable returns the value of property "SerialPort.writable".
//
// It returns ok=false if there is no such property.
func (this SerialPort) Writable() (ret WritableStream, ok bool) {
	ok = js.True == bindings.GetSerialPortWritable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetInfo returns true if the method "SerialPort.getInfo" exists.
func (this SerialPort) HasGetInfo() bool {
	return js.True == bindings.HasSerialPortGetInfo(
		this.Ref(),
	)
}

// GetInfoFunc returns the method "SerialPort.getInfo".
func (this SerialPort) GetInfoFunc() (fn js.Func[func() SerialPortInfo]) {
	return fn.FromRef(
		bindings.SerialPortGetInfoFunc(
			this.Ref(),
		),
	)
}

// GetInfo calls the method "SerialPort.getInfo".
func (this SerialPort) GetInfo() (ret SerialPortInfo) {
	bindings.CallSerialPortGetInfo(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetInfo calls the method "SerialPort.getInfo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SerialPort) TryGetInfo() (ret SerialPortInfo, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySerialPortGetInfo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasOpen returns true if the method "SerialPort.open" exists.
func (this SerialPort) HasOpen() bool {
	return js.True == bindings.HasSerialPortOpen(
		this.Ref(),
	)
}

// OpenFunc returns the method "SerialPort.open".
func (this SerialPort) OpenFunc() (fn js.Func[func(options SerialOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SerialPortOpenFunc(
			this.Ref(),
		),
	)
}

// Open calls the method "SerialPort.open".
func (this SerialPort) Open(options SerialOptions) (ret js.Promise[js.Void]) {
	bindings.CallSerialPortOpen(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryOpen calls the method "SerialPort.open"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SerialPort) TryOpen(options SerialOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySerialPortOpen(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasSetSignals returns true if the method "SerialPort.setSignals" exists.
func (this SerialPort) HasSetSignals() bool {
	return js.True == bindings.HasSerialPortSetSignals(
		this.Ref(),
	)
}

// SetSignalsFunc returns the method "SerialPort.setSignals".
func (this SerialPort) SetSignalsFunc() (fn js.Func[func(signals SerialOutputSignals) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SerialPortSetSignalsFunc(
			this.Ref(),
		),
	)
}

// SetSignals calls the method "SerialPort.setSignals".
func (this SerialPort) SetSignals(signals SerialOutputSignals) (ret js.Promise[js.Void]) {
	bindings.CallSerialPortSetSignals(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&signals),
	)

	return
}

// TrySetSignals calls the method "SerialPort.setSignals"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SerialPort) TrySetSignals(signals SerialOutputSignals) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySerialPortSetSignals(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&signals),
	)

	return
}

// HasSetSignals1 returns true if the method "SerialPort.setSignals" exists.
func (this SerialPort) HasSetSignals1() bool {
	return js.True == bindings.HasSerialPortSetSignals1(
		this.Ref(),
	)
}

// SetSignals1Func returns the method "SerialPort.setSignals".
func (this SerialPort) SetSignals1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SerialPortSetSignals1Func(
			this.Ref(),
		),
	)
}

// SetSignals1 calls the method "SerialPort.setSignals".
func (this SerialPort) SetSignals1() (ret js.Promise[js.Void]) {
	bindings.CallSerialPortSetSignals1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySetSignals1 calls the method "SerialPort.setSignals"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SerialPort) TrySetSignals1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySerialPortSetSignals1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetSignals returns true if the method "SerialPort.getSignals" exists.
func (this SerialPort) HasGetSignals() bool {
	return js.True == bindings.HasSerialPortGetSignals(
		this.Ref(),
	)
}

// GetSignalsFunc returns the method "SerialPort.getSignals".
func (this SerialPort) GetSignalsFunc() (fn js.Func[func() js.Promise[SerialInputSignals]]) {
	return fn.FromRef(
		bindings.SerialPortGetSignalsFunc(
			this.Ref(),
		),
	)
}

// GetSignals calls the method "SerialPort.getSignals".
func (this SerialPort) GetSignals() (ret js.Promise[SerialInputSignals]) {
	bindings.CallSerialPortGetSignals(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSignals calls the method "SerialPort.getSignals"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SerialPort) TryGetSignals() (ret js.Promise[SerialInputSignals], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySerialPortGetSignals(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "SerialPort.close" exists.
func (this SerialPort) HasClose() bool {
	return js.True == bindings.HasSerialPortClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "SerialPort.close".
func (this SerialPort) CloseFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SerialPortCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "SerialPort.close".
func (this SerialPort) Close() (ret js.Promise[js.Void]) {
	bindings.CallSerialPortClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "SerialPort.close"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SerialPort) TryClose() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySerialPortClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasForget returns true if the method "SerialPort.forget" exists.
func (this SerialPort) HasForget() bool {
	return js.True == bindings.HasSerialPortForget(
		this.Ref(),
	)
}

// ForgetFunc returns the method "SerialPort.forget".
func (this SerialPort) ForgetFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SerialPortForgetFunc(
			this.Ref(),
		),
	)
}

// Forget calls the method "SerialPort.forget".
func (this SerialPort) Forget() (ret js.Promise[js.Void]) {
	bindings.CallSerialPortForget(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryForget calls the method "SerialPort.forget"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SerialPort) TryForget() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySerialPortForget(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SerialPortFilter struct {
	// UsbVendorId is "SerialPortFilter.usbVendorId"
	//
	// Optional
	//
	// NOTE: FFI_USE_UsbVendorId MUST be set to true to make this field effective.
	UsbVendorId uint16
	// UsbProductId is "SerialPortFilter.usbProductId"
	//
	// Optional
	//
	// NOTE: FFI_USE_UsbProductId MUST be set to true to make this field effective.
	UsbProductId uint16
	// BluetoothServiceClassId is "SerialPortFilter.bluetoothServiceClassId"
	//
	// Optional
	BluetoothServiceClassId BluetoothServiceUUID

	FFI_USE_UsbVendorId  bool // for UsbVendorId.
	FFI_USE_UsbProductId bool // for UsbProductId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SerialPortFilter with all fields set.
func (p SerialPortFilter) FromRef(ref js.Ref) SerialPortFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SerialPortFilter in the application heap.
func (p SerialPortFilter) New() js.Ref {
	return bindings.SerialPortFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SerialPortFilter) UpdateFrom(ref js.Ref) {
	bindings.SerialPortFilterJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SerialPortFilter) Update(ref js.Ref) {
	bindings.SerialPortFilterJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type SerialPortRequestOptions struct {
	// Filters is "SerialPortRequestOptions.filters"
	//
	// Optional
	Filters js.Array[SerialPortFilter]
	// AllowedBluetoothServiceClassIds is "SerialPortRequestOptions.allowedBluetoothServiceClassIds"
	//
	// Optional
	AllowedBluetoothServiceClassIds js.Array[BluetoothServiceUUID]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SerialPortRequestOptions with all fields set.
func (p SerialPortRequestOptions) FromRef(ref js.Ref) SerialPortRequestOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SerialPortRequestOptions in the application heap.
func (p SerialPortRequestOptions) New() js.Ref {
	return bindings.SerialPortRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SerialPortRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.SerialPortRequestOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SerialPortRequestOptions) Update(ref js.Ref) {
	bindings.SerialPortRequestOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type Serial struct {
	EventTarget
}

func (this Serial) Once() Serial {
	this.Ref().Once()
	return this
}

func (this Serial) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this Serial) FromRef(ref js.Ref) Serial {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this Serial) Free() {
	this.Ref().Free()
}

// HasGetPorts returns true if the method "Serial.getPorts" exists.
func (this Serial) HasGetPorts() bool {
	return js.True == bindings.HasSerialGetPorts(
		this.Ref(),
	)
}

// GetPortsFunc returns the method "Serial.getPorts".
func (this Serial) GetPortsFunc() (fn js.Func[func() js.Promise[js.Array[SerialPort]]]) {
	return fn.FromRef(
		bindings.SerialGetPortsFunc(
			this.Ref(),
		),
	)
}

// GetPorts calls the method "Serial.getPorts".
func (this Serial) GetPorts() (ret js.Promise[js.Array[SerialPort]]) {
	bindings.CallSerialGetPorts(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetPorts calls the method "Serial.getPorts"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Serial) TryGetPorts() (ret js.Promise[js.Array[SerialPort]], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySerialGetPorts(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRequestPort returns true if the method "Serial.requestPort" exists.
func (this Serial) HasRequestPort() bool {
	return js.True == bindings.HasSerialRequestPort(
		this.Ref(),
	)
}

// RequestPortFunc returns the method "Serial.requestPort".
func (this Serial) RequestPortFunc() (fn js.Func[func(options SerialPortRequestOptions) js.Promise[SerialPort]]) {
	return fn.FromRef(
		bindings.SerialRequestPortFunc(
			this.Ref(),
		),
	)
}

// RequestPort calls the method "Serial.requestPort".
func (this Serial) RequestPort(options SerialPortRequestOptions) (ret js.Promise[SerialPort]) {
	bindings.CallSerialRequestPort(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestPort calls the method "Serial.requestPort"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Serial) TryRequestPort(options SerialPortRequestOptions) (ret js.Promise[SerialPort], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySerialRequestPort(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasRequestPort1 returns true if the method "Serial.requestPort" exists.
func (this Serial) HasRequestPort1() bool {
	return js.True == bindings.HasSerialRequestPort1(
		this.Ref(),
	)
}

// RequestPort1Func returns the method "Serial.requestPort".
func (this Serial) RequestPort1Func() (fn js.Func[func() js.Promise[SerialPort]]) {
	return fn.FromRef(
		bindings.SerialRequestPort1Func(
			this.Ref(),
		),
	)
}

// RequestPort1 calls the method "Serial.requestPort".
func (this Serial) RequestPort1() (ret js.Promise[SerialPort]) {
	bindings.CallSerialRequestPort1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestPort1 calls the method "Serial.requestPort"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Serial) TryRequestPort1() (ret js.Promise[SerialPort], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySerialRequestPort1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type MediaDecodingType uint32

const (
	_ MediaDecodingType = iota

	MediaDecodingType_FILE
	MediaDecodingType_MEDIA_SOURCE
	MediaDecodingType_WEBRTC
)

func (MediaDecodingType) FromRef(str js.Ref) MediaDecodingType {
	return MediaDecodingType(bindings.ConstOfMediaDecodingType(str))
}

func (x MediaDecodingType) String() (string, bool) {
	switch x {
	case MediaDecodingType_FILE:
		return "file", true
	case MediaDecodingType_MEDIA_SOURCE:
		return "media-source", true
	case MediaDecodingType_WEBRTC:
		return "webrtc", true
	default:
		return "", false
	}
}

type KeySystemTrackConfiguration struct {
	// Robustness is "KeySystemTrackConfiguration.robustness"
	//
	// Optional, defaults to "".
	Robustness js.String
	// EncryptionScheme is "KeySystemTrackConfiguration.encryptionScheme"
	//
	// Optional, defaults to null.
	EncryptionScheme js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a KeySystemTrackConfiguration with all fields set.
func (p KeySystemTrackConfiguration) FromRef(ref js.Ref) KeySystemTrackConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new KeySystemTrackConfiguration in the application heap.
func (p KeySystemTrackConfiguration) New() js.Ref {
	return bindings.KeySystemTrackConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p KeySystemTrackConfiguration) UpdateFrom(ref js.Ref) {
	bindings.KeySystemTrackConfigurationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p KeySystemTrackConfiguration) Update(ref js.Ref) {
	bindings.KeySystemTrackConfigurationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaCapabilitiesKeySystemConfiguration struct {
	// KeySystem is "MediaCapabilitiesKeySystemConfiguration.keySystem"
	//
	// Required
	KeySystem js.String
	// InitDataType is "MediaCapabilitiesKeySystemConfiguration.initDataType"
	//
	// Optional, defaults to "".
	InitDataType js.String
	// DistinctiveIdentifier is "MediaCapabilitiesKeySystemConfiguration.distinctiveIdentifier"
	//
	// Optional, defaults to "optional".
	DistinctiveIdentifier MediaKeysRequirement
	// PersistentState is "MediaCapabilitiesKeySystemConfiguration.persistentState"
	//
	// Optional, defaults to "optional".
	PersistentState MediaKeysRequirement
	// SessionTypes is "MediaCapabilitiesKeySystemConfiguration.sessionTypes"
	//
	// Optional
	SessionTypes js.Array[js.String]
	// Audio is "MediaCapabilitiesKeySystemConfiguration.audio"
	//
	// Optional
	//
	// NOTE: Audio.FFI_USE MUST be set to true to get Audio used.
	Audio KeySystemTrackConfiguration
	// Video is "MediaCapabilitiesKeySystemConfiguration.video"
	//
	// Optional
	//
	// NOTE: Video.FFI_USE MUST be set to true to get Video used.
	Video KeySystemTrackConfiguration

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaCapabilitiesKeySystemConfiguration with all fields set.
func (p MediaCapabilitiesKeySystemConfiguration) FromRef(ref js.Ref) MediaCapabilitiesKeySystemConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaCapabilitiesKeySystemConfiguration in the application heap.
func (p MediaCapabilitiesKeySystemConfiguration) New() js.Ref {
	return bindings.MediaCapabilitiesKeySystemConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaCapabilitiesKeySystemConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MediaCapabilitiesKeySystemConfigurationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaCapabilitiesKeySystemConfiguration) Update(ref js.Ref) {
	bindings.MediaCapabilitiesKeySystemConfigurationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type HdrMetadataType uint32

const (
	_ HdrMetadataType = iota

	HdrMetadataType_SMPTE_ST2086
	HdrMetadataType_SMPTE_ST2094_10
	HdrMetadataType_SMPTE_ST2094_40
)

func (HdrMetadataType) FromRef(str js.Ref) HdrMetadataType {
	return HdrMetadataType(bindings.ConstOfHdrMetadataType(str))
}

func (x HdrMetadataType) String() (string, bool) {
	switch x {
	case HdrMetadataType_SMPTE_ST2086:
		return "smpteSt2086", true
	case HdrMetadataType_SMPTE_ST2094_10:
		return "smpteSt2094-10", true
	case HdrMetadataType_SMPTE_ST2094_40:
		return "smpteSt2094-40", true
	default:
		return "", false
	}
}

type TransferFunction uint32

const (
	_ TransferFunction = iota

	TransferFunction_SRGB
	TransferFunction_PQ
	TransferFunction_HLG
)

func (TransferFunction) FromRef(str js.Ref) TransferFunction {
	return TransferFunction(bindings.ConstOfTransferFunction(str))
}

func (x TransferFunction) String() (string, bool) {
	switch x {
	case TransferFunction_SRGB:
		return "srgb", true
	case TransferFunction_PQ:
		return "pq", true
	case TransferFunction_HLG:
		return "hlg", true
	default:
		return "", false
	}
}

type VideoConfiguration struct {
	// ContentType is "VideoConfiguration.contentType"
	//
	// Required
	ContentType js.String
	// Width is "VideoConfiguration.width"
	//
	// Required
	Width uint32
	// Height is "VideoConfiguration.height"
	//
	// Required
	Height uint32
	// Bitrate is "VideoConfiguration.bitrate"
	//
	// Required
	Bitrate uint64
	// Framerate is "VideoConfiguration.framerate"
	//
	// Required
	Framerate float64
	// HasAlphaChannel is "VideoConfiguration.hasAlphaChannel"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasAlphaChannel MUST be set to true to make this field effective.
	HasAlphaChannel bool
	// HdrMetadataType is "VideoConfiguration.hdrMetadataType"
	//
	// Optional
	HdrMetadataType HdrMetadataType
	// ColorGamut is "VideoConfiguration.colorGamut"
	//
	// Optional
	ColorGamut ColorGamut
	// TransferFunction is "VideoConfiguration.transferFunction"
	//
	// Optional
	TransferFunction TransferFunction
	// ScalabilityMode is "VideoConfiguration.scalabilityMode"
	//
	// Optional
	ScalabilityMode js.String
	// SpatialScalability is "VideoConfiguration.spatialScalability"
	//
	// Optional
	//
	// NOTE: FFI_USE_SpatialScalability MUST be set to true to make this field effective.
	SpatialScalability bool

	FFI_USE_HasAlphaChannel    bool // for HasAlphaChannel.
	FFI_USE_SpatialScalability bool // for SpatialScalability.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoConfiguration with all fields set.
func (p VideoConfiguration) FromRef(ref js.Ref) VideoConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoConfiguration in the application heap.
func (p VideoConfiguration) New() js.Ref {
	return bindings.VideoConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoConfiguration) UpdateFrom(ref js.Ref) {
	bindings.VideoConfigurationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoConfiguration) Update(ref js.Ref) {
	bindings.VideoConfigurationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaDecodingConfiguration struct {
	// Type is "MediaDecodingConfiguration.type"
	//
	// Required
	Type MediaDecodingType
	// KeySystemConfiguration is "MediaDecodingConfiguration.keySystemConfiguration"
	//
	// Optional
	//
	// NOTE: KeySystemConfiguration.FFI_USE MUST be set to true to get KeySystemConfiguration used.
	KeySystemConfiguration MediaCapabilitiesKeySystemConfiguration
	// Video is "MediaDecodingConfiguration.video"
	//
	// Optional
	//
	// NOTE: Video.FFI_USE MUST be set to true to get Video used.
	Video VideoConfiguration
	// Audio is "MediaDecodingConfiguration.audio"
	//
	// Optional
	//
	// NOTE: Audio.FFI_USE MUST be set to true to get Audio used.
	Audio AudioConfiguration

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaDecodingConfiguration with all fields set.
func (p MediaDecodingConfiguration) FromRef(ref js.Ref) MediaDecodingConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaDecodingConfiguration in the application heap.
func (p MediaDecodingConfiguration) New() js.Ref {
	return bindings.MediaDecodingConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaDecodingConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MediaDecodingConfigurationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaDecodingConfiguration) Update(ref js.Ref) {
	bindings.MediaDecodingConfigurationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaCapabilitiesDecodingInfo struct {
	// KeySystemAccess is "MediaCapabilitiesDecodingInfo.keySystemAccess"
	//
	// Required
	KeySystemAccess MediaKeySystemAccess
	// Configuration is "MediaCapabilitiesDecodingInfo.configuration"
	//
	// Optional
	//
	// NOTE: Configuration.FFI_USE MUST be set to true to get Configuration used.
	Configuration MediaDecodingConfiguration
	// Supported is "MediaCapabilitiesDecodingInfo.supported"
	//
	// Required
	Supported bool
	// Smooth is "MediaCapabilitiesDecodingInfo.smooth"
	//
	// Required
	Smooth bool
	// PowerEfficient is "MediaCapabilitiesDecodingInfo.powerEfficient"
	//
	// Required
	PowerEfficient bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaCapabilitiesDecodingInfo with all fields set.
func (p MediaCapabilitiesDecodingInfo) FromRef(ref js.Ref) MediaCapabilitiesDecodingInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaCapabilitiesDecodingInfo in the application heap.
func (p MediaCapabilitiesDecodingInfo) New() js.Ref {
	return bindings.MediaCapabilitiesDecodingInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaCapabilitiesDecodingInfo) UpdateFrom(ref js.Ref) {
	bindings.MediaCapabilitiesDecodingInfoJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaCapabilitiesDecodingInfo) Update(ref js.Ref) {
	bindings.MediaCapabilitiesDecodingInfoJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaEncodingType uint32

const (
	_ MediaEncodingType = iota

	MediaEncodingType_RECORD
	MediaEncodingType_WEBRTC
)

func (MediaEncodingType) FromRef(str js.Ref) MediaEncodingType {
	return MediaEncodingType(bindings.ConstOfMediaEncodingType(str))
}

func (x MediaEncodingType) String() (string, bool) {
	switch x {
	case MediaEncodingType_RECORD:
		return "record", true
	case MediaEncodingType_WEBRTC:
		return "webrtc", true
	default:
		return "", false
	}
}

type MediaEncodingConfiguration struct {
	// Type is "MediaEncodingConfiguration.type"
	//
	// Required
	Type MediaEncodingType
	// Video is "MediaEncodingConfiguration.video"
	//
	// Optional
	//
	// NOTE: Video.FFI_USE MUST be set to true to get Video used.
	Video VideoConfiguration
	// Audio is "MediaEncodingConfiguration.audio"
	//
	// Optional
	//
	// NOTE: Audio.FFI_USE MUST be set to true to get Audio used.
	Audio AudioConfiguration

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaEncodingConfiguration with all fields set.
func (p MediaEncodingConfiguration) FromRef(ref js.Ref) MediaEncodingConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaEncodingConfiguration in the application heap.
func (p MediaEncodingConfiguration) New() js.Ref {
	return bindings.MediaEncodingConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaEncodingConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MediaEncodingConfigurationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaEncodingConfiguration) Update(ref js.Ref) {
	bindings.MediaEncodingConfigurationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaCapabilitiesEncodingInfo struct {
	// Configuration is "MediaCapabilitiesEncodingInfo.configuration"
	//
	// Optional
	//
	// NOTE: Configuration.FFI_USE MUST be set to true to get Configuration used.
	Configuration MediaEncodingConfiguration
	// Supported is "MediaCapabilitiesEncodingInfo.supported"
	//
	// Required
	Supported bool
	// Smooth is "MediaCapabilitiesEncodingInfo.smooth"
	//
	// Required
	Smooth bool
	// PowerEfficient is "MediaCapabilitiesEncodingInfo.powerEfficient"
	//
	// Required
	PowerEfficient bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaCapabilitiesEncodingInfo with all fields set.
func (p MediaCapabilitiesEncodingInfo) FromRef(ref js.Ref) MediaCapabilitiesEncodingInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaCapabilitiesEncodingInfo in the application heap.
func (p MediaCapabilitiesEncodingInfo) New() js.Ref {
	return bindings.MediaCapabilitiesEncodingInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaCapabilitiesEncodingInfo) UpdateFrom(ref js.Ref) {
	bindings.MediaCapabilitiesEncodingInfoJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaCapabilitiesEncodingInfo) Update(ref js.Ref) {
	bindings.MediaCapabilitiesEncodingInfoJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaCapabilities struct {
	ref js.Ref
}

func (this MediaCapabilities) Once() MediaCapabilities {
	this.Ref().Once()
	return this
}

func (this MediaCapabilities) Ref() js.Ref {
	return this.ref
}

func (this MediaCapabilities) FromRef(ref js.Ref) MediaCapabilities {
	this.ref = ref
	return this
}

func (this MediaCapabilities) Free() {
	this.Ref().Free()
}

// HasDecodingInfo returns true if the method "MediaCapabilities.decodingInfo" exists.
func (this MediaCapabilities) HasDecodingInfo() bool {
	return js.True == bindings.HasMediaCapabilitiesDecodingInfo(
		this.Ref(),
	)
}

// DecodingInfoFunc returns the method "MediaCapabilities.decodingInfo".
func (this MediaCapabilities) DecodingInfoFunc() (fn js.Func[func(configuration MediaDecodingConfiguration) js.Promise[MediaCapabilitiesDecodingInfo]]) {
	return fn.FromRef(
		bindings.MediaCapabilitiesDecodingInfoFunc(
			this.Ref(),
		),
	)
}

// DecodingInfo calls the method "MediaCapabilities.decodingInfo".
func (this MediaCapabilities) DecodingInfo(configuration MediaDecodingConfiguration) (ret js.Promise[MediaCapabilitiesDecodingInfo]) {
	bindings.CallMediaCapabilitiesDecodingInfo(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&configuration),
	)

	return
}

// TryDecodingInfo calls the method "MediaCapabilities.decodingInfo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaCapabilities) TryDecodingInfo(configuration MediaDecodingConfiguration) (ret js.Promise[MediaCapabilitiesDecodingInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaCapabilitiesDecodingInfo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&configuration),
	)

	return
}

// HasEncodingInfo returns true if the method "MediaCapabilities.encodingInfo" exists.
func (this MediaCapabilities) HasEncodingInfo() bool {
	return js.True == bindings.HasMediaCapabilitiesEncodingInfo(
		this.Ref(),
	)
}

// EncodingInfoFunc returns the method "MediaCapabilities.encodingInfo".
func (this MediaCapabilities) EncodingInfoFunc() (fn js.Func[func(configuration MediaEncodingConfiguration) js.Promise[MediaCapabilitiesEncodingInfo]]) {
	return fn.FromRef(
		bindings.MediaCapabilitiesEncodingInfoFunc(
			this.Ref(),
		),
	)
}

// EncodingInfo calls the method "MediaCapabilities.encodingInfo".
func (this MediaCapabilities) EncodingInfo(configuration MediaEncodingConfiguration) (ret js.Promise[MediaCapabilitiesEncodingInfo]) {
	bindings.CallMediaCapabilitiesEncodingInfo(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&configuration),
	)

	return
}

// TryEncodingInfo calls the method "MediaCapabilities.encodingInfo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaCapabilities) TryEncodingInfo(configuration MediaEncodingConfiguration) (ret js.Promise[MediaCapabilitiesEncodingInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaCapabilitiesEncodingInfo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&configuration),
	)

	return
}

type UserActivation struct {
	ref js.Ref
}

func (this UserActivation) Once() UserActivation {
	this.Ref().Once()
	return this
}

func (this UserActivation) Ref() js.Ref {
	return this.ref
}

func (this UserActivation) FromRef(ref js.Ref) UserActivation {
	this.ref = ref
	return this
}

func (this UserActivation) Free() {
	this.Ref().Free()
}

// HasBeenActive returns the value of property "UserActivation.hasBeenActive".
//
// It returns ok=false if there is no such property.
func (this UserActivation) HasBeenActive() (ret bool, ok bool) {
	ok = js.True == bindings.GetUserActivationHasBeenActive(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsActive returns the value of property "UserActivation.isActive".
//
// It returns ok=false if there is no such property.
func (this UserActivation) IsActive() (ret bool, ok bool) {
	ok = js.True == bindings.GetUserActivationIsActive(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type PermissionStatus struct {
	EventTarget
}

func (this PermissionStatus) Once() PermissionStatus {
	this.Ref().Once()
	return this
}

func (this PermissionStatus) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this PermissionStatus) FromRef(ref js.Ref) PermissionStatus {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this PermissionStatus) Free() {
	this.Ref().Free()
}

// State returns the value of property "PermissionStatus.state".
//
// It returns ok=false if there is no such property.
func (this PermissionStatus) State() (ret PermissionState, ok bool) {
	ok = js.True == bindings.GetPermissionStatusState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "PermissionStatus.name".
//
// It returns ok=false if there is no such property.
func (this PermissionStatus) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPermissionStatusName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type Permissions struct {
	ref js.Ref
}

func (this Permissions) Once() Permissions {
	this.Ref().Once()
	return this
}

func (this Permissions) Ref() js.Ref {
	return this.ref
}

func (this Permissions) FromRef(ref js.Ref) Permissions {
	this.ref = ref
	return this
}

func (this Permissions) Free() {
	this.Ref().Free()
}

// HasQuery returns true if the method "Permissions.query" exists.
func (this Permissions) HasQuery() bool {
	return js.True == bindings.HasPermissionsQuery(
		this.Ref(),
	)
}

// QueryFunc returns the method "Permissions.query".
func (this Permissions) QueryFunc() (fn js.Func[func(permissionDesc js.Object) js.Promise[PermissionStatus]]) {
	return fn.FromRef(
		bindings.PermissionsQueryFunc(
			this.Ref(),
		),
	)
}

// Query calls the method "Permissions.query".
func (this Permissions) Query(permissionDesc js.Object) (ret js.Promise[PermissionStatus]) {
	bindings.CallPermissionsQuery(
		this.Ref(), js.Pointer(&ret),
		permissionDesc.Ref(),
	)

	return
}

// TryQuery calls the method "Permissions.query"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Permissions) TryQuery(permissionDesc js.Object) (ret js.Promise[PermissionStatus], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPermissionsQuery(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		permissionDesc.Ref(),
	)

	return
}

// HasRevoke returns true if the method "Permissions.revoke" exists.
func (this Permissions) HasRevoke() bool {
	return js.True == bindings.HasPermissionsRevoke(
		this.Ref(),
	)
}

// RevokeFunc returns the method "Permissions.revoke".
func (this Permissions) RevokeFunc() (fn js.Func[func(permissionDesc js.Object) js.Promise[PermissionStatus]]) {
	return fn.FromRef(
		bindings.PermissionsRevokeFunc(
			this.Ref(),
		),
	)
}

// Revoke calls the method "Permissions.revoke".
func (this Permissions) Revoke(permissionDesc js.Object) (ret js.Promise[PermissionStatus]) {
	bindings.CallPermissionsRevoke(
		this.Ref(), js.Pointer(&ret),
		permissionDesc.Ref(),
	)

	return
}

// TryRevoke calls the method "Permissions.revoke"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Permissions) TryRevoke(permissionDesc js.Object) (ret js.Promise[PermissionStatus], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPermissionsRevoke(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		permissionDesc.Ref(),
	)

	return
}

// HasRequest returns true if the method "Permissions.request" exists.
func (this Permissions) HasRequest() bool {
	return js.True == bindings.HasPermissionsRequest(
		this.Ref(),
	)
}

// RequestFunc returns the method "Permissions.request".
func (this Permissions) RequestFunc() (fn js.Func[func(permissionDesc js.Object) js.Promise[PermissionStatus]]) {
	return fn.FromRef(
		bindings.PermissionsRequestFunc(
			this.Ref(),
		),
	)
}

// Request calls the method "Permissions.request".
func (this Permissions) Request(permissionDesc js.Object) (ret js.Promise[PermissionStatus]) {
	bindings.CallPermissionsRequest(
		this.Ref(), js.Pointer(&ret),
		permissionDesc.Ref(),
	)

	return
}

// TryRequest calls the method "Permissions.request"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Permissions) TryRequest(permissionDesc js.Object) (ret js.Promise[PermissionStatus], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPermissionsRequest(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		permissionDesc.Ref(),
	)

	return
}

type ContactProperty uint32

const (
	_ ContactProperty = iota

	ContactProperty_ADDRESS
	ContactProperty_EMAIL
	ContactProperty_ICON
	ContactProperty_NAME
	ContactProperty_TEL
)

func (ContactProperty) FromRef(str js.Ref) ContactProperty {
	return ContactProperty(bindings.ConstOfContactProperty(str))
}

func (x ContactProperty) String() (string, bool) {
	switch x {
	case ContactProperty_ADDRESS:
		return "address", true
	case ContactProperty_EMAIL:
		return "email", true
	case ContactProperty_ICON:
		return "icon", true
	case ContactProperty_NAME:
		return "name", true
	case ContactProperty_TEL:
		return "tel", true
	default:
		return "", false
	}
}

type ContactAddress struct {
	ref js.Ref
}

func (this ContactAddress) Once() ContactAddress {
	this.Ref().Once()
	return this
}

func (this ContactAddress) Ref() js.Ref {
	return this.ref
}

func (this ContactAddress) FromRef(ref js.Ref) ContactAddress {
	this.ref = ref
	return this
}

func (this ContactAddress) Free() {
	this.Ref().Free()
}

// City returns the value of property "ContactAddress.city".
//
// It returns ok=false if there is no such property.
func (this ContactAddress) City() (ret js.String, ok bool) {
	ok = js.True == bindings.GetContactAddressCity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Country returns the value of property "ContactAddress.country".
//
// It returns ok=false if there is no such property.
func (this ContactAddress) Country() (ret js.String, ok bool) {
	ok = js.True == bindings.GetContactAddressCountry(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DependentLocality returns the value of property "ContactAddress.dependentLocality".
//
// It returns ok=false if there is no such property.
func (this ContactAddress) DependentLocality() (ret js.String, ok bool) {
	ok = js.True == bindings.GetContactAddressDependentLocality(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Organization returns the value of property "ContactAddress.organization".
//
// It returns ok=false if there is no such property.
func (this ContactAddress) Organization() (ret js.String, ok bool) {
	ok = js.True == bindings.GetContactAddressOrganization(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Phone returns the value of property "ContactAddress.phone".
//
// It returns ok=false if there is no such property.
func (this ContactAddress) Phone() (ret js.String, ok bool) {
	ok = js.True == bindings.GetContactAddressPhone(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PostalCode returns the value of property "ContactAddress.postalCode".
//
// It returns ok=false if there is no such property.
func (this ContactAddress) PostalCode() (ret js.String, ok bool) {
	ok = js.True == bindings.GetContactAddressPostalCode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Recipient returns the value of property "ContactAddress.recipient".
//
// It returns ok=false if there is no such property.
func (this ContactAddress) Recipient() (ret js.String, ok bool) {
	ok = js.True == bindings.GetContactAddressRecipient(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Region returns the value of property "ContactAddress.region".
//
// It returns ok=false if there is no such property.
func (this ContactAddress) Region() (ret js.String, ok bool) {
	ok = js.True == bindings.GetContactAddressRegion(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SortingCode returns the value of property "ContactAddress.sortingCode".
//
// It returns ok=false if there is no such property.
func (this ContactAddress) SortingCode() (ret js.String, ok bool) {
	ok = js.True == bindings.GetContactAddressSortingCode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AddressLine returns the value of property "ContactAddress.addressLine".
//
// It returns ok=false if there is no such property.
func (this ContactAddress) AddressLine() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetContactAddressAddressLine(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "ContactAddress.toJSON" exists.
func (this ContactAddress) HasToJSON() bool {
	return js.True == bindings.HasContactAddressToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "ContactAddress.toJSON".
func (this ContactAddress) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.ContactAddressToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "ContactAddress.toJSON".
func (this ContactAddress) ToJSON() (ret js.Object) {
	bindings.CallContactAddressToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "ContactAddress.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ContactAddress) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryContactAddressToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ContactInfo struct {
	// Address is "ContactInfo.address"
	//
	// Optional
	Address js.Array[ContactAddress]
	// Email is "ContactInfo.email"
	//
	// Optional
	Email js.Array[js.String]
	// Icon is "ContactInfo.icon"
	//
	// Optional
	Icon js.Array[Blob]
	// Name is "ContactInfo.name"
	//
	// Optional
	Name js.Array[js.String]
	// Tel is "ContactInfo.tel"
	//
	// Optional
	Tel js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContactInfo with all fields set.
func (p ContactInfo) FromRef(ref js.Ref) ContactInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ContactInfo in the application heap.
func (p ContactInfo) New() js.Ref {
	return bindings.ContactInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ContactInfo) UpdateFrom(ref js.Ref) {
	bindings.ContactInfoJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ContactInfo) Update(ref js.Ref) {
	bindings.ContactInfoJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ContactsSelectOptions struct {
	// Multiple is "ContactsSelectOptions.multiple"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Multiple MUST be set to true to make this field effective.
	Multiple bool

	FFI_USE_Multiple bool // for Multiple.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContactsSelectOptions with all fields set.
func (p ContactsSelectOptions) FromRef(ref js.Ref) ContactsSelectOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ContactsSelectOptions in the application heap.
func (p ContactsSelectOptions) New() js.Ref {
	return bindings.ContactsSelectOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ContactsSelectOptions) UpdateFrom(ref js.Ref) {
	bindings.ContactsSelectOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ContactsSelectOptions) Update(ref js.Ref) {
	bindings.ContactsSelectOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ContactsManager struct {
	ref js.Ref
}

func (this ContactsManager) Once() ContactsManager {
	this.Ref().Once()
	return this
}

func (this ContactsManager) Ref() js.Ref {
	return this.ref
}

func (this ContactsManager) FromRef(ref js.Ref) ContactsManager {
	this.ref = ref
	return this
}

func (this ContactsManager) Free() {
	this.Ref().Free()
}

// HasGetProperties returns true if the method "ContactsManager.getProperties" exists.
func (this ContactsManager) HasGetProperties() bool {
	return js.True == bindings.HasContactsManagerGetProperties(
		this.Ref(),
	)
}

// GetPropertiesFunc returns the method "ContactsManager.getProperties".
func (this ContactsManager) GetPropertiesFunc() (fn js.Func[func() js.Promise[js.Array[ContactProperty]]]) {
	return fn.FromRef(
		bindings.ContactsManagerGetPropertiesFunc(
			this.Ref(),
		),
	)
}

// GetProperties calls the method "ContactsManager.getProperties".
func (this ContactsManager) GetProperties() (ret js.Promise[js.Array[ContactProperty]]) {
	bindings.CallContactsManagerGetProperties(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetProperties calls the method "ContactsManager.getProperties"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ContactsManager) TryGetProperties() (ret js.Promise[js.Array[ContactProperty]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryContactsManagerGetProperties(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSelect returns true if the method "ContactsManager.select" exists.
func (this ContactsManager) HasSelect() bool {
	return js.True == bindings.HasContactsManagerSelect(
		this.Ref(),
	)
}

// SelectFunc returns the method "ContactsManager.select".
func (this ContactsManager) SelectFunc() (fn js.Func[func(properties js.Array[ContactProperty], options ContactsSelectOptions) js.Promise[js.Array[ContactInfo]]]) {
	return fn.FromRef(
		bindings.ContactsManagerSelectFunc(
			this.Ref(),
		),
	)
}

// Select calls the method "ContactsManager.select".
func (this ContactsManager) Select(properties js.Array[ContactProperty], options ContactsSelectOptions) (ret js.Promise[js.Array[ContactInfo]]) {
	bindings.CallContactsManagerSelect(
		this.Ref(), js.Pointer(&ret),
		properties.Ref(),
		js.Pointer(&options),
	)

	return
}

// TrySelect calls the method "ContactsManager.select"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ContactsManager) TrySelect(properties js.Array[ContactProperty], options ContactsSelectOptions) (ret js.Promise[js.Array[ContactInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryContactsManagerSelect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		properties.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasSelect1 returns true if the method "ContactsManager.select" exists.
func (this ContactsManager) HasSelect1() bool {
	return js.True == bindings.HasContactsManagerSelect1(
		this.Ref(),
	)
}

// Select1Func returns the method "ContactsManager.select".
func (this ContactsManager) Select1Func() (fn js.Func[func(properties js.Array[ContactProperty]) js.Promise[js.Array[ContactInfo]]]) {
	return fn.FromRef(
		bindings.ContactsManagerSelect1Func(
			this.Ref(),
		),
	)
}

// Select1 calls the method "ContactsManager.select".
func (this ContactsManager) Select1(properties js.Array[ContactProperty]) (ret js.Promise[js.Array[ContactInfo]]) {
	bindings.CallContactsManagerSelect1(
		this.Ref(), js.Pointer(&ret),
		properties.Ref(),
	)

	return
}

// TrySelect1 calls the method "ContactsManager.select"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ContactsManager) TrySelect1(properties js.Array[ContactProperty]) (ret js.Promise[js.Array[ContactInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryContactsManagerSelect1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		properties.Ref(),
	)

	return
}

type KeyboardLayoutMap struct {
	ref js.Ref
}

func (this KeyboardLayoutMap) Once() KeyboardLayoutMap {
	this.Ref().Once()
	return this
}

func (this KeyboardLayoutMap) Ref() js.Ref {
	return this.ref
}

func (this KeyboardLayoutMap) FromRef(ref js.Ref) KeyboardLayoutMap {
	this.ref = ref
	return this
}

func (this KeyboardLayoutMap) Free() {
	this.Ref().Free()
}

type Keyboard struct {
	EventTarget
}

func (this Keyboard) Once() Keyboard {
	this.Ref().Once()
	return this
}

func (this Keyboard) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this Keyboard) FromRef(ref js.Ref) Keyboard {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this Keyboard) Free() {
	this.Ref().Free()
}

// HasLock returns true if the method "Keyboard.lock" exists.
func (this Keyboard) HasLock() bool {
	return js.True == bindings.HasKeyboardLock(
		this.Ref(),
	)
}

// LockFunc returns the method "Keyboard.lock".
func (this Keyboard) LockFunc() (fn js.Func[func(keyCodes js.Array[js.String]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.KeyboardLockFunc(
			this.Ref(),
		),
	)
}

// Lock calls the method "Keyboard.lock".
func (this Keyboard) Lock(keyCodes js.Array[js.String]) (ret js.Promise[js.Void]) {
	bindings.CallKeyboardLock(
		this.Ref(), js.Pointer(&ret),
		keyCodes.Ref(),
	)

	return
}

// TryLock calls the method "Keyboard.lock"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Keyboard) TryLock(keyCodes js.Array[js.String]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardLock(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		keyCodes.Ref(),
	)

	return
}

// HasLock1 returns true if the method "Keyboard.lock" exists.
func (this Keyboard) HasLock1() bool {
	return js.True == bindings.HasKeyboardLock1(
		this.Ref(),
	)
}

// Lock1Func returns the method "Keyboard.lock".
func (this Keyboard) Lock1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.KeyboardLock1Func(
			this.Ref(),
		),
	)
}

// Lock1 calls the method "Keyboard.lock".
func (this Keyboard) Lock1() (ret js.Promise[js.Void]) {
	bindings.CallKeyboardLock1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryLock1 calls the method "Keyboard.lock"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Keyboard) TryLock1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardLock1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasUnlock returns true if the method "Keyboard.unlock" exists.
func (this Keyboard) HasUnlock() bool {
	return js.True == bindings.HasKeyboardUnlock(
		this.Ref(),
	)
}

// UnlockFunc returns the method "Keyboard.unlock".
func (this Keyboard) UnlockFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.KeyboardUnlockFunc(
			this.Ref(),
		),
	)
}

// Unlock calls the method "Keyboard.unlock".
func (this Keyboard) Unlock() (ret js.Void) {
	bindings.CallKeyboardUnlock(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryUnlock calls the method "Keyboard.unlock"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Keyboard) TryUnlock() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardUnlock(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetLayoutMap returns true if the method "Keyboard.getLayoutMap" exists.
func (this Keyboard) HasGetLayoutMap() bool {
	return js.True == bindings.HasKeyboardGetLayoutMap(
		this.Ref(),
	)
}

// GetLayoutMapFunc returns the method "Keyboard.getLayoutMap".
func (this Keyboard) GetLayoutMapFunc() (fn js.Func[func() js.Promise[KeyboardLayoutMap]]) {
	return fn.FromRef(
		bindings.KeyboardGetLayoutMapFunc(
			this.Ref(),
		),
	)
}

// GetLayoutMap calls the method "Keyboard.getLayoutMap".
func (this Keyboard) GetLayoutMap() (ret js.Promise[KeyboardLayoutMap]) {
	bindings.CallKeyboardGetLayoutMap(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetLayoutMap calls the method "Keyboard.getLayoutMap"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Keyboard) TryGetLayoutMap() (ret js.Promise[KeyboardLayoutMap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardGetLayoutMap(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type MediaSessionAction uint32

const (
	_ MediaSessionAction = iota

	MediaSessionAction_PLAY
	MediaSessionAction_PAUSE
	MediaSessionAction_SEEKBACKWARD
	MediaSessionAction_SEEKFORWARD
	MediaSessionAction_PREVIOUSTRACK
	MediaSessionAction_NEXTTRACK
	MediaSessionAction_SKIPAD
	MediaSessionAction_STOP
	MediaSessionAction_SEEKTO
	MediaSessionAction_TOGGLEMICROPHONE
	MediaSessionAction_TOGGLECAMERA
	MediaSessionAction_HANGUP
	MediaSessionAction_PREVIOUSSLIDE
	MediaSessionAction_NEXTSLIDE
)

func (MediaSessionAction) FromRef(str js.Ref) MediaSessionAction {
	return MediaSessionAction(bindings.ConstOfMediaSessionAction(str))
}

func (x MediaSessionAction) String() (string, bool) {
	switch x {
	case MediaSessionAction_PLAY:
		return "play", true
	case MediaSessionAction_PAUSE:
		return "pause", true
	case MediaSessionAction_SEEKBACKWARD:
		return "seekbackward", true
	case MediaSessionAction_SEEKFORWARD:
		return "seekforward", true
	case MediaSessionAction_PREVIOUSTRACK:
		return "previoustrack", true
	case MediaSessionAction_NEXTTRACK:
		return "nexttrack", true
	case MediaSessionAction_SKIPAD:
		return "skipad", true
	case MediaSessionAction_STOP:
		return "stop", true
	case MediaSessionAction_SEEKTO:
		return "seekto", true
	case MediaSessionAction_TOGGLEMICROPHONE:
		return "togglemicrophone", true
	case MediaSessionAction_TOGGLECAMERA:
		return "togglecamera", true
	case MediaSessionAction_HANGUP:
		return "hangup", true
	case MediaSessionAction_PREVIOUSSLIDE:
		return "previousslide", true
	case MediaSessionAction_NEXTSLIDE:
		return "nextslide", true
	default:
		return "", false
	}
}

type MediaSessionActionHandlerFunc func(this js.Ref, details MediaSessionActionDetails) js.Ref

func (fn MediaSessionActionHandlerFunc) Register() js.Func[func(details MediaSessionActionDetails)] {
	return js.RegisterCallback[func(details MediaSessionActionDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn MediaSessionActionHandlerFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		MediaSessionActionDetails{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type MediaSessionActionHandler[T any] struct {
	Fn  func(arg T, this js.Ref, details MediaSessionActionDetails) js.Ref
	Arg T
}

func (cb *MediaSessionActionHandler[T]) Register() js.Func[func(details MediaSessionActionDetails)] {
	return js.RegisterCallback[func(details MediaSessionActionDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *MediaSessionActionHandler[T]) DispatchCallback(
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

		MediaSessionActionDetails{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type MediaSessionActionDetails struct {
	// Action is "MediaSessionActionDetails.action"
	//
	// Required
	Action MediaSessionAction
	// SeekOffset is "MediaSessionActionDetails.seekOffset"
	//
	// Optional
	//
	// NOTE: FFI_USE_SeekOffset MUST be set to true to make this field effective.
	SeekOffset float64
	// SeekTime is "MediaSessionActionDetails.seekTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_SeekTime MUST be set to true to make this field effective.
	SeekTime float64
	// FastSeek is "MediaSessionActionDetails.fastSeek"
	//
	// Optional
	//
	// NOTE: FFI_USE_FastSeek MUST be set to true to make this field effective.
	FastSeek bool

	FFI_USE_SeekOffset bool // for SeekOffset.
	FFI_USE_SeekTime   bool // for SeekTime.
	FFI_USE_FastSeek   bool // for FastSeek.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaSessionActionDetails with all fields set.
func (p MediaSessionActionDetails) FromRef(ref js.Ref) MediaSessionActionDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaSessionActionDetails in the application heap.
func (p MediaSessionActionDetails) New() js.Ref {
	return bindings.MediaSessionActionDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaSessionActionDetails) UpdateFrom(ref js.Ref) {
	bindings.MediaSessionActionDetailsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaSessionActionDetails) Update(ref js.Ref) {
	bindings.MediaSessionActionDetailsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaPositionState struct {
	// Duration is "MediaPositionState.duration"
	//
	// Optional
	//
	// NOTE: FFI_USE_Duration MUST be set to true to make this field effective.
	Duration float64
	// PlaybackRate is "MediaPositionState.playbackRate"
	//
	// Optional
	//
	// NOTE: FFI_USE_PlaybackRate MUST be set to true to make this field effective.
	PlaybackRate float64
	// Position is "MediaPositionState.position"
	//
	// Optional
	//
	// NOTE: FFI_USE_Position MUST be set to true to make this field effective.
	Position float64

	FFI_USE_Duration     bool // for Duration.
	FFI_USE_PlaybackRate bool // for PlaybackRate.
	FFI_USE_Position     bool // for Position.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaPositionState with all fields set.
func (p MediaPositionState) FromRef(ref js.Ref) MediaPositionState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaPositionState in the application heap.
func (p MediaPositionState) New() js.Ref {
	return bindings.MediaPositionStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaPositionState) UpdateFrom(ref js.Ref) {
	bindings.MediaPositionStateJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaPositionState) Update(ref js.Ref) {
	bindings.MediaPositionStateJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaImage struct {
	// Src is "MediaImage.src"
	//
	// Required
	Src js.String
	// Sizes is "MediaImage.sizes"
	//
	// Optional, defaults to "".
	Sizes js.String
	// Type is "MediaImage.type"
	//
	// Optional, defaults to "".
	Type js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaImage with all fields set.
func (p MediaImage) FromRef(ref js.Ref) MediaImage {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaImage in the application heap.
func (p MediaImage) New() js.Ref {
	return bindings.MediaImageJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaImage) UpdateFrom(ref js.Ref) {
	bindings.MediaImageJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaImage) Update(ref js.Ref) {
	bindings.MediaImageJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaMetadataInit struct {
	// Title is "MediaMetadataInit.title"
	//
	// Optional, defaults to "".
	Title js.String
	// Artist is "MediaMetadataInit.artist"
	//
	// Optional, defaults to "".
	Artist js.String
	// Album is "MediaMetadataInit.album"
	//
	// Optional, defaults to "".
	Album js.String
	// Artwork is "MediaMetadataInit.artwork"
	//
	// Optional, defaults to [].
	Artwork js.Array[MediaImage]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaMetadataInit with all fields set.
func (p MediaMetadataInit) FromRef(ref js.Ref) MediaMetadataInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaMetadataInit in the application heap.
func (p MediaMetadataInit) New() js.Ref {
	return bindings.MediaMetadataInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaMetadataInit) UpdateFrom(ref js.Ref) {
	bindings.MediaMetadataInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaMetadataInit) Update(ref js.Ref) {
	bindings.MediaMetadataInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewMediaMetadata(init MediaMetadataInit) (ret MediaMetadata) {
	ret.ref = bindings.NewMediaMetadataByMediaMetadata(
		js.Pointer(&init))
	return
}

func NewMediaMetadataByMediaMetadata1() (ret MediaMetadata) {
	ret.ref = bindings.NewMediaMetadataByMediaMetadata1()
	return
}

type MediaMetadata struct {
	ref js.Ref
}

func (this MediaMetadata) Once() MediaMetadata {
	this.Ref().Once()
	return this
}

func (this MediaMetadata) Ref() js.Ref {
	return this.ref
}

func (this MediaMetadata) FromRef(ref js.Ref) MediaMetadata {
	this.ref = ref
	return this
}

func (this MediaMetadata) Free() {
	this.Ref().Free()
}

// Title returns the value of property "MediaMetadata.title".
//
// It returns ok=false if there is no such property.
func (this MediaMetadata) Title() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaMetadataTitle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTitle sets the value of property "MediaMetadata.title" to val.
//
// It returns false if the property cannot be set.
func (this MediaMetadata) SetTitle(val js.String) bool {
	return js.True == bindings.SetMediaMetadataTitle(
		this.Ref(),
		val.Ref(),
	)
}

// Artist returns the value of property "MediaMetadata.artist".
//
// It returns ok=false if there is no such property.
func (this MediaMetadata) Artist() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaMetadataArtist(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetArtist sets the value of property "MediaMetadata.artist" to val.
//
// It returns false if the property cannot be set.
func (this MediaMetadata) SetArtist(val js.String) bool {
	return js.True == bindings.SetMediaMetadataArtist(
		this.Ref(),
		val.Ref(),
	)
}

// Album returns the value of property "MediaMetadata.album".
//
// It returns ok=false if there is no such property.
func (this MediaMetadata) Album() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaMetadataAlbum(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlbum sets the value of property "MediaMetadata.album" to val.
//
// It returns false if the property cannot be set.
func (this MediaMetadata) SetAlbum(val js.String) bool {
	return js.True == bindings.SetMediaMetadataAlbum(
		this.Ref(),
		val.Ref(),
	)
}

// Artwork returns the value of property "MediaMetadata.artwork".
//
// It returns ok=false if there is no such property.
func (this MediaMetadata) Artwork() (ret js.FrozenArray[MediaImage], ok bool) {
	ok = js.True == bindings.GetMediaMetadataArtwork(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetArtwork sets the value of property "MediaMetadata.artwork" to val.
//
// It returns false if the property cannot be set.
func (this MediaMetadata) SetArtwork(val js.FrozenArray[MediaImage]) bool {
	return js.True == bindings.SetMediaMetadataArtwork(
		this.Ref(),
		val.Ref(),
	)
}

type MediaSessionPlaybackState uint32

const (
	_ MediaSessionPlaybackState = iota

	MediaSessionPlaybackState_NONE
	MediaSessionPlaybackState_PAUSED
	MediaSessionPlaybackState_PLAYING
)

func (MediaSessionPlaybackState) FromRef(str js.Ref) MediaSessionPlaybackState {
	return MediaSessionPlaybackState(bindings.ConstOfMediaSessionPlaybackState(str))
}

func (x MediaSessionPlaybackState) String() (string, bool) {
	switch x {
	case MediaSessionPlaybackState_NONE:
		return "none", true
	case MediaSessionPlaybackState_PAUSED:
		return "paused", true
	case MediaSessionPlaybackState_PLAYING:
		return "playing", true
	default:
		return "", false
	}
}

type MediaSession struct {
	ref js.Ref
}

func (this MediaSession) Once() MediaSession {
	this.Ref().Once()
	return this
}

func (this MediaSession) Ref() js.Ref {
	return this.ref
}

func (this MediaSession) FromRef(ref js.Ref) MediaSession {
	this.ref = ref
	return this
}

func (this MediaSession) Free() {
	this.Ref().Free()
}

// Metadata returns the value of property "MediaSession.metadata".
//
// It returns ok=false if there is no such property.
func (this MediaSession) Metadata() (ret MediaMetadata, ok bool) {
	ok = js.True == bindings.GetMediaSessionMetadata(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMetadata sets the value of property "MediaSession.metadata" to val.
//
// It returns false if the property cannot be set.
func (this MediaSession) SetMetadata(val MediaMetadata) bool {
	return js.True == bindings.SetMediaSessionMetadata(
		this.Ref(),
		val.Ref(),
	)
}

// PlaybackState returns the value of property "MediaSession.playbackState".
//
// It returns ok=false if there is no such property.
func (this MediaSession) PlaybackState() (ret MediaSessionPlaybackState, ok bool) {
	ok = js.True == bindings.GetMediaSessionPlaybackState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPlaybackState sets the value of property "MediaSession.playbackState" to val.
//
// It returns false if the property cannot be set.
func (this MediaSession) SetPlaybackState(val MediaSessionPlaybackState) bool {
	return js.True == bindings.SetMediaSessionPlaybackState(
		this.Ref(),
		uint32(val),
	)
}

// HasSetActionHandler returns true if the method "MediaSession.setActionHandler" exists.
func (this MediaSession) HasSetActionHandler() bool {
	return js.True == bindings.HasMediaSessionSetActionHandler(
		this.Ref(),
	)
}

// SetActionHandlerFunc returns the method "MediaSession.setActionHandler".
func (this MediaSession) SetActionHandlerFunc() (fn js.Func[func(action MediaSessionAction, handler js.Func[func(details MediaSessionActionDetails)])]) {
	return fn.FromRef(
		bindings.MediaSessionSetActionHandlerFunc(
			this.Ref(),
		),
	)
}

// SetActionHandler calls the method "MediaSession.setActionHandler".
func (this MediaSession) SetActionHandler(action MediaSessionAction, handler js.Func[func(details MediaSessionActionDetails)]) (ret js.Void) {
	bindings.CallMediaSessionSetActionHandler(
		this.Ref(), js.Pointer(&ret),
		uint32(action),
		handler.Ref(),
	)

	return
}

// TrySetActionHandler calls the method "MediaSession.setActionHandler"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaSession) TrySetActionHandler(action MediaSessionAction, handler js.Func[func(details MediaSessionActionDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaSessionSetActionHandler(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(action),
		handler.Ref(),
	)

	return
}

// HasSetPositionState returns true if the method "MediaSession.setPositionState" exists.
func (this MediaSession) HasSetPositionState() bool {
	return js.True == bindings.HasMediaSessionSetPositionState(
		this.Ref(),
	)
}

// SetPositionStateFunc returns the method "MediaSession.setPositionState".
func (this MediaSession) SetPositionStateFunc() (fn js.Func[func(state MediaPositionState)]) {
	return fn.FromRef(
		bindings.MediaSessionSetPositionStateFunc(
			this.Ref(),
		),
	)
}

// SetPositionState calls the method "MediaSession.setPositionState".
func (this MediaSession) SetPositionState(state MediaPositionState) (ret js.Void) {
	bindings.CallMediaSessionSetPositionState(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&state),
	)

	return
}

// TrySetPositionState calls the method "MediaSession.setPositionState"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaSession) TrySetPositionState(state MediaPositionState) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaSessionSetPositionState(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&state),
	)

	return
}

// HasSetPositionState1 returns true if the method "MediaSession.setPositionState" exists.
func (this MediaSession) HasSetPositionState1() bool {
	return js.True == bindings.HasMediaSessionSetPositionState1(
		this.Ref(),
	)
}

// SetPositionState1Func returns the method "MediaSession.setPositionState".
func (this MediaSession) SetPositionState1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaSessionSetPositionState1Func(
			this.Ref(),
		),
	)
}

// SetPositionState1 calls the method "MediaSession.setPositionState".
func (this MediaSession) SetPositionState1() (ret js.Void) {
	bindings.CallMediaSessionSetPositionState1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySetPositionState1 calls the method "MediaSession.setPositionState"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaSession) TrySetPositionState1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaSessionSetPositionState1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetMicrophoneActive returns true if the method "MediaSession.setMicrophoneActive" exists.
func (this MediaSession) HasSetMicrophoneActive() bool {
	return js.True == bindings.HasMediaSessionSetMicrophoneActive(
		this.Ref(),
	)
}

// SetMicrophoneActiveFunc returns the method "MediaSession.setMicrophoneActive".
func (this MediaSession) SetMicrophoneActiveFunc() (fn js.Func[func(active bool)]) {
	return fn.FromRef(
		bindings.MediaSessionSetMicrophoneActiveFunc(
			this.Ref(),
		),
	)
}

// SetMicrophoneActive calls the method "MediaSession.setMicrophoneActive".
func (this MediaSession) SetMicrophoneActive(active bool) (ret js.Void) {
	bindings.CallMediaSessionSetMicrophoneActive(
		this.Ref(), js.Pointer(&ret),
		js.Bool(bool(active)),
	)

	return
}

// TrySetMicrophoneActive calls the method "MediaSession.setMicrophoneActive"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaSession) TrySetMicrophoneActive(active bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaSessionSetMicrophoneActive(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(active)),
	)

	return
}

// HasSetCameraActive returns true if the method "MediaSession.setCameraActive" exists.
func (this MediaSession) HasSetCameraActive() bool {
	return js.True == bindings.HasMediaSessionSetCameraActive(
		this.Ref(),
	)
}

// SetCameraActiveFunc returns the method "MediaSession.setCameraActive".
func (this MediaSession) SetCameraActiveFunc() (fn js.Func[func(active bool)]) {
	return fn.FromRef(
		bindings.MediaSessionSetCameraActiveFunc(
			this.Ref(),
		),
	)
}

// SetCameraActive calls the method "MediaSession.setCameraActive".
func (this MediaSession) SetCameraActive(active bool) (ret js.Void) {
	bindings.CallMediaSessionSetCameraActive(
		this.Ref(), js.Pointer(&ret),
		js.Bool(bool(active)),
	)

	return
}

// TrySetCameraActive calls the method "MediaSession.setCameraActive"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaSession) TrySetCameraActive(active bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaSessionSetCameraActive(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(active)),
	)

	return
}

type DevicePostureType uint32

const (
	_ DevicePostureType = iota

	DevicePostureType_CONTINUOUS
	DevicePostureType_FOLDED
)

func (DevicePostureType) FromRef(str js.Ref) DevicePostureType {
	return DevicePostureType(bindings.ConstOfDevicePostureType(str))
}

func (x DevicePostureType) String() (string, bool) {
	switch x {
	case DevicePostureType_CONTINUOUS:
		return "continuous", true
	case DevicePostureType_FOLDED:
		return "folded", true
	default:
		return "", false
	}
}

type DevicePosture struct {
	EventTarget
}

func (this DevicePosture) Once() DevicePosture {
	this.Ref().Once()
	return this
}

func (this DevicePosture) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this DevicePosture) FromRef(ref js.Ref) DevicePosture {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this DevicePosture) Free() {
	this.Ref().Free()
}

// Type returns the value of property "DevicePosture.type".
//
// It returns ok=false if there is no such property.
func (this DevicePosture) Type() (ret DevicePostureType, ok bool) {
	ok = js.True == bindings.GetDevicePostureType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}
