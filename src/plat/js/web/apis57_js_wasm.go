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

func NewWebGLContextEvent(typ js.String, eventInit WebGLContextEventInit) WebGLContextEvent {
	return WebGLContextEvent{}.FromRef(
		bindings.NewWebGLContextEventByWebGLContextEvent(
			typ.Ref(),
			js.Pointer(&eventInit)),
	)
}

func NewWebGLContextEventByWebGLContextEvent1(typ js.String) WebGLContextEvent {
	return WebGLContextEvent{}.FromRef(
		bindings.NewWebGLContextEventByWebGLContextEvent1(
			typ.Ref()),
	)
}

type WebGLContextEvent struct {
	Event
}

func (this WebGLContextEvent) Once() WebGLContextEvent {
	this.Ref().Once()
	return this
}

func (this WebGLContextEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this WebGLContextEvent) FromRef(ref js.Ref) WebGLContextEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this WebGLContextEvent) Free() {
	this.Ref().Free()
}

// StatusMessage returns the value of property "WebGLContextEvent.statusMessage".
//
// The returned bool will be false if there is no such property.
func (this WebGLContextEvent) StatusMessage() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWebGLContextEventStatusMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type WebGLObject struct {
	ref js.Ref
}

func (this WebGLObject) Once() WebGLObject {
	this.Ref().Once()
	return this
}

func (this WebGLObject) Ref() js.Ref {
	return this.ref
}

func (this WebGLObject) FromRef(ref js.Ref) WebGLObject {
	this.ref = ref
	return this
}

func (this WebGLObject) Free() {
	this.Ref().Free()
}

const (
	WebSocket_CONNECTING uint16 = 0
	WebSocket_OPEN       uint16 = 1
	WebSocket_CLOSING    uint16 = 2
	WebSocket_CLOSED     uint16 = 3
)

func NewWebSocket(url js.String, protocols OneOf_String_ArrayString) WebSocket {
	return WebSocket{}.FromRef(
		bindings.NewWebSocketByWebSocket(
			url.Ref(),
			protocols.Ref()),
	)
}

func NewWebSocketByWebSocket1(url js.String) WebSocket {
	return WebSocket{}.FromRef(
		bindings.NewWebSocketByWebSocket1(
			url.Ref()),
	)
}

type WebSocket struct {
	EventTarget
}

func (this WebSocket) Once() WebSocket {
	this.Ref().Once()
	return this
}

func (this WebSocket) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this WebSocket) FromRef(ref js.Ref) WebSocket {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this WebSocket) Free() {
	this.Ref().Free()
}

