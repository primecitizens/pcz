// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
	this.ref.Once()
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
	this.ref.Free()
}

// StatusMessage returns the value of property "WebGLContextEvent.statusMessage".
//
// It returns ok=false if there is no such property.
func (this WebGLContextEvent) StatusMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWebGLContextEventStatusMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

type WebGLObject struct {
	ref js.Ref
}

func (this WebGLObject) Once() WebGLObject {
	this.ref.Once()
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
	this.ref.Free()
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
	this.ref.Once()
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
	this.ref.Free()
}

// Url returns the value of property "WebSocket.url".
//
// It returns ok=false if there is no such property.
func (this WebSocket) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWebSocketUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ReadyState returns the value of property "WebSocket.readyState".
//
// It returns ok=false if there is no such property.
func (this WebSocket) ReadyState() (ret uint16, ok bool) {
	ok = js.True == bindings.GetWebSocketReadyState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BufferedAmount returns the value of property "WebSocket.bufferedAmount".
//
// It returns ok=false if there is no such property.
func (this WebSocket) BufferedAmount() (ret uint64, ok bool) {
	ok = js.True == bindings.GetWebSocketBufferedAmount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Extensions returns the value of property "WebSocket.extensions".
//
// It returns ok=false if there is no such property.
func (this WebSocket) Extensions() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWebSocketExtensions(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Protocol returns the value of property "WebSocket.protocol".
//
// It returns ok=false if there is no such property.
func (this WebSocket) Protocol() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWebSocketProtocol(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BinaryType returns the value of property "WebSocket.binaryType".
//
// It returns ok=false if there is no such property.
func (this WebSocket) BinaryType() (ret BinaryType, ok bool) {
	ok = js.True == bindings.GetWebSocketBinaryType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBinaryType sets the value of property "WebSocket.binaryType" to val.
//
// It returns false if the property cannot be set.
func (this WebSocket) SetBinaryType(val BinaryType) bool {
	return js.True == bindings.SetWebSocketBinaryType(
		this.ref,
		uint32(val),
	)
}

// HasFuncClose returns true if the method "WebSocket.close" exists.
func (this WebSocket) HasFuncClose() bool {
	return js.True == bindings.HasFuncWebSocketClose(
		this.ref,
	)
}

// FuncClose returns the method "WebSocket.close".
func (this WebSocket) FuncClose() (fn js.Func[func(code uint16, reason js.String)]) {
	bindings.FuncWebSocketClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "WebSocket.close".
func (this WebSocket) Close(code uint16, reason js.String) (ret js.Void) {
	bindings.CallWebSocketClose(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(code),
		reason.Ref(),
	)

	return
}

// HasFuncClose1 returns true if the method "WebSocket.close" exists.
func (this WebSocket) HasFuncClose1() bool {
	return js.True == bindings.HasFuncWebSocketClose1(
		this.ref,
	)
}

// FuncClose1 returns the method "WebSocket.close".
func (this WebSocket) FuncClose1() (fn js.Func[func(code uint16)]) {
	bindings.FuncWebSocketClose1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close1 calls the method "WebSocket.close".
func (this WebSocket) Close1(code uint16) (ret js.Void) {
	bindings.CallWebSocketClose1(
		this.ref, js.Pointer(&ret),
		uint32(code),
	)

	return
}

// TryClose1 calls the method "WebSocket.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebSocket) TryClose1(code uint16) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebSocketClose1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(code),
	)

	return
}

// HasFuncClose2 returns true if the method "WebSocket.close" exists.
func (this WebSocket) HasFuncClose2() bool {
	return js.True == bindings.HasFuncWebSocketClose2(
		this.ref,
	)
}

// FuncClose2 returns the method "WebSocket.close".
func (this WebSocket) FuncClose2() (fn js.Func[func()]) {
	bindings.FuncWebSocketClose2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close2 calls the method "WebSocket.close".
func (this WebSocket) Close2() (ret js.Void) {
	bindings.CallWebSocketClose2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose2 calls the method "WebSocket.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebSocket) TryClose2() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebSocketClose2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSend returns true if the method "WebSocket.send" exists.
func (this WebSocket) HasFuncSend() bool {
	return js.True == bindings.HasFuncWebSocketSend(
		this.ref,
	)
}

// FuncSend returns the method "WebSocket.send".
func (this WebSocket) FuncSend() (fn js.Func[func(data OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String)]) {
	bindings.FuncWebSocketSend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Send calls the method "WebSocket.send".
func (this WebSocket) Send(data OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) (ret js.Void) {
	bindings.CallWebSocketSend(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySend calls the method "WebSocket.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebSocket) TrySend(data OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebSocketSend(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *WebTransportHash) UpdateFrom(ref js.Ref) {
	bindings.WebTransportHashJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WebTransportHash) Update(ref js.Ref) {
	bindings.WebTransportHashJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WebTransportHash) FreeMembers(recursive bool) {
	js.Free(
		p.Algorithm.Ref(),
		p.Value.Ref(),
	)
	p.Algorithm = p.Algorithm.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
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
func (p *WebTransportOptions) UpdateFrom(ref js.Ref) {
	bindings.WebTransportOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WebTransportOptions) Update(ref js.Ref) {
	bindings.WebTransportOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WebTransportOptions) FreeMembers(recursive bool) {
	js.Free(
		p.ServerCertificateHashes.Ref(),
	)
	p.ServerCertificateHashes = p.ServerCertificateHashes.FromRef(js.Undefined)
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
func (p *WebTransportDatagramStats) UpdateFrom(ref js.Ref) {
	bindings.WebTransportDatagramStatsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WebTransportDatagramStats) Update(ref js.Ref) {
	bindings.WebTransportDatagramStatsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WebTransportDatagramStats) FreeMembers(recursive bool) {
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
func (p *WebTransportStats) UpdateFrom(ref js.Ref) {
	bindings.WebTransportStatsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WebTransportStats) Update(ref js.Ref) {
	bindings.WebTransportStatsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WebTransportStats) FreeMembers(recursive bool) {
	if recursive {
		p.Datagrams.FreeMembers(true)
	}
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
func (p *WebTransportCloseInfo) UpdateFrom(ref js.Ref) {
	bindings.WebTransportCloseInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WebTransportCloseInfo) Update(ref js.Ref) {
	bindings.WebTransportCloseInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WebTransportCloseInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Reason.Ref(),
	)
	p.Reason = p.Reason.FromRef(js.Undefined)
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
func (p *WebTransportReceiveStreamStats) UpdateFrom(ref js.Ref) {
	bindings.WebTransportReceiveStreamStatsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WebTransportReceiveStreamStats) Update(ref js.Ref) {
	bindings.WebTransportReceiveStreamStatsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WebTransportReceiveStreamStats) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncGetStats returns true if the method "WebTransportReceiveStream.getStats" exists.
func (this WebTransportReceiveStream) HasFuncGetStats() bool {
	return js.True == bindings.HasFuncWebTransportReceiveStreamGetStats(
		this.ref,
	)
}

// FuncGetStats returns the method "WebTransportReceiveStream.getStats".
func (this WebTransportReceiveStream) FuncGetStats() (fn js.Func[func() js.Promise[WebTransportReceiveStreamStats]]) {
	bindings.FuncWebTransportReceiveStreamGetStats(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetStats calls the method "WebTransportReceiveStream.getStats".
func (this WebTransportReceiveStream) GetStats() (ret js.Promise[WebTransportReceiveStreamStats]) {
	bindings.CallWebTransportReceiveStreamGetStats(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetStats calls the method "WebTransportReceiveStream.getStats"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransportReceiveStream) TryGetStats() (ret js.Promise[WebTransportReceiveStreamStats], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportReceiveStreamGetStats(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *WebTransportSendStreamStats) UpdateFrom(ref js.Ref) {
	bindings.WebTransportSendStreamStatsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WebTransportSendStreamStats) Update(ref js.Ref) {
	bindings.WebTransportSendStreamStatsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WebTransportSendStreamStats) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// SendOrder returns the value of property "WebTransportSendStream.sendOrder".
//
// It returns ok=false if there is no such property.
func (this WebTransportSendStream) SendOrder() (ret int64, ok bool) {
	ok = js.True == bindings.GetWebTransportSendStreamSendOrder(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSendOrder sets the value of property "WebTransportSendStream.sendOrder" to val.
//
// It returns false if the property cannot be set.
func (this WebTransportSendStream) SetSendOrder(val int64) bool {
	return js.True == bindings.SetWebTransportSendStreamSendOrder(
		this.ref,
		float64(val),
	)
}

// HasFuncGetStats returns true if the method "WebTransportSendStream.getStats" exists.
func (this WebTransportSendStream) HasFuncGetStats() bool {
	return js.True == bindings.HasFuncWebTransportSendStreamGetStats(
		this.ref,
	)
}

// FuncGetStats returns the method "WebTransportSendStream.getStats".
func (this WebTransportSendStream) FuncGetStats() (fn js.Func[func() js.Promise[WebTransportSendStreamStats]]) {
	bindings.FuncWebTransportSendStreamGetStats(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetStats calls the method "WebTransportSendStream.getStats".
func (this WebTransportSendStream) GetStats() (ret js.Promise[WebTransportSendStreamStats]) {
	bindings.CallWebTransportSendStreamGetStats(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetStats calls the method "WebTransportSendStream.getStats"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransportSendStream) TryGetStats() (ret js.Promise[WebTransportSendStreamStats], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportSendStreamGetStats(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type WebTransportBidirectionalStream struct {
	ref js.Ref
}

func (this WebTransportBidirectionalStream) Once() WebTransportBidirectionalStream {
	this.ref.Once()
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
	this.ref.Free()
}

// Readable returns the value of property "WebTransportBidirectionalStream.readable".
//
// It returns ok=false if there is no such property.
func (this WebTransportBidirectionalStream) Readable() (ret WebTransportReceiveStream, ok bool) {
	ok = js.True == bindings.GetWebTransportBidirectionalStreamReadable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Writable returns the value of property "WebTransportBidirectionalStream.writable".
//
// It returns ok=false if there is no such property.
func (this WebTransportBidirectionalStream) Writable() (ret WebTransportSendStream, ok bool) {
	ok = js.True == bindings.GetWebTransportBidirectionalStreamWritable(
		this.ref, js.Pointer(&ret),
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
func (p *WebTransportSendStreamOptions) UpdateFrom(ref js.Ref) {
	bindings.WebTransportSendStreamOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WebTransportSendStreamOptions) Update(ref js.Ref) {
	bindings.WebTransportSendStreamOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WebTransportSendStreamOptions) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// Readable returns the value of property "WebTransportDatagramDuplexStream.readable".
//
// It returns ok=false if there is no such property.
func (this WebTransportDatagramDuplexStream) Readable() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagramDuplexStreamReadable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Writable returns the value of property "WebTransportDatagramDuplexStream.writable".
//
// It returns ok=false if there is no such property.
func (this WebTransportDatagramDuplexStream) Writable() (ret WritableStream, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagramDuplexStreamWritable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxDatagramSize returns the value of property "WebTransportDatagramDuplexStream.maxDatagramSize".
//
// It returns ok=false if there is no such property.
func (this WebTransportDatagramDuplexStream) MaxDatagramSize() (ret uint32, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagramDuplexStreamMaxDatagramSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IncomingMaxAge returns the value of property "WebTransportDatagramDuplexStream.incomingMaxAge".
//
// It returns ok=false if there is no such property.
func (this WebTransportDatagramDuplexStream) IncomingMaxAge() (ret float64, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagramDuplexStreamIncomingMaxAge(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIncomingMaxAge sets the value of property "WebTransportDatagramDuplexStream.incomingMaxAge" to val.
//
// It returns false if the property cannot be set.
func (this WebTransportDatagramDuplexStream) SetIncomingMaxAge(val float64) bool {
	return js.True == bindings.SetWebTransportDatagramDuplexStreamIncomingMaxAge(
		this.ref,
		float64(val),
	)
}

// OutgoingMaxAge returns the value of property "WebTransportDatagramDuplexStream.outgoingMaxAge".
//
// It returns ok=false if there is no such property.
func (this WebTransportDatagramDuplexStream) OutgoingMaxAge() (ret float64, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagramDuplexStreamOutgoingMaxAge(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetOutgoingMaxAge sets the value of property "WebTransportDatagramDuplexStream.outgoingMaxAge" to val.
//
// It returns false if the property cannot be set.
func (this WebTransportDatagramDuplexStream) SetOutgoingMaxAge(val float64) bool {
	return js.True == bindings.SetWebTransportDatagramDuplexStreamOutgoingMaxAge(
		this.ref,
		float64(val),
	)
}

// IncomingHighWaterMark returns the value of property "WebTransportDatagramDuplexStream.incomingHighWaterMark".
//
// It returns ok=false if there is no such property.
func (this WebTransportDatagramDuplexStream) IncomingHighWaterMark() (ret float64, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagramDuplexStreamIncomingHighWaterMark(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIncomingHighWaterMark sets the value of property "WebTransportDatagramDuplexStream.incomingHighWaterMark" to val.
//
// It returns false if the property cannot be set.
func (this WebTransportDatagramDuplexStream) SetIncomingHighWaterMark(val float64) bool {
	return js.True == bindings.SetWebTransportDatagramDuplexStreamIncomingHighWaterMark(
		this.ref,
		float64(val),
	)
}

// OutgoingHighWaterMark returns the value of property "WebTransportDatagramDuplexStream.outgoingHighWaterMark".
//
// It returns ok=false if there is no such property.
func (this WebTransportDatagramDuplexStream) OutgoingHighWaterMark() (ret float64, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagramDuplexStreamOutgoingHighWaterMark(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetOutgoingHighWaterMark sets the value of property "WebTransportDatagramDuplexStream.outgoingHighWaterMark" to val.
//
// It returns false if the property cannot be set.
func (this WebTransportDatagramDuplexStream) SetOutgoingHighWaterMark(val float64) bool {
	return js.True == bindings.SetWebTransportDatagramDuplexStreamOutgoingHighWaterMark(
		this.ref,
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
	this.ref.Once()
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
	this.ref.Free()
}

// Ready returns the value of property "WebTransport.ready".
//
// It returns ok=false if there is no such property.
func (this WebTransport) Ready() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetWebTransportReady(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Reliability returns the value of property "WebTransport.reliability".
//
// It returns ok=false if there is no such property.
func (this WebTransport) Reliability() (ret WebTransportReliabilityMode, ok bool) {
	ok = js.True == bindings.GetWebTransportReliability(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CongestionControl returns the value of property "WebTransport.congestionControl".
//
// It returns ok=false if there is no such property.
func (this WebTransport) CongestionControl() (ret WebTransportCongestionControl, ok bool) {
	ok = js.True == bindings.GetWebTransportCongestionControl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Closed returns the value of property "WebTransport.closed".
//
// It returns ok=false if there is no such property.
func (this WebTransport) Closed() (ret js.Promise[WebTransportCloseInfo], ok bool) {
	ok = js.True == bindings.GetWebTransportClosed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Draining returns the value of property "WebTransport.draining".
//
// It returns ok=false if there is no such property.
func (this WebTransport) Draining() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetWebTransportDraining(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Datagrams returns the value of property "WebTransport.datagrams".
//
// It returns ok=false if there is no such property.
func (this WebTransport) Datagrams() (ret WebTransportDatagramDuplexStream, ok bool) {
	ok = js.True == bindings.GetWebTransportDatagrams(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IncomingBidirectionalStreams returns the value of property "WebTransport.incomingBidirectionalStreams".
//
// It returns ok=false if there is no such property.
func (this WebTransport) IncomingBidirectionalStreams() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetWebTransportIncomingBidirectionalStreams(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IncomingUnidirectionalStreams returns the value of property "WebTransport.incomingUnidirectionalStreams".
//
// It returns ok=false if there is no such property.
func (this WebTransport) IncomingUnidirectionalStreams() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetWebTransportIncomingUnidirectionalStreams(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetStats returns true if the method "WebTransport.getStats" exists.
func (this WebTransport) HasFuncGetStats() bool {
	return js.True == bindings.HasFuncWebTransportGetStats(
		this.ref,
	)
}

// FuncGetStats returns the method "WebTransport.getStats".
func (this WebTransport) FuncGetStats() (fn js.Func[func() js.Promise[WebTransportStats]]) {
	bindings.FuncWebTransportGetStats(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetStats calls the method "WebTransport.getStats".
func (this WebTransport) GetStats() (ret js.Promise[WebTransportStats]) {
	bindings.CallWebTransportGetStats(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetStats calls the method "WebTransport.getStats"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransport) TryGetStats() (ret js.Promise[WebTransportStats], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportGetStats(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClose returns true if the method "WebTransport.close" exists.
func (this WebTransport) HasFuncClose() bool {
	return js.True == bindings.HasFuncWebTransportClose(
		this.ref,
	)
}

// FuncClose returns the method "WebTransport.close".
func (this WebTransport) FuncClose() (fn js.Func[func(closeInfo WebTransportCloseInfo)]) {
	bindings.FuncWebTransportClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "WebTransport.close".
func (this WebTransport) Close(closeInfo WebTransportCloseInfo) (ret js.Void) {
	bindings.CallWebTransportClose(
		this.ref, js.Pointer(&ret),
		js.Pointer(&closeInfo),
	)

	return
}

// TryClose calls the method "WebTransport.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransport) TryClose(closeInfo WebTransportCloseInfo) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&closeInfo),
	)

	return
}

// HasFuncClose1 returns true if the method "WebTransport.close" exists.
func (this WebTransport) HasFuncClose1() bool {
	return js.True == bindings.HasFuncWebTransportClose1(
		this.ref,
	)
}

// FuncClose1 returns the method "WebTransport.close".
func (this WebTransport) FuncClose1() (fn js.Func[func()]) {
	bindings.FuncWebTransportClose1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close1 calls the method "WebTransport.close".
func (this WebTransport) Close1() (ret js.Void) {
	bindings.CallWebTransportClose1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose1 calls the method "WebTransport.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransport) TryClose1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportClose1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateBidirectionalStream returns true if the method "WebTransport.createBidirectionalStream" exists.
func (this WebTransport) HasFuncCreateBidirectionalStream() bool {
	return js.True == bindings.HasFuncWebTransportCreateBidirectionalStream(
		this.ref,
	)
}

// FuncCreateBidirectionalStream returns the method "WebTransport.createBidirectionalStream".
func (this WebTransport) FuncCreateBidirectionalStream() (fn js.Func[func(options WebTransportSendStreamOptions) js.Promise[WebTransportBidirectionalStream]]) {
	bindings.FuncWebTransportCreateBidirectionalStream(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateBidirectionalStream calls the method "WebTransport.createBidirectionalStream".
func (this WebTransport) CreateBidirectionalStream(options WebTransportSendStreamOptions) (ret js.Promise[WebTransportBidirectionalStream]) {
	bindings.CallWebTransportCreateBidirectionalStream(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryCreateBidirectionalStream calls the method "WebTransport.createBidirectionalStream"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransport) TryCreateBidirectionalStream(options WebTransportSendStreamOptions) (ret js.Promise[WebTransportBidirectionalStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportCreateBidirectionalStream(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncCreateBidirectionalStream1 returns true if the method "WebTransport.createBidirectionalStream" exists.
func (this WebTransport) HasFuncCreateBidirectionalStream1() bool {
	return js.True == bindings.HasFuncWebTransportCreateBidirectionalStream1(
		this.ref,
	)
}

// FuncCreateBidirectionalStream1 returns the method "WebTransport.createBidirectionalStream".
func (this WebTransport) FuncCreateBidirectionalStream1() (fn js.Func[func() js.Promise[WebTransportBidirectionalStream]]) {
	bindings.FuncWebTransportCreateBidirectionalStream1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateBidirectionalStream1 calls the method "WebTransport.createBidirectionalStream".
func (this WebTransport) CreateBidirectionalStream1() (ret js.Promise[WebTransportBidirectionalStream]) {
	bindings.CallWebTransportCreateBidirectionalStream1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateBidirectionalStream1 calls the method "WebTransport.createBidirectionalStream"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransport) TryCreateBidirectionalStream1() (ret js.Promise[WebTransportBidirectionalStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportCreateBidirectionalStream1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateUnidirectionalStream returns true if the method "WebTransport.createUnidirectionalStream" exists.
func (this WebTransport) HasFuncCreateUnidirectionalStream() bool {
	return js.True == bindings.HasFuncWebTransportCreateUnidirectionalStream(
		this.ref,
	)
}

// FuncCreateUnidirectionalStream returns the method "WebTransport.createUnidirectionalStream".
func (this WebTransport) FuncCreateUnidirectionalStream() (fn js.Func[func(options WebTransportSendStreamOptions) js.Promise[WebTransportSendStream]]) {
	bindings.FuncWebTransportCreateUnidirectionalStream(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateUnidirectionalStream calls the method "WebTransport.createUnidirectionalStream".
func (this WebTransport) CreateUnidirectionalStream(options WebTransportSendStreamOptions) (ret js.Promise[WebTransportSendStream]) {
	bindings.CallWebTransportCreateUnidirectionalStream(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryCreateUnidirectionalStream calls the method "WebTransport.createUnidirectionalStream"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransport) TryCreateUnidirectionalStream(options WebTransportSendStreamOptions) (ret js.Promise[WebTransportSendStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportCreateUnidirectionalStream(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncCreateUnidirectionalStream1 returns true if the method "WebTransport.createUnidirectionalStream" exists.
func (this WebTransport) HasFuncCreateUnidirectionalStream1() bool {
	return js.True == bindings.HasFuncWebTransportCreateUnidirectionalStream1(
		this.ref,
	)
}

// FuncCreateUnidirectionalStream1 returns the method "WebTransport.createUnidirectionalStream".
func (this WebTransport) FuncCreateUnidirectionalStream1() (fn js.Func[func() js.Promise[WebTransportSendStream]]) {
	bindings.FuncWebTransportCreateUnidirectionalStream1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateUnidirectionalStream1 calls the method "WebTransport.createUnidirectionalStream".
func (this WebTransport) CreateUnidirectionalStream1() (ret js.Promise[WebTransportSendStream]) {
	bindings.CallWebTransportCreateUnidirectionalStream1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateUnidirectionalStream1 calls the method "WebTransport.createUnidirectionalStream"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebTransport) TryCreateUnidirectionalStream1() (ret js.Promise[WebTransportSendStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebTransportCreateUnidirectionalStream1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *WebTransportErrorOptions) UpdateFrom(ref js.Ref) {
	bindings.WebTransportErrorOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WebTransportErrorOptions) Update(ref js.Ref) {
	bindings.WebTransportErrorOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WebTransportErrorOptions) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// Source returns the value of property "WebTransportError.source".
//
// It returns ok=false if there is no such property.
func (this WebTransportError) Source() (ret WebTransportErrorSource, ok bool) {
	ok = js.True == bindings.GetWebTransportErrorSource(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StreamErrorCode returns the value of property "WebTransportError.streamErrorCode".
//
// It returns ok=false if there is no such property.
func (this WebTransportError) StreamErrorCode() (ret uint32, ok bool) {
	ok = js.True == bindings.GetWebTransportErrorStreamErrorCode(
		this.ref, js.Pointer(&ret),
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
func (p *WheelEventInit) UpdateFrom(ref js.Ref) {
	bindings.WheelEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WheelEventInit) Update(ref js.Ref) {
	bindings.WheelEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WheelEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.RelatedTarget.Ref(),
		p.View.Ref(),
	)
	p.RelatedTarget = p.RelatedTarget.FromRef(js.Undefined)
	p.View = p.View.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// DeltaX returns the value of property "WheelEvent.deltaX".
//
// It returns ok=false if there is no such property.
func (this WheelEvent) DeltaX() (ret float64, ok bool) {
	ok = js.True == bindings.GetWheelEventDeltaX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DeltaY returns the value of property "WheelEvent.deltaY".
//
// It returns ok=false if there is no such property.
func (this WheelEvent) DeltaY() (ret float64, ok bool) {
	ok = js.True == bindings.GetWheelEventDeltaY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DeltaZ returns the value of property "WheelEvent.deltaZ".
//
// It returns ok=false if there is no such property.
func (this WheelEvent) DeltaZ() (ret float64, ok bool) {
	ok = js.True == bindings.GetWheelEventDeltaZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DeltaMode returns the value of property "WheelEvent.deltaMode".
//
// It returns ok=false if there is no such property.
func (this WheelEvent) DeltaMode() (ret uint32, ok bool) {
	ok = js.True == bindings.GetWheelEventDeltaMode(
		this.ref, js.Pointer(&ret),
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
func (p *WindowControlsOverlayGeometryChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.WindowControlsOverlayGeometryChangeEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WindowControlsOverlayGeometryChangeEventInit) Update(ref js.Ref) {
	bindings.WindowControlsOverlayGeometryChangeEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WindowControlsOverlayGeometryChangeEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.TitlebarAreaRect.Ref(),
	)
	p.TitlebarAreaRect = p.TitlebarAreaRect.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// TitlebarAreaRect returns the value of property "WindowControlsOverlayGeometryChangeEvent.titlebarAreaRect".
//
// It returns ok=false if there is no such property.
func (this WindowControlsOverlayGeometryChangeEvent) TitlebarAreaRect() (ret DOMRect, ok bool) {
	ok = js.True == bindings.GetWindowControlsOverlayGeometryChangeEventTitlebarAreaRect(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Visible returns the value of property "WindowControlsOverlayGeometryChangeEvent.visible".
//
// It returns ok=false if there is no such property.
func (this WindowControlsOverlayGeometryChangeEvent) Visible() (ret bool, ok bool) {
	ok = js.True == bindings.GetWindowControlsOverlayGeometryChangeEventVisible(
		this.ref, js.Pointer(&ret),
	)
	return
}

type WorkerLocation struct {
	ref js.Ref
}

func (this WorkerLocation) Once() WorkerLocation {
	this.ref.Once()
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
	this.ref.Free()
}

// Href returns the value of property "WorkerLocation.href".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Href() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationHref(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Origin returns the value of property "WorkerLocation.origin".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Protocol returns the value of property "WorkerLocation.protocol".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Protocol() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationProtocol(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Host returns the value of property "WorkerLocation.host".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Host() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationHost(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Hostname returns the value of property "WorkerLocation.hostname".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Hostname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationHostname(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Port returns the value of property "WorkerLocation.port".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Port() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationPort(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Pathname returns the value of property "WorkerLocation.pathname".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Pathname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationPathname(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Search returns the value of property "WorkerLocation.search".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Search() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationSearch(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Hash returns the value of property "WorkerLocation.hash".
//
// It returns ok=false if there is no such property.
func (this WorkerLocation) Hash() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerLocationHash(
		this.ref, js.Pointer(&ret),
	)
	return
}

type WorkerNavigator struct {
	ref js.Ref
}

func (this WorkerNavigator) Once() WorkerNavigator {
	this.ref.Once()
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
	this.ref.Free()
}

// Hid returns the value of property "WorkerNavigator.hid".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Hid() (ret HID, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorHid(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Serial returns the value of property "WorkerNavigator.serial".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Serial() (ret Serial, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorSerial(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Permissions returns the value of property "WorkerNavigator.permissions".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Permissions() (ret Permissions, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorPermissions(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Usb returns the value of property "WorkerNavigator.usb".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Usb() (ret USB, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorUsb(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MediaCapabilities returns the value of property "WorkerNavigator.mediaCapabilities".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) MediaCapabilities() (ret MediaCapabilities, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorMediaCapabilities(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ServiceWorker returns the value of property "WorkerNavigator.serviceWorker".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) ServiceWorker() (ret ServiceWorkerContainer, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorServiceWorker(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UserAgentData returns the value of property "WorkerNavigator.userAgentData".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) UserAgentData() (ret NavigatorUAData, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorUserAgentData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DeviceMemory returns the value of property "WorkerNavigator.deviceMemory".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) DeviceMemory() (ret float64, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorDeviceMemory(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Storage returns the value of property "WorkerNavigator.storage".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Storage() (ret StorageManager, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorStorage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StorageBuckets returns the value of property "WorkerNavigator.storageBuckets".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) StorageBuckets() (ret StorageBucketManager, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorStorageBuckets(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Locks returns the value of property "WorkerNavigator.locks".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Locks() (ret LockManager, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorLocks(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AppCodeName returns the value of property "WorkerNavigator.appCodeName".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) AppCodeName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorAppCodeName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AppName returns the value of property "WorkerNavigator.appName".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) AppName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorAppName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AppVersion returns the value of property "WorkerNavigator.appVersion".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) AppVersion() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorAppVersion(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Platform returns the value of property "WorkerNavigator.platform".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Platform() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorPlatform(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Product returns the value of property "WorkerNavigator.product".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Product() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorProduct(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ProductSub returns the value of property "WorkerNavigator.productSub".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) ProductSub() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorProductSub(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UserAgent returns the value of property "WorkerNavigator.userAgent".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) UserAgent() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorUserAgent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Vendor returns the value of property "WorkerNavigator.vendor".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Vendor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorVendor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// VendorSub returns the value of property "WorkerNavigator.vendorSub".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) VendorSub() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorVendorSub(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Oscpu returns the value of property "WorkerNavigator.oscpu".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Oscpu() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorOscpu(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Language returns the value of property "WorkerNavigator.language".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Language() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorLanguage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Languages returns the value of property "WorkerNavigator.languages".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Languages() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorLanguages(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OnLine returns the value of property "WorkerNavigator.onLine".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) OnLine() (ret bool, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorOnLine(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HardwareConcurrency returns the value of property "WorkerNavigator.hardwareConcurrency".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) HardwareConcurrency() (ret uint64, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorHardwareConcurrency(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Connection returns the value of property "WorkerNavigator.connection".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Connection() (ret NetworkInformation, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorConnection(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Ml returns the value of property "WorkerNavigator.ml".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Ml() (ret ML, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorMl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Gpu returns the value of property "WorkerNavigator.gpu".
//
// It returns ok=false if there is no such property.
func (this WorkerNavigator) Gpu() (ret GPU, ok bool) {
	ok = js.True == bindings.GetWorkerNavigatorGpu(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncTaintEnabled returns true if the method "WorkerNavigator.taintEnabled" exists.
func (this WorkerNavigator) HasFuncTaintEnabled() bool {
	return js.True == bindings.HasFuncWorkerNavigatorTaintEnabled(
		this.ref,
	)
}

// FuncTaintEnabled returns the method "WorkerNavigator.taintEnabled".
func (this WorkerNavigator) FuncTaintEnabled() (fn js.Func[func() bool]) {
	bindings.FuncWorkerNavigatorTaintEnabled(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TaintEnabled calls the method "WorkerNavigator.taintEnabled".
func (this WorkerNavigator) TaintEnabled() (ret bool) {
	bindings.CallWorkerNavigatorTaintEnabled(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTaintEnabled calls the method "WorkerNavigator.taintEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerNavigator) TryTaintEnabled() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerNavigatorTaintEnabled(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetAppBadge returns true if the method "WorkerNavigator.setAppBadge" exists.
func (this WorkerNavigator) HasFuncSetAppBadge() bool {
	return js.True == bindings.HasFuncWorkerNavigatorSetAppBadge(
		this.ref,
	)
}

// FuncSetAppBadge returns the method "WorkerNavigator.setAppBadge".
func (this WorkerNavigator) FuncSetAppBadge() (fn js.Func[func(contents uint64) js.Promise[js.Void]]) {
	bindings.FuncWorkerNavigatorSetAppBadge(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetAppBadge calls the method "WorkerNavigator.setAppBadge".
func (this WorkerNavigator) SetAppBadge(contents uint64) (ret js.Promise[js.Void]) {
	bindings.CallWorkerNavigatorSetAppBadge(
		this.ref, js.Pointer(&ret),
		float64(contents),
	)

	return
}

// TrySetAppBadge calls the method "WorkerNavigator.setAppBadge"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerNavigator) TrySetAppBadge(contents uint64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerNavigatorSetAppBadge(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(contents),
	)

	return
}

// HasFuncSetAppBadge1 returns true if the method "WorkerNavigator.setAppBadge" exists.
func (this WorkerNavigator) HasFuncSetAppBadge1() bool {
	return js.True == bindings.HasFuncWorkerNavigatorSetAppBadge1(
		this.ref,
	)
}

// FuncSetAppBadge1 returns the method "WorkerNavigator.setAppBadge".
func (this WorkerNavigator) FuncSetAppBadge1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncWorkerNavigatorSetAppBadge1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetAppBadge1 calls the method "WorkerNavigator.setAppBadge".
func (this WorkerNavigator) SetAppBadge1() (ret js.Promise[js.Void]) {
	bindings.CallWorkerNavigatorSetAppBadge1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySetAppBadge1 calls the method "WorkerNavigator.setAppBadge"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerNavigator) TrySetAppBadge1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerNavigatorSetAppBadge1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClearAppBadge returns true if the method "WorkerNavigator.clearAppBadge" exists.
func (this WorkerNavigator) HasFuncClearAppBadge() bool {
	return js.True == bindings.HasFuncWorkerNavigatorClearAppBadge(
		this.ref,
	)
}

// FuncClearAppBadge returns the method "WorkerNavigator.clearAppBadge".
func (this WorkerNavigator) FuncClearAppBadge() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncWorkerNavigatorClearAppBadge(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearAppBadge calls the method "WorkerNavigator.clearAppBadge".
func (this WorkerNavigator) ClearAppBadge() (ret js.Promise[js.Void]) {
	bindings.CallWorkerNavigatorClearAppBadge(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClearAppBadge calls the method "WorkerNavigator.clearAppBadge"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerNavigator) TryClearAppBadge() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerNavigatorClearAppBadge(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type WorkerGlobalScope struct {
	EventTarget
}

func (this WorkerGlobalScope) Once() WorkerGlobalScope {
	this.ref.Once()
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
	this.ref.Free()
}

// Self returns the value of property "WorkerGlobalScope.self".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Self() (ret WorkerGlobalScope, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeSelf(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Location returns the value of property "WorkerGlobalScope.location".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Location() (ret WorkerLocation, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeLocation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Navigator returns the value of property "WorkerGlobalScope.navigator".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Navigator() (ret WorkerNavigator, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeNavigator(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Fonts returns the value of property "WorkerGlobalScope.fonts".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Fonts() (ret FontFaceSet, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeFonts(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Origin returns the value of property "WorkerGlobalScope.origin".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsSecureContext returns the value of property "WorkerGlobalScope.isSecureContext".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) IsSecureContext() (ret bool, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeIsSecureContext(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CrossOriginIsolated returns the value of property "WorkerGlobalScope.crossOriginIsolated".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) CrossOriginIsolated() (ret bool, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeCrossOriginIsolated(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Scheduler returns the value of property "WorkerGlobalScope.scheduler".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Scheduler() (ret Scheduler, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeScheduler(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TrustedTypes returns the value of property "WorkerGlobalScope.trustedTypes".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) TrustedTypes() (ret TrustedTypePolicyFactory, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeTrustedTypes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Caches returns the value of property "WorkerGlobalScope.caches".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Caches() (ret CacheStorage, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeCaches(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Crypto returns the value of property "WorkerGlobalScope.crypto".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Crypto() (ret Crypto, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeCrypto(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IndexedDB returns the value of property "WorkerGlobalScope.indexedDB".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) IndexedDB() (ret IDBFactory, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopeIndexedDB(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Performance returns the value of property "WorkerGlobalScope.performance".
//
// It returns ok=false if there is no such property.
func (this WorkerGlobalScope) Performance() (ret Performance, ok bool) {
	ok = js.True == bindings.GetWorkerGlobalScopePerformance(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncImportScripts returns true if the method "WorkerGlobalScope.importScripts" exists.
func (this WorkerGlobalScope) HasFuncImportScripts() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeImportScripts(
		this.ref,
	)
}

// FuncImportScripts returns the method "WorkerGlobalScope.importScripts".
func (this WorkerGlobalScope) FuncImportScripts() (fn js.Func[func(urls ...js.String)]) {
	bindings.FuncWorkerGlobalScopeImportScripts(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ImportScripts calls the method "WorkerGlobalScope.importScripts".
func (this WorkerGlobalScope) ImportScripts(urls ...js.String) (ret js.Void) {
	bindings.CallWorkerGlobalScopeImportScripts(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(urls),
		js.SizeU(len(urls)),
	)

	return
}

// HasFuncReportError returns true if the method "WorkerGlobalScope.reportError" exists.
func (this WorkerGlobalScope) HasFuncReportError() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeReportError(
		this.ref,
	)
}

// FuncReportError returns the method "WorkerGlobalScope.reportError".
func (this WorkerGlobalScope) FuncReportError() (fn js.Func[func(e js.Any)]) {
	bindings.FuncWorkerGlobalScopeReportError(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReportError calls the method "WorkerGlobalScope.reportError".
func (this WorkerGlobalScope) ReportError(e js.Any) (ret js.Void) {
	bindings.CallWorkerGlobalScopeReportError(
		this.ref, js.Pointer(&ret),
		e.Ref(),
	)

	return
}

// TryReportError calls the method "WorkerGlobalScope.reportError"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryReportError(e js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeReportError(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		e.Ref(),
	)

	return
}

// HasFuncBtoa returns true if the method "WorkerGlobalScope.btoa" exists.
func (this WorkerGlobalScope) HasFuncBtoa() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeBtoa(
		this.ref,
	)
}

// FuncBtoa returns the method "WorkerGlobalScope.btoa".
func (this WorkerGlobalScope) FuncBtoa() (fn js.Func[func(data js.String) js.String]) {
	bindings.FuncWorkerGlobalScopeBtoa(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Btoa calls the method "WorkerGlobalScope.btoa".
func (this WorkerGlobalScope) Btoa(data js.String) (ret js.String) {
	bindings.CallWorkerGlobalScopeBtoa(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryBtoa calls the method "WorkerGlobalScope.btoa"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryBtoa(data js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeBtoa(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncAtob returns true if the method "WorkerGlobalScope.atob" exists.
func (this WorkerGlobalScope) HasFuncAtob() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeAtob(
		this.ref,
	)
}

// FuncAtob returns the method "WorkerGlobalScope.atob".
func (this WorkerGlobalScope) FuncAtob() (fn js.Func[func(data js.String) js.String]) {
	bindings.FuncWorkerGlobalScopeAtob(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Atob calls the method "WorkerGlobalScope.atob".
func (this WorkerGlobalScope) Atob(data js.String) (ret js.String) {
	bindings.CallWorkerGlobalScopeAtob(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryAtob calls the method "WorkerGlobalScope.atob"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryAtob(data js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeAtob(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncSetTimeout returns true if the method "WorkerGlobalScope.setTimeout" exists.
func (this WorkerGlobalScope) HasFuncSetTimeout() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeSetTimeout(
		this.ref,
	)
}

// FuncSetTimeout returns the method "WorkerGlobalScope.setTimeout".
func (this WorkerGlobalScope) FuncSetTimeout() (fn js.Func[func(handler TimerHandler, timeout int32, arguments ...js.Any) int32]) {
	bindings.FuncWorkerGlobalScopeSetTimeout(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetTimeout calls the method "WorkerGlobalScope.setTimeout".
func (this WorkerGlobalScope) SetTimeout(handler TimerHandler, timeout int32, arguments ...js.Any) (ret int32) {
	bindings.CallWorkerGlobalScopeSetTimeout(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// HasFuncSetTimeout1 returns true if the method "WorkerGlobalScope.setTimeout" exists.
func (this WorkerGlobalScope) HasFuncSetTimeout1() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeSetTimeout1(
		this.ref,
	)
}

// FuncSetTimeout1 returns the method "WorkerGlobalScope.setTimeout".
func (this WorkerGlobalScope) FuncSetTimeout1() (fn js.Func[func(handler TimerHandler) int32]) {
	bindings.FuncWorkerGlobalScopeSetTimeout1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetTimeout1 calls the method "WorkerGlobalScope.setTimeout".
func (this WorkerGlobalScope) SetTimeout1(handler TimerHandler) (ret int32) {
	bindings.CallWorkerGlobalScopeSetTimeout1(
		this.ref, js.Pointer(&ret),
		handler.Ref(),
	)

	return
}

// TrySetTimeout1 calls the method "WorkerGlobalScope.setTimeout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TrySetTimeout1(handler TimerHandler) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeSetTimeout1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
	)

	return
}

// HasFuncClearTimeout returns true if the method "WorkerGlobalScope.clearTimeout" exists.
func (this WorkerGlobalScope) HasFuncClearTimeout() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeClearTimeout(
		this.ref,
	)
}

// FuncClearTimeout returns the method "WorkerGlobalScope.clearTimeout".
func (this WorkerGlobalScope) FuncClearTimeout() (fn js.Func[func(id int32)]) {
	bindings.FuncWorkerGlobalScopeClearTimeout(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearTimeout calls the method "WorkerGlobalScope.clearTimeout".
func (this WorkerGlobalScope) ClearTimeout(id int32) (ret js.Void) {
	bindings.CallWorkerGlobalScopeClearTimeout(
		this.ref, js.Pointer(&ret),
		int32(id),
	)

	return
}

// TryClearTimeout calls the method "WorkerGlobalScope.clearTimeout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryClearTimeout(id int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeClearTimeout(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
	)

	return
}

// HasFuncClearTimeout1 returns true if the method "WorkerGlobalScope.clearTimeout" exists.
func (this WorkerGlobalScope) HasFuncClearTimeout1() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeClearTimeout1(
		this.ref,
	)
}

// FuncClearTimeout1 returns the method "WorkerGlobalScope.clearTimeout".
func (this WorkerGlobalScope) FuncClearTimeout1() (fn js.Func[func()]) {
	bindings.FuncWorkerGlobalScopeClearTimeout1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearTimeout1 calls the method "WorkerGlobalScope.clearTimeout".
func (this WorkerGlobalScope) ClearTimeout1() (ret js.Void) {
	bindings.CallWorkerGlobalScopeClearTimeout1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClearTimeout1 calls the method "WorkerGlobalScope.clearTimeout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryClearTimeout1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeClearTimeout1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetInterval returns true if the method "WorkerGlobalScope.setInterval" exists.
func (this WorkerGlobalScope) HasFuncSetInterval() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeSetInterval(
		this.ref,
	)
}

// FuncSetInterval returns the method "WorkerGlobalScope.setInterval".
func (this WorkerGlobalScope) FuncSetInterval() (fn js.Func[func(handler TimerHandler, timeout int32, arguments ...js.Any) int32]) {
	bindings.FuncWorkerGlobalScopeSetInterval(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetInterval calls the method "WorkerGlobalScope.setInterval".
func (this WorkerGlobalScope) SetInterval(handler TimerHandler, timeout int32, arguments ...js.Any) (ret int32) {
	bindings.CallWorkerGlobalScopeSetInterval(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// HasFuncSetInterval1 returns true if the method "WorkerGlobalScope.setInterval" exists.
func (this WorkerGlobalScope) HasFuncSetInterval1() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeSetInterval1(
		this.ref,
	)
}

// FuncSetInterval1 returns the method "WorkerGlobalScope.setInterval".
func (this WorkerGlobalScope) FuncSetInterval1() (fn js.Func[func(handler TimerHandler) int32]) {
	bindings.FuncWorkerGlobalScopeSetInterval1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetInterval1 calls the method "WorkerGlobalScope.setInterval".
func (this WorkerGlobalScope) SetInterval1(handler TimerHandler) (ret int32) {
	bindings.CallWorkerGlobalScopeSetInterval1(
		this.ref, js.Pointer(&ret),
		handler.Ref(),
	)

	return
}

// TrySetInterval1 calls the method "WorkerGlobalScope.setInterval"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TrySetInterval1(handler TimerHandler) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeSetInterval1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
	)

	return
}

// HasFuncClearInterval returns true if the method "WorkerGlobalScope.clearInterval" exists.
func (this WorkerGlobalScope) HasFuncClearInterval() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeClearInterval(
		this.ref,
	)
}

// FuncClearInterval returns the method "WorkerGlobalScope.clearInterval".
func (this WorkerGlobalScope) FuncClearInterval() (fn js.Func[func(id int32)]) {
	bindings.FuncWorkerGlobalScopeClearInterval(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearInterval calls the method "WorkerGlobalScope.clearInterval".
func (this WorkerGlobalScope) ClearInterval(id int32) (ret js.Void) {
	bindings.CallWorkerGlobalScopeClearInterval(
		this.ref, js.Pointer(&ret),
		int32(id),
	)

	return
}

// TryClearInterval calls the method "WorkerGlobalScope.clearInterval"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryClearInterval(id int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeClearInterval(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
	)

	return
}

// HasFuncClearInterval1 returns true if the method "WorkerGlobalScope.clearInterval" exists.
func (this WorkerGlobalScope) HasFuncClearInterval1() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeClearInterval1(
		this.ref,
	)
}

// FuncClearInterval1 returns the method "WorkerGlobalScope.clearInterval".
func (this WorkerGlobalScope) FuncClearInterval1() (fn js.Func[func()]) {
	bindings.FuncWorkerGlobalScopeClearInterval1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearInterval1 calls the method "WorkerGlobalScope.clearInterval".
func (this WorkerGlobalScope) ClearInterval1() (ret js.Void) {
	bindings.CallWorkerGlobalScopeClearInterval1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClearInterval1 calls the method "WorkerGlobalScope.clearInterval"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryClearInterval1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeClearInterval1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncQueueMicrotask returns true if the method "WorkerGlobalScope.queueMicrotask" exists.
func (this WorkerGlobalScope) HasFuncQueueMicrotask() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeQueueMicrotask(
		this.ref,
	)
}

// FuncQueueMicrotask returns the method "WorkerGlobalScope.queueMicrotask".
func (this WorkerGlobalScope) FuncQueueMicrotask() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncWorkerGlobalScopeQueueMicrotask(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QueueMicrotask calls the method "WorkerGlobalScope.queueMicrotask".
func (this WorkerGlobalScope) QueueMicrotask(callback js.Func[func()]) (ret js.Void) {
	bindings.CallWorkerGlobalScopeQueueMicrotask(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryQueueMicrotask calls the method "WorkerGlobalScope.queueMicrotask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryQueueMicrotask(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeQueueMicrotask(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncCreateImageBitmap returns true if the method "WorkerGlobalScope.createImageBitmap" exists.
func (this WorkerGlobalScope) HasFuncCreateImageBitmap() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeCreateImageBitmap(
		this.ref,
	)
}

// FuncCreateImageBitmap returns the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) FuncCreateImageBitmap() (fn js.Func[func(image ImageBitmapSource, options ImageBitmapOptions) js.Promise[ImageBitmap]]) {
	bindings.FuncWorkerGlobalScopeCreateImageBitmap(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateImageBitmap calls the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) CreateImageBitmap(image ImageBitmapSource, options ImageBitmapOptions) (ret js.Promise[ImageBitmap]) {
	bindings.CallWorkerGlobalScopeCreateImageBitmap(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncCreateImageBitmap1 returns true if the method "WorkerGlobalScope.createImageBitmap" exists.
func (this WorkerGlobalScope) HasFuncCreateImageBitmap1() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeCreateImageBitmap1(
		this.ref,
	)
}

// FuncCreateImageBitmap1 returns the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) FuncCreateImageBitmap1() (fn js.Func[func(image ImageBitmapSource) js.Promise[ImageBitmap]]) {
	bindings.FuncWorkerGlobalScopeCreateImageBitmap1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateImageBitmap1 calls the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) CreateImageBitmap1(image ImageBitmapSource) (ret js.Promise[ImageBitmap]) {
	bindings.CallWorkerGlobalScopeCreateImageBitmap1(
		this.ref, js.Pointer(&ret),
		image.Ref(),
	)

	return
}

// TryCreateImageBitmap1 calls the method "WorkerGlobalScope.createImageBitmap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryCreateImageBitmap1(image ImageBitmapSource) (ret js.Promise[ImageBitmap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeCreateImageBitmap1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
	)

	return
}

// HasFuncCreateImageBitmap2 returns true if the method "WorkerGlobalScope.createImageBitmap" exists.
func (this WorkerGlobalScope) HasFuncCreateImageBitmap2() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeCreateImageBitmap2(
		this.ref,
	)
}

// FuncCreateImageBitmap2 returns the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) FuncCreateImageBitmap2() (fn js.Func[func(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32, options ImageBitmapOptions) js.Promise[ImageBitmap]]) {
	bindings.FuncWorkerGlobalScopeCreateImageBitmap2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateImageBitmap2 calls the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) CreateImageBitmap2(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32, options ImageBitmapOptions) (ret js.Promise[ImageBitmap]) {
	bindings.CallWorkerGlobalScopeCreateImageBitmap2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
		js.Pointer(&options),
	)

	return
}

// HasFuncCreateImageBitmap3 returns true if the method "WorkerGlobalScope.createImageBitmap" exists.
func (this WorkerGlobalScope) HasFuncCreateImageBitmap3() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeCreateImageBitmap3(
		this.ref,
	)
}

// FuncCreateImageBitmap3 returns the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) FuncCreateImageBitmap3() (fn js.Func[func(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32) js.Promise[ImageBitmap]]) {
	bindings.FuncWorkerGlobalScopeCreateImageBitmap3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateImageBitmap3 calls the method "WorkerGlobalScope.createImageBitmap".
func (this WorkerGlobalScope) CreateImageBitmap3(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32) (ret js.Promise[ImageBitmap]) {
	bindings.CallWorkerGlobalScopeCreateImageBitmap3(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
	)

	return
}

// HasFuncStructuredClone returns true if the method "WorkerGlobalScope.structuredClone" exists.
func (this WorkerGlobalScope) HasFuncStructuredClone() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeStructuredClone(
		this.ref,
	)
}

// FuncStructuredClone returns the method "WorkerGlobalScope.structuredClone".
func (this WorkerGlobalScope) FuncStructuredClone() (fn js.Func[func(value js.Any, options StructuredSerializeOptions) js.Any]) {
	bindings.FuncWorkerGlobalScopeStructuredClone(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StructuredClone calls the method "WorkerGlobalScope.structuredClone".
func (this WorkerGlobalScope) StructuredClone(value js.Any, options StructuredSerializeOptions) (ret js.Any) {
	bindings.CallWorkerGlobalScopeStructuredClone(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncStructuredClone1 returns true if the method "WorkerGlobalScope.structuredClone" exists.
func (this WorkerGlobalScope) HasFuncStructuredClone1() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeStructuredClone1(
		this.ref,
	)
}

// FuncStructuredClone1 returns the method "WorkerGlobalScope.structuredClone".
func (this WorkerGlobalScope) FuncStructuredClone1() (fn js.Func[func(value js.Any) js.Any]) {
	bindings.FuncWorkerGlobalScopeStructuredClone1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StructuredClone1 calls the method "WorkerGlobalScope.structuredClone".
func (this WorkerGlobalScope) StructuredClone1(value js.Any) (ret js.Any) {
	bindings.CallWorkerGlobalScopeStructuredClone1(
		this.ref, js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryStructuredClone1 calls the method "WorkerGlobalScope.structuredClone"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryStructuredClone1(value js.Any) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeStructuredClone1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasFuncFetch returns true if the method "WorkerGlobalScope.fetch" exists.
func (this WorkerGlobalScope) HasFuncFetch() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeFetch(
		this.ref,
	)
}

// FuncFetch returns the method "WorkerGlobalScope.fetch".
func (this WorkerGlobalScope) FuncFetch() (fn js.Func[func(input RequestInfo, init RequestInit) js.Promise[Response]]) {
	bindings.FuncWorkerGlobalScopeFetch(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Fetch calls the method "WorkerGlobalScope.fetch".
func (this WorkerGlobalScope) Fetch(input RequestInfo, init RequestInit) (ret js.Promise[Response]) {
	bindings.CallWorkerGlobalScopeFetch(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&init),
	)

	return
}

// HasFuncFetch1 returns true if the method "WorkerGlobalScope.fetch" exists.
func (this WorkerGlobalScope) HasFuncFetch1() bool {
	return js.True == bindings.HasFuncWorkerGlobalScopeFetch1(
		this.ref,
	)
}

// FuncFetch1 returns the method "WorkerGlobalScope.fetch".
func (this WorkerGlobalScope) FuncFetch1() (fn js.Func[func(input RequestInfo) js.Promise[Response]]) {
	bindings.FuncWorkerGlobalScopeFetch1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Fetch1 calls the method "WorkerGlobalScope.fetch".
func (this WorkerGlobalScope) Fetch1(input RequestInfo) (ret js.Promise[Response]) {
	bindings.CallWorkerGlobalScopeFetch1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryFetch1 calls the method "WorkerGlobalScope.fetch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkerGlobalScope) TryFetch1(input RequestInfo) (ret js.Promise[Response], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkerGlobalScopeFetch1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// AnimatorName returns the value of property "WorkletAnimation.animatorName".
//
// It returns ok=false if there is no such property.
func (this WorkletAnimation) AnimatorName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWorkletAnimationAnimatorName(
		this.ref, js.Pointer(&ret),
	)
	return
}

type WorkletAnimationEffect struct {
	ref js.Ref
}

func (this WorkletAnimationEffect) Once() WorkletAnimationEffect {
	this.ref.Once()
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
	this.ref.Free()
}

// LocalTime returns the value of property "WorkletAnimationEffect.localTime".
//
// It returns ok=false if there is no such property.
func (this WorkletAnimationEffect) LocalTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetWorkletAnimationEffectLocalTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLocalTime sets the value of property "WorkletAnimationEffect.localTime" to val.
//
// It returns false if the property cannot be set.
func (this WorkletAnimationEffect) SetLocalTime(val float64) bool {
	return js.True == bindings.SetWorkletAnimationEffectLocalTime(
		this.ref,
		float64(val),
	)
}

// HasFuncGetTiming returns true if the method "WorkletAnimationEffect.getTiming" exists.
func (this WorkletAnimationEffect) HasFuncGetTiming() bool {
	return js.True == bindings.HasFuncWorkletAnimationEffectGetTiming(
		this.ref,
	)
}

// FuncGetTiming returns the method "WorkletAnimationEffect.getTiming".
func (this WorkletAnimationEffect) FuncGetTiming() (fn js.Func[func() EffectTiming]) {
	bindings.FuncWorkletAnimationEffectGetTiming(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetTiming calls the method "WorkletAnimationEffect.getTiming".
func (this WorkletAnimationEffect) GetTiming() (ret EffectTiming) {
	bindings.CallWorkletAnimationEffectGetTiming(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetTiming calls the method "WorkletAnimationEffect.getTiming"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkletAnimationEffect) TryGetTiming() (ret EffectTiming, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkletAnimationEffectGetTiming(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetComputedTiming returns true if the method "WorkletAnimationEffect.getComputedTiming" exists.
func (this WorkletAnimationEffect) HasFuncGetComputedTiming() bool {
	return js.True == bindings.HasFuncWorkletAnimationEffectGetComputedTiming(
		this.ref,
	)
}

// FuncGetComputedTiming returns the method "WorkletAnimationEffect.getComputedTiming".
func (this WorkletAnimationEffect) FuncGetComputedTiming() (fn js.Func[func() ComputedEffectTiming]) {
	bindings.FuncWorkletAnimationEffectGetComputedTiming(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetComputedTiming calls the method "WorkletAnimationEffect.getComputedTiming".
func (this WorkletAnimationEffect) GetComputedTiming() (ret ComputedEffectTiming) {
	bindings.CallWorkletAnimationEffectGetComputedTiming(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetComputedTiming calls the method "WorkletAnimationEffect.getComputedTiming"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkletAnimationEffect) TryGetComputedTiming() (ret ComputedEffectTiming, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkletAnimationEffectGetComputedTiming(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type WorkletGlobalScope struct {
	ref js.Ref
}

func (this WorkletGlobalScope) Once() WorkletGlobalScope {
	this.ref.Once()
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
	this.ref.Free()
}

type WorkletGroupEffect struct {
	ref js.Ref
}

func (this WorkletGroupEffect) Once() WorkletGroupEffect {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncGetChildren returns true if the method "WorkletGroupEffect.getChildren" exists.
func (this WorkletGroupEffect) HasFuncGetChildren() bool {
	return js.True == bindings.HasFuncWorkletGroupEffectGetChildren(
		this.ref,
	)
}

// FuncGetChildren returns the method "WorkletGroupEffect.getChildren".
func (this WorkletGroupEffect) FuncGetChildren() (fn js.Func[func() js.Array[WorkletAnimationEffect]]) {
	bindings.FuncWorkletGroupEffectGetChildren(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetChildren calls the method "WorkletGroupEffect.getChildren".
func (this WorkletGroupEffect) GetChildren() (ret js.Array[WorkletAnimationEffect]) {
	bindings.CallWorkletGroupEffectGetChildren(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetChildren calls the method "WorkletGroupEffect.getChildren"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkletGroupEffect) TryGetChildren() (ret js.Array[WorkletAnimationEffect], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkletGroupEffectGetChildren(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
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
	this.ref.Once()
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
	this.ref.Free()
}

// ReadyState returns the value of property "XMLHttpRequest.readyState".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) ReadyState() (ret uint16, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestReadyState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Timeout returns the value of property "XMLHttpRequest.timeout".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) Timeout() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestTimeout(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTimeout sets the value of property "XMLHttpRequest.timeout" to val.
//
// It returns false if the property cannot be set.
func (this XMLHttpRequest) SetTimeout(val uint32) bool {
	return js.True == bindings.SetXMLHttpRequestTimeout(
		this.ref,
		uint32(val),
	)
}

// WithCredentials returns the value of property "XMLHttpRequest.withCredentials".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) WithCredentials() (ret bool, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestWithCredentials(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWithCredentials sets the value of property "XMLHttpRequest.withCredentials" to val.
//
// It returns false if the property cannot be set.
func (this XMLHttpRequest) SetWithCredentials(val bool) bool {
	return js.True == bindings.SetXMLHttpRequestWithCredentials(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Upload returns the value of property "XMLHttpRequest.upload".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) Upload() (ret XMLHttpRequestUpload, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestUpload(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ResponseURL returns the value of property "XMLHttpRequest.responseURL".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) ResponseURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestResponseURL(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Status returns the value of property "XMLHttpRequest.status".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) Status() (ret uint16, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestStatus(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StatusText returns the value of property "XMLHttpRequest.statusText".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) StatusText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestStatusText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ResponseType returns the value of property "XMLHttpRequest.responseType".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) ResponseType() (ret XMLHttpRequestResponseType, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestResponseType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetResponseType sets the value of property "XMLHttpRequest.responseType" to val.
//
// It returns false if the property cannot be set.
func (this XMLHttpRequest) SetResponseType(val XMLHttpRequestResponseType) bool {
	return js.True == bindings.SetXMLHttpRequestResponseType(
		this.ref,
		uint32(val),
	)
}

// Response returns the value of property "XMLHttpRequest.response".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) Response() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestResponse(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ResponseText returns the value of property "XMLHttpRequest.responseText".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) ResponseText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestResponseText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ResponseXML returns the value of property "XMLHttpRequest.responseXML".
//
// It returns ok=false if there is no such property.
func (this XMLHttpRequest) ResponseXML() (ret Document, ok bool) {
	ok = js.True == bindings.GetXMLHttpRequestResponseXML(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncOpen returns true if the method "XMLHttpRequest.open" exists.
func (this XMLHttpRequest) HasFuncOpen() bool {
	return js.True == bindings.HasFuncXMLHttpRequestOpen(
		this.ref,
	)
}

// FuncOpen returns the method "XMLHttpRequest.open".
func (this XMLHttpRequest) FuncOpen() (fn js.Func[func(method js.String, url js.String)]) {
	bindings.FuncXMLHttpRequestOpen(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open calls the method "XMLHttpRequest.open".
func (this XMLHttpRequest) Open(method js.String, url js.String) (ret js.Void) {
	bindings.CallXMLHttpRequestOpen(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		method.Ref(),
		url.Ref(),
	)

	return
}

// HasFuncOpen1 returns true if the method "XMLHttpRequest.open" exists.
func (this XMLHttpRequest) HasFuncOpen1() bool {
	return js.True == bindings.HasFuncXMLHttpRequestOpen1(
		this.ref,
	)
}

// FuncOpen1 returns the method "XMLHttpRequest.open".
func (this XMLHttpRequest) FuncOpen1() (fn js.Func[func(method js.String, url js.String, async bool, username js.String, password js.String)]) {
	bindings.FuncXMLHttpRequestOpen1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open1 calls the method "XMLHttpRequest.open".
func (this XMLHttpRequest) Open1(method js.String, url js.String, async bool, username js.String, password js.String) (ret js.Void) {
	bindings.CallXMLHttpRequestOpen1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		method.Ref(),
		url.Ref(),
		js.Bool(bool(async)),
		username.Ref(),
		password.Ref(),
	)

	return
}

// HasFuncOpen2 returns true if the method "XMLHttpRequest.open" exists.
func (this XMLHttpRequest) HasFuncOpen2() bool {
	return js.True == bindings.HasFuncXMLHttpRequestOpen2(
		this.ref,
	)
}

// FuncOpen2 returns the method "XMLHttpRequest.open".
func (this XMLHttpRequest) FuncOpen2() (fn js.Func[func(method js.String, url js.String, async bool, username js.String)]) {
	bindings.FuncXMLHttpRequestOpen2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open2 calls the method "XMLHttpRequest.open".
func (this XMLHttpRequest) Open2(method js.String, url js.String, async bool, username js.String) (ret js.Void) {
	bindings.CallXMLHttpRequestOpen2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		method.Ref(),
		url.Ref(),
		js.Bool(bool(async)),
		username.Ref(),
	)

	return
}

// HasFuncOpen3 returns true if the method "XMLHttpRequest.open" exists.
func (this XMLHttpRequest) HasFuncOpen3() bool {
	return js.True == bindings.HasFuncXMLHttpRequestOpen3(
		this.ref,
	)
}

// FuncOpen3 returns the method "XMLHttpRequest.open".
func (this XMLHttpRequest) FuncOpen3() (fn js.Func[func(method js.String, url js.String, async bool)]) {
	bindings.FuncXMLHttpRequestOpen3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open3 calls the method "XMLHttpRequest.open".
func (this XMLHttpRequest) Open3(method js.String, url js.String, async bool) (ret js.Void) {
	bindings.CallXMLHttpRequestOpen3(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		method.Ref(),
		url.Ref(),
		js.Bool(bool(async)),
	)

	return
}

// HasFuncSetRequestHeader returns true if the method "XMLHttpRequest.setRequestHeader" exists.
func (this XMLHttpRequest) HasFuncSetRequestHeader() bool {
	return js.True == bindings.HasFuncXMLHttpRequestSetRequestHeader(
		this.ref,
	)
}

// FuncSetRequestHeader returns the method "XMLHttpRequest.setRequestHeader".
func (this XMLHttpRequest) FuncSetRequestHeader() (fn js.Func[func(name js.String, value js.String)]) {
	bindings.FuncXMLHttpRequestSetRequestHeader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetRequestHeader calls the method "XMLHttpRequest.setRequestHeader".
func (this XMLHttpRequest) SetRequestHeader(name js.String, value js.String) (ret js.Void) {
	bindings.CallXMLHttpRequestSetRequestHeader(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncSend returns true if the method "XMLHttpRequest.send" exists.
func (this XMLHttpRequest) HasFuncSend() bool {
	return js.True == bindings.HasFuncXMLHttpRequestSend(
		this.ref,
	)
}

// FuncSend returns the method "XMLHttpRequest.send".
func (this XMLHttpRequest) FuncSend() (fn js.Func[func(body OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String)]) {
	bindings.FuncXMLHttpRequestSend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Send calls the method "XMLHttpRequest.send".
func (this XMLHttpRequest) Send(body OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) (ret js.Void) {
	bindings.CallXMLHttpRequestSend(
		this.ref, js.Pointer(&ret),
		body.Ref(),
	)

	return
}

// TrySend calls the method "XMLHttpRequest.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TrySend(body OneOf_Document_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestSend(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		body.Ref(),
	)

	return
}

// HasFuncSend1 returns true if the method "XMLHttpRequest.send" exists.
func (this XMLHttpRequest) HasFuncSend1() bool {
	return js.True == bindings.HasFuncXMLHttpRequestSend1(
		this.ref,
	)
}

// FuncSend1 returns the method "XMLHttpRequest.send".
func (this XMLHttpRequest) FuncSend1() (fn js.Func[func()]) {
	bindings.FuncXMLHttpRequestSend1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Send1 calls the method "XMLHttpRequest.send".
func (this XMLHttpRequest) Send1() (ret js.Void) {
	bindings.CallXMLHttpRequestSend1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySend1 calls the method "XMLHttpRequest.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TrySend1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestSend1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAbort returns true if the method "XMLHttpRequest.abort" exists.
func (this XMLHttpRequest) HasFuncAbort() bool {
	return js.True == bindings.HasFuncXMLHttpRequestAbort(
		this.ref,
	)
}

// FuncAbort returns the method "XMLHttpRequest.abort".
func (this XMLHttpRequest) FuncAbort() (fn js.Func[func()]) {
	bindings.FuncXMLHttpRequestAbort(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Abort calls the method "XMLHttpRequest.abort".
func (this XMLHttpRequest) Abort() (ret js.Void) {
	bindings.CallXMLHttpRequestAbort(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAbort calls the method "XMLHttpRequest.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TryAbort() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestAbort(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetResponseHeader returns true if the method "XMLHttpRequest.getResponseHeader" exists.
func (this XMLHttpRequest) HasFuncGetResponseHeader() bool {
	return js.True == bindings.HasFuncXMLHttpRequestGetResponseHeader(
		this.ref,
	)
}

// FuncGetResponseHeader returns the method "XMLHttpRequest.getResponseHeader".
func (this XMLHttpRequest) FuncGetResponseHeader() (fn js.Func[func(name js.String) js.String]) {
	bindings.FuncXMLHttpRequestGetResponseHeader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetResponseHeader calls the method "XMLHttpRequest.getResponseHeader".
func (this XMLHttpRequest) GetResponseHeader(name js.String) (ret js.String) {
	bindings.CallXMLHttpRequestGetResponseHeader(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetResponseHeader calls the method "XMLHttpRequest.getResponseHeader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TryGetResponseHeader(name js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestGetResponseHeader(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncGetAllResponseHeaders returns true if the method "XMLHttpRequest.getAllResponseHeaders" exists.
func (this XMLHttpRequest) HasFuncGetAllResponseHeaders() bool {
	return js.True == bindings.HasFuncXMLHttpRequestGetAllResponseHeaders(
		this.ref,
	)
}

// FuncGetAllResponseHeaders returns the method "XMLHttpRequest.getAllResponseHeaders".
func (this XMLHttpRequest) FuncGetAllResponseHeaders() (fn js.Func[func() js.String]) {
	bindings.FuncXMLHttpRequestGetAllResponseHeaders(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAllResponseHeaders calls the method "XMLHttpRequest.getAllResponseHeaders".
func (this XMLHttpRequest) GetAllResponseHeaders() (ret js.String) {
	bindings.CallXMLHttpRequestGetAllResponseHeaders(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetAllResponseHeaders calls the method "XMLHttpRequest.getAllResponseHeaders"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TryGetAllResponseHeaders() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestGetAllResponseHeaders(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncOverrideMimeType returns true if the method "XMLHttpRequest.overrideMimeType" exists.
func (this XMLHttpRequest) HasFuncOverrideMimeType() bool {
	return js.True == bindings.HasFuncXMLHttpRequestOverrideMimeType(
		this.ref,
	)
}

// FuncOverrideMimeType returns the method "XMLHttpRequest.overrideMimeType".
func (this XMLHttpRequest) FuncOverrideMimeType() (fn js.Func[func(mime js.String)]) {
	bindings.FuncXMLHttpRequestOverrideMimeType(
		this.ref, js.Pointer(&fn),
	)
	return
}

// OverrideMimeType calls the method "XMLHttpRequest.overrideMimeType".
func (this XMLHttpRequest) OverrideMimeType(mime js.String) (ret js.Void) {
	bindings.CallXMLHttpRequestOverrideMimeType(
		this.ref, js.Pointer(&ret),
		mime.Ref(),
	)

	return
}

// TryOverrideMimeType calls the method "XMLHttpRequest.overrideMimeType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TryOverrideMimeType(mime js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestOverrideMimeType(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		mime.Ref(),
	)

	return
}

// HasFuncSetAttributionReporting returns true if the method "XMLHttpRequest.setAttributionReporting" exists.
func (this XMLHttpRequest) HasFuncSetAttributionReporting() bool {
	return js.True == bindings.HasFuncXMLHttpRequestSetAttributionReporting(
		this.ref,
	)
}

// FuncSetAttributionReporting returns the method "XMLHttpRequest.setAttributionReporting".
func (this XMLHttpRequest) FuncSetAttributionReporting() (fn js.Func[func(options AttributionReportingRequestOptions)]) {
	bindings.FuncXMLHttpRequestSetAttributionReporting(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetAttributionReporting calls the method "XMLHttpRequest.setAttributionReporting".
func (this XMLHttpRequest) SetAttributionReporting(options AttributionReportingRequestOptions) (ret js.Void) {
	bindings.CallXMLHttpRequestSetAttributionReporting(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TrySetAttributionReporting calls the method "XMLHttpRequest.setAttributionReporting"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TrySetAttributionReporting(options AttributionReportingRequestOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestSetAttributionReporting(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncSetPrivateToken returns true if the method "XMLHttpRequest.setPrivateToken" exists.
func (this XMLHttpRequest) HasFuncSetPrivateToken() bool {
	return js.True == bindings.HasFuncXMLHttpRequestSetPrivateToken(
		this.ref,
	)
}

// FuncSetPrivateToken returns the method "XMLHttpRequest.setPrivateToken".
func (this XMLHttpRequest) FuncSetPrivateToken() (fn js.Func[func(privateToken PrivateToken)]) {
	bindings.FuncXMLHttpRequestSetPrivateToken(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetPrivateToken calls the method "XMLHttpRequest.setPrivateToken".
func (this XMLHttpRequest) SetPrivateToken(privateToken PrivateToken) (ret js.Void) {
	bindings.CallXMLHttpRequestSetPrivateToken(
		this.ref, js.Pointer(&ret),
		js.Pointer(&privateToken),
	)

	return
}

// TrySetPrivateToken calls the method "XMLHttpRequest.setPrivateToken"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLHttpRequest) TrySetPrivateToken(privateToken PrivateToken) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLHttpRequestSetPrivateToken(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

type XMLSerializer struct {
	ref js.Ref
}

func (this XMLSerializer) Once() XMLSerializer {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncSerializeToString returns true if the method "XMLSerializer.serializeToString" exists.
func (this XMLSerializer) HasFuncSerializeToString() bool {
	return js.True == bindings.HasFuncXMLSerializerSerializeToString(
		this.ref,
	)
}

// FuncSerializeToString returns the method "XMLSerializer.serializeToString".
func (this XMLSerializer) FuncSerializeToString() (fn js.Func[func(root Node) js.String]) {
	bindings.FuncXMLSerializerSerializeToString(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SerializeToString calls the method "XMLSerializer.serializeToString".
func (this XMLSerializer) SerializeToString(root Node) (ret js.String) {
	bindings.CallXMLSerializerSerializeToString(
		this.ref, js.Pointer(&ret),
		root.Ref(),
	)

	return
}

// TrySerializeToString calls the method "XMLSerializer.serializeToString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XMLSerializer) TrySerializeToString(root Node) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXMLSerializerSerializeToString(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		root.Ref(),
	)

	return
}

type XPathEvaluator struct {
	ref js.Ref
}

func (this XPathEvaluator) Once() XPathEvaluator {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncCreateExpression returns true if the method "XPathEvaluator.createExpression" exists.
func (this XPathEvaluator) HasFuncCreateExpression() bool {
	return js.True == bindings.HasFuncXPathEvaluatorCreateExpression(
		this.ref,
	)
}

// FuncCreateExpression returns the method "XPathEvaluator.createExpression".
func (this XPathEvaluator) FuncCreateExpression() (fn js.Func[func(expression js.String, resolver js.Func[func(prefix js.String) js.String]) XPathExpression]) {
	bindings.FuncXPathEvaluatorCreateExpression(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateExpression calls the method "XPathEvaluator.createExpression".
func (this XPathEvaluator) CreateExpression(expression js.String, resolver js.Func[func(prefix js.String) js.String]) (ret XPathExpression) {
	bindings.CallXPathEvaluatorCreateExpression(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		resolver.Ref(),
	)

	return
}

// HasFuncCreateExpression1 returns true if the method "XPathEvaluator.createExpression" exists.
func (this XPathEvaluator) HasFuncCreateExpression1() bool {
	return js.True == bindings.HasFuncXPathEvaluatorCreateExpression1(
		this.ref,
	)
}

// FuncCreateExpression1 returns the method "XPathEvaluator.createExpression".
func (this XPathEvaluator) FuncCreateExpression1() (fn js.Func[func(expression js.String) XPathExpression]) {
	bindings.FuncXPathEvaluatorCreateExpression1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateExpression1 calls the method "XPathEvaluator.createExpression".
func (this XPathEvaluator) CreateExpression1(expression js.String) (ret XPathExpression) {
	bindings.CallXPathEvaluatorCreateExpression1(
		this.ref, js.Pointer(&ret),
		expression.Ref(),
	)

	return
}

// TryCreateExpression1 calls the method "XPathEvaluator.createExpression"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XPathEvaluator) TryCreateExpression1(expression js.String) (ret XPathExpression, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathEvaluatorCreateExpression1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
	)

	return
}

// HasFuncCreateNSResolver returns true if the method "XPathEvaluator.createNSResolver" exists.
func (this XPathEvaluator) HasFuncCreateNSResolver() bool {
	return js.True == bindings.HasFuncXPathEvaluatorCreateNSResolver(
		this.ref,
	)
}

// FuncCreateNSResolver returns the method "XPathEvaluator.createNSResolver".
func (this XPathEvaluator) FuncCreateNSResolver() (fn js.Func[func(nodeResolver Node) Node]) {
	bindings.FuncXPathEvaluatorCreateNSResolver(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateNSResolver calls the method "XPathEvaluator.createNSResolver".
func (this XPathEvaluator) CreateNSResolver(nodeResolver Node) (ret Node) {
	bindings.CallXPathEvaluatorCreateNSResolver(
		this.ref, js.Pointer(&ret),
		nodeResolver.Ref(),
	)

	return
}

// TryCreateNSResolver calls the method "XPathEvaluator.createNSResolver"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XPathEvaluator) TryCreateNSResolver(nodeResolver Node) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathEvaluatorCreateNSResolver(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		nodeResolver.Ref(),
	)

	return
}

// HasFuncEvaluate returns true if the method "XPathEvaluator.evaluate" exists.
func (this XPathEvaluator) HasFuncEvaluate() bool {
	return js.True == bindings.HasFuncXPathEvaluatorEvaluate(
		this.ref,
	)
}

// FuncEvaluate returns the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) FuncEvaluate() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16, result XPathResult) XPathResult]) {
	bindings.FuncXPathEvaluatorEvaluate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Evaluate calls the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) Evaluate(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16, result XPathResult) (ret XPathResult) {
	bindings.CallXPathEvaluatorEvaluate(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
		result.Ref(),
	)

	return
}

// HasFuncEvaluate1 returns true if the method "XPathEvaluator.evaluate" exists.
func (this XPathEvaluator) HasFuncEvaluate1() bool {
	return js.True == bindings.HasFuncXPathEvaluatorEvaluate1(
		this.ref,
	)
}

// FuncEvaluate1 returns the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) FuncEvaluate1() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16) XPathResult]) {
	bindings.FuncXPathEvaluatorEvaluate1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Evaluate1 calls the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) Evaluate1(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16) (ret XPathResult) {
	bindings.CallXPathEvaluatorEvaluate1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
	)

	return
}

// HasFuncEvaluate2 returns true if the method "XPathEvaluator.evaluate" exists.
func (this XPathEvaluator) HasFuncEvaluate2() bool {
	return js.True == bindings.HasFuncXPathEvaluatorEvaluate2(
		this.ref,
	)
}

// FuncEvaluate2 returns the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) FuncEvaluate2() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String]) XPathResult]) {
	bindings.FuncXPathEvaluatorEvaluate2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Evaluate2 calls the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) Evaluate2(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String]) (ret XPathResult) {
	bindings.CallXPathEvaluatorEvaluate2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
	)

	return
}

// HasFuncEvaluate3 returns true if the method "XPathEvaluator.evaluate" exists.
func (this XPathEvaluator) HasFuncEvaluate3() bool {
	return js.True == bindings.HasFuncXPathEvaluatorEvaluate3(
		this.ref,
	)
}

// FuncEvaluate3 returns the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) FuncEvaluate3() (fn js.Func[func(expression js.String, contextNode Node) XPathResult]) {
	bindings.FuncXPathEvaluatorEvaluate3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Evaluate3 calls the method "XPathEvaluator.evaluate".
func (this XPathEvaluator) Evaluate3(expression js.String, contextNode Node) (ret XPathResult) {
	bindings.CallXPathEvaluatorEvaluate3(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
	)

	return
}

type XRBoundedReferenceSpace struct {
	XRReferenceSpace
}

func (this XRBoundedReferenceSpace) Once() XRBoundedReferenceSpace {
	this.ref.Once()
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
	this.ref.Free()
}

// BoundsGeometry returns the value of property "XRBoundedReferenceSpace.boundsGeometry".
//
// It returns ok=false if there is no such property.
func (this XRBoundedReferenceSpace) BoundsGeometry() (ret js.FrozenArray[DOMPointReadOnly], ok bool) {
	ok = js.True == bindings.GetXRBoundedReferenceSpaceBoundsGeometry(
		this.ref, js.Pointer(&ret),
	)
	return
}
