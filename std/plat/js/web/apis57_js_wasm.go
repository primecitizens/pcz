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

func NewWebGLContextEvent(typ js.String, eventInit WebGLContextEventInit) (ret WebGLContextEvent) {
	ret.ref = bindings.NewWebGLContextEventByWebGLContextEvent(
		typ.Ref(),
		js.Pointer(&eventInit))
	return
}

func NewWebGLContextEventByWebGLContextEvent1(typ js.String) (ret WebGLContextEvent) {
	ret.ref = bindings.NewWebGLContextEventByWebGLContextEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this WebGLContextEvent) StatusMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWebGLContextEventStatusMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewWebSocket(url js.String, protocols OneOf_String_ArrayString) (ret WebSocket) {
	ret.ref = bindings.NewWebSocketByWebSocket(
		url.Ref(),
		protocols.Ref())
	return
}

func NewWebSocketByWebSocket1(url js.String) (ret WebSocket) {
	ret.ref = bindings.NewWebSocketByWebSocket1(
		url.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this WebSocket) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWebSocketUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ReadyState returns the value of property "WebSocket.readyState".
//
// It returns ok=false if there is no such property.
func (this WebSocket) ReadyState() (ret uint16, ok bool) {
	ok = js.True == bindings.GetWebSocketReadyState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BufferedAmount returns the value of property "WebSocket.bufferedAmount".
//
// It returns ok=false if there is no such property.
func (this WebSocket) BufferedAmount() (ret uint64, ok bool) {
	ok = js.True == bindings.GetWebSocketBufferedAmount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Extensions returns the value of property "WebSocket.extensions".
//
// It returns ok=false if there is no such property.
func (this WebSocket) Extensions() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWebSocketExtensions(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Protocol returns the value of property "WebSocket.protocol".
//
// It returns ok=false if there is no such property.
func (this WebSocket) Protocol() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWebSocketProtocol(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BinaryType returns the value of property "WebSocket.binaryType".
//
// It returns ok=false if there is no such property.
func (this WebSocket) BinaryType() (ret BinaryType, ok bool) {
	ok = js.True == bindings.GetWebSocketBinaryType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBinaryType sets the value of property "WebSocket.binaryType" to val.
//
// It returns false if the property cannot be set.
func (this WebSocket) SetBinaryType(val BinaryType) bool {
	return js.True == bindings.SetWebSocketBinaryType(
		this.Ref(),
		uint32(val),
	)
}

// HasClose returns true if the method "WebSocket.close" exists.
func (this WebSocket) HasClose() bool {
	return js.True == bindings.HasWebSocketClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "WebSocket.close".
func (this WebSocket) CloseFunc() (fn js.Func[func(code uint16, reason js.String)]) {
	return fn.FromRef(
		bindings.WebSocketCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "WebSocket.close".
func (this WebSocket) Close(code uint16, reason js.String) (ret js.Void) {
	bindings.CallWebSocketClose(
		this.Ref(), js.Pointer(&ret),
		uint32(code),
		reason.Ref(),
	)

	return
}

// TryClose calls the method "WebSocket.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebSocket) TryClose(code uint16, reason js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebSocketClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(code),
		reason.Ref(),
	)

	return
}

// HasClose1 returns true if the method "WebSocket.close" exists.
func (this WebSocket) HasClose1() bool {
	return js.True == bindings.HasWebSocketClose1(
		this.Ref(),
	)
}

// Close1Func returns the method "WebSocket.close".
func (this WebSocket) Close1Func() (fn js.Func[func(code uint16)]) {
	return fn.FromRef(
		bindings.WebSocketClose1Func(
			this.Ref(),
		),
	)
}

// Close1 calls the method "WebSocket.close".
func (this WebSocket) Close1(code uint16) (ret js.Void) {
	bindings.CallWebSocketClose1(
		this.Ref(), js.Pointer(&ret),
		uint32(code),
	)

	return
}

// TryClose1 calls the method "WebSocket.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebSocket) TryClose1(code uint16) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebSocketClose1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(code),
	)

	return
}

// HasClose2 returns true if the method "WebSocket.close" exists.
func (this WebSocket) HasClose2() bool {
	return js.True == bindings.HasWebSocketClose2(
		this.Ref(),
	)
}

// Close2Func returns the method "WebSocket.close".
func (this WebSocket) Close2Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebSocketClose2Func(
			this.Ref(),
		),
	)
}

// Close2 calls the method "WebSocket.close".
func (this WebSocket) Close2() (ret js.Void) {
	bindings.CallWebSocketClose2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose2 calls the method "WebSocket.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebSocket) TryClose2() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebSocketClose2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSend returns true if the method "WebSocket.send" exists.
func (this WebSocket) HasSend() bool {
	return js.True == bindings.HasWebSocketSend(
		this.Ref(),
	)
}

// SendFunc returns the method "WebSocket.send".
func (this WebSocket) SendFunc() (fn js.Func[func(data OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String)]) {
	return fn.FromRef(
		bindings.WebSocketSendFunc(
			this.Ref(),
		),
	)
}

// Send calls the method "WebSocket.send".
func (this WebSocket) Send(data OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) (ret js.Void) {
	bindings.CallWebSocketSend(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySend calls the method "WebSocket.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebSocket) TrySend(data OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebSocketSend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
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
	//
	// NOTE: Datagrams.FFI_USE MUST be set to true to get Datagrams used.
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

func NewWebTransportReceiveStream(underlyingSource js.Object, strategy QueuingStrategy) (ret WebTransportReceiveStream) {
	ret.ref = bindings.NewWebTransportReceiveStreamByWebTransportReceiveStream(
		underlyingSource.Ref(),
		js.Pointer(&strategy))
	return
}

func NewWebTransportReceiveStreamByWebTransportReceiveStream1(underlyingSource js.Object) (ret WebTransportReceiveStream) {
	ret.ref = bindings.NewWebTransportReceiveStreamByWebTransportReceiveStream1(
		underlyingSource.Ref())
	return
}

func NewWebTransportReceiveStreamByWebTransportReceiveStream2() (ret WebTransportReceiveStream) {
	ret.ref = bindings.NewWebTransportReceiveStreamByWebTransportReceiveStream2()
	return
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

// HasGetStats returns true if the method "WebTransportReceiveStream.getStats" exists.
func (this WebTransportReceiveStream) HasGetStats() bool {
	return js.True == bindings.HasWebTransportReceiveStreamGetStats(
		this.Ref(),
	)
}

// GetStatsFunc returns the method "WebTransportReceiveStream.getStats".
func (this WebTransportReceiveStream) GetStatsFunc() (fn js.Func[func() js.Promise[WebTransportReceiveStreamStats]]) {
	return fn.FromRef(
		bindings.WebTransportReceiveStreamGetStatsFunc(
			this.Ref(),
		),
	)
}

// GetStats calls the method "WebTransportReceiveStream.getStats".
func (this WebTransportReceiveStream) GetStats() (ret js.Promise[WebTransportReceiveStreamStats]) {
	bindings.CallWebTransportReceiveStreamGetStats(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetStats calls the method "WebTransportReceiveStream.getStats"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransportReceiveStream) TryGetStats() (ret js.Promise[WebTransportReceiveStreamStats], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportReceiveStreamGetStats(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewWebTransportSendStream(underlyingSink js.Object, strategy QueuingStrategy) (ret WebTransportSendStream) {
	ret.ref = bindings.NewWebTransportSendStreamByWebTransportSendStream(
		underlyingSink.Ref(),
		js.Pointer(&strategy))
	return
}

func NewWebTransportSendStreamByWebTransportSendStream1(underlyingSink js.Object) (ret WebTransportSendStream) {
	ret.ref = bindings.NewWebTransportSendStreamByWebTransportSendStream1(
		underlyingSink.Ref())
	return
}

func NewWebTransportSendStreamByWebTransportSendStream2() (ret WebTransportSendStream) {
	ret.ref = bindings.NewWebTransportSendStreamByWebTransportSendStream2()
	return
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
// It returns ok=false if there is no such property.
func (this WebTransportSendStream) SendOrder() (ret int64, ok bool) {
	ok = js.True == bindings.GetWebTransportSendStreamSendOrder(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSendOrder sets the value of property "WebTransportSendStream.sendOrder" to val.
//
// It returns false if the property cannot be set.
func (this WebTransportSendStream) SetSendOrder(val int64) bool {
	return js.True == bindings.SetWebTransportSendStreamSendOrder(
		this.Ref(),
		float64(val),
	)
}

// HasGetStats returns true if the method "WebTransportSendStream.getStats" exists.
func (this WebTransportSendStream) HasGetStats() bool {
	return js.True == bindings.HasWebTransportSendStreamGetStats(
		this.Ref(),
	)
}

// GetStatsFunc returns the method "WebTransportSendStream.getStats".
func (this WebTransportSendStream) GetStatsFunc() (fn js.Func[func() js.Promise[WebTransportSendStreamStats]]) {
	return fn.FromRef(
		bindings.WebTransportSendStreamGetStatsFunc(
			this.Ref(),
		),
	)
}

// GetStats calls the method "WebTransportSendStream.getStats".
func (this WebTransportSendStream) GetStats() (ret js.Promise[WebTransportSendStreamStats]) {
	bindings.CallWebTransportSendStreamGetStats(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetStats calls the method "WebTransportSendStream.getStats"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransportSendStream) TryGetStats() (ret js.Promise[WebTransportSendStreamStats], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportSendStreamGetStats(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this WebTransportBidirectionalStream) Readable() (ret WebTransportReceiveStream, ok bool) {
	ok = js.True == bindings.GetWebTransportBidirectionalStreamReadable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Writable returns the value of property "WebTransportBidirectionalStream.writable".
//
// It returns ok=false if there is no such property.
func (this WebTransportBidirectionalStream) Writable() (ret WebTransportSendStream, ok bool) {
	ok = js.True == bindings.GetWebTransportBidirectionalStreamWritable(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this WebTransportDatagramDuplexStream) Readable() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagramDuplexStreamReadable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Writable returns the value of property "WebTransportDatagramDuplexStream.writable".
//
// It returns ok=false if there is no such property.
func (this WebTransportDatagramDuplexStream) Writable() (ret WritableStream, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagramDuplexStreamWritable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxDatagramSize returns the value of property "WebTransportDatagramDuplexStream.maxDatagramSize".
//
// It returns ok=false if there is no such property.
func (this WebTransportDatagramDuplexStream) MaxDatagramSize() (ret uint32, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagramDuplexStreamMaxDatagramSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IncomingMaxAge returns the value of property "WebTransportDatagramDuplexStream.incomingMaxAge".
//
// It returns ok=false if there is no such property.
func (this WebTransportDatagramDuplexStream) IncomingMaxAge() (ret float64, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagramDuplexStreamIncomingMaxAge(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetIncomingMaxAge sets the value of property "WebTransportDatagramDuplexStream.incomingMaxAge" to val.
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
// It returns ok=false if there is no such property.
func (this WebTransportDatagramDuplexStream) OutgoingMaxAge() (ret float64, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagramDuplexStreamOutgoingMaxAge(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetOutgoingMaxAge sets the value of property "WebTransportDatagramDuplexStream.outgoingMaxAge" to val.
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
// It returns ok=false if there is no such property.
func (this WebTransportDatagramDuplexStream) IncomingHighWaterMark() (ret float64, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagramDuplexStreamIncomingHighWaterMark(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetIncomingHighWaterMark sets the value of property "WebTransportDatagramDuplexStream.incomingHighWaterMark" to val.
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
// It returns ok=false if there is no such property.
func (this WebTransportDatagramDuplexStream) OutgoingHighWaterMark() (ret float64, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagramDuplexStreamOutgoingHighWaterMark(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetOutgoingHighWaterMark sets the value of property "WebTransportDatagramDuplexStream.outgoingHighWaterMark" to val.
//
// It returns false if the property cannot be set.
func (this WebTransportDatagramDuplexStream) SetOutgoingHighWaterMark(val float64) bool {
	return js.True == bindings.SetWebTransportDatagramDuplexStreamOutgoingHighWaterMark(
		this.Ref(),
		float64(val),
	)
}

func NewWebTransport(url js.String, options WebTransportOptions) (ret WebTransport) {
	ret.ref = bindings.NewWebTransportByWebTransport(
		url.Ref(),
		js.Pointer(&options))
	return
}

func NewWebTransportByWebTransport1(url js.String) (ret WebTransport) {
	ret.ref = bindings.NewWebTransportByWebTransport1(
		url.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this WebTransport) Ready() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetWebTransportReady(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Reliability returns the value of property "WebTransport.reliability".
//
// It returns ok=false if there is no such property.
func (this WebTransport) Reliability() (ret WebTransportReliabilityMode, ok bool) {
	ok = js.True == bindings.GetWebTransportReliability(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CongestionControl returns the value of property "WebTransport.congestionControl".
//
// It returns ok=false if there is no such property.
func (this WebTransport) CongestionControl() (ret WebTransportCongestionControl, ok bool) {
	ok = js.True == bindings.GetWebTransportCongestionControl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Closed returns the value of property "WebTransport.closed".
//
// It returns ok=false if there is no such property.
func (this WebTransport) Closed() (ret js.Promise[WebTransportCloseInfo], ok bool) {
	ok = js.True == bindings.GetWebTransportClosed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Draining returns the value of property "WebTransport.draining".
//
// It returns ok=false if there is no such property.
func (this WebTransport) Draining() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetWebTransportDraining(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Datagrams returns the value of property "WebTransport.datagrams".
//
// It returns ok=false if there is no such property.
func (this WebTransport) Datagrams() (ret WebTransportDatagramDuplexStream, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagrams(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IncomingBidirectionalStreams returns the value of property "WebTransport.incomingBidirectionalStreams".
//
// It returns ok=false if there is no such property.
func (this WebTransport) IncomingBidirectionalStreams() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetWebTransportIncomingBidirectionalStreams(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IncomingUnidirectionalStreams returns the value of property "WebTransport.incomingUnidirectionalStreams".
//
// It returns ok=false if there is no such property.
func (this WebTransport) IncomingUnidirectionalStreams() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetWebTransportIncomingUnidirectionalStreams(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetStats returns true if the method "WebTransport.getStats" exists.
func (this WebTransport) HasGetStats() bool {
	return js.True == bindings.HasWebTransportGetStats(
		this.Ref(),
	)
}

// GetStatsFunc returns the method "WebTransport.getStats".
func (this WebTransport) GetStatsFunc() (fn js.Func[func() js.Promise[WebTransportStats]]) {
	return fn.FromRef(
		bindings.WebTransportGetStatsFunc(
			this.Ref(),
		),
	)
}

// GetStats calls the method "WebTransport.getStats".
func (this WebTransport) GetStats() (ret js.Promise[WebTransportStats]) {
	bindings.CallWebTransportGetStats(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetStats calls the method "WebTransport.getStats"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransport) TryGetStats() (ret js.Promise[WebTransportStats], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportGetStats(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "WebTransport.close" exists.
func (this WebTransport) HasClose() bool {
	return js.True == bindings.HasWebTransportClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "WebTransport.close".
func (this WebTransport) CloseFunc() (fn js.Func[func(closeInfo WebTransportCloseInfo)]) {
	return fn.FromRef(
		bindings.WebTransportCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "WebTransport.close".
func (this WebTransport) Close(closeInfo WebTransportCloseInfo) (ret js.Void) {
	bindings.CallWebTransportClose(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&closeInfo),
	)

	return
}

// TryClose calls the method "WebTransport.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransport) TryClose(closeInfo WebTransportCloseInfo) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&closeInfo),
	)

	return
}

// HasClose1 returns true if the method "WebTransport.close" exists.
func (this WebTransport) HasClose1() bool {
	return js.True == bindings.HasWebTransportClose1(
		this.Ref(),
	)
}

// Close1Func returns the method "WebTransport.close".
func (this WebTransport) Close1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebTransportClose1Func(
			this.Ref(),
		),
	)
}

// Close1 calls the method "WebTransport.close".
func (this WebTransport) Close1() (ret js.Void) {
	bindings.CallWebTransportClose1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose1 calls the method "WebTransport.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransport) TryClose1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportClose1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateBidirectionalStream returns true if the method "WebTransport.createBidirectionalStream" exists.
func (this WebTransport) HasCreateBidirectionalStream() bool {
	return js.True == bindings.HasWebTransportCreateBidirectionalStream(
		this.Ref(),
	)
}

// CreateBidirectionalStreamFunc returns the method "WebTransport.createBidirectionalStream".
func (this WebTransport) CreateBidirectionalStreamFunc() (fn js.Func[func(options WebTransportSendStreamOptions) js.Promise[WebTransportBidirectionalStream]]) {
	return fn.FromRef(
		bindings.WebTransportCreateBidirectionalStreamFunc(
			this.Ref(),
		),
	)
}

// CreateBidirectionalStream calls the method "WebTransport.createBidirectionalStream".
func (this WebTransport) CreateBidirectionalStream(options WebTransportSendStreamOptions) (ret js.Promise[WebTransportBidirectionalStream]) {
	bindings.CallWebTransportCreateBidirectionalStream(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryCreateBidirectionalStream calls the method "WebTransport.createBidirectionalStream"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransport) TryCreateBidirectionalStream(options WebTransportSendStreamOptions) (ret js.Promise[WebTransportBidirectionalStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportCreateBidirectionalStream(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasCreateBidirectionalStream1 returns true if the method "WebTransport.createBidirectionalStream" exists.
func (this WebTransport) HasCreateBidirectionalStream1() bool {
	return js.True == bindings.HasWebTransportCreateBidirectionalStream1(
		this.Ref(),
	)
}

// CreateBidirectionalStream1Func returns the method "WebTransport.createBidirectionalStream".
func (this WebTransport) CreateBidirectionalStream1Func() (fn js.Func[func() js.Promise[WebTransportBidirectionalStream]]) {
	return fn.FromRef(
		bindings.WebTransportCreateBidirectionalStream1Func(
			this.Ref(),
		),
	)
}

// CreateBidirectionalStream1 calls the method "WebTransport.createBidirectionalStream".
func (this WebTransport) CreateBidirectionalStream1() (ret js.Promise[WebTransportBidirectionalStream]) {
	bindings.CallWebTransportCreateBidirectionalStream1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateBidirectionalStream1 calls the method "WebTransport.createBidirectionalStream"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransport) TryCreateBidirectionalStream1() (ret js.Promise[WebTransportBidirectionalStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportCreateBidirectionalStream1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateUnidirectionalStream returns true if the method "WebTransport.createUnidirectionalStream" exists.
func (this WebTransport) HasCreateUnidirectionalStream() bool {
	return js.True == bindings.HasWebTransportCreateUnidirectionalStream(
		this.Ref(),
	)
}

// CreateUnidirectionalStreamFunc returns the method "WebTransport.createUnidirectionalStream".
func (this WebTransport) CreateUnidirectionalStreamFunc() (fn js.Func[func(options WebTransportSendStreamOptions) js.Promise[WebTransportSendStream]]) {
	return fn.FromRef(
		bindings.WebTransportCreateUnidirectionalStreamFunc(
			this.Ref(),
		),
	)
}

// CreateUnidirectionalStream calls the method "WebTransport.createUnidirectionalStream".
func (this WebTransport) CreateUnidirectionalStream(options WebTransportSendStreamOptions) (ret js.Promise[WebTransportSendStream]) {
	bindings.CallWebTransportCreateUnidirectionalStream(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryCreateUnidirectionalStream calls the method "WebTransport.createUnidirectionalStream"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransport) TryCreateUnidirectionalStream(options WebTransportSendStreamOptions) (ret js.Promise[WebTransportSendStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportCreateUnidirectionalStream(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasCreateUnidirectionalStream1 returns true if the method "WebTransport.createUnidirectionalStream" exists.
func (this WebTransport) HasCreateUnidirectionalStream1() bool {
	return js.True == bindings.HasWebTransportCreateUnidirectionalStream1(
		this.Ref(),
	)
}

// CreateUnidirectionalStream1Func returns the method "WebTransport.createUnidirectionalStream".
func (this WebTransport) CreateUnidirectionalStream1Func() (fn js.Func[func() js.Promise[WebTransportSendStream]]) {
	return fn.FromRef(
		bindings.WebTransportCreateUnidirectionalStream1Func(
			this.Ref(),
		),
	)
}

// CreateUnidirectionalStream1 calls the method "WebTransport.createUnidirectionalStream".
func (this WebTransport) CreateUnidirectionalStream1() (ret js.Promise[WebTransportSendStream]) {
	bindings.CallWebTransportCreateUnidirectionalStream1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateUnidirectionalStream1 calls the method "WebTransport.createUnidirectionalStream"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransport) TryCreateUnidirectionalStream1() (ret js.Promise[WebTransportSendStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportCreateUnidirectionalStream1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewWebTransportError(message js.String, options WebTransportErrorOptions) (ret WebTransportError) {
	ret.ref = bindings.NewWebTransportErrorByWebTransportError(
		message.Ref(),
		js.Pointer(&options))
	return
}

func NewWebTransportErrorByWebTransportError1(message js.String) (ret WebTransportError) {
	ret.ref = bindings.NewWebTransportErrorByWebTransportError1(
		message.Ref())
	return
}

func NewWebTransportErrorByWebTransportError2() (ret WebTransportError) {
	ret.ref = bindings.NewWebTransportErrorByWebTransportError2()
	return
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
// It returns ok=false if there is no such property.
func (this WebTransportError) Source() (ret WebTransportErrorSource, ok bool) {
	ok = js.True == bindings.GetWebTransportErrorSource(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StreamErrorCode returns the value of property "WebTransportError.streamErrorCode".
//
// It returns ok=false if there is no such property.
func (this WebTransportError) StreamErrorCode() (ret uint32, ok bool) {
	ok = js.True == bindings.GetWebTransportErrorStreamErrorCode(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewWheelEvent(typ js.String, eventInitDict WheelEventInit) (ret WheelEvent) {
	ret.ref = bindings.NewWheelEventByWheelEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewWheelEventByWheelEvent1(typ js.String) (ret WheelEvent) {
	ret.ref = bindings.NewWheelEventByWheelEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this WheelEvent) DeltaX() (ret float64, ok bool) {
	ok = js.True == bindings.GetWheelEventDeltaX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DeltaY returns the value of property "WheelEvent.deltaY".
//
// It returns ok=false if there is no such property.
func (this WheelEvent) DeltaY() (ret float64, ok bool) {
	ok = js.True == bindings.GetWheelEventDeltaY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DeltaZ returns the value of property "WheelEvent.deltaZ".
//
// It returns ok=false if there is no such property.
func (this WheelEvent) DeltaZ() (ret float64, ok bool) {
	ok = js.True == bindings.GetWheelEventDeltaZ(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DeltaMode returns the value of property "WheelEvent.deltaMode".
//
// It returns ok=false if there is no such property.
func (this WheelEvent) DeltaMode() (ret uint32, ok bool) {
	ok = js.True == bindings.GetWheelEventDeltaMode(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewWindowControlsOverlayGeometryChangeEvent(typ js.String, eventInitDict WindowControlsOverlayGeometryChangeEventInit) (ret WindowControlsOverlayGeometryChangeEvent) {
	ret.ref = bindings.NewWindowControlsOverlayGeometryChangeEventByWindowControlsOverlayGeometryChangeEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
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
// It returns ok=false if there is no such property.
func (this WindowControlsOverlayGeometryChangeEvent) TitlebarAreaRect() (ret DOMRect, ok bool) {
	ok = js.True == bindings.GetWindowControlsOverlayGeometryChangeEventTitlebarAreaRect(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Visible returns the value of property "WindowControlsOverlayGeometryChangeEvent.visible".
//
// It returns ok=false if there is no such property.
func (this WindowControlsOverlayGeometryChangeEvent) Visible() (ret bool, ok bool) {
	ok = js.True == bindings.GetWindowControlsOverlayGeometryChangeEventVisible(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this WorkerLocation) Href() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationHref(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Origin returns the value of property "WorkerLocation.origin".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Protocol returns the value of property "WorkerLocation.protocol".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Protocol() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationProtocol(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Host returns the value of property "WorkerLocation.host".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Host() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationHost(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Hostname returns the value of property "WorkerLocation.hostname".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Hostname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationHostname(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Port returns the value of property "WorkerLocation.port".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Port() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationPort(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Pathname returns the value of property "WorkerLocation.pathname".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Pathname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationPathname(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Search returns the value of property "WorkerLocation.search".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Search() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationSearch(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Hash returns the value of property "WorkerLocation.hash".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Hash() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationHash(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Hid() (ret HID, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorHid(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Serial returns the value of property "WorkerNavigator.serial".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Serial() (ret Serial, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorSerial(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Permissions returns the value of property "WorkerNavigator.permissions".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Permissions() (ret Permissions, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorPermissions(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Usb returns the value of property "WorkerNavigator.usb".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Usb() (ret USB, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorUsb(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MediaCapabilities returns the value of property "WorkerNavigator.mediaCapabilities".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) MediaCapabilities() (ret MediaCapabilities, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorMediaCapabilities(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ServiceWorker returns the value of property "WorkerNavigator.serviceWorker".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) ServiceWorker() (ret ServiceWorkerContainer, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorServiceWorker(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UserAgentData returns the value of property "WorkerNavigator.userAgentData".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) UserAgentData() (ret NavigatorUAData, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorUserAgentData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DeviceMemory returns the value of property "WorkerNavigator.deviceMemory".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) DeviceMemory() (ret float64, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorDeviceMemory(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Storage returns the value of property "WorkerNavigator.storage".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Storage() (ret StorageManager, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorStorage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StorageBuckets returns the value of property "WorkerNavigator.storageBuckets".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) StorageBuckets() (ret StorageBucketManager, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorStorageBuckets(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Locks returns the value of property "WorkerNavigator.locks".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Locks() (ret LockManager, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorLocks(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AppCodeName returns the value of property "WorkerNavigator.appCodeName".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) AppCodeName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorAppCodeName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AppName returns the value of property "WorkerNavigator.appName".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) AppName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorAppName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AppVersion returns the value of property "WorkerNavigator.appVersion".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) AppVersion() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorAppVersion(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Platform returns the value of property "WorkerNavigator.platform".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Platform() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorPlatform(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Product returns the value of property "WorkerNavigator.product".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Product() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorProduct(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ProductSub returns the value of property "WorkerNavigator.productSub".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) ProductSub() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorProductSub(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UserAgent returns the value of property "WorkerNavigator.userAgent".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) UserAgent() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorUserAgent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Vendor returns the value of property "WorkerNavigator.vendor".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Vendor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorVendor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// VendorSub returns the value of property "WorkerNavigator.vendorSub".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) VendorSub() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorVendorSub(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Oscpu returns the value of property "WorkerNavigator.oscpu".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Oscpu() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorOscpu(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Language returns the value of property "WorkerNavigator.language".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Language() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorLanguage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Languages returns the value of property "WorkerNavigator.languages".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Languages() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorLanguages(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OnLine returns the value of property "WorkerNavigator.onLine".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) OnLine() (ret bool, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorOnLine(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HardwareConcurrency returns the value of property "WorkerNavigator.hardwareConcurrency".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) HardwareConcurrency() (ret uint64, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorHardwareConcurrency(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Connection returns the value of property "WorkerNavigator.connection".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Connection() (ret NetworkInformation, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorConnection(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Ml returns the value of property "WorkerNavigator.ml".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Ml() (ret ML, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorMl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Gpu returns the value of property "WorkerNavigator.gpu".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Gpu() (ret GPU, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorGpu(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasTaintEnabled returns true if the method "WorkerNavigator.taintEnabled" exists.
func (this WorkerNavigator) HasTaintEnabled() bool {
	return js.True == bindings.HasWorkerNavigatorTaintEnabled(
		this.Ref(),
	)
}

// TaintEnabledFunc returns the method "WorkerNavigator.taintEnabled".
func (this WorkerNavigator) TaintEnabledFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.WorkerNavigatorTaintEnabledFunc(
			this.Ref(),
		),
	)
}

// TaintEnabled calls the method "WorkerNavigator.taintEnabled".
func (this WorkerNavigator) TaintEnabled() (ret bool) {
	bindings.CallWorkerNavigatorTaintEnabled(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTaintEnabled calls the method "WorkerNavigator.taintEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerNavigator) TryTaintEnabled() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerNavigatorTaintEnabled(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetAppBadge returns true if the method "WorkerNavigator.setAppBadge" exists.
func (this WorkerNavigator) HasSetAppBadge() bool {
	return js.True == bindings.HasWorkerNavigatorSetAppBadge(
		this.Ref(),
	)
}

// SetAppBadgeFunc returns the method "WorkerNavigator.setAppBadge".
func (this WorkerNavigator) SetAppBadgeFunc() (fn js.Func[func(contents uint64) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WorkerNavigatorSetAppBadgeFunc(
			this.Ref(),
		),
	)
}

// SetAppBadge calls the method "WorkerNavigator.setAppBadge".
func (this WorkerNavigator) SetAppBadge(contents uint64) (ret js.Promise[js.Void]) {
	bindings.CallWorkerNavigatorSetAppBadge(
		this.Ref(), js.Pointer(&ret),
		float64(contents),
	)

	return
}

// TrySetAppBadge calls the method "WorkerNavigator.setAppBadge"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerNavigator) TrySetAppBadge(contents uint64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerNavigatorSetAppBadge(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(contents),
	)

	return
}

// HasSetAppBadge1 returns true if the method "WorkerNavigator.setAppBadge" exists.
func (this WorkerNavigator) HasSetAppBadge1() bool {
	return js.True == bindings.HasWorkerNavigatorSetAppBadge1(
		this.Ref(),
	)
}

// SetAppBadge1Func returns the method "WorkerNavigator.setAppBadge".
func (this WorkerNavigator) SetAppBadge1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WorkerNavigatorSetAppBadge1Func(
			this.Ref(),
		),
	)
}

// SetAppBadge1 calls the method "WorkerNavigator.setAppBadge".
func (this WorkerNavigator) SetAppBadge1() (ret js.Promise[js.Void]) {
	bindings.CallWorkerNavigatorSetAppBadge1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySetAppBadge1 calls the method "WorkerNavigator.setAppBadge"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerNavigator) TrySetAppBadge1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerNavigatorSetAppBadge1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClearAppBadge returns true if the method "WorkerNavigator.clearAppBadge" exists.
func (this WorkerNavigator) HasClearAppBadge() bool {
	return js.True == bindings.HasWorkerNavigatorClearAppBadge(
		this.Ref(),
	)
}

// ClearAppBadgeFunc returns the method "WorkerNavigator.clearAppBadge".
func (this WorkerNavigator) ClearAppBadgeFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WorkerNavigatorClearAppBadgeFunc(
			this.Ref(),
		),
	)
}

// ClearAppBadge calls the method "WorkerNavigator.clearAppBadge".
func (this WorkerNavigator) ClearAppBadge() (ret js.Promise[js.Void]) {
	bindings.CallWorkerNavigatorClearAppBadge(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClearAppBadge calls the method "WorkerNavigator.clearAppBadge"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerNavigator) TryClearAppBadge() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerNavigatorClearAppBadge(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Self() (ret WorkerGlobalScope, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeSelf(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Location returns the value of property "WorkerGlobalScope.location".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Location() (ret WorkerLocation, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeLocation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Navigator returns the value of property "WorkerGlobalScope.navigator".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Navigator() (ret WorkerNavigator, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeNavigator(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Fonts returns the value of property "WorkerGlobalScope.fonts".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Fonts() (ret FontFaceSet, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeFonts(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Origin returns the value of property "WorkerGlobalScope.origin".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsSecureContext returns the value of property "WorkerGlobalScope.isSecureContext".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) IsSecureContext() (ret bool, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeIsSecureContext(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CrossOriginIsolated returns the value of property "WorkerGlobalScope.crossOriginIsolated".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) CrossOriginIsolated() (ret bool, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeCrossOriginIsolated(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Scheduler returns the value of property "WorkerGlobalScope.scheduler".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Scheduler() (ret Scheduler, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeScheduler(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TrustedTypes returns the value of property "WorkerGlobalScope.trustedTypes".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) TrustedTypes() (ret TrustedTypePolicyFactory, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeTrustedTypes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Caches returns the value of property "WorkerGlobalScope.caches".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Caches() (ret CacheStorage, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeCaches(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Crypto returns the value of property "WorkerGlobalScope.crypto".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Crypto() (ret Crypto, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeCrypto(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IndexedDB returns the value of property "WorkerGlobalScope.indexedDB".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) IndexedDB() (ret IDBFactory, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeIndexedDB(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Performance returns the value of property "WorkerGlobalScope.performance".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Performance() (ret Performance, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopePerformance(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasImportScripts returns true if the method "WorkerGlobalScope.importScripts" exists.
func (this WorkerGlobalScope) HasImportScripts() bool {
	return js.True == bindings.HasWorkerGlobalScopeImportScripts(
		this.Ref(),
	)
}

// ImportScriptsFunc returns the method "WorkerGlobalScope.importScripts".
func (this WorkerGlobalScope) ImportScriptsFunc() (fn js.Func[func(urls ...js.String)]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeImportScriptsFunc(
			this.Ref(),
		),
	)
}

// ImportScripts calls the method "WorkerGlobalScope.importScripts".
func (this WorkerGlobalScope) ImportScripts(urls ...js.String) (ret js.Void) {
	bindings.CallWorkerGlobalScopeImportScripts(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(urls),
		js.SizeU(len(urls)),
	)

	return
}

// TryImportScripts calls the method "WorkerGlobalScope.importScripts"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryImportScripts(urls ...js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeImportScripts(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(urls),
		js.SizeU(len(urls)),
	)

	return
}

// HasReportError returns true if the method "WorkerGlobalScope.reportError" exists.
func (this WorkerGlobalScope) HasReportError() bool {
	return js.True == bindings.HasWorkerGlobalScopeReportError(
		this.Ref(),
	)
}

// ReportErrorFunc returns the method "WorkerGlobalScope.reportError".
func (this WorkerGlobalScope) ReportErrorFunc() (fn js.Func[func(e js.Any)]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeReportErrorFunc(
			this.Ref(),
		),
	)
}

// ReportError calls the method "WorkerGlobalScope.reportError".
func (this WorkerGlobalScope) ReportError(e js.Any) (ret js.Void) {
	bindings.CallWorkerGlobalScopeReportError(
		this.Ref(), js.Pointer(&ret),
		e.Ref(),
	)

	return
}

// TryReportError calls the method "WorkerGlobalScope.reportError"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryReportError(e js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeReportError(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		e.Ref(),
	)

	return
}

// HasBtoa returns true if the method "WorkerGlobalScope.btoa" exists.
func (this WorkerGlobalScope) HasBtoa() bool {
	return js.True == bindings.HasWorkerGlobalScopeBtoa(
		this.Ref(),
	)
}

// BtoaFunc returns the method "WorkerGlobalScope.btoa".
func (this WorkerGlobalScope) BtoaFunc() (fn js.Func[func(data js.String) js.String]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeBtoaFunc(
			this.Ref(),
		),
	)
}

// Btoa calls the method "WorkerGlobalScope.btoa".
func (this WorkerGlobalScope) Btoa(data js.String) (ret js.String) {
	bindings.CallWorkerGlobalScopeBtoa(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryBtoa calls the method "WorkerGlobalScope.btoa"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryBtoa(data js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeBtoa(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasAtob returns true if the method "WorkerGlobalScope.atob" exists.
func (this WorkerGlobalScope) HasAtob() bool {
	return js.True == bindings.HasWorkerGlobalScopeAtob(
		this.Ref(),
	)
}

// AtobFunc returns the method "WorkerGlobalScope.atob".
func (this WorkerGlobalScope) AtobFunc() (fn js.Func[func(data js.String) js.String]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeAtobFunc(
			this.Ref(),
		),
	)
}

// Atob calls the method "WorkerGlobalScope.atob".
func (this WorkerGlobalScope) Atob(data js.String) (ret js.String) {
	bindings.CallWorkerGlobalScopeAtob(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryAtob calls the method "WorkerGlobalScope.atob"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryAtob(data js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeAtob(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasSetTimeout returns true if the method "WorkerGlobalScope.setTimeout" exists.
func (this WorkerGlobalScope) HasSetTimeout() bool {
	return js.True == bindings.HasWorkerGlobalScopeSetTimeout(
		this.Ref(),
	)
}

// SetTimeoutFunc returns the method "WorkerGlobalScope.setTimeout".
func (this WorkerGlobalScope) SetTimeoutFunc() (fn js.Func[func(handler TimerHandler, timeout int32, arguments ...js.Any) int32]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeSetTimeoutFunc(
			this.Ref(),
		),
	)
}

// SetTimeout calls the method "WorkerGlobalScope.setTimeout".
func (this WorkerGlobalScope) SetTimeout(handler TimerHandler, timeout int32, arguments ...js.Any) (ret int32) {
	bindings.CallWorkerGlobalScopeSetTimeout(
		this.Ref(), js.Pointer(&ret),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// TrySetTimeout calls the method "WorkerGlobalScope.setTimeout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TrySetTimeout(handler TimerHandler, timeout int32, arguments ...js.Any) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeSetTimeout(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// HasSetTimeout1 returns true if the method "WorkerGlobalScope.setTimeout" exists.
func (this WorkerGlobalScope) HasSetTimeout1() bool {
	return js.True == bindings.HasWorkerGlobalScopeSetTimeout1(
		this.Ref(),
	)
}

// SetTimeout1Func returns the method "WorkerGlobalScope.setTimeout".
func (this WorkerGlobalScope) SetTimeout1Func() (fn js.Func[func(handler TimerHandler) int32]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeSetTimeout1Func(
			this.Ref(),
		),
	)
}

// SetTimeout1 calls the method "WorkerGlobalScope.setTimeout".
func (this WorkerGlobalScope) SetTimeout1(handler TimerHandler) (ret int32) {
	bindings.CallWorkerGlobalScopeSetTimeout1(
		this.Ref(), js.Pointer(&ret),
		handler.Ref(),
	)

	return
}

// TrySetTimeout1 calls the method "WorkerGlobalScope.setTimeout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TrySetTimeout1(handler TimerHandler) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeSetTimeout1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
	)

	return
}

// HasClearTimeout returns true if the method "WorkerGlobalScope.clearTimeout" exists.
func (this WorkerGlobalScope) HasClearTimeout() bool {
	return js.True == bindings.HasWorkerGlobalScopeClearTimeout(
		this.Ref(),
	)
}

// ClearTimeoutFunc returns the method "WorkerGlobalScope.clearTimeout".
func (this WorkerGlobalScope) ClearTimeoutFunc() (fn js.Func[func(id int32)]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeClearTimeoutFunc(
			this.Ref(),
		),
	)
}

// ClearTimeout calls the method "WorkerGlobalScope.clearTimeout".
func (this WorkerGlobalScope) ClearTimeout(id int32) (ret js.Void) {
	bindings.CallWorkerGlobalScopeClearTimeout(
		this.Ref(), js.Pointer(&ret),
		int32(id),
	)

	return
}

// TryClearTimeout calls the method "WorkerGlobalScope.clearTimeout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryClearTimeout(id int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeClearTimeout(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
	)

	return
}

// HasClearTimeout1 returns true if the method "WorkerGlobalScope.clearTimeout" exists.
func (this WorkerGlobalScope) HasClearTimeout1() bool {
	return js.True == bindings.HasWorkerGlobalScopeClearTimeout1(
		this.Ref(),
	)
}

// ClearTimeout1Func returns the method "WorkerGlobalScope.clearTimeout".
func (this WorkerGlobalScope) ClearTimeout1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeClearTimeout1Func(
			this.Ref(),
		),
	)
}

// ClearTimeout1 calls the method "WorkerGlobalScope.clearTimeout".
func (this WorkerGlobalScope) ClearTimeout1() (ret js.Void) {
	bindings.CallWorkerGlobalScopeClearTimeout1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClearTimeout1 calls the method "WorkerGlobalScope.clearTimeout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryClearTimeout1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeClearTimeout1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetInterval returns true if the method "WorkerGlobalScope.setInterval" exists.
func (this WorkerGlobalScope) HasSetInterval() bool {
	return js.True == bindings.HasWorkerGlobalScopeSetInterval(
		this.Ref(),
	)
}

// SetIntervalFunc returns the method "WorkerGlobalScope.setInterval".
func (this WorkerGlobalScope) SetIntervalFunc() (fn js.Func[func(handler TimerHandler, timeout int32, arguments ...js.Any) int32]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeSetIntervalFunc(
			this.Ref(),
		),
	)
}

// SetInterval calls the method "WorkerGlobalScope.setInterval".
func (this WorkerGlobalScope) SetInterval(handler TimerHandler, timeout int32, arguments ...js.Any) (ret int32) {
	bindings.CallWorkerGlobalScopeSetInterval(
		this.Ref(), js.Pointer(&ret),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// TrySetInterval calls the method "WorkerGlobalScope.setInterval"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TrySetInterval(handler TimerHandler, timeout int32, arguments ...js.Any) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeSetInterval(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// HasSetInterval1 returns true if the method "WorkerGlobalScope.setInterval" exists.
func (this WorkerGlobalScope) HasSetInterval1() bool {
	return js.True == bindings.HasWorkerGlobalScopeSetInterval1(
		this.Ref(),
	)
}

// SetInterval1Func returns the method "WorkerGlobalScope.setInterval".
func (this WorkerGlobalScope) SetInterval1Func() (fn js.Func[func(handler TimerHandler) int32]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeSetInterval1Func(
			this.Ref(),
		),
	)
}

// SetInterval1 calls the method "WorkerGlobalScope.setInterval".
func (this WorkerGlobalScope) SetInterval1(handler TimerHandler) (ret int32) {
	bindings.CallWorkerGlobalScopeSetInterval1(
		this.Ref(), js.Pointer(&ret),
		handler.Ref(),
	)

	return
}

// TrySetInterval1 calls the method "WorkerGlobalScope.setInterval"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TrySetInterval1(handler TimerHandler) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeSetInterval1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
	)

	return
}

// HasClearInterval returns true if the method "WorkerGlobalScope.clearInterval" exists.
func (this WorkerGlobalScope) HasClearInterval() bool {
	return js.True == bindings.HasWorkerGlobalScopeClearInterval(
		this.Ref(),
	)
}

// ClearIntervalFunc returns the method "WorkerGlobalScope.clearInterval".
func (this WorkerGlobalScope) ClearIntervalFunc() (fn js.Func[func(id int32)]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeClearIntervalFunc(
			this.Ref(),
		),
	)
}

// ClearInterval calls the method "WorkerGlobalScope.clearInterval".
func (this WorkerGlobalScope) ClearInterval(id int32) (ret js.Void) {
	bindings.CallWorkerGlobalScopeClearInterval(
		this.Ref(), js.Pointer(&ret),
		int32(id),
	)

	return
}

// TryClearInterval calls the method "WorkerGlobalScope.clearInterval"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryClearInterval(id int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeClearInterval(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
	)

	return
}

// HasClearInterval1 returns true if the method "WorkerGlobalScope.clearInterval" exists.
func (this WorkerGlobalScope) HasClearInterval1() bool {
	return js.True == bindings.HasWorkerGlobalScopeClearInterval1(
		this.Ref(),
	)
}

// ClearInterval1Func returns the method "WorkerGlobalScope.clearInterval".
func (this WorkerGlobalScope) ClearInterval1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeClearInterval1Func(
			this.Ref(),
		),
	)
}

// ClearInterval1 calls the method "WorkerGlobalScope.clearInterval".
func (this WorkerGlobalScope) ClearInterval1() (ret js.Void) {
	bindings.CallWorkerGlobalScopeClearInterval1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClearInterval1 calls the method "WorkerGlobalScope.clearInterval"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryClearInterval1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeClearInterval1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasQueueMicrotask returns true if the method "WorkerGlobalScope.queueMicrotask" exists.
func (this WorkerGlobalScope) HasQueueMicrotask() bool {
	return js.True == bindings.HasWorkerGlobalScopeQueueMicrotask(
		this.Ref(),
	)
}

// QueueMicrotaskFunc returns the method "WorkerGlobalScope.queueMicrotask".
func (this WorkerGlobalScope) QueueMicrotaskFunc() (fn js.Func[func(callback js.Func[func()])]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeQueueMicrotaskFunc(
			this.Ref(),
		),
	)
}

// QueueMicrotask calls the method "WorkerGlobalScope.queueMicrotask".
func (this WorkerGlobalScope) QueueMicrotask(callback js.Func[func()]) (ret js.Void) {
	bindings.CallWorkerGlobalScopeQueueMicrotask(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryQueueMicrotask calls the method "WorkerGlobalScope.queueMicrotask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryQueueMicrotask(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeQueueMicrotask(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasCreateImageBitmap returns true if the method "WorkerGlobalScope.createImageBitmap" exists.
func (this WorkerGlobalScope) HasCreateImageBitmap() bool {
	return js.True == bindings.HasWorkerGlobalScopeCreateImageBitmap(
		this.Ref(),
	)
}

// CreateImageBitmapFunc returns the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) CreateImageBitmapFunc() (fn js.Func[func(image ImageBitmapSource, options ImageBitmapOptions) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeCreateImageBitmapFunc(
			this.Ref(),
		),
	)
}

// CreateImageBitmap calls the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) CreateImageBitmap(image ImageBitmapSource, options ImageBitmapOptions) (ret js.Promise[ImageBitmap]) {
	bindings.CallWorkerGlobalScopeCreateImageBitmap(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryCreateImageBitmap calls the method "WorkerGlobalScope.createImageBitmap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryCreateImageBitmap(image ImageBitmapSource, options ImageBitmapOptions) (ret js.Promise[ImageBitmap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeCreateImageBitmap(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasCreateImageBitmap1 returns true if the method "WorkerGlobalScope.createImageBitmap" exists.
func (this WorkerGlobalScope) HasCreateImageBitmap1() bool {
	return js.True == bindings.HasWorkerGlobalScopeCreateImageBitmap1(
		this.Ref(),
	)
}

// CreateImageBitmap1Func returns the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) CreateImageBitmap1Func() (fn js.Func[func(image ImageBitmapSource) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeCreateImageBitmap1Func(
			this.Ref(),
		),
	)
}

// CreateImageBitmap1 calls the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) CreateImageBitmap1(image ImageBitmapSource) (ret js.Promise[ImageBitmap]) {
	bindings.CallWorkerGlobalScopeCreateImageBitmap1(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
	)

	return
}

// TryCreateImageBitmap1 calls the method "WorkerGlobalScope.createImageBitmap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryCreateImageBitmap1(image ImageBitmapSource) (ret js.Promise[ImageBitmap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeCreateImageBitmap1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
	)

	return
}

// HasCreateImageBitmap2 returns true if the method "WorkerGlobalScope.createImageBitmap" exists.
func (this WorkerGlobalScope) HasCreateImageBitmap2() bool {
	return js.True == bindings.HasWorkerGlobalScopeCreateImageBitmap2(
		this.Ref(),
	)
}

// CreateImageBitmap2Func returns the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) CreateImageBitmap2Func() (fn js.Func[func(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32, options ImageBitmapOptions) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeCreateImageBitmap2Func(
			this.Ref(),
		),
	)
}

// CreateImageBitmap2 calls the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) CreateImageBitmap2(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32, options ImageBitmapOptions) (ret js.Promise[ImageBitmap]) {
	bindings.CallWorkerGlobalScopeCreateImageBitmap2(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
		js.Pointer(&options),
	)

	return
}

// TryCreateImageBitmap2 calls the method "WorkerGlobalScope.createImageBitmap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryCreateImageBitmap2(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32, options ImageBitmapOptions) (ret js.Promise[ImageBitmap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeCreateImageBitmap2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
		js.Pointer(&options),
	)

	return
}

// HasCreateImageBitmap3 returns true if the method "WorkerGlobalScope.createImageBitmap" exists.
func (this WorkerGlobalScope) HasCreateImageBitmap3() bool {
	return js.True == bindings.HasWorkerGlobalScopeCreateImageBitmap3(
		this.Ref(),
	)
}

// CreateImageBitmap3Func returns the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) CreateImageBitmap3Func() (fn js.Func[func(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeCreateImageBitmap3Func(
			this.Ref(),
		),
	)
}

// CreateImageBitmap3 calls the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) CreateImageBitmap3(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32) (ret js.Promise[ImageBitmap]) {
	bindings.CallWorkerGlobalScopeCreateImageBitmap3(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
	)

	return
}

// TryCreateImageBitmap3 calls the method "WorkerGlobalScope.createImageBitmap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryCreateImageBitmap3(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32) (ret js.Promise[ImageBitmap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeCreateImageBitmap3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
	)

	return
}

// HasStructuredClone returns true if the method "WorkerGlobalScope.structuredClone" exists.
func (this WorkerGlobalScope) HasStructuredClone() bool {
	return js.True == bindings.HasWorkerGlobalScopeStructuredClone(
		this.Ref(),
	)
}

// StructuredCloneFunc returns the method "WorkerGlobalScope.structuredClone".
func (this WorkerGlobalScope) StructuredCloneFunc() (fn js.Func[func(value js.Any, options StructuredSerializeOptions) js.Any]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeStructuredCloneFunc(
			this.Ref(),
		),
	)
}

// StructuredClone calls the method "WorkerGlobalScope.structuredClone".
func (this WorkerGlobalScope) StructuredClone(value js.Any, options StructuredSerializeOptions) (ret js.Any) {
	bindings.CallWorkerGlobalScopeStructuredClone(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryStructuredClone calls the method "WorkerGlobalScope.structuredClone"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryStructuredClone(value js.Any, options StructuredSerializeOptions) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeStructuredClone(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasStructuredClone1 returns true if the method "WorkerGlobalScope.structuredClone" exists.
func (this WorkerGlobalScope) HasStructuredClone1() bool {
	return js.True == bindings.HasWorkerGlobalScopeStructuredClone1(
		this.Ref(),
	)
}

// StructuredClone1Func returns the method "WorkerGlobalScope.structuredClone".
func (this WorkerGlobalScope) StructuredClone1Func() (fn js.Func[func(value js.Any) js.Any]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeStructuredClone1Func(
			this.Ref(),
		),
	)
}

// StructuredClone1 calls the method "WorkerGlobalScope.structuredClone".
func (this WorkerGlobalScope) StructuredClone1(value js.Any) (ret js.Any) {
	bindings.CallWorkerGlobalScopeStructuredClone1(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryStructuredClone1 calls the method "WorkerGlobalScope.structuredClone"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryStructuredClone1(value js.Any) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeStructuredClone1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasFetch returns true if the method "WorkerGlobalScope.fetch" exists.
func (this WorkerGlobalScope) HasFetch() bool {
	return js.True == bindings.HasWorkerGlobalScopeFetch(
		this.Ref(),
	)
}

// FetchFunc returns the method "WorkerGlobalScope.fetch".
func (this WorkerGlobalScope) FetchFunc() (fn js.Func[func(input RequestInfo, init RequestInit) js.Promise[Response]]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeFetchFunc(
			this.Ref(),
		),
	)
}

// Fetch calls the method "WorkerGlobalScope.fetch".
func (this WorkerGlobalScope) Fetch(input RequestInfo, init RequestInit) (ret js.Promise[Response]) {
	bindings.CallWorkerGlobalScopeFetch(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&init),
	)

	return
}

// TryFetch calls the method "WorkerGlobalScope.fetch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryFetch(input RequestInfo, init RequestInit) (ret js.Promise[Response], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeFetch(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&init),
	)

	return
}

// HasFetch1 returns true if the method "WorkerGlobalScope.fetch" exists.
func (this WorkerGlobalScope) HasFetch1() bool {
	return js.True == bindings.HasWorkerGlobalScopeFetch1(
		this.Ref(),
	)
}

// Fetch1Func returns the method "WorkerGlobalScope.fetch".
func (this WorkerGlobalScope) Fetch1Func() (fn js.Func[func(input RequestInfo) js.Promise[Response]]) {
	return fn.FromRef(
		bindings.WorkerGlobalScopeFetch1Func(
			this.Ref(),
		),
	)
}

// Fetch1 calls the method "WorkerGlobalScope.fetch".
func (this WorkerGlobalScope) Fetch1(input RequestInfo) (ret js.Promise[Response]) {
	bindings.CallWorkerGlobalScopeFetch1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryFetch1 calls the method "WorkerGlobalScope.fetch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryFetch1(input RequestInfo) (ret js.Promise[Response], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeFetch1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
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

func NewWorkletAnimation(animatorName js.String, effects OneOf_AnimationEffect_ArrayAnimationEffect, timeline AnimationTimeline, options js.Any) (ret WorkletAnimation) {
	ret.ref = bindings.NewWorkletAnimationByWorkletAnimation(
		animatorName.Ref(),
		effects.Ref(),
		timeline.Ref(),
		options.Ref())
	return
}

func NewWorkletAnimationByWorkletAnimation1(animatorName js.String, effects OneOf_AnimationEffect_ArrayAnimationEffect, timeline AnimationTimeline) (ret WorkletAnimation) {
	ret.ref = bindings.NewWorkletAnimationByWorkletAnimation1(
		animatorName.Ref(),
		effects.Ref(),
		timeline.Ref())
	return
}

func NewWorkletAnimationByWorkletAnimation2(animatorName js.String, effects OneOf_AnimationEffect_ArrayAnimationEffect) (ret WorkletAnimation) {
	ret.ref = bindings.NewWorkletAnimationByWorkletAnimation2(
		animatorName.Ref(),
		effects.Ref())
	return
}

func NewWorkletAnimationByWorkletAnimation3(animatorName js.String) (ret WorkletAnimation) {
	ret.ref = bindings.NewWorkletAnimationByWorkletAnimation3(
		animatorName.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this WorkletAnimation) AnimatorName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkletAnimationAnimatorName(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this WorkletAnimationEffect) LocalTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetWorkletAnimationEffectLocalTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLocalTime sets the value of property "WorkletAnimationEffect.localTime" to val.
//
// It returns false if the property cannot be set.
func (this WorkletAnimationEffect) SetLocalTime(val float64) bool {
	return js.True == bindings.SetWorkletAnimationEffectLocalTime(
		this.Ref(),
		float64(val),
	)
}

// HasGetTiming returns true if the method "WorkletAnimationEffect.getTiming" exists.
func (this WorkletAnimationEffect) HasGetTiming() bool {
	return js.True == bindings.HasWorkletAnimationEffectGetTiming(
		this.Ref(),
	)
}

// GetTimingFunc returns the method "WorkletAnimationEffect.getTiming".
func (this WorkletAnimationEffect) GetTimingFunc() (fn js.Func[func() EffectTiming]) {
	return fn.FromRef(
		bindings.WorkletAnimationEffectGetTimingFunc(
			this.Ref(),
		),
	)
}

// GetTiming calls the method "WorkletAnimationEffect.getTiming".
func (this WorkletAnimationEffect) GetTiming() (ret EffectTiming) {
	bindings.CallWorkletAnimationEffectGetTiming(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetTiming calls the method "WorkletAnimationEffect.getTiming"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkletAnimationEffect) TryGetTiming() (ret EffectTiming, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkletAnimationEffectGetTiming(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetComputedTiming returns true if the method "WorkletAnimationEffect.getComputedTiming" exists.
func (this WorkletAnimationEffect) HasGetComputedTiming() bool {
	return js.True == bindings.HasWorkletAnimationEffectGetComputedTiming(
		this.Ref(),
	)
}

// GetComputedTimingFunc returns the method "WorkletAnimationEffect.getComputedTiming".
func (this WorkletAnimationEffect) GetComputedTimingFunc() (fn js.Func[func() ComputedEffectTiming]) {
	return fn.FromRef(
		bindings.WorkletAnimationEffectGetComputedTimingFunc(
			this.Ref(),
		),
	)
}

// GetComputedTiming calls the method "WorkletAnimationEffect.getComputedTiming".
func (this WorkletAnimationEffect) GetComputedTiming() (ret ComputedEffectTiming) {
	bindings.CallWorkletAnimationEffectGetComputedTiming(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetComputedTiming calls the method "WorkletAnimationEffect.getComputedTiming"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkletAnimationEffect) TryGetComputedTiming() (ret ComputedEffectTiming, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkletAnimationEffectGetComputedTiming(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// HasGetChildren returns true if the method "WorkletGroupEffect.getChildren" exists.
func (this WorkletGroupEffect) HasGetChildren() bool {
	return js.True == bindings.HasWorkletGroupEffectGetChildren(
		this.Ref(),
	)
}

// GetChildrenFunc returns the method "WorkletGroupEffect.getChildren".
func (this WorkletGroupEffect) GetChildrenFunc() (fn js.Func[func() js.Array[WorkletAnimationEffect]]) {
	return fn.FromRef(
		bindings.WorkletGroupEffectGetChildrenFunc(
			this.Ref(),
		),
	)
}

// GetChildren calls the method "WorkletGroupEffect.getChildren".
func (this WorkletGroupEffect) GetChildren() (ret js.Array[WorkletAnimationEffect]) {
	bindings.CallWorkletGroupEffectGetChildren(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetChildren calls the method "WorkletGroupEffect.getChildren"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkletGroupEffect) TryGetChildren() (ret js.Array[WorkletAnimationEffect], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkletGroupEffectGetChildren(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) ReadyState() (ret uint16, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestReadyState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Timeout returns the value of property "XMLHttpRequest.timeout".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) Timeout() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestTimeout(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTimeout sets the value of property "XMLHttpRequest.timeout" to val.
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
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) WithCredentials() (ret bool, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestWithCredentials(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWithCredentials sets the value of property "XMLHttpRequest.withCredentials" to val.
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
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) Upload() (ret XMLHttpRequestUpload, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestUpload(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ResponseURL returns the value of property "XMLHttpRequest.responseURL".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) ResponseURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestResponseURL(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Status returns the value of property "XMLHttpRequest.status".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) Status() (ret uint16, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestStatus(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StatusText returns the value of property "XMLHttpRequest.statusText".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) StatusText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestStatusText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ResponseType returns the value of property "XMLHttpRequest.responseType".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) ResponseType() (ret XMLHttpRequestResponseType, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestResponseType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetResponseType sets the value of property "XMLHttpRequest.responseType" to val.
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
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) Response() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestResponse(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ResponseText returns the value of property "XMLHttpRequest.responseText".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) ResponseText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestResponseText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ResponseXML returns the value of property "XMLHttpRequest.responseXML".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) ResponseXML() (ret Document, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestResponseXML(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasOpen returns true if the method "XMLHttpRequest.open" exists.
func (this XMLHttpRequest) HasOpen() bool {
	return js.True == bindings.HasXMLHttpRequestOpen(
		this.Ref(),
	)
}

// OpenFunc returns the method "XMLHttpRequest.open".
func (this XMLHttpRequest) OpenFunc() (fn js.Func[func(method js.String, url js.String)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestOpenFunc(
			this.Ref(),
		),
	)
}

// Open calls the method "XMLHttpRequest.open".
func (this XMLHttpRequest) Open(method js.String, url js.String) (ret js.Void) {
	bindings.CallXMLHttpRequestOpen(
		this.Ref(), js.Pointer(&ret),
		method.Ref(),
		url.Ref(),
	)

	return
}

// TryOpen calls the method "XMLHttpRequest.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TryOpen(method js.String, url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestOpen(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		method.Ref(),
		url.Ref(),
	)

	return
}

// HasOpen1 returns true if the method "XMLHttpRequest.open" exists.
func (this XMLHttpRequest) HasOpen1() bool {
	return js.True == bindings.HasXMLHttpRequestOpen1(
		this.Ref(),
	)
}

// Open1Func returns the method "XMLHttpRequest.open".
func (this XMLHttpRequest) Open1Func() (fn js.Func[func(method js.String, url js.String, async bool, username js.String, password js.String)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestOpen1Func(
			this.Ref(),
		),
	)
}

// Open1 calls the method "XMLHttpRequest.open".
func (this XMLHttpRequest) Open1(method js.String, url js.String, async bool, username js.String, password js.String) (ret js.Void) {
	bindings.CallXMLHttpRequestOpen1(
		this.Ref(), js.Pointer(&ret),
		method.Ref(),
		url.Ref(),
		js.Bool(bool(async)),
		username.Ref(),
		password.Ref(),
	)

	return
}

// TryOpen1 calls the method "XMLHttpRequest.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TryOpen1(method js.String, url js.String, async bool, username js.String, password js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestOpen1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		method.Ref(),
		url.Ref(),
		js.Bool(bool(async)),
		username.Ref(),
		password.Ref(),
	)

	return
}

// HasOpen2 returns true if the method "XMLHttpRequest.open" exists.
func (this XMLHttpRequest) HasOpen2() bool {
	return js.True == bindings.HasXMLHttpRequestOpen2(
		this.Ref(),
	)
}

// Open2Func returns the method "XMLHttpRequest.open".
func (this XMLHttpRequest) Open2Func() (fn js.Func[func(method js.String, url js.String, async bool, username js.String)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestOpen2Func(
			this.Ref(),
		),
	)
}

// Open2 calls the method "XMLHttpRequest.open".
func (this XMLHttpRequest) Open2(method js.String, url js.String, async bool, username js.String) (ret js.Void) {
	bindings.CallXMLHttpRequestOpen2(
		this.Ref(), js.Pointer(&ret),
		method.Ref(),
		url.Ref(),
		js.Bool(bool(async)),
		username.Ref(),
	)

	return
}

// TryOpen2 calls the method "XMLHttpRequest.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TryOpen2(method js.String, url js.String, async bool, username js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestOpen2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		method.Ref(),
		url.Ref(),
		js.Bool(bool(async)),
		username.Ref(),
	)

	return
}

// HasOpen3 returns true if the method "XMLHttpRequest.open" exists.
func (this XMLHttpRequest) HasOpen3() bool {
	return js.True == bindings.HasXMLHttpRequestOpen3(
		this.Ref(),
	)
}

// Open3Func returns the method "XMLHttpRequest.open".
func (this XMLHttpRequest) Open3Func() (fn js.Func[func(method js.String, url js.String, async bool)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestOpen3Func(
			this.Ref(),
		),
	)
}

// Open3 calls the method "XMLHttpRequest.open".
func (this XMLHttpRequest) Open3(method js.String, url js.String, async bool) (ret js.Void) {
	bindings.CallXMLHttpRequestOpen3(
		this.Ref(), js.Pointer(&ret),
		method.Ref(),
		url.Ref(),
		js.Bool(bool(async)),
	)

	return
}

// TryOpen3 calls the method "XMLHttpRequest.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TryOpen3(method js.String, url js.String, async bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestOpen3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		method.Ref(),
		url.Ref(),
		js.Bool(bool(async)),
	)

	return
}

// HasSetRequestHeader returns true if the method "XMLHttpRequest.setRequestHeader" exists.
func (this XMLHttpRequest) HasSetRequestHeader() bool {
	return js.True == bindings.HasXMLHttpRequestSetRequestHeader(
		this.Ref(),
	)
}

// SetRequestHeaderFunc returns the method "XMLHttpRequest.setRequestHeader".
func (this XMLHttpRequest) SetRequestHeaderFunc() (fn js.Func[func(name js.String, value js.String)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestSetRequestHeaderFunc(
			this.Ref(),
		),
	)
}

// SetRequestHeader calls the method "XMLHttpRequest.setRequestHeader".
func (this XMLHttpRequest) SetRequestHeader(name js.String, value js.String) (ret js.Void) {
	bindings.CallXMLHttpRequestSetRequestHeader(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		value.Ref(),
	)

	return
}

// TrySetRequestHeader calls the method "XMLHttpRequest.setRequestHeader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TrySetRequestHeader(name js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestSetRequestHeader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
	)

	return
}

// HasSend returns true if the method "XMLHttpRequest.send" exists.
func (this XMLHttpRequest) HasSend() bool {
	return js.True == bindings.HasXMLHttpRequestSend(
		this.Ref(),
	)
}

// SendFunc returns the method "XMLHttpRequest.send".
func (this XMLHttpRequest) SendFunc() (fn js.Func[func(body OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestSendFunc(
			this.Ref(),
		),
	)
}

// Send calls the method "XMLHttpRequest.send".
func (this XMLHttpRequest) Send(body OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) (ret js.Void) {
	bindings.CallXMLHttpRequestSend(
		this.Ref(), js.Pointer(&ret),
		body.Ref(),
	)

	return
}

// TrySend calls the method "XMLHttpRequest.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TrySend(body OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestSend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		body.Ref(),
	)

	return
}

// HasSend1 returns true if the method "XMLHttpRequest.send" exists.
func (this XMLHttpRequest) HasSend1() bool {
	return js.True == bindings.HasXMLHttpRequestSend1(
		this.Ref(),
	)
}

// Send1Func returns the method "XMLHttpRequest.send".
func (this XMLHttpRequest) Send1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.XMLHttpRequestSend1Func(
			this.Ref(),
		),
	)
}

// Send1 calls the method "XMLHttpRequest.send".
func (this XMLHttpRequest) Send1() (ret js.Void) {
	bindings.CallXMLHttpRequestSend1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySend1 calls the method "XMLHttpRequest.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TrySend1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestSend1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasAbort returns true if the method "XMLHttpRequest.abort" exists.
func (this XMLHttpRequest) HasAbort() bool {
	return js.True == bindings.HasXMLHttpRequestAbort(
		this.Ref(),
	)
}

// AbortFunc returns the method "XMLHttpRequest.abort".
func (this XMLHttpRequest) AbortFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.XMLHttpRequestAbortFunc(
			this.Ref(),
		),
	)
}

// Abort calls the method "XMLHttpRequest.abort".
func (this XMLHttpRequest) Abort() (ret js.Void) {
	bindings.CallXMLHttpRequestAbort(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAbort calls the method "XMLHttpRequest.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TryAbort() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestAbort(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetResponseHeader returns true if the method "XMLHttpRequest.getResponseHeader" exists.
func (this XMLHttpRequest) HasGetResponseHeader() bool {
	return js.True == bindings.HasXMLHttpRequestGetResponseHeader(
		this.Ref(),
	)
}

// GetResponseHeaderFunc returns the method "XMLHttpRequest.getResponseHeader".
func (this XMLHttpRequest) GetResponseHeaderFunc() (fn js.Func[func(name js.String) js.String]) {
	return fn.FromRef(
		bindings.XMLHttpRequestGetResponseHeaderFunc(
			this.Ref(),
		),
	)
}

// GetResponseHeader calls the method "XMLHttpRequest.getResponseHeader".
func (this XMLHttpRequest) GetResponseHeader(name js.String) (ret js.String) {
	bindings.CallXMLHttpRequestGetResponseHeader(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetResponseHeader calls the method "XMLHttpRequest.getResponseHeader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TryGetResponseHeader(name js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestGetResponseHeader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasGetAllResponseHeaders returns true if the method "XMLHttpRequest.getAllResponseHeaders" exists.
func (this XMLHttpRequest) HasGetAllResponseHeaders() bool {
	return js.True == bindings.HasXMLHttpRequestGetAllResponseHeaders(
		this.Ref(),
	)
}

// GetAllResponseHeadersFunc returns the method "XMLHttpRequest.getAllResponseHeaders".
func (this XMLHttpRequest) GetAllResponseHeadersFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.XMLHttpRequestGetAllResponseHeadersFunc(
			this.Ref(),
		),
	)
}

// GetAllResponseHeaders calls the method "XMLHttpRequest.getAllResponseHeaders".
func (this XMLHttpRequest) GetAllResponseHeaders() (ret js.String) {
	bindings.CallXMLHttpRequestGetAllResponseHeaders(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAllResponseHeaders calls the method "XMLHttpRequest.getAllResponseHeaders"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TryGetAllResponseHeaders() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestGetAllResponseHeaders(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasOverrideMimeType returns true if the method "XMLHttpRequest.overrideMimeType" exists.
func (this XMLHttpRequest) HasOverrideMimeType() bool {
	return js.True == bindings.HasXMLHttpRequestOverrideMimeType(
		this.Ref(),
	)
}

// OverrideMimeTypeFunc returns the method "XMLHttpRequest.overrideMimeType".
func (this XMLHttpRequest) OverrideMimeTypeFunc() (fn js.Func[func(mime js.String)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestOverrideMimeTypeFunc(
			this.Ref(),
		),
	)
}

// OverrideMimeType calls the method "XMLHttpRequest.overrideMimeType".
func (this XMLHttpRequest) OverrideMimeType(mime js.String) (ret js.Void) {
	bindings.CallXMLHttpRequestOverrideMimeType(
		this.Ref(), js.Pointer(&ret),
		mime.Ref(),
	)

	return
}

// TryOverrideMimeType calls the method "XMLHttpRequest.overrideMimeType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TryOverrideMimeType(mime js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestOverrideMimeType(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		mime.Ref(),
	)

	return
}

// HasSetAttributionReporting returns true if the method "XMLHttpRequest.setAttributionReporting" exists.
func (this XMLHttpRequest) HasSetAttributionReporting() bool {
	return js.True == bindings.HasXMLHttpRequestSetAttributionReporting(
		this.Ref(),
	)
}

// SetAttributionReportingFunc returns the method "XMLHttpRequest.setAttributionReporting".
func (this XMLHttpRequest) SetAttributionReportingFunc() (fn js.Func[func(options AttributionReportingRequestOptions)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestSetAttributionReportingFunc(
			this.Ref(),
		),
	)
}

// SetAttributionReporting calls the method "XMLHttpRequest.setAttributionReporting".
func (this XMLHttpRequest) SetAttributionReporting(options AttributionReportingRequestOptions) (ret js.Void) {
	bindings.CallXMLHttpRequestSetAttributionReporting(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TrySetAttributionReporting calls the method "XMLHttpRequest.setAttributionReporting"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TrySetAttributionReporting(options AttributionReportingRequestOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestSetAttributionReporting(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasSetPrivateToken returns true if the method "XMLHttpRequest.setPrivateToken" exists.
func (this XMLHttpRequest) HasSetPrivateToken() bool {
	return js.True == bindings.HasXMLHttpRequestSetPrivateToken(
		this.Ref(),
	)
}

// SetPrivateTokenFunc returns the method "XMLHttpRequest.setPrivateToken".
func (this XMLHttpRequest) SetPrivateTokenFunc() (fn js.Func[func(privateToken PrivateToken)]) {
	return fn.FromRef(
		bindings.XMLHttpRequestSetPrivateTokenFunc(
			this.Ref(),
		),
	)
}

// SetPrivateToken calls the method "XMLHttpRequest.setPrivateToken".
func (this XMLHttpRequest) SetPrivateToken(privateToken PrivateToken) (ret js.Void) {
	bindings.CallXMLHttpRequestSetPrivateToken(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&privateToken),
	)

	return
}

// TrySetPrivateToken calls the method "XMLHttpRequest.setPrivateToken"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TrySetPrivateToken(privateToken PrivateToken) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestSetPrivateToken(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&privateToken),
	)

	return
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

// HasSerializeToString returns true if the method "XMLSerializer.serializeToString" exists.
func (this XMLSerializer) HasSerializeToString() bool {
	return js.True == bindings.HasXMLSerializerSerializeToString(
		this.Ref(),
	)
}

// SerializeToStringFunc returns the method "XMLSerializer.serializeToString".
func (this XMLSerializer) SerializeToStringFunc() (fn js.Func[func(root Node) js.String]) {
	return fn.FromRef(
		bindings.XMLSerializerSerializeToStringFunc(
			this.Ref(),
		),
	)
}

// SerializeToString calls the method "XMLSerializer.serializeToString".
func (this XMLSerializer) SerializeToString(root Node) (ret js.String) {
	bindings.CallXMLSerializerSerializeToString(
		this.Ref(), js.Pointer(&ret),
		root.Ref(),
	)

	return
}

// TrySerializeToString calls the method "XMLSerializer.serializeToString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLSerializer) TrySerializeToString(root Node) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLSerializerSerializeToString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		root.Ref(),
	)

	return
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

// HasCreateExpression returns true if the method "XPathEvaluator.createExpression" exists.
func (this XPathEvaluator) HasCreateExpression() bool {
	return js.True == bindings.HasXPathEvaluatorCreateExpression(
		this.Ref(),
	)
}

// CreateExpressionFunc returns the method "XPathEvaluator.createExpression".
func (this XPathEvaluator) CreateExpressionFunc() (fn js.Func[func(expression js.String, resolver js.Func[func(prefix js.String) js.String]) XPathExpression]) {
	return fn.FromRef(
		bindings.XPathEvaluatorCreateExpressionFunc(
			this.Ref(),
		),
	)
}

// CreateExpression calls the method "XPathEvaluator.createExpression".
func (this XPathEvaluator) CreateExpression(expression js.String, resolver js.Func[func(prefix js.String) js.String]) (ret XPathExpression) {
	bindings.CallXPathEvaluatorCreateExpression(
		this.Ref(), js.Pointer(&ret),
		expression.Ref(),
		resolver.Ref(),
	)

	return
}

// TryCreateExpression calls the method "XPathEvaluator.createExpression"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XPathEvaluator) TryCreateExpression(expression js.String, resolver js.Func[func(prefix js.String) js.String]) (ret XPathExpression, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathEvaluatorCreateExpression(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		resolver.Ref(),
	)

	return
}

// HasCreateExpression1 returns true if the method "XPathEvaluator.createExpression" exists.
func (this XPathEvaluator) HasCreateExpression1() bool {
	return js.True == bindings.HasXPathEvaluatorCreateExpression1(
		this.Ref(),
	)
}

// CreateExpression1Func returns the method "XPathEvaluator.createExpression".
func (this XPathEvaluator) CreateExpression1Func() (fn js.Func[func(expression js.String) XPathExpression]) {
	return fn.FromRef(
		bindings.XPathEvaluatorCreateExpression1Func(
			this.Ref(),
		),
	)
}

// CreateExpression1 calls the method "XPathEvaluator.createExpression".
func (this XPathEvaluator) CreateExpression1(expression js.String) (ret XPathExpression) {
	bindings.CallXPathEvaluatorCreateExpression1(
		this.Ref(), js.Pointer(&ret),
		expression.Ref(),
	)

	return
}

// TryCreateExpression1 calls the method "XPathEvaluator.createExpression"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XPathEvaluator) TryCreateExpression1(expression js.String) (ret XPathExpression, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathEvaluatorCreateExpression1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
	)

	return
}

// HasCreateNSResolver returns true if the method "XPathEvaluator.createNSResolver" exists.
func (this XPathEvaluator) HasCreateNSResolver() bool {
	return js.True == bindings.HasXPathEvaluatorCreateNSResolver(
		this.Ref(),
	)
}

// CreateNSResolverFunc returns the method "XPathEvaluator.createNSResolver".
func (this XPathEvaluator) CreateNSResolverFunc() (fn js.Func[func(nodeResolver Node) Node]) {
	return fn.FromRef(
		bindings.XPathEvaluatorCreateNSResolverFunc(
			this.Ref(),
		),
	)
}

// CreateNSResolver calls the method "XPathEvaluator.createNSResolver".
func (this XPathEvaluator) CreateNSResolver(nodeResolver Node) (ret Node) {
	bindings.CallXPathEvaluatorCreateNSResolver(
		this.Ref(), js.Pointer(&ret),
		nodeResolver.Ref(),
	)

	return
}

// TryCreateNSResolver calls the method "XPathEvaluator.createNSResolver"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XPathEvaluator) TryCreateNSResolver(nodeResolver Node) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathEvaluatorCreateNSResolver(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		nodeResolver.Ref(),
	)

	return
}

// HasEvaluate returns true if the method "XPathEvaluator.evaluate" exists.
func (this XPathEvaluator) HasEvaluate() bool {
	return js.True == bindings.HasXPathEvaluatorEvaluate(
		this.Ref(),
	)
}

// EvaluateFunc returns the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) EvaluateFunc() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16, result XPathResult) XPathResult]) {
	return fn.FromRef(
		bindings.XPathEvaluatorEvaluateFunc(
			this.Ref(),
		),
	)
}

// Evaluate calls the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) Evaluate(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16, result XPathResult) (ret XPathResult) {
	bindings.CallXPathEvaluatorEvaluate(
		this.Ref(), js.Pointer(&ret),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
		result.Ref(),
	)

	return
}

// TryEvaluate calls the method "XPathEvaluator.evaluate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XPathEvaluator) TryEvaluate(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16, result XPathResult) (ret XPathResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathEvaluatorEvaluate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
		result.Ref(),
	)

	return
}

// HasEvaluate1 returns true if the method "XPathEvaluator.evaluate" exists.
func (this XPathEvaluator) HasEvaluate1() bool {
	return js.True == bindings.HasXPathEvaluatorEvaluate1(
		this.Ref(),
	)
}

// Evaluate1Func returns the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) Evaluate1Func() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16) XPathResult]) {
	return fn.FromRef(
		bindings.XPathEvaluatorEvaluate1Func(
			this.Ref(),
		),
	)
}

// Evaluate1 calls the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) Evaluate1(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16) (ret XPathResult) {
	bindings.CallXPathEvaluatorEvaluate1(
		this.Ref(), js.Pointer(&ret),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
	)

	return
}

// TryEvaluate1 calls the method "XPathEvaluator.evaluate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XPathEvaluator) TryEvaluate1(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16) (ret XPathResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathEvaluatorEvaluate1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
	)

	return
}

// HasEvaluate2 returns true if the method "XPathEvaluator.evaluate" exists.
func (this XPathEvaluator) HasEvaluate2() bool {
	return js.True == bindings.HasXPathEvaluatorEvaluate2(
		this.Ref(),
	)
}

// Evaluate2Func returns the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) Evaluate2Func() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String]) XPathResult]) {
	return fn.FromRef(
		bindings.XPathEvaluatorEvaluate2Func(
			this.Ref(),
		),
	)
}

// Evaluate2 calls the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) Evaluate2(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String]) (ret XPathResult) {
	bindings.CallXPathEvaluatorEvaluate2(
		this.Ref(), js.Pointer(&ret),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
	)

	return
}

// TryEvaluate2 calls the method "XPathEvaluator.evaluate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XPathEvaluator) TryEvaluate2(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String]) (ret XPathResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathEvaluatorEvaluate2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
	)

	return
}

// HasEvaluate3 returns true if the method "XPathEvaluator.evaluate" exists.
func (this XPathEvaluator) HasEvaluate3() bool {
	return js.True == bindings.HasXPathEvaluatorEvaluate3(
		this.Ref(),
	)
}

// Evaluate3Func returns the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) Evaluate3Func() (fn js.Func[func(expression js.String, contextNode Node) XPathResult]) {
	return fn.FromRef(
		bindings.XPathEvaluatorEvaluate3Func(
			this.Ref(),
		),
	)
}

// Evaluate3 calls the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) Evaluate3(expression js.String, contextNode Node) (ret XPathResult) {
	bindings.CallXPathEvaluatorEvaluate3(
		this.Ref(), js.Pointer(&ret),
		expression.Ref(),
		contextNode.Ref(),
	)

	return
}

// TryEvaluate3 calls the method "XPathEvaluator.evaluate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XPathEvaluator) TryEvaluate3(expression js.String, contextNode Node) (ret XPathResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathEvaluatorEvaluate3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this XRBoundedReferenceSpace) BoundsGeometry() (ret js.FrozenArray[DOMPointReadOnly], ok bool) {
	ok = js.True == bindings.GetXRBoundedReferenceSpaceBoundsGeometry(
		this.Ref(), js.Pointer(&ret),
	)
	return
}
