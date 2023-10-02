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

//go:wasmimport plat/js/web store_SyncEventInit
//go:noescape

func SyncEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SyncEventInit
//go:noescape

func SyncEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_SyncEvent_SyncEvent
//go:noescape

func NewSyncEventBySyncEvent(
	typ js.Ref,
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_SyncEvent_Tag
//go:noescape

func GetSyncEventTag(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SyncEvent_LastChance
//go:noescape

func GetSyncEventLastChance(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web constof_TableKind
//go:noescape

func ConstOfTableKind(str js.Ref) uint32

//go:wasmimport plat/js/web store_TableDescriptor
//go:noescape

func TableDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TableDescriptor
//go:noescape

func TableDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_Table_Table
//go:noescape

func NewTableByTable(
	descriptor unsafe.Pointer,
	value js.Ref) js.Ref

//go:wasmimport plat/js/web new_Table_Table1
//go:noescape

func NewTableByTable1(
	descriptor unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_Table_Length
//go:noescape

func GetTableLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_Table_Grow
//go:noescape

func CallTableGrow(
	this js.Ref,
	ptr unsafe.Pointer,

	delta uint32,
	value js.Ref,
) uint32

//go:wasmimport plat/js/web func_Table_Grow
//go:noescape

func TableGrowFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Table_Grow1
//go:noescape

func CallTableGrow1(
	this js.Ref,
	ptr unsafe.Pointer,

	delta uint32,
) uint32

//go:wasmimport plat/js/web func_Table_Grow1
//go:noescape

func TableGrow1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Table_Get
//go:noescape

func CallTableGet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_Table_Get
//go:noescape

func TableGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Table_Set
//go:noescape

func CallTableSet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Table_Set
//go:noescape

func TableSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Table_Set1
//go:noescape

func CallTableSet1(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_Table_Set1
//go:noescape

func TableSet1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_TaskControllerInit
//go:noescape

func TaskControllerInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TaskControllerInit
//go:noescape

func TaskControllerInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_TaskController_TaskController
//go:noescape

func NewTaskControllerByTaskController(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_TaskController_TaskController1
//go:noescape

func NewTaskControllerByTaskController1() js.Ref

//go:wasmimport plat/js/web call_TaskController_SetPriority
//go:noescape

func CallTaskControllerSetPriority(
	this js.Ref,
	ptr unsafe.Pointer,

	priority uint32,
) js.Ref

//go:wasmimport plat/js/web func_TaskController_SetPriority
//go:noescape

func TaskControllerSetPriorityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_TaskPriorityChangeEventInit
//go:noescape

func TaskPriorityChangeEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TaskPriorityChangeEventInit
//go:noescape

func TaskPriorityChangeEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_TaskPriorityChangeEvent_TaskPriorityChangeEvent
//go:noescape

func NewTaskPriorityChangeEventByTaskPriorityChangeEvent(
	typ js.Ref,
	priorityChangeEventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_TaskPriorityChangeEvent_PreviousPriority
//go:noescape

func GetTaskPriorityChangeEventPreviousPriority(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web store_TaskSignalAnyInit
//go:noescape

func TaskSignalAnyInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TaskSignalAnyInit
//go:noescape

func TaskSignalAnyInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_TaskSignal_Priority
//go:noescape

func GetTaskSignalPriority(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_TaskSignal_Any
//go:noescape

func CallTaskSignalAny(
	this js.Ref,
	ptr unsafe.Pointer,

	signals js.Ref,
	init unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_TaskSignal_Any
//go:noescape

func TaskSignalAnyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TaskSignal_Any1
//go:noescape

func CallTaskSignalAny1(
	this js.Ref,
	ptr unsafe.Pointer,

	signals js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TaskSignal_Any1
//go:noescape

func TaskSignalAny1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TestUtils_Gc
//go:noescape

func CallTestUtilsGc(
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TestUtils_Gc
//go:noescape

func TestUtilsGcFunc() js.Ref

//go:wasmimport plat/js/web store_TextDecodeOptions
//go:noescape

func TextDecodeOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TextDecodeOptions
//go:noescape

func TextDecodeOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_TextDecoderOptions
//go:noescape

func TextDecoderOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TextDecoderOptions
//go:noescape

func TextDecoderOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_TextDecoder_TextDecoder
//go:noescape

func NewTextDecoderByTextDecoder(
	label js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_TextDecoder_TextDecoder1
//go:noescape

func NewTextDecoderByTextDecoder1(
	label js.Ref) js.Ref

//go:wasmimport plat/js/web new_TextDecoder_TextDecoder2
//go:noescape

func NewTextDecoderByTextDecoder2() js.Ref

//go:wasmimport plat/js/web get_TextDecoder_Encoding
//go:noescape

func GetTextDecoderEncoding(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextDecoder_Fatal
//go:noescape

func GetTextDecoderFatal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextDecoder_IgnoreBOM
//go:noescape

func GetTextDecoderIgnoreBOM(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_TextDecoder_Decode
//go:noescape

func CallTextDecoderDecode(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_TextDecoder_Decode
//go:noescape

func TextDecoderDecodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TextDecoder_Decode1
//go:noescape

func CallTextDecoderDecode1(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TextDecoder_Decode1
//go:noescape

func TextDecoderDecode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TextDecoder_Decode2
//go:noescape

func CallTextDecoderDecode2(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TextDecoder_Decode2
//go:noescape

func TextDecoderDecode2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_TextDecoderStream_TextDecoderStream
//go:noescape

func NewTextDecoderStreamByTextDecoderStream(
	label js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_TextDecoderStream_TextDecoderStream1
//go:noescape

func NewTextDecoderStreamByTextDecoderStream1(
	label js.Ref) js.Ref

//go:wasmimport plat/js/web new_TextDecoderStream_TextDecoderStream2
//go:noescape

func NewTextDecoderStreamByTextDecoderStream2() js.Ref

//go:wasmimport plat/js/web get_TextDecoderStream_Encoding
//go:noescape

func GetTextDecoderStreamEncoding(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextDecoderStream_Fatal
//go:noescape

func GetTextDecoderStreamFatal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextDecoderStream_IgnoreBOM
//go:noescape

func GetTextDecoderStreamIgnoreBOM(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextDecoderStream_Readable
//go:noescape

func GetTextDecoderStreamReadable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextDecoderStream_Writable
//go:noescape

func GetTextDecoderStreamWritable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_TextDetector_Detect
//go:noescape

func CallTextDetectorDetect(
	this js.Ref,
	ptr unsafe.Pointer,

	image js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TextDetector_Detect
//go:noescape

func TextDetectorDetectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_TextEncoderEncodeIntoResult
//go:noescape

func TextEncoderEncodeIntoResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TextEncoderEncodeIntoResult
//go:noescape

func TextEncoderEncodeIntoResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_TextEncoder_Encoding
//go:noescape

func GetTextEncoderEncoding(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_TextEncoder_Encode
//go:noescape

func CallTextEncoderEncode(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TextEncoder_Encode
//go:noescape

func TextEncoderEncodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TextEncoder_Encode1
//go:noescape

func CallTextEncoderEncode1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TextEncoder_Encode1
//go:noescape

func TextEncoderEncode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TextEncoder_EncodeInto
//go:noescape

func CallTextEncoderEncodeInto(
	this js.Ref,
	ptr unsafe.Pointer,

	source js.Ref,
	destination js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TextEncoder_EncodeInto
//go:noescape

func TextEncoderEncodeIntoFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_TextEncoderStream_Encoding
//go:noescape

func GetTextEncoderStreamEncoding(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextEncoderStream_Readable
//go:noescape

func GetTextEncoderStreamReadable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextEncoderStream_Writable
//go:noescape

func GetTextEncoderStreamWritable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_TextFormatInit
//go:noescape

func TextFormatInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TextFormatInit
//go:noescape

func TextFormatInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_TextFormat_TextFormat
//go:noescape

func NewTextFormatByTextFormat(
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_TextFormat_TextFormat1
//go:noescape

func NewTextFormatByTextFormat1() js.Ref

//go:wasmimport plat/js/web get_TextFormat_RangeStart
//go:noescape

func GetTextFormatRangeStart(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_TextFormat_RangeEnd
//go:noescape

func GetTextFormatRangeEnd(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_TextFormat_UnderlineStyle
//go:noescape

func GetTextFormatUnderlineStyle(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextFormat_UnderlineThickness
//go:noescape

func GetTextFormatUnderlineThickness(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_TextFormatUpdateEventInit
//go:noescape

func TextFormatUpdateEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TextFormatUpdateEventInit
//go:noescape

func TextFormatUpdateEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_TextFormatUpdateEvent_TextFormatUpdateEvent
//go:noescape

func NewTextFormatUpdateEventByTextFormatUpdateEvent(
	typ js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_TextFormatUpdateEvent_TextFormatUpdateEvent1
//go:noescape

func NewTextFormatUpdateEventByTextFormatUpdateEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web call_TextFormatUpdateEvent_GetTextFormats
//go:noescape

func CallTextFormatUpdateEventGetTextFormats(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TextFormatUpdateEvent_GetTextFormats
//go:noescape

func TextFormatUpdateEventGetTextFormatsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_TextUpdateEventInit
//go:noescape

func TextUpdateEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TextUpdateEventInit
//go:noescape

func TextUpdateEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_TextUpdateEvent_TextUpdateEvent
//go:noescape

func NewTextUpdateEventByTextUpdateEvent(
	typ js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_TextUpdateEvent_TextUpdateEvent1
//go:noescape

func NewTextUpdateEventByTextUpdateEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_TextUpdateEvent_UpdateRangeStart
//go:noescape

func GetTextUpdateEventUpdateRangeStart(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_TextUpdateEvent_UpdateRangeEnd
//go:noescape

func GetTextUpdateEventUpdateRangeEnd(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_TextUpdateEvent_Text
//go:noescape

func GetTextUpdateEventText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextUpdateEvent_SelectionStart
//go:noescape

func GetTextUpdateEventSelectionStart(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_TextUpdateEvent_SelectionEnd
//go:noescape

func GetTextUpdateEventSelectionEnd(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_TextUpdateEvent_CompositionStart
//go:noescape

func GetTextUpdateEventCompositionStart(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_TextUpdateEvent_CompositionEnd
//go:noescape

func GetTextUpdateEventCompositionEnd(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web new_TimeEvent_TimeEvent
//go:noescape

func NewTimeEventByTimeEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_TimeEvent_TimeEvent1
//go:noescape

func NewTimeEventByTimeEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_TimeEvent_View
//go:noescape

func GetTimeEventView(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TimeEvent_Detail
//go:noescape

func GetTimeEventDetail(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web call_TimeEvent_InitTimeEvent
//go:noescape

func CallTimeEventInitTimeEvent(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
) js.Ref

//go:wasmimport plat/js/web func_TimeEvent_InitTimeEvent
//go:noescape

func TimeEventInitTimeEventFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_ToggleEventInit
//go:noescape

func ToggleEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ToggleEventInit
//go:noescape

func ToggleEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ToggleEvent_ToggleEvent
//go:noescape

func NewToggleEventByToggleEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ToggleEvent_ToggleEvent1
//go:noescape

func NewToggleEventByToggleEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_ToggleEvent_OldState
//go:noescape

func GetToggleEventOldState(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ToggleEvent_NewState
//go:noescape

func GetToggleEventNewState(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_TokenBinding
//go:noescape

func TokenBindingJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TokenBinding
//go:noescape

func TokenBindingJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_TokenBindingStatus
//go:noescape

func ConstOfTokenBindingStatus(str js.Ref) uint32

//go:wasmimport plat/js/web store_TopLevelStorageAccessPermissionDescriptor
//go:noescape

func TopLevelStorageAccessPermissionDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TopLevelStorageAccessPermissionDescriptor
//go:noescape

func TopLevelStorageAccessPermissionDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_TouchType
//go:noescape

func ConstOfTouchType(str js.Ref) uint32

//go:wasmimport plat/js/web store_TouchInit
//go:noescape

func TouchInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TouchInit
//go:noescape

func TouchInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_Touch_Touch
//go:noescape

func NewTouchByTouch(
	touchInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_Touch_Identifier
//go:noescape

func GetTouchIdentifier(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_Touch_Target
//go:noescape

func GetTouchTarget(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Touch_ScreenX
//go:noescape

func GetTouchScreenX(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_Touch_ScreenY
//go:noescape

func GetTouchScreenY(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_Touch_ClientX
//go:noescape

func GetTouchClientX(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_Touch_ClientY
//go:noescape

func GetTouchClientY(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_Touch_PageX
//go:noescape

func GetTouchPageX(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_Touch_PageY
//go:noescape

func GetTouchPageY(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_Touch_RadiusX
//go:noescape

func GetTouchRadiusX(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_Touch_RadiusY
//go:noescape

func GetTouchRadiusY(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_Touch_RotationAngle
//go:noescape

func GetTouchRotationAngle(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_Touch_Force
//go:noescape

func GetTouchForce(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_Touch_AltitudeAngle
//go:noescape

func GetTouchAltitudeAngle(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_Touch_AzimuthAngle
//go:noescape

func GetTouchAzimuthAngle(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_Touch_TouchType
//go:noescape

func GetTouchTouchType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web store_TouchEventInit
//go:noescape

func TouchEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TouchEventInit
//go:noescape

func TouchEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_TouchList_Length
//go:noescape

func GetTouchListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_TouchList_Item
//go:noescape

func CallTouchListItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_TouchList_Item
//go:noescape

func TouchListItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_TouchEvent_TouchEvent
//go:noescape

func NewTouchEventByTouchEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_TouchEvent_TouchEvent1
//go:noescape

func NewTouchEventByTouchEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_TouchEvent_Touches
//go:noescape

func GetTouchEventTouches(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TouchEvent_TargetTouches
//go:noescape

func GetTouchEventTargetTouches(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TouchEvent_ChangedTouches
//go:noescape

func GetTouchEventChangedTouches(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TouchEvent_AltKey
//go:noescape

func GetTouchEventAltKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TouchEvent_MetaKey
//go:noescape

func GetTouchEventMetaKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TouchEvent_CtrlKey
//go:noescape

func GetTouchEventCtrlKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TouchEvent_ShiftKey
//go:noescape

func GetTouchEventShiftKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_TouchEvent_GetModifierState
//go:noescape

func CallTouchEventGetModifierState(
	this js.Ref,
	ptr unsafe.Pointer,

	keyArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TouchEvent_GetModifierState
//go:noescape

func TouchEventGetModifierStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_TrackEventInit
//go:noescape

func TrackEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TrackEventInit
//go:noescape

func TrackEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_TrackEvent_TrackEvent
//go:noescape

func NewTrackEventByTrackEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_TrackEvent_TrackEvent1
//go:noescape

func NewTrackEventByTrackEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_TrackEvent_Track
//go:noescape

func GetTrackEventTrack(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web new_TransformStream_TransformStream
//go:noescape

func NewTransformStreamByTransformStream(
	transformer js.Ref,
	writableStrategy unsafe.Pointer,
	readableStrategy unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_TransformStream_TransformStream1
//go:noescape

func NewTransformStreamByTransformStream1(
	transformer js.Ref,
	writableStrategy unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_TransformStream_TransformStream2
//go:noescape

func NewTransformStreamByTransformStream2(
	transformer js.Ref) js.Ref

//go:wasmimport plat/js/web new_TransformStream_TransformStream3
//go:noescape

func NewTransformStreamByTransformStream3() js.Ref

//go:wasmimport plat/js/web get_TransformStream_Readable
//go:noescape

func GetTransformStreamReadable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TransformStream_Writable
//go:noescape

func GetTransformStreamWritable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TransformStreamDefaultController_DesiredSize
//go:noescape

func GetTransformStreamDefaultControllerDesiredSize(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web call_TransformStreamDefaultController_Enqueue
//go:noescape

func CallTransformStreamDefaultControllerEnqueue(
	this js.Ref,
	ptr unsafe.Pointer,

	chunk js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TransformStreamDefaultController_Enqueue
//go:noescape

func TransformStreamDefaultControllerEnqueueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TransformStreamDefaultController_Enqueue1
//go:noescape

func CallTransformStreamDefaultControllerEnqueue1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TransformStreamDefaultController_Enqueue1
//go:noescape

func TransformStreamDefaultControllerEnqueue1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TransformStreamDefaultController_Error
//go:noescape

func CallTransformStreamDefaultControllerError(
	this js.Ref,
	ptr unsafe.Pointer,

	reason js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TransformStreamDefaultController_Error
//go:noescape

func TransformStreamDefaultControllerErrorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TransformStreamDefaultController_Error1
//go:noescape

func CallTransformStreamDefaultControllerError1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TransformStreamDefaultController_Error1
//go:noescape

func TransformStreamDefaultControllerError1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TransformStreamDefaultController_Terminate
//go:noescape

func CallTransformStreamDefaultControllerTerminate(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TransformStreamDefaultController_Terminate
//go:noescape

func TransformStreamDefaultControllerTerminateFunc(this js.Ref) js.Ref
