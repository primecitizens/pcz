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

//go:wasmimport plat/js/web get_GPURenderPassEncoder_Label
//go:noescape

func GetGPURenderPassEncoderLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_GPURenderPassEncoder_Label
//go:noescape

func SetGPURenderPassEncoderLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetViewport
//go:noescape

func CallGPURenderPassEncoderSetViewport(
	this js.Ref,
	ptr unsafe.Pointer,

	x float32,
	y float32,
	width float32,
	height float32,
	minDepth float32,
	maxDepth float32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetViewport
//go:noescape

func GPURenderPassEncoderSetViewportFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetScissorRect
//go:noescape

func CallGPURenderPassEncoderSetScissorRect(
	this js.Ref,
	ptr unsafe.Pointer,

	x uint32,
	y uint32,
	width uint32,
	height uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetScissorRect
//go:noescape

func GPURenderPassEncoderSetScissorRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetBlendConstant
//go:noescape

func CallGPURenderPassEncoderSetBlendConstant(
	this js.Ref,
	ptr unsafe.Pointer,

	color js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetBlendConstant
//go:noescape

func GPURenderPassEncoderSetBlendConstantFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetStencilReference
//go:noescape

func CallGPURenderPassEncoderSetStencilReference(
	this js.Ref,
	ptr unsafe.Pointer,

	reference uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetStencilReference
//go:noescape

func GPURenderPassEncoderSetStencilReferenceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_BeginOcclusionQuery
//go:noescape

func CallGPURenderPassEncoderBeginOcclusionQuery(
	this js.Ref,
	ptr unsafe.Pointer,

	queryIndex uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_BeginOcclusionQuery
//go:noescape

func GPURenderPassEncoderBeginOcclusionQueryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_EndOcclusionQuery
//go:noescape

func CallGPURenderPassEncoderEndOcclusionQuery(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_EndOcclusionQuery
//go:noescape

func GPURenderPassEncoderEndOcclusionQueryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_ExecuteBundles
//go:noescape

func CallGPURenderPassEncoderExecuteBundles(
	this js.Ref,
	ptr unsafe.Pointer,

	bundles js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_ExecuteBundles
//go:noescape

func GPURenderPassEncoderExecuteBundlesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_End
//go:noescape

func CallGPURenderPassEncoderEnd(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_End
//go:noescape

func GPURenderPassEncoderEndFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetPipeline
//go:noescape

func CallGPURenderPassEncoderSetPipeline(
	this js.Ref,
	ptr unsafe.Pointer,

	pipeline js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetPipeline
//go:noescape

func GPURenderPassEncoderSetPipelineFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetIndexBuffer
//go:noescape

func CallGPURenderPassEncoderSetIndexBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
	indexFormat uint32,
	offset float64,
	size float64,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetIndexBuffer
//go:noescape

func GPURenderPassEncoderSetIndexBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetIndexBuffer1
//go:noescape

func CallGPURenderPassEncoderSetIndexBuffer1(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
	indexFormat uint32,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetIndexBuffer1
//go:noescape

func GPURenderPassEncoderSetIndexBuffer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetIndexBuffer2
//go:noescape

func CallGPURenderPassEncoderSetIndexBuffer2(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
	indexFormat uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetIndexBuffer2
//go:noescape

func GPURenderPassEncoderSetIndexBuffer2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetVertexBuffer
//go:noescape

func CallGPURenderPassEncoderSetVertexBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	slot uint32,
	buffer js.Ref,
	offset float64,
	size float64,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetVertexBuffer
//go:noescape

func GPURenderPassEncoderSetVertexBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetVertexBuffer1
//go:noescape

func CallGPURenderPassEncoderSetVertexBuffer1(
	this js.Ref,
	ptr unsafe.Pointer,

	slot uint32,
	buffer js.Ref,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetVertexBuffer1
//go:noescape

func GPURenderPassEncoderSetVertexBuffer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetVertexBuffer2
//go:noescape

func CallGPURenderPassEncoderSetVertexBuffer2(
	this js.Ref,
	ptr unsafe.Pointer,

	slot uint32,
	buffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetVertexBuffer2
//go:noescape

func GPURenderPassEncoderSetVertexBuffer2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_Draw
//go:noescape

func CallGPURenderPassEncoderDraw(
	this js.Ref,
	ptr unsafe.Pointer,

	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32,
	firstInstance uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_Draw
//go:noescape

func GPURenderPassEncoderDrawFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_Draw1
//go:noescape

func CallGPURenderPassEncoderDraw1(
	this js.Ref,
	ptr unsafe.Pointer,

	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_Draw1
//go:noescape

func GPURenderPassEncoderDraw1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_Draw2
//go:noescape

func CallGPURenderPassEncoderDraw2(
	this js.Ref,
	ptr unsafe.Pointer,

	vertexCount uint32,
	instanceCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_Draw2
//go:noescape

func GPURenderPassEncoderDraw2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_Draw3
//go:noescape

func CallGPURenderPassEncoderDraw3(
	this js.Ref,
	ptr unsafe.Pointer,

	vertexCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_Draw3
//go:noescape

func GPURenderPassEncoderDraw3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_DrawIndexed
//go:noescape

func CallGPURenderPassEncoderDrawIndexed(
	this js.Ref,
	ptr unsafe.Pointer,

	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	baseVertex int32,
	firstInstance uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_DrawIndexed
//go:noescape

func GPURenderPassEncoderDrawIndexedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_DrawIndexed1
//go:noescape

func CallGPURenderPassEncoderDrawIndexed1(
	this js.Ref,
	ptr unsafe.Pointer,

	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	baseVertex int32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_DrawIndexed1
//go:noescape

func GPURenderPassEncoderDrawIndexed1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_DrawIndexed2
//go:noescape

func CallGPURenderPassEncoderDrawIndexed2(
	this js.Ref,
	ptr unsafe.Pointer,

	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_DrawIndexed2
//go:noescape

func GPURenderPassEncoderDrawIndexed2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_DrawIndexed3
//go:noescape

func CallGPURenderPassEncoderDrawIndexed3(
	this js.Ref,
	ptr unsafe.Pointer,

	indexCount uint32,
	instanceCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_DrawIndexed3
//go:noescape

func GPURenderPassEncoderDrawIndexed3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_DrawIndexed4
//go:noescape

func CallGPURenderPassEncoderDrawIndexed4(
	this js.Ref,
	ptr unsafe.Pointer,

	indexCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_DrawIndexed4
//go:noescape

func GPURenderPassEncoderDrawIndexed4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_DrawIndirect
//go:noescape

func CallGPURenderPassEncoderDrawIndirect(
	this js.Ref,
	ptr unsafe.Pointer,

	indirectBuffer js.Ref,
	indirectOffset float64,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_DrawIndirect
//go:noescape

func GPURenderPassEncoderDrawIndirectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_DrawIndexedIndirect
//go:noescape

func CallGPURenderPassEncoderDrawIndexedIndirect(
	this js.Ref,
	ptr unsafe.Pointer,

	indirectBuffer js.Ref,
	indirectOffset float64,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_DrawIndexedIndirect
//go:noescape

func GPURenderPassEncoderDrawIndexedIndirectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetBindGroup
//go:noescape

func CallGPURenderPassEncoderSetBindGroup(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	bindGroup js.Ref,
	dynamicOffsets js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetBindGroup
//go:noescape

func GPURenderPassEncoderSetBindGroupFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetBindGroup1
//go:noescape

func CallGPURenderPassEncoderSetBindGroup1(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	bindGroup js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetBindGroup1
//go:noescape

func GPURenderPassEncoderSetBindGroup1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetBindGroup2
//go:noescape

func CallGPURenderPassEncoderSetBindGroup2(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	bindGroup js.Ref,
	dynamicOffsetsData js.Ref,
	dynamicOffsetsDataStart float64,
	dynamicOffsetsDataLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetBindGroup2
//go:noescape

func GPURenderPassEncoderSetBindGroup2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_PushDebugGroup
//go:noescape

func CallGPURenderPassEncoderPushDebugGroup(
	this js.Ref,
	ptr unsafe.Pointer,

	groupLabel js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_PushDebugGroup
//go:noescape

func GPURenderPassEncoderPushDebugGroupFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_PopDebugGroup
//go:noescape

func CallGPURenderPassEncoderPopDebugGroup(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_PopDebugGroup
//go:noescape

func GPURenderPassEncoderPopDebugGroupFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderPassEncoder_InsertDebugMarker
//go:noescape

func CallGPURenderPassEncoderInsertDebugMarker(
	this js.Ref,
	ptr unsafe.Pointer,

	markerLabel js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_InsertDebugMarker
//go:noescape

func GPURenderPassEncoderInsertDebugMarkerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GPULoadOp
//go:noescape

func ConstOfGPULoadOp(str js.Ref) uint32

//go:wasmimport plat/js/web constof_GPUStoreOp
//go:noescape

func ConstOfGPUStoreOp(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPURenderPassColorAttachment
//go:noescape

func GPURenderPassColorAttachmentJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPURenderPassColorAttachment
//go:noescape

func GPURenderPassColorAttachmentJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPURenderPassDepthStencilAttachment
//go:noescape

func GPURenderPassDepthStencilAttachmentJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPURenderPassDepthStencilAttachment
//go:noescape

func GPURenderPassDepthStencilAttachmentJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GPUQueryType
//go:noescape

func ConstOfGPUQueryType(str js.Ref) uint32

//go:wasmimport plat/js/web get_GPUQuerySet_Type
//go:noescape

func GetGPUQuerySetType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUQuerySet_Count
//go:noescape

func GetGPUQuerySetCount(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUQuerySet_Label
//go:noescape

func GetGPUQuerySetLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_GPUQuerySet_Label
//go:noescape

func SetGPUQuerySetLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_GPUQuerySet_Destroy
//go:noescape

func CallGPUQuerySetDestroy(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUQuerySet_Destroy
//go:noescape

func GPUQuerySetDestroyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPURenderPassTimestampWrites
//go:noescape

func GPURenderPassTimestampWritesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPURenderPassTimestampWrites
//go:noescape

func GPURenderPassTimestampWritesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPURenderPassDescriptor
//go:noescape

func GPURenderPassDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPURenderPassDescriptor
//go:noescape

func GPURenderPassDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPUComputePassEncoder_Label
//go:noescape

func GetGPUComputePassEncoderLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_GPUComputePassEncoder_Label
//go:noescape

func SetGPUComputePassEncoderLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_GPUComputePassEncoder_SetPipeline
//go:noescape

func CallGPUComputePassEncoderSetPipeline(
	this js.Ref,
	ptr unsafe.Pointer,

	pipeline js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_SetPipeline
//go:noescape

func GPUComputePassEncoderSetPipelineFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUComputePassEncoder_DispatchWorkgroups
//go:noescape

func CallGPUComputePassEncoderDispatchWorkgroups(
	this js.Ref,
	ptr unsafe.Pointer,

	workgroupCountX uint32,
	workgroupCountY uint32,
	workgroupCountZ uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_DispatchWorkgroups
//go:noescape

func GPUComputePassEncoderDispatchWorkgroupsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUComputePassEncoder_DispatchWorkgroups1
//go:noescape

func CallGPUComputePassEncoderDispatchWorkgroups1(
	this js.Ref,
	ptr unsafe.Pointer,

	workgroupCountX uint32,
	workgroupCountY uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_DispatchWorkgroups1
//go:noescape

func GPUComputePassEncoderDispatchWorkgroups1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUComputePassEncoder_DispatchWorkgroups2
//go:noescape

func CallGPUComputePassEncoderDispatchWorkgroups2(
	this js.Ref,
	ptr unsafe.Pointer,

	workgroupCountX uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_DispatchWorkgroups2
//go:noescape

func GPUComputePassEncoderDispatchWorkgroups2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUComputePassEncoder_DispatchWorkgroupsIndirect
//go:noescape

func CallGPUComputePassEncoderDispatchWorkgroupsIndirect(
	this js.Ref,
	ptr unsafe.Pointer,

	indirectBuffer js.Ref,
	indirectOffset float64,
) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_DispatchWorkgroupsIndirect
//go:noescape

func GPUComputePassEncoderDispatchWorkgroupsIndirectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUComputePassEncoder_End
//go:noescape

func CallGPUComputePassEncoderEnd(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_End
//go:noescape

func GPUComputePassEncoderEndFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUComputePassEncoder_PushDebugGroup
//go:noescape

func CallGPUComputePassEncoderPushDebugGroup(
	this js.Ref,
	ptr unsafe.Pointer,

	groupLabel js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_PushDebugGroup
//go:noescape

func GPUComputePassEncoderPushDebugGroupFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUComputePassEncoder_PopDebugGroup
//go:noescape

func CallGPUComputePassEncoderPopDebugGroup(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_PopDebugGroup
//go:noescape

func GPUComputePassEncoderPopDebugGroupFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUComputePassEncoder_InsertDebugMarker
//go:noescape

func CallGPUComputePassEncoderInsertDebugMarker(
	this js.Ref,
	ptr unsafe.Pointer,

	markerLabel js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_InsertDebugMarker
//go:noescape

func GPUComputePassEncoderInsertDebugMarkerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUComputePassEncoder_SetBindGroup
//go:noescape

func CallGPUComputePassEncoderSetBindGroup(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	bindGroup js.Ref,
	dynamicOffsets js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_SetBindGroup
//go:noescape

func GPUComputePassEncoderSetBindGroupFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUComputePassEncoder_SetBindGroup1
//go:noescape

func CallGPUComputePassEncoderSetBindGroup1(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	bindGroup js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_SetBindGroup1
//go:noescape

func GPUComputePassEncoderSetBindGroup1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUComputePassEncoder_SetBindGroup2
//go:noescape

func CallGPUComputePassEncoderSetBindGroup2(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	bindGroup js.Ref,
	dynamicOffsetsData js.Ref,
	dynamicOffsetsDataStart float64,
	dynamicOffsetsDataLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_SetBindGroup2
//go:noescape

func GPUComputePassEncoderSetBindGroup2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUComputePassTimestampWrites
//go:noescape

func GPUComputePassTimestampWritesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUComputePassTimestampWrites
//go:noescape

func GPUComputePassTimestampWritesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUComputePassDescriptor
//go:noescape

func GPUComputePassDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUComputePassDescriptor
//go:noescape

func GPUComputePassDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUImageCopyBuffer
//go:noescape

func GPUImageCopyBufferJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUImageCopyBuffer
//go:noescape

func GPUImageCopyBufferJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUOrigin3DDict
//go:noescape

func GPUOrigin3DDictJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUOrigin3DDict
//go:noescape

func GPUOrigin3DDictJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUImageCopyTexture
//go:noescape

func GPUImageCopyTextureJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUImageCopyTexture
//go:noescape

func GPUImageCopyTextureJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPUCommandBuffer_Label
//go:noescape

func GetGPUCommandBufferLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_GPUCommandBuffer_Label
//go:noescape

func SetGPUCommandBufferLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_GPUCommandBufferDescriptor
//go:noescape

func GPUCommandBufferDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUCommandBufferDescriptor
//go:noescape

func GPUCommandBufferDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPUCommandEncoder_Label
//go:noescape

func GetGPUCommandEncoderLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_GPUCommandEncoder_Label
//go:noescape

func SetGPUCommandEncoderLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_BeginRenderPass
//go:noescape

func CallGPUCommandEncoderBeginRenderPass(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_BeginRenderPass
//go:noescape

func GPUCommandEncoderBeginRenderPassFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_BeginComputePass
//go:noescape

func CallGPUCommandEncoderBeginComputePass(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_BeginComputePass
//go:noescape

func GPUCommandEncoderBeginComputePassFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_BeginComputePass1
//go:noescape

func CallGPUCommandEncoderBeginComputePass1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_BeginComputePass1
//go:noescape

func GPUCommandEncoderBeginComputePass1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_CopyBufferToBuffer
//go:noescape

func CallGPUCommandEncoderCopyBufferToBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	source js.Ref,
	sourceOffset float64,
	destination js.Ref,
	destinationOffset float64,
	size float64,
) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_CopyBufferToBuffer
//go:noescape

func GPUCommandEncoderCopyBufferToBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_CopyBufferToTexture
//go:noescape

func CallGPUCommandEncoderCopyBufferToTexture(
	this js.Ref,
	ptr unsafe.Pointer,

	source unsafe.Pointer,
	destination unsafe.Pointer,
	copySize js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_CopyBufferToTexture
//go:noescape

func GPUCommandEncoderCopyBufferToTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_CopyTextureToBuffer
//go:noescape

func CallGPUCommandEncoderCopyTextureToBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	source unsafe.Pointer,
	destination unsafe.Pointer,
	copySize js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_CopyTextureToBuffer
//go:noescape

func GPUCommandEncoderCopyTextureToBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_CopyTextureToTexture
//go:noescape

func CallGPUCommandEncoderCopyTextureToTexture(
	this js.Ref,
	ptr unsafe.Pointer,

	source unsafe.Pointer,
	destination unsafe.Pointer,
	copySize js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_CopyTextureToTexture
//go:noescape

func GPUCommandEncoderCopyTextureToTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_ClearBuffer
//go:noescape

func CallGPUCommandEncoderClearBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
	offset float64,
	size float64,
) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_ClearBuffer
//go:noescape

func GPUCommandEncoderClearBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_ClearBuffer1
//go:noescape

func CallGPUCommandEncoderClearBuffer1(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_ClearBuffer1
//go:noescape

func GPUCommandEncoderClearBuffer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_ClearBuffer2
//go:noescape

func CallGPUCommandEncoderClearBuffer2(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_ClearBuffer2
//go:noescape

func GPUCommandEncoderClearBuffer2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_WriteTimestamp
//go:noescape

func CallGPUCommandEncoderWriteTimestamp(
	this js.Ref,
	ptr unsafe.Pointer,

	querySet js.Ref,
	queryIndex uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_WriteTimestamp
//go:noescape

func GPUCommandEncoderWriteTimestampFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_ResolveQuerySet
//go:noescape

func CallGPUCommandEncoderResolveQuerySet(
	this js.Ref,
	ptr unsafe.Pointer,

	querySet js.Ref,
	firstQuery uint32,
	queryCount uint32,
	destination js.Ref,
	destinationOffset float64,
) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_ResolveQuerySet
//go:noescape

func GPUCommandEncoderResolveQuerySetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_Finish
//go:noescape

func CallGPUCommandEncoderFinish(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_Finish
//go:noescape

func GPUCommandEncoderFinishFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_Finish1
//go:noescape

func CallGPUCommandEncoderFinish1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_Finish1
//go:noescape

func GPUCommandEncoderFinish1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_PushDebugGroup
//go:noescape

func CallGPUCommandEncoderPushDebugGroup(
	this js.Ref,
	ptr unsafe.Pointer,

	groupLabel js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_PushDebugGroup
//go:noescape

func GPUCommandEncoderPushDebugGroupFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_PopDebugGroup
//go:noescape

func CallGPUCommandEncoderPopDebugGroup(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_PopDebugGroup
//go:noescape

func GPUCommandEncoderPopDebugGroupFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCommandEncoder_InsertDebugMarker
//go:noescape

func CallGPUCommandEncoderInsertDebugMarker(
	this js.Ref,
	ptr unsafe.Pointer,

	markerLabel js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_InsertDebugMarker
//go:noescape

func GPUCommandEncoderInsertDebugMarkerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUCommandEncoderDescriptor
//go:noescape

func GPUCommandEncoderDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUCommandEncoderDescriptor
//go:noescape

func GPUCommandEncoderDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPURenderBundleDescriptor
//go:noescape

func GPURenderBundleDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPURenderBundleDescriptor
//go:noescape

func GPURenderBundleDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPURenderBundleEncoder_Label
//go:noescape

func GetGPURenderBundleEncoderLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_GPURenderBundleEncoder_Label
//go:noescape

func SetGPURenderBundleEncoderLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_Finish
//go:noescape

func CallGPURenderBundleEncoderFinish(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_Finish
//go:noescape

func GPURenderBundleEncoderFinishFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_Finish1
//go:noescape

func CallGPURenderBundleEncoderFinish1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_Finish1
//go:noescape

func GPURenderBundleEncoderFinish1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetPipeline
//go:noescape

func CallGPURenderBundleEncoderSetPipeline(
	this js.Ref,
	ptr unsafe.Pointer,

	pipeline js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetPipeline
//go:noescape

func GPURenderBundleEncoderSetPipelineFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetIndexBuffer
//go:noescape

func CallGPURenderBundleEncoderSetIndexBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
	indexFormat uint32,
	offset float64,
	size float64,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetIndexBuffer
//go:noescape

func GPURenderBundleEncoderSetIndexBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetIndexBuffer1
//go:noescape

func CallGPURenderBundleEncoderSetIndexBuffer1(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
	indexFormat uint32,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetIndexBuffer1
//go:noescape

func GPURenderBundleEncoderSetIndexBuffer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetIndexBuffer2
//go:noescape

func CallGPURenderBundleEncoderSetIndexBuffer2(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
	indexFormat uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetIndexBuffer2
//go:noescape

func GPURenderBundleEncoderSetIndexBuffer2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetVertexBuffer
//go:noescape

func CallGPURenderBundleEncoderSetVertexBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	slot uint32,
	buffer js.Ref,
	offset float64,
	size float64,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetVertexBuffer
//go:noescape

func GPURenderBundleEncoderSetVertexBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetVertexBuffer1
//go:noescape

func CallGPURenderBundleEncoderSetVertexBuffer1(
	this js.Ref,
	ptr unsafe.Pointer,

	slot uint32,
	buffer js.Ref,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetVertexBuffer1
//go:noescape

func GPURenderBundleEncoderSetVertexBuffer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetVertexBuffer2
//go:noescape

func CallGPURenderBundleEncoderSetVertexBuffer2(
	this js.Ref,
	ptr unsafe.Pointer,

	slot uint32,
	buffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetVertexBuffer2
//go:noescape

func GPURenderBundleEncoderSetVertexBuffer2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_Draw
//go:noescape

func CallGPURenderBundleEncoderDraw(
	this js.Ref,
	ptr unsafe.Pointer,

	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32,
	firstInstance uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_Draw
//go:noescape

func GPURenderBundleEncoderDrawFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_Draw1
//go:noescape

func CallGPURenderBundleEncoderDraw1(
	this js.Ref,
	ptr unsafe.Pointer,

	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_Draw1
//go:noescape

func GPURenderBundleEncoderDraw1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_Draw2
//go:noescape

func CallGPURenderBundleEncoderDraw2(
	this js.Ref,
	ptr unsafe.Pointer,

	vertexCount uint32,
	instanceCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_Draw2
//go:noescape

func GPURenderBundleEncoderDraw2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_Draw3
//go:noescape

func CallGPURenderBundleEncoderDraw3(
	this js.Ref,
	ptr unsafe.Pointer,

	vertexCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_Draw3
//go:noescape

func GPURenderBundleEncoderDraw3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_DrawIndexed
//go:noescape

func CallGPURenderBundleEncoderDrawIndexed(
	this js.Ref,
	ptr unsafe.Pointer,

	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	baseVertex int32,
	firstInstance uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_DrawIndexed
//go:noescape

func GPURenderBundleEncoderDrawIndexedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_DrawIndexed1
//go:noescape

func CallGPURenderBundleEncoderDrawIndexed1(
	this js.Ref,
	ptr unsafe.Pointer,

	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	baseVertex int32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_DrawIndexed1
//go:noescape

func GPURenderBundleEncoderDrawIndexed1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_DrawIndexed2
//go:noescape

func CallGPURenderBundleEncoderDrawIndexed2(
	this js.Ref,
	ptr unsafe.Pointer,

	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_DrawIndexed2
//go:noescape

func GPURenderBundleEncoderDrawIndexed2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_DrawIndexed3
//go:noescape

func CallGPURenderBundleEncoderDrawIndexed3(
	this js.Ref,
	ptr unsafe.Pointer,

	indexCount uint32,
	instanceCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_DrawIndexed3
//go:noescape

func GPURenderBundleEncoderDrawIndexed3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_DrawIndexed4
//go:noescape

func CallGPURenderBundleEncoderDrawIndexed4(
	this js.Ref,
	ptr unsafe.Pointer,

	indexCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_DrawIndexed4
//go:noescape

func GPURenderBundleEncoderDrawIndexed4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_DrawIndirect
//go:noescape

func CallGPURenderBundleEncoderDrawIndirect(
	this js.Ref,
	ptr unsafe.Pointer,

	indirectBuffer js.Ref,
	indirectOffset float64,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_DrawIndirect
//go:noescape

func GPURenderBundleEncoderDrawIndirectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_DrawIndexedIndirect
//go:noescape

func CallGPURenderBundleEncoderDrawIndexedIndirect(
	this js.Ref,
	ptr unsafe.Pointer,

	indirectBuffer js.Ref,
	indirectOffset float64,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_DrawIndexedIndirect
//go:noescape

func GPURenderBundleEncoderDrawIndexedIndirectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetBindGroup
//go:noescape

func CallGPURenderBundleEncoderSetBindGroup(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	bindGroup js.Ref,
	dynamicOffsets js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetBindGroup
//go:noescape

func GPURenderBundleEncoderSetBindGroupFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetBindGroup1
//go:noescape

func CallGPURenderBundleEncoderSetBindGroup1(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	bindGroup js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetBindGroup1
//go:noescape

func GPURenderBundleEncoderSetBindGroup1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetBindGroup2
//go:noescape

func CallGPURenderBundleEncoderSetBindGroup2(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	bindGroup js.Ref,
	dynamicOffsetsData js.Ref,
	dynamicOffsetsDataStart float64,
	dynamicOffsetsDataLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetBindGroup2
//go:noescape

func GPURenderBundleEncoderSetBindGroup2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_PushDebugGroup
//go:noescape

func CallGPURenderBundleEncoderPushDebugGroup(
	this js.Ref,
	ptr unsafe.Pointer,

	groupLabel js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_PushDebugGroup
//go:noescape

func GPURenderBundleEncoderPushDebugGroupFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_PopDebugGroup
//go:noescape

func CallGPURenderBundleEncoderPopDebugGroup(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_PopDebugGroup
//go:noescape

func GPURenderBundleEncoderPopDebugGroupFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_InsertDebugMarker
//go:noescape

func CallGPURenderBundleEncoderInsertDebugMarker(
	this js.Ref,
	ptr unsafe.Pointer,

	markerLabel js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_InsertDebugMarker
//go:noescape

func GPURenderBundleEncoderInsertDebugMarkerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPURenderBundleEncoderDescriptor
//go:noescape

func GPURenderBundleEncoderDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPURenderBundleEncoderDescriptor
//go:noescape

func GPURenderBundleEncoderDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUQuerySetDescriptor
//go:noescape

func GPUQuerySetDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUQuerySetDescriptor
//go:noescape

func GPUQuerySetDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GPUErrorFilter
//go:noescape

func ConstOfGPUErrorFilter(str js.Ref) uint32

//go:wasmimport plat/js/web get_GPUError_Message
//go:noescape

func GetGPUErrorMessage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxTextureDimension1D
//go:noescape

func GetGPUSupportedLimitsMaxTextureDimension1D(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxTextureDimension2D
//go:noescape

func GetGPUSupportedLimitsMaxTextureDimension2D(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxTextureDimension3D
//go:noescape

func GetGPUSupportedLimitsMaxTextureDimension3D(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxTextureArrayLayers
//go:noescape

func GetGPUSupportedLimitsMaxTextureArrayLayers(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxBindGroups
//go:noescape

func GetGPUSupportedLimitsMaxBindGroups(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxBindGroupsPlusVertexBuffers
//go:noescape

func GetGPUSupportedLimitsMaxBindGroupsPlusVertexBuffers(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxBindingsPerBindGroup
//go:noescape

func GetGPUSupportedLimitsMaxBindingsPerBindGroup(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxDynamicUniformBuffersPerPipelineLayout
//go:noescape

func GetGPUSupportedLimitsMaxDynamicUniformBuffersPerPipelineLayout(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxDynamicStorageBuffersPerPipelineLayout
//go:noescape

func GetGPUSupportedLimitsMaxDynamicStorageBuffersPerPipelineLayout(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxSampledTexturesPerShaderStage
//go:noescape

func GetGPUSupportedLimitsMaxSampledTexturesPerShaderStage(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxSamplersPerShaderStage
//go:noescape

func GetGPUSupportedLimitsMaxSamplersPerShaderStage(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxStorageBuffersPerShaderStage
//go:noescape

func GetGPUSupportedLimitsMaxStorageBuffersPerShaderStage(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxStorageTexturesPerShaderStage
//go:noescape

func GetGPUSupportedLimitsMaxStorageTexturesPerShaderStage(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxUniformBuffersPerShaderStage
//go:noescape

func GetGPUSupportedLimitsMaxUniformBuffersPerShaderStage(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxUniformBufferBindingSize
//go:noescape

func GetGPUSupportedLimitsMaxUniformBufferBindingSize(
	this js.Ref,
	ptr unsafe.Pointer,
) uint64

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxStorageBufferBindingSize
//go:noescape

func GetGPUSupportedLimitsMaxStorageBufferBindingSize(
	this js.Ref,
	ptr unsafe.Pointer,
) uint64

//go:wasmimport plat/js/web get_GPUSupportedLimits_MinUniformBufferOffsetAlignment
//go:noescape

func GetGPUSupportedLimitsMinUniformBufferOffsetAlignment(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MinStorageBufferOffsetAlignment
//go:noescape

func GetGPUSupportedLimitsMinStorageBufferOffsetAlignment(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxVertexBuffers
//go:noescape

func GetGPUSupportedLimitsMaxVertexBuffers(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxBufferSize
//go:noescape

func GetGPUSupportedLimitsMaxBufferSize(
	this js.Ref,
	ptr unsafe.Pointer,
) uint64

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxVertexAttributes
//go:noescape

func GetGPUSupportedLimitsMaxVertexAttributes(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxVertexBufferArrayStride
//go:noescape

func GetGPUSupportedLimitsMaxVertexBufferArrayStride(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxInterStageShaderComponents
//go:noescape

func GetGPUSupportedLimitsMaxInterStageShaderComponents(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxInterStageShaderVariables
//go:noescape

func GetGPUSupportedLimitsMaxInterStageShaderVariables(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxColorAttachments
//go:noescape

func GetGPUSupportedLimitsMaxColorAttachments(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxColorAttachmentBytesPerSample
//go:noescape

func GetGPUSupportedLimitsMaxColorAttachmentBytesPerSample(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxComputeWorkgroupStorageSize
//go:noescape

func GetGPUSupportedLimitsMaxComputeWorkgroupStorageSize(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxComputeInvocationsPerWorkgroup
//go:noescape

func GetGPUSupportedLimitsMaxComputeInvocationsPerWorkgroup(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxComputeWorkgroupSizeX
//go:noescape

func GetGPUSupportedLimitsMaxComputeWorkgroupSizeX(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxComputeWorkgroupSizeY
//go:noescape

func GetGPUSupportedLimitsMaxComputeWorkgroupSizeY(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxComputeWorkgroupSizeZ
//go:noescape

func GetGPUSupportedLimitsMaxComputeWorkgroupSizeZ(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxComputeWorkgroupsPerDimension
//go:noescape

func GetGPUSupportedLimitsMaxComputeWorkgroupsPerDimension(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web store_GPUImageDataLayout
//go:noescape

func GPUImageDataLayoutJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUImageDataLayout
//go:noescape

func GPUImageDataLayoutJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUOrigin2DDict
//go:noescape

func GPUOrigin2DDictJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUOrigin2DDict
//go:noescape

func GPUOrigin2DDictJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUImageCopyExternalImage
//go:noescape

func GPUImageCopyExternalImageJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUImageCopyExternalImage
//go:noescape

func GPUImageCopyExternalImageJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUImageCopyTextureTagged
//go:noescape

func GPUImageCopyTextureTaggedJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUImageCopyTextureTagged
//go:noescape

func GPUImageCopyTextureTaggedJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPUQueue_Label
//go:noescape

func GetGPUQueueLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_GPUQueue_Label
//go:noescape

func SetGPUQueueLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_GPUQueue_Submit
//go:noescape

func CallGPUQueueSubmit(
	this js.Ref,
	ptr unsafe.Pointer,

	commandBuffers js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPUQueue_Submit
//go:noescape

func GPUQueueSubmitFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUQueue_OnSubmittedWorkDone
//go:noescape

func CallGPUQueueOnSubmittedWorkDone(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUQueue_OnSubmittedWorkDone
//go:noescape

func GPUQueueOnSubmittedWorkDoneFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUQueue_WriteBuffer
//go:noescape

func CallGPUQueueWriteBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
	bufferOffset float64,
	data js.Ref,
	dataOffset float64,
	size float64,
) js.Ref

//go:wasmimport plat/js/web func_GPUQueue_WriteBuffer
//go:noescape

func GPUQueueWriteBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUQueue_WriteBuffer1
//go:noescape

func CallGPUQueueWriteBuffer1(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
	bufferOffset float64,
	data js.Ref,
	dataOffset float64,
) js.Ref

//go:wasmimport plat/js/web func_GPUQueue_WriteBuffer1
//go:noescape

func GPUQueueWriteBuffer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUQueue_WriteBuffer2
//go:noescape

func CallGPUQueueWriteBuffer2(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
	bufferOffset float64,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPUQueue_WriteBuffer2
//go:noescape

func GPUQueueWriteBuffer2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUQueue_WriteTexture
//go:noescape

func CallGPUQueueWriteTexture(
	this js.Ref,
	ptr unsafe.Pointer,

	destination unsafe.Pointer,
	data js.Ref,
	dataLayout unsafe.Pointer,
	size js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPUQueue_WriteTexture
//go:noescape

func GPUQueueWriteTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUQueue_CopyExternalImageToTexture
//go:noescape

func CallGPUQueueCopyExternalImageToTexture(
	this js.Ref,
	ptr unsafe.Pointer,

	source unsafe.Pointer,
	destination unsafe.Pointer,
	copySize js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_GPUQueue_CopyExternalImageToTexture
//go:noescape

func GPUQueueCopyExternalImageToTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GPUDeviceLostReason
//go:noescape

func ConstOfGPUDeviceLostReason(str js.Ref) uint32

//go:wasmimport plat/js/web get_GPUDeviceLostInfo_Reason
//go:noescape

func GetGPUDeviceLostInfoReason(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUDeviceLostInfo_Message
//go:noescape

func GetGPUDeviceLostInfoMessage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_GPUDevice_Features
//go:noescape

func GetGPUDeviceFeatures(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_GPUDevice_Limits
//go:noescape

func GetGPUDeviceLimits(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_GPUDevice_Queue
//go:noescape

func GetGPUDeviceQueue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_GPUDevice_Lost
//go:noescape

func GetGPUDeviceLost(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_GPUDevice_Label
//go:noescape

func GetGPUDeviceLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_GPUDevice_Label
//go:noescape

func SetGPUDeviceLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_Destroy
//go:noescape

func CallGPUDeviceDestroy(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_Destroy
//go:noescape

func GPUDeviceDestroyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreateBuffer
//go:noescape

func CallGPUDeviceCreateBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateBuffer
//go:noescape

func GPUDeviceCreateBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreateTexture
//go:noescape

func CallGPUDeviceCreateTexture(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateTexture
//go:noescape

func GPUDeviceCreateTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreateSampler
//go:noescape

func CallGPUDeviceCreateSampler(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateSampler
//go:noescape

func GPUDeviceCreateSamplerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreateSampler1
//go:noescape

func CallGPUDeviceCreateSampler1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateSampler1
//go:noescape

func GPUDeviceCreateSampler1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_ImportExternalTexture
//go:noescape

func CallGPUDeviceImportExternalTexture(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_ImportExternalTexture
//go:noescape

func GPUDeviceImportExternalTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreateBindGroupLayout
//go:noescape

func CallGPUDeviceCreateBindGroupLayout(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateBindGroupLayout
//go:noescape

func GPUDeviceCreateBindGroupLayoutFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreatePipelineLayout
//go:noescape

func CallGPUDeviceCreatePipelineLayout(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreatePipelineLayout
//go:noescape

func GPUDeviceCreatePipelineLayoutFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreateBindGroup
//go:noescape

func CallGPUDeviceCreateBindGroup(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateBindGroup
//go:noescape

func GPUDeviceCreateBindGroupFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreateShaderModule
//go:noescape

func CallGPUDeviceCreateShaderModule(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateShaderModule
//go:noescape

func GPUDeviceCreateShaderModuleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreateComputePipeline
//go:noescape

func CallGPUDeviceCreateComputePipeline(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateComputePipeline
//go:noescape

func GPUDeviceCreateComputePipelineFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreateRenderPipeline
//go:noescape

func CallGPUDeviceCreateRenderPipeline(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateRenderPipeline
//go:noescape

func GPUDeviceCreateRenderPipelineFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreateComputePipelineAsync
//go:noescape

func CallGPUDeviceCreateComputePipelineAsync(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateComputePipelineAsync
//go:noescape

func GPUDeviceCreateComputePipelineAsyncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreateRenderPipelineAsync
//go:noescape

func CallGPUDeviceCreateRenderPipelineAsync(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateRenderPipelineAsync
//go:noescape

func GPUDeviceCreateRenderPipelineAsyncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreateCommandEncoder
//go:noescape

func CallGPUDeviceCreateCommandEncoder(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateCommandEncoder
//go:noescape

func GPUDeviceCreateCommandEncoderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreateCommandEncoder1
//go:noescape

func CallGPUDeviceCreateCommandEncoder1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateCommandEncoder1
//go:noescape

func GPUDeviceCreateCommandEncoder1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreateRenderBundleEncoder
//go:noescape

func CallGPUDeviceCreateRenderBundleEncoder(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateRenderBundleEncoder
//go:noescape

func GPUDeviceCreateRenderBundleEncoderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_CreateQuerySet
//go:noescape

func CallGPUDeviceCreateQuerySet(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateQuerySet
//go:noescape

func GPUDeviceCreateQuerySetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_PushErrorScope
//go:noescape

func CallGPUDevicePushErrorScope(
	this js.Ref,
	ptr unsafe.Pointer,

	filter uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_PushErrorScope
//go:noescape

func GPUDevicePushErrorScopeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUDevice_PopErrorScope
//go:noescape

func CallGPUDevicePopErrorScope(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_PopErrorScope
//go:noescape

func GPUDevicePopErrorScopeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GPUCanvasAlphaMode
//go:noescape

func ConstOfGPUCanvasAlphaMode(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPUCanvasConfiguration
//go:noescape

func GPUCanvasConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUCanvasConfiguration
//go:noescape

func GPUCanvasConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPUCanvasContext_Canvas
//go:noescape

func GetGPUCanvasContextCanvas(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_GPUCanvasContext_Configure
//go:noescape

func CallGPUCanvasContextConfigure(
	this js.Ref,
	ptr unsafe.Pointer,

	configuration unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUCanvasContext_Configure
//go:noescape

func GPUCanvasContextConfigureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCanvasContext_Unconfigure
//go:noescape

func CallGPUCanvasContextUnconfigure(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUCanvasContext_Unconfigure
//go:noescape

func GPUCanvasContextUnconfigureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUCanvasContext_GetCurrentTexture
//go:noescape

func CallGPUCanvasContextGetCurrentTexture(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUCanvasContext_GetCurrentTexture
//go:noescape

func GPUCanvasContextGetCurrentTextureFunc(this js.Ref) js.Ref
