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

type RTCRtpEncodingParameters struct {
	// Active is "RTCRtpEncodingParameters.active"
	//
	// Optional, defaults to true.
	Active bool
	// MaxBitrate is "RTCRtpEncodingParameters.maxBitrate"
	//
	// Optional
	MaxBitrate uint32
	// MaxFramerate is "RTCRtpEncodingParameters.maxFramerate"
	//
	// Optional
	MaxFramerate float64
	// ScaleResolutionDownBy is "RTCRtpEncodingParameters.scaleResolutionDownBy"
	//
	// Optional
	ScaleResolutionDownBy float64
	// Rid is "RTCRtpEncodingParameters.rid"
	//
	// Optional
	Rid js.String
	// Priority is "RTCRtpEncodingParameters.priority"
	//
	// Optional, defaults to "low".
	Priority RTCPriorityType
	// NetworkPriority is "RTCRtpEncodingParameters.networkPriority"
	//
	// Optional
	NetworkPriority RTCPriorityType
	// ScalabilityMode is "RTCRtpEncodingParameters.scalabilityMode"
	//
	// Optional
	ScalabilityMode js.String

	FFI_USE_Active                bool // for Active.
	FFI_USE_MaxBitrate            bool // for MaxBitrate.
	FFI_USE_MaxFramerate          bool // for MaxFramerate.
	FFI_USE_ScaleResolutionDownBy bool // for ScaleResolutionDownBy.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpEncodingParameters with all fields set.
func (p RTCRtpEncodingParameters) FromRef(ref js.Ref) RTCRtpEncodingParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCRtpEncodingParameters RTCRtpEncodingParameters [// RTCRtpEncodingParameters] [0x140009d43c0 0x140009d4500 0x140009d4640 0x140009d4780 0x140009d48c0 0x140009d4960 0x140009d4a00 0x140009d4aa0 0x140009d4460 0x140009d45a0 0x140009d46e0 0x140009d4820] 0x1400099ce58 {0 0}} in the application heap.
func (p RTCRtpEncodingParameters) New() js.Ref {
	return bindings.RTCRtpEncodingParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtpEncodingParameters) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpEncodingParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtpEncodingParameters) Update(ref js.Ref) {
	bindings.RTCRtpEncodingParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCRtpHeaderExtensionParameters struct {
	// Uri is "RTCRtpHeaderExtensionParameters.uri"
	//
	// Required
	Uri js.String
	// Id is "RTCRtpHeaderExtensionParameters.id"
	//
	// Required
	Id uint16
	// Encrypted is "RTCRtpHeaderExtensionParameters.encrypted"
	//
	// Optional, defaults to false.
	Encrypted bool

	FFI_USE_Encrypted bool // for Encrypted.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpHeaderExtensionParameters with all fields set.
func (p RTCRtpHeaderExtensionParameters) FromRef(ref js.Ref) RTCRtpHeaderExtensionParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCRtpHeaderExtensionParameters RTCRtpHeaderExtensionParameters [// RTCRtpHeaderExtensionParameters] [0x140009d4be0 0x140009d4c80 0x140009d4d20 0x140009d4dc0] 0x1400099ceb8 {0 0}} in the application heap.
func (p RTCRtpHeaderExtensionParameters) New() js.Ref {
	return bindings.RTCRtpHeaderExtensionParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtpHeaderExtensionParameters) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpHeaderExtensionParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtpHeaderExtensionParameters) Update(ref js.Ref) {
	bindings.RTCRtpHeaderExtensionParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCRtcpParameters struct {
	// Cname is "RTCRtcpParameters.cname"
	//
	// Optional
	Cname js.String
	// ReducedSize is "RTCRtcpParameters.reducedSize"
	//
	// Optional
	ReducedSize bool

	FFI_USE_ReducedSize bool // for ReducedSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtcpParameters with all fields set.
func (p RTCRtcpParameters) FromRef(ref js.Ref) RTCRtcpParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCRtcpParameters RTCRtcpParameters [// RTCRtcpParameters] [0x140009d4f00 0x140009d4fa0 0x140009d5040] 0x1400099cf48 {0 0}} in the application heap.
func (p RTCRtcpParameters) New() js.Ref {
	return bindings.RTCRtcpParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtcpParameters) UpdateFrom(ref js.Ref) {
	bindings.RTCRtcpParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtcpParameters) Update(ref js.Ref) {
	bindings.RTCRtcpParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCRtpCodecParameters struct {
	// PayloadType is "RTCRtpCodecParameters.payloadType"
	//
	// Required
	PayloadType uint8
	// MimeType is "RTCRtpCodecParameters.mimeType"
	//
	// Required
	MimeType js.String
	// ClockRate is "RTCRtpCodecParameters.clockRate"
	//
	// Required
	ClockRate uint32
	// Channels is "RTCRtpCodecParameters.channels"
	//
	// Optional
	Channels uint16
	// SdpFmtpLine is "RTCRtpCodecParameters.sdpFmtpLine"
	//
	// Optional
	SdpFmtpLine js.String

	FFI_USE_Channels bool // for Channels.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpCodecParameters with all fields set.
func (p RTCRtpCodecParameters) FromRef(ref js.Ref) RTCRtpCodecParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCRtpCodecParameters RTCRtpCodecParameters [// RTCRtpCodecParameters] [0x140009d5180 0x140009d5220 0x140009d52c0 0x140009d5360 0x140009d54a0 0x140009d5400] 0x1400099cf90 {0 0}} in the application heap.
func (p RTCRtpCodecParameters) New() js.Ref {
	return bindings.RTCRtpCodecParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtpCodecParameters) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpCodecParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtpCodecParameters) Update(ref js.Ref) {
	bindings.RTCRtpCodecParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCRtpSendParameters struct {
	// TransactionId is "RTCRtpSendParameters.transactionId"
	//
	// Required
	TransactionId js.String
	// Encodings is "RTCRtpSendParameters.encodings"
	//
	// Required
	Encodings js.Array[RTCRtpEncodingParameters]
	// HeaderExtensions is "RTCRtpSendParameters.headerExtensions"
	//
	// Required
	HeaderExtensions js.Array[RTCRtpHeaderExtensionParameters]
	// Rtcp is "RTCRtpSendParameters.rtcp"
	//
	// Required
	Rtcp RTCRtcpParameters
	// Codecs is "RTCRtpSendParameters.codecs"
	//
	// Required
	Codecs js.Array[RTCRtpCodecParameters]
	// DegradationPreference is "RTCRtpSendParameters.degradationPreference"
	//
	// Optional
	DegradationPreference RTCDegradationPreference

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpSendParameters with all fields set.
func (p RTCRtpSendParameters) FromRef(ref js.Ref) RTCRtpSendParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCRtpSendParameters RTCRtpSendParameters [// RTCRtpSendParameters] [0x140009d4320 0x140009d4b40 0x140009d4e60 0x140009d50e0 0x140009d5540 0x140009d55e0] 0x1400099ce40 {0 0}} in the application heap.
func (p RTCRtpSendParameters) New() js.Ref {
	return bindings.RTCRtpSendParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtpSendParameters) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpSendParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtpSendParameters) Update(ref js.Ref) {
	bindings.RTCRtpSendParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCSetParameterOptions struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCSetParameterOptions with all fields set.
func (p RTCSetParameterOptions) FromRef(ref js.Ref) RTCSetParameterOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCSetParameterOptions RTCSetParameterOptions [// RTCSetParameterOptions] [] 0x1400099cfc0 {0 0}} in the application heap.
func (p RTCSetParameterOptions) New() js.Ref {
	return bindings.RTCSetParameterOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCSetParameterOptions) UpdateFrom(ref js.Ref) {
	bindings.RTCSetParameterOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCSetParameterOptions) Update(ref js.Ref) {
	bindings.RTCSetParameterOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCStatsReport struct {
	ref js.Ref
}

func (this RTCStatsReport) Once() RTCStatsReport {
	this.Ref().Once()
	return this
}

func (this RTCStatsReport) Ref() js.Ref {
	return this.ref
}

func (this RTCStatsReport) FromRef(ref js.Ref) RTCStatsReport {
	this.ref = ref
	return this
}

func (this RTCStatsReport) Free() {
	this.Ref().Free()
}

type SFrameTransformRole uint32

const (
	_ SFrameTransformRole = iota

	SFrameTransformRole_ENCRYPT
	SFrameTransformRole_DECRYPT
)

func (SFrameTransformRole) FromRef(str js.Ref) SFrameTransformRole {
	return SFrameTransformRole(bindings.ConstOfSFrameTransformRole(str))
}

func (x SFrameTransformRole) String() (string, bool) {
	switch x {
	case SFrameTransformRole_ENCRYPT:
		return "encrypt", true
	case SFrameTransformRole_DECRYPT:
		return "decrypt", true
	default:
		return "", false
	}
}

type SFrameTransformOptions struct {
	// Role is "SFrameTransformOptions.role"
	//
	// Optional, defaults to "encrypt".
	Role SFrameTransformRole

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SFrameTransformOptions with all fields set.
func (p SFrameTransformOptions) FromRef(ref js.Ref) SFrameTransformOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 SFrameTransformOptions SFrameTransformOptions [// SFrameTransformOptions] [0x140009d5720] 0x1400099d020 {0 0}} in the application heap.
func (p SFrameTransformOptions) New() js.Ref {
	return bindings.SFrameTransformOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SFrameTransformOptions) UpdateFrom(ref js.Ref) {
	bindings.SFrameTransformOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SFrameTransformOptions) Update(ref js.Ref) {
	bindings.SFrameTransformOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewSFrameTransform(options SFrameTransformOptions) SFrameTransform {
	return SFrameTransform{}.FromRef(
		bindings.NewSFrameTransformBySFrameTransform(
			js.Pointer(&options)),
	)
}

func NewSFrameTransformBySFrameTransform1() SFrameTransform {
	return SFrameTransform{}.FromRef(
		bindings.NewSFrameTransformBySFrameTransform1(),
	)
}

type SFrameTransform struct {
	EventTarget
}

func (this SFrameTransform) Once() SFrameTransform {
	this.Ref().Once()
	return this
}

func (this SFrameTransform) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this SFrameTransform) FromRef(ref js.Ref) SFrameTransform {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this SFrameTransform) Free() {
	this.Ref().Free()
}

// Readable returns the value of property "SFrameTransform.readable".
//
// The returned bool will be false if there is no such property.
func (this SFrameTransform) Readable() (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.GetSFrameTransformReadable(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadableStream{}.FromRef(_ret), _ok
}

// Writable returns the value of property "SFrameTransform.writable".
//
// The returned bool will be false if there is no such property.
func (this SFrameTransform) Writable() (WritableStream, bool) {
	var _ok bool
	_ret := bindings.GetSFrameTransformWritable(
		this.Ref(), js.Pointer(&_ok),
	)
	return WritableStream{}.FromRef(_ret), _ok
}

// SetEncryptionKey calls the method "SFrameTransform.setEncryptionKey".
//
// The returned bool will be false if there is no such method.
func (this SFrameTransform) SetEncryptionKey(key CryptoKey, keyID CryptoKeyID) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallSFrameTransformSetEncryptionKey(
		this.Ref(), js.Pointer(&_ok),
		key.Ref(),
		keyID.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetEncryptionKeyFunc returns the method "SFrameTransform.setEncryptionKey".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SFrameTransform) SetEncryptionKeyFunc() (fn js.Func[func(key CryptoKey, keyID CryptoKeyID) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SFrameTransformSetEncryptionKeyFunc(
			this.Ref(),
		),
	)
}

// SetEncryptionKey1 calls the method "SFrameTransform.setEncryptionKey".
//
// The returned bool will be false if there is no such method.
func (this SFrameTransform) SetEncryptionKey1(key CryptoKey) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallSFrameTransformSetEncryptionKey1(
		this.Ref(), js.Pointer(&_ok),
		key.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetEncryptionKey1Func returns the method "SFrameTransform.setEncryptionKey".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SFrameTransform) SetEncryptionKey1Func() (fn js.Func[func(key CryptoKey) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SFrameTransformSetEncryptionKey1Func(
			this.Ref(),
		),
	)
}

type WorkerOptions struct {
	// Type is "WorkerOptions.type"
	//
	// Optional, defaults to "classic".
	Type WorkerType
	// Credentials is "WorkerOptions.credentials"
	//
	// Optional, defaults to "same-origin".
	Credentials RequestCredentials
	// Name is "WorkerOptions.name"
	//
	// Optional, defaults to "".
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WorkerOptions with all fields set.
func (p WorkerOptions) FromRef(ref js.Ref) WorkerOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 WorkerOptions WorkerOptions [// WorkerOptions] [0x140009d5860 0x140009d5900 0x140009d59a0] 0x1400099d098 {0 0}} in the application heap.
func (p WorkerOptions) New() js.Ref {
	return bindings.WorkerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WorkerOptions) UpdateFrom(ref js.Ref) {
	bindings.WorkerOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WorkerOptions) Update(ref js.Ref) {
	bindings.WorkerOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewWorker(scriptURL js.String, options WorkerOptions) Worker {
	return Worker{}.FromRef(
		bindings.NewWorkerByWorker(
			scriptURL.Ref(),
			js.Pointer(&options)),
	)
}

func NewWorkerByWorker1(scriptURL js.String) Worker {
	return Worker{}.FromRef(
		bindings.NewWorkerByWorker1(
			scriptURL.Ref()),
	)
}

type Worker struct {
	EventTarget
}

func (this Worker) Once() Worker {
	this.Ref().Once()
	return this
}

func (this Worker) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this Worker) FromRef(ref js.Ref) Worker {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this Worker) Free() {
	this.Ref().Free()
}

// Terminate calls the method "Worker.terminate".
//
// The returned bool will be false if there is no such method.
func (this Worker) Terminate() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWorkerTerminate(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TerminateFunc returns the method "Worker.terminate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Worker) TerminateFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WorkerTerminateFunc(
			this.Ref(),
		),
	)
}

// PostMessage calls the method "Worker.postMessage".
//
// The returned bool will be false if there is no such method.
func (this Worker) PostMessage(message js.Any, transfer js.Array[js.Object]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWorkerPostMessage(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		transfer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessageFunc returns the method "Worker.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Worker) PostMessageFunc() (fn js.Func[func(message js.Any, transfer js.Array[js.Object])]) {
	return fn.FromRef(
		bindings.WorkerPostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "Worker.postMessage".
//
// The returned bool will be false if there is no such method.
func (this Worker) PostMessage1(message js.Any, options StructuredSerializeOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWorkerPostMessage1(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessage1Func returns the method "Worker.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Worker) PostMessage1Func() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	return fn.FromRef(
		bindings.WorkerPostMessage1Func(
			this.Ref(),
		),
	)
}

// PostMessage2 calls the method "Worker.postMessage".
//
// The returned bool will be false if there is no such method.
func (this Worker) PostMessage2(message js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWorkerPostMessage2(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessage2Func returns the method "Worker.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Worker) PostMessage2Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.WorkerPostMessage2Func(
			this.Ref(),
		),
	)
}

func NewRTCRtpScriptTransform(worker Worker, options js.Any, transfer js.Array[js.Object]) RTCRtpScriptTransform {
	return RTCRtpScriptTransform{}.FromRef(
		bindings.NewRTCRtpScriptTransformByRTCRtpScriptTransform(
			worker.Ref(),
			options.Ref(),
			transfer.Ref()),
	)
}

func NewRTCRtpScriptTransformByRTCRtpScriptTransform1(worker Worker, options js.Any) RTCRtpScriptTransform {
	return RTCRtpScriptTransform{}.FromRef(
		bindings.NewRTCRtpScriptTransformByRTCRtpScriptTransform1(
			worker.Ref(),
			options.Ref()),
	)
}

func NewRTCRtpScriptTransformByRTCRtpScriptTransform2(worker Worker) RTCRtpScriptTransform {
	return RTCRtpScriptTransform{}.FromRef(
		bindings.NewRTCRtpScriptTransformByRTCRtpScriptTransform2(
			worker.Ref()),
	)
}

type RTCRtpScriptTransform struct {
	ref js.Ref
}

func (this RTCRtpScriptTransform) Once() RTCRtpScriptTransform {
	this.Ref().Once()
	return this
}

func (this RTCRtpScriptTransform) Ref() js.Ref {
	return this.ref
}

func (this RTCRtpScriptTransform) FromRef(ref js.Ref) RTCRtpScriptTransform {
	this.ref = ref
	return this
}

func (this RTCRtpScriptTransform) Free() {
	this.Ref().Free()
}

type OneOf_SFrameTransform_RTCRtpScriptTransform struct {
	ref js.Ref
}

func (x OneOf_SFrameTransform_RTCRtpScriptTransform) Ref() js.Ref {
	return x.ref
}

func (x OneOf_SFrameTransform_RTCRtpScriptTransform) Free() {
	x.ref.Free()
}

func (x OneOf_SFrameTransform_RTCRtpScriptTransform) FromRef(ref js.Ref) OneOf_SFrameTransform_RTCRtpScriptTransform {
	return OneOf_SFrameTransform_RTCRtpScriptTransform{
		ref: ref,
	}
}

func (x OneOf_SFrameTransform_RTCRtpScriptTransform) SFrameTransform() SFrameTransform {
	return SFrameTransform{}.FromRef(x.ref)
}

func (x OneOf_SFrameTransform_RTCRtpScriptTransform) RTCRtpScriptTransform() RTCRtpScriptTransform {
	return RTCRtpScriptTransform{}.FromRef(x.ref)
}

type RTCRtpTransform = OneOf_SFrameTransform_RTCRtpScriptTransform

type RTCRtpSender struct {
	ref js.Ref
}

func (this RTCRtpSender) Once() RTCRtpSender {
	this.Ref().Once()
	return this
}

func (this RTCRtpSender) Ref() js.Ref {
	return this.ref
}

func (this RTCRtpSender) FromRef(ref js.Ref) RTCRtpSender {
	this.ref = ref
	return this
}

func (this RTCRtpSender) Free() {
	this.Ref().Free()
}

// Track returns the value of property "RTCRtpSender.track".
//
// The returned bool will be false if there is no such property.
func (this RTCRtpSender) Track() (MediaStreamTrack, bool) {
	var _ok bool
	_ret := bindings.GetRTCRtpSenderTrack(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaStreamTrack{}.FromRef(_ret), _ok
}

// Transport returns the value of property "RTCRtpSender.transport".
//
// The returned bool will be false if there is no such property.
func (this RTCRtpSender) Transport() (RTCDtlsTransport, bool) {
	var _ok bool
	_ret := bindings.GetRTCRtpSenderTransport(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCDtlsTransport{}.FromRef(_ret), _ok
}

// Dtmf returns the value of property "RTCRtpSender.dtmf".
//
// The returned bool will be false if there is no such property.
func (this RTCRtpSender) Dtmf() (RTCDTMFSender, bool) {
	var _ok bool
	_ret := bindings.GetRTCRtpSenderDtmf(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCDTMFSender{}.FromRef(_ret), _ok
}

// Transform returns the value of property "RTCRtpSender.transform".
//
// The returned bool will be false if there is no such property.
func (this RTCRtpSender) Transform() (RTCRtpTransform, bool) {
	var _ok bool
	_ret := bindings.GetRTCRtpSenderTransform(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCRtpTransform{}.FromRef(_ret), _ok
}

// Transform sets the value of property "RTCRtpSender.transform" to val.
//
// It returns false if the property cannot be set.
func (this RTCRtpSender) SetTransform(val RTCRtpTransform) bool {
	return js.True == bindings.SetRTCRtpSenderTransform(
		this.Ref(),
		val.Ref(),
	)
}

// GetCapabilities calls the staticmethod "RTCRtpSender.getCapabilities".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpSender) GetCapabilities(kind js.String) (RTCRtpCapabilities, bool) {
	var _ret RTCRtpCapabilities
	_ok := js.True == bindings.CallRTCRtpSenderGetCapabilities(
		this.Ref(), js.Pointer(&_ret),
		kind.Ref(),
	)

	return _ret, _ok
}

// GetCapabilitiesFunc returns the staticmethod "RTCRtpSender.getCapabilities".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpSender) GetCapabilitiesFunc() (fn js.Func[func(kind js.String) RTCRtpCapabilities]) {
	return fn.FromRef(
		bindings.RTCRtpSenderGetCapabilitiesFunc(
			this.Ref(),
		),
	)
}

// SetParameters calls the method "RTCRtpSender.setParameters".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpSender) SetParameters(parameters RTCRtpSendParameters, setParameterOptions RTCSetParameterOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCRtpSenderSetParameters(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&parameters),
		js.Pointer(&setParameterOptions),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetParametersFunc returns the method "RTCRtpSender.setParameters".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpSender) SetParametersFunc() (fn js.Func[func(parameters RTCRtpSendParameters, setParameterOptions RTCSetParameterOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCRtpSenderSetParametersFunc(
			this.Ref(),
		),
	)
}

// SetParameters1 calls the method "RTCRtpSender.setParameters".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpSender) SetParameters1(parameters RTCRtpSendParameters) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCRtpSenderSetParameters1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&parameters),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetParameters1Func returns the method "RTCRtpSender.setParameters".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpSender) SetParameters1Func() (fn js.Func[func(parameters RTCRtpSendParameters) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCRtpSenderSetParameters1Func(
			this.Ref(),
		),
	)
}

// GetParameters calls the method "RTCRtpSender.getParameters".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpSender) GetParameters() (RTCRtpSendParameters, bool) {
	var _ret RTCRtpSendParameters
	_ok := js.True == bindings.CallRTCRtpSenderGetParameters(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetParametersFunc returns the method "RTCRtpSender.getParameters".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpSender) GetParametersFunc() (fn js.Func[func() RTCRtpSendParameters]) {
	return fn.FromRef(
		bindings.RTCRtpSenderGetParametersFunc(
			this.Ref(),
		),
	)
}

// ReplaceTrack calls the method "RTCRtpSender.replaceTrack".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpSender) ReplaceTrack(withTrack MediaStreamTrack) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCRtpSenderReplaceTrack(
		this.Ref(), js.Pointer(&_ok),
		withTrack.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ReplaceTrackFunc returns the method "RTCRtpSender.replaceTrack".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpSender) ReplaceTrackFunc() (fn js.Func[func(withTrack MediaStreamTrack) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCRtpSenderReplaceTrackFunc(
			this.Ref(),
		),
	)
}

// SetStreams calls the method "RTCRtpSender.setStreams".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpSender) SetStreams(streams ...MediaStream) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCRtpSenderSetStreams(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(streams),
		js.SizeU(len(streams)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetStreamsFunc returns the method "RTCRtpSender.setStreams".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpSender) SetStreamsFunc() (fn js.Func[func(streams ...MediaStream)]) {
	return fn.FromRef(
		bindings.RTCRtpSenderSetStreamsFunc(
			this.Ref(),
		),
	)
}

// GetStats calls the method "RTCRtpSender.getStats".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpSender) GetStats() (js.Promise[RTCStatsReport], bool) {
	var _ok bool
	_ret := bindings.CallRTCRtpSenderGetStats(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[RTCStatsReport]{}.FromRef(_ret), _ok
}

// GetStatsFunc returns the method "RTCRtpSender.getStats".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpSender) GetStatsFunc() (fn js.Func[func() js.Promise[RTCStatsReport]]) {
	return fn.FromRef(
		bindings.RTCRtpSenderGetStatsFunc(
			this.Ref(),
		),
	)
}

// GenerateKeyFrame calls the method "RTCRtpSender.generateKeyFrame".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpSender) GenerateKeyFrame(rids js.Array[js.String]) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCRtpSenderGenerateKeyFrame(
		this.Ref(), js.Pointer(&_ok),
		rids.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// GenerateKeyFrameFunc returns the method "RTCRtpSender.generateKeyFrame".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpSender) GenerateKeyFrameFunc() (fn js.Func[func(rids js.Array[js.String]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCRtpSenderGenerateKeyFrameFunc(
			this.Ref(),
		),
	)
}

// GenerateKeyFrame1 calls the method "RTCRtpSender.generateKeyFrame".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpSender) GenerateKeyFrame1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCRtpSenderGenerateKeyFrame1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// GenerateKeyFrame1Func returns the method "RTCRtpSender.generateKeyFrame".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpSender) GenerateKeyFrame1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCRtpSenderGenerateKeyFrame1Func(
			this.Ref(),
		),
	)
}

type RTCRtpReceiveParameters struct {
	// HeaderExtensions is "RTCRtpReceiveParameters.headerExtensions"
	//
	// Required
	HeaderExtensions js.Array[RTCRtpHeaderExtensionParameters]
	// Rtcp is "RTCRtpReceiveParameters.rtcp"
	//
	// Required
	Rtcp RTCRtcpParameters
	// Codecs is "RTCRtpReceiveParameters.codecs"
	//
	// Required
	Codecs js.Array[RTCRtpCodecParameters]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpReceiveParameters with all fields set.
func (p RTCRtpReceiveParameters) FromRef(ref js.Ref) RTCRtpReceiveParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCRtpReceiveParameters RTCRtpReceiveParameters [// RTCRtpReceiveParameters] [0x140009d5a40 0x140009d5ae0 0x140009d5b80] 0x1400099d140 {0 0}} in the application heap.
func (p RTCRtpReceiveParameters) New() js.Ref {
	return bindings.RTCRtpReceiveParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtpReceiveParameters) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpReceiveParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtpReceiveParameters) Update(ref js.Ref) {
	bindings.RTCRtpReceiveParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCRtpContributingSource struct {
	// Timestamp is "RTCRtpContributingSource.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Source is "RTCRtpContributingSource.source"
	//
	// Required
	Source uint32
	// AudioLevel is "RTCRtpContributingSource.audioLevel"
	//
	// Optional
	AudioLevel float64
	// RtpTimestamp is "RTCRtpContributingSource.rtpTimestamp"
	//
	// Required
	RtpTimestamp uint32

	FFI_USE_AudioLevel bool // for AudioLevel.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpContributingSource with all fields set.
func (p RTCRtpContributingSource) FromRef(ref js.Ref) RTCRtpContributingSource {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCRtpContributingSource RTCRtpContributingSource [// RTCRtpContributingSource] [0x140009d5c20 0x140009d5cc0 0x140009d5d60 0x140009d5ea0 0x140009d5e00] 0x1400099d170 {0 0}} in the application heap.
func (p RTCRtpContributingSource) New() js.Ref {
	return bindings.RTCRtpContributingSourceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtpContributingSource) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpContributingSourceJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtpContributingSource) Update(ref js.Ref) {
	bindings.RTCRtpContributingSourceJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCRtpSynchronizationSource struct {
	// Timestamp is "RTCRtpSynchronizationSource.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Source is "RTCRtpSynchronizationSource.source"
	//
	// Required
	Source uint32
	// AudioLevel is "RTCRtpSynchronizationSource.audioLevel"
	//
	// Optional
	AudioLevel float64
	// RtpTimestamp is "RTCRtpSynchronizationSource.rtpTimestamp"
	//
	// Required
	RtpTimestamp uint32

	FFI_USE_AudioLevel bool // for AudioLevel.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpSynchronizationSource with all fields set.
func (p RTCRtpSynchronizationSource) FromRef(ref js.Ref) RTCRtpSynchronizationSource {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCRtpSynchronizationSource RTCRtpSynchronizationSource [// RTCRtpSynchronizationSource] [0x140009d5f40 0x140009da000 0x140009da0a0 0x140009da1e0 0x140009da140] 0x1400099d1a0 {0 0}} in the application heap.
func (p RTCRtpSynchronizationSource) New() js.Ref {
	return bindings.RTCRtpSynchronizationSourceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtpSynchronizationSource) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpSynchronizationSourceJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtpSynchronizationSource) Update(ref js.Ref) {
	bindings.RTCRtpSynchronizationSourceJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCRtpReceiver struct {
	ref js.Ref
}

func (this RTCRtpReceiver) Once() RTCRtpReceiver {
	this.Ref().Once()
	return this
}

func (this RTCRtpReceiver) Ref() js.Ref {
	return this.ref
}

func (this RTCRtpReceiver) FromRef(ref js.Ref) RTCRtpReceiver {
	this.ref = ref
	return this
}

func (this RTCRtpReceiver) Free() {
	this.Ref().Free()
}

// Track returns the value of property "RTCRtpReceiver.track".
//
// The returned bool will be false if there is no such property.
func (this RTCRtpReceiver) Track() (MediaStreamTrack, bool) {
	var _ok bool
	_ret := bindings.GetRTCRtpReceiverTrack(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaStreamTrack{}.FromRef(_ret), _ok
}

// Transport returns the value of property "RTCRtpReceiver.transport".
//
// The returned bool will be false if there is no such property.
func (this RTCRtpReceiver) Transport() (RTCDtlsTransport, bool) {
	var _ok bool
	_ret := bindings.GetRTCRtpReceiverTransport(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCDtlsTransport{}.FromRef(_ret), _ok
}

// Transform returns the value of property "RTCRtpReceiver.transform".
//
// The returned bool will be false if there is no such property.
func (this RTCRtpReceiver) Transform() (RTCRtpTransform, bool) {
	var _ok bool
	_ret := bindings.GetRTCRtpReceiverTransform(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCRtpTransform{}.FromRef(_ret), _ok
}

// Transform sets the value of property "RTCRtpReceiver.transform" to val.
//
// It returns false if the property cannot be set.
func (this RTCRtpReceiver) SetTransform(val RTCRtpTransform) bool {
	return js.True == bindings.SetRTCRtpReceiverTransform(
		this.Ref(),
		val.Ref(),
	)
}

// GetCapabilities calls the staticmethod "RTCRtpReceiver.getCapabilities".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpReceiver) GetCapabilities(kind js.String) (RTCRtpCapabilities, bool) {
	var _ret RTCRtpCapabilities
	_ok := js.True == bindings.CallRTCRtpReceiverGetCapabilities(
		this.Ref(), js.Pointer(&_ret),
		kind.Ref(),
	)

	return _ret, _ok
}

// GetCapabilitiesFunc returns the staticmethod "RTCRtpReceiver.getCapabilities".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpReceiver) GetCapabilitiesFunc() (fn js.Func[func(kind js.String) RTCRtpCapabilities]) {
	return fn.FromRef(
		bindings.RTCRtpReceiverGetCapabilitiesFunc(
			this.Ref(),
		),
	)
}

// GetParameters calls the method "RTCRtpReceiver.getParameters".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpReceiver) GetParameters() (RTCRtpReceiveParameters, bool) {
	var _ret RTCRtpReceiveParameters
	_ok := js.True == bindings.CallRTCRtpReceiverGetParameters(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetParametersFunc returns the method "RTCRtpReceiver.getParameters".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpReceiver) GetParametersFunc() (fn js.Func[func() RTCRtpReceiveParameters]) {
	return fn.FromRef(
		bindings.RTCRtpReceiverGetParametersFunc(
			this.Ref(),
		),
	)
}

// GetContributingSources calls the method "RTCRtpReceiver.getContributingSources".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpReceiver) GetContributingSources() (js.Array[RTCRtpContributingSource], bool) {
	var _ok bool
	_ret := bindings.CallRTCRtpReceiverGetContributingSources(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[RTCRtpContributingSource]{}.FromRef(_ret), _ok
}

// GetContributingSourcesFunc returns the method "RTCRtpReceiver.getContributingSources".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpReceiver) GetContributingSourcesFunc() (fn js.Func[func() js.Array[RTCRtpContributingSource]]) {
	return fn.FromRef(
		bindings.RTCRtpReceiverGetContributingSourcesFunc(
			this.Ref(),
		),
	)
}

// GetSynchronizationSources calls the method "RTCRtpReceiver.getSynchronizationSources".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpReceiver) GetSynchronizationSources() (js.Array[RTCRtpSynchronizationSource], bool) {
	var _ok bool
	_ret := bindings.CallRTCRtpReceiverGetSynchronizationSources(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[RTCRtpSynchronizationSource]{}.FromRef(_ret), _ok
}

// GetSynchronizationSourcesFunc returns the method "RTCRtpReceiver.getSynchronizationSources".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpReceiver) GetSynchronizationSourcesFunc() (fn js.Func[func() js.Array[RTCRtpSynchronizationSource]]) {
	return fn.FromRef(
		bindings.RTCRtpReceiverGetSynchronizationSourcesFunc(
			this.Ref(),
		),
	)
}

// GetStats calls the method "RTCRtpReceiver.getStats".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpReceiver) GetStats() (js.Promise[RTCStatsReport], bool) {
	var _ok bool
	_ret := bindings.CallRTCRtpReceiverGetStats(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[RTCStatsReport]{}.FromRef(_ret), _ok
}

// GetStatsFunc returns the method "RTCRtpReceiver.getStats".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpReceiver) GetStatsFunc() (fn js.Func[func() js.Promise[RTCStatsReport]]) {
	return fn.FromRef(
		bindings.RTCRtpReceiverGetStatsFunc(
			this.Ref(),
		),
	)
}

type RTCRtpTransceiverDirection uint32

const (
	_ RTCRtpTransceiverDirection = iota

	RTCRtpTransceiverDirection_SENDRECV
	RTCRtpTransceiverDirection_SENDONLY
	RTCRtpTransceiverDirection_RECVONLY
	RTCRtpTransceiverDirection_INACTIVE
	RTCRtpTransceiverDirection_STOPPED
)

func (RTCRtpTransceiverDirection) FromRef(str js.Ref) RTCRtpTransceiverDirection {
	return RTCRtpTransceiverDirection(bindings.ConstOfRTCRtpTransceiverDirection(str))
}

func (x RTCRtpTransceiverDirection) String() (string, bool) {
	switch x {
	case RTCRtpTransceiverDirection_SENDRECV:
		return "sendrecv", true
	case RTCRtpTransceiverDirection_SENDONLY:
		return "sendonly", true
	case RTCRtpTransceiverDirection_RECVONLY:
		return "recvonly", true
	case RTCRtpTransceiverDirection_INACTIVE:
		return "inactive", true
	case RTCRtpTransceiverDirection_STOPPED:
		return "stopped", true
	default:
		return "", false
	}
}

type RTCRtpTransceiver struct {
	ref js.Ref
}

func (this RTCRtpTransceiver) Once() RTCRtpTransceiver {
	this.Ref().Once()
	return this
}

func (this RTCRtpTransceiver) Ref() js.Ref {
	return this.ref
}

func (this RTCRtpTransceiver) FromRef(ref js.Ref) RTCRtpTransceiver {
	this.ref = ref
	return this
}

func (this RTCRtpTransceiver) Free() {
	this.Ref().Free()
}

// Mid returns the value of property "RTCRtpTransceiver.mid".
//
// The returned bool will be false if there is no such property.
func (this RTCRtpTransceiver) Mid() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCRtpTransceiverMid(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Sender returns the value of property "RTCRtpTransceiver.sender".
//
// The returned bool will be false if there is no such property.
func (this RTCRtpTransceiver) Sender() (RTCRtpSender, bool) {
	var _ok bool
	_ret := bindings.GetRTCRtpTransceiverSender(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCRtpSender{}.FromRef(_ret), _ok
}

// Receiver returns the value of property "RTCRtpTransceiver.receiver".
//
// The returned bool will be false if there is no such property.
func (this RTCRtpTransceiver) Receiver() (RTCRtpReceiver, bool) {
	var _ok bool
	_ret := bindings.GetRTCRtpTransceiverReceiver(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCRtpReceiver{}.FromRef(_ret), _ok
}

// Direction returns the value of property "RTCRtpTransceiver.direction".
//
// The returned bool will be false if there is no such property.
func (this RTCRtpTransceiver) Direction() (RTCRtpTransceiverDirection, bool) {
	var _ok bool
	_ret := bindings.GetRTCRtpTransceiverDirection(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCRtpTransceiverDirection(_ret), _ok
}

// Direction sets the value of property "RTCRtpTransceiver.direction" to val.
//
// It returns false if the property cannot be set.
func (this RTCRtpTransceiver) SetDirection(val RTCRtpTransceiverDirection) bool {
	return js.True == bindings.SetRTCRtpTransceiverDirection(
		this.Ref(),
		uint32(val),
	)
}

// CurrentDirection returns the value of property "RTCRtpTransceiver.currentDirection".
//
// The returned bool will be false if there is no such property.
func (this RTCRtpTransceiver) CurrentDirection() (RTCRtpTransceiverDirection, bool) {
	var _ok bool
	_ret := bindings.GetRTCRtpTransceiverCurrentDirection(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCRtpTransceiverDirection(_ret), _ok
}

// Stop calls the method "RTCRtpTransceiver.stop".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpTransceiver) Stop() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCRtpTransceiverStop(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StopFunc returns the method "RTCRtpTransceiver.stop".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpTransceiver) StopFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RTCRtpTransceiverStopFunc(
			this.Ref(),
		),
	)
}

// SetCodecPreferences calls the method "RTCRtpTransceiver.setCodecPreferences".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpTransceiver) SetCodecPreferences(codecs js.Array[RTCRtpCodecCapability]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCRtpTransceiverSetCodecPreferences(
		this.Ref(), js.Pointer(&_ok),
		codecs.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetCodecPreferencesFunc returns the method "RTCRtpTransceiver.setCodecPreferences".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpTransceiver) SetCodecPreferencesFunc() (fn js.Func[func(codecs js.Array[RTCRtpCodecCapability])]) {
	return fn.FromRef(
		bindings.RTCRtpTransceiverSetCodecPreferencesFunc(
			this.Ref(),
		),
	)
}

type OneOf_MediaStreamTrack_String struct {
	ref js.Ref
}

func (x OneOf_MediaStreamTrack_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_MediaStreamTrack_String) Free() {
	x.ref.Free()
}

func (x OneOf_MediaStreamTrack_String) FromRef(ref js.Ref) OneOf_MediaStreamTrack_String {
	return OneOf_MediaStreamTrack_String{
		ref: ref,
	}
}

func (x OneOf_MediaStreamTrack_String) MediaStreamTrack() MediaStreamTrack {
	return MediaStreamTrack{}.FromRef(x.ref)
}

func (x OneOf_MediaStreamTrack_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type RTCRtpTransceiverInit struct {
	// Direction is "RTCRtpTransceiverInit.direction"
	//
	// Optional, defaults to "sendrecv".
	Direction RTCRtpTransceiverDirection
	// Streams is "RTCRtpTransceiverInit.streams"
	//
	// Optional, defaults to [].
	Streams js.Array[MediaStream]
	// SendEncodings is "RTCRtpTransceiverInit.sendEncodings"
	//
	// Optional, defaults to [].
	SendEncodings js.Array[RTCRtpEncodingParameters]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpTransceiverInit with all fields set.
func (p RTCRtpTransceiverInit) FromRef(ref js.Ref) RTCRtpTransceiverInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCRtpTransceiverInit RTCRtpTransceiverInit [// RTCRtpTransceiverInit] [0x140009da3c0 0x140009da460 0x140009da500] 0x1400099d230 {0 0}} in the application heap.
func (p RTCRtpTransceiverInit) New() js.Ref {
	return bindings.RTCRtpTransceiverInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtpTransceiverInit) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpTransceiverInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtpTransceiverInit) Update(ref js.Ref) {
	bindings.RTCRtpTransceiverInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewRTCSessionDescription(descriptionInitDict RTCSessionDescriptionInit) RTCSessionDescription {
	return RTCSessionDescription{}.FromRef(
		bindings.NewRTCSessionDescriptionByRTCSessionDescription(
			js.Pointer(&descriptionInitDict)),
	)
}

type RTCSessionDescription struct {
	ref js.Ref
}

func (this RTCSessionDescription) Once() RTCSessionDescription {
	this.Ref().Once()
	return this
}

func (this RTCSessionDescription) Ref() js.Ref {
	return this.ref
}

func (this RTCSessionDescription) FromRef(ref js.Ref) RTCSessionDescription {
	this.ref = ref
	return this
}

func (this RTCSessionDescription) Free() {
	this.Ref().Free()
}

// Type returns the value of property "RTCSessionDescription.type".
//
// The returned bool will be false if there is no such property.
func (this RTCSessionDescription) Type() (RTCSdpType, bool) {
	var _ok bool
	_ret := bindings.GetRTCSessionDescriptionType(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCSdpType(_ret), _ok
}

// Sdp returns the value of property "RTCSessionDescription.sdp".
//
// The returned bool will be false if there is no such property.
func (this RTCSessionDescription) Sdp() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCSessionDescriptionSdp(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ToJSON calls the method "RTCSessionDescription.toJSON".
//
// The returned bool will be false if there is no such method.
func (this RTCSessionDescription) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallRTCSessionDescriptionToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "RTCSessionDescription.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCSessionDescription) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.RTCSessionDescriptionToJSONFunc(
			this.Ref(),
		),
	)
}

type RTCSignalingState uint32

const (
	_ RTCSignalingState = iota

	RTCSignalingState_STABLE
	RTCSignalingState_HAVE_LOCAL_OFFER
	RTCSignalingState_HAVE_REMOTE_OFFER
	RTCSignalingState_HAVE_LOCAL_PRANSWER
	RTCSignalingState_HAVE_REMOTE_PRANSWER
	RTCSignalingState_CLOSED
)

func (RTCSignalingState) FromRef(str js.Ref) RTCSignalingState {
	return RTCSignalingState(bindings.ConstOfRTCSignalingState(str))
}

func (x RTCSignalingState) String() (string, bool) {
	switch x {
	case RTCSignalingState_STABLE:
		return "stable", true
	case RTCSignalingState_HAVE_LOCAL_OFFER:
		return "have-local-offer", true
	case RTCSignalingState_HAVE_REMOTE_OFFER:
		return "have-remote-offer", true
	case RTCSignalingState_HAVE_LOCAL_PRANSWER:
		return "have-local-pranswer", true
	case RTCSignalingState_HAVE_REMOTE_PRANSWER:
		return "have-remote-pranswer", true
	case RTCSignalingState_CLOSED:
		return "closed", true
	default:
		return "", false
	}
}

type RTCPeerConnectionState uint32

const (
	_ RTCPeerConnectionState = iota

	RTCPeerConnectionState_CLOSED
	RTCPeerConnectionState_FAILED
	RTCPeerConnectionState_DISCONNECTED
	RTCPeerConnectionState_NEW
	RTCPeerConnectionState_CONNECTING
	RTCPeerConnectionState_CONNECTED
)

func (RTCPeerConnectionState) FromRef(str js.Ref) RTCPeerConnectionState {
	return RTCPeerConnectionState(bindings.ConstOfRTCPeerConnectionState(str))
}

func (x RTCPeerConnectionState) String() (string, bool) {
	switch x {
	case RTCPeerConnectionState_CLOSED:
		return "closed", true
	case RTCPeerConnectionState_FAILED:
		return "failed", true
	case RTCPeerConnectionState_DISCONNECTED:
		return "disconnected", true
	case RTCPeerConnectionState_NEW:
		return "new", true
	case RTCPeerConnectionState_CONNECTING:
		return "connecting", true
	case RTCPeerConnectionState_CONNECTED:
		return "connected", true
	default:
		return "", false
	}
}

type RTCSctpTransportState uint32

const (
	_ RTCSctpTransportState = iota

	RTCSctpTransportState_CONNECTING
	RTCSctpTransportState_CONNECTED
	RTCSctpTransportState_CLOSED
)

func (RTCSctpTransportState) FromRef(str js.Ref) RTCSctpTransportState {
	return RTCSctpTransportState(bindings.ConstOfRTCSctpTransportState(str))
}

func (x RTCSctpTransportState) String() (string, bool) {
	switch x {
	case RTCSctpTransportState_CONNECTING:
		return "connecting", true
	case RTCSctpTransportState_CONNECTED:
		return "connected", true
	case RTCSctpTransportState_CLOSED:
		return "closed", true
	default:
		return "", false
	}
}

type RTCSctpTransport struct {
	EventTarget
}

func (this RTCSctpTransport) Once() RTCSctpTransport {
	this.Ref().Once()
	return this
}

func (this RTCSctpTransport) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this RTCSctpTransport) FromRef(ref js.Ref) RTCSctpTransport {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this RTCSctpTransport) Free() {
	this.Ref().Free()
}

// Transport returns the value of property "RTCSctpTransport.transport".
//
// The returned bool will be false if there is no such property.
func (this RTCSctpTransport) Transport() (RTCDtlsTransport, bool) {
	var _ok bool
	_ret := bindings.GetRTCSctpTransportTransport(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCDtlsTransport{}.FromRef(_ret), _ok
}

// State returns the value of property "RTCSctpTransport.state".
//
// The returned bool will be false if there is no such property.
func (this RTCSctpTransport) State() (RTCSctpTransportState, bool) {
	var _ok bool
	_ret := bindings.GetRTCSctpTransportState(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCSctpTransportState(_ret), _ok
}

// MaxMessageSize returns the value of property "RTCSctpTransport.maxMessageSize".
//
// The returned bool will be false if there is no such property.
func (this RTCSctpTransport) MaxMessageSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetRTCSctpTransportMaxMessageSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// MaxChannels returns the value of property "RTCSctpTransport.maxChannels".
//
// The returned bool will be false if there is no such property.
func (this RTCSctpTransport) MaxChannels() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetRTCSctpTransportMaxChannels(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

func NewRTCPeerConnection(configuration RTCConfiguration) RTCPeerConnection {
	return RTCPeerConnection{}.FromRef(
		bindings.NewRTCPeerConnectionByRTCPeerConnection(
			js.Pointer(&configuration)),
	)
}

func NewRTCPeerConnectionByRTCPeerConnection1() RTCPeerConnection {
	return RTCPeerConnection{}.FromRef(
		bindings.NewRTCPeerConnectionByRTCPeerConnection1(),
	)
}

type RTCPeerConnection struct {
	EventTarget
}

func (this RTCPeerConnection) Once() RTCPeerConnection {
	this.Ref().Once()
	return this
}

func (this RTCPeerConnection) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this RTCPeerConnection) FromRef(ref js.Ref) RTCPeerConnection {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this RTCPeerConnection) Free() {
	this.Ref().Free()
}

// LocalDescription returns the value of property "RTCPeerConnection.localDescription".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnection) LocalDescription() (RTCSessionDescription, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionLocalDescription(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCSessionDescription{}.FromRef(_ret), _ok
}

// CurrentLocalDescription returns the value of property "RTCPeerConnection.currentLocalDescription".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnection) CurrentLocalDescription() (RTCSessionDescription, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionCurrentLocalDescription(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCSessionDescription{}.FromRef(_ret), _ok
}

// PendingLocalDescription returns the value of property "RTCPeerConnection.pendingLocalDescription".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnection) PendingLocalDescription() (RTCSessionDescription, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionPendingLocalDescription(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCSessionDescription{}.FromRef(_ret), _ok
}

// RemoteDescription returns the value of property "RTCPeerConnection.remoteDescription".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnection) RemoteDescription() (RTCSessionDescription, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionRemoteDescription(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCSessionDescription{}.FromRef(_ret), _ok
}

// CurrentRemoteDescription returns the value of property "RTCPeerConnection.currentRemoteDescription".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnection) CurrentRemoteDescription() (RTCSessionDescription, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionCurrentRemoteDescription(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCSessionDescription{}.FromRef(_ret), _ok
}

// PendingRemoteDescription returns the value of property "RTCPeerConnection.pendingRemoteDescription".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnection) PendingRemoteDescription() (RTCSessionDescription, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionPendingRemoteDescription(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCSessionDescription{}.FromRef(_ret), _ok
}

// SignalingState returns the value of property "RTCPeerConnection.signalingState".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnection) SignalingState() (RTCSignalingState, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionSignalingState(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCSignalingState(_ret), _ok
}

// IceGatheringState returns the value of property "RTCPeerConnection.iceGatheringState".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnection) IceGatheringState() (RTCIceGatheringState, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionIceGatheringState(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCIceGatheringState(_ret), _ok
}

// IceConnectionState returns the value of property "RTCPeerConnection.iceConnectionState".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnection) IceConnectionState() (RTCIceConnectionState, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionIceConnectionState(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCIceConnectionState(_ret), _ok
}

// ConnectionState returns the value of property "RTCPeerConnection.connectionState".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnection) ConnectionState() (RTCPeerConnectionState, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionConnectionState(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCPeerConnectionState(_ret), _ok
}

// CanTrickleIceCandidates returns the value of property "RTCPeerConnection.canTrickleIceCandidates".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnection) CanTrickleIceCandidates() (bool, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionCanTrickleIceCandidates(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Sctp returns the value of property "RTCPeerConnection.sctp".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnection) Sctp() (RTCSctpTransport, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionSctp(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCSctpTransport{}.FromRef(_ret), _ok
}

// PeerIdentity returns the value of property "RTCPeerConnection.peerIdentity".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnection) PeerIdentity() (js.Promise[RTCIdentityAssertion], bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionPeerIdentity(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[RTCIdentityAssertion]{}.FromRef(_ret), _ok
}

// IdpLoginUrl returns the value of property "RTCPeerConnection.idpLoginUrl".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnection) IdpLoginUrl() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionIdpLoginUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// IdpErrorInfo returns the value of property "RTCPeerConnection.idpErrorInfo".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnection) IdpErrorInfo() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionIdpErrorInfo(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CreateOffer calls the method "RTCPeerConnection.createOffer".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) CreateOffer(options RTCOfferOptions) (js.Promise[RTCSessionDescriptionInit], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionCreateOffer(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[RTCSessionDescriptionInit]{}.FromRef(_ret), _ok
}

// CreateOfferFunc returns the method "RTCPeerConnection.createOffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) CreateOfferFunc() (fn js.Func[func(options RTCOfferOptions) js.Promise[RTCSessionDescriptionInit]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateOfferFunc(
			this.Ref(),
		),
	)
}

// CreateOffer1 calls the method "RTCPeerConnection.createOffer".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) CreateOffer1() (js.Promise[RTCSessionDescriptionInit], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionCreateOffer1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[RTCSessionDescriptionInit]{}.FromRef(_ret), _ok
}

// CreateOffer1Func returns the method "RTCPeerConnection.createOffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) CreateOffer1Func() (fn js.Func[func() js.Promise[RTCSessionDescriptionInit]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateOffer1Func(
			this.Ref(),
		),
	)
}

