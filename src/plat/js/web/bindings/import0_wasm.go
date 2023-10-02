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

//go:wasmimport plat/js/web call_ANGLE_instanced_arrays_DrawArraysInstancedANGLE
//go:noescape

func CallANGLE_instanced_arraysDrawArraysInstancedANGLE(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	first int32,
	count int32,
	primcount int32,
) js.Ref

//go:wasmimport plat/js/web func_ANGLE_instanced_arrays_DrawArraysInstancedANGLE
//go:noescape

func ANGLE_instanced_arraysDrawArraysInstancedANGLEFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ANGLE_instanced_arrays_DrawElementsInstancedANGLE
//go:noescape

func CallANGLE_instanced_arraysDrawElementsInstancedANGLE(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	count int32,
	typ uint32,
	offset float64,
	primcount int32,
) js.Ref

//go:wasmimport plat/js/web func_ANGLE_instanced_arrays_DrawElementsInstancedANGLE
//go:noescape

func ANGLE_instanced_arraysDrawElementsInstancedANGLEFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ANGLE_instanced_arrays_VertexAttribDivisorANGLE
//go:noescape

func CallANGLE_instanced_arraysVertexAttribDivisorANGLE(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	divisor uint32,
) js.Ref

//go:wasmimport plat/js/web func_ANGLE_instanced_arrays_VertexAttribDivisorANGLE
//go:noescape

func ANGLE_instanced_arraysVertexAttribDivisorANGLEFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_AV1EncoderConfig
//go:noescape

func AV1EncoderConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AV1EncoderConfig
//go:noescape

func AV1EncoderConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_AacBitstreamFormat
//go:noescape

func ConstOfAacBitstreamFormat(str js.Ref) uint32

//go:wasmimport plat/js/web store_AacEncoderConfig
//go:noescape

func AacEncoderConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AacEncoderConfig
//go:noescape

func AacEncoderConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_EventInit
//go:noescape

func EventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EventInit
//go:noescape

func EventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AddEventListenerOptions
//go:noescape

func AddEventListenerOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AddEventListenerOptions
//go:noescape

func AddEventListenerOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_EventListenerOptions
//go:noescape

func EventListenerOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EventListenerOptions
//go:noescape

func EventListenerOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_EventTarget_AddEventListener
//go:noescape

func CallEventTargetAddEventListener(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	callback js.Ref,
	options js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_EventTarget_AddEventListener
//go:noescape

func EventTargetAddEventListenerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_EventTarget_AddEventListener1
//go:noescape

func CallEventTargetAddEventListener1(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	callback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_EventTarget_AddEventListener1
//go:noescape

func EventTargetAddEventListener1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_EventTarget_RemoveEventListener
//go:noescape

func CallEventTargetRemoveEventListener(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	callback js.Ref,
	options js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_EventTarget_RemoveEventListener
//go:noescape

func EventTargetRemoveEventListenerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_EventTarget_RemoveEventListener1
//go:noescape

func CallEventTargetRemoveEventListener1(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	callback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_EventTarget_RemoveEventListener1
//go:noescape

func EventTargetRemoveEventListener1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_EventTarget_DispatchEvent
//go:noescape

func CallEventTargetDispatchEvent(
	this js.Ref,
	ptr unsafe.Pointer,

	event js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_EventTarget_DispatchEvent
//go:noescape

func EventTargetDispatchEventFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_Event_Event
//go:noescape

func NewEventByEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_Event_Event1
//go:noescape

func NewEventByEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_Event_Type
//go:noescape

func GetEventType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Event_Target
//go:noescape

func GetEventTarget(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Event_SrcElement
//go:noescape

func GetEventSrcElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Event_CurrentTarget
//go:noescape

func GetEventCurrentTarget(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Event_EventPhase
//go:noescape

func GetEventEventPhase(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_Event_CancelBubble
//go:noescape

func GetEventCancelBubble(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Event_CancelBubble
//go:noescape

func SetEventCancelBubble(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Event_Bubbles
//go:noescape

func GetEventBubbles(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Event_Cancelable
//go:noescape

func GetEventCancelable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Event_ReturnValue
//go:noescape

func GetEventReturnValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Event_ReturnValue
//go:noescape

func SetEventReturnValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Event_DefaultPrevented
//go:noescape

func GetEventDefaultPrevented(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Event_Composed
//go:noescape

func GetEventComposed(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Event_IsTrusted
//go:noescape

func GetEventIsTrusted(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Event_TimeStamp
//go:noescape

func GetEventTimeStamp(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web call_Event_ComposedPath
//go:noescape

func CallEventComposedPath(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Event_ComposedPath
//go:noescape

func EventComposedPathFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Event_StopPropagation
//go:noescape

func CallEventStopPropagation(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Event_StopPropagation
//go:noescape

func EventStopPropagationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Event_StopImmediatePropagation
//go:noescape

func CallEventStopImmediatePropagation(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Event_StopImmediatePropagation
//go:noescape

func EventStopImmediatePropagationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Event_PreventDefault
//go:noescape

func CallEventPreventDefault(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Event_PreventDefault
//go:noescape

func EventPreventDefaultFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Event_InitEvent
//go:noescape

func CallEventInitEvent(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Event_InitEvent
//go:noescape

func EventInitEventFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Event_InitEvent1
//go:noescape

func CallEventInitEvent1(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Event_InitEvent1
//go:noescape

func EventInitEvent1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Event_InitEvent2
//go:noescape

func CallEventInitEvent2(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Event_InitEvent2
//go:noescape

func EventInitEvent2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_AbortSignal_Aborted
//go:noescape

func GetAbortSignalAborted(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_AbortSignal_Reason
//go:noescape

func GetAbortSignalReason(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_AbortSignal_Abort
//go:noescape

func CallAbortSignalAbort(
	this js.Ref,
	ptr unsafe.Pointer,

	reason js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_AbortSignal_Abort
//go:noescape

func AbortSignalAbortFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AbortSignal_Abort1
//go:noescape

func CallAbortSignalAbort1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AbortSignal_Abort1
//go:noescape

func AbortSignalAbort1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AbortSignal_Timeout
//go:noescape

func CallAbortSignalTimeout(
	this js.Ref,
	ptr unsafe.Pointer,

	milliseconds float64,
) js.Ref

//go:wasmimport plat/js/web func_AbortSignal_Timeout
//go:noescape

func AbortSignalTimeoutFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AbortSignal_Any
//go:noescape

func CallAbortSignalAny(
	this js.Ref,
	ptr unsafe.Pointer,

	signals js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_AbortSignal_Any
//go:noescape

func AbortSignalAnyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AbortSignal_ThrowIfAborted
//go:noescape

func CallAbortSignalThrowIfAborted(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AbortSignal_ThrowIfAborted
//go:noescape

func AbortSignalThrowIfAbortedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_AbortController_Signal
//go:noescape

func GetAbortControllerSignal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_AbortController_Abort
//go:noescape

func CallAbortControllerAbort(
	this js.Ref,
	ptr unsafe.Pointer,

	reason js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_AbortController_Abort
//go:noescape

func AbortControllerAbortFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AbortController_Abort1
//go:noescape

func CallAbortControllerAbort1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AbortController_Abort1
//go:noescape

func AbortControllerAbort1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_AbsoluteOrientationReadingValues
//go:noescape

func AbsoluteOrientationReadingValuesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AbsoluteOrientationReadingValues
//go:noescape

func AbsoluteOrientationReadingValuesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_OrientationSensorLocalCoordinateSystem
//go:noescape

func ConstOfOrientationSensorLocalCoordinateSystem(str js.Ref) uint32

//go:wasmimport plat/js/web store_OrientationSensorOptions
//go:noescape

func OrientationSensorOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_OrientationSensorOptions
//go:noescape

func OrientationSensorOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_AbsoluteOrientationSensor_AbsoluteOrientationSensor
//go:noescape

func NewAbsoluteOrientationSensorByAbsoluteOrientationSensor(
	sensorOptions unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_AbsoluteOrientationSensor_AbsoluteOrientationSensor1
//go:noescape

func NewAbsoluteOrientationSensorByAbsoluteOrientationSensor1() js.Ref

//go:wasmimport plat/js/web store_GetRootNodeOptions
//go:noescape

func GetRootNodeOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GetRootNodeOptions
//go:noescape

func GetRootNodeOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_Attr_NamespaceURI
//go:noescape

func GetAttrNamespaceURI(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Attr_Prefix
//go:noescape

func GetAttrPrefix(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Attr_LocalName
//go:noescape

func GetAttrLocalName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Attr_Name
//go:noescape

func GetAttrName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Attr_Value
//go:noescape

func GetAttrValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Attr_Value
//go:noescape

func SetAttrValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Attr_OwnerElement
//go:noescape

func GetAttrOwnerElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Attr_Specified
//go:noescape

func GetAttrSpecified(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web constof_FillMode
//go:noescape

func ConstOfFillMode(str js.Ref) uint32

//go:wasmimport plat/js/web constof_PlaybackDirection
//go:noescape

func ConstOfPlaybackDirection(str js.Ref) uint32

//go:wasmimport plat/js/web new_CSSUnitValue_CSSUnitValue
//go:noescape

func NewCSSUnitValueByCSSUnitValue(
	value float64,
	unit js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSUnitValue_Value
//go:noescape

func GetCSSUnitValueValue(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_CSSUnitValue_Value
//go:noescape

func SetCSSUnitValueValue(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_CSSUnitValue_Unit
//go:noescape

func GetCSSUnitValueUnit(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSSNumericArray_Length
//go:noescape

func GetCSSNumericArrayLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_CSSNumericArray_Get
//go:noescape

func CallCSSNumericArrayGet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_CSSNumericArray_Get
//go:noescape

func CSSNumericArrayGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSMathSum_CSSMathSum
//go:noescape

func NewCSSMathSumByCSSMathSum(
	argsPtr unsafe.Pointer,
	argsCount uint32) js.Ref

//go:wasmimport plat/js/web get_CSSMathSum_Values
//go:noescape

func GetCSSMathSumValues(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web constof_CSSNumericBaseType
//go:noescape

func ConstOfCSSNumericBaseType(str js.Ref) uint32

//go:wasmimport plat/js/web store_CSSNumericType
//go:noescape

func CSSNumericTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CSSNumericType
//go:noescape

func CSSNumericTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSNumericValue_Add
//go:noescape

func CallCSSNumericValueAdd(
	this js.Ref,
	ptr unsafe.Pointer,

	valuesPtr unsafe.Pointer,
	valuesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Add
//go:noescape

func CSSNumericValueAddFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSNumericValue_Sub
//go:noescape

func CallCSSNumericValueSub(
	this js.Ref,
	ptr unsafe.Pointer,

	valuesPtr unsafe.Pointer,
	valuesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Sub
//go:noescape

func CSSNumericValueSubFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSNumericValue_Mul
//go:noescape

func CallCSSNumericValueMul(
	this js.Ref,
	ptr unsafe.Pointer,

	valuesPtr unsafe.Pointer,
	valuesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Mul
//go:noescape

func CSSNumericValueMulFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSNumericValue_Div
//go:noescape

func CallCSSNumericValueDiv(
	this js.Ref,
	ptr unsafe.Pointer,

	valuesPtr unsafe.Pointer,
	valuesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Div
//go:noescape

func CSSNumericValueDivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSNumericValue_Min
//go:noescape

func CallCSSNumericValueMin(
	this js.Ref,
	ptr unsafe.Pointer,

	valuesPtr unsafe.Pointer,
	valuesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Min
//go:noescape

func CSSNumericValueMinFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSNumericValue_Max
//go:noescape

func CallCSSNumericValueMax(
	this js.Ref,
	ptr unsafe.Pointer,

	valuesPtr unsafe.Pointer,
	valuesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Max
//go:noescape

func CSSNumericValueMaxFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSNumericValue_Equals
//go:noescape

func CallCSSNumericValueEquals(
	this js.Ref,
	ptr unsafe.Pointer,

	valuePtr unsafe.Pointer,
	valueCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Equals
//go:noescape

func CSSNumericValueEqualsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSNumericValue_To
//go:noescape

func CallCSSNumericValueTo(
	this js.Ref,
	ptr unsafe.Pointer,

	unit js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_To
//go:noescape

func CSSNumericValueToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSNumericValue_ToSum
//go:noescape

func CallCSSNumericValueToSum(
	this js.Ref,
	ptr unsafe.Pointer,

	unitsPtr unsafe.Pointer,
	unitsCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_ToSum
//go:noescape

func CSSNumericValueToSumFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSNumericValue_Type
//go:noescape

func CallCSSNumericValueType(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Type
//go:noescape

func CSSNumericValueTypeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSNumericValue_Parse
//go:noescape

func CallCSSNumericValueParse(
	this js.Ref,
	ptr unsafe.Pointer,

	cssText js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Parse
//go:noescape

func CSSNumericValueParseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_EffectTiming
//go:noescape

func EffectTimingJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EffectTiming
//go:noescape

func EffectTimingJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ComputedEffectTiming
//go:noescape

func ComputedEffectTimingJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ComputedEffectTiming
//go:noescape

func ComputedEffectTimingJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
