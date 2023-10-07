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

//go:wasmimport plat/js/web has_ANGLE_instanced_arrays_DrawArraysInstancedANGLE
//go:noescape
func HasFuncANGLE_instanced_arraysDrawArraysInstancedANGLE(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ANGLE_instanced_arrays_DrawArraysInstancedANGLE
//go:noescape
func FuncANGLE_instanced_arraysDrawArraysInstancedANGLE(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ANGLE_instanced_arrays_DrawArraysInstancedANGLE
//go:noescape
func CallANGLE_instanced_arraysDrawArraysInstancedANGLE(
	this js.Ref, retPtr unsafe.Pointer,
	mode uint32,
	first int32,
	count int32,
	primcount int32)

//go:wasmimport plat/js/web try_ANGLE_instanced_arrays_DrawArraysInstancedANGLE
//go:noescape
func TryANGLE_instanced_arraysDrawArraysInstancedANGLE(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mode uint32,
	first int32,
	count int32,
	primcount int32) (ok js.Ref)

//go:wasmimport plat/js/web has_ANGLE_instanced_arrays_DrawElementsInstancedANGLE
//go:noescape
func HasFuncANGLE_instanced_arraysDrawElementsInstancedANGLE(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ANGLE_instanced_arrays_DrawElementsInstancedANGLE
//go:noescape
func FuncANGLE_instanced_arraysDrawElementsInstancedANGLE(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ANGLE_instanced_arrays_DrawElementsInstancedANGLE
//go:noescape
func CallANGLE_instanced_arraysDrawElementsInstancedANGLE(
	this js.Ref, retPtr unsafe.Pointer,
	mode uint32,
	count int32,
	typ uint32,
	offset float64,
	primcount int32)

//go:wasmimport plat/js/web try_ANGLE_instanced_arrays_DrawElementsInstancedANGLE
//go:noescape
func TryANGLE_instanced_arraysDrawElementsInstancedANGLE(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mode uint32,
	count int32,
	typ uint32,
	offset float64,
	primcount int32) (ok js.Ref)

//go:wasmimport plat/js/web has_ANGLE_instanced_arrays_VertexAttribDivisorANGLE
//go:noescape
func HasFuncANGLE_instanced_arraysVertexAttribDivisorANGLE(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ANGLE_instanced_arrays_VertexAttribDivisorANGLE
//go:noescape
func FuncANGLE_instanced_arraysVertexAttribDivisorANGLE(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ANGLE_instanced_arrays_VertexAttribDivisorANGLE
//go:noescape
func CallANGLE_instanced_arraysVertexAttribDivisorANGLE(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	divisor uint32)

//go:wasmimport plat/js/web try_ANGLE_instanced_arrays_VertexAttribDivisorANGLE
//go:noescape
func TryANGLE_instanced_arraysVertexAttribDivisorANGLE(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	divisor uint32) (ok js.Ref)

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

//go:wasmimport plat/js/web has_EventTarget_AddEventListener
//go:noescape
func HasFuncEventTargetAddEventListener(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EventTarget_AddEventListener
//go:noescape
func FuncEventTargetAddEventListener(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EventTarget_AddEventListener
//go:noescape
func CallEventTargetAddEventListener(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	callback js.Ref,
	options js.Ref)

//go:wasmimport plat/js/web try_EventTarget_AddEventListener
//go:noescape
func TryEventTargetAddEventListener(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	callback js.Ref,
	options js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_EventTarget_AddEventListener1
//go:noescape
func HasFuncEventTargetAddEventListener1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EventTarget_AddEventListener1
//go:noescape
func FuncEventTargetAddEventListener1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EventTarget_AddEventListener1
//go:noescape
func CallEventTargetAddEventListener1(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/web try_EventTarget_AddEventListener1
//go:noescape
func TryEventTargetAddEventListener1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_EventTarget_RemoveEventListener
//go:noescape
func HasFuncEventTargetRemoveEventListener(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EventTarget_RemoveEventListener
//go:noescape
func FuncEventTargetRemoveEventListener(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EventTarget_RemoveEventListener
//go:noescape
func CallEventTargetRemoveEventListener(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	callback js.Ref,
	options js.Ref)

//go:wasmimport plat/js/web try_EventTarget_RemoveEventListener
//go:noescape
func TryEventTargetRemoveEventListener(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	callback js.Ref,
	options js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_EventTarget_RemoveEventListener1
//go:noescape
func HasFuncEventTargetRemoveEventListener1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EventTarget_RemoveEventListener1
//go:noescape
func FuncEventTargetRemoveEventListener1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EventTarget_RemoveEventListener1
//go:noescape
func CallEventTargetRemoveEventListener1(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/web try_EventTarget_RemoveEventListener1
//go:noescape
func TryEventTargetRemoveEventListener1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_EventTarget_DispatchEvent
//go:noescape
func HasFuncEventTargetDispatchEvent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EventTarget_DispatchEvent
//go:noescape
func FuncEventTargetDispatchEvent(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EventTarget_DispatchEvent
//go:noescape
func CallEventTargetDispatchEvent(
	this js.Ref, retPtr unsafe.Pointer,
	event js.Ref)

//go:wasmimport plat/js/web try_EventTarget_DispatchEvent
//go:noescape
func TryEventTargetDispatchEvent(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	event js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Event_Target
//go:noescape
func GetEventTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Event_SrcElement
//go:noescape
func GetEventSrcElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Event_CurrentTarget
//go:noescape
func GetEventCurrentTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Event_EventPhase
//go:noescape
func GetEventEventPhase(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Event_CancelBubble
//go:noescape
func GetEventCancelBubble(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Event_CancelBubble
//go:noescape
func SetEventCancelBubble(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Event_Bubbles
//go:noescape
func GetEventBubbles(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Event_Cancelable
//go:noescape
func GetEventCancelable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Event_ReturnValue
//go:noescape
func GetEventReturnValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Event_ReturnValue
//go:noescape
func SetEventReturnValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Event_DefaultPrevented
//go:noescape
func GetEventDefaultPrevented(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Event_Composed
//go:noescape
func GetEventComposed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Event_IsTrusted
//go:noescape
func GetEventIsTrusted(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Event_TimeStamp
//go:noescape
func GetEventTimeStamp(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Event_ComposedPath
//go:noescape
func HasFuncEventComposedPath(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Event_ComposedPath
//go:noescape
func FuncEventComposedPath(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Event_ComposedPath
//go:noescape
func CallEventComposedPath(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Event_ComposedPath
//go:noescape
func TryEventComposedPath(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Event_StopPropagation
//go:noescape
func HasFuncEventStopPropagation(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Event_StopPropagation
//go:noescape
func FuncEventStopPropagation(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Event_StopPropagation
//go:noescape
func CallEventStopPropagation(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Event_StopPropagation
//go:noescape
func TryEventStopPropagation(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Event_StopImmediatePropagation
//go:noescape
func HasFuncEventStopImmediatePropagation(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Event_StopImmediatePropagation
//go:noescape
func FuncEventStopImmediatePropagation(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Event_StopImmediatePropagation
//go:noescape
func CallEventStopImmediatePropagation(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Event_StopImmediatePropagation
//go:noescape
func TryEventStopImmediatePropagation(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Event_PreventDefault
//go:noescape
func HasFuncEventPreventDefault(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Event_PreventDefault
//go:noescape
func FuncEventPreventDefault(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Event_PreventDefault
//go:noescape
func CallEventPreventDefault(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Event_PreventDefault
//go:noescape
func TryEventPreventDefault(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Event_InitEvent
//go:noescape
func HasFuncEventInitEvent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Event_InitEvent
//go:noescape
func FuncEventInitEvent(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Event_InitEvent
//go:noescape
func CallEventInitEvent(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref)

//go:wasmimport plat/js/web try_Event_InitEvent
//go:noescape
func TryEventInitEvent(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Event_InitEvent1
//go:noescape
func HasFuncEventInitEvent1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Event_InitEvent1
//go:noescape
func FuncEventInitEvent1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Event_InitEvent1
//go:noescape
func CallEventInitEvent1(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref)

//go:wasmimport plat/js/web try_Event_InitEvent1
//go:noescape
func TryEventInitEvent1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Event_InitEvent2
//go:noescape
func HasFuncEventInitEvent2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Event_InitEvent2
//go:noescape
func FuncEventInitEvent2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Event_InitEvent2
//go:noescape
func CallEventInitEvent2(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref)

//go:wasmimport plat/js/web try_Event_InitEvent2
//go:noescape
func TryEventInitEvent2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_AbortSignal_Aborted
//go:noescape
func GetAbortSignalAborted(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AbortSignal_Reason
//go:noescape
func GetAbortSignalReason(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AbortSignal_Abort
//go:noescape
func HasFuncAbortSignalAbort(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AbortSignal_Abort
//go:noescape
func FuncAbortSignalAbort(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AbortSignal_Abort
//go:noescape
func CallAbortSignalAbort(
	this js.Ref, retPtr unsafe.Pointer,
	reason js.Ref)

//go:wasmimport plat/js/web try_AbortSignal_Abort
//go:noescape
func TryAbortSignalAbort(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	reason js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AbortSignal_Abort1
//go:noescape
func HasFuncAbortSignalAbort1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AbortSignal_Abort1
//go:noescape
func FuncAbortSignalAbort1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AbortSignal_Abort1
//go:noescape
func CallAbortSignalAbort1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AbortSignal_Abort1
//go:noescape
func TryAbortSignalAbort1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AbortSignal_Timeout
//go:noescape
func HasFuncAbortSignalTimeout(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AbortSignal_Timeout
//go:noescape
func FuncAbortSignalTimeout(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AbortSignal_Timeout
//go:noescape
func CallAbortSignalTimeout(
	this js.Ref, retPtr unsafe.Pointer,
	milliseconds float64)

//go:wasmimport plat/js/web try_AbortSignal_Timeout
//go:noescape
func TryAbortSignalTimeout(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	milliseconds float64) (ok js.Ref)

//go:wasmimport plat/js/web has_AbortSignal_Any
//go:noescape
func HasFuncAbortSignalAny(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AbortSignal_Any
//go:noescape
func FuncAbortSignalAny(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AbortSignal_Any
//go:noescape
func CallAbortSignalAny(
	this js.Ref, retPtr unsafe.Pointer,
	signals js.Ref)

//go:wasmimport plat/js/web try_AbortSignal_Any
//go:noescape
func TryAbortSignalAny(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	signals js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AbortSignal_ThrowIfAborted
//go:noescape
func HasFuncAbortSignalThrowIfAborted(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AbortSignal_ThrowIfAborted
//go:noescape
func FuncAbortSignalThrowIfAborted(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AbortSignal_ThrowIfAborted
//go:noescape
func CallAbortSignalThrowIfAborted(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AbortSignal_ThrowIfAborted
//go:noescape
func TryAbortSignalThrowIfAborted(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AbortController_Signal
//go:noescape
func GetAbortControllerSignal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AbortController_Abort
//go:noescape
func HasFuncAbortControllerAbort(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AbortController_Abort
//go:noescape
func FuncAbortControllerAbort(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AbortController_Abort
//go:noescape
func CallAbortControllerAbort(
	this js.Ref, retPtr unsafe.Pointer,
	reason js.Ref)

//go:wasmimport plat/js/web try_AbortController_Abort
//go:noescape
func TryAbortControllerAbort(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	reason js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AbortController_Abort1
//go:noescape
func HasFuncAbortControllerAbort1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AbortController_Abort1
//go:noescape
func FuncAbortControllerAbort1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AbortController_Abort1
//go:noescape
func CallAbortControllerAbort1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AbortController_Abort1
//go:noescape
func TryAbortControllerAbort1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Attr_Prefix
//go:noescape
func GetAttrPrefix(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Attr_LocalName
//go:noescape
func GetAttrLocalName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Attr_Name
//go:noescape
func GetAttrName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Attr_Value
//go:noescape
func GetAttrValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Attr_Value
//go:noescape
func SetAttrValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Attr_OwnerElement
//go:noescape
func GetAttrOwnerElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Attr_Specified
//go:noescape
func GetAttrSpecified(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSUnitValue_Value
//go:noescape
func SetCSSUnitValueValue(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_CSSUnitValue_Unit
//go:noescape
func GetCSSUnitValueUnit(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSNumericArray_Length
//go:noescape
func GetCSSNumericArrayLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSNumericArray_Get
//go:noescape
func HasFuncCSSNumericArrayGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSNumericArray_Get
//go:noescape
func FuncCSSNumericArrayGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSNumericArray_Get
//go:noescape
func CallCSSNumericArrayGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_CSSNumericArray_Get
//go:noescape
func TryCSSNumericArrayGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSMathSum_CSSMathSum
//go:noescape
func NewCSSMathSumByCSSMathSum(
	argsPtr unsafe.Pointer,
	argsCount uint32) js.Ref

//go:wasmimport plat/js/web get_CSSMathSum_Values
//go:noescape
func GetCSSMathSumValues(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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

//go:wasmimport plat/js/web has_CSSNumericValue_Add
//go:noescape
func HasFuncCSSNumericValueAdd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Add
//go:noescape
func FuncCSSNumericValueAdd(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSNumericValue_Add
//go:noescape
func CallCSSNumericValueAdd(
	this js.Ref, retPtr unsafe.Pointer,
	valuesPtr unsafe.Pointer,
	valuesCount uint32)

//go:wasmimport plat/js/web try_CSSNumericValue_Add
//go:noescape
func TryCSSNumericValueAdd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	valuesPtr unsafe.Pointer,
	valuesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSNumericValue_Sub
//go:noescape
func HasFuncCSSNumericValueSub(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Sub
//go:noescape
func FuncCSSNumericValueSub(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSNumericValue_Sub
//go:noescape
func CallCSSNumericValueSub(
	this js.Ref, retPtr unsafe.Pointer,
	valuesPtr unsafe.Pointer,
	valuesCount uint32)

//go:wasmimport plat/js/web try_CSSNumericValue_Sub
//go:noescape
func TryCSSNumericValueSub(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	valuesPtr unsafe.Pointer,
	valuesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSNumericValue_Mul
//go:noescape
func HasFuncCSSNumericValueMul(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Mul
//go:noescape
func FuncCSSNumericValueMul(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSNumericValue_Mul
//go:noescape
func CallCSSNumericValueMul(
	this js.Ref, retPtr unsafe.Pointer,
	valuesPtr unsafe.Pointer,
	valuesCount uint32)

//go:wasmimport plat/js/web try_CSSNumericValue_Mul
//go:noescape
func TryCSSNumericValueMul(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	valuesPtr unsafe.Pointer,
	valuesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSNumericValue_Div
//go:noescape
func HasFuncCSSNumericValueDiv(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Div
//go:noescape
func FuncCSSNumericValueDiv(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSNumericValue_Div
//go:noescape
func CallCSSNumericValueDiv(
	this js.Ref, retPtr unsafe.Pointer,
	valuesPtr unsafe.Pointer,
	valuesCount uint32)

//go:wasmimport plat/js/web try_CSSNumericValue_Div
//go:noescape
func TryCSSNumericValueDiv(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	valuesPtr unsafe.Pointer,
	valuesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSNumericValue_Min
//go:noescape
func HasFuncCSSNumericValueMin(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Min
//go:noescape
func FuncCSSNumericValueMin(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSNumericValue_Min
//go:noescape
func CallCSSNumericValueMin(
	this js.Ref, retPtr unsafe.Pointer,
	valuesPtr unsafe.Pointer,
	valuesCount uint32)

//go:wasmimport plat/js/web try_CSSNumericValue_Min
//go:noescape
func TryCSSNumericValueMin(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	valuesPtr unsafe.Pointer,
	valuesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSNumericValue_Max
//go:noescape
func HasFuncCSSNumericValueMax(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Max
//go:noescape
func FuncCSSNumericValueMax(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSNumericValue_Max
//go:noescape
func CallCSSNumericValueMax(
	this js.Ref, retPtr unsafe.Pointer,
	valuesPtr unsafe.Pointer,
	valuesCount uint32)

//go:wasmimport plat/js/web try_CSSNumericValue_Max
//go:noescape
func TryCSSNumericValueMax(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	valuesPtr unsafe.Pointer,
	valuesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSNumericValue_Equals
//go:noescape
func HasFuncCSSNumericValueEquals(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Equals
//go:noescape
func FuncCSSNumericValueEquals(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSNumericValue_Equals
//go:noescape
func CallCSSNumericValueEquals(
	this js.Ref, retPtr unsafe.Pointer,
	valuePtr unsafe.Pointer,
	valueCount uint32)

//go:wasmimport plat/js/web try_CSSNumericValue_Equals
//go:noescape
func TryCSSNumericValueEquals(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	valuePtr unsafe.Pointer,
	valueCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSNumericValue_To
//go:noescape
func HasFuncCSSNumericValueTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_To
//go:noescape
func FuncCSSNumericValueTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSNumericValue_To
//go:noescape
func CallCSSNumericValueTo(
	this js.Ref, retPtr unsafe.Pointer,
	unit js.Ref)

//go:wasmimport plat/js/web try_CSSNumericValue_To
//go:noescape
func TryCSSNumericValueTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	unit js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSNumericValue_ToSum
//go:noescape
func HasFuncCSSNumericValueToSum(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_ToSum
//go:noescape
func FuncCSSNumericValueToSum(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSNumericValue_ToSum
//go:noescape
func CallCSSNumericValueToSum(
	this js.Ref, retPtr unsafe.Pointer,
	unitsPtr unsafe.Pointer,
	unitsCount uint32)

//go:wasmimport plat/js/web try_CSSNumericValue_ToSum
//go:noescape
func TryCSSNumericValueToSum(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	unitsPtr unsafe.Pointer,
	unitsCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSNumericValue_Type
//go:noescape
func HasFuncCSSNumericValueType(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Type
//go:noescape
func FuncCSSNumericValueType(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSNumericValue_Type
//go:noescape
func CallCSSNumericValueType(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CSSNumericValue_Type
//go:noescape
func TryCSSNumericValueType(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSNumericValue_Parse
//go:noescape
func HasFuncCSSNumericValueParse(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSNumericValue_Parse
//go:noescape
func FuncCSSNumericValueParse(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSNumericValue_Parse
//go:noescape
func CallCSSNumericValueParse(
	this js.Ref, retPtr unsafe.Pointer,
	cssText js.Ref)

//go:wasmimport plat/js/web try_CSSNumericValue_Parse
//go:noescape
func TryCSSNumericValueParse(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cssText js.Ref) (ok js.Ref)

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
