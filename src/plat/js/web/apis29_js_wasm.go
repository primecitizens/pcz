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
// The returned bool will be false if there is no such property.
func (this SerialPort) Readable() (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.GetSerialPortReadable(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadableStream{}.FromRef(_ret), _ok
}

// Writable returns the value of property "SerialPort.writable".
//
// The returned bool will be false if there is no such property.
func (this SerialPort) Writable() (WritableStream, bool) {
	var _ok bool
	_ret := bindings.GetSerialPortWritable(
		this.Ref(), js.Pointer(&_ok),
	)
	return WritableStream{}.FromRef(_ret), _ok
}

// GetInfo calls the method "SerialPort.getInfo".
//
// The returned bool will be false if there is no such method.
func (this SerialPort) GetInfo() (SerialPortInfo, bool) {
	var _ret SerialPortInfo
	_ok := js.True == bindings.CallSerialPortGetInfo(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetInfoFunc returns the method "SerialPort.getInfo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SerialPort) GetInfoFunc() (fn js.Func[func() SerialPortInfo]) {
	return fn.FromRef(
		bindings.SerialPortGetInfoFunc(
			this.Ref(),
		),
	)
}

// Open calls the method "SerialPort.open".
//
// The returned bool will be false if there is no such method.
func (this SerialPort) Open(options SerialOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallSerialPortOpen(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// OpenFunc returns the method "SerialPort.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SerialPort) OpenFunc() (fn js.Func[func(options SerialOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SerialPortOpenFunc(
			this.Ref(),
		),
	)
}

// SetSignals calls the method "SerialPort.setSignals".
//
// The returned bool will be false if there is no such method.
func (this SerialPort) SetSignals(signals SerialOutputSignals) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallSerialPortSetSignals(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&signals),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetSignalsFunc returns the method "SerialPort.setSignals".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SerialPort) SetSignalsFunc() (fn js.Func[func(signals SerialOutputSignals) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SerialPortSetSignalsFunc(
			this.Ref(),
		),
	)
}

// SetSignals1 calls the method "SerialPort.setSignals".
//
// The returned bool will be false if there is no such method.
func (this SerialPort) SetSignals1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallSerialPortSetSignals1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetSignals1Func returns the method "SerialPort.setSignals".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SerialPort) SetSignals1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SerialPortSetSignals1Func(
			this.Ref(),
		),
	)
}

// GetSignals calls the method "SerialPort.getSignals".
//
// The returned bool will be false if there is no such method.
func (this SerialPort) GetSignals() (js.Promise[SerialInputSignals], bool) {
	var _ok bool
	_ret := bindings.CallSerialPortGetSignals(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[SerialInputSignals]{}.FromRef(_ret), _ok
}

// GetSignalsFunc returns the method "SerialPort.getSignals".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SerialPort) GetSignalsFunc() (fn js.Func[func() js.Promise[SerialInputSignals]]) {
	return fn.FromRef(
		bindings.SerialPortGetSignalsFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "SerialPort.close".
//
// The returned bool will be false if there is no such method.
func (this SerialPort) Close() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallSerialPortClose(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CloseFunc returns the method "SerialPort.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SerialPort) CloseFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SerialPortCloseFunc(
			this.Ref(),
		),
	)
}

// Forget calls the method "SerialPort.forget".
//
// The returned bool will be false if there is no such method.
func (this SerialPort) Forget() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallSerialPortForget(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ForgetFunc returns the method "SerialPort.forget".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SerialPort) ForgetFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SerialPortForgetFunc(
			this.Ref(),
		),
	)
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

// GetPorts calls the method "Serial.getPorts".
//
// The returned bool will be false if there is no such method.
func (this Serial) GetPorts() (js.Promise[js.Array[SerialPort]], bool) {
	var _ok bool
	_ret := bindings.CallSerialGetPorts(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[SerialPort]]{}.FromRef(_ret), _ok
}

// GetPortsFunc returns the method "Serial.getPorts".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Serial) GetPortsFunc() (fn js.Func[func() js.Promise[js.Array[SerialPort]]]) {
	return fn.FromRef(
		bindings.SerialGetPortsFunc(
			this.Ref(),
		),
	)
}