// CreateAnswer calls the method "RTCPeerConnection.createAnswer".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) CreateAnswer(options RTCAnswerOptions) (js.Promise[RTCSessionDescriptionInit], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionCreateAnswer(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[RTCSessionDescriptionInit]{}.FromRef(_ret), _ok
}

// CreateAnswerFunc returns the method "RTCPeerConnection.createAnswer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) CreateAnswerFunc() (fn js.Func[func(options RTCAnswerOptions) js.Promise[RTCSessionDescriptionInit]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateAnswerFunc(
			this.Ref(),
		),
	)
}

// CreateAnswer1 calls the method "RTCPeerConnection.createAnswer".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) CreateAnswer1() (js.Promise[RTCSessionDescriptionInit], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionCreateAnswer1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[RTCSessionDescriptionInit]{}.FromRef(_ret), _ok
}

// CreateAnswer1Func returns the method "RTCPeerConnection.createAnswer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) CreateAnswer1Func() (fn js.Func[func() js.Promise[RTCSessionDescriptionInit]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateAnswer1Func(
			this.Ref(),
		),
	)
}

// SetLocalDescription calls the method "RTCPeerConnection.setLocalDescription".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) SetLocalDescription(description RTCLocalSessionDescriptionInit) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionSetLocalDescription(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&description),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetLocalDescriptionFunc returns the method "RTCPeerConnection.setLocalDescription".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) SetLocalDescriptionFunc() (fn js.Func[func(description RTCLocalSessionDescriptionInit) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetLocalDescriptionFunc(
			this.Ref(),
		),
	)
}

