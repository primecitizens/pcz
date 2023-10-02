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

//go:wasmimport plat/js/web constof_NavigationFocusReset
//go:noescape

func ConstOfNavigationFocusReset(str js.Ref) uint32

//go:wasmimport plat/js/web constof_NavigationScrollBehavior
//go:noescape

func ConstOfNavigationScrollBehavior(str js.Ref) uint32

//go:wasmimport plat/js/web store_NavigationInterceptOptions
//go:noescape

func NavigationInterceptOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NavigationInterceptOptions
//go:noescape

func NavigationInterceptOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_NavigateEvent_NavigateEvent
//go:noescape

func NewNavigateEventByNavigateEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_NavigateEvent_NavigationType
//go:noescape

func GetNavigateEventNavigationType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_NavigateEvent_Destination
//go:noescape

func GetNavigateEventDestination(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NavigateEvent_CanIntercept
//go:noescape

func GetNavigateEventCanIntercept(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NavigateEvent_UserInitiated
//go:noescape

func GetNavigateEventUserInitiated(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NavigateEvent_HashChange
//go:noescape

func GetNavigateEventHashChange(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NavigateEvent_Signal
//go:noescape

func GetNavigateEventSignal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NavigateEvent_FormData
//go:noescape

func GetNavigateEventFormData(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NavigateEvent_DownloadRequest
//go:noescape

func GetNavigateEventDownloadRequest(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NavigateEvent_Info
//go:noescape

func GetNavigateEventInfo(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NavigateEvent_HasUAVisualTransition
//go:noescape

func GetNavigateEventHasUAVisualTransition(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_NavigateEvent_Intercept
//go:noescape

func CallNavigateEventIntercept(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_NavigateEvent_Intercept
//go:noescape

func NavigateEventInterceptFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NavigateEvent_Intercept1
//go:noescape

func CallNavigateEventIntercept1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_NavigateEvent_Intercept1
//go:noescape

func NavigateEventIntercept1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NavigateEvent_Scroll
//go:noescape

func CallNavigateEventScroll(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_NavigateEvent_Scroll
//go:noescape

func NavigateEventScrollFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_NavigationCurrentEntryChangeEventInit
//go:noescape

func NavigationCurrentEntryChangeEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NavigationCurrentEntryChangeEventInit
//go:noescape

func NavigationCurrentEntryChangeEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_NavigationCurrentEntryChangeEvent_NavigationCurrentEntryChangeEvent
//go:noescape

func NewNavigationCurrentEntryChangeEventByNavigationCurrentEntryChangeEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_NavigationCurrentEntryChangeEvent_NavigationType
//go:noescape

func GetNavigationCurrentEntryChangeEventNavigationType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_NavigationCurrentEntryChangeEvent_From
//go:noescape

func GetNavigationCurrentEntryChangeEventFrom(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_NavigationEventInit
//go:noescape

func NavigationEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NavigationEventInit
//go:noescape

func NavigationEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_NavigationEvent_NavigationEvent
//go:noescape

func NewNavigationEventByNavigationEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_NavigationEvent_NavigationEvent1
//go:noescape

func NewNavigationEventByNavigationEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_NavigationEvent_Dir
//go:noescape

func GetNavigationEventDir(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_NavigationEvent_RelatedTarget
//go:noescape

func GetNavigationEventRelatedTarget(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web constof_NavigationTimingType
//go:noescape

func ConstOfNavigationTimingType(str js.Ref) uint32

//go:wasmimport plat/js/web store_NotificationEventInit
//go:noescape

func NotificationEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NotificationEventInit
//go:noescape

func NotificationEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_NotificationEvent_NotificationEvent
//go:noescape

func NewNotificationEventByNotificationEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_NotificationEvent_Notification
//go:noescape

func GetNotificationEventNotification(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NotificationEvent_Action
//go:noescape

func GetNotificationEventAction(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_OES_draw_buffers_indexed_EnableiOES
//go:noescape

func CallOES_draw_buffers_indexedEnableiOES(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_OES_draw_buffers_indexed_EnableiOES
//go:noescape

func OES_draw_buffers_indexedEnableiOESFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OES_draw_buffers_indexed_DisableiOES
//go:noescape

func CallOES_draw_buffers_indexedDisableiOES(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_OES_draw_buffers_indexed_DisableiOES
//go:noescape

func OES_draw_buffers_indexedDisableiOESFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OES_draw_buffers_indexed_BlendEquationiOES
//go:noescape

func CallOES_draw_buffers_indexedBlendEquationiOES(
	this js.Ref,
	ptr unsafe.Pointer,

	buf uint32,
	mode uint32,
) js.Ref

//go:wasmimport plat/js/web func_OES_draw_buffers_indexed_BlendEquationiOES
//go:noescape

func OES_draw_buffers_indexedBlendEquationiOESFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OES_draw_buffers_indexed_BlendEquationSeparateiOES
//go:noescape

func CallOES_draw_buffers_indexedBlendEquationSeparateiOES(
	this js.Ref,
	ptr unsafe.Pointer,

	buf uint32,
	modeRGB uint32,
	modeAlpha uint32,
) js.Ref

//go:wasmimport plat/js/web func_OES_draw_buffers_indexed_BlendEquationSeparateiOES
//go:noescape

func OES_draw_buffers_indexedBlendEquationSeparateiOESFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OES_draw_buffers_indexed_BlendFunciOES
//go:noescape

func CallOES_draw_buffers_indexedBlendFunciOES(
	this js.Ref,
	ptr unsafe.Pointer,

	buf uint32,
	src uint32,
	dst uint32,
) js.Ref

//go:wasmimport plat/js/web func_OES_draw_buffers_indexed_BlendFunciOES
//go:noescape

func OES_draw_buffers_indexedBlendFunciOESFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OES_draw_buffers_indexed_BlendFuncSeparateiOES
//go:noescape

func CallOES_draw_buffers_indexedBlendFuncSeparateiOES(
	this js.Ref,
	ptr unsafe.Pointer,

	buf uint32,
	srcRGB uint32,
	dstRGB uint32,
	srcAlpha uint32,
	dstAlpha uint32,
) js.Ref

//go:wasmimport plat/js/web func_OES_draw_buffers_indexed_BlendFuncSeparateiOES
//go:noescape

func OES_draw_buffers_indexedBlendFuncSeparateiOESFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OES_draw_buffers_indexed_ColorMaskiOES
//go:noescape

func CallOES_draw_buffers_indexedColorMaskiOES(
	this js.Ref,
	ptr unsafe.Pointer,

	buf uint32,
	r js.Ref,
	g js.Ref,
	b js.Ref,
	a js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_OES_draw_buffers_indexed_ColorMaskiOES
//go:noescape

func OES_draw_buffers_indexedColorMaskiOESFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OES_vertex_array_object_CreateVertexArrayOES
//go:noescape

func CallOES_vertex_array_objectCreateVertexArrayOES(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_OES_vertex_array_object_CreateVertexArrayOES
//go:noescape

func OES_vertex_array_objectCreateVertexArrayOESFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OES_vertex_array_object_DeleteVertexArrayOES
//go:noescape

func CallOES_vertex_array_objectDeleteVertexArrayOES(
	this js.Ref,
	ptr unsafe.Pointer,

	arrayObject js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_OES_vertex_array_object_DeleteVertexArrayOES
//go:noescape

func OES_vertex_array_objectDeleteVertexArrayOESFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OES_vertex_array_object_IsVertexArrayOES
//go:noescape

func CallOES_vertex_array_objectIsVertexArrayOES(
	this js.Ref,
	ptr unsafe.Pointer,

	arrayObject js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_OES_vertex_array_object_IsVertexArrayOES
//go:noescape

func OES_vertex_array_objectIsVertexArrayOESFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OES_vertex_array_object_BindVertexArrayOES
//go:noescape

func CallOES_vertex_array_objectBindVertexArrayOES(
	this js.Ref,
	ptr unsafe.Pointer,

	arrayObject js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_OES_vertex_array_object_BindVertexArrayOES
//go:noescape

func OES_vertex_array_objectBindVertexArrayOESFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_OTPCredential_Code
//go:noescape

func GetOTPCredentialCode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_OVR_multiview2_FramebufferTextureMultiviewOVR
//go:noescape

func CallOVR_multiview2FramebufferTextureMultiviewOVR(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	attachment uint32,
	texture js.Ref,
	level int32,
	baseViewIndex int32,
	numViews int32,
) js.Ref

//go:wasmimport plat/js/web func_OVR_multiview2_FramebufferTextureMultiviewOVR
//go:noescape

func OVR_multiview2FramebufferTextureMultiviewOVRFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_OfflineAudioCompletionEventInit
//go:noescape

func OfflineAudioCompletionEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_OfflineAudioCompletionEventInit
//go:noescape

func OfflineAudioCompletionEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_OfflineAudioCompletionEvent_OfflineAudioCompletionEvent
//go:noescape

func NewOfflineAudioCompletionEventByOfflineAudioCompletionEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_OfflineAudioCompletionEvent_RenderedBuffer
//go:noescape

func GetOfflineAudioCompletionEventRenderedBuffer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_OfflineAudioContextOptions
//go:noescape

func OfflineAudioContextOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_OfflineAudioContextOptions
//go:noescape

func OfflineAudioContextOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_OfflineAudioContext_OfflineAudioContext
//go:noescape

func NewOfflineAudioContextByOfflineAudioContext(
	contextOptions unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_OfflineAudioContext_OfflineAudioContext1
//go:noescape

func NewOfflineAudioContextByOfflineAudioContext1(
	numberOfChannels uint32,
	length uint32,
	sampleRate float32) js.Ref

//go:wasmimport plat/js/web get_OfflineAudioContext_Length
//go:noescape

func GetOfflineAudioContextLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_OfflineAudioContext_StartRendering
//go:noescape

func CallOfflineAudioContextStartRendering(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_OfflineAudioContext_StartRendering
//go:noescape

func OfflineAudioContextStartRenderingFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OfflineAudioContext_Resume
//go:noescape

func CallOfflineAudioContextResume(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_OfflineAudioContext_Resume
//go:noescape

func OfflineAudioContextResumeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OfflineAudioContext_Suspend
//go:noescape

func CallOfflineAudioContextSuspend(
	this js.Ref,
	ptr unsafe.Pointer,

	suspendTime float64,
) js.Ref

//go:wasmimport plat/js/web func_OfflineAudioContext_Suspend
//go:noescape

func OfflineAudioContextSuspendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_OrientationSensor_Quaternion
//go:noescape

func GetOrientationSensorQuaternion(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_OrientationSensor_PopulateMatrix
//go:noescape

func CallOrientationSensorPopulateMatrix(
	this js.Ref,
	ptr unsafe.Pointer,

	targetMatrix js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_OrientationSensor_PopulateMatrix
//go:noescape

func OrientationSensorPopulateMatrixFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_OverconstrainedError_OverconstrainedError
//go:noescape

func NewOverconstrainedErrorByOverconstrainedError(
	constraint js.Ref,
	message js.Ref) js.Ref

//go:wasmimport plat/js/web new_OverconstrainedError_OverconstrainedError1
//go:noescape

func NewOverconstrainedErrorByOverconstrainedError1(
	constraint js.Ref) js.Ref

//go:wasmimport plat/js/web get_OverconstrainedError_Constraint
//go:noescape

func GetOverconstrainedErrorConstraint(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_PageTransitionEventInit
//go:noescape

func PageTransitionEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PageTransitionEventInit
//go:noescape

func PageTransitionEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PageTransitionEvent_PageTransitionEvent
//go:noescape

func NewPageTransitionEventByPageTransitionEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_PageTransitionEvent_PageTransitionEvent1
//go:noescape

func NewPageTransitionEventByPageTransitionEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_PageTransitionEvent_Persisted
//go:noescape

func GetPageTransitionEventPersisted(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PaintRenderingContext2D_LineWidth
//go:noescape

func GetPaintRenderingContext2DLineWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_PaintRenderingContext2D_LineWidth
//go:noescape

func SetPaintRenderingContext2DLineWidth(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PaintRenderingContext2D_LineCap
//go:noescape

func GetPaintRenderingContext2DLineCap(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_PaintRenderingContext2D_LineCap
//go:noescape

func SetPaintRenderingContext2DLineCap(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_PaintRenderingContext2D_LineJoin
//go:noescape

func GetPaintRenderingContext2DLineJoin(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_PaintRenderingContext2D_LineJoin
//go:noescape

func SetPaintRenderingContext2DLineJoin(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_PaintRenderingContext2D_MiterLimit
//go:noescape

func GetPaintRenderingContext2DMiterLimit(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_PaintRenderingContext2D_MiterLimit
//go:noescape

func SetPaintRenderingContext2DMiterLimit(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PaintRenderingContext2D_LineDashOffset
//go:noescape

func GetPaintRenderingContext2DLineDashOffset(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_PaintRenderingContext2D_LineDashOffset
//go:noescape

func SetPaintRenderingContext2DLineDashOffset(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PaintRenderingContext2D_ShadowOffsetX
//go:noescape

func GetPaintRenderingContext2DShadowOffsetX(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_PaintRenderingContext2D_ShadowOffsetX
//go:noescape

func SetPaintRenderingContext2DShadowOffsetX(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PaintRenderingContext2D_ShadowOffsetY
//go:noescape

func GetPaintRenderingContext2DShadowOffsetY(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_PaintRenderingContext2D_ShadowOffsetY
//go:noescape

func SetPaintRenderingContext2DShadowOffsetY(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PaintRenderingContext2D_ShadowBlur
//go:noescape

func GetPaintRenderingContext2DShadowBlur(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_PaintRenderingContext2D_ShadowBlur
//go:noescape

func SetPaintRenderingContext2DShadowBlur(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PaintRenderingContext2D_ShadowColor
//go:noescape

func GetPaintRenderingContext2DShadowColor(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_PaintRenderingContext2D_ShadowColor
//go:noescape

func SetPaintRenderingContext2DShadowColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_PaintRenderingContext2D_StrokeStyle
//go:noescape

func GetPaintRenderingContext2DStrokeStyle(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_PaintRenderingContext2D_StrokeStyle
//go:noescape

func SetPaintRenderingContext2DStrokeStyle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_PaintRenderingContext2D_FillStyle
//go:noescape

func GetPaintRenderingContext2DFillStyle(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_PaintRenderingContext2D_FillStyle
//go:noescape

func SetPaintRenderingContext2DFillStyle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_PaintRenderingContext2D_ImageSmoothingEnabled
//go:noescape

func GetPaintRenderingContext2DImageSmoothingEnabled(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_PaintRenderingContext2D_ImageSmoothingEnabled
//go:noescape

func SetPaintRenderingContext2DImageSmoothingEnabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_PaintRenderingContext2D_ImageSmoothingQuality
//go:noescape

func GetPaintRenderingContext2DImageSmoothingQuality(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_PaintRenderingContext2D_ImageSmoothingQuality
//go:noescape

func SetPaintRenderingContext2DImageSmoothingQuality(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_PaintRenderingContext2D_GlobalAlpha
//go:noescape

func GetPaintRenderingContext2DGlobalAlpha(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_PaintRenderingContext2D_GlobalAlpha
//go:noescape

func SetPaintRenderingContext2DGlobalAlpha(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_PaintRenderingContext2D_GlobalCompositeOperation
//go:noescape

func GetPaintRenderingContext2DGlobalCompositeOperation(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_PaintRenderingContext2D_GlobalCompositeOperation
//go:noescape

func SetPaintRenderingContext2DGlobalCompositeOperation(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_ClosePath
//go:noescape

func CallPaintRenderingContext2DClosePath(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_ClosePath
//go:noescape

func PaintRenderingContext2DClosePathFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_MoveTo
//go:noescape

func CallPaintRenderingContext2DMoveTo(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_MoveTo
//go:noescape

func PaintRenderingContext2DMoveToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_LineTo
//go:noescape

func CallPaintRenderingContext2DLineTo(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_LineTo
//go:noescape

func PaintRenderingContext2DLineToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_QuadraticCurveTo
//go:noescape

func CallPaintRenderingContext2DQuadraticCurveTo(
	this js.Ref,
	ptr unsafe.Pointer,

	cpx float64,
	cpy float64,
	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_QuadraticCurveTo
//go:noescape

func PaintRenderingContext2DQuadraticCurveToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_BezierCurveTo
//go:noescape

func CallPaintRenderingContext2DBezierCurveTo(
	this js.Ref,
	ptr unsafe.Pointer,

	cp1x float64,
	cp1y float64,
	cp2x float64,
	cp2y float64,
	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_BezierCurveTo
//go:noescape

func PaintRenderingContext2DBezierCurveToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_ArcTo
//go:noescape

func CallPaintRenderingContext2DArcTo(
	this js.Ref,
	ptr unsafe.Pointer,

	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
	radius float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_ArcTo
//go:noescape

func PaintRenderingContext2DArcToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Rect
//go:noescape

func CallPaintRenderingContext2DRect(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
	w float64,
	h float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Rect
//go:noescape

func PaintRenderingContext2DRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_RoundRect
//go:noescape

func CallPaintRenderingContext2DRoundRect(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
	w float64,
	h float64,
	radii js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_RoundRect
//go:noescape

func PaintRenderingContext2DRoundRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_RoundRect1
//go:noescape

func CallPaintRenderingContext2DRoundRect1(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
	w float64,
	h float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_RoundRect1
//go:noescape

func PaintRenderingContext2DRoundRect1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Arc
//go:noescape

func CallPaintRenderingContext2DArc(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
	radius float64,
	startAngle float64,
	endAngle float64,
	counterclockwise js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Arc
//go:noescape

func PaintRenderingContext2DArcFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Arc1
//go:noescape

func CallPaintRenderingContext2DArc1(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
	radius float64,
	startAngle float64,
	endAngle float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Arc1
//go:noescape

func PaintRenderingContext2DArc1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Ellipse
//go:noescape

func CallPaintRenderingContext2DEllipse(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
	radiusX float64,
	radiusY float64,
	rotation float64,
	startAngle float64,
	endAngle float64,
	counterclockwise js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Ellipse
//go:noescape

func PaintRenderingContext2DEllipseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Ellipse1
//go:noescape

func CallPaintRenderingContext2DEllipse1(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
	radiusX float64,
	radiusY float64,
	rotation float64,
	startAngle float64,
	endAngle float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Ellipse1
//go:noescape

func PaintRenderingContext2DEllipse1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_SetLineDash
//go:noescape

func CallPaintRenderingContext2DSetLineDash(
	this js.Ref,
	ptr unsafe.Pointer,

	segments js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_SetLineDash
//go:noescape

func PaintRenderingContext2DSetLineDashFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_GetLineDash
//go:noescape

func CallPaintRenderingContext2DGetLineDash(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_GetLineDash
//go:noescape

func PaintRenderingContext2DGetLineDashFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_DrawImage
//go:noescape

func CallPaintRenderingContext2DDrawImage(
	this js.Ref,
	ptr unsafe.Pointer,

	image js.Ref,
	dx float64,
	dy float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_DrawImage
//go:noescape

func PaintRenderingContext2DDrawImageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_DrawImage1
//go:noescape

func CallPaintRenderingContext2DDrawImage1(
	this js.Ref,
	ptr unsafe.Pointer,

	image js.Ref,
	dx float64,
	dy float64,
	dw float64,
	dh float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_DrawImage1
//go:noescape

func PaintRenderingContext2DDrawImage1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_DrawImage2
//go:noescape

func CallPaintRenderingContext2DDrawImage2(
	this js.Ref,
	ptr unsafe.Pointer,

	image js.Ref,
	sx float64,
	sy float64,
	sw float64,
	sh float64,
	dx float64,
	dy float64,
	dw float64,
	dh float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_DrawImage2
//go:noescape

func PaintRenderingContext2DDrawImage2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_BeginPath
//go:noescape

func CallPaintRenderingContext2DBeginPath(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_BeginPath
//go:noescape

func PaintRenderingContext2DBeginPathFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Fill
//go:noescape

func CallPaintRenderingContext2DFill(
	this js.Ref,
	ptr unsafe.Pointer,

	fillRule uint32,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Fill
//go:noescape

func PaintRenderingContext2DFillFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Fill1
//go:noescape

func CallPaintRenderingContext2DFill1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Fill1
//go:noescape

func PaintRenderingContext2DFill1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Fill2
//go:noescape

func CallPaintRenderingContext2DFill2(
	this js.Ref,
	ptr unsafe.Pointer,

	path js.Ref,
	fillRule uint32,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Fill2
//go:noescape

func PaintRenderingContext2DFill2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Fill3
//go:noescape

func CallPaintRenderingContext2DFill3(
	this js.Ref,
	ptr unsafe.Pointer,

	path js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Fill3
//go:noescape

func PaintRenderingContext2DFill3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Stroke
//go:noescape

func CallPaintRenderingContext2DStroke(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Stroke
//go:noescape

func PaintRenderingContext2DStrokeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Stroke1
//go:noescape

func CallPaintRenderingContext2DStroke1(
	this js.Ref,
	ptr unsafe.Pointer,

	path js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Stroke1
//go:noescape

func PaintRenderingContext2DStroke1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Clip
//go:noescape

func CallPaintRenderingContext2DClip(
	this js.Ref,
	ptr unsafe.Pointer,

	fillRule uint32,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Clip
//go:noescape

func PaintRenderingContext2DClipFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Clip1
//go:noescape

func CallPaintRenderingContext2DClip1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Clip1
//go:noescape

func PaintRenderingContext2DClip1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Clip2
//go:noescape

func CallPaintRenderingContext2DClip2(
	this js.Ref,
	ptr unsafe.Pointer,

	path js.Ref,
	fillRule uint32,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Clip2
//go:noescape

func PaintRenderingContext2DClip2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Clip3
//go:noescape

func CallPaintRenderingContext2DClip3(
	this js.Ref,
	ptr unsafe.Pointer,

	path js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Clip3
//go:noescape

func PaintRenderingContext2DClip3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_IsPointInPath
//go:noescape

func CallPaintRenderingContext2DIsPointInPath(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
	fillRule uint32,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_IsPointInPath
//go:noescape

func PaintRenderingContext2DIsPointInPathFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_IsPointInPath1
//go:noescape

func CallPaintRenderingContext2DIsPointInPath1(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_IsPointInPath1
//go:noescape

func PaintRenderingContext2DIsPointInPath1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_IsPointInPath2
//go:noescape

func CallPaintRenderingContext2DIsPointInPath2(
	this js.Ref,
	ptr unsafe.Pointer,

	path js.Ref,
	x float64,
	y float64,
	fillRule uint32,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_IsPointInPath2
//go:noescape

func PaintRenderingContext2DIsPointInPath2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_IsPointInPath3
//go:noescape

func CallPaintRenderingContext2DIsPointInPath3(
	this js.Ref,
	ptr unsafe.Pointer,

	path js.Ref,
	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_IsPointInPath3
//go:noescape

func PaintRenderingContext2DIsPointInPath3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_IsPointInStroke
//go:noescape

func CallPaintRenderingContext2DIsPointInStroke(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_IsPointInStroke
//go:noescape

func PaintRenderingContext2DIsPointInStrokeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_IsPointInStroke1
//go:noescape

func CallPaintRenderingContext2DIsPointInStroke1(
	this js.Ref,
	ptr unsafe.Pointer,

	path js.Ref,
	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_IsPointInStroke1
//go:noescape

func PaintRenderingContext2DIsPointInStroke1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_ClearRect
//go:noescape

func CallPaintRenderingContext2DClearRect(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
	w float64,
	h float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_ClearRect
//go:noescape

func PaintRenderingContext2DClearRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_FillRect
//go:noescape

func CallPaintRenderingContext2DFillRect(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
	w float64,
	h float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_FillRect
//go:noescape

func PaintRenderingContext2DFillRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_StrokeRect
//go:noescape

func CallPaintRenderingContext2DStrokeRect(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
	w float64,
	h float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_StrokeRect
//go:noescape

func PaintRenderingContext2DStrokeRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_CreateLinearGradient
//go:noescape

func CallPaintRenderingContext2DCreateLinearGradient(
	this js.Ref,
	ptr unsafe.Pointer,

	x0 float64,
	y0 float64,
	x1 float64,
	y1 float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_CreateLinearGradient
//go:noescape

func PaintRenderingContext2DCreateLinearGradientFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_CreateRadialGradient
//go:noescape

func CallPaintRenderingContext2DCreateRadialGradient(
	this js.Ref,
	ptr unsafe.Pointer,

	x0 float64,
	y0 float64,
	r0 float64,
	x1 float64,
	y1 float64,
	r1 float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_CreateRadialGradient
//go:noescape

func PaintRenderingContext2DCreateRadialGradientFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_CreateConicGradient
//go:noescape

func CallPaintRenderingContext2DCreateConicGradient(
	this js.Ref,
	ptr unsafe.Pointer,

	startAngle float64,
	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_CreateConicGradient
//go:noescape

func PaintRenderingContext2DCreateConicGradientFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_CreatePattern
//go:noescape

func CallPaintRenderingContext2DCreatePattern(
	this js.Ref,
	ptr unsafe.Pointer,

	image js.Ref,
	repetition js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_CreatePattern
//go:noescape

func PaintRenderingContext2DCreatePatternFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Scale
//go:noescape

func CallPaintRenderingContext2DScale(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Scale
//go:noescape

func PaintRenderingContext2DScaleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Rotate
//go:noescape

func CallPaintRenderingContext2DRotate(
	this js.Ref,
	ptr unsafe.Pointer,

	angle float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Rotate
//go:noescape

func PaintRenderingContext2DRotateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Translate
//go:noescape

func CallPaintRenderingContext2DTranslate(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Translate
//go:noescape

func PaintRenderingContext2DTranslateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Transform
//go:noescape

func CallPaintRenderingContext2DTransform(
	this js.Ref,
	ptr unsafe.Pointer,

	a float64,
	b float64,
	c float64,
	d float64,
	e float64,
	f float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Transform
//go:noescape

func PaintRenderingContext2DTransformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_GetTransform
//go:noescape

func CallPaintRenderingContext2DGetTransform(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_GetTransform
//go:noescape

func PaintRenderingContext2DGetTransformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_SetTransform
//go:noescape

func CallPaintRenderingContext2DSetTransform(
	this js.Ref,
	ptr unsafe.Pointer,

	a float64,
	b float64,
	c float64,
	d float64,
	e float64,
	f float64,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_SetTransform
//go:noescape

func PaintRenderingContext2DSetTransformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_SetTransform1
//go:noescape

func CallPaintRenderingContext2DSetTransform1(
	this js.Ref,
	ptr unsafe.Pointer,

	transform unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_SetTransform1
//go:noescape

func PaintRenderingContext2DSetTransform1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_SetTransform2
//go:noescape

func CallPaintRenderingContext2DSetTransform2(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_SetTransform2
//go:noescape

func PaintRenderingContext2DSetTransform2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_ResetTransform
//go:noescape

func CallPaintRenderingContext2DResetTransform(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_ResetTransform
//go:noescape

func PaintRenderingContext2DResetTransformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Save
//go:noescape

func CallPaintRenderingContext2DSave(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Save
//go:noescape

func PaintRenderingContext2DSaveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Restore
//go:noescape

func CallPaintRenderingContext2DRestore(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Restore
//go:noescape

func PaintRenderingContext2DRestoreFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_Reset
//go:noescape

func CallPaintRenderingContext2DReset(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_Reset
//go:noescape

func PaintRenderingContext2DResetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaintRenderingContext2D_IsContextLost
//go:noescape

func CallPaintRenderingContext2DIsContextLost(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PaintRenderingContext2D_IsContextLost
//go:noescape

func PaintRenderingContext2DIsContextLostFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_PaintRenderingContext2DSettings
//go:noescape

func PaintRenderingContext2DSettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaintRenderingContext2DSettings
//go:noescape

func PaintRenderingContext2DSettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_PaintSize_Width
//go:noescape

func GetPaintSizeWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_PaintSize_Height
//go:noescape

func GetPaintSizeHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_PaintWorkletGlobalScope_DevicePixelRatio
//go:noescape

func GetPaintWorkletGlobalScopeDevicePixelRatio(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web call_PaintWorkletGlobalScope_RegisterPaint
//go:noescape

func CallPaintWorkletGlobalScopeRegisterPaint(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	paintCtor js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PaintWorkletGlobalScope_RegisterPaint
//go:noescape

func PaintWorkletGlobalScopeRegisterPaintFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_PasswordCredential_PasswordCredential
//go:noescape

func NewPasswordCredentialByPasswordCredential(
	form js.Ref) js.Ref

//go:wasmimport plat/js/web new_PasswordCredential_PasswordCredential1
//go:noescape

func NewPasswordCredentialByPasswordCredential1(
	data unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_PasswordCredential_Password
//go:noescape

func GetPasswordCredentialPassword(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PasswordCredential_Name
//go:noescape

func GetPasswordCredentialName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PasswordCredential_IconURL
//go:noescape

func GetPasswordCredentialIconURL(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web constof_PaymentComplete
//go:noescape

func ConstOfPaymentComplete(str js.Ref) uint32

//go:wasmimport plat/js/web store_PaymentCompleteDetails
//go:noescape

func PaymentCompleteDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentCompleteDetails
//go:noescape

func PaymentCompleteDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PaymentItem
//go:noescape

func PaymentItemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentItem
//go:noescape

func PaymentItemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