// Url returns the value of property "WebSocket.url".
//
// The returned bool will be false if there is no such property.
func (this WebSocket) Url() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWebSocketUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ReadyState returns the value of property "WebSocket.readyState".
//
// The returned bool will be false if there is no such property.
func (this WebSocket) ReadyState() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetWebSocketReadyState(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// BufferedAmount returns the value of property "WebSocket.bufferedAmount".
//
// The returned bool will be false if there is no such property.
func (this WebSocket) BufferedAmount() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetWebSocketBufferedAmount(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// Extensions returns the value of property "WebSocket.extensions".
//
// The returned bool will be false if there is no such property.
func (this WebSocket) Extensions() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWebSocketExtensions(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Protocol returns the value of property "WebSocket.protocol".
//
// The returned bool will be false if there is no such property.
func (this WebSocket) Protocol() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWebSocketProtocol(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// BinaryType returns the value of property "WebSocket.binaryType".
//
// The returned bool will be false if there is no such property.
func (this WebSocket) BinaryType() (BinaryType, bool) {
	var _ok bool
	_ret := bindings.GetWebSocketBinaryType(
		this.Ref(), js.Pointer(&_ok),
	)
	return BinaryType(_ret), _ok
}

// BinaryType sets the value of property "WebSocket.binaryType" to val.
//
// It returns false if the property cannot be set.
func (this WebSocket) SetBinaryType(val BinaryType) bool {
	return js.True == bindings.SetWebSocketBinaryType(
		this.Ref(),
		uint32(val),
	)
}

// Close calls the method "WebSocket.close".
//
// The returned bool will be false if there is no such method.
func (this WebSocket) Close(code uint16, reason js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebSocketClose(
		this.Ref(), js.Pointer(&_ok),
		uint32(code),
		reason.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "WebSocket.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebSocket) CloseFunc() (fn js.Func[func(code uint16, reason js.String)]) {
	return fn.FromRef(
		bindings.WebSocketCloseFunc(
			this.Ref(),
		),
	)
}

// Close1 calls the method "WebSocket.close".
//
// The returned bool will be false if there is no such method.
func (this WebSocket) Close1(code uint16) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebSocketClose1(
		this.Ref(), js.Pointer(&_ok),
		uint32(code),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Close1Func returns the method "WebSocket.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebSocket) Close1Func() (fn js.Func[func(code uint16)]) {
	return fn.FromRef(
		bindings.WebSocketClose1Func(
			this.Ref(),
		),
	)
}

// Close2 calls the method "WebSocket.close".
//
// The returned bool will be false if there is no such method.
func (this WebSocket) Close2() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebSocketClose2(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Close2Func returns the method "WebSocket.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebSocket) Close2Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebSocketClose2Func(
			this.Ref(),
		),
	)
}

// Send calls the method "WebSocket.send".
//
// The returned bool will be false if there is no such method.
func (this WebSocket) Send(data OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebSocketSend(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SendFunc returns the method "WebSocket.send".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebSocket) SendFunc() (fn js.Func[func(data OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String)]) {
	return fn.FromRef(
		bindings.WebSocketSendFunc(
			this.Ref(),
		),
	)
}

type WebTransportHash struct {
	// Algorithm is "WebTransportHash.algorithm"
	//
	// Optional
	Algorithm js.String
	// Value is "WebTransportHash.value"
	//
	// Optional
	Value BufferSource

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebTransportHash with all fields set.
func (p WebTransportHash) FromRef(ref js.Ref) WebTransportHash {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebTransportHash in the application heap.
func (p WebTransportHash) New() js.Ref {
	return bindings.WebTransportHashJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WebTransportHash) UpdateFrom(ref js.Ref) {
	bindings.WebTransportHashJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WebTransportHash) Update(ref js.Ref) {
	bindings.WebTransportHashJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type WebTransportCongestionControl uint32

const (
	_ WebTransportCongestionControl = iota

	WebTransportCongestionControl_DEFAULT
	WebTransportCongestionControl_THROUGHPUT
	WebTransportCongestionControl_LOW_LATENCY
)

func (WebTransportCongestionControl) FromRef(str js.Ref) WebTransportCongestionControl {
	return WebTransportCongestionControl(bindings.ConstOfWebTransportCongestionControl(str))
}

func (x WebTransportCongestionControl) String() (string, bool) {
	switch x {
	case WebTransportCongestionControl_DEFAULT:
		return "default", true
	case WebTransportCongestionControl_THROUGHPUT:
		return "throughput", true
	case WebTransportCongestionControl_LOW_LATENCY:
		return "low-latency", true
	default:
		return "", false
	}
}

type WebTransportOptions struct {
	// AllowPooling is "WebTransportOptions.allowPooling"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_AllowPooling MUST be set to true to make this field effective.
	AllowPooling bool
	// RequireUnreliable is "WebTransportOptions.requireUnreliable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_RequireUnreliable MUST be set to true to make this field effective.
	RequireUnreliable bool
	// ServerCertificateHashes is "WebTransportOptions.serverCertificateHashes"
	//
	// Optional
	ServerCertificateHashes js.Array[WebTransportHash]
	// CongestionControl is "WebTransportOptions.congestionControl"
	//
	// Optional, defaults to "default".
	CongestionControl WebTransportCongestionControl

	FFI_USE_AllowPooling      bool // for AllowPooling.
	FFI_USE_RequireUnreliable bool // for RequireUnreliable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebTransportOptions with all fields set.
func (p WebTransportOptions) FromRef(ref js.Ref) WebTransportOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebTransportOptions in the application heap.
func (p WebTransportOptions) New() js.Ref {
	return bindings.WebTransportOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WebTransportOptions) UpdateFrom(ref js.Ref) {
	bindings.WebTransportOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WebTransportOptions) Update(ref js.Ref) {
	bindings.WebTransportOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type WebTransportDatagramStats struct {
	// Timestamp is "WebTransportDatagramStats.timestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timestamp MUST be set to true to make this field effective.
	Timestamp DOMHighResTimeStamp
	// ExpiredOutgoing is "WebTransportDatagramStats.expiredOutgoing"
	//
	// Optional
	//
	// NOTE: FFI_USE_ExpiredOutgoing MUST be set to true to make this field effective.
	ExpiredOutgoing uint64
	// DroppedIncoming is "WebTransportDatagramStats.droppedIncoming"
	//
	// Optional
	//
	// NOTE: FFI_USE_DroppedIncoming MUST be set to true to make this field effective.
	DroppedIncoming uint64
	// LostOutgoing is "WebTransportDatagramStats.lostOutgoing"
	//
	// Optional
	//
	// NOTE: FFI_USE_LostOutgoing MUST be set to true to make this field effective.
	LostOutgoing uint64

	FFI_USE_Timestamp       bool // for Timestamp.
	FFI_USE_ExpiredOutgoing bool // for ExpiredOutgoing.
	FFI_USE_DroppedIncoming bool // for DroppedIncoming.
	FFI_USE_LostOutgoing    bool // for LostOutgoing.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebTransportDatagramStats with all fields set.
func (p WebTransportDatagramStats) FromRef(ref js.Ref) WebTransportDatagramStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebTransportDatagramStats in the application heap.
func (p WebTransportDatagramStats) New() js.Ref {
	return bindings.WebTransportDatagramStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WebTransportDatagramStats) UpdateFrom(ref js.Ref) {
	bindings.WebTransportDatagramStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WebTransportDatagramStats) Update(ref js.Ref) {
	bindings.WebTransportDatagramStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type WebTransportStats struct {
	// Timestamp is "WebTransportStats.timestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timestamp MUST be set to true to make this field effective.
	Timestamp DOMHighResTimeStamp
	// BytesSent is "WebTransportStats.bytesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesSent MUST be set to true to make this field effective.
	BytesSent uint64
	// PacketsSent is "WebTransportStats.packetsSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsSent MUST be set to true to make this field effective.
	PacketsSent uint64
	// PacketsLost is "WebTransportStats.packetsLost"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsLost MUST be set to true to make this field effective.
	PacketsLost uint64
	// NumOutgoingStreamsCreated is "WebTransportStats.numOutgoingStreamsCreated"
	//
	// Optional
	//
	// NOTE: FFI_USE_NumOutgoingStreamsCreated MUST be set to true to make this field effective.
	NumOutgoingStreamsCreated uint32
	// NumIncomingStreamsCreated is "WebTransportStats.numIncomingStreamsCreated"
	//
	// Optional
	//
	// NOTE: FFI_USE_NumIncomingStreamsCreated MUST be set to true to make this field effective.
	NumIncomingStreamsCreated uint32
	// BytesReceived is "WebTransportStats.bytesReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesReceived MUST be set to true to make this field effective.
	BytesReceived uint64
	// PacketsReceived is "WebTransportStats.packetsReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_PacketsReceived MUST be set to true to make this field effective.
	PacketsReceived uint64
	// SmoothedRtt is "WebTransportStats.smoothedRtt"
	//
	// Optional
	//
	// NOTE: FFI_USE_SmoothedRtt MUST be set to true to make this field effective.
	SmoothedRtt DOMHighResTimeStamp
	// RttVariation is "WebTransportStats.rttVariation"
	//
	// Optional
	//
	// NOTE: FFI_USE_RttVariation MUST be set to true to make this field effective.
	RttVariation DOMHighResTimeStamp
	// MinRtt is "WebTransportStats.minRtt"
	//
	// Optional
	//
	// NOTE: FFI_USE_MinRtt MUST be set to true to make this field effective.
	MinRtt DOMHighResTimeStamp
	// Datagrams is "WebTransportStats.datagrams"
	//
	// Optional
	Datagrams WebTransportDatagramStats
	// EstimatedSendRate is "WebTransportStats.estimatedSendRate"
	//
	// Optional
	//
	// NOTE: FFI_USE_EstimatedSendRate MUST be set to true to make this field effective.
	EstimatedSendRate uint64

	FFI_USE_Timestamp                 bool // for Timestamp.
	FFI_USE_BytesSent                 bool // for BytesSent.
	FFI_USE_PacketsSent               bool // for PacketsSent.
	FFI_USE_PacketsLost               bool // for PacketsLost.
	FFI_USE_NumOutgoingStreamsCreated bool // for NumOutgoingStreamsCreated.
	FFI_USE_NumIncomingStreamsCreated bool // for NumIncomingStreamsCreated.
	FFI_USE_BytesReceived             bool // for BytesReceived.
	FFI_USE_PacketsReceived           bool // for PacketsReceived.
	FFI_USE_SmoothedRtt               bool // for SmoothedRtt.
	FFI_USE_RttVariation              bool // for RttVariation.
	FFI_USE_MinRtt                    bool // for MinRtt.
	FFI_USE_EstimatedSendRate         bool // for EstimatedSendRate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebTransportStats with all fields set.
func (p WebTransportStats) FromRef(ref js.Ref) WebTransportStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebTransportStats in the application heap.
func (p WebTransportStats) New() js.Ref {
	return bindings.WebTransportStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WebTransportStats) UpdateFrom(ref js.Ref) {
	bindings.WebTransportStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WebTransportStats) Update(ref js.Ref) {
	bindings.WebTransportStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type WebTransportCloseInfo struct {
	// CloseCode is "WebTransportCloseInfo.closeCode"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_CloseCode MUST be set to true to make this field effective.
	CloseCode uint32
	// Reason is "WebTransportCloseInfo.reason"
	//
	// Optional, defaults to "".
	Reason js.String

	FFI_USE_CloseCode bool // for CloseCode.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebTransportCloseInfo with all fields set.
func (p WebTransportCloseInfo) FromRef(ref js.Ref) WebTransportCloseInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebTransportCloseInfo in the application heap.
func (p WebTransportCloseInfo) New() js.Ref {
	return bindings.WebTransportCloseInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WebTransportCloseInfo) UpdateFrom(ref js.Ref) {
	bindings.WebTransportCloseInfoJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WebTransportCloseInfo) Update(ref js.Ref) {
	bindings.WebTransportCloseInfoJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type WebTransportReceiveStreamStats struct {
	// Timestamp is "WebTransportReceiveStreamStats.timestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timestamp MUST be set to true to make this field effective.
	Timestamp DOMHighResTimeStamp
	// BytesReceived is "WebTransportReceiveStreamStats.bytesReceived"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesReceived MUST be set to true to make this field effective.
	BytesReceived uint64
	// BytesRead is "WebTransportReceiveStreamStats.bytesRead"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesRead MUST be set to true to make this field effective.
	BytesRead uint64

	FFI_USE_Timestamp     bool // for Timestamp.
	FFI_USE_BytesReceived bool // for BytesReceived.
	FFI_USE_BytesRead     bool // for BytesRead.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebTransportReceiveStreamStats with all fields set.
func (p WebTransportReceiveStreamStats) FromRef(ref js.Ref) WebTransportReceiveStreamStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebTransportReceiveStreamStats in the application heap.
func (p WebTransportReceiveStreamStats) New() js.Ref {
	return bindings.WebTransportReceiveStreamStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WebTransportReceiveStreamStats) UpdateFrom(ref js.Ref) {
	bindings.WebTransportReceiveStreamStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WebTransportReceiveStreamStats) Update(ref js.Ref) {
	bindings.WebTransportReceiveStreamStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewWebTransportReceiveStream(underlyingSource js.Object, strategy QueuingStrategy) WebTransportReceiveStream {
	return WebTransportReceiveStream{}.FromRef(
		bindings.NewWebTransportReceiveStreamByWebTransportReceiveStream(
			underlyingSource.Ref(),
			js.Pointer(&strategy)),
	)
}

func NewWebTransportReceiveStreamByWebTransportReceiveStream1(underlyingSource js.Object) WebTransportReceiveStream {
	return WebTransportReceiveStream{}.FromRef(
		bindings.NewWebTransportReceiveStreamByWebTransportReceiveStream1(
			underlyingSource.Ref()),
	)
}

func NewWebTransportReceiveStreamByWebTransportReceiveStream2() WebTransportReceiveStream {
	return WebTransportReceiveStream{}.FromRef(
		bindings.NewWebTransportReceiveStreamByWebTransportReceiveStream2(),
	)
}

type WebTransportReceiveStream struct {
	ReadableStream
}

func (this WebTransportReceiveStream) Once() WebTransportReceiveStream {
	this.Ref().Once()
	return this
}

func (this WebTransportReceiveStream) Ref() js.Ref {
	return this.ReadableStream.Ref()
}

func (this WebTransportReceiveStream) FromRef(ref js.Ref) WebTransportReceiveStream {
	this.ReadableStream = this.ReadableStream.FromRef(ref)
	return this
}

func (this WebTransportReceiveStream) Free() {
	this.Ref().Free()
}

// GetStats calls the method "WebTransportReceiveStream.getStats".
//
// The returned bool will be false if there is no such method.
func (this WebTransportReceiveStream) GetStats() (js.Promise[WebTransportReceiveStreamStats], bool) {
	var _ok bool
	_ret := bindings.CallWebTransportReceiveStreamGetStats(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[WebTransportReceiveStreamStats]{}.FromRef(_ret), _ok
}

// GetStatsFunc returns the method "WebTransportReceiveStream.getStats".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebTransportReceiveStream) GetStatsFunc() (fn js.Func[func() js.Promise[WebTransportReceiveStreamStats]]) {
	return fn.FromRef(
		bindings.WebTransportReceiveStreamGetStatsFunc(
			this.Ref(),
		),
	)
}

type WebTransportSendStreamStats struct {
	// Timestamp is "WebTransportSendStreamStats.timestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timestamp MUST be set to true to make this field effective.
	Timestamp DOMHighResTimeStamp
	// BytesWritten is "WebTransportSendStreamStats.bytesWritten"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesWritten MUST be set to true to make this field effective.
	BytesWritten uint64
	// BytesSent is "WebTransportSendStreamStats.bytesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesSent MUST be set to true to make this field effective.
	BytesSent uint64
	// BytesAcknowledged is "WebTransportSendStreamStats.bytesAcknowledged"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesAcknowledged MUST be set to true to make this field effective.
	BytesAcknowledged uint64

	FFI_USE_Timestamp         bool // for Timestamp.
	FFI_USE_BytesWritten      bool // for BytesWritten.
	FFI_USE_BytesSent         bool // for BytesSent.
	FFI_USE_BytesAcknowledged bool // for BytesAcknowledged.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebTransportSendStreamStats with all fields set.
func (p WebTransportSendStreamStats) FromRef(ref js.Ref) WebTransportSendStreamStats {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebTransportSendStreamStats in the application heap.
func (p WebTransportSendStreamStats) New() js.Ref {
	return bindings.WebTransportSendStreamStatsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WebTransportSendStreamStats) UpdateFrom(ref js.Ref) {
	bindings.WebTransportSendStreamStatsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WebTransportSendStreamStats) Update(ref js.Ref) {
	bindings.WebTransportSendStreamStatsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewWebTransportSendStream(underlyingSink js.Object, strategy QueuingStrategy) WebTransportSendStream {
	return WebTransportSendStream{}.FromRef(
		bindings.NewWebTransportSendStreamByWebTransportSendStream(
			underlyingSink.Ref(),
			js.Pointer(&strategy)),
	)
}

func NewWebTransportSendStreamByWebTransportSendStream1(underlyingSink js.Object) WebTransportSendStream {
	return WebTransportSendStream{}.FromRef(
		bindings.NewWebTransportSendStreamByWebTransportSendStream1(
			underlyingSink.Ref()),
	)
}

func NewWebTransportSendStreamByWebTransportSendStream2() WebTransportSendStream {
	return WebTransportSendStream{}.FromRef(
		bindings.NewWebTransportSendStreamByWebTransportSendStream2(),
	)
}

type WebTransportSendStream struct {
	WritableStream
}

func (this WebTransportSendStream) Once() WebTransportSendStream {
	this.Ref().Once()
	return this
}

func (this WebTransportSendStream) Ref() js.Ref {
	return this.WritableStream.Ref()
}

func (this WebTransportSendStream) FromRef(ref js.Ref) WebTransportSendStream {
	this.WritableStream = this.WritableStream.FromRef(ref)
	return this
}

func (this WebTransportSendStream) Free() {
	this.Ref().Free()
}

// SendOrder returns the value of property "WebTransportSendStream.sendOrder".
//
// The returned bool will be false if there is no such property.
func (this WebTransportSendStream) SendOrder() (int64, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportSendStreamSendOrder(
		this.Ref(), js.Pointer(&_ok),
	)
	return int64(_ret), _ok
}

// SendOrder sets the value of property "WebTransportSendStream.sendOrder" to val.
//
// It returns false if the property cannot be set.
func (this WebTransportSendStream) SetSendOrder(val int64) bool {
	return js.True == bindings.SetWebTransportSendStreamSendOrder(
		this.Ref(),
		float64(val),
	)
}

// GetStats calls the method "WebTransportSendStream.getStats".
//
// The returned bool will be false if there is no such method.
func (this WebTransportSendStream) GetStats() (js.Promise[WebTransportSendStreamStats], bool) {
	var _ok bool
	_ret := bindings.CallWebTransportSendStreamGetStats(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[WebTransportSendStreamStats]{}.FromRef(_ret), _ok
}

// GetStatsFunc returns the method "WebTransportSendStream.getStats".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebTransportSendStream) GetStatsFunc() (fn js.Func[func() js.Promise[WebTransportSendStreamStats]]) {
	return fn.FromRef(
		bindings.WebTransportSendStreamGetStatsFunc(
			this.Ref(),
		),
	)
}

type WebTransportBidirectionalStream struct {
	ref js.Ref
}

func (this WebTransportBidirectionalStream) Once() WebTransportBidirectionalStream {
	this.Ref().Once()
	return this
}

func (this WebTransportBidirectionalStream) Ref() js.Ref {
	return this.ref
}

func (this WebTransportBidirectionalStream) FromRef(ref js.Ref) WebTransportBidirectionalStream {
	this.ref = ref
	return this
}

func (this WebTransportBidirectionalStream) Free() {
	this.Ref().Free()
}

// Readable returns the value of property "WebTransportBidirectionalStream.readable".
//
// The returned bool will be false if there is no such property.
func (this WebTransportBidirectionalStream) Readable() (WebTransportReceiveStream, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportBidirectionalStreamReadable(
		this.Ref(), js.Pointer(&_ok),
	)
	return WebTransportReceiveStream{}.FromRef(_ret), _ok
}

// Writable returns the value of property "WebTransportBidirectionalStream.writable".
//
// The returned bool will be false if there is no such property.
func (this WebTransportBidirectionalStream) Writable() (WebTransportSendStream, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportBidirectionalStreamWritable(
		this.Ref(), js.Pointer(&_ok),
	)
	return WebTransportSendStream{}.FromRef(_ret), _ok
}

type WebTransportSendStreamOptions struct {
	// SendOrder is "WebTransportSendStreamOptions.sendOrder"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_SendOrder MUST be set to true to make this field effective.
	SendOrder int64

	FFI_USE_SendOrder bool // for SendOrder.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebTransportSendStreamOptions with all fields set.
func (p WebTransportSendStreamOptions) FromRef(ref js.Ref) WebTransportSendStreamOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebTransportSendStreamOptions in the application heap.
func (p WebTransportSendStreamOptions) New() js.Ref {
	return bindings.WebTransportSendStreamOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WebTransportSendStreamOptions) UpdateFrom(ref js.Ref) {
	bindings.WebTransportSendStreamOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WebTransportSendStreamOptions) Update(ref js.Ref) {
	bindings.WebTransportSendStreamOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type WebTransportReliabilityMode uint32

const (
	_ WebTransportReliabilityMode = iota

	WebTransportReliabilityMode_PENDING
	WebTransportReliabilityMode_RELIABLE_ONLY
	WebTransportReliabilityMode_SUPPORTS_UNRELIABLE
)

func (WebTransportReliabilityMode) FromRef(str js.Ref) WebTransportReliabilityMode {
	return WebTransportReliabilityMode(bindings.ConstOfWebTransportReliabilityMode(str))
}

func (x WebTransportReliabilityMode) String() (string, bool) {
	switch x {
	case WebTransportReliabilityMode_PENDING:
		return "pending", true
	case WebTransportReliabilityMode_RELIABLE_ONLY:
		return "reliable-only", true
	case WebTransportReliabilityMode_SUPPORTS_UNRELIABLE:
		return "supports-unreliable", true
	default:
		return "", false
	}
}

type WebTransportDatagramDuplexStream struct {
	ref js.Ref
}

func (this WebTransportDatagramDuplexStream) Once() WebTransportDatagramDuplexStream {
	this.Ref().Once()
	return this
}

func (this WebTransportDatagramDuplexStream) Ref() js.Ref {
	return this.ref
}

func (this WebTransportDatagramDuplexStream) FromRef(ref js.Ref) WebTransportDatagramDuplexStream {
	this.ref = ref
	return this
}

func (this WebTransportDatagramDuplexStream) Free() {
	this.Ref().Free()
}

// Readable returns the value of property "WebTransportDatagramDuplexStream.readable".
//
// The returned bool will be false if there is no such property.
func (this WebTransportDatagramDuplexStream) Readable() (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportDatagramDuplexStreamReadable(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadableStream{}.FromRef(_ret), _ok
}

// Writable returns the value of property "WebTransportDatagramDuplexStream.writable".
//
// The returned bool will be false if there is no such property.
func (this WebTransportDatagramDuplexStream) Writable() (WritableStream, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportDatagramDuplexStreamWritable(
		this.Ref(), js.Pointer(&_ok),
	)
	return WritableStream{}.FromRef(_ret), _ok
}

// MaxDatagramSize returns the value of property "WebTransportDatagramDuplexStream.maxDatagramSize".
//
// The returned bool will be false if there is no such property.
func (this WebTransportDatagramDuplexStream) MaxDatagramSize() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportDatagramDuplexStreamMaxDatagramSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// IncomingMaxAge returns the value of property "WebTransportDatagramDuplexStream.incomingMaxAge".
//
// The returned bool will be false if there is no such property.
func (this WebTransportDatagramDuplexStream) IncomingMaxAge() (float64, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportDatagramDuplexStreamIncomingMaxAge(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// IncomingMaxAge sets the value of property "WebTransportDatagramDuplexStream.incomingMaxAge" to val.
//
// It returns false if the property cannot be set.
func (this WebTransportDatagramDuplexStream) SetIncomingMaxAge(val float64) bool {
	return js.True == bindings.SetWebTransportDatagramDuplexStreamIncomingMaxAge(
		this.Ref(),
		float64(val),
	)
}

// OutgoingMaxAge returns the value of property "WebTransportDatagramDuplexStream.outgoingMaxAge".
//
// The returned bool will be false if there is no such property.
func (this WebTransportDatagramDuplexStream) OutgoingMaxAge() (float64, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportDatagramDuplexStreamOutgoingMaxAge(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// OutgoingMaxAge sets the value of property "WebTransportDatagramDuplexStream.outgoingMaxAge" to val.
//
// It returns false if the property cannot be set.
func (this WebTransportDatagramDuplexStream) SetOutgoingMaxAge(val float64) bool {
	return js.True == bindings.SetWebTransportDatagramDuplexStreamOutgoingMaxAge(
		this.Ref(),
		float64(val),
	)
}

// IncomingHighWaterMark returns the value of property "WebTransportDatagramDuplexStream.incomingHighWaterMark".
//
// The returned bool will be false if there is no such property.
func (this WebTransportDatagramDuplexStream) IncomingHighWaterMark() (float64, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportDatagramDuplexStreamIncomingHighWaterMark(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// IncomingHighWaterMark sets the value of property "WebTransportDatagramDuplexStream.incomingHighWaterMark" to val.
//
// It returns false if the property cannot be set.
func (this WebTransportDatagramDuplexStream) SetIncomingHighWaterMark(val float64) bool {
	return js.True == bindings.SetWebTransportDatagramDuplexStreamIncomingHighWaterMark(
		this.Ref(),
		float64(val),
	)
}

// OutgoingHighWaterMark returns the value of property "WebTransportDatagramDuplexStream.outgoingHighWaterMark".
//
// The returned bool will be false if there is no such property.
func (this WebTransportDatagramDuplexStream) OutgoingHighWaterMark() (float64, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportDatagramDuplexStreamOutgoingHighWaterMark(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// OutgoingHighWaterMark sets the value of property "WebTransportDatagramDuplexStream.outgoingHighWaterMark" to val.
//
// It returns false if the property cannot be set.
func (this WebTransportDatagramDuplexStream) SetOutgoingHighWaterMark(val float64) bool {
	return js.True == bindings.SetWebTransportDatagramDuplexStreamOutgoingHighWaterMark(
		this.Ref(),
		float64(val),
	)
}

func NewWebTransport(url js.String, options WebTransportOptions) WebTransport {
	return WebTransport{}.FromRef(
		bindings.NewWebTransportByWebTransport(
			url.Ref(),
			js.Pointer(&options)),
	)
}

func NewWebTransportByWebTransport1(url js.String) WebTransport {
	return WebTransport{}.FromRef(
		bindings.NewWebTransportByWebTransport1(
			url.Ref()),
	)
}

type WebTransport struct {
	ref js.Ref
}

func (this WebTransport) Once() WebTransport {
	this.Ref().Once()
	return this
}

func (this WebTransport) Ref() js.Ref {
	return this.ref
}

func (this WebTransport) FromRef(ref js.Ref) WebTransport {
	this.ref = ref
	return this
}

func (this WebTransport) Free() {
	this.Ref().Free()
}

// Ready returns the value of property "WebTransport.ready".
//
// The returned bool will be false if there is no such property.
func (this WebTransport) Ready() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.GetWebTransportReady(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Reliability returns the value of property "WebTransport.reliability".
//
// The returned bool will be false if there is no such property.
func (this WebTransport) Reliability() (WebTransportReliabilityMode, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportReliability(
		this.Ref(), js.Pointer(&_ok),
	)
	return WebTransportReliabilityMode(_ret), _ok
}

// CongestionControl returns the value of property "WebTransport.congestionControl".
//
// The returned bool will be false if there is no such property.
func (this WebTransport) CongestionControl() (WebTransportCongestionControl, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportCongestionControl(
		this.Ref(), js.Pointer(&_ok),
	)
	return WebTransportCongestionControl(_ret), _ok
}

// Closed returns the value of property "WebTransport.closed".
//
// The returned bool will be false if there is no such property.
func (this WebTransport) Closed() (js.Promise[WebTransportCloseInfo], bool) {
	var _ok bool
	_ret := bindings.GetWebTransportClosed(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[WebTransportCloseInfo]{}.FromRef(_ret), _ok
}

// Draining returns the value of property "WebTransport.draining".
//
// The returned bool will be false if there is no such property.
func (this WebTransport) Draining() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.GetWebTransportDraining(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Datagrams returns the value of property "WebTransport.datagrams".
//
// The returned bool will be false if there is no such property.
func (this WebTransport) Datagrams() (WebTransportDatagramDuplexStream, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportDatagrams(
		this.Ref(), js.Pointer(&_ok),
	)
	return WebTransportDatagramDuplexStream{}.FromRef(_ret), _ok
}

// IncomingBidirectionalStreams returns the value of property "WebTransport.incomingBidirectionalStreams".
//
// The returned bool will be false if there is no such property.
func (this WebTransport) IncomingBidirectionalStreams() (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportIncomingBidirectionalStreams(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadableStream{}.FromRef(_ret), _ok
}

// IncomingUnidirectionalStreams returns the value of property "WebTransport.incomingUnidirectionalStreams".
//
// The returned bool will be false if there is no such property.
func (this WebTransport) IncomingUnidirectionalStreams() (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportIncomingUnidirectionalStreams(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadableStream{}.FromRef(_ret), _ok
}

// GetStats calls the method "WebTransport.getStats".
//
// The returned bool will be false if there is no such method.
func (this WebTransport) GetStats() (js.Promise[WebTransportStats], bool) {
	var _ok bool
	_ret := bindings.CallWebTransportGetStats(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[WebTransportStats]{}.FromRef(_ret), _ok
}

// GetStatsFunc returns the method "WebTransport.getStats".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebTransport) GetStatsFunc() (fn js.Func[func() js.Promise[WebTransportStats]]) {
	return fn.FromRef(
		bindings.WebTransportGetStatsFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "WebTransport.close".
//
// The returned bool will be false if there is no such method.
func (this WebTransport) Close(closeInfo WebTransportCloseInfo) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebTransportClose(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&closeInfo),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "WebTransport.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebTransport) CloseFunc() (fn js.Func[func(closeInfo WebTransportCloseInfo)]) {
	return fn.FromRef(
		bindings.WebTransportCloseFunc(
			this.Ref(),
		),
	)
}

// Close1 calls the method "WebTransport.close".
//
// The returned bool will be false if there is no such method.
func (this WebTransport) Close1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebTransportClose1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Close1Func returns the method "WebTransport.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebTransport) Close1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebTransportClose1Func(
			this.Ref(),
		),
	)
}

// CreateBidirectionalStream calls the method "WebTransport.createBidirectionalStream".
//
// The returned bool will be false if there is no such method.
func (this WebTransport) CreateBidirectionalStream(options WebTransportSendStreamOptions) (js.Promise[WebTransportBidirectionalStream], bool) {
	var _ok bool
	_ret := bindings.CallWebTransportCreateBidirectionalStream(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[WebTransportBidirectionalStream]{}.FromRef(_ret), _ok
}

// CreateBidirectionalStreamFunc returns the method "WebTransport.createBidirectionalStream".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebTransport) CreateBidirectionalStreamFunc() (fn js.Func[func(options WebTransportSendStreamOptions) js.Promise[WebTransportBidirectionalStream]]) {
	return fn.FromRef(
		bindings.WebTransportCreateBidirectionalStreamFunc(
			this.Ref(),
		),
	)
}

// CreateBidirectionalStream1 calls the method "WebTransport.createBidirectionalStream".
//
// The returned bool will be false if there is no such method.
func (this WebTransport) CreateBidirectionalStream1() (js.Promise[WebTransportBidirectionalStream], bool) {
	var _ok bool
	_ret := bindings.CallWebTransportCreateBidirectionalStream1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[WebTransportBidirectionalStream]{}.FromRef(_ret), _ok
}

// CreateBidirectionalStream1Func returns the method "WebTransport.createBidirectionalStream".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebTransport) CreateBidirectionalStream1Func() (fn js.Func[func() js.Promise[WebTransportBidirectionalStream]]) {
	return fn.FromRef(
		bindings.WebTransportCreateBidirectionalStream1Func(
			this.Ref(),
		),
	)
}

// CreateUnidirectionalStream calls the method "WebTransport.createUnidirectionalStream".
//
// The returned bool will be false if there is no such method.
func (this WebTransport) CreateUnidirectionalStream(options WebTransportSendStreamOptions) (js.Promise[WebTransportSendStream], bool) {
	var _ok bool
	_ret := bindings.CallWebTransportCreateUnidirectionalStream(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[WebTransportSendStream]{}.FromRef(_ret), _ok
}

// CreateUnidirectionalStreamFunc returns the method "WebTransport.createUnidirectionalStream".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebTransport) CreateUnidirectionalStreamFunc() (fn js.Func[func(options WebTransportSendStreamOptions) js.Promise[WebTransportSendStream]]) {
	return fn.FromRef(
		bindings.WebTransportCreateUnidirectionalStreamFunc(
			this.Ref(),
		),
	)
}

// CreateUnidirectionalStream1 calls the method "WebTransport.createUnidirectionalStream".
//
// The returned bool will be false if there is no such method.
func (this WebTransport) CreateUnidirectionalStream1() (js.Promise[WebTransportSendStream], bool) {
	var _ok bool
	_ret := bindings.CallWebTransportCreateUnidirectionalStream1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[WebTransportSendStream]{}.FromRef(_ret), _ok
}

// CreateUnidirectionalStream1Func returns the method "WebTransport.createUnidirectionalStream".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebTransport) CreateUnidirectionalStream1Func() (fn js.Func[func() js.Promise[WebTransportSendStream]]) {
	return fn.FromRef(
		bindings.WebTransportCreateUnidirectionalStream1Func(
			this.Ref(),
		),
	)
}

type WebTransportErrorSource uint32

const (
	_ WebTransportErrorSource = iota

	WebTransportErrorSource_STREAM
	WebTransportErrorSource_SESSION
)

func (WebTransportErrorSource) FromRef(str js.Ref) WebTransportErrorSource {
	return WebTransportErrorSource(bindings.ConstOfWebTransportErrorSource(str))
}

func (x WebTransportErrorSource) String() (string, bool) {
	switch x {
	case WebTransportErrorSource_STREAM:
		return "stream", true
	case WebTransportErrorSource_SESSION:
		return "session", true
	default:
		return "", false
	}
}

type WebTransportErrorOptions struct {
	// Source is "WebTransportErrorOptions.source"
	//
	// Optional, defaults to "stream".
	Source WebTransportErrorSource
	// StreamErrorCode is "WebTransportErrorOptions.streamErrorCode"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_StreamErrorCode MUST be set to true to make this field effective.
	StreamErrorCode uint32

	FFI_USE_StreamErrorCode bool // for StreamErrorCode.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebTransportErrorOptions with all fields set.
func (p WebTransportErrorOptions) FromRef(ref js.Ref) WebTransportErrorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebTransportErrorOptions in the application heap.
func (p WebTransportErrorOptions) New() js.Ref {
	return bindings.WebTransportErrorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WebTransportErrorOptions) UpdateFrom(ref js.Ref) {
	bindings.WebTransportErrorOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WebTransportErrorOptions) Update(ref js.Ref) {
	bindings.WebTransportErrorOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewWebTransportError(message js.String, options WebTransportErrorOptions) WebTransportError {
	return WebTransportError{}.FromRef(
		bindings.NewWebTransportErrorByWebTransportError(
			message.Ref(),
			js.Pointer(&options)),
	)
}

func NewWebTransportErrorByWebTransportError1(message js.String) WebTransportError {
	return WebTransportError{}.FromRef(
		bindings.NewWebTransportErrorByWebTransportError1(
			message.Ref()),
	)
}

func NewWebTransportErrorByWebTransportError2() WebTransportError {
	return WebTransportError{}.FromRef(
		bindings.NewWebTransportErrorByWebTransportError2(),
	)
}

type WebTransportError struct {
	DOMException
}

func (this WebTransportError) Once() WebTransportError {
	this.Ref().Once()
	return this
}

func (this WebTransportError) Ref() js.Ref {
	return this.DOMException.Ref()
}

func (this WebTransportError) FromRef(ref js.Ref) WebTransportError {
	this.DOMException = this.DOMException.FromRef(ref)
	return this
}

func (this WebTransportError) Free() {
	this.Ref().Free()
}

// Source returns the value of property "WebTransportError.source".
//
// The returned bool will be false if there is no such property.
func (this WebTransportError) Source() (WebTransportErrorSource, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportErrorSource(
		this.Ref(), js.Pointer(&_ok),
	)
	return WebTransportErrorSource(_ret), _ok
}

// StreamErrorCode returns the value of property "WebTransportError.streamErrorCode".
//
// The returned bool will be false if there is no such property.
func (this WebTransportError) StreamErrorCode() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetWebTransportErrorStreamErrorCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

const (
	WheelEvent_DOM_DELTA_PIXEL uint32 = 0x00
	WheelEvent_DOM_DELTA_LINE  uint32 = 0x01
	WheelEvent_DOM_DELTA_PAGE  uint32 = 0x02
)

type WheelEventInit struct {
	// DeltaX is "WheelEventInit.deltaX"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_DeltaX MUST be set to true to make this field effective.
	DeltaX float64
	// DeltaY is "WheelEventInit.deltaY"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_DeltaY MUST be set to true to make this field effective.
	DeltaY float64
	// DeltaZ is "WheelEventInit.deltaZ"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_DeltaZ MUST be set to true to make this field effective.
	DeltaZ float64
	// DeltaMode is "WheelEventInit.deltaMode"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_DeltaMode MUST be set to true to make this field effective.
	DeltaMode uint32
	// ScreenX is "WheelEventInit.screenX"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ScreenX MUST be set to true to make this field effective.
	ScreenX int32
	// ScreenY is "WheelEventInit.screenY"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ScreenY MUST be set to true to make this field effective.
	ScreenY int32
	// ClientX is "WheelEventInit.clientX"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ClientX MUST be set to true to make this field effective.
	ClientX int32
	// ClientY is "WheelEventInit.clientY"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ClientY MUST be set to true to make this field effective.
	ClientY int32
	// Button is "WheelEventInit.button"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Button MUST be set to true to make this field effective.
	Button int16
	// Buttons is "WheelEventInit.buttons"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Buttons MUST be set to true to make this field effective.
	Buttons uint16
	// RelatedTarget is "WheelEventInit.relatedTarget"
	//
	// Optional, defaults to null.
	RelatedTarget EventTarget
	// CtrlKey is "WheelEventInit.ctrlKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_CtrlKey MUST be set to true to make this field effective.
	CtrlKey bool
	// ShiftKey is "WheelEventInit.shiftKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ShiftKey MUST be set to true to make this field effective.
	ShiftKey bool
	// AltKey is "WheelEventInit.altKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_AltKey MUST be set to true to make this field effective.
	AltKey bool
	// MetaKey is "WheelEventInit.metaKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_MetaKey MUST be set to true to make this field effective.
	MetaKey bool
	// ModifierAltGraph is "WheelEventInit.modifierAltGraph"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierAltGraph MUST be set to true to make this field effective.
	ModifierAltGraph bool
	// ModifierCapsLock is "WheelEventInit.modifierCapsLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierCapsLock MUST be set to true to make this field effective.
	ModifierCapsLock bool
	// ModifierFn is "WheelEventInit.modifierFn"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierFn MUST be set to true to make this field effective.
	ModifierFn bool
	// ModifierFnLock is "WheelEventInit.modifierFnLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierFnLock MUST be set to true to make this field effective.
	ModifierFnLock bool
	// ModifierHyper is "WheelEventInit.modifierHyper"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierHyper MUST be set to true to make this field effective.
	ModifierHyper bool
	// ModifierNumLock is "WheelEventInit.modifierNumLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierNumLock MUST be set to true to make this field effective.
	ModifierNumLock bool
	// ModifierScrollLock is "WheelEventInit.modifierScrollLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierScrollLock MUST be set to true to make this field effective.
	ModifierScrollLock bool
	// ModifierSuper is "WheelEventInit.modifierSuper"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierSuper MUST be set to true to make this field effective.
	ModifierSuper bool
	// ModifierSymbol is "WheelEventInit.modifierSymbol"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierSymbol MUST be set to true to make this field effective.
	ModifierSymbol bool
	// ModifierSymbolLock is "WheelEventInit.modifierSymbolLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierSymbolLock MUST be set to true to make this field effective.
	ModifierSymbolLock bool
	// View is "WheelEventInit.view"
	//
	// Optional, defaults to null.
	View Window
	// Detail is "WheelEventInit.detail"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Detail MUST be set to true to make this field effective.
	Detail int32
	// Bubbles is "WheelEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "WheelEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "WheelEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_DeltaX             bool // for DeltaX.
	FFI_USE_DeltaY             bool // for DeltaY.
	FFI_USE_DeltaZ             bool // for DeltaZ.
	FFI_USE_DeltaMode          bool // for DeltaMode.
	FFI_USE_ScreenX            bool // for ScreenX.
	FFI_USE_ScreenY            bool // for ScreenY.
	FFI_USE_ClientX            bool // for ClientX.
	FFI_USE_ClientY            bool // for ClientY.
	FFI_USE_Button             bool // for Button.
	FFI_USE_Buttons            bool // for Buttons.
	FFI_USE_CtrlKey            bool // for CtrlKey.
	FFI_USE_ShiftKey           bool // for ShiftKey.
	FFI_USE_AltKey             bool // for AltKey.
	FFI_USE_MetaKey            bool // for MetaKey.
	FFI_USE_ModifierAltGraph   bool // for ModifierAltGraph.
	FFI_USE_ModifierCapsLock   bool // for ModifierCapsLock.
	FFI_USE_ModifierFn         bool // for ModifierFn.
	FFI_USE_ModifierFnLock     bool // for ModifierFnLock.
	FFI_USE_ModifierHyper      bool // for ModifierHyper.
	FFI_USE_ModifierNumLock    bool // for ModifierNumLock.
	FFI_USE_ModifierScrollLock bool // for ModifierScrollLock.
	FFI_USE_ModifierSuper      bool // for ModifierSuper.
	FFI_USE_ModifierSymbol     bool // for ModifierSymbol.
	FFI_USE_ModifierSymbolLock bool // for ModifierSymbolLock.
	FFI_USE_Detail             bool // for Detail.
	FFI_USE_Bubbles            bool // for Bubbles.
	FFI_USE_Cancelable         bool // for Cancelable.
	FFI_USE_Composed           bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WheelEventInit with all fields set.
func (p WheelEventInit) FromRef(ref js.Ref) WheelEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WheelEventInit in the application heap.
func (p WheelEventInit) New() js.Ref {
	return bindings.WheelEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WheelEventInit) UpdateFrom(ref js.Ref) {
	bindings.WheelEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WheelEventInit) Update(ref js.Ref) {
	bindings.WheelEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewWheelEvent(typ js.String, eventInitDict WheelEventInit) WheelEvent {
	return WheelEvent{}.FromRef(
		bindings.NewWheelEventByWheelEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewWheelEventByWheelEvent1(typ js.String) WheelEvent {
	return WheelEvent{}.FromRef(
		bindings.NewWheelEventByWheelEvent1(
			typ.Ref()),
	)
}

type WheelEvent struct {
	MouseEvent
}

func (this WheelEvent) Once() WheelEvent {
	this.Ref().Once()
	return this
}

func (this WheelEvent) Ref() js.Ref {
	return this.MouseEvent.Ref()
}

func (this WheelEvent) FromRef(ref js.Ref) WheelEvent {
	this.MouseEvent = this.MouseEvent.FromRef(ref)
	return this
}

func (this WheelEvent) Free() {
	this.Ref().Free()
}

// DeltaX returns the value of property "WheelEvent.deltaX".
//
// The returned bool will be false if there is no such property.
func (this WheelEvent) DeltaX() (float64, bool) {
	var _ok bool
	_ret := bindings.GetWheelEventDeltaX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// DeltaY returns the value of property "WheelEvent.deltaY".
//
// The returned bool will be false if there is no such property.
func (this WheelEvent) DeltaY() (float64, bool) {
	var _ok bool
	_ret := bindings.GetWheelEventDeltaY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// DeltaZ returns the value of property "WheelEvent.deltaZ".
//
// The returned bool will be false if there is no such property.
func (this WheelEvent) DeltaZ() (float64, bool) {
	var _ok bool
	_ret := bindings.GetWheelEventDeltaZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// DeltaMode returns the value of property "WheelEvent.deltaMode".
//
// The returned bool will be false if there is no such property.
func (this WheelEvent) DeltaMode() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetWheelEventDeltaMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

type WindowControlsOverlayGeometryChangeEventInit struct {
	// TitlebarAreaRect is "WindowControlsOverlayGeometryChangeEventInit.titlebarAreaRect"
	//
	// Required
	TitlebarAreaRect DOMRect
	// Visible is "WindowControlsOverlayGeometryChangeEventInit.visible"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Visible MUST be set to true to make this field effective.
	Visible bool
	// Bubbles is "WindowControlsOverlayGeometryChangeEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "WindowControlsOverlayGeometryChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "WindowControlsOverlayGeometryChangeEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Visible    bool // for Visible.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WindowControlsOverlayGeometryChangeEventInit with all fields set.
func (p WindowControlsOverlayGeometryChangeEventInit) FromRef(ref js.Ref) WindowControlsOverlayGeometryChangeEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WindowControlsOverlayGeometryChangeEventInit in the application heap.
func (p WindowControlsOverlayGeometryChangeEventInit) New() js.Ref {
	return bindings.WindowControlsOverlayGeometryChangeEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WindowControlsOverlayGeometryChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.WindowControlsOverlayGeometryChangeEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WindowControlsOverlayGeometryChangeEventInit) Update(ref js.Ref) {
	bindings.WindowControlsOverlayGeometryChangeEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewWindowControlsOverlayGeometryChangeEvent(typ js.String, eventInitDict WindowControlsOverlayGeometryChangeEventInit) WindowControlsOverlayGeometryChangeEvent {
	return WindowControlsOverlayGeometryChangeEvent{}.FromRef(
		bindings.NewWindowControlsOverlayGeometryChangeEventByWindowControlsOverlayGeometryChangeEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type WindowControlsOverlayGeometryChangeEvent struct {
	Event
}

func (this WindowControlsOverlayGeometryChangeEvent) Once() WindowControlsOverlayGeometryChangeEvent {
	this.Ref().Once()
	return this
}

func (this WindowControlsOverlayGeometryChangeEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this WindowControlsOverlayGeometryChangeEvent) FromRef(ref js.Ref) WindowControlsOverlayGeometryChangeEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this WindowControlsOverlayGeometryChangeEvent) Free() {
	this.Ref().Free()
}

// TitlebarAreaRect returns the value of property "WindowControlsOverlayGeometryChangeEvent.titlebarAreaRect".
//
// The returned bool will be false if there is no such property.
func (this WindowControlsOverlayGeometryChangeEvent) TitlebarAreaRect() (DOMRect, bool) {
	var _ok bool
	_ret := bindings.GetWindowControlsOverlayGeometryChangeEventTitlebarAreaRect(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMRect{}.FromRef(_ret), _ok
}

// Visible returns the value of property "WindowControlsOverlayGeometryChangeEvent.visible".
//
// The returned bool will be false if there is no such property.
func (this WindowControlsOverlayGeometryChangeEvent) Visible() (bool, bool) {
	var _ok bool
	_ret := bindings.GetWindowControlsOverlayGeometryChangeEventVisible(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

type WorkerLocation struct {
	ref js.Ref
}

func (this WorkerLocation) Once() WorkerLocation {
	this.Ref().Once()
	return this
}

func (this WorkerLocation) Ref() js.Ref {
	return this.ref
}

func (this WorkerLocation) FromRef(ref js.Ref) WorkerLocation {
	this.ref = ref
	return this
}

func (this WorkerLocation) Free() {
	this.Ref().Free()
}

// Href returns the value of property "WorkerLocation.href".
//
// The returned bool will be false if there is no such property.
func (this WorkerLocation) Href() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerLocationHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Origin returns the value of property "WorkerLocation.origin".
//
// The returned bool will be false if there is no such property.
func (this WorkerLocation) Origin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerLocationOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Protocol returns the value of property "WorkerLocation.protocol".
//
// The returned bool will be false if there is no such property.
func (this WorkerLocation) Protocol() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerLocationProtocol(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Host returns the value of property "WorkerLocation.host".
//
// The returned bool will be false if there is no such property.
func (this WorkerLocation) Host() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerLocationHost(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hostname returns the value of property "WorkerLocation.hostname".
//
// The returned bool will be false if there is no such property.
func (this WorkerLocation) Hostname() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerLocationHostname(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Port returns the value of property "WorkerLocation.port".
//
// The returned bool will be false if there is no such property.
func (this WorkerLocation) Port() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerLocationPort(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Pathname returns the value of property "WorkerLocation.pathname".
//
// The returned bool will be false if there is no such property.
func (this WorkerLocation) Pathname() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerLocationPathname(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Search returns the value of property "WorkerLocation.search".
//
// The returned bool will be false if there is no such property.
func (this WorkerLocation) Search() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerLocationSearch(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hash returns the value of property "WorkerLocation.hash".
//
// The returned bool will be false if there is no such property.
func (this WorkerLocation) Hash() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerLocationHash(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type WorkerNavigator struct {
	ref js.Ref
}

func (this WorkerNavigator) Once() WorkerNavigator {
	this.Ref().Once()
	return this
}

func (this WorkerNavigator) Ref() js.Ref {
	return this.ref
}

func (this WorkerNavigator) FromRef(ref js.Ref) WorkerNavigator {
	this.ref = ref
	return this
}

func (this WorkerNavigator) Free() {
	this.Ref().Free()
}

// Hid returns the value of property "WorkerNavigator.hid".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) Hid() (HID, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorHid(
		this.Ref(), js.Pointer(&_ok),
	)
	return HID{}.FromRef(_ret), _ok
}

// Serial returns the value of property "WorkerNavigator.serial".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) Serial() (Serial, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorSerial(
		this.Ref(), js.Pointer(&_ok),
	)
	return Serial{}.FromRef(_ret), _ok
}

// Permissions returns the value of property "WorkerNavigator.permissions".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) Permissions() (Permissions, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorPermissions(
		this.Ref(), js.Pointer(&_ok),
	)
	return Permissions{}.FromRef(_ret), _ok
}

// Usb returns the value of property "WorkerNavigator.usb".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) Usb() (USB, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorUsb(
		this.Ref(), js.Pointer(&_ok),
	)
	return USB{}.FromRef(_ret), _ok
}

// MediaCapabilities returns the value of property "WorkerNavigator.mediaCapabilities".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) MediaCapabilities() (MediaCapabilities, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorMediaCapabilities(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaCapabilities{}.FromRef(_ret), _ok
}

// ServiceWorker returns the value of property "WorkerNavigator.serviceWorker".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) ServiceWorker() (ServiceWorkerContainer, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorServiceWorker(
		this.Ref(), js.Pointer(&_ok),
	)
	return ServiceWorkerContainer{}.FromRef(_ret), _ok
}

// UserAgentData returns the value of property "WorkerNavigator.userAgentData".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) UserAgentData() (NavigatorUAData, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorUserAgentData(
		this.Ref(), js.Pointer(&_ok),
	)
	return NavigatorUAData{}.FromRef(_ret), _ok
}

// DeviceMemory returns the value of property "WorkerNavigator.deviceMemory".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) DeviceMemory() (float64, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorDeviceMemory(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Storage returns the value of property "WorkerNavigator.storage".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) Storage() (StorageManager, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorStorage(
		this.Ref(), js.Pointer(&_ok),
	)
	return StorageManager{}.FromRef(_ret), _ok
}

// StorageBuckets returns the value of property "WorkerNavigator.storageBuckets".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) StorageBuckets() (StorageBucketManager, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorStorageBuckets(
		this.Ref(), js.Pointer(&_ok),
	)
	return StorageBucketManager{}.FromRef(_ret), _ok
}

// Locks returns the value of property "WorkerNavigator.locks".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) Locks() (LockManager, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorLocks(
		this.Ref(), js.Pointer(&_ok),
	)
	return LockManager{}.FromRef(_ret), _ok
}

// AppCodeName returns the value of property "WorkerNavigator.appCodeName".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) AppCodeName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorAppCodeName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AppName returns the value of property "WorkerNavigator.appName".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) AppName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorAppName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AppVersion returns the value of property "WorkerNavigator.appVersion".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) AppVersion() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorAppVersion(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Platform returns the value of property "WorkerNavigator.platform".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) Platform() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorPlatform(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Product returns the value of property "WorkerNavigator.product".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) Product() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorProduct(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ProductSub returns the value of property "WorkerNavigator.productSub".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) ProductSub() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorProductSub(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// UserAgent returns the value of property "WorkerNavigator.userAgent".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) UserAgent() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorUserAgent(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Vendor returns the value of property "WorkerNavigator.vendor".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) Vendor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorVendor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// VendorSub returns the value of property "WorkerNavigator.vendorSub".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) VendorSub() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorVendorSub(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Oscpu returns the value of property "WorkerNavigator.oscpu".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) Oscpu() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorOscpu(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Language returns the value of property "WorkerNavigator.language".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) Language() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorLanguage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Languages returns the value of property "WorkerNavigator.languages".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) Languages() (js.FrozenArray[js.String], bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorLanguages(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[js.String]{}.FromRef(_ret), _ok
}

// OnLine returns the value of property "WorkerNavigator.onLine".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) OnLine() (bool, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorOnLine(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// HardwareConcurrency returns the value of property "WorkerNavigator.hardwareConcurrency".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) HardwareConcurrency() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorHardwareConcurrency(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// Connection returns the value of property "WorkerNavigator.connection".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) Connection() (NetworkInformation, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorConnection(
		this.Ref(), js.Pointer(&_ok),
	)
	return NetworkInformation{}.FromRef(_ret), _ok
}

// Ml returns the value of property "WorkerNavigator.ml".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) Ml() (ML, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorMl(
		this.Ref(), js.Pointer(&_ok),
	)
	return ML{}.FromRef(_ret), _ok
}

// Gpu returns the value of property "WorkerNavigator.gpu".
//
// The returned bool will be false if there is no such property.
func (this WorkerNavigator) Gpu() (GPU, bool) {
	var _ok bool
	_ret := bindings.GetWorkerNavigatorGpu(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPU{}.FromRef(_ret), _ok
}

// TaintEnabled calls the method "WorkerNavigator.taintEnabled".
//
// The returned bool will be false if there is no such method.
func (this WorkerNavigator) TaintEnabled() (bool, bool) {
	var _ok bool
	_ret := bindings.CallWorkerNavigatorTaintEnabled(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// TaintEnabledFunc returns the method "WorkerNavigator.taintEnabled".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerNavigator) TaintEnabledFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.WorkerNavigatorTaintEnabledFunc(
			this.Ref(),
		),
	)
}

// SetAppBadge calls the method "WorkerNavigator.setAppBadge".
//
// The returned bool will be false if there is no such method.
func (this WorkerNavigator) SetAppBadge(contents uint64) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWorkerNavigatorSetAppBadge(
		this.Ref(), js.Pointer(&_ok),
		float64(contents),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetAppBadgeFunc returns the method "WorkerNavigator.setAppBadge".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerNavigator) SetAppBadgeFunc() (fn js.Func[func(contents uint64) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WorkerNavigatorSetAppBadgeFunc(
			this.Ref(),
		),
	)
}

// SetAppBadge1 calls the method "WorkerNavigator.setAppBadge".
//
// The returned bool will be false if there is no such method.
func (this WorkerNavigator) SetAppBadge1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWorkerNavigatorSetAppBadge1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetAppBadge1Func returns the method "WorkerNavigator.setAppBadge".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerNavigator) SetAppBadge1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WorkerNavigatorSetAppBadge1Func(
			this.Ref(),
		),
	)
}

// ClearAppBadge calls the method "WorkerNavigator.clearAppBadge".
//
// The returned bool will be false if there is no such method.
func (this WorkerNavigator) ClearAppBadge() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWorkerNavigatorClearAppBadge(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ClearAppBadgeFunc returns the method "WorkerNavigator.clearAppBadge".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerNavigator) ClearAppBadgeFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WorkerNavigatorClearAppBadgeFunc(
			this.Ref(),
		),
	)
}

type WorkerGlobalScope struct {
	EventTarget
}

func (this WorkerGlobalScope) Once() WorkerGlobalScope {
	this.Ref().Once()
	return this
}

func (this WorkerGlobalScope) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this WorkerGlobalScope) FromRef(ref js.Ref) WorkerGlobalScope {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this WorkerGlobalScope) Free() {
	this.Ref().Free()
}

// Self returns the value of property "WorkerGlobalScope.self".
//
// The returned bool will be false if there is no such property.
func (this WorkerGlobalScope) Self() (WorkerGlobalScope, bool) {
	var _ok bool
	_ret := bindings.GetWorkerGlobalScopeSelf(
		this.Ref(), js.Pointer(&_ok),
	)
	return WorkerGlobalScope{}.FromRef(_ret), _ok
}

// Location returns the value of property "WorkerGlobalScope.location".
//
// The returned bool will be false if there is no such property.
func (this WorkerGlobalScope) Location() (WorkerLocation, bool) {
	var _ok bool
	_ret := bindings.GetWorkerGlobalScopeLocation(
		this.Ref(), js.Pointer(&_ok),
	)
	return WorkerLocation{}.FromRef(_ret), _ok
}

// Navigator returns the value of property "WorkerGlobalScope.navigator".
//
// The returned bool will be false if there is no such property.
func (this WorkerGlobalScope) Navigator() (WorkerNavigator, bool) {
	var _ok bool
	_ret := bindings.GetWorkerGlobalScopeNavigator(
		this.Ref(), js.Pointer(&_ok),
	)
	return WorkerNavigator{}.FromRef(_ret), _ok
}

// Fonts returns the value of property "WorkerGlobalScope.fonts".
//
// The returned bool will be false if there is no such property.
func (this WorkerGlobalScope) Fonts() (FontFaceSet, bool) {
	var _ok bool
	_ret := bindings.GetWorkerGlobalScopeFonts(
		this.Ref(), js.Pointer(&_ok),
	)
	return FontFaceSet{}.FromRef(_ret), _ok
}

// Origin returns the value of property "WorkerGlobalScope.origin".
//
// The returned bool will be false if there is no such property.
func (this WorkerGlobalScope) Origin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkerGlobalScopeOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// IsSecureContext returns the value of property "WorkerGlobalScope.isSecureContext".
//
// The returned bool will be false if there is no such property.
func (this WorkerGlobalScope) IsSecureContext() (bool, bool) {
	var _ok bool
	_ret := bindings.GetWorkerGlobalScopeIsSecureContext(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// CrossOriginIsolated returns the value of property "WorkerGlobalScope.crossOriginIsolated".
//
// The returned bool will be false if there is no such property.
func (this WorkerGlobalScope) CrossOriginIsolated() (bool, bool) {
	var _ok bool
	_ret := bindings.GetWorkerGlobalScopeCrossOriginIsolated(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Scheduler returns the value of property "WorkerGlobalScope.scheduler".
//
// The returned bool will be false if there is no such property.
func (this WorkerGlobalScope) Scheduler() (Scheduler, bool) {
	var _ok bool
	_ret := bindings.GetWorkerGlobalScopeScheduler(
		this.Ref(), js.Pointer(&_ok),
	)
	return Scheduler{}.FromRef(_ret), _ok
}

// TrustedTypes returns the value of property "WorkerGlobalScope.trustedTypes".
//
// The returned bool will be false if there is no such property.
func (this WorkerGlobalScope) TrustedTypes() (TrustedTypePolicyFactory, bool) {
	var _ok bool
	_ret := bindings.GetWorkerGlobalScopeTrustedTypes(
		this.Ref(), js.Pointer(&_ok),
	)
	return TrustedTypePolicyFactory{}.FromRef(_ret), _ok
}

// Caches returns the value of property "WorkerGlobalScope.caches".
//
// The returned bool will be false if there is no such property.
func (this WorkerGlobalScope) Caches() (CacheStorage, bool) {
	var _ok bool
	_ret := bindings.GetWorkerGlobalScopeCaches(
		this.Ref(), js.Pointer(&_ok),
	)
	return CacheStorage{}.FromRef(_ret), _ok
}

// Crypto returns the value of property "WorkerGlobalScope.crypto".
//
// The returned bool will be false if there is no such property.
func (this WorkerGlobalScope) Crypto() (Crypto, bool) {
	var _ok bool
	_ret := bindings.GetWorkerGlobalScopeCrypto(
		this.Ref(), js.Pointer(&_ok),
	)
	return Crypto{}.FromRef(_ret), _ok
}

// IndexedDB returns the value of property "WorkerGlobalScope.indexedDB".
//
// The returned bool will be false if there is no such property.
func (this WorkerGlobalScope) IndexedDB() (IDBFactory, bool) {
	var _ok bool
	_ret := bindings.GetWorkerGlobalScopeIndexedDB(
		this.Ref(), js.Pointer(&_ok),
	)
	return IDBFactory{}.FromRef(_ret), _ok
}

// Performance returns the value of property "WorkerGlobalScope.performance".
//
// The returned bool will be false if there is no such property.
func (this WorkerGlobalScope) Performance() (Performance, bool) {
	var _ok bool
	_ret := bindings.GetWorkerGlobalScopePerformance(
		this.Ref(), js.Pointer(&_ok),
	)
	return Performance{}.FromRef(_ret), _ok
}

// ImportScripts calls the method "WorkerGlobalScope.importScripts".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) ImportScripts(urls ...js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeImportScripts(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(urls),
		js.SizeU(len(urls)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ImportScriptsFunc returns the method "WorkerGlobalScope.importScripts".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) ImportScriptsFunc() (fn js.Func[func(urls ...js.String)]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeImportScriptsFunc(
			this.Ref(),
		),
	)
}

// ReportError calls the method "WorkerGlobalScope.reportError".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) ReportError(e js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeReportError(
		this.Ref(), js.Pointer(&_ok),
		e.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReportErrorFunc returns the method "WorkerGlobalScope.reportError".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) ReportErrorFunc() (fn js.Func[func(e js.Any)]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeReportErrorFunc(
			this.Ref(),
		),
	)
}

// Btoa calls the method "WorkerGlobalScope.btoa".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) Btoa(data js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeBtoa(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// BtoaFunc returns the method "WorkerGlobalScope.btoa".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) BtoaFunc() (fn js.Func[func(data js.String) js.String]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeBtoaFunc(
			this.Ref(),
		),
	)
}

// Atob calls the method "WorkerGlobalScope.atob".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) Atob(data js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeAtob(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// AtobFunc returns the method "WorkerGlobalScope.atob".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) AtobFunc() (fn js.Func[func(data js.String) js.String]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeAtobFunc(
			this.Ref(),
		),
	)
}

// SetTimeout calls the method "WorkerGlobalScope.setTimeout".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) SetTimeout(handler TimerHandler, timeout int32, arguments ...js.Any) (int32, bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeSetTimeout(
		this.Ref(), js.Pointer(&_ok),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return int32(_ret), _ok
}

// SetTimeoutFunc returns the method "WorkerGlobalScope.setTimeout".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) SetTimeoutFunc() (fn js.Func[func(handler TimerHandler, timeout int32, arguments ...js.Any) int32]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeSetTimeoutFunc(
			this.Ref(),
		),
	)
}

// SetTimeout1 calls the method "WorkerGlobalScope.setTimeout".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) SetTimeout1(handler TimerHandler) (int32, bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeSetTimeout1(
		this.Ref(), js.Pointer(&_ok),
		handler.Ref(),
	)

	return int32(_ret), _ok
}

// SetTimeout1Func returns the method "WorkerGlobalScope.setTimeout".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) SetTimeout1Func() (fn js.Func[func(handler TimerHandler) int32]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeSetTimeout1Func(
			this.Ref(),
		),
	)
}

// ClearTimeout calls the method "WorkerGlobalScope.clearTimeout".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) ClearTimeout(id int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeClearTimeout(
		this.Ref(), js.Pointer(&_ok),
		int32(id),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearTimeoutFunc returns the method "WorkerGlobalScope.clearTimeout".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) ClearTimeoutFunc() (fn js.Func[func(id int32)]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeClearTimeoutFunc(
			this.Ref(),
		),
	)
}

// ClearTimeout1 calls the method "WorkerGlobalScope.clearTimeout".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) ClearTimeout1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeClearTimeout1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearTimeout1Func returns the method "WorkerGlobalScope.clearTimeout".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) ClearTimeout1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeClearTimeout1Func(
			this.Ref(),
		),
	)
}

// SetInterval calls the method "WorkerGlobalScope.setInterval".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) SetInterval(handler TimerHandler, timeout int32, arguments ...js.Any) (int32, bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeSetInterval(
		this.Ref(), js.Pointer(&_ok),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return int32(_ret), _ok
}

// SetIntervalFunc returns the method "WorkerGlobalScope.setInterval".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) SetIntervalFunc() (fn js.Func[func(handler TimerHandler, timeout int32, arguments ...js.Any) int32]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeSetIntervalFunc(
			this.Ref(),
		),
	)
}

// SetInterval1 calls the method "WorkerGlobalScope.setInterval".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) SetInterval1(handler TimerHandler) (int32, bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeSetInterval1(
		this.Ref(), js.Pointer(&_ok),
		handler.Ref(),
	)

	return int32(_ret), _ok
}

