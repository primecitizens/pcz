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

//go:wasmimport plat/js/web get_LargestContentfulPaint_RenderTime
//go:noescape
func GetLargestContentfulPaintRenderTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LargestContentfulPaint_LoadTime
//go:noescape
func GetLargestContentfulPaintLoadTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LargestContentfulPaint_Size
//go:noescape
func GetLargestContentfulPaintSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LargestContentfulPaint_Id
//go:noescape
func GetLargestContentfulPaintId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LargestContentfulPaint_Url
//go:noescape
func GetLargestContentfulPaintUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LargestContentfulPaint_Element
//go:noescape
func GetLargestContentfulPaintElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_LargestContentfulPaint_ToJSON
//go:noescape
func HasLargestContentfulPaintToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_LargestContentfulPaint_ToJSON
//go:noescape
func LargestContentfulPaintToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_LargestContentfulPaint_ToJSON
//go:noescape
func CallLargestContentfulPaintToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_LargestContentfulPaint_ToJSON
//go:noescape
func TryLargestContentfulPaintToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_LatencyMode
//go:noescape
func ConstOfLatencyMode(str js.Ref) uint32

//go:wasmimport plat/js/web get_LayoutConstraints_AvailableInlineSize
//go:noescape
func GetLayoutConstraintsAvailableInlineSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutConstraints_AvailableBlockSize
//go:noescape
func GetLayoutConstraintsAvailableBlockSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutConstraints_FixedInlineSize
//go:noescape
func GetLayoutConstraintsFixedInlineSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutConstraints_FixedBlockSize
//go:noescape
func GetLayoutConstraintsFixedBlockSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutConstraints_PercentageInlineSize
//go:noescape
func GetLayoutConstraintsPercentageInlineSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutConstraints_PercentageBlockSize
//go:noescape
func GetLayoutConstraintsPercentageBlockSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutConstraints_BlockFragmentationOffset
//go:noescape
func GetLayoutConstraintsBlockFragmentationOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutConstraints_BlockFragmentationType
//go:noescape
func GetLayoutConstraintsBlockFragmentationType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutConstraints_Data
//go:noescape
func GetLayoutConstraintsData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutEdges_InlineStart
//go:noescape
func GetLayoutEdgesInlineStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutEdges_InlineEnd
//go:noescape
func GetLayoutEdgesInlineEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutEdges_BlockStart
//go:noescape
func GetLayoutEdgesBlockStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutEdges_BlockEnd
//go:noescape
func GetLayoutEdgesBlockEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutEdges_Inline
//go:noescape
func GetLayoutEdgesInline(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutEdges_Block
//go:noescape
func GetLayoutEdgesBlock(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_LayoutSizingMode
//go:noescape
func ConstOfLayoutSizingMode(str js.Ref) uint32

//go:wasmimport plat/js/web store_LayoutOptions
//go:noescape
func LayoutOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_LayoutOptions
//go:noescape
func LayoutOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_LayoutShiftAttribution_Node
//go:noescape
func GetLayoutShiftAttributionNode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutShiftAttribution_PreviousRect
//go:noescape
func GetLayoutShiftAttributionPreviousRect(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutShiftAttribution_CurrentRect
//go:noescape
func GetLayoutShiftAttributionCurrentRect(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutShift_Value
//go:noescape
func GetLayoutShiftValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutShift_HadRecentInput
//go:noescape
func GetLayoutShiftHadRecentInput(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutShift_LastInputTime
//go:noescape
func GetLayoutShiftLastInputTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutShift_Sources
//go:noescape
func GetLayoutShiftSources(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_LayoutShift_ToJSON
//go:noescape
func HasLayoutShiftToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_LayoutShift_ToJSON
//go:noescape
func LayoutShiftToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_LayoutShift_ToJSON
//go:noescape
func CallLayoutShiftToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_LayoutShift_ToJSON
//go:noescape
func TryLayoutShiftToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_LayoutWorkletGlobalScope_RegisterLayout
//go:noescape
func HasLayoutWorkletGlobalScopeRegisterLayout(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_LayoutWorkletGlobalScope_RegisterLayout
//go:noescape
func LayoutWorkletGlobalScopeRegisterLayoutFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_LayoutWorkletGlobalScope_RegisterLayout
//go:noescape
func CallLayoutWorkletGlobalScopeRegisterLayout(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	layoutCtor js.Ref)

//go:wasmimport plat/js/web try_LayoutWorkletGlobalScope_RegisterLayout
//go:noescape
func TryLayoutWorkletGlobalScopeRegisterLayout(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	layoutCtor js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_LineAlignSetting
//go:noescape
func ConstOfLineAlignSetting(str js.Ref) uint32

//go:wasmimport plat/js/web store_LinearAccelerationReadingValues
//go:noescape
func LinearAccelerationReadingValuesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_LinearAccelerationReadingValues
//go:noescape
func LinearAccelerationReadingValuesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_LinearAccelerationSensor_LinearAccelerationSensor
//go:noescape
func NewLinearAccelerationSensorByLinearAccelerationSensor(
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_LinearAccelerationSensor_LinearAccelerationSensor1
//go:noescape
func NewLinearAccelerationSensorByLinearAccelerationSensor1() js.Ref

//go:wasmimport plat/js/web constof_MIDIPortType
//go:noescape
func ConstOfMIDIPortType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_MIDIPortDeviceState
//go:noescape
func ConstOfMIDIPortDeviceState(str js.Ref) uint32

//go:wasmimport plat/js/web constof_MIDIPortConnectionState
//go:noescape
func ConstOfMIDIPortConnectionState(str js.Ref) uint32

//go:wasmimport plat/js/web get_MIDIPort_Id
//go:noescape
func GetMIDIPortId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MIDIPort_Manufacturer
//go:noescape
func GetMIDIPortManufacturer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MIDIPort_Name
//go:noescape
func GetMIDIPortName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MIDIPort_Type
//go:noescape
func GetMIDIPortType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MIDIPort_Version
//go:noescape
func GetMIDIPortVersion(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MIDIPort_State
//go:noescape
func GetMIDIPortState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MIDIPort_Connection
//go:noescape
func GetMIDIPortConnection(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MIDIPort_Open
//go:noescape
func HasMIDIPortOpen(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MIDIPort_Open
//go:noescape
func MIDIPortOpenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MIDIPort_Open
//go:noescape
func CallMIDIPortOpen(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MIDIPort_Open
//go:noescape
func TryMIDIPortOpen(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MIDIPort_Close
//go:noescape
func HasMIDIPortClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MIDIPort_Close
//go:noescape
func MIDIPortCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MIDIPort_Close
//go:noescape
func CallMIDIPortClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MIDIPort_Close
//go:noescape
func TryMIDIPortClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_MIDIConnectionEventInit
//go:noescape
func MIDIConnectionEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MIDIConnectionEventInit
//go:noescape
func MIDIConnectionEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MIDIConnectionEvent_MIDIConnectionEvent
//go:noescape
func NewMIDIConnectionEventByMIDIConnectionEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_MIDIConnectionEvent_MIDIConnectionEvent1
//go:noescape
func NewMIDIConnectionEventByMIDIConnectionEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_MIDIConnectionEvent_Port
//go:noescape
func GetMIDIConnectionEventPort(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_MIDIMessageEventInit
//go:noescape
func MIDIMessageEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MIDIMessageEventInit
//go:noescape
func MIDIMessageEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MIDIMessageEvent_MIDIMessageEvent
//go:noescape
func NewMIDIMessageEventByMIDIMessageEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_MIDIMessageEvent_MIDIMessageEvent1
//go:noescape
func NewMIDIMessageEventByMIDIMessageEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_MIDIMessageEvent_Data
//go:noescape
func GetMIDIMessageEventData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MIDIOutput_Send
//go:noescape
func HasMIDIOutputSend(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MIDIOutput_Send
//go:noescape
func MIDIOutputSendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MIDIOutput_Send
//go:noescape
func CallMIDIOutputSend(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref,
	timestamp float64)

//go:wasmimport plat/js/web try_MIDIOutput_Send
//go:noescape
func TryMIDIOutputSend(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref,
	timestamp float64) (ok js.Ref)

//go:wasmimport plat/js/web has_MIDIOutput_Send1
//go:noescape
func HasMIDIOutputSend1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MIDIOutput_Send1
//go:noescape
func MIDIOutputSend1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MIDIOutput_Send1
//go:noescape
func CallMIDIOutputSend1(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_MIDIOutput_Send1
//go:noescape
func TryMIDIOutputSend1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MIDIOutput_Clear
//go:noescape
func HasMIDIOutputClear(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MIDIOutput_Clear
//go:noescape
func MIDIOutputClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MIDIOutput_Clear
//go:noescape
func CallMIDIOutputClear(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MIDIOutput_Clear
//go:noescape
func TryMIDIOutputClear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_MLAutoPad
//go:noescape
func ConstOfMLAutoPad(str js.Ref) uint32

//go:wasmimport plat/js/web store_MLBatchNormalizationOptions
//go:noescape
func MLBatchNormalizationOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLBatchNormalizationOptions
//go:noescape
func MLBatchNormalizationOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLBufferResourceView
//go:noescape
func MLBufferResourceViewJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLBufferResourceView
//go:noescape
func MLBufferResourceViewJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLClampOptions
//go:noescape
func MLClampOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLClampOptions
//go:noescape
func MLClampOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MLConv2dFilterOperandLayout
//go:noescape
func ConstOfMLConv2dFilterOperandLayout(str js.Ref) uint32

//go:wasmimport plat/js/web constof_MLInputOperandLayout
//go:noescape
func ConstOfMLInputOperandLayout(str js.Ref) uint32

//go:wasmimport plat/js/web store_MLConv2dOptions
//go:noescape
func MLConv2dOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLConv2dOptions
//go:noescape
func MLConv2dOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MLConvTranspose2dFilterOperandLayout
//go:noescape
func ConstOfMLConvTranspose2dFilterOperandLayout(str js.Ref) uint32

//go:wasmimport plat/js/web store_MLConvTranspose2dOptions
//go:noescape
func MLConvTranspose2dOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLConvTranspose2dOptions
//go:noescape
func MLConvTranspose2dOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLEluOptions
//go:noescape
func MLEluOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLEluOptions
//go:noescape
func MLEluOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLGemmOptions
//go:noescape
func MLGemmOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLGemmOptions
//go:noescape
func MLGemmOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
