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

//go:wasmimport plat/js/web constof_GPUAutoLayoutMode
//go:noescape
func ConstOfGPUAutoLayoutMode(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPUShaderModuleCompilationHint
//go:noescape
func GPUShaderModuleCompilationHintJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUShaderModuleCompilationHint
//go:noescape
func GPUShaderModuleCompilationHintJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUShaderModuleDescriptor
//go:noescape
func GPUShaderModuleDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUShaderModuleDescriptor
//go:noescape
func GPUShaderModuleDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPUComputePipeline_Label
//go:noescape
func GetGPUComputePipelineLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPUComputePipeline_Label
//go:noescape
func SetGPUComputePipelineLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_GPUComputePipeline_GetBindGroupLayout
//go:noescape
func HasFuncGPUComputePipelineGetBindGroupLayout(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUComputePipeline_GetBindGroupLayout
//go:noescape
func FuncGPUComputePipelineGetBindGroupLayout(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPUComputePipeline_GetBindGroupLayout
//go:noescape
func CallGPUComputePipelineGetBindGroupLayout(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_GPUComputePipeline_GetBindGroupLayout
//go:noescape
func TryGPUComputePipelineGetBindGroupLayout(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web store_GPUProgrammableStage
//go:noescape
func GPUProgrammableStageJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUProgrammableStage
//go:noescape
func GPUProgrammableStageJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUComputePipelineDescriptor
//go:noescape
func GPUComputePipelineDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUComputePipelineDescriptor
//go:noescape
func GPUComputePipelineDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPURenderPipeline_Label
//go:noescape
func GetGPURenderPipelineLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPURenderPipeline_Label
//go:noescape
func SetGPURenderPipelineLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_GPURenderPipeline_GetBindGroupLayout
//go:noescape
func HasFuncGPURenderPipelineGetBindGroupLayout(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPURenderPipeline_GetBindGroupLayout
//go:noescape
func FuncGPURenderPipelineGetBindGroupLayout(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GPURenderPipeline_GetBindGroupLayout
//go:noescape
func CallGPURenderPipelineGetBindGroupLayout(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_GPURenderPipeline_GetBindGroupLayout
//go:noescape
func TryGPURenderPipelineGetBindGroupLayout(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web constof_GPUVertexStepMode
//go:noescape
func ConstOfGPUVertexStepMode(str js.Ref) uint32

//go:wasmimport plat/js/web constof_GPUVertexFormat
//go:noescape
func ConstOfGPUVertexFormat(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPUVertexAttribute
//go:noescape
func GPUVertexAttributeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUVertexAttribute
//go:noescape
func GPUVertexAttributeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUVertexBufferLayout
//go:noescape
func GPUVertexBufferLayoutJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUVertexBufferLayout
//go:noescape
func GPUVertexBufferLayoutJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUVertexState
//go:noescape
func GPUVertexStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUVertexState
//go:noescape
func GPUVertexStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GPUPrimitiveTopology
//go:noescape
func ConstOfGPUPrimitiveTopology(str js.Ref) uint32

//go:wasmimport plat/js/web constof_GPUIndexFormat
//go:noescape
func ConstOfGPUIndexFormat(str js.Ref) uint32

//go:wasmimport plat/js/web constof_GPUFrontFace
//go:noescape
func ConstOfGPUFrontFace(str js.Ref) uint32

//go:wasmimport plat/js/web constof_GPUCullMode
//go:noescape
func ConstOfGPUCullMode(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPUPrimitiveState
//go:noescape
func GPUPrimitiveStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUPrimitiveState
//go:noescape
func GPUPrimitiveStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GPUStencilOperation
//go:noescape
func ConstOfGPUStencilOperation(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPUStencilFaceState
//go:noescape
func GPUStencilFaceStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUStencilFaceState
//go:noescape
func GPUStencilFaceStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUDepthStencilState
//go:noescape
func GPUDepthStencilStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUDepthStencilState
//go:noescape
func GPUDepthStencilStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUMultisampleState
//go:noescape
func GPUMultisampleStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUMultisampleState
//go:noescape
func GPUMultisampleStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GPUBlendOperation
//go:noescape
func ConstOfGPUBlendOperation(str js.Ref) uint32

//go:wasmimport plat/js/web constof_GPUBlendFactor
//go:noescape
func ConstOfGPUBlendFactor(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPUBlendComponent
//go:noescape
func GPUBlendComponentJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUBlendComponent
//go:noescape
func GPUBlendComponentJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUBlendState
//go:noescape
func GPUBlendStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUBlendState
//go:noescape
func GPUBlendStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUColorTargetState
//go:noescape
func GPUColorTargetStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUColorTargetState
//go:noescape
func GPUColorTargetStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUFragmentState
//go:noescape
func GPUFragmentStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUFragmentState
//go:noescape
func GPUFragmentStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPURenderPipelineDescriptor
//go:noescape
func GPURenderPipelineDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPURenderPipelineDescriptor
//go:noescape
func GPURenderPipelineDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUColorDict
//go:noescape
func GPUColorDictJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUColorDict
//go:noescape
func GPUColorDictJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPURenderBundle_Label
//go:noescape
func GetGPURenderBundleLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPURenderBundle_Label
//go:noescape
func SetGPURenderBundleLabel(
	this js.Ref,
	val js.Ref,
) js.Ref