// SetInterval1Func returns the method "WorkerGlobalScope.setInterval".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) SetInterval1Func() (fn js.Func[func(handler TimerHandler) int32]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeSetInterval1Func(
			this.Ref(),
		),
	)
}

// ClearInterval calls the method "WorkerGlobalScope.clearInterval".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) ClearInterval(id int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeClearInterval(
		this.Ref(), js.Pointer(&_ok),
		int32(id),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearIntervalFunc returns the method "WorkerGlobalScope.clearInterval".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) ClearIntervalFunc() (fn js.Func[func(id int32)]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeClearIntervalFunc(
			this.Ref(),
		),
	)
}

// ClearInterval1 calls the method "WorkerGlobalScope.clearInterval".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) ClearInterval1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeClearInterval1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearInterval1Func returns the method "WorkerGlobalScope.clearInterval".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) ClearInterval1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeClearInterval1Func(
			this.Ref(),
		),
	)
}

// QueueMicrotask calls the method "WorkerGlobalScope.queueMicrotask".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) QueueMicrotask(callback js.Func[func()]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeQueueMicrotask(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// QueueMicrotaskFunc returns the method "WorkerGlobalScope.queueMicrotask".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) QueueMicrotaskFunc() (fn js.Func[func(callback js.Func[func()])]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeQueueMicrotaskFunc(
			this.Ref(),
		),
	)
}

