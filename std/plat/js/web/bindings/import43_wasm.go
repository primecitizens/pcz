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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaRecorder_MimeType
//go:noescape
func GetMediaRecorderMimeType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaRecorder_State
//go:noescape
func GetMediaRecorderState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaRecorder_VideoBitsPerSecond
//go:noescape
func GetMediaRecorderVideoBitsPerSecond(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaRecorder_AudioBitsPerSecond
//go:noescape
func GetMediaRecorderAudioBitsPerSecond(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaRecorder_AudioBitrateMode
//go:noescape
func GetMediaRecorderAudioBitrateMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaRecorder_Start
//go:noescape
func HasFuncMediaRecorderStart(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaRecorder_Start
//go:noescape
func FuncMediaRecorderStart(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaRecorder_Start
//go:noescape
func CallMediaRecorderStart(
	this js.Ref, retPtr unsafe.Pointer,
	timeslice uint32)

//go:wasmimport plat/js/web try_MediaRecorder_Start
//go:noescape
func TryMediaRecorderStart(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	timeslice uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaRecorder_Start1
//go:noescape
func HasFuncMediaRecorderStart1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaRecorder_Start1
//go:noescape
func FuncMediaRecorderStart1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaRecorder_Start1
//go:noescape
func CallMediaRecorderStart1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaRecorder_Start1
//go:noescape
func TryMediaRecorderStart1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaRecorder_Stop
//go:noescape
func HasFuncMediaRecorderStop(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaRecorder_Stop
//go:noescape
func FuncMediaRecorderStop(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaRecorder_Stop
//go:noescape
func CallMediaRecorderStop(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaRecorder_Stop
//go:noescape
func TryMediaRecorderStop(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaRecorder_Pause
//go:noescape
func HasFuncMediaRecorderPause(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaRecorder_Pause
//go:noescape
func FuncMediaRecorderPause(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaRecorder_Pause
//go:noescape
func CallMediaRecorderPause(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaRecorder_Pause
//go:noescape
func TryMediaRecorderPause(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaRecorder_Resume
//go:noescape
func HasFuncMediaRecorderResume(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaRecorder_Resume
//go:noescape
func FuncMediaRecorderResume(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaRecorder_Resume
//go:noescape
func CallMediaRecorderResume(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaRecorder_Resume
//go:noescape
func TryMediaRecorderResume(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaRecorder_RequestData
//go:noescape
func HasFuncMediaRecorderRequestData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaRecorder_RequestData
//go:noescape
func FuncMediaRecorderRequestData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaRecorder_RequestData
//go:noescape
func CallMediaRecorderRequestData(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaRecorder_RequestData
//go:noescape
func TryMediaRecorderRequestData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaRecorder_IsTypeSupported
//go:noescape
func HasFuncMediaRecorderIsTypeSupported(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaRecorder_IsTypeSupported
//go:noescape
func FuncMediaRecorderIsTypeSupported(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaRecorder_IsTypeSupported
//go:noescape
func CallMediaRecorderIsTypeSupported(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref)

//go:wasmimport plat/js/web try_MediaRecorder_IsTypeSupported
//go:noescape
func TryMediaRecorderIsTypeSupported(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Memory_Grow
//go:noescape
func HasFuncMemoryGrow(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Memory_Grow
//go:noescape
func FuncMemoryGrow(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Memory_Grow
//go:noescape
func CallMemoryGrow(
	this js.Ref, retPtr unsafe.Pointer,
	delta uint32)

//go:wasmimport plat/js/web try_Memory_Grow
//go:noescape
func TryMemoryGrow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	delta uint32) (ok js.Ref)

//go:wasmimport plat/js/web get_MessageChannel_Port1
//go:noescape
func GetMessageChannelPort1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MessageChannel_Port2
//go:noescape
func GetMessageChannelPort2(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MessageEvent_Origin
//go:noescape
func GetMessageEventOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MessageEvent_LastEventId
//go:noescape
func GetMessageEventLastEventId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MessageEvent_Source
//go:noescape
func GetMessageEventSource(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MessageEvent_Ports
//go:noescape
func GetMessageEventPorts(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MessageEvent_InitMessageEvent
//go:noescape
func HasFuncMessageEventInitMessageEvent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent
//go:noescape
func FuncMessageEventInitMessageEvent(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent
//go:noescape
func CallMessageEventInitMessageEvent(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	data js.Ref,
	origin js.Ref,
	lastEventId js.Ref,
	source js.Ref,
	ports js.Ref)

//go:wasmimport plat/js/web try_MessageEvent_InitMessageEvent
//go:noescape
func TryMessageEventInitMessageEvent(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	data js.Ref,
	origin js.Ref,
	lastEventId js.Ref,
	source js.Ref,
	ports js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MessageEvent_InitMessageEvent1
//go:noescape
func HasFuncMessageEventInitMessageEvent1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent1
//go:noescape
func FuncMessageEventInitMessageEvent1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent1
//go:noescape
func CallMessageEventInitMessageEvent1(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	data js.Ref,
	origin js.Ref,
	lastEventId js.Ref,
	source js.Ref)

//go:wasmimport plat/js/web try_MessageEvent_InitMessageEvent1
//go:noescape
func TryMessageEventInitMessageEvent1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	data js.Ref,
	origin js.Ref,
	lastEventId js.Ref,
	source js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MessageEvent_InitMessageEvent2
//go:noescape
func HasFuncMessageEventInitMessageEvent2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent2
//go:noescape
func FuncMessageEventInitMessageEvent2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent2
//go:noescape
func CallMessageEventInitMessageEvent2(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	data js.Ref,
	origin js.Ref,
	lastEventId js.Ref)

//go:wasmimport plat/js/web try_MessageEvent_InitMessageEvent2
//go:noescape
func TryMessageEventInitMessageEvent2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	data js.Ref,
	origin js.Ref,
	lastEventId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MessageEvent_InitMessageEvent3
//go:noescape
func HasFuncMessageEventInitMessageEvent3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent3
//go:noescape
func FuncMessageEventInitMessageEvent3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent3
//go:noescape
func CallMessageEventInitMessageEvent3(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	data js.Ref,
	origin js.Ref)

//go:wasmimport plat/js/web try_MessageEvent_InitMessageEvent3
//go:noescape
func TryMessageEventInitMessageEvent3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	data js.Ref,
	origin js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MessageEvent_InitMessageEvent4
//go:noescape
func HasFuncMessageEventInitMessageEvent4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent4
//go:noescape
func FuncMessageEventInitMessageEvent4(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent4
//go:noescape
func CallMessageEventInitMessageEvent4(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	data js.Ref)

//go:wasmimport plat/js/web try_MessageEvent_InitMessageEvent4
//go:noescape
func TryMessageEventInitMessageEvent4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MessageEvent_InitMessageEvent5
//go:noescape
func HasFuncMessageEventInitMessageEvent5(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent5
//go:noescape
func FuncMessageEventInitMessageEvent5(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent5
//go:noescape
func CallMessageEventInitMessageEvent5(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref)

//go:wasmimport plat/js/web try_MessageEvent_InitMessageEvent5
//go:noescape
func TryMessageEventInitMessageEvent5(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MessageEvent_InitMessageEvent6
//go:noescape
func HasFuncMessageEventInitMessageEvent6(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent6
//go:noescape
func FuncMessageEventInitMessageEvent6(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent6
//go:noescape
func CallMessageEventInitMessageEvent6(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref)

//go:wasmimport plat/js/web try_MessageEvent_InitMessageEvent6
//go:noescape
func TryMessageEventInitMessageEvent6(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MessageEvent_InitMessageEvent7
//go:noescape
func HasFuncMessageEventInitMessageEvent7(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MessageEvent_InitMessageEvent7
//go:noescape
func FuncMessageEventInitMessageEvent7(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MessageEvent_InitMessageEvent7
//go:noescape
func CallMessageEventInitMessageEvent7(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref)

//go:wasmimport plat/js/web try_MessageEvent_InitMessageEvent7
//go:noescape
func TryMessageEventInitMessageEvent7(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_ScreenY
//go:noescape
func GetMouseEventScreenY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_ClientX
//go:noescape
func GetMouseEventClientX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_ClientY
//go:noescape
func GetMouseEventClientY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_CtrlKey
//go:noescape
func GetMouseEventCtrlKey(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_ShiftKey
//go:noescape
func GetMouseEventShiftKey(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_AltKey
//go:noescape
func GetMouseEventAltKey(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_MetaKey
//go:noescape
func GetMouseEventMetaKey(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_Button
//go:noescape
func GetMouseEventButton(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_Buttons
//go:noescape
func GetMouseEventButtons(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_RelatedTarget
//go:noescape
func GetMouseEventRelatedTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_MovementX
//go:noescape
func GetMouseEventMovementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_MovementY
//go:noescape
func GetMouseEventMovementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_PageX
//go:noescape
func GetMouseEventPageX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_PageY
//go:noescape
func GetMouseEventPageY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_X
//go:noescape
func GetMouseEventX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_Y
//go:noescape
func GetMouseEventY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_OffsetX
//go:noescape
func GetMouseEventOffsetX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MouseEvent_OffsetY
//go:noescape
func GetMouseEventOffsetY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_GetModifierState
//go:noescape
func HasFuncMouseEventGetModifierState(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_GetModifierState
//go:noescape
func FuncMouseEventGetModifierState(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_GetModifierState
//go:noescape
func CallMouseEventGetModifierState(
	this js.Ref, retPtr unsafe.Pointer,
	keyArg js.Ref)

//go:wasmimport plat/js/web try_MouseEvent_GetModifierState
//go:noescape
func TryMouseEventGetModifierState(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	keyArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_InitMouseEvent
//go:noescape
func HasFuncMouseEventInitMouseEvent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent
//go:noescape
func FuncMouseEventInitMouseEvent(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent
//go:noescape
func CallMouseEventInitMouseEvent(
	this js.Ref, retPtr unsafe.Pointer,
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
	relatedTargetArg js.Ref)

//go:wasmimport plat/js/web try_MouseEvent_InitMouseEvent
//go:noescape
func TryMouseEventInitMouseEvent(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
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
	relatedTargetArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_InitMouseEvent1
//go:noescape
func HasFuncMouseEventInitMouseEvent1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent1
//go:noescape
func FuncMouseEventInitMouseEvent1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent1
//go:noescape
func CallMouseEventInitMouseEvent1(
	this js.Ref, retPtr unsafe.Pointer,
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
	buttonArg int32)

//go:wasmimport plat/js/web try_MouseEvent_InitMouseEvent1
//go:noescape
func TryMouseEventInitMouseEvent1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
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
	buttonArg int32) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_InitMouseEvent2
//go:noescape
func HasFuncMouseEventInitMouseEvent2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent2
//go:noescape
func FuncMouseEventInitMouseEvent2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent2
//go:noescape
func CallMouseEventInitMouseEvent2(
	this js.Ref, retPtr unsafe.Pointer,
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
	metaKeyArg js.Ref)

//go:wasmimport plat/js/web try_MouseEvent_InitMouseEvent2
//go:noescape
func TryMouseEventInitMouseEvent2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
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
	metaKeyArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_InitMouseEvent3
//go:noescape
func HasFuncMouseEventInitMouseEvent3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent3
//go:noescape
func FuncMouseEventInitMouseEvent3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent3
//go:noescape
func CallMouseEventInitMouseEvent3(
	this js.Ref, retPtr unsafe.Pointer,
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
	shiftKeyArg js.Ref)

//go:wasmimport plat/js/web try_MouseEvent_InitMouseEvent3
//go:noescape
func TryMouseEventInitMouseEvent3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
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
	shiftKeyArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_InitMouseEvent4
//go:noescape
func HasFuncMouseEventInitMouseEvent4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent4
//go:noescape
func FuncMouseEventInitMouseEvent4(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent4
//go:noescape
func CallMouseEventInitMouseEvent4(
	this js.Ref, retPtr unsafe.Pointer,
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
	altKeyArg js.Ref)

//go:wasmimport plat/js/web try_MouseEvent_InitMouseEvent4
//go:noescape
func TryMouseEventInitMouseEvent4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
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
	altKeyArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_InitMouseEvent5
//go:noescape
func HasFuncMouseEventInitMouseEvent5(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent5
//go:noescape
func FuncMouseEventInitMouseEvent5(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent5
//go:noescape
func CallMouseEventInitMouseEvent5(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32,
	clientXArg int32,
	clientYArg int32,
	ctrlKeyArg js.Ref)

//go:wasmimport plat/js/web try_MouseEvent_InitMouseEvent5
//go:noescape
func TryMouseEventInitMouseEvent5(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32,
	clientXArg int32,
	clientYArg int32,
	ctrlKeyArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_InitMouseEvent6
//go:noescape
func HasFuncMouseEventInitMouseEvent6(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent6
//go:noescape
func FuncMouseEventInitMouseEvent6(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent6
//go:noescape
func CallMouseEventInitMouseEvent6(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32,
	clientXArg int32,
	clientYArg int32)

//go:wasmimport plat/js/web try_MouseEvent_InitMouseEvent6
//go:noescape
func TryMouseEventInitMouseEvent6(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32,
	clientXArg int32,
	clientYArg int32) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_InitMouseEvent7
//go:noescape
func HasFuncMouseEventInitMouseEvent7(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent7
//go:noescape
func FuncMouseEventInitMouseEvent7(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent7
//go:noescape
func CallMouseEventInitMouseEvent7(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32,
	clientXArg int32)

//go:wasmimport plat/js/web try_MouseEvent_InitMouseEvent7
//go:noescape
func TryMouseEventInitMouseEvent7(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32,
	clientXArg int32) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_InitMouseEvent8
//go:noescape
func HasFuncMouseEventInitMouseEvent8(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent8
//go:noescape
func FuncMouseEventInitMouseEvent8(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent8
//go:noescape
func CallMouseEventInitMouseEvent8(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32)

//go:wasmimport plat/js/web try_MouseEvent_InitMouseEvent8
//go:noescape
func TryMouseEventInitMouseEvent8(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32,
	screenYArg int32) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_InitMouseEvent9
//go:noescape
func HasFuncMouseEventInitMouseEvent9(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent9
//go:noescape
func FuncMouseEventInitMouseEvent9(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent9
//go:noescape
func CallMouseEventInitMouseEvent9(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32)

//go:wasmimport plat/js/web try_MouseEvent_InitMouseEvent9
//go:noescape
func TryMouseEventInitMouseEvent9(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32,
	screenXArg int32) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_InitMouseEvent10
//go:noescape
func HasFuncMouseEventInitMouseEvent10(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent10
//go:noescape
func FuncMouseEventInitMouseEvent10(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent10
//go:noescape
func CallMouseEventInitMouseEvent10(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32)

//go:wasmimport plat/js/web try_MouseEvent_InitMouseEvent10
//go:noescape
func TryMouseEventInitMouseEvent10(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_InitMouseEvent11
//go:noescape
func HasFuncMouseEventInitMouseEvent11(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent11
//go:noescape
func FuncMouseEventInitMouseEvent11(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent11
//go:noescape
func CallMouseEventInitMouseEvent11(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref)

//go:wasmimport plat/js/web try_MouseEvent_InitMouseEvent11
//go:noescape
func TryMouseEventInitMouseEvent11(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_InitMouseEvent12
//go:noescape
func HasFuncMouseEventInitMouseEvent12(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent12
//go:noescape
func FuncMouseEventInitMouseEvent12(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent12
//go:noescape
func CallMouseEventInitMouseEvent12(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref)

//go:wasmimport plat/js/web try_MouseEvent_InitMouseEvent12
//go:noescape
func TryMouseEventInitMouseEvent12(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_InitMouseEvent13
//go:noescape
func HasFuncMouseEventInitMouseEvent13(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent13
//go:noescape
func FuncMouseEventInitMouseEvent13(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent13
//go:noescape
func CallMouseEventInitMouseEvent13(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref)

//go:wasmimport plat/js/web try_MouseEvent_InitMouseEvent13
//go:noescape
func TryMouseEventInitMouseEvent13(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MouseEvent_InitMouseEvent14
//go:noescape
func HasFuncMouseEventInitMouseEvent14(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MouseEvent_InitMouseEvent14
//go:noescape
func FuncMouseEventInitMouseEvent14(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MouseEvent_InitMouseEvent14
//go:noescape
func CallMouseEventInitMouseEvent14(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref)

//go:wasmimport plat/js/web try_MouseEvent_InitMouseEvent14
//go:noescape
func TryMouseEventInitMouseEvent14(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_MutationRecord_Type
//go:noescape
func GetMutationRecordType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MutationRecord_Target
//go:noescape
func GetMutationRecordTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MutationRecord_AddedNodes
//go:noescape
func GetMutationRecordAddedNodes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MutationRecord_RemovedNodes
//go:noescape
func GetMutationRecordRemovedNodes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MutationRecord_PreviousSibling
//go:noescape
func GetMutationRecordPreviousSibling(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MutationRecord_NextSibling
//go:noescape
func GetMutationRecordNextSibling(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MutationRecord_AttributeName
//go:noescape
func GetMutationRecordAttributeName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MutationRecord_AttributeNamespace
//go:noescape
func GetMutationRecordAttributeNamespace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MutationRecord_OldValue
//go:noescape
func GetMutationRecordOldValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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

//go:wasmimport plat/js/web has_MutationObserver_Observe
//go:noescape
func HasFuncMutationObserverObserve(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MutationObserver_Observe
//go:noescape
func FuncMutationObserverObserve(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MutationObserver_Observe
//go:noescape
func CallMutationObserverObserve(
	this js.Ref, retPtr unsafe.Pointer,
	target js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MutationObserver_Observe
//go:noescape
func TryMutationObserverObserve(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	target js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MutationObserver_Observe1
//go:noescape
func HasFuncMutationObserverObserve1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MutationObserver_Observe1
//go:noescape
func FuncMutationObserverObserve1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MutationObserver_Observe1
//go:noescape
func CallMutationObserverObserve1(
	this js.Ref, retPtr unsafe.Pointer,
	target js.Ref)

//go:wasmimport plat/js/web try_MutationObserver_Observe1
//go:noescape
func TryMutationObserverObserve1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	target js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MutationObserver_Disconnect
//go:noescape
func HasFuncMutationObserverDisconnect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MutationObserver_Disconnect
//go:noescape
func FuncMutationObserverDisconnect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MutationObserver_Disconnect
//go:noescape
func CallMutationObserverDisconnect(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MutationObserver_Disconnect
//go:noescape
func TryMutationObserverDisconnect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MutationObserver_TakeRecords
//go:noescape
func HasFuncMutationObserverTakeRecords(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MutationObserver_TakeRecords
//go:noescape
func FuncMutationObserverTakeRecords(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MutationObserver_TakeRecords
//go:noescape
func CallMutationObserverTakeRecords(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MutationObserver_TakeRecords
//go:noescape
func TryMutationObserverTakeRecords(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MutationEvent_PrevValue
//go:noescape
func GetMutationEventPrevValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MutationEvent_NewValue
//go:noescape
func GetMutationEventNewValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MutationEvent_AttrName
//go:noescape
func GetMutationEventAttrName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MutationEvent_AttrChange
//go:noescape
func GetMutationEventAttrChange(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MutationEvent_InitMutationEvent
//go:noescape
func HasFuncMutationEventInitMutationEvent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent
//go:noescape
func FuncMutationEventInitMutationEvent(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent
//go:noescape
func CallMutationEventInitMutationEvent(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	relatedNodeArg js.Ref,
	prevValueArg js.Ref,
	newValueArg js.Ref,
	attrNameArg js.Ref,
	attrChangeArg uint32)

//go:wasmimport plat/js/web try_MutationEvent_InitMutationEvent
//go:noescape
func TryMutationEventInitMutationEvent(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	relatedNodeArg js.Ref,
	prevValueArg js.Ref,
	newValueArg js.Ref,
	attrNameArg js.Ref,
	attrChangeArg uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_MutationEvent_InitMutationEvent1
//go:noescape
func HasFuncMutationEventInitMutationEvent1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent1
//go:noescape
func FuncMutationEventInitMutationEvent1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent1
//go:noescape
func CallMutationEventInitMutationEvent1(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	relatedNodeArg js.Ref,
	prevValueArg js.Ref,
	newValueArg js.Ref,
	attrNameArg js.Ref)

//go:wasmimport plat/js/web try_MutationEvent_InitMutationEvent1
//go:noescape
func TryMutationEventInitMutationEvent1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	relatedNodeArg js.Ref,
	prevValueArg js.Ref,
	newValueArg js.Ref,
	attrNameArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MutationEvent_InitMutationEvent2
//go:noescape
func HasFuncMutationEventInitMutationEvent2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent2
//go:noescape
func FuncMutationEventInitMutationEvent2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent2
//go:noescape
func CallMutationEventInitMutationEvent2(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	relatedNodeArg js.Ref,
	prevValueArg js.Ref,
	newValueArg js.Ref)

//go:wasmimport plat/js/web try_MutationEvent_InitMutationEvent2
//go:noescape
func TryMutationEventInitMutationEvent2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	relatedNodeArg js.Ref,
	prevValueArg js.Ref,
	newValueArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MutationEvent_InitMutationEvent3
//go:noescape
func HasFuncMutationEventInitMutationEvent3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent3
//go:noescape
func FuncMutationEventInitMutationEvent3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent3
//go:noescape
func CallMutationEventInitMutationEvent3(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	relatedNodeArg js.Ref,
	prevValueArg js.Ref)

//go:wasmimport plat/js/web try_MutationEvent_InitMutationEvent3
//go:noescape
func TryMutationEventInitMutationEvent3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	relatedNodeArg js.Ref,
	prevValueArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MutationEvent_InitMutationEvent4
//go:noescape
func HasFuncMutationEventInitMutationEvent4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent4
//go:noescape
func FuncMutationEventInitMutationEvent4(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent4
//go:noescape
func CallMutationEventInitMutationEvent4(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	relatedNodeArg js.Ref)

//go:wasmimport plat/js/web try_MutationEvent_InitMutationEvent4
//go:noescape
func TryMutationEventInitMutationEvent4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	relatedNodeArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MutationEvent_InitMutationEvent5
//go:noescape
func HasFuncMutationEventInitMutationEvent5(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent5
//go:noescape
func FuncMutationEventInitMutationEvent5(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent5
//go:noescape
func CallMutationEventInitMutationEvent5(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref)

//go:wasmimport plat/js/web try_MutationEvent_InitMutationEvent5
//go:noescape
func TryMutationEventInitMutationEvent5(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MutationEvent_InitMutationEvent6
//go:noescape
func HasFuncMutationEventInitMutationEvent6(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent6
//go:noescape
func FuncMutationEventInitMutationEvent6(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent6
//go:noescape
func CallMutationEventInitMutationEvent6(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref)

//go:wasmimport plat/js/web try_MutationEvent_InitMutationEvent6
//go:noescape
func TryMutationEventInitMutationEvent6(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MutationEvent_InitMutationEvent7
//go:noescape
func HasFuncMutationEventInitMutationEvent7(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MutationEvent_InitMutationEvent7
//go:noescape
func FuncMutationEventInitMutationEvent7(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MutationEvent_InitMutationEvent7
//go:noescape
func CallMutationEventInitMutationEvent7(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref)

//go:wasmimport plat/js/web try_MutationEvent_InitMutationEvent7
//go:noescape
func TryMutationEventInitMutationEvent7(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NDEFRecord_MediaType
//go:noescape
func GetNDEFRecordMediaType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NDEFRecord_Id
//go:noescape
func GetNDEFRecordId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NDEFRecord_Data
//go:noescape
func GetNDEFRecordData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NDEFRecord_Encoding
//go:noescape
func GetNDEFRecordEncoding(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NDEFRecord_Lang
//go:noescape
func GetNDEFRecordLang(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_NDEFRecord_ToRecords
//go:noescape
func HasFuncNDEFRecordToRecords(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NDEFRecord_ToRecords
//go:noescape
func FuncNDEFRecordToRecords(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_NDEFRecord_ToRecords
//go:noescape
func CallNDEFRecordToRecords(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_NDEFRecord_ToRecords
//go:noescape
func TryNDEFRecordToRecords(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_NDEFMessage_NDEFMessage
//go:noescape
func NewNDEFMessageByNDEFMessage(
	messageInit unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_NDEFMessage_Records
//go:noescape
func GetNDEFMessageRecords(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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

//go:wasmimport plat/js/web has_NDEFReader_Scan
//go:noescape
func HasFuncNDEFReaderScan(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NDEFReader_Scan
//go:noescape
func FuncNDEFReaderScan(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_NDEFReader_Scan
//go:noescape
func CallNDEFReaderScan(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_NDEFReader_Scan
//go:noescape
func TryNDEFReaderScan(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_NDEFReader_Scan1
//go:noescape
func HasFuncNDEFReaderScan1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NDEFReader_Scan1
//go:noescape
func FuncNDEFReaderScan1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_NDEFReader_Scan1
//go:noescape
func CallNDEFReaderScan1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_NDEFReader_Scan1
//go:noescape
func TryNDEFReaderScan1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_NDEFReader_Write
//go:noescape
func HasFuncNDEFReaderWrite(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NDEFReader_Write
//go:noescape
func FuncNDEFReaderWrite(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_NDEFReader_Write
//go:noescape
func CallNDEFReaderWrite(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_NDEFReader_Write
//go:noescape
func TryNDEFReaderWrite(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_NDEFReader_Write1
//go:noescape
func HasFuncNDEFReaderWrite1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NDEFReader_Write1
//go:noescape
func FuncNDEFReaderWrite1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_NDEFReader_Write1
//go:noescape
func CallNDEFReaderWrite1(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/web try_NDEFReader_Write1
//go:noescape
func TryNDEFReaderWrite1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_NDEFReader_MakeReadOnly
//go:noescape
func HasFuncNDEFReaderMakeReadOnly(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NDEFReader_MakeReadOnly
//go:noescape
func FuncNDEFReaderMakeReadOnly(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_NDEFReader_MakeReadOnly
//go:noescape
func CallNDEFReaderMakeReadOnly(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_NDEFReader_MakeReadOnly
//go:noescape
func TryNDEFReaderMakeReadOnly(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_NDEFReader_MakeReadOnly1
//go:noescape
func HasFuncNDEFReaderMakeReadOnly1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NDEFReader_MakeReadOnly1
//go:noescape
func FuncNDEFReaderMakeReadOnly1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_NDEFReader_MakeReadOnly1
//go:noescape
func CallNDEFReaderMakeReadOnly1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_NDEFReader_MakeReadOnly1
//go:noescape
func TryNDEFReaderMakeReadOnly1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NDEFReadingEvent_Message
//go:noescape
func GetNDEFReadingEventMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NamedFlow_Name
//go:noescape
func GetNamedFlowName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NamedFlow_Overset
//go:noescape
func GetNamedFlowOverset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NamedFlow_FirstEmptyRegionIndex
//go:noescape
func GetNamedFlowFirstEmptyRegionIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_NamedFlow_GetRegions
//go:noescape
func HasFuncNamedFlowGetRegions(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NamedFlow_GetRegions
//go:noescape
func FuncNamedFlowGetRegions(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_NamedFlow_GetRegions
//go:noescape
func CallNamedFlowGetRegions(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_NamedFlow_GetRegions
//go:noescape
func TryNamedFlowGetRegions(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_NamedFlow_GetContent
//go:noescape
func HasFuncNamedFlowGetContent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NamedFlow_GetContent
//go:noescape
func FuncNamedFlowGetContent(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_NamedFlow_GetContent
//go:noescape
func CallNamedFlowGetContent(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_NamedFlow_GetContent
//go:noescape
func TryNamedFlowGetContent(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_NamedFlow_GetRegionsByContent
//go:noescape
func HasFuncNamedFlowGetRegionsByContent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NamedFlow_GetRegionsByContent
//go:noescape
func FuncNamedFlowGetRegionsByContent(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_NamedFlow_GetRegionsByContent
//go:noescape
func CallNamedFlowGetRegionsByContent(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref)

//go:wasmimport plat/js/web try_NamedFlow_GetRegionsByContent
//go:noescape
func TryNamedFlowGetRegionsByContent(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_NavigationDestination_Url
//go:noescape
func GetNavigationDestinationUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NavigationDestination_Key
//go:noescape
func GetNavigationDestinationKey(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NavigationDestination_Id
//go:noescape
func GetNavigationDestinationId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NavigationDestination_Index
//go:noescape
func GetNavigationDestinationIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NavigationDestination_SameDocument
//go:noescape
func GetNavigationDestinationSameDocument(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_NavigationDestination_GetState
//go:noescape
func HasFuncNavigationDestinationGetState(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NavigationDestination_GetState
//go:noescape
func FuncNavigationDestinationGetState(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_NavigationDestination_GetState
//go:noescape
func CallNavigationDestinationGetState(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_NavigationDestination_GetState
//go:noescape
func TryNavigationDestinationGetState(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_NavigateEventInit
//go:noescape
func NavigateEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NavigateEventInit
//go:noescape
func NavigateEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
