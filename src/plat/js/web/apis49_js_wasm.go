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
	//
	// NOTE: FFI_USE_Active MUST be set to true to make this field effective.
	Active bool
	// MaxBitrate is "RTCRtpEncodingParameters.maxBitrate"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxBitrate MUST be set to true to make this field effective.
	MaxBitrate uint32
	// MaxFramerate is "RTCRtpEncodingParameters.maxFramerate"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxFramerate MUST be set to true to make this field effective.
	MaxFramerate float64
	// ScaleResolutionDownBy is "RTCRtpEncodingParameters.scaleResolutionDownBy"
	//
	// Optional
	//
	// NOTE: FFI_USE_ScaleResolutionDownBy MUST be set to true to make this field effective.
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

// New creates a new RTCRtpEncodingParameters in the application heap.
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
	//
	// NOTE: FFI_USE_Encrypted MUST be set to true to make this field effective.
	Encrypted bool

	FFI_USE_Encrypted bool // for Encrypted.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpHeaderExtensionParameters with all fields set.
func (p RTCRtpHeaderExtensionParameters) FromRef(ref js.Ref) RTCRtpHeaderExtensionParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCRtpHeaderExtensionParameters in the application heap.
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
	//
	// NOTE: FFI_USE_ReducedSize MUST be set to true to make this field effective.
	ReducedSize bool

	FFI_USE_ReducedSize bool // for ReducedSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtcpParameters with all fields set.
func (p RTCRtcpParameters) FromRef(ref js.Ref) RTCRtcpParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCRtcpParameters in the application heap.
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
	//
	// NOTE: FFI_USE_Channels MUST be set to true to make this field effective.
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

// New creates a new RTCRtpCodecParameters in the application heap.
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
	//
	// NOTE: Rtcp.FFI_USE MUST be set to true to get Rtcp used.
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

// New creates a new RTCRtpSendParameters in the application heap.
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

// New creates a new RTCSetParameterOptions in the application heap.
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

// New creates a new SFrameTransformOptions in the application heap.
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

func NewSFrameTransform(options SFrameTransformOptions) (ret SFrameTransform) {
	ret.ref = bindings.NewSFrameTransformBySFrameTransform(
		js.Pointer(&options))
	return
}

func NewSFrameTransformBySFrameTransform1() (ret SFrameTransform) {
	ret.ref = bindings.NewSFrameTransformBySFrameTransform1()
	return
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
// It returns ok=false if there is no such property.
func (this SFrameTransform) Readable() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetSFrameTransformReadable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Writable returns the value of property "SFrameTransform.writable".
//
// It returns ok=false if there is no such property.
func (this SFrameTransform) Writable() (ret WritableStream, ok bool) {
	ok = js.True == bindings.GetSFrameTransformWritable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasSetEncryptionKey returns true if the method "SFrameTransform.setEncryptionKey" exists.
func (this SFrameTransform) HasSetEncryptionKey() bool {
	return js.True == bindings.HasSFrameTransformSetEncryptionKey(
		this.Ref(),
	)
}

// SetEncryptionKeyFunc returns the method "SFrameTransform.setEncryptionKey".
func (this SFrameTransform) SetEncryptionKeyFunc() (fn js.Func[func(key CryptoKey, keyID CryptoKeyID) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SFrameTransformSetEncryptionKeyFunc(
			this.Ref(),
		),
	)
}

// SetEncryptionKey calls the method "SFrameTransform.setEncryptionKey".
func (this SFrameTransform) SetEncryptionKey(key CryptoKey, keyID CryptoKeyID) (ret js.Promise[js.Void]) {
	bindings.CallSFrameTransformSetEncryptionKey(
		this.Ref(), js.Pointer(&ret),
		key.Ref(),
		keyID.Ref(),
	)

	return
}

// TrySetEncryptionKey calls the method "SFrameTransform.setEncryptionKey"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SFrameTransform) TrySetEncryptionKey(key CryptoKey, keyID CryptoKeyID) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySFrameTransformSetEncryptionKey(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
		keyID.Ref(),
	)

	return
}

// HasSetEncryptionKey1 returns true if the method "SFrameTransform.setEncryptionKey" exists.
func (this SFrameTransform) HasSetEncryptionKey1() bool {
	return js.True == bindings.HasSFrameTransformSetEncryptionKey1(
		this.Ref(),
	)
}

// SetEncryptionKey1Func returns the method "SFrameTransform.setEncryptionKey".
func (this SFrameTransform) SetEncryptionKey1Func() (fn js.Func[func(key CryptoKey) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SFrameTransformSetEncryptionKey1Func(
			this.Ref(),
		),
	)
}

// SetEncryptionKey1 calls the method "SFrameTransform.setEncryptionKey".
func (this SFrameTransform) SetEncryptionKey1(key CryptoKey) (ret js.Promise[js.Void]) {
	bindings.CallSFrameTransformSetEncryptionKey1(
		this.Ref(), js.Pointer(&ret),
		key.Ref(),
	)

	return
}

// TrySetEncryptionKey1 calls the method "SFrameTransform.setEncryptionKey"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SFrameTransform) TrySetEncryptionKey1(key CryptoKey) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySFrameTransformSetEncryptionKey1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
	)

	return
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

// New creates a new WorkerOptions in the application heap.
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

func NewWorker(scriptURL js.String, options WorkerOptions) (ret Worker) {
	ret.ref = bindings.NewWorkerByWorker(
		scriptURL.Ref(),
		js.Pointer(&options))
	return
}

func NewWorkerByWorker1(scriptURL js.String) (ret Worker) {
	ret.ref = bindings.NewWorkerByWorker1(
		scriptURL.Ref())
	return
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

// HasTerminate returns true if the method "Worker.terminate" exists.
func (this Worker) HasTerminate() bool {
	return js.True == bindings.HasWorkerTerminate(
		this.Ref(),
	)
}

// TerminateFunc returns the method "Worker.terminate".
func (this Worker) TerminateFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WorkerTerminateFunc(
			this.Ref(),
		),
	)
}