// CreateImageBitmap calls the method "WorkerGlobalScope.createImageBitmap".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) CreateImageBitmap(image ImageBitmapSource, options ImageBitmapOptions) (js.Promise[ImageBitmap], bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeCreateImageBitmap(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[ImageBitmap]{}.FromRef(_ret), _ok
}

// CreateImageBitmapFunc returns the method "WorkerGlobalScope.createImageBitmap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) CreateImageBitmapFunc() (fn js.Func[func(image ImageBitmapSource, options ImageBitmapOptions) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeCreateImageBitmapFunc(
			this.Ref(),
		),
	)
}

// CreateImageBitmap1 calls the method "WorkerGlobalScope.createImageBitmap".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) CreateImageBitmap1(image ImageBitmapSource) (js.Promise[ImageBitmap], bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeCreateImageBitmap1(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
	)

	return js.Promise[ImageBitmap]{}.FromRef(_ret), _ok
}

// CreateImageBitmap1Func returns the method "WorkerGlobalScope.createImageBitmap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) CreateImageBitmap1Func() (fn js.Func[func(image ImageBitmapSource) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeCreateImageBitmap1Func(
			this.Ref(),
		),
	)
}

// CreateImageBitmap2 calls the method "WorkerGlobalScope.createImageBitmap".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) CreateImageBitmap2(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32, options ImageBitmapOptions) (js.Promise[ImageBitmap], bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeCreateImageBitmap2(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
		js.Pointer(&options),
	)

	return js.Promise[ImageBitmap]{}.FromRef(_ret), _ok
}