// RequestPort calls the method "Serial.requestPort".
//
// The returned bool will be false if there is no such method.
func (this Serial) RequestPort(options SerialPortRequestOptions) (js.Promise[SerialPort], bool) {
	var _ok bool
	_ret := bindings.CallSerialRequestPort(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[SerialPort]{}.FromRef(_ret), _ok
}

// RequestPortFunc returns the method "Serial.requestPort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Serial) RequestPortFunc() (fn js.Func[func(options SerialPortRequestOptions) js.Promise[SerialPort]]) {
	return fn.FromRef(
		bindings.SerialRequestPortFunc(
			this.Ref(),
		),
	)
}

// RequestPort1 calls the method "Serial.requestPort".
//
// The returned bool will be false if there is no such method.
func (this Serial) RequestPort1() (js.Promise[SerialPort], bool) {
	var _ok bool
	_ret := bindings.CallSerialRequestPort1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[SerialPort]{}.FromRef(_ret), _ok
}

// RequestPort1Func returns the method "Serial.requestPort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Serial) RequestPort1Func() (fn js.Func[func() js.Promise[SerialPort]]) {
	return fn.FromRef(
		bindings.SerialRequestPort1Func(
			this.Ref(),
		),
	)
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
	Audio KeySystemTrackConfiguration
	// Video is "MediaCapabilitiesKeySystemConfiguration.video"
	//
	// Optional
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
	KeySystemConfiguration MediaCapabilitiesKeySystemConfiguration
	// Video is "MediaDecodingConfiguration.video"
	//
	// Optional
	Video VideoConfiguration
	// Audio is "MediaDecodingConfiguration.audio"
	//
	// Optional
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
	Video VideoConfiguration
	// Audio is "MediaEncodingConfiguration.audio"
	//
	// Optional
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

// DecodingInfo calls the method "MediaCapabilities.decodingInfo".
//
// The returned bool will be false if there is no such method.
func (this MediaCapabilities) DecodingInfo(configuration MediaDecodingConfiguration) (js.Promise[MediaCapabilitiesDecodingInfo], bool) {
	var _ok bool
	_ret := bindings.CallMediaCapabilitiesDecodingInfo(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&configuration),
	)

	return js.Promise[MediaCapabilitiesDecodingInfo]{}.FromRef(_ret), _ok
}

// DecodingInfoFunc returns the method "MediaCapabilities.decodingInfo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaCapabilities) DecodingInfoFunc() (fn js.Func[func(configuration MediaDecodingConfiguration) js.Promise[MediaCapabilitiesDecodingInfo]]) {
	return fn.FromRef(
		bindings.MediaCapabilitiesDecodingInfoFunc(
			this.Ref(),
		),
	)
}

// EncodingInfo calls the method "MediaCapabilities.encodingInfo".
//
// The returned bool will be false if there is no such method.
func (this MediaCapabilities) EncodingInfo(configuration MediaEncodingConfiguration) (js.Promise[MediaCapabilitiesEncodingInfo], bool) {
	var _ok bool
	_ret := bindings.CallMediaCapabilitiesEncodingInfo(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&configuration),
	)

	return js.Promise[MediaCapabilitiesEncodingInfo]{}.FromRef(_ret), _ok
}