// SetLocalDescription1 calls the method "RTCPeerConnection.setLocalDescription".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) SetLocalDescription1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionSetLocalDescription1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetLocalDescription1Func returns the method "RTCPeerConnection.setLocalDescription".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) SetLocalDescription1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetLocalDescription1Func(
			this.Ref(),
		),
	)
}

// SetRemoteDescription calls the method "RTCPeerConnection.setRemoteDescription".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) SetRemoteDescription(description RTCSessionDescriptionInit) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionSetRemoteDescription(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&description),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetRemoteDescriptionFunc returns the method "RTCPeerConnection.setRemoteDescription".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) SetRemoteDescriptionFunc() (fn js.Func[func(description RTCSessionDescriptionInit) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetRemoteDescriptionFunc(
			this.Ref(),
		),
	)
}

// AddIceCandidate calls the method "RTCPeerConnection.addIceCandidate".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) AddIceCandidate(candidate RTCIceCandidateInit) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionAddIceCandidate(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&candidate),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// AddIceCandidateFunc returns the method "RTCPeerConnection.addIceCandidate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) AddIceCandidateFunc() (fn js.Func[func(candidate RTCIceCandidateInit) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionAddIceCandidateFunc(
			this.Ref(),
		),
	)
}

// AddIceCandidate1 calls the method "RTCPeerConnection.addIceCandidate".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) AddIceCandidate1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionAddIceCandidate1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// AddIceCandidate1Func returns the method "RTCPeerConnection.addIceCandidate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) AddIceCandidate1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionAddIceCandidate1Func(
			this.Ref(),
		),
	)
}