// CreateImageBitmap2Func returns the method "WorkerGlobalScope.createImageBitmap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) CreateImageBitmap2Func() (fn js.Func[func(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32, options ImageBitmapOptions) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeCreateImageBitmap2Func(
			this.Ref(),
		),
	)
}

// CreateImageBitmap3 calls the method "WorkerGlobalScope.createImageBitmap".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) CreateImageBitmap3(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32) (js.Promise[ImageBitmap], bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeCreateImageBitmap3(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
	)

	return js.Promise[ImageBitmap]{}.FromRef(_ret), _ok
}

// CreateImageBitmap3Func returns the method "WorkerGlobalScope.createImageBitmap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) CreateImageBitmap3Func() (fn js.Func[func(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeCreateImageBitmap3Func(
			this.Ref(),
		),
	)
}

// StructuredClone calls the method "WorkerGlobalScope.structuredClone".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) StructuredClone(value js.Any, options StructuredSerializeOptions) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeStructuredClone(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
		js.Pointer(&options),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// StructuredCloneFunc returns the method "WorkerGlobalScope.structuredClone".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) StructuredCloneFunc() (fn js.Func[func(value js.Any, options StructuredSerializeOptions) js.Any]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeStructuredCloneFunc(
			this.Ref(),
		),
	)
}

