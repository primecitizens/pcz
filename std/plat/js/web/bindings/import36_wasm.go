// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

func _() {
	var (
		_ js.Void
		_ unsafe.Pointer
	)
}

//go:wasmimport plat/js/web new_GPUInternalError_GPUInternalError
//go:noescape
func NewGPUInternalErrorByGPUInternalError(
	message js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUObjectDescriptorBase
//go:noescape
func GPUObjectDescriptorBaseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUObjectDescriptorBase
//go:noescape
func GPUObjectDescriptorBaseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_GPUOutOfMemoryError_GPUOutOfMemoryError
//go:noescape
func NewGPUOutOfMemoryErrorByGPUOutOfMemoryError(
	message js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUPipelineDescriptorBase
//go:noescape
func GPUPipelineDescriptorBaseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUPipelineDescriptorBase
//go:noescape
func GPUPipelineDescriptorBaseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GPUPipelineErrorReason
//go:noescape
func ConstOfGPUPipelineErrorReason(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPUPipelineErrorInit
//go:noescape
func GPUPipelineErrorInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUPipelineErrorInit
//go:noescape
func GPUPipelineErrorInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_GPUPipelineError_GPUPipelineError
//go:noescape
func NewGPUPipelineErrorByGPUPipelineError(
	message js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_GPUPipelineError_GPUPipelineError1
//go:noescape
func NewGPUPipelineErrorByGPUPipelineError1() js.Ref

//go:wasmimport plat/js/web get_GPUPipelineError_Reason
//go:noescape
func GetGPUPipelineErrorReason(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_GPURenderPassLayout
//go:noescape
func GPURenderPassLayoutJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPURenderPassLayout
//go:noescape
func GPURenderPassLayoutJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUUncapturedErrorEventInit
//go:noescape
func GPUUncapturedErrorEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUUncapturedErrorEventInit
//go:noescape
func GPUUncapturedErrorEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_GPUUncapturedErrorEvent_GPUUncapturedErrorEvent
//go:noescape
func NewGPUUncapturedErrorEventByGPUUncapturedErrorEvent(
	typ js.Ref,
	gpuUncapturedErrorEventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_GPUUncapturedErrorEvent_Error
//go:noescape
func GetGPUUncapturedErrorEventError(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_GPUValidationError_GPUValidationError
//go:noescape
func NewGPUValidationErrorByGPUValidationError(
	message js.Ref) js.Ref

//go:wasmimport plat/js/web store_GamepadEventInit
//go:noescape
func GamepadEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GamepadEventInit
//go:noescape
func GamepadEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_GamepadEvent_GamepadEvent
//go:noescape
func NewGamepadEventByGamepadEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_GamepadEvent_Gamepad
//go:noescape
func GetGamepadEventGamepad(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_RTCIdentityProviderDetails
//go:noescape
func RTCIdentityProviderDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCIdentityProviderDetails
//go:noescape
func RTCIdentityProviderDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCIdentityAssertionResult
//go:noescape
func RTCIdentityAssertionResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCIdentityAssertionResult
//go:noescape
func RTCIdentityAssertionResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCIdentityProviderOptions
//go:noescape
func RTCIdentityProviderOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCIdentityProviderOptions
//go:noescape
func RTCIdentityProviderOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GenerateBidInterestGroup
//go:noescape
func GenerateBidInterestGroupJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GenerateBidInterestGroup
//go:noescape
func GenerateBidInterestGroupJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GenerateBidOutput
//go:noescape
func GenerateBidOutputJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GenerateBidOutput
//go:noescape
func GenerateBidOutputJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GenerateTestReportParameters
//go:noescape
func GenerateTestReportParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GenerateTestReportParameters
//go:noescape
func GenerateTestReportParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GeolocationReadingValues
//go:noescape
func GeolocationReadingValuesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GeolocationReadingValues
//go:noescape
func GeolocationReadingValuesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GeolocationSensorOptions
//go:noescape
func GeolocationSensorOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GeolocationSensorOptions
//go:noescape
func GeolocationSensorOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GeolocationSensorReading
//go:noescape
func GeolocationSensorReadingJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GeolocationSensorReading
//go:noescape
func GeolocationSensorReadingJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ReadOptions
//go:noescape
func ReadOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ReadOptions
//go:noescape
func ReadOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_GeolocationSensor_GeolocationSensor
//go:noescape
func NewGeolocationSensorByGeolocationSensor(
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_GeolocationSensor_GeolocationSensor1
//go:noescape
func NewGeolocationSensorByGeolocationSensor1() js.Ref

//go:wasmimport plat/js/web get_GeolocationSensor_Latitude
//go:noescape
func GetGeolocationSensorLatitude(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationSensor_Longitude
//go:noescape
func GetGeolocationSensorLongitude(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationSensor_Altitude
//go:noescape
func GetGeolocationSensorAltitude(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationSensor_Accuracy
//go:noescape
func GetGeolocationSensorAccuracy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationSensor_AltitudeAccuracy
//go:noescape
func GetGeolocationSensorAltitudeAccuracy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationSensor_Heading
//go:noescape
func GetGeolocationSensorHeading(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationSensor_Speed
//go:noescape
func GetGeolocationSensorSpeed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GeolocationSensor_Read
//go:noescape
func HasGeolocationSensorRead(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GeolocationSensor_Read
//go:noescape
func GeolocationSensorReadFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GeolocationSensor_Read
//go:noescape
func CallGeolocationSensorRead(
	this js.Ref, retPtr unsafe.Pointer,
	readOptions unsafe.Pointer)

//go:wasmimport plat/js/web try_GeolocationSensor_Read
//go:noescape
func TryGeolocationSensorRead(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	readOptions unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GeolocationSensor_Read1
//go:noescape
func HasGeolocationSensorRead1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GeolocationSensor_Read1
//go:noescape
func GeolocationSensorRead1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GeolocationSensor_Read1
//go:noescape
func CallGeolocationSensorRead1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GeolocationSensor_Read1
//go:noescape
func TryGeolocationSensorRead1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_ValueType
//go:noescape
func ConstOfValueType(str js.Ref) uint32

//go:wasmimport plat/js/web store_GlobalDescriptor
//go:noescape
func GlobalDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GlobalDescriptor
//go:noescape
func GlobalDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_Global_Global
//go:noescape
func NewGlobalByGlobal(
	descriptor unsafe.Pointer,
	v js.Ref) js.Ref

//go:wasmimport plat/js/web new_Global_Global1
//go:noescape
func NewGlobalByGlobal1(
	descriptor unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_Global_Value
//go:noescape
func GetGlobalValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Global_Value
//go:noescape
func SetGlobalValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_Global_ValueOf
//go:noescape
func HasGlobalValueOf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Global_ValueOf
//go:noescape
func GlobalValueOfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Global_ValueOf
//go:noescape
func CallGlobalValueOf(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Global_ValueOf
//go:noescape
func TryGlobalValueOf(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_GravityReadingValues
//go:noescape
func GravityReadingValuesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GravityReadingValues
//go:noescape
func GravityReadingValuesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_GravitySensor_GravitySensor
//go:noescape
func NewGravitySensorByGravitySensor(
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_GravitySensor_GravitySensor1
//go:noescape
func NewGravitySensorByGravitySensor1() js.Ref

//go:wasmimport plat/js/web constof_GyroscopeLocalCoordinateSystem
//go:noescape
func ConstOfGyroscopeLocalCoordinateSystem(str js.Ref) uint32

//go:wasmimport plat/js/web store_GyroscopeSensorOptions
//go:noescape
func GyroscopeSensorOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GyroscopeSensorOptions
//go:noescape
func GyroscopeSensorOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_Gyroscope_Gyroscope
//go:noescape
func NewGyroscopeByGyroscope(
	sensorOptions unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_Gyroscope_Gyroscope1
//go:noescape
func NewGyroscopeByGyroscope1() js.Ref

//go:wasmimport plat/js/web get_Gyroscope_X
//go:noescape
func GetGyroscopeX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Gyroscope_Y
//go:noescape
func GetGyroscopeY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Gyroscope_Z
//go:noescape
func GetGyroscopeZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_GyroscopeReadingValues
//go:noescape
func GyroscopeReadingValuesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GyroscopeReadingValues
//go:noescape
func GyroscopeReadingValuesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_HIDConnectionEventInit
//go:noescape
func HIDConnectionEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_HIDConnectionEventInit
//go:noescape
func HIDConnectionEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_HIDConnectionEvent_HIDConnectionEvent
//go:noescape
func NewHIDConnectionEventByHIDConnectionEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_HIDConnectionEvent_Device
//go:noescape
func GetHIDConnectionEventDevice(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_HIDInputReportEventInit
//go:noescape
func HIDInputReportEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_HIDInputReportEventInit
//go:noescape
func HIDInputReportEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_HIDInputReportEvent_HIDInputReportEvent
//go:noescape
func NewHIDInputReportEventByHIDInputReportEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_HIDInputReportEvent_Device
//go:noescape
func GetHIDInputReportEventDevice(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HIDInputReportEvent_ReportId
//go:noescape
func GetHIDInputReportEventReportId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HIDInputReportEvent_Data
//go:noescape
func GetHIDInputReportEventData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLAnchorElement_HTMLAnchorElement
//go:noescape
func NewHTMLAnchorElementByHTMLAnchorElement() js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Target
//go:noescape
func GetHTMLAnchorElementTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Target
//go:noescape
func SetHTMLAnchorElementTarget(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Download
//go:noescape
func GetHTMLAnchorElementDownload(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Download
//go:noescape
func SetHTMLAnchorElementDownload(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Ping
//go:noescape
func GetHTMLAnchorElementPing(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Ping
//go:noescape
func SetHTMLAnchorElementPing(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Rel
//go:noescape
func GetHTMLAnchorElementRel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Rel
//go:noescape
func SetHTMLAnchorElementRel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_RelList
//go:noescape
func GetHTMLAnchorElementRelList(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLAnchorElement_Hreflang
//go:noescape
func GetHTMLAnchorElementHreflang(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Hreflang
//go:noescape
func SetHTMLAnchorElementHreflang(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Type
//go:noescape
func GetHTMLAnchorElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Type
//go:noescape
func SetHTMLAnchorElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Text
//go:noescape
func GetHTMLAnchorElementText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Text
//go:noescape
func SetHTMLAnchorElementText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_ReferrerPolicy
//go:noescape
func GetHTMLAnchorElementReferrerPolicy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_ReferrerPolicy
//go:noescape
func SetHTMLAnchorElementReferrerPolicy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_AttributionSourceId
//go:noescape
func GetHTMLAnchorElementAttributionSourceId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_AttributionSourceId
//go:noescape
func SetHTMLAnchorElementAttributionSourceId(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Coords
//go:noescape
func GetHTMLAnchorElementCoords(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Coords
//go:noescape
func SetHTMLAnchorElementCoords(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Charset
//go:noescape
func GetHTMLAnchorElementCharset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Charset
//go:noescape
func SetHTMLAnchorElementCharset(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Name
//go:noescape
func GetHTMLAnchorElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Name
//go:noescape
func SetHTMLAnchorElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Rev
//go:noescape
func GetHTMLAnchorElementRev(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Rev
//go:noescape
func SetHTMLAnchorElementRev(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Shape
//go:noescape
func GetHTMLAnchorElementShape(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Shape
//go:noescape
func SetHTMLAnchorElementShape(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Href
//go:noescape
func GetHTMLAnchorElementHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Href
//go:noescape
func SetHTMLAnchorElementHref(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Origin
//go:noescape
func GetHTMLAnchorElementOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLAnchorElement_Protocol
//go:noescape
func GetHTMLAnchorElementProtocol(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Protocol
//go:noescape
func SetHTMLAnchorElementProtocol(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Username
//go:noescape
func GetHTMLAnchorElementUsername(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Username
//go:noescape
func SetHTMLAnchorElementUsername(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Password
//go:noescape
func GetHTMLAnchorElementPassword(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Password
//go:noescape
func SetHTMLAnchorElementPassword(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Host
//go:noescape
func GetHTMLAnchorElementHost(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Host
//go:noescape
func SetHTMLAnchorElementHost(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Hostname
//go:noescape
func GetHTMLAnchorElementHostname(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Hostname
//go:noescape
func SetHTMLAnchorElementHostname(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Port
//go:noescape
func GetHTMLAnchorElementPort(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Port
//go:noescape
func SetHTMLAnchorElementPort(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Pathname
//go:noescape
func GetHTMLAnchorElementPathname(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Pathname
//go:noescape
func SetHTMLAnchorElementPathname(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Search
//go:noescape
func GetHTMLAnchorElementSearch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Search
//go:noescape
func SetHTMLAnchorElementSearch(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_Hash
//go:noescape
func GetHTMLAnchorElementHash(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_Hash
//go:noescape
func SetHTMLAnchorElementHash(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAnchorElement_AttributionSrc
//go:noescape
func GetHTMLAnchorElementAttributionSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAnchorElement_AttributionSrc
//go:noescape
func SetHTMLAnchorElementAttributionSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLAreaElement_HTMLAreaElement
//go:noescape
func NewHTMLAreaElementByHTMLAreaElement() js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Alt
//go:noescape
func GetHTMLAreaElementAlt(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Alt
//go:noescape
func SetHTMLAreaElementAlt(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Coords
//go:noescape
func GetHTMLAreaElementCoords(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Coords
//go:noescape
func SetHTMLAreaElementCoords(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Shape
//go:noescape
func GetHTMLAreaElementShape(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Shape
//go:noescape
func SetHTMLAreaElementShape(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Target
//go:noescape
func GetHTMLAreaElementTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Target
//go:noescape
func SetHTMLAreaElementTarget(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Download
//go:noescape
func GetHTMLAreaElementDownload(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Download
//go:noescape
func SetHTMLAreaElementDownload(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Ping
//go:noescape
func GetHTMLAreaElementPing(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Ping
//go:noescape
func SetHTMLAreaElementPing(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Rel
//go:noescape
func GetHTMLAreaElementRel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Rel
//go:noescape
func SetHTMLAreaElementRel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_RelList
//go:noescape
func GetHTMLAreaElementRelList(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLAreaElement_ReferrerPolicy
//go:noescape
func GetHTMLAreaElementReferrerPolicy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_ReferrerPolicy
//go:noescape
func SetHTMLAreaElementReferrerPolicy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_NoHref
//go:noescape
func GetHTMLAreaElementNoHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_NoHref
//go:noescape
func SetHTMLAreaElementNoHref(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Href
//go:noescape
func GetHTMLAreaElementHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Href
//go:noescape
func SetHTMLAreaElementHref(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Origin
//go:noescape
func GetHTMLAreaElementOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLAreaElement_Protocol
//go:noescape
func GetHTMLAreaElementProtocol(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Protocol
//go:noescape
func SetHTMLAreaElementProtocol(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Username
//go:noescape
func GetHTMLAreaElementUsername(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Username
//go:noescape
func SetHTMLAreaElementUsername(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Password
//go:noescape
func GetHTMLAreaElementPassword(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Password
//go:noescape
func SetHTMLAreaElementPassword(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Host
//go:noescape
func GetHTMLAreaElementHost(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Host
//go:noescape
func SetHTMLAreaElementHost(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Hostname
//go:noescape
func GetHTMLAreaElementHostname(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Hostname
//go:noescape
func SetHTMLAreaElementHostname(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Port
//go:noescape
func GetHTMLAreaElementPort(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Port
//go:noescape
func SetHTMLAreaElementPort(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Pathname
//go:noescape
func GetHTMLAreaElementPathname(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Pathname
//go:noescape
func SetHTMLAreaElementPathname(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Search
//go:noescape
func GetHTMLAreaElementSearch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Search
//go:noescape
func SetHTMLAreaElementSearch(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLAreaElement_Hash
//go:noescape
func GetHTMLAreaElementHash(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLAreaElement_Hash
//go:noescape
func SetHTMLAreaElementHash(
	this js.Ref,
	val js.Ref,
) js.Ref
