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
	Bubbles bool
	// Cancelable is "RTCTrackEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "RTCTrackEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 RTCTrackEventInit RTCTrackEventInit [// RTCTrackEventInit] [0x140009eb180 0x140009eb220 0x140009eb2c0 0x140009eb360 0x140009eb400 0x140009eb540 0x140009eb680 0x140009eb4a0 0x140009eb5e0 0x140009eb720] 0x1400099d998 {0 0}} in the application heap.
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

func NewRTCTrackEvent(typ js.String, eventInitDict RTCTrackEventInit) RTCTrackEvent {
	return RTCTrackEvent{}.FromRef(
		bindings.NewRTCTrackEventByRTCTrackEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
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
// The returned bool will be false if there is no such property.
func (this RTCTrackEvent) Receiver() (RTCRtpReceiver, bool) {
	var _ok bool
	_ret := bindings.GetRTCTrackEventReceiver(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCRtpReceiver{}.FromRef(_ret), _ok
}

// Track returns the value of property "RTCTrackEvent.track".
//
// The returned bool will be false if there is no such property.
func (this RTCTrackEvent) Track() (MediaStreamTrack, bool) {
	var _ok bool
	_ret := bindings.GetRTCTrackEventTrack(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaStreamTrack{}.FromRef(_ret), _ok
}

// Streams returns the value of property "RTCTrackEvent.streams".
//
// The returned bool will be false if there is no such property.
func (this RTCTrackEvent) Streams() (js.FrozenArray[MediaStream], bool) {
	var _ok bool
	_ret := bindings.GetRTCTrackEventStreams(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[MediaStream]{}.FromRef(_ret), _ok
}

// Transceiver returns the value of property "RTCTrackEvent.transceiver".
//
// The returned bool will be false if there is no such property.
func (this RTCTrackEvent) Transceiver() (RTCRtpTransceiver, bool) {
	var _ok bool
	_ret := bindings.GetRTCTrackEventTransceiver(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCRtpTransceiver{}.FromRef(_ret), _ok
}

func NewRTCTransformEvent(typ js.String, eventInitDict EventInit) RTCTransformEvent {
	return RTCTransformEvent{}.FromRef(
		bindings.NewRTCTransformEventByRTCTransformEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewRTCTransformEventByRTCTransformEvent1(typ js.String) RTCTransformEvent {
	return RTCTransformEvent{}.FromRef(
		bindings.NewRTCTransformEventByRTCTransformEvent1(
			typ.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this RTCTransformEvent) Transformer() (RTCRtpScriptTransformer, bool) {
	var _ok bool
	_ret := bindings.GetRTCTransformEventTransformer(
		this.Ref(), js.Pointer(&_ok),
	)
	return RTCRtpScriptTransformer{}.FromRef(_ret), _ok
}

type RTCTransportStats struct {
	// PacketsSent is "RTCTransportStats.packetsSent"
	//
	// Optional
	PacketsSent uint64
	// PacketsReceived is "RTCTransportStats.packetsReceived"
	//
	// Optional
	PacketsReceived uint64
	// BytesSent is "RTCTransportStats.bytesSent"
	//
	// Optional
	BytesSent uint64
	// BytesReceived is "RTCTransportStats.bytesReceived"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 RTCTransportStats RTCTransportStats [// RTCTransportStats] [0x140009eb900 0x140009eba40 0x140009ebb80 0x140009ebcc0 0x140009ebe00 0x140009ebea0 0x140009ebf40 0x140009f0000 0x140009f00a0 0x140009f0140 0x140009f01e0 0x140009f0280 0x140009f0320 0x140009f03c0 0x140009f0460 0x140009f0500 0x140009f0640 0x140009f06e0 0x140009f0780 0x140009eb9a0 0x140009ebae0 0x140009ebc20 0x140009ebd60 0x140009f05a0] 0x1400099da10 {0 0}} in the application heap.
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
	Width uint32
	// Height is "RTCVideoSourceStats.height"
	//
	// Optional
	Height uint32
	// Frames is "RTCVideoSourceStats.frames"
	//
	// Optional
	Frames uint32
	// FramesPerSecond is "RTCVideoSourceStats.framesPerSecond"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 RTCVideoSourceStats RTCVideoSourceStats [// RTCVideoSourceStats] [0x140009f0820 0x140009f0960 0x140009f0aa0 0x140009f0be0 0x140009f0d20 0x140009f0dc0 0x140009f0e60 0x140009f0f00 0x140009f0fa0 0x140009f08c0 0x140009f0a00 0x140009f0b40 0x140009f0c80] 0x1400099dad0 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this ReadableStreamBYOBRequest) View() (ArrayBufferView, bool) {
	var _ok bool
	_ret := bindings.GetReadableStreamBYOBRequestView(
		this.Ref(), js.Pointer(&_ok),
	)
	return ArrayBufferView{}.FromRef(_ret), _ok
}

// Respond calls the method "ReadableStreamBYOBRequest.respond".
//
// The returned bool will be false if there is no such method.
func (this ReadableStreamBYOBRequest) Respond(bytesWritten uint64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamBYOBRequestRespond(
		this.Ref(), js.Pointer(&_ok),
		float64(bytesWritten),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RespondFunc returns the method "ReadableStreamBYOBRequest.respond".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStreamBYOBRequest) RespondFunc() (fn js.Func[func(bytesWritten uint64)]) {
	return fn.FromRef(
		bindings.ReadableStreamBYOBRequestRespondFunc(
			this.Ref(),
		),
	)
}

// RespondWithNewView calls the method "ReadableStreamBYOBRequest.respondWithNewView".
//
// The returned bool will be false if there is no such method.
func (this ReadableStreamBYOBRequest) RespondWithNewView(view ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamBYOBRequestRespondWithNewView(
		this.Ref(), js.Pointer(&_ok),
		view.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RespondWithNewViewFunc returns the method "ReadableStreamBYOBRequest.respondWithNewView".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStreamBYOBRequest) RespondWithNewViewFunc() (fn js.Func[func(view ArrayBufferView)]) {
	return fn.FromRef(
		bindings.ReadableStreamBYOBRequestRespondWithNewViewFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this ReadableByteStreamController) ByobRequest() (ReadableStreamBYOBRequest, bool) {
	var _ok bool
	_ret := bindings.GetReadableByteStreamControllerByobRequest(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadableStreamBYOBRequest{}.FromRef(_ret), _ok
}

// DesiredSize returns the value of property "ReadableByteStreamController.desiredSize".
//
// The returned bool will be false if there is no such property.
func (this ReadableByteStreamController) DesiredSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetReadableByteStreamControllerDesiredSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Close calls the method "ReadableByteStreamController.close".
//
// The returned bool will be false if there is no such method.
func (this ReadableByteStreamController) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallReadableByteStreamControllerClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "ReadableByteStreamController.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableByteStreamController) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReadableByteStreamControllerCloseFunc(
			this.Ref(),
		),
	)
}

// Enqueue calls the method "ReadableByteStreamController.enqueue".
//
// The returned bool will be false if there is no such method.
func (this ReadableByteStreamController) Enqueue(chunk ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallReadableByteStreamControllerEnqueue(
		this.Ref(), js.Pointer(&_ok),
		chunk.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EnqueueFunc returns the method "ReadableByteStreamController.enqueue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableByteStreamController) EnqueueFunc() (fn js.Func[func(chunk ArrayBufferView)]) {
	return fn.FromRef(
		bindings.ReadableByteStreamControllerEnqueueFunc(
			this.Ref(),
		),
	)
}

// Error calls the method "ReadableByteStreamController.error".
//
// The returned bool will be false if there is no such method.
func (this ReadableByteStreamController) Error(e js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallReadableByteStreamControllerError(
		this.Ref(), js.Pointer(&_ok),
		e.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ErrorFunc returns the method "ReadableByteStreamController.error".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableByteStreamController) ErrorFunc() (fn js.Func[func(e js.Any)]) {
	return fn.FromRef(
		bindings.ReadableByteStreamControllerErrorFunc(
			this.Ref(),
		),
	)
}

// Error1 calls the method "ReadableByteStreamController.error".
//
// The returned bool will be false if there is no such method.
func (this ReadableByteStreamController) Error1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallReadableByteStreamControllerError1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Error1Func returns the method "ReadableByteStreamController.error".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableByteStreamController) Error1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReadableByteStreamControllerError1Func(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this ReadableStreamDefaultController) DesiredSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetReadableStreamDefaultControllerDesiredSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Close calls the method "ReadableStreamDefaultController.close".
//
// The returned bool will be false if there is no such method.
func (this ReadableStreamDefaultController) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamDefaultControllerClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "ReadableStreamDefaultController.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStreamDefaultController) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultControllerCloseFunc(
			this.Ref(),
		),
	)
}

// Enqueue calls the method "ReadableStreamDefaultController.enqueue".
//
// The returned bool will be false if there is no such method.
func (this ReadableStreamDefaultController) Enqueue(chunk js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamDefaultControllerEnqueue(
		this.Ref(), js.Pointer(&_ok),
		chunk.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EnqueueFunc returns the method "ReadableStreamDefaultController.enqueue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStreamDefaultController) EnqueueFunc() (fn js.Func[func(chunk js.Any)]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultControllerEnqueueFunc(
			this.Ref(),
		),
	)
}

// Enqueue1 calls the method "ReadableStreamDefaultController.enqueue".
//
// The returned bool will be false if there is no such method.
func (this ReadableStreamDefaultController) Enqueue1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamDefaultControllerEnqueue1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Enqueue1Func returns the method "ReadableStreamDefaultController.enqueue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStreamDefaultController) Enqueue1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultControllerEnqueue1Func(
			this.Ref(),
		),
	)
}

// Error calls the method "ReadableStreamDefaultController.error".
//
// The returned bool will be false if there is no such method.
func (this ReadableStreamDefaultController) Error(e js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamDefaultControllerError(
		this.Ref(), js.Pointer(&_ok),
		e.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ErrorFunc returns the method "ReadableStreamDefaultController.error".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStreamDefaultController) ErrorFunc() (fn js.Func[func(e js.Any)]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultControllerErrorFunc(
			this.Ref(),
		),
	)
}

// Error1 calls the method "ReadableStreamDefaultController.error".
//
// The returned bool will be false if there is no such method.
func (this ReadableStreamDefaultController) Error1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamDefaultControllerError1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Error1Func returns the method "ReadableStreamDefaultController.error".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStreamDefaultController) Error1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultControllerError1Func(
			this.Ref(),
		),
	)
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
	PreventCancel bool

	FFI_USE_PreventCancel bool // for PreventCancel.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReadableStreamIteratorOptions with all fields set.
func (p ReadableStreamIteratorOptions) FromRef(ref js.Ref) ReadableStreamIteratorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ReadableStreamIteratorOptions ReadableStreamIteratorOptions [// ReadableStreamIteratorOptions] [0x140009f10e0 0x140009f1180] 0x1400099db60 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 RelativeOrientationReadingValues RelativeOrientationReadingValues [// RelativeOrientationReadingValues] [0x140009f1220] 0x1400099db90 {0 0}} in the application heap.
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

func NewRelativeOrientationSensor(sensorOptions OrientationSensorOptions) RelativeOrientationSensor {
	return RelativeOrientationSensor{}.FromRef(
		bindings.NewRelativeOrientationSensorByRelativeOrientationSensor(
			js.Pointer(&sensorOptions)),
	)
}

func NewRelativeOrientationSensorByRelativeOrientationSensor1() RelativeOrientationSensor {
	return RelativeOrientationSensor{}.FromRef(
		bindings.NewRelativeOrientationSensorByRelativeOrientationSensor1(),
	)
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

// ToJSON calls the method "ReportBody.toJSON".
//
// The returned bool will be false if there is no such method.
func (this ReportBody) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallReportBodyToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "ReportBody.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReportBody) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.ReportBodyToJSONFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this Report) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetReportType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Url returns the value of property "Report.url".
//
// The returned bool will be false if there is no such property.
func (this Report) Url() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetReportUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Body returns the value of property "Report.body".
//
// The returned bool will be false if there is no such property.
func (this Report) Body() (ReportBody, bool) {
	var _ok bool
	_ret := bindings.GetReportBody(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReportBody{}.FromRef(_ret), _ok
}

// ToJSON calls the method "Report.toJSON".
//
// The returned bool will be false if there is no such method.
func (this Report) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallReportToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "Report.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Report) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.ReportToJSONFunc(
			this.Ref(),
		),
	)
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
	ModifiedBid float64
	// DataVersion is "ReportResultBrowserSignals.dataVersion"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 ReportResultBrowserSignals ReportResultBrowserSignals [// ReportResultBrowserSignals] [0x140009f1360 0x140009f1400 0x140009f14a0 0x140009f15e0 0x140009f1720 0x140009f17c0 0x140009f1860 0x140009f1900 0x140009f19a0 0x140009f1a40 0x140009f1ae0 0x140009f1b80 0x140009f1c20 0x140009f1cc0 0x140009f1540 0x140009f1680] 0x1400099dbc0 {0 0}} in the application heap.
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
	AdCost float64
	// Seller is "ReportWinBrowserSignals.seller"
	//
	// Optional
	Seller js.String
	// MadeHighestScoringOtherBid is "ReportWinBrowserSignals.madeHighestScoringOtherBid"
	//
	// Optional
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
	ModelingSignals uint16
	// DataVersion is "ReportWinBrowserSignals.dataVersion"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 ReportWinBrowserSignals ReportWinBrowserSignals [// ReportWinBrowserSignals] [0x140009f1d60 0x140009f1ea0 0x140009f1f40 0x140009f60a0 0x140009f6140 0x140009f61e0 0x140009f6320 0x140009f6460 0x140009f6500 0x140009f65a0 0x140009f6640 0x140009f66e0 0x140009f6780 0x140009f6820 0x140009f68c0 0x140009f6960 0x140009f6a00 0x140009f1e00 0x140009f6000 0x140009f6280 0x140009f63c0] 0x1400099dc20 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 ReportingBrowserSignals ReportingBrowserSignals [// ReportingBrowserSignals] [0x140009f6aa0 0x140009f6b40 0x140009f6be0 0x140009f6c80 0x140009f6d20 0x140009f6dc0 0x140009f6e60 0x140009f6f00 0x140009f6fa0 0x140009f7040] 0x1400099dc80 {0 0}} in the application heap.
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
	Buffered bool

	FFI_USE_Buffered bool // for Buffered.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReportingObserverOptions with all fields set.
func (p ReportingObserverOptions) FromRef(ref js.Ref) ReportingObserverOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ReportingObserverOptions ReportingObserverOptions [// ReportingObserverOptions] [0x140009f70e0 0x140009f7180 0x140009f7220] 0x1400099dcc8 {0 0}} in the application heap.
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

func NewReportingObserver(callback js.Func[func(reports js.Array[Report], observer ReportingObserver)], options ReportingObserverOptions) ReportingObserver {
	return ReportingObserver{}.FromRef(
		bindings.NewReportingObserverByReportingObserver(
			callback.Ref(),
			js.Pointer(&options)),
	)
}

func NewReportingObserverByReportingObserver1(callback js.Func[func(reports js.Array[Report], observer ReportingObserver)]) ReportingObserver {
	return ReportingObserver{}.FromRef(
		bindings.NewReportingObserverByReportingObserver1(
			callback.Ref()),
	)
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

// Observe calls the method "ReportingObserver.observe".
//
// The returned bool will be false if there is no such method.
func (this ReportingObserver) Observe() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallReportingObserverObserve(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ObserveFunc returns the method "ReportingObserver.observe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReportingObserver) ObserveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReportingObserverObserveFunc(
			this.Ref(),
		),
	)
}

// Disconnect calls the method "ReportingObserver.disconnect".
//
// The returned bool will be false if there is no such method.
func (this ReportingObserver) Disconnect() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallReportingObserverDisconnect(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DisconnectFunc returns the method "ReportingObserver.disconnect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReportingObserver) DisconnectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReportingObserverDisconnectFunc(
			this.Ref(),
		),
	)
}

// TakeRecords calls the method "ReportingObserver.takeRecords".
//
// The returned bool will be false if there is no such method.
func (this ReportingObserver) TakeRecords() (ReportList, bool) {
	var _ok bool
	_ret := bindings.CallReportingObserverTakeRecords(
		this.Ref(), js.Pointer(&_ok),
	)

	return ReportList{}.FromRef(_ret), _ok
}

// TakeRecordsFunc returns the method "ReportingObserver.takeRecords".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReportingObserver) TakeRecordsFunc() (fn js.Func[func() ReportList]) {
	return fn.FromRef(
		bindings.ReportingObserverTakeRecordsFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this ResizeObserverSize) InlineSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetResizeObserverSizeInlineSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// BlockSize returns the value of property "ResizeObserverSize.blockSize".
//
// The returned bool will be false if there is no such property.
func (this ResizeObserverSize) BlockSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetResizeObserverSizeBlockSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this ResizeObserverEntry) Target() (Element, bool) {
	var _ok bool
	_ret := bindings.GetResizeObserverEntryTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// ContentRect returns the value of property "ResizeObserverEntry.contentRect".
//
// The returned bool will be false if there is no such property.
func (this ResizeObserverEntry) ContentRect() (DOMRectReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetResizeObserverEntryContentRect(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMRectReadOnly{}.FromRef(_ret), _ok
}

// BorderBoxSize returns the value of property "ResizeObserverEntry.borderBoxSize".
//
// The returned bool will be false if there is no such property.
func (this ResizeObserverEntry) BorderBoxSize() (js.FrozenArray[ResizeObserverSize], bool) {
	var _ok bool
	_ret := bindings.GetResizeObserverEntryBorderBoxSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[ResizeObserverSize]{}.FromRef(_ret), _ok
}

// ContentBoxSize returns the value of property "ResizeObserverEntry.contentBoxSize".
//
// The returned bool will be false if there is no such property.
func (this ResizeObserverEntry) ContentBoxSize() (js.FrozenArray[ResizeObserverSize], bool) {
	var _ok bool
	_ret := bindings.GetResizeObserverEntryContentBoxSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[ResizeObserverSize]{}.FromRef(_ret), _ok
}

// DevicePixelContentBoxSize returns the value of property "ResizeObserverEntry.devicePixelContentBoxSize".
//
// The returned bool will be false if there is no such property.
func (this ResizeObserverEntry) DevicePixelContentBoxSize() (js.FrozenArray[ResizeObserverSize], bool) {
	var _ok bool
	_ret := bindings.GetResizeObserverEntryDevicePixelContentBoxSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[ResizeObserverSize]{}.FromRef(_ret), _ok
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

// New creates a new {0x140004cc0e0 ResizeObserverOptions ResizeObserverOptions [// ResizeObserverOptions] [0x140009f7400] 0x1400099dd28 {0 0}} in the application heap.
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

func NewResizeObserver(callback js.Func[func(entries js.Array[ResizeObserverEntry], observer ResizeObserver)]) ResizeObserver {
	return ResizeObserver{}.FromRef(
		bindings.NewResizeObserverByResizeObserver(
			callback.Ref()),
	)
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

// Observe calls the method "ResizeObserver.observe".
//
// The returned bool will be false if there is no such method.
func (this ResizeObserver) Observe(target Element, options ResizeObserverOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallResizeObserverObserve(
		this.Ref(), js.Pointer(&_ok),
		target.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ObserveFunc returns the method "ResizeObserver.observe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ResizeObserver) ObserveFunc() (fn js.Func[func(target Element, options ResizeObserverOptions)]) {
	return fn.FromRef(
		bindings.ResizeObserverObserveFunc(
			this.Ref(),
		),
	)
}

// Observe1 calls the method "ResizeObserver.observe".
//
// The returned bool will be false if there is no such method.
func (this ResizeObserver) Observe1(target Element) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallResizeObserverObserve1(
		this.Ref(), js.Pointer(&_ok),
		target.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Observe1Func returns the method "ResizeObserver.observe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ResizeObserver) Observe1Func() (fn js.Func[func(target Element)]) {
	return fn.FromRef(
		bindings.ResizeObserverObserve1Func(
			this.Ref(),
		),
	)
}

// Unobserve calls the method "ResizeObserver.unobserve".
//
// The returned bool will be false if there is no such method.
func (this ResizeObserver) Unobserve(target Element) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallResizeObserverUnobserve(
		this.Ref(), js.Pointer(&_ok),
		target.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UnobserveFunc returns the method "ResizeObserver.unobserve".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ResizeObserver) UnobserveFunc() (fn js.Func[func(target Element)]) {
	return fn.FromRef(
		bindings.ResizeObserverUnobserveFunc(
			this.Ref(),
		),
	)
}

// Disconnect calls the method "ResizeObserver.disconnect".
//
// The returned bool will be false if there is no such method.
func (this ResizeObserver) Disconnect() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallResizeObserverDisconnect(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DisconnectFunc returns the method "ResizeObserver.disconnect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ResizeObserver) DisconnectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ResizeObserverDisconnectFunc(
			this.Ref(),
		),
	)
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

// New creates a new {0x140004cc0e0 RsaHashedImportParams RsaHashedImportParams [// RsaHashedImportParams] [0x140009f74a0 0x140009f7540] 0x1400099ddd0 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 RsaHashedKeyAlgorithm RsaHashedKeyAlgorithm [// RsaHashedKeyAlgorithm] [0x140009f75e0 0x140009f7680 0x140009f7720 0x140009f77c0] 0x1400099de30 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 RsaHashedKeyGenParams RsaHashedKeyGenParams [// RsaHashedKeyGenParams] [0x140009f7860 0x140009f7900 0x140009f79a0 0x140009f7a40] 0x1400099de90 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 RsaKeyAlgorithm RsaKeyAlgorithm [// RsaKeyAlgorithm] [0x140009f7ae0 0x140009f7b80 0x140009f7c20] 0x1400099def0 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 RsaKeyGenParams RsaKeyGenParams [// RsaKeyGenParams] [0x140009f7cc0 0x140009f7d60 0x140009f7e00] 0x1400099df38 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 RsaOaepParams RsaOaepParams [// RsaOaepParams] [0x140009f7ea0 0x140009f7f40] 0x1400099df68 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 RsaPssParams RsaPssParams [// RsaPssParams] [0x14000a00000 0x14000a000a0] 0x1400099df98 {0 0}} in the application heap.
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
	Bubbles bool
	// Cancelable is "SFrameTransformErrorEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "SFrameTransformErrorEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 SFrameTransformErrorEventInit SFrameTransformErrorEventInit [// SFrameTransformErrorEventInit] [0x14000a00140 0x14000a001e0 0x14000a00280 0x14000a00320 0x14000a00460 0x14000a005a0 0x14000a003c0 0x14000a00500 0x14000a00640] 0x1400099dfc8 {0 0}} in the application heap.
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

func NewSFrameTransformErrorEvent(typ js.String, eventInitDict SFrameTransformErrorEventInit) SFrameTransformErrorEvent {
	return SFrameTransformErrorEvent{}.FromRef(
		bindings.NewSFrameTransformErrorEventBySFrameTransformErrorEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
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
// The returned bool will be false if there is no such property.
func (this SFrameTransformErrorEvent) ErrorType() (SFrameTransformErrorEventType, bool) {
	var _ok bool
	_ret := bindings.GetSFrameTransformErrorEventErrorType(
		this.Ref(), js.Pointer(&_ok),
	)
	return SFrameTransformErrorEventType(_ret), _ok
}

// KeyID returns the value of property "SFrameTransformErrorEvent.keyID".
//
// The returned bool will be false if there is no such property.
func (this SFrameTransformErrorEvent) KeyID() (CryptoKeyID, bool) {
	var _ok bool
	_ret := bindings.GetSFrameTransformErrorEventKeyID(
		this.Ref(), js.Pointer(&_ok),
	)
	return CryptoKeyID{}.FromRef(_ret), _ok
}

// Frame returns the value of property "SFrameTransformErrorEvent.frame".
//
// The returned bool will be false if there is no such property.
func (this SFrameTransformErrorEvent) Frame() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetSFrameTransformErrorEventFrame(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) Target() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// Download returns the value of property "SVGAElement.download".
//
// The returned bool will be false if there is no such property.
func (this SVGAElement) Download() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementDownload(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Download sets the value of property "SVGAElement.download" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) Ping() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementPing(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Ping sets the value of property "SVGAElement.ping" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) Rel() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementRel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Rel sets the value of property "SVGAElement.rel" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) RelList() (DOMTokenList, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementRelList(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMTokenList{}.FromRef(_ret), _ok
}

// Hreflang returns the value of property "SVGAElement.hreflang".
//
// The returned bool will be false if there is no such property.
func (this SVGAElement) Hreflang() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementHreflang(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hreflang sets the value of property "SVGAElement.hreflang" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "SVGAElement.type" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) Text() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Text sets the value of property "SVGAElement.text" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) ReferrerPolicy() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementReferrerPolicy(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ReferrerPolicy sets the value of property "SVGAElement.referrerPolicy" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) Origin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Protocol returns the value of property "SVGAElement.protocol".
//
// The returned bool will be false if there is no such property.
func (this SVGAElement) Protocol() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementProtocol(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Protocol sets the value of property "SVGAElement.protocol" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) Username() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementUsername(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Username sets the value of property "SVGAElement.username" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) Password() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementPassword(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Password sets the value of property "SVGAElement.password" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) Host() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementHost(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Host sets the value of property "SVGAElement.host" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) Hostname() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementHostname(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hostname sets the value of property "SVGAElement.hostname" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) Port() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementPort(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Port sets the value of property "SVGAElement.port" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) Pathname() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementPathname(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Pathname sets the value of property "SVGAElement.pathname" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) Search() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementSearch(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Search sets the value of property "SVGAElement.search" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) Hash() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementHash(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hash sets the value of property "SVGAElement.hash" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAElement) Href() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGAElementHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGAnimatedAngle) BaseVal() (SVGAngle, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedAngleBaseVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAngle{}.FromRef(_ret), _ok
}

// AnimVal returns the value of property "SVGAnimatedAngle.animVal".
//
// The returned bool will be false if there is no such property.
func (this SVGAnimatedAngle) AnimVal() (SVGAngle, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedAngleAnimVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAngle{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGAnimatedBoolean) BaseVal() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedBooleanBaseVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// BaseVal sets the value of property "SVGAnimatedBoolean.baseVal" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAnimatedBoolean) AnimVal() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedBooleanAnimVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}