// StructuredClone1 calls the method "WorkerGlobalScope.structuredClone".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) StructuredClone1(value js.Any) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeStructuredClone1(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// StructuredClone1Func returns the method "WorkerGlobalScope.structuredClone".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) StructuredClone1Func() (fn js.Func[func(value js.Any) js.Any]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeStructuredClone1Func(
			this.Ref(),
		),
	)
}

// Fetch calls the method "WorkerGlobalScope.fetch".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) Fetch(input RequestInfo, init RequestInit) (js.Promise[Response], bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeFetch(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&init),
	)

	return js.Promise[Response]{}.FromRef(_ret), _ok
}

// FetchFunc returns the method "WorkerGlobalScope.fetch".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) FetchFunc() (fn js.Func[func(input RequestInfo, init RequestInit) js.Promise[Response]]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeFetchFunc(
			this.Ref(),
		),
	)
}

// Fetch1 calls the method "WorkerGlobalScope.fetch".
//
// The returned bool will be false if there is no such method.
func (this WorkerGlobalScope) Fetch1(input RequestInfo) (js.Promise[Response], bool) {
	var _ok bool
	_ret := bindings.CallWorkerGlobalScopeFetch1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return js.Promise[Response]{}.FromRef(_ret), _ok
}

// Fetch1Func returns the method "WorkerGlobalScope.fetch".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkerGlobalScope) Fetch1Func() (fn js.Func[func(input RequestInfo) js.Promise[Response]]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeFetch1Func(
			this.Ref(),
		),
	)
}

type OneOf_AnimationEffect_ArrayAnimationEffect struct {
	ref js.Ref
}

func (x OneOf_AnimationEffect_ArrayAnimationEffect) Ref() js.Ref {
	return x.ref
}

func (x OneOf_AnimationEffect_ArrayAnimationEffect) Free() {
	x.ref.Free()
}

func (x OneOf_AnimationEffect_ArrayAnimationEffect) FromRef(ref js.Ref) OneOf_AnimationEffect_ArrayAnimationEffect {
	return OneOf_AnimationEffect_ArrayAnimationEffect{
		ref: ref,
	}
}

func (x OneOf_AnimationEffect_ArrayAnimationEffect) AnimationEffect() AnimationEffect {
	return AnimationEffect{}.FromRef(x.ref)
}

func (x OneOf_AnimationEffect_ArrayAnimationEffect) ArrayAnimationEffect() js.Array[AnimationEffect] {
	return js.Array[AnimationEffect]{}.FromRef(x.ref)
}

func NewWorkletAnimation(animatorName js.String, effects OneOf_AnimationEffect_ArrayAnimationEffect, timeline AnimationTimeline, options js.Any) WorkletAnimation {
	return WorkletAnimation{}.FromRef(
		bindings.NewWorkletAnimationByWorkletAnimation(
			animatorName.Ref(),
			effects.Ref(),
			timeline.Ref(),
			options.Ref()),
	)
}

func NewWorkletAnimationByWorkletAnimation1(animatorName js.String, effects OneOf_AnimationEffect_ArrayAnimationEffect, timeline AnimationTimeline) WorkletAnimation {
	return WorkletAnimation{}.FromRef(
		bindings.NewWorkletAnimationByWorkletAnimation1(
			animatorName.Ref(),
			effects.Ref(),
			timeline.Ref()),
	)
}

func NewWorkletAnimationByWorkletAnimation2(animatorName js.String, effects OneOf_AnimationEffect_ArrayAnimationEffect) WorkletAnimation {
	return WorkletAnimation{}.FromRef(
		bindings.NewWorkletAnimationByWorkletAnimation2(
			animatorName.Ref(),
			effects.Ref()),
	)
}

func NewWorkletAnimationByWorkletAnimation3(animatorName js.String) WorkletAnimation {
	return WorkletAnimation{}.FromRef(
		bindings.NewWorkletAnimationByWorkletAnimation3(
			animatorName.Ref()),
	)
}

type WorkletAnimation struct {
	Animation
}

func (this WorkletAnimation) Once() WorkletAnimation {
	this.Ref().Once()
	return this
}

func (this WorkletAnimation) Ref() js.Ref {
	return this.Animation.Ref()
}

func (this WorkletAnimation) FromRef(ref js.Ref) WorkletAnimation {
	this.Animation = this.Animation.FromRef(ref)
	return this
}

func (this WorkletAnimation) Free() {
	this.Ref().Free()
}

// AnimatorName returns the value of property "WorkletAnimation.animatorName".
//
// The returned bool will be false if there is no such property.
func (this WorkletAnimation) AnimatorName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWorkletAnimationAnimatorName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type WorkletAnimationEffect struct {
	ref js.Ref
}

func (this WorkletAnimationEffect) Once() WorkletAnimationEffect {
	this.Ref().Once()
	return this
}

func (this WorkletAnimationEffect) Ref() js.Ref {
	return this.ref
}

func (this WorkletAnimationEffect) FromRef(ref js.Ref) WorkletAnimationEffect {
	this.ref = ref
	return this
}

func (this WorkletAnimationEffect) Free() {
	this.Ref().Free()
}

