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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Role returns the value of property "RTCIceTransport.role".
//
// The returned bool will be false if there is no such property.
func (this RTCIceTransport) Role() (RTCIceRole, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceTransportRole(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCIceRole(_ret), _ok
}

// Component returns the value of property "RTCIceTransport.component".
//
// The returned bool will be false if there is no such property.
func (this RTCIceTransport) Component() (RTCIceComponent, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceTransportComponent(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCIceComponent(_ret), _ok
}

// State returns the value of property "RTCIceTransport.state".
//
// The returned bool will be false if there is no such property.
func (this RTCIceTransport) State() (RTCIceTransportState, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceTransportState(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCIceTransportState(_ret), _ok
}

// GatheringState returns the value of property "RTCIceTransport.gatheringState".
//
// The returned bool will be false if there is no such property.
func (this RTCIceTransport) GatheringState() (RTCIceGathererState, bool) {
	var _ok bool
	_ret := bindings.GetRTCIceTransportGatheringState(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCIceGathererState(_ret), _ok
}

// GetLocalCandidates calls the method "RTCIceTransport.getLocalCandidates".
//
// The returned bool will be false if there is no such method.
func (this RTCIceTransport) GetLocalCandidates() (js.Array[RTCIceCandidate], bool) {
	var _ok bool
	_ret := bindings.CallRTCIceTransportGetLocalCandidates(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[RTCIceCandidate]{}.FromRef(_ret), _ok
}

// GetLocalCandidatesFunc returns the method "RTCIceTransport.getLocalCandidates".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCIceTransport) GetLocalCandidatesFunc() (fn js.Func[func() js.Array[RTCIceCandidate]]) {
	return fn.FromRef(
		bindings.RTCIceTransportGetLocalCandidatesFunc(
			this.Ref(),
		),
	)
}

// GetRemoteCandidates calls the method "RTCIceTransport.getRemoteCandidates".
//
// The returned bool will be false if there is no such method.
func (this RTCIceTransport) GetRemoteCandidates() (js.Array[RTCIceCandidate], bool) {
	var _ok bool
	_ret := bindings.CallRTCIceTransportGetRemoteCandidates(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[RTCIceCandidate]{}.FromRef(_ret), _ok
}

// GetRemoteCandidatesFunc returns the method "RTCIceTransport.getRemoteCandidates".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCIceTransport) GetRemoteCandidatesFunc() (fn js.Func[func() js.Array[RTCIceCandidate]]) {
	return fn.FromRef(
		bindings.RTCIceTransportGetRemoteCandidatesFunc(
			this.Ref(),
		),
	)
}

// GetSelectedCandidatePair calls the method "RTCIceTransport.getSelectedCandidatePair".
//
// The returned bool will be false if there is no such method.
func (this RTCIceTransport) GetSelectedCandidatePair() (RTCIceCandidatePair, bool) {
	var _ret RTCIceCandidatePair
	_ok := js.True == bindings.CallRTCIceTransportGetSelectedCandidatePair(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetSelectedCandidatePairFunc returns the method "RTCIceTransport.getSelectedCandidatePair".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCIceTransport) GetSelectedCandidatePairFunc() (fn js.Func[func() RTCIceCandidatePair]) {
	return fn.FromRef(
		bindings.RTCIceTransportGetSelectedCandidatePairFunc(
			this.Ref(),
		),
	)
}

// GetLocalParameters calls the method "RTCIceTransport.getLocalParameters".
//
// The returned bool will be false if there is no such method.
func (this RTCIceTransport) GetLocalParameters() (RTCIceParameters, bool) {
	var _ret RTCIceParameters
	_ok := js.True == bindings.CallRTCIceTransportGetLocalParameters(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetLocalParametersFunc returns the method "RTCIceTransport.getLocalParameters".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCIceTransport) GetLocalParametersFunc() (fn js.Func[func() RTCIceParameters]) {
	return fn.FromRef(
		bindings.RTCIceTransportGetLocalParametersFunc(
			this.Ref(),
		),
	)
}

// GetRemoteParameters calls the method "RTCIceTransport.getRemoteParameters".
//
// The returned bool will be false if there is no such method.
func (this RTCIceTransport) GetRemoteParameters() (RTCIceParameters, bool) {
	var _ret RTCIceParameters
	_ok := js.True == bindings.CallRTCIceTransportGetRemoteParameters(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetRemoteParametersFunc returns the method "RTCIceTransport.getRemoteParameters".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCIceTransport) GetRemoteParametersFunc() (fn js.Func[func() RTCIceParameters]) {
	return fn.FromRef(
		bindings.RTCIceTransportGetRemoteParametersFunc(
			this.Ref(),
		),
	)
}

// Gather calls the method "RTCIceTransport.gather".
//
// The returned bool will be false if there is no such method.
func (this RTCIceTransport) Gather(options RTCIceGatherOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCIceTransportGather(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GatherFunc returns the method "RTCIceTransport.gather".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCIceTransport) GatherFunc() (fn js.Func[func(options RTCIceGatherOptions)]) {
	return fn.FromRef(
		bindings.RTCIceTransportGatherFunc(
			this.Ref(),
		),
	)
}

// Gather1 calls the method "RTCIceTransport.gather".
//
// The returned bool will be false if there is no such method.
func (this RTCIceTransport) Gather1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCIceTransportGather1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Gather1Func returns the method "RTCIceTransport.gather".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCIceTransport) Gather1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RTCIceTransportGather1Func(
			this.Ref(),
		),
	)
}

// Start calls the method "RTCIceTransport.start".
//
// The returned bool will be false if there is no such method.
func (this RTCIceTransport) Start(remoteParameters RTCIceParameters, role RTCIceRole) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCIceTransportStart(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&remoteParameters),
		uint32(role),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StartFunc returns the method "RTCIceTransport.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCIceTransport) StartFunc() (fn js.Func[func(remoteParameters RTCIceParameters, role RTCIceRole)]) {
	return fn.FromRef(
		bindings.RTCIceTransportStartFunc(
			this.Ref(),
		),
	)
}

// Start1 calls the method "RTCIceTransport.start".
//
// The returned bool will be false if there is no such method.
func (this RTCIceTransport) Start1(remoteParameters RTCIceParameters) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCIceTransportStart1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&remoteParameters),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Start1Func returns the method "RTCIceTransport.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCIceTransport) Start1Func() (fn js.Func[func(remoteParameters RTCIceParameters)]) {
	return fn.FromRef(
		bindings.RTCIceTransportStart1Func(
			this.Ref(),
		),
	)
}

// Start2 calls the method "RTCIceTransport.start".
//
// The returned bool will be false if there is no such method.
func (this RTCIceTransport) Start2() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCIceTransportStart2(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Start2Func returns the method "RTCIceTransport.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCIceTransport) Start2Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RTCIceTransportStart2Func(
			this.Ref(),
		),
	)
}

