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

//go:wasmimport plat/js/web store_IsInputPendingOptions
//go:noescape

func IsInputPendingOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IsInputPendingOptions
//go:noescape

func IsInputPendingOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_Scheduling_IsInputPending
//go:noescape

func CallSchedulingIsInputPending(
	this js.Ref,
	ptr unsafe.Pointer,

	isInputPendingOptions unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Scheduling_IsInputPending
//go:noescape

func SchedulingIsInputPendingFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Scheduling_IsInputPending1
//go:noescape

func CallSchedulingIsInputPending1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Scheduling_IsInputPending1
//go:noescape

func SchedulingIsInputPending1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_WakeLockType
//go:noescape

func ConstOfWakeLockType(str js.Ref) uint32

//go:wasmimport plat/js/web get_WakeLockSentinel_Released
//go:noescape

func GetWakeLockSentinelReleased(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WakeLockSentinel_Type
//go:noescape

func GetWakeLockSentinelType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_WakeLockSentinel_Release
//go:noescape

func CallWakeLockSentinelRelease(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WakeLockSentinel_Release
//go:noescape

func WakeLockSentinelReleaseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WakeLock_Request
//go:noescape

func CallWakeLockRequest(
	this js.Ref,
	ptr unsafe.Pointer,

	typ uint32,
) js.Ref

//go:wasmimport plat/js/web func_WakeLock_Request
//go:noescape

func WakeLockRequestFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WakeLock_Request1
//go:noescape

func CallWakeLockRequest1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WakeLock_Request1
//go:noescape

func WakeLockRequest1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_PointerEventInit
//go:noescape

func PointerEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PointerEventInit
//go:noescape

func PointerEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PointerEvent_PointerEvent
//go:noescape

func NewPointerEventByPointerEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_PointerEvent_PointerEvent1
//go:noescape

func NewPointerEventByPointerEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_PointerEvent_PointerId
//go:noescape

func GetPointerEventPointerId(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_PointerEvent_Width
//go:noescape

func GetPointerEventWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_PointerEvent_Height
//go:noescape

func GetPointerEventHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_PointerEvent_Pressure
//go:noescape

func GetPointerEventPressure(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_PointerEvent_TangentialPressure
//go:noescape

func GetPointerEventTangentialPressure(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_PointerEvent_TiltX
//go:noescape

func GetPointerEventTiltX(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_PointerEvent_TiltY
//go:noescape

func GetPointerEventTiltY(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_PointerEvent_Twist
//go:noescape

func GetPointerEventTwist(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_PointerEvent_AltitudeAngle
//go:noescape

func GetPointerEventAltitudeAngle(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_PointerEvent_AzimuthAngle
//go:noescape

func GetPointerEventAzimuthAngle(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_PointerEvent_PointerType
//go:noescape

func GetPointerEventPointerType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PointerEvent_IsPrimary
//go:noescape

func GetPointerEventIsPrimary(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_PointerEvent_GetCoalescedEvents
//go:noescape

func CallPointerEventGetCoalescedEvents(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PointerEvent_GetCoalescedEvents
//go:noescape

func PointerEventGetCoalescedEventsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PointerEvent_GetPredictedEvents
//go:noescape

func CallPointerEventGetPredictedEvents(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PointerEvent_GetPredictedEvents
//go:noescape

func PointerEventGetPredictedEventsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_InkTrailStyle
//go:noescape

func InkTrailStyleJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_InkTrailStyle
//go:noescape

func InkTrailStyleJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_InkPresenter_PresentationArea
//go:noescape

func GetInkPresenterPresentationArea(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_InkPresenter_ExpectedImprovement
//go:noescape

func GetInkPresenterExpectedImprovement(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_InkPresenter_UpdateInkTrailStartPoint
//go:noescape

func CallInkPresenterUpdateInkTrailStartPoint(
	this js.Ref,
	ptr unsafe.Pointer,

	event js.Ref,
	style unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_InkPresenter_UpdateInkTrailStartPoint
//go:noescape

func InkPresenterUpdateInkTrailStartPointFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_InkPresenterParam
//go:noescape

func InkPresenterParamJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_InkPresenterParam
//go:noescape

func InkPresenterParamJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_Ink_RequestPresenter
//go:noescape

func CallInkRequestPresenter(
	this js.Ref,
	ptr unsafe.Pointer,

	param unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Ink_RequestPresenter
//go:noescape

func InkRequestPresenterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Ink_RequestPresenter1
//go:noescape

func CallInkRequestPresenter1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Ink_RequestPresenter1
//go:noescape

func InkRequestPresenter1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_PresentationConnectionState
//go:noescape

func ConstOfPresentationConnectionState(str js.Ref) uint32

//go:wasmimport plat/js/web get_PresentationConnection_Id
//go:noescape

func GetPresentationConnectionId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PresentationConnection_Url
//go:noescape

func GetPresentationConnectionUrl(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PresentationConnection_State
//go:noescape

func GetPresentationConnectionState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_PresentationConnection_BinaryType
//go:noescape

func GetPresentationConnectionBinaryType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_PresentationConnection_BinaryType
//go:noescape

func SetPresentationConnectionBinaryType(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web call_PresentationConnection_Close
//go:noescape

func CallPresentationConnectionClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PresentationConnection_Close
//go:noescape

func PresentationConnectionCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PresentationConnection_Terminate
//go:noescape

func CallPresentationConnectionTerminate(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PresentationConnection_Terminate
//go:noescape

func PresentationConnectionTerminateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PresentationConnection_Send
//go:noescape

func CallPresentationConnectionSend(
	this js.Ref,
	ptr unsafe.Pointer,

	message js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PresentationConnection_Send
//go:noescape

func PresentationConnectionSendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PresentationConnection_Send1
//go:noescape

func CallPresentationConnectionSend1(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PresentationConnection_Send1
//go:noescape

func PresentationConnectionSend1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PresentationConnection_Send2
//go:noescape

func CallPresentationConnectionSend2(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PresentationConnection_Send2
//go:noescape

func PresentationConnectionSend2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PresentationConnection_Send3
//go:noescape

func CallPresentationConnectionSend3(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PresentationConnection_Send3
//go:noescape

func PresentationConnectionSend3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_PresentationAvailability_Value
//go:noescape

func GetPresentationAvailabilityValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web new_PresentationRequest_PresentationRequest
//go:noescape

func NewPresentationRequestByPresentationRequest(
	url js.Ref) js.Ref

//go:wasmimport plat/js/web new_PresentationRequest_PresentationRequest1
//go:noescape

func NewPresentationRequestByPresentationRequest1(
	urls js.Ref) js.Ref

//go:wasmimport plat/js/web call_PresentationRequest_Start
//go:noescape

func CallPresentationRequestStart(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PresentationRequest_Start
//go:noescape

func PresentationRequestStartFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PresentationRequest_Reconnect
//go:noescape

func CallPresentationRequestReconnect(
	this js.Ref,
	ptr unsafe.Pointer,

	presentationId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PresentationRequest_Reconnect
//go:noescape

func PresentationRequestReconnectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PresentationRequest_GetAvailability
//go:noescape

func CallPresentationRequestGetAvailability(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PresentationRequest_GetAvailability
//go:noescape

func PresentationRequestGetAvailabilityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_PresentationConnectionList_Connections
//go:noescape

func GetPresentationConnectionListConnections(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PresentationReceiver_ConnectionList
//go:noescape

func GetPresentationReceiverConnectionList(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Presentation_DefaultRequest
//go:noescape

func GetPresentationDefaultRequest(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Presentation_DefaultRequest
//go:noescape

func SetPresentationDefaultRequest(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Presentation_Receiver
//go:noescape

func GetPresentationReceiver(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_VirtualKeyboard_BoundingRect
//go:noescape

func GetVirtualKeyboardBoundingRect(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_VirtualKeyboard_OverlaysContent
//go:noescape

func GetVirtualKeyboardOverlaysContent(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_VirtualKeyboard_OverlaysContent
//go:noescape

func SetVirtualKeyboardOverlaysContent(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_VirtualKeyboard_Show
//go:noescape

func CallVirtualKeyboardShow(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_VirtualKeyboard_Show
//go:noescape

func VirtualKeyboardShowFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VirtualKeyboard_Hide
//go:noescape

func CallVirtualKeyboardHide(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_VirtualKeyboard_Hide
//go:noescape

func VirtualKeyboardHideFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_NavigatorUABrandVersion
//go:noescape

func NavigatorUABrandVersionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NavigatorUABrandVersion
//go:noescape

func NavigatorUABrandVersionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_UADataValues
//go:noescape

func UADataValuesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_UADataValues
//go:noescape

func UADataValuesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_UALowEntropyJSON
//go:noescape

func UALowEntropyJSONJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_UALowEntropyJSON
//go:noescape

func UALowEntropyJSONJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_NavigatorUAData_Brands
//go:noescape

func GetNavigatorUADataBrands(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NavigatorUAData_Mobile
//go:noescape

func GetNavigatorUADataMobile(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NavigatorUAData_Platform
//go:noescape

func GetNavigatorUADataPlatform(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_NavigatorUAData_GetHighEntropyValues
//go:noescape

func CallNavigatorUADataGetHighEntropyValues(
	this js.Ref,
	ptr unsafe.Pointer,

	hints js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_NavigatorUAData_GetHighEntropyValues
//go:noescape

func NavigatorUADataGetHighEntropyValuesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NavigatorUAData_ToJSON
//go:noescape

func CallNavigatorUADataToJSON(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_NavigatorUAData_ToJSON
//go:noescape

func NavigatorUADataToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_StorageEstimate
//go:noescape

func StorageEstimateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_StorageEstimate
//go:noescape

func StorageEstimateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageManager_Persisted
//go:noescape

func CallStorageManagerPersisted(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_StorageManager_Persisted
//go:noescape

func StorageManagerPersistedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageManager_Persist
//go:noescape

func CallStorageManagerPersist(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_StorageManager_Persist
//go:noescape

func StorageManagerPersistFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageManager_Estimate
//go:noescape

func CallStorageManagerEstimate(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_StorageManager_Estimate
//go:noescape

func StorageManagerEstimateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageManager_GetDirectory
//go:noescape

func CallStorageManagerGetDirectory(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_StorageManager_GetDirectory
//go:noescape

func StorageManagerGetDirectoryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_StorageBucketDurability
//go:noescape

func ConstOfStorageBucketDurability(str js.Ref) uint32

//go:wasmimport plat/js/web store_IDBDatabaseInfo
//go:noescape

func IDBDatabaseInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IDBDatabaseInfo
//go:noescape

func IDBDatabaseInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBFactory_Open
//go:noescape

func CallIDBFactoryOpen(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	version float64,
) js.Ref

//go:wasmimport plat/js/web func_IDBFactory_Open
//go:noescape

func IDBFactoryOpenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBFactory_Open1
//go:noescape

func CallIDBFactoryOpen1(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_IDBFactory_Open1
//go:noescape

func IDBFactoryOpen1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBFactory_DeleteDatabase
//go:noescape

func CallIDBFactoryDeleteDatabase(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_IDBFactory_DeleteDatabase
//go:noescape

func IDBFactoryDeleteDatabaseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBFactory_Databases
//go:noescape

func CallIDBFactoryDatabases(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_IDBFactory_Databases
//go:noescape

func IDBFactoryDatabasesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBFactory_Cmp
//go:noescape

func CallIDBFactoryCmp(
	this js.Ref,
	ptr unsafe.Pointer,

	first js.Ref,
	second js.Ref,
) int32

//go:wasmimport plat/js/web func_IDBFactory_Cmp
//go:noescape

func IDBFactoryCmpFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_StorageBucket_Name
//go:noescape

func GetStorageBucketName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_StorageBucket_IndexedDB
//go:noescape

func GetStorageBucketIndexedDB(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_StorageBucket_Caches
//go:noescape

func GetStorageBucketCaches(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_StorageBucket_Persist
//go:noescape

func CallStorageBucketPersist(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_StorageBucket_Persist
//go:noescape

func StorageBucketPersistFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageBucket_Persisted
//go:noescape

func CallStorageBucketPersisted(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_StorageBucket_Persisted
//go:noescape

func StorageBucketPersistedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageBucket_Estimate
//go:noescape

func CallStorageBucketEstimate(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_StorageBucket_Estimate
//go:noescape

func StorageBucketEstimateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageBucket_Durability
//go:noescape

func CallStorageBucketDurability(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_StorageBucket_Durability
//go:noescape

func StorageBucketDurabilityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageBucket_SetExpires
//go:noescape

func CallStorageBucketSetExpires(
	this js.Ref,
	ptr unsafe.Pointer,

	expires float64,
) js.Ref

//go:wasmimport plat/js/web func_StorageBucket_SetExpires
//go:noescape

func StorageBucketSetExpiresFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageBucket_Expires
//go:noescape

func CallStorageBucketExpires(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_StorageBucket_Expires
//go:noescape

func StorageBucketExpiresFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageBucket_GetDirectory
//go:noescape

func CallStorageBucketGetDirectory(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_StorageBucket_GetDirectory
//go:noescape

func StorageBucketGetDirectoryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_StorageBucketOptions
//go:noescape

func StorageBucketOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_StorageBucketOptions
//go:noescape

func StorageBucketOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageBucketManager_Open
//go:noescape

func CallStorageBucketManagerOpen(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_StorageBucketManager_Open
//go:noescape

func StorageBucketManagerOpenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageBucketManager_Open1
//go:noescape

func CallStorageBucketManagerOpen1(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_StorageBucketManager_Open1
//go:noescape

func StorageBucketManagerOpen1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageBucketManager_Keys
//go:noescape

func CallStorageBucketManagerKeys(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_StorageBucketManager_Keys
//go:noescape

func StorageBucketManagerKeysFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageBucketManager_Delete
//go:noescape

func CallStorageBucketManagerDelete(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_StorageBucketManager_Delete
//go:noescape

func StorageBucketManagerDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_LockMode
//go:noescape

func ConstOfLockMode(str js.Ref) uint32

//go:wasmimport plat/js/web get_Lock_Name
//go:noescape

func GetLockName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Lock_Mode
//go:noescape

func GetLockMode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web store_LockOptions
//go:noescape

func LockOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_LockOptions
//go:noescape

func LockOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_LockInfo
//go:noescape

func LockInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_LockInfo
//go:noescape

func LockInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_LockManagerSnapshot
//go:noescape

func LockManagerSnapshotJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_LockManagerSnapshot
//go:noescape

func LockManagerSnapshotJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_LockManager_Request
//go:noescape

func CallLockManagerRequest(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	callback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_LockManager_Request
//go:noescape

func LockManagerRequestFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_LockManager_Request1
//go:noescape

func CallLockManagerRequest1(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	options unsafe.Pointer,
	callback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_LockManager_Request1
//go:noescape

func LockManagerRequest1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_LockManager_Query
//go:noescape

func CallLockManagerQuery(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_LockManager_Query
//go:noescape

func LockManagerQueryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_MimeType_Type
//go:noescape

func GetMimeTypeType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MimeType_Description
//go:noescape

func GetMimeTypeDescription(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MimeType_Suffixes
//go:noescape

func GetMimeTypeSuffixes(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MimeType_EnabledPlugin
//go:noescape

func GetMimeTypeEnabledPlugin(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Plugin_Name
//go:noescape

func GetPluginName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Plugin_Description
//go:noescape

func GetPluginDescription(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Plugin_Filename
//go:noescape

func GetPluginFilename(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Plugin_Length
//go:noescape

func GetPluginLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_Plugin_Item
//go:noescape

func CallPluginItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_Plugin_Item
//go:noescape

func PluginItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Plugin_NamedItem
//go:noescape

func CallPluginNamedItem(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Plugin_NamedItem
//go:noescape

func PluginNamedItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_PluginArray_Length
//go:noescape

func GetPluginArrayLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_PluginArray_Refresh
//go:noescape

func CallPluginArrayRefresh(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PluginArray_Refresh
//go:noescape

func PluginArrayRefreshFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PluginArray_Item
//go:noescape

func CallPluginArrayItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_PluginArray_Item
//go:noescape

func PluginArrayItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PluginArray_NamedItem
//go:noescape

func CallPluginArrayNamedItem(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PluginArray_NamedItem
//go:noescape

func PluginArrayNamedItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_MimeTypeArray_Length
//go:noescape

func GetMimeTypeArrayLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_MimeTypeArray_Item
//go:noescape

func CallMimeTypeArrayItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_MimeTypeArray_Item
//go:noescape

func MimeTypeArrayItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MimeTypeArray_NamedItem
//go:noescape

func CallMimeTypeArrayNamedItem(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MimeTypeArray_NamedItem
//go:noescape

func MimeTypeArrayNamedItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLComputeResult
//go:noescape

func MLComputeResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLComputeResult
//go:noescape

func MLComputeResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
