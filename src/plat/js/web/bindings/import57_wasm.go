// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package bindings

import (
	"unsafe"

	"github.com/primecitizens/std/ffi/js"
)

func _() {
	var (
		_ js.Void
		_ unsafe.Pointer
	)
}

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WebSocket_ReadyState
//go:noescape

func GetWebSocketReadyState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_WebSocket_BufferedAmount
//go:noescape

func GetWebSocketBufferedAmount(
	this js.Ref,
	ptr unsafe.Pointer,
) uint64

//go:wasmimport plat/js/web get_WebSocket_Extensions
//go:noescape

func GetWebSocketExtensions(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WebSocket_Protocol
//go:noescape

func GetWebSocketProtocol(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WebSocket_BinaryType
//go:noescape

func GetWebSocketBinaryType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_WebSocket_BinaryType
//go:noescape

func SetWebSocketBinaryType(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web call_WebSocket_Close
//go:noescape

func CallWebSocketClose(
	this js.Ref,
	ptr unsafe.Pointer,

	code uint32,
	reason js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebSocket_Close
//go:noescape

func WebSocketCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebSocket_Close1
//go:noescape

func CallWebSocketClose1(
	this js.Ref,
	ptr unsafe.Pointer,

	code uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebSocket_Close1
//go:noescape

func WebSocketClose1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebSocket_Close2
//go:noescape

func CallWebSocketClose2(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebSocket_Close2
//go:noescape

func WebSocketClose2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebSocket_Send
//go:noescape

func CallWebSocketSend(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebSocket_Send
//go:noescape

func WebSocketSendFunc(this js.Ref) js.Ref

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

//go:wasmimport plat/js/web call_WebTransportReceiveStream_GetStats
//go:noescape

func CallWebTransportReceiveStreamGetStats(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebTransportReceiveStream_GetStats
//go:noescape

func WebTransportReceiveStreamGetStatsFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) int64

//go:wasmimport plat/js/web set_WebTransportSendStream_SendOrder
//go:noescape

func SetWebTransportSendStreamSendOrder(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web call_WebTransportSendStream_GetStats
//go:noescape

func CallWebTransportSendStreamGetStats(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebTransportSendStream_GetStats
//go:noescape

func WebTransportSendStreamGetStatsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_WebTransportBidirectionalStream_Readable
//go:noescape

func GetWebTransportBidirectionalStreamReadable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WebTransportBidirectionalStream_Writable
//go:noescape

func GetWebTransportBidirectionalStreamWritable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WebTransportDatagramDuplexStream_Writable
//go:noescape

func GetWebTransportDatagramDuplexStreamWritable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WebTransportDatagramDuplexStream_MaxDatagramSize
//go:noescape

func GetWebTransportDatagramDuplexStreamMaxDatagramSize(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_WebTransportDatagramDuplexStream_IncomingMaxAge
//go:noescape

func GetWebTransportDatagramDuplexStreamIncomingMaxAge(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_WebTransportDatagramDuplexStream_IncomingMaxAge
//go:noescape

func SetWebTransportDatagramDuplexStreamIncomingMaxAge(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_WebTransportDatagramDuplexStream_OutgoingMaxAge
//go:noescape

func GetWebTransportDatagramDuplexStreamOutgoingMaxAge(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_WebTransportDatagramDuplexStream_OutgoingMaxAge
//go:noescape

func SetWebTransportDatagramDuplexStreamOutgoingMaxAge(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_WebTransportDatagramDuplexStream_IncomingHighWaterMark
//go:noescape

func GetWebTransportDatagramDuplexStreamIncomingHighWaterMark(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_WebTransportDatagramDuplexStream_IncomingHighWaterMark
//go:noescape

func SetWebTransportDatagramDuplexStreamIncomingHighWaterMark(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_WebTransportDatagramDuplexStream_OutgoingHighWaterMark
//go:noescape

func GetWebTransportDatagramDuplexStreamOutgoingHighWaterMark(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WebTransport_Reliability
//go:noescape

func GetWebTransportReliability(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_WebTransport_CongestionControl
//go:noescape

func GetWebTransportCongestionControl(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_WebTransport_Closed
//go:noescape

func GetWebTransportClosed(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WebTransport_Draining
//go:noescape

func GetWebTransportDraining(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WebTransport_Datagrams
//go:noescape

func GetWebTransportDatagrams(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WebTransport_IncomingBidirectionalStreams
//go:noescape

func GetWebTransportIncomingBidirectionalStreams(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WebTransport_IncomingUnidirectionalStreams
//go:noescape

func GetWebTransportIncomingUnidirectionalStreams(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_WebTransport_GetStats
//go:noescape

func CallWebTransportGetStats(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebTransport_GetStats
//go:noescape

func WebTransportGetStatsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebTransport_Close
//go:noescape

func CallWebTransportClose(
	this js.Ref,
	ptr unsafe.Pointer,

	closeInfo unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_WebTransport_Close
//go:noescape

func WebTransportCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebTransport_Close1
//go:noescape

func CallWebTransportClose1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebTransport_Close1
//go:noescape

func WebTransportClose1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebTransport_CreateBidirectionalStream
//go:noescape

func CallWebTransportCreateBidirectionalStream(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_WebTransport_CreateBidirectionalStream
//go:noescape

func WebTransportCreateBidirectionalStreamFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebTransport_CreateBidirectionalStream1
//go:noescape

func CallWebTransportCreateBidirectionalStream1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebTransport_CreateBidirectionalStream1
//go:noescape

func WebTransportCreateBidirectionalStream1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebTransport_CreateUnidirectionalStream
//go:noescape

func CallWebTransportCreateUnidirectionalStream(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_WebTransport_CreateUnidirectionalStream
//go:noescape

func WebTransportCreateUnidirectionalStreamFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebTransport_CreateUnidirectionalStream1
//go:noescape

func CallWebTransportCreateUnidirectionalStream1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebTransport_CreateUnidirectionalStream1
//go:noescape

func WebTransportCreateUnidirectionalStream1Func(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_WebTransportError_StreamErrorCode
//go:noescape

func GetWebTransportErrorStreamErrorCode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

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
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_WheelEvent_DeltaY
//go:noescape

func GetWheelEventDeltaY(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_WheelEvent_DeltaZ
//go:noescape

func GetWheelEventDeltaZ(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_WheelEvent_DeltaMode
//go:noescape

func GetWheelEventDeltaMode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WindowControlsOverlayGeometryChangeEvent_Visible
//go:noescape

func GetWindowControlsOverlayGeometryChangeEventVisible(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerLocation_Href
//go:noescape

func GetWorkerLocationHref(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerLocation_Origin
//go:noescape

func GetWorkerLocationOrigin(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerLocation_Protocol
//go:noescape

func GetWorkerLocationProtocol(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerLocation_Host
//go:noescape

func GetWorkerLocationHost(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerLocation_Hostname
//go:noescape

func GetWorkerLocationHostname(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerLocation_Port
//go:noescape

func GetWorkerLocationPort(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerLocation_Pathname
//go:noescape

func GetWorkerLocationPathname(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerLocation_Search
//go:noescape

func GetWorkerLocationSearch(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerLocation_Hash
//go:noescape

func GetWorkerLocationHash(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_Hid
//go:noescape

func GetWorkerNavigatorHid(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_Serial
//go:noescape

func GetWorkerNavigatorSerial(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_Permissions
//go:noescape

func GetWorkerNavigatorPermissions(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_Usb
//go:noescape

func GetWorkerNavigatorUsb(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_MediaCapabilities
//go:noescape

func GetWorkerNavigatorMediaCapabilities(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_ServiceWorker
//go:noescape

func GetWorkerNavigatorServiceWorker(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_UserAgentData
//go:noescape

func GetWorkerNavigatorUserAgentData(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_DeviceMemory
//go:noescape

func GetWorkerNavigatorDeviceMemory(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_WorkerNavigator_Storage
//go:noescape

func GetWorkerNavigatorStorage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_StorageBuckets
//go:noescape

func GetWorkerNavigatorStorageBuckets(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_Locks
//go:noescape

func GetWorkerNavigatorLocks(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_AppCodeName
//go:noescape

func GetWorkerNavigatorAppCodeName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_AppName
//go:noescape

func GetWorkerNavigatorAppName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_AppVersion
//go:noescape

func GetWorkerNavigatorAppVersion(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_Platform
//go:noescape

func GetWorkerNavigatorPlatform(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_Product
//go:noescape

func GetWorkerNavigatorProduct(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_ProductSub
//go:noescape

func GetWorkerNavigatorProductSub(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_UserAgent
//go:noescape

func GetWorkerNavigatorUserAgent(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_Vendor
//go:noescape

func GetWorkerNavigatorVendor(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_VendorSub
//go:noescape

func GetWorkerNavigatorVendorSub(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_Oscpu
//go:noescape

func GetWorkerNavigatorOscpu(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_Language
//go:noescape

func GetWorkerNavigatorLanguage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_Languages
//go:noescape

func GetWorkerNavigatorLanguages(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_OnLine
//go:noescape

func GetWorkerNavigatorOnLine(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_HardwareConcurrency
//go:noescape

func GetWorkerNavigatorHardwareConcurrency(
	this js.Ref,
	ptr unsafe.Pointer,
) uint64

//go:wasmimport plat/js/web get_WorkerNavigator_Connection
//go:noescape

func GetWorkerNavigatorConnection(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_Ml
//go:noescape

func GetWorkerNavigatorMl(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerNavigator_Gpu
//go:noescape

func GetWorkerNavigatorGpu(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_WorkerNavigator_TaintEnabled
//go:noescape

func CallWorkerNavigatorTaintEnabled(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WorkerNavigator_TaintEnabled
//go:noescape

func WorkerNavigatorTaintEnabledFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerNavigator_SetAppBadge
//go:noescape

func CallWorkerNavigatorSetAppBadge(
	this js.Ref,
	ptr unsafe.Pointer,

	contents float64,
) js.Ref

//go:wasmimport plat/js/web func_WorkerNavigator_SetAppBadge
//go:noescape

func WorkerNavigatorSetAppBadgeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerNavigator_SetAppBadge1
//go:noescape

func CallWorkerNavigatorSetAppBadge1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WorkerNavigator_SetAppBadge1
//go:noescape

func WorkerNavigatorSetAppBadge1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerNavigator_ClearAppBadge
//go:noescape

func CallWorkerNavigatorClearAppBadge(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WorkerNavigator_ClearAppBadge
//go:noescape

func WorkerNavigatorClearAppBadgeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_WorkerGlobalScope_Self
//go:noescape

func GetWorkerGlobalScopeSelf(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerGlobalScope_Location
//go:noescape

func GetWorkerGlobalScopeLocation(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerGlobalScope_Navigator
//go:noescape

func GetWorkerGlobalScopeNavigator(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerGlobalScope_Fonts
//go:noescape

func GetWorkerGlobalScopeFonts(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerGlobalScope_Origin
//go:noescape

func GetWorkerGlobalScopeOrigin(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerGlobalScope_IsSecureContext
//go:noescape

func GetWorkerGlobalScopeIsSecureContext(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerGlobalScope_CrossOriginIsolated
//go:noescape

func GetWorkerGlobalScopeCrossOriginIsolated(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerGlobalScope_Scheduler
//go:noescape

func GetWorkerGlobalScopeScheduler(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerGlobalScope_TrustedTypes
//go:noescape

func GetWorkerGlobalScopeTrustedTypes(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerGlobalScope_Caches
//go:noescape

func GetWorkerGlobalScopeCaches(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerGlobalScope_Crypto
//go:noescape

func GetWorkerGlobalScopeCrypto(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerGlobalScope_IndexedDB
//go:noescape

func GetWorkerGlobalScopeIndexedDB(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkerGlobalScope_Performance
//go:noescape

func GetWorkerGlobalScopePerformance(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_ImportScripts
//go:noescape

func CallWorkerGlobalScopeImportScripts(
	this js.Ref,
	ptr unsafe.Pointer,

	urlsPtr unsafe.Pointer,
	urlsCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_ImportScripts
//go:noescape

func WorkerGlobalScopeImportScriptsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_ReportError
//go:noescape

func CallWorkerGlobalScopeReportError(
	this js.Ref,
	ptr unsafe.Pointer,

	e js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_ReportError
//go:noescape

func WorkerGlobalScopeReportErrorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_Btoa
//go:noescape

func CallWorkerGlobalScopeBtoa(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_Btoa
//go:noescape

func WorkerGlobalScopeBtoaFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_Atob
//go:noescape

func CallWorkerGlobalScopeAtob(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_Atob
//go:noescape

func WorkerGlobalScopeAtobFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_SetTimeout
//go:noescape

func CallWorkerGlobalScopeSetTimeout(
	this js.Ref,
	ptr unsafe.Pointer,

	handler js.Ref,
	timeout int32,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32,
) int32

//go:wasmimport plat/js/web func_WorkerGlobalScope_SetTimeout
//go:noescape

func WorkerGlobalScopeSetTimeoutFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_SetTimeout1
//go:noescape

func CallWorkerGlobalScopeSetTimeout1(
	this js.Ref,
	ptr unsafe.Pointer,

	handler js.Ref,
) int32

//go:wasmimport plat/js/web func_WorkerGlobalScope_SetTimeout1
//go:noescape

func WorkerGlobalScopeSetTimeout1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_ClearTimeout
//go:noescape

func CallWorkerGlobalScopeClearTimeout(
	this js.Ref,
	ptr unsafe.Pointer,

	id int32,
) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_ClearTimeout
//go:noescape

func WorkerGlobalScopeClearTimeoutFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_ClearTimeout1
//go:noescape

func CallWorkerGlobalScopeClearTimeout1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_ClearTimeout1
//go:noescape

func WorkerGlobalScopeClearTimeout1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_SetInterval
//go:noescape

func CallWorkerGlobalScopeSetInterval(
	this js.Ref,
	ptr unsafe.Pointer,

	handler js.Ref,
	timeout int32,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32,
) int32

//go:wasmimport plat/js/web func_WorkerGlobalScope_SetInterval
//go:noescape

func WorkerGlobalScopeSetIntervalFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_SetInterval1
//go:noescape

func CallWorkerGlobalScopeSetInterval1(
	this js.Ref,
	ptr unsafe.Pointer,

	handler js.Ref,
) int32

//go:wasmimport plat/js/web func_WorkerGlobalScope_SetInterval1
//go:noescape

func WorkerGlobalScopeSetInterval1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_ClearInterval
//go:noescape

func CallWorkerGlobalScopeClearInterval(
	this js.Ref,
	ptr unsafe.Pointer,

	id int32,
) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_ClearInterval
//go:noescape

func WorkerGlobalScopeClearIntervalFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_ClearInterval1
//go:noescape

func CallWorkerGlobalScopeClearInterval1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_ClearInterval1
//go:noescape

func WorkerGlobalScopeClearInterval1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_QueueMicrotask
//go:noescape

func CallWorkerGlobalScopeQueueMicrotask(
	this js.Ref,
	ptr unsafe.Pointer,

	callback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_QueueMicrotask
//go:noescape

func WorkerGlobalScopeQueueMicrotaskFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_CreateImageBitmap
//go:noescape

func CallWorkerGlobalScopeCreateImageBitmap(
	this js.Ref,
	ptr unsafe.Pointer,

	image js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_CreateImageBitmap
//go:noescape

func WorkerGlobalScopeCreateImageBitmapFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_CreateImageBitmap1
//go:noescape

func CallWorkerGlobalScopeCreateImageBitmap1(
	this js.Ref,
	ptr unsafe.Pointer,

	image js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_CreateImageBitmap1
//go:noescape

func WorkerGlobalScopeCreateImageBitmap1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_CreateImageBitmap2
//go:noescape

func CallWorkerGlobalScopeCreateImageBitmap2(
	this js.Ref,
	ptr unsafe.Pointer,

	image js.Ref,
	sx int32,
	sy int32,
	sw int32,
	sh int32,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_CreateImageBitmap2
//go:noescape

func WorkerGlobalScopeCreateImageBitmap2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_CreateImageBitmap3
//go:noescape

func CallWorkerGlobalScopeCreateImageBitmap3(
	this js.Ref,
	ptr unsafe.Pointer,

	image js.Ref,
	sx int32,
	sy int32,
	sw int32,
	sh int32,
) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_CreateImageBitmap3
//go:noescape

func WorkerGlobalScopeCreateImageBitmap3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_StructuredClone
//go:noescape

func CallWorkerGlobalScopeStructuredClone(
	this js.Ref,
	ptr unsafe.Pointer,

	value js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_StructuredClone
//go:noescape

func WorkerGlobalScopeStructuredCloneFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_StructuredClone1
//go:noescape

func CallWorkerGlobalScopeStructuredClone1(
	this js.Ref,
	ptr unsafe.Pointer,

	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_StructuredClone1
//go:noescape

func WorkerGlobalScopeStructuredClone1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_Fetch
//go:noescape

func CallWorkerGlobalScopeFetch(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
	init unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_Fetch
//go:noescape

func WorkerGlobalScopeFetchFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkerGlobalScope_Fetch1
//go:noescape

func CallWorkerGlobalScopeFetch1(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WorkerGlobalScope_Fetch1
//go:noescape

func WorkerGlobalScopeFetch1Func(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WorkletAnimationEffect_LocalTime
//go:noescape

func GetWorkletAnimationEffectLocalTime(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_WorkletAnimationEffect_LocalTime
//go:noescape

func SetWorkletAnimationEffectLocalTime(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web call_WorkletAnimationEffect_GetTiming
//go:noescape

func CallWorkletAnimationEffectGetTiming(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WorkletAnimationEffect_GetTiming
//go:noescape

func WorkletAnimationEffectGetTimingFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkletAnimationEffect_GetComputedTiming
//go:noescape

func CallWorkletAnimationEffectGetComputedTiming(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WorkletAnimationEffect_GetComputedTiming
//go:noescape

func WorkletAnimationEffectGetComputedTimingFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkletGroupEffect_GetChildren
//go:noescape

func CallWorkletGroupEffectGetChildren(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WorkletGroupEffect_GetChildren
//go:noescape

func WorkletGroupEffectGetChildrenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_XMLHttpRequestResponseType
//go:noescape

func ConstOfXMLHttpRequestResponseType(str js.Ref) uint32

//go:wasmimport plat/js/web get_XMLHttpRequest_ReadyState
//go:noescape

func GetXMLHttpRequestReadyState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XMLHttpRequest_Timeout
//go:noescape

func GetXMLHttpRequestTimeout(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_XMLHttpRequest_Timeout
//go:noescape

func SetXMLHttpRequestTimeout(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_XMLHttpRequest_WithCredentials
//go:noescape

func GetXMLHttpRequestWithCredentials(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_XMLHttpRequest_WithCredentials
//go:noescape

func SetXMLHttpRequestWithCredentials(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XMLHttpRequest_Upload
//go:noescape

func GetXMLHttpRequestUpload(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XMLHttpRequest_ResponseURL
//go:noescape

func GetXMLHttpRequestResponseURL(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XMLHttpRequest_Status
//go:noescape

func GetXMLHttpRequestStatus(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XMLHttpRequest_StatusText
//go:noescape

func GetXMLHttpRequestStatusText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XMLHttpRequest_ResponseType
//go:noescape

func GetXMLHttpRequestResponseType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_XMLHttpRequest_ResponseType
//go:noescape

func SetXMLHttpRequestResponseType(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_XMLHttpRequest_Response
//go:noescape

func GetXMLHttpRequestResponse(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XMLHttpRequest_ResponseText
//go:noescape

func GetXMLHttpRequestResponseText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XMLHttpRequest_ResponseXML
//go:noescape

func GetXMLHttpRequestResponseXML(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_XMLHttpRequest_Open
//go:noescape

func CallXMLHttpRequestOpen(
	this js.Ref,
	ptr unsafe.Pointer,

	method js.Ref,
	url js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_Open
//go:noescape

func XMLHttpRequestOpenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XMLHttpRequest_Open1
//go:noescape

func CallXMLHttpRequestOpen1(
	this js.Ref,
	ptr unsafe.Pointer,

	method js.Ref,
	url js.Ref,
	async js.Ref,
	username js.Ref,
	password js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_Open1
//go:noescape

func XMLHttpRequestOpen1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XMLHttpRequest_Open2
//go:noescape

func CallXMLHttpRequestOpen2(
	this js.Ref,
	ptr unsafe.Pointer,

	method js.Ref,
	url js.Ref,
	async js.Ref,
	username js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_Open2
//go:noescape

func XMLHttpRequestOpen2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XMLHttpRequest_Open3
//go:noescape

func CallXMLHttpRequestOpen3(
	this js.Ref,
	ptr unsafe.Pointer,

	method js.Ref,
	url js.Ref,
	async js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_Open3
//go:noescape

func XMLHttpRequestOpen3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XMLHttpRequest_SetRequestHeader
//go:noescape

func CallXMLHttpRequestSetRequestHeader(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_SetRequestHeader
//go:noescape

func XMLHttpRequestSetRequestHeaderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XMLHttpRequest_Send
//go:noescape

func CallXMLHttpRequestSend(
	this js.Ref,
	ptr unsafe.Pointer,

	body js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_Send
//go:noescape

func XMLHttpRequestSendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XMLHttpRequest_Send1
//go:noescape

func CallXMLHttpRequestSend1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_Send1
//go:noescape

func XMLHttpRequestSend1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XMLHttpRequest_Abort
//go:noescape

func CallXMLHttpRequestAbort(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_Abort
//go:noescape

func XMLHttpRequestAbortFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XMLHttpRequest_GetResponseHeader
//go:noescape

func CallXMLHttpRequestGetResponseHeader(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_GetResponseHeader
//go:noescape

func XMLHttpRequestGetResponseHeaderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XMLHttpRequest_GetAllResponseHeaders
//go:noescape

func CallXMLHttpRequestGetAllResponseHeaders(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_GetAllResponseHeaders
//go:noescape

func XMLHttpRequestGetAllResponseHeadersFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XMLHttpRequest_OverrideMimeType
//go:noescape

func CallXMLHttpRequestOverrideMimeType(
	this js.Ref,
	ptr unsafe.Pointer,

	mime js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_OverrideMimeType
//go:noescape

func XMLHttpRequestOverrideMimeTypeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XMLHttpRequest_SetAttributionReporting
//go:noescape

func CallXMLHttpRequestSetAttributionReporting(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_SetAttributionReporting
//go:noescape

func XMLHttpRequestSetAttributionReportingFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XMLHttpRequest_SetPrivateToken
//go:noescape

func CallXMLHttpRequestSetPrivateToken(
	this js.Ref,
	ptr unsafe.Pointer,

	privateToken unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_XMLHttpRequest_SetPrivateToken
//go:noescape

func XMLHttpRequestSetPrivateTokenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XMLSerializer_SerializeToString
//go:noescape

func CallXMLSerializerSerializeToString(
	this js.Ref,
	ptr unsafe.Pointer,

	root js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XMLSerializer_SerializeToString
//go:noescape

func XMLSerializerSerializeToStringFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathEvaluator_CreateExpression
//go:noescape

func CallXPathEvaluatorCreateExpression(
	this js.Ref,
	ptr unsafe.Pointer,

	expression js.Ref,
	resolver js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XPathEvaluator_CreateExpression
//go:noescape

func XPathEvaluatorCreateExpressionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathEvaluator_CreateExpression1
//go:noescape

func CallXPathEvaluatorCreateExpression1(
	this js.Ref,
	ptr unsafe.Pointer,

	expression js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XPathEvaluator_CreateExpression1
//go:noescape

func XPathEvaluatorCreateExpression1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathEvaluator_CreateNSResolver
//go:noescape

func CallXPathEvaluatorCreateNSResolver(
	this js.Ref,
	ptr unsafe.Pointer,

	nodeResolver js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XPathEvaluator_CreateNSResolver
//go:noescape

func XPathEvaluatorCreateNSResolverFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathEvaluator_Evaluate
//go:noescape

func CallXPathEvaluatorEvaluate(
	this js.Ref,
	ptr unsafe.Pointer,

	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref,
	typ uint32,
	result js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XPathEvaluator_Evaluate
//go:noescape

func XPathEvaluatorEvaluateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathEvaluator_Evaluate1
//go:noescape

func CallXPathEvaluatorEvaluate1(
	this js.Ref,
	ptr unsafe.Pointer,

	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref,
	typ uint32,
) js.Ref

//go:wasmimport plat/js/web func_XPathEvaluator_Evaluate1
//go:noescape

func XPathEvaluatorEvaluate1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathEvaluator_Evaluate2
//go:noescape

func CallXPathEvaluatorEvaluate2(
	this js.Ref,
	ptr unsafe.Pointer,

	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XPathEvaluator_Evaluate2
//go:noescape

func XPathEvaluatorEvaluate2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathEvaluator_Evaluate3
//go:noescape

func CallXPathEvaluatorEvaluate3(
	this js.Ref,
	ptr unsafe.Pointer,

	expression js.Ref,
	contextNode js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XPathEvaluator_Evaluate3
//go:noescape

func XPathEvaluatorEvaluate3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_XRBoundedReferenceSpace_BoundsGeometry
//go:noescape

func GetXRBoundedReferenceSpaceBoundsGeometry(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref
