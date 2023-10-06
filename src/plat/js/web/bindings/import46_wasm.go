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

//go:wasmimport plat/js/web new_PresentationConnectionAvailableEvent_PresentationConnectionAvailableEvent
//go:noescape
func NewPresentationConnectionAvailableEventByPresentationConnectionAvailableEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_PresentationConnectionAvailableEvent_Connection
//go:noescape
func GetPresentationConnectionAvailableEventConnection(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_PresentationConnectionCloseReason
//go:noescape
func ConstOfPresentationConnectionCloseReason(str js.Ref) uint32

//go:wasmimport plat/js/web store_PresentationConnectionCloseEventInit
//go:noescape
func PresentationConnectionCloseEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PresentationConnectionCloseEventInit
//go:noescape
func PresentationConnectionCloseEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PresentationConnectionCloseEvent_PresentationConnectionCloseEvent
//go:noescape
func NewPresentationConnectionCloseEventByPresentationConnectionCloseEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_PresentationConnectionCloseEvent_Reason
//go:noescape
func GetPresentationConnectionCloseEventReason(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PresentationConnectionCloseEvent_Message
//go:noescape
func GetPresentationConnectionCloseEventMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_PressureSource
//go:noescape
func ConstOfPressureSource(str js.Ref) uint32

//go:wasmimport plat/js/web constof_PressureState
//go:noescape
func ConstOfPressureState(str js.Ref) uint32

//go:wasmimport plat/js/web get_PressureRecord_Source
//go:noescape
func GetPressureRecordSource(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PressureRecord_State
//go:noescape
func GetPressureRecordState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PressureRecord_Time
//go:noescape
func GetPressureRecordTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PressureRecord_ToJSON
//go:noescape
func HasPressureRecordToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PressureRecord_ToJSON
//go:noescape
func PressureRecordToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PressureRecord_ToJSON
//go:noescape
func CallPressureRecordToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PressureRecord_ToJSON
//go:noescape
func TryPressureRecordToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PressureObserverOptions
//go:noescape
func PressureObserverOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PressureObserverOptions
//go:noescape
func PressureObserverOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PressureObserver_PressureObserver
//go:noescape
func NewPressureObserverByPressureObserver(
	callback js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_PressureObserver_PressureObserver1
//go:noescape
func NewPressureObserverByPressureObserver1(
	callback js.Ref) js.Ref

//go:wasmimport plat/js/web get_PressureObserver_SupportedSources
//go:noescape
func GetPressureObserverSupportedSources(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PressureObserver_Observe
//go:noescape
func HasPressureObserverObserve(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PressureObserver_Observe
//go:noescape
func PressureObserverObserveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PressureObserver_Observe
//go:noescape
func CallPressureObserverObserve(
	this js.Ref, retPtr unsafe.Pointer,
	source uint32)

//go:wasmimport plat/js/web try_PressureObserver_Observe
//go:noescape
func TryPressureObserverObserve(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_PressureObserver_Unobserve
//go:noescape
func HasPressureObserverUnobserve(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PressureObserver_Unobserve
//go:noescape
func PressureObserverUnobserveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PressureObserver_Unobserve
//go:noescape
func CallPressureObserverUnobserve(
	this js.Ref, retPtr unsafe.Pointer,
	source uint32)

//go:wasmimport plat/js/web try_PressureObserver_Unobserve
//go:noescape
func TryPressureObserverUnobserve(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_PressureObserver_Disconnect
//go:noescape
func HasPressureObserverDisconnect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PressureObserver_Disconnect
//go:noescape
func PressureObserverDisconnectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PressureObserver_Disconnect
//go:noescape
func CallPressureObserverDisconnect(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PressureObserver_Disconnect
//go:noescape
func TryPressureObserverDisconnect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PressureObserver_TakeRecords
//go:noescape
func HasPressureObserverTakeRecords(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PressureObserver_TakeRecords
//go:noescape
func PressureObserverTakeRecordsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PressureObserver_TakeRecords
//go:noescape
func CallPressureObserverTakeRecords(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PressureObserver_TakeRecords
//go:noescape
func TryPressureObserverTakeRecords(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PrivateNetworkAccessPermissionDescriptor
//go:noescape
func PrivateNetworkAccessPermissionDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PrivateNetworkAccessPermissionDescriptor
//go:noescape
func PrivateNetworkAccessPermissionDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ProfilerInitOptions
//go:noescape
func ProfilerInitOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ProfilerInitOptions
//go:noescape
func ProfilerInitOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ProfilerFrame
//go:noescape
func ProfilerFrameJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ProfilerFrame
//go:noescape
func ProfilerFrameJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ProfilerStack
//go:noescape
func ProfilerStackJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ProfilerStack
//go:noescape
func ProfilerStackJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ProfilerSample
//go:noescape
func ProfilerSampleJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ProfilerSample
//go:noescape
func ProfilerSampleJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ProfilerTrace
//go:noescape
func ProfilerTraceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ProfilerTrace
//go:noescape
func ProfilerTraceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_Profiler_Profiler
//go:noescape
func NewProfilerByProfiler(
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_Profiler_SampleInterval
//go:noescape
func GetProfilerSampleInterval(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Profiler_Stopped
//go:noescape
func GetProfilerStopped(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Profiler_Stop
//go:noescape
func HasProfilerStop(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Profiler_Stop
//go:noescape
func ProfilerStopFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Profiler_Stop
//go:noescape
func CallProfilerStop(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Profiler_Stop
//go:noescape
func TryProfilerStop(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ProgressEventInit
//go:noescape
func ProgressEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ProgressEventInit
//go:noescape
func ProgressEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ProgressEvent_ProgressEvent
//go:noescape
func NewProgressEventByProgressEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ProgressEvent_ProgressEvent1
//go:noescape
func NewProgressEventByProgressEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_ProgressEvent_LengthComputable
//go:noescape
func GetProgressEventLengthComputable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ProgressEvent_Loaded
//go:noescape
func GetProgressEventLoaded(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ProgressEvent_Total
//go:noescape
func GetProgressEventTotal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PromiseRejectionEventInit
//go:noescape
func PromiseRejectionEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PromiseRejectionEventInit
//go:noescape
func PromiseRejectionEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PromiseRejectionEvent_PromiseRejectionEvent
//go:noescape
func NewPromiseRejectionEventByPromiseRejectionEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_PromiseRejectionEvent_Promise
//go:noescape
func GetPromiseRejectionEventPromise(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PromiseRejectionEvent_Reason
//go:noescape
func GetPromiseRejectionEventReason(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ProximityReadingValues
//go:noescape
func ProximityReadingValuesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ProximityReadingValues
//go:noescape
func ProximityReadingValuesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ProximitySensor_ProximitySensor
//go:noescape
func NewProximitySensorByProximitySensor(
	sensorOptions unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ProximitySensor_ProximitySensor1
//go:noescape
func NewProximitySensorByProximitySensor1() js.Ref

//go:wasmimport plat/js/web get_ProximitySensor_Distance
//go:noescape
func GetProximitySensorDistance(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ProximitySensor_Max
//go:noescape
func GetProximitySensorMax(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ProximitySensor_Near
//go:noescape
func GetProximitySensorNear(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_RegistrationResponseJSON
//go:noescape
func RegistrationResponseJSONJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RegistrationResponseJSON
//go:noescape
func RegistrationResponseJSONJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PublicKeyCredentialDescriptorJSON
//go:noescape
func PublicKeyCredentialDescriptorJSONJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PublicKeyCredentialDescriptorJSON
//go:noescape
func PublicKeyCredentialDescriptorJSONJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PublicKeyCredentialRequestOptionsJSON
//go:noescape
func PublicKeyCredentialRequestOptionsJSONJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PublicKeyCredentialRequestOptionsJSON
//go:noescape
func PublicKeyCredentialRequestOptionsJSONJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PublicKeyCredentialUserEntityJSON
//go:noescape
func PublicKeyCredentialUserEntityJSONJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PublicKeyCredentialUserEntityJSON
//go:noescape
func PublicKeyCredentialUserEntityJSONJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PublicKeyCredentialCreationOptionsJSON
//go:noescape
func PublicKeyCredentialCreationOptionsJSONJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PublicKeyCredentialCreationOptionsJSON
//go:noescape
func PublicKeyCredentialCreationOptionsJSONJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_PublicKeyCredential_RawId
//go:noescape
func GetPublicKeyCredentialRawId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PublicKeyCredential_Response
//go:noescape
func GetPublicKeyCredentialResponse(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PublicKeyCredential_AuthenticatorAttachment
//go:noescape
func GetPublicKeyCredentialAuthenticatorAttachment(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PublicKeyCredential_GetClientExtensionResults
//go:noescape
func HasPublicKeyCredentialGetClientExtensionResults(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PublicKeyCredential_GetClientExtensionResults
//go:noescape
func PublicKeyCredentialGetClientExtensionResultsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PublicKeyCredential_GetClientExtensionResults
//go:noescape
func CallPublicKeyCredentialGetClientExtensionResults(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PublicKeyCredential_GetClientExtensionResults
//go:noescape
func TryPublicKeyCredentialGetClientExtensionResults(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PublicKeyCredential_IsConditionalMediationAvailable
//go:noescape
func HasPublicKeyCredentialIsConditionalMediationAvailable(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PublicKeyCredential_IsConditionalMediationAvailable
//go:noescape
func PublicKeyCredentialIsConditionalMediationAvailableFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PublicKeyCredential_IsConditionalMediationAvailable
//go:noescape
func CallPublicKeyCredentialIsConditionalMediationAvailable(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PublicKeyCredential_IsConditionalMediationAvailable
//go:noescape
func TryPublicKeyCredentialIsConditionalMediationAvailable(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PublicKeyCredential_ToJSON
//go:noescape
func HasPublicKeyCredentialToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PublicKeyCredential_ToJSON
//go:noescape
func PublicKeyCredentialToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PublicKeyCredential_ToJSON
//go:noescape
func CallPublicKeyCredentialToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PublicKeyCredential_ToJSON
//go:noescape
func TryPublicKeyCredentialToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PublicKeyCredential_ParseRequestOptionsFromJSON
//go:noescape
func HasPublicKeyCredentialParseRequestOptionsFromJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PublicKeyCredential_ParseRequestOptionsFromJSON
//go:noescape
func PublicKeyCredentialParseRequestOptionsFromJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PublicKeyCredential_ParseRequestOptionsFromJSON
//go:noescape
func CallPublicKeyCredentialParseRequestOptionsFromJSON(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_PublicKeyCredential_ParseRequestOptionsFromJSON
//go:noescape
func TryPublicKeyCredentialParseRequestOptionsFromJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PublicKeyCredential_IsPasskeyPlatformAuthenticatorAvailable
//go:noescape
func HasPublicKeyCredentialIsPasskeyPlatformAuthenticatorAvailable(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PublicKeyCredential_IsPasskeyPlatformAuthenticatorAvailable
//go:noescape
func PublicKeyCredentialIsPasskeyPlatformAuthenticatorAvailableFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PublicKeyCredential_IsPasskeyPlatformAuthenticatorAvailable
//go:noescape
func CallPublicKeyCredentialIsPasskeyPlatformAuthenticatorAvailable(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PublicKeyCredential_IsPasskeyPlatformAuthenticatorAvailable
//go:noescape
func TryPublicKeyCredentialIsPasskeyPlatformAuthenticatorAvailable(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PublicKeyCredential_ParseCreationOptionsFromJSON
//go:noescape
func HasPublicKeyCredentialParseCreationOptionsFromJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PublicKeyCredential_ParseCreationOptionsFromJSON
//go:noescape
func PublicKeyCredentialParseCreationOptionsFromJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PublicKeyCredential_ParseCreationOptionsFromJSON
//go:noescape
func CallPublicKeyCredentialParseCreationOptionsFromJSON(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_PublicKeyCredential_ParseCreationOptionsFromJSON
//go:noescape
func TryPublicKeyCredentialParseCreationOptionsFromJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PublicKeyCredential_IsUserVerifyingPlatformAuthenticatorAvailable
//go:noescape
func HasPublicKeyCredentialIsUserVerifyingPlatformAuthenticatorAvailable(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PublicKeyCredential_IsUserVerifyingPlatformAuthenticatorAvailable
//go:noescape
func PublicKeyCredentialIsUserVerifyingPlatformAuthenticatorAvailableFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PublicKeyCredential_IsUserVerifyingPlatformAuthenticatorAvailable
//go:noescape
func CallPublicKeyCredentialIsUserVerifyingPlatformAuthenticatorAvailable(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PublicKeyCredential_IsUserVerifyingPlatformAuthenticatorAvailable
//go:noescape
func TryPublicKeyCredentialIsUserVerifyingPlatformAuthenticatorAvailable(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PublicKeyCredentialEntity
//go:noescape
func PublicKeyCredentialEntityJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PublicKeyCredentialEntity
//go:noescape
func PublicKeyCredentialEntityJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_PublicKeyCredentialHints
//go:noescape
func ConstOfPublicKeyCredentialHints(str js.Ref) uint32

//go:wasmimport plat/js/web constof_PublicKeyCredentialType
//go:noescape
func ConstOfPublicKeyCredentialType(str js.Ref) uint32

//go:wasmimport plat/js/web store_PushEventInit
//go:noescape
func PushEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PushEventInit
//go:noescape
func PushEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_PushMessageData_ArrayBuffer
//go:noescape
func HasPushMessageDataArrayBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PushMessageData_ArrayBuffer
//go:noescape
func PushMessageDataArrayBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushMessageData_ArrayBuffer
//go:noescape
func CallPushMessageDataArrayBuffer(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PushMessageData_ArrayBuffer
//go:noescape
func TryPushMessageDataArrayBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PushMessageData_Blob
//go:noescape
func HasPushMessageDataBlob(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PushMessageData_Blob
//go:noescape
func PushMessageDataBlobFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushMessageData_Blob
//go:noescape
func CallPushMessageDataBlob(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PushMessageData_Blob
//go:noescape
func TryPushMessageDataBlob(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PushMessageData_Json
//go:noescape
func HasPushMessageDataJson(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PushMessageData_Json
//go:noescape
func PushMessageDataJsonFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushMessageData_Json
//go:noescape
func CallPushMessageDataJson(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PushMessageData_Json
//go:noescape
func TryPushMessageDataJson(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PushMessageData_Text
//go:noescape
func HasPushMessageDataText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PushMessageData_Text
//go:noescape
func PushMessageDataTextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushMessageData_Text
//go:noescape
func CallPushMessageDataText(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PushMessageData_Text
//go:noescape
func TryPushMessageDataText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_PushEvent_PushEvent
//go:noescape
func NewPushEventByPushEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_PushEvent_PushEvent1
//go:noescape
func NewPushEventByPushEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_PushEvent_Data
//go:noescape
func GetPushEventData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PushPermissionDescriptor
//go:noescape
func PushPermissionDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PushPermissionDescriptor
//go:noescape
func PushPermissionDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PushSubscriptionChangeEventInit
//go:noescape
func PushSubscriptionChangeEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PushSubscriptionChangeEventInit
//go:noescape
func PushSubscriptionChangeEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PushSubscriptionChangeEvent_PushSubscriptionChangeEvent
//go:noescape
func NewPushSubscriptionChangeEventByPushSubscriptionChangeEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_PushSubscriptionChangeEvent_PushSubscriptionChangeEvent1
//go:noescape
func NewPushSubscriptionChangeEventByPushSubscriptionChangeEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_PushSubscriptionChangeEvent_NewSubscription
//go:noescape
func GetPushSubscriptionChangeEventNewSubscription(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PushSubscriptionChangeEvent_OldSubscription
//go:noescape
func GetPushSubscriptionChangeEventOldSubscription(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_RTCAnswerOptions
//go:noescape
func RTCAnswerOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCAnswerOptions
//go:noescape
func RTCAnswerOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RTCStatsType
//go:noescape
func ConstOfRTCStatsType(str js.Ref) uint32