// LocalTime returns the value of property "WorkletAnimationEffect.localTime".
//
// The returned bool will be false if there is no such property.
func (this WorkletAnimationEffect) LocalTime() (float64, bool) {
	var _ok bool
	_ret := bindings.GetWorkletAnimationEffectLocalTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// LocalTime sets the value of property "WorkletAnimationEffect.localTime" to val.
//
// It returns false if the property cannot be set.
func (this WorkletAnimationEffect) SetLocalTime(val float64) bool {
	return js.True == bindings.SetWorkletAnimationEffectLocalTime(
		this.Ref(),
		float64(val),
	)
}

// GetTiming calls the method "WorkletAnimationEffect.getTiming".
//
// The returned bool will be false if there is no such method.
func (this WorkletAnimationEffect) GetTiming() (EffectTiming, bool) {
	var _ret EffectTiming
	_ok := js.True == bindings.CallWorkletAnimationEffectGetTiming(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetTimingFunc returns the method "WorkletAnimationEffect.getTiming".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkletAnimationEffect) GetTimingFunc() (fn js.Func[func() EffectTiming]) {
	return fn.FromRef(
		bindings.WorkletAnimationEffectGetTimingFunc(
			this.Ref(),
		),
	)
}

// GetComputedTiming calls the method "WorkletAnimationEffect.getComputedTiming".
//
// The returned bool will be false if there is no such method.
func (this WorkletAnimationEffect) GetComputedTiming() (ComputedEffectTiming, bool) {
	var _ret ComputedEffectTiming
	_ok := js.True == bindings.CallWorkletAnimationEffectGetComputedTiming(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetComputedTimingFunc returns the method "WorkletAnimationEffect.getComputedTiming".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkletAnimationEffect) GetComputedTimingFunc() (fn js.Func[func() ComputedEffectTiming]) {
	return fn.FromRef(
		bindings.WorkletAnimationEffectGetComputedTimingFunc(
			this.Ref(),
		),
	)
}

type WorkletGlobalScope struct {
	ref js.Ref
}

func (this WorkletGlobalScope) Once() WorkletGlobalScope {
	this.Ref().Once()
	return this
}

func (this WorkletGlobalScope) Ref() js.Ref {
	return this.ref
}

func (this WorkletGlobalScope) FromRef(ref js.Ref) WorkletGlobalScope {
	this.ref = ref
	return this
}

func (this WorkletGlobalScope) Free() {
	this.Ref().Free()
}

type WorkletGroupEffect struct {
	ref js.Ref
}

func (this WorkletGroupEffect) Once() WorkletGroupEffect {
	this.Ref().Once()
	return this
}

func (this WorkletGroupEffect) Ref() js.Ref {
	return this.ref
}

func (this WorkletGroupEffect) FromRef(ref js.Ref) WorkletGroupEffect {
	this.ref = ref
	return this
}

func (this WorkletGroupEffect) Free() {
	this.Ref().Free()
}

// GetChildren calls the method "WorkletGroupEffect.getChildren".
//
// The returned bool will be false if there is no such method.
func (this WorkletGroupEffect) GetChildren() (js.Array[WorkletAnimationEffect], bool) {
	var _ok bool
	_ret := bindings.CallWorkletGroupEffectGetChildren(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[WorkletAnimationEffect]{}.FromRef(_ret), _ok
}

// GetChildrenFunc returns the method "WorkletGroupEffect.getChildren".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkletGroupEffect) GetChildrenFunc() (fn js.Func[func() js.Array[WorkletAnimationEffect]]) {
	return fn.FromRef(
		bindings.WorkletGroupEffectGetChildrenFunc(
			this.Ref(),
		),
	)
}

const (
	XMLHttpRequest_UNSENT           uint16 = 0
	XMLHttpRequest_OPENED           uint16 = 1
	XMLHttpRequest_HEADERS_RECEIVED uint16 = 2
	XMLHttpRequest_LOADING          uint16 = 3
	XMLHttpRequest_DONE             uint16 = 4
)

type OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String struct {
	ref js.Ref
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) Free() {
	x.ref.Free()
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) FromRef(ref js.Ref) OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String {
	return OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String{
		ref: ref,
	}
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) Document() Document {
	return Document{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) Blob() Blob {
	return Blob{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) ArrayBuffer() js.ArrayBuffer {
	return js.ArrayBuffer{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) FormData() FormData {
	return FormData{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) URLSearchParams() URLSearchParams {
	return URLSearchParams{}.FromRef(x.ref)
}

func (x OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type XMLHttpRequestUpload struct {
	XMLHttpRequestEventTarget
}

func (this XMLHttpRequestUpload) Once() XMLHttpRequestUpload {
	this.Ref().Once()
	return this
}

func (this XMLHttpRequestUpload) Ref() js.Ref {
	return this.XMLHttpRequestEventTarget.Ref()
}

func (this XMLHttpRequestUpload) FromRef(ref js.Ref) XMLHttpRequestUpload {
	this.XMLHttpRequestEventTarget = this.XMLHttpRequestEventTarget.FromRef(ref)
	return this
}

func (this XMLHttpRequestUpload) Free() {
	this.Ref().Free()
}

type XMLHttpRequestResponseType uint32

const (
	_ XMLHttpRequestResponseType = iota

	XMLHttpRequestResponseType_
	XMLHttpRequestResponseType_ARRAYBUFFER
	XMLHttpRequestResponseType_BLOB
	XMLHttpRequestResponseType_DOCUMENT
	XMLHttpRequestResponseType_JSON
	XMLHttpRequestResponseType_TEXT
)

func (XMLHttpRequestResponseType) FromRef(str js.Ref) XMLHttpRequestResponseType {
	return XMLHttpRequestResponseType(bindings.ConstOfXMLHttpRequestResponseType(str))
}

func (x XMLHttpRequestResponseType) String() (string, bool) {
	switch x {
	case XMLHttpRequestResponseType_:
		return "", true
	case XMLHttpRequestResponseType_ARRAYBUFFER:
		return "arraybuffer", true
	case XMLHttpRequestResponseType_BLOB:
		return "blob", true
	case XMLHttpRequestResponseType_DOCUMENT:
		return "document", true
	case XMLHttpRequestResponseType_JSON:
		return "json", true
	case XMLHttpRequestResponseType_TEXT:
		return "text", true
	default:
		return "", false
	}
}

type XMLHttpRequest struct {
	XMLHttpRequestEventTarget
}

func (this XMLHttpRequest) Once() XMLHttpRequest {
	this.Ref().Once()
	return this
}

func (this XMLHttpRequest) Ref() js.Ref {
	return this.XMLHttpRequestEventTarget.Ref()
}

func (this XMLHttpRequest) FromRef(ref js.Ref) XMLHttpRequest {
	this.XMLHttpRequestEventTarget = this.XMLHttpRequestEventTarget.FromRef(ref)
	return this
}

func (this XMLHttpRequest) Free() {
	this.Ref().Free()
}

// ReadyState returns the value of property "XMLHttpRequest.readyState".
//
// The returned bool will be false if there is no such property.
func (this XMLHttpRequest) ReadyState() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetXMLHttpRequestReadyState(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Timeout returns the value of property "XMLHttpRequest.timeout".
//
// The returned bool will be false if there is no such property.
func (this XMLHttpRequest) Timeout() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXMLHttpRequestTimeout(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Timeout sets the value of property "XMLHttpRequest.timeout" to val.
//
// It returns false if the property cannot be set.
func (this XMLHttpRequest) SetTimeout(val uint32) bool {
	return js.True == bindings.SetXMLHttpRequestTimeout(
		this.Ref(),
		uint32(val),
	)
}

// WithCredentials returns the value of property "XMLHttpRequest.withCredentials".
//
// The returned bool will be false if there is no such property.
func (this XMLHttpRequest) WithCredentials() (bool, bool) {
	var _ok bool
	_ret := bindings.GetXMLHttpRequestWithCredentials(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// WithCredentials sets the value of property "XMLHttpRequest.withCredentials" to val.
//
// It returns false if the property cannot be set.
func (this XMLHttpRequest) SetWithCredentials(val bool) bool {
	return js.True == bindings.SetXMLHttpRequestWithCredentials(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Upload returns the value of property "XMLHttpRequest.upload".
//
// The returned bool will be false if there is no such property.
func (this XMLHttpRequest) Upload() (XMLHttpRequestUpload, bool) {
	var _ok bool
	_ret := bindings.GetXMLHttpRequestUpload(
		this.Ref(), js.Pointer(&_ok),
	)
	return XMLHttpRequestUpload{}.FromRef(_ret), _ok
}

// ResponseURL returns the value of property "XMLHttpRequest.responseURL".
//
// The returned bool will be false if there is no such property.
func (this XMLHttpRequest) ResponseURL() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetXMLHttpRequestResponseURL(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Status returns the value of property "XMLHttpRequest.status".
//
// The returned bool will be false if there is no such property.
func (this XMLHttpRequest) Status() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetXMLHttpRequestStatus(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// StatusText returns the value of property "XMLHttpRequest.statusText".
//
// The returned bool will be false if there is no such property.
func (this XMLHttpRequest) StatusText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetXMLHttpRequestStatusText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ResponseType returns the value of property "XMLHttpRequest.responseType".
//
// The returned bool will be false if there is no such property.
func (this XMLHttpRequest) ResponseType() (XMLHttpRequestResponseType, bool) {
	var _ok bool
	_ret := bindings.GetXMLHttpRequestResponseType(
		this.Ref(), js.Pointer(&_ok),
	)
	return XMLHttpRequestResponseType(_ret), _ok
}

// ResponseType sets the value of property "XMLHttpRequest.responseType" to val.
//
// It returns false if the property cannot be set.
func (this XMLHttpRequest) SetResponseType(val XMLHttpRequestResponseType) bool {
	return js.True == bindings.SetXMLHttpRequestResponseType(
		this.Ref(),
		uint32(val),
	)
}

// Response returns the value of property "XMLHttpRequest.response".
//
// The returned bool will be false if there is no such property.
func (this XMLHttpRequest) Response() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetXMLHttpRequestResponse(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// ResponseText returns the value of property "XMLHttpRequest.responseText".
//
// The returned bool will be false if there is no such property.
func (this XMLHttpRequest) ResponseText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetXMLHttpRequestResponseText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ResponseXML returns the value of property "XMLHttpRequest.responseXML".
//
// The returned bool will be false if there is no such property.
func (this XMLHttpRequest) ResponseXML() (Document, bool) {
	var _ok bool
	_ret := bindings.GetXMLHttpRequestResponseXML(
		this.Ref(), js.Pointer(&_ok),
	)
	return Document{}.FromRef(_ret), _ok
}

// Open calls the method "XMLHttpRequest.open".
//
// The returned bool will be false if there is no such method.
func (this XMLHttpRequest) Open(method js.String, url js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXMLHttpRequestOpen(
		this.Ref(), js.Pointer(&_ok),
		method.Ref(),
		url.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// OpenFunc returns the method "XMLHttpRequest.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XMLHttpRequest) OpenFunc() (fn js.Func[func(method js.String, url js.String)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestOpenFunc(
			this.Ref(),
		),
	)
}

// Open1 calls the method "XMLHttpRequest.open".
//
// The returned bool will be false if there is no such method.
func (this XMLHttpRequest) Open1(method js.String, url js.String, async bool, username js.String, password js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXMLHttpRequestOpen1(
		this.Ref(), js.Pointer(&_ok),
		method.Ref(),
		url.Ref(),
		js.Bool(bool(async)),
		username.Ref(),
		password.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Open1Func returns the method "XMLHttpRequest.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XMLHttpRequest) Open1Func() (fn js.Func[func(method js.String, url js.String, async bool, username js.String, password js.String)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestOpen1Func(
			this.Ref(),
		),
	)
}

// Open2 calls the method "XMLHttpRequest.open".
//
// The returned bool will be false if there is no such method.
func (this XMLHttpRequest) Open2(method js.String, url js.String, async bool, username js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXMLHttpRequestOpen2(
		this.Ref(), js.Pointer(&_ok),
		method.Ref(),
		url.Ref(),
		js.Bool(bool(async)),
		username.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Open2Func returns the method "XMLHttpRequest.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XMLHttpRequest) Open2Func() (fn js.Func[func(method js.String, url js.String, async bool, username js.String)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestOpen2Func(
			this.Ref(),
		),
	)
}

// Open3 calls the method "XMLHttpRequest.open".
//
// The returned bool will be false if there is no such method.
func (this XMLHttpRequest) Open3(method js.String, url js.String, async bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXMLHttpRequestOpen3(
		this.Ref(), js.Pointer(&_ok),
		method.Ref(),
		url.Ref(),
		js.Bool(bool(async)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Open3Func returns the method "XMLHttpRequest.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XMLHttpRequest) Open3Func() (fn js.Func[func(method js.String, url js.String, async bool)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestOpen3Func(
			this.Ref(),
		),
	)
}

// SetRequestHeader calls the method "XMLHttpRequest.setRequestHeader".
//
// The returned bool will be false if there is no such method.
func (this XMLHttpRequest) SetRequestHeader(name js.String, value js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXMLHttpRequestSetRequestHeader(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		value.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetRequestHeaderFunc returns the method "XMLHttpRequest.setRequestHeader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XMLHttpRequest) SetRequestHeaderFunc() (fn js.Func[func(name js.String, value js.String)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestSetRequestHeaderFunc(
			this.Ref(),
		),
	)
}

// Send calls the method "XMLHttpRequest.send".
//
// The returned bool will be false if there is no such method.
func (this XMLHttpRequest) Send(body OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXMLHttpRequestSend(
		this.Ref(), js.Pointer(&_ok),
		body.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SendFunc returns the method "XMLHttpRequest.send".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XMLHttpRequest) SendFunc() (fn js.Func[func(body OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestSendFunc(
			this.Ref(),
		),
	)
}

// Send1 calls the method "XMLHttpRequest.send".
//
// The returned bool will be false if there is no such method.
func (this XMLHttpRequest) Send1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXMLHttpRequestSend1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Send1Func returns the method "XMLHttpRequest.send".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XMLHttpRequest) Send1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.XMLHttpRequestSend1Func(
			this.Ref(),
		),
	)
}

// Abort calls the method "XMLHttpRequest.abort".
//
// The returned bool will be false if there is no such method.
func (this XMLHttpRequest) Abort() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXMLHttpRequestAbort(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AbortFunc returns the method "XMLHttpRequest.abort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XMLHttpRequest) AbortFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.XMLHttpRequestAbortFunc(
			this.Ref(),
		),
	)
}

// GetResponseHeader calls the method "XMLHttpRequest.getResponseHeader".
//
// The returned bool will be false if there is no such method.
func (this XMLHttpRequest) GetResponseHeader(name js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallXMLHttpRequestGetResponseHeader(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetResponseHeaderFunc returns the method "XMLHttpRequest.getResponseHeader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XMLHttpRequest) GetResponseHeaderFunc() (fn js.Func[func(name js.String) js.String]) {
	return fn.FromRef(
		bindings.XMLHttpRequestGetResponseHeaderFunc(
			this.Ref(),
		),
	)
}

// GetAllResponseHeaders calls the method "XMLHttpRequest.getAllResponseHeaders".
//
// The returned bool will be false if there is no such method.
func (this XMLHttpRequest) GetAllResponseHeaders() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallXMLHttpRequestGetAllResponseHeaders(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetAllResponseHeadersFunc returns the method "XMLHttpRequest.getAllResponseHeaders".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XMLHttpRequest) GetAllResponseHeadersFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.XMLHttpRequestGetAllResponseHeadersFunc(
			this.Ref(),
		),
	)
}

// OverrideMimeType calls the method "XMLHttpRequest.overrideMimeType".
//
// The returned bool will be false if there is no such method.
func (this XMLHttpRequest) OverrideMimeType(mime js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXMLHttpRequestOverrideMimeType(
		this.Ref(), js.Pointer(&_ok),
		mime.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// OverrideMimeTypeFunc returns the method "XMLHttpRequest.overrideMimeType".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XMLHttpRequest) OverrideMimeTypeFunc() (fn js.Func[func(mime js.String)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestOverrideMimeTypeFunc(
			this.Ref(),
		),
	)
}

// SetAttributionReporting calls the method "XMLHttpRequest.setAttributionReporting".
//
// The returned bool will be false if there is no such method.
func (this XMLHttpRequest) SetAttributionReporting(options AttributionReportingRequestOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXMLHttpRequestSetAttributionReporting(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetAttributionReportingFunc returns the method "XMLHttpRequest.setAttributionReporting".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XMLHttpRequest) SetAttributionReportingFunc() (fn js.Func[func(options AttributionReportingRequestOptions)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestSetAttributionReportingFunc(
			this.Ref(),
		),
	)
}

// SetPrivateToken calls the method "XMLHttpRequest.setPrivateToken".
//
// The returned bool will be false if there is no such method.
func (this XMLHttpRequest) SetPrivateToken(privateToken PrivateToken) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXMLHttpRequestSetPrivateToken(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&privateToken),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPrivateTokenFunc returns the method "XMLHttpRequest.setPrivateToken".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XMLHttpRequest) SetPrivateTokenFunc() (fn js.Func[func(privateToken PrivateToken)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestSetPrivateTokenFunc(
			this.Ref(),
		),
	)
}

type OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String struct {
	ref js.Ref
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) Free() {
	x.ref.Free()
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) FromRef(ref js.Ref) OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String {
	return OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String{
		ref: ref,
	}
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) Blob() Blob {
	return Blob{}.FromRef(x.ref)
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) ArrayBuffer() js.ArrayBuffer {
	return js.ArrayBuffer{}.FromRef(x.ref)
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) FormData() FormData {
	return FormData{}.FromRef(x.ref)
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) URLSearchParams() URLSearchParams {
	return URLSearchParams{}.FromRef(x.ref)
}

func (x OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type XMLHttpRequestBodyInit = OneOf_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String

type XMLHttpRequestEventTarget struct {
	EventTarget
}

func (this XMLHttpRequestEventTarget) Once() XMLHttpRequestEventTarget {
	this.Ref().Once()
	return this
}

func (this XMLHttpRequestEventTarget) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this XMLHttpRequestEventTarget) FromRef(ref js.Ref) XMLHttpRequestEventTarget {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this XMLHttpRequestEventTarget) Free() {
	this.Ref().Free()
}

type XMLSerializer struct {
	ref js.Ref
}

func (this XMLSerializer) Once() XMLSerializer {
	this.Ref().Once()
	return this
}

func (this XMLSerializer) Ref() js.Ref {
	return this.ref
}

func (this XMLSerializer) FromRef(ref js.Ref) XMLSerializer {
	this.ref = ref
	return this
}

func (this XMLSerializer) Free() {
	this.Ref().Free()
}

// SerializeToString calls the method "XMLSerializer.serializeToString".
//
// The returned bool will be false if there is no such method.
func (this XMLSerializer) SerializeToString(root Node) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallXMLSerializerSerializeToString(
		this.Ref(), js.Pointer(&_ok),
		root.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// SerializeToStringFunc returns the method "XMLSerializer.serializeToString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XMLSerializer) SerializeToStringFunc() (fn js.Func[func(root Node) js.String]) {
	return fn.FromRef(
		bindings.XMLSerializerSerializeToStringFunc(
			this.Ref(),
		),
	)
}

type XPathEvaluator struct {
	ref js.Ref
}

func (this XPathEvaluator) Once() XPathEvaluator {
	this.Ref().Once()
	return this
}

func (this XPathEvaluator) Ref() js.Ref {
	return this.ref
}

func (this XPathEvaluator) FromRef(ref js.Ref) XPathEvaluator {
	this.ref = ref
	return this
}

func (this XPathEvaluator) Free() {
	this.Ref().Free()
}

// CreateExpression calls the method "XPathEvaluator.createExpression".
//
// The returned bool will be false if there is no such method.
func (this XPathEvaluator) CreateExpression(expression js.String, resolver js.Func[func(prefix js.String) js.String]) (XPathExpression, bool) {
	var _ok bool
	_ret := bindings.CallXPathEvaluatorCreateExpression(
		this.Ref(), js.Pointer(&_ok),
		expression.Ref(),
		resolver.Ref(),
	)

	return XPathExpression{}.FromRef(_ret), _ok
}

// CreateExpressionFunc returns the method "XPathEvaluator.createExpression".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XPathEvaluator) CreateExpressionFunc() (fn js.Func[func(expression js.String, resolver js.Func[func(prefix js.String) js.String]) XPathExpression]) {
	return fn.FromRef(
		bindings.XPathEvaluatorCreateExpressionFunc(
			this.Ref(),
		),
	)
}

// CreateExpression1 calls the method "XPathEvaluator.createExpression".
//
// The returned bool will be false if there is no such method.
func (this XPathEvaluator) CreateExpression1(expression js.String) (XPathExpression, bool) {
	var _ok bool
	_ret := bindings.CallXPathEvaluatorCreateExpression1(
		this.Ref(), js.Pointer(&_ok),
		expression.Ref(),
	)

	return XPathExpression{}.FromRef(_ret), _ok
}

// CreateExpression1Func returns the method "XPathEvaluator.createExpression".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XPathEvaluator) CreateExpression1Func() (fn js.Func[func(expression js.String) XPathExpression]) {
	return fn.FromRef(
		bindings.XPathEvaluatorCreateExpression1Func(
			this.Ref(),
		),
	)
}

// CreateNSResolver calls the method "XPathEvaluator.createNSResolver".
//
// The returned bool will be false if there is no such method.
func (this XPathEvaluator) CreateNSResolver(nodeResolver Node) (Node, bool) {
	var _ok bool
	_ret := bindings.CallXPathEvaluatorCreateNSResolver(
		this.Ref(), js.Pointer(&_ok),
		nodeResolver.Ref(),
	)

	return Node{}.FromRef(_ret), _ok
}

// CreateNSResolverFunc returns the method "XPathEvaluator.createNSResolver".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XPathEvaluator) CreateNSResolverFunc() (fn js.Func[func(nodeResolver Node) Node]) {
	return fn.FromRef(
		bindings.XPathEvaluatorCreateNSResolverFunc(
			this.Ref(),
		),
	)
}

// Evaluate calls the method "XPathEvaluator.evaluate".
//
// The returned bool will be false if there is no such method.
func (this XPathEvaluator) Evaluate(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16, result XPathResult) (XPathResult, bool) {
	var _ok bool
	_ret := bindings.CallXPathEvaluatorEvaluate(
		this.Ref(), js.Pointer(&_ok),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
		result.Ref(),
	)

	return XPathResult{}.FromRef(_ret), _ok
}

// EvaluateFunc returns the method "XPathEvaluator.evaluate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XPathEvaluator) EvaluateFunc() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16, result XPathResult) XPathResult]) {
	return fn.FromRef(
		bindings.XPathEvaluatorEvaluateFunc(
			this.Ref(),
		),
	)
}

// Evaluate1 calls the method "XPathEvaluator.evaluate".
//
// The returned bool will be false if there is no such method.
func (this XPathEvaluator) Evaluate1(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16) (XPathResult, bool) {
	var _ok bool
	_ret := bindings.CallXPathEvaluatorEvaluate1(
		this.Ref(), js.Pointer(&_ok),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
	)

	return XPathResult{}.FromRef(_ret), _ok
}

// Evaluate1Func returns the method "XPathEvaluator.evaluate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XPathEvaluator) Evaluate1Func() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16) XPathResult]) {
	return fn.FromRef(
		bindings.XPathEvaluatorEvaluate1Func(
			this.Ref(),
		),
	)
}

// Evaluate2 calls the method "XPathEvaluator.evaluate".
//
// The returned bool will be false if there is no such method.
func (this XPathEvaluator) Evaluate2(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String]) (XPathResult, bool) {
	var _ok bool
	_ret := bindings.CallXPathEvaluatorEvaluate2(
		this.Ref(), js.Pointer(&_ok),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
	)

	return XPathResult{}.FromRef(_ret), _ok
}

// Evaluate2Func returns the method "XPathEvaluator.evaluate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XPathEvaluator) Evaluate2Func() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String]) XPathResult]) {
	return fn.FromRef(
		bindings.XPathEvaluatorEvaluate2Func(
			this.Ref(),
		),
	)
}

// Evaluate3 calls the method "XPathEvaluator.evaluate".
//
// The returned bool will be false if there is no such method.
func (this XPathEvaluator) Evaluate3(expression js.String, contextNode Node) (XPathResult, bool) {
	var _ok bool
	_ret := bindings.CallXPathEvaluatorEvaluate3(
		this.Ref(), js.Pointer(&_ok),
		expression.Ref(),
		contextNode.Ref(),
	)

	return XPathResult{}.FromRef(_ret), _ok
}

// Evaluate3Func returns the method "XPathEvaluator.evaluate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XPathEvaluator) Evaluate3Func() (fn js.Func[func(expression js.String, contextNode Node) XPathResult]) {
	return fn.FromRef(
		bindings.XPathEvaluatorEvaluate3Func(
			this.Ref(),
		),
	)
}

type XRBoundedReferenceSpace struct {
	XRReferenceSpace
}

func (this XRBoundedReferenceSpace) Once() XRBoundedReferenceSpace {
	this.Ref().Once()
	return this
}

func (this XRBoundedReferenceSpace) Ref() js.Ref {
	return this.XRReferenceSpace.Ref()
}

func (this XRBoundedReferenceSpace) FromRef(ref js.Ref) XRBoundedReferenceSpace {
	this.XRReferenceSpace = this.XRReferenceSpace.FromRef(ref)
	return this
}

func (this XRBoundedReferenceSpace) Free() {
	this.Ref().Free()
}

// BoundsGeometry returns the value of property "XRBoundedReferenceSpace.boundsGeometry".
//
// The returned bool will be false if there is no such property.
func (this XRBoundedReferenceSpace) BoundsGeometry() (js.FrozenArray[DOMPointReadOnly], bool) {
	var _ok bool
	_ret := bindings.GetXRBoundedReferenceSpaceBoundsGeometry(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[DOMPointReadOnly]{}.FromRef(_ret), _ok
}
