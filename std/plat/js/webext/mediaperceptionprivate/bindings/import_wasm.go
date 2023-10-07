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

//go:wasmimport plat/js/webext/mediaperceptionprivate store_AudioSpectrogram
//go:noescape
func AudioSpectrogramJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_AudioSpectrogram
//go:noescape
func AudioSpectrogramJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_AudioHumanPresenceDetection
//go:noescape
func AudioHumanPresenceDetectionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_AudioHumanPresenceDetection
//go:noescape
func AudioHumanPresenceDetectionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_AudioLocalization
//go:noescape
func AudioLocalizationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_AudioLocalization
//go:noescape
func AudioLocalizationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate constof_HotwordType
//go:noescape
func ConstOfHotwordType(str js.Ref) uint32

//go:wasmimport plat/js/webext/mediaperceptionprivate store_Hotword
//go:noescape
func HotwordJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_Hotword
//go:noescape
func HotwordJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_HotwordDetection
//go:noescape
func HotwordDetectionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_HotwordDetection
//go:noescape
func HotwordDetectionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_AudioPerception
//go:noescape
func AudioPerceptionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_AudioPerception
//go:noescape
func AudioPerceptionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_AudioVisualHumanPresenceDetection
//go:noescape
func AudioVisualHumanPresenceDetectionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_AudioVisualHumanPresenceDetection
//go:noescape
func AudioVisualHumanPresenceDetectionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_AudioVisualPerception
//go:noescape
func AudioVisualPerceptionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_AudioVisualPerception
//go:noescape
func AudioVisualPerceptionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_Point
//go:noescape
func PointJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_Point
//go:noescape
func PointJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_BoundingBox
//go:noescape
func BoundingBoxJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_BoundingBox
//go:noescape
func BoundingBoxJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate constof_ComponentType
//go:noescape
func ConstOfComponentType(str js.Ref) uint32

//go:wasmimport plat/js/webext/mediaperceptionprivate store_Component
//go:noescape
func ComponentJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_Component
//go:noescape
func ComponentJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate constof_ComponentInstallationError
//go:noescape
func ConstOfComponentInstallationError(str js.Ref) uint32

//go:wasmimport plat/js/webext/mediaperceptionprivate constof_ComponentStatus
//go:noescape
func ConstOfComponentStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/mediaperceptionprivate store_ComponentState
//go:noescape
func ComponentStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_ComponentState
//go:noescape
func ComponentStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate constof_ServiceError
//go:noescape
func ConstOfServiceError(str js.Ref) uint32

//go:wasmimport plat/js/webext/mediaperceptionprivate constof_EntityType
//go:noescape
func ConstOfEntityType(str js.Ref) uint32

//go:wasmimport plat/js/webext/mediaperceptionprivate constof_DistanceUnits
//go:noescape
func ConstOfDistanceUnits(str js.Ref) uint32

//go:wasmimport plat/js/webext/mediaperceptionprivate store_Distance
//go:noescape
func DistanceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_Distance
//go:noescape
func DistanceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_Entity
//go:noescape
func EntityJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_Entity
//go:noescape
func EntityJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_PacketLatency
//go:noescape
func PacketLatencyJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_PacketLatency
//go:noescape
func PacketLatencyJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate constof_LightCondition
//go:noescape
func ConstOfLightCondition(str js.Ref) uint32

//go:wasmimport plat/js/webext/mediaperceptionprivate store_VideoHumanPresenceDetection
//go:noescape
func VideoHumanPresenceDetectionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_VideoHumanPresenceDetection
//go:noescape
func VideoHumanPresenceDetectionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate constof_FramePerceptionType
//go:noescape
func ConstOfFramePerceptionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/mediaperceptionprivate store_FramePerception
//go:noescape
func FramePerceptionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_FramePerception
//go:noescape
func FramePerceptionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate constof_ImageFormat
//go:noescape
func ConstOfImageFormat(str js.Ref) uint32

//go:wasmimport plat/js/webext/mediaperceptionprivate store_ImageFrame
//go:noescape
func ImageFrameJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_ImageFrame
//go:noescape
func ImageFrameJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_Metadata
//go:noescape
func MetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_Metadata
//go:noescape
func MetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_PerceptionSample
//go:noescape
func PerceptionSampleJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_PerceptionSample
//go:noescape
func PerceptionSampleJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_Diagnostics
//go:noescape
func DiagnosticsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_Diagnostics
//go:noescape
func DiagnosticsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate constof_Feature
//go:noescape
func ConstOfFeature(str js.Ref) uint32