// Terminate calls the method "Worker.terminate".
func (this Worker) Terminate() (ret js.Void) {
	bindings.CallWorkerTerminate(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTerminate calls the method "Worker.terminate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Worker) TryTerminate() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerTerminate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPostMessage returns true if the method "Worker.postMessage" exists.
func (this Worker) HasPostMessage() bool {
	return js.True == bindings.HasWorkerPostMessage(
		this.Ref(),
	)
}

// PostMessageFunc returns the method "Worker.postMessage".
func (this Worker) PostMessageFunc() (fn js.Func[func(message js.Any, transfer js.Array[js.Object])]) {
	return fn.FromRef(
		bindings.WorkerPostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage calls the method "Worker.postMessage".
func (this Worker) PostMessage(message js.Any, transfer js.Array[js.Object]) (ret js.Void) {
	bindings.CallWorkerPostMessage(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		transfer.Ref(),
	)

	return
}

// TryPostMessage calls the method "Worker.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Worker) TryPostMessage(message js.Any, transfer js.Array[js.Object]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerPostMessage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		transfer.Ref(),
	)

	return
}

// HasPostMessage1 returns true if the method "Worker.postMessage" exists.
func (this Worker) HasPostMessage1() bool {
	return js.True == bindings.HasWorkerPostMessage1(
		this.Ref(),
	)
}

// PostMessage1Func returns the method "Worker.postMessage".
func (this Worker) PostMessage1Func() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	return fn.FromRef(
		bindings.WorkerPostMessage1Func(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "Worker.postMessage".
func (this Worker) PostMessage1(message js.Any, options StructuredSerializeOptions) (ret js.Void) {
	bindings.CallWorkerPostMessage1(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPostMessage1 calls the method "Worker.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Worker) TryPostMessage1(message js.Any, options StructuredSerializeOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerPostMessage1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasPostMessage2 returns true if the method "Worker.postMessage" exists.
func (this Worker) HasPostMessage2() bool {
	return js.True == bindings.HasWorkerPostMessage2(
		this.Ref(),
	)
}

// PostMessage2Func returns the method "Worker.postMessage".
func (this Worker) PostMessage2Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.WorkerPostMessage2Func(
			this.Ref(),
		),
	)
}

// PostMessage2 calls the method "Worker.postMessage".
func (this Worker) PostMessage2(message js.Any) (ret js.Void) {
	bindings.CallWorkerPostMessage2(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPostMessage2 calls the method "Worker.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Worker) TryPostMessage2(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerPostMessage2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

func NewRTCRtpScriptTransform(worker Worker, options js.Any, transfer js.Array[js.Object]) (ret RTCRtpScriptTransform) {
	ret.ref = bindings.NewRTCRtpScriptTransformByRTCRtpScriptTransform(
		worker.Ref(),
		options.Ref(),
		transfer.Ref())
	return
}

func NewRTCRtpScriptTransformByRTCRtpScriptTransform1(worker Worker, options js.Any) (ret RTCRtpScriptTransform) {
	ret.ref = bindings.NewRTCRtpScriptTransformByRTCRtpScriptTransform1(
		worker.Ref(),
		options.Ref())
	return
}

func NewRTCRtpScriptTransformByRTCRtpScriptTransform2(worker Worker) (ret RTCRtpScriptTransform) {
	ret.ref = bindings.NewRTCRtpScriptTransformByRTCRtpScriptTransform2(
		worker.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this RTCRtpSender) Track() (ret MediaStreamTrack, ok bool) {
	ok = js.True == bindings.GetRTCRtpSenderTrack(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Transport returns the value of property "RTCRtpSender.transport".
//
// It returns ok=false if there is no such property.
func (this RTCRtpSender) Transport() (ret RTCDtlsTransport, ok bool) {
	ok = js.True == bindings.GetRTCRtpSenderTransport(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Dtmf returns the value of property "RTCRtpSender.dtmf".
//
// It returns ok=false if there is no such property.
func (this RTCRtpSender) Dtmf() (ret RTCDTMFSender, ok bool) {
	ok = js.True == bindings.GetRTCRtpSenderDtmf(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Transform returns the value of property "RTCRtpSender.transform".
//
// It returns ok=false if there is no such property.
func (this RTCRtpSender) Transform() (ret RTCRtpTransform, ok bool) {
	ok = js.True == bindings.GetRTCRtpSenderTransform(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTransform sets the value of property "RTCRtpSender.transform" to val.
//
// It returns false if the property cannot be set.
func (this RTCRtpSender) SetTransform(val RTCRtpTransform) bool {
	return js.True == bindings.SetRTCRtpSenderTransform(
		this.Ref(),
		val.Ref(),
	)
}

// HasGetCapabilities returns true if the staticmethod "RTCRtpSender.getCapabilities" exists.
func (this RTCRtpSender) HasGetCapabilities() bool {
	return js.True == bindings.HasRTCRtpSenderGetCapabilities(
		this.Ref(),
	)
}

// GetCapabilitiesFunc returns the staticmethod "RTCRtpSender.getCapabilities".
func (this RTCRtpSender) GetCapabilitiesFunc() (fn js.Func[func(kind js.String) RTCRtpCapabilities]) {
	return fn.FromRef(
		bindings.RTCRtpSenderGetCapabilitiesFunc(
			this.Ref(),
		),
	)
}

// GetCapabilities calls the staticmethod "RTCRtpSender.getCapabilities".
func (this RTCRtpSender) GetCapabilities(kind js.String) (ret RTCRtpCapabilities) {
	bindings.CallRTCRtpSenderGetCapabilities(
		this.Ref(), js.Pointer(&ret),
		kind.Ref(),
	)

	return
}

// TryGetCapabilities calls the staticmethod "RTCRtpSender.getCapabilities"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpSender) TryGetCapabilities(kind js.String) (ret RTCRtpCapabilities, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpSenderGetCapabilities(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		kind.Ref(),
	)

	return
}

// HasSetParameters returns true if the method "RTCRtpSender.setParameters" exists.
func (this RTCRtpSender) HasSetParameters() bool {
	return js.True == bindings.HasRTCRtpSenderSetParameters(
		this.Ref(),
	)
}

// SetParametersFunc returns the method "RTCRtpSender.setParameters".
func (this RTCRtpSender) SetParametersFunc() (fn js.Func[func(parameters RTCRtpSendParameters, setParameterOptions RTCSetParameterOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCRtpSenderSetParametersFunc(
			this.Ref(),
		),
	)
}

// SetParameters calls the method "RTCRtpSender.setParameters".
func (this RTCRtpSender) SetParameters(parameters RTCRtpSendParameters, setParameterOptions RTCSetParameterOptions) (ret js.Promise[js.Void]) {
	bindings.CallRTCRtpSenderSetParameters(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&parameters),
		js.Pointer(&setParameterOptions),
	)

	return
}

// TrySetParameters calls the method "RTCRtpSender.setParameters"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpSender) TrySetParameters(parameters RTCRtpSendParameters, setParameterOptions RTCSetParameterOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpSenderSetParameters(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
		js.Pointer(&setParameterOptions),
	)

	return
}

// HasSetParameters1 returns true if the method "RTCRtpSender.setParameters" exists.
func (this RTCRtpSender) HasSetParameters1() bool {
	return js.True == bindings.HasRTCRtpSenderSetParameters1(
		this.Ref(),
	)
}

// SetParameters1Func returns the method "RTCRtpSender.setParameters".
func (this RTCRtpSender) SetParameters1Func() (fn js.Func[func(parameters RTCRtpSendParameters) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCRtpSenderSetParameters1Func(
			this.Ref(),
		),
	)
}

// SetParameters1 calls the method "RTCRtpSender.setParameters".
func (this RTCRtpSender) SetParameters1(parameters RTCRtpSendParameters) (ret js.Promise[js.Void]) {
	bindings.CallRTCRtpSenderSetParameters1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TrySetParameters1 calls the method "RTCRtpSender.setParameters"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpSender) TrySetParameters1(parameters RTCRtpSendParameters) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpSenderSetParameters1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasGetParameters returns true if the method "RTCRtpSender.getParameters" exists.
func (this RTCRtpSender) HasGetParameters() bool {
	return js.True == bindings.HasRTCRtpSenderGetParameters(
		this.Ref(),
	)
}

// GetParametersFunc returns the method "RTCRtpSender.getParameters".
func (this RTCRtpSender) GetParametersFunc() (fn js.Func[func() RTCRtpSendParameters]) {
	return fn.FromRef(
		bindings.RTCRtpSenderGetParametersFunc(
			this.Ref(),
		),
	)
}

// GetParameters calls the method "RTCRtpSender.getParameters".
func (this RTCRtpSender) GetParameters() (ret RTCRtpSendParameters) {
	bindings.CallRTCRtpSenderGetParameters(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetParameters calls the method "RTCRtpSender.getParameters"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpSender) TryGetParameters() (ret RTCRtpSendParameters, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpSenderGetParameters(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReplaceTrack returns true if the method "RTCRtpSender.replaceTrack" exists.
func (this RTCRtpSender) HasReplaceTrack() bool {
	return js.True == bindings.HasRTCRtpSenderReplaceTrack(
		this.Ref(),
	)
}

// ReplaceTrackFunc returns the method "RTCRtpSender.replaceTrack".
func (this RTCRtpSender) ReplaceTrackFunc() (fn js.Func[func(withTrack MediaStreamTrack) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCRtpSenderReplaceTrackFunc(
			this.Ref(),
		),
	)
}

// ReplaceTrack calls the method "RTCRtpSender.replaceTrack".
func (this RTCRtpSender) ReplaceTrack(withTrack MediaStreamTrack) (ret js.Promise[js.Void]) {
	bindings.CallRTCRtpSenderReplaceTrack(
		this.Ref(), js.Pointer(&ret),
		withTrack.Ref(),
	)

	return
}

// TryReplaceTrack calls the method "RTCRtpSender.replaceTrack"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpSender) TryReplaceTrack(withTrack MediaStreamTrack) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpSenderReplaceTrack(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		withTrack.Ref(),
	)

	return
}

// HasSetStreams returns true if the method "RTCRtpSender.setStreams" exists.
func (this RTCRtpSender) HasSetStreams() bool {
	return js.True == bindings.HasRTCRtpSenderSetStreams(
		this.Ref(),
	)
}

// SetStreamsFunc returns the method "RTCRtpSender.setStreams".
func (this RTCRtpSender) SetStreamsFunc() (fn js.Func[func(streams ...MediaStream)]) {
	return fn.FromRef(
		bindings.RTCRtpSenderSetStreamsFunc(
			this.Ref(),
		),
	)
}

// SetStreams calls the method "RTCRtpSender.setStreams".
func (this RTCRtpSender) SetStreams(streams ...MediaStream) (ret js.Void) {
	bindings.CallRTCRtpSenderSetStreams(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(streams),
		js.SizeU(len(streams)),
	)

	return
}

// TrySetStreams calls the method "RTCRtpSender.setStreams"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpSender) TrySetStreams(streams ...MediaStream) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpSenderSetStreams(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(streams),
		js.SizeU(len(streams)),
	)

	return
}

// HasGetStats returns true if the method "RTCRtpSender.getStats" exists.
func (this RTCRtpSender) HasGetStats() bool {
	return js.True == bindings.HasRTCRtpSenderGetStats(
		this.Ref(),
	)
}

// GetStatsFunc returns the method "RTCRtpSender.getStats".
func (this RTCRtpSender) GetStatsFunc() (fn js.Func[func() js.Promise[RTCStatsReport]]) {
	return fn.FromRef(
		bindings.RTCRtpSenderGetStatsFunc(
			this.Ref(),
		),
	)
}

// GetStats calls the method "RTCRtpSender.getStats".
func (this RTCRtpSender) GetStats() (ret js.Promise[RTCStatsReport]) {
	bindings.CallRTCRtpSenderGetStats(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetStats calls the method "RTCRtpSender.getStats"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpSender) TryGetStats() (ret js.Promise[RTCStatsReport], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpSenderGetStats(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGenerateKeyFrame returns true if the method "RTCRtpSender.generateKeyFrame" exists.
func (this RTCRtpSender) HasGenerateKeyFrame() bool {
	return js.True == bindings.HasRTCRtpSenderGenerateKeyFrame(
		this.Ref(),
	)
}

// GenerateKeyFrameFunc returns the method "RTCRtpSender.generateKeyFrame".
func (this RTCRtpSender) GenerateKeyFrameFunc() (fn js.Func[func(rids js.Array[js.String]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCRtpSenderGenerateKeyFrameFunc(
			this.Ref(),
		),
	)
}

// GenerateKeyFrame calls the method "RTCRtpSender.generateKeyFrame".
func (this RTCRtpSender) GenerateKeyFrame(rids js.Array[js.String]) (ret js.Promise[js.Void]) {
	bindings.CallRTCRtpSenderGenerateKeyFrame(
		this.Ref(), js.Pointer(&ret),
		rids.Ref(),
	)

	return
}

// TryGenerateKeyFrame calls the method "RTCRtpSender.generateKeyFrame"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpSender) TryGenerateKeyFrame(rids js.Array[js.String]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpSenderGenerateKeyFrame(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rids.Ref(),
	)

	return
}

// HasGenerateKeyFrame1 returns true if the method "RTCRtpSender.generateKeyFrame" exists.
func (this RTCRtpSender) HasGenerateKeyFrame1() bool {
	return js.True == bindings.HasRTCRtpSenderGenerateKeyFrame1(
		this.Ref(),
	)
}

// GenerateKeyFrame1Func returns the method "RTCRtpSender.generateKeyFrame".
func (this RTCRtpSender) GenerateKeyFrame1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCRtpSenderGenerateKeyFrame1Func(
			this.Ref(),
		),
	)
}

// GenerateKeyFrame1 calls the method "RTCRtpSender.generateKeyFrame".
func (this RTCRtpSender) GenerateKeyFrame1() (ret js.Promise[js.Void]) {
	bindings.CallRTCRtpSenderGenerateKeyFrame1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGenerateKeyFrame1 calls the method "RTCRtpSender.generateKeyFrame"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpSender) TryGenerateKeyFrame1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpSenderGenerateKeyFrame1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type RTCRtpReceiveParameters struct {
	// HeaderExtensions is "RTCRtpReceiveParameters.headerExtensions"
	//
	// Required
	HeaderExtensions js.Array[RTCRtpHeaderExtensionParameters]
	// Rtcp is "RTCRtpReceiveParameters.rtcp"
	//
	// Required
	//
	// NOTE: Rtcp.FFI_USE MUST be set to true to get Rtcp used.
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

// New creates a new RTCRtpReceiveParameters in the application heap.
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
	//
	// NOTE: FFI_USE_AudioLevel MUST be set to true to make this field effective.
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

// New creates a new RTCRtpContributingSource in the application heap.
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
	//
	// NOTE: FFI_USE_AudioLevel MUST be set to true to make this field effective.
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

// New creates a new RTCRtpSynchronizationSource in the application heap.
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
// It returns ok=false if there is no such property.
func (this RTCRtpReceiver) Track() (ret MediaStreamTrack, ok bool) {
	ok = js.True == bindings.GetRTCRtpReceiverTrack(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Transport returns the value of property "RTCRtpReceiver.transport".
//
// It returns ok=false if there is no such property.
func (this RTCRtpReceiver) Transport() (ret RTCDtlsTransport, ok bool) {
	ok = js.True == bindings.GetRTCRtpReceiverTransport(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Transform returns the value of property "RTCRtpReceiver.transform".
//
// It returns ok=false if there is no such property.
func (this RTCRtpReceiver) Transform() (ret RTCRtpTransform, ok bool) {
	ok = js.True == bindings.GetRTCRtpReceiverTransform(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTransform sets the value of property "RTCRtpReceiver.transform" to val.
//
// It returns false if the property cannot be set.
func (this RTCRtpReceiver) SetTransform(val RTCRtpTransform) bool {
	return js.True == bindings.SetRTCRtpReceiverTransform(
		this.Ref(),
		val.Ref(),
	)
}

// HasGetCapabilities returns true if the staticmethod "RTCRtpReceiver.getCapabilities" exists.
func (this RTCRtpReceiver) HasGetCapabilities() bool {
	return js.True == bindings.HasRTCRtpReceiverGetCapabilities(
		this.Ref(),
	)
}

// GetCapabilitiesFunc returns the staticmethod "RTCRtpReceiver.getCapabilities".
func (this RTCRtpReceiver) GetCapabilitiesFunc() (fn js.Func[func(kind js.String) RTCRtpCapabilities]) {
	return fn.FromRef(
		bindings.RTCRtpReceiverGetCapabilitiesFunc(
			this.Ref(),
		),
	)
}

// GetCapabilities calls the staticmethod "RTCRtpReceiver.getCapabilities".
func (this RTCRtpReceiver) GetCapabilities(kind js.String) (ret RTCRtpCapabilities) {
	bindings.CallRTCRtpReceiverGetCapabilities(
		this.Ref(), js.Pointer(&ret),
		kind.Ref(),
	)

	return
}

// TryGetCapabilities calls the staticmethod "RTCRtpReceiver.getCapabilities"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpReceiver) TryGetCapabilities(kind js.String) (ret RTCRtpCapabilities, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpReceiverGetCapabilities(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		kind.Ref(),
	)

	return
}

// HasGetParameters returns true if the method "RTCRtpReceiver.getParameters" exists.
func (this RTCRtpReceiver) HasGetParameters() bool {
	return js.True == bindings.HasRTCRtpReceiverGetParameters(
		this.Ref(),
	)
}

// GetParametersFunc returns the method "RTCRtpReceiver.getParameters".
func (this RTCRtpReceiver) GetParametersFunc() (fn js.Func[func() RTCRtpReceiveParameters]) {
	return fn.FromRef(
		bindings.RTCRtpReceiverGetParametersFunc(
			this.Ref(),
		),
	)
}

// GetParameters calls the method "RTCRtpReceiver.getParameters".
func (this RTCRtpReceiver) GetParameters() (ret RTCRtpReceiveParameters) {
	bindings.CallRTCRtpReceiverGetParameters(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetParameters calls the method "RTCRtpReceiver.getParameters"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpReceiver) TryGetParameters() (ret RTCRtpReceiveParameters, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpReceiverGetParameters(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetContributingSources returns true if the method "RTCRtpReceiver.getContributingSources" exists.
func (this RTCRtpReceiver) HasGetContributingSources() bool {
	return js.True == bindings.HasRTCRtpReceiverGetContributingSources(
		this.Ref(),
	)
}

// GetContributingSourcesFunc returns the method "RTCRtpReceiver.getContributingSources".
func (this RTCRtpReceiver) GetContributingSourcesFunc() (fn js.Func[func() js.Array[RTCRtpContributingSource]]) {
	return fn.FromRef(
		bindings.RTCRtpReceiverGetContributingSourcesFunc(
			this.Ref(),
		),
	)
}

// GetContributingSources calls the method "RTCRtpReceiver.getContributingSources".
func (this RTCRtpReceiver) GetContributingSources() (ret js.Array[RTCRtpContributingSource]) {
	bindings.CallRTCRtpReceiverGetContributingSources(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetContributingSources calls the method "RTCRtpReceiver.getContributingSources"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpReceiver) TryGetContributingSources() (ret js.Array[RTCRtpContributingSource], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpReceiverGetContributingSources(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetSynchronizationSources returns true if the method "RTCRtpReceiver.getSynchronizationSources" exists.
func (this RTCRtpReceiver) HasGetSynchronizationSources() bool {
	return js.True == bindings.HasRTCRtpReceiverGetSynchronizationSources(
		this.Ref(),
	)
}

// GetSynchronizationSourcesFunc returns the method "RTCRtpReceiver.getSynchronizationSources".
func (this RTCRtpReceiver) GetSynchronizationSourcesFunc() (fn js.Func[func() js.Array[RTCRtpSynchronizationSource]]) {
	return fn.FromRef(
		bindings.RTCRtpReceiverGetSynchronizationSourcesFunc(
			this.Ref(),
		),
	)
}

// GetSynchronizationSources calls the method "RTCRtpReceiver.getSynchronizationSources".
func (this RTCRtpReceiver) GetSynchronizationSources() (ret js.Array[RTCRtpSynchronizationSource]) {
	bindings.CallRTCRtpReceiverGetSynchronizationSources(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSynchronizationSources calls the method "RTCRtpReceiver.getSynchronizationSources"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpReceiver) TryGetSynchronizationSources() (ret js.Array[RTCRtpSynchronizationSource], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpReceiverGetSynchronizationSources(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetStats returns true if the method "RTCRtpReceiver.getStats" exists.
func (this RTCRtpReceiver) HasGetStats() bool {
	return js.True == bindings.HasRTCRtpReceiverGetStats(
		this.Ref(),
	)
}

// GetStatsFunc returns the method "RTCRtpReceiver.getStats".
func (this RTCRtpReceiver) GetStatsFunc() (fn js.Func[func() js.Promise[RTCStatsReport]]) {
	return fn.FromRef(
		bindings.RTCRtpReceiverGetStatsFunc(
			this.Ref(),
		),
	)
}

// GetStats calls the method "RTCRtpReceiver.getStats".
func (this RTCRtpReceiver) GetStats() (ret js.Promise[RTCStatsReport]) {
	bindings.CallRTCRtpReceiverGetStats(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetStats calls the method "RTCRtpReceiver.getStats"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpReceiver) TryGetStats() (ret js.Promise[RTCStatsReport], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpReceiverGetStats(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this RTCRtpTransceiver) Mid() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCRtpTransceiverMid(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Sender returns the value of property "RTCRtpTransceiver.sender".
//
// It returns ok=false if there is no such property.
func (this RTCRtpTransceiver) Sender() (ret RTCRtpSender, ok bool) {
	ok = js.True == bindings.GetRTCRtpTransceiverSender(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Receiver returns the value of property "RTCRtpTransceiver.receiver".
//
// It returns ok=false if there is no such property.
func (this RTCRtpTransceiver) Receiver() (ret RTCRtpReceiver, ok bool) {
	ok = js.True == bindings.GetRTCRtpTransceiverReceiver(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Direction returns the value of property "RTCRtpTransceiver.direction".
//
// It returns ok=false if there is no such property.
func (this RTCRtpTransceiver) Direction() (ret RTCRtpTransceiverDirection, ok bool) {
	ok = js.True == bindings.GetRTCRtpTransceiverDirection(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDirection sets the value of property "RTCRtpTransceiver.direction" to val.
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
// It returns ok=false if there is no such property.
func (this RTCRtpTransceiver) CurrentDirection() (ret RTCRtpTransceiverDirection, ok bool) {
	ok = js.True == bindings.GetRTCRtpTransceiverCurrentDirection(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasStop returns true if the method "RTCRtpTransceiver.stop" exists.
func (this RTCRtpTransceiver) HasStop() bool {
	return js.True == bindings.HasRTCRtpTransceiverStop(
		this.Ref(),
	)
}

// StopFunc returns the method "RTCRtpTransceiver.stop".
func (this RTCRtpTransceiver) StopFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RTCRtpTransceiverStopFunc(
			this.Ref(),
		),
	)
}

// Stop calls the method "RTCRtpTransceiver.stop".
func (this RTCRtpTransceiver) Stop() (ret js.Void) {
	bindings.CallRTCRtpTransceiverStop(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStop calls the method "RTCRtpTransceiver.stop"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpTransceiver) TryStop() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpTransceiverStop(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetCodecPreferences returns true if the method "RTCRtpTransceiver.setCodecPreferences" exists.
func (this RTCRtpTransceiver) HasSetCodecPreferences() bool {
	return js.True == bindings.HasRTCRtpTransceiverSetCodecPreferences(
		this.Ref(),
	)
}

// SetCodecPreferencesFunc returns the method "RTCRtpTransceiver.setCodecPreferences".
func (this RTCRtpTransceiver) SetCodecPreferencesFunc() (fn js.Func[func(codecs js.Array[RTCRtpCodecCapability])]) {
	return fn.FromRef(
		bindings.RTCRtpTransceiverSetCodecPreferencesFunc(
			this.Ref(),
		),
	)
}

// SetCodecPreferences calls the method "RTCRtpTransceiver.setCodecPreferences".
func (this RTCRtpTransceiver) SetCodecPreferences(codecs js.Array[RTCRtpCodecCapability]) (ret js.Void) {
	bindings.CallRTCRtpTransceiverSetCodecPreferences(
		this.Ref(), js.Pointer(&ret),
		codecs.Ref(),
	)

	return
}

// TrySetCodecPreferences calls the method "RTCRtpTransceiver.setCodecPreferences"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpTransceiver) TrySetCodecPreferences(codecs js.Array[RTCRtpCodecCapability]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpTransceiverSetCodecPreferences(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		codecs.Ref(),
	)

	return
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

// New creates a new RTCRtpTransceiverInit in the application heap.
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

func NewRTCSessionDescription(descriptionInitDict RTCSessionDescriptionInit) (ret RTCSessionDescription) {
	ret.ref = bindings.NewRTCSessionDescriptionByRTCSessionDescription(
		js.Pointer(&descriptionInitDict))
	return
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
// It returns ok=false if there is no such property.
func (this RTCSessionDescription) Type() (ret RTCSdpType, ok bool) {
	ok = js.True == bindings.GetRTCSessionDescriptionType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Sdp returns the value of property "RTCSessionDescription.sdp".
//
// It returns ok=false if there is no such property.
func (this RTCSessionDescription) Sdp() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCSessionDescriptionSdp(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "RTCSessionDescription.toJSON" exists.
func (this RTCSessionDescription) HasToJSON() bool {
	return js.True == bindings.HasRTCSessionDescriptionToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "RTCSessionDescription.toJSON".
func (this RTCSessionDescription) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.RTCSessionDescriptionToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "RTCSessionDescription.toJSON".
func (this RTCSessionDescription) ToJSON() (ret js.Object) {
	bindings.CallRTCSessionDescriptionToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "RTCSessionDescription.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCSessionDescription) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCSessionDescriptionToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this RTCSctpTransport) Transport() (ret RTCDtlsTransport, ok bool) {
	ok = js.True == bindings.GetRTCSctpTransportTransport(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// State returns the value of property "RTCSctpTransport.state".
//
// It returns ok=false if there is no such property.
func (this RTCSctpTransport) State() (ret RTCSctpTransportState, ok bool) {
	ok = js.True == bindings.GetRTCSctpTransportState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxMessageSize returns the value of property "RTCSctpTransport.maxMessageSize".
//
// It returns ok=false if there is no such property.
func (this RTCSctpTransport) MaxMessageSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetRTCSctpTransportMaxMessageSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxChannels returns the value of property "RTCSctpTransport.maxChannels".
//
// It returns ok=false if there is no such property.
func (this RTCSctpTransport) MaxChannels() (ret uint16, ok bool) {
	ok = js.True == bindings.GetRTCSctpTransportMaxChannels(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewRTCPeerConnection(configuration RTCConfiguration) (ret RTCPeerConnection) {
	ret.ref = bindings.NewRTCPeerConnectionByRTCPeerConnection(
		js.Pointer(&configuration))
	return
}

func NewRTCPeerConnectionByRTCPeerConnection1() (ret RTCPeerConnection) {
	ret.ref = bindings.NewRTCPeerConnectionByRTCPeerConnection1()
	return
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
// It returns ok=false if there is no such property.
func (this RTCPeerConnection) LocalDescription() (ret RTCSessionDescription, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionLocalDescription(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CurrentLocalDescription returns the value of property "RTCPeerConnection.currentLocalDescription".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnection) CurrentLocalDescription() (ret RTCSessionDescription, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionCurrentLocalDescription(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PendingLocalDescription returns the value of property "RTCPeerConnection.pendingLocalDescription".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnection) PendingLocalDescription() (ret RTCSessionDescription, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionPendingLocalDescription(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RemoteDescription returns the value of property "RTCPeerConnection.remoteDescription".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnection) RemoteDescription() (ret RTCSessionDescription, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionRemoteDescription(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CurrentRemoteDescription returns the value of property "RTCPeerConnection.currentRemoteDescription".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnection) CurrentRemoteDescription() (ret RTCSessionDescription, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionCurrentRemoteDescription(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PendingRemoteDescription returns the value of property "RTCPeerConnection.pendingRemoteDescription".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnection) PendingRemoteDescription() (ret RTCSessionDescription, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionPendingRemoteDescription(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SignalingState returns the value of property "RTCPeerConnection.signalingState".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnection) SignalingState() (ret RTCSignalingState, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionSignalingState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IceGatheringState returns the value of property "RTCPeerConnection.iceGatheringState".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnection) IceGatheringState() (ret RTCIceGatheringState, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionIceGatheringState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IceConnectionState returns the value of property "RTCPeerConnection.iceConnectionState".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnection) IceConnectionState() (ret RTCIceConnectionState, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionIceConnectionState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ConnectionState returns the value of property "RTCPeerConnection.connectionState".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnection) ConnectionState() (ret RTCPeerConnectionState, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionConnectionState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CanTrickleIceCandidates returns the value of property "RTCPeerConnection.canTrickleIceCandidates".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnection) CanTrickleIceCandidates() (ret bool, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionCanTrickleIceCandidates(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Sctp returns the value of property "RTCPeerConnection.sctp".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnection) Sctp() (ret RTCSctpTransport, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionSctp(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PeerIdentity returns the value of property "RTCPeerConnection.peerIdentity".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnection) PeerIdentity() (ret js.Promise[RTCIdentityAssertion], ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionPeerIdentity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IdpLoginUrl returns the value of property "RTCPeerConnection.idpLoginUrl".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnection) IdpLoginUrl() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionIdpLoginUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IdpErrorInfo returns the value of property "RTCPeerConnection.idpErrorInfo".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnection) IdpErrorInfo() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionIdpErrorInfo(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasCreateOffer returns true if the method "RTCPeerConnection.createOffer" exists.
func (this RTCPeerConnection) HasCreateOffer() bool {
	return js.True == bindings.HasRTCPeerConnectionCreateOffer(
		this.Ref(),
	)
}

// CreateOfferFunc returns the method "RTCPeerConnection.createOffer".
func (this RTCPeerConnection) CreateOfferFunc() (fn js.Func[func(options RTCOfferOptions) js.Promise[RTCSessionDescriptionInit]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateOfferFunc(
			this.Ref(),
		),
	)
}

// CreateOffer calls the method "RTCPeerConnection.createOffer".
func (this RTCPeerConnection) CreateOffer(options RTCOfferOptions) (ret js.Promise[RTCSessionDescriptionInit]) {
	bindings.CallRTCPeerConnectionCreateOffer(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryCreateOffer calls the method "RTCPeerConnection.createOffer"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryCreateOffer(options RTCOfferOptions) (ret js.Promise[RTCSessionDescriptionInit], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionCreateOffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasCreateOffer1 returns true if the method "RTCPeerConnection.createOffer" exists.
func (this RTCPeerConnection) HasCreateOffer1() bool {
	return js.True == bindings.HasRTCPeerConnectionCreateOffer1(
		this.Ref(),
	)
}

// CreateOffer1Func returns the method "RTCPeerConnection.createOffer".
func (this RTCPeerConnection) CreateOffer1Func() (fn js.Func[func() js.Promise[RTCSessionDescriptionInit]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateOffer1Func(
			this.Ref(),
		),
	)
}

// CreateOffer1 calls the method "RTCPeerConnection.createOffer".
func (this RTCPeerConnection) CreateOffer1() (ret js.Promise[RTCSessionDescriptionInit]) {
	bindings.CallRTCPeerConnectionCreateOffer1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateOffer1 calls the method "RTCPeerConnection.createOffer"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryCreateOffer1() (ret js.Promise[RTCSessionDescriptionInit], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionCreateOffer1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateAnswer returns true if the method "RTCPeerConnection.createAnswer" exists.
func (this RTCPeerConnection) HasCreateAnswer() bool {
	return js.True == bindings.HasRTCPeerConnectionCreateAnswer(
		this.Ref(),
	)
}

// CreateAnswerFunc returns the method "RTCPeerConnection.createAnswer".
func (this RTCPeerConnection) CreateAnswerFunc() (fn js.Func[func(options RTCAnswerOptions) js.Promise[RTCSessionDescriptionInit]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateAnswerFunc(
			this.Ref(),
		),
	)
}

// CreateAnswer calls the method "RTCPeerConnection.createAnswer".
func (this RTCPeerConnection) CreateAnswer(options RTCAnswerOptions) (ret js.Promise[RTCSessionDescriptionInit]) {
	bindings.CallRTCPeerConnectionCreateAnswer(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryCreateAnswer calls the method "RTCPeerConnection.createAnswer"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryCreateAnswer(options RTCAnswerOptions) (ret js.Promise[RTCSessionDescriptionInit], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionCreateAnswer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasCreateAnswer1 returns true if the method "RTCPeerConnection.createAnswer" exists.
func (this RTCPeerConnection) HasCreateAnswer1() bool {
	return js.True == bindings.HasRTCPeerConnectionCreateAnswer1(
		this.Ref(),
	)
}

// CreateAnswer1Func returns the method "RTCPeerConnection.createAnswer".
func (this RTCPeerConnection) CreateAnswer1Func() (fn js.Func[func() js.Promise[RTCSessionDescriptionInit]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateAnswer1Func(
			this.Ref(),
		),
	)
}

// CreateAnswer1 calls the method "RTCPeerConnection.createAnswer".
func (this RTCPeerConnection) CreateAnswer1() (ret js.Promise[RTCSessionDescriptionInit]) {
	bindings.CallRTCPeerConnectionCreateAnswer1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateAnswer1 calls the method "RTCPeerConnection.createAnswer"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryCreateAnswer1() (ret js.Promise[RTCSessionDescriptionInit], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionCreateAnswer1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetLocalDescription returns true if the method "RTCPeerConnection.setLocalDescription" exists.
func (this RTCPeerConnection) HasSetLocalDescription() bool {
	return js.True == bindings.HasRTCPeerConnectionSetLocalDescription(
		this.Ref(),
	)
}

// SetLocalDescriptionFunc returns the method "RTCPeerConnection.setLocalDescription".
func (this RTCPeerConnection) SetLocalDescriptionFunc() (fn js.Func[func(description RTCLocalSessionDescriptionInit) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetLocalDescriptionFunc(
			this.Ref(),
		),
	)
}

// SetLocalDescription calls the method "RTCPeerConnection.setLocalDescription".
func (this RTCPeerConnection) SetLocalDescription(description RTCLocalSessionDescriptionInit) (ret js.Promise[js.Void]) {
	bindings.CallRTCPeerConnectionSetLocalDescription(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&description),
	)

	return
}

// TrySetLocalDescription calls the method "RTCPeerConnection.setLocalDescription"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TrySetLocalDescription(description RTCLocalSessionDescriptionInit) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionSetLocalDescription(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&description),
	)

	return
}

// HasSetLocalDescription1 returns true if the method "RTCPeerConnection.setLocalDescription" exists.
func (this RTCPeerConnection) HasSetLocalDescription1() bool {
	return js.True == bindings.HasRTCPeerConnectionSetLocalDescription1(
		this.Ref(),
	)
}

// SetLocalDescription1Func returns the method "RTCPeerConnection.setLocalDescription".
func (this RTCPeerConnection) SetLocalDescription1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetLocalDescription1Func(
			this.Ref(),
		),
	)
}

// SetLocalDescription1 calls the method "RTCPeerConnection.setLocalDescription".
func (this RTCPeerConnection) SetLocalDescription1() (ret js.Promise[js.Void]) {
	bindings.CallRTCPeerConnectionSetLocalDescription1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySetLocalDescription1 calls the method "RTCPeerConnection.setLocalDescription"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TrySetLocalDescription1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionSetLocalDescription1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetRemoteDescription returns true if the method "RTCPeerConnection.setRemoteDescription" exists.
func (this RTCPeerConnection) HasSetRemoteDescription() bool {
	return js.True == bindings.HasRTCPeerConnectionSetRemoteDescription(
		this.Ref(),
	)
}

// SetRemoteDescriptionFunc returns the method "RTCPeerConnection.setRemoteDescription".
func (this RTCPeerConnection) SetRemoteDescriptionFunc() (fn js.Func[func(description RTCSessionDescriptionInit) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetRemoteDescriptionFunc(
			this.Ref(),
		),
	)
}

// SetRemoteDescription calls the method "RTCPeerConnection.setRemoteDescription".
func (this RTCPeerConnection) SetRemoteDescription(description RTCSessionDescriptionInit) (ret js.Promise[js.Void]) {
	bindings.CallRTCPeerConnectionSetRemoteDescription(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&description),
	)

	return
}

// TrySetRemoteDescription calls the method "RTCPeerConnection.setRemoteDescription"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TrySetRemoteDescription(description RTCSessionDescriptionInit) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionSetRemoteDescription(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&description),
	)

	return
}

// HasAddIceCandidate returns true if the method "RTCPeerConnection.addIceCandidate" exists.
func (this RTCPeerConnection) HasAddIceCandidate() bool {
	return js.True == bindings.HasRTCPeerConnectionAddIceCandidate(
		this.Ref(),
	)
}

// AddIceCandidateFunc returns the method "RTCPeerConnection.addIceCandidate".
func (this RTCPeerConnection) AddIceCandidateFunc() (fn js.Func[func(candidate RTCIceCandidateInit) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionAddIceCandidateFunc(
			this.Ref(),
		),
	)
}

// AddIceCandidate calls the method "RTCPeerConnection.addIceCandidate".
func (this RTCPeerConnection) AddIceCandidate(candidate RTCIceCandidateInit) (ret js.Promise[js.Void]) {
	bindings.CallRTCPeerConnectionAddIceCandidate(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&candidate),
	)

	return
}

// TryAddIceCandidate calls the method "RTCPeerConnection.addIceCandidate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryAddIceCandidate(candidate RTCIceCandidateInit) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionAddIceCandidate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&candidate),
	)

	return
}

// HasAddIceCandidate1 returns true if the method "RTCPeerConnection.addIceCandidate" exists.
func (this RTCPeerConnection) HasAddIceCandidate1() bool {
	return js.True == bindings.HasRTCPeerConnectionAddIceCandidate1(
		this.Ref(),
	)
}

// AddIceCandidate1Func returns the method "RTCPeerConnection.addIceCandidate".
func (this RTCPeerConnection) AddIceCandidate1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionAddIceCandidate1Func(
			this.Ref(),
		),
	)
}

// AddIceCandidate1 calls the method "RTCPeerConnection.addIceCandidate".
func (this RTCPeerConnection) AddIceCandidate1() (ret js.Promise[js.Void]) {
	bindings.CallRTCPeerConnectionAddIceCandidate1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAddIceCandidate1 calls the method "RTCPeerConnection.addIceCandidate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryAddIceCandidate1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionAddIceCandidate1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRestartIce returns true if the method "RTCPeerConnection.restartIce" exists.
func (this RTCPeerConnection) HasRestartIce() bool {
	return js.True == bindings.HasRTCPeerConnectionRestartIce(
		this.Ref(),
	)
}

// RestartIceFunc returns the method "RTCPeerConnection.restartIce".
func (this RTCPeerConnection) RestartIceFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionRestartIceFunc(
			this.Ref(),
		),
	)
}

// RestartIce calls the method "RTCPeerConnection.restartIce".
func (this RTCPeerConnection) RestartIce() (ret js.Void) {
	bindings.CallRTCPeerConnectionRestartIce(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRestartIce calls the method "RTCPeerConnection.restartIce"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryRestartIce() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionRestartIce(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetConfiguration returns true if the method "RTCPeerConnection.getConfiguration" exists.
func (this RTCPeerConnection) HasGetConfiguration() bool {
	return js.True == bindings.HasRTCPeerConnectionGetConfiguration(
		this.Ref(),
	)
}

// GetConfigurationFunc returns the method "RTCPeerConnection.getConfiguration".
func (this RTCPeerConnection) GetConfigurationFunc() (fn js.Func[func() RTCConfiguration]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGetConfigurationFunc(
			this.Ref(),
		),
	)
}

// GetConfiguration calls the method "RTCPeerConnection.getConfiguration".
func (this RTCPeerConnection) GetConfiguration() (ret RTCConfiguration) {
	bindings.CallRTCPeerConnectionGetConfiguration(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetConfiguration calls the method "RTCPeerConnection.getConfiguration"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryGetConfiguration() (ret RTCConfiguration, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionGetConfiguration(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetConfiguration returns true if the method "RTCPeerConnection.setConfiguration" exists.
func (this RTCPeerConnection) HasSetConfiguration() bool {
	return js.True == bindings.HasRTCPeerConnectionSetConfiguration(
		this.Ref(),
	)
}

// SetConfigurationFunc returns the method "RTCPeerConnection.setConfiguration".
func (this RTCPeerConnection) SetConfigurationFunc() (fn js.Func[func(configuration RTCConfiguration)]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetConfigurationFunc(
			this.Ref(),
		),
	)
}

// SetConfiguration calls the method "RTCPeerConnection.setConfiguration".
func (this RTCPeerConnection) SetConfiguration(configuration RTCConfiguration) (ret js.Void) {
	bindings.CallRTCPeerConnectionSetConfiguration(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&configuration),
	)

	return
}

// TrySetConfiguration calls the method "RTCPeerConnection.setConfiguration"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TrySetConfiguration(configuration RTCConfiguration) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionSetConfiguration(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&configuration),
	)

	return
}

// HasSetConfiguration1 returns true if the method "RTCPeerConnection.setConfiguration" exists.
func (this RTCPeerConnection) HasSetConfiguration1() bool {
	return js.True == bindings.HasRTCPeerConnectionSetConfiguration1(
		this.Ref(),
	)
}

// SetConfiguration1Func returns the method "RTCPeerConnection.setConfiguration".
func (this RTCPeerConnection) SetConfiguration1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetConfiguration1Func(
			this.Ref(),
		),
	)
}

// SetConfiguration1 calls the method "RTCPeerConnection.setConfiguration".
func (this RTCPeerConnection) SetConfiguration1() (ret js.Void) {
	bindings.CallRTCPeerConnectionSetConfiguration1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySetConfiguration1 calls the method "RTCPeerConnection.setConfiguration"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TrySetConfiguration1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionSetConfiguration1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "RTCPeerConnection.close" exists.
func (this RTCPeerConnection) HasClose() bool {
	return js.True == bindings.HasRTCPeerConnectionClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "RTCPeerConnection.close".
func (this RTCPeerConnection) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "RTCPeerConnection.close".
func (this RTCPeerConnection) Close() (ret js.Void) {
	bindings.CallRTCPeerConnectionClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "RTCPeerConnection.close"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateOffer2 returns true if the method "RTCPeerConnection.createOffer" exists.
func (this RTCPeerConnection) HasCreateOffer2() bool {
	return js.True == bindings.HasRTCPeerConnectionCreateOffer2(
		this.Ref(),
	)
}

// CreateOffer2Func returns the method "RTCPeerConnection.createOffer".
func (this RTCPeerConnection) CreateOffer2Func() (fn js.Func[func(successCallback js.Func[func(description RTCSessionDescriptionInit)], failureCallback js.Func[func(err DOMException)], options RTCOfferOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateOffer2Func(
			this.Ref(),
		),
	)
}

// CreateOffer2 calls the method "RTCPeerConnection.createOffer".
func (this RTCPeerConnection) CreateOffer2(successCallback js.Func[func(description RTCSessionDescriptionInit)], failureCallback js.Func[func(err DOMException)], options RTCOfferOptions) (ret js.Promise[js.Void]) {
	bindings.CallRTCPeerConnectionCreateOffer2(
		this.Ref(), js.Pointer(&ret),
		successCallback.Ref(),
		failureCallback.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryCreateOffer2 calls the method "RTCPeerConnection.createOffer"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryCreateOffer2(successCallback js.Func[func(description RTCSessionDescriptionInit)], failureCallback js.Func[func(err DOMException)], options RTCOfferOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionCreateOffer2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
		failureCallback.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasCreateOffer3 returns true if the method "RTCPeerConnection.createOffer" exists.
func (this RTCPeerConnection) HasCreateOffer3() bool {
	return js.True == bindings.HasRTCPeerConnectionCreateOffer3(
		this.Ref(),
	)
}

// CreateOffer3Func returns the method "RTCPeerConnection.createOffer".
func (this RTCPeerConnection) CreateOffer3Func() (fn js.Func[func(successCallback js.Func[func(description RTCSessionDescriptionInit)], failureCallback js.Func[func(err DOMException)]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateOffer3Func(
			this.Ref(),
		),
	)
}

// CreateOffer3 calls the method "RTCPeerConnection.createOffer".
func (this RTCPeerConnection) CreateOffer3(successCallback js.Func[func(description RTCSessionDescriptionInit)], failureCallback js.Func[func(err DOMException)]) (ret js.Promise[js.Void]) {
	bindings.CallRTCPeerConnectionCreateOffer3(
		this.Ref(), js.Pointer(&ret),
		successCallback.Ref(),
		failureCallback.Ref(),
	)

	return
}

// TryCreateOffer3 calls the method "RTCPeerConnection.createOffer"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryCreateOffer3(successCallback js.Func[func(description RTCSessionDescriptionInit)], failureCallback js.Func[func(err DOMException)]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionCreateOffer3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
		failureCallback.Ref(),
	)

	return
}

// HasSetLocalDescription2 returns true if the method "RTCPeerConnection.setLocalDescription" exists.
func (this RTCPeerConnection) HasSetLocalDescription2() bool {
	return js.True == bindings.HasRTCPeerConnectionSetLocalDescription2(
		this.Ref(),
	)
}

// SetLocalDescription2Func returns the method "RTCPeerConnection.setLocalDescription".
func (this RTCPeerConnection) SetLocalDescription2Func() (fn js.Func[func(description RTCLocalSessionDescriptionInit, successCallback js.Func[func()], failureCallback js.Func[func(err DOMException)]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetLocalDescription2Func(
			this.Ref(),
		),
	)
}

// SetLocalDescription2 calls the method "RTCPeerConnection.setLocalDescription".
func (this RTCPeerConnection) SetLocalDescription2(description RTCLocalSessionDescriptionInit, successCallback js.Func[func()], failureCallback js.Func[func(err DOMException)]) (ret js.Promise[js.Void]) {
	bindings.CallRTCPeerConnectionSetLocalDescription2(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&description),
		successCallback.Ref(),
		failureCallback.Ref(),
	)

	return
}

// TrySetLocalDescription2 calls the method "RTCPeerConnection.setLocalDescription"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TrySetLocalDescription2(description RTCLocalSessionDescriptionInit, successCallback js.Func[func()], failureCallback js.Func[func(err DOMException)]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionSetLocalDescription2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&description),
		successCallback.Ref(),
		failureCallback.Ref(),
	)

	return
}

// HasCreateAnswer2 returns true if the method "RTCPeerConnection.createAnswer" exists.
func (this RTCPeerConnection) HasCreateAnswer2() bool {
	return js.True == bindings.HasRTCPeerConnectionCreateAnswer2(
		this.Ref(),
	)
}

// CreateAnswer2Func returns the method "RTCPeerConnection.createAnswer".
func (this RTCPeerConnection) CreateAnswer2Func() (fn js.Func[func(successCallback js.Func[func(description RTCSessionDescriptionInit)], failureCallback js.Func[func(err DOMException)]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateAnswer2Func(
			this.Ref(),
		),
	)
}

// CreateAnswer2 calls the method "RTCPeerConnection.createAnswer".
func (this RTCPeerConnection) CreateAnswer2(successCallback js.Func[func(description RTCSessionDescriptionInit)], failureCallback js.Func[func(err DOMException)]) (ret js.Promise[js.Void]) {
	bindings.CallRTCPeerConnectionCreateAnswer2(
		this.Ref(), js.Pointer(&ret),
		successCallback.Ref(),
		failureCallback.Ref(),
	)

	return
}

// TryCreateAnswer2 calls the method "RTCPeerConnection.createAnswer"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryCreateAnswer2(successCallback js.Func[func(description RTCSessionDescriptionInit)], failureCallback js.Func[func(err DOMException)]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionCreateAnswer2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
		failureCallback.Ref(),
	)

	return
}

// HasSetRemoteDescription1 returns true if the method "RTCPeerConnection.setRemoteDescription" exists.
func (this RTCPeerConnection) HasSetRemoteDescription1() bool {
	return js.True == bindings.HasRTCPeerConnectionSetRemoteDescription1(
		this.Ref(),
	)
}

// SetRemoteDescription1Func returns the method "RTCPeerConnection.setRemoteDescription".
func (this RTCPeerConnection) SetRemoteDescription1Func() (fn js.Func[func(description RTCSessionDescriptionInit, successCallback js.Func[func()], failureCallback js.Func[func(err DOMException)]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetRemoteDescription1Func(
			this.Ref(),
		),
	)
}

// SetRemoteDescription1 calls the method "RTCPeerConnection.setRemoteDescription".
func (this RTCPeerConnection) SetRemoteDescription1(description RTCSessionDescriptionInit, successCallback js.Func[func()], failureCallback js.Func[func(err DOMException)]) (ret js.Promise[js.Void]) {
	bindings.CallRTCPeerConnectionSetRemoteDescription1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&description),
		successCallback.Ref(),
		failureCallback.Ref(),
	)

	return
}

// TrySetRemoteDescription1 calls the method "RTCPeerConnection.setRemoteDescription"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TrySetRemoteDescription1(description RTCSessionDescriptionInit, successCallback js.Func[func()], failureCallback js.Func[func(err DOMException)]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionSetRemoteDescription1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&description),
		successCallback.Ref(),
		failureCallback.Ref(),
	)

	return
}

// HasAddIceCandidate2 returns true if the method "RTCPeerConnection.addIceCandidate" exists.
func (this RTCPeerConnection) HasAddIceCandidate2() bool {
	return js.True == bindings.HasRTCPeerConnectionAddIceCandidate2(
		this.Ref(),
	)
}

// AddIceCandidate2Func returns the method "RTCPeerConnection.addIceCandidate".
func (this RTCPeerConnection) AddIceCandidate2Func() (fn js.Func[func(candidate RTCIceCandidateInit, successCallback js.Func[func()], failureCallback js.Func[func(err DOMException)]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionAddIceCandidate2Func(
			this.Ref(),
		),
	)
}

// AddIceCandidate2 calls the method "RTCPeerConnection.addIceCandidate".
func (this RTCPeerConnection) AddIceCandidate2(candidate RTCIceCandidateInit, successCallback js.Func[func()], failureCallback js.Func[func(err DOMException)]) (ret js.Promise[js.Void]) {
	bindings.CallRTCPeerConnectionAddIceCandidate2(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&candidate),
		successCallback.Ref(),
		failureCallback.Ref(),
	)

	return
}

// TryAddIceCandidate2 calls the method "RTCPeerConnection.addIceCandidate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryAddIceCandidate2(candidate RTCIceCandidateInit, successCallback js.Func[func()], failureCallback js.Func[func(err DOMException)]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionAddIceCandidate2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&candidate),
		successCallback.Ref(),
		failureCallback.Ref(),
	)

	return
}

// HasCreateDataChannel returns true if the method "RTCPeerConnection.createDataChannel" exists.
func (this RTCPeerConnection) HasCreateDataChannel() bool {
	return js.True == bindings.HasRTCPeerConnectionCreateDataChannel(
		this.Ref(),
	)
}

// CreateDataChannelFunc returns the method "RTCPeerConnection.createDataChannel".
func (this RTCPeerConnection) CreateDataChannelFunc() (fn js.Func[func(label js.String, dataChannelDict RTCDataChannelInit) RTCDataChannel]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateDataChannelFunc(
			this.Ref(),
		),
	)
}

// CreateDataChannel calls the method "RTCPeerConnection.createDataChannel".
func (this RTCPeerConnection) CreateDataChannel(label js.String, dataChannelDict RTCDataChannelInit) (ret RTCDataChannel) {
	bindings.CallRTCPeerConnectionCreateDataChannel(
		this.Ref(), js.Pointer(&ret),
		label.Ref(),
		js.Pointer(&dataChannelDict),
	)

	return
}

// TryCreateDataChannel calls the method "RTCPeerConnection.createDataChannel"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryCreateDataChannel(label js.String, dataChannelDict RTCDataChannelInit) (ret RTCDataChannel, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionCreateDataChannel(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		label.Ref(),
		js.Pointer(&dataChannelDict),
	)

	return
}

// HasCreateDataChannel1 returns true if the method "RTCPeerConnection.createDataChannel" exists.
func (this RTCPeerConnection) HasCreateDataChannel1() bool {
	return js.True == bindings.HasRTCPeerConnectionCreateDataChannel1(
		this.Ref(),
	)
}

// CreateDataChannel1Func returns the method "RTCPeerConnection.createDataChannel".
func (this RTCPeerConnection) CreateDataChannel1Func() (fn js.Func[func(label js.String) RTCDataChannel]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionCreateDataChannel1Func(
			this.Ref(),
		),
	)
}

// CreateDataChannel1 calls the method "RTCPeerConnection.createDataChannel".
func (this RTCPeerConnection) CreateDataChannel1(label js.String) (ret RTCDataChannel) {
	bindings.CallRTCPeerConnectionCreateDataChannel1(
		this.Ref(), js.Pointer(&ret),
		label.Ref(),
	)

	return
}

// TryCreateDataChannel1 calls the method "RTCPeerConnection.createDataChannel"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryCreateDataChannel1(label js.String) (ret RTCDataChannel, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionCreateDataChannel1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		label.Ref(),
	)

	return
}

// HasGetSenders returns true if the method "RTCPeerConnection.getSenders" exists.
func (this RTCPeerConnection) HasGetSenders() bool {
	return js.True == bindings.HasRTCPeerConnectionGetSenders(
		this.Ref(),
	)
}

// GetSendersFunc returns the method "RTCPeerConnection.getSenders".
func (this RTCPeerConnection) GetSendersFunc() (fn js.Func[func() js.Array[RTCRtpSender]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGetSendersFunc(
			this.Ref(),
		),
	)
}

// GetSenders calls the method "RTCPeerConnection.getSenders".
func (this RTCPeerConnection) GetSenders() (ret js.Array[RTCRtpSender]) {
	bindings.CallRTCPeerConnectionGetSenders(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSenders calls the method "RTCPeerConnection.getSenders"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryGetSenders() (ret js.Array[RTCRtpSender], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionGetSenders(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetReceivers returns true if the method "RTCPeerConnection.getReceivers" exists.
func (this RTCPeerConnection) HasGetReceivers() bool {
	return js.True == bindings.HasRTCPeerConnectionGetReceivers(
		this.Ref(),
	)
}

// GetReceiversFunc returns the method "RTCPeerConnection.getReceivers".
func (this RTCPeerConnection) GetReceiversFunc() (fn js.Func[func() js.Array[RTCRtpReceiver]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGetReceiversFunc(
			this.Ref(),
		),
	)
}

// GetReceivers calls the method "RTCPeerConnection.getReceivers".
func (this RTCPeerConnection) GetReceivers() (ret js.Array[RTCRtpReceiver]) {
	bindings.CallRTCPeerConnectionGetReceivers(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetReceivers calls the method "RTCPeerConnection.getReceivers"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryGetReceivers() (ret js.Array[RTCRtpReceiver], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionGetReceivers(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetTransceivers returns true if the method "RTCPeerConnection.getTransceivers" exists.
func (this RTCPeerConnection) HasGetTransceivers() bool {
	return js.True == bindings.HasRTCPeerConnectionGetTransceivers(
		this.Ref(),
	)
}

// GetTransceiversFunc returns the method "RTCPeerConnection.getTransceivers".
func (this RTCPeerConnection) GetTransceiversFunc() (fn js.Func[func() js.Array[RTCRtpTransceiver]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGetTransceiversFunc(
			this.Ref(),
		),
	)
}

// GetTransceivers calls the method "RTCPeerConnection.getTransceivers".
func (this RTCPeerConnection) GetTransceivers() (ret js.Array[RTCRtpTransceiver]) {
	bindings.CallRTCPeerConnectionGetTransceivers(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetTransceivers calls the method "RTCPeerConnection.getTransceivers"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryGetTransceivers() (ret js.Array[RTCRtpTransceiver], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionGetTransceivers(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasAddTrack returns true if the method "RTCPeerConnection.addTrack" exists.
func (this RTCPeerConnection) HasAddTrack() bool {
	return js.True == bindings.HasRTCPeerConnectionAddTrack(
		this.Ref(),
	)
}

// AddTrackFunc returns the method "RTCPeerConnection.addTrack".
func (this RTCPeerConnection) AddTrackFunc() (fn js.Func[func(track MediaStreamTrack, streams ...MediaStream) RTCRtpSender]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionAddTrackFunc(
			this.Ref(),
		),
	)
}

// AddTrack calls the method "RTCPeerConnection.addTrack".
func (this RTCPeerConnection) AddTrack(track MediaStreamTrack, streams ...MediaStream) (ret RTCRtpSender) {
	bindings.CallRTCPeerConnectionAddTrack(
		this.Ref(), js.Pointer(&ret),
		track.Ref(),
		js.SliceData(streams),
		js.SizeU(len(streams)),
	)

	return
}

// TryAddTrack calls the method "RTCPeerConnection.addTrack"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryAddTrack(track MediaStreamTrack, streams ...MediaStream) (ret RTCRtpSender, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionAddTrack(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		track.Ref(),
		js.SliceData(streams),
		js.SizeU(len(streams)),
	)

	return
}

// HasRemoveTrack returns true if the method "RTCPeerConnection.removeTrack" exists.
func (this RTCPeerConnection) HasRemoveTrack() bool {
	return js.True == bindings.HasRTCPeerConnectionRemoveTrack(
		this.Ref(),
	)
}

// RemoveTrackFunc returns the method "RTCPeerConnection.removeTrack".
func (this RTCPeerConnection) RemoveTrackFunc() (fn js.Func[func(sender RTCRtpSender)]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionRemoveTrackFunc(
			this.Ref(),
		),
	)
}

// RemoveTrack calls the method "RTCPeerConnection.removeTrack".
func (this RTCPeerConnection) RemoveTrack(sender RTCRtpSender) (ret js.Void) {
	bindings.CallRTCPeerConnectionRemoveTrack(
		this.Ref(), js.Pointer(&ret),
		sender.Ref(),
	)

	return
}

// TryRemoveTrack calls the method "RTCPeerConnection.removeTrack"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryRemoveTrack(sender RTCRtpSender) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionRemoveTrack(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sender.Ref(),
	)

	return
}

// HasAddTransceiver returns true if the method "RTCPeerConnection.addTransceiver" exists.
func (this RTCPeerConnection) HasAddTransceiver() bool {
	return js.True == bindings.HasRTCPeerConnectionAddTransceiver(
		this.Ref(),
	)
}

// AddTransceiverFunc returns the method "RTCPeerConnection.addTransceiver".
func (this RTCPeerConnection) AddTransceiverFunc() (fn js.Func[func(trackOrKind OneOf_MediaStreamTrack_String, init RTCRtpTransceiverInit) RTCRtpTransceiver]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionAddTransceiverFunc(
			this.Ref(),
		),
	)
}

// AddTransceiver calls the method "RTCPeerConnection.addTransceiver".
func (this RTCPeerConnection) AddTransceiver(trackOrKind OneOf_MediaStreamTrack_String, init RTCRtpTransceiverInit) (ret RTCRtpTransceiver) {
	bindings.CallRTCPeerConnectionAddTransceiver(
		this.Ref(), js.Pointer(&ret),
		trackOrKind.Ref(),
		js.Pointer(&init),
	)

	return
}

// TryAddTransceiver calls the method "RTCPeerConnection.addTransceiver"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryAddTransceiver(trackOrKind OneOf_MediaStreamTrack_String, init RTCRtpTransceiverInit) (ret RTCRtpTransceiver, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionAddTransceiver(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		trackOrKind.Ref(),
		js.Pointer(&init),
	)

	return
}

// HasAddTransceiver1 returns true if the method "RTCPeerConnection.addTransceiver" exists.
func (this RTCPeerConnection) HasAddTransceiver1() bool {
	return js.True == bindings.HasRTCPeerConnectionAddTransceiver1(
		this.Ref(),
	)
}

// AddTransceiver1Func returns the method "RTCPeerConnection.addTransceiver".
func (this RTCPeerConnection) AddTransceiver1Func() (fn js.Func[func(trackOrKind OneOf_MediaStreamTrack_String) RTCRtpTransceiver]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionAddTransceiver1Func(
			this.Ref(),
		),
	)
}

// AddTransceiver1 calls the method "RTCPeerConnection.addTransceiver".
func (this RTCPeerConnection) AddTransceiver1(trackOrKind OneOf_MediaStreamTrack_String) (ret RTCRtpTransceiver) {
	bindings.CallRTCPeerConnectionAddTransceiver1(
		this.Ref(), js.Pointer(&ret),
		trackOrKind.Ref(),
	)

	return
}

// TryAddTransceiver1 calls the method "RTCPeerConnection.addTransceiver"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryAddTransceiver1(trackOrKind OneOf_MediaStreamTrack_String) (ret RTCRtpTransceiver, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionAddTransceiver1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		trackOrKind.Ref(),
	)

	return
}

// HasGetStats returns true if the method "RTCPeerConnection.getStats" exists.
func (this RTCPeerConnection) HasGetStats() bool {
	return js.True == bindings.HasRTCPeerConnectionGetStats(
		this.Ref(),
	)
}

// GetStatsFunc returns the method "RTCPeerConnection.getStats".
func (this RTCPeerConnection) GetStatsFunc() (fn js.Func[func(selector MediaStreamTrack) js.Promise[RTCStatsReport]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGetStatsFunc(
			this.Ref(),
		),
	)
}

// GetStats calls the method "RTCPeerConnection.getStats".
func (this RTCPeerConnection) GetStats(selector MediaStreamTrack) (ret js.Promise[RTCStatsReport]) {
	bindings.CallRTCPeerConnectionGetStats(
		this.Ref(), js.Pointer(&ret),
		selector.Ref(),
	)

	return
}

// TryGetStats calls the method "RTCPeerConnection.getStats"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryGetStats(selector MediaStreamTrack) (ret js.Promise[RTCStatsReport], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionGetStats(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		selector.Ref(),
	)

	return
}

// HasGetStats1 returns true if the method "RTCPeerConnection.getStats" exists.
func (this RTCPeerConnection) HasGetStats1() bool {
	return js.True == bindings.HasRTCPeerConnectionGetStats1(
		this.Ref(),
	)
}

// GetStats1Func returns the method "RTCPeerConnection.getStats".
func (this RTCPeerConnection) GetStats1Func() (fn js.Func[func() js.Promise[RTCStatsReport]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGetStats1Func(
			this.Ref(),
		),
	)
}

// GetStats1 calls the method "RTCPeerConnection.getStats".
func (this RTCPeerConnection) GetStats1() (ret js.Promise[RTCStatsReport]) {
	bindings.CallRTCPeerConnectionGetStats1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetStats1 calls the method "RTCPeerConnection.getStats"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryGetStats1() (ret js.Promise[RTCStatsReport], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionGetStats1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGenerateCertificate returns true if the staticmethod "RTCPeerConnection.generateCertificate" exists.
func (this RTCPeerConnection) HasGenerateCertificate() bool {
	return js.True == bindings.HasRTCPeerConnectionGenerateCertificate(
		this.Ref(),
	)
}

// GenerateCertificateFunc returns the staticmethod "RTCPeerConnection.generateCertificate".
func (this RTCPeerConnection) GenerateCertificateFunc() (fn js.Func[func(keygenAlgorithm AlgorithmIdentifier) js.Promise[RTCCertificate]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGenerateCertificateFunc(
			this.Ref(),
		),
	)
}

// GenerateCertificate calls the staticmethod "RTCPeerConnection.generateCertificate".
func (this RTCPeerConnection) GenerateCertificate(keygenAlgorithm AlgorithmIdentifier) (ret js.Promise[RTCCertificate]) {
	bindings.CallRTCPeerConnectionGenerateCertificate(
		this.Ref(), js.Pointer(&ret),
		keygenAlgorithm.Ref(),
	)

	return
}

// TryGenerateCertificate calls the staticmethod "RTCPeerConnection.generateCertificate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryGenerateCertificate(keygenAlgorithm AlgorithmIdentifier) (ret js.Promise[RTCCertificate], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionGenerateCertificate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		keygenAlgorithm.Ref(),
	)

	return
}

// HasSetIdentityProvider returns true if the method "RTCPeerConnection.setIdentityProvider" exists.
func (this RTCPeerConnection) HasSetIdentityProvider() bool {
	return js.True == bindings.HasRTCPeerConnectionSetIdentityProvider(
		this.Ref(),
	)
}

// SetIdentityProviderFunc returns the method "RTCPeerConnection.setIdentityProvider".
func (this RTCPeerConnection) SetIdentityProviderFunc() (fn js.Func[func(provider js.String, options RTCIdentityProviderOptions)]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetIdentityProviderFunc(
			this.Ref(),
		),
	)
}

// SetIdentityProvider calls the method "RTCPeerConnection.setIdentityProvider".
func (this RTCPeerConnection) SetIdentityProvider(provider js.String, options RTCIdentityProviderOptions) (ret js.Void) {
	bindings.CallRTCPeerConnectionSetIdentityProvider(
		this.Ref(), js.Pointer(&ret),
		provider.Ref(),
		js.Pointer(&options),
	)

	return
}

// TrySetIdentityProvider calls the method "RTCPeerConnection.setIdentityProvider"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TrySetIdentityProvider(provider js.String, options RTCIdentityProviderOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionSetIdentityProvider(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		provider.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasSetIdentityProvider1 returns true if the method "RTCPeerConnection.setIdentityProvider" exists.
func (this RTCPeerConnection) HasSetIdentityProvider1() bool {
	return js.True == bindings.HasRTCPeerConnectionSetIdentityProvider1(
		this.Ref(),
	)
}

// SetIdentityProvider1Func returns the method "RTCPeerConnection.setIdentityProvider".
func (this RTCPeerConnection) SetIdentityProvider1Func() (fn js.Func[func(provider js.String)]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionSetIdentityProvider1Func(
			this.Ref(),
		),
	)
}

// SetIdentityProvider1 calls the method "RTCPeerConnection.setIdentityProvider".
func (this RTCPeerConnection) SetIdentityProvider1(provider js.String) (ret js.Void) {
	bindings.CallRTCPeerConnectionSetIdentityProvider1(
		this.Ref(), js.Pointer(&ret),
		provider.Ref(),
	)

	return
}

// TrySetIdentityProvider1 calls the method "RTCPeerConnection.setIdentityProvider"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TrySetIdentityProvider1(provider js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionSetIdentityProvider1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		provider.Ref(),
	)

	return
}

// HasGetIdentityAssertion returns true if the method "RTCPeerConnection.getIdentityAssertion" exists.
func (this RTCPeerConnection) HasGetIdentityAssertion() bool {
	return js.True == bindings.HasRTCPeerConnectionGetIdentityAssertion(
		this.Ref(),
	)
}

// GetIdentityAssertionFunc returns the method "RTCPeerConnection.getIdentityAssertion".
func (this RTCPeerConnection) GetIdentityAssertionFunc() (fn js.Func[func() js.Promise[js.String]]) {
	return fn.FromRef(
		bindings.RTCPeerConnectionGetIdentityAssertionFunc(
			this.Ref(),
		),
	)
}

// GetIdentityAssertion calls the method "RTCPeerConnection.getIdentityAssertion".
func (this RTCPeerConnection) GetIdentityAssertion() (ret js.Promise[js.String]) {
	bindings.CallRTCPeerConnectionGetIdentityAssertion(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetIdentityAssertion calls the method "RTCPeerConnection.getIdentityAssertion"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCPeerConnection) TryGetIdentityAssertion() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCPeerConnectionGetIdentityAssertion(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type RTCPeerConnectionIceErrorEventInit struct {
	// Address is "RTCPeerConnectionIceErrorEventInit.address"
	//
	// Optional
	Address js.String
	// Port is "RTCPeerConnectionIceErrorEventInit.port"
	//
	// Optional
	//
	// NOTE: FFI_USE_Port MUST be set to true to make this field effective.
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
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "RTCPeerConnectionIceErrorEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "RTCPeerConnectionIceErrorEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new RTCPeerConnectionIceErrorEventInit in the application heap.
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

func NewRTCPeerConnectionIceErrorEvent(typ js.String, eventInitDict RTCPeerConnectionIceErrorEventInit) (ret RTCPeerConnectionIceErrorEvent) {
	ret.ref = bindings.NewRTCPeerConnectionIceErrorEventByRTCPeerConnectionIceErrorEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
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
// It returns ok=false if there is no such property.
func (this RTCPeerConnectionIceErrorEvent) Address() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionIceErrorEventAddress(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Port returns the value of property "RTCPeerConnectionIceErrorEvent.port".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnectionIceErrorEvent) Port() (ret uint16, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionIceErrorEventPort(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Url returns the value of property "RTCPeerConnectionIceErrorEvent.url".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnectionIceErrorEvent) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionIceErrorEventUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ErrorCode returns the value of property "RTCPeerConnectionIceErrorEvent.errorCode".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnectionIceErrorEvent) ErrorCode() (ret uint16, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionIceErrorEventErrorCode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ErrorText returns the value of property "RTCPeerConnectionIceErrorEvent.errorText".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnectionIceErrorEvent) ErrorText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionIceErrorEventErrorText(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "RTCPeerConnectionIceEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "RTCPeerConnectionIceEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new RTCPeerConnectionIceEventInit in the application heap.
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

func NewRTCPeerConnectionIceEvent(typ js.String, eventInitDict RTCPeerConnectionIceEventInit) (ret RTCPeerConnectionIceEvent) {
	ret.ref = bindings.NewRTCPeerConnectionIceEventByRTCPeerConnectionIceEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewRTCPeerConnectionIceEventByRTCPeerConnectionIceEvent1(typ js.String) (ret RTCPeerConnectionIceEvent) {
	ret.ref = bindings.NewRTCPeerConnectionIceEventByRTCPeerConnectionIceEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this RTCPeerConnectionIceEvent) Candidate() (ret RTCIceCandidate, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionIceEventCandidate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Url returns the value of property "RTCPeerConnectionIceEvent.url".
//
// It returns ok=false if there is no such property.
func (this RTCPeerConnectionIceEvent) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCPeerConnectionIceEventUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type RTCPeerConnectionStats struct {
	// DataChannelsOpened is "RTCPeerConnectionStats.dataChannelsOpened"
	//
	// Optional
	//
	// NOTE: FFI_USE_DataChannelsOpened MUST be set to true to make this field effective.
	DataChannelsOpened uint32
	// DataChannelsClosed is "RTCPeerConnectionStats.dataChannelsClosed"
	//
	// Optional
	//
	// NOTE: FFI_USE_DataChannelsClosed MUST be set to true to make this field effective.
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

// New creates a new RTCPeerConnectionStats in the application heap.
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
	//
	// NOTE: FFI_USE_PacketsReceived MUST be set to true to make this field effective.
	PacketsReceived uint64
	// PacketsLost is "RTCReceivedRtpStreamStats.packetsLost"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsLost MUST be set to true to make this field effective.
	PacketsLost int64
	// Jitter is "RTCReceivedRtpStreamStats.jitter"
	//
	// Optional
	//
	// NOTE: FFI_USE_Jitter MUST be set to true to make this field effective.
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

// New creates a new RTCReceivedRtpStreamStats in the application heap.
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
	//
	// NOTE: FFI_USE_RoundTripTime MUST be set to true to make this field effective.
	RoundTripTime float64
	// TotalRoundTripTime is "RTCRemoteInboundRtpStreamStats.totalRoundTripTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalRoundTripTime MUST be set to true to make this field effective.
	TotalRoundTripTime float64
	// FractionLost is "RTCRemoteInboundRtpStreamStats.fractionLost"
	//
	// Optional
	//
	// NOTE: FFI_USE_FractionLost MUST be set to true to make this field effective.
	FractionLost float64
	// RoundTripTimeMeasurements is "RTCRemoteInboundRtpStreamStats.roundTripTimeMeasurements"
	//
	// Optional
	//
	// NOTE: FFI_USE_RoundTripTimeMeasurements MUST be set to true to make this field effective.
	RoundTripTimeMeasurements uint64
	// PacketsReceived is "RTCRemoteInboundRtpStreamStats.packetsReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsReceived MUST be set to true to make this field effective.
	PacketsReceived uint64
	// PacketsLost is "RTCRemoteInboundRtpStreamStats.packetsLost"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsLost MUST be set to true to make this field effective.
	PacketsLost int64
	// Jitter is "RTCRemoteInboundRtpStreamStats.jitter"
	//
	// Optional
	//
	// NOTE: FFI_USE_Jitter MUST be set to true to make this field effective.
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

// New creates a new RTCRemoteInboundRtpStreamStats in the application heap.
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
	//
	// NOTE: FFI_USE_RemoteTimestamp MUST be set to true to make this field effective.
	RemoteTimestamp DOMHighResTimeStamp
	// ReportsSent is "RTCRemoteOutboundRtpStreamStats.reportsSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_ReportsSent MUST be set to true to make this field effective.
	ReportsSent uint64
	// RoundTripTime is "RTCRemoteOutboundRtpStreamStats.roundTripTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_RoundTripTime MUST be set to true to make this field effective.
	RoundTripTime float64
	// TotalRoundTripTime is "RTCRemoteOutboundRtpStreamStats.totalRoundTripTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalRoundTripTime MUST be set to true to make this field effective.
	TotalRoundTripTime float64
	// RoundTripTimeMeasurements is "RTCRemoteOutboundRtpStreamStats.roundTripTimeMeasurements"
	//
	// Optional
	//
	// NOTE: FFI_USE_RoundTripTimeMeasurements MUST be set to true to make this field effective.
	RoundTripTimeMeasurements uint64
	// PacketsSent is "RTCRemoteOutboundRtpStreamStats.packetsSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsSent MUST be set to true to make this field effective.
	PacketsSent uint64
	// BytesSent is "RTCRemoteOutboundRtpStreamStats.bytesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesSent MUST be set to true to make this field effective.
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

// New creates a new RTCRemoteOutboundRtpStreamStats in the application heap.
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
	//
	// NOTE: FFI_USE_Channels MUST be set to true to make this field effective.
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

// New creates a new RTCRtpCodec in the application heap.
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

// New creates a new RTCRtpCodingParameters in the application heap.
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
	//
	// NOTE: Rtcp.FFI_USE MUST be set to true to get Rtcp used.
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

// New creates a new RTCRtpParameters in the application heap.
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
// It returns ok=false if there is no such property.
func (this RTCRtpScriptTransformer) Readable() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetRTCRtpScriptTransformerReadable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Writable returns the value of property "RTCRtpScriptTransformer.writable".
//
// It returns ok=false if there is no such property.
func (this RTCRtpScriptTransformer) Writable() (ret WritableStream, ok bool) {
	ok = js.True == bindings.GetRTCRtpScriptTransformerWritable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Options returns the value of property "RTCRtpScriptTransformer.options".
//
// It returns ok=false if there is no such property.
func (this RTCRtpScriptTransformer) Options() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetRTCRtpScriptTransformerOptions(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGenerateKeyFrame returns true if the method "RTCRtpScriptTransformer.generateKeyFrame" exists.
func (this RTCRtpScriptTransformer) HasGenerateKeyFrame() bool {
	return js.True == bindings.HasRTCRtpScriptTransformerGenerateKeyFrame(
		this.Ref(),
	)
}

// GenerateKeyFrameFunc returns the method "RTCRtpScriptTransformer.generateKeyFrame".
func (this RTCRtpScriptTransformer) GenerateKeyFrameFunc() (fn js.Func[func(rid js.String) js.Promise[js.BigInt[uint64]]]) {
	return fn.FromRef(
		bindings.RTCRtpScriptTransformerGenerateKeyFrameFunc(
			this.Ref(),
		),
	)
}

// GenerateKeyFrame calls the method "RTCRtpScriptTransformer.generateKeyFrame".
func (this RTCRtpScriptTransformer) GenerateKeyFrame(rid js.String) (ret js.Promise[js.BigInt[uint64]]) {
	bindings.CallRTCRtpScriptTransformerGenerateKeyFrame(
		this.Ref(), js.Pointer(&ret),
		rid.Ref(),
	)

	return
}

// TryGenerateKeyFrame calls the method "RTCRtpScriptTransformer.generateKeyFrame"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpScriptTransformer) TryGenerateKeyFrame(rid js.String) (ret js.Promise[js.BigInt[uint64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpScriptTransformerGenerateKeyFrame(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rid.Ref(),
	)

	return
}

// HasGenerateKeyFrame1 returns true if the method "RTCRtpScriptTransformer.generateKeyFrame" exists.
func (this RTCRtpScriptTransformer) HasGenerateKeyFrame1() bool {
	return js.True == bindings.HasRTCRtpScriptTransformerGenerateKeyFrame1(
		this.Ref(),
	)
}

// GenerateKeyFrame1Func returns the method "RTCRtpScriptTransformer.generateKeyFrame".
func (this RTCRtpScriptTransformer) GenerateKeyFrame1Func() (fn js.Func[func() js.Promise[js.BigInt[uint64]]]) {
	return fn.FromRef(
		bindings.RTCRtpScriptTransformerGenerateKeyFrame1Func(
			this.Ref(),
		),
	)
}

// GenerateKeyFrame1 calls the method "RTCRtpScriptTransformer.generateKeyFrame".
func (this RTCRtpScriptTransformer) GenerateKeyFrame1() (ret js.Promise[js.BigInt[uint64]]) {
	bindings.CallRTCRtpScriptTransformerGenerateKeyFrame1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGenerateKeyFrame1 calls the method "RTCRtpScriptTransformer.generateKeyFrame"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpScriptTransformer) TryGenerateKeyFrame1() (ret js.Promise[js.BigInt[uint64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpScriptTransformerGenerateKeyFrame1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSendKeyFrameRequest returns true if the method "RTCRtpScriptTransformer.sendKeyFrameRequest" exists.
func (this RTCRtpScriptTransformer) HasSendKeyFrameRequest() bool {
	return js.True == bindings.HasRTCRtpScriptTransformerSendKeyFrameRequest(
		this.Ref(),
	)
}

// SendKeyFrameRequestFunc returns the method "RTCRtpScriptTransformer.sendKeyFrameRequest".
func (this RTCRtpScriptTransformer) SendKeyFrameRequestFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RTCRtpScriptTransformerSendKeyFrameRequestFunc(
			this.Ref(),
		),
	)
}

// SendKeyFrameRequest calls the method "RTCRtpScriptTransformer.sendKeyFrameRequest".
func (this RTCRtpScriptTransformer) SendKeyFrameRequest() (ret js.Promise[js.Void]) {
	bindings.CallRTCRtpScriptTransformerSendKeyFrameRequest(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySendKeyFrameRequest calls the method "RTCRtpScriptTransformer.sendKeyFrameRequest"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RTCRtpScriptTransformer) TrySendKeyFrameRequest() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCRtpScriptTransformerSendKeyFrameRequest(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// New creates a new RTCRtpStreamStats in the application heap.
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
	//
	// NOTE: FFI_USE_PacketsSent MUST be set to true to make this field effective.
	PacketsSent uint64
	// BytesSent is "RTCSentRtpStreamStats.bytesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesSent MUST be set to true to make this field effective.
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

// New creates a new RTCSentRtpStreamStats in the application heap.
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

// New creates a new RTCStats in the application heap.
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