// RestartIce calls the method "RTCPeerConnection.restartIce".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) RestartIce() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionRestartIce(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RestartIceFunc returns the method "RTCPeerConnection.restartIce".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) RestartIceFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionRestartIceFunc(
			this.Ref(),
		),
	)
}

// GetConfiguration calls the method "RTCPeerConnection.getConfiguration".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) GetConfiguration() (RTCConfiguration, bool) {
	var _ret RTCConfiguration
	_ok := js.True == bindings.CallRTCPeerConnectionGetConfiguration(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetConfigurationFunc returns the method "RTCPeerConnection.getConfiguration".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) GetConfigurationFunc() (fn js.Func[func() RTCConfiguration]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGetConfigurationFunc(
			this.Ref(),
		),
	)
}

// SetConfiguration calls the method "RTCPeerConnection.setConfiguration".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) SetConfiguration(configuration RTCConfiguration) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionSetConfiguration(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&configuration),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetConfigurationFunc returns the method "RTCPeerConnection.setConfiguration".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) SetConfigurationFunc() (fn js.Func[func(configuration RTCConfiguration)]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetConfigurationFunc(
			this.Ref(),
		),
	)
}

// SetConfiguration1 calls the method "RTCPeerConnection.setConfiguration".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) SetConfiguration1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionSetConfiguration1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetConfiguration1Func returns the method "RTCPeerConnection.setConfiguration".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) SetConfiguration1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetConfiguration1Func(
			this.Ref(),
		),
	)
}

