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

//go:wasmimport plat/js/web get_GPURenderPassEncoder_Label
//go:noescape
func GetGPURenderPassEncoderLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPURenderPassEncoder_Label
//go:noescape
func SetGPURenderPassEncoderLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_GPURenderPassEncoder_SetViewport
//go:noescape
func HasFuncGPURenderPassEncoderSetViewport(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetViewport
//go:noescape
func FuncGPURenderPassEncoderSetViewport(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetViewport
//go:noescape
func CallGPURenderPassEncoderSetViewport(
	this js.Ref, retPtr unsafe.Pointer,
	x float32,
	y float32,
	width float32,
	height float32,
	minDepth float32,
	maxDepth float32)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_SetViewport
//go:noescape
func TryGPURenderPassEncoderSetViewport(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float32,
	y float32,
	width float32,
	height float32,
	minDepth float32,
	maxDepth float32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_SetScissorRect
//go:noescape
func HasFuncGPURenderPassEncoderSetScissorRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetScissorRect
//go:noescape
func FuncGPURenderPassEncoderSetScissorRect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetScissorRect
//go:noescape
func CallGPURenderPassEncoderSetScissorRect(
	this js.Ref, retPtr unsafe.Pointer,
	x uint32,
	y uint32,
	width uint32,
	height uint32)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_SetScissorRect
//go:noescape
func TryGPURenderPassEncoderSetScissorRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x uint32,
	y uint32,
	width uint32,
	height uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_SetBlendConstant
//go:noescape
func HasFuncGPURenderPassEncoderSetBlendConstant(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetBlendConstant
//go:noescape
func FuncGPURenderPassEncoderSetBlendConstant(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetBlendConstant
//go:noescape
func CallGPURenderPassEncoderSetBlendConstant(
	this js.Ref, retPtr unsafe.Pointer,
	color js.Ref)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_SetBlendConstant
//go:noescape
func TryGPURenderPassEncoderSetBlendConstant(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	color js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_SetStencilReference
//go:noescape
func HasFuncGPURenderPassEncoderSetStencilReference(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetStencilReference
//go:noescape
func FuncGPURenderPassEncoderSetStencilReference(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetStencilReference
//go:noescape
func CallGPURenderPassEncoderSetStencilReference(
	this js.Ref, retPtr unsafe.Pointer,
	reference uint32)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_SetStencilReference
//go:noescape
func TryGPURenderPassEncoderSetStencilReference(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	reference uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_BeginOcclusionQuery
//go:noescape
func HasFuncGPURenderPassEncoderBeginOcclusionQuery(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_BeginOcclusionQuery
//go:noescape
func FuncGPURenderPassEncoderBeginOcclusionQuery(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_BeginOcclusionQuery
//go:noescape
func CallGPURenderPassEncoderBeginOcclusionQuery(
	this js.Ref, retPtr unsafe.Pointer,
	queryIndex uint32)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_BeginOcclusionQuery
//go:noescape
func TryGPURenderPassEncoderBeginOcclusionQuery(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	queryIndex uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_EndOcclusionQuery
//go:noescape
func HasFuncGPURenderPassEncoderEndOcclusionQuery(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_EndOcclusionQuery
//go:noescape
func FuncGPURenderPassEncoderEndOcclusionQuery(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_EndOcclusionQuery
//go:noescape
func CallGPURenderPassEncoderEndOcclusionQuery(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_EndOcclusionQuery
//go:noescape
func TryGPURenderPassEncoderEndOcclusionQuery(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_ExecuteBundles
//go:noescape
func HasFuncGPURenderPassEncoderExecuteBundles(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_ExecuteBundles
//go:noescape
func FuncGPURenderPassEncoderExecuteBundles(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_ExecuteBundles
//go:noescape
func CallGPURenderPassEncoderExecuteBundles(
	this js.Ref, retPtr unsafe.Pointer,
	bundles js.Ref)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_ExecuteBundles
//go:noescape
func TryGPURenderPassEncoderExecuteBundles(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bundles js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_End
//go:noescape
func HasFuncGPURenderPassEncoderEnd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_End
//go:noescape
func FuncGPURenderPassEncoderEnd(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_End
//go:noescape
func CallGPURenderPassEncoderEnd(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_End
//go:noescape
func TryGPURenderPassEncoderEnd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_SetPipeline
//go:noescape
func HasFuncGPURenderPassEncoderSetPipeline(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetPipeline
//go:noescape
func FuncGPURenderPassEncoderSetPipeline(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetPipeline
//go:noescape
func CallGPURenderPassEncoderSetPipeline(
	this js.Ref, retPtr unsafe.Pointer,
	pipeline js.Ref)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_SetPipeline
//go:noescape
func TryGPURenderPassEncoderSetPipeline(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	pipeline js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_SetIndexBuffer
//go:noescape
func HasFuncGPURenderPassEncoderSetIndexBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetIndexBuffer
//go:noescape
func FuncGPURenderPassEncoderSetIndexBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetIndexBuffer
//go:noescape
func CallGPURenderPassEncoderSetIndexBuffer(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref,
	indexFormat uint32,
	offset float64,
	size float64)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_SetIndexBuffer
//go:noescape
func TryGPURenderPassEncoderSetIndexBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref,
	indexFormat uint32,
	offset float64,
	size float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_SetIndexBuffer1
//go:noescape
func HasFuncGPURenderPassEncoderSetIndexBuffer1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetIndexBuffer1
//go:noescape
func FuncGPURenderPassEncoderSetIndexBuffer1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetIndexBuffer1
//go:noescape
func CallGPURenderPassEncoderSetIndexBuffer1(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref,
	indexFormat uint32,
	offset float64)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_SetIndexBuffer1
//go:noescape
func TryGPURenderPassEncoderSetIndexBuffer1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref,
	indexFormat uint32,
	offset float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_SetIndexBuffer2
//go:noescape
func HasFuncGPURenderPassEncoderSetIndexBuffer2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetIndexBuffer2
//go:noescape
func FuncGPURenderPassEncoderSetIndexBuffer2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetIndexBuffer2
//go:noescape
func CallGPURenderPassEncoderSetIndexBuffer2(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref,
	indexFormat uint32)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_SetIndexBuffer2
//go:noescape
func TryGPURenderPassEncoderSetIndexBuffer2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref,
	indexFormat uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_SetVertexBuffer
//go:noescape
func HasFuncGPURenderPassEncoderSetVertexBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetVertexBuffer
//go:noescape
func FuncGPURenderPassEncoderSetVertexBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetVertexBuffer
//go:noescape
func CallGPURenderPassEncoderSetVertexBuffer(
	this js.Ref, retPtr unsafe.Pointer,
	slot uint32,
	buffer js.Ref,
	offset float64,
	size float64)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_SetVertexBuffer
//go:noescape
func TryGPURenderPassEncoderSetVertexBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	slot uint32,
	buffer js.Ref,
	offset float64,
	size float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_SetVertexBuffer1
//go:noescape
func HasFuncGPURenderPassEncoderSetVertexBuffer1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetVertexBuffer1
//go:noescape
func FuncGPURenderPassEncoderSetVertexBuffer1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetVertexBuffer1
//go:noescape
func CallGPURenderPassEncoderSetVertexBuffer1(
	this js.Ref, retPtr unsafe.Pointer,
	slot uint32,
	buffer js.Ref,
	offset float64)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_SetVertexBuffer1
//go:noescape
func TryGPURenderPassEncoderSetVertexBuffer1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	slot uint32,
	buffer js.Ref,
	offset float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_SetVertexBuffer2
//go:noescape
func HasFuncGPURenderPassEncoderSetVertexBuffer2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetVertexBuffer2
//go:noescape
func FuncGPURenderPassEncoderSetVertexBuffer2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetVertexBuffer2
//go:noescape
func CallGPURenderPassEncoderSetVertexBuffer2(
	this js.Ref, retPtr unsafe.Pointer,
	slot uint32,
	buffer js.Ref)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_SetVertexBuffer2
//go:noescape
func TryGPURenderPassEncoderSetVertexBuffer2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	slot uint32,
	buffer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_Draw
//go:noescape
func HasFuncGPURenderPassEncoderDraw(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_Draw
//go:noescape
func FuncGPURenderPassEncoderDraw(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_Draw
//go:noescape
func CallGPURenderPassEncoderDraw(
	this js.Ref, retPtr unsafe.Pointer,
	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32,
	firstInstance uint32)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_Draw
//go:noescape
func TryGPURenderPassEncoderDraw(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32,
	firstInstance uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_Draw1
//go:noescape
func HasFuncGPURenderPassEncoderDraw1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_Draw1
//go:noescape
func FuncGPURenderPassEncoderDraw1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_Draw1
//go:noescape
func CallGPURenderPassEncoderDraw1(
	this js.Ref, retPtr unsafe.Pointer,
	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_Draw1
//go:noescape
func TryGPURenderPassEncoderDraw1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_Draw2
//go:noescape
func HasFuncGPURenderPassEncoderDraw2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_Draw2
//go:noescape
func FuncGPURenderPassEncoderDraw2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_Draw2
//go:noescape
func CallGPURenderPassEncoderDraw2(
	this js.Ref, retPtr unsafe.Pointer,
	vertexCount uint32,
	instanceCount uint32)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_Draw2
//go:noescape
func TryGPURenderPassEncoderDraw2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	vertexCount uint32,
	instanceCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_Draw3
//go:noescape
func HasFuncGPURenderPassEncoderDraw3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_Draw3
//go:noescape
func FuncGPURenderPassEncoderDraw3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_Draw3
//go:noescape
func CallGPURenderPassEncoderDraw3(
	this js.Ref, retPtr unsafe.Pointer,
	vertexCount uint32)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_Draw3
//go:noescape
func TryGPURenderPassEncoderDraw3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	vertexCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_DrawIndexed
//go:noescape
func HasFuncGPURenderPassEncoderDrawIndexed(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_DrawIndexed
//go:noescape
func FuncGPURenderPassEncoderDrawIndexed(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_DrawIndexed
//go:noescape
func CallGPURenderPassEncoderDrawIndexed(
	this js.Ref, retPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	baseVertex int32,
	firstInstance uint32)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_DrawIndexed
//go:noescape
func TryGPURenderPassEncoderDrawIndexed(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	baseVertex int32,
	firstInstance uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_DrawIndexed1
//go:noescape
func HasFuncGPURenderPassEncoderDrawIndexed1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_DrawIndexed1
//go:noescape
func FuncGPURenderPassEncoderDrawIndexed1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_DrawIndexed1
//go:noescape
func CallGPURenderPassEncoderDrawIndexed1(
	this js.Ref, retPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	baseVertex int32)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_DrawIndexed1
//go:noescape
func TryGPURenderPassEncoderDrawIndexed1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	baseVertex int32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_DrawIndexed2
//go:noescape
func HasFuncGPURenderPassEncoderDrawIndexed2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_DrawIndexed2
//go:noescape
func FuncGPURenderPassEncoderDrawIndexed2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_DrawIndexed2
//go:noescape
func CallGPURenderPassEncoderDrawIndexed2(
	this js.Ref, retPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_DrawIndexed2
//go:noescape
func TryGPURenderPassEncoderDrawIndexed2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_DrawIndexed3
//go:noescape
func HasFuncGPURenderPassEncoderDrawIndexed3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_DrawIndexed3
//go:noescape
func FuncGPURenderPassEncoderDrawIndexed3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_DrawIndexed3
//go:noescape
func CallGPURenderPassEncoderDrawIndexed3(
	this js.Ref, retPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_DrawIndexed3
//go:noescape
func TryGPURenderPassEncoderDrawIndexed3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_DrawIndexed4
//go:noescape
func HasFuncGPURenderPassEncoderDrawIndexed4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_DrawIndexed4
//go:noescape
func FuncGPURenderPassEncoderDrawIndexed4(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_DrawIndexed4
//go:noescape
func CallGPURenderPassEncoderDrawIndexed4(
	this js.Ref, retPtr unsafe.Pointer,
	indexCount uint32)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_DrawIndexed4
//go:noescape
func TryGPURenderPassEncoderDrawIndexed4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	indexCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_DrawIndirect
//go:noescape
func HasFuncGPURenderPassEncoderDrawIndirect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_DrawIndirect
//go:noescape
func FuncGPURenderPassEncoderDrawIndirect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_DrawIndirect
//go:noescape
func CallGPURenderPassEncoderDrawIndirect(
	this js.Ref, retPtr unsafe.Pointer,
	indirectBuffer js.Ref,
	indirectOffset float64)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_DrawIndirect
//go:noescape
func TryGPURenderPassEncoderDrawIndirect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	indirectBuffer js.Ref,
	indirectOffset float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_DrawIndexedIndirect
//go:noescape
func HasFuncGPURenderPassEncoderDrawIndexedIndirect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_DrawIndexedIndirect
//go:noescape
func FuncGPURenderPassEncoderDrawIndexedIndirect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_DrawIndexedIndirect
//go:noescape
func CallGPURenderPassEncoderDrawIndexedIndirect(
	this js.Ref, retPtr unsafe.Pointer,
	indirectBuffer js.Ref,
	indirectOffset float64)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_DrawIndexedIndirect
//go:noescape
func TryGPURenderPassEncoderDrawIndexedIndirect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	indirectBuffer js.Ref,
	indirectOffset float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_SetBindGroup
//go:noescape
func HasFuncGPURenderPassEncoderSetBindGroup(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetBindGroup
//go:noescape
func FuncGPURenderPassEncoderSetBindGroup(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetBindGroup
//go:noescape
func CallGPURenderPassEncoderSetBindGroup(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref,
	dynamicOffsets js.Ref)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_SetBindGroup
//go:noescape
func TryGPURenderPassEncoderSetBindGroup(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref,
	dynamicOffsets js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_SetBindGroup1
//go:noescape
func HasFuncGPURenderPassEncoderSetBindGroup1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetBindGroup1
//go:noescape
func FuncGPURenderPassEncoderSetBindGroup1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetBindGroup1
//go:noescape
func CallGPURenderPassEncoderSetBindGroup1(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_SetBindGroup1
//go:noescape
func TryGPURenderPassEncoderSetBindGroup1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_SetBindGroup2
//go:noescape
func HasFuncGPURenderPassEncoderSetBindGroup2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_SetBindGroup2
//go:noescape
func FuncGPURenderPassEncoderSetBindGroup2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_SetBindGroup2
//go:noescape
func CallGPURenderPassEncoderSetBindGroup2(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref,
	dynamicOffsetsData js.Ref,
	dynamicOffsetsDataStart float64,
	dynamicOffsetsDataLength uint32)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_SetBindGroup2
//go:noescape
func TryGPURenderPassEncoderSetBindGroup2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref,
	dynamicOffsetsData js.Ref,
	dynamicOffsetsDataStart float64,
	dynamicOffsetsDataLength uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_PushDebugGroup
//go:noescape
func HasFuncGPURenderPassEncoderPushDebugGroup(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_PushDebugGroup
//go:noescape
func FuncGPURenderPassEncoderPushDebugGroup(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_PushDebugGroup
//go:noescape
func CallGPURenderPassEncoderPushDebugGroup(
	this js.Ref, retPtr unsafe.Pointer,
	groupLabel js.Ref)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_PushDebugGroup
//go:noescape
func TryGPURenderPassEncoderPushDebugGroup(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	groupLabel js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_PopDebugGroup
//go:noescape
func HasFuncGPURenderPassEncoderPopDebugGroup(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_PopDebugGroup
//go:noescape
func FuncGPURenderPassEncoderPopDebugGroup(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_PopDebugGroup
//go:noescape
func CallGPURenderPassEncoderPopDebugGroup(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_PopDebugGroup
//go:noescape
func TryGPURenderPassEncoderPopDebugGroup(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderPassEncoder_InsertDebugMarker
//go:noescape
func HasFuncGPURenderPassEncoderInsertDebugMarker(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPassEncoder_InsertDebugMarker
//go:noescape
func FuncGPURenderPassEncoderInsertDebugMarker(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPassEncoder_InsertDebugMarker
//go:noescape
func CallGPURenderPassEncoderInsertDebugMarker(
	this js.Ref, retPtr unsafe.Pointer,
	markerLabel js.Ref)

//go:wasmimport plat/js/web try_GPURenderPassEncoder_InsertDebugMarker
//go:noescape
func TryGPURenderPassEncoderInsertDebugMarker(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	markerLabel js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUQuerySet_Count
//go:noescape
func GetGPUQuerySetCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUQuerySet_Label
//go:noescape
func GetGPUQuerySetLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPUQuerySet_Label
//go:noescape
func SetGPUQuerySetLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_GPUQuerySet_Destroy
//go:noescape
func HasFuncGPUQuerySetDestroy(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUQuerySet_Destroy
//go:noescape
func FuncGPUQuerySetDestroy(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUQuerySet_Destroy
//go:noescape
func CallGPUQuerySetDestroy(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUQuerySet_Destroy
//go:noescape
func TryGPUQuerySetDestroy(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPUComputePassEncoder_Label
//go:noescape
func SetGPUComputePassEncoderLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_GPUComputePassEncoder_SetPipeline
//go:noescape
func HasFuncGPUComputePassEncoderSetPipeline(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_SetPipeline
//go:noescape
func FuncGPUComputePassEncoderSetPipeline(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUComputePassEncoder_SetPipeline
//go:noescape
func CallGPUComputePassEncoderSetPipeline(
	this js.Ref, retPtr unsafe.Pointer,
	pipeline js.Ref)

//go:wasmimport plat/js/web try_GPUComputePassEncoder_SetPipeline
//go:noescape
func TryGPUComputePassEncoderSetPipeline(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	pipeline js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUComputePassEncoder_DispatchWorkgroups
//go:noescape
func HasFuncGPUComputePassEncoderDispatchWorkgroups(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_DispatchWorkgroups
//go:noescape
func FuncGPUComputePassEncoderDispatchWorkgroups(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUComputePassEncoder_DispatchWorkgroups
//go:noescape
func CallGPUComputePassEncoderDispatchWorkgroups(
	this js.Ref, retPtr unsafe.Pointer,
	workgroupCountX uint32,
	workgroupCountY uint32,
	workgroupCountZ uint32)

//go:wasmimport plat/js/web try_GPUComputePassEncoder_DispatchWorkgroups
//go:noescape
func TryGPUComputePassEncoderDispatchWorkgroups(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	workgroupCountX uint32,
	workgroupCountY uint32,
	workgroupCountZ uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUComputePassEncoder_DispatchWorkgroups1
//go:noescape
func HasFuncGPUComputePassEncoderDispatchWorkgroups1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_DispatchWorkgroups1
//go:noescape
func FuncGPUComputePassEncoderDispatchWorkgroups1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUComputePassEncoder_DispatchWorkgroups1
//go:noescape
func CallGPUComputePassEncoderDispatchWorkgroups1(
	this js.Ref, retPtr unsafe.Pointer,
	workgroupCountX uint32,
	workgroupCountY uint32)

//go:wasmimport plat/js/web try_GPUComputePassEncoder_DispatchWorkgroups1
//go:noescape
func TryGPUComputePassEncoderDispatchWorkgroups1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	workgroupCountX uint32,
	workgroupCountY uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUComputePassEncoder_DispatchWorkgroups2
//go:noescape
func HasFuncGPUComputePassEncoderDispatchWorkgroups2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_DispatchWorkgroups2
//go:noescape
func FuncGPUComputePassEncoderDispatchWorkgroups2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUComputePassEncoder_DispatchWorkgroups2
//go:noescape
func CallGPUComputePassEncoderDispatchWorkgroups2(
	this js.Ref, retPtr unsafe.Pointer,
	workgroupCountX uint32)

//go:wasmimport plat/js/web try_GPUComputePassEncoder_DispatchWorkgroups2
//go:noescape
func TryGPUComputePassEncoderDispatchWorkgroups2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	workgroupCountX uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUComputePassEncoder_DispatchWorkgroupsIndirect
//go:noescape
func HasFuncGPUComputePassEncoderDispatchWorkgroupsIndirect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_DispatchWorkgroupsIndirect
//go:noescape
func FuncGPUComputePassEncoderDispatchWorkgroupsIndirect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUComputePassEncoder_DispatchWorkgroupsIndirect
//go:noescape
func CallGPUComputePassEncoderDispatchWorkgroupsIndirect(
	this js.Ref, retPtr unsafe.Pointer,
	indirectBuffer js.Ref,
	indirectOffset float64)

//go:wasmimport plat/js/web try_GPUComputePassEncoder_DispatchWorkgroupsIndirect
//go:noescape
func TryGPUComputePassEncoderDispatchWorkgroupsIndirect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	indirectBuffer js.Ref,
	indirectOffset float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUComputePassEncoder_End
//go:noescape
func HasFuncGPUComputePassEncoderEnd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_End
//go:noescape
func FuncGPUComputePassEncoderEnd(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUComputePassEncoder_End
//go:noescape
func CallGPUComputePassEncoderEnd(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUComputePassEncoder_End
//go:noescape
func TryGPUComputePassEncoderEnd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUComputePassEncoder_PushDebugGroup
//go:noescape
func HasFuncGPUComputePassEncoderPushDebugGroup(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_PushDebugGroup
//go:noescape
func FuncGPUComputePassEncoderPushDebugGroup(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUComputePassEncoder_PushDebugGroup
//go:noescape
func CallGPUComputePassEncoderPushDebugGroup(
	this js.Ref, retPtr unsafe.Pointer,
	groupLabel js.Ref)

//go:wasmimport plat/js/web try_GPUComputePassEncoder_PushDebugGroup
//go:noescape
func TryGPUComputePassEncoderPushDebugGroup(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	groupLabel js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUComputePassEncoder_PopDebugGroup
//go:noescape
func HasFuncGPUComputePassEncoderPopDebugGroup(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_PopDebugGroup
//go:noescape
func FuncGPUComputePassEncoderPopDebugGroup(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUComputePassEncoder_PopDebugGroup
//go:noescape
func CallGPUComputePassEncoderPopDebugGroup(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUComputePassEncoder_PopDebugGroup
//go:noescape
func TryGPUComputePassEncoderPopDebugGroup(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUComputePassEncoder_InsertDebugMarker
//go:noescape
func HasFuncGPUComputePassEncoderInsertDebugMarker(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_InsertDebugMarker
//go:noescape
func FuncGPUComputePassEncoderInsertDebugMarker(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUComputePassEncoder_InsertDebugMarker
//go:noescape
func CallGPUComputePassEncoderInsertDebugMarker(
	this js.Ref, retPtr unsafe.Pointer,
	markerLabel js.Ref)

//go:wasmimport plat/js/web try_GPUComputePassEncoder_InsertDebugMarker
//go:noescape
func TryGPUComputePassEncoderInsertDebugMarker(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	markerLabel js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUComputePassEncoder_SetBindGroup
//go:noescape
func HasFuncGPUComputePassEncoderSetBindGroup(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_SetBindGroup
//go:noescape
func FuncGPUComputePassEncoderSetBindGroup(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUComputePassEncoder_SetBindGroup
//go:noescape
func CallGPUComputePassEncoderSetBindGroup(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref,
	dynamicOffsets js.Ref)

//go:wasmimport plat/js/web try_GPUComputePassEncoder_SetBindGroup
//go:noescape
func TryGPUComputePassEncoderSetBindGroup(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref,
	dynamicOffsets js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUComputePassEncoder_SetBindGroup1
//go:noescape
func HasFuncGPUComputePassEncoderSetBindGroup1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_SetBindGroup1
//go:noescape
func FuncGPUComputePassEncoderSetBindGroup1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUComputePassEncoder_SetBindGroup1
//go:noescape
func CallGPUComputePassEncoderSetBindGroup1(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref)

//go:wasmimport plat/js/web try_GPUComputePassEncoder_SetBindGroup1
//go:noescape
func TryGPUComputePassEncoderSetBindGroup1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUComputePassEncoder_SetBindGroup2
//go:noescape
func HasFuncGPUComputePassEncoderSetBindGroup2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUComputePassEncoder_SetBindGroup2
//go:noescape
func FuncGPUComputePassEncoderSetBindGroup2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUComputePassEncoder_SetBindGroup2
//go:noescape
func CallGPUComputePassEncoderSetBindGroup2(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref,
	dynamicOffsetsData js.Ref,
	dynamicOffsetsDataStart float64,
	dynamicOffsetsDataLength uint32)

//go:wasmimport plat/js/web try_GPUComputePassEncoder_SetBindGroup2
//go:noescape
func TryGPUComputePassEncoderSetBindGroup2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref,
	dynamicOffsetsData js.Ref,
	dynamicOffsetsDataStart float64,
	dynamicOffsetsDataLength uint32) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPUCommandEncoder_Label
//go:noescape
func SetGPUCommandEncoderLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_GPUCommandEncoder_BeginRenderPass
//go:noescape
func HasFuncGPUCommandEncoderBeginRenderPass(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_BeginRenderPass
//go:noescape
func FuncGPUCommandEncoderBeginRenderPass(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_BeginRenderPass
//go:noescape
func CallGPUCommandEncoderBeginRenderPass(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUCommandEncoder_BeginRenderPass
//go:noescape
func TryGPUCommandEncoderBeginRenderPass(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_BeginComputePass
//go:noescape
func HasFuncGPUCommandEncoderBeginComputePass(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_BeginComputePass
//go:noescape
func FuncGPUCommandEncoderBeginComputePass(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_BeginComputePass
//go:noescape
func CallGPUCommandEncoderBeginComputePass(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUCommandEncoder_BeginComputePass
//go:noescape
func TryGPUCommandEncoderBeginComputePass(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_BeginComputePass1
//go:noescape
func HasFuncGPUCommandEncoderBeginComputePass1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_BeginComputePass1
//go:noescape
func FuncGPUCommandEncoderBeginComputePass1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_BeginComputePass1
//go:noescape
func CallGPUCommandEncoderBeginComputePass1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUCommandEncoder_BeginComputePass1
//go:noescape
func TryGPUCommandEncoderBeginComputePass1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_CopyBufferToBuffer
//go:noescape
func HasFuncGPUCommandEncoderCopyBufferToBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_CopyBufferToBuffer
//go:noescape
func FuncGPUCommandEncoderCopyBufferToBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_CopyBufferToBuffer
//go:noescape
func CallGPUCommandEncoderCopyBufferToBuffer(
	this js.Ref, retPtr unsafe.Pointer,
	source js.Ref,
	sourceOffset float64,
	destination js.Ref,
	destinationOffset float64,
	size float64)

//go:wasmimport plat/js/web try_GPUCommandEncoder_CopyBufferToBuffer
//go:noescape
func TryGPUCommandEncoderCopyBufferToBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source js.Ref,
	sourceOffset float64,
	destination js.Ref,
	destinationOffset float64,
	size float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_CopyBufferToTexture
//go:noescape
func HasFuncGPUCommandEncoderCopyBufferToTexture(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_CopyBufferToTexture
//go:noescape
func FuncGPUCommandEncoderCopyBufferToTexture(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_CopyBufferToTexture
//go:noescape
func CallGPUCommandEncoderCopyBufferToTexture(
	this js.Ref, retPtr unsafe.Pointer,
	source unsafe.Pointer,
	destination unsafe.Pointer,
	copySize js.Ref)

//go:wasmimport plat/js/web try_GPUCommandEncoder_CopyBufferToTexture
//go:noescape
func TryGPUCommandEncoderCopyBufferToTexture(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source unsafe.Pointer,
	destination unsafe.Pointer,
	copySize js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_CopyTextureToBuffer
//go:noescape
func HasFuncGPUCommandEncoderCopyTextureToBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_CopyTextureToBuffer
//go:noescape
func FuncGPUCommandEncoderCopyTextureToBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_CopyTextureToBuffer
//go:noescape
func CallGPUCommandEncoderCopyTextureToBuffer(
	this js.Ref, retPtr unsafe.Pointer,
	source unsafe.Pointer,
	destination unsafe.Pointer,
	copySize js.Ref)

//go:wasmimport plat/js/web try_GPUCommandEncoder_CopyTextureToBuffer
//go:noescape
func TryGPUCommandEncoderCopyTextureToBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source unsafe.Pointer,
	destination unsafe.Pointer,
	copySize js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_CopyTextureToTexture
//go:noescape
func HasFuncGPUCommandEncoderCopyTextureToTexture(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_CopyTextureToTexture
//go:noescape
func FuncGPUCommandEncoderCopyTextureToTexture(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_CopyTextureToTexture
//go:noescape
func CallGPUCommandEncoderCopyTextureToTexture(
	this js.Ref, retPtr unsafe.Pointer,
	source unsafe.Pointer,
	destination unsafe.Pointer,
	copySize js.Ref)

//go:wasmimport plat/js/web try_GPUCommandEncoder_CopyTextureToTexture
//go:noescape
func TryGPUCommandEncoderCopyTextureToTexture(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source unsafe.Pointer,
	destination unsafe.Pointer,
	copySize js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_ClearBuffer
//go:noescape
func HasFuncGPUCommandEncoderClearBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_ClearBuffer
//go:noescape
func FuncGPUCommandEncoderClearBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_ClearBuffer
//go:noescape
func CallGPUCommandEncoderClearBuffer(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref,
	offset float64,
	size float64)

//go:wasmimport plat/js/web try_GPUCommandEncoder_ClearBuffer
//go:noescape
func TryGPUCommandEncoderClearBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref,
	offset float64,
	size float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_ClearBuffer1
//go:noescape
func HasFuncGPUCommandEncoderClearBuffer1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_ClearBuffer1
//go:noescape
func FuncGPUCommandEncoderClearBuffer1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_ClearBuffer1
//go:noescape
func CallGPUCommandEncoderClearBuffer1(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref,
	offset float64)

//go:wasmimport plat/js/web try_GPUCommandEncoder_ClearBuffer1
//go:noescape
func TryGPUCommandEncoderClearBuffer1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref,
	offset float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_ClearBuffer2
//go:noescape
func HasFuncGPUCommandEncoderClearBuffer2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_ClearBuffer2
//go:noescape
func FuncGPUCommandEncoderClearBuffer2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_ClearBuffer2
//go:noescape
func CallGPUCommandEncoderClearBuffer2(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref)

//go:wasmimport plat/js/web try_GPUCommandEncoder_ClearBuffer2
//go:noescape
func TryGPUCommandEncoderClearBuffer2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_WriteTimestamp
//go:noescape
func HasFuncGPUCommandEncoderWriteTimestamp(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_WriteTimestamp
//go:noescape
func FuncGPUCommandEncoderWriteTimestamp(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_WriteTimestamp
//go:noescape
func CallGPUCommandEncoderWriteTimestamp(
	this js.Ref, retPtr unsafe.Pointer,
	querySet js.Ref,
	queryIndex uint32)

//go:wasmimport plat/js/web try_GPUCommandEncoder_WriteTimestamp
//go:noescape
func TryGPUCommandEncoderWriteTimestamp(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	querySet js.Ref,
	queryIndex uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_ResolveQuerySet
//go:noescape
func HasFuncGPUCommandEncoderResolveQuerySet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_ResolveQuerySet
//go:noescape
func FuncGPUCommandEncoderResolveQuerySet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_ResolveQuerySet
//go:noescape
func CallGPUCommandEncoderResolveQuerySet(
	this js.Ref, retPtr unsafe.Pointer,
	querySet js.Ref,
	firstQuery uint32,
	queryCount uint32,
	destination js.Ref,
	destinationOffset float64)

//go:wasmimport plat/js/web try_GPUCommandEncoder_ResolveQuerySet
//go:noescape
func TryGPUCommandEncoderResolveQuerySet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	querySet js.Ref,
	firstQuery uint32,
	queryCount uint32,
	destination js.Ref,
	destinationOffset float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_Finish
//go:noescape
func HasFuncGPUCommandEncoderFinish(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_Finish
//go:noescape
func FuncGPUCommandEncoderFinish(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_Finish
//go:noescape
func CallGPUCommandEncoderFinish(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUCommandEncoder_Finish
//go:noescape
func TryGPUCommandEncoderFinish(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_Finish1
//go:noescape
func HasFuncGPUCommandEncoderFinish1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_Finish1
//go:noescape
func FuncGPUCommandEncoderFinish1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_Finish1
//go:noescape
func CallGPUCommandEncoderFinish1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUCommandEncoder_Finish1
//go:noescape
func TryGPUCommandEncoderFinish1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_PushDebugGroup
//go:noescape
func HasFuncGPUCommandEncoderPushDebugGroup(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_PushDebugGroup
//go:noescape
func FuncGPUCommandEncoderPushDebugGroup(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_PushDebugGroup
//go:noescape
func CallGPUCommandEncoderPushDebugGroup(
	this js.Ref, retPtr unsafe.Pointer,
	groupLabel js.Ref)

//go:wasmimport plat/js/web try_GPUCommandEncoder_PushDebugGroup
//go:noescape
func TryGPUCommandEncoderPushDebugGroup(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	groupLabel js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_PopDebugGroup
//go:noescape
func HasFuncGPUCommandEncoderPopDebugGroup(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_PopDebugGroup
//go:noescape
func FuncGPUCommandEncoderPopDebugGroup(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_PopDebugGroup
//go:noescape
func CallGPUCommandEncoderPopDebugGroup(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUCommandEncoder_PopDebugGroup
//go:noescape
func TryGPUCommandEncoderPopDebugGroup(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCommandEncoder_InsertDebugMarker
//go:noescape
func HasFuncGPUCommandEncoderInsertDebugMarker(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCommandEncoder_InsertDebugMarker
//go:noescape
func FuncGPUCommandEncoderInsertDebugMarker(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCommandEncoder_InsertDebugMarker
//go:noescape
func CallGPUCommandEncoderInsertDebugMarker(
	this js.Ref, retPtr unsafe.Pointer,
	markerLabel js.Ref)

//go:wasmimport plat/js/web try_GPUCommandEncoder_InsertDebugMarker
//go:noescape
func TryGPUCommandEncoderInsertDebugMarker(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	markerLabel js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPURenderBundleEncoder_Label
//go:noescape
func SetGPURenderBundleEncoderLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_Finish
//go:noescape
func HasFuncGPURenderBundleEncoderFinish(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_Finish
//go:noescape
func FuncGPURenderBundleEncoderFinish(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_Finish
//go:noescape
func CallGPURenderBundleEncoderFinish(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_Finish
//go:noescape
func TryGPURenderBundleEncoderFinish(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_Finish1
//go:noescape
func HasFuncGPURenderBundleEncoderFinish1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_Finish1
//go:noescape
func FuncGPURenderBundleEncoderFinish1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_Finish1
//go:noescape
func CallGPURenderBundleEncoderFinish1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_Finish1
//go:noescape
func TryGPURenderBundleEncoderFinish1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_SetPipeline
//go:noescape
func HasFuncGPURenderBundleEncoderSetPipeline(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetPipeline
//go:noescape
func FuncGPURenderBundleEncoderSetPipeline(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetPipeline
//go:noescape
func CallGPURenderBundleEncoderSetPipeline(
	this js.Ref, retPtr unsafe.Pointer,
	pipeline js.Ref)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_SetPipeline
//go:noescape
func TryGPURenderBundleEncoderSetPipeline(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	pipeline js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_SetIndexBuffer
//go:noescape
func HasFuncGPURenderBundleEncoderSetIndexBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetIndexBuffer
//go:noescape
func FuncGPURenderBundleEncoderSetIndexBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetIndexBuffer
//go:noescape
func CallGPURenderBundleEncoderSetIndexBuffer(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref,
	indexFormat uint32,
	offset float64,
	size float64)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_SetIndexBuffer
//go:noescape
func TryGPURenderBundleEncoderSetIndexBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref,
	indexFormat uint32,
	offset float64,
	size float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_SetIndexBuffer1
//go:noescape
func HasFuncGPURenderBundleEncoderSetIndexBuffer1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetIndexBuffer1
//go:noescape
func FuncGPURenderBundleEncoderSetIndexBuffer1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetIndexBuffer1
//go:noescape
func CallGPURenderBundleEncoderSetIndexBuffer1(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref,
	indexFormat uint32,
	offset float64)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_SetIndexBuffer1
//go:noescape
func TryGPURenderBundleEncoderSetIndexBuffer1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref,
	indexFormat uint32,
	offset float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_SetIndexBuffer2
//go:noescape
func HasFuncGPURenderBundleEncoderSetIndexBuffer2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetIndexBuffer2
//go:noescape
func FuncGPURenderBundleEncoderSetIndexBuffer2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetIndexBuffer2
//go:noescape
func CallGPURenderBundleEncoderSetIndexBuffer2(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref,
	indexFormat uint32)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_SetIndexBuffer2
//go:noescape
func TryGPURenderBundleEncoderSetIndexBuffer2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref,
	indexFormat uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_SetVertexBuffer
//go:noescape
func HasFuncGPURenderBundleEncoderSetVertexBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetVertexBuffer
//go:noescape
func FuncGPURenderBundleEncoderSetVertexBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetVertexBuffer
//go:noescape
func CallGPURenderBundleEncoderSetVertexBuffer(
	this js.Ref, retPtr unsafe.Pointer,
	slot uint32,
	buffer js.Ref,
	offset float64,
	size float64)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_SetVertexBuffer
//go:noescape
func TryGPURenderBundleEncoderSetVertexBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	slot uint32,
	buffer js.Ref,
	offset float64,
	size float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_SetVertexBuffer1
//go:noescape
func HasFuncGPURenderBundleEncoderSetVertexBuffer1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetVertexBuffer1
//go:noescape
func FuncGPURenderBundleEncoderSetVertexBuffer1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetVertexBuffer1
//go:noescape
func CallGPURenderBundleEncoderSetVertexBuffer1(
	this js.Ref, retPtr unsafe.Pointer,
	slot uint32,
	buffer js.Ref,
	offset float64)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_SetVertexBuffer1
//go:noescape
func TryGPURenderBundleEncoderSetVertexBuffer1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	slot uint32,
	buffer js.Ref,
	offset float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_SetVertexBuffer2
//go:noescape
func HasFuncGPURenderBundleEncoderSetVertexBuffer2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetVertexBuffer2
//go:noescape
func FuncGPURenderBundleEncoderSetVertexBuffer2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetVertexBuffer2
//go:noescape
func CallGPURenderBundleEncoderSetVertexBuffer2(
	this js.Ref, retPtr unsafe.Pointer,
	slot uint32,
	buffer js.Ref)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_SetVertexBuffer2
//go:noescape
func TryGPURenderBundleEncoderSetVertexBuffer2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	slot uint32,
	buffer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_Draw
//go:noescape
func HasFuncGPURenderBundleEncoderDraw(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_Draw
//go:noescape
func FuncGPURenderBundleEncoderDraw(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_Draw
//go:noescape
func CallGPURenderBundleEncoderDraw(
	this js.Ref, retPtr unsafe.Pointer,
	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32,
	firstInstance uint32)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_Draw
//go:noescape
func TryGPURenderBundleEncoderDraw(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32,
	firstInstance uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_Draw1
//go:noescape
func HasFuncGPURenderBundleEncoderDraw1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_Draw1
//go:noescape
func FuncGPURenderBundleEncoderDraw1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_Draw1
//go:noescape
func CallGPURenderBundleEncoderDraw1(
	this js.Ref, retPtr unsafe.Pointer,
	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_Draw1
//go:noescape
func TryGPURenderBundleEncoderDraw1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_Draw2
//go:noescape
func HasFuncGPURenderBundleEncoderDraw2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_Draw2
//go:noescape
func FuncGPURenderBundleEncoderDraw2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_Draw2
//go:noescape
func CallGPURenderBundleEncoderDraw2(
	this js.Ref, retPtr unsafe.Pointer,
	vertexCount uint32,
	instanceCount uint32)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_Draw2
//go:noescape
func TryGPURenderBundleEncoderDraw2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	vertexCount uint32,
	instanceCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_Draw3
//go:noescape
func HasFuncGPURenderBundleEncoderDraw3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_Draw3
//go:noescape
func FuncGPURenderBundleEncoderDraw3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_Draw3
//go:noescape
func CallGPURenderBundleEncoderDraw3(
	this js.Ref, retPtr unsafe.Pointer,
	vertexCount uint32)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_Draw3
//go:noescape
func TryGPURenderBundleEncoderDraw3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	vertexCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_DrawIndexed
//go:noescape
func HasFuncGPURenderBundleEncoderDrawIndexed(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_DrawIndexed
//go:noescape
func FuncGPURenderBundleEncoderDrawIndexed(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_DrawIndexed
//go:noescape
func CallGPURenderBundleEncoderDrawIndexed(
	this js.Ref, retPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	baseVertex int32,
	firstInstance uint32)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_DrawIndexed
//go:noescape
func TryGPURenderBundleEncoderDrawIndexed(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	baseVertex int32,
	firstInstance uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_DrawIndexed1
//go:noescape
func HasFuncGPURenderBundleEncoderDrawIndexed1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_DrawIndexed1
//go:noescape
func FuncGPURenderBundleEncoderDrawIndexed1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_DrawIndexed1
//go:noescape
func CallGPURenderBundleEncoderDrawIndexed1(
	this js.Ref, retPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	baseVertex int32)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_DrawIndexed1
//go:noescape
func TryGPURenderBundleEncoderDrawIndexed1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	baseVertex int32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_DrawIndexed2
//go:noescape
func HasFuncGPURenderBundleEncoderDrawIndexed2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_DrawIndexed2
//go:noescape
func FuncGPURenderBundleEncoderDrawIndexed2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_DrawIndexed2
//go:noescape
func CallGPURenderBundleEncoderDrawIndexed2(
	this js.Ref, retPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_DrawIndexed2
//go:noescape
func TryGPURenderBundleEncoderDrawIndexed2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_DrawIndexed3
//go:noescape
func HasFuncGPURenderBundleEncoderDrawIndexed3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_DrawIndexed3
//go:noescape
func FuncGPURenderBundleEncoderDrawIndexed3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_DrawIndexed3
//go:noescape
func CallGPURenderBundleEncoderDrawIndexed3(
	this js.Ref, retPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_DrawIndexed3
//go:noescape
func TryGPURenderBundleEncoderDrawIndexed3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	indexCount uint32,
	instanceCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_DrawIndexed4
//go:noescape
func HasFuncGPURenderBundleEncoderDrawIndexed4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_DrawIndexed4
//go:noescape
func FuncGPURenderBundleEncoderDrawIndexed4(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_DrawIndexed4
//go:noescape
func CallGPURenderBundleEncoderDrawIndexed4(
	this js.Ref, retPtr unsafe.Pointer,
	indexCount uint32)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_DrawIndexed4
//go:noescape
func TryGPURenderBundleEncoderDrawIndexed4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	indexCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_DrawIndirect
//go:noescape
func HasFuncGPURenderBundleEncoderDrawIndirect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_DrawIndirect
//go:noescape
func FuncGPURenderBundleEncoderDrawIndirect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_DrawIndirect
//go:noescape
func CallGPURenderBundleEncoderDrawIndirect(
	this js.Ref, retPtr unsafe.Pointer,
	indirectBuffer js.Ref,
	indirectOffset float64)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_DrawIndirect
//go:noescape
func TryGPURenderBundleEncoderDrawIndirect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	indirectBuffer js.Ref,
	indirectOffset float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_DrawIndexedIndirect
//go:noescape
func HasFuncGPURenderBundleEncoderDrawIndexedIndirect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_DrawIndexedIndirect
//go:noescape
func FuncGPURenderBundleEncoderDrawIndexedIndirect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_DrawIndexedIndirect
//go:noescape
func CallGPURenderBundleEncoderDrawIndexedIndirect(
	this js.Ref, retPtr unsafe.Pointer,
	indirectBuffer js.Ref,
	indirectOffset float64)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_DrawIndexedIndirect
//go:noescape
func TryGPURenderBundleEncoderDrawIndexedIndirect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	indirectBuffer js.Ref,
	indirectOffset float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_SetBindGroup
//go:noescape
func HasFuncGPURenderBundleEncoderSetBindGroup(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetBindGroup
//go:noescape
func FuncGPURenderBundleEncoderSetBindGroup(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetBindGroup
//go:noescape
func CallGPURenderBundleEncoderSetBindGroup(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref,
	dynamicOffsets js.Ref)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_SetBindGroup
//go:noescape
func TryGPURenderBundleEncoderSetBindGroup(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref,
	dynamicOffsets js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_SetBindGroup1
//go:noescape
func HasFuncGPURenderBundleEncoderSetBindGroup1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetBindGroup1
//go:noescape
func FuncGPURenderBundleEncoderSetBindGroup1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetBindGroup1
//go:noescape
func CallGPURenderBundleEncoderSetBindGroup1(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_SetBindGroup1
//go:noescape
func TryGPURenderBundleEncoderSetBindGroup1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_SetBindGroup2
//go:noescape
func HasFuncGPURenderBundleEncoderSetBindGroup2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_SetBindGroup2
//go:noescape
func FuncGPURenderBundleEncoderSetBindGroup2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_SetBindGroup2
//go:noescape
func CallGPURenderBundleEncoderSetBindGroup2(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref,
	dynamicOffsetsData js.Ref,
	dynamicOffsetsDataStart float64,
	dynamicOffsetsDataLength uint32)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_SetBindGroup2
//go:noescape
func TryGPURenderBundleEncoderSetBindGroup2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	bindGroup js.Ref,
	dynamicOffsetsData js.Ref,
	dynamicOffsetsDataStart float64,
	dynamicOffsetsDataLength uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_PushDebugGroup
//go:noescape
func HasFuncGPURenderBundleEncoderPushDebugGroup(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_PushDebugGroup
//go:noescape
func FuncGPURenderBundleEncoderPushDebugGroup(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_PushDebugGroup
//go:noescape
func CallGPURenderBundleEncoderPushDebugGroup(
	this js.Ref, retPtr unsafe.Pointer,
	groupLabel js.Ref)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_PushDebugGroup
//go:noescape
func TryGPURenderBundleEncoderPushDebugGroup(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	groupLabel js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_PopDebugGroup
//go:noescape
func HasFuncGPURenderBundleEncoderPopDebugGroup(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_PopDebugGroup
//go:noescape
func FuncGPURenderBundleEncoderPopDebugGroup(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_PopDebugGroup
//go:noescape
func CallGPURenderBundleEncoderPopDebugGroup(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_PopDebugGroup
//go:noescape
func TryGPURenderBundleEncoderPopDebugGroup(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPURenderBundleEncoder_InsertDebugMarker
//go:noescape
func HasFuncGPURenderBundleEncoderInsertDebugMarker(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderBundleEncoder_InsertDebugMarker
//go:noescape
func FuncGPURenderBundleEncoderInsertDebugMarker(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderBundleEncoder_InsertDebugMarker
//go:noescape
func CallGPURenderBundleEncoderInsertDebugMarker(
	this js.Ref, retPtr unsafe.Pointer,
	markerLabel js.Ref)

//go:wasmimport plat/js/web try_GPURenderBundleEncoder_InsertDebugMarker
//go:noescape
func TryGPURenderBundleEncoderInsertDebugMarker(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	markerLabel js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxTextureDimension1D
//go:noescape
func GetGPUSupportedLimitsMaxTextureDimension1D(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxTextureDimension2D
//go:noescape
func GetGPUSupportedLimitsMaxTextureDimension2D(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxTextureDimension3D
//go:noescape
func GetGPUSupportedLimitsMaxTextureDimension3D(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxTextureArrayLayers
//go:noescape
func GetGPUSupportedLimitsMaxTextureArrayLayers(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxBindGroups
//go:noescape
func GetGPUSupportedLimitsMaxBindGroups(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxBindGroupsPlusVertexBuffers
//go:noescape
func GetGPUSupportedLimitsMaxBindGroupsPlusVertexBuffers(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxBindingsPerBindGroup
//go:noescape
func GetGPUSupportedLimitsMaxBindingsPerBindGroup(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxDynamicUniformBuffersPerPipelineLayout
//go:noescape
func GetGPUSupportedLimitsMaxDynamicUniformBuffersPerPipelineLayout(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxDynamicStorageBuffersPerPipelineLayout
//go:noescape
func GetGPUSupportedLimitsMaxDynamicStorageBuffersPerPipelineLayout(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxSampledTexturesPerShaderStage
//go:noescape
func GetGPUSupportedLimitsMaxSampledTexturesPerShaderStage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxSamplersPerShaderStage
//go:noescape
func GetGPUSupportedLimitsMaxSamplersPerShaderStage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxStorageBuffersPerShaderStage
//go:noescape
func GetGPUSupportedLimitsMaxStorageBuffersPerShaderStage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxStorageTexturesPerShaderStage
//go:noescape
func GetGPUSupportedLimitsMaxStorageTexturesPerShaderStage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxUniformBuffersPerShaderStage
//go:noescape
func GetGPUSupportedLimitsMaxUniformBuffersPerShaderStage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxUniformBufferBindingSize
//go:noescape
func GetGPUSupportedLimitsMaxUniformBufferBindingSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxStorageBufferBindingSize
//go:noescape
func GetGPUSupportedLimitsMaxStorageBufferBindingSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MinUniformBufferOffsetAlignment
//go:noescape
func GetGPUSupportedLimitsMinUniformBufferOffsetAlignment(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MinStorageBufferOffsetAlignment
//go:noescape
func GetGPUSupportedLimitsMinStorageBufferOffsetAlignment(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxVertexBuffers
//go:noescape
func GetGPUSupportedLimitsMaxVertexBuffers(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxBufferSize
//go:noescape
func GetGPUSupportedLimitsMaxBufferSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxVertexAttributes
//go:noescape
func GetGPUSupportedLimitsMaxVertexAttributes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxVertexBufferArrayStride
//go:noescape
func GetGPUSupportedLimitsMaxVertexBufferArrayStride(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxInterStageShaderComponents
//go:noescape
func GetGPUSupportedLimitsMaxInterStageShaderComponents(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxInterStageShaderVariables
//go:noescape
func GetGPUSupportedLimitsMaxInterStageShaderVariables(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxColorAttachments
//go:noescape
func GetGPUSupportedLimitsMaxColorAttachments(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxColorAttachmentBytesPerSample
//go:noescape
func GetGPUSupportedLimitsMaxColorAttachmentBytesPerSample(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxComputeWorkgroupStorageSize
//go:noescape
func GetGPUSupportedLimitsMaxComputeWorkgroupStorageSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxComputeInvocationsPerWorkgroup
//go:noescape
func GetGPUSupportedLimitsMaxComputeInvocationsPerWorkgroup(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxComputeWorkgroupSizeX
//go:noescape
func GetGPUSupportedLimitsMaxComputeWorkgroupSizeX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxComputeWorkgroupSizeY
//go:noescape
func GetGPUSupportedLimitsMaxComputeWorkgroupSizeY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxComputeWorkgroupSizeZ
//go:noescape
func GetGPUSupportedLimitsMaxComputeWorkgroupSizeZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUSupportedLimits_MaxComputeWorkgroupsPerDimension
//go:noescape
func GetGPUSupportedLimitsMaxComputeWorkgroupsPerDimension(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPUQueue_Label
//go:noescape
func SetGPUQueueLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_GPUQueue_Submit
//go:noescape
func HasFuncGPUQueueSubmit(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUQueue_Submit
//go:noescape
func FuncGPUQueueSubmit(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUQueue_Submit
//go:noescape
func CallGPUQueueSubmit(
	this js.Ref, retPtr unsafe.Pointer,
	commandBuffers js.Ref)

//go:wasmimport plat/js/web try_GPUQueue_Submit
//go:noescape
func TryGPUQueueSubmit(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	commandBuffers js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUQueue_OnSubmittedWorkDone
//go:noescape
func HasFuncGPUQueueOnSubmittedWorkDone(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUQueue_OnSubmittedWorkDone
//go:noescape
func FuncGPUQueueOnSubmittedWorkDone(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUQueue_OnSubmittedWorkDone
//go:noescape
func CallGPUQueueOnSubmittedWorkDone(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUQueue_OnSubmittedWorkDone
//go:noescape
func TryGPUQueueOnSubmittedWorkDone(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUQueue_WriteBuffer
//go:noescape
func HasFuncGPUQueueWriteBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUQueue_WriteBuffer
//go:noescape
func FuncGPUQueueWriteBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUQueue_WriteBuffer
//go:noescape
func CallGPUQueueWriteBuffer(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref,
	bufferOffset float64,
	data js.Ref,
	dataOffset float64,
	size float64)

//go:wasmimport plat/js/web try_GPUQueue_WriteBuffer
//go:noescape
func TryGPUQueueWriteBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref,
	bufferOffset float64,
	data js.Ref,
	dataOffset float64,
	size float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUQueue_WriteBuffer1
//go:noescape
func HasFuncGPUQueueWriteBuffer1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUQueue_WriteBuffer1
//go:noescape
func FuncGPUQueueWriteBuffer1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUQueue_WriteBuffer1
//go:noescape
func CallGPUQueueWriteBuffer1(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref,
	bufferOffset float64,
	data js.Ref,
	dataOffset float64)

//go:wasmimport plat/js/web try_GPUQueue_WriteBuffer1
//go:noescape
func TryGPUQueueWriteBuffer1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref,
	bufferOffset float64,
	data js.Ref,
	dataOffset float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUQueue_WriteBuffer2
//go:noescape
func HasFuncGPUQueueWriteBuffer2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUQueue_WriteBuffer2
//go:noescape
func FuncGPUQueueWriteBuffer2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUQueue_WriteBuffer2
//go:noescape
func CallGPUQueueWriteBuffer2(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref,
	bufferOffset float64,
	data js.Ref)

//go:wasmimport plat/js/web try_GPUQueue_WriteBuffer2
//go:noescape
func TryGPUQueueWriteBuffer2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref,
	bufferOffset float64,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUQueue_WriteTexture
//go:noescape
func HasFuncGPUQueueWriteTexture(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUQueue_WriteTexture
//go:noescape
func FuncGPUQueueWriteTexture(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUQueue_WriteTexture
//go:noescape
func CallGPUQueueWriteTexture(
	this js.Ref, retPtr unsafe.Pointer,
	destination unsafe.Pointer,
	data js.Ref,
	dataLayout unsafe.Pointer,
	size js.Ref)

//go:wasmimport plat/js/web try_GPUQueue_WriteTexture
//go:noescape
func TryGPUQueueWriteTexture(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destination unsafe.Pointer,
	data js.Ref,
	dataLayout unsafe.Pointer,
	size js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUQueue_CopyExternalImageToTexture
//go:noescape
func HasFuncGPUQueueCopyExternalImageToTexture(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUQueue_CopyExternalImageToTexture
//go:noescape
func FuncGPUQueueCopyExternalImageToTexture(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUQueue_CopyExternalImageToTexture
//go:noescape
func CallGPUQueueCopyExternalImageToTexture(
	this js.Ref, retPtr unsafe.Pointer,
	source unsafe.Pointer,
	destination unsafe.Pointer,
	copySize js.Ref)

//go:wasmimport plat/js/web try_GPUQueue_CopyExternalImageToTexture
//go:noescape
func TryGPUQueueCopyExternalImageToTexture(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source unsafe.Pointer,
	destination unsafe.Pointer,
	copySize js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_GPUDeviceLostReason
//go:noescape
func ConstOfGPUDeviceLostReason(str js.Ref) uint32

//go:wasmimport plat/js/web get_GPUDeviceLostInfo_Reason
//go:noescape
func GetGPUDeviceLostInfoReason(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUDeviceLostInfo_Message
//go:noescape
func GetGPUDeviceLostInfoMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUDevice_Features
//go:noescape
func GetGPUDeviceFeatures(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUDevice_Limits
//go:noescape
func GetGPUDeviceLimits(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUDevice_Queue
//go:noescape
func GetGPUDeviceQueue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUDevice_Lost
//go:noescape
func GetGPUDeviceLost(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUDevice_Label
//go:noescape
func GetGPUDeviceLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPUDevice_Label
//go:noescape
func SetGPUDeviceLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_GPUDevice_Destroy
//go:noescape
func HasFuncGPUDeviceDestroy(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_Destroy
//go:noescape
func FuncGPUDeviceDestroy(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_Destroy
//go:noescape
func CallGPUDeviceDestroy(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_Destroy
//go:noescape
func TryGPUDeviceDestroy(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreateBuffer
//go:noescape
func HasFuncGPUDeviceCreateBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateBuffer
//go:noescape
func FuncGPUDeviceCreateBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreateBuffer
//go:noescape
func CallGPUDeviceCreateBuffer(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreateBuffer
//go:noescape
func TryGPUDeviceCreateBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreateTexture
//go:noescape
func HasFuncGPUDeviceCreateTexture(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateTexture
//go:noescape
func FuncGPUDeviceCreateTexture(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreateTexture
//go:noescape
func CallGPUDeviceCreateTexture(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreateTexture
//go:noescape
func TryGPUDeviceCreateTexture(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreateSampler
//go:noescape
func HasFuncGPUDeviceCreateSampler(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateSampler
//go:noescape
func FuncGPUDeviceCreateSampler(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreateSampler
//go:noescape
func CallGPUDeviceCreateSampler(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreateSampler
//go:noescape
func TryGPUDeviceCreateSampler(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreateSampler1
//go:noescape
func HasFuncGPUDeviceCreateSampler1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateSampler1
//go:noescape
func FuncGPUDeviceCreateSampler1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreateSampler1
//go:noescape
func CallGPUDeviceCreateSampler1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreateSampler1
//go:noescape
func TryGPUDeviceCreateSampler1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_ImportExternalTexture
//go:noescape
func HasFuncGPUDeviceImportExternalTexture(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_ImportExternalTexture
//go:noescape
func FuncGPUDeviceImportExternalTexture(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_ImportExternalTexture
//go:noescape
func CallGPUDeviceImportExternalTexture(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_ImportExternalTexture
//go:noescape
func TryGPUDeviceImportExternalTexture(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreateBindGroupLayout
//go:noescape
func HasFuncGPUDeviceCreateBindGroupLayout(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateBindGroupLayout
//go:noescape
func FuncGPUDeviceCreateBindGroupLayout(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreateBindGroupLayout
//go:noescape
func CallGPUDeviceCreateBindGroupLayout(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreateBindGroupLayout
//go:noescape
func TryGPUDeviceCreateBindGroupLayout(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreatePipelineLayout
//go:noescape
func HasFuncGPUDeviceCreatePipelineLayout(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreatePipelineLayout
//go:noescape
func FuncGPUDeviceCreatePipelineLayout(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreatePipelineLayout
//go:noescape
func CallGPUDeviceCreatePipelineLayout(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreatePipelineLayout
//go:noescape
func TryGPUDeviceCreatePipelineLayout(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreateBindGroup
//go:noescape
func HasFuncGPUDeviceCreateBindGroup(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateBindGroup
//go:noescape
func FuncGPUDeviceCreateBindGroup(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreateBindGroup
//go:noescape
func CallGPUDeviceCreateBindGroup(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreateBindGroup
//go:noescape
func TryGPUDeviceCreateBindGroup(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreateShaderModule
//go:noescape
func HasFuncGPUDeviceCreateShaderModule(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateShaderModule
//go:noescape
func FuncGPUDeviceCreateShaderModule(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreateShaderModule
//go:noescape
func CallGPUDeviceCreateShaderModule(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreateShaderModule
//go:noescape
func TryGPUDeviceCreateShaderModule(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreateComputePipeline
//go:noescape
func HasFuncGPUDeviceCreateComputePipeline(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateComputePipeline
//go:noescape
func FuncGPUDeviceCreateComputePipeline(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreateComputePipeline
//go:noescape
func CallGPUDeviceCreateComputePipeline(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreateComputePipeline
//go:noescape
func TryGPUDeviceCreateComputePipeline(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreateRenderPipeline
//go:noescape
func HasFuncGPUDeviceCreateRenderPipeline(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateRenderPipeline
//go:noescape
func FuncGPUDeviceCreateRenderPipeline(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreateRenderPipeline
//go:noescape
func CallGPUDeviceCreateRenderPipeline(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreateRenderPipeline
//go:noescape
func TryGPUDeviceCreateRenderPipeline(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreateComputePipelineAsync
//go:noescape
func HasFuncGPUDeviceCreateComputePipelineAsync(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateComputePipelineAsync
//go:noescape
func FuncGPUDeviceCreateComputePipelineAsync(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreateComputePipelineAsync
//go:noescape
func CallGPUDeviceCreateComputePipelineAsync(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreateComputePipelineAsync
//go:noescape
func TryGPUDeviceCreateComputePipelineAsync(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreateRenderPipelineAsync
//go:noescape
func HasFuncGPUDeviceCreateRenderPipelineAsync(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateRenderPipelineAsync
//go:noescape
func FuncGPUDeviceCreateRenderPipelineAsync(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreateRenderPipelineAsync
//go:noescape
func CallGPUDeviceCreateRenderPipelineAsync(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreateRenderPipelineAsync
//go:noescape
func TryGPUDeviceCreateRenderPipelineAsync(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreateCommandEncoder
//go:noescape
func HasFuncGPUDeviceCreateCommandEncoder(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateCommandEncoder
//go:noescape
func FuncGPUDeviceCreateCommandEncoder(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreateCommandEncoder
//go:noescape
func CallGPUDeviceCreateCommandEncoder(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreateCommandEncoder
//go:noescape
func TryGPUDeviceCreateCommandEncoder(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreateCommandEncoder1
//go:noescape
func HasFuncGPUDeviceCreateCommandEncoder1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateCommandEncoder1
//go:noescape
func FuncGPUDeviceCreateCommandEncoder1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreateCommandEncoder1
//go:noescape
func CallGPUDeviceCreateCommandEncoder1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreateCommandEncoder1
//go:noescape
func TryGPUDeviceCreateCommandEncoder1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreateRenderBundleEncoder
//go:noescape
func HasFuncGPUDeviceCreateRenderBundleEncoder(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateRenderBundleEncoder
//go:noescape
func FuncGPUDeviceCreateRenderBundleEncoder(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreateRenderBundleEncoder
//go:noescape
func CallGPUDeviceCreateRenderBundleEncoder(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreateRenderBundleEncoder
//go:noescape
func TryGPUDeviceCreateRenderBundleEncoder(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_CreateQuerySet
//go:noescape
func HasFuncGPUDeviceCreateQuerySet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_CreateQuerySet
//go:noescape
func FuncGPUDeviceCreateQuerySet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_CreateQuerySet
//go:noescape
func CallGPUDeviceCreateQuerySet(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_CreateQuerySet
//go:noescape
func TryGPUDeviceCreateQuerySet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_PushErrorScope
//go:noescape
func HasFuncGPUDevicePushErrorScope(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_PushErrorScope
//go:noescape
func FuncGPUDevicePushErrorScope(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_PushErrorScope
//go:noescape
func CallGPUDevicePushErrorScope(
	this js.Ref, retPtr unsafe.Pointer,
	filter uint32)

//go:wasmimport plat/js/web try_GPUDevice_PushErrorScope
//go:noescape
func TryGPUDevicePushErrorScope(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUDevice_PopErrorScope
//go:noescape
func HasFuncGPUDevicePopErrorScope(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUDevice_PopErrorScope
//go:noescape
func FuncGPUDevicePopErrorScope(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUDevice_PopErrorScope
//go:noescape
func CallGPUDevicePopErrorScope(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUDevice_PopErrorScope
//go:noescape
func TryGPUDevicePopErrorScope(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCanvasContext_Configure
//go:noescape
func HasFuncGPUCanvasContextConfigure(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCanvasContext_Configure
//go:noescape
func FuncGPUCanvasContextConfigure(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCanvasContext_Configure
//go:noescape
func CallGPUCanvasContextConfigure(
	this js.Ref, retPtr unsafe.Pointer,
	configuration unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUCanvasContext_Configure
//go:noescape
func TryGPUCanvasContextConfigure(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	configuration unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCanvasContext_Unconfigure
//go:noescape
func HasFuncGPUCanvasContextUnconfigure(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCanvasContext_Unconfigure
//go:noescape
func FuncGPUCanvasContextUnconfigure(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCanvasContext_Unconfigure
//go:noescape
func CallGPUCanvasContextUnconfigure(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUCanvasContext_Unconfigure
//go:noescape
func TryGPUCanvasContextUnconfigure(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUCanvasContext_GetCurrentTexture
//go:noescape
func HasFuncGPUCanvasContextGetCurrentTexture(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUCanvasContext_GetCurrentTexture
//go:noescape
func FuncGPUCanvasContextGetCurrentTexture(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUCanvasContext_GetCurrentTexture
//go:noescape
func CallGPUCanvasContextGetCurrentTexture(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUCanvasContext_GetCurrentTexture
//go:noescape
func TryGPUCanvasContextGetCurrentTexture(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
