// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type RTCIceGathererState uint32

const (
	_ RTCIceGathererState = iota

	RTCIceGathererState_NEW
	RTCIceGathererState_GATHERING
	RTCIceGathererState_COMPLETE
)

func (RTCIceGathererState) FromRef(str js.Ref) RTCIceGathererState {
	return RTCIceGathererState(bindings.ConstOfRTCIceGathererState(str))
}

func (x RTCIceGathererState) String() (string, bool) {
	switch x {
	case RTCIceGathererState_NEW:
		return "new", true
	case RTCIceGathererState_GATHERING:
		return "gathering", true
	case RTCIceGathererState_COMPLETE:
		return "complete", true
	default:
		return "", false
	}
}

type RTCIceTransport struct {
	EventTarget
}

func (this RTCIceTransport) Once() RTCIceTransport {
	this.ref.Once()
	return this
}

func (this RTCIceTransport) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this RTCIceTransport) FromRef(ref js.Ref) RTCIceTransport {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this RTCIceTransport) Free() {
	this.ref.Free()
}

// Role returns the value of property "RTCIceTransport.role".
//
// It returns ok=false if there is no such property.
func (this RTCIceTransport) Role() (ret RTCIceRole, ok bool) {
	ok = js.True == bindings.GetRTCIceTransportRole(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Component returns the value of property "RTCIceTransport.component".
//
// It returns ok=false if there is no such property.
func (this RTCIceTransport) Component() (ret RTCIceComponent, ok bool) {
	ok = js.True == bindings.GetRTCIceTransportComponent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// State returns the value of property "RTCIceTransport.state".
//
// It returns ok=false if there is no such property.
func (this RTCIceTransport) State() (ret RTCIceTransportState, ok bool) {
	ok = js.True == bindings.GetRTCIceTransportState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// GatheringState returns the value of property "RTCIceTransport.gatheringState".
//
// It returns ok=false if there is no such property.
func (this RTCIceTransport) GatheringState() (ret RTCIceGathererState, ok bool) {
	ok = js.True == bindings.GetRTCIceTransportGatheringState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetLocalCandidates returns true if the method "RTCIceTransport.getLocalCandidates" exists.
func (this RTCIceTransport) HasFuncGetLocalCandidates() bool {
	return js.True == bindings.HasFuncRTCIceTransportGetLocalCandidates(
		this.ref,
	)
}

// FuncGetLocalCandidates returns the method "RTCIceTransport.getLocalCandidates".
func (this RTCIceTransport) FuncGetLocalCandidates() (fn js.Func[func() js.Array[RTCIceCandidate]]) {
	bindings.FuncRTCIceTransportGetLocalCandidates(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetLocalCandidates calls the method "RTCIceTransport.getLocalCandidates".
func (this RTCIceTransport) GetLocalCandidates() (ret js.Array[RTCIceCandidate]) {
	bindings.CallRTCIceTransportGetLocalCandidates(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetLocalCandidates calls the method "RTCIceTransport.getLocalCandidates"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCIceTransport) TryGetLocalCandidates() (ret js.Array[RTCIceCandidate], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCIceTransportGetLocalCandidates(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetRemoteCandidates returns true if the method "RTCIceTransport.getRemoteCandidates" exists.
func (this RTCIceTransport) HasFuncGetRemoteCandidates() bool {
	return js.True == bindings.HasFuncRTCIceTransportGetRemoteCandidates(
		this.ref,
	)
}

// FuncGetRemoteCandidates returns the method "RTCIceTransport.getRemoteCandidates".
func (this RTCIceTransport) FuncGetRemoteCandidates() (fn js.Func[func() js.Array[RTCIceCandidate]]) {
	bindings.FuncRTCIceTransportGetRemoteCandidates(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetRemoteCandidates calls the method "RTCIceTransport.getRemoteCandidates".
func (this RTCIceTransport) GetRemoteCandidates() (ret js.Array[RTCIceCandidate]) {
	bindings.CallRTCIceTransportGetRemoteCandidates(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetRemoteCandidates calls the method "RTCIceTransport.getRemoteCandidates"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCIceTransport) TryGetRemoteCandidates() (ret js.Array[RTCIceCandidate], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCIceTransportGetRemoteCandidates(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetSelectedCandidatePair returns true if the method "RTCIceTransport.getSelectedCandidatePair" exists.
func (this RTCIceTransport) HasFuncGetSelectedCandidatePair() bool {
	return js.True == bindings.HasFuncRTCIceTransportGetSelectedCandidatePair(
		this.ref,
	)
}

// FuncGetSelectedCandidatePair returns the method "RTCIceTransport.getSelectedCandidatePair".
func (this RTCIceTransport) FuncGetSelectedCandidatePair() (fn js.Func[func() RTCIceCandidatePair]) {
	bindings.FuncRTCIceTransportGetSelectedCandidatePair(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSelectedCandidatePair calls the method "RTCIceTransport.getSelectedCandidatePair".
func (this RTCIceTransport) GetSelectedCandidatePair() (ret RTCIceCandidatePair) {
	bindings.CallRTCIceTransportGetSelectedCandidatePair(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetSelectedCandidatePair calls the method "RTCIceTransport.getSelectedCandidatePair"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCIceTransport) TryGetSelectedCandidatePair() (ret RTCIceCandidatePair, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCIceTransportGetSelectedCandidatePair(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetLocalParameters returns true if the method "RTCIceTransport.getLocalParameters" exists.
func (this RTCIceTransport) HasFuncGetLocalParameters() bool {
	return js.True == bindings.HasFuncRTCIceTransportGetLocalParameters(
		this.ref,
	)
}

// FuncGetLocalParameters returns the method "RTCIceTransport.getLocalParameters".
func (this RTCIceTransport) FuncGetLocalParameters() (fn js.Func[func() RTCIceParameters]) {
	bindings.FuncRTCIceTransportGetLocalParameters(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetLocalParameters calls the method "RTCIceTransport.getLocalParameters".
func (this RTCIceTransport) GetLocalParameters() (ret RTCIceParameters) {
	bindings.CallRTCIceTransportGetLocalParameters(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetLocalParameters calls the method "RTCIceTransport.getLocalParameters"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCIceTransport) TryGetLocalParameters() (ret RTCIceParameters, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCIceTransportGetLocalParameters(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetRemoteParameters returns true if the method "RTCIceTransport.getRemoteParameters" exists.
func (this RTCIceTransport) HasFuncGetRemoteParameters() bool {
	return js.True == bindings.HasFuncRTCIceTransportGetRemoteParameters(
		this.ref,
	)
}

// FuncGetRemoteParameters returns the method "RTCIceTransport.getRemoteParameters".
func (this RTCIceTransport) FuncGetRemoteParameters() (fn js.Func[func() RTCIceParameters]) {
	bindings.FuncRTCIceTransportGetRemoteParameters(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetRemoteParameters calls the method "RTCIceTransport.getRemoteParameters".
func (this RTCIceTransport) GetRemoteParameters() (ret RTCIceParameters) {
	bindings.CallRTCIceTransportGetRemoteParameters(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetRemoteParameters calls the method "RTCIceTransport.getRemoteParameters"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCIceTransport) TryGetRemoteParameters() (ret RTCIceParameters, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCIceTransportGetRemoteParameters(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGather returns true if the method "RTCIceTransport.gather" exists.
func (this RTCIceTransport) HasFuncGather() bool {
	return js.True == bindings.HasFuncRTCIceTransportGather(
		this.ref,
	)
}

// FuncGather returns the method "RTCIceTransport.gather".
func (this RTCIceTransport) FuncGather() (fn js.Func[func(options RTCIceGatherOptions)]) {
	bindings.FuncRTCIceTransportGather(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Gather calls the method "RTCIceTransport.gather".
func (this RTCIceTransport) Gather(options RTCIceGatherOptions) (ret js.Void) {
	bindings.CallRTCIceTransportGather(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGather calls the method "RTCIceTransport.gather"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCIceTransport) TryGather(options RTCIceGatherOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCIceTransportGather(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGather1 returns true if the method "RTCIceTransport.gather" exists.
func (this RTCIceTransport) HasFuncGather1() bool {
	return js.True == bindings.HasFuncRTCIceTransportGather1(
		this.ref,
	)
}

// FuncGather1 returns the method "RTCIceTransport.gather".
func (this RTCIceTransport) FuncGather1() (fn js.Func[func()]) {
	bindings.FuncRTCIceTransportGather1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Gather1 calls the method "RTCIceTransport.gather".
func (this RTCIceTransport) Gather1() (ret js.Void) {
	bindings.CallRTCIceTransportGather1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGather1 calls the method "RTCIceTransport.gather"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCIceTransport) TryGather1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCIceTransportGather1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStart returns true if the method "RTCIceTransport.start" exists.
func (this RTCIceTransport) HasFuncStart() bool {
	return js.True == bindings.HasFuncRTCIceTransportStart(
		this.ref,
	)
}

// FuncStart returns the method "RTCIceTransport.start".
func (this RTCIceTransport) FuncStart() (fn js.Func[func(remoteParameters RTCIceParameters, role RTCIceRole)]) {
	bindings.FuncRTCIceTransportStart(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start calls the method "RTCIceTransport.start".
func (this RTCIceTransport) Start(remoteParameters RTCIceParameters, role RTCIceRole) (ret js.Void) {
	bindings.CallRTCIceTransportStart(
		this.ref, js.Pointer(&ret),
		js.Pointer(&remoteParameters),
		uint32(role),
	)

	return
}

// TryStart calls the method "RTCIceTransport.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCIceTransport) TryStart(remoteParameters RTCIceParameters, role RTCIceRole) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCIceTransportStart(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&remoteParameters),
		uint32(role),
	)

	return
}

// HasFuncStart1 returns true if the method "RTCIceTransport.start" exists.
func (this RTCIceTransport) HasFuncStart1() bool {
	return js.True == bindings.HasFuncRTCIceTransportStart1(
		this.ref,
	)
}

// FuncStart1 returns the method "RTCIceTransport.start".
func (this RTCIceTransport) FuncStart1() (fn js.Func[func(remoteParameters RTCIceParameters)]) {
	bindings.FuncRTCIceTransportStart1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start1 calls the method "RTCIceTransport.start".
func (this RTCIceTransport) Start1(remoteParameters RTCIceParameters) (ret js.Void) {
	bindings.CallRTCIceTransportStart1(
		this.ref, js.Pointer(&ret),
		js.Pointer(&remoteParameters),
	)

	return
}

// TryStart1 calls the method "RTCIceTransport.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCIceTransport) TryStart1(remoteParameters RTCIceParameters) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCIceTransportStart1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&remoteParameters),
	)

	return
}

// HasFuncStart2 returns true if the method "RTCIceTransport.start" exists.
func (this RTCIceTransport) HasFuncStart2() bool {
	return js.True == bindings.HasFuncRTCIceTransportStart2(
		this.ref,
	)
}

// FuncStart2 returns the method "RTCIceTransport.start".
func (this RTCIceTransport) FuncStart2() (fn js.Func[func()]) {
	bindings.FuncRTCIceTransportStart2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start2 calls the method "RTCIceTransport.start".
func (this RTCIceTransport) Start2() (ret js.Void) {
	bindings.CallRTCIceTransportStart2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStart2 calls the method "RTCIceTransport.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCIceTransport) TryStart2() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCIceTransportStart2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStop returns true if the method "RTCIceTransport.stop" exists.
func (this RTCIceTransport) HasFuncStop() bool {
	return js.True == bindings.HasFuncRTCIceTransportStop(
		this.ref,
	)
}

// FuncStop returns the method "RTCIceTransport.stop".
func (this RTCIceTransport) FuncStop() (fn js.Func[func()]) {
	bindings.FuncRTCIceTransportStop(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Stop calls the method "RTCIceTransport.stop".
func (this RTCIceTransport) Stop() (ret js.Void) {
	bindings.CallRTCIceTransportStop(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStop calls the method "RTCIceTransport.stop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCIceTransport) TryStop() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCIceTransportStop(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAddRemoteCandidate returns true if the method "RTCIceTransport.addRemoteCandidate" exists.
func (this RTCIceTransport) HasFuncAddRemoteCandidate() bool {
	return js.True == bindings.HasFuncRTCIceTransportAddRemoteCandidate(
		this.ref,
	)
}

// FuncAddRemoteCandidate returns the method "RTCIceTransport.addRemoteCandidate".
func (this RTCIceTransport) FuncAddRemoteCandidate() (fn js.Func[func(remoteCandidate RTCIceCandidateInit)]) {
	bindings.FuncRTCIceTransportAddRemoteCandidate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddRemoteCandidate calls the method "RTCIceTransport.addRemoteCandidate".
func (this RTCIceTransport) AddRemoteCandidate(remoteCandidate RTCIceCandidateInit) (ret js.Void) {
	bindings.CallRTCIceTransportAddRemoteCandidate(
		this.ref, js.Pointer(&ret),
		js.Pointer(&remoteCandidate),
	)

	return
}

// TryAddRemoteCandidate calls the method "RTCIceTransport.addRemoteCandidate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCIceTransport) TryAddRemoteCandidate(remoteCandidate RTCIceCandidateInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCIceTransportAddRemoteCandidate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&remoteCandidate),
	)

	return
}

// HasFuncAddRemoteCandidate1 returns true if the method "RTCIceTransport.addRemoteCandidate" exists.
func (this RTCIceTransport) HasFuncAddRemoteCandidate1() bool {
	return js.True == bindings.HasFuncRTCIceTransportAddRemoteCandidate1(
		this.ref,
	)
}

// FuncAddRemoteCandidate1 returns the method "RTCIceTransport.addRemoteCandidate".
func (this RTCIceTransport) FuncAddRemoteCandidate1() (fn js.Func[func()]) {
	bindings.FuncRTCIceTransportAddRemoteCandidate1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddRemoteCandidate1 calls the method "RTCIceTransport.addRemoteCandidate".
func (this RTCIceTransport) AddRemoteCandidate1() (ret js.Void) {
	bindings.CallRTCIceTransportAddRemoteCandidate1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAddRemoteCandidate1 calls the method "RTCIceTransport.addRemoteCandidate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCIceTransport) TryAddRemoteCandidate1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCIceTransportAddRemoteCandidate1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type RTCDtlsTransportState uint32

const (
	_ RTCDtlsTransportState = iota

	RTCDtlsTransportState_NEW
	RTCDtlsTransportState_CONNECTING
	RTCDtlsTransportState_CONNECTED
	RTCDtlsTransportState_CLOSED
	RTCDtlsTransportState_FAILED
)

func (RTCDtlsTransportState) FromRef(str js.Ref) RTCDtlsTransportState {
	return RTCDtlsTransportState(bindings.ConstOfRTCDtlsTransportState(str))
}

func (x RTCDtlsTransportState) String() (string, bool) {
	switch x {
	case RTCDtlsTransportState_NEW:
		return "new", true
	case RTCDtlsTransportState_CONNECTING:
		return "connecting", true
	case RTCDtlsTransportState_CONNECTED:
		return "connected", true
	case RTCDtlsTransportState_CLOSED:
		return "closed", true
	case RTCDtlsTransportState_FAILED:
		return "failed", true
	default:
		return "", false
	}
}

type RTCDtlsTransport struct {
	EventTarget
}

func (this RTCDtlsTransport) Once() RTCDtlsTransport {
	this.ref.Once()
	return this
}

func (this RTCDtlsTransport) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this RTCDtlsTransport) FromRef(ref js.Ref) RTCDtlsTransport {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this RTCDtlsTransport) Free() {
	this.ref.Free()
}

// IceTransport returns the value of property "RTCDtlsTransport.iceTransport".
//
// It returns ok=false if there is no such property.
func (this RTCDtlsTransport) IceTransport() (ret RTCIceTransport, ok bool) {
	ok = js.True == bindings.GetRTCDtlsTransportIceTransport(
		this.ref, js.Pointer(&ret),
	)
	return
}

// State returns the value of property "RTCDtlsTransport.state".
//
// It returns ok=false if there is no such property.
func (this RTCDtlsTransport) State() (ret RTCDtlsTransportState, ok bool) {
	ok = js.True == bindings.GetRTCDtlsTransportState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetRemoteCertificates returns true if the method "RTCDtlsTransport.getRemoteCertificates" exists.
func (this RTCDtlsTransport) HasFuncGetRemoteCertificates() bool {
	return js.True == bindings.HasFuncRTCDtlsTransportGetRemoteCertificates(
		this.ref,
	)
}

// FuncGetRemoteCertificates returns the method "RTCDtlsTransport.getRemoteCertificates".
func (this RTCDtlsTransport) FuncGetRemoteCertificates() (fn js.Func[func() js.Array[js.ArrayBuffer]]) {
	bindings.FuncRTCDtlsTransportGetRemoteCertificates(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetRemoteCertificates calls the method "RTCDtlsTransport.getRemoteCertificates".
func (this RTCDtlsTransport) GetRemoteCertificates() (ret js.Array[js.ArrayBuffer]) {
	bindings.CallRTCDtlsTransportGetRemoteCertificates(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetRemoteCertificates calls the method "RTCDtlsTransport.getRemoteCertificates"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCDtlsTransport) TryGetRemoteCertificates() (ret js.Array[js.ArrayBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCDtlsTransportGetRemoteCertificates(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type RTCEncodedAudioFrameMetadata struct {
	// SynchronizationSource is "RTCEncodedAudioFrameMetadata.synchronizationSource"
	//
	// Optional
	//
	// NOTE: FFI_USE_SynchronizationSource MUST be set to true to make this field effective.
	SynchronizationSource uint32
	// PayloadType is "RTCEncodedAudioFrameMetadata.payloadType"
	//
	// Optional
	//
	// NOTE: FFI_USE_PayloadType MUST be set to true to make this field effective.
	PayloadType uint8
	// ContributingSources is "RTCEncodedAudioFrameMetadata.contributingSources"
	//
	// Optional
	ContributingSources js.Array[uint32]
	// SequenceNumber is "RTCEncodedAudioFrameMetadata.sequenceNumber"
	//
	// Optional
	//
	// NOTE: FFI_USE_SequenceNumber MUST be set to true to make this field effective.
	SequenceNumber int16

	FFI_USE_SynchronizationSource bool // for SynchronizationSource.
	FFI_USE_PayloadType           bool // for PayloadType.
	FFI_USE_SequenceNumber        bool // for SequenceNumber.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCEncodedAudioFrameMetadata with all fields set.
func (p RTCEncodedAudioFrameMetadata) FromRef(ref js.Ref) RTCEncodedAudioFrameMetadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCEncodedAudioFrameMetadata in the application heap.
func (p RTCEncodedAudioFrameMetadata) New() js.Ref {
	return bindings.RTCEncodedAudioFrameMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCEncodedAudioFrameMetadata) UpdateFrom(ref js.Ref) {
	bindings.RTCEncodedAudioFrameMetadataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCEncodedAudioFrameMetadata) Update(ref js.Ref) {
	bindings.RTCEncodedAudioFrameMetadataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCEncodedAudioFrameMetadata) FreeMembers(recursive bool) {
	js.Free(
		p.ContributingSources.Ref(),
	)
	p.ContributingSources = p.ContributingSources.FromRef(js.Undefined)
}

type RTCEncodedAudioFrame struct {
	ref js.Ref
}

func (this RTCEncodedAudioFrame) Once() RTCEncodedAudioFrame {
	this.ref.Once()
	return this
}

func (this RTCEncodedAudioFrame) Ref() js.Ref {
	return this.ref
}

func (this RTCEncodedAudioFrame) FromRef(ref js.Ref) RTCEncodedAudioFrame {
	this.ref = ref
	return this
}

func (this RTCEncodedAudioFrame) Free() {
	this.ref.Free()
}

// Timestamp returns the value of property "RTCEncodedAudioFrame.timestamp".
//
// It returns ok=false if there is no such property.
func (this RTCEncodedAudioFrame) Timestamp() (ret uint32, ok bool) {
	ok = js.True == bindings.GetRTCEncodedAudioFrameTimestamp(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Data returns the value of property "RTCEncodedAudioFrame.data".
//
// It returns ok=false if there is no such property.
func (this RTCEncodedAudioFrame) Data() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetRTCEncodedAudioFrameData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetData sets the value of property "RTCEncodedAudioFrame.data" to val.
//
// It returns false if the property cannot be set.
func (this RTCEncodedAudioFrame) SetData(val js.ArrayBuffer) bool {
	return js.True == bindings.SetRTCEncodedAudioFrameData(
		this.ref,
		val.Ref(),
	)
}

// HasFuncGetMetadata returns true if the method "RTCEncodedAudioFrame.getMetadata" exists.
func (this RTCEncodedAudioFrame) HasFuncGetMetadata() bool {
	return js.True == bindings.HasFuncRTCEncodedAudioFrameGetMetadata(
		this.ref,
	)
}

// FuncGetMetadata returns the method "RTCEncodedAudioFrame.getMetadata".
func (this RTCEncodedAudioFrame) FuncGetMetadata() (fn js.Func[func() RTCEncodedAudioFrameMetadata]) {
	bindings.FuncRTCEncodedAudioFrameGetMetadata(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetMetadata calls the method "RTCEncodedAudioFrame.getMetadata".
func (this RTCEncodedAudioFrame) GetMetadata() (ret RTCEncodedAudioFrameMetadata) {
	bindings.CallRTCEncodedAudioFrameGetMetadata(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetMetadata calls the method "RTCEncodedAudioFrame.getMetadata"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCEncodedAudioFrame) TryGetMetadata() (ret RTCEncodedAudioFrameMetadata, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCEncodedAudioFrameGetMetadata(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type RTCEncodedVideoFrameMetadata struct {
	// FrameId is "RTCEncodedVideoFrameMetadata.frameId"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameId MUST be set to true to make this field effective.
	FrameId uint64
	// Dependencies is "RTCEncodedVideoFrameMetadata.dependencies"
	//
	// Optional
	Dependencies js.Array[uint64]
	// Width is "RTCEncodedVideoFrameMetadata.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width uint16
	// Height is "RTCEncodedVideoFrameMetadata.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height uint16
	// SpatialIndex is "RTCEncodedVideoFrameMetadata.spatialIndex"
	//
	// Optional
	//
	// NOTE: FFI_USE_SpatialIndex MUST be set to true to make this field effective.
	SpatialIndex uint32
	// TemporalIndex is "RTCEncodedVideoFrameMetadata.temporalIndex"
	//
	// Optional
	//
	// NOTE: FFI_USE_TemporalIndex MUST be set to true to make this field effective.
	TemporalIndex uint32
	// SynchronizationSource is "RTCEncodedVideoFrameMetadata.synchronizationSource"
	//
	// Optional
	//
	// NOTE: FFI_USE_SynchronizationSource MUST be set to true to make this field effective.
	SynchronizationSource uint32
	// PayloadType is "RTCEncodedVideoFrameMetadata.payloadType"
	//
	// Optional
	//
	// NOTE: FFI_USE_PayloadType MUST be set to true to make this field effective.
	PayloadType uint8
	// ContributingSources is "RTCEncodedVideoFrameMetadata.contributingSources"
	//
	// Optional
	ContributingSources js.Array[uint32]
	// Timestamp is "RTCEncodedVideoFrameMetadata.timestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timestamp MUST be set to true to make this field effective.
	Timestamp int64

	FFI_USE_FrameId               bool // for FrameId.
	FFI_USE_Width                 bool // for Width.
	FFI_USE_Height                bool // for Height.
	FFI_USE_SpatialIndex          bool // for SpatialIndex.
	FFI_USE_TemporalIndex         bool // for TemporalIndex.
	FFI_USE_SynchronizationSource bool // for SynchronizationSource.
	FFI_USE_PayloadType           bool // for PayloadType.
	FFI_USE_Timestamp             bool // for Timestamp.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCEncodedVideoFrameMetadata with all fields set.
func (p RTCEncodedVideoFrameMetadata) FromRef(ref js.Ref) RTCEncodedVideoFrameMetadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCEncodedVideoFrameMetadata in the application heap.
func (p RTCEncodedVideoFrameMetadata) New() js.Ref {
	return bindings.RTCEncodedVideoFrameMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCEncodedVideoFrameMetadata) UpdateFrom(ref js.Ref) {
	bindings.RTCEncodedVideoFrameMetadataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCEncodedVideoFrameMetadata) Update(ref js.Ref) {
	bindings.RTCEncodedVideoFrameMetadataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCEncodedVideoFrameMetadata) FreeMembers(recursive bool) {
	js.Free(
		p.Dependencies.Ref(),
		p.ContributingSources.Ref(),
	)
	p.Dependencies = p.Dependencies.FromRef(js.Undefined)
	p.ContributingSources = p.ContributingSources.FromRef(js.Undefined)
}

type RTCEncodedVideoFrameType uint32

const (
	_ RTCEncodedVideoFrameType = iota

	RTCEncodedVideoFrameType_EMPTY
	RTCEncodedVideoFrameType_KEY
	RTCEncodedVideoFrameType_DELTA
)

func (RTCEncodedVideoFrameType) FromRef(str js.Ref) RTCEncodedVideoFrameType {
	return RTCEncodedVideoFrameType(bindings.ConstOfRTCEncodedVideoFrameType(str))
}

func (x RTCEncodedVideoFrameType) String() (string, bool) {
	switch x {
	case RTCEncodedVideoFrameType_EMPTY:
		return "empty", true
	case RTCEncodedVideoFrameType_KEY:
		return "key", true
	case RTCEncodedVideoFrameType_DELTA:
		return "delta", true
	default:
		return "", false
	}
}

type RTCEncodedVideoFrame struct {
	ref js.Ref
}

func (this RTCEncodedVideoFrame) Once() RTCEncodedVideoFrame {
	this.ref.Once()
	return this
}

func (this RTCEncodedVideoFrame) Ref() js.Ref {
	return this.ref
}

func (this RTCEncodedVideoFrame) FromRef(ref js.Ref) RTCEncodedVideoFrame {
	this.ref = ref
	return this
}

func (this RTCEncodedVideoFrame) Free() {
	this.ref.Free()
}

// Type returns the value of property "RTCEncodedVideoFrame.type".
//
// It returns ok=false if there is no such property.
func (this RTCEncodedVideoFrame) Type() (ret RTCEncodedVideoFrameType, ok bool) {
	ok = js.True == bindings.GetRTCEncodedVideoFrameType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Timestamp returns the value of property "RTCEncodedVideoFrame.timestamp".
//
// It returns ok=false if there is no such property.
func (this RTCEncodedVideoFrame) Timestamp() (ret uint32, ok bool) {
	ok = js.True == bindings.GetRTCEncodedVideoFrameTimestamp(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Data returns the value of property "RTCEncodedVideoFrame.data".
//
// It returns ok=false if there is no such property.
func (this RTCEncodedVideoFrame) Data() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetRTCEncodedVideoFrameData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetData sets the value of property "RTCEncodedVideoFrame.data" to val.
//
// It returns false if the property cannot be set.
func (this RTCEncodedVideoFrame) SetData(val js.ArrayBuffer) bool {
	return js.True == bindings.SetRTCEncodedVideoFrameData(
		this.ref,
		val.Ref(),
	)
}

// HasFuncGetMetadata returns true if the method "RTCEncodedVideoFrame.getMetadata" exists.
func (this RTCEncodedVideoFrame) HasFuncGetMetadata() bool {
	return js.True == bindings.HasFuncRTCEncodedVideoFrameGetMetadata(
		this.ref,
	)
}

// FuncGetMetadata returns the method "RTCEncodedVideoFrame.getMetadata".
func (this RTCEncodedVideoFrame) FuncGetMetadata() (fn js.Func[func() RTCEncodedVideoFrameMetadata]) {
	bindings.FuncRTCEncodedVideoFrameGetMetadata(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetMetadata calls the method "RTCEncodedVideoFrame.getMetadata".
func (this RTCEncodedVideoFrame) GetMetadata() (ret RTCEncodedVideoFrameMetadata) {
	bindings.CallRTCEncodedVideoFrameGetMetadata(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetMetadata calls the method "RTCEncodedVideoFrame.getMetadata"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCEncodedVideoFrame) TryGetMetadata() (ret RTCEncodedVideoFrameMetadata, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCEncodedVideoFrameGetMetadata(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type RTCErrorDetailType uint32

const (
	_ RTCErrorDetailType = iota

	RTCErrorDetailType_DATA_CHANNEL_FAILURE
	RTCErrorDetailType_DTLS_FAILURE
	RTCErrorDetailType_FINGERPRINT_FAILURE
	RTCErrorDetailType_SCTP_FAILURE
	RTCErrorDetailType_SDP_SYNTAX_ERROR
	RTCErrorDetailType_HARDWARE_ENCODER_NOT_AVAILABLE
	RTCErrorDetailType_HARDWARE_ENCODER_ERROR
)

func (RTCErrorDetailType) FromRef(str js.Ref) RTCErrorDetailType {
	return RTCErrorDetailType(bindings.ConstOfRTCErrorDetailType(str))
}

func (x RTCErrorDetailType) String() (string, bool) {
	switch x {
	case RTCErrorDetailType_DATA_CHANNEL_FAILURE:
		return "data-channel-failure", true
	case RTCErrorDetailType_DTLS_FAILURE:
		return "dtls-failure", true
	case RTCErrorDetailType_FINGERPRINT_FAILURE:
		return "fingerprint-failure", true
	case RTCErrorDetailType_SCTP_FAILURE:
		return "sctp-failure", true
	case RTCErrorDetailType_SDP_SYNTAX_ERROR:
		return "sdp-syntax-error", true
	case RTCErrorDetailType_HARDWARE_ENCODER_NOT_AVAILABLE:
		return "hardware-encoder-not-available", true
	case RTCErrorDetailType_HARDWARE_ENCODER_ERROR:
		return "hardware-encoder-error", true
	default:
		return "", false
	}
}

type RTCErrorInit struct {
	// ErrorDetail is "RTCErrorInit.errorDetail"
	//
	// Required
	ErrorDetail RTCErrorDetailType
	// SdpLineNumber is "RTCErrorInit.sdpLineNumber"
	//
	// Optional
	//
	// NOTE: FFI_USE_SdpLineNumber MUST be set to true to make this field effective.
	SdpLineNumber int32
	// SctpCauseCode is "RTCErrorInit.sctpCauseCode"
	//
	// Optional
	//
	// NOTE: FFI_USE_SctpCauseCode MUST be set to true to make this field effective.
	SctpCauseCode int32
	// ReceivedAlert is "RTCErrorInit.receivedAlert"
	//
	// Optional
	//
	// NOTE: FFI_USE_ReceivedAlert MUST be set to true to make this field effective.
	ReceivedAlert uint32
	// SentAlert is "RTCErrorInit.sentAlert"
	//
	// Optional
	//
	// NOTE: FFI_USE_SentAlert MUST be set to true to make this field effective.
	SentAlert uint32
	// HttpRequestStatusCode is "RTCErrorInit.httpRequestStatusCode"
	//
	// Optional
	//
	// NOTE: FFI_USE_HttpRequestStatusCode MUST be set to true to make this field effective.
	HttpRequestStatusCode int32

	FFI_USE_SdpLineNumber         bool // for SdpLineNumber.
	FFI_USE_SctpCauseCode         bool // for SctpCauseCode.
	FFI_USE_ReceivedAlert         bool // for ReceivedAlert.
	FFI_USE_SentAlert             bool // for SentAlert.
	FFI_USE_HttpRequestStatusCode bool // for HttpRequestStatusCode.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCErrorInit with all fields set.
func (p RTCErrorInit) FromRef(ref js.Ref) RTCErrorInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCErrorInit in the application heap.
func (p RTCErrorInit) New() js.Ref {
	return bindings.RTCErrorInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCErrorInit) UpdateFrom(ref js.Ref) {
	bindings.RTCErrorInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCErrorInit) Update(ref js.Ref) {
	bindings.RTCErrorInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCErrorInit) FreeMembers(recursive bool) {
}

func NewRTCError(init RTCErrorInit, message js.String) (ret RTCError) {
	ret.ref = bindings.NewRTCErrorByRTCError(
		js.Pointer(&init),
		message.Ref())
	return
}

func NewRTCErrorByRTCError1(init RTCErrorInit) (ret RTCError) {
	ret.ref = bindings.NewRTCErrorByRTCError1(
		js.Pointer(&init))
	return
}

type RTCError struct {
	DOMException
}

func (this RTCError) Once() RTCError {
	this.ref.Once()
	return this
}

func (this RTCError) Ref() js.Ref {
	return this.DOMException.Ref()
}

func (this RTCError) FromRef(ref js.Ref) RTCError {
	this.DOMException = this.DOMException.FromRef(ref)
	return this
}

func (this RTCError) Free() {
	this.ref.Free()
}

// ErrorDetail returns the value of property "RTCError.errorDetail".
//
// It returns ok=false if there is no such property.
func (this RTCError) ErrorDetail() (ret RTCErrorDetailType, ok bool) {
	ok = js.True == bindings.GetRTCErrorErrorDetail(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SdpLineNumber returns the value of property "RTCError.sdpLineNumber".
//
// It returns ok=false if there is no such property.
func (this RTCError) SdpLineNumber() (ret int32, ok bool) {
	ok = js.True == bindings.GetRTCErrorSdpLineNumber(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SctpCauseCode returns the value of property "RTCError.sctpCauseCode".
//
// It returns ok=false if there is no such property.
func (this RTCError) SctpCauseCode() (ret int32, ok bool) {
	ok = js.True == bindings.GetRTCErrorSctpCauseCode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ReceivedAlert returns the value of property "RTCError.receivedAlert".
//
// It returns ok=false if there is no such property.
func (this RTCError) ReceivedAlert() (ret uint32, ok bool) {
	ok = js.True == bindings.GetRTCErrorReceivedAlert(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SentAlert returns the value of property "RTCError.sentAlert".
//
// It returns ok=false if there is no such property.
func (this RTCError) SentAlert() (ret uint32, ok bool) {
	ok = js.True == bindings.GetRTCErrorSentAlert(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HttpRequestStatusCode returns the value of property "RTCError.httpRequestStatusCode".
//
// It returns ok=false if there is no such property.
func (this RTCError) HttpRequestStatusCode() (ret int32, ok bool) {
	ok = js.True == bindings.GetRTCErrorHttpRequestStatusCode(
		this.ref, js.Pointer(&ret),
	)
	return
}

type RTCErrorDetailTypeIdp uint32

const (
	_ RTCErrorDetailTypeIdp = iota

	RTCErrorDetailTypeIdp_IDP_BAD_SCRIPT_FAILURE
	RTCErrorDetailTypeIdp_IDP_EXECUTION_FAILURE
	RTCErrorDetailTypeIdp_IDP_LOAD_FAILURE
	RTCErrorDetailTypeIdp_IDP_NEED_LOGIN
	RTCErrorDetailTypeIdp_IDP_TIMEOUT
	RTCErrorDetailTypeIdp_IDP_TLS_FAILURE
	RTCErrorDetailTypeIdp_IDP_TOKEN_EXPIRED
	RTCErrorDetailTypeIdp_IDP_TOKEN_INVALID
)

func (RTCErrorDetailTypeIdp) FromRef(str js.Ref) RTCErrorDetailTypeIdp {
	return RTCErrorDetailTypeIdp(bindings.ConstOfRTCErrorDetailTypeIdp(str))
}

func (x RTCErrorDetailTypeIdp) String() (string, bool) {
	switch x {
	case RTCErrorDetailTypeIdp_IDP_BAD_SCRIPT_FAILURE:
		return "idp-bad-script-failure", true
	case RTCErrorDetailTypeIdp_IDP_EXECUTION_FAILURE:
		return "idp-execution-failure", true
	case RTCErrorDetailTypeIdp_IDP_LOAD_FAILURE:
		return "idp-load-failure", true
	case RTCErrorDetailTypeIdp_IDP_NEED_LOGIN:
		return "idp-need-login", true
	case RTCErrorDetailTypeIdp_IDP_TIMEOUT:
		return "idp-timeout", true
	case RTCErrorDetailTypeIdp_IDP_TLS_FAILURE:
		return "idp-tls-failure", true
	case RTCErrorDetailTypeIdp_IDP_TOKEN_EXPIRED:
		return "idp-token-expired", true
	case RTCErrorDetailTypeIdp_IDP_TOKEN_INVALID:
		return "idp-token-invalid", true
	default:
		return "", false
	}
}

type RTCErrorEventInit struct {
	// Error is "RTCErrorEventInit.error"
	//
	// Required
	Error RTCError
	// Bubbles is "RTCErrorEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "RTCErrorEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "RTCErrorEventInit.composed"
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

// FromRef calls UpdateFrom and returns a RTCErrorEventInit with all fields set.
func (p RTCErrorEventInit) FromRef(ref js.Ref) RTCErrorEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCErrorEventInit in the application heap.
func (p RTCErrorEventInit) New() js.Ref {
	return bindings.RTCErrorEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCErrorEventInit) UpdateFrom(ref js.Ref) {
	bindings.RTCErrorEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCErrorEventInit) Update(ref js.Ref) {
	bindings.RTCErrorEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCErrorEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Error.Ref(),
	)
	p.Error = p.Error.FromRef(js.Undefined)
}

func NewRTCErrorEvent(typ js.String, eventInitDict RTCErrorEventInit) (ret RTCErrorEvent) {
	ret.ref = bindings.NewRTCErrorEventByRTCErrorEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type RTCErrorEvent struct {
	Event
}

func (this RTCErrorEvent) Once() RTCErrorEvent {
	this.ref.Once()
	return this
}

func (this RTCErrorEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this RTCErrorEvent) FromRef(ref js.Ref) RTCErrorEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this RTCErrorEvent) Free() {
	this.ref.Free()
}

// Error returns the value of property "RTCErrorEvent.error".
//
// It returns ok=false if there is no such property.
func (this RTCErrorEvent) Error() (ret RTCError, ok bool) {
	ok = js.True == bindings.GetRTCErrorEventError(
		this.ref, js.Pointer(&ret),
	)
	return
}

type RTCStatsIceCandidatePairState uint32

const (
	_ RTCStatsIceCandidatePairState = iota

	RTCStatsIceCandidatePairState_FROZEN
	RTCStatsIceCandidatePairState_WAITING
	RTCStatsIceCandidatePairState_IN_PROGRESS
	RTCStatsIceCandidatePairState_FAILED
	RTCStatsIceCandidatePairState_SUCCEEDED
)

func (RTCStatsIceCandidatePairState) FromRef(str js.Ref) RTCStatsIceCandidatePairState {
	return RTCStatsIceCandidatePairState(bindings.ConstOfRTCStatsIceCandidatePairState(str))
}

func (x RTCStatsIceCandidatePairState) String() (string, bool) {
	switch x {
	case RTCStatsIceCandidatePairState_FROZEN:
		return "frozen", true
	case RTCStatsIceCandidatePairState_WAITING:
		return "waiting", true
	case RTCStatsIceCandidatePairState_IN_PROGRESS:
		return "in-progress", true
	case RTCStatsIceCandidatePairState_FAILED:
		return "failed", true
	case RTCStatsIceCandidatePairState_SUCCEEDED:
		return "succeeded", true
	default:
		return "", false
	}
}

type RTCIceCandidatePairStats struct {
	// TransportId is "RTCIceCandidatePairStats.transportId"
	//
	// Required
	TransportId js.String
	// LocalCandidateId is "RTCIceCandidatePairStats.localCandidateId"
	//
	// Required
	LocalCandidateId js.String
	// RemoteCandidateId is "RTCIceCandidatePairStats.remoteCandidateId"
	//
	// Required
	RemoteCandidateId js.String
	// State is "RTCIceCandidatePairStats.state"
	//
	// Required
	State RTCStatsIceCandidatePairState
	// Nominated is "RTCIceCandidatePairStats.nominated"
	//
	// Optional
	//
	// NOTE: FFI_USE_Nominated MUST be set to true to make this field effective.
	Nominated bool
	// PacketsSent is "RTCIceCandidatePairStats.packetsSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsSent MUST be set to true to make this field effective.
	PacketsSent uint64
	// PacketsReceived is "RTCIceCandidatePairStats.packetsReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsReceived MUST be set to true to make this field effective.
	PacketsReceived uint64
	// BytesSent is "RTCIceCandidatePairStats.bytesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesSent MUST be set to true to make this field effective.
	BytesSent uint64
	// BytesReceived is "RTCIceCandidatePairStats.bytesReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesReceived MUST be set to true to make this field effective.
	BytesReceived uint64
	// LastPacketSentTimestamp is "RTCIceCandidatePairStats.lastPacketSentTimestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_LastPacketSentTimestamp MUST be set to true to make this field effective.
	LastPacketSentTimestamp DOMHighResTimeStamp
	// LastPacketReceivedTimestamp is "RTCIceCandidatePairStats.lastPacketReceivedTimestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_LastPacketReceivedTimestamp MUST be set to true to make this field effective.
	LastPacketReceivedTimestamp DOMHighResTimeStamp
	// TotalRoundTripTime is "RTCIceCandidatePairStats.totalRoundTripTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalRoundTripTime MUST be set to true to make this field effective.
	TotalRoundTripTime float64
	// CurrentRoundTripTime is "RTCIceCandidatePairStats.currentRoundTripTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_CurrentRoundTripTime MUST be set to true to make this field effective.
	CurrentRoundTripTime float64
	// AvailableOutgoingBitrate is "RTCIceCandidatePairStats.availableOutgoingBitrate"
	//
	// Optional
	//
	// NOTE: FFI_USE_AvailableOutgoingBitrate MUST be set to true to make this field effective.
	AvailableOutgoingBitrate float64
	// AvailableIncomingBitrate is "RTCIceCandidatePairStats.availableIncomingBitrate"
	//
	// Optional
	//
	// NOTE: FFI_USE_AvailableIncomingBitrate MUST be set to true to make this field effective.
	AvailableIncomingBitrate float64
	// RequestsReceived is "RTCIceCandidatePairStats.requestsReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestsReceived MUST be set to true to make this field effective.
	RequestsReceived uint64
	// RequestsSent is "RTCIceCandidatePairStats.requestsSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestsSent MUST be set to true to make this field effective.
	RequestsSent uint64
	// ResponsesReceived is "RTCIceCandidatePairStats.responsesReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResponsesReceived MUST be set to true to make this field effective.
	ResponsesReceived uint64
	// ResponsesSent is "RTCIceCandidatePairStats.responsesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResponsesSent MUST be set to true to make this field effective.
	ResponsesSent uint64
	// ConsentRequestsSent is "RTCIceCandidatePairStats.consentRequestsSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_ConsentRequestsSent MUST be set to true to make this field effective.
	ConsentRequestsSent uint64
	// PacketsDiscardedOnSend is "RTCIceCandidatePairStats.packetsDiscardedOnSend"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsDiscardedOnSend MUST be set to true to make this field effective.
	PacketsDiscardedOnSend uint32
	// BytesDiscardedOnSend is "RTCIceCandidatePairStats.bytesDiscardedOnSend"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesDiscardedOnSend MUST be set to true to make this field effective.
	BytesDiscardedOnSend uint64
	// Timestamp is "RTCIceCandidatePairStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCIceCandidatePairStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCIceCandidatePairStats.id"
	//
	// Required
	Id js.String

	FFI_USE_Nominated                   bool // for Nominated.
	FFI_USE_PacketsSent                 bool // for PacketsSent.
	FFI_USE_PacketsReceived             bool // for PacketsReceived.
	FFI_USE_BytesSent                   bool // for BytesSent.
	FFI_USE_BytesReceived               bool // for BytesReceived.
	FFI_USE_LastPacketSentTimestamp     bool // for LastPacketSentTimestamp.
	FFI_USE_LastPacketReceivedTimestamp bool // for LastPacketReceivedTimestamp.
	FFI_USE_TotalRoundTripTime          bool // for TotalRoundTripTime.
	FFI_USE_CurrentRoundTripTime        bool // for CurrentRoundTripTime.
	FFI_USE_AvailableOutgoingBitrate    bool // for AvailableOutgoingBitrate.
	FFI_USE_AvailableIncomingBitrate    bool // for AvailableIncomingBitrate.
	FFI_USE_RequestsReceived            bool // for RequestsReceived.
	FFI_USE_RequestsSent                bool // for RequestsSent.
	FFI_USE_ResponsesReceived           bool // for ResponsesReceived.
	FFI_USE_ResponsesSent               bool // for ResponsesSent.
	FFI_USE_ConsentRequestsSent         bool // for ConsentRequestsSent.
	FFI_USE_PacketsDiscardedOnSend      bool // for PacketsDiscardedOnSend.
	FFI_USE_BytesDiscardedOnSend        bool // for BytesDiscardedOnSend.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCIceCandidatePairStats with all fields set.
func (p RTCIceCandidatePairStats) FromRef(ref js.Ref) RTCIceCandidatePairStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCIceCandidatePairStats in the application heap.
func (p RTCIceCandidatePairStats) New() js.Ref {
	return bindings.RTCIceCandidatePairStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCIceCandidatePairStats) UpdateFrom(ref js.Ref) {
	bindings.RTCIceCandidatePairStatsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCIceCandidatePairStats) Update(ref js.Ref) {
	bindings.RTCIceCandidatePairStatsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCIceCandidatePairStats) FreeMembers(recursive bool) {
	js.Free(
		p.TransportId.Ref(),
		p.LocalCandidateId.Ref(),
		p.RemoteCandidateId.Ref(),
		p.Id.Ref(),
	)
	p.TransportId = p.TransportId.FromRef(js.Undefined)
	p.LocalCandidateId = p.LocalCandidateId.FromRef(js.Undefined)
	p.RemoteCandidateId = p.RemoteCandidateId.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
}

type RTCIceCandidateStats struct {
	// TransportId is "RTCIceCandidateStats.transportId"
	//
	// Required
	TransportId js.String
	// Address is "RTCIceCandidateStats.address"
	//
	// Optional
	Address js.String
	// Port is "RTCIceCandidateStats.port"
	//
	// Optional
	//
	// NOTE: FFI_USE_Port MUST be set to true to make this field effective.
	Port int32
	// Protocol is "RTCIceCandidateStats.protocol"
	//
	// Optional
	Protocol js.String
	// CandidateType is "RTCIceCandidateStats.candidateType"
	//
	// Required
	CandidateType RTCIceCandidateType
	// Priority is "RTCIceCandidateStats.priority"
	//
	// Optional
	//
	// NOTE: FFI_USE_Priority MUST be set to true to make this field effective.
	Priority int32
	// Url is "RTCIceCandidateStats.url"
	//
	// Optional
	Url js.String
	// RelayProtocol is "RTCIceCandidateStats.relayProtocol"
	//
	// Optional
	RelayProtocol RTCIceServerTransportProtocol
	// Foundation is "RTCIceCandidateStats.foundation"
	//
	// Optional
	Foundation js.String
	// RelatedAddress is "RTCIceCandidateStats.relatedAddress"
	//
	// Optional
	RelatedAddress js.String
	// RelatedPort is "RTCIceCandidateStats.relatedPort"
	//
	// Optional
	//
	// NOTE: FFI_USE_RelatedPort MUST be set to true to make this field effective.
	RelatedPort int32
	// UsernameFragment is "RTCIceCandidateStats.usernameFragment"
	//
	// Optional
	UsernameFragment js.String
	// TcpType is "RTCIceCandidateStats.tcpType"
	//
	// Optional
	TcpType RTCIceTcpCandidateType
	// Timestamp is "RTCIceCandidateStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCIceCandidateStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCIceCandidateStats.id"
	//
	// Required
	Id js.String

	FFI_USE_Port        bool // for Port.
	FFI_USE_Priority    bool // for Priority.
	FFI_USE_RelatedPort bool // for RelatedPort.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCIceCandidateStats with all fields set.
func (p RTCIceCandidateStats) FromRef(ref js.Ref) RTCIceCandidateStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCIceCandidateStats in the application heap.
func (p RTCIceCandidateStats) New() js.Ref {
	return bindings.RTCIceCandidateStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCIceCandidateStats) UpdateFrom(ref js.Ref) {
	bindings.RTCIceCandidateStatsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCIceCandidateStats) Update(ref js.Ref) {
	bindings.RTCIceCandidateStatsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCIceCandidateStats) FreeMembers(recursive bool) {
	js.Free(
		p.TransportId.Ref(),
		p.Address.Ref(),
		p.Protocol.Ref(),
		p.Url.Ref(),
		p.Foundation.Ref(),
		p.RelatedAddress.Ref(),
		p.UsernameFragment.Ref(),
		p.Id.Ref(),
	)
	p.TransportId = p.TransportId.FromRef(js.Undefined)
	p.Address = p.Address.FromRef(js.Undefined)
	p.Protocol = p.Protocol.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Foundation = p.Foundation.FromRef(js.Undefined)
	p.RelatedAddress = p.RelatedAddress.FromRef(js.Undefined)
	p.UsernameFragment = p.UsernameFragment.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
}

type RTCIceConnectionState uint32

const (
	_ RTCIceConnectionState = iota

	RTCIceConnectionState_CLOSED
	RTCIceConnectionState_FAILED
	RTCIceConnectionState_DISCONNECTED
	RTCIceConnectionState_NEW
	RTCIceConnectionState_CHECKING
	RTCIceConnectionState_COMPLETED
	RTCIceConnectionState_CONNECTED
)

func (RTCIceConnectionState) FromRef(str js.Ref) RTCIceConnectionState {
	return RTCIceConnectionState(bindings.ConstOfRTCIceConnectionState(str))
}

func (x RTCIceConnectionState) String() (string, bool) {
	switch x {
	case RTCIceConnectionState_CLOSED:
		return "closed", true
	case RTCIceConnectionState_FAILED:
		return "failed", true
	case RTCIceConnectionState_DISCONNECTED:
		return "disconnected", true
	case RTCIceConnectionState_NEW:
		return "new", true
	case RTCIceConnectionState_CHECKING:
		return "checking", true
	case RTCIceConnectionState_COMPLETED:
		return "completed", true
	case RTCIceConnectionState_CONNECTED:
		return "connected", true
	default:
		return "", false
	}
}

type RTCIceGatheringState uint32

const (
	_ RTCIceGatheringState = iota

	RTCIceGatheringState_NEW
	RTCIceGatheringState_GATHERING
	RTCIceGatheringState_COMPLETE
)

func (RTCIceGatheringState) FromRef(str js.Ref) RTCIceGatheringState {
	return RTCIceGatheringState(bindings.ConstOfRTCIceGatheringState(str))
}

func (x RTCIceGatheringState) String() (string, bool) {
	switch x {
	case RTCIceGatheringState_NEW:
		return "new", true
	case RTCIceGatheringState_GATHERING:
		return "gathering", true
	case RTCIceGatheringState_COMPLETE:
		return "complete", true
	default:
		return "", false
	}
}

func NewRTCIdentityAssertion(idp js.String, name js.String) (ret RTCIdentityAssertion) {
	ret.ref = bindings.NewRTCIdentityAssertionByRTCIdentityAssertion(
		idp.Ref(),
		name.Ref())
	return
}

type RTCIdentityAssertion struct {
	ref js.Ref
}

func (this RTCIdentityAssertion) Once() RTCIdentityAssertion {
	this.ref.Once()
	return this
}

func (this RTCIdentityAssertion) Ref() js.Ref {
	return this.ref
}

func (this RTCIdentityAssertion) FromRef(ref js.Ref) RTCIdentityAssertion {
	this.ref = ref
	return this
}

func (this RTCIdentityAssertion) Free() {
	this.ref.Free()
}

// Idp returns the value of property "RTCIdentityAssertion.idp".
//
// It returns ok=false if there is no such property.
func (this RTCIdentityAssertion) Idp() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCIdentityAssertionIdp(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIdp sets the value of property "RTCIdentityAssertion.idp" to val.
//
// It returns false if the property cannot be set.
func (this RTCIdentityAssertion) SetIdp(val js.String) bool {
	return js.True == bindings.SetRTCIdentityAssertionIdp(
		this.ref,
		val.Ref(),
	)
}

// Name returns the value of property "RTCIdentityAssertion.name".
//
// It returns ok=false if there is no such property.
func (this RTCIdentityAssertion) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRTCIdentityAssertionName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "RTCIdentityAssertion.name" to val.
//
// It returns false if the property cannot be set.
func (this RTCIdentityAssertion) SetName(val js.String) bool {
	return js.True == bindings.SetRTCIdentityAssertionName(
		this.ref,
		val.Ref(),
	)
}

type ValidateAssertionCallbackFunc func(this js.Ref, assertion js.String, origin js.String) js.Ref

func (fn ValidateAssertionCallbackFunc) Register() js.Func[func(assertion js.String, origin js.String) js.Promise[RTCIdentityValidationResult]] {
	return js.RegisterCallback[func(assertion js.String, origin js.String) js.Promise[RTCIdentityValidationResult]](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ValidateAssertionCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ValidateAssertionCallback[T any] struct {
	Fn  func(arg T, this js.Ref, assertion js.String, origin js.String) js.Ref
	Arg T
}

func (cb *ValidateAssertionCallback[T]) Register() js.Func[func(assertion js.String, origin js.String) js.Promise[RTCIdentityValidationResult]] {
	return js.RegisterCallback[func(assertion js.String, origin js.String) js.Promise[RTCIdentityValidationResult]](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ValidateAssertionCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RTCIdentityValidationResult struct {
	// Identity is "RTCIdentityValidationResult.identity"
	//
	// Required
	Identity js.String
	// Contents is "RTCIdentityValidationResult.contents"
	//
	// Required
	Contents js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCIdentityValidationResult with all fields set.
func (p RTCIdentityValidationResult) FromRef(ref js.Ref) RTCIdentityValidationResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCIdentityValidationResult in the application heap.
func (p RTCIdentityValidationResult) New() js.Ref {
	return bindings.RTCIdentityValidationResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCIdentityValidationResult) UpdateFrom(ref js.Ref) {
	bindings.RTCIdentityValidationResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCIdentityValidationResult) Update(ref js.Ref) {
	bindings.RTCIdentityValidationResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCIdentityValidationResult) FreeMembers(recursive bool) {
	js.Free(
		p.Identity.Ref(),
		p.Contents.Ref(),
	)
	p.Identity = p.Identity.FromRef(js.Undefined)
	p.Contents = p.Contents.FromRef(js.Undefined)
}

type RTCIdentityProvider struct {
	// GenerateAssertion is "RTCIdentityProvider.generateAssertion"
	//
	// Required
	GenerateAssertion js.Func[func(contents js.String, origin js.String, options *RTCIdentityProviderOptions) js.Promise[RTCIdentityAssertionResult]]
	// ValidateAssertion is "RTCIdentityProvider.validateAssertion"
	//
	// Required
	ValidateAssertion js.Func[func(assertion js.String, origin js.String) js.Promise[RTCIdentityValidationResult]]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCIdentityProvider with all fields set.
func (p RTCIdentityProvider) FromRef(ref js.Ref) RTCIdentityProvider {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCIdentityProvider in the application heap.
func (p RTCIdentityProvider) New() js.Ref {
	return bindings.RTCIdentityProviderJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCIdentityProvider) UpdateFrom(ref js.Ref) {
	bindings.RTCIdentityProviderJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCIdentityProvider) Update(ref js.Ref) {
	bindings.RTCIdentityProviderJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCIdentityProvider) FreeMembers(recursive bool) {
	js.Free(
		p.GenerateAssertion.Ref(),
		p.ValidateAssertion.Ref(),
	)
	p.GenerateAssertion = p.GenerateAssertion.FromRef(js.Undefined)
	p.ValidateAssertion = p.ValidateAssertion.FromRef(js.Undefined)
}

type RTCIdentityProviderRegistrar struct {
	ref js.Ref
}

func (this RTCIdentityProviderRegistrar) Once() RTCIdentityProviderRegistrar {
	this.ref.Once()
	return this
}

func (this RTCIdentityProviderRegistrar) Ref() js.Ref {
	return this.ref
}

func (this RTCIdentityProviderRegistrar) FromRef(ref js.Ref) RTCIdentityProviderRegistrar {
	this.ref = ref
	return this
}

func (this RTCIdentityProviderRegistrar) Free() {
	this.ref.Free()
}

// HasFuncRegister returns true if the method "RTCIdentityProviderRegistrar.register" exists.
func (this RTCIdentityProviderRegistrar) HasFuncRegister() bool {
	return js.True == bindings.HasFuncRTCIdentityProviderRegistrarRegister(
		this.ref,
	)
}

// FuncRegister returns the method "RTCIdentityProviderRegistrar.register".
func (this RTCIdentityProviderRegistrar) FuncRegister() (fn js.Func[func(idp RTCIdentityProvider)]) {
	bindings.FuncRTCIdentityProviderRegistrarRegister(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Register calls the method "RTCIdentityProviderRegistrar.register".
func (this RTCIdentityProviderRegistrar) Register(idp RTCIdentityProvider) (ret js.Void) {
	bindings.CallRTCIdentityProviderRegistrarRegister(
		this.ref, js.Pointer(&ret),
		js.Pointer(&idp),
	)

	return
}

// TryRegister calls the method "RTCIdentityProviderRegistrar.register"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RTCIdentityProviderRegistrar) TryRegister(idp RTCIdentityProvider) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRTCIdentityProviderRegistrarRegister(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&idp),
	)

	return
}

type RTCIdentityProviderGlobalScope struct {
	WorkerGlobalScope
}

func (this RTCIdentityProviderGlobalScope) Once() RTCIdentityProviderGlobalScope {
	this.ref.Once()
	return this
}

func (this RTCIdentityProviderGlobalScope) Ref() js.Ref {
	return this.WorkerGlobalScope.Ref()
}

func (this RTCIdentityProviderGlobalScope) FromRef(ref js.Ref) RTCIdentityProviderGlobalScope {
	this.WorkerGlobalScope = this.WorkerGlobalScope.FromRef(ref)
	return this
}

func (this RTCIdentityProviderGlobalScope) Free() {
	this.ref.Free()
}

// RtcIdentityProvider returns the value of property "RTCIdentityProviderGlobalScope.rtcIdentityProvider".
//
// It returns ok=false if there is no such property.
func (this RTCIdentityProviderGlobalScope) RtcIdentityProvider() (ret RTCIdentityProviderRegistrar, ok bool) {
	ok = js.True == bindings.GetRTCIdentityProviderGlobalScopeRtcIdentityProvider(
		this.ref, js.Pointer(&ret),
	)
	return
}

type RTCInboundRtpStreamStats struct {
	// TrackIdentifier is "RTCInboundRtpStreamStats.trackIdentifier"
	//
	// Required
	TrackIdentifier js.String
	// Mid is "RTCInboundRtpStreamStats.mid"
	//
	// Optional
	Mid js.String
	// RemoteId is "RTCInboundRtpStreamStats.remoteId"
	//
	// Optional
	RemoteId js.String
	// FramesDecoded is "RTCInboundRtpStreamStats.framesDecoded"
	//
	// Optional
	//
	// NOTE: FFI_USE_FramesDecoded MUST be set to true to make this field effective.
	FramesDecoded uint32
	// KeyFramesDecoded is "RTCInboundRtpStreamStats.keyFramesDecoded"
	//
	// Optional
	//
	// NOTE: FFI_USE_KeyFramesDecoded MUST be set to true to make this field effective.
	KeyFramesDecoded uint32
	// FramesRendered is "RTCInboundRtpStreamStats.framesRendered"
	//
	// Optional
	//
	// NOTE: FFI_USE_FramesRendered MUST be set to true to make this field effective.
	FramesRendered uint32
	// FramesDropped is "RTCInboundRtpStreamStats.framesDropped"
	//
	// Optional
	//
	// NOTE: FFI_USE_FramesDropped MUST be set to true to make this field effective.
	FramesDropped uint32
	// FrameWidth is "RTCInboundRtpStreamStats.frameWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameWidth MUST be set to true to make this field effective.
	FrameWidth uint32
	// FrameHeight is "RTCInboundRtpStreamStats.frameHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameHeight MUST be set to true to make this field effective.
	FrameHeight uint32
	// FramesPerSecond is "RTCInboundRtpStreamStats.framesPerSecond"
	//
	// Optional
	//
	// NOTE: FFI_USE_FramesPerSecond MUST be set to true to make this field effective.
	FramesPerSecond float64
	// QpSum is "RTCInboundRtpStreamStats.qpSum"
	//
	// Optional
	//
	// NOTE: FFI_USE_QpSum MUST be set to true to make this field effective.
	QpSum uint64
	// TotalDecodeTime is "RTCInboundRtpStreamStats.totalDecodeTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalDecodeTime MUST be set to true to make this field effective.
	TotalDecodeTime float64
	// TotalInterFrameDelay is "RTCInboundRtpStreamStats.totalInterFrameDelay"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalInterFrameDelay MUST be set to true to make this field effective.
	TotalInterFrameDelay float64
	// TotalSquaredInterFrameDelay is "RTCInboundRtpStreamStats.totalSquaredInterFrameDelay"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalSquaredInterFrameDelay MUST be set to true to make this field effective.
	TotalSquaredInterFrameDelay float64
	// PauseCount is "RTCInboundRtpStreamStats.pauseCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_PauseCount MUST be set to true to make this field effective.
	PauseCount uint32
	// TotalPausesDuration is "RTCInboundRtpStreamStats.totalPausesDuration"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalPausesDuration MUST be set to true to make this field effective.
	TotalPausesDuration float64
	// FreezeCount is "RTCInboundRtpStreamStats.freezeCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_FreezeCount MUST be set to true to make this field effective.
	FreezeCount uint32
	// TotalFreezesDuration is "RTCInboundRtpStreamStats.totalFreezesDuration"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalFreezesDuration MUST be set to true to make this field effective.
	TotalFreezesDuration float64
	// LastPacketReceivedTimestamp is "RTCInboundRtpStreamStats.lastPacketReceivedTimestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_LastPacketReceivedTimestamp MUST be set to true to make this field effective.
	LastPacketReceivedTimestamp DOMHighResTimeStamp
	// HeaderBytesReceived is "RTCInboundRtpStreamStats.headerBytesReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_HeaderBytesReceived MUST be set to true to make this field effective.
	HeaderBytesReceived uint64
	// PacketsDiscarded is "RTCInboundRtpStreamStats.packetsDiscarded"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsDiscarded MUST be set to true to make this field effective.
	PacketsDiscarded uint64
	// FecBytesReceived is "RTCInboundRtpStreamStats.fecBytesReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_FecBytesReceived MUST be set to true to make this field effective.
	FecBytesReceived uint64
	// FecPacketsReceived is "RTCInboundRtpStreamStats.fecPacketsReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_FecPacketsReceived MUST be set to true to make this field effective.
	FecPacketsReceived uint64
	// FecPacketsDiscarded is "RTCInboundRtpStreamStats.fecPacketsDiscarded"
	//
	// Optional
	//
	// NOTE: FFI_USE_FecPacketsDiscarded MUST be set to true to make this field effective.
	FecPacketsDiscarded uint64
	// BytesReceived is "RTCInboundRtpStreamStats.bytesReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesReceived MUST be set to true to make this field effective.
	BytesReceived uint64
	// NackCount is "RTCInboundRtpStreamStats.nackCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_NackCount MUST be set to true to make this field effective.
	NackCount uint32
	// FirCount is "RTCInboundRtpStreamStats.firCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_FirCount MUST be set to true to make this field effective.
	FirCount uint32
	// PliCount is "RTCInboundRtpStreamStats.pliCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_PliCount MUST be set to true to make this field effective.
	PliCount uint32
	// TotalProcessingDelay is "RTCInboundRtpStreamStats.totalProcessingDelay"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalProcessingDelay MUST be set to true to make this field effective.
	TotalProcessingDelay float64
	// EstimatedPlayoutTimestamp is "RTCInboundRtpStreamStats.estimatedPlayoutTimestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_EstimatedPlayoutTimestamp MUST be set to true to make this field effective.
	EstimatedPlayoutTimestamp DOMHighResTimeStamp
	// JitterBufferDelay is "RTCInboundRtpStreamStats.jitterBufferDelay"
	//
	// Optional
	//
	// NOTE: FFI_USE_JitterBufferDelay MUST be set to true to make this field effective.
	JitterBufferDelay float64
	// JitterBufferTargetDelay is "RTCInboundRtpStreamStats.jitterBufferTargetDelay"
	//
	// Optional
	//
	// NOTE: FFI_USE_JitterBufferTargetDelay MUST be set to true to make this field effective.
	JitterBufferTargetDelay float64
	// JitterBufferEmittedCount is "RTCInboundRtpStreamStats.jitterBufferEmittedCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_JitterBufferEmittedCount MUST be set to true to make this field effective.
	JitterBufferEmittedCount uint64
	// JitterBufferMinimumDelay is "RTCInboundRtpStreamStats.jitterBufferMinimumDelay"
	//
	// Optional
	//
	// NOTE: FFI_USE_JitterBufferMinimumDelay MUST be set to true to make this field effective.
	JitterBufferMinimumDelay float64
	// TotalSamplesReceived is "RTCInboundRtpStreamStats.totalSamplesReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalSamplesReceived MUST be set to true to make this field effective.
	TotalSamplesReceived uint64
	// ConcealedSamples is "RTCInboundRtpStreamStats.concealedSamples"
	//
	// Optional
	//
	// NOTE: FFI_USE_ConcealedSamples MUST be set to true to make this field effective.
	ConcealedSamples uint64
	// SilentConcealedSamples is "RTCInboundRtpStreamStats.silentConcealedSamples"
	//
	// Optional
	//
	// NOTE: FFI_USE_SilentConcealedSamples MUST be set to true to make this field effective.
	SilentConcealedSamples uint64
	// ConcealmentEvents is "RTCInboundRtpStreamStats.concealmentEvents"
	//
	// Optional
	//
	// NOTE: FFI_USE_ConcealmentEvents MUST be set to true to make this field effective.
	ConcealmentEvents uint64
	// InsertedSamplesForDeceleration is "RTCInboundRtpStreamStats.insertedSamplesForDeceleration"
	//
	// Optional
	//
	// NOTE: FFI_USE_InsertedSamplesForDeceleration MUST be set to true to make this field effective.
	InsertedSamplesForDeceleration uint64
	// RemovedSamplesForAcceleration is "RTCInboundRtpStreamStats.removedSamplesForAcceleration"
	//
	// Optional
	//
	// NOTE: FFI_USE_RemovedSamplesForAcceleration MUST be set to true to make this field effective.
	RemovedSamplesForAcceleration uint64
	// AudioLevel is "RTCInboundRtpStreamStats.audioLevel"
	//
	// Optional
	//
	// NOTE: FFI_USE_AudioLevel MUST be set to true to make this field effective.
	AudioLevel float64
	// TotalAudioEnergy is "RTCInboundRtpStreamStats.totalAudioEnergy"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalAudioEnergy MUST be set to true to make this field effective.
	TotalAudioEnergy float64
	// TotalSamplesDuration is "RTCInboundRtpStreamStats.totalSamplesDuration"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalSamplesDuration MUST be set to true to make this field effective.
	TotalSamplesDuration float64
	// FramesReceived is "RTCInboundRtpStreamStats.framesReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_FramesReceived MUST be set to true to make this field effective.
	FramesReceived uint32
	// DecoderImplementation is "RTCInboundRtpStreamStats.decoderImplementation"
	//
	// Optional
	DecoderImplementation js.String
	// PlayoutId is "RTCInboundRtpStreamStats.playoutId"
	//
	// Optional
	PlayoutId js.String
	// PowerEfficientDecoder is "RTCInboundRtpStreamStats.powerEfficientDecoder"
	//
	// Optional
	//
	// NOTE: FFI_USE_PowerEfficientDecoder MUST be set to true to make this field effective.
	PowerEfficientDecoder bool
	// FramesAssembledFromMultiplePackets is "RTCInboundRtpStreamStats.framesAssembledFromMultiplePackets"
	//
	// Optional
	//
	// NOTE: FFI_USE_FramesAssembledFromMultiplePackets MUST be set to true to make this field effective.
	FramesAssembledFromMultiplePackets uint32
	// TotalAssemblyTime is "RTCInboundRtpStreamStats.totalAssemblyTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalAssemblyTime MUST be set to true to make this field effective.
	TotalAssemblyTime float64
	// RetransmittedPacketsReceived is "RTCInboundRtpStreamStats.retransmittedPacketsReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_RetransmittedPacketsReceived MUST be set to true to make this field effective.
	RetransmittedPacketsReceived uint64
	// RetransmittedBytesReceived is "RTCInboundRtpStreamStats.retransmittedBytesReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_RetransmittedBytesReceived MUST be set to true to make this field effective.
	RetransmittedBytesReceived uint64
	// RtxSsrc is "RTCInboundRtpStreamStats.rtxSsrc"
	//
	// Optional
	//
	// NOTE: FFI_USE_RtxSsrc MUST be set to true to make this field effective.
	RtxSsrc uint32
	// FecSsrc is "RTCInboundRtpStreamStats.fecSsrc"
	//
	// Optional
	//
	// NOTE: FFI_USE_FecSsrc MUST be set to true to make this field effective.
	FecSsrc uint32
	// PacketsReceived is "RTCInboundRtpStreamStats.packetsReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsReceived MUST be set to true to make this field effective.
	PacketsReceived uint64
	// PacketsLost is "RTCInboundRtpStreamStats.packetsLost"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsLost MUST be set to true to make this field effective.
	PacketsLost int64
	// Jitter is "RTCInboundRtpStreamStats.jitter"
	//
	// Optional
	//
	// NOTE: FFI_USE_Jitter MUST be set to true to make this field effective.
	Jitter float64
	// Ssrc is "RTCInboundRtpStreamStats.ssrc"
	//
	// Required
	Ssrc uint32
	// Kind is "RTCInboundRtpStreamStats.kind"
	//
	// Required
	Kind js.String
	// TransportId is "RTCInboundRtpStreamStats.transportId"
	//
	// Optional
	TransportId js.String
	// CodecId is "RTCInboundRtpStreamStats.codecId"
	//
	// Optional
	CodecId js.String
	// Timestamp is "RTCInboundRtpStreamStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCInboundRtpStreamStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCInboundRtpStreamStats.id"
	//
	// Required
	Id js.String

	FFI_USE_FramesDecoded                      bool // for FramesDecoded.
	FFI_USE_KeyFramesDecoded                   bool // for KeyFramesDecoded.
	FFI_USE_FramesRendered                     bool // for FramesRendered.
	FFI_USE_FramesDropped                      bool // for FramesDropped.
	FFI_USE_FrameWidth                         bool // for FrameWidth.
	FFI_USE_FrameHeight                        bool // for FrameHeight.
	FFI_USE_FramesPerSecond                    bool // for FramesPerSecond.
	FFI_USE_QpSum                              bool // for QpSum.
	FFI_USE_TotalDecodeTime                    bool // for TotalDecodeTime.
	FFI_USE_TotalInterFrameDelay               bool // for TotalInterFrameDelay.
	FFI_USE_TotalSquaredInterFrameDelay        bool // for TotalSquaredInterFrameDelay.
	FFI_USE_PauseCount                         bool // for PauseCount.
	FFI_USE_TotalPausesDuration                bool // for TotalPausesDuration.
	FFI_USE_FreezeCount                        bool // for FreezeCount.
	FFI_USE_TotalFreezesDuration               bool // for TotalFreezesDuration.
	FFI_USE_LastPacketReceivedTimestamp        bool // for LastPacketReceivedTimestamp.
	FFI_USE_HeaderBytesReceived                bool // for HeaderBytesReceived.
	FFI_USE_PacketsDiscarded                   bool // for PacketsDiscarded.
	FFI_USE_FecBytesReceived                   bool // for FecBytesReceived.
	FFI_USE_FecPacketsReceived                 bool // for FecPacketsReceived.
	FFI_USE_FecPacketsDiscarded                bool // for FecPacketsDiscarded.
	FFI_USE_BytesReceived                      bool // for BytesReceived.
	FFI_USE_NackCount                          bool // for NackCount.
	FFI_USE_FirCount                           bool // for FirCount.
	FFI_USE_PliCount                           bool // for PliCount.
	FFI_USE_TotalProcessingDelay               bool // for TotalProcessingDelay.
	FFI_USE_EstimatedPlayoutTimestamp          bool // for EstimatedPlayoutTimestamp.
	FFI_USE_JitterBufferDelay                  bool // for JitterBufferDelay.
	FFI_USE_JitterBufferTargetDelay            bool // for JitterBufferTargetDelay.
	FFI_USE_JitterBufferEmittedCount           bool // for JitterBufferEmittedCount.
	FFI_USE_JitterBufferMinimumDelay           bool // for JitterBufferMinimumDelay.
	FFI_USE_TotalSamplesReceived               bool // for TotalSamplesReceived.
	FFI_USE_ConcealedSamples                   bool // for ConcealedSamples.
	FFI_USE_SilentConcealedSamples             bool // for SilentConcealedSamples.
	FFI_USE_ConcealmentEvents                  bool // for ConcealmentEvents.
	FFI_USE_InsertedSamplesForDeceleration     bool // for InsertedSamplesForDeceleration.
	FFI_USE_RemovedSamplesForAcceleration      bool // for RemovedSamplesForAcceleration.
	FFI_USE_AudioLevel                         bool // for AudioLevel.
	FFI_USE_TotalAudioEnergy                   bool // for TotalAudioEnergy.
	FFI_USE_TotalSamplesDuration               bool // for TotalSamplesDuration.
	FFI_USE_FramesReceived                     bool // for FramesReceived.
	FFI_USE_PowerEfficientDecoder              bool // for PowerEfficientDecoder.
	FFI_USE_FramesAssembledFromMultiplePackets bool // for FramesAssembledFromMultiplePackets.
	FFI_USE_TotalAssemblyTime                  bool // for TotalAssemblyTime.
	FFI_USE_RetransmittedPacketsReceived       bool // for RetransmittedPacketsReceived.
	FFI_USE_RetransmittedBytesReceived         bool // for RetransmittedBytesReceived.
	FFI_USE_RtxSsrc                            bool // for RtxSsrc.
	FFI_USE_FecSsrc                            bool // for FecSsrc.
	FFI_USE_PacketsReceived                    bool // for PacketsReceived.
	FFI_USE_PacketsLost                        bool // for PacketsLost.
	FFI_USE_Jitter                             bool // for Jitter.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCInboundRtpStreamStats with all fields set.
func (p RTCInboundRtpStreamStats) FromRef(ref js.Ref) RTCInboundRtpStreamStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCInboundRtpStreamStats in the application heap.
func (p RTCInboundRtpStreamStats) New() js.Ref {
	return bindings.RTCInboundRtpStreamStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCInboundRtpStreamStats) UpdateFrom(ref js.Ref) {
	bindings.RTCInboundRtpStreamStatsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCInboundRtpStreamStats) Update(ref js.Ref) {
	bindings.RTCInboundRtpStreamStatsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCInboundRtpStreamStats) FreeMembers(recursive bool) {
	js.Free(
		p.TrackIdentifier.Ref(),
		p.Mid.Ref(),
		p.RemoteId.Ref(),
		p.DecoderImplementation.Ref(),
		p.PlayoutId.Ref(),
		p.Kind.Ref(),
		p.TransportId.Ref(),
		p.CodecId.Ref(),
		p.Id.Ref(),
	)
	p.TrackIdentifier = p.TrackIdentifier.FromRef(js.Undefined)
	p.Mid = p.Mid.FromRef(js.Undefined)
	p.RemoteId = p.RemoteId.FromRef(js.Undefined)
	p.DecoderImplementation = p.DecoderImplementation.FromRef(js.Undefined)
	p.PlayoutId = p.PlayoutId.FromRef(js.Undefined)
	p.Kind = p.Kind.FromRef(js.Undefined)
	p.TransportId = p.TransportId.FromRef(js.Undefined)
	p.CodecId = p.CodecId.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
}

type RTCSdpType uint32

const (
	_ RTCSdpType = iota

	RTCSdpType_OFFER
	RTCSdpType_PRANSWER
	RTCSdpType_ANSWER
	RTCSdpType_ROLLBACK
)

func (RTCSdpType) FromRef(str js.Ref) RTCSdpType {
	return RTCSdpType(bindings.ConstOfRTCSdpType(str))
}

func (x RTCSdpType) String() (string, bool) {
	switch x {
	case RTCSdpType_OFFER:
		return "offer", true
	case RTCSdpType_PRANSWER:
		return "pranswer", true
	case RTCSdpType_ANSWER:
		return "answer", true
	case RTCSdpType_ROLLBACK:
		return "rollback", true
	default:
		return "", false
	}
}

type RTCLocalSessionDescriptionInit struct {
	// Type is "RTCLocalSessionDescriptionInit.type"
	//
	// Optional
	Type RTCSdpType
	// Sdp is "RTCLocalSessionDescriptionInit.sdp"
	//
	// Optional, defaults to "".
	Sdp js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCLocalSessionDescriptionInit with all fields set.
func (p RTCLocalSessionDescriptionInit) FromRef(ref js.Ref) RTCLocalSessionDescriptionInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCLocalSessionDescriptionInit in the application heap.
func (p RTCLocalSessionDescriptionInit) New() js.Ref {
	return bindings.RTCLocalSessionDescriptionInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCLocalSessionDescriptionInit) UpdateFrom(ref js.Ref) {
	bindings.RTCLocalSessionDescriptionInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCLocalSessionDescriptionInit) Update(ref js.Ref) {
	bindings.RTCLocalSessionDescriptionInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCLocalSessionDescriptionInit) FreeMembers(recursive bool) {
	js.Free(
		p.Sdp.Ref(),
	)
	p.Sdp = p.Sdp.FromRef(js.Undefined)
}

type RTCMediaSourceStats struct {
	// TrackIdentifier is "RTCMediaSourceStats.trackIdentifier"
	//
	// Required
	TrackIdentifier js.String
	// Kind is "RTCMediaSourceStats.kind"
	//
	// Required
	Kind js.String
	// Timestamp is "RTCMediaSourceStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCMediaSourceStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCMediaSourceStats.id"
	//
	// Required
	Id js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCMediaSourceStats with all fields set.
func (p RTCMediaSourceStats) FromRef(ref js.Ref) RTCMediaSourceStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCMediaSourceStats in the application heap.
func (p RTCMediaSourceStats) New() js.Ref {
	return bindings.RTCMediaSourceStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCMediaSourceStats) UpdateFrom(ref js.Ref) {
	bindings.RTCMediaSourceStatsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCMediaSourceStats) Update(ref js.Ref) {
	bindings.RTCMediaSourceStatsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCMediaSourceStats) FreeMembers(recursive bool) {
	js.Free(
		p.TrackIdentifier.Ref(),
		p.Kind.Ref(),
		p.Id.Ref(),
	)
	p.TrackIdentifier = p.TrackIdentifier.FromRef(js.Undefined)
	p.Kind = p.Kind.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
}

type RTCOfferAnswerOptions struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCOfferAnswerOptions with all fields set.
func (p RTCOfferAnswerOptions) FromRef(ref js.Ref) RTCOfferAnswerOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCOfferAnswerOptions in the application heap.
func (p RTCOfferAnswerOptions) New() js.Ref {
	return bindings.RTCOfferAnswerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCOfferAnswerOptions) UpdateFrom(ref js.Ref) {
	bindings.RTCOfferAnswerOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCOfferAnswerOptions) Update(ref js.Ref) {
	bindings.RTCOfferAnswerOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCOfferAnswerOptions) FreeMembers(recursive bool) {
}

type RTCOfferOptions struct {
	// IceRestart is "RTCOfferOptions.iceRestart"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IceRestart MUST be set to true to make this field effective.
	IceRestart bool
	// OfferToReceiveAudio is "RTCOfferOptions.offerToReceiveAudio"
	//
	// Optional
	//
	// NOTE: FFI_USE_OfferToReceiveAudio MUST be set to true to make this field effective.
	OfferToReceiveAudio bool
	// OfferToReceiveVideo is "RTCOfferOptions.offerToReceiveVideo"
	//
	// Optional
	//
	// NOTE: FFI_USE_OfferToReceiveVideo MUST be set to true to make this field effective.
	OfferToReceiveVideo bool

	FFI_USE_IceRestart          bool // for IceRestart.
	FFI_USE_OfferToReceiveAudio bool // for OfferToReceiveAudio.
	FFI_USE_OfferToReceiveVideo bool // for OfferToReceiveVideo.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCOfferOptions with all fields set.
func (p RTCOfferOptions) FromRef(ref js.Ref) RTCOfferOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCOfferOptions in the application heap.
func (p RTCOfferOptions) New() js.Ref {
	return bindings.RTCOfferOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCOfferOptions) UpdateFrom(ref js.Ref) {
	bindings.RTCOfferOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCOfferOptions) Update(ref js.Ref) {
	bindings.RTCOfferOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCOfferOptions) FreeMembers(recursive bool) {
}

type RTCQualityLimitationReason uint32

const (
	_ RTCQualityLimitationReason = iota

	RTCQualityLimitationReason_NONE
	RTCQualityLimitationReason_CPU
	RTCQualityLimitationReason_BANDWIDTH
	RTCQualityLimitationReason_OTHER
)

func (RTCQualityLimitationReason) FromRef(str js.Ref) RTCQualityLimitationReason {
	return RTCQualityLimitationReason(bindings.ConstOfRTCQualityLimitationReason(str))
}

func (x RTCQualityLimitationReason) String() (string, bool) {
	switch x {
	case RTCQualityLimitationReason_NONE:
		return "none", true
	case RTCQualityLimitationReason_CPU:
		return "cpu", true
	case RTCQualityLimitationReason_BANDWIDTH:
		return "bandwidth", true
	case RTCQualityLimitationReason_OTHER:
		return "other", true
	default:
		return "", false
	}
}

type RTCOutboundRtpStreamStats struct {
	// Mid is "RTCOutboundRtpStreamStats.mid"
	//
	// Optional
	Mid js.String
	// MediaSourceId is "RTCOutboundRtpStreamStats.mediaSourceId"
	//
	// Optional
	MediaSourceId js.String
	// RemoteId is "RTCOutboundRtpStreamStats.remoteId"
	//
	// Optional
	RemoteId js.String
	// Rid is "RTCOutboundRtpStreamStats.rid"
	//
	// Optional
	Rid js.String
	// HeaderBytesSent is "RTCOutboundRtpStreamStats.headerBytesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_HeaderBytesSent MUST be set to true to make this field effective.
	HeaderBytesSent uint64
	// RetransmittedPacketsSent is "RTCOutboundRtpStreamStats.retransmittedPacketsSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_RetransmittedPacketsSent MUST be set to true to make this field effective.
	RetransmittedPacketsSent uint64
	// RetransmittedBytesSent is "RTCOutboundRtpStreamStats.retransmittedBytesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_RetransmittedBytesSent MUST be set to true to make this field effective.
	RetransmittedBytesSent uint64
	// RtxSsrc is "RTCOutboundRtpStreamStats.rtxSsrc"
	//
	// Optional
	//
	// NOTE: FFI_USE_RtxSsrc MUST be set to true to make this field effective.
	RtxSsrc uint32
	// TargetBitrate is "RTCOutboundRtpStreamStats.targetBitrate"
	//
	// Optional
	//
	// NOTE: FFI_USE_TargetBitrate MUST be set to true to make this field effective.
	TargetBitrate float64
	// TotalEncodedBytesTarget is "RTCOutboundRtpStreamStats.totalEncodedBytesTarget"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalEncodedBytesTarget MUST be set to true to make this field effective.
	TotalEncodedBytesTarget uint64
	// FrameWidth is "RTCOutboundRtpStreamStats.frameWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameWidth MUST be set to true to make this field effective.
	FrameWidth uint32
	// FrameHeight is "RTCOutboundRtpStreamStats.frameHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameHeight MUST be set to true to make this field effective.
	FrameHeight uint32
	// FramesPerSecond is "RTCOutboundRtpStreamStats.framesPerSecond"
	//
	// Optional
	//
	// NOTE: FFI_USE_FramesPerSecond MUST be set to true to make this field effective.
	FramesPerSecond float64
	// FramesSent is "RTCOutboundRtpStreamStats.framesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_FramesSent MUST be set to true to make this field effective.
	FramesSent uint32
	// HugeFramesSent is "RTCOutboundRtpStreamStats.hugeFramesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_HugeFramesSent MUST be set to true to make this field effective.
	HugeFramesSent uint32
	// FramesEncoded is "RTCOutboundRtpStreamStats.framesEncoded"
	//
	// Optional
	//
	// NOTE: FFI_USE_FramesEncoded MUST be set to true to make this field effective.
	FramesEncoded uint32
	// KeyFramesEncoded is "RTCOutboundRtpStreamStats.keyFramesEncoded"
	//
	// Optional
	//
	// NOTE: FFI_USE_KeyFramesEncoded MUST be set to true to make this field effective.
	KeyFramesEncoded uint32
	// QpSum is "RTCOutboundRtpStreamStats.qpSum"
	//
	// Optional
	//
	// NOTE: FFI_USE_QpSum MUST be set to true to make this field effective.
	QpSum uint64
	// TotalEncodeTime is "RTCOutboundRtpStreamStats.totalEncodeTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalEncodeTime MUST be set to true to make this field effective.
	TotalEncodeTime float64
	// TotalPacketSendDelay is "RTCOutboundRtpStreamStats.totalPacketSendDelay"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalPacketSendDelay MUST be set to true to make this field effective.
	TotalPacketSendDelay float64
	// QualityLimitationReason is "RTCOutboundRtpStreamStats.qualityLimitationReason"
	//
	// Optional
	QualityLimitationReason RTCQualityLimitationReason
	// QualityLimitationDurations is "RTCOutboundRtpStreamStats.qualityLimitationDurations"
	//
	// Optional
	QualityLimitationDurations js.Record[float64]
	// QualityLimitationResolutionChanges is "RTCOutboundRtpStreamStats.qualityLimitationResolutionChanges"
	//
	// Optional
	//
	// NOTE: FFI_USE_QualityLimitationResolutionChanges MUST be set to true to make this field effective.
	QualityLimitationResolutionChanges uint32
	// NackCount is "RTCOutboundRtpStreamStats.nackCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_NackCount MUST be set to true to make this field effective.
	NackCount uint32
	// FirCount is "RTCOutboundRtpStreamStats.firCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_FirCount MUST be set to true to make this field effective.
	FirCount uint32
	// PliCount is "RTCOutboundRtpStreamStats.pliCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_PliCount MUST be set to true to make this field effective.
	PliCount uint32
	// EncoderImplementation is "RTCOutboundRtpStreamStats.encoderImplementation"
	//
	// Optional
	EncoderImplementation js.String
	// PowerEfficientEncoder is "RTCOutboundRtpStreamStats.powerEfficientEncoder"
	//
	// Optional
	//
	// NOTE: FFI_USE_PowerEfficientEncoder MUST be set to true to make this field effective.
	PowerEfficientEncoder bool
	// Active is "RTCOutboundRtpStreamStats.active"
	//
	// Optional
	//
	// NOTE: FFI_USE_Active MUST be set to true to make this field effective.
	Active bool
	// ScalabilityMode is "RTCOutboundRtpStreamStats.scalabilityMode"
	//
	// Optional
	ScalabilityMode js.String
	// PacketsSent is "RTCOutboundRtpStreamStats.packetsSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsSent MUST be set to true to make this field effective.
	PacketsSent uint64
	// BytesSent is "RTCOutboundRtpStreamStats.bytesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesSent MUST be set to true to make this field effective.
	BytesSent uint64
	// Ssrc is "RTCOutboundRtpStreamStats.ssrc"
	//
	// Required
	Ssrc uint32
	// Kind is "RTCOutboundRtpStreamStats.kind"
	//
	// Required
	Kind js.String
	// TransportId is "RTCOutboundRtpStreamStats.transportId"
	//
	// Optional
	TransportId js.String
	// CodecId is "RTCOutboundRtpStreamStats.codecId"
	//
	// Optional
	CodecId js.String
	// Timestamp is "RTCOutboundRtpStreamStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCOutboundRtpStreamStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCOutboundRtpStreamStats.id"
	//
	// Required
	Id js.String

	FFI_USE_HeaderBytesSent                    bool // for HeaderBytesSent.
	FFI_USE_RetransmittedPacketsSent           bool // for RetransmittedPacketsSent.
	FFI_USE_RetransmittedBytesSent             bool // for RetransmittedBytesSent.
	FFI_USE_RtxSsrc                            bool // for RtxSsrc.
	FFI_USE_TargetBitrate                      bool // for TargetBitrate.
	FFI_USE_TotalEncodedBytesTarget            bool // for TotalEncodedBytesTarget.
	FFI_USE_FrameWidth                         bool // for FrameWidth.
	FFI_USE_FrameHeight                        bool // for FrameHeight.
	FFI_USE_FramesPerSecond                    bool // for FramesPerSecond.
	FFI_USE_FramesSent                         bool // for FramesSent.
	FFI_USE_HugeFramesSent                     bool // for HugeFramesSent.
	FFI_USE_FramesEncoded                      bool // for FramesEncoded.
	FFI_USE_KeyFramesEncoded                   bool // for KeyFramesEncoded.
	FFI_USE_QpSum                              bool // for QpSum.
	FFI_USE_TotalEncodeTime                    bool // for TotalEncodeTime.
	FFI_USE_TotalPacketSendDelay               bool // for TotalPacketSendDelay.
	FFI_USE_QualityLimitationResolutionChanges bool // for QualityLimitationResolutionChanges.
	FFI_USE_NackCount                          bool // for NackCount.
	FFI_USE_FirCount                           bool // for FirCount.
	FFI_USE_PliCount                           bool // for PliCount.
	FFI_USE_PowerEfficientEncoder              bool // for PowerEfficientEncoder.
	FFI_USE_Active                             bool // for Active.
	FFI_USE_PacketsSent                        bool // for PacketsSent.
	FFI_USE_BytesSent                          bool // for BytesSent.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCOutboundRtpStreamStats with all fields set.
func (p RTCOutboundRtpStreamStats) FromRef(ref js.Ref) RTCOutboundRtpStreamStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCOutboundRtpStreamStats in the application heap.
func (p RTCOutboundRtpStreamStats) New() js.Ref {
	return bindings.RTCOutboundRtpStreamStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCOutboundRtpStreamStats) UpdateFrom(ref js.Ref) {
	bindings.RTCOutboundRtpStreamStatsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCOutboundRtpStreamStats) Update(ref js.Ref) {
	bindings.RTCOutboundRtpStreamStatsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCOutboundRtpStreamStats) FreeMembers(recursive bool) {
	js.Free(
		p.Mid.Ref(),
		p.MediaSourceId.Ref(),
		p.RemoteId.Ref(),
		p.Rid.Ref(),
		p.QualityLimitationDurations.Ref(),
		p.EncoderImplementation.Ref(),
		p.ScalabilityMode.Ref(),
		p.Kind.Ref(),
		p.TransportId.Ref(),
		p.CodecId.Ref(),
		p.Id.Ref(),
	)
	p.Mid = p.Mid.FromRef(js.Undefined)
	p.MediaSourceId = p.MediaSourceId.FromRef(js.Undefined)
	p.RemoteId = p.RemoteId.FromRef(js.Undefined)
	p.Rid = p.Rid.FromRef(js.Undefined)
	p.QualityLimitationDurations = p.QualityLimitationDurations.FromRef(js.Undefined)
	p.EncoderImplementation = p.EncoderImplementation.FromRef(js.Undefined)
	p.ScalabilityMode = p.ScalabilityMode.FromRef(js.Undefined)
	p.Kind = p.Kind.FromRef(js.Undefined)
	p.TransportId = p.TransportId.FromRef(js.Undefined)
	p.CodecId = p.CodecId.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
}

type RTCSessionDescriptionInit struct {
	// Type is "RTCSessionDescriptionInit.type"
	//
	// Required
	Type RTCSdpType
	// Sdp is "RTCSessionDescriptionInit.sdp"
	//
	// Optional, defaults to "".
	Sdp js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCSessionDescriptionInit with all fields set.
func (p RTCSessionDescriptionInit) FromRef(ref js.Ref) RTCSessionDescriptionInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCSessionDescriptionInit in the application heap.
func (p RTCSessionDescriptionInit) New() js.Ref {
	return bindings.RTCSessionDescriptionInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCSessionDescriptionInit) UpdateFrom(ref js.Ref) {
	bindings.RTCSessionDescriptionInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCSessionDescriptionInit) Update(ref js.Ref) {
	bindings.RTCSessionDescriptionInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCSessionDescriptionInit) FreeMembers(recursive bool) {
	js.Free(
		p.Sdp.Ref(),
	)
	p.Sdp = p.Sdp.FromRef(js.Undefined)
}

type RTCSessionDescriptionCallbackFunc func(this js.Ref, description *RTCSessionDescriptionInit) js.Ref

func (fn RTCSessionDescriptionCallbackFunc) Register() js.Func[func(description *RTCSessionDescriptionInit)] {
	return js.RegisterCallback[func(description *RTCSessionDescriptionInit)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RTCSessionDescriptionCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RTCSessionDescriptionInit
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

type RTCSessionDescriptionCallback[T any] struct {
	Fn  func(arg T, this js.Ref, description *RTCSessionDescriptionInit) js.Ref
	Arg T
}

func (cb *RTCSessionDescriptionCallback[T]) Register() js.Func[func(description *RTCSessionDescriptionInit)] {
	return js.RegisterCallback[func(description *RTCSessionDescriptionInit)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RTCSessionDescriptionCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RTCSessionDescriptionInit
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

type RTCPeerConnectionErrorCallbackFunc func(this js.Ref, err DOMException) js.Ref

func (fn RTCPeerConnectionErrorCallbackFunc) Register() js.Func[func(err DOMException)] {
	return js.RegisterCallback[func(err DOMException)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RTCPeerConnectionErrorCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		DOMException{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RTCPeerConnectionErrorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, err DOMException) js.Ref
	Arg T
}

func (cb *RTCPeerConnectionErrorCallback[T]) Register() js.Func[func(err DOMException)] {
	return js.RegisterCallback[func(err DOMException)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RTCPeerConnectionErrorCallback[T]) DispatchCallback(
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

		DOMException{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RTCRtpCodecCapability struct {
	// MimeType is "RTCRtpCodecCapability.mimeType"
	//
	// Required
	MimeType js.String
	// ClockRate is "RTCRtpCodecCapability.clockRate"
	//
	// Required
	ClockRate uint32
	// Channels is "RTCRtpCodecCapability.channels"
	//
	// Optional
	//
	// NOTE: FFI_USE_Channels MUST be set to true to make this field effective.
	Channels uint16
	// SdpFmtpLine is "RTCRtpCodecCapability.sdpFmtpLine"
	//
	// Optional
	SdpFmtpLine js.String

	FFI_USE_Channels bool // for Channels.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpCodecCapability with all fields set.
func (p RTCRtpCodecCapability) FromRef(ref js.Ref) RTCRtpCodecCapability {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCRtpCodecCapability in the application heap.
func (p RTCRtpCodecCapability) New() js.Ref {
	return bindings.RTCRtpCodecCapabilityJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCRtpCodecCapability) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpCodecCapabilityJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCRtpCodecCapability) Update(ref js.Ref) {
	bindings.RTCRtpCodecCapabilityJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCRtpCodecCapability) FreeMembers(recursive bool) {
	js.Free(
		p.MimeType.Ref(),
		p.SdpFmtpLine.Ref(),
	)
	p.MimeType = p.MimeType.FromRef(js.Undefined)
	p.SdpFmtpLine = p.SdpFmtpLine.FromRef(js.Undefined)
}

type RTCRtpHeaderExtensionCapability struct {
	// Uri is "RTCRtpHeaderExtensionCapability.uri"
	//
	// Required
	Uri js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpHeaderExtensionCapability with all fields set.
func (p RTCRtpHeaderExtensionCapability) FromRef(ref js.Ref) RTCRtpHeaderExtensionCapability {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCRtpHeaderExtensionCapability in the application heap.
func (p RTCRtpHeaderExtensionCapability) New() js.Ref {
	return bindings.RTCRtpHeaderExtensionCapabilityJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCRtpHeaderExtensionCapability) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpHeaderExtensionCapabilityJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCRtpHeaderExtensionCapability) Update(ref js.Ref) {
	bindings.RTCRtpHeaderExtensionCapabilityJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCRtpHeaderExtensionCapability) FreeMembers(recursive bool) {
	js.Free(
		p.Uri.Ref(),
	)
	p.Uri = p.Uri.FromRef(js.Undefined)
}

type RTCRtpCapabilities struct {
	// Codecs is "RTCRtpCapabilities.codecs"
	//
	// Required
	Codecs js.Array[RTCRtpCodecCapability]
	// HeaderExtensions is "RTCRtpCapabilities.headerExtensions"
	//
	// Required
	HeaderExtensions js.Array[RTCRtpHeaderExtensionCapability]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCRtpCapabilities with all fields set.
func (p RTCRtpCapabilities) FromRef(ref js.Ref) RTCRtpCapabilities {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCRtpCapabilities in the application heap.
func (p RTCRtpCapabilities) New() js.Ref {
	return bindings.RTCRtpCapabilitiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCRtpCapabilities) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpCapabilitiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCRtpCapabilities) Update(ref js.Ref) {
	bindings.RTCRtpCapabilitiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCRtpCapabilities) FreeMembers(recursive bool) {
	js.Free(
		p.Codecs.Ref(),
		p.HeaderExtensions.Ref(),
	)
	p.Codecs = p.Codecs.FromRef(js.Undefined)
	p.HeaderExtensions = p.HeaderExtensions.FromRef(js.Undefined)
}
