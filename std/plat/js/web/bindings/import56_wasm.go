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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VideoEncoder_EncodeQueueSize
//go:noescape
func GetVideoEncoderEncodeQueueSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_VideoEncoder_Configure
//go:noescape
func HasVideoEncoderConfigure(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Configure
//go:noescape
func VideoEncoderConfigureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoEncoder_Configure
//go:noescape
func CallVideoEncoderConfigure(
	this js.Ref, retPtr unsafe.Pointer,
	config unsafe.Pointer)

//go:wasmimport plat/js/web try_VideoEncoder_Configure
//go:noescape
func TryVideoEncoderConfigure(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	config unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_VideoEncoder_Encode
//go:noescape
func HasVideoEncoderEncode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Encode
//go:noescape
func VideoEncoderEncodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoEncoder_Encode
//go:noescape
func CallVideoEncoderEncode(
	this js.Ref, retPtr unsafe.Pointer,
	frame js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_VideoEncoder_Encode
//go:noescape
func TryVideoEncoderEncode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	frame js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_VideoEncoder_Encode1
//go:noescape
func HasVideoEncoderEncode1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Encode1
//go:noescape
func VideoEncoderEncode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoEncoder_Encode1
//go:noescape
func CallVideoEncoderEncode1(
	this js.Ref, retPtr unsafe.Pointer,
	frame js.Ref)

//go:wasmimport plat/js/web try_VideoEncoder_Encode1
//go:noescape
func TryVideoEncoderEncode1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	frame js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_VideoEncoder_Flush
//go:noescape
func HasVideoEncoderFlush(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Flush
//go:noescape
func VideoEncoderFlushFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoEncoder_Flush
//go:noescape
func CallVideoEncoderFlush(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_VideoEncoder_Flush
//go:noescape
func TryVideoEncoderFlush(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_VideoEncoder_Reset
//go:noescape
func HasVideoEncoderReset(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Reset
//go:noescape
func VideoEncoderResetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoEncoder_Reset
//go:noescape
func CallVideoEncoderReset(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_VideoEncoder_Reset
//go:noescape
func TryVideoEncoderReset(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_VideoEncoder_Close
//go:noescape
func HasVideoEncoderClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Close
//go:noescape
func VideoEncoderCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoEncoder_Close
//go:noescape
func CallVideoEncoderClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_VideoEncoder_Close
//go:noescape
func TryVideoEncoderClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_VideoEncoder_IsConfigSupported
//go:noescape
func HasVideoEncoderIsConfigSupported(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_IsConfigSupported
//go:noescape
func VideoEncoderIsConfigSupportedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoEncoder_IsConfigSupported
//go:noescape
func CallVideoEncoderIsConfigSupported(
	this js.Ref, retPtr unsafe.Pointer,
	config unsafe.Pointer)

//go:wasmimport plat/js/web try_VideoEncoder_IsConfigSupported
//go:noescape
func TryVideoEncoderIsConfigSupported(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	config unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_VideoFacingModeEnum
//go:noescape
func ConstOfVideoFacingModeEnum(str js.Ref) uint32

//go:wasmimport plat/js/web constof_VideoResizeModeEnum
//go:noescape
func ConstOfVideoResizeModeEnum(str js.Ref) uint32

//go:wasmimport plat/js/web get_VideoTrackGenerator_Writable
//go:noescape
func GetVideoTrackGeneratorWritable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VideoTrackGenerator_Muted
//go:noescape
func GetVideoTrackGeneratorMuted(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VideoTrackGenerator_Muted
//go:noescape
func SetVideoTrackGeneratorMuted(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_VideoTrackGenerator_Track
//go:noescape
func GetVideoTrackGeneratorTrack(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ViewTimeline_StartOffset
//go:noescape
func GetViewTimelineStartOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ViewTimeline_EndOffset
//go:noescape
func GetViewTimelineEndOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VisibilityStateEntry_Name
//go:noescape
func GetVisibilityStateEntryName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VisibilityStateEntry_EntryType
//go:noescape
func GetVisibilityStateEntryEntryType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VisibilityStateEntry_StartTime
//go:noescape
func GetVisibilityStateEntryStartTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VisibilityStateEntry_Duration
//go:noescape
func GetVisibilityStateEntryDuration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WEBGL_compressed_texture_astc_GetSupportedProfiles
//go:noescape
func HasWEBGL_compressed_texture_astcGetSupportedProfiles(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_compressed_texture_astc_GetSupportedProfiles
//go:noescape
func WEBGL_compressed_texture_astcGetSupportedProfilesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_compressed_texture_astc_GetSupportedProfiles
//go:noescape
func CallWEBGL_compressed_texture_astcGetSupportedProfiles(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WEBGL_compressed_texture_astc_GetSupportedProfiles
//go:noescape
func TryWEBGL_compressed_texture_astcGetSupportedProfiles(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WEBGL_debug_shaders_GetTranslatedShaderSource
//go:noescape
func HasWEBGL_debug_shadersGetTranslatedShaderSource(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_debug_shaders_GetTranslatedShaderSource
//go:noescape
func WEBGL_debug_shadersGetTranslatedShaderSourceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_debug_shaders_GetTranslatedShaderSource
//go:noescape
func CallWEBGL_debug_shadersGetTranslatedShaderSource(
	this js.Ref, retPtr unsafe.Pointer,
	shader js.Ref)

//go:wasmimport plat/js/web try_WEBGL_debug_shaders_GetTranslatedShaderSource
//go:noescape
func TryWEBGL_debug_shadersGetTranslatedShaderSource(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	shader js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WEBGL_draw_buffers_DrawBuffersWEBGL
//go:noescape
func HasWEBGL_draw_buffersDrawBuffersWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_draw_buffers_DrawBuffersWEBGL
//go:noescape
func WEBGL_draw_buffersDrawBuffersWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_draw_buffers_DrawBuffersWEBGL
//go:noescape
func CallWEBGL_draw_buffersDrawBuffersWEBGL(
	this js.Ref, retPtr unsafe.Pointer,
	buffers js.Ref)

//go:wasmimport plat/js/web try_WEBGL_draw_buffers_DrawBuffersWEBGL
//go:noescape
func TryWEBGL_draw_buffersDrawBuffersWEBGL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffers js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WEBGL_draw_instanced_base_vertex_base_instance_DrawArraysInstancedBaseInstanceWEBGL
//go:noescape
func HasWEBGL_draw_instanced_base_vertex_base_instanceDrawArraysInstancedBaseInstanceWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_draw_instanced_base_vertex_base_instance_DrawArraysInstancedBaseInstanceWEBGL
//go:noescape
func WEBGL_draw_instanced_base_vertex_base_instanceDrawArraysInstancedBaseInstanceWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_draw_instanced_base_vertex_base_instance_DrawArraysInstancedBaseInstanceWEBGL
//go:noescape
func CallWEBGL_draw_instanced_base_vertex_base_instanceDrawArraysInstancedBaseInstanceWEBGL(
	this js.Ref, retPtr unsafe.Pointer,
	mode uint32,
	first int32,
	count int32,
	instanceCount int32,
	baseInstance uint32)

//go:wasmimport plat/js/web try_WEBGL_draw_instanced_base_vertex_base_instance_DrawArraysInstancedBaseInstanceWEBGL
//go:noescape
func TryWEBGL_draw_instanced_base_vertex_base_instanceDrawArraysInstancedBaseInstanceWEBGL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mode uint32,
	first int32,
	count int32,
	instanceCount int32,
	baseInstance uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_WEBGL_draw_instanced_base_vertex_base_instance_DrawElementsInstancedBaseVertexBaseInstanceWEBGL
//go:noescape
func HasWEBGL_draw_instanced_base_vertex_base_instanceDrawElementsInstancedBaseVertexBaseInstanceWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_draw_instanced_base_vertex_base_instance_DrawElementsInstancedBaseVertexBaseInstanceWEBGL
//go:noescape
func WEBGL_draw_instanced_base_vertex_base_instanceDrawElementsInstancedBaseVertexBaseInstanceWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_draw_instanced_base_vertex_base_instance_DrawElementsInstancedBaseVertexBaseInstanceWEBGL
//go:noescape
func CallWEBGL_draw_instanced_base_vertex_base_instanceDrawElementsInstancedBaseVertexBaseInstanceWEBGL(
	this js.Ref, retPtr unsafe.Pointer,
	mode uint32,
	count int32,
	typ uint32,
	offset float64,
	instanceCount int32,
	baseVertex int32,
	baseInstance uint32)

//go:wasmimport plat/js/web try_WEBGL_draw_instanced_base_vertex_base_instance_DrawElementsInstancedBaseVertexBaseInstanceWEBGL
//go:noescape
func TryWEBGL_draw_instanced_base_vertex_base_instanceDrawElementsInstancedBaseVertexBaseInstanceWEBGL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mode uint32,
	count int32,
	typ uint32,
	offset float64,
	instanceCount int32,
	baseVertex int32,
	baseInstance uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_WEBGL_lose_context_LoseContext
//go:noescape
func HasWEBGL_lose_contextLoseContext(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_lose_context_LoseContext
//go:noescape
func WEBGL_lose_contextLoseContextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_lose_context_LoseContext
//go:noescape
func CallWEBGL_lose_contextLoseContext(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WEBGL_lose_context_LoseContext
//go:noescape
func TryWEBGL_lose_contextLoseContext(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WEBGL_lose_context_RestoreContext
//go:noescape
func HasWEBGL_lose_contextRestoreContext(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_lose_context_RestoreContext
//go:noescape
func WEBGL_lose_contextRestoreContextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_lose_context_RestoreContext
//go:noescape
func CallWEBGL_lose_contextRestoreContext(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WEBGL_lose_context_RestoreContext
//go:noescape
func TryWEBGL_lose_contextRestoreContext(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WEBGL_multi_draw_MultiDrawArraysWEBGL
//go:noescape
func HasWEBGL_multi_drawMultiDrawArraysWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_MultiDrawArraysWEBGL
//go:noescape
func WEBGL_multi_drawMultiDrawArraysWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_multi_draw_MultiDrawArraysWEBGL
//go:noescape
func CallWEBGL_multi_drawMultiDrawArraysWEBGL(
	this js.Ref, retPtr unsafe.Pointer,
	mode uint32,
	firstsList js.Ref,
	firstsOffset uint32,
	countsList js.Ref,
	countsOffset uint32,
	drawcount int32)

//go:wasmimport plat/js/web try_WEBGL_multi_draw_MultiDrawArraysWEBGL
//go:noescape
func TryWEBGL_multi_drawMultiDrawArraysWEBGL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mode uint32,
	firstsList js.Ref,
	firstsOffset uint32,
	countsList js.Ref,
	countsOffset uint32,
	drawcount int32) (ok js.Ref)

//go:wasmimport plat/js/web has_WEBGL_multi_draw_MultiDrawElementsWEBGL
//go:noescape
func HasWEBGL_multi_drawMultiDrawElementsWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_MultiDrawElementsWEBGL
//go:noescape
func WEBGL_multi_drawMultiDrawElementsWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_multi_draw_MultiDrawElementsWEBGL
//go:noescape
func CallWEBGL_multi_drawMultiDrawElementsWEBGL(
	this js.Ref, retPtr unsafe.Pointer,
	mode uint32,
	countsList js.Ref,
	countsOffset uint32,
	typ uint32,
	offsetsList js.Ref,
	offsetsOffset uint32,
	drawcount int32)

//go:wasmimport plat/js/web try_WEBGL_multi_draw_MultiDrawElementsWEBGL
//go:noescape
func TryWEBGL_multi_drawMultiDrawElementsWEBGL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mode uint32,
	countsList js.Ref,
	countsOffset uint32,
	typ uint32,
	offsetsList js.Ref,
	offsetsOffset uint32,
	drawcount int32) (ok js.Ref)

//go:wasmimport plat/js/web has_WEBGL_multi_draw_MultiDrawArraysInstancedWEBGL
//go:noescape
func HasWEBGL_multi_drawMultiDrawArraysInstancedWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_MultiDrawArraysInstancedWEBGL
//go:noescape
func WEBGL_multi_drawMultiDrawArraysInstancedWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_multi_draw_MultiDrawArraysInstancedWEBGL
//go:noescape
func CallWEBGL_multi_drawMultiDrawArraysInstancedWEBGL(
	this js.Ref, retPtr unsafe.Pointer,
	mode uint32,
	firstsList js.Ref,
	firstsOffset uint32,
	countsList js.Ref,
	countsOffset uint32,
	instanceCountsList js.Ref,
	instanceCountsOffset uint32,
	drawcount int32)

//go:wasmimport plat/js/web try_WEBGL_multi_draw_MultiDrawArraysInstancedWEBGL
//go:noescape
func TryWEBGL_multi_drawMultiDrawArraysInstancedWEBGL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mode uint32,
	firstsList js.Ref,
	firstsOffset uint32,
	countsList js.Ref,
	countsOffset uint32,
	instanceCountsList js.Ref,
	instanceCountsOffset uint32,
	drawcount int32) (ok js.Ref)

//go:wasmimport plat/js/web has_WEBGL_multi_draw_MultiDrawElementsInstancedWEBGL
//go:noescape
func HasWEBGL_multi_drawMultiDrawElementsInstancedWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_MultiDrawElementsInstancedWEBGL
//go:noescape
func WEBGL_multi_drawMultiDrawElementsInstancedWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_multi_draw_MultiDrawElementsInstancedWEBGL
//go:noescape
func CallWEBGL_multi_drawMultiDrawElementsInstancedWEBGL(
	this js.Ref, retPtr unsafe.Pointer,
	mode uint32,
	countsList js.Ref,
	countsOffset uint32,
	typ uint32,
	offsetsList js.Ref,
	offsetsOffset uint32,
	instanceCountsList js.Ref,
	instanceCountsOffset uint32,
	drawcount int32)

//go:wasmimport plat/js/web try_WEBGL_multi_draw_MultiDrawElementsInstancedWEBGL
//go:noescape
func TryWEBGL_multi_drawMultiDrawElementsInstancedWEBGL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mode uint32,
	countsList js.Ref,
	countsOffset uint32,
	typ uint32,
	offsetsList js.Ref,
	offsetsOffset uint32,
	instanceCountsList js.Ref,
	instanceCountsOffset uint32,
	drawcount int32) (ok js.Ref)

//go:wasmimport plat/js/web has_WEBGL_multi_draw_instanced_base_vertex_base_instance_MultiDrawArraysInstancedBaseInstanceWEBGL
//go:noescape
func HasWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawArraysInstancedBaseInstanceWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_instanced_base_vertex_base_instance_MultiDrawArraysInstancedBaseInstanceWEBGL
//go:noescape
func WEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawArraysInstancedBaseInstanceWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_multi_draw_instanced_base_vertex_base_instance_MultiDrawArraysInstancedBaseInstanceWEBGL
//go:noescape
func CallWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawArraysInstancedBaseInstanceWEBGL(
	this js.Ref, retPtr unsafe.Pointer,
	mode uint32,
	firstsList js.Ref,
	firstsOffset uint32,
	countsList js.Ref,
	countsOffset uint32,
	instanceCountsList js.Ref,
	instanceCountsOffset uint32,
	baseInstancesList js.Ref,
	baseInstancesOffset uint32,
	drawcount int32)

//go:wasmimport plat/js/web try_WEBGL_multi_draw_instanced_base_vertex_base_instance_MultiDrawArraysInstancedBaseInstanceWEBGL
//go:noescape
func TryWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawArraysInstancedBaseInstanceWEBGL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mode uint32,
	firstsList js.Ref,
	firstsOffset uint32,
	countsList js.Ref,
	countsOffset uint32,
	instanceCountsList js.Ref,
	instanceCountsOffset uint32,
	baseInstancesList js.Ref,
	baseInstancesOffset uint32,
	drawcount int32) (ok js.Ref)

//go:wasmimport plat/js/web has_WEBGL_multi_draw_instanced_base_vertex_base_instance_MultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL
//go:noescape
func HasWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_instanced_base_vertex_base_instance_MultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL
//go:noescape
func WEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_multi_draw_instanced_base_vertex_base_instance_MultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL
//go:noescape
func CallWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL(
	this js.Ref, retPtr unsafe.Pointer,
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
	drawcount int32)

//go:wasmimport plat/js/web try_WEBGL_multi_draw_instanced_base_vertex_base_instance_MultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL
//go:noescape
func TryWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
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
	drawcount int32) (ok js.Ref)

//go:wasmimport plat/js/web has_WEBGL_provoking_vertex_ProvokingVertexWEBGL
//go:noescape
func HasWEBGL_provoking_vertexProvokingVertexWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_provoking_vertex_ProvokingVertexWEBGL
//go:noescape
func WEBGL_provoking_vertexProvokingVertexWEBGLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WEBGL_provoking_vertex_ProvokingVertexWEBGL
//go:noescape
func CallWEBGL_provoking_vertexProvokingVertexWEBGL(
	this js.Ref, retPtr unsafe.Pointer,
	provokeMode uint32)

//go:wasmimport plat/js/web try_WEBGL_provoking_vertex_ProvokingVertexWEBGL
//go:noescape
func TryWEBGL_provoking_vertexProvokingVertexWEBGL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	provokeMode uint32) (ok js.Ref)

//go:wasmimport plat/js/web store_WebAssemblyInstantiatedSource
//go:noescape
func WebAssemblyInstantiatedSourceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WebAssemblyInstantiatedSource
//go:noescape
func WebAssemblyInstantiatedSourceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_WebAssembly_Validate
//go:noescape
func HasWebAssemblyValidate() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Validate
//go:noescape
func WebAssemblyValidateFunc() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_Validate
//go:noescape
func CallWebAssemblyValidate(
	retPtr unsafe.Pointer,
	bytes js.Ref)

//go:wasmimport plat/js/web try_WebAssembly_Validate
//go:noescape
func TryWebAssemblyValidate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bytes js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WebAssembly_Compile
//go:noescape
func HasWebAssemblyCompile() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Compile
//go:noescape
func WebAssemblyCompileFunc() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_Compile
//go:noescape
func CallWebAssemblyCompile(
	retPtr unsafe.Pointer,
	bytes js.Ref)

//go:wasmimport plat/js/web try_WebAssembly_Compile
//go:noescape
func TryWebAssemblyCompile(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bytes js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WebAssembly_Instantiate
//go:noescape
func HasWebAssemblyInstantiate() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Instantiate
//go:noescape
func WebAssemblyInstantiateFunc() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_Instantiate
//go:noescape
func CallWebAssemblyInstantiate(
	retPtr unsafe.Pointer,
	bytes js.Ref,
	importObject js.Ref)

//go:wasmimport plat/js/web try_WebAssembly_Instantiate
//go:noescape
func TryWebAssemblyInstantiate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bytes js.Ref,
	importObject js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WebAssembly_Instantiate1
//go:noescape
func HasWebAssemblyInstantiate1() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Instantiate1
//go:noescape
func WebAssemblyInstantiate1Func() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_Instantiate1
//go:noescape
func CallWebAssemblyInstantiate1(
	retPtr unsafe.Pointer,
	bytes js.Ref)

//go:wasmimport plat/js/web try_WebAssembly_Instantiate1
//go:noescape
func TryWebAssemblyInstantiate1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bytes js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WebAssembly_Instantiate2
//go:noescape
func HasWebAssemblyInstantiate2() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Instantiate2
//go:noescape
func WebAssemblyInstantiate2Func() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_Instantiate2
//go:noescape
func CallWebAssemblyInstantiate2(
	retPtr unsafe.Pointer,
	moduleObject js.Ref,
	importObject js.Ref)

//go:wasmimport plat/js/web try_WebAssembly_Instantiate2
//go:noescape
func TryWebAssemblyInstantiate2(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	moduleObject js.Ref,
	importObject js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WebAssembly_Instantiate3
//go:noescape
func HasWebAssemblyInstantiate3() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Instantiate3
//go:noescape
func WebAssemblyInstantiate3Func() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_Instantiate3
//go:noescape
func CallWebAssemblyInstantiate3(
	retPtr unsafe.Pointer,
	moduleObject js.Ref)

//go:wasmimport plat/js/web try_WebAssembly_Instantiate3
//go:noescape
func TryWebAssemblyInstantiate3(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	moduleObject js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WebAssembly_CompileStreaming
//go:noescape
func HasWebAssemblyCompileStreaming() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_CompileStreaming
//go:noescape
func WebAssemblyCompileStreamingFunc() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_CompileStreaming
//go:noescape
func CallWebAssemblyCompileStreaming(
	retPtr unsafe.Pointer,
	source js.Ref)

//go:wasmimport plat/js/web try_WebAssembly_CompileStreaming
//go:noescape
func TryWebAssemblyCompileStreaming(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WebAssembly_InstantiateStreaming
//go:noescape
func HasWebAssemblyInstantiateStreaming() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_InstantiateStreaming
//go:noescape
func WebAssemblyInstantiateStreamingFunc() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_InstantiateStreaming
//go:noescape
func CallWebAssemblyInstantiateStreaming(
	retPtr unsafe.Pointer,
	source js.Ref,
	importObject js.Ref)

//go:wasmimport plat/js/web try_WebAssembly_InstantiateStreaming
//go:noescape
func TryWebAssemblyInstantiateStreaming(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source js.Ref,
	importObject js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WebAssembly_InstantiateStreaming1
//go:noescape
func HasWebAssemblyInstantiateStreaming1() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_InstantiateStreaming1
//go:noescape
func WebAssemblyInstantiateStreaming1Func() js.Ref

//go:wasmimport plat/js/web call_WebAssembly_InstantiateStreaming1
//go:noescape
func CallWebAssemblyInstantiateStreaming1(
	retPtr unsafe.Pointer,
	source js.Ref)

//go:wasmimport plat/js/web try_WebAssembly_InstantiateStreaming1
//go:noescape
func TryWebAssemblyInstantiateStreaming1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_WebGLContextEventInit
//go:noescape
func WebGLContextEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WebGLContextEventInit
//go:noescape
func WebGLContextEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
