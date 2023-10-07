// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

type (
	_ unsafe.Pointer
	_ js.Ref
)

//go:wasmimport plat/js/web new_WebGLContextEvent_WebGLContextEvent
//go:noescape
func NewWebGLContextEventByWebGLContextEvent(
	typ js.Ref,
	eventInit unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_WebGLContextEvent_WebGLContextEvent1
//go:noescape
func NewWebGLContextEventByWebGLContextEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_WebGLContextEvent_StatusMessage
//go:noescape
func GetWebGLContextEventStatusMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_WebSocket_WebSocket
//go:noescape
func NewWebSocketByWebSocket(
	url js.Ref,
	protocols js.Ref) js.Ref

//go:wasmimport plat/js/web new_WebSocket_WebSocket1
//go:noescape
func NewWebSocketByWebSocket1(
	url js.Ref) js.Ref

//go:wasmimport plat/js/web get_WebSocket_Url
//go:noescape
func GetWebSocketUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebSocket_ReadyState
//go:noescape
func GetWebSocketReadyState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebSocket_BufferedAmount
//go:noescape
func GetWebSocketBufferedAmount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebSocket_Extensions
//go:noescape
func GetWebSocketExtensions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebSocket_Protocol
//go:noescape
func GetWebSocketProtocol(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebSocket_BinaryType
//go:noescape
func GetWebSocketBinaryType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_WebSocket_BinaryType
//go:noescape
func SetWebSocketBinaryType(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web has_WebSocket_Close
//go:noescape
func HasFuncWebSocketClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WebSocket_Close
//go:noescape
func FuncWebSocketClose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebSocket_Close
//go:noescape
func CallWebSocketClose(
	this js.Ref, retPtr unsafe.Pointer,
	code uint32,
	reason js.Ref)

//go:wasmimport plat/js/web try_WebSocket_Close
//go:noescape
func TryWebSocketClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	code uint32,
	reason js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WebSocket_Close1
//go:noescape
func HasFuncWebSocketClose1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WebSocket_Close1
//go:noescape
func FuncWebSocketClose1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebSocket_Close1
//go:noescape
func CallWebSocketClose1(
	this js.Ref, retPtr unsafe.Pointer,
	code uint32)

//go:wasmimport plat/js/web try_WebSocket_Close1
//go:noescape
func TryWebSocketClose1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	code uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_WebSocket_Close2
//go:noescape
func HasFuncWebSocketClose2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WebSocket_Close2
//go:noescape
func FuncWebSocketClose2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebSocket_Close2
//go:noescape
func CallWebSocketClose2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WebSocket_Close2
//go:noescape
func TryWebSocketClose2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WebSocket_Send
//go:noescape
func HasFuncWebSocketSend(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WebSocket_Send
//go:noescape
func FuncWebSocketSend(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebSocket_Send
//go:noescape
func CallWebSocketSend(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_WebSocket_Send
//go:noescape
func TryWebSocketSend(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_WebTransportHash
//go:noescape
func WebTransportHashJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WebTransportHash
//go:noescape
func WebTransportHashJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_WebTransportCongestionControl
//go:noescape
func ConstOfWebTransportCongestionControl(str js.Ref) uint32

//go:wasmimport plat/js/web store_WebTransportOptions
//go:noescape
func WebTransportOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WebTransportOptions
//go:noescape
func WebTransportOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_WebTransportDatagramStats
//go:noescape
func WebTransportDatagramStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WebTransportDatagramStats
//go:noescape
func WebTransportDatagramStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_WebTransportStats
//go:noescape
func WebTransportStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WebTransportStats
//go:noescape
func WebTransportStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_WebTransportCloseInfo
//go:noescape
func WebTransportCloseInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WebTransportCloseInfo
//go:noescape
func WebTransportCloseInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_WebTransportReceiveStreamStats
//go:noescape
func WebTransportReceiveStreamStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WebTransportReceiveStreamStats
//go:noescape
func WebTransportReceiveStreamStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_WebTransportReceiveStream_WebTransportReceiveStream
//go:noescape
func NewWebTransportReceiveStreamByWebTransportReceiveStream(
	underlyingSource js.Ref,
	strategy unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_WebTransportReceiveStream_WebTransportReceiveStream1
//go:noescape
func NewWebTransportReceiveStreamByWebTransportReceiveStream1(
	underlyingSource js.Ref) js.Ref

//go:wasmimport plat/js/web new_WebTransportReceiveStream_WebTransportReceiveStream2
//go:noescape
func NewWebTransportReceiveStreamByWebTransportReceiveStream2() js.Ref

//go:wasmimport plat/js/web has_WebTransportReceiveStream_GetStats
//go:noescape
func HasFuncWebTransportReceiveStreamGetStats(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WebTransportReceiveStream_GetStats
//go:noescape
func FuncWebTransportReceiveStreamGetStats(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebTransportReceiveStream_GetStats
//go:noescape
func CallWebTransportReceiveStreamGetStats(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WebTransportReceiveStream_GetStats
//go:noescape
func TryWebTransportReceiveStreamGetStats(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_WebTransportSendStreamStats
//go:noescape
func WebTransportSendStreamStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WebTransportSendStreamStats
//go:noescape
func WebTransportSendStreamStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_WebTransportSendStream_WebTransportSendStream
//go:noescape
func NewWebTransportSendStreamByWebTransportSendStream(
	underlyingSink js.Ref,
	strategy unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_WebTransportSendStream_WebTransportSendStream1
//go:noescape
func NewWebTransportSendStreamByWebTransportSendStream1(
	underlyingSink js.Ref) js.Ref

//go:wasmimport plat/js/web new_WebTransportSendStream_WebTransportSendStream2
//go:noescape
func NewWebTransportSendStreamByWebTransportSendStream2() js.Ref

//go:wasmimport plat/js/web get_WebTransportSendStream_SendOrder
//go:noescape
func GetWebTransportSendStreamSendOrder(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_WebTransportSendStream_SendOrder
//go:noescape
func SetWebTransportSendStreamSendOrder(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web has_WebTransportSendStream_GetStats
//go:noescape
func HasFuncWebTransportSendStreamGetStats(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WebTransportSendStream_GetStats
//go:noescape
func FuncWebTransportSendStreamGetStats(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebTransportSendStream_GetStats
//go:noescape
func CallWebTransportSendStreamGetStats(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WebTransportSendStream_GetStats
//go:noescape
func TryWebTransportSendStreamGetStats(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebTransportBidirectionalStream_Readable
//go:noescape
func GetWebTransportBidirectionalStreamReadable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebTransportBidirectionalStream_Writable
//go:noescape
func GetWebTransportBidirectionalStreamWritable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_WebTransportSendStreamOptions
//go:noescape
func WebTransportSendStreamOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WebTransportSendStreamOptions
//go:noescape
func WebTransportSendStreamOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_WebTransportReliabilityMode
//go:noescape
func ConstOfWebTransportReliabilityMode(str js.Ref) uint32

//go:wasmimport plat/js/web get_WebTransportDatagramDuplexStream_Readable
//go:noescape
func GetWebTransportDatagramDuplexStreamReadable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebTransportDatagramDuplexStream_Writable
//go:noescape
func GetWebTransportDatagramDuplexStreamWritable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebTransportDatagramDuplexStream_MaxDatagramSize
//go:noescape
func GetWebTransportDatagramDuplexStreamMaxDatagramSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebTransportDatagramDuplexStream_IncomingMaxAge
//go:noescape
func GetWebTransportDatagramDuplexStreamIncomingMaxAge(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_WebTransportDatagramDuplexStream_IncomingMaxAge
//go:noescape
func SetWebTransportDatagramDuplexStreamIncomingMaxAge(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_WebTransportDatagramDuplexStream_OutgoingMaxAge
//go:noescape
func GetWebTransportDatagramDuplexStreamOutgoingMaxAge(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_WebTransportDatagramDuplexStream_OutgoingMaxAge
//go:noescape
func SetWebTransportDatagramDuplexStreamOutgoingMaxAge(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_WebTransportDatagramDuplexStream_IncomingHighWaterMark
//go:noescape
func GetWebTransportDatagramDuplexStreamIncomingHighWaterMark(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_WebTransportDatagramDuplexStream_IncomingHighWaterMark
//go:noescape
func SetWebTransportDatagramDuplexStreamIncomingHighWaterMark(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_WebTransportDatagramDuplexStream_OutgoingHighWaterMark
//go:noescape
func GetWebTransportDatagramDuplexStreamOutgoingHighWaterMark(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_WebTransportDatagramDuplexStream_OutgoingHighWaterMark
//go:noescape
func SetWebTransportDatagramDuplexStreamOutgoingHighWaterMark(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web new_WebTransport_WebTransport
//go:noescape
func NewWebTransportByWebTransport(
	url js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_WebTransport_WebTransport1
//go:noescape
func NewWebTransportByWebTransport1(
	url js.Ref) js.Ref

//go:wasmimport plat/js/web get_WebTransport_Ready
//go:noescape
func GetWebTransportReady(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebTransport_Reliability
//go:noescape
func GetWebTransportReliability(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebTransport_CongestionControl
//go:noescape
func GetWebTransportCongestionControl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebTransport_Closed
//go:noescape
func GetWebTransportClosed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebTransport_Draining
//go:noescape
func GetWebTransportDraining(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebTransport_Datagrams
//go:noescape
func GetWebTransportDatagrams(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebTransport_IncomingBidirectionalStreams
//go:noescape
func GetWebTransportIncomingBidirectionalStreams(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebTransport_IncomingUnidirectionalStreams
//go:noescape
func GetWebTransportIncomingUnidirectionalStreams(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WebTransport_GetStats
//go:noescape
func HasFuncWebTransportGetStats(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WebTransport_GetStats
//go:noescape
func FuncWebTransportGetStats(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebTransport_GetStats
//go:noescape
func CallWebTransportGetStats(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WebTransport_GetStats
//go:noescape
func TryWebTransportGetStats(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WebTransport_Close
//go:noescape
func HasFuncWebTransportClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WebTransport_Close
//go:noescape
func FuncWebTransportClose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebTransport_Close
//go:noescape
func CallWebTransportClose(
	this js.Ref, retPtr unsafe.Pointer,
	closeInfo unsafe.Pointer)

//go:wasmimport plat/js/web try_WebTransport_Close
//go:noescape
func TryWebTransportClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	closeInfo unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WebTransport_Close1
//go:noescape
func HasFuncWebTransportClose1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WebTransport_Close1
//go:noescape
func FuncWebTransportClose1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebTransport_Close1
//go:noescape
func CallWebTransportClose1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WebTransport_Close1
//go:noescape
func TryWebTransportClose1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WebTransport_CreateBidirectionalStream
//go:noescape
func HasFuncWebTransportCreateBidirectionalStream(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WebTransport_CreateBidirectionalStream
//go:noescape
func FuncWebTransportCreateBidirectionalStream(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebTransport_CreateBidirectionalStream
//go:noescape
func CallWebTransportCreateBidirectionalStream(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_WebTransport_CreateBidirectionalStream
//go:noescape
func TryWebTransportCreateBidirectionalStream(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WebTransport_CreateBidirectionalStream1
//go:noescape
func HasFuncWebTransportCreateBidirectionalStream1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WebTransport_CreateBidirectionalStream1
//go:noescape
func FuncWebTransportCreateBidirectionalStream1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebTransport_CreateBidirectionalStream1
//go:noescape
func CallWebTransportCreateBidirectionalStream1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WebTransport_CreateBidirectionalStream1
//go:noescape
func TryWebTransportCreateBidirectionalStream1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WebTransport_CreateUnidirectionalStream
//go:noescape
func HasFuncWebTransportCreateUnidirectionalStream(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WebTransport_CreateUnidirectionalStream
//go:noescape
func FuncWebTransportCreateUnidirectionalStream(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebTransport_CreateUnidirectionalStream
//go:noescape
func CallWebTransportCreateUnidirectionalStream(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_WebTransport_CreateUnidirectionalStream
//go:noescape
func TryWebTransportCreateUnidirectionalStream(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WebTransport_CreateUnidirectionalStream1
//go:noescape
func HasFuncWebTransportCreateUnidirectionalStream1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WebTransport_CreateUnidirectionalStream1
//go:noescape
func FuncWebTransportCreateUnidirectionalStream1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebTransport_CreateUnidirectionalStream1
//go:noescape
func CallWebTransportCreateUnidirectionalStream1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WebTransport_CreateUnidirectionalStream1
//go:noescape
func TryWebTransportCreateUnidirectionalStream1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_WebTransportErrorSource
//go:noescape
func ConstOfWebTransportErrorSource(str js.Ref) uint32

//go:wasmimport plat/js/web store_WebTransportErrorOptions
//go:noescape
func WebTransportErrorOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WebTransportErrorOptions
//go:noescape
func WebTransportErrorOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_WebTransportError_WebTransportError
//go:noescape
func NewWebTransportErrorByWebTransportError(
	message js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_WebTransportError_WebTransportError1
//go:noescape
func NewWebTransportErrorByWebTransportError1(
	message js.Ref) js.Ref

//go:wasmimport plat/js/web new_WebTransportError_WebTransportError2
//go:noescape
func NewWebTransportErrorByWebTransportError2() js.Ref

//go:wasmimport plat/js/web get_WebTransportError_Source
//go:noescape
func GetWebTransportErrorSource(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebTransportError_StreamErrorCode
//go:noescape
func GetWebTransportErrorStreamErrorCode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_WheelEventInit
//go:noescape
func WheelEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WheelEventInit
//go:noescape
func WheelEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_WheelEvent_WheelEvent
//go:noescape
func NewWheelEventByWheelEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_WheelEvent_WheelEvent1
//go:noescape
func NewWheelEventByWheelEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_WheelEvent_DeltaX
//go:noescape
func GetWheelEventDeltaX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WheelEvent_DeltaY
//go:noescape
func GetWheelEventDeltaY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WheelEvent_DeltaZ
//go:noescape
func GetWheelEventDeltaZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WheelEvent_DeltaMode
//go:noescape
func GetWheelEventDeltaMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_WindowControlsOverlayGeometryChangeEventInit
//go:noescape
func WindowControlsOverlayGeometryChangeEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WindowControlsOverlayGeometryChangeEventInit
//go:noescape
func WindowControlsOverlayGeometryChangeEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_WindowControlsOverlayGeometryChangeEvent_WindowControlsOverlayGeometryChangeEvent
//go:noescape
func NewWindowControlsOverlayGeometryChangeEventByWindowControlsOverlayGeometryChangeEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_WindowControlsOverlayGeometryChangeEvent_TitlebarAreaRect
//go:noescape
func GetWindowControlsOverlayGeometryChangeEventTitlebarAreaRect(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WindowControlsOverlayGeometryChangeEvent_Visible
//go:noescape
func GetWindowControlsOverlayGeometryChangeEventVisible(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerLocation_Href
//go:noescape
func GetWorkerLocationHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerLocation_Origin
//go:noescape
func GetWorkerLocationOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerLocation_Protocol
//go:noescape
func GetWorkerLocationProtocol(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerLocation_Host
//go:noescape
func GetWorkerLocationHost(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerLocation_Hostname
//go:noescape
func GetWorkerLocationHostname(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerLocation_Port
//go:noescape
func GetWorkerLocationPort(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerLocation_Pathname
//go:noescape
func GetWorkerLocationPathname(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerLocation_Search
//go:noescape
func GetWorkerLocationSearch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerLocation_Hash
//go:noescape
func GetWorkerLocationHash(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_Hid
//go:noescape
func GetWorkerNavigatorHid(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_Serial
//go:noescape
func GetWorkerNavigatorSerial(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_Permissions
//go:noescape
func GetWorkerNavigatorPermissions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_Usb
//go:noescape
func GetWorkerNavigatorUsb(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_MediaCapabilities
//go:noescape
func GetWorkerNavigatorMediaCapabilities(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_ServiceWorker
//go:noescape
func GetWorkerNavigatorServiceWorker(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_UserAgentData
//go:noescape
func GetWorkerNavigatorUserAgentData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_DeviceMemory
//go:noescape
func GetWorkerNavigatorDeviceMemory(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_Storage
//go:noescape
func GetWorkerNavigatorStorage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_StorageBuckets
//go:noescape
func GetWorkerNavigatorStorageBuckets(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_Locks
//go:noescape
func GetWorkerNavigatorLocks(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_AppCodeName
//go:noescape
func GetWorkerNavigatorAppCodeName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_AppName
//go:noescape
func GetWorkerNavigatorAppName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_AppVersion
//go:noescape
func GetWorkerNavigatorAppVersion(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_Platform
//go:noescape
func GetWorkerNavigatorPlatform(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_Product
//go:noescape
func GetWorkerNavigatorProduct(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_ProductSub
//go:noescape
func GetWorkerNavigatorProductSub(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_UserAgent
//go:noescape
func GetWorkerNavigatorUserAgent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_Vendor
//go:noescape
func GetWorkerNavigatorVendor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_VendorSub
//go:noescape
func GetWorkerNavigatorVendorSub(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_Oscpu
//go:noescape
func GetWorkerNavigatorOscpu(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_Language
//go:noescape
func GetWorkerNavigatorLanguage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_Languages
//go:noescape
func GetWorkerNavigatorLanguages(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_OnLine
//go:noescape
func GetWorkerNavigatorOnLine(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_HardwareConcurrency
//go:noescape
func GetWorkerNavigatorHardwareConcurrency(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_Connection
//go:noescape
func GetWorkerNavigatorConnection(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_Ml
//go:noescape
func GetWorkerNavigatorMl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerNavigator_Gpu
//go:noescape
func GetWorkerNavigatorGpu(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerNavigator_TaintEnabled
//go:noescape
func HasFuncWorkerNavigatorTaintEnabled(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerNavigator_TaintEnabled
//go:noescape
func FuncWorkerNavigatorTaintEnabled(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerNavigator_TaintEnabled
//go:noescape
func CallWorkerNavigatorTaintEnabled(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WorkerNavigator_TaintEnabled
//go:noescape
func TryWorkerNavigatorTaintEnabled(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerNavigator_SetAppBadge
//go:noescape
func HasFuncWorkerNavigatorSetAppBadge(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerNavigator_SetAppBadge
//go:noescape
func FuncWorkerNavigatorSetAppBadge(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerNavigator_SetAppBadge
//go:noescape
func CallWorkerNavigatorSetAppBadge(
	this js.Ref, retPtr unsafe.Pointer,
	contents float64)

//go:wasmimport plat/js/web try_WorkerNavigator_SetAppBadge
//go:noescape
func TryWorkerNavigatorSetAppBadge(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	contents float64) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerNavigator_SetAppBadge1
//go:noescape
func HasFuncWorkerNavigatorSetAppBadge1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerNavigator_SetAppBadge1
//go:noescape
func FuncWorkerNavigatorSetAppBadge1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerNavigator_SetAppBadge1
//go:noescape
func CallWorkerNavigatorSetAppBadge1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WorkerNavigator_SetAppBadge1
//go:noescape
func TryWorkerNavigatorSetAppBadge1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerNavigator_ClearAppBadge
//go:noescape
func HasFuncWorkerNavigatorClearAppBadge(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerNavigator_ClearAppBadge
//go:noescape
func FuncWorkerNavigatorClearAppBadge(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerNavigator_ClearAppBadge
//go:noescape
func CallWorkerNavigatorClearAppBadge(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WorkerNavigator_ClearAppBadge
//go:noescape
func TryWorkerNavigatorClearAppBadge(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerGlobalScope_Self
//go:noescape
func GetWorkerGlobalScopeSelf(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerGlobalScope_Location
//go:noescape
func GetWorkerGlobalScopeLocation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerGlobalScope_Navigator
//go:noescape
func GetWorkerGlobalScopeNavigator(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerGlobalScope_Fonts
//go:noescape
func GetWorkerGlobalScopeFonts(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerGlobalScope_Origin
//go:noescape
func GetWorkerGlobalScopeOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerGlobalScope_IsSecureContext
//go:noescape
func GetWorkerGlobalScopeIsSecureContext(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerGlobalScope_CrossOriginIsolated
//go:noescape
func GetWorkerGlobalScopeCrossOriginIsolated(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerGlobalScope_Scheduler
//go:noescape
func GetWorkerGlobalScopeScheduler(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerGlobalScope_TrustedTypes
//go:noescape
func GetWorkerGlobalScopeTrustedTypes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerGlobalScope_Caches
//go:noescape
func GetWorkerGlobalScopeCaches(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerGlobalScope_Crypto
//go:noescape
func GetWorkerGlobalScopeCrypto(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerGlobalScope_IndexedDB
//go:noescape
func GetWorkerGlobalScopeIndexedDB(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkerGlobalScope_Performance
//go:noescape
func GetWorkerGlobalScopePerformance(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_ImportScripts
//go:noescape
func HasFuncWorkerGlobalScopeImportScripts(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_ImportScripts
//go:noescape
func FuncWorkerGlobalScopeImportScripts(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_ImportScripts
//go:noescape
func CallWorkerGlobalScopeImportScripts(
	this js.Ref, retPtr unsafe.Pointer,
	urlsPtr unsafe.Pointer,
	urlsCount uint32)

//go:wasmimport plat/js/web try_WorkerGlobalScope_ImportScripts
//go:noescape
func TryWorkerGlobalScopeImportScripts(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	urlsPtr unsafe.Pointer,
	urlsCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_ReportError
//go:noescape
func HasFuncWorkerGlobalScopeReportError(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_ReportError
//go:noescape
func FuncWorkerGlobalScopeReportError(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_ReportError
//go:noescape
func CallWorkerGlobalScopeReportError(
	this js.Ref, retPtr unsafe.Pointer,
	e js.Ref)

//go:wasmimport plat/js/web try_WorkerGlobalScope_ReportError
//go:noescape
func TryWorkerGlobalScopeReportError(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	e js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_Btoa
//go:noescape
func HasFuncWorkerGlobalScopeBtoa(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_Btoa
//go:noescape
func FuncWorkerGlobalScopeBtoa(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_Btoa
//go:noescape
func CallWorkerGlobalScopeBtoa(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_WorkerGlobalScope_Btoa
//go:noescape
func TryWorkerGlobalScopeBtoa(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_Atob
//go:noescape
func HasFuncWorkerGlobalScopeAtob(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_Atob
//go:noescape
func FuncWorkerGlobalScopeAtob(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_Atob
//go:noescape
func CallWorkerGlobalScopeAtob(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_WorkerGlobalScope_Atob
//go:noescape
func TryWorkerGlobalScopeAtob(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_SetTimeout
//go:noescape
func HasFuncWorkerGlobalScopeSetTimeout(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_SetTimeout
//go:noescape
func FuncWorkerGlobalScopeSetTimeout(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_SetTimeout
//go:noescape
func CallWorkerGlobalScopeSetTimeout(
	this js.Ref, retPtr unsafe.Pointer,
	handler js.Ref,
	timeout int32,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32)

//go:wasmimport plat/js/web try_WorkerGlobalScope_SetTimeout
//go:noescape
func TryWorkerGlobalScopeSetTimeout(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handler js.Ref,
	timeout int32,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_SetTimeout1
//go:noescape
func HasFuncWorkerGlobalScopeSetTimeout1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_SetTimeout1
//go:noescape
func FuncWorkerGlobalScopeSetTimeout1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_SetTimeout1
//go:noescape
func CallWorkerGlobalScopeSetTimeout1(
	this js.Ref, retPtr unsafe.Pointer,
	handler js.Ref)

//go:wasmimport plat/js/web try_WorkerGlobalScope_SetTimeout1
//go:noescape
func TryWorkerGlobalScopeSetTimeout1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handler js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_ClearTimeout
//go:noescape
func HasFuncWorkerGlobalScopeClearTimeout(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_ClearTimeout
//go:noescape
func FuncWorkerGlobalScopeClearTimeout(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_ClearTimeout
//go:noescape
func CallWorkerGlobalScopeClearTimeout(
	this js.Ref, retPtr unsafe.Pointer,
	id int32)

//go:wasmimport plat/js/web try_WorkerGlobalScope_ClearTimeout
//go:noescape
func TryWorkerGlobalScopeClearTimeout(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id int32) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_ClearTimeout1
//go:noescape
func HasFuncWorkerGlobalScopeClearTimeout1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_ClearTimeout1
//go:noescape
func FuncWorkerGlobalScopeClearTimeout1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_ClearTimeout1
//go:noescape
func CallWorkerGlobalScopeClearTimeout1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WorkerGlobalScope_ClearTimeout1
//go:noescape
func TryWorkerGlobalScopeClearTimeout1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_SetInterval
//go:noescape
func HasFuncWorkerGlobalScopeSetInterval(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_SetInterval
//go:noescape
func FuncWorkerGlobalScopeSetInterval(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_SetInterval
//go:noescape
func CallWorkerGlobalScopeSetInterval(
	this js.Ref, retPtr unsafe.Pointer,
	handler js.Ref,
	timeout int32,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32)

//go:wasmimport plat/js/web try_WorkerGlobalScope_SetInterval
//go:noescape
func TryWorkerGlobalScopeSetInterval(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handler js.Ref,
	timeout int32,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_SetInterval1
//go:noescape
func HasFuncWorkerGlobalScopeSetInterval1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_SetInterval1
//go:noescape
func FuncWorkerGlobalScopeSetInterval1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_SetInterval1
//go:noescape
func CallWorkerGlobalScopeSetInterval1(
	this js.Ref, retPtr unsafe.Pointer,
	handler js.Ref)

//go:wasmimport plat/js/web try_WorkerGlobalScope_SetInterval1
//go:noescape
func TryWorkerGlobalScopeSetInterval1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handler js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_ClearInterval
//go:noescape
func HasFuncWorkerGlobalScopeClearInterval(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_ClearInterval
//go:noescape
func FuncWorkerGlobalScopeClearInterval(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_ClearInterval
//go:noescape
func CallWorkerGlobalScopeClearInterval(
	this js.Ref, retPtr unsafe.Pointer,
	id int32)

//go:wasmimport plat/js/web try_WorkerGlobalScope_ClearInterval
//go:noescape
func TryWorkerGlobalScopeClearInterval(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id int32) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_ClearInterval1
//go:noescape
func HasFuncWorkerGlobalScopeClearInterval1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_ClearInterval1
//go:noescape
func FuncWorkerGlobalScopeClearInterval1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_ClearInterval1
//go:noescape
func CallWorkerGlobalScopeClearInterval1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WorkerGlobalScope_ClearInterval1
//go:noescape
func TryWorkerGlobalScopeClearInterval1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_QueueMicrotask
//go:noescape
func HasFuncWorkerGlobalScopeQueueMicrotask(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_QueueMicrotask
//go:noescape
func FuncWorkerGlobalScopeQueueMicrotask(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_QueueMicrotask
//go:noescape
func CallWorkerGlobalScopeQueueMicrotask(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/web try_WorkerGlobalScope_QueueMicrotask
//go:noescape
func TryWorkerGlobalScopeQueueMicrotask(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_CreateImageBitmap
//go:noescape
func HasFuncWorkerGlobalScopeCreateImageBitmap(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_CreateImageBitmap
//go:noescape
func FuncWorkerGlobalScopeCreateImageBitmap(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_CreateImageBitmap
//go:noescape
func CallWorkerGlobalScopeCreateImageBitmap(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_WorkerGlobalScope_CreateImageBitmap
//go:noescape
func TryWorkerGlobalScopeCreateImageBitmap(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_CreateImageBitmap1
//go:noescape
func HasFuncWorkerGlobalScopeCreateImageBitmap1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_CreateImageBitmap1
//go:noescape
func FuncWorkerGlobalScopeCreateImageBitmap1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_CreateImageBitmap1
//go:noescape
func CallWorkerGlobalScopeCreateImageBitmap1(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref)

//go:wasmimport plat/js/web try_WorkerGlobalScope_CreateImageBitmap1
//go:noescape
func TryWorkerGlobalScopeCreateImageBitmap1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_CreateImageBitmap2
//go:noescape
func HasFuncWorkerGlobalScopeCreateImageBitmap2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_CreateImageBitmap2
//go:noescape
func FuncWorkerGlobalScopeCreateImageBitmap2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_CreateImageBitmap2
//go:noescape
func CallWorkerGlobalScopeCreateImageBitmap2(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref,
	sx int32,
	sy int32,
	sw int32,
	sh int32,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_WorkerGlobalScope_CreateImageBitmap2
//go:noescape
func TryWorkerGlobalScopeCreateImageBitmap2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref,
	sx int32,
	sy int32,
	sw int32,
	sh int32,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_CreateImageBitmap3
//go:noescape
func HasFuncWorkerGlobalScopeCreateImageBitmap3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_CreateImageBitmap3
//go:noescape
func FuncWorkerGlobalScopeCreateImageBitmap3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_CreateImageBitmap3
//go:noescape
func CallWorkerGlobalScopeCreateImageBitmap3(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref,
	sx int32,
	sy int32,
	sw int32,
	sh int32)

//go:wasmimport plat/js/web try_WorkerGlobalScope_CreateImageBitmap3
//go:noescape
func TryWorkerGlobalScopeCreateImageBitmap3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref,
	sx int32,
	sy int32,
	sw int32,
	sh int32) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_StructuredClone
//go:noescape
func HasFuncWorkerGlobalScopeStructuredClone(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_StructuredClone
//go:noescape
func FuncWorkerGlobalScopeStructuredClone(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_StructuredClone
//go:noescape
func CallWorkerGlobalScopeStructuredClone(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_WorkerGlobalScope_StructuredClone
//go:noescape
func TryWorkerGlobalScopeStructuredClone(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_StructuredClone1
//go:noescape
func HasFuncWorkerGlobalScopeStructuredClone1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_StructuredClone1
//go:noescape
func FuncWorkerGlobalScopeStructuredClone1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_StructuredClone1
//go:noescape
func CallWorkerGlobalScopeStructuredClone1(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_WorkerGlobalScope_StructuredClone1
//go:noescape
func TryWorkerGlobalScopeStructuredClone1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_Fetch
//go:noescape
func HasFuncWorkerGlobalScopeFetch(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_Fetch
//go:noescape
func FuncWorkerGlobalScopeFetch(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_Fetch
//go:noescape
func CallWorkerGlobalScopeFetch(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	init unsafe.Pointer)

//go:wasmimport plat/js/web try_WorkerGlobalScope_Fetch
//go:noescape
func TryWorkerGlobalScopeFetch(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	init unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkerGlobalScope_Fetch1
//go:noescape
func HasFuncWorkerGlobalScopeFetch1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_Fetch1
//go:noescape
func FuncWorkerGlobalScopeFetch1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkerGlobalScope_Fetch1
//go:noescape
func CallWorkerGlobalScopeFetch1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_WorkerGlobalScope_Fetch1
//go:noescape
func TryWorkerGlobalScopeFetch1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web new_WorkletAnimation_WorkletAnimation
//go:noescape
func NewWorkletAnimationByWorkletAnimation(
	animatorName js.Ref,
	effects js.Ref,
	timeline js.Ref,
	options js.Ref) js.Ref

//go:wasmimport plat/js/web new_WorkletAnimation_WorkletAnimation1
//go:noescape
func NewWorkletAnimationByWorkletAnimation1(
	animatorName js.Ref,
	effects js.Ref,
	timeline js.Ref) js.Ref

//go:wasmimport plat/js/web new_WorkletAnimation_WorkletAnimation2
//go:noescape
func NewWorkletAnimationByWorkletAnimation2(
	animatorName js.Ref,
	effects js.Ref) js.Ref

//go:wasmimport plat/js/web new_WorkletAnimation_WorkletAnimation3
//go:noescape
func NewWorkletAnimationByWorkletAnimation3(
	animatorName js.Ref) js.Ref

//go:wasmimport plat/js/web get_WorkletAnimation_AnimatorName
//go:noescape
func GetWorkletAnimationAnimatorName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WorkletAnimationEffect_LocalTime
//go:noescape
func GetWorkletAnimationEffectLocalTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_WorkletAnimationEffect_LocalTime
//go:noescape
func SetWorkletAnimationEffectLocalTime(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web has_WorkletAnimationEffect_GetTiming
//go:noescape
func HasFuncWorkletAnimationEffectGetTiming(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkletAnimationEffect_GetTiming
//go:noescape
func FuncWorkletAnimationEffectGetTiming(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkletAnimationEffect_GetTiming
//go:noescape
func CallWorkletAnimationEffectGetTiming(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WorkletAnimationEffect_GetTiming
//go:noescape
func TryWorkletAnimationEffectGetTiming(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkletAnimationEffect_GetComputedTiming
//go:noescape
func HasFuncWorkletAnimationEffectGetComputedTiming(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkletAnimationEffect_GetComputedTiming
//go:noescape
func FuncWorkletAnimationEffectGetComputedTiming(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkletAnimationEffect_GetComputedTiming
//go:noescape
func CallWorkletAnimationEffectGetComputedTiming(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WorkletAnimationEffect_GetComputedTiming
//go:noescape
func TryWorkletAnimationEffectGetComputedTiming(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WorkletGroupEffect_GetChildren
//go:noescape
func HasFuncWorkletGroupEffectGetChildren(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WorkletGroupEffect_GetChildren
//go:noescape
func FuncWorkletGroupEffectGetChildren(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WorkletGroupEffect_GetChildren
//go:noescape
func CallWorkletGroupEffectGetChildren(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WorkletGroupEffect_GetChildren
//go:noescape
func TryWorkletGroupEffectGetChildren(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_XMLHttpRequestResponseType
//go:noescape
func ConstOfXMLHttpRequestResponseType(str js.Ref) uint32

//go:wasmimport plat/js/web get_XMLHttpRequest_ReadyState
//go:noescape
func GetXMLHttpRequestReadyState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XMLHttpRequest_Timeout
//go:noescape
func GetXMLHttpRequestTimeout(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XMLHttpRequest_Timeout
//go:noescape
func SetXMLHttpRequestTimeout(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_XMLHttpRequest_WithCredentials
//go:noescape
func GetXMLHttpRequestWithCredentials(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XMLHttpRequest_WithCredentials
//go:noescape
func SetXMLHttpRequestWithCredentials(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XMLHttpRequest_Upload
//go:noescape
func GetXMLHttpRequestUpload(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XMLHttpRequest_ResponseURL
//go:noescape
func GetXMLHttpRequestResponseURL(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XMLHttpRequest_Status
//go:noescape
func GetXMLHttpRequestStatus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XMLHttpRequest_StatusText
//go:noescape
func GetXMLHttpRequestStatusText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XMLHttpRequest_ResponseType
//go:noescape
func GetXMLHttpRequestResponseType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XMLHttpRequest_ResponseType
//go:noescape
func SetXMLHttpRequestResponseType(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_XMLHttpRequest_Response
//go:noescape
func GetXMLHttpRequestResponse(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XMLHttpRequest_ResponseText
//go:noescape
func GetXMLHttpRequestResponseText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XMLHttpRequest_ResponseXML
//go:noescape
func GetXMLHttpRequestResponseXML(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XMLHttpRequest_Open
//go:noescape
func HasFuncXMLHttpRequestOpen(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_Open
//go:noescape
func FuncXMLHttpRequestOpen(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XMLHttpRequest_Open
//go:noescape
func CallXMLHttpRequestOpen(
	this js.Ref, retPtr unsafe.Pointer,
	method js.Ref,
	url js.Ref)

//go:wasmimport plat/js/web try_XMLHttpRequest_Open
//go:noescape
func TryXMLHttpRequestOpen(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	method js.Ref,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XMLHttpRequest_Open1
//go:noescape
func HasFuncXMLHttpRequestOpen1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_Open1
//go:noescape
func FuncXMLHttpRequestOpen1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XMLHttpRequest_Open1
//go:noescape
func CallXMLHttpRequestOpen1(
	this js.Ref, retPtr unsafe.Pointer,
	method js.Ref,
	url js.Ref,
	async js.Ref,
	username js.Ref,
	password js.Ref)

//go:wasmimport plat/js/web try_XMLHttpRequest_Open1
//go:noescape
func TryXMLHttpRequestOpen1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	method js.Ref,
	url js.Ref,
	async js.Ref,
	username js.Ref,
	password js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XMLHttpRequest_Open2
//go:noescape
func HasFuncXMLHttpRequestOpen2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_Open2
//go:noescape
func FuncXMLHttpRequestOpen2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XMLHttpRequest_Open2
//go:noescape
func CallXMLHttpRequestOpen2(
	this js.Ref, retPtr unsafe.Pointer,
	method js.Ref,
	url js.Ref,
	async js.Ref,
	username js.Ref)

//go:wasmimport plat/js/web try_XMLHttpRequest_Open2
//go:noescape
func TryXMLHttpRequestOpen2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	method js.Ref,
	url js.Ref,
	async js.Ref,
	username js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XMLHttpRequest_Open3
//go:noescape
func HasFuncXMLHttpRequestOpen3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_Open3
//go:noescape
func FuncXMLHttpRequestOpen3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XMLHttpRequest_Open3
//go:noescape
func CallXMLHttpRequestOpen3(
	this js.Ref, retPtr unsafe.Pointer,
	method js.Ref,
	url js.Ref,
	async js.Ref)

//go:wasmimport plat/js/web try_XMLHttpRequest_Open3
//go:noescape
func TryXMLHttpRequestOpen3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	method js.Ref,
	url js.Ref,
	async js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XMLHttpRequest_SetRequestHeader
//go:noescape
func HasFuncXMLHttpRequestSetRequestHeader(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_SetRequestHeader
//go:noescape
func FuncXMLHttpRequestSetRequestHeader(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XMLHttpRequest_SetRequestHeader
//go:noescape
func CallXMLHttpRequestSetRequestHeader(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_XMLHttpRequest_SetRequestHeader
//go:noescape
func TryXMLHttpRequestSetRequestHeader(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XMLHttpRequest_Send
//go:noescape
func HasFuncXMLHttpRequestSend(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_Send
//go:noescape
func FuncXMLHttpRequestSend(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XMLHttpRequest_Send
//go:noescape
func CallXMLHttpRequestSend(
	this js.Ref, retPtr unsafe.Pointer,
	body js.Ref)

//go:wasmimport plat/js/web try_XMLHttpRequest_Send
//go:noescape
func TryXMLHttpRequestSend(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	body js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XMLHttpRequest_Send1
//go:noescape
func HasFuncXMLHttpRequestSend1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_Send1
//go:noescape
func FuncXMLHttpRequestSend1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XMLHttpRequest_Send1
//go:noescape
func CallXMLHttpRequestSend1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XMLHttpRequest_Send1
//go:noescape
func TryXMLHttpRequestSend1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XMLHttpRequest_Abort
//go:noescape
func HasFuncXMLHttpRequestAbort(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_Abort
//go:noescape
func FuncXMLHttpRequestAbort(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XMLHttpRequest_Abort
//go:noescape
func CallXMLHttpRequestAbort(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XMLHttpRequest_Abort
//go:noescape
func TryXMLHttpRequestAbort(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XMLHttpRequest_GetResponseHeader
//go:noescape
func HasFuncXMLHttpRequestGetResponseHeader(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_GetResponseHeader
//go:noescape
func FuncXMLHttpRequestGetResponseHeader(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XMLHttpRequest_GetResponseHeader
//go:noescape
func CallXMLHttpRequestGetResponseHeader(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_XMLHttpRequest_GetResponseHeader
//go:noescape
func TryXMLHttpRequestGetResponseHeader(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XMLHttpRequest_GetAllResponseHeaders
//go:noescape
func HasFuncXMLHttpRequestGetAllResponseHeaders(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_GetAllResponseHeaders
//go:noescape
func FuncXMLHttpRequestGetAllResponseHeaders(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XMLHttpRequest_GetAllResponseHeaders
//go:noescape
func CallXMLHttpRequestGetAllResponseHeaders(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XMLHttpRequest_GetAllResponseHeaders
//go:noescape
func TryXMLHttpRequestGetAllResponseHeaders(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XMLHttpRequest_OverrideMimeType
//go:noescape
func HasFuncXMLHttpRequestOverrideMimeType(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_OverrideMimeType
//go:noescape
func FuncXMLHttpRequestOverrideMimeType(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XMLHttpRequest_OverrideMimeType
//go:noescape
func CallXMLHttpRequestOverrideMimeType(
	this js.Ref, retPtr unsafe.Pointer,
	mime js.Ref)

//go:wasmimport plat/js/web try_XMLHttpRequest_OverrideMimeType
//go:noescape
func TryXMLHttpRequestOverrideMimeType(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mime js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XMLHttpRequest_SetAttributionReporting
//go:noescape
func HasFuncXMLHttpRequestSetAttributionReporting(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_SetAttributionReporting
//go:noescape
func FuncXMLHttpRequestSetAttributionReporting(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XMLHttpRequest_SetAttributionReporting
//go:noescape
func CallXMLHttpRequestSetAttributionReporting(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_XMLHttpRequest_SetAttributionReporting
//go:noescape
func TryXMLHttpRequestSetAttributionReporting(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XMLHttpRequest_SetPrivateToken
//go:noescape
func HasFuncXMLHttpRequestSetPrivateToken(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_SetPrivateToken
//go:noescape
func FuncXMLHttpRequestSetPrivateToken(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XMLHttpRequest_SetPrivateToken
//go:noescape
func CallXMLHttpRequestSetPrivateToken(
	this js.Ref, retPtr unsafe.Pointer,
	privateToken unsafe.Pointer)

//go:wasmimport plat/js/web try_XMLHttpRequest_SetPrivateToken
//go:noescape
func TryXMLHttpRequestSetPrivateToken(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	privateToken unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XMLSerializer_SerializeToString
//go:noescape
func HasFuncXMLSerializerSerializeToString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XMLSerializer_SerializeToString
//go:noescape
func FuncXMLSerializerSerializeToString(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XMLSerializer_SerializeToString
//go:noescape
func CallXMLSerializerSerializeToString(
	this js.Ref, retPtr unsafe.Pointer,
	root js.Ref)

//go:wasmimport plat/js/web try_XMLSerializer_SerializeToString
//go:noescape
func TryXMLSerializerSerializeToString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	root js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XPathEvaluator_CreateExpression
//go:noescape
func HasFuncXPathEvaluatorCreateExpression(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathEvaluator_CreateExpression
//go:noescape
func FuncXPathEvaluatorCreateExpression(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XPathEvaluator_CreateExpression
//go:noescape
func CallXPathEvaluatorCreateExpression(
	this js.Ref, retPtr unsafe.Pointer,
	expression js.Ref,
	resolver js.Ref)

//go:wasmimport plat/js/web try_XPathEvaluator_CreateExpression
//go:noescape
func TryXPathEvaluatorCreateExpression(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expression js.Ref,
	resolver js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XPathEvaluator_CreateExpression1
//go:noescape
func HasFuncXPathEvaluatorCreateExpression1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathEvaluator_CreateExpression1
//go:noescape
func FuncXPathEvaluatorCreateExpression1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XPathEvaluator_CreateExpression1
//go:noescape
func CallXPathEvaluatorCreateExpression1(
	this js.Ref, retPtr unsafe.Pointer,
	expression js.Ref)

//go:wasmimport plat/js/web try_XPathEvaluator_CreateExpression1
//go:noescape
func TryXPathEvaluatorCreateExpression1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expression js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XPathEvaluator_CreateNSResolver
//go:noescape
func HasFuncXPathEvaluatorCreateNSResolver(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathEvaluator_CreateNSResolver
//go:noescape
func FuncXPathEvaluatorCreateNSResolver(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XPathEvaluator_CreateNSResolver
//go:noescape
func CallXPathEvaluatorCreateNSResolver(
	this js.Ref, retPtr unsafe.Pointer,
	nodeResolver js.Ref)

//go:wasmimport plat/js/web try_XPathEvaluator_CreateNSResolver
//go:noescape
func TryXPathEvaluatorCreateNSResolver(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	nodeResolver js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XPathEvaluator_Evaluate
//go:noescape
func HasFuncXPathEvaluatorEvaluate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathEvaluator_Evaluate
//go:noescape
func FuncXPathEvaluatorEvaluate(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XPathEvaluator_Evaluate
//go:noescape
func CallXPathEvaluatorEvaluate(
	this js.Ref, retPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref,
	typ uint32,
	result js.Ref)

//go:wasmimport plat/js/web try_XPathEvaluator_Evaluate
//go:noescape
func TryXPathEvaluatorEvaluate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref,
	typ uint32,
	result js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XPathEvaluator_Evaluate1
//go:noescape
func HasFuncXPathEvaluatorEvaluate1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathEvaluator_Evaluate1
//go:noescape
func FuncXPathEvaluatorEvaluate1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XPathEvaluator_Evaluate1
//go:noescape
func CallXPathEvaluatorEvaluate1(
	this js.Ref, retPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref,
	typ uint32)

//go:wasmimport plat/js/web try_XPathEvaluator_Evaluate1
//go:noescape
func TryXPathEvaluatorEvaluate1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref,
	typ uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_XPathEvaluator_Evaluate2
//go:noescape
func HasFuncXPathEvaluatorEvaluate2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathEvaluator_Evaluate2
//go:noescape
func FuncXPathEvaluatorEvaluate2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XPathEvaluator_Evaluate2
//go:noescape
func CallXPathEvaluatorEvaluate2(
	this js.Ref, retPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref)

//go:wasmimport plat/js/web try_XPathEvaluator_Evaluate2
//go:noescape
func TryXPathEvaluatorEvaluate2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XPathEvaluator_Evaluate3
//go:noescape
func HasFuncXPathEvaluatorEvaluate3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathEvaluator_Evaluate3
//go:noescape
func FuncXPathEvaluatorEvaluate3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XPathEvaluator_Evaluate3
//go:noescape
func CallXPathEvaluatorEvaluate3(
	this js.Ref, retPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref)

//go:wasmimport plat/js/web try_XPathEvaluator_Evaluate3
//go:noescape
func TryXPathEvaluatorEvaluate3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_XRBoundedReferenceSpace_BoundsGeometry
//go:noescape
func GetXRBoundedReferenceSpaceBoundsGeometry(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)
