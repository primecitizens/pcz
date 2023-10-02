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

//go:wasmimport plat/js/web store_VideoEncoderEncodeOptionsForAv1
//go:noescape

func VideoEncoderEncodeOptionsForAv1JSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoEncoderEncodeOptionsForAv1
//go:noescape

func VideoEncoderEncodeOptionsForAv1JSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_VideoEncoderEncodeOptionsForAvc
//go:noescape

func VideoEncoderEncodeOptionsForAvcJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoEncoderEncodeOptionsForAvc
//go:noescape

func VideoEncoderEncodeOptionsForAvcJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_VideoEncoderEncodeOptions
//go:noescape

func VideoEncoderEncodeOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoEncoderEncodeOptions
//go:noescape

func VideoEncoderEncodeOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_VideoEncoderSupport
//go:noescape

func VideoEncoderSupportJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoEncoderSupport
//go:noescape

func VideoEncoderSupportJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_VideoEncoder_VideoEncoder
//go:noescape

func NewVideoEncoderByVideoEncoder(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_VideoEncoder_State
//go:noescape

func GetVideoEncoderState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_VideoEncoder_EncodeQueueSize
//go:noescape

func GetVideoEncoderEncodeQueueSize(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_VideoEncoder_Configure
//go:noescape

func CallVideoEncoderConfigure(
	this js.Ref,
	ptr unsafe.Pointer,

	config unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Configure
//go:noescape

func VideoEncoderConfigureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoEncoder_Encode
//go:noescape

func CallVideoEncoderEncode(
	this js.Ref,
	ptr unsafe.Pointer,

	frame js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Encode
//go:noescape

func VideoEncoderEncodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoEncoder_Encode1
//go:noescape

func CallVideoEncoderEncode1(
	this js.Ref,
	ptr unsafe.Pointer,

	frame js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Encode1
//go:noescape

func VideoEncoderEncode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoEncoder_Flush
//go:noescape

func CallVideoEncoderFlush(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Flush
//go:noescape

func VideoEncoderFlushFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoEncoder_Reset
//go:noescape

func CallVideoEncoderReset(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Reset
//go:noescape

func VideoEncoderResetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoEncoder_Close
//go:noescape

func CallVideoEncoderClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Close
//go:noescape

func VideoEncoderCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoEncoder_IsConfigSupported
//go:noescape

func CallVideoEncoderIsConfigSupported(
	this js.Ref,
	ptr unsafe.Pointer,

	config unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_IsConfigSupported
//go:noescape

func VideoEncoderIsConfigSupportedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_VideoFacingModeEnum
//go:noescape

func ConstOfVideoFacingModeEnum(str js.Ref) uint32

//go:wasmimport plat/js/web constof_VideoResizeModeEnum
//go:noescape

func ConstOfVideoResizeModeEnum(str js.Ref) uint32

//go:wasmimport plat/js/web get_VideoTrackGenerator_Writable
//go:noescape

func GetVideoTrackGeneratorWritable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_VideoTrackGenerator_Muted
//go:noescape

func GetVideoTrackGeneratorMuted(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_VideoTrackGenerator_Muted
//go:noescape

func SetVideoTrackGeneratorMuted(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_VideoTrackGenerator_Track
//go:noescape

func GetVideoTrackGeneratorTrack(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_ViewTimelineOptions
//go:noescape

func ViewTimelineOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ViewTimelineOptions
//go:noescape

func ViewTimelineOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ViewTimeline_ViewTimeline
//go:noescape

func NewViewTimelineByViewTimeline(
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ViewTimeline_ViewTimeline1
//go:noescape

func NewViewTimelineByViewTimeline1() js.Ref

//go:wasmimport plat/js/web get_ViewTimeline_Subject
//go:noescape

func GetViewTimelineSubject(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ViewTimeline_StartOffset
//go:noescape

func GetViewTimelineStartOffset(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ViewTimeline_EndOffset
//go:noescape

func GetViewTimelineEndOffset(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_VisibilityStateEntry_Name
//go:noescape

func GetVisibilityStateEntryName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_VisibilityStateEntry_EntryType
//go:noescape

func GetVisibilityStateEntryEntryType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_VisibilityStateEntry_StartTime
//go:noescape

func GetVisibilityStateEntryStartTime(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_VisibilityStateEntry_Duration
//go:noescape

func GetVisibilityStateEntryDuration(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_WEBGL_compressed_texture_astc_GetSupportedProfiles
//go:noescape

func CallWEBGL_compressed_texture_astcGetSupportedProfiles(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WEBGL_compressed_texture_astc_GetSupportedProfiles
//go:noescape

func WEBGL_compressed_texture_astcGetSupportedProfilesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_debug_shaders_GetTranslatedShaderSource
//go:noescape

func CallWEBGL_debug_shadersGetTranslatedShaderSource(
	this js.Ref,
	ptr unsafe.Pointer,

	shader js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WEBGL_debug_shaders_GetTranslatedShaderSource
//go:noescape

func WEBGL_debug_shadersGetTranslatedShaderSourceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_draw_buffers_DrawBuffersWEBGL
//go:noescape

func CallWEBGL_draw_buffersDrawBuffersWEBGL(
	this js.Ref,
	ptr unsafe.Pointer,

	buffers js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WEBGL_draw_buffers_DrawBuffersWEBGL
//go:noescape

func WEBGL_draw_buffersDrawBuffersWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_draw_instanced_base_vertex_base_instance_DrawArraysInstancedBaseInstanceWEBGL
//go:noescape

func CallWEBGL_draw_instanced_base_vertex_base_instanceDrawArraysInstancedBaseInstanceWEBGL(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	first int32,
	count int32,
	instanceCount int32,
	baseInstance uint32,
) js.Ref

//go:wasmimport plat/js/web func_WEBGL_draw_instanced_base_vertex_base_instance_DrawArraysInstancedBaseInstanceWEBGL
//go:noescape

func WEBGL_draw_instanced_base_vertex_base_instanceDrawArraysInstancedBaseInstanceWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_draw_instanced_base_vertex_base_instance_DrawElementsInstancedBaseVertexBaseInstanceWEBGL
//go:noescape

func CallWEBGL_draw_instanced_base_vertex_base_instanceDrawElementsInstancedBaseVertexBaseInstanceWEBGL(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	count int32,
	typ uint32,
	offset float64,
	instanceCount int32,
	baseVertex int32,
	baseInstance uint32,
) js.Ref

//go:wasmimport plat/js/web func_WEBGL_draw_instanced_base_vertex_base_instance_DrawElementsInstancedBaseVertexBaseInstanceWEBGL
//go:noescape

func WEBGL_draw_instanced_base_vertex_base_instanceDrawElementsInstancedBaseVertexBaseInstanceWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_lose_context_LoseContext
//go:noescape

func CallWEBGL_lose_contextLoseContext(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WEBGL_lose_context_LoseContext
//go:noescape

func WEBGL_lose_contextLoseContextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_lose_context_RestoreContext
//go:noescape

func CallWEBGL_lose_contextRestoreContext(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WEBGL_lose_context_RestoreContext
//go:noescape

func WEBGL_lose_contextRestoreContextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_multi_draw_MultiDrawArraysWEBGL
//go:noescape

func CallWEBGL_multi_drawMultiDrawArraysWEBGL(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	firstsList js.Ref,
	firstsOffset uint32,
	countsList js.Ref,
	countsOffset uint32,
	drawcount int32,
) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_MultiDrawArraysWEBGL
//go:noescape

func WEBGL_multi_drawMultiDrawArraysWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_multi_draw_MultiDrawElementsWEBGL
//go:noescape

func CallWEBGL_multi_drawMultiDrawElementsWEBGL(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	countsList js.Ref,
	countsOffset uint32,
	typ uint32,
	offsetsList js.Ref,
	offsetsOffset uint32,
	drawcount int32,
) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_MultiDrawElementsWEBGL
//go:noescape

func WEBGL_multi_drawMultiDrawElementsWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_multi_draw_MultiDrawArraysInstancedWEBGL
//go:noescape

func CallWEBGL_multi_drawMultiDrawArraysInstancedWEBGL(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	firstsList js.Ref,
	firstsOffset uint32,
	countsList js.Ref,
	countsOffset uint32,
	instanceCountsList js.Ref,
	instanceCountsOffset uint32,
	drawcount int32,
) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_MultiDrawArraysInstancedWEBGL
//go:noescape

func WEBGL_multi_drawMultiDrawArraysInstancedWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_multi_draw_MultiDrawElementsInstancedWEBGL
//go:noescape

func CallWEBGL_multi_drawMultiDrawElementsInstancedWEBGL(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	countsList js.Ref,
	countsOffset uint32,
	typ uint32,
	offsetsList js.Ref,
	offsetsOffset uint32,
	instanceCountsList js.Ref,
	instanceCountsOffset uint32,
	drawcount int32,
) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_MultiDrawElementsInstancedWEBGL
//go:noescape

func WEBGL_multi_drawMultiDrawElementsInstancedWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_multi_draw_instanced_base_vertex_base_instance_MultiDrawArraysInstancedBaseInstanceWEBGL
//go:noescape

func CallWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawArraysInstancedBaseInstanceWEBGL(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	firstsList js.Ref,
	firstsOffset uint32,
	countsList js.Ref,
	countsOffset uint32,
	instanceCountsList js.Ref,
	instanceCountsOffset uint32,
	baseInstancesList js.Ref,
	baseInstancesOffset uint32,
	drawcount int32,
) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_instanced_base_vertex_base_instance_MultiDrawArraysInstancedBaseInstanceWEBGL
//go:noescape

func WEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawArraysInstancedBaseInstanceWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_multi_draw_instanced_base_vertex_base_instance_MultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL
//go:noescape

func CallWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	countsList js.Ref,
	countsOffset uint32,
	typ uint32,
	offsetsList js.Ref,
	offsetsOffset uint32,
	instanceCountsList js.Ref,
	instanceCountsOffset uint32,
	baseVerticesList js.Ref,
	baseVerticesOffset uint32,
	baseInstancesList js.Ref,
	baseInstancesOffset uint32,
	drawcount int32,
) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_instanced_base_vertex_base_instance_MultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL
//go:noescape

func WEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_provoking_vertex_ProvokingVertexWEBGL
//go:noescape

func CallWEBGL_provoking_vertexProvokingVertexWEBGL(
	this js.Ref,
	ptr unsafe.Pointer,

	provokeMode uint32,
) js.Ref

//go:wasmimport plat/js/web func_WEBGL_provoking_vertex_ProvokingVertexWEBGL
//go:noescape

func WEBGL_provoking_vertexProvokingVertexWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_WebAssemblyInstantiatedSource
//go:noescape

func WebAssemblyInstantiatedSourceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WebAssemblyInstantiatedSource
//go:noescape

func WebAssemblyInstantiatedSourceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebAssembly_Validate
//go:noescape

func CallWebAssemblyValidate(
	ptr unsafe.Pointer,

	bytes js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Validate
//go:noescape

func WebAssemblyValidateFunc() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_Compile
//go:noescape

func CallWebAssemblyCompile(
	ptr unsafe.Pointer,

	bytes js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Compile
//go:noescape

func WebAssemblyCompileFunc() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_Instantiate
//go:noescape

func CallWebAssemblyInstantiate(
	ptr unsafe.Pointer,

	bytes js.Ref,
	importObject js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Instantiate
//go:noescape

func WebAssemblyInstantiateFunc() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_Instantiate1
//go:noescape

func CallWebAssemblyInstantiate1(
	ptr unsafe.Pointer,

	bytes js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Instantiate1
//go:noescape

func WebAssemblyInstantiate1Func() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_Instantiate2
//go:noescape

func CallWebAssemblyInstantiate2(
	ptr unsafe.Pointer,

	moduleObject js.Ref,
	importObject js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Instantiate2
//go:noescape

func WebAssemblyInstantiate2Func() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_Instantiate3
//go:noescape

func CallWebAssemblyInstantiate3(
	ptr unsafe.Pointer,

	moduleObject js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Instantiate3
//go:noescape

func WebAssemblyInstantiate3Func() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_CompileStreaming
//go:noescape

func CallWebAssemblyCompileStreaming(
	ptr unsafe.Pointer,

	source js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebAssembly_CompileStreaming
//go:noescape

func WebAssemblyCompileStreamingFunc() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_InstantiateStreaming
//go:noescape

func CallWebAssemblyInstantiateStreaming(
	ptr unsafe.Pointer,

	source js.Ref,
	importObject js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebAssembly_InstantiateStreaming
//go:noescape

func WebAssemblyInstantiateStreamingFunc() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_InstantiateStreaming1
//go:noescape

func CallWebAssemblyInstantiateStreaming1(
	ptr unsafe.Pointer,

	source js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebAssembly_InstantiateStreaming1
//go:noescape

func WebAssemblyInstantiateStreaming1Func() js.Ref

//go:wasmimport plat/js/web store_WebGLContextEventInit
//go:noescape

func WebGLContextEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WebGLContextEventInit
//go:noescape

func WebGLContextEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