// EncodingInfoFunc returns the method "MediaCapabilities.encodingInfo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaCapabilities) EncodingInfoFunc() (fn js.Func[func(configuration MediaEncodingConfiguration) js.Promise[MediaCapabilitiesEncodingInfo]]) {
	return fn.FromRef(
		bindings.MediaCapabilitiesEncodingInfoFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this UserActivation) HasBeenActive() (bool, bool) {
	var _ok bool
	_ret := bindings.GetUserActivationHasBeenActive(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// IsActive returns the value of property "UserActivation.isActive".
//
// The returned bool will be false if there is no such property.
func (this UserActivation) IsActive() (bool, bool) {
	var _ok bool
	_ret := bindings.GetUserActivationIsActive(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
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
// The returned bool will be false if there is no such property.
func (this PermissionStatus) State() (PermissionState, bool) {
	var _ok bool
	_ret := bindings.GetPermissionStatusState(
		this.Ref(), js.Pointer(&_ok),
	)
	return PermissionState(_ret), _ok
}

// Name returns the value of property "PermissionStatus.name".
//
// The returned bool will be false if there is no such property.
func (this PermissionStatus) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPermissionStatusName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
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

// Query calls the method "Permissions.query".
//
// The returned bool will be false if there is no such method.
func (this Permissions) Query(permissionDesc js.Object) (js.Promise[PermissionStatus], bool) {
	var _ok bool
	_ret := bindings.CallPermissionsQuery(
		this.Ref(), js.Pointer(&_ok),
		permissionDesc.Ref(),
	)

	return js.Promise[PermissionStatus]{}.FromRef(_ret), _ok
}

// QueryFunc returns the method "Permissions.query".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Permissions) QueryFunc() (fn js.Func[func(permissionDesc js.Object) js.Promise[PermissionStatus]]) {
	return fn.FromRef(
		bindings.PermissionsQueryFunc(
			this.Ref(),
		),
	)
}

// Revoke calls the method "Permissions.revoke".
//
// The returned bool will be false if there is no such method.
func (this Permissions) Revoke(permissionDesc js.Object) (js.Promise[PermissionStatus], bool) {
	var _ok bool
	_ret := bindings.CallPermissionsRevoke(
		this.Ref(), js.Pointer(&_ok),
		permissionDesc.Ref(),
	)

	return js.Promise[PermissionStatus]{}.FromRef(_ret), _ok
}

// RevokeFunc returns the method "Permissions.revoke".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Permissions) RevokeFunc() (fn js.Func[func(permissionDesc js.Object) js.Promise[PermissionStatus]]) {
	return fn.FromRef(
		bindings.PermissionsRevokeFunc(
			this.Ref(),
		),
	)
}

// Request calls the method "Permissions.request".
//
// The returned bool will be false if there is no such method.
func (this Permissions) Request(permissionDesc js.Object) (js.Promise[PermissionStatus], bool) {
	var _ok bool
	_ret := bindings.CallPermissionsRequest(
		this.Ref(), js.Pointer(&_ok),
		permissionDesc.Ref(),
	)

	return js.Promise[PermissionStatus]{}.FromRef(_ret), _ok
}

// RequestFunc returns the method "Permissions.request".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Permissions) RequestFunc() (fn js.Func[func(permissionDesc js.Object) js.Promise[PermissionStatus]]) {
	return fn.FromRef(
		bindings.PermissionsRequestFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this ContactAddress) City() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetContactAddressCity(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Country returns the value of property "ContactAddress.country".
//
// The returned bool will be false if there is no such property.
func (this ContactAddress) Country() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetContactAddressCountry(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DependentLocality returns the value of property "ContactAddress.dependentLocality".
//
// The returned bool will be false if there is no such property.
func (this ContactAddress) DependentLocality() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetContactAddressDependentLocality(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Organization returns the value of property "ContactAddress.organization".
//
// The returned bool will be false if there is no such property.
func (this ContactAddress) Organization() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetContactAddressOrganization(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Phone returns the value of property "ContactAddress.phone".
//
// The returned bool will be false if there is no such property.
func (this ContactAddress) Phone() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetContactAddressPhone(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// PostalCode returns the value of property "ContactAddress.postalCode".
//
// The returned bool will be false if there is no such property.
func (this ContactAddress) PostalCode() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetContactAddressPostalCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Recipient returns the value of property "ContactAddress.recipient".
//
// The returned bool will be false if there is no such property.
func (this ContactAddress) Recipient() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetContactAddressRecipient(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Region returns the value of property "ContactAddress.region".
//
// The returned bool will be false if there is no such property.
func (this ContactAddress) Region() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetContactAddressRegion(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SortingCode returns the value of property "ContactAddress.sortingCode".
//
// The returned bool will be false if there is no such property.
func (this ContactAddress) SortingCode() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetContactAddressSortingCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AddressLine returns the value of property "ContactAddress.addressLine".
//
// The returned bool will be false if there is no such property.
func (this ContactAddress) AddressLine() (js.FrozenArray[js.String], bool) {
	var _ok bool
	_ret := bindings.GetContactAddressAddressLine(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[js.String]{}.FromRef(_ret), _ok
}

// ToJSON calls the method "ContactAddress.toJSON".
//
// The returned bool will be false if there is no such method.
func (this ContactAddress) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallContactAddressToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "ContactAddress.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ContactAddress) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.ContactAddressToJSONFunc(
			this.Ref(),
		),
	)
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

// GetProperties calls the method "ContactsManager.getProperties".
//
// The returned bool will be false if there is no such method.
func (this ContactsManager) GetProperties() (js.Promise[js.Array[ContactProperty]], bool) {
	var _ok bool
	_ret := bindings.CallContactsManagerGetProperties(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[ContactProperty]]{}.FromRef(_ret), _ok
}

// GetPropertiesFunc returns the method "ContactsManager.getProperties".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ContactsManager) GetPropertiesFunc() (fn js.Func[func() js.Promise[js.Array[ContactProperty]]]) {
	return fn.FromRef(
		bindings.ContactsManagerGetPropertiesFunc(
			this.Ref(),
		),
	)
}

// Select calls the method "ContactsManager.select".
//
// The returned bool will be false if there is no such method.
func (this ContactsManager) Select(properties js.Array[ContactProperty], options ContactsSelectOptions) (js.Promise[js.Array[ContactInfo]], bool) {
	var _ok bool
	_ret := bindings.CallContactsManagerSelect(
		this.Ref(), js.Pointer(&_ok),
		properties.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Array[ContactInfo]]{}.FromRef(_ret), _ok
}

// SelectFunc returns the method "ContactsManager.select".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ContactsManager) SelectFunc() (fn js.Func[func(properties js.Array[ContactProperty], options ContactsSelectOptions) js.Promise[js.Array[ContactInfo]]]) {
	return fn.FromRef(
		bindings.ContactsManagerSelectFunc(
			this.Ref(),
		),
	)
}

// Select1 calls the method "ContactsManager.select".
//
// The returned bool will be false if there is no such method.
func (this ContactsManager) Select1(properties js.Array[ContactProperty]) (js.Promise[js.Array[ContactInfo]], bool) {
	var _ok bool
	_ret := bindings.CallContactsManagerSelect1(
		this.Ref(), js.Pointer(&_ok),
		properties.Ref(),
	)

	return js.Promise[js.Array[ContactInfo]]{}.FromRef(_ret), _ok
}

// Select1Func returns the method "ContactsManager.select".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ContactsManager) Select1Func() (fn js.Func[func(properties js.Array[ContactProperty]) js.Promise[js.Array[ContactInfo]]]) {
	return fn.FromRef(
		bindings.ContactsManagerSelect1Func(
			this.Ref(),
		),
	)
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

// Lock calls the method "Keyboard.lock".
//
// The returned bool will be false if there is no such method.
func (this Keyboard) Lock(keyCodes js.Array[js.String]) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallKeyboardLock(
		this.Ref(), js.Pointer(&_ok),
		keyCodes.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// LockFunc returns the method "Keyboard.lock".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Keyboard) LockFunc() (fn js.Func[func(keyCodes js.Array[js.String]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.KeyboardLockFunc(
			this.Ref(),
		),
	)
}

// Lock1 calls the method "Keyboard.lock".
//
// The returned bool will be false if there is no such method.
func (this Keyboard) Lock1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallKeyboardLock1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Lock1Func returns the method "Keyboard.lock".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Keyboard) Lock1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.KeyboardLock1Func(
			this.Ref(),
		),
	)
}

// Unlock calls the method "Keyboard.unlock".
//
// The returned bool will be false if there is no such method.
func (this Keyboard) Unlock() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallKeyboardUnlock(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UnlockFunc returns the method "Keyboard.unlock".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Keyboard) UnlockFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.KeyboardUnlockFunc(
			this.Ref(),
		),
	)
}

