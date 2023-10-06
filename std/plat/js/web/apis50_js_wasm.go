// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

type RTCTrackEventInit struct {
	// Receiver is "RTCTrackEventInit.receiver"
	//
	// Required
	Receiver RTCRtpReceiver
	// Track is "RTCTrackEventInit.track"
	//
	// Required
	Track MediaStreamTrack
	// Streams is "RTCTrackEventInit.streams"
	//
	// Optional, defaults to [].
	Streams js.Array[MediaStream]
	// Transceiver is "RTCTrackEventInit.transceiver"
	//
	// Required
	Transceiver RTCRtpTransceiver
	// Bubbles is "RTCTrackEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "RTCTrackEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "RTCTrackEventInit.composed"
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

// FromRef calls UpdateFrom and returns a RTCTrackEventInit with all fields set.
func (p RTCTrackEventInit) FromRef(ref js.Ref) RTCTrackEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCTrackEventInit in the application heap.
func (p RTCTrackEventInit) New() js.Ref {
	return bindings.RTCTrackEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCTrackEventInit) UpdateFrom(ref js.Ref) {
	bindings.RTCTrackEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCTrackEventInit) Update(ref js.Ref) {
	bindings.RTCTrackEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewRTCTrackEvent(typ js.String, eventInitDict RTCTrackEventInit) (ret RTCTrackEvent) {
	ret.ref = bindings.NewRTCTrackEventByRTCTrackEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type RTCTrackEvent struct {
	Event
}

func (this RTCTrackEvent) Once() RTCTrackEvent {
	this.Ref().Once()
	return this
}

func (this RTCTrackEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this RTCTrackEvent) FromRef(ref js.Ref) RTCTrackEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this RTCTrackEvent) Free() {
	this.Ref().Free()
}

// Receiver returns the value of property "RTCTrackEvent.receiver".
//
// It returns ok=false if there is no such property.
func (this RTCTrackEvent) Receiver() (ret RTCRtpReceiver, ok bool) {
	ok = js.True == bindings.GetRTCTrackEventReceiver(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Track returns the value of property "RTCTrackEvent.track".
//
// It returns ok=false if there is no such property.
func (this RTCTrackEvent) Track() (ret MediaStreamTrack, ok bool) {
	ok = js.True == bindings.GetRTCTrackEventTrack(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Streams returns the value of property "RTCTrackEvent.streams".
//
// It returns ok=false if there is no such property.
func (this RTCTrackEvent) Streams() (ret js.FrozenArray[MediaStream], ok bool) {
	ok = js.True == bindings.GetRTCTrackEventStreams(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Transceiver returns the value of property "RTCTrackEvent.transceiver".
//
// It returns ok=false if there is no such property.
func (this RTCTrackEvent) Transceiver() (ret RTCRtpTransceiver, ok bool) {
	ok = js.True == bindings.GetRTCTrackEventTransceiver(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewRTCTransformEvent(typ js.String, eventInitDict EventInit) (ret RTCTransformEvent) {
	ret.ref = bindings.NewRTCTransformEventByRTCTransformEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewRTCTransformEventByRTCTransformEvent1(typ js.String) (ret RTCTransformEvent) {
	ret.ref = bindings.NewRTCTransformEventByRTCTransformEvent1(
		typ.Ref())
	return
}

type RTCTransformEvent struct {
	Event
}

func (this RTCTransformEvent) Once() RTCTransformEvent {
	this.Ref().Once()
	return this
}

func (this RTCTransformEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this RTCTransformEvent) FromRef(ref js.Ref) RTCTransformEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this RTCTransformEvent) Free() {
	this.Ref().Free()
}

// Transformer returns the value of property "RTCTransformEvent.transformer".
//
// It returns ok=false if there is no such property.
func (this RTCTransformEvent) Transformer() (ret RTCRtpScriptTransformer, ok bool) {
	ok = js.True == bindings.GetRTCTransformEventTransformer(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type RTCTransportStats struct {
	// PacketsSent is "RTCTransportStats.packetsSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsSent MUST be set to true to make this field effective.
	PacketsSent uint64
	// PacketsReceived is "RTCTransportStats.packetsReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsReceived MUST be set to true to make this field effective.
	PacketsReceived uint64
	// BytesSent is "RTCTransportStats.bytesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesSent MUST be set to true to make this field effective.
	BytesSent uint64
	// BytesReceived is "RTCTransportStats.bytesReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesReceived MUST be set to true to make this field effective.
	BytesReceived uint64
	// IceRole is "RTCTransportStats.iceRole"
	//
	// Optional
	IceRole RTCIceRole
	// IceLocalUsernameFragment is "RTCTransportStats.iceLocalUsernameFragment"
	//
	// Optional
	IceLocalUsernameFragment js.String
	// DtlsState is "RTCTransportStats.dtlsState"
	//
	// Required
	DtlsState RTCDtlsTransportState
	// IceState is "RTCTransportStats.iceState"
	//
	// Optional
	IceState RTCIceTransportState
	// SelectedCandidatePairId is "RTCTransportStats.selectedCandidatePairId"
	//
	// Optional
	SelectedCandidatePairId js.String
	// LocalCertificateId is "RTCTransportStats.localCertificateId"
	//
	// Optional
	LocalCertificateId js.String
	// RemoteCertificateId is "RTCTransportStats.remoteCertificateId"
	//
	// Optional
	RemoteCertificateId js.String
	// TlsVersion is "RTCTransportStats.tlsVersion"
	//
	// Optional
	TlsVersion js.String
	// DtlsCipher is "RTCTransportStats.dtlsCipher"
	//
	// Optional
	DtlsCipher js.String
	// DtlsRole is "RTCTransportStats.dtlsRole"
	//
	// Optional
	DtlsRole RTCDtlsRole
	// SrtpCipher is "RTCTransportStats.srtpCipher"
	//
	// Optional
	SrtpCipher js.String
	// SelectedCandidatePairChanges is "RTCTransportStats.selectedCandidatePairChanges"
	//
	// Optional
	//
	// NOTE: FFI_USE_SelectedCandidatePairChanges MUST be set to true to make this field effective.
	SelectedCandidatePairChanges uint32
	// Timestamp is "RTCTransportStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCTransportStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCTransportStats.id"
	//
	// Required
	Id js.String

	FFI_USE_PacketsSent                  bool // for PacketsSent.
	FFI_USE_PacketsReceived              bool // for PacketsReceived.
	FFI_USE_BytesSent                    bool // for BytesSent.
	FFI_USE_BytesReceived                bool // for BytesReceived.
	FFI_USE_SelectedCandidatePairChanges bool // for SelectedCandidatePairChanges.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCTransportStats with all fields set.
func (p RTCTransportStats) FromRef(ref js.Ref) RTCTransportStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCTransportStats in the application heap.
func (p RTCTransportStats) New() js.Ref {
	return bindings.RTCTransportStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCTransportStats) UpdateFrom(ref js.Ref) {
	bindings.RTCTransportStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCTransportStats) Update(ref js.Ref) {
	bindings.RTCTransportStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCVideoSourceStats struct {
	// Width is "RTCVideoSourceStats.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width uint32
	// Height is "RTCVideoSourceStats.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height uint32
	// Frames is "RTCVideoSourceStats.frames"
	//
	// Optional
	//
	// NOTE: FFI_USE_Frames MUST be set to true to make this field effective.
	Frames uint32
	// FramesPerSecond is "RTCVideoSourceStats.framesPerSecond"
	//
	// Optional
	//
	// NOTE: FFI_USE_FramesPerSecond MUST be set to true to make this field effective.
	FramesPerSecond float64
	// TrackIdentifier is "RTCVideoSourceStats.trackIdentifier"
	//
	// Required
	TrackIdentifier js.String
	// Kind is "RTCVideoSourceStats.kind"
	//
	// Required
	Kind js.String
	// Timestamp is "RTCVideoSourceStats.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// Type is "RTCVideoSourceStats.type"
	//
	// Required
	Type RTCStatsType
	// Id is "RTCVideoSourceStats.id"
	//
	// Required
	Id js.String

	FFI_USE_Width           bool // for Width.
	FFI_USE_Height          bool // for Height.
	FFI_USE_Frames          bool // for Frames.
	FFI_USE_FramesPerSecond bool // for FramesPerSecond.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCVideoSourceStats with all fields set.
func (p RTCVideoSourceStats) FromRef(ref js.Ref) RTCVideoSourceStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCVideoSourceStats in the application heap.
func (p RTCVideoSourceStats) New() js.Ref {
	return bindings.RTCVideoSourceStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCVideoSourceStats) UpdateFrom(ref js.Ref) {
	bindings.RTCVideoSourceStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCVideoSourceStats) Update(ref js.Ref) {
	bindings.RTCVideoSourceStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ReadableStreamBYOBRequest struct {
	ref js.Ref
}

func (this ReadableStreamBYOBRequest) Once() ReadableStreamBYOBRequest {
	this.Ref().Once()
	return this
}

func (this ReadableStreamBYOBRequest) Ref() js.Ref {
	return this.ref
}

func (this ReadableStreamBYOBRequest) FromRef(ref js.Ref) ReadableStreamBYOBRequest {
	this.ref = ref
	return this
}

func (this ReadableStreamBYOBRequest) Free() {
	this.Ref().Free()
}

// View returns the value of property "ReadableStreamBYOBRequest.view".
//
// It returns ok=false if there is no such property.
func (this ReadableStreamBYOBRequest) View() (ret js.ArrayBufferView, ok bool) {
	ok = js.True == bindings.GetReadableStreamBYOBRequestView(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRespond returns true if the method "ReadableStreamBYOBRequest.respond" exists.
func (this ReadableStreamBYOBRequest) HasRespond() bool {
	return js.True == bindings.HasReadableStreamBYOBRequestRespond(
		this.Ref(),
	)
}

// RespondFunc returns the method "ReadableStreamBYOBRequest.respond".
func (this ReadableStreamBYOBRequest) RespondFunc() (fn js.Func[func(bytesWritten uint64)]) {
	return fn.FromRef(
		bindings.ReadableStreamBYOBRequestRespondFunc(
			this.Ref(),
		),
	)
}

// Respond calls the method "ReadableStreamBYOBRequest.respond".
func (this ReadableStreamBYOBRequest) Respond(bytesWritten uint64) (ret js.Void) {
	bindings.CallReadableStreamBYOBRequestRespond(
		this.Ref(), js.Pointer(&ret),
		float64(bytesWritten),
	)

	return
}

// TryRespond calls the method "ReadableStreamBYOBRequest.respond"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamBYOBRequest) TryRespond(bytesWritten uint64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamBYOBRequestRespond(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(bytesWritten),
	)

	return
}

// HasRespondWithNewView returns true if the method "ReadableStreamBYOBRequest.respondWithNewView" exists.
func (this ReadableStreamBYOBRequest) HasRespondWithNewView() bool {
	return js.True == bindings.HasReadableStreamBYOBRequestRespondWithNewView(
		this.Ref(),
	)
}

// RespondWithNewViewFunc returns the method "ReadableStreamBYOBRequest.respondWithNewView".
func (this ReadableStreamBYOBRequest) RespondWithNewViewFunc() (fn js.Func[func(view js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.ReadableStreamBYOBRequestRespondWithNewViewFunc(
			this.Ref(),
		),
	)
}

// RespondWithNewView calls the method "ReadableStreamBYOBRequest.respondWithNewView".
func (this ReadableStreamBYOBRequest) RespondWithNewView(view js.ArrayBufferView) (ret js.Void) {
	bindings.CallReadableStreamBYOBRequestRespondWithNewView(
		this.Ref(), js.Pointer(&ret),
		view.Ref(),
	)

	return
}

// TryRespondWithNewView calls the method "ReadableStreamBYOBRequest.respondWithNewView"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamBYOBRequest) TryRespondWithNewView(view js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamBYOBRequestRespondWithNewView(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		view.Ref(),
	)

	return
}

type ReadableByteStreamController struct {
	ref js.Ref
}

func (this ReadableByteStreamController) Once() ReadableByteStreamController {
	this.Ref().Once()
	return this
}

func (this ReadableByteStreamController) Ref() js.Ref {
	return this.ref
}

func (this ReadableByteStreamController) FromRef(ref js.Ref) ReadableByteStreamController {
	this.ref = ref
	return this
}

func (this ReadableByteStreamController) Free() {
	this.Ref().Free()
}

// ByobRequest returns the value of property "ReadableByteStreamController.byobRequest".
//
// It returns ok=false if there is no such property.
func (this ReadableByteStreamController) ByobRequest() (ret ReadableStreamBYOBRequest, ok bool) {
	ok = js.True == bindings.GetReadableByteStreamControllerByobRequest(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DesiredSize returns the value of property "ReadableByteStreamController.desiredSize".
//
// It returns ok=false if there is no such property.
func (this ReadableByteStreamController) DesiredSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetReadableByteStreamControllerDesiredSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasClose returns true if the method "ReadableByteStreamController.close" exists.
func (this ReadableByteStreamController) HasClose() bool {
	return js.True == bindings.HasReadableByteStreamControllerClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "ReadableByteStreamController.close".
func (this ReadableByteStreamController) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReadableByteStreamControllerCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "ReadableByteStreamController.close".
func (this ReadableByteStreamController) Close() (ret js.Void) {
	bindings.CallReadableByteStreamControllerClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "ReadableByteStreamController.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableByteStreamController) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableByteStreamControllerClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasEnqueue returns true if the method "ReadableByteStreamController.enqueue" exists.
func (this ReadableByteStreamController) HasEnqueue() bool {
	return js.True == bindings.HasReadableByteStreamControllerEnqueue(
		this.Ref(),
	)
}

// EnqueueFunc returns the method "ReadableByteStreamController.enqueue".
func (this ReadableByteStreamController) EnqueueFunc() (fn js.Func[func(chunk js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.ReadableByteStreamControllerEnqueueFunc(
			this.Ref(),
		),
	)
}

// Enqueue calls the method "ReadableByteStreamController.enqueue".
func (this ReadableByteStreamController) Enqueue(chunk js.ArrayBufferView) (ret js.Void) {
	bindings.CallReadableByteStreamControllerEnqueue(
		this.Ref(), js.Pointer(&ret),
		chunk.Ref(),
	)

	return
}

// TryEnqueue calls the method "ReadableByteStreamController.enqueue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableByteStreamController) TryEnqueue(chunk js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableByteStreamControllerEnqueue(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		chunk.Ref(),
	)

	return
}

// HasError returns true if the method "ReadableByteStreamController.error" exists.
func (this ReadableByteStreamController) HasError() bool {
	return js.True == bindings.HasReadableByteStreamControllerError(
		this.Ref(),
	)
}

// ErrorFunc returns the method "ReadableByteStreamController.error".
func (this ReadableByteStreamController) ErrorFunc() (fn js.Func[func(e js.Any)]) {
	return fn.FromRef(
		bindings.ReadableByteStreamControllerErrorFunc(
			this.Ref(),
		),
	)
}

// Error calls the method "ReadableByteStreamController.error".
func (this ReadableByteStreamController) Error(e js.Any) (ret js.Void) {
	bindings.CallReadableByteStreamControllerError(
		this.Ref(), js.Pointer(&ret),
		e.Ref(),
	)

	return
}

// TryError calls the method "ReadableByteStreamController.error"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableByteStreamController) TryError(e js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableByteStreamControllerError(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		e.Ref(),
	)

	return
}

// HasError1 returns true if the method "ReadableByteStreamController.error" exists.
func (this ReadableByteStreamController) HasError1() bool {
	return js.True == bindings.HasReadableByteStreamControllerError1(
		this.Ref(),
	)
}

// Error1Func returns the method "ReadableByteStreamController.error".
func (this ReadableByteStreamController) Error1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReadableByteStreamControllerError1Func(
			this.Ref(),
		),
	)
}

// Error1 calls the method "ReadableByteStreamController.error".
func (this ReadableByteStreamController) Error1() (ret js.Void) {
	bindings.CallReadableByteStreamControllerError1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryError1 calls the method "ReadableByteStreamController.error"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableByteStreamController) TryError1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableByteStreamControllerError1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ReadableStreamDefaultController struct {
	ref js.Ref
}

func (this ReadableStreamDefaultController) Once() ReadableStreamDefaultController {
	this.Ref().Once()
	return this
}

func (this ReadableStreamDefaultController) Ref() js.Ref {
	return this.ref
}

func (this ReadableStreamDefaultController) FromRef(ref js.Ref) ReadableStreamDefaultController {
	this.ref = ref
	return this
}

func (this ReadableStreamDefaultController) Free() {
	this.Ref().Free()
}

// DesiredSize returns the value of property "ReadableStreamDefaultController.desiredSize".
//
// It returns ok=false if there is no such property.
func (this ReadableStreamDefaultController) DesiredSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetReadableStreamDefaultControllerDesiredSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasClose returns true if the method "ReadableStreamDefaultController.close" exists.
func (this ReadableStreamDefaultController) HasClose() bool {
	return js.True == bindings.HasReadableStreamDefaultControllerClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "ReadableStreamDefaultController.close".
func (this ReadableStreamDefaultController) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultControllerCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "ReadableStreamDefaultController.close".
func (this ReadableStreamDefaultController) Close() (ret js.Void) {
	bindings.CallReadableStreamDefaultControllerClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "ReadableStreamDefaultController.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamDefaultController) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamDefaultControllerClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasEnqueue returns true if the method "ReadableStreamDefaultController.enqueue" exists.
func (this ReadableStreamDefaultController) HasEnqueue() bool {
	return js.True == bindings.HasReadableStreamDefaultControllerEnqueue(
		this.Ref(),
	)
}

// EnqueueFunc returns the method "ReadableStreamDefaultController.enqueue".
func (this ReadableStreamDefaultController) EnqueueFunc() (fn js.Func[func(chunk js.Any)]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultControllerEnqueueFunc(
			this.Ref(),
		),
	)
}

// Enqueue calls the method "ReadableStreamDefaultController.enqueue".
func (this ReadableStreamDefaultController) Enqueue(chunk js.Any) (ret js.Void) {
	bindings.CallReadableStreamDefaultControllerEnqueue(
		this.Ref(), js.Pointer(&ret),
		chunk.Ref(),
	)

	return
}

// TryEnqueue calls the method "ReadableStreamDefaultController.enqueue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamDefaultController) TryEnqueue(chunk js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamDefaultControllerEnqueue(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		chunk.Ref(),
	)

	return
}

// HasEnqueue1 returns true if the method "ReadableStreamDefaultController.enqueue" exists.
func (this ReadableStreamDefaultController) HasEnqueue1() bool {
	return js.True == bindings.HasReadableStreamDefaultControllerEnqueue1(
		this.Ref(),
	)
}

// Enqueue1Func returns the method "ReadableStreamDefaultController.enqueue".
func (this ReadableStreamDefaultController) Enqueue1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultControllerEnqueue1Func(
			this.Ref(),
		),
	)
}

// Enqueue1 calls the method "ReadableStreamDefaultController.enqueue".
func (this ReadableStreamDefaultController) Enqueue1() (ret js.Void) {
	bindings.CallReadableStreamDefaultControllerEnqueue1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryEnqueue1 calls the method "ReadableStreamDefaultController.enqueue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamDefaultController) TryEnqueue1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamDefaultControllerEnqueue1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasError returns true if the method "ReadableStreamDefaultController.error" exists.
func (this ReadableStreamDefaultController) HasError() bool {
	return js.True == bindings.HasReadableStreamDefaultControllerError(
		this.Ref(),
	)
}

// ErrorFunc returns the method "ReadableStreamDefaultController.error".
func (this ReadableStreamDefaultController) ErrorFunc() (fn js.Func[func(e js.Any)]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultControllerErrorFunc(
			this.Ref(),
		),
	)
}

// Error calls the method "ReadableStreamDefaultController.error".
func (this ReadableStreamDefaultController) Error(e js.Any) (ret js.Void) {
	bindings.CallReadableStreamDefaultControllerError(
		this.Ref(), js.Pointer(&ret),
		e.Ref(),
	)

	return
}

// TryError calls the method "ReadableStreamDefaultController.error"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamDefaultController) TryError(e js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamDefaultControllerError(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		e.Ref(),
	)

	return
}

// HasError1 returns true if the method "ReadableStreamDefaultController.error" exists.
func (this ReadableStreamDefaultController) HasError1() bool {
	return js.True == bindings.HasReadableStreamDefaultControllerError1(
		this.Ref(),
	)
}

// Error1Func returns the method "ReadableStreamDefaultController.error".
func (this ReadableStreamDefaultController) Error1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultControllerError1Func(
			this.Ref(),
		),
	)
}

// Error1 calls the method "ReadableStreamDefaultController.error".
func (this ReadableStreamDefaultController) Error1() (ret js.Void) {
	bindings.CallReadableStreamDefaultControllerError1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryError1 calls the method "ReadableStreamDefaultController.error"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamDefaultController) TryError1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamDefaultControllerError1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OneOf_ReadableStreamDefaultController_ReadableByteStreamController struct {
	ref js.Ref
}

func (x OneOf_ReadableStreamDefaultController_ReadableByteStreamController) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ReadableStreamDefaultController_ReadableByteStreamController) Free() {
	x.ref.Free()
}

func (x OneOf_ReadableStreamDefaultController_ReadableByteStreamController) FromRef(ref js.Ref) OneOf_ReadableStreamDefaultController_ReadableByteStreamController {
	return OneOf_ReadableStreamDefaultController_ReadableByteStreamController{
		ref: ref,
	}
}

func (x OneOf_ReadableStreamDefaultController_ReadableByteStreamController) ReadableStreamDefaultController() ReadableStreamDefaultController {
	return ReadableStreamDefaultController{}.FromRef(x.ref)
}

func (x OneOf_ReadableStreamDefaultController_ReadableByteStreamController) ReadableByteStreamController() ReadableByteStreamController {
	return ReadableByteStreamController{}.FromRef(x.ref)
}

type ReadableStreamController = OneOf_ReadableStreamDefaultController_ReadableByteStreamController

type ReadableStreamIteratorOptions struct {
	// PreventCancel is "ReadableStreamIteratorOptions.preventCancel"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_PreventCancel MUST be set to true to make this field effective.
	PreventCancel bool

	FFI_USE_PreventCancel bool // for PreventCancel.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReadableStreamIteratorOptions with all fields set.
func (p ReadableStreamIteratorOptions) FromRef(ref js.Ref) ReadableStreamIteratorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReadableStreamIteratorOptions in the application heap.
func (p ReadableStreamIteratorOptions) New() js.Ref {
	return bindings.ReadableStreamIteratorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ReadableStreamIteratorOptions) UpdateFrom(ref js.Ref) {
	bindings.ReadableStreamIteratorOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ReadableStreamIteratorOptions) Update(ref js.Ref) {
	bindings.ReadableStreamIteratorOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ReadableStreamType uint32

const (
	_ ReadableStreamType = iota

	ReadableStreamType_BYTES
)

func (ReadableStreamType) FromRef(str js.Ref) ReadableStreamType {
	return ReadableStreamType(bindings.ConstOfReadableStreamType(str))
}

func (x ReadableStreamType) String() (string, bool) {
	switch x {
	case ReadableStreamType_BYTES:
		return "bytes", true
	default:
		return "", false
	}
}

type RelativeOrientationReadingValues struct {
	// Quaternion is "RelativeOrientationReadingValues.quaternion"
	//
	// Required
	Quaternion js.FrozenArray[float64]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RelativeOrientationReadingValues with all fields set.
func (p RelativeOrientationReadingValues) FromRef(ref js.Ref) RelativeOrientationReadingValues {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RelativeOrientationReadingValues in the application heap.
func (p RelativeOrientationReadingValues) New() js.Ref {
	return bindings.RelativeOrientationReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RelativeOrientationReadingValues) UpdateFrom(ref js.Ref) {
	bindings.RelativeOrientationReadingValuesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RelativeOrientationReadingValues) Update(ref js.Ref) {
	bindings.RelativeOrientationReadingValuesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewRelativeOrientationSensor(sensorOptions OrientationSensorOptions) (ret RelativeOrientationSensor) {
	ret.ref = bindings.NewRelativeOrientationSensorByRelativeOrientationSensor(
		js.Pointer(&sensorOptions))
	return
}

func NewRelativeOrientationSensorByRelativeOrientationSensor1() (ret RelativeOrientationSensor) {
	ret.ref = bindings.NewRelativeOrientationSensorByRelativeOrientationSensor1()
	return
}

type RelativeOrientationSensor struct {
	OrientationSensor
}

func (this RelativeOrientationSensor) Once() RelativeOrientationSensor {
	this.Ref().Once()
	return this
}

func (this RelativeOrientationSensor) Ref() js.Ref {
	return this.OrientationSensor.Ref()
}

func (this RelativeOrientationSensor) FromRef(ref js.Ref) RelativeOrientationSensor {
	this.OrientationSensor = this.OrientationSensor.FromRef(ref)
	return this
}

func (this RelativeOrientationSensor) Free() {
	this.Ref().Free()
}

type ReportBody struct {
	ref js.Ref
}

func (this ReportBody) Once() ReportBody {
	this.Ref().Once()
	return this
}

func (this ReportBody) Ref() js.Ref {
	return this.ref
}

func (this ReportBody) FromRef(ref js.Ref) ReportBody {
	this.ref = ref
	return this
}

func (this ReportBody) Free() {
	this.Ref().Free()
}

// HasToJSON returns true if the method "ReportBody.toJSON" exists.
func (this ReportBody) HasToJSON() bool {
	return js.True == bindings.HasReportBodyToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "ReportBody.toJSON".
func (this ReportBody) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.ReportBodyToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "ReportBody.toJSON".
func (this ReportBody) ToJSON() (ret js.Object) {
	bindings.CallReportBodyToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "ReportBody.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReportBody) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportBodyToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type Report struct {
	ref js.Ref
}

func (this Report) Once() Report {
	this.Ref().Once()
	return this
}

func (this Report) Ref() js.Ref {
	return this.ref
}

func (this Report) FromRef(ref js.Ref) Report {
	this.ref = ref
	return this
}

func (this Report) Free() {
	this.Ref().Free()
}

// Type returns the value of property "Report.type".
//
// It returns ok=false if there is no such property.
func (this Report) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetReportType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Url returns the value of property "Report.url".
//
// It returns ok=false if there is no such property.
func (this Report) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetReportUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Body returns the value of property "Report.body".
//
// It returns ok=false if there is no such property.
func (this Report) Body() (ret ReportBody, ok bool) {
	ok = js.True == bindings.GetReportBody(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "Report.toJSON" exists.
func (this Report) HasToJSON() bool {
	return js.True == bindings.HasReportToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "Report.toJSON".
func (this Report) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.ReportToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "Report.toJSON".
func (this Report) ToJSON() (ret js.Object) {
	bindings.CallReportToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "Report.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Report) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ReportList = js.Array[Report]

type ReportResultBrowserSignals struct {
	// Desirability is "ReportResultBrowserSignals.desirability"
	//
	// Required
	Desirability float64
	// TopLevelSellerSignals is "ReportResultBrowserSignals.topLevelSellerSignals"
	//
	// Optional
	TopLevelSellerSignals js.String
	// ModifiedBid is "ReportResultBrowserSignals.modifiedBid"
	//
	// Optional
	//
	// NOTE: FFI_USE_ModifiedBid MUST be set to true to make this field effective.
	ModifiedBid float64
	// DataVersion is "ReportResultBrowserSignals.dataVersion"
	//
	// Optional
	//
	// NOTE: FFI_USE_DataVersion MUST be set to true to make this field effective.
	DataVersion uint32
	// TopWindowHostname is "ReportResultBrowserSignals.topWindowHostname"
	//
	// Required
	TopWindowHostname js.String
	// InterestGroupOwner is "ReportResultBrowserSignals.interestGroupOwner"
	//
	// Required
	InterestGroupOwner js.String
	// RenderURL is "ReportResultBrowserSignals.renderURL"
	//
	// Required
	RenderURL js.String
	// Bid is "ReportResultBrowserSignals.bid"
	//
	// Required
	Bid float64
	// HighestScoringOtherBid is "ReportResultBrowserSignals.highestScoringOtherBid"
	//
	// Required
	HighestScoringOtherBid float64
	// BidCurrency is "ReportResultBrowserSignals.bidCurrency"
	//
	// Optional
	BidCurrency js.String
	// HighestScoringOtherBidCurrency is "ReportResultBrowserSignals.highestScoringOtherBidCurrency"
	//
	// Optional
	HighestScoringOtherBidCurrency js.String
	// TopLevelSeller is "ReportResultBrowserSignals.topLevelSeller"
	//
	// Optional
	TopLevelSeller js.String
	// ComponentSeller is "ReportResultBrowserSignals.componentSeller"
	//
	// Optional
	ComponentSeller js.String
	// BuyerAndSellerReportingId is "ReportResultBrowserSignals.buyerAndSellerReportingId"
	//
	// Optional
	BuyerAndSellerReportingId js.String

	FFI_USE_ModifiedBid bool // for ModifiedBid.
	FFI_USE_DataVersion bool // for DataVersion.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReportResultBrowserSignals with all fields set.
func (p ReportResultBrowserSignals) FromRef(ref js.Ref) ReportResultBrowserSignals {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReportResultBrowserSignals in the application heap.
func (p ReportResultBrowserSignals) New() js.Ref {
	return bindings.ReportResultBrowserSignalsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ReportResultBrowserSignals) UpdateFrom(ref js.Ref) {
	bindings.ReportResultBrowserSignalsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ReportResultBrowserSignals) Update(ref js.Ref) {
	bindings.ReportResultBrowserSignalsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ReportWinBrowserSignals struct {
	// AdCost is "ReportWinBrowserSignals.adCost"
	//
	// Optional
	//
	// NOTE: FFI_USE_AdCost MUST be set to true to make this field effective.
	AdCost float64
	// Seller is "ReportWinBrowserSignals.seller"
	//
	// Optional
	Seller js.String
	// MadeHighestScoringOtherBid is "ReportWinBrowserSignals.madeHighestScoringOtherBid"
	//
	// Optional
	//
	// NOTE: FFI_USE_MadeHighestScoringOtherBid MUST be set to true to make this field effective.
	MadeHighestScoringOtherBid bool
	// InterestGroupName is "ReportWinBrowserSignals.interestGroupName"
	//
	// Optional
	InterestGroupName js.String
	// BuyerReportingId is "ReportWinBrowserSignals.buyerReportingId"
	//
	// Optional
	BuyerReportingId js.String
	// ModelingSignals is "ReportWinBrowserSignals.modelingSignals"
	//
	// Optional
	//
	// NOTE: FFI_USE_ModelingSignals MUST be set to true to make this field effective.
	ModelingSignals uint16
	// DataVersion is "ReportWinBrowserSignals.dataVersion"
	//
	// Optional
	//
	// NOTE: FFI_USE_DataVersion MUST be set to true to make this field effective.
	DataVersion uint32
	// TopWindowHostname is "ReportWinBrowserSignals.topWindowHostname"
	//
	// Required
	TopWindowHostname js.String
	// InterestGroupOwner is "ReportWinBrowserSignals.interestGroupOwner"
	//
	// Required
	InterestGroupOwner js.String
	// RenderURL is "ReportWinBrowserSignals.renderURL"
	//
	// Required
	RenderURL js.String
	// Bid is "ReportWinBrowserSignals.bid"
	//
	// Required
	Bid float64
	// HighestScoringOtherBid is "ReportWinBrowserSignals.highestScoringOtherBid"
	//
	// Required
	HighestScoringOtherBid float64
	// BidCurrency is "ReportWinBrowserSignals.bidCurrency"
	//
	// Optional
	BidCurrency js.String
	// HighestScoringOtherBidCurrency is "ReportWinBrowserSignals.highestScoringOtherBidCurrency"
	//
	// Optional
	HighestScoringOtherBidCurrency js.String
	// TopLevelSeller is "ReportWinBrowserSignals.topLevelSeller"
	//
	// Optional
	TopLevelSeller js.String
	// ComponentSeller is "ReportWinBrowserSignals.componentSeller"
	//
	// Optional
	ComponentSeller js.String
	// BuyerAndSellerReportingId is "ReportWinBrowserSignals.buyerAndSellerReportingId"
	//
	// Optional
	BuyerAndSellerReportingId js.String

	FFI_USE_AdCost                     bool // for AdCost.
	FFI_USE_MadeHighestScoringOtherBid bool // for MadeHighestScoringOtherBid.
	FFI_USE_ModelingSignals            bool // for ModelingSignals.
	FFI_USE_DataVersion                bool // for DataVersion.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReportWinBrowserSignals with all fields set.
func (p ReportWinBrowserSignals) FromRef(ref js.Ref) ReportWinBrowserSignals {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReportWinBrowserSignals in the application heap.
func (p ReportWinBrowserSignals) New() js.Ref {
	return bindings.ReportWinBrowserSignalsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ReportWinBrowserSignals) UpdateFrom(ref js.Ref) {
	bindings.ReportWinBrowserSignalsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ReportWinBrowserSignals) Update(ref js.Ref) {
	bindings.ReportWinBrowserSignalsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ReportingBrowserSignals struct {
	// TopWindowHostname is "ReportingBrowserSignals.topWindowHostname"
	//
	// Required
	TopWindowHostname js.String
	// InterestGroupOwner is "ReportingBrowserSignals.interestGroupOwner"
	//
	// Required
	InterestGroupOwner js.String
	// RenderURL is "ReportingBrowserSignals.renderURL"
	//
	// Required
	RenderURL js.String
	// Bid is "ReportingBrowserSignals.bid"
	//
	// Required
	Bid float64
	// HighestScoringOtherBid is "ReportingBrowserSignals.highestScoringOtherBid"
	//
	// Required
	HighestScoringOtherBid float64
	// BidCurrency is "ReportingBrowserSignals.bidCurrency"
	//
	// Optional
	BidCurrency js.String
	// HighestScoringOtherBidCurrency is "ReportingBrowserSignals.highestScoringOtherBidCurrency"
	//
	// Optional
	HighestScoringOtherBidCurrency js.String
	// TopLevelSeller is "ReportingBrowserSignals.topLevelSeller"
	//
	// Optional
	TopLevelSeller js.String
	// ComponentSeller is "ReportingBrowserSignals.componentSeller"
	//
	// Optional
	ComponentSeller js.String
	// BuyerAndSellerReportingId is "ReportingBrowserSignals.buyerAndSellerReportingId"
	//
	// Optional
	BuyerAndSellerReportingId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReportingBrowserSignals with all fields set.
func (p ReportingBrowserSignals) FromRef(ref js.Ref) ReportingBrowserSignals {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReportingBrowserSignals in the application heap.
func (p ReportingBrowserSignals) New() js.Ref {
	return bindings.ReportingBrowserSignalsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ReportingBrowserSignals) UpdateFrom(ref js.Ref) {
	bindings.ReportingBrowserSignalsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ReportingBrowserSignals) Update(ref js.Ref) {
	bindings.ReportingBrowserSignalsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ReportingObserverCallbackFunc func(this js.Ref, reports js.Array[Report], observer ReportingObserver) js.Ref

func (fn ReportingObserverCallbackFunc) Register() js.Func[func(reports js.Array[Report], observer ReportingObserver)] {
	return js.RegisterCallback[func(reports js.Array[Report], observer ReportingObserver)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ReportingObserverCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Report]{}.FromRef(args[0+1]),
		ReportingObserver{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ReportingObserverCallback[T any] struct {
	Fn  func(arg T, this js.Ref, reports js.Array[Report], observer ReportingObserver) js.Ref
	Arg T
}

func (cb *ReportingObserverCallback[T]) Register() js.Func[func(reports js.Array[Report], observer ReportingObserver)] {
	return js.RegisterCallback[func(reports js.Array[Report], observer ReportingObserver)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ReportingObserverCallback[T]) DispatchCallback(
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

		js.Array[Report]{}.FromRef(args[0+1]),
		ReportingObserver{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ReportingObserverOptions struct {
	// Types is "ReportingObserverOptions.types"
	//
	// Optional
	Types js.Array[js.String]
	// Buffered is "ReportingObserverOptions.buffered"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Buffered MUST be set to true to make this field effective.
	Buffered bool

	FFI_USE_Buffered bool // for Buffered.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReportingObserverOptions with all fields set.
func (p ReportingObserverOptions) FromRef(ref js.Ref) ReportingObserverOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReportingObserverOptions in the application heap.
func (p ReportingObserverOptions) New() js.Ref {
	return bindings.ReportingObserverOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ReportingObserverOptions) UpdateFrom(ref js.Ref) {
	bindings.ReportingObserverOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ReportingObserverOptions) Update(ref js.Ref) {
	bindings.ReportingObserverOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewReportingObserver(callback js.Func[func(reports js.Array[Report], observer ReportingObserver)], options ReportingObserverOptions) (ret ReportingObserver) {
	ret.ref = bindings.NewReportingObserverByReportingObserver(
		callback.Ref(),
		js.Pointer(&options))
	return
}

func NewReportingObserverByReportingObserver1(callback js.Func[func(reports js.Array[Report], observer ReportingObserver)]) (ret ReportingObserver) {
	ret.ref = bindings.NewReportingObserverByReportingObserver1(
		callback.Ref())
	return
}

type ReportingObserver struct {
	ref js.Ref
}

func (this ReportingObserver) Once() ReportingObserver {
	this.Ref().Once()
	return this
}

func (this ReportingObserver) Ref() js.Ref {
	return this.ref
}

func (this ReportingObserver) FromRef(ref js.Ref) ReportingObserver {
	this.ref = ref
	return this
}

func (this ReportingObserver) Free() {
	this.Ref().Free()
}

// HasObserve returns true if the method "ReportingObserver.observe" exists.
func (this ReportingObserver) HasObserve() bool {
	return js.True == bindings.HasReportingObserverObserve(
		this.Ref(),
	)
}

// ObserveFunc returns the method "ReportingObserver.observe".
func (this ReportingObserver) ObserveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReportingObserverObserveFunc(
			this.Ref(),
		),
	)
}

// Observe calls the method "ReportingObserver.observe".
func (this ReportingObserver) Observe() (ret js.Void) {
	bindings.CallReportingObserverObserve(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryObserve calls the method "ReportingObserver.observe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReportingObserver) TryObserve() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportingObserverObserve(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDisconnect returns true if the method "ReportingObserver.disconnect" exists.
func (this ReportingObserver) HasDisconnect() bool {
	return js.True == bindings.HasReportingObserverDisconnect(
		this.Ref(),
	)
}

// DisconnectFunc returns the method "ReportingObserver.disconnect".
func (this ReportingObserver) DisconnectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReportingObserverDisconnectFunc(
			this.Ref(),
		),
	)
}

// Disconnect calls the method "ReportingObserver.disconnect".
func (this ReportingObserver) Disconnect() (ret js.Void) {
	bindings.CallReportingObserverDisconnect(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDisconnect calls the method "ReportingObserver.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReportingObserver) TryDisconnect() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportingObserverDisconnect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasTakeRecords returns true if the method "ReportingObserver.takeRecords" exists.
func (this ReportingObserver) HasTakeRecords() bool {
	return js.True == bindings.HasReportingObserverTakeRecords(
		this.Ref(),
	)
}

// TakeRecordsFunc returns the method "ReportingObserver.takeRecords".
func (this ReportingObserver) TakeRecordsFunc() (fn js.Func[func() ReportList]) {
	return fn.FromRef(
		bindings.ReportingObserverTakeRecordsFunc(
			this.Ref(),
		),
	)
}

// TakeRecords calls the method "ReportingObserver.takeRecords".
func (this ReportingObserver) TakeRecords() (ret ReportList) {
	bindings.CallReportingObserverTakeRecords(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTakeRecords calls the method "ReportingObserver.takeRecords"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReportingObserver) TryTakeRecords() (ret ReportList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportingObserverTakeRecords(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ResidentKeyRequirement uint32

const (
	_ ResidentKeyRequirement = iota

	ResidentKeyRequirement_DISCOURAGED
	ResidentKeyRequirement_PREFERRED
	ResidentKeyRequirement_REQUIRED
)

func (ResidentKeyRequirement) FromRef(str js.Ref) ResidentKeyRequirement {
	return ResidentKeyRequirement(bindings.ConstOfResidentKeyRequirement(str))
}

func (x ResidentKeyRequirement) String() (string, bool) {
	switch x {
	case ResidentKeyRequirement_DISCOURAGED:
		return "discouraged", true
	case ResidentKeyRequirement_PREFERRED:
		return "preferred", true
	case ResidentKeyRequirement_REQUIRED:
		return "required", true
	default:
		return "", false
	}
}

type ResizeObserverCallbackFunc func(this js.Ref, entries js.Array[ResizeObserverEntry], observer ResizeObserver) js.Ref

func (fn ResizeObserverCallbackFunc) Register() js.Func[func(entries js.Array[ResizeObserverEntry], observer ResizeObserver)] {
	return js.RegisterCallback[func(entries js.Array[ResizeObserverEntry], observer ResizeObserver)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ResizeObserverCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ResizeObserverEntry]{}.FromRef(args[0+1]),
		ResizeObserver{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ResizeObserverCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[ResizeObserverEntry], observer ResizeObserver) js.Ref
	Arg T
}

func (cb *ResizeObserverCallback[T]) Register() js.Func[func(entries js.Array[ResizeObserverEntry], observer ResizeObserver)] {
	return js.RegisterCallback[func(entries js.Array[ResizeObserverEntry], observer ResizeObserver)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ResizeObserverCallback[T]) DispatchCallback(
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

		js.Array[ResizeObserverEntry]{}.FromRef(args[0+1]),
		ResizeObserver{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ResizeObserverSize struct {
	ref js.Ref
}

func (this ResizeObserverSize) Once() ResizeObserverSize {
	this.Ref().Once()
	return this
}

func (this ResizeObserverSize) Ref() js.Ref {
	return this.ref
}

func (this ResizeObserverSize) FromRef(ref js.Ref) ResizeObserverSize {
	this.ref = ref
	return this
}

func (this ResizeObserverSize) Free() {
	this.Ref().Free()
}

// InlineSize returns the value of property "ResizeObserverSize.inlineSize".
//
// It returns ok=false if there is no such property.
func (this ResizeObserverSize) InlineSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetResizeObserverSizeInlineSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BlockSize returns the value of property "ResizeObserverSize.blockSize".
//
// It returns ok=false if there is no such property.
func (this ResizeObserverSize) BlockSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetResizeObserverSizeBlockSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type ResizeObserverEntry struct {
	ref js.Ref
}

func (this ResizeObserverEntry) Once() ResizeObserverEntry {
	this.Ref().Once()
	return this
}

func (this ResizeObserverEntry) Ref() js.Ref {
	return this.ref
}

func (this ResizeObserverEntry) FromRef(ref js.Ref) ResizeObserverEntry {
	this.ref = ref
	return this
}

func (this ResizeObserverEntry) Free() {
	this.Ref().Free()
}

// Target returns the value of property "ResizeObserverEntry.target".
//
// It returns ok=false if there is no such property.
func (this ResizeObserverEntry) Target() (ret Element, ok bool) {
	ok = js.True == bindings.GetResizeObserverEntryTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContentRect returns the value of property "ResizeObserverEntry.contentRect".
//
// It returns ok=false if there is no such property.
func (this ResizeObserverEntry) ContentRect() (ret DOMRectReadOnly, ok bool) {
	ok = js.True == bindings.GetResizeObserverEntryContentRect(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BorderBoxSize returns the value of property "ResizeObserverEntry.borderBoxSize".
//
// It returns ok=false if there is no such property.
func (this ResizeObserverEntry) BorderBoxSize() (ret js.FrozenArray[ResizeObserverSize], ok bool) {
	ok = js.True == bindings.GetResizeObserverEntryBorderBoxSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContentBoxSize returns the value of property "ResizeObserverEntry.contentBoxSize".
//
// It returns ok=false if there is no such property.
func (this ResizeObserverEntry) ContentBoxSize() (ret js.FrozenArray[ResizeObserverSize], ok bool) {
	ok = js.True == bindings.GetResizeObserverEntryContentBoxSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DevicePixelContentBoxSize returns the value of property "ResizeObserverEntry.devicePixelContentBoxSize".
//
// It returns ok=false if there is no such property.
func (this ResizeObserverEntry) DevicePixelContentBoxSize() (ret js.FrozenArray[ResizeObserverSize], ok bool) {
	ok = js.True == bindings.GetResizeObserverEntryDevicePixelContentBoxSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type ResizeObserverBoxOptions uint32

const (
	_ ResizeObserverBoxOptions = iota

	ResizeObserverBoxOptions_BORDER_BOX
	ResizeObserverBoxOptions_CONTENT_BOX
	ResizeObserverBoxOptions_DEVICE_PIXEL_CONTENT_BOX
)

func (ResizeObserverBoxOptions) FromRef(str js.Ref) ResizeObserverBoxOptions {
	return ResizeObserverBoxOptions(bindings.ConstOfResizeObserverBoxOptions(str))
}

func (x ResizeObserverBoxOptions) String() (string, bool) {
	switch x {
	case ResizeObserverBoxOptions_BORDER_BOX:
		return "border-box", true
	case ResizeObserverBoxOptions_CONTENT_BOX:
		return "content-box", true
	case ResizeObserverBoxOptions_DEVICE_PIXEL_CONTENT_BOX:
		return "device-pixel-content-box", true
	default:
		return "", false
	}
}

type ResizeObserverOptions struct {
	// Box is "ResizeObserverOptions.box"
	//
	// Optional, defaults to "content-box".
	Box ResizeObserverBoxOptions

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ResizeObserverOptions with all fields set.
func (p ResizeObserverOptions) FromRef(ref js.Ref) ResizeObserverOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ResizeObserverOptions in the application heap.
func (p ResizeObserverOptions) New() js.Ref {
	return bindings.ResizeObserverOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ResizeObserverOptions) UpdateFrom(ref js.Ref) {
	bindings.ResizeObserverOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ResizeObserverOptions) Update(ref js.Ref) {
	bindings.ResizeObserverOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewResizeObserver(callback js.Func[func(entries js.Array[ResizeObserverEntry], observer ResizeObserver)]) (ret ResizeObserver) {
	ret.ref = bindings.NewResizeObserverByResizeObserver(
		callback.Ref())
	return
}

type ResizeObserver struct {
	ref js.Ref
}

func (this ResizeObserver) Once() ResizeObserver {
	this.Ref().Once()
	return this
}

func (this ResizeObserver) Ref() js.Ref {
	return this.ref
}

func (this ResizeObserver) FromRef(ref js.Ref) ResizeObserver {
	this.ref = ref
	return this
}

func (this ResizeObserver) Free() {
	this.Ref().Free()
}

// HasObserve returns true if the method "ResizeObserver.observe" exists.
func (this ResizeObserver) HasObserve() bool {
	return js.True == bindings.HasResizeObserverObserve(
		this.Ref(),
	)
}

// ObserveFunc returns the method "ResizeObserver.observe".
func (this ResizeObserver) ObserveFunc() (fn js.Func[func(target Element, options ResizeObserverOptions)]) {
	return fn.FromRef(
		bindings.ResizeObserverObserveFunc(
			this.Ref(),
		),
	)
}

// Observe calls the method "ResizeObserver.observe".
func (this ResizeObserver) Observe(target Element, options ResizeObserverOptions) (ret js.Void) {
	bindings.CallResizeObserverObserve(
		this.Ref(), js.Pointer(&ret),
		target.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryObserve calls the method "ResizeObserver.observe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ResizeObserver) TryObserve(target Element, options ResizeObserverOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResizeObserverObserve(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		target.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasObserve1 returns true if the method "ResizeObserver.observe" exists.
func (this ResizeObserver) HasObserve1() bool {
	return js.True == bindings.HasResizeObserverObserve1(
		this.Ref(),
	)
}

// Observe1Func returns the method "ResizeObserver.observe".
func (this ResizeObserver) Observe1Func() (fn js.Func[func(target Element)]) {
	return fn.FromRef(
		bindings.ResizeObserverObserve1Func(
			this.Ref(),
		),
	)
}

// Observe1 calls the method "ResizeObserver.observe".
func (this ResizeObserver) Observe1(target Element) (ret js.Void) {
	bindings.CallResizeObserverObserve1(
		this.Ref(), js.Pointer(&ret),
		target.Ref(),
	)

	return
}

// TryObserve1 calls the method "ResizeObserver.observe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ResizeObserver) TryObserve1(target Element) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResizeObserverObserve1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		target.Ref(),
	)

	return
}

// HasUnobserve returns true if the method "ResizeObserver.unobserve" exists.
func (this ResizeObserver) HasUnobserve() bool {
	return js.True == bindings.HasResizeObserverUnobserve(
		this.Ref(),
	)
}

// UnobserveFunc returns the method "ResizeObserver.unobserve".
func (this ResizeObserver) UnobserveFunc() (fn js.Func[func(target Element)]) {
	return fn.FromRef(
		bindings.ResizeObserverUnobserveFunc(
			this.Ref(),
		),
	)
}

// Unobserve calls the method "ResizeObserver.unobserve".
func (this ResizeObserver) Unobserve(target Element) (ret js.Void) {
	bindings.CallResizeObserverUnobserve(
		this.Ref(), js.Pointer(&ret),
		target.Ref(),
	)

	return
}

// TryUnobserve calls the method "ResizeObserver.unobserve"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ResizeObserver) TryUnobserve(target Element) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResizeObserverUnobserve(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		target.Ref(),
	)

	return
}

// HasDisconnect returns true if the method "ResizeObserver.disconnect" exists.
func (this ResizeObserver) HasDisconnect() bool {
	return js.True == bindings.HasResizeObserverDisconnect(
		this.Ref(),
	)
}

// DisconnectFunc returns the method "ResizeObserver.disconnect".
func (this ResizeObserver) DisconnectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ResizeObserverDisconnectFunc(
			this.Ref(),
		),
	)
}

// Disconnect calls the method "ResizeObserver.disconnect".
func (this ResizeObserver) Disconnect() (ret js.Void) {
	bindings.CallResizeObserverDisconnect(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDisconnect calls the method "ResizeObserver.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ResizeObserver) TryDisconnect() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResizeObserverDisconnect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type RsaHashedImportParams struct {
	// Hash is "RsaHashedImportParams.hash"
	//
	// Required
	Hash HashAlgorithmIdentifier
	// Name is "RsaHashedImportParams.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RsaHashedImportParams with all fields set.
func (p RsaHashedImportParams) FromRef(ref js.Ref) RsaHashedImportParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RsaHashedImportParams in the application heap.
func (p RsaHashedImportParams) New() js.Ref {
	return bindings.RsaHashedImportParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RsaHashedImportParams) UpdateFrom(ref js.Ref) {
	bindings.RsaHashedImportParamsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RsaHashedImportParams) Update(ref js.Ref) {
	bindings.RsaHashedImportParamsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RsaHashedKeyAlgorithm struct {
	// Hash is "RsaHashedKeyAlgorithm.hash"
	//
	// Required
	//
	// NOTE: Hash.FFI_USE MUST be set to true to get Hash used.
	Hash KeyAlgorithm
	// ModulusLength is "RsaHashedKeyAlgorithm.modulusLength"
	//
	// Required
	ModulusLength uint32
	// PublicExponent is "RsaHashedKeyAlgorithm.publicExponent"
	//
	// Required
	PublicExponent BigInteger
	// Name is "RsaHashedKeyAlgorithm.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RsaHashedKeyAlgorithm with all fields set.
func (p RsaHashedKeyAlgorithm) FromRef(ref js.Ref) RsaHashedKeyAlgorithm {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RsaHashedKeyAlgorithm in the application heap.
func (p RsaHashedKeyAlgorithm) New() js.Ref {
	return bindings.RsaHashedKeyAlgorithmJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RsaHashedKeyAlgorithm) UpdateFrom(ref js.Ref) {
	bindings.RsaHashedKeyAlgorithmJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RsaHashedKeyAlgorithm) Update(ref js.Ref) {
	bindings.RsaHashedKeyAlgorithmJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RsaHashedKeyGenParams struct {
	// Hash is "RsaHashedKeyGenParams.hash"
	//
	// Required
	Hash HashAlgorithmIdentifier
	// ModulusLength is "RsaHashedKeyGenParams.modulusLength"
	//
	// Required
	ModulusLength uint32
	// PublicExponent is "RsaHashedKeyGenParams.publicExponent"
	//
	// Required
	PublicExponent BigInteger
	// Name is "RsaHashedKeyGenParams.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RsaHashedKeyGenParams with all fields set.
func (p RsaHashedKeyGenParams) FromRef(ref js.Ref) RsaHashedKeyGenParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RsaHashedKeyGenParams in the application heap.
func (p RsaHashedKeyGenParams) New() js.Ref {
	return bindings.RsaHashedKeyGenParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RsaHashedKeyGenParams) UpdateFrom(ref js.Ref) {
	bindings.RsaHashedKeyGenParamsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RsaHashedKeyGenParams) Update(ref js.Ref) {
	bindings.RsaHashedKeyGenParamsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RsaKeyAlgorithm struct {
	// ModulusLength is "RsaKeyAlgorithm.modulusLength"
	//
	// Required
	ModulusLength uint32
	// PublicExponent is "RsaKeyAlgorithm.publicExponent"
	//
	// Required
	PublicExponent BigInteger
	// Name is "RsaKeyAlgorithm.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RsaKeyAlgorithm with all fields set.
func (p RsaKeyAlgorithm) FromRef(ref js.Ref) RsaKeyAlgorithm {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RsaKeyAlgorithm in the application heap.
func (p RsaKeyAlgorithm) New() js.Ref {
	return bindings.RsaKeyAlgorithmJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RsaKeyAlgorithm) UpdateFrom(ref js.Ref) {
	bindings.RsaKeyAlgorithmJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RsaKeyAlgorithm) Update(ref js.Ref) {
	bindings.RsaKeyAlgorithmJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RsaKeyGenParams struct {
	// ModulusLength is "RsaKeyGenParams.modulusLength"
	//
	// Required
	ModulusLength uint32
	// PublicExponent is "RsaKeyGenParams.publicExponent"
	//
	// Required
	PublicExponent BigInteger
	// Name is "RsaKeyGenParams.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RsaKeyGenParams with all fields set.
func (p RsaKeyGenParams) FromRef(ref js.Ref) RsaKeyGenParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RsaKeyGenParams in the application heap.
func (p RsaKeyGenParams) New() js.Ref {
	return bindings.RsaKeyGenParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RsaKeyGenParams) UpdateFrom(ref js.Ref) {
	bindings.RsaKeyGenParamsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RsaKeyGenParams) Update(ref js.Ref) {
	bindings.RsaKeyGenParamsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RsaOaepParams struct {
	// Label is "RsaOaepParams.label"
	//
	// Optional
	Label BufferSource
	// Name is "RsaOaepParams.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RsaOaepParams with all fields set.
func (p RsaOaepParams) FromRef(ref js.Ref) RsaOaepParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RsaOaepParams in the application heap.
func (p RsaOaepParams) New() js.Ref {
	return bindings.RsaOaepParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RsaOaepParams) UpdateFrom(ref js.Ref) {
	bindings.RsaOaepParamsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RsaOaepParams) Update(ref js.Ref) {
	bindings.RsaOaepParamsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RsaPssParams struct {
	// SaltLength is "RsaPssParams.saltLength"
	//
	// Required
	SaltLength uint32
	// Name is "RsaPssParams.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RsaPssParams with all fields set.
func (p RsaPssParams) FromRef(ref js.Ref) RsaPssParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RsaPssParams in the application heap.
func (p RsaPssParams) New() js.Ref {
	return bindings.RsaPssParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RsaPssParams) UpdateFrom(ref js.Ref) {
	bindings.RsaPssParamsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RsaPssParams) Update(ref js.Ref) {
	bindings.RsaPssParamsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type SFrameTransformErrorEventType uint32

const (
	_ SFrameTransformErrorEventType = iota

	SFrameTransformErrorEventType_AUTHENTICATION
	SFrameTransformErrorEventType_KEY_ID
	SFrameTransformErrorEventType_SYNTAX
)

func (SFrameTransformErrorEventType) FromRef(str js.Ref) SFrameTransformErrorEventType {
	return SFrameTransformErrorEventType(bindings.ConstOfSFrameTransformErrorEventType(str))
}

func (x SFrameTransformErrorEventType) String() (string, bool) {
	switch x {
	case SFrameTransformErrorEventType_AUTHENTICATION:
		return "authentication", true
	case SFrameTransformErrorEventType_KEY_ID:
		return "keyID", true
	case SFrameTransformErrorEventType_SYNTAX:
		return "syntax", true
	default:
		return "", false
	}
}

type SFrameTransformErrorEventInit struct {
	// ErrorType is "SFrameTransformErrorEventInit.errorType"
	//
	// Required
	ErrorType SFrameTransformErrorEventType
	// Frame is "SFrameTransformErrorEventInit.frame"
	//
	// Required
	Frame js.Any
	// KeyID is "SFrameTransformErrorEventInit.keyID"
	//
	// Optional
	KeyID CryptoKeyID
	// Bubbles is "SFrameTransformErrorEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "SFrameTransformErrorEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "SFrameTransformErrorEventInit.composed"
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

// FromRef calls UpdateFrom and returns a SFrameTransformErrorEventInit with all fields set.
func (p SFrameTransformErrorEventInit) FromRef(ref js.Ref) SFrameTransformErrorEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SFrameTransformErrorEventInit in the application heap.
func (p SFrameTransformErrorEventInit) New() js.Ref {
	return bindings.SFrameTransformErrorEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SFrameTransformErrorEventInit) UpdateFrom(ref js.Ref) {
	bindings.SFrameTransformErrorEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SFrameTransformErrorEventInit) Update(ref js.Ref) {
	bindings.SFrameTransformErrorEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewSFrameTransformErrorEvent(typ js.String, eventInitDict SFrameTransformErrorEventInit) (ret SFrameTransformErrorEvent) {
	ret.ref = bindings.NewSFrameTransformErrorEventBySFrameTransformErrorEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type SFrameTransformErrorEvent struct {
	Event
}

func (this SFrameTransformErrorEvent) Once() SFrameTransformErrorEvent {
	this.Ref().Once()
	return this
}

func (this SFrameTransformErrorEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this SFrameTransformErrorEvent) FromRef(ref js.Ref) SFrameTransformErrorEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this SFrameTransformErrorEvent) Free() {
	this.Ref().Free()
}

// ErrorType returns the value of property "SFrameTransformErrorEvent.errorType".
//
// It returns ok=false if there is no such property.
func (this SFrameTransformErrorEvent) ErrorType() (ret SFrameTransformErrorEventType, ok bool) {
	ok = js.True == bindings.GetSFrameTransformErrorEventErrorType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// KeyID returns the value of property "SFrameTransformErrorEvent.keyID".
//
// It returns ok=false if there is no such property.
func (this SFrameTransformErrorEvent) KeyID() (ret CryptoKeyID, ok bool) {
	ok = js.True == bindings.GetSFrameTransformErrorEventKeyID(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Frame returns the value of property "SFrameTransformErrorEvent.frame".
//
// It returns ok=false if there is no such property.
func (this SFrameTransformErrorEvent) Frame() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetSFrameTransformErrorEventFrame(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGAElement struct {
	SVGGraphicsElement
}

func (this SVGAElement) Once() SVGAElement {
	this.Ref().Once()
	return this
}

func (this SVGAElement) Ref() js.Ref {
	return this.SVGGraphicsElement.Ref()
}

func (this SVGAElement) FromRef(ref js.Ref) SVGAElement {
	this.SVGGraphicsElement = this.SVGGraphicsElement.FromRef(ref)
	return this
}

func (this SVGAElement) Free() {
	this.Ref().Free()
}

// Target returns the value of property "SVGAElement.target".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Target() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGAElementTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Download returns the value of property "SVGAElement.download".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Download() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementDownload(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDownload sets the value of property "SVGAElement.download" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetDownload(val js.String) bool {
	return js.True == bindings.SetSVGAElementDownload(
		this.Ref(),
		val.Ref(),
	)
}

// Ping returns the value of property "SVGAElement.ping".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Ping() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementPing(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPing sets the value of property "SVGAElement.ping" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetPing(val js.String) bool {
	return js.True == bindings.SetSVGAElementPing(
		this.Ref(),
		val.Ref(),
	)
}

// Rel returns the value of property "SVGAElement.rel".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Rel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementRel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetRel sets the value of property "SVGAElement.rel" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetRel(val js.String) bool {
	return js.True == bindings.SetSVGAElementRel(
		this.Ref(),
		val.Ref(),
	)
}

// RelList returns the value of property "SVGAElement.relList".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) RelList() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetSVGAElementRelList(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Hreflang returns the value of property "SVGAElement.hreflang".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Hreflang() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementHreflang(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHreflang sets the value of property "SVGAElement.hreflang" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetHreflang(val js.String) bool {
	return js.True == bindings.SetSVGAElementHreflang(
		this.Ref(),
		val.Ref(),
	)
}

// Type returns the value of property "SVGAElement.type".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "SVGAElement.type" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetType(val js.String) bool {
	return js.True == bindings.SetSVGAElementType(
		this.Ref(),
		val.Ref(),
	)
}

// Text returns the value of property "SVGAElement.text".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Text() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetText sets the value of property "SVGAElement.text" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetText(val js.String) bool {
	return js.True == bindings.SetSVGAElementText(
		this.Ref(),
		val.Ref(),
	)
}

// ReferrerPolicy returns the value of property "SVGAElement.referrerPolicy".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) ReferrerPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementReferrerPolicy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetReferrerPolicy sets the value of property "SVGAElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetSVGAElementReferrerPolicy(
		this.Ref(),
		val.Ref(),
	)
}

// Origin returns the value of property "SVGAElement.origin".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Protocol returns the value of property "SVGAElement.protocol".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Protocol() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementProtocol(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetProtocol sets the value of property "SVGAElement.protocol" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetProtocol(val js.String) bool {
	return js.True == bindings.SetSVGAElementProtocol(
		this.Ref(),
		val.Ref(),
	)
}

// Username returns the value of property "SVGAElement.username".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Username() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementUsername(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetUsername sets the value of property "SVGAElement.username" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetUsername(val js.String) bool {
	return js.True == bindings.SetSVGAElementUsername(
		this.Ref(),
		val.Ref(),
	)
}

// Password returns the value of property "SVGAElement.password".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Password() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementPassword(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPassword sets the value of property "SVGAElement.password" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetPassword(val js.String) bool {
	return js.True == bindings.SetSVGAElementPassword(
		this.Ref(),
		val.Ref(),
	)
}

// Host returns the value of property "SVGAElement.host".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Host() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementHost(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHost sets the value of property "SVGAElement.host" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetHost(val js.String) bool {
	return js.True == bindings.SetSVGAElementHost(
		this.Ref(),
		val.Ref(),
	)
}

// Hostname returns the value of property "SVGAElement.hostname".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Hostname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementHostname(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHostname sets the value of property "SVGAElement.hostname" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetHostname(val js.String) bool {
	return js.True == bindings.SetSVGAElementHostname(
		this.Ref(),
		val.Ref(),
	)
}

// Port returns the value of property "SVGAElement.port".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Port() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementPort(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPort sets the value of property "SVGAElement.port" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetPort(val js.String) bool {
	return js.True == bindings.SetSVGAElementPort(
		this.Ref(),
		val.Ref(),
	)
}

// Pathname returns the value of property "SVGAElement.pathname".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Pathname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementPathname(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPathname sets the value of property "SVGAElement.pathname" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetPathname(val js.String) bool {
	return js.True == bindings.SetSVGAElementPathname(
		this.Ref(),
		val.Ref(),
	)
}

// Search returns the value of property "SVGAElement.search".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Search() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementSearch(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSearch sets the value of property "SVGAElement.search" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetSearch(val js.String) bool {
	return js.True == bindings.SetSVGAElementSearch(
		this.Ref(),
		val.Ref(),
	)
}

// Hash returns the value of property "SVGAElement.hash".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Hash() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAElementHash(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHash sets the value of property "SVGAElement.hash" to val.
//
// It returns false if the property cannot be set.
func (this SVGAElement) SetHash(val js.String) bool {
	return js.True == bindings.SetSVGAElementHash(
		this.Ref(),
		val.Ref(),
	)
}

// Href returns the value of property "SVGAElement.href".
//
// It returns ok=false if there is no such property.
func (this SVGAElement) Href() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGAElementHref(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGAnimateElement struct {
	SVGAnimationElement
}

func (this SVGAnimateElement) Once() SVGAnimateElement {
	this.Ref().Once()
	return this
}

func (this SVGAnimateElement) Ref() js.Ref {
	return this.SVGAnimationElement.Ref()
}

func (this SVGAnimateElement) FromRef(ref js.Ref) SVGAnimateElement {
	this.SVGAnimationElement = this.SVGAnimationElement.FromRef(ref)
	return this
}

func (this SVGAnimateElement) Free() {
	this.Ref().Free()
}

type SVGAnimateMotionElement struct {
	SVGAnimationElement
}

func (this SVGAnimateMotionElement) Once() SVGAnimateMotionElement {
	this.Ref().Once()
	return this
}

func (this SVGAnimateMotionElement) Ref() js.Ref {
	return this.SVGAnimationElement.Ref()
}

func (this SVGAnimateMotionElement) FromRef(ref js.Ref) SVGAnimateMotionElement {
	this.SVGAnimationElement = this.SVGAnimationElement.FromRef(ref)
	return this
}

func (this SVGAnimateMotionElement) Free() {
	this.Ref().Free()
}

type SVGAnimateTransformElement struct {
	SVGAnimationElement
}

func (this SVGAnimateTransformElement) Once() SVGAnimateTransformElement {
	this.Ref().Once()
	return this
}

func (this SVGAnimateTransformElement) Ref() js.Ref {
	return this.SVGAnimationElement.Ref()
}

func (this SVGAnimateTransformElement) FromRef(ref js.Ref) SVGAnimateTransformElement {
	this.SVGAnimationElement = this.SVGAnimationElement.FromRef(ref)
	return this
}

func (this SVGAnimateTransformElement) Free() {
	this.Ref().Free()
}

type SVGAnimatedAngle struct {
	ref js.Ref
}

func (this SVGAnimatedAngle) Once() SVGAnimatedAngle {
	this.Ref().Once()
	return this
}

func (this SVGAnimatedAngle) Ref() js.Ref {
	return this.ref
}

func (this SVGAnimatedAngle) FromRef(ref js.Ref) SVGAnimatedAngle {
	this.ref = ref
	return this
}

func (this SVGAnimatedAngle) Free() {
	this.Ref().Free()
}

// BaseVal returns the value of property "SVGAnimatedAngle.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedAngle) BaseVal() (ret SVGAngle, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedAngleBaseVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AnimVal returns the value of property "SVGAnimatedAngle.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedAngle) AnimVal() (ret SVGAngle, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedAngleAnimVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGAnimatedBoolean struct {
	ref js.Ref
}

func (this SVGAnimatedBoolean) Once() SVGAnimatedBoolean {
	this.Ref().Once()
	return this
}

func (this SVGAnimatedBoolean) Ref() js.Ref {
	return this.ref
}

func (this SVGAnimatedBoolean) FromRef(ref js.Ref) SVGAnimatedBoolean {
	this.ref = ref
	return this
}

func (this SVGAnimatedBoolean) Free() {
	this.Ref().Free()
}

// BaseVal returns the value of property "SVGAnimatedBoolean.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedBoolean) BaseVal() (ret bool, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedBooleanBaseVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBaseVal sets the value of property "SVGAnimatedBoolean.baseVal" to val.
//
// It returns false if the property cannot be set.
func (this SVGAnimatedBoolean) SetBaseVal(val bool) bool {
	return js.True == bindings.SetSVGAnimatedBooleanBaseVal(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// AnimVal returns the value of property "SVGAnimatedBoolean.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedBoolean) AnimVal() (ret bool, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedBooleanAnimVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}
