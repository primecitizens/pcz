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

type RTCAudioPlayoutStats struct {
	// Kind is "RTCAudioPlayoutStats.kind"
	//
	// Required
	Kind js.String
	// SynthesizedSamplesDuration is "RTCAudioPlayoutStats.synthesizedSamplesDuration"
	//
	// Optional
	//
	// NOTE: FFI_USE_SynthesizedSamplesDuration MUST be set to true to make this field effective.
	SynthesizedSamplesDuration float64
	// SynthesizedSamplesEvents is "RTCAudioPlayoutStats.synthesizedSamplesEvents"
	//
	// Optional
	//
	// NOTE: FFI_USE_SynthesizedSamplesEvents MUST be set to true to make this field effective.
	SynthesizedSamplesEvents uint32
	// TotalSamplesDuration is "RTCAudioPlayoutStats.totalSamplesDuration"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalSamplesDuration MUST be set to true to make this field effective.
	TotalSamplesDuration float64
	// TotalPlayoutDelay is "RTCAudioPlayoutStats.totalPlayoutDelay"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalPlayoutDelay MUST be set to true to make this field effective.
	TotalPlayoutDelay float64
	// TotalSamplesCount is "RTCAudioPlayoutStats.totalSamplesCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalSamplesCount MUST be set to true to make this field effective.
	TotalSamplesCount uint64
	// Timestamp is "RTCAudioPlayoutStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCAudioPlayoutStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCAudioPlayoutStats.id"
	//
	// Required
	Id js.String

	FFI_USE_SynthesizedSamplesDuration bool // for SynthesizedSamplesDuration.
	FFI_USE_SynthesizedSamplesEvents   bool // for SynthesizedSamplesEvents.
	FFI_USE_TotalSamplesDuration       bool // for TotalSamplesDuration.
	FFI_USE_TotalPlayoutDelay          bool // for TotalPlayoutDelay.
	FFI_USE_TotalSamplesCount          bool // for TotalSamplesCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCAudioPlayoutStats with all fields set.
func (p RTCAudioPlayoutStats) FromRef(ref js.Ref) RTCAudioPlayoutStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCAudioPlayoutStats in the application heap.
func (p RTCAudioPlayoutStats) New() js.Ref {
	return bindings.RTCAudioPlayoutStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCAudioPlayoutStats) UpdateFrom(ref js.Ref) {
	bindings.RTCAudioPlayoutStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCAudioPlayoutStats) Update(ref js.Ref) {
	bindings.RTCAudioPlayoutStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCAudioSourceStats struct {
	// AudioLevel is "RTCAudioSourceStats.audioLevel"
	//
	// Optional
	//
	// NOTE: FFI_USE_AudioLevel MUST be set to true to make this field effective.
	AudioLevel float64
	// TotalAudioEnergy is "RTCAudioSourceStats.totalAudioEnergy"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalAudioEnergy MUST be set to true to make this field effective.
	TotalAudioEnergy float64
	// TotalSamplesDuration is "RTCAudioSourceStats.totalSamplesDuration"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalSamplesDuration MUST be set to true to make this field effective.
	TotalSamplesDuration float64
	// EchoReturnLoss is "RTCAudioSourceStats.echoReturnLoss"
	//
	// Optional
	//
	// NOTE: FFI_USE_EchoReturnLoss MUST be set to true to make this field effective.
	EchoReturnLoss float64
	// EchoReturnLossEnhancement is "RTCAudioSourceStats.echoReturnLossEnhancement"
	//
	// Optional
	//
	// NOTE: FFI_USE_EchoReturnLossEnhancement MUST be set to true to make this field effective.
	EchoReturnLossEnhancement float64
	// DroppedSamplesDuration is "RTCAudioSourceStats.droppedSamplesDuration"
	//
	// Optional
	//
	// NOTE: FFI_USE_DroppedSamplesDuration MUST be set to true to make this field effective.
	DroppedSamplesDuration float64
	// DroppedSamplesEvents is "RTCAudioSourceStats.droppedSamplesEvents"
	//
	// Optional
	//
	// NOTE: FFI_USE_DroppedSamplesEvents MUST be set to true to make this field effective.
	DroppedSamplesEvents uint32
	// TotalCaptureDelay is "RTCAudioSourceStats.totalCaptureDelay"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalCaptureDelay MUST be set to true to make this field effective.
	TotalCaptureDelay float64
	// TotalSamplesCaptured is "RTCAudioSourceStats.totalSamplesCaptured"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalSamplesCaptured MUST be set to true to make this field effective.
	TotalSamplesCaptured uint64
	// TrackIdentifier is "RTCAudioSourceStats.trackIdentifier"
	//
	// Required
	TrackIdentifier js.String
	// Kind is "RTCAudioSourceStats.kind"
	//
	// Required
	Kind js.String
	// Timestamp is "RTCAudioSourceStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCAudioSourceStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCAudioSourceStats.id"
	//
	// Required
	Id js.String

	FFI_USE_AudioLevel                bool // for AudioLevel.
	FFI_USE_TotalAudioEnergy          bool // for TotalAudioEnergy.
	FFI_USE_TotalSamplesDuration      bool // for TotalSamplesDuration.
	FFI_USE_EchoReturnLoss            bool // for EchoReturnLoss.
	FFI_USE_EchoReturnLossEnhancement bool // for EchoReturnLossEnhancement.
	FFI_USE_DroppedSamplesDuration    bool // for DroppedSamplesDuration.
	FFI_USE_DroppedSamplesEvents      bool // for DroppedSamplesEvents.
	FFI_USE_TotalCaptureDelay         bool // for TotalCaptureDelay.
	FFI_USE_TotalSamplesCaptured      bool // for TotalSamplesCaptured.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCAudioSourceStats with all fields set.
func (p RTCAudioSourceStats) FromRef(ref js.Ref) RTCAudioSourceStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCAudioSourceStats in the application heap.
func (p RTCAudioSourceStats) New() js.Ref {
	return bindings.RTCAudioSourceStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCAudioSourceStats) UpdateFrom(ref js.Ref) {
	bindings.RTCAudioSourceStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCAudioSourceStats) Update(ref js.Ref) {
	bindings.RTCAudioSourceStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCBundlePolicy uint32

const (
	_ RTCBundlePolicy = iota

	RTCBundlePolicy_BALANCED
	RTCBundlePolicy_MAX_COMPAT
	RTCBundlePolicy_MAX_BUNDLE
)

func (RTCBundlePolicy) FromRef(str js.Ref) RTCBundlePolicy {
	return RTCBundlePolicy(bindings.ConstOfRTCBundlePolicy(str))
}

func (x RTCBundlePolicy) String() (string, bool) {
	switch x {
	case RTCBundlePolicy_BALANCED:
		return "balanced", true
	case RTCBundlePolicy_MAX_COMPAT:
		return "max-compat", true
	case RTCBundlePolicy_MAX_BUNDLE:
		return "max-bundle", true
	default:
		return "", false
	}
}

type RTCDtlsFingerprint struct {
	// Algorithm is "RTCDtlsFingerprint.algorithm"
	//
	// Optional
	Algorithm js.String
	// Value is "RTCDtlsFingerprint.value"
	//
	// Optional
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCDtlsFingerprint with all fields set.
func (p RTCDtlsFingerprint) FromRef(ref js.Ref) RTCDtlsFingerprint {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCDtlsFingerprint in the application heap.
func (p RTCDtlsFingerprint) New() js.Ref {
	return bindings.RTCDtlsFingerprintJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCDtlsFingerprint) UpdateFrom(ref js.Ref) {
	bindings.RTCDtlsFingerprintJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCDtlsFingerprint) Update(ref js.Ref) {
	bindings.RTCDtlsFingerprintJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCCertificate struct {
	ref js.Ref
}

func (this RTCCertificate) Once() RTCCertificate {
	this.Ref().Once()
	return this
}

func (this RTCCertificate) Ref() js.Ref {
	return this.ref
}

func (this RTCCertificate) FromRef(ref js.Ref) RTCCertificate {
	this.ref = ref
	return this
}

func (this RTCCertificate) Free() {
	this.Ref().Free()
}

// Expires returns the value of property "RTCCertificate.expires".
//
// The returned bool will be false if there is no such property.
func (this RTCCertificate) Expires() (EpochTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetRTCCertificateExpires(
		this.Ref(), js.Pointer(&_ok),
	)
	return EpochTimeStamp(_ret), _ok
}

// GetFingerprints calls the method "RTCCertificate.getFingerprints".
//
// The returned bool will be false if there is no such method.
func (this RTCCertificate) GetFingerprints() (js.Array[RTCDtlsFingerprint], bool) {
	var _ok bool
	_ret := bindings.CallRTCCertificateGetFingerprints(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[RTCDtlsFingerprint]{}.FromRef(_ret), _ok
}

// GetFingerprintsFunc returns the method "RTCCertificate.getFingerprints".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCCertificate) GetFingerprintsFunc() (fn js.Func[func() js.Array[RTCDtlsFingerprint]]) {
	return fn.FromRef(
		bindings.RTCCertificateGetFingerprintsFunc(
			this.Ref(),
		),
	)
}

type RTCCertificateExpiration struct {
	// Expires is "RTCCertificateExpiration.expires"
	//
	// Optional
	//
	// NOTE: FFI_USE_Expires MUST be set to true to make this field effective.
	Expires uint64

	FFI_USE_Expires bool // for Expires.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCCertificateExpiration with all fields set.
func (p RTCCertificateExpiration) FromRef(ref js.Ref) RTCCertificateExpiration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCCertificateExpiration in the application heap.
func (p RTCCertificateExpiration) New() js.Ref {
	return bindings.RTCCertificateExpirationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCCertificateExpiration) UpdateFrom(ref js.Ref) {
	bindings.RTCCertificateExpirationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCCertificateExpiration) Update(ref js.Ref) {
	bindings.RTCCertificateExpirationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCCertificateStats struct {
	// Fingerprint is "RTCCertificateStats.fingerprint"
	//
	// Required
	Fingerprint js.String
	// FingerprintAlgorithm is "RTCCertificateStats.fingerprintAlgorithm"
	//
	// Required
	FingerprintAlgorithm js.String
	// Base64Certificate is "RTCCertificateStats.base64Certificate"
	//
	// Required
	Base64Certificate js.String
	// IssuerCertificateId is "RTCCertificateStats.issuerCertificateId"
	//
	// Optional
	IssuerCertificateId js.String
	// Timestamp is "RTCCertificateStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCCertificateStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCCertificateStats.id"
	//
	// Required
	Id js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCCertificateStats with all fields set.
func (p RTCCertificateStats) FromRef(ref js.Ref) RTCCertificateStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCCertificateStats in the application heap.
func (p RTCCertificateStats) New() js.Ref {
	return bindings.RTCCertificateStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCCertificateStats) UpdateFrom(ref js.Ref) {
	bindings.RTCCertificateStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCCertificateStats) Update(ref js.Ref) {
	bindings.RTCCertificateStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCCodecStats struct {
	// PayloadType is "RTCCodecStats.payloadType"
	//
	// Required
	PayloadType uint32
	// TransportId is "RTCCodecStats.transportId"
	//
	// Required
	TransportId js.String
	// MimeType is "RTCCodecStats.mimeType"
	//
	// Required
	MimeType js.String
	// ClockRate is "RTCCodecStats.clockRate"
	//
	// Optional
	//
	// NOTE: FFI_USE_ClockRate MUST be set to true to make this field effective.
	ClockRate uint32
	// Channels is "RTCCodecStats.channels"
	//
	// Optional
	//
	// NOTE: FFI_USE_Channels MUST be set to true to make this field effective.
	Channels uint32
	// SdpFmtpLine is "RTCCodecStats.sdpFmtpLine"
	//
	// Optional
	SdpFmtpLine js.String
	// Timestamp is "RTCCodecStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCCodecStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCCodecStats.id"
	//
	// Required
	Id js.String

	FFI_USE_ClockRate bool // for ClockRate.
	FFI_USE_Channels  bool // for Channels.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCCodecStats with all fields set.
func (p RTCCodecStats) FromRef(ref js.Ref) RTCCodecStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCCodecStats in the application heap.
func (p RTCCodecStats) New() js.Ref {
	return bindings.RTCCodecStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCCodecStats) UpdateFrom(ref js.Ref) {
	bindings.RTCCodecStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCCodecStats) Update(ref js.Ref) {
	bindings.RTCCodecStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCIceServer struct {
	// Urls is "RTCIceServer.urls"
	//
	// Required
	Urls OneOf_String_ArrayString
	// Username is "RTCIceServer.username"
	//
	// Optional
	Username js.String
	// Credential is "RTCIceServer.credential"
	//
	// Optional
	Credential js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCIceServer with all fields set.
func (p RTCIceServer) FromRef(ref js.Ref) RTCIceServer {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCIceServer in the application heap.
func (p RTCIceServer) New() js.Ref {
	return bindings.RTCIceServerJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCIceServer) UpdateFrom(ref js.Ref) {
	bindings.RTCIceServerJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCIceServer) Update(ref js.Ref) {
	bindings.RTCIceServerJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCIceTransportPolicy uint32

const (
	_ RTCIceTransportPolicy = iota

	RTCIceTransportPolicy_RELAY
	RTCIceTransportPolicy_ALL
)

func (RTCIceTransportPolicy) FromRef(str js.Ref) RTCIceTransportPolicy {
	return RTCIceTransportPolicy(bindings.ConstOfRTCIceTransportPolicy(str))
}

func (x RTCIceTransportPolicy) String() (string, bool) {
	switch x {
	case RTCIceTransportPolicy_RELAY:
		return "relay", true
	case RTCIceTransportPolicy_ALL:
		return "all", true
	default:
		return "", false
	}
}

type RTCRtcpMuxPolicy uint32

const (
	_ RTCRtcpMuxPolicy = iota

	RTCRtcpMuxPolicy_REQUIRE
)

func (RTCRtcpMuxPolicy) FromRef(str js.Ref) RTCRtcpMuxPolicy {
	return RTCRtcpMuxPolicy(bindings.ConstOfRTCRtcpMuxPolicy(str))
}

func (x RTCRtcpMuxPolicy) String() (string, bool) {
	switch x {
	case RTCRtcpMuxPolicy_REQUIRE:
		return "require", true
	default:
		return "", false
	}
}

type RTCConfiguration struct {
	// IceServers is "RTCConfiguration.iceServers"
	//
	// Optional, defaults to [].
	IceServers js.Array[RTCIceServer]
	// IceTransportPolicy is "RTCConfiguration.iceTransportPolicy"
	//
	// Optional, defaults to "all".
	IceTransportPolicy RTCIceTransportPolicy
	// BundlePolicy is "RTCConfiguration.bundlePolicy"
	//
	// Optional, defaults to "balanced".
	BundlePolicy RTCBundlePolicy
	// RtcpMuxPolicy is "RTCConfiguration.rtcpMuxPolicy"
	//
	// Optional, defaults to "require".
	RtcpMuxPolicy RTCRtcpMuxPolicy
	// Certificates is "RTCConfiguration.certificates"
	//
	// Optional, defaults to [].
	Certificates js.Array[RTCCertificate]
	// IceCandidatePoolSize is "RTCConfiguration.iceCandidatePoolSize"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_IceCandidatePoolSize MUST be set to true to make this field effective.
	IceCandidatePoolSize uint8
	// PeerIdentity is "RTCConfiguration.peerIdentity"
	//
	// Optional
	PeerIdentity js.String

	FFI_USE_IceCandidatePoolSize bool // for IceCandidatePoolSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCConfiguration with all fields set.
func (p RTCConfiguration) FromRef(ref js.Ref) RTCConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCConfiguration in the application heap.
func (p RTCConfiguration) New() js.Ref {
	return bindings.RTCConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCConfiguration) UpdateFrom(ref js.Ref) {
	bindings.RTCConfigurationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCConfiguration) Update(ref js.Ref) {
	bindings.RTCConfigurationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCDTMFSender struct {
	EventTarget
}

func (this RTCDTMFSender) Once() RTCDTMFSender {
	this.Ref().Once()
	return this
}

func (this RTCDTMFSender) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this RTCDTMFSender) FromRef(ref js.Ref) RTCDTMFSender {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this RTCDTMFSender) Free() {
	this.Ref().Free()
}

// CanInsertDTMF returns the value of property "RTCDTMFSender.canInsertDTMF".
//
// The returned bool will be false if there is no such property.
func (this RTCDTMFSender) CanInsertDTMF() (bool, bool) {
	var _ok bool
	_ret := bindings.GetRTCDTMFSenderCanInsertDTMF(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ToneBuffer returns the value of property "RTCDTMFSender.toneBuffer".
//
// The returned bool will be false if there is no such property.
func (this RTCDTMFSender) ToneBuffer() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCDTMFSenderToneBuffer(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// InsertDTMF calls the method "RTCDTMFSender.insertDTMF".
//
// The returned bool will be false if there is no such method.
func (this RTCDTMFSender) InsertDTMF(tones js.String, duration uint32, interToneGap uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCDTMFSenderInsertDTMF(
		this.Ref(), js.Pointer(&_ok),
		tones.Ref(),
		uint32(duration),
		uint32(interToneGap),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InsertDTMFFunc returns the method "RTCDTMFSender.insertDTMF".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCDTMFSender) InsertDTMFFunc() (fn js.Func[func(tones js.String, duration uint32, interToneGap uint32)]) {
	return fn.FromRef(
		bindings.RTCDTMFSenderInsertDTMFFunc(
			this.Ref(),
		),
	)
}

// InsertDTMF1 calls the method "RTCDTMFSender.insertDTMF".
//
// The returned bool will be false if there is no such method.
func (this RTCDTMFSender) InsertDTMF1(tones js.String, duration uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCDTMFSenderInsertDTMF1(
		this.Ref(), js.Pointer(&_ok),
		tones.Ref(),
		uint32(duration),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InsertDTMF1Func returns the method "RTCDTMFSender.insertDTMF".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCDTMFSender) InsertDTMF1Func() (fn js.Func[func(tones js.String, duration uint32)]) {
	return fn.FromRef(
		bindings.RTCDTMFSenderInsertDTMF1Func(
			this.Ref(),
		),
	)
}

// InsertDTMF2 calls the method "RTCDTMFSender.insertDTMF".
//
// The returned bool will be false if there is no such method.
func (this RTCDTMFSender) InsertDTMF2(tones js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCDTMFSenderInsertDTMF2(
		this.Ref(), js.Pointer(&_ok),
		tones.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InsertDTMF2Func returns the method "RTCDTMFSender.insertDTMF".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCDTMFSender) InsertDTMF2Func() (fn js.Func[func(tones js.String)]) {
	return fn.FromRef(
		bindings.RTCDTMFSenderInsertDTMF2Func(
			this.Ref(),
		),
	)
}

type RTCDTMFToneChangeEventInit struct {
	// Tone is "RTCDTMFToneChangeEventInit.tone"
	//
	// Optional, defaults to "".
	Tone js.String
	// Bubbles is "RTCDTMFToneChangeEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "RTCDTMFToneChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "RTCDTMFToneChangeEventInit.composed"
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

// FromRef calls UpdateFrom and returns a RTCDTMFToneChangeEventInit with all fields set.
func (p RTCDTMFToneChangeEventInit) FromRef(ref js.Ref) RTCDTMFToneChangeEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCDTMFToneChangeEventInit in the application heap.
func (p RTCDTMFToneChangeEventInit) New() js.Ref {
	return bindings.RTCDTMFToneChangeEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCDTMFToneChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.RTCDTMFToneChangeEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCDTMFToneChangeEventInit) Update(ref js.Ref) {
	bindings.RTCDTMFToneChangeEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewRTCDTMFToneChangeEvent(typ js.String, eventInitDict RTCDTMFToneChangeEventInit) RTCDTMFToneChangeEvent {
	return RTCDTMFToneChangeEvent{}.FromRef(
		bindings.NewRTCDTMFToneChangeEventByRTCDTMFToneChangeEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewRTCDTMFToneChangeEventByRTCDTMFToneChangeEvent1(typ js.String) RTCDTMFToneChangeEvent {
	return RTCDTMFToneChangeEvent{}.FromRef(
		bindings.NewRTCDTMFToneChangeEventByRTCDTMFToneChangeEvent1(
			typ.Ref()),
	)
}

type RTCDTMFToneChangeEvent struct {
	Event
}

func (this RTCDTMFToneChangeEvent) Once() RTCDTMFToneChangeEvent {
	this.Ref().Once()
	return this
}

func (this RTCDTMFToneChangeEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this RTCDTMFToneChangeEvent) FromRef(ref js.Ref) RTCDTMFToneChangeEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this RTCDTMFToneChangeEvent) Free() {
	this.Ref().Free()
}

// Tone returns the value of property "RTCDTMFToneChangeEvent.tone".
//
// The returned bool will be false if there is no such property.
func (this RTCDTMFToneChangeEvent) Tone() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCDTMFToneChangeEventTone(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type RTCDataChannelState uint32

const (
	_ RTCDataChannelState = iota

	RTCDataChannelState_CONNECTING
	RTCDataChannelState_OPEN
	RTCDataChannelState_CLOSING
	RTCDataChannelState_CLOSED
)

func (RTCDataChannelState) FromRef(str js.Ref) RTCDataChannelState {
	return RTCDataChannelState(bindings.ConstOfRTCDataChannelState(str))
}

func (x RTCDataChannelState) String() (string, bool) {
	switch x {
	case RTCDataChannelState_CONNECTING:
		return "connecting", true
	case RTCDataChannelState_OPEN:
		return "open", true
	case RTCDataChannelState_CLOSING:
		return "closing", true
	case RTCDataChannelState_CLOSED:
		return "closed", true
	default:
		return "", false
	}
}

type RTCPriorityType uint32

const (
	_ RTCPriorityType = iota

	RTCPriorityType_VERY_LOW
	RTCPriorityType_LOW
	RTCPriorityType_MEDIUM
	RTCPriorityType_HIGH
)

func (RTCPriorityType) FromRef(str js.Ref) RTCPriorityType {
	return RTCPriorityType(bindings.ConstOfRTCPriorityType(str))
}

func (x RTCPriorityType) String() (string, bool) {
	switch x {
	case RTCPriorityType_VERY_LOW:
		return "very-low", true
	case RTCPriorityType_LOW:
		return "low", true
	case RTCPriorityType_MEDIUM:
		return "medium", true
	case RTCPriorityType_HIGH:
		return "high", true
	default:
		return "", false
	}
}

type RTCDataChannel struct {
	EventTarget
}

func (this RTCDataChannel) Once() RTCDataChannel {
	this.Ref().Once()
	return this
}

func (this RTCDataChannel) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this RTCDataChannel) FromRef(ref js.Ref) RTCDataChannel {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this RTCDataChannel) Free() {
	this.Ref().Free()
}

// Label returns the value of property "RTCDataChannel.label".
//
// The returned bool will be false if there is no such property.
func (this RTCDataChannel) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCDataChannelLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Ordered returns the value of property "RTCDataChannel.ordered".
//
// The returned bool will be false if there is no such property.
func (this RTCDataChannel) Ordered() (bool, bool) {
	var _ok bool
	_ret := bindings.GetRTCDataChannelOrdered(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// MaxPacketLifeTime returns the value of property "RTCDataChannel.maxPacketLifeTime".
//
// The returned bool will be false if there is no such property.
func (this RTCDataChannel) MaxPacketLifeTime() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetRTCDataChannelMaxPacketLifeTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// MaxRetransmits returns the value of property "RTCDataChannel.maxRetransmits".
//
// The returned bool will be false if there is no such property.
func (this RTCDataChannel) MaxRetransmits() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetRTCDataChannelMaxRetransmits(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Protocol returns the value of property "RTCDataChannel.protocol".
//
// The returned bool will be false if there is no such property.
func (this RTCDataChannel) Protocol() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCDataChannelProtocol(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Negotiated returns the value of property "RTCDataChannel.negotiated".
//
// The returned bool will be false if there is no such property.
func (this RTCDataChannel) Negotiated() (bool, bool) {
	var _ok bool
	_ret := bindings.GetRTCDataChannelNegotiated(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Id returns the value of property "RTCDataChannel.id".
//
// The returned bool will be false if there is no such property.
func (this RTCDataChannel) Id() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetRTCDataChannelId(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// ReadyState returns the value of property "RTCDataChannel.readyState".
//
// The returned bool will be false if there is no such property.
func (this RTCDataChannel) ReadyState() (RTCDataChannelState, bool) {
	var _ok bool
	_ret := bindings.GetRTCDataChannelReadyState(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCDataChannelState(_ret), _ok
}

// BufferedAmount returns the value of property "RTCDataChannel.bufferedAmount".
//
// The returned bool will be false if there is no such property.
func (this RTCDataChannel) BufferedAmount() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetRTCDataChannelBufferedAmount(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// BufferedAmountLowThreshold returns the value of property "RTCDataChannel.bufferedAmountLowThreshold".
//
// The returned bool will be false if there is no such property.
func (this RTCDataChannel) BufferedAmountLowThreshold() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetRTCDataChannelBufferedAmountLowThreshold(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// BufferedAmountLowThreshold sets the value of property "RTCDataChannel.bufferedAmountLowThreshold" to val.
//
// It returns false if the property cannot be set.
func (this RTCDataChannel) SetBufferedAmountLowThreshold(val uint32) bool {
	return js.True == bindings.SetRTCDataChannelBufferedAmountLowThreshold(
		this.Ref(),
		uint32(val),
	)
}

// BinaryType returns the value of property "RTCDataChannel.binaryType".
//
// The returned bool will be false if there is no such property.
func (this RTCDataChannel) BinaryType() (BinaryType, bool) {
	var _ok bool
	_ret := bindings.GetRTCDataChannelBinaryType(
		this.Ref(), js.Pointer(&_ok),
	)
	return BinaryType(_ret), _ok
}

// BinaryType sets the value of property "RTCDataChannel.binaryType" to val.
//
// It returns false if the property cannot be set.
func (this RTCDataChannel) SetBinaryType(val BinaryType) bool {
	return js.True == bindings.SetRTCDataChannelBinaryType(
		this.Ref(),
		uint32(val),
	)
}

// Priority returns the value of property "RTCDataChannel.priority".
//
// The returned bool will be false if there is no such property.
func (this RTCDataChannel) Priority() (RTCPriorityType, bool) {
	var _ok bool
	_ret := bindings.GetRTCDataChannelPriority(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCPriorityType(_ret), _ok
}

// Close calls the method "RTCDataChannel.close".
//
// The returned bool will be false if there is no such method.
func (this RTCDataChannel) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCDataChannelClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "RTCDataChannel.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCDataChannel) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RTCDataChannelCloseFunc(
			this.Ref(),
		),
	)
}

// Send calls the method "RTCDataChannel.send".
//
// The returned bool will be false if there is no such method.
func (this RTCDataChannel) Send(data js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCDataChannelSend(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SendFunc returns the method "RTCDataChannel.send".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCDataChannel) SendFunc() (fn js.Func[func(data js.String)]) {
	return fn.FromRef(
		bindings.RTCDataChannelSendFunc(
			this.Ref(),
		),
	)
}

// Send1 calls the method "RTCDataChannel.send".
//
// The returned bool will be false if there is no such method.
func (this RTCDataChannel) Send1(data Blob) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCDataChannelSend1(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Send1Func returns the method "RTCDataChannel.send".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCDataChannel) Send1Func() (fn js.Func[func(data Blob)]) {
	return fn.FromRef(
		bindings.RTCDataChannelSend1Func(
			this.Ref(),
		),
	)
}

// Send2 calls the method "RTCDataChannel.send".
//
// The returned bool will be false if there is no such method.
func (this RTCDataChannel) Send2(data js.ArrayBuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCDataChannelSend2(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Send2Func returns the method "RTCDataChannel.send".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCDataChannel) Send2Func() (fn js.Func[func(data js.ArrayBuffer)]) {
	return fn.FromRef(
		bindings.RTCDataChannelSend2Func(
			this.Ref(),
		),
	)
}

// Send3 calls the method "RTCDataChannel.send".
//
// The returned bool will be false if there is no such method.
func (this RTCDataChannel) Send3(data ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCDataChannelSend3(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Send3Func returns the method "RTCDataChannel.send".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCDataChannel) Send3Func() (fn js.Func[func(data ArrayBufferView)]) {
	return fn.FromRef(
		bindings.RTCDataChannelSend3Func(
			this.Ref(),
		),
	)
}

type RTCDataChannelEventInit struct {
	// Channel is "RTCDataChannelEventInit.channel"
	//
	// Required
	Channel RTCDataChannel
	// Bubbles is "RTCDataChannelEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "RTCDataChannelEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "RTCDataChannelEventInit.composed"
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

// FromRef calls UpdateFrom and returns a RTCDataChannelEventInit with all fields set.
func (p RTCDataChannelEventInit) FromRef(ref js.Ref) RTCDataChannelEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCDataChannelEventInit in the application heap.
func (p RTCDataChannelEventInit) New() js.Ref {
	return bindings.RTCDataChannelEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCDataChannelEventInit) UpdateFrom(ref js.Ref) {
	bindings.RTCDataChannelEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCDataChannelEventInit) Update(ref js.Ref) {
	bindings.RTCDataChannelEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewRTCDataChannelEvent(typ js.String, eventInitDict RTCDataChannelEventInit) RTCDataChannelEvent {
	return RTCDataChannelEvent{}.FromRef(
		bindings.NewRTCDataChannelEventByRTCDataChannelEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type RTCDataChannelEvent struct {
	Event
}

func (this RTCDataChannelEvent) Once() RTCDataChannelEvent {
	this.Ref().Once()
	return this
}

func (this RTCDataChannelEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this RTCDataChannelEvent) FromRef(ref js.Ref) RTCDataChannelEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this RTCDataChannelEvent) Free() {
	this.Ref().Free()
}

// Channel returns the value of property "RTCDataChannelEvent.channel".
//
// The returned bool will be false if there is no such property.
func (this RTCDataChannelEvent) Channel() (RTCDataChannel, bool) {
	var _ok bool
	_ret := bindings.GetRTCDataChannelEventChannel(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCDataChannel{}.FromRef(_ret), _ok
}

type RTCDataChannelInit struct {
	// Ordered is "RTCDataChannelInit.ordered"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Ordered MUST be set to true to make this field effective.
	Ordered bool
	// MaxPacketLifeTime is "RTCDataChannelInit.maxPacketLifeTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxPacketLifeTime MUST be set to true to make this field effective.
	MaxPacketLifeTime uint16
	// MaxRetransmits is "RTCDataChannelInit.maxRetransmits"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxRetransmits MUST be set to true to make this field effective.
	MaxRetransmits uint16
	// Protocol is "RTCDataChannelInit.protocol"
	//
	// Optional, defaults to "".
	Protocol js.String
	// Negotiated is "RTCDataChannelInit.negotiated"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Negotiated MUST be set to true to make this field effective.
	Negotiated bool
	// Id is "RTCDataChannelInit.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id uint16
	// Priority is "RTCDataChannelInit.priority"
	//
	// Optional, defaults to "low".
	Priority RTCPriorityType

	FFI_USE_Ordered           bool // for Ordered.
	FFI_USE_MaxPacketLifeTime bool // for MaxPacketLifeTime.
	FFI_USE_MaxRetransmits    bool // for MaxRetransmits.
	FFI_USE_Negotiated        bool // for Negotiated.
	FFI_USE_Id                bool // for Id.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCDataChannelInit with all fields set.
func (p RTCDataChannelInit) FromRef(ref js.Ref) RTCDataChannelInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCDataChannelInit in the application heap.
func (p RTCDataChannelInit) New() js.Ref {
	return bindings.RTCDataChannelInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCDataChannelInit) UpdateFrom(ref js.Ref) {
	bindings.RTCDataChannelInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCDataChannelInit) Update(ref js.Ref) {
	bindings.RTCDataChannelInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCDataChannelStats struct {
	// Label is "RTCDataChannelStats.label"
	//
	// Optional
	Label js.String
	// Protocol is "RTCDataChannelStats.protocol"
	//
	// Optional
	Protocol js.String
	// DataChannelIdentifier is "RTCDataChannelStats.dataChannelIdentifier"
	//
	// Optional
	//
	// NOTE: FFI_USE_DataChannelIdentifier MUST be set to true to make this field effective.
	DataChannelIdentifier uint16
	// State is "RTCDataChannelStats.state"
	//
	// Required
	State RTCDataChannelState
	// MessagesSent is "RTCDataChannelStats.messagesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_MessagesSent MUST be set to true to make this field effective.
	MessagesSent uint32
	// BytesSent is "RTCDataChannelStats.bytesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesSent MUST be set to true to make this field effective.
	BytesSent uint64
	// MessagesReceived is "RTCDataChannelStats.messagesReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_MessagesReceived MUST be set to true to make this field effective.
	MessagesReceived uint32
	// BytesReceived is "RTCDataChannelStats.bytesReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesReceived MUST be set to true to make this field effective.
	BytesReceived uint64
	// Timestamp is "RTCDataChannelStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCDataChannelStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCDataChannelStats.id"
	//
	// Required
	Id js.String

	FFI_USE_DataChannelIdentifier bool // for DataChannelIdentifier.
	FFI_USE_MessagesSent          bool // for MessagesSent.
	FFI_USE_BytesSent             bool // for BytesSent.
	FFI_USE_MessagesReceived      bool // for MessagesReceived.
	FFI_USE_BytesReceived         bool // for BytesReceived.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCDataChannelStats with all fields set.
func (p RTCDataChannelStats) FromRef(ref js.Ref) RTCDataChannelStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCDataChannelStats in the application heap.
func (p RTCDataChannelStats) New() js.Ref {
	return bindings.RTCDataChannelStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCDataChannelStats) UpdateFrom(ref js.Ref) {
	bindings.RTCDataChannelStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCDataChannelStats) Update(ref js.Ref) {
	bindings.RTCDataChannelStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCDegradationPreference uint32

const (
	_ RTCDegradationPreference = iota

	RTCDegradationPreference_MAINTAIN_FRAMERATE
	RTCDegradationPreference_MAINTAIN_RESOLUTION
	RTCDegradationPreference_BALANCED
)

func (RTCDegradationPreference) FromRef(str js.Ref) RTCDegradationPreference {
	return RTCDegradationPreference(bindings.ConstOfRTCDegradationPreference(str))
}

func (x RTCDegradationPreference) String() (string, bool) {
	switch x {
	case RTCDegradationPreference_MAINTAIN_FRAMERATE:
		return "maintain-framerate", true
	case RTCDegradationPreference_MAINTAIN_RESOLUTION:
		return "maintain-resolution", true
	case RTCDegradationPreference_BALANCED:
		return "balanced", true
	default:
		return "", false
	}
}

type RTCDtlsRole uint32

const (
	_ RTCDtlsRole = iota

	RTCDtlsRole_CLIENT
	RTCDtlsRole_SERVER
	RTCDtlsRole_UNKNOWN
)

func (RTCDtlsRole) FromRef(str js.Ref) RTCDtlsRole {
	return RTCDtlsRole(bindings.ConstOfRTCDtlsRole(str))
}

func (x RTCDtlsRole) String() (string, bool) {
	switch x {
	case RTCDtlsRole_CLIENT:
		return "client", true
	case RTCDtlsRole_SERVER:
		return "server", true
	case RTCDtlsRole_UNKNOWN:
		return "unknown", true
	default:
		return "", false
	}
}

type RTCIceCandidateInit struct {
	// Candidate is "RTCIceCandidateInit.candidate"
	//
	// Optional, defaults to "".
	Candidate js.String
	// SdpMid is "RTCIceCandidateInit.sdpMid"
	//
	// Optional, defaults to null.
	SdpMid js.String
	// SdpMLineIndex is "RTCIceCandidateInit.sdpMLineIndex"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_SdpMLineIndex MUST be set to true to make this field effective.
	SdpMLineIndex uint16
	// UsernameFragment is "RTCIceCandidateInit.usernameFragment"
	//
	// Optional, defaults to null.
	UsernameFragment js.String

	FFI_USE_SdpMLineIndex bool // for SdpMLineIndex.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCIceCandidateInit with all fields set.
func (p RTCIceCandidateInit) FromRef(ref js.Ref) RTCIceCandidateInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCIceCandidateInit in the application heap.
func (p RTCIceCandidateInit) New() js.Ref {
	return bindings.RTCIceCandidateInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCIceCandidateInit) UpdateFrom(ref js.Ref) {
	bindings.RTCIceCandidateInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCIceCandidateInit) Update(ref js.Ref) {
	bindings.RTCIceCandidateInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCIceComponent uint32

const (
	_ RTCIceComponent = iota

	RTCIceComponent_RTP
	RTCIceComponent_RTCP
)

func (RTCIceComponent) FromRef(str js.Ref) RTCIceComponent {
	return RTCIceComponent(bindings.ConstOfRTCIceComponent(str))
}

func (x RTCIceComponent) String() (string, bool) {
	switch x {
	case RTCIceComponent_RTP:
		return "rtp", true
	case RTCIceComponent_RTCP:
		return "rtcp", true
	default:
		return "", false
	}
}

type RTCIceProtocol uint32

const (
	_ RTCIceProtocol = iota

	RTCIceProtocol_UDP
	RTCIceProtocol_TCP
)

func (RTCIceProtocol) FromRef(str js.Ref) RTCIceProtocol {
	return RTCIceProtocol(bindings.ConstOfRTCIceProtocol(str))
}

func (x RTCIceProtocol) String() (string, bool) {
	switch x {
	case RTCIceProtocol_UDP:
		return "udp", true
	case RTCIceProtocol_TCP:
		return "tcp", true
	default:
		return "", false
	}
}

type RTCIceCandidateType uint32

const (
	_ RTCIceCandidateType = iota

	RTCIceCandidateType_HOST
	RTCIceCandidateType_SRFLX
	RTCIceCandidateType_PRFLX
	RTCIceCandidateType_RELAY
)

func (RTCIceCandidateType) FromRef(str js.Ref) RTCIceCandidateType {
	return RTCIceCandidateType(bindings.ConstOfRTCIceCandidateType(str))
}

func (x RTCIceCandidateType) String() (string, bool) {
	switch x {
	case RTCIceCandidateType_HOST:
		return "host", true
	case RTCIceCandidateType_SRFLX:
		return "srflx", true
	case RTCIceCandidateType_PRFLX:
		return "prflx", true
	case RTCIceCandidateType_RELAY:
		return "relay", true
	default:
		return "", false
	}
}

type RTCIceTcpCandidateType uint32

const (
	_ RTCIceTcpCandidateType = iota

	RTCIceTcpCandidateType_ACTIVE
	RTCIceTcpCandidateType_PASSIVE
	RTCIceTcpCandidateType_SO
)

func (RTCIceTcpCandidateType) FromRef(str js.Ref) RTCIceTcpCandidateType {
	return RTCIceTcpCandidateType(bindings.ConstOfRTCIceTcpCandidateType(str))
}

func (x RTCIceTcpCandidateType) String() (string, bool) {
	switch x {
	case RTCIceTcpCandidateType_ACTIVE:
		return "active", true
	case RTCIceTcpCandidateType_PASSIVE:
		return "passive", true
	case RTCIceTcpCandidateType_SO:
		return "so", true
	default:
		return "", false
	}
}

type RTCIceServerTransportProtocol uint32

const (
	_ RTCIceServerTransportProtocol = iota

	RTCIceServerTransportProtocol_UDP
	RTCIceServerTransportProtocol_TCP
	RTCIceServerTransportProtocol_TLS
)

func (RTCIceServerTransportProtocol) FromRef(str js.Ref) RTCIceServerTransportProtocol {
	return RTCIceServerTransportProtocol(bindings.ConstOfRTCIceServerTransportProtocol(str))
}

func (x RTCIceServerTransportProtocol) String() (string, bool) {
	switch x {
	case RTCIceServerTransportProtocol_UDP:
		return "udp", true
	case RTCIceServerTransportProtocol_TCP:
		return "tcp", true
	case RTCIceServerTransportProtocol_TLS:
		return "tls", true
	default:
		return "", false
	}
}

func NewRTCIceCandidate(candidateInitDict RTCIceCandidateInit) RTCIceCandidate {
	return RTCIceCandidate{}.FromRef(
		bindings.NewRTCIceCandidateByRTCIceCandidate(
			js.Pointer(&candidateInitDict)),
	)
}

func NewRTCIceCandidateByRTCIceCandidate1() RTCIceCandidate {
	return RTCIceCandidate{}.FromRef(
		bindings.NewRTCIceCandidateByRTCIceCandidate1(),
	)
}

type RTCIceCandidate struct {
	ref js.Ref
}

func (this RTCIceCandidate) Once() RTCIceCandidate {
	this.Ref().Once()
	return this
}

func (this RTCIceCandidate) Ref() js.Ref {
	return this.ref
}

func (this RTCIceCandidate) FromRef(ref js.Ref) RTCIceCandidate {
	this.ref = ref
	return this
}

func (this RTCIceCandidate) Free() {
	this.Ref().Free()
}

// Candidate returns the value of property "RTCIceCandidate.candidate".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) Candidate() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidateCandidate(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SdpMid returns the value of property "RTCIceCandidate.sdpMid".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) SdpMid() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidateSdpMid(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SdpMLineIndex returns the value of property "RTCIceCandidate.sdpMLineIndex".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) SdpMLineIndex() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidateSdpMLineIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Foundation returns the value of property "RTCIceCandidate.foundation".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) Foundation() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidateFoundation(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Component returns the value of property "RTCIceCandidate.component".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) Component() (RTCIceComponent, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidateComponent(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCIceComponent(_ret), _ok
}

// Priority returns the value of property "RTCIceCandidate.priority".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) Priority() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidatePriority(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Address returns the value of property "RTCIceCandidate.address".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) Address() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidateAddress(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Protocol returns the value of property "RTCIceCandidate.protocol".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) Protocol() (RTCIceProtocol, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidateProtocol(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCIceProtocol(_ret), _ok
}

// Port returns the value of property "RTCIceCandidate.port".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) Port() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidatePort(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Type returns the value of property "RTCIceCandidate.type".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) Type() (RTCIceCandidateType, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidateType(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCIceCandidateType(_ret), _ok
}

// TcpType returns the value of property "RTCIceCandidate.tcpType".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) TcpType() (RTCIceTcpCandidateType, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidateTcpType(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCIceTcpCandidateType(_ret), _ok
}

// RelatedAddress returns the value of property "RTCIceCandidate.relatedAddress".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) RelatedAddress() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidateRelatedAddress(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// RelatedPort returns the value of property "RTCIceCandidate.relatedPort".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) RelatedPort() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidateRelatedPort(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// UsernameFragment returns the value of property "RTCIceCandidate.usernameFragment".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) UsernameFragment() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidateUsernameFragment(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// RelayProtocol returns the value of property "RTCIceCandidate.relayProtocol".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) RelayProtocol() (RTCIceServerTransportProtocol, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidateRelayProtocol(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCIceServerTransportProtocol(_ret), _ok
}

// Url returns the value of property "RTCIceCandidate.url".
//
// The returned bool will be false if there is no such property.
func (this RTCIceCandidate) Url() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceCandidateUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ToJSON calls the method "RTCIceCandidate.toJSON".
//
// The returned bool will be false if there is no such method.
func (this RTCIceCandidate) ToJSON() (RTCIceCandidateInit, bool) {
	var _ret RTCIceCandidateInit
	_ok := js.True == bindings.CallRTCIceCandidateToJSON(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// ToJSONFunc returns the method "RTCIceCandidate.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCIceCandidate) ToJSONFunc() (fn js.Func[func() RTCIceCandidateInit]) {
	return fn.FromRef(
		bindings.RTCIceCandidateToJSONFunc(
			this.Ref(),
		),
	)
}

type RTCIceCandidatePair struct {
	// Local is "RTCIceCandidatePair.local"
	//
	// Optional
	Local RTCIceCandidate
	// Remote is "RTCIceCandidatePair.remote"
	//
	// Optional
	Remote RTCIceCandidate

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCIceCandidatePair with all fields set.
func (p RTCIceCandidatePair) FromRef(ref js.Ref) RTCIceCandidatePair {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCIceCandidatePair in the application heap.
func (p RTCIceCandidatePair) New() js.Ref {
	return bindings.RTCIceCandidatePairJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCIceCandidatePair) UpdateFrom(ref js.Ref) {
	bindings.RTCIceCandidatePairJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCIceCandidatePair) Update(ref js.Ref) {
	bindings.RTCIceCandidatePairJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCIceParameters struct {
	// UsernameFragment is "RTCIceParameters.usernameFragment"
	//
	// Optional
	UsernameFragment js.String
	// Password is "RTCIceParameters.password"
	//
	// Optional
	Password js.String
	// IceLite is "RTCIceParameters.iceLite"
	//
	// Optional
	//
	// NOTE: FFI_USE_IceLite MUST be set to true to make this field effective.
	IceLite bool

	FFI_USE_IceLite bool // for IceLite.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCIceParameters with all fields set.
func (p RTCIceParameters) FromRef(ref js.Ref) RTCIceParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCIceParameters in the application heap.
func (p RTCIceParameters) New() js.Ref {
	return bindings.RTCIceParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCIceParameters) UpdateFrom(ref js.Ref) {
	bindings.RTCIceParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCIceParameters) Update(ref js.Ref) {
	bindings.RTCIceParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCIceGatherOptions struct {
	// GatherPolicy is "RTCIceGatherOptions.gatherPolicy"
	//
	// Optional, defaults to "all".
	GatherPolicy RTCIceTransportPolicy
	// IceServers is "RTCIceGatherOptions.iceServers"
	//
	// Optional
	IceServers js.Array[RTCIceServer]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCIceGatherOptions with all fields set.
func (p RTCIceGatherOptions) FromRef(ref js.Ref) RTCIceGatherOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCIceGatherOptions in the application heap.
func (p RTCIceGatherOptions) New() js.Ref {
	return bindings.RTCIceGatherOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCIceGatherOptions) UpdateFrom(ref js.Ref) {
	bindings.RTCIceGatherOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCIceGatherOptions) Update(ref js.Ref) {
	bindings.RTCIceGatherOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCIceRole uint32

const (
	_ RTCIceRole = iota

	RTCIceRole_UNKNOWN
	RTCIceRole_CONTROLLING
	RTCIceRole_CONTROLLED
)

func (RTCIceRole) FromRef(str js.Ref) RTCIceRole {
	return RTCIceRole(bindings.ConstOfRTCIceRole(str))
}

func (x RTCIceRole) String() (string, bool) {
	switch x {
	case RTCIceRole_UNKNOWN:
		return "unknown", true
	case RTCIceRole_CONTROLLING:
		return "controlling", true
	case RTCIceRole_CONTROLLED:
		return "controlled", true
	default:
		return "", false
	}
}

type RTCIceTransportState uint32

const (
	_ RTCIceTransportState = iota

	RTCIceTransportState_CLOSED
	RTCIceTransportState_FAILED
	RTCIceTransportState_DISCONNECTED
	RTCIceTransportState_NEW
	RTCIceTransportState_CHECKING
	RTCIceTransportState_COMPLETED
	RTCIceTransportState_CONNECTED
)

func (RTCIceTransportState) FromRef(str js.Ref) RTCIceTransportState {
	return RTCIceTransportState(bindings.ConstOfRTCIceTransportState(str))
}

func (x RTCIceTransportState) String() (string, bool) {
	switch x {
	case RTCIceTransportState_CLOSED:
		return "closed", true
	case RTCIceTransportState_FAILED:
		return "failed", true
	case RTCIceTransportState_DISCONNECTED:
		return "disconnected", true
	case RTCIceTransportState_NEW:
		return "new", true
	case RTCIceTransportState_CHECKING:
		return "checking", true
	case RTCIceTransportState_COMPLETED:
		return "completed", true
	case RTCIceTransportState_CONNECTED:
		return "connected", true
	default:
		return "", false
	}
}
