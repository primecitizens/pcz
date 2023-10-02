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

//go:wasmimport plat/js/web constof_RecordingState
//go:noescape

func ConstOfRecordingState(str js.Ref) uint32

//go:wasmimport plat/js/web new_MediaRecorder_MediaRecorder
//go:noescape

func NewMediaRecorderByMediaRecorder(
	stream js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_MediaRecorder_MediaRecorder1
//go:noescape

func NewMediaRecorderByMediaRecorder1(
	stream js.Ref) js.Ref

//go:wasmimport plat/js/web get_MediaRecorder_Stream
//go:noescape

func GetMediaRecorderStream(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MediaRecorder_MimeType
//go:noescape

func GetMediaRecorderMimeType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MediaRecorder_State
//go:noescape

func GetMediaRecorderState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_MediaRecorder_VideoBitsPerSecond
//go:noescape

func GetMediaRecorderVideoBitsPerSecond(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_MediaRecorder_AudioBitsPerSecond
//go:noescape

func GetMediaRecorderAudioBitsPerSecond(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_MediaRecorder_AudioBitrateMode
//go:noescape

func GetMediaRecorderAudioBitrateMode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_MediaRecorder_Start
//go:noescape

func CallMediaRecorderStart(
	this js.Ref,
	ptr unsafe.Pointer,

	timeslice uint32,
) js.Ref

//go:wasmimport plat/js/web func_MediaRecorder_Start
//go:noescape

func MediaRecorderStartFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaRecorder_Start1
//go:noescape

func CallMediaRecorderStart1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaRecorder_Start1
//go:noescape

func MediaRecorderStart1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaRecorder_Stop
//go:noescape

func CallMediaRecorderStop(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaRecorder_Stop
//go:noescape

func MediaRecorderStopFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaRecorder_Pause
//go:noescape

func CallMediaRecorderPause(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaRecorder_Pause
//go:noescape

func MediaRecorderPauseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaRecorder_Resume
//go:noescape

func CallMediaRecorderResume(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaRecorder_Resume
//go:noescape

func MediaRecorderResumeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaRecorder_RequestData
//go:noescape

func CallMediaRecorderRequestData(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaRecorder_RequestData
//go:noescape

func MediaRecorderRequestDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaRecorder_IsTypeSupported
//go:noescape

func CallMediaRecorderIsTypeSupported(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaRecorder_IsTypeSupported
//go:noescape

func MediaRecorderIsTypeSupportedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaStreamTrackEventInit
//go:noescape

func MediaStreamTrackEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaStreamTrackEventInit
//go:noescape

func MediaStreamTrackEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MediaStreamTrackEvent_MediaStreamTrackEvent
//go:noescape

func NewMediaStreamTrackEventByMediaStreamTrackEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_MediaStreamTrackEvent_Track
//go:noescape

func GetMediaStreamTrackEventTrack(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_MediaStreamTrackProcessorInit
//go:noescape

func MediaStreamTrackProcessorInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaStreamTrackProcessorInit
//go:noescape

func MediaStreamTrackProcessorInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MediaStreamTrackProcessor_MediaStreamTrackProcessor
//go:noescape

func NewMediaStreamTrackProcessorByMediaStreamTrackProcessor(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_MediaStreamTrackProcessor_Readable
//go:noescape

func GetMediaStreamTrackProcessorReadable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_MediaStreamTrackProcessor_Readable
//go:noescape

func SetMediaStreamTrackProcessorReadable(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_MemoryDescriptor
//go:noescape

func MemoryDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MemoryDescriptor
//go:noescape

func MemoryDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_Memory_Memory
//go:noescape

func NewMemoryByMemory(
	descriptor unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_Memory_Buffer
//go:noescape

func GetMemoryBuffer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_Memory_Grow
//go:noescape

func CallMemoryGrow(
	this js.Ref,
	ptr unsafe.Pointer,

	delta uint32,
) uint32

//go:wasmimport plat/js/web func_Memory_Grow
//go:noescape

func MemoryGrowFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_MessageChannel_Port1
//go:noescape

func GetMessageChannelPort1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MessageChannel_Port2
//go:noescape

func GetMessageChannelPort2(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_MessageEventInit
//go:noescape

func MessageEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MessageEventInit
//go:noescape

func MessageEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MessageEvent_MessageEvent
//go:noescape

func NewMessageEventByMessageEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_MessageEvent_MessageEvent1
//go:noescape

func NewMessageEventByMessageEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_MessageEvent_Data
//go:noescape

func GetMessageEventData(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MessageEvent_Origin
//go:noescape

func GetMessageEventOrigin(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MessageEvent_LastEventId
//go:noescape

func GetMessageEventLastEventId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MessageEvent_Source
//go:noescape

func GetMessageEventSource(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MessageEvent_Ports
//go:noescape

func GetMessageEventPorts(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent
//go:noescape

func CallMessageEventInitMessageEvent(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	data js.Ref,
	origin js.Ref,
	lastEventId js.Ref,
	source js.Ref,
	ports js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent
//go:noescape

func MessageEventInitMessageEventFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent1
//go:noescape

func CallMessageEventInitMessageEvent1(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	data js.Ref,
	origin js.Ref,
	lastEventId js.Ref,
	source js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent1
//go:noescape

func MessageEventInitMessageEvent1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent2
//go:noescape

func CallMessageEventInitMessageEvent2(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	data js.Ref,
	origin js.Ref,
	lastEventId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent2
//go:noescape

func MessageEventInitMessageEvent2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent3
//go:noescape

func CallMessageEventInitMessageEvent3(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	data js.Ref,
	origin js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent3
//go:noescape

func MessageEventInitMessageEvent3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent4
//go:noescape

func CallMessageEventInitMessageEvent4(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent4
//go:noescape

func MessageEventInitMessageEvent4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent5
//go:noescape

func CallMessageEventInitMessageEvent5(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent5
//go:noescape

func MessageEventInitMessageEvent5Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent6
//go:noescape

func CallMessageEventInitMessageEvent6(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent6
//go:noescape

func MessageEventInitMessageEvent6Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent7
//go:noescape

func CallMessageEventInitMessageEvent7(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent7
//go:noescape

func MessageEventInitMessageEvent7Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MeteringMode
//go:noescape

func ConstOfMeteringMode(str js.Ref) uint32

//go:wasmimport plat/js/web store_MidiPermissionDescriptor
//go:noescape

func MidiPermissionDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MidiPermissionDescriptor
//go:noescape

func MidiPermissionDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MockCameraConfiguration
//go:noescape

func MockCameraConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MockCameraConfiguration
//go:noescape

func MockCameraConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MockCaptureDeviceConfiguration
//go:noescape

func MockCaptureDeviceConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MockCaptureDeviceConfiguration
//go:noescape

func MockCaptureDeviceConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MockCapturePromptResult
//go:noescape

func ConstOfMockCapturePromptResult(str js.Ref) uint32

//go:wasmimport plat/js/web store_MockCapturePromptResultConfiguration
//go:noescape

func MockCapturePromptResultConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MockCapturePromptResultConfiguration
//go:noescape

func MockCapturePromptResultConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MockMicrophoneConfiguration
//go:noescape

func MockMicrophoneConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MockMicrophoneConfiguration
//go:noescape

func MockMicrophoneConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MockSensor
//go:noescape

func MockSensorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MockSensor
//go:noescape

func MockSensorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MockSensorType
//go:noescape

func ConstOfMockSensorType(str js.Ref) uint32

//go:wasmimport plat/js/web store_MockSensorConfiguration
//go:noescape

func MockSensorConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MockSensorConfiguration
//go:noescape

func MockSensorConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MockSensorReadingValues
//go:noescape

func MockSensorReadingValuesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MockSensorReadingValues
//go:noescape

func MockSensorReadingValuesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MouseEventInit
//go:noescape

func MouseEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MouseEventInit
//go:noescape

func MouseEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MouseEvent_MouseEvent
//go:noescape

func NewMouseEventByMouseEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_MouseEvent_MouseEvent1
//go:noescape

func NewMouseEventByMouseEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_MouseEvent_ScreenX
//go:noescape

func GetMouseEventScreenX(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_MouseEvent_ScreenY
//go:noescape

func GetMouseEventScreenY(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_MouseEvent_ClientX
//go:noescape

func GetMouseEventClientX(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_MouseEvent_ClientY
//go:noescape

func GetMouseEventClientY(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_MouseEvent_CtrlKey
//go:noescape

func GetMouseEventCtrlKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MouseEvent_ShiftKey
//go:noescape

func GetMouseEventShiftKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MouseEvent_AltKey
//go:noescape

func GetMouseEventAltKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MouseEvent_MetaKey
//go:noescape

func GetMouseEventMetaKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MouseEvent_Button
//go:noescape

func GetMouseEventButton(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_MouseEvent_Buttons
//go:noescape

func GetMouseEventButtons(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_MouseEvent_RelatedTarget
//go:noescape

func GetMouseEventRelatedTarget(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MouseEvent_MovementX
//go:noescape

func GetMouseEventMovementX(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_MouseEvent_MovementY
//go:noescape

func GetMouseEventMovementY(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_MouseEvent_PageX
//go:noescape

func GetMouseEventPageX(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_MouseEvent_PageY
//go:noescape

func GetMouseEventPageY(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_MouseEvent_X
//go:noescape

func GetMouseEventX(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_MouseEvent_Y
//go:noescape

func GetMouseEventY(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_MouseEvent_OffsetX
//go:noescape

func GetMouseEventOffsetX(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_MouseEvent_OffsetY
//go:noescape

func GetMouseEventOffsetY(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web call_MouseEvent_GetModifierState
//go:noescape

func CallMouseEventGetModifierState(
	this js.Ref,
	ptr unsafe.Pointer,

	keyArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_GetModifierState
//go:noescape

func MouseEventGetModifierStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent
//go:noescape

func CallMouseEventInitMouseEvent(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32,
	clientXArg int32,
	clientYArg int32,
	ctrlKeyArg js.Ref,
	altKeyArg js.Ref,
	shiftKeyArg js.Ref,
	metaKeyArg js.Ref,
	buttonArg int32,
	relatedTargetArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent
//go:noescape

func MouseEventInitMouseEventFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent1
//go:noescape

func CallMouseEventInitMouseEvent1(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32,
	clientXArg int32,
	clientYArg int32,
	ctrlKeyArg js.Ref,
	altKeyArg js.Ref,
	shiftKeyArg js.Ref,
	metaKeyArg js.Ref,
	buttonArg int32,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent1
//go:noescape

func MouseEventInitMouseEvent1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent2
//go:noescape

func CallMouseEventInitMouseEvent2(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32,
	clientXArg int32,
	clientYArg int32,
	ctrlKeyArg js.Ref,
	altKeyArg js.Ref,
	shiftKeyArg js.Ref,
	metaKeyArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent2
//go:noescape

func MouseEventInitMouseEvent2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent3
//go:noescape

func CallMouseEventInitMouseEvent3(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32,
	clientXArg int32,
	clientYArg int32,
	ctrlKeyArg js.Ref,
	altKeyArg js.Ref,
	shiftKeyArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent3
//go:noescape

func MouseEventInitMouseEvent3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent4
//go:noescape

func CallMouseEventInitMouseEvent4(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32,
	clientXArg int32,
	clientYArg int32,
	ctrlKeyArg js.Ref,
	altKeyArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent4
//go:noescape

func MouseEventInitMouseEvent4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent5
//go:noescape

func CallMouseEventInitMouseEvent5(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32,
	clientXArg int32,
	clientYArg int32,
	ctrlKeyArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent5
//go:noescape

func MouseEventInitMouseEvent5Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent6
//go:noescape

func CallMouseEventInitMouseEvent6(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32,
	clientXArg int32,
	clientYArg int32,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent6
//go:noescape

func MouseEventInitMouseEvent6Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent7
//go:noescape

func CallMouseEventInitMouseEvent7(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32,
	clientXArg int32,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent7
//go:noescape

func MouseEventInitMouseEvent7Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent8
//go:noescape

func CallMouseEventInitMouseEvent8(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent8
//go:noescape

func MouseEventInitMouseEvent8Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent9
//go:noescape

func CallMouseEventInitMouseEvent9(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent9
//go:noescape

func MouseEventInitMouseEvent9Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent10
//go:noescape

func CallMouseEventInitMouseEvent10(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent10
//go:noescape

func MouseEventInitMouseEvent10Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent11
//go:noescape

func CallMouseEventInitMouseEvent11(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent11
//go:noescape

func MouseEventInitMouseEvent11Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent12
//go:noescape

func CallMouseEventInitMouseEvent12(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent12
//go:noescape

func MouseEventInitMouseEvent12Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent13
//go:noescape

func CallMouseEventInitMouseEvent13(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent13
//go:noescape

func MouseEventInitMouseEvent13Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent14
//go:noescape

func CallMouseEventInitMouseEvent14(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent14
//go:noescape

func MouseEventInitMouseEvent14Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_MutationRecord_Type
//go:noescape

func GetMutationRecordType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MutationRecord_Target
//go:noescape

func GetMutationRecordTarget(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MutationRecord_AddedNodes
//go:noescape

func GetMutationRecordAddedNodes(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MutationRecord_RemovedNodes
//go:noescape

func GetMutationRecordRemovedNodes(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MutationRecord_PreviousSibling
//go:noescape

func GetMutationRecordPreviousSibling(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MutationRecord_NextSibling
//go:noescape

func GetMutationRecordNextSibling(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MutationRecord_AttributeName
//go:noescape

func GetMutationRecordAttributeName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MutationRecord_AttributeNamespace
//go:noescape

func GetMutationRecordAttributeNamespace(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MutationRecord_OldValue
//go:noescape

func GetMutationRecordOldValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_MutationObserverInit
//go:noescape

func MutationObserverInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MutationObserverInit
//go:noescape

func MutationObserverInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MutationObserver_MutationObserver
//go:noescape

func NewMutationObserverByMutationObserver(
	callback js.Ref) js.Ref

//go:wasmimport plat/js/web call_MutationObserver_Observe
//go:noescape

func CallMutationObserverObserve(
	this js.Ref,
	ptr unsafe.Pointer,

	target js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_MutationObserver_Observe
//go:noescape

func MutationObserverObserveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MutationObserver_Observe1
//go:noescape

func CallMutationObserverObserve1(
	this js.Ref,
	ptr unsafe.Pointer,

	target js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MutationObserver_Observe1
//go:noescape

func MutationObserverObserve1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MutationObserver_Disconnect
//go:noescape

func CallMutationObserverDisconnect(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MutationObserver_Disconnect
//go:noescape

func MutationObserverDisconnectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MutationObserver_TakeRecords
//go:noescape

func CallMutationObserverTakeRecords(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MutationObserver_TakeRecords
//go:noescape

func MutationObserverTakeRecordsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_MutationEvent_MutationEvent
//go:noescape

func NewMutationEventByMutationEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_MutationEvent_MutationEvent1
//go:noescape

func NewMutationEventByMutationEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_MutationEvent_RelatedNode
//go:noescape

func GetMutationEventRelatedNode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MutationEvent_PrevValue
//go:noescape

func GetMutationEventPrevValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MutationEvent_NewValue
//go:noescape

func GetMutationEventNewValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MutationEvent_AttrName
//go:noescape

func GetMutationEventAttrName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MutationEvent_AttrChange
//go:noescape

func GetMutationEventAttrChange(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent
//go:noescape

func CallMutationEventInitMutationEvent(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	relatedNodeArg js.Ref,
	prevValueArg js.Ref,
	newValueArg js.Ref,
	attrNameArg js.Ref,
	attrChangeArg uint32,
) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent
//go:noescape

func MutationEventInitMutationEventFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent1
//go:noescape

func CallMutationEventInitMutationEvent1(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	relatedNodeArg js.Ref,
	prevValueArg js.Ref,
	newValueArg js.Ref,
	attrNameArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent1
//go:noescape

func MutationEventInitMutationEvent1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent2
//go:noescape

func CallMutationEventInitMutationEvent2(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	relatedNodeArg js.Ref,
	prevValueArg js.Ref,
	newValueArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent2
//go:noescape

func MutationEventInitMutationEvent2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent3
//go:noescape

func CallMutationEventInitMutationEvent3(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	relatedNodeArg js.Ref,
	prevValueArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent3
//go:noescape

func MutationEventInitMutationEvent3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent4
//go:noescape

func CallMutationEventInitMutationEvent4(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	relatedNodeArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent4
//go:noescape

func MutationEventInitMutationEvent4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent5
//go:noescape

func CallMutationEventInitMutationEvent5(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent5
//go:noescape

func MutationEventInitMutationEvent5Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent6
//go:noescape

func CallMutationEventInitMutationEvent6(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent6
//go:noescape

func MutationEventInitMutationEvent6Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent7
//go:noescape

func CallMutationEventInitMutationEvent7(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent7
//go:noescape

func MutationEventInitMutationEvent7Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_NDEFMakeReadOnlyOptions
//go:noescape

func NDEFMakeReadOnlyOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NDEFMakeReadOnlyOptions
//go:noescape

func NDEFMakeReadOnlyOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_NDEFRecordInit
//go:noescape

func NDEFRecordInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NDEFRecordInit
//go:noescape

func NDEFRecordInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_NDEFMessageInit
//go:noescape

func NDEFMessageInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NDEFMessageInit
//go:noescape

func NDEFMessageInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_NDEFRecord_NDEFRecord
//go:noescape

func NewNDEFRecordByNDEFRecord(
	recordInit unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_NDEFRecord_RecordType
//go:noescape

func GetNDEFRecordRecordType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NDEFRecord_MediaType
//go:noescape

func GetNDEFRecordMediaType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NDEFRecord_Id
//go:noescape

func GetNDEFRecordId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NDEFRecord_Data
//go:noescape

func GetNDEFRecordData(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NDEFRecord_Encoding
//go:noescape

func GetNDEFRecordEncoding(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NDEFRecord_Lang
//go:noescape

func GetNDEFRecordLang(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_NDEFRecord_ToRecords
//go:noescape

func CallNDEFRecordToRecords(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_NDEFRecord_ToRecords
//go:noescape

func NDEFRecordToRecordsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_NDEFMessage_NDEFMessage
//go:noescape

func NewNDEFMessageByNDEFMessage(
	messageInit unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_NDEFMessage_Records
//go:noescape

func GetNDEFMessageRecords(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_NDEFScanOptions
//go:noescape

func NDEFScanOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NDEFScanOptions
//go:noescape

func NDEFScanOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_NDEFWriteOptions
//go:noescape

func NDEFWriteOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NDEFWriteOptions
//go:noescape

func NDEFWriteOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_NDEFReader_Scan
//go:noescape

func CallNDEFReaderScan(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_NDEFReader_Scan
//go:noescape

func NDEFReaderScanFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NDEFReader_Scan1
//go:noescape

func CallNDEFReaderScan1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_NDEFReader_Scan1
//go:noescape

func NDEFReaderScan1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NDEFReader_Write
//go:noescape

func CallNDEFReaderWrite(
	this js.Ref,
	ptr unsafe.Pointer,

	message js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_NDEFReader_Write
//go:noescape

func NDEFReaderWriteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NDEFReader_Write1
//go:noescape

func CallNDEFReaderWrite1(
	this js.Ref,
	ptr unsafe.Pointer,

	message js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_NDEFReader_Write1
//go:noescape

func NDEFReaderWrite1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NDEFReader_MakeReadOnly
//go:noescape

func CallNDEFReaderMakeReadOnly(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_NDEFReader_MakeReadOnly
//go:noescape

func NDEFReaderMakeReadOnlyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NDEFReader_MakeReadOnly1
//go:noescape

func CallNDEFReaderMakeReadOnly1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_NDEFReader_MakeReadOnly1
//go:noescape

func NDEFReaderMakeReadOnly1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_NDEFReadingEventInit
//go:noescape

func NDEFReadingEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NDEFReadingEventInit
//go:noescape

func NDEFReadingEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_NDEFReadingEvent_NDEFReadingEvent
//go:noescape

func NewNDEFReadingEventByNDEFReadingEvent(
	typ js.Ref,
	readingEventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_NDEFReadingEvent_SerialNumber
//go:noescape

func GetNDEFReadingEventSerialNumber(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NDEFReadingEvent_Message
//go:noescape

func GetNDEFReadingEventMessage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NamedFlow_Name
//go:noescape

func GetNamedFlowName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NamedFlow_Overset
//go:noescape

func GetNamedFlowOverset(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NamedFlow_FirstEmptyRegionIndex
//go:noescape

func GetNamedFlowFirstEmptyRegionIndex(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web call_NamedFlow_GetRegions
//go:noescape

func CallNamedFlowGetRegions(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_NamedFlow_GetRegions
//go:noescape

func NamedFlowGetRegionsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NamedFlow_GetContent
//go:noescape

func CallNamedFlowGetContent(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_NamedFlow_GetContent
//go:noescape

func NamedFlowGetContentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NamedFlow_GetRegionsByContent
//go:noescape

func CallNamedFlowGetRegionsByContent(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_NamedFlow_GetRegionsByContent
//go:noescape

func NamedFlowGetRegionsByContentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_NavigationDestination_Url
//go:noescape

func GetNavigationDestinationUrl(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NavigationDestination_Key
//go:noescape

func GetNavigationDestinationKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NavigationDestination_Id
//go:noescape

func GetNavigationDestinationId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NavigationDestination_Index
//go:noescape

func GetNavigationDestinationIndex(
	this js.Ref,
	ptr unsafe.Pointer,
) int64

//go:wasmimport plat/js/web get_NavigationDestination_SameDocument
//go:noescape

func GetNavigationDestinationSameDocument(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_NavigationDestination_GetState
//go:noescape

func CallNavigationDestinationGetState(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_NavigationDestination_GetState
//go:noescape

func NavigationDestinationGetStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_NavigateEventInit
//go:noescape

func NavigateEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NavigateEventInit
//go:noescape

func NavigateEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
