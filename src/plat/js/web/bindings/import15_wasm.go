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

//go:wasmimport plat/js/web store_GPUExtent3DDict
//go:noescape
func GPUExtent3DDictJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUExtent3DDict
//go:noescape
func GPUExtent3DDictJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUTextureDescriptor
//go:noescape
func GPUTextureDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUTextureDescriptor
//go:noescape
func GPUTextureDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPUSampler_Label
//go:noescape
func GetGPUSamplerLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPUSampler_Label
//go:noescape
func SetGPUSamplerLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web constof_GPUAddressMode
//go:noescape
func ConstOfGPUAddressMode(str js.Ref) uint32

//go:wasmimport plat/js/web constof_GPUFilterMode
//go:noescape
func ConstOfGPUFilterMode(str js.Ref) uint32

//go:wasmimport plat/js/web constof_GPUMipmapFilterMode
//go:noescape
func ConstOfGPUMipmapFilterMode(str js.Ref) uint32

//go:wasmimport plat/js/web constof_GPUCompareFunction
//go:noescape
func ConstOfGPUCompareFunction(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPUSamplerDescriptor
//go:noescape
func GPUSamplerDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUSamplerDescriptor
//go:noescape
func GPUSamplerDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPUExternalTexture_Label
//go:noescape
func GetGPUExternalTextureLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPUExternalTexture_Label
//go:noescape
func SetGPUExternalTextureLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_GPUExternalTextureDescriptor
//go:noescape
func GPUExternalTextureDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUExternalTextureDescriptor
//go:noescape
func GPUExternalTextureDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPUBindGroupLayout_Label
//go:noescape
func GetGPUBindGroupLayoutLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPUBindGroupLayout_Label
//go:noescape
func SetGPUBindGroupLayoutLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web constof_GPUBufferBindingType
//go:noescape
func ConstOfGPUBufferBindingType(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPUBufferBindingLayout
//go:noescape
func GPUBufferBindingLayoutJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUBufferBindingLayout
//go:noescape
func GPUBufferBindingLayoutJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GPUSamplerBindingType
//go:noescape
func ConstOfGPUSamplerBindingType(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPUSamplerBindingLayout
//go:noescape
func GPUSamplerBindingLayoutJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUSamplerBindingLayout
//go:noescape
func GPUSamplerBindingLayoutJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GPUTextureSampleType
//go:noescape
func ConstOfGPUTextureSampleType(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPUTextureBindingLayout
//go:noescape
func GPUTextureBindingLayoutJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUTextureBindingLayout
//go:noescape
func GPUTextureBindingLayoutJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GPUStorageTextureAccess
//go:noescape
func ConstOfGPUStorageTextureAccess(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPUStorageTextureBindingLayout
//go:noescape
func GPUStorageTextureBindingLayoutJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUStorageTextureBindingLayout
//go:noescape
func GPUStorageTextureBindingLayoutJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUExternalTextureBindingLayout
//go:noescape
func GPUExternalTextureBindingLayoutJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUExternalTextureBindingLayout
//go:noescape
func GPUExternalTextureBindingLayoutJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUBindGroupLayoutEntry
//go:noescape
func GPUBindGroupLayoutEntryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUBindGroupLayoutEntry
//go:noescape
func GPUBindGroupLayoutEntryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUBindGroupLayoutDescriptor
//go:noescape
func GPUBindGroupLayoutDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUBindGroupLayoutDescriptor
//go:noescape
func GPUBindGroupLayoutDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPUPipelineLayout_Label
//go:noescape
func GetGPUPipelineLayoutLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPUPipelineLayout_Label
//go:noescape
func SetGPUPipelineLayoutLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_GPUPipelineLayoutDescriptor
//go:noescape
func GPUPipelineLayoutDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUPipelineLayoutDescriptor
//go:noescape
func GPUPipelineLayoutDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPUBindGroup_Label
//go:noescape
func GetGPUBindGroupLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPUBindGroup_Label
//go:noescape
func SetGPUBindGroupLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_GPUBufferBinding
//go:noescape
func GPUBufferBindingJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUBufferBinding
//go:noescape
func GPUBufferBindingJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUBindGroupEntry
//go:noescape
func GPUBindGroupEntryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUBindGroupEntry
//go:noescape
func GPUBindGroupEntryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUBindGroupDescriptor
//go:noescape
func GPUBindGroupDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUBindGroupDescriptor
//go:noescape
func GPUBindGroupDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GPUCompilationMessageType
//go:noescape
func ConstOfGPUCompilationMessageType(str js.Ref) uint32

//go:wasmimport plat/js/web get_GPUCompilationMessage_Message
//go:noescape
func GetGPUCompilationMessageMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUCompilationMessage_Type
//go:noescape
func GetGPUCompilationMessageType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUCompilationMessage_LineNum
//go:noescape
func GetGPUCompilationMessageLineNum(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUCompilationMessage_LinePos
//go:noescape
func GetGPUCompilationMessageLinePos(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUCompilationMessage_Offset
//go:noescape
func GetGPUCompilationMessageOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUCompilationMessage_Length
//go:noescape
func GetGPUCompilationMessageLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUCompilationInfo_Messages
//go:noescape
func GetGPUCompilationInfoMessages(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUShaderModule_Label
//go:noescape
func GetGPUShaderModuleLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_GPUShaderModule_Label
//go:noescape
func SetGPUShaderModuleLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_GPUShaderModule_GetCompilationInfo
//go:noescape
func HasGPUShaderModuleGetCompilationInfo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUShaderModule_GetCompilationInfo
//go:noescape
func GPUShaderModuleGetCompilationInfoFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUShaderModule_GetCompilationInfo
//go:noescape
func CallGPUShaderModuleGetCompilationInfo(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUShaderModule_GetCompilationInfo
//go:noescape
func TryGPUShaderModuleGetCompilationInfo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
