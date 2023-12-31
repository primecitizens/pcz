// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
func (p *RTCAudioPlayoutStats) UpdateFrom(ref js.Ref) {
	bindings.RTCAudioPlayoutStatsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCAudioPlayoutStats) Update(ref js.Ref) {
	bindings.RTCAudioPlayoutStatsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCAudioPlayoutStats) FreeMembers(recursive bool) {
	js.Free(
		p.Kind.Ref(),
		p.Id.Ref(),
	)
	p.Kind = p.Kind.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
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
func (p *RTCAudioSourceStats) UpdateFrom(ref js.Ref) {
	bindings.RTCAudioSourceStatsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCAudioSourceStats) Update(ref js.Ref) {
	bindings.RTCAudioSourceStatsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCAudioSourceStats) FreeMembers(recursive bool) {
	js.Free(
		p.TrackIdentifier.Ref(),
		p.Kind.Ref(),
		p.Id.Ref(),
	)
	p.TrackIdentifier = p.TrackIdentifier.FromRef(js.Undefined)
	p.Kind = p.Kind.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
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
func (p *RTCDtlsFingerprint) UpdateFrom(ref js.Ref) {
	bindings.RTCDtlsFingerprintJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCDtlsFingerprint) Update(ref js.Ref) {
	bindings.RTCDtlsFingerprintJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCDtlsFingerprint) FreeMembers(recursive bool) {
	js.Free(
		p.Algorithm.Ref(),
		p.Value.Ref(),
	)
	p.Algorithm = p.Algorithm.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type RTCCertificate struct {
	ref js.Ref
}

func (this RTCCertificate) Once() RTCCertificate {
	this.ref.Once()
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
	this.ref.Free()
}

// Expires returns the value of property "RTCCertificate.expires".
//
// It returns ok=false if there is no such property.
func (this RTCCertificate) Expires() (ret EpochTimeStamp, ok bool) {
	ok = js.True == bindings.GetRTCCertificateExpires(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetFingerprints returns true if the method "RTCCertificate.getFingerprints" exists.
func (this RTCCertificate) HasFuncGetFingerprints() bool {
	return js.True == bindings.HasFuncRTCCertificateGetFingerprints(
		this.ref,
	)
}

// FuncGetFingerprints returns the method "RTCCertificate.getFingerprints".
func (this RTCCertificate) FuncGetFingerprints() (fn js.Func[func() js.Array[RTCDtlsFingerprint]]) {
	bindings.FuncRTCCertificateGetFingerprints(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFingerprints calls the method "RTCCertificate.getFingerprints".
func (this RTCCertificate) GetFingerprints() (ret js.Array[RTCDtlsFingerprint]) {
	bindings.CallRTCCertificateGetFingerprints(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetFingerprints calls the method "RTCCertificate.getFingerprints"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCCertificate) TryGetFingerprints() (ret js.Array[RTCDtlsFingerprint], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCCertificateGetFingerprints(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
func (p *RTCCertificateExpiration) UpdateFrom(ref js.Ref) {
	bindings.RTCCertificateExpirationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCCertificateExpiration) Update(ref js.Ref) {
	bindings.RTCCertificateExpirationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCCertificateExpiration) FreeMembers(recursive bool) {
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
func (p *RTCCertificateStats) UpdateFrom(ref js.Ref) {
	bindings.RTCCertificateStatsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCCertificateStats) Update(ref js.Ref) {
	bindings.RTCCertificateStatsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCCertificateStats) FreeMembers(recursive bool) {
	js.Free(
		p.Fingerprint.Ref(),
		p.FingerprintAlgorithm.Ref(),
		p.Base64Certificate.Ref(),
		p.IssuerCertificateId.Ref(),
		p.Id.Ref(),
	)
	p.Fingerprint = p.Fingerprint.FromRef(js.Undefined)
	p.FingerprintAlgorithm = p.FingerprintAlgorithm.FromRef(js.Undefined)
	p.Base64Certificate = p.Base64Certificate.FromRef(js.Undefined)
	p.IssuerCertificateId = p.IssuerCertificateId.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
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
func (p *RTCCodecStats) UpdateFrom(ref js.Ref) {
	bindings.RTCCodecStatsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCCodecStats) Update(ref js.Ref) {
	bindings.RTCCodecStatsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCCodecStats) FreeMembers(recursive bool) {
	js.Free(
		p.TransportId.Ref(),
		p.MimeType.Ref(),
		p.SdpFmtpLine.Ref(),
		p.Id.Ref(),
	)
	p.TransportId = p.TransportId.FromRef(js.Undefined)
	p.MimeType = p.MimeType.FromRef(js.Undefined)
	p.SdpFmtpLine = p.SdpFmtpLine.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
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
func (p *RTCIceServer) UpdateFrom(ref js.Ref) {
	bindings.RTCIceServerJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCIceServer) Update(ref js.Ref) {
	bindings.RTCIceServerJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCIceServer) FreeMembers(recursive bool) {
	js.Free(
		p.Urls.Ref(),
		p.Username.Ref(),
		p.Credential.Ref(),
	)
	p.Urls = p.Urls.FromRef(js.Undefined)
	p.Username = p.Username.FromRef(js.Undefined)
	p.Credential = p.Credential.FromRef(js.Undefined)
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
func (p *RTCConfiguration) UpdateFrom(ref js.Ref) {
	bindings.RTCConfigurationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCConfiguration) Update(ref js.Ref) {
	bindings.RTCConfigurationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCConfiguration) FreeMembers(recursive bool) {
	js.Free(
		p.IceServers.Ref(),
		p.Certificates.Ref(),
		p.PeerIdentity.Ref(),
	)
	p.IceServers = p.IceServers.FromRef(js.Undefined)
	p.Certificates = p.Certificates.FromRef(js.Undefined)
	p.PeerIdentity = p.PeerIdentity.FromRef(js.Undefined)
}

type RTCDTMFSender struct {
	EventTarget
}

func (this RTCDTMFSender) Once() RTCDTMFSender {
	this.ref.Once()
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
	this.ref.Free()
}

// CanInsertDTMF returns the value of property "RTCDTMFSender.canInsertDTMF".
//
// It returns ok=false if there is no such property.
func (this RTCDTMFSender) CanInsertDTMF() (ret bool, ok bool) {
	ok = js.True == bindings.GetRTCDTMFSenderCanInsertDTMF(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ToneBuffer returns the value of property "RTCDTMFSender.toneBuffer".
//
// It returns ok=false if there is no such property.
func (this RTCDTMFSender) ToneBuffer() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCDTMFSenderToneBuffer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncInsertDTMF returns true if the method "RTCDTMFSender.insertDTMF" exists.
func (this RTCDTMFSender) HasFuncInsertDTMF() bool {
	return js.True == bindings.HasFuncRTCDTMFSenderInsertDTMF(
		this.ref,
	)
}

// FuncInsertDTMF returns the method "RTCDTMFSender.insertDTMF".
func (this RTCDTMFSender) FuncInsertDTMF() (fn js.Func[func(tones js.String, duration uint32, interToneGap uint32)]) {
	bindings.FuncRTCDTMFSenderInsertDTMF(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertDTMF calls the method "RTCDTMFSender.insertDTMF".
func (this RTCDTMFSender) InsertDTMF(tones js.String, duration uint32, interToneGap uint32) (ret js.Void) {
	bindings.CallRTCDTMFSenderInsertDTMF(
		this.ref, js.Pointer(&ret),
		tones.Ref(),
		uint32(duration),
		uint32(interToneGap),
	)

	return
}

// TryInsertDTMF calls the method "RTCDTMFSender.insertDTMF"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCDTMFSender) TryInsertDTMF(tones js.String, duration uint32, interToneGap uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCDTMFSenderInsertDTMF(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		tones.Ref(),
		uint32(duration),
		uint32(interToneGap),
	)

	return
}

// HasFuncInsertDTMF1 returns true if the method "RTCDTMFSender.insertDTMF" exists.
func (this RTCDTMFSender) HasFuncInsertDTMF1() bool {
	return js.True == bindings.HasFuncRTCDTMFSenderInsertDTMF1(
		this.ref,
	)
}

// FuncInsertDTMF1 returns the method "RTCDTMFSender.insertDTMF".
func (this RTCDTMFSender) FuncInsertDTMF1() (fn js.Func[func(tones js.String, duration uint32)]) {
	bindings.FuncRTCDTMFSenderInsertDTMF1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertDTMF1 calls the method "RTCDTMFSender.insertDTMF".
func (this RTCDTMFSender) InsertDTMF1(tones js.String, duration uint32) (ret js.Void) {
	bindings.CallRTCDTMFSenderInsertDTMF1(
		this.ref, js.Pointer(&ret),
		tones.Ref(),
		uint32(duration),
	)

	return
}

// TryInsertDTMF1 calls the method "RTCDTMFSender.insertDTMF"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCDTMFSender) TryInsertDTMF1(tones js.String, duration uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCDTMFSenderInsertDTMF1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		tones.Ref(),
		uint32(duration),
	)

	return
}

// HasFuncInsertDTMF2 returns true if the method "RTCDTMFSender.insertDTMF" exists.
func (this RTCDTMFSender) HasFuncInsertDTMF2() bool {
	return js.True == bindings.HasFuncRTCDTMFSenderInsertDTMF2(
		this.ref,
	)
}

// FuncInsertDTMF2 returns the method "RTCDTMFSender.insertDTMF".
func (this RTCDTMFSender) FuncInsertDTMF2() (fn js.Func[func(tones js.String)]) {
	bindings.FuncRTCDTMFSenderInsertDTMF2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertDTMF2 calls the method "RTCDTMFSender.insertDTMF".
func (this RTCDTMFSender) InsertDTMF2(tones js.String) (ret js.Void) {
	bindings.CallRTCDTMFSenderInsertDTMF2(
		this.ref, js.Pointer(&ret),
		tones.Ref(),
	)

	return
}

// TryInsertDTMF2 calls the method "RTCDTMFSender.insertDTMF"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCDTMFSender) TryInsertDTMF2(tones js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCDTMFSenderInsertDTMF2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		tones.Ref(),
	)

	return
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
func (p *RTCDTMFToneChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.RTCDTMFToneChangeEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCDTMFToneChangeEventInit) Update(ref js.Ref) {
	bindings.RTCDTMFToneChangeEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCDTMFToneChangeEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Tone.Ref(),
	)
	p.Tone = p.Tone.FromRef(js.Undefined)
}

func NewRTCDTMFToneChangeEvent(typ js.String, eventInitDict RTCDTMFToneChangeEventInit) (ret RTCDTMFToneChangeEvent) {
	ret.ref = bindings.NewRTCDTMFToneChangeEventByRTCDTMFToneChangeEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewRTCDTMFToneChangeEventByRTCDTMFToneChangeEvent1(typ js.String) (ret RTCDTMFToneChangeEvent) {
	ret.ref = bindings.NewRTCDTMFToneChangeEventByRTCDTMFToneChangeEvent1(
		typ.Ref())
	return
}

type RTCDTMFToneChangeEvent struct {
	Event
}

func (this RTCDTMFToneChangeEvent) Once() RTCDTMFToneChangeEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// Tone returns the value of property "RTCDTMFToneChangeEvent.tone".
//
// It returns ok=false if there is no such property.
func (this RTCDTMFToneChangeEvent) Tone() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCDTMFToneChangeEventTone(
		this.ref, js.Pointer(&ret),
	)
	return
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
	this.ref.Once()
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
	this.ref.Free()
}

// Label returns the value of property "RTCDataChannel.label".
//
// It returns ok=false if there is no such property.
func (this RTCDataChannel) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCDataChannelLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Ordered returns the value of property "RTCDataChannel.ordered".
//
// It returns ok=false if there is no such property.
func (this RTCDataChannel) Ordered() (ret bool, ok bool) {
	ok = js.True == bindings.GetRTCDataChannelOrdered(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxPacketLifeTime returns the value of property "RTCDataChannel.maxPacketLifeTime".
//
// It returns ok=false if there is no such property.
func (this RTCDataChannel) MaxPacketLifeTime() (ret uint16, ok bool) {
	ok = js.True == bindings.GetRTCDataChannelMaxPacketLifeTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxRetransmits returns the value of property "RTCDataChannel.maxRetransmits".
//
// It returns ok=false if there is no such property.
func (this RTCDataChannel) MaxRetransmits() (ret uint16, ok bool) {
	ok = js.True == bindings.GetRTCDataChannelMaxRetransmits(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Protocol returns the value of property "RTCDataChannel.protocol".
//
// It returns ok=false if there is no such property.
func (this RTCDataChannel) Protocol() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCDataChannelProtocol(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Negotiated returns the value of property "RTCDataChannel.negotiated".
//
// It returns ok=false if there is no such property.
func (this RTCDataChannel) Negotiated() (ret bool, ok bool) {
	ok = js.True == bindings.GetRTCDataChannelNegotiated(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "RTCDataChannel.id".
//
// It returns ok=false if there is no such property.
func (this RTCDataChannel) Id() (ret uint16, ok bool) {
	ok = js.True == bindings.GetRTCDataChannelId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ReadyState returns the value of property "RTCDataChannel.readyState".
//
// It returns ok=false if there is no such property.
func (this RTCDataChannel) ReadyState() (ret RTCDataChannelState, ok bool) {
	ok = js.True == bindings.GetRTCDataChannelReadyState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BufferedAmount returns the value of property "RTCDataChannel.bufferedAmount".
//
// It returns ok=false if there is no such property.
func (this RTCDataChannel) BufferedAmount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetRTCDataChannelBufferedAmount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BufferedAmountLowThreshold returns the value of property "RTCDataChannel.bufferedAmountLowThreshold".
//
// It returns ok=false if there is no such property.
func (this RTCDataChannel) BufferedAmountLowThreshold() (ret uint32, ok bool) {
	ok = js.True == bindings.GetRTCDataChannelBufferedAmountLowThreshold(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBufferedAmountLowThreshold sets the value of property "RTCDataChannel.bufferedAmountLowThreshold" to val.
//
// It returns false if the property cannot be set.
func (this RTCDataChannel) SetBufferedAmountLowThreshold(val uint32) bool {
	return js.True == bindings.SetRTCDataChannelBufferedAmountLowThreshold(
		this.ref,
		uint32(val),
	)
}

// BinaryType returns the value of property "RTCDataChannel.binaryType".
//
// It returns ok=false if there is no such property.
func (this RTCDataChannel) BinaryType() (ret BinaryType, ok bool) {
	ok = js.True == bindings.GetRTCDataChannelBinaryType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBinaryType sets the value of property "RTCDataChannel.binaryType" to val.
//
// It returns false if the property cannot be set.
func (this RTCDataChannel) SetBinaryType(val BinaryType) bool {
	return js.True == bindings.SetRTCDataChannelBinaryType(
		this.ref,
		uint32(val),
	)
}

// Priority returns the value of property "RTCDataChannel.priority".
//
// It returns ok=false if there is no such property.
func (this RTCDataChannel) Priority() (ret RTCPriorityType, ok bool) {
	ok = js.True == bindings.GetRTCDataChannelPriority(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncClose returns true if the method "RTCDataChannel.close" exists.
func (this RTCDataChannel) HasFuncClose() bool {
	return js.True == bindings.HasFuncRTCDataChannelClose(
		this.ref,
	)
}

// FuncClose returns the method "RTCDataChannel.close".
func (this RTCDataChannel) FuncClose() (fn js.Func[func()]) {
	bindings.FuncRTCDataChannelClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "RTCDataChannel.close".
func (this RTCDataChannel) Close() (ret js.Void) {
	bindings.CallRTCDataChannelClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "RTCDataChannel.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCDataChannel) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCDataChannelClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSend returns true if the method "RTCDataChannel.send" exists.
func (this RTCDataChannel) HasFuncSend() bool {
	return js.True == bindings.HasFuncRTCDataChannelSend(
		this.ref,
	)
}

// FuncSend returns the method "RTCDataChannel.send".
func (this RTCDataChannel) FuncSend() (fn js.Func[func(data js.String)]) {
	bindings.FuncRTCDataChannelSend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Send calls the method "RTCDataChannel.send".
func (this RTCDataChannel) Send(data js.String) (ret js.Void) {
	bindings.CallRTCDataChannelSend(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySend calls the method "RTCDataChannel.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCDataChannel) TrySend(data js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCDataChannelSend(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncSend1 returns true if the method "RTCDataChannel.send" exists.
func (this RTCDataChannel) HasFuncSend1() bool {
	return js.True == bindings.HasFuncRTCDataChannelSend1(
		this.ref,
	)
}

// FuncSend1 returns the method "RTCDataChannel.send".
func (this RTCDataChannel) FuncSend1() (fn js.Func[func(data Blob)]) {
	bindings.FuncRTCDataChannelSend1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Send1 calls the method "RTCDataChannel.send".
func (this RTCDataChannel) Send1(data Blob) (ret js.Void) {
	bindings.CallRTCDataChannelSend1(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySend1 calls the method "RTCDataChannel.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCDataChannel) TrySend1(data Blob) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCDataChannelSend1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncSend2 returns true if the method "RTCDataChannel.send" exists.
func (this RTCDataChannel) HasFuncSend2() bool {
	return js.True == bindings.HasFuncRTCDataChannelSend2(
		this.ref,
	)
}

// FuncSend2 returns the method "RTCDataChannel.send".
func (this RTCDataChannel) FuncSend2() (fn js.Func[func(data js.ArrayBuffer)]) {
	bindings.FuncRTCDataChannelSend2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Send2 calls the method "RTCDataChannel.send".
func (this RTCDataChannel) Send2(data js.ArrayBuffer) (ret js.Void) {
	bindings.CallRTCDataChannelSend2(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySend2 calls the method "RTCDataChannel.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCDataChannel) TrySend2(data js.ArrayBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCDataChannelSend2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncSend3 returns true if the method "RTCDataChannel.send" exists.
func (this RTCDataChannel) HasFuncSend3() bool {
	return js.True == bindings.HasFuncRTCDataChannelSend3(
		this.ref,
	)
}

// FuncSend3 returns the method "RTCDataChannel.send".
func (this RTCDataChannel) FuncSend3() (fn js.Func[func(data js.ArrayBufferView)]) {
	bindings.FuncRTCDataChannelSend3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Send3 calls the method "RTCDataChannel.send".
func (this RTCDataChannel) Send3(data js.ArrayBufferView) (ret js.Void) {
	bindings.CallRTCDataChannelSend3(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySend3 calls the method "RTCDataChannel.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCDataChannel) TrySend3(data js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCDataChannelSend3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
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
func (p *RTCDataChannelEventInit) UpdateFrom(ref js.Ref) {
	bindings.RTCDataChannelEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCDataChannelEventInit) Update(ref js.Ref) {
	bindings.RTCDataChannelEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCDataChannelEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Channel.Ref(),
	)
	p.Channel = p.Channel.FromRef(js.Undefined)
}

func NewRTCDataChannelEvent(typ js.String, eventInitDict RTCDataChannelEventInit) (ret RTCDataChannelEvent) {
	ret.ref = bindings.NewRTCDataChannelEventByRTCDataChannelEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type RTCDataChannelEvent struct {
	Event
}

func (this RTCDataChannelEvent) Once() RTCDataChannelEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// Channel returns the value of property "RTCDataChannelEvent.channel".
//
// It returns ok=false if there is no such property.
func (this RTCDataChannelEvent) Channel() (ret RTCDataChannel, ok bool) {
	ok = js.True == bindings.GetRTCDataChannelEventChannel(
		this.ref, js.Pointer(&ret),
	)
	return
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
func (p *RTCDataChannelInit) UpdateFrom(ref js.Ref) {
	bindings.RTCDataChannelInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCDataChannelInit) Update(ref js.Ref) {
	bindings.RTCDataChannelInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCDataChannelInit) FreeMembers(recursive bool) {
	js.Free(
		p.Protocol.Ref(),
	)
	p.Protocol = p.Protocol.FromRef(js.Undefined)
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
func (p *RTCDataChannelStats) UpdateFrom(ref js.Ref) {
	bindings.RTCDataChannelStatsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCDataChannelStats) Update(ref js.Ref) {
	bindings.RTCDataChannelStatsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCDataChannelStats) FreeMembers(recursive bool) {
	js.Free(
		p.Label.Ref(),
		p.Protocol.Ref(),
		p.Id.Ref(),
	)
	p.Label = p.Label.FromRef(js.Undefined)
	p.Protocol = p.Protocol.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
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
func (p *RTCIceCandidateInit) UpdateFrom(ref js.Ref) {
	bindings.RTCIceCandidateInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCIceCandidateInit) Update(ref js.Ref) {
	bindings.RTCIceCandidateInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCIceCandidateInit) FreeMembers(recursive bool) {
	js.Free(
		p.Candidate.Ref(),
		p.SdpMid.Ref(),
		p.UsernameFragment.Ref(),
	)
	p.Candidate = p.Candidate.FromRef(js.Undefined)
	p.SdpMid = p.SdpMid.FromRef(js.Undefined)
	p.UsernameFragment = p.UsernameFragment.FromRef(js.Undefined)
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

func NewRTCIceCandidate(candidateInitDict RTCIceCandidateInit) (ret RTCIceCandidate) {
	ret.ref = bindings.NewRTCIceCandidateByRTCIceCandidate(
		js.Pointer(&candidateInitDict))
	return
}

func NewRTCIceCandidateByRTCIceCandidate1() (ret RTCIceCandidate) {
	ret.ref = bindings.NewRTCIceCandidateByRTCIceCandidate1()
	return
}

type RTCIceCandidate struct {
	ref js.Ref
}

func (this RTCIceCandidate) Once() RTCIceCandidate {
	this.ref.Once()
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
	this.ref.Free()
}

// Candidate returns the value of property "RTCIceCandidate.candidate".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) Candidate() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidateCandidate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SdpMid returns the value of property "RTCIceCandidate.sdpMid".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) SdpMid() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidateSdpMid(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SdpMLineIndex returns the value of property "RTCIceCandidate.sdpMLineIndex".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) SdpMLineIndex() (ret uint16, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidateSdpMLineIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Foundation returns the value of property "RTCIceCandidate.foundation".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) Foundation() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidateFoundation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Component returns the value of property "RTCIceCandidate.component".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) Component() (ret RTCIceComponent, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidateComponent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Priority returns the value of property "RTCIceCandidate.priority".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) Priority() (ret uint32, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidatePriority(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Address returns the value of property "RTCIceCandidate.address".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) Address() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidateAddress(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Protocol returns the value of property "RTCIceCandidate.protocol".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) Protocol() (ret RTCIceProtocol, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidateProtocol(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Port returns the value of property "RTCIceCandidate.port".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) Port() (ret uint16, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidatePort(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "RTCIceCandidate.type".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) Type() (ret RTCIceCandidateType, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidateType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TcpType returns the value of property "RTCIceCandidate.tcpType".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) TcpType() (ret RTCIceTcpCandidateType, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidateTcpType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RelatedAddress returns the value of property "RTCIceCandidate.relatedAddress".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) RelatedAddress() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidateRelatedAddress(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RelatedPort returns the value of property "RTCIceCandidate.relatedPort".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) RelatedPort() (ret uint16, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidateRelatedPort(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UsernameFragment returns the value of property "RTCIceCandidate.usernameFragment".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) UsernameFragment() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidateUsernameFragment(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RelayProtocol returns the value of property "RTCIceCandidate.relayProtocol".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) RelayProtocol() (ret RTCIceServerTransportProtocol, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidateRelayProtocol(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Url returns the value of property "RTCIceCandidate.url".
//
// It returns ok=false if there is no such property.
func (this RTCIceCandidate) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCIceCandidateUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncToJSON returns true if the method "RTCIceCandidate.toJSON" exists.
func (this RTCIceCandidate) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncRTCIceCandidateToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "RTCIceCandidate.toJSON".
func (this RTCIceCandidate) FuncToJSON() (fn js.Func[func() RTCIceCandidateInit]) {
	bindings.FuncRTCIceCandidateToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "RTCIceCandidate.toJSON".
func (this RTCIceCandidate) ToJSON() (ret RTCIceCandidateInit) {
	bindings.CallRTCIceCandidateToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "RTCIceCandidate.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCIceCandidate) TryToJSON() (ret RTCIceCandidateInit, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCIceCandidateToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
func (p *RTCIceCandidatePair) UpdateFrom(ref js.Ref) {
	bindings.RTCIceCandidatePairJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCIceCandidatePair) Update(ref js.Ref) {
	bindings.RTCIceCandidatePairJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCIceCandidatePair) FreeMembers(recursive bool) {
	js.Free(
		p.Local.Ref(),
		p.Remote.Ref(),
	)
	p.Local = p.Local.FromRef(js.Undefined)
	p.Remote = p.Remote.FromRef(js.Undefined)
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
func (p *RTCIceParameters) UpdateFrom(ref js.Ref) {
	bindings.RTCIceParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCIceParameters) Update(ref js.Ref) {
	bindings.RTCIceParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCIceParameters) FreeMembers(recursive bool) {
	js.Free(
		p.UsernameFragment.Ref(),
		p.Password.Ref(),
	)
	p.UsernameFragment = p.UsernameFragment.FromRef(js.Undefined)
	p.Password = p.Password.FromRef(js.Undefined)
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
func (p *RTCIceGatherOptions) UpdateFrom(ref js.Ref) {
	bindings.RTCIceGatherOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCIceGatherOptions) Update(ref js.Ref) {
	bindings.RTCIceGatherOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCIceGatherOptions) FreeMembers(recursive bool) {
	js.Free(
		p.IceServers.Ref(),
	)
	p.IceServers = p.IceServers.FromRef(js.Undefined)
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
