// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

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

//go:wasmimport plat/js/web store_RTCTrackEventInit
//go:noescape
func RTCTrackEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCTrackEventInit
//go:noescape
func RTCTrackEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_RTCTrackEvent_RTCTrackEvent
//go:noescape
func NewRTCTrackEventByRTCTrackEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_RTCTrackEvent_Receiver
//go:noescape
func GetRTCTrackEventReceiver(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCTrackEvent_Track
//go:noescape
func GetRTCTrackEventTrack(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCTrackEvent_Streams
//go:noescape
func GetRTCTrackEventStreams(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCTrackEvent_Transceiver
//go:noescape
func GetRTCTrackEventTransceiver(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_RTCTransformEvent_RTCTransformEvent
//go:noescape
func NewRTCTransformEventByRTCTransformEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_RTCTransformEvent_RTCTransformEvent1
//go:noescape
func NewRTCTransformEventByRTCTransformEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_RTCTransformEvent_Transformer
//go:noescape
func GetRTCTransformEventTransformer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_RTCTransportStats
//go:noescape
func RTCTransportStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCTransportStats
//go:noescape
func RTCTransportStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCVideoSourceStats
//go:noescape
func RTCVideoSourceStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCVideoSourceStats
//go:noescape
func RTCVideoSourceStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_ReadableStreamBYOBRequest_View
//go:noescape
func GetReadableStreamBYOBRequestView(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStreamBYOBRequest_Respond
//go:noescape
func HasReadableStreamBYOBRequestRespond(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamBYOBRequest_Respond
//go:noescape
func ReadableStreamBYOBRequestRespondFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStreamBYOBRequest_Respond
//go:noescape
func CallReadableStreamBYOBRequestRespond(
	this js.Ref, retPtr unsafe.Pointer,
	bytesWritten float64)

//go:wasmimport plat/js/web try_ReadableStreamBYOBRequest_Respond
//go:noescape
func TryReadableStreamBYOBRequestRespond(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bytesWritten float64) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStreamBYOBRequest_RespondWithNewView
//go:noescape
func HasReadableStreamBYOBRequestRespondWithNewView(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamBYOBRequest_RespondWithNewView
//go:noescape
func ReadableStreamBYOBRequestRespondWithNewViewFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStreamBYOBRequest_RespondWithNewView
//go:noescape
func CallReadableStreamBYOBRequestRespondWithNewView(
	this js.Ref, retPtr unsafe.Pointer,
	view js.Ref)

//go:wasmimport plat/js/web try_ReadableStreamBYOBRequest_RespondWithNewView
//go:noescape
func TryReadableStreamBYOBRequestRespondWithNewView(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	view js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_ReadableByteStreamController_ByobRequest
//go:noescape
func GetReadableByteStreamControllerByobRequest(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ReadableByteStreamController_DesiredSize
//go:noescape
func GetReadableByteStreamControllerDesiredSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableByteStreamController_Close
//go:noescape
func HasReadableByteStreamControllerClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableByteStreamController_Close
//go:noescape
func ReadableByteStreamControllerCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableByteStreamController_Close
//go:noescape
func CallReadableByteStreamControllerClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableByteStreamController_Close
//go:noescape
func TryReadableByteStreamControllerClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableByteStreamController_Enqueue
//go:noescape
func HasReadableByteStreamControllerEnqueue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableByteStreamController_Enqueue
//go:noescape
func ReadableByteStreamControllerEnqueueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableByteStreamController_Enqueue
//go:noescape
func CallReadableByteStreamControllerEnqueue(
	this js.Ref, retPtr unsafe.Pointer,
	chunk js.Ref)

//go:wasmimport plat/js/web try_ReadableByteStreamController_Enqueue
//go:noescape
func TryReadableByteStreamControllerEnqueue(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	chunk js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableByteStreamController_Error
//go:noescape
func HasReadableByteStreamControllerError(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableByteStreamController_Error
//go:noescape
func ReadableByteStreamControllerErrorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableByteStreamController_Error
//go:noescape
func CallReadableByteStreamControllerError(
	this js.Ref, retPtr unsafe.Pointer,
	e js.Ref)

//go:wasmimport plat/js/web try_ReadableByteStreamController_Error
//go:noescape
func TryReadableByteStreamControllerError(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	e js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableByteStreamController_Error1
//go:noescape
func HasReadableByteStreamControllerError1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableByteStreamController_Error1
//go:noescape
func ReadableByteStreamControllerError1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableByteStreamController_Error1
//go:noescape
func CallReadableByteStreamControllerError1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableByteStreamController_Error1
//go:noescape
func TryReadableByteStreamControllerError1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ReadableStreamDefaultController_DesiredSize
//go:noescape
func GetReadableStreamDefaultControllerDesiredSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStreamDefaultController_Close
//go:noescape
func HasReadableStreamDefaultControllerClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamDefaultController_Close
//go:noescape
func ReadableStreamDefaultControllerCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStreamDefaultController_Close
//go:noescape
func CallReadableStreamDefaultControllerClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableStreamDefaultController_Close
//go:noescape
func TryReadableStreamDefaultControllerClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStreamDefaultController_Enqueue
//go:noescape
func HasReadableStreamDefaultControllerEnqueue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamDefaultController_Enqueue
//go:noescape
func ReadableStreamDefaultControllerEnqueueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStreamDefaultController_Enqueue
//go:noescape
func CallReadableStreamDefaultControllerEnqueue(
	this js.Ref, retPtr unsafe.Pointer,
	chunk js.Ref)

//go:wasmimport plat/js/web try_ReadableStreamDefaultController_Enqueue
//go:noescape
func TryReadableStreamDefaultControllerEnqueue(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	chunk js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStreamDefaultController_Enqueue1
//go:noescape
func HasReadableStreamDefaultControllerEnqueue1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamDefaultController_Enqueue1
//go:noescape
func ReadableStreamDefaultControllerEnqueue1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStreamDefaultController_Enqueue1
//go:noescape
func CallReadableStreamDefaultControllerEnqueue1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableStreamDefaultController_Enqueue1
//go:noescape
func TryReadableStreamDefaultControllerEnqueue1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStreamDefaultController_Error
//go:noescape
func HasReadableStreamDefaultControllerError(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamDefaultController_Error
//go:noescape
func ReadableStreamDefaultControllerErrorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStreamDefaultController_Error
//go:noescape
func CallReadableStreamDefaultControllerError(
	this js.Ref, retPtr unsafe.Pointer,
	e js.Ref)

//go:wasmimport plat/js/web try_ReadableStreamDefaultController_Error
//go:noescape
func TryReadableStreamDefaultControllerError(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	e js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStreamDefaultController_Error1
//go:noescape
func HasReadableStreamDefaultControllerError1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamDefaultController_Error1
//go:noescape
func ReadableStreamDefaultControllerError1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStreamDefaultController_Error1
//go:noescape
func CallReadableStreamDefaultControllerError1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableStreamDefaultController_Error1
//go:noescape
func TryReadableStreamDefaultControllerError1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ReadableStreamIteratorOptions
//go:noescape
func ReadableStreamIteratorOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ReadableStreamIteratorOptions
//go:noescape
func ReadableStreamIteratorOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_ReadableStreamType
//go:noescape
func ConstOfReadableStreamType(str js.Ref) uint32

//go:wasmimport plat/js/web store_RelativeOrientationReadingValues
//go:noescape
func RelativeOrientationReadingValuesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RelativeOrientationReadingValues
//go:noescape
func RelativeOrientationReadingValuesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_RelativeOrientationSensor_RelativeOrientationSensor
//go:noescape
func NewRelativeOrientationSensorByRelativeOrientationSensor(
	sensorOptions unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_RelativeOrientationSensor_RelativeOrientationSensor1
//go:noescape
func NewRelativeOrientationSensorByRelativeOrientationSensor1() js.Ref

//go:wasmimport plat/js/web has_ReportBody_ToJSON
//go:noescape
func HasReportBodyToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReportBody_ToJSON
//go:noescape
func ReportBodyToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReportBody_ToJSON
//go:noescape
func CallReportBodyToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReportBody_ToJSON
//go:noescape
func TryReportBodyToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Report_Type
//go:noescape
func GetReportType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Report_Url
//go:noescape
func GetReportUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Report_Body
//go:noescape
func GetReportBody(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Report_ToJSON
//go:noescape
func HasReportToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Report_ToJSON
//go:noescape
func ReportToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Report_ToJSON
//go:noescape
func CallReportToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Report_ToJSON
//go:noescape
func TryReportToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ReportResultBrowserSignals
//go:noescape
func ReportResultBrowserSignalsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ReportResultBrowserSignals
//go:noescape
func ReportResultBrowserSignalsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ReportWinBrowserSignals
//go:noescape
func ReportWinBrowserSignalsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ReportWinBrowserSignals
//go:noescape
func ReportWinBrowserSignalsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ReportingBrowserSignals
//go:noescape
func ReportingBrowserSignalsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ReportingBrowserSignals
//go:noescape
func ReportingBrowserSignalsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ReportingObserverOptions
//go:noescape
func ReportingObserverOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ReportingObserverOptions
//go:noescape
func ReportingObserverOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ReportingObserver_ReportingObserver
//go:noescape
func NewReportingObserverByReportingObserver(
	callback js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ReportingObserver_ReportingObserver1
//go:noescape
func NewReportingObserverByReportingObserver1(
	callback js.Ref) js.Ref

//go:wasmimport plat/js/web has_ReportingObserver_Observe
//go:noescape
func HasReportingObserverObserve(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReportingObserver_Observe
//go:noescape
func ReportingObserverObserveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReportingObserver_Observe
//go:noescape
func CallReportingObserverObserve(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReportingObserver_Observe
//go:noescape
func TryReportingObserverObserve(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReportingObserver_Disconnect
//go:noescape
func HasReportingObserverDisconnect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReportingObserver_Disconnect
//go:noescape
func ReportingObserverDisconnectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReportingObserver_Disconnect
//go:noescape
func CallReportingObserverDisconnect(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReportingObserver_Disconnect
//go:noescape
func TryReportingObserverDisconnect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReportingObserver_TakeRecords
//go:noescape
func HasReportingObserverTakeRecords(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReportingObserver_TakeRecords
//go:noescape
func ReportingObserverTakeRecordsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReportingObserver_TakeRecords
//go:noescape
func CallReportingObserverTakeRecords(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReportingObserver_TakeRecords
//go:noescape
func TryReportingObserverTakeRecords(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_ResidentKeyRequirement
//go:noescape
func ConstOfResidentKeyRequirement(str js.Ref) uint32

//go:wasmimport plat/js/web get_ResizeObserverSize_InlineSize
//go:noescape
func GetResizeObserverSizeInlineSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ResizeObserverSize_BlockSize
//go:noescape
func GetResizeObserverSizeBlockSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ResizeObserverEntry_Target
//go:noescape
func GetResizeObserverEntryTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ResizeObserverEntry_ContentRect
//go:noescape
func GetResizeObserverEntryContentRect(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ResizeObserverEntry_BorderBoxSize
//go:noescape
func GetResizeObserverEntryBorderBoxSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ResizeObserverEntry_ContentBoxSize
//go:noescape
func GetResizeObserverEntryContentBoxSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ResizeObserverEntry_DevicePixelContentBoxSize
//go:noescape
func GetResizeObserverEntryDevicePixelContentBoxSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_ResizeObserverBoxOptions
//go:noescape
func ConstOfResizeObserverBoxOptions(str js.Ref) uint32

//go:wasmimport plat/js/web store_ResizeObserverOptions
//go:noescape
func ResizeObserverOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ResizeObserverOptions
//go:noescape
func ResizeObserverOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ResizeObserver_ResizeObserver
//go:noescape
func NewResizeObserverByResizeObserver(
	callback js.Ref) js.Ref

//go:wasmimport plat/js/web has_ResizeObserver_Observe
//go:noescape
func HasResizeObserverObserve(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ResizeObserver_Observe
//go:noescape
func ResizeObserverObserveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ResizeObserver_Observe
//go:noescape
func CallResizeObserverObserve(
	this js.Ref, retPtr unsafe.Pointer,
	target js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_ResizeObserver_Observe
//go:noescape
func TryResizeObserverObserve(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	target js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ResizeObserver_Observe1
//go:noescape
func HasResizeObserverObserve1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ResizeObserver_Observe1
//go:noescape
func ResizeObserverObserve1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ResizeObserver_Observe1
//go:noescape
func CallResizeObserverObserve1(
	this js.Ref, retPtr unsafe.Pointer,
	target js.Ref)

//go:wasmimport plat/js/web try_ResizeObserver_Observe1
//go:noescape
func TryResizeObserverObserve1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	target js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ResizeObserver_Unobserve
//go:noescape
func HasResizeObserverUnobserve(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ResizeObserver_Unobserve
//go:noescape
func ResizeObserverUnobserveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ResizeObserver_Unobserve
//go:noescape
func CallResizeObserverUnobserve(
	this js.Ref, retPtr unsafe.Pointer,
	target js.Ref)

//go:wasmimport plat/js/web try_ResizeObserver_Unobserve
//go:noescape
func TryResizeObserverUnobserve(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	target js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ResizeObserver_Disconnect
//go:noescape
func HasResizeObserverDisconnect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ResizeObserver_Disconnect
//go:noescape
func ResizeObserverDisconnectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ResizeObserver_Disconnect
//go:noescape
func CallResizeObserverDisconnect(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ResizeObserver_Disconnect
//go:noescape
func TryResizeObserverDisconnect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_RsaHashedImportParams
//go:noescape
func RsaHashedImportParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RsaHashedImportParams
//go:noescape
func RsaHashedImportParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RsaHashedKeyAlgorithm
//go:noescape
func RsaHashedKeyAlgorithmJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RsaHashedKeyAlgorithm
//go:noescape
func RsaHashedKeyAlgorithmJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RsaHashedKeyGenParams
//go:noescape
func RsaHashedKeyGenParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RsaHashedKeyGenParams
//go:noescape
func RsaHashedKeyGenParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RsaKeyAlgorithm
//go:noescape
func RsaKeyAlgorithmJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RsaKeyAlgorithm
//go:noescape
func RsaKeyAlgorithmJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RsaKeyGenParams
//go:noescape
func RsaKeyGenParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RsaKeyGenParams
//go:noescape
func RsaKeyGenParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RsaOaepParams
//go:noescape
func RsaOaepParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RsaOaepParams
//go:noescape
func RsaOaepParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RsaPssParams
//go:noescape
func RsaPssParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RsaPssParams
//go:noescape
func RsaPssParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_SFrameTransformErrorEventType
//go:noescape
func ConstOfSFrameTransformErrorEventType(str js.Ref) uint32

//go:wasmimport plat/js/web store_SFrameTransformErrorEventInit
//go:noescape
func SFrameTransformErrorEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SFrameTransformErrorEventInit
//go:noescape
func SFrameTransformErrorEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_SFrameTransformErrorEvent_SFrameTransformErrorEvent
//go:noescape
func NewSFrameTransformErrorEventBySFrameTransformErrorEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_SFrameTransformErrorEvent_ErrorType
//go:noescape
func GetSFrameTransformErrorEventErrorType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SFrameTransformErrorEvent_KeyID
//go:noescape
func GetSFrameTransformErrorEventKeyID(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SFrameTransformErrorEvent_Frame
//go:noescape
func GetSFrameTransformErrorEventFrame(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAElement_Target
//go:noescape
func GetSVGAElementTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAElement_Download
//go:noescape
func GetSVGAElementDownload(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_Download
//go:noescape
func SetSVGAElementDownload(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_Ping
//go:noescape
func GetSVGAElementPing(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_Ping
//go:noescape
func SetSVGAElementPing(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_Rel
//go:noescape
func GetSVGAElementRel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_Rel
//go:noescape
func SetSVGAElementRel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_RelList
//go:noescape
func GetSVGAElementRelList(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAElement_Hreflang
//go:noescape
func GetSVGAElementHreflang(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_Hreflang
//go:noescape
func SetSVGAElementHreflang(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_Type
//go:noescape
func GetSVGAElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_Type
//go:noescape
func SetSVGAElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_Text
//go:noescape
func GetSVGAElementText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_Text
//go:noescape
func SetSVGAElementText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_ReferrerPolicy
//go:noescape
func GetSVGAElementReferrerPolicy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_ReferrerPolicy
//go:noescape
func SetSVGAElementReferrerPolicy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_Origin
//go:noescape
func GetSVGAElementOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAElement_Protocol
//go:noescape
func GetSVGAElementProtocol(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_Protocol
//go:noescape
func SetSVGAElementProtocol(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_Username
//go:noescape
func GetSVGAElementUsername(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_Username
//go:noescape
func SetSVGAElementUsername(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_Password
//go:noescape
func GetSVGAElementPassword(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_Password
//go:noescape
func SetSVGAElementPassword(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_Host
//go:noescape
func GetSVGAElementHost(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_Host
//go:noescape
func SetSVGAElementHost(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_Hostname
//go:noescape
func GetSVGAElementHostname(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_Hostname
//go:noescape
func SetSVGAElementHostname(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_Port
//go:noescape
func GetSVGAElementPort(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_Port
//go:noescape
func SetSVGAElementPort(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_Pathname
//go:noescape
func GetSVGAElementPathname(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_Pathname
//go:noescape
func SetSVGAElementPathname(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_Search
//go:noescape
func GetSVGAElementSearch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_Search
//go:noescape
func SetSVGAElementSearch(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_Hash
//go:noescape
func GetSVGAElementHash(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAElement_Hash
//go:noescape
func SetSVGAElementHash(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAElement_Href
//go:noescape
func GetSVGAElementHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedAngle_BaseVal
//go:noescape
func GetSVGAnimatedAngleBaseVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedAngle_AnimVal
//go:noescape
func GetSVGAnimatedAngleAnimVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedBoolean_BaseVal
//go:noescape
func GetSVGAnimatedBooleanBaseVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAnimatedBoolean_BaseVal
//go:noescape
func SetSVGAnimatedBooleanBaseVal(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedBoolean_AnimVal
//go:noescape
func GetSVGAnimatedBooleanAnimVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)