// Close calls the method "RTCPeerConnection.close".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "RTCPeerConnection.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCloseFunc(
			this.Ref(),
		),
	)
}

// CreateOffer2 calls the method "RTCPeerConnection.createOffer".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) CreateOffer2(successCallback js.Func[func(description RTCSessionDescriptionInit)], failureCallback js.Func[func(err DOMException)], options RTCOfferOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionCreateOffer2(
		this.Ref(), js.Pointer(&_ok),
		successCallback.Ref(),
		failureCallback.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CreateOffer2Func returns the method "RTCPeerConnection.createOffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) CreateOffer2Func() (fn js.Func[func(successCallback js.Func[func(description RTCSessionDescriptionInit)], failureCallback js.Func[func(err DOMException)], options RTCOfferOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateOffer2Func(
			this.Ref(),
		),
	)
}

// CreateOffer3 calls the method "RTCPeerConnection.createOffer".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) CreateOffer3(successCallback js.Func[func(description RTCSessionDescriptionInit)], failureCallback js.Func[func(err DOMException)]) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionCreateOffer3(
		this.Ref(), js.Pointer(&_ok),
		successCallback.Ref(),
		failureCallback.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CreateOffer3Func returns the method "RTCPeerConnection.createOffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) CreateOffer3Func() (fn js.Func[func(successCallback js.Func[func(description RTCSessionDescriptionInit)], failureCallback js.Func[func(err DOMException)]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateOffer3Func(
			this.Ref(),
		),
	)
}

// SetLocalDescription2 calls the method "RTCPeerConnection.setLocalDescription".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) SetLocalDescription2(description RTCLocalSessionDescriptionInit, successCallback js.Func[func()], failureCallback js.Func[func(err DOMException)]) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionSetLocalDescription2(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&description),
		successCallback.Ref(),
		failureCallback.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetLocalDescription2Func returns the method "RTCPeerConnection.setLocalDescription".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) SetLocalDescription2Func() (fn js.Func[func(description RTCLocalSessionDescriptionInit, successCallback js.Func[func()], failureCallback js.Func[func(err DOMException)]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetLocalDescription2Func(
			this.Ref(),
		),
	)
}

// CreateAnswer2 calls the method "RTCPeerConnection.createAnswer".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) CreateAnswer2(successCallback js.Func[func(description RTCSessionDescriptionInit)], failureCallback js.Func[func(err DOMException)]) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionCreateAnswer2(
		this.Ref(), js.Pointer(&_ok),
		successCallback.Ref(),
		failureCallback.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CreateAnswer2Func returns the method "RTCPeerConnection.createAnswer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) CreateAnswer2Func() (fn js.Func[func(successCallback js.Func[func(description RTCSessionDescriptionInit)], failureCallback js.Func[func(err DOMException)]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateAnswer2Func(
			this.Ref(),
		),
	)
}

// SetRemoteDescription1 calls the method "RTCPeerConnection.setRemoteDescription".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) SetRemoteDescription1(description RTCSessionDescriptionInit, successCallback js.Func[func()], failureCallback js.Func[func(err DOMException)]) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionSetRemoteDescription1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&description),
		successCallback.Ref(),
		failureCallback.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetRemoteDescription1Func returns the method "RTCPeerConnection.setRemoteDescription".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) SetRemoteDescription1Func() (fn js.Func[func(description RTCSessionDescriptionInit, successCallback js.Func[func()], failureCallback js.Func[func(err DOMException)]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetRemoteDescription1Func(
			this.Ref(),
		),
	)
}

// AddIceCandidate2 calls the method "RTCPeerConnection.addIceCandidate".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) AddIceCandidate2(candidate RTCIceCandidateInit, successCallback js.Func[func()], failureCallback js.Func[func(err DOMException)]) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionAddIceCandidate2(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&candidate),
		successCallback.Ref(),
		failureCallback.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// AddIceCandidate2Func returns the method "RTCPeerConnection.addIceCandidate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) AddIceCandidate2Func() (fn js.Func[func(candidate RTCIceCandidateInit, successCallback js.Func[func()], failureCallback js.Func[func(err DOMException)]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionAddIceCandidate2Func(
			this.Ref(),
		),
	)
}

// CreateDataChannel calls the method "RTCPeerConnection.createDataChannel".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) CreateDataChannel(label js.String, dataChannelDict RTCDataChannelInit) (RTCDataChannel, bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionCreateDataChannel(
		this.Ref(), js.Pointer(&_ok),
		label.Ref(),
		js.Pointer(&dataChannelDict),
	)

	return RTCDataChannel{}.FromRef(_ret), _ok
}

// CreateDataChannelFunc returns the method "RTCPeerConnection.createDataChannel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) CreateDataChannelFunc() (fn js.Func[func(label js.String, dataChannelDict RTCDataChannelInit) RTCDataChannel]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateDataChannelFunc(
			this.Ref(),
		),
	)
}

// CreateDataChannel1 calls the method "RTCPeerConnection.createDataChannel".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) CreateDataChannel1(label js.String) (RTCDataChannel, bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionCreateDataChannel1(
		this.Ref(), js.Pointer(&_ok),
		label.Ref(),
	)

	return RTCDataChannel{}.FromRef(_ret), _ok
}

// CreateDataChannel1Func returns the method "RTCPeerConnection.createDataChannel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) CreateDataChannel1Func() (fn js.Func[func(label js.String) RTCDataChannel]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateDataChannel1Func(
			this.Ref(),
		),
	)
}

// GetSenders calls the method "RTCPeerConnection.getSenders".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) GetSenders() (js.Array[RTCRtpSender], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionGetSenders(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[RTCRtpSender]{}.FromRef(_ret), _ok
}

// GetSendersFunc returns the method "RTCPeerConnection.getSenders".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) GetSendersFunc() (fn js.Func[func() js.Array[RTCRtpSender]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGetSendersFunc(
			this.Ref(),
		),
	)
}

// GetReceivers calls the method "RTCPeerConnection.getReceivers".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) GetReceivers() (js.Array[RTCRtpReceiver], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionGetReceivers(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[RTCRtpReceiver]{}.FromRef(_ret), _ok
}

// GetReceiversFunc returns the method "RTCPeerConnection.getReceivers".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) GetReceiversFunc() (fn js.Func[func() js.Array[RTCRtpReceiver]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGetReceiversFunc(
			this.Ref(),
		),
	)
}

// GetTransceivers calls the method "RTCPeerConnection.getTransceivers".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) GetTransceivers() (js.Array[RTCRtpTransceiver], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionGetTransceivers(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[RTCRtpTransceiver]{}.FromRef(_ret), _ok
}

// GetTransceiversFunc returns the method "RTCPeerConnection.getTransceivers".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) GetTransceiversFunc() (fn js.Func[func() js.Array[RTCRtpTransceiver]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGetTransceiversFunc(
			this.Ref(),
		),
	)
}

// AddTrack calls the method "RTCPeerConnection.addTrack".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) AddTrack(track MediaStreamTrack, streams ...MediaStream) (RTCRtpSender, bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionAddTrack(
		this.Ref(), js.Pointer(&_ok),
		track.Ref(),
		js.SliceData(streams),
		js.SizeU(len(streams)),
	)

	return RTCRtpSender{}.FromRef(_ret), _ok
}

// AddTrackFunc returns the method "RTCPeerConnection.addTrack".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) AddTrackFunc() (fn js.Func[func(track MediaStreamTrack, streams ...MediaStream) RTCRtpSender]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionAddTrackFunc(
			this.Ref(),
		),
	)
}

// RemoveTrack calls the method "RTCPeerConnection.removeTrack".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) RemoveTrack(sender RTCRtpSender) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionRemoveTrack(
		this.Ref(), js.Pointer(&_ok),
		sender.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveTrackFunc returns the method "RTCPeerConnection.removeTrack".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) RemoveTrackFunc() (fn js.Func[func(sender RTCRtpSender)]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionRemoveTrackFunc(
			this.Ref(),
		),
	)
}

// AddTransceiver calls the method "RTCPeerConnection.addTransceiver".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) AddTransceiver(trackOrKind OneOf_MediaStreamTrack_String, init RTCRtpTransceiverInit) (RTCRtpTransceiver, bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionAddTransceiver(
		this.Ref(), js.Pointer(&_ok),
		trackOrKind.Ref(),
		js.Pointer(&init),
	)

	return RTCRtpTransceiver{}.FromRef(_ret), _ok
}

// AddTransceiverFunc returns the method "RTCPeerConnection.addTransceiver".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) AddTransceiverFunc() (fn js.Func[func(trackOrKind OneOf_MediaStreamTrack_String, init RTCRtpTransceiverInit) RTCRtpTransceiver]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionAddTransceiverFunc(
			this.Ref(),
		),
	)
}

// AddTransceiver1 calls the method "RTCPeerConnection.addTransceiver".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) AddTransceiver1(trackOrKind OneOf_MediaStreamTrack_String) (RTCRtpTransceiver, bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionAddTransceiver1(
		this.Ref(), js.Pointer(&_ok),
		trackOrKind.Ref(),
	)

	return RTCRtpTransceiver{}.FromRef(_ret), _ok
}

// AddTransceiver1Func returns the method "RTCPeerConnection.addTransceiver".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) AddTransceiver1Func() (fn js.Func[func(trackOrKind OneOf_MediaStreamTrack_String) RTCRtpTransceiver]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionAddTransceiver1Func(
			this.Ref(),
		),
	)
}

// GetStats calls the method "RTCPeerConnection.getStats".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) GetStats(selector MediaStreamTrack) (js.Promise[RTCStatsReport], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionGetStats(
		this.Ref(), js.Pointer(&_ok),
		selector.Ref(),
	)

	return js.Promise[RTCStatsReport]{}.FromRef(_ret), _ok
}

// GetStatsFunc returns the method "RTCPeerConnection.getStats".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) GetStatsFunc() (fn js.Func[func(selector MediaStreamTrack) js.Promise[RTCStatsReport]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGetStatsFunc(
			this.Ref(),
		),
	)
}

// GetStats1 calls the method "RTCPeerConnection.getStats".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) GetStats1() (js.Promise[RTCStatsReport], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionGetStats1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[RTCStatsReport]{}.FromRef(_ret), _ok
}

// GetStats1Func returns the method "RTCPeerConnection.getStats".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) GetStats1Func() (fn js.Func[func() js.Promise[RTCStatsReport]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGetStats1Func(
			this.Ref(),
		),
	)
}

// GenerateCertificate calls the staticmethod "RTCPeerConnection.generateCertificate".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) GenerateCertificate(keygenAlgorithm AlgorithmIdentifier) (js.Promise[RTCCertificate], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionGenerateCertificate(
		this.Ref(), js.Pointer(&_ok),
		keygenAlgorithm.Ref(),
	)

	return js.Promise[RTCCertificate]{}.FromRef(_ret), _ok
}

// GenerateCertificateFunc returns the staticmethod "RTCPeerConnection.generateCertificate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) GenerateCertificateFunc() (fn js.Func[func(keygenAlgorithm AlgorithmIdentifier) js.Promise[RTCCertificate]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGenerateCertificateFunc(
			this.Ref(),
		),
	)
}

// SetIdentityProvider calls the method "RTCPeerConnection.setIdentityProvider".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) SetIdentityProvider(provider js.String, options RTCIdentityProviderOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionSetIdentityProvider(
		this.Ref(), js.Pointer(&_ok),
		provider.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetIdentityProviderFunc returns the method "RTCPeerConnection.setIdentityProvider".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) SetIdentityProviderFunc() (fn js.Func[func(provider js.String, options RTCIdentityProviderOptions)]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetIdentityProviderFunc(
			this.Ref(),
		),
	)
}

// SetIdentityProvider1 calls the method "RTCPeerConnection.setIdentityProvider".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) SetIdentityProvider1(provider js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionSetIdentityProvider1(
		this.Ref(), js.Pointer(&_ok),
		provider.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetIdentityProvider1Func returns the method "RTCPeerConnection.setIdentityProvider".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) SetIdentityProvider1Func() (fn js.Func[func(provider js.String)]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetIdentityProvider1Func(
			this.Ref(),
		),
	)
}

// GetIdentityAssertion calls the method "RTCPeerConnection.getIdentityAssertion".
//
// The returned bool will be false if there is no such method.
func (this RTCPeerConnection) GetIdentityAssertion() (js.Promise[js.String], bool) {
	var _ok bool
	_ret := bindings.CallRTCPeerConnectionGetIdentityAssertion(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.String]{}.FromRef(_ret), _ok
}

// GetIdentityAssertionFunc returns the method "RTCPeerConnection.getIdentityAssertion".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCPeerConnection) GetIdentityAssertionFunc() (fn js.Func[func() js.Promise[js.String]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGetIdentityAssertionFunc(
			this.Ref(),
		),
	)
}