//go:wasmimport plat/js/webext/mediaperceptionprivate store_MediaPerception
//go:noescape
func MediaPerceptionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_MediaPerception
//go:noescape
func MediaPerceptionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_NamedTemplateArgument
//go:noescape
func NamedTemplateArgumentJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_NamedTemplateArgument
//go:noescape
func NamedTemplateArgumentJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate constof_ProcessStatus
//go:noescape
func ConstOfProcessStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/mediaperceptionprivate store_ProcessState
//go:noescape
func ProcessStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_ProcessState
//go:noescape
func ProcessStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate constof_Status
//go:noescape
func ConstOfStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/mediaperceptionprivate store_VideoStreamParam
//go:noescape
func VideoStreamParamJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_VideoStreamParam
//go:noescape
func VideoStreamParamJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_Whiteboard
//go:noescape
func WhiteboardJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_Whiteboard
//go:noescape
func WhiteboardJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate store_State
//go:noescape
func StateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate load_State
//go:noescape
func StateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate has_GetDiagnostics
//go:noescape
func HasFuncGetDiagnostics() js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate func_GetDiagnostics
//go:noescape
func FuncGetDiagnostics(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaperceptionprivate call_GetDiagnostics
//go:noescape
func CallGetDiagnostics(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaperceptionprivate try_GetDiagnostics
//go:noescape
func TryGetDiagnostics(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate has_GetState
//go:noescape
func HasFuncGetState() js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate func_GetState
//go:noescape
func FuncGetState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaperceptionprivate call_GetState
//go:noescape
func CallGetState(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaperceptionprivate try_GetState
//go:noescape
func TryGetState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate has_OnMediaPerception
//go:noescape
func HasFuncOnMediaPerception() js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate func_OnMediaPerception
//go:noescape
func FuncOnMediaPerception(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaperceptionprivate call_OnMediaPerception
//go:noescape
func CallOnMediaPerception(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate try_OnMediaPerception
//go:noescape
func TryOnMediaPerception(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate has_OffMediaPerception
//go:noescape
func HasFuncOffMediaPerception() js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate func_OffMediaPerception
//go:noescape
func FuncOffMediaPerception(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaperceptionprivate call_OffMediaPerception
//go:noescape
func CallOffMediaPerception(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate try_OffMediaPerception
//go:noescape
func TryOffMediaPerception(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate has_HasOnMediaPerception
//go:noescape
func HasFuncHasOnMediaPerception() js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate func_HasOnMediaPerception
//go:noescape
func FuncHasOnMediaPerception(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaperceptionprivate call_HasOnMediaPerception
//go:noescape
func CallHasOnMediaPerception(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate try_HasOnMediaPerception
//go:noescape
func TryHasOnMediaPerception(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate has_SetAnalyticsComponent
//go:noescape
func HasFuncSetAnalyticsComponent() js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate func_SetAnalyticsComponent
//go:noescape
func FuncSetAnalyticsComponent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaperceptionprivate call_SetAnalyticsComponent
//go:noescape
func CallSetAnalyticsComponent(
	retPtr unsafe.Pointer,
	component unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaperceptionprivate try_SetAnalyticsComponent
//go:noescape
func TrySetAnalyticsComponent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	component unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate has_SetComponentProcessState
//go:noescape
func HasFuncSetComponentProcessState() js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate func_SetComponentProcessState
//go:noescape
func FuncSetComponentProcessState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaperceptionprivate call_SetComponentProcessState
//go:noescape
func CallSetComponentProcessState(
	retPtr unsafe.Pointer,
	processState unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaperceptionprivate try_SetComponentProcessState
//go:noescape
func TrySetComponentProcessState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	processState unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/mediaperceptionprivate has_SetState
//go:noescape
func HasFuncSetState() js.Ref

//go:wasmimport plat/js/webext/mediaperceptionprivate func_SetState
//go:noescape
func FuncSetState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaperceptionprivate call_SetState
//go:noescape
func CallSetState(
	retPtr unsafe.Pointer,
	state unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaperceptionprivate try_SetState
//go:noescape
func TrySetState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	state unsafe.Pointer) (ok js.Ref)