// Stop calls the method "RTCIceTransport.stop".
//
// The returned bool will be false if there is no such method.
func (this RTCIceTransport) Stop() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCIceTransportStop(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StopFunc returns the method "RTCIceTransport.stop".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCIceTransport) StopFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RTCIceTransportStopFunc(
			this.Ref(),
		),
	)
}

// AddRemoteCandidate calls the method "RTCIceTransport.addRemoteCandidate".
//
// The returned bool will be false if there is no such method.
func (this RTCIceTransport) AddRemoteCandidate(remoteCandidate RTCIceCandidateInit) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCIceTransportAddRemoteCandidate(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&remoteCandidate),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddRemoteCandidateFunc returns the method "RTCIceTransport.addRemoteCandidate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCIceTransport) AddRemoteCandidateFunc() (fn js.Func[func(remoteCandidate RTCIceCandidateInit)]) {
	return fn.FromRef(
		bindings.RTCIceTransportAddRemoteCandidateFunc(
			this.Ref(),
		),
	)
}

// AddRemoteCandidate1 calls the method "RTCIceTransport.addRemoteCandidate".
//
// The returned bool will be false if there is no such method.
func (this RTCIceTransport) AddRemoteCandidate1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCIceTransportAddRemoteCandidate1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddRemoteCandidate1Func returns the method "RTCIceTransport.addRemoteCandidate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCIceTransport) AddRemoteCandidate1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RTCIceTransportAddRemoteCandidate1Func(
			this.Ref(),
		),
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// IceTransport returns the value of property "RTCDtlsTransport.iceTransport".
//
// The returned bool will be false if there is no such property.
func (this RTCDtlsTransport) IceTransport() (RTCIceTransport, bool) {
	var _ok bool
	_ret := bindings.GetRTCDtlsTransportIceTransport(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCIceTransport{}.FromRef(_ret), _ok
}

// State returns the value of property "RTCDtlsTransport.state".
//
// The returned bool will be false if there is no such property.
func (this RTCDtlsTransport) State() (RTCDtlsTransportState, bool) {
	var _ok bool
	_ret := bindings.GetRTCDtlsTransportState(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCDtlsTransportState(_ret), _ok
}

// GetRemoteCertificates calls the method "RTCDtlsTransport.getRemoteCertificates".
//
// The returned bool will be false if there is no such method.
func (this RTCDtlsTransport) GetRemoteCertificates() (js.Array[js.ArrayBuffer], bool) {
	var _ok bool
	_ret := bindings.CallRTCDtlsTransportGetRemoteCertificates(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[js.ArrayBuffer]{}.FromRef(_ret), _ok
}

// GetRemoteCertificatesFunc returns the method "RTCDtlsTransport.getRemoteCertificates".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCDtlsTransport) GetRemoteCertificatesFunc() (fn js.Func[func() js.Array[js.ArrayBuffer]]) {
	return fn.FromRef(
		bindings.RTCDtlsTransportGetRemoteCertificatesFunc(
			this.Ref(),
		),
	)
}

type RTCEncodedAudioFrameMetadata struct {
	// SynchronizationSource is "RTCEncodedAudioFrameMetadata.synchronizationSource"
	//
	// Optional
	SynchronizationSource uint32
	// PayloadType is "RTCEncodedAudioFrameMetadata.payloadType"
	//
	// Optional
	PayloadType uint8
	// ContributingSources is "RTCEncodedAudioFrameMetadata.contributingSources"
	//
	// Optional
	ContributingSources js.Array[uint32]
	// SequenceNumber is "RTCEncodedAudioFrameMetadata.sequenceNumber"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 RTCEncodedAudioFrameMetadata RTCEncodedAudioFrameMetadata [// RTCEncodedAudioFrameMetadata] [0x1400098de00 0x1400098df40 0x1400099a0a0 0x1400099a140 0x1400098dea0 0x1400099a000 0x1400099a1e0] 0x14000921fb0 {0 0}} in the application heap.
func (p RTCEncodedAudioFrameMetadata) New() js.Ref {
	return bindings.RTCEncodedAudioFrameMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCEncodedAudioFrameMetadata) UpdateFrom(ref js.Ref) {
	bindings.RTCEncodedAudioFrameMetadataJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCEncodedAudioFrameMetadata) Update(ref js.Ref) {
	bindings.RTCEncodedAudioFrameMetadataJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCEncodedAudioFrame struct {
	ref js.Ref
}

func (this RTCEncodedAudioFrame) Once() RTCEncodedAudioFrame {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Timestamp returns the value of property "RTCEncodedAudioFrame.timestamp".
//
// The returned bool will be false if there is no such property.
func (this RTCEncodedAudioFrame) Timestamp() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetRTCEncodedAudioFrameTimestamp(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Data returns the value of property "RTCEncodedAudioFrame.data".
//
// The returned bool will be false if there is no such property.
func (this RTCEncodedAudioFrame) Data() (js.ArrayBuffer, bool) {
	var _ok bool
	_ret := bindings.GetRTCEncodedAudioFrameData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.ArrayBuffer{}.FromRef(_ret), _ok
}

// Data sets the value of property "RTCEncodedAudioFrame.data" to val.
//
// It returns false if the property cannot be set.
func (this RTCEncodedAudioFrame) SetData(val js.ArrayBuffer) bool {
	return js.True == bindings.SetRTCEncodedAudioFrameData(
		this.Ref(),
		val.Ref(),
	)
}

// GetMetadata calls the method "RTCEncodedAudioFrame.getMetadata".
//
// The returned bool will be false if there is no such method.
func (this RTCEncodedAudioFrame) GetMetadata() (RTCEncodedAudioFrameMetadata, bool) {
	var _ret RTCEncodedAudioFrameMetadata
	_ok := js.True == bindings.CallRTCEncodedAudioFrameGetMetadata(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetMetadataFunc returns the method "RTCEncodedAudioFrame.getMetadata".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCEncodedAudioFrame) GetMetadataFunc() (fn js.Func[func() RTCEncodedAudioFrameMetadata]) {
	return fn.FromRef(
		bindings.RTCEncodedAudioFrameGetMetadataFunc(
			this.Ref(),
		),
	)
}

type RTCEncodedVideoFrameMetadata struct {
	// FrameId is "RTCEncodedVideoFrameMetadata.frameId"
	//
	// Optional
	FrameId uint64
	// Dependencies is "RTCEncodedVideoFrameMetadata.dependencies"
	//
	// Optional
	Dependencies js.Array[uint64]
	// Width is "RTCEncodedVideoFrameMetadata.width"
	//
	// Optional
	Width uint16
	// Height is "RTCEncodedVideoFrameMetadata.height"
	//
	// Optional
	Height uint16
	// SpatialIndex is "RTCEncodedVideoFrameMetadata.spatialIndex"
	//
	// Optional
	SpatialIndex uint32
	// TemporalIndex is "RTCEncodedVideoFrameMetadata.temporalIndex"
	//
	// Optional
	TemporalIndex uint32
	// SynchronizationSource is "RTCEncodedVideoFrameMetadata.synchronizationSource"
	//
	// Optional
	SynchronizationSource uint32
	// PayloadType is "RTCEncodedVideoFrameMetadata.payloadType"
	//
	// Optional
	PayloadType uint8
	// ContributingSources is "RTCEncodedVideoFrameMetadata.contributingSources"
	//
	// Optional
	ContributingSources js.Array[uint32]
	// Timestamp is "RTCEncodedVideoFrameMetadata.timestamp"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 RTCEncodedVideoFrameMetadata RTCEncodedVideoFrameMetadata [// RTCEncodedVideoFrameMetadata] [0x1400099a320 0x1400099a500 0x1400099a5a0 0x1400099a6e0 0x1400099a820 0x1400099a960 0x1400099aaa0 0x1400099abe0 0x1400099ad20 0x1400099adc0 0x1400099a3c0 0x1400099a640 0x1400099a780 0x1400099a8c0 0x1400099aa00 0x1400099ab40 0x1400099ac80 0x1400099ae60] 0x1400099c000 {0 0}} in the application heap.
func (p RTCEncodedVideoFrameMetadata) New() js.Ref {
	return bindings.RTCEncodedVideoFrameMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCEncodedVideoFrameMetadata) UpdateFrom(ref js.Ref) {
	bindings.RTCEncodedVideoFrameMetadataJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCEncodedVideoFrameMetadata) Update(ref js.Ref) {
	bindings.RTCEncodedVideoFrameMetadataJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Type returns the value of property "RTCEncodedVideoFrame.type".
//
// The returned bool will be false if there is no such property.
func (this RTCEncodedVideoFrame) Type() (RTCEncodedVideoFrameType, bool) {
	var _ok bool
	_ret := bindings.GetRTCEncodedVideoFrameType(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCEncodedVideoFrameType(_ret), _ok
}

// Timestamp returns the value of property "RTCEncodedVideoFrame.timestamp".
//
// The returned bool will be false if there is no such property.
func (this RTCEncodedVideoFrame) Timestamp() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetRTCEncodedVideoFrameTimestamp(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Data returns the value of property "RTCEncodedVideoFrame.data".
//
// The returned bool will be false if there is no such property.
func (this RTCEncodedVideoFrame) Data() (js.ArrayBuffer, bool) {
	var _ok bool
	_ret := bindings.GetRTCEncodedVideoFrameData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.ArrayBuffer{}.FromRef(_ret), _ok
}

// Data sets the value of property "RTCEncodedVideoFrame.data" to val.
//
// It returns false if the property cannot be set.
func (this RTCEncodedVideoFrame) SetData(val js.ArrayBuffer) bool {
	return js.True == bindings.SetRTCEncodedVideoFrameData(
		this.Ref(),
		val.Ref(),
	)
}

// GetMetadata calls the method "RTCEncodedVideoFrame.getMetadata".
//
// The returned bool will be false if there is no such method.
func (this RTCEncodedVideoFrame) GetMetadata() (RTCEncodedVideoFrameMetadata, bool) {
	var _ret RTCEncodedVideoFrameMetadata
	_ok := js.True == bindings.CallRTCEncodedVideoFrameGetMetadata(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetMetadataFunc returns the method "RTCEncodedVideoFrame.getMetadata".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCEncodedVideoFrame) GetMetadataFunc() (fn js.Func[func() RTCEncodedVideoFrameMetadata]) {
	return fn.FromRef(
		bindings.RTCEncodedVideoFrameGetMetadataFunc(
			this.Ref(),
		),
	)
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
	SdpLineNumber int32
	// SctpCauseCode is "RTCErrorInit.sctpCauseCode"
	//
	// Optional
	SctpCauseCode int32
	// ReceivedAlert is "RTCErrorInit.receivedAlert"
	//
	// Optional
	ReceivedAlert uint32
	// SentAlert is "RTCErrorInit.sentAlert"
	//
	// Optional
	SentAlert uint32
	// HttpRequestStatusCode is "RTCErrorInit.httpRequestStatusCode"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 RTCErrorInit RTCErrorInit [// RTCErrorInit] [0x1400099afa0 0x1400099b040 0x1400099b180 0x1400099b2c0 0x1400099b400 0x1400099b540 0x1400099b0e0 0x1400099b220 0x1400099b360 0x1400099b4a0 0x1400099b5e0] 0x1400099c078 {0 0}} in the application heap.
func (p RTCErrorInit) New() js.Ref {
	return bindings.RTCErrorInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCErrorInit) UpdateFrom(ref js.Ref) {
	bindings.RTCErrorInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCErrorInit) Update(ref js.Ref) {
	bindings.RTCErrorInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewRTCError(init RTCErrorInit, message js.String) RTCError {
	return RTCError{}.FromRef(
		bindings.NewRTCErrorByRTCError(
			js.Pointer(&init),
			message.Ref()),
	)
}

func NewRTCErrorByRTCError1(init RTCErrorInit) RTCError {
	return RTCError{}.FromRef(
		bindings.NewRTCErrorByRTCError1(
			js.Pointer(&init)),
	)
}

type RTCError struct {
	DOMException
}

func (this RTCError) Once() RTCError {
	this.Ref().Once()
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
	this.Ref().Free()
}

// ErrorDetail returns the value of property "RTCError.errorDetail".
//
// The returned bool will be false if there is no such property.
func (this RTCError) ErrorDetail() (RTCErrorDetailType, bool) {
	var _ok bool
	_ret := bindings.GetRTCErrorErrorDetail(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCErrorDetailType(_ret), _ok
}

// SdpLineNumber returns the value of property "RTCError.sdpLineNumber".
//
// The returned bool will be false if there is no such property.
func (this RTCError) SdpLineNumber() (int32, bool) {
	var _ok bool
	_ret := bindings.GetRTCErrorSdpLineNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// SctpCauseCode returns the value of property "RTCError.sctpCauseCode".
//
// The returned bool will be false if there is no such property.
func (this RTCError) SctpCauseCode() (int32, bool) {
	var _ok bool
	_ret := bindings.GetRTCErrorSctpCauseCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// ReceivedAlert returns the value of property "RTCError.receivedAlert".
//
// The returned bool will be false if there is no such property.
func (this RTCError) ReceivedAlert() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetRTCErrorReceivedAlert(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// SentAlert returns the value of property "RTCError.sentAlert".
//
// The returned bool will be false if there is no such property.
func (this RTCError) SentAlert() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetRTCErrorSentAlert(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// HttpRequestStatusCode returns the value of property "RTCError.httpRequestStatusCode".
//
// The returned bool will be false if there is no such property.
func (this RTCError) HttpRequestStatusCode() (int32, bool) {
	var _ok bool
	_ret := bindings.GetRTCErrorHttpRequestStatusCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
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
	Bubbles bool
	// Cancelable is "RTCErrorEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "RTCErrorEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 RTCErrorEventInit RTCErrorEventInit [// RTCErrorEventInit] [0x1400099b720 0x1400099b7c0 0x1400099b900 0x1400099ba40 0x1400099b860 0x1400099b9a0 0x1400099bae0] 0x1400099c420 {0 0}} in the application heap.
func (p RTCErrorEventInit) New() js.Ref {
	return bindings.RTCErrorEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCErrorEventInit) UpdateFrom(ref js.Ref) {
	bindings.RTCErrorEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCErrorEventInit) Update(ref js.Ref) {
	bindings.RTCErrorEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewRTCErrorEvent(typ js.String, eventInitDict RTCErrorEventInit) RTCErrorEvent {
	return RTCErrorEvent{}.FromRef(
		bindings.NewRTCErrorEventByRTCErrorEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type RTCErrorEvent struct {
	Event
}

func (this RTCErrorEvent) Once() RTCErrorEvent {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Error returns the value of property "RTCErrorEvent.error".
//
// The returned bool will be false if there is no such property.
func (this RTCErrorEvent) Error() (RTCError, bool) {
	var _ok bool
	_ret := bindings.GetRTCErrorEventError(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCError{}.FromRef(_ret), _ok
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
	Nominated bool
	// PacketsSent is "RTCIceCandidatePairStats.packetsSent"
	//
	// Optional
	PacketsSent uint64
	// PacketsReceived is "RTCIceCandidatePairStats.packetsReceived"
	//
	// Optional
	PacketsReceived uint64
	// BytesSent is "RTCIceCandidatePairStats.bytesSent"
	//
	// Optional
	BytesSent uint64
	// BytesReceived is "RTCIceCandidatePairStats.bytesReceived"
	//
	// Optional
	BytesReceived uint64
	// LastPacketSentTimestamp is "RTCIceCandidatePairStats.lastPacketSentTimestamp"
	//
	// Optional
	LastPacketSentTimestamp DOMHighResTimeStamp
	// LastPacketReceivedTimestamp is "RTCIceCandidatePairStats.lastPacketReceivedTimestamp"
	//
	// Optional
	LastPacketReceivedTimestamp DOMHighResTimeStamp
	// TotalRoundTripTime is "RTCIceCandidatePairStats.totalRoundTripTime"
	//
	// Optional
	TotalRoundTripTime float64
	// CurrentRoundTripTime is "RTCIceCandidatePairStats.currentRoundTripTime"
	//
	// Optional
	CurrentRoundTripTime float64
	// AvailableOutgoingBitrate is "RTCIceCandidatePairStats.availableOutgoingBitrate"
	//
	// Optional
	AvailableOutgoingBitrate float64
	// AvailableIncomingBitrate is "RTCIceCandidatePairStats.availableIncomingBitrate"
	//
	// Optional
	AvailableIncomingBitrate float64
	// RequestsReceived is "RTCIceCandidatePairStats.requestsReceived"
	//
	// Optional
	RequestsReceived uint64
	// RequestsSent is "RTCIceCandidatePairStats.requestsSent"
	//
	// Optional
	RequestsSent uint64
	// ResponsesReceived is "RTCIceCandidatePairStats.responsesReceived"
	//
	// Optional
	ResponsesReceived uint64
	// ResponsesSent is "RTCIceCandidatePairStats.responsesSent"
	//
	// Optional
	ResponsesSent uint64
	// ConsentRequestsSent is "RTCIceCandidatePairStats.consentRequestsSent"
	//
	// Optional
	ConsentRequestsSent uint64
	// PacketsDiscardedOnSend is "RTCIceCandidatePairStats.packetsDiscardedOnSend"
	//
	// Optional
	PacketsDiscardedOnSend uint32
	// BytesDiscardedOnSend is "RTCIceCandidatePairStats.bytesDiscardedOnSend"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 RTCIceCandidatePairStats RTCIceCandidatePairStats [// RTCIceCandidatePairStats] [0x1400099bb80 0x1400099bc20 0x1400099bcc0 0x1400099bd60 0x1400099be00 0x1400099bf40 0x140009a00a0 0x140009a01e0 0x140009a0320 0x140009a0460 0x140009a05a0 0x140009a06e0 0x140009a0820 0x140009a0960 0x140009a0aa0 0x140009a0be0 0x140009a0d20 0x140009a0e60 0x140009a0fa0 0x140009a10e0 0x140009a1220 0x140009a1360 0x140009a14a0 0x140009a1540 0x140009a15e0 0x1400099bea0 0x140009a0000 0x140009a0140 0x140009a0280 0x140009a03c0 0x140009a0500 0x140009a0640 0x140009a0780 0x140009a08c0 0x140009a0a00 0x140009a0b40 0x140009a0c80 0x140009a0dc0 0x140009a0f00 0x140009a1040 0x140009a1180 0x140009a12c0 0x140009a1400] 0x1400099c450 {0 0}} in the application heap.
func (p RTCIceCandidatePairStats) New() js.Ref {
	return bindings.RTCIceCandidatePairStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCIceCandidatePairStats) UpdateFrom(ref js.Ref) {
	bindings.RTCIceCandidatePairStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCIceCandidatePairStats) Update(ref js.Ref) {
	bindings.RTCIceCandidatePairStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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

// New creates a new {0x140004cc0e0 RTCIceCandidateStats RTCIceCandidateStats [// RTCIceCandidateStats] [0x140009a1680 0x140009a1720 0x140009a17c0 0x140009a1900 0x140009a19a0 0x140009a1a40 0x140009a1b80 0x140009a1c20 0x140009a1cc0 0x140009a1d60 0x140009a1e00 0x140009a1f40 0x140009a8000 0x140009a80a0 0x140009a8140 0x140009a81e0 0x140009a1860 0x140009a1ae0 0x140009a1ea0] 0x1400099c570 {0 0}} in the application heap.
func (p RTCIceCandidateStats) New() js.Ref {
	return bindings.RTCIceCandidateStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCIceCandidateStats) UpdateFrom(ref js.Ref) {
	bindings.RTCIceCandidateStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCIceCandidateStats) Update(ref js.Ref) {
	bindings.RTCIceCandidateStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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

func NewRTCIdentityAssertion(idp js.String, name js.String) RTCIdentityAssertion {
	return RTCIdentityAssertion{}.FromRef(
		bindings.NewRTCIdentityAssertionByRTCIdentityAssertion(
			idp.Ref(),
			name.Ref()),
	)
}

type RTCIdentityAssertion struct {
	ref js.Ref
}

func (this RTCIdentityAssertion) Once() RTCIdentityAssertion {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Idp returns the value of property "RTCIdentityAssertion.idp".
//
// The returned bool will be false if there is no such property.
func (this RTCIdentityAssertion) Idp() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCIdentityAssertionIdp(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Idp sets the value of property "RTCIdentityAssertion.idp" to val.
//
// It returns false if the property cannot be set.
func (this RTCIdentityAssertion) SetIdp(val js.String) bool {
	return js.True == bindings.SetRTCIdentityAssertionIdp(
		this.Ref(),
		val.Ref(),
	)
}

// Name returns the value of property "RTCIdentityAssertion.name".
//
// The returned bool will be false if there is no such property.
func (this RTCIdentityAssertion) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRTCIdentityAssertionName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "RTCIdentityAssertion.name" to val.
//
// It returns false if the property cannot be set.
func (this RTCIdentityAssertion) SetName(val js.String) bool {
	return js.True == bindings.SetRTCIdentityAssertionName(
		this.Ref(),
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
		assert.Throw("invalid", "callback", "invocation")
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

// New creates a new {0x140004cc0e0 RTCIdentityValidationResult RTCIdentityValidationResult [// RTCIdentityValidationResult] [0x140009a83c0 0x140009a8460] 0x1400099c648 {0 0}} in the application heap.
func (p RTCIdentityValidationResult) New() js.Ref {
	return bindings.RTCIdentityValidationResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCIdentityValidationResult) UpdateFrom(ref js.Ref) {
	bindings.RTCIdentityValidationResultJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCIdentityValidationResult) Update(ref js.Ref) {
	bindings.RTCIdentityValidationResultJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCIdentityProvider struct {
	// GenerateAssertion is "RTCIdentityProvider.generateAssertion"
	//
	// Required
	GenerateAssertion js.Func[func(contents js.String, origin js.String, options RTCIdentityProviderOptions) js.Promise[RTCIdentityAssertionResult]]
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

// New creates a new {0x140004cc0e0 RTCIdentityProvider RTCIdentityProvider [// RTCIdentityProvider] [0x140009a8320 0x140009a8500] 0x1400099c630 {0 0}} in the application heap.
func (p RTCIdentityProvider) New() js.Ref {
	return bindings.RTCIdentityProviderJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCIdentityProvider) UpdateFrom(ref js.Ref) {
	bindings.RTCIdentityProviderJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCIdentityProvider) Update(ref js.Ref) {
	bindings.RTCIdentityProviderJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCIdentityProviderRegistrar struct {
	ref js.Ref
}

func (this RTCIdentityProviderRegistrar) Once() RTCIdentityProviderRegistrar {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Register calls the method "RTCIdentityProviderRegistrar.register".
//
// The returned bool will be false if there is no such method.
func (this RTCIdentityProviderRegistrar) Register(idp RTCIdentityProvider) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRTCIdentityProviderRegistrarRegister(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&idp),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RegisterFunc returns the method "RTCIdentityProviderRegistrar.register".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RTCIdentityProviderRegistrar) RegisterFunc() (fn js.Func[func(idp RTCIdentityProvider)]) {
	return fn.FromRef(
		bindings.RTCIdentityProviderRegistrarRegisterFunc(
			this.Ref(),
		),
	)
}

type RTCIdentityProviderGlobalScope struct {
	WorkerGlobalScope
}

func (this RTCIdentityProviderGlobalScope) Once() RTCIdentityProviderGlobalScope {
	this.Ref().Once()
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
	this.Ref().Free()
}

// RtcIdentityProvider returns the value of property "RTCIdentityProviderGlobalScope.rtcIdentityProvider".
//
// The returned bool will be false if there is no such property.
func (this RTCIdentityProviderGlobalScope) RtcIdentityProvider() (RTCIdentityProviderRegistrar, bool) {
	var _ok bool
	_ret := bindings.GetRTCIdentityProviderGlobalScopeRtcIdentityProvider(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCIdentityProviderRegistrar{}.FromRef(_ret), _ok
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
	FramesDecoded uint32
	// KeyFramesDecoded is "RTCInboundRtpStreamStats.keyFramesDecoded"
	//
	// Optional
	KeyFramesDecoded uint32
	// FramesRendered is "RTCInboundRtpStreamStats.framesRendered"
	//
	// Optional
	FramesRendered uint32
	// FramesDropped is "RTCInboundRtpStreamStats.framesDropped"
	//
	// Optional
	FramesDropped uint32
	// FrameWidth is "RTCInboundRtpStreamStats.frameWidth"
	//
	// Optional
	FrameWidth uint32
	// FrameHeight is "RTCInboundRtpStreamStats.frameHeight"
	//
	// Optional
	FrameHeight uint32
	// FramesPerSecond is "RTCInboundRtpStreamStats.framesPerSecond"
	//
	// Optional
	FramesPerSecond float64
	// QpSum is "RTCInboundRtpStreamStats.qpSum"
	//
	// Optional
	QpSum uint64
	// TotalDecodeTime is "RTCInboundRtpStreamStats.totalDecodeTime"
	//
	// Optional
	TotalDecodeTime float64
	// TotalInterFrameDelay is "RTCInboundRtpStreamStats.totalInterFrameDelay"
	//
	// Optional
	TotalInterFrameDelay float64
	// TotalSquaredInterFrameDelay is "RTCInboundRtpStreamStats.totalSquaredInterFrameDelay"
	//
	// Optional
	TotalSquaredInterFrameDelay float64
	// PauseCount is "RTCInboundRtpStreamStats.pauseCount"
	//
	// Optional
	PauseCount uint32
	// TotalPausesDuration is "RTCInboundRtpStreamStats.totalPausesDuration"
	//
	// Optional
	TotalPausesDuration float64
	// FreezeCount is "RTCInboundRtpStreamStats.freezeCount"
	//
	// Optional
	FreezeCount uint32
	// TotalFreezesDuration is "RTCInboundRtpStreamStats.totalFreezesDuration"
	//
	// Optional
	TotalFreezesDuration float64
	// LastPacketReceivedTimestamp is "RTCInboundRtpStreamStats.lastPacketReceivedTimestamp"
	//
	// Optional
	LastPacketReceivedTimestamp DOMHighResTimeStamp
	// HeaderBytesReceived is "RTCInboundRtpStreamStats.headerBytesReceived"
	//
	// Optional
	HeaderBytesReceived uint64
	// PacketsDiscarded is "RTCInboundRtpStreamStats.packetsDiscarded"
	//
	// Optional
	PacketsDiscarded uint64
	// FecBytesReceived is "RTCInboundRtpStreamStats.fecBytesReceived"
	//
	// Optional
	FecBytesReceived uint64
	// FecPacketsReceived is "RTCInboundRtpStreamStats.fecPacketsReceived"
	//
	// Optional
	FecPacketsReceived uint64
	// FecPacketsDiscarded is "RTCInboundRtpStreamStats.fecPacketsDiscarded"
	//
	// Optional
	FecPacketsDiscarded uint64
	// BytesReceived is "RTCInboundRtpStreamStats.bytesReceived"
	//
	// Optional
	BytesReceived uint64
	// NackCount is "RTCInboundRtpStreamStats.nackCount"
	//
	// Optional
	NackCount uint32
	// FirCount is "RTCInboundRtpStreamStats.firCount"
	//
	// Optional
	FirCount uint32
	// PliCount is "RTCInboundRtpStreamStats.pliCount"
	//
	// Optional
	PliCount uint32
	// TotalProcessingDelay is "RTCInboundRtpStreamStats.totalProcessingDelay"
	//
	// Optional
	TotalProcessingDelay float64
	// EstimatedPlayoutTimestamp is "RTCInboundRtpStreamStats.estimatedPlayoutTimestamp"
	//
	// Optional
	EstimatedPlayoutTimestamp DOMHighResTimeStamp
	// JitterBufferDelay is "RTCInboundRtpStreamStats.jitterBufferDelay"
	//
	// Optional
	JitterBufferDelay float64
	// JitterBufferTargetDelay is "RTCInboundRtpStreamStats.jitterBufferTargetDelay"
	//
	// Optional
	JitterBufferTargetDelay float64
	// JitterBufferEmittedCount is "RTCInboundRtpStreamStats.jitterBufferEmittedCount"
	//
	// Optional
	JitterBufferEmittedCount uint64
	// JitterBufferMinimumDelay is "RTCInboundRtpStreamStats.jitterBufferMinimumDelay"
	//
	// Optional
	JitterBufferMinimumDelay float64
	// TotalSamplesReceived is "RTCInboundRtpStreamStats.totalSamplesReceived"
	//
	// Optional
	TotalSamplesReceived uint64
	// ConcealedSamples is "RTCInboundRtpStreamStats.concealedSamples"
	//
	// Optional
	ConcealedSamples uint64
	// SilentConcealedSamples is "RTCInboundRtpStreamStats.silentConcealedSamples"
	//
	// Optional
	SilentConcealedSamples uint64
	// ConcealmentEvents is "RTCInboundRtpStreamStats.concealmentEvents"
	//
	// Optional
	ConcealmentEvents uint64
	// InsertedSamplesForDeceleration is "RTCInboundRtpStreamStats.insertedSamplesForDeceleration"
	//
	// Optional
	InsertedSamplesForDeceleration uint64
	// RemovedSamplesForAcceleration is "RTCInboundRtpStreamStats.removedSamplesForAcceleration"
	//
	// Optional
	RemovedSamplesForAcceleration uint64
	// AudioLevel is "RTCInboundRtpStreamStats.audioLevel"
	//
	// Optional
	AudioLevel float64
	// TotalAudioEnergy is "RTCInboundRtpStreamStats.totalAudioEnergy"
	//
	// Optional
	TotalAudioEnergy float64
	// TotalSamplesDuration is "RTCInboundRtpStreamStats.totalSamplesDuration"
	//
	// Optional
	TotalSamplesDuration float64
	// FramesReceived is "RTCInboundRtpStreamStats.framesReceived"
	//
	// Optional
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
	PowerEfficientDecoder bool
	// FramesAssembledFromMultiplePackets is "RTCInboundRtpStreamStats.framesAssembledFromMultiplePackets"
	//
	// Optional
	FramesAssembledFromMultiplePackets uint32
	// TotalAssemblyTime is "RTCInboundRtpStreamStats.totalAssemblyTime"
	//
	// Optional
	TotalAssemblyTime float64
	// RetransmittedPacketsReceived is "RTCInboundRtpStreamStats.retransmittedPacketsReceived"
	//
	// Optional
	RetransmittedPacketsReceived uint64
	// RetransmittedBytesReceived is "RTCInboundRtpStreamStats.retransmittedBytesReceived"
	//
	// Optional
	RetransmittedBytesReceived uint64
	// RtxSsrc is "RTCInboundRtpStreamStats.rtxSsrc"
	//
	// Optional
	RtxSsrc uint32
	// FecSsrc is "RTCInboundRtpStreamStats.fecSsrc"
	//
	// Optional
	FecSsrc uint32
	// PacketsReceived is "RTCInboundRtpStreamStats.packetsReceived"
	//
	// Optional
	PacketsReceived uint64
	// PacketsLost is "RTCInboundRtpStreamStats.packetsLost"
	//
	// Optional
	PacketsLost int64
	// Jitter is "RTCInboundRtpStreamStats.jitter"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 RTCInboundRtpStreamStats RTCInboundRtpStreamStats [// RTCInboundRtpStreamStats] [0x140009a85a0 0x140009a8640 0x140009a86e0 0x140009a8780 0x140009a88c0 0x140009a8a00 0x140009a8b40 0x140009a8c80 0x140009a8dc0 0x140009a8f00 0x140009a9040 0x140009a9180 0x140009a92c0 0x140009a9400 0x140009a9540 0x140009a9680 0x140009a97c0 0x140009a9900 0x140009a9a40 0x140009a9b80 0x140009a9cc0 0x140009a9e00 0x140009a9f40 0x140009aa0a0 0x140009aa1e0 0x140009aa320 0x140009aa460 0x140009aa5a0 0x140009aa6e0 0x140009aa820 0x140009aa960 0x140009aaaa0 0x140009aabe0 0x140009aad20 0x140009aae60 0x140009aafa0 0x140009ab0e0 0x140009ab220 0x140009ab360 0x140009ab4a0 0x140009ab5e0 0x140009ab720 0x140009ab860 0x140009ab9a0 0x140009abae0 0x140009abb80 0x140009abc20 0x140009abd60 0x140009abea0 0x140009ac000 0x140009ac140 0x140009ac280 0x140009ac3c0 0x140009ac500 0x140009ac640 0x140009ac780 0x140009ac8c0 0x140009ac960 0x140009aca00 0x140009acaa0 0x140009acb40 0x140009acbe0 0x140009acc80 0x140009a8820 0x140009a8960 0x140009a8aa0 0x140009a8be0 0x140009a8d20 0x140009a8e60 0x140009a8fa0 0x140009a90e0 0x140009a9220 0x140009a9360 0x140009a94a0 0x140009a95e0 0x140009a9720 0x140009a9860 0x140009a99a0 0x140009a9ae0 0x140009a9c20 0x140009a9d60 0x140009a9ea0 0x140009aa000 0x140009aa140 0x140009aa280 0x140009aa3c0 0x140009aa500 0x140009aa640 0x140009aa780 0x140009aa8c0 0x140009aaa00 0x140009aab40 0x140009aac80 0x140009aadc0 0x140009aaf00 0x140009ab040 0x140009ab180 0x140009ab2c0 0x140009ab400 0x140009ab540 0x140009ab680 0x140009ab7c0 0x140009ab900 0x140009aba40 0x140009abcc0 0x140009abe00 0x140009abf40 0x140009ac0a0 0x140009ac1e0 0x140009ac320 0x140009ac460 0x140009ac5a0 0x140009ac6e0 0x140009ac820] 0x1400099c660 {0 0}} in the application heap.
func (p RTCInboundRtpStreamStats) New() js.Ref {
	return bindings.RTCInboundRtpStreamStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCInboundRtpStreamStats) UpdateFrom(ref js.Ref) {
	bindings.RTCInboundRtpStreamStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCInboundRtpStreamStats) Update(ref js.Ref) {
	bindings.RTCInboundRtpStreamStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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

// New creates a new {0x140004cc0e0 RTCLocalSessionDescriptionInit RTCLocalSessionDescriptionInit [// RTCLocalSessionDescriptionInit] [0x140009acd20 0x140009acdc0] 0x1400099c8d0 {0 0}} in the application heap.
func (p RTCLocalSessionDescriptionInit) New() js.Ref {
	return bindings.RTCLocalSessionDescriptionInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCLocalSessionDescriptionInit) UpdateFrom(ref js.Ref) {
	bindings.RTCLocalSessionDescriptionInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCLocalSessionDescriptionInit) Update(ref js.Ref) {
	bindings.RTCLocalSessionDescriptionInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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

// New creates a new {0x140004cc0e0 RTCMediaSourceStats RTCMediaSourceStats [// RTCMediaSourceStats] [0x140009ace60 0x140009acf00 0x140009acfa0 0x140009ad040 0x140009ad0e0] 0x1400099c930 {0 0}} in the application heap.
func (p RTCMediaSourceStats) New() js.Ref {
	return bindings.RTCMediaSourceStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCMediaSourceStats) UpdateFrom(ref js.Ref) {
	bindings.RTCMediaSourceStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCMediaSourceStats) Update(ref js.Ref) {
	bindings.RTCMediaSourceStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCOfferAnswerOptions struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCOfferAnswerOptions with all fields set.
func (p RTCOfferAnswerOptions) FromRef(ref js.Ref) RTCOfferAnswerOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCOfferAnswerOptions RTCOfferAnswerOptions [// RTCOfferAnswerOptions] [] 0x1400099c9a8 {0 0}} in the application heap.
func (p RTCOfferAnswerOptions) New() js.Ref {
	return bindings.RTCOfferAnswerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCOfferAnswerOptions) UpdateFrom(ref js.Ref) {
	bindings.RTCOfferAnswerOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCOfferAnswerOptions) Update(ref js.Ref) {
	bindings.RTCOfferAnswerOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCOfferOptions struct {
	// IceRestart is "RTCOfferOptions.iceRestart"
	//
	// Optional, defaults to false.
	IceRestart bool
	// OfferToReceiveAudio is "RTCOfferOptions.offerToReceiveAudio"
	//
	// Optional
	OfferToReceiveAudio bool
	// OfferToReceiveVideo is "RTCOfferOptions.offerToReceiveVideo"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 RTCOfferOptions RTCOfferOptions [// RTCOfferOptions] [0x140009ad180 0x140009ad2c0 0x140009ad400 0x140009ad220 0x140009ad360 0x140009ad4a0] 0x1400099c9d8 {0 0}} in the application heap.
func (p RTCOfferOptions) New() js.Ref {
	return bindings.RTCOfferOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCOfferOptions) UpdateFrom(ref js.Ref) {
	bindings.RTCOfferOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCOfferOptions) Update(ref js.Ref) {
	bindings.RTCOfferOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	HeaderBytesSent uint64
	// RetransmittedPacketsSent is "RTCOutboundRtpStreamStats.retransmittedPacketsSent"
	//
	// Optional
	RetransmittedPacketsSent uint64
	// RetransmittedBytesSent is "RTCOutboundRtpStreamStats.retransmittedBytesSent"
	//
	// Optional
	RetransmittedBytesSent uint64
	// RtxSsrc is "RTCOutboundRtpStreamStats.rtxSsrc"
	//
	// Optional
	RtxSsrc uint32
	// TargetBitrate is "RTCOutboundRtpStreamStats.targetBitrate"
	//
	// Optional
	TargetBitrate float64
	// TotalEncodedBytesTarget is "RTCOutboundRtpStreamStats.totalEncodedBytesTarget"
	//
	// Optional
	TotalEncodedBytesTarget uint64
	// FrameWidth is "RTCOutboundRtpStreamStats.frameWidth"
	//
	// Optional
	FrameWidth uint32
	// FrameHeight is "RTCOutboundRtpStreamStats.frameHeight"
	//
	// Optional
	FrameHeight uint32
	// FramesPerSecond is "RTCOutboundRtpStreamStats.framesPerSecond"
	//
	// Optional
	FramesPerSecond float64
	// FramesSent is "RTCOutboundRtpStreamStats.framesSent"
	//
	// Optional
	FramesSent uint32
	// HugeFramesSent is "RTCOutboundRtpStreamStats.hugeFramesSent"
	//
	// Optional
	HugeFramesSent uint32
	// FramesEncoded is "RTCOutboundRtpStreamStats.framesEncoded"
	//
	// Optional
	FramesEncoded uint32
	// KeyFramesEncoded is "RTCOutboundRtpStreamStats.keyFramesEncoded"
	//
	// Optional
	KeyFramesEncoded uint32
	// QpSum is "RTCOutboundRtpStreamStats.qpSum"
	//
	// Optional
	QpSum uint64
	// TotalEncodeTime is "RTCOutboundRtpStreamStats.totalEncodeTime"
	//
	// Optional
	TotalEncodeTime float64
	// TotalPacketSendDelay is "RTCOutboundRtpStreamStats.totalPacketSendDelay"
	//
	// Optional
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
	QualityLimitationResolutionChanges uint32
	// NackCount is "RTCOutboundRtpStreamStats.nackCount"
	//
	// Optional
	NackCount uint32
	// FirCount is "RTCOutboundRtpStreamStats.firCount"
	//
	// Optional
	FirCount uint32
	// PliCount is "RTCOutboundRtpStreamStats.pliCount"
	//
	// Optional
	PliCount uint32
	// EncoderImplementation is "RTCOutboundRtpStreamStats.encoderImplementation"
	//
	// Optional
	EncoderImplementation js.String
	// PowerEfficientEncoder is "RTCOutboundRtpStreamStats.powerEfficientEncoder"
	//
	// Optional
	PowerEfficientEncoder bool
	// Active is "RTCOutboundRtpStreamStats.active"
	//
	// Optional
	Active bool
	// ScalabilityMode is "RTCOutboundRtpStreamStats.scalabilityMode"
	//
	// Optional
	ScalabilityMode js.String
	// PacketsSent is "RTCOutboundRtpStreamStats.packetsSent"
	//
	// Optional
	PacketsSent uint64
	// BytesSent is "RTCOutboundRtpStreamStats.bytesSent"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 RTCOutboundRtpStreamStats RTCOutboundRtpStreamStats [// RTCOutboundRtpStreamStats] [0x140009ad540 0x140009ad5e0 0x140009ad680 0x140009ad720 0x140009ad7c0 0x140009ad900 0x140009ada40 0x140009adb80 0x140009adcc0 0x140009ade00 0x140009adf40 0x140009b00a0 0x140009b01e0 0x140009b0320 0x140009b0460 0x140009b05a0 0x140009b06e0 0x140009b0820 0x140009b0960 0x140009b0aa0 0x140009b0be0 0x140009b0c80 0x140009b0d20 0x140009b0e60 0x140009b0fa0 0x140009b10e0 0x140009b1220 0x140009b12c0 0x140009b1400 0x140009b1540 0x140009b15e0 0x140009b1720 0x140009b1860 0x140009b1900 0x140009b19a0 0x140009b1a40 0x140009b1ae0 0x140009b1b80 0x140009b1c20 0x140009ad860 0x140009ad9a0 0x140009adae0 0x140009adc20 0x140009add60 0x140009adea0 0x140009b0000 0x140009b0140 0x140009b0280 0x140009b03c0 0x140009b0500 0x140009b0640 0x140009b0780 0x140009b08c0 0x140009b0a00 0x140009b0b40 0x140009b0dc0 0x140009b0f00 0x140009b1040 0x140009b1180 0x140009b1360 0x140009b14a0 0x140009b1680 0x140009b17c0] 0x1400099ca08 {0 0}} in the application heap.
func (p RTCOutboundRtpStreamStats) New() js.Ref {
	return bindings.RTCOutboundRtpStreamStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCOutboundRtpStreamStats) UpdateFrom(ref js.Ref) {
	bindings.RTCOutboundRtpStreamStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCOutboundRtpStreamStats) Update(ref js.Ref) {
	bindings.RTCOutboundRtpStreamStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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

// New creates a new {0x140004cc0e0 RTCSessionDescriptionInit RTCSessionDescriptionInit [// RTCSessionDescriptionInit] [0x140009b1cc0 0x140009b1d60] 0x1400099cc18 {0 0}} in the application heap.
func (p RTCSessionDescriptionInit) New() js.Ref {
	return bindings.RTCSessionDescriptionInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCSessionDescriptionInit) UpdateFrom(ref js.Ref) {
	bindings.RTCSessionDescriptionInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCSessionDescriptionInit) Update(ref js.Ref) {
	bindings.RTCSessionDescriptionInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCSessionDescriptionCallbackFunc func(this js.Ref, description RTCSessionDescriptionInit) js.Ref

func (fn RTCSessionDescriptionCallbackFunc) Register() js.Func[func(description RTCSessionDescriptionInit)] {
	return js.RegisterCallback[func(description RTCSessionDescriptionInit)](
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

	if ctx.Return(fn(
		args[0],

		RTCSessionDescriptionInit{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RTCSessionDescriptionCallback[T any] struct {
	Fn  func(arg T, this js.Ref, description RTCSessionDescriptionInit) js.Ref
	Arg T
}

func (cb *RTCSessionDescriptionCallback[T]) Register() js.Func[func(description RTCSessionDescriptionInit)] {
	return js.RegisterCallback[func(description RTCSessionDescriptionInit)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RTCSessionDescriptionCallback[T]) DispatchCallback(
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

		RTCSessionDescriptionInit{}.FromRef(args[0+1]),
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
		assert.Throw("invalid", "callback", "invocation")
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

// New creates a new {0x140004cc0e0 RTCRtpCodecCapability RTCRtpCodecCapability [// RTCRtpCodecCapability] [0x140009b1e00 0x140009b1ea0 0x140009b1f40 0x140009d40a0 0x140009d4000] 0x1400099cde0 {0 0}} in the application heap.
func (p RTCRtpCodecCapability) New() js.Ref {
	return bindings.RTCRtpCodecCapabilityJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtpCodecCapability) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpCodecCapabilityJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtpCodecCapability) Update(ref js.Ref) {
	bindings.RTCRtpCodecCapabilityJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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

// New creates a new {0x140004cc0e0 RTCRtpHeaderExtensionCapability RTCRtpHeaderExtensionCapability [// RTCRtpHeaderExtensionCapability] [0x140009d41e0] 0x1400099cdf8 {0 0}} in the application heap.
func (p RTCRtpHeaderExtensionCapability) New() js.Ref {
	return bindings.RTCRtpHeaderExtensionCapabilityJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtpHeaderExtensionCapability) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpHeaderExtensionCapabilityJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtpHeaderExtensionCapability) Update(ref js.Ref) {
	bindings.RTCRtpHeaderExtensionCapabilityJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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

// New creates a new {0x140004cc0e0 RTCRtpCapabilities RTCRtpCapabilities [// RTCRtpCapabilities] [0x140009d4140 0x140009d4280] 0x1400099cdb0 {0 0}} in the application heap.
func (p RTCRtpCapabilities) New() js.Ref {
	return bindings.RTCRtpCapabilitiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCRtpCapabilities) UpdateFrom(ref js.Ref) {
	bindings.RTCRtpCapabilitiesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCRtpCapabilities) Update(ref js.Ref) {
	bindings.RTCRtpCapabilitiesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}