type RTCPeerConnectionIceErrorEventInit struct {
	// Address is "RTCPeerConnectionIceErrorEventInit.address"
	//
	// Optional
	Address js.String
	// Port is "RTCPeerConnectionIceErrorEventInit.port"
	//
	// Optional
	Port uint16
	// Url is "RTCPeerConnectionIceErrorEventInit.url"
	//
	// Optional
	Url js.String
	// ErrorCode is "RTCPeerConnectionIceErrorEventInit.errorCode"
	//
	// Required
	ErrorCode uint16
	// ErrorText is "RTCPeerConnectionIceErrorEventInit.errorText"
	//
	// Optional
	ErrorText js.String
	// Bubbles is "RTCPeerConnectionIceErrorEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "RTCPeerConnectionIceErrorEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "RTCPeerConnectionIceErrorEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Port       bool // for Port.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCPeerConnectionIceErrorEventInit with all fields set.
func (p RTCPeerConnectionIceErrorEventInit) FromRef(ref js.Ref) RTCPeerConnectionIceErrorEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCPeerConnectionIceErrorEventInit RTCPeerConnectionIceErrorEventInit [// RTCPeerConnectionIceErrorEventInit] [0x140009da780 0x140009da820 0x140009da960 0x140009daa00 0x140009daaa0 0x140009dab40 0x140009dac80 0x140009dadc0 0x140009da8c0 0x140009dabe0 0x140009dad20 0x140009dae60] 0x1400099d3c8 {0 0}} in the application heap.
func (p RTCPeerConnectionIceErrorEventInit) New() js.Ref {
	return bindings.RTCPeerConnectionIceErrorEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCPeerConnectionIceErrorEventInit) UpdateFrom(ref js.Ref) {
	bindings.RTCPeerConnectionIceErrorEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCPeerConnectionIceErrorEventInit) Update(ref js.Ref) {
	bindings.RTCPeerConnectionIceErrorEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewRTCPeerConnectionIceErrorEvent(typ js.String, eventInitDict RTCPeerConnectionIceErrorEventInit) RTCPeerConnectionIceErrorEvent {
	return RTCPeerConnectionIceErrorEvent{}.FromRef(
		bindings.NewRTCPeerConnectionIceErrorEventByRTCPeerConnectionIceErrorEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type RTCPeerConnectionIceErrorEvent struct {
	Event
}

func (this RTCPeerConnectionIceErrorEvent) Once() RTCPeerConnectionIceErrorEvent {
	this.Ref().Once()
	return this
}

func (this RTCPeerConnectionIceErrorEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this RTCPeerConnectionIceErrorEvent) FromRef(ref js.Ref) RTCPeerConnectionIceErrorEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this RTCPeerConnectionIceErrorEvent) Free() {
	this.Ref().Free()
}

// Address returns the value of property "RTCPeerConnectionIceErrorEvent.address".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnectionIceErrorEvent) Address() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionIceErrorEventAddress(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Port returns the value of property "RTCPeerConnectionIceErrorEvent.port".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnectionIceErrorEvent) Port() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionIceErrorEventPort(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Url returns the value of property "RTCPeerConnectionIceErrorEvent.url".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnectionIceErrorEvent) Url() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionIceErrorEventUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ErrorCode returns the value of property "RTCPeerConnectionIceErrorEvent.errorCode".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnectionIceErrorEvent) ErrorCode() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionIceErrorEventErrorCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// ErrorText returns the value of property "RTCPeerConnectionIceErrorEvent.errorText".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnectionIceErrorEvent) ErrorText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionIceErrorEventErrorText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type RTCPeerConnectionIceEventInit struct {
	// Candidate is "RTCPeerConnectionIceEventInit.candidate"
	//
	// Optional
	Candidate RTCIceCandidate
	// Url is "RTCPeerConnectionIceEventInit.url"
	//
	// Optional
	Url js.String
	// Bubbles is "RTCPeerConnectionIceEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "RTCPeerConnectionIceEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "RTCPeerConnectionIceEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCPeerConnectionIceEventInit with all fields set.
func (p RTCPeerConnectionIceEventInit) FromRef(ref js.Ref) RTCPeerConnectionIceEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCPeerConnectionIceEventInit RTCPeerConnectionIceEventInit [// RTCPeerConnectionIceEventInit] [0x140009dafa0 0x140009db040 0x140009db0e0 0x140009db220 0x140009db360 0x140009db180 0x140009db2c0 0x140009db400] 0x1400099d440 {0 0}} in the application heap.
func (p RTCPeerConnectionIceEventInit) New() js.Ref {
	return bindings.RTCPeerConnectionIceEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCPeerConnectionIceEventInit) UpdateFrom(ref js.Ref) {
	bindings.RTCPeerConnectionIceEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCPeerConnectionIceEventInit) Update(ref js.Ref) {
	bindings.RTCPeerConnectionIceEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewRTCPeerConnectionIceEvent(typ js.String, eventInitDict RTCPeerConnectionIceEventInit) RTCPeerConnectionIceEvent {
	return RTCPeerConnectionIceEvent{}.FromRef(
		bindings.NewRTCPeerConnectionIceEventByRTCPeerConnectionIceEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewRTCPeerConnectionIceEventByRTCPeerConnectionIceEvent1(typ js.String) RTCPeerConnectionIceEvent {
	return RTCPeerConnectionIceEvent{}.FromRef(
		bindings.NewRTCPeerConnectionIceEventByRTCPeerConnectionIceEvent1(
			typ.Ref()),
	)
}

type RTCPeerConnectionIceEvent struct {
	Event
}

func (this RTCPeerConnectionIceEvent) Once() RTCPeerConnectionIceEvent {
	this.Ref().Once()
	return this
}

func (this RTCPeerConnectionIceEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this RTCPeerConnectionIceEvent) FromRef(ref js.Ref) RTCPeerConnectionIceEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this RTCPeerConnectionIceEvent) Free() {
	this.Ref().Free()
}

// Candidate returns the value of property "RTCPeerConnectionIceEvent.candidate".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnectionIceEvent) Candidate() (RTCIceCandidate, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionIceEventCandidate(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCIceCandidate{}.FromRef(_ret), _ok
}

// Url returns the value of property "RTCPeerConnectionIceEvent.url".
//
// The returned bool will be false if there is no such property.
func (this RTCPeerConnectionIceEvent) Url() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCPeerConnectionIceEventUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type RTCPeerConnectionStats struct {
	// DataChannelsOpened is "RTCPeerConnectionStats.dataChannelsOpened"
	//
	// Optional
	DataChannelsOpened uint32
	// DataChannelsClosed is "RTCPeerConnectionStats.dataChannelsClosed"
	//
	// Optional
	DataChannelsClosed uint32
	// Timestamp is "RTCPeerConnectionStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCPeerConnectionStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCPeerConnectionStats.id"
	//
	// Required
	Id js.String

	FFI_USE_DataChannelsOpened bool // for DataChannelsOpened.
	FFI_USE_DataChannelsClosed bool // for DataChannelsClosed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCPeerConnectionStats with all fields set.
func (p RTCPeerConnectionStats) FromRef(ref js.Ref) RTCPeerConnectionStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCPeerConnectionStats RTCPeerConnectionStats [// RTCPeerConnectionStats] [0x140009db540 0x140009db680 0x140009db7c0 0x140009db860 0x140009db900 0x140009db5e0 0x140009db720] 0x1400099d488 {0 0}} in the application heap.
func (p RTCPeerConnectionStats) New() js.Ref {
	return bindings.RTCPeerConnectionStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCPeerConnectionStats) UpdateFrom(ref js.Ref) {
	bindings.RTCPeerConnectionStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCPeerConnectionStats) Update(ref js.Ref) {
	bindings.RTCPeerConnectionStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCReceivedRtpStreamStats struct {
	// PacketsReceived is "RTCReceivedRtpStreamStats.packetsReceived"
	//
	// Optional
	PacketsReceived uint64
	// PacketsLost is "RTCReceivedRtpStreamStats.packetsLost"
	//
	// Optional
	PacketsLost int64
	// Jitter is "RTCReceivedRtpStreamStats.jitter"
	//
	// Optional
	Jitter float64
	// Ssrc is "RTCReceivedRtpStreamStats.ssrc"
	//
	// Required
	Ssrc uint32
	// Kind is "RTCReceivedRtpStreamStats.kind"
	//
	// Required
	Kind js.String
	// TransportId is "RTCReceivedRtpStreamStats.transportId"
	//
	// Optional
	TransportId js.String
	// CodecId is "RTCReceivedRtpStreamStats.codecId"
	//
	// Optional
	CodecId js.String
	// Timestamp is "RTCReceivedRtpStreamStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCReceivedRtpStreamStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCReceivedRtpStreamStats.id"
	//
	// Required
	Id js.String

	FFI_USE_PacketsReceived bool // for PacketsReceived.
	FFI_USE_PacketsLost     bool // for PacketsLost.
	FFI_USE_Jitter          bool // for Jitter.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCReceivedRtpStreamStats with all fields set.
func (p RTCReceivedRtpStreamStats) FromRef(ref js.Ref) RTCReceivedRtpStreamStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCReceivedRtpStreamStats RTCReceivedRtpStreamStats [// RTCReceivedRtpStreamStats] [0x140009db9a0 0x140009dbb80 0x140009dbcc0 0x140009dbe00 0x140009dbea0 0x140009dbf40 0x140009e8000 0x140009e80a0 0x140009e8140 0x140009e81e0 0x140009dba40 0x140009dbc20 0x140009dbd60] 0x1400099d4d0 {0 0}} in the application heap.
func (p RTCReceivedRtpStreamStats) New() js.Ref {
	return bindings.RTCReceivedRtpStreamStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCReceivedRtpStreamStats) UpdateFrom(ref js.Ref) {
	bindings.RTCReceivedRtpStreamStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCReceivedRtpStreamStats) Update(ref js.Ref) {
	bindings.RTCReceivedRtpStreamStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCRemoteInboundRtpStreamStats struct {
	// LocalId is "RTCRemoteInboundRtpStreamStats.localId"
	//
	// Optional
	LocalId js.String
	// RoundTripTime is "RTCRemoteInboundRtpStreamStats.roundTripTime"
	//
	// Optional
	RoundTripTime float64
	// TotalRoundTripTime is "RTCRemoteInboundRtpStreamStats.totalRoundTripTime"
	//
	// Optional
	TotalRoundTripTime float64
	// FractionLost is "RTCRemoteInboundRtpStreamStats.fractionLost"
	//
	// Optional
	FractionLost float64
	// RoundTripTimeMeasurements is "RTCRemoteInboundRtpStreamStats.roundTripTimeMeasurements"
	//
	// Optional
	RoundTripTimeMeasurements uint64
	// PacketsReceived is "RTCRemoteInboundRtpStreamStats.packetsReceived"
	//
	// Optional
	PacketsReceived uint64
	// PacketsLost is "RTCRemoteInboundRtpStreamStats.packetsLost"
	//
	// Optional
	PacketsLost int64
	// Jitter is "RTCRemoteInboundRtpStreamStats.jitter"
	//
	// Optional
	Jitter float64
	// Ssrc is "RTCRemoteInboundRtpStreamStats.ssrc"
	//
	// Required
	Ssrc uint32
	// Kind is "RTCRemoteInboundRtpStreamStats.kind"
	//
	// Required
	Kind js.String
	// TransportId is "RTCRemoteInboundRtpStreamStats.transportId"
	//
	// Optional
	TransportId js.String
	// CodecId is "RTCRemoteInboundRtpStreamStats.codecId"
	//
	// Optional
	CodecId js.String
	// Timestamp is "RTCRemoteInboundRtpStreamStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCRemoteInboundRtpStreamStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCRemoteInboundRtpStreamStats.id"
	//
	// Required
	Id js.String

	FFI_USE_RoundTripTime             bool // for RoundTripTime.
	FFI_USE_TotalRoundTripTime        bool // for TotalRoundTripTime.
	FFI_USE_FractionLost              bool // for FractionLost.
	FFI_USE_RoundTripTimeMeasurements bool // for RoundTripTimeMeasurements.
	FFI_USE_PacketsReceived           bool // for PacketsReceived.
	FFI_USE_PacketsLost               bool // for PacketsLost.
	FFI_USE_Jitter                    bool // for Jitter.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRemoteInboundRtpStreamStats with all fields set.
func (p RTCRemoteInboundRtpStreamStats) FromRef(ref js.Ref) RTCRemoteInboundRtpStreamStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCRemoteInboundRtpStreamStats RTCRemoteInboundRtpStreamStats [// RTCRemoteInboundRtpStreamStats] [0x140009e8280 0x140009e8320 0x140009e8460 0x140009e85a0 0x140009e86e0 0x140009e8820 0x140009e8960 0x140009e8aa0 0x140009e8be0 0x140009e8c80 0x140009e8d20 0x140009e8dc0 0x140009e8e60 0x140009e8f00 0x140009e8fa0 0x140009e83c0 0x140009e8500 0x140009e8640 0x140009e8780 0x140009e88c0 0x140009e8a00 0x140009e8b40] 0x1400099d578 {0 0}} in the application heap.
func (p RTCRemoteInboundRtpStreamStats) New() js.Ref {
	return bindings.RTCRemoteInboundRtpStreamStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRemoteInboundRtpStreamStats) UpdateFrom(ref js.Ref) {
	bindings.RTCRemoteInboundRtpStreamStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRemoteInboundRtpStreamStats) Update(ref js.Ref) {
	bindings.RTCRemoteInboundRtpStreamStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCRemoteOutboundRtpStreamStats struct {
	// LocalId is "RTCRemoteOutboundRtpStreamStats.localId"
	//
	// Optional
	LocalId js.String
	// RemoteTimestamp is "RTCRemoteOutboundRtpStreamStats.remoteTimestamp"
	//
	// Optional
	RemoteTimestamp DOMHighResTimeStamp
	// ReportsSent is "RTCRemoteOutboundRtpStreamStats.reportsSent"
	//
	// Optional
	ReportsSent uint64
	// RoundTripTime is "RTCRemoteOutboundRtpStreamStats.roundTripTime"
	//
	// Optional
	RoundTripTime float64
	// TotalRoundTripTime is "RTCRemoteOutboundRtpStreamStats.totalRoundTripTime"
	//
	// Optional
	TotalRoundTripTime float64
	// RoundTripTimeMeasurements is "RTCRemoteOutboundRtpStreamStats.roundTripTimeMeasurements"
	//
	// Optional
	RoundTripTimeMeasurements uint64
	// PacketsSent is "RTCRemoteOutboundRtpStreamStats.packetsSent"
	//
	// Optional
	PacketsSent uint64
	// BytesSent is "RTCRemoteOutboundRtpStreamStats.bytesSent"
	//
	// Optional
	BytesSent uint64
	// Ssrc is "RTCRemoteOutboundRtpStreamStats.ssrc"
	//
	// Required
	Ssrc uint32
	// Kind is "RTCRemoteOutboundRtpStreamStats.kind"
	//
	// Required
	Kind js.String
	// TransportId is "RTCRemoteOutboundRtpStreamStats.transportId"
	//
	// Optional
	TransportId js.String
	// CodecId is "RTCRemoteOutboundRtpStreamStats.codecId"
	//
	// Optional
	CodecId js.String
	// Timestamp is "RTCRemoteOutboundRtpStreamStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCRemoteOutboundRtpStreamStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCRemoteOutboundRtpStreamStats.id"
	//
	// Required
	Id js.String

	FFI_USE_RemoteTimestamp           bool // for RemoteTimestamp.
	FFI_USE_ReportsSent               bool // for ReportsSent.
	FFI_USE_RoundTripTime             bool // for RoundTripTime.
	FFI_USE_TotalRoundTripTime        bool // for TotalRoundTripTime.
	FFI_USE_RoundTripTimeMeasurements bool // for RoundTripTimeMeasurements.
	FFI_USE_PacketsSent               bool // for PacketsSent.
	FFI_USE_BytesSent                 bool // for BytesSent.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRemoteOutboundRtpStreamStats with all fields set.
func (p RTCRemoteOutboundRtpStreamStats) FromRef(ref js.Ref) RTCRemoteOutboundRtpStreamStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCRemoteOutboundRtpStreamStats RTCRemoteOutboundRtpStreamStats [// RTCRemoteOutboundRtpStreamStats] [0x140009e9040 0x140009e90e0 0x140009e9220 0x140009e9360 0x140009e94a0 0x140009e95e0 0x140009e9720 0x140009e9860 0x140009e99a0 0x140009e9a40 0x140009e9ae0 0x140009e9b80 0x140009e9c20 0x140009e9cc0 0x140009e9d60 0x140009e9180 0x140009e92c0 0x140009e9400 0x140009e9540 0x140009e9680 0x140009e97c0 0x140009e9900] 0x1400099d650 {0 0}} in the application heap.
func (p RTCRemoteOutboundRtpStreamStats) New() js.Ref {
	return bindings.RTCRemoteOutboundRtpStreamStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRemoteOutboundRtpStreamStats) UpdateFrom(ref js.Ref) {
	bindings.RTCRemoteOutboundRtpStreamStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRemoteOutboundRtpStreamStats) Update(ref js.Ref) {
	bindings.RTCRemoteOutboundRtpStreamStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCRtpCodec struct {
	// MimeType is "RTCRtpCodec.mimeType"
	//
	// Required
	MimeType js.String
	// ClockRate is "RTCRtpCodec.clockRate"
	//
	// Required
	ClockRate uint32
	// Channels is "RTCRtpCodec.channels"
	//
	// Optional
	Channels uint16
	// SdpFmtpLine is "RTCRtpCodec.sdpFmtpLine"
	//
	// Optional
	SdpFmtpLine js.String

	FFI_USE_Channels bool // for Channels.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpCodec with all fields set.
func (p RTCRtpCodec) FromRef(ref js.Ref) RTCRtpCodec {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCRtpCodec RTCRtpCodec [// RTCRtpCodec] [0x140009e9e00 0x140009e9ea0 0x140009e9f40 0x140009ea0a0 0x140009ea000] 0x1400099d740 {0 0}} in the application heap.
func (p RTCRtpCodec) New() js.Ref {
	return bindings.RTCRtpCodecJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtpCodec) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpCodecJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtpCodec) Update(ref js.Ref) {
	bindings.RTCRtpCodecJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCRtpCodingParameters struct {
	// Rid is "RTCRtpCodingParameters.rid"
	//
	// Optional
	Rid js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpCodingParameters with all fields set.
func (p RTCRtpCodingParameters) FromRef(ref js.Ref) RTCRtpCodingParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCRtpCodingParameters RTCRtpCodingParameters [// RTCRtpCodingParameters] [0x140009ea140] 0x1400099d758 {0 0}} in the application heap.
func (p RTCRtpCodingParameters) New() js.Ref {
	return bindings.RTCRtpCodingParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtpCodingParameters) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpCodingParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtpCodingParameters) Update(ref js.Ref) {
	bindings.RTCRtpCodingParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCRtpParameters struct {
	// HeaderExtensions is "RTCRtpParameters.headerExtensions"
	//
	// Required
	HeaderExtensions js.Array[RTCRtpHeaderExtensionParameters]
	// Rtcp is "RTCRtpParameters.rtcp"
	//
	// Required
	Rtcp RTCRtcpParameters
	// Codecs is "RTCRtpParameters.codecs"
	//
	// Required
	Codecs js.Array[RTCRtpCodecParameters]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpParameters with all fields set.
func (p RTCRtpParameters) FromRef(ref js.Ref) RTCRtpParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCRtpParameters RTCRtpParameters [// RTCRtpParameters] [0x140009ea1e0 0x140009ea280 0x140009ea320] 0x1400099d7a0 {0 0}} in the application heap.
func (p RTCRtpParameters) New() js.Ref {
	return bindings.RTCRtpParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtpParameters) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtpParameters) Update(ref js.Ref) {
	bindings.RTCRtpParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCRtpScriptTransformer struct {
	ref js.Ref
}

func (this RTCRtpScriptTransformer) Once() RTCRtpScriptTransformer {
	this.Ref().Once()
	return this
}

func (this RTCRtpScriptTransformer) Ref() js.Ref {
	return this.ref
}

func (this RTCRtpScriptTransformer) FromRef(ref js.Ref) RTCRtpScriptTransformer {
	this.ref = ref
	return this
}

func (this RTCRtpScriptTransformer) Free() {
	this.Ref().Free()
}

// Readable returns the value of property "RTCRtpScriptTransformer.readable".
//
// The returned bool will be false if there is no such property.
func (this RTCRtpScriptTransformer) Readable() (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.GetRTCRtpScriptTransformerReadable(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadableStream{}.FromRef(_ret), _ok
}

// Writable returns the value of property "RTCRtpScriptTransformer.writable".
//
// The returned bool will be false if there is no such property.
func (this RTCRtpScriptTransformer) Writable() (WritableStream, bool) {
	var _ok bool
	_ret := bindings.GetRTCRtpScriptTransformerWritable(
		this.Ref(), js.Pointer(&_ok),
	)
	return WritableStream{}.FromRef(_ret), _ok
}

// Options returns the value of property "RTCRtpScriptTransformer.options".
//
// The returned bool will be false if there is no such property.
func (this RTCRtpScriptTransformer) Options() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetRTCRtpScriptTransformerOptions(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// GenerateKeyFrame calls the method "RTCRtpScriptTransformer.generateKeyFrame".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpScriptTransformer) GenerateKeyFrame(rid js.String) (js.Promise[js.BigInt[uint64]], bool) {
	var _ok bool
	_ret := bindings.CallRTCRtpScriptTransformerGenerateKeyFrame(
		this.Ref(), js.Pointer(&_ok),
		rid.Ref(),
	)

	return js.Promise[js.BigInt[uint64]]{}.FromRef(_ret), _ok
}

// GenerateKeyFrameFunc returns the method "RTCRtpScriptTransformer.generateKeyFrame".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpScriptTransformer) GenerateKeyFrameFunc() (fn js.Func[func(rid js.String) js.Promise[js.BigInt[uint64]]]) {
	return fn.FromRef(
		bindings.RTCRtpScriptTransformerGenerateKeyFrameFunc(
			this.Ref(),
		),
	)
}

// GenerateKeyFrame1 calls the method "RTCRtpScriptTransformer.generateKeyFrame".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpScriptTransformer) GenerateKeyFrame1() (js.Promise[js.BigInt[uint64]], bool) {
	var _ok bool
	_ret := bindings.CallRTCRtpScriptTransformerGenerateKeyFrame1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.BigInt[uint64]]{}.FromRef(_ret), _ok
}

// GenerateKeyFrame1Func returns the method "RTCRtpScriptTransformer.generateKeyFrame".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpScriptTransformer) GenerateKeyFrame1Func() (fn js.Func[func() js.Promise[js.BigInt[uint64]]]) {
	return fn.FromRef(
		bindings.RTCRtpScriptTransformerGenerateKeyFrame1Func(
			this.Ref(),
		),
	)
}

// SendKeyFrameRequest calls the method "RTCRtpScriptTransformer.sendKeyFrameRequest".
//
// The returned bool will be false if there is no such method.
func (this RTCRtpScriptTransformer) SendKeyFrameRequest() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRTCRtpScriptTransformerSendKeyFrameRequest(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SendKeyFrameRequestFunc returns the method "RTCRtpScriptTransformer.sendKeyFrameRequest".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCRtpScriptTransformer) SendKeyFrameRequestFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCRtpScriptTransformerSendKeyFrameRequestFunc(
			this.Ref(),
		),
	)
}

type RTCRtpStreamStats struct {
	// Ssrc is "RTCRtpStreamStats.ssrc"
	//
	// Required
	Ssrc uint32
	// Kind is "RTCRtpStreamStats.kind"
	//
	// Required
	Kind js.String
	// TransportId is "RTCRtpStreamStats.transportId"
	//
	// Optional
	TransportId js.String
	// CodecId is "RTCRtpStreamStats.codecId"
	//
	// Optional
	CodecId js.String
	// Timestamp is "RTCRtpStreamStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCRtpStreamStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCRtpStreamStats.id"
	//
	// Required
	Id js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpStreamStats with all fields set.
func (p RTCRtpStreamStats) FromRef(ref js.Ref) RTCRtpStreamStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCRtpStreamStats RTCRtpStreamStats [// RTCRtpStreamStats] [0x140009ea460 0x140009ea500 0x140009ea5a0 0x140009ea640 0x140009ea6e0 0x140009ea780 0x140009ea820] 0x1400099d800 {0 0}} in the application heap.
func (p RTCRtpStreamStats) New() js.Ref {
	return bindings.RTCRtpStreamStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtpStreamStats) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpStreamStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtpStreamStats) Update(ref js.Ref) {
	bindings.RTCRtpStreamStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCSentRtpStreamStats struct {
	// PacketsSent is "RTCSentRtpStreamStats.packetsSent"
	//
	// Optional
	PacketsSent uint64
	// BytesSent is "RTCSentRtpStreamStats.bytesSent"
	//
	// Optional
	BytesSent uint64
	// Ssrc is "RTCSentRtpStreamStats.ssrc"
	//
	// Required
	Ssrc uint32
	// Kind is "RTCSentRtpStreamStats.kind"
	//
	// Required
	Kind js.String
	// TransportId is "RTCSentRtpStreamStats.transportId"
	//
	// Optional
	TransportId js.String
	// CodecId is "RTCSentRtpStreamStats.codecId"
	//
	// Optional
	CodecId js.String
	// Timestamp is "RTCSentRtpStreamStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCSentRtpStreamStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCSentRtpStreamStats.id"
	//
	// Required
	Id js.String

	FFI_USE_PacketsSent bool // for PacketsSent.
	FFI_USE_BytesSent   bool // for BytesSent.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCSentRtpStreamStats with all fields set.
func (p RTCSentRtpStreamStats) FromRef(ref js.Ref) RTCSentRtpStreamStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCSentRtpStreamStats RTCSentRtpStreamStats [// RTCSentRtpStreamStats] [0x140009ea8c0 0x140009eaa00 0x140009eab40 0x140009eabe0 0x140009eac80 0x140009ead20 0x140009eadc0 0x140009eae60 0x140009eaf00 0x140009ea960 0x140009eaaa0] 0x1400099d890 {0 0}} in the application heap.
func (p RTCSentRtpStreamStats) New() js.Ref {
	return bindings.RTCSentRtpStreamStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCSentRtpStreamStats) UpdateFrom(ref js.Ref) {
	bindings.RTCSentRtpStreamStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCSentRtpStreamStats) Update(ref js.Ref) {
	bindings.RTCSentRtpStreamStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCStats struct {
	// Timestamp is "RTCStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCStats.id"
	//
	// Required
	Id js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCStats with all fields set.
func (p RTCStats) FromRef(ref js.Ref) RTCStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCStats RTCStats [// RTCStats] [0x140009eafa0 0x140009eb040 0x140009eb0e0] 0x1400099d938 {0 0}} in the application heap.
func (p RTCStats) New() js.Ref {
	return bindings.RTCStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCStats) UpdateFrom(ref js.Ref) {
	bindings.RTCStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCStats) Update(ref js.Ref) {
	bindings.RTCStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}