// GetLayoutMap calls the method "Keyboard.getLayoutMap".
//
// The returned bool will be false if there is no such method.
func (this Keyboard) GetLayoutMap() (js.Promise[KeyboardLayoutMap], bool) {
	var _ok bool
	_ret := bindings.CallKeyboardGetLayoutMap(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[KeyboardLayoutMap]{}.FromRef(_ret), _ok
}

// GetLayoutMapFunc returns the method "Keyboard.getLayoutMap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Keyboard) GetLayoutMapFunc() (fn js.Func[func() js.Promise[KeyboardLayoutMap]]) {
	return fn.FromRef(
		bindings.KeyboardGetLayoutMapFunc(
			this.Ref(),
		),
	)
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

func NewMediaMetadata(init MediaMetadataInit) MediaMetadata {
	return MediaMetadata{}.FromRef(
		bindings.NewMediaMetadataByMediaMetadata(
			js.Pointer(&init)),
	)
}

func NewMediaMetadataByMediaMetadata1() MediaMetadata {
	return MediaMetadata{}.FromRef(
		bindings.NewMediaMetadataByMediaMetadata1(),
	)
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
// The returned bool will be false if there is no such property.
func (this MediaMetadata) Title() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaMetadataTitle(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Title sets the value of property "MediaMetadata.title" to val.
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
// The returned bool will be false if there is no such property.
func (this MediaMetadata) Artist() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaMetadataArtist(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Artist sets the value of property "MediaMetadata.artist" to val.
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
// The returned bool will be false if there is no such property.
func (this MediaMetadata) Album() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaMetadataAlbum(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Album sets the value of property "MediaMetadata.album" to val.
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
// The returned bool will be false if there is no such property.
func (this MediaMetadata) Artwork() (js.FrozenArray[MediaImage], bool) {
	var _ok bool
	_ret := bindings.GetMediaMetadataArtwork(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[MediaImage]{}.FromRef(_ret), _ok
}

// Artwork sets the value of property "MediaMetadata.artwork" to val.
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
// The returned bool will be false if there is no such property.
func (this MediaSession) Metadata() (MediaMetadata, bool) {
	var _ok bool
	_ret := bindings.GetMediaSessionMetadata(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaMetadata{}.FromRef(_ret), _ok
}

// Metadata sets the value of property "MediaSession.metadata" to val.
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
// The returned bool will be false if there is no such property.
func (this MediaSession) PlaybackState() (MediaSessionPlaybackState, bool) {
	var _ok bool
	_ret := bindings.GetMediaSessionPlaybackState(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaSessionPlaybackState(_ret), _ok
}

// PlaybackState sets the value of property "MediaSession.playbackState" to val.
//
// It returns false if the property cannot be set.
func (this MediaSession) SetPlaybackState(val MediaSessionPlaybackState) bool {
	return js.True == bindings.SetMediaSessionPlaybackState(
		this.Ref(),
		uint32(val),
	)
}

// SetActionHandler calls the method "MediaSession.setActionHandler".
//
// The returned bool will be false if there is no such method.
func (this MediaSession) SetActionHandler(action MediaSessionAction, handler js.Func[func(details MediaSessionActionDetails)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaSessionSetActionHandler(
		this.Ref(), js.Pointer(&_ok),
		uint32(action),
		handler.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetActionHandlerFunc returns the method "MediaSession.setActionHandler".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaSession) SetActionHandlerFunc() (fn js.Func[func(action MediaSessionAction, handler js.Func[func(details MediaSessionActionDetails)])]) {
	return fn.FromRef(
		bindings.MediaSessionSetActionHandlerFunc(
			this.Ref(),
		),
	)
}

// SetPositionState calls the method "MediaSession.setPositionState".
//
// The returned bool will be false if there is no such method.
func (this MediaSession) SetPositionState(state MediaPositionState) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaSessionSetPositionState(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&state),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPositionStateFunc returns the method "MediaSession.setPositionState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaSession) SetPositionStateFunc() (fn js.Func[func(state MediaPositionState)]) {
	return fn.FromRef(
		bindings.MediaSessionSetPositionStateFunc(
			this.Ref(),
		),
	)
}

// SetPositionState1 calls the method "MediaSession.setPositionState".
//
// The returned bool will be false if there is no such method.
func (this MediaSession) SetPositionState1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaSessionSetPositionState1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPositionState1Func returns the method "MediaSession.setPositionState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaSession) SetPositionState1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaSessionSetPositionState1Func(
			this.Ref(),
		),
	)
}

// SetMicrophoneActive calls the method "MediaSession.setMicrophoneActive".
//
// The returned bool will be false if there is no such method.
func (this MediaSession) SetMicrophoneActive(active bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaSessionSetMicrophoneActive(
		this.Ref(), js.Pointer(&_ok),
		js.Bool(bool(active)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetMicrophoneActiveFunc returns the method "MediaSession.setMicrophoneActive".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaSession) SetMicrophoneActiveFunc() (fn js.Func[func(active bool)]) {
	return fn.FromRef(
		bindings.MediaSessionSetMicrophoneActiveFunc(
			this.Ref(),
		),
	)
}

// SetCameraActive calls the method "MediaSession.setCameraActive".
//
// The returned bool will be false if there is no such method.
func (this MediaSession) SetCameraActive(active bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaSessionSetCameraActive(
		this.Ref(), js.Pointer(&_ok),
		js.Bool(bool(active)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetCameraActiveFunc returns the method "MediaSession.setCameraActive".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaSession) SetCameraActiveFunc() (fn js.Func[func(active bool)]) {
	return fn.FromRef(
		bindings.MediaSessionSetCameraActiveFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this DevicePosture) Type() (DevicePostureType, bool) {
	var _ok bool
	_ret := bindings.GetDevicePostureType(
		this.Ref(), js.Pointer(&_ok),
	)
	return DevicePostureType(_ret), _ok
}
